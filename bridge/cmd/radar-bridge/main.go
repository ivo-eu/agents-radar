package main

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"html"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"syscall"
	"time"
)

const (
	defaultPagesURL = "https://ivo-eu.github.io/agents-radar"
	defaultListen   = "127.0.0.1:4399"
	workerCount     = 4
)

var (
	datePattern   = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	reportPattern = regexp.MustCompile(`^[a-z0-9-]+$`)
	spacePattern  = regexp.MustCompile(`\s+`)
	hex64Pattern  = regexp.MustCompile(`^[0-9a-f]{64}$`)
)

type Config struct {
	VaultPath string `json:"vault_path"`
	PagesURL  string `json:"pages_url"`
	Listen    string `json:"listen"`
	// RepoPath is a local clone of the agents-radar repo that Bridge owns and
	// pushes favorites.json from. Optional: if empty, favorites sync is disabled
	// and the endpoint returns a clear error so the website keeps items pending.
	RepoPath string `json:"repo_path,omitempty"`
}

type State struct {
	Token             string            `json:"token"`
	LastSync          string            `json:"last_sync,omitempty"`
	LastError         string            `json:"last_error,omitempty"`
	Files             map[string]string `json:"files"`
	Favorites         map[string]string `json:"favorites"` // legacy (Phase 1 per-file map); unused in Phase 2
	FavoritesLastSync string            `json:"favorites_last_sync,omitempty"`
	FavoritesError    string            `json:"favorites_error,omitempty"`
}

type SyncFile struct {
	Path   string `json:"path"`
	Size   int64  `json:"size"`
	SHA256 string `json:"sha256"`
}

type SyncManifest struct {
	Version   int        `json:"version"`
	Generated string     `json:"generated"`
	Files     []SyncFile `json:"files"`
}

type SyncResult struct {
	Updated int `json:"updated"`
	Skipped int `json:"skipped"`
}

// FavoriteCandidate is the legacy (Phase 1) per-item model. Retained only for
// the cross-language ID contract tests; the live sync path uses FavoriteItem.
type FavoriteCandidate struct {
	Kind       string `json:"kind"`
	Title      string `json:"title"`
	URL        string `json:"url,omitempty"`
	Excerpt    string `json:"excerpt"`
	Section    string `json:"section,omitempty"`
	ReportDate string `json:"report_date"`
	ReportType string `json:"report_type"`
	ReportURL  string `json:"report_url"`
}

const favoritesFileName = "favorites.json"

// FavoriteSource records where a favorite came from (the daily report) and the
// anchor used for jump-back highlighting on the website.
type FavoriteSource struct {
	Date   string `json:"date"`
	Report string `json:"report"`
	Label  string `json:"label,omitempty"`
	URL    string `json:"url,omitempty"`
	Anchor string `json:"anchor,omitempty"`
}

// FavoriteItem is the Phase 2 unified model. The id is computed by the browser
// (SHA-256, see index.html) and trusted here. A non-empty DeletedAt is a
// tombstone so cross-device removals do not resurrect on the next merge.
//
// CreatedAt is the FIRST-favorite time and is used only for stable sorting.
// RevivedAt records the most recent re-favorite (after a delete). Life/death is
// decided by max(CreatedAt, RevivedAt) vs max(DeletedAt) — keeping the revive
// event time means a stale tombstone from another device cannot re-kill an item
// that was re-favorited after that tombstone. (See review §5.2.)
type FavoriteItem struct {
	ID        string         `json:"id"`
	Kind      string         `json:"kind"`
	Title     string         `json:"title,omitempty"`
	URL       string         `json:"url,omitempty"`
	Site      string         `json:"site,omitempty"`
	Excerpt   string         `json:"excerpt,omitempty"`
	Section   string         `json:"section,omitempty"`
	CreatedAt string         `json:"created_at,omitempty"`
	RevivedAt string         `json:"revived_at,omitempty"`
	DeletedAt string         `json:"deleted_at,omitempty"`
	Source    FavoriteSource `json:"source"`
}

// FavoritesFile is the on-disk / on-wire shape of favorites.json.
type FavoritesFile struct {
	Version   int            `json:"version"`
	UpdatedAt string         `json:"updated_at"`
	Items     []FavoriteItem `json:"items"`
}

type App struct {
	config        Config
	configPath    string
	statePath     string
	client        *http.Client
	mu            sync.Mutex
	state         State
	syncing       atomic.Bool
	syncMu        sync.Mutex
	favMu         sync.Mutex
	allowedOrigin string
}

func main() {
	for _, key := range []string{"HTTP_PROXY", "HTTPS_PROXY", "ALL_PROXY", "http_proxy", "https_proxy", "all_proxy"} {
		_ = os.Unsetenv(key)
	}

	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	switch os.Args[1] {
	case "install":
		if err := installCommand(os.Args[2:]); err != nil {
			log.Fatal(err)
		}
	case "serve":
		app, err := loadApp()
		if err != nil {
			log.Fatal(err)
		}
		if err := app.serve(); err != nil {
			log.Fatal(err)
		}
	case "sync":
		app, err := loadApp()
		if err != nil {
			log.Fatal(err)
		}
		if result, handled, err := app.syncViaServer(); handled {
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("updated=%d skipped=%d\n", result.Updated, result.Skipped)
			return
		}
		result, err := app.sync(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("updated=%d skipped=%d\n", result.Updated, result.Skipped)
	case "status":
		app, err := loadApp()
		if err != nil {
			log.Fatal(err)
		}
		data, _ := json.MarshalIndent(app.status(), "", "  ")
		fmt.Println(string(data))
	default:
		usage()
		os.Exit(2)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: radar-bridge <install|serve|sync|status>")
}

func appSupportDir() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, "Library", "Application Support", "RadarBridge"), nil
}

