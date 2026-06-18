# OpenClaw Ecosystem Digest 2026-06-18

> Issues: 252 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-18 03:43 UTC

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

# OpenClaw Project Digest – 2026-06-18

## 1. Today’s Overview
OpenClaw is experiencing extremely high activity: **252 issues** and **500 pull requests** were updated in the last 24 hours, with **10 issues closed** and **91 PRs merged or closed**. The project remains in a heavy development and triage phase, with a strong focus on security hardening, session-state reliability, and multi-channel delivery correctness. Several critical (P1/platinum‑hermit) bugs are open around message loss, OAuth failures, and crash loops, while a large backlog of feature requests continues to accumulate. No new releases were published today.

---

## 2. Releases
*No new releases were published on 2026-06-18.*

---

## 3. Project Progress
**91 PRs were merged or closed today**, reflecting substantial forward momentum. Notable merged/closed items include:

- [#94396](https://github.com/openclaw/openclaw/issues/94396) (closed) – **Sandbox skill sync fallback** now uses config‑resolved workspace instead of a hardcoded default.
- [#68936](https://github.com/openclaw/openclaw/issues/68936) (closed) – **PR review autofix pipeline + Windows daemon** – a large (XL) script that adds automated review‑comment resolution and a Windows background supervisor.
- [#94309](https://github.com/openclaw/openclaw/issues/94309) (closed) – Telegram “Quote & Reply” support confirmed as OpenClaw-specific and fixed.

Among the **409 open PRs**, several have reached “ready for maintainer look” status, including:
- [#93276](https://github.com/openclaw/openclaw/pull/93276) – Prevents tool‑discovery plugin loads from wiping active providers (P1, platinum hermit).
- [#94343](https://github.com/openclaw/openclaw/pull/94343) – Fixes plugin method scope resolution across all surfaces (closes #92044, P1).
- [#94020](https://github.com/openclaw/openclaw/pull/94020) – Resolves documentation contradiction around `networkidle` in browser tools.

Many small‑to‑medium fixes were merged for agent loop abortion, cron job defaults, memory indexer alignment, and TUI local-mode handling.

---

## 4. Community Hot Topics
The most active discussions (by comment count and reactions) highlight three core community concerns:

| Issue | Comments | Topic | Community Pulse |
|-------|----------|-------|-----------------|
| [#25592](https://github.com/openclaw/openclaw/issues/25592) | 32 | Text between tool calls leaks to messaging channels | **Highest priority** – a security/UX bug causing internal processing text to appear in Slack, iMessage, etc. Labelled `platinum hermit` with `impact:security`. |
| [#9443](https://github.com/openclaw/openclaw/issues/9443) | 25 | Request for prebuilt Android APK releases | Strong demand from mobile users; 2 👍. Despite being P2, it has been open since February. |
| [#10659](https://github.com/openclaw/openclaw/issues/10659) | 13 | Masked secrets to prevent agents from seeing raw API keys | 4 👍, P1 security feature with `impact:auth-provider`. Several related issues on secret handling. |
| [#42475](https://github.com/openclaw/openclaw/issues/42475) | 12 | Per-agent cost budget enforcement at gateway level | Operators wanting to cap spend without external monitoring. |
| [#7707](https://github.com/openclaw/openclaw/issues/7707) | 12 | Memory trust tagging by source | Preventing memory poisoning from untrusted web scrapes. |
| [#6731](https://github.com/openclaw/openclaw/issues/6731) | 12 | Safe/unsafe ClawdBot (Rust rewrite proposal) | Controversial feature request to rewrite in Rust for safety. |
| [#13583](https://github.com/openclaw/openclaw/issues/13583) | 11 | Pre‑response enforcement hooks (hard gates) | High‑stakes workflows (finance, security) demanding mechanical enforcement. |

**Underlying needs:** The community is demanding **security by default** – preventing accidental data leaks, controlling tool execution, and isolating untrusted content. There is also a strong desire for **production readiness** (cost budgets, deployment guides, crash resilience) and **mobile/Android support**.

---

## 5. Bugs & Stability
Several high‑severity bugs remain open or were newly reported today. They are ranked by severity (P1 / platinum hermit first):

### P1 / Platinum Hermit (critical)
- [#25592](https://github.com/openclaw/openclaw/issues/25592) – **Text between tool calls leaks** → internal processing text routed to messaging channels. No fix PR yet.
- [#86215](https://github.com/openclaw/openclaw/issues/86215) – **Codex OAuth refresh failures wedge agents** for hours without alerting or profile rotation. Still needs a fix.
- [#41372](https://github.com/openclaw/openclaw/issues/41372) – **25 production findings** from a 4‑week self‑hosted deployment (config crashes, CLI docs, Discord, cron). Field report with actionable issues.
- [#45224](https://github.com/openclaw/openclaw/issues/45224) – **Unhandled Playwright assertion crashes Gateway** process. Crash loop impact.
- [#41949](https://github.com/openclaw/openclaw/issues/41949) – **Browser interactions exhaust model context** by injecting too much page content.
- [#41346](https://github.com/openclaw/openclaw/issues/41346) – **Externally registered cron jobs auto‑enable** without consent, inherit expensive default model, no circuit breaker.
- [#93858](https://github.com/openclaw/openclaw/issues/93858) – **Foreground reply fence interleaving** – new bug opened yesterday; tracks three states of the issue.

### P1 / Diamond Lobster (high)
- [#41744](https://github.com/openclaw/openclaw/issues/41744) – Feishu image read tool loses media before delivery.
- [#41165](https://github.com/openclaw/openclaw/issues/41165) – Telegram DMs still routing to main session after #40519.
- [#40255](https://github.com/openclaw/openclaw/issues/40255) – User‑configured heartbeat prompt no longer respected (regression).
- [#44502](https://github.com/openclaw/openclaw/issues/44502) – Discord routing/mention‑gating regression.
- [#43374](https://github.com/openclaw/openclaw/issues/43374) – All LLM API calls time out simultaneously under multi‑agent concurrency.
- [#42510](https://github.com/openclaw/openclaw/issues/42510) – Google Chat `replyToMode: "off"` has no effect.
- [#42157](https://github.com/openclaw/openclaw/issues/42157) – Inbound Telegram webhook payloads lost on gateway restart.
- [#41201](https://github.com/openclaw/openclaw/issues/41201) – Control UI avatar not displaying (regression).

### P2 / Notable
- [#43747](https://github.com/openclaw/openclaw/issues/43747) – Memory management in chaos (regression, 3 users reporting inconsistent behavior).
- [#42273](https://github.com/openclaw/openclaw/issues/42273) – `backup create` stalls on large installations (4GB+).
- [#41899](https://github.com/openclaw/openclaw/issues/41899) – Plugin circuit breaker for graceful degradation (still open).

**Fix PRs exist for some bugs:**
- [#94344](https://github.com/openclaw/openclaw/pull/94344) → Fixes memory‑core OpenAI embedding SSRF policy.
- [#94355](https://github.com/openclaw/openclaw/pull/94355) → Fallback to generic embedding provider registry.
- [#94378](https://github.com/openclaw/openclaw/pull/94378) → Fixes image‑generation parsing of invalid entries.
- [#93276](https://github.com/openclaw/openclaw/pull/93276) → Stops plugin discovery from clearing active providers.
- [#94354](https://github.com/openclaw/openclaw/pull/94354) → Handles `/status` and `/compact` in TUI local mode.
- [#94044](https://github.com/openclaw/openclaw/pull/94044) → Prevents Matrix gateway halt on decryption failures.
- [#94235](https://github.com/openclaw/openclaw/pull/94235) → Allows cron lane silent replies.

---

## 6. Feature Requests & Roadmap Signals
The following features have strong community support (👍) or are frequently requested. The next minor release (expected within 2–4 weeks) may include:

| Feature | Issue | 👍 | Likelihood |
|---------|-------|----|------------|
| **Prebuilt Android APK releases** | [#9443](https://github.com/openclaw/openclaw/issues/9443) | 2 | High – mobile usage is growing, and the source code already exists. |
| **Masked secrets (agents cannot see raw keys)** | [#10659](https://github.com/openclaw/openclaw/issues/10659) | 4 | High – aligns with security focus; many related PRs. |
| **Per‑agent cost budget enforcement** | [#42475](https://github.com/openclaw/openclaw/issues/42475) | 1 | Medium – operators are asking for it, but no PR yet. |
| **Memory trust tagging by source** | [#7707](https://github.com/openclaw/openclaw/issues/7707) | 0 | Medium – addresses memory poisoning concerns. |
| **Pre‑response enforcement hooks** | [#13583](https://github.com/openclaw/openclaw/issues/13583) | 2 | Medium – high demand from finance/security users. |
| **Session snapshots (save/load)** | [#13700](https://github.com/openclaw/openclaw/issues/13700) | 0 | Low – useful but not critical. |
| **Theme customization system** | [#28300](https://github.com/openclaw/openclaw/issues/28300) | 5 | High – control UI theme studio with presets; strong visual feedback. |
| **TUI Shift+Enter for newline** | [#10118](https://github.com/openclaw/openclaw/issues/10118) | 4 | Medium – UX improvement for TUI power users. |
| **Activity‑aware run timeout** | [#41588](https://github.com/openclaw/openclaw/issues/41588) | 1 | Medium – prevents killing active multi‑tool chains. |
| **AWS deployment guide** | [#13597](https://github.com/openclaw/openclaw/issues/13597) | 3 | High – reduces adoption friction. |

**Prediction for next release (v2026.7.x):**
- Masked secrets (likely as part of a broader security pass).
- Android APK builds (CI integration already in progress).
- Theme customization (Control UI improvements).
- Fixes for the top P1 bugs (text leak, OAuth refresh, Telegram routing).

---

## 7. User Feedback Summary

**Pain points voiced today:**
- “Memory management is in chaos” – inconsistent behaviour across machines (#43747).
- “All LLM API calls time out simultaneously” under concurrency (#43374).
- “Text between tool calls leaks to messaging channels” – unacceptable in production (#25592).
- “Control UI avatar not displaying” – broken regression (#41201).
- “Gateway crashes on unhandled Playwright assertion” – crash loop (#45224).
- “Backup create stalls on large installations” – silent failure (#42273).
- “Heartbeat prompt no longer respected” – regression (#40255).
- “Field report: 25 findings from 4 weeks of self‑hosted use” – comprehensive pain list (#41372).

**Satisfaction signals:**
- High engagement on feature requests (themes, Android, cost budgets) shows a healthy, invested community.
- Several PRs submitted by external contributors (yuxuan-7814, liuhao1024, Pandah97, zhiqiang26) indicate a welcoming development environment.
- “ClawSweeper” bot is actively fixing issues (#75403), suggesting automated quality tooling is paying off.

**Use cases driving feedback:**
- **Self‑hosted production deployments** (multi‑agent, multi‑channel, large `.openclaw` directories).
- **Finance/security workflows** requiring hard gates and cost enforcement.
- **Mobile and cross‑platform usage** (Android, Telegram, Slack, Feishu).
- **TUI power users** wanting multi‑line input and persistent configuration.

---

## 8. Backlog Watch
Several important issues and PRs have been awaiting maintainer action for months:

| Item | Created | Age | Status | Why Stuck |
|------|---------|-----|--------|-----------|
| [#25592](https://github.com/openclaw/openclaw/issues/25592) – Text leak to channels | 2026-02-24 | ~4 months | Open, P1 | Multiple `clawsweeper:needs-*` tags – awaits maintainer review and security review. |
| [#9443](https://github.com/openclaw/openclaw/issues/9443) – Android APK | 2026-02-05 | >4 months | Open, P2 | Needs maintainer review and product decision. |
| [#10659](https://github.com/openclaw/openclaw/issues/10659) – Masked secrets | 2026-02-06 | >4 months | Open, P1 | Needs security review and product decision. |
| [#7707](https://github.com/openclaw/openclaw/issues/7707) – Memory trust tagging | 2026-02-03 | >4 months | Open, P2 | Needs security review and product decision. |
| [#6731](https://github.com/openclaw/openclaw/issues/6731) – Safe/unsafe ClawdBot | 2026-02-02 | >4 months | Open, P1 | Needs security review and product decision. |
| [#13583](https://github.com/openclaw/openclaw/issues/13583) – Pre‑response hooks | 2026-02-10 | >4 months | Open, P1 | Needs security review. |
| [#13751](https://github.com/openclaw/openclaw/issues/13751) – Feishu permission reduction | 2026-02-11 | >4 months | Open, P1 | Needs security review and product decision. |
| [#14785](https://github.com/openclaw/openclaw/issues/14785) – Reduce tool schema token overhead | 2026-02-12 | >4 months | Open, P2 | Needs maintainer review. |
| [#40255](https://github.com/openclaw/openclaw/issues/40255) – Heartbeat prompt regression | 2026-03-08 | ~3 months | Open, P1 | Linked PR open (#? ), but still needs product decision. |
| [#86215](https://github.com/openclaw/openclaw/issues/86215) – OAuth refresh failures | 2026-05-24 | ~3 weeks | Open, P1 | Needs maintainer review, live repro. |

**PRs needing maintainer attention:**
- [#92725](https://github.com/openclaw/openclaw/pull/92725) – External reranker (XL, 13 days old, waiting on author).
- [#93956](https://github.com/openclaw/openclaw/pull/93956) – Ollama skip auto‑discovery for remote URLs (M, waiting on author).
- [#94054](https://github.com/openclaw/openclaw/pull/94054) – Reject empty cert paths (S, waiting on author).
- [#75403](https://github.com/openclaw/openclaw/pull/75403) – Typing indicator regression fix (S, waiting on author).

**Recommendation:** The maintainers should prioritise the four oldest P1 security‑related issues (#25592, #10659, #6731, #13583) as they have been open for months with no maintainer decision. The OAuth failure bug (#86215) is newer but production‑critical.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: Personal AI Agent Open-Source Ecosystem

## 1. Ecosystem Overview

The personal AI agent open-source landscape is experiencing a phase of intense feature development coupled with a maturing focus on production reliability. Projects are converging on essential architectural patterns—multi-channel gateways, plugin systems, memory management, and tool-use orchestration—while diverging on target deployment profiles (lightweight embedded vs. full-stack desktop). Security hardening has emerged as a top cross-project priority, with multiple critical vulnerabilities reported simultaneously, signaling that the ecosystem is moving from rapid prototyping to operational trustworthiness. Community engagement remains healthily distributed, with both core teams and external contributors driving meaningful improvements across almost all projects.

## 2. Activity Comparison

| Project | Issues (24h) | PRs (24h) | Merged/Closed PRs | Release (24h) | Health Score* |
|---------|-------------|-----------|-------------------|----------------|---------------|
| **OpenClaw** | 252 | 500 | 91 merged/closed | None | Medium |
| **NanoBot** | 9 | 31 | 17 merged/closed | None | High |
| **Hermes Agent** | 10 | 50 | 2 merged/closed | None | Medium |
| **PicoClaw** | 4 | 10 | 6 merged/closed | None | High |
| **NanoClaw** | 5 | 20 | 3 merged/closed | **v2.1.17** | High |
| **NullClaw** | 3 | 2 | 0 merged/closed | None | Low |
| **IronClaw** | 11 | 50 | 17 merged/closed | None | High |
| **LobsterAI** | 1 | 10 | 10 merged/closed | **v2026.6.15** | High |
| **TinyClaw** | 3 | 0 | 0 merged/closed | None | Critical |
| **Moltis** | 2 | 1 | 0 merged/closed | None | Medium |
| **CoPaw** | 45 | 50 | 33 merged/closed | **v1.1.12** | High |
| **ZeptoClaw** | 0 | 0 | 0 | None | Dormant |
| **ZeroClaw** | 17 | 50 | 11 merged/closed | None | High |

*Health Score: High = active triage, critical bugs fixed same day; Medium = steady progress with some unresolved blockers; Low = minimal activity with stale issues; Critical = security vulnerabilities unaddressed*

## 3. OpenClaw's Position

| Dimension | OpenClaw | Peer Average (active projects) |
|-----------|----------|-------------------------------|
| Daily issue volume | 252 | 12 |
| Daily PR volume | 500 | 33 |
| Open P1/P2 bugs | ~40 | ~5 |
| Community contributors (24h) | 5+ external | 2-3 external |
| Release cadence | Irregular | Monthly average |
| Backlog age (oldest critical) | 4 months | 2 weeks - 1 month |

**Advantages:**
- **Largest community mass** by orders of magnitude—5x more PRs than next most active project (Hermes/ZeroClaw). This drives faster bug discovery and wider integration support.
- **Most comprehensive multi-channel support** (Telegram, Slack, Feishu, Discord, Matrix, iMessage) with dedicated fixes in active flight.
- **Reference implementation status**—other projects (LobsterAI, CoPaw) explicitly build compatibility layers for OpenClaw's compaction, skill formats, and gateway protocols.

**Technical Approach Differences:**
- **Heavier, more opinionated architecture** compared to NanoBot/PicoClaw—uses a dedicated Gateway process, skill-sync daemon, and proprietary session compaction.
- **Tool-discovery system** that other projects (ZeroClaw, IronClaw) are re-implementing with WASM plugins or explicit registries to avoid OpenClaw's provider-wipe bug.
- **"Platinum Hermit" severity tier**—unique to OpenClaw, indicating a more granular (some might say complex) triage system.

**Key Weakness:** Backlog bloat. Critical P1 issues like message text leakage (#25592) have been open for 4 months, and OAuth refresh failures (#86215) for 3 weeks without maintainer decision. This contrasts with NanoClaw, which fixed a critical session stall bug within hours of report.

## 4. Shared Technical Focus Areas

| Focus Area | Projects Affected | Specific Needs |
|------------|------------------|----------------|
| **Security by Default** | OpenClaw, TinyClaw, LobsterAI, NanoClaw, Hermes Agent | Arbitrary file reads (TinyClaw #282, LobsterAI #2176), secret masking (OpenClaw #10659), memory injection bypass (Hermes Agent #48181), path traversal (NanoClaw #2800) |
| **Message Delivery Reliability** | OpenClaw, NanoClaw, IronClaw, CoPaw | Session stall kills all delivery (NanoClaw #2796), text between tool calls leaks to channels (OpenClaw #25592), subagent approval routing fails (CoPaw #5295) |
| **Multi-Platform Gateway** | OpenClaw, PicoClaw, ZeroClaw, CoPaw | Feishu WebSocket cards (PicoClaw), Matrix infinite loop protection (CoPaw #5204), Mattermost WebSocket listener (ZeroClaw #7098) |
| **Cost / Budget Governance** | OpenClaw, ZeroClaw, IronClaw, NanoBot | Per-agent cost budgets (OpenClaw #42475), cached input token pricing (ZeroClaw #7492), context window limits per model (NanoBot #4389) |
| **Context & Memory Management** | OpenClaw, NanoBot, IronClaw, CoPaw, Hermes Agent | Context compression decorator (ZeroClaw #7673), memory trust tagging by source (OpenClaw #7707), compaction crashes (CoPaw #5287), replay window preservation (NanoBot #4349) |
| **Agent-to-Agent Security** | CoPaw, NanoClaw, Hermes Agent | Approval delegation (Hermes #47863), per-message approval gates (NanoClaw #2793), subagent approval routing to external channels (CoPaw #5295) |
| **Onboarding & Documentation** | NullClaw, NanoBot, OpenClaw, CoPaw | Headless VPS guides (NullClaw #861), interactive setup wizard (NanoBot #4376), Android APK releases (OpenClaw #9443) |

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | Hermes Agent | PicoClaw | NanoClaw | IronClaw | CoPaw | ZeroClaw |
|-----------|----------|---------|--------------|----------|----------|----------|-------|----------|
| **Primary Target User** | Power users, self-hosters | Lightweight end-users | Desktop-first enthusiasts | Embedded/RISC-V hackers | Production operators | Enterprise teams | Chinese-market users | Advanced CLI/API devs |
| **Deployment Model** | Full-stack gateway + daemon | Minimal binary | Desktop app + remote runtime | Tiny single-binary | Container/multi-instance | Local + cloud hybrid | Docker-heavy | CLI-first + WASM plugins |
| **Signature Feature** | Session compaction | Git workspace security | OTLP observability | Gemini 3.5 Flash fix | Per-message approval | Projects entity & skill evolution | XiaoYi channel support | Alias rename cascade |
| **Language Stack** | TypeScript/Python | TypeScript | TypeScript | Rust | Rust + TypeScript | Rust | Python/TypeScript | Rust |
| **Plugin Architecture** | Skill sync daemon | MCP plugins | Plugin hooks + LUMEN | Built-in only | CLI-derived skills | WASM plugins | AgentScope framework | WASM plugin lifecycle |
| **Community Language** | English + Chinese | English | English | Chinese-heavy | Mixed | Chinese-heavy | Chinese-dominant | English |
| **Release Maturity** | Pre-stable, high churn | Stable, incremental | Pre-stable, bursty | Early, conservative | Production-ready v2.x | Pre-stable v0.29 | Stable v1.1.12 | Pre-stable v0.8 |
| **Unique Risk** | Backlog starvation | Windows Safari UI bug | Desktop thin-client gap | Dependency on libolm | Tripwire breaking changes | 74 Windows test failures | Docker upgrade regression | Android compile failure |

## 6. Community Momentum & Maturity

### Tier 1: High Momentum (rapid iteration, strong maintainer response)

| Project | Velocity Signal | Release Cadence | Community Growth Indicator |
|---------|----------------|-----------------|---------------------------|
| **OpenClaw** | 91 PRs merged in 24h | Irregular (none today) | 5+ external contributors, but 4-month-old critical bugs |
| **NanoClaw** | 3 PRs merged + critical bug fixed same day | v2.1.17 today | Korean README contributor, fast triage |
| **IronClaw** | 17 PRs merged, 4 issues closed | Automated release pipeline | Dogfooding tracker, new contributor (PR #5061) |
| **CoPaw** | 33 PRs merged, 2 releases | v1.1.12 stable + beta same day | Multiple first-time contributors, active Chinese community |
| **ZeroClaw** | 11 PRs merged, stacked PR series completed | v0.8.x milestone | Active RFC discussions, contributor @NiuBlibing fixing Windows |

### Tier 2: Steady Growth (consistent progress, manageable backlog)

| Project | Velocity Signal | Risks | 
|---------|----------------|-------|
| **NanoBot** | 17 PRs merged, quick fixes for workspace security | Light community discussion (low comments) |
| **Hermes Agent** | 50 PRs open, OTLP plugin merged | 2 PRs merged only; thin client request (#38602) unaddressed |
| **PicoClaw** | 6 PRs merged, security fix (#3140) | 2 open issues without maintainer response |
| **LobsterAI** | 10 PRs merged, computer use feature released | Security vulnerability (#2176) reported today, no fix yet |

### Tier 3: Low Activity / Stalled

| Project | Signal | Recommendation |
|---------|--------|----------------|
| **NullClaw** | 0 PRs merged, 3 issues open 30+ days without maintainer reply | Maintainer needs to triage scheduler bug (#915) |
| **TinyClaw** | 0 PRs, 3 critical security vulnerabilities filed today | Immediate maintainer response required |
| **Moltis** | 1 open PR, 2 enhancement requests, no maintainer comment | Acknowledge TTS config request (#1126) |
| **ZeptoClaw** | No activity at all | May be abandoned |

## 7. Trend Signals for AI Agent Developers

### Security Hardening as Prerequisite for Production
- **5 of 12 projects** reported or fixed critical security issues today—file read vulnerabilities (TinyClaw, LobsterAI, NanoClaw), credential leaks (OpenClaw, Hermes Agent), and authentication bypass (TinyClaw). **Implication:** Any new agent framework must ship with sandboxed file access, secret masking, and API authentication by default. The era of "trust the agent" is over.

### The "Reliability Wall" at 10+ Agents
- Multiple projects report system-level failures under concurrency: session stalls (NanoClaw), simultaneous timeouts (OpenClaw), OOM crashes (LobsterAI), and approval attribution races (ZeroClaw). **Implication:** The ecosystem is hitting scaling limits of single-process architectures. Developers building multi-agent deployments should expect to implement circuit breakers, per-session isolation, and external monitoring (OTLP in Hermes Agent points this way).

### Multi-Channel Gateway is Table Stakes
- Every active project supports at least 3 messaging channels. ZeroClaw, OpenClaw, and CoPaw are adding WebSocket listeners, native polls, and platform-specific UX (iMessage polls in Hermes, Feishu QR scanning in NanoBot). **Implication:** Agent-to-user interaction is converging on a "universal inbox" pattern—projects that don't abstract channel handling (TinyClaw, NullClaw) are falling behind.

### Context Management Remains Unsolved
- Context compression (ZeroClaw RFC #7673), memory trust tagging (OpenClaw #7707), compaction crashes (CoPaw #5287), and replay-window history (NanoBot #4349) are all active pain points. **Implication:** There is no dominant solution for bounded context windows with reliable memory recall. This is the single highest-ROI area for a new tool or library.

### Governance Features Signal Enterprise Demand
- Approval delegation (Hermes Agent), per-message gates (NanoClaw), per-agent cost budgets (OpenClaw), and security policy reload without restart (ZeroClaw) are being built concurrently across projects. **Implication:** The "personal assistant" label is increasingly serving team/enterprise deployments. Expect SSO, audit logs, and role-based skill access to become requirements within 6 months.

### WASM and Plugin Systems are Converging
- ZeroClaw is implementing WASM plugin lifecycle hooks; IronClaw has a WASM-based skill extraction PR; CoPaw supports MCP skills. OpenClaw's skill-sync daemon and Hermes Agent's LUMEN binary protocol represent alternative approaches. **Implication:** The plugin interface is not standardized—developers face a binding choice. WASM offers portability but adds complexity; OpenClaw's approach favors speed but risks lock-in.

### Mobile Support is a Growing Gap
- Android APK requests (OpenClaw #9443), iOS Safari zoom bugs (NanoBot #4388), Termux compile failures (ZeroClaw #7911), and thin-client requests (Hermes #38602) all point to unmet mobile demand. **Implication:** No project delivers a first-class mobile experience. This is a strategic opportunity—the first agent project with a polished, cross-platform mobile client will capture significant share from power users who want "agent in my pocket."

---

*Report generated from GitHub data snapshot 2026-06-18. Reflects activity across 13 projects in the open-source personal AI assistant ecosystem.*

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest – 2026-06-18

## 1. Today’s Overview
NanoBot saw a day of high activity, with **31 pull requests updated** in the last 24 hours—**17 merged or closed**—and **9 issues touched**, of which 7 remain open. The project continues to smooth out rough edges from recent workspace security (#4202, #4380) and channel-specific bugs (Feishu, WhatsApp, Anthropic). A surge of merged fixes indicates strong maintainer responsiveness, though no new release was published. Community discussion remains quiet in terms of comments, but several feature requests signal a growing demand for multi-instance support and better onboarding.

## 2. Releases
**No new releases** in the last 24 hours. The latest published version remains unchanged.

## 3. Project Progress
The following notable PRs were merged or closed today (all closed as valid/bug fixes or feature additions unless noted):

| PR | Title | Key Outcome |
|----|-------|-------------|
| #4380 | fix: allow git commands in workspace subdirectories | Resolves a workspace security guard that blocked `git add/commit/push` from subdirectories (#4375). |
| #4381 | fix: recover failed Feishu streaming updates | Retries Feishu CardKit content updates after streaming failures. |
| #4385 | fix: log primary model error before fallback | Includes primary model error text in logs before attempting fallback models. |
| #4350 | feat(web): add Keenable search provider | New built-in web search provider (Keenable). |
| #4351 | feat(providers): better Mistral support | Fixes Mistral API strictness around `reasoning_effort`, `tools/functions`, `stop`, `min_p`. |
| #4354 | feat(bridge): send read receipts (blue ticks) for incoming WhatsApp messages | Marks messages as read via WhatsApp bridge. |
| #4202 | Clarify filesystem workspace write policy | Aligns `apply_patch` and tool paths with explicit read/write allowed directories. |
| #4349 | fix(session): preserve user turns in replay-window history | Prevents LLM replay from starting mid-turn by keeping full user turns. |
| #4283 | fix(webui): correct activity duration display | Uses final assistant latency, labels mixed tool activity as “work”. |
| #4347 | fix(WebUI): fix `my tool` model preset switching | Handles model preset selection in the My Tool skill. |
| #4356 | fix(anthropic): sanitize tool_use/tool_result IDs to API pattern | Strips invalid characters from tool IDs to avoid 400 errors. |
| #4367 | fix(providers): disable proxy for local endpoints, respect env proxy for cloud | Local model servers (Ollama, etc.) no longer forced through HTTP proxy. |
| #4053 | fix(tools): keep read-only roots out of write paths | Ensures extra allowed dirs are read-only; stops write tools inheriting media-dir access. |

Other open PRs advancing features include #4342 (Feishu WebSocket card content), #4391 (Feishu QR scan-to-create bot), #4394 (OpenAI image reference edits), #4303 (MCP generator cleanup), #4373 (memory delivery preservation), and #4392 (tool microcompaction configurable).

## 4. Community Hot Topics
Issues and PRs with the most discussion or reactions (comments count is low across the board; we highlight those with at least one comment or a reaction):

- **#4374** – [feature request] Project workspaces: SOUL.md/USER.md read/write asymmetry  
  *1 comment, 0 reactions*  
  Highlights a design flaw: bootstrap files are read from the project per turn but written to the default workspace. Community expects symmetry.  
  [GitHub Issue #4374](https://github.com/HKUDS/nanobot/issues/4374)

- **#4389** – [feature request] Per-model `contextWindowTokens` for fallback models  
  *1 comment, 0 reactions*  
  When fallback models have smaller context windows, the prompt is not trimmed. Users want per-model token limits.  
  [GitHub Issue #4389](https://github.com/HKUDS/nanobot/issues/4389)

- **#4376** – [enhancement] User-friendly wizard  
  *1 comment, 1 reaction (👍)*  
  New/non-technical users find `nanobot onboard --wizard` too technical. Calls for a guided, interactive onboarding UX.  
  [GitHub Issue #4376](https://github.com/HKUDS/nanobot/issues/4376)

- **#4390** – [feature request] Multi-instances for normies  
  *0 comments, 0 reactions*  
  A user wants to run multiple NanoBot instances on one machine, organized per folder with per-folder config, and hide UI settings. No maintainer response yet.  
  [GitHub Issue #4390](https://github.com/HKUDS/nanobot/issues/4390)

- **#936** – [feature request] Multi-Tenant Gateway for Multiple Agents (opened Feb 2026)  
  *1 comment, 0 reactions*  
  Long-standing request for a single gateway to manage multiple agents. Still open with no recent activity.  
  [GitHub Issue #936](https://github.com/HKUDS/nanobot/issues/936)

## 5. Bugs & Stability
Ranked by severity (critical → low):

| Severity | Issue | Status | Fix PR |
|----------|-------|--------|--------|
| **Critical** | **#4375** – Git commands blocked by workspace security guard (subdirectories) | Closed | Merged #4380 |
| **High** | **#4388** – iOS Safari: tapping input triggers page zoom, UI distortion | Open | None yet |
| **High** | **#4303** – MCP `streamableHttp` crash: `RuntimeError: cancel scope in different task` | Open PR #4303 (ready to merge) | Not yet merged |
| **High** | **#4373** – Memory consolidation loses delivery context (replay window) | Open PR #4373 | Not yet merged |
| **Medium** | **#4356** – Anthropic tool ID sanitisation → 400 errors | Closed | Merged #4356 |
| **Medium** | **#4367** – Proxy blocking local model servers | Closed | Merged #4367 |
| **Medium** | **#4322** – `NameError: 'session_key'` after merge (stale) | Closed | Merged earlier |
| **Medium** | **#4381** – Feishu streaming card update failures | Closed | Merged #4381 |
| **Low** | **#4385** – Primary model error not logged before fallback | Closed | Merged #4385 |

Two open bugs (iOS zoom, MCP crash) lack a merged fix and could affect mobile and MCP-heavy users respectively.

## 6. Feature Requests & Roadmap Signals
Several open issues point to clear user demands that could shape the next minor/major version:

- **Workspace symmetry** (#4374) – Likely next milestone after workspace security rework; consistency between read and write paths.
- **Per-model context window** (#4389) – Essential for reliable fallback behaviour; could be combined with the existing fallback refactoring PRs.
- **User-friendly wizard** (#4376) – High community engagement (👍); a classic “first experience” improvement that lowers barriers.
- **Multi-instance / multi-tenant** (#4390, #936) – Both requests target the same pain: running multiple agents efficiently. #936 is older and bigger; #4390 is simpler. The maintainers may choose to implement a lightweight folder-based approach first.
- **On-demand heartbeat** (#3437) – Debugging feature for `HEARTBEAT.md` iteration; open since April, no recent activity.

**Prediction for next release:** The workspace asymmetry fix (#4374) and per-model context window (#4389) seem plausible as small, self-contained improvements. The onboarding wizard (#4376) or multi-instance support (#4390) might appear as experimental features.

## 7. User Feedback Summary
- **Pain points**  
  - Workspace security is too restrictive (git commands blocked in subdirs) – fixed today.  
  - Feishu cards display incorrectly via WebSocket – fixed today.  
  - iOS Safari zoom breaks the WebUI – open, mobile users affected.  
  - Onboarding requires too much technical knowledge (#4376).  
  - Fallback models cannot use smaller context windows (#4389).  

- **Use cases**  
  - Single machine, multi-instance per folder (#4390).  
  - Multi-agent gateway for production deployments (#936).  
  - Feishu bot creation without manual app configuration (#4391).  

- **Satisfaction signals**  
  - Quick turnaround on git security fix (#4380) – shows maintainer attentiveness.  
  - Integration of new search provider (Keenable) without community push.  
  - Read receipts (WhatsApp) and Mistral support added promptly.  

Overall, the community seems engaged but light on discussion; the steady stream of merged fixes suggests a healthy project with responsive maintainers.

## 8. Backlog Watch
Issues that have been open for weeks or months without a maintainer response or significant progress:

| Issue | Age | Topic | Action Needed |
|-------|-----|-------|---------------|
| **#936** – Multi-Tenant Gateway | ~4 months | Infrastructure | Needs spec or triage; could be merged with #4390. |
| **#3437** – On-demand heartbeat | ~2 months | Debugging | Low complexity; could be a good first PR for contributors. |
| **#4376** – User-friendly wizard | 1 day | Onboarding | Already has one reaction; maintainers should acknowledge and invite design. |
| **#4390** – Multi-instance for normies | 1 day | Configuration | New but duplicates #936’s spirit; recommend merging into one meta-issue. |

**Maintainer attention requested:** #936 and #3437 are languishing. A brief comment with “looking into it” or “help wanted” would signal project health.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest – 2026-06-18

## 1. Today's Overview
Activity remains high, with **10 issues** and **50 pull requests** updated in the last 24 hours. All issues remain open; 2 PRs were merged/closed today (details not listed), while 48 open PRs continue to advance. The project is in a burst of feature work and critical bug fixes, particularly around OAuth billing compatibility for Anthropic clients and security hardening of memory tool injection. No new releases were cut. The community is actively shaping the desktop client, gateway UX, and plugin ecosystem.

## 2. Releases
**No new releases** today. The last published release is not specified in this dataset.

## 3. Project Progress
Two PRs were **merged or closed** today, though their specific subjects are not shown in the top-20 list. The open PR landscape indicates significant forward momentum:

- **Anthropic OAuth billing fix** ([PR #48177](https://github.com/NousResearch/hermes-agent/pull/48177)) – adds the required `x-anthropic-billing-header` system block to enable Pro/Max/Team OAuth requests. A companion PR ([#48192](https://github.com/NousResearch/hermes-agent/pull/48192)) normalises tool names on the wire.
- **Memory security bypass** ([Issue #48181](https://github.com/NousResearch/hermes-agent/issues/48181)) – identified and reported; fix PR not yet linked.
- **Desktop clipboard on WSL** ([PR #48186](https://github.com/NousResearch/hermes-agent/pull/48186)) – resolves “No image found” errors on WSL2.
- **Observability plugin (OTLP)** ([PR #48184](https://github.com/NousResearch/hermes-agent/pull/48184)) – adds OpenTelemetry export capability.
- **LUMEN binary protocol for MCP** ([PR #47740](https://github.com/NousResearch/hermes-agent/pull/47740)) – brings wire compression and multi-agent streaming.
- **ComfyUI skill fix** ([PR #48145](https://github.com/NousResearch/hermes-agent/pull/48145)) – strips `_comment` keys that broke workflow submission.
- **Four new optional skills** ([PR #47576](https://github.com/NousResearch/hermes-agent/pull/47576)) – graphify, ui-ux-pro-max, impl-validator, suede-promo.

## 4. Community Hot Topics
### Most Active Issues
| Issue | Comments | Reactions | Summary |
|-------|----------|-----------|---------|
| [#38602 – Desktop Client-Only Installation](https://github.com/NousResearch/hermes-agent/issues/38602) | 6 | 18 👍 | Request to install Hermes Desktop as a thin client (no local agent runtime). |
| [#41808 – Dashboard Chat React error #301](https://github.com/NousResearch/hermes-agent/issues/41808) | 2 | 0 | Maximum update depth exceeded on external (Tailscale) connections. |
| [#48175 – Discord system prompts from external files](https://github.com/NousResearch/hermes-agent/issues/48175) | 2 | 0 | How to configure separate AI personas per Discord channel. |

The **thin client feature** (#38602) draws the strongest community support, indicating a desire for remote agent management. The **React crash** (#41808) affects external network users and is a stability concern.

### Most Active PRs (by potential impact)
- [#47863 – Native approval delegation](https://github.com/NousResearch/hermes-agent/pull/47863) – cross-platform (WeChat/WeCom → Feishu) admin approval for dangerous commands.
- [#48194 – iMessage native polls for clarify](https://github.com/NousResearch/hermes-agent/pull/48194) – renders multi-choice questions as native iMessage polls.
- [#48177 – Anthropic OAuth billing fix](https://github.com/NousResearch/hermes-agent/pull/48177) – directly addresses a P1 bug blocking Pro/Max/Team users.

## 5. Bugs & Stability
### Critical (P1)
- **OAuth billing header missing** ([#48176](https://github.com/NousResearch/hermes-agent/issues/48176)) – Claude Pro/Max/Team OAuth requests return HTTP 400. **Fix PR exists** ([#48177](https://github.com/NousResearch/hermes-agent/pull/48177)).

### High (P2)
- **Memory toolset bypass** ([#48181](https://github.com/NousResearch/hermes-agent/issues/48181)) – disabled `memory` toolsets can be re-enabled via late injection. **No fix PR yet** – security advisory.
- **Desktop profile/session mismatch** ([#48183](https://github.com/NousResearch/hermes-agent/issues/48183)) – after switching profile, session list shows empty; data exists in correct state.db but UI fails to load. No fix PR yet.
- **Dashboard React error on Tailscale** ([#41808](https://github.com/NousResearch/hermes-agent/issues/41808)) – only impacts external connections; localhost works. No fix PR in data.

### Medium (P3)
- **Desktop update fails in Docker-backed setups** ([PR #48188](https://github.com/NousResearch/hermes-agent/pull/48188)) – install-method stamp scoping fix.
- **WSL clipboard paste always fails** ([PR #48186](https://github.com/NousResearch/hermes-agent/pull/48186)) – fix in progress.
- **Skill ComfyUI workflows broken** ([PR #48145](https://github.com/NousResearch/hermes-agent/pull/48145)) – fix ready.
- **Browser tools hidden on macOS** ([PR #48185](https://github.com/NousResearch/hermes-agent/pull/48185)) – detection of app-bundle Chrome.

### Other
- `model.context_length` global bleed ([PR #48187](https://github.com/NousResearch/hermes-agent/pull/48187)) – session-scoped fix.
- Hindsight API port dropped ([PR #48191](https://github.com/NousResearch/hermes-agent/pull/48191)) – env materialisation fix.

## 6. Feature Requests & Roadmap Signals
### Strongest Community Wants
- **Desktop thin client** (#38602) – would enable remote agent control without local runtime. Likely to land in next minor release given high reaction count.
- **Session ↔ workspace binding** ([#48190](https://github.com/NousResearch/hermes-agent/issues/48190)) – record `cwd`, repo, and branch per session. Roadmap signal for TUI/CLI productivity.
- **Desktop Gateway Command Center** ([#48189](https://github.com/NousResearch/hermes-agent/issues/48189)) – start/stop/install messaging gateway from desktop UI.
- **Better interactive experience** ([#48182](https://github.com/NousResearch/hermes-agent/issues/48182)) – skill highlight, inline preview, UI polish.
- **Managed system generalisation** ([#48179](https://github.com/NousResearch/hermes-agent/issues/48179)) – extend beyond NixOS/Homebrew to Fedora `dnf`.

### PRs That Signal Roadmap Direction
- **Approval delegation** ([#47863](https://github.com/NousResearch/hermes-agent/pull/47863)) – enterprise/team governance.
- **OTLP observability** ([#48184](https://github.com/NousResearch/hermes-agent/pull/48184)) – production monitoring.
- **LUMEN binary MCP** ([#47740](https://github.com/NousResearch/hermes-agent/pull/47740)) – performance and multi-agent support.
- **iMessage native polls** ([#48194](https://github.com/NousResearch/hermes-agent/pull/48194)) – platform-specific UX improvement.
- **System tray support** ([#48163](https://github.com/NousResearch/hermes-agent/pull/48163)) – desktop UX QoL.

**Prediction for next version:** At least two of the P1/P2 bug fixes (OAuth billing, memory bypass) will be merged, alongside the thin-client feature (#38602) and the approval delegation PR (#47863). The OTLP plugin (#48184) may also land as a new optional module.

## 7. User Feedback Summary
### Pain Points Expressed
- **Cannot use Claude Pro/Max/Team OAuth** – users report blocked usage. Workaround unclear.
- **Desktop is too heavy for remote setups** – users want a thin client that connects to a server-side agent.
- **Session management is flat** – lack of workspace context (cwd, repo) makes resuming work difficult.
- **Discord multi-persona configuration cumbersome** – users want clean external prompt files per channel.
- **React crash on Tailscale** – external network users hit a UI-breaking bug.
- **Desktop profile switching breaks session list** – data integrity but UI failure.
- **Memory tool bypass is a security concern** – administrators need assurance that disabled tools stay disabled.

### Satisfaction Indicators
- Positive response to **new optional skills** (graphify, ux, etc.) and **approval delegation**.
- Community engagement is high – many first-time contributors submitting PRs (e.g., #48194, #48186, #48185).
- **OTLP plugin**, **LUMEN protocol**, and **native iMessage polls** show the project is investing in both power-user and platform-native experiences.

## 8. Backlog Watch
### Long-Open Issues Needing Maintainer Attention
- [#38602 – Desktop thin client](https://github.com/NousResearch/hermes-agent/issues/38602) – created June 4, high reaction count, no maintainer response yet. Should be prioritised.
- [#41808 – Dashboard React error on Tailscale](https://github.com/NousResearch/hermes-agent/issues/41808) – June 8, 2 comments, no fix PR. Impacting external network users.

### Old PRs Awaiting Review/Merge
- [#27208 – Agent loop stopped plugin hook](https://github.com/NousResearch/hermes-agent/pull/27208) – opened May 17, 30 days stale. Adds useful interrupt lifecycle hook.
- [#19331 – Cognee query tool](https://github.com/NousResearch/hermes-agent/pull/19331) – opened May 3, over 6 weeks old. Adds read-only knowledge-graph tool.
- [#45929 – Clear npm advisories + scope tags](https://github.com/NousResearch/hermes-agent/pull/45929) – June 14, addresses security vulnerabilities and global state pollution.

### Security Advisory
- [#48181 – Memory tool bypass](https://github.com/NousResearch/hermes-agent/issues/48181) – filed today, P2, no fix PR. Should be escalated to maintainers for prompt remediation.

---

*Generated from GitHub data snapshot 2026-06-18. All links point to NousResearch/hermes-agent.*

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest — 2026-06-18

## Today’s Overview
Project activity was moderate on June 18, with 4 issues and 10 pull requests updated in the last 24 hours. The team closed 6 PRs (including one security fix) and merged a new LLM provider integration, while 2 open issues remain active. Focus appears to be on bug fixing, security hardening, and expanding gateway compatibility. No new releases were published.

## Releases
No new releases. The project currently has no tagged version in the reported window.

## Project Progress
**Merged/Closed PRs (last 24h):**  
- **#3140** — Fixed OneBot inbound media URL handling to block private network fetches (security fix).  
- **#3139** — Updated Sogou search regex to match changed HTML structure, restoring web search functionality.  
- **#3136** — Added both camelCase and snake_case `thought_signature` in Gemini tool call requests, resolving a 400 error with Gemini 3.5 Flash.  
- **#2990** — Fixed Web UI session history display to show all user messages instead of only the last one.  
- **#2917** — Merged new [NEAR AI Cloud](https://github.com/sipeed/picoclaw/pull/2917) provider as a first-class OpenAI-compatible LLM backend.  
- **#3138** — Added a review feature (Korean-language title; details unclear from summary).

**Open PRs with recent updates:**  
- **#3142** — Fixes duplicate message delivery in sub-agent `ToolResult` (preventing double push).  
- **#3141** — Adds diagnostic logging for Brave Search API empty results.  
- **#3092** — Improves type assertion safety in skill installation (still open after 8 days).  
- **#3063** — Implements [DeltaChat gateway](https://github.com/sipeed/picoclaw/pull/3063) (feature, open since June 8).

## Community Hot Topics
- **#3088** — *[Feature] Use vodozemac instead of libolm*  
  High priority issue with 2 👍 and 1 comment. The community wants to replace the unmaintained and insecure `libolm` with the official `vodozemac` library. This is the most upvoted open issue and signals strong user demand for cryptographic upgrade.  
  [Issue #3088](https://github.com/sipeed/picoclaw/issues/3088)

- **#3093** — *[Feature] I need SimpleX or Tox*  
  Request for adding alternative communication gateways (SimpleX, Wire, or Tox). No maintainer response yet.  
  [Issue #3093](https://github.com/sipeed/picoclaw/issues/3093)

- **#3142** — *fix(spawn): clear ForUser in sub-turn ToolResult to prevent duplicate messages*  
  Active PR with recent update; addresses a root cause that caused message duplication in async sub-agent workflows — a likely user-facing annoyance.  
  [PR #3142](https://github.com/sipeed/picoclaw/pull/3142)

## Bugs & Stability
| Severity | Issue/PR | Description | Status |
|----------|----------|-------------|--------|
| **High** | #3070 (closed) | OneBot media URL handling allowed host-side arbitrary fetch (SSRF-like vulnerability) | Fixed in PR #3140 |
| **Medium** | #3111 (closed) | Tool execution fails with Gemini 3.5 Flash due to missing `thought_signature` in schema | Fixed in PR #3136 |
| **Low** | #3141 (open PR) | Brave Search API returning empty results silently – diagnostic logging added | PR open |
| **Low** | #3092 (open PR) | Skill installation silently ignores wrong types for `version`/`force` arguments | PR open (stale) |

The **security vulnerability** (#3070) has been patched. The **Gemini 3.5 Flash** compatibility issue was resolved. A **Sogou search parsing regression** was also fixed (#3139).

## Feature Requests & Roadmap Signals
- **Vodozemac replacement (libolm → vodozemac)** — #3088 is high priority, labelled `help wanted`. Likely candidate for next release given security implications and official replacement status.
- **DeltaChat gateway** — PR #3063 is open and not merged; may appear in a future release if maintainers accept it.
- **NEAR AI Cloud provider** — Already merged (PR #2917), so it will be available in the next release.
- **SimpleX / Tox gateway** — #3093 has no maintainer interaction; less likely for immediate roadmap.

## User Feedback Summary
- **Pain points**: Users report frustration with Gemini 3.5 Flash tool execution failure (#3111) and silent failures in web search (#3139, #3141). The `libolm` deprecation (no longer maintained) is a security concern that users actively want addressed (#3088).
- **Positive signals**: The community appreciates the quick fix for the SSRF vulnerability (#3070 → #3140). The new NEAR AI Cloud provider was well-received (merged after a month of development).
- **Missing maintainer response**: Issue #3093 (SimpleX/Tox) and PR #3092 (skill type assertions) have been open for 8–10 days without recent maintainer attention.

## Backlog Watch
- **Issue #3088** — High-priority feature request (vodozemac). No maintainer comment since creation on June 9. The `help wanted` label suggests external contribution is encouraged.
- **PR #3063** — DeltaChat gateway feature PR open since June 8, with no new commits or reviewer activity since June 17. Needs maintainer review to avoid stagnation.
- **Issue #3093** — Gateway request (SimpleX/Tox) uncommented by maintainers. May need prioritization if the project aims to expand communication channels.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-18

## 1. Today's Overview

NanoClaw saw a **very active day** with 20 pull requests updated (3 merged/closed) and 5 issues updated (1 closed). Two rollup releases landed today—v2.1.0 and v2.1.17—both carrying breaking changes. The community is highly engaged, contributing documentation translations, security fixes, and new features. Overall project health is strong, with rapid responses to critical bugs (a delivery stall issue was opened and fixed within hours).

## 2. Releases

**v2.1.17** ([release notes](https://github.com/nanocoai/nanoclaw/releases/tag/v2.1.17))  
Rollup covering `v2.1.1` through `v2.1.17`.

**Breaking changes:**
- [`@onecli-sh/sdk` 0.5.0 → 2.2.1] (https://github.com/nanocoai/nanoclaw/issues?q=is%3Aissue+%40onecli-sh%2Fsdk) — requires a OneCLI server with the `/v1` API. Older servers will return 404 on every SDK call. The sanctioned gateway and CLI versions are now pinned.
- **Startup now requires an upgrade marker** ([v2.1.0](https://github.com/nanocoai/nanoclaw/releases/tag/v2.1.0) also carries this): the host refuses to boot unless `data/upgrade-state.json` records that the install reached the current version through a sanctioned upgrade path.

**Migration notes:**
- Update your OneCLI server to a version that supports the `/v1` API before upgrading the NanoClaw SDK.
- If you run managed fleets or immutable images, set the environment variable `NANOCLAW_DISABLE_UPGRADE_TRIPWIRE=1` to bypass the startup tripwire (see PR [#2780](https://github.com/nanocoai/nanoclaw/pull/2780)).

## 3. Project Progress

**Merged/closed PRs today:**

| PR | Title | Status |
|----|-------|--------|
| [#2797](https://github.com/nanocoai/nanoclaw/pull/2797) | fix(delivery): isolate per-session failures | **Merged** – Fixes critical stall bug |
| [#2794](https://github.com/nanocoai/nanoclaw/pull/2794) | fix(providers): restore env-var gateway auth for managed fleets | **Merged** – Fixes authentication regression |
| [#2780](https://github.com/nanocoai/nanoclaw/pull/2780) | feat(upgrade-state): env opt-out for the startup tripwire | **Merged** – Enables managed-fleet deployments |

**Also closed:** Issue [#2796](https://github.com/nanocoai/nanoclaw/issues/2796) (session stall) was resolved by PR #2797.

## 4. Community Hot Topics

- **Session delivery stall** – Issue [#2796](https://github.com/nanocoai/nanoclaw/issues/2796) (1 comment) described a critical bug where a single unhealthy session could halt all message delivery. The community reacted quickly: the fix PR [#2797](https://github.com/nanocoai/nanoclaw/pull/2797) was proposed and merged on the same day. Strong maintainer responsiveness.

- **Per-message approval policies** – PR [#2793](https://github.com/nanocoai/nanoclaw/pull/2793) (open, 0 comments) introduces a new feature for agent-to-agent communication: optional, directed approval gates. This addresses a common governance need in multi-agent deployments.

- **Korean README translation** – PR [#2806](https://github.com/nanocoai/nanoclaw/pull/2806) (open) adds `README_ko.md`. Demonstrates growing international community engagement.

- **Atlas Cloud LLM backend** – PR [#2717](https://github.com/nanocoai/nanoclaw/pull/2717) (open, 1.5 weeks old) adds documentation for using Atlas Cloud as an OpenAI-compatible backend. Indicates interest in alternative LLM providers.

**Analysis:** The community is focused on reliability (delivery bug), governance (approval policies), localization, and expanding backend options – healthy mix of operational and feature concerns.

## 5. Bugs & Stability

**High severity (fix available):**
- **Session stall kills all delivery** – Issue [#2796](https://github.com/nanocoai/nanoclaw/issues/2796) (critical). **Fixed** by PR [#2797](https://github.com/nanocoai/nanoclaw/pull/2797) (merged).
- **Path traversal in group creation** – PR [#2800](https://github.com/nanocoai/nanoclaw/pull/2800) (open) fixes CWE-22: `ncl groups create --folder ../../etc` can escape `GROUPS_DIR`. Security vulnerability.
- **Unauthorized file read via send_file** – PR [#2799](https://github.com/nanocoai/nanoclaw/pull/2799) (open) fixes CVE-2026-29611: `send_file` can read any file in the container, including credentials. Security vulnerability.

**Medium severity (fix in progress):**
- **CLI `messaging-groups create` always throws** – Issue [#2804](https://github.com/nanocoai/nanoclaw/pull/2804) (PR open). The command fails with `NOT NULL constraint failed`.
- **No request timeout in socket client** – PR [#2802](https://github.com/nanocoai/nanoclaw/pull/2802) (open) fixes a potential indefinite hang.
- **Non-object JSON parsing** – PR [#2801](https://github.com/nanocoai/nanoclaw/pull/2801) (open) fixes `safeParseContent` returning primitives.

**Low severity (documentation):**
- `add-imessage` skill fails if `src/channels/` directory missing – Issue [#2791](https://github.com/nanocoai/nanoclaw/issues/2791) fixed by PR [#2792](https://github.com/nanocoai/nanoclaw/pull/2792) (open).
- Setup skill is a 10-line stub – Issue [#2789](https://github.com/nanocoai/nanoclaw/issues/2789) fixed by PR [#2790](https://github.com/nanocoai/nanoclaw/pull/2790) (open).
- OneCLI port 10254 appears only in troubleshooting – Issue [#2787](https://github.com/nanocoai/nanoclaw/issues/2787) fixed by PR [#2788](https://github.com/nanocoai/nanoclaw/pull/2788) (open).
- Generic title in migrate-nanoclaw skill – Issue [#2785](https://github.com/nanocoai/nanoclaw/issues/2785) fixed by PR [#2786](https://github.com/nanocoai/nanoclaw/pull/2786) (open).

## 6. Feature Requests & Roadmap Signals

| Feature | PR/Issue | Status | Likely in next version |
|---------|----------|--------|------------------------|
| Per-message approval policies for agent-to-agent | [#2793](https://github.com/nanocoai/nanoclaw/pull/2793) | Open | **Strong candidate** – active PR with no major objections |
| CLI-derived dashboard skill (`/add-clidash`) | [#2795](https://github.com/nanocoai/nanoclaw/pull/2795) | Open | Possible – utility skill, low complexity |
| Atlas Cloud LLM backend documentation | [#2717](https://github.com/nanocoai/nanoclaw/pull/2717) | Open (1.5 wks) | Likely – low risk, adds ecosystem |
| Korean README | [#2806](https://github.com/nanocoai/nanoclaw/pull/2806) | Open | High – trivial translation merge |
| Managed-fleet opt-out for upgrade tripwire | [#2780](https://github.com/nanocoai/nanoclaw/pull/2780) | **Merged** | Already in v2.1.17 |
| Restore env-var gateway auth for managed fleets | [#2794](https://github.com/nanocoai/nanoclaw/pull/2794) | **Merged** | Already in v2.1.17 |

**Predictions for next release (v2.2.x):**  
- Per-message approval policies (PR #2793)  
- Security hardening (path traversal fix #2800, file read fix #2799)  
- Continued documentation improvements from community PRs  

## 7. User Feedback Summary

- **Pain point – delivery reliability:** User `mashkovtsevlx` reported that a single bad session could halt all message delivery. The fix was met with approval (PR #2797 merged). This suggests production users are hitting edge cases with multi-agent workloads.
- **Pain point – documentation gaps:** Several issues from user `specterslient95-lgtm` (4 documentation bugs) indicate that new users find skill instructions incomplete or unclear. Quick acceptance of fix PRs shows the project values onboarding quality.
- **Positive – internationalization:** A Korean README translation was contributed and welcomed. The user `arkjun` is a new contributor.
- **Feature demand – LLM flexibility:** The Atlas Cloud PR (#2717) from `lucaszhu-hue` suggests users want to avoid vendor lock-in and seek low-cost or specialized backends.
- **Governance desire:** The approval-policy feature (#2793) by `moshe-nanoco` indicates enterprise/team use cases are emerging.

## 8. Backlog Watch

| Item | Age | Last Updated | Notes |
|------|-----|--------------|-------|
| [#2750](https://github.com/nanocoai/nanoclaw/pull/2750) – fix stale outbound.db journals | 6 days | Updated today | Open PR with fix for two related bugs (#2516, #2640). Maintainers should prioritize review – it addresses crash-after-kill scenarios. |
| [#2717](https://github.com/nanocoai/nanoclaw/pull/2717) – Atlas Cloud LLM backend docs | 9 days | Updated today | No maintainer comment yet. Low risk, high value. |
| [#2793](https://github.com/nanocoai/nanoclaw/pull/2793) – per-message approval policies | 1 day | Today | Large feature PR; needs maintainer review and possibly design discussion. |
| [#2800](https://github.com/nanocoai/nanoclaw/pull/2800) – path traversal fix | 1 day | Today | Security fix – should be merged quickly. |
| [#2799](https://github.com/nanocoai/nanoclaw/pull/2799) – send_file read restriction | 1 day | Today | Security fix (CVE assigned) – urgent. |

No abandoned issues or PRs older than 2 weeks were identified.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest – 2026-06-18

## Today’s Overview

The NullClaw repository saw moderate activity over the past 24 hours, with no new releases and no merged pull requests. Three issues remain open and were recently updated, while two pull requests (both still open) address key usability and feature gaps: memory recall configuration and CLI arrow‑key handling. The project’s core functionality appears functional, but lingering bugs and documentation gaps — especially around scheduling and headless Web UI setup — continue to affect user experience. Overall, the project is in a maintenance‑plus‑feature phase, with community contributions helping to fill critical gaps.

## Releases

No new releases were published today.

## Project Progress

No pull requests were merged or closed in the last 24 hours. The two open PRs (see below) are awaiting review or further refinement.

## Community Hot Topics

The following issues and PRs have generated the most comment activity (2 comments each):

- **[#915 – [bug] Problem with scheduler unauthorized](https://github.com/nullclaw/nullclaw/issues/915)**  
  The scheduler fails when using an external Ollama host on the same network. The user reports that tool calling works generally, but the scheduler does not function in Telegram or the chatbot CLI. This points to a possible authentication or API‑routing issue between NullClaw and the remote LLM.

- **[#865 – [bug] CLI shows ctrl characters for up/down/left/right keys](https://github.com/nullclaw/nullclaw/issues/865)**  
  A CLI usability bug where arrow keys are not interpreted as terminal navigation commands but instead print control characters. This prevents users from using command history or editing lines effectively.

- **[#861 – How to enable the Web UI on headless VPS server?](https://github.com/nullclaw/nullclaw/issues/861)**  
  A documentation request (1 comment) asking for a plain‑language explanation of setting up the Web UI on a headless VPS, particularly the “tunneled browser” concept.

The underlying need across these threads is **better documentation and smoother out‑of‑the‑box user experience**, especially for non‑expert Linux users deploying the agent remotely.

## Bugs & Stability

Two bugs were actively discussed today, ranked by severity:

| Severity | Issue | Description | Fix PR exists? |
|----------|-------|-------------|----------------|
| High     | [#915](https://github.com/nullclaw/nullclaw/issues/915) | Scheduler fails with external Ollama – blocks core automation feature. | No |
| Medium   | [#865](https://github.com/nullclaw/nullclaw/issues/865) | CLI shows raw control characters – impairs interactive use. | Yes: [#960](https://github.com/nullclaw/nullclaw/pull/960) |

The fix for the CLI arrow‑key issue is proposed in PR #960, which introduces a line editor using raw‑mode input. No crashes or regressions were reported in the last 24 hours.

## Feature Requests & Roadmap Signals

Two clear feature signals emerged:

- **Configurable memory recall** (PR [#961](https://github.com/nullclaw/nullclaw/pull/961)) adds three configuration keys (`auto_recall`, `recall_limit`, `max_context_bytes`). This gives users control over memory enrichment and context window sizing — a practical enhancement for power users managing token budgets and performance.

- **Headless Web UI deployment** (Issue [#861](https://github.com/nullclaw/nullclaw/issues/861)) is a pain point that suggests the next version may need either a simplified setup wizard, Docker‑based defaults, or a dedicated configuration option for tunneling (e.g., Tailscale, Cloudflare Tunnel).

Both changes are likely candidates for the next minor release.

## User Feedback Summary

Users report the following real‑world pain points:

- **External LLM integration not fully reliable** – the feature works for standard tool calls, but the scheduler fails without clear error messages.
- **CLI is unusable for non‑expert terminal users** – missing basic line editing and history navigation makes daily interaction frustrating.
- **Setup documentation for remote/headless scenarios is too technical** – users want step‑by‑step instructions, not jargon‑heavy references.

Satisfaction appears moderate: users are actively trying NullClaw and contributing bug reports/PRs, but several open issues go unanswered for weeks (e.g., #861, #865, #915), which may erode confidence. The quick PR from the community for the CLI fix (#960) is a positive sign.

## Backlog Watch

Three issues have been open for over a month without a maintainer reply:

- **[#861](https://github.com/nullclaw/nullclaw/issues/861)** – Web UI on headless VPS (created Apr 22, 1 comment, no maintainer response)
- **[#865](https://github.com/nullclaw/nullclaw/issues/865)** – CLI arrow keys (created Apr 23, 2 comments, now with a fix PR)
- **[#915](https://github.com/nullclaw/nullclaw/issues/915)** – Scheduler unauthorized (created May 15, 2 comments, no fix PR yet)

These issues represent **high‑impact usability and reliability blockers** that are likely deterring new users. A maintainer response or priority triage is recommended.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-06-18

## 1. Today's Overview
The project remains highly active with 50 pull requests and 11 issues updated in the last 24 hours, including 17 closed/merged PRs and 4 closed issues. No new releases were published. Development momentum is strong, driven by a mix of core team feature work (Projects stack, Slack OAuth hardening, no-progress detection) and community contributions (Bedrock integration, skill extraction). A long-running dependency bump PR (#4876) continues to linger, and several user-reported bugs around agent recovery and UI feedback remain open, indicating ongoing polish work.

## 2. Releases
No new releases were recorded today. The last release appears to be PR #3708 (closed yesterday) which bumped `ironclaw` from 0.24.0 → 0.29.1 and included breaking changes in `ironclaw_common` and `ironclaw_skills`. Users upgrading should review the [changelog](https://github.com/nearai/ironclaw/pull/3708) for migration notes.

## 3. Project Progress
The following PRs were merged or closed today (selected notable items):

- **#3708** – `chore: release` (automated release bump, closed)  
- **#5000** – `feat(agent-loop): content-digest plumbing for output-aware progress (PR2)` (inert plumbing merged)  
- **#5022** – `feat(agent-loop): output-aware no-progress detection (PR3)` (keystone of no-progress redesign, merged)  
- **#5052** – `fix(reborn): live Slack OAuth path structural DM-parity` (closes #5009, security hardening)

Additionally, four issues were closed: #4823 (UI feedback on delete), #4793 (first-run onboarding question), #4952 (Slack auth stale record), #4974 (duplicate action buttons). The Projects stack (PRs #5015–#5019, size XL, 5 parts) remains open and under review.

## 4. Community Hot Topics
While comment counts are low across the board, the following items attracted the most attention:

- **Issue #4761** – `[bug] [Reborn] Agent stops after repeated tool failures instead of recovering` [🔗](https://github.com/nearai/ironclaw/issues/4761)  
  3 comments. This is a week-old bug that blocks agents from completing multi-step tasks. The underlying need is robust error recovery.

- **Issue #4879** – `IronClaw Reborn Local Dogfooding Findings 06/15–06/21` [🔗](https://github.com/nearai/ironclaw/issues/4879)  
  A central tracking issue for usability problems found during internal dogfooding. Signals that the team is actively stress-testing the Reborn experience.

- **PR #5061** – `feat(reborn): skill extraction & self-evolution with activation controls` [🔗](https://github.com/nearai/ironclaw/pull/5061)  
  A new contributor PR introducing Hermes-style skill distillation. No comments yet, but the feature is ambitious and likely to draw reviewer attention.

- **PR #4876** – `build(deps): bump the everything-else group with 43 updates` [🔗](https://github.com/nearai/ironclaw/pull/4876)  
  A large dependency update PR that has been open since June 14. Currently has no comments but is critical for keeping the project’s supply chain current.

## 5. Bugs & Stability
New bug reports filed today, ranked by severity:

| Bug | Severity | Description | Fix PR? |
|-----|----------|-------------|---------|
| **#5060** – GitHub analysis workflows may enter repeated approval loops [🔗](https://github.com/nearai/ironclaw/issues/5060) | **High** | Blocks automated issue tracking analysis; no results produced. | None yet |
| **#5058** – Bedrock unreachable from ironclaw-reborn binary + Converse tool-schema rejects top-level combinators [🔗](https://github.com/nearai/ironclaw/issues/5058) | **High** | AWS Bedrock completely unusable with standalone binary. | **PR #5059** (open) |
| **#5007** – Skills validation error does not clear after required fields filled [🔗](https://github.com/nearai/ironclaw/issues/5007) | **Medium** | UI validation state bug, confusing for users. | None yet |
| **#4761** – Agent stops after repeated tool failures (updated yesterday, still open) [🔗](https://github.com/nearai/ironclaw/issues/4761) | **Medium** | Agent fails to recover from transient tool errors. | None yet |

Previously reported bugs fixed today: #4823 (delete feedback), #4974 (duplicate buttons), #4952 (stale auth flow), #4793 (onboarding question closed).

## 6. Feature Requests & Roadmap Signals
Several significant features are in flight, indicating the direction of the next release:

- **Skill extraction and self-evolution** – PR #5061 (new contributor) proposes automatic skill distillation from successful conversations, with safety scans. If merged, this could be a major usability enhancement.
- **Projects entity** – Stack of 5 PRs (#5015–#5019) adds full CRUD for projects and project membership, introduced by the core team. Likely to land in the next release.
- **Read-only filesystem viewer** – PR #5057 adds an agent file browser to WebChat v2, improving transparency.
- **AWS Bedrock support** – PR #5059 fixes the missing `bedrock` feature in the standalone binary, answering user demand for cloud LLM backends.
- **Output-aware no-progress detection** – PRs #5000 and #5022 (already merged) improve agent loop robustness by detecting stalled progress via content digests.

These features suggest that the upcoming release will focus on **agent introspection** (file viewer, skill logs), **team collaboration** (projects), and **platform extensibility** (Bedrock).

## 7. User Feedback Summary
User pain points (primarily reported by `sunglow666` and dogfooding logs) include:

- **Agent recovery failures** – Issue #4761: agents hang after repeated tool failures instead of retrying or degrading gracefully. This is a core reliability concern.
- **UI feedback gaps** – Closing conversations that are still running (#4823, now fixed) and missing error clearing in validation forms (#5007) erode trust.
- **First-run friction** – Dogfooding issue #4879 tracks onboarding hurdles such as the welcome page blocking Extensions/Automations (#4793, resolved by discussion).
- **Cloud provider integration** – AWS Bedrock ( #5058 ) and `auto` model resolution (PR #5045, still open) show that users want seamless multi-provider support.

Overall sentiment: active frustration with agent stability and UI polish, but enthusiasm for new features like skill extraction and project management.

## 8. Backlog Watch
The following items require maintainer attention:

- **PR #4876** – `build(deps): bump the everything-else group with 43 updates` [🔗](https://github.com/nearai/ironclaw/pull/4876)  
  Open since June 14 with no reviewer activity. A large batch update that risks merge conflicts and should be prioritized to avoid dependency drift.

- **Issue #4761** – `Agent stops after repeated tool failures` [🔗](https://github.com/nearai/ironclaw/issues/4761)  
  Open for a week with no fix PR. Core reliability issue that affects all users running complex workflows.

- **Issue #5007** – `Skills validation error does not clear` [🔗](https://github.com/nearai/ironclaw/issues/5007)  
  Minor but reported twice; no PR assigned yet.

- **PR #5061** – New contributor skill extraction feature [🔗](https://github.com/nearai/ironclaw/pull/5061)  
  Needs code review and design feedback; no comments yet despite being filed today. Should be triaged quickly to encourage first-time contributors.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-06-18

## 1. Today's Overview
Project activity remains high, with 10 pull requests merged/closed in the last 24 hours and one new release (2026.6.15) published three days ago. The repository currently has one open issue, a **high-severity security vulnerability** reported today (#2176). The majority of merged PRs focus on the `cowork` subsystem—fixing scroll behavior, voice input, model selection, streaming metadata, and context compaction. The development pace is steady, but the disclosed security flaw demands immediate maintainer attention.

---

## 2. Releases
**Latest release: [LobsterAI 2026.6.15](https://github.com/netease-youdao/LobsterAI/releases/tag/2026.6.15)** (June 15, 2026)

**What’s Changed:**
- **feat: add computer use** – Enables the agent to interact with desktop applications (similar to Anthropic’s computer use feature).
- **feat(cowork): add realtime ASR voice input** – Live speech recognition for the Cowork mode.
- **feat(cowork): improve post-compaction context continuity** – Better conversation history compression to maintain agent task coherence after OpenClaw compaction.

**Breaking Changes:** None explicitly documented.  
**Migration Notes:** Users should update to 2026.6.15 to benefit from the new features and stability fixes. No manual migration steps are required.

---

## 3. Project Progress (Merged/Closed PRs in Last 24h)
All 10 pull requests were merged or closed within the past day. Major areas of improvement:

| PR # | Title & Area | Summary |
|------|--------------|---------|
| #2175 | [chore: optimize readme](https://github.com/netease-youdao/LobsterAI/pull/2175) | Documentation polishing. |
| #2174 | [fix(cowork): settle scroll-to-bottom position](https://github.com/netease-youdao/LobsterAI/pull/2174) | Keeps scroll aligned with latest message; cleans up timers on session change. |
| #2162 | [fix(cowork): preserve voice input cancel guard](https://github.com/netease-youdao/LobsterAI/pull/2162) | Resolves merge conflict to keep realtime-only ASR flow while preserving cancel/guard logic. |
| #2153 | [fix(cowork): preserve same-name package model selection](https://github.com/netease-youdao/LobsterAI/pull/2153) | Fixes model selection ambiguity when package and custom models share names. |
| #2154 | [fix(cowork): show model metadata after stopped streams](https://github.com/netease-youdao/LobsterAI/pull/2154) | Ensures model metadata is displayed after user stops a stream manually. |
| #2149 | [fix(openclaw): raise gateway heap limit](https://github.com/netease-youdao/LobsterAI/pull/2149) | Prevents OOM crashes under heavy multi-channel workloads. |
| #2147 | [fix(cowork): prevent stopped startup turns from sending chat](https://github.com/netease-youdao/LobsterAI/pull/2147) | Cancels OpenClaw turn startup when a stop command arrives early. |
| #2145 | [feat(cowork): improve post-compaction context continuity](https://github.com/netease-youdao/LobsterAI/pull/2145) | Adds LobsterAI-owned continuity layer around OpenClaw compaction. |
| #2144 | [fix(auth): update portal fallback URLs](https://github.com/netease-youdao/LobsterAI/pull/2144) | Points local fallbacks to new production portal domains. |
| #1463 | [fix long modal titles (stale)](https://github.com/netease-youdao/LobsterAI/pull/1463) | Truncates long runtime titles in modal headers to prevent layout overflow. |

**Key theme:** Continuous reliability improvements for the Cowork/OpenClaw subsystem, with emphasis on voice input integration, scroll handling, and memory management.

---

## 4. Community Hot Topics
The only open issue is a **security advisory** filed today:

- **#2176 [Security] LobsterAI automatic artifact loading allows message-derived arbitrary local file reads**  
  [Issue Link](https://github.com/netease-youdao/LobsterAI/issues/2176)  
  **Reporter:** YLChen-007 | **Comments:** 1 | **Reactions:** 0  
  **Summary:** LobsterAI automatically parses `MEDIA:` file references from assistant or tool output and forwards the resulting file path into a privileged Electron context, enabling arbitrary local file reads.

**Analysis:** This is a critical vulnerability that could allow an attacker (via crafted assistant output) to read arbitrary files on the host system. The single comment likely comes from a maintainer acknowledging the report. No fix PR exists yet. This will likely be the top priority for the next patch release.

---

## 5. Bugs & Stability
**Critical:**
- **#2176 – Arbitrary local file read via automatic artifact loading** – Severity: **High/Critical**. No fix available yet. Users are advised to avoid using automated artifact loading until a patch is released.

**Medium (addressed in PRs):**
- **OpenClaw OOM crashes under high load** – Fixed by raising the gateway heap limit ([PR #2149](https://github.com/netease-youdao/LobsterAI/pull/2149)).
- **Stopped streams missing model metadata** – Fixed ([PR #2154](https://github.com/netease-youdao/LobsterAI/pull/2154)).
- **Startup race causing unintended chat sends** – Fixed ([PR #2147](https://github.com/netease-youdao/LobsterAI/pull/2147)).

**Low:**
- Long modal titles causing layout overflow – Fixed ([PR #1463](https://github.com/netease-youdao/LobsterAI/pull/1463)).
- Stale scroll-to-bottom behavior – Fixed ([PR #2174](https://github.com/netease-youdao/LobsterAI/pull/2174)).

---

## 6. Feature Requests & Roadmap Signals
No feature requests were filed as issues today. However, the latest release and recent PRs signal the following roadmap directions:

- **Computer Use** – The agent now has the ability to control desktop applications (released in 2026.6.15).
- **Enhanced Voice Interaction** – Real-time ASR voice input for Cowork mode is now live.
- **Context Continuity** – The team is investing heavily in making the agent remember task state across compressed conversations, suggesting a focus on long-running agent tasks.

**Prediction for next release:** The next version (2026.6.18 or later) will almost certainly include a fix for the security vulnerability #2176. Additional stability patches for Cowork and OpenClaw are likely.

---

## 7. User Feedback Summary
- **Pain point – Security:** The discovery of #2176 indicates that at least one user (likely a security researcher) has identified a dangerous bypass of sandboxing. This reflects a potential trust concern for users who rely on automatic artifact loading.
- **Pain point – Stability:** Multiple fixes today target crashes (OOM, stream metadata loss, unintended chat sends) – these affect users running long Cowork sessions or using voice input.
- **Use case – Voice input & computer use:** The release notes show strong demand for multimodal interaction (speech + screen control). Early adopters are likely satisfied with the new capabilities but may encounter edge-case bugs.

No explicit satisfaction/dissatisfaction comments are present in the data.

---

## 8. Backlog Watch
- **PR #1463** (stale, opened April 4, closed today) – While it was successfully merged after 2.5 months, the long delay suggests a pattern of slow review for non-critical fixes. No other long-unanswered issues are visible in today’s snapshot.
- **No unresolved issues older than 24h** – The only open issue (#2176) was filed today and is actively under discussion. Maintainer responsiveness appears good for security reports.

**Recommendation:** Monitor the security issue’s progress and ensure a hotfix release is published promptly. The project health is otherwise robust, with frequent commits and releases.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

# TinyClaw Project Digest – 2026-06-18

## 1. Today’s Overview
The TinyClaw project saw no new releases or pull requests in the last 24 hours, indicating a low development activity day. However, three critical security vulnerabilities were reported by an external researcher (YLChen‑007) in a single batch, all opened today and remaining unaddressed. These issues expose the project’s authentication and file-handling mechanisms, requiring immediate maintainer attention. Overall project health is stable in terms of velocity, but the security posture is concerning and will likely dominate the next sprint.

## 2. Releases
*No new releases were published.*
Omit section.

## 3. Project Progress
No pull requests were merged or closed today. No feature implementations or bug fixes were advanced.

## 4. Community Hot Topics
All three open issues were created today and have zero comments or reactions, so there is no active community discussion yet. The underlying need is clear: each issue describes a serious security flaw that could allow remote attackers to perform unauthorized actions or exfiltrate sensitive data. The community (and the security researcher) are likely awaiting maintainer acknowledgment and patches.

- **Issue #284** – [Security] Unauthenticated API allows Claude invocation with provider permission checks disabled  
  https://github.com/TinyAGI/tinyagi/issues/284
- **Issue #283** – [Security] Unauthenticated `prompt_file` config allows local file disclosure to the model provider  
  https://github.com/TinyAGI/tinyagi/issues/283
- **Issue #282** – [Security] Untrusted `[send_file: …]` response tags allow arbitrary host file attachment delivery  
  https://github.com/TinyAGI/tinyagi/issues/282

## 5. Bugs & Stability
All three reports are security vulnerabilities, each ranked as **critical** due to potential remote exploitation without authentication. No fix PRs exist yet. Summary by severity (all critical):

| ID | Title | Impact |
|----|-------|--------|
| #284 | Unauthenticated API messages invoke Claude with disabled permission checks | Attacker can call Claude endpoints without auth, bypassing provider permission checks. |
| #283 | `prompt_file` agent config allows arbitrary file disclosure | Unauthenticated attacker can read arbitrary local files and send them to an external LLM provider. |
| #282 | `[send_file: …]` response tags allow arbitrary file attachment | Attacker can influence provider output to deliver files from the host to the user (or vice versa). |

No other bugs (crashes, regressions) were reported.

## 6. Feature Requests & Roadmap Signals
No feature requests were recorded today. The only signals are the security issues, which likely require urgent architectural changes to authentication and input validation before any new features can be safely added.

## 7. User Feedback Summary
No direct user feedback was captured today. The three issues, however, originate from a security researcher (YLChen‑007) and highlight real pain points:
- Lack of authentication on critical API endpoints.
- Insufficient validation of file paths and response tags.
- Default configurations that disable security checks.

These points indicate dissatisfaction with the project’s current security posture and serve as a strong directive for the maintainers.

## 8. Backlog Watch
All three open issues are brand new (created 2026-06-18) and have received no maintainer response. They should be triaged immediately. No older, long‑unanswered items were updated today. The maintainers should prioritize at least acknowledging these reports to prevent further risk of exploitation.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

## Moltis Project Digest — 2026-06-18

### Today's Overview
The project shows low but focused activity over the past 24 hours. Two enhancement issues remain open, and one open pull request aims to improve configurability of the WebUI RPC timeout. No new releases have been published, and no bugs or regressions were reported. The community is primarily requesting user-facing quality-of-life improvements, indicating the project is in a steady feature-request phase rather than a crisis management mode.

---

### Releases
No new releases today. The latest published version remains unchanged.

---

### Project Progress
No pull requests were merged or closed in the last 24 hours. The only open PR is **#1130** (linked below), which addresses configurable RPC timeouts for the WebUI — a response to issue #1127 (not shown in today’s data). This PR is still under review.

- [PR #1130: feat: make webui rpc timeout configurable](https://github.com/moltis-org/moltis/pull/1130)

---

### Community Hot Topics
The most active discussion is **Issue #1126** with three comments, requesting configurability of the TTS output format. The request indicates users want to control whether TTS returns raw audio bytes, a file path, or another structured format. No maintainer response has been recorded yet.

- [Issue #1126: allow to configure the format of tts output](https://github.com/moltis-org/moltis/issues/1126)

**Issue #1131** (copy + export as Markdown) has zero comments, but represents a clear user desire to export agent conversations in a portable, human-readable format.

- [Issue #1131: Add copy + export as Markdown](https://github.com/moltis-org/moltis/issues/1131)

Both issues are tagged as `enhancement`, and the community seems to be focusing on output formatting and interoperability features.

---

### Bugs & Stability
No bug reports, crashes, regressions, or stability-related issues were updated in the last 24 hours. The project appears stable on the reported front.

---

### Feature Requests & Roadmap Signals
Two feature requests surfaced today:
- **Configurable TTS output format** (#1126): likely driven by users integrating Moltis with custom pipelines or external audio players.
- **Copy/export conversations as Markdown** (#1131): suggests demand for better shareability and documentation of agent interactions.

Additionally, **PR #1130** (configurable RPC timeout) shows the team is actively improving WebUI reliability. If merged, this could be part of the next minor release.

*Prediction:* The TTS format configuration (Issue #1126) and the Markdown export (Issue #1131) are moderate-effort features that could land in the next version, assuming maintainer bandwidth allows. The RPC timeout PR is likely to be merged soon given it already has a fix.

---

### User Feedback Summary
User feedback today is entirely feature-oriented:
- **Pain point:** TTS output is currently fixed-format, limiting integration with custom audio backends or storage.
- **Use case:** Exporting chat transcripts as Markdown for documentation, note-taking, or review.
- **Satisfaction:** No negative feedback was recorded; the community appears constructive.

---

### Backlog Watch
No long-unanswered issues or PRs are present in today’s data. The oldest open item is Issue #1126 (created June 16), which has received comments but no maintainer reply yet. It would benefit from a quick acknowledgment or a request for more details from the team.

- [Issue #1126 (2 days old, no maintainer response)](https://github.com/moltis-org/moltis/issues/1126)

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest – 2026-06-18

## 1. Today's Overview
CoPaw showed high activity on 2026-06-18, with **45 issues updated** (24h) — 34 closed and 11 still open — and **50 pull requests updated** — 33 merged/closed and 17 open. Two new releases landed: **v1.1.12** (stable) and **v1.1.12-beta.2** (pre-release), bringing major UI improvements and performance optimizations. The community reported several regressions in Docker upgrades, subagent approval routing, and model selection, though the project team merged fixes quickly for many of these. Overall project health appears strong, with active triage and a growing contributor base (multiple first-time contributors in today's PRs).

## 2. Releases
**v1.1.12** (stable, [release notes](https://github.com/agentscope-ai/QwenPaw/releases/tag/v1.1.12))  
- **Console**: Entire Models Page overhaul – provider aggregation, unified card UI, layout redesign ([#5203](https://github.com/agentscope-ai/QwenPaw/pull/5203)).  
- **Simple Mode**: Added flat navigation and session list sorted by update time ([#5222](https://github.com/agentscope-ai/QwenPaw/pull/5222)).  
- No breaking changes noted; migration is a normal update.

**v1.1.12-beta.2** (pre-release, [release notes](https://github.com/agentscope-ai/QwenPaw/releases/tag/v1.1.12-beta.2))  
- Performance: removed unnecessary deep copy in agent config handling ([#5240](https://github.com/agentscope-ai/QwenPaw/pull/5240)).  
- Console: added session filter by title ([#5178](https://github.com/agentscope-ai/QwenPaw/pull/5178)).  
- Chore: dependency bumps.

Both releases are compatible with existing configurations and skill/MCP assets.

## 3. Project Progress
Today **33 PRs were merged/closed** (plus 17 open). Key merged changes:  
- **Backup robustness**: skip unreadable files instead of failing whole backup ([#5041](https://github.com/agentscope-ai/QwenPaw/pull/5041), fixes [#4916](https://github.com/agentscope-ai/QwenPaw/issues/4916)).  
- **Memory stability**: rename ChromaDB probe collection to `'probe-test'` to avoid conflicts ([#5289](https://github.com/agentscope-ai/QwenPaw/pull/5289)).  
- **XiaoYi channel**: dual WebSocket refactor and A2A protocol alignment ([#5274](https://github.com/agentscope-ai/QwenPaw/pull/5274), [#3839](https://github.com/agentscope-ai/QwenPaw/pull/3839)).  
- **Duplicate session_id in filenames**: fixed when `user_id` equals `session_id` ([#5026](https://github.com/agentscope-ai/QwenPaw/pull/5026), fixes [#5025](https://github.com/agentscope-ai/QwenPaw/issues/5025)).  
- **Proactive responder**: prevented cache pollution of `load_agent_config()` ([#5275](https://github.com/agentscope-ai/QwenPaw/pull/5275)).  
- **Version bumps**: released v1.1.12 and opened `2.0.0a1` milestone branch ([#5280](https://github.com/agentscope-ai/QwenPaw/pull/5280), [#5281](https://github.com/agentscope-ai/QwenPaw/pull/5281)).  
- **Scripts**: corrected prerelease arguments expansion ([#5288](https://github.com/agentscope-ai/QwenPaw/pull/5288)).  

Open PRs of interest:  
- Chat history right-side panel ([#5293](https://github.com/agentscope-ai/QwenPaw/pull/5293)).  
- Cron update CLI command (`cron update <job_id>`) ([#5210](https://github.com/agentscope-ai/QwenPaw/pull/5210)).  
- OpenClaw config migration tool ([#5276](https://github.com/agentscope-ai/QwenPaw/pull/5276)).  
- Timeout protection for context compaction ([#5242](https://github.com/agentscope-ai/QwenPaw/pull/5242)).  
- Customizable column order in sessions page ([#4975](https://github.com/agentscope-ai/QwenPaw/pull/4975)).  

## 4. Community Hot Topics
The most active discussion today:  

- **#280 – "Which Skills and MCPs Can Be Built-in?"** (27 comments, closed, [link](https://github.com/agentscope-ai/QwenPaw/issues/280))  
  *Underlying need*: Users want popular skills/MCPs pre-installed for an out-of-the-box experience. Although closed, this continues to attract discussion as more providers emerge.  

- **#4108 – "为什么新版本的webui这么卡"** (8 comments, closed, [link](https://github.com/agentscope-ai/QwenPaw/issues/4108))  
  *Severe UI lag on Windows* during reply generation (CPU/memory spikes). User reported it makes multi-tasking impossible.  

- **#5262 – "每次升级之后，被禁用的内置技能又会重新变回启用"** (5 comments, open, [link](https://github.com/agentscope-ai/QwenPaw/issues/5262))  
  Regression: built-in skills re-enabled after every upgrade. Users want persistent disable state.  

- **#5295 – "Subagent approval request not pushed to external channel"** (1 comment, open, [link](https://github.com/agentscope-ai/QwenPaw/issues/5295))  
  *Cross-agent safety gap*: subagent approval notifications are not forwarded to external channels (QQ, etc.), causing silent failures.  

- **#5204 – "两个 QwenPaw Agent 通过 Matrix 互聊时陷入无限循环"** (2 comments, open, [link](https://github.com/agentscope-ai/QwenPaw/issues/5204))  
  *Infinite loop*: two agents via Matrix endlessly wake each other – no runtime circuit breaker exists for cross-agent feedback loops.  

## 5. Bugs & Stability
High severity:  
1. **Subagent approval route failure** ([#5295](https://github.com/agentscope-ai/QwenPaw/issues/5295)) – critical for safety. No fix PR yet.  
2. **UI freeze during generation** ([#4108](https://github.com/agentscope-ai/QwenPaw/issues/4108)) – high user impact on Windows. Closed but root cause likely still latent.  
3. **Matrix cross-agent infinite loop** ([#5204](https://github.com/agentscope-ai/QwenPaw/issues/5204)) – no fix proposed.  

Medium severity:  
4. **Upgrade resets disabled built-in skills** ([#5262](https://github.com/agentscope-ai/QwenPaw/issues/5262)) – annoyance for power users.  
5. **Manual model providers not shown in chat page** ([#5292](https://github.com/agentscope-ai/QwenPaw/issues/5292), closed) – fixed? status unclear.  
6. **Docker upgrade failure (v0.2.0)** – several duplicates (e.g., [#2229](https://github.com/agentscope-ai/QwenPaw/issues/2229), [#2254](https://github.com/agentscope-ai/QwenPaw/issues/2254)) – closed, likely resolved in newer versions.  
7. **Context compaction crash when summary exceeds maxLength** – fix PR [#5287](https://github.com/agentscope-ai/QwenPaw/pull/5287) open.  
8. **DingTalk SSL with `uv tool install`** – fix PR [#5291](https://github.com/agentscope-ai/QwenPaw/pull/5291) open.  

Low severity:  
- Cron misfire grace window too short ([#5241](https://github.com/agentscope-ai/QwenPaw/pull/5241) – PR increases from 60s to 3600s).  
- Discord.py import error on Windows ([#5290](https://github.com/agentscope-ai/QwenPaw/issues/5290), open).  

## 6. Feature Requests & Roadmap Signals
Requests with strong user demand (based on PRs and issue upvotes):  

- **Built-in skills/MCPs catalog** ([#280](https://github.com/agentscope-ai/QwenPaw/issues/280)) – likely to appear as a community-driven registry in v1.2.  
- **HTTP API for headless operation** ([#2202](https://github.com/agentscope-ai/QwenPaw/issues/2202)) – requested for multi-agent automation.  
- **Web console upgrade mechanism** ([#2235](https://github.com/agentscope-ai/QwenPaw/issues/2235)) – remote one-click update.  
- **Merge same-brand provider cards** ([#4965](https://github.com/agentscope-ai/QwenPaw/issues/4965)) – UI clutter reduction, already linked to PR [#5203](https://github.com/agentscope-ai/QwenPaw/pull/5203) (merged).  
- **Cron update/modify command** ([#4939](https://github.com/agentscope-ai/QwenPaw/issues/4939), PR [#5210](https://github.com/agentscope-ai/QwenPaw/pull/5210) open) – likely in next minor release.  
- **Scheduled task history** ([#1366](https://github.com/agentscope-ai/QwenPaw/issues/1366)) – user even contributed a local branch.  
- **OpenClaw migration** ([#5254](https://github.com/agentscope-ai/QwenPaw/issues/5254), PR [#5276](https://github.com/agentscope-ai/QwenPaw/pull/5276) open) – will bring users from similar ecosystems.  

Roadmap update PR [#5277](https://github.com/agentscope-ai/QwenPaw/pull/5277) hints at a **v2.0.0a1** milestone (alpha) already started ([#5281](https://github.com/agentscope-ai/QwenPaw/pull/5281)).

## 7. User Feedback Summary
Real pain points captured today:  
- **Docker upgrade breaks** – multiple users reported web UI unresponsive after upgrading to v0.2.0 (e.g., [#2229](https://github.com/agentscope-ai/QwenPaw/issues/2229), [#2254](https://github.com/agentscope-ai/QwenPaw/issues/2254)).  
- **Write_file truncation** ([#1563](https://github.com/agentscope-ai/QwenPaw/issues/1563)) – large files only partially written; tool forgets `file_path` param.  
- **Multi-agent concurrency** – several reports about model loading failures in non-default agents ([#2174](https://github.com/agentscope-ai/QwenPaw/issues/2174)), request serialization when opening multiple tabs ([#2116](https://github.com/agentscope-ai/QwenPaw/issues/2116)).  
- **Skill auto-disable on import** – users find it annoying that custom skills are disabled after each upgrade ([#3090](https://github.com/agentscope-ai/QwenPaw/issues/3090)).  
- **Positive feedback**: Simple mode and model page redesign are well-received; community contributions (e.g., Venice AI provider, XiaoYi channel improvements) show growing engagement.  

## 8. Backlog Watch
Issues/PRs that need maintainer attention:  

- **#280 – Built-in skills discussion** (closed but still gathering comments) may need a formal proposal.  
- **#1258 – iFlow model no response** (closed, but no conclusive resolution) – community may benefit from a documented workaround.  
- **#1167 – httpx[socks] missing on Ubuntu 24.04** (closed, but one-line install still affected in older versions).  
- **#5204 – Matrix infinite loop** (open, 2 comments) – critical design flaw with no assignee.  
- **#5295 – Subagent approval routing** (open, 1 comment) – security-sensitive, needs prompt triage.  
- **#5276 – OpenClaw migration PR** (open, no reviewer yet) – valuable for onboarding users.  
- **#5287 – Compaction summary overflow** (open, first-time contributor) – needs review to encourage new contributors.  

No issue older than 30 days without any comment from maintainers; triage appears healthy.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest – 2026-06-18

## Today’s Overview

ZeroClaw saw **very high activity** on 2026-06-18: 17 issues and 50 pull requests were updated in the last 24 hours, with 11 PRs merged or closed. No new releases were published. The project continues to make strong progress on its **v0.8.2 (skills + WASM plugins)** and **v0.9.0 (auth/security)** milestones. A major stacked PR series for alias rename-and-cascade utility was completed and merged, while several critical bugs—particularly around Windows self-update, approval attribution, and session history—received fixes. The community remains engaged, with several RFCs and feature proposals under discussion.

---

## Releases

No new releases today. (Latest release information not available.)

---

## Project Progress

### Merged/Closed PRs Today (11 total)

Notable merged/closed pull requests from the last 24 hours:

- **#7842** (merged) – `feat(cli): agents/providers/channels CRUD + skill-bundle cascade` – Completes the 8-PR stacked series for typed delete-with-cascade and alias rename.
- **#7841** (merged) – `feat(gateway): agent owned-state rename cascade + rename wiring` – Core rename logic for agents with owned state.
- **#7840** (merged) – `feat(config): rename_with_cascade for aliased entries` – Config-level support for renaming aliased providers/agents.
- **#7684** (merged) – `fix(acp): surface history-pruner and turn-cancel as visible events` – Improves ACP event visibility for users.

Other closed PRs include minor fixes and test additions.

### Key Open PRs (still in review, not yet merged)

- **#7853** – `fix(update): repair Windows self-update and harden the update pipeline` – A high-impact fix for a broken binary-swap path on Windows.
- **#7492** – `feat(cost): support cached input token pricing from OpenAI-compatible` – Adds cost-awareness for cached tokens.
- **#7821** – `feat(config): add schema struct & risk field` – Introduces SandboxPolicyConfig; needs author action.
- **#7098** – `feat(channel/mattermost): add optional WebSocket listener mode` – Reduces latency by replacing REST polling; needs author action.

---

## Community Hot Topics

### Most Active Issues (by comment count)

1. **#7673 – RFC: Native context compression as a provider pipeline decorator** (3 comments)  
   *Proposes a `CompressionDecorator` to compress ChatRequest payloads before sending to providers. High risk, RFC stage.*  
   [zeroclaw-labs/zeroclaw Issue #7673](https://github.com/zeroclaw-labs/zeroclaw/issues/7673)

2. **#6970 – Tracker: v0.8.1 integration/channel/provider/tool queue and history** (3 comments)  
   *Operational tracker for the v0.8.1 release queue; covers additive channels, providers, tools, and integration-adjacent work.*  
   [zeroclaw-labs/zeroclaw Issue #6970](https://github.com/zeroclaw-labs/zeroclaw/issues/6970)

3. **#7175 – feat(config): typed delete-with-cascade for aliased entries** (2 comments)  
   *Accepted enhancement; the merged PR series addressed this feature.*  
   [zeroclaw-labs/zeroclaw Issue #7175](https://github.com/zeroclaw-labs/zeroclaw/issues/7175)

4. **#7675 – RFC: Hardened CI pipeline — supply-chain scanning, provenance, and SBOM generation** (2 comments)  
   *Proposes a CI security gate for supply-chain security.*  
   [zeroclaw-labs/zeroclaw Issue #7675](https://github.com/zeroclaw-labs/zeroclaw/issues/7675)

### Underlying Needs

The most active RFCs (#7673, #7675) indicate a growing concern with **operational efficiency** (context compression) and **security hardening** (supply-chain CI). The tracker #6970 reflects the community’s continued interest in a stable v0.8.1 release. The wide discussion around alias rename (#7175) has been largely resolved with the merged PR series.

---

## Bugs & Stability

### Bugs Reported or Updated Today (2026-06-18)

| Issue | Severity | Summary | Fix PR? |
|-------|----------|---------|---------|
| **#7907** | **S1 – workflow blocked** | `rename_agent_cascade()` mutates owned state before config persistence, causing data loss or inconsistency. | No fix PR yet. |
| **#7462** | **S2 – degraded behavior** | 74 test failures on Windows (Unix-only commands, path semantics, console encoding). | No fix PR; test suite still Linux-only. |
| **#7737** | **S2 – degraded behavior** | Approval attribution depends on channel-global side channel; concurrent approvals can be lost. | No dedicated fix PR; likely addressed by ongoing work. |
| **#7563** | **S1 – workflow blocked (CLOSED today)** | Canvas-store regression in WS chat/ACP sessions broke `/canvas` after PR #6986. | Fixed and closed. |

### Other Stability Items

- **#7910** (new) – Add Windows runtime test coverage for self-update swap/rollback/sidecar paths – follow-up to #7853.
- **#7897** – Enhancement (not a bug) – Apply security policy updates without full daemon reload.
- **#7901** – `fix(runtime): bound repeated shell approval loops` – Prevents infinite approval prompts for identical shell requests.

**Takeaway:** Windows support remains fragile (#7462, #7910). The newly reported S1 bug #7907 (agent rename persistence) is critical and has no fix in flight. ACP approval attribution (#7737) and the canvas regression (#7563, now fixed) show ongoing runtime stability challenges.

---

## Feature Requests & Roadmap Signals

### New RFCs and Enhancements Today

- **#7897** – `[Feature]: Apply security policy and config updates without full daemon reload` – Zero-downtime reload for security/channel config.
- **#7883** – `[Feature]: Expose intra-family provider fallback notices` – Let users see when fallback switches model within same provider.
- **#7822** – `[RFC]: WASM plugin lifecycle hook subscriptions (PluginCapability::Hook)` – Expand WASM plugin capabilities.
- **#7673** – `RFC: Native context compression` – Provider pipeline decorator for storage/bandwidth savings.
- **#7675** – `RFC: Hardened CI pipeline` – SBOM, supply-chain scanning.

### Milestone Predictions

Based on active trackers and accepted RFCs, **v0.8.2** (skills platform + WASM plugin program) and **v0.9.0** (auth, security, gateway hardening) are the next major releases. The alias rename cascade (#7175) and CLI CRUD (#7842) are likely to ship in **v0.8.2**. The context compression (#7673) and zero-downtime reload (#7897) may target **v0.9.0** or later.

---

## User Feedback Summary

- **Android/Termux support** – Issue #7911 reports that both precompiled binaries and local compilation fail on Android aarch64. User wants to run ZeroClaw on mobile; currently blocked.
- **Windows instability** – Issue #7462 (74 test failures) and the now-fixed #7563 (canvas regression) indicate poor Windows testing coverage. User @NiuBlibing has been actively reporting and fixing Windows issues (#7853, #7910).
- **Approval attribution confusion** – Issue #7737 describes a real-world pain where concurrent approvals can overwrite each other, causing lost decisions.
- **Secret prompt UX** – PR #7856 fixes lack of feedback when entering secrets via dialoguer: users reported not knowing if input was received.

Overall, the community is engaged but frustrated by **Windows first-class support** and certain **race conditions** in the runtime.

---

## Backlog Watch

### Issues/PRs Needing Maintainer Attention

- **#7673 – RFC: CompressionDecorator** – Labeled `needs-author-action`. Author (@ConYel) may need to follow up on feedback.
- **#7821 – feat(config): add risk field** – `needs-author-action` on the PR. Currently blocked from review.
- **#7098 – feat(channel/mattermost): WebSocket listener** – `needs-author-action`. Stalled since June 2; maintainer review requested.
- **#7462 – Windows test failures** – No fix PR despite being accepted and labelled `priority:p1`. Active for over a week; root cause identified but no traction.
- **#7907 – Agent rename persistence** – New S1 bug with no assignee or fix in progress. Should be escalated.

### Long-Standing Items

- **#6970 – v0.8.1 tracker** – Being actively updated; not stale.
- **#7432 – v0.9.0 tracker** – No updates in last 24h but is milestone coordination.
- **#7320 – v0.8.3 tracker** – Last updated 2026-06-17; progressing.

**Note:** Several PRs (e.g., #7835, #7819) carry the `needs-author-action` label, indicating the author has not responded to review comments. Maintainers should prompt authors to unblock these.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*