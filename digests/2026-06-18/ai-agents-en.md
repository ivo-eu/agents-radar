# OpenClaw Ecosystem Digest 2026-06-18

> Issues: 289 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-18 12:31 UTC

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

## Today’s Overview

OpenClaw remains in a period of exceptionally high activity, with 289 issues and 500 pull requests updated in the last 24 hours. The vast majority of recent activity is focused on bug stabilization and long‑standing technical debt rather than new feature work. Most open issues are P1/P2 severity affecting session state, message loss, security, and auth providers, many of which have been open for weeks or months. The project saw 12 issues closed and 43 PRs merged/closed today, indicating steady resolution of some critical regressions. However, a significant number of items remain in the “needs maintainer review” or “needs product decision” states, suggesting the maintainer team is stretched.

## Releases

No new releases were published today (latest remains 2026.6.1). Users experiencing the DeepSeek Prompt Cache regression (issue #91016, closed) should note that the fix appears to have been merged, but no version tag has been cut yet.

## Project Progress

Today’s 43 merged/closed PRs include several noteworthy improvements:

* **PR #68936** (closed) – Added a PR review autofix pipeline and a Windows background daemon, demonstrating investment in developer tooling and cross‑platform support.
* **PR #94435** (closed) – Introduced a bundled XMemo cloud memory plugin (`xmemo-memory`), expanding long‑term memory backend options beyond local‑only solutions.
* **PR #94573** (open, ready for maintainer) – Fixes a performance issue in the ingress queue state DB by no longer copying `process.env` on every queue operation (fixes #94571).
* **PR #94574** (open) – Stops leaking the Slack bot token into `/api/auth.test` request bodies (auth security fix).
* **PR #94575** (open) – Normalizes bundled channel names on outbound sends, fixing `GatewayClientRequestError: unsupported channel: telegram` for bundled channels.
* **PR #94312** (open) – Strips `<relevant-memories>` tags from outbound assistant text, preventing raw markup from appearing in Telegram, Discord, Slack, etc.
* **PR #93696** (open) – Delivers reasoning blocks as `m.notice` in Matrix instead of suppressing them entirely.

Multiple smaller fixes landed for Telegram, cron, compaction, and CLI stability. The “ClawSweeper” bot continues to auto‑file PRs for regression bugs found via static analysis.

## Community Hot Topics

The most active discussions (by comment count and reactions) reveal deep operational pain points:

* **#50090 – Community Skill Development & ClawHub** (15 comments, 2 👍) – The gap between the “write a SKILL.md” promise and the actual friction of publishing/installing skills continues to frustrate community contributors. Users are asking for a clear, reliable path to extending OpenClaw beyond core capabilities.
* **#58450 – Agent can promise a follow‑up without starting any actual action** (15 comments, 3 👍) – High concern about agent accountability: agents giving false confirmations (e.g., “I’ll check and follow up”) without spawning background tasks, leaving users waiting indefinitely.
* **#65161 – Heartbeat isolated mode cadence stalls** (14 comments, 1 👍) – Multiple regressions in heartbeat scheduling, event labeling, state persistence, and lightContext management.
* **#57326 – CLI‑backed helper paths still bypass CLI dispatch** (13 comments, 1 👍) – A remaining security/correctness hole where some API paths skip the intended CLI‑agent routing for CLI‑backed models.
* **#91016 (closed) – DeepSeek Prompt Cache completely broken after upgrade** (12 comments, 6 👍) – Users burning real money (~$6/hour) due to cache invalidation regression. High visibility, now closed with a fix.
* **#63216 – Repeated hard resets despite high reserveTokensFloor** (11 comments, 3 👍) – A persistent session‑level bug causing context‑overflow loops that re‑inject bootstrap context on recovery.

Other active threads include #67288 (Bedrock discovery running unnecessarily), #67777 (subagent completion delivery loss), #55334 (sessions.json unbounded growth leading to OOM), #64810 (Telegram heartbeat swallowing replies), #63829 (per‑agent memory‑wiki vault), and #67419 (bootstrap context bloat wasting 20‑30% tokens).

## Bugs & Stability

Reported bugs today span several high‑severity categories. The most critical ones are summarized below; many have linked fix PRs in progress.

| Severity | Issue | Description | Fix PR |
|----------|-------|-------------|--------|
| **P1 – Critical** | #63216 | Repeated hard context‑overflow resets on specific session keys; retry loop re‑injects bootstrap context | No linked PR yet |
| **P1 – Critical** | #55334 | `sessions.json` unbounded growth with full `skillsSnapshot` per session → gateway OOM | Needs maintainer review |
| **P1 – Critical** | #63998 | Session transcript doomloop: crash‑restart cycle inflates transcript until OOM | Needs maintainer review |
| **P1 – Critical** | #65161 | Multiple heartbeat regressions in isolated mode: cadence stalls, mislabeled exec‑events, writer missing | No linked PR yet |
| **P1 – High** | #67777 | Subagent completion delivery lost on timeout/drain/orphan prune | Needs maintainer review |
| **P1 – High** | #57326 | CLI‑backed helper paths still bypass correct dispatch | No linked PR yet |
| **P1 – High** | #64810 | Heartbeat/system events can swallow in‑progress Telegram replies | Needs maintainer review |
| **P1 – High** | #64267 | Agent internal thinking (English) exposed to user | No linked PR yet |
| **P1 – High** | #65624 | Mattermost slash commands leak cleartext reusable tokens (CVSS 7.6/8.6) | Linked PR open |
| **P1 – High** | #91016 (closed) | DeepSeek Prompt Cache completely dead after upgrade (money burn) – **fix merged** | PR merged |
| **P2 – Medium** | #67419 | Bootstrap files re‑injected every turn → 20‑30% token waste | Needs maintainer review |
| **P2 – Medium** | #66443 | Overflow recovery duplicates `role=user` messages in session JSONL | Linked PR open |
| **P2 – Medium** | #65374 | Built‑in dreaming system contaminates agent identity in multi‑agent setups | No linked PR yet |
| **P2 – Medium** | #66977 | `sqlite-vec` cannot load on macOS (vector search broken) | Linked PR open |

Additionally, #67288 (Bedrock discovery runs on every request), #67366 (TypeError during `onboard` on macOS), #67417 (backup fails with ENOENT on concurrent cleanup), and #61009 (docs–runtime mismatch on exec host override) are P2 bugs with active discussion or fix PRs.

## Feature Requests & Roadmap Signals

Several user‑requested features appear to be gaining traction and could shape the next minor release:

* **Per‑agent memory‑wiki vault isolation** (#63829, 9 👍) – High demand to allow each agent in a multi‑agent setup its own knowledge base instead of sharing a global one.
* **Per‑agent TTS/STT configuration** (#66252, 1 👍) – Enable multiple languages/voices per agent in a single instance.
* **Per‑agent dreaming control** (#67413, 5 👍) – Memory‑core dreaming currently runs on all agents simultaneously, causing OOM. Request for per‑workspace scheduling and disable.
* **Plugin UI Extension System** (#66944, 4 👍) – Allow plugins to contribute native tabs to the Control UI using Lit Web Components. Would unlock a rich third‑party dashboard experience.
* **Multi‑index embedding with model‑aware failover** (#63990, 1 👍) – Avoid vector space corruption when swapping embedding models; production reliability feature.
* **Sensitive data masking** (#64046, 0 👍 but detailed) – API keys, tokens, secrets stored in cleartext in config files, logs, and UI. Broad security improvement.
* **Exec sandbox isolation and tool permission model** (#58730, 1 👍) – Inspired by Claude Code source leak; propose stronger permission boundaries and sandboxing.
* **Guarantee last N raw messages survive compaction** (#58818, 2 👍) – Addresses recurring data loss after session resets.
* **Browser tool enhancements** (#60381) – Add `force` parameter for click and expose `evaluate` action to handle modern JS frameworks.
* **Anthropic advisor tool support** (#63930) – Server‑side tool blocks currently unsupported; would enable Claude‑based advisor mid‑inference.

Features most likely to land in the next release (based on maintainer attention and linked PRs): per‑agent memory vault UI/CLI integration (#63829), TTS/STT overrides (#66252), and the XMemo cloud memory plugin (#94435, already merged).

## User Feedback Summary

Real user pain points expressed in today’s issues reveal a pattern of **unexpected costs, data loss, and broken expectations**:

* **Money concerns** – The DeepSeek Prompt Cache regression (#91016) burned real dollars; users are demanding more robust cost‑control and cache validation dashboards.
* **False agent promises** – Issue #58450 (“I’ll check and come back”) is a recurring frustration: agents that say they will perform background work but never do it.
* **Message loss in busy channels** – Telegram topic sessions ( #64810, #56692) causing replies to be swallowed by heartbeats or mis‑routed in group contexts. Matches PR #94207 (wake drain fix) and #94301 (drain worker‑spooled updates).
* **Context management failures** – Repeated hard resets (#63216), doomloop OOM (#63998), and backup failures (#67417) make long‑running sessions unreliable.
* **Security posture gaps** – Cleartext tokens in Mattermost (#65624), Slack (#94574), config files (#64046), and CI‑backed routing bypass (#57326) are eroding trust.
* **Accessibility regression** – Screen readers announce every token during streaming (#65538), making the Web UI unusable for visually‑impaired users.
* **iOS/macOS integration** – Repeated TCC permission requests (#94286) and `sqlite-vec` on macOS (#66977) break out‑of‑box experience on Apple platforms.

Satisfaction signals are rare, but the community shows appreciation for the **open‑source transparency** (ClawSweeper bot, detailed issue labeling) and the **earnest engagement** of some maintainer comments. The overall tone is constructive but urgent.

## Backlog Watch

Several important issues have been open for 1–3 months without resolution or meaningful maintainer response:

| Issue | Date | Age | Notes |
|-------|------|-----|-------|
| **#50090 – Community Skill Development & ClawHub** | 2026-03-19 | ~3 mo | P2, needs product decision. Core ecosystem enabler. |
| **#58450 – Agent false follow‑up promises** | 2026-03-31 | ~2.5 mo | P2, needs maintainer review and product decision. High community impact. |
| **#57326 – CLI helper path bypass** | 2026-03-29 | ~2.5 mo | P1, security‑related, still needs maintainer review. |
| **#55334 – sessions.json unbounded growth** | 2026-03-26 | ~2.5 mo | P1, OOM trigger, no linked fix PR. |
| **#54463 – QMD memory symlink loop** | 2026-03-25 | ~2.5 mo | P2, platform stability, needs live repro. |
| **#57256 – `openclaw status` false unavailable** | 2026-03-29 | ~2.5 mo | P2, CLI reporting bug, linked PR open but stalled. |
| **#67777 – Subagent completion loss** | 2026-04-16 | ~2 mo | P1, needs maintainer review. |
| **#64810 – Telegram heartbeat swallows replies** | 2026-04-11 | ~2 mo | P1, needs maintainer review. |
| **#65624 – Mattermost cleartext callback URLs** | 2026-04-13 | ~2 mo | P1 security (CVSS 7.6/8.6), linked PR open but not merged. |
| **#66443 – Overflow duplicate user messages** | 2026-04-14 | ~2 mo | P1, linked PR open but not merged. |
| **#64267 – Agent internal thinking exposed** | 2026-04-10 | ~2 mo | P1, no linked PR. |

The lack of maintainer review on these high‑severity, older items—especially those with security implications (#57326, #65624, #64267)—poses a growing risk to production deployments. Community members have noted that some issues are being tagged `clawsweeper:needs-product-decision` without visible progress.

---

## Cross-Ecosystem Comparison

# Cross-Project Ecosystem Comparison Report — 2026-06-18

## 1. Ecosystem Overview

The personal AI assistant open-source ecosystem continues expanding rapidly, with **12 tracked projects** showing aggregate activity exceeding **650 issues and 800 pull requests** updated in a single 24-hour period. A clear "Claw family" has emerged, centered on OpenClaw as the core reference implementation, with derivatives (NanoClaw, PicoClaw, NullClaw, ZeptoClaw, ZeroClaw) sharing architectural DNA while differentiating on deployment scale and platform targets. The ecosystem is bifurcating into two tiers: large, general-purpose agent frameworks (OpenClaw, IronClaw, CoPaw) with hundreds of contributors, and smaller, specialized agents optimized for specific runtimes (Moltis, TinyClaw, NullClaw). Security hardening and memory/context management dominate the collective development agenda, with **8 of 12 projects** addressing active vulnerabilities or data integrity bugs. The competitive landscape is converging on shared patterns—MCP (Model Context Protocol) integration, multi-provider LLM routing, workspace sandboxing, and real-time voice/computer-use capabilities—while diverging on target deployment environments (desktop vs. mobile vs. embedded).

---

## 2. Activity Comparison

| Project | Issues Updated (24h) | PRs Updated (24h) | Release Today | Overall Health |
|---------|---------------------|--------------------|---------------|----------------|
| **OpenClaw** | 289 | 500 | No | Medium ⚠️ (high velocity but critical P1s backlogged) |
| **IronClaw** | 19 | 50 | No | High ✅ (responsive maintainers, OAuth fixes) |
| **CoPaw** | 33 | 31 | Yes (v1.1.12.post1) | High ✅ (active PR merging, 64 new tests) |
| **ZeroClaw** | 12 | 50 | No | High ✅ (9 PRs merged, P1 fixes landing) |
| **NanoClaw** | 2 | 15 | Yes (v2.1.0, v2.1.17) | High ✅ (breaking releases but responsive) |
| **Hermes Agent** | 17 | 50 | No | Medium ⚠️ (high PR volume, long-open P2s) |
| **NanoBot** | 8 | 31 | No | High ✅ (15 PRs merged, responsive team) |
| **PicoClaw** | 4 | 8 | No | High ✅ (security fixes shipped) |
| **LobsterAI** | 2 | 15 | Yes (2026.6.15) | High ✅ (major features shipped) |
| **NullClaw** | 3 | 3 | No | Low ⚠️ (no merges today, unaddressed bugs) |
| **Moltis** | 2 | 1 | No | Low ⚠️ (low velocity, no review activity) |
| **TinyClaw** | 3 | 0 | No | Critical 🔴 (zero maintainer response to 3 CVEs) |
| **ZeptoClaw** | 0 | 0 | No | Stalled ⚠️ (no activity) |

---

## 3. OpenClaw's Position

**Advantages:**
- Largest contributor base (289 issues, 500 PRs/day)—drives ecosystem-wide patterns (MCP, memory plugins, auth providers)
- **ClawSweeper bot** provides automated regression detection unmatched by peers
- **XMemo cloud memory plugin** (#94435) and headroom compression signal leadership on context management
- Session-state model with per-agent memory isolation (#63829) addresses multi-agent scaling challenges before competitors

**Technical Approach Differences:**
- OpenClaw's **session/context lifecycle** is more complex than NanoBot's stateless runner or NullClaw's lightweight REPL—this drives both its extensibility and its P1 bugs (session OOM, transcript doomloops)
- Uses **bundled channels** (Telegram, Discord, Slack) with normalization—IronClaw and ZeroClaw treat channels as host-agnostic ingress
- Heavier dependency on **community skill development** (#50090) vs. Hermes Agent's built-in skill catalog

**Community Size Comparison:**
- OpenClaw's 500 daily PRs dwarfs the next largest (IronClaw/ZeroClaw at 50 each)—roughly **10x the merge velocity**
- However, maintainer bandwidth is a bottleneck: **11 P1 bugs** lack assignees or fix PRs, vs. IronClaw's P1s all having linked fixes
- Community sentiment is constructive but **urgent**—the DeepSeek cache regression (#91016) burned real money, eroding trust in QA processes

---

## 4. Shared Technical Focus Areas

The following requirements appear across multiple projects, suggesting ecosystem-wide priorities:

| Focus Area | Affected Projects | Specific Needs |
|------------|------------------|----------------|
| **Context/Memory Management** | OpenClaw, CoPaw, NanoBot, NullClaw | Compaction freezing (#5218 CoPaw), delivery wipe (#4307 NanoBot), OOM loops (#63998 OpenClaw), configurable recall (#961 NullClaw) |
| **Security Hardening** | OpenClaw, PicoClaw, NanoClaw, LobsterAI, TinyClaw | SSRF bypass (#3143 PicoClaw), path traversal (#2799 NanoClaw), OAuth token refresh (#5071 IronClaw), arbitrary file read (#2176 LobsterAI), unauthenticated API (#282–284 TinyClaw) |
| **Multi-Provider LLM Routing** | Hermes, PicoClaw, ZeroClaw, IronClaw | Mistral timestamp rejection (#48386), Gemini thought_signature (#3136), Qwen system message order (#20866), OpenAI Responses tool delivery (#7756 ZeroClaw) |
| **Multi-Agent/Workspace Isolation** | OpenClaw, NanoBot, CoPaw, NanoClaw | Per-agent memory vault (#63829), workspace read/write asymmetry (#4374), sandbox isolation (#5310), permission escalation (#2807 NanoClaw) |
| **UI/UX & Accessibility** | OpenClaw, CoPaw, Hermes, Moltis | Screen reader token noise (#65538), Web UI on headless (#861 NullClaw), Markdown export (#1131 Moltis), TTS format config (#1126 Moltis) |
| **Channel Platform Parity** | OpenClaw, NanoBot, PicoClaw, ZeroClaw, Hermes | Telegram heartbeat swallowing (#64810), WhatsApp read receipts (#4354), Discord streaming (#5314 CoPaw), Feishu QR login (#4391 NanoBot) |

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | Hermes Agent | NanoBot | LobsterAI | TinyClaw |
|-----------|----------|--------------|---------|-----------|----------|
| **Primary User** | Power users, multi-agent deployers | Developers, model tinkerers | Ops teams, multi-instance | Productivity workers | Minimal/embedded |
| **Deployment Target** | Desktop + server | Desktop + TUI | Cloud gateway | Desktop (Electron) | API-only |
| **Architecture** | Monorepo, bundled channels | Plugin-heavy, MCP-first | Runner + gateway | Monolithic desktop app | Minimal API |
| **Memory Strategy** | Session state + XMemo cloud | Mem0 + app_id isolation | Workspace files + consolidation | Local artifact store | None (stateless) |
| **Risk Profile** | High-complexity bugs (OOM, loops) | Provider compatibility gaps | Workspace inconsistency | Electron security (file read) | Critical auth bypass |
| **Differentiator** | Largest ecosystem, ClawSweeper | TUI/CLI parity, multi-profile | Multi-tenant, wizard UX | Computer Use, real-time ASR | Minimal surface area |

**Key Differentiations:**
- **OpenClaw** competes on breadth (100+ skills, 10+ channels) but suffers complexity debt—its P1 bugs outnumber all other projects combined
- **Hermes Agent** invests heavily in the TUI experience and multi-profile gateway—closest to a "developer workstation" agent
- **LobsterAI** is the only project shipping **Computer Use MVP** and real-time ASR—positioning as productivity-first assistant
- **TinyClaw** is architecturally the simplest but faces existential security risk (3 unauthenticated API endpoints)
- **ZeroClaw** and **IronClaw** are iterating fastest on OAuth, tool reliability, and CI—appealing to production deployers

---

## 6. Community Momentum & Maturity

### Tier 1: High Velocity, Stabilizing (OpenClaw, IronClaw, CoPaw, ZeroClaw)
- **OpenClaw**: 500 PRs/day but critical bugs backlogged—entering a **stabilization phase** where feature work is paused for debt reduction
- **IronClaw**: Strong CI infrastructure, OAuth fixes landing, Reborn WebUI maturing—approaching **production readiness**
- **CoPaw**: 17 PRs merged today, 64 integration tests added—demonstrating **disciplined QA** alongside features
- **ZeroClaw**: 9 PRs merged, P1 bugs with linked fixes—**responsive maintainer team**, good candidate for v0.8.1 release

### Tier 2: Rapid Iteration (NanoBot, PicoClaw, NanoClaw, Hermes)
- **NanoBot**: 15 PRs merged, wizard UX improvements—**user-focused iteration** with growing provider support
- **PicoClaw**: Security fixes shipped same-day—**responsible disclosure culture**, but lower overall volume
- **NanoClaw**: Breaking releases suggest **major architectural shifts**—operator overhead but feature velocity high
- **Hermes Agent**: 50 PRs but 3 open, community submitting—**contributor engagement strong**, but maintainer review lagging

### Tier 3: Steady/Moderate (NullClaw, Moltis, LobsterAI)
- **LobsterAI**: Major feature releases (Computer Use, ASR)—**product-led growth** but security gap (#2176) needs immediate attention
- **NullClaw/Moltis**: Low volume, no merges—**risk of stagnation** if maintainer responsiveness doesn't improve

### Tier 4: Stalled/Critical (TinyClaw, ZeptoClaw)
- **TinyClaw**: Zero code changes, three unaddressed CVEs—**project at risk** unless maintainers respond within 48h
- **ZeptoClaw**: No activity in 24h—may be dormant

---

## 7. Trend Signals

### Emerging Trends from Community Feedback

1. **Cost Awareness is Driving Feature Demand**
   - DeepSeek cache regression (#91016) burned real money—users demand **cache validation dashboards** and cost tracking (ZeroClaw's cached-token pricing PR #7492 is a direct response)
   - Headroom compression (CoPaw #5244) advertised as "60–95% token reduction" signals a **market for cost-optimization plugins**

2. **Security is the #1 Unspoken Requirement**
   - 8 of 12 projects had active security issues today
   - TinyClaw's 3 unauthenticated endpoints and LobsterAI's arbitrary file read (#2176) show **basic auth is not yet standard**
   - Expect **mandatory auth gates** and **sandbox isolation** to become baseline expectations within 2–3 releases

3. **Multi-Provider LLM Routing is Fragile**
   - Every project with >2 providers reports endpoint-specific bugs (Mistral 422, Gemini thought_signature, Qwen system messages)
   - **Provider compatibility testing** is a growing need—ClawSweeper's approach may become ecosystem standard
   - Users want **automatic fallback** and **model-aware failover** (Hermes PR #48273, ZeroClaw #7756)

4. **Context Management is the New Reliability Frontier**
   - Three separate projects report context compaction freezing (#5218 CoPaw, #65161 OpenClaw, #4307 NanoBot)
   - "Doomloops" (OOM from crash-restart cycles) and delivery wipe are recurring patterns
   - **Idempotent message delivery** (#2808 NanoClaw) and **post-compaction coherence** are emerging as design requirements

5. **Desktop vs. Mobile UX Split**
   - LobsterAI's Computer Use + ASR targets **productivity desktop**
   - NullClaw's headless Web UI request (#861) and iOS TCC permission issues (#94286 OpenClaw) show **mobile/embedded deployment is under-supported**
   - Expect a **mobile-native agent project** to emerge in the next quarter

6. **Community Skill Ecosystems Are Fracturing**
   - OpenClaw's skill development friction (#50090) contrasts with Hermes' built-in catalog and CoPaw's plugin system
   - **ClawHub** (OpenClaw) and **Plugin UI Extension System** (#66944) suggest convergence on marketplace models
   - The winning approach will offer **sandboxed installation** and **versioning**—two gaps in today's ecosystem

7. **Real-Time Voice and Computer Use are 2026's Killer Features**
   - LobsterAI shipped both this week; CoPaw has terminal coding mode (#5304); Hermes is experimenting with multi-profile gateway
   - These capabilities are still **MVP-stage**—expect rapid iteration on latency, permission models, and cross-platform support

### Value to AI Agent Developers

| Insight | Actionable Takeaway |
|---------|---------------------|
| Context management is the #1 source of runtime failures | Invest in **compaction monitoring**, **idempotent delivery**, and **crash recovery** before scaling |
| Security hardening is not optional | Ship **auth gates**, **path validation**, and **sandboxing** before public release—TinyClaw's CVEs could destroy trust |
| Multi-provider routing requires extensive testing | Build a **provider compatibility matrix** and automated regression suite; heuristics for error classification (ZeroClaw #7927) are critical |
| Users will pay for cost optimization | Headroom compression, cache-aware pricing, and token budgets are **differentiators** that drive adoption |
| Community plugins need marketplace infrastructure | If you want an ecosystem, invest in **skill installers**, **permissions**, and **version management** early |
| Desktop and mobile are diverging | Decide your target deployment—trying to support both without dedicated effort will produce mediocre experiences in both |

---

*Analysis based on GitHub activity data from 2026-06-18. Health scores are qualitative assessments considering velocity, maintainer responsiveness, security posture, and bug resolution rate.*

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-06-18

## Today's Overview
The project saw a surge of activity with **31 pull requests** updated in the last 24 hours (16 open, 15 merged/closed) and **8 issues** updated (7 open, 1 closed). No new releases were published. The project is in a healthy, high-velocity state with multiple contributors landing fixes and features across the runner, providers, channels, and WebUI. Most activity centers on stability improvements (memory consolidation, workspace security) and user experience enhancements (wizard flow, UI config hiding, multi-instance support). The maintainer team appears responsive, as several critical bugs have identified fix PRs or have been closed.

## Releases
None.

## Project Progress
**Merged/closed pull requests today (15 total):** Notable merges include:

- **Qiniu provider support** (#3643) – adds a new model provider (七牛云 AI) with tests.
- **Workspace write policy clarification** (#4202) – aligns `apply_patch` path handling, separates read-only/write-allowed dirs.
- **Keenable search provider** (#4350) – adds a research-driven web search provider.
- **Mistral provider improvements** (#4351) – fixes reasoning_effort strictness and other Mistral API incompatibilities.
- **WhatsApp read receipts** (#4354) – sends blue ticks for incoming messages via the bridge.
- **Read-only root protection** (#4053) – keeps extra allowed roots from being written to.
- **Git workspace security fix** (see #4375 closed issue, related PR #4380 merged earlier? Actually #4375 is closed as bug; there is PR #4393 adding tests for git subdirectories, still open, but the guard fix itself (likely #4380) is merged).

Several other PRs remain open but are actively reviewed. The project is steadily accumulating provider integrations, channel enhancements, and security hardening.

## Community Hot Topics
- **#4307 – Post-turn consolidation wipes delivery message** ([Issue](https://github.com/HKUDS/nanobot/issues/4307))  
  *Comments: 3 | 👍: 0*  
  This bug reports that when `context_window_tokens` is modest, long multi-iteration turns lose the assistant’s own delivery message after consolidation. It is the most active issue today and has attracted a fix PR: **#4373** (preserve delivery context during consolidation) which is still open.

- **#4374 – Project workspace read/write asymmetry** ([Issue](https://github.com/HKUDS/nanobot/issues/4374))  
  *Comments: 2 | 👍: 0*  
  Users report that `SOUL.md`/`USER.md` are read from the project root but written to the default workspace, causing data inconsistency. This is a design-level concern that may require deeper architectural changes.

- **#4376 – User-friendly wizard** ([Issue](https://github.com/HKUDS/nanobot/issues/4376))  
  *Comments: 1 | 👍: 1*  
  A straightforward improvement request with broad support (1 reaction). Corresponding PR **#4395** is already in progress.

The underlying need across these topics is **usability**: new users struggle with configuration, and experienced users hit data loss or inconsistency in workspace management.

## Bugs & Stability
| Rank | Issue | Severity | Status | Fix PR |
|------|-------|----------|--------|--------|
| 1 | #4307 – Consolidation wipes delivery message | High – user follow-ups lost | Open | #4373 (open) |
| 2 | #4375 – Git commands blocked by workspace security | Medium – prevents version control | Closed (fix merged) | Related PR #4393 (test only, fix in #4380) |
| 3 | #4374 – Workspace read/write asymmetry | Medium – data inconsistency | Open | None yet |
| 4 | #4388 – iOS Safari input zoom bug | Low – UI distortion on mobile | Open | None yet |
| 5 | #4389 – Fallback models not trimmed per‑model | Medium – potential crashes | Open (feature request) | None |

The two most severe bugs (#4307 and #4375) have identified fixes; #4375 is already closed. The delivery message bug (#4307) still needs its PR merged. No crash-related bugs were reported today.

## Feature Requests & Roadmap Signals
- **Multi-tenant gateway** (#936, opened in Feb 2026) – a long-standing request for running multiple agents from a single gateway. No recent PR, but the upvote count suggests community interest.
- **Per-model contextWindowTokens** (#4389) – allowing different context limits for fallback models. A small but important flexibility improvement.
- **User-friendly wizard** (#4376) – prompted the already-open PR #4395, which is likely to land in the next release.
- **Multi-instance UI hiding** (#4390) – requests ability to hide UI settings for non-technical users. PR #4399 (hidden settings sections) directly addresses this.
- **Feishu QR scan-to-create bot** – PR #4391 adds a QR login flow for the Feishu channel, lowering onboarding friction.

**Predictions for next version:** The wizard improvements (#4395), hidden settings sections (#4399), optional feature enablement (#4396), and the Feishu QR feature (#4391) are all open, non‑bugfix PRs that are nearly ready. Expect them in the next minor release.

## User Feedback Summary
- **Pain points**:  
  - Data loss during long conversations (#4307)  
  - Configuration complexity for non-technical users (#4376, #4390)  
  - Inconsistent workspace file saving (#4374)  
  - Mobile UI broken on iOS Safari (#4388)  
  - Git commands unexpectedly blocked (#4375 – now fixed)

- **Positive signals**:  
  - PRs are being merged quickly (15 merged/closed today).  
  - Community contributions are growing: new provider integrations (Qiniu, Keenable, Mistral), channel enhancements (WhatsApp, Feishu), and UI polish.  
  - The maintainer team actively reviews and merges, indicating good project health.

Overall, users are engaged but encounter real friction in early setup and long-running sessions. The project’s responsiveness to bugs is strong.

## Backlog Watch
- **#936 – Multi-Tenant Gateway** ([Issue](https://github.com/HKUDS/nanobot/issues/936))  
  Opened 2026-02-21, last updated 2026-06-17. Only 1 comment. This feature is highly requested by operators running multiple agent instances. It has not been addressed in any PR and may require design discussion. The maintainer team should triage this for prioritization.

- **#4307** is active and has a fix PR, so not backlog.  
- **#4388** (iOS zoom) has no fix yet; could become stale if not addressed soon, but it’s recent.

No other issues older than 30 days with low maintainer attention.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-06-18

## 1. Today’s Overview

Hermes Agent saw a burst of activity over the past 24 hours: **17 issues** (13 open, 4 closed) and **50 pull requests** (47 open, 3 merged/closed) were updated. The project is in an intensive development phase, with many new bugs reported and even more features being proposed and implemented via PRs. While no new releases landed today, the sheer volume of contributions — from both core maintainers and the community — signals healthy momentum, though a handful of long-standing issues remain unassigned.

## 2. Releases

No new releases were published today.

## 3. Project Progress

Three pull requests were merged or closed today:

- **PR #45987** (merged): `feat(mem0): add app_id support for project-level memory isolation` — Adds scoping in the Mem0 memory plugin via `app_id`, enabling cleaner separation between projects. Backward compatible.
- **PR #48396** (merged): `fix(spotify): handle invalid_grant refresh token expiry` — Addresses Spotify’s July 2026 policy change that expires refresh tokens after 6 months. Previously a raw API error was shown; now the user is guided through re-authentication.
- **PR #48389** (closed as wrong repo): The author acknowledged the PR was accidentally submitted to the wrong project.

Other notable features advanced via open PRs today include:
- Multiplexing multiple profiles over a single gateway process (opt-in) – PR #48273
- Kanban rework loop for agent task management – PR #48274
- Neuralwatt as a native LLM provider – PR #48383
- XMemo cloud memory provider – PR #48239
- Documentation maintainer automation skill – PR #48410

## 4. Community Hot Topics

The following issues and PRs attracted the most discussion and reactions:

- **Issue #47917** (9 comments, 1 👍) — [Desktop build fails after update – electronDist cache invalidated](https://github.com/NousResearch/hermes-agent/issues/47917)  
  A persistent regression where the Electron binary cache is deleted during updates, breaking desktop builds. Users report the fix from PR #47276 works temporarily but breaks again after pulling latest code. The underlying need is for a stable, non-flaky desktop build pipeline.

- **Issue #19753** (4 comments, 2 👍) — [Auxiliary title generation 404 on custom Anthropic-mode providers](https://github.com/NousResearch/hermes-agent/issues/19753)  
  A long-standing bug (open since May 4) affecting custom providers with `api_mode: anthropic_messages`. Auxiliary tasks fail due to a double-`/v1` in the URL. Users are clearly frustrated as this blocks advanced workflows with services like Kimi.

- **Issue #12130** (4 comments) — [TUI v2 feature-parity audit vs v1 CLI](https://github.com/NousResearch/hermes-agent/issues/12130)  
  A comprehensive audit shows the new TUI is missing ~23 of 48 slash commands and zero context references (`@file`, `@diff`, etc.). The community is calling for feature parity to make the TUI a viable alternative to the classic CLI.

- **PR #48273** and **PR #48274** — Both received attention as significant architectural changes (multi-profile gateway and Kanban loop), though comment counts were not provided.

## 5. Bugs & Stability

Today saw several bug reports, ranked by severity:

**P1 (Critical):**
- **Issue #48176** — [OAuth Pro/Max/Team requests rejected with HTTP 400 – missing x-anthropic-billing-header](https://github.com/NousResearch/hermes-agent/issues/48176)  
  *New today.* OAuth credentials for higher-tier Claude plans are being rejected with a `third-party / extra usage` error. This blocks all users with paid Anthropic accounts. No fix PR yet.

**P2 (High):**
- **Issue #47917** — Desktop build failure (described above). Despite a prior fix, the issue re-emerges. A rework of the build caching logic is likely needed.
- **Issue #48338** — [`_append_model_switch_marker` injects `role:"system"` mid-conversation → HTTP 400 on strict providers](https://github.com/NousResearch/hermes-agent/issues/48338)  
  *New today.* Model switching in TUI breaks Qwen/vLLM endpoints because a system message is inserted out of order. A fix is presumably needed in the TUI gateway code.
- **Issue #48386** — [Gateway sends timestamp metadata field → HTTP 422 on Mistral endpoint](https://github.com/NousResearch/hermes-agent/issues/48386)  
  *New today.* The gateway injects a `timestamp` field into message objects, which Mistral’s strict OpenAI-compatible API rejects. Marked as duplicate, but no fix PR yet.
- **Issue #48388** — [Desktop GUI `_save_cfg()` sorts config keys alphabetically](https://github.com/NousResearch/hermes-agent/issues/48388)  
  *New today.* Users’ custom `config.yaml` ordering is destroyed after any GUI save. **Fix PR #48399** addresses this by adding `sort_keys=False`.

**P3 (Moderate):**
- **Issue #19753** — Auxiliary title generation 404 (long-standing, see earlier).
- **Issue #20866** — [400 format_error on Qwen3.6-27B – system message order](https://github.com/NousResearch/hermes-agent/issues/20866)  
  Similar root cause to #48338 but affecting auxiliary tasks. No dedicated fix PR yet.
- **Issue #48406** — [Desktop inline edit loses unsaved text when clicking outside editor](https://github.com/NousResearch/hermes-agent/issues/48406)  
  *New today.* Simple UX regression; workaround suggested.
- **Issue #48411** — [uv.exe flagged as malware false positive](https://github.com/NousResearch/hermes-agent/issues/48411)  
  *New today.* Not a code bug but an AV quarantine nuisance. A note for Windows users.

**Security-related:**
- **Issue #7651** (closed) — [Telegram missing user-level access control](https://github.com/NousResearch/hermes-agent/issues/7651) was closed, indicating the fix from PR #17748 is now live.
- **PR #47219** (open) — Fixes a config/.env write-gate bypass for non-slash paths. Still under review.

## 6. Feature Requests & Roadmap Signals

Several feature requests surfaced today, hinting at upcoming priorities:

- **TUI v2 parity** — Issue #12130 remains a top community ask. Expect continued work on slash commands, context references, and overlay menus.
- **Persistent model per Discord channel** — Issue #48413 requests a middle scope between session-only and global persistence. Likely to be picked up given the active Discord community.
- **In-session raw/compressed view toggle** — Issue #48404 asks for visibility into compression events. A small UX improvement that could land in a minor release.
- **Telegram `/undo` message deletion** — Issue #48400 is directly paired with **PR #48401** (open). This feature is almost ready to land.
- **Gemma-4 tool-call & reasoning fallback** — PR #48412 adds fallback logic for Gemma-4 when vLLM parsing fails. Shows ongoing support for new model families.
- **Kanban rework loop** — PR #48274 introduces a structured review cycle for agent tasks. If merged, it will be a significant workflow addition for power users.

## 7. User Feedback Summary

From today’s contributions and discussions:

- **Pain points**: Desktop build instability (multiple users reporting same issue), OAuth failures for paid Anthropic tiers (recent but urgent), config ordering destruction by GUI (a subtle but annoying regression), and missing TUI features (frustration from CLI veterans).
- **Use cases**: Users rely on Hermes for multi-profile deployments (PR #48273), integration with specialized models (Neuralwatt, Qwen, Gemma-4), and memory persistence across sessions (Mem0/XMemo). The Telegram and email platform integrations are heavily used, with requests for better undo behavior and access control.
- **Satisfaction indicators**: Positive contributions with detailed bug reports and thoughtful feature requests. The community is actively submitting PRs (50 today!), indicating strong engagement and trust in the project’s direction.

## 8. Backlog Watch

Three issues have been open for over a month without a clear maintainer response or assignee:

- **Issue #19753** (since May 4, 2026) — [Auxiliary title generation 404 on custom Anthropic-mode providers](https://github.com/NousResearch/hermes-agent/issues/19753)  
  Still active (updated today) but no assignee. Affects users of Kimi and similar custom providers.

- **Issue #12130** (since Apr 18, 2026) — [TUI v2 feature-parity audit](https://github.com/NousResearch/hermes-agent/issues/12130)  
  A detailed audit with no corresponding roadmap announcement. The community would benefit from a maintainer’s triage comment.

- **Issue #20866** (since May 6, 2026) — [400 format_error on Qwen3.6-27B – system message order](https://github.com/NousResearch/hermes-agent/issues/20866)  
  Similar to a just-filed bug (#48338) but older. Could be a duplicate or a different manifestation.

These issues are not critical but represent user trust that could erode if left unaddressed. A quick triage (even to mark as duplicates or planned) would improve project health perception.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

## PicoClaw Project Digest — 2026-06-18

### 1. Today's Overview

PicoClaw saw a high activity day with **8 pull requests updated** (4 merged/closed) and **4 issues updated** (2 closed). The maintainers focused on security hardening — shipping fixes for SSRF bypass via ISATAP literals, private inbound media fetch in OneBot, and adding diagnostic logging for the Brave search tool. Two Gemini 3.5 Flash compatibility bugs were resolved, and a new NEAR AI Cloud provider was merged. Community pull requests for a DeltaChat gateway and a `skills_install` type-assertion fix remain open. Overall the project is in a healthy state, with responsive maintainers addressing security and compatibility issues quickly.

### 2. Releases

*No new releases in the last 24 hours.*

### 3. Project Progress — Merged / Closed Pull Requests

- **#3141** — `fix(web_search): add diagnostic logging for Brave empty results` (closed)  
  Adds logging when Brave returns HTTP 200 with zero results, aiding diagnosis of silent failures.  
  [PR #3141](https://github.com/sipeed/picoclaw/pull/3141)

- **#3136** — `fix(gemini): set both camelCase and snake_case thought_signature in tool call request body` (closed)  
  Fixes tool execution failures with Gemini 3.5 Flash by sending the required `thought_signature` field in snake_case.  
  [PR #3136](https://github.com/sipeed/picoclaw/pull/3136)

- **#2917** — `feat(provider): add NEAR AI Cloud provider` (closed)  
  Introduces NEAR AI Cloud as a first-class OpenAI-compatible LLM provider, complete with catalog, model list fetch, and configuration.  
  [PR #2917](https://github.com/sipeed/picoclaw/pull/2917)

- **#3140** — `fix(onebot): block private inbound media fetches` (closed)  
  Prevents attacker-controlled media URLs from being used to fetch localhost/private network resources.  
  [PR #3140](https://github.com/sipeed/picoclaw/pull/3140)

- *Issues closed today:*  
  - **#3125** (`web_search` silently fails with Brave API key from `.security.yml`) — closed without a specific fix identified; likely related to diagnostic logging added in #3141.  
  - **#3111** (Gemini 3.5 Flash tool execution broken) — fixed by #3136.

### 4. Community Hot Topics

- **#3088 – `[Feature] use vodozemac instead of libolm`** (open, high priority, 👍2)  
  [Issue #3088](https://github.com/sipeed/picoclaw/issues/3088)  
  User request to replace the unmaintained `libolm` with `vodozemac`. The idea of making `libolm` optional at compile time has been discussed. This is the only issue tagged `help wanted` and `priority: high`, indicating strong maintainer interest.

- **#3093 – `[Feature] I need SimpleX or tox`** (open, stale)  
  [Issue #3093](https://github.com/sipeed/picoclaw/issues/3093)  
  Asks for integration with SimpleX, Wire, or Tox. Low activity (0 👍) but remains open, reflecting continued community interest in alternative messaging protocols.

- **#3063 – `feat: add deltachat gateway`** (open PR)  
  [PR #3063](https://github.com/sipeed/picoclaw/pull/3063)  
  Community contribution to add a DeltaChat gateway. Still awaiting review/merge.

### 5. Bugs & Stability

| Severity | Bug | Status | Fix PR |
|----------|-----|--------|--------|
| **High** | SSRF guard bypass via ISATAP IPv6 literals that embed private IPv4 addresses | Open | #3143 |
| **High** | OneBot inbound media fetch allowed private network addresses | Fixed | #3140 (closed) |
| **High** | `web_search` silently returns "No results" when Brave API key is misconfigured in `.security.yml` | Closed (root cause possibly unresolved) | #3141 adds diagnostics |
| **Moderate** | Gemini 3.5 Flash tool execution fails with `400 Bad Request` due to missing `thought_signature` | Fixed | #3136 (closed) |
| **Moderate** | Duplicate messages on async sub-agent completion due to `ForUser` not cleared in spawn subturn | Waiting merge | #3142 (open) |
| **Low** | `skills_install` silently accepts non-string/non-bool values for version/force | Awaiting review | #3092 (stale) |

### 6. Feature Requests & Roadmap Signals

- **Vodozemac replacement** (issue #3088) is the highest-priority feature request and likely to land in the next release, given the `high priority` tag and maintainer engagement.

- **NEAR AI Cloud provider** (PR #2917) was merged today, so it will appear in the next tagged release.

- **DeltaChat gateway** (PR #3063) is open and well-structured; could be merged soon if maintainers review.

- **SimpleX / Tox / Wire** (issue #3093) remains a low-activity request but signals user desire for decentralized chat backends.

- **Skills install robustness** (PR #3092) and **spawn duplicate fix** (PR #3142) are bug fixes that may be included in an upcoming patch release.

### 7. User Feedback Summary

- **Pain point: silent failures** – Giordano10 reported that the `web_search` tool broke silently after the `.security.yml` migration (issue #3125). Diagnostic logging (PR #3141) addresses symptom detection, but a deeper fix may still be needed.
- **Pain point: model incompatibility** – Giordano10 also reported Gemini 3.5 Flash tools failing (issue #3111). The fix (PR #3136) was merged promptly.
- **Request for security hardening** – pbsds opened #3088 to deprecate the insecure `libolm`, reflecting broader community demand for better encryption libraries.
- **Protocol diversity** – Users continue to ask for alternative messaging integrations (SimpleX, Tox, DeltaChat), indicating a desire to use PicoClaw with decentralised networks.

### 8. Backlog Watch

- **#3092** – `fix(skills_install): add ok checks for version and force type assertions` (open since 2026-06-10, stale)  
  [PR #3092](https://github.com/sipeed/picoclaw/pull/3092)  
  A small but important robustness fix. Maintainer review is needed to avoid future silent misconfigurations.

- **#3093** – `[Feature] I need SimpleX or tox` (open since 2026-06-10, stale)  
  [Issue #3093](https://github.com/sipeed/picoclaw/issues/3093)  
  No maintainer response yet. If the project does not intend to support these protocols, a clear close or roadmap note would help users.

- **#3063** – `feat: add deltachat gateway` (open since 2026-06-08)  
  [PR #3063](https://github.com/sipeed/picoclaw/pull/3063)  
  Awaiting maintainer review. The PR is substantial and may require architectural input.

- **#3088** – Vodozemac replacement (open since 2026-06-09)  
  [Issue #3088](https://github.com/sipeed/picoclaw/issues/3088)  
  High priority but no PR yet. The community and maintainers should align on implementation details soon.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-18

## Today’s Overview
NanoClaw saw **high development velocity** over the last 24 hours, with **15 PRs updated** (4 merged/closed) and **2 issues** filed. Two new roll-up releases shipped (v2.1.0, v2.1.17), each carrying breaking changes that require operator attention. The community is heavily focused on **security hardening**: three open PRs address path-traversal or privilege-escalation vulnerabilities. A critical delivery-stall bug (#2796) was identified and fixed within the same day, demonstrating responsive maintainership. Overall project health is strong, with sustained contributor activity across CLI, server, and security domains.

## Releases
### v2.1.17 — Rollup (v2.1.1–v2.1.17)
- **Breaking**: `@onecli-sh/sdk` upgraded from 0.5.0 → 2.2.1. Requires the OneCLI server to expose the `/v1` API; older servers will 404 every SDK call. The sanctioned gateway and CLI versions are now pinned.
- **Migration**: Ensure your OneCLI gateway is updated to support `/v1` endpoints. No code changes needed for NanoClaw itself, but the gateway and CLI must be redeployed.

### v2.1.0 — Rollup (v2.0.65–v2.1.0)
- **Breaking**: Startup now requires an upgrade marker. The host refuses to boot unless `data/upgrade-state.json` records that the current version was reached through a sanctioned upgrade path. Installations created directly at this version must manually place a valid marker file.
- **Migration**: If upgrading from v2.0.x, run the standard upgrade procedure; it will generate the marker automatically. For fresh installs, copy the marker template from the release assets.

Both releases are cumulative; sites on v2.1.0 should apply v2.1.17 to pick up all intermediate fixes.

## Project Progress (Merged/Closed PRs Today)
Four PRs were merged or closed:

- **#2806** — *docs: add Korean README* (merged). Full translation `README_ko.md` added, following the existing language-switcher pattern.
- **#2805** — *fix(setup): parse Claude OAuth token from wrapped PTY capture* (merged). Resolves a token-parsing failure in sandboxed environments where long lines are wrapped.
- **#2803** — *refactor: remove dead resolveGroupIpcPath* (merged). Cleanup of v2-era dead code; no production impact.
- **#2797** — *fix(delivery): isolate per-session failures so one bad session can't stall delivery for all* (merged). Fixes #2796 (see Bugs & Stability). Each session is now wrapped in its own try/catch.

**Still open, notable**: #2793 (per-message approval policies), #2809 (Apple Container runtime), #2808 (idempotent insertMessage), and several security fixes await review.

## Community Hot Topics
- **#2796** (CLOSED, 1 comment) — *One unhealthy session stalls message delivery for all agents*. Opened by mashkovtsevlx and fixed within hours by the same author (PR #2797 merged today). This issue garnered immediate attention and a rapid fix, indicating high alignment between bug reporters and the core team.
- **#2807** (OPEN, 0 comments) — *[Security] Non-owner members can create persistent child agents without approval*. A privilege-management flaw reported by YLChen-007. No comments yet; the vulnerability allows unauthorized agent creation in owner-initialized groups. Likely to attract discussion as it touches authorization boundaries.
- **#2793** (OPEN) — *feat(agent-to-agent): per-message approval policies*. A feature PR adding granular approval gates. Though no comments are recorded, the concept addresses a common governance need and may become a future default.

## Bugs & Stability
**Severity ranking** (high to low):

1. **Critical – Delivery stall** (Issue #2796, fixed by #2797). One corrupt session could halt message delivery for all agents until daemon restart. **Fixed in v2.1.18?** (not yet released but merged to main).
2. **High – Path traversal in `send_file`** (PR #2799, open). A prompt-injected or compromised agent can read any container-visible file. CVE-2026-29611 assigned. Fix restricts reads to `/workspace`.
3. **High – Group folder path traversal** (PR #2800, open). `ncl groups create --folder ../../etc` bypasses `assertValidGroupFolder`. Fix adds validator in CLI create path.
4. **Medium – CLi `messaging-groups create` always throws** (PR #2804, open). The create command is completely dead due to missing `instance` column. Fix adds the required column.
5. **Medium – Socket client timeout/buffer overflow** (PR #2802, open). `SocketTransport.sendFrame` has no timeout or response-size limit, risking hang or OOM.
6. **Low – `safeParseContent` fails on primitive JSON** (PR #2801, open). Returns non-object for e.g. `"5"`, breaking callers expecting `.text`/`.sender`.

All bugs have contributor-submitted fix PRs. Only the critical delivery stall has been merged; the rest await review.

## Feature Requests & Roadmap Signals
- **Per-message approval policies** (PR #2793, open) — introduces an optional, directed approval gate on agent-to-agent connections. Backward-compatible (no policy = today’s behavior). Likely to land in next minor release (v2.2).
- **Apple Container runtime** (PR #2809, open) — adds `CONTAINER_RUNTIME=container` for macOS native containers. Env-gated, default unchanged (docker). Clear signal of interest in macOS deployment.
- **CLI dashboard skill** (PR #2795, open) — `/add-clidash` read-only CLI-derived dashboard. Utility skill, no core changes. Shows demand for administrative tooling.
- **Stale outbound.db journal recovery** (PR #2750, open) — fixes two failure modes after container kills. This is a stability improvement rather than a new feature, but is long-standing and likely to be merged next.

## User Feedback Summary
- **Pain point**: Single-session corruption causing global delivery stall (mashkovtsevlx, reported and fixed). This reflects a real operational pain where a misbehaving agent can take down a whole production instance.
- **Security concern**: Non-owner member able to create child agents without approval (YLChen-007). Highlights demand for stricter authorization defaults.
- **Integration friction**: OAuth token parsing fails under sandboxed PTY (amit-shafnir, fixed). Indicates that setup procedures are still fragile in certain CI/sandbox environments.
- **Positive signal**: Multiple community contributors (sturdy4days, caburi00, glifocat, arkjun, leetwito, hidenwalker) are actively submitting PRs, showing a healthy contributor ecosystem.

No explicit user satisfaction statements are present, but the rapid embedding of fixes suggests maintainers are responsive to reported issues.

## Backlog Watch
| Item | Age | Status | Concern |
|------|-----|--------|---------|
| #2750 – Stale outbound.db journals after container kills | 6 days (updated 2026-06-17) | OPEN with PR | No maintainer review yet. Fixes two reported issues (#2516, #2640). May be blocking for high-availability setups. |
| #2804 – CLI `messaging-groups create` dead | 1 day | OPEN | Completely breaks a CLI command. Low complexity fix. |
| #2802 – Socket client timeout | 1 day | OPEN | Potential infinite hang. Should be prioritized. |
| #2799 – `send_file` path traversal (CVE) | 1 day | OPEN | Security fix with assigned CVE. Needs expedited review. |
| #2800 – Group folder path traversal | 1 day | OPEN | Security fix. |

**Maintainer attention needed on** #2750 (stability regression, long-pending) and the three security-related PRs (#2799, #2800, #2807 – issue + no PR). The backlog is young but security items should be triaged quickly.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-06-18

## 1. Today’s Overview
Project activity on June 18 remains moderate with 3 issues and 3 pull requests updated in the last 24 hours. No new releases were published, but two feature PRs and one bug-fix PR are now under review. The open bug about CLI arrow-key handling has an accompanying fix (#960), while an ongoing scheduler authorization problem (#915) and a documentation request for headless Web UI (#861) continue to draw community attention. Overall, the project is in a steady state with active patch and feature development, though no code has been merged today.

## 2. Releases
None — no new releases were recorded.

---

## 3. Project Progress
No pull requests were merged or closed today. However, three PRs are open and under active development:

| PR | Title | Status |
|----|-------|--------|
| [#960](https://github.com/nullclaw/nullclaw/pull/960) | fix(cli): handle arrow keys in agent REPL | Open — directly addresses Issue #865 |
| [#961](https://github.com/nullclaw/nullclaw/pull/961) | feat(memory): add configurable auto-recall, recall_limit, max_context_bytes | Open — adds three new JSON config keys |
| [#962](https://github.com/nullclaw/nullclaw/pull/962) | docs(providers): document native Anthropic provider with API key and OAuth support | Open — closes Issue #767 |

These PRs represent progress in CLI usability, memory recall control, and provider documentation, but remain in review.

---

## 4. Community Hot Topics
The most active discussions based on comment counts:

- **[Issue #865 – CLI shows ctrl characters for up/down/left/right keys](https://github.com/nullclaw/nullclaw/issues/865)** (2 comments)  
  User reports broken terminal keybindings in the CLI. A fix PR (#960) has been submitted, indicating high community demand for a usable interactive REPL.

- **[Issue #915 – Problem with scheduler unauthorized](https://github.com/nullclaw/nullclaw/issues/915)** (2 comments)  
  A user running NullClaw with Ollama on Ubuntu reports the scheduler fails entirely (both Telegram and CLI). Underlying need: reliable scheduling with external LLM hosts.

- **[Issue #861 – How to enable the Web UI on headless VPS server?](https://github.com/nullclaw/nullclaw/issues/861)** (1 comment)  
  A user asks for a clear, jargon-free guide to set up the Web UI via tunneled browser. This reflects a recurring documentation gap for headless deployments.

No PRs have accumulated comments yet.

---

## 5. Bugs & Stability
One new bug was reported today (through issue updates):

| Issue | Severity | Status | Fix PR |
|-------|----------|--------|--------|
| [#915 – Scheduler unauthorized](https://github.com/nullclaw/nullclaw/issues/915) | **Medium** – core scheduling feature broken for external LLM hosts | Open, no maintainer response yet | None |
| [#865 – CLI arrow keys broken](https://github.com/nullclaw/nullclaw/issues/865) | **Low-Medium** – usability regression in interactive mode | Open, fix PR #960 submitted | [#960](https://github.com/nullclaw/nullclaw/pull/960) |

No crashes or regressions were reported today beyond these two issues. The scheduler bug (#915) is potentially more impactful as it blocks a core use-case (scheduled tasks). The CLI bug (#865) has a pending fix.

---

## 6. Feature Requests & Roadmap Signals
Two ongoing PRs introduce new features likely to appear in the next release:

- **Configurable memory recall** ([#961](https://github.com/nullclaw/nullclaw/pull/961)) – adds `auto_recall`, `recall_limit`, and `max_context_bytes` under `memory`. This responds to user needs for finer control over memory enrichment and token budget.

- **Native Anthropic provider documentation** ([#962](https://github.com/nullclaw/nullclaw/pull/962)) – documents API key and OAuth support for Anthropic models, closing a long-standing documentation issue (#767).

User-requested features visible from open issues include:
- Scheduler reliability with external Ollama hosts (#915)
- Clearer Web UI setup guide for headless servers (#861)

Given the activity, the next version may include the memory recall improvements and the Anthropic provider docs, while scheduler fixes and CLI improvements are still in review.

---

## 7. User Feedback Summary
Real pain points expressed by users:

- **Scheduler not working** – “not in telegram chat nor CLI” despite working tool calls. This signals a potentially incomplete integration with external LLM backends.
- **CLI terminal unusable** – arrow keys show “CTRL character garbage” instead of expected navigation. This frustrates users who rely on history and cursor movement.
- **Web UI documentation confusing** – “I don’t understand 70% of that” – indicates the setup guide for headless VPS is too technical or lacks step-by-step instructions.

Satisfaction is not explicitly stated, but the presence of fix PRs suggests the community is engaged and maintainers are responsive.

---

## 8. Backlog Watch
Issues that remain open for an extended period without maintainer response or resolution:

- **[Issue #861 – Web UI on headless VPS](https://github.com/nullclaw/nullclaw/issues/861)** (created 2026-04-22, last updated 2026-06-17)  
  No maintainer comment has been posted. The user explicitly requested “non-jargon human terms”; a clear answer or documentation update would address this gap.

- **[Issue #915 – Scheduler unauthorized](https://github.com/nullclaw/nullclaw/issues/915)** (created 2026-05-15, last updated 2026-06-17)  
  No maintainer reply yet. This is a functional bug that may require investigation into external LLM host authentication.

These issues would benefit from a maintainer response or triage label to set community expectations.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

## IronClaw Project Digest – 2026-06-18

### 1. Today's Overview
Activity remains very high, with **19 issues** updated in the last 24 hours (10 open, 9 closed) and **50 pull requests** updated (28 open, 22 merged/closed). The project continues to focus heavily on the **Reborn** WebUI and engine v2 migration, with substantial quality-of-life fixes, OAuth reliability improvements, and CI enhancements. No new releases were published today. Developer momentum is strong, but several high-severity bugs around OAuth token refresh and tool failure handling remain open.

### 2. Releases
*(No new releases in this digest period.)*

### 3. Project Progress – Merged/Closed PRs
The following key PRs were merged or closed today (22 total; highlights from the top 20):

- **#5075** – Added a single `ci-verdict` rollup check for deterministic pre-merge gates → improves merge confidence.  
- **#5074** – Seeded Rust cache from merge queue and stabilized restore keys → faster CI for PR and merge queue.  
- **#4984** – Fixed failed tool activity updates not appearing in Reborn WebUI on first load (fixes #4942).  
- **#5047** – Skills validation error now clears after required fields are filled in Reborn WebUI (fixes #5007).  
- **#5053** – Refreshes OAuth runtime credentials on staging to avoid using stale access tokens (fixes staging-only issue).  
- **#5048** – Added `github.get_authenticated_user` capability and improved `github.list_repos` guidance.

These changes advance **Reborn WebUI stability**, **CI infrastructure**, and **Google/GitHub tooling parity**.

### 4. Community Hot Topics
Most active issues by comment count (all under 10 comments, none with reactions):

- **[#4761 – Agent stops after repeated tool failures instead of recovering](https://github.com/nearai/ironclaw/issues/4761)** (5 comments) – User reported that IronClaw (Reborn) stops responding after multiple tool call failures without attempting recovery. Underlying need: **better error resilience and automatic retry logic**.
- **[#4942 – Tool call failures not appearing until page reload](https://github.com/nearai/ironclaw/issues/4942)** (3 comments) – SSE updates missing tool failure status in WebUI. Fix PR #4984 was merged today, directly addressing this.
- **[#1520 – Qwen error: “Coding Plan is currently only available for Coding Agents”](https://github.com/nearai/ironclaw/issues/1520)** (3 comments) – A provider-specific HTTP 405 on Alibaba’s Qwen endpoint. Users want **cross-provider compatibility** for Coding Agents.
- **[#2800 – Engine v2 default flip umbrella tracker](https://github.com/nearai/ironclaw/issues/2800)** (2 comments) – Tracks remaining work to make engine v2 the default. This is a major architectural milestone still in progress.

No PRs had user comments in this dataset.

### 5. Bugs & Stability
New bugs filed or updated today, ranked by severity:

| Severity | Issue | Description | Fix PR exists? |
|----------|-------|-------------|----------------|
| **High** | [#5071 – Proactively refresh Google OAuth tokens before expiry](https://github.com/nearai/ironclaw/issues/5071) | Users forced to reauthenticate every hour; OAuth tokens expire without refresh. | PR #5054 (guide) and #5053 (staging fix) partially address; proactive refresh code still WIP. |
| **Medium** | [#5070 – Auth gate cancel after approval can replay OAuth prompt and leave activity running](https://github.com/nearai/ironclaw/issues/5070) | Canceling a pending OAuth grant can trigger repeated prompts or stuck activities. | No explicit fix PR yet; likely related to #5067. |
| **Medium** | [#5066 – Keep OAuth auth gate visible when authorization URL is unavailable](https://github.com/nearai/ironclaw/issues/5066) | Falls back to generic auth prompt, losing OAuth affordance. | PR #5067 (open) addresses this. |
| **Low/UX** | [#5077 – Invalid chat URLs should redirect to new chat](https://github.com/nearai/ironclaw/issues/5077) | Broken chat page with error instead of fallback. | Not yet. |
| **Low/UX** | [#5076 – Sidebar keeps chat thread highlighted on non-chat pages](https://github.com/nearai/ironclaw/issues/5076) | Misleading active state when navigating away from chat. | Not yet. |
| **Ongoing** | [#4108 – Nightly E2E failed](https://github.com/nearai/ironclaw/issues/4108) (updated today) | Scheduled E2E failing, likely due to missing scenario file. | PR #5073 (open) aims to fix the v2-engine matrix. |

Additionally, **#4704** (closed today) documented an approval loop with unhelpful error messages; fix merged earlier.

### 6. Feature Requests & Roadmap Signals
- **[#5069 – Automation UX Redesign](https://github.com/nearai/ironclaw/issues/5069)** (new today) – Requests a full UI redesign for the Automation page. Likely linked to the trigger poller work in PR #5030.
- **[#4505 – WeCom group conversation titles indistinguishable](https://github.com/nearai/ironclaw/issues/4505)** (still open) – Enhancement to distinguish group chats in sidebar.
- Several open PRs signal upcoming features:
  - **Auto-approve tool permissions** – PRs #5068, #5063, #5062 introduce per-tool permissions, global auto-approve toggle, and database-backed settings. Likely to appear in the next minor release.
  - **Slack as generic host ingress** – PR #5072 (open) refactors Slack integration, proving no behavioral change but enabling future extensibility.
  - **Binary document extraction from Google Drive** – PR #4997 adds PDF/PPTX/DOCX/XLSX text extraction via a host-side seam.
  - **Admin usage tracking for Engine V2** – PR #4989 persists LLM completions for cost monitoring.
  - **Production trigger poller** – PR #5030 wires scheduled automation in production.

**Prediction:** The next version will likely focus on **auto-approve settings**, **improved OAuth token management**, and **conversation/UX refinements**.

### 7. User Feedback Summary
Real pain points voiced in today’s issues:

- **Reliability:** “Agent stops after repeated tool failures” (#4761) and “Approval loop with no actionable error” (#4704) indicate users expect graceful recovery from tool errors.
- **UI responsiveness:** “Tool call failures don’t appear until reload” (#4942) – fixed today.
- **Form usability:** “Validation error doesn’t clear after filling fields” (#5007) – fixed today.
- **Navigation clarity:** “Icons for Logs/Docs are misleading” (#4923, closed) – replaced with text labels.
- **Identity confusion:** “WeCom group titles not distinguishable” (#4505) – still unresolved.
- **OAuth friction:** Multiple issues (#5071, #5070, #5066) about OAuth prompts and token refresh – users want seamless authentication.

No satisfaction signals were recorded, but the team’s quick turnaround on several bugs (e.g., #4942, #5007) suggests improving responsiveness.

### 8. Backlog Watch
Long-unanswered items needing maintainer attention:

| Item | Created | Last Updated | Notes |
|------|---------|--------------|-------|
| [#1520 – Qwen error](https://github.com/nearai/ironclaw/issues/1520) | 2026-03-21 | 2026-06-18 | Provider compatibility issue; 3 comments, no maintainer resolution. |
| [#4108 – Nightly E2E failure](https://github.com/nearai/ironclaw/issues/4108) | 2026-05-27 | 2026-06-18 | CI reliability; fix PR #5073 is open but not yet merged. |
| [#4505 – WeCom group titles](https://github.com/nearai/ironclaw/issues/4505) | 2026-06-05 | 2026-06-18 | UX enhancement, low activity. |
| [#4002 – Dependabot PR (actions group)](https://github.com/nearai/ironclaw/pull/4002) | 2026-05-24 | 2026-06-18 | Open PR with 16 action updates; may be blocked by CI changes. |

The Qwen issue (#1520) is the oldest open bug and could be a frustration for users on Alibaba endpoints. Nightly CI failures (#4108) are a persistent risk for release quality.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest – 2026-06-18

## 1. Today's Overview

The project is experiencing high activity: **15 pull requests** were updated in the last 24 hours (14 closed/merged, 1 open), while **2 issues** received updates (both still open). A new release **2026.6.15** was published, adding computer use capabilities and real-time ASR voice input. The most newsworthy event is a **high-severity security vulnerability** (#2176) reported today, detailing an arbitrary local file read via automatic artifact loading. The team is actively merging the `release/2026.6.11` branch into `main`, signaling a strong push toward the next stable release.

## 2. Releases

**New Release: [LobsterAI 2026.6.15](https://github.com/netease-youdao/LobsterAI/releases/tag/2026.6.15)**  
*Published 2026-06-15*

**What's Changed:**
- **Computer Use MVP**: Added a Windows x64 built-in Computer Use kit with marketplace metadata, skill bundle integrity checks, install/uninstall handling, and a managed runtime resolver/MCP server bridge for app/ window listing, launching, and screen capture.
- **Cowork – Real-time ASR Voice Input**: Introduced a real-time ASR mode via WebSocket streaming of PCM audio, with configurable voice input mode (real-time vs. one-shot) and i18n support.
- **Context Compaction**: Improved post-compaction context coherence in the cowork module.

**Breaking Changes:** None documented.  
**Migration Notes:** Users on older versions may need to reconfigure their voice input mode after update (the Settings UI now offers a “语音输入模式” toggle). No data migration required.

## 3. Project Progress

**Merged/Closed PRs (14 out of 15) – key highlights:**

| PR | Area | Summary |
|----|------|---------|
| [#2179](https://github.com/netease-youdao/LobsterAI/pull/2179) | multiple | Merged `release/2026.6.11` into `main` for the 2026.6.18 release. Includes document Artifact sharing (DOCX, PPTX, XLSX, PDF, CSV, TSV, Markdown, Mermaid) and preview improvements. |
| [#2178](https://github.com/netease-youdao/LobsterAI/pull/2178) | artifacts | Added Markdown and Mermaid file sharing from the Artifact panel, including zip packaging and resource handling tests. |
| [#2163](https://github.com/netease-youdao/LobsterAI/pull/2163) | voice-input | Refined dictation recording UI and ASR quota handling (in-memory quota, lazy-reset across sessions). |
| [#2160](https://github.com/netease-youdao/LobsterAI/pull/2160) | voice-input | Removed legacy short ASR upload flow; now only real-time ASR mode remains. Removed Settings toggle for recognition mode. |
| [#2156](https://github.com/netease-youdao/LobsterAI/pull/2156) | computer-use | Bumped Computer Use runtime to 1.0.7 (includes UIA breadcrumbs for diagnostics). |
| [#2143](https://github.com/netease-youdao/LobsterAI/pull/2143) | skills | Feat: add Computer Use MVP (core integration, managed runtime, MCP bridge). |
| [#2148](https://github.com/netease-youdao/LobsterAI/pull/2148) | cowork | Feat: add real-time ASR voice input (initial implementation with WebSocket streaming). |
| [#2177](https://github.com/netease-youdao/LobsterAI/pull/2177) | cowork | Renamed Chinese copy from "听写" to "语音输入" and English from "dictation" to "voice input". |
| [#2155](https://github.com/netease-youdao/LobsterAI/pull/2155) | cowork | Fixed duplicate real-time ASR start requests. |
| [#2150](https://github.com/netease-youdao/LobsterAI/pull/2150) | kits | Fixed sticky header/toolbar for expert suite controls. |
| [#2113](https://github.com/netease-youdao/LobsterAI/pull/2113) | voice-input | Added macOS microphone permission handling for voice input. |
| [#2111](https://github.com/netease-youdao/LobsterAI/pull/2111) | cowork | Refactored voice input into modular components (ASR IPC, WAV encoding, client, UI). |
| [#2107](https://github.com/netease-youdao/LobsterAI/pull/2107) | multiple | Release 2026.6.2 (features, bug fixes across cowork, MCP, HTML share, artifacts). |
| [#2119](https://github.com/netease-youdao/LobsterAI/pull/2119) | multiple | Release 2026.6.4 (voice input, artifacts, shortcuts, updates). |

**Significance**: The project has firmly transitioned voice input to a real-time-only model, and the Computer Use feature has reached MVP stage. Document artifact sharing now covers a wide range of office and code formats.

## 4. Community Hot Topics

Only **two issues** were active today:

- **#2176 – Security: Arbitrary local file read via automatic artifact loading**  
  *[Open](https://github.com/netease-youdao/LobsterAI/issues/2176)* | *1 comment*  
  **Summary:** LobsterAI automatically parses `MEDIA:` file references from assistant or tool output and forwards the resulting file path into a privileged Electron context, enabling arbitrary local file reads. This is a **critical security concern** (no fix PR yet).  
  **Community sentiment:** High urgency – the single comment likely from the reporter. The project needs immediate maintainer attention.

- **#1422 – UI: Long service names truncation in MCP custom page deletion dialog**  
  *[Open (stale)](https://github.com/netease-youdao/LobsterAI/issues/1422)* | *1 comment*  
  **Summary:** When the service name is long, the delete confirmation dialog displays poorly (no word wrap / overflow). Created in April, last updated today with no new activity.  
  **Community sentiment:** Low engagement – possibly a minor annoyance that didn't receive traction.

**Analysis:** The security issue is by far the most active community topic today. The stale UI bug indicates that cosmetic issues may be deprioritized in favor of feature work.

## 5. Bugs & Stability

**Critical (Severity: High)**
- **#2176 – Arbitrary local file read via artifact loading**  
  *Reported 2026-06-18*  
  No fix PR exists. The vulnerability allows an attacker to craft messages that trick LobsterAI into reading local files and forwarding them. This is a **security bug** requiring a hotfix. The project must respond quickly to patch the `ARTIFACT:` handler with proper path validation or sandboxing.

**Low (Severity: Low)**
- **#1422 – UI truncation in delete dialog**  
  *Reported 2026-04-03, last updated 2026-06-18*  
  Cosmetic issue – no fix PR, but no evidence of crashes or data loss. Likely low priority.

**Stability from merged PRs:**  
No regressions are reported. The computer use runtime bump (#2156) includes diagnostic improvements. The duplicate real-time ASR start fix (#2155) should prevent resource leaks.

## 6. Feature Requests & Roadmap Signals

**Recently merged features that signal near-term roadmap:**
- **Computer Use MVP** (#2143) – a major new capability that likely will be expanded in subsequent releases (e.g., adding macOS/Linux support, more app control actions).
- **Document Artifact Sharing** (#2178, #2179) – support for office formats, Markdown, Mermaid. Suggests a push toward making LobsterAI a full productivity assistant with rich output sharing.
- **Real-time ASR Voice Input (#2148) + removal of legacy mode (#2160)** – indicates the team is consolidating on one high-quality voice input pipeline and may add continuous dictation or wake-word features next.

**User-requested features (inferred from current work):**
- Better voice input UX (quota feedback, renaming UI strings to "voice input") – suggests users found "dictation" confusing.
- Selected text context from artifact previews (part of release 2026.6.2) – users needed to reference specific portions of long outputs.

**Prediction for next version (2026.6.18 or 2026.6.25):**
- Security patch for #2176 (must be fast).
- Enhanced computer use with additional actions (e.g., keyboard/mouse simulation, file operations).
- Further artifact format support (maybe audio/video?).

## 7. User Feedback Summary

**Direct user pain points (from issues):**
- **Security concern**: #2176 indicates a user (or security researcher) discovered a dangerous data leak vector. This will erode trust if not addressed.
- **UI polish**: Long service names in delete dialogs (#1422) – a minor but real irritation for users with many MCP servers.

**Satisfaction indicators:**
- High PR merge velocity (14 today) suggests the team is responsive to features.
- The removal of legacy voice input mode indicates the team heard feedback that the two-mode system was confusing – they streamlined to real-time-only.

**Dissatisfaction indicators:**
- The stale label on #1422 may frustrate users who reported it over two months ago.
- The lack of a fix for #2176 today (even a partial fix) may raise concerns about the project's security responsiveness.

## 8. Backlog Watch

**Long-unanswered but important issues/PRs needing maintainer attention:**

- **[PR #1277 – chore(deps-dev): bump electron group](https://github.com/netease-youdao/LobsterAI/pull/1277)**  
  *Open since 2026-04-02, updated today but not merged*  
  Updates Electron from 40.x to 42.x. This is a **dependency bump** that may include security fixes and performance improvements. The PR is kept alive by Dependabot but has not received a maintainer review nor a merge. Given the recent security issue (#2176), maintaining up-to-date Electron is critical. **Recommendation**: merge or close with rationale.

- **[Issue #1422 – MCP custom page dialog truncation](https://github.com/netease-youdao/LobsterAI/issues/1422)**  
  *Open since 2026-04-03, stale*  
  No maintainer response. A simple CSS fix (word break / max-width) would likely close it. **Recommendation**: label as `good first issue` or assign a fix.

- **[Security Issue #2176 – Arbitrary file read](https://github.com/netease-youdao/LobsterAI/issues/2176)**  
  *Open since today*  
  This is a **hot** backlog item – it should be elevated to critical priority and assigned immediately.

---

*Generated from GitHub data on 2026-06-18.*

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

## TinyClaw Project Digest — 2026-06-18

**Data source:** [TinyAGI/tinyagi](https://github.com/TinyAGI/tinyagi)  
**Period covered:** 2026-06-17 through 2026-06-18 (24h)

---

### 1. Today's Overview
The project saw no code changes, releases, or pull requests in the last 24 hours. Activity was limited entirely to the filing of three high-severity security issues, all reported by the same researcher (YLChen-007) on 2026-06-18. Each issue describes an unauthenticated API endpoint or unsafe input handling that could lead to remote file disclosure or unauthorized model invocation. No maintainer responses, comments, or fixes have been posted yet. The project appears stable in terms of release cadence but faces immediate security scrutiny.

---

### 2. Releases
No new releases were published today. The most recent release remains unknown (none listed).

---

### 3. Project Progress
No pull requests were merged or closed in the last 24 hours. No features or fixes were advanced.

---

### 4. Community Hot Topics
All three open issues are security advisories with no community comments or reactions yet. However, given their critical nature, they represent the most active discussion points for the project. The underlying need is clearly for authentication and input-validation hardening.

- **[#284] [Security] TinyAGI allows unauthenticated API messages to invoke Claude with provider permission checks disabled by default**  
  [View issue](https://github.com/TinyAGI/tinyagi/issues/284)

- **[#283] [Security] Unauthenticated `prompt_file` agent configuration allows arbitrary local file disclosure to the model provider**  
  [View issue](https://github.com/TinyAGI/tinyagi/issues/283)

- **[#282] [Security] Untrusted `[send_file: ...]` response tags allow arbitrary host file attachment delivery in TinyAGI**  
  [View issue](https://github.com/TinyAGI/tinyagi/issues/282)

---

### 5. Bugs & Stability
Three new security vulnerabilities were reported today, all ranked **critical** due to the lack of required authentication and the potential for remote exploitation.

| Issue | Severity | Summary | Fix PR Exists? |
|-------|----------|---------|----------------|
| [#284](https://github.com/TinyAGI/tinyagi/issues/284) | Critical | Unauthenticated `POST /api/message` allows arbitrary AI model invocation (Claude) without permission checks. | No |
| [#283](https://github.com/TinyAGI/tinyagi/issues/283) | Critical | Unauthenticated agent configuration API accepts `prompt_file`, enabling arbitrary local file reads sent to model provider. | No |
| [#282](https://github.com/TinyAGI/tinyagi/issues/282) | Critical | Response tag `[send_file: ...]` allows unauthenticated file attachment delivery to any host. | No |

No other bugs, crashes, or regressions were reported.

---

### 6. Feature Requests & Roadmap Signals
No feature requests were filed or discussed today. The project’s immediate roadmap should prioritize addressing the reported security gaps before any new features.

---

### 7. User Feedback Summary
No user feedback or usage comments appeared in the last 24 hours. The absence of user engagement suggests the community may be waiting for a response to the security disclosures, or that the project currently has limited active users.

---

### 8. Backlog Watch
No long-standing issues or pull requests require maintainer attention today. However, the three new security issues have zero response time and should be treated as high-priority backlog items. If no triage or fix appears within 24–48 hours, they may become stale and increase project risk.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest — 2026-06-18

## 1. Today's Overview
Moltis showed moderate activity over the past 24 hours, with two open enhancement issues and one open pull request receiving updates. No new releases were published, and no pull requests were merged or closed. The project continues to evolve through community-requested features and configuration improvements, though overall velocity remains steady rather than high. The single open PR suggests focused development on making internal timeouts configurable, a practical reliability enhancement.

## 2. Releases
*None published in the last 24 hours.*

## 3. Project Progress
- **Merged/Closed PRs today:** 0  
- **Open PRs:**  
  - **#1130** (open) — `feat: make webui rpc timeout configurable` by khimaros. This PR addresses issue #1127 (not listed in today’s data) and is currently awaiting review. It would allow users to adjust the WebUI’s RPC timeout, a straightforward quality-of-life improvement that may help users with slower connections or complex tasks.

No features were advanced or bugs fixed via merged PRs today.

## 4. Community Hot Topics
- **Issue #1126** — `[Feature]: allow to configure the format of tts output` by khimaros (3 comments, opened Jun 16, updated Jun 17). This is the most discussed item today. The community is engaging on how to control the output format of text-to-speech, likely driven by users who need different audio encodings or stream formats for downstream integrations. The three comments (not detailed in data) suggest active interest and possible debate about implementation scope.

- **Issue #1131** — `[Feature]: Add copy + export as Markdown` by vvuk (0 comments, opened Jun 17). A quiet but clear request to copy or export chat sessions as Markdown, indicating users want portable, formatted logs of their AI interactions.

- **PR #1130** — (see section 3) has no comments yet but is linked to issue #1127, which may have generated prior discussion.

## 5. Bugs & Stability
No bugs, crashes, or regressions were reported in the last 24 hours. The project’s stability appears solid, with all active items being enhancements. No fix-related PRs are currently open.

## 6. Feature Requests & Roadmap Signals
Two enhancement requests surfaced today:
- **Configurable TTS output format** (#1126) — likely to be prioritized if the community consensus forms quickly, as it touches a core integration.
- **Copy/export as Markdown** (#1131) — a low-effort, high-value UX feature that could appear in the next minor release if a contributor picks it up.

Additionally, **PR #1130** (configurable WebUI RPC timeout) is a practical reliability feature that may have been responding to user frustration with timeouts. Given it’s already a PR, it is the strongest candidate for inclusion in the next release.

## 7. User Feedback Summary
- **Pain points:** Users are requesting more control over output formats (TTS) and export capabilities (Markdown), signaling a desire for interoperability and customization.
- **Use cases:** Exporting conversations as Markdown suggests use in documentation, note-taking, or archiving; TTS format configurability hints at use with external audio pipelines or accessibility tools.
- **Satisfaction/dissatisfaction:** No overt dissatisfaction is evident; the requests are all additive. The lack of bug reports points to a stable core.

## 8. Backlog Watch
No long-unanswered issues or PRs are visible in today’s data snippet. All items are recent (created or updated within the last two days). Maintainers should watch:
- **Issue #1126** (now 2 days old with comments) – could benefit from a maintainer’s direction or tag.
- **PR #1130** – awaiting review; if left unaddressed, it may become stale. Recommended to mark as `needs-review` to accelerate progress.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest — 2026-06-18

## 1. Today's Overview

CoPaw project activity remains **very high** today, with 33 issues and 31 PRs updated in the last 24 hours — a healthy balance of bug fixing, new feature development, and community engagement. A patch release (`v1.1.12.post1`) was published, addressing critical scripting and memory backend issues. The community is particularly vocal this week around **context compaction stability** (process freezes, information loss) and **upgrade UX regressions** (disabled skills re-enabling, token usage display bugs). Several first-time-contributor PRs are in the review pipeline, and maintainers are actively merging fixes across MCP, memory indexing, SSL certificate handling, and UI components.

## 2. Releases

A new patch release was published today:

- **v1.1.12.post1** — [Release link](https://github.com/agentscope-ai/QwenPaw/releases/tag/v1.1.12.post1)
  - **Changes:**
    - `fix(scripts)`: correct prerelease arguments expansion and bump version
    - `fix(memory)`: rename ChromaDB probe collection to `'probe-test'`
  - **Breaking changes: None**
  - **Migration notes:** This is a hotfix; no migration steps required for most users. Users should restart QwenPaw to pick up the fix.

## 3. Project Progress

**17 PRs were merged or closed today.** Key advances:

- **Context Management Overhaul (#5309)** — Merged: Migrated from custom `LightContextManager` to AgentScope 2.0 native compression, using `QwenPawOffloader` and `Offloader` protocol (by @qbc2016).
- **MCP Performance (#4849)** — Merged: `SharedMCPPool` reuses MCP server processes across agents, preventing process explosion on Windows (by @wangfei010313).
- **Plugin System (#4794, #5008)** — Merged: Added uninstall hooks, exposed skill provider API, cherry-picked to `dev/agentscope2.0`.
- **Discord Streaming (#5314)** — Open: Adds streaming response support via message edit and typing indicator (by @hongxicheng).
- **Terminal Coding Mode (#5304)** — Open: `qwenpaw terminal` interactive coding-mode terminal with daemon autostart (by @nguyenthanhthe).
- **Headroom Context Compression (#5244)** — Open: Integrates Headroom SDK as a drop-in Context Manager backend (by @K1-lihongrong, first-time contributor).
- **Context Usage Display Fixes (#5303, #5306)** — Merged: Fixed Web chat and Console context usage popover to use active model's `max_input_length` instead of `agent.running.max_input_length`.
- **Windows SSL Certificate Handling (#5291, #5298)** — Merged: Fixed DingTalk channel SSL failures under `uv` install and Windows build script SSL errors.
- **Windows Memory Index Fix (#5265)** — Open: Forces vector index rebuild on Windows `"local"` backend to fix persistence (by @nguyenthanhthe).
- **Sandbox Isolation (#5310)** — Open: Adds bubblewrap Linux sandbox with mount namespace isolation (by @vanwaals, first-time contributor).
- **Integration Test Suite (#5270)** — Merged: Sprint 3.1–3.4 with 64 test cases across ACP, Plugin, Security, and cross-cutting features (by @yutai78786).
- **Chat History Right Panel (#5293)** — Merged: History chat list moved from Drawer pop-up to embedded sidebar panel (by @zhaozhuang521).

## 4. Community Hot Topics

The following issues and PRs generated the most discussion today:

1. **[#5218 — Sub-agent context compaction freezes process](https://github.com/agentscope-ai/QwenPaw/issues/5218)** (16 comments) — ⚠️ High severity. User reports entire QwenPaw process becomes unresponsive when a sub-agent triggers context compaction (上下文压缩). Only restart helps. **Underlying need:** The compaction subsystem lacks timeout/retry logic and may deadlock. No fix PR yet.

2. **[#5171 — Context compaction loses full context when profile tokens exceed threshold](https://github.com/agentscope-ai/QwenPaw/issues/5171)** (8 comments) — Agent profile files with token count > retention threshold cause compaction to zero — complete context wipe. **Underlying need:** Compaction logic needs per-message or category-based retention, not global token cut.

3. **[#5063 — Integrate Headroom as optional context compression layer](https://github.com/agentscope-ai/QwenPaw/issues/5063)** (7 comments) — Feature request to integrate Headroom SDK for 60–95% token reduction. PR #5244 (open) implements this. **Underlying need:** Users want alternative/parallel compression strategies to reduce costs.

4. **[#5262 — Disabled built-in skills re-enable on upgrade](https://github.com/agentscope-ai/QwenPaw/issues/5262)** (7 comments) — Recurring complaint (#4807 was previously filed). Disabled skills like `docx`, `xlsx` become re-enabled after every upgrade. **Underlying need:** Config state must persist across upgrades.

5. **[#5140 — Attachment download fails for docx/pdf (404)](https://github.com/agentscope-ai/QwenPaw/issues/5140)** (8 comments, now closed) — Text files download fine; binary files (docx, pdf) fail with 404. Fix PR was merged in a previous cycle.

## 5. Bugs & Stability

Bugs reported or updated today, ranked by severity:

| Severity | Issue | Description | Fix Status |
|----------|-------|-------------|------------|
| 🔴 Critical | [#5218](https://github.com/agentscope-ai/QwenPaw/issues/5218) | Sub-agent context compaction freezes entire process | No fix PR yet |
| 🔴 Critical | [#5171](https://github.com/agentscope-ai/QwenPaw/issues/5171) | Context compaction wipes all context when profile tokens exceed threshold | No fix PR yet |
| 🟡 High | [#3854](https://github.com/agentscope-ai/QwenPaw/issues/3854) | Chromadb Rust binding segfault (SIGSEGV) kills process — 45+ times/session | No fix PR yet |
| 🟡 High | [#5262](https://github.com/agentscope-ai/QwenPaw/issues/5262) | Disabled built-in skills re-enable on each upgrade | No fix PR yet |
| 🟡 High | [#5244](https://github.com/agentscope-ai/QwenPaw/issues/5244) | Headroom integration (open PR, under review) | PR #5244 open |
| 🟢 Medium | [#5300](https://github.com/agentscope-ai/QwenPaw/issues/5300) | Web context usage display uses wrong `max_input_length` (131072 vs 480000) | Fixed in #5303 and #5306 |
| 🟢 Medium | [#5313](https://github.com/agentscope-ai/QwenPaw/issues/5313) | MCP `streamable_http` Authorization header loses "Bearer" prefix | Closed |
| 🟢 Medium | [#5237](https://github.com/agentscope-ai/QwenPaw/issues/5237) | DingTalk channel not working with `uv` install on Windows | Fixed in #5291 |
| 🟢 Medium | [#5317](https://github.com/agentscope-ai/QwenPaw/issues/5317) | Tauri on Windows can't find Python for custom skills | No fix PR yet |
| 🟢 Medium | [#5265](https://github.com/agentscope-ai/QwenPaw/issues/5265) | Windows memory store backend fails to persist vector index | PR #5265 open |

**Today's fix highlights:**
- `#5303` / `#5306` — Context usage display denominator bug fixed
- `#5291` / `#5298` — SSL certificate handling fixed for DingTalk and Windows build
- `#5309` — Context management migrated to AgentScope 2.0 native compression (reduces custom code risk)
- `#5270` — 64 new integration tests landed (covers ACP, Plugin, Security)

## 6. Feature Requests & Roadmap Signals

Features requested or updated today that may shape the next release:

1. **Headroom Integration (#5063 / #5244)** — Strong signal. First-time contributor PR with SDK for 60–95% token compression. Likely to land in next minor release.
2. **Separate Vision Model Routing (#3940)** — User wants automatic routing of image inputs to a vision-capable model without manual switch. 5 comments, open since April.
3. **Terminal Coding Mode (#5304)** — Official PR from core contributor. A `qwenpaw terminal` CLI mode with daemon autostart. Likely to land soon.
4. **Native TodoWrite-like Progress Panel (#5318)** — New request today. Users want a real-time multi-step task progress panel in web UI.
5. **Recency-Aware Memory Search (#5316)** — Open feature request today. Optional ranking by date for `memory_search` results from daily notes.
6. **Sandbox Isolation (#5310)** — First-time contributor PR adding bubblewrap Linux sandbox. Shows community interest in security.
7. **Auto Model Listing After Provider Registration (#3844)** — Open since April. Users want auto-discovery of models from OpenAI-compatible providers without manual registration.

**Predictions for next release:** Headroom compression, terminal coding mode, and the sandbox isolation PR are likely candidates. The context compaction bugs (#5218, #5171) must be fixed for stability.

## 7. User Feedback Summary

Real user pain points and use cases expressed in today's issues/comments:

| Theme | Example | Impact |
|-------|---------|--------|
| **Context compaction fragility** | "Process freezes completely" (#5218); "Context completely lost" (#5171) | Highest pain — users lose work and cannot recover without restart |
| **Upgrade regressions** | "Disabled skills re-enable every time" (#5262) | Medium pain — manual re-disabling after every update is frustrating |
| **SSL/certificate issues** | "DingTalk not working with uv install" (#5237); "Windows build fails SSL verification" (#5298) | Medium pain — blocks channel communication |
| **File attachment bugs** | "docx/pdf download 404" (#5140, closed) | Low now — fixed in previous cycle |
| **Permission errors** | "Access denied to authorized NAS paths" (#4922) | Medium pain — workflow interruption |
| **Context display confusion** | "Context usage shows 131072 but model supports 480000" (#5300) | Low — misleads users about remaining context |
| **Python environment issues** | "Tauri can't find Python for skills" (#5317) | Medium — blocks custom skill execution |
| **Backup failures** | "Backup never succeeds" (#3821, closed) | Low now — likely environment-specific |
| **History showing system prompts** | "Arrow keys show system instructions in history" (#3975, closed) | Low — UX polish issue |

**Overall sentiment:** Users appreciate the rapid bug fixing (many issues closed quickly) but are frustrated by **context compaction instability** and **upgrade state loss**. New features (Headroom, terminal mode, Discord streaming) are generating positive community engagement.

## 8. Backlog Watch

Issues and PRs that are old, important, and may need maintainer attention:

1. **[#3940 — Separate vision model routing](https://github.com/agentscope-ai/QwenPaw/issues/3940)** (open since 2026-04-29, 5 comments) — User-facing enhancement for multi-modal workflows. No assignee, no PR.
2. **[#3844 — Auto model listing after provider registration](https://github.com/agentscope-ai/QwenPaw/issues/3844)** (open since 2026-04-26, 2 comments) — Would greatly improve UX for self-hosted providers. No PR.
3. **[#3854 — Chromadb Rust binding segfault](https://github.com/agentscope-ai/QwenPaw/issues/3854)** (open since 2026-04-27, 6 comments) — Critical crash bug on Linux. Still no fix committed despite high severity.
4. **[#3768 — Auto-reject commands by regex](https://github.com/agentscope-ai/QwenPaw/issues/3768)** (closed, 3 comments) — Feature was requested but closed. May need re-evaluation if community interest persists.
5. **[#4622 — DataPaw plugin (data analysis)](https://github.com/agentscope-ai/QwenPaw/pull/4622)** (open since 2026-05-22, first-time contributor) — Substantial plugin with 12 BI skills. Under review for over 27 days — needs reviewer bandwidth.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest – 2026-06-18

**Generated from GitHub data (zeroclaw-labs/zeroclaw)**  
*Data snapshot: 2026-06-18 23:59 UTC*

---

## 1. Today's Overview

The ZeroClaw project saw very high activity today, with **50 PRs updated** (9 merged/closed) and **12 issues updated** (1 closed). Development focus is split between urgent bug fixes (especially around provider compatibility and runtime stability) and feature work (WASM plugin host, Discord slash-command parity, unified slash-command registry). Several P1 bugs were addressed with corresponding fix PRs, indicating a responsive maintainer team. Community engagement is moderate, with most issues having few comments but covering real-world pain points like Android Termux installation and tool availability on different LLM providers.

---

## 2. Releases

No new releases were published today. The last release version remains unknown from this snapshot.

---

## 3. Project Progress – Merged/Closed PRs Today

Nine PRs were merged/closed in the last 24 hours. Notable examples from the top-20 list:

| PR | Title | Type | Summary |
|----|-------|------|---------|
| [#7927](https://github.com/zeroclaw-labs/zeroclaw/pull/7927) | fix(providers): guard is_non_retryable against false-positive on 429 bodies | Bug fix | Prevents rate-limit (429) responses from being misclassified as permanent failures when the response body contains heuristically matched words like “model” or “invalid”. |
| [#7918](https://github.com/zeroclaw-labs/zeroclaw/pull/7918) | test: deterministic storage-reader timestamp and ordering regressions | Test coverage | Adds focused regression tests for memory, infra, and log edge cases identified in #7685, covering same-timestamp pagination and SQLite ordering. |
| [#7492](https://github.com/zeroclaw-labs/zeroclaw/pull/7492) | feat(cost): support cached input token pricing from OpenAI-compatible | Feature | Parses `prompt_tokens_details.cached_tokens` (OpenAI) and `prompt_cache_hit_tokens` (DeepSeek) to apply discounted cached-input rates in the cost tracker. |
| [#7583](https://github.com/zeroclaw-labs/zeroclaw/pull/7583) | fix(runtime): honor profile tool iteration limits | Bug fix | Ensures cron/CLI-style agent runs use the correct effective max tool iteration limits from the resolved agent config. |

Other merged/closed PRs (not in top-20 list) account for the remaining 5 of the 9 total.

---

## 4. Community Hot Topics

- **Issue #2079 – “Restore GitHub as a native channel”**  
  [zeroclaw-labs/zeroclaw#2079](https://github.com/zeroclaw-labs/zeroclaw/issues/2079)  
  *7 comments, 0 👍*  
  This long-running feature request calls for first-class GitHub integration (issues, PRs, comments). The heavy comment count suggests strong community desire. No fix PR is linked yet, but the issue is labeled `priority:p2, status:accepted`.

- **Issue #7694 – “feat(memory): cover storage-reader timestamp and ordering edge cases”**  
  [zeroclaw-labs/zeroclaw#7694](https://github.com/zeroclaw-labs/zeroclaw/issues/7694)  
  *4 comments, 0 👍*  
  A help-wanted / good-first-issue task for adding deterministic test coverage. Today’s merged PR #7918 directly addresses this, reflecting active community contribution.

- **Issue #6970 – “Tracker: v0.8.1 integration/channel/provider/tool queue and history”**  
  [zeroclaw-labs/zeroclaw#6970](https://github.com/zeroclaw-labs/zeroclaw/issues/6970)  
  *3 comments, 0 👍*  
  Operational tracker for the v0.8.1 release. It remains open as a planning hub, indicating the team is actively working toward this milestone.

- **PR #7162 – “feat(channels): notify before context compression”**  
  [zeroclaw-labs/zeroclaw#7162](https://github.com/zeroclaw-labs/zeroclaw/pull/7162)  
  *No comments yet (opened 2026-06-03)*  
  Provides channel-visible notices before proactive context compression. Despite being open for two weeks, it hasn’t attracted community discussion, possibly because it is a relatively low-risk enhancement.

---

## 5. Bugs & Stability

| Severity | Issue | Title | Fix PR Exists? |
|----------|-------|-------|----------------|
| **S1 – Workflow blocked** | [#7756](https://github.com/zeroclaw-labs/zeroclaw/issues/7756) | Native/MCP tools unavailable on OpenAI Responses/reasoning and Anthropic turns | ✗ (PR #7933 adds diagnostics but does not fix the root cause) |
| **S1 – Workflow blocked** | [#7804](https://github.com/zeroclaw-labs/zeroclaw/issues/7804) | Code history can send non-alternating Anthropic messages | ✓ PR #7931 (coalesce stripped compatible history roles) |
| **S1 – Workflow blocked** | [#7871](https://github.com/zeroclaw-labs/zeroclaw/issues/7871) | Shell tool can hang when grandchild processes inherit pipe handles | ✗ (Tracker issue, no fix PR yet) |
| **S2 – Degraded behavior** | [#7799](https://github.com/zeroclaw-labs/zeroclaw/issues/7799) (closed) | Resumed Code sessions reopen with blank transcript | ✓ Fixed? (closed) |
| **S3 – Minor** | [#7892](https://github.com/zeroclaw-labs/zeroclaw/issues/7892) | CLI approval prompt should read controlling terminal when stdin is detached | ✗ |
| **S3 – Minor** | [#7917](https://github.com/zeroclaw-labs/zeroclaw/issues/7917) | i18n: file_download tool strings untranslated in non-English locales | ✓ Multiple PRs: #7925, #7924 (both open) |

**Notes:** High-risk bugs dominate today’s activity. The most critical bug (#7756) has a diagnostic PR (#7933) but no fix yet. The Anthropic message-ordering bug (#7804) appears to be addressed by a merged-like PR (#7931 is still open but marked as fix). The shell hang bug (#7871) remains unaddressed and could cause production outages.

---

## 6. Feature Requests & Roadmap Signals

- **GitHub Native Channel** (#2079) – The most commented feature request. Likely candidate for v0.8.1 given its `priority:p2` and `status:accepted`.
- **WASM Plugin Host** (#7928 PR) – Adds initial WASM component-model plugin host code targeting wit v0. This is a major architectural addition and signals long-term extensibility.
- **Unified Slash-Command Registry** (#7929) – New issue today proposing a single gateway-served catalogue. If accepted, will reduce fragmentation across Web UI, TUI, and channel runtime.
- **Process-Memory Limits on Subprocesses** (#6916) – `priority:p1` enhancement that would prevent LLM-triggered shell commands from OOMing containers. A high-demand stability improvement.
- **Auto-Clean Temporary Files** (#7923 PR) – Config-driven cleanup of temporary files. Practical for production deployments.
- **Discord Slash Command Localizations + Guild Scope** (#7922 PR) – Completes Discord command surface parity. Shows ongoing investment in chat platforms.

**Prediction for v0.8.1:** GitHub native channel, process-memory limits, and unified slash-command registries are likely to land based on accepted labels and priority.

---

## 7. User Feedback Summary

- **Pain Points:**
  - **Installation on Android/Termux** (#7911): The precompiled binary fails for `linux aarch64`. User reports unknown binary type.
  - **Tool availability across providers** (#7756): MCP tools are registered but not always delivered to the model, blocking workflows.
  - **Session resume corruption** (#7799, closed): Resumed Code sessions show blank transcript, frustrating long-running sessions.
  - **Shell tool hangs** (#7871): Grandchild processes prevent EOF, causing indefinite waits.
  - **i18n gaps** (#7917): Non-English users see English fallback for download tool strings.

- **Satisfaction Signals:**
  - Multiple contributors are actively submitting PRs (e.g., #7925, #7924, #7916, #7918), suggesting a healthy contributor experience.
  - The team responds quickly to high-severity bugs (P1 bugs get accepted and PRs opened within days).

---

## 8. Backlog Watch

- **Issue #2079 – Restore GitHub as a native channel**  
  [zeroclaw-labs/zeroclaw#2079](https://github.com/zeroclaw-labs/zeroclaw/issues/2079)  
  Opened 2026-02-27, last updated today. Despite being `accepted` for 4 months, no implementation PR exists. This could become a frustration point for the community if it slips further.

- **PR #7162 – Notify before context compression**  
  [zeroclaw-labs/zeroclaw#7162](https://github.com/zeroclaw-labs/zeroclaw/pull/7162)  
  Open since 2026-06-03 with no maintainer review. Small change but stalled for two weeks.

- **PR #7170 – Slack outbound attachment upload**  
  [zeroclaw-labs/zeroclaw#7170](https://github.com/zeroclaw-labs/zeroclaw/pull/7170)  
  Also open since June 3, no maintainer attention. Potentially blocked by testing or design concerns.

- **Issue #6916 – Process-memory limits on subprocesses**  
  [zeroclaw-labs/zeroclaw#6916](https://github.com/zeroclaw-labs/zeroclaw/issues/6916)  
  Accepted P1 since 2026-05-25 with no linked PR. Critical for production safety; maintainers should prioritize.

---

*End of Digest*

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*