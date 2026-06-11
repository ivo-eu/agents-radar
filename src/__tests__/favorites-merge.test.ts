import { readFileSync } from "node:fs";
import { fileURLToPath } from "node:url";
import { dirname, resolve } from "node:path";
import { describe, it, expect, beforeAll } from "vitest";

// Loads the real Phase 2 merge engine out of index.html (the website side of the
// favorites sync contract) and checks it against the same rules the Go Bridge
// enforces: earliest created_at, revive-after-delete, and — crucially — that an
// in-flight reconcile never drops changes the user made during a sync (review §5.1)
// and a revive survives a stale tombstone (review §5.2).

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
  created_at?: string;
  revived_at?: string;
  deleted_at?: string;
  source?: Record<string, string>;
}

let mergeItem: (a: FavItem, b: FavItem) => FavItem;

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

beforeAll(async () => {
  const html = readFileSync(INDEX_HTML, "utf8");
  const names = ["favPreferNonEmpty", "favEarliest", "favLatest", "mergeSource", "mergeItem"];
  const src = names.map((n) => grabFunction(html, n)).join("\n\n") + "\nexport { mergeItem };";
  const mod = await import("data:text/javascript," + encodeURIComponent(src));
  mergeItem = mod.mergeItem;
});

// Mirror of the production fold in reconcileWithServer / loadFavorites: union by
// id, mergeItem on collision. Used here to prove the in-flight reconcile keeps
// every change. (reconcileWithServer itself touches localStorage/DOM globals.)
function fold(...lists: FavItem[][]): Map<string, FavItem> {
  const byId = new Map<string, FavItem>();
  for (const list of lists) {
    for (const it of list) {
      if (!it || !it.id) continue;
      byId.set(it.id, byId.has(it.id) ? mergeItem(byId.get(it.id)!, it) : it);
    }
  }
  return byId;
}

const src = { date: "2026-06-10", report: "ai-hn" };
const A = "a".repeat(64);
const B = "b".repeat(64);

describe("mergeItem matches the Go merge rules", () => {
  it("keeps the earliest created_at", () => {
    const out = mergeItem(
      { id: A, created_at: "2026-06-10T09:00:00Z", source: src },
      { id: A, created_at: "2026-06-10T08:00:00Z", source: src },
    );
    expect(out.created_at).toBe("2026-06-10T08:00:00Z");
  });

  it("fills missing fields from the other side", () => {
    const out = mergeItem(
      { id: A, created_at: "2026-06-10T08:00:00Z", source: src },
      { id: A, title: "Cursor", site: "github.com", created_at: "2026-06-10T08:00:00Z", source: { ...src, label: "HN" } },
    );
    expect(out.title).toBe("Cursor");
    expect(out.site).toBe("github.com");
    expect(out.source?.label).toBe("HN");
  });

  it("tombstones when delete is the latest event", () => {
    const out = mergeItem(
      { id: A, created_at: "2026-06-10T08:00:00Z", source: src },
      { id: A, created_at: "2026-06-10T08:00:00Z", deleted_at: "2026-06-10T09:00:00Z", source: src },
    );
    expect(out.deleted_at).toBe("2026-06-10T09:00:00Z");
  });

  it("revive after delete clears the tombstone and survives a stale one", () => {
    const deleted = { id: A, created_at: "2026-06-10T08:00:00Z", deleted_at: "2026-06-10T09:00:00Z", source: src };
    const revived = { id: A, created_at: "2026-06-10T08:00:00Z", revived_at: "2026-06-10T10:00:00Z", source: src };
    const live = mergeItem(deleted, revived);
    expect(live.deleted_at).toBeFalsy();
    expect(live.created_at).toBe("2026-06-10T08:00:00Z"); // sort key unchanged
    expect(live.revived_at).toBe("2026-06-10T10:00:00Z");
    // Re-merging the stale 09:00 tombstone must NOT re-kill it.
    const after = mergeItem(live, deleted);
    expect(after.deleted_at).toBeFalsy();
  });
});

describe("in-flight reconcile keeps changes made during a sync (review §5.1)", () => {
  it("preserves an item added while the request was in flight", () => {
    // Snapshot sent to Bridge held only A; user added B mid-flight.
    const serverReturned: FavItem[] = [{ id: A, kind: "link", created_at: "2026-06-10T08:00:00Z", source: src }];
    const currentStore: FavItem[] = [
      { id: A, kind: "link", created_at: "2026-06-10T08:00:00Z", source: src },
      { id: B, kind: "link", created_at: "2026-06-10T08:05:00Z", source: src }, // in-flight add
    ];
    const merged = fold(serverReturned, currentStore);
    expect(merged.has(B)).toBe(true);
    expect([...merged.keys()].sort()).toEqual([A, B]);
  });

  it("preserves a delete made while the request was in flight", () => {
    // Server still thinks A is alive; user deleted it mid-flight.
    const serverReturned: FavItem[] = [{ id: A, kind: "link", created_at: "2026-06-10T08:00:00Z", source: src }];
    const currentStore: FavItem[] = [
      { id: A, kind: "link", created_at: "2026-06-10T08:00:00Z", deleted_at: "2026-06-10T08:05:00Z", source: src },
    ];
    const merged = fold(serverReturned, currentStore);
    expect(merged.get(A)?.deleted_at).toBe("2026-06-10T08:05:00Z");
  });

  it("is order-independent (fold base-then-current == current-then-base for content)", () => {
    const a: FavItem = { id: A, kind: "link", title: "from server", created_at: "2026-06-10T08:00:00Z", source: src };
    const b: FavItem = { id: A, kind: "link", created_at: "2026-06-10T07:00:00Z", source: src };
    const m1 = mergeItem(a, b);
    const m2 = mergeItem(b, a);
    expect(m1.created_at).toBe(m2.created_at);
    expect(m1.deleted_at ?? "").toBe(m2.deleted_at ?? "");
  });
});