func loadApp() (*App, error) {
	dir, err := appSupportDir()
	if err != nil {
		return nil, err
	}
	configPath := filepath.Join(dir, "config.json")
	statePath := filepath.Join(dir, "state.json")

	var config Config
	if err := readJSON(configPath, &config); err != nil {
		return nil, fmt.Errorf("read config: %w; run radar-bridge install first", err)
	}
	config.PagesURL = strings.TrimRight(config.PagesURL, "/")
	if config.Listen == "" {
		config.Listen = defaultListen
	}
	if config.VaultPath == "" || config.PagesURL == "" {
		return nil, errors.New("config requires vault_path and pages_url")
	}
	if config.RepoPath != "" {
		config.RepoPath = filepath.Clean(config.RepoPath)
	}
	pagesURL, err := url.Parse(config.PagesURL)
	if err != nil || pagesURL.Scheme != "https" || pagesURL.Host == "" {
		return nil, errors.New("pages_url must be an absolute HTTPS URL")
	}
	allowedOrigin := pagesURL.Scheme + "://" + pagesURL.Host

	state := State{Files: map[string]string{}, Favorites: map[string]string{}}
	if err := readJSON(statePath, &state); err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("read state: %w", err)
	}
	if state.Token == "" {
		state.Token, err = randomToken()
		if err != nil {
			return nil, err
		}
	}
	if state.Files == nil {
		state.Files = map[string]string{}
	}
	if state.Favorites == nil {
		state.Favorites = map[string]string{}
	}

	app := &App{
		config:        config,
		configPath:    configPath,
		statePath:     statePath,
		state:         state,
		allowedOrigin: allowedOrigin,
		client: &http.Client{
			Timeout: 90 * time.Second,
			Transport: &http.Transport{
				Proxy:               nil,
				MaxIdleConns:        8,
				MaxIdleConnsPerHost: 4,
				IdleConnTimeout:     30 * time.Second,
			},
		},
	}
	if err := app.saveState(); err != nil {
		return nil, err
	}
	return app, nil
}

func installCommand(args []string) error {
	flags := flag.NewFlagSet("install", flag.ContinueOnError)
	vault := flags.String("vault", "", "absolute SecondBrain vault path")
	pagesURL := flags.String("pages-url", defaultPagesURL, "agents-radar Pages URL")
	listen := flags.String("listen", defaultListen, "local listen address")
	repoPath := flags.String("repo", "", "absolute path to an existing agents-radar clone Bridge may push from")
	repoURL := flags.String("repo-url", "", "git URL to clone into a dedicated Bridge-owned agents-radar clone")
	if err := flags.Parse(args); err != nil {
		return err
	}
	if *vault == "" || !filepath.IsAbs(*vault) {
		return errors.New("--vault must be an absolute path")
	}
	info, err := os.Stat(*vault)
	if err != nil || !info.IsDir() {
		return fmt.Errorf("vault directory does not exist: %s", *vault)
	}
	host, _, err := net.SplitHostPort(*listen)
	if err != nil || host != "127.0.0.1" {
		return errors.New("--listen must use 127.0.0.1")
	}

	dir, err := appSupportDir()
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Join(dir, "logs"), 0o700); err != nil {
		return err
	}

	resolvedRepo, err := resolveRepoPath(*repoPath, *repoURL, dir)
	if err != nil {
		return err
	}

	config := Config{
		VaultPath: filepath.Clean(*vault),
		PagesURL:  strings.TrimRight(*pagesURL, "/"),
		Listen:    *listen,
		RepoPath:  resolvedRepo,
	}
	if err := writeJSONAtomic(filepath.Join(dir, "config.json"), config, 0o600); err != nil {
		return err
	}

	// Deploy the binary to a stable location and point LaunchAgents at THAT copy,
	// not at wherever `install` happened to run from (e.g. a repo build product
	// that later gets cleaned). See review §5.6.
	stableBin, err := deployStableBinary(dir)
	if err != nil {
		return err
	}
	if err := installLaunchAgents(stableBin, dir); err != nil {
		return err
	}
	fmt.Printf("Radar Bridge installed for %s\n", config.VaultPath)
	fmt.Printf("Binary: %s\n", stableBin)
	if config.RepoPath != "" {
		fmt.Printf("Favorites sync repo: %s\n", config.RepoPath)
		if origin, err := runGit(config.RepoPath, "remote", "get-url", "origin"); err == nil {
			fmt.Printf("Favorites push origin: %s\n", strings.TrimSpace(origin))
		}
	} else {
		fmt.Println("No --repo/--repo-url given: favorites sync to GitHub is disabled.")
	}
	return nil
}

// resolveRepoPath validates an existing clone (--repo) or clones one (--repo-url)
// into a dedicated Bridge-owned directory. Returns "" when neither is given
// (favorites sync stays disabled). The resulting path must be a git work tree
// with an "origin" remote so Bridge can push favorites.json.
func resolveRepoPath(repoPath, repoURL, supportDir string) (string, error) {
	if repoPath != "" && repoURL != "" {
		return "", errors.New("use either --repo or --repo-url, not both")
	}
	if repoPath != "" {
		if !filepath.IsAbs(repoPath) {
			return "", errors.New("--repo must be an absolute path")
		}
		clean := filepath.Clean(repoPath)
		if err := assertRepoUsable(clean); err != nil {
			return "", err
		}
		return clean, nil
	}
	if repoURL != "" {
		dest := filepath.Join(supportDir, "agents-radar")
		if _, err := os.Stat(filepath.Join(dest, ".git")); err == nil {
			// Already cloned; reuse it — but make sure its origin matches the URL the
			// user just passed, or we would silently keep pushing to the OLD repo.
			if err := assertRepoUsable(dest); err != nil {
				return "", err
			}
			existing, gerr := runGit(dest, "remote", "get-url", "origin")
			if gerr != nil {
				return "", fmt.Errorf("git remote get-url origin: %w: %s", gerr, strings.TrimSpace(existing))
			}
			if !sameGitURL(existing, repoURL) {
				// Only adopt a new origin on a Bridge-owned clone; never silently
				// rewrite the remote of a directory we cannot confirm we own.
				if !repoIsOwned(dest) {
					return "", fmt.Errorf("%s 已存在但不是 Bridge 专用 clone，且 origin（%s）与 --repo-url（%s）不一致；请手动删除该目录或确认后重试", dest, strings.TrimSpace(existing), repoURL)
				}
				if out, err := runGit(dest, "remote", "set-url", "origin", repoURL); err != nil {
					return "", fmt.Errorf("更新 origin 到 %s 失败：%w: %s", repoURL, err, strings.TrimSpace(out))
				}
				fmt.Printf("已将专用 clone 的 origin 更新为 %s\n", repoURL)
			}
			if out, err := runGit(dest, "fetch", "origin"); err != nil {
				return "", fmt.Errorf("git fetch in existing clone: %w: %s", err, out)
			}
			if err := markRepoOwned(dest); err != nil {
				return "", err
			}
			return dest, nil
		}
		if out, err := runGit("", "clone", "--depth", "1", repoURL, dest); err != nil {
			return "", fmt.Errorf("git clone failed: %w: %s", err, out)
		}
		if err := assertRepoUsable(dest); err != nil {
			return "", err
		}
		// Mark this as a Bridge-owned clone: only owned clones may be hard-reset.
		if err := markRepoOwned(dest); err != nil {
			return "", err
		}
		return dest, nil
	}
	return "", nil
}

