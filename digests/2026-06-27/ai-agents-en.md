# OpenClaw Ecosystem Digest 2026-06-27

> Issues: 213 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-27 09:15 UTC

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

# OpenClaw Project Digest — 2026-06-27

## 1. Today’s Overview

Project activity remains **very high**: 213 issues and 500 pull requests were updated in the last 24 hours. However, no new releases were published today, and the majority of open issues (203) are still awaiting triage or decision. Security and stability concerns dominate the conversation, with **diamond lobster**-rated bugs (highest severity) accounting for a large share of top issues. The community is pushing hard for per-agent cost controls, better media handling across channels, and tighter sandbox isolation. While the project is clearly under heavy development, the backlog of “needs-maintainer-review” and “needs-product-decision” items suggests that maintainer bandwidth is becoming a bottleneck.

## 2. Releases

**None.** No new versions were released on 2026-06-27.

## 3. Project Progress

A total of **27 PRs were merged or closed** today. Among the top 30 PRs by comment count, only one is closed:

- **PR #68936** – *Autofix: add PR review autofix pipeline + Windows daemon* — closed (likely merged), adding a pipeline that uses the Claude Agent SDK to address review comments automatically and a Windows background daemon for gateway supervision.

Most of the open PRs in the top 30 are still in review or awaiting author feedback. Notable open PRs that advanced today include:

- **PR #89038** – Fix for QQbot outbound resolution and delivery drain on reconnect (in re-review loop).
- **PR #96701** – Bounds Anthropic SSE streaming reads at 16 MiB (ready for maintainer look).
- **PR #96723** – Same bound fix for the legacy provider path.
- **PR #94945** – Fix for DeepSeek cache boundary suffix causing collapsed hit rates (waiting on author).

The activity pattern shows a clear focus on **memory safety / resource exhaustion** fixes (many “bound read” PRs) and on **channel-specific reliability** (QQbot, Codex isolation, cron handling).

## 4. Community Hot Topics

The most discussed issues (by comment count) reveal deep concerns around **message leakage, security boundaries, and multi-agent coordination**.

