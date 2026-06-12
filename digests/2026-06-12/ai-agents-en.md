# OpenClaw Ecosystem Digest 2026-06-12

> Issues: 404 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-12 06:22 UTC

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

# OpenClaw Project Digest – 2026-06-12

## 1. Today's Overview

OpenClaw sees exceptionally high activity today with **404 issues** and **500 pull requests** updated in the last 24 hours. Of those, 253 issues were closed and 142 PRs were merged or closed, indicating a strong release and fix push. A new beta release **v2026.6.6-beta.2** was cut, focused on hardening security boundaries across a wide surface. The community remains highly engaged, with several long-standing feature requests and regressions drawing concentrated discussion. Overall project health appears strong, though the sheer volume of open items (151 open issues, 358 open PRs) suggests the maintainers are working through a significant backlog.

## 2. Releases

**v2026.6.6-beta.2** – released today (2026-06-12)

- **Highlights**: Security boundaries are substantially tighter across transcripts, sandbox binds, host environment inheritance, MCP stdio, Codex HTTP access, native search policy, elevated sender checks, deleted-agent ACP bypasses, loopback tools, Discord moderation, and Teams group actions.

No detailed changelog or migration notes were provided in the data, but given the sweeping security fixes, users should test this beta in isolated environments before upgrading production instances. Breaking changes are likely for any custom deployments that relied on the now-tightened surfaces (e.g., sandbox escapes, MCP stdio, loopback tools).

## 3. Project Progress

Today **142 pull requests** were closed or merged. Notable among them:

