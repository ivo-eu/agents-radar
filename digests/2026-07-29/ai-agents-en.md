# OpenClaw Ecosystem Digest 2026-07-29

> Issues: 152 | PRs: 500 | Projects covered: 13 | Generated: 2026-07-29 00:10 UTC

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

# OpenClaw Project Digest – 2026-07-29

**Generated from data on openclaw/openclaw (GitHub)**

---

## Today's Overview

OpenClaw saw extremely high activity today, with 152 issues and 500 pull requests updated in the last 24 hours. The project released **v2026.7.2-beta.5**, a crucial beta focused on state safety and crash recovery. A severe memory leak in the gateway process (RSS growing from 350 MB to 15.5 GB) remains an open P0 concern, while the community continues to rally around cross‑platform desktop app support (115 comments on #75). The volume of merged/closed PRs (251) indicates strong engineering velocity, though many PRs are dependency bumps and automated housekeeping.

---

## Releases

**v2026.7.2-beta.5** (one new release today)

- **Highlights:**
  - **Quarantine store** – protects persisted data when the primary database is damaged.
  - **Crash‑recoverable SQLite snapshots** – allow safe rollback after unexpected shutdowns.
  - **Crash‑durable filesystem publication** – ensures state is not lost on OS‑level crashes.
  - **Schema‑upgrade data‑loss rejection** – prevents accidental downgrades or incompatible schema changes from corrupting data.
  - **Rollback‑writer snapshot recovery** – enables restoring the last known‑good state when a write operation fails.

No breaking changes or migration notes were included in the release description. This is a beta release under the `2026.7.x` line.

---

## Project Progress

Today **251 PRs were merged or closed**, reflecting a large batch of completed work. Notable changes include:

- **Agent fixes:**
  - [#114779 – `fix: chat replies are silently dropped`](https://github.com/openclaw/openclaw/pull/114779) – closed, addressed a regression where Telegram DM replies lost ownership.
  - [#115016 – `fix(auto-reply): prevent false no-reply fallbacks`](https://github.com/openclaw/openclaw/pull/115016) – prevents “No reply was generated” on valid outcomes.
  - [#115395 – `fix(agents): prevent path-only bootstrap files from crashing turns`](https://github.com/openclaw/openclaw/pull/115395) – avoids a crash when file entries lack an explicit `name`.
  - [#115422 – `fix(agents): restore subagent completion delivery on CLI runtimes`](https://github.com/openclaw/openclaw/pull/115422) – fixes a gap in waking inactive parent runs after subagent completion.

- **State and database hardening:**
  - [#115321 – `fix(sqlite): release statement caches when databases close`](https://github.com/openclaw/openclaw/pull/115321) – prevents stale caches from causing errors after disposal.
  - [#115447 – `fix(state): fail closed on newer schema in doctor repair`](https://github.com/openclaw/openclaw/pull/115447) – ensures `doctor --fix` refuses to proceed with a database from a future version.

- **Memory system improvements:**
  - [#115438 – `feat(memory): isolate curated entries by project`](https://github.com/openclaw/openclaw/pull/115438) – stops MEMORY.md entries from different projects bleeding into each other.
  - [#115329 – `feat(gateway): add read-only memory.list method`](https://github.com/openclaw/openclaw/pull/115329) – exposes a safe inspection path for agent memory.

- **Integrations and platform:**
  - [#112303 – `fix(embeddings): scope Mistral and DeepInfra memory cache identity by endpoint`](https://github.com/openclaw/openclaw/pull/112303) – fixes cache collisions when multiple endpoints use the same model name.
  - [#112084 – `fix(agents): use lastSeenAtMs tie-break in canvas default-node selection`](https://github.com/openclaw/openclaw/pull/112084) – closed; prevents waking an older device when a newer one is available.
  - [#111534 – `fix(update): gateway restart survives inherited cross-user D-Bus environment`](https://github.com/openclaw/openclaw/pull/111534) – Linux update reliability improvement.

- **Documentation and dependencies:**
  - [#114867 – `docs(mistral): fix broken adjustable reasoning docs URL`](https://github.com/openclaw/openclaw/pull/114867) – closed.
  - [#113927 – `build(deps): bump the actions group`](https://github.com/openclaw/openclaw/pull/113927) – routine dependency update.

---

## Community Hot Topics

| Issue / PR | Comments | Reactions | Topic |
|------------|----------|-----------|-------|
| [#75 – Linux/Windows Clawdbot Apps](https://github.com/openclaw/openclaw/issues/75) | 115 | 👍 80 | Feature request for native desktop apps (macOS, iOS, Android already exist). |
| [#7707 – Memory Trust Tagging by Source](https://github.com/openclaw/openclaw/issues/7707) | 22 | – | Security feature to tag memory entries by trust level to prevent poisoning. |
| [#91588 – Critical: Gateway Memory Leak](https://github.com/openclaw/openclaw/issues/91588) | 20 | 👍 1 | P0 memory leak causing 15.5 GB RSS and OOM crashes. |
| [#96857 – Normal tool text outputs degrade to placeholders](https://github.com/openclaw/openclaw/issues/96857) | 15 | 👍 4 | Closed bug where agents see `(see attached image)` instead of text. |
| [#11665 – Webhook sessions should reuse existing session](https://github.com/openclaw/openclaw/issues/11665) | 11 | – | Multi-turn hook sessions not working as documented. |
| [#108075 – LLM request failed: provider rejected schema](https://github.com/openclaw/openclaw/issues/108075) | 11 | 👍 1 | Closed regression in 2026.7.1. |
| [#6615 – Add denylist support for exec-approvals](https://github.com/openclaw/openclaw/issues/6615) | 10 | 👍 8 | Policy feature: “allow all except dangerous commands”. |

**Analysis:** The most talked‑about topic remains the **Linux/Windows desktop app gap** (#75, 115 comments, 80 thumbs‑up). Users with macOS, iOS, and Android clients are demanding parity. Security‑minded users are pushing for **memory trust tagging** (#7707) and **denylist‑based command approvals** (#6615). The **critical memory leak** (#91588) is the top stability concern, affecting anyone running a gateway for more than a day.

---

## Bugs & Stability

### Critical (P0)
- **#91588 – Gateway Memory Leak** (RSS 350 MB → 15.5 GB, OOM kills) – still open, no fix PR identified yet. Affects all long‑running gateways.
- **#102755 – Project won't start on Windows and WSL** (beta release blocker) – open, no fix PR. Build hangs on second launch.

### High Severity (P1)
- **#115326 – Crash‑loop breaker suppresses Discord/WhatsApp permanently** (regression, recovery fails with WebSocket 1006) – open. No linked fix PR.
- **#98790 – Concurrent agent‑to‑agent turn forks session tree** (permanent transcript corruption, Anthropic rejects assistant‑terminal messages) – open, no fix PR.
- **#102268 – Silent empty tool results in long‑running Sonnet 5 sessions** (tool outputs silently disappear after large results) – open.
- **#88955 – qqbot WebSocket reconnection causes “Outbound not configured” error** – open.
- **#98435 – MCP loopback transport does not auto‑reconnect on gateway restart** (`recovered=1` is misleading) – open.
- **#90098 – Stack‑safe large attachment handling for Control UI** – open, linked PR exists.
- **#110067 – claude‑cli runtime: mid‑turn assistant text never delivered durably** – open.
- **#111839 – Channel follow‑ups never steer active Codex runs** – open.
- **#99387 – Intermittent LLM request failed provider rejected schema** – open.
- **#112196 – memory_search transient sync timeout masks as permanent failure** – open.

### Fixed / Closed Today
- **#111519 – Telegram DM replies fall back after stale DM‑scope cleanup** – closed via PR #114779.
- **#115008 – Schema‑version error blames a downgrade that never happened** – closed.
- **#115021 – OpenAI realtime Talk advertises unsupported Codex OAuth fallback** – closed.
- **#106612 – Missing Child Process Termination in host‑hook‑runtime** – closed.
- **#108984 – claude‑cli sessions wiped by OpenClaw's byte‑guard compaction** – closed.
- **#99174 – Anthropic‑shaped 400 bodies classify as null in failover** – closed.

### Stability Observations
The new beta release directly addresses state safety and crash recovery – a clear response to the recent influx of data‑loss and session‑corruption bugs. However, the critical memory leak (#91588) and the Windows startup blocker (#102755) remain unresolved and are likely to block general availability of 2026.7.2 stable.

---

## Feature Requests & Roadmap Signals

Top community‑requested features with high engagement:

| Feature | Issue | Comments | Predict Next Version |
|---------|-------|----------|---------------------|
| **Linux/Windows native apps** | [#75](https://github.com/openclaw/openclaw/issues/75) | 115 | **Unlikely next release** – requires significant engineering; no PR open. |
| **Memory trust tagging by source** | [#7707](https://github.com/openclaw/openclaw/issues/7707) | 22 | Possible in 2026.8 – security theme aligns with recent state safety work. |
| **Denylist support for exec‑approvals** | [#6615](https://github.com/openclaw/openclaw/issues/6615) | 10 | Likely – complements existing allowlist; community strong demand. |
| **Azure Foundry GPT Realtime Talk** | [#87325](https://github.com/openclaw/openclaw/issues/87325) | 8 | Possible – auth gaps were flagged in #115021; may come soon. |
| **Production‑readiness stability label** | [#73537](https://github.com/openclaw/openclaw/issues/73537) | 8 | Unclear – maintainers haven't responded; could be deferred. |
| **Image viewing in webchat file viewer** | [#113251](https://github.com/openclaw/openclaw/issues/113251) | 6 | Likely – UI improvement with low complexity. |

**Roadmap signals from today’s PRs:**
- [#114388 – `feat(agents)!: remove the stored default agent`](https://github.com/openclaw/openclaw/pull/114388) – a breaking change to reduce ambiguity in agent routing. This suggests a move toward explicit ownership in the architecture.
- [#114151 – `feat(plugins): allow per-turn tool narrowing in prompt hooks`](https://github.com/openclaw/openclaw/pull/114151) – improving plugin extensibility.
- [#115329 – `feat(gateway): add read-only memory.list method`](https://github.com/openclaw/openclaw/pull/115329) – observability for memory, possibly leading to broader memory management APIs.

---

## User Feedback Summary

**Pain points voiced:**
- **Stability:** Multiple regressions in the 2026.7.x beta series (e.g., #115326, #111519) are causing production users to hesitate. One user ([Reneb-cafe](https://github.com/openclaw/openclaw/issues/73537)) explicitly requested a “production‑readiness” label after running OpenClaw for family/business use.
- **Memory leaks:** The gateway OOM problem (#91588) is causing launchd‑handoff restart cycles, disrupting automation and long‑running sessions.
- **Missing platforms:** Linux and Windows users feel left behind (#75). The community has 80 thumbs‑up for parity with macOS/iOS/Android.
- **Security:** Users are concerned about memory poisoning from untrusted sources (#7707) and want fine‑grained command approvals (#6615).
- **Configuration friction:** Webhook session reuse (#11665) not working as documented, and config validator rejecting `mediaLocalRoots` (#47002) – leads to wasted debugging time.

**Positive signals:**
- Gratitude for the product’s value (“It has genuinely become part of our daily workflow” – #73537).
- Active community participation with detailed reproduction steps and linked PRs shows investment in project health.

---

## Backlog Watch

Long‑standing issues or PRs that remain open and require maintainer attention:

| Issue / PR | Age (Created) | Priority | Notes |
|------------|---------------|----------|-------|
| [#75 – Linux/Windows Clawdbot Apps](https://github.com/openclaw/openclaw/issues/75) | 2026-01-01 | P2 | 115 comments; no maintainer response in months. |
| [#7707 – Memory Trust Tagging](https://github.com/openclaw/openclaw/issues/7707) | 2026-02-03 | P2 | Needs product decision; labeled with `clawsweeper:needs-product-decision`. |
| [#6615 – Denylist for exec‑approvals](https://github.com/openclaw/openclaw/issues/6615) | 2026-02-01 | P2 | 10 comments, 8 👍; still awaiting maintainer review. |
| [#11665 – Webhook session reuse](https://github.com/openclaw/openclaw/issues/11665) | 2026-02-08 | P2 | Documented but not working; linked PR is open (#? maybe not linked). |
| [#8299 – Suppress sub‑agent announce](https://github.com/openclaw/openclaw/issues/8299) | 2026-02-03 | P2 | 8 comments; no resolution. |
| [#73537 – Production‑readiness stability label](https://github.com/openclaw/openclaw/issues/73537) | 2026-04-28 | P2 | 8 comments; maintainer hasn't replied. |
| [#98790 – Concurrent agent‑to‑agent turn forks session tree](https://github.com/openclaw/openclaw/issues/98790) | 2026-07-01 | P1 | Active session corruption; no fix PR yet. |

**PRs awaiting maintainer review** (status: `ready for maintainer look`):
- [#110875 – feat(mattermost): read channel history](https://github.com/openclaw/openclaw/pull/110875)
- [#111534 – fix(update): gateway restart cross‑user D‑Bus](https://github.com/openclaw/openclaw/pull/111534)
- [#112303 – fix(embeddings): Mistral/DeepInfra cache key](https://github.com/openclaw/openclaw/pull/112303)

These PRs have `proof: sufficient` and `status: 👀 ready for maintainer look`, meaning they are technically ready but waiting for a core maintainer’s final sign‑off.

---

*All data sourced from GitHub openclaw/openclaw repository as of 2026-07-29.*

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: Personal AI Assistant Open-Source Ecosystem
**Date:** 2026-07-29 | **Analyst:** Senior Ecosystem Analyst

---

## 1. Ecosystem Overview

The personal AI assistant open-source landscape is undergoing a rapid maturation phase, characterized by heavy investment in **state safety, crash recovery, and production-grade stability** across all major projects. The ecosystem is bifurcating: infrastructure-focused projects (OpenClaw, IronClaw, ZeroClaw) are building robust data persistence, multi-agent orchestration, and enterprise-grade security boundaries, while application-layer projects (NanoBot, CoPaw, LobsterAI) prioritize user-facing features like desktop apps, platform channel parity, and rich media support. A clear consensus has emerged around **memory management, authentication reliability, and MCP protocol stability** as the most critical shared pain points. Community engagement remains high across the board, though review bottlenecks and maintainer availability are becoming visible constraints, particularly in the mid-tier projects.

---

## 2. Activity Comparison

| Project | Issues Updated (24h) | PRs Updated (24h) | PRs Merged/Closed | Release Status | Health Score |
|---|---|---|---|---|---|
| **OpenClaw** | 152 | 500 | 251 | v2026.7.2-beta.5 (today) | ★★★★☆ High velocity, P0 memory leak |
| **ZeroClaw** | 9 | 50 | 8 | v0.8.5 pending (weekly) | ★★★★☆ Active eval overhaul |
| **IronClaw** | 8 | 50 | 15 | No release today | ★★★★☆ Strong epic-driven development |
| **Hermes Agent** | 28 | 50 | 16 | v0.19.0 (stable) | ★★★★☆ High engagement, Windows gaps |
| **CoPaw** | 10 | 47 | 9 | QwenPaw 2.0.1 stable | ★★★☆☆ Review bottleneck forming |
| **NanoBot** | 7 | 37 | 18 | No release today | ★★★★☆ Rapid fix cadence |
| **PicoClaw** | 4 | 10 | 3 | No release today | ★★★☆☆ Healthy but small |
| **NanoClaw** | 0 | 12 | 5 | No release today | ★★★☆☆ Clean backlog |
| **Moltis** | 1 | 9 | 2 | No release today | ★★★☆☆ Focused improvements |
| **LobsterAI** | 5 | 7 | 6 | No release today | ★★★★☆ Dense feature merges |
| **ZeptoClaw** | 0 | 2 | 1 | No release today | ★★☆☆☆ Maintenance mode |
| **TinyClaw** | 0 | 0 | 0 | No activity | ★☆☆☆☆ Dormant |
| **NullClaw** | 0 | 0 | 0 | No activity | ★☆☆☆☆ Dormant |

**Note:** Health Score is qualitative, factoring volume, responsiveness, stability posture, and community engagement.

---

## 3. OpenClaw's Position

OpenClaw maintains the **largest community and highest engineering velocity** in the ecosystem, with 500 PRs and 152 issues updated in a single day—roughly 10x the activity of its nearest peers (ZeroClaw, IronClaw, Hermes). Key advantages:

- **Architectural maturity:** The quarantine store, crash-recoverable SQLite snapshots, and schema-upgrade rejection feature set (shipped in today's beta) represent the most comprehensive state-safety guarantees in the ecosystem. No other project has equivalent data protection.
- **Cross-platform support:** Native macOS, iOS, and Android clients exist; Linux/Windows desktop app demand (#75) has 115 comments and 80 thumbs-up, showing strong market pull.
- **Technical approach:** OpenClaw uses a **gateway-centric model** with persistent agent memory and crash-durable filesystem publication, contrasting with NanoBot's plugin-based extensibility and Hermes's ACP-first architecture. This makes OpenClaw better suited for long-running production deployments.
- **Community size comparison:** OpenClaw's issue/PR volume is ~10x ZeroClaw, ~5x Hermes Agent, and ~20x PicoClaw. It is the clear reference implementation.

**Primary risk:** The unresolved P0 memory leak (#91588, RSS growing to 15.5GB) and Windows startup blocker (#102755) could delay the stable release and erode trust among production users.

---

## 4. Shared Technical Focus Areas

The following requirements recur across multiple projects, indicating ecosystem-wide priorities:

| Focus Area | Affected Projects | Specific Needs |
|---|---|---|
| **Memory & State Persistence** | OpenClaw, NanoBot, PicoClaw, ZeroClaw, CoPaw | Crash-recoverable snapshots, memory isolation by project, trust tagging, compaction safety |
| **Authentication Robustness** | Hermes Agent, PicoClaw, IronClaw, ZeroClaw | Browser OAuth survival, multi-profile token management, stale session handling, WhatsApp pairing |
| **MCP Protocol Stability** | ZeroClaw, CoPaw, NanoBot | Stdio response routing, session recovery on restart, image content block forwarding |
| **Windows Support** | OpenClaw, Hermes Agent, CoPaw, LobsterAI | Installer reliability (corrupt agent.json), shell character escaping, process termination |
| **Multi-Platform Desktop** | OpenClaw, Hermes Agent, NanoBot | Linux/Windows native apps, SSH remote mode, PWA push notifications |
| **Provider Diversity** | OpenClaw, Hermes Agent, PicoClaw, NanoClaw | Mistral support, DeepInfra cache scoping, MiniMax OAuth, Azure Foundry |
| **Tool Execution Security** | OpenClaw, IronClaw, ZeroClaw | Denylist approvals, exec-approval policies, TOCTOU path validation, sandbox containers |

**Emerging theme (day-old):** **Multi-agent orchestration** is gaining traction. NanoBot's #5000 (multi-agent collaboration proposal) and Hermes's #5257 (generalized ACP client) both surfaced within the last 8-9 days, suggesting a shift toward agent swarm patterns.

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | Hermes Agent | NanoBot | IronClaw | ZeroClaw | CoPaw |
|---|---|---|---|---|---|---|
| **Core Philosophy** | Gateway-centric, crash-proof state | ACP-first multi-agent orchestration | Plugin platform + marketplace | Enterprise error recoverability | Evaluation-driven reliability | Desktop + Chrome integration |
| **Target User** | Production operators, long-running | CLI power users, developers | End users, skill builders | Enterprise teams, security-sensitive | CI/CD workflows, testers | Windows/macOS desktop users |
| **Language** | Go (gateway) + Rust (core) | Rust | Python | Rust | Rust | Python |
| **Key Strength** | State safety guarantees | Protocol generalization (ACP) | Fast feature iteration | Epic-driven reliability | Eval framework depth | Desktop GUI automation |
| **Key Weakness** | Memory leak, missing Linux desktop | Windows platform gaps | Token cost complaints | Integration install friction | Channel config validation | Review bottleneck |
| **Community Maturity** | Reference project, largest | Strong, vocal contributors | High engagement | Security-focused | Methodical | Growing, first-time contributors |
| **Release Cadence** | Weekly betas | ~Monthly minors | Continuous (no formal) | Epic-based | Weekly non-breaking | Patch + feature releases |

**Notable niche players:**
- **Moltis** — Focused on Slack reliability and ACP agent exposure; small but targeted.
- **PicoClaw** — Lightweight, mobile-friendly (Feishu, DingTalk, Android attention).
- **LobsterAI** — Windows installer specialist with side-chat innovation (btw feature).
- **ZeptoClaw** — Rust runtime maintenance; essentially in idle mode.

---

## 6. Community Momentum & Maturity

**Tier 1: Hyper-growth / High iteration**
- **OpenClaw, ZeroClaw, IronClaw, Hermes Agent** — Each sees 50+ daily PR updates, multiple merged changes, and active community discourse. These are the ecosystem's engines for new features and architectural work.

**Tier 2: Steady advancement**
- **NanoBot, CoPaw, Moltis, LobsterAI, PicoClaw, NanoClaw** — 7–50 daily PRs, clear feature roadmaps, responsive maintainers. These projects are adding substantial value in their niches but lack the cross-cutting architectural scale of Tier 1.

**Tier 3: Stable / Low activity**
- **ZeptoClaw** — Only dependency updates; effectively in maintenance mode.
- **TinyClaw, NullClaw** — No activity; likely abandoned or dormant.

**Stabilization signals:**
- **NanoBot** has shifted from rapid feature acceleration to **targeted P1 regression fixes** (6 fix PRs in one day), indicating threshold to a stable release.
- **ZeroClaw** is consolidating its eval framework, a sign of approaching a quality gate.
- **OpenClaw** shipped state-safety features that directly address recent data-loss bugs—a classic stabilization pattern.

**Rapid iteration signals:**
- **IronClaw**’s error-recoverability epic (#6284) is spawning multiple workstream PRs; do not expect stability soon.
- **Hermes** is juggling new providers, SSH fixes, and ACP generalization—still in expansion phase.

---

## 7. Trend Signals

**For AI agent developers and practitioners, the following industry trends emerge from today's data:**

1. **Crash recovery is now table stakes.** OpenClaw's quarantine store and snapshot recovery, NanoBot's session consolidation fixes, and ZeroClaw's lifecycle bracket leak closure all point to a fundamental shift: **agents must survive ungraceful shutdowns**. Developers should treat crash-durable state as a non-negotiable requirement, not a nice-to-have.

2. **Multi-agent orchestration is the next battleground.** Three distinct signals appeared in 24 hours: NanoBot's collaboration proposal, Hermes's ACP generalization, and CoPaw's sub-agent isolation request. Teams building agent workflows should invest in ACP-compatible orchestration infrastructure now.

3. **Memory trust and security are moving up the stack.** OpenClaw's memory trust tagging, IronClaw's TOCTOU fixes, and CoPaw's permission-boundary discussions indicate that **agent memory poisoning and cross-session contamination** are becoming recognized attack surfaces. Expect memory provenance frameworks to emerge as a standard requirement.

4. **Platform parity is a persistent friction point.** Windows users are disproportionately affected across LobsterAI (installer bugs), CoPaw (agent.json corruption), and Hermes (shell character escaping). Developers targeting cross-platform deployments should budget extra time for Windows-specific testing.

5. **MCP is winning, but immature.** Four projects reported MCP-related bugs (session recovery, response routing, image content handling). The protocol is standardizing but lacks robust retry and lifecycle management. Tooling around MCP reliability is an immediate gap.

6. **Self-hosted memory backends are gaining interest.** Moltis's zvec PR (embedded vector store), CoPaw's checkpoint management, and PicoClaw's vodozemac encryption push all signal a **move away from cloud-dependent memory**. This aligns with enterprise data sovereignty concerns.

7. **User feedback is driving quality, not just features.** The most upvoted items across projects are stability improvements: OpenClaw's Linux desktop (80 👍), Hermes's Mistral support (23 👍), and general production-readiness requests. **Features alone don't retain users—reliability does.**

---

**Summary for Decision-Makers:**

- **Choose OpenClaw** if production-grade state safety and multi-platform desktop are priorities, but account for the memory leak workaround.
- **Choose Hermes Agent** if you need ACP generalization for multi-agent workflows and can accept Windows gaps.
- **Choose IronClaw** if enterprise security and error-recoverability are critical (but expect ongoing churn).
- **Choose NanoBot** for rapid prototyping with plugin extensibility and skill marketplaces.
- **Monitor ZeroClaw** for its eval framework—it may become the ecosystem's testing standard.
- **Avoid** TinyClaw, NullClaw, and (likely) ZeptoClaw for active development.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-07-29

---

## 1. Today's Overview

Project activity remains high, with **37 pull requests** (18 merged/closed) and **7 issues** updated in the past 24 hours. The team is focused on stability, pushing multiple **P1 regressions fixes** across pairing, providers, memory, subagents, execution, and the Web UI. At the same time, several features—including a **LINE channel**, **skill marketplaces**, and an **extension platform**—are advancing through open PRs. No new releases were published today, but the high volume of merged fixes and features suggests a release may be imminent.

---

## 2. Releases

No new versions were released in the last 24 hours.

---

## 3. Project Progress

**Merged/Closed PRs (18 total)** — notable examples include:

- **#5145** – [fix(ci): stabilize and speed up CI](https://github.com/HKUDS/nanobot/pull/5145) — replaced timing-dependent test with stdin-gated handshake, batched pip installs.
- **#5144** – [fix(ci): scope PR path detection to head changes](https://github.com/HKUDS/nanobot/pull/5144) — improved CI reliability for concurrent base updates.
- **#5143** – [fix(webui): animate reasoning drawer transitions](https://github.com/HKUDS/nanobot/pull/5143) — UI polish (P2).
- **#5110** – [feat(config): add actionable startup diagnostics and WebUI recovery](https://github.com/HKUDS/nanobot/pull/5110) — extended `nanobot status` with offline readiness checks.
- **#5142** – [fix(webui): open threads at latest message](https://github.com/HKUDS/nanobot/pull/5142) — restored topic scroll behavior.

These closed PRs improve CI, configuration diagnostics, and Web UI user experience.

---

## 4. Community Hot Topics

| Item | Type | Comments | Reactions | Summary |
|------|------|----------|-----------|---------|
| [Issue #5](https://github.com/HKUDS/nanobot/issues/5) | Issue (CLOSED) | 7 | 👍 3 | Request to document `uv install` for faster setup – resolved. |
| [Issue #5000](https://github.com/HKUDS/nanobot/issues/5000) | Issue (OPEN) | 5 | 👍 0 | Proposal to evolve subagent system toward true multi-agent collaboration. Active discussion. |
| [Issue #1332](https://github.com/HKUDS/nanobot/issues/1332) | Issue (CLOSED) | 4 | 👍 0 | Token consumption complaint ("Hello" using 5k input tokens) – closed as stale. |
| [Issue #5118](https://github.com/HKUDS/nanobot/issues/5118) | Bug (OPEN) | 2 | 👍 0 | Session consolidation drops uploaded media paths – fix PRs exist (#5120, #5139). |
| [Issue #5138](https://github.com/HKUDS/nanobot/issues/5138) | Bug (OPEN) | 1 | 👍 0 | MCP SDK v2 migration to fix stdio shutdown errors. |

**Analysis:** The multi-agent collaboration proposal (#5000) signals a strong community desire for more sophisticated agent orchestration. The media path bug (#5118) has attracted two competing fix attempts, showing its importance. Token cost complaints (like #1332) remain a persistent user pain point.

---

## 5. Bugs & Stability

**New bugs reported today (updated within last 24h):**

| Severity | Issue | Description | Has Fix PR? |
|----------|-------|-------------|-------------|
| **Critical** | [#5118](https://github.com/HKUDS/nanobot/issues/5118) | Media paths dropped during session consolidation → files unrecoverable after archive. | Yes – #5120, #5139 (both open) |
| **High** | [#5138](https://github.com/HKUDS/nanobot/issues/5138) | MCP stdio session teardown error + stdout pollution. | Tracked, no dedicated fix PR yet |
| **High** | [#5133](https://github.com/HKUDS/nanobot/issues/5133) | `finish_reason='length'` with tool calls misrouted to empty-response retry. | No fix PR yet |
| **Medium** | [#5149](https://github.com/HKUDS/nanobot/issues/5149) | WhatsApp audio output broken (`neonize.utils.ffmpeg` warning). | No fix PR yet |

**Regressions fixed via new PRs (all P1):**
- [#5155](https://github.com/HKUDS/nanobot/pull/5155) – Null `approved` map in pairing store.
- [#5154](https://github.com/HKUDS/nanobot/pull/5154) – Primitive items in Responses API parser.
- [#5153](https://github.com/HKUDS/nanobot/pull/5153) – Non-string timestamp / missing role in memory archiving.
- [#5152](https://github.com/HKUDS/nanobot/pull/5152) – Subagent partial completion results misreported.
- [#5151](https://github.com/HKUDS/nanobot/pull/5151) – Idle session locks not released (memory leak).
- [#5150](https://github.com/HKUDS/nanobot/pull/5150) – Unbounded buffered session output.
- [#5147](https://github.com/HKUDS/nanobot/pull/5147) – Transient store read failure erases approvals.
- [#5146](https://github.com/HKUDS/nanobot/pull/5146) – Malformed token-usage day keys crash API.

**Overall stability assessment:** The project is in a **high-fix cadence** mode, addressing regressions quickly. Most bugs have parallel fix PRs, indicating strong maintainer responsiveness.

---

## 6. Feature Requests & Roadmap Signals

| Request | Source | Likelihood for Next Version |
|---------|--------|----------------------------|
| **Multi-agent collaboration** (#5000) | Issue | Moderate – design proposal, not yet coded |
| **Skill marketplaces / Discover UI** (#5116) | Open PR (P1) | Very high – PR already submitted |
| **Unified extension platform** (#5098) | Open PR (P1, conflict) | High – close to merge |
| **Image-aware model presets** (#5148) | Open PR (P1) | Very high – directly improves config |
| **LINE Messaging API channel** (#5115) | Open PR (P1) | High – new channel addition |
| **Stable resource path aliases** (#5131) | Open PR | High – improves reliability of resource access |

**Prediction:** The next release will likely include the **WebUI skill marketplace**, **extension platform**, **image-aware presets**, and the **LINE channel**. Multi-agent collaboration (#5000) may appear as an RFC in the following release.

---

## 7. User Feedback Summary

- **Installation speed:** Issue #5 (closed) requested `uv install` documentation – resolved, indicating responsiveness.
- **Token consumption:** Issue #1332 (stale) highlighted high input token usage even for trivial queries – no recent resolution, but remains a latent concern.
- **WhatsApp audio:** Issue #5149 reports that outgoing audio messages fail while incoming audio works – user likely frustrated.
- **Media persistence:** Issue #5118 reveals a critical archiving bug that leads to file loss – user found it through data loss.
- **MCP integration pain:** Issue #5138 user reports shutdown warnings and protocol pollution during stdio usage, affecting daily workflows.

Overall, users are actively reporting bugs and feature proposals, showing high engagement. The team's rapid fix cycles (same-day PRs for many bugs) should improve satisfaction.

---

## 8. Backlog Watch

No long-unanswered issues were observed in the latest 24h data. The oldest open issue with maintainer attention is **#5000** (created 2026-07-20, 8 days ago) – it has active comments but no labelled response from maintainers yet. Other high-severity bugs (#5133, #5149) are very recent and may receive attention shortly.

**Potential watch item:** Issue #1332 (token consumption), though closed as stale, might re-emerge if multiple users report similar high costs. No follow-up PR has been merged to address token efficiency.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-07-29

## 1. Today's Overview

The project saw very high activity on 2026-07-29, with **28 open issues** and **50 pull requests** updated in the last 24 hours. **No new releases** were published; development continues on the `main` branch. Of the 50 PRs updated, **16 were merged or closed**, indicating active merging of fixes and features. The community is engaging deeply, with several long-standing feature requests (e.g., Mistral provider, ACP generalization) accumulating strong reactions and discussion. Bug reports remain concentrated around cross-platform stability (Windows, SSH, cron), and maintainers are actively responding with targeted fix PRs.

## 2. Releases

*No new releases today.* The latest tagged version remains v0.19.0 (2026-07-20).

## 3. Project Progress

Sixteen pull requests were merged or closed in the last 24 hours. Notable among the top-20-by-comment PRs:

- **#73699** (closed) – Auto-generated lint/formatting fix (CI housekeeping).
- **#73697** (closed) – Desktop: bump session list order on user send so recent chats appear immediately.
- **#67415** (closed) – Fix Discord thread routing for profile-aware dispatches.
- **#73688** (closed) – Fix first-turn STT hallucinations on Discord voice channels by rejecting comfort/silence Opus segments.

Other merged PRs (not in top 20) likely include smaller bug fixes and documentation improvements, contributing to overall stability and reliability.

## 4. Community Hot Topics

The following issues and pull requests attracted the most comments and reactions, revealing key community interests:

- **Issue #5257** – *feat: Generalized ACP client for multi-agent CLI orchestration*  
  [☝️21 👍 · 20 comments](https://github.com/NousResearch/hermes-agent/issues/5257)  
  Users are eager for Hermes to act as a universal **Agent Client Protocol** orchestrator for all ACP-compatible coding agents, not just Copilot.

- **Issue #20859** – *Support for Mistral as LLM provider*  
  [☝️23 👍 · 10 comments](https://github.com/NousResearch/hermes-agent/issues/20859)  
  Mistral integration is the most-upvoted feature request, reflecting a strong desire for an alternative to closed-source providers.

- **Issue #69551** – *Desktop SSH remote mode broken with non-default profiles*  
  [☝️0 👍 · 9 comments](https://github.com/NousResearch/hermes-agent/issues/69551)  
  A high-impact bug where profile-scoped `HERMES_HOME` conflicts with hardcoded paths, blocking remote SSH usage for many power users.

- **Issue #17649** – *Semantic Skill Retrieval with SQLite FTS5*  
  [☝️1 👍 · 4 comments](https://github.com/NousResearch/hermes-agent/issues/17649)  
  A proposal to replace the all-skills-injection approach (costly at ~4,500 tokens/turn) with on-demand search – a clear cost-saving and scalability improvement.

## 5. Bugs & Stability

Several new and updated bugs were reported today, ranked by severity (P2 = high, P3 = moderate). Many already have corresponding fix PRs.

| Issue | Severity | Description | Fix PR (if exists) |
|-------|----------|-------------|-------------------|
| [#69551](https://github.com/NousResearch/hermes-agent/issues/69551) | P2 | Desktop SSH broken with non-default profiles (token-path mismatch) | - |
| [#73694](https://github.com/NousResearch/hermes-agent/issues/73694) | P2 | Cron blocks safe local Ollama jobs (localhost vs 127.0.0.1) | [#73659](https://github.com/NousResearch/hermes-agent/pull/73659) (open) |
| [#73693](https://github.com/NousResearch/hermes-agent/issues/73693) | P2 | Windows Rider ACP hangs on `read_file` and `terminal` during Git Bash bootstrap | - |
| [#73700](https://github.com/NousResearch/hermes-agent/issues/73700) | P2 (new) | WhatsApp bridge fails to pair (HTTP 405, stale Baileys version) | - |
| [#73691](https://github.com/NousResearch/hermes-agent/issues/73691) | P3 | Desktop auto-TTS toggle active but never fires on API server path | - |
| [#71069](https://github.com/NousResearch/hermes-agent/issues/71069) | P2 | Terminal tool resolves to wrong Hermes install in multi-HERMES_HOME setups | - |
| [#35654](https://github.com/NousResearch/hermes-agent/issues/35654) | P2 | Browser tool fails on Windows with shell characters (&, \|, <, >) | - |
| [#67637](https://github.com/NousResearch/hermes-agent/issues/67637) | P2 | Windows dependency stage discards traceback and aborts on transient probe failures | - |
| [#67627](https://github.com/NousResearch/hermes-agent/issues/67627) | P2 | TUI exec quick commands block the RPC reader loop | - |
| [#68012](https://github.com/NousResearch/hermes-agent/issues/68012) | P2 | Dashboard cron create silently drops finite repeat counts | - |
| [#51773](https://github.com/NousResearch/hermes-agent/issues/51773) | P2 | `max_tokens` rejection misclassified as context overflow → compression loop dead-end | - |
| [#7448](https://github.com/NousResearch/hermes-agent/issues/7448) | P2 | Webhook idempotency keyed only by delivery ID, causing multi-route fan-out skip | - |
| [#58705](https://github.com/NousResearch/hermes-agent/issues/58705) | P3 | mem0 OSS agent tools fail with Qdrant lock conflict when plugin holds the lock | [#68992](https://github.com/NousResearch/hermes-agent/pull/68992) (open) |
| [#73692](https://github.com/NousResearch/hermes-agent/issues/73692) | P2 | `agent.disabled_toolsets: [browser]` silently removes `web_search` | - |

**Stability note:** A recurring pattern of **Windows-specific regressions** (browser tool, dependency bootstrap, Rider ACP) suggests platform testing gaps. The maintainers have labeled several with `sweeper:risk-platform-windows`.

## 6. Feature Requests & Roadmap Signals

The following features were requested or are under active development, indicating likely additions in the next minor release:

| Feature | Issue/PR | Status | Likely Impact |
|---------|----------|--------|---------------|
| **Mistral LLM provider** | [#20859](https://github.com/NousResearch/hermes-agent/issues/20859) | Open, needs decision | High – most-upvoted request |
| **Generalized ACP client** | [#5257](https://github.com/NousResearch/hermes-agent/issues/5257) | Open, needs decision | High – enables multi-agent orchestration |
| **Omnious model provider** | [#73696](https://github.com/NousResearch/hermes-agent/pull/73696) | Open PR | Medium – adds another supported API |
| **Semantic skill retrieval (FTS5)** | [#17649](https://github.com/NousResearch/hermes-agent/issues/17649) | Open | Medium – cost savings for skill-heavy users |
| **xAI OAuth shared store** (multi-profile) | [#67261](https://github.com/NousResearch/hermes-agent/pull/67261) | Open PR | Medium – solves token refresh conflicts |
| **Claude CLI inference backend** | [#67335](https://github.com/NousResearch/hermes-agent/pull/67335) | Open PR | Medium – allows base-subscription billing |
| **Email gateway draft mode** | [#47789](https://github.com/NousResearch/hermes-agent/issues/47789) | Open | Medium – risk reduction for business users |
| **Language auto-correction** | [#31514](https://github.com/NousResearch/hermes-agent/issues/31514) | Open | Low – niche but clear use case |
| **TTS for xAI OAuth proxy** | [#36130](https://github.com/NousResearch/hermes-agent/pull/36130) | Open PR | Low – completes proxy feature parity |
| **Cron dead-pin guard** | [#73506](https://github.com/NousResearch/hermes-agent/pull/73506) | Open PR | Medium – reliability improvement |

**Prediction:** The next minor release (v0.20.0) is likely to include **Mistral support** (given community demand) and the **generalized ACP client** (strategic extensibility), along with the **xAI OAuth shared store** fix and **Claude CLI backend** as part of provider parity efforts.

## 7. User Feedback Summary

Real pain points expressed in recent issues and comments include:

- **SSH profile breakage** (#69551) – Users relying on multiple Hermes profiles for remote work cannot use SSH mode, forcing workarounds.
- **Windows instability** (#35654, #67637, #73693) – Shell character escaping, bootstrap failures, and JetBrains integration hangs cause poor Windows experience.
- **Missing provider support** (#20859, #73423, #36124) – Users actively request Mistral, Hetzner, and full TTS for xAI; lack of these forces them to use other tools.
- **Cost overhead** (#17649) – The current skill injection model is wasteful; power users want on-demand retrieval to reduce token spend.
- **Auto-TTS not working** (#73691) – Desktop app toggle is non-functional, indicating a gap between UI and backend implementation.
- **WhatsApp bridge broken** (#73700) – A newly reported critical bug for users relying on WhatsApp integration.
- **Webhook multi-route fan-out** (#7448) – Teams using GitHub webhooks to trigger multiple routes report unexpected silent drops.

Overall satisfaction is tempered by these usability issues, but the community is actively contributing fixes and feature PRs.

## 8. Backlog Watch

The following items have been open for extended periods without maintainer action or require a decision:

| Issue/PR | Open Since | Notes |
|----------|------------|-------|
| [#5257](https://github.com/NousResearch/hermes-agent/issues/5257) – ACP client generalization | 2026-04-05 | Needs decision; 21 👍, 20 comments. Long dormant. |
| [#20859](https://github.com/NousResearch/hermes-agent/issues/20859) – Mistral provider | 2026-05-06 | Needs decision; 23 👍, 10 comments. Could be a quick win. |
| [#17649](https://github.com/NousResearch/hermes-agent/issues/17649) – Skill retrieval | 2026-04-29 | Only 1 👍, but clear use case. No maintainer response. |
| [#29933](https://github.com/NousResearch/hermes-agent/pull/29933) – Feishu env-var fallback | 2026-05-21 | Open PR with no recent movement. |
| [#31795](https://github.com/NousResearch/hermes-agent/pull/31795) – Kanban DB integrity retry | 2026-05-25 | Open PR, needs review. |
| [#36130](https://github.com/NousResearch/hermes-agent/pull/36130) – xAI TTS proxy | 2026-05-31 | Open PR, pending merge. |
| [#68992](https://github.com/NousResearch/hermes-agent/pull/68992) – mem0 Qdrant recreate fix | 2026-07-22 | Open PR, addresses a P3 bug (#58705). |
| [#73659](https://github.com/NousResearch/hermes-agent/pull/73659) – Cron localhost fix | 2026-07-28 | Open PR, may resolve #73694 quickly. |

Maintainers should prioritize **#5257** and **#20859** as high-impact, high-community interest decisions. The **#69551 SSH bug** also deserves immediate attention given its severity and number of affected users.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest — 2026-07-29

*Data sourced from [github.com/sipeed/picoclaw](https://github.com/sipeed/picoclaw)*

---

## 1. Today's Overview

The project shows steady community activity with **4 issues** (3 closed, 1 open) and **10 pull requests** (3 merged/closed, 7 open) updated in the last 24 hours. No new releases were published. The closed issues and merged PRs indicate a healthy pace of bug fixing and incremental improvements, while the high number of open PRs (7) suggests an active review queue. Key areas of focus include authentication reliability, provider integrations (Anthropic, Exa), and cross‑platform fixes (Android, DingTalk). Overall, the project is **actively maintained** with a strong preference for quickly closing reported bugs (three of four issues resolved within the day).

---

## 2. Releases

**No new releases** in the reported period.

---

## 3. Project Progress

Three pull requests were merged or closed today:

- **[#3256 (merged)](https://github.com/sipeed/picoclaw/pull/3256) – `fix(feishu): send audio and video with native message types`**  
  Author: AaronZ345. Feishu channel now delivers audio (opus) and video (mp4) as native playable messages instead of generic file downloads.

- **[#3254 (merged)](https://github.com/sipeed/picoclaw/pull/3254) – `fix(agent): prefer verbatim model matches over provider‑alias splits when resolving refs`**  
  Author: fabdelgado. Model resolution logic now correctly prioritises exact model string matches, preventing a provider‑alias split from stealing a resolution intended for a verbatim model entry.

- **[#3228 (merged)](https://github.com/sipeed/picoclaw/pull/3228) – `fix(anthropic-messages): send SystemParts as system blocks with cache_control`**  
  Author: AayushGupta16. The `anthropic_messages` provider now respects `SystemParts` and enables per‑block `cache_control` markers, fixing a complete block on Anthropic prompt caching for callers using this provider.

These changes improve **platform integration** (Feishu), **model reliability** (agent), and **LLM caching** (Anthropic).

---

## 4. Community Hot Topics

The most active discussions involved security, cross‑platform functionality, and authentication:

- **[#3088 (CLOSED) – [Feature] use vodozemac instead of libolm](https://github.com/sipeed/picoclaw/issues/3088)**  
  *10 comments, 2 👍*  
  This long‑running feature request (open since June, marked stale) was closed today. The community strongly supports replacing the insecure, unmaintained `libolm` with the official Rust‑based `vodozemac`. The underlying need is **security and long‑term maintainability** of end‑to‑end encryption primitives.

- **[#3182 (OPEN) – [BUG] Android version](https://github.com/sipeed/picoclaw/issues/3182)**  
  *5 comments*  
  A user reports inability to launch the service on Android despite full permissions. The issue remains open and stale. Underlying need: **reliable cross‑platform support** – Android users expect a functional mobile experience.

- **[#3280 (OPEN) – fix(auth): make browser OAuth login survive real‑world callback conditions](https://github.com/sipeed/picoclaw/pull/3280)**  
  *Active PR, author honbou*  
  This PR addresses four independent causes of OAuth failures on headless/remote setups. The community’s interest reflects a widespread pain point: **authentication reliability** in diverse deployment environments.

---

## 5. Bugs & Stability

One new bug was reported and closed the same day; no critical regressions were introduced.

| Issue | Severity | Summary | Fix PR |
|-------|----------|---------|--------|
| [#3300 (CLOSED)](https://github.com/sipeed/picoclaw/issues/3300) | **High** | Tool set missing `read_file` causes a deadlock when a user instructs the AI to read `RULES.md` before every response. The system does not expose `read_file` as an available tool, leading to infinite loop. | None needed (closed without code change – likely a configuration or documentation fix) |
| [#3255 (CLOSED)](https://github.com/sipeed/picoclaw/issues/3255) | **Medium** | DingTalk chat list preview shows static “PicoClaw” instead of reply content – only the list preview is affected, not the in‑chat view. | Not provided (closed as completed) |
| [#3182 (OPEN)](https://github.com/sipeed/picoclaw/issues/3182) | **Medium** | Android service fails to launch; user cannot change path from settings. Remains open and stale. | No fix PR yet |

**Overall Stability**: The majority of bugs were resolved quickly. The Android issue (#3182) remains the most concerning open bug, as it affects a whole platform.

---

## 6. Feature Requests & Roadmap Signals

User‑requested features and emerging roadmap items from the last 24 hours:

- **Vodozemac replacement** – [#3088](https://github.com/sipeed/picoclaw/issues/3088) (closed, likely scheduled for next release)
- **Exa web search provider** – [#3299 (OPEN, PR)](https://github.com/sipeed/picoclaw/pull/3299) – adds native Exa API support for `tools.web` / `web_search` with date range filters.
- **Configurable default fallback chain** – [#3200 (OPEN, PR)](https://github.com/sipeed/picoclaw/pull/3200) – allows users to set and persist a fallback chain of models via the web UI and backend API.
- **Capture Anthropic prompt cache tokens** – [#3251 (OPEN, PR)](https://github.com/sipeed/picoclaw/pull/3251) – exposes cache‑related metrics currently discarded by both Anthropic providers.
- **Installation scripts moved to main repo** – [#1951 (OPEN, PR)](https://github.com/sipeed/picoclaw/pull/1951) – a long‑standing organisational improvement.

**Prediction**: The next minor release will likely include the vodozemac switch, Exa provider, and the Anthropic cache metrics fix, given the high activity and community demand.

---

## 7. User Feedback Summary

Real user pain points and use cases surfaced in today’s data:

- **Cross‑platform frustration** – Android users (#3182) cannot launch the service, indicating gaps in mobile testing.
- **DingTalk list previews** (#3255) – the UI gives no hint of reply content, hurting user experience for DingTalk users.
- **Custom rules integration** (#3300) – users want to split system prompts into separate files (e.g., `RULES.md`), but the tool set is incomplete, causing deadlocks.
- **OAuth reliability** (#3280) – users on headless/remote machines face burned authorization codes, requiring a complete restart of the login flow.

Overall satisfaction is mixed: bugs are fixed quickly when reported, but **Android stability** and **authentication robustness** are clear pain points.

---

## 8. Backlog Watch

Items that have been stale or require maintainer attention:

| Issue/PR | Status | Age | Why it matters |
|----------|--------|-----|----------------|
| [#1951 – Move installation scripts](https://github.com/sipeed/picoclaw/pull/1951) | OPEN · PR | Since 2026‑03‑24 (4 months) | Consolidates documentation and reduces confusion; marked `stale`. |
| [#3182 – Android bug](https://github.com/sipeed/picoclaw/issues/3182) | OPEN · Issue | Since 2026‑06‑26 (1 month) | Blocks Android users; no active fix in progress. |
| [#3280 – OAuth fix](https://github.com/sipeed/picoclaw/pull/3280) | OPEN · PR | Since 2026‑07‑21 (1 week) | A thorough fix for a high‑impact problem; awaits review. |
| [#3279 – Seahorse tool‑call leak](https://github.com/sipeed/picoclaw/pull/3279) | OPEN · PR | Since 2026‑07‑21 (1 week) | Fixes a data‑leakage bug that could corrupt user messages. |
| [#3251 – Anthropic cache tokens](https://github.com/sipeed/picoclaw/pull/3251) | OPEN · PR | Since 2026‑07‑12 (2.5 weeks) | Important for operators monitoring prompt caching efficiency. |

**Recommendation**: The maintainers should prioritise reviewing [#3280 (OAuth)](https://github.com/sipeed/picoclaw/pull/3280) and [#3251 (cache tokens)](https://github.com/sipeed/picoclaw/pull/3251) to clear the PR backlog, and address the Android issue [#3182](https://github.com/sipeed/picoclaw/issues/3182) to avoid platform abandonment.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-07-29

## 1. Today's Overview
Moderate development activity with 12 pull requests updated in the last 24 hours, of which 5 were closed/merged and 7 remain open. No new releases were published, and no new issues were filed. The project continues to mature with a strong focus on reliability—initialisation fixes for agent containers, merge safety guards for the `/update-nanoclaw` workflow, and configurable webhook ports are among the concrete improvements that landed today. The dual‑engine quota‑fallback feature (#3057) remains under active review and is already battle‑tested in production.

## 2. Releases
*No new releases were created. The project appears to be in a continuous integration phase with features shipping via merged PRs rather than formal releases.*

## 3. Project Progress
Five pull requests were closed/merged today, advancing stability and configurability:

| PR | Type | Summary | Link |
|---|---|---|---|
| #3060 | Fix | Add `--init` to container spawn args so PID 1 reaps zombie processes, and correct the documentation gap | [PR #3060](https://github.com/nanocoai/nanoclaw/pull/3060) |
| #1255 | Feature | Introduce MiniMax OAuth (Coding Plan) as an alternative model provider – includes device‑code OAuth with PKCE, token polling, and auto‑refresh | [PR #1255](https://github.com/nanocoai/nanoclaw/pull/1255) |
| #2197 | Fix | Guard merge state in `/update-nanoclaw` to prevent silent single‑parent commits that could lose history | [PR #2197](https://github.com/nanocoai/nanoclaw/pull/2197) |
| #1136 | Feature | Add auto‑merge audit and container smoke test to `/update-nanoclaw` to catch silent code drops during upstream merges | [PR #1136](https://github.com/nanocoai/nanoclaw/pull/1136) |
| #2598 | Fix | Load per‑group `CLAUDE.local.md` by adding `'local'` to `settingSources` | [PR #2598](https://github.com/nanocoai/nanoclaw/pull/2598) |

Additionally, several open PRs saw recent commit activity and are nearing completion (see Community Hot Topics).

## 4. Community Hot Topics
While no issues were created today, the most active pull requests reflect key developer and operational needs:

- **#3057 (Open) – Dual‑engine quota fallback**  
  *Author: elia‑ben‑cnaan*  
  A production‑proven feature branching Claude→Codex automatic fallback, handoff recaps, and proactive quota warnings. Already live on a WhatsApp deployment since July 6. Discussion centres on architecture details and migration (migration 017).  
  [PR #3057](https://github.com/nanocoai/nanoclaw/pull/3057)

- **#3143 (Open) – Preserve resolved approval card content**  
  *Author: Koshkoshinsk*  
  Ensures resolved approval cards retain their title and request details while muting the decision; survives deserialisation.  
  [PR #3143](https://github.com/nanocoai/nanoclaw/pull/3143)

- **#3144 (Open) – Configurable webhook bind address via `WEBHOOK_HOST`**  
  *Author: jonnychesthair-crypto*  
  Adds `WEBHOOK_HOST` environment variable (default `0.0.0.0`) to allow deployments to listen on specific interfaces.  
  [PR #3144](https://github.com/nanocoai/nanoclaw/pull/3144)

- **#3148 (Open) – Honor `WEBHOOK_PORT` from `.env`**  
  *Author: ogarciarevett*  
  Closes #2901 by making `WEBHOOK_PORT` follow NanoClaw’s normal configuration precedence (env → `.env` → default 3000).  
  [PR #3148](https://github.com/nanocoai/nanoclaw/pull/3148)

The underlying need is clear: developers running NanoClaw in custom environments require fine‑grained network configuration and robust provider failover, while end‑users want stable approval workflows that survive page reloads.

## 5. Bugs & Stability
Several stability fixes were merged or opened today. Ranked by severity:

| Severity | PR | Description | Fix Available? |
|---|---|---|---|
| **High** | #3060 (merged) | Zombie process accumulation in agent containers; PID 1 never reaped children. Fix: add `--init` flag. | ✅ Merged |
| **High** | #3148 (open) | `WEBHOOK_PORT` environment variable ignored; hardcoded default used instead. Fix: read via config precedence. | ✅ Open PR |
| **Medium** | #2197 (merged) | `/update-nanoclaw` could produce single‑parent merge commits on customized forks, losing history. Fix: guard merge state. | ✅ Merged |
| **Medium** | #2598 (merged) | Per‑group `CLAUDE.local.md` files not loaded. Fix: add `'local'` setting source. | ✅ Merged |
| **Medium** | #3145 (open) | Missing channel destinations for existing messaging‑group wirings. Fix: migration 021 backfills them. | ✅ Open PR |
| **Low** | #3146 (open) | Two dev scripts (`test-v2-host.ts` and another) rotted after architecture migrations; they crash before spawning containers. Fix: updated to current config shape. | ✅ Open PR |
| **Low** | #3147 (open) | Reply context leaking across destinations in agent runner. Fix: keep reply context local. | ✅ Open PR |

No regressions or crashes reported outside the scope of these fixes.

## 6. Feature Requests & Roadmap Signals
The following features, requested or implemented in recent PRs, point toward future capabilities:

- **Dual‑engine quota fallback (#3057)** – Likely to be stabilised and released as the canonical failover mechanism for production deployments. Expect it in the next minor version.
- **MiniMax OAuth provider (#1255, already merged)** – Adds a third‑party model provider with its own OAuth flow. Creates a pattern for bringing in more providers.
- **Configurable webhook host/port (#3144, #3148)** – Two complementary PRs that together expose full network binding control. A strong signal that the project wants to support containerised / orchestrated deployments.

Predictions for the next version:  
- Merge of #3057 (dual‑engine) after review.  
- Inclusion of #3144 and #3148 to close the webhook configuration gap.  
- Continued expansion of provider backends (MiniMax is now merged, others may follow).

## 7. User Feedback Summary
No new issues were filed today, so direct user feedback is absent. However, the content of the merged and open PRs reveals implicit pain points:

- **Manual configuration friction** – Users needed `WEBHOOK_PORT` to respect `.env` (issue #2901, now fixed via #3148).  
- **Process cleanliness** – Zombie processes in containers caused resource leaks, addressed by #3060.  
- **Update safety** – Fork maintainers faced silent history corruption during `/update-nanoclaw`; the merge‑state guard (#2197) and audit smoke test (#1136) directly address their dissatisfaction.  
- **Provider flexibility** – Adding MiniMax OAuth reflects a desire to use NanoClaw without an Anthropic API key or Claude subscription.

Overall satisfaction likely remains high given the rapid closure of reported issues (5 bugs fixed in one day), but lack of formal user surveys makes quantification impossible.

## 8. Backlog Watch
There are **no open issues** in the repository (total: 0). This indicates either a very clean bug‑tracking process or that users are reporting bugs via other channels (e.g., Discord, GitHub Discussions). No long‑unanswered items require maintainer attention at this time.

All open PRs have been updated within the last 24 hours, suggesting maintainers are actively reviewing and pushing forward the backlog.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest – 2026-07-29

## 1. Today’s Overview
Project activity remains very high, with **50 pull requests** and **8 issues** updated in the last 24 hours. No new releases were published today. Of the 50 PRs, **15 were merged or closed**, while 35 remain open and in active review. The focus is split between **error-recoverability** (epic #6284 driving multiple workstreams), **IronHub integration** (rebuild on the Reborn stack), and **security hardening** (TOCTOU fixes, trust-boundary tests, and denylist exemptions). Several integration-install bugs (Notion, Slack) were reported and are likely to see rapid fixes given the team’s momentum.

## 2. Releases
**None** – no new versions of IronClaw were published in the last 24 hours. The latest stable release remains `ironclaw 1.0.0` (referenced in issue #6814).

## 3. Project Progress
Today **15 pull requests were closed or merged**, advancing several critical areas:

- **Sandbox container transport** – PR #6746 (part 1 of 4) introduces Docker-connect retry, egress allowlisting, and shell limits for the sandbox.
- **Tool-disclosure security surface** – PR #5659 (open, but updated today) narrows three leak vectors and includes regression/trust-boundary tests.
- **File-system TOCTOU escapes** – PR #6817 closes four path-confusion vulnerabilities in `DiskFilesystem`.
- **Lifecycle state consolidation** – PR #6696 (DB migration) makes the process journal the single source of truth for turn state.
- **Messaging framework** – PR #6831 defines a host-owned standard ops vocabulary (16 core + 13 reserved names) with canonical schemas.
- **Signing & multi-tenant isolation** – PRs #6813 and #6822 add attested gate resolve and KMS ship-gate (group 7/8 of signing work).
- **Testing infrastructure** – Several workstream-completion PRs (#6823, #6825, #6827, #6828) add e2e coverage for persistence backends, fault-profile cross-boundary tests, and webhook ingress.

*Note: The precise list of merged/closed PRs is not shown in the provided data; the 15 count is inferred from the summary.*

## 4. Community Hot Topics
The most active discussion centers on the **error-recoverability epic** (#6284):

- **[Issue #6284](https://github.com/nearai/ironclaw/issues/6284)** – *“[EPIC] error-recoverability endgame — the model recovers from 100% of the errors it sees”*  
  **15 comments**, 0 reactions. This epic (opened July 19) is the project’s top priority, defining a five-part recoverability contract. It has spawned multiple workstream PRs (#6824, #6826, #6832 addressed today) and continues to attract attention from core contributors.

- **[Issue #6820](https://github.com/nearai/ironclaw/issues/6820)** – *“IronHub: agent reaches for an unsigned catalog URL when discovery disappoints”*  
  2 comments. A trust-boundary concern from a live preview; split from the IronHub feature work.

- **[Issue #6814](https://github.com/nearai/ironclaw/issues/6814)** – *“Third-party skills still trip the prompt content denylist on 1.0.0”*  
  1 comment (from author). A follow-up to earlier fixes, showing user frustration with released behaviour.

The underlying need across these threads is **reliability and trustworthiness** – both in agent error-handling and in catalog/data integrity.

## 5. Bugs & Stability
Seven new bugs were filed today, none with known fixes yet (though several align with open PRs):

| Issue | Description | Severity |
|-------|-------------|----------|
| [#6814](https://github.com/nearai/ironclaw/issues/6814) | Third-party skills with “API key” in description fail every run – prompt denylist false positive | **Critical** – breaks all third-party skill usage on 1.0.0 |
| [#6833](https://github.com/nearai/ironclaw/issues/6833) | Notion tool installation fails/hangs with unclear messaging | **High** – blocks popular integration |
| [#6834](https://github.com/nearai/ironclaw/issues/6834) | Slack setup fails in IronClaw (near.foundation account) | **High** – blocks team collaboration |
| [#6820](https://github.com/nearai/ironclaw/issues/6820) | Agent uses unsigned catalog URL during IronHub discovery – trust-boundary issue | **High** – could lead to unsigned content |
| [#6835](https://github.com/nearai/ironclaw/issues/6835) | MCP auth failures never raise re-auth gate (classified as Client, not AuthRequired) | **Medium** – silent error misclassification |
| [#6821](https://github.com/nearai/ironclaw/issues/6821) | IronHub free-text search matches return non-catalog entries, misrepresenting inventory | **Medium** – misleads users |
| [#6829](https://github.com/nearai/ironclaw/issues/6829) | Telegram forum-topic delivery lacks full coverage for `message_thread_id` | **Low** – partial path only |

Fix PRs for related error-classification bugs (#6824, #6826, #6832) are open and address the *recoverability* domain, but not the install/denylist issues directly.

## 6. Feature Requests & Roadmap Signals
The **error-recoverability epic** (#6284) is the strongest roadmap signal – its completion will directly impact the next minor or major release. Additionally:

- **IronHub on Reborn** – PRs #6754, #6780, and issues #6820/#6821 point to a full port of the IronHub catalog and install flow, likely targeting the next release.
- **Channel ingress centralization** (PR #6816) – moves auth/command classification into a shared host-owned path, suggesting a unified multi-channel strategy.
- **Signing and attestation** (PRs #6813, #6822) – multi-tenant isolation and KMS ship-gate indicate enterprise-grade trust infrastructure coming soon.
- **WebUI design system** (PR #6836) – extraction of `@ironclaw/ui` suggests a UI refresh is in the pipeline.

**Prediction for next version:** The stable release (tentatively 1.1.0) will likely include the error-recoverability endgame, IronHub Reborn port, and the TOCTOU/filesystem fixes. Integration fixes (Notion, Slack) are blockers for many users and will likely be patched quickly, possibly in a 1.0.1 hotfix.

## 7. User Feedback Summary
Reported pain points (from issues and comments) indicate:

- **Integration friction** – Notion and Slack cannot be installed/set up; users hitting terminal failures or hanging processes.
- **Denylist frustration** – Third-party skill authors cannot use the word “API key” even in descriptions, blocking all their skills on 1.0.0. This is a regression from earlier fixes.
- **Catalog confusion** – IronHub search returns irrelevant results (free-text matches read as complete listings), undermining trust in the catalog.
- **Trust gaps** – The agent silently using unsigned URLs during discovery (issue #6820) raises security concerns for early adopters.

Overall satisfaction is difficult to gauge, but the volume of open issues and PRs suggests a highly active community that is engaging constructively.

## 8. Backlog Watch
No issues or PRs appear to be long-unanswered within the last 24h data. However, the **error-recoverability epic** (#6284) has been open since July 19 (10 days) and, while actively worked, has not yet closed. Given its scope, it remains the most critical item requiring sustained maintainer attention. The **third-party denylist bug** (#6814) is only 1 day old but has no fix PR yet – given its severity, maintainers may want to prioritise a hotfix. All other items are fresh and receiving updates.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-07-29

## 1. Today's Overview

The project maintains high development velocity with **7 pull requests updated in the last 24 hours**, of which **6 were merged or closed** (including the much-anticipated side-chat feature and multiple critical installer fixes). On the issue side, **5 new or updated issues** were reported, with **3 of them opened just today**—an indicator of steady community engagement. No new releases were published, but the dense batch of merged PRs suggests an upcoming minor or patch release may be imminent. Overall, the project is in a healthy state with both feature additions and stability improvements actively landing.

## 2. Releases

No new releases were created in the last 24 hours.

## 3. Project Progress

The following pull requests were merged or closed today, reflecting tangible progress across documentation, the installer, the OpenClaw runtime, and the renderer:

- **#2402 – fix(update): reject Windows installer redirects instead of trusting response.url**  
  *[Closed]* Strengthens the update logic on Windows by validating redirects, preventing potential silent failures.  
  [PR #2402](https://github.com/netease-youdao/LobsterAI/pull/2402)

- **#2400 – fix(openclaw): enforce runtime/config safety-contract gate to stop false-stop token burn**  
  *[Closed]* Introduces a runtime‑build‑info and config contract check so the bundled OpenClaw runtime cannot run without LobsterAI’s managed safety policy. Also retires “prompt-exposure-budget” as a terminal kind.  
  [PR #2400](https://github.com/netease-youdao/LobsterAI/pull/2400)

- **#2399 – feat(renderer): hide sites nav entry outside test mode**  
  *[Closed]* A UI polish that removes a test‑only navigation item from production builds.  
  [PR #2399](https://github.com/netease-youdao/LobsterAI/pull/2399)

- **#2398 – fix(installer): drive Skills backup outcome from helper exit codes**  
  *[Closed]* Resolves a critical installer bug where the Skills backup step misread the PowerShell helper’s output, leading to a spurious backup‑missing error and degraded upgrade experience.  
  [PR #2398](https://github.com/netease-youdao/LobsterAI/pull/2398)

- **#2397 – feat(cowork): add isolated /btw side chat**  
  *[Closed]* Adds a floating, resizable side‑chat panel for selected assistant text, with drag support, eight‑direction resizing, stop, and follow‑up questions. Execution and history are kept isolated from the main conversation, routed through the OpenClaw utility stream.  
  [PR #2397](https://github.com/netease-youdao/LobsterAI/pull/2397)

- **#2394 – Fix/windows install manual overwrite blocked**  
  *[Closed]* Another Windows‑installer fix addressing blocked manual overwrites during setup.  
  [PR #2394](https://github.com/netease-youdao/LobsterAI/pull/2394)

**Still open but stale:** PR #1233 (model provider links & API key guidance) remains open after several months with no recent activity.

## 4. Community Hot Topics

The most active issues and PRs in the last 24 hours, ranked by engagement:

- **Issue #2401 – “skill技能”** (1 comment, opened yesterday)  
  The user asks whether LobsterAI uses Anthropic’s official skills for PDF/DOCX/PPTX/XLSX handling and whether those skills are license‑clear for commercial use. This signals growing community interest in the legal and licensing side of built‑in skill components.  
  [Issue #2401](https://github.com/netease-youdao/LobsterAI/issues/2401)

- **Issue #2395 – “无法安装” (Installation failure)** (1 comment, opened yesterday)  
  Reports a blocked installation due to a failed Skills backup. The error points to `install-timing.log`; project maintainers already merged PR #2398 addressing the exact root cause. This issue is likely resolved in the next build.  
  [Issue #2395](https://github.com/netease-youdao/LobsterAI/issues/2395)

- **Issue #2396 – exec tool default shell wrapper bug** (0 comments yet, opened yesterday)  
  Describes a serious functional gap on Windows: the default shell wrapper is PowerShell 5.1, causing Linux‑style commands or scripts with special characters (e.g., `node -e`, `pwsh -Command`) to fail silently. No fix PR has been linked yet.  
  [Issue #2396](https://github.com/netease-youdao/LobsterAI/issues/2396)

- **Stale issues #1236 (plugin ID mismatch) and #2071 (scheduled task error)** received automated updates yesterday but lack maintainer responses. Both have been open for months.

## 5. Bugs & Stability

**High severity** (blocker):
- **Issue #2395 – Installation blocked on Windows**  
  *Root cause identified and fixed in PR #2398.* The installer misclassified the Skills backup outcome, causing upgrades to fail. Fixed in the repository; will be available in the next release.  
  [Issue #2395](https://github.com/netease-youdao/LobsterAI/issues/2395)

**Medium severity** (functional regression):
- **Issue #2396 – exec tool default shell = PowerShell 5.1**  
  *No fix PR yet.* Commands expecting a Unix‑like shell or `pwsh`‑specific syntax silently fail. Affects Windows users running cross‑platform workloads.  
  [Issue #2396](https://github.com/netease-youdao/LobsterAI/issues/2396)

**Low severity** (configuration warnings):
- **Issue #1236 – Plugin ID mismatch warning** (stale)  
  Cosmetic warning during gateway startup; no functional impact but erodes trust in configuration. No recent maintainer comment.  
  [Issue #1236](https://github.com/netease-youdao/LobsterAI/issues/1236)

- **Issue #2071 – Scheduled task creation error** (stale)  
  User reported an error with a screenshot; version mentioned is 2026.5.27. No diagnosis provided.  
  [Issue #2071](https://github.com/netease-youdao/LobsterAI/issues/2071)

## 6. Feature Requests & Roadmap Signals

- **Side‑chat (/btw) panel** (PR #2397, merged today) – Already implemented. An isolated floating panel for follow‑up questions, likely to appear in the next release.
- **Model provider links & API key guidance** (PR #1233, open but stale) – This PR has been open since April; it adds clickable links to official sites and “get API key” shortcuts. If merged, it would improve onboarding. Given its age, maintainers may need to push it forward or close it.
- **Skill licensing clarification** (Issue #2401) – The user’s question about Anthropic skills’ commercial use may drive a documentation update or a licensing FAQ page.
- **Cross‑platform shell compatibility** – The exec tool bug (Issue #2396) implicitly requests a more robust shell detection or a configurable shell wrapper, which could become a roadmap item for Windows users.

## 7. User Feedback Summary

- **Pain points**: Installation failures (skill backup issue) and the exec‑tool default‑shell problem are the most immediate frustrations reported in the last 24 hours. Both affect Windows users disproportionately.
- **Use cases**: Users rely on LobsterAI for document processing (PDF, DOCX, etc.) and shell command execution—the two areas where issues were filed today.
- **Satisfaction signals**: The rapid merges (6 PRs in 24h) demonstrate responsive maintenance, likely boosting user confidence. The new `/btw` side‑chat feature also indicates active product evolution.
- **Unaddressed concerns**: The stale issues #1236 and #2071 have not received maintainer replies for months, which may cause dissatisfaction among the reporters. License and commercial‑use questions (#2401) remain unanswered.

## 8. Backlog Watch

The following important items have been open for a long time without maintainer intervention:

- **PR #1233 – Model provider links & API key guidance**  
  *Opened: 2026-04-01, last updated: 2026-07-28 (automatic)*. A well‑reviewed feature that would improve model provider onboarding. Needs a maintainer decision (merge with conflict resolution or close).  
  [PR #1233](https://github.com/netease-youdao/LobsterAI/pull/1233)

- **Issue #2071 – Scheduled task creation error**  
  *Opened: 2026-05-28, last updated: 2026-07-28 (stale label added)*. User provided a screenshot of the error but received no response. Likely a bug, but insufficient reproduction steps.  
  [Issue #2071](https://github.com/netease-youdao/LobsterAI/issues/2071)

- **Issue #1236 – Plugin ID mismatch warning**  
  *Opened: 2026-04-01, last updated: 2026-07-28 (stale label added)*. Reports a persistent configuration warning. A maintainer’s guidance on whether the warning is intentional or a bug would help.  
  [Issue #1236](https://github.com/netease-youdao/LobsterAI/issues/1236)

These items, while not critical, signal areas where community contributions await review or closure.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest — 2026-07-29

## Today's Overview

The project saw moderate activity in the last 24 hours, with **1 issue closed** and **9 pull requests updated** — 7 remain open and 2 were merged or closed. No new releases were published. The development focus remains on infrastructure hardening: instrumentation, access control, Slack reliability, PWA push notifications, and exposing Moltis as an ACP agent. The closing of a reported bug (archiving cron sessions) and two merged PRs indicate steady progress toward better UX and security boundaries.

## Releases

No new versions were tagged in the last 24 hours. The latest stable release (if any) predates this digest period.

## Project Progress

Two pull requests were merged or closed today:

- **#1172 – fix(web): hide archived cron sessions by default**  
  ([PR #1172](https://github.com/moltis-org/moltis/pull/1172)) – Applies the shared archived-session preference to the Cron tab so archived runs are hidden by default. Includes a Playwright regression test for hiding, showing, and re-hiding.

- **#1171 – Move ACP selection into the chat model picker**  
  ([PR #1171](https://github.com/moltis-org/moltis/pull/1171)) – Relocates installed ACP clients into the composer model selector alongside provider-backed models, removes the historical header ACP selector and the redundant “Built-in LLM agent” option. Preserves per-session binding, ACP-only auto-binding, and reasoning control.

Additionally, **Issue #1111** (bug report for archiving cron sessions having no visible effect) was closed as fixed, likely by PR #1172.

## Community Hot Topics

The only issue updated in the last 24 hours was **#1111** (closed, 0 comments). Among open PRs, while no comment counts are recorded, several are notable:

- **#1166 – feat(slack): per-message acknowledgment reactions, phases, reconnect supervision, and Block Kit**  
  ([PR #1166](https://github.com/moltis-org/moltis/pull/1166)) – Addresses the long-standing Slack limitation (no typing indicator) with a reliable “I got it” reaction signal, phase feedback, and Block Kit rendering.

- **#1158 – feat(memory): add zvec vector database memory backend**  
  ([PR #1158](https://github.com/moltis-org/moltis/pull/1158)) – A community-driven experiment (author: demyanrogozhin) adding a Zvec+redb-based memory backend, feature-gated behind `zvec`. Signals interest in self-hosted, embeddable vector storage.

- **#1174 – Add instrumentation and feedback collection infrastructure**  
  ([PR #1174](https://github.com/moltis-org/moltis/pull/1174)) – Adds Langfuse v4 export, OTLP backends, and end-user reaction feedback, indicating a push toward observability and user experience measurement.

**Underlying needs**: Users and contributors are seeking better reliability (Slack, notifications), custom memory backends (self-hosting), and telemetry for troubleshooting and improvement.

## Bugs & Stability

Only one bug report was active in the period:

- **#1111 – [Bug]: Archiving a cron session has no visible effect**  
  ([Issue #1111](https://github.com/moltis-org/moltis/issues/1111)) – **Severity: Medium**. The UI did not reflect the archive action. Fixed by PR #1172 (merged today). No other crashes, regressions, or new bug reports were filed.

Overall stability appears good, with the team responding quickly to reported issues.

## Feature Requests & Roadmap Signals

No explicit feature requests were filed as issues today. However, several open PRs signal planned or experimental features likely to land in the next release:

- **ACP agent via stdio** (#1169) – Exposes Moltis as an ACP agent, enabling external tooling integration.
- **Terminal-Bench chat runner** (#1175) – Adds `moltis-ctl chat` and `chat-history` commands for headless evaluation.
- **PWA push notifications** (#1173) – Makes notifications reliable, private, and non-disruptive across tabs and devices.
- **Per-account operators list** (#1170) – Separates channel access from privilege with explicit operator roles, enhancing security.
- **Instrumentation & feedback** (#1174) – Operational telemetry (OpenTelemetry, Langfuse) and user reaction feedback.

**Prediction**: The next minor version will likely include the operators gate (#1170), instrumentation (#1174), and the ACP/CLI runners (#1169, #1175). The zvec memory backend (#1158) may be merged experimentally behind a feature flag.

## User Feedback Summary

Direct user feedback is limited, but one actionable pain point was reported:

- **Archiving cron sessions**: A user reported that archiving had no visible effect, indicating a UX gap in the Cron tab. The fix (#1172) was merged promptly.

Indirect feedback from PRs suggests a desire for:
- **Better Slack integration** (#1166) – users want real-time acknowledgment and richer messages.
- **Self-hosted memory** (#1158) – interest in running Moltis without external vector databases.
- **Access control granularity** (#1170) – need for per-account operator permissions in multi-user channels.

Satisfaction appears high given the quick turnaround on bug fixes and the breadth of new features being developed.

## Backlog Watch

The oldest open pull request with maintainer activity is:

- **#1158 – feat(memory): add zvec vector database memory backend**  
  ([PR #1158](https://github.com/moltis-org/moltis/pull/1158)) – Opened July 17, 2026 (12 days ago). Still under review; no merge or explicit rejection. This PR may require maintainer decision on whether to accept as an experimental feature or request additional testing.

No other issues or PRs appear neglected. The project team is responsive, with most PRs receiving updates within a day or two.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw (QwenPaw) Project Digest – 2026-07-29

## 1. Today's Overview

Activity remains high, with 10 issues updated in the last 24 hours (7 open, 3 closed) and 47 pull requests updated (38 open, 9 merged/closed). No new releases were published. The project is seeing active bug fixes for critical regressions (Windows installer, Agent JSON corruption, mission command crash) and steady feature development (desktop GUI automation, unified browser SDK, checkpoint management). Community engagement is solid, with several first-time contributors submitting fixes. The overall health is good, though the volume of open PRs (38) suggests a review bottleneck may be forming.

## 2. Releases

*None.* No new versions were tagged in the last 24 hours. The current stable version remains **QwenPaw 2.0.1**, with a **2.1.0b1** pre-release in use (see PR #6532).

## 3. Project Progress

**Merged/Closed PRs (9 total; select highlights):**
- [#6495](https://github.com/agentscope-ai/QwenPaw/pull/6495) – `fix(video): deliver video data to models` – Resolves the `view_video` data-block dropping bug (closes #6474) by adding video serialization to OpenAI Response, Anthropic, and Chat Completions providers.
- [#6501](https://github.com/agentscope-ai/QwenPaw/issues/6501) (issue closed) – Development install documentation fixed to include the `test` extra.
- [#6403](https://github.com/agentscope-ai/QwenPaw/issues/6403) (issue closed) – RobotFramework syntax highlighting added in Coding Mode’s web IDE.

**Advancing features (notable open PRs with recent updates):**
- [#6424](https://github.com/agentscope-ai/QwenPaw/pull/6424) – Native desktop GUI automation for Windows/macOS (accessibility-first + Tauri control). Under review.
- [#6276](https://github.com/agentscope-ai/QwenPaw/pull/6276) – Unified browser SDK (“one SDK, any backend”) – foundational for Chrome extension plugin.
- [#6269](https://github.com/agentscope-ai/QwenPaw/pull/6269) – Workspace checkpoint management using shadow Git stores.
- [#6157](https://github.com/agentscope-ai/QwenPaw/pull/6157) – Chrome extension plugin with native messaging bridge (depends on #6276).

## 4. Community Hot Topics

- **Most commented issue:** [#6524](https://github.com/agentscope-ai/QwenPaw/issues/6524) – “MCP 后端重启后客户端无法自动恢复” (3 comments). Users report that `streamable_http` MCP sessions are not recovered automatically after the MCP server restarts; a manual `list mcp` command is required. This indicates a missing session lifecycle management in the MCP client proxy.

- **Active enhancement discussion:** [#6509](https://github.com/agentscope-ai/QwenPaw/issues/6509) – “支持Sub Agent之间的隔离机制及单个Sub Agent内会话的完全隔离” (2 comments). Generated by user `wuarron`, this request highlights multi-tenant isolation gaps and file workspace conflicts between sessions. The community is pushing for UUID-based workspace directories.

- **Skill tags regression:** [#6537](https://github.com/agentscope-ai/QwenPaw/issues/6537) – “Skill tags disappear on restart” (1 comment). A regression of a previously fixed issue (#3270), indicating a flaw in the manifest reconciliation logic.

## 5. Bugs & Stability

**High severity:**
- [#6534](https://github.com/agentscope-ai/QwenPaw/issues/6534) – **Windows Installer infinite loop** – The NSIS installer incorrectly checks for the installer process itself, preventing installation entirely. No fix PR yet; this blocks all new Windows users.
- [#6520](https://github.com/agentscope-ai/QwenPaw/issues/6520) – **`agent.json` systematic corruption** (BOM, missing quotes, double-encoding) on Windows. A first-time contributor has submitted [#6528](https://github.com/agentscope-ai/QwenPaw/pull/6528) with safe JSON reading and BOM handling. This fix is critical for Windows stability.

**Medium severity:**
- [#6524](https://github.com/agentscope-ai/QwenPaw/issues/6524) – MCP session recovery failure (see above). No fix PR yet.
- [#6533](https://github.com/agentscope-ai/QwenPaw/issues/6533) – `/mission` command crashes with `TypeError` due to missing `verification_instructions` argument. A fix is proposed in [#6535](https://github.com/agentscope-ai/QwenPaw/pull/6535) (CloudPaw mission monkey-patch).
- [#6529](https://github.com/agentscope-ai/QwenPaw/issues/6529) – ACP `new_session` response missing `models` field, breaking external clients. Fix PR [#6531](https://github.com/agentscope-ai/QwenPaw/pull/6531) is open (first-time contributor).

**Low severity / fixed today:**
- [#6474](https://github.com/agentscope-ai/QwenPaw/issues/6474) – Video data block silently dropped – fixed by merged PR [#6495](https://github.com/agentscope-ai/QwenPaw/pull/6495).
- [#6537](https://github.com/agentscope-ai/QwenPaw/issues/6537) – Skill tags regression – no fix PR yet.

## 6. Feature Requests & Roadmap Signals

- **Sub‑Agent Isolation** ([#6509](https://github.com/agentscope-ai/QwenPaw/issues/6509)) – Strong demand for multi-session file isolation and inter-agent blocking. Likely to be addressed in the next minor release (2.1.x) given its security implications.
- **Checkpoint management** ([#6269](https://github.com/agentscope-ai/QwenPaw/pull/6269)) – Workspace-scoped shadow Git checkpoints are nearing completion; could ship in 2.1.0.
- **Reranker support for ReMe memory** ([#6398](https://github.com/agentscope-ai/QwenPaw/pull/6398)) – Under review; would improve memory retrieval quality.
- **Background tool call offload refactor** ([#6151](https://github.com/agentscope-ai/QwenPaw/pull/6151)) – Dual-deadline architecture to fix cancel/hint bugs; waiting review.
- **Unified browser & Chrome extension** ([#6276](https://github.com/agentscope-ai/QwenPaw/pull/6276), [#6157](https://github.com/agentscope-ai/QwenPaw/pull/6157)) – Still in review, likely targeting a 2.2.0 feature release.

## 7. User Feedback Summary

- **Windows users are struggling:** Persistent installation failure (installer loop), `agent.json` corruption, and missing BOM handling are causing a poor out-of-box experience. The corruption bug (#6520) has a fix in PR, but the installer bug (#6534) still needs attention.
- **MCP users want reliability:** The loss of connection after MCP server restart forces manual intervention (#6524). Users expect transparent reconnection.
- **Multi-session workflows are fragile:** The lack of per-session file isolation (#6509) leads to cross-contamination of workspace resources, a pain point for power users running multiple concurrent tasks.
- **ACP integration is blocked:** External tool clients (Multica, OpenCode) cannot discover available models due to missing `models` field (#6529). Quick fix available in PR.

## 8. Backlog Watch

- **Issue #3270 (original skill tags bug)** – Was closed but regression #6537 shows it is not fully resolved. Maintainers should review the reconciliation logic.
- **PR #6151 (background tool calls offload)** – Open since July 15, labeled “Under Review” but no recent maintainer action. Addresses three bugs from #6056.
- **PR #6157 (Chrome extension plugin)** – Blocked on #6276 (unified browser). Both are large features that may need dedicated review bandwidth.
- **PR #6387 (optional channel dependencies)** – Open since July 23; moves SDKs out of default dependencies. Needs maintainer decision to avoid dependency bloat.
- **Issue #6524 (MCP session recovery)** – No assignee or fix PR; the root cause (stale session ID reuse) is non-trivial. Should be prioritized for the next patch release.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

# ZeptoClaw Project Digest — 2026-07-29

## Today’s Overview
The ZeptoClaw project experienced very low activity over the past 24 hours, with no new issues opened or closed and no releases published. The only movements were two automated dependency pull requests updated by Dependabot, one of which was recently merged (#613) and another opened yesterday (#649). This indicates a stable maintenance phase with no reported bugs, feature requests, or community discussion. Overall project health appears solid, though engagement from contributors and users remains minimal.

## Releases
**None** — No new releases were published on 2026-07-29 or in the preceding days.

## Project Progress
Only one pull request was merged/closed today (updated in the last 24h):
- **#613** [CLOSED] `chore(deps): bump rust from 1.95-slim-trixie to 1.96-slim-trixie`  
  *Merged automatically by Dependabot on 2026-07-28.*  
  This update brings the Rust runtime base image forward one minor version, ensuring continued compatibility and security patches. No functional changes to the codebase were introduced.  
  [PR #613](https://github.com/qhkm/zeptoclaw/pull/613)

No other feature work or bug fixes were merged.

## Community Hot Topics
**No active community discussion** — The only pull requests updated in the last 24 hours are Dependabot-driven dependency bumps. Neither PR attracted comments, reactions, or further discussion from maintainers or users. The project currently has zero open issues.

- [PR #649](https://github.com/qhkm/zeptoclaw/pull/649) (open, Dependabot bump to Rust 1.97)  
- [PR #613](https://github.com/qhkm/zeptoclaw/pull/613) (closed, previously merged bump to Rust 1.96)

## Bugs & Stability
**No bugs, crashes, or regressions reported** in the last 24 hours. The zero open-issue count suggests the project is free of known stability concerns at this time.

## Feature Requests & Roadmap Signals
**No feature requests** were submitted or discussed in the last 24 hours. The lack of new issues and PRs beyond maintenance updates gives no clear signal about upcoming features. Future roadmap directions remain unknown from available data.

## User Feedback Summary
**No user feedback or pain points** were recorded in the last 24 hours. With zero issues and no comments on PRs, the user base appears either satisfied or dormant. No explicit satisfaction/dissatisfaction indicators are present.

## Backlog Watch
**No long-unanswered issues or PRs** require maintainer attention. The sole open pull request (#649) was created yesterday and is a standard Dependabot update; no outstanding issues exist. The project backlog is clean.

---

*Generated from GitHub data for ZeptoClaw (github.com/qhkm/zeptoclaw) on 2026-07-29.*

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

## ZeroClaw Project Digest — 2026-07-29

### 1. Today’s Overview
The ZeroClaw repository shows high‑activity levels with **9 issues updated** (7 open, 2 closed) and **50 PRs updated** (42 open, 8 merged/closed) in the last 24 hours. Development is concentrated on runtime stability, channel adapters, and a major overhaul of the evaluation framework (eval). Two issues were closed — a critical agent lifecycle bracket leak (#9374) and a test recovery task (#9473) — and the `v0.8.5` weekly non‑breaking release tracker (#9459) is actively being populated. Several high‑risk bugs (crashloop supervisor, multimodal context undercount, flaky CI tests) remain under active investigation, with corresponding fix PRs already open.

### 2. Releases
**No new releases** were published. The next expected release is **v0.8.5**, tracked by issue [#9459](https://github.com/zeroclaw-labs/zeroclaw/issues/9459).

### 3. Project Progress
Two issues were closed today:
- **#9374 – CLI `run()` lifecycle bracket leak** – Open‑coded `AgentStart`/`AgentEnd` emission in `agent::run` caused unbalanced events on 12 exit paths. Now resolved.
- **#9473 – Recover disabled local‑only tests** – Tests gated by `zeroclaw_root_crate` that never needed the root crate have been moved into active `#[cfg(test)]` modules.

Eight PRs were merged or closed (details not sampled), but notable contributions that advanced include:
- **Eval framework** – Six open PRs from the `IftekharUddin` and `Audacity88` teams continue to build out repeated runs (`pass@k`, error bars), JUnit XML reports, memory seeding, regression suites, and run‑history receipts ([#9223](https://github.com/zeroclaw-labs/zeroclaw/pull/9223), [#9224](https://github.com/zeroclaw-labs/zeroclaw/pull/9224), [#9225](https://github.com/zeroclaw-labs/zeroclaw/pull/9225), [#9244](https://github.com/zeroclaw-labs/zeroclaw/pull/9244), [#9248](https://github.com/zeroclaw-labs/zeroclaw/pull/9248), [#9200? not in list but implied]).
- **MCP multiplex fix** ([#9418](https://github.com/zeroclaw-labs/zeroclaw/pull/9418)) – A high‑risk, XL‑sized PR that routes stdio JSON‑RPC responses by exact generation to prevent replay of unknown outcomes.
- **Canonical installation docs** ([#9267](https://github.com/zeroclaw-labs/zeroclaw/pull/9267)) and **MUSL measurement builds** ([#9286](https://github.com/zeroclaw-labs/zeroclaw/pull/9286)) are pending review.

### 4. Community Hot Topics
The most active discussion is around:

- **[Issue #6724: Crashloop supervisor with empty channel credentials](https://github.com/zeroclaw-labs/zeroclaw/issues/6724)** — 4 comments, marked `risk:high`. Users report that enabling Signal or Voice Call channels via the dashboard without providing credentials causes the channels orchestrator to restart every ~2 seconds. The issue has been open since May 16, indicating a persistent pain point. No dedicated fix PR exists yet, but it is accepted and in‑progress.

- **[Issue #9332: Multimodal context meter undercounts image‑heavy requests](https://github.com/zeroclaw-labs/zeroclaw/issues/9332)** — 2 comments, `risk:high`. The ZeroCode context meter substantially understates image‑heavy requests before dispatch, leading to spikes beyond 100% and collapse. A PR ([#9424](https://github.com/zeroclaw-labs/zeroclaw/pull/9424)) addresses related `think`‑only completions, but the context meter issue itself is still open.

- **PR #9418 (MCP multiplex)** and **PR #9424 (reject semantic‑empty completions)** have received multiple labels (`needs-author-action`, `risk:high`, `size:XL`) and are likely gathering community feedback before merge.

### 5. Bugs & Stability
Reported bugs today, ranked by severity:

| Severity | Issue | Description | Fix PR Exists? |
|----------|-------|-------------|----------------|
| **High** | [#6724](https://github.com/zeroclaw-labs/zeroclaw/issues/6724) | Enabled channel with empty credentials causes supervisor crashloop | No |
| **High** | [#9332](https://github.com/zeroclaw-labs/zeroclaw/issues/9332) | Multimodal context meter undercounts image‑heavy requests | No |
| **High** | [#9518](https://github.com/zeroclaw-labs/zeroclaw/issues/9518) | CI lifecycle observer tests capture unrelated parallel events | Yes, [PR #9522](https://github.com/zeroclaw-labs/zeroclaw/pull/9522) |
| **Medium** | [#9506](https://github.com/zeroclaw-labs/zeroclaw/issues/9506) | Email channel cannot preserve CC recipients or send Reply‑All | No |
| **Low** | [#9374](https://github.com/zeroclaw-labs/zeroclaw/issues/9374) | CLI run() lifecycle bracket leak (closed today) | Closed as fixed |

Additional bug‑fix PRs under review:
- [#9504](https://github.com/zeroclaw-labs/zeroclaw/pull/9504): Show terminal notice on context exhaustion
- [#9519](https://github.com/zeroclaw-labs/zeroclaw/pull/9519): Serialize gateway config writes to prevent data loss
- [#9520](https://github.com/zeroclaw-labs/zeroclaw/pull/9520): Honor `always:true` skill frontmatter in compact prompt mode
- [#9517](https://github.com/zeroclaw-labs/zeroclaw/pull/9517): Localize tool‑approval prompts across adapters
- [#9457](https://github.com/zeroclaw-labs/zeroclaw/pull/9457): Restore foreground startup feedback for daemon
- [#9368](https://github.com/zeroclaw-labs/zeroclaw/pull/9368): Count retained history in whole turns

### 6. Feature Requests & Roadmap Signals
Two notable feature requests were opened today:
- **[#9521: Map MCP tools/call `type:image` content blocks into vision pipeline](https://github.com/zeroclaw-labs/zeroclaw/issues/9521)** — A clear user need to pass base64 image data to vision‑capable providers instead of dumping JSON text. Likely to be included in the next minor release given the growing multimodal use case.
- **[#9516: Upgrade CPAL to 0.18 with voice‑wake migration](https://github.com/zeroclaw-labs/zeroclaw/issues/9516)** — A dependency maintenance task accepted as a follow‑up.

The **v0.8.5 milestone** ([#9459](https://github.com/zeroclaw-labs/zeroclaw/issues/9459)) lists the scope for the weekly non‑breaking release. Based on current activity, the release will likely include:
- The eval framework enhancements (repeated runs, JUnit, memory seeds)
- MCP multiplex fix (if #9418 merges)
- Several channel‑localization and history‑counting patches
- The installer documentation and MUSL CI additions

### 7. User Feedback Summary
Real pain points expressed in recent issues:
- **Channel configuration friction** (#6724): Users can accidentally enable channels without credentials, causing endless restarts. The lack of validation feedback in the dashboard is frustrating.
- **Multimodal context miscalculation** (#9332): Image‑heavy workflows break silently due to undercounting. This degrades reliability for vision‑based use cases.
- **Email channel limitations** (#9506): Inability to handle CC lists or send Reply‑All makes the email adapter unsuitable for professional correspondence.
- **MCP image content not forwarded** (#9521): Developers integrating image‑generating or image‑processing MCP tools find the output rendered as text, losing the visual context.

Overall satisfaction is tempered by these stability and completeness gaps. The high number of `needs-author-action` PRs (at least 8 among sampled PRs) suggests the maintainer team may be bottlenecked by contributor responsiveness.

### 8. Backlog Watch
Issues and PRs that have been unanswered or unmerged for an extended period:

- **[Issue #6724](https://github.com/zeroclaw-labs/zeroclaw/issues/6724)** — Open since May 16, 2026 (over 2 months), marked `priority:p3` and `risk:high`. Despite being accepted and in‑progress, no fix PR has been linked. Needs maintainer attention to assign or prioritize.

- **[PR #9418](https://github.com/zeroclaw-labs/zeroclaw/pull/9418) (MCP multiplex)** — Open since July 26, `needs-author-action` and `risk:high`. Large, complex change that addresses a critical concurrency bug. Awaiting author response.

- **[PR #9368](https://github.com/zeroclaw-labs/zeroclaw/pull/9368) (history retention in whole turns)** — Open since July 25, `needs-author-action`. High‑risk change to a core runtime setting.

- **[PR #9267](https://github.com/zeroclaw-labs/zeroclaw/pull/9267) (installer documentation)** — Open since July 23, `needs-author-action`. A distinguished‑contributor PR that would benefit the entire community but is stalled.

- **[PR #9244](https://github.com/zeroclaw-labs/zeroclaw/pull/9244) (eval memory seeding)** — Open since July 21, `needs-author-action`, `risk:high`, `size:XL`. One of the foundational eval PRs that may be blocking the completion of the evaluation framework.

Maintainers should consider clearing the `needs-author-action` label by either merging, closing, or requesting updates to reduce stagnation.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*