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
)

type Config struct {
	VaultPath string `json:"vault_path"`
	PagesURL  string `json:"pages_url"`
	Listen    string `json:"listen"`
}

type State struct {
	Token     string            `json:"token"`
	LastSync string            `json:"last_sync,omitempty"`
	LastError string           `json:"last_error,omitempty"`
	Files     map[string]string `json:"files"`
	Favorites map[string]string `json:"favorites"`
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

type App struct {
	config        Config
	configPath    string
	statePath     string
	client        *http.Client
	mu            sync.Mutex
	state         State
	syncing       atomic.Bool
	syncMu        sync.Mutex
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
	config := Config{
		VaultPath: filepath.Clean(*vault),
		PagesURL:  strings.TrimRight(*pagesURL, "/"),
		Listen:    *listen,
	}
	if err := writeJSONAtomic(filepath.Join(dir, "config.json"), config, 0o600); err != nil {
		return err
	}

	executable, err := os.Executable()
	if err != nil {
		return err
	}
	if err := installLaunchAgents(executable, dir); err != nil {
		return err
	}
	fmt.Printf("Radar Bridge installed for %s\n", config.VaultPath)
	return nil
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
	mux.HandleFunc("/api/favorites/check", a.auth(a.handleFavoriteCheck))
	mux.HandleFunc("/api/favorites", a.auth(a.handleFavorite))
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
		"connected":  true,
		"syncing":    a.syncing.Load(),
		"last_sync":  a.state.LastSync,
		"last_error": a.state.LastError,
		"files":      len(a.state.Files),
		"favorites":  len(a.state.Favorites),
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

func (a *App) handleFavoriteCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "POST required")
		return
	}
	var request struct {
		Items []FavoriteCandidate `json:"items"`
	}
	if err := decodeJSON(r, &request); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	saved := make([]bool, len(request.Items))
	a.mu.Lock()
	for i, candidate := range request.Items {
		id, err := favoriteID(candidate)
		if err != nil {
			continue
		}
		relative, ok := a.state.Favorites[id]
		if !ok {
			continue
		}
		_, err = os.Stat(filepath.Join(a.config.VaultPath, relative))
		saved[i] = err == nil
	}
	a.mu.Unlock()
	writeJSON(w, http.StatusOK, map[string]any{"saved": saved})
}

func (a *App) handleFavorite(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, http.StatusMethodNotAllowed, "POST required")
		return
	}
	var candidate FavoriteCandidate
	if err := decodeJSON(r, &candidate); err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := favoriteID(candidate)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	a.mu.Lock()
	if relative, ok := a.state.Favorites[id]; ok {
		if _, statErr := os.Stat(filepath.Join(a.config.VaultPath, relative)); statErr == nil {
			a.mu.Unlock()
			writeJSON(w, http.StatusOK, map[string]any{"id": id, "path": relative, "created": false})
			return
		}
		delete(a.state.Favorites, id)
	}
	a.mu.Unlock()

	month := candidate.ReportDate[:7]
	title := strings.TrimSpace(candidate.Title)
	if title == "" {
		title = "情报收藏"
	}
	filename := safeFilename(title) + "-" + id[:8] + ".md"
	relative := filepath.Join("收藏", candidate.ReportDate[:4], month, filename)
	dest, err := safeJoin(a.config.VaultPath, relative)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}
	if err := os.MkdirAll(filepath.Dir(dest), 0o755); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	if err := writeFileAtomic(dest, []byte(renderFavorite(candidate, id)), 0o644); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	a.mu.Lock()
	a.state.Favorites[id] = relative
	a.mu.Unlock()
	if err := a.saveState(); err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}
	writeJSON(w, http.StatusCreated, map[string]any{"id": id, "path": relative, "created": true})
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

func renderFavorite(candidate FavoriteCandidate, id string) string {
	title := strings.TrimSpace(candidate.Title)
	if title == "" {
		title = "情报收藏"
	}
	sourceWiki := fmt.Sprintf("[[日报/AI情报/%s/%s|%s %s]]",
		candidate.ReportDate, candidate.ReportType, candidate.ReportDate, candidate.ReportType)
	var builder strings.Builder
	builder.WriteString("---\n")
	builder.WriteString("type: favorite\n")
	builder.WriteString("status: unread\n")
	builder.WriteString("favorite_kind: " + candidate.Kind + "\n")
	builder.WriteString("favorite_id: " + id + "\n")
	builder.WriteString("title: " + jsonString(title) + "\n")
	if candidate.URL != "" {
		builder.WriteString("url: " + jsonString(candidate.URL) + "\n")
	}
	builder.WriteString("source_report: " + jsonString(sourceWiki) + "\n")
	builder.WriteString("report_date: " + candidate.ReportDate + "\n")
	builder.WriteString("report_type: " + candidate.ReportType + "\n")
	builder.WriteString("captured: " + time.Now().Format(time.RFC3339) + "\n")
	builder.WriteString("tags:\n  - 收藏\n  - AI情报\n")
	builder.WriteString("---\n\n# " + title + "\n\n")
	if candidate.URL != "" {
		builder.WriteString("原始链接：[" + title + "](" + candidate.URL + ")\n\n")
	}
	builder.WriteString("来源报告：" + sourceWiki + "\n\n")
	if candidate.Section != "" {
		builder.WriteString("来源章节：" + candidate.Section + "\n\n")
	}
	builder.WriteString("## 收藏时上下文\n\n" + strings.TrimSpace(candidate.Excerpt) + "\n\n")
	builder.WriteString("## 我的笔记\n")
	return builder.String()
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
