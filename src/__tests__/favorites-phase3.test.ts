import { readFileSync } from "node:fs";
import { fileURLToPath } from "node:url";
import { dirname, resolve } from "node:path";
import { describe, it, expect, beforeAll } from "vitest";

// Loads the real Phase 3 favorites code out of index.html (filters, tag parsing,
// tags/note LWW) and checks the behaviours the Phase 3 review (P2-3.1/3.3/3.4,
// P3-3.5) said were missing coverage. The extracted module declares the filter
// state `let`s and a setFilters() so favItemMatches can be exercised in isolation.

const HERE = dirname(fileURLToPath(import.meta.url));
const INDEX_HTML = resolve(HERE, "../../index.html");

interface FavItem {
  id: string;
  kind?: string;
  title?: string;
  url?: string;
  site?: string;
  excerpt?: string;
  section?: string;
  note?: string;
  tags?: string[];
  created_at?: string;
  revived_at?: string;
  edited_at?: string;
  deleted_at?: string;
  source?: Record<string, string>;
}

function grabFunction(source: string, name: string): string {
  const re = new RegExp(`(async\\s+)?function ${name}\\s*\\([^)]*\\)\\s*\\{`);
  const m = re.exec(source);
  if (!m) throw new Error(`function not found in index.html: ${name}`);
  let depth = 0;
  let i = m.index + m[0].length - 1;
  for (; i < source.length; i++) {
    if (source[i] === "{") depth++;
    else if (source[i] === "}") {
      depth--;
      if (depth === 0) {
        i++;
        break;
      }
    }
  }
  return source.slice(m.index, i);
}

let parseTagsInput: (s: string) => string[];
let mergeItem: (a: FavItem, b: FavItem) => FavItem;
let favLastActivity: (it: FavItem) => string;
let favItemMatches: (it: FavItem) => boolean;
let setFilters: (s: { q?: string; date?: string; tags?: string[] }) => void;

beforeAll(async () => {
  const html = readFileSync(INDEX_HTML, "utf8");
  const fns = [
    "favPreferNonEmpty",
    "favEarliest",
    "favLatest",
    "favTagsNoteKey",
    "favKeyGreaterEq",
    "mergeSource",
    "mergeItem",
    "parseTagsInput",
    "favLastActivity",
    "favItemMatches",
  ];
  const src =
    "const favTextEncoder = new TextEncoder();\n" + // favKeyGreaterEq depends on this module-level const
    "let favSearchQuery='', favSourceDateFilter='', favTagFilter=new Set();\n" +
    "function setFilters(s){favSearchQuery=s.q||'';favSourceDateFilter=s.date||'';favTagFilter=new Set(s.tags||[]);}\n" +
    fns.map((n) => grabFunction(html, n)).join("\n\n") +
    "\nexport { parseTagsInput, mergeItem, favLastActivity, favItemMatches, setFilters };";
  const mod = await import("data:text/javascript," + encodeURIComponent(src));
  parseTagsInput = mod.parseTagsInput;
  mergeItem = mod.mergeItem;
  favLastActivity = mod.favLastActivity;
  favItemMatches = mod.favItemMatches;
  setFilters = mod.setFilters;
});

const src = { date: "2026-06-10", report: "ai-hn", label: "HN 社区动态" };
const item = (over: Partial<FavItem>): FavItem => ({
  id: "a".repeat(64),
  kind: "link",
  source: src,
  ...over,
});

describe("parseTagsInput", () => {
  it("splits, strips '#', trims, dedupes", () => {
    expect(parseTagsInput("#agent, b  c #agent")).toEqual(["agent", "b", "c"]);
  });
  it("caps at 20 tags", () => {
    const many = Array.from({ length: 30 }, (_, i) => `t${i}`).join(" ");
    expect(parseTagsInput(many)).toHaveLength(20);
  });
  it("truncates a tag by Unicode code point (32)", () => {
    const tag = "😀".repeat(40); // 40 code points
    expect([...(parseTagsInput(tag)[0] ?? "")]).toHaveLength(32);
  });
});

