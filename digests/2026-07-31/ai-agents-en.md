# OpenClaw Ecosystem Digest 2026-07-31

> Issues: 370 | PRs: 500 | Projects covered: 13 | Generated: 2026-07-31 00:15 UTC

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

# OpenClaw Project Digest — 2026-07-31

## Today’s Overview

OpenClaw continues to show very high community and development activity, with **370 issues** and **500 pull requests** updated in the last 24 hours. The vast majority of these remain open (366 issues, 424 PRs), indicating a large backlog of open work and active triage. No new releases were published today. The most active discussions centre on critical UX and reliability bugs — especially where the agent’s internal processing text leaks to messaging channels, crash-loop breakers permanently suppress channels, or bootstrap configuration files are silently ignored. Several high-severity security and session‑state issues also dominate the conversation. On the PR side, maintainers are focused on fixing sandbox workspace misconfiguration, command hang issues with `/steer`, and CLI flag handling, while larger changes (Android lifecycle preservation, policy previews) await review.

## Releases

None today.

## Project Progress

Today **76 PRs were merged or closed** (out of the 500 updated). While not all are individually listed in the top 30, the following key merged/closed items were observed:

- **#18778** (closed) – Discord canvas support (merged previously, now closed)
- **#84978** (closed) – Discord desktop proof draft (merged previously, now closed)

Among the open PRs that advanced with significant activity:

- **#116223** fix(agents): honor configured sandbox skill workspace – addresses custom‑workspace agents using the default workspace
- **#116565** fix(codex): `/steer` hangs indefinitely when app‑server does not answer – resolves a Telegram control lane lockup
- **#116587** fix(cli): accept inherited flags after nested subcommands – fixes `--agent` and `--json` placement inconsistency
- **#116220** fix(trajectory): record conversation state once per turn, preserve prompt in truncated events – prevents silent prompt loss in oversized events
- **#116593** fix(android): preserve device work across lifecycle changes – closes issue #116592 for Android/WearOS
- **#116579** fix(gateway): avoid false port‑busy reports behind Tailscale Serve – resolves a misleading startup error
- **#116591** fix(exec): arm the gateway node invoke deadline from exec – follow‑up to #115248

Several PRs remain in “waiting on author” or “needs proof” status, indicating that maintainers are actively requesting more evidence or revisions.

## Community Hot Topics

The following issues and PRs gathered the most comments and reactions over the past 24 hours, revealing the community’s top concerns:

- **#25592** (38 comments, 1 👍) – “Text between tool calls leaks to messaging channels”  
  _P1, diamond lobster._ A critical UX and security bug: internal agent processing output is sent to Slack/iMessage. Underlying need: isolation of internal reasoning from user‑visible messages.

- **#115326** (20 comments) – “Crash-loop breaker suppresses Discord/WhatsApp permanently; documented recovery fails with WebSocket 1006”  
  _Bug, P1, silver shellfish._ A regression where the gateway permanently disables channels after a crash loop, and the official `channels.start` command fails. Users are effectively locked out.

- **#22438** (17 comments) – “Tiered bootstrap file loading for progressive context control”  
  _P2, off‑meta tidepool._ Feature request to allow per‑session and per‑agent bootstrap file selection to save token budget. Strong consensus on need.

- **#29387** (15 comments, 5 👍) – “Bootstrap files in agentDir are silently ignored – only workspace files injected”  
  _Bug, P1, diamond lobster._ Users discover that per‑agent SOUL.md etc. are not loaded. Frustration with inconsistent documentation/implementation.

- **#50090** (15 comments, 2 👍) – “Community Skill Development & ClawHub”  
  _P2, silver shellfish._ Discussion of the gap between the ClawHub promise and reality: skill publishing, installation, and ecosystem health.

- **#48003** (15 comments, 4 👍) – “Steer mode does not inject messages mid‑turn for main sessions”  
  _P1, platinum hermit._ A regression since commit 9889c6da5 that breaks the steer queue mechanic.

- **#116201** (8 comments) – “Realtime voice work can retain unbounded provider and consult state”  
  _Bug, P1, gold shrimp._ Memory leak in voice sessions – maintainer‑tagged, likely related to the PR #116589 fixing meeting bot playback.

