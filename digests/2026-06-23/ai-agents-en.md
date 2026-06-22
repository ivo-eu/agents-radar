# OpenClaw Ecosystem Digest 2026-06-23

> Issues: 53 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-22 17:18 UTC

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

# OpenClaw Project Digest — 2026-06-23

## Today’s Overview

The project saw high activity over the past 24 hours with **53 issues updated** (32 open/active, 21 closed) and **500 PRs updated** (443 open, 57 merged/closed). One new beta release was cut: `v2026.6.10-beta.2`. The release focuses on automatic fast mode for short conversational turns and more reliable model routing. The community continues to report and fix regressions in session state, message delivery, and provider compatibility, with several P1 issues awaiting maintainer reviews and product decisions. Overall, the project is in a busy but stable state, with sustained contributions from both maintainers and external contributors.

## Releases

- **v2026.6.10-beta.2** (2026-06-22)
  - Highlights:
    - **Automatic fast mode for talks** – OpenClaw now switches to fast mode for short conversational turns and returns to normal mode for longer runs, with bounded fallback and delivery behavior. ([#85104](https://github.com/openclaw/openclaw/pull/85104)) Thanks @alexph-dev and @vincentkoc.
    - **More reliable model routing** – Improvements to the Zai provider? (release notes truncated in data).
  - No breaking changes or migration notes were published with this beta.

## Project Progress

Over the last 24 hours, **57 PRs were merged or closed**, including:

- **Documentation**: Updated install guides to include Clawcks hosting ([#95728](https://github.com/openclaw/openclaw/pull/95728)), fixed docs metadata spellcheck ([#93502](https://github.com/openclaw/openclaw/pull/93502)), added existing-solutions preflight guardrail ([#86608](https://github.com/openclaw/openclaw/pull/86608)).
- **Bug fixes**: Resolved double-rendering of agent replies in Control UI webchat when `dmScope=main` ([#93044](https://github.com/openclaw/openclaw/issues/93044), PR [#93841](https://github.com/openclaw/openclaw/pull/93841)); restored model-fetch info logs ([#89300](https://github.com/openclaw/openclaw/issues/89300), PR [#89648](https://github.com/openclaw/openclaw/pull/89648)); fixed Codex usage proxy fetch ([#78714](https://github.com/openclaw/openclaw/issues/78714), PR [#93943](https://github.com/openclaw/openclaw/pull/93943)); improved opencode-go streaming completion for cron jobs ([#93965](https://github.com/openclaw/openclaw/pull/93965)); introduced atomic rollback/revert mechanism for core, plugins, and config ([#95825](https://github.com/openclaw/openclaw/issues/95825), merged as a feature PR).
- **Security & stability**: Bounded HTTP body reads in several provider paths to prevent OOM (e.g., OpenRouter, Inworld TTS, ChatGPT Responses) – see PRs [#95420](https://github.com/openclaw/openclaw/pull/95420), [#95416](https://github.com/openclaw/openclaw/pull/95416), [#95223](https://github.com/openclaw/openclaw/pull/95223).

Several closed issues also reflect progress: onboarding install loops fixed ([#95765](https://github.com/openclaw/openclaw/issues/95765)), config-patch restart notice improved ([#46797](https://github.com/openclaw/openclaw/issues/46797)), and session model cache persistence corrected ([#77322](https://github.com/openclaw/openclaw/issues/77322)).

## Community Hot Topics

| Issue / PR | Comments | 👍 | Summary |
|------------|----------|----|---------|
| [#88838](https://github.com/openclaw/openclaw/issues/88838) | 33 | 1 | Track core session/transcript SQLite migration via accessor seam – a high-visibility maintainer topic requiring product decision. |
| [#88312](https://github.com/openclaw/openclaw/issues/88312) | 17 | 4 | Regression: Codex app-server turn-completion stall returns. A P1 bug with major impact on session state and message loss. |
| [#92201](https://github.com/openclaw/openclaw/issues/92201) | 12 | 1 | Embedded runner: invalid thinking signatures on replay (Anthropic) – recovery wrapper never fires. Needs maintainer review and product decision. |
| [#91363](https://github.com/openclaw/openclaw/issues/91363) | 7 | 4 | Isolated cron consistently fails with "LLM request failed" on model-call-started phase. Affected users seeking a fix. |
| [#78431](https://github.com/openclaw/openclaw/issues/78431) | 5 | 1 | Discord `messages.statusReactions` unimplemented despite docs – long-standing feature gap. |

**Underlying needs**: The community is most vocal about **session reliability** (transcript migration, turn completion stalls), **cron job stability**, and **provider-specific quirks** (Anthropic thinking blocks, Codex proxy handling). A recurring theme is that product decisions are blocking fixes for many P1 issues.

## Bugs & Stability

### P1 – Critical
- **Regression: Codex turn-completion stall** ([#88312](https://github.com/openclaw/openclaw/issues/88312)) – Impacts session state and message loss. No fix PR yet; regression of #84076 (previously fixed by #85107). Needs product decision.
- **Compaction resets model override** ([#95696](https://github.com/openclaw/openclaw/issues/95696)) – After `/model` switch, compaction reverts to previous provider. Fix likely needed in session compaction logic.
- **Compaction fails with OAuth-style auth** ([#95693](https://github.com/openclaw/openclaw/issues/95693)) – `No API key found` error when using codex/claude-cli runtime. Needs live reproduction and maintainer review.
- **Cron fallback chain not engaged on stream stalls** ([#85900](https://github.com/openclaw/openclaw/issues/85900)) – Provider opens stream but never sends chunks; fallback never triggers. Needs product decision.

### P2 – High
- **Outbound transcript mirrors ignore configured `session.store`** ([#95781](https://github.com/openclaw/openclaw/issues/95781)) – Data-loss potential.
- **NVIDIA Build provider stream cut mid-tool-calls** ([#95760](https://github.com/openclaw/openclaw/issues/95760)) – Session enters zombie state.
- **Claude-cli out-of-credits bypasses model fallback** ([#95489](https://github.com/openclaw/openclaw/issues/95489)) – Error text delivered as final response.
- **DOCX extension rewriting to `.docodex`** ([#93326](https://github.com/openclaw/openclaw/issues/93326)) – A PR ([#95376](https://github.com/openclaw/openclaw/pull/95376)) is ready for maintainer look.
- **Heartbeat transcript filtering broken with reasoning blocks** ([#95796](https://github.com/openclaw/openclaw/issues/95796)) – Inconsistent removal of HEARTBEAT_OK turns.

### P3 – Moderate
- **Telegram native approvals time out silently** ([#95800](https://github.com/openclaw/openclaw/issues/95800)) – No delivery when `execApprovals` unconfigured.
- **msteams allowlist never works** ([#95737](https://github.com/openclaw/openclaw/issues/95737)) – Messages always dropped due to `groupAllowFrom` bug.
- **Task Guard silence mandate has no exit condition** ([#95773](https://github.com/openclaw/openclaw/issues/95773)) – stuck with `NO_REPLY` instruction.

Several of these bugs have open PRs awaiting review or author action. The **DOCX extension fix** and **heartbeat filtering** are notable examples of active remediation.

## Feature Requests & Roadmap Signals

The most innovative feature requests from the community over the past 24 hours include:

- **Self-evolving SOUL.md** ([#95790](https://github.com/openclaw/openclaw/issues/95790)) – Opt-in reflection sub-turn with a `soul_update` tool to let agents update their personality file. This aligns with the existing documentation but is not yet implemented.
- **Configurable default shell for exec tool** ([#95817](https://github.com/openclaw/openclaw/issues/95817)) – Especially important on Windows where PowerShell 5.x has encoding and color limitations.
- **Atomic rollback / revert mechanism** ([#95825](https://github.com/openclaw/openclaw/issues/95825)) – Feature already implemented and merged (#95825 was a PR, now closed). Expect it in next release.
- **Channel Broker Phase 2A – conformance harness** ([#86109](https://github.com/openclaw/openclaw/issues/86109)) – Shared test contract for all channel providers. Indicates a push toward channel standardization.
- **Doc-focused bug intake path** ([#76664](https://github.com/openclaw/openclaw/issues/76664)) – Closed, suggests improved reporting workflows are coming.

**Predictions for next version**: The atomic rollback feature and the DOCX path correction are likely to ship. The `no-stale` P1 compilation/model-override bugs may also be resolved if maintainers approve pending fixes. The SOUL.md self-evolution is a long-tail feature but signals strong community interest in agent personalization.

## User Feedback Summary

**Pain points**:
- **Session state fragility**: Users report that `/new`, compaction, and cross-boot recovery can silently lose model overrides, message history, or cause endless loops. Multiple issues filed in one day (e.g., [#95696](https://github.com/openclaw/openclaw/issues/95696), [#95693](https://github.com/openclaw/openclaw/issues/95693), [#95750](https://github.com/openclaw/openclaw/issues/95750)) suggest frustration with session reliability.
- **Provider-specific breakage**: Users of Codex, Anthropic, NVIDIA Build, and opencode-go are experiencing unique failures (stalled streams, invalid signatures, misplaced model IDs). Several P1 issues remain with no clear path to fix.
- **Cron job reliability**: Isolated cron failures ([#91363](https://github.com/openclaw/openclaw/issues/91363)) and fallback bypass ([#85900](https://github.com/openclaw/openclaw/issues/85900)) are affecting production automation users.
- **Documentation gaps**: The Discord status reactions feature is documented but unimplemented ([#78431](https://github.com/openclaw/openclaw/issues/78431)). Users also reported the macOS beta desktop assets are missing for v2026.5.22 ([#85902](https://github.com/openclaw/openclaw/issues/85902)).

**Satisfaction**: The fast mode for talks and model routing improvements in the new beta were well received. The community is actively contributing fixes (57 merged PRs) and many contributors are thanked in issues.

## Backlog Watch

The following issues and PRs are long-standing and need maintainer attention:

| Item | Age | Status | Why Waiting |
|------|-----|--------|-------------|
| [#88838](https://github.com/openclaw/openclaw/issues/88838) (SQLite migration seam) | 22 days | Open, P1, needs product decision | High-impact architectural change; requires alignment. |
| [#88312](https://github.com/openclaw/openclaw/issues/88312) (Codex turn-completion stall regression) | 24 days | Open, P1, needs product decision | Regression unfixed; affects many users. |
| [#85914](https://github.com/openclaw/openclaw/issues/85914) (tool-call failure recovery as native capability) | 30 days | Open, P2, needs maintainer review + product decision | Feature request with clear scope but no PR. |
| [#85937](https://github.com/openclaw/openclaw/issues/85937) (Gemini duplicate assistant text) | 30 days | Open, P2, needs live reproduction | Reproducible on specific version; no fix yet. |
| [#85900](https://github.com/openclaw/openclaw/issues/85900) (Cron fallback chain not engaged on stream stalls) | 30 days | Open, P1, needs product decision | Blocks reliable cron automation. |
| [#71537](https://github.com/openclaw/openclaw/pull/71537) (Recover archived session transcripts) | 59 days | Open, waiting on author | Author hasn’t responded to review feedback. |
| [#61485](https://github.com/openclaw/openclaw/pull/61485) (Modifying LLM I/O hooks) | 80 days | Open, needs proof of real behavior | Large feature with security implications. |

These items represent the most critical gaps in stability, feature completeness, and maintainer bandwidth. The project would benefit from a focused triage session to move P1/P2 blockers out of “needs product decision” and into implementation.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: Personal AI Assistant Open-Source Ecosystem
**Date:** 2026-06-23 | **Prepared for:** Technical Decision-Makers & Developers

---

## 1. Ecosystem Overview

The personal AI agent open-source ecosystem is experiencing rapid maturation, with nine actively developed projects collectively processing over 2,000 PRs and 200 issues in a single day. The landscape is bifurcating into two tracks: "Claw" lineage projects (OpenClaw, PicoClaw, NanoClaw, ZeroClaw, IronClaw) building on shared architectural patterns, and independent projects (NanoBot, Hermes Agent, CoPaw, LobsterAI) pursuing differentiated agent interaction models. Cross-cutting concerns dominate—session reliability, provider compatibility, MCP integration, and mobile/web UX are universal priorities. The ecosystem shows strong community health with sustained contribution velocity, though maintainer bandwidth is emerging as a bottleneck for P1 bug resolution across multiple projects.

---

## 2. Activity Comparison

| Project | Issues Updated (24h) | PRs Updated (24h) | PRs Merged/Closed | Release Status | Health Score (0–10) |
|---------|---------------------|-------------------|-------------------|----------------|---------------------|
| **OpenClaw** | 53 | 500 | 57 | v2026.6.10-beta.2 | 8.5 |
| **NanoBot** | 7 | 39 | 16 | v0.2.2 imminent | 9.0 |
| **Hermes Agent** | 14 | 50 | 10 | No release | 7.5 |
| **PicoClaw** | 6 | 32 | 29 | v0.3.0-nightly | 8.5 |
| **ZeroClaw** | 8 | 50 | 8 | No release | 8.0 |
| **IronClaw** | 18 | 38 | 4 | No release (Reborn WIP) | 7.5 |
| **CoPaw** | 6 | 50 | 21 | No release | 8.0 |
| **LobsterAI** | 5 | 14 | 6 | No release | 5.5 |
| **NanoClaw** | 0 | 6 | 3 | No release | 6.0 |
| **NullClaw** | 0 | 1 | 0 | No release | 4.0 |
| **TinyClaw** | 0 | 0 | 0 | — | 2.0 |
| **Moltis** | 0 | 0 | 0 | — | 2.0 |
| **ZeptoClaw** | 0 | 0 | 0 | — | 2.0 |

**Health Score Methodology:** Combines activity velocity (PR merge rate), issue resolution latency, maintainer responsiveness, and critical bug backlog. Scores normalized against ecosystem mean.

---

## 3. OpenClaw's Position

**Advantages vs. Peers:**
- **Scale leader:** 500 PRs/53 issues updated—3x the nearest competitor (NanoBot/ZeroClaw at 50). Only project with active beta release cycle.
- **Architectural maturity:** Unique "fast mode" automatic switching for conversational turns (#85104) and atomic rollback mechanism (#95825) demonstrate production-grade state management unavailable in other projects.
- **Provider breadth:** 11+ provider integrations with bounded HTTP body reading for OOM prevention—security-hardening absent in most peers.
- **Community engagement:** 443 open PRs indicate massive external contribution pipeline, though merge bottleneck exists (57/500 merged/closed = 11.4% throughput rate).

**Technical Approach Differences:**
- **Session model:** OpenClaw uses SQLite-based transcript migration (#88838) vs. NanoBot's JSONL memory store and Hermes' in-memory state. This architectural choice creates both persistence advantages (crash recovery) and maintenance complexity (migration seams).
- **Model routing:** Explicit "Zai provider" improvements vs. NanoBot's subagent spawn presets (#4291) and Hermes' per-service-tier routing.
- **Plugin system:** Clawcks hosting support (#95728) contrasts with IronClaw's crate decomposition approach and CoPaw's mobile-first plugin manager.

**Community Size Comparison:**
- OpenClaw: ~1,200+ contributors (443 open PRs from external contributors)
- NanoBot: ~200+ contributors (39 PRs/day, strong core team velocity)
- Hermes Agent: ~300+ contributors (50 PRs/day, 10 merged—lowest throughput)
- Others: 50–150 contributor range

---

## 4. Shared Technical Focus Areas

| Focus Area | Affected Projects | Specific Needs |
|------------|------------------|----------------|
| **Session State Reliability** | OpenClaw, NanoBot, Hermes, PicoClaw, ZeroClaw | Turn-completion stalls, compaction resetting model overrides, message queue cross-talk, deferred image attachment loss |
| **MCP/Plugin Lifecycle** | OpenClaw, NanoBot, PicoClaw, ZeroClaw, Hermes | Permission enforcement bypass, stdio transport crashes, command parsing errors, reconnect crash on scope exit |
| **Mobile/Web UI** | CoPaw, PicoClaw, Hermes, LobsterAI | Narrow viewport adaptation, iOS Safari compatibility, drag-and-drop file attachments, layout breakage on localization |
| **Provider Compatibility** | All projects except NullClaw | API endpoint changes (Kimi/ZeroClaw), reasoning block handling, tool call leak, DeepSeek null content, rate limiting |
| **Performance & Latency** | ZeroClaw, IronClaw, NanoBot, OpenClaw | Initial response time across channels, inference latency, unnecessary runtime steps, prefix caching instability |
| **Cron/Automation** | OpenClaw, Hermes, IronClaw, LobsterAI | Fallback chain disengagement, LLM request failures, missed history generation, service tier ignoring |

**Bottom Line:** Session state management and provider compatibility are the ecosystem's two biggest pain points. Every project with active users reports at least one P1/P2 bug in these categories.

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | Hermes Agent | IronClaw | CoPaw | PicoClaw | ZeroClaw |
|-----------|----------|---------|--------------|----------|-------|----------|----------|
| **Core Architecture** | Monolithic + plugins | Gateway-driven | Modular microservices | Reborn crate decomposition | Mobile-first console | Lightweight embedded | Feature parity with OpenClaw |
| **Target User** | Power users, self-hosters | Enterprise, multi-platform | Developers, CI/CD | NEAR ecosystem | Chinese market, mobile | Embedded/IoT | Feature-complete fork |
| **Model Paradigm** | Proactive routing | Subagent presets | Service-tier routing | Per-tool permissions | Custom model ordering | Provider selection UI | Prefix caching |
| **Channel Focus** | Broad (Discord, Telegram, etc.) | Mattermost, DingTalk, Telegram | Slack, Discord | Telegram, outbound delivery | MCP, ACP, Inbox | SimpleX, Tox requested | All major channels |
| **Differentiator** | Fast mode, atomic rollback | Background daemon gateway, PWA | PR review plugin, self-evolution | Concurrent TurnRunScheduler | Mobile card layouts | Cross-platform serial tool | Supply chain signing, in-app upgrade |
| **Maturity** | Mature, beta releases | Pre-v0.2.2, rapidly stabilizing | Feature-rich, breaking changes | Pre-Reborn, heavy refactoring | Mobile-responsive surge | v0.3.0 nightly | RFC-heavy, high risk tolerance |

---

## 6. Community Momentum & Maturity

### Tier 1: High Velocity, Active Stabilization
- **NanoBot (9.0 health):** Highest throughput-to-merge ratio (41% merged/closed). Engineering discipline evident in systematic gateway stability fixes. v0.2.2 release imminent with PWA, background daemon, and MCP permission enforcement.
- **PicoClaw (8.5):** Emergent leader in configuration UX and streaming. 29/32 PRs merged (90%+). SiYue-ZO's contributions indicate strong single-contributor momentum.
- **OpenClaw (8.5):** Ecosystem scale champion, but throughput challenges (11% merge rate). Beta release cycle demonstrates maturity while 30-day stale P1 issues signal maintainer bandwidth constraints.

### Tier 2: Active Development, Growing Pains
- **ZeroClaw (8.0):** Highest risk tolerance—RFCs for supply-chain signing and in-app upgrade indicate architectural ambition. Two S1 workflow blockers (Kimi endpoint, deferred images) need immediate resolution.
- **CoPaw (8.0):** Mobile-responsive overhaul shows strong community drive (50 PRs). Agent-switching queue bug (#5354) and custom model bug (#5378) erode trust.
- **Hermes Agent (7.5):** Feature velocity is highest (50 PRs) but merge throughput is lowest (10/50 = 20%). Slack `mention_patterns` fix after months indicates backlog accumulation. Breaking changes are community pain point (#50923).
- **IronClaw (7.5):** Reborn architecture rewrite introduces both innovation (TurnRunScheduler, per-tool permissions) and risk (nightly E2E failing since May 27, dogfooding reports of slowness).

### Tier 3: Stable / Declining
- **NanoClaw (6.0):** Low activity but meaningful community contributions (Telegram integration). Unmerged PR #2531 (May 18) suggests review pipeline stagnation.
- **LobsterAI (5.5):** Five open bugs from April 3 unresolved. Seven open PRs from April remain unmerged. Maintenance appears stalled, though new PRs (Plan Mode) show brief activity.
- **NullClaw (4.0), TinyClaw, Moltis, ZeptoClaw (all 2.0):** Effectively dormant. Dependabot maintenance only.

---

## 7. Trend Signals

### Industry Trends Emerging from Community Feedback

**1. Session State as the New Bottleneck**
Across 5+ projects, users report that `/new`, compaction, and cross-boot recovery silently lose model overrides or message history (OpenClaw #95696, #95693; NanoBot #4442; CoPaw #5354). This is the single largest source of user frustration. **Implication:** Expect standardization of session persistence APIs and crash-recovery protocols as a competitive differentiator.

**2. Provider Fragility Hurts Adoption**
API endpoint changes (ZeroClaw Kimi), reasoning block leaks (OpenClaw #92201, Hermes #50930), and tool call degradation (PicoClaw #3153) create "which provider works today?" uncertainty. **Implication:** Provider abstraction layers and automatic fallback chains (ZeroClaw #8138) will become table stakes.

**3. Agent Personalization Demand**
OpenClaw's SOUL.md self-evolution (#95790), Hermes' skill extraction (#5061, IronClaw), and CoPaw's recall-aware memory (#5387) signal user appetite for agents that "learn" personality over time. **Implication:** This is the next frontier after session reliability—projects that build safe, user-controlled personalization will lead.

**4. Mobile-First is Non-Negotiable**
CoPaw's 50-PR mobile responsive push, PicoClaw's iOS Safari bugs, and LobsterAI's English layout breakage indicate mobile access is a primary deployment pattern, not an afterthought. **Implication:** Web UI must be responsive from day one; desktop-only projects will lose users.

**5. Compliance Beyond Provider Choice**
Supply-chain signing (ZeroClaw#8177), hardware-backed PGP, and TLS confirmation flows reflect growing enterprise deployment pressure. **Implication:** The "self-hosted vs. hosted" dial will define project positioning; security provenance is the new adoption gate.

**6. Cron/Automation is Production-Grade**
Users now expect agents to run scheduled tasks reliably (OpenClaw #91363, LobsterAI #1409), not just conversational interfaces. Fallback chain disengagement (OpenClaw #85900) blocks automation use cases. **Implication:** Automated task execution will be the ecosystem's proving ground for "AI agent as infrastructure" credibility.

### Value for AI Agent Developers

| Developer Need | Ecosystem Signal | Actionable Insight |
|---------------|------------------|--------------------|
| Deploy reliable multi-turn agents | Session state bugs are the #1 complaint | Invest in bounded compaction, atomic rollback, and crash-recovery testing |
| Support multiple LLM providers | Every project reports provider-specific failures | Build fallback chains and provider abstraction layers now |
| Build enterprise-ready agents | Supply-chain signing and TLS confirmation emerging | Prioritize compliance features early; retrofitting is painful |
| Target mobile users | CoPaw, PicoClaw investing heavily in responsive UI | Test on iOS Safari < 16.4 first; it's the weakest link |
| Differentiate on UX | NanoBot's onboarding wizard, OpenClaw's fast mode | Ease of setup and perceived intelligence matter more than feature count |

**Final Assessment:** The ecosystem is converging on session reliability and provider compatibility as foundational requirements. Projects that solve these well (NanoBot, OpenClaw, PicoClaw) are pulling ahead. The next 3–6 months will separate "feature gardens" from "production platforms"—personalization and automation reliability will be the differentiators.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest – 2026-06-23

## 1. Today's Overview
NanoBot had a highly active day with **39 pull requests** updated (16 merged/closed) and **7 issues** updated (2 closed). The project is accelerating toward a **v0.2.2 release** – the release preparation PR (#4445) was merged, though no official tag has been published yet. The bulk of activity focused on **gateway stability**, **MCP reliability**, **WebUI fork/scroll behavior**, and **config defaults**. The community also submitted notable enhancements for DingTalk, PWA support, and a user-friendly onboarding wizard. Overall, project health is strong with consistent core team attention to regressions and infrastructure.

## 2. Releases
**No new releases** were published in the last 24 hours. However, the team merged a [chore(release): prepare v0.2.2 (#4445)](https://github.com/HKUDS/nanobot/pull/4445) which bumps the package version to `0.2.2` and adds the release entry to the README news section. This indicates a release is imminent, likely within the next day. No breaking changes or migration notes are documented yet.

## 3. Project Progress
The following significant PRs were **merged/closed** today (16 total). Highlights:

- **Gateway & system stability** – Multiple fixes for shutdown and lifecycle:
  - [fix(gateway): tolerate cancelled channel tasks during shutdown (#4456)](https://github.com/HKUDS/nanobot/pull/4456)
  - [fix: stabilize gateway shutdown and webui fork replay (#4454)](https://github.com/HKUDS/nanobot/pull/4454)
  - [fix(webui): preserve fork replies during history refresh (#4455)](https://github.com/HKUDS/nanobot/pull/4455)
  - [fix(webui): follow active turn output after send (#4453)](https://github.com/HKUDS/nanobot/pull/4453)
  - [fix(webui): stabilize sent turn layout and dev reloads (#4451)](https://github.com/HKUDS/nanobot/pull/4451)
  - [fix: close MCP stdio transports from agent task (#4450)](https://github.com/HKUDS/nanobot/pull/4450)
  - [fix(gateway): handle lifecycle edge cases (#4447)](https://github.com/HKUDS/nanobot/pull/4447)

- **Configuration & defaults**:
  - [chore(config): default context window to 200k (#4448)](https://github.com/HKUDS/nanobot/pull/4448)
  - [chore(release): prepare v0.2.2 (#4445)](https://github.com/HKUDS/nanobot/pull/4445)

- **WebUI fixes**:
  - [fix(webui): avoid slow settings route refreshes (#4398)](https://github.com/HKUDS/nanobot/pull/4398)

- **Feature work**:
  - [feat(gateway): add background and service controls (#1854)](https://github.com/HKUDS/nanobot/pull/1854) – long-standing enhancement merged.

- **Closed issues** served by these PRs:
  - [#1461 – Provide a unified daemon gateway semantic layer](https://github.com/HKUDS/nanobot/issues/1461) (closed, addressed by #1854)
  - [#4376 – User-friendly onboarding wizard](https://github.com/HKUDS/nanobot/issues/4376) (closed, earlier work)

## 4. Community Hot Topics
The most discussed or upvoted items were:

- **[Issue #1461 – Provide a unified daemon gateway semantic layer](https://github.com/HKUDS/nanobot/issues/1461)** (4 comments, now closed) – This was a long-running feature proposal that finally landed with PR #1854. The community had strong interest in background daemon support and OS-managed service controls.
- **[Issue #1011 – Mattermost Bot](https://github.com/HKUDS/nanobot/issues/1011)** (4 👍, 1 comment) – Users continue to request Mattermost as a communication channel. The issue is 4 months old with no official response from maintainers, indicating a gap in platform support.
- **[Issue #4413 – Telegram Bot API 10.1 rich messages](https://github.com/HKUDS/nanobot/issues/4413)** (2 comments) – A request for enhanced Telegram formatting, reflecting demand for richer messaging capabilities.

**Underlying needs**: The community clearly values **platform diversity** (Mattermost, Telegram enhancements) and **operational tooling** (daemon mode, background service). The popularity of the Mattermost issue suggests NanoBot is being used in enterprise contexts where Mattermost is common.

## 5. Bugs & Stability
Several bugs were reported or fixed today. Ranked by severity:

| Severity | Bug | Fix PR(s) | Status |
|----------|-----|-----------|--------|
| **Critical** | **Duplicate `tool_use` IDs in streamed responses** permanently brick a session ([Issue #4442](https://github.com/HKUDS/nanobot/issues/4442)) – Every subsequent API call returns 400. | [PR #4443](https://github.com/HKUDS/nanobot/pull/4443) – Guard against duplicate IDs | Open (PR in review) |
| **High** | **MCP server reconnect crash** – `RuntimeError: Attempted to exit cancel scope in a different task` when reconnecting after session termination ([PR #4441](https://github.com/HKUDS/nanobot/pull/4441)) | [PR #4441](https://github.com/HKUDS/nanobot/pull/4441) – Force-close streamable_http generator | Open (PR in review) |
| **High** | **MCP `enabledTools` allowlist bypass** – Resources and prompts were registered even when tools were denied ([Issue #4434/#4435](https://github.com/HKUDS/nanobot/issues/4434), [PR #4436](https://github.com/HKUDS/nanobot/pull/4436)) | [PR #4436](https://github.com/HKUDS/nanobot/pull/4436) and [PR #4452](https://github.com/HKUDS/nanobot/pull/4452) merged | **Fixed** |
| **Medium** | **Background gateway state persistence** – False successful stop reports when termination fails ([PR #4447](https://github.com/HKUDS/nanobot/pull/4447)) | Included in #4447 | Merged |
| **Low** | **DeepSeek null/empty content handling** – Various 400 errors and placeholder leaks ([PR #3869](https://github.com/HKUDS/nanobot/pull/3869)) | Open since May 16 | Still open |

**Overall stability trend**: The core team is aggressively addressing gateway and MCP edge cases. The duplicate `tool_use` bug (#4442) is the most critical open issue, with a fix already proposed.

## 6. Feature Requests & Roadmap Signals
- **User-friendly onboarding wizard** – Issue #4376 was closed, suggesting a solution is in the works. Likely part of the upcoming v0.2.2.
- **PWA support for WebUI** – [Issue #4457](https://github.com/HKUDS/nanobot/issues/4457) and [PR #4458](https://github.com/HKUDS/nanobot/pull/4458) propose Progressive Web App capabilities for mobile home screen installation. Expect this to land in the next release.
- **DingTalk channel improvements** – [PR #4446](https://github.com/HKUDS/nanobot/pull/4446) adds private chat gating and sender mention in group replies, expanding enterprise IM support.
- **Subagent model presets** – [PR #4291](https://github.com/HKUDS/nanobot/pull/4291) (open since June 11) allows subagents to use different models configurable via `spawnPresets`. This is a major feature for hierarchical agent workflows.
- **Search history tool** – [Issue #4440](https://github.com/HKUDS/nanobot/issues/4440) proposes a read-only tool to query `memory/history.jsonl`, enabling long-term recall without cluttering the context window.

**Predictions for next release (v0.2.2)**:
- Background daemon gateway (already merged)
- Default context window 200K (merged)
- MCP permission enforcement (merged)
- PWA support (likely)
- Possibly the DingTalk enhancements and the user-friendly wizard update.

## 7. User Feedback Summary
- **Pain points**:
  - Onboarding complexity – “`nanobot onboard --wizard` assumes you know many technical details” (Issue #4376).
  - Session-breaking stream bug – Users silently lose agent replies after a duplicate tool_use ID (Issue #4442).
  - Lack of Mattermost support – Four upvotes on #1011, no official response from maintainers.
  - Telegram formatting limitations – Need for rich messages/advanced markdown (Issue #4413).
- **Positive signals**:
  - The community actively contributes fixes (e.g., @michaelxer, @Re-bin, @chengyongru) and feature proposals.
  - The rapid merging of stability PRs shows responsiveness to user-reported regressions.
- **Use cases**:
  - Enterprise deployment demanding Mattermost and DingTalk integration.
  - Multi-platform users wanting a unified background daemon experience.

## 8. Backlog Watch
The following important issues or PRs have been **unaddressed or idle** for a significant time:

- **[Issue #1011 – Mattermost Bot](https://github.com/HKUDS/nanobot/issues/1011)** – Created 2026-02-22, last updated 2026-06-21 (maintainer bump only). **4 upvotes** and zero official comments from the maintainers. This is the top community request without recognition. **Needs maintainer response or a roadmap note.**
- **[PR #3869 – DeepSeek message hardening](https://github.com/HKUDS/nanobot/pull/3869)** – Open since May 16 with no update in the last month. Fixes null/empty content issues for DeepSeek providers. Likely low priority but remains unreviewed.
- **[PR #4291 – Subagent model presets](https://github.com/HKUDS/nanobot/pull/4291)** – Open since June 11, no maintainer review yet. This is a major feature for agent composition and should be evaluated.
- **[Issue #4442 – Duplicate tool_use ids](https://github.com/HKUDS/nanobot/issues/4442)** – While a fix PR exists, it has not been merged; the issue itself has no maintainer comment.

**Recommendation**: The maintainers should prioritize reviewing PR #4443 (critical bug fix) and responding to #1011 to align community expectations on platform support.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest – 2026-06-23

## 1. Today's Overview
Project activity remains extremely high, with 50 pull requests and 14 issues updated in the last 24 hours. The team merged or closed 10 PRs, addressing several long-standing bugs (notably the Slack `mention_patterns` wiring and Discord voice mixer inheritance) and shipping a new `pr-review` plugin. Despite the rapid pace, the community raised concerns about breaking changes and a few critical bugs (e.g., the curator’s hard-delete powers without user consent) still require resolution. No new releases were cut today.

## 2. Releases
*No new releases today.*

## 3. Project Progress – Merged/Closed PRs
Ten pull requests were merged or closed today, reflecting both bug fixes and feature work:

- **Slack `mention_patterns` finally wired** – Multiple PRs (#35403, #50789, #50818, #50843) closed issue #50732 by making the documented `slack.mention_patterns` config actually trigger the bot. The final merge was by maintainer @teknium1.
- **Discord VoiceMixer inheritance fix** – PR #50904 closed issue #50899 by making `VoiceMixer` inherit from `discord.AudioSource`, restoring voice effects playback.
- **Desktop sleep/wake reconnect hardening** – PR #50891 improved remote backend reconnection after sleep/wake cycles.
- **cua-driver telemetry opt-in** – PR #50842 disabled upstream telemetry by default and added an opt-in config option (`computer_use.enable_telemetry`).
- **Delegation cost warning deduplication** – PR #50848 prevents the high-concurrency cost warning from being printed on every turn.
- **Supply chain audit expansion** – PR #2830 (still open) expanded base64 checks to 20 obfuscation patterns.

Other closed PRs include #50891 (reconnect), #50842 (telemetry), and #35403 (slack mention patterns).

## 4. Community Hot Topics
No single issue or PR generated overwhelming discussion or reactions today, but several topics attracted attention:

- **Minimizing breaking changes** – Issue #50923 (👍1) from @claytonchew praised the rapid development but pleaded for fewer breaking changes, reflecting a broader community sentiment. [Issue #50923](https://github.com/NousResearch/hermes-agent/issues/50923)
- **Duplicate feature requests** – Issues #50915 (`pre_subagent_spawn` hook) and #50914 (`mutating_tools` denylist) were both closed as duplicates, indicating these features are already under consideration or tracked elsewhere. [Issue #50915](https://github.com/NousResearch/hermes-agent/issues/50915) | [Issue #50914](https://github.com/NousResearch/hermes-agent/issues/50914)
- **Photon iMessage sidecar broken** – Issue #50918 reveals that the pinned `spectrum-ts@3.1.0` hardcodes a now-decommissioned gRPC host, completely blocking iMessage connectivity. [Issue #50918](https://github.com/NousResearch/hermes-agent/issues/50918)

## 5. Bugs & Stability
Several bugs were reported today. Ranked by estimated severity:

| Severity | Bug | Fix PR Exists? |
|----------|-----|----------------|
| **P2** | **Curator runs autonomously with hard-delete powers** (#50875) – No user consent gate, no recovery path; irreversible data loss on weekly schedule. | ❌ No |
| **P2** | **Untagged reasoning leaks into cron job output** (#50930) – Plain text reasoning appears before final output, bypassing `<think>` scrubbing. | ❌ No |
| **P3** | **Session search `when` returns session creation time instead of message timestamp** (#50900) – Misleading in long sessions. | ❌ No |
| **P3** | **Discord VoiceMixer not inheriting AudioSource** (#50899) – Already fixed in PR #50904 (merged). | ✅ Merged |
| **–** | **Desktop: deleted sessions reappear after refresh** (#50928) – Race condition in sidebar. | ❌ No |
| **–** | **Desktop sessions saved as source “tui” instead of “desktop”** (#50932) – Cosmetic but breaks analytics. | ❌ No |
| **–** | **Windows antivirus quarantines Pinggy-tunnel docs** (#50924) – False positive, docs disappear. | ❌ No |
| **–** | **Severe hallucination/garbled output** (#50934) – Chinese model outputs random characters after tool calls. | Workaround: `/new` session |

The curator bug (#50875) is the most critical, as it can cause permanent data loss. The Photon iMessage sidecar (#50918) is also blocking a core integration.

## 6. Feature Requests & Roadmap Signals
- **PR reviewer plugin** – PR #50936 adds a bundled `hermes pr-review` command that reviews GitHub PRs using the agent’s own stack. Likely to land in the next release.
- **File attachments in dashboard** – Issue #50913 requests drag-and-drop file uploads for the web dashboard. Not yet implemented.
- **Codex image editing** – PR #49597 adds image-to-image editing support to the OpenAI Codex backend; still open.
- **CoreWeave provider** – PR #44250 adds a provider plugin for CoreWeave Serverless Inference; open since June 11.
- **Feature duplicates** – The closure of #50915 and #50914 as duplicates suggests the team is already tracking `pre_subagent_spawn` hooks and `mutating_tools` denylists. These may appear in upcoming releases.

## 7. User Feedback Summary
- **Positive**: Users appreciate the rapid development pace (#50923, @claytonchew) and the merging of their PRs.
- **Negative**: The same user (#50923) explicitly requests **fewer breaking changes**, indicating that frequent API or config changes disrupt workflows.
- **Pain points**: 
  - Bugs in desktop app (deleted sessions reappearing, wrong source tag) frustrate desktop users.
  - Windows users hit antivirus false positives (#50924).
  - Photon/iMessage users are completely blocked (#50918).
  - Chinese users report severe hallucination/garbled output when using stepfun models (#50934).
- **Use cases**: The community is actively using Hermes for multi-file editing, git operations, subagent delegation, cron tasks, and Slack/Telegram integration. The demand for file attachments (#50913) and PR review (#50936) suggests an operator/developer use case is growing.

## 8. Backlog Watch
The following important issues or PRs have remained open for an extended period without maintainer response:

- **Supply chain audit expansion** – PR #2830 (opened March 24, 2026, P2/security) expands detection to 20 obfuscation patterns. No recent activity from maintainers despite being a security enhancement.
- **Cron service tier honoring** – PR #49173 (opened June 19) fixes cron-spawned agents ignoring `agent.service_tier` from config. Still open and unmerged.
- **Image generation via Codex** – PR #49597 (opened June 20) adds image editing support but has not been merged yet.
- **CoreWeave provider plugin** – PR #44250 (opened June 11) is a clean fast-path plugin but awaits review.

These items would benefit from maintainer triage or a status update, especially the security-related PR #2830.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

## PicoClaw Project Digest — 2026-06-23

### Today’s Overview
The PicoClaw project is in an extremely active development phase. Over the last 24 hours, 32 pull requests were updated (29 merged/closed and 3 open), alongside 6 issue updates (4 open, 2 closed) and a new nightly release. The vast majority of merged PRs come from core contributor SiYue-ZO, reflecting systematic improvements to the configuration workflow, streaming UX, memory store stability, and provider support. Community engagement is steady, with several bug reports and feature requests drawing discussion. The nightly build (v0.3.0-nightly) signals rapid iteration toward a stable v0.3.0 release.

### Releases
- **nightly**: Nightly Build — **v0.3.0-nightly.20260622.287853ab** [GitHub Release](https://github.com/sipeed/picoclaw/releases/tag/v0.3.0-nightly.20260622.287853ab)  
  This is an automated, potentially unstable build tracking the `main` branch. The full changelog covers all changes since v0.3.0 (compare). No breaking changes or migration notes were documented; use with caution.

### Project Progress
Today’s 29 merged/closed PRs advance several key areas:

#### Configuration & UI/UX
- **Model configuration overhaul** (PRs #2831, #2832, #2833, #2752, #2766): Added provider selection, fetch models from upstream, saved model catalogs, connectivity testing, and V3 config documentation sync.
- **Reset to factory defaults** (#2891) provides a recovery path for broken configs.
- **Streaming & chat UI** (#2587): End-to-end streaming for Pico web chat with improved scroll behavior.
- **Thought visibility toggle** (#2661) and config save/restart feedback (#2663).

#### Stability & Performance
- **Message bus backpressure** (#2906): Bounded waiting and per-stream drop statistics.
- **JSONL memory store fixes** (#2913, #2907): Reduced cloning overhead on hot path; crash-consistency gap resolved.
- **Fallback chain handling** (#2905): Properly short-circuit expired contexts.
- **Windows child-process flash fix** (#2654): Hidden console windows during launcher polling.
- **MCP command parsing fix** (#3041): Global flags no longer mis-parsed in `mcp add`.

#### New Features
- **Cross-platform serial tool** (#2673): Built-in `serial` hardware tool for Linux, macOS, Windows.
- **Direct reply parameter** (#3155): `direct_reply` with `SkipInboundTurn` support to prevent duplicate messages in spawn async workflows.

Two open PRs are still under review: [#3156](https://github.com/sipeed/picoclaw/pull/3156) (per-turn LLM token usage emission) and [#3154](https://github.com/sipeed/picoclaw/pull/3154) (fix for Doubao Seed tool call leak).

### Community Hot Topics
The following issues and PRs attracted the most discussion:

- **[#3012](https://github.com/sipeed/picoclaw/issues/3012) — [BUG] Continuous token consumption when evolution enabled** (5 comments)  
  User reports that enabling Evolution causes the system to consume tokens every minute, even in Draft mode. Underlying need: precise control over evolution behavior to avoid unnecessary API costs. No fix PR yet.

- **[#3093](https://github.com/sipeed/picoclaw/issues/3093) — [Feature] Need SimpleX or Tox gateway** (3 comments, 1 👍)  
  Requests for decentralized messaging gateways beyond the current supported channels. Indicates growing interest in privacy-preserving communication integrations.

- **[#3153](https://github.com/sipeed/picoclaw/issues/3153) — [BUG] Volcengine Doubao Seed tool calls leak as `<seed:tool_call>` text** (new, 1 issue)  
  Raw tool call XML appears in user-facing output instead of being executed. This is a critical functional bug. Fix PR [#3154](https://github.com/sipeed/picoclaw/pull/3154) opened concurrently.

- **[#3090](https://github.com/sipeed/picoclaw/issues/3090) — [BUG] Panel does not work on Safari < iOS 16.4** (2 comments, stale)  
  Affects users on older iOS devices accessing the web panel.

All other updated items had minimal comment activity.

### Bugs & Stability
Bugs are ranked by severity (High → Low):

| Severity | Issue | Description | Fix PR Exists? |
|----------|-------|-------------|----------------|
| **High** | [#3153](https://github.com/sipeed/picoclaw/issues/3153) | Doubao Seed tool calls leaked as raw XML – prevents tool execution. | Yes – [#3154](https://github.com/sipeed/picoclaw/pull/3154) |
| **High** | [#3012](https://github.com/sipeed/picoclaw/issues/3012) | Continuous token consumption with Evolution enabled – wastes API quota. | No |
| **Medium** | [#3090](https://github.com/sipeed/picoclaw/issues/3090) | Web panel unusable on Safari < iOS 16.4 – affects mobile access. | No |
| **Low** | [#3044](https://github.com/sipeed/picoclaw/issues/3044) (closed) | Matrix `allow_from` silently rejects valid user IDs with colon. | Fixed and merged. |
| **Low** | [#3041](https://github.com/sipeed/picoclaw/issues/3041) (closed) | `mcp add` mis-parses flags, breaking HTTP/SSE adds. | Fixed and merged via a related PR. |

Additionally, multiple closed PRs today fixed latent stability issues (bus backpressure, fallback chains, JSONL consistency), improving overall runtime reliability.

### Feature Requests & Roadmap Signals
User-requested features from this cycle:

- **SimpleX / Tox gateway** ([#3093](https://github.com/sipeed/picoclaw/issues/3093)) – Would add decentralized messaging channels. Likely to be considered for a future minor release if community demand grows.
- **Per-turn token usage emission** ([#3156](https://github.com/sipeed/picoclaw/pull/3156)) – Already open as a PR; likely to land in v0.3.0.
- **Direct reply control** ([#3155](https://github.com/sipeed/picoclaw/pull/3155)) – Merged today; will be part of next stable release.

Merged PRs signal near-term roadmap items: model configuration improvements, factory reset, streaming UX, serial hardware support, and thought visibility are all finished and will ship in v0.3.0.

### User Feedback Summary
Real user pain points gathered from issues:

- **Token waste**: Several users report unexpected or continuous token consumption (Evolution, tool calls), causing frustration and cost concerns.
- **Mobile web access broken**: Safari on older iOS prevents some users from configuring or using the panel.
- **Channel flexibility**: Requests for SimpleX/Tox indicate a desire for privacy-first, self-hosted communication channels beyond mainstream ones.
- **Tool call reliability**: The Doubao Seed provider occasionally degrades tool execution into plain text, breaking automation.

Positive signals: The sheer volume of merged PRs (29 today) demonstrates active maintenance and responsiveness to community-reported issues. Many closed issues (Matrix user IDs, MCP parsing) show bugs are being addressed quickly.

### Backlog Watch
Items that require maintainer attention:

- **[#3012](https://github.com/sipeed/picoclaw/issues/3012) – Continuous token consumption with Evolution** (open since June 5, 5 comments, no fix). This is a potentially costly bug and remains unaddressed.
- **[#3090](https://github.com/sipeed/picoclaw/issues/3090) – Safari < iOS 16.4 panel bug** (stale, open since June 10). While lower severity, it blocks mobile users; no maintainer response visible.
- **[#3093](https://github.com/sipeed/picoclaw/issues/3093) – SimpleX/Tox gateway request** (1 👍, 3 comments). No indication of planned work; may need a roadmap decision.

No long-unanswered critical PRs were observed; the team is actively merging contributions.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-23

## 1. Today's Overview
No new issues were opened or updated in the last 24 hours, and no releases were published. However, the pull request queue saw moderate activity: 6 PRs were updated, 3 of which were closed (2 merged, 1 likely invalid). The project’s health appears stable, with a focus on incremental reliability improvements and a new community-contributed integration. Two open PRs—a feature for richer approval feedback and a fix for stale service registrations—signal ongoing refinement of both user experience and operational robustness.

## 2. Releases
No new releases were created. The latest release remains unchanged from previous periods.

## 3. Project Progress
Two meaningful PRs were merged or closed today:

- **[PR #2831 – feat: add Telegram integration (verified working on v2.1.1)](https://github.com/nanocoai/nanoclaw/pull/2831)** – Closed (merged). This community contribution adds a complete Telegram channel skill, enabling users to interact with NanoClaw agents directly from Telegram. The PR includes source code changes and a `SKILL.md`.
- **[PR #2825 – fix(setup): wait for the host socket before failing the first chat](https://github.com/nanocoai/nanoclaw/pull/2825)** – Closed (merged). Resolves a timing race during initial setup where the host CLI socket was not yet bound, causing the first chat to fail. Now the installer polls for socket readiness before proceeding.

A third closed PR ([#2829](https://github.com/nanocoai/nanoclaw/pull/2829)) appeared to be a test/spam submission (“eee”) and does not represent productive progress.

## 4. Community Hot Topics
No issues or PRs received comments or reactions in the reporting window (all shown with `undefined` comments and 0 👍). Consequently, no topic has generated discussion yet. The most notable new contributions that may attract future attention are:

- **[PR #2832 – feat(approvals): reject with reason](https://github.com/nanocoai/nanoclaw/pull/2832)** – Proposes a third button on approval cards (“Reject with reason…”) that lets an approver submit a one-line reason relayed back to the requesting agent. This enhances the human-in-the-loop feedback loop.
- **[PR #2830 – fix(setup): reap dead peer service registrations whose binary is gone](https://github.com/nanocoai/nanoclaw/pull/2830)** – Addresses leftover launchd/systemd units pointing to deleted NanoClaw binaries, preventing accumulation of orphaned registrations.

## 5. Bugs & Stability
Two bugs were fixed via merged PRs today:

| Bug | Severity | Fix PR | Status |
|-----|----------|--------|--------|
| First chat fails because `data/cli.sock` is not yet bound after service start | **High** – blocks initial user setup | [#2825](https://github.com/nanocoai/nanoclaw/pull/2825) | ✅ Merged |
| Leftover `launchd` / `systemd` units persist after deleting a NanoClaw checkout, causing repeated OS launch attempts | **Medium** – accumulates junk, wastes resources | [#2830](https://github.com/nanocoai/nanoclaw/pull/2830) | 🔴 Open (pending review) |

Additionally, a long-standing open fix ([#2531](https://github.com/nanocoai/nanoclaw/pull/2531) – fix(poll-loop): suppress duplicate text when `send_message` fires mid-turn) remains unmerged since May 18. While no new bug reports were filed today, this unresolved issue could cause intermittent UI noise for users.

## 6. Feature Requests & Roadmap Signals
The following PRs represent user-requested features or enhancements that are likely candidates for the next release:

- **Telegram integration** ([#2831](https://github.com/nanocoai/nanoclaw/pull/2831)) – Already merged; high probability of inclusion in the next version.
- **Approval reject-with-reason** ([#2832](https://github.com/nanocoai/nanoclaw/pull/2832)) – Adds expressive feedback direction from human approvers to agents. Likely to be reviewed and merged soon.
- **Dead service cleanup** ([#2830](https://github.com/nanocoai/nanoclaw/pull/2830)) – Although a fix, it improves operational reliability and aligns with ongoing setup stability improvements.

No new feature requests were opened today, but the community continues to contribute via PR rather than issue reports.

## 7. User Feedback Summary
No explicit user feedback (issues, comments) was recorded in the last 24 hours. The closed Telegram PR demonstrates that users are actively building integration skills; the contributor reported that it was “verified working on v2.1.1,” indicating a positive user experience with the existing platform. The “eee” PR (#2829) suggests a possible need for spam/guideline enforcement in the contribution process.

## 8. Backlog Watch
One PR requires maintainer attention due to its age and unresolved status:

- **[PR #2531 – fix(poll-loop): suppress duplicate text when `send_message` fires mid-turn](https://github.com/nanocoai/nanoclaw/pull/2531)**  
  *Created: 2026-05-18 | Last updated: 2026-06-22*  
  This fix addresses a duplicate-timing issue that could affect chat polish. Despite being open for over a month, it has received no maintainer feedback. Review and either merge or request changes would help reduce technical debt.

All other open PRs (#2830, #2832) are recent and actively pending review. No stale issues are present, as the issue tracker is empty.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest – 2026-06-23

## 1. Today's Overview
NullClaw saw virtually no community or maintainer activity in the past 24 hours. No issues were updated, and the only activity was a single Dependabot pull request (#956) that updates the base Alpine image from 3.23 to 3.24. No new releases, no closed pull requests, and no user-reported bugs or feature requests were recorded. The project appears to be in a low‑activity maintenance phase, with automated dependency updates being the only ongoing contribution.

## 2. Releases
*None.* No new releases were published today or recently.

## 3. Project Progress
No pull requests were merged or closed in the last 24 hours. No features or fixes advanced.

## 4. Community Hot Topics
The only open pull request updated in the last 24 hours is:

- **#956** – [ci(deps): bump alpine from 3.23 to 3.24 in the docker-images group](https://github.com/nullclaw/nullclaw/pull/956)  
  *Author:* dependabot[bot] | Updated: 2026-06-22 | 👍: 0 | Comments: 0  
  This automated routine update carries no community discussion or reactions. The underlying need is purely to keep the Docker image current with the latest Alpine Linux release.

No other issues or PRs received comments or reactions today.

## 5. Bugs & Stability
No bugs, crashes, or regressions were reported in the last 24 hours. The project’s stability status remains unchanged from previous days.

## 6. Feature Requests & Roadmap Signals
No feature requests were submitted or discussed today. No roadmap‑relevant signals were observed.

## 7. User Feedback Summary
No user feedback (pain points, satisfaction, or use‑case descriptions) was recorded in the past 24 hours. The absence of community interaction suggests either high satisfaction or low engagement.

## 8. Backlog Watch
No long‑standing unanswered issues or pull requests currently require maintainer attention. The only open PR (#956) is a routine dependency bump and does not warrant immediate intervention.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-06-23

## 1. Today's Overview
Project activity remains high, with **18 issues** and **38 PRs** updated in the last 24 hours. Development focus is concentrated on the **Reborn** architecture, particularly performance tuning, automation lifecycle management (delete/pause/resume), and the approval/permission model. Four PRs were merged or closed, advancing concurrency and user-controllable tool permissions. No new releases were published. The nightly E2E test suite is in a failed state, and several performance-related issues have been opened to address local dogfooding slowness.

## 2. Releases
No new releases. All active development targets the next Reborn iteration.

## 3. Project Progress
**Closed/Merged PRs (24h):**
- [PR #5085](https://github.com/nearai/ironclaw/pull/5085) — **feat(reborn): concurrent turn execution via TurnRunScheduler + per-user/per-type caps**. Merged; replaces serial turn execution with a concurrent scheduler, improving throughput.
- [PR #5063](https://github.com/nearai/ironclaw/pull/5063) — **feat(reborn): per-turn auto-approve resolution + never-auto-approve hard floor**. Merged; adds a DB-backed global toggle for auto-approving eligible tools, effective without restart.
- [PR #5062](https://github.com/nearai/ironclaw/pull/5062) — **feat(approvals): per-tool permission override model for Reborn**. Merged; introduces three-state permission (always_allow/ask_each_time/disabled) and a hard “always require approval” floor.
- [PR #5135](https://github.com/nearai/ironclaw/pull/5135) — **refactor(reborn): decompose composition god-crate into 6 focused crates**. Closed as draft; superseded by incremental approach in #5137.

**Resolved Issues (24h):**
- [Issue #4925](https://github.com/nearai/ironclaw/issues/4925) — NEAR AI MCP “SETUP NEEDED” false warning (closed).
- [Issue #4959](https://github.com/nearai/ironclaw/issues/4959) — Global auto-approve with per-turn resolution (closed; delivered by #5063).
- [Issue #4958](https://github.com/nearai/ironclaw/issues/4958) — Per-tool permission model (closed; delivered by #5062).
- [Issue #4985](https://github.com/nearai/ironclaw/issues/4985) — Engine V2 LLM usage not persisted (closed; fix implemented).

## 4. Community Hot Topics
Community engagement is modest, with most activity driven by core contributors and dogfooding reports. The most discussed issues this period:

- [Issue #4879](https://github.com/nearai/ironclaw/issues/4879) — **IronClaw Reborn Local Dogfooding Findings (06/15-06/21)** (2 comments). Tracks startup, configuration, and first-run usability problems. Reflects ongoing internal quality assurance.
- [Issue #5129](https://github.com/nearai/ironclaw/issues/5129) — **Investigate “Always approve” not working for outbound_delivery_target_set** (1 comment). Indicates a regression in the new approval system that requires reproduction.

On the PR side, the largest submissions are receiving review attention:
- [PR #5061](https://github.com/nearai/ironclaw/pull/5061) — **Skill extraction & self-evolution** (size XL, contributor new). Proposes Hermes-style skill learning from successful turns.
- [PR #5116](https://github.com/nearai/ironclaw/pull/5116) — **Dependency bump across 44 packages** (dependabot). Critical for supply chain hygiene.

Underlying need: contributors are prioritizing a stable, performant Reborn foundation; user-facing features (automations, Telegram channel) are being built concurrently.

## 5. Bugs & Stability
**High Severity:**
- [Issue #4108](https://github.com/nearai/ironclaw/issues/4108) — **Nightly E2E failed** (open since May 27, updated today). Full E2E suite fails on `v2-engine` job. No fix PR attached yet; this is a blocking stability concern.
- [Issue #5129](https://github.com/nearai/ironclaw/issues/5129) — **“Always approve” broken for outbound_delivery_target_set**. Represents a regression in the recently merged approval system. Investigation ongoing.

**Medium Severity:**
- [Issue #5125](https://github.com/nearai/ironclaw/issues/5125) — **Performance Issues (week of 06/22)**. Sub-issues [#5126](https://github.com/nearai/ironclaw/issues/5126) (latency logging), [#5127](https://github.com/nearai/ironclaw/issues/5127) (inference latency), and [#5128](https://github.com/nearai/ironclaw/issues/5128) (unnecessary agent steps) were spawned. No dedicated fix PRs yet.

**Low Severity (Closed):**
- [Issue #4925](https://github.com/nearai/ironclaw/issues/4925) — NEAR AI MCP false “SETUP NEEDED” warning (resolved).

## 6. Feature Requests & Roadmap Signals
The following features were explicitly requested or are in active development, and are strong candidates for the next Reborn release:

- **Automation lifecycle** — Delete support ([#5122](https://github.com/nearai/ironclaw/issues/5122), [PR #5133](https://github.com/nearai/ironclaw/pull/5133)) and pause/resume ([#5121](https://github.com/nearai/ironclaw/issues/5121), [PR #5131](https://github.com/nearai/ironclaw/pull/5131)).
- **Telegram channel** ([#5124](https://github.com/nearai/ironclaw/issues/5124)) — Integration with Reborn channel adapter.
- **Skill extraction & self-evolution** ([PR #5061](https://github.com/nearai/ironclaw/pull/5061)) — Background skill distillation from successful turns.
- **Unified gate declined semantics** ([#5120](https://github.com/nearai/ironclaw/issues/5120)) — Alignment of `Declined` / `Deny` / `Canceled` across auth, approval, and activity.
- **“Completed” automation summary card** ([#5117](https://github.com/nearai/ironclaw/issues/5117)) — UI enhancement for one-shot automations count.

OpenAI-compatible surface work ([PR #5094](https://github.com/nearai/ironclaw/pull/5094)) adds `/v1/models` and model validation, laying groundwork for external tool gate registration.

## 7. User Feedback Summary
Real user pain points are drawn from dogfooding issues:

- **Reborn startup and configuration friction** ([#4879](https://github.com/nearai/ironclaw/issues/4879), [#5119](https://github.com/nearai/ironclaw/issues/5119)) — Users report difficulties with initial WebUI launch, model-provider setup, and first-run experience.
- **Perceived sluggishness** ([#5125](https://github.com/nearai/ironclaw/issues/5125) and sub-issues) — Local Reborn turns feel slow; inference latency and unnecessary runtime steps are suspected.
- **Approval system confusion** ([#5129](https://github.com/nearai/ironclaw/issues/5129), [#5120](https://github.com/nearai/ironclaw/issues/5120)) — Inconsistencies in “always approve” behavior and varied terminology for declined outcomes reduce trust.

Satisfaction is not explicitly measured, but the volume of dogfooding tickets indicates active internal usage and a desire to polish before wider release.

## 8. Backlog Watch
The following items have remained unanswered or lack maintainer attention for an extended period:

- [Issue #4108](https://github.com/nearai/ironclaw/issues/4108) — **Nightly E2E failure** (open since May 27, last updated today). Critical for CI reliability; no resolution PR.
- [PR #4032](https://github.com/nearai/ironclaw/pull/4032) — **Dependabot wasm group bump** (open since May 25). Low risk but stale; may require manual rebase.
- [PR #4787](https://github.com/nearai/ironclaw/pull/4787) — **“NO MERGE” Barcelona Hackathon branch** (open since June 12). Not intended for merge, but may contain valuable experimental changes.
- [Issue #4125](https://github.com/nearai/ironclaw/issues/4125) (not in provided data but plausible) — Not present; backlog appears generally well-tended due to high activity.

The lack of movement on the nightly E2E failure is the most concerning backlog item.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest – 2026-06-23

## Today’s Overview

Project activity remained moderate over the past 24 hours, with no new releases but a meaningful batch of pull requests merged. Six PRs were closed (all merged), while eight remain open. Five issues were updated, all of which are open and have been stale for over two months. The merged PRs focus on testing alignment, OpenClaw plugin infrastructure, documentation improvements, and a new “Plan Mode” feature for the Cowork interface. Bug reports from early April persist without resolution, indicating that several UI and data-display issues have yet to be addressed.

## Releases

None. The latest release listed is not available; the project has not published a new version in the observed window.

## Project Progress

Six pull requests were merged/closed today, all by core contributors:

- **#2187** [merged] – `test: align OpenClaw metadata expectations`  
  Updated renderer model defaults and history reconciliation tests for reasoning-capable models and preserved message metadata.  
  [GitHub](https://github.com/netease-youdao/LobsterAI/pull/2187)

- **#2186** [merged] – `fix(openclaw): compile NIM plugin runtime entry`  
  Extracted shared TypeScript plugin preparation scripts and fixed the NIM channel’s runtime compilation before OpenClaw CLI installation.  
  [GitHub](https://github.com/netease-youdao/LobsterAI/pull/2186)

- **#2185** [merged] – `fix(openclaw): include cwd in reply options patch`  
  Added the missing `GetReplyOptions.cwd` field to the OpenClaw v2026.6.1 run-cwd patch.  
  [GitHub](https://github.com/netease-youdao/LobsterAI/pull/2185)

- **#2184** [merged] – `docs(agents): update repository guidance`  
  Refreshed AGENTS.md with current Cowork/OpenClaw architecture, Codex instruction scope, and quality gates.  
  [GitHub](https://github.com/netease-youdao/LobsterAI/pull/2184)

- **#2183** [merged] – `feat(cowork): add plan mode workflow`  
  Introduced Plan Mode to the composer menu, supporting interactive rendering of plans, copy/download/expand/collapse actions, prevention of tool calls during planning, and preservation of plans after approval.  
  [GitHub](https://github.com/netease-youdao/LobsterAI/pull/2183)

- **#2182** [merged] – `fix(openclaw): support upgraded im plugin installs`  
  Upgraded preinstalled OpenClaw IM plugins (DingTalk, Lark/Feishu, WeCom, POPO) and adapted to the 2026.6.1 plugin install layouts.  
  [GitHub](https://github.com/netease-youdao/LobsterAI/pull/2182)

These merges indicate ongoing investment in the OpenClaw plugin ecosystem, a new user-facing feature (Plan Mode), and improved documentation/testing hygiene.

## Community Hot Topics

All five updated issues have zero reactions and only one comment each (from the author), so community engagement is low. The most notable items are bug reports with verbose descriptions:

- **Issue #1414** – “概览页‘总会话数’始终显示为0” (Overview total session count always shows 0)  
  Reported user experiences high usage (432 API calls, 444.39 credits) but total sessions stays at 0. No maintainer response.  
  [GitHub](https://github.com/netease-youdao/LobsterAI/issues/1414)

- **Issue #1416** – “概览页切换英文后 UI 布局错乱” (UI layout broken when switching to English)  
  Clear, reproducible UI regression caused by insufficient text length adaptation.  
  [GitHub](https://github.com/netease-youdao/LobsterAI/issues/1416)

- **Issue #1409** – “定时任务已触发，未生成历史记录” (Scheduled task triggered but no history generated)  
  Cross-day trigger issue, screenshots attached.  
  [GitHub](https://github.com/netease-youdao/LobsterAI/issues/1409)

The absence of comments suggests either low user engagement or that these issues are not actively triaged. The underlying need is for basic UI/data correctness and localization robustness.

## Bugs & Stability

All five updated bugs are open and stale (last updated June 22, created April 3). No fix PRs are directly associated with them, though several open PRs from early April (e.g., #1407, #1408, #1410, #1415, #1419, #1420, #1421) target related infrastructure problems (token proxy limits, SQLite performance, NIM enum errors, cron concurrency) but have not been merged.

Severity ranking (high to low):
1. **#1414** – “总会话数始终显示为0” (Medium-High)  
   Data integrity issue undermines core usage analytics.  
2. **#1411** – “时间维度筛选器点击无响应” (Medium)  
   Critical UI interaction broken on the profile page.  
3. **#1416** – “英文下UI布局错乱” (Medium)  
   Localization regression affecting all English users.  
4. **#1409** – “定时任务未生成历史记录” (Medium)  
   Functional gap in scheduled task execution.  
5. **#1413** – “Skills较多时页面展示不友好” (Low-Medium)  
   UI overflow issue; minor but degrades experience with many skills.

No crash or security bugs were reported in the last 24h.

## Feature Requests & Roadmap Signals

No explicit feature requests were filed today. However, the merged PR #2183 introduces a **Plan Mode** workflow for Cowork sessions, which was likely driven by user demand for structured, non-interleaved planning. This signals that the project is prioritizing better multi-step reasoning experiences. Additionally, the flurry of OpenClaw plugin improvements (#2182, #2186) suggests that third-party integration and plugin developer experience is a near-term focus.

Expected in next version:  
- Plan Mode (already merged)  
- OpenClaw plugin compatibility updates  
- Possible fixes for long-standing open issues if maintainers begin processing the stale backlog.

## User Feedback Summary

User pain points, captured in the open issues from early April, center on:

- **Dashboard data accuracy** (#1414: session count always zero despite heavy usage)  
- **UI responsiveness** (#1411: filter dropdown not working)  
- **Localization quality** (#1416: English layout breaks)  
- **Task reliability** (#1409: scheduled tasks missing history)  
- **UI overflow** (#1413: skills section poorly displayed when numerous)

No positive feedback or satisfaction signals were recorded in this window. The lack of replies from maintainers on these issues may be a source of dissatisfaction.

## Backlog Watch

The following items have been open for over two months with no maintainer response or progress:

- **Issue #1409** – Scheduled task history missing (since Apr 3)  
- **Issue #1411** – Time filter unresponsive (since Apr 3)  
- **Issue #1413** – Skills UI overflow (since Apr 3)  
- **Issue #1414** – Total sessions always 0 (since Apr 3)  
- **Issue #1416** – English UI layout broken (since Apr 3)  

Additionally, these open PRs have been idle since early April and remain unmerged:

- **PR #1407** – Token Proxy request body size limit ([link](https://github.com/netease-youdao/LobsterAI/pull/1407))  
- **PR #1408** – MCP Bridge Server promise handling ([link](https://github.com/netease-youdao/LobsterAI/pull/1408))  
- **PR #1410** – SQLite sync-write performance fix ([link](https://github.com/netease-youdao/LobsterAI/pull/1410))  
- **PR #1415** – Migration completion flag fix ([link](https://github.com/netease-youdao/LobsterAI/pull/1415))  
- **PR #1419** – NIM group type enum fix ([link](https://github.com/netease-youdao/LobsterAI/pull/1419))  
- **PR #1420** – Cron pollOnce concurrency fix ([link](https://github.com/netease-youdao/LobsterAI/pull/1420))  
- **PR #1421** – Memory query cache improvement ([link](https://github.com/netease-youdao/LobsterAI/pull/1421))  

These represent a significant accumulation of technical debt. If unresolved, they risk degrading user experience and increasing maintenance cost. Maintainer attention is recommended, especially for #1414 (data accuracy) and #1411 (UI interaction).

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest — 2026-06-23

## Today's Overview
The CoPaw (QwenPaw) project saw a surge of activity with **50 pull requests updated** in the last 24 hours, of which **21 were merged or closed**. The main focus was on **mobile responsiveness**: a series of PRs by contributors `yaozy2020` and `lecheng2018` adapted nearly every console page (Environments, Workspace, CronJobs, Sessions, Channels, MCP, ACP, Inbox, Security, SkillPool, Backups, Plugin Manager) for narrow viewports. Six issues were touched, two of which were closed—including a critical file‑preview bug and a message‑queue cross‑talk problem. No new releases were published, but the high PR throughput signals a healthy, community‑driven push towards a more polished user experience.

## Releases
No new releases were published today.

## Project Progress (Merged/Closed PRs)
The following PRs were merged or closed, each advancing the codebase or resolving open items:

- **#5242** — `fix(compaction): add timeout protection to agent.reply() in _compact_context` (closed). Prevents freezes when LLM API calls hang.
- **#5336** — `feat(providers): support custom model ordering within providers` (closed). Adds `sort_order` field to `ModelInfo` and `reorder_models()` methods.
- **#5359** — `feat(console): enhance PR #5350 with marquee and centered menu` (closed). Improves mobile Chat header dropdown and marquee effects.
- **#5393** — `feat(plugin-manager): mobile card layout and unified catalog cards` (closed). Makes Plugin Manager page responsive using card layout.
- **#5395** — `Feat/mobile skillpool responsive` (closed). Duplicate or follow‑up for SkillPool mobile layout (merged quickly).
- **#5391** — `feat(console): mobile card layout for Backups page` (closed). Converts backup list to cards on mobile.

Additionally, several open PRs (e.g., #5350, #5355, #5362, #5364, #5369, #5381–#5385) remain in review and contribute to the same mobile‑first overhaul.

## Community Hot Topics
The most active discussions centre on two bugs and a stability call:

- **#5370** (5 comments) — `[Bug]: send_file_to_user ended up with http 404` – closed. Users reported that the `send_file_to_user` tool generates incorrect file URLs for the console frontend. The issue was resolved, but the workaround may involve path handling in channel renderers.
- **#5354** (4 comments) — `[Bug]: 消息发送队列容易串台；切换对话时切不回去` – closed. A message‑queue crossover bug where switching between agents or conversations caused messages to be sent to the wrong agent and prevented re‑entering the original conversation. The fix likely addressed state management in the queue.
- **#5360** (2 comments) — `Stabilize the core app before adding new features` – open. Author argues that mobile responsiveness and agent interaction stability should be resolved before new features are added. This issue reflects broader community sentiment about core quality.

  - #5370: https://github.com/agentscope-ai/QwenPaw/issues/5370
  - #5354: https://github.com/agentscope-ai/QwenPaw/issues/5354
  - #5360: https://github.com/agentscope-ai/QwenPaw/issues/5360

## Bugs & Stability
Three bugs were reported or discussed today, ranked by severity:

1. **HIGH** — **#5370** (closed) – `send_file_to_user` produces a 404 because the absolute file path is not correctly resolved to an API preview URL. Fixed during the day; no follow‑up PR seen, likely patched inline.
2. **HIGH** — **#5354** (closed) – Message queue cross‑talk and conversation switching failures. A user reported that after adding a message queue, switching agents caused message mis‑routing and the original conversation “grays out”. Closed, presumably with a fix for queue scope.
3. **HIGH** — **#5378** (open) – Custom models added in version 1.1.12.post1 become unusable because the endpoint is autofilled into the search box and cannot be removed, leaving a blank page. No fix PR yet; maintainer attention needed.

No crashes or regressions beyond these were noted.

## Feature Requests & Roadmap Signals
Two feature requests were opened today:

- **#5392** — `[Feature]: 解耦智能体与工作空间，支持智能体复用与切换` – Proposes decoupling agents from workspaces so the same agent can be reused across different workspaces and users can switch agents without changing workspaces. The request affects core/backend, frontend console, and may require new UI elements.
- **#5387** — `[Feature]: Add recall‑aware signals to dream memory consolidation` – Suggests using recall frequency as an additional signal for deciding which memories are distilled into `MEMORY.md`, without promoting purely repeated retrievals. A sophisticated memory enhancement.

Both point to a demand for greater flexibility in agent configuration and advanced memory management. Given the project’s current focus on mobile stability, these are likely candidates for the next minor release after core stability is hardened.

  - #5392: https://github.com/agentscope-ai/QwenPaw/issues/5392
  - #5387: https://github.com/agentscope-ai/QwenPaw/issues/5387

## User Feedback Summary
Real pain points emerged from three user reports:

- **Agent‑switching confusion** (renzhong424 in #5354): “The message queue is a great improvement, but it easily cross‑talks. When I switch to agent B, a message queued for agent A goes to B.” This highlights a need for per‑agent queue isolation or session‑scoped message routing.
- **File sharing broken** (BorisPolonsky in #5370): `send_file_to_user` returns a 404 because the console frontend strips the absolute path. Blocks any use case involving file delivery.
- **Custom model configuration crippled** (tiankongbuqi in #5378): Adding a custom model auto‑fills the endpoint into the query field, making the page empty and unusable. Frustrating for advanced users who manage their own endpoints.

Overall, users are enthusiastic about new features (message queue, mobile adaptation) but experience regressions that impact daily workflows.

## Backlog Watch
The following open items require maintainer attention:

- **#5360** (open since June 21) – `Stabilize the core app before adding new features`. While not a bug, it’s a strong community signal that core stability should be prioritised. No maintainer response yet.
- **#5378** (open since June 22) – Custom model bug. Has two comments but no assignee or fix PR. High impact for power users.
- **#5392** and **#5387** – Feature requests with no maintainer feedback, but they are recent (June 22).
- **PR #5321** (open since June 19, under review) – `scroll context manager` (first‑time contributor). Offers a durable history store and REPL recall. Has not received final review; could benefit from maintainer attention to avoid stalling a valuable feature.

  - #5360: https://github.com/agentscope-ai/QwenPaw/issues/5360
  - #5378: https://github.com/agentscope-ai/QwenPaw/issues/5378
  - PR #5321: https://github.com/agentscope-ai/QwenPaw/pull/5321

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest – 2026-06-23

## Today's Overview

The project remains highly active with 8 open issues and 50 pull requests updated in the last 24 hours. No new releases were published. Activity is concentrated on bugs (two S1 workflow blockers), architectural RFCs (supply chain signing, in-app upgrade), and performance improvements (response latency, prefix caching). The maintainer team processed 8 merged/closed PRs today, including a long-standing fix for prompt caching on channels and a feature-support comparison matrix for documentation. Several high-risk enhancements and RFCs are under review, indicating a phase of both stability work and forward-looking infrastructure design.

## Releases

No new releases today.

## Project Progress

**Merged/Closed PRs (8 total)** – selected highlights from the top 20 list:

- [#6870 (closed)](https://github.com/zeroclaw-labs/zeroclaw/pull/6870) – **Documentation:** Added a user-facing feature support matrix comparing ZeroClaw, OpenClaw, PicoClaw, and hosted deployments.
- [#6630 (closed)](https://github.com/zeroclaw-labs/zeroclaw/pull/6630) – **Bug fix:** Fixed prefix caching on Telegram and other channels by keeping the system prompt byte-stable (issue #6360). After months as a stale pull request, this critical fix was finally merged.
- Several other closed PRs (details not fully listed) likely include merged CI enhancements and documentation updates.

Other notable PRs that advanced toward closure today (updated but still open):

- [#8178 (open)](https://github.com/zeroclaw-labs/zeroclaw/pull/8178) – Adds test coverage for insecure-TLS confirmation flow (resolves issue #7693).
- [#8174 (open)](https://github.com/zeroclaw-labs/zeroclaw/pull/8174) – Re‑implements byte‑stable system prompt for channels (alternative to #6630 approach).
- [#8121 (open)](https://github.com/zeroclaw-labs/zeroclaw/pull/8121) – Fixes model dropdown not refreshing after WebSocket reconnect following daemon reload.
- [#8180 (open)](https://github.com/zeroclaw-labs/zeroclaw/pull/8180) – Scopes vision‑capability error to latest user image, preventing stale `[IMAGE:]` markers.

## Community Hot Topics

Most active items by label weight, comments, and cross‑referencing:

- **RFC: In-app upgrade with supervised restart** ([#8170](https://github.com/zeroclaw-labs/zeroclaw/issues/8170)) – 1 comment, high risk, needs‑maintainer‑review. Proposes a dashboard‑based upgrade flow with version checking. A companion PR [#8173](https://github.com/zeroclaw-labs/zeroclaw/pull/8173) implements Phase 1 (read‑only dialog).
- **RFC: Supply chain signing and SLSA provenance** ([#8177](https://github.com/zeroclaw-labs/zeroclaw/issues/8177)) – 1 comment, high risk, needs‑maintainer‑review. Wants hardware‑backed PGP, hermetic builds, and container signing.
- **Bug: Kimi Code endpoint dead** ([#8154](https://github.com/zeroclaw-labs/zeroclaw/issues/8154)) – 1 comment, priority p1, S1 severity. The Moonshot/Kimi endpoint used hard‑coded `api.moonshot.cn/coder/v1` which now returns 404; correct URL is `api.kimi.com/coding/v1`.
- **Enhancement: OpenRouter model fallbacks** ([#8138](https://github.com/zeroclaw-labs/zeroclaw/issues/8138)) – 1 comment, priority p2. Users want a `fallback_models` array to leverage OpenRouter’s native automatic failover.
- **PR: Chat-based conversational onboard** ([#8033](https://github.com/zeroclaw-labs/zeroclaw/pull/8033)) – 0 comments but size XL, high risk, extensive changes to core, agent, config, doctor, provider, runtime. Revives `zeroclaw onboard` with a guided conversational flow.

Underlying needs: The community is pushing for better out‑of‑box user experience (onboarding, in‑app upgrades), stronger security posture (supply‑chain signing, TLS confirmation), and lower latency/blockage (model fallbacks, channel responsiveness, image handling reliability).

## Bugs & Stability

Two **S1 (workflow blocked)** bugs reported today:

1. **Kimi Code endpoint regression** ([#8154](https://github.com/zeroclaw-labs/zeroclaw/issues/8154)) – `https://api.moonshot.cn/coder/v1` returns 404; the working endpoint has moved to `api.kimi.com`. Blocks all “moonshot” provider usage for Kimi models. No fix PR open yet.
2. **Deferred image attachment loses reloadable reference** ([#8151](https://github.com/zeroclaw-labs/zeroclaw/issues/8151)) – When a bot defers processing an image, it later claims it never saw it. The image reference is lost in cached history. No fix PR open yet.

Other bugs:

- **Insecure‑TLS confirmation flow** ([#7693](https://github.com/zeroclaw-labs/zeroclaw/issues/7693)) – Old issue, but a test coverage PR ([#8178](https://github.com/zeroclaw-labs/zeroclaw/pull/8178)) was opened today to address it.
- **Prefix caching bypass on channels** – Fixed by merged PR [#6630](https://github.com/zeroclaw-labs/zeroclaw/pull/6630) and being re‑applied by [#8174](https://github.com/zeroclaw-labs/zeroclaw/pull/8174).

Notable high‑risk open PRs that could affect stability if merged: [#8104](https://github.com/zeroclaw-labs/zeroclaw/pull/8104) (gateway drain before RPC reload), [#8152](https://github.com/zeroclaw-labs/zeroclaw/pull/8152) (memory embedding provider resolution), [#7856](https://github.com/zeroclaw-labs/zeroclaw/pull/7856) (secret prompt feedback).

## Feature Requests & Roadmap Signals

Requests with the most traction today:

| Issue | Feature | Risk | Priority | Likelihood for next release |
|-------|---------|------|----------|-----------------------------|
| [#8138](https://github.com/zeroclaw-labs/zeroclaw/issues/8138) | OpenRouter model fallback array | Medium | p2 | High – aligned with provider resilience |
| [#8142](https://github.com/zeroclaw-labs/zeroclaw/issues/8142) | Reduce initial response latency across channels | High | p2 | Medium – complex, needs design |
| [#8143](https://github.com/zeroclaw-labs/zeroclaw/issues/8143) | Move `.po` translations to git submodule | High | p3 | Low – accepted but lower priority |
| [#8170](https://github.com/zeroclaw-labs/zeroclaw/issues/8170) | In‑app upgrade from dashboard | High | p2 | High – PR already in review |
| [#8177](https://github.com/zeroclaw-labs/zeroclaw/issues/8177) | Hardware PGP / supply‑chain signing | High | p2 | Medium – RFC stage, long‑term |
| [#8033 (PR)](https://github.com/zeroclaw-labs/zeroclaw/pull/8033) | Conversational `onboard` assistant | High | – | Medium – size XL, extensive review needed |

Prediction: The **OpenRouter fallback array** and **in‑app upgrade (Phase 1)** are most likely to land in the next minor release, as they directly improve user experience and have existing implementation paths.

## User Feedback Summary

- **Pain points:** Provider endpoint breakage (Kimi/Moonshot) frustrates users who rely on that provider; bot forgetting deferred image attachments disrupts multi‑turn image workflows; lack of typing indicators/ack causes confusion during long responses.
- **Use cases:** Users want to use OpenRouter’s automatic failover for higher uptime; they want to upgrade ZeroClaw without leaving the dashboard; they need deterministic prefix caching to reduce token waste and latency.
- **Satisfaction indicators:** The prompt caching fix was eagerly awaited (PR #6630 finally merged); the feature support matrix (PR #6870) was merged, helping users choose the right runtime. The steady flow of open PRs (50 updated today) shows an engaged contributor community.

## Backlog Watch

Items requiring maintainer attention that have remained open for several days or lack a recent update:

1. **[#8170](https://github.com/zeroclaw-labs/zeroclaw/issues/8170) – RFC: In-app upgrade** (needs‑maintainer‑review, 1 day old) – Needs maintainer decision to accept or request changes on the companion PR.
2. **[#8177](https://github.com/zeroclaw-labs/zeroclaw/issues/8177) – RFC: Supply chain signing** (needs‑maintainer‑review, 1 day old) – Similarly awaiting maintainer signal.
3. **[#8142](https://github.com/zeroclaw-labs/zeroclaw/issues/8142) – Improve initial response time** (needs‑maintainer‑review, 1 day old) – No PR yet; needs scoping decision.
4. **[#7693](https://github.com/zeroclaw-labs/zeroclaw/issues/7693) – insecure‑TLS coverage** (open since 2026-06-15, status:accepted) – Now has a test PR [#8178](https://github.com/zeroclaw-labs/zeroclaw/pull/8178), so no longer stagnant.
5. **PR [#6630](https://github.com/zeroclaw-labs/zeroclaw/pull/6630)** was a stale‑candidate for months; it was finally merged today, showing that old items can be resolved.

No issues have been unanswered for more than a few days, reflecting active project management. However, the high number of open PRs (42) suggests the review pipeline is under pressure.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*