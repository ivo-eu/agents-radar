import { readFileSync } from "node:fs";
import { fileURLToPath } from "node:url";
import { dirname, resolve } from "node:path";
import { webcrypto } from "node:crypto";
import { describe, it, expect, beforeAll } from "vitest";

// These tests load the real favorite-ID functions out of index.html (the single
// source of truth for the website bookmark system) and assert they produce the
// exact SHA-256 IDs the Go Bridge produces. The expected IDs below were emitted
// by the Go implementation (bridge favoriteID / normalizeFavoriteURL / normalizeText).
// If the JS normalization drifts from Go, these byte-for-byte IDs break — which
// is the whole point: Phase 2 multi-device sync relies on identical IDs.

const HERE = dirname(fileURLToPath(import.meta.url));
const INDEX_HTML = resolve(HERE, "../../index.html");

// crypto.subtle is a secure-context global in browsers; provide it in Node.
if (!(globalThis as { crypto?: unknown }).crypto) {
  (globalThis as { crypto?: unknown }).crypto = webcrypto;
}
// new URL(href, location.href) needs a base in the page; shim it for Node.
(globalThis as { location?: unknown }).location = {
  href: "https://ivo-eu.github.io/agents-radar/",
};

type FavId = (kind: string, parts: Record<string, string>) => Promise<string>;
type Norm = (value: string) => string;

let favId: FavId;
let normalizeUrl: Norm;
let normalizeText: Norm;

// Extract a top-level function declaration (optionally async) by name, balancing braces.
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
  const names = [
    "normalizeUrl",
    "encodeGoQuery",
    "goQueryEscape",
    "normalizeText",
    "sha256Hex",
    "favKeyString",
    "favId",
  ];
  const src =
    names.map((n) => grabFunction(html, n)).join("\n\n") + "\nexport { favId, normalizeUrl, normalizeText };";
  const mod = await import("data:text/javascript," + encodeURIComponent(src));
  favId = mod.favId;
  normalizeUrl = mod.normalizeUrl;
  normalizeText = mod.normalizeText;
});

describe("favorite URL normalization matches Go Bridge", () => {
  const cases: Array<[string, string]> = [
    ["HTTPS://Example.COM/article/?utm_source=radar&b=2&a=1#section", "https://example.com/article/?a=1&b=2"],
    ["HTTP://EXAMPLE.com/", "http://example.com"],
    ["https://example.com/p?ref=x&source=y&keep=1", "https://example.com/p?keep=1"],
    ["https://example.com/s?q=hello world&x=a+b", "https://example.com/s?q=hello+world&x=a+b"],
    ["https://example.com:8443/x?z=1", "https://example.com:8443/x?z=1"],
  ];
  it.each(cases)("normalizes %s", (raw, expected) => {
    expect(normalizeUrl(raw)).toBe(expected);
  });
});

describe("favorite text normalization matches Go Bridge", () => {
  it("collapses whitespace, trims, lowercases", () => {
    expect(normalizeText("Agent   memory\narchitecture")).toBe("agent memory architecture");
    expect(normalizeText("  Provider COMPATIBILITY Notes ")).toBe("provider compatibility notes");
  });
});

describe("favorite ID is identical to the Go Bridge SHA-256", () => {
  // Expected IDs emitted by the Go implementation (see bridge favoriteID).
  it("link ignores tracking params, fragment, query order", async () => {
    expect(await favId("link", { url: "https://example.com/post?utm_source=one&b=2&a=1#frag" })).toBe(
      "c7d086f6048104063f2da1ddc84b454f1a59e5e32c184771b605afae4d30cd7b",
    );
  });
  it("link preserves a non-root trailing slash", async () => {
    expect(
      await favId("link", { url: "https://Example.COM/article/?utm_source=radar&b=2&a=1#section" }),
    ).toBe("96a07bb55d0639622c75f8ba1d63114bc6ddcb985c0cea16841294bda715d331");
  });
  it("link collapses the root path slash", async () => {
    expect(await favId("link", { url: "HTTP://EXAMPLE.com/" })).toBe(
      "2774e215cd358e72cc2709391daa99d7d98b0631b643575d7729a4b46a693811",
    );
  });
  it("link strips ref/source params", async () => {
    expect(await favId("link", { url: "https://example.com/p?ref=x&source=y&keep=1" })).toBe(
      "600663b46a20576ace4737869b8d489244beebddd394df5ac40e47a3813a73d2",
    );
  });
  it("link encodes spaces and plus signs like Go", async () => {
    expect(await favId("link", { url: "https://example.com/s?q=hello world&x=a+b" })).toBe(
      "f6038fd489bb7df8cf07eade211bd039765d071dae15d5b4b58b5abfaed3863c",
    );
  });
  it("link percent-encodes unicode path and query", async () => {
    expect(await favId("link", { url: "https://example.com/路径/page?名=值" })).toBe(
      "ff92cca91ef946be3a939b3b42465f8f45206b5c2037f7b4e11bce48fb6c0267",
    );
  });
  it("link keeps non-default ports", async () => {
    expect(await favId("link", { url: "https://example.com:8443/x?z=1" })).toBe(
      "ea73621b36e539f39eee4f960c10e7a7305429f3087c2005e5f5f69f340659f0",
    );
  });

  it("excerpt normalizes whitespace before hashing", async () => {
    expect(
      await favId("excerpt", {
        date: "2026-06-09",
        report: "ai-agents",
        excerpt: "Agent   memory\narchitecture",
      }),
    ).toBe("ee7cf0e9715741daf27c0c223665e81871bd6354f3daa6c3e385e5d4f6334a6c");
  });
  it("excerpt lowercases before hashing", async () => {
    expect(
      await favId("excerpt", {
        date: "2026-06-10",
        report: "ai-agents",
        excerpt: "Provider COMPATIBILITY Notes",
      }),
    ).toBe("87743bf98cbc5fc085e39d6088893760c46c584ae5d5a76f4b7e3f52deef6027");
  });
  it("excerpt handles unicode", async () => {
    expect(
      await favId("excerpt", {
        date: "2026-06-09",
        report: "ai-trending",
        excerpt: "关于 GitHub Trending 的趋势判断",
      }),
    ).toBe("edd2b2c7cc7f2c071424aaf1fb5c256632295bb0e228a69d31c98cb28df77b01");
  });
});