// sameGitURL compares two git remote URLs ignoring a trailing ".git", a trailing
// slash, and surrounding whitespace — enough to tell "same repo" from "different
// repo" for the origin-mismatch guard.
func sameGitURL(a, b string) bool {
	norm := func(s string) string {
		s = strings.TrimSpace(s)
		s = strings.TrimSuffix(s, "/")
		s = strings.TrimSuffix(s, ".git")
		return s
	}
	return norm(a) == norm(b)
}

func assertRepoUsable(path string) error {
	info, err := os.Stat(filepath.Join(path, ".git"))
	if err != nil || !(info.IsDir() || info.Mode().IsRegular()) {
		return fmt.Errorf("not a git repository: %s", path)
	}
	if out, err := runGit(path, "remote", "get-url", "origin"); err != nil {
		return fmt.Errorf("repo has no 'origin' remote: %s: %s", path, strings.TrimSpace(out))
	}
	return nil
}

// deployStableBinary copies the currently-running executable into
// <supportDir>/bin/radar-bridge via atomic rename, so the LaunchAgent always
// points at a Bridge-owned copy that survives cleaning the build dir. Returns
// the stable path. A no-op (still atomic) when already running from there.
func deployStableBinary(supportDir string) (string, error) {
	src, err := os.Executable()
	if err != nil {
		return "", err
	}
	src, err = filepath.EvalSymlinks(src)
	if err != nil {
		return "", err
	}
	binDir := filepath.Join(supportDir, "bin")
	if err := os.MkdirAll(binDir, 0o755); err != nil {
		return "", err
	}
	dest := filepath.Join(binDir, "radar-bridge")
	if resolved, derr := filepath.EvalSymlinks(dest); derr == nil && resolved == src {
		return dest, nil // already running from the stable location
	}
	data, err := os.ReadFile(src)
	if err != nil {
		return "", fmt.Errorf("read current binary: %w", err)
	}
	if err := writeFileAtomic(dest, data, 0o755); err != nil {
		return "", fmt.Errorf("install binary to %s: %w", dest, err)
	}
	return dest, nil
}

func installLaunchAgents(executable, dir string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	agentsDir := filepath.Join(home, "Library", "LaunchAgents")
	if err := os.MkdirAll(agentsDir, 0o755); err != nil {
		return err
	}
	logDir := filepath.Join(dir, "logs")
	servePath := filepath.Join(agentsDir, "com.ivo.radar-bridge.plist")
	syncPath := filepath.Join(agentsDir, "com.ivo.radar-bridge.sync.plist")

	servePlist := plist("com.ivo.radar-bridge", executable, "serve", logDir, true, false)
	syncPlist := plist("com.ivo.radar-bridge.sync", executable, "sync", logDir, false, true)
	if err := os.WriteFile(servePath, []byte(servePlist), 0o644); err != nil {
		return err
	}
	if err := os.WriteFile(syncPath, []byte(syncPlist), 0o644); err != nil {
		return err
	}

	domain := fmt.Sprintf("gui/%d", os.Getuid())
	for _, path := range []string{servePath, syncPath} {
		_ = exec.Command("launchctl", "bootout", domain, path).Run()
		if output, err := exec.Command("launchctl", "bootstrap", domain, path).CombinedOutput(); err != nil {
			return fmt.Errorf("launchctl bootstrap %s: %w: %s", path, err, output)
		}
	}
	return nil
}

