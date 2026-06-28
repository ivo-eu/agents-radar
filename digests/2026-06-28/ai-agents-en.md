# OpenClaw Ecosystem Digest 2026-06-28

> Issues: 150 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-28 10:09 UTC

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

# OpenClaw Project Digest – 2026-06-28

## Today's Overview

OpenClaw shows extremely high activity with **150 issues** and **500 pull requests** updated in the last 24 hours. Of those, 140 issues remain open and 365 PRs are still in progress, indicating a fast-moving project with substantial maintainer workload. No new releases were cut today, but the community and development teams are actively addressing a broad range of bugs, features, and regressions. The project health is robust but strained by the volume of high-severity issues, particularly around session-state integrity, message delivery, and security boundaries.

## Releases

No new releases on 2026-06-28.

## Project Progress

**135 pull requests were merged or closed today**, reflecting significant forward momentum. Notable closures include:

- **[#76171](openclaw/openclaw/issues/76171)** (P1, impact:crash-loop) – Stale worker process accumulation causing high host load; now closed (presumably fixed).
- **[#97216](openclaw/openclaw/issues/97216)** (P3, WebChat sidebar) – Agent display name feature request closed.
- **[#97454](openclaw/openclaw/issues/97454)** (fix provider-transport) – SSE body handling under JSON content-type, merged.
- **[#97460](openclaw/openclaw/issues/97460)** (fix CLI) – Skip gateway secrets.resolve for exec-based commands, merged.
- **[#97461](openclaw/openclaw/issues/97461)** (fix Matrix) – Deduplicate poll answers before max_selections, merged.
- **[#68936](openclaw/openclaw/issues/68936)** (autofix pipeline + Windows daemon) – Scripts PR closed.

A major ongoing refactor stands out:

- **[#96625](openclaw/openclaw/issues/96625)** (size: XL, proof: sufficient, status: ⏳ waiting on author) – Flips sessions and transcripts to SQLite storage. This is a foundational change that could significantly improve reliability and query performance, but remains in author-review state.

Other long-running PRs continue to see updates, such as the Mattermost security fix **[#64546](openclaw/openclaw/issues/64546)** and the gateway memory flush fix **[#88968](openclaw/openclaw/issues/88968)**.

## Community Hot Topics

The most active issues (by comment count) reveal deep-seated concerns about session reliability, security, and missing features:

| Issue | Comments | Underlying Need |
|-------|----------|-----------------|
| [#69208](openclaw/openclaw/issues/69208) Umbrella: duplicate transcript, replay, and context assembly across channels | 12 | Cross-channel consistency and message integrity. Users want **no duplicate or lost messages** regardless of runtime path. |
| [#67777](openclaw/openclaw/issues/67777) Subagent completion delivery lost on timeout/drain/orphan | 10 | Reliable **subagent orchestration** – completion must survive transient failures. |
| [#71736](openclaw/openclaw/issues/71736) Control UI plugin contribution slots (RFC) | 9 | **Plugin extensibility** for chat modes, approval cards, input guards – requested by 100yenadmin. |
| [#72015](openclaw/openclaw/issues/72015) Active-memory blocks replies, overloads multi-agent gateways | 8 | **Stability under load** – active-memory plugin causing slowdowns/reliability issues. |
| [#76171](openclaw/openclaw/issues/76171) Stale worker processes → high load (now closed) | 7 | **Process lifecycle management** – worker leaks causing multi-minute response times. |
| [#71142](openclaw/openclaw/issues/71142) Configurable upload size limit | 7 | **Usability** – hardcoded 5MB limit blocks legitimate larger uploads. |
| [#73182](openclaw/openclaw/issues/73182) Reasoning default silently flipped to on for Claude | 6 | **Cost and privacy** – silent config change doubled Anthropic spend and leaked thinking tokens. |
| [#72031](openclaw/openclaw/issues/72031) image tool fails for Bedrock with AWS SDK creds | 6 | **Provider compatibility** – Bedrock auth regression blocks image understanding. |
| [#74077](openclaw/openclaw/issues/74077) Slash command for preview streaming mode | 5 | **Runtime configuration** – users want per-chat streaming control without restart. |
| [#73602](openclaw/openclaw/issues/73602) WhatsApp/Telegram reliability on WSL2 | 5 | **Platform-specific channel reliability** – WSL2 users experiencing persistent disconnects. |

Top PRs by activity (based on size, labels, and maintainer engagement):

- **[#96625](openclaw/openclaw/issues/96625)** – SQLite sessions/transcripts refactor (XL, proof sufficient, waiting on author)
- **[#88968](openclaw/openclaw/issues/88968)** – Prevent memory flush failure from aborting user reply (P1, ready for maintainer look)
- **[#97175](openclaw/openclaw/issues/97175)** – Bound deferred turn maintenance with per-task timeout (P1, ready for maintainer look)
- **[#96229](openclaw/openclaw/issues/96229)** – Per-agent env subprocess contract (XL, proof sufficient, ready for maintainer look)
- **[#97340](openclaw/openclaw/issues/97340)** – MSTeams multi-account support (XL, waiting on author)

## Bugs & Stability

Today’s updates surface several high-severity bugs, many with fix PRs in flight. Sorted by impact:

### Critical / High Severity (P1)

- **[#69208](openclaw/openclaw/issues/69208)** (P1) – Umbrella for duplicate transcript/replay across channels. **No fix PR yet.** 
- **[#67777](openclaw/openclaw/issues/67777)** (P1) – Subagent completion delivery lost on timeout/drain/orphan. **Linked PR open.**
- **[#72015](openclaw/openclaw/issues/72015)** (P1) – Active-memory blocks replies, overloads multi-agent gateways. **Linked PR open.**
- **[#73182](openclaw/openclaw/issues/73182)** (P1) – Reasoning default silently flipped on for Claude, doubling spend. **No new fix PR; linked PR open.**
- **[#72031](openclaw/openclaw/issues/72031)** (P1) – `image` tool fails for Bedrock with AWS SDK auth. **Linked PR open.**
- **[#76171](openclaw/openclaw/issues/76171)** (P1) – Stale worker process accumulation (now **closed** – fixed).
- **[#71699](openclaw/openclaw/issues/71699)** (P1) – Windows hard crash 0xC0000409 (stack buffer overrun) during Mattermost streaming. **Needs live repro.**
- **[#72418](openclaw/openclaw/issues/72418)** (P1) – `shouldSkipBackendSelfPairing` allows loopback clients to bypass device pairing (CVSS 9.3). **Linked PR open.**
- **[#69943](openclaw/openclaw/issues/69943)** (P1) – Session-memory hook re-injects raw tokens, creating poisoning loop. **Needs product decision.**
- **[#73910](openclaw/openclaw/issues/73910)** (P1) – OpenClaw-managed Codex ACP uses isolated CODEX_HOME without auth bridge. **Linked PR open.**
- **[#71066](openclaw/openclaw/issues/71066)** (P1) – Telegram polling silently non-functional. **Needs live repro.**
- **[#73801](openclaw/openclaw/issues/73801)** (P1) – Active Memory with Cerebras gpt-oss-120b times out and can pin CPU. **Needs live repro.**
- **[#71326](openclaw/openclaw/issues/71326)** (P1) – Cross-exec stale file reads (vnode/dentry cache race) – regression in 2026.4.20. **Needs info.**
- **[#72021](openclaw/openclaw/issues/72021)** (P2) – Short-term memory promotion mixes daily/session signals with real recalls, skewing recall quality.

### Medium Severity (P2)

- **[#73602](openclaw/openclaw/issues/73602)** – WhatsApp flaps and Telegram stalls on WSL2. **Needs live repro.**
- **[#69572](openclaw/openclaw/issues/69572)** – Feishu typing indicator uses wrong API (Message Reaction instead of Typing). **Needs live repro.**
- **[#72704](openclaw/openclaw/issues/72704)** – Excessive JSON metadata in Telegram messages degrades model comprehension. **Linked PR open.**
- **[#71417](openclaw/openclaw/issues/71417)** – `openclaw agent` defaults to last channel, silently resumes old session. **Linked PR open.**
- **[#73432](openclaw/openclaw/issues/73432)** – QMD embedding never triggered per configured interval. **Linked PR open.**
- **[#72176](openclaw/openclaw/issues/72176)** – Intermittent duplicate message delivery in 2026.4.24 across all channels. **Linked PR open.**
- **[#73574](openclaw/openclaw/issues/73574)** – Wiki lint does not resolve Obsidian link variants. **Linked PR open.**
- **[#69582](openclaw/openclaw/issues/69582)** – Parameter injection `{}` causes infinite loop in tool invocation. **Linked PR open.**
- **[#73049](openclaw/openclaw/issues/73049)** – memory-core breaks agent by using wrong API key for embeddings (regression 2026.4.25). **Needs live repro.**

### Low / Cosmetic

- **[#73743](openclaw/openclaw/issues/73743)** (P1?) – CLI subprocess startup ~25s on idle macOS; worker-tick fan-out compounds to 5+ min ticks. **Linked PR open.**
- **[#73676](openclaw/openclaw/issues/73676)** (P2) – CLI crashes when current working directory is deleted (uv_cwd error not handled). **Linked PR open.**
- **[#73478](openclaw/openclaw/issues/73478)** (P2) – Gateway does not output image info (transcript missing). **Linked PR open.**

**Key stability signal:** The majority of P1 bugs now have linked fix PRs, and one major P1 (worker accumulation) is already closed. However, many issues remain in `needs-live-repro` or `needs-product-decision` state, indicating the project is still triaging root causes.

## Feature Requests & Roadmap Signals

Several high-quality feature requests have emerged, many with detailed RFCs and proof-of-concept PRs. Likely candidates for the next release:

- **[#71736](openclaw/openclaw/issues/71736)** – Control UI plugin contribution slots (RFC, P2, needs product decision). If adopted, this would create a **standardized SDK surface** for plugins to extend chat modes, approval cards, input guards, etc. Expected to land after the session/transcripts refactor.
- **[#71142](openclaw/openclaw/issues/71142)** – Configurable upload size limit (P2, enhancement). Simple config change likely to ship soon given community demand.
- **[#74077](openclaw/openclaw/issues/74077)** – Slash command `/stream` for preview streaming mode (P3). Low risk, high usability – could appear in a point release.
- **[#71195](openclaw/openclaw/issues/71195)** – macOS Talk Mode speech-to-speech (P2, enhancement, linked PR open). Parity with voice-call plugin – sub-second turns. High value for mobile/desktop users.
- **[#71712](openclaw/openclaw/issues/71712)** – Agent-facing scheduling API with non-forgeable provenance (RFC, P2). Enables agents to manage cron jobs securely. Would unlock powerful autonomous workflows.
- **[#82450](openclaw/openclaw/issues/82450)** – Linear Persistent Workspace Mode for blind users (P2, accessibility). Includes linear layout, persistent workspace, no scrolling. Strong community support.
- **[#73537](openclaw/openclaw/issues/73537)** – Production-readiness stability label for releases (P2). Users want clear guidance on which releases are safe for production.
- **[#73638](openclaw/openclaw/issues/73638)** – Non-interactive onboarding for trusted-proxy auth (P2, feature). Essential for automated deployments behind reverse proxies.
- **[#72717](openclaw/openclaw/issues/72717)** – SQLite FTS index for wiki_search (P2, enhancement). Performance boost for large wikis.
- **[#71301](openclaw/openclaw/issues/71301)** – Version-matched bundled docs + native docs retrieval (P3). Improves agent autonomy.

**Prediction:** Next version (likely 2026.5.x or 2026.6.x) will include the SQLite sessions refactor (#96625), which will unblock many other improvements. Configurable upload size and the `/stream` command are low-effort high-reward and could land as hotfixes. The Control UI plugin slots are more architecturally significant and likely scheduled for a subsequent minor release.

## User Feedback Summary

**Positive signals:**
- A fully blind user (#82450) calls OpenClaw "one of the most powerful AI work interfaces I have ever used," using it daily for video workflows, browser automation, social media, and music market research.
- A family/business user (#73537) thanks the team and says OpenClaw has "genuinely become part of our daily workflow" with Telegram integration, automations, and Home Assistant control.
- Multiple users show appreciation for the project's ambition, especially around multi-agent orchestration and plugin extensibility.

**Pain points and frustrations:**
- **Performance degradation:** Several users report 2-3+ minute response times (#76171, #72015) and high CPU load (#73801), especially with Active Memory plugin or on multi-agent gateways.
- **Unreliable message delivery:** Duplicate messages (#69208, #72176), lost subagent completions (#67777), and silent polling failures (#71066) erode trust in channel reliability.
- **Security concerns:** Silent config changes (#73182), auth bypass vectors (#72418, #64546), and data poisoning (#69943) worry users deploying in production.
- **Platform-specific issues:** Windows crashes (#71699), WSL2 flakiness (#73602), macOS CLI slowness (#73743) – users on non-Linux environments experience disproportionate bugs.
- **Missing configuration knobs:** Hardcoded upload limits (#71142), no per-chat streaming control (#74077), no non-interactive onboarding (#73638) – users want more operational flexibility.
- **Regression anxiety:** Multiple regressions reported (e.g., Bedrock image tool #72031, duplicate messages in 2026.4.24 #72176, memory-core API key issue in 2026.4.25 #73049). Users are cautious about upgrading without clear changelogs.

## Backlog Watch

Several high-importance items are languishing or need maintainer attention:

### Issues Stalled Without Resolution

- **[#45655](openclaw/openclaw/issues/45655)** (P2, open since Mar 14) – Poe image/video-output models accepted in config but fail at runtime. No fix PR yet. **Important for media-workspace users.**
- **[#69572](openclaw/openclaw/issues/69572)** (P2, open since Apr 21) – Feishu typing indicator uses wrong API. Needs product decision. **Blocks Feishu user experience.**
- **[#68264](openclaw/openclaw/issues/68264)** (P2, open since Apr 17) – Canvas/Browser UI visualization fails to render. Regression. **Waiting on maintainer review.**
- **[#68187](openclaw/openclaw/issues/68187)** (P2, open since Apr 17) – SSE-backed MCP sessions go stale after server restart. **Needs maintainer review.**
- **[#67750](openclaw/openclaw/issues/67750)** (P2, open since Apr 16) – Successful auto-compaction can still timeout and cause `/new` fallback. **Linked PR open but needs maintainer review.**
- **[#73049](openclaw/openclaw/issues/73049)** (P2, open since Apr 27) – memory-core regression 2026.4.25: wrong API key used for embeddings. **Needs live repro.**

### PRs Awaiting Maintainer Action

- **[#96625](openclaw/openclaw/issues/96625)** (XL, SQLite refactor) – **Waiting on author** since Jun 25. This is the largest architectural change in flight and may be blocking other work.
- **[#88968](openclaw/openclaw/issues/88968)** (P1, memory flush fix) – **Ready for maintainer look** but has been open since Jun 1. Needs maintainer sign-off.
- **[#97175](openclaw/openclaw/issues/97175)** (P1, deferred turn maintenance timeout) – **Ready for maintainer look** since Jun 27. Could solve a class of session hangs.
- **[#96229](openclaw/openclaw/issues/96229)** (XL, per-agent env subprocess contract) – **Ready for maintainer look** since Jun 24. Long-running PR with security implications.
- **[#64546](openclaw/openclaw/issues/64546)** (P1, Mattermost HMAC security) – **Waiting on author** since Apr 11. Severity is high (forgeable interaction tokens) but author may have stalled.
- **[#73751](openclaw/openclaw/issues/73751)** (P2, Windows exec codepage fix) – **Waiting on author** since Apr 28. Blocks CJK users on Windows.
- **[#73403](openclaw/openclaw/issues/73403)** (P1, agent routing per recipient) – **Waiting on author** since Apr 28. Fixes a fundamental session-key bug.

**Recommendation:** The maintainers should prioritize reviewing the P1 ready-for-look PRs (#88968, #97175, #96229) as they directly address user-facing reliability issues (memory flush abort, deferred turn hangs, agent environment contracts). The Mattermost security PR (#64546) should be escalated due to its criticality. The stalled SQLite refactor (#96625) may need maintainer guidance to unblock the author.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report – 2026-06-28

## 1. Ecosystem Overview

The open-source personal AI assistant ecosystem remains highly fragmented but intensely active, with projects ranging from monolithic reference implementations (OpenClaw) to ultra-lightweight agents (NanoBot) and enterprise-grade multi-tenant systems (IronClaw). A clear convergence is emerging around channel-matrix reliability, session-state integrity, and local-model compatibility – reflecting user demand for production-grade dependability across diverse deployment environments. The landscape is characterized by rapid feature iteration, with several projects merging 10–50+ PRs per day, while others consolidate after major refactors.

## 2. Activity Comparison

| Project | Issues Updated (24h) | PRs Updated (24h) | Release Today? | Health Score (1-5) |
|---|---|---|---|---|
| **OpenClaw** | 150 (140 open) | 500 (365 open) | No | 4 – high throughput, but P1 bug backlog stresses maintainers |
| **NanoBot** | 4 (3 open) | 24 (19 open) | No | 4 – fast bug fixes, active feature pipeline |
| **Hermes Agent** | 10 (9 open) | 50 (34 open) | No | 5 – high activity, rapid merges, strong contributor engagement |
| **PicoClaw** | 3 (1 open) | 3 (2 open) | No | 3 – stable but low momentum; one open bug unaddressed |
| **NanoClaw** | 1 (1 open) | 7 (7 open) | No | 2 – no merges today; single critical bug; 3 PRs stalled 8 days |
| **NullClaw** | 1 (1 open) | 1 (1 open) | No | 2 – very low activity; Android build issue stale 2 months |
| **IronClaw** | 9 (2 open) | 50 (26 open) | No | 5 – extremely productive; capability-policy epic complete; OAuth fix merged |
| **LobsterAI** | 4 (0 open, all closed) | 6 (2 open) | No | 3 – moderate; merged long-pending artifact preview; compatibility gap unresolved |
| **TinyClaw** | 0 | 0 | - | 1 – no activity observed |
| **Moltis** | 1 (1 open) | 4 (4 open) | No | 2 – steady but no merges; PRs target local-model quirks |
| **CoPaw** | 2 (2 open) | 6 (6 open) | No | 2 – no PRs merged; high-severity regression; test PRs awaiting review |
| **ZeptoClaw** | 0 | 0 | - | 1 – no activity |
| **ZeroClaw** | 4 (3 open) | 50 (40 open) | No | 4 – very high PR volume; 10 merges; active channel and config fixes |

*Health score factors: code velocity, bug responsiveness, community engagement, backlog health.*

## 3. OpenClaw's Position

OpenClaw is the ecosystem's largest reference project by a wide margin – 150 issues and 500 PRs touched in 24h dwarf all peers (next highest: Hermes Agent and IronClaw at 50 PRs each). Its primary advantages are:

- **Community scale**: Deepest user base, with feature requests from blind users to enterprise operators.
- **Architectural ambition**: The SQLite sessions refactor (#96625) is a foundational change that reduces complexity compared to NanoBot's Python-only approach or Moltis's Apple-native focus.
- **Cross-channel robustness**: Addresses message duplication (#69208) and subagent reliability (#67777) at a systemic level – a challenge others face piecemeal.

However, OpenClaw's scale also introduces risk: 140 open issues and 365 in-progress PRs create reviewer bottlenecks. P1 bugs like the silent reasoning-cost leak (#73182) and auth bypass (#72418) reflect the difficulty of maintaining security boundaries across a large codebase. Compared to IronClaw (which ships a complete capability-policy framework) or ZeroClaw (rapid channel fixes), OpenClaw's bug-triage velocity lags behind its feature-merge pace.

## 4. Shared Technical Focus Areas

Several requirements appear across three or more projects:

| Requirement | Projects | Details |
|---|---|---|
| **Session-state reliability** | OpenClaw, NanoBot, Hermes, PicoClaw, NullClaw | Duplicate messages, lost completions, stale sessions, per-session skill isolation |
| **Local/small-model compatibility** | OpenClaw, Moltis, CoPaw, ZeroClaw | Stringified tool args, null optional params, context budget overflow, model-specific config overrides |
| **Channel diversity & robustness** | OpenClaw, NanoBot, Hermes, PicoClaw, ZeroClaw, LobsterAI | Telegram, WhatsApp, Matrix, Discord, Signal, Weixin – each with unique reliability bugs |
| **Plugin/extensibility SDK** | OpenClaw (RFC #71736), ZeroClaw (WASM plugin host), NullClaw (approval flow) | Standardized contributions vs. per-project bespoke |
| **Multi-agent orchestration** | OpenClaw, NanoBot (A2A delegation), PicoClaw (agent collaboration bus), IronClaw (capability policy) | Delegation, subagent lifecycle, role-based access |

The strongest signal is **session-state integrity**, cited by four projects as a top user pain point – suggesting an ecosystem-wide infrastructure gap that could be filled by a shared library or protocol.

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | Hermes Agent | IronClaw | ZeroClaw | Others |
|---|---|---|---|---|---|---|
| **Target user** | Broad – desktop, server, office | Lightweight / CLI-first | Power users, multi-platform | Enterprise / multi-tenant | Channel-heavy, self-hosters | Niche (Moltis: Apple; CoPaw: QwenPaw backend) |
| **Architectural choice** | Monolithic + plugin slots | Python-only, dual runtime (Node) | Modular, many platforms | Crate-based, Reborn stack | WASM plugin host | Varies (Go, Zig, TypeScript) |
| **Differentiator** | Largest community, most channels | "Ultra-lightweight" label vs. reality | Rapid bug fix turnaround | Complete authorization framework | Wire-protocol-first RFC | Local-model focus (Moltis), test coverage (CoPaw) |
| **Release cadence** | Rolling, no named release today | v0.2.x expected soon | No release today | 0.30.x pending breaking changes | No release today | Staggered point releases |

The key axis of differentiation is **scale vs. simplicity**: OpenClaw optimizes for breadth, IronClaw for authorization depth, NanoBot for minimal footprint (despite Node.js dependency), and ZeroClaw for self-hosted automation.

## 6. Community Momentum & Maturity

**Tier 1 – High velocity, large contributor base**  
- **OpenClaw**: 500 PRs touched, 135 merged – extremely active but struggling with issue volume.  
- **Hermes Agent**: 50 PRs, 16 merged – healthy balance of fixes and features.  
- **IronClaw**: 50 PRs, 24 merged – strong maintainer throughput, epic completed.  
- **ZeroClaw**: 50 PRs, 10 merged – high feature churn, some PRs risk stagnation.

**Tier 2 – Moderate, steady iteration**  
- **NanoBot**: 24 PRs, 5 merged – fixing bugs and adding features; branding tension may affect engagement.  
- **LobsterAI**: 6 PRs, 4 merged – merging long-stale contributions; compatibility gap unresolved.  
- **PicoClaw**: 3 PRs, 2 merged – stable but low ambient activity.

**Tier 3 – Low activity / consolidation**  
- **NanoClaw**, **NullClaw**, **Moltis**, **CoPaw**: <10 PRs total, many unmerged (8+ days stale).  
- **TinyClaw**, **ZeptoClaw**: No activity – effectively dormant.

**Maturity signals**: IronClaw is the most mature in terms of testing infrastructure (new integration-test framework) and authorization design. OpenClaw is mature in community size but not in stability (many open P1 bugs). NanoBot and ZeroClaw are still in rapid feature-addition phase.

## 7. Trend Signals

From community feedback across all projects, five industry trends emerge:

1. **Multi-agent orchestration is the next frontier** – OpenClaw (subagent reliability), NanoBot (A2A delegation), PicoClaw (agent collaboration bus), and IronClaw (capability policy) are all investing in agent-to-agent workflows. This suggests a move beyond single-turn chatbots toward persistent, coordinated task execution.

2. **Local/private models gain traction** – Moltis, CoPaw, and OpenClaw all report issues specific to small local models (Gemma, oMLX, Cerebras). Users want to reduce API costs and retain data privacy, but the ecosystem lacks standardized tool-call formatting for non-OpenAI models.

3. **Channel reliability is table stakes** – Every project with Telegram, WhatsApp, or Matrix users reported at least one critical bug (stale locks, polling failures, encryption gaps). Users are demanding zero-message-loss guarantees before trusting assistants in production.

4. **Security and configurability demand is high** – Silent config toggles (#73182 in OpenClaw), auth bypass vectors (#72418), and permission tiers (#3114 in PicoClaw) show that users are deploying these assistants in sensitive environments and require admin controls.

5. **User experience fragmentation** – No project has converged on a standard UI pattern: OpenClaw has plugin slots, ZeroClaw has SOP cron, NanoBot has WebUI stucks. The ecosystem would benefit from shared UX patterns (e.g., approval cards, streaming indicators).

**Value for AI agent developers**: Projects that prioritize **local-model compatibility**, **plug-in authorization**, and **session-state persistence** will capture the growing self-hosted enterprise market. The rapid iteration seen in OpenClaw, Hermes, and ZeroClaw indicates that features are commoditizing fast – differentiation will come from reliability and security, not feature count.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-06-28

## 1. Today’s Overview
NanoBot saw **very high activity** on June 28: 24 pull requests were updated (19 open, 5 merged/closed) alongside 4 issues. The team is aggressively fixing bugs and shipping features, with **no new release** today. Significant efforts are concentrated on caching performance, UI reliability, and core agent robustness, while community discussions highlight tension around the project’s “ultra-lightweight” branding.

---

## 2. Releases
No new releases were published today.

---

## 3. Project Progress
**5 pull requests were merged or closed today**, reflecting steady forward movement:

- **[#4575](https://github.com/HKUDS/nanobot/pull/4575)** – Added repository guidelines (closed).  
- **[#4572](https://github.com/HKUDS/nanobot/pull/4572)** – Fixed CLI to allow OAuth login to become the main provider (closed, superseded by #4573).  
- **[#4510](https://github.com/HKUDS/nanobot/pull/4510)** – Fixed agent to drop malformed tool calls with invalid names (merged).  
- **[#4225](https://github.com/HKUDS/nanobot/pull/4225) & [#4357](https://github.com/HKUDS/nanobot/pull/4357)** – Added silent mode and lock_recipient for cron jobs (both merged).  

**Key feature advances** (still open):  
- **A2A peer delegation** ([#4571](https://github.com/HKUDS/nanobot/pull/4571)) – Native agent-to-agent collaboration with depth guards.  
- **Per-subagent model override** ([#4570](https://github.com/HKUDS/nanobot/pull/4570)) – `spawn` tool now accepts a `model` parameter.  
- **Mattermost channel** ([#4459](https://github.com/HKUDS/nanobot/pull/4459)) – WebSocket + REST integration.  
- **MCP image content** ([#4542](https://github.com/HKUDS/nanobot/pull/4542)) – Deliver images from MCP tools as proper artifacts.  
- **Subdirectory skills** ([#4504](https://github.com/HKUDS/nanobot/pull/4504)) – Organise skills into nested folders.  

---

## 4. Community Hot Topics

| Item | Type | Comments | 👍 | Summary |
|------|------|----------|----|---------|
| [#660](https://github.com/HKUDS/nanobot/issues/660) | Issue (closed) | 14 | 5 | Claims the project is “ultra-lightweight” but requires both Python and Node.js in Dockerfile; community wants a pure Python runtime. |
| [#4500](https://github.com/HKUDS/nanobot/issues/4500) | Issue (open) | 2 | 0 | WebUI stuck in “processing” after self-restart; stop button reports “No active task”. |
| [#4231](https://github.com/HKUDS/nanobot/pull/4231) | Issue (open) | 1 | 0 | Request to add model parameter to `spawn` tool for subagent override (now implemented in #4570). |
| [#4222](https://github.com/HKUDS/nanobot/issues/4222) | Issue (open) | 1 | 0 | Continuous prefix-cache invalidation caused by `max_messages` truncation and micro-compaction. |

**Analysis**: The most commented item, [#660](https://github.com/HKUDS/nanobot/issues/660), reveals a persistent user pain point: the “ultra-lightweight” tagline feels misleading when Node.js is a hard dependency. While the issue was closed, no public resolution was offered, which may leave some users dissatisfied. The other three issues are actively being addressed by linked PRs.

---

## 5. Bugs & Stability
Several bugs reported today have corresponding fix PRs:

| Bug | Severity | Fix PR | Status |
|-----|----------|--------|--------|
| WebUI self-restart leaves stuck streaming; stop button broken ([#4500](https://github.com/HKUDS/nanobot/issues/4500)) | **High** – UI becomes unusable | [#4565](https://github.com/HKUDS/nanobot/pull/4565) | Open |
| Prefix/prompt caching continuously invalidated by message truncation ([#4222](https://github.com/HKUDS/nanobot/issues/4222)) | **High** – degrades latency/performance | [#4568](https://github.com/HKUDS/nanobot/pull/4568) | Open |
| Malformed relay responses crash agent or feed bad demonstrations ([#4569](https://github.com/HKUDS/nanobot/pull/4569)) | **Medium** – crashes on certain upstream APIs | [#4569](https://github.com/HKUDS/nanobot/pull/4569) | Open |
| Weixin channel forced non-streaming due to missing config field ([#4567](https://github.com/HKUDS/nanobot/pull/4567)) | **Medium** – broken relay for Anthropic-compatible backends | [#4567](https://github.com/HKUDS/nanobot/pull/4567) | Open |
| Corrupt legacy session files silently dropped from `list_sessions` ([#4566](https://github.com/HKUDS/nanobot/pull/4566)) | **Low** – affects upgrade path for old installations | [#4566](https://github.com/HKUDS/nanobot/pull/4566) | Open |

All reported bugs are being actively patched, indicating a healthy maintenance pace.

---

## 6. Feature Requests & Roadmap Signals
- **Per-subagent model override** ([#4231](https://github.com/HKUDS/nanobot/issues/4231)) – Already implemented in PR [#4570](https://github.com/HKUDS/nanobot/pull/4570); likely to land in the next minor release.
- **Agent-to-Agent delegation** ([#4179](https://github.com/HKUDS/nanobot/issues/4179) via [#4571](https://github.com/HKUDS/nanobot/pull/4571)) – Native A2A with cross-delegation depth guard; a major architectural feature.
- **Mattermost channel** ([#4459](https://github.com/HKUDS/nanobot/pull/4459)) – New communication platform integration.
- **OAuth login as main provider** ([#4573](https://github.com/HKUDS/nanobot/pull/4573)) – Improves setup UX.
- **Reliability layer** ([#4534](https://github.com/HKUDS/nanobot/pull/4534)) – Long-output handling, verification feedback, service execution, runtime budget convergence.

**Prediction**: The next version (likely v0.2.3) will include the subagent model override, the A2A delegation framework (at least the basic mechanics), and the WebUI stability fixes. Mattermost integration and subdirectory skills are strong candidates for the following release.

---

## 7. User Feedback Summary
- **Lightweight branding**: Issue [#660](https://github.com/HKUDS/nanobot/issues/660) (closed) attracted strong pushback on the “ultra-lightweight” description when Node.js remains a dependency. Users expect a cleaner Python-only stack.
- **WebUI reliability**: Bug [#4500](https://github.com/HKUDS/nanobot/issues/4500) describes a frustrating experience where the UI appears stuck after a server restart, with no feedback from the stop button.
- **Subagent flexibility**: Contributor @jsapede (in [#4231](https://github.com/HKUDS/nanobot/issues/4231)) wants to assign different models to sub-agents – a common pattern for cost/quality trade-offs.
- **MCP tool content**: Developers using MCP servers with image outputs want those images rendered as inline artifacts, not base64 blobs.

Overall, users are pushing for more predictable performance, richer multimodal support, and clearer project positioning.

---

## 8. Backlog Watch
- **[#660](https://github.com/HKUDS/nanobot/issues/660)** – Though closed, this high-reaction issue received no maintainer comment explaining the Node.js dependency or whether a pure-Python path is planned. The community may consider this unresolved.
- **[#4222](https://github.com/HKUDS/nanobot/issues/4222)** – Opened 22 days ago, now has a fix PR ([#4568](https://github.com/HKUDS/nanobot/pull/4568)) but still open. No official assignment or milestone.
- **[#4231](https://github.com/HKUDS/nanobot/issues/4231)** – Opened 21 days ago, now addressed by PR [#4570](https://github.com/HKUDS/nanobot/pull/4570). No maintainer comment on the issue itself.

No truly stale issues were identified; all important items have received attention within the last few days.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# 📋 Hermes Agent Project Digest – 2026-06-28

**Data snapshot** (24h to 2026-06-28 23:59 UTC):
- Issues updated: 10 (9 open, 1 closed)
- PRs updated: 50 (34 open, 16 merged/closed)
- New releases: **none**

---

## 1. Today’s Overview

The project is experiencing **sustained high activity** with 50 PRs and 10 issues touched in the last day. A large batch of new contributions arrived – 16 PRs merged/closed and 17 new issues filed, many with detailed feature proposals and bug reports. The single closed issue (#39714) is a P1 venv path bug that was resolved after three comments and two upvotes. The PR pipeline shows a healthy mix of bug fixes, security hardening, new feature implementations, and infrastructure improvements (MCP, Kanban, Telegram). No release was published, but the code churn suggests a significant release may be forthcoming once the current wave of fixes and features stabilises.

---

## 2. Releases

*No new releases were published during the reporting period.*

---

## 3. Project Progress (Merged/Closed PRs)

16 PRs were merged or closed today. Notable advances include:

- **CLI & dependency management** – #44772 (open) addresses root npm dependency pruning during `hermes update`, but remains unmerged; several other PRs with `comp/cli` label were merged, suggesting incremental progress.
- **MCP reconnect resilience** – #54139 (open) fixes MCP reconnect retries after disconnect, with tool deregistration/re-registration logic (merged).
- **Kanban memory-aware dispatch** – #54145 (merged) adds persistence of reap registry to disk (`recent_exits.json`), lowers `max_in_progress`, and includes model key canonicalization.
- **Gateway & platform fixes** – Multiple P1/P2 bug-fix PRs were merged (see §5).
- **Telegram agent skill** – #54177 (merged) introduces a structured skill for Telegram workstream discipline.
- **Security hardening** – #54134 (open) pins n8n catalog refs, stops token fragment leaks, aligns core dependency pins.
- **Test coverage** – #54158 adds regression tests for Nvidia provider `extra_body` serialisation.

Full list: all 50 PRs have GitHub links in the provided data (top 20 detailed above).

---

## 4. Community Hot Topics

### Most-active issues (by comments and reactions)

| Issue | Title | Comments | 👍 | Status |
|-------|-------|----------|---|--------|
| [#39714](https://github.com/NousResearch/hermes-agent/issues/39714) | Bug: `hermes update` installs deps into hardcoded `venv/` while active env is `.venv` | 3 | 2 | **Closed** |
| [#35966](https://github.com/NousResearch/hermes-agent/issues/35966) | Feature: Native desktop/mobile client app | 2 | 2 | Open |
| [#50745](https://github.com/NousResearch/hermes-agent/issues/50745) | Feature: Mobile App for Hermes Agent | 2 | 0 | Open |

**Analysis:** The closed #39714 highlights a uv‑specific installation mismatch that was quickly addressed. The two mobile‑app feature requests (#35966, #50745) continue to draw community attention – users clearly want a first‑party Hermes client that bypasses third‑party chat platforms. The lack of maintainer response on these open feature requests (both last updated today but no assigned label or milestone) may indicate they are being tracked for a future roadmap item.

### Most-active PRs (by nature of discussion)

All top‑20 PRs show zero explicit comments in the data, but many are multi-author contributions with extensive commit histories (e.g., #54145, #54164, #54162). No single PR has disproportionate discussion; activity is spread across many small, focused changes.

---

## 5. Bugs & Stability

### New bugs reported today (ranked by severity)

| Issue | Title | Severity | Fix PR exists? |
|-------|-------|----------|----------------|
| [#54167](https://github.com/NousResearch/hermes-agent/issues/54167) | Stale Telegram bot token lock permanently kills platform | **P1** | ✅ [#54176](https://github.com/NousResearch/hermes-agent/pull/54176) (merged) |
| [#54174](https://github.com/NousResearch/hermes-agent/issues/54174) | `profile install` in Docker doesn't register s6 gateway slot | **P1** | ❌ No fix PR yet |
| [#54153](https://github.com/NousResearch/hermes-agent/issues/54153) | [Feature] Soft-warning at 80% tool-call iteration budget | P3 | ❌ |
| [#54154](https://github.com/NousResearch/hermes-agent/issues/54154) | Kanban dispatcher should capture worker crash diagnostics | P3 | ❌ |
| [#54155](https://github.com/NousResearch/hermes-agent/issues/54155) | Command allowlist path-prefix/glob patterns | P3 | ❌ |
| [#54156](https://github.com/NousResearch/hermes-agent/issues/54156) | Kanban auto-block failure summary annotation | P3 | ❌ |

### Bug fixes merged today (highlights)

| PR | Fixes | Severity |
|----|-------|----------|
| [#54176](https://github.com/NousResearch/hermes-agent/pull/54176) | Mark Telegram lock failure as **retryable** instead of permanently fatal | P1 (#54167) |
| [#54162](https://github.com/NousResearch/hermes-agent/pull/54162) | Treat zombie PIDs as dead in `_pid_exists` to unblock `--replace` | P1 (#42126) |
| [#54164](https://github.com/NousResearch/hermes-agent/pull/54164) | Reject unauthorized Telegram users before event construction | P1 (#40863) |
| [#54157](https://github.com/NousResearch/hermes-agent/pull/54157) | Block `hermes gateway start` from inside a running gateway session | P2 |
| [#54165](https://github.com/NousResearch/hermes-agent/pull/54165) | Fix QQAdapter `connect()` missing `is_reconnect` parameter (reconnect loop) | P2 |
| [#54159](https://github.com/NousResearch/hermes-agent/pull/54159) | Bootout stale launchd registration before macOS gateway bootstrap | P2 |
| [#54160](https://github.com/NousResearch/hermes-agent/pull/54160) | Accept typed replies to pending multi-choice clarify prompts | P2 |
| [#54161](https://github.com/NousResearch/hermes-agent/pull/54161) | Start MCP discovery for websocket sessions (desktop/dashboard) | P2 |
| [#54166](https://github.com/NousResearch/hermes-agent/pull/54166) | Non‑reusable sentinel for prefix‑matched secrets in file reads (security) | P2 |
| [#54168](https://github.com/NousResearch/hermes-agent/pull/54168) | Recognize Forgejo/Gitea webhook event headers | P2 |

**Assessment:** Today’s bug landscape is dominated by **telegram‑related issues** (stale locks, unauthorized users, multi‑choice clarify) – all addressed by merged PRs. The Docker profile‑install bug (#54174) remains without a fix, which could be a critical blocker for containerised deployments. No regression in stable functionalities was reported.

---

## 6. Feature Requests & Roadmap Signals

### New feature requests today

- **Per‑call child_timeout override** for `delegate_task` (#54152) – allows subagents different wall‑clock budgets.
- **Soft‑warning at ~80% tool‑call budget** (#54153) – agents can checkpoint state before max iterations.
- **Kanban crash diagnostics** (#54154) – attach stderr/dmesg/journalctl to task thread on worker death.
- **Command allowlist path‑prefix/glob patterns** (#54155) – auto‑approve chmod/mkdir under `~/.hermes/**`.
- **Kanban auto‑block failure summary** (#54156) – structured annotation with decomposition suggestion.

### Persistent user requests

- **Native desktop/mobile client** (#35966, #50745) – remains the most‑voted feature category (4 total 👍). No maintainer comment on either issue yet.

### Predictions for next version

The following features from PRs merged today are strong candidates for inclusion in the next release:
- **Memory‑aware Kanban dispatch** (#54145)
- **Telegram agent cockpit skill** (#54177)
- **app.nz provider** (#54146)
- **MCP reconnect resilience** (#54139)
- **Entry‑point plugin discovery** (#54150)
- **Command allowlist glob patterns** (#54155) – if the corresponding issue is accepted.

---

## 7. User Feedback Summary

- **Pain points**  
  - uv‑based installs cause confusion with hardcoded venv paths (#39714, now fixed).  
  - Telegram gateways can be permanently killed by stale lock files (#54167, PR merged).  
  - Docker profile install does not register s6 gateway slots (#54174, unfixed).  
  - Agents lack budget‑aware checkpointing (#54153) – users want to avoid wasted work.

- **Use cases**  
  - Sub‑delegation with different timeouts (#54152) – multi‑step research / firmware builds.  
  - Mobile/desktop direct interaction without Telegram/Discord (#35966, #50745).  
  - Self‑hosted git workflows (Forgejo/Gitea webhook support in #54168).  
  - Nvidia LLM provider with custom `chat_template_kwargs` (#54158).

- **Satisfaction signals**  
  - Quick turnaround on the venv bug (#39714 closed in 23 days with 3 comments).  
  - Many one‑day bug‑fix PRs (e.g., lock retry, zombie PID, unauthorized users) show responsive maintainers.  
  - Community contributors are actively submitting both bug fixes and features (c03rad0r, jerryfang527, mmysior, chadsm-sys, etc.).

---

## 8. Backlog Watch

### Issues needing maintainer attention

| Issue | Last Updated | Age (days) | Reason for Watch |
|-------|--------------|------------|------------------|
| [#35966](https://github.com/NousResearch/hermes-agent/issues/35966) | 2026-06-28 | 28 days | Native client app (2 👍, no maintainer reply) |
| [#50745](https://github.com/NousResearch/hermes-agent/issues/50745) | 2026-06-28 | 6 days | Duplicate mobile app request; should be marked as duplicate or merged with #35966 |
| [#44772](https://github.com/NousResearch/hermes-agent/pull/44772) | 2026-06-28 | 16 days | `comp/cli` fix for npm workspace deps – still open with no review comments |

### PRs lingering without merge

- **#44772** – critical for Node.js dependency management; submitted 16 days ago, no reviewer activity.  
- **#54134** (security, P2) – hardens catalog refs and auth diagnostics; merged today per data? Actually it's still open (status `OPEN`). Needs final review.

---

**Overall project health:** **High activity, low friction.** Bugs are being fixed rapidly, the community is engaged, and the feature pipeline is rich. The only concerns are the unaddressed Docker profile‑install bug and the two client‑app requests lacking maintainer feedback. The volume of P1 bug‑fix PRs merged today indicates strong focus on stability.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest — 2026-06-28

## 1. Today's Overview

PicoClaw shows measured activity with 3 issues and 3 pull requests updated in the last 24 hours. A new open bug (#3194) concerning unhandled encrypted Matrix messages was reported, while two older issues and two PRs were closed, including a significant feature for inter-agent collaboration (#2937) and a fix for MCP argument parsing (#3048). A new PR (#3193) introducing a “simplex” channel type remains open and under review. No new releases were published. Overall the project is in a stable patch-and-feature cycle, with maintainers addressing both bug reports and community-driven enhancements.

## 2. Releases

*No new releases in the last 24 hours.*

## 3. Project Progress

Two pull requests were merged/closed today:

- **#3048** (closed, merged) – **fix(mcp): reject unknown pre-positional flags in `add`**  
  *Author: afjcjsbx*  
  Fixes an argument parsing bug where root-level persistent flags like `--no-color` leaked into the MCP subcommand parser.  
  [GitHub](https://github.com/sipeed/picoclaw/pull/3048)

- **#2937** (closed, merged) – **Feat/agent collaboration**  
  *Author: afjcjsbx*  
  Introduces a first-class Agent Collaboration Bus with per-agent mailboxes, collaboration threads, structured message envelopes, and permission-aware delivery. This is a major architectural addition enabling multi-agent workflows.  
  [GitHub](https://github.com/sipeed/picoclaw/pull/2937)

The open PR **#3193** (simplex channel type) is still awaiting review but indicates ongoing interest in expanding channel support.

## 4. Community Hot Topics

- **#2472** (closed, 7 comments, 1 👍) – **BUG: `list_dir` fails on Windows with path separator mismatch**  
  *Author: ut2or1*  
  The most discussed item. The user reported that `list_dir` crashes on Windows because backslashes aren’t converted to forward slashes for Go’s `os.Root`. This is a cross‑platform compatibility pain point for Windows users. The issue was closed, implying a fix was applied.  
  [GitHub](https://github.com/sipeed/picoclaw/issues/2472)

- **#3114** (closed, 2 comments, 1 👍) – **Future Request: Telegram channel permission levels by conversation type**  
  *Author: v2up-32mb*  
  A detailed proposal to add per‑channel-type access control (private chat / group / channel) to the Telegram adapter. The request highlights security concerns (e.g., preventing shell execution in groups). Currently closed as stale, but the underlying need for granular permissions remains unaddressed.  
  [GitHub](https://github.com/sipeed/picoclaw/issues/3114)

## 5. Bugs & Stability

| Severity | Issue | Description | Status |
|----------|-------|-------------|--------|
| **High** | [#3194](https://github.com/sipeed/picoclaw/issues/3194) | "Received encrypted message but crypto is not enabled" – Matrix adapter receives an encrypted event but fails to decrypt when `crypto` is disabled. Potentially blocks Matrix usage for users with end‑to‑end encryption enabled on their homeserver. | Open, no fix PR yet |
| **Medium** | #2472 (closed) | Windows path separator bug in `list_dir`. Resolved, but Windows users may have encountered data access failures prior to fix. | Fixed and closed |
| **Low** | #3048 (merged) | Pre‑positional flag leak in `mcp add` could cause confusing errors. Fixed. | Merged and closed |

**#3194** is the only open bug reported today. It affects Matrix channel users who have E2EE enabled. No workaround or fix has been provided in the issue comments.

## 6. Feature Requests & Roadmap Signals

- **Agent Collaboration** – Now merged (#2937), this foundational feature enables inter‑agent communication. Future versions may expose this as a user‑facing API or integrate it with tool orchestration.
- **Simplex Channel** – The open PR #3193 proposes a new channel type (likely for the Simplex messaging protocol). If merged, it expands PicoClaw’s multi‑channel reach.
- **Telegram Permission Tiers** – Issue #3114, though closed as stale, reflects a recurring request for context‑aware access control. The community may push for a re‑open or alternative implementation.
- **Windows Compatibility** – The closed bug #2472 shows a clear need for thorough cross‑platform testing. Future releases should include Windows‑specific CI or documentation.

**Prediction**: The next minor release (v0.2.7) will likely include the agent collaboration bus, the MCP flag fix, and possibly the simplex channel if #3193 is merged soon. A patch release for the Matrix encryption bug (#3194) is also probable.

## 7. User Feedback Summary

- **Pain Points**:
  - Windows users face path separator issues (now fixed in #2472).
  - Matrix users with encrypted rooms cannot use PicoClaw if crypto is disabled (#3194).
  - Telegram users desire per‑group permission scopes to prevent accidental dangerous commands.
- **Use Cases**:
  - Multi‑agent workflows (agent collaboration feature).
  - Cross‑platform automation (Windows + Linux).
  - Secure group chat bots (Telegram, Matrix).
- **Satisfaction**: The quick closure of #2472 suggests responsive maintainers for confirmed bugs. However, the absence of any comment on #3194 (even to acknowledge) may cause concern.

## 8. Backlog Watch

| Item | Last Maintainer Action | Days Since Update | Concern |
|------|------------------------|-------------------|---------|
| [#3114](https://github.com/sipeed/picoclaw/issues/3114) (Telegram permission tiers) | Closed as stale | 1 day (closed 2026-06-27) | No maintainer rationale given; feature request may be worth revisiting. |
| [#3194](https://github.com/sipeed/picoclaw/issues/3194) (Matrix crypto bug) | None | 0 days | No maintainer response yet. If left unanswered for >48h, it becomes a stale candidate. |
| [#2472](https://github.com/sipeed/picoclaw/issues/2472) (Windows path bug) | Closed | 1 day | Resolved; no further action needed. |

The only item requiring immediate maintainer attention is **#3194** – a clean bug repro with environment details. A quick triage or workaround note would improve project responsiveness.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-28

## 1. Today’s Overview
Activity remains moderate, with seven open pull requests updated in the last 24 hours but no merges or new releases. A single new issue was opened, highlighting a critical crash when using the OpenAI provider with agent-runner containers. The project is clearly in an iterative stabilization phase, with several fixes and one feature PR in progress. Despite low issue volume, the presence of container- and deployment-related PRs suggests ongoing infrastructure hardening. No releases were published today.

## 2. Releases
No new releases today.

---

## 3. Project Progress
No pull requests were merged or closed today. However, the following open PRs represent ongoing work:

- **#2877** – [feat(telegram): native rich rendering via Bot API 10.1 sendRichMessage](https://github.com/nanocoai/nanoclaw/pull/2877) (new, created today)  
  Implements Telegram’s new rich message API, a notable feature addition.
- **#2875** – [Deploy/coolify](https://github.com/nanocoai/nanoclaw/pull/2875) (updated today)  
  Adds a new deployment target (Coolify).
- **#2873** – [fix(skills): split pre-flight from credentials](https://github.com/nanocoai/nanoclaw/pull/2873)  
  Separates credential validation from skill code refresh, enabling cleaner `/update-skills` workflows.
- **#2874** – [fix(signal): survive signal-cli boot flaps](https://github.com/nanocoai/nanoclaw/pull/2874)  
  Prevents crash-looping when the Signal CLI backend experiences transient failures.
- **#2822** – [refactor(container-runner): drop dead /workspace/global mount](https://github.com/nanocoai/nanoclaw/pull/2822) (updated recently)  
  Cleans up unused container mounts.
- **#2823** – [fix: remove groups/global/CLAUDE.md (host deletes it on every startup)](https://github.com/nanocoai/nanoclaw/pull/2823)  
  Addresses an annoying startup-side effect where the host deletes a critical instruction file.
- **#2824** – [fix: drop stale "Global Memory" instruction from main seed prompt](https://github.com/nanocoai/nanoclaw/pull/2824)  
  Removes outdated system prompt content.

---

## 4. Community Hot Topics
The only issue created today (#2876) has drawn attention due to its impact on a core flow (OpenAI provider):

- **[#2876] – Add OpenAI provider to agent-runner – container crash on spawn**  
  https://github.com/nanocoai/nanoclaw/issues/2876  
  The user reports that changing a group’s provider to `openai` via CLI succeeds and persists in the database, but the container immediately crashes on the next message. This exposes a gap in container-side provider support and suggests the agent-runner is not fully wired for third-party providers. Expect this to be a high-priority fix.

Among PRs, **#2877** (Telegram rich rendering) is the most prominent feature addition, signaling community interest in richer messaging channels. **#2874** (Signal crash resilience) addresses a common pain point for users of the Signal integration.

No PRs have received comments or reactions yet (all show 0), likely because they have not been widely noticed or are still early-stage.

---

## 5. Bugs & Stability
One new bug was reported today:

- **Critical: #2876 – Container crash when using OpenAI provider**  
  https://github.com/nanocoai/nanoclaw/issues/2876  
  *Severity: High* – Blocks users from using any OpenAI models after switching provider. No fix PR exists yet; likely requires changes in container runtime configuration or environment variable injection.

Other bugs being addressed by open PRs:

- **#2874 – Signal client boot flaps cause agent crash-loop** (medium severity, fix PR exists)
- **#2823 – Host deletes CLAUDE.md on startup** (low severity, fix PR exists)

No regressions or previously fixed issues have resurfaced today.

---

## 6. Feature Requests & Roadmap Signals

- **OpenAI provider support** – The crash in #2876 reveals unmet demand for a working OpenAI integration. A fix is likely needed before this can be considered stable.
- **Telegram rich rendering** – PR #2877 brings native `sendRichMessage` support, aligning with Telegram’s latest API. This may land in the next minor release (2.2.x).
- **Coolify deployment** – PR #2875 adds a new deployment target, suggesting growing self-hosting requirements.
- **Skill code refresh improvements** – PR #2873 makes it safer to update skill code without re-entering credentials, a quality‑of‑life improvement likely to be included soon.

Predictions for next release (v2.2): Inclusion of fixes for #2874 and #2823, plus the Telegram feature (#2877) if reviewed and merged quickly.

---

## 7. User Feedback Summary

The single issue #2876 highlights a clear pain point: users who switch from the default provider to OpenAI experience a broken setup with no error feedback other than a crash. The user appears frustrated with the silent failure after a seemingly successful CLI command. No positive feedback or praise was visible in today’s data.

PRs such as #2874 (Signal reliability) and #2877 (Telegram richness) indicate that users are deploying NanoClaw across multiple messaging channels and value stability as much as new features. The absence of comments suggests that the community may be waiting for these PRs to land before testing.

---

## 8. Backlog Watch

The following open PRs from **CutSnake01** have been pending for eight days (since June 20) and have received no comments or maintainer responses:

- **#2822** – [refactor(container-runner): drop dead /workspace/global mount](https://github.com/nanocoai/nanoclaw/pull/2822)
- **#2823** – [fix: remove groups/global/CLAUDE.md (host deletes it on startup)](https://github.com/nanocoai/nanoclaw/pull/2823)
- **#2824** – [fix: drop stale "Global Memory" instruction from main seed prompt](https://github.com/nanocoai/nanoclaw/pull/2824)

All three are labelled as “follows-guidelines” and address straightforward cleanup or bugfixes. They may be waiting for code review or CI validation. Their age risks merge conflicts if left unattended. A maintainer should triage these as a batch to keep the backlog healthy.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest – 2026-06-28

## 1. Today’s Overview
The project shows low activity over the last 24 hours, with one new pull request and one existing issue receiving updates. No releases or merged PRs occurred today. The single open PR introduces a structured approval flow for agent tools, while the open issue reports a build failure on Android/Termux that has been lingering for over two months. Overall project health appears stable but with limited momentum; community contributions are focused on core agent interaction mechanics, while Android compatibility remains a persistent pain point.

## 2. Releases
*No new releases were published today. The latest available release remains **v2026.4.17** (referenced in Issue #868).*

## 3. Project Progress
**Merged/closed PRs today:** None.  
**Notable new PR (open):**  
- **#969 – feat(agent): structured approval_request / approval_response flow**  
  *Author: valonmulolli*  
  Implements a two-turn approval mechanism for tools that raise `error.ApprovalRequired`. The agent catches the error, stores a pending approval, and emits an SSE event for the channel to render a UI.  
  → [Link to PR](https://github.com/nullclaw/nullclaw/pull/969)

No other feature advancements or fixes were merged today.

## 4. Community Hot Topics
- **Issue #868 – [bug] zig build fails on Android/Termux (aarch64)**  
  *4 comments, last updated 27 Jun*  
  User reports a build failure (`AccessDenied on options.zig linkat`) using Zig 0.16.0 on LineageOS 22.2 (Android aarch64). The issue was created on 23 Apr and has received limited maintainer engagement. Community comments suggest possible workarounds but no official resolution.  
  → [Link to Issue](https://github.com/nullclaw/nullclaw/issues/868)  
  *Underlying need:* Reliable cross‑platform build support, specifically for Android/Termux environments where native compilation is desired.

- **PR #969 – Structured approval flow**  
  *Created and updated today, no comments yet*  
  This feature targets improved user‑agent interaction by adding a formal approval step before dangerous tool execution. It is likely to generate discussion as it touches security and UX.  
  → [Link to PR](https://github.com/nullclaw/nullclaw/pull/969)

## 5. Bugs & Stability
**Active bug (open):**  
- **Issue #868 – Build failure on Android/Termux (aarch64)** – **Severity: Medium**  
  The bug prevents compilation of the project on a major mobile platform. The error involves a file-linking permission issue (`AccessDenied`) during Zig’s LLD linker step. No fix PR exists yet. The bug has been open for over two months with only community comments.  
  → [Link to Issue](https://github.com/nullclaw/nullclaw/issues/868)

No new crashes, regressions, or other bug reports were filed today.

## 6. Feature Requests & Roadmap Signals
**Explicit feature request embedded in PR #969:**  
The structured `approval_request`/`approval_response` flow is a user‑facing feature likely driven by the need for safer shell tool usage. This aligns with typical agent‑safety roadmaps.  

**Implicit signals from Issue #868:**  
Users request native Android support for development and testing. While not a formal feature request, the build failure indicates a gap in cross‑platform compatibility that may become a higher priority if mobile usage grows.

**Prediction for next version:**  
If PR #969 is merged and refined, the next release could include the structured approval mechanism as a core agent feature. Android build fixes may also be included if maintainers address #868.

## 7. User Feedback Summary
- **Pain points:**  
  - Build process is broken on Android (Termux / aarch64) – user expressed frustration with access‑denied errors when linking.  
  - The lengthy open status of #868 suggests a lack of maintainer responsiveness to platform‑specific issues.

- **Use cases:**  
  - Development on mobile Linux environments (Termux) – likely for on‑the‑go testing or edge deployments.  
  - Agent interaction safety – the new PR addresses a clear use case where users want to confirm dangerous tool operations.

- **Satisfaction/Dissatisfaction:**  
  - No explicit satisfaction signals today.  
  - Dissatisfaction is implied by the lack of progress on #868 (zero maintainer replies in >2 months, only community workarounds).

## 8. Backlog Watch
- **Issue #868 – Build failure on Android**  
  *Created: 2026-04-23 | Last updated: 2026-06-27 | Comments: 4 | No maintainer response*  
  This issue is approaching three months without a maintainer reply or fix. It blocks Android users from building the project. Given the low activity of the project, it risks becoming stale.  
  → [Link to Issue](https://github.com/nullclaw/nullclaw/issues/868)

No other important issues or PRs are languishing unanswered. PR #969 is very new and should be reviewed promptly to maintain contributor momentum.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-06-28

## 1. Today's Overview

IronClaw saw extremely high development activity on 2026-06-28, with **50 pull requests updated** in the last 24 hours (26 open, 24 merged/closed) and **9 issues updated** (7 closed, 2 open). The overwhelming focus remains the **capability-policy epic (#5261)**, which appears to have reached code-complete status — all five sub-issues (#5266, #5267, #5268, #5272, #5273) were closed yesterday after the foundational policy model crate (#5262) was merged. A significant new integration-test framework for the Reborn stack was also delivered (#5381). No new releases were cut; a pending release PR (#5311) with breaking changes in `ironclaw_common` and `ironclaw_skills` is still open. Project health is strong with high throughput, though two open issues warrant attention: a recurring nightly E2E failure (#4108, open since May 27) and a newly reported Google OAuth token refresh blocker (#5378).

## 2. Releases

**No new releases.** The last release PR (#5311) remains open and contains:
- ⚠ API breaking changes in `ironclaw_common` (0.4.2 → 0.5.0)
- ⚠ API breaking changes in `ironclaw_skills` (0.3.0 → 0.4.0)
- `ironclaw` binary version bump from 0.24.0 → 0.29.1
- Minor compatible changes in `ironclaw_safety` (0.2.2 → 0.2.3) and `ironclaw_skill_learning` (0.1.0 → 0.1.1)

**Migration notes:** Teams consuming `ironclaw_common` or `ironclaw_skills` should prepare for API surface changes once this release merges.

## 3. Project Progress

The capability-policy epic (#5261) reached a major milestone — **all five sub-issues closed** and the foundational `ironclaw_capability_policy` crate merged (#5262). This introduces the four-dimension policy model (configuration, identity, approval, availability) with a precedence cascade and in-memory store-backed resolver.

**Key merged/closed PRs:**
- **#5262** — `ironclaw_capability_policy` crate (policy model, part of epic #5261)
- **#5381** — New **Reborn integration-test framework** (slices 1–2): scripted-SDK seam, tool-call/egress assertions, full design. Enables in-process testing through the real internal stack with only model providers faked.
- **#4970** — Fix: record `Delivered` (not `Skipped`) for non-OAuth auth-denied triggered runs
- **#5384** — Pin WebUI v2 frontend Node tooling (Node 22, `.node-version`, `.nvmrc`)
- **#5378** — Google OAuth token refresh failure filed (closed with fix in PR #5388)

## 4. Community Hot Topics

The most active discussion centers on the **capability-policy epic** (#5261) with 5 sub-issues all authored by `zetyquickly`, each generating 1–2 comments. The epic defines a comprehensive per-user authorization surface for the Reborn stack.

**Most commented activity:**
- **#5272** — REST-created local users + dynamic auth (2 comments)
- **#5268** — Admin REST surface for granting permissions (1 comment)
- **#5273** — Four-dimension policy delta store + resolver (1 comment)
- **#5267** — Availability resolver at dispatch seam (1 comment)
- **#5266** — DB-backed user role + admin gate (1 comment)
- **#5261** — Epic: Reborn capability policy (1 comment)

**Underlying need:** The project is building toward a complete **multi-tenant, role-based authorization system** that distinguishes between Owner, Admin, and Member users with granular tool/skill access control. The new issue **#5385** (open, 0 comments) explicitly requests the final state of fine-grained user configuration.

**Active PR drawing attention:** 
- **#5354** — Reborn WebUI v2 live QA canary (size XL, risk medium, core contributor `serrrfirat`). Extends Playwright-driven canary testing with live LLM/tool integrations.
- **#5279** — Fix Reborn queued message steering (size XL, by `ilblackdragon`). Preserves active-run UI behavior around send, approvals, and backend error display.

## 5. Bugs & Stability

**High severity:**
- **#5378** — Google OAuth token refresh fails with `BackendUnavailable` on `hosted-single-tenant` and `local-dev` profiles. Forces re-authentication roughly every hour for all Google OAuth-backed capabilities (Gmail, Calendar, Drive). **Fix exists:** PR #5388 ([nearai/ironclaw PR #5388](https://github.com/nearai/ironclaw/pull/5388)) addresses RS256 `id_token` decoding after `jsonwebtoken` 10.x bump.
- **#4108** — Nightly E2E scheduled run failing (open since **2026-05-27**, 32 days). Workflow run: [actions/runs/28311379145](https://github.com/nearai/ironclaw/actions/runs/28311379145). No comments or assignee — appears stale.

**Medium severity:**
- **#5385** — New issue reporting lack of fine-grained user configuration (Owner, Admin, Member roles). The owner is currently set via env vars only; no admin surface exists for managing non-owner users.
- **#5306** — Fix for `ask-each-time` approval resume loop (PR open by `italic-jinxin`). Prevents re-approval gates being applied to already-approved capability leases.
- **#5338** — Fix for generic `invalid_input` error detail (PR open). Was hiding real failure reasons behind vague "driver protocol error" messages.

**Low severity:**
- **#5297** — Stale gate projection rows in WebUI stream (fix PR open by `hanakannzashi`)
- **#5365** — WebUI v2 chat Retry button was wired to a no-op stub (fix PR open by `henrypark133`)
- **#5084** — Automations page redesign (PR open by `achalvs`) — functional but layout improvement

## 6. Feature Requests & Roadmap Signals

**Explicitly requested:**
- **#5385** — "Add Capability Policy" requests fine-grained user configuration with Owner, Admin, and Member roles. The capability-policy epic (#5261) already addresses this, suggesting this feature is **imminent** for the next release.
- **#5279** (PR) — Queued message steering for busy-thread user messages, preserving real-time UI behavior. Likely targeting the v2.0 release.

**Emerging themes:**
- **Authorization & multi-tenant readiness**: The capability-policy epic is the strongest roadmap signal — IronClaw is building toward enterprise-grade user management.
- **Testing infrastructure**: Two major test framework efforts (Reborn integration-test framework #5381, WebUI v2 QA canary #5354) suggest a build-up to a stable v2 release.
- **CI reliability**: The 32-day-old nightly E2E failure (#4108) may block confidence in automated testing, though no maintainer has commented.

**Prediction:** The next release (0.30.x) will likely include:
1. Full capability-policy implementation (Owner/Admin/Member roles + per-user tool/skill auth)
2. The Reborn integration-test framework
3. Google OAuth token refresh fix
4. Breaking changes in `ironclaw_common` and `ironclaw_skills`

## 7. User Feedback Summary

**Expressed pain points:**
- **OAuth re-authentication every ~1 hour** (issue #5378 by `thisisjoshford`) — "forces a re-authentication roughly once an hour" — a critical user experience blocker for any Google-integrated workflow.
- **Generic error messages** — PR #5338 addresses "vague 'driver protocol error'" and "only the bare kind" shown to users on tool failures.
- **Broken Retry button** — PR #5365 confirms the WebUI v2 Retry button "was wired to a truthy no-op stub, so it rendered but did nothing" — indicates UI polish gaps in the v2 frontend.
- **Stale gate row display** — PR #5297 fixes "stale blocked-gate suppression" in the WebUI stream, suggesting users saw outdated approval states.

**Satisfaction signals:**
- High developer engagement — 50 PRs updated in 24 hours indicates active, motivated contributor base
- The automations page redesign (#5084) by new contributor `achalvs` shows community investment in UI quality
- 24 merged/closed PRs in one day demonstrates strong maintainer responsiveness

## 8. Backlog Watch

**Critical attention needed:**
- **#4108** — Nightly E2E failure ([Issue #4108](https://github.com/nearai/ironclaw/issues/4108)) — Open since **2026-05-27 (32 days)**. Zero comments, no assignee. This is a **recurring scheduled run failure** that undermines confidence in CI/CD. Severity increases daily — the pipeline has been broken for a month without maintainer acknowledgment.
- **#3706** — Dependencies bump PR ([PR #3706](https://github.com/nearai/ironclaw/pull/3706)) — Open since **2026-05-16 (43 days)**. Bumps `postcss`, `@remotion/cli`, `@remotion/tailwind-v4` for the architecture video docs. Blocked by security and compatibility concerns? No comments from maintainers.

**Stale but important:**
- **#5297** — Stale gate projection rows fix (PR open 2 days). While not long-stale, the underlying issue #5218 could have existed for weeks, impacting WebUI v2 users.
- **#5084** — Automations page redesign (PR open since 2026-06-18, 10 days). No maintainer feedback on the contributor's work.

**Recommendation:** The nightly E2E failure (#4108) should be the team's top operational priority — a month-long CI breakage creates risk for regressions in the active capability-policy and OAuth work.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-06-28

## 1. Today's Overview

Project activity is moderate but concentrated on housekeeping and merging long-pending contributions. No new releases or fresh issues were created today; instead the maintainers closed 4 stale issues (all from early April) and merged/closed 4 PRs, including a major conflict-resolved artifact preview pipeline (#1441) and a hotfix for MCP SSE/streaming HTTP support (#1001). Two open PRs remain from the same period, suggesting a backlog of feature work is still awaiting review. Overall, the project appears to be in a steady state with incremental improvements being folded in, though community-facing compatibility updates (e.g., OpenClaw v2026.3.24) remain unaddressed.

## 2. Releases

**No new releases today.** The latest release remains unnamed (last release date unknown from data).

## 3. Project Progress

Four pull requests were closed or merged in the last 24 hours, representing both feature additions and critical fixes:

- **#1440** – [`feat(cowork): 将已选技能标签移至输入框内顶部展示`](https://github.com/netease-youdao/LobsterAI/pull/1440)  
  *Author: gongzhi-netease* – Relocated active skill badges from a crowded bottom toolbar to a dedicated position above the textarea, improving layout clarity for users with many selected skills.

- **#1441** – [`feat(artifacts): add extensible preview pipeline for HTML, React and Mermaid`](https://github.com/netease-youdao/LobsterAI/pull/1441)  
  *Author: febugcoder* – A conflict-resolved and bug-fixed revival of a previously stalled PR (#1011). Adds a pluggable preview system for code artifacts (HTML, React components, Mermaid diagrams) in Cowork sessions. This is a significant feature enhancement.

- **#1445** – [`fix(skills): 修复技能重复导入无校验及 zip 导入目录名异常的问题`](https://github.com/netease-youdao/LobsterAI/pull/1445)  
  *Author: gongzhi-netease* – Fixes two bugs: (1) zip-imported skills getting random temporary directory names, and (2) duplicate skills being silently installed with `-1` suffixes. Now reads the skill name from `SKILL.md` frontmatter and properly blocks duplicates.

- **#1001** – [`hotfix：增加对 sse 和 流式http 的mcp支持`](https://github.com/netease-youdao/LobsterAI/pull/1001)  
  *Author: callmekeyboardman* – Patches the MCP server manager to actually start servers using SSE and streaming HTTP transports (previously only `stdio` worked, silently ignoring others). This closes a significant gap in MCP configuration reliability.

Additionally, two PRs remain **open** (see Backlog Watch).

## 4. Community Hot Topics

All four closed issues received only 2–3 comments each, so activity is low. The most discussed was:

- **#1443** – [有计划支持新版本的openclaw吗？](https://github.com/netease-youdao/LobsterAI/issues/1443) (3 comments, 👍0)  
  *Author: Juzisuan965* – Request for compatibility with OpenClaw v2026.3.24, which introduced breaking changes. The issue was closed as stale without a resolution. This underscores a latent compatibility gap that may resurface if users continue to upgrade their OpenClaw deployment.

Other issues (#1437, #1439, #1442) discussed UI/UX bugs around scheduled tasks, skill deactivation, and skill display after conversation. These were all closed as stale, suggesting the maintainers considered them either fixed by recent PRs or not reproducible.

## 5. Bugs & Stability

No new bugs were reported today. However, the bugs that were fixed by the merged PRs today reflect previously unresolved stability concerns:

| Severity | Bug Description | Fixed in PR |
|----------|----------------|------------|
| **High** | MCP configurations with SSE/streaming HTTP transport silently did not work – users believed sync completed but servers never started. | [#1001](https://github.com/netease-youdao/LobsterAI/pull/1001) |
| **Medium** | Zip-imported skills received random directory names instead of the intended skill name from `SKILL.md`. | [#1445](https://github.com/netease-youdao/LobsterAI/pull/1445) |
| **Medium** | Duplicate skills could be imported without validation, injecting redundant system prompts and affecting LLM routing stability. | [#1445](https://github.com/netease-youdao/LobsterAI/pull/1445) |
| **Low** | Active skill badges displayed in a crowded bottom toolbar alongside attachment buttons, causing visual clutter when many skills were selected. | [#1440](https://github.com/netease-youdao/LobsterAI/pull/1440) |

All four issues that were closed today (unrelated to the above PRs) remain unlinked to fix PRs – likely closed due to staleness.

## 6. Feature Requests & Roadmap Signals

Two signals point to features that may land in the next release:

- **OpenClaw v2026.3.24 compatibility** (#1443) – While closed as stale, the user’s local failure (crash on upgrade) indicates demand for breaking change adaptation. No maintainer response was recorded, so this may be deferred.
- **Scheduled task UI overhaul** – Open PR [#1488](https://github.com/netease-youdao/LobsterAI/pull/1488) (by gongzhi-netease) proposes a complete card-grid redesign, search, date filtering, and history grouping. This is a substantial UX enhancement that has been open since April 5 – if merged, it will be a highlight of the next version.
- **Per-session skill selection** – Open PR [#1494](https://github.com/netease-youdao/LobsterAI/pull/1494) decouples active skills from a global store to per-session state. This directly addresses user confusion in #1442, where skills selected in one session carried over to others.

Given the author (gongzhi-netease) is a frequent contributor and both PRs are almost a month old, they likely require a final review pass and could be merged soon.

## 7. User Feedback Summary

From the closed issues and PR descriptions, the following pain points were expressed:

- **Breaking-change anxiety** – Users operating OpenClaw v2026.3.24 experienced launch failures with no clear migration path or official statement from LobsterAI.
- **Confusing skill behavior** – Skills selected for an Agent either disappeared after a conversation (#1442) or could still be invoked even when toggled “off” (#1439). Users questioned the purpose of agent-level skill selection.
- **UI friction** – A scheduled task creation bug (#1437): clearing the calendar with “No Repeat” selected caused the create button to silently fail, with no error feedback.
- **Disconnect between user expectation and artifact preview** – The original PR #1011 (now merged as #1441) was requested by users wanting richer previews of code outputs in Cowork sessions.

Overall satisfaction appears mixed: the project is actively merging fixes, but communication around compatibility and stale issues may leave users uncertain.

## 8. Backlog Watch

Two open PRs from early April have not been merged or closed despite being updated today:

- **[#1488](https://github.com/netease-youdao/LobsterAI/pull/1488) – Scheduled task UI upgrade (open since April 5)**  
  *Author: gongzhi-netease* – No reviewer comments or approvals. The lack of movement on a feature this size (card grid, search, history enhancements) suggests either blockers or low reviewer bandwidth.

- **[#1494](https://github.com/netease-youdao/LobsterAI/pull/1494) – Per-session skill selection (open since April 6)**  
  *Author: gongzhi-netease* – Similarly unreviewed. This PR directly addresses a reported bug (#1442) and aligns with the project’s trend toward session-isolated state.

These two items represent the longest-pending active contributions and would benefit from maintainer attention to reduce backlog and deliver user-requested improvements.

---

*Generated from GitHub data for netease-youdao/LobsterAI, snapshot 2026-06-28.*

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest – 2026-06-28

## Today's Overview
Project activity remains steady with four open pull requests and one new bug issue updated in the last 24 hours. No code was merged today, but the maintainers are actively iterating on fixes for local model compatibility and build dependencies. The two newest PRs (#1139, #1138) were both created today, indicating ongoing development momentum. Community engagement is low, with zero comments or reactions on current items.

## Releases
No new releases today.

## Project Progress
No pull requests were merged or closed today. The following open PRs represent the current focus of development:
- **#1139** – Fixes a build issue where enabling the `metrics` feature inadvertently pulls in the entire `matrix-sdk` dependency (even when the Matrix channel is disabled).  
- **#1138** – Addresses a critical agent‑side problem: full‑resolution images embedded as base‑64 can consume ~350K tokens, exceeding the context budget and causing prompt rejection on every turn.  
- **#1136** – Coerces stringified scalar tool arguments (e.g., `"true"` instead of `true`) emitted by small local models (Gemma 4, oMLX) before validation, preventing pre‑dispatch failures.  
- **#1098** – Handles `null` values for optional browser‑tool parameters, another quirk of smaller local models.

All fixes target robustness for local and resource‑constrained environments.

## Community Hot Topics
No issues or PRs generated comments or reactions today. The single open issue (#1137) has zero discussion. The most substantial underlying need appears to be **compatibility with Apple‑signed environments** (Issue #1137) and **reliable operation with small local LLMs** (PRs #1136, #1098).

## Bugs & Stability
| Issue | Severity | Description | Fix PR Exists? |
|-------|----------|-------------|----------------|
| [#1137 – Apple Container ID exceeds name limit](https://github.com/moltis-org/moltis/issues/1137) | **Medium** | Container ID length violates system limits on Apple platforms, likely preventing proper deployment or execution on macOS/iOS. | No fix PR yet. |
| *(Implicit)* Large image → context overflow | **High** | Addressed by PR #1138 – currently under review. | Yes (#1138) |
| *(Implicit)* Stringified tool args & null params | **Medium** | Addressed by PRs #1136 and #1098 – cause silent failures with local models. | Yes (#1136, #1098) |

## Feature Requests & Roadmap Signals
No explicit feature requests were filed today. However, the pattern of PRs (#1138, #1136, #1098) strongly indicates a roadmap focus on **first‑class support for small/local models** such as Gemma 4 and oMLX. The image‑downscaling fix (#1138) also suggests a planned optimization for reducing token waste, likely to land in the next minor release.

## User Feedback Summary
Direct user feedback is sparse, but the bug report and PR summaries reveal recurring pain points:
- **Apple platform constraints** – Container IDs hitting length limits (Issue #1137).
- **Local model quirks** – Models emitting `null` for optional params and stringified scalars (PRs #1098, #1136).
- **Build dependency bloat** – Unnecessary compilation of `matrix-sdk` when metrics are enabled (PR #1139).

No satisfaction signals (e.g., solved confirmations, thanks) were recorded today.

## Backlog Watch
- **[PR #1098](https://github.com/moltis-org/moltis/pull/1098)** – `fix(browser): tolerate null optional params in browser tool calls` has been open since June 3 (25 days) with no maintainer comments or merge activity. It addresses a real pain point for local models and may need a final review.  
- No other issues or PRs have been left unanswered for an extended period.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

## CoPaw Project Digest – 2026-06-28

### 1. Today's Overview
Project activity today is moderate, with 2 open issues and 6 open pull requests updated in the last 24 hours. No new releases were published. The focus remains on backend unit-test coverage for the Agentscope 2.0 migration, with three substantial test PRs (totaling 120+ test cases) still under review. Two user-reported issues highlight connectivity and UI clarity concerns. No PRs were merged today, indicating a review bottleneck that maintainers should address.

### 2. Releases
*No new releases today.*

### 3. Project Progress
No pull requests were merged or closed today. The following open PRs represent active contributions that have not yet been integrated:
- **#5581** – Unit tests for `app-infra` backend layer (31 cases, Agentscope 2.0) – still open.
- **#5422** – Unit tests for `chats` module (38 cases) – still open.
- **#5423** – Unit tests for `crons` module (51 cases) – still open.
- **#5586** – Fix for context compaction threshold ignoring conversation-level model override – still open.
- **#5585** – Add streaming mode for Matrix channel – still open.
- **#5568** – Fix official plugin installation failures on QwenPaw 2.0 – still open.

These PRs represent important foundational improvements and bug fixes that have not yet advanced to merge.

### 4. Community Hot Topics
The two open issues have limited engagement (1 comment each, no reactions). Neither has generated significant discussion. The most noteworthy is:
- **[#5584]** – *Unable to connect custom ascend-vllm model* – User reports that after upgrading from v1.1.7, QwenPaw fails to connect to a custom vLLM endpoint even though configuration tests pass. This suggests a regression in the model connection layer.  
  → [Issue #5584](agentscope-ai/QwenPaw%20Issue%20%235584)
- **[#5583]** – *Chat interface right-side popup default selection background not obvious* – A UI/UX clarity concern from a user.  
  → [Issue #5583](agentscope-ai/QwenPaw%20Issue%20%235583)

### 5. Bugs & Stability
Two bugs surfaced today, with one rated higher severity:

- **High severity**: [#5584] – Connection failure to custom ascend-vllm model. User states v1.1.7 worked; newer versions consistently fail with `openai.APIConnectionError`. This is a regression affecting model connectivity – a core functionality. No fix PR exists yet.
- **Low severity**: [#5583] – UI: default selected element background in chat popup is hard to distinguish. Cosmetic issue.

Additionally, a fix PR exists for a related bug: **[#5586]** addresses a scenario where the context compaction threshold incorrectly uses static configuration instead of the per-conversation model override. This PR (from a first-time contributor) is awaiting review and merge.

### 6. Feature Requests & Roadmap Signals
The only new feature request today is **PR #5585** which adds streaming mode support to the Matrix channel, similar to the Discord integration. This suggests growing interest in real-time multi-platform messaging. If merged, it could extend CoPaw’s channel support.

The sustained push for unit tests (three PRs from the same author) signals a systematic effort to improve test coverage and reliability for the Agentscope 2.0 codebase. These PRs are likely part of a planned sprint (W1–W3) and may be merged together in the next release.

### 7. User Feedback Summary
Users expressed two distinct pain points:
- **Connectivity regression**: A user who successfully used v1.1.7 with custom vLLM models is now blocked after upgrading. This directly impacts usability for advanced users who run custom backends.
- **UI feedback**: A user noted the default selection in the chat popup lacks visual contrast, indicating that recent UI changes may need refinement.

No positive feedback or satisfaction signals were captured in today’s data.

### 8. Backlog Watch
Several open PRs require maintainer review and potential merging:
- **#5568** – Fix for official plugin installation failures (opened 2 days ago) – critical for plugin ecosystem.
- **#5586** – Context compaction threshold fix (opened today) – core runtime correctness.
- **#5585** – Matrix streaming mode (opened yesterday) – feature addition.
- **#5581, #5422, #5423** – Unit test PRs (opened 1–5 days ago) – foundational for code quality.

No issues were flagged as long-unanswered; however, the lack of comments or reactions on the two open issues may indicate that users are awaiting a timely response.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest – 2026-06-28

## 1. Today’s Overview
The ZeroClaw project shows **very high development velocity** today: 50 pull requests were updated (40 open, 10 merged/closed) and 4 issues were updated (3 open, 1 closed). No new releases were published. Activity is concentrated on **channel bug fixes**, **core config improvements**, and several **large feature PRs** (auto-clean, WASM plugin host, SOP cron triggers). The community is quiet on discussion (most items have 0–1 comments), but the sheer number of updated PRs indicates a productive sprint cycle. The project’s health appears strong with rapid bug resolution and ambitious feature work advancing simultaneously.

## 2. Releases
No new releases were published today. The latest stable version remains unchanged.

## 3. Project Progress
Today **10 pull requests were merged or closed** (the list is not shown in the data, but the count is given). Among the **40 open PRs**, several are near completion and have been updated in the last 24 hours, indicating steady progress:

- **Channel fixes**: #8414 (WhatsApp Web model commands), #7858 (localize runtime command replies), #8110 (Discord autocomplete), #8319 (restore agent precheck controls).
- **Config & stability**: #7529 (gateway dashboard URL only when available), #8115 (daemon fail-fast on address in use), #8339 (promote tool-result image markers for native tools), #8350 (cache regex for strip_tags).
- **CI & docs**: #8343 (build release artifacts from feature registry), #8341 (sync docs with workspace layout), #8276 (Scoop manifest register `zerocode.exe`).
- **Features**: #8400 (wire cron triggers into SOP maintenance tick), #8235 (per-agent `prompt_injection_mode` override), #7923 (automatic temporary file cleanup), #8389 (passive WhatsApp group context), #8368 (WASMtime component-model host – marked DO NOT MERGE).

These cover both bug fixes and enhancements. The project’s **feature pipeline is active**, especially around channels, multi-agent config, and plugin execution.

## 4. Community Hot Topics
Engagement today is low: most issues and PRs have 0–1 comments and no reactions. The most notable item is:

- **[Issue #8396 – RFC: Wire-Protocol-First Provider Model](https://github.com/zeroclaw-labs/zeroclaw/issues/8396)** – This is a high-risk RFC proposing a major architectural shift to make the wire API the primary organizing axis. It has 1 comment and is tagged `needs-maintainer-review`. The underlying need is to **reduce confusion between provider backend types** (e.g., OpenAI vs. Azure) and simplify the config model. This could significantly affect future releases.

Among PRs, the **largest and riskiest** are:

- **#8368** (WASMtime host, size:XL, risk:high) – a foundational change to replace Extism with a WASMtime component model. Marked DO NOT MERGE, indicating it’s still experimental.
- **#7923** (auto-clean, size:XL, risk:high) – adds a configurable temporary file cleanup subsystem.
- **#8400** (SOP cron triggers, size:M, risk:high) – wires cron triggers into the maintenance tick, a key feature for SOP workflows.

These items signal **deep interest in plugin isolation, file lifecycle management, and event-driven SOP execution**.

## 5. Bugs & Stability
Several bugs were reported or fixed today. The **closed issue** #7808 (CLI secret prompt no feedback) has been fixed. The remaining bugs in PRs are being actively addressed:

| Bug | Severity | Status | Fix PR |
|-----|----------|--------|--------|
| WhatsApp Web model commands not working for certain channel names | Medium | Fixed in #8414 | #8414 |
| Hard-coded English replies in core runtime commands | Medium | Fixed in #7858 | #7858 |
| Gateway prints dashboard URL even when web dist is missing | Low | Fixed in #7529 | #7529 |
| Discord autocomplete fails in parent-allowlisted threads | Medium | Fixed in #8110 | #8110 |
| Daemon silently fails when gateway address is already in use | **High** | Fixed in #8115 | #8115 |
| Tool-result image markers not promoted to `image_url` for native tool calls | **High** | Fixed in #8339 | #8339 |
| `strip_tags` recompiles regex each call and panics on invalid pattern | **High** (panic risk) | Fixed in #8350 | #8350 |
| Agent precheck config dropped before runtime | Medium | Fixed in #8319 | #8319 |
| Scoop manifest missing `zerocode.exe` | Low | Fixed in #8276 | #8276 |

No new critical bugs were reported today. The project is **aggressively fixing stability issues**, especially those affecting channels and the daemon.

## 6. Feature Requests & Roadmap Signals
Two new feature requests were opened today:

- **[Issue #8415 – Telegram Bot API 10.1 Rich Messages](https://github.com/zeroclaw-labs/zeroclaw/issues/8415)** – User wants better table rendering and rich message support for Telegram. Likely to be implemented in a future channel update.
- **[Issue #8413 – channel-filesystem SOP event source](https://github.com/zeroclaw-labs/zeroclaw/issues/8413)** – Add a filesystem watcher to trigger SOP workflows. This aligns with the existing `cron` and MQTT event sources and could arrive in the next minor release.

Features already in PR pipeline that are likely to land in the next version:
- **Per-agent `prompt_injection_mode` override** (#8235) – enables multi-agent setups with different security modes.
- **Passive WhatsApp group context** (#8389) – improves history handling for WhatsApp groups.
- **SOP cron triggers** (#8400) – production-ready cron for SOP workflows.
- **Notify before context compression** (#7162) – better UX for agents with large contexts.

The **RFC #8396** could reshape the provider model in a future major version, but it is still in early review.

## 7. User Feedback Summary
User feedback today is limited but reveals clear pain points:

- **CLI usability**: Issue #7808 (now closed) highlighted that the secret prompt gives no feedback after pasting – users were confused. Fix is merged.
- **Telegram client compatibility**: Issue #8415 shows a user frustrated that tables are not rendered correctly on Telegram, leading to a feature request for Rich Messages.
- **Workflow automation**: Issue #8413 demonstrates demand for filesystem-triggered SOP workflows, indicating users want to integrate ZeroClaw with file-based pipelines.

Overall, **users are satisfied with rapid bug fixes** (the CLI prompt fix was addressed within 12 days) but **demand richer channel UX** and more event sources.

## 8. Backlog Watch
Several important issues and PRs require maintainer attention or have been open for a while:

- **[Issue #8396 – RFC: Wire-Protocol-First Provider Model](https://github.com/zeroclaw-labs/zeroclaw/issues/8396)** – tagged `needs-maintainer-review`. Filed 1 day ago; no maintainer response yet.
- **[PR #7162 – Notify before context compression](https://github.com/zeroclaw-labs/zeroclaw/pull/7162)** – open since June 3. Updated today but not merged; may be waiting for final review.
- **[PR #7529 – Fix gateway dashboard URL print](https://github.com/zeroclaw-labs/zeroclaw/pull/7529)** – open since June 12. Low risk, small fix; should be mergeable.
- **[PR #7923 – Auto-clean](https://github.com/zeroclaw-labs/zeroclaw/pull/7923)** – open since June 18. Large feature, risk:high. Needs careful review.
- **[PR #8030 – Doctor warn on OpenAI Codex profile/slot mismatch](https://github.com/zeroclaw-labs/zeroclaw/pull/8030)** – open since June 19. Medium risk; could improve diagnostics.
- **[PR #8350 – Cache strip_tags regex](https://github.com/zeroclaw-labs/zeroclaw/pull/8350)** – tagged `needs-author-action` – awaiting author response.

No long-unanswered issues (the oldest open issues are from June 16–28). The project’s **maintainer bandwidth may be stretched** given the high volume of PRs, but response times appear reasonable.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*