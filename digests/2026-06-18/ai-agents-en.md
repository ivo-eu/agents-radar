# OpenClaw Ecosystem Digest 2026-06-18

> Issues: 200 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-18 03:18 UTC

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

# OpenClaw Project Digest — 2026-06-18

**Generated from GitHub data (openclaw/openclaw) – activity in last 24h.**

---

## Today’s Overview

OpenClaw remains in an intensely active development cycle. In the past 24 hours, **200 issues** were updated (184 open, 16 closed) and **500 pull requests** were updated (414 open, 86 merged/closed). No new releases were published. The project shows high contributor engagement but also a significant maintenance backlog: many high-severity security and stability issues (rated “platinum hermit” or “diamond lobster”) are stalled awaiting maintainer review or product decisions. The PR pipeline is healthy, with numerous small fixes and feature additions merged, but several larger PRs are blocked on proof-of-behaviour requests. Overall, the community is actively building and reporting problems, while maintainer bandwidth appears to be a bottleneck for moving critical items forward.

---

## Releases

No new releases today.

---

## Project Progress – Merged/Closed PRs Today

Key pull requests that were merged or closed today (selected from top-commented PRs):

- **`#94366`** (closed) – `fix(cron): allow empty assistant replies in cron lane` – Prevents `FailoverError` on intentionally silent cron watchers.
- **`#94397`** (closed) – `fix(agents): render identity name in runtime prompt` – Resolves `#8126` by injecting agent identity into system prompts.
- **`#94398`** (closed) – `fix(gateway): expose idempotencyKey in chat history metadata` – Improves deduplication in web UI.
- **`#94063`** (closed) – `fix(doctor): drop inert legacy cron notify when cron.webhook is unset` – Removes noisy warning in `doctor` checks.
- **`#92682`** (closed) – `fix(read): use system encoding fallback for non-UTF-8 text files on Windows` – Important for users on Chinese Windows (GBK encoding).
- **`#94209`** (closed) – `fix(model): cap contextWindow at native runtime catalog limit` – Prevents configuration mismatch (e.g., Codex gpt-5.5).
- **`#94205`** (closed) – `fix(ui): ensure raw config textarea is visible and scrollable` – UI glitch fix.
- **`#94196`** (closed) – `fix(auth): auto-correct type:token to type:api_key for Anthropic Enterprise` – Fixes 529 errors.
- **`#94200`** (closed) – `fix(docker): add optional Claude CLI install and persistent ~/.claude directory` – Improves Docker experience.
- **`#94234`** (closed) – `fix(anthropic): allow failover for thinking signature replay errors` – Fixes long-turn sessions with stale signatures.
- **`#94277`** (closed) – `fix(exec): add advisory file lock to prevent concurrent approval writes from losing entries` – Prevents race condition in `exec-approvals.json`.

Other open PRs that advanced today include parallelised MCP connections (`#94382`), direct text delivery for subagent completion (`#94375`), and IPv4 loopback fix for port checks (`#94399`).

---

## Community Hot Topics

Most active issues and PRs by comment count (with links):