func plist(label, executable, command, logDir string, runAtLoad, calendar bool) string {
	calendarXML := ""
	if calendar {
		calendarXML = `
  <key>StartCalendarInterval</key>
  <dict><key>Hour</key><integer>9</integer><key>Minute</key><integer>15</integer></dict>`
	}
	keepAliveXML := ""
	if runAtLoad {
		keepAliveXML = "\n  <key>KeepAlive</key><true/>"
	}
	return fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>Label</key><string>%s</string>
  <key>ProgramArguments</key>
  <array><string>%s</string><string>%s</string></array>
  <key>RunAtLoad</key><%s/>%s%s
  <key>StandardOutPath</key><string>%s</string>
  <key>StandardErrorPath</key><string>%s</string>
  <key>ProcessType</key><string>Background</string>
</dict>
</plist>
`, html.EscapeString(label), html.EscapeString(executable), html.EscapeString(command),
		map[bool]string{true: "true", false: "false"}[runAtLoad], keepAliveXML, calendarXML,
		html.EscapeString(filepath.Join(logDir, label+".log")),
		html.EscapeString(filepath.Join(logDir, label+".error.log")))
}

func (a *App) serve() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/pair", a.handlePair)
	mux.HandleFunc("/api/status", a.auth(a.handleStatus))
	mux.HandleFunc("/api/sync", a.auth(a.handleSync))
	mux.HandleFunc("/api/favorites/sync", a.auth(a.handleFavoritesSync))
	mux.HandleFunc("/api/favorites", a.auth(a.handleFavoritesGet))
	handler := a.cors(mux)

	go func() {
		_, _ = a.sync(context.Background())
	}()

	server := &http.Server{
		Addr:              a.config.Listen,
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       15 * time.Second,
		WriteTimeout:      10 * time.Minute,
		IdleTimeout:       60 * time.Second,
	}
	log.Printf("Radar Bridge listening on http://%s", a.config.Listen)
	return server.ListenAndServe()
}

func (a *App) cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := strings.TrimRight(r.Header.Get("Origin"), "/")
		if origin != "" && origin != a.allowedOrigin {
			http.Error(w, "origin not allowed", http.StatusForbidden)
			return
		}
		if origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
		}
		w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		if r.Header.Get("Access-Control-Request-Private-Network") == "true" {
			w.Header().Set("Access-Control-Allow-Private-Network", "true")
		}
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func (a *App) auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		a.mu.Lock()
		token := a.state.Token
		a.mu.Unlock()
		if r.Header.Get("Authorization") != "Bearer "+token {
			writeError(w, http.StatusUnauthorized, "invalid pairing token")
			return
		}
		next(w, r)
	}
}

func (a *App) handlePair(w http.ResponseWriter, r *http.Request) {
	returnURL := strings.TrimRight(r.URL.Query().Get("return"), "/")
	if returnURL != a.config.PagesURL {
		writeError(w, http.StatusBadRequest, "invalid return URL")
		return
	}
	a.mu.Lock()
	token := a.state.Token
	a.mu.Unlock()
	target := returnURL + "/?bridge_token=" + url.QueryEscape(token)
	http.Redirect(w, r, target, http.StatusFound)
}

func (a *App) handleStatus(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, a.status())
}

func (a *App) status() map[string]any {
	a.mu.Lock()
	defer a.mu.Unlock()
	return map[string]any{
		"connected":           true,
		"syncing":             a.syncing.Load(),
		"last_sync":           a.state.LastSync,
		"last_error":          a.state.LastError,
		"files":               len(a.state.Files),
		"favorites_enabled":   a.config.RepoPath != "",
		"favorites_last_sync": a.state.FavoritesLastSync,
		"favorites_error":     a.state.FavoritesError,
	}
}

func (a *App) handleSync(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "POST required")
		return
	}
	result, err := a.sync(r.Context())
	if err != nil {
		writeError(w, http.StatusBadGateway, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, result)
}

func (a *App) sync(ctx context.Context) (SyncResult, error) {
	a.syncMu.Lock()
	defer a.syncMu.Unlock()
	lockFile, err := os.OpenFile(filepath.Join(filepath.Dir(a.statePath), "sync.lock"), os.O_CREATE|os.O_RDWR, 0o600)
	if err != nil {
		return SyncResult{}, err
	}
	defer lockFile.Close()
	if err := syscall.Flock(int(lockFile.Fd()), syscall.LOCK_EX); err != nil {
		return SyncResult{}, err
	}
	defer syscall.Flock(int(lockFile.Fd()), syscall.LOCK_UN)

	var diskState State
	if err := readJSON(a.statePath, &diskState); err == nil {
		a.mu.Lock()
		a.state = diskState
		if a.state.Files == nil {
			a.state.Files = map[string]string{}
		}
		if a.state.Favorites == nil {
			a.state.Favorites = map[string]string{}
		}
		a.mu.Unlock()
	}
	a.syncing.Store(true)
	defer a.syncing.Store(false)

	manifestURL := a.config.PagesURL + "/sync-manifest.json"
	var manifest SyncManifest
	if err := a.fetchJSONWithRetry(ctx, manifestURL, &manifest); err != nil {
		a.recordSyncError(err)
		return SyncResult{}, err
	}

	targetRoot := filepath.Join(a.config.VaultPath, "日报", "AI情报")
	if err := os.MkdirAll(targetRoot, 0o755); err != nil {
		return SyncResult{}, err
	}

	type job struct {
		file SyncFile
		dest string
	}
	jobs := make(chan job)
	results := make(chan error, workerCount)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	a.mu.Lock()
	knownFiles := make(map[string]string, len(a.state.Files))
	for key, value := range a.state.Files {
		knownFiles[key] = value
	}
	a.mu.Unlock()

	result := SyncResult{}
	var resultMu sync.Mutex
	var workers sync.WaitGroup
	for i := 0; i < workerCount; i++ {
		workers.Add(1)
		go func() {
			defer workers.Done()
			for item := range jobs {
				if knownFiles[item.file.Path] == item.file.SHA256 {
					if _, err := os.Stat(item.dest); err == nil {
						resultMu.Lock()
						result.Skipped++
						resultMu.Unlock()
						continue
					}
				}
				if err := a.downloadFile(ctx, item.file, item.dest); err != nil {
					select {
					case results <- err:
					default:
					}
					cancel()
					return
				}
				resultMu.Lock()
				result.Updated++
				resultMu.Unlock()
				a.mu.Lock()
				a.state.Files[item.file.Path] = item.file.SHA256
				a.mu.Unlock()
			}
		}()
	}

	go func() {
		defer close(jobs)
		for _, file := range manifest.Files {
			dest, err := safeJoin(targetRoot, file.Path)
			if err != nil {
				select {
				case results <- err:
				default:
				}
				cancel()
				return
			}
			select {
			case jobs <- job{file: file, dest: dest}:
			case <-ctx.Done():
				return
			}
		}
	}()

	done := make(chan struct{})
	go func() {
		workers.Wait()
		close(done)
	}()

	select {
	case err := <-results:
		<-done
		a.recordSyncError(err)
		return result, err
	case <-done:
	}

	a.mu.Lock()
	a.state.LastSync = time.Now().Format(time.RFC3339)
	a.state.LastError = ""
	a.mu.Unlock()
	if err := a.saveState(); err != nil {
		return result, err
	}
	return result, nil
}

func (a *App) syncViaServer() (SyncResult, bool, error) {
	a.mu.Lock()
	token := a.state.Token
	a.mu.Unlock()
	request, err := http.NewRequest(http.MethodPost, "http://"+a.config.Listen+"/api/sync", nil)
	if err != nil {
		return SyncResult{}, true, err
	}
	request.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{
		Timeout: 10 * time.Minute,
		Transport: &http.Transport{
			Proxy: nil,
		},
	}
	response, err := client.Do(request)
	if err != nil {
		var networkError net.Error
		if errors.As(err, &networkError) {
			return SyncResult{}, false, nil
		}
		return SyncResult{}, true, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		var payload map[string]string
		_ = json.NewDecoder(response.Body).Decode(&payload)
		return SyncResult{}, true, fmt.Errorf("local bridge: %s", payload["error"])
	}
	var result SyncResult
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		return SyncResult{}, true, err
	}
	return result, true, nil
}

func (a *App) downloadFile(ctx context.Context, file SyncFile, dest string) error {
	parts := strings.Split(file.Path, "/")
	for i := range parts {
		parts[i] = url.PathEscape(parts[i])
	}
	sourceURL := a.config.PagesURL + "/digests/" + strings.Join(parts, "/")
	var lastErr error
	for attempt := 0; attempt < 2; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, sourceURL, nil)
		if err != nil {
			return err
		}
		resp, err := a.client.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			content, readErr := io.ReadAll(io.LimitReader(resp.Body, file.Size+1))
			_ = resp.Body.Close()
			if readErr == nil && int64(len(content)) == file.Size {
				sum := sha256.Sum256(content)
				if hex.EncodeToString(sum[:]) == file.SHA256 {
					if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
						return err
					}
					return writeFileAtomic(dest, content, 0o644)
				}
				lastErr = fmt.Errorf("checksum mismatch: %s", file.Path)
			} else {
				lastErr = fmt.Errorf("size mismatch: %s", file.Path)
			}
		} else {
			if resp != nil {
				_ = resp.Body.Close()
				lastErr = fmt.Errorf("download %s: HTTP %d", file.Path, resp.StatusCode)
			} else {
				lastErr = fmt.Errorf("download %s: %w", file.Path, err)
			}
		}
		if attempt == 0 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(time.Second):
			}
		}
	}
	return lastErr
}

func (a *App) fetchJSONWithRetry(ctx context.Context, sourceURL string, target any) error {
	var lastErr error
	for attempt := 0; attempt < 3; attempt++ {
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, sourceURL, nil)
		if err != nil {
			return err
		}
		resp, err := a.client.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			decodeErr := json.NewDecoder(resp.Body).Decode(target)
			_ = resp.Body.Close()
			if decodeErr == nil {
				return nil
			}
			lastErr = decodeErr
		} else if resp != nil {
			_ = resp.Body.Close()
			lastErr = fmt.Errorf("fetch %s: HTTP %d", sourceURL, resp.StatusCode)
		} else {
			lastErr = err
		}
		if attempt < 2 {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(time.Duration(1<<attempt) * time.Second):
			}
		}
	}
	return lastErr
}

func (a *App) recordSyncError(err error) {
	a.mu.Lock()
	a.state.LastError = err.Error()
	a.mu.Unlock()
	_ = a.saveState()
}

// handleFavoritesGet returns the repo's current favorites.json (empty when sync
// is disabled or the file does not exist yet).
func (a *App) handleFavoritesGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, http.StatusMethodNotAllowed, "GET required")
		return
	}
	if a.config.RepoPath == "" {
		writeJSON(w, http.StatusOK, FavoritesFile{Version: 1, Items: []FavoriteItem{}})
		return
	}
	file, err := readFavoritesFile(filepath.Join(a.config.RepoPath, favoritesFileName))
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusOK, file)
}

// handleFavoritesSync merges the website's full favorites set (alive + tombstones)
// into the repo favorites.json, writes the Obsidian markdown, and pushes
// favorites.json to origin/master. The browser id is trusted (validated 64-hex).
func (a *App) handleFavoritesSync(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "POST required")
		return
	}
	if a.config.RepoPath == "" {
		writeError(w, http.StatusBadRequest, "收藏同步未启用：请先 radar-bridge install --repo-url <git地址>")
		return
	}
	var request struct {
		Items []FavoriteItem `json:"items"`
	}
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	incoming := validFavorites(request.Items)

	a.favMu.Lock()
	result := a.syncFavorites(incoming)
	a.favMu.Unlock()

	// A vault-write failure is a partial success: GitHub may have published, but the
	// SecondBrain markdown is missing. Surface it as a persisted error so status()
	// keeps reporting it and the front-end keeps the sync pending until it heals.
	vaultFailed := result.pushed && !result.vaultWritten
	a.mu.Lock()
	switch {
	case result.err != nil:
		a.state.FavoritesError = result.err.Error()
	case vaultFailed:
		a.state.FavoritesLastSync = time.Now().Format(time.RFC3339)
		a.state.FavoritesError = favVaultErrorText(result.vaultErr)
	default:
		a.state.FavoritesLastSync = time.Now().Format(time.RFC3339)
		a.state.FavoritesError = ""
	}
	a.mu.Unlock()
	_ = a.saveState()

	resp := map[string]any{
		"items":         result.items,
		"total":         len(result.items),
		"alive":         countAlive(result.items),
		"pushed":        result.pushed,
		"vault_written": result.vaultWritten,
	}
	if result.err != nil {
		resp["error"] = result.err.Error()
	}
	if vaultFailed {
		resp["vault_error"] = favVaultErrorText(result.vaultErr)
	}
	writeJSON(w, http.StatusOK, resp)
}

func favVaultErrorText(detail string) string {
	if strings.TrimSpace(detail) == "" {
		return "SecondBrain 写入失败"
	}
	return "SecondBrain 写入失败：" + detail
}

// favSyncResult is the outcome of one favorites sync. pushed = reached
// origin/master; vaultWritten = the Obsidian markdown was written successfully.
// Both products must succeed for a sync to be fully complete (technical doc:
// "一次同步两份产物一起写"). vaultErr carries the markdown write error text even
// when the GitHub push succeeded, so the API/front-end can surface and retry it.
type favSyncResult struct {
	items        []FavoriteItem
	pushed       bool
	vaultWritten bool
	vaultErr     string
	err          error
}

// syncFavorites aligns the dedicated clone to origin/master, merges incoming
// items on top, writes favorites.json + the Obsidian md, then commits and pushes.
// Returns the merged set, whether the push reached origin, and any error.
//
// The clone MUST be Bridge-owned: we refuse a working tree that is dirty with
// anything other than favorites.json, and otherwise `git reset --hard
// origin/master` to guarantee fast-forward pushes (no rebase conflicts).
func (a *App) syncFavorites(incoming []FavoriteItem) favSyncResult {
	repo := a.config.RepoPath
	favPath := filepath.Join(repo, favoritesFileName)

	if err := assertCleanExceptFavorites(repo); err != nil {
		return favSyncResult{err: err}
	}

	// writeProducts writes favorites.json + the Obsidian markdown for a merged set.
	// A favorites.json write failure is fatal (returned as err). A markdown failure
	// does NOT abort (GitHub may still publish) but is reported via vaultWritten +
	// vaultErr so the API/front-end can flag the partial result and retry it.
	writeProducts := func(merged []FavoriteItem) (vaultWritten bool, vaultErr string, fatal error) {
		if err := writeFavoritesFile(favPath, merged); err != nil {
			return false, "", err
		}
		if err := a.writeFavoritesMarkdown(merged); err != nil {
			log.Printf("favorites: write markdown failed: %v", err)
			return false, err.Error(), nil
		}
		return true, "", nil
	}

	if out, err := runGit(repo, "fetch", "origin", "master"); err != nil {
		// Offline: merge against the local working file, write json+md, do not push.
		log.Printf("favorites: git fetch failed (offline?): %v: %s", err, strings.TrimSpace(out))
		base, rerr := readFavoritesItems(favPath)
		if rerr != nil {
			return favSyncResult{err: fmt.Errorf("读取本地 favorites.json 失败，已中止同步：%w", rerr)}
		}
		merged := mergeFavorites(base, incoming)
		vaultWritten, vaultErr, werr := writeProducts(merged)
		if werr != nil {
			return favSyncResult{items: merged, vaultWritten: vaultWritten, vaultErr: vaultErr, err: werr}
		}
		return favSyncResult{items: merged, vaultWritten: vaultWritten, vaultErr: vaultErr, err: errors.New("离线：已写本地，待联网后自动推送")}
	}

	var merged []FavoriteItem
	var vaultWritten bool
	var vaultErr string
	var lastErr error
	for attempt := 0; attempt < 2; attempt++ {
		if err := a.gitAlignToOrigin(repo); err != nil {
			return favSyncResult{items: merged, err: err}
		}
		// Read the freshly-aligned baseline. A corrupt/unreadable file fails closed:
		// abort rather than overwrite an existing library with just the incoming set.
		base, rerr := readFavoritesItems(favPath)
		if rerr != nil {
			return favSyncResult{err: fmt.Errorf("读取远端 favorites.json 失败，已中止同步（避免覆盖）：%w", rerr)}
		}
		merged = mergeFavorites(base, incoming)
		var werr error
		if vaultWritten, vaultErr, werr = writeProducts(merged); werr != nil {
			return favSyncResult{items: merged, vaultWritten: vaultWritten, vaultErr: vaultErr, err: werr}
		}

		if out, err := runGit(repo, "add", "--", favoritesFileName); err != nil {
			return favSyncResult{items: merged, vaultWritten: vaultWritten, vaultErr: vaultErr, err: fmt.Errorf("git add: %w: %s", err, strings.TrimSpace(out))}
		}
		// Nothing staged vs origin/master → already in sync (md was still rewritten above).
		if _, err := runGit(repo, "diff", "--cached", "--quiet", "--", favoritesFileName); err == nil {
			return favSyncResult{items: merged, pushed: true, vaultWritten: vaultWritten, vaultErr: vaultErr}
		}
		msg := fmt.Sprintf("favorites: sync (%d alive)", countAlive(merged))
		if out, err := runGit(repo, "commit", "-m", msg, "--", favoritesFileName); err != nil {
			return favSyncResult{items: merged, vaultWritten: vaultWritten, vaultErr: vaultErr, err: fmt.Errorf("git commit: %w: %s", err, strings.TrimSpace(out))}
		}
		if out, err := runGit(repo, "push", "origin", "HEAD:master"); err == nil {
			return favSyncResult{items: merged, pushed: true, vaultWritten: vaultWritten, vaultErr: vaultErr}
		} else {
			lastErr = fmt.Errorf("git push: %s", strings.TrimSpace(out))
			log.Printf("favorites: push attempt %d failed: %s", attempt+1, strings.TrimSpace(out))
			if out2, err2 := runGit(repo, "fetch", "origin", "master"); err2 != nil {
				return favSyncResult{items: merged, vaultWritten: vaultWritten, vaultErr: vaultErr, err: fmt.Errorf("git fetch on retry: %w: %s", err2, strings.TrimSpace(out2))}
			}
			// next loop: gitAlignToOrigin drops our commit and re-merges on the new base.
		}
	}
	return favSyncResult{items: merged, vaultWritten: vaultWritten, vaultErr: vaultErr, err: lastErr}
}

func runGit(dir string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	cmd := exec.CommandContext(ctx, "git", args...)
	if dir != "" {
		cmd.Dir = dir
	}
	// Never block on interactive credential prompts (loopback daemon).
	cmd.Env = append(os.Environ(), "GIT_TERMINAL_PROMPT=0")
	out, err := cmd.CombinedOutput()
	return string(out), err
}

// ownershipMarker lives inside .git/ (never committed, never in the work tree).
// Its presence means this clone was created/adopted by Bridge specifically for
// favorites sync, so destructive alignment (git reset --hard) is safe here.
const ownershipMarker = "radar-bridge-owned"

func ownershipMarkerPath(repo string) string {
	return filepath.Join(repo, ".git", ownershipMarker)
}
func markRepoOwned(repo string) error {
	gitDir := filepath.Join(repo, ".git")
	// Only mark a real .git directory (skip submodule/.git-file edge cases).
	if info, err := os.Stat(gitDir); err != nil || !info.IsDir() {
		return nil
	}
	return os.WriteFile(ownershipMarkerPath(repo), []byte("favorites sync clone\n"), 0o600)
}
func repoIsOwned(repo string) bool {
	_, err := os.Stat(ownershipMarkerPath(repo))
	return err == nil
}

// gitAlignToOrigin brings the working tree to origin/master before merging.
//
//   - Bridge-owned clone (has ownership marker): `git reset --hard origin/master`
//     is safe — Bridge created it solely for favorites, there is nothing to lose.
//   - User-supplied --repo (no marker): NEVER reset --hard. We refuse unless the
//     clone is on master with no commits ahead of origin/master, then fast-forward
//     only. This protects a user who pointed --repo at their own working clone
//     (feature branch, unpushed commits, etc. — see review §5.3).
func (a *App) gitAlignToOrigin(repo string) error {
	if repoIsOwned(repo) {
		if out, err := runGit(repo, "reset", "--hard", "origin/master"); err != nil {
			return fmt.Errorf("git reset: %w: %s", err, strings.TrimSpace(out))
		}
		return nil
	}
	// Unowned repo: only a safe fast-forward, no history rewrite.
	branch, err := runGit(repo, "rev-parse", "--abbrev-ref", "HEAD")
	if err != nil {
		return fmt.Errorf("git rev-parse: %w: %s", err, strings.TrimSpace(branch))
	}
	if strings.TrimSpace(branch) != "master" {
		return fmt.Errorf("仓库当前不在 master 分支（%s）；为保护你的工作区，已拒绝同步。请改用 Bridge 专用 clone（install --repo-url）", strings.TrimSpace(branch))
	}
	if ahead, _ := runGit(repo, "rev-list", "--count", "origin/master..HEAD"); strings.TrimSpace(ahead) != "0" {
		return errors.New("本地 master 领先 origin/master（有未推送提交）；为保护你的提交，已拒绝同步。请改用 Bridge 专用 clone（install --repo-url）")
	}
	if out, err := runGit(repo, "merge", "--ff-only", "origin/master"); err != nil {
		return fmt.Errorf("git merge --ff-only: %w: %s", err, strings.TrimSpace(out))
	}
	return nil
}

// assertCleanExceptFavorites refuses to run if the working tree has uncommitted
// changes to anything other than favorites.json — protecting a user who points
// --repo at their primary working clone instead of a dedicated one.
func assertCleanExceptFavorites(repo string) error {
	out, err := runGit(repo, "status", "--porcelain")
	if err != nil {
		return fmt.Errorf("git status failed: %w: %s", err, strings.TrimSpace(out))
	}
	for _, line := range strings.Split(out, "\n") {
		line = strings.TrimRight(line, "\r")
		if strings.TrimSpace(line) == "" {
			continue
		}
		path := line
		if len(line) > 3 {
			path = line[3:]
		}
		if i := strings.Index(path, " -> "); i >= 0 {
			path = path[i+4:]
		}
		path = strings.Trim(strings.TrimSpace(path), "\"")
		if path != favoritesFileName {
			return fmt.Errorf("仓库工作区有未提交改动（%s）；请用 Bridge 专用 clone（install --repo-url）", path)
		}
	}
	return nil
}

func validFavorites(items []FavoriteItem) []FavoriteItem {
	out := make([]FavoriteItem, 0, len(items))
	for _, it := range items {
		if !hex64Pattern.MatchString(it.ID) {
			continue
		}
		if it.Kind != "report" && it.Kind != "link" && it.Kind != "excerpt" {
			continue
		}
		if !datePattern.MatchString(it.Source.Date) || !reportPattern.MatchString(it.Source.Report) {
			continue
		}
		out = append(out, it)
	}
	return out
}

func countAlive(items []FavoriteItem) int {
	n := 0
	for _, it := range items {
		if it.DeletedAt == "" {
			n++
		}
	}
	return n
}

// mergeFavorites unions base and incoming by id (incoming folded last so it wins
// content fields), keeps the earliest created_at, applies the tombstone/revive
// rule, and sorts newest-first.
func mergeFavorites(base, incoming []FavoriteItem) []FavoriteItem {
	byID := map[string]FavoriteItem{}
	order := []string{}
	fold := func(it FavoriteItem) {
		if it.ID == "" {
			return
		}
		existing, ok := byID[it.ID]
		if !ok {
			byID[it.ID] = it
			order = append(order, it.ID)
			return
		}
		byID[it.ID] = mergeTwo(existing, it)
	}
	for _, it := range base {
		fold(it)
	}
	for _, it := range incoming {
		fold(it)
	}
	out := make([]FavoriteItem, 0, len(order))
	for _, id := range order {
		out = append(out, byID[id])
	}
	sortFavorites(out)
	return out
}

// mergeTwo combines two records of the same favorite id; b is the newer input.
func mergeTwo(a, b FavoriteItem) FavoriteItem {
	out := a
	if out.Kind == "" {
		out.Kind = b.Kind
	}
	out.Title = preferNonEmpty(b.Title, a.Title)
	out.URL = preferNonEmpty(b.URL, a.URL)
	out.Site = preferNonEmpty(b.Site, a.Site)
	out.Excerpt = preferNonEmpty(b.Excerpt, a.Excerpt)
	out.Section = preferNonEmpty(b.Section, a.Section)
	out.Source = mergeSource(a.Source, b.Source)

	// created_at = earliest first-favorite (stable sort key).
	out.CreatedAt = earliest(a.CreatedAt, b.CreatedAt)
	// revived_at = most recent re-favorite event seen on either side.
	out.RevivedAt = latest(a.RevivedAt, b.RevivedAt)
	delTime := latest(a.DeletedAt, b.DeletedAt)
	// Life/death uses the latest ADD-class event (first favorite OR revive). A
	// revive after a delete keeps the item alive even when a stale tombstone is
	// merged in later, because revived_at preserves that event time.
	aliveTime := latest(latest(a.CreatedAt, b.CreatedAt), out.RevivedAt)
	if delTime != "" && (aliveTime == "" || aliveTime <= delTime) {
		out.DeletedAt = delTime // tombstone
	} else {
		out.DeletedAt = "" // alive: a (re-)favorite happened at/after the deletion
	}
	return out
}

func mergeSource(a, b FavoriteSource) FavoriteSource {
	return FavoriteSource{
		Date:   preferNonEmpty(b.Date, a.Date),
		Report: preferNonEmpty(b.Report, a.Report),
		Label:  preferNonEmpty(b.Label, a.Label),
		URL:    preferNonEmpty(b.URL, a.URL),
		Anchor: preferNonEmpty(b.Anchor, a.Anchor),
	}
}

func sortFavorites(items []FavoriteItem) {
	sort.SliceStable(items, func(i, j int) bool {
		if items[i].CreatedAt != items[j].CreatedAt {
			return items[i].CreatedAt > items[j].CreatedAt // newest first
		}
		return items[i].ID < items[j].ID
	})
}

func preferNonEmpty(primary, fallback string) string {
	if strings.TrimSpace(primary) != "" {
		return primary
	}
	return fallback
}
func earliest(a, b string) string {
	switch {
	case a == "":
		return b
	case b == "":
		return a
	case a < b:
		return a
	default:
		return b
	}
}
func latest(a, b string) string {
	if a > b {
		return a
	}
	return b
}

func readFavoritesItems(path string) ([]FavoriteItem, error) {
	f, err := readFavoritesFile(path)
	if err != nil {
		return nil, err
	}
	return f.Items, nil
}

// readFavoritesFile reads favorites.json. A MISSING file is treated as an empty
// library (first run). Any other read/parse error is returned so the caller can
// fail closed — never silently overwriting an existing-but-unreadable library.
func readFavoritesFile(path string) (FavoritesFile, error) {
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		return FavoritesFile{Version: 1, Items: []FavoriteItem{}}, nil
	}
	if err != nil {
		return FavoritesFile{}, fmt.Errorf("read %s: %w", filepath.Base(path), err)
	}
	items, perr := parseFavorites(data)
	if perr != nil {
		return FavoritesFile{}, perr
	}
	return FavoritesFile{Version: 1, Items: items}, nil
}

