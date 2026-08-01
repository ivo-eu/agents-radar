# OpenClaw Ecosystem Digest 2026-08-01

> Issues: 214 | PRs: 500 | Projects covered: 13 | Generated: 2026-08-01 00:12 UTC

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

# OpenClaw Project Digest — 2026-08-01

## 1. Today's Overview

OpenClaw remains highly active with 214 issues updated in the last 24 hours (193 open, 21 closed) and 500 PRs updated (372 open, 128 merged/closed). No new release was published today, so fixes and features are accumulating toward a future cut. The dominant themes are reliability around session state and message delivery, performance work on large history stores, and a large localization push. Many high-severity P1 bugs still carry `no-new-fix-pr` / `needs-maintainer-review` labels, indicating active triage but slower resolution on core architectural issues.

## 2. Releases

None today. No release notes, breaking changes, or migration notes to report.

## 3. Project Progress

From the top-30 PR snapshot, the following PRs were closed/merged in the last 24 hours:

- [#113192](https://github.com/openclaw/openclaw/pull/113192) — fix(auto-reply): make session reset granular for CLI sessions
- [#107909](https://github.com/openclaw/openclaw/pull/107909) — fix(ui): show live quick settings channel status
- [#116737](https://github.com/openclaw/openclaw/pull/116737) — fix(ollama): honor stop sequences and asynchronous payload hooks
- [#116742](https://github.com/openclaw/openclaw/pull/116742) — fix(agents): compact session status change text
- [#117090](https://github.com/openclaw/openclaw/pull/117090) — fix(ui): prevent stale sidebar sessions after switching filters

Several issues were also closed, including:

- [#116868](https://github.com/openclaw/openclaw/issues/116868) — SQLite-backed sessions can fall back to frozen legacy JSONL and resurrect completed tasks
- [#116973](https://github.com/openclaw/openclaw/issues/116973) — shipped docs reference retired `gateway.reload` tuning paths
- [#113059](https://github.com/openclaw/openclaw/issues/113059) — iPad Magic Keyboard Return not inserting newline in chat composer
- [#115311](https://github.com/openclaw/openclaw/issues/115311) — bounded Code Mode repair for smaller models

Notable open PRs advancing features/architecture:

- [#117034](https://github.com/openclaw/openclaw/pull/117034) — feat(audit): add execution identity inspection
- [#117040](https://github.com/openclaw/openclaw/pull/117040) — improve: make session lists faster on large stores
- [#115138](https://github.com/openclaw/openclaw/pull/115138) — perf(sqlite): memory-map reads on local-filesystem databases
- [#115761](https://github.com/openclaw/openclaw/pull/115761) — fix(cron): preserve watched commands across gateway hot reload
- [#116900](https://github.com/openclaw/openclaw/pull/116900) — fix(feishu): deliver `ask_user` question cards
- [#117089](https://github.com/openclaw/openclaw/pull/117089) — fix(outbound): preserve fenced tool call examples
- [#117119](https://github.com/openclaw/openclaw/pull/117119) — fix: prevent lost subagent replies and incorrect running session states
- [#117118](https://github.com/openclaw/openclaw/pull/117118) — perf: count large histories before Gateway prewarm
- [#117097](https://github.com/openclaw/openclaw/pull/117097) — refactor(channels): share Matrix and Slack progress draft composition

## 4. Community Hot Topics

The most active issues by comment count cluster around delivery reliability, data integrity, and token efficiency:

- [#86519](https://github.com/openclaw/openclaw/issues/86519) — [P1] Agent repeats identical replies 2–10× on Telegram after 5.20 update — 14 comments, 1 👍  
  Users are still experiencing duplicate-message regressions on a widely used channel.

- [#113306](https://github.com/openclaw/openclaw/issues/113306) — [P1] SQLite snapshot restore lacks end-to-end crash and identity guarantees — 13 comments  
  Discussion focuses on durability guarantees and cleanup edge cases.

- [#67419](https://github.com/openclaw/openclaw/issues/67419) — [P2] Session context bloat: bootstrap files re-injected every turn — 11 comments, 2 👍  
  Strong support from users concerned about 20–30% token waste on every session.

- [#115908](https://github.com/openclaw/openclaw/issues/115908) — [P1] Session transcript projection livelock under sustained writes — 11 comments  
  A serious availability issue: the Node main thread can stall for tens of seconds.

- [#114137](https://github.com/openclaw/openclaw/issues/114137) — [P1] Visible channel turns dispatch with no queued reply payloads — 11 comments  
  Final text is persisted but never delivered to the user.

- [#87109](https://github.com/openclaw/openclaw/issues/87109) — [P1] Gateway heap grows to 1073MB+ at idle; cron jobs fail silently — 10 comments, 1 👍  
  Long-running memory growth is disrupting cron automation.

- [#113251](https://github.com/openclaw/openclaw/issues/113251) — [Feature] Image viewing in webchat file viewer — 10 comments  
  Pure UX demand with no controversy.

- [#10687](https://github.com/openclaw/openclaw/issues/10687) — [P2] Fully dynamic model discovery (OpenRouter + beyond) — 9 comments, 3 👍  
  Most-reacted issue in the snapshot; users want fast-moving catalogs supported.

Underlying need: users are asking for stable message delivery, lower token/cost overhead, better observability, and more flexible provider/model support.

## 5. Bugs & Stability

Ranked by severity from the data available:

### P0
- [#112395](https://github.com/openclaw/openclaw/issues/112395) — Startup migration preflight blocks gateway after upgrade from 6.11 to 7.1; migration tables and leases are empty. Labeled `crash-loop` and `ux-release-blocker`; a linked PR is open.

### P1 — Session state / message loss / delivery
- [#86519](https://github.com/openclaw/openclaw/issues/86519) — Telegram duplicate replies after 5.20 update; still open.
- [#114137](https://github.com/openclaw/openclaw/issues/114137) — Final reply persisted but never delivered on Signal/visible channel turns.
- [#115908](https://github.com/openclaw/openclaw/issues/115908) — Transcript projection livelock blocks all channel transports.
- [#114211](https://github.com/openclaw/openclaw/issues/114211) — Matrix room agents loop on no-reply output and stale session replay.
- [#115476](https://github.com/openclaw/openclaw/issues/115476) — Context refresh after compaction replays old Telegram `message_id`; missing dedup.
- [#114653](https://github.com/openclaw/openclaw/issues/114653) — `sessions_send`/`sessions_history` transient visibility failures indistinguishable from policy denials.
- [#86963](https://github.com/openclaw/openclaw/issues/86963) — Orphaned/oversized native Codex thread wedges a session permanently.
- [#49889](https://github.com/openclaw/openclaw/issues/49889) — Telegram partial-stream delivery-mode observability gap.

### P1 — Data durability / memory / automation
- [#113306](https://github.com/openclaw/openclaw/issues/113306) — SQLite snapshot restore lacks crash and identity guarantees.
- [#87109](https://github.com/openclaw/openclaw/issues/87109) — Gateway heap grows to 1073MB+ at idle; cron jobs silently fail.
- [#90098](https://github.com/openclaw/openclaw/issues/90098) — Large attachments overflow Control UI/gateway stack.
- [#85844](https://github.com/openclaw/openclaw/issues/85844) — Auto-update leaves running gateway with stale hashed bundle imports.
- [#91892](https://github.com/openclaw/openclaw/issues/91892) — Cron jobs stall during AI model calls; `stream_progress` never completes.
- [#89228](https://github.com/openclaw/openclaw/issues/89228) — `exec` intermittently unavailable in isolated cron sessions.
- [#114615](https://github.com/openclaw/openclaw/issues/114615) — Every CLI invocation pays ~6s of eager plugin-graph init.

### P1 — Security / config
- [#54416](https://github.com/openclaw/openclaw/issues/54416) — Agent automatically executes config commands from example code without confirmation.

### Regression and config bugs
- [#114654](https://github.com/openclaw/openclaw/issues/114654) — `agents.defaults.compaction.*` reload is classified as no-op; edits silently never apply.
- [#115152](https://github.com/openclaw/openclaw/issues/115152) — Regression from #95939: `bootstrapMaxChars`/`bootstrapTotalMaxChars` deleted on every restart.
- [#114192](https://github.com/openclaw/openclaw/issues/114192) — TUI conversation history disappears after compaction.

Fix PRs are in motion for some of these: [#117119](https://github.com/openclaw/openclaw/pull/117119) addresses lost subagent replies, [#117118](https://github.com/openclaw/openclaw/pull/117118) improves large-history prewarm, [#117040](https://github.com/openclaw/openclaw/pull/117040) speeds session lists, and [#115138](https://github.com/openclaw/openclaw/pull/115138) reduces SQLite event-loop blocking. However, many P1s still lack a fix PR.

## 6. Feature Requests & Roadmap Signals

High-signal feature requests in the last 24 hours:

- [#10687](https://github.com/openclaw/openclaw/issues/10687) — Fully dynamic model discovery for OpenRouter and fast-moving catalogs.
- [#87325](https://github.com/openclaw/openclaw/issues/87325) — Azure AI Foundry GPT Realtime Talk via gateway relay.
- [#81913](https://github.com/openclaw/openclaw/issues/81913) — Stable plugin SDK surface for installed skill workflows.
- [#88032](https://github.com/openclaw/openclaw/issues/88032) — Telegram quote/reply as a durable first-class inbound contract.
- [#50205](https://github.com/openclaw/openclaw/issues/50205) — Configurable Gemini API request labels for GCP billing.
- [#10944](https://github.com/openclaw/openclaw/issues/10944) — Telegram `parseMode` config.
- [#7476](https://github.com/openclaw/openclaw/issues/7476) — WhatsApp sticker send support.
- [#9764](https://github.com/openclaw/openclaw/issues/9764) — Google Chat user OAuth for reactions and media uploads.
- [#57404](https://github.com/openclaw/openclaw/issues/57404) — Expose per-run token usage on WebSocket lifecycle events.
- [#57307](https://github.com/openclaw/openclaw/issues/57307) — Memory importance scoring + time decay for long-term memory.
- [#115924](https://github.com/openclaw/openclaw/issues/115924) — “Idea Shower” parallel thought collector while agent is working.
- [#53654](https://github.com/openclaw/openclaw/issues/53654) — Discord `messageUpdate`/`messageDelete` support for edit/delete workflows.
- [#113251](https://github.com/openclaw/openclaw/issues/113251) — Image viewing in the webchat file viewer.

In-progress roadmap signals from PRs: the localization trilogy ([#111544](https://github.com/openclaw/openclaw/pull/111544), [#111545](https://github.com/openclaw/openclaw/pull/111545), [#111542](https://github.com/openclaw/openclaw/pull/111542)) is waiting on author but appears close; audit execution identity ([#117034](https://github.com/openclaw/openclaw/pull/117034)) is ready for maintainer look; and the Tencent provider externalization ([#116300](https://github.com/openclaw/openclaw/pull/116300)) has compatibility and security implications.

Prediction: the next release will likely prioritize stability fixes around SQLite, session-list performance, cron hot reload, and the Telegram/Matrix message-delivery regressions. Dynamic model discovery and localization are strong candidates for the following minor release.

## 7. User Feedback Summary

User sentiment is mixed: there is clearly heavy adoption and active contribution, but also frustration with update regressions and silent failures.

Key pain points:

- **Update regressions** keep appearing: #86519 (5.20 duplicate replies), #90786 (6.1 memory embedding provider regression), #112395 (6.11 → 7.1 upgrade broken), and #115152 (post-#95939 config deletion).
- **Token/cost concerns** are a recurring theme: #67419 estimates 20–30% token waste; #95610 and #95840 show OpenAI prompt caching and `contextPruning` are not working as intended.
- **Silent failure modes** damage trust: cron jobs fail without output (#87109, #91892), messages are persisted but never delivered (#114137), and permission-like errors are indistinguishable from transient failures (#114653).
- **Platform gaps** are common: WhatsApp stickers, Discord edit/delete, Google Chat OAuth, Telegram `parseMode`, and image preview in webchat all remain requested.
- **Positive signal**: many PRs are labeled `ready for maintainer look`, and the maintainer bot labels (`clawsweeper:needs-maintainer-review`, `issue-rating`) show structured triage is active.

## 8. Backlog Watch

These important issues have been open for a long time and still need maintainer/fix attention:

- [#67419](https://github.com/openclaw/openclaw/issues/67419) — [P2] Session context bloat; open since Apr 15, 11 comments, 2 👍, no new fix PR.
- [#10687](https://github.com/openclaw/openclaw/issues/10687) — [P2] Dynamic model discovery; open since Feb 6, 9 comments, 3 👍.
- [#86519](https://github.com/openclaw/openclaw/issues/86519) — [P1] Telegram duplicate replies; open since May 25, 14 comments.
- [#87109](https://github.com/openclaw/openclaw/issues/87109) — [P1] Gateway heap growth / silent cron failures; open since May 27.
- [#90098](https://github.com/openclaw/openclaw/issues/90098) — [P1] Large attachment stack overflow; open since Jun 4.
- [#89228](https://github.com/openclaw/openclaw/issues/89228) — [P1] `exec` unavailable in isolated cron sessions; open since Jun 1.
- [#90786](https://github.com/openclaw/openclaw/issues/90786) — [P1] `memory status --index` fails with unknown embedding provider; open since Jun 5.
- [#91892](https://github.com/openclaw/openclaw/issues/91892) — [P1] Cron jobs stall during model calls; open since Jun 10.
- [#53654](https://github.com/openclaw/openclaw/issues/53654) — [P2] Discord edit/delete support; open since Mar 24, 3 👍.
- [#54416](https://github.com/openclaw/openclaw/issues/54416) — [P1] Agent auto-executes config commands without confirmation; open since Mar 25.
- [#95610](https://github.com/openclaw/openclaw/issues/95610) / [#95840](https://github.com/openclaw/openclaw/issues/95840) — OpenAI prompt-cache and cache-TTL pruning issues; open since Jun 21–22.
- [#81156](https://github.com/openclaw/openclaw/issues/81156) — MiniMax usage percentage semantics inverted; open since May 12.
- [#47979](https://github.com/openclaw/openclaw/issues/47979) — Control UI Dashboard freezes on Chrome 146; open since Mar 16.

PRs waiting on author/maintainer response:

- [#111544](https://github.com/openclaw/openclaw/pull/111544), [#111545](https://github.com/openclaw/openclaw/pull/111545), [#111542](https://github.com/openclaw/openclaw/pull/111542) — Localization trilogy, waiting on author since Jul 19.
- [#111755](https://github.com/openclaw/openclaw/pull/111755) — Yuanbao plugin catalog sync to 2.17.0, waiting on author.
- [#111115](https://github.com/openclaw/openclaw/pull/111115) — memory-core QMD multi-collection search fix, needs proof.

Overall, OpenClaw is in a high-activity period with strong community contribution, but the backlog of P1 session-state and message-delivery bugs suggests the next release should be a stability-focused one.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report — Personal AI Assistant / Agent Ecosystem
**Date:** 2026-08-01 | **Scope:** 13 tracked projects

---

## 1. Ecosystem Overview

The open-source personal AI assistant landscape remains densely populated and pre-consolidation: 13 projects are actively competing with heavily overlapping feature sets (channel adapters, session stores, provider integrations), while no project shipped a release on 2026-08-01. The dominant engineering themes across all projects are session-state durability (JSONL→SQLite migrations), message-delivery reliability, token/cost efficiency, and security hardening. The field is splitting into three architectural camps: general-purpose gateway platforms (OpenClaw), lightweight personal assistants (NanoBot, NanoClaw, PicoClaw, NullClaw), and enterprise/multi-tenant or developer-workflow platforms (IronClaw, ZeroClaw, Hermes Agent, CoPaw). A notable dynamic is the emergence of OpenClaw-compatible derivatives (LobsterAI) and OpenClaw-adjacent branding (NanoClaw, PicoClaw, NullClaw, IronClaw), indicating OpenClaw's ecosystem gravity — even as its stability backlog creates openings for lighter competitors.

---

## 2. Activity Comparison

| Project | Issues (24h) | PRs (24h) | PRs Merged/Closed | Release Status | Health Score* |
|---|---|---|---|---|---|
| OpenClaw | 214 | 500 | 128 | None (accumulating) | 6/10 |
| IronClaw | 16 | 50 | 32 | None (release PR blocked ~1 mo) | 6/10 |
| ZeroClaw | 27 | 50 | 12 | None | 8/10 |
| Hermes Agent | 16 | 50 | 4 | None | 8/10 |
| CoPaw (QwenPaw) | 18 | 43 | 13 | None | 6/10 |
| NanoBot | 4 | 17 | 7 | None | 8/10 |
| LobsterAI | 4 | 12 | 11 | None | 7/10 |
| NanoClaw | 8 | 9 | 3 | None (pipeline restored) | 6/10 |
| Moltis | 2 | 7 | 1 | None | 8/10 |
| PicoClaw | 0 | 3 | 0 | None | 6/10 |
| NullClaw | 0 | 1 | 0 | None | 7/10 |
| TinyClaw | 0 | 0 | 0 | None | 3/10 |
| ZeptoClaw | 0 | 0 | 0 | None | 3/10 |

*Health Score = composite of fix velocity, open-bug severity (P0/P1), security posture, and maintainer responsiveness, derived from the 24h digest data.

**Key observations:**
- **OpenClaw** dominates raw volume by 4–10× but carries the largest P1 backlog (duplicate replies, livelocks, memory growth) with many items lacking fix PRs.
- **IronClaw** shows the highest merge velocity (32 PRs) but is shipping a P0 cross-user memory leak and has a month-old blocked release PR.
- **NanoBot, Moltis, and ZeroClaw** have the best health-to-activity ratio: security-critical bugs are closed with fixes in the same window, and architectural work is advancing.
- **TinyClaw and ZeptoClaw** are effectively dormant.

---

## 3. OpenClaw's Position

**Advantages vs. peers:**
- **Community scale:** 214 issues / 500 PRs per day vs. 4–50 for all peers; unmatched contributor base and triage infrastructure (bot-labeled reviews, issue ratings).
- **Channel breadth:** Telegram, Matrix, Slack, Feishu, Discord, WhatsApp, Signal, Google Chat — the widest surface of any project in the set.
- **Feature maturity:** audit identity inspection, cron with hot-reload, localization trilogy, plugin SDK signals, and a large PR pipeline advancing SQLite performance and large-history handling.
- **Ecosystem gravity:** spawned an explicit derivative (LobsterAI) and multiple name-alike projects (NanoClaw, PicoClaw, NullClaw, IronClaw), cementing it as the reference architecture.

**Technical approach differences:** OpenClaw uses a gateway + SQLite session store + channel-adapter architecture with a heavy plugin/audit surface. Peers are deliberately simpler: NanoBot converges on lightweight personal-assistant UX; NullClaw delegates to official vendor CLIs; PicoClaw is a minimal Go implementation. OpenClaw's complexity cost shows in its P1 cluster — session-state fallback to legacy JSONL, transcript-projection livelock, heap growth to 1GB+ — which smaller projects avoid by construction.

**Community size comparison:** OpenClaw's issue volume alone (214/day) exceeds the *total tracked activity* of all other projects combined. However, its size also creates triage lag: several P1 issues (Telegram duplicates, #86519; context bloat, #67419) have been open for months. This is the key vulnerability window for lightweight competitors.

---

## 4. Shared Technical Focus Areas

Requirements emerging independently across multiple projects:

1. **SQLite / session-store durability** — OpenClaw (SQLite fallback resurrects completed tasks), NanoBot (JSONL→SQLite migration, #5173), IronClaw (libSQL write-stall pathology), CoPaw (agent.json corruption on Windows), ZeroClaw (session-persistence contract ownership). *All five* are converging on SQLite as the runtime session backbone with crash/identity guarantees.

2. **Channel delivery reliability** — OpenClaw (Telegram duplicates, Signal replies persisted but never delivered), NanoBot (WeChat errcode -14 60-min pauses), NanoClaw (Telegram pairing silently broken), CoPaw (WeChat cron reports success but never delivers). Trust in "did the message actually arrive?" is the #1 user pain point across the ecosystem.

3. **Token / cost efficiency** — OpenClaw (20–30% token waste from bootstrap re-injection; prompt caching broken), LobsterAI (DeepSeek prefix-cache hit rate collapse 100%→57%), Hermes (subagents inherit all 21 parent toolsets), CoPaw (44M tokens burned on undelivered WeChat push). Byte-stable prompt construction and cache-aware projections are becoming core engineering requirements.

4. **Provider/model flexibility** — OpenClaw (dynamic model discovery), ZeroClaw (OpenAI-compatible endpoint for Open WebUI/LobeChat/Aider), NanoBot (DeepSeek Responses API), CoPaw (NVIDIA NIM), NullClaw (Grok CLI provider), IronClaw (hosted MCP server registration).

5. **Security hardening as adoption prerequisite** — IronClaw (P0 cross-user memory leak, TOCTOU FS fix), Moltis (zip path traversal, missing pairing verification, privileged-command exposure), ZeroClaw (Landlock sandbox breaking shell), NanoClaw (interactive-card spoofing, log secret redaction), OpenClaw (auto-exec of config commands), Hermes (session-ID randomization).

6. **Desktop / WebUI UX** — CoPaw (Windows data corruption, quick-input window, bundled Python), NanoBot (Windows MIME-type WebUI failure), OpenClaw (iPad keyboard, webchat image viewer), LobsterAI (stale-closed sidebar improvements).

7. **Deployment flexibility beyond Docker** — NanoClaw (K8s/Sealos, native runner, Docker-less), CoPaw (bundled Python runtime), IronClaw (systemd linger for headless).

---

## 5. Differentiation Analysis

| Project | Core Differentiator | Target User | Architecture |
|---|---|---|---|
| **OpenClaw** | Broadest channel + plugin ecosystem | Self-hosters, power users | Gateway + SQLite + adapters (Node) |
| **IronClaw** | Multi-tenant security & contract extraction (WS1) | Enterprises, SaaS | Rust workspace, tenant-isolated FS, hosted MCP |
| **ZeroClaw** | RFC-driven design; Wasm sandbox; ZeroCode | Developers, coding agents, robotics/IoT | Rust, WASM plugin runtime, OpenAI-compat gateway |
| **CoPaw (QwenPaw)** | Chinese ecosystem (Feishu/WeChat, AgentScope) | Chinese-speaking desktop users | Python/AgentScope, desktop app, ReMe/Auto-Dream memory |
| **Hermes Agent** | Desktop/TUI + developer ergonomics | Developers, desktop-first | Gateway profiles, Discord auth, LSP integration |
| **NanoBot** | Lightweight personal assistant, fast iteration | Individual users | Python, WebUI, WeChat/Weixin focus |
| **NanoClaw** | Container-runtime flexibility | Restricted/K8s environments | Docker-centric w/ Apple Container + native runner signals |
| **Moltis** | Security-gated modularity; Nostr/Buzz | Privacy-conscious teams | Rust, NIP-29 group chat, zvec memory |
| **LobsterAI** | DeepSeek cache-optimized OpenClaw derivative | DeepSeek users | OpenClaw-compatible + cache-stable prompt projection |
| **PicoClaw** | Minimal Go, niche channels | IRC/DeltaChat/Simplex users | Go, lightweight |
| **NullClaw** | CLI-provider delegation | CLI power users | Spawn-per-request to vendor CLIs (codex, gemini, grok) |
| **TinyClaw / ZeptoClaw** | — | — | Inactive |

---

## 6. Community Momentum & Maturity

**Tier 1 — High velocity, high complexity (rapid iteration with stability risk):**
OpenClaw, IronClaw, ZeroClaw, CoPaw, Hermes Agent. All five are accumulating features/fixes toward unreleased cuts; IronClaw and ZeroClaw show the most disciplined processes (stacked WS1 contract PRs; RFC-tracker + follow-up issue waves).

**Tier 2 — Steady, healthy (good fix velocity, manageable backlog):**
NanoBot (fastest bug-to-fix turnaround; major SQLite migration landed), Moltis (security fixes in review), NanoClaw, LobsterAI.

**Tier 3 — Stabilizing / low throughput:**
PicoClaw (stable, no regressions, but 3 PRs awaiting review), NullClaw (clean but review-velocity low), LobsterAI (maintenance mode, auto-closing UX work as stale).

**Tier 4 — Inactive:**
TinyClaw, ZeptoClaw (no activity in 24h window).

**Release-blocked:** IronClaw (release PR #5598 open ~1 month, API-breaking bumps) and NanoClaw (release pipeline was broken; fix merged #3163). OpenClaw's next release is widely expected to be stability-focused, per community signal.

---

## 7. Trend Signals

1. **SQLite is winning the session-store war.** NanoBot's JSONL→SQLite migration, OpenClaw's SQLite perf PRs, and ZeroClaw/IronClaw contract work all point to SQLite as the ecosystem-standard local session backbone. Developers should design around SQLite durability, snapshot-restore integrity, and WAL-mode performance from day one.

2. **Delivery-reliability guarantees are the new competitive battleground.** Duplicate replies, silently dropped messages, and cron pushes that "succeed" but never arrive were reported in 6 of 13 projects. Build idempotent delivery, explicit delivery receipts, and dedup keys into channel adapters.

3. **Cache-aware prompt construction is a hard requirement.** LobsterAI's DeepSeek hit-rate collapse (100%→57%) and OpenClaw's byte-instability issues show that prefix caching fails under naive prompt rewriting. Prompt projections must be byte-stable across turns; this directly affects user cost-per-session.

4. **OpenAI API compatibility is becoming table stakes for interop.** ZeroClaw's endpoint request (#8550) explicitly names Open WebUI, LobeChat, LangChain, Continue.dev, and Aider as blocked without it. New agents should expose an OpenAI-compatible surface to inherit the tooling ecosystem.

5. **Deployment flexibility beyond Docker is a rising adoption barrier.** NanoClaw's K8s/Docker-less thread (4 active issues) and CoPaw's bundled-Python request reflect real enterprise and desktop friction. Runtime abstraction (native/K8s/WASM) is emerging as a differentiator.

6. **Security hardening is a prerequisite, not a feature.** Cross-user memory leaks (IronClaw P0), zip path traversal (Moltis), sandbox breakage (ZeroClaw Landlock), and secret-in-log leaks (NanoClaw) show the bar rising. Projects that ship failsafe-by-default isolation will win team/enterprise evaluations.

7. **CLI-provider delegation and local tooling are gaining traction.** NullClaw's grok-cli PR extends a pattern (codex-cli, gemini-cli, claude-cli) that avoids API-integration maintenance; ZeroClaw's LSP RFC and pre-submission gate target local-model codegen quality. Expect "shell out to the official vendor CLI" and "language-server backstops" to spread.

8. **The Chinese ecosystem channel stack is under-served globally.** CoPaw (Feishu/WeChat, Chinese-language issue base) and NanoBot's WeChat/Weixin fixes show concentrated demand. Global projects treating WeChat/Feishu as second-class will lose that user base.

**Bottom line for decision-makers:** OpenClaw remains the ecosystem reference and safest hiring/contribution target, but its stability debt is real. NanoBot and ZeroClaw offer the best health-to-momentum ratios; IronClaw is the strongest enterprise-security bet once its P0 leak and release PR clear. New projects should differentiate on delivery guarantees, cache-stable prompts, OpenAI compatibility, and deployment flexibility rather than re-implementing channel adapters.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-08-01

## 1. Today's Overview

NanoBot saw steady, high-velocity maintenance on 2026-08-01: 4 issues were touched in the last 24 hours (2 open, 2 closed) and 17 pull requests were updated, with 7 merged or closed. No new releases were published. The day’s work centered on WeChat/Weixin channel reliability, Windows/Termux platform compatibility, WebUI fixes, and a major session-storage migration from JSONL to SQLite. Overall project health looks strong: critical bugs are being fixed quickly, and a broad set of feature PRs are in flight.

## 2. Releases

No new releases were published in this window.

## 3. Project Progress

Multiple meaningful fixes and improvements were merged or closed today:

- **#5173 — Migrate session storage from JSONL to SQLite** ([PR](https://github.com/HKUDS/nanobot/pull/5173)): Makes `sessions.db` the runtime session store, imports existing JSONL sessions transactionally on first startup, and keeps JSONL as rollback backups. This is a major architectural advancement.
- **#5196 — Fix Weixin recover refreshed state after session expiry** ([PR](https://github.com/HKUDS/nanobot/pull/5196)): Fixes the WeChat `errcode -14` 60-minute pause issue by reloading persisted state after pause expiry.
- **#4223 — Reload Weixin session state after pause expiry** ([PR](https://github.com/HKUDS/nanobot/pull/4223)): Related Weixin fix, closed alongside the newer #5196.
- **#5192 — Fix Slack channel thread openers scoped to their own session** ([PR](https://github.com/HKUDS/nanobot/pull/5192)): Prevents unrelated threads from sharing a channel-wide session.
- **#5193 — Fix WebUI preserve user scroll ownership near tail** ([PR](https://github.com/HKUDS/nanobot/pull/5193)): Improves chat scroll behavior near the bottom of long threads.
- **#5189 — Install timezone data on all platforms** ([PR](https://github.com/HKUDS/nanobot/pull/5189)): Resolves Termux and minimal-Linux timezone startup failures by using `tzdata` as fallback.
- **#5145 — Stabilize and speed up CI** ([PR](https://github.com/HKUDS/nanobot/pull/5145)): Replaces a timing-dependent test with a handshake and batches dependency installs.

## 4. Community Hot Topics

The most actively discussed item in the available data is:

- **#5195 — [bug] [weixin] Re-scan QR login overwrites new token with old one in `stop()`, causing immediate errcode -14** ([Issue](https://github.com/HKUDS/nanobot/issues/5195)) — 2 comments, closed. This captured real user pain around WeChat re-login and session invalidation. It directly led to fix PR #5196.

Also notable because they are open and unresolved:

- **#5198 — Not possible to change models in a specific session unless reconfiguring the entire instance** ([Issue](https://github.com/HKUDS/nanobot/issues/5198))
- **#5190 — Module script loading fails with MIME type "text/plain"** ([Issue](https://github.com/HKUDS/nanobot/issues/5190))

Underlying need: users want more flexible session-level model selection, and Windows users need the WebUI to load correctly without manual MIME workarounds.

## 5. Bugs & Stability

Ranked by severity and user impact:

1. **Frontend completely blocked on Windows** — #5190 ([Issue](https://github.com/HKUDS/nanobot/issues/5190)): Module scripts fail because Windows registry maps `.js` to `text/plain`. An open fix exists: #5191 ([PR](https://github.com/HKUDS/nanobot/pull/5191)).
2. **In-session model switching impossible** — #5198 ([Issue](https://github.com/HKUDS/nanobot/issues/5198)): Users cannot change models without reconfiguring the whole instance; `/model` with another model ID reportedly does not work. No fix PR yet.
3. **Weixin session expiry after QR re-login** — #5195 ([Issue](https://github.com/HKUDS/nanobot/issues/5195)): Closed. Fix landed in #5196 ([PR](https://github.com/HKUDS/nanobot/pull/5196)).
4. **Termux startup failure due to timezone validation** — #5187 ([Issue](https://github.com/HKUDS/nanobot/issues/5187)): Closed. Fix landed in #5189 ([PR](https://github.com/HKUDS/nanobot/pull/5189)).

Additional stability fixes are pending review:

- **#5201 — Tolerate malformed persisted session summary** ([PR](https://github.com/HKUDS/nanobot/pull/5201))
- **#5200 — Preserve exec wait targets across response truncation** ([PR](https://github.com/HKUDS/nanobot/pull/5200))

## 6. Feature Requests & Roadmap Signals

The following are strong signals for future NanoBot versions:

- **Session-level model switching** — #5198 ([Issue](https://github.com/HKUDS/nanobot/issues/5198)): Users expect a simple UI to switch models per session, similar to commercial cloud assistants.
- **Quick Chat and Temporary Chat** — #5184 ([PR](https://github.com/HKUDS/nanobot/pull/5184)): Adds persistent Quick Chat and opt-in Temporary Chat with in-memory history. Likely to land soon.
- **DeepSeek Responses API support** — #5197 ([PR](https://github.com/HKUDS/nanobot/pull/5197)): Routes `deepseek-v4-flash` through DeepSeek’s native Responses API while preserving other models on Chat Completions.
- **Session export/import/search/stats commands** — #1565 ([PR](https://github.com/HKUDS/nanobot/pull/1565)): A long-open CLI feature, currently blocked by merge conflicts.
- **Skill status CLI command** — #1319 ([PR](https://github.com/HKUDS/nanobot/pull/1319)): Would help diagnose why skills are unavailable; also conflict-blocked.

The next release will likely include SQLite-backed sessions, the WebUI quick-chat experience, DeepSeek Responses support, and platform compatibility fixes.

## 7. User Feedback Summary

Real user pain points this cycle:

- **WeChat/Weixin reliability remains fragile**: Users re-scan QR codes, but the old token can overwrite the new one, causing silent 60-minute pauses. This was reported in #5195 and partially addressed by #5196.
- **Termux users cannot start NanoBot without system timezone data**: #5187 was quickly fixed by bundling `tzdata`.
- **Windows users cannot load the WebUI at all** due to MIME type misdetection: #5190. A fix is proposed but not yet merged.
- **Slack users saw cross-thread context leakage**, which was fixed in #5192.
- **Model switching UX is a repeated expectation**: #5198 indicates that users want per-session model selection without global reconfiguration.

Overall, users are engaged and reporting concrete edge-case bugs, while maintainers are responding rapidly with targeted fixes and architectural improvements.

## 8. Backlog Watch

Several older PRs remain open and are labeled with `conflict`, requiring rebase or maintainer attention:

- **#1656 — fix(validation): handle None value in string schema validation** ([PR](https://github.com/HKUDS/nanobot/pull/1656)) — open since 2026-03-07.
- **#1565 — feat(session): add session export, import, search and stats commands** ([PR](https://github.com/HKUDS/nanobot/pull/1565)) — open since 2026-03-05; overlaps with the new SQLite session work and should be re-evaluated.
- **#1319 — feat: add skill status command** ([PR](https://github.com/HKUDS/nanobot/pull/1319)) — open since 2026-02-28; no longer fresh but still useful for diagnostics.

These should be rebased, reviewed, or explicitly closed to keep the backlog healthy. The Weixin-related #4223 ([PR](https://github.com/HKUDS/nanobot/pull/4223)) was also conflict-labeled but is now closed, presumably superseded by #5196.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-08-01

## 1. Today's Overview

Hermes Agent is in a high-activity phase: 16 issues were updated in the last 24 hours (15 open, 1 closed) and 50 PRs were updated (46 open, 4 merged/closed). No new release was published in this window, so today’s activity is entirely issue triage, PR review, and fix preparation. The tracker shows a concentrated cluster of P2 bugs around update/backup reliability, desktop session-state handling, and gateway profile isolation, with several corresponding fix PRs already open. Overall, the project looks healthy: contributors are actively submitting targeted fixes, and maintainers are moving quickly on the most visible stability problems.

## 2. Releases

No new releases in this window.

## 3. Project Progress

- 4 PRs were merged or closed in the last 24 hours, though their titles are not itemized in the top-20 snapshot.
- 1 issue was closed: [#71045](https://github.com/NousResearch/hermes-agent/issues/71045) — feature request withdrawn by author.

Notable open PRs that advanced or entered review today:

- [#75735](https://github.com/NousResearch/hermes-agent/pull/75735) — Refactors DiscordAdapter’s 7k-line authorization cluster into a dedicated mixin, closing [#75734](https://github.com/NousResearch/hermes-agent/issues/75734).
- [#75720](https://github.com/NousResearch/hermes-agent/pull/75720) — Adds `laravel-lsp` for `.blade.php` files, closing [#75718](https://github.com/NousResearch/hermes-agent/issues/75718).
- [#75728](https://github.com/NousResearch/hermes-agent/pull/75728) — Fixes full pre-update backup aborting on non-SQLite `.db` files, addressing [#75724](https://github.com/NousResearch/hermes-agent/issues/75724).
- [#75730](https://github.com/NousResearch/hermes-agent/pull/75730) — Makes TUI/desktop gateway honor administrator-managed `config.set` keys, matching CLI behavior.
- [#75733](https://github.com/NousResearch/hermes-agent/pull/75733) — Generalizes desktop optimistic user-turn deduplication beyond images, fixing duplicate `@file:` attachments.
- [#75729](https://github.com/NousResearch/hermes-agent/pull/75729) — Scopes routed `/memory` and `/skills` writes to the correct profile home, fixing a profile-isolation bug.

## 4. Community Hot Topics

Most-commented issues this window:

- [#75598](https://github.com/NousResearch/hermes-agent/issues/75598) — [Bug]: issue with updates (4 comments). Users report update instability and conflicting gateway instances across profiles. This is the clearest signal of upgrade-path pain.
- [#66084](https://github.com/NousResearch/hermes-agent/issues/66084) — `_tui_need_npm_install()` compares against entire monorepo lockfile (3 comments). A workspace-scoping bug that likely triggers unnecessary TUI re-installs on every dashboard start.
- [#43800](https://github.com/NousResearch/hermes-agent/issues/43800) — Honcho plugin ignores `endpoint.baseUrl` (3 comments). Self-hosted users are being silently routed to Honcho cloud.
- [#75725](https://github.com/NousResearch/hermes-agent/issues/75725) — MiniMax-M3 interleaved thinking stops after first tool-call turn (2 comments).

Underlying themes: users want reliable upgrades, explicit self-hosted configuration, and consistent multi-profile/session isolation. The most active threads are not feature debates but real stability/config bugs.

## 5. Bugs & Stability

Ranked by severity:

1. **High — Update instability and backup failure**
   - [#75598](https://github.com/NousResearch/hermes-agent/issues/75598): updates cause program instability; multiple gateways conflict across profiles.
   - [#75724](https://github.com/NousResearch/hermes-agent/issues/75724): full pre-update backup aborts on Windows when `HERMES_HOME` contains a non-SQLite `.db` file. Fix PR exists: [#75728](https://github.com/NousResearch/hermes-agent/pull/75728).

2. **High — Desktop session and model defaults**
   - [#75727](https://github.com/NousResearch/hermes-agent/issues/75727): new desktop sessions use a localStorage-sticky model instead of the profile’s configured default.

3. **Medium — Profile/session isolation**
   - [#75708](https://github.com/NousResearch/hermes-agent/issues/75708): mem0 plugin ignores `gateway_session_key`, storing memories under `hermes-user` on API-server paths.
   - Related gateway fix PR: [#75729](https://github.com/NousResearch/hermes-agent/pull/75729) scopes routed memory/skill writes to the correct profile.

4. **Medium — Provider/gateway failures**
   - [#75725](https://github.com/NousResearch/hermes-agent/issues/75725): MiniMax-M3 stops thinking after the first tool call via `/anthropic` endpoint.
   - [#75693](https://github.com/NousResearch/hermes-agent/issues/75693): provider unreachable and fallback switching caused Telegram gateway errors.

5. **Needs repro — System-level breakage reports**
   - [#75694](https://github.com/NousResearch/hermes-agent/issues/75694): user reports Hermes made their computer unusable after `chown` during SFTP setup.
   - [#75695](https://github.com/NousResearch/hermes-agent/issues/75695): dashboard stopped working with "Frontend not built" error after the same incident.
   These are labeled `needs-repro` but suggest a failsafe gap in setup/ownership operations.

6. **Low/UI — Desktop rendering**
   - [#75731](https://github.com/NousResearch/hermes-agent/issues/75731): user message with `@file:` attachment rendered twice after history hydration; fix PR [#75733](https://github.com/NousResearch/hermes-agent/pull/75733).
   - [#75710](https://github.com/NousResearch/hermes-agent/issues/75710): sidebar session titles truncate to one line, especially problematic for RTL/Persian titles.

## 6. Feature Requests & Roadmap Signals

- **Laravel/Blade LSP support** — [#75718](https://github.com/NousResearch/hermes-agent/issues/75718), implemented in [#75720](https://github.com/NousResearch/hermes-agent/pull/75720). Likely next-version candidate.
- **Per-subagent toolset restriction** — [#75737](https://github.com/NousResearch/hermes-agent/issues/75737): `delegate_task` currently inherits all parent toolsets, bloating subagent prompts. High usability value for power users.
- **Discord authorization refactor** — [#75734](https://github.com/NousResearch/hermes-agent/issues/75734), implemented in [#75735](https://github.com/NousResearch/hermes-agent/pull/75735). Architecture/health improvement, not user-facing.
- **MCP config whitespace warnings** — [#75736](https://github.com/NousResearch/hermes-agent/pull/75736): warns on hidden whitespace in MCP config values, preventing mysterious 401/connect failures.
- **Discord voice barge-in** — [#75325](https://github.com/NousResearch/hermes-agent/pull/75325): opt-in conservative barge-in while bot TTS is playing.
- **Skills YAML date coercion** — [#56048](https://github.com/NousResearch/hermes-agent/pull/56048): coerce YAML date/time frontmatter to ISO strings, fixing `skill_view`.

Most of these are well-scoped incremental features; the per-subagent toolset restriction and Laravel LSP are the most visibly user-requested.

## 7. User Feedback Summary

- **Upgrade anxiety is real**: users who previously had smooth updates now report post-update instability and profile/gateway conflicts ([#75598](https://github.com/NousResearch/hermes-agent/issues/75598)).
- **Self-hosted configuration trust**: Honcho plugin users expect `baseUrl` in `~/.honcho/config.json` to be respected; instead they are silently routed to production cloud ([#43800](https://github.com/NousResearch/hermes-agent/issues/43800)).
- **Desktop defaults are confusing**: localStorage-sticky model selection overriding profile defaults causes session-to-session inconsistency ([#75727](https://github.com/NousResearch/hermes-agent/issues/75727)).
- **Power users want smaller, focused subagents**: loading 21 toolsets for a simple web-research subtask is seen as wasteful and prompt-bloating ([#75737](https://github.com/NousResearch/hermes-agent/issues/75737)).
- **Positive signal**: contributors are filing well-structured bugs with root-cause analysis, and maintainers are providing fix PRs in the same window (e.g., [#75724](https://github.com/NousResearch/hermes-agent/issues/75724) → [#75728](https://github.com/NousResearch/hermes-agent/pull/75728)).

## 8. Backlog Watch

Long-running or under-reviewed items that need maintainer attention:

- [#66084](https://github.com/NousResearch/hermes-agent/issues/66084) — TUI `_tui_need_npm_install()` lockfile comparison bug (open since 2026-07-17, updated 2026-08-01, 3 comments).
- [#43800](https://github.com/NousResearch/hermes-agent/issues/43800) — Honcho plugin wrong config key for self-hosted `baseUrl` (open since 2026-06-10, 3 comments).
- [#39643](https://github.com/NousResearch/hermes-agent/pull/39643) — Randomize fallback API session IDs for `/v1/chat/completions` (open since 2026-06-05, security-relevant).
- [#56048](https://github.com/NousResearch/hermes-agent/pull/56048) — Coerce YAML date/time frontmatter to ISO strings (open since 2026-07-01).
- [#56833](https://github.com/NousResearch/hermes-agent/pull/56833) — Soften MCP circuit-breaker error messages to avoid model over-adherence (open since 2026-07-02).

These are not all urgent by label, but several touch security/session-state boundaries and would benefit from maintainer review or merge decisions soon.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest — 2026-08-01

## Today's Overview

PicoClaw had a quiet day in the last 24 hours: no new releases, no PRs merged or closed, and no newly filed issues. Two open issues and three open PRs received updates, all from the prior day. The most notable activity is on an IRC long-message feature request and a CPU usage bug in the chat input. With three PRs still open and no merge activity, the project appears stable but maintainer throughput is currently low.

## Releases

No new releases were published in the data window.

## Project Progress

No pull requests were merged or closed today. Three PRs were updated and remain open:

- [PR #3222](https://github.com/sipeed/picoclaw/pull/3222) — `refactor(deltachat): cleanup implementation, documentation -200LOC` — removes legacy features, hardcoded relay list, and password-based email config; renames `invite_link` to `join_invite_link` and adds `show_invite_link`.
- [PR #3193](https://github.com/sipeed/picoclaw/pull/3193) — `Added simplex channel type` — adds a new messaging channel backend for Simplex.
- [PR #3200](https://github.com/sipeed/picoclaw/pull/3200) — `feat(models): add configurable default fallback chain` — adds a configurable fallback model chain in the web UI, persisted through the backend API.

These features are still awaiting review/merge.

## Community Hot Topics

The most active item in the last 24 hours is the feature request [Issue #3287](https://github.com/sipeed/picoclaw/issues/3287) — “Better support long messages in IRC,” with 2 comments. It asks PicoClaw to treat IRCv3 long messages as a single cohesive message despite IRC’s 512-byte limit and client-side splitting. The underlying need is better real-world IRC interoperability for users who send longer messages.

Also active is [Issue #3292](https://github.com/sipeed/picoclaw/issues/3292) — a bug report about high CPU usage when focusing the input box in the web chat interface. This appears to be a UI performance issue on Firefox/Linux.

No PRs had recorded comments or reactions in the provided data.

## Bugs & Stability

One bug report is currently active:

- [Issue #3292](https://github.com/sipeed/picoclaw/issues/3292) — `[BUG] CPU usage too high when focus on input box in chat interface` — reported on PicoClaw 0.3.1, Go 1.26, using deepseek-v4-flash through Firefox on Debian Linux x64. Marked as `stale` and updated on 2026-07-31.  
  **Severity:** Moderate — localized UI performance issue, no crash or data loss. No linked fix PR currently exists.

No new bugs, crashes, or regressions were filed in the last 24 hours.

## Feature Requests & Roadmap Signals

The open feature-related signals point toward improved messaging interoperability and model configuration:

- [Issue #3287](https://github.com/sipeed/picoclaw/issues/3287) — Better handling of long IRCv3 messages as a single message.
- [PR #3193](https://github.com/sipeed/picoclaw/pull/3193) — New Simplex channel type.
- [PR #3200](https://github.com/sipeed/picoclaw/pull/3200) — Configurable model fallback chain in the web UI.

If PR #3200 is merged, it is likely to be included in the next minor version, since it directly improves model fallback configuration for end users. Long-message IRC support may follow as a community-requested usability fix.

## User Feedback Summary

User-reported pain points from the active issues:

- **IRC fragmentation:** Messages over 512 bytes are split by IRC clients into multiple newline-separated chunks; users want PicoClaw to handle them as one logical message.
- **Web UI performance:** Focusing the chat input box causes high CPU usage, which affects daily chat interactions.

The open PRs also show community interest in expanding PicoClaw’s supported channels and improving model fallback workflows. No direct satisfaction/dissatisfaction ratings were available in this data window.

## Backlog Watch

Several important PRs have been open for weeks without recorded review comments:

- [PR #3193](https://github.com/sipeed/picoclaw/pull/3193) — Simplex channel support, open since 2026-06-27.
- [PR #3200](https://github.com/sipeed/picoclaw/pull/3200) — Default model fallback chain, open since 2026-07-01.
- [PR #3222](https://github.com/sipeed/picoclaw/pull/3222) — DeltaChat cleanup/refactor, open since 2026-07-03.

These are substantial, user-facing changes and would benefit from maintainer review or explicit feedback. The stale-labelled [Issue #3292](https://github.com/sipeed/picoclaw/issues/3292) also needs triage to avoid being closed without a fix.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-08-01

## 1. Today's Overview

NanoClaw saw sustained activity over the last 24 hours: **8 issues updated** (all still open) and **9 PRs updated** (6 open, 3 closed/merged). There were **no new releases**, and the latest release list is empty. The project is currently focused on container-runtime flexibility, new channel integrations (iMessage, Dial, Telegram), and security hardening. A new high-priority Telegram pairing bug was reported, while several long-running feature requests around Docker-less and Kubernetes execution remain open.

## 2. Releases

**No new releases in the last 24 hours.** The latest releases section is empty. A recently closed release-path fix ([#3163](https://github.com/nanocoai/nanoclaw/pull/3163)) suggests a release pipeline restoration may be in progress, but no versioned artifacts were published.

## 3. Project Progress

Three PRs moved to closed/merged state in the last 24 hours:

- **[#3163](https://github.com/nanocoai/nanoclaw/pull/3163) — fix(release): restore the v2.1.54 release path**  
  Closed/merged. Addresses a regression in the release pipeline, likely unblocking the next NanoClaw release.

- **[#3076](https://github.com/nanocoai/nanoclaw/pull/3076) — feat(imessage): unified local+hosted adapter targeting spectrum-ts v11**  
  Closed/merged. Advances iMessage support by consolidating local and hosted adapter paths.

- **[#1678](https://github.com/nanocoai/nanoclaw/pull/1678) — docs(skills): update voice transcription skills for Telegram + Linux**  
  Closed/merged. Improves documentation for voice transcription, removing the WhatsApp-only restriction and updating Whisper usage.

Also active in review: **#3164** (hosted iMessage via Photon), **#3041** (Dial SMS/voice channel), **#2809** (Apple Container runtime), **#2651** (security hardening for interactive cards), and **#3161** (log secret redaction).

## 4. Community Hot Topics

Most active issues by comments/reactions:

- **[#1184](https://github.com/nanocoai/nanoclaw/issues/1184) — Challenges deploying NanoClaw in restricted K8s environments (Sealos)**  
  *Comments: 3, 👍: 1*  
  The community is clearly interested in running NanoClaw in production on Kubernetes, but restricted environments make the current Docker-centric model difficult. The author praises NanoClaw as a “lightweight, secure alternative to the more bloated agent frameworks.”

- **[#1732](https://github.com/nanocoai/nanoclaw/issues/1732) — Native runner mode: bypass Docker for host-tool access**  
  *Comments: 3*  
  Frequent request from users who need tmux, headed browsers, and macOS APIs. Docker isolation is valued but is seen as a blocker for host-integrated agent use cases.

- **[#1225](https://github.com/nanocoai/nanoclaw/issues/1225) — Run it without Docker**  
  *Comments: 2*  
  A direct request for Docker-free operation on Windows and Linux.

- **[#2354](https://github.com/nanocoai/nanoclaw/issues/2354) — Kubernetes container runtime for agent spawning**  
  *Comments: 1, 👍: 1*  
  User want to spawn per-session agent pods on a user-provided cluster instead of local Docker.

Underlying need: users increasingly want NanoClaw to be deployable in managed/restricted environments and in host-integrated modes, not just Docker-based local containers.

## 5. Bugs & Stability

Ranked by severity:

- **[#3162](https://github.com/nanocoai/nanoclaw/issues/3162) — [High] Telegram pairing silently broken if boot-time getMe fails**  
  *Reported 2026-07-31, verified on `channels` branch.* A single failed HTTP call at boot can permanently break Telegram pairing, with no user-facing error. No dedicated fix PR exists yet.

- **[#2923](https://github.com/nanocoai/nanoclaw/issues/2923) — [security] `ask_user_question` card can be defaced by a forged click before origin authz**  
  A display/integrity spoof issue: a forged click can overwrite card text even when the response is rejected. Related fix PR **[#2651](https://github.com/nanocoai/nanoclaw/pull/2651)** is still open.

- **[#2588](https://github.com/nanocoai/nanoclaw/issues/2588) — `skill/apple-container` branch is substantially out of sync with mainline**  
  The documented `/convert-to-apple-container` skill will fail immediately due to missing APIs/modules and stale Node/tsc assumptions. Needs maintainer sync or documentation update.

- **[#2589](https://github.com/nanocoai/nanoclaw/issues/2589) — Apple Container: `host.docker.internal` doesn't resolve from inside the microVM**  
  Breaks OneCLI proxy URL resolution in Apple Container. Related Apple Container PR **[#2809](https://github.com/nanocoai/nanoclaw/pull/2809)** is open.

- **[#3161](https://github.com/nanocoai/nanoclaw/pull/3161) — [Fix] redact secrets from host structured logs**  
  Open fix PR for a log-leak issue: structured logs were serializing credentials verbatim into `nanoclaw.log`.

## 6. Feature Requests & Roadmap Signals

Active feature requests and in-flight feature PRs:

- **Native/host runner mode** — [#1732](https://github.com/nanocoai/nanoclaw/issues/1732): bypass Docker for host-tool access.
- **Docker-less operation** — [#1225](https://github.com/nanocoai/nanoclaw/issues/1225).
- **Kubernetes container runtime** — [#2354](https://github.com/nanocoai/nanoclaw/issues/2354).
- **Apple Container runtime + remote OneCLI gateway** — [#2809](https://github.com/nanocoai/nanoclaw/pull/2809).
- **New channel adapters** — Dial SMS/voice ([#3041](https://github.com/nanocoai/nanoclaw/pull/3041)), hosted iMessage via Photon ([#3164](https://github.com/nanocoai/nanoclaw/pull/3164)), unified iMessage adapter ([#3076](https://github.com/nanocoai/nanoclaw/pull/3076)).

Given the release-path fix in [#3163](https://github.com/nanocoai/nanoclaw/pull/3163), a new version may be imminent. The most likely next-version candidates are container-runtime abstraction, the iMessage/Dial channel work, and security hardening fixes.

## 7. User Feedback Summary

- **Positive sentiment:** NanoClaw is appreciated for its minimalist approach and as a lightweight, secure alternative to heavier agent frameworks ([#1184](https://github.com/nanocoai/nanoclaw/issues/1184)).
- **Main pain points:**
  - Hard dependency on Docker / inability to run on restricted Kubernetes environments ([#1184](https://github.com/nanocoai/nanoclaw/issues/1184), [#1225](https://github.com/nanocoai/nanoclaw/issues/1225), [#1732](https://github.com/nanocoai/nanoclaw/issues/1732)).
  - Apple Container support is broken or unreliable ([#2588](https://github.com/nanocoai/nanoclaw/issues/2588), [#2589](https://github.com/nanocoai/nanoclaw/issues/2589)).
  - Security-sensitive behaviors around Telegram pairing and interactive card origin validation ([#3162](https://github.com/nanocoai/nanoclaw/issues/3162), [#2923](https://github.com/nanocoai/nanoclaw/issues/2923)).

Overall, users are enthusiastic about the project’s design goals but are hitting deployment and integration friction in non-Docker and host-connected scenarios.

## 8. Backlog Watch

Long-running or otherwise unresolved items that need maintainer attention:

- **[#1184](https://github.com/nanocoai/nanoclaw/issues/1184)** — Open since 2026-03-17. K8s/Sealos deployment question; still no clear path forward.
- **[#1225](https://github.com/nanocoai/nanoclaw/issues/1225)** — Open since 2026-03-18. Docker-less operation; low priority but frequently requested.
- **[#1732](https://github.com/nanocoai/nanoclaw/issues/1732)** — Open since 2026-04-10. Native runner mode; major feature request.
- **[#2354](https://github.com/nanocoai/nanoclaw/issues/2354)** — Open since 2026-05-08. Kubernetes runtime for agent pods.
- **[#2588](https://github.com/nanocoai/nanoclaw/issues/2588)** — Open since 2026-05-22. Apple Container branch is broken against mainline.
- **[#2651](https://github.com/nanocoai/nanoclaw/pull/2651)** — Open since 2026-05-30. Security fix for interactive response origin validation; directly relevant to [#2923](https://github.com/nanocoai/nanoclaw/issues/2923).
- **[#2809](https://github.com/nanocoai/nanoclaw/pull/2809)** — Open since 2026-06-18. Apple Container runtime + remote OneCLI gateway; would partially address [#2589](https://github.com/nanocoai/nanoclaw/issues/2589).
- **[#2954](https://github.com/nanocoai/nanoclaw/pull/2954)** — Open since 2026-07-04. Security reporting/triage policy docs; waiting on maintainer review.

These items cover the highest-impact gaps in deployment flexibility, security hardening, and documentation consistency.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-08-01

## 1. Today's Overview

NullClaw's activity over the past 24 hours was minimal: zero issues were updated (0 open/active, 0 closed) and no releases were published. The only movement was on one open pull request, **#981**, which was last touched on 2026-07-31. There is no open issue backlog and no reported bugs, indicating a stable—if quiet—project state. The lone in-flight PR adding Grok CLI provider support signals that development continues, but maintainer review velocity appears low. Overall project health is positive: no regressions, no backlog, and one constructive feature contribution awaiting attention.

## 2. Releases

No new releases were published in the last 24 hours, and the latest-releases list is empty. There are no version changes, breaking changes, or migration notes to report.

## 3. Project Progress

- **No PRs were merged or closed** in the last 24 hours.
- The only active item is **[PR #981 — feat(provider): add grok-cli provider for xAI Grok CLI](https://github.com/nullclaw/nullclaw/pull/981)**, still open. It adds an optional CLI-based provider that delegates to the local `grok` executable using the same spawn-per-request pattern as the existing `codex-cli`, `gemini-cli`, and `claude-cli` providers. The feature is implemented and awaiting review/merge but has not landed.

## 4. Community Hot Topics

There is minimal discussion activity to analyze. The only updated PR, **[#981](https://github.com/nullclaw/nullclaw/pull/981)**, has no reported comments or reactions. The underlying signal is nonetheless meaningful: a contributor invested effort to add xAI Grok as a provider, following the established multi-CLI-provider architecture. This indicates user demand for breadth of backend support and a preference for delegating to locally installed official CLIs rather than maintaining API integrations in-house.

## 5. Bugs & Stability

No bugs, crashes, or regressions were reported or updated in the last 24 hours, and no fix PRs are pending. There are no stability concerns to rank at this time.

## 6. Feature Requests & Roadmap Signals

- The sole feature signal is **[PR #981](https://github.com/nullclaw/nullclaw/pull/981)**, requesting/implementing a `grok-cli` provider for xAI Grok.
- The PR deliberately mirrors the existing `codex-cli` / `gemini-cli` / `claude-cli` pattern, suggesting the roadmap is trending toward an ever-widening set of CLI-backed providers. If merged, the next release could reasonably include Grok support alongside the existing providers.
- Longer-term, the repeated use of the spawn-per-request pattern may create pressure to extract shared provider infrastructure (auth-handling, process management) into common utilities.

## 7. User Feedback Summary

Direct user feedback in this window is limited to the contributor of **[PR #981](https://github.com/nullclaw/nullclaw/pull/981)**. The expressed use case is clear: users want to route NullClaw workloads through xAI's official `grok` CLI with the same ergonomics as other CLI providers. The PR also signals awareness of a real pain point—setup friction—by explicitly documenting that the provider is **optional** and requires the `grok` CLI to be installed and authenticated. No broader satisfaction or dissatisfaction signals were captured in the last 24 hours.

## 8. Backlog Watch

- The issue tracker is effectively empty (0 total issues), so there are no long-unanswered issues demanding maintainer attention.
- **[PR #981](https://github.com/nullclaw/nullclaw/pull/981)** has been open since 2026-07-29 (~3 days) with no visible maintainer comments. It is not yet a backlog concern, but if it remains unreviewed for another week, it would become the top item needing maintainer attention.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-08-01

## 1. Today's Overview

IronClaw saw heavy architectural activity in the last 24 hours: 50 PRs were updated (32 merged/closed, 18 open) and 16 issues were touched (15 open, 1 closed). The dominant effort is the WS1 target-architecture refactoring — a stacked series of contract-extraction PRs ([#6967](https://github.com/nearai/ironclaw/pull/6967) → [#6975](https://github.com/nearai/ironclaw/pull/6975) → [#6977](https://github.com/nearai/ironclaw/pull/6977) → [#6980](https://github.com/nearai/ironclaw/pull/6980) → [#6981](https://github.com/nearai/ironclaw/pull/6981)) carving neutral loop/extension/product contracts out of `ironclaw_host_api`, with WS1.1–1.3 now merged. Also merged: hosted MCP server registration ([#6930](https://github.com/nearai/ironclaw/pull/6930)) and a long-pending TOCTOU hardening of the filesystem backend ([#3952](https://github.com/nearai/ironclaw/pull/3952)). Stability is the counterweight: a P0 cross-user memory leak ([#6900](https://github.com/nearai/ironclaw/issues/6900)), a Postgres API capacity regression ([#6973](https://github.com/nearai/ironclaw/pull/6973)), and a libSQL write-stall pathology ([#6974](https://github.com/nearai/ironclaw/issues/6974)) are all active. No new releases were published.

## 2. Releases

No new releases were published in this window. Note: the release PR [#5598](https://github.com/nearai/ironclaw/pull/5598) (bumping `ironclaw_common` to 0.5.0 and `ironclaw_skills` to 0.4.0, both API-breaking) remains open after ~a month (see Backlog Watch).

## 3. Project Progress

**Merged/closed PRs this cycle (WS1 architecture refactor):**
- **WS1.1 — Turn vocabulary completed** ([#6967](https://github.com/nearai/ironclaw/pull/6967)): Completed neutral turn vocabulary in `ironclaw_host_api` and retired the turns shims.
- **WS1.2 — Loop contracts extracted** ([#6975](https://github.com/nearai/ironclaw/pull/6975)): Carved `ironclaw_loop_contracts` out of the turn kernel, flipped `ironclaw_agent_loop` onto it, and landed enforcement + CI registration.
- **WS1.3 — Extension contracts extracted** ([#6977](https://github.com/nearai/ironclaw/pull/6977)): Created `crates/ironclaw_extension_contracts`, repointed every consumer, and closed the dual import paths.

**Feature work merged:**
- **Hosted MCP server registration** ([#6930](https://github.com/nearai/ironclaw/pull/6930), henrypark133): Tenant-runtime registration for hosted MCP servers with automatic no-auth/bearer/OAuth detection, wired into the full extension install→removal lifecycle.
- **TOCTOU-hardened LocalFilesystem** ([#3952](https://github.com/nearai/ironclaw/pull/3952), zmanian): Kernel-race-free tenant FS boundary via fd-relative `openat2`/`O_NOFOLLOW` traversal — a high-leverage multi-tenant security fix, merged after two months.
- **Tool error classification regression fixed** ([#4022](https://github.com/nearai/ironclaw/pull/4022), zmanian): Host-side HTTP response errors are again recoverable/model-visible instead of run-aborting (regression from #4014).
- **Trace refactor follow-up** ([#3942](https://github.com/nearai/ironclaw/pull/3942), zmanian): PilotAllowlist moved from string matching to a serde-driven enum with caller-level error-branch tests.
- **Docs reconciliation** ([#6979](https://github.com/nearai/ironclaw/pull/6979)): Target-architecture docs updated to match the merged #6930 (5 markdown files, +27/−11).
- **Dependency bump** ([#6932](https://github.com/nearai/ironclaw/pull/6932), dependabot): 34-update "everything-else" group bump.

**Still open in the WS1 series:** [#6980](https://github.com/nearai/ironclaw/pull/6980) (WS1.4 — product contracts + adapter half) and [#6981](https://github.com/nearai/ironclaw/pull/6981) (WS1.5 — sealed evidence minting behind witness grants, security-sensitive).

Also closed: issue [#6920](https://github.com/nearai/ironclaw/issues/6920) (target-architecture baselines / WS0 prerequisites), signaling the cleanup foundation is complete.

## 4. Community Hot Topics

- **Path-keyed CI gates tracking** ([#6963](https://github.com/nearai/ironclaw/issues/6963)) — **5 comments**, the most active issue. Filed in response to a review on #6946: a tracking ticket for eight discovered CI/dev gates that still resolve scope from the flat `crates/ironclaw_*` tree shape. Underlying need: the workspace restructuring keeps breaking path-based CI logic, and maintainers are consolidating the defects into one accountable checklist.
- **IronHub skill CTA 404** ([#6940](https://github.com/nearai/ironclaw/issues/6940)) — 2 comments. Every skill's call-to-action 404s; high user visibility, low diagnostic detail so far.
- **Target-architecture baselines** ([#6920](https://github.com/nearai/ironclaw/issues/6920)) — closed with 2 comments; the WS0 prerequisite slice after the #6918 decision record is considered complete.

The stacked WS1 PR series ([#6980](https://github.com/nearai/ironclaw/pull/6980), [#6977](https://github.com/nearai/ironclaw/pull/6977), [#6975](https://github.com/nearai/ironclaw/pull/6975), [#6967](https://github.com/nearai/ironclaw/pull/6967)) is the main source of cross-PR discussion and review load.

## 5. Bugs & Stability

Ranked by severity:

1. **P0 — Cross-user memory leak via shared channels** ([#6900](https://github.com/nearai/ironclaw/issues/6900), security): Unrouted shared conversations (e.g., shared Slack) bind all actors to the operator's memory namespace. Reads/writes must fail closed or bind per-actor. **No fix PR yet** — highest-priority open security bug.
2. **Postgres API capacity regression** ([#6973](https://github.com/nearai/ironclaw/pull/6973), open PR by serrrfirat): Post-#6696 regression took p95 from 3.74s → 12.0s (send_message p95 275ms → 4.78s). **Fix PR open and reportedly recovers capacity.**
3. **libSQL write-stall pathology** ([#6974](https://github.com/nearai/ironclaw/issues/6974)): Tool-heavy stress cases at p95 37–135s post-#6696, split out of #6973. **No dedicated fix yet.**
4. **CI roll-up structurally fails on workflow_dispatch** ([#6978](https://github.com/nearai/ironclaw/issues/6978)): `critical-mutation` is skipped on manual runs but disallowed, so the Tests (Reborn) roll-up is always red even with zero real lane failures. Blocking clean CI signal for the WS1 series.
5. **P2 — Shared home directory / visible workspaces** ([#6866](https://github.com/nearai/ironclaw/issues/6866), privacy): All users share one home directory and can see each other's workspaces. **No fix PR yet.**
6. **IronHub skill CTA 404** ([#6940](https://github.com/nearai/ironclaw/issues/6940)): Broken CTA across every skill.
7. **New-account email auth broken** ([#6972](https://github.com/nearai/ironclaw/issues/6972)): Users cannot authenticate after email-based signup — blocks onboarding.
8. **Linux service install lacks lingering** ([#6976](https://github.com/nearai/ironclaw/issues/6976)): `ironclaw service install` doesn't enable systemd user lingering, so headless/VM/unattended deployments stop running after logout.

## 6. Feature Requests & Roadmap Signals

- **Migration tool for legacy agents** ([#6939](https://github.com/nearai/ironclaw/issues/6939)): Users of Hermes/Openclaw cite high switching costs — no way to carry over setup, config, or memory. Strong adoption-barrier signal; plausible candidate for a near-term tooling workstream.
- **Skills epic — model-driven selection and self-created skills** ([#6941](https://github.com/nearai/ironclaw/issues/6941)): Deliberately carved out of the oversized #6565 with 21 acceptance criteria; pairs with open PR [#6938](https://github.com/nearai/ironclaw/pull/6938) ("the model chooses the skill, not a keyword scorer").
- **"Tools" vs "Extensions" terminology decision** ([#6971](https://github.com/nearai/ironclaw/issues/6971)): Users ask for one consistent term; maintainers need to confirm whether tools/channels remain types of extensions.
- **Admin-Managed Agents as UserId Subjects epic** ([#6578](https://github.com/nearai/ironclaw/issues/6578), p1, security): Tenant admins need non-human subjects without a second identity hierarchy.
- **Rebrand consistency** ([#6854](https://github.com/nearai/ironclaw/issues/6854)): Extension descriptions still say "Reborn"; open PR [#6970](https://github.com/nearai/ironclaw/pull/6970) removes "Reborn" from public docs and is likely to close this.

## 7. User Feedback Summary

Real-user pain points in this window cluster around onboarding, privacy, and naming:

- **Broken first-run flows**: Every skill CTA 404s ([#6940](https://github.com/nearai/ironclaw/issues/6940)) and email-based signup produces accounts that can't authenticate ([#6972](https://github.com/nearai/ironclaw/issues/6972)) — both erode trust at the worst possible moment.
- **Privacy blockers for multi-tenant use**: Shared home directories exposing other users' workspaces ([#6866](https://github.com/nearai/ironclaw/issues/6866)) is a serious objection for any team/enterprise evaluation.
- **Naming and branding confusion**: "Tools" vs "Extensions" ([#6971](https://github.com/nearai/ironclaw/issues/6971)) and lingering "Reborn" branding ([#6854](https://github.com/nearai/ironclaw/issues/6854)) signal the product surface hasn't fully caught up to the "Ironclaw 1.0" external narrative.
- **Switching-cost resistance**: Legacy Hermes/Openclaw users won't start from scratch without a migration path ([#6939](https://github.com/nearai/ironclaw/issues/6939)).
- **Operational reliability**: Headless/VM users need `loginctl enable-linger` behavior in the installer ([#6976](https://github.com/nearai/ironclaw/issues/6976)) for unattended operation.

Overall sentiment: users are engaged and testing deeply, but encountering onboarding, privacy, and naming issues that should be prioritized ahead of 1.0 scale-out.

## 8. Backlog Watch

- **Release PR #5598** ([#5598](https://github.com/nearai/ironclaw/pull/5598), opened 2026-07-03): The `chore: release` PR carrying API-breaking bumps (`ironclaw_common` 0.4.2→0.5.0, `ironclaw_skills` 0.3.0→0.4.0) has been open for ~a month with no merge or visible maintainer response. This is the longest-unaddressed item and is blocking the next published release.
- **Standardized messaging framework** ([#6831](https://github.com/nearai/ironclaw/pull/6831), opened 2026-07-28): Large XL PR adding host-owned standard ops (16 core operations, canonical schemas, 12-code error taxonomy); no recent discussion visible.
- **IronHub deep-link register/install gateway** ([#6780](https://github.com/nearai/ironclaw/pull/6780), opened 2026-07-28): Re-port of #5409 with HMAC register handshake and private manifest sources; waiting on review.
- **Admin-Managed Agents epic** ([#6578](https://github.com/nearai/ironclaw/issues/6578), opened 2026-07-23): Security-relevant p1 epic with only 1 comment — under-visited relative to its importance, and it intersects with the P0 memory-leak issue's identity-binding problem.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-08-01

## Today’s Overview

Project activity in the last 24 hours was maintenance-focused: **4 issues were updated and all 4 were closed as stale**, while **12 PRs were updated — 1 still open and 11 closed/merged**. No new releases were published. The most substantive changes landed around **OpenClaw stability fixes**, especially preserving DeepSeek long-session prefix-cache hit rates, plus small renderer UX improvements. Overall health looks stable, though community discussion volume is low and several user-facing feature requests were auto-closed as stale rather than implemented.

## Releases

**No new releases detected for 2026-08-01.**  
No version notes, breaking changes, or migration guidance to report. A release-prep PR exists ([PR #2416](https://github.com/netease-youdao/LobsterAI/pull/2416)), but no release object was published.

## Project Progress

Merged/closed PRs in the last 24 hours include both fresh fixes and stale cleanup.

**Active/recently merged fixes:**
- [PR #2413](https://github.com/netease-youdao/LobsterAI/pull/2413) — `fix(openclaw): keep live prompt tool-result history byte-stable across turns`  
  Prevents repeated rewriting of cached history, stabilizing DeepSeek prefix-cache behavior.
- [PR #2415](https://github.com/netease-youdao/LobsterAI/pull/2415) — `fix(openclaw): drop aggregate cap in live tool-result prompt projection`  
  Addresses cache-hit-rate collapse by passing `aggregateMaxCharsOverride=null` on live requests.
- [PR #2414](https://github.com/netease-youdao/LobsterAI/pull/2414) — `fix(cowork): prevent BTW tool protocol leakage`  
  Sanitizes provider tool-call markup from side-chat results and preserves error metadata.
- [PR #2417](https://github.com/netease-youdao/LobsterAI/pull/2417) — `fix(sites): add copy success feedback`  
  Reuses conversation copy icon/interaction for site URLs and share codes.

**Release / project management:**
- [PR #2416](https://github.com/netease-youdao/LobsterAI/pull/2416) — `Release/2026.7.31` in `docs`, `main`, and `openclaw` areas.

**Stale PRs closed this cycle:**  
Several older PRs were closed as stale, suggesting they were not merged in their current form:
- [PR #172](https://github.com/netease-youdao/LobsterAI/pull/172) — Antigravity OAuth integration
- [PR #1308](https://github.com/netease-youdao/LobsterAI/pull/1308) — Isolate home-screen input draft per agent
- [PR #1315](https://github.com/netease-youdao/LobsterAI/pull/1315) — Draggable sidebar width
- [PR #1318](https://github.com/netease-youdao/LobsterAI/pull/1318) — Keyboard shortcut kbd hints in sidebar
- [PR #1320](https://github.com/netease-youdao/LobsterAI/pull/1320) — Skeleton loading state for session list
- [PR #1321](https://github.com/netease-youdao/LobsterAI/pull/1321) — Dismiss settings overlays when switching tabs

## Community Hot Topics

There are **no strongly active discussions** in the current data. All four updated issues have exactly **2 comments** and **0 reactions**; PRs report no comment data.

The most “discussed” items are all UI/UX feature requests, now closed as stale:
- [Issue #1311](https://github.com/netease-youdao/LobsterAI/issues/1311) — Table rendering shows raw tags on line breaks; long text needs hover-to-see-full-content.
- [Issue #1314](https://github.com/netease-youdao/LobsterAI/issues/1314) — Sidebar width should be draggable.
- [Issue #1317](https://github.com/netease-youdao/LobsterAI/issues/1317) — Sidebar buttons should show keyboard shortcut hints.
- [Issue #1319](https://github.com/netease-youdao/LobsterAI/issues/1319) — Session list should distinguish “loading” from “empty” state.

Underlying needs: users want **better information density, discoverability, and perceived performance** in the main UI. These were not resolved; they were closed by stale-bot after inactivity.

## Bugs & Stability

No new bug issues were filed today, but several fix PRs were closed/merged for stabilizations.

Ranked by severity:

1. **High — DeepSeek long-session prefix-cache hit-rate regression**  
   Live prompt projection was re-applying an aggregate character cap on every request, rewriting already-cached history whenever new tool results arrived. This dropped DeepSeek cache hit rates from ~100% to ~57%.  
   Fixes: [PR #2413](https://github.com/netease-youdao/LobsterAI/pull/2413) and [PR #2415](https://github.com/netease-youdao/LobsterAI/pull/2415).

2. **Medium — BTW tool protocol leakage in side-chat results**  
   Tool-call markup could leak from provider results into user-visible or downstream contexts.  
   Fix: [PR #2414](https://github.com/netease-youdao/LobsterAI/pull/2414).

3. **Medium / Open — Cron yield descendant finalization**  
   [PR #2234](https://github.com/netease-youdao/LobsterAI/pull/2234) remains open: after `sessions_yield`, descendant-agent completion events may not drive the parent agent, and the finalization loop needs continuation support. This is still waiting for review/merge.

4. **Low — Missing copy success feedback**  
   Site URL / share-code copy actions lacked confirmation.  
   Fix: [PR #2417](https://github.com/netease-youdao/LobsterAI/pull/2417).

## Feature Requests & Roadmap Signals

User-requested features visible in the data all relate to **sidebar and session-list UX**:

- Draggable sidebar width: [Issue #1314](https://github.com/netease-youdao/LobsterAI/issues/1314)
- Keyboard shortcut badges on sidebar buttons: [Issue #1317](https://github.com/netease-youdao/LobsterAI/issues/1317)
- Skeleton loading state for session list: [Issue #1319](https://github.com/netease-youdao/LobsterAI/issues/1319)
- Table content hover-to-expand / raw-tag cleanup: [Issue #1311](https://github.com/netease-youdao/LobsterAI/issues/1311)

All four had associated PRs, but both issues and PRs were closed as stale. This suggests the maintainers did not prioritize them in the current cycle. Near-term roadmap appears focused on **OpenClaw/DeepSeek stability and release hardening** rather than UI feature work. If these UX issues are still desired, they would need to be reopened or re-submitted.

## User Feedback Summary

User-reported pain points in the closed issues were concrete and reasonably common for desktop AI assistants:

- Fixed 240px sidebar is too wide on small screens and too narrow on large screens; long session titles get truncated without a way to expand.
- Keyboard shortcuts exist but are invisible; users discover them only through settings, increasing onboarding friction.
- During app startup, the session list flashes “暂无会话” / “No sessions” while data is still loading, leading users to think their history was lost.
- Table cells render raw HTML/markup on line breaks, and truncated long text has no way to view the full content inline.
- Older PR [PR #1321](https://github.com/netease-youdao/LobsterAI/pull/1321) highlighted a settings-tab bug where modal overlays remained mounted, making the UI appear read-only.

No positive/negative sentiment metrics were available. The most notable signal is that **user-facing feature requests were closed without landing**, which may leave users with unresolved friction unless these items resurface.

## Backlog Watch

Important items needing maintainer attention:

- **Open PR: [PR #2234](https://github.com/netease-youdao/LobsterAI/pull/2234)** — `fix(openclaw): cron yield descendant finalization`  
  Created 2026-06-30, updated 2026-07-31, labeled `stale`, still open. This addresses a real cron/parallel-agent execution bug and should be reviewed before another stale cycle.

- **Stale-closed UX issues/PRs that may need reopening:**  
  - [Issue #1311](https://github.com/netease-youdao/LobsterAI/issues/1311) — table raw tags / hover expansion  
  - [Issue #1314](https://github.com/netease-youdao/LobsterAI/issues/1314) / [PR #1315](https://github.com/netease-youdao/LobsterAI/pull/1315) — draggable sidebar  
  - [Issue #1317](https://github.com/netease-youdao/LobsterAI/issues/1317) / [PR #1318](https://github.com/netease-youdao/LobsterAI/pull/1318) — shortcut hints  
  - [Issue #1319](https://github.com/netease-youdao/LobsterAI/issues/1319) / [PR #1320](https://github.com/netease-youdao/LobsterAI/pull/1320) — skeleton loading  

- **Stale OAuth feature PR: [PR #172](https://github.com/netease-youdao/LobsterAI/pull/172)** — Antigravity OAuth integration  
  Closed as stale; if provider expansion is still planned, this needs an updated follow-up.

All stale closures are from 2026-07-31, indicating automatic lifecycle cleanup rather than active rejection.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest — 2026-08-01

## 1. Today's Overview
Moltis saw moderate activity in the last 24 hours: 2 issues were updated (1 open, 1 closed) and 7 PRs were updated (6 open, 1 closed/merged). No new releases were published. The project’s momentum is concentrated around security hardening, account-level privilege controls, and optional infrastructure features such as Nostr group chat, vector memory, and instrumentation. The closed Markdown copy/export feature also indicates steady progress on user-facing usability.

## 2. Releases
No new releases were published in this window.

## 3. Project Progress
- **Markdown copy and session export** — [PR #1176](https://github.com/moltis-org/moltis/pull/1176) was closed/merged, adding:
  - Preserved Markdown when copying assistant replies.
  - Session-level “Save as Markdown” with full paginated history export.
  - Markdown-safe handling of images.
- This PR closes the long-standing enhancement request [Issue #1131](https://github.com/moltis-org/moltis/issues/1131), which was also closed today.

## 4. Community Hot Topics
Comment activity is low in the visible data, so attention is measured by PR churn, issue status, and reactions. The most notable discussion drivers:

- **Security fixes before adoption** — [PR #1179](https://github.com/moltis-org/moltis/pull/1179) ("verify node pairing signatures") and [PR #1180](https://github.com/moltis-org/moltis/pull/1180) ("harden model and zip paths") were both opened by a contributor who explicitly wants to use Moltis but wants security issues addressed first.
- **Nostr/Buzz group chat** — [PR #1168](https://github.com/moltis-org/moltis/pull/1168) adds NIP-29 group chat support for Buzz channels, expanding Moltis into Block’s open-source agent workspace ecosystem.
- **Privileged command exposure** — [PR #1170](https://github.com/moltis-org/moltis/pull/1170) addresses a real vulnerability class: allowlisted channel users could previously reach privileged commands and host tools without being operators.
- **New GPT 5.6 Luna bug report** — [Issue #1181](https://github.com/moltis-org/moltis/issues/1181) is a fresh open bug report, likely related to model compatibility/behavior.

## 5. Bugs & Stability
Ranked by severity:

- **High — Arbitrary file write via malicious zips / HuggingFace repos** — [PR #1180](https://github.com/moltis-org/moltis/pull/1180) fixes two bug classes that allow writing outside the intended directory, potentially overwriting config, credentials, or scripts and leading to code execution. A fix PR already exists.
- **High — Missing node pairing signature verification** — [PR #1179](https://github.com/moltis-org/moltis/pull/1179) binds `node.pair.verify` to the server-issued pending request, preventing callers from supplying their own key/challenge. Fix PR exists.
- **Medium — Privileged tools accessible to non-operators** — [PR #1170](https://github.com/moltis-org/moltis/pull/1170) fixes a boundary issue where channel access allowlists were conflated with operator privilege. Fix PR exists.
- **Open bug — GPT 5.6 Luna** — [Issue #1181](https://github.com/moltis-org/moltis/issues/1181) is a newly reported bug with no comments or fix PR yet. Exact symptoms are not fully visible in the summary.

## 6. Feature Requests & Roadmap Signals
- **Copy/export as Markdown** — [Issue #1131](https://github.com/moltis-org/moltis/issues/1131) was requested by a user and has now been implemented via [PR #1176](https://github.com/moltis-org/moltis/pull/1176). This is likely to appear in the next release.
- **Nostr NIP-29 group chat support** — [PR #1168](https://github.com/moltis-org/moltis/pull/1168) suggests a strong roadmap direction toward interoperability with Buzz / Nostr-based agent team channels.
- **Zvec vector memory backend** — [PR #1158](https://github.com/moltis-org/moltis/pull/1158) adds an experimental memory backend using Zvec + redb, behind a `zvec` cargo feature. This may become part of the “full” feature set if accepted.
- **Instrumentation and feedback infrastructure** — [PR #1174](https://github.com/moltis-org/moltis/pull/1174) proposes backend-neutral agent instrumentation, Langfuse v4 export, OTLP support, and end-user reaction feedback. This signals interest in production observability and user feedback loops.

## 7. User Feedback Summary
- **Security concerns are a barrier to adoption**: [PR #1179](https://github.com/moltis-org/moltis/pull/1179) author explicitly states they want to use Moltis but need security fixes landed first. This reflects real user friction around trust and safe defaults.
- **Usability demand is being addressed**: The Markdown copy/export feature ([PR #1176](https://github.com/moltis-org/moltis/pull/1176)) had a positive reaction on the original issue ([#1131](https://github.com/moltis-org/moltis/issues/1131), 👍1), indicating users want better ways to export session content.
- **Bleeding-edge model compatibility concerns**: [Issue #1181](https://github.com/moltis-org/moltis/issues/1181) shows users are running newer/edge models like GPT 5.6 Luna and expect Moltis to handle them correctly.

## 8. Backlog Watch
- [PR #1158](https://github.com/moltis-org/moltis/pull/1158) — **Zvec memory backend**, open since 2026-07-17. This is the oldest open PR in the current window and may need a maintainer review or decision, especially since it is described as an experiment.
- [PR #1168](https://github.com/moltis-org/moltis/pull/1168) — **Nostr NIP-29 group chat**, open since 2026-07-25. Significant new functionality; still awaiting merge or review feedback.
- [PR #1170](https://github.com/moltis-org/moltis/pull/1170) — **Operator privileges fix**, open since 2026-07-26. Security-relevant and worth prioritizing for fast review.
- [Issue #1131](https://github.com/moltis-org/moltis/issues/1131) — This enhancement sat open from 2026-06-17 until 2026-07-31. Its closure is positive, but the ~6 week wait may indicate issue triage latency for older items.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest — 2026-08-01

## 1. Today's Overview

QwenPaw/CoPaw activity remained very high over the last 24 hours: 18 issues were updated (13 open, 5 closed) and 43 PRs were updated (30 open, 13 closed/merged). No new release was published. The largest cluster of activity centers on stability fixes for the 2.0.x line: shell-command execution, AgentScope compatibility, memory/compression behavior, and Windows-specific data corruption. The contributor mix includes a notable number of first-time contributors, several with fix PRs already under review. Maintainer review bandwidth appears to be a bottleneck for older but ready-for-review PRs.

## 2. Releases

**None.** No new releases were published in the last 24 hours. The latest referenced versions in issue reports remain `qwenpaw==2.0.1`, `v2.0.0.post2`, and `agentscope==2.0.4.post1`; no changelog, breaking-change notes, or migration guidance to report.

## 3. Project Progress

No releases shipped, but 13 PRs were closed/merged in the last 24 hours. Notable visible closures:

- [PR #6602 — Fix/issue 6558 session integrity](https://github.com/agentscope-ai/QwenPaw/pull/6602) — fixes chat-session UI data integrity issues, preserves in-flight responses when switching Coding/Chat modes.
- [PR #6592 — fix(memory): flush Auto-Memory before Scroll context eviction](https://github.com/agentscope-ai/QwenPaw/pull/6592) — addresses the Dream/memory early-event loss reported in #6555.
- [PR #6573 — fix(audio): restore transcription for channel audio messages](https://github.com/agentscope-ai/QwenPaw/pull/6573) — fixes Feishu audio message transcription failure in 2.x.
- [PR #6606 — fix(read_file): accept numeric string line ranges](https://github.com/agentscope-ai/QwenPaw/pull/6606) — tool input robustness fix.
- [PR #6604 — docs(memory): explain ReMe self-evolving knowledge base](https://github.com/agentscope-ai/QwenPaw/pull/6604) — documentation for memory lifecycle and Auto-Dream.

Open PRs advancing features/fixes include:

- [PR #6610 — fix: shell command execution hangs and UI freezes](https://github.com/agentscope-ai/QwenPaw/pull/6610) — targets #6608 and #6589.
- [PR #6609 — Fix spawn subagent schema](https://github.com/agentscope-ai/QwenPaw/pull/6609) — fixes #6588.
- [PR #6615 — fix(agentscope): compatibility and config loading issues](https://github.com/agentscope-ai/QwenPaw/pull/6615) — targets #6612.
- [PR #6528 — fix: resolve agent.json corruption](https://github.com/agentscope-ai/QwenPaw/pull/6528) — targets #6520.
- [PR #6526 — feat: Add NVIDIA NIM provider support](https://github.com/agentscope-ai/QwenPaw/pull/6526) — new provider integration.
- [PR #6607 — feat(desktop): global-hotkey floating quick-input window](https://github.com/agentscope-ai/QwenPaw/pull/6607) — desktop UX feature.

## 4. Community Hot Topics

The most active discussion remains around long-session reliability and configuration durability:

- [Issue #6537 — Skill tags disappear on restart](https://github.com/agentscope-ai/QwenPaw/issues/6537) — 10 comments. A regression of #3270; tags are saved but lost on manifest reconciliation. High user impact because configuration does not survive restarts.
- [Issue #6601 — QwenPaw does not report empty-response errors](https://github.com/agentscope-ai/QwenPaw/issues/6601) — 5 comments. In long sessions the model can return empty responses without any framework-level error, effectively killing the session.
- [Issue #6588 — `spawn_subagent` single-task mode unusable](https://github.com/agentscope-ai/QwenPaw/issues/6588) — 4 comments. Tool schema incorrectly makes `batch` required, blocking a documented use case.
- [Issue #6083 — Desktop workspace output quick-access button](https://github.com/agentscope-ai/QwenPaw/issues/6083) — 4 comments. Users want one-click access to generated files.
- [Issue #6160 — Bundled/independent Python runtime for Desktop](https://github.com/agentscope-ai/QwenPaw/issues/6160) — 4 comments. A recurring ask from Windows users without a system Python.
- [Issue #6260 — Improve result presentation](https://github.com/agentscope-ai/QwenPaw/issues/6260) — 2 comments, 1 reaction. Users want thinking/tool-call output collapsed so delivered results are not buried.

Underlying needs: users are pushing for more production-grade reliability, clearer output presentation, and less fragile desktop/configuration behavior.

## 5. Bugs & Stability

Ranked roughly by severity:

1. **[Issue #6520 — `agent.json` systematic corruption](https://github.com/agentscope-ai/QwenPaw/issues/6520)** — BOM, missing quotes, double-encoded Chinese text across ~20+ fields; causes complete system failure on Windows. Fix PR: [PR #6528](https://github.com/agentscope-ai/QwenPaw/pull/6528) open.
2. **[Issue #6612 — Incompatibility with `agentscope==2.0.4.post1`](https://github.com/agentscope-ai/QwenPaw/issues/6612)** — proactive/memory subsystem crashes and tool-permission deadlock caused by AgentScope API changes. Fix PR: [PR #6615](https://github.com/agentscope-ai/QwenPaw/pull/6615) open.
3. **[Issue #6608 — Long-running shell commands bypass `shell_command_timeout` and block Feishu sessions](https://github.com/agentscope-ai/QwenPaw/issues/6608)** — a dedup script blocked a channel session for 1.5 hours; orphan subprocess on cancel. Fix PR: [PR #6610](https://github.com/agentscope-ai/QwenPaw/pull/6610) open.
4. **[Issue #6589 — `execute_shell_command` large output freezes UI](https://github.com/agentscope-ai/QwenPaw/issues/6589)** — tens of thousands of lines render synchronously and block the frontend. Fix PR: [PR #6610](https://github.com/agentscope-ai/QwenPaw/pull/6610) open.
5. **[Issue #6614 — WeChat cron push silently never delivers](https://github.com/agentscope-ai/QwenPaw/issues/6614)** — tasks report `success` but WeChat returns `ret=-2` / invalid `context_token`; user reports ~44M tokens burned. No visible fix PR yet.
6. **[Issue #6537 — Skill tags disappear on restart](https://github.com/agentscope-ai/QwenPaw/issues/6537)** — configuration regression; no visible fix PR yet.
7. **[Issue #6588 — `spawn_subagent` schema blocks single-task mode](https://github.com/agentscope-ai/QwenPaw/issues/6588)** — fix PR: [PR #6609](https://github.com/agentscope-ai/QwenPaw/pull/6609) open.
8. **[Issue #6601 — Empty responses not reported in long sessions](https://github.com/agentscope-ai/QwenPaw/issues/6601)** — no visible fix PR yet.
9. **[Issue #6555 — Dream/memory compression misses early-session events](https://github.com/agentscope-ai/QwenPaw/issues/6555)** — fixed in [PR #6592](https://github.com/agentscope-ai/QwenPaw/pull/6592) closed; alternate fix [PR #6564](https://github.com/agentscope-ai/QwenPaw/pull/6564) still open.
10. **[Issue #6544 — Feishu audio messages silently fail transcription](https://github.com/agentscope-ai/QwenPaw/issues/6544)** — fixed in [PR #6573](https://github.com/agentscope-ai/QwenPaw/pull/6573) closed.
11. **[Issue #6558 — Multiple chat session UI data-integrity issues](https://github.com/agentscope-ai/QwenPaw/issues/6558)** — fixed in [PR #6602](https://github.com/agentscope-ai/QwenPaw/pull/6602) closed.
12. **[Issue #6512 — `execute_shell_command` large output truncation / internal error](https://github.com/agentscope-ai/QwenPaw/issues/6512)** — related to shell output handling; no dedicated visible fix PR yet.
13. **[Issue #6529 — ACP `new_session` response missing `models` field](https://github.com/agentscope-ai/QwenPaw/issues/6529)** — closed; no visible fix PR in this sample.

## 6. Feature Requests & Roadmap Signals

User-requested features in the last 24 hours:

- [Issue #6083 — Workspace output quick-access button in Desktop](https://github.com/agentscope-ai/QwenPaw/issues/6083) — one-click access to agent-generated files.
- [Issue #6160 — Independent or bundled Python runtime for Desktop](https://github.com/agentscope-ai/QwenPaw/issues/6160) — avoid dependence on a system Python for generated scripts.
- [Issue #6260 — Collapsible thinking/tool-call details, result-first presentation](https://github.com/agentscope-ai/QwenPaw/issues/6260).
- [Issue #6512 — Large shell output should auto-write to a file or support streaming reads](https://github.com/agentscope-ai/QwenPaw/issues/6512).
- [Issue #6587 — Rename Desktop app from “QwenPaw Desktop” to “QwenPaw”](https://github.com/agentscope-ai/QwenPaw/issues/6587).

Roadmap signals from open PRs:

- [PR #6526 — NVIDIA NIM provider support](https://github.com/agentscope-ai/QwenPaw/pull/6526) — likely to land in a future provider-focused release.
- [PR #6607 — Global-hotkey quick-input floating window](https://github.com/agentscope-ai/QwenPaw/pull/6607) — desktop UX accelerator.
- [PR #6302 — Unified provider discovery, model metadata, routing, and Agent controls](https://github.com/agentscope-ai/QwenPaw/pull/6302) — larger architectural change.
- [PR #6611 — Align Scroll and memory with AgentScope lifecycle](https://github.com/agentscope-ai/QwenPaw/pull/6611) — likely required for memory stability.

Prediction: the next version will likely prioritize the already-open fix PRs for AgentScope compatibility (#6615), shell-command stability (#6610), and agent.json corruption (#6528), with NVIDIA NIM support (#6526) and the desktop quick-input window (#6607) as likely feature additions.

## 7. User Feedback Summary

Real user pain points visible in this 24h sample:

- **Windows desktop fragility** is over-represented: corrupted `agent.json`, missing Python environments, obscured input box, and UI freezes from large shell output.
- **Silent failures are especially damaging**: WeChat cron reports success but never delivers; Feishu audio messages fail transcription without errors; long shell commands hang sessions indefinitely.
- **Non-technical users frequently mention workflow interruption** — e.g. having to leave the desktop app and navigate to `~/.qwenpaw/workspaces/...` to access generated reports, CSVs, and images.
- **Result presentation is a recurring dissatisfaction**: users feel that chain-of-thought and tool logs overwhelm the actual delivered result.
- **Chinese-language user feedback dominates the current issue set**, indicating a strong user base among Chinese-speaking desktop/Feishu/WeChat users.
- No positive feedback was captured in this 24h window; the current sentiment is predominantly bug-report and usability-request driven.

## 8. Backlog Watch

Issues and PRs that need maintainer attention:

- [Issue #6083 — Desktop workspace quick access](https://github.com/agentscope-ai/QwenPaw/issues/6083) — open since 2026-07-14, 4 comments, no visible maintainer response.
- [PR #6203 — fix(utils): bound and hide Windows tasklist liveness probe](https://github.com/agentscope-ai/QwenPaw/pull/6203) — open since 2026-07-16, marked “Under Review / ready-for-human-review” but still not merged.
- [Issue #6160 — Bundled Python runtime](https://github.com/agentscope-ai/QwenPaw/issues/6160) — open since 2026-07-16, 4 comments, no clear resolution path.
- [PR #6302 — Unify provider discovery, model metadata, routing, Agent controls](https://github.com/agentscope-ai/QwenPaw/pull/6302) — open since 2026-07-21; large architectural PR needing review.
- [Issue #6260 — Improve result presentation](https://github.com/agentscope-ai/QwenPaw/issues/6260) — open since 2026-07-19, user support via reaction, no roadmap answer yet.
- [PR #6528 — fix agent.json corruption](https://github.com/agentscope-ai/QwenPaw/pull/6528) — open since 2026-07-28, targets a critical issue (#6520), should be prioritized.
- [Issue #6512 — Large shell output truncation / internal error](https://github.com/agentscope-ai/QwenPaw/issues/6512) — open since 2026-07-28, no dedicated fix PR visible.
- [Issue #6537 — Skill tags disappear on restart](https://github.com/agentscope-ai/QwenPaw/issues/6537) — high comment count, no fix PR visible; needs maintainer triage.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-08-01

## Today's Overview

ZeroClaw is in a period of sustained high activity: **27 issues** (25 open) and **50 PRs** (38 open, 12 merged/closed) were touched in the last 24 hours, with **no new releases**. The project is heavily focused on architecture via RFC — at least 15 RFCs or RFC-trackers are active, covering WASM plugin sandboxing, OpenAI API compatibility, LSP support, and security-policy contracts. Notably, the maintainers appear to be converting accepted PR review feedback into structured "follow-up" tracking issues (a wave of issues #9590–#9602 filed 2026-07-31), indicating a disciplined, process-heavy contribution workflow with risk/size labeling and maintainer decision queues. Overall health signals are strong: security bugs are being closed with fix PRs, while the RFC pipeline is the main bottleneck for new capability delivery.

## Releases

**No new releases** in the last 24 hours. The latest release stream is unchanged; most recent activity is on `master` via PRs targeting the next unreleased version.

## Project Progress

**12 PRs merged/closed** in the past 24 hours. From the observable set, the notable closures:

- **[#9114 — fix(runtime/security): allow various devices and files on landlock sandbox](https://github.com/zeroclaw-labs/zeroclaw/pull/9114)** *(closed, P1, high risk)* — Policy-hardening follow-up that unblocks the shell tool under Landlock; closes bug [#8973](https://github.com/zeroclaw-labs/zeroclaw/issues/8973) (Fedora `/dev/null` access failure).
- **[#9294 — fix(web): keep sidebar selection mutually exclusive](https://github.com/zeroclaw-labs/zeroclaw/pull/9294)** *(closed)* — Resolves route-segment matching so `/config/agents` no longer also highlights its Config parent.
- **[#9087 — fix(robot-kit): bound TTS and audio playback subprocess waits](https://github.com/zeroclaw-labs/zeroclaw/pull/9087)** *(closed)* — Adds deadlines to `piper`/`aplay`/`paplay` subprocess waits, preventing hangs on headless or slow robot hardware.
- **[#9369 — fix(containers): serialize shared Cargo caches](https://github.com/zeroclaw-labs/zeroclaw/pull/9369)** *(closed)* — Uses `sharing=locked` on BuildKit cache mounts to stop concurrent-stage Cargo extraction races.
- **[#9367 — test(runtime): isolate safety-net test workspaces](https://github.com/zeroclaw-labs/zeroclaw/pull/9367)** *(closed)* — Moves five agent-construction safety-net tests off shared `/tmp` into dedicated temporary workspaces.
- **[#9344 — chore(deps): bump anchore/sbom-action to v0.24.0](https://github.com/zeroclaw-labs/zeroclaw/pull/9344)** *(closed)* — SBOM action pinned to full SHA v0.24.0 for stable releases.
- **[#9525 — refactor(agent): split history for before-llm-call hook — PR-A](https://github.com/zeroclaw-labs/zeroclaw/pull/9525)** *(closed, behavioral-equivalence refactor)* — Separates `loop_history` from `loop_new_messages`/`round_added`, preparing for a future `before_llm_call` hook without changing runtime behavior.

Two issues were also closed: the Landlock bug **[#8973](https://github.com/zeroclaw-labs/zeroclaw/issues/8973)** (fixed by #9114) and **[#9539](https://github.com/zeroclaw-labs/zeroclaw/issues/9539)** enabling Dependabot security updates for transitive lockfile advisories.

## Community Hot Topics

Most-commented issues in the last 24 hours:

- **[#8550 — Add OpenAI-compatible chat completions endpoint](https://github.com/zeroclaw-labs/zeroclaw/issues/8550)** *(6 comments)* — Demand for standard client compatibility (Open WebUI, LobeChat, OpenAI SDK, Continue.dev, Aider). A large counterpart PR **[#8486](https://github.com/zeroclaw-labs/zeroclaw/pull/8486)** (XL, needs-author-action) implementing the endpoint is open — this is the single clearest interoperability ask.
- **[#5907 — RFC: Opt-in LSP support for ZeroCode coding workflows](https://github.com/zeroclaw-labs/zeroclaw/issues/5907)** *(5 comments)* — Language-server support to reduce hallucination in local-model codegen, referencing Claude Code/OpenCode precedent. Open since 2026-04-19 with `needs-author-action`.
- **[#8692 — Tracker: Maintainer decision queue for RFCs and design issues](https://github.com/zeroclaw-labs/zeroclaw/issues/8692)** *(5 comments)* — The project's own bottleneck: an issue-level queue to route RFCs to maintainer attention. High engagement signals a backlog of decisions pending.
- **[#8135 — RFC: Wasm-first plugin runtime](https://github.com/zeroclaw-labs/zeroclaw/issues/8135)**, **[#8078 — RFC: zerocode local pre-submission gate](https://github.com/zeroclaw-labs/zeroclaw/issues/8078)**, **[#8973 — Landlock blocks shell on Fedora](https://github.com/zeroclaw-labs/zeroclaw/issues/8973)**, and **[#9246 — RFC: preserve Todo tracker config during ZeroCode migration](https://github.com/zeroclaw-labs/zeroclaw/issues/9246)** — all at 4 comments.

Underlying need: users and contributors are pushing ZeroClaw toward (a) **drop-in compatibility with the OpenAI ecosystem**, (b) **stronger local/coding workflows** (LSP, pre-submission gates), and (c) **secure sandboxing** that does not break everyday functionality like shell execution.

## Bugs & Stability

Bugs updated in the last 24 hours, ranked by severity:

**S1 — workflow blocked:**
- **[#9591 — fix(channels): clear delivery registry when reload removes all channels](https://github.com/zeroclaw-labs/zeroclaw/issues/9591)** *(P1, accepted, follow-up)* — A reload that removes all configured channels leaves a stale `CRON_CHANNEL_REGISTRY`, so deliveries can target vanished channels.
- **[#9601 — ci(security): diagnose missing Dependabot PRs for transitive Cargo alerts](https://github.com/zeroclaw-labs/zeroclaw/issues/9601)** *(P1, accepted, follow-up)* — Direct npm alerts produce Dependabot PRs but resolvable transitive Cargo alerts do not; four occurrences found in verification for #9539.

**S1/S2 — security:**
- **[#8973 — Landlock blocks shell access to required system files on Fedora](https://github.com/zeroclaw-labs/zeroclaw/issues/8973)** *(closed)* — `sh` could not access `/dev/null` under the sandbox; resolved by merged PR **[#9114](https://github.com/zeroclaw-labs/zeroclaw/pull/9114)**.

**S2 — degraded behavior / correctness:**
- **[#9592 — fix(tools): probe the saved provider alias after model-routing updates](https://github.com/zeroclaw-labs/zeroclaw/issues/9592)** *(P1, accepted)* — The probe reads from the pre-update runtime snapshot rather than the newly saved alias config.
- **[#9546 — updater web-dist test depends on host installation state](https://github.com/zeroclaw-labs/zeroclaw/issues/9546)** *(P2)* — CI test assumes binary-adjacent `web/dist` path; fails on dev machines with host-installed artifacts.
- **[#9594 — Coding-agent tools charge the action budget twice](https://github.com/zeroclaw-labs/zeroclaw/issues/9594)** *(P2, accepted)* — `SecurityPolicy::enforce_tool_operation` records one action at the tool layer and again at production wrapper level.

**Lower severity:**
- **[#9590 — Concurrent `models refresh` runs can lose cache entries](https://github.com/zeroclaw-labs/zeroclaw/issues/9590)** *(P3, accepted)* — Unlocked read-modify-write on the shared cache file across processes.

Also active in the PR pipeline (open, fixing channel/cron reliability): **[#9524](https://github.com/zeroclaw-labs/zeroclaw/pull/9524)** (Signal/Voice Call channels crash-loop on missing credentials), **[#9313](https://github.com/zeroclaw-labs/zeroclaw/pull/9313)** (WeChat sync-cursor persistence before enqueue), **[#9320](https://github.com/zeroclaw-labs/zeroclaw/pull/9320)** (cron jobs hold sqlite lock forever on hung runs), and **[#9002](https://github.com/zeroclaw-labs/zeroclaw/pull/9002)** (agent turns cancelled on viewer disconnect).

## Feature Requests & Roadmap Signals

The open RFC portfolio defines the likely roadmap:

- **OpenAI-compatible gateway** — [#8550](https://github.com/zeroclaw-labs/zeroclaw/issues/8550) with implementation PR [#8486](https://github.com/zeroclaw-labs/zeroclaw/pull/8486) in flight. High probability for near-term release.
- **Wasm-first plugin runtime** — [#8135](https://github.com/zeroclaw-labs/zeroclaw/issues/8135) (default-on Wasm plugins, signed distribution) and [#8187](https://github.com/zeroclaw-labs/zeroclaw/issues/8187) (capability-gated WASI hardware functions) would eliminate Node.js native plugin risk and enable GPIO/SPI/I2C/USB access via sandboxed modules.
- **Coding workflow enhancements** — [#5907](https://github.com/zeroclaw-labs/zeroclaw/issues/5907) (LSP support) and [#8078](https://github.com/zeroclaw-labs/zeroclaw/issues/8078) (local `zerocode` pre-submission gate) would materially improve ZeroCode's credibility for agentic software engineering.
- **AI-assisted PR review** — [#9330](https://github.com/zeroclaw-labs/zeroclaw/issues/9330) proposes CI-triggered AI pre-review/re-review with human-owned final approval; consistent with the project's existing risk-based review lanes.
- **Contract/architecture consolidation** — [#9346](https://github.com/zeroclaw-labs/zeroclaw/issues/9346) (unified package/capability/config catalog), [#9598](https://github.com/zeroclaw-labs/zeroclaw/issues/9598) (SOP capability permission contract), [#9593](https://github.com/zeroclaw-labs/zeroclaw/issues/9593) (`TaskRecord` as single lifecycle owner), [#9595](https://github.com/zeroclaw-labs/zeroclaw/issues/9595) (single provider family registry), and [#9600](https://github.com/zeroclaw-labs/zeroclaw/issues/9600) (session-persistence contract ownership) — a batch of accepted follow-ups indicating a refactoring wave after recent review cycles.
- **Observability** — PR [#9352](https://github.com/zeroclaw-labs/zeroclaw/pull/9352) (XL) adds cross-turn `conversation_id` correlation to OTel, aligned with RFC #8933.

Prediction: the next release will likely include the OpenAI chat completions endpoint, the Landlock hardening (#9114), channel reliability fixes, and the merged history-split refactor (#9525), with the Wasm-first runtime and LSP support landing in a later milestone once RFC decisions finalize.

## User Feedback Summary

- **Interoperability is the top ask.** Issue #8550 (OpenAI-compatible endpoint) comes from a user who explicitly documents that Open WebUI, LobeChat, LangChain, OpenAI SDK, Continue.dev, and Aider cannot connect to ZeroClaw without custom work. This is a recurring real-world adoption blocker.
- **Sandboxing must not break basic functionality.** The Fedora Landlock report (#8973) was a concrete S2 regression ("shell tool always fails because `sh` cannot access `/dev/null`") and was resolved promptly — a positive signal for user confidence in security defaults.
- **Local-model users want hallucination guardrails.** #5907 frames LSP support as a "useful backstop" for local models, implicitly reporting that current codegen quality from local models without LSP is a pain point.
- **Operational reliability concerns surfaced in PRs:** WeChat sync-cursor data-loss window (#9313), cron jobs holding locks forever on hung providers (#9320), Signal/Voice Call channels crash-looping on empty credentials (#9524), and agent turns being cancelled when a dashboard viewer disconnects (#9002) all reflect real deployments hitting edge conditions.
- **Resource accounting complaints:** #9594 (double action-budget charging) and #9590 (concurrent `models refresh` losing cache entries) are correctness/trust issues that can erode confidence in usage metering and CLI behavior.

## Backlog Watch

- **[#5907 — RFC: Opt-in LSP support](https://github.com/zeroclaw-labs/zeroclaw/issues/5907)** — Open since **2026-04-19** (~3.5 months), 5 comments, labeled `needs-author-action` with `risk:high`. One of the oldest open RFCs; requires author response or maintainer decision.
- **[#8692 — Maintainer decision queue tracker](https://github.com/zeroclaw-labs/zeroclaw/issues/8692)** — The meta-issue for unblocking RFCs; its own open status indicates the decision pipeline is the critical path for the project's architecture work.
- **[#8691 — Restore ADR baseline and audit accepted RFC decision records](https://github.com/zeroclaw-labs/zeroclaw/issues/8691)** — Docs/architecture tracker, `status:accepted` but no recent movement; accepted RFCs may lack follow-through decision records.
- **Stalled XL PRs** needing author action: **[#9214 — live eval execution mode](https://github.com/zeroclaw-labs/zeroclaw/pull/9214)**, **[#9352 — OTel conversation correlation](https://github.com/zeroclaw-labs/zeroclaw/pull/9352)**, and **[#8486 — OpenAI chat completions endpoint](https://github.com/zeroclaw-labs/zeroclaw/pull/8486)** are large, high-value contributions blocked on author responses; maintainers may need to intervene or shepherd these to completion.
- **[#9544 — fix(delegate): honor configured provider fallbacks](https://github.com/zeroclaw-labs/zeroclaw/pull/9544)** — Trusted-contributor fix for delegated sessions not materializing root-configured route/fallback/retry settings; `needs-author-action` despite high risk and real correctness impact.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*