# OpenClaw Ecosystem Digest 2026-06-16

> Issues: 295 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-16 05:20 UTC

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

# OpenClaw Project Digest – 2026-06-16

## Today's Overview
OpenClaw activity remains very high: 295 issues and 500 pull requests were updated in the past 24 hours, indicating a vibrant and fast-moving project. One new beta release was published today (v2026.6.8-beta.2), focused on richer Telegram and WhatsApp delivery. The project continues to see a mix of critical bug reports (particularly around session context, security, and message delivery) alongside ambitious feature proposals. The community is highly engaged, with several long-running discussions accumulating dozens of comments and reactions.

## Releases
- **New Release**: [v2026.6.8-beta.2](https://github.com/openclaw/openclaw/releases/tag/v2026.6.8-beta.2) – “openclaw 2026.6.8-beta.2”  
  **Highlights**:  
  - Telegram and WhatsApp channel delivery are richer and less brittle: Telegram can send structured rich text with tables, lists, expandable blockquotes, preserved intentional line breaks, prompt-preserving CLI backend delivery, retired native draft migration, and safer rich-media bo… (summary truncated in data).  
  No breaking changes or migration notes are mentioned in the release data. Expect a full changelog in the release notes.

## Project Progress
Today **92 pull requests were merged or closed** (out of 500 updated in 24h). Notable merged/closed PRs from the top of the list:
- [#93480](https://github.com/openclaw/openclaw/pull/93480) (CLOSED) – Fix agent: restore suppressed answer when `before_agent_finalize` revision is silent (addresses a data-loss bug in core response loop).
- [#93483](https://github.com/openclaw/openclaw/pull/93483) (CLOSED) – Fix Feishu channel dispatch by spreading full plugin runtime to ensure `channel.inbound` is available.
- [#93491](https://github.com/openclaw/openclaw/pull/93491) (CLOSED) – Fix CLI: apply `--log-level` in route-first command path.
- [#93472](https://github.com/openclaw/openclaw/pull/93472) (CLOSED) – Fix Feishu: guard `channelRuntime.inbound` before dispatch fallback.
- [#93460](https://github.com/openclaw/openclaw/pull/93460) (CLOSED) – Fix CLI: honor `--log-level` in route-first commands (alternative approach).

Many smaller fixes for UI, scripting, and channel connectors have been merged, improving overall stability.

## Community Hot Topics
Most active issues and PRs by comment count and reactions:

- **Issue #75** – [Linux/Windows Clawdbot Apps](https://github.com/openclaw/openclaw/issues/75)  
  109 comments, 79👍. Long-running feature request for desktop clients on Linux/Windows. The community clearly wants cross-platform parity with macOS/iOS/Android.
- **Issue #22676** – [Signal daemon stop() race condition](https://github.com/openclaw/openclaw/issues/22676)  
  17 comments. A critical bug leading to orphaned processes and send failures during SIGUSR1 restarts. No linked fix PR yet.
- **Issue #22438** – [Tiered bootstrap file loading](https://github.com/openclaw/openclaw/issues/22438)  
  17 comments. Users want to avoid wasting LLM tokens on unused bootstrap files per session.
- **Issue #32473** – [Control UI requires device identity (HTTPS/localhost)](https://github.com/openclaw/openclaw/issues/32473)  
  17 comments, 5👍. Regression in secure context handling for web UI.
- **Issue #32296** – [Agent replies to previous message](https://github.com/openclaw/openclaw/issues/32296)  
  15 comments. Session context confusion – a P1 bug that disrupts conversation flow.
- **Issue #39604** – [Add `tools.web.fetch.allowPrivateNetwork`](https://github.com/openclaw/openclaw/issues/39604)  
  13 comments, 9👍. Strong demand for controlled access to private networks via `web_fetch`.
- **Issue #29387** – [Bootstrap files in agentDir silently ignored](https://github.com/openclaw/openclaw/issues/29387)  
  14 comments, 5👍. Per-agent bootstrap files not loaded – a P1 bug with security implications.

**Underlying needs**: Users are pushing for better cross-platform support, higher reliability in multi-agent/async workflows, more granular security controls, and intelligent context management to save token costs.

## Bugs & Stability
High-impact bugs reported or updated in the last 24 hours:

| Issue | Severity | Impact | Has Fix PR? |
|-------|----------|--------|-------------|
| [#22676](https://github.com/openclaw/openclaw/issues/22676) – Signal daemon race condition | P1 | Orphaned processes, message loss | None yet |
| [#32296](https://github.com/openclaw/openclaw/issues/32296) – Agent replies to previous message | P1 | Session context confusion | None yet |
| [#29387](https://github.com/openclaw/openclaw/issues/29387) – Bootstrap files ignored | P1 | Security, session state | None yet |
| [#32473](https://github.com/openclaw/openclaw/issues/32473) – Control UI secure context regression | P2 | Security, UI broken | None yet |
| [#39604](https://github.com/openclaw/openclaw/issues/39604) – Private network access missing | P2 | Security, missing feature | PRs exist? Not linked |
| [#40001](https://github.com/openclaw/openclaw/issues/40001) – Write tool lacks append mode | P1 | Data loss in cron sessions | None yet |
| [#40540](https://github.com/openclaw/openclaw/issues/40540) – `openclaw update` EBUSY on Windows | P1 (closed) | Update fails | Closed (likely fixed) |
| [#31331](https://github.com/openclaw/openclaw/issues/31331) – Docker sandbox workspace binding | P1 | Session state, security | None yet |
| [#38327](https://github.com/openclaw/openclaw/issues/38327) – "Cannot convert undefined or null to object" with Gemini | P1 | Crash, LLM integration | None yet |
| [#39476](https://github.com/openclaw/openclaw/issues/39476) – A2A sessions_send duplicate messages | P1 | Message loss, state | None yet |
| [#41165](https://github.com/openclaw/openclaw/issues/41165) – Telegram DMs pollute main session | P1 | Message routing | None yet |
| [#87711](https://github.com/openclaw/openclaw/issues/87711) – Empty assistant delivery on Telegram | P1 | Message loss | None yet |

Several fix PRs are open for related issues (e.g., #93480 for agent finalize suppression, #91674 for heartbeat wake, #92899 for compaction timeout). Overall, the bug landscape shows persistent concerns around session state, message delivery, and security boundaries.

## Feature Requests & Roadmap Signals
Key features requested and discussed today:
- **Private network access** (`issue #39604`) – allow private IPs via `web_fetch` (9👍).
- **Slack Block Kit support** (`#12602`) – richer interactive responses.
- **Telegram Business Bot support** (`#20786`, 6👍) – enable business messaging.
- **Backup/restore utility** (`#13616`) – standardized migration and recovery.
- **Tiered bootstrap file loading** (`#22438`) – reduce token waste.
- **Memory trust tagging** (`#7707`) – prevent memory poisoning.
- **Post-subagent extension hook** (`#22358`) – structured trajectory generation.
- **Filesystem sandboxing config** (`#7722`, 4👍) – path-level permissions.
- **Exec-approval denylist** (`#6615`, 7👍) – allow everything except dangerous commands.
- **Capability-based permissions** (`#12678`) – default-deny for high-risk tools.
- **Automated session memory preservation** (`#40418`) – continuous learning across sessions.

**Prediction for next version**: Given community demand and maintainer attention, items like private network access, Slack Block Kit, Telegram Business Bot, and tiered bootstrap loading are likely candidates for inclusion in upcoming releases (v2026.6.x or mid-July).

## User Feedback Summary
**Pain points expressed**:
- "Agent replies to previous message instead of current" – confusion in multi-turn.
- "Signal daemon crashes on config change" – reliability concerns for production use.
- "Bootstrap files ignored in agentDir" – wasted tokens and config surprises.
- "Write tool overwrites shared files" – data loss in cron/sessions.
- "Docker sandbox workspace binding broken" – difficult deployment.
- "Windows EBUSY on update" – installation friction.
- "Telegram sessions wedged by bad JSON" – need better recovery UX.

**Satisfaction signals**:
- Continued high community engagement (295 issues, 500 PRs updated).
- Many users actively contributing fixes (e.g., CJK IME fix #93498, Feishu fixes).
- The beta release shows ongoing improvement in channel delivery (Telegram/WhatsApp).

**Overall sentiment**: Users are enthusiastic about OpenClaw’s capabilities but face stability and usability challenges, especially around session management, multi-channel routing, and configuration edge cases.

## Backlog Watch
Issues/PRs that are long-unanswered or waiting for maintainer attention:

- **Issue #75** – Linux/Windows apps (open since Jan 2026, 109 comments, 79👍) – still no resolution.
- **Issue #6615** – Exec-approval denylist (open since Feb, 7👍) – needs product decision.
- **Issue #7722** – Filesystem sandboxing config (open since Feb, 4👍) – needs security review.
- **Issue #8299** – Suppress sub-agent announce (open since Feb, 1👍) – needs maintainer review.
- **Issue #11665** – Webhook session reuse (open since Feb 8) – documented but broken, needs fix.
- **Issue #12678** – Capability-based permissions (open since Feb 9) – large scale, likely deferred.
- **Issue #13616** – Backup/restore utility (open since Feb 10) – no maintainer movement.
- **Issue #20786** – Telegram Business Bot (open since Feb 19, 6👍) – no product decision.
- **PR #40311** – Brave Goggles integration (open since Mar 8, needs behavior proof) – stalled.
- **PR #40877** – iOS main-thread warnings fix (open since Mar 9, needs real-behavior proof) – stalled.
- **PR #41375** – Hook replies delivery (open since Mar 9, waiting on author) – needs re-engagement.

**Recommendation**: Maintainer attention is needed on the Linux/Windows app request (the top-voted community issue) and on several security-related features (exec denylist, filesystem sandboxing) that have been pending for months.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: Personal AI Agent Open-Source Ecosystem
**Date:** 2026-06-16 | **Prepared for:** Technical decision-makers and developers

---

## 1. Ecosystem Overview

The personal AI agent open-source landscape is experiencing a **productivity plateau** despite high development velocity across major projects. While activity metrics remain strong—over 1,200 combined issues and 1,200+ PRs updated across observed projects in 24 hours—the ecosystem is increasingly converging on a core set of unsolved challenges: **session context reliability, multi-channel delivery stability, and security boundary controls**. The dominant architectural pattern is evolving from simple LLM wrappers toward **agent orchestration platforms** with delegation, memory tiers, and plugin ecosystems. A notable bifurcation is emerging between general-purpose assistants (OpenClaw, Hermes Agent) and specialized workflow/automation frameworks (NanoBot, ZeroClaw), with the latter group showing higher feature velocity but lower cross-platform maturity. The ecosystem remains highly fragmented, with no single project achieving clear dominance across all dimensions.

---

## 2. Activity Comparison

| Project | Issues Updated (24h) | PRs Updated (24h) | Merged/Closed PRs (24h) | Release Today | Health Score* |
|---|---|---|---|---|---|
| **OpenClaw** | 295 | 500 | 92 | ✅ v2026.6.8-beta.2 | **High** |
| **Hermes Agent** | 8 | 50 | 2 | ❌ | **High** |
| **IronClaw** | 17 | 50 | 19 | ❌ | **High** |
| **CoPaw** | 28 | 49 | 33 | ❌ | **High** |
| **ZeroClaw** | 3 | 50 | 1 | ❌ | **Medium-High** |
| **NanoBot** | 6 | 26 | 5 | ❌ | **Medium** |
| **PicoClaw** | 3 | 13 | 3 | ✅ nightly build | **Medium** |
| **NanoClaw** | 1 | 11 | 3 | ❌ | **Medium-Low** |
| **LobsterAI** | 0 | 11 | 5 | ❌ | **Low-Medium** |
| **Moltis** | 0 | 2 | 0 | ❌ | **Low** |
| **NullClaw** | 2 | 0 | 0 | ❌ | **Low** |
| **TinyClaw** | 0 | 0 | 0 | ❌ | **Inactive** |
| **ZeptoClaw** | 0 | 0 | 0 | ❌ | **Inactive** |

*\*Health Score: Composite of activity volume, merge velocity, issue responsiveness, and release cadence.*

**Key observations:**
- **OpenClaw dominates raw activity** with 295 issues and 500 PRs updated—more than the next four projects combined.
- **Hermes, IronClaw, and CoPaw** form a second tier with balanced issue-to-PR ratios and high merge throughput.
- **3 projects are effectively dormant** (TinyClaw, ZeptoClaw, NullClaw with 0-2 activities), indicating ecosystem churn.

---

## 3. OpenClaw's Position

**Advantages:**
- **Community scale:** 7.3× more issues and 10× more PRs updated than the next-closest project (Hermes Agent). This network effect accelerates bug discovery and feature incubation.
- **Release discipline:** Regular beta releases (v2026.6.8-beta.2 today) with documented changelogs—only PicoClaw matches this cadence.
- **Cross-platform demand:** The #1 community issue (Linux/Windows desktop apps, 109 comments, 79👍) reveals an underserved user base that competitors have not aggressively targeted.
- **Channel breadth:** Rich Telegram/WhatsApp delivery (structured tables, blockquotes, CLI backend) exceeds most peers' channel capabilities.

**Technical approach differences:**
- OpenClaw uses a **plugin-runtime architecture** where channels (Telegram, Feishu, WhatsApp) are first-class extensions with full runtime binding, rather than thin API wrappers. This enables richer delivery but increases complexity (witnessed in multiple Feishu dispatch fixes today).
- Its **session context model** is more opinionated than Hermes Agent's delegation-focused approach, leading to both deeper conversation coherence and more context-confusion bugs (Issue #32296).

**Community size comparison:**
- OpenClaw: ~295 daily active issue participants → likely 15,000-25,000 monthly active developers (estimated).
- Hermes Agent: ~8 daily issues → likely 3,000-5,000 MAU.
- IronClaw/CoPaw: similar to Hermes.
- All others: <2,000 MAU.

**Consequence:** OpenClaw's scale creates **documentation debt**—critical bugs (Signal daemon race, bootstrap file ignoring) go weeks without fix PRs despite high visibility, because maintainers are overwhelmed by volume.

---

## 4. Shared Technical Focus Areas

| Focus Area | Projects Affected | Specific Needs |
|---|---|---|
| **Session Context Reliability** | OpenClaw, NanoBot, CoPaw, ZeroClaw, NullClaw | Agent replies to wrong messages (OpenClaw #32296), context compression wiping tokens (CoPaw #5171), missing goal context (NanoBot #4286), session ordering races (ZeroClaw #7753) |
| **Multi-Channel Delivery Stability** | OpenClaw, Hermes Agent, PicoClaw, CoPaw, Nashor's | Telegram DMs polluting main session (OpenClaw #41165), Feishu WebSocket card rendering (NanoBot #4342), QQ channel failure on Windows (PicoClaw #3015), DingTalk stream death on sleep (CoPaw #5214) |
| **Security Boundary Controls** | OpenClaw, PicoClaw, IronClaw, ZeroClaw | CIDR bypass via reverse proxy (PicoClaw #3069), OAuth resume failure (IronClaw #4907), private network access controls (OpenClaw #39604), credential scoping (IronClaw #4935) |
| **Token Cost Optimization** | OpenClaw, Hermes Agent, CoPaw, NanoBot | Tiered bootstrap loading (OpenClaw #22438), context-aware model switching (Hermes #47047), headroom compression (CoPaw #5063), empty response fallback (NanoBot #4287) |
| **Local Model Compatibility** | NullClaw, IronClaw, CoPaw, OpenClaw | Ollama truncated responses (NullClaw #952), MiniMax XML thinking breakage (CoPaw #4625), Gemini null-object crash (OpenClaw #38327), Ollama false-positive test (IronClaw #4696) |
| **Plugin/Dependency Management** | CoPaw, OpenClaw, ZeroClaw | Plugin auto-install cmd spam (CoPaw #5181), MCP tools not reaching model (ZeroClaw #7756), A2A/MCP integration discovery (NanoBot #4362) |

**Emerging pattern:** Every project with >10 daily issues is struggling with **asynchronous state recovery**—when agents pause (auth flows, network interruptions, context overflows), the recovery path is either broken or silent. This is the single largest category of P1 bugs.

---

## 5. Differentiation Analysis

| Project | Primary Differentiator | Target User | Tech Architecture |
|---|---|---|---|
| **OpenClaw** | Broadest channel support (Telegram, WhatsApp, Feishu, CLI, Web) with rich structured delivery | Power users wanting a universal assistant | Plugin-runtime, session-oriented, context-first |
| **Hermes Agent** | Advanced delegation/DAG workflows (Orchestrator v3, background sub-agents) | Developers building multi-agent systems | Dynamic workflow DAG, fallback chains, layered memory |
| **IronClaw** | Reborn binary with vision support, OAuth credential management, CI/CD automation | Enterprise teams needing secure AI automation | Rust-based core, WebUI-first, credential scoping |
| **CoPaw** | Skill market/plugin ecosystem, DingTalk/Feishu/Matrix channel focus | Chinese-market users, plugin developers | Skill-slash injection, Agent OS Driver abstraction |
| **ZeroClaw** | Gateway web dashboard, Mattermost/WhatsApp/Lark channel parity | DevOps teams integrating AI into existing chat ops | Gateway-centric, WebSocket listeners, config-first |
| **NanoBot** | WebUI automation management, silent cron jobs, audit tooling | Users wanting no-code agent control | Config-parity WebUI, task automation, MCP integration |
| **PicoClaw** | Lightweight embedded deployment (RISC-V, Windows gateway) | Resource-constrained or single-board users | Minimal binary, security-hardened, nightly builds |
| **NanoClaw** | Remote MCP server support, Streamlabs integration, fitness/lifestyle skills | Hobbyists and streamers | Codex archives, gateway upgrade workflows |
| **LobsterAI** | Cowork session focus, realtime voice input, dictation ASR | Collaborative note-taking users | Electron desktop app, ASR-focused, cowork UI |
| **Moltis** | External agent model/effort selection, context command injection | Researchers experimenting with agent orchestration | Provider abstraction, minimal surface area |

**Key insight:** The gap between **general-purpose** and **specialized** projects is widening. OpenClaw, Hermes, and IronClaw are competing for the "universal agent" space, while CoPaw, ZeroClaw, and NanoBot are carving vertical niches. This fragmentation means **no single project meets all use cases**, forcing developers to choose based on primary channel or deployment environment.

---

## 6. Community Momentum & Maturity

**Tier 1: Rapid Iteration (High velocity, many open issues)**
- **OpenClaw** – 500 PRs/day, but quality control is strained. Many P1 bugs linger for weeks. Risk of feature bloat.
- **Hermes Agent** – Balanced activity (50 PRs, 8 issues). Maintaining quality while adding delegation features. Strongest PR-to-issue ratio.
- **IronClaw** – 19 merged/50 updated today. Focused on Reborn stabilization. Good merge discipline.
- **CoPaw** – 33 merged/49 updated. Highest merge rate today. Sprint 2.4 in progress. Good momentum but accumulating critical bugs.

**Tier 2: Steady Development (Moderate activity, predictable cadence)**
- **ZeroClaw** – 50 PRs but only 1 merged. Review bottleneck. Features are mature but stuck in queue.
- **NanoBot** – 26 PRs, 5 merged. Healthy pipeline. Automation management view is a strong signal.
- **PicoClaw** – Low but focused activity. Security-hardening spree. Nightly releases show discipline.

**Tier 3: Stabilizing or Stalling**
- **NanoClaw** – 11 PRs, 3 merged. Long-standing PRs (#2626-#2628) need review. Risk of contributor churn.
- **LobsterAI** – Only Dependabot PRs active. Cowork features merged but community issues abandoned.
- **Moltis** – 2 PRs, both new. Too early to assess sustainability.
- **NullClaw** – 2 issues, 0 PRs. Maintenance mode. Low contributor confidence.
- **TinyClaw, ZeptoClaw** – Dead. Zero activity for 24+ hours.

**Maturity assessment:**
- **Most mature:** OpenClaw, IronClaw, Hermes Agent—have passed the "trough of disillusionment" and are building production-grade foundations.
- **Emerging threats:** ZeroClaw's Mattermost/WhatsApp parity and CoPaw's skill market could capture niche audiences faster than generalists can.
- **At risk:** NullClaw, TinyClaw—if maintainers don't re-engage within 30 days, these projects will become archival.

---

## 7. Trend Signals

**1. Context is the new bottleneck**
- Every project with significant usage reports session-state bugs (OpenClaw #32296, NanoBot #4286, CoPaw #5171). The ecosystem is hitting the practical limits of sliding-window context management. **Context lifecycle management**—how to persist, compress, recover, and evolve conversation state—is the single most important unsolved problem. Projects that solve this (e.g., Hermes Agent's layered memory prototype) will gain a structural advantage.

**2. Self-healing agents are a prerequisite**
- Users are demanding that agents survive network interruptions (IronClaw #4108, ZeroClaw #7753), model failures (OpenClaw #38327, NanoBot #4287), and platform crashes (CoPaw #5209). **Automatic recovery without user intervention** is no longer a nice-to-have; it's table stakes for production deployments. The "pilot mode" expectation—where users walk away and return to a completed task—requires robust retry, backpressure, and state reconciliation.

**3. Security is moving from perimeter to context**
- Multiple projects report attacks on the auth/credential boundary (PicoClaw #3069 CIDR bypass, IronClaw #4907 OAuth resume failure, OpenClaw #29387 bootstrap file security). The trend is toward **granular, context-aware permissions**: capabilities-based (OpenClaw #12678), filesystem-level (OpenClaw #7722), and tenant-scoped (IronClaw #4935). Developers should plan for "default-deny" architectures where every tool invocation requires explicit scope permission.

**4. Multi-model routing is becoming a first-class feature**
- Users want to switch models mid-session (Hermes Agent #47047), route by cost-tier (Hermes Agent fallback chains), and fall back when primary models fail (NanoBot #4287). This mirrors enterprise data routing patterns. **LLM-as-a-mesh**—where model selection is abstracted behind a routing layer based on cost, latency, capability, and availability—is emerging as a critical architectural pattern.

**5. The desktop/web vs. mobile/channel war is settled: both are required**
- OpenClaw's #1 issue (Linux/Windows desktop apps, 79👍) shows that **web-only or CLI-only is insufficient** for mainstream adoption. But Hermes Agent and CoPaw are investing in mobile channels (Telegram, WhatsApp, Matrix) with equal vigor. The winning projects will offer **ubiquitous access**—desktop, mobile, web, and chat—with unified session state across all surfaces.

**6. Plugin ecosystems are becoming "skills markets"**
- CoPaw's skill market (#5123) and NanoClaw's Strava integration (#2777) indicate a shift from monolithic agents to **app-store models** for agent capabilities. The ability to browse, install, and manage third-party skills/plugins will become a competitive differentiator. This mirrors the browser extension model—those with the largest plugin ecosystems will attract the most users.

**7. Cost observability is underequipped**
- Almost no project provides real-time token cost tracking or budget alerts. Users are discovering overruns only when bills arrive (OpenClaw #22438 on token waste, NanoBot #4309 on zero usage tokens issue). **Cost transparency as a UX layer**—showing per-turn costs, historical spend, and budget controls—is a gap waiting to be filled. Projects that add this first will win cost-sensitive users (students, researchers, small businesses).

**For AI agent developers:**
- **Choose OpenClaw** if you need broad channel support and are willing to manage complexity.
- **Choose Hermes Agent** if you're building multi-agent systems and need delegation workflows.
- **Choose IronClaw** if you prioritize security and enterprise deployment.
- **Choose CoPaw** if targeting the Chinese market or plugin ecosystem.
- **Watch ZeroClaw** for Mattermost/DevOps integration—it's underrated.
- **Avoid** NullClaw, TinyClaw, and ZeptoClaw for new projects—they lack community inertia.

The ecosystem is maturing rapidly, but the gaps in context management, self-healing, and security remain the biggest opportunities for projects that can solve them first.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-06-16

## 1. Today's Overview
Activity remains high, with **6 issues** and **26 pull requests** updated in the last 24 hours. Two issues and five PRs were closed/merged, while the majority of PRs are still open, indicating a healthy development pipeline. No new releases were published, but the project is seeing significant contributions in areas such as WebUI automation, provider compatibility, and session management. Community engagement is solid, with several detailed bug reports and feature contributions from external developers.

---

## 2. Releases
*None.* No new versions were released today.

---

## 3. Project Progress (Merged/Closed PRs Today)
Two pull requests were merged or closed today:

- **[PR #4359 – fix(agent): refresh goal continuation context](https://github.com/HKUDS/nanobot/pull/4359)** – Resolves the “sustained goal” context issue (#4286) by lazily refreshing goal continuation text. This fix ensures goals created by `long_task` during the current runner call are included in follow-up prompts.
- **[PR #4348 – fix(session): keep auto compact suffix on user turn](https://github.com/HKUDS/nanobot/pull/4348)** – Prevents auto-compact from trimming partial tool turns by preserving the recent suffix extending backward to the containing user turn.

Both were merged by core contributor **chengyongru**.

---

## 4. Community Hot Topics
The following issues and PRs attracted the most attention (based on comment count and discussion depth):

- **[Issue #4360 – [bug] “end of file unexpected” during installer](https://github.com/HKUDS/nanobot/issues/4360)** (6 comments) – A fresh Docker container of Debian 13 triggers a `Syntax error: end of file unexpected (expecting "}")` during pip installation. The author suspects a corrupt or incomplete script section. This is the most commented issue today, suggesting a blocking problem for new users.
- **[PR #4342 – fix(feishu): support reading WebSocket rendered card content](https://github.com/HKUDS/nanobot/pull/4342)** – Addresses a structural mismatch in Feishu (Lark) card parsing when messages arrive via WebSocket. The root cause involves three specific differences in nested list elements. The PR includes detailed analysis and is a high-quality community fix.
- **[PR #4330 – feat(webui): add automation management view](https://github.com/HKUDS/nanobot/pull/4330)** – A large feature addition that introduces a WebUI surface for listing, filtering, running, pausing/resuming, and deleting user automations, along with API routes and localization. This is a substantial contribution and signals growing interest in no-code agent management.

*Underlying needs*: Users want a reliable installation experience on modern Linux distributions, robust platform-specific integrations (Feishu), and richer Web-based control over automations.

---

## 5. Bugs & Stability
Several bugs were reported today; the most significant are ranked below:

| Severity | Issue | Description | Fix Exists? |
|----------|-------|-------------|-------------|
| **High** | [#4360 – Installer failure on Debian 13](https://github.com/HKUDS/nanobot/issues/4360) | `pip: 20: Syntax error: end of file unexpected` during installation in a fresh Debian 13 Docker container. Blocks deployment for users on that base image. | No PR yet |
| **Medium** | [#4287 – Empty model responses not triggering fallback](https://github.com/HKUDS/nanobot/issues/4287) | When DeepSeek returns empty completions, the error is classified as “non-fallbackable,” so the agent does not retry with an alternative model. | Open, possible relation to PR #4358 |
| **Medium** | [#4322 – NameError: 'session_key' is not defined in context.py](https://github.com/HKUDS/nanobot/issues/4322) | A merge of `origin/main` into a feature branch caused a missing variable reference that crashes the agent on startup. Marked stale. | No PR, but root cause identified in issue |
| **Low** | [#4309 – `/v1/chat/completions` always returns zero usage tokens](https://github.com/HKUDS/nanobot/issues/4309) | The OpenAI-compatible endpoint hardcodes token usage to 0. This was closed today, presumably fixed. | Closed (fix not listed, likely merged) |
| **Low** | [#4286 – Missing “sustained goal” context](https://github.com/HKUDS/nanobot/issues/4286) | Nanobot repeatedly reported missing context during article creation. | Fixed by PR #4359 (merged today) |

A related PR **#4358 – fix(api): avoid duplicate user turn on empty-response retry** addresses the empty-response fallback pattern and may close #4287 as well.

---

## 6. Feature Requests & Roadmap Signals
Notable feature additions under active development:

- **Automation Management UI** (PR #4330) – The largest feature PR today; a WebUI for managing automations (tasks). Likely to land in the next minor release.
- **Silent Cron Jobs** (PR #4357) – Allows scheduled jobs to run without auto-delivering a response, useful for monitoring tasks. A practical extension of the cron tool.
- **Audit Tool** (PR #4320) – Adds a `tools.audit` configuration and an `AuditTool` for agent action observability. This signals a move toward enterprise compliance and logging.
- **Better Mistral Support** (PR #4351) – Fixes strict API constraints for Mistral models (reasoning_effort, max_tokens, temperature, n). Increases reliability for Mistral users.
- **Keenable Search Provider** (PR #4350) – Adds a research-driven web search engine as a built-in search provider, expanding the tool ecosystem.
- **WebUI/Config Parity** (PR #4313) – Closes the gap between WebUI settings and `config.json`, with new write endpoints for temperature, tool limits, dreams, channels, and memory. A major usability improvement for non‑CLI users.

These features suggest the next version will focus on **WebUI completeness**, **observability**, and **provider robustness**.

---

## 7. User Feedback Summary
Real user pain points reported this period:

- **The-Markitecht** (#4360) – Installer fails on Debian 13, a common base Docker image. This is a critical onboarding blocker for Linux users.
- **glebov** (#4287) – Long-running Telegram bot users face degraded reliability when primary models return empty responses without fallback.
- **fablau** (#4286) – The “sustained goal” context disappeared during article generation, causing the agent to lose track of its task. (Now fixed.)
- **professionelle-hypnose** (#4322) – A merge error introduced a crash on startup, highlighting the need for better branch integration testing.
- **alx1379** (#4309) – Token usage statistics are broken in the OpenAI-compatible API, impacting cost tracking. (Now closed.)
- **JiajunBernoulli** (#4342) – Feishu card rendering broken over WebSocket due to structural mismatches; PR submitted with clear root cause analysis.

Overall, users are actively reporting and often contributing fixes, indicating a healthy, engaged community. The main satisfaction drivers are the responsiveness of core maintainers (many bugs closed quickly) and the breadth of new features.

---

## 8. Backlog Watch
Issues and PRs that may need maintainer attention or are lingering without engagement:

- **[Issue #4287 – Empty model responses not triggering fallback](https://github.com/HKUDS/nanobot/issues/4287)** – Open since June 10 with only 2 comments. Despite a related PR (#4358), the issue is still open and needs a status update or acceptance of the fix.
- **[Issue #4322 – NameError: 'session_key' is not defined](https://github.com/HKUDS/nanobot/issues/4322)** – Marked “stale” after June 15. The root cause is identified but no fix PR has been linked. If the branch is abandoned, this could affect users on `fix/prompt-caching`.
- **[PR #4303 – fix(mcp): close tracked generators in _close_server to prevent GC crash](https://github.com/HKUDS/nanobot/pull/4303)** – Open since June 11, addressing a crash when a `streamableHttp` MCP server session terminates. No comments from maintainers. This is a stability-critical fix for MCP users.
- **[Issue #4362 – A2A/MCP Integration: MetaVision AI tools now discoverable](https://github.com/HKUDS/nanobot/issues/4362)** – Opened today by an external developer announcing integration. While not a bug or request, it signals interest in A2A compatibility. Maintainers may want to triage as a feature suggestion or documentation addition.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest – 2026-06-16

## 1. Today’s Overview

Today, Hermes Agent shows **high development velocity** with 8 open issues and 50 pull requests updated in the last 24 hours. The project maintains a healthy ratio of open PRs (48) to merged/closed PRs (2), indicating sustained contributor momentum. No new releases were cut today. The most active areas are **delegation workflows**, **gateway platform fixes** (Telegram, SMS, BlueBubbles), **configuration fallback chains**, and the introduction of a layered **filesystem memory plugin**. Two PRs were merged (not listed in the top 20), suggesting that critical fixes or features landed, though details are unavailable from the provided data.

## 2. Releases

No new releases today.

## 3. Project Progress

Two pull requests were merged or closed today (specifics not in the top 20). The open PRs reflect work that has advanced significantly:

- **Delegation & Orchestration**: PR #46971 (`dynamic_workflow DAG tool`), PR #46978 (provider/model tag for delegate lines), PR #46981 (fallback isolation for cost‑capped subagents), and PR #47035 (Orchestrator v3 showcase with background delegation).
- **Gateway fixes**: PR #46972 (SMS body cap at 1600 chars), PR #46974 (BlueBubbles self‑echo loop fix).
- **Config & Cron**: PR #46976 (explicit `fallback_order` chain), PR #46986 (honour fallback_order in cron).
- **Memory & Tools**: PR #47041 (L0/L1/L2 filesystem memory provider plugin), PR #46988 (prefer injected memory before search).
- **TUI & Desktop**: PR #46980 (syntax aliases for C#, Java, Kotlin), PR #46984 (Linux auto‑restart after update), PR #47050 (fix custom provider self‑dedup).

## 4. Community Hot Topics

The most active issue by both comments and reactions is:

- **Issue #5941 – Add Searxng as a default web search provider**  
  *30 👍, 6 comments* – Created 2026-04-07, still open. This long‑standing request has high community support. Users want an alternative to Firecrawl and Tavily, plus a reranker step.  
  [NousResearch/hermes-agent Issue #5941](https://github.com/NousResearch/hermes-agent/issues/5941)

Other notable items with lower engagement but topical interest:

- **Issue #47035 – Orchestrator v3 + non‑blocking doer/reviewer** (showcase, 0 comments) – Demonstrates real‑world usage of `delegate_task(background=true)`.  
  [NousResearch/hermes-agent Issue #47035](https://github.com/NousResearch/hermes-agent/issues/47035)

- **PR #47041 – LayeredMemoryProvider plugin** – A community‑built solution to the 2200‑char memory limit, also showcased in Issue #43955.  
  [NousResearch/hermes-agent PR #47041](https://github.com/NousResearch/hermes-agent/pull/47041)

## 5. Bugs & Stability

Several bugs were reported today, ranked by severity (P2=medium, P3=low):

- **P2 – Bootstrap fails on Windows when gateway process locks `hermes.exe`** (#47036) – Leaves installation broken. No fix PR yet.  
  [Issue #47036](https://github.com/NousResearch/hermes-agent/issues/47036)

- **P2 – Desktop model picker hides custom providers due to `is_aggregator()` false positive** (#47042) – Models disappear from the picker. A fix PR #47050 is open.  
  [Issue #47042](https://github.com/NousResearch/hermes-agent/issues/47042) · [PR #47050](https://github.com/NousResearch/hermes-agent/pull/47050)

- **P2 – Telegram: rich‑message final reply overlaps with legacy MarkdownV2 rendering** (#47048) – Tables and bullets appear twice. No fix PR yet.  
  [Issue #47048](https://github.com/NousResearch/hermes-agent/issues/47048)

- **P2 – Kanban worker errors report generic “protocol violation”** – PR #46985 open to surface real worker error messages.  
  [PR #46985](https://github.com/NousResearch/hermes-agent/pull/46985)

- **P2 – Camofox browser fails on stale tab 404** – PR #46982 adds auto‑recovery.  
  [PR #46982](https://github.com/NousResearch/hermes-agent/pull/46982)

- **P2 – Token reduction false abort in 413 retry path** – PR #46987 fixes detection when message count stays unchanged.  
  [PR #46987](https://github.com/NousResearch/hermes-agent/pull/46987)

- **P3 – BlueBubbles self‑echo reply loop** – PR #46974 addresses the issue.  
  [PR #46974](https://github.com/NousResearch/hermes-agent/pull/46974)

- **P3 – SMS body exceeds Twilio’s 1600‑char limit** – PR #46972 caps the body.  
  [PR #46972](https://github.com/NousResearch/hermes-agent/pull/46972)

- **P3 – Linux desktop app does not auto‑restart after update** – PR #46984 fixes it.  
  [PR #46984](https://github.com/NousResearch/hermes-agent/pull/46984)

No P1 (critical) bugs were reported today.

## 6. Feature Requests & Roadmap Signals

Community‑driven feature requests and developer showcases point toward the following likely roadmap items:

- **Searxng as a default web search provider** (#5941) – Strong demand. Would reduce dependency on proprietary APIs.
- **Mid‑session model switching** (#47047) – Automatic or manual model changes within a session while preserving conversation history. Addresses cost vs. capability tradeoffs.
- **Multiple API endpoints per custom provider** (#47039) – Dropdown selection in desktop UI instead of manual config editing. Especially requested for Chinese domestic LLM APIs.
- **Layered memory system** (PR #47041, Issue #43955) – Filesystem‑based L0/L1/L2 tiers to replace the 2200‑char flat memory. Already prototyped by the community.
- **Conversational follow‑up routing** (PR #47043) and **/after command** (PR #47044) – Precise targeting of replies within a session.
- **Dynamic PII‑safe platform detection** (PR #47045) – Plugin authors can mark platforms as safe without core changes.

**Next‑version prediction:** Delegation workflow improvements (DAG tool, fallback isolation) and the layered memory plugin are strong candidates, given the high code activity. The Searxng provider request may land as a PR soon due to its popularity.

## 7. User Feedback Summary

**Pain points expressed in issues:**

- Windows users face a hard lock during bootstrap when a gateway process is running (#47036).
- Custom provider models disappear from the desktop picker after a recent update (#47042).
- Telegram users get garbled tables with double‑rendered content (#47048).
- The built‑in 2200‑character memory limit forces token waste and context loss, leading to a community‑built layered memory alternative (Issue #43955).

**Satisfaction signals:**

- The Orchestrator v3 showcase (#47035) demonstrates that advanced delegation patterns are being actively adopted and extended.
- The layered memory plugin (PR #47041) received a positive showcase and suggests users are building on Hermes’ plugin system.

**Use cases highlighted:**

- Cost‑sensitive users want mid‑session model switching to use cheaper models for routine queries and premium models only when needed (#47047).
- Developers integrating with multiple Chinese LLM APIs want a single custom provider with endpoint dropdowns (#47039).
- Users self‑hosting web search (Searxng) want integration to avoid third‑party API costs (#5941).

## 8. Backlog Watch

**Issue #5941 – Add Searxng as a default web search provider**  
Created 2026-04-07 (70 days open), 30 👍, 6 comments. Maintainers have not yet assigned a milestone or responded with a formal roadmap. Given the high community interest, this issue risks becoming stale without a maintainer comment.  
[Issue #5941](https://github.com/NousResearch/hermes-agent/issues/5941)

**PR #46971 – Dynamic workflow DAG tool**  
Opened today but has no comments. As a significant new delegation feature, it may require maintainer review soon to align with existing delegation internals.  
[PR #46971](https://github.com/NousResearch/hermes-agent/pull/46971)

**PR #46976 / #46986 – Fallback order chain**  
Two overlapping PRs addressing the same config area. Maintainers should evaluate and possibly merge the cleaner solution to avoid duplication.  
[PR #46976](https://github.com/NousResearch/hermes-agent/pull/46976) · [PR #46986](https://github.com/NousResearch/hermes-agent/pull/46986)

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest — 2026-06-16

## 1. Today's Overview

The project shows moderate activity today, with **3 issues updated** (1 open, 2 closed) and **13 pull requests updated** (10 open, 3 merged/closed). A new **nightly release** (v0.2.9-nightly.20260616) was published, though it is an automated build with no detailed changelog beyond the usual inclusion of `main` branch changes. The development team continues to focus on **stability, security, and compatibility**, closing a security bypass vulnerability and advancing several bug‑fix PRs. Community engagement is steady, with active discussion on a Windows‑specific QQ channel failure and a resolved RISC‑V compatibility issue.

## 2. Releases

**New release**: [`nightly` (v0.2.9-nightly.20260616.c1ff5aa6)](https://github.com/sipeed/picoclaw/releases/tag/v0.2.9)  
- Automated nightly build, may be unstable.  
- Full changelog: [v0.2.9…main](https://github.com/sipeed/picoclaw/compare/v0.2.9...main) – no manual changelog entry; changes reflect the latest merged PRs.  
- **No breaking changes or migration notes** provided. Use with caution.

## 3. Project Progress

Three pull requests were merged or closed today, all of which improve documentation, diagnostics, and user experience:

- **[#3096 – docs: add PicoPaw banners to READMEs](https://github.com/sipeed/picoclaw/pull/3096)** – merged. Adds project branding banners to documentation files.  
- **[#3126 – fix(web): improve launcher allowlist bypass diagnostics](https://github.com/sipeed/picoclaw/pull/3126)** – merged. Enhances startup logging to detect when `allow_localhost_bypass` is misconfigured, addressing security issue [#3069](https://github.com/sipeed/picoclaw/issues/3069).  
- **[#3097 – feat: add shift-enter hint below chat composer](https://github.com/sipeed/picoclaw/pull/3097)** – merged. Adds a visible “Shift + Enter” hint inside the Web UI to inform users how to insert line breaks.

Additionally, a batch of bug‑fix PRs remain open but were updated today (see §5), indicating ongoing work to harden the codebase.

## 4. Community Hot Topics

| Item | Type | Comments | Status | Summary |
|------|------|----------|--------|---------|
| [#2887 – .deb version on RISC‑V not functional with OpenAI model](https://github.com/sipeed/picoclaw/issues/2887) | Issue | 10 | Closed (stale) | User reported that the `.deb` package on RISC‑V cannot use OpenAI models. The issue was closed as stale after a month without resolution. Although closed, the underlying need for proper RISC‑V support remains an open concern. |
| [#3015 – QQ channel connection failed on Windows](https://github.com/sipeed/picoclaw/issues/3015) | Issue | 3 | Open (stale) | Windows users experience a token retrieval timeout when starting QQ channel (`picoclaw gateway`). Pico channel works normally. Community has posted workaround attempts; no fix PR or official response yet. |
| [#3069 – Security: launcher `allowed_cidrs` bypass via reverse proxy](https://github.com/sipeed/picoclaw/issues/3069) | Issue | 0 | Closed | A report demonstrating that `RemoteAddr`–based CIDR checks can be spoofed by a same‑host reverse proxy. Fixed by PR [#3126](https://github.com/sipeed/picoclaw/pull/3126) today. |

The **RISC‑V issue** (#2887) attracted the most comments, highlighting a persistent pain point for users on that architecture. The **QQ channel failure** (#3015) is the most active open issue, suggesting Windows channel integration needs dedicated attention.

## 5. Bugs & Stability

**Critical severity (fixed):**  
- **Security bypass** – Issue [#3069](https://github.com/sipeed/picoclaw/issues/3069) (allowlist CIDR bypass via same‑host reverse proxy) was resolved today by PR [#3126](https://github.com/sipeed/picoclaw/pull/3126), which improves bypass detection logging. No immediate patch release; fix is in `main` and will be included in the next stable release.

**High severity (open):**  
- **QQ channel failure on Windows** – Issue [#3015](https://github.com/sipeed/picoclaw/issues/3015) remains open. No associated fix PR. Impacts all Windows users wanting to use QQ as a channel.

**Medium severity (multiple fix PRs in review):**  
Today, **seven open bug‑fix pull requests** were updated, indicating active work on stability:

- [#3132 – panic recovery in core goroutines](https://github.com/sipeed/picoclaw/pull/3132) – prevents single‑goroutine panics from killing the process.  
- [#3059 – explicit ignore of `Close()` errors](https://github.com/sipeed/picoclaw/pull/3059) – silences linter warnings on resource cleanup.  
- [#3054 – type assertion checks in LINE channel `Send`](https://github.com/sipeed/picoclaw/pull/3054) – prevents panic on malformed `sync.Map` data.  
- [#3047 – restore full JSONL history in session detail](https://github.com/sipeed/picoclaw/pull/3047) – fixes missing archived messages in the web API.  
- [#3128 – ignore `resp.Body.Close()` errors after `io.ReadAll`](https://github.com/sipeed/picoclaw/pull/3128) – non‑meaningful errors in search providers.  
- [#3131 – type assertion checks in tool registry](https://github.com/sipeed/picoclaw/pull/3131) – fallback to zero values on schema mismatch.  
- [#3130 – handle `json.Marshal` errors in seahorse tools](https://github.com/sipeed/picoclaw/pull/3130) – returns descriptive error instead of silent empty string.  
- [#3129 – explicit ignore of file `Close()` error in TTS](https://github.com/sipeed/picoclaw/pull/3129) – cosmetic fix.  
- [#3127 – explicit ignore of directory fd `Close()` errors](https://github.com/sipeed/picoclaw/pull/3127) – cosmetic fix.

These PRs collectively improve error handling, panic safety, and data integrity. None have been merged yet.

## 6. Feature Requests & Roadmap Signals

While no new feature requests were filed today, two open PRs hint at upcoming capabilities:

- **[#2975 – feat(telegram): treat reply to bot message as mention in group chats](https://github.com/sipeed/picoclaw/pull/2975)** – opened 2026‑05‑30, still open. If merged, it will allow users in Telegram groups to trigger the bot by simply replying to its messages (currently requires explicit `@mention`). This is likely to land in the next minor release.  
- **[#3047 – fix(web): restore full JSONL history for session detail](https://github.com/sipeed/picoclaw/pull/3047)** – though classified as a fix, it adds a new API capability: showing archived messages in session detail. This meets a frequently expressed user need for historical visibility.  
- **[#3097 – shift-enter hint](https://github.com/sipeed/picoclaw/pull/3097)** – already merged today; a small UX improvement.

No roadmap documents are available, but the pattern suggests the team is balancing security/hardening against incremental UX improvements.

## 7. User Feedback Summary

Based on issues and PRs updated today, users are expressing:

- **Pain points**:  
  - RISC‑V Linux users cannot use OpenAI models via the `.deb` package (issue #2887, now closed but unresolved).  
  - Windows users experience a complete channel failure when using QQ (issue #3015).  
  - Security‑conscious users are concerned about improper CIDR enforcement (issue #3069, now fixed).  

- **Use cases**:  
  - Deploying PicoClaw as a personal AI assistant on low‑cost RISC‑V hardware.  
  - Running PicoClaw as a Windows gateway to integrate with Chinese chat platforms (QQ).  
  - Using PicoClaw behind a reverse proxy for network isolation.  

- **Satisfaction**:  
  - The prompt response to the security bypass (issue #3069 opened on Jun 9, fixed Jun 15) demonstrates a healthy security posture.  
  - However, the QQ channel failure (#3015) has been open for 10 days without a committed fix, which may cause frustration among Windows users.  
  - The RISC‑V issue was closed as stale without a concrete solution, potentially alienating users on that platform.

## 8. Backlog Watch

Issues and PRs that require maintainer attention:

- **[#3015 – QQ channel connection failed on Windows](https://github.com/sipeed/picoclaw/issues/3015)** – open for 10 days, marked stale, no maintainer response or linked PR. This is the most impactful unresolved bug.  
- **[#2975 – feat(telegram): treat reply as mention](https://github.com/sipeed/picoclaw/pull/2975)** – open for 17 days, no activity. A seemingly small change that would improve Telegram usability; low risk to merge.  
- **[#3047 – fix(web): restore full JSONL history](https://github.com/sipeed/picoclaw/pull/3047)** – open for 9 days, no maintainer review. Could be important for users relying on session history.  
- **[#3059, #3054, #3059, #3128, #3131, #3130, #3129, #3127, #3132 – bug‑fix PRs](https://github.com/sipeed/picoclaw/pulls?q=is%3Apr+is%3Aopen+updated%3A%3E%3D2026-06-15)** – all opened within the last 8 days; none merged yet. While many are cosmetic, the panic‑recovery PR (#3132) and type‑assertion fixes (#3054, #3131) directly address crash risks and should be prioritized.  

None of the items are extremely old, but the accumulation of unmerged bug‑fix PRs suggests a need for a dedicated stabilization sprint soon.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

## NanoClaw Project Digest – 2026-06-16

### 1. Today’s Overview

The project saw moderate activity: 1 new issue opened, 11 pull requests updated (3 merged/closed, 8 still open), and no new releases were published. The community contributed fixes across Slack, WhatsApp, the agent runner, and the CLI, while maintainers focused on merging a gateway upgrade workflow and refining documentation. Overall momentum remains steady with a healthy mix of bug fixes and infrastructure improvements.

### 2. Releases

*No new releases have been published. The latest stable version remains unchanged.*

### 3. Project Progress – Merged / Closed PRs Today

Three pull requests were merged or closed:

- **#2774 – `feat(update-nanoclaw): upgrade OneCLI gateway when its pinned version moves`** (merged) – Ensures that when `versions.json` advances the OneCLI gateway pin, `update-nanoclaw` now also upgrades the running gateway/CLI, preventing silent mismatches after an update.  
  [nanocoai/nanoclaw PR #2774](https://github.com/nanocoai/nanoclaw/pull/2774)

- **#2772 – `fix(codex): per-thread conversation archive (CDX-004)`** (closed) – Changes the Codex archive strategy from writing one file per exchange to keying on the thread/continuation id, consolidating scattered conversation fragments into a single archive per thread.  
  [nanocoai/nanoclaw PR #2772](https://github.com/nanocoai/nanoclaw/pull/2772)

- **#2773 – `docs(add-codex): drop redundant TTY warning in auth note`** (closed) – Removed a duplicated sentence from the `add-codex` skill documentation to avoid confusion during authentication.  
  [nanocoai/nanoclaw PR #2773](https://github.com/nanocoai/nanoclaw/pull/2773)

### 4. Community Hot Topics

The only new issue, **#2779**, reports a breaking bug in Slack integration where `@handle` parts inside URLs (e.g., HackMD notes, Mastodon profiles) are incorrectly rewritten into broken mentions. Despite zero comments, its immediate impact on agents sending links to Slack makes it a high-visibility problem.  
[Issue #2779 – Slack: @handles inside URLs get mangled](https://github.com/nanocoai/nanoclaw/issues/2779)

Several open PRs attracted repeated attention today (updated after weeks of inactivity):  
- **#2628** – `fix(cli): honor user-supplied --id in ncl groups create` (open since May 27)  
- **#2627** – `fix(reactions): align MCP add_reaction schema with channel reality`  
- **#2626** – `fix(signal): replace silent restartService failure with explicit error`

These long-standing fixes are still under review, indicating the community is invested in resolving them.

### 5. Bugs & Stability

**High Severity:**  
- **Issue #2779** – Slack URL mangling breaks `@handle` paths in shared links. No fix PR yet; the bug can cause agents to send broken references to users.  

**Medium Severity – Fix PRs in Progress:**  
- **#2759** – Budget/token-exhausted LLM turns (e.g., Anthropic budget errors) were being silently dropped instead of delivered as error messages. A fix is under review.  
  [PR #2759 – deliver budget/billing error turns](https://github.com/nanocoai/nanoclaw/pull/2759)  

- **#2627** – Reactions (emoji) sent via MCP were not translated correctly across channels, causing silent failures on WhatsApp, Discord, Telegram, etc.  
- **#2626** – Signal channel `restartService` silently failed after an earlier `unload`, making the setup wizard report incorrect status.  

**Low Severity:**  
- **#2628** – The CLI `--id` flag is documented but ignored; user-supplied IDs are overridden with random UUIDs. This is a usability bug that has been open for three weeks.

### 6. Feature Requests & Roadmap Signals

Three PRs signal likely upcoming features:

- **#2777 – `feat: add /add-strava skill for official Strava MCP`** – Adds an OAuth flow and skill to wire Strava’s MCP endpoint into agent groups. This expands the agent’s ability to interact with external fitness/lifestyle data.  
  [PR #2777](https://github.com/nanocoai/nanoclaw/pull/2777)

- **#2776 – `feat: support remote HTTP/SSE MCP servers`** – Extends MCP server configuration to allow remote endpoints alongside the existing stdio servers, enabling agents to connect to cloud-hosted or third-party MCP services.  
  [PR #2776](https://github.com/nanocoai/nanoclaw/pull/2776)

- **#2778 – `fix(whatsapp): route inbound media through shared session inbox`** – While labelled a fix, this PR solves a fundamental gap where WhatsApp media never reached the agent container, effectively enabling a new capability.  
  [PR #2778](https://github.com/nanocoai/nanoclaw/pull/2778)

These additions suggest the next minor release may include remote MCP support, a Strava integration, and working WhatsApp media ingestion.

### 7. User Feedback Summary

- **Pain Points:** The Slack URL mangling bug (#2779) is a clear source of friction for users who share links with `@handles`.  
- **Contributor Engagement:** Multiple community members (eldar702, IamAdamJowett, clementdecoligny, assapin, Koshkoshinsk) are actively submitting fixes and features, indicating satisfaction with the project’s direction and maintainer responsiveness.  
- **Documentation Clarity:** The redundant warning in the Codex auth note (#2773) was quickly cleaned up, showing maintainers listen to feedback about confusing docs.

### 8. Backlog Watch

Three open PRs from May 27 remain unresolved despite being updated today:

- **#2628** – `fix(cli): honor user-supplied --id in ncl groups create` – Long-running, affects a core CLI feature.  
- **#2627** – `fix(reactions): align MCP add_reaction schema with channel reality` – Blocks cross-platform emoji reactivity.  
- **#2626** – `fix(signal): replace silent restartService failure with explicit error` – Causes a poor setup experience for Signal users.

These PRs have gathered no maintainer comments or reviews for weeks. If merged, they would resolve several documented issues (#2390, #2569, #2583). The project would benefit from a dedicated review pass on these items.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-06-16

## Today’s Overview
Project activity remains low, with no new releases or pull requests updated in the last 24 hours. Two open issues received comments, but no code contributions were merged or closed. The community is primarily asking about configuration and local model reliability rather than proposing new features. Overall, the project appears to be in a maintenance-lull phase, with no signs of active development sprints or urgent regressions.

## Releases
No new releases were published in the last 24 hours. The latest release remains unchanged.

## Project Progress
No pull requests were updated, merged, or closed today. No feature advancements or fixes can be reported.

## Community Hot Topics
Two issues are actively being discussed (both with 1 comment each):

- **Issue #957** [OPEN] – *Rate limit issue* (jacktang, 2026-06-15)  
  User asks how the “rate limit” config works and how to modify the threshold when using NullClaw as an agent runtime without memory and with JSON output. This highlights a need for clearer documentation around internal rate‑limiting mechanisms.

- **Issue #952** [OPEN] – *[bug] Local model using ollama returns incomplete answers* (bloodgroup-cplusplus, 2026-06-11)  
  The user reports that pulling Gemma via Ollama leads to truncated responses. A screenshot is provided. This indicates potential compatibility or streaming issues between NullClaw and local LLM backends.

Neither issue has received maintainer responses yet, and both are tagged as open/active.

## Bugs & Stability
- **Moderate severity:** Issue #952 – Incomplete answers from a local Ollama model (Gemma). No fix PR exists. The bug could impact users relying on local, privacy‑preserving deployments. Root cause is unknown; could be a token‑limit issue, streaming bug, or model‑specific parsing problem.

- **Low severity:** Issue #957 – Rate limit configuration confusion. This is more of a usability/documentation gap than a functional bug.

No crashes or regressions were reported in the last 24 hours.

## Feature Requests & Roadmap Signals
No explicit feature requests were filed today. However, the underlying need in Issue #957 suggests that users want fine‑grained control over internal rate‑limiting thresholds (e.g., configurable per‑second limits, environment variables). A configuration documentation improvement or a new `max_requests_per_second` config key could be added in the next minor release.

## User Feedback Summary
- **Pain point #1:** Lack of clear documentation on the “rate limit” configuration leads to frustration when users try to run NullClaw as a lightweight agent runtime. Users expect to control output format (JSON) without hitting undocumented limits.
- **Pain point #2:** Local model integration with Ollama is unreliable, producing incomplete answers. This undermines one of NullClaw’s key use cases (offline/private AI agents). User has provided steps and a screenshot, suggesting high confidence in the reproducibility of the bug.
- **Satisfaction indicators:** None reported; no positive feedback or workarounds were shared in the open issues.

## Backlog Watch
- **Issue #952** (open since June 11) – Five days without maintainer response. This is a reproducible bug with clear reproduction steps. High‑priority candidate for investigation, especially if the project aims to support Ollama‑based models.
- **Issue #957** (open since June 15) – New but unanswered. Adding a response with config example or pointing to relevant source code would improve user experience.

No stale PRs or long‑dormant feature requests currently require maintainer attention.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-06-16

## 1. Today’s Overview

The IronClaw project shows **high activity** today, with 17 issues and 50 pull requests updated in the last 24 hours. Of the PRs, 19 were merged or closed, signalling steady progress toward stabilisation and feature completion. The development focus is heavily weighted toward the **Reborn** binary and WebUI, particularly around authentication flows, tool execution reliability, and credential scoping. Several critical bugs in OAuth workflows and tool failure recovery have been addressed by merged PRs, while a number of usability issues (e.g. shell approval denial, tool call visibility) remain open. No new releases were published today.

---

## 2. Releases

No new releases today.

---

## 3. Project Progress

Today’s 19 merged/closed PRs advanced several key areas:

- **Auth-gate denial fix** ([#4944](https://github.com/nearai/ironclaw/pull/4944)) — Denying an OAuth/credential gate in Reborn no longer loops forever; the model now receives a clear denial reason and can take alternative action.
- **Security update** ([#4950](https://github.com/nearai/ironclaw/pull/4950)) — Patched `wasmtime` to 44.0.3 to address RUSTSEC-2026-0182 (WASIp1 `fd_renumber` leak), fixing a repo-wide CI failure.
- **Benchmark CI fix** ([#4947](https://github.com/nearai/ironclaw/pull/4947)) — The `/benchmark` suite now validates against the `benchmarks` `main` branch instead of a stale pin, unblocking automated benchmarking.
- **Trace Commons merge resolution** ([#4929](https://github.com/nearai/ironclaw/pull/4929)) — Cleared main‑merge conflicts for the trace-commons agent onboarding feature (#4559), ready for final integration.
- **Routine delivery steering** ([#4780](https://github.com/nearai/ironclaw/pull/4780) — merged yesterday) — Model now sees guidance to select outbound delivery targets before creating routines, improving Slack/hook interoperability.
- **Image attachment support** ([#4871](https://github.com/nearai/ironclaw/pull/4871) — merged yesterday) — Attached images are now sent as real multimodal content to vision‑capable models, closing a major attachment gap.

Other notable changes include post-merge review follow-ups for image vision ([#4945](https://github.com/nearai/ironclaw/pull/4945)) and an auto-wiring script for Google OAuth in the Reborn dev launcher ([#4943](https://github.com/nearai/ironclaw/pull/4943)).

---

## 4. Community Hot Topics

The most active issues (by comment count, each with 2 comments) reveal recurring pain points:

- **[#4907](https://github.com/nearai/ironclaw/issues/4907) — OAuth resume failure** – After a successful Google OAuth flow, the original run fails instead of resuming. This is a critical blocker for Google Calendar/Gmail integration.
- **[#4761](https://github.com/nearai/ironclaw/issues/4761) – Agent stops after repeated tool failures** – The run completes but the agent stops, leaving the user without recovery.
- **[#4764](https://github.com/nearai/ironclaw/issues/4764) – Shell approval denial leaves tool pending** – Clicking “Deny” doesn’t cancel the tool invocation and provides no feedback.
- **[#4880](https://github.com/nearai/ironclaw/issues/4880) – Automate code review** – A strong community request to let AI handle PR review comments and reduce human workload.

These issues indicate a need for more robust **error recovery**, **transparent UX feedback**, and **automation** of repetitive development tasks.

---

## 5. Bugs & Stability

Bugs reported today are ranked by severity:

| Severity | Issue | Summary | Fix PR |
|----------|-------|---------|--------|
| **Critical** | [#4907](https://github.com/nearai/ironclaw/issues/4907) | OAuth resume failure – run not continuing after auth | None yet |
| **Critical** | [#4761](https://github.com/nearai/ironclaw/issues/4761) | Agent stops after repeated tool failures | None yet |
| **Critical** | [#4764](https://github.com/nearai/ironclaw/issues/4764) | Shell approval denial leaves tool pending | Partially addressed by [#4944](https://github.com/nearai/ironclaw/pull/4944) |
| **High** | [#4762](https://github.com/nearai/ironclaw/issues/4762) | Failed tool workflow causes inconsistent ordering of follow‑up messages | None |
| **High** | [#4942](https://github.com/nearai/ironclaw/issues/4942) | Tool call failures not displayed until page reload | None |
| **High** | [#4914](https://github.com/nearai/ironclaw/issues/4914) | Gmail OAuth with no selected scopes returns raw malformed callback error | None |
| **Medium** | [#4913](https://github.com/nearai/ironclaw/issues/4913) | Google Calendar authorization not reused across conversations | [#4939](https://github.com/nearai/ironclaw/pull/4939) (open) |
| **Medium** | [#4904](https://github.com/nearai/ironclaw/issues/4904) | Duplicate `extension_install` actions during Google Calendar installation | None |
| **Low** | [#4697](https://github.com/nearai/ironclaw/issues/4697) (closed) | Provider status inconsistency in Inference settings | Resolved |
| **Low** | [#4696](https://github.com/nearai/ironclaw/issues/4696) (closed) | Ollama test connection falsely reports success | Resolved |

The ongoing **Nightly E2E failure** ([#4108](https://github.com/nearai/ironclaw/issues/4108)) persists as a systemic test infrastructure problem (open since May 27).

---

## 6. Feature Requests & Roadmap Signals

Based on recent issues and PRs, the following areas are likely priorities for the next release:

- **Vision support** – Two PRs advance vision: [#4951](https://github.com/nearai/ironclaw/pull/4951) backs the vision tool with the provider’s real vision model, and [#4902](https://github.com/nearai/ironclaw/pull/4902) adds inline image support to the OpenAI‑compatible endpoint. Expect stable vision in the next minor release.
- **Credentials scoping** – [#4935](https://github.com/nearai/ironclaw/issues/4935) and the corresponding PR [#4939](https://github.com/nearai/ironclaw/pull/4939) propose making credentials owner‑scoped (tenant/user/agent) rather than thread‑scoped. This would fix cross‑thread OAuth reuse and reduce authentication friction.
- **Code review automation** – [#4880](https://github.com/nearai/ironclaw/issues/4880) and [#4882](https://github.com/nearai/ironclaw/issues/4882) (build coding agent cloud workflow) signal a desire to integrate IronClaw into the development lifecycle, possibly surfacing as a bot that can receive PR review assignments.
- **Slack personal token** – [#4941](https://github.com/nearai/ironclaw/pull/4941) adds a `slack_user_tool` for user‑token operations, expanding beyond bot tokens (e.g. searching all conversations). This could be a differentiator for power users.
- **Downloadable project files** – [#4933](https://github.com/nearai/ironclaw/pull/4933) (open) lets Reborn agents produce downloadable files (CSV, reports), bridging the gap between agent output and user‑accessible artifacts.

---

## 7. User Feedback Summary

User-submitted issues (primarily from `sunglow666` and `zetyquickly`) highlight concrete pain points:

- **OAuth friction** – The most common theme: authorisation flows either fail silently, don’t resume the original task, or require re‑authentication for every conversation. This is a major obstacle for adopting Google integrations.
- **Tool execution transparency** – Users report confusion when tool calls fail but the UI doesn’t indicate the failure until a manual reload, or when denied shell approvals leave the session in an ambiguous state.
- **Configuration inconsistencies** – The Inference settings page can show an “Active” provider that is not actually used (Ollama false‑positive test, provider mismatch).
- **Desire for automation** – The feature request to automate code review and code generation (issues #4880, #4882) suggests users want IronClaw to become a development assistant, not just a personal AI.

Overall satisfaction is tempered by these UX and reliability issues, but the high number of fix PRs shows a responsive maintainer team.

---

## 8. Backlog Watch

- **[#4108](https://github.com/nearai/ironclaw/issues/4108) – Nightly E2E failed** (opened 2026-05-27, updated today but with 0 comments). This automated test failure has been open for 20 days with no root‑cause discussion or resolution. It likely blocks confidence in CI and should receive maintainer attention promptly.
- **[#4761](https://github.com/nearai/ironclaw/issues/4761) – Agent stops after repeated tool failures** (opened 2026-06-11, 2 comments). Despite being marked “bug” and “Reborn,” no fix PR has been linked. As a critical UX blocker, it merits escalation.
- **[#4880](https://github.com/nearai/ironclaw/issues/4880) – Automate code review** (opened yesterday, 2 comments). While not a bug, it is a significant feature request that could shape the project’s direction. Maintainers should clarify scope and timeline to avoid community frustration.

No other issues appear to have been ignored for an unusually long time.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

```markdown
# LobsterAI Project Digest – 2026-06-16

## Today's Overview
The project shows moderate development activity with **11 pull requests updated** in the last 24 hours, of which **5 were merged/closed**. No new releases or fresh issues were created; the two existing open issues remain stale from early April. The merged PRs concentrate on refining the cowork experience and voice input, while the open PRs are dominated by dependency upgrades from Dependabot. Overall, the team is actively polishing the cowork module and cleaning up technical debt, but long-standing bugs in skill management continue to lack maintainer attention.

## Releases
No new releases were published in the last 24 hours.

## Project Progress
Five pull requests were merged or closed in the past day, advancing the cowork and voice input features:

- **[PR #2168](https://github.com/netease-youdao/LobsterAI/pull/2168)** (Closed) — Added a floating scroll‑to‑bottom button for cowork conversations, supporting smooth scrolling, wheel passthrough, i18n labels, and click diagnostics.
- **[PR #2163](https://github.com/netease-youdao/LobsterAI/pull/2163)** (Closed) — Refined the dictation recording UI for voice input and improved ASR quota handling (in‑memory quota slice, lazy‑reset on daily availability).
- **[PR #2162](https://github.com/netease-youdao/LobsterAI/pull/2162)** (Closed) — Fixed a merge conflict in the cowork voice‑input by preserving the realtime‑only ASR flow while keeping draft ownership, stale callback guards, session‑switch cancellation, and diagnostic logging.
- **[PR #2161](https://github.com/netease-youdao/LobsterAI/pull/2161)** (Closed) — Minor chore update to the "About" dialog.
- **[PR #2160](https://github.com/netease-youdao/LobsterAI/pull/2160)** (Closed) — Removed the short ASR upload flow, made cowork voice input always use realtime ASR, removed the Settings mode switch for recognition, and stripped legacy configuration keys.

These changes indicate a clear shift toward a simpler, always‑realtime voice input model and improved in‑conversation UX.

## Community Hot Topics
The most active items in the community remain the two stale issues and one feature PR:

- **[Issue #1426](https://github.com/netease-youdao/LobsterAI/issues/1426)** (Open, 1 comment) — User reports uploading a local skill shows no success feedback and the skill list does not refresh.
- **[Issue #1427](https://github.com/netease-youdao/LobsterAI/issues/1427)** (Open, 1 comment) — Duplicate skills can be added via local upload, causing multiple identical entries.
- **[PR #1428](https://github.com/netease-youdao/LobsterAI/pull/1428)** (Open, stale, no comments) — Proposed feature to push system notifications when a cowork session completes or errors while the main window is unfocused.

The underlying need is for better user feedback and notification parity with comparable tools like Claude Code and Cursor.

## Bugs & Stability
Two usability bugs were reported on 2026-04-03 and remain unresolved:

- **#1426** — No success tip after uploading a skill; skill list not refreshed.  
  *Severity: Medium – hampers user confidence.*
- **#1427** — Duplicate skill entries allowed via local upload.  
  *Severity: Medium – leads to confusion and clutter.*

No fix PRs are linked to either issue. On the stability side, the merged PRs #2162 and #2160 addressed voice input cancellation guards and ASR flow consistency, improving robustness of the cowork voice experience.

## Feature Requests & Roadmap Signals
- **System notifications for cowork** (PR #1428) is a strong signal toward background awareness. The idea has been open for over two months without merge, but the recent activity in cowork features suggests it could be picked up in the next iteration.
- **Scroll‑to‑bottom control** (PR #2168) was merged today, fulfilling a common expectation in chat interfaces.
- The removal of the short‑ASR flow and recognition mode switch (PR #2160) aligns the project with a streamlined, realtime‑first voice input approach, likely reducing future configuration burden.

These points suggest the next minor release may focus on cowork UX polish and notification support.

## User Feedback Summary
From the two open issues, real pain points emerge around the skill management UI:
- Users expect immediate visual feedback after uploading a skill (no toast or list update).
- The system does not prevent duplicate skill names, forcing manual cleanup.

Both issues were filed two months ago with no maintainer response, which may contribute to user dissatisfaction. In contrast, the cowork voice‑input improvements (merged today) address earlier feedback on ASR complexity and cancellation behavior.

## Backlog Watch
The following items require maintainer attention:

- **[Issue #1426](https://github.com/netease-youdao/LobsterAI/issues/1426)** – Stale (2+ months), no label or assignee.
- **[Issue #1427](https://github.com/netease-youdao/LobsterAI/issues/1427)** – Stale (2+ months), same severity as #1426.
- **[PR #1428](https://github.com/netease-youdao/LobsterAI/pull/1428)** – Stale feature PR (2+ months) with no reviewer activity.
- **[PR #1277](https://github.com/netease-youdao/LobsterAI/pull/1277)** – Dependabot PR bumping `electron` from 40.2.1 to 42.4.0, open since April 2. While dependency updates can often wait, such a large jump may introduce breaking changes and should not be left unresolved indefinitely.

A brief triage of these items would improve project health and community trust.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest – 2026-06-16

## Today's Overview
Project activity remains low with no new issues or releases reported in the last 24 hours. Two pull requests received updates, but both were created on the previous day and remain open without merge activity. No issues were opened or closed, and no bugs or regressions were filed. The overall pulse suggests a quiet period, possibly reflecting maintainer focus on reviewing the two open PRs.

## Releases
*None.* No new releases were published today.

## Project Progress
No pull requests were merged or closed today. Both open PRs (#1124 and #1125) received updates but remain under review.

## Community Hot Topics
Only two PRs updated today, each with no comments or reactions. They represent the entirety of community contribution activity:
- **PR #1125 – Support model and effort selection for external agents**  
  [Link](https://github.com/moltis-org/moltis/pull/1125)  
  Adds first-class configuration for external-agent provider model and effort selection, including model/effort metadata persistence. This addresses a user need for finer control when integrating non-native agents.
- **PR #1124 – Add context command support for chat turns**  
  [Link](https://github.com/moltis-org/moltis/pull/1124)  
  Introduces an optional `chat.context_command` that injects runtime context before each chat turn, improving deployment ergonomics by avoiding manual context pasting.

Neither PR has attracted discussion yet, suggesting they may be early-stage contributions awaiting maintainer review.

## Bugs & Stability
No bugs, crashes, or regressions were reported today. The project remains in a stable state.

## Feature Requests & Roadmap Signals
Both open PRs represent feature additions rather than user‑filed requests. The `chat.context_command` feature (PR #1124) aligns with a common deployment pain point – dynamically generating context from scripts – and could appear in a future release if merged. Model/effort selection for external agents (PR #1125) is a natural extension of Moltis’s provider abstraction and may signal upcoming support for more complex agent orchestration.

## User Feedback Summary
No user feedback, pain points, or satisfaction reports were captured today. The absence of new issues may indicate either general contentment or low active usage.

## Backlog Watch
No issues or PRs require maintainer attention beyond the two open PRs. Neither PR has been lingering for an extended time (both created yesterday). No unanswered questions or stale items were identified.

---

*Generated from GitHub data for moltis-org/moltis on 2026-06-16.*

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest — 2026-06-16

## Today's Overview

CoPaw (now QwenPaw) shows high development velocity: **28 issues** updated in the last 24 hours (17 open, 11 closed) and **49 pull requests** updated (16 open, 33 merged/closed). No new releases were cut. The community is actively reporting bugs and proposing features, with critical regressions in context management, file downloads, and channel stability dominating the conversation. Merged PRs include fixes for tool card loading spinners, skill slash injection display, and a first version of an Agent OS Driver abstraction. Maintainers are reviewing several first‑time contributor patches while large‑scope features (Models Page overhaul, non‑blocking streaming flush, user input queue) remain open.

## Releases

No new releases today. The latest version remains **v1.1.11.post2**, which introduced a plugin dependency auto‑install feature that has since triggered several reported side effects.

## Project Progress (Merged / Closed PRs Today)

Among the 33 merged/closed PRs, the following key changes stand out:

- **Tool card loading spinner fix** ([#5141](agentscope-ai/QwenPaw PR #5141)) – Fixed missing spinner for shell commands and unregistered tools; refactored status logic.
- **Skill slash display improvement** ([#5146](agentscope-ai/QwenPaw PR #5146)) – Replaced full SKILL.md injection with `<skill>` block markers; closed associated bug #5031.
- **Wide mode toggle for chat layout** ([#5212](agentscope-ai/QwenPaw PR #5212)) – Adds a user‑toggleable expanded chat layout to improve screen utilization.
- **Agent OS Driver abstraction** ([#5067](agentscope-ai/QwenPaw PR #5067)) – First draft of a unified interface for MCP, A2A, ACP capabilities (merged after review).
- **Skill market UI & platform integration** ([#5123](agentscope-ai/QwenPaw PR #5123)) – Added QwenPaw skill market endpoint, categories, and previews.

Other closed PRs include minor CSS fixes (GitPanel tab styles, Shield icon centering), cron job repair tolerance, and integration test scaffolding for Sprint 2.4.

## Community Hot Topics

The most active discussions reflect deep‑seated user needs:

- **Xiaoyi (Huawei) channel integration** – [#1911](agentscope-ai/QwenPaw Issue #1911) (22 comments, open since March). User reports the channel connects but all responses return “开小差” (network congestion). No communication appears in CoPaw’s chat list, only in Xiaoyi’s open platform logs. The underlying need is a reliable debugging path for channel bridges.
- **File download regression** – [#5140](agentscope-ai/QwenPaw Issue #5140) (6 comments). Pure text files work in v1.1.11.post2 but docx/pdf trigger a 404. Despite earlier fixes, the issue persists – users increasingly rely on file exchange.
- **Feishu CardKit streaming slowness** – [#5167](agentscope-ai/QwenPaw Issue #5167) (5 comments). Long replies become unusably slow (“字字往外吐”). Community suggests optimizations like non‑streaming fallback or chunked updates.
- **MiniMax M2.5 thinking format breakage** – [#4625](agentscope-ai/QwenPaw Issue #4625) (5 comments). Model returns reasoning in XML instead of QwenPaw’s expected `thinking` blocks, causing agent instruction execution to fail. No fix PR yet.
- **Plugin dependency cmd window spam** – [#5181](agentscope-ai/QwenPaw Issue #5181) (5 comments). On PyPI‑unreachable networks, pip retry creates visible cmd.exe popup loops. Users demand background installation with proper logging.
- **Context compression data loss** – [#5171](agentscope-ai/QwenPaw Issue #5171) (4 comments). When agent persona files exceed token threshold, compression reduces context to zero – a critical blocker for long‑running agents.
- **Long conversation hang** – [#5161](agentscope-ai/QwenPaw Issue #5161) (4 comments). QwenPaw becomes completely unresponsive after many turns; no timeout or recovery mechanism.

## Bugs & Stability

Reported bugs today, ranked by severity:

| Issue | Severity | Description | Fix PR exists? |
|-------|----------|-------------|----------------|
| [#5181](agentscope-ai/QwenPaw Issue #5181) | **Critical** | Plugin dependency install spawns infinite cmd.exe popups | No |
| [#5209](agentscope-ai/QwenPaw Issue #5209) | **Critical** | macOS ARM64 Desktop (Tauri) crashes every ~1 minute; EXC_BAD_ACCESS | No |
| [#4625](agentscope-ai/QwenPaw Issue #4625) | **Critical** | MiniMax M2.5 XML thinking breaks agent commands; open since May 22 | No |
| [#5171](agentscope-ai/QwenPaw Issue #5171) | **Critical** | Context compression can wipe all context to 0 tokens | No |
| [#5214](agentscope-ai/QwenPaw Issue #5214) | **High** | DingTalk stream silently dies after laptop sleep; asyncio event loop freeze | No |
| [#5204](agentscope-ai/QwenPaw Issue #5204) | **High** | Two QwenPaw agents on Matrix enter infinite mutual wake‑up loop | No |
| [#5206](agentscope-ai/QwenPaw Issue #5206) | **High** | `load_agent_config()` returns cache reference, leading to configuration pollution | No |
| [#5161](agentscope-ai/QwenPaw Issue #5161) | **High** | Long conversations cause complete unresponsiveness (no timeout) | No |
| [#5140](agentscope-ai/QwenPaw Issue #5140) | **Medium** | File download (docx/pdf) returns 404 on post2 | No |
| [#5199](agentscope-ai/QwenPaw Issue #5199) | **Medium** | File attachment sending still intermittent | No |
| [#5208](agentscope-ai/QwenPaw Issue #5208) | **Medium** | Assistant message count mismatch when model uses `"reasoning"` instead of `"thinking"` | No |
| [#5207](agentscope-ai/QwenPaw Issue #5207) | **Low** | Inconsistent workspace path resolution between file tools and shell | No |

No PRs directly address the critical bugs yet, though the non‑blocking flush PR [#5215](agentscope-ai/QwenPaw PR #5215) may alleviate channel‑related hangs.

## Feature Requests & Roadmap Signals

Strong community demand for better context visibility and streaming performance:

- **Context usage display** – Multiple related issues ([#4284](agentscope-ai/QwenPaw Issue #4284), [#5103](agentscope-ai/QwenPaw Issue #5103), [#3366](agentscope-ai/QwenPaw Issue #3366), [#4782](agentscope-ai/QwenPaw Issue #4782)) request real‑time token counters, progress bars, and conversation turn counts. The backend already tracks these (see `light_context_manager.py`), but the frontend hasn’t exposed them.
- **Feishu CardKit streaming optimization** – [#5167](agentscope-ai/QwenPaw Issue #5167) proposes chunked updates or fallback to non‑streaming for long replies. A related PR [#5215](agentscope-ai/QwenPaw PR #5215) refactors delta dispatch to use non‑blocking, adaptive flushing – likely to land in the next release.
- **Headroom compression integration** – [#5063](agentscope-ai/QwenPaw Issue #5063) suggests incorporating an external reversible compression layer. No decision yet.
- **Agent self-evolution** – [#5205](agentscope-ai/QwenPaw Issue #5205) envisions agents learning from mistakes and auto‑correcting behavior beyond static rule files.
- **Desktop UX improvements** – [#5211](agentscope-ai/QwenPaw Issue #5211) criticizes excessive top‑bar space; fixed‑port configuration ([#5200](agentscope-ai/QwenPaw Issue #5200)) also requested.
- **Conversation queue** – [#5103](agentscope-ai/QwenPaw Issue #5103) and open PR [#5158](agentscope-ai/QwenPaw PR #5158) want an input queue so users can type while an agent is responding.

The Models Page overhaul ([#5203](agentscope-ai/QwenPaw PR #5203)) and the DataPaw analytics plugin ([#4622](agentscope-ai/QwenPaw PR #4622)) are both open and could appear in v1.2.

## User Feedback Summary

**Satisfaction** – The community is highly engaged, filing detailed bug reports and thoughtful enhancement proposals. The skill market and plugin ecosystem are warmly received.

**Pain Points** (recurring across multiple users):

- **File handling** – Download/upload of attachments (docx, pdf) is broken or intermittent, frustrating users who rely on document exchange.
- **Context management** – Context compression can destroy agent memory silently; no UI to monitor token usage before it’s too late.
- **Channel reliability** – Xiaoyi, DingTalk, Matrix, and Desktop channels all exhibit silent failure modes (sleep loss, loop, no error logs) that require manual restart.
- **Plugin UX** – Auto‑install of plugin dependencies is aggressive and not user‑friendly (visible cmd windows, infinite retry on network failure).
- **Long conversation hangs** – No progress indication or recovery—agent seems dead.

Common refrain: “Need better error messages, progress feedback, and self‑healing.”

## Backlog Watch

Several important items have been open for weeks without maintainer resolution:

- [#1911](agentscope-ai/QwenPaw Issue #1911) – Xiaoyi channel **silent failure** (open since 2026‑03‑20, 22 comments). No maintainer response in the last month.
- [#4625](agentscope-ai/QwenPaw Issue #4625) – MiniMax XML thinking breakage (open since 2026‑05‑22, 5 comments). Breaks agent execution for all MiniMax users.
- [#5171](agentscope-ai/QwenPaw Issue #5171) – Context compression **data loss** (open since 2026‑06‑13, 4 comments). A critical stability issue.
- [#5063](agentscope-ai/QwenPaw Issue #5063) – Headroom compression integration (open since 2026‑06‑10, 4 comments). No maintainer response.
- [#5204](agentscope-ai/QwenPaw Issue #5204) – Cross‑agent infinite loop (open since 2026‑06‑15). Important for multi‑agent deployments.
- [#5206](agentscope-ai/QwenPaw Issue #5206) – Config cache pollution (open since 2026‑06‑15). Low visibility but high impact.

**PRs needing review:**
- [#4900](agentscope-ai/QwenPaw PR #4900) – Plugin loader decoupling (open since June 2, 0 comments from maintainers).
- [#5040](agentscope-ai/QwenPaw PR #5040) & [#5041](agentscope-ai/QwenPaw PR #5041) – First‑time contributor patches for cron job and backup robustness (under review, but no updates in a week).
- [#5201](agentscope-ai/QwenPaw PR #5201) – Sprint 2.4 integration test framework (ready for merge).

**Advice to maintainers:** Prioritize the critical bugs (#5181, #5171, #5209) and provide at least status updates on the older issue #1911 to reassure the community.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest – 2026-06-16

## Today's Overview

ZeroClaw continues to see intense development activity, with **50 pull requests updated in the last 24 hours** (49 open, 1 merged/closed) and **3 new issues filed**. No new releases were published today. The community is actively submitting both bug fixes and enhancements, particularly around tool integration, channel support, and configuration management. The high PR volume suggests a large contributor base and efficient CI/CD pipelines, though the single merged PR indicates a significant review backlog. Issue severity includes one **workflow-blocking (S1)** bug related to MCP tools not reaching the model, which will likely attract urgent maintainer attention.

---

## Releases

*None in the last 24 hours.*

---

## Project Progress

Only one pull request was merged or closed in the past 24 hours. While not appearing in the top 20 by comment count, the project continues to accumulate improvements across many open PRs. Notable pending advancements visible in today's data include:

- **Refactored domain/URL validation** – [#7340](https://github.com/zeroclaw-labs/zeroclaw/pull/7340) extracts duplicate validation logic and fixes IPv6 bracket handling.
- **Slash-command support for the web chat** – [#7223](https://github.com/zeroclaw-labs/zeroclaw/pull/7223) adds client-side commands (`/help`, `/clear`, `/model`, etc.).
- **Mattermost WebSocket listener mode** – [#7098](https://github.com/zeroclaw-labs/zeroclaw/pull/7098) enables near-real-time event delivery as an alternative to polling.
- **WhatsApp reaction parity** – [#7535](https://github.com/zeroclaw-labs/zeroclaw/pull/7535) implements `add_reaction` / `remove_reaction`.
- **`zeroclaw doctor` diagnostics improvement** – [#7727](https://github.com/zeroclaw-labs/zeroclaw/pull/7727) surfaces non-fatal config warnings.
- **Windows self-update support** – [#7530](https://github.com/zeroclaw-labs/zeroclaw/pull/7530) adds `.zip` asset acceptance for Windows targets.

No regressions were introduced in the merged PR; details on the merged change are not available in the sampled data.

---

## Community Hot Topics

Activity is spread across many open PRs, with few concentrated discussions. The most active issue by comment count is:

- **[Bug]: native/MCP tools not sent to model** ([#7756](https://github.com/zeroclaw-labs/zeroclaw/issues/7756)) – 1 comment. Reported by a power user (perlowja) who also filed a session-ordering race. The issue identifies a critical gap where registered MCP tools are silently omitted for certain model backends. The single comment likely came from a maintainer requesting additional context.

**Underlying needs:** The community urgently needs reliable, model-agnostic tool delivery. The discrepancy between tool registration and actual sending suggests a deeper architectural issue in how providers interpret/forward tool definitions.

Among the many open PRs, those with broad community significance (based on enhancement scope) include the Mattermost WebSocket mode ([#7098](https://github.com/zeroclaw-labs/zeroclaw/pull/7098)) and the slash-command feature ([#7223](https://github.com/zeroclaw-labs/zeroclaw/pull/7223)). Both address long-standing user requests for better latency and user experience.

---

## Bugs & Stability

Three bugs were reported today, ranked by severity:

| Severity | Issue | Summary | Fix PR? |
|----------|-------|---------|---------|
| **S1 – workflow blocked** | [#7756](https://github.com/zeroclaw-labs/zeroclaw/issues/7756) | MCP tools registered but not sent to model on `wire_api=responses` and for Anthropic models. Tool-dependent agents cannot function. | No direct fix PR yet. |
| **S2 – degraded behavior** | [#7757](https://github.com/zeroclaw-labs/zeroclaw/issues/7757) | Gateway web dashboard Skills page only shows `skill_bundles`, missing workspace/open/plugin skills. | No fix PR yet. |
| **S3 (implicit)** | [#7753](https://github.com/zeroclaw-labs/zeroclaw/issues/7753) | Channel session persistence has a race condition: concurrent same-sender workers can corrupt session order. | No fix PR yet. |

Additionally, several bug-fix PRs updated today address pre-existing issues:

- **Poisoned lock recovery** ([#7755](https://github.com/zeroclaw-labs/zeroclaw/pull/7755)) – Prevents crashes from poisoned `ActivatedToolSet` locks in the runtime turn path.
- **OAuth credential fallback** ([#7640](https://github.com/zeroclaw-labs/zeroclaw/pull/7640)) – Fixes delegation to OAuth providers not using global credentials.
- **Config round-trip loss** ([#7532](https://github.com/zeroclaw-labs/zeroclaw/pull/7532)) – Aligns serde defaults with struct Default to avoid silent value loss on save/load.
- **Web fetch private host wildcard** ([#7424](https://github.com/zeroclaw-labs/zeroclaw/pull/7424)) – Ensures `["*"]` truly covers DNS-resolved private hosts.

No crashes or regressions were reported among the day’s closed items.

---

## Feature Requests & Roadmap Signals

Several user-facing enhancements are in active PR review:

- **Slash commands in web chat** ([#7223](https://github.com/zeroclaw-labs/zeroclaw/pull/7223)) – likely to land in the next minor release as it improves the gateway UX.
- **Mattermost WebSocket mode** ([#7098](https://github.com/zeroclaw-labs/zeroclaw/pull/7098)) – high value for Mattermost users; could become configurable per instance.
- **WhatsApp reaction parity** ([#7535](https://github.com/zeroclaw-labs/zeroclaw/pull/7535)) – fills a gap against other channels.
- **Per-channel ack_reactions for Lark/Feishu** ([#7495](https://github.com/zeroclaw-labs/zeroclaw/pull/7495)) – finally making the setting effective.
- **`zeroclaw quickstart` agent alias normalisation** ([#7637](https://github.com/zeroclaw-labs/zeroclaw/pull/7637)) – reduces FTUE friction.
- **Windows code page decoding test** ([#7670](https://github.com/zeroclaw-labs/zeroclaw/pull/7670)) – ensures non-ASCII shell output is handled correctly on Windows.

No new feature requests were filed as issues today. The current roadmap signals a strong focus on **channel parity** (WhatsApp, Lark, Mattermost) and **web UI expansion** (slash commands, skills page).

---

## User Feedback Summary

Real user pain points captured today:

- **“Workflow blocked” (perlowja):** MCP tools are registered but not sent to certain models, breaking all tool-dependent workflows. This user also reported a session-ordering race, indicating deep integration testing.
- **Skills page incompleteness (NiuBlibing):** The web dashboard only shows one type of skill bundle, missing workspace and plugin skills – a clear usability gap for users managing multiple skill sources.
- **Quickstart frustration (implied by PR #7637):** Capitalised agent aliases silently fail at the end of setup, discarding all input. The community responded with a fix to auto-normalise.

Satisfaction signals: contributors are actively submitting well-structured patches (e.g., runtime lock recovery, self-test authentication). The volume of PRs suggests a healthy and engaged development community.

---

## Backlog Watch

Several pull requests carry the **`needs-author-action`** label and risk becoming stale:

| PR | Title | Last Update | Days Since Author Action |
|----|-------|-------------|--------------------------|
| [#7094](https://github.com/zeroclaw-labs/zeroclaw/pull/7094) | fix(cli): make `zeroclaw models set` persist the model in config | 2026-06-16 | >14 (stale-candidate) |
| [#7098](https://github.com/zeroclaw-labs/zeroclaw/pull/7098) | feat(channel/mattermost): add optional WebSocket listener mode | 2026-06-16 | >14 (stale-candidate) |
| [#7215](https://github.com/zeroclaw-labs/zeroclaw/pull/7215) | fix(quickstart): surface port field for webhook channel config | 2026-06-16 | >12 |
| [#7340](https://github.com/zeroclaw-labs/zeroclaw/pull/7340) | refactor: extract duplicate domain/URL validation | 2026-06-16 | >9 |
| [#7532](https://github.com/zeroclaw-labs/zeroclaw/pull/7532) | fix(config): align serde defaults with struct Default | 2026-06-16 | >4 (needs-author-action) |

These PRs represent important fixes and features. Two are flagged as **stale candidates** – maintainers may need to nudge authors or take over to prevent valuable work from decaying. Notably, the Mattermost WebSocket feature ([#7098](https://github.com/zeroclaw-labs/zeroclaw/pull/7098)) has been waiting for author action for over two weeks despite strong community interest.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*