// parseFavorites unmarshals favorites.json. An empty file is an empty library;
// malformed JSON is a hard error (do NOT treat as empty, or a corrupt remote
// file would be silently overwritten on the next sync).
func parseFavorites(data []byte) ([]FavoriteItem, error) {
	if len(strings.TrimSpace(string(data))) == 0 {
		return []FavoriteItem{}, nil
	}
	var f FavoritesFile
	if err := json.Unmarshal(data, &f); err != nil {
		return nil, fmt.Errorf("favorites.json 解析失败（可能已损坏）：%w", err)
	}
	if f.Items == nil {
		return []FavoriteItem{}, nil
	}
	return f.Items, nil
}
func writeFavoritesFile(path string, items []FavoriteItem) error {
	if items == nil {
		items = []FavoriteItem{}
	}
	file := FavoritesFile{Version: 1, UpdatedAt: time.Now().UTC().Format(time.RFC3339), Items: items}
	return writeJSONAtomic(path, file, 0o644)
}

// ── Obsidian markdown: 收藏/AI情报收藏.md ──
func (a *App) writeFavoritesMarkdown(items []FavoriteItem) error {
	dest, err := safeJoin(a.config.VaultPath, filepath.Join("收藏", "AI情报收藏.md"))
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		return err
	}
	return writeFileAtomic(dest, []byte(renderFavoritesMarkdown(items, a.config.PagesURL)), 0o644)
}