- **`#25592`** (32 comments) – [Text between tool calls leaks to messaging channels](https://github.com/openclaw/openclaw/issues/25592)  
  *P1 security – could expose internal outputs to Slack, iMessage, etc. This is a long-running pain point (opened Feb 2026) with high community concern.*

- **`#9443`** (25 comments) – [Request: Prebuilt Android APK releases](https://github.com/openclaw/openclaw/issues/9443)  
  *Strong demand for official APK builds; no maintainer decision yet.*

- **`#68596`** (15 comments) – [Configurable streaming watchdog timeout threshold](https://github.com/openclaw/openclaw/issues/68596)  
  *Users running reasoning models (Kimi, DeepSeek) hit frequent watchdog resets; 8 👍.*

- **`#7707`** (12 comments) – [Memory Trust Tagging by Source](https://github.com/openclaw/openclaw/issues/7707)  
  *Security feature to prevent memory poisoning; 12 comments but no maintainer traction.*

- **`#31583`** (12 comments) – [`exec` tool does not inherit `skills.entries.*.env` environment variables](https://github.com/openclaw/openclaw/issues/31583)  
  *Regression affecting secret injection; linked PR open but stalled.*

- **`#27445`** (11 comments) – [`announceTarget` option for sub-agent completion announce routing](https://github.com/openclaw/openclaw/issues/27445)  
  *Feature to let parent agent orchestrate multi-step workflows; 5 👍.*

- **`#86215`** (9 comments) – [Codex OAuth refresh failures can wedge an agent for hours](https://github.com/openclaw/openclaw/issues/86215)  
  *P1 platinum – OAuth timeout causes hours of silent failure; no aggressive profile rotation.*

- **`#43747`** (9 comments) – [Memory management is in chaos](https://github.com/openclaw/openclaw/issues/43747)  
  *Three users report inconsistent memory storage behaviour – regression.*

**Underlying needs:** Security (text leakage, memory poisoning, OAuth, env injection) remains the dominant theme. Users also want better control over tool outputs, streaming behaviour, and memory semantics.

---

## Bugs & Stability

### High Severity (P1 / Platinum Hermit / Diamond Lobster)

| Issue | Description | Status | Fix PR? |
|-------|-------------|--------|---------|
| [#25592](https://github.com/openclaw/openclaw/issues/25592) | Text between tool calls leaks to messaging channels | Open, needs maintainer review | None |
| [#86215](https://github.com/openclaw/openclaw/issues/86215) | Codex OAuth refresh failures wedge agents for hours | Open, needs live repro | None |
| [#45224](https://github.com/openclaw/openclaw/issues/45224) | Unhandled Playwright assertion error crashes entire Gateway | Open, needs maintainer review | None |
| [#44502](https://github.com/openclaw/openclaw/issues/44502) | Discord routing / mention-gating regression | Open, linked PR open | None listed |
| [#43374](https://github.com/openclaw/openclaw/issues/43374) | All LLM API calls time out simultaneously (concurrency bug) | Open, needs info | None |
| [#40255](https://github.com/openclaw/openclaw/issues/40255) | Regression: user-configured heartbeat prompt no longer respected | Open, linked PR open | None |
| [#38091](https://github.com/openclaw/openclaw/issues/38091) | WebSocket reconnect terminates sessions (TUI/Web) | Open, needs maintainer review | None |
| [#31583](https://github.com/openclaw/openclaw/issues/31583) | `exec` tool missing `skills.entries.*.env` variables | Open, linked PR open | None |
| [#92451](https://github.com/openclaw/openclaw/issues/92451) | v2026.6.x system prompt bloat degrades smaller models | Open, needs live repro | None |

### Bugs Reported Today (2026-06-18)

- **`#94309`** (closed) – [Telegram Desktop does not offer Quote & Reply on OpenClaw bot messages](https://github.com/openclaw/openclaw/issues/94309) – *Minor UX, closed with note? Actually open status is "CLOSED"? It's closed, but mentioned as issue. Low impact.*
- **`#94360`** (open) – [Feishu: exec stderr output sent as user-visible reply](https://github.com/openclaw/openclaw/issues/94360) – *Medium: stderr blocks agent's actual response.*
- **`#94032`** (open) – [exec private-LAN access fails while same user GUI succeeds](https://github.com/openclaw/openclaw/issues/94032) – *P2, platinum – puzzling network isolation issue.*
- **`#93794`** (open) – [Messages on v2026.6.8 no longer supported on Telegram Web](https://github.com/openclaw/openclaw/issues/93794) – *P1 regression, 8 👍, needs immediate investigation.*
- **`#93884`** (open) – [Document gateway host agent runtime boundary](https://github.com/openclaw/openclaw/issues/93884) – *Documentation decision capture.*

### Regressions (noted in title)

- `#31583`, `#44502`, `#40255`, `#43747`, `#92451`, `#93794`

---

## Feature Requests & Roadmap Signals

Top community-requested features and likely next-version candidates:

| Issue | Feature | Votes/Activity | Predict Next Version? |
|-------|---------|----------------|-----------------------|
| [#9443](https://github.com/openclaw/openclaw/issues/9443) | Prebuilt Android APK releases | 25 comments, 2 👍 | Likely soon (ongoing pressure) |
| [#68596](https://github.com/openclaw/openclaw/issues/68596) | Configurable streaming watchdog timeout | 15 comments, 8 👍 | Possible – high demand |
| [#7707](https://github.com/openclaw/openclaw/issues/7707) | Memory Trust Tagging by Source | 12 comments, 0 👍 | Lower priority (no 👍) |
| [#27445](https://github.com/openclaw/openclaw/issues/27445) | Sub-agent completion announce routing option | 11 comments, 5 👍 | Likely – dev interest |
| [#13700](https://github.com/openclaw/openclaw/issues/13700) | Session snapshots (save/load checkpoints) | 6 comments, 0 👍 | Possible |
| [#28300](https://github.com/openclaw/openclaw/issues/28300) | Theme customization system (control UI) | 6 comments, 5 👍 | Likely – strong UX demand |
| [#13219](https://github.com/openclaw/openclaw/issues/13219) | Per-model cost tracking | 5 comments, 1 👍 | Possible – admin need |
| [#8892](https://github.com/openclaw/openclaw/issues/8892) | `--agent` flag for TUI | 4 comments, 3 👍 | Likely – straightforward |
| [#7456](https://github.com/openclaw/openclaw/issues/7456) | Go back navigation in onboarding wizard | 4 comments, 0 👍 | Minor improvement |
| [#11955](https://github.com/openclaw/openclaw/issues/11955) | Memory/Context improvements (metrics, search, chaining) | 4 comments, 0 👍 | Larger effort – possible |
| [#14785](https://github.com/openclaw/openclaw/issues/14785) | Reduce tool schema token overhead (~3,500 tok/session) | 7 comments, 0 👍 | High impact – maintainer likely interested |
| [#16670](https://github.com/openclaw/openclaw/issues/16670) | Onboarding wizard: include memory/embedding setup | 8 comments, 1 👍 | Likely – reduces confusion |
| [#7443](https://github.com/openclaw/openclaw/issues/7443) (likely #7433) | WhatsApp group message reliability | 4 comments, 1 👍 | Important for WhatsApp users |

**Signal:** There is strong demand for better **memory management** and **observability** (cost tracking, metrics, logging). Many feature requests relate to **multi-agent orchestration** (sub-agent routing, session snapshots). The maintainers should prioritise the streaming watchdog and memory trust tagging given security implications, but user feedback suggests they are not yet acted upon.

---

## User Feedback Summary

Real pain points expressed by users in today’s activity:

- **“Text between tool calls leaks to channels”** – hurts trust, especially in collaborative spaces.  
- **“Memory management is in chaos”** – different roles, different stores, no unified behaviour.  
- **“Streaming watchdog resets every 30s during reasoning”** – prevents use with extended-think models.  
- **“Codex OAuth can wedge for hours without alerting”** – operations nightmare.  
- **“WebSocket reconnect kills active sessions”** – frequent with poor connections.  
- **“exec tool doesn’t inherit env vars from skills config”** – breaks secret injection.  
- **“v2026.6.x system prompt bloat degrades smaller models”** – regression affecting edge devices.  
- **“Telegram Web broke on v2026.6.8”** – high impact (8 👍, P1).  
- **“Feishu exec stderr covers agent responses”** – medium annoyance.  
- **“exec private-LAN access fails”** – confusing for local deployments.  

**Satisfaction:** Users appreciate the rapid PR churn but are frustrated by slow resolution of critical regressions and security issues.

---

## Backlog Watch

These important issues and PRs have gone long-unanswered or lack maintainer traction:

### Critical/High Priority Stalled Issues

- **[#25592](https://github.com/openclaw/openclaw/issues/25592)** – P1 security (text leakage) – open since Feb 2026, needs product decision & security review.  
- **[#86215](https://github.com/openclaw/openclaw/issues/86215)** – P1 OAuth wedge – open since May 2026, needs live repro & decision.  
- **[#45224](https://github.com/openclaw/openclaw/issues/45224)** – P1 Gateway crash from Playwright – open since March 2026, needs maintainer review.  
- **[#43374](https://github.com/openclaw/openclaw/issues/43374)** – P1 all-timeout concurrency bug – open since March 2026, needs info.  
- **[#40255](https://github.com/openclaw/openclaw/issues/40255)** – P1 heartbeat prompt regression – linked PR open, but product decision missing.  
- **[#38091](https://github.com/openclaw/openclaw/issues/38091)** – P1 WebSocket reconnect termination – open since March 2026.  
- **[#31583](https://github.com/openclaw/openclaw/issues/31583)** – P1 `exec` env regression – linked PR open, but product decision missing.  
- **[#92451](https://github.com/openclaw/openclaw/issues/92451)** – P2 system prompt bloat degradation – open since June 12, needs live repro.  
- **[#43747](https://github.com/openclaw/openclaw/issues/43747)** – P2 memory management chaos – open since March 2026, source repro provided.

### PRs Needing Maintainer Action

- **[#94351](https://github.com/openclaw/openclaw/pull/94351)** – fix: rewrite asset hrefs for reverse proxy – status: “waiting on author”.  
- **[#94057](https://github.com/openclaw/openclaw/pull/94057)** – fix sandbox skills mount in Docker – needs real behaviour proof.  
- **[#94362](https://github.com/openclaw/openclaw/pull/94362)** – fix ollama thinking profile resolution – needs real behaviour proof.  
- **[#94107](https://github.com/openclaw/openclaw/pull/94107)** – fix Telegram reserved target rejection – needs proof, merge-risk: 🚨 message-delivery.  
- **[#44111](https://github.com/openclaw/openclaw/pull/44111)** – Backup encrypted snapshot flow – large, needs serious review.  
- **[#40311](https://github.com/openclaw/openclaw/pull/40311)** – feat: Brave Goggles support – waiting for behaviour proof.

**Recommendation:** The project would benefit from a focused maintainer “bug bash” on the P1 platinum/diamond issues and clearing the grooming queue (`clawsweeper:needs-product-decision`, `clawsweeper:needs-maintainer-review`). The high number of tagged but unresolved items suggests a need for dedicated triage cycles.

---

*This digest was generated from GitHub data retrieved on 2026-06-18 23:59 UTC. All links point to the [OpenClaw repository](https://github.com/openclaw/openclaw).*

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: Personal AI Assistant Open-Source Ecosystem
**Date:** 2026-06-18 | **Period:** Last 24 hours

---

## 1. Ecosystem Overview

The personal AI agent open-source landscape is experiencing a Cambrian explosion, with eight actively developed projects processing a combined **~900 issues and ~750 pull requests** in a single day. The ecosystem is bifurcating into two camps: heavyweight reference implementations (OpenClaw, ZeroClaw, CoPaw) that prioritize feature breadth and extensibility, and lean, purpose-built agents (NanoBot, NanoClaw, PicoClaw) optimized for specific deployment scenarios. A common thread across all projects is the tension between rapid feature iteration and security/stability debt—every project has critical vulnerabilities or regressions that remain unpatched, suggesting maintainer bandwidth is the primary bottleneck across the ecosystem. The healthiest projects by closure velocity (IronClaw, NanoBot) share a pattern of dedicated bug-triage cycles and community-contributed fix PRs arriving within hours of issue creation.

---

## 2. Activity Comparison

| Project | Issues Updated (Open/Closed) | PRs Updated (Open/Merged) | Releases (24h) | Health Score | Notes |
|---------|------------------------------|---------------------------|----------------|--------------|-------|
| **OpenClaw** | 200 (184/16) | 500 (414/86) | 0 | ⚠️ **Strained** | High volume but 80% of items stalled; P1 security bugs linger for months |
| **ZeroClaw** | 17 (16/1) | 50 (33/17) | 0 | 🟢 **Healthy** | 11 PRs merged; stacked feature work landing cleanly |
| **CoPaw** | 42 (32/10) | 50 (17/33) | 2 (v1.1.12 + beta) | 🟢 **Productive** | 76% closure rate on issues; 10 first-time contributor PRs |
| **Hermes Agent** | 10 (10/0) | 50 (46/4) | 0 | 🟡 **Moderate** | P1 OAuth fix underway but P2 memory bypass unaddressed |
| **NanoBot** | 6 (5/1) | 32 (14/18) | 0 | 🟢 **Efficient** | 18 PRs merged; rapid turnaround on bugs (avg <24h) |
| **IronClaw** | 11 (7/4) | 50 (33/17) | 0 | 🟢 **High Velocity** | 17 PRs merged; strong dogfooding culture catching bugs early |
| **LobsterAI** | 1 (1/0) | 11 (0/11) | 1 (2026.6.15) | 🟢 **Focused** | 11/11 PRs merged; security vuln reported today needs patch |
| **NanoClaw** | 5 (4/1) | 20 (17/3) | 2 (v2.1.17) | 🟢 **Responsive** | All reported bugs have fix PRs; breaking releases managed clearly |
| **PicoClaw** | 4 (4/0) | 10 (4/6) | 0 | 🟢 **Stable** | Critical SSRF fix merged same-day; low issue volume |
| **NullClaw** | 3 (3/0) | 2 (2/0) | 0 | 🟡 **Low Activity** | Only 2 PRs open; scheduler bug unaddressed for weeks |

**Health Score Legend:**
- 🟢 Healthy: Closure rate >50%, critical issues have active fix PRs, maintainers responsive
- 🟡 Moderate: Moderate activity but stalled high-severity items or low closure rate
- ⚠️ Strained: Significant maintenance debt, high issue-to-fix ratio, long-standing critical bugs

---

## 3. OpenClaw's Position

**Advantages vs. Peers:**
- **Scale & Breadth:** With 500 daily PRs and 200 issues, OpenClaw dwarfs all competitors. Its issue taxonomy (platinum hermit, diamond lobster) indicates a mature severity classification system absent in most peers.
- **Community Depth:** The most active feature requests (Android APK, streaming watchdog, memory trust tagging) have 8–25 comments each, reflecting a large, engaged user base. No other project approaches this level of participation.
- **Ecosystem Gravity:** Multiple projects (LobsterAI, CoPaw) explicitly depend on OpenClaw components (gateway, provider SDK), making it a de facto infrastructure layer for the ecosystem.

**Technical Approach Differences:**
- OpenClaw employs a **microservice-style architecture** (separate gateway, agent runtime, doctor, cron subsystems), whereas IronClaw and Hermes Agent favor monolithic Rust/Go binaries. This gives OpenClaw better horizontal scaling but at the cost of operational complexity.
- Its **cron lane and sub-agent orchestration** model is the most sophisticated among peers—no other project has equivalent multi-agent workflow primitives.

**Community Size Comparison:**
- Issue/PR volume is **5–10x higher** than the next busiest projects (CoPaw, ZeroClaw, IronClaw at ~50 daily PRs each).
- However, **closure rate is the worst** among active projects (43 PRs merged vs 414 open = 9% throughput ratio vs NanoBot's 56%).
- Maintainer bandwidth appears to be a structural constraint—critical items marked `needs-product-decision` have been untouched for 2–4 months, a pattern not observed in smaller projects like NanoClaw or PicoClaw.

---

## 4. Shared Technical Focus Areas

### Security & Data Protection
- **Text/Tool Output Leakage:** OpenClaw (#25592), CoPaw (#5204), Hermes Agent (#48181) all have bugs where internal agent outputs leak to messaging channels or cross-agent boundaries.
- **OAuth/Auth Stability:** OpenClaw (#86215), Hermes Agent (#48176), NanoClaw (#2794) all experienced OAuth authentication failures that completely block paid-tier users.
- **File System SSRF:** PicoClaw fixed a critical OneBot SSRF (#3140); NanoClaw has an open path traversal (#2800); LobsterAI has an arbitrary file read vulnerability (#2176).

### Streaming & Model Compatibility
- **Streaming Watchdog / Timeout:** OpenClaw (#68596), CoPaw (#5218), IronClaw (no-progress detection stack #5022) are all building or tuning streaming resilience mechanisms.
- **Model-Specific Schema Issues:** PicoClaw fixed Gemini 3.5 Flash tool execution (#3136); OpenClaw capped contextWindow for Codex gpt-5.5 (#94209); NanoBot added Mistral reasoning_effort validation (#4351).

### Memory & Context Management
- **Memory Chaos / Inconsistency:** OpenClaw (#43747), CoPaw (#5218), and Hermes Agent (#48181) all have open issues about memory behavior varying unpredictably across stores.
- **Context Compression & Prompt Bloat:** OpenClaw (#92451), IronClaw (#5022), LobsterAI (#2145) are actively working on context compaction algorithms.

### Multi-Agent & Collaboration
- **Cross-Agent Loops:** CoPaw (#5204), OpenClaw (#25592) report infinite loops or message leakage between agents—no circuit-breaker pattern exists yet.
- **Sub-Agent Routing & Approval:** OpenClaw (#27445), Hermes Agent (#12794), NanoClaw (#2793) are all building directed approval gates and delegation controls.

### Platform Parity
- **Windows Gaps:** ZeroClaw (#7462—74 test failures), OpenClaw (#92682—GBK encoding fix) highlight persistent Windows ecosystem issues.
- **Android/Mobile:** OpenClaw (#9443—Android APK), ZeroClaw (#7911—Termux) show unmet demand for ARM mobile targets.

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | ZeroClaw | CoPaw | IronClaw | NanoBot/NanoClaw | PicoClaw |
|-----------|----------|----------|-------|----------|------------------|----------|
| **Primary Language** | TypeScript/Python | Rust | TypeScript | Rust | TypeScript | Go |
| **Target User** | Power users, self-hosters | Enterprise/security-focused | Chinese-market, cloud-native | Researchers, agent developers | Lightweight, managed-fleet ops | Embedded/edge devices |
| **Architecture** | Microservice (gateway + agents) | Monolithic binary | Monolithic + Web console | Monolithic Rust binary | Modular Node.js + CLI | Single binary + plugins |
| **Unique Feature** | Cron lanes, sub-agent orchestration | WASM plugin system, SBOM/build provenance | Built-in skills marketplace | Reborn WebUI, output-aware progress detection | Managed-fleet upgrade tripwire | NEAR AI TEE provider, DeltaChat gateway |
| **Release Cadence** | Daily (patch-level); no stable tag | Milestone-tracked (v0.8.x) | Weekly (v1.1.x stable) | Continuous; no version bumps | Breaking changes in rollup releases | Sporadic; fix-driven |
| **Security Posture** | High severity awareness but slow patching (4-month gaps) | Proactive: SSRF pinning, CI hardening RFC | Reactive: OOM fix, ChromaDB probe | Reactive but fast: hotfix within hours | Proactive: security fixes auto-PR'd with issues | Quick on critical (SSRF fixed in <24h) |
| **Community Contribution** | Large but contributor bottleneck | Healthy; Windows & Android contributions | Strong; 10 first-time contributors/day | Active dogfooders; few external contributors | Moderate; CLI-focused contributions | Low; PRs from 2-3 consistent contributors |

**Key Architectural Divergence:**
- **WASM-based plugins** (ZeroClaw) vs **NPM package ecosystem** (OpenClaw, NanoBot) vs **Native Rust/TUI** (IronClaw). ZeroClaw's WASM plugin program (#7314, #7822) is the only project explicitly moving toward sandboxed third-party extensions.
- **Managed-fleet focus** (NanoClaw) vs **self-hosted power user** (OpenClaw, CoPaw) vs **desktop thin-client** (Hermes Agent #38602). These represent fundamentally different deployment philosophies affecting packaging, update mechanics, and auth models.

---

## 6. Community Momentum & Maturity

### Tier 1: Rapid Iteration (PR merger rate >40% of open+updated)
- **IronClaw** — 17 PRs merged/50 updated (34%), with a dogfooding pipeline catching bugs pre-release. The no-progress detection stack (3 PRs) merging within hours suggests strong review velocity.
- **NanoBot** — 18 PRs merged/32 updated (56%), the highest throughput ratio. The team is closing issues as fast as they open, indicating disciplined triage.
- **NanoClaw** — 3 PRs merged/20 updated (15%) but all critical bugs fixed same-day. Breaking releases are clearly communicated with migration paths.

### Tier 2: High-Volume, Bottlenecked
- **OpenClaw** — 86 PRs merged/500 updated (17%). Despite enormous absolute output, the open/closed ratio reveals structural inefficiency. The 4-month-old security issues indicate the team lacks bandwidth for deep architectural fixes.
- **CoPaw** — 33 PRs merged/50 updated (66% issue closure). The 10 first-time contributor PRs signal healthy community growth, but three critical process-freeze bugs remain unresolved.
- **ZeroClaw** — 11 PRs merged/50 updated (22%). Stacked PRs and milestone trackers demonstrate good planning, but the Windows/Android gaps and `needs-author-action` backlog slow momentum.

### Tier 3: Niche Maintenance
- **PicoClaw** — Low volume but high responsiveness. The project is stable and focused, likely targeting a narrow deployment surface (embedded, NEAR ecosystem).
- **NullClaw** — Dormant: only 2 PRs, scheduler bug open since May. This project may be winding down or maintainer-constrained.
- **LobsterAI** — Burst activity: 11 PRs merged in one day after a major release. Post-release stabilization mode; security vulnerability (#2176) needs urgent attention.

### Maturity Indicators
- **Release discipline:** Only OpenClaw lacks a clear release strategy (no tags, no changelog). All other projects have semver or milestone-tracked releases.
- **Test infrastructure:** IronClaw, ZeroClaw, and CoPaw have CI/CD pipelines. ZeroClaw's missing Windows CI (#7462) is a known gap being actively discussed.
- **Security audit cadence:** No project mentions external security audits or bug bounty programs. Security fixes are reactive (reported by users) rather than proactive.

---

## 7. Trend Signals for AI Agent Developers

### 1. The Post-Chat Paradigm: Agents as Infrastructure
The most active features across projects are not chat improvements but **scheduled workflows** (OpenClaw cron, ZeroClaw cron CLI, NanoBot cron), **agent-to-agent communication** (CoPaw Matrix, OpenClaw sub-agent routing), and **automation triggers** (IronClaw dogfooding tracker, NanoClaw messaging-groups). The community is treating agents as infrastructure—scheduled, observable, and composable—rather than conversational interfaces.

**Value for developers:** Invest in agent orchestration patterns (directed acyclic graphs, approval gates, circuit breakers) over chat UX. The next wave of value will come from autonomous multi-agent workflows, not single-turn conversations.

### 2. Security Debt is the Dominant Risk
Every project in this sample has at least one **unpatched vulnerability that could exfiltrate data or crash the system**. The most dangerous pattern is **output channel leakage** (OpenClaw #25592, CoPaw #5204)—agents that connect to Slack, Telegram, or email can expose internal agent reasoning or tool outputs to unintended audiences. SSRF and arbitrary file read vulnerabilities (PicoClaw #3070, LobsterAI #2176) are the second most common class.

**Value for developers:** Prioritize input validation pipelines and output channel isolation. Treat every agent as a potential data breach vector until proven otherwise. ZeroClaw's approach (pinning DNS for http_request, typed delete-with-cascade) is a pattern to emulate.

### 3. The Rise of Local-First, Thin-Client Architectures
Hermes Agent's most-upvoted issue (#38602—desktop thin-client), IronClaw's Reborn WebUI, and NanoClaw's managed-fleet agent model all point to a **separation of runtime and UI**. Developers want lightweight frontends (CLI, TUI, WebUI) that connect to remote or containerized agent runtimes rather than monolithic local installations.

**Value for developers:** Design your agent runtime as a headless service with a well-defined API. Multiple frontends (CLI, Web, mobile) can connect to the same backend. This pattern also enables the microservice scaling that OpenClaw is pioneering.

### 4. Multi-Agent Orchestration is Becoming a Commodity
Sub-agent delegation, context compaction, and memory consolidation are now being implemented by **every project simultaneously**. This is the fast-follower phase of the ecosystem—the first project to ship a reliable, secure multi-agent orchestration layer will capture significant mindshare.

**Value for developers:** Watch OpenClaw's sub-agent routing (#27445) and ZeroClaw's WASM hooks (#7822) as bellwethers. The battle will be won by the project that solves the **infinite loop problem** (CoPaw #5204, OpenClaw #25592) first.

### 5. Platform Lock-in is Breaking Down
Multiple projects are building **migration tools** (CoPaw `qwenpaw migrate openclaw` PR #5276) and **provider-agnostic gateways** (LLM provider swaps in NanoBot #4351, PicoClaw #2917). Users are treating agents as interchangeable infrastructure, not platform-specific products.

**Value for developers:** Invest in open formats (OpenAI-compatible APIs, common skill definitions) rather than proprietary protocols. The ecosystem is consolidating around composable, swappable components—not walled gardens.

### 6. Unmet Demand: Reliable Context Memory
Despite every project working on memory management, **no project has a solution for predictable, consistent memory behavior**. OpenClaw (#43747), CoPaw (#5218), and Hermes Agent (#48181) all have open "memory chaos" issues. This is the single largest unsolved technical challenge in the ecosystem.

**Value for developers:** This is a greenfield opportunity. A clean, auditable, source-tagged memory system (like OpenClaw's proposed trust tagging in #7707) would be a differentiating feature. Current approaches (vector stores, key-value histories, recursive summarization) all have compounding error rates that make long-running agents unreliable.

---

## Summary

The ecosystem is healthy and accelerating, but quality and security gaps are widening as feature velocity outpaces maintenance capacity. OpenClaw remains the most influential project by scale and ecosystem dependency, but its maintenance debt makes it a risky foundation for production deployments without dedicated internal support. Smaller projects (NanoBot, IronClaw, NanoClaw) demonstrate that **rapid, responsive development is possible** with disciplined triage and a focused scope. For technical decision-makers, the choice between projects should hinge on: (1) your tolerance for unpatched security issues, (2) your need for multi-agent orchestration vs. single-agent reliability, and (3) your willingness to invest in migration paths as the ecosystem consolidates.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest – 2026-06-18

## 1. Today’s Overview

NanoBot saw a surge of activity over the past 24 hours, with **18 merged/closed pull requests** and **6 open issues** updated (1 closed). This high PR velocity reflects a strong focus on bug fixes, provider enhancements, and infrastructure hardening. No new releases were published, but the volume of merged work — especially in areas like session replay, tool security, and channel integrations — suggests a release may be imminent. Community engagement remains moderate, with several feature requests and enhancement proposals receiving fresh attention.

## 2. Releases

*No new releases have been published during this period.*

## 3. Project Progress

**18 pull requests were merged or closed** in the last day, representing significant advances across multiple areas:

- **Web & Search** – Added [Keenable as a built-in search provider](https://github.com/HKUDS/nanobot/pull/4350) and [better Mistral support](https://github.com/HKUDS/nanobot/pull/4351) (strict parameter handling, reasoning_effort validation).
- **Messaging Channels** – Merged [WhatsApp read receipts (blue ticks)](https://github.com/HKUDS/nanobot/pull/4354) and [Feishu streaming update recovery](https://github.com/HKUDS/nanobot/pull/4381). A [Feishu QR scan-to-create CLI command](https://github.com/HKUDS/nanobot/pull/4391) is still open.
- **Session & Memory** – Fixed [replay-window history trimming](https://github.com/HKUDS/nanobot/pull/4349) to avoid cutting user turns, and [preserved delivery context during memory consolidation](https://github.com/HKUDS/nanobot/pull/4373).
- **Tool Security** – Merged [read-only protection for extra allowed directories](https://github.com/HKUDS/nanobot/pull/4053) and [filesystem workspace write policy clarification](https://github.com/HKUDS/nanobot/pull/4202). Allowed [git commands from workspace subdirectories](https://github.com/HKUDS/nanobot/pull/4380).
- **Stability & Logging** – Fixed [proxy interference with local endpoints](https://github.com/HKUDS/nanobot/pull/4367), [Anthropic tool ID sanitization](https://github.com/HKUDS/nanobot/pull/4356), and added [primary model error logging before fallback](https://github.com/HKUDS/nanobot/pull/4385).
- **WebUI** – Corrected [activity duration display](https://github.com/HKUDS/nanobot/pull/4283) and fixed [tool model preset switching](https://github.com/HKUDS/nanobot/pull/4347).
- **MCP** – A [fix for generator tracking during server close](https://github.com/HKUDS/nanobot/pull/4303) is still open.

## 4. Community Hot Topics

- **Multi-tenant Gateway** ([#936](https://github.com/HKUDS/nanobot/issues/936)) – This Feb 2026 feature request was updated again, highlighting ongoing demand for a single gateway serving multiple agents. The proposal includes a central config with an agent list – a common pain point among power users.
- **User‑friendly wizard** ([#4376](https://github.com/HKUDS/nanobot/issues/4376)) – Received one 👍, suggesting many users find the `nanobot onboard --wizard` too technical. The request asks for a guided, non‑technical setup flow.
- **Per‑model contextWindowTokens** ([#4389](https://github.com/HKUDS/nanobot/issues/4389)) – A recent request to set context windows per fallback model, important for mixed‑capability model stacks.
- **Multi‑instances for normies** ([#4390](https://github.com/HKUDS/nanobot/issues/4390)) – Asks for simpler multi‑instance management, e.g., folder‑based configs with hidden UI settings.
- **Feishu WebSocket card fix** ([#4342](https://github.com/HKUDS/nanobot/pull/4342)) – While comment count is not shown, the PR addresses a critical rendering bug for Feishu cards received via WebSocket.

## 5. Bugs & Stability

**BUGS REPORTED TODAY (high severity):**
- **[iOS Safari input zoom](https://github.com/HKUDS/nanobot/issues/4388)** – On iPhone Air (iOS 26.5), tapping the WebUI input box causes automatic page zoom and UI distortion. **Status:** Open, no fix PR yet.
- **[NameError: session_key](https://github.com/HKUDS/nanobot/issues/4322)** – A crash after merging `fix/prompt-caching`. **Status:** Closed – root cause identified in extracted method `_build_memory_context`.

**BUG FIXES MERGED TODAY (notable):**
- **Feishu streaming recovery** (#4381) – Retries failed card updates by reopening stream mode.
- **Git subdirectory guard** (#4380) – Shell safety now correctly compares paths against the command’s working directory.
- **Proxy blocking local endpoints** (#4367) – Disables proxy for localhost/LAN addresses; respects env proxy for cloud.
- **Anthropic tool ID sanitization** (#4356) – Rejects invalid characters in tool IDs to avoid 400 errors.
- **Replay‑window history truncation** (#4349) – Prevents LLM replay from starting mid‑user‑turn.
- **Filesystem write policy** (#4202) – Clarifies read‑only vs write‑allowed directories.

**Stability posture:** Extensive bug‑fix PRs indicate a project in active maintenance. The iOS zoom bug is the only reported regression without an immediate fix.

## 6. Feature Requests & Roadmap Signals

**Demanded features with strong signals:**
- **Multi‑tenant gateway** (#936) – Centralised gateway config would reduce resource overhead. Could land in next minor release if maintainers pick it up.
- **User‑friendly wizard** (#4376) – Aligns with NanoBot’s goal of lowering the barrier for new users. Likely to be prioritised.
- **Per‑model context window tokens** (#4389) – A targeted improvement for fallback reliability; relatively simple to implement.
- **On‑demand heartbeat trigger** (#3437) – Debugging tool for `HEARTBEAT.md` iteration. Low complexity, high developer value.
- **Multi‑instances for normies** (#4390) – Echoes #936 but from a usage/setup perspective. Could be folded into a future “profiles” feature.

**Merged features pointing to the next version:**
- Keenable search provider (#4350)
- WhatsApp read receipts (#4354)
- Mistral provider improvements (#4351)
- Feishu QR‑scan bot creation (#4391 – still open, but likely merged soon)

Taken together, the upcoming release seems focused on **channel parity** (Feishu, WhatsApp), **search diversity**, and **provider resilience**.

## 7. User Feedback Summary

**Pain points expressed:**
- “Onboarding is too technical” – #4376
- “Running multiple agents requires too many gateway containers” – #936
- “I want multiple instances with simple folder‑based configs” – #4390
- “WebUI zooms on iOS Safari and breaks layout” – #4388
- “Fallback model with smaller context window doesn’t trim prompts” – #4389

**Satisfaction indicators:**
- Rapid bug‑fix turnaround (many PRs merged within 24h of issue creation).
- The community is contributing actively (32 PRs updated, 14 open) – a sign of a healthy ecosystem.
- No widespread complaints about stability or missing core features.

## 8. Backlog Watch

Items needing maintainer attention:
- **[#936 – Multi‑tenant Gateway](https://github.com/HKUDS/nanobot/issues/936)** – Opened 2026-02-21, still open with one comment. This is a recurring request that should be evaluated for roadmap inclusion.
- **[#3437 – On‑demand heartbeat trigger](https://github.com/HKUDS/nanobot/issues/3437)** – Opened 2026-04-25, no maintainer response. A small but valuable debug feature.
- **[#4303 – MCP generator GC crash](https://github.com/HKUDS/nanobot/pull/4303)** – Open PR fixing a runtime crash during server reconnect. May require review/merge soon.
- **[#4342 – Feishu WebSocket card content](https://github.com/HKUDS/nanobot/pull/4342)** – Open PR with structural fixes, no comments from maintainers yet.
- **[#4391 – Feishu QR scan‑to‑create](https://github.com/HKUDS/nanobot/pull/4391)** – New PR, large feature, needs review.

*Generating this digest from raw GitHub data represents a contribution to the NanoBot project’s transparency and community communication.*

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent – Project Digest for 2026-06-18

## 1. Today's Overview
Activity remains high: **10 issues** were updated in the last 24 hours (all still open), and **50 pull requests** saw activity—**46 open** and **4 merged/closed**. No new releases were cut today. The project is processing a steady flow of bug reports and feature contributions across multiple components (gateway, agent, CLI, TUI, tools). Community engagement is strong, with several long‑standing feature requests (e.g., Desktop thin‑client installation) still gathering reactions. The security and stability teams are active: a P1 OAuth‑billing regression and a P2 memory‑toolset bypass are both under investigation with corresponding fix PRs already opened.

## 2. Releases
No new releases were published today. The latest available version remains unchanged.

## 3. Project Progress
Four PRs were merged or closed today:

- **#43051** (merged) – `fix(approval): honor glob command allowlist entries` – Adds support for shell‑style glob patterns (e.g., `podman *`) in the `command_allowlist` configuration, improving approval‑based security for user‑defined commands.
  [PR #43051](https://github.com/NousResearch/hermes-agent/pull/43051)

- Three other closed PRs (not fully detailed in the top‑20 list) – likely minor fixes or documentation updates.

Notable open PRs that advanced significantly (still open) and represent major feature work:

- **#47740** – LUMEN binary protocol transport for MCP servers (32–80% wire compression, 34 tools across 3 reference servers).  
  [PR #47740](https://github.com/NousResearch/hermes-agent/pull/47740)
- **#45929** – Clears npm advisories and scopes Nous Portal tags to prevent stale global state.  
  [PR #45929](https://github.com/NousResearch/hermes-agent/pull/45929)
- **#48184** – New OTLP observability plugin for OpenTelemetry export.  
  [PR #48184](https://github.com/NousResearch/hermes-agent/pull/48184)
- **#12794** – Per‑subagent model/provider overrides for `delegate_task` + model observability.  
  [PR #12794](https://github.com/NousResearch/hermes-agent/pull/12794)

## 4. Community Hot Topics

### Most Active Issue (by reactions)
**#38602 – Desktop Client‑Only Installation**  
18 👍, 6 comments. Users want to install the Hermes Desktop app as a thin client that connects to a remote Hermes runtime, instead of always bootstrapping the agent locally. This is a long‑standing pain point for network‑isolated or lightweight setups.  
[Issue #38602](https://github.com/NousResearch/hermes-agent/issues/38602)

### Most Active Issue (by comments)
**#48175 – Separate system prompts per Discord channel**  
2 comments, but filed only hours ago. The user wants to configure different AI personas per Discord channel using external `.md` files. This reflects a widespread need for multi‑persona support in chat integrations.  
[Issue #48175](https://github.com/NousResearch/hermes-agent/issues/48175)

### Most Active Pull Request (by comment count)
Among the top 20, none show explicit comment counts; however, the large feature PR **#47740 (LUMEN protocol)** and the security‑related **#48177 (OAuth billing fix)** are drawing attention from maintainers and reviewers.

### Analysis of Underlying Needs
- **Thin‑client mode** – Users want separation between the lightweight UI and the heavy runtime, enabling remote agent management.
- **Multi‑persona per channel** – Growing demand for flexible, context‑aware AI assistants across different conversation environments.
- **Robust MCP integration** – Active community contribution for binary protocol support shows appetite for efficient tool communication.

## 5. Bugs & Stability

| Priority | Issue | Description | Fix PR Exists? |
|----------|-------|-------------|----------------|
| **P1** | [#48176](https://github.com/NousResearch/hermes-agent/issues/48176) | OAuth Pro/Max/Team requests fail with HTTP 400 due to missing `x-anthropic-billing-header` system block. | Yes – [#48177](https://github.com/NousResearch/hermes-agent/pull/48177) (open, fix by same reporter) |
| **P2** | [#48181](https://github.com/NousResearch/hermes-agent/issues/48181) | Disabled `memory` toolsets can be bypassed by late memory‑provider tool injection – a security vulnerability. | No dedicated fix PR yet (reported today). |
| **P2** | [#48183](https://github.com/NousResearch/hermes-agent/issues/48183) | Desktop profile switching breaks session list visibility; cross‑profile data scattering. | No fix PR yet. |
| **P2** | [#48173](https://github.com/NousResearch/hermes-agent/issues/48173) | Mid‑session model switch leaves stale Model/Provider in persisted system prompt. | No fix PR yet. |
| **P2** | [#48172](https://github.com/NousResearch/hermes-agent/issues/48172) | macOS local browser tool fails to detect Chrome app bundle on first run. | Yes – [#48185](https://github.com/NousResearch/hermes-agent/pull/48185) (open) |
| **P3** | [#41808](https://github.com/NousResearch/hermes-agent/issues/41808) | Dashboard Chat tab React error #301 (max update depth) on external connections (Tailscale). | No fix PR yet. |
| **P3** | [#48182](https://github.com/NousResearch/hermes-agent/issues/48182) | TUI interactive experience lacks skill highlighting. | No fix PR yet (feature request). |
| **P3** | Multiple ComfyUI workflow bugs | Bundled workflows ship with `_comment` keys causing HTTP 500 on submit – three separate PRs addressing this. | Yes – [#48143](https://github.com/NousResearch/hermes-agent/pull/48143), [#48144](https://github.com/NousResearch/hermes-agent/pull/48144), [#48145](https://github.com/NousResearch/hermes-agent/pull/48145) (all open) |

**Critical**: The P1 OAuth regression blocks Claude Pro/Max/Team users entirely. The P2 memory bypass needs urgent security review.

## 6. Feature Requests & Roadmap Signals

### User‑Requested Features (new today)
- **#48182** – TUI skill highlighting on `/` keypress (like Codex/Claude).  
- **#48179** – Extend managed‑system concept beyond NixOS/Homebrew to Fedora/DNF and other package managers.  
- **#48163** – System tray support for Desktop (close to tray instead of quit) – already a PR.  
- **#48180** – Linux backend for computer‑use tool (currently macOS‑only).  

### Likely Candidates for Next Release
- The **OTLP observability plugin** (#48184) – completes the observability triad next to Langfuse and LangSmith.
- **LUMEN binary protocol** (#47740) – addresses performance and multi‑agent sharing for MCP.
- **MacOS Chrome detection fix** (#48185) – trivial user‑experience fix for first‑time browser tool usage.
- **Per‑subagent model overrides** (#12794) – highly requested for delegation workflows.

### Roadmap Signals
- **Packaging/distribution**: The Fedora packaging effort (#48179) indicates a push toward Linux distribution integration.
- **AI persona management**: The Discord multi‑prompt question (#48175) combined with the TUI enhancements suggests the team is investing in flexible user‑facing configuration.
- **Security hardening**: The P2 memory bypass (#48181) shows ongoing focus on tool‑policy enforcement.

## 7. User Feedback Summary

**Common Pain Points:**
- Desktop client always requires a full Hermes runtime, preventing lightweight/remote usage (#38602).
- OAuth authentication broken for high‑tier Anthropic accounts (P1, #48176) – immediate blocking for Pro/Max/Team subscribers.
- Profile switching in the Desktop app causes session data to become invisible or scattered (#48183).
- Local browser tool fails on macOS without manual environment variable configuration (#48172).
- ComfyUI workflows non‑functional out of the box due to `_comment` metadata (#48143 et al.).

**Positive Sentiment (inferred from feature requests):**
- Active community contributions (34 tools in #47740, new plugins like OTLP) indicate strong developer engagement.
- Users are migrating from Codex/Claude and requesting parity features (#48182), suggesting Hermes is seen as a viable alternative.

**Satisfaction Signals:**
- The high number of open PRs (50 updated today) reflects a healthy, responsive maintainer team.
- Issue #38602 has 18 👍 – a clear, long‑standing desire that the team has not yet addressed (last updated today, but created 2 weeks ago).

## 8. Backlog Watch

| Issue/PR | Age | Reason for Concern |
|----------|-----|-------------------|
| [#19331](https://github.com/NousResearch/hermes-agent/pull/19331) – `cognee_query` tool | Created 2026‑05‑03 (46 days) | Feature PR with no recent maintainer activity; may be blocked or abandoned. |
| [#12794](https://github.com/NousResearch/hermes-agent/pull/12794) – per‑subagent overrides | Created 2026‑04‑20 (59 days) | Large feature with many comments; no update from author in 2 months – risk of stale. |
| [#24923](https://github.com/NousResearch/hermes-agent/pull/24923) – treat timeout as refusal | Created 2026‑05‑13 (36 days) | Security‑related fix (P1) with no merge – unclear if design disagreement. |
| [#48173](https://github.com/NousResearch/hermes-agent/issues/48173) – stale model after switch | Reported today | No fix PR yet; P2 bug that could confuse users. |
| [#41808](https://github.com/NousResearch/hermes-agent/issues/41808) – Dashboard React error | Created 2026‑06‑08 (10 days) | Affects external network users; no maintainer response visible. |

**Call to Action:** Maintainers should review the two long‑standing feature PRs (#19331, #12794) and the security/timeout fix (#24923) to decide on next steps. The P2 memory bypass (#48181) needs rapid triage.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest — 2026-06-18

## Today's Overview
The project saw a flurry of activity with **10 pull requests updated** and **4 issues touched** in the last 24 hours. No new releases were published, but the development pulse is strong: **6 PRs were merged or closed**, including a critical security fix for OneBot’s media fetching vulnerability and a fix enabling tool execution with Gemini 3.5 Flash. Community attention is focused on a high-priority request to replace the unmaintained `libolm` with `vodozemac`. The merged NEAR AI Cloud provider and a new review feature signal continued expansion of platform support and user-facing functionality.

## Releases
*None — no new versions were released in this period.*

## Project Progress
The following PRs were merged or closed today (within the last 24h), representing concrete improvements:

- **Security fix (OneBot)** — [#3140](https://github.com/sipeed/picoclaw/pull/3140): Blocks attacker-controlled media URLs from fetching private/local addresses, closing [#3070](https://github.com/sipeed/picoclaw/issues/3070).
- **Gemini 3.5 Flash tool execution** — [#3136](https://github.com/sipeed/picoclaw/pull/3136): Adds the required `thought_signature` (snake_case) field to tool call requests, fixing [#3111](https://github.com/sipeed/picoclaw/issues/3111).
- **Sogou search parsing** — [#3139](https://github.com/sipeed/picoclaw/pull/3139): Updates HTML regex to match changed Sogou WAP page structure.
- **Session history display** — [#2990](https://github.com/sipeed/picoclaw/pull/2990): Fixes issue where only the last user message was shown in conversation view.
- **NEAR AI Cloud provider** — [#2917](https://github.com/sipeed/picoclaw/pull/2917): Adds a new OpenAI-compatible LLM provider with TEE-capable model suggestions.
- **Review feature (리뷰기능 추가)** — [#3138](https://github.com/sipeed/picoclaw/pull/3138): Introduces a review/feedback UI element (details in Korean).

Additionally, two open PRs are being worked on: diagnostic logging for Brave empty results ([#3141](https://github.com/sipeed/picoclaw/pull/3141)) and a spawn duplicate message fix ([#3142](https://github.com/sipeed/picoclaw/pull/3142)).

## Community Hot Topics
- **#3088 – Feature: Use vodozemac instead of libolm**  
  [Link](https://github.com/sipeed/picoclaw/issues/3088)  
  *2 thumbs up, 1 comment, label: `priority: high`*  
  This issue elicits the strongest community signal. Users demand replacing the insecure, unmaintained `libolm` with `vodozemac`, the official successor. The high-priority tag suggests the maintainers are actively considering it for the next release.

- **#3093 – Feature: SimpleX or Tox gateway**  
  [Link](https://github.com/sipeed/picoclaw/issues/3093)  
  *1 comment, 0 thumbs up, marked `stale`*  
  A user requests integration with SimpleX or Tox for decentralized messaging. Lack of engagement may indicate low priority or uncertainty about scope.

- **#3063 – PR: Add DeltaChat gateway**  
  [Link](https://github.com/sipeed/picoclaw/pull/3063)  
  *No comments, open since 2026-06-08*  
  An open pull request to add a DeltaChat gateway. This is likely the most concrete alternative being developed, though it has not been merged yet.

## Bugs & Stability
Several bugs were addressed or remain open; severity ranking is based on impact.

| Severity | Bug / Issue | Status | Fix PR |
|----------|-------------|--------|--------|
| **Critical** | OneBot inbound media fetch allows arbitrary server-side request; attacker can probe private networks. | Closed via [#3140](https://github.com/sipeed/picoclaw/pull/3140) | ✅ |
| **High** | Tool execution fails on Gemini 3.5 Flash due to missing `thought_signature` in schema → HTTP 400. | Closed via [#3136](https://github.com/sipeed/picoclaw/pull/3136) | ✅ |
| **Medium** | Sogou web search tool fails to parse any results (HTML structure changed). | Closed via [#3139](https://github.com/sipeed/picoclaw/pull/3139) | ✅ |
| **Medium** | `spawn` sub‑agent `ToolResult` sets both `ForUser` and `ForLLM`, causing duplicate message delivery. | Open PR [#3142](https://github.com/sipeed/picoclaw/pull/3142) | 🔧 in review |
| **Low** | `skills_install` discards type assertion `ok` for `version`/`force`; silent fallback to defaults. | Open PR [#3092](https://github.com/sipeed/picoclaw/pull/3092) | 🔧 stale |
| **Low** | Brave Web Search returns HTTP 200 with zero results; no diagnostic log. | Open PR [#3141](https://github.com/sipeed/picoclaw/pull/3141) adds logging | 🔧 in review |

No crashes or regressions were reported in the last 24 hours.

## Feature Requests & Roadmap Signals
The most notable feature signals point toward an upcoming **encryption stack upgrade** and **new messaging backends**:

- **vodozemac adoption ([#3088](https://github.com/sipeed/picoclaw/issues/3088))** — Marked `priority: high` and `help wanted`. Given the maintainer label, this is likely to land in the next minor or patch release.
- **DeltaChat gateway ([#3063](https://github.com/sipeed/picoclaw/pull/3063))** — An open PR that could be merged soon, adding a robust communication channel.
- **SimpleX / Tox ([#3093](https://github.com/sipeed/picoclaw/issues/3093))** — Only a suggestion so far, but echoes the desire for more privacy‑oriented gateways.
- **NEAR AI Cloud provider ([#2917](https://github.com/sipeed/picoclaw/pull/2917))** — Already merged; the next release will include it as a first‑class provider.

The merged **review feature ([#3138](https://github.com/sipeed/picoclaw/pull/3138))** suggests UI/UX maturity improvements are also on the roadmap.

## User Feedback Summary
Users are voicing concrete pain points:

- **Security anxiety** — The private‑network fetch vulnerability in OneBot ([#3070](https://github.com/sipeed/picoclaw/issues/3070)) was quickly fixed, but underscores demand for hardened input validation.
- **New model compatibility** — Adopters of Gemini 3.5 Flash were blocked from using tools entirely ([#3111](https://github.com/sipeed/picoclaw/issues/3111)). The rapid turnaround (issue → fix PR in 5 days) shows responsiveness.
- **Search reliability** — Sogou users hit a silent parsing failure ([#3139](https://github.com/sipeed/picoclaw/pull/3139)). Diagnostic logging for Brave ([#3141](https://github.com/sipeed/picoclaw/pull/3141)) proactively addresses similar blind spots.
- **Cryptographic trust** — The call to replace `libolm` ([#3088](https://github.com/sipeed/picoclaw/issues/3088)) reflects a broader dissatisfaction with using unmaintained dependencies.

Overall, satisfaction appears high when fixes arrive quickly, but the community expects proactive handling of dependency rot and third‑party API changes.

## Backlog Watch
Items that have been open for over a week without significant maintainer activity or resolution:

- **#3093 – Feature request: SimpleX or Tox**  
  Created 2026-06-10, no maintainer comment. Labeled `stale`. Low community engagement may be the reason, but a statement of feasibility would be helpful.
- **#3092 – PR: skills_install type assertion fix**  
  Created 2026-06-10, no updates since. Not merged nor commented on by maintainers. Although low‑severity, it’s a clean improvement that could be reviewed.
- **#3063 – PR: DeltaChat gateway**  
  Created 2026-06-08, no maintainer feedback. Since it adds significant functionality, it warrants timely review.

No other issues or PRs appear to be languishing without attention.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-18

## 1. Today's Overview

NanoClaw saw very high activity in the last 24 hours, with **20 pull requests** (17 open, 3 merged/closed) and **5 issues** (4 open, 1 closed) updated. Two rollup releases were published — v2.1.17 (latest) and v2.1.0 — both carrying **breaking changes** that affect SDK compatibility and startup flow. The project team responded quickly to critical bugs, including a delivery-stall regression affecting all agents, which was fixed and merged the same day. Several security fixes (path traversal, untrusted file reads) were also submitted.

## 2. Releases

Two new tags were issued:

- **[v2.1.17](https://github.com/nanocoai/nanoclaw/releases/tag/v2.1.17)** — Rollup release covering v2.1.1 through v2.1.17.  
  **Breaking change:** The `@onecli-sh/sdk` was bumped from **0.5.0 → 2.2.1**. This requires a OneCLI server that exposes the `/v1` API; older servers will return 404 on every SDK call. The sanctioned gateway and CLI versions are now pinned.  
  **Migration:** Ensure your OneCLI gateway is on a version supporting the `/v1` API.

- **[v2.1.0](https://github.com/nanocoai/nanoclaw/releases/tag/v2.1.0)** — Rollup covering v2.0.65 through v2.1.0.  
  **Breaking change:** The host now refuses to boot unless `data/upgrade-state.json` records that the install reached the current version through a sanctioned upgrade path.  
  **Migration:** Run `bash nanoclaw.sh` (or the appropriate upgrade script) to generate the required marker; managed fleets can set `NANOCLAW_DISABLE_UPGRADE_TRIPWIRE=1` as an opt-out (see PR #2780).

## 3. Project Progress — Merged/Closed PRs

Three PRs were merged or closed today:

- **[#2797](https://github.com/nanocoai/nanoclaw/pull/2797) `fix(delivery): isolate per-session failures`** — Merged. Fixes issue #2796 where a single bad session could stall message delivery for all agents. The delivery loop now catches per-session errors and continues.
- **[#2794](https://github.com/nanocoai/nanoclaw/pull/2794) `fix(providers): restore env-var gateway auth for managed fleets`** — Merged. Managed-fleet agents (immutable VM images) could no longer authenticate to the LLM after v2.1.17; this restores authentication via environment variables.
- **[#2780](https://github.com/nanocoai/nanoclaw/pull/2780) `feat(upgrade-state): env opt-out for the startup tripwire (managed fleets)`** — Merged. Adds `NANOCLAW_DISABLE_UPGRADE_TRIPWIRE=1` to allow managed fleets to bypass the new upgrade marker requirement.

## 4. Community Hot Topics

- **🔴 Critical delivery stall ([#2796](https://github.com/nanocoai/nanoclaw/issues/2796))** — A single unhealthy session could halt message delivery for every agent. Closed the same day with fix PR #2797. Generated the most discussion (1 comment) and was the top priority.
- **🔧 Agent-to-agent approval policies ([#2793](https://github.com/nanocoai/nanoclaw/pull/2793))** — A feature PR adding optional per-message approval gates on connected agents. This is a significant enhancement that could shape future agent interaction workflows.
- **🌐 Korean README translation ([#2806](https://github.com/nanocoai/nanoclaw/pull/2806))** — Community contribution adding full Korean translation, reflecting growing international interest.

## 5. Bugs & Stability

Ranked by severity:

| Severity | Issue | Summary | Fix PR exists? |
|----------|-------|---------|----------------|
| **Critical** | [#2796](https://github.com/nanocoai/nanoclaw/issues/2796) | One bad session stalls delivery for all agents | ✅ [#2797](https://github.com/nanocoai/nanoclaw/pull/2797) merged |
| **High** | [#2804](https://github.com/nanocoai/nanoclaw/issues/2804) | `ncl messaging-groups create` always throws NOT NULL constraint | ✅ [#2804](https://github.com/nanocoai/nanoclaw/pull/2804) open |
| **High** | [#2802](https://github.com/nanocoai/nanoclaw/issues/2802) | `ncl` socket client has no request timeout, unbounded buffer | ✅ [#2802](https://github.com/nanocoai/nanoclaw/pull/2802) open |
| **Medium** | [#2801](https://github.com/nanocoai/nanoclaw/issues/2801) | `safeParseContent` returns non-object JSON, callers get undefined | ✅ [#2801](https://github.com/nanocoai/nanoclaw/pull/2801) open |
| **Medium** | [#2800](https://github.com/nanocoai/nanoclaw/issues/2800) | Path traversal via `--folder` in `ncl groups create` (CWE-22) | ✅ [#2800](https://github.com/nanocoai/nanoclaw/pull/2800) open |
| **Medium** | [#2799](https://github.com/nanocoai/nanoclaw/issues/2799) | `send_file` can read arbitrary container files (CVE-2026-29611) | ✅ [#2799](https://github.com/nanocoai/nanoclaw/pull/2799) open |
| **Low** | [#2791](https://github.com/nanocoai/nanoclaw/issues/2791) | `add-imessage` step fails if `src/channels/` doesn't exist | ✅ [#2792](https://github.com/nanocoai/nanoclaw/pull/2792) open |
| **Low** | [#2789](https://github.com/nanocoai/nanoclaw/issues/2789) | `setup` skill only 600 bytes, no concrete steps | ✅ [#2790](https://github.com/nanocoai/nanoclaw/pull/2790) open |
| **Low** | [#2787](https://github.com/nanocoai/nanoclaw/issues/2787) | Port 10254 only mentioned in troubleshooting | ✅ [#2788](https://github.com/nanocoai/nanoclaw/pull/2788) open |
| **Low** | [#2785](https://github.com/nanocoai/nanoclaw/issues/2785) | `migrate-nanoclaw` skill has generic "Context" H1 | ✅ [#2786](https://github.com/nanocoai/nanoclaw/pull/2786) open |

All reported bugs have corresponding fix PRs, indicating strong maintainer responsiveness.

## 6. Feature Requests & Roadmap Signals

- **Per-message approval policies** ([#2793](https://github.com/nanocoai/nanoclaw/pull/2793)) — Adds directed approval gates on agent-to-agent connections. Likely to be included in the next minor release given its maturity.
- **CLI dashboard skill** ([#2795](https://github.com/nanocoai/nanoclaw/pull/2795)) — Read-only dashboard via CLI. A utility skill with no source changes, easy to merge.
- **Atlas Cloud LLM backend** ([#2717](https://github.com/nanocoai/nanoclaw/pull/2717)) — Documentation PR adding Atlas Cloud as an OpenAI-compatible option. Open since June 9; may be waiting for review or further validation.
- **Korean README** ([#2806](https://github.com/nanocoai/nanoclaw/pull/2806)) — Translation contribution, likely to be merged soon following existing language-switcher patterns.

Predictions: v2.2.0 will likely include the agent-to-agent approval feature and several of the security fixes currently in open PRs.

## 7. User Feedback Summary

- **Pain points** center on onboarding and documentation: setup skill too vague, missing directory errors, undisclosed port numbers, and generic headings. Users expressed this through issues #2789, #2791, #2787, #2785.
- **Reliability concerns** are high: the delivery-stall bug (#2796) would have caused silent message loss — its rapid fix (same day) should improve trust.
- **Managed-fleet users** needed workarounds for the startup tripwire and gateway auth; both were addressed within 24 hours (PRs #2780, #2794).

Overall sentiment appears constructive: users are reporting issues clearly, and the team is addressing them promptly.

## 8. Backlog Watch

- **[#2750](https://github.com/nanocoai/nanoclaw/pull/2750) `fix: recover stale outbound.db journals`** — Open since June 12, updated June 17. This PR addresses two related failure modes (#2516, #2640) that cause host–container communication issues. Despite its age, it has received recent updates and may be close to merge. No maintainer response is blocking it; the PR author is actively pushing.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-06-18

## Today’s Overview
The project maintains a low but steady level of activity with three open issues and two open pull requests updated in the last 24 hours; no items were closed or merged today. The two new PRs target long-standing community pain points (CLI key handling and memory recall configurability), signaling responsive maintenance despite no new releases. The three active issues span a confirmed bug (scheduler), a CLI usability problem, and a documentation/configuration question—none have been resolved yet, but the PRs directly address two of them. Overall project health appears stable, with maintainer attention focused on quality-of-life improvements and user-reported regressions.

## Releases
No new releases were published in the last 24 hours.

## Project Progress
No pull requests were merged or closed today. Two PRs remain open:
- **[#961 – feat(memory): add configurable auto-recall, recall_limit, max_context_bytes](https://github.com/nullclaw/nullclaw/pull/961)** (opened today by valonmulolli) introduces three new JSON config keys under `memory` to let users disable memory enrichment, cap the number of recalled entries, and limit context token usage. This addresses a feature gap long requested by power users who need finer control over memory overhead.
- **[#960 – fix(cli): handle arrow keys in agent REPL](https://github.com/nullclaw/nullclaw/pull/960)** (opened yesterday by vernonstinebaker) adds a small, allocation-free line editor for the interactive CLI, enabling proper up/down history navigation and cursor movement instead of printing control characters. This directly fixes issue #865.

## Community Hot Topics
- **[Issue #861 – How to enable the Web UI on headless VPS server?](https://github.com/nullclaw/nullclaw/issues/861)** (updated 2026-06-17, 1 comment) — A user asks for a plain‑language guide to setting up the Web UI on a remote server without a browser. Despite being open for nearly two months, it has only one comment, suggesting a gap in onboarding documentation that could frustrate new adopters.
- **[Issue #915 – [bug] Problem with scheduler unauthorized](https://github.com/nullclaw/nullclaw/issues/915)** (2 comments) — A user running NullClaw on Ubuntu with an external Ollama host reports the scheduler fails in both Telegram and CLI. The exact error is not described but points to an authentication/authorization issue. No maintainer response yet.
- **[Issue #865 – [bug] CLI shows ctrl characters for arrow keys](https://github.com/nullclaw/nullclaw/issues/865)** (2 comments) — Now linked to PR #960, this is the most actionable community concern and the PR author is likely the reporter or a contributor.

## Bugs & Stability
Two bugs have been reported and updated in the last 24 hours (both remain open):

| Issue | Description | Severity | Has Fix PR? |
|-------|-------------|----------|-------------|
| [#865](https://github.com/nullclaw/nullclaw/issues/865) | CLI control characters on arrow key input | **Medium** – usability regression for terminal users | Yes (PR #960) |
| [#915](https://github.com/nullclaw/nullclaw/issues/915) | Scheduler unauthorized (likely auth/token mismatch) | **High** – core scheduling feature broken for external LLM setups | No |

The scheduler bug (#915) is more severe because it blocks automated workflows. No regression reports from recent releases appear (no releases were made).

## Feature Requests & Roadmap Signals
The only feature-level change visible today is **PR #961** (memory configurability), which is likely to land in the next patch release. It reflects a common request: users want to disable memory recall to reduce latency or token cost, and to cap context size. This feature is straightforward to integrate and has no breaking changes (defaults preserve current behavior).

No other explicit feature requests are present in the latest issues, but issue #861 implicitly requests better documentation for headless Web UI setup—a documentation gap that could be addressed in a minor update.

## User Feedback Summary
- **Pain Points:** CLI responsiveness (#865), scheduler authentication (#915), and lack of clear Web UI setup instructions for headless servers (#861).
- **Use Cases:** Running NullClaw on Ubuntu with external Ollama instances (e.g., RTX 3090 / Qwen 3.6:27B) for local-first AI agent tasks; headless VPS deployments.
- **Satisfaction:** Neutral. No explicit praise or frustration, but the presence of two unresolved bugs from April/May suggests some users may be waiting for fixes. The new PRs indicate maintainers are listening.

## Backlog Watch
Three open issues have not received a maintainer response in weeks:
- **[#865](https://github.com/nullclaw/nullclaw/issues/865)** (CLI arrow keys) – opened 2026-04-23, now has a fix PR (#960) that needs review.
- **[#861](https://github.com/nullclaw/nullclaw/issues/861)** (Web UI headless) – opened 2026-04-22, no maintainer response; could benefit from a documentation update or a config example.
- **[#915](https://github.com/nullclaw/nullclaw/issues/915)** (scheduler unauthorized) – opened 2026-05-15, last updated by user on 2026-06-17; no maintainer comment. This is the most critical backlog item due to severity.

None of these have been marked as stale or closed. The PR #960 should be reviewed and merged promptly to close #865. For #915, maintainers should request logs or provide troubleshooting steps.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest – 2026-06-18

## Today’s Overview

The project saw **high activity** over the past 24 hours: **50 pull requests updated** (33 open, 17 merged or closed) and **11 issues updated** (7 open, 4 closed). No new releases were published. The Reborn WebUI and agent-loop areas dominate the change log, with a major **5‑PR stack introducing the Projects entity** and a multi‑PR overhaul of **no‑progress detection** now merging. The community is actively dogfooding the Reborn build, surfacing usability bugs (e.g., Bedrock unreachable, `NEARAI_MODEL=auto` rejection) that are being addressed immediately by fresh fix PRs. Overall project health is **strong**, with a clear cadence of feature work and rapid bug resolution.

## Releases

*None in this digest period.*

## Project Progress

**Merged / Closed PRs (17 total)** include several key fixes:

- [#3708](https://github.com/nearai/ironclaw/pull/3708) – Release chore (bumped multiple crates).  
- [#5022](https://github.com/nearai/ironclaw/pull/5022) – **Output-aware no-progress detection (PR3)**: the keystone of the redesigned progress heuristic, merging a functional `ContentDigest`‑based detection.  
- [#5000](https://github.com/nearai/ironclaw/pull/5000) – **Content-digest plumbing (PR2)**: inert plumbing for the new detection, now merged into a stacked chain.  
- [#5052](https://github.com/nearai/ironclaw/pull/5052) – **Live Slack OAuth DM‑parity**: closes [#5009](https://github.com/nearai/ironclaw/issues/5009) by hardening the non‑triggered Slack OAuth path.  
- [#4952](https://github.com/nearai/ironclaw/issues/4952) – Slack auth‑flow `cancel_flow` stale record bug was fixed (closed as part of the review chain).  

**Closed issues (4)** reflect resolved user bugs:

- [#4823](https://github.com/nearai/ironclaw/issues/4823) – “No UI feedback when deleting a running conversation” (fixed).  
- [#4974](https://github.com/nearai/ironclaw/issues/4974) – “Duplicate `...` action buttons in Activity” (fixed).  
- [#4793](https://github.com/nearai/ironclaw/issues/4793) – “Should first-run onboarding block Extensions/Automations?” (closed as question, no code change).  
- [#4952](https://github.com/nearai/ironclaw/issues/4952) – Slack auto‑deny stale AuthFlow record (fixed).  

**Open but advancing** – the **Projects feature stack** ([#5015](https://github.com/nearai/ironclaw/pull/5015) → [#5019](https://github.com/nearai/ironclaw/pull/5019)) is in active review, adding a first‑class `Project` entity, CRUD endpoints, and WebChat v2 frontend. Also under review: skill extraction & self‑evolution ([#5061](https://github.com/nearai/ironclaw/pull/5061)), read‑only agent filesystem viewer ([#5057](https://github.com/nearai/ironclaw/pull/5057)), Bedrock support ([#5059](https://github.com/nearai/ironclaw/pull/5059)), and a large dependency bump ([#4876](https://github.com/nearai/ironclaw/pull/4876)).

## Community Hot Topics

- **Dogfooding tracker** [#4879](https://github.com/nearai/ironclaw/issues/4879) – A weekly thread tracking local Reborn issues. While low‑comment (1), it aggregates first‑run usability problems and has been updated today, signalling active user testing.  
- **`NEARAI_MODEL=auto` rejection** [#5044](https://github.com/nearai/ironclaw/issues/5044) – A critical desktop‑sidecar issue quickly addressed by two fix PRs ([#5043](https://github.com/nearai/ironclaw/pull/5043), [#5045](https://github.com/nearai/ironclaw/pull/5045)).  
- **Bedrock unreachable** [#5058](https://github.com/nearai/ironclaw/issues/5058) – Another high‑impact bug with a matching fix PR [#5059](https://github.com/nearai/ironclaw/pull/5059) opened the same day.  
- **Skills validation error** [#5007](https://github.com/nearai/ironclaw/issues/5007) – A UX pain point where validation messages persist after field fill, still open without a confirmed fix.  

The **Projects stack** (5 PRs) and **no‑progress detection stack** (3 PRs) dominate review discussions, indicating heavy maintainer and reviewer attention.

## Bugs & Stability

**High severity (blocking or breakage):**

| Bug | Status | Fix PR? |
|-----|--------|---------|
| [#5060](https://github.com/nearai/ironclaw/issues/5060) – GitHub analysis workflows enter repeated approval loops | **Open** | None yet |
| [#5058](https://github.com/nearai/ironclaw/issues/5058) – Bedrock unreachable from `ironclaw-reborn` binary | **Open** | [#5059](https://github.com/nearai/ironclaw/pull/5059) |
| [#5044](https://github.com/nearai/ironclaw/issues/5044) – `NEARAI_MODEL=auto` rejected by cloud API | **Open** | [#5043](https://github.com/nearai/ironclaw/pull/5043), [#5045](https://github.com/nearai/ironclaw/pull/5045) |

**Medium severity (usability / workflow):**

| Bug | Status | Fix PR? |
|-----|--------|---------|
| [#5007](https://github.com/nearai/ironclaw/issues/5007) – Skills validation error persists after required fields filled | **Open** | None yet |
| [#5056](https://github.com/nearai/ironclaw/issues/5056) – Test issue (author `ilblackdragon`) | **Open** | Likely internal, not a real bug |

**Low severity (cosmetic/fixed):**

| Bug | Status | Fix PR? |
|-----|--------|---------|
| [#4823](https://github.com/nearai/ironclaw/issues/4823) – No UI feedback on delete running conversation | **Closed** | Fixed |
| [#4974](https://github.com/nearai/ironclaw/issues/4974) – Duplicate `...` buttons in Activity | **Closed** | Fixed |
| [#4952](https://github.com/nearai/ironclaw/issues/4952) – Slack auto‑deny stale AuthFlow record | **Closed** | [#5052](https://github.com/nearai/ironclaw/pull/5052) |

## Feature Requests & Roadmap Signals

- **Projects** (5‑PR stack) – Introducing a full Project entity with membership and role‑gated access. Expected to land in next release.  
- **Skill extraction & self‑evolution** ([#5061](https://github.com/nearai/ironclaw/pull/5061)) – Background skill learning from successful conversations. A major new capability.  
- **Agent filesystem viewer** ([#5057](https://github.com/nearai/ironclaw/pull/5057)) – Read‑only browse of agent memory and home directories in WebChat v2.  
- **Bedrock support** ([#5059](https://github.com/nearai/ironclaw/pull/5059)) – Wires the `bedrock` feature through the Reborn binary and fixes tool‑schema issues.  
- **No‑progress detection redesign** (stack [#4993](https://github.com/nearai/ironclaw/pull/4993), [#5000](https://github.com/nearai/ironclaw/pull/5000), [#5022](https://github.com/nearai/ironclaw/pull/5022)) – Now fully merged; will reduce agent hang loops.  
- **Improved `NEARAI_MODEL=auto` resolution** ([#5045](https://github.com/nearai/ironclaw/pull/5045)) – Resolves `auto` to a real model (`z-ai/glm-5.2`) to fix desktop startup.  
- **Google OAuth refresh guidance** ([#5054](https://github.com/nearai/ironclaw/pull/5054)) – Better handling of missing refresh tokens.

**Prediction for next version:** The Projects feature, skill extraction, Bedrock support, and the no‑progress improvements are likely candidates for the next release.

## User Feedback Summary

**Pain points reported (all from Reborn testing):**

- **Onboarding confusion**: First‑run flow blocks Extensions/Automations until a provider is configured (question [#4793](https://github.com/nearai/ironclaw/issues/4793)).  
- **Model configuration errors**: Desktop users hitting `NEARAI_MODEL=auto` rejection ([#5044](https://github.com/nearai/ironclaw/issues/5044)) – a serious first‑run blocker.  
- **UI feedback gaps**: Deleting a running conversation gives no feedback (now fixed [#4823](https://github.com/nearai/ironclaw/issues/4823)). Activity tool rows show duplicate buttons (fixed [#4974](https://github.com/nearai/ironclaw/issues/4974)).  
- **Validation persistence**: Skills form shows stale error messages until submit ([#5007](https://github.com/nearai/ironclaw/issues/5007)).  
- **Slack OAuth**: Triggered auth auto‑deny left stale records (fixed [#4952](https://github.com/nearai/ironclaw/issues/4952)); live OAuth path now hardened ([#5009](https://github.com/nearai/ironclaw/issues/5009), fixed in [#5052](https://github.com/nearai/ironclaw/pull/5052)).  

**Satisfaction signals**: The community is actively using the Reborn build (dogfooding), reporting issues early, and seeing rapid turnaround (most bugs get a fix PR within 24 hours). The no‑progress detection stack merging is a positive quality‑of‑life improvement for automation reliability.

## Backlog Watch

No long‑unanswered important issues are currently outstanding. The dogfooding tracker [#4879](https://github.com/nearai/ironclaw/issues/4879) is actively maintained and should continue to receive maintainer attention as weekly entries accumulate. The three high‑severity bugs ([#5060](https://github.com/nearai/ironclaw/issues/5060), [#5058](https://github.com/nearai/ironclaw/issues/5058), [#5044](https://github.com/nearai/ironclaw/issues/5044)) are all fresh (< 24h) and already have associated fix PRs, so the backlog is effectively zero. No stale PRs needing maintainer attention were identified.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-06-18

## 1. Today's Overview

The project is in a highly active phase with **11 pull requests merged within the last 24 hours**, all closed/merged, and **one new vulnerability issue (#2176) reported today**. A new release **2026.6.15** landed three days ago, introducing computer use and realtime ASR voice input. The merged PRs this cycle concentrate on fixing cowork (agent collaboration) stability, model selection, streaming, and memory management. The single open issue—a critical security flaw in automatic artifact loading—currently has no associated fix PR and warrants immediate attention.

---

## 2. Releases

**[LobsterAI 2026.6.15](https://github.com/netease-youdao/LobsterAI/releases/tag/2026.6.15)** (released 2026-06-15)

**What's Changed:**
- `feat: add computer use` – Enables agent-driven GUI automation (e.g., control other apps).
- `feat(cowork): add realtime ASR voice input` – Live speech recognition during cowork sessions.
- `feat(cowork): improve post-compaction context continuity` – Better history summarization to preserve task state after context compression.

**Breaking changes:** None documented.  
**Migration notes:** Users upgrading from earlier versions may need to review the new computer use feature’s permissions; no breaking API changes are mentioned.

---

## 3. Project Progress

All 11 PRs updated in the last 24 hours were **closed/merged**. Key advances and fixes include:

| PR | Area | Summary |
|----|------|---------|
| [#2175](https://github.com/netease-youdao/LobsterAI/pull/2175) | renderer, docs | Chore: optimize README |
| [#2174](https://github.com/netease-youdao/LobsterAI/pull/2174) | renderer, cowork | Fix scroll-to-bottom alignment with latest message height; clean up timers on session change/unmount |
| [#2173](https://github.com/netease-youdao/LobsterAI/pull/2173) | renderer, cowork | Render user messages as plain text (preserve line breaks); add diagnostic logging |
| [#2162](https://github.com/netease-youdao/LobsterAI/pull/2162) | renderer, docs, cowork | Preserve voice input cancel guard after merge; keep realtime-only ASR flow, fix draft ownership and session-switch cancellation |
| [#2153](https://github.com/netease-youdao/LobsterAI/pull/2153) | renderer, main, openclaw, cowork | Fix same-name package model selection; add regression tests and debug logging |
| [#2154](https://github.com/netease-youdao/LobsterAI/pull/2154) | renderer, main | Show model metadata after stopped streams; preserve metadata for manually stopped partial replies |
| [#2149](https://github.com/netease-youdao/LobsterAI/pull/2149) | main, openclaw | Raise OpenClaw gateway heap limit to reduce OOM crashes under long-running workloads |
| [#2147](https://github.com/netease-youdao/LobsterAI/pull/2147) | renderer, main | Prevent stopped startup turns from sending chat; cancel OpenClaw turn on early stop, emit idle state |
| [#2145](https://github.com/netease-youdao/LobsterAI/pull/2145) | docs, main | Improve cowork context compaction quality; add continuity layer, diagnostics, session-scoped task state |
| [#2144](https://github.com/netease-youdao/LobsterAI/pull/2144) | renderer, docs, main | Update portal fallback URLs to new domains |
| [#1463](https://github.com/netease-youdao/LobsterAI/pull/1463) | stale (codex) | Fix long modal titles (truncate with tooltip); applies to agent, skill, MCP, and scheduled task modals |

**Overall progress:** The team is aggressively stabilizing the cowork feature—improving voice input robustness, scroll behavior, model selection, streaming metadata, and memory management. The OOM fix and the long-dormant modal title fix (#1463, open since April) were also merged today.

---

## 4. Community Hot Topics

Only **one issue** is currently open and active:

**[#2176 – [Security] LobsterAI automatic artifact loading allows message-derived arbitrary local file reads](https://github.com/netease-youdao/LobsterAI/issues/2176)**  
- Author: YLChen-007 | Created/Updated: 2026-06-18 | Comments: 1  
- **Underlying need:** The vulnerability lets an attacker craft a message that, when parsed, causes LobsterAI to read arbitrary local files via the automatic `MEDIA:` file reference parser. The reporter suggests this could lead to data exfiltration.  
- **Analysis:** No fix PR exists yet; this is the highest-priority item in the community right now.

No other issues or PRs have received comments or reactions in the last 24 hours beyond this one.

---

## 5. Bugs & Stability

**Critical (unfixed)**
- [#2176 – Arbitrary local file read via automatic artifact loading](https://github.com/netease-youdao/LobsterAI/issues/2176) — Reported today, no patch yet. This is a security bug that could expose user files. **Severity: critical.**

**Fixed in today’s merges (all resolved)**
- **OOM crashes** in OpenClaw gateway under long-running multi-channel workloads → fixed by [#2149](https://github.com/netease-youdao/LobsterAI/pull/2149) (raise heap limit).
- **Scroll-to-bottom misalignment** after message updates → fixed by [#2174](https://github.com/netease-youdao/LobsterAI/pull/2174).
- **User message formatting** — line breaks lost in sent bubbles → fixed by [#2173](https://github.com/netease-youdao/LobsterAI/pull/2173).
- **Voice input cancel guard broken** after merge conflict → fixed by [#2162](https://github.com/netease-youdao/LobsterAI/pull/2162).
- **Model selection** for same-named packages (custom vs. built-in) → fixed by [#2153](https://github.com/netease-youdao/LobsterAI/pull/2153).
- **Stream metadata lost** when user manually stops a partial reply → fixed by [#2154](https://github.com/netease-youdao/LobsterAI/pull/2154).
- **Startup-turn send race** where a stop arrived before the gateway run became active → fixed by [#2147](https://github.com/netease-youdao/LobsterAI/pull/2147).

**Stability improvements:**
- Context compaction quality enhanced via [#2145](https://github.com/netease-youdao/LobsterAI/pull/2145) (adds continuity layer, diagnostics, session-scoped task state).

---

## 6. Feature Requests & Roadmap Signals

No explicit feature requests were filed in the last 24 hours. However, based on merged PRs and the latest release, several signals emerge:

- **Computer use** (release 2026.6.15) — a major new capability for GUI automation. Likely to see further integration with cowork and voice input.
- **Realtime ASR voice input** — now live in cowork. Future versions may extend it to standalone sessions or add multi-language support.
- **Context continuity improvements** — the team is investing heavily in preserving agent state across history compression. This points toward longer-running, more autonomous agent tasks.
- **Model selection** enhancements (PR #2153) suggest users were confused by duplicate model names; future iterations may bring a UI redesign for model browsing.
- **Security** — given the severity of #2176, a hotfix release (2026.6.16 or 2026.6.17) is likely to address the artifact loading vulnerability.

**Prediction for next version:** A security patch release containing a fix for #2176, plus further stability tweaks for voice input and cowork turn handling.

---

## 7. User Feedback Summary

**Direct feedback (from issues):**
- **Critical security concern:** “LobsterAI automatically parses `MEDIA:` file references … allowing arbitrary local file reads.” (Issue #2176) — the highest-priority user pain point today.

**Inferred pain points from fix PRs:**
- **Voice input instability:** Users experienced lost cancel guards, session-switch errors, and merge conflicts in ASR flow (PR #2162).
- **Model confusion:** Users could not distinguish between custom models and built-in packages with the same name (PR #2153).
- **Broken streaming UX:** Manually stopping a partial assistant reply lost model metadata (PR #2154).
- **OOM crashes during long cowork sessions** — fixed by heap limit increase (PR #2149).
- **Irritating UI:** Scroll did not stay at the bottom after new messages; user messages lost line breaks (PRs #2174, #2173).
- **Old bug:** Long modal titles overflowed dialogs (PR #1463, open since April, finally fixed today).

**Satisfaction indicators:** The team is merging fixes rapidly (11 PRs in one day), which suggests high responsiveness to reported issues.

---

## 8. Backlog Watch

| Issue/PR | Age | Status | Attention needed |
|----------|-----|--------|------------------|
| [#1463](https://github.com/netease-youdao/LobsterAI/pull/1463) – Fix long modal titles | Created 2026-04-04 (10 weeks) | **Merged today** after being stale | ✅ Resolved |
| [#2176](https://github.com/netease-youdao/LobsterAI/issues/2176) – Security vulnerability | Created today | Open, no fix PR | **Urgent** – needs maintainer review and patch |
| *No other long-stale items* | – | – | – |

**Note:** No open issues older than 2 weeks remain unaddressed. The backlog is clean except for the newly reported security bug. Maintainers should prioritize a fix for #2176, possibly as a hotfix release.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest — 2026-06-18

## Today's Overview
The Moltis project saw moderate activity over the past 24 hours, with two open feature requests and one open pull request. No issues or PRs were closed or merged, indicating a stable but unmerged work-in-progress state. The project has no new releases, and all activity remains in the proposal and implementation stage. Overall health appears steady, with continued community engagement around configuration and export capabilities.

## Releases
*No new releases in the last 24 hours.*

## Project Progress
No pull requests were merged or closed today.  
One new PR was opened:
- **PR #1130** *(open)* – **feat: make webui rpc timeout configurable**  
  Author: khimaros  
  Summary: “written on the tin. fixes #1127”  
  This PR addresses a configuration improvement for the WebUI RPC timeout, linking to an unresolved issue (#1127). It is the only active code change under review.  
  [View PR](https://github.com/moltis-org/moltis/pull/1130)

## Community Hot Topics
The most active discussion this period centers on **Issue #1126**:
- **#1126** – `[Feature]: allow to configure the format of tts output`  
  Author: khimaros | Created: 2026-06-16 | Updated: 2026-06-17 | Comments: 3 | 👍: 0  
  This request has attracted three comments, indicating a clear community interest in controlling the output format of the Text-to-Speech engine (e.g., file type, encoding). The author also contributed the related PR #1130, suggesting they are an active contributor.  
  [View Issue](https://github.com/moltis-org/moltis/issues/1126)

The second issue, **#1131** (copy + export as Markdown), has no comments yet but aligns with a common user need for interoperability.  
[View Issue](https://github.com/moltis-org/moltis/issues/1131)

## Bugs & Stability
No bugs, crashes, or regressions were reported in the last 24 hours. The project currently shows zero open bug issues. Stability remains untroubled by new defects.

## Feature Requests & Roadmap Signals
Two enhancement requests surfaced today:
1. **Configurable TTS output format** (#1126) – users want to choose the output format (e.g., WAV, MP3, text). Given that a PR for a configurable timeout (#1130) was already submitted by the same author, this feature may be prioritized in a future minor release.
2. **Copy and export as Markdown** (#1131) – users want to export conversation history or tool output in Markdown format. This aligns with standard interoperability expectations and could land in the next feature batch if community demand grows.

Both requests are labeled `enhancement` and have no pushback from maintainers, making them candidates for inclusion in an upcoming release.

## User Feedback Summary
Real user pain points visible in today’s activity:
- *Lack of TTS output configuration* – users are constrained to a single output format, limiting integration with different audio pipelines.
- *Inability to export structured text* – the request for Markdown export indicates users need to share or archive assistant outputs in a readable, portable format.
- *WebUI timeout adjustment* (implied by PR #1130) – suggests users need finer control over network stability in RPC interfaces.

No negative satisfaction signals were observed; all requests are constructive enhancement ideas with clear use cases.

## Backlog Watch
The following issue, referenced but not listed in the “Latest Issues” section, deserves attention:
- **#1127** – linked from PR #1130 (fixes #1127). Its content is unknown from this snapshot, but it is the target of a pending change. If this issue remains unresolved or lacking maintainer comment, it may be a blocking item for the PR. A maintainer review of #1127 is recommended before merging.  
  [View Issue #1127](https://github.com/moltis-org/moltis/issues/1127)

No other long-unanswered issues are apparent from the available data.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest — 2026-06-18

## Today's Overview
The project is experiencing very high activity: **42 issues** and **50 PRs** were updated in the last 24 hours, with a closure/merge rate of 76% and 66% respectively. Two releases were published today — v1.1.12 (stable) and v1.1.12-beta.2 (pre-release). The majority of work focused on bug fixes, UI improvements, and community contributions (10 PRs authored by first-time contributors). While the project demonstrates strong community engagement, several stability-critical bugs (process freezes, infinite loops, and channel misrouting) remain open and require maintainer attention.

## Releases

### v1.1.12 (Latest)  
- **Console:** Overhauled Models Page with provider aggregation, unified card UI, and layout redesign ([#5203](https://github.com/agentscope-ai/QwenPaw/pull/5203))  
- **Console:** Introduced "Simple Mode" with flat navigation and sorted session list by update time ([#5222](https://github.com/agentscope-ai/QwenPaw/pull/5222))  
- **Breaking changes:** None reported.  
- **Migration:** Standard in-place upgrade; no config changes required.

### v1.1.12-beta.2 (Pre-release)  
- Removed unnecessary deep copy operations in agent config ([#5240](https://github.com/agentscope-ai/QwenPaw/pull/5240))  
- Added session filter by title ([#5178](https://github.com/agentscope-ai/QwenPaw/pull/5178))  
- Various dependency bumps and chore updates.

*Note: v1.1.12-beta.2 was published earlier today; v1.1.12 supersedes it as the stable release.*

## Project Progress
**33 PRs were merged/closed today**, advancing the following areas:

### Fixes
- **ChromaDB runtime probe collection name** – renamed `_probe` to `probe-test` (fixes #5284, PR [#5289](https://github.com/agentscope-ai/QwenPaw/pull/5289))  
- **Proactive responder cache pollution** – prevents `load_agent_config()` mutation (PR [#5275](https://github.com/agentscope-ai/QwenPaw/pull/5275))  
- **Backup fails on unreadable files** – skips problematic files instead of aborting (PR [#5041](https://github.com/agentscope-ai/QwenPaw/pull/5041), closes #4916)  
- **Duplicated `session_id` in filenames** – avoids `{sid}_{sid}.json` when `user_id` equals `session_id` (PR [#5026](https://github.com/agentscope-ai/QwenPaw/pull/5026))  
- **XiaoYi A2A channel** – refactored to dual WebSocket connections (PR [#5274](https://github.com/agentscope-ai/QwenPaw/pull/5274) and [#3839](https://github.com/agentscope-ai/QwenPaw/pull/3839))

### Features
- **New provider: Venice AI** – reuses OpenAI compatibility layer (PR [#1088](https://github.com/agentscope-ai/QwenPaw/pull/1088))  
- **Migration tool (`qwenpaw migrate openclaw`)** – imports configs from OpenClaw installations (PR [#5276](https://github.com/agentscope-ai/QwenPaw/pull/5276), related issue #5254)  
- **Cron update CLI** – `qwenpaw cron update <job_id>` for modifying existing cron jobs (PR [#5210](https://github.com/agentscope-ai/QwenPaw/pull/5210), closes #4939)  
- **Chat history right panel** – embeds history as a permanent sidebar (open PR [#5293](https://github.com/agentscope-ai/QwenPaw/pull/5293))

### Documentation & Tooling
- Updated roadmap ([#5277](https://github.com/agentscope-ai/QwenPaw/pull/5277))  
- Release automation and version bumps ([#5280](https://github.com/agentscope-ai/QwenPaw/pull/5280), [#5288](https://github.com/agentscope-ai/QwenPaw/pull/5288))

## Community Hot Topics

| Issue/PR | Type | Comments | Summary |
|----------|------|----------|---------|
| [#280 – Built-in Skills & MCPs](https://github.com/agentscope-ai/QwenPaw/issues/280) | Discussion (Closed) | 27 | Community strongly supports pre-installing popular skills/MCPs for better OOTB experience |
| [#5218 – Context Compaction Freeze](https://github.com/agentscope-ai/QwenPaw/issues/5218) | Bug (Open) | 16 | Sub-agent context compaction freezes entire process – **critical stability concern** |
| [#4108 – WebUI lag](https://github.com/agentscope-ai/QwenPaw/issues/4108) | Question (Closed) | 8 | Users report severe desktop stuttering while generating responses |
| [#4967 – Infinite execution loop](https://github.com/agentscope-ai/QwenPaw/issues/4967) | Bug (Open) | 6 | Agent enters unrecoverable infinite loop – similar to #5162 and #4873 |
| [#5262 – Disabled skills re-enable on upgrade](https://github.com/agentscope-ai/QwenPaw/issues/5262) | Bug (Open) | 4 | Every upgrade resets disabled built-in skills (persistent frustration) |
| [#5204 – Two agents infinite loop via Matrix](https://github.com/agentscope-ai/QwenPaw/issues/5204) | Bug (Open) | 2 | Cross-agent feedback loop – no runtime circuit breaker exists |

**Underlying needs:** The community demands better built-in capabilities (skills/MCPs), robust process isolation (preventing freezes and infinite loops), and configuration persistence across upgrades.

## Bugs & Stability
**Ranked by severity** (all reported today or still open):

| Severity | Issue | Description | Fix PR Exists? |
|----------|-------|-------------|----------------|
| **Critical** | [#5218](https://github.com/agentscope-ai/QwenPaw/issues/5218) | Process freeze on sub-agent context compaction | No (related timeout PR [#5242](https://github.com/agentscope-ai/QwenPaw/pull/5242) *open*) |
| **Critical** | [#4967](https://github.com/agentscope-ai/QwenPaw/issues/4967) | Agent infinite loop (ReAct pattern) | No |
| **Critical** | [#5204](https://github.com/agentscope-ai/QwenPaw/issues/5204) | Cross-agent infinite loop via Matrix | No |
| **High** | [#5262](https://github.com/agentscope-ai/QwenPaw/issues/5262) | Disabled skills re-enable on upgrade | No |
| **High** | [#5264](https://github.com/agentscope-ai/QwenPaw/issues/5264) | Feishu group chat replies misrouted to private chat | No |
| **Medium** | [#5292](https://github.com/agentscope-ai/QwenPaw/issues/5292) | Manually added provider not shown in model selector | No |
| **Medium** | [#5284](https://github.com/agentscope-ai/QwenPaw/issues/5284) | ChromaDB probe fails due to collection name `_probe` | Yes ([#5289](https://github.com/agentscope-ai/QwenPaw/pull/5289) merged) |

**Notable:** The `_compact_context` timeout protection PR [#5242](https://github.com/agentscope-ai/QwenPaw/pull/5242) directly addresses the freeze in #5218, but remains open. The ChromaDB probe fix was merged today.

## Feature Requests & Roadmap Signals

| Request | Issue/PR | Likely Next Version? |
|---------|----------|---------------------|
| **Built-in skills/MCPs** (popular pre-installed) | [#280](https://github.com/agentscope-ai/QwenPaw/issues/280) | High – strong community demand, closed as discussion |
| **HTTP API** (direct API without Web UI) | [#2202](https://github.com/agentscope-ai/QwenPaw/issues/2202) | Medium – multiple users request |
| **Web console upgrade** (remote update via UI) | [#2235](https://github.com/agentscope-ai/QwenPaw/issues/2235) | Low – no current activity |
| **Merge provider cards** (e.g., Zhipu variants) | [#4965](https://github.com/agentscope-ai/QwenPaw/issues/4965) | Medium – UI quality-of-life |
| **Scheduled task history** | [#1366](https://github.com/agentscope-ai/QwenPaw/issues/1366) | Low – has local branch, but not merged |
| **OpenClaw migration tool** | [#5276](https://github.com/agentscope-ai/QwenPaw/pull/5276) | Already merged – will be in next release |
| **Customizable column order in sessions** | [#4975](https://github.com/agentscope-ai/QwenPaw/pull/4975) | Medium – PR open, under review |

**Prediction:** The next release (likely v1.1.13 or v2.0.0a line) will focus on built-in skills (based on #280 momentum), HTTP API exposure, and the high-priority stability fixes (#5218, #4967).

## User Feedback Summary
- **Satisfaction:** Positive contributions from first-timers (Venice AI provider, backup fix, cron update). Users appreciate the new Models Page and Simple Mode.
- **Pain points:**
  - **Performance regressions** – WebUI causes system-wide lag during generation (#4108).
  - **Configuration bleeding** – Upgrades reset disabled skills (#5262) and skills re-disable after edit (#3090).
  - **Stability** – Process freezes (#5218), infinite loops (#4967, #5204), and channel misrouting (#5264) erode trust.
  - **Tooling inconsistencies** – Path resolution between `read_file` and `execute_shell_command` (#5207), missing model providers in selector (#5292).
- **Common demand:** More refined user experience out-of-the-box (built-in skills, better channel routing), and robust error handling.

## Backlog Watch
Open issues requiring maintainer attention due to high impact or long neglect:

| Issue | Age | Priority |
|-------|-----|----------|
| [#4967 – Infinite loop](https://github.com/agentscope-ai/QwenPaw/issues/4967) | 13 days | Critical – no fix PR yet |
| [#5204 – Cross-agent Matrix loop](https://github.com/agentscope-ai/QwenPaw/issues/5204) | 3 days | Critical – no fix PR yet |
| [#5262 – Disabled skills reset](https://github.com/agentscope-ai/QwenPaw/issues/5262) | 1 day (regression) | High – affects upgrade UX |
| [#5264 – Feishu routing bug](https://github.com/agentscope-ai/QwenPaw/issues/5264) | 1 day | High – platform-specific |
| [#5292 – Missing provider](https://github.com/agentscope-ai/QwenPaw/issues/5292) | 0 days | Medium – reported today, no PR |
| [#3839 – XiaoYi A2A (closed but may need backport?)](https://github.com/agentscope-ai/QwenPaw/pull/3839) | Closed | Already merged; no issue. |
| [#1088 – Venice AI provider](https://github.com/agentscope-ai/QwenPaw/pull/1088) | Closed | Merged. |

**Recommendation:** Prioritize fixes for process freezes (#5218, #4967) and cross-agent loops (#5204) as they directly affect basic usability. The disabled skills reset (#5262) should be a straightforward config migration save.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-06-18

---

## 1. Today’s Overview

ZeroClaw saw heavy development activity over the past 24 hours, with **50 pull requests updated** (11 merged/closed) and **17 issues updated** (16 still open). No new releases were published. The team is actively pushing forward multiple milestone trackers (**v0.8.1**, **v0.8.2**, **v0.8.3**, and **v0.9.0**), while addressing a steady stream of bugs and regressions. The most notable merged work involves a stacked series of config management improvements (typed delete-with-cascade, alias rename, and CLI CRUD for agents/providers/channels), alongside fixes for ACP event visibility. However, several high-severity issues remain open, including a Windows test regression (74 failures) and a newly discovered state-persistence race in agent rename. Overall, the project is in an active refinement phase with clear roadmap signals toward WASM plugins, skills platform, and security hardening.

---

## 2. Releases

**No new releases in the last 24 hours.** The most recent release remains the v0.8.x series. Upcoming milestones (v0.8.2, v0.8.3, v0.9.0) are tracked in dedicated issues (see Section 6).

---

## 3. Project Progress — Merged/Closed PRs

**11 PRs were merged or closed today.** The most significant among those captured in the data are:

- **#7842** — [feat(cli): agents/providers/channels CRUD + skill-bundle cascade](https://github.com/zeroclaw-labs/zeroclaw/pull/7842)  
  Final slice of the stacked series (#7468/#7175). Adds CLI commands for typed delete-with-cascade and alias rename, including cascade for skill bundles.

- **#7841** — [feat(gateway): agent owned-state rename cascade + rename wiring](https://github.com/zeroclaw-labs/zeroclaw/pull/7841)  
  Enables renaming agents to cascade updates to owned state (e.g., memory, cron jobs, gateways). This PR implements the core rename path.

- **#7840** — [feat(config): rename_with_cascade for aliased entries](https://github.com/zeroclaw-labs/zeroclaw/pull/7840)  
  Provides the config-layer implementation for propagating alias renames through provider and channel references.

- **#7684** — [fix(acp): surface history-pruner and turn-cancel as visible events](https://github.com/zeroclaw-labs/zeroclaw/pull/7684)  
  Previously, the history pruner’s collapsed messages appeared as plain bot output; now they render as styled system events, improving UX for ACP users.

- **#7563** — [closed issue] canvas-store regression in WS chat/ACP sessions (S1, fixed by preceding PR #6986 follow-up). The issue was closed, indicating the fix is live.

These merges mark solid progress on configuration management and user-facing ACP improvements, but also introduce a newly reported bug (see #7907, Section 5).

---

## 4. Community Hot Topics

The most active issues (by comment count) reflect deep architectural discussions and cross-cutting concerns:

- **#7673** — [RFC: Native context compression as a provider pipeline decorator](https://github.com/zeroclaw-labs/zeroclaw/issues/7673) (3 comments)  
  Proposes a `CompressionDecorator` to compress `ChatRequest` payloads before forwarding to LLMs. This is a high-risk, high-impact enhancement that could reduce latency and cost. Active discussion on trade-offs between compression and provider compatibility.

- **#6970** — [v0.8.1 integration/channel/provider/tool queue and history](https://github.com/zeroclaw-labs/zeroclaw/issues/6970) (3 comments)  
  Operational tracker for the current release queue. Community members are tracking the status of additive channels, providers, and tool improvements.

- **#7175** — [feat(config): typed delete-with-cascade for aliased entries](https://github.com/zeroclaw-labs/zeroclaw/issues/7175) (2 comments)  
  Now largely implemented by the merged PR stack. The conversation focused on ensuring cascading deletions do not corrupt cross-type references.

- **#7675** — [RFC: Hardened CI pipeline — supply-chain scanning, provenance, SBOM](https://github.com/zeroclaw-labs/zeroclaw/issues/7675) (2 comments)  
  Calls for a new CI security gate. Community members are interested in SBOM generation and dependency provenance verification.

- **#7822** — [RFC: WASM plugin lifecycle hook subscriptions](https://github.com/zeroclaw-labs/zeroclaw/issues/7822) (1 comment)  
  Proposes `PluginCapability::Hook` to let WASM plugins subscribe to agent lifecycle events. This is a key enabling feature for the v0.8.2 WASM plugin program.

**Underlying needs:** The community is pushing for better operational security (CI hardening, SBOM), performance optimization (context compression), and extensibility (WASM hooks). The discussion around #7673 and #7822 indicates a desire for more modular, efficient, and secure plugin and provider architectures.

---

## 5. Bugs & Stability

Several bugs were reported or actively fixed today. Ranked by severity:

| Severity | Issue / PR | Description | Status |
|----------|------------|-------------|--------|
| **S1 – Blocked** | [#7907](https://github.com/zeroclaw-labs/zeroclaw/issues/7907) | Agent rename moves owned state (e.g., memory, cron) before config persistence is durable. Introduced by #7841 merge. Workflow blocked until order is corrected. | Open, no fix PR yet |
| **S1 – Blocked** | [#7563](https://github.com/zeroclaw-labs/zeroclaw/issues/7563) (closed) | Canvas-store regression in WS chat/ACP sessions – fixed and closed. | Resolved |
| **S2 – Degraded** | [#7462](https://github.com/zeroclaw-labs/zeroclaw/issues/7462) | 74 test failures on Windows (Unix-only commands, path semantics, console encoding). CI only runs on Linux. | Open, needs maintainer action |
| **S2 – Degraded** | [#7737](https://github.com/zeroclaw-labs/zeroclaw/issues/7737) | Approval attribution still uses channel-global side channel – concurrent approvals can overwrite state before consumption. | Open, status:accepted |
| **S2 – Degraded** | [#7819](https://github.com/zeroclaw-labs/zeroclaw/pull/7819) | Missing-skill suggestions incorrectly considered all registered tools, suppressing install suggestions when a tool was excluded per-turn. | Fix PR open |
| **S2 – Degraded** | [#7853](https://github.com/zeroclaw-labs/zeroclaw/pull/7853) | Windows self-update fundamentally broken (remove-then-copy fails because Windows locks running image). PR includes swap fix and rollback hardening. | Fix PR open |

Other notable bug-fix PRs:
- **#7901** — Bounds repeated shell approval loops (turn-local guard) → [PR](https://github.com/zeroclaw-labs/zeroclaw/pull/7901)
- **#7902** — Pins `http_request` to vetted DNS addresses (SSRF prevention) → [PR](https://github.com/zeroclaw-labs/zeroclaw/pull/7902)
- **#7903** — Fixes ACP session history not replaying in `session/messages` → [PR](https://github.com/zeroclaw-labs/zeroclaw/pull/7903)
- **#7893** — Persists manual cron trigger results → [PR](https://github.com/zeroclaw-labs/zeroclaw/pull/7893)
- **#7908** — Repairs WebDriver snapshot returns and CSS selector escaping → [PR](https://github.com/zeroclaw-labs/zeroclaw/pull/7908)

**Summary:** The most critical unresolved bug is #7907 (state persistence race in agent rename). Windows stability (#7462, #7853) remains a pain point, though a fix PR exists for self-update. Several S2 issues have active PRs.

---

## 6. Feature Requests & Roadmap Signals

Requests and RFCs filed or updated today point to upcoming priorities:

- **Context compression decorator** ([#7673](https://github.com/zeroclaw-labs/zeroclaw/issues/7673)) – A high-impact, high-risk RFC. Likely candidate for v0.9.0 (auth/security/breaking-change milestone) given its pipeline change.

- **Hardened CI pipeline** ([#7675](https://github.com/zeroclaw-labs/zeroclaw/issues/7675)) – Supply-chain scanning and SBOM generation. Could land in v0.8.x as a security uplift.

- **WASM plugin lifecycle hook subscriptions** ([#7822](https://github.com/zeroclaw-labs/zeroclaw/issues/7822)) – Status:accepted, tracked under v0.8.2 WASM plugin program (#7314). Likely to be implemented in the coming weeks.

- **Zero-downtime security/channel reload** ([#7897](https://github.com/zeroclaw-labs/zeroclaw/issues/7897)) – Scoped reload without full daemon restart. Part of v0.9.0.

- **Intra-family provider fallback notices** ([#7883](https://github.com/zeroclaw-labs/zeroclaw/issues/7883)) – Low priority (P3) but accepted. Likely v0.8.3 or later.

**Roadmap signals from tracker issues:**
- **v0.8.2** (skills platform + WASM plugin program) – #7852, #7314
- **v0.8.3** (MCP dashboard, web/plugin-management) – #7320
- **v0.9.0** (auth, security, gateway boundary) – #7432

**Prediction:** Context compression (#7673) and WASM hooks (#7822) are the most likely features to land in the next minor release (v0.8.2), given their high community interest and alignment with the skills/plugin platform.

---

## 7. User Feedback Summary

Real pain points surfaced by users and contributors:

- **Android Termux installation failure** ([#7911](https://github.com/zeroclaw-labs/zeroclaw/issues/7911)) – The precompiled binary and local compilation both fail because ZeroClaw does not recognise `aarch64-linux-android`. User is blocked.

- **Windows parity issues** – 74 test failures ([#7462](https://github.com/zeroclaw-labs/zeroclaw/issues/7462)) and broken self-update ([#7853](https://github.com/zeroclaw-labs/zeroclaw/pull/7853)) are ongoing frustrations. The community is actively contributing fixes (e.g., NiuBlibing’s Windows work), but CI still only runs Linux, making Windows quality a blind spot.

- **Canvas UI regression** ([#7563](https://github.com/zeroclaw-labs/zeroclaw/issues/7563), closed) – The `/canvas` page went blank after a WebSocket chat session. Users reported a blocked workflow; the fix has been merged.

- **Secret prompt feedback** ([#7856](https://github.com/zeroclaw-labs/zeroclaw/pull/7856)) – Users complained that password prompts provided zero confirmation; the PR adds post-input echo for feedback.

- **Git operations error** ([#7835](https://github.com/zeroclaw-labs/zeroclaw/pull/7835)) – Bare “Not in a git repository” message with no context or recovery hint; PR adds Fluent i18n message with path context.

Overall satisfaction is tempered by the Windows situation, but users appreciate the rapid response to bugs (e.g., canvas fix, secret prompt). The Android Termux request is unmet and may require additional packaging work.

---

## 8. Backlog Watch

Issues and PRs that have been unanswered or need maintainer attention:

- **#7911** (Android Termux support) – Filed today, zero comments. Needs a maintainer to identify if cross-compilation for Android is planned or if a workaround exists.

- **#7462** (74 Windows test failures) – Filed June 10, updated today but only one comment. Despite ongoing Windows fixes (e.g., #7853), the root CI gap remains. This issue likely needs a champion to add a Windows CI job.

- **#7673** (context compression RFC) – Needs more community input or maintainer decision on direction. Label includes `needs-author-action`, meaning the author may need to provide more detail.

- **#7675** (hardened CI RFC) – Similarly `needs-author-action`. No movement beyond initial proposal.

- **#7821** (feat(config): add schema struct & risk field) – PR with `needs-author-action`. The PR adds infrastructure but is waiting on author response for review feedback.

- **#7098** (Mattermost WebSocket listener) – Open since June 2, updated today but still `needs-author-action`. The feature is large (XL) and likely needs maintainer guidance to move forward.

- **#7835** (git_operations error fix) – Also `needs-author-action`. The PR is straightforward but awaiting author’s response to review comments.

**Recommendation:** Maintainers should prioritize replying to the `needs-author-action` items to unblock these contributions, and assess the Windows CI gap (#7462) as a high-priority infrastructure improvement. The Android support request (#7911) is a new but potentially valuable expansion.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*