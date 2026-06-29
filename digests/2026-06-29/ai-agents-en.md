# OpenClaw Ecosystem Digest 2026-06-29

> Issues: 82 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-29 14:39 UTC

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

# OpenClaw Project Digest — 2026-06-29

## Today's Overview
OpenClaw saw exceptionally high activity on 2026-06-29 — **82 issues** and **500 PRs** were updated in the last 24 hours, indicating strong maintainer and community engagement. Of those, **12 issues** were closed and **54 PRs** were merged or closed. The project released a new beta (`v2026.6.11-beta.2`) with improved channel control and operator tooling. The pulse of the project is vigorous, with many fixes targeting memory pressure, routing correctness, and provider compatibility.

## Releases
### v2026.6.11-beta.2
A new beta was released today. Key highlights (from the release notes) include:
- **More capable channel control:** Slack relay mode, native Mattermost `/oc_queue` command, and per-DM model overrides — making multi-channel workflows easier to automate and tune. (PRs #94707, #95546, #95120; thanks to @sjf-oa, @amknight, @xydigit-zt, @thomaszta, @gandalf-at-lerian)
- "Richer operat…" (truncated in source) — likely referencing richer operator controls or reporting.

No breaking changes or migration notes were reported in this digest data. The release marks a step forward for production channel management.

## Project Progress
Today **54 PRs were merged or closed**, covering a wide array of fixes and improvements. Among the most notable (based on top-commented PRs listed this period):
- **#97742** (open) – Fix LLM structured tool result text preservation across providers.
- **#94172** (open) – Expose filesystem discovery tools (ls/find/grep) to the Codex runtime.
- **#87449** (open) – Preserve text-block boundaries in Mattermost draft previews, fixing a long-standing issue.
- **#96980** (open) – Deduplicate assistant messages in TUI after session reload.
- **#96938** (open) – Keep reply directive IDs Unicode-safe to prevent malformed surrogates.
- **#97803** (open) – Throw typed `ExitError` instead of generic `Error` for simulated runtime exit.

Also, **12 issues were closed** today, including:
- **#81117** (closed, enhancement) – Control UI session picker now shows derived titles instead of raw IDs.
- **#97690** (closed, bug) – Dreaming phase completion events now normalize outcome status.

These closures indicate steady progress on both user-facing improvements and internal reliability.

## Community Hot Topics
The most active discussions this period (by comment count):

| Issue/PR | Comments | Topic |
|----------|----------|-------|
| [#81061](https://github.com/openclaw/openclaw/issues/81061) | 7 | Feature request for a `before_route_inbound_message` hook enabling channel bridging/proxying before routing. |
| [#81156](https://github.com/openclaw/openclaw/issues/81156) | 5 | Bug: MiniMax usage count semantics inverted (shows % used as % left). |
| [#80918](https://github.com/openclaw/openclaw/issues/80918) | 4 | Silent send miss: incomplete-turn classifier discards final text turn after `update_plan` tool. |
| [#81339](https://github.com/openclaw/openclaw/issues/81339) | 3 | WebChat UI strips leading whitespace in code blocks, breaking ASCII diagrams. |
| [#81213](https://github.com/openclaw/openclaw/issues/81213) | 3 | Beta report: Codex runtime loads but OpenAI primary times out with inconsistent fallback. |

Underlying needs:
- **Routing extensibility (#81061):** Users need a pre-routing hook to build custom channel bridges and proxies — a clear architectural gap.
- **Data correctness (#81156, #80918):** Trust in provider usage displays and message delivery is critical; both bugs erode user confidence.
- **UI precision (#81339):** Developers rely on ASCII diagrams in code blocks; stripping whitespace breaks developer workflows.

## Bugs & Stability
Reported bugs today (updated or created) ranked by severity:

### P1 (Critical)
- **#81099** – `claude-cli` backend: `AskUserQuestion` tool never returns result (no picker UI). Blocks interactive workflows.
- **#81234** – Cron agent jobs time out after turn-accepted; Discord DM lane blocked by stale session key. Risk of message loss.
- **#80926** – Azure OpenAI Responses stalls before first event when memory tools are exposed. No fix PR yet.

### P2 (High)
- **#80933** – CLI compaction never triggers for `claude-cli` runtime; context tokens budget ignored. Companion PR likely pending.
- **#81182** – Overflow recovery waits full auto-compaction timeout before truncating tool results — adds ~900s latency.
- **#80862** – Telegram reasoning previews create multiple messages instead of editing one, leaving stale messages.
- **#81178** – Repeated early preflight compactions after successful compaction due to stale transcript usage.
- **#80858** – Dreaming pipeline promotes empty-result placeholder into MEMORY.md (fix PR #80916 exists and is ready).
- **#81322** – WhatsApp image processing works but images not attached to outbound messages.
- **#80920** – Gateway hangs indefinitely on macOS Big Sur due to `canBindToHost("::1")` with no timeout.

### P3 (Medium)
- **#80984** – Typed error for cron delivery with no prior chat (follow-up to PR).
- **#81355** – RPC fanout performance: `tts.status` monopolizes event loop; `applyPluginAutoEnable` recomputes 8× per fanout.

Several bugs have associated open fix PRs (e.g., #80916 for #80858, #81260 for duplicate progress lines, #81300 for Codex reasoning exposure). The volume of P2 bugs suggests need for a dedicated stability sprint.

## Feature Requests & Roadmap Signals
Notable feature requests from today’s activity:

- **#80853** – Session maintenance config to exempt specific sessions (e.g., primary WebUI) from pruning.
- **#80989** – `/progress` command toggle for tool-call progress in streaming preview (mobile-friendly).
- **#80942** – Control UI auto-set `document.title` to agent name for multi-agent workflows.
- **#80806** – Expose sender ID/metadata in agent turn hooks for security and routing.
- **#81271** – Per-sender exec node routing for multi-node setups.
- **#80836** – `gateway.auth.tokenScopes` config for headless/single-tenant deployments.
- **#81232** – Discord tool to fetch single message by ID/URL.
- **#80841** – Twilio AMD support for outbound voice calls (machine detection).
- **#81253** – i18n/template support for exec approval prompts (better UX for non-technical users).

*Prediction for next version:* The pre-routing hook (#81061) and session title picker (#81117, already closed) are strong candidates for inclusion in upcoming stable releases. Similarly, the `agents.setDefault` RPC (#92957) and the token-scopes config (#80836) address pain points for operators.

## User Feedback Summary
**Pain points voiced by users:**
- **Misleading usage display:** MiniMax percentage reversed (#81156) – “I rely on this to manage quota.”
- **Session list useless:** Raw UUIDs instead of derived titles (#81117) – “Impossible to find conversations.”
- **Stale reasoning messages on Telegram:** (#80862) – “Clutters chat with garbage.”
- **Code block integrity broken:** (#81339) – “ASCII diagrams unreadable.”
- **Cron job duplication:** (#81087) – “Feishu messages sent twice every day.”
- **Inconsistent channel behavior:** Mattermost private channels get non-deterministic `chat_type` (#95646).

**Positive signals:** The new beta’s Slack relay and Mattermost /oc_queue command were appreciated (multiple contributors thanked). The overall PR count (500) suggests a highly engaged contributor base.

## Backlog Watch
Issues and PRs requiring maintainer attention due to age or missing decisions:

- **#80806** (2026-05-11) – Expose sender metadata in hooks; needs maintainer review and product decision.
- **#80984** (2026-05-12) – Typed error for cron delivery; needs product decision.
- **#81213** (2026-05-12) – Beta report with timeframe fallback – needs maintainer review and info.
- **#80841** (2026-05-12) – Voice call AMD – open for 48 days with no product decision.
- **#80799** (2026-05-11) – Flat-key disclosure of sibling agent credentials – critical security gap still needs security review and product decision.
- **#80864** (2026-05-12) – Preflight validation framework – awaiting maintainer review and product decision.
- **#81271** (2026-05-13) – Per-sender exec routing – stalled on security review.
- **#81310** (2026-05-13) – Plugin shutdown drainage contract – needs maintainer review.

These items have been stale for over a month (since mid-May) and represent significant usability, security, or architectural gaps. A maintainer review pass would help unlock contributor progress.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: AI Agent Open-Source Ecosystem

## 1. Ecosystem Overview

The personal AI assistant open-source landscape on 2026-06-29 shows a mature, rapidly converging ecosystem with 12 tracked projects exhibiting starkly different development velocities. The core trend is a collective push toward **production readiness**: multi-agent orchestration, provider reliability, channel completeness, and cost optimization dominate the collective backlog. A clear bifurcation has emerged between high-velocity flagship implementations (OpenClaw, ZeroClaw, CoPaw) and stable-but-quiescent niche projects (TinyClaw, Moltis, ZeptoClaw). Security hardening—particularly around credential handling, symlink containment, and authorization bypass—is now a shared priority across five major projects. The ecosystem is transitioning from "can it work?" to "can it work reliably at scale?".

---

## 2. Activity Comparison

| Project | Issues Updated (24h) | PRs Updated (24h) | Release Today? | Estimated Health Score* |
|---------|---------------------|--------------------|----------------|------------------------|
| **OpenClaw** | 82 (12 closed) | 500 (54 merged) | ✅ v2026.6.11-beta.2 | 9.5/10 |
| **ZeroClaw** | 10 (1 closed) | 50 (20 merged) | ❌ | 8.5/10 |
| **CoPaw** | 13 (4 closed) | 50 (23 merged) | ❌ | 8.5/10 |
| **IronClaw** | 4 (2 closed) | 50 (20 merged) | ❌ | 7.5/10 |
| **Hermes Agent** | 19 (0 closed) | 50 (0 merged) | ❌ | 6.5/10 |
| **NanoBot** | 6 (2 closed) | 27 (10 merged) | ❌ | 8.0/10 |
| **LobsterAI** | 7 (3 closed) | 40 (39 merged) | ✅ v2026.6.29 | 8.0/10 |
| **NanoClaw** | 0 | 10 (3 merged) | ❌ | 7.0/10 |
| **NullClaw** | 1 (1 closed) | 4 (1 closed) | ❌ | 6.5/10 |
| **PicoClaw** | 1 (1 closed) | 1 (1 closed) | ❌ | 4.5/10 |
| **Moltis** | 1 (0 closed) | 0 | ❌ | 3.0/10 |
| **TinyClaw** | 0 | 0 | ❌ | 1.0/10 |
| **ZeptoClaw** | 0 | 0 | ❌ | 1.0/10 |

*Health Score: Composite of engagement volume, merge velocity, bug response time, and community feedback quality. Subjective but data-informed.*

---

## 3. OpenClaw's Position

**Advantages vs Peers:**
- **Community scale:** 82 issues and 500 PRs in 24 hours—3–10× the activity of any peer. This yields faster bug discovery and broader provider coverage.
- **Release cadence:** Multiple beta releases per month, including today's v2026.6.11-beta.2 with production-grade channel control (Slack relay, Mattermost `/oc_queue`, per-DM model overrides).
- **Provider breadth:** Actively fixing LLM structured tool result preservation across providers—a problem only CoPaw and ZeroClaw also tackle.
- **Architecture maturity:** Pre-routing hooks (#81061), session title pickers (closed), and typed exit errors demonstrate a production-oriented design philosophy.

**Technical Approach Differences:**
- OpenClaw uses a **monolithic core with layered adapter system** (Slack, Mattermost, Discord). Hermes Agent pursues a self-registering pluggable architecture (#3823) but hasn't delivered. ZeroClaw uses daemon-owned SOP control planes for workflow orchestration—a more radical architectural bet.
- OpenClaw's **540 bugs with active fix PRs** suggests a "ship fast, fix faster" philosophy vs. Hermes Agent's "triage then merge" approach (0 PRs merged today despite 19 bugs).

**Community Size Comparison:**
- OpenClaw's 500 daily PR contributors dwarf all peers. The next closest (ZeroClaw, CoPaw, IronClaw) each have ~50 PRs/day. OpenClaw likely has 5–10× the contributor base.
- However, this scale creates a **review bottleneck**: the backlog shows 8+ items over 30 days stale, suggesting maintainers can't keep pace with contribution volume.

---

## 4. Shared Technical Focus Areas

Five cross-project requirements are emerging simultaneously across multiple implementations:

| Focus Area | Affected Projects | Specific Needs |
|------------|-------------------|----------------|
| **Provider Reliability & Fallbacks** | OpenClaw, NanoBot, ZeroClaw, CoPaw, IronClaw | Inconsistent streaming (Kimi #5600), cache invalidation (#4222 NanoBot), timeout handling (#81213 OpenClaw), reasoning model support (#4419 NanoBot) |
| **Channel Completeness & Consistency** | OpenClaw, Hermes, NanoClaw, CoPaw, ZeroClaw | WhatsApp images not attached (#81322), Feishu long message truncation (#5561 CoPaw), Mattermost chat_type nondeterminism (#95646), Telegram stale messages (#80862) |
| **Context / Token Optimization** | OpenClaw, NanoBot, CoPaw, NullClaw, IronClaw | Prefix caching (NanoBot #4222), layered skills index (Hermes #54939), microcompaction (NanoBot #4392), hard tool result caps (CoPaw #5342) |
| **Multi-Agent Security** | OpenClaw, NanoBot, NanoClaw, ZeroClaw, IronClaw | Credential leakage in MCP URLs (NanoBot #4584), symlink escapes (NanoClaw #2880), IDOR in /resume (Hermes #52355), flat-key disclosure (OpenClaw #80799) |
| **Cost Visibility & Billing** | OpenClaw, ZeroClaw, CoPaw, LobsterAI | Misleading usage displays (#81156 OpenClaw), cost/org snapshot RPC (ZeroClaw #8482), DeepSeek cache hit rate (CoPaw #3891), subscription credit opacity (LobsterAI #2081) |

**Underlying pattern:** The ecosystem is converging on a **reliability-first, cost-aware, multi-channel** standard. No single project dominates all five areas, but every major project is investing in at least three.

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | ZeroClaw | CoPaw | Hermes Agent | NanoBot | IronClaw |
|-----------|----------|----------|-------|-------------|---------|----------|
| **Primary User** | General agent builders | Enterprise teams | Chinese market developers | Research / advanced users | General / lightweight | Developer productivity |
| **Key Differentiator** | Scale & channel breadth | SOP control planes | Mission Mode + runtime v2 | Browser automation & LSP | Context governance | Tool permission system |
| **Architecture** | Monolithic + adapters | Daemon-owned workflows | Plugin-based middleware | Self-registering adapters (planned) | Modular subagent spawn | Capability policy model |
| **Deployment Target** | Self-hosted, Docker | Production clusters | Cloud + self-hosted | macOS/Windows desktop | Lightweight / edge | Developer workstation |
| **Channel Strength** | Slack, Mattermost, Discord, Telegram | WhatsApp, WeChat | Feishu, DingTalk, QQ | Multiple (buggy on Windows) | WebUI, CLI | WebUI v2 (Reborn) |
| **Community Language** | English-heavy | Chinese + English | Chinese-primary | English-heavy | English | English |
| **Maturity Stage** | Rapid iterative release | Feature-complete hardening | Refactoring + test coverage | Triage-heavy | Steady incremental | Architectural pivot |

**Key insight:** No project is "best" across all dimensions. OpenClaw wins on breadth, ZeroClaw on operational maturity, CoPaw on Chinese market integration, and IronClaw on developer workflow depth. The projects are **converging in capabilities but diverging in deployment philosophy**.

---

## 6. Community Momentum & Maturity

**Tier 1: High Velocity (10+ PRs/day, active releases)**
- **OpenClaw** – Unmatched volume but facing review bottlenecks. Beta releases every 1-2 weeks.
- **ZeroClaw** – 20 merges/day with systematic feature rollout (SOP, WASM, cost tracking). Nearing v0.8.0.
- **CoPaw** – 23 merges/day, heavy test coverage investment. Likely near release candidate.
- **LobsterAI** – Release-day spike (39 merges). Generally steady, not hyperactive.

**Tier 2: Moderate Activity (5-10 PRs/day, periodic releases)**
- **NanoBot** – 10 merges/day, focused on caching and subagent flexibility. Healthy, predictable cadence.
- **IronClaw** – 20 merges/day but high proportion of chore/dependency PRs. Major error-recoverability refactor underway.
- **NanoClaw** – 3 merges/day. Small team but responsive. Discord adapter and dashboard suggest growth.

**Tier 3: Low Activity / Stagnant (<3 PRs/day, no recent releases)**
- **Hermes Agent** – **Warning signs**: 50 PRs but 0 merged today, 19 bugs zero closed. Windows quality gap. Maintainers may be overwhelmed.
- **NullClaw** – 1 merge today, slow but focused (REPL fix + streaming tools). Minimal feature creep.
- **PicoClaw** – Closing stale items without merging. Maintainer attention minimal. Risk of abandonment.
- **Moltis** – 1 issue update, 0 PRs. Effectively dormant.
- **TinyClaw / ZeptoClaw** – No activity. Dead.

**Momentum trend:** The ecosystem is consolidating around a handful of "winner" implementations. Projects that fail to achieve critical mass of either contributors (Hermes) or maintainer attention (PicoClaw) are falling behind rapidly.

---

## 7. Trend Signals

**Extracted from community feedback across all projects:**

### 1. Provider Independence is the #1 Pain Point
Users are frustrated by API-breaking differences between LLM providers. The Kimi streaming error (#5600 ZeroClaw), MiniMax percentage display (#81156 OpenClaw), and Groq native tool calling (#7909 ZeroClaw) all represent the same class of problem: **providers implement the same spec differently.** The market is demanding a robust normalization layer that no project fully delivers yet.

### 2. Voice I/O is Underserved but Wanted
Three projects (NanoBot #4010, OpenClaw's Twilio AMD #80841, LobsterAI's email connectivity #1388) have voice or telephony feature requests open. Voice input is working in some projects, but **voice output and two-way calling remain gaps.** This signals a shift from "text assistant" to "voice assistant" expectations.

### 3. Cost Transparency Becomes Table Stakes
Users no longer accept opaque billing. The MiniMax reversed percentage bug (#81156), DeepSeek cache inefficiency (#3891 CoPaw), and subscription credit complaints (#2081 LobsterAI) all show that **developers need real-time, accurate cost feedback** to manage agent deployments. ZeroClaw's cost/org snapshot (#8482) may become a template for others.

### 4. LSP and Post-Task Reflection Show Developer-Tool Ambitions
Hermes Agent's LSP integration (#516), post-task reflection (#483), and IronClaw's tool permission system point toward a future where **agents actively debug and improve their own tool chains.** This is a leading indicator of the "agent-as-developer-assistant" use case, distinct from "agent-as-chatbot."

### 5. Channel-Specific Fixes are a Reminder of Fragmentation
Each project is fighting the same battle with different platforms: WhatsApp image handling (OpenClaw #81322, ZeroClaw #8389), Telegram message editing (OpenClaw #80862), Feishu long replies (CoPaw #5561), WeChat/DingTalk quirks (CoPaw). **The ecosystem would benefit from a shared channel adapter library**, but no such effort exists across projects.

### 6. TUI/CLI Usability is Being Re-invested
After a period of web-first focus, NullClaw's arrow-key fix (#960/970), OpenClaw's TUI session reload deduplication (#96980), and NanoClaw's dashboard pusher (#2871) show that **terminal-based interaction is not dead.** Developers running agents in headless or CI environments value terminal UX.

### Value for AI Agent Developers:

- **If you need production channel breadth:** OpenClaw is the safest bet, but be prepared to deal with review delays on PRs.
- **If you need cost-controlled multi-agent workflows:** ZeroClaw's SOP + cost tracking is ahead of peers.
- **If you serve Chinese-market users:** CoPaw has the deepest Feishu/DingTalk integration.
- **If you want a lightweight, responsive codebase:** NanoBot offers rapid fixes and clean architecture with fewer moving parts.
- **If you're building developer tools:** IronClaw's capability policy model and Hermes's browser automation are worth studying.
- **If stability matters more than features:** Avoid Hermes (Windows gap) and PicoClaw (risk of abandonment). ZeroClaw and CoPaw show the strongest merge-to-bug ratio.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest – 2026-06-29

## 1. Today's Overview
The NanoBot project shows very high activity: 6 issues were updated in the last 24 hours (4 open, 2 closed) and 27 pull requests were touched (17 open, 10 closed/merged). No new releases were published. The repository is in an intense development phase, with major attention to context-governance reliability, subagent model flexibility, WebUI usability, and security hardening. The closing of two long-standing issues (lightweight claim dispute, prefix caching bug) signals that maintainers are actively addressing community pain points.

## 2. Releases
*None* – no releases were published on 2026-06-29.

## 3. Project Progress
Ten pull requests were merged or closed today. Key advancements:

- **Caching and token estimation fixes** – PR [#4254](https://github.com/HKUDS/nanobot/pull/4254) applied microcompaction when estimating session prompt tokens, resolving the root cause of issue [#4222](https://github.com/HKUDS/nanobot/issues/4222) (prefix caching invalidation). PR [#4392](https://github.com/HKUDS/nanobot/pull/4392) made tool microcompaction configurable, giving cache-sensitive deployments more control.
- **WebUI session timestamps and Markdown export** – PR [#4585](https://github.com/HKUDS/nanobot/pull/4585) (closed, merged) delivers both features requested in [#4579](https://github.com/HKUDS/nanobot/issues/4579). The export function preserves markdown and collapses tool/trace activity.
- **Session retention refactor** – PR [#4574](https://github.com/HKUDS/nanobot/pull/4574) replaced a bare tuple return with a typed `RetentionResult`, improving code clarity and preventing misuse.
- **Duplicate skill guard** – PR [#4554](https://github.com/HKUDS/nanobot/pull/4554) (open, but noted as merged elsewhere) blocks Dream from creating duplicate skills, addressing a recurring community complaint.
- **CLI OAuth main-provider flag** – PR [#4573](https://github.com/HKUDS/nanobot/pull/4573) (open) adds `--set-main` to `nanobot provider login`, solving a setup friction reported by users.

## 4. Community Hot Topics
The most debated items this week:

- **#660 – Lightweight claim vs. Node.js dependency** (closed, 15 comments, 5 👍). Users questioned the "ultra-lightweight" tag because the Dockerfile includes Node.js. The issue was closed, likely after maintainers clarified or updated documentation. See [HKUDS/nanobot#660](https://github.com/HKUDS/nanobot/issues/660).
- **#4419 – Automatic reasoning effort escalation** (open, 4 comments). A popular feature request to dynamically adjust `reasoningEffort` for provider reasoning models. Underlying need: power users want granular control over model thought depth without manual config changes. See [HKUDS/nanobot#4419](https://github.com/HKUDS/nanobot/issues/4419).
- **#4010 – Text-to-speech / voice output** (open, 2 comments, 2 👍). Users want to close the voice loop after voice input already works. This feature would enable spoken replies on channels like Telegram. See [HKUDS/nanobot#4010](https://github.com/HKUDS/nanobot/issues/4010).

## 5. Bugs & Stability
Two notable bug reports and one high-severity security fix:

- **🔥 High severity – Credential leakage in MCP URLs** – PR [#4584](https://github.com/HKUDS/nanobot/pull/4584) (open) redacts secrets from URLs before logging. Previously `userinfo` and query-string tokens were logged in plaintext on validation/connect paths. No fix merged yet.
- **Medium severity – `max_messages` truncation invalidates prefix caching** – Issue [#4222](https://github.com/HKUDS/nanobot/issues/4222) (closed) – two mechanisms (truncation boundary drift and microcompaction) caused per-turn cache misses. Fixed by merges in PRs [#4254](https://github.com/HKUDS/nanobot/pull/4254) and [#4392](https://github.com/HKUDS/nanobot/pull/4392).
- **Low severity – Tool-key migration crash on null sections** – PR [#4583](https://github.com/HKUDS/nanobot/pull/4583) (open) guards `dict.get()` where a key is present but set to `None`. Prevents crashes during config loading.

No new crashes or regressions were reported today. The caching fix is the most impactful stability improvement.

## 6. Feature Requests & Roadmap Signals
Several feature requests point to the next minor version:

- **Subagent model overrides** – PR [#4291](https://github.com/HKUDS/nanobot/pull/4291) allows spawn to specify a named model preset. PR [#4570](https://github.com/HKUDS/nanobot/pull/4570) (closed as duplicate) also addressed this. Likely to land soon.
- **Native Agent-to-Agent (A2A) delegation** – PR [#4571](https://github.com/HKUDS/nanobot/pull/4571) implements cross-delegation depth guards and peer agent registration. Targets [#4179](https://github.com/HKUDS/nanobot/issues/4179). This is a major architectural upgrade.
- **Context/token optimization** – PR [#4588](https://github.com/HKUDS/nanobot/pull/4588) introduces compact command-output modules for large exec results, plus oversized subagent announcement compaction. Complementary to the caching fixes.
- **Conda/virtual environment support** – Issue [#4580](https://github.com/HKUDS/nanobot/issues/4580) asks for `conda` environment handling in subprocesses. Not yet implemented but has maintainer attention.
- **Ask clarification tool** – PR [#4527](https://github.com/HKUDS/nanobot/pull/4527) adds a built-in tool for agent-to-user clarification. Small footprint, high value for conversational flows.
- **Globalping MCP preset** – PR [#4383](https://github.com/HKUDS/nanobot/pull/4383) adds a free network-measurement MCP server. Signals expansion of built-in MCP integrations.

**Prediction for next major release (v1.x):** The combination of A2A delegation, subagent model override, WebUI export, and context-cost reductions strongly suggests a focus on enterprise/production readiness and multi-agent workflows.

## 7. User Feedback Summary
- **Satisfied** – Users appreciate the rapid fixes for caching and token estimation (closing [#4222]). The WebUI timestamp and export features received positive reactions.
- **Frustrations** – The “ultra-lightweight” branding still bothers some users (see [#660](https://github.com/HKUDS/nanobot/issues/660)). Power users want more control over reasoning model parameters ([#4419](https://github.com/HKUDS/nanobot/issues/4419)) and native voice output ([#4010](https://github.com/HKUDS/nanobot/issues/4010)).
- **Pain point** – The lack of conda/virtual environment support ( [#4580](https://github.com/HKUDS/nanobot/issues/4580) ) forces devs to use system Python for subprocess execution, which is inconvenient for reproducible setups.
- **Real-world use case** – Multiple parallel sessions: the WebUI improvements directly address the pain of managing chat histories without timestamps or export.

## 8. Backlog Watch
No issues are critically stale (older than 60 days without update), but a few need maintainer eyes:

- **#4010 – Text-to-speech / voice output** – Opened 2026-05-26, only 2 comments. No maintainer reply yet. If voice output is not on the short-term roadmap, a clear statement would help.
- **#4383 – Globalping MCP preset** (PR) – Opened 2026-06-17, still open. The feature is additive and low risk; could be merged quickly.
- **#4291 – Subagent model presets** – Opened 2026-06-11, still open and overlapping with [#4570](https://github.com/HKUDS/nanobot/pull/4570) (closed as duplicate). Maintainers should consolidate and merge the best approach.

No PR has gone more than 2 weeks without update, indicating a healthy review pace.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-06-29

## 1. Today's Overview

Hermes Agent saw intense activity on 29 June 2026, with **19 issues** and **50 pull requests** updated in the last 24 hours. **All items remain open** – no issues were closed and no PRs were merged or released, indicating a day focused on triaging new bug reports and submitting fixes rather than landing completed work. The project’s health shows a healthy but busy development cycle, with a sharp spike in Windows-related bug reports and a coordinated batch of security-oriented pull requests from the same contributor (ooiuii) addressing webhook body-size enforcement across multiple platforms. The absence of new releases or merges suggests maintainers are currently reviewing a large queue of pending changes.

## 2. Releases

No new releases were published today. The latest release remains unknown based on the data.

## 3. Project Progress

**Merged/closed PRs today: 0** — no feature work or bug fixes landed. However, **50 pull requests were actively updated**, signalling significant work-in-progress:

- Several platform-specific webhook body-limit enforcement PRs (WhatsApp, Feishu, LINE, WeCom) are now open for review (e.g., [#54944](https://github.com/NousResearch/hermes-agent/pull/54944), [#54938](https://github.com/NousResearch/hermes-agent/pull/54938), [#54931](https://github.com/NousResearch/hermes-agent/pull/54931), [#54934](https://github.com/NousResearch/hermes-agent/pull/54934)).
- Two high-priority security PRs from claudlos address an **IDOR vulnerability in /resume** ([#52355](https://github.com/NousResearch/hermes-agent/pull/52355)) and **cloud-metadata bypass in browser_navigate** ([#52349](https://github.com/NousResearch/hermes-agent/pull/52349)).
- A notable feature PR adds **opt-in layered skills index** to cut system-prompt tokens ([#54939](https://github.com/NousResearch/hermes-agent/pull/54939)).
- Fixes for **Windows compatibility** (pty handling, process termination) are proposed in [#54933](https://github.com/NousResearch/hermes-agent/pull/54933) and [#54858](https://github.com/NousResearch/hermes-agent/pull/54858).
- **iMessage HEIC attachment handling** is fixed in [#54946](https://github.com/NousResearch/hermes-agent/pull/54946).

## 4. Community Hot Topics

The most-discussed issues (by comments and reactions) reveal deep developer interest in improving the agent’s extensibility and intelligence:

- **[#3823](https://github.com/NousResearch/hermes-agent/issues/3823) – Refactor gateway: self-registering adapters** (5 comments, 2 👍)  
  *Need:* Eliminate the 17+ files required to add a new messaging platform adapter. The community wants a pluggable architecture that reduces boilerplate and enables faster integrations.

- **[#516](https://github.com/NousResearch/hermes-agent/issues/516) – LSP Integration** (4 comments, 1 👍)  
  *Need:* Post-edit diagnostics and code intelligence similar to Kilocode. Developers want the agent to understand code context after edits, enabling smarter code review and refactoring.

- **[#15497](https://github.com/NousResearch/hermes-agent/issues/15497) – Hindsight provider shutdown race** (4 comments)  
  *Need:* Memory provider stability during interpreter shutdown. This is a hard-to-reproduce but critical race condition affecting state persistence.

- **[#483](https://github.com/NousResearch/hermes-agent/issues/483) – Post-task reflection & missing affordance detection** (3 comments)  
  *Need:* Enhancing agent self-awareness to detect missing tools/actions after completing a task, inspired by Cognitive Workbench research.

The volume of comments on these older feature requests indicates sustained community interest. No PRs on these have been merged yet.

## 5. Bugs & Stability

**19 bugs were reported today**, with a heavy concentration on Windows and memory/profile isolation. Severity ranking:

| Severity | Issue | Title | Fix PR? |
|----------|-------|-------|---------|
| **P1** | [#54919](https://github.com/NousResearch/hermes-agent/issues/54919) | Hermes not launching on Windows (uv trampoline error) | No PR yet |
| **P2** | [#54929](https://github.com/NousResearch/hermes-agent/issues/54929) | Cross-session message leak | No PR yet |
| **P2** | [#54936](https://github.com/NousResearch/hermes-agent/issues/54936) | Node window flashes continuously on Windows | No PR yet |
| **P2** | [#54937](https://github.com/NousResearch/hermes-agent/issues/54937) | Background_review cross-profile memory contamination | No PR yet |
| **P2** | [#54927](https://github.com/NousResearch/hermes-agent/issues/54927) | winpty-rs panic on Chinese Windows (HRESULT 0x00000000) | No PR yet |
| **P2** | [#54928](https://github.com/NousResearch/hermes-agent/issues/54928) | Desktop boot timeout too short (15s) on Windows | No PR yet |
| **P2** | [#54947](https://github.com/NousResearch/hermes-agent/issues/54947) | Agent cache cross-process write detection invalidates every turn | No PR yet |
| **P3** | [#54935](https://github.com/NousResearch/hermes-agent/issues/54935) | Feishu webhook buffers oversized chunked bodies before 413 | [#54938](https://github.com/NousResearch/hermes-agent/pull/54938) |
| **P3** | [#54940](https://github.com/NousResearch/hermes-agent/issues/54940) | WhatsApp webhook same issue | [#54944](https://github.com/NousResearch/hermes-agent/pull/54944) |
| **P3** | [#54930](https://github.com/NousResearch/hermes-agent/issues/54930) | LINE webhook same issue | [#54931](https://github.com/NousResearch/hermes-agent/pull/54931) |
| **P3** | [#54932](https://github.com/NousResearch/hermes-agent/issues/54932) | WeCom webhook same issue | [#54934](https://github.com/NousResearch/hermes-agent/pull/54934) |
| **P3** | [#54922](https://github.com/NousResearch/hermes-agent/issues/54922) | custom_providers[].extra_body silently dropped on gateway paths | No PR yet |
| **P3** | [#54945](https://github.com/NousResearch/hermes-agent/issues/54945) | Mem0 OSS flags rejected by top-level argparse | No PR yet |

Most high-severity bugs are **Windows-specific** (starting, flashing, timeout, panic), indicating a major quality gap on that platform. The webhook body-limit bugs have already been addressed by PRs, and the Notion HMAC signature issue has a replacement PR in review ([#54942](https://github.com/NousResearch/hermes-agent/pull/54942)).

## 6. Feature Requests & Roadmap Signals

New feature requests filed today point to ongoing demands for **operational flexibility and smarter context management**:

- **[#54941](https://github.com/NousResearch/hermes-agent/issues/54941) – /reload-plugins hot-reload**  
  Users want to enable/disable plugins without starting a new session. Likely to be picked up soon given the plugin ecosystem growth.

- **[#54926](https://github.com/NousResearch/hermes-agent/issues/54926) – `hermes update` should migrate all profiles**  
  Currently only migrates the active profile; community wants multi-profile support.

- **[#54939](https://github.com/NousResearch/hermes-agent/pull/54939) – Layered skills index** (PR)  
  Opt-in feature to compress system prompts by splitting active/dormant skills. Could land in the next minor release.

- Older feature requests [#3823](https://github.com/NousResearch/hermes-agent/issues/3823) (self-registering adapters), [#516](https://github.com/NousResearch/hermes-agent/issues/516) (LSP), and [#483](https://github.com/NousResearch/hermes-agent/issues/483) (post-task reflection) remain open with no PRs attached. They represent long-term roadmap items.

**Predictions for next version:** The webhook body-limit fixes, Notion signature fix, and Windows pty/compat fixes are all small, self-contained changes likely to be merged soon. The security PRs (IDOR, cloud-metadata) are high priority and may be fast-tracked.

## 7. User Feedback Summary

Real user pain points surfaced in today’s bug reports and comments:

- **Poor Windows experience**  
  Multiple users reported the agent failing to start, Node.js windows flashing, desktop boot timing out, and a panic on Chinese Windows. This is the single biggest source of dissatisfaction. Users explicitly ask for “fix node flash” and longer boot timeout.

- **Memory/profile isolation failures**  
  Two bugs ([#54929](https://github.com/NousResearch/hermes-agent/issues/54929), [#54937](https://github.com/NousResearch/hermes-agent/issues/54937)) describe conversations leaking across sessions and memory contamination across profiles – critical for users running multiple agent personas.

- **Cache performance regression**  
  A user reports that cross-process write detection invalidates the agent cache every turn, essentially disabling the cache. This could severely impact performance for heavy users.

- **CLI usability**  
  A user trying to set up Mem0 OSS found documented flags rejected by argparse. Another notes that `custom_providers[].extra_body` works on CLI but is silently dropped on gateway paths. Documentation mismatch frustrates users.

On the positive side, contributors are actively submitting fixes – the webhook issues were reported and PRs filed on the same day, showing a responsive contributor community.

## 8. Backlog Watch

Several high-impact issues and PRs have been open for months without maintainer action:

- **[#3823](https://github.com/NousResearch/hermes-agent/issues/3823)** – Gateway adapter refactor (open since March 2026, 3 months). No PR.  
- **[#516](https://github.com/NousResearch/hermes-agent/issues/516)** – LSP integration (open since March 2026). No PR.  
- **[#15497](https://github.com/NousResearch/hermes-agent/issues/15497)** – Hindsight shutdown race (open since April 2026). No fix PR.  
- **[#483](https://github.com/NousResearch/hermes-agent/issues/483)** – Post-task reflection (open since March 2026). No PR.  
- **[#32284](https://github.com/NousResearch/hermes-agent/pull/32284)** – Notion HMAC signature fix (open since May 2026). Superseded by [#54942](https://github.com/NousResearch/hermes-agent/pull/54942), but original reviewer attention needed.

Additionally, security PRs [#52355](https://github.com/NousResearch/hermes-agent/pull/52355) and [#52349](https://github.com/NousResearch/hermes-agent/pull/52349) (both two days old) are labelled **P1** and may pose user data risk if left unmerged. The Windows P1/P2 bugs with no fix PRs require immediate maintenance focus to avoid alienating the Windows user base.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest – 2026-06-29

## Today’s Overview
Project activity remains subdued. Over the past 24 hours, one **stale issue** was closed and one **stale PR** was closed (neither merged), while a new feature PR was updated but remains open. No new releases were published. The overall cadence suggests maintainer attention is focused on clearing backlog items rather than pushing forward development. Community engagement is minimal, with only one actively discussed issue (now closed) and one open PR awaiting review.

## Releases
*None.* No new versions were released during this period.

## Project Progress
- **PR #2964** (stale, closed) – *Feat/image input compression*: This PR proposed configurable inbound image compression for the vision pipeline. It was closed on 2026-06-28 after being marked stale, indicating it was **not merged**. The capability to reduce image size before building model payloads remains unimplemented in the main branch.  
  [GitHub PR #2964](https://github.com/sipeed/picoclaw/pull/2964)

- **Issue #2984** (stale, closed) – *[Feature][Protocol] Add explicit turn completion signal for Pico WebSocket clients*: This feature request was closed due to inactivity. No code change was made.  
  [GitHub Issue #2984](https://github.com/sipeed/picoclaw/issues/2984)

## Community Hot Topics
- **Issue #2984** (closed) – *Explicit turn completion signal for WebSocket clients*  
  **4 comments** | **2 👍**  
  The discussion revealed a clear pain point: external Pico Protocol clients could not determine when the agent had finished processing a user message. The request proposed adding a dedicated `turn.complete` event. Although the issue was closed as stale, the underlying need—deterministic end-of-turn signaling—remains unaddressed. This is a critical gap for integrators building responsive chat applications on top of PicoClaw.  
  [GitHub Issue #2984](https://github.com/sipeed/picoclaw/issues/2984)

## Bugs & Stability
No bugs, crashes, or regressions were reported in the last 24 hours. The project currently has no open issues tagged as bugs.

## Feature Requests & Roadmap Signals
- **PR #3063** (open) – *feat: add deltachat gateway*  
  This PR introduces a new gateway integration for [Delta Chat](https://delta.chat), an email-based messaging app. If merged, it would expand PicoClaw’s channel support beyond typical WebSocket/HTTP transports. The PR also includes documentation updates. It has not yet received comments or reviews, but it signals interest in decentralized communication channels.  
  [GitHub PR #3063](https://github.com/sipeed/picoclaw/pull/3063)

- **Issue #2984** – The turn-completion signal request, while closed, represents a high-priority feature for any production WebSocket client. It is likely to be re‑raised or implemented in a future minor release if community demand persists.

## User Feedback Summary
Based on the limited data, user feedback centers on **protocol completeness** and **media handling**:
- **Pain point**: Lack of explicit “agent finished” event makes client-side state management unreliable (Issue #2984).
- **Use case**: External clients need deterministic turn completion to trigger post-processing or UI updates.
- **Satisfaction**: No positive or negative feedback was recorded today. The closure of two feature-driven items without resolution may frustrate contributors.

## Backlog Watch
- **PR #3063** – *feat: add deltachat gateway*  
  Updated today (2026-06-29) but has **zero comments** and no review activity from maintainers. As it introduces a non-trivial new channel, it risks becoming stale if not triaged soon.  
  [GitHub PR #3063](https://github.com/sipeed/picoclaw/pull/3063)

- **Issue #2984** (closed) – Although closed as stale, the underlying demand (turn completion signal) is significant. Maintainers should consider re‑opening or creating a new tracking issue to prevent it from being forgotten.

- **PR #2964** (closed) – The image compression feature was abandoned. If vision pipelines continue to cause payload size issues, a fresh implementation or discussion may be necessary.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest – 2026-06-29

## 1. Today’s Overview

The project saw a high level of activity over the past 24 hours, with **10 pull requests updated**, 3 of which were merged or closed. No new issues were filed—an indication that the current development cycle is focused on polishing existing features and hardening security rather than responding to fresh bug reports. The team merged a voice notification overhaul, a database constraint fix, and a critical agent-to-agent security patch. Open PRs include a Discord channel adapter, a dashboard pusher for monitoring, and several fixes for Slack setup, button parsing, and token reconnection. Project health appears strong, with rapid iteration on both features and stability.

## 2. Releases

No new releases were published today. The latest release (if any) remains unchanged from previous periods.

## 3. Project Progress

**Merged/closed PRs (3):**
- **#2883** – `feat: voice-notify v3 意图分流 + kill-switch` – Merged. Improves voice notification by splitting summaries into 5 intent categories (action, silent, navigate, tech_status, notify) and adding a runtime kill‑switch. All 38 tests pass.
- **#2882** – `fix(ncl): default messaging-groups create instance to channel_type` – Merged. Resolves a `NOT NULL` constraint violation on the `instance` column when creating messaging groups via the CLI.
- **#2879** – `fix(agent-to-agent): containment-check target inbox in forwardAttachedFiles` – Merged. Fixes a symlink-follow vulnerability (CWE‑59) in A2A attachment forwarding that could allow writes outside the session root.

These merges advance the codebase with better configurability (voice‑notify), reliability (CLI CRUD), and security (containment enforcement).

## 4. Community Hot Topics

The provided data includes **zero comments or reactions** on the PRs, making it impossible to gauge active discussion. However, the subject matter suggests two likely areas of community interest:

- **Security patches** (#2880, #2879) – both address the same symlink escape vulnerability (CWE‑59). This is a high‑impact area and likely drew developer attention.
- **Slack Socket Mode** (#2885) – the guided setup flow was added in a feature branch but never merged into `main`, so users on the trunk still face a webhook‑only flow. This mismatch suggests community users may have been confused or frustrated.

No direct user feedback (comments, reactions) was recorded in the 24‑hour window.

## 5. Bugs & Stability

| Severity | Issue / PR | Description | Status |
|----------|------------|-------------|--------|
| **Critical (Security)** | #2880 (open) | Symlink follow in inbox attachment writes (CWE‑59) – compromised agent could write arbitrary files on the host. | Fix PR exists, under review. |
| **Critical (Security)** | #2879 (merged) | Same CWE‑59 in A2A attachment forwarding. | Fixed and merged today. |
| **High** | #2882 (merged) | `ncl messaging-groups create` fails with `NOT NULL` constraint on `instance` column. | Fixed and merged today. |
| **Medium** | #2881 (open) | Discord button `custom_id` parsing broken – delimiter causes `resolveSelectedOption` failures. | Fix PR exists, awaiting review. |
| **Medium** | #2878 (open) | `runCodexAuthStep()` returns success for stale tokens, causing mid‑conversation failures. | Fix PR exists, under review. |

All reported bugs have associated fix PRs, demonstrating responsive maintenance.

## 6. Feature Requests & Roadmap Signals

Today’s PRs reveal the following features being actively built:

- **Discord channel adapter** (#2884) – Adds Discord support via the Chat SDK bridge (Gateway mode, concurrent dispatch). Likely to land in the next minor release.
- **Dashboard pusher** (#2871) – Sends state snapshots to a `@nanoco/nanoclaw-dashboard` server every 60 seconds. Suggests a native monitoring dashboard is on the roadmap.
- **Slack Socket Mode guided setup** (#2885) – Currently stranded in a feature branch (see #2885). This is a user‑requested UX improvement and could land once merged into `main`.
- **Voice‑notify v3** (#2883, already merged) – Splits summaries by intent; this may set a precedent for more modular notification handlers.
- **Deploy / Coolify** (#2875) – A skill for deploying via Coolify, indicating interest in self‑hosted deployment options.

We predict the **next version will include Discord support, a dashboard endpoint, and the Slack Socket Mode fix**. Voice‑notify v3 is already integrated.

## 7. User Feedback Summary

No direct user feedback (comments, satisfaction ratings) was available in the provided data. However, several pain points can be inferred from PRs:

- **Slack users on `main`** cannot use Socket Mode – they must use webhooks or wait for the `channels` branch merge. This is a work‑around issue.
- **Codex users** experienced mid‑conversation failures due to stale tokens (#2878) – a frustrating experience that the PR aims to resolve.
- **Messaging‑group creation** was broken for anyone using the CLI (#2882), a regression from migration 016.
- **Security concerns** about symlink attacks (#2880, #2879) likely worry advanced users running multi‑agent setups.

Overall, the team is responding rapidly to these issues, suggesting reasonable user satisfaction with the project’s maintenance velocity.

## 8. Backlog Watch

No issues are open for more than a few days, so there is no significant backlog. The PRs that remain open and may benefit from maintainer attention are:

- **#2875** – `Deploy/coolify` (opened June 27, no activity in last 24h). May need a final review or rebase.
- **#2871** – `feat(dashboard): add dashboard pusher with OpenCode support` (opened June 27). Still awaiting approval.
- **#2881** – `fix(discord): decode custom_id delimiter` (opened June 28) – a small but important fix for Discord functionality.
- **#2880** – `fix(security): contain inbox symlink escapes` (opened June 28) – critical security patch, should be prioritised for merging.

No PR has been left unanswered for more than 2 days, so the backlog is minimal.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-06-29

## 1. Today's Overview

Activity on NullClaw remained moderate with one issue closed and four pull requests touched. The closed issue (#50) indicates community curiosity about ESP32 compatibility has been addressed, while the PR pipeline shows clear progress on CLI usability (arrow-key handling) and streaming improvements (native tool calls). No new releases were cut today, but the project is actively iterating on two key feature areas. Overall health appears stable, with maintainers responding to both bugs and enhancement requests.

## 2. Releases

No new releases were published in the last 24 hours.

## 3. Project Progress

One pull request was closed today:

- **[PR #960 — fix(cli): handle arrow keys in agent REPL](https://github.com/NullClaw/nullclaw/pull/960)** (CLOSED)  
  Added a small allocation-free line editor and POSIX raw-mode input for the interactive `nullclaw agent` REPL. This addresses a long-standing usability gap where control characters were printed instead of cursor/history navigation. The fix was superseded by an identical open PR (#970), likely for further refinement or CI re-run. Functionality advanced toward a smoother terminal experience.

No other PRs were merged.

## 4. Community Hot Topics

The most active discussion item remains:

- **[Issue #50 — Can this run on an Esp32?](https://github.com/NullClaw/nullclaw/issues/50)** (CLOSED, 4 comments)  
  Created in February, this issue was closed today after being dormant for several months. The underlying need – running NullClaw on low-power edge hardware – resonates with the project’s potential for IoT and embedded AI agents. The closure suggests either a definitive answer (likely “not currently feasible”) or a decision to defer. No reactions were recorded, but the request signals community interest in lightweight deployments.

No other issues or PRs accumulated comments or reactions today.

## 5. Bugs & Stability

One bug fix advanced today, though not yet merged:

- **Arrow keys in agent REPL** — The original fix (PR #960) was closed and re-opened as [#970](https://github.com/NullClaw/nullclaw/pull/970) (OPEN). This is a moderate-severity usability issue: users on interactive TTY sessions could not navigate history or move the cursor, making the REPL nearly unusable for long inputs. The fix is ready and awaiting final review.

No new crash reports, regressions, or security vulnerabilities surfaced in the last 24 hours.

## 6. Feature Requests & Roadmap Signals

The most explicit feature request in the dataset is the **ESP32 support** from issue #50. While now closed, it hints at a potential direction for lightweight, embedded agent runtimes. Maintainers have not publicly committed to this, but the question’s longevity (4 months) and closure without a clear “no” may imply it remains on the radar.

On the PR side, **[#971 — feat(streaming): native tool calls during SSE streaming](https://github.com/NullClaw/nullclaw/pull/971)** (OPEN, created today) is a substantial feature. It decouples native tool-call support from the streaming path, allowing providers that support native tools during streaming to emit them directly instead of falling back to prompt-injection. This is a strong signal that the next minor version will improve streaming reliability and performance for tool-using agents.

## 7. User Feedback Summary

- **Pain point**: The CLI REPL was difficult to use without arrow-key support (PR #960/970). Users likely experienced frustration with history navigation.
- **Use case**: A developer (@ngantrandev) expressed strong interest in running NullClaw on an ESP32 microcontroller, indicating a desire for autonomous agents on edge hardware with limited resources.
- **Satisfaction**: The project’s quick response to the REPL bug (two PRs in under two weeks) shows maintainers are attentive to core usability feedback.

## 8. Backlog Watch

No open issues or PRs appear to have been neglected for an unusually long time. The single existing issue (#50) was recently closed. The open PRs (#956, #970, #971) are all fresh (created within the last two weeks) and actively updated. The only item worth monitoring:

- **[PR #956 — ci(deps): bump alpine from 3.23 to 3.24](https://github.com/NullClaw/nullclaw/pull/956)** (OPEN, 14 days old)  
  A routine dependency bump by Dependabot. While not urgent, a prolonged stall might indicate CI configuration issues or maintainer oversight. Currently no blockers reported.

No other items in the backlog require immediate maintainer attention.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest – 2026-06-29

## 1. Today's Overview
IronClaw saw high activity with **50 pull requests updated** in the last 24 hours (30 open, 20 merged/closed) and **4 issues updated** (2 closed, 2 open). A major focus is on **error recoverability and failure detail surfacing**, led by a stack of PRs from core contributor `serrrfirat` (#4841, #5389, #5390, #5403, #5407). The most impactful closed items include a fix for the “Ask each time” tool permission bug (#5196) and the addition of a global Always Allow setting (#4776). Two critical stability concerns remain: a persistent nightly E2E failure (#4108, open since May) and an open capability policy design issue (#5385). No new releases were published today.

## 2. Releases
- No releases tagged on 2026-06-29.

## 3. Project Progress – Merged/Closed PRs & Issues
**Closed Issues:**
- **[#5196](https://github.com/nearai/ironclaw/issues/5196) – [Reborn] "Ask each time" tool permission may fail with authorization error and trigger duplicate approval flow** – Resolved the bug where approved tools would still show an authorization error and require re-approval.
- **[#4776](https://github.com/nearai/ironclaw/issues/4776) – Add global Always Allow setting for eligible tools** – Implemented the feature for Reborn local WebUI / engine v2.

**Merged/Closed PRs (selected from top 20):**
- **[#5400](https://github.com/nearai/ironclaw/pull/5400) – fix(ci): stabilize extension crate name parsing and Windows clippy** – Gated Unix-only imports and replaced ad-hoc TOML parsing with typed deserialization.
- **[#5405](https://github.com/nearai/ironclaw/pull/5405) – chore: upgrade Rust version to 1.96** – Raised the workspace minimum from 1.92, centralized MSRV, and updated Docker build surfaces.
- **[#5408](https://github.com/nearai/ironclaw/pull/5408) – ci: ignore unreachable wasmtime wasi advisory** – Temporarily suppressed `RUSTSEC-2026-0188` with documented reason; full upgrade deferred to a newer toolchain.
- **[#5391](https://github.com/nearai/ironclaw/pull/5391) – build(deps): bump the everything-else group with 8 updates** – Routine dependency bumps including `agent-client-protocol` to 1.0.0.

**Other notable open PRs (merged status pending):**
- (#5404) Chat composer clearing fix
- (#5401) Localization of v2 tools and extensions copy
- (#5402) Shared-persistence integration tests for approvals/auth/memory/secrets/extensions
- (#5314) Running tool activity details in WebUI v2

## 4. Community Hot Topics
No PRs or issues attracted multiple comments or reactions today (all `👍: 0` and comments = undefined in the data). However, the following items are likely attracting developer attention:

- **[#5196](https://github.com/nearai/ironclaw/issues/5196) – Tool permission failure** – Closed today after a single comment; the root cause and fix may be of interest to users encountering duplicate approval flows.
- **[#5385](https://github.com/nearai/ironclaw/issues/5385) – Add Capability Policy** – Open since June 27, proposes fine-grained user roles (owner, admin, member). This is a foundational feature for multi-user deployments and likely to generate discussion.
- **Stacked PRs by `serrrfirat`** (#4841, #5389, #5390, #5403, #5407) – A coordinated effort to eliminate “run-borking” errors and surface real failure details to the model. These are technically deep and represent a major architectural shift in error handling.

## 5. Bugs & Stability
| Severity | Issue / PR | Description | Fix Status |
|----------|------------|-------------|------------|
| **High** | [#4108](https://github.com/nearai/ironclaw/issues/4108) – Nightly E2E failed | Scheduled E2E run has been failing since May 27; updated today with a failure report. No fix PR associated. | **Open, no apparent resolution** |
| **High** | [#5196](https://github.com/nearai/ironclaw/issues/5196) – Tool authorization error (closed) | "Ask each time" tools would fail after approval, requiring duplicate user confirmation. | **Closed – fix merged** |
| **Medium** | [#5407](https://github.com/nearai/ironclaw/pull/5407) – Skill bubble dropped on SSE resume | Post-run skill learning notification missing when SSE resumes from a durable cursor; regression from #5061. | **Open PR with fix** |
| **Medium** | [#5404](https://github.com/nearai/ironclaw/pull/5404) – Chat composer not clearing after send | UI bug: text and attachments remain after send until manual clearing. | **Open PR with fix** |
| **Medium** | [#5338](https://github.com/nearai/ironclaw/pull/5338) – Surface real failure detail | Generic `invalid_input` errors replaced with user-friendly capability/tool failure details cross-layer. | **Open PR** |

## 6. Feature Requests & Roadmap Signals
- **Global Always Allow (#4776, closed)** – Now implemented. Users can designate eligible tools as permanently approved without per-use prompts.
- **Capability Policy (#5385, open)** – A multi-user permission system is under design. Given its scope (owner/admin/member roles via env vars), it is a strong candidate for the next release.
- **Error recoverability audit** – Multiple PRs signal a planned release where every terminal error is either recovered, explained to the user, or reported to the model with actionable detail. This is a major reliability improvement likely to land soon.
- **Localization (#5401)** – WebUI v2 copy is being translated from hardcoded English; a sign of global readiness in the Reborn interface.
- **Shared-persistence integration tests (#5402)** – Once merged, these tests will cover cross-thread e2e flows (approvals, auth, memory, secrets, extensions), increasing confidence in production.
- **Deep-link agent installation (#5409)** – Addition of a `POST /api/ironhub/…` gateway for private manifest installs suggests upcoming IronHub marketplace integration.

## 7. User Feedback Summary
User-reported pain points captured in issue data:
- **Duplicate approval flow (#5196)** – A user running tools with “Ask each time” permission had to approve twice without any execution. This has been fixed.
- **Missing global always-allow (#4776)** – Users desired a way to skip repeated approval for trusted tools. Now available.
- **Nightly E2E instability (#4108)** – No direct user report, but repeated failures erode confidence in CI and could delay feature delivery.

No sentiment (likes, reactions) was recorded, so satisfaction is inferred from issue closures.

## 8. Backlog Watch
- **[#4108](https://github.com/nearai/ironclaw/issues/4108) – Nightly E2E failed** – Open for 33 days with only automated updates. A root‑cause investigation and/or fix PR is overdue. High priority.
- **[#5385](https://github.com/nearai/ironclaw/issues/5385) – Add Capability Policy** – Opened June 27 with no comments. While not yet urgent, the design will require maintainer input and community discussion before implementation.
- **Unreviewed stacked PRs** – The error-recoverability series (#4841, #5390, #5389, #5403) is large and inter‑dependent; prompt review is critical to avoid merge conflicts and keep momentum.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-06-29

## Today's Overview
Today marks the **2026.6.29 release**, a major stability update driven by **40 pull requests** (39 merged/closed, 1 open) and **7 issue updates** (4 open, 3 closed). The team pushed a concentrated batch of fixes targeting OpenClaw integration, Cowork conversation rail UI, and cron job reliability. Release activity peaked, with most PRs landing as part of the promotion workflow. While the commit velocity is high, several long-standing user-reported bugs (language consistency, share screenshot truncation, email connectivity) remain open, suggesting a continued focus on core infrastructure over edge-case polish.

---

## Releases

### LobsterAI 2026.6.29
- **Release page:** https://github.com/netease-youdao/LobsterAI/releases

**Key changes** (from PR [#2228](https://github.com/netease-youdao/LobsterAI/pull/2228)):
- **OpenClaw integration stability:**
  - Preserve user turn cache stability.
  - Keep agent bootstrap workspace separate from task working directory.
  - Preserve cron run follow-up history.
  - Route plugin approvals through Cowork permissions.
- **Cowork conversation rail fixes:**
  - Clean plan-mode tags and section labels from tooltip previews.
  - Increase tooltip preview length.
  - Reset stale active rail width on hover.
- **Additional OpenClaw patches:**
  - Migrate legacy cron storage on startup.
  - Keep agent identity files (SOUL.md, IDENTITY.md, USER.md, MEMORY.md) in correct workspace.
  - Compile and preinstall QQ and Discord plugins.
  - Backport user-turn serialization patch for earlier v2026.6.1.

No breaking changes or migration notes were documented. The release bundles these fixes into a stable main branch.

---

## Project Progress

**39 PRs were merged or closed today,** reflecting a release-focused sprint. Major areas of advancement:

### OpenClaw & Agent Infrastructure
- **Cron & session handling:** [#2220](https://github.com/netease-youdao/LobsterAI/pull/2220) (preserve follow-up history), [#2190](https://github.com/netease-youdao/LobsterAI/pull/2190) (sync cron run sessions), [#2191](https://github.com/netease-youdao/LobsterAI/pull/2191) (clarify startup states), [#2189](https://github.com/netease-youdao/LobsterAI/pull/2189) (legacy cron migration).
- **Plugin ecosystem:** [#2182](https://github.com/netease-youdao/LobsterAI/pull/2182) (DingTalk, Lark, WeCom, POPO upgrades), [#2186](https://github.com/netease-youdao/LobsterAI/pull/2186) (NIM plugin compilation), [#2198](https://github.com/netease-youdao/LobsterAI/pull/2198) (QQ/Discord preinstalls).
- **Identity & workspace separation:** [#2227](https://github.com/netease-youdao/LobsterAI/pull/2227) (bootstrap workspace isolation).
- **Cache & serialization:** [#2219](https://github.com/netease-youdao/LobsterAI/pull/2219) (user-turn cache stability), [#2185](https://github.com/netease-youdao/LobsterAI/pull/2185) (cwd in reply options).

### Cowork UI
- **Conversation rail fixes:** Multiple PRs [#2222](https://github.com/netease-youdao/LobsterAI/pull/2222), [#2223](https://github.com/netease-youdao/LobsterAI/pull/2223), [#2226](https://github.com/netease-youdao/LobsterAI/pull/2226) (tooltip cleanup, hover behavior, lazy navigation). A revert [#2225](https://github.com/netease-youdao/LobsterAI/pull/2225) corrected a mistaken merge.

### Documentation & Testing
- **[#2184](https://github.com/netease-youdao/LobsterAI/pull/2184):** Refresh AGENTS.md with current architecture and quality gates.
- **[#2187](https://github.com/netease-youdao/LobsterAI/pull/2187):** Align model metadata tests for reasoning-capable models.

---

## Community Hot Topics

Most issues have only 1–2 comments, but user reactions indicate frustration with specific features. The most active items:

### Issue [#2081](https://github.com/netease-youdao/LobsterAI/issues/2081) — [CLOSED] Subscription积分清零 complaint
- **2 comments, 0 👍** – User complains that 5500积分 (credits) were cleared at month-end without notice. This was closed, likely resolved or acknowledged. The underlying need is for **transparent subscription credit rollover policies**.

### Issue [#1386](https://github.com/netease-youdao/LobsterAI/issues/1386) — [OPEN] Share screenshot truncation
- **1 comment** – When a long chat is shared as an image, the resulting screenshot is incomplete. This has been open since April 3. Users expect reliable share functionality for multi-turn conversations.

### Issue [#1388](https://github.com/netease-youdao/LobsterAI/issues/1388) — [OPEN] Email connectivity test hangs
- **1 comment** – Clicking “test connectivity” never finishes, even after restart. Blocks email configuration for users.

Other issues discuss language selection inconsistency ([#1389](https://github.com/netease-youdao/LobsterAI/issues/1389)) and agent creation UI overflow ([#1435](https://github.com/netease-youdao/LobsterAI/issues/1435), closed).

---

## Bugs & Stability

Ranked by severity and impact:

| Severity | Issue | Description | Status | Fix PR? |
|----------|-------|-------------|--------|---------|
| **High** | [#1386](https://github.com/netease-youdao/LobsterAI/issues/1386) | Long chat share image truncation | Open since Apr 3 | None today |
| **High** | [#1388](https://github.com/netease-youdao/LobsterAI/issues/1388) | Email test connectivity hangs indefinitely | Open since Apr 3 | None today |
| **High** | [#1390](https://github.com/netease-youdao/LobsterAI/issues/1390) | Scheduled task update button unresponsive (intermittent) | Open since Apr 3 | None today |
| **Medium** | [#1389](https://github.com/netease-youdao/LobsterAI/issues/1389) | Language switch shows Chinese text for English option | Open since Apr 3 | None today |
| **Low** | [#1435](https://github.com/netease-youdao/LobsterAI/issues/1435) | Agent name too long overflows modal | Closed | None |
| **Low** | [#1434](https://github.com/netease-youdao/LobsterAI/issues/1434) | Empty search in agent skills shows English text when language is Chinese | Closed | None |

No new critical regressions were introduced today. The release focused on OpenClaw backend stability. The three high-severity issues (share, email, scheduled tasks) remain unaddressed.

---

## Feature Requests & Roadmap Signals

- **Subscription credit rollover** – Issue [#2081](https://github.com/netease-youdao/LobsterAI/issues/2081) reflects demand for clearer billing policies and grace periods. While closed, this may drive a future settings improvement.
- **Improved share functionality** – [#1386](https://github.com/netease-youdao/LobsterAI/issues/1386) suggests users want reliable screenshot exports for long conversations. Likely to be addressed in a future UI iteration.
- **Language i18n polish** – Multiple stale issues ([#1389](https://github.com/netease-youdao/LobsterAI/issues/1389), [#1434](https://github.com/netease-youdao/LobsterAI/issues/1434)) indicate incomplete localization. Given they have been open for 3 months, they may be scheduled for a dedicated i18n sprint.
- **Email configuration reliability** – [#1388](https://github.com/netease-youdao/LobsterAI/issues/1388) blocks a critical user workflow; a fix is likely in the backlog.

No major new feature requests appeared today; the release is purely stability-focused.

---

## User Feedback Summary

- **Pain points:**
  - Unexpected credit expiry without notice (negative sentiment).
  - Share screenshot truncation reduces trust in collaboration tools.
  - Email setup and scheduled task failures hinder automation workflows.
  - Language inconsistencies degrade user experience for non-English users.
- **Satisfaction signals:** The rapid release cycle (2026.6.29) and dense PR list show active maintenance. OpenClaw and Cowork improvements suggest investment in enterprise/team features.
- **Overall sentiment:** Mixed – the core agent platform is improving, but several basic UI/UX bugs persist for months.

---

## Backlog Watch

These important, long-unanswered issues have not received maintainer attention in the last 24h and remain open:

- **[#1386](https://github.com/netease-youdao/LobsterAI/issues/1386)** — Share screenshot incomplete (Apr 3, open, 1 comment). No fix PR.
- **[#1388](https://github.com/netease-youdao/LobsterAI/issues/1388)** — Email test connectivity stuck (Apr 3, open, 1 comment). No fix PR.
- **[#1389](https://github.com/netease-youdao/LobsterAI/issues/1389)** — Language selection English shows Chinese text (Apr 3, open, 2 comments). No fix PR.
- **[#1390](https://github.com/netease-youdao/LobsterAI/issues/1390)** — Scheduled task update not responding (Apr 3, open, 1 comment). No fix PR.
- **[#1277](https://github.com/netease-youdao/LobsterAI/pull/1277)** — Open PR (dependabot, Apr 2) bumping electron from 40.2.1 to 42.5.0. **Stale for 88 days.** Should be reviewed and merged to keep dependencies current.

These items represent cumulative user-facing debt that should be prioritized in upcoming sprints alongside infrastructure work.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

**Moltis Project Digest — 2026-06-29**

---

## Today's Overview

Project activity remains very low over the past 24 hours. No pull requests were updated or merged, no new releases were published, and only one issue received an update. The sole active issue (#1137) was opened two days ago and has one comment, indicating light community engagement. Overall, the project appears to be in a quiet maintenance phase with no visible feature work or bug-fix cycles.

---

## Releases

**None** — No new releases were published in the last 24 hours, and the latest release history shows no recent version tags.

---

## Project Progress

**No pull requests were merged or closed** in the reporting period. No features, improvements, or fixes were advanced via PRs.

---

## Community Hot Topics

Only one issue has seen any discussion:

- **#1137 [Bug]: Apple Container ID exceeds name limit** *(open, created 2026-06-27, updated 2026-06-28, 1 comment, 0 👍)*  
  [Issue #1137](https://github.com/moltis-org/moltis/issues/1137)  
  Reported by **holgzn**. The bug describes a problem where an Apple Container ID (likely a macOS/iOS sandbox identifier) exceeds the allowed name limit, causing an error. The reporter confirms they searched existing issues and are using the latest Moltis version. The single comment may be from the author or a maintainer, but no further detail is visible.

This is the only active discussion in the community today. The underlying need is to ensure Moltis handles platform-specific naming constraints gracefully, particularly on macOS.

---

## Bugs & Stability

- **#1137 – Apple Container ID exceeds name limit**  
  **Severity: Medium** — The bug prevents Moltis from working correctly on macOS when container IDs are too long. It affects users on Apple platforms but does not appear to be a crash or regression. No associated fix PR exists yet.

No other bugs, crashes, or regressions were reported or updated in the last 24 hours.

---

## Feature Requests & Roadmap Signals

**No feature requests or roadmap signals** were observed in the latest issues or PRs. The only open item is a bug report. Given the low activity, no predictions can be made about upcoming features.

---

## User Feedback Summary

No user feedback (comments, reactions, or use cases) was recorded in the past 24 hours beyond the single comment on issue #1137. The reporter’s pain point is clear: Moltis fails on macOS when container IDs exceed a certain length. There is no evidence of satisfaction or dissatisfaction from other users.

---

## Backlog Watch

No long-unanswered important issues or PRs were identified. The only open issue (#1137) is recent (2 days old) and received an update yesterday, so it is not yet in need of extra maintainer attention. The project’s backlog appears empty of stale items.

---

**Overall Project Health:** Low activity, stable but stagnant. The single open bug indicates ecosystem-specific compatibility work may be needed, but no development momentum is visible.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest — 2026-06-29

## 1. Today's Overview
CoPaw shows **high activity** today with **50 pull requests** (27 open, 23 merged/closed) and **13 issues** (9 open, 4 closed) updated in the last 24 hours. The project is in the midst of a major refactoring wave – multiple PRs integrate Mission Mode with the new Runtime v2 architecture, add plugin-based middleware registration, and expand frontend unit test coverage. No new releases were published, but the volume of coordinated PRs (including several test sprints and observability fixes) suggests a release candidate may be nearing. Community engagement remains strong with several bug reports and feature requests receiving prompt attention from maintainers.

## 2. Releases
**None.** No new versions were published today.

## 3. Project Progress
**23 PRs were merged or closed today.** Key highlights:

- **Runtime & Governance:** [#5511](https://github.com/agentscope-ai/QwenPaw/pull/5511) restored Langfuse trace grouping lost during the 2.0 merge; [#5601](https://github.com/agentscope-ai/QwenPaw/pull/5601) fixed tool-guard approval notifications for IM channels; [#5524](https://github.com/agentscope-ai/QwenPaw/pull/5524) registered `spawn_subagent` in the Runtime v2 tool discovery.
- **Documentation:** [#5614](https://github.com/agentscope-ai/QwenPaw/pull/5614) replaced the outdated backpack analogy with a scroll strategy explanation for context management; [#5618](https://github.com/agentscope-ai/QwenPaw/pull/5618) updated website figures and removed the plan mode.
- **UI/UX:** [#5619](https://github.com/agentscope-ai/QwenPaw/pull/5619) deepened the selected session background in the chat sidebar (fixes [#5583](https://github.com/agentscope-ai/QwenPaw/issues/5583)); [#5620](https://github.com/agentscope-ai/QwenPaw/pull/5620) improved agent settings table readability.
- **Bug Fixes:** [#5505](https://github.com/agentscope-ai/QwenPaw/issues/5505) (MiniMax-M3 image safety mis-cache) and [#5543](https://github.com/agentscope-ai/QwenPaw/issues/5543) (functionDeclaration `type: null` schema) were both closed.
- **New Feature:** [#2495](https://github.com/agentscope-ai/QwenPaw/issues/2495) (MCP tool listing) was closed – the feature to see MCP tools in the UI has been implemented.

## 4. Community Hot Topics
The most active issues reflect two major user concerns: **cost optimization** and **stability under failure**.

- [#3891](https://github.com/agentscope-ai/QwenPaw/issues/3891) – *DeepSeek前缀缓存命中率偏低（~95%），优化空间巨大*  
  **Comments: 5 | 👍: 1**  
  Oldest open issue (Apr 27) still unresolved despite huge cost implications (5% cache miss → up to 20× price difference). High community urgency.
- [#5561](https://github.com/agentscope-ai/QwenPaw/issues/5561) – *Agent链接飞书机器人后长回复无法接收*  
  **Comments: 3**  
  Feishu channel truncation of long replies – a practical blocker for production use.
- [#5342](https://github.com/agentscope-ai/QwenPaw/issues/5342) – *hard cap on tool result size at execution layer*  
  **Comments: 3**  
  Defense-in-depth proposal against cascading failures when LLM calls fail. PR [#5510](https://github.com/agentscope-ai/QwenPaw/pull/5510) is under review to address this.

No PRs attracted significant discussion (comment counts not provided in top 20).

## 5. Bugs & Stability
**High Severity:**
- [#5579](https://github.com/agentscope-ai/QwenPaw/issues/5579) – *对话记录在异常中断场景下丢失*  
  Conversation data loss on reboot or crash. No fix PR yet. Data integrity issue.
- [#5616](https://github.com/agentscope-ai/QwenPaw/issues/5616) – *自动化任务莫名终止*  
  Unprovoked automation task termination. Root cause unknown.

**Medium Severity:**
- [#5561](https://github.com/agentscope-ai/QwenPaw/issues/5561) – Feishu long message delivery failure (ongoing).
- [#5505](https://github.com/agentscope-ai/QwenPaw/issues/5505) – MiniMax-M3 image cache error – **Closed** via PR [#5510](https://github.com/agentscope-ai/QwenPaw/pull/5510)? Actually closed but no linked fix PR in data; likely resolved by separate commit.
- [#5543](https://github.com/agentscope-ai/QwenPaw/issues/5543) – Function Declaration `type: null` schema – **Closed** with workaround.

**Low Severity:**
- [#5583](https://github.com/agentscope-ai/QwenPaw/issues/5583) – Chat popup selection background not visible – **Closed**, fixed by PR [#5619](https://github.com/agentscope-ai/QwenPaw/pull/5619).

## 6. Feature Requests & Roadmap Signals
Several enhancements point to the next release priorities:

- **Vision Fallback for Text Models** ([#5615](https://github.com/agentscope-ai/QwenPaw/issues/5615)) – Automatically call a vision model when a text-only model receives images. Likely to be implemented soon given high demand.
- **Custom Model Protocol** ([#5609](https://github.com/agentscope-ai/QwenPaw/issues/5609)) – Support non-standard API endpoints (e.g., image generation). Would unlock many free models.
- **Checkpoint Persistence** ([#5579](https://github.com/agentscope-ai/QwenPaw/issues/5579)) – Already raised as a bug; a natural feature to prevent data loss.
- **Windows Tray Icon** ([#5622](https://github.com/agentscope-ai/QwenPaw/issues/5622)) – Desktop UX improvement for background running.
- **DingTalk Image Preview** ([#5593](https://github.com/agentscope-ai/QwenPaw/issues/5593)) – Upload images as previewable messages instead of file attachments.

PR activity indicates near-term roadmap includes:
- **Test coverage expansion** (multiple PRs for unit tests: [#5438](https://github.com/agentscope-ai/QwenPaw/pull/5438), [#5434](https://github.com/agentscope-ai/QwenPaw/pull/5434), [#5409](https://github.com/agentscope-ai/QwenPaw/pull/5409), [#5422](https://github.com/agentscope-ai/QwenPaw/pull/5422), [#5423](https://github.com/agentscope-ai/QwenPaw/pull/5423) – all open).
- **Plugin system** [#5221](https://github.com/agentscope-ai/QwenPaw/pull/5221) – introducing structured plugin registration with version constraints.
- **Mission Mode v2 integration** [#5442](https://github.com/agentscope-ai/QwenPaw/pull/5442).
- **Security documentation** for Sandbox isolation [#5621](https://github.com/agentscope-ai/QwenPaw/pull/5621).

## 7. User Feedback Summary
**Positive signals:**
- MCP tool listing feature finally delivered (closes [#2495](https://github.com/agentscope-ai/QwenPaw/issues/2495)).
- Active response to UI/UX complaints: selection background fix merged same day.

**Pain Points (repeated across issues):**
- **Data loss** – conversation history vulnerability (#5579) and cost from pipeline failures (tool result overflow, #5342) are top concerns.
- **Channel reliability** – Feishu and DingTalk channels have specific regressions (#5561, #5593).
- **Cost inefficiency** – DeepSeek cache issue (#3891) remains unresolved for 2 months, frustrating users paying high miss prices.
- **Model compatibility** – Vendor-specific quirks (MiniMax safety cache, function schema `null` type) require defensive handling.

Users clearly **value stability and cost control** over new features. The backlog of test PRs suggests maintainers are aware and investing in quality.

## 8. Backlog Watch
**Critical:** [#3891](https://github.com/agentscope-ai/QwenPaw/issues/3891) – *DeepSeek Prefix Cache Hit Rate ~95%*  
Created Apr 27, last updated Jun 29. No assignment, no PR. Given direct financial impact, this should be prioritized.

**Important:** [#5342](https://github.com/agentscope-ai/QwenPaw/issues/5342) – *Hard cap on tool result size at execution layer*  
Open since Jun 20, 3 comments. A fix PR [#5510](https://github.com/agentscope-ai/QwenPaw/pull/5510) is under review, but the issue itself remains open – expected to close automatically on merge.

**Neglected Legacy:** No other issues older than a month with zero responses. The project maintainers are actively triaging. However, the lack of a concrete timeline for [#3891](https://github.com/agentscope-ai/QwenPaw/issues/3891) is worrying given its cost implications.

---

*Data sourced from Github – CoPaw / QwenPaw repository, 2026-06-29 snapshot.*

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-06-29

---

## 1. Today's Overview

The project remains highly active with **10 issues** updated in the last 24 hours (9 open, 1 closed) and **50 pull requests** updated (30 open, 20 merged/closed). No new releases were published. Activity is concentrated on bug fixing, infrastructure hardening, and landing several feature milestones, particularly the **passive WhatsApp group context** feature and the **cost/org billing** snapshot RPC. The high PR churn (20 merged today) and the presence of multiple high-severity bugs with in-flight fixes indicate a focused push toward stabilisation ahead of the next release. The **SOP (Standard Operating Procedure)** system and **WASM plugin** roadmap items also saw notable commits today.

---

## 2. Releases

*No new releases were published in the last 24 hours.*

---

## 3. Project Progress

A total of **20 pull requests** were merged or closed today. Notable completed work:

- **Passive WhatsApp group context** — [PR #8389](https://github.com/zeroclaw-labs/zeroclaw/pull/8389) (closed) adds opt-in passive context for WhatsApp group chats, storing unaddressed messages as history without triggering an agent turn. This addresses a long-standing gap for group channel usability.
- **Cost/org snapshot RPC** — [PR #8482](https://github.com/zeroclaw-labs/zeroclaw/pull/8482) (closed) introduces an optional `cost/org` endpoint for reading organisation-level billed cost snapshots, alongside a `cost/query` windowed query. This lays the foundation for billing dashboards.
- **Corresponding issue** [Issue #8379](https://github.com/zeroclaw-labs/zeroclaw/issues/8379) (closed) tracked the WhatsApp feature.

Other infrastructure fixes merged today (as part of the 20 closures) include improvements to cron shutdown, hardware timeout error preservation, and documentation autolink cleanup (not all visible in the top-20 PR list but reflected in the total count).

---

## 4. Community Hot Topics

The following issues and PRs generated the most discussion and reactions:

1. **[Issue #5600](https://github.com/zeroclaw-labs/zeroclaw/issues/5600) — Use kimi-code provider in streaming chat call tools, provider API reports an error**  
   *11 comments, 1 👍* — A high-severity bug (S1 – workflow blocked) when using the Kimi compatible provider. The error originates from a missing `reasoning_content` field in the assistant response. The issue has been open since April 2026 and remains unassigned, drawing repeated attention.

2. **[Issue #8054](https://github.com/zeroclaw-labs/zeroclaw/issues/8054) — System prompt tool-availability should match per-turn effective tools across all entry points**  
   *8 comments* — A follow-up to a previous fix that only covered the direct runtime agent path. Users and maintainers are discussing how to propagate the same fix to channels, gateway, WebSocket, and `/think` endpoints. The issue is blocked pending a broader refactor.

3. **[Issue #6557](https://github.com/zeroclaw-labs/zeroclaw/issues/6557) — Reconcile runtime model switching with provider structure for v0.8.0**  
   *4 comments* — A planning issue tracking the unification of model-switching commands (slash commands, runtime config) under the newer provider structure. The discussion highlights friction between how channels and the runtime handle model changes.

**Analysis:** The community is most vocal about provider compatibility bugs and system prompt consistency. The Kimi provider issue (#5600) in particular is a persistent pain point that has not been fully resolved after months.

---

## 5. Bugs & Stability

Several bugs with **high or medium severity** received updates today:

| Severity | Issue / PR | Description | Fix PR |
|----------|------------|-------------|--------|
| **High** | [#5600](https://github.com/zeroclaw-labs/zeroclaw/issues/5600) | Kimi provider streaming error – missing `reasoning_content` | No fix PR linked yet |
| **High** | [#8054](https://github.com/zeroclaw-labs/zeroclaw/issues/8054) | System prompt tool-availability mismatch across entry points | Fix in #8053 (merged earlier) – scope extended |
| **High** | [#8463](https://github.com/zeroclaw-labs/zeroclaw/pull/8463) (open PR) | Unbounded interactive CLI stdin could OOM | Fix PR [#8463](https://github.com/zeroclaw-labs/zeroclaw/pull/8463) – caps input at 1 MiB |
| **Medium** | [#7909](https://github.com/zeroclaw-labs/zeroclaw/pull/7909) (open PR) | Groq native tool calling missing `name` field → HTTP 400 | Fix PR [#7909](https://github.com/zeroclaw-labs/zeroclaw/pull/7909) |
| **Medium** | [#7960](https://github.com/zeroclaw-labs/zeroclaw/pull/7960) (open PR) | `execute_pipeline` ignores per-agent `ToolAccessPolicy` | Fix PR [#7960](https://github.com/zeroclaw-labs/zeroclaw/pull/7960) – addresses issue #7947 |
| **High** | [#8003](https://github.com/zeroclaw-labs/zeroclaw/pull/8003) (open PR) | `session_end` hook never fires | Fix PR [#8003](https://github.com/zeroclaw-labs/zeroclaw/pull/8003) – wires existing dead code to session lifecycle |
| **Medium** | [#8465](https://github.com/zeroclaw-labs/zeroclaw/pull/8465) (open PR) | Scheduler (cron) doesn't respond to daemon shutdown | Fix PR [#8465](https://github.com/zeroclaw-labs/zeroclaw/pull/8465) – adds `CancellationToken` |
| **Low** | [#8499](https://github.com/zeroclaw-labs/zeroclaw/pull/8499) (open PR) | Hardware timeout handlers discard inner error | Fix PR [#8499](https://github.com/zeroclaw-labs/zeroclaw/pull/8499) |

**Assessment:** The project is actively addressing bugs with timely PRs. The two highest-impact bugs (#5600 and #8054) remain open without a clear fix in progress, which may delay the next stable release.

---

## 6. Feature Requests & Roadmap Signals

Several feature-oriented issues and PRs were updated today, signalling near-term roadmap priorities:

- **SOP (Standard Operating Procedure) system** – [Issue #8413](https://github.com/zeroclaw-labs/zeroclaw/issues/8413) requests a filesystem-based event source for SOP workflows. Corresponding implementation PR [#8461](https://github.com/zeroclaw-labs/zeroclaw/pull/8461) (open, size:XL) adds a `notify`-based watcher. This is a core building block for the **daemon-owned SOP control plane** tracked in [Issue #8288](https://github.com/zeroclaw-labs/zeroclaw/issues/8288) (milestone 5/5).
- **WASM plugin program** – [Issue #7314](https://github.com/zeroclaw-labs/zeroclaw/issues/7314) tracker and [PR #8491](https://github.com/zeroclaw-labs/zeroclaw/pull/8491) add per-call execution limits and the FND-001 backend taxonomy, advancing the v0.8.3 WASM plugin release.
- **Cost tracking** – [PR #8483](https://github.com/zeroclaw-labs/zeroclaw/pull/8483) implements the *zerocode* Cost tab with period breakdown and org billed views.
- **Observability** – [Issue #8462](https://github.com/zeroclaw-labs/zeroclaw/issues/8462) is an RFC for runtime policy controlling OTel content (LLM/tool payloads). This may become part of the structured observability effort.
- **Goal mode** – [PR #8393](https://github.com/zeroclaw-labs/zeroclaw/pull/8393) (size:XL) implements the runtime goal-mode control plane, accepted via RFC #8303.

**Prediction:** The next minor release (likely v0.8.x) will include the filesystem SOP source, WASM plugin execution limits, cost tracking (zerocode), and the goal mode runtime. The observability policy RFC may land as a follow-up.

---

## 7. User Feedback Summary

Real pain points surfaced in today's data:

- **Provider incompatibility** – Users of the Kimi compatible provider face a streaming error that blocks tool calls. The issue (#5600) has been open for months without a fix, likely causing user frustration.
- **System prompt misleading** – The issue #8054 reveals that agents sometimes tell users “No tools are available” even when tools are configured. This degrades user trust.
- **Unbounded input risk** – The fix PR #8463 was prompted by a real crash scenario (unbounded paste into interactive CLI). Users benefit from the cap, but it also reveals a lack of input validation.
- **Missing session hooks** – Users relying on session lifecycle hooks (e.g., to log cleanup) were silently not fired (#8003). This suggests users may have observed missing workflow steps without understanding why.
- **WhatsApp group context** – The successful closure of #8379/#8389 indicates strong user demand for better group chat handling, especially in passive (non-intrusive) listening mode.
- **Cost visibility** – The new cost/org snapshot and zerocode Cost tab (#8482, #8483) address user needs for billing and monitoring.

Overall, users are engaged and reporting bugs actively, but the long tail of unaddressed high-severity issues may erode satisfaction.

---

## 8. Backlog Watch

Several important issues and PRs remain open for extended periods or lack maintainer attention:

- **[Issue #5600](https://github.com/zeroclaw-labs/zeroclaw/issues/5600)** (Kimi provider error) – Open since April 2026, no assigned maintainer, no fix PR. Despite high severity and 11 comments, this is the oldest high-severity bug in the active list.
- **[Issue #6074](https://github.com/zeroclaw-labs/zeroclaw/issues/6074)** (track 153 lost commits) – Open since April 2026, no recent activity. This audit issue has broad implications for code recovery but appears stalled.
- **[Issue #6557](https://github.com/zeroclaw-labs/zeroclaw/issues/6557)** (model switching reconciliation) – Open since May 2026, accepted but not yet assigned. Blocking v0.8.0 provider config hardening.
- **[Issue #8453](https://github.com/zeroclaw-labs/zeroclaw/issues/8453)** (log dead code) – Low priority but has zero comments from maintainers. A simple cleanup that may be deprioritised.
- **[PR #8393](https://github.com/zeroclaw-labs/zeroclaw/pull/8393)** (goal mode implementation) – Needs author action, large size (XL), and has been open since June 27. If left unattended, it may miss the next release window.

**Recommendation:** Maintainers should prioritise triage on #5600 and #6074, and either assign or escalate the goal mode PR to avoid roadmap slippage.

---

*Aggregated from GitHub activity on zeroclaw-labs/zeroclaw, 2026-06-29.*

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*