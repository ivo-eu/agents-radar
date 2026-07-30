# OpenClaw Ecosystem Digest 2026-07-30

> Issues: 199 | PRs: 500 | Projects covered: 13 | Generated: 2026-07-30 00:11 UTC

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

# OpenClaw Project Digest — 2026-07-30

## 1. Today’s Overview

OpenClaw is in a period of very high churn: **199 issues and 500 pull requests** were updated in the last 24 hours. With **87 PRs merged/closed** today, the pace of fixes is keeping up with the volume of incoming reports, but the backlog remains heavy — 187 open issues and 413 open PRs. The majority of active issues carry **P1/P2** severity labels and reflect a wave of **regressions, message-loss bugs, and platform-specific crash loops** that surfaced after recent releases. No new official releases were cut today, but the community’s focus is firmly on stability, with many fix PRs in the pipeline. The maintainer queue is under strain, as dozens of high-priority items still await triage or product decisions.

## 2. Releases
No new releases as of 2026-07-30.

## 3. Project Progress
Today’s merged/closed PRs (87 total) show concentrated cleanup in several areas:

- **Gateway & control plane** – Multiple small perf improvements: `fix(gateway-client): reset outer event sequence on reconnect` ([#116132](https://github.com/openclaw/openclaw/pull/116132)), `perf(gateway): cache lazy server methods import` ([#116060](https://github.com/openclaw/openclaw/pull/116060)), and `fix: reconcile ambiguous Control UI updates` ([#116098](https://github.com/openclaw/openclaw/pull/116098)).
- **Channel delivery fixes** – `fix(discord): keep activity receipts in adopted threads` ([#116119](https://github.com/openclaw/openclaw/pull/116119)) and `fix(feishu): preserve card actions through mention gating` ([#116105](https://github.com/openclaw/openclaw/pull/116105)) address message-loss scenarios in group chats.
- **Context engine & tool support** – `improve(codex): reduce long-history prompt allocations` ([#115048](https://github.com/openclaw/openclaw/pull/115048)), `fix(google-vertex): add Google Vertex AI onboarding wizard, fix ADC auth` ([#87800](https://github.com/openclaw/openclaw/pull/87800)) – still open but waiting maintainer review.
- **Backup & upgrades** – `fix(doctor): stop advisory state-dir skips from wedging gateway startup` ([#114678](https://github.com/openclaw/openclaw/pull/114678)) resolves a critical boot-blocker on Linux.
- **Feature work** – `feat(msteams): support multiple bot accounts` ([#112811](https://github.com/openclaw/openclaw/pull/112811)) and `feat(queue): persist followup queues across gateway restarts` ([#82572](https://github.com/openclaw/openclaw/pull/82572)) are still open but show progress.

A notable trend is the regular appearance of **bot-automated PRs** (Dependabot, ClawSweeper) – 6 of the top-30 PRs by comment count are automated, suggesting the CI/CD pipeline is active but may be generating noise.

## 4. Community Hot Topics

### Most Active Issues (by comment count)
| Issue | Comments | Summary |
|-------|----------|---------|
| [#88657](https://github.com/openclaw/openclaw/issues/88657) 🐚 | 10 | DeepSeek V4 Flash incomplete turn regression (OpenRouter) |
| [#89278](https://github.com/openclaw/openclaw/issues/89278) 🐚 | 9 | Codex OAuth refresh succeeds but cron/heartbeat fails with 10s timeout |
| [#112423](https://github.com/openclaw/openclaw/issues/112423) 🦪 | 9 | Large SQLite transcript cleanup blocks gateway event loop |
| [#87756](https://github.com/openclaw/openclaw/issues/87756) 🐚 | 8 | Prompt-launched Lobster workflow hangs on nested /tools/invoke |
| [#92433](https://github.com/openclaw/openclaw/issues/92433) 🐚 | 8 | Subagent completion silently dropped when announce steers into requester run |
| [#8299](https://github.com/openclaw/openclaw/issues/8299) 🌊 | 7 | Feature: config option to suppress sub-agent announce (long-running, Feb 2026) |
| [#88154](https://github.com/openclaw/openclaw/issues/88154) 🌊 | 7 | Feature: Slack Modal Support for Interactive Workflows |
| [#89315](https://github.com/openclaw/openclaw/issues/89315) 🦪 | 6 | Gateway heap grows unbounded, killed by cgroup OOM on long-running Linux |

**Analysis:** The community is vocal about **reliability regressions** after the `.27`/`.28` releases. Many users are hitting the same pain points: incomplete AI model responses, lost messages across channels (WhatsApp, Discord, Telegram), and session-state corruption. The volume of comments on feature requests (#8299, #88154) shows demand for better control over agent behaviour, especially around **sub-agent lifecycle** and **UI interactivity**.

### Notable PRs (by type/labels)
- [#90836](https://github.com/openclaw/openclaw/pull/90836) – `fix(cron): block self-narrating auto-announce replies` (merge-risk: 🚨 message-delivery) – extremely critical for cron reliability.
- [#114678](https://github.com/openclaw/openclaw/pull/114678) – `fix(doctor): stop advisory state-dir skips from wedging gateway startup` – addresses a **boot loop** on Linux after upgrade.
- [#112811](https://github.com/openclaw/openclaw/pull/112811) – `feat(msteams): support multiple bot accounts` – long-awaited multi-tenant Teams support.

## 5. Bugs & Stability

Priority P0 and P1 bugs dominate today’s reports:

| Issue | Severity | Impact | Fix PR exists? |
|-------|----------|--------|----------------|
| [#112962](https://github.com/openclaw/openclaw/issues/112962) | **P0** | `gateway.bind` config normalises all bind modes to localhost, causing perpetual crash-restart loop | No open PR |
| [#87928](https://github.com/openclaw/openclaw/issues/87928) | **P0** | macOS update leaves stale node host → Gateway restart storm (blocker) | No PR linked |
| [#89315](https://github.com/openclaw/openclaw/issues/89315) | **P1** | Gateway heap grows unbounded → OOM killed on long-running Linux deployments | No PR linked |
| [#89473](https://github.com/openclaw/openclaw/issues/89473) | **P1** | Reasoning tokens leak to chat channels (Discord etc.) when models stream interleaved text/thinking blocks | No PR linked |
| [#87756](https://github.com/openclaw/openclaw/issues/87756) | **P1** | Prompt-launched Lobster workflow hangs on nested /tools/invoke | No PR linked |
| [#112698](https://github.com/openclaw/openclaw/issues/112698) | **P1** | Codex app-server notification path starves Gateway main thread ~22s (quadratic snapshots) | No PR linked |
| [#116010](https://github.com/openclaw/openclaw/issues/116010) | **P2** | All persistent sessions capped at 128k context regardless of model *[new today]* | No PR linked |
| [#105528](https://github.com/openclaw/openclaw/issues/105528) | P2 | exec/read tools silently return empty output on Windows (regression) | No PR linked |

**Critical regressions** from `2026.5.27→2026.5.28` continue to plague users: DeepSeek V4 Flash incomplete turns ([#88657](https://github.com/openclaw/openclaw/issues/88657)), Bedrock provider registration broken ([#88707](https://github.com/openclaw/openclaw/issues/88707)), and webhook `action: "wake"` silently ignoring `agentId/sessionKey` ([#64556](https://github.com/openclaw/openclaw/issues/64556)). A cluster of **systemd and upgrade-related bugs** (#79375, #87928, #112962) indicate the 2026.7.x series introduced deployment pain for Linux and macOS users.

Today’s fresh bugs include a **128k context cap regression** ([#116010](https://github.com/openclaw/openclaw/issues/116010)) that affects all persistent sessions regardless of model, and a codex notification path that can lock the gateway for ~22 seconds ([#112698](https://github.com/openclaw/openclaw/issues/112698)).

## 6. Feature Requests & Roadmap Signals

The most upvoted and discussed feature requests this week:

- **LTS Version** ([#87295](https://github.com/openclaw/openclaw/issues/87295), 4 👍) – Users want a stable, long-term supported release for critical deployments. This is the top-voted enhancement.
- **Slack Modal Support** ([#88154](https://github.com/openclaw/openclaw/issues/88154), 7 comments) – Would enable structured form input in Slack; requested since May.
- **Memory Therapy** ([#105494](https://github.com/openclaw/openclaw/issues/105494)) – Interactive loop to resolve memory-wiki open questions with the user; no PR yet.
- **Suppress Sub-Agent Announce** ([#8299](https://github.com/openclaw/openclaw/issues/8299), oldest active enhancement, Feb 2026) – Persistent demand for configurable sub-agent behaviour.
- **Inherit Requester Session Model for Subagents** ([#89522](https://github.com/openclaw/openclaw/issues/89522)) – Allows subagents to use the parent’s active model.
- **Source-Aware Instruction Tracking** ([#87714](https://github.com/openclaw/openclaw/issues/87714)) – Architectural mitigation for indirect prompt injection.

**Prediction for next version:** Given the high engagement and maintainer attention, expect at minimum **sub-agent announce suppression** and **LTS release discussion** to appear in the next minor release. Slack modal support is complex but has clear community support. The **memory therapy** feature is early-stage but aligns with the project’s memory-wiki direction.

## 7. User Feedback Summary

Real user pain points from today’s data:

- **“My bot misses replies in group chats”** – WhatsApp reply fence ([#92186](https://github.com/openclaw/openclaw/issues/92186)), Feishu card callbacks dropped ([#116105](https://github.com/openclaw/openclaw/issues/116105)), Telegram duplicate sub-sessions ([#112342](https://github.com/openclaw/openclaw/issues/112342)).
- **“Upgrade broke my setup”** – Multiple users report that upgrading from 2026.5.26→2026.5.27/28 caused incomplete model responses, broken Bedrock auth, and systemd conflicts.
- **“My agent loses context after a few turns”** – Persistent 128k context cap ([#116010](https://github.com/openclaw/openclaw/issues/116010)) and Windows exec/read empty output ([#105528](https://github.com/openclaw/openclaw/issues/105528)) frustrate day-to-day use.
- **“Gateway keeps crashing on long runs”** – Heap OOM ([#89315](https://github.com/openclaw/openclaw/issues/89315)) and SQLite transcript blocking ([#112423](https://github.com/openclaw/openclaw/issues/112423)) hit users with heavy workloads.
- **“I want more control over agent behaviour”** – The sub-agent announce suppression feature ([#8299](https://github.com/openclaw/openclaw/issues/8299)) and ability to see model after reset ([#89274](https://github.com/openclaw/openclaw/issues/89274)) show users want less magic and more configurability.

Satisfaction signals are harder to find, but the fact that many users invest time in detailed bug reports (reproduction steps, logs) indicates strong engagement and willingness to see the project succeed.

## 8. Backlog Watch

The following issues and PRs have been waiting for maintainer attention for weeks or months, with no fix PR in sight:

**Issues (stale, high priority)**
- [#8299](https://github.com/openclaw/openclaw/issues/8299) – Config option to suppress sub-agent announce (created Feb 2026, needs product decision)
- [#69086](https://github.com/openclaw/openclaw/issues/69086) – `attempt-execution` session-history guard too broad (created Apr 2026, P2)
- [#87756](https://github.com/openclaw/openclaw/issues/87756) – Lobster workflow hang (created May 2026, P1, still open)
- [#88657](https://github.com/openclaw/openclaw/issues/88657) – DeepSeek V4 Flash incomplete turn (created May 2026, no fix PR)
- [#88087](https://github.com/openclaw/openclaw/issues/88087) – Poor UX for long-running background tasks (created May 2026, user abandoning droplet)

**PRs (waiting for review)**
- [#82572](https://github.com/openclaw/openclaw/pull/82572) – Persist followup queues across gateway restarts (created May 2026, status: needs proof)
- [#87764](https://github.com/openclaw/openclaw/pull/87764) – Support owner-scoped ClawHub skill refs (created May 2026, needs proof)
- [#87800](https://github.com/openclaw/openclaw/pull/87800) – Google Vertex AI onboarding (created May 2026, ready for maintainer look)
- [#90603](https://github.com/openclaw/openclaw/pull/90603) – Use configured default agent ID in auth/model path discovery (created Jun 2026, ready for maintainer look)
- [#97175](https://github.com/openclaw/openclaw/pull/97175) – Bound deferred turn maintenance with per-task timeout (created Jun 2026, needs proof)

**Risk:** The combination of **17 issues tagged "needs-maintainer-review"** and **35 PRs with status "ready for maintainer look" or "needs proof"** suggests the review pipeline is the bottleneck. Several P1 regressions (#88657, #87756, #89315) have no fix PR at all, which could delay a stable release.

---

*Generated from OpenClaw public GitHub data for 2026-07-30. Data snapshot: 199 issues, 500 PRs updated in last 24h.*

---

## Cross-Ecosystem Comparison

# Cross-Project Ecosystem Comparison Report — 2026-07-30

## 1. Ecosystem Overview

The open-source personal AI assistant ecosystem remains in a high-velocity stabilization phase, where the largest projects (OpenClaw, IronClaw, CoPaw, ZeroClaw) are simultaneously dealing with regression hangovers from recent releases while pushing forward major architectural features like inter-agent communication (A2A), attested signing, and declarative skill systems. A second tier of projects (NanoBot, Hermes Agent, LobsterAI, Moltis) is focusing on targeted UX hardening and infrastructure maturity, with smaller projects (PicoClaw, NanoClaw, NullClaw) making incremental gains on specific integrations. A common pattern across all active projects is the tension between rapid feature iteration and the need for production-grade reliability—users are increasingly vocal about regressions that break basic workflows. The ecosystem is consolidating around a core set of shared challenges: message delivery guarantees, context management across sessions, secure tool execution, and provider diversity.

---

## 2. Activity Comparison

| Project | Open Issues | Updated Issues (24h) | Open PRs | Updated PRs (24h) | PRs Merged/Closed (24h) | Release Status | Health Score |
|---|---|---|---|---|---|---|---|
| **OpenClaw** | 187 | 199 | 413 | 500 | 87 | No release | ⚠️ High churn, heavy regression load |
| **IronClaw** | *12* | 12 | *50* | 50 | 14 | No release (RC imminent) | ✅ Strong, rapid issue resolution |
| **CoPaw** | 19 | 19 | 36 | 50 | 14 | No release | ✅ Fast-moving, healthy contribution |
| **ZeroClaw** | 7 touched | 7 | *50* | 50 | 2 | No release (v0.8.4 due Jul 31) | ✅ Active, security-focused |
| **NanoBot** | 3 open | 5 | 15 open | 33 | 18 | No release | ✅ Very responsive maintainers |
| **Hermes Agent** | *12* | 12 | *50* | 50 | 14 | No release | ✅ Stabilizing with fast fixes |
| **LobsterAI** | 0 updated | 0 | 2 open | *15* | 13 | No release (2026.7.24 released) | ✅ Polished, low bug count |
| **Moltis** | 0 | 0 | 4 | 5 | 1 | No release | ✅ Clean backlog, steady |
| **NanoClaw** | 2 | 2 | 3 | 7 | 4 | No release | ⚠️ Moderate, quiet |
| **NullClaw** | 1 | 1 | 2 | 4 | 2 | No release | ⚠️ Low activity, targeted fixes |
| **PicoClaw** | 1 | 1 | 2 | 2 | 0 | No release | 🔴 Stale, no maintainer interaction |
| **TinyClaw** | — | 0 | — | 0 | 0 | No release | 🔴 Inactive |
| **ZeptoClaw** | — | 0 | — | 0 | 0 | No release | 🔴 Inactive |

*Note: Issue/PR counts marked with * are estimated from the "50" cap in the digests; actual totals may be higher.*

---

## 3. OpenClaw's Position

**Advantages vs. peers:**
- Largest community and contributor base by a wide margin (199 issues, 500 PRs updated in 24h)—dwarfing even the next most active projects (IronClaw, ZeroClaw at ~50 PRs each).
- Most mature ecosystem: the reference implementation, with the broadest channel support (Discord, Feishu, WhatsApp, Telegram, Microsoft Teams) and the deepest integration feature set.
- The "LTS release" discussion (#87295) signals that OpenClaw is the default choice for production deployments where stability is paramount.

**Technical approach differences:**
- OpenClaw uses a **centralized Gateway architecture** with a heavy control plane, while projects like NanoBot and LobsterAI are more lightweight. This gives OpenClaw more power but also more surface area for regressions (heap OOM, event loop blocking, SQLite contention).
- OpenClaw's **Codex context engine** and memory-wiki system are more sophisticated than most peers, but also a source of complex bugs (prompt allocation, OAuth timing, subagent completion drops).

**Community size comparison:**
- OpenClaw's community engagement is an order of magnitude larger than any peer. The next most active projects (IronClaw, ZeroClaw, Hermes Agent) each move about 50 PRs/day—OpenClaw moves 500.
- However, OpenClaw's backlog is also proportionally heavier: 187 open issues and 413 open PRs, compared to zero open issues for Moltis and LobsterAI.

**Risk:** OpenClaw's velocity is outpacing its maintainer capacity. The "needs-maintainer-review" and "ready for maintainer look" queues are growing, and several P1/P0 regressions lack fix PRs.

---

## 4. Shared Technical Focus Areas

The following requirements are emerging across multiple projects simultaneously, indicating ecosystem-wide pain points:

| Requirement | Projects Affected | Specific Needs |
|---|---|---|
| **Message delivery reliability** | OpenClaw (#88657, #116105, #92186), CoPaw (#6524), NullClaw (#915), IronClaw (#6805) | At-least-once delivery, reconnection guarantees, no silent drops in group chats |
| **Context/window management** | OpenClaw (#116010 — 128k cap), CoPaw (#6541 — compression errors), Hermes Agent (#74462 — cold start), ZeroClaw (#9278 — compression defaults) | Configurable context size, session persistence, predictable compression |
| **Provider diversity & fallback** | NanoClaw (#1350 — Copilot SDK), ZeroClaw (#8965 — declarative skill provider), IronClaw (#3057 — dual-engine quota fallback), NullClaw (#981 — Grok CLI) | Multi-provider routing, automatic quota overflow, unified abstraction |
| **Sub-agent lifecycle control** | OpenClaw (#8299 — suppress announce), NanoBot (#5000 — multi-agent collaboration), CoPaw (#6475 — notice_after_complete) | Configurable sub-agent behavior, parent-child model inheritance, structured delegation |
| **Security hardening** | ZeroClaw (#9384, #9401, #9433 — sandbox/symlink/tool allowlist), CoPaw (#6487 — path exfiltration), IronClaw (#6813–6822 — attested signing), Moltis (#1170 — operator gating) | Privilege separation, sandboxing, policy enforcement, cryptographic signing |
| **Observability & feedback** | Moltis (#1174 — instrumentation, Langfuse), IronClaw (#6524 — hermetic testing), Hermes Agent (#74460 — voice confirmation), NanoBot (#5162 — delivery status) | Telemetry, user feedback loops, traceability, testing automation |
| **Desktop/mobile UX** | Hermes Agent (#69551 — SSH profiles), LobsterAI (#2405 — selected text tags), CoPaw (#6424 — desktop GUI automation) | Native integrations, progressive web apps, cross-platform consistency |

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | IronClaw | ZeroClaw | CoPaw | NanoBot | Hermes Agent | LobsterAI | Moltis |
|---|---|---|---|---|---|---|---|---|
| **Target user** | Power users, self-hosters, enterprise | Security-critical deployments, multi-chain | Developer-friendly, declarative ops | Desktop-first, Chinese market | Lightweight, Python-first | Desktop agent, voice & SSH | Electron desktop, IM integration | PWA, multi-channel agent |
| **Architecture** | Centralized Gateway + Codex engine | Process-based, attested signing, Railway | Modular, A2A protocol, declarative skills | Desktop app + MCP server | Python modules, strict typing | Node.js, web-based TUI | Electron, React, IM platforms | Go-based, PWA-first |
| **Key differentiator** | Broadcast ecosystem, reference impl | Cryptographic security, bug bash discipline | Security hardening, documentation gaps | Desktop automation, Chinese community | Very responsive maintenance | SSH remote mode, voice | Cowork features, Electron polish | Zero open issues, clean backlog |
| **Release cadence** | Continuous churn (no LTS yet) | RC imminent (common 0.5.0, skills 0.4.0) | v0.8.4/0.8.5 dual track | 2.0.1 current, 2.1.0 likely | Head-only, no release | v0.19.0 latest | 2026.7.24 just released | No new releases |
| **Community health** | High engagement, maintainer bottleneck | Strong, structured bug bashes | Active, many PRs need author action | High contribution, regressions irritate | Very responsive, fast turnarounds | Active, decision bottlenecks | Low bug count, automated deps stale | Sole contributor, clean |

---

## 6. Community Momentum & Maturity

**Tier 1 — Critical momentum, high churn, heavy regression load:**
- **OpenClaw** — Ecosystem leader by volume, but the regression wave from .27/.28 releases is eroding user trust. The gap between feature velocity and stability is widening. A "stabilization sprint" is overdue.

**Tier 2 — High iteration, rapid issue resolution, growing maturity:**
- **IronClaw** — Best-in-class bug bash discipline (all P1s closed within 24-48h). The attested signing stack and hermetic testing epic signal a project moving toward enterprise readiness.
- **ZeroClaw** — Dual release tracks (maintenance + weekly non-breaking) show operational maturity. Heavy security focus. The A2A PR indicates architectural ambition.
- **NanoBot** — Most responsive maintenance among mid-tier projects. Media path and PowerShell bugs closed within 3 days. Strict typing enforcement improves code quality.
- **CoPaw** — High contribution volume (14 PRs merged/day). Desktop GUI automation and checkpoint system are major features. But regressions (skill tags, MCP reconnect) annoy users.

**Tier 3 — Steady build, low bug count, polished:**
- **LobsterAI** — Zero open issues, 13 PRs merged (all fixes). The reverted run-safety feature shows discipline: they backed out a problematic feature rather than shipping broken. The cleanest backlog in the ecosystem.
- **Hermes Agent** — Stabilizing after v0.19.0. Fast turnaround on MCP shutdown bug (report-to-fix in 3 weeks). SSH profiles and voice gateway are maturing.
- **Moltis** — Zero open issues, four open PRs all from the same maintainer. Clean backlog but sole contributor risk.

**Tier 4 — Low activity, stable features, limited contribution:**
- **NanoClaw** — Moderate activity, dual-engine fallback is a notable feature. Provider diversity demand (#1350) is the biggest gap.
- **NullClaw** — Targeted fixes (scheduler token persistence, Grok provider). Low volume but responsive on the one active bug.

**Tier 5 — Stale/Inactive:**
- **PicoClaw** — No maintainer interaction on open PRs (both stale since March/July). Single new bug report with zero response.
- **TinyClaw, ZeptoClaw** — No activity in 24h. Effectively dormant.

---

## 7. Trend Signals

### From user feedback (value for AI agent developers):

1. **"Upgrade broke my setup" is the #1 pain point** — Across OpenClaw, CoPaw, and Hermes Agent, users report that minor releases introduce regressions that break core workflows (message delivery, authentication, context management). The ecosystem needs better regression testing and LTS release tracks. Developers building on these frameworks should pin versions and test upgrades in staging.

2. **Context management is the next frontier** — The 128k context cap regression (OpenClaw), compression errors on DeepSeek (CoPaw), and cold start latency (Hermes Agent) all point to the same challenge: AI agents need intelligent, configurable context windows that respect model limits while retaining essential history. The ecosystem is moving from "infinite context" promises to pragmatic, user-configurable policies.

3. **Multi-agent orchestration is emerging as a core requirement** — OpenClaw's sub-agent announce suppression, NanoBot's multi-agent proposal, CoPaw's notice_after_complete, and ZeroClaw's goal tools all indicate that developers want structured delegation with lifecycle control. The ad-hoc sub-agent model is being replaced by designed multi-agent architectures.

4. **Provider diversity is non-negotiable** — Users across projects (NanoClaw #1350, NullClaw #981, IronClaw #3057, ZeroClaw #8965) demand the ability to switch between LLM providers, use multiple backends, and fall back on quota exhaustion. Single-provider lock-in is increasingly unacceptable for production use.

5. **Voice and mobile are growing but immature** — Telegram voice typing indicators (Hermes Agent), voice confirmation mode (Hermes Agent #74460), PWA notifications (Moltis #1173), and microphone silence errors (NanoBot #5165) show that voice/mobile UX is still rough. This is a differentiator opportunity for projects that invest in it.

6. **Security hardening is accelerating** — ZeroClaw's three concurrent sandbox/allowlist fixes, CoPaw's path exfiltration patch, IronClaw's attested signing campaign, and Moltis's operator gating all signal that the ecosystem is moving from "works on my machine" to "production-hardened." AI agent developers should expect more tool-call restrictions and identity verification layers.

7. **Observability is becoming table stakes** — Moltis's Langfuse instrumentation, IronClaw's hermetic testing epic, NanoBot's delivery status tracking, and Hermes Agent's patch parser improvements all show that developers need insight into agent behavior. The agents that succeed will be those that make their reasoning and actions auditable.

### Bottom line for developers:
- **If you need stability today**, pin OpenClaw pre-.27 or use LobsterAI (zero open bugs) or Moltis (clean backlog).
- **If you need cutting-edge features**, watch OpenClaw for architecture innovation, IronClaw for security, and ZeroClaw for A2A.
- **If you are building production agents**, invest in multi-provider fallback, context management policies, and observability—these will be the differentiators in 12 months.
- **If you are contributing**, the most impactful areas are: regression testing infrastructure, documentation (especially for Telegram and MCP), and sub-agent lifecycle controls.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-07-30

Based on the 24-hour activity window ending 2026-07-30, data sourced from [github.com/HKUDS/nanobot](https://github.com/HKUDS/nanobot).

---

## 1. Today’s Overview

The project saw an exceptionally high burst of activity, with **33 pull requests updated** (18 merged/closed, 15 open) and **5 issues updated** (3 open, 2 closed). The majority of effort was concentrated on **stability and hardening**: critical bugs such as media‑path loss during session consolidation (Issue [#5118](https://github.com/HKUDS/nanobot/issues/5118)) and UTF‑8 corruption under Windows PowerShell 5.1 (Issue [#5159](https://github.com/HKUDS/nanobot/issues/5159)) were resolved. Infrastructure also advanced — the codebase was brought under **BasedPyright strict type checking** (PR [#5158](https://github.com/HKUDS/nanobot/pull/5158)), and new WebUI features (skill marketplaces, message delivery tracking) reached `main`. No new releases were cut today.

---

## 2. Releases

**None.** The last published release remains the previous version; today’s changes are head‑only.

---

## 3. Project Progress

The following pull requests were merged or closed today, reflecting both new features and regressions fixes:

- **feat(webui): add skill marketplaces and management** (PR [#5116](https://github.com/HKUDS/nanobot/pull/5116)) — Adds a Discover view for browsing and installing third‑party skills from skills.sh and SkillHub, with instant filtering and install‑history sparklines.
- **feat(webui): track optimistic message delivery status** (PR [#5162](https://github.com/HKUDS/nanobot/pull/5162)) — Implements `sending → accepted → failed` lifecycle for user messages, with inline failure UI and gateway error details on hover.
- **fix(shell): preserve UTF‑8 native input on PowerShell 5** (PR [#5160](https://github.com/HKUDS/nanobot/pull/5160)) — Fixes `$OutputEncoding` misconfiguration that caused non‑ASCII characters to be corrupted when using Windows PowerShell 5.1.
- **fix(memory): expose media references to session consolidation** (PR [#5157](https://github.com/HKUDS/nanobot/pull/5157)) — Resolves the root cause of Issue [#5118](https://github.com/HKUDS/nanobot/issues/5118) by sharing the media‑breadcrumb renderer between live replay and consolidation; uploaded file paths are no longer silently dropped.
- **refactor: enforce BasedPyright strict type checking** (PR [#5158](https://github.com/HKUDS/nanobot/pull/5158)) — All 273 Python modules under `nanobot/` are now strict‑clean, with 31 file‑level suppressions that will be incrementally removed.

Additionally, several critical‑regression fixes remain **open** but have active PRs: session lock leaks (PR [#5151](https://github.com/HKUDS/nanobot/pull/5151)), unbounded exec buffered output (PR [#5150](https://github.com/HKUDS/nanobot/pull/5150)), and subagent partial completion (PR [#5152](https://github.com/HKUDS/nanobot/pull/5152)).

---

## 4. Community Hot Topics

- **Issue #5000**: [Proposal: evolve the current subagent system toward multi‑agent collaboration](https://github.com/HKUDS/nanobot/issues/5000)  
  *Comments: 6 | 👍: 0*  
  The most‑discussed open issue. The author argues that the current subagent model is more “background task delegation” than true multi‑agent collaboration, lacking persistent identity and shared state. The discussion has drawn input from several contributors, indicating strong interest in redesigning the agent‑orchestration architecture.

- **PR #5034**: [feat(goal): add durable state‑graph planning and recovery](https://github.com/HKUDS/nanobot/pull/5034)  
  A long‑running enhancement PR that complements Issue #5000 by introducing structured execution plans and recovery paths for `/goal` sessions. Although it has not yet been merged, its duration (open since July 22) and the number of dependency labels suggest it is a high‑priority roadmap item.

- **PR #5116 (closed)**: Skill marketplaces were a highly anticipated feature, and its closure today signals that the ecosystem is now open for third‑party skill discovery.

---

## 5. Bugs & Stability

| Bug (Issue) | Severity | Status | Fix PR |
|-------------|----------|--------|--------|
| Session consolidation drops media paths (Issue [#5118](https://github.com/HKUDS/nanobot/issues/5118)) | **Critical** – files become unrecoverable after archiving | Closed | [#5157](https://github.com/HKUDS/nanobot/pull/5157) |
| PowerShell 5.1 corrupts non‑ASCII input (Issue [#5159](https://github.com/HKUDS/nanobot/issues/5159)) | **High** – breaks all non‑English pipelines on legacy Windows | Closed | [#5160](https://github.com/HKUDS/nanobot/pull/5160) |
| Manual cron runs lose completion state on WebUI polling (Issue [#5163](https://github.com/HKUDS/nanobot/issues/5163)) | **Medium** – state inconsistency, but execution still succeeds | Open | None yet |
| Subagent partial‑completion model inference (Related PR [#5152](https://github.com/HKUDS/nanobot/pull/5152)) | **Medium** – model may assume unfinished sub‑tasks are complete | Open | [#5152](https://github.com/HKUDS/nanobot/pull/5152) |
| Idle session lock leaks (PR [#5151](https://github.com/HKUDS/nanobot/pull/5151)) | **Low‑Medium** – eventual memory pressure under long‑running sessions | Open | [#5151](https://github.com/HKUDS/nanobot/pull/5151) |
| WebUI false microphone silence errors (PR [#5165](https://github.com/HKUDS/nanobot/pull/5165)) | **Low** – UI noise, audio transcription still works | Open | [#5165](https://github.com/HKUDS/nanobot/pull/5165) |
| Malformed token‑usage day keys crash settings API (PR [#5146](https://github.com/HKUDS/nanobot/pull/5146)) | **Medium** – blocks `/api/settings` when one key is corrupt | Open | [#5146](https://github.com/HKUDS/nanobot/pull/5146) |

The team has been **highly responsive**, with dedicated fix PRs for every reported regression except the cron race (Issue [#5163](https://github.com/HKUDS/nanobot/issues/5163)), which remains unassigned.

---

## 6. Feature Requests & Roadmap Signals

- **Multi‑agent collaboration** (Issue [#5000](https://github.com/HKUDS/nanobot/issues/5000)) — The proposal has gained momentum and is likely to be advanced in the next minor release, as it directly enables complex workflows and task decomposition.
- **Durable state‑graph goal planning** (PR [#5034](https://github.com/HKUDS/nanobot/pull/5034)) — Already in implementation stage; if merged, it would provide structured recovery for long‑running `/goal` tasks, addressing a common pain point.
- **Custom Telegram Bot API base URL** (PR [#4919](https://github.com/HKUDS/nanobot/pull/4919)) — Despite being open since July 14, it continues to receive updates. Enterprise users needing self‑hosted Bot API servers should expect this in the next release.
- **Skill marketplaces** (PR [#5116](https://github.com/HKUDS/nanobot/pull/5116)) — Now merged, this feature will allow users to discover and install community skills directly from the WebUI, opening the door for an ecosystem of plugins.

---

## 7. User Feedback Summary

### Pain points addressed today:
- **Media path loss** — Users reported that uploaded files became unrecoverable after archiving. The fix (PR [#5157](https://github.com/HKUDS/nanobot/pull/5157)) directly resolves this, and the community expressed appreciation in the issue thread.
- **PowerShell 5.1 encoding** — Non‑ASCII input (e.g., CJK characters) was silently corrupted. The fix (PR [#5160](https://github.com/HKUDS/nanobot/pull/5160)) was contributed by a user who encountered the problem on legacy Windows setups.
- **WebUI mis‑silence** — Users relying on microphone input saw false “silence” errors that prevented transcription. PR [#5165](https://github.com/HKUDS/nanobot/pull/5165) addresses the root cause (Web Audio analyser reporting flat waveform on non‑silent audio).

### Satisfaction signals:
- The quick turnaround on the media‑path and PowerShell bugs (both opened and closed within 3 days) indicates a **responsive maintainer team**.
- The addition of **optimistic message delivery status** (PR [#5162](https://github.com/HKUDS/nanobot/pull/5162)) and **skill marketplaces** (PR [#5116](https://github.com/HKUDS/nanobot/pull/5116)) were met with positive reactions in chat.

### Persistent dissatisfaction:
- The **cron‑state race condition** (Issue [#5163](https://github.com/HKUDS/nanobot/issues/5163)) is still open with no assigned fix. Users may experience confusion when a manually triggered cron job succeeds but the WebUI still shows “Failed”.
- The **multi‑agent proposal** (Issue [#5000](https://github.com/HKUDS/nanobot/issues/5000)) highlights a structural limitation that some users consider a blocker for production‑grade automations.

---

## 8. Backlog Watch

The following items have been open for an extended period and may require maintainer attention:

- **Issue #4812**: [fix(memory): use .get() for role key to prevent KeyError](https://github.com/HKUDS/nanobot/pull/4812)  
  *Opened: July 6 | Status: Open with conflicts*  
  A defensive fix for malformed history entries, but conflicts with recent memory refactoring have stalled it. A rebase would likely be straightforward.

- **PR #4919**: [feat(telegram): support custom Bot API base URL and extra headers](https://github.com/HKUDS/nanobot/pull/4919)  
  *Opened: July 14 | Status: Open, conflict*  
  This feature is high‑demand for enterprise deployments, but it has been blocked by merge conflicts for over two weeks. It has received several updates, indicating the author is active, but maintainer review is needed to unblock.

- **Issue #5000**: [Proposal: evolve subagent system](https://github.com/HKUDS/nanobot/issues/5000)  
  *Opened: July 20 | Status: Open, discussion ongoing*  
  While not “unanswered,” the proposal has not yet received an official roadmap commitment from maintainers. Community members are awaiting guidance on whether to contribute a design document.

- **PR #5094**: [fix(providers): use canonical OpenRouter app URL](https://github.com/HKUDS/nanobot/pull/5094)  
  *Opened: July 26 | Status: Open, conflict*  
  A low‑risk attribution improvement; stale conflicts may have prevented merge. Should be trivial to resolve.

---

*Generated from GitHub activity data for the 24‑hour period ending 2026-07-30.*

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-07-30

## 1. Today's Overview

Hermes Agent is experiencing a high-activity day with 50 pull requests and 12 issues updated in the last 24 hours, reflecting strong community engagement and ongoing stabilization efforts. The project closed 14 PRs (including six merged/closed in the top-20 set) and saw one issue resolved. No new releases were cut; the latest tagged version remains v2026.7.20 (v0.19.0). Key areas of focus include desktop SSH profile isolation, MCP server lifecycle fixes, configuration dual-storage inconsistencies, and several file-tool validation improvements. The pulse suggests the maintainers are prioritizing bug-squashing ahead of a potential v0.19.1 release.

## 2. Releases

No new releases today. The latest release remains **v2026.7.20** (package version 0.19.0). Note: issue #74448 flags a divergence between the tag and `__version__` strings, which may be worth tracking.

## 3. Project Progress

**Merged/closed PRs (top-20 subset):**

| PR | Title | Outcome |
|----|-------|---------|
| [#74459](https://github.com/nousresearch/hermes-agent/pull/74459) | fmt(js): auto-fix | Merged |
| [#74420](https://github.com/nousresearch/hermes-agent/pull/74420) | fix(managed_uv): repair vulnerable SQLite runtime in .venv installs too | Merged |
| [#74454](https://github.com/nousresearch/hermes-agent/pull/74454) | fmt(js): auto-fix | Merged |
| [#72054](https://github.com/nousresearch/hermes-agent/pull/72054) | fix(mcp): reap orphaned parked server task in _connect_server | Merged |
| [#74446](https://github.com/nousresearch/hermes-agent/pull/74446) | feat(pairing): profile-correct approvals, and a desktop surface to do them from | Merged |

**Notable feature advances in open PRs (still under review):**
- [#27040](https://github.com/nousresearch/hermes-agent/pull/27040) — Generic `voice_server` gateway platform for external voice runtimes (WebSocket protocol).
- [#65982](https://github.com/nousresearch/hermes-agent/pull/65982) — `claude-agent-sdk` as a first-class Hermes runtime under subscription OAuth.
- [#73861](https://github.com/nousresearch/hermes-agent/pull/73861) — French locale for Hermes Desktop.

**Bug fixes merged today:**
- SQLite WAL-reset corruption in `.venv` installs (PR #74420)
- Orphaned MCP server tasks on connection errors (PR #72054)
- Desktop pairing approvals now profile-correct with a GUI surface (PR #74446)

## 4. Community Hot Topics

**Most active issues (by comment count):**

- [#71298](https://github.com/nousresearch/hermes-agent/issues/71298) (13 comments) — **Bug: providers vs custom_providers dual storage causes CLI/GUI mismatch + model version stuck in profile**. Users report confusion when `hermes setup model` and Desktop GUI show different provider configurations. The underlying need is for a single source of truth in `config.yaml`.

- [#69551](https://github.com/nousresearch/hermes-agent/issues/69551) (10 comments) — **Desktop SSH remote mode broken with non-default profile**. Token-path validation resolves against profile-scoped `HERMES_HOME` while client hardcodes `~/.hermes/desktop-ssh`. Users working with multiple profiles cannot use SSH remote mode.

- [#60197](https://github.com/nousresearch/hermes-agent/issues/60197) (8 comments) — **[CLOSED] RuntimeError: Event loop is closed during /exit (MCPServerTask.shutdown)**. Multiple MCP server tasks fire ignored exceptions when the agent exits. The fix was merged via PR #72054 today.

- [#18659](https://github.com/nousresearch/hermes-agent/issues/18659) (5 comments) — **scan_skill_commands unconditionally clears _skill_commands before try block**. All 90+ skill slash commands silently lost on scan failure. This has been open since May and still needs a maintainer decision.

- [#69663](https://github.com/nousresearch/hermes-agent/issues/69663) (3 comments) — **Desktop stuck "An update is finishing…" after successful self-update**. Updater process never exits, boot gate blocks forever. A high-priority stability concern for desktop users.

**Most active PRs (by discussion, though comment counts are undefined):**
- [#74455](https://github.com/nousresearch/hermes-agent/pull/74455) — Fix for composer chips demoting to plaintext (serious UX regression).
- [#27040](https://github.com/nousresearch/hermes-agent/pull/27040) — Voice server gateway (waiting for decision).
- [#65982](https://github.com/nousresearch/hermes-agent/pull/65982) — claude-agent-sdk provider (large feature, needs-decision).

## 5. Bugs & Stability

**New bugs reported today (2026-07-30):**

| Issue | Severity | Description | Fix PR Exists? |
|-------|----------|-------------|----------------|
| [#74456](https://github.com/nousresearch/hermes-agent/issues/74456) | P2 | Hermes Agent installation broken in Termux (curl pipe bash fails) | No |
| [#74448](https://github.com/nousresearch/hermes-agent/issues/74448) | P2 | `search_files` silently falls through to content-search for invalid `target` enum values | Yes — [#74457](https://github.com/nousresearch/hermes-agent/pull/74457) open |
| [#74462](https://github.com/nousresearch/hermes-agent/issues/74462) | (unlabeled, ~P2) | Cold start latency >16 seconds for first chat in session | No |
| [#74460](https://github.com/nousresearch/hermes-agent/issues/74460) | (feature request) | Voice confirmation mode — STT transcribe → user confirms → agent receives | No |

**Critical regressions or lingering high-severity bugs updated today:**
- [#71298](https://github.com/nousresearch/hermes-agent/issues/71298) (P2, needs-repro) — providers vs custom_providers dual storage mismatch.
- [#69551](https://github.com/nousresearch/hermes-agent/issues/69551) (P2, needs-decision) — Desktop SSH broken with non-default profile.
- [#69663](https://github.com/nousresearch/hermes-agent/issues/69663) (P2, needs-repro) — Desktop update stuck loop.
- [#70637](https://github.com/nousresearch/hermes-agent/issues/70637) (P2, needs-repro) — Telegram typing indicator persists when agent is idle.

**Bugs fixed today:**
- MCP orphaned task leak ([#60197](https://github.com/nousresearch/hermes-agent/issues/60197) → PR #72054 merged).
- SQLite vulnerability in `.venv` installs (PR #74420 merged).
- Patch parser V4A boundary marker false positives (PR #74458 open fix).

## 6. Feature Requests & Roadmap Signals

**New feature request today:**
- [#74460](https://github.com/nousresearch/hermes-agent/issues/74460) — **Voice Confirmation Mode**: STT transcription presented for user confirmation before sending to agent. Especially requested by Telegram/Discord voice message users with accents or noisy environments.

**Previously open features with recent activity:**
- [#27040](https://github.com/nousresearch/hermes-agent/pull/27040) — Generic voice server gateway (open PR since May 16, labelled `needs-decision`). Likely candidate for next minor release if maintainers approve.
- [#65982](https://github.com/nousresearch/hermes-agent/pull/65982) — Claude Agent SDK provider (open PR since July 16, heavy label set including `needs-decision`). Represents a major new runtime option.
- [#73861](https://github.com/nousresearch/hermes-agent/pull/73861) — French locale for Desktop (open PR, no strong opposition seen).
- [#57295](https://github.com/nousresearch/hermes-agent/issues/57295) (P3) — Desktop updater should notify user when local changes are stashed/applied/conflicted. Unanswered since July 2, but relevant to the update-stuck bug.

**Roadmap prediction:** The next version (v0.19.1 or v0.20.0) is likely to include:
- Fixes for search_files validation, patch parser, MCP shutdown
- pairing approval UI (already merged today)
- Potential inclusion of the voice gateway PR if it passes review – the number of companion issues (Telegram typing indicator) suggests voice is a growing use case.

## 7. User Feedback Summary

**Pain points expressed today:**
- **Cold start latency**: One user reports 16 seconds for the first message in TUI mode – far above the expected 2 seconds. This suggests session initialization or model loading overhead is excessive for some setups.
- **Installation failure on Termux**: The `curl | bash` installer fails on Termux (Android) with no obvious workaround. Likely due to unsupported platform detection or missing dependencies.
- **Desktop dual-storage confusion**: Users trying to manage providers via both CLI and GUI experience inconsistent states and stuck model versions. The root cause is `providers` vs `custom_providers` being displayed differently.
- **SSH remote mode profile mismatch**: Power users with multiple profiles cannot use SSH remote desktop at all – a blocker for enterprise adoption.
- **Update stuck indefinitely**: Even after successful binary update, the desktop remains stuck on a "finishing" screen, requiring manual intervention.
- **Composer chips demoting to plaintext**: A UI regression that frustrates users who rely on slash commands (e.g., `/work @folder`).

**Use cases highlighted:**
- Voice messaging on Telegram/Discord with confirmation before agent processing.
- Multi-user sessions in group chats (PR #63089 addresses turn isolation).
- Kanban worker terminal environment isolation (PR #73117).
- Cross-platform installation (Termux, macOS/Linux desktop).

**Satisfaction signals:** High community activity (50 PRs, 12 issues updated) indicates an engaged user base. The fast turnaround on the MCP shutdown bug (from report to merged fix in 3 weeks) shows maintainers are responsive to crashes.

## 8. Backlog Watch

**Long-unanswered issues needing maintainer attention:**

| Issue | Age | Status | Why it matters |
|-------|-----|--------|----------------|
| [#18659](https://github.com/nousresearch/hermes-agent/issues/18659) (scan_skill_commands bug) | Created 2026-05-02 (3 months), last updated today | needs-decision | All skill commands silently lost on scan failure – a major UX degradation that has been open for months. Likely a simple fix but requires maintainer to prioritize. |
| [#38359](https://github.com/nousresearch/hermes-agent/issues/38359) (TUI dark theme diff colors) | Created 2026-06-03 (2 months), last updated 2026-07-29 | P3, no decision | Visual inconsistency that diminishes the polish of the TUI experience. |
| [#57295](https://github.com/nousresearch/hermes-agent/issues/57295) (Desktop updater notifications) | Created 2026-07-02 (4 weeks) | P3, zero comments | Related to the update-stuck bug; user-facing feedback on local changes is still missing. |

**Open PRs needing maintainer decision:**
- [#27040](https://github.com/nousresearch/hermes-agent/pull/27040) – voice server gateway (since May 16, labelled `needs-decision`). Stalled for over 2 months.
- [#65982](https://github.com/nousresearch/hermes-agent/pull/65982) – claude-agent-sdk provider (since July 16, labelled `needs-decision`). Large feature with 4 related PRs.
- [#63089](https://github.com/nousresearch/hermes-agent/pull/63089) – telegram busy group turn isolation (since July 12, labelled `needs-decision`).
- [#63263](https://github.com/nousresearch/hermes-agent/pull/63263) – Gemini content-policy mapping (since July 12, labelled `needs-decision`).
- [#63258](https://github.com/nousresearch/hermes-agent/pull/63258) – Cron DST transition fix (since July 12, labelled `needs-decision`).

These long-pending items risk community frustration if not addressed soon, especially #18659 which impacts all skill command users.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest – 2026-07-30

## 1. Today's Overview
Project activity remains low, with only **1 issue** and **2 pull requests** updated in the last 24 hours, and no merged or closed items. No new releases have been published. The single new issue (#3301) reports a notable bug where `/clear` and session auto-compression fail when chats are routed to a non-default agent via dispatch rules. Two older PRs received updates (#3283, #1951) but remain open, indicating a backlog that may require maintainer attention. Overall, the project appears to be in a quiet maintenance phase with community engagement limited to bug reporting and incremental feature work.

## 2. Releases
No new releases as of 2026-07-30.

## 3. Project Progress
No pull requests were merged or closed today. Two PRs were updated:
- **#3283** – [fix(dingtalk): support picture/image message inbound](https://github.com/sipeed/picoclaw/pull/3283) (by MrTreasure, created Jul 22, updated Jul 29). This is a feature enhancement for the DingTalk channel, adding image message handling. Despite being labeled “stale,” it received an update yesterday.
- **#1951** – [chore: move installation scripts from docs repo to here](https://github.com/sipeed/picoclaw/pull/1951) (by lc6464, created Mar 24, updated Jul 29). A long‑standing chore PR that brings scripts into the main repository. Its update today may signal renewed interest.

No functional fixes or feature merges occurred.

## 4. Community Hot Topics
The only active discussion revolves around the newly filed bug report:
- **#3301** – [[BUG] /clear and session auto-compression don't work in chats routed to non-default agent via dispatch rules](https://github.com/sipeed/picoclaw/issues/3301)  
  *Author: j-v | 0 comments | 0 reactions*  
  **Analysis:** The issue describes a clear user pain point: multi‑agent dispatch rules break core chat management functions. Although it has no comments yet, its specificity (affecting Discord/Telegram channels, Raspberry Pi, DeepSeek model) suggests a configuration‑level regression. The lack of community discussion may indicate the bug is either niche or newly discovered.

Both open PRs (#3283, #1951) have no comments, so no active debate exists.

## 5. Bugs & Stability
**High severity:**
- **#3301** – `/clear` and session auto‑compression fail under non‑default dispatch rules.  
  **Impact:** Affects multi‑agent workflows on popular channels (Discord, Telegram). Users cannot clear conversation history or rely on automatic compression, likely leading to memory/resource issues over time.  
  **Status:** No associated fix PR yet; the bug is newly open and unreviewed.

No other crashes or regressions were reported in the last 24h.

## 6. Feature Requests & Roadmap Signals
No explicit feature requests were filed today. However, the updated PRs hint at future directions:
- **DingTalk picture support (#3283)** – If merged, this would complete inbound image handling for the DingTalk channel, an enhancement for users of that platform.
- **Installation script consolidation (#1951)** – Moving scripts into the main repo simplifies deployment and is a common quality‑of‑life improvement that often precedes a release.

These changes are likely candidates for inclusion in the next version (e.g. 0.3.2), assuming they are reviewed and merged.

## 7. User Feedback Summary
Only one piece of user feedback exists from the reported bug:
> **Pain point:** Configuration of custom dispatch rules breaks basic session management – `/clear` and auto‑compression become non‑functional.  
> **Use case:** Users running PicoClaw on Raspberry Pi with multiple agents (triggered by dispatch rules) cannot maintain clean chat histories.  
> **Satisfaction:** Implicitly negative – the issue was filed without a workaround and no immediate resolution.

No praise or other feedback was recorded in the last 24 hours.

## 8. Backlog Watch
Two important items remain unresolved and may require maintainer attention:

- **#3283** – [fix(dingtalk): support picture/image message inbound](https://github.com/sipeed/picoclaw/pull/3283)  
  *Open since Jul 22, labeled “stale” – no maintainer review or comments.*  
  This PR adds meaningful functionality but risks being abandoned. A decision (review/merge/close) would clarify its status.

- **#1951** – [chore: move installation scripts from docs repo to here](https://github.com/sipeed/picoclaw/pull/1951)  
  *Open since Mar 24 – updated yesterday but still unmerged.*  
  Being a chore with a long lifespan, it may block other workflow improvements or cause confusion if scripts exist in both locations.

Both PRs have zero maintainer interaction, which could signal a bottleneck in project governance. Acknowledgment or triage would help the community understand the project’s direction.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest – 2026-07-30

## 1. Today's Overview

NanoClaw shows moderate activity with 2 open issues and 7 pull requests updated in the last 24 hours, of which 4 have been merged or closed. No new releases were cut today. The team is focusing on container infrastructure hardening, bug fixes in polling and Slack integration, and advancing the large dual-engine quota fallback feature. A new Telegram bug (empty rich messages) and a long-standing feature request for GitHub Copilot SDK support remain open, signaling ongoing community and maintainer attention.

## 2. Releases

_No new releases today._

## 3. Project Progress

**Merged / Closed PRs (4):**

- **[#3150 [CLOSED] setup: fetch a hardened agent image instead of building it**](https://github.com/nanocoai/nanoclaw/pull/3150) – Adds an alternative method to pull a pre‑built, hardened agent image from the NanoClaw registry (provided by Echo.ai). Local building remains the default and only account‑free path. This operational improvement reduces setup time and improves security.

- **[#2440 [CLOSED] fix(poll-loop) + feat(agent): session routing fix and pre-compaction notification**](https://github.com/nanocoai/nanoclaw/pull/2440) – Fixes session routing when the container restarts with pending inbound messages, ensuring the correct reply channel is used. Also adds pre‑compaction notifications for agent sessions.

- **[#2904 [CLOSED] fix(slack): reload thread history from platform on @mention**](https://github.com/nanocoai/nanoclaw/pull/2904) – Ensures that when a bot is re‑tagged in an existing Slack thread, the full thread history (including all human messages since the last bot reply) is fetched, eliminating message blindness.

- **[#3060 [CLOSED] fix(container): add --init to agent container spawn args so PID 1 reaps zombie processes**](https://github.com/nanocoai/nanoclaw/pull/3060) – Adds `--init` flag to container spawn to prevent zombie process accumulation, correcting a documented gap.

**Open PRs (3):**  
- [#3145 fix(db): backfill destinations for existing wirings](https://github.com/nanocoai/nanoclaw/pull/3145) – Migration to add missing channel destinations.  
- [#3149 fix(cli): add --rw flag to groups config add-mount](https://github.com/nanocoai/nanoclaw/pull/3149) – CLI usability fix.  
- [#3057 Dual-engine quota fallback: Claude→Codex overflow, handoff recaps, proactive quota warning](https://github.com/nanocoai/nanoclaw/pull/3057) – Large feature branch for automatic fallback between providers, already tested in production.

## 4. Community Hot Topics

- **[Issue #1350 – Add GitHub Copilot SDK as alternative AI backend](https://github.com/nanocoai/nanoclaw/issues/1350)**  
  _8 👍, 3 comments_ – This open feature request has strong community backing. The user wants to use Copilot models (e.g., GPT‑4.1) alongside or instead of the Anthropic Claude Agent SDK. The discussion highlights a desire for provider diversity and cost flexibility. The issue has been open since March 2026, indicating sustained interest but slow progress.

- **[PR #3057 – Dual-engine quota fallback: Claude→Codex overflow](https://github.com/nanocoai/nanoclaw/pull/3057)**  
  Author: elia-ben-cnaan, updated today – This comprehensive feature branch introduces automatic fallback from Claude to Codex on quota exhaustion, with handoff recaps and proactive warnings. Already battle‑tested on a live WhatsApp deployment. The PR has likely attracted significant review attention (no comment count provided). It represents a major roadmap item.

- **[Issue #3151 – Telegram: Bot API 10.1 rich_message inbound arrives empty](https://github.com/nanocoai/nanoclaw/issues/3151)**  
  New, no comments yet – A recently reported bug that silently drops formatted content from Telegram. Although not yet hot, its severity (data loss) will likely drive community attention.

## 5. Bugs & Stability

| Severity | Issue/PR | Description | Status |
|----------|----------|-------------|--------|
| **High** | [#3151 (open)](https://github.com/nanocoai/nanoclaw/issues/3151) | Telegram `rich_message` content silently dropped – messages with formatted text arrive empty, no error. Root cause appears to be Bot API 10.1 changes. | No fix PR yet |
| **Medium** | [#3145 (open PR)](https://github.com/nanocoai/nanoclaw/pull/3145) | Missing channel destinations for existing messaging‑group wirings – migration 021 backfills them. | Fix under review |
| **Low** | [#3149 (open PR)](https://github.com/nanocoai/nanoclaw/pull/3149) | CLI `groups config add-mount` missing `--rw` flag – usability bug. | Fix open |
| **Fixed** | [#2904 (closed)](https://github.com/nanocoai/nanoclaw/pull/2904) | Slack thread history invisible on re‑@mention – fixed. | Merged |
| **Fixed** | [#3060 (closed)](https://github.com/nanocoai/nanoclaw/pull/3060) | Zombie process accumulation in agents – fixed with `--init`. | Merged |
| **Fixed** | [#2440 (closed)](https://github.com/nanocoai/nanoclaw/pull/2440) | Session routing bug after container restart – fixed. | Merged |

The Telegram bug (#3151) is the most critical new stability issue, as it causes silent data loss without error logs.

## 6. Feature Requests & Roadmap Signals

- **GitHub Copilot SDK backend** ([#1350](https://github.com/nanocoai/nanoclaw/issues/1350)) – Strong user demand for provider choice. Given the project’s active work on dual‑engine fallback (Claude ↔ Codex in #3057), integrating a third provider (Copilot) could be a natural next step. Likely candidate for v0.12 or later.

- **Dual‑engine quota fallback** ([#3057](https://github.com/nanocoai/nanoclaw/pull/3057)) – Already being merged as a major feature. Production‑tested and includes migration 017. This will significantly improve reliability for users on tight Claude quotas.

- **Pre‑compaction notifications** – Included in [#2440](https://github.com/nanocoai/nanoclaw/pull/2440) (closed), already shipped. Notifies agents before session compaction, improving state management.

- **Hardened agent image** ([#3150](https://github.com/nanocoai/nanoclaw/pull/3150)) – Operational feature merged today. Users can now skip local builds and pull a pre‑hardened image from the registry, lowering barriers for quick deployments.

These signals suggest the next stable release will focus on provider flexibility, reliability, and operational maturity.

## 7. User Feedback Summary

- **Provider dependency pain** – Issue #1350 explicitly states that “NanoClaw currently only supports Claude as the AI backend.” Users want alternatives to avoid vendor lock‑in and reduce costs. The 8 👍 indicate broad agreement.
- **Silent data loss in Telegram** – #3151 describes frustration: “messages carrying Bot API 10.1 rich_message content reach the agent completely empty — no text, no attachments, and no error.” This is a serious UX regression for Telegram users.
- **Slack thread blindness** – The fix in #2904 (now merged) addressed a long‑standing pain: users re‑@mentioning the bot deep in a thread would lose all intervening human messages. The community likely appreciates this fix.
- **Infrastructure improvements welcomed** – PR #3150 (hardened image) and #3060 (zombie reaping) show that users value operational stability and ease of setup.

Overall sentiment appears mixed: strong approval for fixes and performance improvements, but frustration with missing provider diversity and a new data‑loss bug.

## 8. Backlog Watch

| Item | Age | Status | Urgency |
|------|-----|--------|---------|
| [#1350 – GitHub Copilot SDK backend](https://github.com/nanocoai/nanoclaw/issues/1350) | 4 months old (created 2026-03-22) | Open, 3 comments, 8 👍 | High – long‑standing feature request with strong community support. No maintainer response visible. |
| [#3057 – Dual-engine quota fallback](https://github.com/nanocoai/nanoclaw/pull/3057) | 15 days old (created 2026-07-15) | Open, still being reviewed | Moderate – large PR, likely to be merged soon given production‑tested status. |
| [#3145 – DB backfill migration](https://github.com/nanocoai/nanoclaw/pull/3145) | 2 days old (created 2026-07-28) | Open, awaiting review | Low – needed for consistency, but not blocking. |
| [#3149 – CLI --rw flag](https://github.com/nanocoai/nanoclaw/pull/3149) | 1 day old | Open, awaiting review | Low – minor usability fix. |

The most critical backlog item is **#1350** – it has been unanswered for over four months despite high community interest. A status update or roadmap commitment from the maintainers would be beneficial.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-07-30

## Today’s Overview
Project activity was moderate, with four pull requests updated in the past 24 hours—two merged/closed and two open. No new releases were published. A single open bug report (#915) concerning scheduler authentication failure remains active, but a dedicated fix PR (#980) was opened today to address a related underlying issue (token persistence). The merged PRs add a new Grok CLI provider and enhance memory recall configurability, indicating continued focus on expanding provider support and memory management improvements.

## Releases
*No new releases were issued.*

## Project Progress
Two pull requests were merged or closed today:

- **[PR #981](https://github.com/nullclaw/nullclaw/pull/981) (merged/closed)** – feat(provider): add grok-cli provider for xAI Grok CLI. Implements a new provider that delegates to the local `grok` CLI, following the spawn-per-request pattern of the existing `codex-cli` provider. This expands the range of supported LLM backends.

- **[PR #961](https://github.com/nullclaw/nullclaw/pull/961) (merged/closed)** – feat(memory): add configurable auto-recall, recall_limit, max_context_bytes. Introduces three new JSON config keys under `memory` to control memory recall behavior (`auto_recall`, `recall_limit`, `max_context_bytes`). An earlier open PR (#979) covers the same feature and remains open, suggesting iterative refinement.

## Community Hot Topics
The most active item is the only open issue, **[#915](https://github.com/nullclaw/nullclaw/issues/915) [bug] Problem with scheduler unauthorized**, authored by scabros. It has 3 comments and 1 reaction. The user reports that the scheduler fails to work in both Telegram chat and other contexts when using an external Ollama host with qwen3.6:27b on an RTX 3090. Underlying need: reliable scheduler authentication and cross‑platform (Ubuntu) support for external LLM hosts.

No other issues or PRs received comments or reactions in the reporting window.

## Bugs & Stability
One bug is actively being addressed:

- **High severity** – **[Issue #915](https://github.com/nullclaw/nullclaw/issues/915)**: Scheduler unauthorized. The `/pair` endpoint generates a token but never persists it to disk, causing `readPairedToken()` to return `null` and scheduler authentication to fail.  
  **Fix in progress** – **[PR #980](https://github.com/nullclaw/nullclaw/pull/980) (open)** explicitly addresses this by persisting the paired token to disk during `/pair`, aiming to resolve the root cause (#839). No regressions or new crashes were reported.

## Feature Requests & Roadmap Signals
Two open PRs signal upcoming features likely to land in the next release:

- **[PR #979](https://github.com/nullclaw/nullclaw/pull/979) (open)** – Configurable memory recall: `auto_recall`, `recall_limit`, `max_context_bytes`. Gives users finer control over context injection, useful for reducing token cost or limiting noise from recalled memory.

- **[PR #980](https://github.com/nullclaw/nullclaw/pull/980) (open)** – Persistent scheduler token, critical for fixing the scheduler bug and improving reliability of scheduled tasks.

These reflect continued investment in memory management and scheduling stability.

## User Feedback Summary
The only direct user feedback comes from Issue #915 (scabros). The user’s pain point is clear:

- **Pain point:** Scheduler does not function at all when using an external Ollama LLM host on the same network. The `/pair` token is not persisted, so scheduled actions are unauthorized.
- **Use case:** Running NullClaw on Ubuntu with a remote Ollama (qwen3.6:27b, RTX 3090) for self-hosted AI assistant tasks, including scheduled interactions via Telegram.
- **Satisfaction indicator:** The user has not indicated satisfaction; the issue remains open with 3 comments and 1 reaction (positive, likely confirmation from others facing the same problem). The existence of a fix PR (#980) suggests the maintainers are actively addressing the concern.

## Backlog Watch
- **[Issue #915](https://github.com/nullclaw/nullclaw/issues/915)** has been open since May 15, 2026 (over 2 months) but was updated on July 29, indicating recent attention. The fix PR #980 is a positive signal.
- The referenced issue **[#839](https://github.com/nullclaw/nullclaw/issues/839)** (the root cause of token persistence) is not directly shown in the digest data but is explicitly mentioned in PR #980. This older issue likely requires further tracking if not already closed.
- No other long-unanswered issues or PRs were identified in the reporting window.

Maintainers should prioritize merging PR #980 and ensuring the scheduler token persistence fix is released, as the bug directly affects a core feature (scheduled tasks) for users with external LLM setups.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-07-30

## Today's Overview
IronClaw experienced a high-activity day with **50 pull requests and 12 issues updated** in the last 24 hours. The team fixed several critical (P1) bugs from the ongoing bug bash, including a Gmail auto-authorization flaw and a persistent turn-state store failure that required restart. Major feature work continues on the **attested signing** subsystem (8‑PR series, groups 4–8) and on **testing infrastructure** (WS10/12 coverage gates, regression promotion loop). No new releases were published today, but the release automation PR (#5598) remains open, suggesting a version bump may be imminent. Overall project health is strong, with rapid issue resolution and steady feature progress.

## Releases
**No new releases** were recorded in the last 24 hours. The release automation PR (#5598) is still open, targeting `ironclaw_common` 0.5.0 (breaking) and `ironclaw_skills` 0.4.0 (breaking), but no release has been cut.

## Project Progress
**14 PRs were merged or closed** today. Notable merges include:
- **Full‑thread QA artifacts (reborn)** — PR [#6346](nearai/ironclaw PR #6346) (closed) adds caller‑owned, multi‑run thread exports with bounded logs and a fixture importer for production-derived providers.
- **Gmail extension auto‑authorization fix** — Issue [#6348](nearai/ironclaw Issue #6348) (closed) resolved a security bug where reinstalling the extension granted access without user consent.
- **Turn‑state store latch degradation** — Issue [#6815](nearai/ironclaw Issue #6815) (closed) fixed a flush‑failure condition that rendered the instance 503 until manual restart.
- **Intermittent service_unavailability** — Issue [#6805](nearai/ironclaw Issue #6805) (closed) addressed the ~30 minute availability drop on Railway instances.
- **Stop button failure** — Issue [#6720](nearai/ironclaw Issue #6720) (closed) fixed a case where tasks ran indefinitely and the stop button was unresponsive.
- **Process journal kernel migration** — Issue [#6666](nearai/ironclaw Issue #6666) (closed) moves the turn‑run journal into `ironclaw_processes`, aligning with the neutral process architecture.

The open **attested signing** stack (PRs [#6813](nearai/ironclaw PR #6813), [#6818](nearai/ironclaw PR #6818), [#6822](nearai/ironclaw PR #6822), [#6809](nearai/ironclaw PR #6809), [#6811](nearai/ironclaw PR #6811), [#6769](nearai/ironclaw PR #6769)) is progressing: multi‑tenant isolation, Ledger clear‑signing, durable PostgreSQL/libSQL stores, and provider registration all landed in open PRs awaiting review.

## Community Hot Topics
- **Epic: Hermetic capability and journey testing platform** — Issue [#6524](nearai/ironclaw Issue #6524) (4 comments) is the most‑discussed issue. It aims to mechanically verify that every capability and critical user journey has deterministic, meaningful coverage. This epic will define the project’s quality infrastructure for the next quarter.
- **Attested signing PR series** — The 8‑PR signing campaign (groups 4–8) and the earlier PR [#3964](nearai/ironclaw PR #3964) (rebased) continue to attract attention. These PRs implement a hardened, multi‑chain signing flow with Ledger support and KMS ship‑gate. Community members and core contributors are actively reviewing the dense architectural changes.
- **Regression promotion loop** — PR [#6884](nearai/ironclaw PR #6884) adds a WS10 pipeline for promoting live‑caught failures into regression tests. It reconciles and selectively ports the earlier #6346 work, showing the team’s commitment to automated quality gates.

## Bugs & Stability
**Fixed today (critical P1):**
- **Gmail auto‑authorization after reinstall** (#6348) — security vulnerability where OAuth prompt was skipped.
- **Turn‑state store latch forever degraded** (#6815) — instance went 503 for 30+ minutes after a write‑behind flush failure.
- **Intermittent service_unavailable** (#6805) — every 30 minutes on Railway, all requests failed.
- **Stop button fails to cancel execution** (#6720) — tasks ran indefinitely; UI reported “Couldn’t stop this run.”
- **Automations not shown in web chat** (#6806) — P2 UX bug requiring manual navigation to see automation results.

**Newly reported (open, no fix yet):**
- **Test flake: ironclaw_reborn_composition** (#6887) — intermittent timeouts under full parallelism; root cause is run‑time contention, not a code defect.
- **Gemini OAuth tool call 400** (#6880) — tool schemas bypass `shape_tool_schema` entirely, causing 400 errors on every tool call.
- **Automation runs hit‑or‑miss** (#6879) — unattended runs sometimes execute as plain chat turns; structural pipeline issue.
- **Channel command gating** (#6877) — operator‑fallback identity lane lacks an activation guard; latent security trap (not exploitable today).
- **`/model set` silently drops trailing arguments** (#6875) — parsing bug: `/model set opus` sets model name to `"set"`.

The team fixed six P1/P2 issues today, keeping the bug bash momentum. The new Gemini and automation bugs are moderate severity but need prompt attention.

## Feature Requests & Roadmap Signals
- **Epic: Hermetic testing platform** (#6524) signals a major investment in coverage automation. Expect this to appear in the next RC as a set of deterministic test suites for all capabilities.
- **Process journal kernel into ironclaw_processes** (#6666, now closed/merged) aligns with the long‑term architecture of separating process lifecycle from turn management.
- **Channel command gating and identity lanes** (#6877) points to upcoming work on secure operator‑fallback identity resolution and door‑asymmetry decisions for WebUI command palette.
- **Attested signing** (groups 1–8) is the largest feature in flight; the final group (#6818) adds Ledger clear‑signing product. Once merged, this will enable cryptographically secured agent transactions across multiple chains.
- **WS12 scaling, artifact, and coverage gates** (PR [#6881](nearai/ironclaw PR #6881)) will harden CI infrastructure — scheduled stress/soak/mutation lanes and explicit flake handling. Likely to land before the next release.

## User Feedback Summary
User pain points surfaced during the bug bash:
- **Gmail extension reinstall** (#6348) violates user expectation of explicit OAuth consent — a trust and security issue.
- **Service unavailability every 30 minutes** (#6805) degrades production reliability; users cannot depend on the Railway instance.
- **Stop button not working** (#6720) frustrates power users who need to abort runaway tasks.
- **Automations hidden from web chat** (#6806) creates confusion — users expect automation output to appear in the same conversation.
- **Test flakiness** (#6887) affects developer confidence in CI when running tests in parallel.

Satisfaction: The team’s rapid response to bug bash reports (all P1s closed within 24–48 hours) indicates strong operational discipline.

## Backlog Watch
- **Old signing PR #3964** — This group‑4 PR from May 2026 was rebased (1184 commits behind) and remains open. It carries the durable one‑shot challenge store and WebAuthn verifier. With the newer 8‑PR series advancing, this older PR may need to be superseded or explicitly closed.
- **Release PR #5598** — Open since July 3, repeatedly updated as dependencies bump. It blocks the official release of `ironclaw_common` 0.5.0 (breaking) and `ironclaw_skills` 0.4.0. Maintainers should review and either merge or trim scope.
- **Issue #6887** — The `ironclaw_reborn_composition` test suite flake could grow into a CI blocker. A fix PR (#6886) already exists to complete state‑machine testing, but the parallelism contention may need a separate root‑cause investigation.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest – 2026-07-30

**Generated from GitHub data (netease-youdao/LobsterAI) for the last 24 hours.**

---

## 1. Today's Overview

Activity remained high with **13 pull requests merged or closed** in the last 24 hours and **2 open PRs** (one automated dependency update, one long‑standing bug fix). No new releases were published. The focus was on **polishing the cowork feature** – improving side‑chat input handling, adding selected‑text tags, fixing export modal layering, scroll jumps, and IM message flicker. Several infrastructure fixes landed (Windows caption button hover colors, update check interval, login callback resilience). A notable revert of the “run‑safety‑contract” feature indicates quality‑gate issues. Project health appears strong with continuous iterative improvements.

---

## 2. Releases

**None** – no new releases in the last 24 hours.

---

## 3. Project Progress

The following **13 PRs were merged or closed** (all on 2026-07-29), grouped by area:

### Core Features & Enhancements
- **[PR #2405] feat(cowork): add selected text tags to side chat** [merged]  
  Adds removable context tags for selected text, with support for direct sending and follow‑up editing, plus state safeguards and tests.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/2405

- **[PR #2404] refactor/kimi k3 auto only compat** [merged]  
  Compatibility refactoring for Kimi K3 auto‑only mode (areas: renderer, build, docs, main, openclaw).  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/2404

### Bug Fixes
- **[PR #2406] fix(cowork): improve side chat input handling** [merged]  
  Accumulates selected text excerpts, removes product‑level length limit, retains bounded context and transport safety.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/2406

- **[PR #2376] fix(cowork): render export modal above sidebar** [merged]  
  Uses a body portal to avoid stacking context conflicts.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/2376

- **[PR #2364] fix(cowork): prevent scroll jumps on session refresh** [merged]  
  Scopes refresh events by session ID and preserves loaded message history.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/2364

- **[PR #2363] fix(cowork): prevent periodic IM message flicker** [merged]  
  Compares matching history windows during reconciliation, preserves older messages.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/2363

- **[PR #2360] fix(auth): preserve local callback across login retries** [merged]  
  Reuses active callback server for repeated and concurrent login attempts; adds diagnostics and tests.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/2360

- **[PR #2355] fix(window): align Windows caption button hover colors** [merged]  
  Theme‑aware hover states matching sidebar controls.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/2355

- **[PR #2346] fix(cowork): open email diagnostics in a new chat** [merged]  
  Prevents stale history or IM sessions from overriding the new chat.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/2346

### Chores & Housekeeping
- **[PR #2407] Release/2026.7.24** [closed] – Release branch for version 2026.7.24 (containing the above fixes).  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/2407

- **[PR #2347] chore(updater): reduce automatic update check interval** [merged]  
  Changed from 12 hours to 2 hours for quicker user access to updates.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/2347

- **[PR #2403] revert(openclaw): remove run‑safety‑contract gate for no‑progress token burn** [merged]  
  Client‑side Run Safety feature (PR #2400) introduced release‑blocking issues (receipt identity keying, false‑success followups, byte‑accounting mismatches). Reverted to prior behavior.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/2403

### Stale PRs Closed
- **[PR #1322] fix(cowork): true LRU eviction for LLM memory judge cache** [closed] – Fix from April that was finally merged or closed.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/1322

---

## 4. Community Hot Topics

No issues were updated in the last 24 hours. Among the **open PRs**, the most notable are:

- **[PR #1277] chore(deps-dev): bump electron group (2 updates)** – An automated dependency update that bumps `electron` from 40.2.1 to 43.2.0 and `electron‑builder`; remains open after 4 months.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/1277  
  *Underlying need*: Keeping Electron up‑to‑date for security and feature parity.

- **[PR #1232] fix(scheduledTask): first‑run result not pushed to UI** – Chinese‑described bug fix for a long‑standing scheduling issue (created April 1). No comments.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/1232  
  *Underlying need*: Reliable real‑time UI updates for scheduled tasks.

---

## 5. Bugs & Stability

No new bugs were reported as issues. However, the **reverted PR #2403** indicates significant stability concerns:

| Severity | Bug | Status | Fix PR |
|----------|-----|--------|--------|
| High | Run‑safety contract feature (client‑side) caused receipt identity keying, false‑success followups, compaction runId handling, and byte‑accounting mismatches. | Reverted – feature removed. | [#2403](https://github.com/netease-youdao/LobsterAI/pull/2403) |
| Medium | Cowork side‑chat input handling: missing context accumulation, product‑level length limit. | Fixed | [#2406](https://github.com/netease-youdao/LobsterAI/pull/2406) |
| Medium | Scroll jumps during session refresh. | Fixed | [#2364](https://github.com/netease-youdao/LobsterAI/pull/2364) |
| Medium | Periodic IM message flicker during reconciliation. | Fixed | [#2363](https://github.com/netease-youdao/LobsterAI/pull/2363) |
| Low | Login callback not preserved across retries. | Fixed | [#2360](https://github.com/netease-youdao/LobsterAI/pull/2360) |
| Low | Windows caption button hover colors misaligned. | Fixed | [#2355](https://github.com/netease-youdao/LobsterAI/pull/2355) |
| Low | Email diagnostics opening in stale IM session. | Fixed | [#2346](https://github.com/netease-youdao/LobsterAI/pull/2346) |

All bugs with active fixes have been merged.

---

## 6. Feature Requests & Roadmap Signals

No explicit feature requests were filed today. Based on merged PRs and recent activity:

- **Selected text context tags** ([#2405](https://github.com/netease-youdao/LobsterAI/pull/2405)) is a user‑visible enhancement likely to be included in the next release.
- **Kimi K3 auto‑only compatibility** ([#2404](https://github.com/netease-youdao/LobsterAI/pull/2404)) suggests ongoing work to support new hardware or platform profiles.
- The **reverted run‑safety feature** will likely be re‑evaluated and redesigned for a future release after addressing identified issues.

Prediction: The next minor release (after 2026.7.24) will focus on stabilizing the cowork experience and potentially reintroduce a corrected run‑safety mechanism.

---

## 7. User Feedback Summary

No direct user feedback or comments are available in this data set. However, the pattern of fixes in the last 24 hours implies the following real‑world pain points were addressed:

- Cowork side‑chat felt incomplete without visible selected‑text context.
- Export modal could be hidden behind other UI elements.
- Scrolling was jumpy during session refresh.
- IM messages flickered periodically.
- Login retries could fail silently.
- Windows users saw inconsistent title‑bar button hover styles.

With these fixes merged, user satisfaction with the cowork feature and overall UI polish should improve measurably.

---

## 8. Backlog Watch

The **most critical open item** requiring maintainer attention:

- **[PR #1232] fix(scheduledTask): fix first‑run result not pushed to UI** – Created **2026-04-01**, last updated **2026-07-29**. A Chinese‑language bug fix for a core scheduling feature. After nearly 4 months without being merged or commented on, this may need a review or integration into the next release.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/1232

- **[PR #1277] chore(deps-dev): bump electron group** – While automated, it updates Electron from 40 to 43 (major version jump). It remains open since April 2, possibly awaiting manual conflict resolution or testing.  
  🔗 https://github.com/netease-youdao/LobsterAI/pull/1277

No stale issues (0 open issues) – the backlog consists only of these two open PRs.

---

*End of digest.*

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

**Moltis Project Digest – 2026-07-30**

---

### 1. Today's Overview
Project activity remained moderate on 2026-07-30. No new issues were created or updated in the last 24 hours, but five pull requests saw updates, including one merged PR that enhances PWA push notifications. The four open PRs all address substantial features: ACP agent exposure over stdio, Slack acknowledgment lifecycle, instrumentation infrastructure, and privileged tool access gating. The single merged PR (#1173) was closed after being updated yesterday, signaling steady progress toward a more polished user experience. Overall, the project appears to be in a healthy development cadence, with the maintainer (penso) driving multiple concurrent initiatives.

### 2. Releases
No new releases were published today. The latest release remains unchanged; no migration notes or breaking changes are required.

### 3. Project Progress
One pull request was merged/closed today:
- **#1173** ([moltis-org/moltis PR #1173](https://github.com/moltis-org/moltis/pull/1173)) – **feat(pwa): make push notifications reliable and non-disruptive**  
  – This PR improves PWA push notification reliability, privacy, ordering, and cross-tab behavior. It re-alerts for newer messages without losing earlier-message counts, uses a generic privacy-safe title, strips rich formatting, and maintains an app-wide unread badge.

No other closed or merged PRs were recorded today. The four open PRs remain under active iteration.

### 4. Community Hot Topics
No issues or PRs generated substantial comments or reactions today (all have 0 comments and 0 👍). However, the most notable ongoing discussions revolve around the following open PRs:

- **#1169** ([moltis-org/moltis PR #1169](https://github.com/moltis-org/moltis/pull/1169)) – Exposes Moltis as an ACP agent over stdio, routing through the cancellable `LiveChatService`. Enforces session isolation and rate limits.  
  *Underlying need*: Enables third-party tooling and AI orchestration systems to integrate Moltis via the Agent Communication Protocol, expanding interoperability.

- **#1166** ([moltis-org/moltis PR #1166](https://github.com/moltis-org/moltis/pull/1166)) – Adds per-message Slack acknowledgment reactions, phase tracking, reconnection supervision, and Block Kit support.  
  *Underlying need*: Slack bots lack typing indicators; reactions serve as a reliable progress signal. This ensures safe lifecycle handling under queueing, cancellation, and retries.

- **#1174** ([moltis-org/moltis PR #1174](https://github.com/moltis-org/moltis/pull/1174)) – Adds backend-neutral instrumentation, Langfuse v4 export, operational OTLP backends, and end-user reaction feedback. Records immutable turns, cache-aware tokens, and provider failover attribution.  
  *Underlying need*: Observability and feedback loops are critical for production AI agents; this infrastructure supports debugging, cost tracking, and user satisfaction measurement.

- **#1170** ([moltis-org/moltis PR #1170](https://github.com/moltis-org/moltis/pull/1170)) – Gates `/sh` and other privileged tools behind a per-account `operators` list, separating access from privilege.  
  *Underlying need*: Security hardening – previously, any allowed user could run host commands. This change enforces operator-level controls across all execution paths.

### 5. Bugs & Stability
No new bugs, crashes, or regressions were reported today in issues. The one PR that could be considered a security fix is **#1170**, which addresses a privilege escalation vector (access vs. operator separation). It remains open but already has a solution in progress. No high-severity stability issues are currently visible.

### 6. Feature Requests & Roadmap Signals
While no explicit feature requests were filed as issues, the open PRs strongly indicate the following roadmap priorities:
- **ACP agent exposure** (PR #1169) – Likely to be included in the next release, as it is a foundational integration point.
- **Slack UX improvements** (PR #1166) – High demand for reliable chatbot feedback in Slack; expect this to merge soon.
- **Observability & feedback** (PR #1174) – Instrumentation is a common ask for production deployments; probable candidate for the next minor version.
- **PWA notifications** (PR #1173, already merged) – The merged notification work suggests a push toward mobile and desktop PWA support, likely a stable release goal.

### 7. User Feedback Summary
No direct user feedback or pain points were reported in issues today. Given the nature of the open PRs, the implicit user needs include:
- Secure multi-tenant setups (PR #1170)
- Reliable Slack interactions (PR #1166)
- Integration with external AI orchestrators (PR #1169)
- Visibility into agent performance and user reactions (PR #1174)

No explicit dissatisfaction signals were captured in the last 24 hours.

### 8. Backlog Watch
No open issues or PRs were found to be long-unanswered or needing maintainer attention. All items with recent updates are being actively handled by the sole contributor (penso). The project has zero open issues, which suggests a clean backlog and active maintenance. The four open PRs are all from the same author and are being iterated – none are stalled. No action required from the community or maintainer on outstanding items.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest — 2026-07-30

## 1. Today's Overview

Project activity remains very high, with **19 issues** and **50 pull requests** updated in the last 24 hours. Of those, **15 issues** are open/active and **4 closed**; on the PR side **36 are open** and **14 merged/closed**. No new releases were published today. The community is actively reporting regressions, requesting UX improvements, and contributing substantial feature PRs (e.g., desktop GUI automation, reranker support, creator checkpoints). CI and stability issues are drawing maintainer attention, but the overall pulse indicates a healthy, fast-moving open-source project.

## 2. Releases

**No new releases today.** The latest available version remains QwenPaw 2.0.1 (Desktop). No changelog, breaking changes, or migration notes to report.

## 3. Project Progress

**14 PRs were merged or closed** in the last 24 hours. The most notable include:

- **Security fix:** [`#6487`](https://github.com/agentscope-ai/QwenPaw/pull/6487) (Closed) – *fix: restrict import-local source path to prevent arbitrary directory exfiltration* – merges a patch that tightens import path validation.

- **Under review / ready for human review:**
  - [`#6531`](https://github.com/agentscope-ai/QwenPaw/pull/6531) – *fix(acp): add models field to new_session response* – enables external ACP clients to discover available models.
  - [`#6486`](https://github.com/agentscope-ai/QwenPaw/pull/6486) – *fix(matrix): probe vodozemac E2EE backend so encryption works on Python 3.12*.
  - [`#6424`](https://github.com/agentscope-ai/QwenPaw/pull/6424) – *feat(computer-use): native desktop GUI automation for Windows and macOS* – a major new builtin tool.
  - [`#6312`](https://github.com/agentscope-ai/QwenPaw/pull/6312) – *feat(console): configurable theme/skin module (Task 1 draft)*.

- **Feature PRs (open):**
  - [`#6556`](https://github.com/agentscope-ai/QwenPaw/pull/6556) – *feat(creator): creation checkpoints, home redesign, media recovery, export/import, and bilingual guide*.
  - [`#6553`](https://github.com/agentscope-ai/QwenPaw/pull/6553) – *feat: redesign app center* (My Apps, Official Apps, App Market tabs).
  - [`#6398`](https://github.com/agentscope-ai/QwenPaw/pull/6398) – *feat: add reranker support for ReMe memory search (backend)*.
  - [`#6383`](https://github.com/agentscope-ai/QwenPaw/pull/6383) – *feat(sandbox): add unelevated sandbox for windows*.
  - [`#6269`](https://github.com/agentscope-ai/QwenPaw/pull/6269) – *feat(checkpoints): add workspace checkpoint management*.

- **Fix PRs (open, recently updated):**
  - [`#6561`](https://github.com/agentscope-ai/QwenPaw/pull/6561) – *fix(mcp): ensure exposed tool names start with a letter*.
  - [`#6535`](https://github.com/agentscope-ai/QwenPaw/pull/6535) – *fix(cloudpaw): accept mission verification kwargs*.
  - [`#6554`](https://github.com/agentscope-ai/QwenPaw/pull/6554) – *fix(providers): add MiniMax context windows to the static catalog*.
  - [`#6543`](https://github.com/agentscope-ai/QwenPaw/pull/6543) – *fix(onebot): clean text and send local media*.
  - [`#6500`](https://github.com/agentscope-ai/QwenPaw/pull/6500) – *fix(browser): make unauthenticated local CDP exposure opt-in*.

## 4. Community Hot Topics

The most active discussions (by comment count) reflect both bug reports and feature requests:

- **[#6537](https://github.com/agentscope-ai/QwenPaw/issues/6537)** – *[Bug]: Skill tags disappear on restart (regression of #3270)*  
  **9 comments**. Community frustration over a regression that loses user-set skill tags after restart, despite the data being saved to disk. The issue is actively debated, with users asking for a quick fix.

- **[#6460](https://github.com/agentscope-ai/QwenPaw/issues/6460)** – *QwenPaw 2.0.1 首页/会话在 Edge+Wayland 下单标签高 CPU 占用*  
  **4 comments**. High CPU usage on Linux + Wayland + Edge, likely triggered by large result sets or WebSocket push. The user suspects rendering inefficiency.

- **[#6524](https://github.com/agentscope-ai/QwenPaw/issues/6524)** – *[Bug]: MCP 后端重启后客户端无法自动恢复*  
  **3 comments**. When the MCP server restarts, QwenPaw does not automatically reconnect; the user must manually run `list mcp`. Identified as a session reuse problem.

- **[#6056](https://github.com/agentscope-ai/QwenPaw/issues/6056)** – *[Bug]: Background offload kills subprocess immediately — LLM-provided timeout is silently ignored*  
  **3 comments**. A long-standing bug (opened Jul 13) where background shell commands are killed instantly. This is likely the root cause of regression #6245.

- **[#6415?]** Not in provided data, but **#6475** (feature request for `notice_after_complete`) has 2 comments and is gaining interest.

## 5. Bugs & Stability

Ranked by severity (H=High, M=Medium, L=Low):

| Severity | Issue | Description | Fix PR exists? |
|----------|-------|-------------|----------------|
| **H** | [#6537](https://github.com/agentscope-ai/QwenPaw/issues/6537) | Skill tags lost on restart (regression of #3270) – data integrity bug. | No |
| **H** | [#6534](https://github.com/agentscope-ai/QwenPaw/issues/6534) | Windows NSIS installer infinite loop – installation impossible. | No |
| **H** | [#6245](https://github.com/agentscope-ai/QwenPaw/issues/6245) | Session permanently blocked when shell command exceeds coordinator deadline (regression from #6056 fix). | No |
| **H** | [#6555](https://github.com/agentscope-ai/QwenPaw/issues/6555) | Dream/memory compression misses early-session events when context scrolls out. | No |
| **M** | [#6524](https://github.com/agentscope-ai/QwenPaw/issues/6524) | MCP backend restart – client doesn’t auto-reconnect. | No |
| **M** | [#6541](https://github.com/agentscope-ai/QwenPaw/issues/6541) | Scroll context compression triggers MODEL_EXECUTION_ERROR on DeepSeek – wrong `role=user`. | No |
| **M** | [#6557](https://github.com/agentscope-ai/QwenPaw/issues/6557) | MCP tool names starting with hyphen cause strict LLM API 400 errors. | Yes: [#6561](https://github.com/agentscope-ai/QwenPaw/pull/6561) |
| **M** | [#6558](https://github.com/agentscope-ai/QwenPaw/issues/6558), [#6559](https://github.com/agentscope-ai/QwenPaw/issues/6559), [#6560](https://github.com/agentscope-ai/QwenPaw/issues/6560) | Multiple chat session UI data integrity issues (messages lost, unwanted forking, lack of copy/undo). | No |
| **L** | [#6529](https://github.com/agentscope-ai/QwenPaw/issues/6529) | ACP new_session missing `models` field. | Yes: [#6531](https://github.com/agentscope-ai/QwenPaw/pull/6531) |
| **L** | [#6551](https://github.com/agentscope-ai/QwenPaw/issues/6551) | Aliyun coding plan models not aligned with official website. | No |
| **L** | [#6563](https://github.com/agentscope-ai/QwenPaw/issues/6563) | CI workflow blocks all fork PRs (real-behavior-proof). | No |
| **L** | [#6056](https://github.com/agentscope-ai/QwenPaw/issues/6056) | Background offload kills subprocess immediately (underlying cause of #6245). | No |

## 6. Feature Requests & Roadmap Signals

Features requested or prototyped recently that are likely to appear in the next minor release (2.1.0 or 2.0.2):

- **`notice_after_complete` tool** ([#6475](https://github.com/agentscope-ai/QwenPaw/issues/6475)) – Allows agent to start a long task, respond to the user, and notify after completion. High community interest.
- **Chat session UX improvements** ([#6560](https://github.com/agentscope-ai/QwenPaw/issues/6560)) – Copy, undo, stop generation, mission mode, scroll performance, session ID, context transfer.
- **Desktop GUI automation** (PR [#6424](https://github.com/agentscope-ai/QwenPaw/pull/6424)) – Native computer-use tool for Windows and macOS. In review, likely to land soon.
- **Workspace checkpoint management** (PR [#6269](https://github.com/agentscope-ai/QwenPaw/pull/6269)) – Recoverable conversation history via shadow Git store.
- **App center redesign** (PR [#6553](https://github.com/agentscope-ai/QwenPaw/pull/6553)) – Three-tab layout (My Apps, Official, Market) with featured badges.
- **Configurable theme/skin module** (PR [#6312](https://github.com/agentscope-ai/QwenPaw/pull/6312)) – Draft for product branding customization.
- **ReMe memory search reranker support** (PR [#6398](https://github.com/agentscope-ai/QwenPaw/pull/6398)) – Over-fetch + rerank improves memory retrieval quality.

Predictions: The next version (2.1.0) will likely include the computer-use tool, checkpoint system, and the creator plugin improvements (#6556), along with fixes for the skill tag regression and MCP auto-reconnect.

## 7. User Feedback Summary

**Pain points (direct quotes & paraphrases from issues):**
- *“Skill tags disappear on restart – very frustrating.”* (#6537)
- *“CPU usage skyrockets on Edge+Wayland – only on QwenPaw pages.”* (#6460)
- *“Switching sessions loses messages; replies re-render from scratch.”* (#6558)
- *“Session list becomes chaotic with unwanted forks – no parent-child grouping.”* (#6559)
- *“Can’t copy agent reply text, ESC doesn’t stop generation, no undo.”* (#6560)
- *“Windows installer shows 'still running' error in infinite loop – installation impossible.”* (#6534)
- *“When MCP server restarts, client is dead – must manually re-list.”* (#6524)
- *“Shell commands that hang permanently block the session.”* (#6245)

**Positive signals:**
- Active contribution from first-time contributors (PRs #6312, #6562, #6531, #6486, #6543).
- Enthusiasm for new features: computer-use, checkpoints, reranker, creator plugin.
- Users providing detailed reproduction steps (e.g., #6541 with DeepSeek context compression).

**Overall satisfaction:** Users appreciate the rapid development but are frustrated by regressions (especially skill tags and session blocking). The community is proactive in reporting bugs and offering PRs.

## 8. Backlog Watch

Issues or PRs that have remained open for an extended period without maintainer response, or require re-triage:

- **[#6056](https://github.com/agentscope-ai/QwenPaw/issues/6056)** (Jul 13) – Background offload kills subprocess. This is the root cause of a later regression (#6245). Needs a fix (no assignee yet).
- **[#6245](https://github.com/agentscope-ai/QwenPaw/issues/6245)** (Jul 18) – Session permanently blocked after shell timeout. Marked as regression but no fix PR linked.
- **[#6102](https://github.com/agentscope-ai/QwenPaw/pull/6102)** (Jul 14) – Boundary meta-test for isolation failures. Still open, awaiting review.
- **[#6103](https://github.com/agentscope-ai/QwenPaw/pull/6103)** (Jul 14) – Raise frontend vitest coverage thresholds. Simple change but stale.
- **[#6496](https://github.com/agentscope-ai/QwenPaw/issues/6496)** (Jul 27, closed) – Legacy plugins disabled due to version derivation. Already closed, but the underlying behavior may still affect other plugins.

Maintainers should prioritize reviewing the critical regressions (#6537, #6534, #6245) and the CI blocker for fork PRs (#6563) to avoid stalling community contributions.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-07-30

## Today’s Overview
ZeroClaw is in a high-activity phase with 50 PRs updated in the last 24 hours, of which 2 were merged/closed. Seven issues were touched (5 open, 2 closed). The project is preparing two concurrent releases: the maintenance train v0.8.4 (target July 31) and the weekly non-breaking v0.8.5. Security hardening and architectural feature work (A2A, goal tools, declarative skills) dominate the open PR queue, with multiple high‑risk, XL‑sized contributions in flight. The community is active, though many PRs await author action.

## Releases
No new releases were published today.

## Project Progress
- **Closed Issues:**
  - #9186 [[Bug]: MCP stdio: response id not matched, 30s hard timeout vs 180–600s tool budget, Mutex held for whole call](https://github.com/zeroclaw-labs/zeroclaw/issues/9186) — severity S1 (workflow blocked), fixed by a PR not shown in the top 20.
  - #9278 [[Bug]: context_compression.enabled defaults true while runtime ignores it](https://github.com/zeroclaw-labs/zeroclaw/issues/9278) — severity S2, resolved.
- **Merged/Closed PRs (2):** No details available in the displayed data; likely small dependency bumps or CI fixes.

## Community Hot Topics
- **Most Active Issue (by comments):**
  - #8810 [[Bug]: Documentation is wrong - Telegram example](https://github.com/zeroclaw-labs/zeroclaw/issues/8810) — 2 comments. A user reports that the Telegram example code in the docs is incorrect, calling for a dedicated end‑to‑end guide. A related PR #9242 is open to address this.
- **Largest Feature PRs (all remain open):**
  - #8965 [feat(skills): declarative auto‑activation with provider switch and image‑turn tool blocking](https://github.com/zeroclaw-labs/zeroclaw/pull/8965) — XL size, touches agent, channels, runtime, skills.
  - #8687 / #8688 / #8689 — a three‑part series by `vrurg` adding goal tools, admission control, and channel `/goal` commands (each XL, high risk).
  - #9208 [fix(runtime): stop per‑iteration tool‑schema deep clones](https://github.com/zeroclaw-labs/zeroclaw/pull/9208) — XL, high priority P1, addressing performance regression.
  - #9324 [feat(a2a): outbound client config, shared wire‑model, tools](https://github.com/zeroclaw-labs/zeroclaw/pull/9324) — XL, implementing the A2A outbound client RFC.

The community is pushing for **declarative skill activation**, **goal‑driven agent control**, and **inter‑agent communication** — all signals of growing adoption beyond basic chatbots.

## Bugs & Stability
| Issue | Severity | Status | Summary | Fix PR? |
|-------|----------|--------|---------|---------|
| #9186 | S1 – workflow blocked | Closed | MCP stdio response id mismatch, 30s hard timeout vs configured budget, mutex held for whole call | Yes (merged) |
| #9278 | S2 – degraded behavior | Closed | `context_compression.enabled` default true but runtime ignores it | Yes (merged) |
| #8810 | S2 – degraded behavior | Open | Documentation wrong for Telegram example | PR #9242 open |
| #9384 | High risk | Open PR | Shell command path arguments – symlink escape hardening | Open |
| #9401 | High risk | Open PR | Preserve shell cwd across sandbox wrappers | Open |
| #9433 | High risk | Open PR | Tool allowlists not enforced in `ensure_no_escalation_beyond` | Open |
| #9477 | High risk | Open PR | Recover tool invocations wrapped in `<tools>` tag | Open |

Multiple **security‑critical** fixes are in the open PR queue (#9384, #9401, #9433), indicating that the team is proactively hardening the sandbox and policy enforcement.

## Feature Requests & Roadmap Signals
- **#9549** [RFC: Build a community‑powered local model advisor for ZeroClaw](https://github.com/zeroclaw-labs/zeroclaw/issues/9549) — a proposal to help users choose the best local model (Ollama, llama.cpp) based on hardware and use case. Status: needs maintainer review.
- **#9459** [Tracker: v0.8.5 weekly non‑breaking release](https://github.com/zeroclaw-labs/zeroclaw/issues/9459) — the upcoming release that will likely absorb many of the bugfixes and smaller enhancements currently in flight.
- **#8357** [Tracker: v0.8.4 maintenance train](https://github.com/zeroclaw-labs/zeroclaw/issues/8357) — feature‑frozen, target July 31. This release will include the MCP fix, context compression default correction, and possibly security backports.
- **#8965** (PR) declarative skill auto‑activation and **#9324** (PR) A2A tools are strong candidates for v0.8.5 if merged in time.

## User Feedback Summary
- **Documentation confusion:** Issue #8810 explicitly states “the documentation is wrong” for the Telegram channel. The user appreciates the Rust implementation but criticizes incomplete examples. A companion PR #9242 aims to replace the two‑line example with a full guide.
- **Broken external link:** #9550 reports that the LinkedIn link on the GitHub organization page points to a non‑existent page. Minor but affects first impressions for new contributors.
- **Configuration inconsistency:** #9278 (closed) highlights that a default (`context_compression.enabled = true`) is ignored at runtime, causing confusion during onboarding.
- **Performance concerns:** The S1 bug #9186 involved a 30‑second hard timeout conflicting with user‑configured tool budgets – a real friction point for heavy tool users.

Overall sentiment is **cautiously positive**: users value the technical foundations but find gaps in documentation and default behavior. The maintainers are responding quickly with fixes and enhanced guides.

## Backlog Watch
- **#8357** [v0.8.4 maintenance train](https://github.com/zeroclaw-labs/zeroclaw/issues/8357) – open since June 26, updated yesterday but still incomplete. As the target date (July 31) is tomorrow, maintainers should finalise the milestone.
- **PRs with `needs-author-action` label** (many, e.g., #8965, #8687–8689, #9208, #9324, #9433, #9477) – these are large contributions awaiting author responses to review comments. If left unaddressed, they risk stalling the v0.8.5 release.
- **#8810** [Documentation bug](https://github.com/zeroclaw-labs/zeroclaw/issues/8810) – open for 23 days, with a fix PR (#9242) that also carries `needs-author-action`. The original reporter may need to verify the updated guide.
- **#9549** [Local model advisor RFC](https://github.com/zeroclaw-labs/zeroclaw/issues/9549) – awaiting maintainer decision. This could significantly improve the new‑user experience and aligns with the community’s interest in local models.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*