describe("favItemMatches: search + date + tag (AND)", () => {
  it("search matches title/excerpt/site/note/tag", () => {
    const it = item({
      title: "Cursor",
      excerpt: "agent memory",
      site: "github.com",
      note: "回看",
      tags: ["cli"],
    });
    setFilters({ q: "cursor" });
    expect(favItemMatches(it)).toBe(true);
    setFilters({ q: "memory" });
    expect(favItemMatches(it)).toBe(true);
    setFilters({ q: "github" });
    expect(favItemMatches(it)).toBe(true);
    setFilters({ q: "回看" });
    expect(favItemMatches(it)).toBe(true);
    setFilters({ q: "cli" });
    expect(favItemMatches(it)).toBe(true);
    setFilters({ q: "nope" });
    expect(favItemMatches(it)).toBe(false);
  });
  it("date filter is exact match on source.date", () => {
    const it = item({ source: { ...src } });
    setFilters({ date: "2026-06-10" });
    expect(favItemMatches(it)).toBe(true);
    setFilters({ date: "2026-06-09" });
    expect(favItemMatches(it)).toBe(false);
  });
  it("tag filter requires ALL selected tags (AND)", () => {
    const it = item({ tags: ["agent", "cli"] });
    setFilters({ tags: ["agent"] });
    expect(favItemMatches(it)).toBe(true);
    setFilters({ tags: ["agent", "cli"] });
    expect(favItemMatches(it)).toBe(true);
    setFilters({ tags: ["agent", "web"] }); // item lacks "web"
    expect(favItemMatches(it)).toBe(false);
  });
  it("combines all filters with AND", () => {
    const it = item({ title: "Cursor", tags: ["cli"], source: { ...src } });
    setFilters({ q: "cursor", date: "2026-06-10", tags: ["cli"] });
    expect(favItemMatches(it)).toBe(true);
    setFilters({ q: "cursor", date: "2026-06-09", tags: ["cli"] });
    expect(favItemMatches(it)).toBe(false);
  });
});

describe("tags/note LWW", () => {
  it("later edited_at wins the whole tag set + note", () => {
    const a = item({ tags: ["x"], note: "old", edited_at: "2026-06-11T08:00:00Z" });
    const b = item({ tags: ["y", "z"], note: "new", edited_at: "2026-06-11T09:00:00Z" });
    const m = mergeItem(a, b);
    expect(m.tags).toEqual(["y", "z"]);
    expect(m.note).toBe("new");
  });
  it("does not union — a removed tag stays removed", () => {
    const a = item({ tags: ["a", "b"], edited_at: "2026-06-11T08:00:00Z" });
    const b = item({ tags: ["a"], edited_at: "2026-06-11T09:00:00Z" });
    expect(mergeItem(a, b).tags).toEqual(["a"]);
  });
  it("is order-independent on an edited_at tie", () => {
    const a = item({ tags: ["a"], note: "AA", edited_at: "2026-06-11T09:00:00Z" });
    const b = item({ tags: ["b", "c"], note: "BB", edited_at: "2026-06-11T09:00:00Z" });
    const ab = mergeItem(a, b);
    const ba = mergeItem(b, a);
    expect(ab.tags).toEqual(ba.tags);
    expect(ab.note).toBe(ba.note);
  });
  // P2-6.2: tie-break compares UTF-8 BYTES (not UTF-16 code units), so it agrees
  // with the Go Bridge for non-BMP chars. "😀" (F0 9F 98 80) > "￿"/U+FFFF (EF BF BF)
  // in UTF-8; a naive JS >= would pick the other way (D83D < FFFF in UTF-16).
  it("picks the same non-ASCII winner as Go and stays order-independent", () => {
    const ed = "2026-06-15T08:00:00Z";
    const emoji = item({ tags: ["😀"], edited_at: ed });
    const uffff = item({ tags: ["￿"], edited_at: ed });
    expect(mergeItem(emoji, uffff).tags).toEqual(["😀"]);
    expect(mergeItem(uffff, emoji).tags).toEqual(["😀"]); // order-independent
  });
  it("drops empty tags/note/edited_at from the merged item", () => {
    const a = item({ created_at: "2026-06-10T08:00:00Z" });
    const b = item({ created_at: "2026-06-10T09:00:00Z" });
    const m = mergeItem(a, b);
    expect("tags" in m).toBe(false);
    expect("note" in m).toBe(false);
    expect("edited_at" in m).toBe(false);
  });
});

describe("favLastActivity", () => {
  it("includes edited_at so an edit triggers sync", () => {
    const it = item({ created_at: "2026-06-10T08:00:00Z", edited_at: "2026-06-11T10:00:00Z" });
    expect(favLastActivity(it)).toBe("2026-06-11T10:00:00Z");
  });
});