func renderFavoritesMarkdown(items []FavoriteItem, pagesURL string) string {
	alive := make([]FavoriteItem, 0, len(items))
	for _, it := range items {
		if it.DeletedAt == "" {
			alive = append(alive, it)
		}
	}
	sortFavorites(alive)

	kindLabel := map[string]string{"report": "整篇", "link": "链接", "excerpt": "段落"}
	var b strings.Builder
	b.WriteString("# AI 情报收藏\n\n")
	b.WriteString("> 由 Radar Bridge 自动生成，请勿手改。最后更新：")
	b.WriteString(time.Now().Format("2006-01-02 15:04"))
	b.WriteString("\n")

	lastDay := ""
	for _, it := range alive {
		day := favDay(it.CreatedAt)
		if day != lastDay {
			b.WriteString("\n## " + day + "\n")
			lastDay = day
		}
		label := kindLabel[it.Kind]
		if label == "" {
			label = it.Kind
		}
		b.WriteString("\n### [" + label + "] " + mdInline(it.Title) + "\n")
		if it.Kind == "link" && it.Site != "" {
			b.WriteString("- 网站：" + it.Site + "\n")
		}
		src := it.Source.Date
		if it.Source.Label != "" {
			src += " · " + it.Source.Label
		}
		b.WriteString("- 来源：" + src + "\n")
		b.WriteString("- 回到 Radar：<" + radarFocusURL(pagesURL, it) + ">\n")
		if it.Kind == "link" && it.URL != "" {
			b.WriteString("- 原始链接：<" + it.URL + ">\n")
		}
		if strings.TrimSpace(it.Excerpt) != "" {
			b.WriteString("\n> " + mdQuote(it.Excerpt) + "\n")
		}
	}
	return b.String()
}

