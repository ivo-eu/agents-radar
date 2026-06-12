package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestNormalizeFavoriteURL(t *testing.T) {
	got, err := normalizeFavoriteURL("HTTPS://Example.COM/article/?utm_source=radar&b=2&a=1#section")
	if err != nil {
		t.Fatal(err)
	}
	if got != "https://example.com/article/?a=1&b=2" {
		t.Fatalf("unexpected URL: %s", got)
	}
}

func TestFavoriteIDStableForTrackingParameters(t *testing.T) {
	base := FavoriteCandidate{
		Kind:       "link",
		ReportDate: "2026-06-09",
		ReportType: "ai-hn",
	}
	first := base
	first.URL = "https://example.com/post?utm_source=one"
	second := base
	second.URL = "https://example.com/post?utm_source=two"

	firstID, err := favoriteID(first)
	if err != nil {
		t.Fatal(err)
	}
	secondID, err := favoriteID(second)
	if err != nil {
		t.Fatal(err)
	}
	if firstID != secondID {
		t.Fatal("tracking parameters should not create duplicate favorites")
	}
}

func TestSafeFilename(t *testing.T) {
	got := safeFilename(`A/B: "Agent" | Notes?`)
	if strings.ContainsAny(got, `/\:*?"<>|`) {
		t.Fatalf("unsafe filename: %s", got)
	}
}

func TestSafeJoinRejectsTraversal(t *testing.T) {
	root := t.TempDir()
	if _, err := safeJoin(root, "../../secret"); err == nil {
		t.Fatal("expected traversal to be rejected")
	}
	got, err := safeJoin(root, "2026-06-09/ai-hn.md")
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(got, filepath.Clean(root)+string(filepath.Separator)) {
		t.Fatalf("path escaped root: %s", got)
	}
}

// These IDs are also asserted byte-for-byte by the website test
// src/__tests__/favorites-id.test.ts. The two sets MUST stay equal so Phase 2
// multi-device sync matches favorites the browser created against favorites the
// Bridge writes. Changing normalization on either side breaks both tests.
func TestFavoriteIDMatchesWebsiteContract(t *testing.T) {
	cases := []struct {
		name       string
		cand       FavoriteCandidate
		expectedID string
	}{
		{
			"link-tracking",
			FavoriteCandidate{Kind: "link", URL: "https://example.com/post?utm_source=one&b=2&a=1#frag", ReportDate: "2026-06-09", ReportType: "ai-hn"},
			"c7d086f6048104063f2da1ddc84b454f1a59e5e32c184771b605afae4d30cd7b",
		},
		{
			"link-trailing-slash",
			FavoriteCandidate{Kind: "link", URL: "https://Example.COM/article/?utm_source=radar&b=2&a=1#section", ReportDate: "2026-06-09", ReportType: "ai-hn"},
			"96a07bb55d0639622c75f8ba1d63114bc6ddcb985c0cea16841294bda715d331",
		},
		{
			"link-root-slash",
			FavoriteCandidate{Kind: "link", URL: "HTTP://EXAMPLE.com/", ReportDate: "2026-06-09", ReportType: "ai-hn"},
			"2774e215cd358e72cc2709391daa99d7d98b0631b643575d7729a4b46a693811",
		},
		{
			"link-ref-source",
			FavoriteCandidate{Kind: "link", URL: "https://example.com/p?ref=x&source=y&keep=1", ReportDate: "2026-06-09", ReportType: "ai-hn"},
			"600663b46a20576ace4737869b8d489244beebddd394df5ac40e47a3813a73d2",
		},
		{
			"link-space-in-query",
			FavoriteCandidate{Kind: "link", URL: "https://example.com/s?q=hello world&x=a+b", ReportDate: "2026-06-09", ReportType: "ai-hn"},
			"f6038fd489bb7df8cf07eade211bd039765d071dae15d5b4b58b5abfaed3863c",
		},
		{
			"link-unicode-path",
			FavoriteCandidate{Kind: "link", URL: "https://example.com/路径/page?名=值", ReportDate: "2026-06-09", ReportType: "ai-hn"},
			"ff92cca91ef946be3a939b3b42465f8f45206b5c2037f7b4e11bce48fb6c0267",
		},
		{
			"link-port",
			FavoriteCandidate{Kind: "link", URL: "https://example.com:8443/x?z=1", ReportDate: "2026-06-09", ReportType: "ai-hn"},
			"ea73621b36e539f39eee4f960c10e7a7305429f3087c2005e5f5f69f340659f0",
		},
		{
			"excerpt-ws",
			FavoriteCandidate{Kind: "excerpt", Excerpt: "Agent   memory\narchitecture", ReportDate: "2026-06-09", ReportType: "ai-agents"},
			"ee7cf0e9715741daf27c0c223665e81871bd6354f3daa6c3e385e5d4f6334a6c",
		},
		{
			"excerpt-case",
			FavoriteCandidate{Kind: "excerpt", Excerpt: "Provider COMPATIBILITY Notes", ReportDate: "2026-06-10", ReportType: "ai-agents"},
			"87743bf98cbc5fc085e39d6088893760c46c584ae5d5a76f4b7e3f52deef6027",
		},
		{
			"excerpt-unicode",
			FavoriteCandidate{Kind: "excerpt", Excerpt: "关于 GitHub Trending 的趋势判断", ReportDate: "2026-06-09", ReportType: "ai-trending"},
			"edd2b2c7cc7f2c071424aaf1fb5c256632295bb0e228a69d31c98cb28df77b01",
		},
	}
	for _, c := range cases {
		id, err := favoriteID(c.cand)
		if err != nil {
			t.Fatalf("%s: %v", c.name, err)
		}
		if id != c.expectedID {
			t.Fatalf("%s: id mismatch\n  want %s\n  got  %s", c.name, c.expectedID, id)
		}
	}
}

