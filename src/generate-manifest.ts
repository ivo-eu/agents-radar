import fs from "fs";
import crypto from "node:crypto";
import path from "path";
import { marked } from "marked";
import { REPORT_LABELS } from "./i18n.ts";

const DIGESTS_DIR = "digests";
const MANIFEST_PATH = "manifest.json";
const SYNC_MANIFEST_PATH = "sync-manifest.json";
const SEARCH_INDEX_DIR = "search-index";
const FEED_PATH = "feed.xml";
const SITE_URL = process.env["PAGES_URL"] ?? "https://ivo-eu.github.io/agents-radar";
const DATE_RE = /^\d{4}-\d{2}-\d{2}$/;
const REPORT_FILES = [
  "ai-cli",
  "ai-cli-en",
  "ai-agents",
  "ai-agents-en",
  "ai-web",
  "ai-web-en",
  "ai-trending",
  "ai-trending-en",
  "ai-hn",
  "ai-hn-en",
  "ai-ph",
  "ai-ph-en",
  "ai-arxiv",
  "ai-arxiv-en",
  "ai-hf",
  "ai-hf-en",
  "ai-community",
  "ai-community-en",
  "ai-weekly",
  "ai-weekly-en",
  "ai-monthly",
  "ai-monthly-en",
] as const;
const MAX_FEED_ITEMS = 30;

interface DateEntry {
  date: string;
  reports: string[];
}

interface Manifest {
  generated: string;
  dates: DateEntry[];
}

interface SyncFile {
  path: string;
  size: number;
  sha256: string;
}

interface SyncManifest {
  version: number;
  generated: string;
  files: SyncFile[];
}

interface ReportContent {
  summary: string;
  fullHtml: string;
}

const DAYS = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
const MONTHS = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"];

export function toRfc822(date: Date): string {
  return (
    `${DAYS[date.getUTCDay()]}, ${String(date.getUTCDate()).padStart(2, "0")} ` +
    `${MONTHS[date.getUTCMonth()]} ${date.getUTCFullYear()} ` +
    `${String(date.getUTCHours()).padStart(2, "0")}:${String(date.getUTCMinutes()).padStart(2, "0")}:${String(date.getUTCSeconds()).padStart(2, "0")} +0000`
  );
}

