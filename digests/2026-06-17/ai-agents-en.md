# OpenClaw Ecosystem Digest 2026-06-17

> Issues: 197 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-17 03:58 UTC

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

# OpenClaw Project Digest – 2026-06-17

## 1. Today's Overview

OpenClaw saw extremely high activity in the last 24 hours, with **197 issues updated** (188 open/active, 9 closed) and **500 pull requests updated** (410 open, 90 merged/closed). A new release **v2026.6.8** shipped, focusing on richer channel delivery for Telegram and WhatsApp. The project remains in a vibrant development phase, with significant community engagement on long-standing bugs and feature requests, though the sheer volume of open issues (188) and PRs (410) suggests maintainer bandwidth may be strained.

## 2. Releases

**New release: [v2026.6.8](https://github.com/openclaw/openclaw/releases/tag/v2026.6.8)** ✅

- **Highlights:**
  - **Richer channel delivery:** Telegram now renders structured text (tables, lists, expandable blockquotes, preserved line breaks, CLI-backed replies). WhatsApp honors configured ACP bindings.
- **No breaking changes or migration notes** were indicated in the provided release notes.
- **Related issues:** [#92679](https://github.com/openclaw/openclaw/issues/92679), [#931…](https://github.com/openclaw/openclaw/issues/931) (truncated in data).

## 3. Project Progress

In the last 24 hours, **90 pull requests were merged or closed**. Key advancements include:

- **macOS approvals migration fix** – [#93880](https://github.com/openclaw/openclaw/pull/93880) (closed)
- **Telegram /think levels for live-discovered Ollama models** – [#93882](https://github.com/openclaw/openclaw/pull/93882) (open, AI-assisted)
- **Hook payloads now include tool_use blocks** – [#93637](https://github.com/openclaw/openclaw/pull/93637) (auto-fix for bug #93381)
- **Gateway reserved target keywords rejected** – [#93761](https://github.com/openclaw/openclaw/pull/93761) (closed, prevents incidents from literal `"current"` targets)
- **Codex extension now routes context via turn-scoped instructions** – [#93762](https://github.com/openclaw/openclaw/pull/93762) (closed)
- **Neutral billing copy for subscription auth** – [#93763](https://github.com/openclaw/openclaw/pull/93763) (closed)
- **Doctor command skip false-positive warnings** – [#92731](https://github.com/openclaw/openclaw/pull/92731) (closed)
- **Persisted history text rendering in Control UI** – [#93841](https://github.com/openclaw/openclaw/pull/93841) (open)
- **Cron watchdog improvements** – [#93914](https://github.com/openclaw/openclaw/pull/93914) (treats setup phases as progress)
- **Memory pressure thresholds configurable** – [#93910](https://github.com/openclaw/openclaw/pull/93910) (open)

## 4. Community Hot Topics

| Issue / PR | Comments | 👍 | Summary |
|------------|----------|----|---------|
| [#75](https://github.com/openclaw/openclaw/issues/75) – [Linux/Windows Clawdbot Apps] | 109 | 79 | Long-running request for desktop apps beyond macOS/iOS/Android. |
| [#44925](https://github.com/openclaw/openclaw/issues/44925) – [Subagent completion silently lost] | 19 | 1 | Critical reliability bug: agent orchestration failures without retry or notification. |
| [#22676](https://github.com/openclaw/openclaw/issues/22676) – [Signal daemon race condition] | 17 | 0 | Race on SIGUSR1 restart leads to orphaned processes and send failures. |
| [#68596](https://github.com/openclaw/openclaw/issues/68596) – [Configurable streaming watchdog] | 14 | 8 | Request to extend watchdog timeout for models with extended reasoning. |
| [#62505](https://github.com/openclaw/openclaw/issues/62505) – [Coding Agent regression] | 14 | 1 | Agent used for coding stopped completing tasks after 2026.4.2. |

The community’s strongest demand centers on **platform support** (Linux/Windows apps), **subagent reliability**, and **configuration flexibility** (watchdog, private network access). The high reaction count on #75 (79 👍) underscores the importance of cross-platform availability.

## 5. Bugs & Stability

**Critical (P1) bugs reported or updated today:**

| Issue | Summary | Fix PR exists? |
|-------|---------|----------------|
| [#44925](https://github.com/openclaw/openclaw/issues/44925) | Subagent completion silently lost, no retry | linked [PR open](https://github.com/openclaw/openclaw/pull/??) |
| [#22676](https://github.com/openclaw/openclaw/issues/22676) | SIGUSR1 restart race → orphaned processes | linked PR open |
| [#62505](https://github.com/openclaw/openclaw/issues/62505) | Coding Agent never completes (regression) | linked PR open |
| [#48003](https://github.com/openclaw/openclaw/issues/48003) | Steer mode not injecting messages mid-turn | linked PR open |
| [#69118](https://github.com/openclaw/openclaw/issues/69118) | Group channel session reset on every turn | linked PR open |
| [#93375](https://github.com/openclaw/openclaw/issues/93375) | Telegram polling silent crash loop, health monitor can't recover | new, no fix PR yet |
| [#73148](https://github.com/openclaw/openclaw/issues/73148) | Image tool opaque "Failed to optimize image" without sharp | no linked PR |
| [#66443](https://github.com/openclaw/openclaw/issues/66443) | Overflow recovery duplicates user messages, amplifies transcript growth | linked PR open |
| [#39807](https://github.com/openclaw/openclaw/issues/39807) | Billing error infinite retry spiral (no backoff) | linked PR open |
| [#52130](https://github.com/openclaw/openclaw/issues/52130) | Telegram restart storm + misleading SecretRef | no linked PR |
| [#44749](https://github.com/openclaw/openclaw/issues/44749) | Concurrent allow-always lost allowlist entries | linked PR open |
| [#69943](https://github.com/openclaw/openclaw/issues/69943) | Session-memory hook persists raw tokens → poisoning loop | linked PR open |

**Security-impacting bugs:**
- [#39604](https://github.com/openclaw/openclaw/issues/39604) – Private network access still blocked; opt-in requested.
- [#48949](https://github.com/openclaw/openclaw/issues/48949) – Feishu fails with proxy configured.
- [#39847](https://github.com/openclaw/openclaw/issues/39847) – Echo contamination: internal metadata leaked to Discord.

**Stability improvements in progress:** Several of these have associated fix PRs, but many await maintainer review or product decisions.

## 6. Feature Requests & Roadmap Signals

| Issue | Feature | Likelihood for next release |
|-------|---------|----------------------------|
| [#68596](https://github.com/openclaw/openclaw/issues/68596) | Configurable streaming watchdog timeout | High – user demand, simple config extension |
| [#39604](https://github.com/openclaw/openclaw/issues/39604) | `tools.web.fetch.allowPrivateNetwork` | Medium – security review needed |
| [#63930](https://github.com/openclaw/openclaw/issues/63930) | Anthropic advisor tool support | Medium – beta server-side tool |
| [#66252](https://github.com/openclaw/openclaw/issues/66252) | Per-agent TTS/STT overrides | Medium – multi-language use case |
| [#67413](https://github.com/openclaw/openclaw/issues/67413) | Per-agent dreaming configuration | Low – requires memory architecture changes |
| [#54373](https://github.com/openclaw/openclaw/issues/54373) | Context provenance metadata | Low – RFC phase, complex |
| [#75](https://github.com/openclaw/openclaw/issues/75) | Linux/Windows desktop apps | Long-term roadmap |

The release v2026.6.8 already addressed channel delivery, which was a frequent pain point. Next likely additions are **configurable watchdog timeout** (quick win) and **private network access opt-in** (security boundary work).

## 7. User Feedback Summary

**Common pain points expressed in issues:**

- **Subagent orchestration fragility** – silent failures, lost completions, no retry logic.
- **Session state corruption** – duplicate messages, transcript lock-ups, drift in hash-based session resets.
- **Regression frequency** – multiple "worked before, now fails" reports (coding agent, DeepSeek V4 Flash, Feishu cards).
- **Platform gaps** – Linux/Windows users are left out of native app support.
- **Opaque errors** – `"Failed to optimize image"` without hint about missing `sharp`; `"Unavailable — Outside allowed folders"` despite correct config.
- **Configuration complexity** – many features require deep config changes (acp, auth profiles, memory pressure) not discoverable in UI.

**Positive signals:**
- Users actively engage with the project: 79 👍 on the Linux/Windows app request indicates a passionate community.
- The new release improved Telegram/WhatsApp rendering, which the community had been asking for.
- PRs with "AI-assisted" annotations show growing use of automation for fixes.

**Satisfaction note:** The high number of bug reports suggests users rely heavily on OpenClaw in production, but also that reliability is a critical blocker.

## 8. Backlog Watch

The following high-priority items have remained open or unaddressed for extended periods and need maintainer attention:

| Issue | Created | Status | Notes |
|-------|---------|--------|-------|
| [#75](https://github.com/openclaw/openclaw/issues/75) – Linux/Windows Apps | 2026-01-01 | Needs product decision, security review | 109 comments, 79 👍 |
| [#44925](https://github.com/openclaw/openclaw/issues/44925) – Subagent silent loss | 2026-03-13 | Needs maintainer review, linked PR open | P1, critical reliability |
| [#22676](https://github.com/openclaw/openclaw/issues/22676) – Signal daemon race | 2026-02-21 | Linked PR open | Also P1, impact on session state |
| [#68596](https://github.com/openclaw/openclaw/issues/68596) – Configurable watchdog | 2026-04-18 | Needs product decision | 8 👍, no movement on design |
| [#57901](https://github.com/openclaw/openclaw/issues/57901) – Safeguard compaction ignores config | 2026-03-30 | Linked PR open, needs product decision | P2 |
| [#69208](https://github.com/openclaw/openclaw/issues/69208) – Duplicate transcript umbrella | 2026-04-20 | Needs live repro, maintainer review | Multiple channels affected |
| [#43367](https://github.com/openclaw/openclaw/issues/43367) – Multi-agent orchestration unstable | 2026-03-11 | Needs live repro, security review | P1, also impacts auth |
| [#39807](https://github.com/openclaw/openclaw/issues/39807) – Billing error death spiral | 2026-03-08 | Needs product decision, linked PR | P1, no backoff |

**Notable PRs waiting on author or maintainer:**

- [#67820](https://github.com/openclaw/openclaw/pull/67820) – WhatsApp QR reuse (waiting on author)
- [#69822](https://github.com/openclaw/openclaw/pull/69822) – Socket drain / session evictions (waiting on author)
- [#90741](https://github.com/openclaw/openclaw/pull/90741) – Auth-profile fingerprint cache (ready for maintainer look)
- [#93580](https://github.com/openclaw/openclaw/pull/93580) – Cron delivery awareness (ready for maintainer look)
- [#67077](https://github.com/openclaw/openclaw/pull/67077) – Auth-profile saves non-fatal (waiting on author)

The backlog suggests the team should prioritize **critical reliability bugs** (#44925, #22676) and **user-requested configurability** (#68596) while triaging the wave of PRs that have passed initial review.

---

*Data snapshot: 2026-06-17. Sources: GitHub issues & PRs for [openclaw/openclaw](https://github.com/openclaw/openclaw).*

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: Personal AI Agent Open-Source Ecosystem

## 1. Ecosystem Overview

The personal AI agent open-source landscape is experiencing an intense period of competition and maturation, with seven projects showing high activity this week and two emerging as clear scaling leaders (OpenClaw, IronClaw). The ecosystem is bifurcating between "ecosystem reference" projects that aim to be full-stack platforms (OpenClaw, Hermes, ZeroClaw) and "specialized forks" optimizing for specific deployment constraints (NanoClaw, PicoClaw, TinyClaw). A strong shared pattern is emerging: users are treating these agents as production infrastructure, demanding reliability (subagent orchestration, session state persistence), channel parity (Telegram, WhatsApp, Slack, WeCom), and security hardening (credential handling, workspace isolation). The community is self-healing—many critical bugs are being fixed by contributors within 24-48 hours of being reported. The ecosystem is clearly past the "experiment" phase and into "production deployment" concerns, with significant pressure on maintainer bandwidth across the board.

---

## 2. Activity Comparison

| Project | Issues Updated (24h) | PRs Updated (24h) | Merged/Closed PRs (24h) | Release Status | Health Score |
|---------|----------------------|-------------------|------------------------|----------------|--------------|
| **OpenClaw** | 197 (9 closed) | 500 (90 merged) | 90 | v2026.6.8 ✅ | **Very Strong** — High activity but strained maintainer bandwidth (188 open issues) |
| **IronClaw** | 21 (6 closed) | 50 (14 merged) | 14 | No release | **Strong** — Active triage, responsive to bugs |
| **CoPaw** | 27 | 40 (25 merged) | 25 | v1.1.12-beta.1 ✅ | **Very Strong** — Rapid merging, beta release today |
| **ZeroClaw** | 2 | 50 (2 merged) | 2 | No release | **Active** — High PR volume but low merge rate; staging for v0.8.1 |
| **Hermes Agent** | 29 (1 closed) | 50 (3 merged) | 3 | No release | **Moderate** — High activity but low closure rate (47 open PRs) |
| **NanoBot** | 10 (5 closed) | 24 (14 merged) | 14 | No release | **Strong** — Healthy balance, good merging velocity |
| **PicoClaw** | 15 | 15 (12 merged) | 12 | Nightly ✅ | **Very Strong** — Excellent merge-to-open ratio, responsive security triage |
| **NanoClaw** | 6 (1 closed) | 5 (4 merged) | 4 | No release | **Good** — Targeted fixes, lean operation |
| **LobsterAI** | 2 (0 closed) | 5 (5 merged) | 5 | No release | **Moderate** — Small but efficient, maintenance phase |
| **NullClaw** | 2 | 3 (0 merged) | 0 | No release | **Low** — Low activity, no merged PRs |
| **Moltis** | 4 | 2 (0 merged) | 0 | No release | **Low** — Minimal engagement, no progress |
| **TinyClaw** | 0 | 1 (0 merged) | 0 | No release | **Dormant** — No activity; single PR awaiting review |
| **ZeptoClaw** | 0 | 1 (0 merged) | 0 | No release | **Dormant** — Only Dependabot activity |

**Community Size Proxy:** OpenClaw's 79 👍 on a single feature request (#75) versus IronClaw/CoPaw's 0-3 reactions on most issues reveals that OpenClaw commands the largest engaged community by an order of magnitude.

---

## 3. OpenClaw's Position

**Advantages:**
- **Scale leader:** 197 issues + 500 PRs in 24h is 4-10x the next most active project. Community of ~500+ active contributors vs. 20-50 for peers.
- **Release velocity:** v2026.6.8 shipped with Telegram/WhatsApp improvements. 90 PRs merged in one day—maintainers are clearly prioritizing throughput.
- **Ecosystem anchor:** Most forks (PicoClaw, NanoClaw) are explicitly derived from OpenClaw, making it the de facto reference architecture.
- **Channel breadth:** Telegram, WhatsApp, Discord, Signal, Feishu/Lark, and more—the widest chat surface coverage.

**Technical Approach Differences:**
- **Architecture:** OpenClaw uses a gateway + agent runtime + channel adapters model. This is similar to Hermes and ZeroClaw but more modular.
- **Memory/persistence:** Pluggable backends (Postgres, SQLite) vs. IronClaw's in-memory + NEAR cloud. OpenClaw's approach is more self-hosted friendly.
- **Tool ecosystem:** Codex extension, web/extract tools, cron, billing—richer out-of-box than NanoBot or NullClaw.

**Weaknesses vs. Peers:**
- **Maintainer bandwidth is clearly strained:** 188 open issues, 410 open PRs. Compare to PicoClaw (15/15) where most issues get closed within days. OpenClaw has a backlog danger.
- **Stability regressions are frequent:** Multiple "worked before, now fails" reports (Coding Agent, DeepSeek V4 Flash) erode trust.
- **No Linux/Windows desktop apps** (#75 has 79 👍)—PicoClaw and CoPaw are closing this gap.
- **Subagent orchestration is fragile** (#44925) — IronClaw and CoPaw are more mature in multi-agent reliability.

**Community Size Comparison (Indicators):**

| Metric | OpenClaw | IronClaw | CoPaw |
|--------|----------|----------|-------|
| Max 👍 on single issue | 79 | ~3 | ~5 |
| Comments on hottest issue | 109 | 3 | 14 |
| Unique contributors (est. daily) | 50-100 | 10-15 | 10-15 |
| PR merge turnaround | 1-7 days | 1-3 days | 12-24h |

**Verdict:** OpenClaw is the ecosystem's central hub and will remain so due to network effects. However, it is losing ground in reliability and developer experience to smaller, more focused projects like PicoClaw and CoPaw. For mission-critical deployments, consider derivatives.

---

## 4. Shared Technical Focus Areas

Requirements emerging independently across multiple projects:

| Need | Projects Reporting | Specific Pain Points |
|------|-------------------|---------------------|
| **Subagent orchestration reliability** | OpenClaw, Hermes, IronClaw, CoPaw | Silent completion loss, no retry logic, process freezes during compaction |
| **Telegram/WhatsApp rich content** | OpenClaw, Hermes, PicoClaw, ZeroClaw | Bullet lists broken, tables not rendered, photo caching failures |
| **Linux/Windows native apps** | OpenClaw (#75), PicoClaw, TinyClaw | Desktop app gap is the #1 community request across all projects |
| **Configuration flexibility** | All projects | Watchdog timeouts, private network access, per-agent TTS/STT, streaming config |
| **Backward compatibility/regression prevention** | OpenClaw, NanoBot, CoPaw | "Worked in vX.Y, broken in vX.Y+1" pattern erodes trust |
| **Opaque error messages** | OpenClaw, Hermes, NanoBot | "Failed to optimize image" with no hint, "Unavailable" with no context |
| **Workspace security / isolation** | NanoBot, PicoClaw, ZeroClaw, CoPaw | Git commands blocked, stale file writes to wrong paths, credential exfiltration |
| **Serverless/immutable deployment** | NanoClaw, NullClaw | Need env-var-based config injection, upgrade tripwire opt-out |
| **Scheduled/background automation** | LobsterAI, IronClaw, NullClaw, ZeroClaw | Cron tasks interrupting chat, silent failures in scheduled tasks |
| **Multi-tenant / multi-user support** | Hermes, IronClaw, OpenClaw | Shared core with isolated agents, per-user profiles, managed fleets |

The **subagent reliability** and **Telegram rendering** clusters are the most critical—they appear in 4+ projects each and have high user impact.

---

## 5. Differentiation Analysis

| Dimension | OpenClaw / Hermes / ZeroClaw | PicoClaw / NanoClaw | IronClaw / CoPaw | TinyClaw / ZeptoClaw |
|-----------|------------------------------|---------------------|------------------|----------------------|
| **Target User** | Power users, developers, self-hosters | Embedded/edge, minimal-footprint | Enterprise teams, cloud-first | Learners, proof-of-concept |
| **Architecture Philosophy** | Full platform, extensible via plugins | Single-binary, minimal dependencies | Cloud-hosted, managed state | Ultra-light, educational |
| **Memory/Persistence** | Configurable (SQL/Postgres) | SQLite, file-based | In-memory + cloud DB | File only |
| **Chat Channels** | All major (Telegram, Discord, Signal, etc.) | Subset (Telegram, Discord) | Slack, Reborn UI | CLI-only |
| **Multi-Agent** | Yes (subagent orchestration) | Limited | Yes (engine V2) | No |
| **Security Posture** | Mixed: many open CVEs, credential issues | Strong: rapid patch turnaround | Strong: SSO, per-user profiles | Minimal |
| **Release Cadence** | Weekly (major), daily (patch) | Nightly, bi-weekly stable | Monthly | No release in weeks |
| **Community Governance** | BDFL-style (openclaw core team) | Maintainer-driven | NEAR Foundation-backed | Solo maintainer |
| **Key Weakness** | Backlog overload, regression frequency | Channel breadth limited | Cloud dependency, lesser community | Stagnation, no adoption |

**Niche Strategies:**
- **NanoBot** leans heavily into WebUI (automation management view merged today) and Dream memory—differentiating on UX rather than channel breadth.
- **Moltis** and **LobsterAI** are targeting specific verticals (voice-first, cowork/artifacts) rather than competing as general agents.
- **NullClaw** is differentiate on MS Teams/cron focus but has minimal momentum.

---

## 6. Community Momentum & Maturity

**Tier 1: Hyper-accelerating (daily releases, 50+ PRs/day)**
- **OpenClaw** — Despite backlog, the sheer volume proves it's the center of gravity. Releasing weekly with 90 PRs/day merging.
- **CoPaw** — Shipping betas, 25 PRs merged in 24h. Community-driven patches (first-time contributor fixed a critical freeze). Rapidly maturing.

**Tier 2: Rapid iteration (10-50 PRs/day, clear release cycles)**
- **IronClaw** — Engine V2 Milestone 0 closed. Focused on core agent quality over feature breadth.
- **PicoClaw** — 12 PRs merged/closed today. Excellent merge-to-open ratio. Security advisories being addressed.
- **ZeroClaw** — High PR volume (50) but low merge rate. Staging for v0.8.1. Needs to convert backlog into releases.
- **NanoBot** — 14 PRs merged, healthy balance. WebUI and Dream memory are maturing.

**Tier 3: Maintenance mode (1-5 PRs/day, no releases)**
- **NanoClaw** — Focused, but low volume. Budget-billing fix was important. Managed-fleet features requested.
- **Hermes Agent** — High issue count (28 open) but low closure rate. Needs maintainer bandwidth improvement.
- **LobsterAI** — UI polish only. Scheduled tasks bug (#1424) is 2 months stale—concerning for project health.

**Tier 4: Stalled (0-1 PRs/day, no maintainer response)**
- **NullClaw**, **Moltis**, **TinyClaw**, **ZeptoClaw** — Minimal activity. TinyClaw's single Windows PR has been open for 24h with no review. ZeptoClaw is effectively a dependency-update-only project.

**Maturity Assessment:** OpenClaw is the most mature in feature breadth but immature in stability. CoPaw shows the highest maturity growth rate (releasing betas, closing critical bugs within days). PicoClaw is the most operationally mature for embedded use cases. The Tier 3-4 projects are at risk of abandonment unless maintainers re-engage.

---

## 7. Trend Signals

**1. Reliability over features — the "Production Reality" shift**
Multiple projects report users treating agents as production infrastructure (OpenClaw's #44925 subagent loss, CoPaw's #5218 process freeze, Hermes's #6841 tool-call corruption). The community is voting with their bug reports: they want agents that don't silently fail. Expect a wave of "reliability engineering" features: retry logic, timeout configs, health monitors, and observability hooks.

**2. Enterprise multi-tenancy is under-served**
Hermes (#34352) has a detailed proposal for isolated agent instances. IronClaw (#5008) is building per-user profiles. ZeroClaw has per-agent workspace security. The demand for multi-user, multi-agent environments is rising fast—expect this to be a major differentiator in 2026 H2.

**3. Channel parity is table stakes, not a differentiator**
Every major channel now has at least one bug report (Telegram bullets, Slack @handles, WhatsApp QR, WeCom media). Users expect all channels to work equally well. Projects that neglect "unsexy" channel polish will lose users to those that do. OpenClaw's v2026.6.8 (richer Telegram rendering) is a direct response.

**4. Configuration explosion is unsustainable**
Users are requesting configurable watchdog timeouts (OpenClaw #68596), per-agent TTS/STT (#66252), private network opt-in (#39604), streaming flags (PicoClaw #2404), and env-var overrides (NanoClaw #2781). The trend is clear: **"one-size-fits-all" config models are failing**. Expect projects to converge on layered config (global → agent → per-turn) or YAML-based DSLs.

**5. The "Ollama/local model" user base is growing but underserved**
NullClaw (#952 — incomplete local model answers), Moltis (#1128 — whisper.cpp transcription errors), CoPaw (#5233 — Ollama model selection missing). Users running local models face unique bugs that cloud-first projects don't test for. This is an underserved niche with high loyalty potential.

**6. Scheduled tasks/cron is the new "killer app"**
IronClaw, NullClaw, CoPaw, LobsterAI, and ZeroClaw all have active cron/scheduler work. Users want agents to run autonomously, not just respond to chat. The need for background execution (CoPaw #5250) and silent cron jobs (CoPaw #5251, merged today) signals that scheduled, non-interactive agent operations will become a core expectation.

**7. Security is becoming a differentiator — fast**
PicoClaw received 9 security advisories in one batch. CoPaw released v1.1.12-beta.1 specifically for keychain isolation. ZeroClaw (#7284) fixed per-agent workspace directory enforcement. OpenClaw has open security-impact bugs (#39604, #48949). Projects that can demonstrate rapid security patch turnaround (PicoClaw, CoPaw) will gain trust over those with unaddressed CVEs.

**Value for AI Agent Developers:**
- If building for production: Watch OpenClaw for community size, but build on PicoClaw or CoPaw for stability.
- If targeting local/first users: Prioritize same-model bugs across platforms.
- If targeting enterprise: Double down on multi-tenancy, SSO, and per-user profiles — the demand is outpacing supply.
- If building a new project: Choose a niche (voice, embedded, security-hardened) rather than trying to match OpenClaw's breadth.

---

*Data snapshot: 2026-06-17. Sources: GitHub activity for 13 projects in the personal AI agent ecosystem.*

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

## NanoBot Project Digest — 2026-06-17

### 1. Today’s Overview
The project saw elevated activity with **10 issues** and **24 pull requests** updated in the last 24 hours. Of those, 5 issues were closed and 14 PRs were merged/closed, indicating a healthy balance of bug fixing, feature integration, and community contributions. No new release was published today. The most notable developments include fixes for installer breakage on Debian 13, a major rework of stream idle timeout validation, and the merge of a WebUI automation management view (#4330). Community engagement remains strong, with one installer bug (#4360) attracting 9 comments.

### 2. Releases
*No new releases were published on this date.*

### 3. Project Progress — Merged/Closed PRs (14 total)
The following high-impact pull requests were merged or closed today:

- **#4330** — *[enhancement] feat(webui): add automation management view* (merged — adds a full UI for filtering, running, pausing, and editing automations)  
- **#4352** — *[valid] fix(context): cap recent-history digest by tokens, not characters* (merged — prevents overflow for CJK/text-heavy histories)  
- **#4355** — *[valid] chore: ignore bridge/node_modules* (merged — tidy gitignore)  
- **#4358** — *fix(api): avoid duplicate user turn on empty-response retry* (merged — closes #4079)  
- **#4361** — *fix(providers): enable thinking for Kimi K2.7 models* (merged)  
- **#4363** — *[valid] fix(providers): validate stream idle timeout config* (merged — closes #4065)  
- **#4365** — *[documentation] docs: use pipe pattern for curl installer commands* (merged — safer Docker/script installations)  
- **#4368** — *[valid] Fix macOS installer for externally managed Python* (merged)  
- **#4369** — *[valid] Explain empty Dream runs* (merged — better user feedback instead of opaque error)  
- **#4370** — *[valid] Enable idle auto-compact by default* (merged — default `15` minutes; explicit `0` to opt out)  
- **#4379** — *nanobot main* (closed — likely a duplicate or test PR)

These merges advance both usability (automation UI, clearer Dream diagnostics) and stability (timeout validation, retry logic, token-aware context limits).

### 4. Community Hot Topics
The most active discussion this period centered on **Issue #4360** ([HKUDS/nanobot Issue #4360](https://github.com/HKUDS/nanobot/issues/4360)) — *"[bug] 'end of file unexpected' during installer"* — with 9 comments. The issue reports that the one-command installer fails inside a fresh Debian 13 Docker container due to a shell syntax error. The reporter suspects the failure is related to an earlier message about `/root/...`. The high engagement reflects the critical nature of installation reliability for new users. A separate documentation PR (#4365) was merged to change the curl installer pattern, which may mitigate similar issues.

### 5. Bugs & Stability
Several bugs were reported or fixed today, ranked by potential impact:

| # | Issue | Severity | Fix PR (if any) |
|---|-------|----------|----------------|
| #4360 | Installer crashes on Debian 13 (`end of file unexpected`) | **High** – prevents fresh setup | #4365 (doc fix) addresses the pattern but further fix may be needed |
| #4375 | Git commands blocked by workspace security policy ([#4375](https://github.com/HKUDS/nanobot/issues/4375)) | **High** – blocks version control inside workspace | No fix PR yet |
| #4079 | API empty-response retry duplicates user turns (closed) | **Medium** – message log corruption | Fixed by #4358 (merged) |
| #4242 | `dream.enabled: false` still injects recent history ([#4242](https://github.com/HKUDS/nanobot/issues/4242)) | **Medium** – config ignored for system prompt injection | No fix PR yet |
| #4065 | Invalid `NANOBOT_STREAM_IDLE_TIMEOUT_S` crashes streaming (closed) | **Medium** – crash on misconfiguration | Fixed by #4363 (merged) |
| #4374 | Project workspaces: SOUL.md/USER.md read from project but written to default workspace ([#4374](https://github.com/HKUDS/nanobot/issues/4374)) | **Medium** – data integrity asymmetry | No fix PR yet |
| #4366 | Local model servers bypass proxy (closed) | **Low** – connectivity issue | Addressed by open PR #4367 |

Additionally, open PRs target other stability concerns:  
- **#4367** — disable proxy for local endpoints, respect env for cloud  
- **#4371** — add breakpoint before Recent History to stabilize system prefix caching  
- **#4372** — fix MCP malformed progress notifications  
- **#4373** — preserve delivery context during memory consolidation  

### 6. Feature Requests & Roadmap Signals
Two feature requests were opened today:

- **#4376** — *"user friendly wizard"* ([HKUDS/nanobot Issue #4376](https://github.com/HKUDS/nanobot/issues/4376)) proposes a guided onboarding wizard that reduces technical barriers for non-expert users. Given the team’s recent focus on WebUI (#4330) and installer improvements (#4365, #4368), this is likely to be prioritized for the next release.
- **#4378** — *"cron level model/preset"* ([HKUDS/nanobot Issue #4378](https://github.com/HKUDS/nanobot/issues/4378)) requests the ability to switch models or presets via cron. This is a follow-up to a earlier discussion (#4292). Workarounds using cron scripts are possible but a native feature may appear in a future minor release.

A notable feature PR still open:  
- **#4350** — *feat(web): add Keenable search provider* (new built-in web search service) — likely to land in the next minor version.

### 7. User Feedback Summary
Real user pain points surfaced in today’s issues:

- **Installer reliability**: The `end of file unexpected` error on Debian 13 (The‑Markitecht) highlights that the installer shell pattern is fragile in certain environments. The merged PR #4365 may help, but underlying syntax issues remain.
- **Configuration not honoured**: skyline75489 reported that disabling Dream still injects all chat history into the system prompt because the Dream cursor never advances. The merged PR #4369 adds clearer messaging but does not fix the core logic.
- **Workspace security false positive**: jjmanrique found that Git commands (add, commit, push) inside subdirectories are blocked by the workspace security guard, even when those directories are within the allowed root. This is a usability blocker for agent workflows using version control.
- **Data write inconsistency**: maximilize discovered that project workspace files (SOUL.md, USER.md) are read from the per-turn project path but *written* to the default workspace, causing asymmetry. This could lead to lost agent state.

Positive signals: Contributors continue to submit high-quality fixes and features (e.g., Yu‑xin‑c’s delivery context preservation in PR #4373, michaelxer’s proxy fix in #4367), and maintainers are merging rapidly.

### 8. Backlog Watch
The following long-standing issues and PRs remain unanswered or require maintainer attention:

- **Issue #4242** (open since 2026-06-08, 1 comment) — Dream injection bug: no fix PR yet, despite being open for over a week. Deserves prioritisation as it affects a major feature (Dream memory).
- **PR #3662** (open since 2026-05-06) — *fix(tokens): avoid network loads during estimation*: aims to improve offline resilience but has not been updated today. Needs review/merge decision.
- **PR #4053** (open since 2026-05-29) — *fix(tools): keep read-only roots out of write paths*: addresses a security boundary issue but is still open. Given the new workspace security bug (#4375), this PR may be a key part of the solution.
- **Issue #4375** (new today) — Git command blocking: no response from maintainers yet. Should be triaged promptly due to its severity.

All links: [HKUDS/nanobot](https://github.com/HKUDS/nanobot)

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-06-17

## 1. Today's Overview
Project activity remains very high, with **29 issues** and **50 pull requests** updated in the last 24 hours. The open issue count (28) and open PR count (47) indicate an active maintenance cycle, though the number of closed items is low (1 issue, 3 PRs). No new releases were published today. The issue tracker reveals a broad range of bugs across platforms (Telegram, Desktop, CLI, email, Signal) and core components (gateway, agent, credential pool, MCP), alongside several feature suggestions that signal growing enterprise and multi-user adoption. The high PR volume suggests the maintainer team is actively addressing reported problems.

## 2. Releases
**No new releases today.**

## 3. Project Progress
**3 pull requests were merged/closed today** (data available for one):
- **[PR #28981 — fix: exclude .stash directory from skill scanning](https://github.com/NousResearch/hermes-agent/pull/28981)** (closed) – Prevents the skill scanner from picking up working data from the Stash skill-sync tool, reducing noise and potential conflicts.

**1 issue was closed:**
- **[Issue #46789 — Desktop app process execution segfaults on macOS](https://github.com/NousResearch/hermes-agent/issues/46789)** (closed) – The bug has apparently been resolved or superseded; details are not yet public.

Other notable open PRs advancing bug fixes today (not yet merged) include fixes for MCP discovery logging, credential pool retry parsing, and Telegram bullet display – see section 5.

## 4. Community Hot Topics
The most active discussions (by comment count) reveal strong user interest in **multi-tenancy** and **cross-platform display quality**:

- **[Issue #34352 – Solving the Multi-Tenant Hermes Problem](https://github.com/NousResearch/hermes-agent/issues/34352)** (8 comments, 0 👍) – A detailed proposal from NimbleCoAI to enable isolated agent instances sharing the same core, bypassing hook system limitations. This is the most-commented issue today and reflects enterprise demand for concurrent, isolated agent runs.

- **[Issue #6388 – Telegram MarkdownV2 escape breaks bullet list display](https://github.com/NousResearch/hermes-agent/issues/6388)** (6 comments, 1 👍) – Users are frustrated that `-` becomes `\-` in Telegram, breaking list rendering. The root cause is in the gateway’s escape function, and the thread shows multiple workarounds are needed.

- **[Issue #6841 – Hermes tool-calling pipeline corrupts tool names and JSON arguments](https://github.com/NousResearch/hermes-agent/issues/6841)** (2 comments, 3 👍) – A P1 bug from 2026-04-09 with high reaction count. Intermittent malformed payloads cause generic tool-call failures. This issue remains open and is a top candidate for community attention.

- **[Issue #47042 – Desktop model picker hides custom providers](https://github.com/NousResearch/hermes-agent/issues/47042)** (1 comment, 2 👍) – A regression in the model picker’s deduplication logic causing custom providers to disappear.

## 5. Bugs & Stability
Today’s newly reported or updated bugs span multiple components, with severity P2 (moderate) and P3 (minor). No P0 critical issues were reported. Many have corresponding fix PRs already opened.

### Newly filed issues (2026-06-17) with severity P2:

| Issue | Title | Component | Fix PR (if exists) |
|-------|-------|-----------|--------------------|
| [#47609](https://github.com/NousResearch/hermes-agent/issues/47609) | Desktop crashes when sending attached image with vision model | Desktop | None yet |
| [#47605](https://github.com/NousResearch/hermes-agent/issues/47605) | Sandbox PATH shadowing & HOME resolution | cron/config | None |
| [#47599](https://github.com/NousResearch/hermes-agent/issues/47599) | `web_extract` takes minutes to finish | web tool | None |
| [#47595](https://github.com/NousResearch/hermes-agent/issues/47595) | Gateway makes periodic 30K-token API calls (wastes ~$5/mo) | gateway | None |
| [#47515](https://github.com/NousResearch/hermes-agent/issues/47515) | `hermes config set` coerces `off`/`on` to booleans | CLI | [PR #47607](https://github.com/NousResearch/hermes-agent/pull/47607) |
| [#47509](https://github.com/NousResearch/hermes-agent/issues/47509) | MCP discovery failures logged at DEBUG, invisible | gateway/MCP | [PR #47602](https://github.com/NousResearch/hermes-agent/pull/47602) |
| [#47510](https://github.com/NousResearch/hermes-agent/issues/47510) | MCP stdio subprocesses accumulate as zombies on restart | gateway/MCP | None yet |

### Other notable bugs updated today (P2/P3):
- **[#47048](https://github.com/NousResearch/hermes-agent/issues/47048)** – Telegram rich-message final reply overlaps with legacy MarkdownV2 (double render of tables + bullets).
- **[#47093](https://github.com/NousResearch/hermes-agent/issues/47093)** – Telegram photos dropped silently when `get_file()` times out before caching.
- **[#46866](https://github.com/NousResearch/hermes-agent/issues/46866)** – Signal gateway misroutes approval responses as steered mid-turn messages.
- **[#46891](https://github.com/NousResearch/hermes-agent/issues/46891)** – Credential pool retry-delay parser doesn’t handle absolute-datetime rate-limit messages (fix in [PR #47597](https://github.com/NousResearch/hermes-agent/pull/47597) context?).
- **[#46856](https://github.com/NousResearch/hermes-agent/issues/46856)** – OpenRouter generic errors not triggering rate-limit cooldown, causing fallback resets every turn.

## 6. Feature Requests & Roadmap Signals
Several feature requests indicate the community is pushing Hermes toward **multi-user, multi-platform, and self-service deployments**:

- **[#34352](https://github.com/NousResearch/hermes-agent/issues/34352) – Multi-tenant Hermes** (comment count highest). Likely a priority for the next major release given the detailed production-ready solution.
- **[#45779](https://github.com/NousResearch/hermes-agent/issues/45779) – Multi-gateway connections with per-gateway tabs in Desktop** (1 👍). Would allow managing multiple remote agents from one UI.
- **[#47608](https://github.com/NousResearch/hermes-agent/issues/47608) – Appservice for Matrix** (P3) – User wants an application service instead of password-based auth for Matrix bridges, hinting at larger institutional deployments.
- **[#47601](https://github.com/NousResearch/hermes-agent/issues/47601) – Dashboard broken behind reverse proxy** (CORS + systemd env) – Shows demand for production hosting behind proxies.
- **[#47598](https://github.com/NousResearch/hermes-agent/pull/47598) – Per-profile git credential provisioning** (PR, not issue). A boot-time feature for managing multiple GitHub tokens per profile, useful for multi-tenant scenarios.

**Prediction for next minor release (v0.16.x):** The multi-tenant proposal (if refined and merged) could land in the next minor version, along with the MCP visibility fixes, Telegram display corrections, and the credential pool absolute-time parser fix.

## 7. User Feedback Summary
Real user pain points from today’s issue reports:

- **Telegram display is broken** – Bullet lists, tables, and photo caching all have bugs. Users want reliable rich message rendering.
- **Desktop app crashes** – Image attachment with vision model causes crash (100% reproducible). This directly impacts user experience for image-described queries.
- **Configuration friction** – `hermes config set` silently corrupts enum values; custom providers disappear from model picker; `hermes status` shows false negatives for authenticated providers.
- **Invisible failures** – MCP misconfigurations, credential pool timeout parsing, and gateway API inefficiencies (wasteful periodic calls) are all silent, making debugging difficult.
- **Email subject missing** – Agents cannot set custom subjects on outbound emails, limiting integration with email workflows.
- **Slow web extraction** – `web_extract` taking minutes hurts interactivity; users expect sub-second tool responses.

Satisfaction signals: The high number of open PRs (47) suggests contributors are actively engaged and fixes are in the pipeline. The multi-tenant proposal received detailed community engagement.

## 8. Backlog Watch
The following important issues have been open for an extended period with no clear resolution:

- **[Issue #6841](https://github.com/NousResearch/hermes-agent/issues/6841)** – P1 tool-calling corruption (since April 9, 2026). 3 reactions, 2 comments. This is the highest-severity item on the backlog and affects all users. No PR linked.
- **[Issue #6388](https://github.com/NousResearch/hermes-agent/issues/6388)** – Telegram MarkdownV2 list rendering (since April 9, 2026). 1 reaction, 6 comments. Multiple users confirm the issue. No fix PR yet.
- **[Issue #34352](https://github.com/NousResearch/hermes-agent/issues/34352)** – Multi-tenant problem (since May 29, 2026). High engagement but no official response from maintainers? Check if there's a label. The issue is open and needs maintainer feedback to move forward.

All three issues would benefit from a maintainer’s triage or acknowledgement to guide community contributions.

---

*Generated from GitHub data updated as of 2026-06-17T23:59:59Z.*

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest
**Date:** 2026-06-17  
**Data Span:** Last 24 hours (up to 2026-06-17)

---

## 1. Today’s Overview

PicoClaw shows **very high activity** with 15 issues and 15 PRs updated in the last 24 hours, plus a new nightly release. Activity is driven by a major security advisory batch (9 issues filed on June 9, all updated today) and a rapid resolution cycle: 12 PRs were merged or closed, including several important bug fixes and features. The project’s responsiveness to both security reports and community-reported bugs remains strong. One feature request (streaming HTTP support) continues to attract discussion, while the newly published **nightly v0.3.0‑nightly build** provides early access to the latest changes.

---

## 2. Releases

- **nightly** (v0.3.0-nightly.20260617.a16a1e15)  
  An automated nightly build from `main`. Use with caution as it may be unstable.  
  [Full Changelog](https://github.com/sipeed/picoclaw/compare/v0.3.0...main)  

  No accompanying migration notes or documented breaking changes.

---

## 3. Project Progress (Merged/Closed PRs)

In the last 24 hours **12 PRs were merged or closed**, reflecting a healthy development cadence. Key changes include:

- **Feature – Remote cron commands** (#3137): Adds `tools.cron.command_allowed_remotes` config to control which channels can trigger cron commands – important for multi‑channel deployments.
- **Feature – Out‑of‑tree channel registration** (#3120): Introduces `RegisterChannelSettings` hook to let third‑party modules register channel config without forking PicoClaw.
- **Bug fix – Telegram forum topics** (#3135): Ensures replies go to the correct topic thread by using `compositeChatID` in `InboundContext.ChatID`.
- **Stability – Panic recovery** (#3132): Adds defer‑recover guards on critical‑path goroutines to prevent process crashes.
- **Bug fix – Session history** (#2990): Now reads full conversation history instead of showing only the last user message.
- **Bug fix – Context compression config** (#2988): `summarize_token_percent` setting is now respected; the `/context` command shows correct compression threshold.
- **Bug fix – Tool calls filtered during streaming** (#2987): Tool calls are no longer dropped when a stream is active.
- **Bug fix – Empty LLM responses retried** (#2983): Retries when OpenAI‑compatible providers return semantically empty assistant messages.
- **Minor fixes** – Explicitly ignore `Close()` errors on directory file descriptors (#3127), handle `json.Marshal` errors in Seahorse tools (#3130), and explicitly ignore file close errors in TTS (#3129).

A few older stale PRs (#2990, #2988, #2987, #2983) were also merged today, showing maintainers clearing the backlog.

Still open PRs (3):
- **#3116** – Complete `turn.done` lifecycle signaling (fixes issue #2984)
- **#3115** – Fix inline data URL media extraction for generic tool output
- **#3136** – Set both camelCase and snake_case `thought_signature` in Gemini tool call request body

---

## 4. Community Hot Topics

| Title | Type | Engagement | Link |
|-------|------|------------|------|
| **[Feature] Add in config to send streaming HTTP request** (#2404) | Issue (open) | 12 comments, 1 👍 | [View](https://github.com/sipeed/picoclaw/issues/2404) |
| Security advisories (10 issues, all by @YLChen-007) | Issues (open, stale) | 1 comment each | [Full list](#3070–3082) |

**Analysis:**  
The streaming HTTP request feature (#2404) has been open since April and remains the most discussed single issue. It reflects a core need: users want to pipe streaming responses from LLM backends to avoid timeouts and enable real‑time UX. The solution (adding a `"streaming": true` config key) is straightforward, and the PR pipeline shows the project is capable of such config changes – likely to land in a future stable release.

The burst of security issues (9 filed on June 9, updated today) have not received public resolutions yet. They cover SSRF bypasses, command whitelist escapes, and channel policy bypasses. The community and maintainers are aware; these issues are flagged `stale` but are still awaiting triage. They may represent a coordinated security audit disclosure.

---

## 5. Bugs & Stability

**High Severity**  
- **Security advisories** (see Section 8 – Backlog Watch) – Many allow privilege escalation, SSRF, or unauthorized access. No fix PRs visible yet, but the project has been responsive to prior reports.  
- **`su -c 'echo OK'` crashes the agent** (#3134) – `agent gateway` returns “No daemon is currently running!” and the agent crashes when asked to execute this command. Severity is high because it breaks a common Linux command path. No fix PR identified yet. Closed as duplicate? (It was closed shortly after filing.)

**Medium Severity**  
- **Telegram forum topic misrouting** (#3110) – Bot sends to `#General` instead of the correct topic thread. Already fixed by PR #3135 (merged).  
- **Empty LLM response never retried** – Fixed by PR #2983 (merged).  
- **Session history shows only last message** – Fixed by PR #2990 (merged).  
- **Context compression percent ignored** – Fixed by PR #2988 (merged).  

**Low Severity**  
- **Implicit error discards** in file operations – Fixed by PRs #3127, #3129, #3130.  
- **Unprotected goroutine panics** – Fixed by PR #3132.  

**Regressions:** None reported in the last 24 hours.

---

## 6. Feature Requests & Roadmap Signals

| Request | Comments | Likely Roadmap |
|---------|----------|----------------|
| Config‑driven streaming HTTP requests (#2404) | 12 comments, open since April | High – fits the project’s extensibility pattern; likely in next stable release. |
| Remote cron commands (#3137) | Already merged | Available in nightly. |
| Out‑of‑tree channels hook (#3120) | Already merged | Enables third‑party channel ecosystem. |

Predictions for next stable release (v0.3.0+):  
- **Streaming config** (#2404) is the most‑asked feature and implementation is trivial (boolean in config).  
- **`turn.done` lifecycle completion** (#3116) is an open PR that directly fixes a known issue.  
- **Gemini provider improvements** (#3136) – support for both camelCase and snake_case keys, needed for Gemini 3.5 Flash Agentic reasoning.  

No major roadmap document was updated in the observed data.

---

## 7. User Feedback Summary

**Common pain points (resolved or in progress):**  
- *Telegram forum topic replies* – users want correct threading; fix merged.  
- *Session history truncated* – users could only see last message in Web UI; fix merged.  
- *Context compression not honoring config* – users confused by `76800` always shown; fix merged.  
- *Tool calls dropped during streaming* – broke multi‑step LLM interactions; fix merged.  
- *Empty LLM responses not retried* – silent failures; fix merged.  

**Unresolved pain points:**  
- `su -c` command breaks agent (issue #3134, closed without fix – needs verification).  
- Streaming HTTP to LLM backends not possible (feature request #2404 – actively wanted).  

**User satisfaction:**  
High – the rapid merging of fixes (12 PRs in one day) indicates a well‑maintained project responsive to community feedback. The security disclosure reaction (even if not yet fixed) demonstrates transparency.

---

## 8. Backlog Watch

| Item | Age | Status | Action Needed |
|------|-----|--------|---------------|
| **#2404 – Streaming config** (enhancement) | 71 days open | Still open, 12 comments | Needs maintainer decision – should be merged given popularity. |
| **#3082, #3081, #3079, #3078, #3076, #3075, #3074, #3073, #3072, #3071, #3070, #3068 – Security issues** | 8 days open, marked `stale` | All filed by same reporter, no comments from maintainers | High priority – each describes a concrete bypass. Should be triaged and either fixed or acknowledged with mitigation plans. |
| **#3110 – Telegram forum topic** | 5 days open, closed after PR merged | Resolved by #3135 | No further action. |
| **#2984 – `turn.done` lifecycle** | ~15 days old | PR #3116 still open | Needs review and test. |

**Notable:** The security advisory pipeline is concerning – 12 open security issues with zero maintainer responses. Even if they are being handled privately, the public issues should be updated to reflect status. The `stale` label suggests they may have been auto‑marked, which contradicts the urgency of security bugs.

---

*Generated from GitHub activity up to 2026-06-17 23:59 UTC.*

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-17

## 1. Today's Overview

The project saw moderate activity over the past 24 hours, with **6 issues updated** (5 open, 1 closed) and **5 pull requests updated** (1 open, 4 merged/closed). A critical user-facing bug—silent dropping of budget-exhausted LLM turns—was fixed and merged, improving reliability. Community discussion continues around credential proxy compliance risks and Slack URL handling, while two new feature requests signal growing needs for managed-fleet deployments and external credential injection. No new releases were published.

## 2. Releases

*No new releases today.*

## 3. Project Progress

Four PRs were merged/closed today, advancing both reliability and functionality:

- **PR #2069** ([link](nanocoai/nanoclaw PR #2069)) — Merged the **webchat skill v1**, a new feature skill that adds a chat interface channel or integration (closed).
- **PR #2782** ([link](nanocoai/nanoclaw PR #2782)) — Fixed the **tailscale-docker routing service** to self-heal by switching from a `Type=oneshot` unit to a sustained mechanism that reapplies ip rules after Tailscale flushes them (e.g., exit-node reconnect). This addresses a silent connectivity loss in Tailscale-based deployments.
- **PR #2759** ([link](nanocoai/nanoclaw PR #2759)) — Fixed the **agent-runner budget/billing error handling** (closes #2751). LLM turns that hit token or spend limits are now delivered as error messages to the user instead of being silently dropped.
- **PR #2775** ([link](nanocoai/nanoclaw PR #2775)) — Clarified the **OneCLI gateway upgrade documentation**, noting that the `[BREAKING]` notice about version pinning applies only to fresh installs, eliminating a false sense of automated upgrade.

One PR remains open:
- **PR #2780** ([link](nanocoai/nanoclaw PR #2780)) — Adds an environment variable (`NANOCLAW_DISABLE_UPGRADE_TRIPWIRE`) to opt out of the startup upgrade tripwire, enabling deployments with immutable base images.

## 4. Community Hot Topics

The most active discussions (by comment count) are:

- **Issue #1669** ([link](nanocoai/nanoclaw Issue #1669)) — *Does Credential Proxy implementation risk Anthropic account bans?*  
  Author: LCJD99 | Updated: 2026-06-16 | 1 comment  
  Underlying concern: Compliance with Anthropic’s TOS regarding OAuth reverse-proxies. The project’s credential proxy might trigger anti-fraud checks. No maintainer resolution yet.

- **Issue #2779** ([link](nanocoai/nanoclaw Issue #2779)) — *Slack: @handles inside URLs get mangled into broken mentions*  
  Author: GitOnion | Updated: 2026-06-16 | 1 comment  
  User reports that Slack message formatting corrupts URLs containing `@handle` path segments (e.g., HackMD, Mastodon). This is a user experience regression for Slack channel integration.

Both issues lack multiple comments but represent significant functional and trust concerns.

## 5. Bugs & Stability

Reported bugs and their status (ranked by severity):

| Severity | Issue | Description | Fix PR | Status |
|----------|-------|-------------|--------|--------|
| **High** | #2779 | Slack URL mangling breaks links with `@handle` in path. Agent-generated links become unusable. | None open | Open |
| **Medium** | #2784 | Container-runner staleness check only watches `index.ts`, missing changes to `ipc-mcp-stdio.ts`. Leads to stale code in session directories. | None open | Open |
| **Medium** | #2751 (closed) | Budget-exhausted LLM turns silently dropped. **Fixed** by PR #2759 (merged today). | #2759 | Closed |
| **Low** | #2783 | `docs/SECURITY.md` describes retired v1 trust model and references a non-existent skill, causing documentation drift. | None open | Open |

No crashes or regressions reported. The most impactful fix made today resolves the silent-drop issue (#2751).

## 6. Feature Requests & Roadmap Signals

Two new feature requests were filed today:

- **Issue #2781** ([link](nanocoai/nanoclaw Issue #2781)) — *Support NANOCLAW_NATIVE_CREDENTIALS to bypass OneCLI*  
  User story: Downstream packagers want to inject provider credentials directly via environment variables, skipping OneCLI authentication. This would simplify deployments in sandboxed or immutable environments.

- **PR #2780** ([link](nanocoai/nanoclaw PR #2780)) — *feat(upgrade-state): env opt-out for the startup tripwire (managed fleets)*  
  Adds `NANOCLAW_DISABLE_UPGRADE_TRIPWIRE=1` to skip the upgrade tripwire at startup, logged at `warn`. This is intended for fleets that bake NanoClaw into immutable image installs.

These two features together indicate a growing user base that requires non-interactive, configuration-driven deployment — likely in production or CI/CD scenarios. Both are candidates for inclusion in the next minor release.

## 7. User Feedback Summary

- **Pain points expressed:**
  - Budget/billing errors producing silent failures (now fixed, per user @assapin).
  - Slack integration breaks URLs containing `@` symbols, causing frustration for teams using collaborative note-taking platforms.
  - Security documentation is out of date, confusing new contributors or operators.
  - Credential proxy architecture raises legal/compliance doubts for users relying on Anthropic API.

- **Use cases driving requests:**
  - **Managed fleets / immutable infrastructure** — Users want to disable upgrade tripwires and use externally injected credentials without OneCLI.
  - **Slack as primary channel** — Users heavily rely on agent-generated links in Slack; broken mentions undermine trust.

- **Satisfaction signals:**
  - Quick turnaround on the budget-exhaustion fix (PR #2759 filed and merged in 2 days) demonstrates responsive maintenance.
  - No negative sentiment expressed; all feedback is constructive and feature-oriented.

## 8. Backlog Watch

- **Issue #1669** ([link](nanocoai/nanoclaw Issue #1669)) — *Credential Proxy risk*  
  Opened 2026-04-06, last updated 2026-06-16. This open concern about Anthropic TOS compliance has had only one comment and no maintainer reply. Given the legal ramifications, this issue merits a formal assessment or an update in project documentation.

- **Issue #2783** ([link](nanocoai/nanoclaw Issue #2783)) — *docs/SECURITY.md is outdated*  
  Filed today, but as the canonical security document (linked from multiple places), it should be prioritized to avoid misleading readers. No assignee or response yet.

- **Issue #2784** ([link](nanocoai/nanoclaw Issue #2784)) — *Container-runner staleness bug*  
  Filed today, no comments. While not critical, it could cause subtle runtime inconsistencies in group sessions if not addressed.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-06-17

## 1. Today’s Overview
The NullClaw project shows moderate activity with two issues and three pull requests updated in the last 24 hours. No new releases were published. Two open bugs are drawing community attention, and three open PRs—including a long‑standing cron subagent feature—are being actively worked on. Overall, the project is in a steady development phase with focus on user‑reported stability and authentication improvements.

## 2. Releases
*None in the last 24 hours.*

## 3. Project Progress
No PRs were merged or closed today. However, three open PRs advanced:

- **PR #959** – `fix(cron): persist paired token for scheduler tool access (#839)` – Directly addresses the scheduler access bug reported in Issue #839 by persisting bearer tokens after pairing.
- **PR #958** – `fix(teams): accept lowercase serviceurl JWT claim and raise JWKS fetch cap` – Fixes 403 rejection during MS Teams connector‑token validation.
- **PR #783** – `feat(cron): cron subagent, run history, JSON output, security hardening` – A large feature still open after two months; adds a full cron engine with scheduling, history, and authentication hardening.

All three PRs remain under review.

## 4. Community Hot Topics

**Most active issue:**  
[Issue #952](https://github.com/nullclaw/nullclaw/issues/952) – *bug: Local model using ollama returns incomplete answers*  
- 2 comments, 0 👍  
- User reports that the agent does not answer in complete sentences when using Gemma pulled via Ollama. The issue is still open with no maintainer response yet. Underlying need: reliable and complete output from locally hosted models.

**Notable older issue:**  
[Issue #839](https://github.com/nullclaw/nullclaw/issues/839) – *bug: bit has no access to scheduler !?*  
- 1 comment, open since April 18, 2026  
- User cannot use the scheduler feature from the `bit` context. A fix is in progress via PR #959.

**PRs with engagement:**  
The three open PRs currently have zero comments each, suggesting review is still internal or awaiting community input.

## 5. Bugs & Stability

| Severity | Issue | Description | Fix PR? |
|----------|-------|-------------|---------|
| **High** | [#952](https://github.com/nullclaw/nullclaw/issues/952) | Incomplete answers when using local model (Ollama/Gemma) | No fix PR yet |
| **Medium** | [#839](https://github.com/nullclaw/nullclaw/issues/839) | `bit` user cannot access scheduler tool | PR #959 (open) |

No crash or regression was reported today.

## 6. Feature Requests & Roadmap Signals

- **Scheduler token persistence** – PR #959 is a direct response to Issue #839, indicating that the team is prioritising seamless scheduler access for all agents.
- **MS Teams token compatibility** – PR #958 suggests upcoming support for Microsoft Teams integration, likely in the next patch release.
- **Full‑featured cron subagent** – PR #783 (open since 2026-04-07) remains a strong candidate for the next minor or feature release, bringing scheduled tasks, history, JSON output, and improved security.

These signals point to a near‑term focus on **scheduling reliability**, **third‑party integration fixes**, and **user‑facing job management**.

## 7. User Feedback Summary

**Pain points expressed:**
- Incomplete answers from Ollama‑sourced models (Issue #952) – a blocking usability issue for local‑model users.
- Scheduler not accessible from all contexts (Issue #839) – impacts task automation workflows.

**Use cases:** Running local models (e.g., Gemma) and scheduling recurring tasks via cron.

**Satisfaction:** Not explicitly reported, but the swift PR for scheduler access (#959) indicates maintainers are responsive. The lack of response on the incomplete‑answer bug (#952) may cause some frustration.

## 8. Backlog Watch

These important items have been open for extended periods without a maintainer comment or merged fix:

- **[Issue #839](https://github.com/nullclaw/nullclaw/issues/839)** – *bug: bit has no access to scheduler !?* (opened 2026-04-18)  
  A fix PR (#959) was just submitted, so this is now addressed; still awaiting review/merge.

- **[PR #783](https://github.com/nullclaw/nullclaw/pull/783)** – *feat: cron subagent, run history, JSON output, security hardening* (opened 2026-04-07)  
  No comments from maintainers; a large feature that may require more discussion before merging.

Maintainers are encouraged to review PR #959 promptly and provide a status update on the incomplete‑answer bug (#952).

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-06-17

## 1. Today's Overview
Project activity remains very high. Over the past 24 hours, 21 issues were updated (15 open, 6 closed) and 50 pull requests saw activity (36 open, 14 merged/closed). No new releases were published. The development focus is split between stability fixes for the Reborn UI (automations, skills, conversation display) and architectural improvements to the agent engine (multi-route execution, prompt tightening, no-progress detection). A recurring theme is addressing **automation reliability** and **user-facing clarity** — both from bug reports and from proactive UX enhancements.

## 2. Releases
*None.* No new versions were tagged in the last 24 hours.

## 3. Project Progress
Six issues and 14 PRs were closed/merged today, advancing several feature areas:

| Closed Item | Summary |
|-------------|---------|
| [#2721](https://github.com/nearai/ironclaw/issues/2721) | Engine V2 quality: Milestone 0 + multi-route execution (design closed after evaluation) |
| [#2725](https://github.com/nearai/ironclaw/issues/2725) | Milestone 0 evaluation and go/no-go decision |
| [#2724](https://github.com/nearai/ironclaw/issues/2724) | Tightened CodeAct prompt for simple tasks |
| [#2723](https://github.com/nearai/ironclaw/issues/2723) | Tightened orchestrator loop in `default.py` |
| [#4723](https://github.com/nearai/ironclaw/issues/4723) | Fixed composer hover state border highlight |
| [#4852](https://github.com/nearai/ironclaw/issues/4852) | Fixed missing shell command in approval dialog and activity history |

**Key merged PRs:**
- [#4902](https://github.com/nearai/ironclaw/pull/4902) – **Vision support for inline images** on `/v1/chat/completions` (OpenAI-compat).
- [#4858](https://github.com/nearai/ironclaw/pull/4858) – **Sanitized command details** shown in approval prompts and activity display.
- [#4995](https://github.com/nearai/ironclaw/pull/4995) – **Benchmark forwarding** of `NEARAI_API_KEY` to enable NEAR AI cloud routing.
- [#4954](https://github.com/nearai/ironclaw/pull/4954) – **Approval-gate denial** now surfaces to the model instead of cancelling the run.

These changes improve agent loop robustness, developer observability, and benchmark infrastructure.

## 4. Community Hot Topics
The most active issues (by comment count) are older milestones:

- [#2721](https://github.com/nearai/ironclaw/issues/2721) (3 comments) – Engine V2 multi-route execution design. This long-running epic drove several prompt and loop tightening PRs. Now closed, it signals completion of Milestone 0.

- [#4853](https://github.com/nearai/ironclaw/issues/4853) (1 comment) – Tool activity disappears after completion in multi-tenant environments. This is a live bug affecting Railway deployments; a fix is being tracked.

Among PRs, the largest (size XL) and most strategic pull requests are:

- [#5010](https://github.com/nearai/ironclaw/pull/5010) – **Surface tool calls in OpenAI Responses API** (non-streaming). This is a core API feature that makes IronClaw more compatible with OpenAI clients.
- [#5008](https://github.com/nearai/ironclaw/pull/5008) – **Per-user agent-context profile** (timezone/locale/location) – part of a two-part memory split.
- [#4953](https://github.com/nearai/ironclaw/pull/4953) – **Gate triggered Slack OAuth URL on verified DM** – security follow-up addressing a potential leak of OAuth URLs.

These PRs are under active review and represent high-impact additions to the product.

## 5. Bugs & Stability
Multiple bugs were filed today, ranked by severity:

**High severity:**
- [#4992](https://github.com/nearai/ironclaw/issues/4992) – Local-dev SSO mismatch causes Railway automations to fail before run/thread creation. **Fix PR** [#5003](https://github.com/nearai/ironclaw/pull/5003) is already open.
- [#4991](https://github.com/nearai/ironclaw/issues/4991) – WASM Google Drive auth failures dead-end without refresh-retry. **No fix PR yet**.
- [#5009](https://github.com/nearai/ironclaw/issues/5009) – Live (non-triggered) Slack OAuth path lacks structural DM-parity (security). **Fix PR** [#4953](https://github.com/nearai/ironclaw/pull/4953) addresses the triggered path but the live path remains.

**Medium severity (UX blockers):**
- [#5004](https://github.com/nearai/ironclaw/issues/5004) – Automations failure summary card not actionable (can’t see which automation failed).
- [#5005](https://github.com/nearai/ironclaw/issues/5005) – Automations page lacks management actions (pause, edit, delete).
- [#5007](https://github.com/nearai/ironclaw/issues/5007) – Skills validation error does not clear after fields are filled.
- [#4918](https://github.com/nearai/ironclaw/issues/4918) – Automation logs display zero entries for both successful and failed runs.

**Low severity:**
- [#4988](https://github.com/nearai/ironclaw/issues/4988) – Recent runs visualization (colored dots) is unclear.
- [#4982](https://github.com/nearai/ironclaw/issues/4982) – Automation row selection area is too small.

Several of these have corresponding fix PRs in the pipeline, indicating the team is actively addressing the reported issues.

## 6. Feature Requests & Roadmap Signals
User-reported features from today include:

- [#4963](https://github.com/nearai/ironclaw/issues/4963) – Remove redundant participant names/avatars per message (UI simplification).
- [#4981](https://github.com/nearai/ironclaw/issues/4981) – Redesign dashboard status badges (MUTED, SIGNAL, INFO, SUCCESS) to be more intuitive.
- [#4980](https://github.com/nearai/ironclaw/issues/4980) – Add empty-state guidance for creating automations (currently no button or instructions).
- [#5006](https://github.com/nearai/ironclaw/issues/5006) – Add search/filter to the Skills page and fix metadata formatting.

These are likely to land in the next Reborn UI iteration (v2.x). Additionally, two in-progress PRs signal roadmap direction:

- **Per-user agent-context** ([#5008](https://github.com/nearai/ironclaw/pull/5008)): allows IronClaw to know user’s timezone, locale, and location without prompting.
- **Tool call exposure in OpenAI Responses API** ([#5010](https://github.com/nearai/ironclaw/pull/5010)): improves interoperability with OpenAI SDKs.

Both are candidate for the next minor release.

## 7. User Feedback Summary
The majority of feedback comes from QA (sunglow666) and core contributors. Key pain points:

- **Automations are confusing**: users cannot understand why automations fail, cannot manage them, and the failure card provides no actionable information.
- **Onboarding is weak**: new users don’t know how to create automations or what status badges mean.
- **Validation is inconsistent**: skills form validation errors persist even after the user fills in required fields.
- **Logs are missing**: automation run logs show zero entries, making debugging impossible.
- **Visual clarity issues**: redundant participant labels, unclear dot-based run history, and limited row selection hit area.

Overall satisfaction seems low for the Automations page, but the team is responding with focused PRs (e.g., [#5003](https://github.com/nearai/ironclaw/pull/5003) for SSO recovery).

## 8. Backlog Watch
Only a handful of open items are older than a week:

- [#4518](https://github.com/nearai/ironclaw/pull/4518) – Reborn extension lifecycle e2e coverage (opened June 6, no merge yet). This PR adds important test coverage but may be blocked by other dependencies.
- [#4787](https://github.com/nearai/ironclaw/pull/4787) – Barcelona Hackathon fork (opened June 12). Marked `[NO MERGE]`, but contains useful onboarding changes; should be reviewed for upstream extraction.

No critical issues languish without attention. Most older items (April milestones) were closed today, demonstrating strong issue triage. The team should keep an eye on [#4991](https://github.com/nearai/ironclaw/issues/4991) (Google Drive auth dead-end) which has no fix PR yet and affects any user with expired tokens.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

Here is the structured project digest for LobsterAI based on the provided GitHub data.

---

# LobsterAI Project Digest
**Date:** 2026-06-17

## 1. Today's Overview

Project activity is moderate, with the maintainers focusing on quality-of-life fixes and UI/UX refinements rather than new features. The day saw 5 pull requests merged, primarily targeting the Cowork module (rail navigation jank, task search) and the Artifacts/HTML share system. A single new open issue was reported, concerning a lack of validation for duplicate keyboard shortcuts. The community's bug reports regarding silent failures in scheduled tasks remain open, indicating a need for maintainer attention. Overall, the project is in a stable maintenance phase with ongoing efforts to polish the user experience.

## 2. Releases
**None.**

There were no new releases published in the last 24 hours.

## 3. Project Progress

The following PRs were merged/closed today, representing completed work:

- **[Merged] PR #2173** - `fix(cowork): render user messages as plain text`: Preserves user-entered line breaks in sent message bubbles and adds diagnostic tools for future spacing issues.
    - [Link to PR](https://github.com/netease-youdao/LobsterAI/pull/2173)
- **[Merged] PR #2172** - `fix(html-share): 支持恢复因开启数量上限关闭的分享`: Adds the ability to restore HTML shares that were closed due to hitting the quantity limit. Includes backend documentation and tests.
    - [Link to PR](https://github.com/netease-youdao/LobsterAI/pull/2172)
- **[Merged] PR #2171** - `fix(cowork): reduce rail navigation jank in long sessions`: Optimizes rail performance by avoiding smooth scrolling for long jumps and memoizing rail items to prevent repeated scans. Adds debug logs for navigation issues.
    - [Link to PR](https://github.com/netease-youdao/LobsterAI/pull/2171)
- **[Merged] PR #2170** - `fix(cowork): search tasks from database`: Improves Cowork task search by querying the backing SQLite database instead of filtering the preloaded recent sessions in the modal.
    - [Link to PR](https://github.com/netease-youdao/LobsterAI/pull/2170)
- **[Merged] PR #2169** - `feat(artifacts): 优化预览卡片与浏览器预览体验`: Unifies preview card styles, adds hover subtitles for HTML artifacts, optimizes the "open with" menu, and improves the built-in browser preview experience.
    - [Link to PR](https://github.com/netease-youdao/LobsterAI/pull/2169)

## 4. Community Hot Topics

The most active community discussion centers on a user experience regression:

- **Issue #1425 (Open)** - `[stale] 快捷键重复无校验`: This issue reports that users can save duplicate keyboard shortcuts without any validation or error. This is a clear UX gap that degrades workflow reliability.
    - [Link to Issue](https://github.com/netease-youdao/LobsterAI/issues/1425)

## 5. Bugs & Stability

One new bug was reported today, and a critical pre-existing bug remains unaddressed.

- **[Medium] Issue #1425** - **No validation for duplicate keyboard shortcuts.** Users can assign the same shortcut to multiple functions, leading to ambiguous behavior. No fix PR exists yet.
    - *Severity:* Medium. Impacts usability but not stability.
    - [Link to Issue](https://github.com/netease-youdao/LobsterAI/issues/1425)

- **[High] PR #1424 (Open, Stale)** - **Scheduled Task stop handler is a no-op.** The IPC handler for stopping a scheduled task returns `{ success: true }` without actually stopping the task. Furthermore, UI components do not display error states from the backend, meaning users have no feedback when any scheduled task operation (create, delete, toggle) fails.
    - *Severity:* High. Represents a functional bug with a silent failure, undermining trust in the system.
    - [Link to PR](https://github.com/netease-youdao/LobsterAI/pull/1424)

## 6. Feature Requests & Roadmap Signals

No new explicit feature requests appeared today. However, the PRs merged today signal the development team's current priorities:

- **Cowork UX Stability:** PRs #2171 and #2173 show a focus on making the Cowork module more performant and reliable, especially in long sessions.
- **Artifact/Share Polish:** PRs #2169 and #2172 indicate that the team is refining the artifact sharing and preview experience.
- **Search & Data Integrity:** PR #2170 suggests a move toward more robust, database-backed features rather than in-memory only solutions.

These areas are likely candidates for the next minor release.

## 7. User Feedback Summary

User feedback is currently focused on two distinct pain points:

- **Shortcut Configuration:** The current system allows users to create conflicting keyboard shortcuts (Issue #1425), leading to confusion. The user expects validation on save.
- **Silent Failures:** A prolonged issue (PR #1424) highlights a significant dissatisfaction point: the scheduled tasks UI provides zero feedback when operations fail. The user's underlying need is for transparent system feedback and reliable operation state management.

## 8. Backlog Watch

The following long-standing item requires maintainer attention:

- **PR #1424 (Open, Stale)** - `fix(scheduledTasks)`: **CRITICAL.** This PR has been open since **April 3, 2026** and remains unmerged. It addresses a significant bug where the scheduled task stop command fails silently, and error states are not surfaced to the user. The lack of resolution for over two months is concerning for project health.
    - [Link to PR](https://github.com/netease-youdao/LobsterAI/pull/1424)

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

# TinyClaw Project Digest – 2026-06-17

## 1. Today's Overview
Activity on the TinyClaw repository remains very low. No new issues were filed or updated in the past 24 hours, and no pull requests were merged or closed. The only movement is a single open PR (#281) submitted yesterday that addresses Windows-specific CLI bugs. With zero new releases and no open issues, the project appears to be in a quiet maintenance phase, with the sole focus currently on improving cross‑platform compatibility.

## 2. Releases
*No new releases to report.*

## 3. Project Progress
No pull requests were merged or closed today. The only PR under consideration is:
- **[PR #281](https://github.com/TinyAGI/tinyagi/pull/281) [OPEN]** – Fix: Windows cross-platform support in CLI  
  *Author: mperkins0155 | Created: 2026-06-16*  
  This PR aims to fix three Windows‑only bugs that prevented the `tinyagi` CLI from running natively on Windows (non‑WSL). It remains open and has not yet been reviewed or merged.

## 4. Community Hot Topics
The only active discussion item is **PR #281**, which has no comments or reactions so far. Its presence indicates that Windows usability is a concern for some users, and the author has stepped forward to address it. No other issues or PRs are generating conversation.

## 5. Bugs & Stability
**Severity: Medium**  
Three Windows‑specific bugs have been identified and are being addressed in **PR #281**:
1. Doubled drive letter causing `MODULE_NOT_FOUND` errors.
2. Path resolution issues due to `path.resolve` behaving differently on Windows.
3. Unspecified additional Windows incompatibility.

No other bug reports were filed today. The fix appears to be targeted and isolated to the CLI, with no impact on core agent functionality.

## 6. Feature Requests & Roadmap Signals
No new feature requests were submitted. The only signal is the ongoing effort to improve Windows cross‑platform support, which may become part of the next release or minor patch if merged. No major roadmap changes are evident.

## 7. User Feedback Summary
Direct user feedback is absent in the current data. However, the submission of PR #281 suggests that at least one user (or contributor) encountered real pain points running TinyClaw on native Windows. The lack of any other issues or comments implies that either the user base is small, or the core experience on Linux/macOS is stable.

## 8. Backlog Watch
There are no long‑unanswered issues or PRs requiring maintainer attention at this time. The only open item (PR #281) was submitted just yesterday and is unlikely to be considered “stale.” Maintainers should review it promptly to keep the project moving forward.

---

*Generated from GitHub data for TinyAGI/tinyagi on 2026-06-17.*

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest – 2026-06-17

## 1. Today's Overview
Moltis had a moderate day of activity with 4 issues and 2 pull requests updated in the last 24 hours, though no new releases were published. The project is in an active development phase, with two open enhancements and two open bugs being discussed. One bug was closed swiftly, indicating responsive triage. Two open PRs from contributor `gptme-thomas` are awaiting review, adding configurable context commands and external-agent model/effort selection.

---

## 2. Releases
No new releases were created today. The latest release remains prior to this report.

---

## 3. Project Progress
No PRs were merged or closed today. The two open PRs represent ongoing feature work:
- **#1124** – *Add context command support for chat turns* (open, updated 2026-06-16)  
  Adds an optional `chat.context_command` that runs before each turn, injecting generated runtime context into the prompt.
- **#1125** – *Support model and effort selection for external agents* (open, updated 2026-06-16)  
  Adds `models` and `efforts` configuration for external-agent providers, and exposes them via the `/model` command.

Both are key enhancements for deployment flexibility and multi-agent workflows.

---

## 4. Community Hot Topics
The most active issue is:
- **#1126** – [Feature] *allow to configure the format of tts output*  
  [Link](https://github.com/moltis-org/moltis/issues/1126)  
  **2 comments** – The only issue with multiple comments. The user requests the ability to choose TTS output format (e.g., WAV vs MP3, sample rate). This reflects a need for better audio output configurability, likely for integration with different downstream systems.

The two open PRs have no discussion yet, suggesting they may need maintainer attention to move forward.

---

## 5. Bugs & Stability
Two bugs were reported today, ranked by severity:

**High severity**
- **#1129** – *Lack of echo cancellation causes agent to retrigger itself in live mode* (open, 0 comments)  
  [Link](https://github.com/moltis-org/moltis/issues/1129)  
  A live‑mode echo loop can render the agent unusable in real‑time conversations. This is a critical stability issue for voice‑first use cases. No fix PR is open yet.

**Low severity (resolved)**
- **#1128** – *Transcription errors with self‑hosted whisper.cpp* (closed, 1 comment)  
  [Link](https://github.com/moltis-org/moltis/issues/1128)  
  Closed on the same day it was filed, likely indicating a known configuration issue or quick patch. The closure suggests the team is responsive to self‑hosted transcription problems.

---

## 6. Feature Requests & Roadmap Signals
Two feature requests were opened today:
- **#1126** – *Configure TTS output format*  
  [Link](https://github.com/moltis-org/moltis/issues/1126)  
  Likely to be implemented in a future minor release given the project’s current work on voice features.
- **#1127** – *Configure RPC timeout*  
  [Link](https://github.com/moltis-org/moltis/issues/1127)  
  A small but practical quality‑of‑life improvement for deployments with unreliable network links.

Both are straightforward configuration additions and may appear in the next version alongside the ongoing PRs (#1124, #1125).

---

## 7. User Feedback Summary
- **Pain points**  
  - Echo cancellation missing in live mode (#1129) – causes agent self‑triggering.  
  - Inability to control TTS output format (#1126) – needed for downstream audio pipelines.  
  - Fixed transcription issues with self‑hosted whisper.cpp (#1128) – signals users experiment with local models.
- **Use cases**  
  - Voice‑first interaction (live mode, TTS).  
  - Self‑hosted deployments (whisper.cpp).  
  - Multi‑agent systems (external agent model/effort in PR #1125).
- **Satisfaction**  
  The rapid closure of #1128 suggests responsive support. No overt dissatisfaction expressed, but the echo cancellation bug (#1129) is a clear blocker for live voice use.

---

## 8. Backlog Watch
No long‑unanswered issues or PRs were identified in today’s data snapshot. All tracked items were updated within the last 24 hours. The two open PRs (#1124, #1125) have been awaiting review since 2026-06-15 and could benefit from maintainer feedback to progress.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

Here is a structured project digest for CoPaw based on the provided GitHub data for 2026-06-17.

---

## CoPaw Project Digest — 2026-06-17

### 1. Today's Overview
The project is in a period of high activity, with 27 issues and 40 pull requests updated in the last 24 hours, alongside the release of a new beta version. The community is highly engaged, particularly around critical stability issues like macOS crashes and process freezing, while also pushing forward significant new features and internationalization. The maintainers are actively merging PRs and addressing regressions, indicating strong project momentum. The release of `v1.1.12-beta.1` demonstrates a focus on security and desktop stability.

### 2. Releases
A new beta version, **`v1.1.12-beta.1`** , was released today.
- **Changes:**
    - **Security:** Fixed a critical issue where the keychain master key was not isolated per installation ([#5028](https://github.com/agentscope-ai/QwenPaw/pull/5028)).
    - **Desktop (Stability):** Hardened the Tauri Windows CI to better handle `crates.io` fetch failures ([#5125](https://github.com/agentscope-ai/QwenPaw/pull/5125)).
    - **Refactoring:** Contains unspecified refactoring work in the console.
- **Breaking Changes:** None explicitly mentioned.
- **Migration Notes:** As this is a beta release, users are advised to test this version in non-production environments before upgrading.

### 3. Project Progress
Today saw a high volume of merged/closed PRs (25 total), indicating rapid advancement on several fronts.
- **Bug Fixes & Stability:**
    - **Desktop Crash:** A fix for the `SIGSEGV` crash on macOS ARM64 by Tauri plugin dependency repair ([#5238](https://github.com/agentscope-ai/QwenPaw/pull/5238)).
    - **Memory Crashes:** A new fix specifically mitigates the ChromaDB-related `SIGSEGV` on macOS by disabling vector operations ([#5246](https://github.com/agentscope-ai/QwenPaw/pull/5246)).
    - **Context Compaction:** A PR adds timeout protection to agent calls during context compaction to prevent process freezes ([#5242](https://github.com/agentscope-ai/QwenPaw/pull/5242)).
    - **Performance:** An optimization was merged to remove unnecessary deep-copy operations in the agent config caching layer ([#5240](https://github.com/agentscope-ai/QwenPaw/pull/5240)).
- **New Features:**
    - **Cron Jobs:** A PR merges a `silent` option for cron agent jobs to prevent them from interrupting main chat conversations ([#5251](https://github.com/agentscope-ai/QwenPaw/pull/5251)).
    - **Console UI:** A new feature adds a "Filter by Title" input for the session list, addressing long-standing community feedback ([#5178](https://github.com/agentscope-ai/QwenPaw/pull/5178)).
    - **Internationalization (i18n):** Vietnamese language support has been added for both the Console UI ([#5175](https://github.com/agentscope-ai/QwenPaw/pull/5175)) and a full Vietnamese README ([#5245](https://github.com/agentscope-ai/QwenPaw/pull/5245) - pending merge).
- **Testing:** Integration tests for cron execution and tool API were added for the Sprint 2.4 milestone ([#5201](https://github.com/agentscope-ai/QwenPaw/pull/5201)).

### 4. Community Hot Topics
The following issues and PRs generated the most discussion and activity, highlighting key community interests.

- **Most Discussed Issue:** **[#5218 - [Bug] 子Agent触发上下文压缩时QwenPaw进程冻结无响应](https://github.com/agentscope-ai/QwenPaw/issues/5218)** (14 comments)
    - **Context:** A critical user-reported bug where context compaction by a sub-agent causes a complete UI freeze.
    - **Needs:** A reliable, non-blocking context management system, especially for complex multi-agent workflows. This underscores a deep need for process isolation and timeout handling.

- **Feature with High Engagement:** **[#5063 - Integrate Headroom as an optional context compression layer](https://github.com/agentscope-ai/QwenPaw/issues/5063)** (6 comments)
    - **Context:** A feature request for integrating a reversible compression layer to reduce token consumption by 60-95%.
    - **Needs:** Users are seeking pro-active cost and token optimization strategies rather than just reactive bug fixes for context overflow. A corresponding PR ([#5244](https://github.com/agentscope-ai/QwenPaw/pull/5244)) was opened today, showing strong developer interest.

- **Active PR:** **[#5242 - fix(compaction): add timeout protection to agent.reply() in _compact_context](https://github.com/agentscope-ai/QwenPaw/pull/5242)**
    - **Context:** First-time contributor `lecheng2018` is directly addressing the root cause of the process freeze in #5218.
    - **Needs:** Demonstrates a healthy community where users contribute fixes for the problems they report.

### 5. Bugs & Stability
Several critical issues were reported, with corresponding fix PRs submitted.

- **Critical:**
    - **[macOS SIGSEGV Crash Loop](https://github.com/agentscope-ai/QwenPaw/issues/5243):** A pervasive crash on macOS ARM64 traced to ChromaDB Rust bindings. **Fix PR:** [#5246](https://github.com/agentscope-ai/QwenPaw/pull/5246).
    - **[\#5209 - QwenPaw Desktop (Tauri) 崩溃循环 — macOS ARM64](https://github.com/agentscope-ai/QwenPaw/issues/5209):** A desktop crash loop caused by a plugin dependency startup loop. **Fix PR:** [#5238](https://github.com/agentscope-ai/QwenPaw/pull/5238).
    - **[Process Freeze on Context Compaction](https://github.com/agentscope-ai/QwenPaw/issues/5218):** UI freezes when sub-agents trigger context compression. **Fix PR:** [#5242](https://github.com/agentscope-ai/QwenPaw/pull/5242).

- **High Severity:**
    - **[Cron Tasks Interrupting Chat](https://github.com/agentscope-ai/QwenPaw/issues/5250):** Scheduled tasks inject messages into the main chat, interrupting ongoing work. **Fix PR:** [#5251](https://github.com/agentscope-ai/QwenPaw/pull/5251) (merged).
    - **[Gemini Tool Calling Regression](https://github.com/agentscope-ai/QwenPaw/issues/5163):** A regression from v1.1.10 to v1.1.11.post2 broke tool calling for Gemini models.

- **Other Notable Bugs:**
    - [Conversation thinking loop enters deadlock](https://github.com/agentscope-ai/QwenPaw/issues/5162)
    - [Ollama model selection option missing](https://github.com/agentscope-ai/QwenPaw/issues/5233)
    - [Token display unit error (e.g., 1GB shown as 1B)](https://github.com/agentscope-ai/QwenPaw/issues/5239)

### 6. Feature Requests & Roadmap Signals
Several new features are being actively discussed or implemented, signaling future project direction.

- **Agent Self-Evolution:** The [request](https://github.com/agentscope-ai/QwenPaw/issues/5205) for an agent self-evolution mechanism (learning from mistakes) is a strong signal for moving beyond static rule files towards dynamic adaptation.
- **Headroom Integration:** The [feature request](https://github.com/agentscope-ai/QwenPaw/issues/5063) for headroom compression already has a [draft PR](https://github.com/agentscope-ai/QwenPaw/pull/5244), making it highly likely for the next stable release.
- **WeCom Rich Media:** A request to [support combined image and text messages for WeCom (WeChat Work)](https://github.com/agentscope-ai/QwenPaw/issues/5217) highlights a clear user need in the China market.
- **Config Migration:** A [question](https://github.com/agentscope-ai/QwenPaw/issues/5254) about importing configs from OpenClaw/Hermes suggests users are looking for lower switching costs from other agents.
- **Cron Execution:** After the bug fix, the need for **background execution** for cron tasks is now a validated feature request ([#5250](https://github.com/agentscope-ai/QwenPaw/issues/5250)).

### 7. User Feedback Summary
- **Pain Points:**
    - **Stability:** The most vocal pain points are stability, particularly the **macOS crash loop** and **process freezes during context compaction**, which severely disrupt the user experience.
    - **Interruptions:** **Cron tasks interrupting conversations** is a clear workflow disrupter that has been fixed by the community.
    - **Configuration Complexity:** The **sidebar being too complicated** and the **workspace structure lacking conventions** indicate a need for a more streamlined and opinionated out-of-box experience.
    - **Channel Integration:** Users of channels like **DingTalk** and **custom channels** are facing plugin/setup reliability issues.
- **Use Cases:**
    - Users are leveraging CoPaw for **complex multi-agent tasks** (pr #5218).
    - **Scheduled automation** via cron is a key use case being pushed hard by the community.
- **Satisfaction:** The quick release of a beta version and the rapid community-driven fixes for reported bugs show a healthy, responsive ecosystem that users are actively contributing to.

### 8. Backlog Watch
Some important issues are losing visibility and may need maintainer attention.

- **[#4625 - MiniMax-M2.5 模型思考过程返回XML格式导致不兼容](https://github.com/agentscope-ai/QwenPaw/issues/4625):** This issue regarding incompatible XML format output from the MiniMax model has been open since 2026-05-22 and has 6 comments with no PR attached. Long-standing and impacts a specific model provider.
- **[#5088 - feat: initial governance & sandbox interface discussion](https://github.com/agentscope-ai/QwenPaw/pull/5088):** This Breaking Change PR for a governance and sandbox interface has been open since June 10th with no recent activity. It represents a significant architectural change that needs to be resolved before other dependant features can proceed.
- **[#5234 - 云端QwenPaw可通过Prompt Injection链实现完整RCE](https://github.com/agentscope-ai/QwenPaw/issues/5234):** This is a **Critical** security claim about a full RCE via prompt injection. It was opened on 2026-06-16 with only 1 comment. This requires immediate investigation and a public response from the maintainers.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

# ZeptoClaw Project Digest — 2026-06-17

## 1. Today's Overview
ZeptoClaw saw minimal activity in the past 24 hours, with no new issues opened or closed and no releases published. The only movement was a single open pull request (#630) from a Dependabot bot updating the base Docker image from one Debian `trixie-slim` digest to another. No human-authored code changes, feature work, or bug fixes were merged. Overall project velocity appears low, with no signs of urgent maintenance or community engagement.

## 2. Releases
No new releases were published today. The latest release remains unspecified.

## 3. Project Progress
No pull requests were merged or closed today. The only active PR (#630) is an automated dependency bump and has not yet been reviewed or merged.

## 4. Community Hot Topics
There are no issues or pull requests with comments or reactions beyond the single Dependabot PR, which received no discussion. Community activity is effectively zero.

## 5. Bugs & Stability
No bugs, crashes, or regressions were reported today. The dependency bump in PR #630 is a routine security/maintenance update with no expected stability impact.

## 6. Feature Requests & Roadmap Signals
No feature requests were submitted today. Without user input, no predictions can be made regarding the next version's direction.

## 7. User Feedback Summary
No user feedback (pain points, use cases, satisfaction/dissatisfaction) was recorded in the tracked GitHub data today.

## 8. Backlog Watch
No long-unanswered issues or PRs are currently open (the only open PR was created yesterday). Maintainers have no pending community items to address.

---

*All data sourced from [github.com/qhkm/zeptoclaw](https://github.com/qhkm/zeptoclaw).*

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest – 2026-06-17

## 1. Today’s Overview

The ZeroClaw project saw exceptionally high activity on 2026-06-17, with **50 pull requests updated in the last 24 hours** and only 2 open issues. The majority of PRs (48) remain open, indicating a large volume of ongoing development and review. Two PRs were merged or closed during the period. The current release cycle is between versions (no new releases today), and the team is actively working toward the **v0.8.1 milestone**, as tracked in issue #6970. Overall project health is strong, with significant cross‑domain contributions spanning security, configuration, memory, channels, and the CLI.

## 2. Releases

**No new releases today.** The latest published release remains unchanged.

## 3. Project Progress

Two pull requests were merged or closed in the last 24 hours. While the report does not list their IDs explicitly, the PR churn of 50 items suggests a healthy pace of review and integration. Notable open PRs that advanced today include:

- **#7825** – Moves Discord from an opt‑in extra into the `default‑channels` bundle, making it easier for new users to get the most‑requested chat surface.
- **#7284** – Critical security fix: per‑agent workspace directory creation and Android shell support.
- **#7668** – Fixes a nested runtime panic when dropping the Postgres memory client.
- **#7778** – Fixes agent tool call emission so live cards render promptly.
- **#7826** – Moves credential redaction from the tool execution path to the rendering layer, preventing legitimate credentials from being scrubbed in model‑facing output.

A full list of the top‑20 PRs by comment count is provided in the data overview.

## 4. Community Hot Topics

The most active discussion thread today is **Issue #6970**, which serves as the operational tracker for the **v0.8.1 integration/channel/provider/tool queue**. It has 3 comments and is labelled as high‑risk, accepted. The issue complements #6489 and is central to the team’s near‑term planning. No PRs accumulated significant comment volume (all shown as `undefined`), indicating that code review is progressing without protracted debate.

- [#6970 [Tracker]: v0.8.1 integration/channel/provider/tool queue and history](https://github.com/zeroclaw-labs/zeroclaw/issues/6970)

## 5. Bugs & Stability

Several high‑risk bug fixes were updated today, reflecting the project’s focus on reliability:

- **#7284 (risk: high)** – Per‑agent workspace directory creation not being enforced; fix includes Android shell support. *Fix PR exists.*
- **#7668 (risk: medium)** – Postgres client drop panic due to nested Tokio runtime. *Fix PR exists.*
- **#7778 (risk: high)** – Agent not emitting pending `ToolCall` events, causing live cards to not render. *Fix PR exists.*
- **#7826 (risk: high)** – Credential redaction stripped legitimate outputs from tool execution. *Fix PR exists.*
- **#7793 (risk: high)** – Doctor validation incorrectly failing for custom/openai‑compatible providers when no endpoint is configured. *Fix PR exists.*
- **#7424 (risk: high)** – `web_fetch.allowed_private_hosts = ["*"]` not covering DNS‑resolved private hosts. *Fix PR exists.*
- **#7215 (risk: high, needs‑author‑action)** – Quickstart wizard missing the `port` field for webhook channels.
- **#7532 (risk: medium, needs‑author‑action)** – Config save round‑trip loss when serde defaults differ from struct `Default`.
- **#7495 (risk: medium)** – Lark/Feishu `ack_reactions` override had no effect.
- **#7094 (risk: medium, stale‑candidate)** – `zeroclaw models set` did not persist the model in config.
- **#7823 (risk: medium)** – ZeroCode approval overlay background missing after clear.

No new bugs were filed today; the only new issue (#7824) is a feature request.

## 6. Feature Requests & Roadmap Signals

The single new issue today is **#7824**, requesting **WeCom (WeChat Work) WebSocket channel support for proactive messaging and media file sends**. This aligns with a pattern seen in other channels (Telegram, Discord) and suggests the community wants parity across all chat surfaces.

On the PR side, several large‑scale features are under active development:

- **#7098 (risk: high)** – Opt‑in WebSocket listener mode for Mattermost (near‑real‑time events).
- **#7223 (risk: medium)** – Slash‑command support in the gateway web chat input.
- **#7081 (risk: high)** – Memory hygiene extension to prune daily and core database rows.
- **#7535 (risk: medium)** – WhatsApp reaction support (`add_reaction`/`remove_reaction`).
- **#7340 (risk: high)** – Domain/URL validation extraction and IPv6 bracket fix.

These features are likely candidates for inclusion in **v0.8.2** or the next minor release, given their maturity and risk classification.

## 7. User Feedback Summary

User feedback is reflected in the issues and PRs filed:

- **Pain point: wecom_ws cannot send proactive messages.** Issue #7824 expresses a clear need for outbound messaging and media file support on the WeChat Work channel.
- **Config round‑trip loss** (PR #7532) caused silent field flips – a subtle but frustrating bug reported by a contributor.
- **Quickstart setup friction** (PR #7215) – new users were unable to complete first‑time use when a webhook channel required a port.
- **WebSocket vs polling latency** (PR #7098) – Mattermost users want near‑real‑time events to reduce latency.
- **Per‑agent workspace security** (PR #7284) – security‑conscious users discovered the shell tool’s working directory was not properly isolated.

Overall, contributors are actively addressing pain points, and the project shows a positive trend of user‑driven improvements.

## 8. Backlog Watch

Several open PRs have been tagged with `needs‑author‑action` or `stale‑candidate`, indicating they require maintainer or author attention to move forward:

- **#7098** – Mattermost WebSocket mode (stale‑candidate, needs‑author‑action). Awaiting final revisions.
- **#7215** – Quickstart port field fix (needs‑author‑action). Addresses a critical FTUE blocker.
- **#7094** – CLI model persist fix (stale‑candidate, needs‑author‑action). Small but important CLI usability fix.
- **#7532** – Config serde defaults alignment (needs‑author‑action). Medium‑severity config bug.

These items should be prioritised to avoid accumulating technical debt. Additionally, the **v0.8.1 tracker (#6970)** is a central coordination issue that may need regular grooming.

*Data sourced from GitHub: zeroclaw-labs/zeroclaw*

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*