// favDay groups by the favorite's LOCAL calendar date (technical doc §3.6:
// "created_at 当地日期"). A favorite saved just after midnight Shanghai time
// must land on that local day, not the previous UTC day.
func favDay(iso string) string {
	t, err := time.Parse(time.RFC3339, iso)
	if err != nil {
		if len(iso) >= 10 {
			return iso[:10]
		}
		return "未知日期"
	}
	return t.Local().Format("2006-01-02")
}
func radarFocusURL(pagesURL string, it FavoriteItem) string {
	focus := it.ID
	if it.Kind == "excerpt" && it.Source.Anchor != "" {
		focus = it.Source.Anchor
	}
	return fmt.Sprintf("%s/?focus=%s#%s/%s", pagesURL, url.QueryEscape(focus), it.Source.Date, it.Source.Report)
}
func mdInline(s string) string {
	return strings.TrimSpace(spacePattern.ReplaceAllString(s, " "))
}
func mdQuote(s string) string {
	return strings.ReplaceAll(strings.TrimSpace(s), "\n", "\n> ")
}

func favoriteID(candidate FavoriteCandidate) (string, error) {
	if !datePattern.MatchString(candidate.ReportDate) || !reportPattern.MatchString(candidate.ReportType) {
		return "", errors.New("invalid report source")
	}
	switch candidate.Kind {
	case "link":
		normalized, err := normalizeFavoriteURL(candidate.URL)
		if err != nil {
			return "", err
		}
		sum := sha256.Sum256([]byte("link\n" + normalized))
		return hex.EncodeToString(sum[:]), nil
	case "excerpt":
		excerpt := normalizeText(candidate.Excerpt)
		if len(excerpt) < 8 {
			return "", errors.New("excerpt is too short")
		}
		sum := sha256.Sum256([]byte("excerpt\n" + candidate.ReportDate + "\n" + candidate.ReportType + "\n" + excerpt))
		return hex.EncodeToString(sum[:]), nil
	default:
		return "", errors.New("favorite kind must be link or excerpt")
	}
}