export function escapeXml(s: string): string {
  return s.replace(/&/g, "&amp;").replace(/</g, "&lt;").replace(/>/g, "&gt;").replace(/"/g, "&quot;");
}

export function markdownToSearchText(markdown: string): string {
  return markdown
    .replace(/```[\s\S]*?```/g, " ")
    .replace(/https?:\/\/\S+/g, " ")
    .replace(/<[^>]+>/g, " ")
    .replace(/[#>*_`~|[\]()-]/g, " ")
    .replace(/\s+/g, " ")
    .trim()
    .toLowerCase();
}

function walkFiles(dir: string): string[] {
  const files: string[] = [];
  for (const entry of fs.readdirSync(dir, { withFileTypes: true })) {
    if (entry.name.startsWith(".")) continue;
    const fullPath = path.join(dir, entry.name);
    if (entry.isDirectory()) {
      files.push(...walkFiles(fullPath));
    } else if (entry.isFile()) {
      files.push(fullPath);
    }
  }
  return files;
}

function generateSyncManifest(generated: string): void {
  const files = walkFiles(DIGESTS_DIR)
    .map((filePath): SyncFile => {
      const content = fs.readFileSync(filePath);
      return {
        path: path.relative(DIGESTS_DIR, filePath).split(path.sep).join("/"),
        size: content.length,
        sha256: crypto.createHash("sha256").update(content).digest("hex"),
      };
    })
    .sort((a, b) => a.path.localeCompare(b.path));

  const manifest: SyncManifest = { version: 1, generated, files };
  fs.writeFileSync(SYNC_MANIFEST_PATH, JSON.stringify(manifest, null, 2) + "\n");
  console.log(`${SYNC_MANIFEST_PATH} updated: ${files.length} files`);
}

function generateSearchIndexes(entries: DateEntry[], generated: string): void {
  fs.mkdirSync(SEARCH_INDEX_DIR, { recursive: true });
  const months = new Map<string, Record<string, string>>();

  for (const { date, reports } of entries) {
    const month = date.slice(0, 7);
    const monthEntries = months.get(month) ?? {};
    const chunks = reports
      .filter((report) => !report.endsWith("-en"))
      .map((report) => {
        const filePath = path.join(DIGESTS_DIR, date, `${report}.md`);
        return fs.existsSync(filePath) ? fs.readFileSync(filePath, "utf-8") : "";
      });
    monthEntries[date] = markdownToSearchText(chunks.join("\n"));
    months.set(month, monthEntries);
  }

  const monthList = [...months.keys()].sort().reverse();
  for (const month of monthList) {
    fs.writeFileSync(
      path.join(SEARCH_INDEX_DIR, `${month}.json`),
      JSON.stringify({ generated, month, dates: months.get(month) }) + "\n",
    );
  }
  fs.writeFileSync(
    path.join(SEARCH_INDEX_DIR, "manifest.json"),
    JSON.stringify({ generated, months: monthList }, null, 2) + "\n",
  );
  console.log(`${SEARCH_INDEX_DIR} updated: ${monthList.length} months`);
}

async function getReportContent(date: string, report: string): Promise<ReportContent> {
  const filePath = path.join(DIGESTS_DIR, date, `${report}.md`);

  try {
    const markdown = fs.readFileSync(filePath, "utf-8");
    const html = await marked.parse(markdown, { async: false });

    // Extract summary text from original HTML (before CDATA escape)
    const textOnly = html
      .replace(/<[^>]+>/g, "")
      .replace(/\s+/g, " ")
      .trim();
    const summary = textOnly.length > 500 ? textOnly.slice(0, 500) + "..." : textOnly;

    // Escape CDATA end marker to prevent injection
    const safeHtml = html.replace(/]]>/g, "]]]]><![CDATA[");

    return {
      summary: escapeXml(summary), // Plain text, XML-escaped, no CDATA
      fullHtml: `<![CDATA[${safeHtml}]]>`, // HTML in CDATA, no escaping needed
    };
  } catch {
    // Fallback to title-only content on any error
    const label = REPORT_LABELS[report] ?? report;
    const title = `${label} ${date}`;
    return {
      summary: escapeXml(title),
      fullHtml: `<![CDATA[${escapeXml(title)}]]>`,
    };
  }
}

async function main(): Promise<void> {
  const generated = new Date().toISOString();
  const entries = fs
    .readdirSync(DIGESTS_DIR)
    .filter((name) => DATE_RE.test(name) && fs.statSync(path.join(DIGESTS_DIR, name)).isDirectory())
    .sort()
    .reverse()
    .map((date) => {
      const reports = REPORT_FILES.filter((r) => fs.existsSync(path.join(DIGESTS_DIR, date, `${r}.md`)));
      return { date, reports };
    })
    .filter((e) => e.reports.length > 0);

  const manifest: Manifest = {
    generated,
    dates: entries,
  };

  fs.writeFileSync(MANIFEST_PATH, JSON.stringify(manifest, null, 2) + "\n");
  console.log(`manifest.json updated: ${entries.length} dates`);
  generateSyncManifest(generated);
  generateSearchIndexes(entries, generated);

  // ── RSS Feed ──────────────────────────────────────────────────────────────────

  const feedItems: Array<{ date: string; report: string }> = [];
  outer: for (const entry of entries) {
    for (const report of entry.reports) {
      feedItems.push({ date: entry.date, report });
      if (feedItems.length >= MAX_FEED_ITEMS) break outer;
    }
  }

  const buildDate = toRfc822(new Date());

  const itemXmlChunks: string[] = [];
  for (const { date, report } of feedItems) {
    const label = REPORT_LABELS[report] ?? report;
    const title = `${label} ${date}`;
    const link = `${SITE_URL}/#${date}/${report}`;
    const parts = date.split("-").map(Number);
    const pubDate = toRfc822(new Date(Date.UTC(parts[0]!, parts[1]! - 1, parts[2]!)));
    const content = await getReportContent(date, report);
    itemXmlChunks.push(
      [
        "    <item>",
        `      <title>${escapeXml(title)}</title>`,
        `      <link>${escapeXml(link)}</link>`,
        `      <guid isPermaLink="true">${escapeXml(link)}</guid>`,
        `      <pubDate>${pubDate}</pubDate>`,
        `      <description>${content.summary}</description>`,
        `      <content:encoded>${content.fullHtml}</content:encoded>`,
        "    </item>",
      ].join("\n"),
    );
  }
  const itemsXml = itemXmlChunks.join("\n");

  const feedXml =
    `<?xml version="1.0" encoding="UTF-8"?>\n` +
    `<rss version="2.0" xmlns:atom="http://www.w3.org/2005/Atom" xmlns:content="http://purl.org/rss/1.0/modules/content/">\n` +
    `  <channel>\n` +
    `    <title>agents-radar</title>\n` +
    `    <link>${SITE_URL}</link>\n` +
    `    <description>AI 开源生态每日简报 · Daily AI ecosystem digest</description>\n` +
    `    <language>zh-CN</language>\n` +
    `    <atom:link href="${SITE_URL}/feed.xml" rel="self" type="application/rss+xml"/>\n` +
    `    <lastBuildDate>${buildDate}</lastBuildDate>\n` +
    itemsXml +
    `\n  </channel>\n` +
    `</rss>\n`;

  fs.writeFileSync(FEED_PATH, feedXml);
  console.log(`feed.xml updated: ${feedItems.length} items`);
}

// Run only when executed directly (not imported for testing)
const isDirectRun =
  process.argv[1] &&
  (process.argv[1].endsWith("generate-manifest.ts") || process.argv[1].endsWith("generate-manifest.js"));
if (isDirectRun) {
  main().catch((err) => {
    console.error(err);
    process.exit(1);
  });
}
