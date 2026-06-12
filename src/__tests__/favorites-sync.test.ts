import { readFileSync } from "node:fs";
import { fileURLToPath } from "node:url";
import { dirname, resolve } from "node:path";
import { describe, it, expect, beforeEach } from "vitest";

// Drives the REAL flushFavSync (and its pure helpers) extracted from index.html
// inside an injectable sandbox, to verify the review §10.2 contract: a vault-write
// failure must NOT advance the sync watermark, must keep the item pending, and a
// later vault success must clear the pending state.

const HERE = dirname(fileURLToPath(import.meta.url));
const INDEX_HTML = resolve(HERE, "../../index.html");

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

interface FavItem {
  id: string;
  kind?: string;
  created_at?: string;
  revived_at?: string;
  deleted_at?: string;
  source?: Record<string, string>;
}
interface SyncResponse {
  items?: FavItem[];
  pushed?: boolean;
  vault_written?: boolean;
  vault_error?: string;
  error?: string;
  alive?: number;
}

// Build a sandbox whose globals back the extracted functions. bridgeFetch returns
// queued responses; setTimeout is captured (not real) so we can fire retries by hand.
function makeHarness(initialItems: FavItem[], responses: SyncResponse[]) {
  const html = readFileSync(INDEX_HTML, "utf8");
  const pureNames = [
    "favPreferNonEmpty",
    "favEarliest",
    "favLatest",
    "mergeSource",
    "mergeItem",
    "aliveItems",
    "favLastActivity",
    "pendingFavCount",
    "reconcileWithServer",
    "flushFavSync",
  ];
  const body = pureNames.map((n) => grabFunction(html, n)).join("\n\n");

  const store = {
    favoriteStore: { version: 1, updated_at: "", items: initialItems.map((it) => ({ ...it })) },
  };
  const localData = new Map<string, string>();
  let fetchCalls = 0;
  let pendingTimer: (() => void) | null = null;

  const env = {
    favSyncing: false,
    bridgeConnected: true,
    favSyncedAt: "",
    favSyncError: "",
    favSyncTimer: 0 as unknown,
    favoriteIds: new Set<string>(),
    currentFavFilter: null as string | null,
    FAV_KEY: "radar-favorites",
    FAV_SYNCED_KEY: "radar-favorites-synced-at",
    FAV_SYNC_DEBOUNCE: 5000,
    alerts: [] as string[],
    fetchCount: () => fetchCalls,
    fireRetry: async () => {
      const fn = pendingTimer;
      pendingTimer = null;
      if (fn) await fn();
    },
    localStorage: {
      setItem: (k: string, v: string) => void localData.set(k, v),
      getItem: (k: string) => localData.get(k) ?? null,
    },
    setTimeout: (fn: () => void) => {
      pendingTimer = fn;
      return 1;
    },
    clearTimeout: () => {
      pendingTimer = null;
    },
    updateBridgeButtons: () => {},
    renderFavNav: () => {},
    renderFavoritesView: () => {},
    alert: (_m: string) => {},
    bridgeFetch: async () => {
      const r = responses[Math.min(fetchCalls, responses.length - 1)];
      fetchCalls += 1;
      return r;
    },
  };
  env.alert = (m: string) => env.alerts.push(m);

  // Run the extracted code under a scope proxy so reads/writes of favSyncedAt,
  // favSyncError, favoriteStore, … land on our injectable env (see runWithProxy).
  return { env, store, run: () => runWithProxy(body, env as Record<string, unknown>, store) };
}

// Execute the extracted function bodies with a scope proxy so that both reads and
// assignments to favSyncedAt/favSyncError/etc. land on `env`, and favoriteStore on store.
function runWithProxy(body: string, env: Record<string, unknown>, store: { favoriteStore: unknown }) {
  const scope = new Proxy(
    {},
    {
      // with() consults `has` for EVERY identifier; returning true routes them all
      // here. For names we inject, serve env/store; otherwise fall back to the real
      // global (JSON, Array, Set, …) so the extracted code still works.
      has: () => true,
      get: (_t, prop: string | symbol) => {
        if (typeof prop === "symbol") return (globalThis as Record<symbol, unknown>)[prop];
        if (prop === "favoriteStore") return store.favoriteStore;
        if (prop in env) return env[prop];
        return (globalThis as Record<string, unknown>)[prop];
      },
      set: (_t, prop: string | symbol, value) => {
        if (prop === "favoriteStore") store.favoriteStore = value;
        else env[prop as string] = value;
        return true;
      },
    },
  );
  const fn = new Function("scope", `with(scope){ ${body}\n return { flushFavSync, pendingFavCount }; }`);
  return fn(scope) as {
    flushFavSync: (showResult?: boolean) => Promise<void>;
    pendingFavCount: () => number;
  };
}

const src = { date: "2026-06-10", report: "ai-hn" };
const A = "a".repeat(64);

describe("vault-write failure keeps the sync pending and self-heals (review §10.2)", () => {
  let harness: ReturnType<typeof makeHarness>;
  let api: { flushFavSync: (s?: boolean) => Promise<void>; pendingFavCount: () => number };

  beforeEach(() => {
    harness = makeHarness(
      [{ id: A, kind: "link", created_at: "2026-06-10T08:00:00Z", source: src }],
      [
        // 1st sync: GitHub pushed, but SecondBrain markdown failed.
        {
          items: [{ id: A, kind: "link", created_at: "2026-06-10T08:00:00Z", source: src }],
          pushed: true,
          vault_written: false,
          vault_error: "SecondBrain 写入失败：read-only",
        },
        // 2nd sync (retry): both succeed.
        {
          items: [{ id: A, kind: "link", created_at: "2026-06-10T08:00:00Z", source: src }],
          pushed: true,
          vault_written: true,
        },
      ],
    );
    api = harness.run();
  });

  it("does not advance the watermark or clear pending when vault fails", async () => {
    expect(api.pendingFavCount()).toBe(1);
    await api.flushFavSync(false);
    // Watermark NOT advanced → still pending.
    expect(harness.env.favSyncedAt).toBe("");
    expect(api.pendingFavCount()).toBe(1);
    expect(String(harness.env.favSyncError)).toContain("SecondBrain 写入失败");
  });

  it("retries and clears pending only after the vault write succeeds", async () => {
    await api.flushFavSync(false); // attempt 1 → vault fails, schedules retry
    expect(harness.env.fetchCount()).toBe(1);
    await harness.env.fireRetry(); // attempt 2 → both succeed
    expect(harness.env.fetchCount()).toBe(2);
    expect(api.pendingFavCount()).toBe(0);
    expect(harness.env.favSyncError).toBe("");
    expect(harness.env.favSyncedAt).not.toBe("");
  });
});

describe("clean sync advances the watermark and clears pending", () => {
  it("a fully successful sync leaves nothing pending", async () => {
    const h = makeHarness(
      [{ id: A, kind: "link", created_at: "2026-06-10T08:00:00Z", source: src }],
      [
        {
          items: [{ id: A, kind: "link", created_at: "2026-06-10T08:00:00Z", source: src }],
          pushed: true,
          vault_written: true,
        },
      ],
    );
    const api = h.run();
    expect(api.pendingFavCount()).toBe(1);
    await api.flushFavSync(false);
    expect(api.pendingFavCount()).toBe(0);
    expect(h.env.favSyncError).toBe("");
  });
});
