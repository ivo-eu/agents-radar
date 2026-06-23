# OpenClaw Ecosystem Digest 2026-06-23

> Issues: 39 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-23 10:50 UTC

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

# OpenClaw Project Digest — 2026-06-23

## 1. Today’s Overview

The OpenClaw project saw extremely high activity over the past 24 hours, with **500 PRs updated** (84 merged/closed) and **39 issues updated** (27 open, 12 closed). No new releases were published today, but the 2026.6.9 release is still being stabilized with a heavy influx of regression fixes and targeted enhancements. The most debated topics revolve around session state integrity, runtime reliability (especially the ACPX path and Windows support), and long-standing feature requests (e.g., agent-facing scheduling, Telegram reply context). The maintainer queue is under pressure — many issues carry the `clawsweeper:needs-maintainer-review` or `clawsweeper:needs-product-decision` labels, suggesting roadmap decisions are pending.

## 2. Releases

**None.** No new releases were tagged today. The latest published version remains the 2026.6.9 release.

## 3. Project Progress

Today **84 PRs were merged or closed**, the majority being bug fixes and internal improvements. Notable merged/closed contributions include:

- **[Bug #76729]** *(Feishu replies disappearing after compaction)* was closed — a Diamond Lobster P1 fix now shipped.
- **[Legacy cost handling #37878]** *(numeric-string costs)* merged after a long review cycle.
- **[Bug #89868]** *(stop-button abort race with compaction)* closed — likely sourced from a linked PR.
- **[Bug #92260]** *(heartbeat reasoning payload leak)* closed — a fix for `includeReasoning=false` not being respected.
- **[PR #68936]** *(autofix pipeline + Windows daemon)* merged after two months.
- **[PR #95296]** *(fix Anthropic OAuth user-agent)* moved to final review, addressing token validation failures on `claude-cli` OAuth path.

Several long-open PRs also received updates today, indicating maintainers are revisiting the backlog. However, many remain in `waiting on author` or `needs proof` status.

## 4. Community Hot Topics

The following issues and PRs generated the most discussion and reactions:

- **[#92201 – Embedded runner thinking signatures intermittently invalid on replay](https://github.com/openclaw/openclaw/issues/92201)**  
  *13 comments, 1 👍*  
  A **critical P1 Diamond Lobster** bug where Anthropic thinking block signatures fail during replay, causing the recovery wrapper never to fire because error messages are genericized. This is a high-severity session-state/integrity issue for Slack plugin users.

- **[#92043 – 180 s compaction timeout blocks legitimate long compactions permanently](https://github.com/openclaw/openclaw/issues/92043)**  
  *9 comments, 2 👍*  
  Users report that the lowered 180 s timeout transforms slow-but-recoverable compactions into permanent failures. Community is asking for partial-progress reuse or a configurable penalty.

- **[#71712 – RFC: Agent-facing scheduling API with non-forgeable provenance](https://github.com/openclaw/openclaw/issues/71712)**  
  *5 comments, 0 👍 (but still active after two months)*  
  This stale-but-important P2 outlines a roadmap feature for agents to dynamically schedule cron jobs. The discussion is blocked by a need for product decision and security review.

- **[#90404 – ACPX TypeError: Cannot use "in" operator on numeric value](https://github.com/openclaw/openclaw/issues/90404)**  
  *5 comments, 1 👍*  
  A P1 bug causing `sessions_spawn(runtime="acp", agentId="claude")` to fail with a TypeError. Root cause identified, and a linked PR is open.

- **[#95750 – Main-session restart-recovery has no cross-boot retry budget → death-loop across reboots](https://github.com/openclaw/openclaw/issues/95750)**  
  *1 comment, but high-severity Diamond Lobster P1*  
  A freshly reported blocker where a wedged session on Claude Code CLI causes the gateway to crash-loop across reboots. Marked `queueable-fix` and `source-repro`.

**PRs with notable engagement** (older PRs with repeated updates):  
- **[#79401 – Emit structured runtime incidents](https://github.com/openclaw/openclaw/pull/79401)** — ready for maintainer look after months of proof collection.  
- **[#78664 – Cache provider tool schema normalization](https://github.com/openclaw/openclaw/pull/78664)** — a performance improvement waiting for maintainer review.

## 5. Bugs & Stability

A large share of today’s open issues are **P1 Diamond Lobster** bugs impacting session state, message delivery, or causing crash loops. Key reports below, ranked by severity:

| Issue | Severity | Summary | Fix PR? |
|-------|----------|---------|---------|
| [#92201](https://github.com/openclaw/openclaw/issues/92201) | P1 Diamond Lobster | Thinking signatures invalid on replay; recovery never fires | No (linked open) |
| [#92043](https://github.com/openclaw/openclaw/issues/92043) | P1 Diamond Lobster | 180s compaction timeout fails identically every turn | No (linked PR open) |
| [#90404](https://github.com/openclaw/openclaw/issues/90404) | P1 Diamond Lobster | ACPX TypeError on `in` operator | Yes (linked PR open) |
| [#95750](https://github.com/openclaw/openclaw/issues/95750) | P1 Diamond Lobster | Cross-boot death-loop on Claude Code CLI session wedges | No, but `queueable-fix` |
| [#85844](https://github.com/openclaw/openclaw/issues/85844) | P1 Platinum Hermit | Auto-update leaves stale hashed bundle imports | No, needs maintainer review |
| [#93465](https://github.com/openclaw/openclaw/issues/93465) | P1 Platinum Hermit | Windows ACPX runtime `spawn EINVAL` – runtime=acp unusable | No (needs live repro) |
| [#95141](https://github.com/openclaw/openclaw/issues/95141) | P1 Platinum Hermit | Gateway restart during in-flight WebChat turn drops the message | No (needs live repro) |
| [#96046](https://github.com/openclaw/openclaw/issues/96046) | Regression (no rating) | Gateway rejects valid plugin config as "plugin not found" on 2026.6.9 | No (new today) |
| [#95985](https://github.com/openclaw/openclaw/issues/95985) | P2 Diamond Lobster | Bootstrap file writes cause cacheWrite spiral on long sessions | No |

Newly reported bugs today (2026-06-23) include:
- **#95998** – `ensureGlobalUndiciEnvProxyDispatcher()` breaks COS chunked upload via qqbot plugin.
- **#95985** – cacheWrite spiral on bootstrap file writes.
- **#96007** – Discord multi-part reply truncates after error text.
- **#96064** – IRC plugin creates ghost nicks with underscore suffix.
- **#96020** – `sessions_send` delivery.status never updates from "pending".

Several bugs were closed today, including [#95760](https://github.com/openclaw/openclaw/issues/95760) (NVIDIA Build provider stream cut) and [#76729](https://github.com/openclaw/openclaw/issues/76729) (Feishu compaction drop). The resolution cadence is strong, but the high number of open P1 issues suggests stability is still a top concern.

## 6. Feature Requests & Roadmap Signals

Today’s activity reveals several clear user-desired features:

- **[#88032 – Telegram quote/reply context as first-class durable contract](https://github.com/openclaw/openclaw/issues/88032)**  
  *P2 Diamond Lobster, needs product decision.* Users want the quote/reply handling to be a stable inbound contract instead of split across prompt and runtime patches. This has been patched locally by multiple users.

- **[#71712 – Agent-facing scheduling API with non-forgeable provenance](https://github.com/openclaw/openclaw/issues/71712)**  
  *P2, stale.* A roadmap feature that could unlock autonomous agent workflows. Blocked on product decision and security review.

- **[#84724 – Allow isolated cron runs to disable their current job](https://github.com/openclaw/openclaw/issues/84724)**  
  *P2 Diamond Lobster, needs security review.* A small but impactful enhancement for recurring job lifecycle management.

- **[#96056 – Add non-secret env list support to OpenShell sandbox config](https://github.com/openclaw/openclaw/issues/96056)**  
  *New today.* Users want to pass environment variables (API base URLs, model names) without marking them as secrets. A linked PR (#96073) already exists.

- **[#95752 – Collapse tool activity blocks in Control UI by default](https://github.com/openclaw/openclaw/issues/95752)**  
  *Enhancement, P3.* A UI polish request to reduce visual clutter from tool execution blocks.

**PRs hinting at next version features:**
- **#96062** (feat(copilot): mirror native plan and subagent events) – bringing Copilot parity with Codex for plan visibility.
- **#96068** (fix(acpx): consume acpx 0.11.1 model capability errors) – unblocks ACP harnesses that lack model selection support.
- **#96072** (perf(browser): index role snapshot references) – performance improvement in browser accessibility tree handling.

**Prediction for next minor release (2026.6.10+):** Likely includes the heartbeat reasoning fix (#95798), Telegram markdown splitting (#95903), ACPX model capability error consumption (#96068), and potentially the OpenShell env config (#96073). The Windows ACPX fix (#93465) may be deferred pending live repro.

## 7. User Feedback Summary

**Pain points voiced today:**
- **WebChat UI flakiness:** Users report assistant replies rendered above user messages and duplicated inbound messages ([#95566](https://github.com/openclaw/openclaw/issues/95566)). Another user reports that the `Default (...)` model label is not calculated from active agent config ([#77690](https://github.com/openclaw/openclaw/pull/77690)).
- **Plugin detection regression:** A 2026.6.9 user found that a valid extension plugin is rejected at startup with "plugin not found" ([#96046](https://github.com/openclaw/openclaw/issues/96046)).
- **Windows environment still fragile:** ACPX runtime is completely broken on Windows ([#93465](https://github.com/openclaw/openclaw/issues/93465)); installer-created scheduled task shows a visible console ([#89231](https://github.com/openclaw/openclaw/issues/89231), closed today).
- **Mirror mode creates recursive directories:** OpenShell >=0.0.37 causes nested previous-workspace directories when using mirror ([#96016](https://github.com/openclaw/openclaw/issues/96016)).
- **Config-change gateway restart loses in-flight turn:** A simple `openclaw channels add` while chatting can erase the user’s message ([#95141](https://github.com/openclaw/openclaw/issues/95141)).
- **Inter-session message status never updates:** `sessions_send` always returns `delivery.status = "pending"` even after successful delivery ([#96020](https://github.com/openclaw/openclaw/issues/96020)).

**Satisfaction signals:**
- Several high-impact bugs were closed today (Feishu compaction, NVIDIA provider cut, heartbeat reasoning leak), showing the team is responsive to community reports.
- Users actively contribute PRs with proof and are engaging in long-lived feature discussions (e.g., #71712 scheduling API).

## 8. Backlog Watch

The following important issues and PRs have remained unanswered or stalled for an extended period. They are flagged for maintainer attention:

| Item | Status | Age | Why Important |
|------|--------|-----|---------------|
| [#71712](https://github.com/openclaw/openclaw/issues/71712) – Agent-facing scheduling API | `stale`, needs product & security decision | Since April 25 | Core roadmap feature for autonomous agents |
| [#81041](https://github.com/openclaw/openclaw/issues/81041) – `systemPromptHash` causes phantom claude-cli session restarts | `stale`, P1, needs product decision | Since May 12 | High impact on chat channel stability (Telegram, Discord, etc.) |
| [#85844](https://github.com/openclaw/openclaw/issues/85844) – Auto-update stale hashed bundle imports | P1, needs maintainer review and live repro | Since May 23 | Affects all users after auto-update; gateway may run with corrupted module graph |
| [#84724](https://github.com/openclaw/openclaw/issues/84724) – Cron disable capability | P2, needs security review | Since May 20 | Small change but blocked on security review |
| [#93968](https://github.com/openclaw/openclaw/issues/93968) – Silent config auto-patch breaks cron on macOS without Docker | P1, needs info & security review | Since June 17 | Two mechanisms silently write to `openclaw.json`, breaking cron jobs |
| [#79342](https://github.com/openclaw/openclaw/pull/79342) – Anti-sycophancy stress fixture suite | Open since May 8, `needs proof` | Large script/quality PR waiting for behavioral verification |
| [#78742](https://github.com/openclaw/openclaw/pull/78742) – Audit gateway restart attribution | Open since May 7, `needs proof` | Important for observability and debugging |
| [#78639](https://github.com/openclaw/openclaw/pull/78639) – Async-ify channels.status + messages.send | Open since May 6, waiting on author | Performance fix for WhatsApp / plugin discovery |

**Reminder:** Many of these issues carry `clawsweeper:needs-maintainer-review` or `clawsweeper:needs-product-decision` tags. A triage session to either close, accept, or schedule them would help reduce backlog pressure.

---

## Cross-Ecosystem Comparison

# Cross-Project Ecosystem Comparison Report
**Date:** 2026-06-23 | **Scope:** 12 Open-Source AI Agent Projects

---

## 1. Ecosystem Overview

The personal AI assistant open-source ecosystem is experiencing a **development boom**, with 7 of 12 projects showing high or very high activity today—collectively processing over 700 pull requests and 100 issues in 24 hours. The landscape is bifurcating into **core infrastructure projects** (OpenClaw, ZeroClaw) focused on runtime reliability and protocol standardization, and **application-layer projects** (CoPaw, NanoBot, Hermes Agent) racing to deliver desktop, mobile, and multi-platform user experiences. A shared urgency around **agent-facing scheduling, platform channel expansion, and reasoning block handling** reveals that the ecosystem is maturing from "can agents work?" to "how do we make them reliable and observable in production?" The rapid bug-fix cadence across projects (e.g., Groq provider fixed same-day in ZeroClaw, NanoBot patching v0.2.2 regressions within hours) signals a community that values stability even amid breakneck feature velocity.

---

## 2. Activity Comparison

| Project | Issues Updated | PRs Updated | Release Today | Health Assessment |
|---|---|---|---|---|
| **OpenClaw** | 39 (27 open) | 500 (84 merged) | No | Strong but strained: high merge velocity, but P1 Diamond Lobster bugs pending |
| **ZeroClaw** | 3 | 50 (7 merged) | No | Very high development velocity; critical Groq bug fixed same-day |
| **CoPaw** | 14 (8 open) | 50 (33 merged) | **v1.1.12.post2** | Healthy: rapid iteration with minor cron stability concerns |
| **Hermes Agent** | 21 (20 open) | 50 (5 merged) | No | Growing backlog: high engagement but low merge rate relative to open items |
| **NanoBot** | 8 | 33 (9 merged) | No | Healthy: quick regression fixes, responsive maintainers |
| **IronClaw** | 10 (8 open) | 33 (12 merged) | No | High velocity but critical reborn regression (#5139) threatens stability |
| **PicoClaw** | 2 (2 open) | 18 (6 merged) | No | Moderate, incremental progress |
| **NanoClaw** | 0 | 10 (1 merged) | No | Active dependency maintenance and feature PRs, low engagement |
| **LobsterAI** | 1 | 9 (3 merged) | No | Moderate; stale PR backlog of 81 days is concerning |
| **Moltis** | 0 | 1 (1 merged) | No | Low activity; single image tool addition merged |
| **NullClaw** | 0 | 1 (0 merged) | No | Quiet; single Matrix persistence PR pending |
| **TinyClaw** | 0 | 0 | No | **Inactive** — no activity observed |
| **ZeptoClaw** | 0 | 0 | No | **Inactive** — no activity observed |

**Health Scoring Key:**
- *Strong* = High merge velocity + responsive bug fixing + active maintainers (OpenClaw, ZeroClaw)
- *Healthy* = Good feature/bug balance, regressions addressed quickly (CoPaw, NanoBot, IronClaw)
- *Strained* = High engagement but low merge rate or growing critical backlog (Hermes Agent, LobsterAI)
- *Low Activity* = Minimal engagement or blocked progress (PicoClaw, NanoClaw, Moltis, NullClaw)
- *Inactive* = No detectable activity (TinyClaw, ZeptoClaw)

---

## 3. OpenClaw's Position

**Advantages over peers:**
- **Scale leader:** 500 PRs and 39 issues in 24 hours dwarfs every other project. The nearest competitor (ZeroClaw, CoPaw) each had 50 PRs—an order of magnitude less.
- **Mature severity taxonomy:** "Diamond Lobster" severity classification (shared with IronClaw) enables precise triage. No other project has equivalent granularity for P1 critical bugs.
- **Plugin ecosystem depth:** 10+ active platform channels (Discord, Telegram, Feishu, IRC, WebChat, Slack, etc.) vs. 2-5 for most peers. Telegram reply context (#88032) and IRC ghost nicks (#96064) show real-world deployment diversity.
- **Session integrity focus:** Thinking signature validation (#92201), compaction timeout handling (#92043), and cross-boot death-loop prevention (#95750) demonstrate investment in runtime correctness that rivals lack.

**Technical approach differences:**
- **ACPX runtime:** OpenClaw's proprietary runtime protocol (with Windows issues #93465) enables cross-platform agent execution but adds complexity. ZeroClaw and IronClaw use simpler spawn-based approaches.
- **Compaction-based state management:** Unique among peers. The 180s timeout debate (#92043) and Feishu compaction fix (#76729) reveal a design trade-off no other project has encountered.
- **"Clawsweeper" backlog system:** Tag-based triage (`needs-maintainer-review`, `needs-product-decision`) formalizes decision-making. CoPaw and NanoBot rely on PR comments for similar routing.

**Community size comparison:**
- OpenClaw's 500 PRs from ~84 merges suggests a contributor base of 40-60 active developers. CoPaw (~50 PRs, 33 merges) and ZeroClaw (~50 PRs, 7 merges) have smaller but more focused contributors.
- NanoBot's 21 new contributors in v0.2.2 indicates strong onboarding, while OpenClaw's backlog of `waiting on author` and `needs proof` PRs suggests reviewer bandwidth is the bottleneck.

---

## 4. Shared Technical Focus Areas

The following requirements emerge across **three or more projects**, signaling industry-wide consensus:

| Focus Area | Projects Affected | Specific Needs |
|---|---|---|
| **Agent-facing scheduling & cron reliability** | OpenClaw (#71712), CoPaw (#5064, #5398, #5235), IronClaw (#5131) | Non-forgeable provenance, pause/resume lifecycle, per-agent cron isolation |
| **Telegram/Discord rich message formatting** | OpenClaw (#88032), NanoBot (#4413, #4470), Hermes Agent (#10452) | Reply context as first-class contract, markdown normalization, streaming preview stability |
| **Thinking/reasoning block UI handling** | OpenClaw (#92201), NanoBot (#4465), CoPaw (#5416), ZeroClaw (#8219) | Signature validation on replay, hiding raw tags, reasoning_content extraction |
| **Tool call reliability & provider compatibility** | OpenClaw (#90404 ACPX), ZeroClaw (#8219 Groq), CoPaw (#5345 OpenAI compat), IronClaw (#5139 reborn) | Type-safe `tool_call_id`, provider fallback on reasoning rejection, multi-turn tool loop resilience |
| **Observability & cost tracking** | OpenClaw (#79401 structured incidents), ZeroClaw (#8065 trace_id/cost_usd), PicoClaw (#3156 token usage) | Per-turn cost emission, structured runtime incidents, OTLP integration |
| **Windows & macOS deployment stability** | OpenClaw (#93465 Windows ACPX), CoPaw (#5153 Tauri startup), Hermes Agent (#4707 launchd profile) | Daemon/service lifecycle, scheduled task console visibility, startup optimization |

**Underlying driver:** These clusters reveal that the ecosystem is converging on a **production readiness standard**—agents must execute reliably on schedule, communicate across platforms, handle reasoning models without leaking internals, and be observable for cost and debugging. Projects that solve these first (OpenClaw, ZeroClaw) will set de facto expectations.

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | ZeroClaw | CoPaw | NanoBot | Hermes Agent | IronClaw |
|---|---|---|---|---|---|---|
| **Primary Target User** | Self-hosted advanced operators (Slack, Telegram, IRC) | CLI power users, language model providers (Groq, NIM) | Desktop-first developers (Tauri, pywebview) | Casual Telegram/WebUI users | Multi-channel power users (Discord, Kanban) | Benchmark/enterprise (PinchBench, Reborn automation) |
| **Core Architecture** | Plugin gateway + ACPX runtime | CLI-centric + durable run-state store | Desktop app + SSE notifications | Lightweight agent + rich frontend | Gateway runner + provider abstraction | Reborn self-evolution engine + Trigger scheduler |
| **Release Cadence** | Stable (last release Jun 9) | Fast (daily PRs, no tag yet) | **Fast (patch today)** | Fast (v0.2.2 shipped, next patch imminent) | Fast but closed PRs lag | Moderate (no release today) |
| **Differentiator** | Session compaction & thinking integrity | Observability + per-agent env isolation | Mobile responsive + Todo UI | Heartbeat suppression & read-only sessions | Autopilot engine + Kanban integration | Skill extraction & benchmark tooling |
| **Key Bottleneck** | Maintainer backlog (clawsweeper) | Mutex poison crash (#8149) | Cron stability | Post-release regressions (v0.2.2) | 50 open PRs, only 5 merged | Reborn regression (#5139) |

**Architectural Trade-off Summary:**
- **Monolithic vs. Modular:** OpenClaw's plugin-based architecture allows maximum channel diversity but creates integration debt (e.g., IRC ghost nicks, Feishu compaction). CoPaw's Tauri desktop app provides a consistent UX but limits channel extensibility.
- **CLI-first vs. GUI-first:** ZeroClaw and NanoBot prioritize terminal interactions (doctor, sessions_send). Hermes Agent and IronClaw invest in WebUI and desktop dashboards.
- **Provider abstraction:** ZeroClaw and OpenClaw handle multiple LLM providers (Anthropic, Groq, NVIDIA, OpenAI) with their own compatibility layers. CoPaw and Hermes Agent show friction with non-OpenAI providers (custom OpenAI compat, Kimi K2).

---

## 6. Community Momentum & Maturity

**Tier 1: High Velocity (30+ PRs/day, rapid merges)**
- **OpenClaw** (500 PRs, 84 merges) — The ecosystem's engine. Despite maintainer strain, fixes for critical bugs (heartbeat reasoning, Feishu compaction) ship quickly. Backlog of product decisions is the main drag.
- **CoPaw** (50 PRs, 33 merges, patch release today) — Most efficient merge rate in the ecosystem. Balance of new features (terminal mode, SSE notifications) and bug fixes (dream task misfire) maintains momentum.
- **ZeroClaw** (50 PRs, 7 merges) — High open PR count suggests a "merge later" strategy. Same-day fix for Groq bug shows operational responsiveness.
- **Hermes Agent** (50 PRs, 5 merges) — Growing backlog: 20 open issues + 45 open PRs suggests maintainers are overwhelmed. Community engagement is high but unaddressed.

**Tier 2: Moderate Velocity (10-30 PRs/day)**
- **NanoBot** (33 PRs, 9 merges) — Healthy ratio of merges to open items. Quick regression fixes (Telegram, WebUI) after v0.2.2. Good newcomer onboarding (21 new contributors).
- **IronClaw** (33 PRs, 12 merges) — High merge rate but critical regression (#5139) threatens trust. Closure of long-standing feature requests (#4959, #4958) signals roadmap progress.

**Tier 3: Low Velocity (<10 PRs/day)**
- **PicoClaw** (18 PRs, 6 merges) — Steady but unspectacular. WhatsApp reconnection fix and security PRs (cross-site protection) show incremental quality improvements.
- **NanoClaw** (10 PRs, 1 merge) — Dependency maintenance dominates. Slack Socket Mode (#2837) is a notable exception.
- **LobsterAI** (9 PRs, 3 merges) — Moderate but burdened by 81-day stale PR backlog. The unresolved #1400 critical bug (upgrade crash) is a red flag.

**Tier 4: Stalled (<10 PRs/day, limited merges)**
- **Moltis** (1 PR, 1 merge) — Single feature addition; effectively maintenance mode.
- **NullClaw** (1 PR, 0 merges) — Matrix persistence fix pending; otherwise inactive.
- **TinyClaw, ZeptoClaw** — **Inactive** with no detectable activity.

**Maturity Assessment:**
The ecosystem is **polarizing into maintainers and innovators**. OpenClaw, CoPaw, and ZeroClaw are investing in infrastructure that others will consume (scheduling APIs, observability standards, provider adapters). Hermes Agent and IronClaw may face contributor burnout if merge velocity doesn't catch up to engagement. The inactive projects (TinyClaw, ZeptoClaw) likely lost maintainer attention or were absorbed into other stacks.

---

## 7. Trend Signals

### 1. Platform Ecosystem Expansion
Every active project is adding channels:
- OpenClaw: Discord, IRC, Telegram quote context
- NanoBot: Mattermost (#4459)
- Hermes Agent: Multi-bot Telegram (#10452), Kanban desktop
- ZeroClaw: Knowledge graph restore (#8182)
- PicoClaw: SimpleX/Tox request (#3093)

**Signal:** The "universal AI assistant" vision requires ubiquitous platform support. Development effort is shifting from core agent capability to **channel integration and message fidelity**.

### 2. Reliability-Driven Development
Today's bug reports cluster around **session integrity**:
- OpenClaw: Thinking signatures, compaction timeouts, cross-boot death-loops
- CoPaw: Cron scheduler halting, dream task failures
- ZeroClaw: Groq tool loop, mutex poison crash
- IronClaw: Reborn 0-LLM regression
- NanoBot: Telegram streaming flicker, WebUI thinking tag leakage

**Signal:** The ecosystem has passed the "can it work?" phase and entered the "does it work reliably in production?" phase. **Session state management** and **long-running task resilience** are the new moats.

### 3. Observability as Table Stakes
Multiple projects are adding structured logging, cost tracking, and telemetry:
- OpenClaw: Structured runtime incidents (#79401)
- ZeroClaw: trace_id/cost_usd (#8065), OTLP flush (#8146)
- PicoClaw: Per-turn token emission (#3156)
- IronClaw: Turn-state write convoy prevention (#5142)

**Signal:** Developers deploying AI agents in organizational settings demand **auditability and cost attribution**. This is moving from nice-to-have to must-have for enterprise adoption.

### 4. The Agent Scheduling Race
Three projects (OpenClaw #71712, CoPaw #5064/5398, IronClaw #5131) are independently building agent-facing scheduling APIs with **non-forgeable provenance** and **pause/resume lifecycles**.

**Signal:** **Autonomous agent workflows**—where agents create, manage, and monitor their own cron jobs—are the next frontier. The project that standardizes this (likely OpenClaw given its community size) will shape the ecosystem's future.

### 5. Provider Fragmentation Pain
Users are frustrated with provider-specific bugs:
- OpenClaw: Anthropic OAuth token validation (#95296)
- ZeroClaw: Groq `tool_call_id` serialization (#8219)
- CoPaw: Custom OpenAI compat lacks function calling (#5345)
- Hermes Agent: Z.AI false 429 (#47685)

**Signal:** Multi-provider support is a **differentiation risk**—users expect seamless switching but hit edge cases. The trend toward **provider-agnostic runtime abstractions** (OpenClaw's ACPX, ZeroClaw's provider trait) will accelerate.

### 6. Desktop/Mobile Convergence
CoPaw's mobile-responsive PRs (#5394, #5367) and Tauri auto-updater (#4669), Hermes Agent's Kanban desktop app (#41222), and NanoBot's PWA support (#4458) all point to **unified desktop-mobile agent experiences**.

**Signal:** The "terminal-only AI agent" phase is ending. **Rich graphical interfaces** with offline capability (PWA, Tauri) are becoming expected.

---

## Recommendations for Technical Decision-Makers

1. **If you need maximum platform coverage and runtime reliability today, choose OpenClaw**—but prepare for Windows friction and maintainer backlog delays.
2. **If you prioritize desktop UX and rapid iteration, CoPaw offers the best feature velocity** with acceptable cron stability trade-offs.
3. **If you're building provider-facing tooling or need advanced observability, ZeroClaw's PR trail is the most informative** for infrastructure patterns.
4. **If you're a new contributor, NanoBot's onboarding (21 new contributors in one release) is the most welcoming**—focus on Telegram optimization and dependency PRs.
5. **Monitor the Agent Scheduling API development** in OpenClaw (#71712)—it's likely to become an ecosystem standard that other projects will adopt.

The next 30 days will reveal whether OpenClaw's maintainer bottleneck is addressed (critical for ecosystem health), whether CoPaw stabilizes cron scheduling (critical for user trust), and whether the stalled projects (LobsterAI, NullClaw) regain momentum or fade into irrelevance.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest – 2026-06-23

## Today’s Overview

NanoBot saw high community and development activity today, with **8 issues** and **33 PRs** updated in the last 24 hours. The newly released **v0.2.2** (140 PRs merged, 21 new contributors) brings major durability improvements, but several regressions have been reported, particularly in Telegram message formatting and WebUI thinking-tag rendering. The maintainers responded quickly with fix PRs for the most critical bugs. Notably, two issues were closed (#4410 heartbeat suppression, #2305 reasoning toggle) and nine PRs were merged or closed. The project remains in a healthy, fast-moving state with strong community involvement.

## Releases

### v0.2.2 – Durability & Reliability Release
- **Merged PRs**: 140
- **New contributors**: 21
- **Headline improvements**:
  - WebUI conversations are now **transcript-segmented** (instead of one fragile file)
  - Forked chats preserve replies more reliably
  - General agent sturdiness enhancements
- **No breaking changes or migration notes** indicated in the release notes.
- **Known regressions** (post-release): Several bugs in Telegram channel and WebUI reasoning blocks, with fixes already in PRs (#4472, #4466).

**Link**: [v0.2.2 release](https://github.com/HKUDS/nanobot/releases/tag/v0.2.2)

## Project Progress

**Merged/closed PRs today (9 total):**
- [#4412](https://github.com/HKUDS/nanobot/pull/4412) – **Heartbeat notification suppression**: Ensures routine cron job responses are not sent when there is nothing to report (fixes #4410).
- [#4461](https://github.com/HKUDS/nanobot/pull/4461) – **v0.2.2 release announcement in README** (docs-only).
- Seven other PRs closed (not individually listed in the provided data, but implied by "merged/closed: 9").

**Key functional advances visible in open PRs:**
- **Optional feature enablement** across CLI and WebUI (#4396) – moves Bedrock to optional, adds shared discovery infrastructure.
- **Mattermost channel support** (#4459) – real-time messaging with WebSocket auto-reconnect and streaming.
- **Kimi Coding Plan provider** (#4464) – dedicated endpoint using Anthropic Messages API.
- **Node.js bump to v24** (#4460) – infrastructure update.
- **PWA support for WebUI** (#4458) – offline caching, home screen installation.
- **Read-only session support** (#4271) – disable LLM processing for pinned guides/FAQ.

## Community Hot Topics

1. **#4410 – Heartbeat sending unwanted messages after upgrade** (closed, 2 comments)  
   [Issue](https://github.com/HKUDS/nanobot/issues/4410)  
   User reported that routine cron jobs now always deliver LLM responses even when empty. Fix was merged today (#4412). This was a pain point for users relying on heartbeat for quiet monitoring.

2. **#4413 – Telegram Bot API 10.1 rich messages** (open, 2 comments)  
   [Issue](https://github.com/HKUDS/nanobot/issues/4413)  
   Enhancement request to convert standard markdown to Telegram’s new rich message format for better text rendering.

3. **#4465 – Raw `<thinking/>` tags visible in WebUI** (open, 1 comment, fix PR #4466)  
   [Issue](https://github.com/HKUDS/nanobot/issues/4465)  
   Thinking tags leak into chat UI, breaking user experience. Likely triggered by recent reasoning-block handling. PR #4466 normalizes these tags.

4. **#4470 – Telegram display bug after v0.2.2** (open, 0 comments yet, but reports two regressions)  
   [Issue](https://github.com/HKUDS/nanobot/issues/4470)  
   Line breaks ignored and message flickering due to streaming preview interaction. Fix PR #4472 submitted same day.

## Bugs & Stability

| Severity | Bug | Status | Fix PR |
|----------|-----|--------|--------|
| **High** | Telegram: line breaks not respected, message flickering (#4470) | Open (reported today) | [#4472](https://github.com/HKUDS/nanobot/pull/4472) – skip sendRichMessage when streaming preview exists |
| **High** | WebUI shows raw `<thinking>` tags (#4465) | Open (reported today) | [#4466](https://github.com/HKUDS/nanobot/pull/4466) – normalize thinking tags |
| **Medium** | Heartbeat sends unwanted messages after cron refactor (#4410) | **Closed** (fixed in #4412) | Merged |
| **Medium** | MCP server crash on reconnect (`attempted to exit cancel scope in different task`) (#4441) | Open (PR open) | [#4441](https://github.com/HKUDS/nanobot/pull/4441) – force-close generator on reconnect failure |
| **Medium** | Anthropic 400 due to duplicate tool_use IDs (#4442, fix PR #4444) | Open | [#4444](https://github.com/HKUDS/nanobot/pull/4444) |
| **Low** | Archived keys included in heartbeat targets (#4468) | Open | [#4468](https://github.com/HKUDS/nanobot/pull/4468) |
| **Low** | iOS Safari auto-zoom on WebUI composer (#4388, fix PR #4471) | Open | [#4471](https://github.com/HKUDS/nanobot/pull/4471) |

No critical crashes or data loss reported. All high-severity bugs have fix PRs already submitted within hours of reporting.

## Feature Requests & Roadmap Signals

- **Telegram Rich Messages** (#4413) – likely to be included in next minor release given the existing sendRichMessage code and fix PR #4472.
- **Dream skill merging** (#4467) – user wants Dream to update existing workspace skills instead of creating duplicates. PR #4469 submitted same day.
- **Kimi Coding Plan provider** (#4463) – PR #4464 already open, shows community demand for code-specific subscription models.
- **PWA support** (#4457) – PR #4458 open, low-hanging fruit for mobile UX.
- **Mattermost channel** (#4459) – PR open, adds significant enterprise integration.
- **Read-only sessions** (#4271) – PR open for over two weeks, likely merging soon.
- **Hide reasoning steps toggle** – issue #2305 was **closed today** as completed, indicating the feature is now implemented.

**Prediction for next version (v0.2.3):**  
Telegram rich message fixes, Dream skill merging, iOS zoom fix, and possibly PWA support. The optional feature enablement (#4396) may also land as it touches multiple subsystems.

## User Feedback Summary

**Pain points reported today:**
- Telegram users hit two regressions immediately after v0.2.2 upgrade: broken line breaks and disruptive flickering during streaming.
- WebUI users see raw thinking tags, especially those using models that emit `<thinking>` tags.
- Heartbeat cron users found their quiet setups broken; fix accepted quickly.
- Dream skill duplication frustrates users with active custom workflows – they want incremental improvements not fresh files.

**Positive signals:**
- The v0.2.2 durability improvements are acknowledged (transcript segmentation, forked chat reliability).
- Community contributions are active: 9 PRs merged today including rapid fixes for reported bugs.
- The project’s openness to new channels (Mattermost, Kimi Coding) is well received.

**Overall satisfaction / dissatisfaction:**  
Satisfaction remains high due to fast bug fixes, but the regressions in v0.2.2 cause temporary frustration. The team’s responsiveness suggests a quick patch release.

## Backlog Watch

The following older issues/PRs still need maintainer attention:

- [#3869](https://github.com/HKUDS/nanobot/pull/3869) – **DeepSeek message hardening** (open since May 16). Handling null/empty content for DeepSeek v4. PR has been open for over a month; important for provider stability.
- [#4271](https://github.com/HKUDS/nanobot/pull/4271) – **Read-only sessions** (open since June 10). A feature for administrative sessions; no recent maintainer review comments visible.
- [#4397](https://github.com/HKUDS/nanobot/pull/4397) – **User-attention hint injection** (open since June 18). Complex change to interrupt tool chains when user messages arrive mid-turn. Needs careful code review.
- [#4396](https://github.com/HKUDS/nanobot/pull/4396) – **Optional feature enablement** (open since June 18). Large PR touching CLI, WebUI, and plugin system. Core architectural change.

No issues appear to be completely abandoned, but the above PRs have been waiting for maintainer feedback for over 5 days each. Given the high activity today, they may be reviewed soon.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Project Digest: Hermes Agent – 2026-06-23

**Repository:** [NousResearch/hermes-agent](https://github.com/nousresearch/hermes-agent)  
**Data as of:** 2026-06-23 (UTC)

---

## 1. Today's Overview

Activity remains very high: 21 issues and 50 pull requests were updated in the last 24 hours, reflecting a project in active development. Of those, only 1 issue and 5 PRs were closed/merged — the vast majority remain open, indicating a growing backlog of work-in-progress. No new releases were cut today. The community is highly engaged, with several feature requests and bug reports garnering multiple comments and reactions. The project continues to evolve rapidly across gateway platforms, desktop UI, provider integrations (Z.AI, OpenAI Codex), and core agent capabilities like autopilot and memory providers.

---

## 2. Releases

**None** – No new releases were published in the data window.

---

## 3. Project Progress

**Closed/Merged PRs today (5 total):**  
Only one of the top-20 PRs is closed:  
- **[#51311](https://github.com/nousresearch/hermes-agent/pull/51311)** – *fix(gateway): runner takes over TTS when adapter won't auto-TTS for Discord DM voice_only* (P2, severity: risk-message-delivery). This patch ensures voice replies in Discord DMs are handled by the gateway runner when the base adapter cannot auto-TTS.  

The remaining 4 closed PRs (not shown in the top-20 list) likely include minor fixes and dependency bumps. No major features were merged today.

---

## 4. Community Hot Topics

### Most Discussed Issues (by comment count)

| Issue | Comments | Subject |  
|-------|----------|---------|  
| [#1955](https://github.com/nousresearch/hermes-agent/issues/1955) | 9 | feat: per-channel model and system prompt overrides for gateway platforms |  
| [#10452](https://github.com/nousresearch/hermes-agent/issues/10452) | 7 | Support multi Telegram bots for gateway routing and send_message |  
| [#27038](https://github.com/nousresearch/hermes-agent/issues/27038) | 6 | Bug: Codex Responses API rejects replayed assistant message items with long id fields |  
| [#34390](https://github.com/nousresearch/hermes-agent/issues/34390) | 5 | dashboard: add --allowed-hosts flag for reverse-proxy and Tailscale access |  
| [#47685](https://github.com/nousresearch/hermes-agent/issues/47685) | 5 | Z.AI GLM-5.2 Coding Plan returns 429/code 1305 when system prompt contains exact phrase "Hermes Agent" |  

### Most Reacted (👍)

- **Issue #1955** – 5 👍 – Users want granular model routing per channel, a strong signal for advanced multi-purpose deployments.  
- **Issue #10452** – 4 👍 – Multi-bot Telegram support is a common request for teams running separate bots.  
- **Issue #18430** – 5 👍 – Auto-rename Discord threads from generated titles would reduce manual friction.  
- **Issue #41222** – 6 👍 – Integrating the Kanban board into the desktop app is the most-wanted feature by reaction count.

**Analysis:** The community is pushing for higher-level platform orchestration (per-channel overrides, multi-bot, auto-naming) and better desktop integration (Kanban, reverse-proxy support). There is frustration with provider quirks (Z.AI, Codex) but also active troubleshooting.

---

## 5. Bugs & Stability

### New bugs reported today (2026-06-23)

| Issue | Severity | Description | Fix PR? |
|-------|----------|-------------|---------|
| [#51307](https://github.com/nousresearch/hermes-agent/issues/51307) | P2 | Dashboard virtual scrolling causes blank/empty history | None yet |
| [#51316](https://github.com/nousresearch/hermes-agent/issues/51316) | (default P3 from label? no priority label shown) | MCP tool snapshot misses stdio servers with slow startup in single-query mode | **Yes** – [#51322](https://github.com/nousresearch/hermes-agent/pull/51322) (open) |
| [#51320](https://github.com/nousresearch/hermes-agent/issues/51320) | P3 | Desktop chat input text disappears on send (macOS) | None yet |
| [#51321](https://github.com/nousresearch/hermes-agent/issues/51321) | P3 | Artifacts view shows incorrect file timestamps | None yet |
| [#51318](https://github.com/nousresearch/hermes-agent/issues/51318) | (no severity) | RTL/Bidi first-strong-character heuristic breaks mixed text | None yet |

### Existing bugs updated today (with ongoing activity)

- **P1** – [#4707](https://github.com/nousresearch/hermes-agent/issues/4707) – Cron under profile-scoped launchd gateway falls back to default `~/.hermes` (risk: security-boundary). No fix PR open.  
- **P2** – [#27038](https://github.com/nousresearch/hermes-agent/issues/27038) – Codex Responses API rejects long `id` fields. No fix PR yet.  
- **P2** – [#47685](https://github.com/nousresearch/hermes-agent/issues/47685) – Z.AI 429 on phrase "Hermes Agent". Investigation active.  
- **P2** – [#49200](https://github.com/nousresearch/hermes-agent/issues/49200) – Configured memory provider fails silently; falls back to built-in. No fix PR.  
- **P2** – [#30594](https://github.com/nousresearch/hermes-agent/issues/30594) – `hermes update` fails with PEP 668 on system Python. No fix PR.  
- **P2** – [#51312](https://github.com/nousresearch/hermes-agent/pull/51312) (PR, not issue) – Fix for `/steer` queued when no tool message exists. Open.

**Overall stability:** Several P2 bugs are lingering without associated fix PRs. The Codex and Z.AI provider bugs are particularly concerning for users relying on those backends. The desktop UI has two minor but annoying bugs reported today.

---

## 6. Feature Requests & Roadmap Signals

### Notable feature requests opened today

- [#51288](https://github.com/nousresearch/hermes-agent/issues/51288) – Add `HERMES_TUI_WS_WRITE_TIMEOUT_S` env var (TUI Gateway).  
- [#51319](https://github.com/nousresearch/hermes-agent/issues/51319) – **Add Doubao/Volcano Engine as native provider** (by wangyihang0222, with praise for Hermes).  
- [#51318](https://github.com/nousresearch/hermes-agent/issues/51318) – RTL/Bidi majority-based detection + manual toggle.  
- [#51314](https://github.com/nousresearch/hermes-agent/issues/51314) – Invalid test issue (closed).

### Other active feature requests with traction

| Issue | 👍 | Request |  
|-------|----|---------|  
| [#41222](https://github.com/nousresearch/hermes-agent/issues/41222) | 6 | Integrate Kanban board into desktop app |  
| [#18430](https://github.com/nousresearch/hermes-agent/issues/18430) | 5 | Auto-rename Discord threads from session titles |  
| [#1955](https://github.com/nousresearch/hermes-agent/issues/1955) | 5 | Per-channel model/system prompt overrides |  
| [#10452](https://github.com/nousresearch/hermes-agent/issues/10452) | 4 | Multi Telegram bots |  
| [#34390](https://github.com/nousresearch/hermes-agent/issues/34390) | 0 | Dashboard `--allowed-hosts` flag |  

### Predictions for next release

- The **autopilot PR [#49917](https://github.com/nousresearch/hermes-agent/pull/49917)** (engine-enforced goal-chasing with reviewer) is a large feature likely to merge soon.  
- **Multi-bot Telegram support** (#10452) and **per-channel overrides** (#1955) have been open since March/April and may land in a minor version.  
- **Doubao provider** (#51319) is a fresh request but has zero comments yet; quick implementation depends on upstream SDK availability.

---

## 7. User Feedback Summary

### Positive feedback
- **Thank you** from wangyihang0222 in [#51319](https://github.com/nousresearch/hermes-agent/issues/51319): *"Hermes Agent is an excellent tool… elegant design, greatly improved AI-assisted work efficiency."*  
- **High daily usage** reported in [#51320](https://github.com/nousresearch/hermes-agent/issues/51320): *"Thanks to the team … daily high-frequency usage, overall great experience."*

### Pain points expressed
1. **Input text disappears on send** ([#51320](https://github.com/nousresearch/hermes-agent/issues/51320)) – disrupts writing flow; macOS desktop.  
2. **Artifacts view wrong timestamps** ([#51321](https://github.com/nousresearch/hermes-agent/issues/51321)) – confusing file management.  
3. **Silent memory provider fallback** ([#49200](https://github.com/nousresearch/hermes-agent/issues/49200)) – users unknowingly lose external memory.  
4. **Z.AI false 429** ([#47685](https://github.com/nousresearch/hermes-agent/issues/47685)) – inexplicable rejection with "Hermes Agent" in system prompt.  
5. **Profile cron issue** ([#4707](https://github.com/nousresearch/hermes-agent/issues/4707)) – break profile isolation on macOS.  
6. **MCP slow startup in single-query** ([#51316](https://github.com/nousresearch/hermes-agent/issues/51316)) – workaround PR exists (#51322).  
7. **Dashboard blank scrolling** ([#51307](https://github.com/nousresearch/hermes-agent/issues/51307)) – hinders reviewing history.  
8. **Desktop build fails with codesign identity** ([#41499](https://github.com/nousresearch/hermes-agent/issues/41499)) – blocks local development on macOS.

Users are generally satisfied with the product breadth but frustrated by specific regressions and platform quirks.

---

## 8. Backlog Watch

### Issues needing maintainer attention (long-open, high priority)

| Issue | Created | Priority | Subject | Notes |
|-------|---------|----------|---------|-------|
| [#1955](https://github.com/nousresearch/hermes-agent/issues/1955) | 2026-03-18 | P3 | Per-channel model overrides | 9 comments, 5 👍 – very active but no PR. |
| [#4707](https://github.com/nousresearch/hermes-agent/issues/4707) | 2026-04-03 | **P1** | Cron falls back to default home (security) | No fix PR; risk: security-boundary. |
| [#10452](https://github.com/nousresearch/hermes-agent/issues/10452) | 2026-04-15 | P3 | Multi Telegram bots | 7 comments, 4 👍 – no PR yet. |
| [#27038](https://github.com/nousresearch/hermes-agent/issues/27038) | 2026-05-16 | P2 | Codex Responses API rejects long IDs | 6 comments, no fix PR. |
| [#18430](https://github.com/nousresearch/hermes-agent/issues/18430) | 2026-05-01 | P3 | Auto-rename Discord threads | 5 👍 – no PR. |
| [#27858](https://github.com/nousresearch/hermes-agent/pull/27858) | 2026-05-18 | P3 | Disable eager rate-limit fallback (PR) | Open, awaiting review. |
| [#27040](https://github.com/nousresearch/hermes-agent/pull/27040) | 2026-05-16 | P3 | Generic voice_server gateway (PR) | Open, awaiting review. |

### PRs with no recent maintainer activity

- [#29961](https://github.com/nousresearch/hermes-agent/pull/29961) (May 21) – Discord multi-bot infinite loop fix; open with undefined comments.  
- [#27328](https://github.com/nousresearch/hermes-agent/pull/27328) (May 17) – Feishu markdown tables rendering; open.  
- [#25268](https://github.com/nousresearch/hermes-agent/pull/25268) (May 13) – Hide Codex commentary messages; open.

**Recommendation:** Maintainers should prioritise the **P1 security issue #4707** and the **P2 bug #27038** (Codex) plus the popular feature requests #1955 and #18430 to reduce community frustration. The many open PRs (50 total, only 5 closed today) suggest a need for dedicated review cycles.

---

*Generated by Hermes Agent Project Analysis Tool on 2026-06-23.*

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest – 2026-06-23

## Today's Overview

The project saw moderate activity over the past 24 hours, with 2 issues updated (both open) and 18 pull requests updated (6 merged/closed, 12 open). No new releases were published. The merged PRs were largely focused on dependency bumps, bug fixes, and a minor feature enhancement for skill installation instructions. Open PRs indicate ongoing work on WhatsApp reliability, Android ADB tooling, sandbox path handling, and security hardening (cross-site request protection). Community engagement remains steady, with one new bug report about task repetition and an older feature request gaining traction.

## Releases

*No new releases this period.*

## Project Progress

Six pull requests were merged or closed in the last 24 hours, reflecting both dependency maintenance and targeted fixes:

- **#3162** – `fix(whatsapp): add reconnection and async message processing` – Resolves automatic WebSocket disconnection issues by introducing goroutine-based message handling, pong handlers, read deadlines, and exponential backoff reconnection. *Merged*.
- **#3053** – `fix(evolution): add ok check for LoadOrStore type assertion in lockStoreFile` – Prevents a panic from untyped assertions. *Merged*.
- **#3091** – `fix(openai_compat): add ok check for native_search type assertion` – Ensures a non-boolean value does not silently disable native search. *Merged*.
- **#3101** – `build(deps-dev): bump vite from 8.0.13 to 8.0.16` – Frontend dependency update. *Merged*.
- **#3105** – `build(deps-dev): bump eslint from 10.2.1 to 10.4.1` – Frontend dependency update. *Merged*.
- **#3152** – `add installation instructions to picoclaw skills search` – Enhances the `picoclaw skills search` output by spelling out how to install a skill. *Merged*.

## Community Hot Topics

The most active discussions center on a requested integration and a newly reported bug:

- **Issue #3093** – [“I need SimpleX or tox”](https://github.com/sipeed/picoclaw/issues/3093)  
  *Author: Damian-o2* | *Comments: 3* | *👍: 1*  
  OP requests adding a gateway for SimpleX, Wire, or Tox messaging. This is a long-standing feature request (created June 10) with modest community support. The need for alternative chat protocol bridges reflects a desire for decentralized, privacy-focused communication channels.

- **Issue #3159** – [“经常重复任务” (Frequent task repetition)](https://github.com/sipeed/picoclaw/issues/3159)  
  *Author: oKatTjC* | *Created: 2026-06-23* | *Comments: 0*  
  The user reports that after asking for “today’s US news” then “today’s French news,” the second response repeats the US news task before answering about France. This indicates a potential bug in conversation context or tool execution ordering, possibly related to session history or concurrent task handling.

## Bugs & Stability

One new bug was reported today, and several open PRs aim to improve reliability:

- **High Severity:** [#3159](https://github.com/sipeed/picoclaw/issues/3159) – **Task repetition bug** – AI model (deepseek-v4-flash-free) re-executes previous tasks instead of using context. No fix PR exists yet. Could be related to tool invocation logic or token stream handling.
- **Medium Severity:** [PR #3161](https://github.com/sipeed/picoclaw/pull/3161) – `fix(exec): keep deny patterns active for custom allow rules` – Ensures that `exec` deny patterns are not bypassed when a command matches a custom allow rule. This prevents potential security holes (e.g., allowing `jq` payloads that read environment variables). Open.
- **Medium Severity:** [PR #3160](https://github.com/sipeed/picoclaw/pull/3160) – `fix(auth): reject cross-site launcher setup requests` – Adds browser provenance checks to prevent CSRF-like attacks on first-run dashboard password setup. Open.
- **Low Severity (merged):** [#3162](https://github.com/sipeed/picoclaw/pull/3162) – WhatsApp reconnection fix (merged today).
- **Low Severity (merged):** [#3053](https://github.com/sipeed/picoclaw/pull/3053) – Type assertion panic fix in `lockStoreFile`.
- **Low Severity (merged):** [#3091](https://github.com/sipeed/picoclaw/pull/3091) – Silent search disable fix in OpenAI compat.

## Feature Requests & Roadmap Signals

User-generated feature requests remain focused on expanding communication channel support and improving interaction reliability:

- **#3093** – Support for **SimpleX / Wire / Tox** gateways. If maintainers see value in decentralized protocols, this could land in a future minor release (e.g., v0.3.x).
- **#3157** – **Experimental Android ADB remote operations tool** (PR open). This suggests interest in smartphone automation, which may become a built-in tool if stabilized.
- **#3156** – **Per-turn LLM token usage emission** (PR open). A “Pico channel” feature enabling downstream consumers to track token consumption per conversation. Likely to be merged soon as it adds useful observability.
- **#3118** – **Remote Pico WebSocket mode** for the agent (PR open). Adds `--remote ws://...` flag to connect agent to a remote Pico server. Could be part of a broader multi-agent or cloud deployment capability.

## User Feedback Summary

- **Pain points:** The single new bug report (#3159) highlights frustration with the AI repeatedly performing unrelated tasks before answering a new question—likely a context or session management issue. The user is on Debian 13, using the Web UI with a deepseek model.
- **Use cases:** Activity centers around news summarization across countries, but the bug degrades the experience.
- **Satisfaction:** No explicit dissatisfaction beyond the bug. The three comments on #3093 show moderate interest but no heated debate. Overall, the project appears stable with incremental improvements.

## Backlog Watch

Several open PRs and issues require maintainer attention:

- **Open stale PRs (unmerged for >7 days):**
  - [#3115](https://github.com/sipeed/picoclaw/pull/3115) – Fix inline data URL media extraction (critical for session history corruption). Last updated June 12. *Stale*.
  - [#3131](https://github.com/sipeed/picoclaw/pull/3131) – Add ok checks for tool schema type assertions. Last updated June 15. *Stale*.
  - [#3128](https://github.com/sipeed/picoclaw/pull/3128) – Ignore body.Close errors in search providers. Last updated June 15. *Stale*.
  - [#3104](https://github.com/sipeed/picoclaw/pull/3104), [#3100](https://github.com/sipeed/picoclaw/pull/3100), [#3103](https://github.com/sipeed/picoclaw/pull/3103) – Dependabot PRs for shadcn, plugin-react, typescript-eslint. All marked `[stale]`. Should be reviewed and merged to keep frontend dependencies current.

- **Long-unanswered issues:**
  - [#3093](https://github.com/sipeed/picoclaw/issues/3093) (feature request, created June 10) – No maintainer response. Acknowledgment or status update would be helpful.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

## NanoClaw Project Digest — 2026-06-23

### 1. Today's Overview
The project saw **very high PR activity**, with **10 pull requests updated** in the last 24 hours (9 open, 1 merged). No new issues were filed, and no releases were cut. The day was dominated by a coordinated dependency upgrade of the Chat SDK to **4.29.0** across three registry branches (`main`, `channels`, `providers`), alongside substantive feature work including Slack Socket Mode support and approval-card improvements. Overall health remains strong, with steady feature development and systematic dependency maintenance.

### 2. Releases
*No new releases today.*

### 3. Project Progress
Only one PR was **merged/closed** today:

- **#2833 – `Feat/hook surface guard`** (author: javexed)  
  *Merged.* Adds guard logic to the hook surface. The PR was marked as following contribution guidelines.  
  [PR #2833](https://github.com/nanocoai/nanoclaw/pull/2833)

Notable open PRs that advanced today (updated, not yet merged):

- **#2834** – `chore(deps): move chat SDK + channel-adapter pins to 4.29.0`  
  Core upgrade of the `main` install surface to Chat SDK 4.29.0, ensuring type compatibility between channel adapters and the bridge.  
  [PR #2834](https://github.com/nanocoai/nanoclaw/pull/2834)

- **#2837** – `feat(slack): Socket Mode — adapter + guided setup`  
  Adds outbound WebSocket connectivity for Slack using `SLACK_APP_TOKEN`, enabling local dev without a public endpoint.  
  [PR #2837](https://github.com/nanocoai/nanoclaw/pull/2837)

- **#2832** – `feat(approvals): reject with reason`  
  Introduces a third button on approval cards that lets the approver attach a one-line reason, relayed back to the requesting agent.  
  [PR #2832](https://github.com/nanocoai/nanoclaw/pull/2832)

- **#2830** – `fix(setup): reap dead peer service registrations whose binary is gone`  
  Cleans up orphaned launchd/systemd units after a checkout is deleted, preventing infinite restart loops.  
  [PR #2830](https://github.com/nanocoai/nanoclaw/pull/2830)

### 4. Community Hot Topics
Despite zero issues or PR comments today, several open PRs signal active community contributions:

- **Slack Socket Mode (#2837)** — Highly anticipated by developers behind NAT or firewalls. The “no public endpoint” promise is a significant quality-of-life improvement for local debugging.  
  [PR #2837](https://github.com/nanocoai/nanoclaw/pull/2837)

- **Approvals with reason (#2832)** — Addresses a frequent pain point where agents receive a bare “declined” with no context. Adding a reason field improves agent adaptability.  
  [PR #2832](https://github.com/nanocoai/nanoclaw/pull/2832)

- **Container performance (#2771)** — Proposes `--shm-size=1g` and `--init` for agent containers to avoid Chromium crashes. This PR received updates today after being opened a week ago, suggesting ongoing maintainer interest.  
  [PR #2771](https://github.com/nanocoai/nanoclaw/pull/2771)

### 5. Bugs & Stability
**No new bugs were reported today.** However, two PRs directly improve stability:

- **#2830** (fix) addresses a real-world accumulation of orphaned service registrations that cause OS-level spam. Severity: **Medium** (annoying but not blocking).  
  [PR #2830](https://github.com/nanocoai/nanoclaw/pull/2830)

- **#2771** (perf) mitigates Chromium crashes in containerised agents due to insufficient shared memory. Severity: **High** (can cause agent failures under load).  
  [PR #2771](https://github.com/nanocoai/nanoclaw/pull/2771)

### 6. Feature Requests & Roadmap Signals
Several long-running feature PRs indicate likely roadmap items:

- **Slack Socket Mode (#2837)** — Likely to land in the next minor release given its high community value and the fact that it builds on existing adapter infrastructure.
- **IMAP/SMTP email integration (#1235)** — Still open after 3 months, but remains the most-requested channel addition. The PR exposes 6 MCP tools and could appear in a follow-up release once dependency upgrades settle.
- **CLI-derived dashboard skill (#2795)** — A utility skill that reads CLI output, submitted as a “follows-guidelines” PR. Low complexity, may be merged soon.

### 7. User Feedback Summary
No explicit user feedback (comments, reactions) was recorded today. Inferred pain points from open PRs:

- **Local development friction** – Slack Socket Mode (#2837) directly targets users who cannot expose a public endpoint.
- **Poor agent feedback on rejections** – (#2832) suggests agents currently have no way to learn *why* a module was declined.
- **Stale service registrations** – (#2830) points to cleanup issues when users delete or move their NanoClaw directory.

### 8. Backlog Watch
- **PR #1235 – IMAP/SMTP email integration** (opened 2026-03-18)  
  After 3 months with no recent updates, this is the oldest open feature PR. It represents substantial functionality and may need a rebase or renewed maintainer review.  
  [PR #1235](https://github.com/nanocoai/nanoclaw/pull/1235)

- **PR #2771 – Container performance** (opened 2026-06-15)  
  Updated today but still unmerged. Given its high severity impact on headless browser agents, maintainer attention is warranted.  
  [PR #2771](https://github.com/nanocoai/nanoclaw/pull/2771)

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-06-23

## 1. Today's Overview

The NullClaw project experienced minimal visible activity over the past 24 hours. No issues were updated, and no new releases were published. One pull request (PR #968) remains open, targeting a critical persistence bug in the Matrix channel that causes unnecessary initial syncs after every restart. The absence of merged changes or closed issues indicates a quiet development day, though the lone PR signals ongoing maintenance work. Overall project health appears stable but with low engagement; the long-term impact will depend on whether this PR and other open items receive maintainer attention soon.

## 2. Releases

*No new releases were published in the last 24 hours.*  
The latest release remains unspecified. Users should continue using the most recent tagged version available on the repository.

## 3. Project Progress

- **Merged/Closed PRs today:** 0  
- **Open PRs advanced/fixed:** 1

**PR #968** ([nullclaw/nullclaw#968](https://github.com/nullclaw/nullclaw/pull/968)) — *fix(matrix): persist next_batch across restart + test env isolation*  
  - **Author:** addadi  
  - **Status:** Open (created 2026-06-22, updated 2026-06-22)  
  - **Summary:** The PR addresses a bug where the Matrix channel’s `/sync` cursor (`next_batch`) was stored only in RAM, causing the `buildSyncUrl` function to omit the `&since=` parameter on restart, forcing an initial sync from the homeserver. The fix ensures the cursor is persisted (likely to disk or a database) and adds test environment isolation to prevent state leakage.  
  - **Significance:** This is a functional fix improving startup reliability for Matrix users by avoiding unnecessary data re‑fetching. No conflicts or discussions have been reported yet.

## 4. Community Hot Topics

With only one PR updated (PR #968) and no comments or reactions recorded in the metadata, there are no active discussions or highly engaged items today. The PR itself has received no reactions (👍: 0) and no comments, but its underlying issue—data loss across restarts—is a pain point for users relying on persistent Matrix sessions. If the PR gains momentum, it could become a focus for the community.

## 5. Bugs & Stability

- **Bug reported/being fixed:** (High severity)  
  **Matrix `next_batch` not persisted across restart**  
  - **Impact:** Every restart forces a full initial sync, wasting bandwidth and time; may also cause missed messages during the sync gap.  
  - **Status:** Addressed by open PR #968. No fix has been merged yet, so users running from `main` still experience this bug.  
  - **Mitigation:** Users on builds before this PR should ensure any Matrix sessions are re‑synced manually after restart, or use a persistent state mechanism if available.

No other bugs, crashes, or regressions were reported in the last 24 hours.

## 6. Feature Requests & Roadmap Signals

No feature requests or roadmap signals were observed in the latest data. The only actionable item is a bug fix, not a new feature. The maintainers may be focusing on stability improvements. Future releases might include the Matrix persistence fix (PR #968) as a high-priority patch.

## 7. User Feedback Summary

No user feedback, pain points, or satisfaction indicators were captured in the last 24 hours. The absence of new issues or comments suggests either low usage or no urgent problems beyond the known Matrix bug. Any user frustration related to restart syncs would likely be alleviated once PR #968 is merged.

## 8. Backlog Watch

- **No long-unanswered Issues or PRs** were identified in the last 24 hours.  
  The only open PR (#968) is still fresh (created <2 days ago) and does not yet require maintainer attention beyond normal review. No other backlog items were updated.

*Overall, NullClaw is in a quiet period with a single high-priority bug fix under review. Maintainers should proceed with merging PR #968 to eliminate the Matrix restart issue as soon as possible.*

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-06-23

## 1. Today’s Overview
The project saw high activity with 33 pull requests updated in the last 24 hours (12 merged/closed) and 10 issues updated (8 still open). Two long-standing feature issues—global auto-approve (#4959) and per-tool permissions (#4958)—were closed, signaling completed implementation. However, a serious regression in Reborn web/research tasks (#5139) is causing 0 LLM calls and turn timeouts on 21 of 147 PinchBench tasks, placing stability at risk. Work continues on automation pause/resume (#5131), Slack integration via WebUI (#5152), context management (#5149), and tool permissions UI (#5068). The merge queue is intermittently blocked by a flaky test (#5147).

## 2. Releases
No new releases in the last 24 hours. The latest version remains unchanged.

## 3. Project Progress
**Merged/closed PRs today (5):**
- [#5142 fix(turns): prevent turn-state write convoy](https://github.com/nearai/ironclaw/pull/5142) – Removes per-user turn-state write gate, replaces with versioned CAS + jittered backoff to avoid contention.
- [#5141 fix(triggers): complete once permanent failures](https://github.com/nearai/ironclaw/pull/5141) – Prevents `Once` triggers from re-firing after permanent pre-submission failures.
- [#5150 [codex] restore GSuite duplicate account fallback](https://github.com/nearai/ironclaw/pull/5150) – Restores GSuite resolver’s fallback for duplicate reusable accounts.
- [#5061 feat(reborn): skill extraction & self-evolution with activation controls](https://github.com/nearai/ironclaw/pull/5061) – Adds background skill distillation from successful turns, scoped write + safety scan, with activation controls.
- [#4712 Move Slack setup into WebUI](https://github.com/nearai/ironclaw/pull/4712) – (closed, superseded by #5152) Hard-cuts Slack TOML config down to `[slack].enabled`.

**Closed issues representing implemented features:**
- [#4959 Reborn global auto-approve setting with per-turn approval resolution](https://github.com/nearai/ironclaw/issues/4959) – Now closed, likely delivered through #5068 or related PRs.
- [#4958 Reborn per-tool permission model (always_allow/ask_each_time/disabled)](https://github.com/nearai/ironclaw/issues/4958) – Also closed, enabling tool-level permission states.

## 4. Community Hot Topics
- [#5139 reborn regression: web/research tasks hang at init (0 LLM calls)](https://github.com/nearai/ironclaw/issues/5139) – Most critical open issue. Discussed in 1 comment. Affects 14% of today’s benchmark tasks. Root cause lies in 10 commits between `2b2ccc55` and `704fcd43`.
- [#5151 Claude fails to create Reborn automation after trigger pause/resume tools exposed](https://github.com/nearai/ironclaw/issues/5151) – Reports `claude-sonnet-4-5` failing to call `builtin.trigger_create` after pause/resume tools became available. Likely exposes a model/tool-selection issue.
- [#5148 Turn scheduler heartbeat can self-deadlock while a run holds transition state](https://github.com/nearai/ironclaw/issues/5148) – A running turn may hang forever if scheduler heartbeat fires while executor holds async store lock. Reported during GitHub extension install flow.
- [#5147 Flaky test: trigger_poller_does_not_submit_turn_for_unpaired_actor](https://github.com/nearai/ironclaw/issues/5147) – 1-in-3 flaky failure on `main` blocking merge queue. No fix PR yet, but maintainers are aware.

## 5. Bugs & Stability
| Severity | Issue | Description | Fix PR |
|----------|-------|-------------|--------|
| **Critical** | [#5139](https://github.com/nearai/ironclaw/issues/5139) | Reborn regression: web/research tasks hang with 0 LLM calls, turn timeouts. Zeroed 21/147 tasks. | None yet (under investigation) |
| **High** | [#5148](https://github.com/nearai/ironclaw/issues/5148) | Scheduler heartbeat self-deadlock when executor holds transition state. | None |
| **High** | [#5147](https://github.com/nearai/ironclaw/issues/5147) | Flaky test intermittently blocks merge queue (removes PRs from queue). | None |
| **Medium** | [#5151](https://github.com/nearai/ironclaw/issues/5151) | Claude fails to create automation after trigger pause/resume tools exposed. | Possibly related to #5131 (still open) |
| **Ongoing** | [#4108](https://github.com/nearai/ironclaw/issues/4108) | Nightly E2E scheduled run failing repeatedly since May 27. Latest failure 2026-06-23. | None |

Two fix PRs merged today: #5142 (preventing turn-state write convoy) and #5141 (fixing Once trigger permanent failures). These address recently reported concurrency and trigger edge cases.

## 6. Feature Requests & Roadmap Signals
**New or active feature PRs/requests:**
- [#5131 automation pause/resume](https://github.com/nearai/ironclaw/pull/5131) – Core feature, still open. Likely to land in next release.
- [#5152 Move Slack setup into WebUI](https://github.com/nearai/ironclaw/pull/5152) – Replaces #4712; includes secret store integration and dynamic channel setup.
- [#5149 Context management – progressive tool disclosure](https://github.com/nearai/ironclaw/pull/5149) – Flag-gated, default off. Aims to reduce per-call tool schema tokens from ~25.8k to mitigate NEAR AI timeouts.
- [#5068 WebUI tool permissions + global auto-approve settings](https://github.com/nearai/ironclaw/pull/5068) – Still open; wires UI to durable stores so changes apply without restart.
- [#5144 Show NEAR AI default base URL in provider card](https://github.com/nearai/ironclaw/issues/5144) – UX improvement.
- [#5146 No button to deactivate an extension on Extensions page](https://github.com/nearai/ironclaw/issues/5146) – User-reported missing action.

**Prediction:** The next version will likely include automation pause/resume (#5131), Slack WebUI setup (#5152), context management (#5149), and tool permissions UI (#5068). The reborn regression (#5139) must be fixed before any release.

## 7. User Feedback Summary
- **Pain points:** Reborn tasks timing out with 0 LLM calls (#5139) – severe impact on daily benchmarks. Claude model failing to create automations (#5151) – affects automation workflows. Missing deactivate button for extensions (#5146) – reduces management efficiency. Provider card showing `None` for NEAR AI base URL (#5144) – confusing.
- **Satisfaction indicators:** Closure of auto-approve (#4959) and per-tool permissions (#4958) suggests these long-requested features are now available. Slack integration moving into WebUI (#5152) reduces TOML configuration burden.
- **Overall sentiment:** Development velocity is high, but stability regressions and blocker flaky tests erode user trust. The community is actively reporting edge-case failures, which indicates heavy testing and real usage.

## 8. Backlog Watch
| Item | Age | Impact | Notes |
|------|-----|--------|-------|
| [#4108 Nightly E2E failure](https://github.com/nearai/ironclaw/issues/4108) | 27 days | Blocks confidence in all E2E nightly runs. Multiple failures cited. | No assignee, no linked PR. Needs prioritization. |
| [#2863 Add Manifest as built-in LLM provider](https://github.com/nearai/ironclaw/pull/2863) | 62 days | Adds a new provider; low risk. | Stale PR from April. No recent updates. Maintainer review needed. |
| [#4804 Reborn operator log tail/follow](https://github.com/nearai/ironclaw/pull/4804) | 11 days | XL-size feature, still open. | No recent comments. Risk of bit-rot if not merged soon. |
| [#4860 Local service lifecycle backend](https://github.com/nearai/ironclaw/pull/4860) | 9 days | XL-size, wires systemd/launchd control. | Blocked waiting for persistence service? Needs attention. |
| [#5120 Unify gate declined semantics](https://github.com/nearai/ironclaw/issues/5120) | 1 day (new) | Non-urgent but touches multiple UX surfaces. | Maintainers should agree on terminology soon to avoid accumulation of variations. |

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-06-23

## 1. Today's Overview

Project activity is moderate, with 1 issue and 9 pull requests updated in the last 24 hours. Three PRs were merged/closed today, primarily addressing OpenClaw cron session management and legacy storage migration. No new releases were published. The most notable community attention remains on a critical startup bug (Issue #1400), which was updated today after months of inactivity, indicating ongoing user distress. Six open PRs have remained stale since early April, suggesting a maintenance backlog.

---

## 2. Releases

**No new releases** today. The latest release remains unknown.

---

## 3. Project Progress

Three pull requests were merged/closed today, all created and resolved on 2026-06-23:

- **#2190 (Closed)** — `fix(openclaw): sync cron run sessions`  
  Recognizes `agent:{agentId}:cron:{jobId}:run:{runId}` session keys and normalizes them to a stable per-agent cache key, preventing duplicate Cowork sessions for repeated runs.  
  [PR #2190](https://github.com/netease-youdao/LobsterAI/pull/2190)

- **#2189 (Closed)** — `fix(openclaw): migrate legacy cron storage on startup`  
  Detects legacy OpenClaw cron JSON/run-log storage before gateway startup and runs the official migration using a temporary cron config, then syncs `cron.store`.  
  [PR #2189](https://github.com/netease-youdao/LobsterAI/pull/2189)

- **#2188 (Closed)** — `Liuzhq/rlog`  
  (Summary not provided; tagged with areas including renderer, docs, main, openclaw, cowork; likely a logging-related feature or fix)  
  [PR #2188](https://github.com/netease-youdao/LobsterAI/pull/2188)

These merges advance the OpenClaw subsystem’s reliability for scheduled tasks and migration paths.

---

## 4. Community Hot Topics

The only issue with significant discussion is:

- **#1400 (Open, Stale)** — `4.1版本严重bug，网关反复启动失败，反复重启，无限循环！`  
  👤 danielmonlite | 📅 Created 2026-04-03 | 💬 6 comments | 👍 0  
  The user describes an upgrade from 3.30 to 4.1 causing an infinite restart loop, and a secondary problem where custom LLM (qwen3.5-plus) fails due to an unmet `web-extractor` dependency. The user left contact info (email / WeChat).  
  [Issue #1400](https://github.com/netease-youdao/LobsterAI/issues/1400)

**Underlying need:** Users require stable upgrades without breaking existing custom configurations, and clear migration guidance for LLM provider dependencies.

No other issues or PRs attracted more than 0 comments or reactions today.

---

## 5. Bugs & Stability

| ID | Severity | Description | Status | Fix PR? |
|----|----------|-------------|--------|---------|
| #1400 | **Critical** | 4.1 upgrade causes infinite gateway restart loop; custom LLM cannot be used because `web-extractor` requires `web-search` enabled. User reports system completely unusable. | Open (stale, updated today) | No fix PR linked. |
| (none) | | No other bugs reported in the last 24 hours. | | |

The #1400 issue remains the highest-priority stability threat. The user suggests a possible conflict with LobsterAI’s automatic configuration of qwen3.5 upon login, but even after logging out the problem persists.

---

## 6. Feature Requests & Roadmap Signals

No new feature requests were filed today. The following **long-standing open PRs** (all from April 2026) contain feature/fix work that may land in the next minor release:

- **#1404** — `feat(scheduledTasks): 定时任务创建界面时间控件优化`  
  Improves time picker (replaces native `<input type="time">` with a custom component) and replaces native `<select>` elements with themed custom dropdowns for cron configuration.  
  [PR #1404](https://github.com/netease-youdao/LobsterAI/pull/1404)

- **#1402** — `fix(cowork): keep all files from multi-select attachment picker`  
  Solves a bug where only the last selected file was kept when choosing multiple files in one dialog.  
  [PR #1402](https://github.com/netease-youdao/LobsterAI/pull/1402)

- **#1406** — `fix(scheduled-task): fallback notify channel list when IM filter is empty`  
  Prevents empty dropdown for notification channels when IM settings exist but no platform is toggled on.  
  [PR #1406](https://github.com/netease-youdao/LobsterAI/pull/1406)

- **#1401** — `fix: 修复请求安全性问题`  
  Replaces `Math.random()` with `crypto.randomUUID()` for SSE request IDs to prevent predictable stream IDs.  
  [PR #1401](https://github.com/netease-youdao/LobsterAI/pull/1401)

These PRs indicate ongoing work to polish the scheduled tasks UI, file attachment UX, i18n, and security. None have received maintainer review or merge yet.

---

## 7. User Feedback Summary

The only direct user feedback in the last 24 hours comes from Issue #1400:

- **Pain points:**  
  - Upgrade from 3.30 to 4.1 fails catastrophically (infinite restart loop).  
  - Custom LLM provider (qwen3.5-plus) cannot be used because `web-extractor` is a required subsystem that is blocked when `web-search` is not enabled.  
  - The user expresses frustration: “反正现在是彻底瘫痪了！” (Now it is completely paralyzed).

- **Satisfaction:** Very low for the affected user. They are seeking direct help via email and WeChat.

- **Use case:** The user appears to be a self-hosted operator who relies on a specific third-party LLM and expects backward compatibility with the upgrade.

---

## 8. Backlog Watch

Several **stale items** (last updated >2 months ago, or updated only recently without resolution) require maintainer attention:

| Item | Type | Last Updated | Days Stale | Risk |
|------|------|--------------|------------|------|
| [#1400](https://github.com/netease-youdao/LobsterAI/issues/1400) | Issue (critical bug) | 2026-06-23 | 81 (created Apr 3) | **High** – user still unable to use product |
| [#1401](https://github.com/netease-youdao/LobsterAI/pull/1401) | PR (security fix) | 2026-06-23 (updated, but no review) | 81 | Medium – security improvement unmerged |
| [#1402](https://github.com/netease-youdao/LobsterAI/pull/1402) | PR (cowork bug fix) | 2026-06-23 | 81 | Medium – user-facing attachment bug |
| [#1403](https://github.com/netease-youdao/LobsterAI/pull/1403) | PR (i18n fix) | 2026-06-23 | 81 | Low – cosmetic, but affects multilingual users |
| [#1404](https://github.com/netease-youdao/LobsterAI/pull/1404) | PR (scheduled tasks UI) | 2026-06-23 | 81 | Medium – UX improvement |
| [#1406](https://github.com/netease-youdao/LobsterAI/pull/1406) | PR (scheduled task channels) | 2026-06-23 | 81 | Medium – missing channel list |
| [#1277](https://github.com/netease-youdao/LobsterAI/pull/1277) | PR (deps update) | 2026-06-22 | 82 | Low – dependency bump (electron 40→42) |

**Recommendation:** The project should prioritize triaging Issue #1400 and the associated stale PRs to prevent user churn and security debt.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest – 2026-06-23

## 1. Today's Overview
Project activity was minimal over the past 24 hours, with zero new issues and no releases. A single pull request (#215) was updated and merged/closed, reflecting ongoing refinement of the tooling layer. The lack of open issues or bug reports suggests the codebase is currently stable, though community engagement appears low. Overall, the project is in a maintenance phase with incremental feature integration.

## 2. Releases
**No new releases** since the last digest. The latest available release remains unchanged.

## 3. Project Progress
- **Merged/Closed PRs (1 item):**
  - **#215 – feat(tools): add send_image tool for channel image delivery**  
    *Author:* [maximilize](https://github.com/maximilize) · *Created:* 2026-02-23 · *Merged/Closed:* 2026-06-23  
    *Link:* [PR #215](https://github.com/moltis-org/moltis/pull/215)  
    **Summary:** Introduces a `send_image` tool that allows skills to send local image files (PNG, JPEG, GIF, WebP) to channel targets such as Telegram. It reuses the existing screenshot pipeline, returning a `data:` URI in the `screenshot` key for automatic pickup by the chat runner. Supports an optional `caption` parameter. This feature extends Moltis’s multimodal capabilities for chat‑based interactions.

## 4. Community Hot Topics
No issues or PRs generated significant discussion in the last 24 hours. The only recent PR (#215) had zero comments and reactions. Community engagement is currently quiet.

## 5. Bugs & Stability
No bugs, crashes, or regressions were reported in the past 24 hours. The project appears stable, with no open issues or fix PRs outstanding.

## 6. Feature Requests & Roadmap Signals
No explicit feature requests were filed. The merger of the `send_image` tool (PR #215) indicates an ongoing focus on expanding the tool ecosystem for richer channel interactions. Given the pattern of tool additions (previous features like screenshot), future releases may include further media‑handling tools (e.g., `send_video`, `send_audio`) or enhanced captioning support.

## 7. User Feedback Summary
No user feedback (comments, reactions, or new issues) was recorded in the last 24 hours. The lack of complaints or feature requests suggests either low usage or satisfaction with the current feature set.

## 8. Backlog Watch
No long‑unanswered issues or PRs currently require maintainer attention. The only closed PR (#215) was merged today; no other items are pending review. The project backlog is effectively empty.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

Here is the **CoPaw Project Digest** for **2026-06-23**, generated from the provided GitHub data.

---

## CoPaw Project Digest — 2026-06-23

### 1. Today's Overview

The CoPaw (QwenPaw) project is in a period of **very high activity and rapid iteration**. In the last 24 hours, **50 pull requests** were updated (33 merged/closed), **14 issues** were active (8 still open), and a new patch release (**v1.1.12.post2**) was published. Development is focused on three key areas: **fixing critical cron/dream task scheduling bugs**, **improving desktop and mobile client responsiveness**, and **preparing the major v2.0 architecture merge**. Community engagement is strong, with users actively reporting regressions and requesting specific integrations (e.g., Anthropic-compatible providers, mobile WebUI). The project's health is good, with maintainers responding quickly to bug reports, though a cluster of cron-related issues suggests a stability regression that requires close monitoring.

### 2. Releases

- **Version:** **v1.1.12.post2**
- **Type:** Patch release
- **Key Changes (from PRs):**
    - `fix`: Navigate to new chat after deleting the current session.
    - `feat`: Enhanced file preview to support relative paths in the console.
    - `fix`: Resolved an issue where code for image generation or certain files was not visible.
- **Breaking Changes:** None. This is a minor bug-fix and enhancement release.
- **Migration Notes:** Standard upgrade via `pip install --upgrade qwenpaw`. No configuration changes required.
    - **Link:** [v1.1.12.post2 Release](https://github.com/agentscope-ai/QwenPaw/releases/tag/v1.1.12.post2)

### 3. Project Progress

Today saw significant forward momentum across several tracks, with **33 PRs merged or closed**. Key advancements include:

- **Terminal Coding Mode (CLI):** A major new feature, `qwenpaw terminal`, was merged. It provides an interactive coding-mode terminal that connects to a running QwenPaw daemon, enabling file manipulation and subprocess management. (PR [#5304])
- **Real-Time Notifications:** SSE-based push notifications were added for the console channel, replacing polling latency with sub-50ms delivery for events like cron and API triggers, plus an optional audio beep. (PR [#5331])
- **Memory Search Enhancement:** An optional recency-aware ranking step for `memory_search` on daily notes was implemented, using exponential temporal decay to boost more recent notes. (PR [#5325])
- **Plan Execution Panel (todo_write):** A native `todo_write` tool and frontend panel were added to display multi-step agent task progress in real-time. (PR [#5323])
- **Desktop Client Improvements:**
    - **System Tray Minimization:** The desktop client now minimizes to the system tray on close instead of quitting. (PR [#5326])
    - **Graceful Shutdown:** A new API endpoint for the Tauri sidecar ensures the backend shuts down gracefully before the desktop app closes. (PR [#5432])
- **Mobile Responsive Layout:** Initial PRs were merged to fix mobile layouts for the **Plugin Manager** (PR [#5394]) and **Security settings** pages (PR [#5367]).
- **V2.0 Foundation:** A major merge from `main` (v1.1.12) into the `dev/agentscope2.0` branch was completed, porting 170 commits into the new architecture. (PR [#5412])

### 4. Community Hot Topics

The most active discussions reveal pain points around **cron reliability** and **model compatibility**.

- **Issue #5064 (Closed): [Bug] Agent-generated cron tasks don't trigger.** This issue had **12 comments**. The user reported that tasks created by an agent appeared in the UI but never executed. While closed, the underlying problem of cron scheduling integrity remains a hot topic, as seen in related open issues.
    - [Link to Issue #5064](https://github.com/agentscope-ai/QwenPaw/issues/5064)
- **Issue #5345 (Open): [Bug] Custom OpenAI-compatible providers don't support function calling.** With **6 comments**, users are reporting that providers like OMLX (which implement the full OpenAI API) only return text and fail to invoke tools when added as a custom provider. This highlights a gap in the provider abstraction layer.
    - [Link to Issue #5345](https://github.com/agentscope-ai/QwenPaw/issues/5345)
- **Issue #5398 (Closed): [Bug] Cron scheduler stops dispatching enabled jobs.** With **5 comments**, this issue mirrors the challenges in #5064, where the cron scheduler for a specific agent stopped firing even though the process remained alive.
    - [Link to Issue #5398](https://github.com/agentscope-ai/QwenPaw/issues/5398)

### 5. Bugs & Stability

Today's bug reports indicate a **moderate stability concern**, primarily around the cron/dream scheduler subsystem.

| Severity | Bug | Issue/PR | Status |
| :--- | :--- | :--- | :--- |
| **Critical** | **Dream Task Execution Failed:** All three agents failed to execute a nightly summary task. | [#5402] | **Closed** (Fixed by PR #5426 - set default misfire time) |
| **High** | **Cron Scheduler Stops Dispatching:** Two separate reports of the cron scheduler halting for enabled jobs without crashing the app. | [#5064], [#5398], [#5235] | **Closed** (Root cause under investigation; PR #5347 may help) |
| **Medium** | **Severe Lag Switching Agents/Windows:** Significant UI stuttering when switching chats or agents. | [#5421] | **Open** |
| **Medium** | **DeepSeek Agent Hangs During "Thinking":** Agent gets stuck and requires manual intervention ("stop" then "continue"). | [#5328] | **Open** |
| **Low** | **Internal Server Error on Fresh Install:** New installation via `pip` fails with a `get_remote_addr` exception. | [#5379] | **Open** |
| **Low** | **Thinking Content Not Displayed:** Some models output text in `thinking`/`reasoning_content` fields, leaving the main `content` empty and invisible to users. | [#5416] | **Open** |

**Risk Assessment:** The cluster of cron/dream issues is the most pressing. While PR #5426 fixes the immediate "dream" failure by adjusting misfire grace periods, the root cause of the scheduler stopping entirely (#5398) may need a more robust fix.

### 6. Feature Requests & Roadmap Signals

Several feature requests point to clear user needs that are likely to be prioritized:

- **Mobile-First Experience (High Priority):** Issue #4635 (now closed) requested a mobile-friendly WebUI. This is being actively addressed, with mobile-responsive layout PRs for the Plugin Manager (PR [#5394]) and Security page (PR [#5367]) merged today.
    - **Prediction:** A full mobile-responsive console UI is likely in the **next minor version (v1.2.0)** .
- **Kimi Coding Model Support:** Issue #5427 requests configuration support for Kimi K2 Code, which uses an *Anthropic-compatible* API endpoint, not just OpenAI.
    - **Prediction:** This is a narrow but requested integration. A new provider or provider configuration option for Anthropic API format support could appear in **v1.2.0**.
- **Recency-Aware Memory Search:** Issue #5316 (now closed/merged) requested better ranking for daily notes.
    - **Prediction:** Already merged into `main`. Will be available in the **next patch**.
- **Discord Streaming Responses:** An older request (Issue #1296) from March for streaming responses in the Discord channel remains open.
    - **Prediction:** Not a priority now, but the SSE notification work (PR #5331) may lay the groundwork for this in a future release.

### 7. User Feedback Summary

User sentiment is a mix of **satisfaction with rapid innovation** and **frustration with regressions in core functionality**.

- **Satisfaction:** Users are engaged and requesting advanced features (custom providers, memory ranking, mobile access), indicating they see value in the platform and are investing time in it.
- **Pain Points:**
    - **Reliability:** The top pain point is the unpredictability of cron/dream task execution. Users are losing trust in the scheduling system.
    - **Model Integration Friction:** Users are frustrated that known-working providers (OMLX, Kimi) require manual hacks or do not work at all, especially for function calling.
    - **UI Performance:** Users on desktop are experiencing significant lag and stutter when managing multiple agents, impacting the core workflow.
    - **Configuration Complexity:** New users are hitting "Internal Server Error" on a fresh install, which creates a poor first-time experience.

### 8. Backlog Watch

The following items have been open for some time and may need maintainer attention:

- **Feature: Replicate Tauri startup optimization to pywebview (PR #5153):** Opened on June 12. This PR aims to optimize the startup time of the Windows desktop client. It has not yet been merged.
    - [Link to PR #5153](https://github.com/agentscope-ai/QwenPaw/pull/5153)
- **Feature: Add Tauri auto-updater (PR #4669):** Opened on May 25. This is a significant feature for desktop distribution. It is still under review, possibly waiting for the v2.0 merge to resolve architectural differences.
    - [Link to PR #4669](https://github.com/agentscope-ai/QwenPaw/pull/4669)
- **Feature: CLI `cron update` command (PR #5210):** Opened on June 15. An important workflow improvement to avoid the delete-recreate cycle for cron jobs. Still under review.
    - [Link to PR #5210](https://github.com/agentscope-ai/QwenPaw/pull/5210)
- **Bug: Drop invalid jobs on startup (PR #5347):** This PR provides a long-term fix for cron corruption by cleaning up invalid entries on workspace start. It directly addresses the root cause of several cron issues. It has been open for 3 days and should be prioritized for merge.
    - [Link to PR #5347](https://github.com/agentscope-ai/QwenPaw/pull/5347)

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

## ZeroClaw Project Digest — 2026-06-23

### 1. Today's Overview
ZeroClaw saw **very high activity** over the past 24 hours, with **50 pull requests** updated (43 open, 7 merged/closed) and **3 issues** touched. No new releases were cut. The project’s development velocity remains strong, driven by a large batch of enhancements across observability, runtime stability, provider compatibility, and CI security. The quick turnaround of a critical Groq provider bug (fix merged same day) underscores responsive maintenance. The single closed issue (#8143) implemented a major infrastructure task (moving `.po` translations to a git submodule), further tidying the codebase.

### 2. Releases
*No releases published today.* Omit.

### 3. Project Progress
Seven PRs were merged or closed today. Notable items:

- **#8227** `[fix(providers): expose replay_assistant_reasoning and fall back tool_call_id (#8219)]` — Closed (merged). Directly resolves the critical Groq multi-turn tool loop bug (see Bugs & Stability).
- **#8220** `[test(channels): add regression for JSON tool envelope preservation during proactive trim]` — Closed (merged). Hardens channel orchestrator against a regression scenario with `tool_call_id`.
- **#8143** `[Task]: Implement RFC #7184 (.po only)` — Closed. Moved gettext `.po` translation files to a dedicated git submodule, improving repository hygiene.
- Other closed PRs (4 more) likely include minor fixes and dependency bumps (e.g., #8225, #8227, #8220, plus others not shown in top-20).

The majority of open PRs remain active and are advancing key features (see Feature Requests & Roadmap Signals).

### 4. Community Hot Topics
With no comments or reactions recorded on today’s issues, activity is primarily driven by the bug report and fix cluster:

- **Bug report #8219** — `[Bug]: gpt-oss-120b on Groq fails native multi-turn tool loops` — Critical provider issue. Attracted a same-day fix (#8227), indicating high urgency.
- **Feature request #8226** — `[enhancement]: support per-agent custom environment variables configuration` — Fresh request with no discussion yet; signals demand for per-agent isolation in execution contexts.
- **PR #8066** (opt-in LLM payload capture) and **#8065** (trace_id / cost_usd correlation) — Both by Nillth, represent a significant observability push, likely to be heavily discussed upon review.
- **PR #8173** (in-app upgrade with supervised restart) — A large (size: L), cross-domain change that touches gateway, daemon, and security; will require maintainer review.

**Underlying need:** Users are pushing for better observability (full request logging, tracing, cost tracking), better provider compatibility (Groq, NVIDIA NIM), and operational features (in-app upgrade, per-agent config).

### 5. Bugs & Stability
| Bug | Severity | Status | Notes |
|-----|----------|--------|-------|
| **#8219**: Groq provider – `tool_call_id` serialized as null, `reasoning_content` rejected on second turn | **Critical** | Fixed in #8227 (merged) | Issue still open; fix applied. Affects native tool-calling loops on Groq. |
| **#7345**: Vision routing false positives from path-listing tools (content_search, glob_search) | **High** | Open PR | PR #7345 gates path-listing results from vision routing to avoid incorrect provider selection. |
| **#8048**: `history_pruning` config silently overridden; tool-result content lost under context pressure | **High** | Open PR | Fix honors user setting; size M, risk medium. |
| **#8100**: NVIDIA NIM provider lacks vision support | **Medium** | Open PR | PR #8100 enables vision by overriding `supports_vision`. |
| **#8084**: `doctor` command fails for custom providers due to missing `Config` | **Medium** | Open PR | PR #8084 passes config to create provider correctly. |
| **#8146**: CLI one-shot loses telemetry when OTLP backend configured (process exits before flush) | **High** | Open PR | PR #8146 flushes telemetry before exit. |
| **#7771**: Lifecycle observer events missing `channel`, `agent_alias`, `turn_id` | **High** | Open PR | PR #7771 propagates fields from all paths, not just `Agent::turn`. |
| **#8003**: `session_end` hook never fired (dead code) | **High** | Open PR | PR #8003 wires existing `fire_session_end()` into RPC session teardown. |
| **#8149**: Mutex poison crash in plugin runtime (`HostContext`) | **High** (crash) | Open, needs-author-action | PR #8149 uses `unwrap_or_else` to tolerate poison. |

**Summary:** The Groq bug was the most critical and has been resolved. Multiple high-severity runtime/observability bugs have open PRs with active progress.

### 6. Feature Requests & Roadmap Signals
**New feature request:**  
- **#8226** – Per-agent custom environment variables in `AliasedAgentConfig`. Likely to be considered for a future release given the simplicity of the change (declarative `env` map).

**In-flight features (open PRs) that may land in the next release:**  

| PR | Feature | Size | Impact |
|----|---------|------|--------|
| #8066 | Opt-in LLM request payload capture (default off) | M | Observability – full audit trail of what was asked |
| #8065 | Log correlation by `trace_id` + per-call `cost_usd` | M | Observability – cost tracking and distributed tracing |
| #8185 | Suggest cached extra registry skills in install commands | S | Usability – better missing-skill guidance |
| #8182 | Restore client relationship knowledge graph actions (Postgres) | L | Memory – client, contact, interaction nodes |
| #8206 | Durable SQLite run-state store + live run metrics | L | Runtime – SOP checkpointing and progress tracking |
| #8173 | In-app upgrade with optional supervised restart (RFC #8170) | L | Operations – dashboard upgrade with zero-downtime |
| #8000 | Improved Zerocode UI (browse mode badge, auto-exit browse) | M | Developer tooling |
| #7846 | Wire `before_llm_call` hook into LLM call paths | XS | Extensibility – plugin hook before LLM call |
| #8168 | Trivy container scanning (CVE gates for PR/release) | S | CI/Security – automated vulnerability scanning |

**Prediction:** The next release will likely include observability improvements (#8066, #8065), the SQLite run-state store (#8206), and the in-app upgrade mechanism (#8173), alongside provider fixes.

### 7. User Feedback Summary
No explicit user feedback is captured in the raw data, but actionable signals can be inferred:

- **Pain point:** Groq provider users (especially those using `gpt-oss-120b`) experienced broken multi-turn tool loops – quickly fixed. (Reported by perlowja)
- **Pain point:** Users relying on OTLP telemetry backends (e.g., Langfuse) lose data on CLI one-shot runs – PR #8146 addresses this.
- **Pain point:** Custom provider users encounter validation errors in the `doctor` command – PR #8084 resolves.
- **Use case:** Developers deploying agents in multi-tenant or secure environments need per-agent environment variable isolation (issue #8226).
- **Use case:** Teams managing client relationships want a knowledge graph of contacts and interactions – PR #8182 restores that capability.
- **Satisfaction indicator:** The high number of enhancement PRs (50 in one day) suggests a healthy, engaged contributor community.

### 8. Backlog Watch
The data only shows items updated in the last 24 hours, so no long-untouched issues are visible. However, several open PRs have been pending for over two weeks and require maintainer attention:

- **#7345** (vision false positives) – Created 2026-06-07, still open with no recent comments. Risk: high, size: S. May need review.
- **#7846** (`before_llm_call` hook) – Created 2026-06-17, risk: high, size: XS. Awaiting review.
- **#7771** (observer event propagation) – Created 2026-06-16, size: L, risk: high. A large fix touching multiple components.
- **#8149** (mutex poison) – Tagged `needs-author-action`; blocked on author response.

Maintainers may want to prioritize these older items to prevent staleness.

*All links: https://github.com/zeroclaw-labs/zeroclaw/issues/{id} or /pull/{id}.*

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*