| Issue | Comments | Summary |
|-------|----------|---------|
| [#25592](https://github.com/openclaw/openclaw/issues/25592) 🔥 | **32** | Text between tool calls leaks to messaging channels (e.g., Slack, iMessage). Users report that internal processing output gets routed to active channels. This is a critical UX/security issue requiring product decision. |
| [#39604](https://github.com/openclaw/openclaw/issues/39604) | **13** | Feature request to allow private network access (`allowPrivateNetwork`) via config. Highly upvoted (9 👍). |
| [#42475](https://github.com/openclaw/openclaw/issues/42475) | **12** | Per-agent cost budget enforcement at gateway level – a direct operator need for runaway spend prevention. |
| [#41744](https://github.com/openclaw/openclaw/issues/41744) | **12** | Feishu image tool loses media before final delivery – a channel-specific media handling bug. |
| [#40001](https://github.com/openclaw/openclaw/issues/40001) | **11** | Write tool lacks append mode; cron sessions overwrite shared files, causing silent data loss. |

On the PR side, comment counts are not reported, but the most active PRs by size and merge-risk involve **DeepSeek model compatibility** (PR #97221, #97208) and **session isolation for Codex** (PR #81777). The community is clearly struggling with **model-provider-specific issues** (DeepSeek thinking format, Anthropic streaming bounds) and **multi-channel concurrency** (Telegram lane starvation discussed in #41120).

## 5. Bugs & Stability

Several high-priority bugs were discussed today, many flagged as **diamond lobster** (highest impact):

| Bug | Severity | Summary | Fix PR? |
|-----|----------|---------|---------|
| [#25592](https://github.com/openclaw/openclaw/issues/25592) | P1 – Security / Message loss | Text between tool calls leaks to channels | linked PR open |
| [#39604](https://github.com/openclaw/openclaw/issues/39604) (feature) | P2 – Security | Private network blocked by default | no PR yet |
| [#38327](https://github.com/openclaw/openclaw/issues/38327) | P1 – Crash loop | “Cannot convert undefined or null to object” with Google Vertex/Gemini 3.1 after 2026.3.2 upgrade | no PR yet |
| [#45224](https://github.com/openclaw/openclaw/issues/45224) | P1 – Crash | Unhandled Playwright assertion error crashes Gateway process | no PR yet |
| [#39807](https://github.com/openclaw/openclaw/issues/39807) | P1 – Infinite retry | Billing error (402) causes death spiral with no backoff | linked PR open |
| [#44749](https://github.com/openclaw/openclaw/issues/44749) | P1 – Data loss | Concurrent allow-always approvals lose allowlist entries (race condition) | no PR yet |
| [#39847](https://github.com/openclaw/openclaw/issues/39847) | P1 – Security | Echo contamination: internal metadata leaked to Discord | linked PR open |
| [#94518](https://github.com/openclaw/openclaw/issues/94518) | P1 – Performance | DeepSeek cache hit rate <10% after 6.x upgrade | PR #94945 open (waiting on author) |

A notable regression: **#38439** and **#41201** both report that **agent avatars in the Control UI are broken** (404 for `/avatar/{agentId}`). Both are flagged as P2 regressions with linked PRs still open.

The recurring theme is **unbounded resource consumption** (SSE streams, JSON responses) – multiple PRs today (e.g., #96701, #96723, #96782, #97147) specifically add byte-level caps to prevent OOM. These fixes are now “ready for maintainer look,” suggesting the project is actively hardening against memory exhaustion attacks.

## 6. Feature Requests & Roadmap Signals

The following features received notable community attention (many are P2, diamond lobster):

- **#42475** – Per-agent cost budget enforcement (daily/monthly caps) – likely to land in next minor version given operator demand.
- **#42840** – MathJax/LaTeX support in Control UI – high upvotes (7 👍), relatively simple front-end change.
- **#28300** – Theme customization system (6 presets + custom studio) – community design feedback from #28048.
- **#38626** – Subagent lifecycle observability (events, artifacts, supervision) – aligns with growing multi-agent usage.
- **#43454** – Gateway lifecycle hooks (onSubagentComplete, onToolCallThreshold) – requested by power users.
- **#43564** – ACP session skill context injection – would bridge skills into external agents.
- **#40786** – `.gitignore`-like exclude patterns for `backup create` – needed for large installations.
- **#40418** – Automated session memory preservation on `/new` – addresses a common pain point.

Prediction: **Cost budget enforcement (#42475)** and **avatar fix (#38439/#41201)** are most likely to appear in the next release (2026.7.x), as they directly affect operators’ ability to run the platform reliably. **LaTeX rendering (#42840)** is a relatively low-risk enhancement that could also slip in soon.

## 7. User Feedback Summary

Real pain points expressed in today’s discussions:

- **“Text between tool calls leaks to channels”** (#25592): Users describe the agent’s internal processing being visible to end users – a **critical UX and security failure**.
- **“Shared files get overwritten by cron sessions”** (#40001): Users relying on `write` for memory/logging lose data because there’s no append mode.
- **“Billing errors cause infinite retry loops burning API credits”** (#39807): One operator reported 5,206+ failed agent runs in 6 hours.
- **“Avatar broken in UI”** (#38439): A regression that makes the Control UI look broken – undermines professional deployments.
- **“DeepSeek cache hit rate collapsed after upgrade”** (#94518): A 10× cost increase due to cache-misalignment – a significant financial pain.
- **“Browser-heavy sessions starve all other channels”** (#41120): A concrete concurrency design issue affecting multi-channel operators.

On the positive side, the community is actively contributing fixes: many PRs adding resource bounds and better error messages (e.g., #96790 adds actionable install hints) show user satisfaction with the project’s direction, even as stability issues persist.

## 8. Backlog Watch

Several important items have been languishing without maintainer action, despite high impact:

| Issue | Impact | Status |
|-------|--------|--------|
| [#37634](https://github.com/openclaw/openclaw/issues/37634) – Sandbox with `workspaceAccess: none` leaves workspace read-only | Session state / Security | Needs maintainer review and product decision (since March 6) |
| [#38327](https://github.com/openclaw/openclaw/issues/38327) – “Cannot convert undefined or null to object” regression | Crash loop | Needs live repro, maintainer review, product decision (since March 6) |
| [#40001](https://github.com/openclaw/openclaw/issues/40001) – Write tool lacks append mode | Data loss | Needs product decision, linked PR open (since March 8) |
| [#39476](https://github.com/openclaw/openclaw/issues/39476) – A2A `sessions_send` causes duplicate messages | Session state / Message loss | Needs product decision, linked PR open (since March 8) |
| [#39847](https://github.com/openclaw/openclaw/issues/39847) – Echo contamination (metadata leak) | Security | Stale, needs security review, linked PR open (since March 8) |

These items have been waiting for **product decisions** or **security reviews** for over 3 months. While the project is clearly active, the steady stream of new issues may be overwhelming the maintainers. A dedicated triage sprint or clearer decision-making process could help clear this backlog.

---

*Generated from OpenClaw GitHub data as of 2026-06-27.*

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: Personal AI Assistant Open-Source Ecosystem
**Date: 2026-06-27**

---

## 1. Ecosystem Overview

The personal AI assistant open-source landscape is experiencing intense development velocity, with six of twelve tracked projects showing "very high" daily activity. The ecosystem is converging on several critical requirements: **per-agent cost controls**, **multi-channel reliability**, **security hardening** (message leakage, sandbox isolation, FIPS compliance), and **plugin/extensibility systems** that allow third-party tool integration without core modifications. While projects share common DNA—most derive from or are inspired by OpenClaw's architecture—they are rapidly differentiating around target users: enterprise/security-focused (ZeroClaw, OpenClaw), personal assistant (NanoBot, PicoClaw), and agent collaboration/cowork (CoPaw, LobsterAI). A notable bifurcation exists between projects prioritizing **feature velocity** (IronClaw, CoPaw) and those in **stabilization phases** (Hermes, LobsterAI). The ecosystem remains heavily dependent on a small number of core maintainers, with several projects showing critical backlog bottlenecks.

---

## 2. Activity Comparison

| Project | Issues Updated (24h) | PRs Updated (24h) | Release Today | Health Score | Key Signal |
|---------|----------------------|-------------------|---------------|--------------|------------|
| **OpenClaw** | 213 | 500 | None | ⚠️ High activity, critical backlog | Diamond lobster bugs dominate; maintainer bottleneck |
| **NanoBot** | 25 | 54 | None | ✅ Healthy, rapid iteration | 31 PRs merged/closed; security fixes fast-tracked |
| **Hermes Agent** | 0 | 50 | None | ✅ Stabilizing | 47 PRs merged; zero new issues; systematic bug sweeps |
| **PicoClaw** | 4 | 13 | None | 🟢 Moderate, refining | 8 PRs merged; code hygiene + SSRF fix |
| **NanoClaw** | 3 | 5 | None | 🟡 Moderate, bug surfaced | Critical `/update-skills` regression |
| **NullClaw** | 1 | 0 | None | 🔴 Low, stagnating | Single build issue 65 days open; no maintainer response |
| **IronClaw** | 4 | 50 | None | ⚠️ Very high, intense porting | 16 merged; Reborn WebUI feature parity push |
| **LobsterAI** | ~10 | ~15 | **Yes** (v2026.6.26) | ✅ Healthy, post-release | 10 merged; critical installer/UI freeze bugs |
| **TinyClaw** | 0 | 0 | None | 🔴 Inactive | No activity |
| **Moltis** | 0 | 1 | None | 🔴 Low, quiet | Single open PR, no community engagement |
| **CoPaw** | 11 | 37 | **Yes** (v2.0.0-beta.1) | ⚠️ Very high, beta turbulence | 9 merged; breaking changes causing plugin failures |
| **ZeroClaw** | 31 | 50 | None | ✅ High, structured milestones | 6 merged; v0.8.3 + v0.9.0 trackers active |

**Health Score Key:** ✅ = Healthy iteration | 🟢 = Moderate/Stable | 🟡 = Moderate with concerns | ⚠️ = High activity with risks | 🔴 = Low activity/Stagnating

---

## 3. OpenClaw's Position

**Advantages:**
- **Largest community engagement** by far (213 issues, 500 PRs in 24h)—serves as the ecosystem's core reference implementation
- **Most comprehensive security framework** with diamond lobster severity ratings and systematic hardening (SSE bounding, sandbox isolation, message leakage prevention)
- **Deepest model provider coverage**—fixes span DeepSeek, Anthropic, Google Vertex, OpenAI-compatible endpoints simultaneously
- **Strongest multi-channel support**—QQbot, Feishu, Telegram, Slack, iMessage all receiving channel-specific fixes

**Technical Approach Differences:**
- OpenClaw uses a **monolithic gateway architecture** requiring per-agent cost budgets at the gateway level (vs. NanoBot's per-conversation model overrides)
- Security is **enforcement-first** (private network blocked by default, sandbox `workspaceAccess:none` reality) compared to NanoBot's after-the-fact shell pattern validation
- Community contributions are **overwhelming maintainers**—203 of 213 open issues await triage, creating a bottleneck that smaller projects avoid

**Community Size Comparison:**
- OpenClaw's 500 daily PRs dwarf every other project combined (~260 total across all others)
- However, **merge rate is low** relative to activity (27 PRs merged vs. 500 updated)—suggesting a large review queue
- NanoBot (31 merged/54 updated) and Hermes (47 merged/50 updated) show higher throughput-to-activity ratios

**Key Risk:** OpenClaw's maintainer bandwidth is the ecosystem's single point of failure. Critical items waiting 3+ months for product decisions could push users to more responsive alternatives like NanoBot or ZeroClaw.

---

## 4. Shared Technical Focus Areas

The following requirements are emerging across **multiple projects**, indicating ecosystem-wide priorities:

| Focus Area | Projects Affected | Specific Needs |
|------------|------------------|----------------|
| **Per-Agent Cost Controls** | OpenClaw (#42475), ZeroClaw ($ budgets), CoPaw (model fallback #5572) | Daily/monthly caps, runaway spend prevention, billing error backoff |
| **Security & Sandboxing** | OpenClaw (message leakage #25592), NanoBot (shell bypass #4521), PicoClaw (SSRF #3074), Hermes (FIPS #51962) | Channel leakage prevention, filesystem isolation, FIPS compliance, private network guards |
| **Platform Reliability** | Hermes (clock monotonicity), LobsterAI (UI freeze #2214), CoPaw (conversation loss #5579), NanoBot (Windows shell inconsistency) | Cross-platform encoding, non-blocking I/O, crash recovery, session persistence |
| **Media & File Handling** | OpenClaw (Feishu image loss #41744), CoPaw (file generation rendering #4865), Moltis (auto-screenshot #1135) | Streaming file content, append mode, media delivery guarantees |
| **Plugin / Extensibility** | NanoBot (plugin system #2231), ZeroClaw (Wasm runtime #8368), CoPaw (Skill marketplace) | Manifest-based loaders, capability enforcement, signed distribution |
| **Multi-Provider Model Support** | OpenClaw (DeepSeek cache #94518), CoPaw (DeepSeek V4 errors #5573), ZeroClaw (OpenRouter fallbacks #8138) | Model-specific format fixes, fallback chains, embedding persistence |
| **Channel-Specific UX** | ZeroClaw (WhatsApp group context #8379), NanoBot (WhatsApp self-chat), OpenClaw (Telegram lane starvation #41120) | Typing indicators, group context awareness, lane scheduling |
| **Context/Budget Management** | ZeroClaw (context overflow #5808), OpenClaw (SSE bounds #96701), CoPaw (scroll context #5321) | Token budget enforcement, streaming bounds, retrieval-driven context |

**Observation:** The most uniformly shared need is **security & cost control**—operators running these agents in production are demanding predictable spend and guaranteed isolation. This suggests the ecosystem is maturing from "can it work?" to "can we run it safely at scale?"

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | ZeroClaw | CoPaw | Hermes | IronClaw |
|-----------|----------|---------|----------|-------|--------|----------|
| **Primary Target User** | Enterprise/Security teams | Individual power users | Security-conscious operators | Agent collaboration teams | CI/CD/automation | WebUI developers |
| **Architecture** | Monolithic gateway | Modular plugin host | Wasm-first, milestone-driven | Beta/migration (2.0) | Python agent runtime | Rust-based, Reborn WebUI |
| **Channel Coverage** | Maximal (Slack, iMessage, QQ, Feishu, Discord, Telegram) | Strong (WhatsApp, Telegram, Discord) | Broad (Inkbox, WhatsApp, email, SMS, iMessage) | Strong (Feishu, cron, SSH) | Minimal (primarily API) | Focused (WebUI, Slack) |
| **Security Approach** | Enforcement-first (blocked by default, diamond lobster severity) | After-the-fact pattern validation | Supply-chain provenance (SLSA), Wasm capability enforcement | Beta stability gaps | Systematic hardening (FIPS, encoding, monotonic clocks) | Approval UX (default-on eligible tools) |
| **Innovation Pace** | Bug-fix heavy, slow feature decisions | Rapid iteration, high merge rate | Structured milestone trackers | Beta turbulence, fast test infra | Stabilization, zero new features | Intense porting to Reborn v2 |
| **Community Contribution Model** | Overwhelming maintainers; large backlog | Responsive; security fixes within 24h | Structured RFCs, milestone trackers | Growing; first-time contributors | Systematic sweep PRs from core team | Core-team dominated |
| **Model Provider Strategy** | Deepest breadth, per-provider fixes | Configurable per-conversation | OpenRouter fallbacks, embedding identity | DeepSeek compatibility focus | Minimal (CI use case) | OpenAI-compatible focus |

**Key Differentiation Insight:** These projects are not direct competitors—they serve different tiers of the same ecosystem. OpenClaw is the **battle-tested reference**, NanoBot is the **fast-moving personal assistant**, ZeroClaw is the **security-first reimagining**, and CoPaw is the **collaboration/enterprise play**. A developer choosing between them should weigh **stability vs. velocity** and **feature depth vs. maintainability**.

---

## 6. Community Momentum & Maturity

### Tier 1: Rapidly Iterating (New features + major migrations)
- **CoPaw**: v2.0.0-beta.1 signals a major architectural shift; high community engagement but beta turbulence (plugin failures, conversation loss). Momentum is highest here, but risk is also highest.
- **IronClaw**: Intense Reborn WebUI port—50 PRs/day, 16 merged. Feature parity push suggests an imminent stable release. Core-team dominated, reliable but less community-driven.
- **ZeroClaw**: Structured milestone v0.8.3 + v0.9.0 trackers show disciplined development. Wasm plugin runtime and supply-chain signing are differentiated bets. Community contributions growing.

### Tier 2: Stabilizing (Bug fixes, hardening, no major new features)
- **Hermes Agent**: Zero new issues, 47 systematic bug fixes merged. In a pure stabilization phase—clock monotonicity, FIPS, encoding sweeps. Next release will be a reliability milestone.
- **LobsterAI**: Post-release cleanup (10 merges) but facing critical installer/UI freeze bugs. Mermaid rendering and skill lifecycle fixes show attention to UX polish.
- **PicoClaw**: Refinement phase—code hygiene, SSRF security fix. No new features, but no regressions either. Stable for moderate use cases.

### Tier 3: Moderate Activity (Bug fixes + targeted features)
- **NanoClaw**: Single critical bug (/update-skills no-op) needs attention, but per-group model override and dashboard pusher signal feature-forward direction.
- **Moltis**: Minimal activity; single auto-screenshot PR. Either feature-complete or resource-constrained.

### Tier 4: Dormant / Inactive
- **NullClaw**: 65-day-old build issue with no maintainer response. Signals project abandonment risk.
- **TinyClaw**: No activity.
- **ZeptoClaw**: No activity.

**Momentum Verdict:** The ecosystem is **healthy but fragmented**. Core OpenClaw is the bottleneck preventing even faster ecosystem growth; its backlog could drive power users to NanoBot or ZeroClaw. The three most promising projects for long-term investment (based on velocity + responsiveness) are **NanoBot**, **ZeroClaw**, and **CoPaw**.

---

## 7. Trend Signals

### Industry Trends Extracted from Community Feedback

1. **Per-Agent Cost Control Is Non-Negotiable**
   - Multiple projects (OpenClaw #42475, ZeroClaw budget enforcement, CoPaw model fallback #5572) are building cost caps, billing error backoff, and spend monitoring. **Agent operators are demanding predictable costs** before deploying at scale.

2. **Security Is Shifting from "Can We" to "How Do We Guarantee"**
   - FIPS compliance (Hermes), SLSA provenance (ZeroClaw #8177), message channel leakage prevention (OpenClaw #25592), and shell injection prevention (NanoBot #4521) signal that **enterprise and government adoption is the next frontier**. Projects without a security roadmap will be left behind.

3. **Multi-Provider Resilience Is Table Stakes**
   - DeepSeek cache collapses (OpenClaw #94518), Vertex crash loops (#38327), and OpenRouter fallback requests (ZeroClaw #8138) show that **relying on a single AI provider is untenable**. Model-agnostic architecture and automatic fallback chains are becoming requirements, not differentiators.

4. **Context Management Is the New Memory**
   - Context budget overflow (ZeroClaw #5808), scroll/retrieval-driven context (CoPaw #5321), and SSE bounding (OpenClaw #96701) all point to **context window management as the critical UX and cost lever**. Projects that implement first-class context governors (trim, budget, persistence) will win user trust.

5. **Channel Integration Is Commoditizing**
   - Every major project supports Telegram, WhatsApp, Discord, Slack. The differentiator is now **channel-specific UX**—typing indicators (NanoBot), passive group context (ZeroClaw #8379), lane scheduling (OpenClaw #41120), and media delivery guarantees. **Deep channel integration, not breadth, is the competitive edge.**

6. **Plugin/Extensibility Systems Are the Next Battleground**
   - NanoBot (#2231), ZeroClaw (#8368 Wasm), and CoPaw (Skill marketplace) are all investing in plugin systems. The standard is not yet set (manifest-based vs. Wasm vs. scripting), but the **direction is clear: modular, capability-enforced, signed plugin distribution**.

### Value for AI Agent Developers

- **Build cost governance early**—retrofitting it is painful (ask OpenClaw #39807, 5,206 failed runs in 6 hours).
- **Assume multi-provider from day one**—abstract provider interface, plan for format differences, implement fallback chains.
- **Prioritize context budget enforcement**—the #1 silent killer of agent reliability (ZeroClaw #5808 shows system prompt alone can exceed default budget).
- **Invest in platform testing**—Windows, FIPS, non-UTF-8 locales. Hermes's systematic sweeps are the right approach; one-time fixes don't scale.
- **Choose your community model**—OpenClaw shows that popularity without maintainer bandwidth creates bottlenecks. Smaller projects like NanoBot achieve higher throughput-to-activity ratios. Decide whether you want many issues or fast merges.

---

*Report generated from community digest data as of 2026-06-27. Activity metrics are 24-hour snapshots and may not reflect long-term trends.*

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest – 2026-06-27

## Today's Overview
The NanoBot project saw **high activity** over the past 24 hours: **25 issues** were updated (16 open, 9 closed) and **54 pull requests** were updated (23 open, 31 merged/closed). No new releases were published. A wave of **bug fixes and feature PRs** from contributors like `axelray-dev` and `dajiaohuang** were merged, addressing long-standing issues around session collisions, Anthropic provider compliance, tool scheduling, and plugin extensibility. Two security advisories (#4521, #4518) were closed with fix PRs. Overall, the project is in a **healthy, fast-paced development cycle**, with maintainers actively reviewing and merging contributions.

## Releases
**None** – No new versions were tagged today.

## Project Progress
**31 pull requests were merged or closed** today. Key advancements include:

- **Plugin system** (PR #4558, fixes #2231): A minimal plugin system with manifest loader (`plugin.json`), discovered from `~/.nanobot/plugins/` and entry points. Supports tools, skills, and MCP server configs.
- **External agent delegation** (PR #4559, fixes #3436, #3024): New `agent_delegate` tool calls external AI agents (Claude Code, Codex, opencode).
- **Trust LLM parallel tool calls** (PR #4557, fixes #3096): Tool scheduler now runs all LLM-requested tool calls concurrently instead of serializing by tool class.
- **Text-to-speech tool** (PR #4560, fixes #4010): Supports `edge-tts`, macOS `say`, `espeak-ng`, and Windows SAPI.
- **Core bug fixes** (PRs #4533, #4532, #4531, #4530, #4523): Session key collision on disk, Anthropic `type` field validation, stream delta coalescing with `_stream_id`, duplicate tool call IDs in non-stream parser, and a flaky test fix.
- **WhatsApp improvements** (PRs #1450, #4317, #2411, #3761, #3051): Allow self-chat via linked device, typing indicators, configurable group replies.
- **Skill deduplication guard** (PR #4554): Prevents Dream from creating duplicate skills.
- **Security hardening** (PR #4562): Validate each shell segment against `exec.allowPatterns` to prevent shell-chain bypass.

## Community Hot Topics
Most active discussions (by comments / reactions):

1. **Feature: Plugin System** – Issue #2231 (4 comments, 0 👍)  
   User `andrader` requested a plugin system akin to Copilot CLI. **Now merged** in PR #4558.  
   [Issue #2231](https://github.com/HKUDS/nanobot/issues/2231)

2. **Override Model Per Conversation** – Issue #4253 (4 comments, 0 👍)  
   User `rombert` wants to switch between OpenRouter and local LlamaCpp models per task. No PR yet; high demand.  
   [Issue #4253](https://github.com/HKUDS/nanobot/issues/4253)

3. **Automatic Reasoning Effort Escalation** – Issue #4419 (3 comments, 0 👍)  
   `orrinwitt` proposes dynamic escalation of `reasoningEffort` for complex tasks. No PR yet.  
   [Issue #4419](https://github.com/HKUDS/nanobot/issues/4419)

4. **Heartbeat Tasks Deliver to Original Channel** – Issue #4418 (2 comments, 0 👍)  
   Heartbeat results go to the last active channel, not the one where the task was added. Open.  
   [Issue #4418](https://github.com/HKUDS/nanobot/issues/4418)

5. **Call External Agent** – Issue #3436 (2 comments, 0 👍)  
   `jsapede` asked for delegating to external agents. **Now merged** in PR #4559.  
   [Issue #3436](https://github.com/HKUDS/nanobot/issues/3436)

6. **Tool Scheduling – Trust Parallel Calls** – Issue #3096 (2 comments, 0 👍)  
   `chenyahui` highlighted serialization of tool calls. **Now merged** in PR #4557.  
   [Issue #3096](https://github.com/HKUDS/nanobot/issues/3096)

7. **Lightweight Claim Contradiction** – Issue #660 (12 comments, 5 👍)  
   User `besoeasy` disputes the “ultra-lightweight” tag due to Node.js dependency. **Closed** – no further action taken.  
   [Issue #660](https://github.com/HKUDS/nanobot/issues/660)

## Bugs & Stability
**Bugs reported today (2026-06-26/27), ranked by severity:**

| Severity | Issue | Summary | Status & Fix PR |
|----------|-------|---------|----------------|
| **Critical** | #4521 | `exec.allowPatterns` shell-chain bypass allows unintended command execution | **Closed** – fixed by PR #4562 |
| **Critical** | #4518 | Default login-shell execution in `exec` reintroduces secrets from shell startup files | **Closed** – fixed by PR #4562 (or separate) |
| **High** | #4544 | Windows `exec` uses cmd.exe vs PowerShell inconsistently for single/multi-line commands | Open, no fix PR yet |
| **Medium** | #4539 | Telegram messages not rendering on Telegram Web v0.2.2 | **Closed** (likely fixed) |
| **Low** | #4057 | Session key collision on disk (distinct keys map to same file) | **Closed** – fixed by PR #4533 |
| **Low** | #4060 | Anthropic provider emits content blocks without required `type` | **Closed** – fixed by PR #4532 |
| **Low** | #4063 | Stream delta coalescing ignores `_stream_id`, merges distinct streams | **Closed** – fixed by PR #4531 |
| **Low** | #4059 | Non-stream parser preserves duplicate tool call IDs | **Closed** – fixed by PR #4530 |
| **Low** | #4511 | Windows `--background` flag causes inconsistent state after `/restart` | Open |
| **Low** | #4513 | NSSM service + `/restart` causes port conflicts and service stop | Open |

The two security issues (#4521, #4518) were **quickly patched** within 24 hours.

## Feature Requests & Roadmap Signals
User-requested features with **no merged PR yet** (potentially in next release):

- **Model override per conversation** (#4253) – highly actionable, similar to existing provider override patterns.
- **Automatic reasoning effort escalation** (#4419) – aligns with recent `reasoningEffort` support.
- **Heartbeat results delivered to original channel** (#4418) – a UX improvement likely to be picked up soon.
- **Heartbeat-specific model override** (#4431) – follow-up to heartbeat feature.
- **Provider override for Dream model** (#4029) – separate provider for Dream vs. main agent.
- **Crawl4AI web fetching support** (#2700) – alternative to Jina/fallback.
- **API authentication on wildcard bind** (#4490) – security parity with WS gateway.
- **TTS voice output** (#4010) – **already merged** in PR #4560, expect it in next release.
- **Plugin system** (#2231) – merged.
- **External agent delegation** (#3436) – merged.
- **Parallel tool calls** (#3096) – merged.

**Prediction:** The next minor version (v0.2.3 or v0.3.0) will likely include the plugin system, agent delegation, TTS, parallel tool calls, and the security fixes. Per-conversation model override and reasoning escalation may follow soon after.

## User Feedback Summary
- **Positive:** Contributors appreciate the rapid bug fixes (e.g., session collisions, Anthropic provider issues) and the new plugin system. WhatsApp improvements (self-chat, typing indicators) fulfill community requests.
- **Pain Points:**
  - Windows compatibility issues: `--background` restart bug (#4511), NSSM service integration (#4513), inconsistent shell semantics (#4544).
  - Heartbeat always uses a separate session (#1899, #4418) – users want it to share the main session or deliver results correctly.
  - The project’s “ultra-lightweight” branding is disputed due to Node.js dependency (#660), though closed without resolution.
- **Use Cases:** Users are running NanoBot in mixed environments (local + cloud models), as a personal assistant via Telegram/WhatsApp, and as a code companion. The demand for plugin extensibility and external agent delegation shows a growing need to integrate with specialized AI tools.

## Backlog Watch
- **Issue #1899** (March 2026) – Heartbeat session isolation from main session. Two comments, no maintainer response. Still open.
- **Issue #2700** (April 2026) – Crawl4AI support for web fetching. One comment, no action.
- **Issue #4253** (June 2026) – Model override per conversation. No PR yet, but moderate community interest.
- **Issue #4418** (June 2026) – Heartbeat channel delivery. No PR.
- **Issue #4511, #4513** (June 25) – Windows bugs. No fix PR yet, but maintainer activity high.

Maintainers have shown excellent responsiveness today; these older items may be addressed in the coming weeks.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-06-27

## Today’s Overview

The project saw no new issues filed in the last 24 hours, but 50 pull requests were updated, of which 47 were closed or merged — a strong signal of ongoing cleanup and stabilization. No new releases were published. Activity was dominated by systematic bug fixes spanning multiple components (agent, gateway, CLI, tools), with a particular focus on platform compatibility (Windows, FIPS‑enabled Linux) and reliability improvements (clock monotonicity, encoding, boolean coercion). The high merge rate and low incoming issue volume suggest the team is actively reducing technical debt ahead of an upcoming release.

## Releases

None.

## Project Progress

All 47 merged or closed PRs are bug fixes. The work can be grouped into several systematic sweeps:

- **Clock monotonicity** (PRs #50584, #50593, #50595, #50599, #51301, #51679) – Replaced `time.time()` with `time.monotonic()` across agent runtime helpers, tool executor, display spinner, chat completion helpers, gateway, and TUI modules to eliminate negative or inflated durations caused by NTP/clock adjustments.
- **FIPS compatibility** (PRs #51962, #51973) – Added `usedforsecurity=False` to all `hashlib.md5()`, `sha256()`, and `sha1()` calls (62 total) for non‑security purposes (caching, content hashing) to prevent crashes on FIPS‑enabled systems.
- **Explicit UTF‑8 encoding** (PRs #50655, #50660, #50666, #50679) – Added `encoding="utf-8"` to all `read_text()`/`write_text()` calls across CLI, agent, tools, skills, and gateway (dozens of files) to avoid `UnicodeDecodeError` on Windows and other non‑UTF‑8 locales.
- **Boolean coercion** (PRs #50604, #50646) – Replaced `bool(config.get("enabled"))` with `is_truthy_value()` to correctly handle quoted YAML strings like `"false"`.
- **Assert removal** (PR #50565) – Replaced bare `assert` statements in the WeChat platform module with `RuntimeError` guards to prevent silent failures under `python -O`.
- **Transport‑recovery alignment** (PRs #52227, #51793) – Aligned `_TRANSIENT_TRANSPORT_ERRORS` with the classifier’s `_TRANSPORT_ERROR_TYPES` and removed overly broad `max_tokens` pattern from overflow detection.
- **Windows UX** (PR #53397) – Suppressed console flash in Copilot auth token lookup by adding `creationflags=windows_hide_flags()`.

Additionally, one gate‑narrowing fix for WebSocket reconnect races (#53525) remains open and under review.

## Community Hot Topics

No issues or PRs received significant comments or reactions within the last 24 hours. The two currently open PRs that may attract attention are:

- **#53525** – *fix(gateway): preserve rebound ws sessions during teardown*  
  Author: konsisumer | [Link](https://github.com/nousresearch/hermes-agent/pull/53525)  
  Addresses a narrow race condition where sessions on a disconnected WebSocket are detached after ownership changes. The underlying need is to prevent stale session cleanup from breaking active desktop connections.

- **#53487** – *fix(file_tools): keep repeated-slash tilde paths under home*  
  Author: Harshkamdar67 | [Link](https://github.com/nousresearch/hermes-agent/pull/53487)  
  Fixes a security/functional issue where paths like `~//scratch/file.txt` would resolve under filesystem root instead of the user’s home directory.

## Bugs & Stability

All bug reports resolved today were found and fixed within the same PR set — no new issue was opened. The bugs are ranked by severity:

| Severity | Bug | Fix PR(s) | Description |
|----------|-----|-----------|-------------|
| **High** | Production crash on FIPS‑enabled systems | #51962, #51973 | `hashlib.md5/sha1/sha256` without `usedforsecurity=False` raises `ValueError` on RHEL 8/9 FIPS mode. Affects 62 call sites. |
| **High** | UnicodeDecodeError on non‑UTF‑8 locales | #50655, #50660, #50666, #50679 | `Path.read_text()` without `encoding="utf-8"` fails on Windows (cp1252) and other locales. Affects configs, user data, service scripts. |
| **Medium** | Negative/inflated duration measurements | #50584, #50593, #50595, #50599, #51301, #51679 | `time.time()` causes wildly inaccurate elapsed times under NTP or manual clock changes. Affects agent loop, tool executor, spinner, gateway. |
| **Medium** | Boolean config misparse | #50604, #50646 | `bool("false")` returns `True`, causing wrong behavior for YAML values like `enabled: "false"`. |
| **Medium** | WeChat platform silent failures | #50565 | `assert` stripped under `-O`, leaving uninitialized sessions. |
| **Low** | Windows console flash | #53397 | `subprocess.run()` without hide flags causes visible CMD windows every ~15s during credential polling. |
| **Low** | Reconnect race in WebSocket teardown | #53525 (open) | Sessions detached stale after owner change during reconnection. |

All listed bugs (except #53525) now have a merged fix.

## Feature Requests & Roadmap Signals

No feature requests were submitted in the last 24 hours. The set of merged PRs indicates a strong focus on **reliability** and **platform hardening** rather than new functionality. Based on the systematic nature of the fixes (clock monotonicity, FIPS, encoding), the next version of Hermes Agent is expected to be a **stability release** with improved cross‑platform compatibility and fewer clock‑sensitive anomalies.

## User Feedback Summary

No direct user feedback was recorded in issues or PR comments. However, the bugs fixed today reveal common pain points:

- **Clock drift** – Users running Hermes on systems with NTP sync or manual clock changes saw incorrect telemetry and potential timeouts.
- **FIPS failures** – Government/enterprise users on RHEL 8/9 were unable to run the agent at all without patching hash calls.
- **Windows encoding errors** – A significant portion of Windows users likely encountered `UnicodeDecodeError` when reading JSON configs or skill manifests.
- **Boolean config confusion** – Users who quoted YAML booleans in their configs experienced unexpected activation/deactivation of features.

The absence of new issue reports today could reflect that these dominant pain points have been addressed by the merged fixes.

## Backlog Watch

No issues are languishing unanswered — the last 24 hours saw zero new issues. Two open PRs may require maintainer attention:

- **#53525** (open, created 2026-06-27) — WebSocket session teardown race, no comments yet. Should be reviewed promptly as it relates to a known reconnect bug.
- **#53487** (open, created 2026-06-27) — Tilde path expansion fix, no comments yet. Involves a functional regression that could affect file tool behavior.

Neither appears to be a long‑standing item; both are recent and under review.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest – 2026-06-27

## Today's Overview
The project saw moderate activity with 4 issues updated and 13 pull requests updated in the last 24 hours. Of those PRs, 8 were merged or closed, reflecting a strong focus on code hygiene (explicit error ignoring) and targeted bug fixes. No new releases were made. The community contributed several fixes for Android startup issues and a security enhancement for SSRF guards. The overall pace suggests a period of refinement and stability improvements rather than major new features.

## Releases
No new releases today.

## Project Progress
**Merged/Closed Pull Requests (8 total):**
- [#3181 – `fix(gateway): guard startup info assertions`](https://github.com/sipeed/picoclaw/pull/3181) – Improves gateway stability by handling missing or malformed startup info sections.
- [#3143 – `fix(web): block private IPv4 embeds in ISATAP literals`](https://github.com/sipeed/picoclaw/pull/3143) – Closes an SSRF vulnerability (#3074) by teaching the IP classifier to detect embedded private IPv4 addresses within ISATAP IPv6 literals.
- [#3187 – `test(utils): explicitly ignore resp.Body.Close() errors in tests`](https://github.com/sipeed/picoclaw/pull/3187) – Code hygiene across test helpers.
- [#3188 – `fix(health): explicitly ignore json.Encode errors`](https://github.com/sipeed/picoclaw/pull/3188) – Suppresses secondary errors in health server responses.
- [#3186 – `fix(membench): explicitly ignore resp.Body.Close() error`](https://github.com/sipeed/picoclaw/pull/3186) – Cleanup in membench LLM retry loop.
- [#3185 – `fix(updater): explicitly ignore resp2.Body.Close() error`](https://github.com/sipeed/picoclaw/pull/3185) – Cleanup in updater checksum download path.
- [#3184 – `fix(channels): explicitly ignore resp.Body.Close() errors in websocket dial cleanup`](https://github.com/sipeed/picoclaw/pull/3184) – Cleanup across Pico and WhatsApp channels.
- [#3183 – `fix(onebot): explicitly ignore resp.Body.Close() error after websocket dial`](https://github.com/sipeed/picoclaw/pull/3183) – Cleanup for OneBot channel.

These merges represent a systematic effort to eliminate unhandled error warnings, improving code quality and reducing noise for static analysis tools.

## Community Hot Topics
- **Issue #3088 – `[Feature] use vodozemac instead of libolm`**  
  [🔗](https://github.com/sipeed/picoclaw/issues/3088)  
  *2 reactions, 3 comments, labeled “help wanted, priority: high”*  
  This request to replace the deprecated `libolm` with `vodozemac` has gathered the most community support. It reflects a strong desire for security and compliance with modern Matrix standards. The issue has been open since June 9 but was updated yesterday, indicating ongoing interest.

- **Issue #3150 – `[BUG]它给自己整失忆了 (AI forgets itself)`**  
  [🔗](https://github.com/sipeed/picoclaw/issues/3150)  
  *3 comments, 0 reactions, marked stale*  
  A user reports that the AI agent loses context (“memory loss”). While not heavily upvoted, the number of comments suggests the reporter and maintainers are actively discussing possible causes.

- **Issue #3182 – `[BUG] Android version`**  
  [🔗](https://github.com/sipeed/picoclaw/issues/3182)  
  *0 comments, new*  
  Fresh bug report about an Android service launch failure. The user attached a screenshot showing a path change issue. This is likely to gain attention soon.

## Bugs & Stability
| Bug | Priority | Description | Fix Available? |
|-----|----------|-------------|----------------|
| **#3182 – Android launch failure** | **High** | Service cannot start; user cannot change settings path. New report. | No fix PR yet. |
| **#3150 – AI memory loss** | **Medium** | Agent forgets context; stale but updated. May be related to chat history handling. | No open fix. |
| **#3094 (closed) – Duplicate messages from sub-agents** | **Medium** | Resolved by a previous merge; confirmed fixed. | Already merged. |
| **SSRF bypass (fixed in #3143)** | **Critical** | ISATAP literal allowed access to private IPs. | Merged today. |

No regression reports were observed. The Android bug (#3182) is the most urgent unaddressed stability issue.

## Feature Requests & Roadmap Signals
The **vodozemac** replacement (#3088) is the most prominent feature request. It aligns with the project’s security roadmap and is likely to be prioritized, especially given its “priority: high” label and 2 thumbs-up. No other new feature requests appeared today. The recent merges (e.g., SSRF guard, startup stability) suggest the team is focusing on hardening existing functionality before introducing new capabilities.

## User Feedback Summary
- **Pain points expressed:**
  - AI losing context over time (#3150) – frustration with inconsistent agent memory.
  - Android version broken on some devices (#3182) – cannot launch the service, blocking mobile use.
  - Duplicate message spam from sub-agent tasks (#3094, now fixed) – user satisfaction improved by the fix.
- **Use case notes:**
  - Reporter of #3088 emphasizes security compliance, likely using PicoClaw in a production environment that requires up-to-date cryptographic libraries.
- **Overall sentiment:** Mixed – users appreciate responsive bug fixing (duplicate messages solved, SSRF fixed) but face ongoing stability issues, especially on Android.

## Backlog Watch
- **Issue #3150** (open since June 19, updated June 26) – The “AI forgets” bug remains unresolved and has been marked stale. No assignee or priority label. Needs maintainer triage.
- **Issue #3088** (open since June 9, updated June 26) – While actively discussed, no PR has been submitted for the vodozemac switch. Given its high priority and maintainer attention, this should be addressed soon.
- **Pull Request #3180** (open, `fix(cli): skip tool calls with invalid arguments`) – Opened June 26 but not yet merged. Could prevent CLI crashes; consider prioritizing.

No other long-stale items are apparent. The current development pace indicates the team is responsive to both security and quality issues.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-27

## Today's Overview

The project shows steady development activity with 5 pull requests and 3 issues updated in the last 24 hours. One critical bug surfaced in the `/update-skills` command that silently breaks channel refreshes, while two feature PRs (per‑group model override and a dashboard pusher) signal growing infrastructure for multi‑agent orchestration and observability. One migration‑related fix was merged, improving upgrade reliability. Overall, the project is in a healthy, feature‑forward state, though the new bug needs prompt attention.

## Releases

No new releases were published today.

## Project Progress

- **#2859 (merged)** – [fix(migrate-v2): don't SELECT is_main from v1 registered_groups](https://github.com/nanocoai/nanoclaw/pull/2859)  
  Fixed a migration crash on older v1 databases (e.g., 1.1.0) where the `is_main` column did not exist. Ensures a smooth upgrade path to v2.

- **#2870 (open)** – [fix(whatsapp): keep native participant addressing for group encryption](https://github.com/nanocoai/nanoclaw/pull/2870)  
  Addresses a WhatsApp group reply issue: replies were logged as delivered but never appeared due to incorrect metadata in the Baileys socket hook. A promising fix awaiting review.

## Community Hot Topics

No issues or PRs attracted heavy commentary today. The most discussed item is the new `/update-skills` bug:

- **#2868 (open)** – [Bug: /update-skills silent no‑op for already‑installed channels](https://github.com/nanocoai/nanoclaw/issues/2868)  
  1 comment. User reports that running the update command does not actually refresh adapter code or pinned dependencies, nullifying the intended `[Unreleased]` migration workflow.

## Bugs & Stability

| Severity | Issue / PR | Description | Fix Available? |
|----------|------------|-------------|---------------|
| **High** | [#2868](https://github.com/nanocoai/nanoclaw/issues/2868) | `/update-skills` is a silent no‑op; pre‑flight skips code/deps refresh | No open fix PR yet |
| Medium | [#2870](https://github.com/nanocoai/nanoclaw/pull/2870) | WhatsApp group replies not delivered – encryption metadata mismatch | Fix PR submitted |
| Low | [#2860](https://github.com/nanocoai/nanoclaw/pull/2860) | `libsignal` debug logging spills session keys into console | Fix PR submitted |

The high‑severity bug (#2868) directly impacts users who rely on the upgrade path described in the changelog. No fix PR has been proposed yet.

## Feature Requests & Roadmap Signals

Two new feature PRs were opened today, pointing toward enhanced configuration flexibility and observability:

- **#2872** – [feat(opencode): per‑group model override via container_configs.model](https://github.com/nanocoai/nanoclaw/pull/2872)  
  Allows each OpenCode agent group to run a different model by injecting `OPENCODE_MODEL` into the container. This is a strong signal that multi‑model deployments are becoming a priority.

- **#2871** – [feat(dashboard): add dashboard pusher with OpenCode support](https://github.com/nanocoai/nanoclaw/pull/2871)  
  Collects NanoClaw state snapshots and posts them to a dedicated dashboard server every 60 seconds. Likely to be merged soon, given the project’s increasing complexity.

A closed feature request (#1275) for auto‑prompt when added to a new group was updated yesterday, possibly reactivated, but remains closed. It may resurface if users continue to ask for group onboarding improvements.

## User Feedback Summary

The primary pain point expressed today is the frustrating silent failure of `/update-skills`. Users who followed the migration instructions expecting skill updates to work are left with stale code. No positive feedback was recorded. The WhatsApp group fix (#2870) addresses a real user‑reported problem of messages not appearing, indicating user satisfaction concerns on that channel.

## Backlog Watch

- **#2868** (open, no comments after initial report) – No maintainer response yet. Should be triaged and assigned to prevent upgrade confusion.
- **#2860** (open PR) – A relatively simple logging fix that has been open for a day; could be merged quickly to reduce noise.
- **#1275** (closed) – Although closed, it was updated recently, possibly due to user interest. The project may want to consider a lighter version of auto‑registration as a low‑hanging feature.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest – 2026-06-27

## Today's Overview
Project activity remains low, with no new pull requests or releases in the last 24 hours. The only update was a comment on a long-standing Android/Termux build failure (Issue #868), which was last touched by a user today. This suggests minimal maintenance velocity, though the single open issue signals a possible platform compatibility gap that has not yet been addressed. Overall project health is stable but shows signs of stagnation in community contributions.

## Releases
*No new releases in the reporting period.*

## Project Progress
- **Merged/closed PRs today:** None  
- **Features advanced:** None  

No code changes were merged or closed in the last 24 hours.

## Community Hot Topics
The only active discussion is:

- **Issue #868** – [zig build fails on Android/Termux (aarch64) with AccessDenied on options.zig linkat](https://github.com/nullclaw/nullclaw/issues/868)  
  *Author: NOTJuangamer10* | Created: 2026-04-23 | Updated: 2026-06-27 | Comments: 4  
  The user reports that `zig build -Doptimize=ReleaseSmall` fails on Android/aarch64 with an `AccessDenied` error related to `options.zig linkat`. The issue has been open for over two months but received a comment today, possibly a bump or a suggestion. No maintainer response is visible in the provided data. The underlying need is reliable cross‑platform build support, especially on Android environments like Termux.

## Bugs & Stability
- **Active bug (medium severity):** Issue #868 describes a build failure that prevents compiling NullClaw on Android/Termux (aarch64). The error is a filesystem permission issue during linking, which blocks users on that platform entirely. No fix PR has been submitted yet.  
- **No other crashes or regressions** reported in the last 24 hours.

## Feature Requests & Roadmap Signals
No explicit feature requests were recorded in the last 24 hours. However, the existence of Issue #868 highlights a demand for better Android packaging/CI testing. If the maintainer prioritizes cross‑platform builds, a future release may include a corrected build script for Termux or a workaround for the `linkat` permission error.

## User Feedback Summary
- **Pain point:** Users on Android (aarch64) cannot build NullClaw from source after the v2026.4.17 release, even with the recommended Zig 0.16.0.  
- **Satisfaction impact:** The lack of resolution for over two months may frustrate mobile developers and reduce adoption on non‑standard Linux environments.  
- **Use case:** Likely power‑users running NullClaw in Termux for on‑device development or automation.

## Backlog Watch
- **Issue #868** – Open for 65 days, no maintainer comment. A fix or workaround would significantly improve platform coverage.  
- **No other long‑unanswered items** currently in the backlog.

**Project health indicator:** 🟢 Stable but low activity – one blocking issue remains unaddressed, and community engagement is minimal.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest – 2026-06-27

## 1. Today’s Overview

IronClaw saw very high activity over the past 24 hours: **4 issues** were updated (3 still open), and **50 pull requests** were touched (34 open, 16 merged/closed). No new releases were cut. The bulk of the work continues to be the **massive port of legacy E2E/Playwright coverage** to the Reborn WebUI v2 stack, alongside targeted fixes for tool-approval UX, error surfacing, and API alignment. A nightly E2E test failure (Issue #4108) indicates a potential regression, but no release is blocked. Overall, the project is in an intense **feature‑completion and stabilisation phase** for the Reborn platform.

## 2. Releases

No new releases today. The only release‑related PR (#5311) is still open; it would bump several crates (`ironclaw_common` 0.4.2→0.5.0, `ironclaw` 0.24.0→0.29.1) with API‑breaking changes.

## 3. Project Progress

Sixteen PRs were merged or closed today. Key merged/closed items that advanced the codebase:

| PR / Issue | Summary |
|------------|---------|
| [#5366](https://github.com/nearai/ironclaw/pull/5366) (merged) | **Feat(approvals):** Default “Always allow eligible tools” to ON, closing Issue #5364. |
| [#5347](https://github.com/nearai/ironclaw/pull/5347) (closed) | **Port Reborn Responses API input handling** – adds OpenAI‑compatible input handling and workflow contracts. |
| [#5346](https://github.com/nearai/ironclaw/pull/5346) (closed) | **Align Reborn runtime tool surface** – family wiring and local‑dev external‑tool capability fixes. |
| [#5364](https://github.com/nearai/ironclaw/issues/5364) (closed) | Issue: “Make ‘Always allow eligible tools’ the default” – resolved by PR #5366. |

These merge/close actions show the team is steadily **eliminating legacy gaps** and **tightening default UX** (no more per‑call approval prompts out of the box).

## 4. Community Hot Topics

No issues or PRs accumulated more than **zero comments or reactions** today. The liveliest area remains the ongoing **Reborn WebUI port wave** (PRs #5374, #5375, #5372, #5373, #5371, #5370, #5376 all opened today by `ilblackdragon`). Though not “hot” in discussion, the volume signals sustained core‑team focus.

The **Nightly E2E failure** (Issue #4108) is a CI‑triggered alert without community discussion, but its recurring nature makes it a silent hot topic for stability.

## 5. Bugs & Stability

One **high‑severity** automated bug report and three fix PRs targeting stability surfaced:

- **🚨 High: Nightly E2E failed** – Issue [#4108](https://github.com/nearai/ironclaw/issues/4108) (updated today, still open). The `Full E2E / E2E (extensions)` workflow failed on commit `5298504a`. No fix PR linked yet; this may be a transient infrastructure issue or a real regression from the day’s many changes.

- **🟡 Medium: Vague tool error messages** – PR [#5338](https://github.com/nearai/ironclaw/pull/5338) (open) surfaces real failure detail instead of generic `"invalid_input"`. Fix end‑to‑end at terminal and per‑tool activity card.

- **🟡 Medium: Approval resume loop** – PR [#5306](https://github.com/nearai/ironclaw/pull/5306) (open) prevents infinite `ask_each_time` approval gates after a one‑shot lease approval.

- **🟢 Low: WebUI Retry button no‑op** – PR [#5365](https://github.com/nearai/ironclaw/pull/5365) (open) fixes the Retry button that was wired to a stub and did not actually re‑send.

- **🟢 Low: Cranelift debug log floods** – PR [#5369](https://github.com/nearai/ironclaw/pull/5369) (open) suppresses noisy compiler logs, keeping `IRONCLAW_REBORN_LOG=debug` useful.

All fix PRs are from core contributors, suggesting good maintainer attention.

## 6. Feature Requests & Roadmap Signals

- **Default tool‑approval off** (Issue [#5364](https://github.com/nearai/ironclaw/issues/5364)) – already shipped in PR #5366. Expect this in the next release.
- **Non‑Slack channel pairing** – Issue [#5368](https://github.com/nearai/ironclaw/issues/5368) requests an end‑to‑end wire for generic channel pairing in WebUI (currently only Slack works). Related PRs [#5373](https://github.com/nearai/ironclaw/pull/5373) and [#5372](https://github.com/nearai/ironclaw/pull/5372) are already open, so this may land in days.
- **Reborn capability policy epic** – Issue [#5261](https://github.com/nearai/ironclaw/issues/5261) (open 2 days) continues the admin‑shared tools & skills with per‑user auth. No PRs yet, but it is a tracked epic for the Reborn stack.

The consistent pattern of “port legacy browser coverage” PRs (#5371–#5376) strongly indicates the next minor release will **reach feature parity** between the old WebUI and Reborn WebUI v2.

## 7. User Feedback Summary

No direct user comments or reactions are present in today’s data. However, two implicit pain points are visible:

- **Approval fatigue** – The closure of Issue #5364 (“Always allow eligible tools” default on) reflects a user request to reduce per‑call prompts. The fix is already merged.
- **Broken Retry button** – PR #5365 addresses a UX regression that would frustrate chat users. It is still open but actively being fixed.
- **Obscured error details** – PR #5338 tackles the developer/end‑user irritation of seeing “invalid_input” instead of the actual failure reason.

## 8. Backlog Watch

- **Nightly E2E failure** – Issue [#4108](https://github.com/nearai/ironclaw/issues/4108) (open 31 days, updated today) has no assignee and no fix PR. If persistent, it could mask real regressions.
- **Mass dependency update** – PR [#5271](https://github.com/nearai/ironclaw/pull/5271) (open 2 days, 47 package updates) touches nearly every crate. Given the high risk (marked `risk: high`), it needs careful review.
- **External‑tool Responses round‑trip** – PR [#5099](https://github.com/nearai/ironclaw/pull/5099) (open 8 days, size XL) is a foundational feature for OpenAI‑compatible external tools. Delays may block dependent work.
- **Release PR #5311** – Open for 1 day; it includes API‑breaking changes. Maintainers should consider it once the porting wave stabilises.

No issues or PRs are critically lagging without any maintainer response; the team appears highly responsive.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

Here is the LobsterAI project digest for 2026-06-27.

---

# LobsterAI Project Digest — 2026-06-27

## 1. Today's Overview

The project shows **high activity** today, with a significant focus on stabilization and bug fixing following a new release. The team merged 10 pull requests (PRs), many of which were old "stale" fixes for long-standing bugs, alongside several new urgent patches for the renderer and Mermaid artifact handling. The open issue count remains low at 2, indicating the community is currently reporting critical, user-facing problems rather than minor requests. The release of **LobsterAI 2026.6.26** introduces a new "plan mode" for cowork, suggesting a push toward more complex workflow features, while the team is actively cleaning up technical debt (stale PRs).

## 2. Releases

- **Version:** [LobsterAI 2026.6.26](https://github.com/netease-youdao/LobsterAI/releases/tag/2026.6.26) (Released 2026-06-26)
- **Key Changes:**
    - `feat(openclaw)`: Upgraded the OpenClaw runtime to `v2026.6.1`.
    - `feat(cowork)`: Added a new **"plan mode"** workflow for collaborative agent tasks.
    - `fix(openclaw)`: Support for an upgraded Instant Messaging (IM) plugin.
- **Breaking Changes / Migration Notes:** No breaking changes or migration notes were provided in the release data. The runtime upgrade appears to be a minor version bump.

## 3. Project Progress

Today, the project resolved a clean-up of 10 merged/closed PRs. Beyond the routine stale PR closures, several focused fixes were landed:

- **Renderer & Artifact Stability:**
    - [PR #2213](https://github.com/netease-youdao/LobsterAI/pull/2213) and [PR #2210](https://github.com/netease-youdao/LobsterAI/pull/2210) addressed unstable Mermaid diagram rendering. The fix prevents raw error SVGs from leaking into the DOM and ensures proper cleanup of rendering containers.
    - [PR #2212](https://github.com/netease-youdao/LobsterAI/pull/2212) and [PR #2213](https://github.com/netease-youdao/LobsterAI/pull/2213) stabilized the skill search popover, preventing it from closing while a user's focus is still inside the menu.
- **Main Process:**
    - [PR #2211](https://github.com/netease-youdao/LobsterAI/pull/2211) fixed a code quality issue by sorting imports in OpenClaw patch decision tests.

## 4. Community Hot Topics

The community is currently focused on two **critical, user-facing bugs**, both reported by the same user (`woxinsj`). There are no issues with high comment counts, suggesting these are isolated but severe experiences.

1.  **[#2215 - Installation Failure](https://github.com/netease-youdao/LobsterAI/issues/2215) (Open):** A user has exhaustively debugged a persistent `Resource extraction failed` error during installation. Their analysis reveals a multi-layered problem involving environment variables (`ERROR_BAD_ENVIRONMENT`) and confusing real vs. false installation paths. Underlying Need: **Robust, clean installer** that can handle complex Windows environments and provide clear error messages.
2.  **[#2214 - Desktop UI Freeze on Backup](https://github.com/netease-youdao/LobsterAI/issues/2214) (Open):** A user reports that the "Data Backup" function causes the entire main process to become unresponsive (100% reproducible). Underlying Need: **Non-blocking background operations** for I/O heavy tasks like database backup, especially for users with large SQLite databases (71.6 MB).

## 5. Bugs & Stability

Two new bugs were reported today, both rated **high severity**.

| Issue | Severity | Description | Fix PR Exists? |
| :--- | :--- | :--- | :--- |
| [#2214 - Desktop UI Freeze on Backup](https://github.com/netease-youdao/LobsterAI/issues/2214) | **HIGH** | 100% reproducible. The main window becomes unresponsive when clicking "Backup Data". Likely a main-thread blocking I/O issue. | No |
| [#2215 - Installation Failure](https://github.com/netease-youdao/LobsterAI/issues/2215) | **HIGH** | Persistent "Resource extraction failed" error. User traced it to environment issues and multiple installation directories. | No |

A notable **stability regression** was fixed today: The stale [PR #1446](https://github.com/netease-youdao/LobsterAI/pull/1446) was merged, tackling an infinite restart loop of the OpenClaw gateway that could bring the entire app down. Its merge today implies this fix is now in the latest release.

## 6. Feature Requests & Roadmap Signals

- **Cron Job Session Management:** The user need for better session organization is now addressed. The merge of [PR #1449](https://github.com/netease-youdao/LobsterAI/pull/1449) (folded grouped display for repeated cron job executions) suggests the next version (2026.7.x) will have improved sidebar management for recurring tasks.
- **Non-Blocking Backup:** The crash report in [#2214](https://github.com/netease-youdao/LobsterAI/issues/2214) strongly signals a user need for a **background/async backup process**. This is a strong candidate for a hotfix in the next point or minor release.
- **Unified Skill State:** The merge of [PR #1453](https://github.com/netease-youdao/LobsterAI/pull/1453) fixes a critical "skill state sync" issue where disabled skills were still being injected into prompts. This indicates the team is working on a more robust skill management lifecycle, which could be a feature in a future roadmap.

## 7. User Feedback Summary

- **Pain Points (Installation):** The user `woxinsj` provided an incredibly detailed debug log for the installation failure (Issue #2215). The complexity (checking anti-viruses, manual path discovery) indicates a frustrating user experience for first-time setup.
- **Pain Points (Stability):** The same user reported a complete UI freeze (Issue #2214) during a routine backup operation. This is a clear sign of poor responsiveness for a core feature, leading to user dissatisfaction.
- **Satisfaction Signals:** The closure of 10 stale PRs (including fixes for i18n errors, skill injection, and shortcut conflicts) shows the team is listening to long-standing community reports and fixing them. This improves overall user trust.

## 8. Backlog Watch

- **[PR #2065 - Agent ID Generation](https://github.com/netease-youdao/LobsterAI/pull/2065) (Open, Stale):** This PR has been open for a month (since May 28) and addresses a serious **data safety** issue—deleted agents can be "resurrected" if a new one with the same name is created because ID is derived from the name. This fix uses short UUIDs. It's a significant risk for users who frequently test or recreate agents. The PR remains open and needs maintainer attention to merge or update.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest – 2026-06-27

## Today’s Overview
Project activity was minimal over the past 24 hours. No new issues were opened or updated, and no releases were published. A single open pull request (#1135) received an update, but it remains unmerged and has attracted no comments or reactions. The overall pulse suggests a low-activity period, possibly following a recent release cycle or maintenance window. No bugs, regressions, or stability issues were reported, indicating a relatively quiet phase for the codebase.

## Releases
None in the last 24 hours.

## Project Progress
No pull requests were merged or closed today. The sole open PR (#1135) did not advance to a merge state.

## Community Hot Topics
- **#1135 – browser: optional auto-screenshot after each action**  
  [GitHub PR #1135](https://github.com/moltis-org/moltis/pull/1135)  
  *Author: resumeparseeval | Updated: 2026-06-26 | Comments: 0 | 👍: 0*  
  This pull request proposes a new feature that automatically captures a screenshot after every state‑changing browser action and attaches it to the action’s tool result. The capture logic is placed at the single dispatch point in `BrowserManager::execute_action` (`crates/browser/...`). Although the PR currently has no community engagement (no comments or reactions), it signals interest in richer per‑step visual feedback for chat clients. The feature could be used to build step‑by‑step screenshot timelines, improving debugging and end‑user transparency. Maintainers may want to solicit feedback to gauge demand and finalize the design.

## Bugs & Stability
No bugs, crashes, or regressions were reported in the last 24 hours.

## Feature Requests & Roadmap Signals
The only feature‑related signal today is PR #1135 (auto‑screenshots after browser actions). If accepted, this would likely land in the next minor release (e.g., vX.Y+1) and could be accompanied by a configuration toggle (e.g., `auto_screenshot: bool`) to avoid breaking existing workflows. No other feature requests appeared in issues or PRs.

## User Feedback Summary
No user feedback (issues, comments, or reactions) was recorded in the past 24 hours. The project appears to be in a low‑engagement phase.

## Backlog Watch
No important issues or PRs with long inactivity require immediate maintainer attention. The oldest open contributions (not shown in today’s data) are outside the 24‑hour window and should be reviewed separately if they have been dormant for weeks.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest — 2026-06-27

## Today’s Overview

CoPaw (the QwenPaw agent platform) shows very high activity today, driven by the release of `v2.0.0-beta.1` and a surge in community bug reports and feature requests. A total of **11 issues** were updated (10 still open) and **37 pull requests** saw activity (28 open, 9 merged/closed). The release marks a major refactoring step—migrating the agent runtime to AgentScope 2.0—but the beta carries breaking changes that are causing plugin and integration failures. Simultaneously, the project is investing heavily in test infrastructure: four unit-test PRs (backend and frontend) were advanced, covering modules like crons, runner, app-infra, and console stores. Community engagement is strong, with users reporting real-world pain points around message spam, conversation loss, and channel integration issues.

## Releases

**v2.0.0-beta.1** was published today (tag: `agentscope-ai/QwenPaw v2.0.0-beta.1`). This is an **early beta for QwenPaw 2.0.0**, explicitly marked as unstable and not intended for production. The only documented change is a single refactor commit: “refactor: migrate agent”. However, the migration underlies many subsequent issues—several official plugins fail to install on 2.0 (see [#5568](https://github.com/agentscope-ai/QwenPaw/pull/5568)), and multiple CI/CD fixes were required (e.g., [#5578](https://github.com/agentscope-ai/QwenPaw/pull/5578)). There are no migration notes beyond the warning about breaking changes. Users on stable 1.x installations are advised to remain on the old series until 2.0 reaches release candidate.

## Project Progress

Nine pull requests were merged or closed today, reflecting progress in several areas:

- **Test infrastructure** – The backend unit-test coverage plan continued with three PRs merged:
  - [#5580](https://github.com/agentscope-ai/QwenPaw/issues/5580) (closed) – adds test plan for `app` infra layer.
  - [#5423](https://github.com/agentscope-ai/QwenPaw/pull/5423) (merged) – 51 test cases for the crons module.
  - [#5422](https://github.com/agentscope-ai/QwenPaw/pull/5422) (merged) – 47 test cases for the runner module.
  - Multiple frontend test PRs (e.g., [#5409](https://github.com/agentscope-ai/QwenPaw/pull/5409), [#5434](https://github.com/agentscope-ai/QwenPaw/pull/5434), [#5438](https://github.com/agentscope-ai/QwenPaw/pull/5438)) were also merged, adding 300+ cases across stores, hooks, inbox, and API modules.

- **Desktop lifecycle** – [#5265](https://github.com/agentscope-ai/QwenPaw/pull/5265) (merged) implements graceful shutdown for the desktop app (Tauri) by sending a raw TCP shutdown request and respecting a 5-second deadline, fixing abrupt termination.

- **Model management** – [#5297](https://github.com/agentscope-ai/QwenPaw/pull/5297) (merged) adds batch test and batch delete endpoints for models, leveraging `asyncio.gather` for parallel execution.

- **File upload** – [#5436](https://github.com/agentscope-ai/QwenPaw/pull/5436) (merged) enables drag-and-drop file upload onto the chat sender area, a frequently requested UX improvement.

- **CI fixes** – [#5578](https://github.com/agentscope-ai/QwenPaw/pull/5578) (open, but likely to merge quickly) removes `BOOTSTRAP.md` after Tauri initialization to unblock Windows/macOS package verification.

## Community Hot Topics

The most active discussions today reveal three major user pain points:

1. **Message aggregation** – Issue [#5563](https://github.com/agentscope-ai/QwenPaw/issues/5563) (5 comments, 0 👍) describes how multi-step agent actions send one message per step, causing chat spam. Users want an option to batch or collapse intermediate outputs. This has already triggered two related PRs: [#5577](https://github.com/agentscope-ai/QwenPaw/pull/5577) adds an `aggregate_message_replies` channel setting, and [#5575](https://github.com/agentscope-ai/QwenPaw/pull/5575) makes the no-text debounce configurable.

2. **SSH plugin installation loop** – [#5550](https://github.com/agentscope-ai/QwenPaw/issues/5550) (4 comments) reports that the Remote SSH plugin (0.1.3) enters an infinite dependency-installation loop on macOS, plus leaves stale backend processes. This is a blocker for remote development workflows and is tagged as a bug with high impact.

3. **File generation rendering** – [#4865](https://github.com/agentscope-ai/QwenPaw/issues/4865) (3 comments, 2 👍) has been open since June 1 but remains a hot issue. When an agent uses `write_file` to produce large files (HTML, Python, etc.), the web console does not stream the content token-by-token; users see a long loading spinner before the file appears. This is a top-voted UX issue.

A wildcard is [#5567](https://github.com/agentscope-ai/QwenPaw/issues/5567) – a community member submitted a Skill that helps users convert verbal bug reports into standard GitHub issues, showing active community tooling development. The issue has 1 👍 and 2 comments.

## Bugs & Stability

| Severity | Issue | Description | Fix PR Exists? |
|----------|-------|-------------|----------------|
| **Critical** | [#5579](https://github.com/agentscope-ai/QwenPaw/issues/5579) | Conversation records are lost entirely under abnormal interruption (host reboot, service crash). No checkpoint persistence. Users lose all progress. | No |
| **High** | [#5550](https://github.com/agentscope-ai/QwenPaw/issues/5550) | Remote SSH plugin dependency installation loop + stuck old backend processes on macOS. | No |
| **Medium** | [#5573](https://github.com/agentscope-ai/QwenPaw/issues/5573) | DeepSeek V4 thinking mode causes 400 errors on OpenAI-compatible endpoints due to missing `reasoning_content` stream field and null type schemas. | No |
| **Medium** | [#5566](https://github.com/agentscope-ai/QwenPaw/issues/5566) | Cron tasks cannot be silent (empty reply still generates notification); `channels send` command unreachable from background scripts. | Partial: [#5575](https://github.com/agentscope-ai/QwenPaw/pull/5575) addresses debounce behavior. |
| **Low** | [#5561](https://github.com/agentscope-ai/QwenPaw/issues/5561) | Long reply messages to Feishu bot are not delivered as text; only file attachments work. | No |

Additionally, the v2.0.0-beta.1 release has its own verification checklist ([#5571](https://github.com/agentscope-ai/QwenPaw/issues/5571)), which was not completed within the 4-hour deadline, suggesting the beta may not have passed all installation checks.

## Feature Requests & Roadmap Signals

The following features were requested in the last 24 hours and are likely candidates for the next minor release:

- **Automatic message aggregation for multi-step agent responses** (`#5563`) – Already has a work-in-progress PR (`#5577`). Expect it in v2.0.0-beta.2 or later.
- **Model auto-fallback** (`#5572`) – Users want automatic fallback to backup models when primary model hits quota, timeout, or failure. No PR yet, but the need is clear for production deployments.
- **Streaming file content rendering** (`#4865`) – Requested over three weeks ago, but gaining traction. A fix would require changes to the web console tool-call display layer.
- **Cron channel improvements** (`#5566`) – Ability to suppress empty notifications and to trigger notifications from background scripts. The debounce config PR (`#5575`) partially addresses this, but more may be needed.
- **‘Scroll’ context management** (`#5321`) – A large PR adding retrieval-driven conversation history is still under review; if merged, it would provide an alternative to compression-based context management.

The roadmap appears to prioritize **stability** (unit tests, graceful shutdown) and **channel UX** (aggregation, debounce). Community members are actively contributing features (e.g., drag-and-drop upload, batch model management) that improve daily workflows.

## User Feedback Summary

Today’s feedback paints a mixed picture. Users are excited about the project’s direction (new beta, plugin ecosystem, Skill marketplace) but frustrated by **stability gaps**:

- **Pain point #1 – Conversation reliability** – Two separate issues (`#5579`, `#5566`) highlight that conversations can be silently lost or interrupted, destroying user trust in long-running or unattended tasks.
- **Pain point #2 – Integration friction** – DeepSeek V4 compatibility (Model-as-a-Service), Feishu channel limits, and SSH plugin installation loops show that the platform’s integration layer has rough edges, especially on macOS and with non-standard endpoints.
- **Pain point #3 – UX inconsistency** – Message spam during multi-step tasks (`#5563`) and lack of streaming feedback for file generation (`#4865`) make the agent feel opaque and slow.
- **Positive signals** – Users are taking the time to submit detailed bug reports and even build utilities (e.g., the Issue helper Skill `#5567`). The first-time contributor count (PRs `#5524`, `#4622`, `#5321`, `#5574`) indicates a growing, engaged community.

## Backlog Watch

Several items remain open for extended periods and may need maintainer attention:

| Item | Age | Status | Notes |
|------|-----|--------|-------|
| [#4865](https://github.com/agentscope-ai/QwenPaw/issues/4865) – Streaming file content | 26 days | Open, no PR | Top-voted UX issue (2 👍). No assignee. |
| [#4622](https://github.com/agentscope-ai/QwenPaw/pull/4622) – DataPaw plugin (12 BI skills) | 36 days | Open, “Under Review” | First-time contributor. Has been waiting for review since May 22. |
| [#5567](https://github.com/agentscope-ai/QwenPaw/issues/5567) – GitHub Issue helper Skill | 1 day | Open, 2 comments | Community-created Skill; no official response yet. |
| [#5571](https://github.com/agentscope-ai/QwenPaw/issues/5571) – Release verification | 1 day | Open, 0 comments | Beta release checklist not completed within deadline; may indicate need for automated verification. |
| [#5321](https://github.com/agentscope-ai/QwenPaw/pull/5321) – Scroll context manager | 8 days | Open, “Under Review” | Large feature PR (retrieval-driven context). No recent updates. |

**Recommendation**: Prioritize review of `#4622` (DataPaw plugin) to avoid demotivating first-time contributors, and assign an owner to `#4865` as it is the most-voted UX issue. The conversation loss bug (`#5579`) should be treated as critical and escalated, as it directly affects data integrity.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-06-27

## 1. Today’s Overview
ZeroClaw saw high activity on 2026-06-27, with **31 issues updated** (23 open, 8 closed) and **50 pull requests updated** (44 open, 6 merged/closed). The project is deep in the **v0.8.3 milestone** (runtime, Wasm plugins, channels, and CI) while also preparing **v0.9.0** (auth, security, gateway). No new releases were published today. Major themes include supply‑chain security (SLSA provenance, hardware PGP), Wasm‑first plugin runtime, onboarding improvements, and several critical bug fixes. The community is actively contributing to feature requests (Goal Mode, OpenRouter fallbacks, WhatsApp passive context) and stability patches.

## 2. Releases
No new releases were created today. The latest published version remains **v0.8.2** (see previous digest). No migration notes are available.

## 3. Project Progress
**Merged/closed PRs (6 total):**
- **#8330** — `fix(zerocode): render only the viewport in long sessions` — merged; fixes performance degradation in long conversation transcripts.
- **#8381** — `fix(tools): drop unwrap from hardware_memory_read chip lookup` — merged; replaces a panic‑prone `unwrap` with proper error handling.
- **#8371–#8378** — A series of closed issues implementing the **dms‑gst‑extraction agent** (bootstrap milestone). All six user stories (US1–US6) plus documentation polish were completed. These are labelled `invalid` (closed as completed tasks), not bugs.

**Open PRs advancing features:**
- **#8033** (XL, open) — Two‑path onboard tree (LLM + deterministic) wired end‑to‑end over RPC and CLI. This is the largest upcoming onboarding improvement.
- **#8368** (open) — Wasmtime component‑model host for tool/channel/memory plugins (replaces Extism).
- **#8389** (open) — Passive WhatsApp group context (companion to issue #8379).
- **#8384** (open) — Native Inkbox channel (email, SMS, voice, iMessage) with Quickstart onboarding.
- **#8382** (open) — Persist embedding identity and auto‑migrate vectors on model change (fixes #7948).
- **#8277** (open) — SLSA provenance attestation for release pipeline (Phase A of #8177).

**Trackers updated:** #8071 (v0.8.3 runtime), #8360 (provider serialization), #8362 (channel adapter parity), #8363 (config‑driven policy), #8073 (observability/CI/docs), #8070 (gateway/web/zeroCode).

## 4. Community Hot Topics
The most active discussions (by comment count) reveal strong interest in security, reliability, and new capabilities:

| Issue/PR | Comments | Topic | Link |
|----------|----------|-------|------|
| **#8177** [RFC] | 10 | Supply‑chain signing: hardware PGP, hermetic builds, SLSA provenance. Core to v0.9.0 security roadmap. | [URL](https://github.com/zeroclaw-labs/zeroclaw/issues/8177) |
| **#5808** [Bug] | 6 | Default 32k context budget exceeded by system prompt + tool definitions on iteration 1, causing perpetual preemptive trim. S1 severity. | [URL](https://github.com/zeroclaw-labs/zeroclaw/issues/5808) |
| **#8138** [Feature] | 3 | Support OpenRouter model fallbacks array in provider config. | [URL](https://github.com/zeroclaw-labs/zeroclaw/issues/8138) |
| **#8303** [RFC] | 2 | Goal mode for bounded autonomous session work (pursue objective until completion, pause, cancellation). | [URL](https://github.com/zeroclaw-labs/zeroclaw/issues/8303) |
| **#8135** [RFC] | 2 | Wasm‑first plugin runtime — default‑on, capability enforcement, signed distribution. | [URL](https://github.com/zeroclaw-labs/zeroclaw/issues/8135) |
| **#8379** [Feature] | 2 | Opt‑in passive group context for WhatsApp Web group chats. | [URL](https://github.com/zeroclaw-labs/zeroclaw/issues/8379) |

Users are pushing for **better security defaults** (#8177, #8135), **reliable model switching** (#8138), **long‑running autonomous tasks** (#8303), and **multi‑channel context awareness** (#8379). The high engagement on #5808 indicates the context budget bug is a top pain point.

## 5. Bugs & Stability
**New bugs reported today (ranked by severity):**

| Issue | Severity | Description | Fix PR/Status |
|-------|----------|-------------|---------------|
| #8385 | **S1** – workflow blocked | ZeroCode transcript message highlight traps input after mouse selection; user cannot type again. | No fix PR yet. |
| #8386 | **S2** – degraded behavior | SQLite default memory backend never prompts for embedding model; hybrid search silently degrades to keyword‑only. | No fix PR yet. |
| #8360 (tracker) | – | Multiple provider‑side serialization bugs (part of v0.8.3). | Being addressed in #8355, #8353, #8350. |

**Existing critical bugs with ongoing fixes:**
- **#5808** (S1, context budget overflow) — Fix PR **#7440** (open) rebased onto master; implements remediation (detect floor > budget and block agent start).
- **#6350** (WhatsApp LID→phone resolution) — Fix PR **#6622** (open, needs‑author‑action, stale‑candidate). Persistent store approach.
- **#6434** (shell authorization at `autonomy.level=full`) — Fix PR **#6619** (open, needs‑author‑action, stale‑candidate).

**Stability improvements merged today:** #8330 (viewport rendering fix for long sessions), #8381 (unwrap removal in hardware tools), #8353 (better error messages and unwrap→expect in runtime), #8388 (unwrap→expect in tool‑call parser).

## 6. Feature Requests & Roadmap Signals
**New feature requests today:**
- **#8387** — Add daemon restart controls to ZeroCode.
- **#8383** — Show active runtime context in ZeroCode Dashboard.
- **#8367** (RFC) — Capability‑aware documentation for agent‑visible features.

**Active high‑demand features (likely for v0.8.3 or v0.9.0):**
- **Goal Mode** (#8303) — First‑class durable autonomous sessions. Strong community support.
- **OpenRouter fallbacks** (#8138) — Simple config change to enable model failover.
- **Wasm plugin runtime** (#8135, #8368) — Default‑on, signed distribution. The PR #8368 is a major implementation.
- **Supply‑chain signing** (#8177, #8277) — SLSA L2 provenance already in progress.
- **Embedding identity persistence** (#7948, PR #8382) — Needed to avoid vector corruption on model change.

**What might land next:** The v0.8.3 tracker (#7320) is broad, but the **onboarding tree** (#8033) and **Wasm plugin host** (#8368) are the largest patches. A point release (0.8.3) could ship before the v0.9.0 security overhaul.

## 7. User Feedback Summary
**Pain points reported by users (directly from issue descriptions):**
1. **Context budget thrashes on first turn** (#5808) — “First LLM iteration already exceeds budget by ~3.3x purely from system prompt + tool definitions.” Users are blocked from using default config.
2. **Missing embedding onboarding** (#8386) — Users following quickstart get degraded search (keyword‑only) without realizing they need a separate embedding model step.
3. **ZeroCode UI regressions** (#8385) — Mouse selection traps input, blocking typing entirely in the TUI.
4. **WhatsApp group messages dropped** (#8379) — Messages not addressed to bot are simply discarded; users want them stored as passive context.
5. **WhatsApp LID allowlist bypass** (#6350) — Operators using phone numbers in allowlists see messages silently dropped when Meta uses LIDs.

**Use cases driving new features:**
- **Enterprise security (#8177)** — Operators need verifiable supply chain provenance for audit compliance.
- **Multi‑model resilience (#8138)** — Avoiding downtime when primary provider is rate‑limited.
- **Autonomous long‑running workflows (#8303)** — “Users need goal mode that can start from a single prompt and run until completion, pause, or failure.”
- **Multi‑channel bots (#8384, #8379)** — Native iMessage/SMS/email integration (Inkbox) and group context for WhatsApp.

**Satisfaction signals:** The community is actively contributing fixes and features (many PRs from first‑time contributors). The structured milestone trackers (#7432, #8071) help maintain transparency. No widespread dissatisfaction is visible beyond the critical bugs.

## 8. Backlog Watch
The following important items have been open for **over a month** or lack maintainer action:

| Item | Type | Opened | Status | Notes |
|------|------|--------|--------|-------|
| **#6622** – WhatsApp LID→phone resolution | PR (fix) | 2026-05-13 | `needs-author-action`, `stale-candidate` | Fix for #6350. Maintainer review needed. |
| **#6619** – Shell authorization at `full` autonomy | PR (fix) | 2026-05-13 | `needs-author-action`, `stale-candidate` | Fix for #6434. Blocked on author response. |
| **#5808** – Context budget overflow | Issue (S1) | 2026-04-16 | `in-progress` | Fix PR #7440 open; needs review and merge. |
| **#8177** – Supply‑chain signing RFC | Issue | 2026-06-22 | `blocked` | Requires hardware PGP key management. |
| **#8350** – Regex cache in web‑search | PR (fix) | 2026-06-26 | `needs-author-action` | Small performance fix, author must address CI. |
| **#7440** – Context budget remediation | PR (fix) | 2026-06-09 | Open | Detailed fix for #5808; has been rebased multiple times. |

**Maintainer attention requested:** The two stale‑candidate PRs (#6622, #6619) risk being auto‑closed if authors do not respond. The #5808 / #7440 pair is the most impactful open bug; merging the fix would unblock many users. The supply‑chain RFC (#8177) is blocked on design decisions and hardware key provisioning.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*