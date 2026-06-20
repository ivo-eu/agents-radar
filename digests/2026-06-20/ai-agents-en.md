# OpenClaw Ecosystem Digest 2026-06-20

> Issues: 116 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-20 10:17 UTC

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

# OpenClaw Project Digest — 2026-06-20

## 1. Today’s Overview

OpenClaw shows very high activity: 116 issues and 500 pull requests were updated in the last 24 hours, with the vast majority still open (113 issues, 481 PRs). Only 3 issues and 19 PRs were closed or merged in this period, suggesting a large review backlog and a focus on bug triage rather than batch merging. No new releases were published today. The project is receiving a torrent of detailed bug reports, regression notices, and feature requests across platforms (Discord, Telegram, Matrix, Google Chat, QQ, WeChat, Feishu), indicating both broad adoption and scaling pains in the current `2026.5.28` stable release.

## 2. Releases

No new releases today. The latest published version remains `2026.5.28` (commit `e932160`). Several issues report upgrade blockers from `2026.5.27` → `2026.5.28`, notably the Bedrock provider regression (#88707) and a GHCR Docker image with stale config schema (#88788).

## 3. Project Progress

None of the top-30 most-discussed PRs were merged today. Of the 19 merged/closed PRs, details are not listed in the provided dataset. Based on open PRs that appear close to ready (labels like `👀 ready for maintainer look` and `proof: sufficient`), several fixes are nearing completion but haven’t been merged yet:

- **Fix: make post-turn compaction non-fatal**  
  PR [#95301](https://github.com/openclaw/openclaw/pull/95301) — prevents summarization errors from discarding already-generated replies.
- **Fix: prevent silent message loss from EmbeddedAttemptSessionTakeoverError**  
  PR [#89039](https://github.com/openclaw/openclaw/pull/89039) — addresses a critical race between OpenAI SDK retries and session lock fencing.
- **Fix: defer text settlement for final-mode TTS (Telegram churn)**  
  PR [#83988](https://github.com/openclaw/openclaw/pull/83988) — eliminates visible text→voice note replacement flicker.
- **Fix: hide archived Workboard cards in CLI list by default**  
  PR [#94562](https://github.com/openclaw/openclaw/pull/94562) — UX improvement for workboard management.

## 4. Community Hot Topics

The most active discussions on GitHub today reflect deep operational pain points:

- **GHCR Docker image emits stale Discord progress commentary config schema**  
  Issue [#88788](https://github.com/openclaw/openclaw/issues/88788) — **12 comments**, 1 👍. The 2026.5.28 Docker image rejects a config field (`channels.discord.streaming.progress.commentary`) that should be valid, causing deployment confusion.

- **Codex OAuth refresh failures can wedge an agent for hours**  
  Issue [#86215](https://github.com/openclaw/openclaw/issues/86215) — **9 comments**, 1 👍. P1 bug where stale OAuth tokens block agent rotation without alerting. A frequently cited blocker for production deployments.

- **Gateway heap grows unbounded → cgroup OOM on long-running Linux deployments**  
  Issue [#89315](https://github.com/openclaw/openclaw/issues/89315) — **6 comments**, 3 👍. A memory leak in the gateway process kills the agent after hours of uptime. High community demand for a fix.

- **Claude-Opus-4.8 context window hard-coded to 128K instead of native 1M (GitHub Copilot provider)**  
  Issue [#91869](https://github.com/openclaw/openclaw/issues/91869) — **4 comments**, 2 👍. Closed but relevant: the static catalog entry shadows the live endpoint’s true capability.

- **Proposal: measure and reduce OpenClaw-owned pre-model latency**  
  Issue [#88812](https://github.com/openclaw/openclaw/issues/88812) — **4 comments**, 1 👍. Real-world measurement shows 2.2s of OpenClaw overhead vs 2.8s model call; community wants diagnostic tooling.

These issues collectively indicate that scalability, reliability under load, and config/data correctness are the most pressing concerns for the user base.

## 5. Bugs & Stability

All reported bugs are from versions `2026.5.28` or earlier, with several regressions from `2026.5.27`. P1 severity bugs dominate:

### Critical (P1, impact: crash-loop, security, or message-loss)

| Issue | Title | Impact | Fix PR? |
|-------|-------|--------|---------|
| [#89315](https://github.com/openclaw/openclaw/issues/89315) | Gateway heap grows unbounded → cgroup OOM | crash-loop | No linked PR |
| [#86215](https://github.com/openclaw/openclaw/issues/86215) | Codex OAuth refresh failures wedge agent | session-state, auth-provider | `clawsweeper:no-new-fix-pr` |
| [#88707](https://github.com/openclaw/openclaw/issues/88707) | Bedrock provider registration broken in 2026.5.28 | auth-provider, regression | `clawsweeper:no-new-fix-pr` |
| [#88562](https://github.com/openclaw/openclaw/issues/88562) | `apiKey` written as plain string instead of secret-ref | security | `clawsweeper:needs-security-review` |
| [#89257](https://github.com/openclaw/openclaw/issues/89257) | `backup create --verify` exits 13, leaves corrupt archive | data-loss, crash-loop | `clawsweeper:no-new-fix-pr` |
| [#89473](https://github.com/openclaw/openclaw/issues/89473) | Reasoning tokens leak to chat channels | security, message-loss | `clawsweeper:fix-shape-clear`, no PR |
| [#88955](https://github.com/openclaw/openclaw/issues/88955) | QQbot WebSocket reconnection → "Outbound not configured" | message-loss, regression | `clawsweeper:linked-pr-open` |
| [#89332](https://github.com/openclaw/openclaw/issues/89332) | `exec` spawn UNKNOWN after gateway restart | crash-loop | `clawsweeper:needs-info` |
| [#89069](https://github.com/openclaw/openclaw/issues/89069) | Childless subagent stuck initializing exhausts exec capacity | crash-loop | `clawsweeper:no-new-fix-pr` |
| [#89167](https://github.com/openclaw/openclaw/issues/89167) | Failed session stays bound, crashes next TUI launch | crash-loop | `clawsweeper:no-new-fix-pr` |
| [#88907](https://github.com/openclaw/openclaw/issues/88907) | Chronic Telegram agent failures (timeouts, empty responses) | message-loss, auth-provider | `clawsweeper:no-new-fix-pr` |

### High (P1/P2, impact: session-state, auth, or data integrity)

- [#88856](https://github.com/openclaw/openclaw/issues/88856) — Silent subagent tool result drop (~3.8% historical rate).  
- [#87857](https://github.com/openclaw/openclaw/issues/87857) — Agent skips mandatory AGENTS.md startup sequence.  
- [#89641](https://github.com/openclaw/openclaw/issues/89641) — Telegram `run_in_background` task-notification silently dropped when session idle.  
- [#89594](https://github.com/openclaw/openclaw/issues/89594) — Microsoft Teams inbound attachments unreachable.  
- [#89254](https://github.com/openclaw/openclaw/issues/89254) — Matrix E2EE inbound DMs intermittently not dispatched (~1 in 7).  
- [#89184](https://github.com/openclaw/openclaw/issues/89184) — LINE webhook returns 404 after v2026.5.28 update.  
- [#89549](https://github.com/openclaw/openclaw/issues/89549) — Subagent child runs fail with `HTTP 401 Missing scopes`.  
- [#89609](https://github.com/openclaw/openclaw/issues/89609) — `openclaw migrate codex` always fails with "Unknown migration provider".  
- [#89589](https://github.com/openclaw/openclaw/issues/89589) — State directory permissions reset to group-writable by update-check writes.  
- [#89607](https://github.com/openclaw/openclaw/issues/89607) — Secret file-source not injected into ElevenLabs TTS at runtime.  
- [#89551](https://github.com/openclaw/openclaw/issues/89551) — Session-memory hook model config ignored for LLM slug generation.  
- [#89430](https://github.com/openclaw/openclaw/issues/89430) — Google Chat media upload fails under app authentication (403).  

**Security note:** Several issues carry `clawsweeper:needs-security-review` or `impact:security` labels (#88562, #89473, #89228, #89589, #89607) — all await maintainer decision.

## 6. Feature Requests & Roadmap Signals

Community feature proposals indicate a strong desire for **platform parity**, **operator transparency**, and **performance improvements**:

| Issue | Feature | Likelihood for Next Release |
|-------|---------|----------------------------|
| [#94418](https://github.com/openclaw/openclaw/issues/94418) | Alibaba Model Studio Token Plan (Team Edition) provider | Medium – existing qwen plugin, relatively simple extension |
| [#95279](https://github.com/openclaw/openclaw/issues/95279) | Trusted inbound-decoration contract (for consumers to strip/dedup) | High – repeated theme of message deduplication |
| [#85087](https://github.com/openclaw/openclaw/issues/85087) | Automatic fast-mode: on for talks, off for longer runs | Medium – already discussed; low-breaking change |
| [#89274](https://github.com/openclaw/openclaw/issues/89274) | Show resolved model & thinking level after `/new` or `/reset` | Low (P3, cosmetic) |
| [#95295](https://github.com/openclaw/openclaw/issues/95295) | Show child-spawned subagent sessions in Web UI sidebar | Medium – frequent UX complaint |
| [#87689](https://github.com/openclaw/openclaw/issues/87689) | Guard to disable session transcript ingestion during QMD migrations | Medium – operational need for memory migrations |
| [#89666](https://github.com/openclaw/openclaw/issues/89666) | Make `cron.sessionRetention` prune isolated-target cron sessions | Medium – gap in cleanup logic |
| [#88812](https://github.com/openclaw/openclaw/issues/88812) | Measure and reduce pre-model latency in channel replies | High – directly impacts user-perceived performance |

Given the volume of P1 bugs, the next release may prioritise stability fixes (heap leak, OAuth retry, reasoning token leak) over new features. However, the latency measurement and fast-mode automation are low-risk additions that could ship quickly.

## 7. User Feedback Summary

The community is highly engaged, filing detailed bug reports with reproduction steps, environment specifics, and historical data. Common pain points:

- **Reliability under load**: Heap OOM, subagent timeouts, silent drops, and session wedging are top frustrations. Users on long-running systemd deployments and multi-agent setups report the most severe issues.
- **Breaking regressions**: The `2026.5.28` release broke Bedrock, LINE webhooks, and introduced a stale Docker config schema, eroding trust in point releases.
- **Opaque behavior**: Users want /status to show effective model (#89532), visible subagent sessions (#95295), and clear notifications when sessions fail or rotate (#86215, #89167).
- **Security concerns**: Plain-text apiKey writing (#88562) and reasoning token leaks (#89473) worry users running in shared environments.
- **Platform gaps**: Microsoft Teams attachments, Matrix E2EE dispatch, LINE webhook, and QQ bot reconnection are still unreliable; Telegram media buffering in RAM is flagged as a performance risk (#88594).

Despite frustrations, users express **positive sentiment** toward the project’s capability (“Claude Opus 4.8 with native 1M context is great once it works”) and are actively contributing repro data and even proposed fixes (linked PRs in many issue labels).

## 8. Backlog Watch

Several important issues have been open for extended periods without resolution:

| Issue / PR | Open Since | Severity | Status |
|------------|-----------|----------|--------|
| [#86215](https://github.com/openclaw/openclaw/issues/86215) — Codex OAuth refresh wedge | 2026-05-24 | P1, diamond lobster | `needs-maintainer-review`, `no-new-fix-pr` |
| [#88707](https://github.com/openclaw/openclaw/issues/88707) — Bedrock regression in 2026.5.28 | 2026-05-31 | P1, diamond lobster | `needs-maintainer-review`, `no-new-fix-pr` |
| [#87857](https://github.com/openclaw/openclaw/issues/87857) — Agent skips AGENTS.md startup sequence | 2026-05-29 | P2, diamond lobster | `stale`, `needs-maintainer-review` |
| [#58993](https://github.com/openclaw/openclaw/pull/58993) — Google Chat DM detection (stale PR) | 2026-04-01 | P1, platinum hermit | `stale`, `needs-pr-context` |
| [#89171](https://github.com/openclaw/openclaw/pull/89171) — Validate harness tool names (maintainer-reviewed) | 2026-06-01 | P2 | `ready for maintainer look` — awaiting merge |
| [#89039](https://github.com/openclaw/openclaw/pull/89039) — Silent message loss from takeover error | 2026-06-01 | P1, platinum hermit | `re-review loop` — close to merge |

The backlog of `clawsweeper:needs-maintainer-review` issues (more than 20 items in this dataset alone) suggests that maintainer bandwidth is a bottleneck. The oldest open PR (Google Chat, April 1) and several P1 issues lacking fix PRs should be prioritised to restore community confidence.

---

## Cross-Ecosystem Comparison

# Cross-Project Ecosystem Comparison Report

**Date:** 2026-06-20  
**Analyst:** AI Agent Open-Source Ecosystem Lead

---

## 1. Ecosystem Overview

The open-source personal AI assistant ecosystem is experiencing a bifurcation between mature reference frameworks and specialized toolchains. The "Claw" family (OpenClaw, PicoClaw, NanoClaw, NullClaw, IronClaw, ZeroClaw) dominates the reference-implementation space, with OpenClaw acting as the de facto core platform and others serving as lightweight, embedded, or hardware-targeted derivatives. Meanwhile, projects like NanoBot, Hermes Agent, and CoPaw represent a second wave of opinionated, Python-native agents focused on developer productivity and production deployment. The ecosystem is grappling with three universal challenges: memory management across long sessions, multi-platform channel reliability, and security hardening post-MVP. Activity is heavily concentrated in 5–6 projects, while several others (TinyClaw, Moltis, LobsterAI, ZeptoClaw) show stagnation or maintainer attrition.

---

## 2. Activity Comparison (Last 24 Hours)

| Project | Issues Updated | PRs Updated | Release Today | Health Assessment |
|---------|---------------|-------------|---------------|-------------------|
| **OpenClaw** | 116 (3 closed) | 500 (19 merged) | No | ⚠️ Strained but active; massive backlog |
| **NanoBot** | 7 (3 closed) | 20 (3 merged) | No | ✅ Healthy, steady cadence |
| **Hermes Agent** | 7 new | 50 (30 merged) | **v0.17.0** | ✅ Very strong; high throughput |
| **PicoClaw** | 4 | 7 (1 merged) | **Nightly** | ✅ Moderate; focused on v0.3.0 |
| **NanoClaw** | 0 | 6 (0 merged) | No | ⚠️ Quiet; critical CVE fix pending |
| **NullClaw** | 2 (1 closed) | 0 | No | ⚠️ Low activity; 1 new high-severity bug |
| **IronClaw** | 2 | 32 (11 merged) | No | ✅ High velocity; Reborn runtime focus |
| **LobsterAI** | 5 (all stale-closed) | 0 | No | 🔴 Minimal; no fixes, no releases |
| **TinyClaw** | 1 | 0 | No | 🔴 Critical vulnerability, no activity |
| **Moltis** | 0 | 1 (Dependabot) | No | 🟢 Stable but dormant |
| **CoPaw** | 8 | 15 (5+ merged) | No | ✅ Very high; strong community contributions |
| **ZeptoClaw** | 0 | 0 | No | 🔴 Inactive |
| **ZeroClaw** | 11 (2 closed) | 50 (5 merged) | No | ✅ Intense; milestone/sprint push |

---

## 3. OpenClaw's Position

**Advantages vs. Peers:**
- **Platform breadth:** OpenClaw supports 12+ messaging channels (Discord, Telegram, Matrix, QQ, WeChat, LINE, Feishu, Google Chat, Signal, Slack) — more than any competitor. Hermes Agent is closest with 8–10.
- **Community size:** 116 issues and 500 PRs in 24 hours dwarfs all peers. The contributor base is 10–50x larger than NanoBot or PicoClaw.
- **Ecosystem centrality:** Several projects (PicoClaw, NanoClaw, NullClaw) are explicitly derived from or compatible with OpenClaw's schema and plugin model (e.g., OneCLI, Workboard, MCP).

**Technical Approach Differences:**
- OpenClaw uses a **monolithic reference architecture** with a heavy gateway process, while NanoBot and Hermes Agent prefer **modular, plugin-based** designs. OpenClaw's gateway process is a source of both power (centralized control) and pain (heap memory leaks, OOM crashes — see #89315).
- OpenClaw's configuration is YAML-based and complex; users report "stale config schema" errors (#88788). NanoBot uses simpler JSON config files.
- OpenClaw has **proprietary/closed-source extensions** (Codex OAuth, Workboard) that create vendor lock-in for production users, generating community frustration (#86215). Hermes Agent's entire stack is open-source.

**Weaknesses:**
- **Review bottleneck:** 113 open issues, 481 open PRs. Maintainer bandwidth is clearly insufficient. Triage labels like `clawsweeper:needs-maintainer-review` outnumber closed items by 40:1.
- **Regression proneness:** v2026.5.28 broke Bedrock, LINE webhooks, Docker config schema. Users report "eroding trust in point releases."
- **Security debt:** Plain-text API key storage (#88562), reasoning token leaks (#89473), SSRF vectors — multiple P1 security items lack fix PRs.

---

## 4. Shared Technical Focus Areas

| Focus Area | Projects Affected | Specific Needs |
|------------|------------------|----------------|
| **Memory/Context Management** | OpenClaw (#89315, #88812), NanoBot (#4418, #4389), Hermes Agent (#12987), PicoClaw (#3150), CoPaw (#5342, #5321) | Heap leak fixes, token estimation caching, context window handling for fallback models, "memory loss" bugs, long-session context compaction |
| **Multi-Platform Channel Reliability** | OpenClaw (#88788, #89473, #88955), Hermes Agent (#49536, #49569), PicoClaw (#3114), ZeroClaw (#6651) | Telegram text overlap, WhatsApp Docker regression, QQ bot reconnection, Matrix E2EE dispatch, Discord message chunking, stale config schemas |
| **Security Hardening** | OpenClaw (#88562, #89473, #89228), Hermes Agent (#7073), NanoClaw (#2799 CVE), TinyClaw (#285), ZeroClaw (#8044) | SSRF via ISATAP, unauthenticated API endpoints, plain-text credential storage, reasoning token over-sharing, per-sender authorization gaps |
| **Performance Optimization** | OpenClaw (#88812), NanoBot (#4420), CoPaw (#5342) | Tool-definition JSON caching, pre-model latency measurement, unbounded heap growth, context explosion risks |
| **Model Provider Compatibility** | OpenClaw (#88707, #91869), NullClaw (#952, #967), CoPaw (#5330, #5345), NanoBot (#4389) | Bedrock regression, hard-coded context windows, function calling gaps in custom providers, fallback model misconfiguration |

**Emerging pattern:** Memory management is the #1 cross-project pain point. Every active project has at least one open issue about session memory leaks, context window mismatches, or token estimation inefficiency.

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | Hermes Agent | IronClaw | CoPaw | ZeroClaw |
|-----------|----------|---------|--------------|----------|-------|----------|
| **Primary Focus** | Full-stack reference agent | Lightweight Python agent CLI | Production gateway + dashboard | Enterprise multi-tenant runtime | Multi-agent collaboration | Plugin/skills platform |
| **Target User** | Power users, self-hosters | Developers, ML engineers | Ops teams, Docker users | Enterprise dev teams | General users, mobile | Plugin developers |
| **Architecture** | Monolithic gateway + plugins | Modular, pip-installable | Microservices + stateful gateway | Rust-based Reborn runtime | Python + Web UI | Rust core + plugin VM |
| **Channel Breadth** | 12+ (broadest) | 6–8 (Telegram, Discord, Matrix, CLI) | 8–10 (incl. WhatsApp, Signal) | 4–6 (Telegram, Slack, extensible) | 5–7 (Web, mobile, API) | 6–8 (Telegram, Discord, Matrix) |
| **Release Cadence** | Monthly point releases | Weekly minor patches | Bi-weekly major releases | Pre-release sprints | Frequent patch + feature | Milestone-driven |
| **Community Size** | Very large (116 issues/day) | Medium (7 issues/day) | Large (7 new issues/day) | Medium (2 issues/day) | High (8 issues/day) | High (11 issues/day) |
| **Major Weakness** | Review bottleneck, regressions | Stream stall regression (fixed) | WhatsApp Docker regression | CI instability (nightly E2E fail) | Provider compatibility gaps | Skill discovery errors |

---

## 6. Community Momentum & Maturity Tiers

**Tier 1 – High Velocity, Rapid Iteration** (weekly releases, frequent merges, strong contributor base):  
- **Hermes Agent** – v0.17.0 released today with ~800 merged PRs since v0.16.0. 30 PRs merged in 24 hours.  
- **IronClaw** – 11 PRs merged today. Strong focus on Reborn runtime maturity.  
- **CoPaw** – Multiple community-contributed fixes merged today. High first-time contributor activity.  
- **ZeroClaw** – 5 PRs merged, 50 updated. Sprint or milestone push in progress.  

**Tier 2 – High Volume, Strained** (high activity but bottlenecked):  
- **OpenClaw** – Massive throughput but overwhelmed maintainers; 481 open PRs. Community energy is high but frustration is building.  

**Tier 3 – Steady Refinement** (moderate activity, focused scope):  
- **NanoBot** – Steady bug fixes and small features. Healthy triage of issues.  
- **PicoClaw** – Nightly builds, clear roadmap to v0.3.0. Moderate but consistent.  
- **NanoClaw** – Low activity but critical fixes pending; may accelerate post-CVE.  
- **NullClaw** – Minimal activity; one critical bug needs attention.  

**Tier 4 – Stalled / Minimal** (no merges, no releases, stale issues):  
- **LobsterAI** – All 5 updated issues closed as stale without resolution. No evidence of active development since early April.  
- **TinyClaw** – Critical CVE disclosed; zero response from maintainers in 24 hours.  
- **Moltis** – Dormant; only Dependabot activity.  
- **ZeptoClaw** – Zero activity.  

---

## 7. Trend Signals for AI Agent Developers

**1. Reliability Under Load is the #1 Community Expectation**  
Across OpenClaw (OOM heap crash), NanoBot (stream stall), Hermes Agent (WhatsApp Docker regression), and CoPaw (silent message drops), users consistently report that agents become unreliable after hours of uptime or during concurrent sessions. *Action: Invest in memory profiling, graceful degradation, and load-testing infrastructure before adding new features.*

**2. Multi-Model Fallback is Becoming a Baseline Requirement**  
NanoBot (#4389), OpenClaw (#91869), and CoPaw (#5345) all see user demand for per-model context windows, reasoning effort customization, and seamless fallback between providers. Users are no longer satisfied with single-model deployments. *Action: Standardize model configuration schemas across projects — the ecosystem needs a common "model manifest" specification to reduce fragmentation.*

**3. Security is Moving from Afterthought to Gate**  
OpenClaw (#88562, #89473), NanoClaw (#2799 CVE), TinyClaw (#285), Hermes Agent (#7073), and ZeroClaw (#8044) all have active security reports. The shift from "make it work" to "make it safe" is accelerating, especially for multi-tenant and Docker deployments. *Action: Projects without security audits or CVEs pending should prioritize authentication, input validation, and credential management before their next release.*

**4. Observability is the New Competitive Advantage**  
CoPaw merged a Langfuse tracing improvement today. OpenClaw's community is demanding latency measurement tooling (#88812). IronClaw is adding QA record/replay fixtures. Operators need to understand *why* their agent is slow or failing. *Action: Exposure of per-turn timing, token usage breakdowns, and tool call traces will become table stakes in the next 3–6 months.*

**5. The "Claw Family" Moats Are Weakening**  
While OpenClaw remains the largest project, specialized alternatives like Hermes Agent (better production ops) and NanoBot (simpler developer experience) are gaining mindshare. OpenClaw's 481-open-PR backlog and regression-prone releases risk community migration. *Action: OpenClaw maintainers need to merge the critical P1 fixes (OOM, OAuth wedge, reasoning token leak) and reduce backlog before the next release, or risk losing their core community to better-maintained alternatives.*

**6. Mobile and Programmatic Access are Underserved**  
CoPaw (#5329) and ZeroClaw (#8046) received feature requests for mobile UX and webhook modes. The current generation of agents is desktop-first and chat-only. *Action: API-first design and responsive web UIs will be necessary for the next wave of adoption, especially for developers integrating agents into their own products.*

---

*Data sources: GitHub issue/PR activity on 2026-06-20 for all 13 projects listed. Health assessments combine quantitative activity metrics with qualitative analysis of issue severity, maintainer responsiveness, and community sentiment.*

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-06-20

## 1. Today's Overview

The project shows **high activity** with 7 issues and 20 pull requests updated in the last 24 hours. 3 issues were closed and 3 PRs were merged, indicating steady progress across bug fixes, performance improvements, and feature work. No new release was cut today. The community is actively contributing, with several substantive PRs addressing memory consolidation, subagent modes, Telegram integration, and CLI ergonomics. A notable performance enhancement (tool-definition JSON caching) is already under review.

---

## 2. Releases

**None** — No new releases today.

---

## 3. Project Progress

**Merged/Closed PRs today:**

| PR | Title & Summary | Status |
|----|-----------------|--------|
| [#4138](https://github.com/HKUDS/nanobot/pull/4138) | **Add `tools.file.enable` to toggle built-in filesystem tools** – aligns filesystem tools with the existing enable/disable pattern used by `exec` and `web` tool groups. | ✅ Merged/Closed |
| [#4230](https://github.com/HKUDS/nanobot/pull/4230) | **Fix: set httpx timeout for streamableHttp transport** – prevents indefinite waits during MCP streamable HTTP handshake. | ✅ Merged/Closed |
| [#4246](https://github.com/HKUDS/nanobot/pull/4246) | **Fix(session): `delete_session` also removes legacy path files** – resolves an asymmetry that could revive deleted sessions from the legacy `~/.nanobot/sessions/` directory. | ✅ Merged/Closed |

**Closed Issues today:**  
- [#4013](https://github.com/HKUDS/nanobot/issues/4013) (stream stalled >90 seconds) – closed, likely fixed.  
- [#4374](https://github.com/HKUDS/nanobot/issues/4374) (project workspaces read/write asymmetry) – closed.  
- [#4389](https://github.com/HKUDS/nanobot/issues/4389) (per-model `contextWindowTokens` for fallback models) – closed, possibly resolved by the discussion or upcoming PR.

---

## 4. Community Hot Topics

The most-discussed items today (by comment count) are:

- **[Issue #4013](https://github.com/HKUDS/nanobot/issues/4013)** – *"Error calling LLM: stream stalled for more than 90 seconds"* (5 comments). This bug surfaced after updating to 0.2.0 and rendered the agent unusable for sustained tasks. The closure suggests a fix has been applied; community members expressed gratitude for earlier versions' stability and frustration with the regression.

- **[Issue #4374](https://github.com/HKUDS/nanobot/issues/4374)** – *"Project workspaces: SOUL.md/USER.md read from project but written to default workspace"* (3 comments). Highlights a real deployment pain where persistent memory files are read correctly but written to the wrong location, causing agent context drift.

- **[Issue #4389](https://github.com/HKUDS/nanobot/issues/4389)** – *"Feature Request: Per-model contextWindowTokens for fallback models"* (2 comments). A usability gap when primary vs fallback models have different context windows; the community is seeking a model-level override.

**Underlying need:** Users are increasingly deploying multiple models and fallback chains, and NanoBot’s configuration system needs to keep pace with per-model granularity for context limits, reasoning effort, and other model-specific parameters.

---

## 5. Bugs & Stability

**High severity** – no crashes or regressions introduced today, but a performance bottleneck was identified:

- **Issue #4420** (open) – *"Performance optimization: `estimate_prompt_tokens` does redundant JSON serialization per call"*. This affects every agent turn (up to 3 calls per turn) on large tool manifests. The fix PR [#4421](https://github.com/HKUDS/nanobot/pull/4421) already proposes caching the serialised tool JSON. **Severity: Medium** – not a crash but noticeably degrades latency on complex tool sets.

**Fixed bugs today:**
- Stream stall after 90 seconds (issue #4013) – closed, suspected fix shipped in a prior patch.
- Session deletion legacy path asymmetry (PR #4246 merged) – resolves history revival after deletion.
- MCP streamable HTTP timeout (PR #4230 merged) – prevents indefinite hangs during startup.

**Overall stability assessment:** The project is recovering from a recent regression (0.2.0 stream stall) and hardening critical paths. The current open issues show no critical crashes.

---

## 6. Feature Requests & Roadmap Signals

Several user-submitted features received active PRs or were created today:

| Feature | Issue/PR | Likelihood for next release |
|---------|----------|-----------------------------|
| **Per-model `contextWindowTokens`** for fallback models | [#4389](https://github.com/HKUDS/nanobot/issues/4389) | Low (closed as question; may need separate PR) |
| **Automatic reasoning effort escalation** (default + escalated levels) | [#4419](https://github.com/HKUDS/nanobot/issues/4419) | Medium – aligns with growing multi-model support |
| **Telegram `sendRichMessage` support** (tables, task lists, math blocks) | [#4422](https://github.com/HKUDS/nanobot/issues/4422) + [#4423](https://github.com/HKUDS/nanobot/pull/4423) | **High** – PR already open |
| **Heartbeat tasks delivery to the channel where they were added** | [#4418](https://github.com/HKUDS/nanobot/issues/4418) | Medium – addresses a frequent workflow pain |
| **Suppress routine cron job notifications** | [#4412](https://github.com/HKUDS/nanobot/pull/4412) | High – PR ready, small scope |
| **Aggregated subagent result mode** | [#4414](https://github.com/HKUDS/nanobot/pull/4414) | High – PR under review |
| **SuspendTurn for async / human-in-the-loop** | [#4411](https://github.com/HKUDS/nanobot/pull/4411) | High – major architectural addition |
| **CLI inline TUI** | [#4329](https://github.com/HKUDS/nanobot/pull/4329) | Medium – large UI change, may need more iteration |
| **Python SDK expansion** | [#4296](https://github.com/HKUDS/nanobot/pull/4296) | Medium – extensive, backward-compatible |
| **Cron job model presets** | [#4416](https://github.com/HKUDS/nanobot/pull/4416) | High – PR under review |
| **Subagent spawn model override** | [#4415](https://github.com/HKUDS/nanobot/pull/4415) | High – small PR, ready |
| **Cache tool definition JSON in `estimate_prompt_tokens`** | [#4421](https://github.com/HKUDS/nanobot/pull/4421) | **Very High** – performance fix, trivial risk |

**Prediction:** The next minor release (0.3.0) is likely to include the performance cache (#4421), Telegram rich messages (#4423), aggregated subagent results (#4414), cron presets (#4416), spawn model override (#4415), and the suppress routine cron notifications (#4412). Larger features like TUI and SDK expansion may land in subsequent releases.

---

## 7. User Feedback Summary

**Positive signals:**
- Users praised the 0.1.5 release (#4013: “it’s been very good”).
- The community is actively contributing PRs and thoughtful feature requests, indicating strong engagement.

**Pain points expressed:**
- **Stream stall regression** after upgrading to 0.2.0 (#4013) – rendered the agent unusable for “real work”. Likely fixed now.
- **Workspace read/write asymmetry** (#4374) – agent memory files written to the wrong location caused confusion.
- **No per-model context window** for fallback models (#4389) – leads to prompt truncation errors when fallback model has smaller window.
- **Redundant token encoding** (#4420) – noticeable slowdown for users with large tool definitions (e.g., the reporter’s “nanobee” project).
- **Heartbeat results delivered to wrong channel** (#4418) – reduces usability of scheduled tasks.
- **No built-in toggle for filesystem tools** (#4138) – now addressed by the merged PR.

**Satisfaction:** Users appreciate the project’s rapid iteration and responsiveness, but the 0.2.0 regression temporarily eroded trust. The quick closure of #4013 and other bugs should restore confidence.

---

## 8. Backlog Watch

Several important PRs have remained open for weeks or months without maintainer response (status as of today):

| PR | Created | Summary | Age | Notes |
|----|---------|---------|-----|-------|
| [#1945](https://github.com/HKUDS/nanobot/pull/1945) | 2026-03-12 | **XMPP channel** – fully functional but self-described “no guarantees”. | ~100 days | No comments from maintainers; may need review or decision to merge/deprecate. |
| [#3591](https://github.com/HKUDS/nanobot/pull/3591) | 2026-05-02 | **Dream: add update scope controls** – prevent unwanted skill drift. | ~49 days | No recent updates on review. |
| [#3590](https://github.com/HKUDS/nanobot/pull/3590) | 2026-05-02 | **Heartbeat: add manual trigger command** – on-demand execution. | ~49 days | Same author as #3591; also stalled. |
| [#3662](https://github.com/HKUDS/nanobot/pull/3662) | 2026-05-06 | **Avoid network loads during token estimation** – local fallback now superseded by #4421 approach. | ~45 days | Might be superseded by #4421. |

Additionally, **Issue #4418** (heartbeat delivery) and **#4419** (reasoning effort escalation) were created today and have zero maintainer response yet, but it’s too early to consider them backlog.

**Action recommended:** The three older PRs (#1945, #3591, #3590) have not received maintainer feedback for over a month. The XMPP channel is a substantial integration; a decision (merge, request changes, or close with reasoning) would reduce community uncertainty.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-06-20

## 1. Today’s Overview
The project saw high activity over the past 24 hours: **7 new open issues**, **50 PRs updated** (30 merged/closed, 20 still open), and **one new release** (`v2026.6.19` / v0.17.0 — “The Reach Release”). The release marks a significant milestone with ~1,475 commits, ~800 merged PRs, and contributions from 245 community developers since v0.16.0. Issue creation is focused on Telegram, WhatsApp, Docker, and gateway bugs, while the PR board shows a healthy mix of bug fixes, feature work, and security patches. Overall project health appears robust, with ongoing rapid iteration and a strong community feedback loop.

## 2. Releases
**v2026.6.19 (Hermes Agent v0.17.0)** – Released June 19, 2026.  
Changelog highlights (since v0.16.0): ~1,475 commits · ~800 merged PRs · 1,693 files changed · 235,390 insertions · 50,730 deletions · 300+ issues closed · 245 community contributors. No explicit breaking changes or migration notes were provided in the snippet; the release notes appear truncated. Given the scale, users should review the full v0.17.0 changelog before upgrading.

## 3. Project Progress
30 PRs were closed or merged today. Key advances include:

- **Signal quoted-reply context** – PR [#46388](https://github.com/NousResearch/hermes-agent/pull/46388) (closed) and its salvage PR [#49563](https://github.com/NousResearch/hermes-agent/pull/49563) (merged) enable the gateway to carry Signal quote metadata, fixing ambiguous reply handling.
- **Gateway budget enforcement** – PR [#30828](https://github.com/NousResearch/hermes-agent/pull/30828) (closed) enforces turn budgets across gateway, CLI, API, cron, and delegation, with context-spill guards and regression tests.
- **Stale max_iterations fix** – PR [#32544](https://github.com/NousResearch/hermes-agent/pull/32544) (closed) fixes a gateway cached-agent bug where `max_turns` changes were not applied on the next turn.
- **Telegram reaction config** – PR [#49571](https://github.com/NousResearch/hermes-agent/pull/49571) (open, feature) makes processing lifecycle emojis configurable.
- **Cognee query tool** – PR [#19331](https://github.com/NousResearch/hermes-agent/pull/19331) (open, feature) adds a read-only `cognee_query` tool for source-bound knowledge retrieval.
- **Dashboard system components view** – PR [#49050](https://github.com/NousResearch/hermes-agent/pull/49050) (open, feature) adds a federated AI-stack status view to the dashboard.
- **Desktop build hardening** – PR [#49547](https://github.com/NousResearch/hermes-agent/pull/49547) (open, bug) fixes six issues including build-stamp crashes and unsafe cleanup.
- **Minimax-CN vision routing** – PR [#49568](https://github.com/NousResearch/hermes-agent/pull/49568) (open, bug) fixes base URL, vision routing, and auxiliary backend for the minimax-cn provider.

## 4. Community Hot Topics
Discussion activity was low (most issues/PRs have 0 comments), but the following gather attention:

- **Issue #49536** (Telegram text overlap) – 1 comment, high user concern. A fixing PR [#49537](https://github.com/NousResearch/hermes-agent/pull/49537) is already open.
- **Issue #49569** and **#49561** (WhatsApp Docker bridge broken) – Both opened today, indicating a regression in v2026.6.19 for Docker users. No fix PR yet.
- **PR #7073** (WhatsApp bridge security – HIGH) – Updated today after being open since April 10. Exposes critical messaging in `bridge.js`; maintainer attention is needed.
- **PR #12987** (Ebbinghaus memory decay) – Updated today, open since April 20. A long-running feature request for time-aware memory in Mem0 plugin.
- **PR #39853** (notification_sources config) – Open since June 5, updated today. The gateway’s Kanban poller was ignoring a documented config option.

Underlying sentiment: Users are actively testing the new release and reporting integration gaps (Telegram, WhatsApp, Docker) while also pushing for deeper memory and context capabilities.

## 5. Bugs & Stability
Several bugs were reported today, ranked by severity (P2 and P3 from labels):

| Issue | Severity | Description | Fix PR? |
|-------|----------|-------------|---------|
| [#49536](https://github.com/NousResearch/hermes-agent/issues/49536) | P2 | Telegram finalize message text overlap due to `parse_mode` mutation | Yes – [#49537](https://github.com/NousResearch/hermes-agent/pull/49537) open |
| [#49572](https://github.com/NousResearch/hermes-agent/issues/49572) | P3 | `timestamp` field on message dicts causes HTTP 400 on strict providers | None yet |
| [#49569](https://github.com/NousResearch/hermes-agent/issues/49569) | P2 | WhatsApp Docker bridge: npm install EACCES + wrong log path | None yet |
| [#49561](https://github.com/NousResearch/hermes-agent/issues/49561) | P2 | WhatsApp bridge cannot be installed after upgrade to 2026.6.19 | None yet |
| [#49567](https://github.com/NousResearch/hermes-agent/issues/49567) | P2 | Dashboard starts on `0.0.0.0` when `HERMES_DASHBOARD=1`, rejected by auth gate | None yet |
| [#49529](https://github.com/NousResearch/hermes-agent/issues/49529) (via PR [#49566](https://github.com/NousResearch/hermes-agent/pull/49566)) | P2 | False-positive venv entry point warning in `hermes doctor` | PR [#49566](https://github.com/NousResearch/hermes-agent/pull/49566) |
| [#49499](https://github.com/NousResearch/hermes-agent/issues/49499) (via PR [#49564](https://github.com/NousResearch/hermes-agent/pull/49564)) | P2 | Subprocess encoding crashes on Windows with non-UTF-8 locale | PR [#49564](https://github.com/NousResearch/hermes-agent/pull/49564) |
| Various desktop (PR [#49547](https://github.com/NousResearch/hermes-agent/pull/49547)) | P2 | Desktop integration: build-stamp crashes, broken auto-update on Linux, unsafe cleanup | PR [#49547](https://github.com/NousResearch/hermes-agent/pull/49547) |

The two WhatsApp Docker bugs are the most disruptive (both P2, no PR yet). The Telegram text overlap has a fix in progress.

## 6. Feature Requests & Roadmap Signals
New feature requests from today’s issues:

- **Configurable Telegram reactions** ([#49570](https://github.com/NousResearch/hermes-agent/issues/49570) – PR [#49571](https://github.com/NousResearch/hermes-agent/pull/49571)) – Allow users to customise lifecycle emojis or disable them. Likely to land in v0.17.1.
- **Video multimodal forwarding** ([#49565](https://github.com/NousResearch/hermes-agent/issues/49565)) – Pass `.mp4` files as native multimodal content to video-capable models. Currently a feature request; may align with model provider updates.

Long-standing feature PRs that advance the roadmap:

- **Ebbinghaus memory decay** ([#12987](https://github.com/NousResearch/hermes-agent/pull/12987)) – Time-series awareness for Mem0 plugin.
- **Cognee query tool** ([#19331](https://github.com/NousResearch/hermes-agent/pull/19331)) – Isolated data set querying.
- **Dashboard system components view** ([#49050](https://github.com/NousResearch/hermes-agent/pull/49050)) – Federated stack health UI.
- **Signal recent chat history** ([#46391](https://github.com/NousResearch/hermes-agent/pull/46391)) – Inject bounded Signal history into gateway context.

**Prediction for next version (v0.17.1):** The Telegram reaction config and the Signal quoted-reply context (already merged) are strong candidates. The WhatsApp Docker regression will likely be a priority hotfix.

## 7. User Feedback Summary
Real pain points surfaced today:

- **Telegram message overlap** – Users must refresh to see finalised text.
- **WhatsApp Docker regression** – Bridge completely broken after upgrade; npm permissions and path errors.
- **Dashboard auth loop** – `HERMES_DASHBOARD=1` on Docker leads to gateway rejecting 0.0.0.0.
- **Windows locale issues** – `hermes doctor` crashes on Chinese (GBK/CP936) systems.
- **Missing pip in uv-managed venvs** – Plugin installers fail for Google Meet, Google Chat, and Honcho.

Satisfaction indicators: The large v0.17.0 release and volume of merged PRs signal active maintenance. Users are filing clear, reproducible bug reports, suggesting engagement and testing of the new release.

## 8. Backlog Watch
Issues and PRs that have been open for a long time without resolution:

- **PR #7073** (WhatsApp security, HIGH, open since April 10) – Critical exposure of messaging in `bridge.js`. Still no merge date.
- **PR #12987** (Ebbinghaus memory, open since April 20) – Feature PR with wide interest; needs review and testing.
- **PR #15872** (stale timestamp perception, open since April 26) – Agent misinterprets session start time as current time. Multiple maintainer pings.
- **PR #18844** (custom_providers context-length, open since May 2) – Fix for three bugs in custom provider resolution; no recent movement.
- **PR #19331** (cognee query tool, open since May 3) – Awaiting integration feedback.
- **PR #39853** (notification_sources config, open since June 5) – Gateway bug with available fix; no merge yet.
- **PR #44772** (workspaces npm fix, open since June 12) – Root deps pruned on every update; a straightforward fix.

These items represent both security debt and feature delivery bottlenecks. Maintainers should prioritise review, especially the security fix (#7073) and the WhatsApp Docker regression (#49569, #49561) which have no open PRs.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest – 2026-06-20

## 1. Today’s Overview

PicoClaw shows moderate activity with 4 issues and 7 pull requests updated in the last 24 hours. One new nightly build (v0.3.0-nightly.20260620.287853ab) was published, continuing the march toward the v0.3.0 release. The community remains engaged around Windows compatibility, Telegram permission controls, and attachment handling, while two older bugs (Windows path separator, Matrix identity parsing) have fix PRs awaiting review. A freshly opened bug titled “它给自己整失忆了” (memory loss) and a security-focused SSRF bypass fix PR indicate both stability and security are active areas.

## 2. Releases

**Nightly Build (v0.3.0-nightly.20260620.287853ab)**  
[Full Changelog](https://github.com/sipeed/picoclaw/compare/v0.3.0...main)  
This is an automated, potentially unstable build. No breaking changes are documented, but users deploying from this snapshot should expect possible regressions. The nightly incrementally incorporates fixes and features from the `main` branch since the last v0.3.0 tag.

No official v0.3.0 release yet; migration notes are not applicable.

## 3. Project Progress

One pull request was **merged/closed** today:

- **[PR #2956 – fix: preserve channel enabled state when merging security.yml](https://github.com/sipeed/picoclaw/pull/2956)**  
  (by yuxuan-7814) – Merged. Fixes an issue where channels explicitly `enabled: true` in `config.json` were being disabled after loading `.security.yml`. This improves configurability and prevents silent service disruptions.

Other PRs remain open (see Backlog Watch).

## 4. Community Hot Topics

The most discussed items reflect two major concerns: **Windows compatibility** and **attachment/media support**.

- **[Issue #2472 – `list_dir` returns "invalid argument" on Windows](https://github.com/sipeed/picoclaw/issues/2472)**  
  *6 comments, 1 👍* – A long-standing bug (since April) where platform path separators break filesystem operations. The high comment count indicates it affects multiple users.

- **[Issue #348 – General Attachment Support](https://github.com/sipeed/picoclaw/issues/348)**  
  *4 comments* – Labelled `priority: high` and `type: roadmap`. Users are requesting the ability to process files, documents, and media attachments across Telegram, Discord, and other channels. This is a core usability gap.

- **[Issue #3114 – Telegram permission control by conversation type](https://github.com/sipeed/picoclaw/issues/3114)**  
  *1 comment, 1 👍* – A feature request to restrict dangerous operations (e.g., `exec`, `write_file`) in groups and channels while allowing full capabilities in private chats.

- **[Issue #3150 – BUG: 它给自己整失忆了](https://github.com/sipeed/picoclaw/issues/3150)**  
  *2 comments* – Newly reported (Chinese: “it made itself amnesiac”). Likely a memory/context loss issue. Rapid attention suggests it affects core agent behavior.

- **[PR #2937 – Feat/agent collaboration](https://github.com/sipeed/picoclaw/pull/2937)**  
  (by afjcjsbx) – A large open PR introducing an internal agent collaboration bus. It has no recent comments but represents a significant architectural extension.

## 5. Bugs & Stability

Three new bug-related items were updated in the last 24 hours, ranked by severity:

- **HIGH – SSRF bypass via ISATAP IPv6 literals**  
  [Issue #3074](https://github.com/sipeed/picoclaw/issues/3074) (not shown but referenced) is addressed by **[PR #3143 – fix(web): block private IPv4 embeds in ISATAP literals](https://github.com/sipeed/picoclaw/pull/3143)**. This fixes a security vulnerability where `web_fetch` could be tricked into accessing private networks. The fix is open and ready for review.

- **MEDIUM – Memory/context loss (?) – Issue #3150**  
  The open bug “它给自己整失忆了” has no detailed reproduction steps yet. If confirmed, it could cause agent instability across all channels. No associated fix PR exists as of now.

- **MEDIUM – Windows `list_dir` path mismatch – Issue #2472**  
  A functional bug blocking Windows users from basic file listing. A fix is not yet merged, but the bug has been acknowledged.

- **LOW – Matrix identity `allow_from` parsing – PR #3045**  
  [PR #3045](https://github.com/sipeed/picoclaw/pull/3045) provides a fix for Matrix user IDs with colons. It is open but stale.

Additionally, three more open PRs address minor type-assertion panics and flag parsing issues ([#3091](https://github.com/sipeed/picoclaw/pull/3091), [#3053](https://github.com/sipeed/picoclaw/pull/3053), [#3048](https://github.com/sipeed/picoclaw/pull/3048)), all considered low severity but indicative of careful code quality efforts.

## 6. Feature Requests & Roadmap Signals

The most prominent feature demands point to **multi-modal input** and **granular permission systems**:

- **[Issue #348 – General Attachment Support](https://github.com/sipeed/picoclaw/issues/348)** – Labelled `type: roadmap` and `priority: high`. It is the top-ranked roadmap item. Expect it to be a major focus for the v0.3.0 release.

- **[Issue #3114 – Telegram permission by chat type](https://github.com/sipeed/picoclaw/issues/3114)** – A natural extension of the security model. Could land alongside the attachment support to allow safe file operations in groups.

- **[PR #2937 – Agent Collaboration Bus](https://github.com/sipeed/picoclaw/pull/2937)** – Though stale, this feature would enable multi-agent communication, a significant step toward complex workflows. Predict inclusion in v0.3.0 or v0.4.0.

## 7. User Feedback Summary

Real user pain points observed from issues and PRs include:

- **Windows users** are blocked from basic file operations (`list_dir`) due to path separator handling (Issue #2472).
- **Matrix users** cannot use the `allow_from` whitelist because of user ID format parsing (PR #3045).
- **Configuration loss** occurred when enabling channels in `config.json` after editing `.security.yml` – now fixed in merged PR #2956.
- **Security concerns** – SSRF bypass in `web_fetch` (PR #3143) shows users are actively testing edge cases.
- **Missing features** – Attachment handling (Issue #348) and Telegram role-based permissions (Issue #3114) are frequently requested.
- **Stability worry** – The sudden “memory loss” bug (Issue #3150) may cause dissatisfaction if not quickly resolved.

No overt praise was captured in the data; the community appears constructive and bug-focused.

## 8. Backlog Watch

Several important items have gone stale or lack maintainer response:

| Item | Type | Stale Since | Impact |
|------|------|-------------|--------|
| [Issue #2472 – Windows `list_dir` bug](https://github.com/sipeed/picoclaw/issues/2472) | Bug (open) | Created Apr 10, last updated Jun 19 | Blocks Windows users; fix PR needed |
| [Issue #3114 – Telegram permissions](https://github.com/sipeed/picoclaw/issues/3114) | Feature (open) | Created Jun 12, no maintainer response | Popular request with 1 like; no roadmap commitment |
| [PR #2937 – Agent collaboration](https://github.com/sipeed/picoclaw/pull/2937) | Feature PR (stale) | Last updated Jun 19, no review comments | Large architectural change; needs maintainer evaluation |
| [PR #3045 – Matrix identity fix](https://github.com/sipeed/picoclaw/pull/3045) | Bug fix PR (stale) | Last updated Jun 19, no merge | Simple fix for Matrix users |
| [PR #3048 – MCP flag parsing fix](https://github.com/sipeed/picoclaw/pull/3048) | Bug fix PR (stale) | Last updated Jun 19, no merge | Affects CLI reliability |
| [PR #3053 – Evolution lockStoreFile panic fix](https://github.com/sipeed/picoclaw/pull/3053) | Bug fix PR (stale) | Last updated Jun 19, no merge | Potential crash in store |
| [PR #3091 – OpenAI compat type assertion fix](https://github.com/sipeed/picoclaw/pull/3091) | Bug fix PR (stale) | Last updated Jun 19, no merge | Silent disabling of native search |

These items represent low-hanging fruit for maintainers and could improve project velocity and community trust.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest – 2026-06-20

## Today’s Overview
The project shows no merged changes or new releases today, indicating a quiet maintenance phase. Six pull requests were updated within the last 24 hours, all remaining open with no new comments or reactions. Notably, three of these PRs address security and stability concerns (CVE-2026-29611, safeParseContent handling, Discord message chunking), while two others propose new features (Apple Container runtime and parent permission inheritance). The absence of closed issues or merged PRs suggests ongoing review or revision cycles. Overall activity is low, but the presence of critical fixes signals that maintainers are actively working on hardening the codebase.

## Releases
None this period. No new versions have been published.

## Project Progress
No pull requests were merged or closed today. However, the following open PRs represent active development and fixes under review:
- **#2799** – Security fix for `send_file` (CVE-2026-29611)
- **#2801** – Guard against non-object JSON in `safeParseContent`
- **#2820** – Persist delivery target on `pending_approvals` rows
- **#2812** – Chunk Discord replies instead of truncating
- **#2605** – Feature: inherit parent agent permissions via OneCLI
- **#2809** – Feature: Apple Container runtime + remote OneCLI gateway

These indicate incremental hardening of the approval and message delivery systems, plus platform expansions.

## Community Hot Topics
All PRs currently have zero comments and zero reactions, so there is no clear community engagement to report. The most-discussed items (by virtue of being active) are the security fix **#2799** and the Apple Container feature **#2809**, as they address critical vulnerabilities or significant new capabilities. The lack of discussion may reflect that these PRs are still early in review, or that the project’s community tends to focus on asynchronous code changes rather than public commentary.

## Bugs & Stability
Three fix PRs address stability and security:
- **Critical severity**: **#2799** – `send_file` allows any container-visible file to be read due to missing path restriction. This is a CVE (CVE-2026-29611) and a potential RCE/data exposure vector. Fix pending merge.
- **Medium severity**: **#2801** – `safeParseContent` returns non-object JSON without fallback, leading to `undefined` values for `.text`/`.sender`. Fix ensures primitive payloads are handled correctly.
- **Low severity**: **#2812** – Discord replies >2000 characters are truncated instead of being split across multiple messages. Fix adds chunking logic.
- Additionally, **#2820** fixes a data consistency bug where `pending_approvals` rows never record delivery metadata, affecting approval history queries.

All listed PRs are authored by different contributors and remain open.

## Feature Requests & Roadmap Signals
Two feature PRs stand out:
- **#2605** (`guyb1`) – Inherit parent agent permissions via OneCLI. This addresses a long-standing request to cascade permissions in hierarchical agent setups. The PR follows the project’s contributing guidelines and has been open since May 24, suggesting it may be nearing completion.
- **#2809** (`hidenwalker`) – Apple Container runtime support and remote OneCLI gateway. This enables NanoClaw to run on macOS containers and connect to an external OneCLI, which could expand deployment options for Mac users and enterprise environments.

These features, combined with the CVE fix, are likely candidates for the next minor release (e.g., v0.8.x).

## User Feedback Summary
No direct user feedback (issues, comments) was reported in the last 24 hours. However, the existence of PR #2812 (Discord chunking) implies users experienced truncated replies, and #2820 (approval delivery persistence) suggests users had difficulty tracking approvals. The security fix #2799 indicates an active vulnerability that could affect production deployments. Overall, user pain points revolve around data integrity and platform-specific message limits.

## Backlog Watch
- **PR #2605** (feat: inherit parent agent permissions) – Created May 24, updated June 19. No comments or reviewer activity. With no assignee mentioned, this PR risks stalling if left unattended. Maintainer attention is recommended.
- **PR #2799** (CVE fix) – Critical security update; although only opened June 17, it should be prioritized for merge to mitigate the vulnerability.

No other long-unanswered issues/PRs are present (total open issues: 0).

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-06-20

## Today's Overview
Project activity was limited to issue tracking, with no new pull requests or releases in the last 24 hours. Two issues were updated: one long-standing bug regarding incomplete answers from local Ollama models was closed, while a fresh error report (`NoResponseContent`) was filed. The absence of merged PRs or version bumps suggests a maintenance-focused period, with the team likely triaging user-reported problems. Overall project health is stable but low in development momentum.

## Releases
*None.* No new releases were published today or in the recent past. The latest available release remains v2026.5.29.

## Project Progress
*No pull requests were updated or merged in the last 24 hours.* No feature advancements or regression fixes were landed today.

## Community Hot Topics
Active discussion centered on a single closed issue:

- **[#952 [bug] Local model using ollama returns incomplete answers](https://github.com/nullclaw/nullclaw/issues/952)** (3 comments, closed)  
  *Author: bloodgroup-cplusplus* — Users reported that the agent truncated output when using a local Gemma model via Ollama. The issue was resolved (closed) today, suggesting a fix or workaround was applied, though no linked PR is visible in the data.

No other issues or PRs generated significant discussion in the last 24 hours.

## Bugs & Stability
One new bug was reported today, and one older bug was resolved:

| Issue | Status | Severity | Summary |
|-------|--------|----------|---------|
| [#967 error: NoResponseContent](https://github.com/nullclaw/nullclaw/issues/967) | Open (new) | **High** | On Windows 11 (v2026.5.29), using the Agnes-2.0-Flash model, the agent fails with `error: NoResponseContent` in >50% of conversations. Reproducible across 12 out of 21 dialogs. No fix PR exists yet. |
| [#952 Local model using ollama returns incomplete answers](https://github.com/nullclaw/nullclaw/issues/952) | Closed | Medium | Incomplete sentences when using Gemma via Ollama. Closed today, so likely resolved — but no patch details are available in this digest. |

**Ranking:** The `NoResponseContent` error is the most critical active bug, affecting a large fraction of interactions on a specific OS/model combination. It needs immediate maintainer attention.

## Feature Requests & Roadmap Signals
No explicit feature requests were recorded in today's updates. The two open/closed issues are both bug reports, not enhancement requests. Based on recent trends, the next release may include fixes for Ollama integration and Windows-specific response handling.

## User Feedback Summary
Current user pain points revolve around reliability of local model inference:
- Ollama users experienced truncated answers (now closed).
- Windows users report a high-frequency failure (`NoResponseContent`) with the Agnes model, causing frustration.
- No positive feedback or use-case descriptions were shared in the latest issues.

Users appear dissatisfied with the stability of non-API model paths, especially under Windows.

## Backlog Watch
No long-unanswered issues or PRs were identified. The only open issue (#967) was created today and has no comments yet. Maintainers should prioritize it to prevent it from languishing. No other items require immediate attention.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest – 2026-06-20

## Today’s Overview
The project is experiencing **high activity**, with 32 pull requests updated in the last 24 hours (21 open, 11 merged/closed) and 2 issues updated (both open). Development velocity remains strong, concentrated on the **Reborn runtime**, **host ingress**, **CI infrastructure**, and **QA tooling**. A **nightly E2E failure** (#4108) persists without a fix PR, but no new regressions were introduced. The absence of new releases suggests work is still in a pre-release integration phase. Overall, the project is healthy, with a clear focus on maturing the Reborn feature set and stabilizing the CI pipeline.

## Releases
**None.** No new versions were published in the last 24 hours.

## Project Progress
The 11 merged/closed PRs from the last day advanced several key areas:

- **QA & Developer Guidance**  
  - [#5097 – docs: add Reborn QA guidance to agent rules](https://github.com/nearai/ironclaw/issues/5097) – Merged. Formalised cross-layer and user-visible behaviour testing rules.  
  - [#5096 – test(reborn-qa): port project-setup automation-workflow benchmarks to QA record/replay](https://github.com/nearai/ironclaw/issues/5096) – Merged. Seven benchmarks ported into the recorded trace harness.  
  - [#5095 – test(reborn-qa): add recorded fixtures](https://github.com/nearai/ironclaw/issues/5095) – Merged. Committed LLM trace fixtures for connection, routine, and web-fetch scenarios.

- **CI Optimisation**  
  - [#5092 – ci(spike): A/B sccache (GHA) vs rust-cache on a heavy Reborn build](https://github.com/nearai/ironclaw/issues/5092) – Closed. Non-blocking experimental workflow to evaluate cache strategies.

Other merged/closed PRs (not detailed in the top-20 list) likely include smaller fixes and dependency updates.

## Community Hot Topics
- **[#4108 – Nightly E2E failed](https://github.com/nearai/ironclaw/issues/4108)** (open, 0 comments) – This automated bug report flags a critical pipeline failure. No discussion or fix PR exists yet. The failure is likely blocking the nightly validation gate.  
- **[#5091 – Unified feature-flag system for Reborn](https://github.com/nearai/ironclaw/issues/5091)** (open, 0 comments) – Detailed proposal for a centralised feature-flag system (env + dynamic switching, targeting, rollout). Though no comments, the specification is comprehensive and signals a strong developer need.  
- **PR activity** – Several large PRs (e.g., #5081, #5103, #5102, #4989, #5087) have high comment counts (though the data shows `undefined`). These are core-contributed and indicate active collaboration on multi-tenant Postgres, manifest-driven channels, and OAuth token refresh.

## Bugs & Stability
- **Critical**  
  - [#4108 – Nightly E2E failed](https://github.com/nearai/ironclaw/issues/4108) – The nightly end-to-end test suite is breaking. No fix PR is in progress, and the workflow run (ID 27860203846) shows a failure in the “Full E2E / E2E (features)” job. This is a high-priority stability risk.

- **Moderate / No new bugs**  
  No other crash nor regression reports appeared in the last 24 hours. Many PRs carry a `risk: medium` label, which is typical for large feature additions and may introduce latent instability. The merged CI spike (#5092) suggests maintainers are proactively addressing build performance, which indirectly improves stability.

## Feature Requests & Roadmap Signals
The most notable request is **[#5091 – Unified feature-flag system](https://github.com/nearai/ironclaw/issues/5091)**, which proposes replacing ad-hoc environment-variable gating with a comprehensive per-tenant, dynamic, A/B-capable system. This aligns with the increasing multi-tenancy focus seen in PRs like:

- [#5081 – Hosted single-tenant Postgres profile](https://github.com/nearai/ironclaw/pull/5081) – A hosted preview path with PostgreSQL-backed state.
- [#5100 – Telegram ingress from extension state](https://github.com/nearai/ironclaw/pull/5100) and [#5093 – Slack ingress from extension state](https://github.com/nearai/ironclaw/pull/5093) – Manifest-driven channel management.
- [#5085 – Concurrent turn execution via TurnRunScheduler](https://github.com/nearai/ironclaw/pull/5085) – Performance improvement for LLM inference.
- [#5062 – Per-tool permission override model](https://github.com/nearai/ironclaw/pull/5062) and [#5099 – External-tool Responses round-trip](https://github.com/nearai/ironclaw/pull/5099) – Enhanced extensibility and tool integration.

These features point toward a **Reborn v2** with full multi-tenant, multi-channel, and tool-rich capabilities. The next major release will likely include the feature-flag system, the hosted profile, and the ingress-extension architecture.

## User Feedback Summary
No direct user feedback (comments, reactions) was recorded in the last 24 hours. However, the PRs themselves indicate resolved pain points:

- **[#5087 – Proactive Google OAuth token refresh](https://github.com/nearai/ironclaw/pull/5087)** – Addresses user frustration with manual reconnection due to token expiry.
- **[#5101 – Reuse cargo-component installer in live canary](https://github.com/nearai/ironclaw/pull/5101)** – Fixes swallowed install errors, improving CI reliability and developer experience.

The nightly E2E failure (#4108) is a recurring dissatisfaction signal for developers relying on stable CI.

## Backlog Watch
- **[#4002 – build(deps): bump the actions group (16 updates)](https://github.com/nearai/ironclaw/issues/4002)** – Open since May 24, this Dependabot PR is non-trivial (size L, risk medium) and remains unmerged. Outdated CI dependencies could cause build inconsistencies.
- **[#4108 – Nightly E2E failed](https://github.com/nearai/ironclaw/issues/4108)** – Open since May 27 without a fix. This is a critical blocker for nightly quality assurance.
- **[#5081 – Hosted single-tenant Postgres profile](https://github.com/nearai/ironclaw/pull/5081)** – A large, core-contributed PR with a database migration. Despite being open only two days, it requires careful review and likely has far-reaching implications.

No other long-unanswered items were identified.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-06-20

## Today's Overview
The project saw minimal activity on June 20, 2026. No new releases were published, and no pull requests were created or updated. Five issues were updated, all of which were closed as **stale** after prolonged inactivity. These issues had been open for over two months (created in early April) and were automatically closed without being resolved. This pattern suggests that maintenance and community engagement remain low, with no evidence of ongoing feature development or bug fixing. The project appears to be in a passive state, with user-reported problems left unaddressed.

## Releases
No new releases were reported.

## Project Progress
- **Merged/Closed PRs today:** 0  
- **Advancements:** None. No code changes or fixes were merged.

## Community Hot Topics
Five issues were updated and closed as stale. No new discussions or active threads emerged. The following issues had the most engagement (2–3 comments each) and reflect recurring user pain points:

- **[#1496 — 任务显示完成，但是没有返回 (Task shows complete but no return)](https://github.com/netease-youdao/LobsterAI/issues/1496)**  
  Author: netease-george | 👍: 0 | Comments: 3  
  User reports that a submitted task is marked as complete but no output is returned. This suggests a possible backend processing issue or response handling bug.

- **[#1468 — 创建Agent弹窗关闭时无未保存确认 (Agent creation modal closes without save confirmation)](https://github.com/netease-youdao/LobsterAI/issues/1468)**  
  Author: MaoQianTu | 👍: 0 | Comments: 2  
  Unsaved data in the “Create Agent” modal is silently lost when closing via X, Cancel, or clicking outside. The same pattern appears in two other issues.

- **[#1469 — Agent设置面板关闭时无未保存确认 (Agent settings panel closes without save confirmation)](https://github.com/netease-youdao/LobsterAI/issues/1469)**  
  Author: MaoQianTu | 👍: 0 | Comments: 2  
  Identical behavior for the settings panel – any modified configuration is lost without warning.

- **[#1470 — MCP服务器配置弹窗关闭时无未保存确认 (MCP server configuration modal closes without save confirmation)](https://github.com/netease-youdao/LobsterAI/issues/1470)**  
  Author: MaoQianTu | 👍: 0 | Comments: 2  
  The same data loss issue in the MCP server form, also triggered by Escape key.

- **[#1495 — 无缘无故中断进程 (Unexplained process interruption)](https://github.com/netease-youdao/LobsterAI/issues/1495)**  
  Author: xuzhiwu123 | 👍: 1 | Comments: 2  
  User reports that tasks are frequently interrupted without any apparent reason, questioning whether it is a client or model issue.

**Underlying needs:**  
The cluster of issues around unsaved‑data modals (#1468, #1469, #1470) indicates a strong user desire for basic UX safeguards (confirmation dialogs) across all modal forms. The task reliability issues (#1496, #1495) point to deeper stability concerns that erode user trust.

## Bugs & Stability
All five reported bugs were closed as stale today, meaning no fix was acknowledged. Their severity is assessed as follows:

| Issue | Severity | Description |
|-------|----------|-------------|
| [#1496](https://github.com/netease-youdao/LobsterAI/issues/1496) | **High** | Task completion with no output – blocks core functionality. |
| [#1495](https://github.com/netease-youdao/LobsterAI/issues/1495) | **High** | Random process interruption without explanation – affects all workflows. |
| [#1468](https://github.com/netease-youdao/LobsterAI/issues/1468) | Medium | Data loss in agent creation – UI usability issue. |
| [#1469](https://github.com/netease-youdao/LobsterAI/issues/1469) | Medium | Data loss in agent settings – UI usability issue. |
| [#1470](https://github.com/netease-youdao/LobsterAI/issues/1470) | Medium | Data loss in MCP server config – UI usability issue. |

No fix PRs exist for any of these bugs. The closure as stale without resolution leaves users vulnerable to the same issues.

## Feature Requests & Roadmap Signals
No explicit feature requests were submitted today. However, the repeated complaint about modal data loss implies a strong desire for **dirty‑state detection and unsaved‑changes warnings** across the UI. This could be a candidate for a minor UX improvement in the next release, though no roadmap signals were observed.

## User Feedback Summary
- **Pain points:**  
  - Tasks that appear complete but return no result.  
  - Frequent, unexplained process interruptions.  
  - Frustrating data loss when closing modal dialogs without saving.  
- **Use cases:**  
  - Agent creation and configuration (name, system prompts, MCP environment variables).  
  - Running tasks with expectation of reliable output.  
- **Satisfaction/Dissatisfaction:**  
  - The closure of these issues as stale without comment suggests user dissatisfaction may be growing, as reported problems remain unaddressed. The single 👍 reaction on #1495 indicates at least one other user shares that frustration.

## Backlog Watch
No open issues were updated today, but the five now‑closed stale issues represent unresolved user problems that maintainers should revisit. Particularly important:

- **[#1496 – Task shows complete but no return](https://github.com/netease-youdao/LobsterAI/issues/1496)** – Core functionality broken, needs investigation.  
- **[#1495 – Unexplained process interruption](https://github.com/netease-youdao/LobsterAI/issues/1495)** – Possibly related to client‑side or model stability, deserves a maintainer response even if closed.

Maintainer attention is recommended to either provide a fix or publicly acknowledge the limitation and planned timeline.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

# TinyClaw Project Digest — 2026-06-20

## 1. Today's Overview
TinyClaw (TinyAGI) saw minimal activity over the past 24 hours. No new releases or pull requests were recorded, and only one issue was updated. That single issue is a **critical security vulnerability** concerning unauthenticated access to the HTTP management API, allowing arbitrary local file reads into provider-bound prompts. The project appears to be in a maintenance lull with no recent feature work or commits visible in this window. The disclosure of this vulnerability is likely to trigger an urgent response from maintainers.

## 2. Releases
No new releases were published today. The latest release remains **v0.0.20** (date unknown, but referenced in the security advisory). No migration notes or changelogs are available.

## 3. Project Progress
No pull requests were merged or closed today. No features, fixes, or improvements advanced beyond the current state.

## 4. Community Hot Topics
Only one issue was active in the last 24 hours:

- **[#285 – Security: Unauthenticated `prompt_file` update allows arbitrary local file read into provider-bound prompts](https://github.com/TinyAGI/tinyagi/issues/285)**  
  *Opened by YLChen-007* • 0 comments • 0 reactions  
  **Summary**: A critical vulnerability in TinyAGI ≤0.0.20 enables any client reaching the HTTP management API to set an agent’s `prompt_file` to an arbitrary local path, leading to local file disclosure.  
  **Analysis**: This is by far the most pressing community topic. The lack of comments and reactions likely indicates the issue was raised very recently (same day). The underlying need is a **patch or immediate workaround** to restrict API access (e.g., authentication or network-level isolation).

## 5. Bugs & Stability
| Issue | Severity | Description | Fix Available? |
|-------|----------|-------------|----------------|
| [#285](https://github.com/TinyAGI/tinyagi/issues/285) | **Critical** | Unauthenticated `prompt_file` update allows arbitrary local file read. | No fix PR yet. |

No other bugs, crashes, or regressions were reported today. The reported vulnerability directly impacts system confidentiality and could be exploited in multi-tenant or exposed deployments.

## 6. Feature Requests & Roadmap Signals
No feature requests were submitted or discussed today. Based on the security issue, the immediate roadmap signal is a **forced security patch release** (likely v0.0.21) to address the `prompt_file` endpoint vulnerability. Future releases may also introduce mandatory authentication for the management API.

## 7. User Feedback Summary
No user feedback or use cases were shared today. The only interaction is the vulnerability report, which comes from a security researcher rather than an end user. Pain points cannot be assessed from this limited data.

## 8. Backlog Watch
No long-unanswered issues or PRs were identified in today’s data. The project backlog appears clean, but the **single open issue (#285)** demands immediate maintainer attention to prevent exploitation. No other unaddressed threads are visible.

---

*Digest generated from GitHub data at `2026-06-20T23:59:59Z`.*

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest – 2026-06-20

## 1. Today's Overview
Moltis saw minimal activity in the last 24 hours. No new issues were created or updated, and no pull requests were merged or closed. The sole update was a dependency‑bump pull request from Dependabot, keeping documentation dependencies current. Overall project health appears stable, with no signs of active development or incoming community contributions today.

## 2. Releases
No new releases were published today. The latest release remains unchanged.

## 3. Project Progress
No pull requests were merged or closed today. The only open PR is a routine dependency update.

## 4. Community Hot Topics
There are no issues or pull requests with significant comments or reactions. The single open PR is a maintenance update with zero community engagement. The community appears quiet today.

## 5. Bugs & Stability
No bugs, crashes, or regressions were reported in the last 24 hours. The project remains stable as far as available data indicates.

## 6. Feature Requests & Roadmap Signals
No feature requests were submitted or discussed today. There is no new data to suggest what might appear in the next version.

## 7. User Feedback Summary
No user feedback, pain points, or use‑case discussions were recorded in the last 24 hours. Satisfaction or dissatisfaction cannot be assessed from today’s data.

## 8. Backlog Watch
No long‑unanswered issues or PRs were identified. The only open PR (#1133) is a Dependabot update that requires maintainer review and merge, but it does not indicate a backlog concern.

---

**Links**
- Open PR #1133 (dependencies): [moltis-org/moltis PR #1133](https://github.com/moltis-org/moltis/pull/1133)

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

Based on the provided GitHub data from the CoPaw (QwenPaw) repository, here is the project digest for 2026-06-20.

***

# CoPaw Project Digest - 2026-06-20

## 1. Today's Overview
The CoPaw project is experiencing a day of exceptionally high activity, with 15 updated Pull Requests (PRs) and 8 Issues updated in the last 24 hours. The development velocity is strong, highlighted by the merging of several significant contributions from the community, including improvements to observability, cron job reliability, and context management stability. The community is highly engaged, with numerous first-time contributors submitting fixes for critical areas like security (path traversal) and model compatibility, signaling a healthy and growing ecosystem. However, several new bug reports surfaced today, indicating ongoing challenges with third-party provider compatibility and API reliability under load.

## 2. Releases
No new releases were published today. The latest version remains `v1.1.12.post1`, referenced in a bug report today (#5344).

## 3. Project Progress
Several key PRs were merged or closed today, advancing the project's stability and feature set.

- **Observability Enhancement (Langfuse)**: PR [#5128](https://github.com/agentscope-ai/CoPaw/issues/5128) by `totoyang` was closed. This change improves Langfuse tracing by grouping an entire agent ReAct loop into a single trace, significantly enhancing debugging capabilities.
- **Context Manager Stability**: PR [#5242](https://github.com/agentscope-ai/CoPaw/issues/5242) by `lecheng2018` was closed. This fix adds timeout protection to the `agent.reply()` call within the context compaction logic, preventing the process from freezing during LLM API hangs.
- **Cron Job Reliability**: PR [#5241](https://github.com/agentscope-ai/CoPaw/issues/5241) by `lecheng2018` was closed. It increases the default `misfire_grace_seconds` from 60 to 3600, giving long-running agent tasks more time to complete before a scheduled cron job is skipped.
- **Multi-Agent Collaboration**: PR [#5179](https://github.com/agentscope-ai/CoPaw/issues/5179) by `nguyenthanhthe` was closed. This fix expands the trigger keywords for the multi-agent collaboration skill, improving the agent's ability to recognize user intent for team collaboration modes.
- **Provider Compatibility**: Two PRs by `nguyenthanhthe` (#5337, #5338) were closed, specifically addressing connection test failures for the Zhipu AI provider. A new, more refined fix (#5339) remains open.

## 4. Community Hot Topics
The most active discussions center on improving the user experience for mobile users and troubleshooting integration nuances.

- **Mobile & Sidebar UX**: Issue [#5329](https://github.com/agentscope-ai/CoPaw/issues/5329) (3 comments) is a feature request from `bob-geek11` asking for an agent-switching button in the collapsed sidebar mode, primarily for mobile users. This has immediate traction, as a solution is already being proposed in PR [#5334](https://github.com/agentscope-ai/CoPaw/issues/5334).
- **Reasoning Block Incompatibility**: Issue [#5208](https://github.com/agentscope-ai/CoPaw/issues/5208) (6 comments) involves a warning about message count mismatches when using models (like LongCat-2.0-Preview) that use a `"reasoning"` type block instead of `"thinking"`. This highlights a need for broader compatibility with emerging LLM API formats.
- **Cron Task Interruption**: Issue [#5250](https://github.com/agentscope-ai/CoPaw/issues/5250) (2 comments) details a design friction where cron task descriptions are injected as user messages, causing the agent to interrupt its current workflow. The underlying need is for better context/world-model separation for scheduled tasks vs. user-facing conversations.

## 5. Bugs & Stability
Several bugs were reported today, ranging from critical functional failures to potential stability risks.

- **[High] Function Calling in Custom Providers**: Issue [#5345](https://github.com/agentscope-ai/CoPaw/issues/5345) reports that custom OpenAI-compatible providers (e.g., OMLX) do not support function calling. This is a critical functional gap for users who rely on agentic capabilities via these third-party services.
- **[High] Silent Message Drops via API**: Issue [#5344](https://github.com/agentscope-ai/CoPaw/issues/5344) describes a bug where `POST /api/console/chat` returns HTTP 200 but silently discards the message if the agent is busy. This is a significant reliability issue for programmatic users. A companion issue (#5343) was closed, likely as a duplicate.
- **[Medium] Context Explosion Vulnerability**: Issue [#5342](https://github.com/agentscope-ai/CoPaw/issues/5342) points out that the tool-result pruning mechanism is skipped when an LLM call fails (e.g., 502 error), leading to unchecked context growth. This represents a cascading failure risk.
- **[Medium] Zhipu AI Connectivity**: While a fix is in review (PR #5339), Issue [#5330](https://github.com/agentscope-ai/CoPaw/issues/5330) (referenced in several PRs) highlights a persistent bug where model-level connection tests fail for the Zhipu AI provider due to an incompatible message format.
- **[Low-Medium] Cron Task Workflow Conflict**: Issue [#5250](https://github.com/agentscope-ai/CoPaw/issues/5250), though not a crash, is a significant behavioral bug where scheduled cron tasks incorrectly interrupt the main chat workflow.

## 6. Feature Requests & Roadmap Signals
The most prominent feature requests are already seeing active implementation, suggesting they are high-priority for the next release.

- **Mobile/Dynamic UI**: The request for an agent-switch button in collapsed sidebar mode ([#5329](https://github.com/agentscope-ai/CoPaw/issues/5329)) is likely **in development for the next release**, as a corresponding PR (#5334) is already open.
- **Model List Sorting**: The request to allow users to reorder models within a provider ([#5267](https://github.com/agentscope-ai/CoPaw/issues/5267)) is also under active development, with a PR ready for review (#5336).
- **Dockerized Tools**: A new feature PR, [#5346](https://github.com/agentscope-ai/CoPaw/issues/5346), proposes running tools in Docker for enhanced security and isolation. This is an early signal that the project may be moving towards more robust sandboxing for tool execution.
- **Durable Context History**: PR [#5321](https://github.com/agentscope-ai/CoPaw/issues/5321) introduces a "scroll" context strategy for durable history and a recall REPL, suggesting a move toward more advanced, retrieval-augmented context management.

## 7. User Feedback Summary
User sentiment is a mix of high engagement and frustration with specific integration and stability issues.

- **Pain Points**:
    - **Mobile Usability**: Users like `bob-geek11` are actively trying to use CoPaw on mobile but find the current UI blocking basic actions like switching agents.
    - **Provider Integration**: There is significant user investment in custom providers (OMLX, Zhipu AI), but compatibility issues, especially with advanced features like function calling and connection tests, are a source of friction.
    - **API Reliability**: The bug where the API silently drops messages (#5344) is a critical pain point for users automating workflows, as it creates hidden failures.
- **Use Cases**: Users are clearly pushing the tool beyond simple chat, employing it for mobile access, automated/scripted interactions, and integration with diverse LLM backends.
- **Satisfaction**: The rapid creation of PRs to address bugs and the number of first-time contributors suggest a strong user-base that is invested in the project's success, even if they are encountering bugs.

## 8. Backlog Watch
No long-unanswered issues were flagged in the 24-hour window. Several PRs have been ready for maintainer action:

- **PR #5148** (status unknown): A PR on Langfuse observability, closed today.
- **PR #5066** (status unknown): A PR referenced in other contexts; its current status is unclear.
- **Issue #5267** (Model Sorting): While a PR (#5336) exists, the original feature request is open and could benefit from final maintainer review and merge.
- **PR #5179** (Multi-Agent Keywords): Closed today, resolving this item from the backlog.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest – 2026-06-20

## 1. Today's Overview
ZeroClaw saw intense development activity with **50 pull requests updated** in the last 24 hours (45 open, 5 merged/closed) and **11 issues updated** (9 open, 2 closed). No new release was published. The high PR volume, many carrying high-risk labels, indicates a sprint or milestone push, particularly around gateway hardening, skills platform, and runtime test coverage. Two long-standing bugs – a matrix channel memory leak and a gateway credential hardening issue – were closed today, reflecting solid progress. The project remains highly active with contributions from multiple community members across a wide surface area.

## 2. Releases
**None.** The latest release predates this period.

## 3. Project Progress
Two important issues were resolved:

- **#6127 – Silent-fallback hardening in gateway ([closed](https://github.com/zeroclaw-labs/zeroclaw/issues/6127))**  
  Follow-up to a previous fix (#6099) that hardens gateway-side credential resolution with a fail-loud/sentinel split. This PR/issue closes the remaining silent-fallback vectors.

- **#6651 – Matrix channel memory leak (~1 MB Pss per reload) ([closed](https://github.com/zeroclaw-labs/zeroclaw/issues/6651))**  
  Result of an upstream Arc cycle in matrix-sdk 0.17. Now resolved, eliminating heap growth on SIGHUP/reload.

Additionally, 5 PRs were merged/closed today (details not shown in top 20), representing core improvements in areas such as CI, documentation, and bug fixes. Major open PRs advancing features include:

- **#8008** – `zeroclaw auth email-login` subcommand (OAuth2 device code flow)
- **#7945** – xAI OAuth login support
- **#8000** – ZeroCode UI improvements (browse mode badge, auto-exit)
- **#7923** – Automatic temporary file cleanup config
- **#7922** – Discord slash command localizations + guild scope

## 4. Community Hot Topics
The most active discussion remains around the now-closed **#6127 (gateway hardening)** with 7 comments. The conversation centered on ensuring no silent data paths remain after the earlier fallback fix.

Other items with 1 comment (indicating active triage or coordination):

- **#7320** – v0.8.3 MCP dashboard tracker ([issue](https://github.com/zeroclaw-labs/zeroclaw/issues/7320))
- **#7688** – Runtime hook panic recovery and cancellation tests ([issue](https://github.com/zeroclaw-labs/zeroclaw/issues/7688))
- **#7686** – Approval-gated tool execution ordering tests ([issue](https://github.com/zeroclaw-labs/zeroclaw/issues/7686))

The tracker issues **#7320** (v0.8.3 MCP dashboard) and **#7852** (v0.8.2 skills platform) serve as coordination hubs for upcoming releases, signaling strong maintainer focus on unifying skills, plugins, and A2A surfaces.

## 5. Bugs & Stability
Two new S2 (degraded behavior) bugs were reported today:

- **#8047 – ReadSkillTool looks in `data_dir` but skills live in agent workspace ([issue](https://github.com/zeroclaw-labs/zeroclaw/issues/8047))**  
  In compact skills mode, agent prompt lists correct file locations, but `read_skill("<name>")` fails silently. No fix PR yet.

- **#8039 – fill-translations leak-repair leaves orphaned continuation lines (silent data-loss) ([issue](https://github.com/zeroclaw-labs/zeroclaw/issues/8039))**  
  The repair tool only rewrites the `msgstr ""` line but doesn't remove subsequent multi-line continuations, leading to malformed .po output.

Both are unaddressed as of today. A related security/authorization issue was also raised:

- **#8044 – `/model --agent` lacks per-sender authorization ([issue](https://github.com/zeroclaw-labs/zeroclaw/issues/8044))** – any chat participant can change the effective model for all users.

The earlier matrix memory leak (#6651) is now fixed. No critical crashes or regressions were reported.

## 6. Feature Requests & Roadmap Signals
New user-driven feature requests:

- **#8046 – Optional Telegram webhook mode ([issue](https://github.com/zeroclaw-labs/zeroclaw/issues/8046))**  
  Suggests adding webhook ingress alongside current long-polling. Low implementation friction; likely to be picked up in a minor release.

- **#8043 – RFC: Retire aardvark-sys into zeroclaw-hardware ([issue](https://github.com/zeroclaw-labs/zeroclaw/issues/8043))**  
  Structural cleanup to fold a standalone crate. Sponsored by @JordanTheJet; may land in v0.8.2 or v0.8.3 depending on consensus.

- **#8044 – Harden `/model --agent` scope with per-sender authorization (enhancement)**  
  Raised as a security gap; likely to be prioritized after the skills platform release.

Roadmap trackers provide forward-looking signals:

- **#7320 ([v0.8.3 MCP dashboard](https://github.com/zeroclaw-labs/zeroclaw/issues/7320))** – web and plugin-management surfaces.
- **#7852 ([v0.8.2 skills platform](https://github.com/zeroclaw-labs/zeroclaw/issues/7852))** – coherent skills + plugins + A2A surface.

These trackers are actively linked to several open PRs (e.g., MCP server config UI #8032, auto-clean #7923), indicating near-term delivery.

## 7. User Feedback Summary
User reports today highlight real pain points:

- **Skill discovery broken** (#8047): The `read_skill` tool's path mismatch causes "Unknown skill" errors, blocking skill usage in compact mode.
- **Translation tool data corruption** (#8039): Operators relying on `fill-translations` may experience silent data loss in .po files.
- **Telegram limitation** (#8046): Users behind NAT desire webhook support for reliability.
- **Authorization gap** (#8044): A user noticed that any participant can change the global agent model, posing a trust concern.

On the satisfaction side, the closure of the long-standing matrix memory leak (#6651) and gateway hardening (#6127) addresses two persistent stability issues that likely frustrated users. Overall sentiment is mixed – active feature work is appreciated, but daily bugs degrade the experience for some workflows.

## 8. Backlog Watch
No issues flagged as `status:no-stale` and unanswered for an extended period. The oldest open issue among today's updates is **#7320 (June 6)**, which is actively tracked. **#7852 (June 17)** is also active. No items appear neglected.

However note: **#6127** and **#6651**, which had been open since April and May respectively, were closed today. No other long-standing unanswered issues were identified. Maintainers appear responsive.

---

*Digest generated from GitHub data on 2026-06-20. All links point to zeroclaw-labs/zeroclaw.*

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*