- **fix: initialize execSecurity from config on cold gateway startup** (#92336) – resolves a regression where `tools.exec.security` set to `"full"` was ignored after gateway restart.
- **fix(webchat): ensure senderLabel normalization** (#92331) – fixes empty sender attribution in inter-session messages.
- **Autofix: add PR review autofix pipeline + Windows daemon** (#68936) – merged infrastructure for automated PR review comment resolution and gateway supervision on Windows.
- **fix(codex): wait for migrated plugin inventory** (#82160) – prevents Codex from failing transiently during plugin migration due to incomplete discovery.
- **fix(subagent): preserve steered task text on restart redispatch** (#77539) – prevents loss of task context when a sub‑agent restarts.
- **feat(feishu): handle VC meeting invites** (#89751, superseded by #92340) – adds support for Feishu voice‑call meeting events.

Several critical bug fixes also landed in open PRs today (see §5), though they remain open pending final review.

## 4. Community Hot Topics

The most active issues (by comment count) reveal key community concerns:

- **#9443 – Prebuilt Android APK releases** (25 comments, 👍2)  
  *User Lysen’s AI assistant requests compiled APKs for the Android companion app. Strong demand for simpler mobile deployment.*  
  [openclaw/openclaw#9443](https://github.com/openclaw/openclaw/issues/9443)

- **#52875 – session_send gives “no session found” regression** (22 comments, 👍1)  
  *After upgrading to 2026-3-22, agents lose ability to contact each other. Indicates a significant session‑routing regression.*  
  [openclaw/openclaw#52875](https://github.com/openclaw/openclaw/issues/52875) *(closed)*

- **#32473 – control UI requires device identity (HTTPS/localhost)** (17 comments, 👍5)  
  *Users on VPS/Docker with custom domains cannot access the control UI due to a missing secure‑context requirement.*  
  [openclaw/openclaw#32473](https://github.com/openclaw/openclaw/issues/32473)

- **#90083 – OpenAI ChatGPT transport fails with `invalid_provider_content_type`** (16 comments, 👍3)  
  *After upgrading to 2026.6.1, OpenAI gpt‑5.4/5.5 inference breaks. A provider compatibility issue.*  
  [openclaw/openclaw#90083](https://github.com/openclaw/openclaw/issues/90083) *(closed)*

- **#57901 – Safeguard compaction ignores compaction.model config** (14 comments, 👍1)  
  *When `compaction.mode: "safeguard"` is set with a custom model, the safeguard extension always uses the session’s main model instead.*  
  [openclaw/openclaw#57901](https://github.com/openclaw/openclaw/issues/57901)

**Underlying needs:**  
- Easier mobile/Android deployment (APK)  
- Reliable inter‑agent session management  
- Better HTTPS/identity handling for non‑localhost deployments  
- Robust OpenAI model compatibility after upgrades  
- Correct configuration inheritance for compaction and other features

## 5. Bugs & Stability

The following active bugs are ranked by severity (P1 = highest):

| Issue | Severity | Summary | Fix PR? |
|-------|----------|---------|---------|
| [#57326](https://github.com/openclaw/openclaw/issues/57326) | P1 | CLI-backed helper paths still bypass CLI dispatch on latest main – security bypass risk. | No open PR yet |
| [#32473](https://github.com/openclaw/openclaw/issues/32473) | P2 (regression) | Control UI requires secure context (HTTPS/localhost) – blocks remote admin. | No open PR yet |
| [#57901](https://github.com/openclaw/openclaw/issues/57901) | P2 | Safeguard compaction ignores `compaction.model` config. | No open PR yet |
| [#30381](https://github.com/openclaw/openclaw/issues/30381) | P2 | chatCompletions validates model even when `x-openclaw-agent-id` header is present – blocks BYO model names. | No open PR yet |
| [#41165](https://github.com/openclaw/openclaw/issues/41165) | P1 | Telegram DMs can still pollute the main session after previous fixes. | No open PR yet |
| [#34528](https://github.com/openclaw/openclaw/issues/34528) | P2 | Feishu reaction `message_id` with suffix causes 400 errors on API calls. | No open PR yet |
| [#64500](https://github.com/openclaw/openclaw/issues/64500) | P2 | Circuit breaker only blocks individual tools, not pair loops – ping‑pong attacks survive. | No open PR yet |

Additionally, fix PRs are open for several issues:
- [#92336](https://github.com/openclaw/openclaw/pull/92336) fixes execSecurity config not applied on cold start.  
- [#92358](https://github.com/openclaw/openclaw/pull/92358) fixes plugin method scope resolution in gateway mode.  
- [#92347](https://github.com/openclaw/openclaw/pull/92347) adds model.usage diagnostic events from cron runs.

A cluster of closed issues today addressed stale bugs: Feishu channel config validation (#63101), proactive compaction scheduler lock (#63892), main session prompt crash on undefined length (#63612), and Slack container sandbox attachment failures (#63905).

## 6. Feature Requests & Roadmap Signals

Top enhancement requests (by comment count / reaction count) suggest the community is pushing for:

- **Prebuilt Android APK** (#9443, 25 comments, 👍2) – high demand for mobile companion deployment.
- **Dynamic model discovery** (#10687, 9 comments, 👍3) – users want OpenRouter model catalogs to be automatically refreshed, not statically compiled.
- **Slack Block Kit support** (#12602, 13 comments) – richer messaging in Slack beyond plain text.
- **Pre‑response enforcement hooks** (#13583, 11 comments, 👍2) – mechanical gating for finance/security workflows.
- **Memory trust tagging by source** (#7707, 8 comments) – prevent memory poisoning from untrusted web content.
- **Slack tool‑level progress in threads** (#33413, 8 comments, 👍3) – agents should indicate which tool is running.
- **Automated session memory preservation & synthesis** (#40418, 7 comments) – retain context across `/new` resets.

**Predictions for next version (v2026.6.x stable):**  
- Dynamic model discovery (#10687) has an open PR (#92247) that improves model list loading, so this could land soon.  
- Backup/restore utility (#13616) and native secrets management (#13610) are high‑value but lack active PRs.  
- Slack tool progress (#33413) and Block Kit (#12602) are repeatedly requested but no PRs exist.

## 7. User Feedback Summary

Real user pain points captured in today’s activity:

- **“Upgrade broke my agents”** – Multiple regressions after upgrading to 2026.3.22 (#52875) and 2026.6.1 (#90083, #91016). Users report lost session communication, invalid provider content types, and broken prompt caching costing ~$6/hour (DeepSeek users).
- **“Can’t run mobile without building from source”** – Android APK request (#9443) from a user’s AI assistant highlights a gap in developer experience for mobile usage.
- **“Config validation too strict after upgrade”** – Feishu (#63101) and Telegram (#62985) config errors after upgrading from v4.5 to v4.8.
- **“Circuit breaker doesn’t stop looping”** – #64500 shows users are hitting tool‑loop exploits despite the protection.
- **“Session state is fragile”** – Issues like #57326 (CLI bypass), #57901 (compaction model), and #41165 (Telegram session pollution) indicate that session isolation and state management remain weak points.
- **“Control UI dead on VPS without HTTPS”** – #32473 has 5 upvotes and 17 comments, signaling many remote users are blocked from managing their instances.

Overall satisfaction is mixed: the project is highly active and responsive, but the volume of regressions in recent releases is frustrating users who rely on stability for production workflows.

## 8. Backlog Watch

Several important issues have gone unanswered or remain open for months, needing maintainer attention:

| Issue / PR | Created | Last Update | Status |
|------------|---------|-------------|--------|
| [#9443](https://github.com/openclaw/openclaw/issues/9443) – Prebuilt Android APK | 2026-02-05 | 2026-06-12 | Open, 25 comments, no maintainer response |
| [#7707](https://github.com/openclaw/openclaw/issues/7707) – Memory trust tagging | 2026-02-03 | 2026-06-12 | Open, 8 comments, needs‑security‑review |
| [#6615](https://github.com/openclaw/openclaw/issues/6615) – Denylist for exec‑approvals | 2026-02-01 | 2026-06-12 | Open, 7 comments, needs‑decision |
| [#11665](https://github.com/openclaw/openclaw/issues/11665) – Webhook session reuse | 2026-02-08 | 2026-06-12 | Open, 8 comments, linked PR open but stalled |
| [#12602](https://github.com/openclaw/openclaw/issues/12602) – Slack Block Kit | 2026-02-09 | 2026-06-12 | Open, 13 comments, needs‑decision |
| [#13616](https://github.com/openclaw/openclaw/issues/13616) – Backup/restore utility | 2026-02-10 | 2026-06-12 | Open, 8 comments, needs‑decision |
| [#13583](https://github.com/openclaw/openclaw/issues/13583) – Pre‑response enforcement hooks | 2026-02-10 | 2026-06-12 | Open, 11 comments, needs‑security‑review |
| [#63330](https://github.com/openclaw/openclaw/pull/63330) – Session followup turn API (PR) | 2026-04-08 | 2026-06-12 | Stale, needs‑real‑behavior‑proof |
| [#82434](https://github.com/openclaw/openclaw/pull/82434) – External plugin approval verification (PR) | 2026-05-16 | 2026-06-12 | Stale, P2, large merge‑risk |

Several of these items (especially #9443, #7707, #6615, #12602) have been open for over four months with no maintainer comment, indicating a gap in community feature integration. The PR #63330 adds a powerful plugin API but has been waiting for real‑world proof since April.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: Personal AI Assistant Open-Source Ecosystem

**Date:** 2026-06-12  
**Prepared for:** Technical decision-makers and developers evaluating open-source AI agent infrastructure

---

## 1. Ecosystem Overview

The personal AI assistant open-source landscape is in a phase of rapid maturing, driven by two parallel forces: security hardening after high-profile exploits, and architectural convergence toward multi-agent, MCP-enabled, and sandboxed runtimes. Projects are increasingly competing on platform portability, with WSL2, macOS, and ARM support emerging as differentiators rather than nice-to-haves. The community is demanding **production-grade reliability**—context persistence, session isolation, and provider compatibility—over raw feature velocity, as evidenced by the surge of regression-related issue reports across nearly every major project. A clear pattern is emerging: the ecosystem is consolidating around modular, gateway-centric architectures that decouple agent logic from channel delivery, with MCP acting as the universal integration protocol.

---

## 2. Activity Comparison

| Project | Issues Updated (24h) | PRs Updated (24h) | Release (24h) | Health Score | Notes |
|---|---|---|---|---|---|
| **OpenClaw** | 404 (253 closed) | 500 (142 merged) | v2026.6.6-beta.2 | **High Flux / Responsive** | Massive scale; security-focused beta |
| **NanoBot** | 4 (1 closed) | 17 (4 merged) | None | **Active / Moderate** | Focused additions; few regressions |
| **Hermes Agent** | 14 (14 closed) | 50 (27 merged) | None | **Very Healthy** | All bugs closed today; fast turnaround |
| **PicoClaw** | 10 (0 closed) | 35 (17 merged) | v0.2.9-nightly | **High Activity / Fragile** | Nightly build; Gemini regression |
| **NanoClaw** | 1 (1 closed) | 13 (9 merged) | None | **Active / Stable** | Signal integration gap |
| **NullClaw** | 1 (0 closed) | 0 | None | **Low / Dormant** | Single open bug; no PR activity |
| **IronClaw** | 19 (3 closed) | 40 (18 merged) | None | **Intense Dev / Stabilizing** | Reborn focus; nightly E2E failing |
| **LobsterAI** | 2 (1 closed) | 19 (13 merged) | None | **Active / Productive** | 13 merged fixes; ASR feature |
| **TinyClaw** | 0 | 0 | None | **Inactive** | No activity |
| **Moltis** | 1 (0 closed) | 1 (0 merged) | None | **Low / Early** | Two open items; no PR merged |
| **CoPaw** | 23 (3 closed) | 37 (16 merged) | v1.1.11.post1/post2 | **High Flux / Fragile** | Desktop regressions in v1.1.11 |
| **ZeptoClaw** | 0 | 0 | None | **Inactive** | No activity |
| **ZeroClaw** | 23 (3 closed) | 50 (3 merged) | v0.8.0 (prior) | **High Flux / Post-Launch** | Major multi-agent release regressions |

**Health Score Definitions:**  
- **Very Healthy** – High activity, low open bugs, fast closure  
- **Active / Stable** – Moderate activity, few breaking issues  
- **High Flux / Fragile** – High activity, but significant open regressions  
- **Low / Dormant** – Minimal change, no development momentum  

---

## 3. OpenClaw's Position

**Advantages over peers:**
- **Community scale**: 404 daily issues and 500 daily PRs dwarf every other project—OpenClaw's contributor base is 4–10× larger than its nearest competitor (IronClaw: 40 PRs; ZeroClaw: 50 PRs).
- **Security maturity**: The v2026.6.6-beta.2 release explicitly hardens transcripts, sandbox binds, MCP stdio, loopback tools, and multiple channel-specific attack surfaces—a level of coordinated security response unmatched elsewhere.
- **Reference implementation status**: OpenClaw is the clear "upstream" reference project; other projects (LobsterAI, PicoClaw, NanoClaw) explicitly fork or integrate its architecture.

**Technical approach differences:**
- OpenClaw's architecture is **monolithic gateway** with plugin support, while many peers (ZeroClaw, NanoBot, IronClaw) are moving toward **multi-agent daemon** or **microservice** patterns.
- Strongest **platform coverage**—Linux, Windows, macOS, Android companion—though Android APK packaging remains an open community demand (#9443).
- **Configuration inheritance** bugs (compaction model, execSecurity) suggest complexity in its layered config model; simpler projects (NanoBot, Moltis) have fewer such issues.

**Community size comparison:**  
OpenClaw's raw activity is 5–10× higher than the next-most-active projects (ZeroClaw, IronClaw, CoPaw). However, the sheer volume also means more regressions per release. Hermes Agent, while smaller, shows a higher **closure velocity**—100% of issues closed today.

---

## 4. Shared Technical Focus Areas

The following requirements are emerging independently across multiple projects, signaling ecosystem-wide priorities:

| Focus Area | Projects Affected | Specific Needs |
|---|---|---|
| **Multi-agent orchestration** | ZeroClaw, OpenClaw, PicoClaw, CoPaw, IronClaw | Subagent model presets, inter-agent mailboxes, session-scoped automation, spawn delegation |
| **MCP stability & extensibility** | OpenClaw, NanoBot, PicoClaw, ZeroClaw, Moltis | Stdio security, auto-reconnect, per-request dynamic headers, config persistence |
| **Platform portability (Windows/macOS/ARM)** | PicoClaw, ZeroClaw, CoPaw, IronClaw | WSL2 OOM, macOS Homebrew breakage, Windows `list_dir` path bugs, Tauri CI |
| **Memory & context management** | OpenClaw, Hermes, CoPaw, NanoClaw | Compaction model inheritance, orphaned tool results, long-term memory vector config |
| **Security hardening** | OpenClaw, IronClaw, CoPaw, ZeroClaw | Keychain isolation, sandbox governance, circuit breaker loops, approval callback hooks |
| **Provider compatibility** | OpenClaw, ZeroClaw, CoPaw, PicoClaw, NullClaw | Gemini `thought_signature`, OpenAI `invalid_provider_content_type`, Ollama truncation |
| **Session state reliability** | OpenClaw, NanoBot, LobsterAI, CoPaw | Session isolation after restart, non-blocking input queues, draft persistence |

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | Hermes | ZeroClaw | CoPaw | IronClaw |
|---|---|---|---|---|---|---|
| **Primary user** | Advanced operators / self-hosters | Developer integrators | Plugin developers | Multi-agent operators | Desktop power users | Enterprise teams |
| **Architecture** | Monolithic gateway + plugins | Modular microservice | Plugin SDK + gateway | Multi-agent daemon | Desktop app + backend | Reborn microservice (v2) |
| **Key strength** | Scale / security / reference | MCP extensibility | Bug closure speed | Multi-agent model | UI polish / desktop | Configuration-as-Code vision |
| **Key weakness** | Config complexity / regressions | Small community | Windows gaps | Post-release stability | Desktop regressions | E2E CI broken |
| **Platform focus** | All (Linux first) | Linux / Docker | Linux / Windows | Linux / WSL2 | Windows / macOS | Linux / macOS |
| **AI provider support** | Broad (OpenAI, Anthropic, Gemini, local) | Broad (custom + OpenAI) | Focused (Anthropic + OpenAI) | Broad + Gemini | Chinese providers + OpenAI | Broad + ChatGPT device flow |
| **Community response** | Fast but overwhelmed | Moderate | Very fast | Post-release firefighting | Rapid patch cycle | Fast on bugs, slow on RFCs |

---

## 6. Community Momentum & Maturity

**Tier 1 – High Momentum, Rapid Iteration**  
- **OpenClaw** – Massive community, but scaling pains. Merging ~140 PRs/day indicates strong maintainer bandwidth, yet 358 open PRs show backlog pressure.  
- **ZeroClaw** – Post-major-release surge. Multi-agent paradigm shift attracting new users, but WSL2/macOS regressions threaten retention.  
- **CoPaw** – Very active, especially on desktop. The AgentScope 2.0 migration (#4727) will be a pivotal fork in the road.  
- **Hermes Agent** – Smallest in this tier, but **best closure metrics** (all issues closed today). Model for high-velocity, low-regret development.

**Tier 2 – Stabilizing / Maturing**  
- **IronClaw** – Deep focus on Reborn hardiness. 19 open issues but rapid PR pairing (e.g., #4789 → #4790 within hours). Moving toward production cutover.  
- **LobsterAI** – Productive day (13 merges), but the Cowork module seems to be the only active area. Chinese-language community feedback is a strength.  
- **NanoBot** – Adding features (cron binding, subagent models) but no release cadence. Medium risk of scope creep.  
- **PicoClaw** – Nightly releases suggest continuous delivery, but the Gemini regression (#3111) is a critical on-ramp blocker for that provider's users.

**Tier 3 – Low Activity / Dormant**  
- **NanoClaw** – Only 1 issue, 13 PRs. Small but focused; Signal integration is the main gap.  
- **Moltis** – Two open items, no merged PRs. Early-stage or maintainer-limited.  
- **NullClaw, TinyClaw, ZeptoClaw** – Essentially inactive. Not suitable for active development needs.

---

## 7. Trend Signals

**1. Multi-agent is the new baseline.**  
ZeroClaw's v0.8.0 (daemon managing many named agents) and PicoClaw's agent collaboration bus PR (#2937) signal that single-agent architectures are being retired. Developers should expect subagent delegation, hierarchical model presets, and session-scoped automation to become table stakes within 6 months.

**2. Security is moving from perimeter to runtime.**  
OpenClaw's beta hardens 10+ attack surfaces. IronClaw adds approval callback hooks. CoPaw isolates keychain per install. The trend is toward **defense-in-depth inside the agent loop**—circuit breakers, pre-response enforcement hooks, and sandbox governance—not just network-level firewalls.

**3. Platform portability is the #1 user frustration.**  
Every project with significant desktop/CLI usage (ZeroClaw, CoPaw, PicoClaw, IronClaw) has at least one critical platform-specific bug open. Windows users face path separator bugs, macOS users get blank dashboards, WSL2 users hit OOMs. **Cross-platform E2E testing is the largest infrastructure gap** across the ecosystem.

**4. MCP is becoming the universal plumbing.**  
MCP issues span OpenClaw (stdio security), NanoBot (reconnect crashes), PicoClaw (dynamic headers), ZeroClaw (auto-reconnect), and Moltis (auth failures). The protocol is winning, but its implementation surface is still fragmented. Expect convergence on a standardized MCP governance layer in coming months.

**5. "Provider lock-in" is being actively dismantled.**  
Multiple projects now support provider-agnostic abstraction (NanoBot's custom provider PR, ZeroClaw's provider-switch tool, OpenClaw's model overlay). Users are demanding the ability to route models per-task, per-agent, or per-channel. **Multi-provider, multi-model routing is the winning architecture.**

**6. Community satisfaction is inversely correlated with release frequency.**  
Projects that ship rapidly (OpenClaw beta, CoPaw post-patches, ZeroClaw v0.8.0) see the most regressions and complaint volume. Projects with slower, more gated releases (Hermes, NanoBot) have fewer complaints. **Value of a stable, well-tested release outweighs feature velocity** for production users.

---

**Key Recommendation for Decision-Makers:**  
- If you need **broadest ecosystem and security maturity**: OpenClaw  
- If you need **fast bug fixes and plugin extensibility**: Hermes Agent  
- If you need **multi-agent out of the box**: ZeroClaw (with macOS caveats)  
- If you need **desktop-first UX with Chinese provider support**: CoPaw  
- If you need **MCP-first, lightweight integration**: NanoBot  

All projects would benefit from **cross-platform test suites** and **provider-agnostic memory/context utilities**—the two largest shared gaps identified in today's data.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-06-12

## 1. Today's Overview
NanoBot saw heavy development activity in the past 24 hours, with **4 issues updated** (3 open, 1 closed) and **17 pull requests updated** (13 open, 4 merged/closed). The project is in a phase of active feature additions and bug fixes, particularly around **MCP stability**, **session management**, and **provider extensibility**. No new releases were published, but several high-impact PRs are approaching merge. The overall health appears robust, with a responsive maintainer team addressing both community-reported bugs and long-standing feature requests.

## 2. Releases
No new releases are recorded in the data.

## 3. Project Progress
Four pull requests were **merged or closed** today:
- **[#4020 – feat(providers): make stream-idle timeout configurable per-provider](https://github.com/HKUDS/nanobot/pull/4020)** – *Closed*  
  Introduces a per-provider `stream_idle_timeout` setting, moving beyond the single env-var knob. Helps local LLM users avoid premature stream aborts.

- **[#4289 – feat(slack): add groupRequireMention to scope allowlist channels to @mentions](https://github.com/HKUDS/nanobot/pull/4289)** – *Closed*  
  Adds a Slack channel option to require @mention even in allowlisted channels, giving teams finer control over bot responses.

- **[#4297 / #4298 – Worktree feature + Hermes research doc](https://github.com/HKUDS/nanobot/pull/4297)** – *Closed* (duplicate PRs)  
  Internal documentation additions (likely non-functional changes).

## 4. Community Hot Topics
No single issue or PR drew exceptional comment or reaction volume, but **a few topics stand out**:

- **[Issue #4305 – Multiple custom providers](https://github.com/HKUDS/nanobot/issues/4305)**  
  *1 comment, 0 reactions*  
  User requests the ability to configure more than one "custom" or "openai" provider. This is a recurring pain point; an older PR [#3239](https://github.com/HKUDS/nanobot/pull/3239) already attempts to address it and has seen recent updates.

- **[PR #4299 – Bind scheduled automations to sessions](https://github.com/HKUDS/nanobot/pull/4299)**  
  A substantial feature that ties cron automations to specific sessions and defers execution until the target session is idle. No comments but multiple commits indicate active discussion among maintainers.

- **[PR #4291 – Allow subagents to use configurable model presets](https://github.com/HKUDS/nanobot/pull/4291)**  
  Enables hierarchical model selection – subagents can run with a different model than their parent. Another feature that directly addresses user requests for more flexible multi-agent setups.

**Underlying need**: Users want **flexible, configurable provider topologies** (multiple custom endpoints, per-subagent models, session-scoped cron) to support complex production deployments and heterogeneous model usage.

## 5. Bugs & Stability
Three bugs were reported today, ranked by severity:

| Severity | Issue | Description | Fix PR Exists? |
|----------|-------|-------------|----------------|
| **High** | [#4307](https://github.com/HKUDS/nanobot/issues/4307) | Post-turn consolidation wipes the assistant’s own delivery message, losing follow-up references. Can corrupt multi-iteration conversations when `context_window_tokens` is set modestly. | No |
| **Medium** | [#4302](https://github.com/HKUDS/nanobot/issues/4302) | Gateway crashes when MCP server reconnects after session termination due to task cancellation scope mismatch. | Yes ([#4303](https://github.com/HKUDS/nanobot/pull/4303)) |
| **Medium** | [#4306](https://github.com/HKUDS/nanobot/issues/4006) – *reported earlier, fix PR active* | Orphaned `role:"tool"` messages in session history cause API rejections. Fix PR [#4306](https://github.com/HKUDS/nanobot/pull/4306) (same number) prevents persistence of unmatchable tool results. | Yes ([#4306](https://github.com/HKUDS/nanobot/pull/4306)) |

Additionally, the **bwrap sandbox bug** ([#4236](https://github.com/HKUDS/nanobot/issues/4236)) on Ubuntu 24.04 has been closed, indicating a fix or workaround has been applied.

## 6. Feature Requests & Roadmap Signals
User-requested features visible in today’s data:

- **Multiple custom providers** ([#4305](https://github.com/HKUDS/nanobot/issues/4305), PR [#3239](https://github.com/HKUDS/nanobot/pull/3239)) – Likely to land soon, as the PR has been updated today.
- **Cron automation bound to sessions** ([#4299](https://github.com/HKUDS/nanobot/pull/4299)) – A significant architectural change; may ship in next minor release.
- **Subagent model presets** ([#4291](https://github.com/HKUDS/nanobot/pull/4291)) – Enhances spawn tool flexibility.
- **Python SDK expansion** ([#4296](https://github.com/HKUDS/nanobot/pull/4296)) – Adds richer `RunResult`, session control, and runtime metadata.
- **Skills caching** ([#4301](https://github.com/HKUDS/nanobot/pull/4301)) – Performance improvement for skill-loading hot paths.

Prediction: The **next NanoBot release** will likely include multiple custom providers, cron session binding, and subagent model presets, as these are mature PRs with recent maintainer activity.

## 7. User Feedback Summary
Real-world pain points and use cases surfaced by the community:

- **“I need more than one custom provider”** – Users running multiple internal or third-party OpenAI-compatible APIs have no clean configuration path.
- **“bwrap sandbox fails on Ubuntu 24.04”** – Now closed, but indicates the challenge of supporting restricted user namespaces.
- **“Gateway crashes on MCP reconnect”** – Stability concern for production deployments using streamable HTTP MCP servers.
- **“Post-turn consolidation wipes assistant messages”** – Severe data loss that breaks multi-turn workflows; reported with detailed steps.
- **“Orphaned tool results pollute history”** – API compatibility and rendering issues when tool calls and results are mismatched.
- **“Subagents should use different models”** – Power users want to delegate expensive reasoning tasks to cheaper/faster models.
- **“Slack channel should require @mention in allowlist mode”** – Teams want to prevent unintended bot responses in noisy channels.

Overall satisfaction is moderate – the project is adding wanted features, but the volume of bug reports suggests stability improvements are still a priority.

## 8. Backlog Watch
The following open issues/PRs have been unanswered or have seen no maintainer response for an extended period:

- **[PR #3239 – Support multiple custom OpenAI-compatible providers](https://github.com/HKUDS/nanobot/pull/3239)**  
  Opened April 17, updated today. While not stale, it has been in review for nearly two months. Active today, but still unmerged.

- **[PR #3538 – Add gateway start/stop/restart commands](https://github.com/HKUDS/nanobot/pull/3538)**  
  Opened April 29, last updated June 11. Introduces basic daemon management but has no maintainer comments.

- **[PR #4021 – Dedup reasoning items before send (Codex)](https://github.com/HKUDS/nanobot/pull/4021)**  
  Opened May 27, updated June 11. Fix for duplicate reasoning item errors in OpenAI Responses API – important for Codex users, but not merged.

- **[Issue #4006 – Orphaned tool results](https://github.com/HKUDS/nanobot/issues/4006)** – *Opened long before, now has a fix PR but the issue itself may still need closure.*  

These items deserve maintainer attention to either merge, close, or provide status updates.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest – 2026-06-12

## Today's Overview

The Hermes Agent project exhibited **very high activity** on June 12, 2026, with **50 pull requests** updated in the last 24 hours (23 open, 27 merged/closed) and **14 issues** closed during the same window. **No new releases** were published, but the near‑instant resolution of all updated issues (0 still open) and a flurry of merged/closed PRs indicate the maintainers are aggressively clearing the backlog. The focus was on stability: the closed issues overwhelmingly represent bug reports (12 out of 14) that were already marked as `sweeper:implemented-on-main`, suggesting fixes were landed on the main branch. Additionally, several open PRs introduce new features (e.g., WSL2 execution backend, rate‑limit guard hooks), pointing to ongoing forward‑looking development alongside the cleanup push.

## Releases

*No new releases were published today.* The most recent release remains v0.13.0 (tagged 2026-05-07).

## Project Progress

**27 pull requests** were merged or closed in the last 24 hours. Key changes include:

| PR | Description | Impact |
|----|-------------|--------|
| [#44676](https://github.com/NousResearch/hermes-agent/pull/44676) | Strip `default` from `$ref` nodes in tool schemas | Fixes HTTP 400 errors with strict backends (Fireworks‑hosted Kimi K2) |
| [#25367](https://github.com/NousResearch/hermes-agent/pull/25367) | Exclude live integration tests from PR CI | Improves CI safety for non‑hermetic tests |
| [#25340](https://github.com/NousResearch/hermes-agent/pull/25340) | Graceful fallback when no Win32 console buffer | Enables Hermes to run in Windows headless services |
| [#25298](https://github.com/NousResearch/hermes-agent/pull/25298) | Preserve systemd unit paths when service user differs from CLI user | Prevents broken unit files after `hermes update` |
| [#25265](https://github.com/NousResearch/hermes-agent/pull/25265) | Omit failed commands from session summaries | Prevents broken tool calls from poisoning summaries |
| [#25261](https://github.com/NousResearch/hermes-agent/pull/25261) | Trigger fallback on 429 rate‑limit and billing errors | Enables automatic fallback when primary provider is exhausted |
| [#25258](https://github.com/NousResearch/hermes-agent/pull/25258) | Fix duplicate Slack final replies after stream finalize failure | Eliminates double messages in Slack |
| [#25236](https://github.com/NousResearch/hermes-agent/pull/25236) | Restore example‑dashboard plugin's `dist/index.js` | Re‑adds a vital SDK reference bundle |

Several other fixes landed for plugin hooks (#25204), Telegram polling (#25221), Mattermost threading (#25181), vision tool with Telegram (#25118), and more.

## Community Hot Topics

The most active issues (by comment count) all saw **4 comments** each and were resolved today:

- **[Issue #25378](https://github.com/NousResearch/hermes-agent/issues/25378)** – Azure OpenAI configuration not working (404 error). The author reported a `NotFoundError` and asked for a workaround. The issue was closed with `sweeper:implemented-on-main`, suggesting a fix was committed.
- **[Issue #25281](https://github.com/NousResearch/hermes-agent/issues/25281)** – Dashboard "Update" button deletes scheduled cron jobs. The reporter mentioned this happened multiple times and requested persistent storage for cron jobs. A fix was implemented on main.
- **[Issue #25351](https://github.com/NousResearch/hermes-agent/issues/25351)** – Dashboard chat breaks behind reverse proxies due to 15‑second `npm run build` blocking the event loop. The bug description included a detailed root‑cause analysis and impact on WebSocket keepalives. Closed with fix.
- **[Issue #25191](https://github.com/NousResearch/hermes-agent/issues/25191)** – Gateway install issue on Windows (emoji in code causing failure). The issue was closed with a fix.

These issues reveal a **desire for robustness in production deployments** (reverse proxies, cron persistence, Windows support) and **low friction with cloud provider integrations** (Azure OpenAI).

## Bugs & Stability

All **14 issues** updated in the last 24 hours were closed (0 open). Severity breakdown (by P‑level assigned in labels):

| Severity | Count | Examples |
|----------|-------|----------|
| **P1 (Critical)** | 4 | Cron deletion (#25281), Gateway memory leak growing to 20–37 GB (#25315), Iterative context compaction overriding active topics (#9631), Telegram dead‑polling state (#25221) |
| **P2 (High)** | 6 | Dashboard chat blocking (#25351), Gateway install failure (#25191), Shell hooks unreliability (#25204), Mattermost threading broken (#25181), Vision tool with Telegram (#25118), `/restart` shutting down Docker containers (#25217) |
| **P3 (Low)** | 4 | Azure OpenAI config (#25378), Kimi‑coding API path error (#25104), Feature request for plugin hooks (#25262), Command display truncated in TUI (#18376) |

**Critical fix deployed:** The gateway memory leak reported in [#25315](https://github.com/NousResearch/hermes-agent/issues/25315) (agent cache not cleaning `_session_messages`) had a corresponding fix PR merged – see [#25298](https://github.com/NousResearch/hermes-agent/pull/25298) and related. The cron‑deletion bug was also resolved.

No regressions were explicitly reported today. The `sweeper:implemented-on-main` label on nearly all closed issues indicates fixes are already in the main branch.

## Feature Requests & Roadmap Signals

One explicit feature request appeared:

- **[#25262](https://github.com/NousResearch/hermes-agent/issues/25262)** – Add `pre/post tool‑call hooks` to the plugin context API. The author works on AgentGuard (security/audit plugin) and needs hooks to inspect tool calls. This was **closed with `sweeper:implemented-on-main`** today, meaning it is already merged.

Additionally, open PRs suggest upcoming features that may land in the next release:

- **[#44158](https://github.com/NousResearch/hermes-agent/pull/44158)** – WSL2 execution backend + per‑backend persistent CWD (Windows users get a native Linux environment without Docker).
- **[#44703](https://github.com/NousResearch/hermes-agent/pull/44703)** – Pre‑send rate‑limit guard hook plus a “CLAUDE‑trio” RFC (borrowing patterns from a popular 182K‑star repo).
- **[#44704](https://github.com/NousResearch/hermes-agent/pull/44704)** – Document `${ENV_VAR}` substitution in MCP `user env:` blocks.

Given that the hook feature (#25262) was already implemented, the current roadmap appears to emphasise **platform compatibility** (WSL2) and **security/auditing** (rate‑limit guard, tool‑call hooks).

## User Feedback Summary

- **Pain point 1 (most common):** Providers and integrations break in non‑standard environments. Users reported issues with Azure OpenAI, Kimi‑coding, Telegram, Mattermost, Slack, and reverse proxies – all were addressed quickly.
- **Pain point 2:** Update process destroys configuration (cron jobs). The community requested persistent storage; the fix ensures cron jobs survive updates.
- **Pain point 3:** Memory / resource leaks in long‑running gateway deployments. The 20–37 GB memory leak was identified and fixed.
- **Satisfaction indicators:** Users filed detailed bug reports with logs and root‑cause analysis (#25351, #25315), indicating a technically engaged community. The feature request for plugin hooks (#25262) was implemented the same day it was filed – a sign of high responsiveness.

## Backlog Watch

No long‑unanswered issues or PRs were observed today; the oldest issue closed today was [#9631](https://github.com/NousResearch/hermes-agent/issues/9631) (filed April 14, 2026 – nearly two months old) – a P1 bug about context compaction overriding active topics. It was finally resolved today. No other dormant items remain visible in the current dataset.

However, several **open PRs** have been pending for weeks without merge:

- **[#12525](https://github.com/NousResearch/hermes-agent/pull/12525)** (April 19) – Verify entity state after Home Assistant service call. P2 bug fix.
- **[#11457](https://github.com/NousResearch/hermes-agent/pull/11457)** (April 17) – Gate OGG conversion on platform for TTS. P2 bug fix.
- **[#11455](https://github.com/NousResearch/hermes-agent/pull/11455)** (April 17) – Mistral API compatibility patches. P3 feature.

These are relatively old (almost two months) and may need maintainer review or conflict resolution. They are not blocked by new issues but have not been merged, possibly due to complexity or lower priority.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest – 2026-06-12

## Today's Overview

The PicoClaw project showed very high activity today with **10 issues** and **35 pull requests** updated in the last 24 hours. A new **nightly build (v0.2.9-nightly)** was released, reflecting steady integration of bug fixes and dependency upgrades. The community remains strongly engaged, with contributions covering Windows cross‑platform fixes, Gemini model compatibility, channel permission scoping, and multi‑agent collaboration infrastructure. The project appears healthy and actively maintained, though several critical bugs require prompt resolution.

## Releases

- **nightly (v0.2.9-nightly.20260612.413d3749)** – Automated nightly build.
  This is an unstable development snapshot and is **not recommended for production use**. No explicit breaking changes or migration notes were provided. Full changelog: [v0.2.9…main](https://github.com/sipeed/picoclaw/compare/v0.2.9...main).

## Project Progress

**17 pull requests were merged or closed today**, including several important fixes:

- **Tool call reliability** – `#2957` fixes `tool_calls` being incorrectly filtered during streaming (closed).
- **Identity & channel parsing** – `#3045` corrects Matrix `allow_from` handling for user IDs containing colons; `#3048` fixes `mcp add` argument parsing when root flags are passed; `#2956` preserves channel enabled state during security.yml merge.
- **Singleton PID check** – `#2955` prevents false‑positive startup failures by verifying process identity.
- **Model ID corrections** – `#2947` fixes the default Claude model ID from `claude-sonnet-4.6` to `claude-sonnet-4-6` (hyphens).
- **WhatsApp native mode** – `#2934` enables native WhatsApp support via the `use_native` flag.
- **MCP dynamic headers** – `#2696` adds support for per‑request HTTP headers forwarded from channel context.
- **Security scoping** – `#3109` introduces channel‑level permission scoping, restricting dangerous operations in group/channel chats (feature merged).
- **Dependency updates** – Multiple bumps by Dependabot: AWS SDK (`#3106`, `#3102`), Go sync ( `#3099`), MCP Go SDK (`#3098`), GitHub Copilot SDK (`#3107`), ESLint (`#3105`), shadcn (`#3104`).

## Community Hot Topics

- **[Issue #2984](https://github.com/sipeed/picoclaw/issues/2984)** – *"Add explicit turn completion signal for Pico WebSocket clients"*  
  2 reactions (👍), 2 comments. External clients need a deterministic signal to know when the agent has finished processing. This reflects a growing demand for robust real‑time protocol semantics.

- **[Issue #2472](https://github.com/sipeed/picoclaw/issues/2472)** – *"[BUG] list_dir returns 'invalid argument' on Windows"*  
  1 reaction, 5 comments. A long‑standing platform compatibility bug that blocks Windows users; still open.

- **[Issue #3111](https://github.com/sipeed/picoclaw/issues/3111)** – *"[BUG] Tool execution fails with Gemini 3.5 Flash"*  
  New bug reported today (0 comments, but high potential impact). The Google API rejects tool calls due to an incompatible schema. Users of Gemini 3.5 Flash are immediately affected.

- **[PR #2937](https://github.com/sipeed/picoclaw/pull/2937)** – *"Feat/agent collaboration"*  
  Still open; introduces a first‑class internal agent collaboration bus with per‑agent mailboxes and collaboration threads. This PR has drawn attention as a major architectural enhancement.

## Bugs & Stability

The following bugs were reported today, ranked by severity:

| Rank | Issue | Summary | Severity | Fix PR exists? |
|------|-------|---------|----------|----------------|
| 1 | [#3111](https://github.com/sipeed/picoclaw/issues/3111) | Tool execution fails with **Gemini 3.5 Flash** due to missing `thought_signature` in schema | **Critical** – blocks all tool usage with that model | No |
| 2 | [#3108](https://github.com/sipeed/picoclaw/issues/3108) | Image description requests **hallucinate** when active model lacks vision support | **High** – produces unrelated text instead of actual description | No |
| 3 | [#3110](https://github.com/sipeed/picoclaw/issues/3110) | Telegram adapter ignores `message_thread_id` in Forum topics – replies go to #General | **High** – breaks Telegram Forum usage | No |
| 4 | [#3094](https://github.com/sipeed/picoclaw/issues/3094) | Asynchronous sub‑agent (`spawn`) tasks cause **duplicate messages** on channels | **Medium** – creates noisy user experience | No |

Additionally, the older Windows path bug (`#2472`) remains unresolved.

## Feature Requests & Roadmap Signals

- **Explicit turn completion signal** ([#2984](https://github.com/sipeed/picoclaw/issues/2984)) – Strong community support; likely to be prioritised for the next minor release to improve WebSocket protocol usability.
- **Agent collaboration bus** ([PR #2937](https://github.com/sipeed/picoclaw/pull/2937)) – Though still open, this major feature is under active development and may be included in a near‑future release.
- **Channel‑level permission scoping** ([#3109](https://github.com/sipeed/picoclaw/issues/3109)) – Already merged, now available in the nightly build.
- **MCP per‑request dynamic headers** ([#2696](https://github.com/sipeed/picoclaw/pull/2696)) – Also merged; enables richer channel ↔ MCP server interactions.

Signals from today’s issues indicate the next version will likely address **Gemini vision reasoning compatibility**, **Telegram Forum support**, and **sub‑agent message deduplication**.

## User Feedback Summary

**Pain points**:
- Windows users are blocked by the `list_dir` path separator bug (still open since April).
- Gemini 3.5 Flash users cannot execute any tools – a showstopper for that provider.
- Telegram Forum users cannot get replies inside the correct thread, making the bot nearly unusable in forum‑style groups.
- Image analysis with text‑only models produces hallucinated descriptions, misleading users.
- Sub‑agent tasks generate duplicate messages, cluttering channels.

**Positive signals**:
- The community is actively contributing fixes (e.g., Matrix identity, WhatsApp mode, tool call streaming).
- The project’s fast pace of dependency updates and security patches (e.g., allowed_cidrs bypass fix #3080) indicates strong maintainer responsiveness.
- Features like permission scoping and MCP dynamic headers are being promptly merged.

## Backlog Watch

| Item | Opened | Last Update | Status |
|------|--------|-------------|--------|
| [Issue #2472](https://github.com/sipeed/picoclaw/issues/2472) – Windows path separator bug | 2026-04-10 | 2026-06-11 | **Open, needs resolution** – has 5 comments, no assignee |
| [Issue #2984](https://github.com/sipeed/picoclaw/issues/2984) – Turn completion signal | 2026-06-02 | 2026-06-12 | **Open, 2 reactions** – awaiting maintainer decision |
| [PR #2937](https://github.com/sipeed/picoclaw/pull/2937) – Agent collaboration bus | 2026-05-24 | 2026-06-11 | **Open, no recent activity** – major feature needs attention |
| [Issue #3094](https://github.com/sipeed/picoclaw/issues/3094) – Duplicate messages from spawn | 2026-06-10 | 2026-06-11 | **Open, no fix yet** – impacts real users |

These items represent the most pressing gaps in the project’s stability and feature completeness. Maintainers should prioritise the Windows compatibility bug and the Gemini 3.5 Flash regression to avoid alienating key user segments.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-12

## Today’s Overview

The NanoClaw project saw a surge of activity today, with 13 pull requests updated in the last 24 hours — including 9 merged/closed and 4 still open. A single bug issue (#2495) was closed after a corresponding fix PR (#2738) was merged. The remaining open issue (#1356) continues to track the design of a scalable agent memory system, having accumulated 6 👍 reactions. No new releases were published. Overall, the project is in a healthy state with a strong focus on stability fixes and infrastructure features, though the large batch of merged PRs today suggests a coordinated push ahead of an upcoming release.

## Releases

*None to report. No new releases were created in the last 24 hours.*

---

## Project Progress – Merged/Closed PRs

Nine pull requests were merged or closed today, spanning fixes, feature additions, and infrastructure improvements.

**Fixes (merged):**

| PR | Summary |
|---|---|
| [#2738 `fix(session-manager)`](https://github.com/nanocoai/nanoclaw/pull/2738) — *writeOutboundDirect opens outbound.db read-only* | Fixes a critical bug where command-gate deny responses were silently dropped because the outbound database was opened in read-only mode. |
| [#2736 `fix(host-sweep)`](https://github.com/nanocoai/nanoclaw/pull/2736) — *grace period for freshly-woken containers* | Prevents premature sweep of stale processing claims when containers are just started. |
| [#2735 `fix(chat-sdk-bridge)`](https://github.com/nanocoai/nanoclaw/pull/2735) — *record the acting user on resolved approval cards* | Adds missing user attribution on approval resolution. |
| [#2741 `fix(setup)`](https://github.com/nanocoai/nanoclaw/pull/2741) — *auto-submit handoff context as Claude's first prompt* | Ensures the interactive setup handoff works correctly by sending a user message to Claude. |

**Features and enhancements (merged):**

| PR | Summary |
|---|---|
| [#2740 `feat(container)`](https://github.com/nanocoai/nanoclaw/pull/2740) — *per-group idle timeout for ephemeral sessions* | Allows clean exit of idle container sessions on a per-group basis. |
| [#2739 `feat(webhook-server)`](https://github.com/nanocoai/nanoclaw/pull/2739) — *raw-route registry for non-Chat-SDK webhooks* | Enables custom webhook routes that bypass the Chat SDK. |
| [#2737 `feat(approvals)`](https://github.com/nanocoai/nanoclaw/pull/2737) — *approval-resolved callback registry* | Modules can now observe approval resolution in a decoupled manner. |
| [#2734 `feat(delivery)`](https://github.com/nanocoai/nanoclaw/pull/2734) — *getDeliveryAction read side* | Adds a read endpoint for the action registry. |
| [#2733 `feat(channels)`](https://github.com/nanocoai/nanoclaw/pull/2733) — *native channel-instance dimension* | Lays groundwork for multi-bot deployments on the same channel. |

---

## Community Hot Topics

**Most active Issue**  
[#1356 – Agent memory system redesign](https://github.com/nanocoai/nanoclaw/issues/1356) (open since March 2026, updated today, 6 👍, 2 comments)  
This long-running discussion has gained traction again. The community recognises the scaling limits of the current memory system (54 files, 83 KB) and clearly desires a more scalable solution. The issue serves as a research tracker rather than a quick fix, indicating it is a deliberate design initiative.

**Most active PR**  
(by discussion or reactions – none of today’s PRs had comments listed, but the open PRs are noteworthy:)  
- [#2744 – fix(signal): deliver agent reactions and forward inbound reactions](https://github.com/nanocoai/nanoclaw/pull/2744) (open, 0 👍) addresses a missing reaction delivery path for the Signal adapter, which is a blocking issue for Signal users.  
- [#2685 – docs(signal): group typing, outbound reactions, quote-reply fix](https://github.com/nanocoai/nanoclaw/pull/2685) (open since June 4) updates Signal documentation and is still awaiting merge.

---

## Bugs & Stability

**High severity (fix merged today):**  
- **`writeOutboundDirect` silently drops command-gate deny responses** – The outbound database was opened read-only, causing INSERT failures to be swallowed. Fix PR [#2738](https://github.com/nanocoai/nanoclaw/pull/2738) was merged today, resolving [#2495](https://github.com/nanocoai/nanoclaw/issues/2495).

**Medium severity (fix merged today):**  
- **Host-sweep premature removal of stale claims** – Freshly-woken containers could have their legitimate processing claims swept. Fixed in [#2736](https://github.com/nanocoai/nanoclaw/pull/2736).  
- **Setup handoff not auto-submitting context** – Interactive Claude handoff did not send a user message, so the context was ignored. Fixed in [#2741](https://github.com/nanocoai/nanoclaw/pull/2741).

**Open bugs (still present):**  
- **Signal adapter: agent reactions not delivered, inbound reactions ignored** – PR [#2744](https://github.com/nanocoai/nanoclaw/pull/2744) is open but not yet merged. This is a functional issue for Signal users.  
- **`ncl wirings create` silently skips `agent_destinations` side effect** – Described in open PR [#2743](https://github.com/nanocoai/nanoclaw/pull/2743). Messages sent to a new chat via a wiring command are dropped. A fix PR is already open.

---

## Feature Requests & Roadmap Signals

The merged PRs today point to several infrastructure features that are likely to land in the next version:

- **Multi-bot channel support** – The new `channel-instance` dimension PR [#2733](https://github.com/nanocoai/nanoclaw/pull/2733) enables running multiple bots on the same channel, a clear enabler for multi-tenant or team deployments.
- **Custom webhook routes** – Raw-route registry [#2739](https://github.com/nanocoai/nanoclaw/pull/2739) will allow integrations that don’t use the Chat SDK.
- **Approval callback hooks** – [#2737](https://github.com/nanocoai/nanoclaw/pull/2737) gives modules a standard way to react to approval outcomes without modifying core logic.
- **Per-group idle timeout** – [#2740](https://github.com/nanocoai/nanoclaw/pull/2740) improves container lifecycle management, especially for ephemeral sessions.

The open issue on memory redesign (#1356) suggests the community wants a more robust long-term memory solution. This could become a major feature in a future milestone.

---

## User Feedback Summary

- **Pain point (Signal users):** Reactions are completely broken on Signal – agents can't react to messages, and incoming reactions are lost. The open PR [#2744](https://github.com/nanocoai/nanoclaw/pull/2744) aims to fix this.
- **Pain point (CLI / wirings):** Creating a wiring silently drops the `agent_destinations` entry, causing sent messages to disappear. Reported in PR [#2743](https://github.com/nanocoai/nanoclaw/pull/2743).
- **Pain point (command-gate denials):** The read-only mode bug caused denial responses to be silently discarded, breaking security/command-gate workflows. Now fixed.
- **Satisfaction signal:** The high volume of well-structured PRs (many labeled `follows-guidelines`) indicates mature contributor processes and a healthy open-source community.

---

## Backlog Watch

| Issue/PR | Age | Why it matters |
|---|---|---|
| [#1356 – Agent memory system redesign](https://github.com/nanocoai/nanoclaw/issues/1356) | ~3 months (created 2026-03-23) | High community interest (6 👍); no decision or implementation yet. May block scaling beyond ~50 agents. |
| [#2685 – docs(signal): group typing, outbound reactions, quote-reply fix](https://github.com/nanocoai/nanoclaw/pull/2685) | 8 days (created 2026-06-04) | Documentation updates for Signal features; still open despite being small. Could be merged once the Signal reaction fix (#2744) lands. |

No other long-unanswered items were identified. The maintainers are clearly very active, as evidenced by today’s large batch of merges.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-06-12

## 1. Today's Overview
Project activity on 2026-06-12 was minimal, with only a single open issue updated in the last 24 hours and no new pull requests or releases. The community remains engaged on a fresh bug report concerning local model behavior, but no merged changes or resolved items were recorded. The project appears to be in a maintenance lull, with no visible development momentum from contributors or the maintainer.

## 2. Releases
No new releases were published today. The latest release remains unspecified.

## 3. Project Progress
No pull requests were updated or merged today. No features, fixes, or refactoring work has been publicly advanced.

## 4. Community Hot Topics
The only active discussion is on **Issue #952** — a bug report regarding incomplete answers when using local models via Ollama. The issue has zero comments and no reactions, so it has not yet sparked community conversation. However, it addresses a core functionality that could affect many users.

- **Issue #952** [bug] Local model using ollama returns incomplete answers  
  Author: bloodgroup-cplusplus | Created: 2026-06-11 | Updated: 2026-06-12 | Comments: 0 | 👍: 0  
  [View on GitHub](https://github.com/nullclaw/nullclaw/issues/952)

## 5. Bugs & Stability
One bug was reported today, with no immediate fix PR:

- **Severity: High** — *Local model (Ollama/Gemma) returns incomplete answers* (Issue #952). The agent truncates responses, making the core feature unusable for affected users. No workaround has been suggested in the issue. Since no corresponding PR exists, this remains an unresolved stability concern.

No crashes or regressions linked to past releases were reported.

## 6. Feature Requests & Roadmap Signals
No new feature requests were filed today. The only open issue is a bug, not a feature. No roadmap signals are present in the latest data.

## 7. User Feedback Summary
One user (bloodgroup-cplusplus) reported a clear pain point: using a locally pulled Gemma model through Ollama results in incomplete sentence answers. The user included a screenshot, indicating frustration with a setup that should “just work.” No positive feedback or satisfaction signals were recorded. The overall sentiment is neutral-to-concerned due to the unresolved bug.

## 8. Backlog Watch
No existing issues or PRs appear to be long-unanswered. The project’s issue tracker currently holds only the newly filed #952, which itself needs maintainer attention. There are no stale items requiring a response.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

## IronClaw Project Digest — 2026-06-12

### 1. Today's Overview
IronClaw remains in an intense development and bug‑fixing phase, with **19 issues updated** and **40 pull requests updated** in the last 24 hours. The Reborn (v2) branch dominates activity: 16 of those issues are open, reflecting a steady stream of user‑reported UI/UX and reliability problems. While no new releases were published today, the team merged/closed **18 PRs**—including several mid‑sized fixes to the Reborn WebUI, capability runtime, and observability wiring. Nightly E2E tests continue to fail, and several high‑severity bugs (driver unavailability, large‑response failures) are being addressed by dedicated fix PRs. Overall, the project is highly active with a clear focus on stabilising the Reborn binary before production cutover.

### 2. Releases
**No new releases** were published in the last 24 hours. The latest release remains **ironclaw v0.29.1** (from PR #3708, still open). The release PR itself (`#3708`) has been open since May 16 and includes breaking API changes in `ironclaw_common` and `ironclaw_skills`; it is periodically updated but not yet merged.

### 3. Project Progress
Today’s merged/closed PRs represent meaningful forward progress on Reborn stability and feature completeness:

- **Production cutover readiness** – [#4763](https://github.com/nearai/ironclaw/pull/4763) (merged) closes the final readiness slice (#4621) with a completeness map and explicit rollback guards.
- **WebUI operator logs** – [#4760](https://github.com/nearai/ironclaw/pull/4760) (merged) wires the v2 Logs page to a real in‑process ring buffer, solving the “empty shell” issue (#4758).
- **Automation run navigation** – [#4757](https://github.com/nearai/ironclaw/pull/4757) (merged) fixes a 404 when clicking triggered runs on the Automations page.
- **Capability runtime resilience** – [#4784](https://github.com/nearai/ironclaw/pull/4784) (merged) converts `HostUnavailable` capability failures into normal tool failures instead of aborting the agent loop.
- **Slack delivery state** – [#4782](https://github.com/nearai/ironclaw/pull/4782) (merged) unifies outbound state stores so WebUI delivery defaults actually reach Slack.
- **Extension activation & GSuite OAuth** – [#4744](https://github.com/nearai/ironclaw/pull/4744) (merged) consolidates extension‑auth‑gating work, fixing the end‑to‑end “connect an extension” user story.
- **QA branch promotion** – [#4786](https://github.com/nearai/ironclaw/pull/4786) (merged) promotes `main` to the QA branch.

Several other PRs were opened but remain open (see Community Hot Topics). The overall quality trajectory shows that each bug report is being quickly followed by a fix PR.

### 4. Community Hot Topics
The most commented issue this week remains the long‑running epic:

- **[#3036 – [EPIC] Configuration‑as‑Code for IronClaw Reborn: tenant blueprints and use‑case harnesses](https://github.com/nearai/ironclaw/issues/3036)** – 7 comments, 1 👍. This epic has been open since April 28 and continues to draw discussion. It envisions a declarative, schema‑driven configuration layer for operators. The lack of movement suggests maintainers are still scoping or waiting for Reborn stablisation.

Other active topics (each with 1 comment, but representative of user engagement):

- **Workspace‑relative path duplication** [#4759](https://github.com/nearai/ironclaw/issues/4759)
- **Agent stops after repeated tool failures** [#4761](https://github.com/nearai/ironclaw/issues/4761)
- **Workspace files not discoverable from WebUI** [#4750](https://github.com/nearai/ironclaw/issues/4750)
- **Large response request fails** [#4751](https://github.com/nearai/ironclaw/issues/4751)
- **Denying shell approval leaves tool pending** [#4764](https://github.com/nearai/ironclaw/issues/4764)

These issues all share a common theme: **the Reborn WebUI and agent execution loop need hardening against realistic user workflows**.

### 5. Bugs & Stability
Today’s bug reports highlight several reliability and UX problems. Ranked by severity:

| Severity | Issue | Description | Fix PR exists? |
|----------|-------|-------------|----------------|
| **Critical** | [#4789](https://github.com/nearai/ironclaw/issues/4789) – `driver_unavailable` after first successful tool call | Agent loop aborts on multi‑tool batches. Root cause in `RefreshingLocalDevCapabilityDriver`. | ✅ [#4790](https://github.com/nearai/ironclaw/pull/4790) (open) |
| **High** | [#4761](https://github.com/nearai/ironclaw/issues/4761) – Agent stops after repeated tool failures | Run completes but the workspace file is not saved because failures are not retried. | No dedicated PR yet |
| **High** | [#4751](https://github.com/nearai/ironclaw/issues/4751) – Large response fails with “tool arguments exceed 16384 bytes” | Provider rejects large tool call payloads. | Not yet |
| **Medium** | [#4783](https://github.com/nearai/ironclaw/issues/4783) – WASM extension capabilities fail with “network” obligation | Pure‑compute extensions blocked before execution. | Not yet |
| **Medium** | [#4762](https://github.com/nearai/ironclaw/issues/4762) – Failed tool workflow causes inconsistent message ordering | Follow‑up messages lose context after a tool failure. | Not yet |
| **Low** | [#4759](https://github.com/nearai/ironclaw/issues/4759) – Workspace path duplication | Paths prefixed with “workspace/” twice when using relative paths. | Not yet |
| **Low** | [#4750](https://github.com/nearai/ironclaw/issues/4750) – Workspace files not discoverable | Files created by agent do not appear in WebUI file browser. | Not yet |
| **Low** | [#4764](https://github.com/nearai/ironclaw/issues/4764) – Deny approval leaves tool pending | No feedback after user denies shell execution. | Not yet |
| **Low** | [#4770](https://github.com/nearai/ironclaw/issues/4770) – Tool activity stops after refresh | Possible SSE reconnect issue. | Not yet |

Additionally, the **Nightly E2E** continues to fail ([#4108](https://github.com/nearai/ironclaw/issues/4108)), last reported today.

### 6. Feature Requests & Roadmap Signals
Several open issues and PRs point toward upcoming functionality:

- **Automated QA for Reborn binary** – [#4775](https://github.com/nearai/ironclaw/issues/4775) (epic) proposes a full CI‑driven QA suite (22 tests already written in PR [#4769](https://github.com/nearai/ironclaw/pull/4769), open).
- **Always‑allow global setting** – [#4776](https://github.com/nearai/ironclaw/issues/4776) (child of #4692) requests a single toggle to auto‑approve eligible tools.
- **Operator log filtering** – [#4771](https://github.com/nearai/ironclaw/issues/4771) and PR [#4788](https://github.com/nearai/ironclaw/pull/4788) (open) add per‑thread/run/turn log filtering.
- **Persistent tenant sandbox** – Design doc PR [#4785](https://github.com/nearai/ironclaw/pull/4785) outlines how to support persistent environments and agent‑built extensions in hosted deployments.
- **Subagent inline prompt budget** – [#4765](https://github.com/nearai/ironclaw/pull/4765) (open) introduces a dedicated type to remove the 512‑byte constraint on subagent prompts.
- **Delivery steering** – [#4780](https://github.com/nearai/ironclaw/pull/4780) (open) adds model guidance for outbound delivery targets before creating triggers.

Taken together, the next release (likely v0.30.x) will probably include: deeper observability, a global “always allow” setting, subagent improvements, and the first iteration of automated QA.

### 7. User Feedback Summary
Real user pain points, mostly from contributors `sunglow666`, `think‑in‑universe`, and `zetyquickly`:

- **Frustration**: ChatGPT subscription device‑code flow is confusing and shows a terminal‑style prompt in the WebUI ([#4711](https://github.com/nearai/ironclaw/issues/4711), now closed and fixed).
- **Dissatisfaction**: File management is broken – created files are not visible in the WebUI ([#4750](https://github.com/nearai/ironclaw/issues/4750)), and relative paths are duplicated ([#4759](https://github.com/nearai/ironclaw/issues/4759)).
- **Reliability concerns**: Agent often stops mid‑workflow or fails silently ([#4761](https://github.com/nearai/ironclaw/issues/4761), [#4789](https://github.com/nearai/ironclaw/issues/4789)), forcing restarts.
- **Approval UX**: Denying a tool approval leaves the user with no feedback and a pending state ([#4764](https://github.com/nearai/ironclaw/issues/4764)).
- **Positive signal**: The team is responding quickly to bug reports – many issues have a matching fix PR within hours (e.g., #4789 → #4790).

### 8. Backlog Watch
Items that have been open for an extended period or are at risk of being forgotten:

- **[#3036 – Configuration‑as‑Code epic](https://github.com/nearai/ironclaw/issues/3036)** – Opened April 28, 2026 (≈1.5 months), 7 comments, no assignee. Despite being a “suggested_P2” and labelled “reborn”, it has not progressed. With Reborn stabilisation still in flight, this may be intentionally deferred, but it remains the largest missing feature for operators.
- **[#4108 – Nightly E2E failure](https://github.com/nearai/ironclaw/issues/4108)** – Opened May 27, 2026, still failing daily. The root cause is not yet documented. This blocks CI confidence.
- **[#3708 – Release PR](https://github.com/nearai/ironclaw/pull/3708)** – Open since May 16, 2026, with breaking API changes. Unmerged for almost a month. While release timing is a conscious decision, the longer it stays open the more risk of drift.
- **[#4692 – IronClaw Reborn Local Testing Findings](https://github.com/nearai/ironclaw/issues/4692)** – Opened June 10, 2026 (2 days ago), already has 3 child issues. This should be monitored as it aggregates first‑run usability bugs that may block adoption.

No issues or PRs have gone completely unanswered by maintainers; the project shows a healthy response cadence.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest – 2026-06-12

## 1. Today's Overview
A busy development day with 19 pull requests updated in the last 24 hours, 13 of which were merged or closed. Only two issues were updated, with one closed (a resolved API error) and one still open (user reporting token waste from repeated output). No new releases were cut today. The high merge volume signals active bug fixing and feature integration, particularly around the Cowork module and gateway stability.

## 2. Releases
None.

## 3. Project Progress
The following PRs were merged or closed today, reflecting significant advances in stability, UX, and new functionality:

**Merged/Closed (13 items):**
- **#2154** – `fix(cowork): show model metadata after stopped streams` – Ensures LLM metadata is preserved when a user manually stops a partial reply.
- **#2153** – `fix(cowork): preserve same-name package model selection` – Resolves model normalization conflicts when custom models share names with built-in packages.
- **#2152** – `fix(cowork): extend pre-send model sync timeout on slow gateways` – Raises timeout to 90s to prevent message drops on cold-start gateways.
- **#2151** – `Feat/2026.6.8 share files` – Adds file sharing capability.
- **#2150** – `fix(kits): keep expert suite controls sticky` – Makes expert suite page title and toolbar fixed/sticky for consistency.
- **#2149** – `fix(openclaw): raise gateway heap limit` – Reduces OOM crashes under long-running multi-channel workloads.
- **#2148** – `feat(cowork): add realtime ASR voice input` – Implements streaming ASR over WebSocket for real-time voice transcription.
- **#2147** – `fix(cowork): prevent stopped startup turns from sending chat` – Cancels turn startup if user stops before gateway becomes active, preventing phantom messages.
- **#1473** – `fix: create Agent modal close unsaved confirmation` – Prevents silent data loss when closing the creation dialog.
- **#1474** – `fix: Agent settings panel close unsaved confirmation` – Same protection for agent configuration panel.
- **#1475** – `fix: MCP server form modal close unsaved confirmation` – Protects MCP server configurations.
- **#1476** – `fix: switch session or navigation view draft persistence` – Ensures input box draft content is saved immediately on view change.
- **#1477** – `fix: re-edit history message overwrite confirmation` – Warns users when editing a past message would overwrite current input.

## 4. Community Hot Topics
- **Issue #2121** ([link](https://github.com/netease-youdao/LobsterAI/issues/2121)) – **"对一个现象的疑问（怀疑是bug）"** (Suspected bug: repeated output wasting tokens)  
  The user reports that the AI repeatedly outputs the same text, consuming tokens needlessly. They attribute it to "claw" behaviour. This issue has 1 comment and remains open since June 7. It highlights a potential usability/performance concern that may affect many users with long-running sessions.

- **Issue #1** ([link](https://github.com/netease-youdao/LobsterAI/issues/1)) – **"hit API error with OpenAI API Type"** (Closed)  
  The user encountered a 400 API error when trying to use the OpenAI API type with a MiniMaxi key. The issue gathered 5 comments and was resolved (closed today). This indicates community engagement and maintainer responsiveness to configuration problems.

## 5. Bugs & Stability
Several critical and medium-severity bugs were fixed today. Ranked by severity:

| Severity | Bug | Fix PR |
|----------|-----|--------|
| **High** | Gateway process OOM under long-running workloads | #2149 (merged) |
| **High** | User-stopped startup turns still sending chat messages | #2147 (merged) |
| **High** | OpenClaw gateway repeatedly restarting in v4.1 | #1446 (open, stale) |
| **Medium** | Disabled skills still injected into conversation prompts | #1453 (open, stale) |
| **Medium** | Scheduled task creation silently fails when date cleared | #1454 (open, stale) |
| **Medium** | Keyboard shortcut conflicts silently override each other | #1456 (open, stale) |
| **Medium** | Agent/MCP/Settings modals losing unsaved changes on close | #1473–#1477 (merged) |
| **Low** | Pre-send model sync timeout too short for slow gateways | #2152 (merged) |
| **Low** | Model metadata lost after manually stopping stream | #2154 (merged) |

The **gateway restart loop** (#1446) and **disabled skills injection** (#1453) are still open but have had recent updates, suggesting they are under review.

## 6. Feature Requests & Roadmap Signals
The following features were merged or are still open as signals for the next release:

- **Realtime ASR voice input** (#2148) – Likely to appear in the next release, adding real-time speech-to-text for Cowork.
- **File sharing** (#2151) – Expanded collaboration features.
- **Scheduled task session folding/grouping** (#1449, open) – Would declutter the session list for recurring tasks. Strong community need.
- **Internationalization (i18n) fixes** (#1448, open) – Adding missing Chinese translations and fixing hardcoded English strings in agent settings.
- **Unsaved changes protection** (#1473–#1477) – Already merged; enhances UX confidence.

Predictions for v4.2 or v4.3: Realtime ASR, file sharing, and scheduled task grouping are the most likely additions given the merge and open PR activity.

## 7. User Feedback Summary
- **Token waste frustration** (#2121): A user suspects the AI is repeating itself and burning tokens. They call it a "bug" and ask for a solution. No maintainer response yet.
- **Configuration confusion** (#1): Setting a MiniMaxi API key with "OpenAI" type caused a 400 error. The issue was resolved, but this indicates users may struggle with API type selection.
- **Data loss concerns**: Multiple users had closed modals losing unsaved changes (addresses by #1473–#1477). The merged fixes directly address this pain point.
- **Gateway instability**: The v4.1 upgrade introduced a gateway restart loop (#1446) that impacted many users (associated issue #1400). The fix is still open.
- **Missing Chinese translations**: Users noticed English text ("Delete", "No matching skills") in the Chinese UI (#1448). This is a reported dissatisfaction with localization.

Overall, users are actively using the product and reporting both stability issues and UX gaps. The maintainer team has been quick to respond with fixes.

## 8. Backlog Watch
The following open items require attention from maintainers:

- **Issue #2121** ([link](https://github.com/netease-youdao/LobsterAI/issues/2121)) – Created June 7, updated June 11, no official response yet. Addresses token waste due to repeated output. High potential impact on user satisfaction.
- **PR #1449** ([link](https://github.com/netease-youdao/LobsterAI/pull/1449)) – `feat(cowork): 定时任务多次执行记录折叠分组展示` – Opened April 3, last updated June 12. Despite recent update, remains open. A feature that would greatly improve user experience for scheduled tasks.
- **PR #1453** ([link](https://github.com/netease-youdao/LobsterAI/pull/1453)) – `fix(skills): 修复已停用技能仍被注入对话提示词的问题` – Critical bug fix for disabled skills still being used. Open since April 3.
- **PR #1446** ([link](https://github.com/netease-youdao/LobsterAI/pull/1446)) – `fix(openclaw): 修复网关反复启动失败导致的无限重启循环` – Major stability fix for gateway crashes. Still open after nearly two months.

These four items represent the most significant unresolved issues from the community perspective.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest — 2026-06-12

## Today's Overview

Activity on the Moltis repository was light but focused, with one open bug report and one open pull request updated in the last 24 hours. No new releases were published. The project remains stable, with maintainers addressing specific delivery issues on WhatsApp and an authentication problem with Fastmail MCP. The single open PR suggests ongoing development, though community engagement is limited to these two items.

## Releases

No new releases were recorded today. The latest release remains unchanged.

## Project Progress

No pull requests were merged or closed today. One PR is currently open:

- **PR #1116** [OPEN] `fix(whatsapp): deliver replies to @lid chats via PN JID rewrite` — Author: `juanlotito`. This PR addresses a silent message drop on WhatsApp when replies are sent to a privacy-enabled sender's `@lid` chat. The fix rewrites the phone number JID to ensure outbound messages reach the user. No comments have been provided yet.

## Community Hot Topics

Only two items were active, neither attracting significant discussion or reactions:

- **Issue #1115** [OPEN] `[bug] [Bug]: Fastmail MCP Authorisation` — Author: `kmath313` (1 comment, 0 👍). The reporter describes a preflight check and indicates the issue occurred on the latest version. No additional context or workarounds have been shared.
- **PR #1116** [OPEN] `fix(whatsapp): deliver replies to @lid chats via PN JID rewrite` — Author: `juanlotito` (0 comments, 0 👍). The PR description notes that replies were visible in the web UI but never delivered to the user; the fix targets the outbound sender logic.

Neither item has drawn community votes or extended threads, suggesting these are early-stage reports.

## Bugs & Stability

One bug was reported today:

- **Issue #1115** — **Fastmail MCP Authorisation** (Severity: Medium). The user confirmed they are on the latest Moltis version and searched for duplicates before reporting. No fix PR has been linked, and no additional reproduction steps or logs are available. This issue may block Fastmail MCP integration for affected users. No other crashes or regressions were reported.

## Feature Requests & Roadmap Signals

No explicit feature requests were filed today. The existing open PR and issue are both bug fixes, indicating the current focus is on stability rather than new functionality. No roadmap signals (e.g., deprecations, new integrations) were detected in the data.

## User Feedback Summary

Both pieces of user feedback reflect concrete pain points:

- **WhatsApp delivery failures**: The fix in PR #1116 addresses a scenario where replies to `@lid` chats are silently dropped, even though they appear in the web UI. This points to a reliability gap in WhatsApp outbound messaging.
- **Fastmail MCP authentication**: The bug report #1115 suggests an authorization failure when using Fastmail's MCP, likely blocking email workflows for users relying on that provider.

No expressions of satisfaction or dissatisfaction beyond the reports were recorded.

## Backlog Watch

No long-unanswered issues or PRs appear in today's 24-hour snapshot. The only open items are recent (created yesterday or today). Maintainers may want to follow up on #1115 with a request for additional logs, and review PR #1116 promptly to close a potential message-loss bug. No stale items were detected.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest — 2026-06-12

## Today's Overview
CoPaw saw **very high activity** on 2026-06-12, with **23 issues** and **37 pull requests** updated in the last 24 hours. Two new patch releases (v1.1.11.post1 and v1.1.11.post2) were published. The community reported several regressions in the v1.1.11 desktop client, including SSL certificate failures, infinite process spawning, and memory leaks. Meanwhile, the development team made progress on multiple fronts: merging security fixes, prepping a major AgentScope 2.0 backend migration, and advancing runtime modularisation. A significant **1.0.13 → 2.0 backend migration** issue (#4727) remains a key strategic focus, and new open PRs signal architectural evolution toward multi-protocol support and sandbox governance.

## Releases
Two patch releases were published today:
- **v1.1.11.post1** — Changelog: bump version, revert a compile-check fix for Discord, and add a release duty checklist. ([GitHub Release](https://github.com/agentscope-ai/QwenPaw/releases/tag/v1.1.11.post1))
- **v1.1.11.post2** — Changes: style fix truncating tool card titles to single line with ellipsis, version bump. ([GitHub Release](https://github.com/agentscope-ai/QwenPaw/releases/tag/v1.1.11.post2))

No breaking changes or migration notes were published. These releases primarily address minor UI polish and packaging stability.

## Project Progress
Today **16 PRs were merged/closed** (of 37 updated). Notable completed work includes:
- **Security**: Merged PR #5028 ([fix(security): isolate keychain master key per install](https://github.com/agentscope-ai/QwenPaw/pull/5028)) — fixes a critical vulnerability where multiple installs shared the same OS keychain master key.
- **Desktop CI**: PR #5125 ([fix(desktop): harden Tauri Windows CI against crates.io fetch failures](https://github.com/agentscope-ai/QwenPaw/pull/5125)) merged — adds retry and caching for Windows Tauri builds.
- **UI Polish**: Three first-time-contributor PRs merged: #5133 (AionUi design language – CSS-only), #5134 (changelog historian agent for dev pipeline), #5136 (complete pt-BR i18n translation).
- **New Plugin Potential**: PR #4622 ([plugin(datapaw): add data-analysis plugin with 12 BI skills](https://github.com/agentscope-ai/QwenPaw/pull/4622)) remains open but has been actively reviewed and touched today.
- **Active architectural work**: PR #5067 ([feat(driver): introduce Agent OS Driver](https://github.com/agentscope-ai/QwenPaw/pull/5067)) and PR #5078 ([feat(runtime): Runtime 2.0 modular architecture](https://github.com/agentscope-ai/QwenPaw/pull/5078)) are under review, both representing major refactoring toward extensibility and protocol abstraction.

## Community Hot Topics
The most active discussions today:

1. **#4727 – [Breaking Change] Migrate backend from AgentScope 1.x to AgentScope 2.0**  
   9 comments, 2 👍. This is the highest-profile open issue, outlining a full dependency upgrade. Community interest is high, but no concrete timeline has been given.  
   [Link](https://github.com/agentscope-ai/QwenPaw/issues/4727)

2. **#5064 – [Bug] Scheduled tasks created by agent fail to trigger**  
   8 comments. A functional regression specific to the Chinese-speaking user base. The agent can create timed tasks but they never execute. No fix PR exists yet.  
   [Link](https://github.com/agentscope-ai/QwenPaw/issues/5064)

3. **#5106 – [Bug] Tauri desktop SSL certificate error + infinite process spawning leading to crash**  
   7 comments. Rapidly closed after a quick investigation, indicating a severe end-user blocker.  
   [Link](https://github.com/agentscope-ai/QwenPaw/issues/5106)

4. **#3817 – [Question] Long-term memory vector model config loses persistence**  
   5 comments. Longstanding issue (since April) still causing user frustration – config resets on container restart.  
   [Link](https://github.com/agentscope-ai/QwenPaw/issues/3817)

5. **#5086 – [Bug] OpenSSL 3.5 regression – desktop fails to start**  
   5 comments, 1 👍. Closed after root cause identified (DER certificate parsing failure). This and #5106 are linked to packaging of the bundled Python.  
   [Link](https://github.com/agentscope-ai/QwenPaw/issues/5086)

Other busy threads: #5137 (memory config loss – 4 comments), #5098 (memory search UI empty – 4 comments), #5138 (Windows client process increase – 3 comments).

## Bugs & Stability
Several stability issues emerged today, concentrated on the desktop client:

- **Critical**: **Windows client process leak** (#5138) – v1.1.11.post2 causes unbounded process creation, saturating memory. No fix PR yet. (Open)
- **Critical**: **SSL certificate crash + infinite spawn** (#5106) – Closed; acknowledged as a severe startup blocker.
- **High**: **OpenSSL regression** (#5086) – Closed; fix involves switching to PEM certificate loading or bundling a patched OpenSSL.
- **High**: **Desktop cannot start** (#5095) – Closed; likely same root cause as #5086/#5106.
- **Medium**: **Memory config loss on collapse** (#5137) – Bug in UI form submission when collapsible sections are not expanded. Fix PR #5144 ([Fix/collapse forcerender memory config loss](https://github.com/agentscope-ai/QwenPaw/pull/5144)) opened same day.
- **Medium**: **Attachment download 404 for docx/pdf** (#5140) – Regression in v1.1.11.post2. No fix PR yet.
- **Medium**: **Coding mode session loss on refresh** (#5142) – Open; fix PR #5147 ([fix: preserve coding session route](https://github.com/agentscope-ai/QwenPaw/pull/5147)) already submitted.
- **Low-Medium**: **Context compression stats mismatch** (#5122) – User reports tool metadata inflating actual input tokens. Open, no fix.
- **Low**: **Math formula rendering: square root display** (#5143) – Closed; likely a frontend KaTeX issue.

Overall, the v1.1.11 release introduced several desktop regressions that are being patched in rapid succession. The “post” release cycle indicates active firefighting.

## Feature Requests & Roadmap Signals
New user-requested features today:

1. **Agent Swarm/Team Collaboration** (#5139) – Explicit request to add multi-agent orchestration similar to WorkBuddy and JiuwenSwarm. This aligns with the open PR #5067 (Agent OS Driver) and hints at a roadmap toward multi-agent collaboration.
2. **Conversation queue, token stats & timestamps** (#5103) – A user who tried OpenClaw wants a non-blocking input queue and per-message token counters. PR #5130 ([feat(chat): add per-turn token and context usage popover](https://github.com/agentscope-ai/QwenPaw/pull/5130)) was opened today, directly addressing token stats.
3. **Coding mode code completion** (#5131) – IDE-like autocomplete integration. This is lower priority but indicates growing use of the Coding mode as a code editor.

Other ongoing feature PRs: customizable column order (#4975), Agent OS Driver (#5067), Runtime 2.0 (#5078), and governance/sandbox interface (#5088). These suggest the next minor version will focus on modularisation, security, and multi-agent extensibility.

## User Feedback Summary
- **Desktop stability remains the #1 pain point**: Multiple users across different regions report that the newest desktop builds (v1.1.11) fail to start or crash severely on Windows. The team is responding with rapid patches.
- **Memory & context management issues frustrate power users**: Config persistence (#3817), memory search rendering (#5098), and token counting inaccuracies (#5122) erode trust in long-term memory features.
- **Attachment handling is inconsistent**: Users report regression from v1.1.10 where file downloads (especially non-text) stopped working (#5102, #5140).
- **Positive signals**: A first-time contributor added complete pt-BR localization (#5136), and another added a historian agent for changelogs (#5134), showing community engagement. The AionUi design language PR (#5133) indicates interest in a more modern UI.

## Backlog Watch
Several important issues and PRs need maintainer attention:

- **#3817 – Vector model config persistence resets on restart** (closed but user still experiencing; verification needed)
- **#5064 – Timed tasks not triggering** – no fix PR after 2 days; could be a core scheduler issue.
- **#5127 – Langfuse traces fragmented** – open bug with no maintainer response; PR #5128 submitted but not yet reviewed.
- **#5098 – Memory search UI empty** – open and affects many users.
- **#4622 – DataPaw plugin** – open since May 22, extensive code, needs final review.
- **#4900 – Decouple plugin loader from agent startup** – open since June 2; critical for PyInstaller/Tauri deployments.
- **#4975 – Customizable column order** – open since June 5, small UI enhancement that could be merged.
- **#5067 and #5078 – Major architecture PRs** – under review but high complexity; need dedicated maintainer cycles.

The team should consider closing or providing timelines on the long-running AgentScope 2.0 migration (#4727) to manage community expectations.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest – 2026-06-12

## 1. Today's Overview
The ZeroClaw project is in a high-activity phase following the **v0.8.0 release** (the “big one” multi-agent rewrite). In the last 24 hours, 23 issues were updated (20 open, 3 closed), and 50 pull requests were touched (47 open, 3 merged/closed). The community is actively testing the new release and reporting blockers on macOS and WSL2, while maintainers are merging fixes for ACP session inheritance and MCP config persistence. Overall project health is stable but stretched – the release queue tracker (#7112) was closed today, signalling that the core v0.8.0 milestone has shipped, but a wave of bug reports and feature requests indicates ongoing stabilization work.

## 2. Releases
**v0.8.0** was published (link: [zeroclaw-labs/zeroclaw/releases/tag/v0.8.0](https://github.com/zeroclaw-labs/zeroclaw/releases/tag/v0.8.0)).  
Key changes:
- **Multi-agent daemon** – one daemon runs many named agents, each with its own workspace, memory, model provider, security policy, channels, and personality.
- **Rewritten configuration schema** with automatic migration from previous versions.
- Likely breaking changes for anyone using the old single-agent config format; migration is automatic but users should verify custom settings.

No migration guide was explicitly mentioned in the data, but the release note states “migrates your existing setup automatically.”

## 3. Project Progress
Three PRs were merged/closed today:
- **[#7517](https://github.com/zeroclaw-labs/zeroclaw/pull/7517)** (closed) – `fix(runtime/subagent): inherit ACP session cwd into spawn_subagent and delegate`. Resolves the subagent workspace jail issue reported in #7263.
- **[#7519](https://github.com/zeroclaw-labs/zeroclaw/pull/7519)** (closed) – `fix(config): persist [[mcp.servers]] per-field edits via natural-key dirty-path walker`. Fixes MCP server config persistence.
- **[#7520](https://github.com/zeroclaw-labs/zeroclaw/pull/7520)** (closed) – `fix(ci): install cross g++ for ARM glibc release builds`. Resolves a build failure on ARM targets for the v0.8.0 release.

Additionally, three issues were closed today: **#7112** (v0.8.0 release queue tracker), **#6443** (Twitch chat channel feature – implemented), and **#7263** (subagent cwd bug – fixed by #7517).

## 4. Community Hot Topics
The most active discussions (4 comments each) revolve around critical bugs and high-risk enhancements:

- **[#5542](https://github.com/zeroclaw-labs/zeroclaw/issues/5542)** – “Consecutive OOM in WSL2” (S0 severity). Continues to see updates; likely requires runtime memory budget improvements.
- **[#6302](https://github.com/zeroclaw-labs/zeroclaw/issues/6302)** – “Gemini 400 – assistant tool_call emitted as first non-system turn”. Provider-specific history serializer issue blocking Gemini users.
- **[#6312](https://github.com/zeroclaw-labs/zeroclaw/issues/6312)** – “per-alias webhook path routing for multi-instance channels”. Feature request for deeper multi-agent channel integration.
- **[#6391](https://github.com/zeroclaw-labs/zeroclaw/issues/6391)** – “real heartbeat tracking for daemon nodes”. Community wants liveness signals beyond static registry entries.

These reflect two key user needs: **stability across platforms/providers** (WSL2, Gemini) and **operational control** (multi-node health, webhook routing).

## 5. Bugs & Stability
Today’s new bugs (all open):
- **[#7523](https://github.com/zeroclaw-labs/zeroclaw/issues/7523)** (S1 – workflow blocked) – **Dashboard not available** after macOS Homebrew install. User reports `cargo web build` is required. Likely a packaging issue.
- **[#7527](https://github.com/zeroclaw-labs/zeroclaw/issues/7527)** (S1 – workflow blocked) – **macOS app doesn’t work** after install – no permission detection, empty page, window disappears. Two separate macOS reports within hours indicate a possible regression in the v0.8.0 macOS build.
- **[#7521](https://github.com/zeroclaw-labs/zeroclaw/issues/7521)** (enhancement/bug?) – `file_read` tool falls back to lossy UTF-8 for non-UTF-8 text (cp1251, Latin-1). Workflow-blocked for users with non-English files. No PR open yet.

Existing high-severity bugs still open (with recent activity):
- [#5808](https://github.com/zeroclaw-labs/zeroclaw/issues/5808) – “Default 32k context budget exceeded on iteration 1” (S1) – no fix PR.
- [#5903](https://github.com/zeroclaw-labs/zeroclaw/issues/5903) – “MCP stdio child processes accumulate with heartbeat” (S1) – no fix PR.
- [#6037](https://github.com/zeroclaw-labs/zeroclaw/issues/6037) – “Cron jobs launched repeatedly” – fix PR #6038 exists but is stale and needs maintainer action.
- [#6173](https://github.com/zeroclaw-labs/zeroclaw/issues/6173) – “model_switch tool not persistent” (S2) – no fix PR.
- [#6350](https://github.com/zeroclaw-labs/zeroclaw/issues/6350) – “WhatsApp allowed-numbers bypassed” (S2) – in progress.
- [#6361](https://github.com/zeroclaw-labs/zeroclaw/issues/6361) – “context_compression drops tool calls for OpenAI-compatible providers” (S1) – in progress.

No fix PRs have been opened for today’s new bugs yet.

## 6. Feature Requests & Roadmap Signals
The following feature/enhancement issues were updated today and signal near-term roadmap priorities:
- **[#6311](https://github.com/zeroclaw-labs/zeroclaw/issues/6311)** – Inject agent alias into system prompt on init (part of multi-agent UX work). Likely to land in v0.8.1.
- **[#7521](https://github.com/zeroclaw-labs/zeroclaw/issues/7521)** – `file_read` charset detection for non-UTF-8 files. Highly demanded by non-English users.
- **[#6390](https://github.com/zeroclaw-labs/zeroclaw/issues/6390)** – `zeroclaw node add <url>` CLI for remote daemon registration. Blocked on node infrastructure.
- **[#6365](https://github.com/zeroclaw-labs/zeroclaw/issues/6365)** – Dashboard “Update ZeroClaw” button. Gateway API enhancement.
- **[#6346](https://github.com/zeroclaw-labs/zeroclaw/issues/6346)** – `zeroclaw node` CLI + dashboard health management. Follow-up to multi-node foundation.
- **[#5618](https://github.com/zeroclaw-labs/zeroclaw/issues/5618)** – Replace DaemonSubsystems with typed Registry API. Architectural refinement to reduce callback complexity.

The **Twitch channel** feature (#6443) was closed today, meaning it shipped or was merged. Community-requested features like **webhook path routing** (#6312) and **heartbeat tracking** (#6391) are accepted but blocked, likely targeting v0.9.0.

## 7. User Feedback Summary
Real pain points from today’s new issues and comments:
- **macOS users** are facing a broken experience post-install – both the CLI dashboard and the GUI app fail (#7523, #7527). This is a critical onboarding blocker for macOS.
- **Windows users** continue to encounter platform-specific issues: Claude Code execution (#7214 – fix PR open), WSL2 OOM (#5542), and setup documentation (#6102 – docs PR open).
- **Non-English users** cannot read their own files via the `file_read` tool (#7521) – a clear localisation gap.
- **Gemini provider users** are blocked by a history serialization bug (#6302) that makes the agent unusable with that model.
- **WhatsApp channel users** report silent message drops when contacts are LID-based (#6350) – a security bypass that undermines configuration trust.

Overall satisfaction is tempered by the v0.8.0 launch – many are excited about multi-agent but frustrated by regressions.

## 8. Backlog Watch
Several important PRs and issues have been ignored or are stuck with the `needs-author-action` label:

- **[#5516](https://github.com/zeroclaw-labs/zeroclaw/pull/5516)** (PR, 2 months old) – Wire fuzz stub targets to real code paths. Stale candidate.
- **[#5661](https://github.com/zeroclaw-labs/zeroclaw/pull/5661)** (PR, 2 months old) – Cron CLI delivery flags. Needs author action, stale.
- **[#5892](https://github.com/zeroclaw-labs/zeroclaw/pull/5892)** (PR, 2 months old) – Three production blockers (tool_choice, orphaned tool_use, vision). Critical fix.
- **[#6038](https://github.com/zeroclaw-labs/zeroclaw/pull/6038)** (PR, 1.5 months old) – Cron lock fix for duplicate execution. Directly addresses #6037.
- **[#6085](https://github.com/zeroclaw-labs/zeroclaw/pull/6085)** (PR, 1.5 months old) – Default session TTL 168h + history trim. Important for memory management.
- **[#6143](https://github.com/zeroclaw-labs/zeroclaw/pull/6143)** (PR, 1.5 months old) – Universal skill registry (agentskills.io, skills.sh). Large feature.
- **[#7214](https://github.com/zeroclaw-labs/zeroclaw/pull/7214)** (PR, 8 days old) – Fix Claude Code on Windows. Needs maintainer review.
- **[#7351](https://github.com/zeroclaw-labs/zeroclaw/pull/7351)** (PR, 5 days old) – MCP auto-reconnect. Needs maintainer review.

Additionally, issue **[#5808](https://github.com/zeroclaw-labs/zeroclaw/issues/5808)** (default context budget exceeded) and **[#5903](https://github.com/zeroclaw-labs/zeroclaw/issues/5903)** (MCP orphan processes) remain open with no direct fix PRs, despite being marked S1 (workflow blocked). Maintainer attention is urgently needed on these to avoid user churn post-v0.8.0.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*