func normalizeFavoriteURL(raw string) (string, error) {
	parsed, err := url.Parse(strings.TrimSpace(raw))
	if err != nil || (parsed.Scheme != "http" && parsed.Scheme != "https") || parsed.Host == "" {
		return "", errors.New("invalid favorite URL")
	}
	parsed.Scheme = strings.ToLower(parsed.Scheme)
	parsed.Host = strings.ToLower(parsed.Host)
	parsed.Fragment = ""
	query := parsed.Query()
	for key := range query {
		lower := strings.ToLower(key)
		if strings.HasPrefix(lower, "utm_") || lower == "ref" || lower == "source" {
			query.Del(key)
		}
	}
	parsed.RawQuery = query.Encode()
	if parsed.Path == "/" {
		parsed.Path = ""
	}
	return parsed.String(), nil
}

func normalizeText(value string) string {
	return strings.ToLower(spacePattern.ReplaceAllString(strings.TrimSpace(value), " "))
}

func safeFilename(value string) string {
	replacer := strings.NewReplacer(
		"/", "-", "\\", "-", ":", "-", "*", "", "?", "", "\"", "", "<", "", ">", "", "|", "-",
		"\n", " ", "\r", " ", "\t", " ",
	)
	value = spacePattern.ReplaceAllString(strings.TrimSpace(replacer.Replace(value)), " ")
	value = strings.Trim(value, ". ")
	runes := []rune(value)
	if len(runes) > 60 {
		value = string(runes[:60])
	}
	if value == "" {
		return "情报收藏"
	}
	return value
}

func safeJoin(root, relative string) (string, error) {
	clean := filepath.Clean(filepath.FromSlash(relative))
	if clean == "." || filepath.IsAbs(clean) || clean == ".." || strings.HasPrefix(clean, ".."+string(filepath.Separator)) {
		return "", errors.New("unsafe relative path")
	}
	full := filepath.Join(root, clean)
	rel, err := filepath.Rel(root, full)
	if err != nil || rel == ".." || strings.HasPrefix(rel, ".."+string(filepath.Separator)) {
		return "", errors.New("path escapes target directory")
	}
	return full, nil
}

func randomToken() (string, error) {
	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func jsonString(value string) string {
	data, _ := json.Marshal(value)
	return string(data)
}

func decodeJSON(r *http.Request, target any) error {
	defer r.Body.Close()
	decoder := json.NewDecoder(io.LimitReader(r.Body, 256*1024))
	decoder.DisallowUnknownFields()
	return decoder.Decode(target)
}

func readJSON(path string, target any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, target)
}

func writeJSONAtomic(path string, value any, mode os.FileMode) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return err
	}
	return writeFileAtomic(path, append(data, '\n'), mode)
}

func writeFileAtomic(path string, content []byte, mode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	temp, err := os.CreateTemp(filepath.Dir(path), ".radar-*")
	if err != nil {
		return err
	}
	tempPath := temp.Name()
	defer os.Remove(tempPath)
	if err := temp.Chmod(mode); err != nil {
		_ = temp.Close()
		return err
	}
	if _, err := temp.Write(content); err != nil {
		_ = temp.Close()
		return err
	}
	if err := temp.Close(); err != nil {
		return err
	}
	return os.Rename(tempPath, path)
}

func (a *App) saveState() error {
	a.mu.Lock()
	defer a.mu.Unlock()
	return writeJSONAtomic(a.statePath, a.state, 0o600)
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(value)
}

func writeError(w http.ResponseWriter, status int, message string) {
	writeJSON(w, status, map[string]string{"error": message})
}
