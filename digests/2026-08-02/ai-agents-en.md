# OpenClaw Ecosystem Digest 2026-08-02

> Issues: 222 | PRs: 500 | Projects covered: 13 | Generated: 2026-08-02 00:13 UTC

- [OpenClaw](https://github.com/openclaw/openclaw)
- [NanoBot](https://github.com/HKUDS/nanobot)
- [Hermes Agent](https://github.com/nousresearch/hermes-agent)
- [PicoClaw](https://github.com/sipeed/picoclaw)
- [NanoClaw](https://github.com/qwibitai/nanoclaw)
- [NullClaw](https://github.com/nullclaw/nullclaw)
- [IronClaw](https://github.com/nearai/ironclaw)
- [LobsterAI](https://github.com/netease-youdao/LobsterAI)
- [TinyClaw](https://github.com/TinyAGI/tinyagi)
- [Moltis](https://github.com/moltis-org/moltis)
- [CoPaw](https://github.com/agentscope-ai/CoPaw)
- [ZeptoClaw](https://github.com/qhkm/zeptoclaw)
- [ZeroClaw](https://github.com/zeroclaw-labs/zeroclaw)

---

## OpenClaw Deep Dive

# OpenClaw Project Digest — 2026-08-02

## 1. Today's Overview

OpenClaw is in a high-activity maintenance and hardening cycle: **222 issues** were updated in the last 24 hours (200 open, 22 closed) and **500 PRs** were touched (397 open, **103 merged/closed**). A new beta, **v2026.7.2-beta.6**, shipped with a significant batch of state-safety and crash-recovery features. The tracker remains dominated by P1 reliability issues — silent message failures, crash-loop recovery gaps, and session-state leaks — with the hottest thread (DeepSeek v4 Flash silent reply failure) at 73 comments. Overall health signals are mixed but trending positive: high merge throughput, quick issue triage, and an active ClawSweeper automation loop, offset by several long-lived P0/P1 items around docs drift, DB migration safety, and provider compatibility.

## 2. Releases

### v2026.7.2-beta.6 (`openclaw 2026.7.2-beta.6`)

This release is squarely focused on **state safety and recovery**:

- **Quarantine store** that survives primary-database damage
- **Crash-recoverable SQLite snapshots**
- **Crash-durable filesystem publication**
- **Schema-upgrade data-loss rejection**
- **Rollback-writer snapshot recovery**

These changes directly target several high-severity issues in the tracker, including schema downgrade/upgrade hazards ([#115421](https://github.com/openclaw/openclaw/issues/115421)) and post-migration state corruption ([#114084](https://github.com/openclaw/openclaw/issues/114084)). No explicit breaking changes or migration notes beyond the schema-safety machinery were included in the visible release notes.

## 3. Project Progress

**103 PRs merged/closed today.** Notable areas of forward motion (from the active pipeline):

- **CLI/Gateway performance**: startup-phase attribution in diagnostics ([PR #117702](https://github.com/openclaw/openclaw/pull/117702)); keeping gateway-backed agent turns cold to cut a ~13s startup tax ([PR #117705](https://github.com/openclaw/openclaw/pull/117705)); `openclaw sessions` now reuses plugin metadata, cutting a 344-session listing from ~31s ([PR #117707](https://github.com/openclaw/openclaw/pull/117707))
- **Security**: skill scans now detect aliased shell execution via TypeScript AST ([PR #117394](https://github.com/openclaw/openclaw/pull/117394)); GitHub Actions dependency group bump ([PR #117180](https://github.com/openclaw/openclaw/pull/117180))
- **Channel reliability**: Tlon channel-history refresh after monitor restart ([PR #114580](https://github.com/openclaw/openclaw/pull/114580)); Slack ingress isolation per thread ([PR #114552](https://github.com/openclaw/openclaw/pull/114552)); QQBot heartbeat-ACK liveness watchdog ([PR #114902](https://github.com/openclaw/openclaw/pull/114902)); Feishu legacy document read support ([PR #81766](https://github.com/openclaw/openclaw/pull/81766))
- **Codex integration**: reduced active-run transcript overhead ([PR #115049](https://github.com/openclaw/openclaw/pull/115049)); preserved upstream token-usage breakdown ([PR #117013](https://github.com/openclaw/openclaw/pull/117013))
- **Memory**: stale memory search results now qualified ([PR #117706](https://github.com/openclaw/openclaw/pull/117706)); recovery of primary embedding provider after transient outages ([PR #116562](https://github.com/openclaw/openclaw/pull/116562))
- **Merged via automerge**: Grok subscription HTTP 426 fix ([PR #117704](https://github.com/openclaw/openclaw/pull/117704), closed today)

## 4. Community Hot Topics

1. **[#116277 — DeepSeek v4 Flash silent reply failure](https://github.com/openclaw/openclaw/issues/116277)** (P1, 73 comments) — The model silently fails to generate replies and OpenClaw posts a generic "No reply was generated" fallback. Community concern centers on absence of model-level failure surfacing. Underlying need: **observable, actionable model failure handling** instead of silent degradation.
2. **[#116201 — Realtime voice retains unbounded provider/consult state](https://github.com/openclaw/openclaw/issues/116201)** (P1, 36 comments) — Resource limits expressed as item counts/cancellation signals rather than hard ownership bounds; slow/bursty providers retain superseded work. Underlying need: **hard resource ownership bounds for realtime voice sessions**.
3. **[#99241 — Tool outputs render as image attachments, unreadable to agent](https://github.com/openclaw/openclaw/issues/99241)** (26 comments, 2👍) — **Closed** this cycle; long-running ANSI-heavy workflows collapsed tool results into `(see attached image)` placeholders. Resolution is a win for evidence-heavy agent workflows.
4. **[#115326 — Crash-loop breaker permanently suppresses Discord/WhatsApp](https://github.com/openclaw/openclaw/issues/115326)** (P1, 24 comments) — Documented recovery (`channels.start`) fails with WebSocket 1006. Underlying need: **reliable, documented channel recovery** after breaker trips.
5. **[#48920 — Live Docs ahead of release](https://github.com/openclaw/openclaw/issues/48920)** (P0, 10 comments, 4👍) — Docs reference `IsolatedSessions` not present in the latest release. Community frustration with **doc/release drift**.
6. **[#95279 — Trusted inbound-decoration contract](https://github.com/openclaw/openclaw/issues/95279)** (5 comments, 4👍) — Users want a non-forgeable way to strip/dedup metadata decorations rather than text heuristics.

## 5. Bugs & Stability

### P0 (release blockers)
- **[#48920 — Live Docs ahead of release](https://github.com/openclaw/openclaw/issues/48920)** — `IsolatedSessions` documented but unavailable in v2026.3.13.
- **[#115421 — Schema downgrade recovery wipes state DB, cron jobs lost](https://github.com/openclaw/openclaw/issues/115421)** — Recovery quarantines the v6 DB and creates a fresh empty one; linked PR open.

### P1 (high-impact reliability)
- **[#116277 — DeepSeek v4 Flash silent reply failure](https://github.com/openclaw/openclaw/issues/116277)** (73 comments) — Generic fallback masks model failure; needs live repro.
- **[#116201 — Realtime voice state retention](https://github.com/openclaw/openclaw/issues/116201)** (36 comments) — Unbounded provider/consult state under bursty conditions.
- **[#115326 — Crash-loop breaker suppresses Discord/WhatsApp permanently](https://github.com/openclaw/openclaw/issues/115326)** (24 comments) — Recovery path broken (WebSocket 1006).
- **[#115424 — Gateway V8 heap OOM → 7-core-dump loop](https://github.com/openclaw/openclaw/issues/115424)** — Restart-recovery hot-resumes a poisoned session, converting one crash into repeated OOMs.
- **[#115546 — CLI-budget compaction fires early (4.9s–50s), 100% failure on large sessions](https://github.com/openclaw/openclaw/issues/115546)** — No retry; wake death-spiral on orchestrator sessions.
- **[#117262 — SQLite contention: 3 write handles cause ~33s event-loop stalls](https://github.com/openclaw/openclaw/issues/117262)** — WAL at max, event-loop blocked; DEF-61.
- **[#115642 — Billing cooldown outlives the outage](https://github.com/openclaw/openclaw/issues/115642)** — Fixed ~5h `disabledUntil` window for subscription auth; no probe-based recovery.
- **[#114234 — Usage-cost refresh lock never released after PID reuse in containers](https://github.com/openclaw/openclaw/issues/114234)** — Permanently frozen cache; linked PR open ([#103961](https://github.com/openclaw/openclaw/pull/103961) addresses the same lock-leak family).
- **[#106231 — Loop detection blocks exec but never terminates the stuck run](https://github.com/openclaw/openclaw/issues/106231)** — Resources burned for hours; linked PR open.
- **[#87763 — SSRF guard pinned DNS vs. autoSelectFamily → 120s timeouts](https://github.com/openclaw/openclaw/issues/87763)** — Gateway unresponsive on affected providers.
- **[#116488 — Superseded reply never released from registry](https://github.com/openclaw/openclaw/issues/116488)** — Session reports active work after `run:completed`; watchdog waits out abort timer.
- **[#117491 — Heartbeat delivery routes to wrong Telegram account](https://github.com/openclaw/openclaw/issues/117491)** — `accountId` cleared by `commitmentDeliveryContext`; linked PR open.
- **[#114084 — "no such table: session_entries" after migration](https://github.com/openclaw/openclaw/issues/114084)** — Diagnostic subsystem still references the old table name on beta.4.
- **[#116010 — All persistent sessions capped at 128k context regardless of model](https://github.com/openclaw/openclaw/issues/116010)** — Likely linked to the same PR family as the schema migration.
- **[#116691 — Volcano Engine long conversations fail (`missing input.status`)](https://github.com/openclaw/openclaw/issues/116691)** — openai-responses regression on the second contextual turn.

**Resolved this cycle:** [#99241](https://github.com/openclaw/openclaw/issues/99241) (tool-output images, closed), [#34528](https://github.com/openclaw/openclaw/issues/34528) (Feishu reaction message_id suffix, closed), [#106730](https://github.com/openclaw/openclaw/issues/106730) (exec pipe misinterpretation, closed), [#115413](https://github.com/openclaw/openclaw/issues/115413) (compaction false-success report, closed), [#90203](https://github.com/openclaw/openclaw/issues/90203) (MEMORY.md false-missing check, closed).

## 6. Feature Requests & Roadmap Signals

Strong signals for upcoming releases:

- **[#115924 — "Idea Shower": parallel thought collector while agent works](https://github.com/openclaw/openclaw/issues/115924)** (new, P3) — A UX shift toward parallel human-agent interaction rather than serial interrupt/queue.
- **[#114264 — Automatic model routing by message type](https://github.com/openclaw/openclaw/issues/114264)** (P2) — Route text/image/audio/TTS to specialized models (e.g., Xiaomi MiMo V2.5 family).
- **[#114146 — `talk.realtime.providers.<id>.baseUrl` for OpenAI Realtime-compatible providers](https://github.com/openclaw/openclaw/issues/114146)** (P2, 1👍) — Needed for Alibaba Bailian Qwen3-ASR-Flash and proxy/middleware services.
- **[#113251 — Image viewing in the webchat file viewer](https://github.com/openclaw/openclaw/issues/113251)** (P2) — Currently impossible to preview images in webchat history.
- **[#115400 — `sessions_send`: synchronous wait + duplicate-delivery fix](https://github.com/openclaw/openclaw/issues/115400)** (P2) — Gatekeeper-agent pattern users need a blocking send option.
- **[#95724 — Memory index by source directory, not agent](https://github.com/openclaw/openclaw/issues/95724)** (P2, 1👍) — Eliminates duplicate vector stores for same-workspace agents.
- **[#86983 — Outbound DM allowlist (`dmAllowTo`)](https://github.com/openclaw/openclaw/issues/86983)** (P2, 1👍) — Technical enforcement for outbound DMs; needs security review.
- **[#95516 — Skill lifecycle management](https://github.com/openclaw/openclaw/issues/95516)** (P2, 2👍) — Auto-optimization on failure plus usage-based retirement.
- **[#95601 — VoiceOver-friendly chat history](https://github.com/openclaw/openclaw/issues/95601)** (P2, 2👍) — Accessibility follow-up; user explicitly thanked the team for usage-display improvements.

**Prediction:** `baseUrl` for realtime providers ([#114146](https://github.com/openclaw/openclaw/issues/114146)) and webchat image viewing ([#113251](https://github.com/openclaw/openclaw/issues/113251)) are contained, high-value changes likely to land soon. The "Idea Shower" concept ([#115924](https://github.com/openclaw/openclaw/issues/115924)) is more speculative but signals growing interest in parallel interaction models.

## 7. User Feedback Summary

**Satisfaction signals:** Users publicly thanked maintainers for accessibility improvements in v2026.6.9 ([#95601](https://github.com/openclaw/openclaw/issues/95601)); 103 PRs merged today demonstrates a responsive maintainer/automation pipeline; the ClawSweeper autogenerated PRs ([#117443](https://github.com/openclaw/openclaw/pull/117443), [#117704](https://github.com/openclaw/openclaw/pull/117704)) show an effective automated fix loop.

**Recurring pain points:**

- **Silent message loss / generic fallbacks** — the dominant theme across [#116277](https://github.com/openclaw/openclaw/issues/116277), [#115326](https://github.com/openclaw/openclaw/issues/115326), [#115476](https://github.com/openclaw/openclaw/issues/115476), [#95566](https://github.com/openclaw/openclaw/issues/95566)
- **Scale performance** — `openclaw sessions` at 31s for 344 sessions and ~13s startup tax per agent turn on plugin-heavy hosts ([PR #117707](https://github.com/openclaw/openclaw/pull/117707), [PR #117705](https://github.com/openclaw/openclaw/pull/117705))
- **Provider compatibility friction** — DeepSeek v4 Flash, Volcano Engine, Anthropic refusals ([#98976](https://github.com/openclaw/openclaw/issues/98976)), Grok HTTP 426 ([PR #117704](https://github.com/openclaw/openclaw/pull/117704))
- **Broken recovery paths** — crash-loop breaker ([#115326](https://github.com/openclaw/openclaw/issues/115326)), compaction false-success ([#115413](https://github.com/openclaw/openclaw/issues/115413)), retired Codex binding tombstones ([#116022](https://github.com/openclaw/openclaw/issues/116022))
- **Windows/PowerShell command incompatibility** — agent emits Unix commands (`head`, tilde expansion) on Windows ([#117644](https://github.com/openclaw/openclaw/issues/117644))
- **Silent plugin-load failures** — enabled plugin invisible on every diagnostic surface ([#117243](https://github.com/openclaw/openclaw/issues/117243)); WeChat plugin broken export ([#115478](https://github.com/openclaw/openclaw/issues/115478))

## 8. Backlog Watch

Long-unanswered items requiring maintainer attention:

- **[PR #80228 — Codex continuity bridge](https://github.com/openclaw/openclaw/pull/80228)** — Open since May 10; tagged needs-real-behavior-proof with multiple merge-risk flags (compatibility, message-delivery, security-boundary).
- **[PR #81766 — Feishu legacy document read support](https://github.com/openclaw/openclaw/pull/81766)** — Open since May 14; still awaiting real-behavior proof.
- **[PR #81470 — Webchat TTS audio in `broadcastChatFinal`](https://github.com/openclaw/openclaw/pull/81470)** — Open since May 13; session-state and message-delivery merge risks.
- **[#48920 — Live Docs ahead of release](https://github.com/openclaw/openclaw/issues/48920)** — **P0**, open since March 17; docs/release drift is a recurring trust issue.
- **[#87763 — SSRF guard DNS timeouts](https://github.com/openclaw/openclaw/issues/87763)** — P1, open since May 28; needs security review plus linked PR.
- **[#86983 — Outbound DM allowlist](https://github.com/openclaw/openclaw/issues/86983)** — Open since May 26; needs product decision and security review.
- **[#88909 — NSUserDefaults bundle-identifier warning on macOS](https://github.com/openclaw/openclaw/issues/88909)** — Open since June 1; needs live repro (a fix PR now exists: [#117690](https://github.com/openclaw/openclaw/pull/117690)).
- **[#95724 — Memory indexing by source directory](https://github.com/openclaw/openclaw/issues/95724)** — Open since June 22; maintainer review pending.
- **[#103961 — Usage-cost refresh lock leak fix](https://github.com/openclaw/openclaw/pull/103961)** — Open since July 10; needs proof, related to the P1 [#114234](https://github.com/openclaw/openclaw/issues/114234).
- **[#110024 — `restartPending` reconciliation](https://github.com/openclaw/openclaw/pull/110024)** — Open since July 17; documented as the cause of two multi-hour production outages (~34h total); still needs proof.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report — Personal AI Assistant / Agent Open-Source Ecosystem
**Date:** 2026-08-02 | **Coverage:** 13 projects in the OpenClaw ecosystem family

---

## 1. Ecosystem Overview

The personal AI assistant open-source landscape is in a **reliability-hardening phase** rather than a feature-expansion phase. Across the 10 active projects, the dominant themes are silent message failures, state-durability gaps, and provider-compatibility friction. The ecosystem is consolidating around a common architectural pattern — a gateway/agent core with pluggable channels (Slack, WhatsApp, Telegram, Matrix, iMessage) and model-provider adapters. Merge velocity remains healthy at the top (103 merges/day in OpenClaw alone), but several mid-tier projects show **maintainer-bottleneck symptoms**: zero merge throughput despite high PR volume (ZeroClaw), stale-bot sweeps without fixes (LobsterAI), and long-unreviewed PRs (CoPaw, IronClaw). Security isolation is rapidly moving from a nice-to-have to an S0 requirement as multi-agent deployments proliferate.

---

## 2. Activity Comparison

*Health score = weighted composite of merge throughput, issue-closure velocity, release cadence, and severity of open issues (10 = best).*

| Project | Issues Updated (Closed) | PRs Updated (Merged/Closed) | Release | Health |
|---|---|---|---|---|
| **OpenClaw** | 222 (22) | 500 (103) | v2026.7.2-beta.6 | **8/10** |
| **Hermes Agent** | 11 (3) | 50 (20) | None | **8/10** |
| **NanoBot** | 5 (4) | 25 (13) | None | **8/10** |
| **IronClaw** | 20 (3) | 25 (9) | None (release PR pending ~30d) | **7/10** |
| **NanoClaw** | 2 (0) | 16 (6) | v2.1.54 (breaking) | **7/10** |
| **Moltis** | 0 (0) | 4 (3) | None | **6/10** |
| **CoPaw** | 9 (0) | 12 (1) | None | **6/10** |
| **PicoClaw** | 1 (0) | 3 (1) | None | **5/10** |
| **ZeroClaw** | 14 (0) | 50 (0) | None (v0.8.5/v0.9.0 tracked) | **4/10** |
| **LobsterAI** | 7 (6 stale) | 2 (0) | None | **3/10** |
| NullClaw | 0 | 0 | None | Dormant |
| TinyClaw | 0 | 0 | None | Dormant |
| ZeptoClaw | 0 | 0 | None | Dormant |

**Key observations:**
- OpenClaw's daily PR volume is **10× the next-busiest project** (500 vs 50).
- ZeroClaw has OpenClaw-scale PR traffic (50) but **zero merges** — a clear review bottleneck.
- Three projects (NullClaw, TinyClaw, ZeptoClaw) show no activity at all, indicating ecosystem consolidation.

---

## 3. OpenClaw's Position

**Advantages:**
- **Reference-project gravity:** Largest community by an order of magnitude (222 issues/day, 500 PRs/day); the de facto standard for channel breadth and gateway architecture.
- **Automated fix loop:** ClawSweeper automation merges fixes (e.g., Grok HTTP 426) without human intervention, enabling scale no peer matches.
- **State-safety engineering:** The v2026.7.2-beta.6 quarantine store and crash-recoverable snapshots are the most advanced durability work in the ecosystem.
- **Channel coverage:** Slack, Discord, WhatsApp, Feishu, QQBot, Tlon, Telegram, and more — unmatched breadth.

**Technical approach differences:**
- Beta-gated, high-cadence releases (v2026.7.2-beta.6) vs. IronClaw's ~30-day release-PR stalls and NanoClaw's large rollup releases (v2.1.18→v2.1.54).
- Maintainer-driven core with heavy automation, contrasted with Hermes's multi-profile/multi-gateway infrastructure focus and ZeroClaw's agent-centric isolation model.

**Risks relative to peers:**
- P0 doc/release drift (#48920) is a trust issue peers handle better (smaller docs surface).
- Reliability debt at scale: 13 P1 items open, including silent-failure and crash-loop clusters. Hermes and NanoBot are closing security/reliability bugs faster *relative to their size*.

---

## 4. Shared Technical Focus Areas

Requirements emerging independently across multiple projects:

| Theme | Projects | Specific Needs |
|---|---|---|
| **State durability & crash recovery** | OpenClaw, NanoBot, NanoClaw, IronClaw | Crash-safe DB snapshots, migration safety, cron state persistence, cache stability — OpenClaw (#115421), NanoBot (#5163, #4801), NanoClaw (#3166), IronClaw (#6984–#6986) |
| **Provider compatibility & routing** | OpenClaw, NanoBot, Hermes, IronClaw, CoPaw, PicoClaw, ZeroClaw | DeepSeek v4 Flash silent failures (OpenClaw), Gemini `ToolCallBlock` crash (CoPaw), `api_base` hijacking (NanoBot), OmniRoute 503 handling (Hermes), OpenRouter session_id caching (ZeroClaw). **OrcaRouter integration appeared in 3 projects on the same day** (IronClaw #7009, CoPaw #6622, PicoClaw #3309) |
| **Security & privilege isolation** | ZeroClaw, Hermes, Moltis, NanoBot, OpenClaw | Per-agent session/channel ownership (ZeroClaw S0 #9646/#9647), profile secret scoping (Hermes), operator allowlists (Moltis), per-sender rate limits (NanoBot), AST-based shell-execution detection (OpenClaw) |
| **Observability & failure surfacing** | OpenClaw, ZeroClaw, NanoBot, NanoClaw | Replace silent degradation with actionable errors: generic "No reply" fallbacks (OpenClaw #116277), approval-timeout logged as denial (ZeroClaw #9642), tool-output leakage (NanoBot #5185), credential-expiry alerts (NanoClaw #3167) |
| **Performance at scale** | OpenClaw, IronClaw, ZeroClaw | 31s session listings and 13s startup tax (OpenClaw), libSQL p95 37–135s regressions (IronClaw #6974), SQLite event-loop stalls (OpenClaw #117262) |
| **Docs/release drift** | OpenClaw, ZeroClaw, NanoClaw | Live docs ahead of release (OpenClaw P0 #48920), nonexistent config keys (ZeroClaw #9640), init docs misalignment (NanoClaw #3046) |
| **Multi-agent discoverability** | CoPaw, OpenClaw, NanoBot | Agents not activating without explicit PROFILE.md references (CoPaw #6621), parallel interaction models (OpenClaw #115924), subagent presets (NanoBot #5207) |

---

## 5. Differentiation Analysis

| Project | Primary Focus | Target User | Architecture Signature |
|---|---|---|---|
| **OpenClaw** | Breadth: channels + providers + autonomy | Power users, self-hosters | Node/TS gateway monolith, beta-gated releases, automation-assisted |
| **Hermes Agent** | Multi-profile/multi-gateway infrastructure | Operators running agent fleets | Profile-scoped secrets, desktop app, systemd integration |
| **IronClaw** | Enterprise architecture quality | Engineering orgs, Rust teams | Rust workspace, `product_contracts` dependency inversion, CI-gate discipline |
| **ZeroClaw** | Security & multi-agent isolation | Security-conscious multi-agent users | Agent-centric ownership model, breaking-change roadmaps (v0.9.0) |
| **NanoClaw** | Consumer channels, hosted backends | Chat-first users (iMessage/WhatsApp) | Channel unification (local + hosted backends), setup-flow UX |
| **CoPaw** | Desktop UX + Chinese ecosystem | Aliyun/Qwen ecosystem users | Desktop app, global hotkey ambitions, ACP transport |
| **NanoBot** | Lightweight WebUI productivity | Casual self-hosters | Fast triage-to-fix, WebUI chat modes, cron/agent workflows |
| **Moltis** | Team collaboration, Nostr | Nostr/Buzz workspace users | NIP-29 group chat, per-account operator governance |
| **PicoClaw** | Embedded/low-power, i18n | Matrix users, Sipeed hardware owners | Minimal footprint, localization breadth (zh-TW just merged) |
| **LobsterAI** | Studio/console companion | OpenClaw console users | Currently stalled: stale-cleanup only, 2 PRs stuck ~4 months |

**Strategic divergence:** The ecosystem is splitting into **generalists** (OpenClaw, Hermes) vs. **specialists** (Moltis→Nostr, PicoClaw→embedded, CoPaw→Chinese desktop). IronClaw and ZeroClaw are competing on the *quality bar* — IronClaw via architecture governance, ZeroClaw via security defaults.

---

## 6. Community Momentum & Maturity

**Tier 1 — Rapid iteration (shipping daily/weekly):**
- **OpenClaw** — 103 merges/day, beta releases, automated fix loops. Reliability debt is the constraint.
- **Hermes Agent** — 20 merges/day equivalent; security-critical fixes (profile secret isolation) closed same-window as reports. Most responsive maintainer team relative to issue load.
- **NanoBot** — 13 merges, 4/5 issues closed; fastest triage-to-fix ratio in the ecosystem.

**Tier 2 — Steady / feature-driven:**
- **IronClaw** — High throughput but mid-migration (Reborn Wave 2); release train stalled ~30 days.
- **NanoClaw** — Shipped a breaking release (v2.1.54) with fast issue→fix loops; backlog of 7+ week-old PRs.
- **Moltis** — Small but completing: 3/4 PRs merged; clean, focused scope.
- **CoPaw** — Active fix wave, but 0 releases means users wait; long-unmerged UX PRs (>1 month).

**Tier 3 — Stabilizing / at-risk:**
- **ZeroClaw** — High contribution interest, **zero merge throughput**. Queue bloat with `stale-candidate`/`needs-author-action` labels; S0 security issues open with no fix PRs.
- **PicoClaw** — Low but alive; one stale Matrix-sync reliability bug dominating.
- **LobsterAI** — Stale-bot is doing more work than maintainers; 2 fix PRs stuck since April.

**Tier 4 — Dormant:** NullClaw, TinyClaw, ZeptoClaw — no activity in 24h; likely candidates for archival or community takeover.

---

## 7. Trend Signals

1. **Reliability is the competitive battleground.** Silent message loss and generic fallbacks are the #1 complaint across OpenClaw (`#116277`, 73 comments), NanoBot (`#5185`), and ZeroClaw (`#9642`). Expect observable failure-handling (model-level error surfacing, audit-true denial logs) to become a differentiator.

2. **State durability is the next architectural frontier.** OpenClaw's quarantine store/snapshots, IronClaw's cache-stability P0s, and NanoClaw's migration crash all point to the same gap: **session state is the asset, and it is not yet crash-safe**. Agent developers should design durable, recoverable state from day one.

3. **Provider multi-homing is table stakes.** Almost every project hit a provider bug this week (DeepSeek, Gemini, Grok, Volcano Engine, Anthropic, OpenRouter). The simultaneous appearance of OrcaRouter PRs in IronClaw, CoPaw, and PicoClaw signals demand for router-style provider abstraction as a built-in.

4. **Security isolation is catching up to multi-agent adoption.** ZeroClaw's S0 per-agent ownership gaps (`#9646`, `#9647`) are the most severe, but Hermes's secret-scope migration and Moltis's operator list show the ecosystem responding. **Per-agent data isolation must be assumed, not configured.**

5. **Cost optimization is moving into the agent layer.** OpenRouter prompt-cache session IDs (ZeroClaw `#9631`), token-accounting bugs (IronClaw `#6989`), and rate limiting (NanoBot `#5108`) indicate economic pressure on multi-turn agent deployments.

6. **Human-agent interaction is shifting from serial to parallel.** OpenClaw's "Idea Shower" (`#115924`) and CoPaw's global-hotkey quick input (`#6568`) both envision background thought collection and lightweight interrupt — a UX model borrowed from launcher apps (Raycast) rather than chat UIs.

7. **Docs drift is a trust eroder at every scale.** OpenClaw's P0 `#48920` (live docs ahead of release) and ZeroClaw's nonexistent config key (`#9640`) show that documentation quality tracks project velocity inversely. Automated doc/release sync will be a quiet differentiator.

**Bottom line for developers:** Build for observable failures, assumption-free state durability, multi-provider routing, and per-agent isolation *upfront*. The ecosystem is converging on these four pillars, and the projects that ship them first — OpenClaw, Hermes, and NanoBot are currently closest — will define the reference patterns for 2026-2027.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-08-02

## Today's Overview
Project activity is high: 25 PRs were updated in the last 24 hours (12 open, 13 closed/merged), and 5 issues were touched (4 closed, 1 open). The current work focuses on WebUI improvements, session/memory hardening, cron reliability, and provider fixes. Several reported bugs already have fix PRs closed, indicating a fast triage-to-fix loop. The main open threads are per-session model switching and a few long-running provider-related PRs. Overall, the project appears healthy and close to a solid patch/minor release.

## Releases
No new releases in the last 24 hours; no changelog or migration notes to report.

## Project Progress
The reported set includes 13 closed/merged PRs. Notable advances:

- [PR #5172](https://github.com/HKUDS/nanobot/pull/5172) — preserve and replay Responses API reasoning state and compact context.
- [PR #5108](https://github.com/HKUDS/nanobot/pull/5108) — add per-sender message rate limiting across channels, addressing abuse/token-cost risks.
- [PR #5183](https://github.com/HKUDS/nanobot/pull/5183) — preserve cron manual run completion state; fixes [Issue #5163](https://github.com/HKUDS/nanobot/issues/5163).
- [PR #5153](https://github.com/HKUDS/nanobot/pull/5153) — handle non-string timestamps and missing `role` fields in raw archive; fixes [Issue #4801](https://github.com/HKUDS/nanobot/issues/4801).
- [PR #5208](https://github.com/HKUDS/nanobot/pull/5208) — advance Dream cron cursor after durable changes, preventing repeated history reprocessing.
- [PR #5200](https://github.com/HKUDS/nanobot/pull/5200) — preserve exec wait targets under response truncation.
- [PR #5201](https://github.com/HKUDS/nanobot/pull/5201) — tolerate malformed persisted session summaries.
- [PR #3732](https://github.com/HKUDS/nanobot/pull/3732) — require `api_base` before local provider wins on keyword match, preventing provider hijacking.
- [PR #5209](https://github.com/HKUDS/nanobot/pull/5209) — WebUI sidebar selection highlight refactor.
- [PR #5199](https://github.com/HKUDS/nanobot/pull/5199) — narrow Pyright suppressions in CLI code.

## Community Hot Topics
- [Issue #5185](https://github.com/HKUDS/nanobot/issues/5185) — 4 comments. User saw tool-call code appearing in model responses. It was closed as invalid/provider-related, but it highlights a need for better provider diagnostics when tool output leaks into visible model text.
- [Issue #5205](https://github.com/HKUDS/nanobot/issues/5205) — 2 comments. `nanobot plugins enable feishu` fails because the uv-managed Python has no `ensurepip`. This is a real environment/packaging friction point for plugin onboarding.
- [Issue #5198](https://github.com/HKUDS/nanobot/issues/5198) — open, 1 comment. Users want per-session model switching; `/model` does not behave as expected and the WebUI model blip is not interactive. This matches a common expectation from commercial chat UIs.

## Bugs & Stability
Ranked by severity:

1. **Cron completion state lost on WebUI polling** — [Issue #5163](https://github.com/HKUDS/nanobot/issues/5163). Manual cron runs succeeded but stayed marked `Failed`. Fixed by [PR #5183](https://github.com/HKUDS/nanobot/pull/5183).
2. **KeyError on malformed session history** — [Issue #4801](https://github.com/HKUDS/nanobot/issues/4801). Unprotected `message['role']` access crashes `MemoryStore._format_messages`. Fixed by [PR #5153](https://github.com/HKUDS/nanobot/pull/5153).
3. **Plugin installation broken in uv environment** — [Issue #5205](https://github.com/HKUDS/nanobot/issues/5205). Missing `ensurepip` blocks Feishu channel enablement. Closed; no corresponding fix PR was listed in the visible snapshot.
4. **Per-session model switching blocked** — [Issue #5198](https://github.com/HKUDS/nanobot/issues/5198). Open bug; no fix PR yet. Related UX work [PR #5202](https://github.com/HKUDS/nanobot/pull/5202) only makes preset switching more discoverable.
5. **Tool-call code in responses** — [Issue #5185](https://github.com/HKUDS/nanobot/issues/5185). Closed as invalid/provider, not a project defect, but worth monitoring if similar reports recur.

## Feature Requests & Roadmap Signals
Open PRs and [Issue #5198](https://github.com/HKUDS/nanobot/issues/5198) point to the next feature themes:

- **WebUI chat modes**: Quick Chat and Temporary Chat — [PR #5184](https://github.com/HKUDS/nanobot/pull/5184).
- **Cross-session search and @-mentions** — [PR #5211](https://github.com/HKUDS/nanobot/pull/5211).
- **Model preset switching UX**: click/tap menu replacing the hidden gesture — [PR #5202](https://github.com/HKUDS/nanobot/pull/5202).
- **Trusted proxy bootstrap auth** for Cloudflare Tunnel/Access deployments — [PR #5210](https://github.com/HKUDS/nanobot/pull/5210).
- **Subagent model presets** via the `spawn` tool — [PR #5207](https://github.com/HKUDS/nanobot/pull/5207).
- **Well-known skills.sh sources** support — [PR #5186](https://github.com/HKUDS/nanobot/pull/5186).
- **JSONL session list/thread loading performance** — [PR #5194](https://github.com/HKUDS/nanobot/pull/5194).

If merged, the next release will likely emphasize WebUI productivity, security hardening, and smoother model/subagent configuration.

## User Feedback Summary
- **Tool output confusion**: [Issue #5185](https://github.com/HKUDS/nanobot/issues/5185) reports user-facing tool-call code; even after triage, this kind of output can erode trust and needs better provider-side safeguards or debug logging.
- **Plugin enablement friction**: [Issue #5205](https://github.com/HKUDS/nanobot/issues/5205) shows a Debian server user cannot enable Feishu without dealing with Python environment issues manually.
- **Model switching expectation**: [Issue #5198](https://github.com/HKUDS/nanobot/issues/5198) indicates users expect session-level model control comparable to commercial chat UIs; the current model blip and `/model` command are not adequate.
- **Cron trust**: [Issue #5163](https://github.com/HKUDS/nanobot/issues/5163) shows users care about accurate execution state; seeing `Failed` after a successful run is a serious satisfaction issue, now fixed.

## Backlog Watch
- [PR #3732](https://github.com/HKUDS/nanobot/pull/3732) — open since 2026-05-11. Provider selection fix requiring `api_base` before local provider wins on keyword match. Prevents silent model hijacking; needs maintainer review.
- [PR #3869](https://github.com/HKUDS/nanobot/pull/3869) — open since 2026-05-16. DeepSeek message hardening: null content, `"(empty)"` leak, assistant text preservation. Still marked with a conflict label.
- [PR #5139](https://github.com/HKUDS/nanobot/pull/5139) — open since 2026-07-28. Priority P1 fix for dropped media paths during session consolidation. Has a conflict label and is important for data preservation.
- [Issue #5198](https://github.com/HKUDS/nanobot/issues/5198) — open since 2026-07-31. Per-session model switching may need a maintainer decision on API/UI design rather than just a small fix.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-08-02

## 1. Today's Overview

This is a **high-activity maintenance period** for Hermes Agent: 11 issues were updated in the last 24 hours (8 open, 3 closed), and 50 PRs were updated (30 still open, 20 closed/merged). No new release was cut, so the day's work is landing directly on `main`. The overall theme is **reliability hardening**: update/install recovery, profile secret isolation, Windows/desktop compatibility, and gateway state correctness. Maintainers appear responsive, with several high-severity fixes already closed or submitted in the same window as their related bug reports.

---

## 2. Releases

**None.** No new releases were published in this window. The changes below are merged/closed PRs that will presumably ship in the next release.

---

## 3. Project Progress

Closed/merged PRs in the last 24 hours show concentrated work on security, update reliability, and desktop/gateway fixes:

- **Security / profile isolation** — [PR #76462](https://github.com/NousResearch/hermes-agent/pull/76462) (closed) completes the profile secret-scope migration: `get_env_value`, the Anthropic adapter, and ACP env scrubbing now route through `agent.secret_scope.get_secret`. This directly addresses the cross-profile credential leaks in [#51603](https://github.com/NousResearch/hermes-agent/issues/51603) and [#75141](https://github.com/NousResearch/hermes-agent/issues/75141).
- **Update reliability** — [PR #76464](https://github.com/NousResearch/hermes-agent/pull/76464) (closed, P1) makes `hermes update` provision a managed Node runtime when system npm fails `engines.npm` checks. [PR #76250](https://github.com/NousResearch/hermes-agent/pull/76250) (closed) makes npm recovery resumable after no-op retries. [PR #75859](https://github.com/NousResearch/hermes-agent/pull/75859) (closed) relaxes the npm engine constraint to support npm 11.11.0.
- **Gateway fixes** — [PR #76465](https://github.com/NousResearch/hermes-agent/pull/76465) (closed) fixes relay thread-rename to carry the parent-channel discriminator.
- **Features** — [PR #76461](https://github.com/NousResearch/hermes-agent/pull/76461) (closed) adds support for custom `gateway.systemd_unit_name` values; [PR #7793](https://github.com/NousResearch/hermes-agent/pull/7793) (closed) adds auto-ingestion of Telegram PDF drops into the knowledge base; [PR #18621](https://github.com/NousResearch/hermes-agent/pull/18621) (closed) adds an opt-in `description_full` skill field to bypass 60-character truncation.
- **Profile cloning** — [PR #24183](https://github.com/NousResearch/hermes-agent/pull/24183) (closed) excludes `.archive` subtrees from profile clones, fixing copy errors from retired skill snapshots.

---

## 4. Community Hot Topics

The most active issues by comment count:

- [Issue #75598](https://github.com/NousResearch/hermes-agent/issues/75598) — **7 comments** (closed). Users report update-related instability starting ~1 week ago: multiple gateways running and conflicting with each other, with profile switching failing to deactivate other components.
- [Issue #51603](https://github.com/NousResearch/hermes-agent/issues/51603) — **5 comments** (closed). Cross-profile Anthropic token leak in multiplex mode. This was a serious security concern and is now addressed by the secret-scoping PR cluster.
- [Issue #43757](https://github.com/NousResearch/hermes-agent/issues/43757) — **3 comments** (open). Responses API `function_call_output` items are stripped from the `input` array, losing tool results across turns.
- [Issue #32887](https://github.com/NousResearch/hermes-agent/issues/32887) — **3 comments** (open). `gateway_state.json` heartbeat tick is missing for idle gateways, causing false liveness failures in cross-container WebUI deployments.
- [Issue #76064](https://github.com/NousResearch/hermes-agent/issues/76064) — **2 comments + 1 👍** (open). Desktop app ships demo/dogfood plugins enabled by default, creating production UI clutter.

PR comment counts were not included in the data extract, but the highest-signal PRs are the P1 update fix [#76464](https://github.com/NousResearch/hermes-agent/pull/76464) and the security-scope completion [#76462](https://github.com/NousResearch/hermes-agent/pull/76462).

Underlying need: users are running Hermes as a **multi-profile, multi-gateway infrastructure component**, and they care most about isolation, update safety, and gateway liveness. The community is also actively reporting desktop polish issues.

---

## 5. Bugs & Stability

Ranked by severity:

- **P1 — `hermes update` broken by npm `engines`/EBADENGINE**  
  Reported in [#75598](https://github.com/NousResearch/hermes-agent/issues/75598) and [#75651](https://github.com/NousResearch/hermes-agent/issues/75651). Fixed by PR [#76464](https://github.com/NousResearch/hermes-agent/pull/76464), with follow-up resumability work in open PR [#76467](https://github.com/NousResearch/hermes-agent/pull/76467) and npm constraint relaxation in [#75859](https://github.com/NousResearch/hermes-agent/pull/75859).

- **P2 — OmniRoute structured `chat_admission_busy` 503 aborts multi-agent turns**  
  [Issue #76468](https://github.com/NousResearch/hermes-agent/issues/76468) — Hermes treats this as a generic provider overload instead of waiting/retrying for capacity. No fix PR yet.

- **P2 — Termux install fails: `nemo-relay<0.7,>=0.6.0` not found**  
  [Issue #76469](https://github.com/NousResearch/hermes-agent/issues/76469) — needs reproduction. No fix PR yet.

- **P2 — Responses API tool results lost across turns**  
  [Issue #43757](https://github.com/NousResearch/hermes-agent/issues/43757) — `function_call_output` items in the `input` array are stripped. No fix PR yet.

- **P2 — Profile `.env` does not clear inherited Hermes env vars**  
  Closed as fixed in [#75141](https://github.com/NousResearch/hermes-agent/issues/75141) via the secret-scope migration.

- **P2 — Windows: system Node can override bundled Node**  
  [PR #76459](https://github.com/NousResearch/hermes-agent/pull/76459) is open to fix the Windows portable layout and PATH precedence.

- **P2 — Windows: device/proc read-guard is a no-op**  
  [PR #69403](https://github.com/NousResearch/hermes-agent/pull/69403) is open; `normpath` on Windows breaks the `/dev/*` blocklist comparison.

- **P3 — Desktop demo plugins enabled by default**  
  [Issue #76064](https://github.com/NousResearch/hermes-agent/issues/76064) — UI clutter from example/counter plugins in status bar.

- **P3 — `gateway_state.json` heartbeat missing for idle gateways**  
  [Issue #32887](https://github.com/NousResearch/hermes-agent/issues/32887) — open, labeled `needs-decision`.

- **P3 — `hermes honcho peers` shows `(not set)` for non-default profiles**  
  [Issue #76414](https://github.com/NousResearch/hermes-agent/issues/76414) — host key built with `.` instead of `_`.

- **P3 — Desktop terminal glyph corruption**  
  [PR #76463](https://github.com/NousResearch/hermes-agent/pull/76463) is open to refresh sibling terminals when the shared WebGL atlas is cleared.

Other fix PRs opened in the last 24 hours include multimodal-aware context estimation ([#76471](https://github.com/NousResearch/hermes-agent/pull/76471)), zero-chunk provider stall recovery ([#76473](https://github.com/NousResearch/hermes-agent/pull/76473)), and desktop clarify-card focus handling ([#76472](https://github.com/NousResearch/hermes-agent/pull/76472)).

---

## 6. Feature Requests & Roadmap Signals

New feature requests and feature PRs:

- **TTS sample-rate support** — [Issue #76466](https://github.com/NousResearch/hermes-agent/issues/76466): use the sample rate returned by OpenAI-compatible TTS endpoints instead of hardcoding 24 kHz. Likely to land soon; local TTS users are actively hitting this.
- **Per-turn provider request budget** — [PR #76458](https://github.com/NousResearch/hermes-agent/pull/76458) is a duplicate/feature PR for an opt-in budget on physical provider requests per user turn.
- **Config array decoding** — [PR #76470](https://github.com/NousResearch/hermes-agent/pull/76470) would make `hermes config set` decode JSON array literals into YAML lists, fixing allowlist configs. Labeled `needs-decision`.
- **Web fallback-provider management** — [PR #55170](https://github.com/NousResearch/hermes-agent/pull/55170) adds a Fallback Chain panel to the dashboard for managing `fallback_providers`.
- **Custom systemd unit names** — already closed in [PR #76461](https://github.com/NousResearch/hermes-agent/pull/76461), likely to appear in the next release.
- **Better skill descriptions** — `description_full` opt-in from [PR #18621](https://github.com/NousResearch/hermes-agent/pull/18621) is also closed and likely shipping.

Prediction: the next release will probably include the update/Node runtime recovery work, the desktop Windows PATH fix, config array-literal handling, and the TTS sample-rate change.

---

## 7. User Feedback Summary

- **Update regression pain is real.** [Issue #75598](https://github.com/NousResearch/hermes-agent/issues/75598) says updates went smoothly until about a week ago, then "problems started making whole program unstable." This matches the `engines.npm`/EBADENGINE cluster and profile/gateway conflict reports.
- **Profile isolation is a top trust concern.** Users of multiplex gateways explicitly worry about credentials leaking between profiles, as seen in [#51603](https://github.com/NousResearch/hermes-agent/issues/51603) and [#75141](https://github.com/NousResearch/hermes-agent/issues/75141).
- **Installation friction on non-mainstream platforms.** Reports include Termux failing on `nemo-relay` ([#76469](https://github.com/NousResearch/hermes-agent/issues/76469)), macOS npm EBADENGINE and LXC/macOS git-sync freezes ([#75651](https://github.com/NousResearch/hermes-agent/issues/75651)), and Windows Node precedence ([#76459](https://github.com/NousResearch/hermes-agent/pull/76459)).
- **Desktop users want production-quality defaults.** The demo/dogfood plugin clutter in [#76064](https://github.com/NousResearch/hermes-agent/issues/76064) is a small but clear signal that desktop polish matters as Hermes expands beyond CLI/power users.
- **API compatibility matters.** The Responses API tool-result loss ([#43757](https://github.com/NousResearch/hermes-agent/issues/43757)) and TTS sample-rate hardcoding ([#76466](https://github.com/NousResearch/hermes-agent/issues/76466)) show users are building real integrations on Hermes's gateway and tooling surface.

---

## 8. Backlog Watch

These items have been open for a while and could use maintainer attention:

- [Issue #43757](https://github.com/NousResearch/hermes-agent/issues/43757) — Responses API `function_call_output` stripped from `input` array. Open since **2026-06-10**, P2, with no fix PR yet. This breaks tool-calling workflows across turns.
- [Issue #32887](https://github.com/NousResearch/hermes-agent/issues/32887) — Missing heartbeat tick in `gateway_state.json`. Open since **2026-05-27**, P3, labeled `needs-decision`. Affects WebUI liveness for idle gateways.
- [PR #55170](https://github.com/NousResearch/hermes-agent/pull/55170) — Fallback-provider management UI. Open since **2026-06-29**; useful feature waiting for review/merge.
- [PR #69403](https://github.com/NousResearch/hermes-agent/pull/69403) — Windows device/proc read-guard fix. Open since **2026-07-22**, P2; important for Windows security boundaries.
- [PR #76470](https://github.com/NousResearch/hermes-agent/pull/76470) — Config array-literal decoding. New but already labeled `needs-decision`; could affect allowlist configuration behavior.

Overall, the project is in a **healthy, fast-moving bug-fixing phase**, with maintainers closing security-critical work quickly. The main ongoing risk is the update/install path across multiple platforms, and the main user-facing opportunity is reducing desktop and gateway integration friction.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest — 2026-08-02

## 1. Today's Overview
PicoClaw saw low but steady activity over the last 24 hours: 1 open issue and 3 pull requests were updated, with no new releases published. The most significant open problem remains a stale Matrix sync reliability bug (#3203), which continues to gather community attention. On the contribution side, a Traditional Chinese locale PR was closed/merged, while two new provider integrations are waiting for review. Overall project health is stable, though bug triage and stale-item handling deserve maintainer focus.

## 2. Releases
No new releases were published on 2026-08-02. The latest release data is empty, so there are no changelog, breaking-change, or migration notes to report.

## 3. Project Progress
- [#3261 — Add zh-TW locale and Traditional Chinese translations](https://github.com/sipeed/picoclaw/pull/3261)  
  Closed/merged in the last 24 hours. This PR extends localization to Traditional Chinese / Taiwanese terminology across the WebUI and documentation.

No other PRs were merged. Two feature PRs remain open:
- [#3299 — Add native Exa web search provider](https://github.com/sipeed/picoclaw/pull/3299)
- [#3309 — feat(providers): add OrcaRouter as an OpenAI-compatible provider](https://github.com/sipeed/picoclaw/pull/3309)

## 4. Community Hot Topics
- [#3203 — [BUG] Matrix sync loop has no reconnection logic — silent death after network/server disruption](https://github.com/sipeed/picoclaw/issues/3203)  
  The most active item: 7 comments and 2 👍 reactions. Users report that the Matrix `/sync` long-polling loop dies permanently after any network or server disruption, and because the main process remains alive, systemd `Restart=on-failure` never triggers. This points to a clear demand for built-in reconnection and health-check logic.

The open provider PRs (#3299, #3309) have not yet attracted comments, but they signal growing interest in expanding PicoClaw's native web-search and OpenAI-compatible provider ecosystem.

## 5. Bugs & Stability
- **High severity:** [#3203 — Matrix sync loop no reconnection logic](https://github.com/sipeed/picoclaw/issues/3203)  
  Silent death of Matrix sync after network/homeserver disruption. No automatic recovery and no systemd restart because the process stays alive. This can cause undetected loss of messaging functionality. The issue is stale-labeled and no fix PR is currently linked.

No other bugs, crashes, or regressions were reported in the last 24 hours.

## 6. Feature Requests & Roadmap Signals
- **Exa web search provider** ([#3299](https://github.com/sipeed/picoclaw/pull/3299)): Adds native `tools.web` / `web_search` support with date-range filters.
- **OrcaRouter provider** ([#3309](https://github.com/sipeed/picoclaw/pull/3309)): Adds a first-class OpenAI-compatible multi-vendor router provider.
- **Traditional Chinese localization** ([#3261](https://github.com/sipeed/picoclaw/pull/3261)): Signals ongoing demand for broader i18n coverage.

If these PRs move forward, the next version could include expanded provider options, improved search capabilities, and better localization — alongside a likely reliability fix for Matrix sync.

## 7. User Feedback Summary
The clearest user pain point is Matrix sync reliability: after any network disruption or homeserver restart, the sync loop dies silently and does not reconnect, leaving the service unusable without manual intervention. Users express frustration that systemd cannot recover the process because the process itself stays alive.

On a positive note, community members are actively contributing localization and new provider integrations, indicating healthy external interest in expanding PicoClaw's feature set.

## 8. Backlog Watch
- [#3203 — Matrix sync loop no reconnection logic](https://github.com/sipeed/picoclaw/issues/3203)  
  Open since 2026-07-02, updated 2026-08-01, stale-labeled, with 7 comments. This is the most important item needing maintainer attention, either as a fix or explicit triage decision.
- [#3261 — Add zh-TW locale](https://github.com/sipeed/picoclaw/pull/3261)  
  Now closed, but stale-labeled; maintainers should confirm it was properly merged or clearly document closure rationale.
- [#3299 — Exa web search provider](https://github.com/sipeed/picoclaw/pull/3299) and [#3309 — OrcaRouter provider](https://github.com/sipeed/picoclaw/pull/3309)  
  Open and awaiting review; these represent low-risk feature contributions that could be moved forward.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-08-02

## 1. Today's Overview

NanoClaw saw a high-activity day: 16 pull requests touched in the last 24 hours (10 still open, 6 merged/closed), 2 issues updated, and a new release cut. The headline is **v2.1.54**, a rollup release that ships a breaking unification of iMessage support — folding the previously separate local and hosted integrations into a single `imessage` channel. Maintainer velocity is strong: reporter **glifocat** filed two issues (#3169, #3171) and immediately followed each with fix PRs (#3170, #3172), both of which advanced or closed within the same day. The cross-sectional PR age suggests some long-running fixes are still awaiting review, but the project remains responsive to fresh regressions and feature requests.

## 2. Releases

**v2.1.54** — *Rollup release covering v2.1.18 through v2.1.54; everything merged since the v2.1.17 tag.*

- **[BREAKING] iMessage unified into a single `imessage` channel with two backends**, installed via one skill (`/add-imessage`):
  - **Local** — this Mac's `chat.db` via the Chat SDK
  - **Hosted** — native Photon (via `spectru...` — description truncated in data)
- **Migration note**: Existing separate iMessage setups should be re-installed through `/add-imessage` and reconfigured with the chosen backend. This change supersedes the earlier standalone channel PRs (#2999), with a corrected registration flow delivered in #3164.

No other release notes were provided beyond the rollup summary.

## 3. Project Progress

Six PRs were merged/closed in the last 24 hours:

- **[#3170 — fix(setup): dispatch failure assist to the picked provider](https://github.com/nanocoai/nanoclaw/pull/3170)** — Merged. Fixes the setup flow so failure diagnostics go to the provider the operator actually selected (e.g., Codex) instead of always offering the Claude CLI. Directly resolves issue #3169.
- **[#3168 — fix(release): close post-merge safety gaps](https://github.com/nanocoai/nanoclaw/pull/3168)** — Merged. Hardens the release pipeline against post-merge edge cases.
- **[#3167 — feat(credentials): alert when a provider credential expires](https://github.com/nanocoai/nanoclaw/pull/3167)** — Merged. Adds proactive alerts for expiring provider credentials (motivated by a real Codex ChatGPT credential expiry on 2026-08-01 that surfaced only as a cryptic WhatsApp reconnect error).
- **[#2999 — feat(channels): unify iMessage into a single `imessage` channel (local + hosted backends)](https://github.com/nanocoai/nanoclaw/pull/2999)** — Closed. Original unification PR, superseded by #3164.
- **[#3164 — Hosted iMessage (Photon): supersede #2999 with a working registration flow](https://github.com/nanocoai/nanoclaw/pull/3164)** — Closed. Lands the shipped version of hosted iMessage with a functional registration flow.
- **[#3165 — Codex/copilot changes](https://github.com/nanocoai/nanoclaw/pull/3165)** — Closed. Provider-related changes, details not fully specified.

## 4. Community Hot Topics

- **iMessage unification saga (#2999 → #3164)** — The original unified-channel PR ([#2999](https://github.com/nanocoai/nanoclaw/pull/2999), opened 07-10) was ultimately superseded by [#3164](https://github.com/nanocoai/nanoclaw/pull/3164) (opened 07-31) because the first lacked a working hosted registration flow. Both closed on 08-01 and the result shipped in v2.1.54. This shows a fast internal iteration loop: an incomplete feature was reworked and released within three weeks rather than lingering.
- **Qodo skills controversy ([#3171](https://github.com/nanocoai/nanoclaw/issues/3171))** — The bundled `get-qodo-rules` and `qodo-pr-resolver` skills require a Qodo SaaS API key from `~/.qodo/config.json` that nothing in the repo provisions, and they intercept normal coding requests. The community fix ([#3172](https://github.com/nanocoai/nanoclaw/pull/3172)) simply removes both skills; it remains open for review.
- **Credential expiry invisibility ([#3167](https://github.com/nanocoai/nanoclaw/pull/3167))** — The PR description paints a real-world pain point: an expired Codex credential produced only `Error: Reconnecting... 2/5: Read-only file system (os error 30)` in WhatsApp, with nothing actionable in `nanoclaw.error.log`. This prompted a new alerting feature — a good example of operational feedback driving product changes.

## 5. Bugs & Stability

Ranked by severity:

1. **[#3166 — fix(migrate-v2): call `insertTaskRow`, not the removed `insertTask`](https://github.com/nanocoai/nanoclaw/pull/3166)** — *High, open.* `setup/migrate-v2/tasks.ts` imports `insertTask`, which no longer exists; the module exports `insertTaskRow`. Because it's a static ESM import, the migration step dies with a `SyntaxError` before executing anything. This fully breaks the v2 migration path. Fix PR is open.
2. **[#3171 — Qodo skills intercept normal coding requests](https://github.com/nanocoai/nanoclaw/issues/3171)** — *High, open.* Two bundled skills activate during normal coding workflows but depend on an integration nothing sets up, causing unexpected behavior for all users. Removal PR #3172 is open.
3. **[#3174 — Agent containers unusable on rootless Docker](https://github.com/nanocoai/nanoclaw/pull/3174)** — *Medium-high, open.* Two independent failures, both invisible when the host user is in the `docker` group. The author deliberately kept the agent account out of the group and hit both. Fix PR is open.
4. **[#3169 — Setup failures always offer the Claude CLI on non-Claude installs](https://github.com/nanocoai/nanoclaw/issues/3169)** — *Medium, closed.* Setup wrongly suggests installing Claude CLI (and even starts Anthropic sign-in) when the operator picked Codex or another provider. Fixed by merged PR #3170.
5. **[#2956 — Duplicate delivery when final output repeats tool-sent content](https://github.com/nanocoai/nanoclaw/pull/2956)** — *Medium, open since 07-05.* Agents that send a reply via `send_message` and restate it in their final output deliver the message twice. Fix PR still waiting.
6. **[#2750 — Stale `outbound.db` journals after container kills](https://github.com/nanocoai/nanoclaw/pull/2750)** — *Open since 06-12.* Host-side READONLY `outbound.db` handles get stale journals after a ceiling/claim-stuck SIGKILL, causing poll races. Fix PR covers both #2516 and #2640.

## 6. Feature Requests & Roadmap Signals

- **Credential expiry alerts** ([#3167](https://github.com/nanocoai/nanoclaw/pull/3167)) — Already merged; proactive credential monitoring is now part of the project.
- **Rootless Docker support** ([#3174](https://github.com/nanocoai/nanoclaw/pull/3174)) — A clear infrastructure gap for container-based agent deployment; likely to merge soon given its precise diagnosis.
- **Best-effort reaction delivery** ([#3121](https://github.com/nanocoai/nanoclaw/pull/3121)) — Open since 07-23; makes reaction delivery non-fatal, improving channel robustness.
- **Egress update** ([#3173](https://github.com/nanocoai/nanoclaw/pull/3173)) — Open; likely updates network egress configuration or docs.
- **Router hardening for untrusted input** ([#2801](https://github.com/nanocoai/nanoclaw/pull/2801)) — `safeParseContent` can return primitives instead of objects, causing undefined reads; a security-adjacent hardening that has been open since 06-17.
- **Hosted iMessage (Photon backend)** — Shipped in v2.1.54, signalling that hosted channel backends are an active strategic direction.

## 7. User Feedback Summary

- **Setup flow frustration** (#3169): Choosing a non-Claude provider still routes diagnostics through the Claude CLI, and can trigger Anthropic sign-in. This is a poor first-run experience; users want provider-matched diagnostics. Fixed same-day.
- **Bundled skill interference** (#3171): Two shipped skills assume a Qodo account that is never set up, and they intercept normal coding requests — a trust/UX problem for default installs.
- **Cryptic operational failures** (#3167): Credential expiry manifested as an unrelated "Read-only file system (os error 30)" reconnect error in WhatsApp; operators had no actionable signal. The merged alerting feature directly addresses this.
- **Container deployment friction** (#3174): Rootless Docker users are silently locked out of agent containers depending on host group membership — a portability concern for modern container setups.
- **Duplicate message delivery** (#2956): Agents using `send_message` MCP tooling can double-send outputs to users, which is confusing in chat channels.
- Overall satisfaction signal is positive: community-contributed PRs are abundant (8 distinct external authors across 16 PRs), and maintainers close issue→fix loops within hours.

## 8. Backlog Watch

PRs and issues that have been open longest and need maintainer attention:

- **[#2750 — Recover stale `outbound.db` journals; classify hot-journal poll races](https://github.com/nanocoai/nanoclaw/pull/2750)** — Open since **06-12**; addresses two diagnosed failure modes (#2516, #2640). No merge signal after ~7 weeks.
- **[#2801 — Harden untrusted router input (`safeParseContent` + `engage_pattern`)](https://github.com/nanocoai/nanoclaw/pull/2801)** — Open since **06-17**; cross-cutting robustness fix for parsing primitive payloads.
- **[#2956 — Suppress duplicate delivery in agent-runner](https://github.com/nanocoai/nanoclaw/pull/2956)** — Open since **07-05**; user-facing duplicate-message bug with a ready fix.
- **[#3046 — Docs: align init-first-agent with current status blocks](https://github.com/nanocoai/nanoclaw/pull/3046)** — Open since **07-14**; documentation drift correction.
- **[#3090 — Fix(templates): prepend all top-level context Markdown](https://github.com/nanocoai/nanoclaw/pull/3090)** — Open since **07-19**; template consistency fix from a core-team author.
- **[#3121 — Make reaction delivery best-effort](https://github.com/nanocoai/nanoclaw/pull/3121)** — Open since **07-23**; channel resilience improvement.
- **[#3172 — Remove the two qodo skills](https://github.com/nanocoai/nanoclaw/pull/3172)** — New but time-sensitive: it resolves an active user-facing bug (#3171) and should ideally land with or before the next release.

---

*Digest generated from GitHub activity for 2026-08-02. All links reference [nanocoai/nanoclaw](https://github.com/nanocoai/nanoclaw).*

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-08-02

## 1. Today's Overview

IronClaw activity on 2026-08-02 is healthy and maintainer-driven: 20 issues and 25 PRs were updated in the last 24 hours. Of those, 17 issues and 16 PRs remain open, while 3 issues and 9 PRs were closed or merged. The dominant themes are the ongoing “Reborn” Wave 2 architecture refactor — dependency inversion onto `ironclaw_product_contracts` — and prompt-cache/compaction work from the pi-harness adoption program. No new releases were published. The pulse is consistent with a project mid-migration: high PR throughput, careful CI-gate hardening, and disciplined issue tracking.

## 2. Releases

**No new releases were published in this window.** The latest release data is empty. A release PR remains open from earlier — [#5598](https://github.com/nearai/ironclaw/pull/5598) — proposing `ironclaw_common` 0.5.0, `ironclaw_safety` 0.2.3, and `ironclaw_skills` 0.4.0, but it has not yet merged.

## 3. Project Progress

The visible closed/merged PR set shows steady progress on the Reborn refactor and CI hardening:

- [#6998](https://github.com/nearai/ironclaw/pull/6998) — **Merged/closed**: WS2.1 contract inversion: `ironclaw_extension_host` now implements `ironclaw_product_contracts` ports instead of `ironclaw_product` ports. Behavior-free refactor.
- [#7002](https://github.com/nearai/ironclaw/pull/7002) — **Merged/closed**: WS5 port inversion for `webui` + `openai_compat` onto `product_contracts`.
- [#6996](https://github.com/nearai/ironclaw/pull/6996) — **Merged/closed**: CI-gate hardening that closes [#6963](https://github.com/nearai/ironclaw/issues/6963), replacing path-keyed gates with inventory-driven, fail-closed discovery.
- [#6995](https://github.com/nearai/ironclaw/pull/6995) — **Merged/closed**: Wave 1 target-architecture truth audit, reconciling decision records with shipped reality.
- [#6761](https://github.com/nearai/ironclaw/pull/6761) — **Merged/closed**: Regression test for generic outbound registration in the channel host.

Issues closed today include:

- [#6963](https://github.com/nearai/ironclaw/issues/6963) — Path-keyed CI gates surviving `#6946` (closed by #6996).
- [#6921](https://github.com/nearai/ironclaw/issues/6921) — Extract neutral loop/extension/product contracts and seal evidence minting.
- [#6903](https://github.com/nearai/ironclaw/issues/6903) — Admin users list unable to load beyond first page.

The open WS2/Wave 2 stack is still advancing: [#7000](https://github.com/nearai/ironclaw/pull/7000), [#7003](https://github.com/nearai/ironclaw/pull/7003), [#7004](https://github.com/nearai/ironclaw/pull/7004), and [#7005](https://github.com/nearai/ironclaw/pull/7005) are all active. Cache-work PRs [#6997](https://github.com/nearai/ironclaw/pull/6997) and [#7001](https://github.com/nearai/ironclaw/pull/7001) are also open and tied to P0 cache-stability issues.

## 4. Community Hot Topics

Among issues with explicit comment counts, the most active are:

- [#6963](https://github.com/nearai/ironclaw/issues/6963) — **7 comments**. Tracking issue for eight path-keyed CI/dev gates that `#6946` did not rewrite. Closed by #6996. The discussion reflects a strong maintainer concern about silent CI no-ops and fail-closed gate discovery.
- [#6974](https://github.com/nearai/ironclaw/issues/6974) — **2 comments**. libSQL `thread_store_writes` pathology: tool-heavy stress cases at p95 37–135s post-`#6696`. Split out of the larger Postgres capacity issue [#6973](https://github.com/nearai/ironclaw/pull/6973). Underlying need is performance recovery after the row-native process journal rewrite.
- [#6921](https://github.com/nearai/ironclaw/issues/6921) — **2 comments**, now closed. The “neutral contracts” extraction is a sign of the project’s dependency-boundary cleanup maturing.

No 👍 reactions were recorded on any listed item, and the PR sample did not include explicit comment counts, so issue activity is the strongest engagement signal in this window.

## 5. Bugs & Stability

Ranked roughly by severity:

| Severity | Item | Description | Fix / Status |
|---|---|---|---|
| **P0 — cache correctness/performance** | [#6984](https://github.com/nearai/ironclaw/issues/6984), [#6985](https://github.com/nearai/ironclaw/issues/6985), [#6986](https://github.com/nearai/ironclaw/issues/6986) | Missing explicit Anthropic `cache_control` breakpoints; the cached prompt prefix mutates; advertised tool array is byte-unstable. These invalidate prompt caches and hurt latency/cost. | [#6997](https://github.com/nearai/ironclaw/pull/6997) fixes #6984; [#7001](https://github.com/nearai/ironclaw/pull/7001) fixes #6985; #6986 still open. |
| **P1 — performance regression** | [#6974](https://github.com/nearai/ironclaw/issues/6974) | libSQL `thread_store_writes` pathology: p95 37–135s in tool-heavy stress cases after #6696. | Open. Related to [#6973](https://github.com/nearai/ironclaw/pull/6973), but no dedicated fix PR yet. |
| **P1 — token accounting bug** | [#6989](https://github.com/nearai/ironclaw/issues/6989) | `ModelWorkRequest::for_assistant` estimates input tokens from `content_ref.as_str().len()` — the reference string, not the referenced content. | Open. |
| **CI structural bug** | [#6978](https://github.com/nearai/ironclaw/issues/6978) | `workflow_dispatch` runs of `reborn-tests.yml` fail the Tests (Reborn) roll-up because `critical-mutation` is skipped but disallowed. | Open. |
| **CI locale bug** | [#6992](https://github.com/nearai/ironclaw/pull/6992) | `comm` input sorting uses `LC_ALL=C`, but `comm` itself runs in ambient locale, causing crate-discovery failure under UTF-8 collation. | Fix PR open. |
| **Coverage gate blocker** | [#7006](https://github.com/nearai/ironclaw/issues/7006) | Queued-message steering PR [#5981](https://github.com/nearai/ironclaw/pull/5981) trips the changed-lines integration coverage gate on fault-injection and error paths the hermetic harness cannot execute. | Open; likely needs gate redesign. |
| **Pre-existing findings in moved code** | [#7011](https://github.com/nearai/ironclaw/issues/7011) | WS2.4 split surfaced five findings in `ironclaw_extension_manager`, including a false `WriteFilesystem` effect and an untested lock predicate. | Open. |
| **Architecture rule gap** | [#6999](https://github.com/nearai/ironclaw/issues/6999) | `reborn_dependency_boundaries` server-lifecycle rule never covered the WebChat v2 route surface it documents. | Open; recorded as an architecture decision, not a silent fix. |

No crashes or security-related issues were reported in this window.

## 6. Feature Requests & Roadmap Signals

- [#7009](https://github.com/nearai/ironclaw/issues/7009) — **Add OrcaRouter as a built-in LLM provider.** `providers.json` already has OpenRouter, Together, Fireworks, Cerebras, SambaNova, NVIDIA, Venice, io.net, and Yandex; OrcaRouter users currently need a fallback path. Likely a small, high-value provider addition.
- [#6983](https://github.com/nearai/ironclaw/issues/6983) — **Add `hub` as an alias for `ironhub`.** User feedback from IronHub dashboard documentation; `ironhub` and `iron-hub` exist, but `hub` does not. This is a simple CLI-compatibility enhancement and could land in a next minor release.
- [#6993](https://github.com/nearai/ironclaw/issues/6993) — **Backend wiring for the OOBE automation-tasks prototype.** The UI prototype itself landed in [#6994](https://github.com/nearai/ironclaw/pull/6994) as a mock-data-only first-run experience; this issue tracks making it real.
- [#7007](https://github.com/nearai/ironclaw/pull/7007) — **Alert live-canary Slack on merge-queue failures.** Internal CI observability improvement.
- [#7008](https://github.com/nearai/ironclaw/issues/7008) — **Split the `product_wire` DTO family.** The file is 1,923 lines and needs owner cleanup to remove its `large_file` exemption.

Roadmap signal: the next likely batch of merged work is the P0 cache-stability pair [#6997](https://github.com/nearai/ironclaw/pull/6997) + [#7001](https://github.com/nearai/ironclaw/pull/7001), followed by the remaining WS2/Wave 2 refactor stack. OrcaRouter and the `hub` alias are the most user-visible small features currently queued.

## 7. User Feedback Summary

Explicit user-facing feedback in this window is modest:

- **OrcaRouter support missing** — [#7009](https://github.com/nearai/ironclaw/issues/7009): users reaching for common multi-provider gateways expect a first-class entry in `providers.json`.
- **`hub` CLI alias missing** — [#6983](https://github.com/nearai/ironclaw/issues/6983): a user preparing IronHub CLI release documentation found the `ironhub` subcommand awkwardly named and expected `hub` as an alias.
- **Admin pagination bug** — [#6903](https://github.com/nearai/ironclaw/issues/6903): admins could not load the user list beyond the first 100 users. This is now closed, so it appears resolved.

There is also product-level feedback embedded in [#6993](https://github.com/nearai/ironclaw/issues/6993): the OOBE automation-tasks prototype is only a UI shell on mock data, and users cannot yet act on it. No satisfaction metrics or 👍 reactions were present in the dataset, so overall sentiment is best described as “maintainer-led, with a few concrete user-pain-point feature requests.”

## 8. Backlog Watch

These items have been open for a while or are blocking important work and need maintainer attention:

- [#5598](https://github.com/nearai/ironclaw/pull/5598) — **Release PR open since 2026-07-03** (~30 days). Contains breaking changes for `ironclaw_common` and `ironclaw_skills`; needs a merge/release decision.
- [#5981](https://github.com/nearai/ironclaw/pull/5981) — **Queued-message steering**, open since 2026-07-11. Large Reborn feature; currently blocked by the coverage-gate issue [#7006](https://github.com/nearai/ironclaw/issues/7006).
- [#5982](https://github.com/nearai/ironclaw/pull/5982) — **Budget approval-as-blocked-gate + usage settings**, open since 2026-07-11. Stacked on #5981 and waiting on that review path.
- [#6780](https://github.com/nearai/ironclaw/pull/6780) — **IronHub deep-link register/install gateway + private manifest source**, open since 2026-07-28. Re-port of #5409; substantive feature needing review.
- [#6917](https://github.com/nearai/ironclaw/pull/6917) — **WebUI workspace file links in authenticated previews**, open since 2026-07-30. User-facing fix with security-relevant link handling.
- [#6973](https://github.com/nearai/ironclaw/pull/6973) — **Postgres API capacity recovery**, open since 2026-07-31. Large performance PR closely tied to [#6974](https://github.com/nearai/ironclaw/issues/6974); important for hosted capacity stability.

Overall, IronClaw is in a healthy but busy state: architecture migration is progressing, CI quality is being defended aggressively, and performance/cache regressions are receiving P0-level attention.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-08-02

## 1. Today's Overview

As of 2026-08-02, LobsterAI is in a low-activity maintenance and triage phase. 7 issues were touched in the last 24 hours, but 6 of them were automatically closed as **stale**, leaving only 1 open/active issue. 2 pull requests were updated and remain open, with **no PRs merged/closed** and **no new releases**. The overall signal is that the project is performing stale-cleanup and backlog hygiene rather than shipping new features. The two open PRs (#1224, #2358) still await maintainer review, and one long-standing i18n/UX issue (#1223) remains unresolved.

## 2. Releases

No new releases were published in the last 24 hours. This section is omitted.

## 3. Project Progress

No PRs were merged or closed in the last 24 hours. Active development is limited to two open PRs:

- [PR #1224](https://github.com/netease-youdao/LobsterAI/pull/1224) — `fix(agent)`: i18n hardcoding fix, Escape key support for Agent modals, and double-click protection when deleting. This PR is linked to issue #1223 but has not been merged.
- [PR #2358](https://github.com/netease-youdao/LobsterAI/pull/2358) — `fix(cowork)`: shows localized feedback when session rename fails, preventing silent failures. Also not yet merged.

Six stale issues were closed, including a feature-oriented issue for code-block line numbers (#1302). No feature work landed.

## 4. Community Hot Topics

The most active item is [Issue #1293](https://github.com/netease-youdao/LobsterAI/issues/1293) (2 comments, 1 👍): custom Studio HTTP MCP servers are not actually updated in the OpenClaw engine, so only SSE-based MCPs can be used. This was the only issue with a positive reaction and reflects a practical integration gap.

The only currently open issue, [Issue #1223](https://github.com/netease-youdao/LobsterAI/issues/1223) (1 comment), covers i18n/UX problems: hardcoded Chinese labels in prompts, missing Escape key close for Agent dialogs, and missing delete double-click protection. It is paired with PR #1224.

The stale-closed issues each had 2 comments, mostly likely from maintainer/stale-bot interaction rather than deep community discussion. No PR comment activity was recorded.

## 5. Bugs & Stability

The following bugs were touched in the last 24 hours. Most were closed as stale, so severity is based on user impact, not fix availability.

- **High — [Issue #1296](https://github.com/netease-youdao/LobsterAI/issues/1296): Uploading a 3MB long image crashes the page and makes new tasks consistently fail.** This is a major stability problem because it can make the whole app unusable. No linked fix PR was found; the issue was closed as stale.
- **High/Medium — [Issue #1307](https://github.com/netease-youdao/LobsterAI/issues/1307): After closing a model provider config panel, other provider configs become read-only and cannot be edited.** This is a clear UI state bug. No fix PR found.
- **Medium — [Issue #1298](https://github.com/netease-youdao/LobsterAI/issues/1298): Model test connection succeeds, but entering a short question triggers "input too long" and exceeds model limits.** This suggests a broken token-counting or context-limit path.
- **Medium — [Issue #1293](https://github.com/netease-youdao/LobsterAI/issues/1293): Custom HTTP MCP servers are not recognized by the OpenClaw engine; only SSE works.** Core MCP interoperability gap.
- **Low — [Issue #1305](https://github.com/netease-youdao/LobsterAI/issues/1305): Scheduled task history shows the wrong title after the task is run successfully and deleted.**
- **Low/UX — [Issue #1223](https://github.com/netease-youdao/LobsterAI/issues/1223): Hardcoded Chinese strings leak into English prompts; missing Escape-key close and delete double-click protection.** A fix PR exists (#1224) but remains open.

## 6. Feature Requests & Roadmap Signals

The most visible feature request was [Issue #1302](https://github.com/netease-youdao/LobsterAI/issues/1302): add a line-number toggle button to code blocks, including language-tagged and plain code blocks. This was closed as stale, suggesting it may not be prioritized soon.

The i18n/UX changes in [Issue #1223](https://github.com/netease-youdao/LobsterAI/issues/1223) and [PR #2358](https://github.com/netease-youdao/LobsterAI/pull/2358) are the most likely near-term roadmap signals. If maintainers resume feature work, #1224 and #2358 could land in the next release, while the line-number feature may need to be resubmitted or reopened.

## 7. User Feedback Summary

User-reported pain points in the last 24 hours include:

- MCP integration is incomplete: custom HTTP MCP servers do not reach the OpenClaw engine, only SSE works.
- Long image uploads can crash the page and poison subsequent tasks.
- Model context-length validation gives false positives even for short input.
- English users still encounter hardcoded Chinese text in AI prompts.
- Session rename failures are silent, leaving users confused.
- The model provider edit panel can become permanently read-only after closing it once.

No positive feedback or explicit satisfaction signals appeared in the data. The tone is mostly bug-reporting with screenshots, indicating real user friction. The stale-closing of several unfixed bugs without visible resolution may increase user frustration.

## 8. Backlog Watch

The following items need maintainer attention and are at risk of being swept away by stale bot:

- [Issue #1223](https://github.com/netease-youdao/LobsterAI/issues/1223) — Open since April 1, still open, with a fix PR waiting. Needs review or an explicit decision.
- [PR #1224](https://github.com/netease-youdao/LobsterAI/pull/1224) — Open since April 1, no comments, no merge. Directly resolves #1223.
- [PR #2358](https://github.com/netease-youdao/LobsterAI/pull/2358) — Open since July 18, no comments, no merge. Fixes a real UX issue (#670).
- [Issue #1296](https://github.com/netease-youdao/LobsterAI/issues/1296) and [Issue #1307](https://github.com/netease-youdao/LobsterAI/issues/1307) — Closed as stale despite being potentially serious bugs. They may need reopening if still reproducible.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest — 2026-08-02

## Today's Overview

Over the last 24 hours, Moltis saw no issue updates and zero new releases, but meaningful pull request activity: 4 PRs were updated, with 3 merged/closed and 1 still open. The merged work advances security boundaries, observability/instrumentation, and Nostr-based group chat support. The single open PR addresses a session-management fix and appears ready for maintainer review. Overall, the project is in an active integration phase, with more merged code than open discussion.

## Releases

No new releases were published in this period.

## Project Progress

Three PRs were merged/closed in the last 24 hours:

- **#1174 — Add instrumentation and feedback collection infrastructure** ([PR #1174](https://github.com/moltis-org/moltis/pull/1174))  
  Merges backend-neutral agent instrumentation, Langfuse v4 export, OTLP backend support, and end-user reaction feedback. Includes streaming/non-streaming parity, provider failover attribution, cache-aware token usage, and reasoning tracking.

- **#1170 — Gate /sh and privileged tools behind a per-account operators list** ([PR #1170](https://github.com/moltis-org/moltis/pull/1170))  
  Closes a privilege separation gap by separating channel access from operator privileges. Enforces the new boundary across commands, callbacks, queue replay, chat execution, and external integrations.

- **#1168 — Add NIP-29 group chat support for Buzz channels** ([PR #1168](https://github.com/moltis-org/moltis/pull/1168))  
  Adds Nostr NIP-29 group chat support, enabling `moltis-nostr` to work with Block’s Buzz workspace over NIP-42-authenticated connections.

Also active:

- **#1182 — fix(sessions): allow deleting and archiving the main session** ([PR #1182](https://github.com/moltis-org/moltis/pull/1182))  
  Open PR that removes the special `main` session guard, aligning it with normal session deletion/archival rules while preserving the active-channel-session restriction.

## Community Hot Topics

No issues or PRs recorded explicit comment/reaction counts in this window. The most notable activity is the open PR **#1182** ([PR #1182](https://github.com/moltis-org/moltis/pull/1182)), which directly addresses user-facing session management limitations. The underlying need is straightforward: users expect the main session to behave like any other session for deletion/archiving, without special protections. The three merged PRs also reflect community-facing concerns:

- Stronger operator/privilege controls (#1170)
- Better observability and user feedback (#1174)
- Nostr/NIP-29 interoperability for Buzz (#1168)

## Bugs & Stability

No new bug reports or regressions were filed in the last 24 hours. However, two PRs relate to stability/security:

- **High:** #1170 ([PR #1170](https://github.com/moltis-org/moltis/pull/1170)) fixes a privilege escalation path where channel senders with access allowlist status could reach privileged commands and host tools. This was a security-boundary bug, now fixed by an explicit per-account `operators` list.
- **Medium:** #1182 ([PR #1182](https://github.com/moltis-org/moltis/pull/1182), open) addresses the `main` session deletion/archival restriction from issue #1132 ([#1132](https://github.com/moltis-org/moltis/issues/1132)). The fix is implemented but not yet merged.

## Feature Requests & Roadmap Signals

The merged PRs suggest the upcoming release may include:

- **Agent instrumentation and feedback collection** (#1174): expects Langfuse v4 export and OTLP backend support; signals a focus on production observability and user reaction capture.
- **Nostr NIP-29 / Buzz group channels** (#1168): signals continued investment in Nostr-based team collaboration and opens the door to broader NIP-29 workspace integration.
- **Per-account operators / privilege separation** (#1170): a governance feature likely to be prominent in the next release, especially for externally-facing channels.

No explicit feature requests from users were filed in this period.

## User Feedback Summary

No direct user issues or comments were recorded in the last 24 hours. Indirect signals from PRs indicate:

- Users wanted to delete/archive the `main` session like any other session (#1182 / #1132).
- Operators wanted a clearer separation between channel access and privileged tool use (#1170).
- There is demand for self-hosted Nostr group chat support, specifically for Buzz/BLIP-29 workflows (#1168).

Overall, the feedback direction is toward more flexible session management, stricter security controls, and richer interoperability/observability for production deployments.

## Backlog Watch

No long-unanswered issues or PRs were visible in this data. The only open PR, #1182 ([PR #1182](https://github.com/moltis-org/moltis/pull/1182)), is recent and likely awaiting maintainer review; it should be tracked for merge in the next cycle.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest — 2026-08-02

## Today's Overview

QwenPaw/CoPaw saw active but mostly unmerged development in the last 24 hours: **9 issues updated** (all open), **12 PRs updated** (11 open, 1 closed), and **0 releases published**. The activity pattern suggests a heavy bug-fixing and community-contribution phase, with several PRs already targeting today’s newly reported issues. First-time contributor participation is notable, especially around ACP transport and provider-compat bugs. No release shipped, so users are still waiting for the accumulated fixes to land.

## Releases

No new releases in the last 24 hours.

## Project Progress

Only one PR was closed/merged today:

- **[#6598 — fix(skills): preserve plugin-sourced skill tags across reconcile cycles](https://github.com/agentscope-ai/QwenPaw/pull/6598)**
  - Fixes [#6537](https://github.com/agentscope-ai/QwenPaw/issues/6537): plugin-sourced skills lost tags on restart because reconciliation deleted manifest entries whose source directories were not on disk.

Other notable open PRs are ready for review or in active iteration, forming a clear “fix wave”:

- [#6629](https://github.com/agentscope-ai/QwenPaw/pull/6629) — fix auto-compression not triggering `summarize_when_compact`
- [#6623](https://github.com/agentscope-ai/QwenPaw/pull/6623) — fix ACP notification/prompt race losing final text
- [#6620](https://github.com/agentscope-ai/QwenPaw/pull/6620) — fix Gemini `ToolCallBlock` extra_content crash
- [#6631](https://github.com/agentscope-ai/QwenPaw/pull/6631) — align Aliyun coding plan models
- [#6622](https://github.com/agentscope-ai/QwenPaw/pull/6622) — add OrcaRouter as a built-in provider

## Community Hot Topics

The most-discussed issues today each have 2 comments:

- **[#6593 — Unified cleaning page for QwenPaw storage](https://github.com/agentscope-ai/QwenPaw/issues/6593)**  
  Users report that long-running agents accumulate memory, tool artifacts, backups, and session data with no good way to clean them. The request is for a global, agent-agnostic cleanup page with manual and optional automatic cleaning. This is a UX/polish concern but also a practical reliability issue for heavy users.

- **[#6480 — `nohup` / background shell commands never return to idle](https://github.com/agentscope-ai/QwenPaw/issues/6480)**  
  Running shell commands with `nohup` or trailing `&` leaves the agent stuck. This has been open since July 26 and affects automation workflows that rely on background processes.

- **[#6568 — Global hotkey floating quick input window](https://github.com/agentscope-ai/QwenPaw/issues/6568)**  
  Users want a Raycast/豆包-style lightweight input overlay, because opening the full 1280×800 desktop window for a quick question is too heavy.

Also generating discussion:

- **[#6621 — Missing multi-agent collaboration guidance](https://github.com/agentscope-ai/QwenPaw/issues/6621)**  
  A user reports 50+ rounds of multi-agent debugging before discovering agents only activate if explicitly referenced in `PROFILE.md`. The underlying need is better discoverability and onboarding for multi-agent workflows.

## Bugs & Stability

Ranked by severity:

1. **Critical crash: `ToolCallBlock` has no field `extra_content`**  
   [#6619](https://github.com/agentscope-ai/QwenPaw/issues/6619) — In QwenPaw 2.0.1 + agentscope 2.0.4.post1, streaming Gemini tool-call responses crashes on every affected request.  
   **Fix PR exists:** [#6620](https://github.com/agentscope-ai/QwenPaw/pull/6620)

2. **ACP delegate sometimes returns “completed without text output”**  
   [#6625](https://github.com/agentscope-ai/QwenPaw/issues/6625) — When `session/update` and `session/prompt` arrive in the same TCP segment, the final text is lost. The issue includes a root-cause analysis.  
   **Fix PR exists:** [#6623](https://github.com/agentscope-ai/QwenPaw/pull/6623)

3. **Scroll auto-compression does not trigger memory summarize flow**  
   [#6624](https://github.com/agentscope-ai/QwenPaw/issues/6624) — Manual `/compact` works, but automatic context eviction skips `summarize_when_compact`, causing memory loss.  
   **Fix PR exists:** [#6629](https://github.com/agentscope-ai/QwenPaw/pull/6629)

4. **CI “Real behavior proof” gate strips fenced Evidence blocks**  
   [#6626](https://github.com/agentscope-ai/QwenPaw/issues/6626) — PRs with only fenced code blocks in the `## Evidence` section are incorrectly rejected. No fix PR has been opened yet.

5. **Long-running `nohup` shell commands hang the agent**  
   [#6480](https://github.com/agentscope-ai/QwenPaw/issues/6480) — Downgraded from critical because it’s an edge-case workflow, but it remains unanswered and affects real automation users.

## Feature Requests & Roadmap Signals

The clearest roadmap signals from today’s issue/PR activity:

- **Global cleanup UI** — [#6593](https://github.com/agentscope-ai/QwenPaw/issues/6593) requests a dedicated page to clean generated files, expired memory, backups, and session history across all agents. A likely candidate for a future desktop/console release.
- **Global-hotkey quick input** — [#6568](https://github.com/agentscope-ai/QwenPaw/issues/6568) aligns QwenPaw desktop with modern launcher-style UX, similar to Raycast or 豆包.
- **Provider ecosystem expansion** — PRs like [#6622](https://github.com/agentscope-ai/QwenPaw/pull/6622) (OrcaRouter built-in), [#6631](https://github.com/agentscope-ai/QwenPaw/pull/6631) (Aliyun coding plan model alignment), and [#6302](https://github.com/agentscope-ai/QwenPaw/pull/6302) (unified provider discovery/metadata) show a strong push toward better provider configuration and routing.
- **Tracing / observability** — [#6627](https://github.com/agentscope-ai/QwenPaw/issues/6627) asks how to use LoongSuite Python traces with QwenPaw, indicating demand for deeper LLM observability.
- **Desktop workspace access** — [#6306](https://github.com/agentscope-ai/QwenPaw/pull/6306) adds a workspace shortcut to the sidebar, useful for users who need to open generated artifacts.

If merged, the current patch wave (memory fix, ACP fix, Gemini crash fix, provider alignment) would likely form the basis of the next 2.0.x patch release.

## User Feedback Summary

Real user pain points from the last 24 hours:

- **Long-term storage bloat is a major concern.** Users find QwenPaw “increasingly bloated” after heavy agent use and are afraid to manually delete files due to risk of breaking something.
- **Multi-agent features are not discoverable enough.** One user reports losing significant time because agents do not automatically activate unless explicitly referenced in `PROFILE.md`, despite reading the official docs.
- **Desktop interaction is too heavyweight for quick questions.** Users want a floating mini-input window rather than opening the full main UI.
- **Background shell execution is broken for common patterns.** `nohup ... &` causes the agent to hang, which blocks legitimate workflows.
- **Timezone and timestamp handling is confusing.** [#6618](https://github.com/agentscope-ai/QwenPaw/pull/6618) fixes a forced-UTC interpretation in the session list, indicating user-visible timezone bugs.

Overall, users are actively engaged and willing to file detailed issue reports, often with root-cause analysis. The dissatisfaction is mostly around stability and discoverability, not core agent capability.

## Backlog Watch

The following items have been open for a while and likely need maintainer attention:

- **[#5490 — feat(console): show tool-card images inline and add gallery navigation](https://github.com/agentscope-ai/QwenPaw/pull/5490)**  
  Open since **June 24**. Large UX improvement for media files in chat, still unreviewed/unmerged after more than a month.

- **[#6302 — feat: unify provider discovery, model metadata, routing, and agent controls](https://github.com/agentscope-ai/QwenPaw/pull/6302)**  
  Open since **July 21**. A broad architectural PR tied to issue #6167. Needs maintainer review because it touches provider onboarding and model management.

- **[#6306 — feat(desktop): add workspace shortcut to sidebar](https://github.com/agentscope-ai/QwenPaw/pull/6306)**  
  Open since **July 21**. Closes #6083. Small, user-facing desktop improvement stuck in review.

- **[#6480 — `nohup` shell command hangs](https://github.com/agentscope-ai/QwenPaw/issues/6480)**  
  Open since **July 26**. No fix PR has been linked, despite the reproduction path being clearly described.

These items, especially the two July PRs, appear to be delayed not because of controversy but because of limited maintainer bandwidth.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-08-02

## 1. Today's Overview

ZeroClaw had a busy but **non-completing** 24 hours: 14 issues and 50 PRs were updated, with **0 issues closed** and **0 PRs merged/closed**. Security and multi-agent isolation concerns dominate new bug reports, while many PRs are stuck in `needs-author-action` or `stale-candidate` states, suggesting a maintainer/review bottleneck. No release was published in this window. Overall project health is **active but at risk of queue bloat**: issue/PR volume is high, but merge throughput is currently zero.

---

## 2. Releases

No new releases were published in the last 24 hours.

The active roadmap signals are the release trackers:
- [v0.8.5 weekly non-breaking release tracker #9459](https://github.com/zeroclaw-labs/zeroclaw/issues/9459)
- [v0.9.0 auth/security/gateway/breaking-change tracker #7432](https://github.com/zeroclaw-labs/zeroclaw/issues/7432)

---

## 3. Project Progress

No PRs were merged or closed in this window. The following large or high-risk PRs were updated and remain open:

- [fix(channels/telegram): skip unauthorized handler for non-mentioned group messages with mention_only #9634](https://github.com/zeroclaw-labs/zeroclaw/pull/9634) — Adds Telegram `allow_groups` config and refines `mention_only` behavior.
- [feat(eval): append-only run-history receipts #9248](https://github.com/zeroclaw-labs/zeroclaw/pull/9248) — Adds eval run-history receipts for trend analysis.
- [feat(computer-use): add native macOS, Linux X11, and Windows drivers #9091](https://github.com/zeroclaw-labs/zeroclaw/pull/9091) — Large feature adding desktop-control tooling with fail-closed OS drivers.
- [feat(relay): secure transport and browser enrollment frontdoor #9080](https://github.com/zeroclaw-labs/zeroclaw/pull/9080) — Adds mutual TLS for remote WSS and a browser enrollment frontdoor.
- [fix(anthropic): support stored OAuth profiles #9420](https://github.com/zeroclaw-labs/zeroclaw/pull/9420) — Adds explicit OAuth profile resolution for Anthropic providers.
- [fix(browser): validate screenshot destination path against workspace policy #9362](https://github.com/zeroclaw-labs/zeroclaw/pull/9362) — Security fix for an arbitrary file write in the browser screenshot action.

No feature or fix reached `master` during the digest window.

---

## 4. Community Hot Topics

### Most active issues by comment count

- [Tracker: Maintainer decision queue for RFCs and design issues #8692](https://github.com/zeroclaw-labs/zeroclaw/issues/8692) — **7 comments**  
  Central decision queue for RFCs, design issues, and release-policy questions. This is the main coordination bottleneck for unblocking maintainer reviews.

- [RFC: Treat an empty WhatsApp Web `allowed_groups` as permit-none #9397](https://github.com/zeroclaw-labs/zeroclaw/issues/9397) — **5 comments**  
  High-priority security RFC, currently `p1` and `in-progress`. It proposes default-deny semantics for WhatsApp Web groups.

- [Tracker: v0.9.0 auth, security, gateway, and breaking-change queue #7432](https://github.com/zeroclaw-labs/zeroclaw/issues/7432) — **2 comments**  
  Public tracker for the next major breaking-change cycle.

- [Feature: Send stable session_id to OpenRouter for prompt-cache savings #9631](https://github.com/zeroclaw-labs/zeroclaw/issues/9631) — **2 comments**  
  User-visible cost concern: OpenRouter chats replay system prompts and tool schemas on every request.

PR comment counts were not exposed in the data. By size, risk, and labels, the most attention-heavy PR threads are:

- [feat(relay): secure transport and browser enrollment frontdoor #9080](https://github.com/zeroclaw-labs/zeroclaw/pull/9080)
- [feat(computer-use): add native macOS, Linux X11, and Windows drivers #9091](https://github.com/zeroclaw-labs/zeroclaw/pull/9091)
- [refactor(zerocode): consolidate Code pane, rails, and prompt drafts #8655](https://github.com/zeroclaw-labs/zeroclaw/pull/8655)
- [feat(eval): append-only run-history receipts #9248](https://github.com/zeroclaw-labs/zeroclaw/pull/9248)
- [feat(config): add schema struct & risk field #7821](https://github.com/zeroclaw-labs/zeroclaw/pull/7821)

Underlying needs: maintainers need to clear decision queues, security defaults need to become stricter, and users need lower LLM costs.

---

## 5. Bugs & Stability

### S0 — Critical security / data isolation

- [Session/channel read+write tools lack per-agent ownership scoping #9646](https://github.com/zeroclaw-labs/zeroclaw/issues/9646)  
  Model-supplied `session_id` / `channel_id` arguments are not checked against per-agent ownership. Any agent can reach another agent’s sessions and channels. **No visible fix PR yet.**

- [Knowledge graph has no per-agent attribution #9647](https://github.com/zeroclaw-labs/zeroclaw/issues/9647)  
  The `knowledge` tool exposes a single shared knowledge graph to all agents, allowing cross-agent reads/mutations. **No visible fix PR yet.**

### P1 — High severity

- [An approval that times out is recorded as an explicit operator denial #9642](https://github.com/zeroclaw-labs/zeroclaw/issues/9642)  
  Falsifies the audit trail by logging a human denial that never happened. **No visible fix PR yet.**

- [Windows null-device redirect `2>nul` rejected by shell security policy #9633](https://github.com/zeroclaw-labs/zeroclaw/issues/9633)  
  Shell policy accepts `/dev/null` but rejects Windows `nul` / `NUL`, breaking legitimate Windows usage. **No visible fix PR yet.**

- [WhatsApp Web policy doc comments cite `allowed_numbers`, a V2 key with no V3 field #9640](https://github.com/zeroclaw-labs/zeroclaw/issues/9640)  
  Documentation points operators to a nonexistent config key. **No visible fix PR yet.**

### P2 — Low severity

- [Blog missing RSS/Atom feed #9628](https://github.com/zeroclaw-labs/zeroclaw/issues/9628)  
  Usability/docs gap for project followers.

No bug-fix PRs were merged in this window. Several security-fix PRs remain open from prior days, including [#9362](https://github.com/zeroclaw-labs/zeroclaw/pull/9362), [#8918](https://github.com/zeroclaw-labs/zeroclaw/pull/8918), and [#8713](https://github.com/zeroclaw-labs/zeroclaw/pull/8713).

---

## 6. Feature Requests & Roadmap Signals

- [Feature: Send stable session_id to OpenRouter for prompt-cache savings #9631](https://github.com/zeroclaw-labs/zeroclaw/issues/9631)  
  Strong user value: reduces LLM costs for multi-turn agent chats. Currently `needs-maintainer-review`; likely candidate for a near-term minor release if approved.

- [Feature: Select a default agent for standalone ACP with --agent #9632](https://github.com/zeroclaw-labs/zeroclaw/issues/9632)  
  `p2`, `in-progress`, `accepted`. Small CLI change likely to land in v0.8.5.

- [RFC: retire the Lucid memory connector at v0.9.0 #9644](https://github.com/zeroclaw-labs/zeroclaw/issues/9644)  
  `p2`, `needs-maintainer-review`, `risk:high`. A breaking-change candidate for v0.9.0.

- [RFC: Treat an empty WhatsApp Web `allowed_groups` as permit-none #9397](https://github.com/zeroclaw-labs/zeroclaw/issues/9397)  
  `p1`, `in-progress`. Security default-change likely to ship soon, possibly as a patch if treated as a bug fix.

- [test(web): dashboard reconnect harness and detached-turn assertions #9641](https://github.com/zeroclaw-labs/zeroclaw/issues/9641)  
  `p1`, `in-progress`, `accepted`; test infrastructure for a prior dashboard change.

- [Blog missing RSS/Atom feed #9628](https://github.com/zeroclaw-labs/zeroclaw/issues/9628)  
  Small accepted docs/feature request; possible quick win.

**Roadmap prediction:** v0.8.5 is likely to contain small accepted items like #9632, #9641, and maybe #9628. v0.9.0 is the likely home for #9644 and other breaking changes tracked in #7432.

---

## 7. User Feedback Summary

- **Multi-agent isolation is the loudest user pain point.** Two S0 reports (#9646, #9647) show users expecting strict per-agent ownership over sessions, channels, and knowledge graphs.
- **LLM cost is a real operational concern.** OpenRouter users want prompt caching via stable `session_id` (#9631).
- **Audit-log trust matters.** Users report that approval timeouts are being misrepresented as explicit operator denials (#9642).
- **Windows parity is missing.** Shell policy treats Unix `/dev/null` as safe but rejects Windows `nul` (#9633).
- **Documentation quality directly affects configuration.** WhatsApp operators can follow docs and configure a nonexistent key (#9640).
- **Contributor friction is visible.** Many PRs are `needs-author-action` or `stale-candidate`, and one maintainer repaired a branch after the author-action deadline in [#8918](https://github.com/zeroclaw-labs/zeroclaw/pull/8918).
- **Positive signal:** the blog itself was described as "quite nice", but missing RSS (#9628).

---

## 8. Backlog Watch

These items are long-running, high-risk, or stuck and need maintainer attention:

- [Tracker: Maintainer decision queue for RFCs and design issues #8692](https://github.com/zeroclaw-labs/zeroclaw/issues/8692)  
  Opened 2026-07-04, updated 2026-08-01, 7 comments. This is the queue that should unblock many of the RFCs and PRs above.

- [feat(config): add schema struct & risk field #7821](https://github.com/zeroclaw-labs/zeroclaw/pull/7821)  
  Opened 2026-06-17, `stale-candidate`, `needs-author-action`, `risk:high`. Large sandbox-policy config change.

- [fix(tools/mcp): centralize deferred-MCP access policy #8496](https://github.com/zeroclaw-labs/zeroclaw/pull/8496)  
  Opened 2026-06-29, `stale-candidate`, `needs-author-action`, `risk:high`. Security-relevant MCP access-policy fix.

- [fix(channels): add env-var fallback for OpenAI STT credentials #8576](https://github.com/zeroclaw-labs/zeroclaw/pull/8576)  
  Opened 2026-07-01, `stale-candidate`, `needs-author-action`, `p3`.

- [refactor(zerocode): consolidate Code pane, rails, and prompt drafts #8655](https://github.com/zeroclaw-labs/zeroclaw/pull/8655)  
  Opened 2026-07-03, `stale-candidate`, `needs-author-action`, `risk:high`, size XL.

- [fix(security): redact Slack tokens in the leak detector #8918](https://github.com/zeroclaw-labs/zeroclaw/pull/8918)  
  Opened 2026-07-09, `needs-maintainer-review`, `risk:high`. Maintainer already repaired the branch; now needs final review.

- [feat(relay): secure transport and browser enrollment frontdoor #9080](https://github.com/zeroclaw-labs/zeroclaw/pull/9080)  
  Opened 2026-07-15, `needs-author-action`, `risk:high`, size XL. Major security/transport feature.

- [feat(computer-use): add native macOS, Linux X11, and Windows drivers #9091](https://github.com/zeroclaw-labs/zeroclaw/pull/9091)  
  Opened 2026-07-15, `needs-author-action`, `risk:high`, size XL. Major desktop automation feature.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*