func TestExcerptFavoriteIDNormalizesWhitespace(t *testing.T) {
	first := FavoriteCandidate{
		Kind:       "excerpt",
		Excerpt:    "Agent   memory\narchitecture",
		ReportDate: "2026-06-09",
		ReportType: "ai-agents",
	}
	second := first
	second.Excerpt = "agent memory architecture"
	firstID, err := favoriteID(first)
	if err != nil {
		t.Fatal(err)
	}
	secondID, err := favoriteID(second)
	if err != nil {
		t.Fatal(err)
	}
	if firstID != secondID {
		t.Fatal("normalized excerpts should have the same ID")
	}
}

// ── Phase 2: favorites sync ──

const id64a = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
const id64b = "bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb"

// favMdPath is where the unified favorites note lands inside a test vault, kept in
// sync with favoritesMarkdownRel (日报/AI情报/收藏/AI情报收藏.md).
func favMdPath(vault string) string { return filepath.Join(vault, favoritesMarkdownRel) }

// blockFavMarkdownWrite makes the markdown's parent dir a regular file so the
// subsequent MkdirAll/write fails — used to exercise the vault-write-failure path.
func blockFavMarkdownWrite(t *testing.T, vault string) {
	t.Helper()
	parent := filepath.Dir(favMdPath(vault)) // …/日报/AI情报/收藏
	if err := os.MkdirAll(filepath.Dir(parent), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(parent, []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
}

func item(id, kind, created string) FavoriteItem {
	return FavoriteItem{
		ID: id, Kind: kind, CreatedAt: created,
		Source: FavoriteSource{Date: "2026-06-10", Report: "ai-trending", Label: "GitHub AI 趋势"},
	}
}

func TestValidFavoritesSkipsInvalid(t *testing.T) {
	in := []FavoriteItem{
		item(id64a, "link", "2026-06-10T00:00:00Z"),                                                  // ok
		item("short", "link", "2026-06-10T00:00:00Z"),                                                // bad id
		{ID: id64b, Kind: "bogus", Source: FavoriteSource{Date: "2026-06-10", Report: "ai-cli"}},     // bad kind
		{ID: id64b, Kind: "link", Source: FavoriteSource{Date: "2026/06/10", Report: "ai-cli"}},      // bad date
		{ID: id64b, Kind: "link", Source: FavoriteSource{Date: "2026-06-10", Report: "Bad Report!"}}, // bad report
	}
	got := validFavorites(in)
	if len(got) != 1 || got[0].ID != id64a {
		t.Fatalf("expected only the one valid item, got %d", len(got))
	}
}

func TestMergeKeepsEarliestCreatedAndFillsFields(t *testing.T) {
	base := []FavoriteItem{{ID: id64a, Kind: "link", CreatedAt: "2026-06-10T09:00:00Z",
		Source: FavoriteSource{Date: "2026-06-10", Report: "ai-trending"}}}
	incoming := []FavoriteItem{{ID: id64a, Kind: "link", Title: "Cursor", Site: "github.com",
		CreatedAt: "2026-06-10T08:00:00Z", Source: FavoriteSource{Date: "2026-06-10", Report: "ai-trending", Label: "GitHub AI 趋势"}}}
	merged := mergeFavorites(base, incoming)
	if len(merged) != 1 {
		t.Fatalf("expected 1 merged item, got %d", len(merged))
	}
	m := merged[0]
	if m.CreatedAt != "2026-06-10T08:00:00Z" {
		t.Fatalf("expected earliest created_at, got %s", m.CreatedAt)
	}
	if m.Title != "Cursor" || m.Site != "github.com" || m.Source.Label != "GitHub AI 趋势" {
		t.Fatalf("incoming fields should fill missing values: %+v", m)
	}
}

func TestMergeTombstoneWinsAndRevives(t *testing.T) {
	// add then delete → tombstoned
	add := item(id64a, "link", "2026-06-10T08:00:00Z")
	del := add
	del.DeletedAt = "2026-06-10T09:00:00Z"
	merged := mergeFavorites([]FavoriteItem{add}, []FavoriteItem{del})
	if merged[0].DeletedAt == "" {
		t.Fatal("delete after add should tombstone the item")
	}
	if countAlive(merged) != 0 {
		t.Fatal("tombstoned item must not count as alive")
	}

	// re-add after delete (created_at later than deleted_at) → revived
	readd := item(id64a, "link", "2026-06-10T10:00:00Z")
	revived := mergeFavorites([]FavoriteItem{del}, []FavoriteItem{readd})
	if revived[0].DeletedAt != "" {
		t.Fatalf("re-add after delete should clear tombstone, got %q", revived[0].DeletedAt)
	}
	if countAlive(revived) != 1 {
		t.Fatal("revived item should be alive")
	}
}

// Review §5.2: add(08:00) → delete(09:00) → revive(10:00) must stay ALIVE even
// when later merged with a stale 09:00 tombstone from another device. The revive
// event time is preserved in revived_at; created_at stays the earliest (sort key).
func TestMergeRevivePersistsAgainstStaleTombstone(t *testing.T) {
	add := item(id64a, "link", "2026-06-10T08:00:00Z")
	del := add
	del.DeletedAt = "2026-06-10T09:00:00Z"

	// Local device: revive at 10:00 (created_at kept at 08:00, revived_at = 10:00).
	revive := add
	revive.RevivedAt = "2026-06-10T10:00:00Z"
	local := mergeFavorites([]FavoriteItem{add, del}, []FavoriteItem{revive})
	if countAlive(local) != 1 {
		t.Fatalf("revive after delete must be alive, got %d alive", countAlive(local))
	}
	if local[0].CreatedAt != "2026-06-10T08:00:00Z" {
		t.Fatalf("created_at must stay the earliest (sort key), got %q", local[0].CreatedAt)
	}
	if local[0].RevivedAt != "2026-06-10T10:00:00Z" {
		t.Fatalf("revived_at must be preserved, got %q", local[0].RevivedAt)
	}

	// Another device still holds only the 09:00 tombstone. Re-merging must NOT
	// resurrect the deletion — the 10:00 revive wins.
	final := mergeFavorites(local, []FavoriteItem{del})
	if final[0].DeletedAt != "" {
		t.Fatalf("stale tombstone re-killed a revived item: deleted_at=%q", final[0].DeletedAt)
	}
	if countAlive(final) != 1 {
		t.Fatal("item revived after the tombstone must remain alive across re-merges")
	}
}

func TestMergeSortsNewestFirst(t *testing.T) {
	merged := mergeFavorites(nil, []FavoriteItem{
		item(id64a, "link", "2026-06-10T08:00:00Z"),
		item(id64b, "link", "2026-06-10T12:00:00Z"),
	})
	if merged[0].CreatedAt < merged[1].CreatedAt {
		t.Fatal("merge output must be sorted newest-first")
	}
}

func TestRenderFavoritesMarkdownGroupsAndSkipsTombstones(t *testing.T) {
	items := []FavoriteItem{
		{ID: id64a, Kind: "link", Title: "Cursor 新功能", Site: "github.com", URL: "https://github.com/x/y",
			CreatedAt: "2026-06-10T15:42:00Z", Source: FavoriteSource{Date: "2026-06-10", Report: "ai-trending", Label: "GitHub AI 趋势"}},
		{ID: id64b, Kind: "excerpt", Title: "Provider 兼容性", Excerpt: "关于 provider 兼容性的分析", Section: "横向对比",
			CreatedAt: "2026-06-09T10:00:00Z", Source: FavoriteSource{Date: "2026-06-09", Report: "ai-agents", Label: "AI Agents 生态", Anchor: id64a}},
		{ID: "ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff", Kind: "report", Title: "Tombstoned",
			CreatedAt: "2026-06-10T01:00:00Z", DeletedAt: "2026-06-10T02:00:00Z", Source: FavoriteSource{Date: "2026-06-10", Report: "ai-ph"}},
	}
	md := renderFavoritesMarkdown(items, "https://ivo-eu.github.io/agents-radar")
	if strings.Contains(md, "Tombstoned") {
		t.Fatal("tombstoned items must be excluded from markdown")
	}
	for _, want := range []string{
		"# AI 情报收藏",
		"## 2026-06-10",
		"### [链接] Cursor 新功能",
		"- 网站：github.com",
		"## 2026-06-09",
		"### [段落] Provider 兼容性",
		"?focus=" + id64a, // excerpt jump-back uses the block anchor
		"https://github.com/x/y",
		// review needs: each favorite also links into the vault report (Obsidian wikilink),
		"- 在库内打开：[[日报/AI情报/2026-06-10/ai-trending]]",
		"- 在库内打开：[[日报/AI情报/2026-06-09/ai-agents#横向对比]]", // excerpt carries a section anchor
	} {
		if !strings.Contains(md, want) {
			t.Fatalf("markdown missing %q\n---\n%s", want, md)
		}
	}
	// newest day must come first
	if strings.Index(md, "## 2026-06-10") > strings.Index(md, "## 2026-06-09") {
		t.Fatal("days must be newest-first")
	}
}

func TestVaultReportLink(t *testing.T) {
	link := vaultReportLink(FavoriteItem{Kind: "link", Source: FavoriteSource{Date: "2026-06-10", Report: "ai-cli"}})
	if link != "[[日报/AI情报/2026-06-10/ai-cli]]" {
		t.Fatalf("link favorite wikilink wrong: %q", link)
	}
	// excerpt with a section gets a heading anchor
	ex := vaultReportLink(FavoriteItem{Kind: "excerpt", Section: "横向对比", Source: FavoriteSource{Date: "2026-06-09", Report: "ai-agents"}})
	if ex != "[[日报/AI情报/2026-06-09/ai-agents#横向对比]]" {
		t.Fatalf("excerpt wikilink with anchor wrong: %q", ex)
	}
	// excerpt without a section: no anchor
	noSec := vaultReportLink(FavoriteItem{Kind: "excerpt", Source: FavoriteSource{Date: "2026-06-09", Report: "ai-agents"}})
	if noSec != "[[日报/AI情报/2026-06-09/ai-agents]]" {
		t.Fatalf("excerpt wikilink without section wrong: %q", noSec)
	}
	// section with wikilink-breaking characters is sanitized out
	dirty := vaultReportLink(FavoriteItem{Kind: "excerpt", Section: "A [b] | c # d", Source: FavoriteSource{Date: "2026-06-09", Report: "ai-agents"}})
	if dirty != "[[日报/AI情报/2026-06-09/ai-agents#A b c d]]" {
		t.Fatalf("section anchor not sanitized: %q", dirty)
	}
	// incomplete source → no link
	if got := vaultReportLink(FavoriteItem{Kind: "link", Source: FavoriteSource{Date: "2026-06-10"}}); got != "" {
		t.Fatalf("incomplete source should yield no link, got %q", got)
	}
}

// ── git fixture: bare origin + Bridge-owned clone ──

func setupRepoFixture(t *testing.T) (app *App, clonePath, vault string) {
	t.Helper()
	root := t.TempDir()
	origin := filepath.Join(root, "origin.git")
	clonePath = filepath.Join(root, "clone")
	vault = filepath.Join(root, "vault")
	if err := os.MkdirAll(vault, 0o755); err != nil {
		t.Fatal(err)
	}
	mustGit(t, "", "init", "--bare", "-b", "master", origin)
	mustGit(t, "", "clone", origin, clonePath)
	mustGit(t, clonePath, "config", "user.email", "bridge@test.local")
	mustGit(t, clonePath, "config", "user.name", "Radar Bridge Test")
	mustGit(t, clonePath, "config", "commit.gpgsign", "false")
	// seed origin/master with an initial commit so fetch master works
	if err := os.WriteFile(filepath.Join(clonePath, "README.md"), []byte("seed\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	mustGit(t, clonePath, "checkout", "-b", "master")
	mustGit(t, clonePath, "add", "README.md")
	mustGit(t, clonePath, "commit", "-m", "seed")
	mustGit(t, clonePath, "push", "-u", "origin", "master")
	// Mark as a Bridge-owned clone so destructive alignment (reset --hard) is allowed.
	if err := markRepoOwned(clonePath); err != nil {
		t.Fatal(err)
	}

	app = &App{config: Config{
		RepoPath:  clonePath,
		VaultPath: vault,
		PagesURL:  "https://ivo-eu.github.io/agents-radar",
	}}
	return app, clonePath, vault
}

func mustGit(t *testing.T, dir string, args ...string) {
	t.Helper()
	if out, err := runGit(dir, args...); err != nil {
		t.Fatalf("git %s: %v: %s", strings.Join(args, " "), err, out)
	}
}

func TestSyncFavoritesWritesCommitsAndPushes(t *testing.T) {
	app, clone, vault := setupRepoFixture(t)
	incoming := []FavoriteItem{
		item(id64a, "link", "2026-06-10T15:42:00Z"),
		item(id64b, "excerpt", "2026-06-10T15:31:00Z"),
	}
	res := app.syncFavorites(incoming)
	if res.err != nil {
		t.Fatalf("sync failed: %v", res.err)
	}
	if !res.pushed {
		t.Fatal("expected push to succeed against bare origin")
	}
	if !res.vaultWritten {
		t.Fatal("expected vault markdown to be written")
	}
	if len(res.items) != 2 {
		t.Fatalf("expected 2 merged items, got %d", len(res.items))
	}
	// favorites.json exists in the clone working tree
	if _, err := os.Stat(filepath.Join(clone, favoritesFileName)); err != nil {
		t.Fatalf("favorites.json not written: %v", err)
	}
	// committed to origin/master
	out, err := runGit(clone, "show", "origin/master:"+favoritesFileName)
	if err != nil || !strings.Contains(out, id64a) {
		t.Fatalf("favorites.json not pushed to origin: %v\n%s", err, out)
	}
	// Obsidian md written to the vault (under the digest tree)
	if _, err := os.Stat(favMdPath(vault)); err != nil {
		t.Fatalf("收藏 md not written: %v", err)
	}
}

func TestSyncFavoritesRefusesDirtyWorkingTree(t *testing.T) {
	app, clone, _ := setupRepoFixture(t)
	if err := os.WriteFile(filepath.Join(clone, "unrelated.txt"), []byte("x"), 0o644); err != nil {
		t.Fatal(err)
	}
	res := app.syncFavorites([]FavoriteItem{item(id64a, "link", "2026-06-10T15:42:00Z")})
	if res.err == nil {
		t.Fatal("expected refusal on dirty working tree")
	}
	if res.pushed {
		t.Fatal("must not push when refusing")
	}
}

func TestSyncFavoritesOfflineWritesLocalNoPush(t *testing.T) {
	app, clone, vault := setupRepoFixture(t)
	// Point origin at a path that does not exist → fetch fails (offline).
	mustGit(t, clone, "remote", "set-url", "origin", filepath.Join(t.TempDir(), "missing.git"))
	res := app.syncFavorites([]FavoriteItem{item(id64a, "link", "2026-06-10T15:42:00Z")})
	if res.pushed {
		t.Fatal("offline sync must not report pushed")
	}
	if res.err == nil || !strings.Contains(res.err.Error(), "离线") {
		t.Fatalf("expected offline error, got %v", res.err)
	}
	if len(res.items) != 1 {
		t.Fatalf("offline sync should still merge locally, got %d", len(res.items))
	}
	// local favorites.json + md still written
	if _, err := os.Stat(filepath.Join(clone, favoritesFileName)); err != nil {
		t.Fatalf("offline: favorites.json not written: %v", err)
	}
	if _, err := os.Stat(favMdPath(vault)); err != nil {
		t.Fatalf("offline: md not written: %v", err)
	}
}

// Review §5.4: a corrupt remote favorites.json must FAIL CLOSED — abort the sync
// rather than treat it as empty and overwrite the library with just the incoming.
func TestSyncFavoritesFailsClosedOnCorruptBaseline(t *testing.T) {
	app, clone, _ := setupRepoFixture(t)
	// Push a malformed favorites.json to origin/master.
	if err := os.WriteFile(filepath.Join(clone, favoritesFileName), []byte("{ not json"), 0o644); err != nil {
		t.Fatal(err)
	}
	mustGit(t, clone, "add", favoritesFileName)
	mustGit(t, clone, "commit", "-m", "corrupt")
	mustGit(t, clone, "push", "origin", "master")

	res := app.syncFavorites([]FavoriteItem{item(id64a, "link", "2026-06-10T15:42:00Z")})
	if res.err == nil {
		t.Fatal("expected sync to fail closed on corrupt baseline")
	}
	if res.pushed {
		t.Fatal("must not push when baseline is unreadable")
	}
	// origin/master must still hold the corrupt file, NOT an overwrite.
	out, _ := runGit(clone, "show", "origin/master:"+favoritesFileName)
	if strings.Contains(out, id64a) {
		t.Fatal("corrupt baseline was overwritten — fail-closed violated")
	}
}

func TestParseFavoritesEmptyVsMalformed(t *testing.T) {
	if items, err := parseFavorites([]byte("   \n")); err != nil || len(items) != 0 {
		t.Fatalf("empty file should be an empty library, got %v err=%v", items, err)
	}
	if _, err := parseFavorites([]byte("{oops")); err == nil {
		t.Fatal("malformed JSON must return an error")
	}
}

// Review §5.5: when the vault markdown cannot be written (e.g. read-only path),
// the sync still pushes to GitHub but reports vault_written=false.
func TestSyncFavoritesReportsVaultWriteFailure(t *testing.T) {
	app, _, vault := setupRepoFixture(t)
	blockFavMarkdownWrite(t, vault)
	res := app.syncFavorites([]FavoriteItem{item(id64a, "link", "2026-06-10T15:42:00Z")})
	if res.err != nil {
		t.Fatalf("git push should still succeed: %v", res.err)
	}
	if !res.pushed {
		t.Fatal("expected push to succeed even when vault write fails")
	}
	if res.vaultWritten {
		t.Fatal("vault_written must be false when the markdown write fails")
	}
	// Review §10.2: the vault error text must be carried out (not swallowed) so the
	// handler can persist it and the front-end can show + retry it.
	if res.vaultErr == "" {
		t.Fatal("vaultErr must carry the markdown write error text")
	}
}

// Review §10.2: when vault write fails but GitHub push succeeds, the handler must
// PERSIST the favorites error (not clear it) so status() keeps reporting it and the
// website keeps the sync pending until a retry heals SecondBrain.
func TestHandleFavoritesSyncPersistsVaultError(t *testing.T) {
	app, _, vault := setupRepoFixture(t)
	app.statePath = filepath.Join(t.TempDir(), "state.json")
	blockFavMarkdownWrite(t, vault)
	body, _ := json.Marshal(map[string]any{"items": []FavoriteItem{item(id64a, "link", "2026-06-10T15:42:00Z")}})
	req := httptest.NewRequest(http.MethodPost, "/api/favorites/sync", strings.NewReader(string(body)))
	rec := httptest.NewRecorder()
	app.handleFavoritesSync(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", rec.Code)
	}
	var resp map[string]any
	if err := json.Unmarshal(rec.Body.Bytes(), &resp); err != nil {
		t.Fatal(err)
	}
	if resp["pushed"] != true {
		t.Fatalf("expected pushed=true, got %v", resp["pushed"])
	}
	if resp["vault_written"] != false {
		t.Fatalf("expected vault_written=false, got %v", resp["vault_written"])
	}
	if ve, _ := resp["vault_error"].(string); !strings.Contains(ve, "SecondBrain 写入失败") {
		t.Fatalf("response must include vault_error, got %q", ve)
	}
	// The persisted state error must NOT be cleared on a vault failure.
	if !strings.Contains(app.state.FavoritesError, "SecondBrain 写入失败") {
		t.Fatalf("FavoritesError must persist the vault failure, got %q", app.state.FavoritesError)
	}
}

func TestSameGitURL(t *testing.T) {
	cases := []struct {
		a, b string
		want bool
	}{
		{"https://github.com/ivo-eu/agents-radar.git", "https://github.com/ivo-eu/agents-radar", true},
		{"https://github.com/ivo-eu/agents-radar/", "https://github.com/ivo-eu/agents-radar.git", true},
		{" https://github.com/ivo-eu/agents-radar.git\n", "https://github.com/ivo-eu/agents-radar", true},
		{"https://github.com/ivo-eu/agents-radar", "https://github.com/someone/fork", false},
		{"git@github.com:ivo-eu/agents-radar.git", "https://github.com/ivo-eu/agents-radar", false},
	}
	for _, c := range cases {
		if got := sameGitURL(c.a, c.b); got != c.want {
			t.Fatalf("sameGitURL(%q,%q)=%v want %v", c.a, c.b, got, c.want)
		}
	}
}

// Review §10.3: a second install --repo-url pointing at a DIFFERENT repo must not
// be silently ignored. On a Bridge-owned clone the origin is updated; on a clone we
// cannot confirm we own, install refuses rather than rewrite the user's remote.
func TestResolveRepoPathOriginMismatch(t *testing.T) {
	// Build two bare origins to act as "old" and "new" repo URLs.
	root := t.TempDir()
	oldOrigin := filepath.Join(root, "old.git")
	newOrigin := filepath.Join(root, "new.git")
	mustGit(t, "", "init", "--bare", "-b", "master", oldOrigin)
	mustGit(t, "", "init", "--bare", "-b", "master", newOrigin)

	// Case A: existing OWNED clone, new URL differs → origin gets updated.
	supportA := t.TempDir()
	dest := filepath.Join(supportA, "agents-radar")
	mustGit(t, "", "clone", oldOrigin, dest)
	if err := markRepoOwned(dest); err != nil {
		t.Fatal(err)
	}
	got, err := resolveRepoPath("", newOrigin, supportA)
	if err != nil {
		t.Fatalf("owned clone should adopt the new origin, got %v", err)
	}
	if got != dest {
		t.Fatalf("expected %s, got %s", dest, got)
	}
	origin, _ := runGit(dest, "remote", "get-url", "origin")
	if !sameGitURL(origin, newOrigin) {
		t.Fatalf("origin should be updated to %s, got %s", newOrigin, strings.TrimSpace(origin))
	}

	// Case B: existing UNOWNED clone, new URL differs → refuse (do not rewrite).
	supportB := t.TempDir()
	destB := filepath.Join(supportB, "agents-radar")
	mustGit(t, "", "clone", oldOrigin, destB)
	// no ownership marker
	if _, err := resolveRepoPath("", newOrigin, supportB); err == nil {
		t.Fatal("must refuse to adopt a new origin on an unowned clone")
	}
	originB, _ := runGit(destB, "remote", "get-url", "origin")
	if !sameGitURL(originB, oldOrigin) {
		t.Fatalf("unowned clone origin must be left untouched, got %s", strings.TrimSpace(originB))
	}

	// Case C: existing owned clone, SAME URL → reused without error.
	supportC := t.TempDir()
	destC := filepath.Join(supportC, "agents-radar")
	mustGit(t, "", "clone", oldOrigin, destC)
	if err := markRepoOwned(destC); err != nil {
		t.Fatal(err)
	}
	if _, err := resolveRepoPath("", oldOrigin, supportC); err != nil {
		t.Fatalf("same origin should reuse cleanly, got %v", err)
	}
}

// Review §5.3: an UNOWNED repo (user pointed --repo at their own clone) must
// never be hard-reset. A feature branch or unpushed commits → refuse, no damage.
func TestSyncRefusesHardResetOnUnownedRepo(t *testing.T) {
	app, clone, _ := setupRepoFixture(t)
	// Drop the ownership marker → behaves like a user-supplied --repo.
	if err := os.Remove(ownershipMarkerPath(clone)); err != nil {
		t.Fatal(err)
	}
	// Case A: on a feature branch.
	mustGit(t, clone, "checkout", "-b", "feature/work")
	res := app.syncFavorites([]FavoriteItem{item(id64a, "link", "2026-06-10T15:42:00Z")})
	if res.err == nil || res.pushed {
		t.Fatalf("must refuse sync on a non-master unowned repo, got err=%v pushed=%v", res.err, res.pushed)
	}
	if !strings.Contains(res.err.Error(), "master") {
		t.Fatalf("error should explain the branch guard, got %v", res.err)
	}

	// Case B: back on master but with an unpushed local commit (ahead of origin).
	mustGit(t, clone, "checkout", "master")
	if err := os.WriteFile(filepath.Join(clone, "local.txt"), []byte("mine\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	mustGit(t, clone, "add", "local.txt")
	mustGit(t, clone, "commit", "-m", "unpushed local work")
	res = app.syncFavorites([]FavoriteItem{item(id64a, "link", "2026-06-10T15:42:00Z")})
	if res.err == nil || res.pushed {
		t.Fatalf("must refuse to hard-reset a repo with unpushed commits, got err=%v", res.err)
	}
	// The local commit must still be there — nothing was reset.
	if out, _ := runGit(clone, "log", "-1", "--pretty=%s"); !strings.Contains(out, "unpushed local work") {
		t.Fatalf("unowned repo lost its local commit: %s", out)
	}
}

// Review §5.7: markdown groups by LOCAL calendar date, not UTC.
func TestFavDayUsesLocalDate(t *testing.T) {
	// 2026-06-10T16:30:00Z is 2026-06-11 00:30 in UTC+8. In an UTC+8 zone the
	// favorite must group under 06-11; the regression was grouping under 06-10.
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Skipf("tz data unavailable: %v", err)
	}
	orig := time.Local
	time.Local = loc
	defer func() { time.Local = orig }()
	if got := favDay("2026-06-10T16:30:00Z"); got != "2026-06-11" {
		t.Fatalf("expected local date 2026-06-11, got %s", got)
	}
}

// Review §5.6: install deploys the running binary to a stable bin/ path so the
// LaunchAgent never points at a transient build product.
func TestDeployStableBinary(t *testing.T) {
	support := t.TempDir()
	dest, err := deployStableBinary(support)
	if err != nil {
		t.Fatalf("deploy failed: %v", err)
	}
	want := filepath.Join(support, "bin", "radar-bridge")
	if dest != want {
		t.Fatalf("expected stable path %s, got %s", want, dest)
	}
	info, err := os.Stat(dest)
	if err != nil {
		t.Fatalf("stable binary not written: %v", err)
	}
	if info.Mode().Perm()&0o100 == 0 {
		t.Fatalf("stable binary is not executable: %v", info.Mode())
	}
}

func TestFavoritesJSONRoundTripsRevivedAt(t *testing.T) {
	in := FavoritesFile{Version: 1, Items: []FavoriteItem{{
		ID: id64a, Kind: "link", CreatedAt: "2026-06-10T08:00:00Z",
		RevivedAt: "2026-06-10T10:00:00Z", Source: FavoriteSource{Date: "2026-06-10", Report: "ai-hn"},
	}}}
	data, err := json.Marshal(in)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(data), `"revived_at":"2026-06-10T10:00:00Z"`) {
		t.Fatalf("revived_at must serialize: %s", data)
	}
	items, err := parseFavorites(data)
	if err != nil || len(items) != 1 || items[0].RevivedAt != "2026-06-10T10:00:00Z" {
		t.Fatalf("revived_at must round-trip, got %+v err=%v", items, err)
	}
}

// Review §11.2: adopting a fresh disk state in the digest-sync path must refresh
// only digest-owned fields and PRESERVE the in-memory favorites fields.
func TestAdoptDigestStateKeepsFavoritesFields(t *testing.T) {
	app := &App{}
	// In-memory: favorites sync just recorded a vault failure + a fresh sync time.
	app.state = State{
		Token:             "in-memory-token",
		FavoritesError:    "SecondBrain 写入失败：read-only",
		FavoritesLastSync: "2026-06-11T09:00:00Z",
	}
	// Disk snapshot is older: digest fields populated, favorites fields empty.
	disk := State{
		Token:     "disk-token",
		LastSync:  "2026-06-11T08:00:00Z",
		LastError: "prev digest error",
		Files:     map[string]string{"a.md": "sha"},
		// FavoritesError / FavoritesLastSync intentionally empty (stale).
	}
	app.adoptDigestStateLocked(disk)

	// Digest-owned fields adopted from disk.
	if app.state.Token != "disk-token" || app.state.LastSync != "2026-06-11T08:00:00Z" ||
		app.state.LastError != "prev digest error" || app.state.Files["a.md"] != "sha" {
		t.Fatalf("digest fields not adopted from disk: %+v", app.state)
	}
	// Favorites-owned fields preserved from memory (NOT clobbered by the stale disk).
	if app.state.FavoritesError != "SecondBrain 写入失败：read-only" {
		t.Fatalf("FavoritesError was clobbered: %q", app.state.FavoritesError)
	}
	if app.state.FavoritesLastSync != "2026-06-11T09:00:00Z" {
		t.Fatalf("FavoritesLastSync was clobbered: %q", app.state.FavoritesLastSync)
	}
}

// Review §11.4: simulate the interleaving where a digest sync reads the old disk
// state, a favorites sync then persists a vault error, and the digest sync later
// adopts its (stale) snapshot + saves. The favorites error must survive in BOTH
// memory and state.json — no stale-snapshot overwrite.
func TestDigestSyncDoesNotClobberFavoritesError(t *testing.T) {
	dir := t.TempDir()
	app := &App{statePath: filepath.Join(dir, "state.json")}
	app.state = State{Token: "t", Files: map[string]string{}}
	if err := app.saveState(); err != nil {
		t.Fatal(err)
	}

	// 1) Digest sync reads the (old, favorites-empty) disk snapshot.
	var diskSnapshot State
	if err := readJSON(app.statePath, &diskSnapshot); err != nil {
		t.Fatal(err)
	}

	// 2) Favorites sync records a vault failure to memory + disk.
	app.mu.Lock()
	app.state.FavoritesError = "SecondBrain 写入失败：read-only"
	app.state.FavoritesLastSync = "2026-06-11T09:00:00Z"
	app.mu.Unlock()
	if err := app.saveState(); err != nil {
		t.Fatal(err)
	}

	// 3) Digest sync now adopts its stale snapshot and saves (the dangerous step).
	app.mu.Lock()
	app.adoptDigestStateLocked(diskSnapshot)
	app.state.LastSync = "2026-06-11T09:01:00Z"
	app.state.LastError = ""
	app.mu.Unlock()
	if err := app.saveState(); err != nil {
		t.Fatal(err)
	}

	// Memory must still hold the favorites error.
	if app.state.FavoritesError == "" {
		t.Fatal("digest sync clobbered FavoritesError in memory")
	}
	// Disk must still hold it too.
	var onDisk State
	if err := readJSON(app.statePath, &onDisk); err != nil {
		t.Fatal(err)
	}
	if onDisk.FavoritesError == "" {
		t.Fatalf("digest sync clobbered FavoritesError on disk: %+v", onDisk)
	}
	if onDisk.LastSync != "2026-06-11T09:01:00Z" {
		t.Fatalf("digest field should still be saved: %q", onDisk.LastSync)
	}

	// Reverse path: a digest sync holds a stale snapshot that STILL CARRIES the old
	// error; meanwhile favorites sync succeeds and clears it. Adopting the stale
	// snapshot must NOT resurrect the cleared error — in memory or on disk.
	//
	// Capture the stale snapshot WHILE the error is on disk (this is the crux the
	// previous version missed: it reused the always-empty diskSnapshot).
	var staleWithError State
	if err := readJSON(app.statePath, &staleWithError); err != nil {
		t.Fatal(err)
	}
	if staleWithError.FavoritesError == "" {
		t.Fatal("test setup: stale snapshot must carry the old error")
	}

	// Favorites sync succeeds → clears the error in memory + disk.
	app.mu.Lock()
	app.state.FavoritesError = ""
	app.mu.Unlock()
	if err := app.saveState(); err != nil {
		t.Fatal(err)
	}

	// Digest sync adopts its stale (error-bearing) snapshot and saves.
	app.mu.Lock()
	app.adoptDigestStateLocked(staleWithError)
	app.state.LastSync = "2026-06-11T09:02:00Z"
	app.mu.Unlock()
	if err := app.saveState(); err != nil {
		t.Fatal(err)
	}
	if app.state.FavoritesError != "" {
		t.Fatalf("stale digest snapshot resurrected a cleared FavoritesError in memory: %q", app.state.FavoritesError)
	}
	var afterClear State
	if err := readJSON(app.statePath, &afterClear); err != nil {
		t.Fatal(err)
	}
	if afterClear.FavoritesError != "" {
		t.Fatalf("stale digest snapshot resurrected a cleared FavoritesError on disk: %q", afterClear.FavoritesError)
	}
}