Several other issues with high comment counts (e.g., #53628, #51429, #54531, #50093) highlight configuration variable expansion bugs, hardcoded workspace paths, and channel delivery failures.

## Bugs & Stability

Bugs reported or updated today, ranked by severity and community impact:

| Issue | Severity | Summary | Fix PR exists? |
|-------|----------|---------|----------------|
| #25592 | **P1 Diamond Lobster** | Internal text leaked to messaging channels (conversation & security) | No specific fix PR yet |
| #115326 | **P1 Silver Shellfish** | Crash‑loop breaker permanently disables Discord/WhatsApp; recovery broken | No |
| #29387 | **P1 Diamond Lobster** | Bootstrap files in agentDir silently ignored | No |
| #48003 | **P1 Platinum Hermit** | Steer mode broken: messages not injected mid‑turn | No |
| #115001 | **P2 Platinum Hermit** | Hybrid memory search returns spurious 1.0 similarity scores via FTS fallback | No |
| #51396 | **P1 Diamond Lobster** | `clearUnboundScopes` strips operator scopes unconditionally for token‑auth clients | No |
| #100778 | **P1 Diamond Lobster** | Preflight compaction failure permanently locks Composer into “terminated” state | No |
| #69118 | **P1 Diamond Lobster** | Claude CLI sessions reset on every turn in group channels due to `extraSystemPromptHash` drift | No |
| #50165 | **P2 Silver Shellfish** | Subagents appear completed before work actually finishes | No |
| #48810 | **P1 Platinum Hermit** | Compaction retry creates orphan fork in parentId chain | No |

No critical (P0) bugs were updated today except #48920 (Live Docs ahead of release, rated P0 but lower impact). The absence of fix PRs for many high‑severity issues suggests a bottleneck in either triage or review capacity.

## Feature Requests & Roadmap Signals

Several feature requests gathered significant attention, signalling clear community desires:

- **#22438** – Tiered bootstrap file loading (progressive context control). Likely to be prioritised given its ties to token cost management.
- **#50090** – Community Skill Development & ClawHub ecosystem improvement. High visibility; may appear as a platform initiative.
- **#22358** – Post‑subagent completion extension hook. Useful for trajectory generation and auditing.
- **#20786** – Telegram Business Bot support (business_message updates). 6 👍 – strong interest from business users.
- **#80213** – Skill author‑defined setup hook (run scripts on install/update). 4 👍 – ecosystem growth.
- **#67413** – Per‑agent dreaming configuration. 5 👍 – memory management.
- **#96675** – Owner‑signed responsibility gates for memory/actions. Privacy and consent feature.
- **#60572** – Multi‑Slot Memory Architecture. 3 👍 – deeper memory customisation.
- **#63990** – Multi‑index embedding memory with model‑aware failover. Production reliability.
- **#53548** – Decouple mode=”session” from thread binding requirement. 3 👍 – session management flexibility.

**Prediction for next release:** Given the volume of P1 regression bugs and the community’s focus on session reliability, the next release (likely a hotfix) will address #25592, #115326, and #48003. Feature work on bootstrap tiering (#22438) and ClawHub (#50090) may also appear as experimental previews.

## User Feedback Summary

Real user pain points and satisfaction signals extracted from today’s data:

- **Frustration with configuration inconsistency**: Multiple reports of bootstrap files being ignored (#29387), XDG_CONFIG_HOME not expanded (#53628), hardcoded workspace paths (#51429), and `cacheRetention` for LiteLLM proxies ignored (#37966).
- **Channel delivery reliability breakdowns**: Users on Discord, WhatsApp, Telegram, and Slack are experiencing message loss or channel suppression after crash loops (#115326, #54531, #50093). The “steer” mode breakage (#48003) also disrupts interactive use.
- **Trust and safety concerns**: Agent internal text leaking to channels (#25592), cron sessions hallucinating output instead of failing (#49876), and subagent completion state inaccuracy (#50165) erode user trust.
- **Positive community engagement**: Users are actively proposing and discussing detailed feature requests (e.g., tiered bootstrap, skill setup hooks, multi‑slot memory), showing a healthy desire to contribute to the project’s direction.
- **Tool and integration gaps**: Lack of Telegram Business support (#20786), missing inline buttons in webchat (#46656), and no pagination for message listing (#71452) limit advanced use cases.

Overall sentiment is a mix of high enthusiasm for the platform’s potential and acute frustration with regressions that break everyday usage. The large volume of open issues (366) suggests the community is very engaged but also that the maintenance team may be stretched.

## Backlog Watch

Long‑standing important issues and PRs that have not received recent maintainer action (based on label clues like `needs-product-decision`, `waiting on author`, or stale flags):

| Issue/PR | Age | Status | Why it matters |
|----------|-----|--------|----------------|
| #25592 | Created Feb 24, 2026 | Open, needs‑product‑decision | **P1** security/UX leak – untouched for months |
| #22438 | Created Feb 21, 2026 | Open, needs‑product‑decision | High‑impact feature with 17 comments and no maintainer decision |
| #50090 | Created Mar 19, 2026 | Open, needs‑product‑decision, stale | Community skill ecosystem – core to project growth |
| #20786 | Created Feb 19, 2026 | Open, needs‑product‑decision | 6 👍 – Telegram Business is widely requested |
| #50199 | Created Mar 19, 2026 | Open, needs‑product‑decision, P3 | Skill priority configuration – basic UX improvement |
| #115326 | Created Jul 28, 2026 | Open, needs‑maintainer‑review | Critical regression – only 2 days old but already requires attention |
| **PR #101916** (policy preview repairs) | Created Jul 7, 2026 | Open, status “ready for maintainer look” | Large PR with compatibility/security‑boundary risk – awaiting merge |
| **PR #101813** (remove local dependency denylist) | Created Jul 7, 2026 | Open, ready for maintainer look | Hits core plugin install trust boundaries |

These items represent either critical 🔴 P1 bugs that have been open for months without resolution, or features that would unblock large parts of the user experience. The project would benefit from a triage push to either close or assign product decisions to these long‑standing items.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report — 2026-07-31

## 1. Ecosystem Overview

The open-source personal AI agent ecosystem continues to expand rapidly, with projects ranging from reference implementations (OpenClaw) to specialized assistants (NanoBot, Hermes Agent) and lightweight forks (NanoClaw, ZeptoClaw). Activity is concentrated around reliability fixes, channel parity (Slack, Telegram, IRC, Matrix), and security hardening — reflecting a maturing ecosystem where users expect production-grade stability. A few projects (NullClaw, TinyClaw) show zero activity, while others like IronClaw and CoPaw are in aggressive development cycles. Common pain points across the board include memory leaks, configuration inconsistencies, and credential exposure in subprocesses.

## 2. Activity Comparison

| Project | Issues Updated (last 24h) | PRs Updated | Merged/Closed PRs | Release (today/recent) | Health Score |
|---|---|---|---|---|---|
| OpenClaw | 370 (366 open) | 500 (424 open) | 76 | None | Medium |
| NanoBot | 7 (5 open) | 49 (16 open) | 33 | None | High |
| Hermes Agent | 14 (all open) | 50 (47 open) | 3 | v0.19.1 (yesterday) | Medium |
| PicoClaw | 7 (4 open) | 17 (12 open) | 5 | None | Medium |
| NanoClaw | 2 (new) | 19 (12 open) | 7 | None | High |
| NullClaw | 0 | 0 | 0 | – | Dormant |
| IronClaw | 19 (all open) | 50 (27 open) | 23 | None | High |
| LobsterAI | 0 | 10 (2 open) | 8 | 2026.7.29 (yesterday) | High |
| TinyClaw | 0 | 0 | 0 | – | Dormant |
| Moltis | 2 (both open) | 5 (3 open) | 2 | None | Medium |
| CoPaw (QwenPaw) | 15 (11 open) | 50 (24 open) | 26 | None | High |
| ZeptoClaw | 0 | 1 (open) | 0 | None | Low |
| ZeroClaw | 2 (new) | 50 (all open) | 0 | None | Medium |

**Notes:** Health Score considers merge velocity, bug severity, release cadence, and community engagement. OpenClaw’s raw volume is highest, but the large backlog of unaddressed P1 bugs and low merge-to-open ratio lowers its relative score.

## 3. OpenClaw’s Position

**Advantages:**  
- Largest ecosystem by far — 370+ issues, 500+ PRs, and highest community engagement (38 comments on top issue).  
- Core reference design — sets architectural patterns (sandbox, steer, trajectory) that forks emulate.  
- Broadest channel support (Discord, WhatsApp, Telegram, Slack, iMessage) and deep integration with MCP.

**Technical Differences:**  
- Uses a heavily layered architecture with dedicated gateway, exec, trajectory, and sandbox services.  
- Steer mode and bootstrap file loading are distinguishing features not fully replicated elsewhere.  
- Memory model includes hybrid FTS+embedding and “dream” compression.

**Community Size Comparison:**  
- OpenClaw’s daily active contributors and issue volume dwarf all other projects. For example, NanoBot (49 PRs) and IronClaw (50 PRs) are active but much smaller.  
- However, slower triage of critical bugs (e.g., text leak #25592 open since February) erodes trust.

**Weakness relative to peers:**  
- NanoBot merged 33 PRs today vs. OpenClaw’s 76 but from a far smaller pool; merge efficiency (76/500 = 15%) is lower than NanoBot (33/49 = 67%).  
- Several P1 issues lack any fix PR, while peers (e.g., ZeroClaw) produce fix PRs the same day for S0 bugs.

## 4. Shared Technical Focus Areas

Requirements emerging across multiple projects:

| Focus Area | Affected Projects | Specific Needs |
|---|---|---|
| **Security hardening** | OpenClaw, ZeroClaw, Moltis, ZeptoClaw | Credential leakage to subprocesses (ZeptoClaw #645); unauthenticated webhooks (ZeroClaw #9565); vault auth (Moltis #1177); tool call leak to channels (OpenClaw #25592). |
| **Channel reliability** | OpenClaw, NanoBot, Hermes, CoPaw, ZeroClaw | Crash-loop breaker suppressing channels (OpenClaw #115326); Telegram polling stalls (NanoBot #5156); Slack memory leak (IronClaw #6900); Matrix E2EE breakage (CoPaw #6476). |
| **Performance regressions** | OpenClaw, CoPaw, IronClaw | ~2s fixed overhead per reply (CoPaw #6307); 4096-token context cap on Ollama (Hermes #43900); UI freezes on large shell output (CoPaw #6589). |
| **Memory & context management** | OpenClaw, Hermes, CoPaw | Hybrid memory similarity spurious (OpenClaw #115001); Dream compression misses early events (CoPaw #6555); unbounded memory growth in exec sessions (NanoBot #5150). |
| **Skill / plugin ecosystem** | OpenClaw, NanoBot, CoPaw, LobsterAI | ClawHub community skill development (OpenClaw #50090); skill install from registry branches broken (NanoClaw #3155); author-defined setup hooks (OpenClaw #80213). |
| **MCP integration** | PicoClaw, CoPaw, ZeroClaw | OAuth 2.1 for MCP servers (PicoClaw #2546); session not recovered after server restart (CoPaw #6524); MCP YAML parsing quirks (Hermes #75093). |
| **Cross-platform parity** | OpenClaw, PicoClaw, Moltis, Hermes | Session management missing on Telegram (PicoClaw #3307); IRC long message splitting (PicoClaw #3287); Telegram inline buttons (Moltis #1178); WeCom silent failures (Hermes #29667). |

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | Hermes Agent | IronClaw | CoPaw (QwenPaw) | ZeroClaw | LobsterAI |
|---|---|---|---|---|---|---|---|
| **Primary focus** | Core reference, full agent platform | Lightweight assistant, WebUI first | Plugin extensibility, TTS, multi-provider | Enterprise Slack/WebUI, architecture overhaul | Computer Use, sandbox, creator tools | Security, eval pipeline, email | Cowork, enterprise isolation, UI polish |
| **Target users** | Developers building custom agents | End users, hobbyists | Power users, plugin developers | Enterprise teams, Slack users | General users, desktop app | Security-conscious ops | Enterprise, Chinese market |
| **Release cadence** | Frequent but patch-heavy | Rapid fixes, no releases | Tagged patches monthly | No recent release, restructuring | v2.0 recently, hotfix expected | No recent release | Steady patch releases |
| **Community size** | Very large | Medium | Medium | Medium-Large | Medium-Large | Small | Small |
| **Architectural strength** | Deepest channel/ tool integration | Fastest issue-to-fix cycle | Plugin system, provider fallback | Cross-channel file flows, componentized | Computer Use, Dream memory | Security-first design | Clean UI, enterprise auth |
| **Weakness** | Bug triage bottleneck, backlog | Channel-specific regressions | Ollama context cap, cost tracking | Cross-user memory leak, release blocked | Performance regression from v1.x | Many open PRs, zero merges today | Low community engagement |

## 6. Community Momentum & Maturity

**Tier 1 – High velocity / rapid iteration:**  
- **NanoBot** – 33 merges in 24h, quick bug fixes, active feature rollout (Quick Chat, SQLite migration).  
- **IronClaw** – 23 merges, executing target crate architecture, new Slack commands, heavy refactoring.  
- **CoPaw (QwenPaw)** – 26 merges, Computer Use, CI fixes, strong community contribution (fork PRs).  
- **NanoClaw** – 7 merges, image hardening, skill ecosystem additions, coordinated push.

**Tier 2 – Steady but gated:**  
- **OpenClaw** – High absolute activity but slow on critical fixes; large backlog dilutes momentum.  
- **Hermes Agent** – Patch release yesterday, but only 3 merges today; many open PRs and old bugs.  
- **PicoClaw** – Moderate merges, dependency bumps, steady but not accelerating.  
- **LobsterAI** – Good merge rate, stable releases, but zero community feedback today.  
- **Moltis** – Small but active core; 2 merges, new security bug and feature request.

**Tier 3 – Low activity / dormant / early:**  
- **NullClaw, TinyClaw** – No activity, likely inactive.  
- **ZeptoClaw** – Single PR addressing security, otherwise dormant.  
- **ZeroClaw** – High PR count but zero merges; security fixes in review; needs maintainer bandwidth.

**Maturity signals:**  
- Projects with recent releases (Hermes v0.19.1, LobsterAI 2026.7.29) are stabilizing.  
- IronClaw and CoPaw are in active restructuring, risking regressions but showing long-term investment.  
- OpenClaw remains the de facto standard but must address the growing backlog to retain developer trust.

## 7. Trend Signals

1. **Security is the top concern.** Three separate projects (ZeroClaw S0 webhook auth, Moltis vault auth, OpenClaw tool call leak) highlight that credential leakage and unauthenticated entry points are the most urgent class of bugs. Developers should prioritize subprocess isolation and webhook verification in their own agents.

2. **Performance regressions from v2.0 migrations are painful.** CoPaw’s 2s overhead and Hermes’ Ollama context cap show that major version jumps can silently degrade UX. The community demands automated performance regression testing.

3. **Channel parity is still incomplete.** Telegram, WhatsApp, IRC, and Slack all have missing features (inline buttons, session management, E2EE). The expectation is that agents should behave identically across platforms.

4. **Skill/plugin ecosystems are immature.** OpenClaw’s ClawHub gaps and NanoClaw’s registry branch drift indicate that “app store” features are still aspirational. Standards for skill packaging and verification are needed.

5. **Memory and context management remain unsolved.** Hybrid search spurious scores, Dream compression gaps, and unbounded memory growth recur across projects. The community is moving toward configurable, tiered memory architectures.

6. **Local model support is growing but fragile.** Hermes’ Ollama context cap, ZeroClaw’s small model streaming fix, and NanoBot’s Termux incompatibility all point to increasing usage of on-device models, but with rough edges.

**Value for AI agent developers:**  
- Invest in **defensive security** (subprocess scrubbing, webhook authentication) early.  
- Build **channel-agnostic core** and treat platform-specific bugs as high priority.  
- Plan for **memory budget management** — users expect long-running sessions without leaks.  
- Monitor **regression detection** in CI, especially for performance and output integrity.  
- Engage with **skill ecosystem tooling** — the gap between promise and reality is a key opportunity.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-07-31

## 1. Today’s Overview

The NanoBot project saw very high development activity in the last 24 hours, with **49 pull requests updated** (16 open, 33 merged/closed) and **7 issues updated** (5 open, 2 closed). No new releases were published. The surge in merged PRs points to a focused push on stability fixes, CI improvements, and the rollout of new WebUI features. Several critical‑severity regressions from the past week have been addressed (e.g., `finish_reason='length'` misrouting, session‑lock leaks, Telegram polling stalls), and the project is actively modernizing its session storage from JSONL to SQLite. Overall project health is strong, with rapid triage of reported bugs and a healthy pipeline of community‑contributed enhancements.

## 2. Releases

No new releases were published in the last 24 hours.

## 3. Project Progress

**33 pull requests were merged or closed** in the last day. Key advances include:

- **CI/CD & Performance**: [#5145](https://github.com/HKUDS/nanobot/pull/5145) stabilised and sped up CI by replacing timing‑dependent tests with stdin‑gated handshakes and batching channel dependency installs.
- **Agent & Session Stability**:
  - [#5136](https://github.com/HKUDS/nanobot/pull/5136) fixed a critical misrouting bug when `finish_reason='length'` arrives with blank content and tool calls (closes [#5133](https://github.com/HKUDS/nanobot/issues/5133)).
  - [#5150](https://github.com/HKUDS/nanobot/pull/5150) bounded buffered session output to prevent unbounded memory growth during exec sessions.
  - [#5151](https://github.com/HKUDS/nanobot/pull/5151) released idle session locks using `WeakValueDictionary` to avoid memory leaks.
  - [#5147](https://github.com/HKUDS/nanobot/pull/5147) protected pairing approvals from being erased by transient store read failures.
  - [#5117](https://github.com/HKUDS/nanobot/pull/5117) tolerated invalid timestamps during idle compaction.
- **WebUI**: Two sequential PRs by Re‑bin introduced a persistent **Quick Chat** entry ([#5181](https://github.com/HKUDS/nanobot/pull/5181)) and later consolidated sidebar selection highlighting ([#5182](https://github.com/HKUDS/nanobot/pull/5182)). A follow‑up PR ([#5184](https://github.com/HKUDS/nanobot/pull/5184)) added **Temporary Chat** (in‑memory only) and refactored the sidebar.
- **Responses API**: [#5172](https://github.com/HKUDS/nanobot/pull/5172) preserved and replayed the opaque Responses output‑item chain (including encrypted reasoning) and compacted context between turns.
- **Session Storage**: [#5173](https://github.com/HKUDS/nanobot/pull/5173) migrated the runtime session store from JSONL files to SQLite, with automatic import and rollback backup.

## 4. Community Hot Topics

- **[#5149 – “no audio?”](https://github.com/HKUDS/nanobot/issues/5149)** (3 comments)  
  User reports that NanoBot receives audio on WhatsApp but never sends it. Logs show ffmpeg warnings. This issue has drawn community attention, but no fix PR is attached yet.

- **[#5185 – “Nanobot returning tool calls code in responses”](https://github.com/HKUDS/nanobot/issues/5185)** (1 comment)  
  A sudden regression where tool call JSON is leaked into the assistant’s text response. Author could not reproduce a specific trigger. No fix PR yet; likely related to recent changes in tool‑call handling.

- **PR [#5156 – Fix Telegram polling stalls](https://github.com/HKUDS/nanobot/pull/5156)** (opens [#5171](https://github.com/HKUDS/nanobot/issues/5171))  
  The stall bug is a high‑impact community pain point (messages lost silently). A fix PR is open, reviewing and awaiting merge.

The underlying needs point to **channel‑specific reliability** (Telegram, WhatsApp) and **response integrity** (tool call injection).

## 5. Bugs & Stability

| Severity | Issue | Summary | Fix PR exists? |
|----------|-------|---------|----------------|
| **Critical** | [#5185](https://github.com/HKUDS/nanobot/issues/5185) | Tool call code leaked into text responses – breaks all tool‑using conversations. | No |
| **Critical** | [#5171](https://github.com/HKUDS/nanobot/issues/5171) | Telegram polling permanently stalls after network blip – messages lost. | [#5156](https://github.com/HKUDS/nanobot/pull/5156) (open) |
| **High** | [#5149](https://github.com/HKUDS/nanobot/issues/5149) | WhatsApp bot cannot send audio files (only receives). | No |
| **Medium** | [#5187](https://github.com/HKUDS/nanobot/issues/5187) | NanoBot fails on Termux due to missing timezone validation. | No |
| **Low** | [#3106](https://github.com/HKUDS/nanobot/issues/3106) | Scheduled tasks with GPT fail with “couldn’t produce final answer” (works with other models). Open since April. | No |

Additional regressions have been **closed with fixes** today:
- [#5133](https://github.com/HKUDS/nanobot/issues/5133) (finish_reason length misrouting) → fixed in [#5136](https://github.com/HKUDS/nanobot/pull/5136).
- [#4791](https://github.com/HKUDS/nanobot/issues/4791) (DoS via no rate limiting) – closed but with no merged fix visible; likely deferred.

## 6. Feature Requests & Roadmap Signals

Several features are moving through the PR pipeline and are strong candidates for the next release:

- **Quick Chat & Temporary Chat** ([#5181](https://github.com/HKUDS/nanobot/pull/5181), [#5184](https://github.com/HKUDS/nanobot/pull/5184)) – first‑class sidebar entries for persistent and ephemeral conversations.
- **Custom Telegram Bot API base URL** ([#4919](https://github.com/HKUDS/nanobot/pull/4919)) – enables self‑hosted gateways.
- **Sub‑agent configurable model presets** ([#4291](https://github.com/HKUDS/nanobot/pull/4291)) – allows LLM‑specified model selection for spawned agents.
- **Session storage migration to SQLite** ([#5173](https://github.com/HKUDS/nanobot/pull/5173)) – foundational for future performance improvements.
- **Preserving OpenAI Responses reasoning state** ([#5172](https://github.com/HKUDS/nanobot/pull/5172)) – aligns with ARC‑AGI‑3 capabilities.
- **Support for well‑known skills.sh sources** ([#5186](https://github.com/HKUDS/nanobot/pull/5186)) – enhances the skills marketplace.
- **Cron manual‑run completion state** ([#5183](https://github.com/HKUDS/nanobot/pull/5183)) – fixes race conditions in scheduled execution.

User‑requested features still open as PRs include **heartbeat shared session** ([#4551](https://github.com/HKUDS/nanobot/pull/4551)) and **codex dedup fix** ([#4021](https://github.com/HKUDS/nanobot/pull/4021)), both flagged as conflicting.

## 7. User Feedback Summary

- **Pain points**:  
  - Audio sending failure on WhatsApp (issue [#5149](https://github.com/HKUDS/nanobot/issues/5149)).  
  - Sudden tool‑call leaking in responses ([#5185](https://github.com/HKUDS/nanobot/issues/5185)) – a showstopper for automation.  
  - Telegram bot going deaf after network hiccup ([#5171](https://github.com/HKUDS/nanobot/issues/5171)) – reliability concern.  
  - Incompatibility with Termux environment ([#5187](https://github.com/HKUDS/nanobot/issues/5187)).  

- **Use cases highlighted**:  
  - WhatsApp and Telegram as primary channels for both sending and receiving media.  
  - Running NanoBot on mobile devices (Termux) for experimentation.  
  - Scheduled tasks with multiple LLM providers (GPT vs. Gemini).  

- **Satisfaction signals**:  
  - Community is actively contributing fixes (33 PRs merged in one day).  
  - Quick Chat and Temporary Chat features address long‑standing requests for simplified UI.  

Overall, users are engaged but facing channel‑specific regressions; the core platform is rapidly improving.

## 8. Backlog Watch

The following items have lacked maintainer attention for an extended period:

- **[#3106](https://github.com/HKUDS/nanobot/issues/3106) – “completed tool steps but couldn’t produce a final answer”** (open since April 13, 2026). Affects GPT scheduled tasks; author notes it works with Gemini. No assignee or recent activity.

- **[#4551](https://github.com/HKUDS/nanobot/pull/4551) – feat(heartbeat): add isolated_session config** (open since June 26, 2026). Labelled conflict, hasn’t been updated in a month.

- **[#4819](https://github.com/HKUDS/nanobot/pull/4819) – fix(memory): replace WeakValueDictionary** (open since July 6, conflict label). A critical fix for lock loss, but blockers remain.

- **[#4021](https://github.com/HKUDS/nanobot/pull/4021) – fix(codex): dedup reasoning items** (open since May 27, conflict label). Addresses a 400 error in multi‑turn responses.

These items would benefit from a maintainer review and either merge or explicit closure.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-07-31

## 1. Today's Overview
Project activity remains high, with **14 issues updated** (all open) and **50 pull requests updated** (47 open, 3 merged/closed) in the last 24 hours. A new patch release **v0.19.1 (v2026.7.30)** was cut yesterday, rolling up ~1,000+ PRs into a stable tag for downstream consumers. The community is actively contributing fixes and features across plugins, memory providers, platform integrations, and the CLI, though several long-standing bugs around cost tracking, context limits, and macOS update reliability still lack maintainer resolution.

## 2. Releases
**Hermes Agent v0.19.1 (v2026.7.30)** — a patch release that bundles all PRs merged since v0.19.0. No breaking changes are documented. The release is intended to provide a stable tagged version for Docker images, hosted deployments, and fresh installs. No migration notes beyond the usual upgrade path (`hermes update` or pulling the new Docker tag) are provided.

- [Release v0.19.1](https://github.com/NousResearch/hermes-agent/releases/tag/v2026.7.30)

## 3. Project Progress
Three PRs were **merged/closed** in the last 24 hours:

- **#41868** (closed) — Fixes the Kanban hard billing/quota crash loop (bug #41805). Workers that hit a hard provider 429 before doing any tool work now exit cleanly instead of being respawned indefinitely.  
  [PR #41868](https://github.com/NousResearch/hermes-agent/pull/41868)

The remaining 47 open PRs represent active development: thinking‑tag parsing, MCP discovery timing, memory sanitisation, Discord voice fixes, and more. The project is in a productive state with a steady stream of contributions.

## 4. Community Hot Topics
The most discussed issues and their underlying needs:

- **#43900** (10 comments) — Ollama local models silently capped at 4096-token context. Users report that Hermes reads the GGUF metadata correctly but never actually sets `num_ctx`, causing garbled responses after the context limit. The community is pushing for a fix that explicitly passes the detected context length.  
  [Issue #43900](https://github.com/NousResearch/hermes-agent/issues/43900)

- **#64900** (6 comments) — Allow plugins to extend `send_message` with platform‑specific schema fields and send handlers. Currently, custom platforms require editing core code. This feature would greatly improve plugin extensibility.  
  [Issue #64900](https://github.com/NousResearch/hermes-agent/issues/64900)

- **#18304** (6 comments, 1 👍) — Usage cost always 0 in `state.db` even when tokens are recorded for Anthropic and Google. Users cannot track spending, undermining billing transparency.  
  [Issue #18304](https://github.com/NousResearch/hermes-agent/issues/18304)

- **#74973** (5 comments) — `hermes update` on macOS silently leaves the gateway dead and unloaded from `launchd` while reporting success. This causes unexpected downtime until manually restarted.  
  [Issue #74973](https://github.com/NousResearch/hermes-agent/issues/74973)

## 5. Bugs & Stability
New bugs reported in the last 24 hours (created or active):

| Issue | Severity | Description | Fix PR Exists? |
|-------|----------|-------------|----------------|
| [#75091](https://github.com/NousResearch/hermes-agent/issues/75091) | **P2** | `extra_body` from the primary provider leaks onto the fallback provider during failover because it is not re‑resolved. | No |
| [#75093](https://github.com/NousResearch/hermes-agent/issues/75093) | **P2** | JSON‑quoted YAML scalar in `mcp_servers.*.args` breaks all MCP connections — the gateway passes the JSON string verbatim as a single argument. | No |
| [#75097](https://github.com/NousResearch/hermes-agent/issues/75097) | **P2** | Iteration‑budget semantics diverge: AIAgent defaults to 90 and `execute_code` only refunds one limiter, causing premature exhaustion. | No |
| [#75088](https://github.com/NousResearch/hermes-agent/issues/75088) | **P3** | Duplicate report of npm vulnerabilities found by two independent Hermes agents (duplicate, likely already fixed). | No |
| [#75098](https://github.com/NousResearch/hermes-agent/issues/75098) | **P3** | Feature request: detect model `response_format` support during onboarding and warn user. | No |
| [#74933](https://github.com/NousResearch/hermes-agent/issues/74933) | **P3** | Hindsight provider rejects `shared` observation scope, fragmenting observations per session. | Yes: [#75092](https://github.com/NousResearch/hermes-agent/pull/75092) |
| [#75096](https://github.com/NousResearch/hermes-agent/issues/75096) | **P2** | Telegram conflict‑retry backoff is reset after transient progress — fix PR submitted. | Yes: [#75096](https://github.com/NousResearch/hermes-agent/pull/75096) |
| [#75081](https://github.com/NousResearch/hermes-agent/issues/75081) | **P2** | Gateway advance‑snapshot logic duplicates transcript writes in FTS‑rebuild retry path. Fix PR submitted. | Yes: [#75081](https://github.com/NousResearch/hermes-agent/pull/75081) |
| [#75094](https://github.com/NousResearch/hermes-agent/issues/75094) | **P1** | Hardened thinking‑tag regex to handle model‑prefixed variants (e.g. `<M2_7think>`) that providers like MiniMax emit. | Yes: [#75094](https://github.com/NousResearch/hermes-agent/pull/75094) |
| [#75095](https://github.com/NousResearch/hermes-agent/issues/75095) | **P3** | SQL injection‑like vulnerability in FTS5 MATCH queries for memory search — fix submitted. | Yes: [#75095](https://github.com/NousResearch/hermes-agent/pull/75095) |

Several of today’s new bugs already have corresponding pull requests, indicating a reactive and fast‑fixing community.

## 6. Feature Requests & Roadmap Signals
Notable feature requests with traction:

- **Plugin extensibility for `send_message`** (#64900) – would allow plugin‑platforms to add custom parameters (voice selection, metadata) without core edits. Likely candidate for v0.20.
- **ElevenLabs generation options** (#39382) – exposes `language_code`, `voice_settings`, and speaker boost. Already has a PR, likely to land soon.
- **`register_aux_provider` for plugins** (#70691) – enables plugins to contribute auxiliary providers to the hardcoded chain. Strong roadmap signal for plugin ecosystem growth.
- **Unified subagent model/reasoning controls** (#74375) – adds CLI and TUI model selection for subagents. Indicates movement toward richer delegation management.
- **OpenAI Responses verbosity** (#72638) – first‑class `text_verbosity` config for GPT‑5. Reflects alignment with evolving provider APIs.

Prediction: The next minor release (v0.20) will likely include plugin send_message extensibility, ElevenLabs TTS options, and subagent model controls, given the maturity of their PRs.

## 7. User Feedback Summary
Real pain points surfacing from today’s data:

- **Ollama users** are frustrated by silent 4K context caps that force manual `num_ctx` overrides (issue #43900).
- **macOS users** encounter silent update failures that leave the gateway dead, requiring manual recovery (#74973).
- **Multi‑provider users** are hit by cost tracking being broken across Anthropic, Google, and likely others (#18304).
- **Hindsight memory users** face observation fragmentation because shared scopes are not honoured (#74933).
- **MCP users** are bitten by YAML parsing quirks that break all connections (#75093).
- **Kanban/usage‑limit users** were stuck in infinite crash loops until today’s fix (#41868 merged).

On the positive side, the community is actively contributing PRs for most of these issues, and the release frequency (v0.19.1 just out) suggests maintainers are responsive to regressions.

## 8. Backlog Watch
Issues that have been open for weeks or months without clear resolution:

- **#18304** (created May 1, 2026) – Usage cost always 0. Despite 6 comments and 1 👍, no PR or maintainer assignment. Users cannot trust billing data.  
  [Issue #18304](https://github.com/NousResearch/hermes-agent/issues/18304)

- **#26770** (created May 16, 2026) – `hermes update` causes OOM on low‑memory servers. Filed 2.5 months ago with 1 comment and no PR.  
  [Issue #26770](https://github.com/NousResearch/hermes-agent/issues/26770)

- **#28796** (created May 19, 2026) – Built-in memory and external provider both inject into system prompt when external provider has content. No comment from maintainers.  
  [Issue #28796](https://github.com/NousResearch/hermes-agent/issues/28796)

- **#29667** (created May 21, 2026) – WeCom (企业微信) silent delivery failures due to ephemeral WebSocket disconnects. Marked P2 but no progress.  
  [Issue #29667](https://github.com/NousResearch/hermes-agent/issues/29667)

- **#43900** (created June 11, 2026) – Ollama context cap. 10 comments, high community demand, but no maintainer response.

These issues represent risk of user dissatisfaction and churn if left unaddressed. The Kanban crash (#41805) was fixed today after being open since June 8, showing that backlog items can eventually be resolved — but a more systematic triage for these oldest open bugs would improve project health.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest – 2026-07-31

---

## 1. Today’s Overview

The project saw moderate activity over the past 24 hours, with **7 issues updated** (4 open, 3 closed) and **17 pull requests updated** (12 open, 5 merged/closed). No new releases were published. The bulk of PR activity consists of automated dependency bumps (e.g., AWS SDK, Anthropic SDK, Pion RTP, GitHub Actions) alongside three human-authored PRs that fix bugs or add features. A notable newly opened issue (#3308) flags concurrency hazards and memory/speed optimisations across core components, indicating on-going quality hardening. Overall project health appears stable, with steady contributions from the community and maintainers.

---

## 2. Releases

**None.**  
No releases were announced in the last 24 hours.

---

## 3. Project Progress

Five pull requests were merged or closed today:

- **#3163** `[CLOSED]` – **feat(bedrock): leverage Converse prompt caching via cache points** (merged)  
  Adds support for AWS Bedrock’s prompt caching feature (cache point markers in `system`, `tools`, `messages`). This reduces cost for repeated conversation prefixes.  
  [PR #3163](https://github.com/sipeed/picoclaw/pull/3163)

- **#3262** `[CLOSED]` – **build(deps): bump actions/setup-go from 6 to 7**  
  [PR #3262](https://github.com/sipeed/picoclaw/pull/3262)

- **#3263** `[CLOSED]` – **build(deps): bump actions/setup-node from 6 to 7**  
  [PR #3263](https://github.com/sipeed/picoclaw/pull/3263)

- **#3288** `[CLOSED]` – **build(deps): bump** `github.com/aws/aws-sdk-go-v2/service/bedrockruntime` **from 1.53.3 to 1.56.0**  
  [PR #3288](https://github.com/sipeed/picoclaw/pull/3288)

- **#3290** `[CLOSED]` – **build(deps): bump** `github.com/aws/aws-sdk-go-v2/config` **from 1.32.25 to 1.32.31**  
  [PR #3290](https://github.com/sipeed/picoclaw/pull/3290)

Additionally, three issues were closed: #2546 (OAuth 2.1 + PKCE feature, see *Feature Requests*), #3258 (process hook bug fix), and #3257 (stateless gateway mode feature).  
Closure of #3257 indicates progress toward a requested stateless mode for gateway sessions.

---

## 4. Community Hot Topics

The most active issue continues to be **#2546** (closed today), which proposed **OAuth 2.1 + PKCE for MCP servers** from the dashboard. With **6 comments** and a duplicate issue (#3302) filed just yesterday, it clearly taps into strong user demand for a non‑technical MCP integration workflow similar to Claude.ai’s “Add connector”.  
[Issue #2546](https://github.com/sipeed/picoclaw/issues/2546)  
[Issue #3302](https://github.com/sipeed/picoclaw/issues/3302)

Several other issues received 2 comments each:
- **#3287** – Better support for long messages over IRCv3  
  [Issue #3287](https://github.com/sipeed/picoclaw/issues/3287)
- **#3258** – Process Hook bug (closed)  
  [Issue #3258](https://github.com/sipeed/picoclaw/issues/3258)
- **#3257** – Stateless/no‑history gateway mode (closed)  
  [Issue #3257](https://github.com/sipeed/picoclaw/issues/3257)

These discussions reveal an engaged community focused on channel parity (IRC, Telegram) and configuration flexibility (OAuth, session management).

---

## 5. Bugs & Stability

**High severity:**  
- **#3308** – *[BUG] Code Review: Concurrency hazards, goroutine leaks, and memory/speed optimizations in SeaHorse, Channel Manager, and Hooks*  
  A comprehensive report identifying potential goroutine leaks and optimisation opportunities across three core packages. No fix PR exists yet; this is a critical quality concern.  
  [Issue #3308](https://github.com/sipeed/picoclaw/issues/3308)

**Fixed bugs:**  
- **#3258** – *Process Hook before_tool modify not working* – **closed** after a 2‑comment discussion. The defect (decision field discarded, args misparsed) appears resolved.  
  [Issue #3258](https://github.com/sipeed/picoclaw/issues/3258)

**Bug fix PRs in flight:**  
- **#3279** – `fix(seahorse): prevent tool‑call format leakage into LLM summaries` – addresses a class of bug where tool‑call formatting leaks into user‑visible summaries.  
  [PR #3279](https://github.com/sipeed/picoclaw/pull/3279)

- **#3283** – `fix(dingtalk): support picture/image message inbound` – adds inbound image handling for DingTalk channel (with graceful fallback).  
  [PR #3283](https://github.com/sipeed/picoclaw/pull/3283)

---

## 6. Feature Requests & Roadmap Signals

Multiple feature issues were opened or discussed:

- **OAuth 2.1 for MCP servers** – #2546 (closed) and duplicate #3302. The implementation has likely been accepted; the duplicate suggests high demand may accelerate delivery in the next release.
- **Session listing/switching for Telegram** – #3307: a parity gap vs. the Web UI. Users on Telegram cannot list or switch conversation sessions. This is a clear next‑step for channel capability.
- **Better IRC long message handling** – #3287: IRCv3 message splitting confuses the agent. Community wants a cohesive handling approach.
- **Stateless gateway mode** – #3257 (closed): the feature was accepted and implemented.

**PRs signalling near‑term features:**  
- **#3270** – `feat: add DashScope TTS provider and WeChat audio file sending` – introduces Alibaba Cloud TTS and audio support for WeChat. Likely to be merged in the coming days.  
  [PR #3270](https://github.com/sipeed/picoclaw/pull/3270)
- **#3271** – `chore(providers): update default model names to 2026-07 latest` – refreshes model IDs for nine providers (OpenAI GPT‑5.6, Anthropic Claude 4.5, etc.).  
  [PR #3271](https://github.com/sipeed/picoclaw/pull/3271)
- **#3200** – `feat(models): add configurable default fallback chain` – allows users to configure model fallback chains via Web UI. Open since July 1.  
  [PR #3200](https://github.com/sipeed/picoclaw/pull/3200)
- **#3222** – `refactor(deltachat): cleanup implementation, documentation -200LOC` – reduces codebase size and improves clarity for the Delta Chat channel.  
  [PR #3222](https://github.com/sipeed/picoclaw/pull/3222)

**Prediction:** The next minor version will likely include OAuth 2.1 support for MCP, session management for Telegram, the DashScope TTS provider, and the model fallback chain UI.

---

## 7. User Feedback Summary

**Pain points expressed:**

- **Channel parity:** Telegram users lack session management commands (#3307); IRC users experience message splitting (#3287).
- **OAuth friction:** Non‑technical users cannot add OAuth‑protected MCP servers without a CLI or Node.js. The “Add connector” UX from Claude.ai is a strong reference point (#2546, #3302).
- **Stateless gateway use‑case:** Users running PicoClaw via gateway mode want a simple way to create fresh conversations without session persistence (#3257).
- **Process hook bug:** The `before_tool` modify hook was silently discarding `decision` field changes (#3258, now fixed).

**Satisfaction signals:**  
The number of open issues and PRs (28 total updated in 24 hours) suggests a healthy, active community. The quick closure of #2546 and #3257 shows maintainer responsiveness to high‑value requests.

---

## 8. Backlog Watch

The following items have been marked **stale** and have received no recent attention from maintainers:

- **PR #3222** – `refactor(deltachat): cleanup implementation, documentation -200LOC`  
  Open since July 3, stale label. No maintainer comments. This is a significant code reduction (‑200 LOC) that could improve Delta Chat channel maintainability.  
  [PR #3222](https://github.com/sipeed/picoclaw/pull/3222)

- **Issue #3287** – *Better support long messages in IRC*  
  Open since July 22, stale label with 2 comments. No assignee. Still waiting for a design decision or implementation.  
  [Issue #3287](https://github.com/sipeed/picoclaw/issues/3287)

- **PR #3291** – `build(deps): bump github.com/github/copilot-sdk/go from 0.2.0 to 1.0.8`  
  A major version jump (0.2 → 1.0.8) with breaking changes. Should be reviewed carefully. No maintainer activity in 7 days.  
  [PR #3291](https://github.com/sipeed/picoclaw/pull/3291)

- **PR #3289** – `build(deps): bump github.com/pion/rtp from 1.10.2 to 1.10.5`  
  Simple patch bump, but also marked stale. Low risk; could be merged quickly.  
  [PR #3289](https://github.com/sipeed/picoclaw/pull/3289)

Additionally, the high‑severity bug report **#3308** (concurrency hazards) is only hours old but will require maintainer attention soon.

---

*Digest generated from GitHub data for sipeed/picoclaw on 2026-07-31.*

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-07-31

## Today's Overview

Activity is extremely high: 19 pull requests were updated in the last 24 hours (12 open, 7 merged/closed), alongside 2 newly filed issues — a clear spike that points to a coordinated push by multiple contributors. No new release was published today, so the mainline tip remains the primary delivery channel. The open issues describe two distinct reliability regressions (one affecting all inbound message reactions/edits, another breaking skill installation from registry branches), both surfaced by core or regular contributors, which suggests the project is hitting integration friction after recent structural changes. The PR pipeline shows a strong mix of critical bugfixes, image hardening, and skill ecosystem additions, with several long-running PRs suddenly gaining updates — possibly indicating a maintainer-driven merge window or a preparation for an upcoming release.

---

## Releases

**No new releases** were published today. The next tagged version remains unknown; the pace of merged PRs (7 today) suggests a release could be imminent.

---

## Project Progress

**7 pull requests were merged or closed in the last 24 hours**, covering bugfixes, infrastructure tightening, and documentation improvements:

- **#3160** (merged) — Repinned the agent image to `hardened-2026-07-30`, reducing layer count from 18 to 8 and shrinking the largest single layer from 39% to 27% of total size. This accelerates image pulls across the fleet.
- **#3159** (merged) — Made the Vercel CLI opt-in (`/add-vercel` skill) rather than baked into every agent image. Reduces credential surface area and trims ~20 MB from default images.
- **#3122** (merged) — Fixed `opencode` skill compatibility with main branch, custom-endpoint transport logic, and memory parity. A previously broken integration is now functional again.
- **#2682** (merged) — Added v2 compatibility detection in `update-skills`: branches with `package.json` version starting with `1.` are now skipped gracefully rather than offered for installation. Prevents operator confusion.
- **#3152** (merged) — Documentation only: added links to `docs/REQUIREMENTS.md` and `docs/SECURITY.md` from the README Architecture section, improving discoverability.
- **#2476** (merged) — Implemented a `/restart no-nanoclaw` skill for restarting services without touching the NanoClaw core process.
- **#3014** (merged) — Fixed `hasIdenticalSend` in `agent-runner` to bound deduplication checks to the current turn, preventing false positives across turns.

---

## Community Hot Topics

**Issue #3153** (1 comment) — `add_reaction` / `edit_message` always fail on inbound messages because the agent-group suffix is not stripped from the platform message ID. On Slack every call returns `message_not_found` and then retries 3× before failing. This is a fresh report (opened 2026-07-30) with high potential impact: it blocks all interactive message reactions and message editing from the agent side. No fix PR exists yet.

**Issue #3155** (0 comments) — Registry branches have drifted from `main`; provider payloads fail their own install gates. The reporter ran a skill installation step from a specific revision and hit a typecheck failure, indicating that the skill payload in the `providers` branch is not compatible with current `main`. This suggests a gap in CI/CD for keeping registry branches in sync after core changes.

**PR #3119** — Reconciles untracked orphan containers that cause duplicate-per-group spawns. The root cause (observed over 5 days of uptime) is that a single agent group accumulated 3 concurrent containers due to a scheduler race. This has been open since July 23 and was updated today, indicating ongoing review.

**PR #2685** — Updates Signal integration docs for group typing indicators, outbound reactions, and quote-reply fix. This PR has been open since June 4 — one of the longest-open documentation improvements, now updated again.

---

## Bugs & Stability

### High Severity

- **#3153 — `add_reaction` / `edit_message` always fail on inbound messages**  
  Platform message IDs retain an agent-group suffix, so the platform never finds the message. Affects all users relying on reactions or message editing via the agent. No fix PR yet.  
  *Link: nanocoai/nanoclaw#3153*

- **#3155 — Registry branches drifted from main; provider payloads fail install gates**  
  A skill installation from the `providers` branch fails at its own typecheck step on current `main`. This is a CI/registry integrity bug that breaks the `/add-codex` workflow and could affect any operator installing skills from registry branches.  
  *Link: nanocoai/nanoclaw#3155*

### Medium Severity

- **#3119 — Duplicate per-group container spawns**  
  Untracked orphan containers accumulate when the scheduler fires multiple times before reconciliation completes. A single agent group reached 3 concurrent containers. A fix PR (#3119) is open and under review.  
  *Link: nanocoai/nanoclaw#3119*

### Low Severity / Niche

- **#3157 — Dangling symlinks crash template skill materialization**  
  `fs.statSync` follows symlinks that point at container-internal paths, causing file-not-found errors when materializing template skills in non-container contexts. Fix PR is open.  
  *Link: nanocoai/nanoclaw#3157*

---

## Feature Requests & Roadmap Signals

The following PRs indicate features currently in development or review:

- **#3154 — Scheduled tasks now receive `current_time`** (core-team)  
  Tasks are rendered with their effective scheduled occurrence time, and a task-only `current_time` including weekday is generated when the task reaches the agent. This improves support for time-aware agent behavior.

- **#3158 — Image verification with publisher identity pinning** (core-team)  
  Wires real Sigstore publisher identity into the verification gate that previously was skipped due to missing env vars. Ensures all agent image pulls are attestation-checked per architecture.

- **#3124 — Report unavailable MCP servers**  
  Enables graceful detection and reporting of MCP servers that are down or unresponsive, rather than silent failures.

- **#3145 — Backfill destinations for existing wirings**  
  Migration 021 adds missing channel destinations for existing messaging-group wirings, preserving custom local names. Likely to be merged soon.

- **Long-running skill additions** (PRs #2301, #2317, #2634) continue to receive updates, indicating sustained interest in GitHub polling mode, voice transcription (Whisper), and AWS credential proxy (paws4claws). These have been open for 2-3 months but show no signs of abandonment.

---

## User Feedback Summary

**Pain Points:**
- **Message reactions/edits broken** — The most immediate user-facing bug: any agent attempting to react to or edit an inbound message will always fail. This blocks common conversational patterns and could degrade user trust in agent responses.
- **Skill installation failures** — Registry branch drift means that users following the documented `/add-codex` workflow may hit build errors. This creates a poor onboarding experience for operators trying to extend their agents.
- **Container duplication** — Operators running agents with `NRestarts=0` over multiple days can silently accumulate duplicate containers, consuming extra resources without visible errors.

**Satisfaction Signals:**
- The rapid merge of 7 PRs in a single day suggests strong maintainer responsiveness and contributor momentum.
- Image hardening (PR #3160) and the Vercel CLI opt-in (PR #3159) were both reviewed and merged quickly — these address real operational overhead and security concerns that fleet operators would feel.

**Use Cases:**
- The Signal integration PR (#2685) shows active demand for group chat, reactions, and wire-format documentation — users are pushing for parity with Slack/Telegram features.
- Multiple new skill PRs (polling GitHub, local Whisper, AWS credential proxy) indicate a community that is actively building agent capabilities for internal deployments, not just consuming pre-built features.

---

## Backlog Watch

The following Issues and PRs have been open for extended periods without maintainer response or merge, and may need attention:

- **Issue #2685 — Signal docs PR** (opened June 4, 2026, 57 days open)  
  A documentation-only PR that has been superseded by later changes? Updated today but still not merged. May only need a final review.  
  *Link: nanocoai/nanoclaw#2685*

- **PR #2301 — GitHub polling mode (+ git access + safe secret merge)** (opened May 6, 2026, 86 days open)  
  One of the most feature-rich skill additions (polling mode for NAT environments, webhook security warnings, safe OneCLI secret merge). Despite high utility, it has been languishing.  
  *Link: nanocoai/nanoclaw#2301*

- **PR #2317 —/add-voice-transcription-free-whisper** (opened May 7, 2026, 85 days open)  
  A complete skill with two backends (openai-whisper, whisper.cpp) and pre-flight detection. Potentially blocked by infrastructure decisions.  
  *Link: nanocoai/nanoclaw#2317*

- **PR #2634 — `add-paws4claws` AWS credential proxy** (opened May 28, 2026, 64 days open)  
  Introduces a new mount-from-outside pattern. Could be a test balloon for a more general operator skill architecture.

- **PR #2537 — Pre-commit hooks CI** (opened May 18, 2026, 73 days open)  
  Comprehensive pre-commit configuration for prettier, eslint, typecheck, vitest. Would improve contributor quality of life. May be deprioritized vs. feature work.

---

*Digest generated from GitHub data snapshot for 2026-07-31. 19 PRs updated, 0 releases, 2 new issues.*

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest – 2026-07-31

## Today’s Overview
IronClaw saw heavy activity in the last 24 hours: 19 issues updated (all open) and 50 pull requests updated (27 open, 23 merged/closed). No new releases were cut. The project is in an intense restructuring phase, with multiple workstreams executing the “target crate architecture” plan (epic #3773). At the same time, several critical bugs and user-reported pain points surfaced, including a cross-user memory leak (#6900) and a Slack integration failure (#6834). The team is actively shipping large feature PRs for durable cross-channel attachments (#6364), hosted MCP server registration (#6930), and a product command train for Slack and WebUI (#6931, #6891).

## Releases
No new releases this period.

## Project Progress (Merged/Closed PRs Today)
Of the 50 PRs updated, 23 were merged or closed. Notable closed PRs include:

- **#6934** – `refactor(host_api): de-wildcard the contract prelude` (WS0 of target architecture) – merged, behavior-free restructuring.
- **#6931** – `feat(slack): native /ironclaw slash commands` – merged, final PR of the command train.
- **#6891** – `feat(webui): role-filtered command palette` – merged, PR-2 of command train.
- **#6862** – `fix(reborn): preserve terminal model error explanations` – merged, with a DB migration.
- **#6874** – `chore(deps): bump everything-else group` – closed (likely superseded by #6932).

These PRs advanced features (Slack commands, WebUI palette, error recovery) and laid the foundation for the ongoing architecture overhaul.

## Community Hot Topics
The most active issues by comment count:

- **#6284** – [EPIC] Error-recoverability endgame (15 comments)  
  The team is iterating on a formal recoverability contract so that *every* mid-run error is visible to the model and actionable. This is a foundational reliability goal.

- **#6524** – Epic: Hermetic capability and journey testing platform (4 comments)  
  A push toward deterministic, mechanical coverage for every supported capability and user journey.

- **#6752** – Bug: Instance deletion fails, UI stuck on re-login (1 comment)  
  Real-world user block reported via Slack; has high visibility.

- **#6900** – Bug: Shared-channel memory leak (new, 0 comments yet)  
  Critical security/privacy issue – users in shared Slack channels leak into each other’s memory namespace.

Pull requests with high activity (large changes, core contributors):

- **#6364** – Durable cross-channel file flows (XL, 31 commits)
- **#6930** – Register hosted MCP servers (XL)
- **#6901** – Agentic activity and streaming UX redesign (XL, new contributor)
- **#6780** – Deep-link register/install gateway for IronHub (XL, re-port)

## Bugs & Stability
Several bugs reported or updated today, ranked by severity:

| Issue | Severity | Description | Fix PR |
|-------|----------|-------------|--------|
| #6900 | **Critical** | Cross-user memory leak in shared channels (Slack). All users collapse into the operator’s memory namespace. | Not yet |
| #6752 | **High** | Instance deletion fails, UI stuck on “Loading your agents…” on re-login. | Not yet |
| #6834 | **High** | Slack integration setup fails for near.foundation accounts. | Not yet |
| #6866 | **Medium** | Same home directory shared across all users; workspaces visible to others (privacy). | Not yet |
| #6916 | **Low** | Markdown `.md`/`.mdx` files rendered as plain text in preview modal. | Not yet |
| #6915 | **Low** | Workspace file links in assistant messages do not open the referenced file. | Not yet |

No fix PRs have been opened yet for these bugs. The team appears focused on structural refactoring and feature work.

## Feature Requests & Roadmap Signals
User requests and roadmap epics updated today:

- **#6901** (PR) – Agentic activity & streaming UX redesign: a new contributor is building a `NearProcessIndicator` with streaming content, based on an approved design mockup. Likely target for next release.
- **#6565** – Epic: Reliable Skill Discovery, Routing, and Activation. A corrected diagnosis from last week shows the auto-activation pipeline is not running on the primary `TurnCoordinator` path. This is a high-risk, high-impact feature.
- **#6284** – Error recoverability endgame: a top-level goal to make the model recover from 100% of errors. Multiple sub-issues expected.
- **#3773** – Target crate architecture epic: a multi-wave program to restructure the entire crate layout. Workstreams #6920–#6927 were created today to implement the plan.

These epics indicate that the next major release will focus on reliability, skill selection, and architectural clarity rather than new user-facing features.

## User Feedback Summary
Real pain points and satisfaction signals from Slack and issue comments:

- **Instance management frustration**: Deleting an instance leaves the UI in a broken state (#6752). Users must refresh or re-login blindly.
- **Slack integration unreliability**: Setup flow fails without clear error (#6834). Shared channels also cause memory leaks (#6900).
- **Privacy concern**: All users see each other’s workspaces due to a shared home directory (#6866). Enterprise users will be affected.
- **Minor UX annoyances**: Markdown files not rendering (#6916) and file links not navigating (#6915) degrade the assistant experience.

Overall, users are encountering blocking stability and security issues that may erode trust, but the steady stream of feature PRs (Slack commands, streaming UX) indicates active development to address the product’s maturity.

## Backlog Watch
Items that have been open for a significant time without resolution:

| Issue/PR | Days Open | Notes |
|----------|-----------|-------|
| #3773 – Epic: Target Crate Architecture | 73 days (since May 19) | Now being executed with daily PRs. |
| #5598 – Release PR | 28 days (since July 3) | Open, blocking new version; breaking changes for `ironclaw_common` and `ironclaw_skills` are ready. |
| #6428 – Dependabot (tokio-ecosystem) | 10 days | Low risk, but stale updates not merged. |
| #5664 – Dependabot (actions group) | 26 days | Many action updates pending; maintainer intervention needed. |

The release PR #5598 is especially notable – its merge would cut a new version with API breaking changes. Its prolonged open state suggests the team is waiting for more changes to land before shipping.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest – 2026-07-31

## 1. Today's Overview
The project showed healthy development activity despite zero Issue updates in the last 24 hours. Ten Pull Requests were updated, with eight merged or closed, indicating steady feature delivery and maintenance. The recent release (2026.7.29) introduced selected-text tagging for side chat, support for Kimi K3 models, and auth hardening. The two open PRs are older items that received fresh attention, suggesting the team is revisiting backlog. Overall project momentum remains strong, with a focus on cowork features, enterprise isolation, and UI/UX polish.

## 2. Releases
**New Release:** [LobsterAI 2026.7.29](https://github.com/netease-youdao/LobsterAI/releases)

Changes included:
- `feat(cowork)`: add selected text tags to side chat
- `feat`: support Kimi K3 model
- `fix(auth)`: harden session lifecycle and token refresh

No breaking changes or migration notes were documented in the release.

## 3. Project Progress – Merged/Closed PRs
Eight PRs were merged/closed in the last 24 hours (all closed as of 2026-07-30). Key advances:

- **[#2412 – fix(nsis): re-kill survivor processes on every stop poll round](https://github.com/netease-youdao/LobsterAI/pull/2412)** (Windows) – Improves installer shutdown reliability by repeatedly killing survivors during polling.
- **[#2411 – feat(sidebar): support check-in and banner carousel](https://github.com/netease-youdao/LobsterAI/pull/2411)** – Adds unified carousel for daily check-in and image banners in the sidebar.
- **[#2410 – style(sites): align page layout with management views](https://github.com/netease-youdao/LobsterAI/pull/2410)** – UI consistency across Sites, Skills, and MCP pages.
- **[#2409 – feat(enterprise): isolate account-scoped auth and service flows](https://github.com/netease-youdao/LobsterAI/pull/2409)** – Prevents cross-account contamination, enforces entitlements, and improves diagnostics.
- **[#2408 – feat(activity): add native daily check-in experience](https://github.com/netease-youdao/LobsterAI/pull/2408)** – Desktop daily check-in with rewards flow, integrated with sidebar and account menu.
- **[#2406 – fix(cowork): improve side chat input handling](https://github.com/netease-youdao/LobsterAI/pull/2406)** – Accumulates selected text excerpts, removes length limits, keeps bounded context.
- **[#2397 – feat(cowork): add isolated /btw side chat](https://github.com/netease-youdao/LobsterAI/pull/2397)** – Floating side-chat panel with drag/resize, isolated from main conversation, routed via OpenClaw.
- **[#2389 – fix(email): prevent attachment path traversal](https://github.com/netease-youdao/LobsterAI/pull/2389)** – Security fix: sanitizes filenames and enforces download boundaries, with cross-platform tests.

## 4. Community Hot Topics
No Issues were updated, and all PRs listed show zero comments and zero reactions, indicating low direct community engagement in this period. However, two older open PRs (see Backlog Watch) received activity, suggesting maintainers are addressing long-standing requests. The underlying need for better conversation management surfaced in the merged side‑chat features and the open “mark as unread” PR (#1228).

## 5. Bugs & Stability
No new bugs were filed in the last 24 hours (0 Issues). Two bug-fix PRs were merged that address stability and security:

- **Attachment path traversal** ([#2389](https://github.com/netease-youdao/LobsterAI/pull/2389)) – Security vulnerability fix; severity moderate.
- **NSIS survivor process kill** ([#2412](https://github.com/netease-youdao/LobsterAI/pull/2412)) – Installer reliability fix; low severity.

The release also included auth hardening. No crashes or regressions were reported.

## 6. Feature Requests & Roadmap Signals
Two open PRs represent user‑facing feature requests that have been pending since April:

- **Mark session as unread** ([#1228](https://github.com/netease-youdao/LobsterAI/pull/1228)) – Adds manual unread status to sessions.
- **AgentCreateModal Escape close & form reset** ([#1231](https://github.com/netease-youdao/LobsterAI/pull/1231)) – Improves modal UX.

Given the recent focus on cowork and side‑chat enhancements, the “mark as unread” feature is likely to be considered for the next minor release. The modal UX fix may also be prioritized as a quality‑of‑life improvement.

## 7. User Feedback Summary
No direct user feedback (comments, reactions, or new Issues) was recorded in this period. Indirect signals from merged features suggest users desire:
- Isolated side conversations (`/btw` side chat)
- Daily check‑in gamification
- Enterprise multi‑account safety

No dissatisfaction or pain points were explicitly voiced.

## 8. Backlog Watch
Two open PRs require maintainer attention:

- **[#1228 – feat(cowork): 新增会话「标记为未读」功能](https://github.com/netease-youdao/LobsterAI/pull/1228)** (Created 2026-04-01, updated 2026-07-30) – Adds “mark session as unread” to context menus and Redux store. Stale for ~4 months; a decision to merge or close is needed.
- **[#1231 – fix(agent): AgentCreateModal 支持 Escape 键关闭](https://github.com/netease-youdao/LobsterAI/pull/1231)** (Created 2026-04-01, updated 2026-07-30) – Modal UX improvements. Same staleness pattern.

Both were updated recently, possibly indicating renewed review. The project would benefit from clear resolution (merge, request changes, or close with explanation).

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest – 2026-07-31

## 1. Today's Overview
The Moltis project shows a steady development cadence with **5 pull requests updated in the last 24 hours**, of which **2 were merged/closed**. There are **2 open issues** updated today, one a feature request and one a security-critical bug report. No new releases were cut today. The project remains actively maintained by core contributor **penso** alongside community contributions from **eddyvlad** (issue) and **Practice100101** (bug report) and **Jonesxq** (PR). Activity is focused on strengthening messaging integrations (Slack, Telegram), adding observability infrastructure, and hardening access control.

## 2. Releases
**No new releases** were published today. The latest public release remains the previous version (no information provided). Users should monitor upcoming releases incorporating the merged PRs and bug fixes described below.

## 3. Project Progress
Two pull requests were **merged/closed** in the last 24 hours:

- **PR #1166** – *feat(slack): per-message acknowledgment reactions, phases, reconnect supervision, and Block Kit*  
  Merged today by **penso**. This builds on earlier acknowledgement work (#1165) to provide safe lifecycle handling for Slack reactions (the equivalent of a typing indicator). It addresses queueing, cancellation, retries, callback bursts, and delivery failures. [View PR](https://github.com/moltis-org/moltis/pull/1166)

- **PR #1169** – *feat(acp): expose Moltis as an ACP agent over stdio*  
  Closed/merged (likely earlier) but still updated recently. Adds a default-on `moltis acp` command that exposes Moltis as an ACP agent with prompt routing, session isolation, concurrency limits, and final-text reconciliation. [View PR](https://github.com/moltis-org/moltis/pull/1169)

Three **open PRs** remain in active review:

- **PR #1170** – *fix(channels): gate /sh and privileged tools behind a per-account operators list*  
  Addresses a privilege escalation risk by introducing an explicit `operators` list per account. [View PR](https://github.com/moltis-org/moltis/pull/1170)

- **PR #1174** – *Add instrumentation and feedback collection infrastructure*  
  Adds backend-neutral agent instrumentation, Langfuse v4 export, OTLP backends, and end-user reaction feedback. A major step toward observability. [View PR](https://github.com/moltis-org/moltis/pull/1174)

- **PR #1176** – *feat(web): add Markdown copy and session export*  
  Preserves Markdown in copied assistant replies and adds a **Save as Markdown** action that exports full session history with images. [View PR](https://github.com/moltis-org/moltis/pull/1176)

## 4. Community Hot Topics
Neither of the two open issues nor any PRs have comments or reactions recorded, indicating early-stage discussion. The absence of community engagement on these items may suggest they were filed recently (both created yesterday) and may see more activity in coming days.

- **Issue #1178** – *[Feature]: Let agents send Telegram inline buttons and receive structured callback responses*  
  No comments yet. Filed by **eddyvlad**, this request would allow agents to create interactive Telegram keyboards and handle structured callbacks. [View Issue](https://github.com/moltis-org/moltis/issues/1178)

- **Issue #1177** – *[Bug]: Vault Unlock/Recovery Endpoints Missing Authentication (CWE-306)*  
  Filed by **Practice100101**. No discussion yet, but the security implications are potentially severe (see Bugs & Stability). [View Issue](https://github.com/moltis-org/moltis/issues/1177)

## 5. Bugs & Stability
One **high-severity bug** was reported today:

- **Issue #1177** – *Vault Unlock/Recovery Endpoints Missing Authentication (CWE-306)*  
  **Severity: Critical** – Unauthenticated access to vault unlock/recovery endpoints could allow an attacker to bypass credential checks. No fix PR exists yet. The issue is brand new (created 2026-07-30) and requires immediate maintainer attention to assess impact and develop a hotfix. [View Issue](https://github.com/moltis-org/moltis/issues/1177)

No other crashes or regressions were reported today.

## 6. Feature Requests & Roadmap Signals
Two feature-oriented items stand out:

- **Telegram inline buttons & structured callbacks (Issue #1178)** – This would extend Moltis’s interactive messaging capabilities to Telegram, mirroring the Slack Block Kit work merged earlier (PR #1166). Likely to be considered for the next minor release given the recent Slack work.

- **Markdown copy & session export (PR #1176)** – Already submitted as a pull request by **Jonesxq**, this is likely to be reviewed and merged soon. It enhances the web UI for power users who want to preserve session context.

- **Instrumentation & feedback (PR #1174)** – If merged, this would add production-grade observability (OpenTelemetry, Langfuse) and user feedback collection. This signals a roadmap priority toward enterprise readiness and usage analytics.

## 7. User Feedback Summary
No explicit user satisfaction/dissatisfaction data was collected in the last 24 hours. However, the bug report (Issue #1177) and feature request (Issue #1178) imply:

- **Pain point**: Security concerns around vault endpoints – a user discovered an authentication gap and proactively reported it, indicating awareness of Moltis’s security posture.
- **Desired capability**: Agents interacting with Telegram using rich inline keyboards – a user would like to build bots that present structured options and process callbacks, likely for workflows like ticket selection, form filling, or navigation.

## 8. Backlog Watch
No issues or PRs appear to have been languishing without maintainer attention. All open items were created or updated within the last 2–5 days. The busiest contributor **penso** continues to push and review code actively. No stale items were identified. Maintainers should keep an eye on **Issue #1177** (security bug) as it may require a fast-tracked fix, potentially bypassing normal release cycles.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# QwenPaw Project Digest — 2026-07-31

## Today’s Overview
The QwenPaw project saw very high activity: **15 issues** and **50 pull requests** were updated in the last 24 hours, with **26 PRs closed/merged** and **11 open issues** remaining active. No new releases were published. The community is heavily focused on bug fixes and regressions introduced by v2.0, particularly a ~2 s fixed overhead per reply, MCP session recovery failures, and UI freezes caused by large shell outputs. Several contributors stepped up to fix CI blockers, Matrix E2EE, and sandbox handling, indicating strong external engagement despite the influx of stability concerns.

## Releases
No new releases were published during this period. The latest release remains v2.0.0.post3 (as referenced in issue #6307).

## Project Progress
26 PRs were merged or closed today. Notable advancements include:

- **Fix Matrix E2EE** ([#6486](https://github.com/agentscope-ai/QwenPaw/pull/6486)) — Probes the modern `vodozemac` backend instead of legacy `olm`, restoring encrypted messaging on Python 3.12.
- **Fix CI block for fork PRs** ([#6563](https://github.com/agentscope-ai/QwenPaw/issues/6563) → PR [#6562](https://github.com/agentscope-ai/QwenPaw/pull/6562) merged) — Resolves the `real-behavior-proof` workflow failure that prevented all fork contributors from passing CI.
- **Computer Use native automation** ([#6424](https://github.com/agentscope-ai/QwenPaw/pull/6424) merged) — Adds a `computer_use` tool with accessibility-first GUI automation for Windows and macOS.
- **Creator plugin iteration** ([#6556](https://github.com/agentscope-ai/QwenPaw/pull/6556) merged) — Introduces creation checkpoints, home redesign, media recovery, and export/import.
- **Sandbox cleanup** ([#6582](https://github.com/agentscope-ai/QwenPaw/pull/6582) merged) — Fixes sandbox handling in cleanup logic.
- **Sandbox fallback configurability** ([#6256](https://github.com/agentscope-ai/QwenPaw/pull/6256) merged) — Makes the fallback action when sandbox is unavailable configurable.
- **Session-level approval inheritance** ([#6506](https://github.com/agentscope-ai/QwenPaw/issues/6506) closed) — Fixed in PR [#6562](https://github.com/agentscope-ai/QwenPaw/pull/6562); child sub-agents now properly respect the parent’s `approval_level = OFF`.

## Community Hot Topics
The most active issues by comment count:

- **[#6307 – v2.0 introduces ~2s fixed overhead per reply](https://github.com/agentscope-ai/QwenPaw/issues/6307)** (7 comments, open)  
  *Underlying need:* Users upgrading from v1.x experience a significant performance regression unrelated to model latency. This is a top-priority blocker for adoption.

- **[#6524 – MCP backend restart breaks client until manual reconnect](https://github.com/agentscope-ai/QwenPaw/issues/6524)** (5 comments, open)  
  *Underlying need:* MCP connections become stale after server restarts; the client does not automatically recover sessions, requiring the user to run `list mcp` to re-establish.

- **[#6563 – CI bug blocks all fork PRs](https://github.com/agentscope-ai/QwenPaw/issues/6563)** (4 comments, closed)  
  *Underlying need:* Contributing from forks was impossible due to a CI workflow permission issue. Quickly resolved by the community.

The most commented PRs (all with undefined comment count in the extract) include the multi‑fix PR [#6562](https://github.com/agentscope-ai/QwenPaw/pull/6562) (closed, fixing 3 bugs) and the Computer Use PR [#6424](https://github.com/agentscope-ai/QwenPaw/pull/6424) (merged).

## Bugs & Stability
Several bugs were reported or addressed today, ranked by severity:

| Severity | Issue | Description | Fix PR exists? |
|----------|-------|-------------|----------------|
| **Critical** | [#6307](https://github.com/agentscope-ai/QwenPaw/issues/6307) | ~2s fixed overhead on every reply in v2.0 | No |
| **High** | [#6589](https://github.com/agentscope-ai/QwenPaw/issues/6589) | `execute_shell_command` large output freezes UI | No |
| **High** | [#6524](https://github.com/agentscope-ai/QwenPaw/issues/6524) | MCP sessions not recovered after server restart | ✅ PR [#6586](https://github.com/agentscope-ai/QwenPaw/pull/6586) |
| **Medium** | [#6555](https://github.com/agentscope-ai/QwenPaw/issues/6555) | Dream memory compression misses early-session events when context scrolls out | No |
| **Medium** | [#6578](https://github.com/agentscope-ai/QwenPaw/issues/6578) | Cron `dispatch.mode: "final"` not honored; all events pushed live | No |
| **Medium** | [#6588](https://github.com/agentscope-ai/QwenPaw/issues/6588) | `spawn_subagent` single-task mode unusable because `batch` is required in schema | No |
| **Low** | [#6476](https://github.com/agentscope-ai/QwenPaw/issues/6476) | Matrix E2EE broken on Python 3.12 | ✅ PR [#6486](https://github.com/agentscope-ai/QwenPaw/pull/6486) (merged) |
| **Low** | [#6506](https://github.com/agentscope-ai/QwenPaw/issues/6506) | Session-level approval not inherited by sub-agents | ✅ In PR [#6562](https://github.com/agentscope-ai/QwenPaw/pull/6562) (merged) |

The performance regression [#6307](https://github.com/agentscope-ai/QwenPaw/issues/6307) remains unaddressed and is the most impactful stability issue.

## Feature Requests & Roadmap Signals
Several user-requested features emerged, mostly around UX polish and tool behavior:

- **Chinese filename preservation** ([#6453](https://github.com/agentscope-ai/QwenPaw/issues/6453)) — Two complementary PRs exist ([#6567](https://github.com/agentscope-ai/QwenPaw/pull/6567), [#6492](https://github.com/agentscope-ai/QwenPaw/pull/6492)), with #6567 already submitted. Likely to land in next patch.
- **Large output handling** ([#6512](https://github.com/agentscope-ai/QwenPaw/issues/6512)) — Suggests automatic file output or streaming for `execute_shell_command`. No PR yet; may be folded into the UB freeze fix.
- **Multimodal warning optimization** ([#6452](https://github.com/agentscope-ai/QwenPaw/issues/6452)) — Users find the “no multimodal capability” banner too aggressive.
- **Character count toggle** ([#6585](https://github.com/agentscope-ai/QwenPaw/issues/6585)) — Request to hide the live character count display while loading.
- **File drag-and-drop display** ([#6583](https://github.com/agentscope-ai/QwenPaw/issues/6583)) — Show multiple filenames on separate lines instead of one long line.
- **Desktop app rename** ([#6587](https://github.com/agentscope-ai/QwenPaw/issues/6587)) — Drop the “Desktop” suffix.
- **Configurable theme system** (PR [#6312](https://github.com/agentscope-ai/QwenPaw/pull/6312) draft) — Still in draft; likely for v2.1.
- **Unified provider model platform** (PR [#6302](https://github.com/agentscope-ai/QwenPaw/pull/6302) open) — Large refactor addressing multiple model-provider pain points; may target next minor release.

Features likely to appear in the next version (v2.0.2 or v2.1): Chinese filename hints, MCP session auto-recovery, sandbox fallback configurability, and a fix for the approval inheritance.

## User Feedback Summary
User sentiment is mixed — strong appreciation for the project’s capabilities (e.g., “非常不错的项目” in [#6585](https://github.com/agentscope-ai/QwenPaw/issues/6585)) but clear frustration with regressions. Key pain points:

- **Performance regression** from v1.x to v2.0 (issue [#6307](https://github.com/agentscope-ai/QwenPaw/issues/6307)) is the most critically felt.
- **MCP connectivity issues** ([#6524](https://github.com/agentscope-ai/QwenPaw/issues/6524)) disrupt workflows where remote MCP servers restart.
- **UI freezes** during large shell command output ([#6589](https://github.com/agentscope-ai/QwenPaw/issues/6589)) make the desktop app unusable in those scenarios.
- **Small UX annoyances** (file name display, character count flicker, app name) accumulate and lower satisfaction.

Several users volunteered to contribute PRs (e.g., [#6512](https://github.com/agentscope-ai/QwenPaw/issues/6512) “Willing to Contribute” checkbox checked), indicating a motivated community willing to help fix pain points.

## Backlog Watch
Issues and PRs that have seen maintainer inaction or remain unresolved:

| Item | Description | Days since last activity | Danger |
|------|-------------|--------------------------|--------|
| [#6307](https://github.com/agentscope-ai/QwenPaw/issues/6307) | Performance regression ~2s overhead | 10 days (created Jul 21, updated Jul 30) | **High** — no fix PR; blocks v2.0 adoption |
| [#6555](https://github.com/agentscope-ai/QwenPaw/issues/6555) | Dream memory compression misses early events | 2 days, no PR | Medium — complex memory bug |
| [#6512](https://github.com/agentscope-ai/QwenPaw/issues/6512) | `execute_shell_command` output truncation | 3 days, no PR | Medium — related to UI freeze issue |
| [#6589](https://github.com/agentscope-ai/QwenPaw/issues/6589) | UI freeze on large output | 1 day, no PR | High severity but recent |
| [#6588](https://github.com/agentscope-ai/QwenPaw/issues/6588) | `spawn_subagent` batch schema bug | 1 day, no PR | Medium |
| [#6586](https://github.com/agentscope-ai/QwenPaw/pull/6586) | MCP session recovery PR | 1 day, open, needs review | Medium — maybe awaiting merge |
| [#6567](https://github.com/agentscope-ai/QwenPaw/pull/6567) | Chinese filename fix PR | 1 day, open, needs review | Low — clear fix |
| [#6302](https://github.com/agentscope-ai/QwenPaw/pull/6302) | Unified provider model platform PR | 10 days, open | Large refactor, needs careful maintainer review |
| [#6312](https://github.com/agentscope-ai/QwenPaw/pull/6312) | Configurable theme draft PR | 10 days, open | Draft, may need scope definition |

The most critical backlog item is the performance regression (#6307), which lacks any associated fix PR. Maintainers should prioritize triaging this issue and providing guidance to the community.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

# ZeptoClaw Project Digest – 2026-07-31

Generated from GitHub activity (github.com/qhkm/zeptoclaw)

---

## 1. Today's Overview

The ZeptoClaw repository shows minimal activity over the last 24 hours: no new issues, no merged pull requests, and no releases. The only item in flight is an open pull request (#645) that has been updated but not yet merged. This PR addresses two significant runtime concerns—credential leakage and incomplete process cleanup—indicating that maintainers are actively working on hardening the shell‑execution subsystem. Overall project velocity is low today, but the single open PR suggests focused maintenance effort on security and stability.

## 2. Releases

No new releases are available. The latest release remains unchanged.

## 3. Project Progress

No pull requests were merged or closed in the last 24 hours.  
The only active PR is **#645** (open, last updated 2026-07-30):

> **fix(runtime): scrub subprocess secrets and reap timed-out process trees**  
> *Author: qhkm*  
> *URL: [PR #645](https://github.com/qhkm/zeptoclaw/pull/645)*

This PR aims to fix two issues:
- Environment variable leakage: model‑authored shell commands were inheriting ZeptoClaw’s full process environment, exposing provider keys and other secrets.
- Process‑tree reaping: runtime timeouts that dropped `Command::output()` futures left zombie descendants, and Docker containers were not consistently terminated.

While not yet merged, this work addresses critical operational risks.

## 4. Community Hot Topics

The only active discussion item is **PR #645**. Though it has zero comments and reactions in the snippet, its summary reveals the underlying need for **secure isolation of subprocesses** (preventing credential leaks to untrusted model commands) and **reliable process lifecycle management** (avoiding zombie processes and container leaks). These are foundational trust and reliability concerns for any AI agent runtime. No other issues or PRs have generated community discussion in the past 24 hours.

## 5. Bugs & Stability

No new bugs or regressions were reported today.  
The open PR #645 directly targets the following stability/security issues (severity: **high**):

| Issue | Severity | Fix PR Exists? |
|-------|----------|----------------|
| Credential leakage to subprocesses (provider keys, environment secrets) | High | Yes (PR #645) |
| Zombie processes / orphaned containers after timeout | Medium‑High | Yes (PR #645) |

No crashes or regressions have been noted.

## 6. Feature Requests & Roadmap Signals

No user‑requested features were filed in the last 24 hours. The activity on PR #645 suggests that the maintainers are prioritising **security hardening** and **resource cleanup** for the next release. If merged, these changes will likely appear in the upcoming version (e.g., 0.x.0). No roadmap‑level signals are visible beyond this PR.

## 7. User Feedback Summary

No explicit user feedback (comments, reviews, or issue reports) was recorded today. However, the nature of PR #645 implies user pain points:

- **Pain point 1:** Users running untrusted model‑authored commands risk exposing environment secrets (e.g., API keys, cloud credentials). This is a critical trust issue.
- **Pain point 2:** Long‑running or timed‑out commands can leave behind zombie processes and dangling containers, wasting resources and potentially blocking future executions.

These are operational rather than feature‑related complaints, and the fix directly addresses them.

## 8. Backlog Watch

There are no long‑unanswered issues or PRs requiring maintainer attention. The only open PR (#645) is recently updated and being actively worked on by the core maintainer (`qhkm`). No items are flagged as stale or neglected.

---

*End of digest – data as of 2026-07-31 23:59 UTC*

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-07-31

## 1. Today's Overview

The project is in a period of **intense, high-risk development** with **50 pull requests updated in the last 24 hours** (all still open) and **two critical security bugs filed** by contributor JordanTheJet. No releases were cut today. Activity is clearly centered on patching immediate attack-surface vulnerabilities and advancing major feature work (eval pipelines, skill injection, plugin validation). The **two new issues** both carry significant severity labels (S0 data-loss/security risk and S2 degraded behavior), and **fix PRs already exist for both** (#9569, #9568), indicating a quick-response cycle. The sheer volume of open PRs (50) and the presence of “needs-author-action” labels on many suggest that maintainer bandwidth remains a bottleneck, but the pipeline is active.

## 2. Releases

No new releases were created on 2026-07-31.

## 3. Project Progress

**No pull requests were merged or closed today.** The 50 PRs updated were all open at the time of the data snapshot. However, several **new fix PRs were introduced** that directly address the bugs reported today:

- **#9569** – `fix(gateway): fail closed when a WhatsApp Cloud or Linq webhook cannot be verified` (fixes #9565)
- **#9568** – `fix(security): match command allowlist entries case-insensitively on Unix` (fixes #9566)
- **#9571** – `chore(channels): remove the WATI channel` (cleanup after vulnerability exposure)

Other meaningful feature-boundary PRs updated today include:

- **#9567** – Multiple To/Cc/Bcc recipients in email channel (stacked on #9506)
- **#9570** (implied by #9567) – email reply-threading fix
- **#9126** – Plugin typed instance config validation (large, high risk)
- **#9248** – Append-only eval run-history receipts
- **#9244** – Seed and grade isolated case memory in evaluation
- **#9224** – Repeated live eval runs with pass@k and error bars

## 4. Community Hot Topics

The **most active issue** today is:

- **[#9565 – [Bug]: gateway webhook handlers do not fail closed](https://github.com/zeroclaw-labs/zeroclaw/issues/9565)**  
  *Comments: 2, Author: JordanTheJet, Severity: S0*  
  This is a **red-flag security issue**: three inbound webhook handlers in the gateway dispatch attacker-controlled messages without authenticating the caller. The discussion likely centers on verification gaps in WhatsApp Cloud, Linq, and the now-being-removed WATI channel. The **fix PR #9569** was opened the same day, showing rapid triage.

- **[#9566 – [Bug]: uppercase allowed_commands entries never match on Unix](https://github.com/zeroclaw-labs/zeroclaw/issues/9566)**  
  *Comments: 0, Severity: S2*  
  Root cause: a case-sensitive comparison gated to Windows only, leaving Unix users unable to enforce allowlists with uppercase characters. **Fix PR #9568** provides a case-insensitive match across all platforms.

The **PRs with highest risk labels** (but few comments) include:

- [#9325](https://github.com/zeroclaw-labs/zeroclaw/pull/9325) – fix streamed user turns for small local models (risk: high, size: M)
- [#8937](https://github.com/zeroclaw-labs/zeroclaw/pull/8937) – stream-hash tool args to avoid deep clone (risk: high)
- [#8688](https://github.com/zeroclaw-labs/zeroclaw/pull/8688) – trusted goal tools and delegation boundaries (risk: high, size: XL)

These indicate the community is pushing on **correctness for local models** and **security boundaries for delegated execution**.

## 5. Bugs & Stability

**Two bugs filed today**, both with fix PRs already in place:

| # | Title | Severity | Status | Fix PR |
|---|-------|----------|--------|--------|
| [#9565](https://github.com/zeroclaw-labs/zeroclaw/issues/9565) | Gateway webhook handlers do not fail closed | **S0 – data loss / security risk** | Open | [#9569](https://github.com/zeroclaw-labs/zeroclaw/pull/9569) |
| [#9566](https://github.com/zeroclaw-labs/zeroclaw/issues/9566) | Uppercase allowed_commands entries never match on Unix | **S2 – degraded behavior** | Open | [#9568](https://github.com/zeroclaw-labs/zeroclaw/pull/9568) |

Both were authored by **JordanTheJet**, indicating a security-conscious contributor. The S0 bug is particularly concerning because it allows **unauthenticated injection of attacker-controlled messages** into the agent pipeline. The **WATI channel is being removed entirely** (PR #9571) as part of the remediation, which suggests the vulnerability may have been systemic to that channel’s implementation.

Additionally, an older **regression** is noted in #9566 (regressed from #4552), meaning a previously fixed case-insensitivity behavior was lost.

## 6. Feature Requests & Roadmap Signals

While no explicit new feature *requests* were filed today (only bugs), the following **enhancement PRs** advanced and signal upcoming capabilities:

- **Evaluation framework maturation** – PRs #9248, #9244, #9224, #9225 are building a robust eval suite with repeat runs, memory seeding, regression tests, and history receipts. This aligns with a likely **v2.0 evaluation module**.
- **Plugin instance config validation** – PR #9126 introduces JSON Schema-based validation for plugin configs, which will be essential for security and stability.
- **Skill injection defaults** – PR #8313 (open since June 25) moves to compact injection by default, deprecating full-mode. This should reduce prompt context consumption.
- **Trusted goal / delegation boundaries** – PR #8688 adds `goal_start`, `goal_objective`, `goal_resume` tools with scoped trust – a major architectural change for multi-agent scenarios.

Predictions for the **next minor release**:
- Fix for #9565 (gateway fail-closed) – mandatory
- Fix for #9566 (allowlist case-insensitive) – high priority
- Removal of WATI channel (#9571)
- Possibly the email multiple-recipient feature (#9567) if #9506 stabilizes

## 7. User Feedback Summary

Real user pain points surfaced today:

1. **Security / trust** – A user (or security researcher) found that webhook handlers silently accept unauthenticated messages. This is a **major trust concern** for anyone using WhatsApp Cloud, Linq, or (until removal) WATI channels. The fix is inbound, but the perception of vulnerability may linger.
2. **Command allowlist silently failing** – Users who configured `allowed_commands` with uppercase characters (e.g., `Write-Output` on Unix) would find those commands silently denied. This is a **configuration breakage** that could cause hours of debugging.
3. **Small local model usability** – PR #9325 addresses a specific complaint: models like Ollama’s `llama3.2` treat streamed user turns as log output, leading to protocol commentary instead of conversation. This reflects a growing segment of users running local models who expect seamless chat behavior.

No explicit satisfaction signals or praise were present in the data, but the quick response with fix PRs suggests the project is responsive to critical reports.

## 8. Backlog Watch

The following important items need maintainer attention:

- **[PR #8688](https://github.com/zeroclaw-labs/zeroclaw/pull/8688) – feat(runtime): add trusted goal tools and delegation boundaries** (opened July 4, needs-author-action)  
  A large (XL) enhancement touching many subsystems. It has been open nearly a month without merging. If this is critical for the next release, it needs renewed review.

- **[PR #8313](https://github.com/zeroclaw-labs/zeroclaw/pull/8313) – feat(skills): default to compact injection, deprecate full mode** (opened June 25, needs-author-action)  
  A behavior-changing feature that has been sitting for over a month. The deprecation window may be causing user confusion.

- **[#9565](https://github.com/zeroclaw-labs/zeroclaw/issues/9565) – Gateway security bug** (opened today)  
  Although PR #9569 exists, the issue itself is a **red flag for production users**. Maintainers should prioritize merging the fix and issuing a patch release.

- **[PR #8937](https://github.com/zeroclaw-labs/zeroclaw/pull/8937) – fix(agent): stream-hash tool args** (opened July 10, needs-author-action)  
  A performance fix for deep-clone overhead in loop detection. Marked risk:high – a review could prevent a future regression.

No issues older than one month with zero comments were identified, but the “needs-author-action” label on many PRs suggests the bottleneck is on the contributor side, not the maintainers. The project health appears strong in terms of activity, but **security fixes must be expedited** to maintain trust.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*