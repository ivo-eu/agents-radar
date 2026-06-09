# OpenClaw Ecosystem Digest 2026-06-09

> Issues: 147 | PRs: 475 | Projects covered: 13 | Generated: 2026-06-09 04:18 UTC

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

# OpenClaw Project Digest — 2026-06-09

## 1. Today’s Overview

OpenClaw remains highly active, with **147 issues** (83 open, 64 closed) and **475 pull requests** (314 open, 161 merged/closed) updated in the last 24 hours. Two new beta releases (v2026.6.5-beta.3 and v2026.6.5-beta.5) landed, bringing urgent fixes for QQBot thinking-scaffold leakage and MCP tool result coercion. The community is reporting a concentrated set of regressions, particularly around OpenAI/ChatGPT Responses transport failures, session context confusion, and memory vector search issues, while maintainers are pushing several large-quality PRs for agent runtime stability, Windows cross‑platform support, and QA taxonomy validation. Overall, the project is in a rapid‑fix cycle with strong contributor engagement.

## 2. Releases

**New releases (2):**
- **v2026.6.5‑beta.5** ([release](https://github.com/openclaw/openclaw/releases/tag/v2026.6.5-beta.5))  
  - QQBot now strips model reasoning/thinking scaffolding before native delivery, preventing raw `<thinking>` content from leaking into channel replies. (Thanks @openperf; fixes #89913, #90132)  
  - MCP tool results now coerce `resource_link`, `resource`, `audio`, malformed image, and future n.
- **v2026.6.5‑beta.3** ([release](https://github.com/openclaw/openclaw/releases/tag/v2026.6.5-beta.3))  
  - Identical QQBot fix and MCP coercion improvements (same highlights as beta.5).

No breaking changes or migration notes were reported for these releases.

## 3. Project Progress

**Merged/closed PRs (161 in last 24h)** – key highlights include:

- **Agent & session stability**  
  - `fix(agents): trim dense text delta snapshots` (#91580) – reduces memory pressure in streaming by dropping full accumulated snapshots from plain text delta events.  
  - `fix(model-fallback): don't rethrow provider-side AbortErrors as user cancellations` (#90908) – prevents false cancellation signals during API connection drops.  
  - `fix(compact): make /compact command cancelable via abortEmbeddedAgentRun` (#90821) – enables `/stop` to abort running compaction.  
  - `fix(usage): return all agent sessions when no agentId specified` (#89034) – restores pre‑v2026.5.17 behaviour for multi‑agent setups.

- **Cross‑platform & test quality**  
  - A series of PRs by @aniruddhaadak80 replaced hard‑coded `win32` skips with dynamic capability checks in browser paths, wait‑for‑download, IRC, LINE, Nextcloud Talk, Telegram, and WhatsApp test suites (#90369, #90374, #90376, #90379, #90384, #90387, #90393).  
  - `fix(browser): replace broad win32 skips with dynamic capability checks in paths tests` (#90369) – improves Windows test coverage.

- **Tooling & documentation**  
  - `fix(secrets): generate secretref reference docs from the credential matrix` (#89142) – automates drift detection for credential surface documentation.  
  - `fix(codex): restore memory recall guidance` (#91594) – re‑enables memory‑core recall instructions for Codex native turns.

- **Performance**  
  - `fix(tui): prewarm agent runtime before first send` (#86981) – reduces first‑message latency in TUI mode.  
  - `perf(control-ui): lazy load slash commands` (#91598) – stops loading all slash commands on startup.

These fixes address high‑priority regressions in session state, message delivery, and Windows compatibility.

## 4. Community Hot Topics

The most active issues (by comment count and reactions) reveal the community’s primary pain points:

| Issue | Comments | 👍 | Summary |
|-------|----------|----|---------|
| [#90083](https://github.com/openclaw/openclaw/issues/90083) [OPEN] | 15 | 3 | OpenAI/ChatGPT Responses transport fails with `invalid_provider_content_type` for GPT‑5.4/5.5 after upgrade to 2026.6.1. |
| [#32296](https://github.com/openclaw/openclaw/issues/32296) [OPEN] | 14 | 1 | Agent replies to previous message instead of current – session context confusion that has been open since March. |
| [#85888](https://github.com/openclaw/openclaw/issues/85888) [OPEN] | 10 | 1 | Cron jobs fail consistently with MiniMax 503 overload during early morning, but manual triggers succeed – suggests a scheduling/request‑shaping issue. |
| [#65156](https://github.com/openclaw/openclaw/issues/65156) [CLOSED] | 7 | 1 | Memory vector search broken in v4.11 due to SQLite ABI mismatch – **fixed** (closed). |
| [#48300](https://github.com/openclaw/openclaw/issues/48300) [CLOSED] | 6 | 2 | `memory_search` hybrid mode not returning FTS matches – **fixed**. |
| [#44294](https://github.com/openclaw/openclaw/issues/44294) [OPEN] | 5 | 1 | ACP backend error kinds all mapped to `end_turn` – a long‑standing protocol design concern. |
| [#44293](https://github.com/openclaw/openclaw/issues/44293) [OPEN] | 5 | 1 | `pnpm check:docs` fails on native Windows PowerShell – contributor friction. |
| [#44297](https://github.com/openclaw/openclaw/issues/44297) [CLOSED] | 6 | 1 | Slack external arg‑menu fallback is silently downgraded – now surfaced as a visible health signal. |
| [#44292](https://github.com/openclaw/openclaw/issues/44292) [CLOSED] | 6 | 1 | Missing config field labels cause CI failures – autofix added. |

**Underlying needs**:  
- **Provider compatibility** – Users upgrading to GPT‑5.4/5.5 are blocked by a transport‑level content‑type check.  
- **Session correctness** – Context confusion (#32296) has been open >3 months, indicating a deep routing issue.  
- **Scheduled job reliability** – Cron tasks failing due to API overload (MiniMax) point to missing retry/back‑off logic.  
- **Platform inclusivity** – Windows contributor tooling (PowerShell, symlink handling) is receiving increased attention.

## 5. Bugs & Stability

**P1 (critical) bugs reported or active today:**

1. **OpenAI ChatGPT Responses transport fails** ([#90083](https://github.com/openclaw/openclaw/issues/90083)) – GPT‑5.4/5.5 inference returns `invalid_provider_content_type`. No fix PR visible; workaround needed. **Severity: High** – blocks all users of these models.
2. **Session context confusion** ([#32296](https://github.com/openclaw/openclaw/issues/32296)) – Agent replies to previous message. Open since March, no fix PR. **Severity: High** – causes persistent conversation misalignment.
3. **DeepSeek V4 Flash incomplete turn** ([#88657](https://github.com/openclaw/openclaw/issues/88657)) – `payloads=0, tools=2, stopReason=stop` after upgrade to 2026.5.27/28, regressed from 2026.5.26. No fix PR. **Severity: High** – model unusable.
4. **WhatsApp monitor drops inbound silently** ([#91191](https://github.com/openclaw/openclaw/issues/91191)) – Messages from one specific number are discarded. Needs maintainer review + info. **Severity: High** – message loss.
5. **Gateway heap grows to 1073MB+ at idle** ([#87109](https://github.com/openclaw/openclaw/issues/87109)) – Causes cron job failures under memory pressure. No fix PR. **Severity: High** – stability risk.
6. **LM Studio sessions start without tools** ([#91434](https://github.com/openclaw/openclaw/issues/91434)) – Zero tools compiled without warning. Needs product decision. **Severity: High** – silent failure for tool‑dependent agents.
7. **MCP OAuth two bugs** ([#91433](https://github.com/openclaw/openclaw/issues/91433)) – Regex failure under `/u` and OAuth errors surfaced as `[object Response]`. No fix PR yet. **Severity: High** – blocks streamable‑http MCP servers.

**P2 regressions with fix PRs in flight**:
- `claude-cli tool-call-then-empty-text` leads to duplicate reply in Telegram ([#91302](https://github.com/openclaw/openclaw/issues/91302)) – **closed** with fix.
- Memory index meta never written during auto‑sync ([#90338](https://github.com/openclaw/openclaw/issues/90338)) – **closed**.
- Telegram `/compact` slash command never appears in logs ([#89525](https://github.com/openclaw/openclaw/issues/89525)) – **closed**.
- `openai-completions` adapter passes empty `content[]` silently ([#91394](https://github.com/openclaw/openclaw/issues/91394)) – **closed**.

Several other P2 bugs have fix PRs linked (e.g., #88933, #90157, #90891, #91051, #91150), indicating a healthy fix velocity.

## 6. Feature Requests & Roadmap Signals

Notable feature requests with ongoing discussion:

- **Configurable Dream Diary language** ([#79223](https://github.com/openclaw/openclaw/issues/79223)) – The hardcoded English prompt in `memory-core` blocks non‑English workspaces. Likely to be addressed in a future plugin version.
- **Skill Workshop targeting shared skills** ([#74601](https://github.com/openclaw/openclaw/issues/74601)) – Enables proposing refinements to shared skills across workspaces. Still open, could appear in a minor release.
- **Expose OPENCLAW_SESSION_KEY env var in exec child processes** ([#66705](https://github.com/openclaw/openclaw/issues/66705)) – Has a linked PR open, so likely near completion.
- **Add native PowerShell smoke coverage** ([#44291](https://github.com/openclaw/openclaw/issues/44291)) – Part of a broader Windows contributor experience push; recent PRs improving test infrastructure align with this.
- **Preserve structured ACP error kinds** ([#44294](https://github.com/openclaw/openclaw/issues/44294)) – A new PR (#91603) has been submitted today to address exactly this, so it is likely for the next release.

**PRs that signal roadmap direction**:
- `feat(subagents): forward toolsAllow from sessions_spawn` ([#78441](https://github.com/openclaw/openclaw/pull/78441)) – Extends subagent control.
- `Discord: add reaction notification wake policy` ([#83169](https://github.com/openclaw/openclaw/pull/83169)) – Enhances Discord integration.
- `Add QA scorecard taxonomy validation` ([#91500](https://github.com/openclaw/openclaw/pull/91500)) – Indicates a maturation of the QA framework.

**Prediction for next versions (v2026.6.x)**:
- ACP error kind preservation (PR #91603).
- Secretref documentation autogeneration (PR #89142).
- PowerShell compatibility improvements (several test PRs).
- Memory core fixes for Dream Diary language.

## 7. User Feedback Summary

**Pain points reported today (real user voices):**

- “After upgrading to OpenClaw 2026.6.1… inference fails for `openai/gpt-5.4` and `openai/gpt-5.5`” (#90083) – *blocks model upgrade path.*
- “The agent responds to the previous user message instead of the current message” (#32296) – *basic conversation coherence not working.*
- “Repeated shared SQLite plugin install metadata conflict warnings for codex/discord” (#90418) – *annoying but non‑fatal noise.*
- “Gateway heap grows to 1073MB+ at idle on macOS, cron jobs fail silently” (#87109) – *resource leak threatening reliability.*
- “/compact skips with ‘no real conversation messages’ while /status shows substantial Context” (#91150) – *UX confusion.*
- “WhatsApp monitor drops inbound silently for one number” (#91191) – *message loss without notification.*

**Satisfaction signals:**  
- Positive comments in issue descriptions (“I think OpenClaw is genuinely brilliant… this project is exceptional” — #88933).  
- Quick turnarounds on P2 fixes (e.g., #91302, #90338 closed within 24 hours).  
- Community contributors submitting multiple test improvement PRs (e.g., @aniruddhaadak80) shows trust in the project’s direction.

**Overall sentiment**: Users are engaged but frustrated by regressions introduced in the 2026.6.x cycle, especially around core transport and session handling. The maintenance team is responding rapidly with patches, which tempers dissatisfaction.

## 8. Backlog Watch

**Long‑unanswered issues needing maintainer attention:**

- **#32296** (Session context confusion, open since March 2) – P1, 14 comments, no fix PR. Appears to be a deep architectural issue.
- **#44289** (Generate secretref docs, open since March 12) – P2, stale, linked to PR #89142 which is now open – **moving forward**.
- **#44291** (PowerShell smoke coverage, open since March 12) – P2, stale, no progress discussion.
- **#44293** (Make `pnpm check:docs` work in PowerShell, open since March 12) – P2, stale.
- **#44294** (Preserve ACP error kinds, open since March 12) – **just got a fix PR today** (#91603) – may be resolved soon.
- **#66705** (Expose session key env vars, open since April 14) – P2, has a linked PR open, but PR appears stalled.
- **#74601** (Skill Workshop shared skills, open since April 29) – Enhancement, no maintainer reply.
- **#78754** (Weixin unsupported channel error, open since May 7) – Regression, no fix PR.
- **#79223** (Dream Diary language, open since May 8) – Enhancement, no maintainer reply.
- **#87109** (Gateway heap growth, open since May 27) – P1, critical, but no fix PR yet – *needs urgent triage*.
- **#87291** (Reply context truncated at 500 chars, open since May 27) – P2, no fix PR.
- **#87229** (WhatsApp media metadata omitted, open since May 27) – P2, no fix PR.
- **#87327** (Isolated agent runs stall, open since May 27) – P1, no fix PR.
- **#88657** (DeepSeek V4 Flash incomplete turn, open since May 31) – P1, no fix PR – *needs maintainer review*.
- **#90083** (OpenAI transport failure, open since June 4) – P1, new, no fix PR yet – *rapidly growing comments*.
- **#91433** (MCP OAuth bugs, open since June 8) – P1, new, no fix PR yet.
- **#91434** (LM Studio no tools, open since June 8) – P1, new, no fix PR yet.

**PRs awaiting maintainer review (status: 📣 needs proof / 👀 ready for maintainer look):**
- **#89045** (fix terminal session status recovery) – P1, ready for review.
- **#89040** (avoid event‑loop stall in embedded_run) – P1, ready.
- **#91580** (trim dense text delta snapshots) – P1, ready.
- **#91500** (QA taxonomy validation) – maintainer, ready.
- **#91557** (iPad/iPhone control surfaces) – maintainer, ready.
- **#78441** (subagents toolsAllow) – large PR, ready.
- **#83169** (Discord reaction notification) – needs proof.

The backlog contains several critical P1 bugs that have been open for over a week without a fix PR (especially #87109, #88657). Maintainers should prioritise reviewing the ready‑to‑merge PRs that address session recovery and event‑loop stalls (#89045, #89040).

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report — AI Agent Open-Source Ecosystem
**Date:** 2026-06-09 | **Prepared for:** Technical Decision-Makers & Developers

---

## 1. Ecosystem Overview

The personal AI assistant and agent open-source ecosystem is experiencing a bifurcation between **battle-tested reference implementations** (OpenClaw) and **rapidly iterating specialized forks** (NanoBot, CoPaw, ZeroClaw). OpenClaw remains the de facto core with the largest contributor base and most mature MCP/session infrastructure, while projects like Hermes Agent and IronClaw push boundaries in desktop UX and enterprise workflow routing. The ecosystem is collectively prioritizing **cross-platform reliability, channel adapter parity, and security hardening** — with over 350 combined PRs across tracked projects in a single day. A notable pattern is the emergence of **WASM plugin systems** (ZeroClaw) and **gamified agent loops** (CoPaw's "learning loop" requests), suggesting the next wave of differentiation will center on extensibility and self-improving agent behaviors.

---

## 2. Activity Comparison

| Project | Issues (Open/Closed) | PRs (Open/Merged) | Release Today | Health Score* |
|---------|---------------------|-------------------|---------------|---------------|
| **OpenClaw** | 83 / 64 | 314 / 161 | ✅ 2 betas | **A** (High activity, rapid fix cycle) |
| **NanoBot** | 3 / 4 | 19 / 15 | ❌ | **B+** (Focused, responsive) |
| **Hermes Agent** | 3 / 1 | 45 / 5 | ❌ | **B** (High feature velocity, moderate bug triage) |
| **PicoClaw** | 2 / 1 | 11 / 9 | ✅ Nightly | **B** (Code quality push, niche platform gaps) |
| **NanoClaw** | 1 / 0 | 1 / 3 | ❌ | **C** (Low activity, critical WhatsApp bug) |
| **NullClaw** | 0 / 0 | 0 / 0 | ❌ | **Inactive** |
| **IronClaw** | 7 / 4 | 27 / 23 | ❌ | **A-** (Enterprise-grade, major refactor in progress) |
| **LobsterAI** | 1 / 0 | 7 / 16 | ❌ | **A-** (Exceptional fix velocity, backlog clearing) |
| **TinyClaw** | 0 / 0 | 1 / 0 | ❌ | **B-** (Quiet, one install fix in progress) |
| **Moltis** | 0 / 0 | 0 / 0 | ❌ | **Inactive** |
| **CoPaw** | 8 / 13 | 47 / 24 | ❌ | **B+** (Very high activity, growing contributor base) |
| **ZeptoClaw** | 0 / 0 | 0 / 0 | ❌ | **Inactive** |
| **ZeroClaw** | 15 / 2 | 40 / 10 | ❌ | **B+** (Active, structured milestone tracking) |

*\*Health Score: Qualitative assessment based on PR velocity, bug triage speed, backlog management, and community responsiveness.*

---

## 3. OpenClaw's Position

### Advantages vs. Peers
- **Largest contributor base:** 475 PRs (314 open) vs. next closest IronClaw (50) and CoPaw (47) — ~6x more active PRs than any single peer.
- **Two beta releases in 24 hours** — demonstrates rapid iteration capability unmatched by any other project.
- **MCP ecosystem maturity:** First-class MCP tool result coercion, OAuth handling, and `streamable-http` support — peer projects are still catching up (e.g., ZeroClaw's MCP dashboard still in v0.8.3 planning).
- **Comprehensive channel support:** QQBot, Discord, Slack, Telegram, LINE, IRC, WhatsApp, Nextcloud Talk — broader than any competitor's adapter matrix.

### Technical Approach Differences
- **Structured session context model** — OpenClaw's `/compact` and embedded agent run infrastructure is more formalized than NanoBot's or CoPaw's session management.
- **Provider abstraction depth:** Separate transports for OpenAI Chat/Responses, streaming, and fallback chains — Hermes Agent and IronClaw are still building equivalent layers.
- **Security-first design:** Secretref auto-documentation, credential matrix drift detection, and ACP error kind preservation — peers are only now adding egress lockdown (NanoClaw) and SSRF protection (NanoBot).

### Community Size Comparison
- **OpenClaw:** 147 issues, 475 PRs — likely 50-100+ active weekly contributors.
- **IronClaw:** 11 issues, 50 PRs — smaller but enterprise-focused.
- **CoPaw:** 21 issues, 47 PRs — growing Chinese-speaking community.
- **NanoBot:** 7 issues, 34 PRs — moderate, academic-backed (HKUDS).
- **ZeroClaw:** 17 issues, 50 PRs — active but smaller core team.

**Verdict:** OpenClaw is the clear **reference implementation** and **community hub**, but its complexity introduces regression risks (P1 bugs in transport, session confusion) that smaller projects may avoid through simpler architectures.

---

## 4. Shared Technical Focus Areas

| Focus Area | Affected Projects | Specific Needs |
|------------|------------------|----------------|
| **Provider Compatibility Regressions** | OpenClaw, CoPaw, PicoClaw, IronClaw | GPT-5.4/5.5 transport failures (OpenClaw #90083), DeepSeek V4 Flash incomplete turns (OpenClaw #88657), Minimax config failure (IronClaw #4587), GPT-5.x hardcoded model lists (CoPaw #2777) |
| **Cross-Platform Parity** | OpenClaw, PicoClaw, ZeroClaw, CoPaw | Windows PowerShell CI failures (OpenClaw #44293), RISC-V non-functional (PicoClaw #2887), Windows console flash (PicoClaw #3061), macOS/iOS control surfaces (OpenClaw #91557) |
| **Session & Context Reliability** | OpenClaw, NanoBot, ZeroClaw | Context confusion (OpenClaw #32296), memory vector search (OpenClaw #65156), orphan tool results (NanoBot #4219), Matrix session isolation (ZeroClaw #6487) |
| **MCP/Plugin Infrastructure** | OpenClaw, NanoBot, ZeroClaw, CoPaw | OAuth bugs (OpenClaw #91433), SSRF loopback (NanoBot #4074), WASM path divergence (ZeroClaw #6254), plugin marketplace (CoPaw #5023) |
| **Security Hardening** | NanoClaw, ZeroClaw, IronClaw, NanoBot | Egress lockdown (NanoClaw #2713), SSRF rejection (NanoBot #4074), SecurityAuditSink (IronClaw #3959), symlink escape prevention (NanoBot #4221) |
| **Channel Adapter Reliability** | OpenClaw, NanoBot, PicoClaw, NanoClaw, ZeroClaw, CoPaw | WhatsApp silent drops (OpenClaw #91191), Telegram code-block splitting (NanoBot #4250), QQ token timeout (PicoClaw #3015), WeChat token expiry (CoPaw #4477), Feishu LLM-only (ZeroClaw #4873) |
| **Observability & Tracing** | OpenClaw, IronClaw, ZeroClaw, CoPaw | Turn metadata (ZeroClaw #7385), gateway heap monitoring (OpenClaw #87109), Langfuse/OpenTelemetry requests (CoPaw #5009), trajectory observer (IronClaw #4588) |

**Key Insight:** **Provider API compatibility** and **session correctness** are the most widespread pain points — affecting 5+ projects each. This suggests the underlying LLM provider APIs are evolving faster than the open-source agent layer can adapt.

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | Hermes Agent | IronClaw | CoPaw | ZeroClaw |
|-----------|----------|---------|--------------|----------|-------|----------|
| **Primary Use Case** | Reference agent orchestration | Lightweight multi-provider assistant | Desktop-first personal agent | Enterprise workflow routing | Chinese-market agent platform | Modular WASM plugin agent |
| **Target User** | Developers, integrators | Researchers, hobbyists | Power users, Mac users | Enterprise teams | Chinese-speaking end users | Tinkerers, security-conscious |
| **Architecture** | Monolithic core + plugins | Plugin-lite, academic | Desktop Tauri app | Rust-based, microservices | Python, Tauri desktop | Rust, WASM sandboxing |
| **Channel Parity** | ✅ 10+ channels | ✅ 5+ channels | ❌ Slack, Discord, iMessage only | ❌ IRC, Telegram, Slack (in progress) | ✅ 8+ channels (WeChat focus) | ✅ 6+ channels (Matrix focus) |
| **Security Maturity** | High (secretref, ACP kinds) | Medium (SSRF, symlink) | Medium (ZAI billing fix) | Very High (audit sink, hooks) | Low-Medium | High (egress, pluggable providers) |
| **Community Language** | English | English/Chinese | English | English | Chinese | English |
| **Release Cadence** | Multiple betas/week | ~Weekly | ~Bi-weekly | ~Monthly | ~Weekly | Milestone-based |
| **Key Weakness** | Regression risk from rapid changes | Limited adaptability | Niche platform focus | Slow to ship (breaking changes) | English docs gap, Windows perf | High backlog of accepted bugs |

**Architectural Divergence Example:**
- **OpenClaw** uses a centralized session context model with `/compact` and embedded agent runs.
- **NanoBot** uses per-provider `extra_query` for flexible routing.
- **IronClaw** implements `ProductWorkflow` as a routing layer for OpenAI-compatible endpoints.
- **ZeroClaw** uses WASM sandboxes with per-alias session isolation.

**Takeaway:** No single project dominates all dimensions. OpenClaw is best for **broad compatibility and community support**, IronClaw for **enterprise production** (at the cost of slower shipping), and CoPaw for **Chinese ecosystem integration**. The choice depends on deployment region, security requirements, and tolerance for breaking changes.

---

## 6. Community Momentum & Maturity

### Tier 1: Rapidly Iterating (High risk, high reward)
| Project | Velocity | Key Risk | Why Watch |
|---------|----------|----------|-----------|
| **OpenClaw** | 🚀🔥 **Extreme** | Regression cascade (7 P1 bugs open) | Core reference; every fix upstream influences forks |
| **CoPaw** | 🚀🔥 **Very High** | Scalability under active development | Growing Chinese-speaking contributor base |
| **IronClaw** | 🚀 **High** | Breaking changes in release pipeline | Enterprise architecture patterns |
| **ZeroClaw** | 🚀 **High** | Milestone debt (v0.8.0 overdue) | WASM plugin model is novel |

### Tier 2: Stabilizing (Lower risk, consolidating)
| Project | Velocity | Stability Signal | Why Watch |
|---------|----------|-----------------|-----------|
| **NanoBot** | 🟡 **Moderate** | Quick fix turnaround (SSRF, symlink) | Good for smaller deployments |
| **LobsterAI** | 🟢 **Refactoring** | Closed 8 stale PRs in one day | Clean codebase, Electron-based |
| **PicoClaw** | 🟢 **Quality focus** | 9 defensive code PRs merged | Embedded/Niche platform focus |

### Tier 3: Quiet / Stalled
| Project | Concern |
|---------|---------|
| **NanoClaw** | Single critical WhatsApp bug with no fix; low contributor engagement |
| **TinyClaw** | Only one install-fix PR open; minimal feature development |
| **NullClaw, Moltis, ZeptoClaw** | No activity in 24h; potentially dead or dormant |

**Ecosystem Observation:** The gap between **Tier 1** and **Tier 3** is widening. Projects that can't keep up with provider API changes (GPT-5.x, MCP evolution) are being abandoned. The few truly "stabilizing" projects (PicoClaw, LobsterAI) are those explicitly choosing to narrow their scope.

---

## 7. Trend Signals

### What Users Are Asking For (Implied Requirements)

| Trend | Evidence (Projects Affected) | Impact on Agent Developers |
|-------|------------------------------|---------------------------|
| **Model Flexibility Per Session** | NanoBot (#4253), CoPaw (#2777), OpenClaw (#90083) | Need per-conversation provider routing, not just per-instance config |
| **Self-Improving Agents** | CoPaw (#5017), ZeroClaw (#5844) | "Learning loops" and configurable memory weight — next UX frontier |
| **Observability as a Feature** | OpenClaw (#87109), IronClaw (#4588), CoPaw (#5009), ZeroClaw (#7385) | Token usage, latency, cost attribution, trajectory traces — production must-haves |
| **Pluggable Security Policies** | NanoClaw (#2713), ZeroClaw (#7142), OpenClaw (#44294) | Egress lockdown, pluggable audit sinks, per-tenant auth |
| **Cross-Platform Parity** | OpenClaw (#44293), PicoClaw (#2887), CoPaw (#5015) | Windows/macOS/Linux/RISC-V — ecosystem is demanding universal runs |
| **Automated Skill/Plugin Discovery** | ZeroClaw (#6254), CoPaw (#5023), OpenClaw (#74601) | Dynamic model fetching, plugin markets, shared skill workshops |
| **Channel-Specific Reliability** | NanoBot (#4250), OpenClaw (#91191), PicoClaw (#3015), CoPaw (#5030) | Telegram code blocks, WhatsApp drops, WeChat token expiry — each channel has its own failure modes |

### Industry-Level Signals

1. **Provider API Fragmentation is the #1 Headache:** GPT-5.x, DeepSeek V4, Minimax, and Xiaomi ASR all broke in different ways across projects. The agent layer is becoming a **translation layer** between rapidly changing provider APIs and stable user expectations.

2. **Security is Moving from Opt-In to Default-On:** Egress lockdown (NanoClaw), SSRF protection (NanoBot), and audit sinks (IronClaw) are becoming standard. Projects without security hardening are increasingly risky for production.

3. **The Desktop is the New Battleground:** Hermes Agent (Tauri), CoPaw (Tauri), and LobsterAI (Electron) all invest in desktop UX. OpenClaw's TUI pre-warming (#86981) shows even the core project acknowledges this shift.

4. **Chinese Ecosystem is a Major Growth Vector:** CoPaw's community has 8+ WeChat-compatible channels and explicit Hermes Agent requests. PicoClaw and ZeroClaw also show Chinese-language issue reporting. Projects ignoring this may miss a significant user base.

5. **WASM Plugin Systems are Emerging as a Differentiator:** ZeroClaw's WASM sandboxing and CoPaw's plugin marketplace suggest the next architectural shift is toward **extensible, sandboxed runtimes** rather than monolithic cores.

### Value for AI Agent Developers

- **If building a production deployment:** Start with OpenClaw for breadth, but plan for provider API shims and session isolation testing. Monitor IronClaw for enterprise patterns.
- **If targeting desktop users:** Use Hermes Agent as UX inspiration, but evaluate CoPaw for Chinese market reach.
- **If prioritizing security:** Study NanoClaw's egress lockdown and ZeroClaw's pluggable providers. Expect OpenClaw to absorb these in future releases.
- **If building an agent orchestration platform:** Watch ZeroClaw's WASM model — it may define how third-party plugins integrate across projects.

---

*Report generated from 13 project digests dated 2026-06-09. All links reference GitHub repositories as noted in source digests.*

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest – 2026-06-09

## 1. Today's Overview
The project remains highly active: **34 PRs were updated in the last 24 hours** (15 merged/closed, 19 open), alongside **7 issues updated** (3 open, 4 closed). No new releases were tagged, but the commit/PR velocity indicates continuous feature integration and bug‑fixing. Key merges landed in transcription providers, extra query support for OpenAI‑compatible backends, and several reliability patches. Community interaction is moderate, with enhancement requests and bug reports receiving prompt attention.

---

## 2. Releases
*None – no releases were published in the last 24 hours.*

---

## 3. Project Progress (Merged/Closed PRs)
Fifteen PRs were merged or closed in the last 24 hours. Notable contributions:

- **Agent Collaboration** – [`#3992`](https://github.com/HKUDS/nanobot/pull/3992) (merged) enables cross‑instance agent messaging via a shared message bus.
- **Transcription Providers**  
  - [`#4113`](https://github.com/HKUDS/nanobot/pull/4113) – OpenRouter as a transcription provider.  
  - [`#4175`](https://github.com/HKUDS/nanobot/pull/4175) – Xiaomi MiMo ASR (`mimo‑v2.5‑asr`).  
  - [`#4224`](https://github.com/HKUDS/nanobot/pull/4224) – AssemblyAI transcription.
- **Shared Voice Input** – [`#4232`](https://github.com/HKUDS/nanobot/pull/4232) promotes transcription from a channel‑only feature to a shared capability used by WebUI and desktop.
- **Extra Query Support** – [`#4217`](https://github.com/HKUDS/nanobot/pull/4217) adds `ProviderConfig.extra_query` for Azure‑style gateways (closes [#4204](https://github.com/HKUDS/nanobot/issues/4204)).
- **Session & Memory Fixes**  
  - [`#4219`](https://github.com/HKUDS/nanobot/pull/4219) – Drops orphan tool results before history trimming.  
  - [`#4221`](https://github.com/HKUDS/nanobot/pull/4221) – Blocks relative symlink escapes in `ExecTool` (closes [#4072](https://github.com/HKUDS/nanobot/issues/4072)).
- **Security** – [`#4074`](https://github.com/HKUDS/nanobot/issues/4074) (closed) addresses SSRF rejection for MCP HTTP/SSE.

---

## 4. Community Hot Topics
| Item | Type | Comments | Summary |
|------|------|----------|---------|
| [#4253](https://github.com/HKUDS/nanobot/issues/4253) | Issue (Open) | 1 | User wants to override model provider per conversation to switch between fast (OpenRouter) and private/local (llamacpp) presets. |
| [#4233](https://github.com/HKUDS/nanobot/issues/4233) | Issue (Open) | 1 | Request to display NanoBot version in WebUI, possibly with update notifications. |
| [#4250](https://github.com/HKUDS/nanobot/issues/4250) | Issue (Open) | 0 | Telegram `split_message` breaks fenced code blocks; fix PR [#4257](https://github.com/HKUDS/nanobot/pull/4257) is open. |
| [#4257](https://github.com/HKUDS/nanobot/pull/4257) | PR (Open) | – | Fenced‑code‑block‑aware splitting for Telegram – directly addresses #4250. |
| [#4252](https://github.com/HKUDS/nanobot/pull/4252) | PR (Open) | – | Adds TeX math rendering (`\(...\)`, `\[...\]`) to WebUI using remark‑math/rehype‑katex. |

**Analysis:** The most active topic is **model flexibility per conversation** – a common requirement for users mixing local and cloud LLMs. The WebUI version display and file‑upload requests also draw interest. The Telegram code‑block bug is quickly getting a targeted fix.

---

## 5. Bugs & Stability
Three bugs were reported today; two already have fix PRs in review.

| Severity | Issue | Description | Status |
|----------|-------|-------------|--------|
| **High** | [#4250](https://github.com/HKUDS/nanobot/issues/4250) | `split_message` splits fenced code blocks, causing broken HTML in Telegram. | Open; fix PR [#4257](https://github.com/HKUDS/nanobot/pull/4257) open. |
| **Medium** | [#4256](https://github.com/HKUDS/nanobot/pull/4256) | Memory cursor may become non‑monotonic after history compaction, risking ID collisions. | Open fix PR. |
| **Low** | [#4223](https://github.com/HKUDS/nanobot/pull/4223) | WeChat channel enters permanent silent loop when session token expires – fix reloads session state after pause. | Open PR. |

Earlier high‑severity security fix [#4074](https://github.com/HKUDS/nanobot/issues/4074) (SSRF loopback in MCP) was closed with a fix merged previously. Overall stability is improving, especially with systematic symlink‑escape and orphan‑tool‑result fixes.

---

## 6. Feature Requests & Roadmap Signals
**Strong signals (likely in next minor release):**
- **Model override per conversation** ([#4253](https://github.com/HKUDS/nanobot/issues/4253)) – closely aligned with the recent `extra_query` and transcription improvements; would fit the provider abstraction layer.
- **File/PDF upload in chat** ([#4251](https://github.com/HKUDS/nanobot/issues/4251)) – voiced in Chinese; user expects image/PDF upload for summarization. Could leverage existing vision/transcription pipelines.
- **WebUI version display** ([#4233](https://github.com/HKUDS/nanobot/issues/4233)) – low effort, high value for users.

**Recently shipped signals:**  
- Transcription provider expansion (Xiaomi, AssemblyAI, OpenRouter) and shared voice input ([#4232](https://github.com/HKUDS/nanobot/pull/4232)) suggest the team is prioritizing multimodal input.
- Cross‑instance agent collaboration ([#3992](https://github.com/HKUDS/nanobot/pull/3992)) hints at a roadmap for distributed agent ecosystems.

**Prediction:** The next version will likely include model per‑conversation support and basic file upload, building on the flexible provider config introduced in 2026-06.

---

## 7. User Feedback Summary
- **Pain points:**
  - Inability to switch between local and cloud models mid‑conversation ([#4253](https://github.com/HKUDS/nanobot/issues/4253)).
  - Telegram code‑block rendering broken ([#4250](https://github.com/HKUDS/nanobot/issues/4250)).
  - Missing WebUI version information ([#4233](https://github.com/HKUDS/nanobot/issues/4233)).
  - WeChat channel silent after token expiry ([#4223](https://github.com/HKUDS/nanobot/pull/4223)).
  - Chinese‑speaking user asks for file/PDF upload support ([#4251](https://github.com/HKUDS/nanobot/issues/4251)).
- **Satisfaction signals:**
  - Rapid implementation of extra query parameters for Azure gateways ([#4217](https://github.com/HKUDS/nanobot/pull/4217)).
  - Multiple new transcription providers (Xiaomi, AssemblyAI, OpenRouter) – a clear answer to community requests for diversity in STT backends.

---

## 8. Backlog Watch
While most open issues and PRs are recent, several important PRs have been open for over a week without significant maintainer activity:

| Item | Age | Description | Notes |
|------|-----|-------------|-------|
| [#4177](https://github.com/HKUDS/nanobot/pull/4177) | 6 days | Onboarding docs for beginners – high community value. | No recent comments. |
| [#4119](https://github.com/HKUDS/nanobot/pull/4119) | 9 days | Symlink escape fix – **superseded** by merged [#4221](https://github.com/HKUDS/nanobot/pull/4221)? Needs closure. | Unclear if still relevant. |
| [#3982](https://github.com/HKUDS/nanobot/pull/3982) | 16 days | Scripted agent runner harness for tests. | Open with no recent activity; core testing infrastructure. |
| [#3983](https://github.com/HKUDS/nanobot/pull/3983) | 16 days | Runner test coverage for blocked tool-call finish reasons. | Same author, both probably waiting for review. |

**Action:** Maintainers should evaluate [#4119](https://github.com/HKUDS/nanobot/pull/4119) for closure (the fix now lives in [#4221](https://github.com/HKUDS/nanobot/pull/4221)) and triage the runner test PRs to merge or update.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-06-09

## 1. Today’s Overview

The Hermes Agent project saw high activity over the past 24 hours, with **50 pull requests updated** (45 open, 5 merged/closed) and **4 issues updated** (3 open, 1 closed). No new releases were published. Development velocity is strong across multiple fronts: CLI, desktop UI, tool integrations (web, browser, MCP, Discord, iMessage), agent core improvements, and security fixes. The backlog is being actively addressed, with several P1/P2 bug-fix PRs submitted. Community engagement is moderate, with one long-standing feature request (#12655) still under discussion and a Slack integration change moving to completion.

## 2. Releases

**None.** No new versions were tagged in the last 24 hours. The latest release (if any) predates this digest.

## 3. Project Progress

**5 pull requests were merged or closed today.** While the specific merged PRs are not listed individually in the top-20 set, the closed issue **#42601 – “Slack: use Block Kit for structured notifications”** (updated today, now closed) suggests the corresponding work has been completed. Additionally, the following PRs (all still open) represent major progress toward feature completion:

- **#42507** – Complete bidirectional iMessage media + interactions (tapbacks, voice, video, read receipts, typing indicators). *Tracked issue #42454.*
- **#42282** – Remote onboarding path for desktop app (first-class “Connect to existing Hermes”).
- **#41402** – Bubble chat layout (dual-column, user/assistant alignment).
- **#42598** – Graphify local notes lane with query presets and note-saving controls.
- **#41871** – Gateway Gateway relay adapter (phases 0–1) for disposable/scale-to-zero gateway.
- **#42595** – Gemini TTS persona prompts and audio tags.
- **#42521** – Resizable themed terminal pane + command-palette hotkeys.

These indicate strong forward momentum on UX, plugin ecosystem, and multi-platform integration.

## 4. Community Hot Topics

**Most commented issue:**
- **[#12655](https://github.com/NousResearch/hermes-agent/issues/12655) – `picker_providers` config for `/model`** (6 comments, 0 👍).  
  This feature request from April 19 gained attention again. Users who rely on custom endpoints want to hide built-in providers (Anthropic, OpenRouter) that auto-detect via OAuth/env. The need for better provider filtering in the model picker is a recurring theme.

**Other active discussions:**
- **[#42601](https://github.com/NousResearch/hermes-agent/issues/42601) – Slack Block Kit** (2 comments, closed today). The resolution distinguishes normal conversational replies (plain text/mrkdwn) from structured notifications (Block Kit). Community desire for more Slack richness is now addressed.

**PRs with undefined comment counts** (likely 0) — no PRs received more than trivial discussion among the top 20. This may indicate a maintainer-driven development pace with less community vetting on individual pulls.

## 5. Bugs & Stability

Three new bugs were reported today:

| Issue | Severity | Description | Fix PR Exists? |
|-------|----------|-------------|----------------|
| [#42585](https://github.com/NousResearch/hermes-agent/issues/42585) (Open) | **P2** | `web_search` remains exposed via the `browser` toolset even when the `web` toolset is disabled, making `hermes tools disable web --platform cli` misleading. | No dedicated fix PR yet, but PR [#42550](https://github.com/NousResearch/hermes-agent/pull/42550) reworks `web_search` priority chain and may address overlap. |
| [#42593](https://github.com/NousResearch/hermes-agent/issues/42593) (Open) | **P3** | File browser hover-reveal trigger zone (14px wide) blocks the chat scrollbar on the right edge. Mouse targeting the scrollbar accidentally opens the file browser. 130ms delay helps but doesn’t eliminate the issue. | No fix PR yet. |
| [#42601](https://github.com/NousResearch/hermes-agent/issues/42601) (Closed) | — | This was a feature request, not a bug. Closed today. | Work completed (Block Kit distinction). |

**Severity ranking:** The `web_search` exposure bug (#42585) is the most concerning because it undermines user trust in tool disable commands. It should be prioritized for a fix.

Additionally, several **bug-fix PRs** were submitted today:
- **[#42547](https://github.com/NousResearch/hermes-agent/pull/42547) (P2)** – Fix ZAI endpoint probing order to prevent Coding Plan keys from being routed to pay-as-you-go billing (silent wallet drain).
- **[#42548](https://github.com/NousResearch/hermes-agent/pull/42548) (P3)** – Pass timeout to title generator, cut `max_tokens` to 64.
- **[#42552](https://github.com/NousResearch/hermes-agent/pull/42552) (P2)** – Conditional MCP init for cron jobs; robust datetime parsing.
- **[#42553](https://github.com/NousResearch/hermes-agent/pull/42553) (P3)** – Migration safety guard for `model_override` column; single-worker claim protection.
- **[#42596](https://github.com/NousResearch/hermes-agent/pull/42596)** – Detect post-tool-call progress placeholders (<60 chars) and trigger nudge (fixes #42503).
- **[#42597](https://github.com/NousResearch/hermes-agent/pull/42597)** – Fix Discord voice meeting responsiveness; reset inactivity timers on inbound audio.
- **[#42600](https://github.com/NousResearch/hermes-agent/pull/42600)** – `/plugins` now shows installed-but-not-enabled plugins correctly.

These PRs demonstrate active bug triage and a healthy pipeline of stability improvements.

## 6. Feature Requests & Roadmap Signals

**Top requested features** (from issues and PR discussions):

- **Model picker provider filtering** ([#12655](https://github.com/NousResearch/hermes-agent/issues/12655)) — High demand from power users with custom endpoints. Likely candidate for next minor release.
- **Slack Block Kit structured notifications** ([#42601](https://github.com/NousResearch/hermes-agent/issues/42601)) — Now closed, so expected in upcoming release.
- **iMessage full media support** — PR [#42507](https://github.com/NousResearch/hermes-agent/pull/42507) is large and stacked. Likely to land in the next version.
- **Web extract LLM summarization thresholds** — PR [#42599](https://github.com/NousResearch/hermes-agent/pull/42599) makes `MAX_OUTPUT_SIZE` configurable. Addresses silent truncation complaints.
- **Tool search deferral** — PR [#42551](https://github.com/NousResearch/hermes-agent/pull/42551) allows platforms to control which core tools appear initially vs. by search. Aims to reduce UX clutter.
- **Gemini TTS persona prompts** — PR [#42595](https://github.com/NousResearch/hermes-agent/pull/42595) adds `tts.gemini.persona_prompt_file` and audio tags. Reflects growing interest in voice personality.
- **Bubble chat layout** — PR [#41402](https://github.com/NousResearch/hermes-agent/pull/41402) (from June 7) adds a dual-column bubble mode. A long-standing usability request.

**Predictions for next version:** The iMessage overhaul (#42507), bubble chat (#41402), and the web extract configurability (#42599) are large, well-reviewed PRs that may be merged soon. The model picker config (#12655) and Slack Block Kit (#42601) are likely part of the same release cycle.

## 7. User Feedback Summary

**Pain points expressed in today’s data:**

- **“`hermes tools disable web` is misleading”** (#42585) — Users expect the web toolset to be fully disabled, but `web_search` still works through the browser toolset. This erodes trust in configuration commands.
- **“File browser hover-reveal blocks scrollbar”** (#42593) — A minor but annoying UI issue when trying to scroll long conversations.
- **“Silent wallet drain with ZAI Coding Plan keys”** (#42547 fix) — Users on Coding Plan billing were incorrectly routed to pay-as-you-go, costing money without warning.
- **“Title generation blocks session when provider is slow”** (#42548 fix) — Timeout wasn’t passed, causing indefinite waits.
- **“Post-tool-call progress messages treated as final answers”** (#42596 fix) — Users saw “working on it…” as final responses, causing confusion.
- **“Discord voice meetings unresponsive”** (#42597 fix) — Inactivity timers not reset on inbound audio, causing early disconnection.

**Satisfaction signals:**
- Closure of [#42601](https://github.com/NousResearch/hermes-agent/issues/42601) (Slack Block Kit) indicates responsiveness to integration requests.
- Many PRs are incremental improvements (bubble layout, TTS persona, Gateway relay) that suggest an active and maturing product.

## 8. Backlog Watch

The following issue has been open for **over a month without a resolution**:

- **[#12655](https://github.com/NousResearch/hermes-agent/issues/12655)** (created 2026-04-19) – `picker_providers` config for `/model`.  
  Despite 6 comments and clear specification, no PR has been attached. This is a relatively small config change and could be a good candidate for a first-time contributor or maintainer quick-win. It would significantly improve UX for power users.

**Other items needing attention:**
- PR [#40663](https://github.com/NousResearch/hermes-agent/pull/40663) (created 2026-06-06, still open) – Security fix for shell-expanded command detection. It’s labeled P1 and addresses issue #36846. It has not been merged in three days; given its severity (dangerous command detection), maintainers should review and prioritize.
- No issues older than ~2 days (other than #12655) are languishing. Current triage appears healthy.

---

*Generated from Hermes Agent GitHub data snapshot on 2026-06-09. All links point to the NousResearch/hermes-agent repository.*

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

## PicoClaw Project Digest – 2026-06-09

### 1. Today's Overview
Project activity is **high**, with 20 pull requests and 3 issues updated in the last 24 hours. The community is submitting a large batch of defensive code improvements (unchecked type assertions, error handling, linter warnings) — 9 PRs were merged or closed, indicating a focused drive for code quality. A new **nightly v0.2.9** build was published, though it is marked as potentially unstable. The open issue count (2) is low, but one long-standing bug on RISC-V remains unresolved.

### 2. Releases

**`nightly` – v0.2.9-nightly.20260609.46b29a0a**  
*Automated build – may be unstable. Use with caution.*  
**Changelog**: [v0.2.9…main](https://github.com/sipeed/picoclaw/compare/v0.2.9...main)  
No breaking changes or migration notes beyond the usual nightly caveat.

### 3. Project Progress
**9 PRs merged/closed today** (all closed/merged in the list):

- **Location message support** – [#3052](https://github.com/sipeed/picoclaw/pull/3052) (merged) fixes Telegram location pins being silently ignored.
- **Health check fix** – [#3062](https://github.com/sipeed/picoclaw/pull/3062) (merged) resolves health endpoint always returning “not ready”.
- **Error wrapping improvements** – [#3051](https://github.com/sipeed/picoclaw/pull/3051), [#3060](https://github.com/sipeed/picoclaw/pull/3060) switch from `%v` to `%w` in channels, MCP, and skill helpers.
- **Structured logging** – [#3050](https://github.com/sipeed/picoclaw/pull/3050) replaces raw `log.Printf`/`fmt.Printf` with structured logger in state management.
- **Defensive type assertion checks** – [#3055](https://github.com/sipeed/picoclaw/pull/3055), [#3056](https://github.com/sipeed/picoclaw/pull/3056), [#3057](https://github.com/sipeed/picoclaw/pull/3057), [#3058](https://github.com/sipeed/picoclaw/pull/3058), [#3018](https://github.com/sipeed/picoclaw/pull/3018) – add `ok` checks in agent, tools, LINE, Evolution, and webfetch to prevent panics.

These changes collectively improve stability across Telegram, LINE, agent loops, and HTTP utilities.

### 4. Community Hot Topics
- **Most discussed issue**: [#2887](https://github.com/sipeed/picoclaw/issues/2887) (9 comments) – `.deb` on RISC‑V fails with OpenAI models. The user reports a complete non-functional state on Debian RISC‑V with `gpt-5.4-2026-03-05`. No fix PR exists yet; the issue is marked `[stale]` but remains open.
- **Active discussion**: [#3015](https://github.com/sipeed/picoclaw/issues/3015) (2 comments) – QQ channel connection fails on Windows with token timeout. The reporter notes Pico channel works fine, suggesting a Windows‑specific regression in the QQ adapter.
- **New feature PR**: [#3063](https://github.com/sipeed/picoclaw/pull/3063) (DeltaChat gateway) generated no comments yet but represents a community‑driven channel extension.

The underlying need is for **cross‑platform parity** (Linux RISC‑V, Windows) and **reliable channel adapters**, especially for niche platforms like QQ.

### 5. Bugs & Stability

| Bug | Severity | Status | Fix PR |
|-----|----------|--------|--------|
| RISC‑V/OpenAI not functional (#2887) | **High** – core feature broken on supported arch | Open since May 17, stale | None |
| QQ channel connection timeout on Windows (#3015) | **High** – blocks Windows users from QQ | Open since Jun 6 | None |
| Telegram location messages ignored (#3049) | **Medium** – missing channel feature | Closed (fixed by #3052) | [#3052](https://github.com/sipeed/picoclaw/pull/3052) |
| Console flash in Windows child processes | **Low** – cosmetic/UX | Open (PR #3061) | [#3061](https://github.com/sipeed/picoclaw/pull/3061) |

The RISC‑V issue is the most critical unresolved bug. The QQ problem is newer and lacks a fix PR. The location bug is now resolved.

### 6. Feature Requests & Roadmap Signals
- **DeltaChat gateway** (PR [#3063](https://github.com/sipeed/picoclaw/pull/3063)) – a new channel adapter, likely to be merged after review. Signals interest in federated/email-style messaging.
- **Session config persistence** (PR [#3067](https://github.com/sipeed/picoclaw/pull/3067)) – fixes the UI not saving `dm_scope`. While a bugfix, it touches front-end/back-end config flow; may be included in the next stable release.
- **Agent loop reload stability** (PR [#2904](https://github.com/sipeed/picoclaw/pull/2904)) – open since May 20; if merged, it will improve runtime reliability during provider hot‑reload.

No explicit user feature requests appear in issues today. The community seems focused on stability rather than new features.

### 7. User Feedback Summary
- **Pain points**: RISC‑V users cannot use OpenAI (s0me0ne-unkn0wn); Windows users see console flashes and QQ channel failure (cuandada); Telegram location pins are now fixed (terurium).
- **Satisfaction**: The rapid response to the Telegram bug within 2 days shows good maintainer attention. However, the lack of movement on the RISC‑V issue after 3+ weeks may cause frustration.
- **Use cases**: Deployment on SBCs (Raspberry Pi ARMv7, RISC‑V) and Windows desktop; multi‑channel messaging (Telegram, QQ, DeltaChat emerging).

### 8. Backlog Watch
- **#2887** – RISC‑V `.deb`/OpenAI bug (9 comments, stale since May 17). Needs maintainer triage and a fix PR.
- **PR #2904** – Agent loop reload stability (open since May 20, no comments). Important for runtime safety, but has not been reviewed.
- **#3015** – QQ Windows token timeout (2 days old). No assignee or comment from maintainers yet; should be watched.
- **PR #3063** – DeltaChat gateway (new, large lines of code). Needs thorough review to avoid integration regressions.

All other open PRs from today are small defensive fixes and likely low risk for merge. Maintainer attention should focus on the RISC‑V blocker and the long‑standing agent reload PR.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-09

## Today’s Overview
Activity was moderate, with one open issue and four pull requests updated in the last 24 hours. Three PRs were merged or closed, one remains open, and no new releases were cut. The open issue (#2715) reports a significant functional regression in WhatsApp media handling that blocks agents from accessing inbound attachments. Meanwhile, two security-focused PRs were completed – an opt-in egress lockdown feature and a set of auth/security hardening changes – indicating active infrastructure and safety improvements. Overall, the project shows healthy commit cadence with clear prioritization of both feature work and bug fixes.

## Releases
No new releases were published today.

## Project Progress
Three pull requests were merged or closed:

- **#2716** (closed) – *[follows-guidelines] resolve check test 0609 once* — Author: zybChaitin. A procedural PR that likely addressed CI or test compliance checks. No source code impact described.
- **#2713** (closed) – *feat(security): egress lockdown (opt-in, off by default)* — Author: omri-maya. Implements an optional “egress lockdown” by placing each agent container on a Docker `--internal` network with only the OneCLI gateway accessible. This prevents agents from making direct outbound internet requests unless explicitly allowed through the gateway. Addresses a long-standing concern about agent network isolation.
- **#2712** (closed) – *[follows-guidelines] pull request* — Author: juhojeon86. Another procedural PR, likely a test or documentation update adhering to the contribution guide.

The merged security PRs represent tangible progress in hardening the agent runtime environment.

## Community Hot Topics
The only open issue and the remaining open PR are the most active items:

- **Issue #2715** (open) – *Inbound WhatsApp media (images/docs/audio) is unreachable by the agent* — Author: jon-ruth. Filed yesterday with zero comments but zero reactions. Describes a concrete bug: WhatsApp attachments are saved to a host directory not mounted into the agent container, so the agent receives a non-existent path (`/workspace/attachments/...`). This is a high-impact functional regression affecting all WhatsApp integrations. Underlying need: reliable multi-modal messaging capability without manual workarounds. [GitHub Issue #2715](https://github.com/nanocoai/nanoclaw/issues/2715)

- **PR #2714** (open) – *security: fix four auth/security issues* — Author: JorellDacasin. Proposes binding the webhook server to `127.0.0.1` by default, replacing `Math.random()` with `crypto.randomUUID()` for approval IDs, and other hardening. No comments yet. Addresses implicit security vulnerabilities that could affect production deployments. Underlying need: improved default security posture without breaking existing configurations. [GitHub PR #2714](https://github.com/nanocoai/nanoclaw/pull/2714)

## Bugs & Stability
One bug reported today, ranked **critical**:

- **Issue #2715** – WhatsApp media attachments cannot be accessed by the agent. The root cause is a filesystem mount mismatch: attachments are written to `DATA_DIR/attachments` on the host, but the agent is given a container-internal path that doesn’t exist. This is a full functional regression for WhatsApp media (images, documents, audio). No fix PR exists yet; the issue is still untriaged. **Severity:** Critical (blocks core messaging feature). [GitHub Issue #2715](https://github.com/nanocoai/nanoclaw/issues/2715)

No other crashes or regressions were reported today.

## Feature Requests & Roadmap Signals
The merged **egress lockdown** feature (PR #2713) suggests that container network isolation is a priority for the maintainers. While not a user request per se, it reflects a roadmap direction toward enterprise-grade security. The open **auth/security fixes** (PR #2714) also align with that pattern. No explicit feature requests appeared in today’s data, but the WhatsApp media bug indirectly calls for a better attachment-mounting mechanism – possibly a fix that will appear in a v2.x patch release.

Prediction: The next minor release may include the egress lockdown (opt-in) and the four auth fixes from PR #2714, along with a fix for the WhatsApp media path issue.

## User Feedback Summary
The only direct user feedback is from **jon-ruth** (Issue #2715), who reports clear dissatisfaction: “the agent cannot open images/documents/audio a user sends.” This indicates a real pain point for users relying on WhatsApp as a communication channel. The lack of comments may mean the issue was just filed, but the severity suggests frustration will grow if not addressed quickly. No positive feedback was recorded today.

## Backlog Watch
No long-unanswered issues or PRs are evident from today’s data snapshot. All open items are recent (created within the last ~24 hours). Maintainers should monitor **Issue #2715** closely and ideally triage it within the next 48 hours, as it blocks a primary integration path.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest – 2026-06-09

## 1. Today's Overview

The project saw very high activity over the past 24 hours, with **50 PRs** (27 open, 23 merged/closed) and **11 issues** (7 open, 4 closed). No new releases were published. The bulk of the work continues to focus on the **Reborn** runtime: routing core OpenAI-compatible endpoints (chat completions, responses, streaming) through the new `ProductWorkflow` layer, adding approval and operator parity, and hardening the hook/security-audit subsystem. Several user-reported bugs were fixed the same day (Google Calendar, Codex model discovery), while a configuration issue for the Minimax provider remains open. The project is clearly in an intensive integration phase, with multiple XL-sized PRs converging.

---

## 2. Releases

No new releases were published today. The most recent release candidate remains tracked in the open PR [#3708](https://github.com/nearai/ironclaw/pull/3708) (release chore, open since May 16), which notes breaking changes in `ironclaw_common` and `ironclaw_skills`.

---

## 3. Project Progress

The following notable PRs were merged or closed today (23 closed/merged total):

### Features & enhancements
- **Scoped outbound delivery defaults** – [#4581](https://github.com/nearai/ironclaw/pull/4581) (merged) implements Phase 2 of trigger delivery defaults with versioned preference writes and scope resolution.
- **Automation run history UI** – [#4580](https://github.com/nearai/ironclaw/pull/4580) (merged) persists bounded trigger run history and exposes it through the WebUI.
- **Channel-neutral outbound target authority** – [#4589](https://github.com/nearai/ironclaw/pull/4589) (merged) adds status and revalidation for outbound target preferences.
- **Planner subagent flavor** – [#4572](https://github.com/nearai/ironclaw/pull/4572) (merged) replaces the `researcher` subagent with `planner` and renames `flavor_id` → `subagent_type`.
- **NormalizingProvider decorator** – [#4576](https://github.com/nearai/ironclaw/pull/4576) (merged) adds Layer-3 decorator for tool‑call finish‑reason normalization (RC3/M9 Phase B).
- **System sentinel round-trip fix** – [#4523](https://github.com/nearai/ironclaw/pull/4523) (merged) fixes `TenantId`/`UserId` deserialization for the `\x1fSYSTEM\x1f` sentinel.

### Bug fixes
- **Google Calendar `list_events`** – [#4578](https://github.com/nearai/ironclaw/pull/4578) (merged) defaults `timeMin` to now, fixing “oldest events first” behaviour.
- **Codex model version** – issue [#4564](https://github.com/nearai/ironclaw/issues/4564) was closed; fix is presumed merged (*no corresponding PR in today’s list*).

### Issues closed
- **Composition maintainability** – [#3958](https://github.com/nearai/ironclaw/issues/3958) (closed) splits `hooks.rs` and simplifies loader machinery.
- **Trace Commons onboarding** – [#4560](https://github.com/nearai/ironclaw/issues/4560) (closed) routes HTTP through host egress policy.

### Still open (key in-progress PRs)
- Chat completions routing ([#4495](https://github.com/nearai/ironclaw/pull/4495)), Responses routing ([#4546](https://github.com/nearai/ironclaw/pull/4546)), SSE translation ([#4552](https://github.com/nearai/ironclaw/pull/4552)), agent‑driven Trace Commons onboarding ([#4559](https://github.com/nearai/ironclaw/pull/4559)), and the trajectory observer hook ([#4588](https://github.com/nearai/ironclaw/pull/4588)).

---

## 4. Community Hot Topics

The following issues and PRs received the most comments or are generating active discussion:

- **Security hardening follow-ups** – [#3957](https://github.com/nearai/ironclaw/issues/3957) (2 comments) and [#3959](https://github.com/nearai/ironclaw/issues/3959) (2 comments) address third‑party hook activation and `SecurityAuditSink` adoption. Both are labelled `security-review-required` and have been open since May 23, reflecting the team’s cautious approach before enabling multi‑tenant prod.
- **Minimax provider bug** – [#4587](https://github.com/nearai/ironclaw/issues/4587) (1 comment) reports a configuration failure. The only reply suggests it may be a secret‑store key‑metadata issue; no fix PR yet.
- **Auth evidence with tenant identity** – [#4585](https://github.com/nearai/ironclaw/issues/4585) (1 comment) is a follow‑up from review on PR [#4495](https://github.com/nearai/ironclaw/pull/4495), requesting that `VerifiedAuthClaim` carry tenant identity for tenant‑aware validation.
- **Release PR** – [#3708](https://github.com/nearai/ironclaw/pull/3708) remains open with ongoing discussion about breaking changes and release timing.

*Note: PR comment counts were not provided in the raw data; the above reflects available issue comment counts.*

---

## 5. Bugs & Stability

| ID | Severity | Description | Status | Fix PR? |
|----|----------|-------------|--------|---------|
| [#4587](https://github.com/nearai/ironclaw/issues/4587) | High | Minimax provider configuration fails with `api_key_set=false` because stored key metadata cannot be read. | Open, 1 comment | No PR yet |
| [#4577](https://github.com/nearai/ironclaw/issues/4577) | Medium | `google_calendar list_events` returns oldest events because `timeMin` is not defaulted. | Closed (fix in [#4578](https://github.com/nearai/ironclaw/pull/4578)) | ✅ Merged |
| [#4564](https://github.com/nearai/ironclaw/issues/4564) | Medium | Codex (ChatGPT subscription) hardcodes `client_version=0.111.0`, blocking newer models like `gpt-5.5`. | Closed | ✅ (presumed merged) |
| [#4586](https://github.com/nearai/ironclaw/issues/4586) | Low | Clarification request: timed-out turns should be bounded by turn admission (no second quota system needed). | Open, 0 comments | Discussion only |
| [#4585](https://github.com/nearai/ironclaw/issues/4585) | Low (enhancement) | Auth evidence missing tenant identity; makes tenant‑aware validation impossible. | Open, 1 comment | No PR yet |

The Minimax provider bug is the most impactful unresolved issue today.

---

## 6. Feature Requests & Roadmap Signals

Several issues and PRs point toward the next phase of IronClaw development:

- **Reborn operator parity** – Epic [#4533](https://github.com/nearai/ironclaw/issues/4533) calls for setup, config, diagnostics, and service lifecycle (created June 7). Epic [#4539](https://github.com/nearai/ironclaw/issues/4539) demands approval parity with V1 (approve once, deny, always allow). These are likely to land in the next minor release.
- **Streaming & SSE** – PRs [#4552](https://github.com/nearai/ironclaw/pull/4552) (projection streams to OpenAI SSE) and [#4546](https://github.com/nearai/ironclaw/pull/4546) (Responses through ProductWorkflow) are in progress and will unify the streaming path.
- **Auth & identity** – [#4585](https://github.com/nearai/ironclaw/issues/4585) and [#4186](https://github.com/nearai/ironclaw/pull/4186) (local-dev approval gates) show ongoing work to strengthen security boundaries.
- **Trajectory observation** – PR [#4588](https://github.com/nearai/ironclaw/pull/4588) (draft) enables downstream consumers to observe agent trajectory in real time, a strong signal for observability tooling.
- **Slack outbound targets** – PR [#4590](https://github.com/nearai/ironclaw/pull/4590) adds Slack shared‑channel outbound targets, extending the triggered delivery framework.

**Predictions**: The next release (likely `ironclaw 0.30.x`) will include Reborn operator setup, approvals parity, and OpenAI‑compatible streaming. The `NormalizingProvider` decorator and `planner` subagent flavor are almost certain to ship.

---

## 7. User Feedback Summary

Two user‑reported bugs were resolved swiftly today, indicating a responsive team:

- **BenKurrek** reported that `list_events` in Google Calendar returned outdated events ([#4577](https://github.com/nearai/ironclaw/issues/4577)) and that Codex model discovery was blocked by a hardcoded version ([#4564](https://github.com/nearai/ironclaw/issues/4564)). Both were fixed within hours – a positive user experience.
- **matiasbenary** reported a configuration failure with the Minimax provider ([#4587](https://github.com/nearai/ironclaw/issues/4587)). This is a real pain point – users cannot use the provider at all. The sole reply suggests a secret‑store issue, but no fix has been merged. Dissatisfaction may grow if this remains unresolved.

Overall, the project appears to be user‑focused, with quick bugfix turnaround for common tools, but the Minimax issue needs attention.

---

## 8. Backlog Watch

These older, important items have not seen recent maintainer action and may need attention:

| ID | Created | Age | Issue/PR | Reason for Concern |
|----|---------|-----|----------|-------------------|
| [#3957](https://github.com/nearai/ironclaw/issues/3957) | 2026-05-23 | 17d | Open | Security‑critical: third‑party hook activation hardening. Labelled `security-review-required` but awaiting follow‑ups. |
| [#3959](https://github.com/nearai/ironclaw/issues/3959) | 2026-05-23 | 17d | Open | `SecurityAuditSink` adoption at remaining boundary call sites. |
| [#4186](https://github.com/nearai/ironclaw/pull/4186) | 2026-05-28 | 12d | Open | Local‑dev approval gates – a key piece of Reborn parity. Last updated June 8, but still open. |
| [#3708](https://github.com/nearai/ironclaw/pull/3708) | 2026-05-16 | 24d | Open | Release PR with breaking changes – has been stalled for over three weeks, blocking downstream consumers. |

The two security issues (#3957, #3959) are the highest priority backlog items and should be resolved before enabling multi‑tenant production use of third‑party hooks.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-06-09

## 1. Today's Overview
The project saw exceptionally high activity with **17 pull requests updated** in the last 24 hours, **16 of which were merged or closed**. This indicates a major push of feature development and bug fixes across multiple areas: data migration, cowork notifications, authentication robustness, and OpenClaw gateway exposure. Only **one new issue** was opened, a feature request for Hermes agent support. No new releases were published. The rapid closure of several previously stale PRs suggests the team is actively clearing the backlog while delivering concrete improvements.

## 2. Releases
*No new releases recorded on this date.*

## 3. Project Progress
The following key features and fixes were merged or closed today:

| PR | Title | Area | Status |
|----|-------|------|--------|
| [#2130](https://github.com/netease-youdao/LobsterAI/pull/2130) | feat(cowork): add task completion notifications | renderer, build, docs, main | **Merged** |
| [#2125](https://github.com/netease-youdao/LobsterAI/pull/2125) | feat(data-migration): add user data backup and restore | renderer, main | **Merged** |
| [#2126](https://github.com/netease-youdao/LobsterAI/pull/2126) | fix(data-migration): restore data in place and preserve runtime locks | renderer, main | **Merged** |
| [#2128](https://github.com/netease-youdao/LobsterAI/pull/2128) | fix(data-migration): exclude Network directory from backup | renderer, main | **Merged** |
| [#2129](https://github.com/netease-youdao/LobsterAI/pull/2129) | chore(auth): add login callback diagnostics | renderer, main | **Merged** |
| [#2127](https://github.com/netease-youdao/LobsterAI/pull/2127) | fix(auth): improve Windows focus after callback login | main | **Merged** |
| [#2123](https://github.com/netease-youdao/LobsterAI/pull/2123) | feat(settings): surface OpenClaw gateway URL and refine runtime status | renderer, main, openclaw | **Merged** |
| [#2124](https://github.com/netease-youdao/LobsterAI/pull/2124) | chore: enhance test mode | renderer | **Merged** |

Additionally, **8 stale PRs** (originally created in April) were closed after being addressed or superseded:
- [#1510](https://github.com/netease-youdao/LobsterAI/pull/1510) – Fix scheduled task IM notification without session selection
- [#1514](https://github.com/netease-youdao/LobsterAI/pull/1514) – Fix QQ Bot group whitelist missing UI
- [#1515](https://github.com/netease-youdao/LobsterAI/pull/1515) – Fix log export timeout issue
- [#1517](https://github.com/netease-youdao/LobsterAI/pull/1517) – Fix GitHub Copilot OAuth polling & token loss
- [#1521](https://github.com/netease-youdao/LobsterAI/pull/1521) – Prevent skills-changed from spurious gateway restart
- [#1522](https://github.com/netease-youdao/LobsterAI/pull/1522) – Add dynamic model list fetching from provider API
- [#1524](https://github.com/netease-youdao/LobsterAI/pull/1524) – Detailed error messages on test connection failure
- [#1526](https://github.com/netease-youdao/LobsterAI/pull/1526) – Add session color labeling for cowork conversations

## 4. Community Hot Topics
- **Issue #2131** ([link](https://github.com/netease-youdao/LobsterAI/issues/2131)): *“LobsterAI 支持 hermes agent有计划吗？”* — The only open issue updated today, asking whether Hermes agent support is planned. With 1 comment and 0 reactions, it is a low-engagement but clear signal of interest in extending the agent ecosystem. The question remains unanswered so far.

- **Open dependabot PR #1277** ([link](https://github.com/netease-youdao/LobsterAI/pull/1277)): Bumps Electron from 40.2.1 to 42.3.3 (and electron-builder). Though created in April, it was updated yesterday and remains open. This is a routine dependency update but has been pending for over two months.

## 5. Bugs & Stability
No new bugs were reported today. However, **critical stability improvements** were delivered through the merged PRs:

| Severity | Issue | Fix PR(s) | Description |
|----------|-------|-----------|-------------|
| High | Scheduled task IM notifications fail silently | [#1510](https://github.com/netease-youdao/LobsterAI/pull/1510) | Form allowed submission without selecting a conversation channel |
| High | Log export times out after 30 seconds | [#1515](https://github.com/netease-youdao/LobsterAI/pull/1515) | Serial compression of large logs exceeded timeout; now fixed |
| Medium | GitHub Copilot OAuth token lost if settings closed during polling | [#1517](https://github.com/netease-youdao/LobsterAI/pull/1517) | Missing cleanup on unmount |
| Medium | QQ Bot group whitelist unusable via UI | [#1514](https://github.com/netease-youdao/LobsterAI/pull/1514) | Input field and add button were missing |
| Low | Spurious OpenClaw gateway restart on skills change | [#1521](https://github.com/netease-youdao/LobsterAI/pull/1521) | State change handler was too broad |

Data migration fixes ([#2125](https://github.com/netease-youdao/LobsterAI/pull/2125), [#2126](https://github.com/netease-youdao/LobsterAI/pull/2126), [#2128](https://github.com/netease-youdao/LobsterAI/pull/2128)) also enhance reliability by preserving runtime locks and excluding network state from backups.

## 6. Feature Requests & Roadmap Signals
- **Hermes agent support** — Issue [#2131](https://github.com/netease-youdao/LobsterAI/pull/2131) explicitly asks about Hermes agent. Given the project’s focus on extensibility, this may be considered for an upcoming minor release.
- **Data backup & restore** ([#2125](https://github.com/netease-youdao/LobsterAI/pull/2125)) — A major portability feature that could land in the next patch release.
- **Dynamic model list fetching** ([#1522](https://github.com/netease-youdao/LobsterAI/pull/1522)) — Will reduce manual configuration effort and keep providers up-to-date.
- **Cowork session color labels** ([#1526](https://github.com/netease-youdao/LobsterAI/pull/1526)) — A user-requested UX improvement.
- **OpenClaw gateway URL exposure** ([#2123](https://github.com/netease-youdao/LobsterAI/pull/2123)) — Helps advanced users integrate or troubleshoot the gateway.

These features, combined with the task completion notifications ([#2130](https://github.com/netease-youdao/LobsterAI/pull/2130)), suggest the roadmap is currently weighted toward **data management, cowork enhancements, and provider flexibility**.

## 7. User Feedback Summary
While no direct user feedback threads were updated today, the merged bug fixes directly address several long-standing pain points:
- *“Log export always times out”* (fixed)
- *“QQ Bot group allowlist can’t be configured via UI”* (fixed)
- *“Scheduled tasks with IM notification silently fail”* (fixed)
- *“Copilot OAuth token is lost when I close settings”* (fixed)
- *“Model list is outdated / I have to add new models manually”* (fixed via dynamic fetching)

These fixes indicate that users have been experiencing real friction in daily workflows, and the project is now systematically resolving them. The one remaining open issue (#2131) signals a desire for more agent types, suggesting the community sees value in expanding beyond current integrations.

## 8. Backlog Watch
- **PR #1277** ([link](https://github.com/netease-youdao/LobsterAI/pull/1277)) – Dependabot update for Electron and electron-builder (from v40 to v42). Open since **April 2, 2026** (68+ days). This is a major dependency bump that could contain security fixes and performance enhancements. Needs maintainer review and merge to prevent falling far behind the Electron release cycle.
- No other issues or PRs remain unanswered for an extended period; the recent closure of 8 stale PRs shows the team is actively managing the backlog.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

# TinyClaw Project Digest – 2026-06-09

## 🧭 Today's Overview

The TinyClaw repository saw minimal activity in the last 24 hours. No new issues were reported or updated, and no pull requests were merged or closed. One open pull request (#280) was updated, focusing on an installation fix for the native `better-sqlite3` dependency. No new releases were published. Overall, the project appears to be in a quiet phase with no active bug reports or feature discussions.

## 🚀 Releases

*None.* No new versions were released in the last 24 hours.

## 📈 Project Progress

No pull requests were merged or closed today. The only activity is an ongoing pull request that aims to improve developer experience.

---

## 🔥 Community Hot Topics

- **#280 – [OPEN] fix(install): add postinstall script to auto-rebuild better-sqlite3**  
  *Author: dsy122 | Created: 2026-06-08 | Updated: 2026-06-08 | Comments: 0 | 👍: 0*  
  [View PR](https://github.com/TinyAGI/tinyagi/pull/280)  
  This pull request addresses a recurring pain point for users when setting up the project after a fresh `npm install`. The `better-sqlite3` native addon requires compilation for the current Node.js runtime, and without an automatic rebuild, users encounter errors. The proposed fix adds a `postinstall` script to eliminate the manual `npm rebuild` step. Although there are no comments or reactions yet, the PR directly targets a common setup friction point.

## 🐛 Bugs & Stability

*No new bugs were reported today.* However, the underlying issue described in PR #280 (installation errors due to missing `better-sqlite3` rebuild) is a known stability concern for new contributors. The fix is under review; once merged, it will resolve this installation regression.

**Severity (if active):** Moderate – affects all new developers or CI setups, but does not impact running systems.

## 💡 Feature Requests & Roadmap Signals

No new feature requests were submitted in the last 24 hours. The only actionable item is the installation improvement in PR #280, which is more of a developer experience fix than a feature. No roadmap signals were detected.

## 💬 User Feedback Summary

Based on the context of PR #280, user feedback is implicit: new contributors or users performing a fresh install encounter build errors because `better-sqlite3` must be manually rebuilt. The frustration is that the default `npm install` step does not handle this automatically, breaking the out-of-the-box experience. The proposed script aims to resolve this dissatisfaction without requiring user intervention.

## 📋 Backlog Watch

- **No long-unanswered issues or PRs** currently warrant maintainer attention.  
  The only open item (#280) was created less than 24 hours ago and has received no comments.

---

*Generated from TinyClaw repository data (github.com/TinyAGI/tinyagi) – 2026-06-09.*

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest — 2026-06-09

## Today's Overview

The CoPaw repository saw very high activity in the past 24 hours, with 47 pull requests and 21 issues updated. 13 issues were closed and 24 PRs were merged or closed, indicating a strong flow of bug fixes and feature work. No new releases were published. The community was engaged on feature suggestions (notably Hermes Agent integration) and several Windows-specific stability problems. Overall, the project appears in a healthy state with frequent, responsive maintainer attention.

---

## Releases

None in this period.

---

## Project Progress

The following notable PRs were merged or closed today, representing concrete progress:

- **Yuanbao typing indicator** — [#5020](https://github.com/agentscope-ai/CoPaw/pull/5020): Adds a heartbeat-based typing indicator for the Yuanbao channel, improving real-time user feedback.
- **OneBot graceful degradation on port conflict** — [#5010](https://github.com/agentscope-ai/CoPaw/pull/5010): Fixes a crash during zero-downtime reload when the WebSocket port is still in use, preventing permanent channel failure.
- **DingTalk cross-user message merging fix** — [#4932](https://github.com/agentscope-ai/CoPaw/pull/4932): Corrects a bug that batch-merged messages from different users when conversation_id suffixes collided.
- **Plugin extension infrastructure (Console)** — [#4997](https://github.com/agentscope-ai/CoPaw/pull/4997): Implements a unified frontend extension point registration mechanism for plugins (menu, routes, slots).
- **DataPaw plugin** — [#4622](https://github.com/agentscope-ai/CoPaw/pull/4622) (still open but actively updated): Adds a data-analysis plugin with 12 BI skills, aimed at bundle inclusion.

Additionally, several first-time-contributor PRs were opened for small bug fixes (empty msg list guard [#5038](https://github.com/agentscope-ai/CoPaw/pull/5038), empty Exec= line parsing [#5037](https://github.com/agentscope-ai/CoPaw/pull/5037), llama.cpp version parsing [#5035](https://github.com/agentscope-ai/CoPaw/pull/5035)), reflecting growing community contributions.

---

## Community Hot Topics

- **Issue #5017 — Hermes Agent feature suggestion** [[link](https://github.com/agentscope-ai/CoPaw/issues/5017)]  
  _7 comments, 2 👍, open._ A user strongly advocates for CoPaw to learn from Hermes Agent's "learning loop" capabilities, where agents automatically create and iterate on skills. This resonated with the community, indicating a desire for self-improving agent behavior.

- **Issue #5003 — Aliyun coding plan Qwen3.7‑plus hangs** [[link](https://github.com/agentscope-ai/CoPaw/issues/5003)]  
  _8 comments, closed._ User reported that using Aliyun's coding plan with Qwen3.7‑plus causes indefinite stalling. The issue was quickly resolved (closed within 24h), showing responsive triage.

- **Issue #4477 — WeChat iLink token expiry** [[link](https://github.com/agentscope-ai/CoPaw/issues/4477)]  
  _15 comments, closed._ Detailed bug report about missing retry logic when context_token expires; also lacked logging for file/image send failures. Closed, presumably fixed.

- **Issue #5030 — Duplicate responses in WeChat channel with active mode** [[link](https://github.com/agentscope-ai/CoPaw/issues/5030)]  
  _1 comment, open._ A user reports that after enabling "active mode", a single question yields two similar but not identical replies. Root cause not yet identified.

---

## Bugs & Stability

Several bugs were reported or updated today. Ranked by severity:

| Issue | Severity | Description | Fix PR Exist? |
|-------|----------|-------------|---------------|
| [#5030](https://github.com/agentscope-ai/CoPaw/issues/5030) | High | Duplicate message responses in WeChat channel when active mode is on. May confuse users and waste API calls. | None yet |
| [#5029](https://github.com/agentscope-ai/CoPaw/issues/5029) | High | Pet feature crashes and severe lag; also default agent cannot be changed (hardcoded `default`). Affects core UX. | None yet |
| [#5025](https://github.com/agentscope-ai/CoPaw/issues/5025) | Medium | `submit_to_agent` fails with `FileNotFoundError` due to duplicated `session_id` in filename when user_id defaults to session_id. | Yes — [#5036](https://github.com/agentscope-ai/CoPaw/pull/5036) and [#5026](https://github.com/agentscope-ai/CoPaw/pull/5026) address the same root cause |
| [#5031](https://github.com/agentscope-ai/CoPaw/issues/5031) | Low | Skill slash invocation displays expanded `SKILL.md` content in Console instead of the intended response. Cosmetic but confusing. | None yet |
| [#2777](https://github.com/agentscope-ai/CoPaw/issues/2777) | Medium | GPT-5.x models fail with `max_tokens` parameter error; hardcoded model list not fetching from API. Open since April 1. | None yet |
| [#5019](https://github.com/agentscope-ai/CoPaw/issues/5019) | Medium | Memory compaction crashes with `AttributeError` when `source` is a string instead of dict. Closed (fix likely merged). | Closed |
| [#5034](https://github.com/agentscope-ai/CoPaw/issues/5034) | Medium | MCP tool names containing dots cause OpenAI API 400 errors. Closed — fix already in later versions. | Closed |

Two fix PRs for the `session_id` duplication bug ([#5036](https://github.com/agentscope-ai/CoPaw/pull/5036), [#5026](https://github.com/agentscope-ai/CoPaw/pull/5026)) are open and ready for review. A security fix for keychain master key isolation ([#5028](https://github.com/agentscope-ai/CoPaw/pull/5028)) was also opened.

---

## Feature Requests & Roadmap Signals

- **Learning Loop / Skill Auto‑iteration** — Issue [#5017](https://github.com/agentscope-ai/CoPaw/issues/5017) has attracted positive reactions, asking CoPaw to absorb Hermes Agent's "learning loop" for automatic skill creation. This is likely to be discussed for a future release.
- **Observability/Tracing Integration** — Issue [#5009](https://github.com/agentscope-ai/CoPaw/issues/5009) requests integration with Langfuse, OpenTelemetry, etc. for token usage, latency, and cost attribution. The maintainer has not yet answered, but the lack of built-in tracing is a gap for production deployments.
- **Plugin Market Tab** — PR [#5023](https://github.com/agentscope-ai/CoPaw/pull/5023) (open) adds a graphical plugin marketplace integrated with AgentScope Platform, indicating a push toward an ecosystem.
- **Goal Mode** — PR [#4443](https://github.com/agentscope-ai/CoPaw/pull/4443) (open since May 16) adds lightweight session‑scoped standing objectives (`/goal`). Still under review.
- **TUI Bundling** — PR [#5032](https://github.com/agentscope-ai/CoPaw/pull/5032) (opened today) integrates the terminal chat UI into the main `qwenpaw` binary, Phase 1 of the CLI‑TUI design.
- **Desktop Auto‑Updater** — PR [#4669](https://github.com/agentscope-ai/CoPaw/pull/4669) (open) adds Tauri‑based auto‑update for desktop builds.

Given the volume, the next minor release may include the Plugin Market tab, `session_id` bugfixes, and Windows desktop stability patches.

---

## User Feedback Summary

**Positive:**  
- "体验非常好" ("Great experience") from a long‑time user in [#5017](https://github.com/agentscope-ai/CoPaw/issues/5017), praising localisation and ease of use.
- Many issues are closed quickly, indicating good maintainer responsiveness.

**Pain Points:**  
- **Performance:** Several Windows users report severe lag and high CPU usage during streaming or task execution ([#5015](https://github.com/agentscope-ai/CoPaw/issues/5015), [#4792](https://github.com/agentscope-ai/CoPaw/issues/4792)). Also the Pet feature crashes and stutters ([#5029](https://github.com/agentscope-ai/CoPaw/issues/5029)).
- **WeChat / Channel reliability:** Token expiry without retry ([#4477]), duplicate responses in active mode ([#5030]), and plugin tools not auto‑discovered in WeCom channels ([#4585]) erode trust for multi‑channel deployments.
- **MCP tool naming:** Users hitting API validation errors due to dots in tool names ([#4918], [#5034]) — a recurring theme.
- **Agent selection hardcoded:** Users want to customize which agent is loaded by default, but `default` is still forced in some flows ([#5029]).

---

## Backlog Watch

- **Issue #2777 (GPT-5.x max_tokens error)** — Open since **April 1, 2026** (over 2 months). The model list is hardcoded and fails with newer models. No PR yet. This blocks users who want to use GPT-5.x models directly.
- **Issue #5009 (Observability/Tracing)** — Only 2 comments, but zero maintainer response. Needs triage as it's a legitimate enterprise requirement.
- **PR #4443 (Goal Mode)** — Open since May 16, last updated June 8. No review feedback visible. If maintainers intend to include it, a decision is needed.
- **PR #4669 (Desktop Auto‑Updater)** — Open since May 25, dependent on #3813. This is a critical feature for desktop users; lack of progress may frustrate those on older builds.
- **Issue #4670** (not shown but implied by PR #4669 dependency) — We recommend the maintainers review and merge the desktop auto‑updater soon to avoid a growing support burden from manual update requests.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-06-09

## 1. Today's Overview

The ZeroClaw project saw **high activity** in the last 24 hours, with **17 issues** (15 open, 2 closed) and **50 PRs** (40 open, 10 merged/closed). No new releases were published. Development continues across multiple parallel milestones (v0.8.0, v0.8.1, v0.8.2, v0.8.3), with significant focus on security, plugin infrastructure, and channel reliability. Notable today: a critical Matrix session isolation fix was merged, a WASM plugin path reconciliation PR was opened, and several new bugs were reported around web-fetch private-host resolution and CI tooling gaps.

## 2. Releases

**None** in the last 24 hours.

## 3. Project Progress (Merged/Closed PRs)

Four PRs were merged/closed in the last 24 hours:

- **[PR #6487 – Bug: Multi-alias Matrix instances share one session store](https://github.com/zeroclaw-labs/zeroclaw/pull/6487)** (closed) — Fixed a `S1` blocker where multiple Matrix channel aliases under one daemon silently shared a single session store, causing wrong account deliveries. The fix isolates state per alias and repairs key backups.

- **[PR #6148 – feat(hardware): smart-room ESP32 demo with Telegram + simulator](https://github.com/zeroclaw-labs/zeroclaw/pull/6148)** (closed) — Hackathon project; an end-to-end demo integration of Telegram bot with ZeroClaw controlling an ESP32 smart-room simulator.

- **[PR #6225 – Feature: Smart Truncation for Telegram](https://github.com/zeroclaw-labs/zeroclaw/pull/6225)** (closed) — Added Markdown-aware message truncation for Telegram to avoid ugly codeblock splits.

- **[PR #7388 – fix(matrix): isolate session state per channel alias and repair key backup](https://github.com/zeroclaw-labs/zeroclaw/pull/7388)** (closed) — A breaking change that moves Matrix session stores to per-alias paths (manual migration required). Resolves the session-clobbering issue #6487.

Additional newly opened PRs that are not yet merged but signal active development: [#7413](https://github.com/zeroclaw-labs/zeroclaw/pull/7413) (WASM plugin path reconciliation), [#7365](https://github.com/zeroclaw-labs/zeroclaw/pull/7365) (documentation rework), [#7367](https://github.com/zeroclaw-labs/zeroclaw/pull/7367) (webhook per-alias routing), [#7385](https://github.com/zeroclaw-labs/zeroclaw/pull/7385) (observability turn metadata), [#7348](https://github.com/zeroclaw-labs/zeroclaw/pull/7348) (cron catch-up fix).

## 4. Community Hot Topics

Issues with the most discussion (by comment count) in the last 24 hours:

- **[Issue #6699 – tool_filter_groups is a no-op for real MCP tools (prefix-check bug)](https://github.com/zeroclaw-labs/zeroclaw/issues/6699)** (7 comments) — Two distinct bugs: the MCP-tool prefix mismatch in the filter gate, and no integration with deferred loading. High-risk, accepted, and actively discussed. Underlying need: reliable tool filtering for MCP surfaces.

- **[Issue #5844 – [Bug]: Too much emphasis on memory](https://github.com/zeroclaw-labs/zeroclaw/issues/5844)** (5 comments) — System prompt over-prioritizes memories over current context, especially in cron jobs. Users want configurable memory emphasis. Accepted, high-risk.

- **[Issue #6254 – WASM plugin install path and runtime scan path diverge](https://github.com/zeroclaw-labs/zeroclaw/issues/6254)** (4 comments) — CLI-installed plugins are invisible to the agent because install and discovery scan different directories. A fix PR [#7413](https://github.com/zeroclaw-labs/zeroclaw/pull/7413) was opened today to address this.

- **[Issue #4832 – Add config option to disable LeakDetector high-entropy token redaction](https://github.com/zeroclaw-labs/zeroclaw/issues/4832)** (4 comments) — False positives on MD5 hashes, random filenames, etc. Users request a config toggle. Accepted, p2.

- **[Issue #7142 – Expose security enforcement layer as a pluggable provider interface](https://github.com/zeroclaw-labs/zeroclaw/issues/7142)** (4 comments) — Umbrella tracking issue for v0.9.0. Community interest in extensible security.

**Underlying needs**: Users are pushing for more controllable security (redaction, pluggable providers), better tool filter reliability, and smoother plugin experiences. The high engagement on #6699 and #5844 indicates pain points in core agent behavior that affect daily usage.

## 5. Bugs & Stability

**New bugs reported today (2026-06-09):**

| Issue | Severity | Summary |
|-------|----------|---------|
| [#7412](https://github.com/zeroclaw-labs/zeroclaw/issues/7412) | S2 - degraded | `web_fetch`/`http_request` rejects domain names that resolve to private IPs even when `allowed_private_hosts = ["*"]` is set. No fix PR yet. |
| [#7409](https://github.com/zeroclaw-labs/zeroclaw/issues/7409) | S2 - degraded | Clippy lint job runs only on Linux, allowing Windows/macOS-gated code to accumulate lint errors. No fix PR yet. |
| [#7408](https://github.com/zeroclaw-labs/zeroclaw/issues/7408) | S2 - degraded | CI labeler still misses runtime crate paths for several tool sub-labels (cron, etc.). A fix PR exists (same issue is a PR). |

**Other active high-risk bugs still open (updated in last 24h):**

- [#6699](https://github.com/zeroclaw-labs/zeroclaw/issues/6699) (S2, tool_filter_groups no-op)
- [#5844](https://github.com/zeroclaw-labs/zeroclaw/issues/5844) (S2, memory emphasis)
- [#6254](https://github.com/zeroclaw-labs/zeroclaw/issues/6254) (S2, plugin path divergence – fix PR [#7413](https://github.com/zeroclaw-labs/zeroclaw/pull/7413) opened)
- [#6645](https://github.com/zeroclaw-labs/zeroclaw/issues/6645) (S2, SkillImprover only handles `SKILL.toml`)
- [#4873](https://github.com/zeroclaw-labs/zeroclaw/issues/4873) (S3, Feishu calls LLM instead of Agent)

**Regression signal**: The newly reported #7412 is a subtle security/sandbox bypass that could break legitimate use cases; no immediate fix. The CI coverage gap (#7409) threatens cross-platform quality.

## 6. Feature Requests & Roadmap Signals

**Feature requests filed today:**

- [#7410](https://github.com/zeroclaw-labs/zeroclaw/issues/7410) – Read gateway webhook signing secrets from AppState at handler time instead of caching at startup (enhancement, follow-up from #7367). Suggests safer secret management.

**Ongoing roadmap trackers (updated today):**

- [#7112](https://github.com/zeroclaw-labs/zeroclaw/issues/7112) – v0.8.0 release queue & Stable-tier blockers (config, tool-call-parser)
- [#6970](https://github.com/zeroclaw-labs/zeroclaw/issues/6970) – v0.8.1 integration/channel/provider/tool PR queue
- [#7314](https://github.com/zeroclaw-labs/zeroclaw/issues/7314) – v0.8.2 WASM plugin program
- [#7320](https://github.com/zeroclaw-labs/zeroclaw/issues/7320) – v0.8.3 MCP dashboard and web/plugin-management surfaces

**Predictions for next version (v0.8.0):** Likely includes the config and tool-call-parser Stable-tier promotion, the Matrix session isolation fix (already merged), and possibly the pluggable security provider interface (#7142) if scoped. The WASM plugin reconciliation (#7413) is a strong candidate for v0.8.1. Webhook per-alias routing (#7367) and observability turn metadata (#7385) may land in v0.8.x.

## 7. User Feedback Summary

**Reported pain points:**

- **Memory over-emphasis** (#5844): “Gives too much value to memories … I have to be very explicit to override it.” — User wants configurable memory weight.
- **MCP tool filter broken** (#6699): “tool_filter_groups has zero effect” — User discovers documented feature doesn’t work for real MCP tools.
- **Plugin install confusion** (#6254): “Plugin install copies to wrong directory, agent can’t find it” — Multiple users affected.
- **LeakDetector false positives** (#4832): “Redacts MD5-hashed filenames” — User requests an opt-out toggle.
- **Feishu integration** (#4873): “Only LLM is called, not Agent” — User expects full agent behavior.

**Satisfaction signals:** The community is actively contributing fixes (e.g., smart truncation, ESP32 demo, session isolation). The large number of PRs (50) and tracker issues suggests a healthy, fast-moving project with clear milestones.

**Dissatisfaction signals:** Several high-risk bugs have been open for weeks (e.g., #5844 since April 17, #6254 since May 1, #6699 since May 16) without being resolved, which may frustrate users relying on these features. However, many received updates today, indicating maintainer attention.

## 8. Backlog Watch

**Notable long-open issues (30+ days) that remain active and important:**

- [#4832](https://github.com/zeroclaw-labs/zeroclaw/issues/4832) – *Add config option to disable LeakDetector redaction* (Mar 27, 72 days open) – Still accepted without a fix PR. 4 comments, last updated June 8.
- [#4873](https://github.com/zeroclaw-labs/zeroclaw/issues/4873) – *Feishu calls LLM instead of Agent* (Mar 28, 71 days) – Unresolved, last updated June 8.
- [#5844](https://github.com/zeroclaw-labs/zeroclaw/issues/5844) – *Too much emphasis on memory* (Apr 17, 53 days) – Updated today, no fix PR yet.
- [#6254](https://github.com/zeroclaw-labs/zeroclaw/issues/6254) – *WASM plugin install/detect path divergence* (May 1, 39 days) – Fix PR [#7413](https://github.com/zeroclaw-labs/zeroclaw/pull/7413) opened today.
- [#6645](https://github.com/zeroclaw-labs/zeroclaw/issues/6645) – *SkillImprover/skill_manage only handle SKILL.toml* (May 14, 26 days) – Accepted but no fix.

**Tracker issues that are overdue for resolution:**

- [#7112](https://github.com/zeroclaw-labs/zeroclaw/issues/7112) – v0.8.0 release blockers (since June 2). Although only 7 days old, the milestone has many open sub-issues that need completion before a release can ship.
- [#6970](https://github.com/zeroclaw-labs/zeroclaw/issues/6970) – v0.8.1 queue (since May 27) – Coordination issue with many linked PRs in flight.

**Maintainers appear active** on these issues (many updated today), so the backlog is being managed. The next week will be critical to see whether the long-standing bugs receive fix PRs or remain open.

</details>