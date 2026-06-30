# OpenClaw Ecosystem Digest 2026-06-30

> Issues: 60 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-30 10:45 UTC

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

# OpenClaw Project Digest — 2026-06-30

## Today's Overview

OpenClaw saw extremely high activity on 2026-06-30, with **60 issues updated** (51 open) and **500 pull requests updated** (467 open, 33 merged/closed). Despite the volume, no new releases were published. The project remains in a heavily active development phase, but the sheer number of open PRs (467) and high-severity issues (many P1) indicates ongoing stability and regression challenges. The majority of today’s merged PRs were small resilience fixes (guardrails for malformed plugin manifests), while open issues cluster around message loss, session state corruption, and authentication provider breakage—suggesting core reliability bottlenecks are being addressed incrementally but not yet resolved.

---

## Releases

*None.* No new versions or releases were published on 2026-06-30.

---

## Project Progress

**33 PRs were merged or closed** today, predominantly authored by maintainers and focused on hardening plugin SDK infrastructure against malformed metadata. Key contributions:

- **Plugin manifest resilience** – Multiple PRs from vincentkoc (e.g., [`#90090`](https://github.com/openclaw/openclaw/pull/90090), [`#90054`](https://github.com/openclaw/openclaw/pull/90054), [`#90052`](https://github.com/openclaw/openclaw/pull/90052), [`#90073`](https://github.com/openclaw/openclaw/pull/90073), [`#90068`](https://github.com/openclaw/openclaw/pull/90068), [`#90064`](https://github.com/openclaw/openclaw/pull/90064), [`#90061`](https://github.com/openclaw/openclaw/pull/90061), [`#90062`](https://github.com/openclaw/openclaw/pull/90062), [`#90063`](https://github.com/openclaw/openclaw/pull/90063), [`#90066`](https://github.com/openclaw/openclaw/pull/90066), [`#90077`](https://github.com/openclaw/openclaw/pull/90077), [`#90085`](https://github.com/openclaw/openclaw/pull/90085), [`#90095`](https://github.com/openclaw/openclaw/pull/90095), [`#90099`](https://github.com/openclaw/openclaw/pull/90099), [`#90102`](https://github.com/openclaw/openclaw/pull/90102)) – each adds per-row isolation for QA runners, facade registries, runtime boundaries, metadata snapshots, provider credentials, tool metadata, etc. These fix crashes when a single malformed plugin row would block all subsequent healthy registrations.

- **Subagent completion reliability** – PR [`#90041`](https://github.com/openclaw/openclaw/pull/90041) (still open) prevents `message_tool_only` delivery mode from swallowing subagent completion messages; it also fixes a lifecycle issue where pending tool boundaries were incorrectly resolved.

- **Telegram delivery guard** – PR [`#90066`](https://github.com/openclaw/openclaw/pull/90066) suppresses reconnect drain re-entry while a delivery is in-flight, reducing log noise and potential duplicate messages.

- **Abort-path session lock release** – PR [`#90065`](https://github.com/openclaw/openclaw/pull/90065) (needs proof) bounds the session lock release during aborts, preventing 5–30 minute unresponsiveness on embedded-run timeouts.

- **Tool result preservation** – PR [`#97742`](https://github.com/openclaw/openclaw/pull/97742) ensures structured tool results (e.g., attachments) are not replaced with literal placeholders like `(see attached image)`.

- **Trajectory export restriction** – PR [`#97840`](https://github.com/openclaw/openclaw/pull/97840) restricts trajectory export to owners only, closing a security gap.

- **Closed bug fixes** – Issues [`#98099`](https://github.com/openclaw/openclaw/issues/98099) (SQLite auth migration leaving empty tables), [`#98120`](https://github.com/openclaw/openclaw/issues/98120) (normalizeAccountId crash), [`#98033`](https://github.com/openclaw/openclaw/issues/98033) (iOS token loss on QR pairing), and [`#98112`](https://github.com/openclaw/openclaw/issues/98112) (orphaned subagent runs) were closed, suggesting fixes were deployed or merged.

---

## Community Hot Topics

The most active issues and PRs by comment count and reactions reveal significant user frustration around message reliability and session state. Top items:

| Issue / PR | Comments | Reactions | Summary |
|---|---|---|---|
| [`#84516`](https://github.com/openclaw/openclaw/issues/84516) (open, P1) | 11 | 👍 2 | Codex app-server truncates long agent replies at ~1000 chars silently – no error reported. |
| [`#92433`](https://github.com/openclaw/openclaw/issues/92433) (open, P1) | 6 | 👍 1 | Subagent completion dropped when announce steers into a requester run that ends early. |
| [`#84569`](https://github.com/openclaw/openclaw/issues/84569) (open, P1) | 6 | 👍 3 | WhatsApp session stalls on long model call, reply never delivered. |
| [`#84536`](https://github.com/openclaw/openclaw/issues/84536) (open, P1) | 5 | 👍 1 | Preemptive context overflow silently kills embedded sessions without notification. |
| [`#84527`](https://github.com/openclaw/openclaw/issues/84527) (open, enhancement, P2) | 3 | 👍 10 | Feature request: add Antigravity CLI (agy) backend as replacement for deprecated Google Gemini CLI. |
| [`#97742`](https://github.com/openclaw/openclaw/pull/97742) (open PR) | undefined | 👍 0 | Fix preserving structured tool result text across providers – directly addresses the silent truncation/loss patterns. |

**Analysis:** The recurrence of **message loss** (truncation, dropped completions, stalled sessions, silent kill) across multiple channels (Codex, WhatsApp, Telegram, Discord) points to a systemic reliability gap in the session/run lifecycle. Users are also pressing for **model provider parity** (#84527, #84528) – integration with subscription-backed models like Antigravity CLI (Google’s replacement for Gemini CLI) and official OAuth paths for premium models.

---

## Bugs & Stability

Today’s bug report volume was high, with **10+ P1 and P2 issues filed**. Many are regressions or reliability crashers. Below are the most critical, ranked by severity:

### 🔴 P1 – Critical / Data Loss

| Issue | Description | Fix Status |
|---|---|---|
| [`#98120`](https://github.com/openclaw/openclaw/issues/98120) (closed) | `normalizeAccountId` TypeError crashes gateway WebSocket every 6–7 min on Windows (regression in v1.5.1). | **Closed** – fix merged. |
| [`#98112`](https://github.com/openclaw/openclaw/issues/98112) (closed) | Full-process gateway restart orphans in-flight delegated subagent runs; subagent answer lost. | **Closed** – fix likely merged. |
| [`#98107`](https://github.com/openclaw/openclaw/issues/98107) (open) | Gateway regenerates `service-env` file on every restart, wiping Telegram bot tokens and other secrets with placeholder values. | No fix PR yet. |
| [`#98081`](https://github.com/openclaw/openclaw/issues/98081) (open) | Discord turns missed or silently finish across gateway restart on 2026.6.10. | No fix PR yet. |
| [`#98076`](https://github.com/openclaw/openclaw/issues/98076) (open) | Telegram inbound documents silently dropped if received during gateway upgrade/restart. | No fix PR yet. |
| [`#98101`](https://github.com/openclaw/openclaw/issues/98101) (open) | HTTP 429 overloaded error misclassified as rate limit – blocking z.ai users. | No fix PR yet. |
| [`#98100`](https://github.com/openclaw/openclaw/issues/98100) (open) | z.ai endpoint auto-detection fails due to missing User-Agent header, blocking onboarding. | No fix PR yet. |
| [`#98116`](https://github.com/openclaw/openclaw/issues/98116) (open) | iOS Chat duplicates final assistant reply briefly before reconciliation. | No fix PR yet. |
| [`#98118`](https://github.com/openclaw/openclaw/issues/98118) (open) | Local Docker source builds OOM during `tsdown` – build blocker. | No fix PR yet. |

### 🟡 P2 – Significant but Non-Critical

| Issue | Description | Fix Status |
|---|---|---|
| [`#97651`](https://github.com/openclaw/openclaw/issues/97651) (open) | Tool call output contaminates conversation prefix, collapsing DeepSeek prefix cache hit rate (regression). | No fix PR. |
| [`#97911`](https://github.com/openclaw/openclaw/issues/97911) (open) | `tools.deny: ["skill_workshop"]` does not hide Skill Workshop from Codex deferred tools – security/gate issue. | No fix PR. |
| [`#98062`](https://github.com/openclaw/openclaw/issues/98062) (open) | iOS app fails to connect over Tailscale CGNAT (100.x.x.x) – WebSocket upgrade silently dropped. | No fix PR. |
| [`#98044`](https://github.com/openclaw/openclaw/issues/98044) (open) | Android onboarding does not request camera permission for QR setup – usability blocker. | No fix PR. |

### 🟢 Regression Highlights

- **#97651** (tool call output contaminating prefix) – regression in v2026.6.10.
- **#96120** (gateway crash) – regression after a recent upgrade.
- **#84567** (isolated cron setup timeout) – regression in v2026.5.18 affecting bundled Codex harness.
- **#84203** (cold auth latency on Windows) – regression in v2026.5.18 + Codex 0.131.0.

**Stability trend:** The abundance of P1 bugs opened today, especially around gateway restarts and iOS connectivity, suggests the v2026.6.10 release introduced several regressions. The large number of open Plugin SDK hardening PRs (still awaiting review) indicates the team is aware of fragility in plugin metadata ingestion but the backlog remains heavy.

---

## Feature Requests & Roadmap Signals

Several feature requests filed today align with long-standing community asks:

- **Memory distillation gate** – [`#98092`](https://github.com/openclaw/openclaw/issues/98092) proposes an optional LLM extraction step for `memory-core` deep-phase promotion, converting raw snippets into atomic facts rather than dumping raw text into MEMORY.md. This could be a high-impact UX improvement for long-running agents.

- **Skill Workshop extensibility** – [`#98069`](https://github.com/openclaw/openclaw/issues/98069) requests runtime validation for skill proposals before apply, and [`#98070`](https://github.com/openclaw/openclaw/issues/98070) asks for a composable low-level variant of Skill Workshop for custom creation workflows. Both indicate power users want more control over skill development.

- **Third-party model provider plugins** – [`#98114`](https://github.com/openclaw/openclaw/issues/98114) (closed as feature request) calls for a defined onboarding path for third-party AI providers via ClawHub – a signal that the community wants plugin market growth.

- **Configurable memory pressure thresholds** – [`#98089`](https://github.com/openclaw/openclaw/issues/98089) (closed) asks for tunable thresholds to avoid premature restarts on low-RAM hosts.

- **Antigravity CLI (agy) backend** – [`#84527`](https://github.com/openclaw/openclaw/issues/84527) (with 10 👍) remains open and is time-sensitive: Google’s Gemini CLI deprecation deadline is June 18, 2026 (already passed). The PR exists ([`#linked-pr-open`](https://github.com/openclaw/openclaw/pulls?q=is%3Apr+84527)), but not yet merged.

**Prediction for next version (2026.6.13 or 2026.7.x):** Expect at least a late merge for Antigravity CLI support, continued plugin hardening PRs, and a focus on fixing the gateway restart regression cluster (issues #98076, #98081, #98107). Memory distillation and Skill Workshop extensions are lower-probability for immediate inclusion but likely in the roadmap.

---

## User Feedback Summary

The tone of today’s issues is **frustrated but constructive**. Users consistently report:

- **Silent failures** as the most painful experience: messages disappear without error, sessions are killed without notification, and replies are truncated without any log indication.
- **Regressions after upgrades**: Several users (e.g., #84569, #84203, #97651) note that things worked before a specific version (v2026.5.18 or v2026.6.10) and broke after.
- **Configuration pitfalls**: Issues like `tools.deny` not working, service-env file wiping tokens, and unresolvable model refs after config edits (#84212) highlight that the configuration system has opaque validation gaps.
- **Cross-platform pain points**: iOS and Android users face connectivity and permission issues; Windows users see high auth latency.

Positive signals include high engagement on feature requests (10 👍 for Antigravity CLI), and the community stepping in with pull requests to fix issues (e.g., #97742, #98104). Users want OpenClaw to be reliable and extensible – not just feature-rich.

---

## Backlog Watch

Long-standing issues that remain unattended or need maintainer escalation:

| Issue | Created | Age | Urgency |
|---|---|---|---|
| [`#27482`](https://github.com/openclaw/openclaw/issues/27482) | 2026-02-26 | ~4 months | Feature: direct video upload to LLM. No fix PR. |
| [`#81099`](https://github.com/openclaw/openclaw/issues/81099) | 2026-05-12 | ~7 weeks | P1: `AskUserQuestion` tool call never returns result in claude-cli backend. Stale label. |
| [`#84127`](https://github.com/openclaw/openclaw/issues/84127) | 2026-05-19 | ~6 weeks | P1: webchat dashboard regression after v5.18 upgrade. No fix PR. |
| [`#84154`](https://github.com/openclaw/openclaw/issues/84154) | 2026-05-19 | ~6 weeks | P2: Telegram group message recorded but run not dispatched until next message. Needs product decision. |
| [`#84536`](https://github.com/openclaw/openclaw/issues/84536) | 2026-05-20 | ~6 weeks | P1: preemptive context overflow kills embedded sessions silently. No fix PR. |
| [`#84567`](https://github.com/openclaw/openclaw/issues/84567) | 2026-05-20 | ~6 weeks | P1: isolated cron harness still hangs in v2026.5.18. Linked PR open but not merged. |

These issues, especially #81099 and #84536, are high-impact and have been open for over a month without resolution. The maintainers should prioritize them alongside the fresh P1 regressions filed today.

---

*Generated from GitHub data snapshot on 2026-06-30. All counts and links are as reported.*

---

## Cross-Ecosystem Comparison

# Cross-Project Ecosystem Comparison Report — 2026-06-30

## 1. Ecosystem Overview

The open-source personal AI assistant ecosystem is in a phase of **rapid, parallel evolution** with over a dozen projects competing and complementing each other. Activity is heavily concentrated around core reliability (message delivery, session state, plugin SDK hardening) while expanding channel support (WeChat, Discord, DingTalk, Telegram, Feishu) and provider integrations (DeepSeek V4, Anthropic OAuth, Antigravity CLI). The community is vocal about silent failures and regression management, and projects that demonstrate high merge rates and responsive triage (e.g., NanoBot, CoPaw, Hermes Agent) are gaining trust. Despite the fragmentation, common challenges in session lifecycle, message loss, and config validation are driving convergent solutions across projects.

## 2. Activity Comparison

| Project | Issues Updated | PRs Updated | Merged/Closed PRs | New Release | Health Score* |
|---|---|---|---|---|---|
| OpenClaw | 60 | 500 | 33 | No | 3 / 5 |
| NanoBot | 14 | 52 | 31 | No | 5 / 5 |
| Hermes Agent | 11 | 50 | 32 | No | 5 / 5 |
| PicoClaw | 6 | 7 | 2 | Nightly | 3 / 5 |
| NanoClaw | 1 | 6 | 0 | No | 1 / 5 |
| NullClaw | 1 | 4 | 1 | No | 2 / 5 |
| IronClaw | 9 | 50 | 20 | No | 3 / 5 |
| LobsterAI | 12 | 16 | 13 | Yes | 5 / 5 |
| CoPaw | 15 | 50 | 24 | No | 4 / 5 |
| ZeroClaw | 8 | 50 | 9 | No | 3 / 5 |
| TinyClaw | 0 | 0 | 0 | No | 1 / 5 |
| Moltis | 0 | 0 | 0 | No | 1 / 5 |
| ZeptoClaw | 0 | 0 | 0 | No | 1 / 5 |

*Health Score: 5 = high merge rate, responsive to bugs, stable CI; 3 = active but with open blockers; 1 = no observable activity.

## 3. OpenClaw’s Position

OpenClaw remains the **largest and most complex** project by raw PR count (500+ updates) and is explicitly the “core reference” implementation. Its community is an order of magnitude larger than peers, but this scale also brings **fragmentation and regression debt** — 467 open PRs and 10+ P1 bugs filed today, many around message truncation, gateway restart crashes, and session state corruption. Compared to peers:

- **NanoBot** and **CoPaw** merge a higher percentage of their PRs daily (60%+ vs. 6.6% for OpenClaw), indicating more efficient maintainer workflows.
- **Hermes Agent** and **LobsterAI** focus on specific verticals (MoA pipeline, Cowork UI) and show cleaner release cadences.
- OpenClaw’s advantage is its **plugin SDK** and **multi-channel breadth**; its disadvantage is **reliability instability** that frustrates power users.

## 4. Shared Technical Focus Areas

Several cross-project requirements surfaced in multiple digests:

| Requirement | Projects Affected |
|---|---|
| **Message delivery reliability** (silent truncation, dropped replies, session stalls) | OpenClaw, Hermes Agent, NanoClaw, NullClaw, LobsterAI |
| **Session state / lifecycle fixes** (orphaned subagents, compression-chain disappearance, abort-path lock release) | OpenClaw, Hermes Agent, CoPaw, ZeroClaw |
| **Provider token / credential persistence** (OAuth tokens lost on config refresh, service-env secrets wiped) | NanoBot, PicoClaw, OpenClaw |
| **Channel integration completeness** (Discord attachments, Feishu long messages, DingTalk streaming speed, WhatsApp context loss) | NanoClaw, CoPaw, Hermes Agent, OpenClaw |
| **Plugin/manifest resilience** (malformed metadata blocking all registrations, SSRF bypass) | OpenClaw, PicoClaw, IronClaw |
| **Memory management** (compaction thresholds, context overflow, distillation gates) | NanoBot, OpenClaw, LobsterAI |
| **Configuration validation gaps** (tools.deny not working, model refs unresolvable, autofill hijacking) | OpenClaw, CoPaw, NanoBot |
| **Cross-platform connectivity** (Tailscale CGNAT, Windows auth latency, iOS/Android quirks) | OpenClaw, CoPaw, PicoClaw |

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | Hermes Agent | CoPaw | IronClaw | LobsterAI |
|---|---|---|---|---|---|---|
| **Primary focus** | Core reference, plugin SDK, 20+ channels | Lightweight automation (cron, scripts, WebUI) | Mixture-of-Agents, security hardening | v2.0 prep, DingTalk/WeChat, desktop native | Reborn stack port, QA infra | Cowork UI, OpenClaw bridge |
| **Target users** | Power users, plugin developers | Self-hosting hobbyists, automation-first | Enterprise, MoA pipelines | Chinese-market users (WeChat, DingTalk) | NEAR AI cloud ecosystem | Desktop-centric workflow (cowork) |
| **Architecture** | Monorepo, plugin-driven, high complexity | Modular, Python + Node.js, TDD-heavy | Gateway-stream, strong type system | Tauri desktop, skill pools, runtime v2 | Legacy-to-Reborn migration, E2E test coverage | Electron desktop, cowoork panels |
| **Community style** | High volume, sometimes chaotic | Responsive, high close rate | Security-conscious, fast bug fixes | First-time contributor friendly, Chinese docs | Internal maintainer-driven | Mix of Western and Chinese feedback |
| **Release velocity** | Irregular, nightly builds | Frequent but no tag today | No release today (but many fixes) | Prepping v2.0.0 | No release today | **2026.6.29 release** just published |

## 6. Community Momentum & Maturity

**Tier 1 – High Momentum (merge-heavy, responsive, healthy CI):**
- **NanoBot, Hermes Agent, LobsterAI** — merge rates >80% today, with rapid issue closure and clear release paths. Their communities are constructive and engaged.
- **CoPaw** — very active with 24 merges, first-time contributors, and a v2.0.0 tracker. Slight risk from PR backlog.

**Tier 2 – Moderate Momentum (active but with impedance):**
- **OpenClaw** — enormous PR volume but low merge ratio and many critical regressions. Maintainer bandwidth is spread thin.
- **IronClaw** — heavy feature porting but critical bugs (routine misrouting, nightly E2E failure) unresolved for weeks.
- **ZeroClaw** — intense internal tracker activity but low community engagement; many PRs are maintainer-driven.
- **PicoClaw** — steady but small; nightly builds suggest iterative progress.

**Tier 3 – Low or Stagnant:**
- **NanoClaw** — no merges today, single bug report, low community visibility. Risk of stalled development.
- **NullClaw** — minimal activity; one merge and one new bug. Project appears to be maintained in spare time.
- **TinyClaw, Moltis, ZeptoClaw** — no activity in 24 hours. May be hibernating or abandoned.

## 7. Trend Signals — What AI Agent Developers Should Watch

- **Silent failures are the #1 user pain point.** Projects that invest in clear error surfacing (e.g., Hermes Agent’s credential redaction, PicoClaw’s auth error messages) earn trust. Developers should prioritize explicit failure notifications over silent truncation or dropped sessions.

- **Session lifecycle management is becoming a core primitive.** Across OpenClaw, CoPaw, and ZeroClaw, improving subagent completion, compression chains, and abort-path safety is essential for multi-step workflows. This is an architectural design space where early investment pays off.

- **Cross-channel consistency is non-negotiable.** Users expect the same file-attachment behavior on Discord, Telegram, and WeChat. Projects that treat channels as first-class citizens (e.g., CoPaw’s DingTalk @mention, OpenClaw’s Telegram guard) gain adoption in multi-platform deployments.

- **Provider parity is a competitive differentiator.** Requests for Antigravity CLI (OpenClaw), Anthropic OAuth (NanoBot), and DeepSeek V4 fixes (CoPaw) show that tying a project to a single model provider limits its audience. Plugin-based provider architectures are winning.

- **Observability and memory management are converging.** Features like per-turn token usage (PicoClaw), reasoning effort escalation (NanoBot), and memory distillation gates (OpenClaw) reflect growing demand for transparent, cost-efficient long-running agents.

- **Security hardening is accelerating.** ZIP bomb fixes, API key redaction, symlink guards, and SSRF bypass patches were all merged today. Developers can expect hardening to become a default expectation, not an afterthought.

- **The “lightweight” claim is under scrutiny.** NanoBot’s Issue #660 challenged its Node.js dependency; OpenClaw’s Docker build OOM. The community values true minimalism, not just marketing.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-06-30

## 1. Today's Overview
NanoBot is in a period of **very high activity**: 14 issues were updated in the last 24 hours (5 open, 9 closed) and **52 pull requests** were touched, with 31 merged or closed. This indicates an intensive development cycle, likely in preparation for a new release. No new version was published today. The project continues to refine its automation framework, WebUI, provider support, and agent internals, while addressing several bugs and community feature requests. The maintainer team appears highly responsive, with many issues closed and PRs merged within the same day.

## 2. Releases
No new releases today.

## 3. Project Progress
The most significant merged/closed PRs from the last 24 hours include:

- **Automation & Memory**
  - [PR #4382](https://github.com/HKUDS/nanobot/pull/4382) (merged) — Introduced **script triggers** alongside existing cron automations, moved cron under `automations` package.
  - [PR #4330](https://github.com/HKUDS/nanobot/pull/4330) (merged) — Added a **WebUI automation management view** with filtering, editing, pausing, and deleting user automations.
  - [PR #4370](https://github.com/HKUDS/nanobot/pull/4370) (merged) — Changed default `idleCompactAfterMinutes` from `0` to **15 minutes**, enabling automatic memory compaction.
  - [PR #4369](https://github.com/HKUDS/nanobot/pull/4369) (merged) — Improved `/dream` command to explain empty runs instead of giving an opaque response.

- **Agent & Tooling**
  - [PR #4359](https://github.com/HKUDS/nanobot/pull/4359) (merged) — Fixed **goal continuation context** to include goals created during long tasks.
  - [PR #4349](https://github.com/HKUDS/nanobot/pull/4349) (merged) — Preserved user turns in the **replay‑window history** to avoid starting replay mid‑turn.
  - [PR #4347](https://github.com/HKUDS/nanobot/pull/4347) (merged) — Fixed **model preset switching** in the MyTool system.
  - [PR #4202](https://github.com/HKUDS/nanobot/pull/4202) (merged) — Clarified **filesystem write policy**, aligning `apply_patch` path handling and adding separate read/write directory controls.

- **Provider & Network**
  - [PR #4578](https://github.com/HKUDS/nanobot/pull/4578) (merged) — Added **provider‑scoped proxy configuration** for OpenAI‑compatible providers and OpenAI Codex.

- **Channel & WebUI Fixes**
  - [PR #4381](https://github.com/HKUDS/nanobot/pull/4381) (merged) — **Feishu streaming recovery**: retrying CardKit updates when the first attempt fails.
  - [PR #4487](https://github.com/HKUDS/nanobot/pull/4487) (merged) — Fixed **multi‑file apply_patch edits** preserving distinct file edit records.
  - [PR #4386](https://github.com/HKUDS/nanobot/pull/4386) (merged) — Silenced **unroutable CLI progress** noise for unknown channels.

- **CI & Maintenance**
  - [PR #4400](https://github.com/HKUDS/nanobot/pull/4400) (merged) — **CI now skips** for docs‑only changes.

Additionally, [PR #4385](https://github.com/HKUDS/nanobot/pull/4385) (merged) improved logging by including the primary model error before fallback, and [PR #4386](https://github.com/HKUDS/nanobot/pull/4386) (merged) made Dream runs silent when manual.

---

## 4. Community Hot Topics
The most actively discussed items (by comments or reactions) in the last 24 hours:

1. **[Issue #660](https://github.com/HKUDS/nanobot/issues/660)** (closed) — “Project claims to be 'ultra-lightweight' but includes bloated Node.js dependency”  
   *15 comments, 5 👍* — Users perceive a contradiction between the project’s description and the Dockerfile requiring both Python and Node.js. Underlying need: either clarify the lightweight claim or offer a Node‑free alternative.

2. **[Issue #4419](https://github.com/HKUDS/nanobot/issues/4419)** (open) — “Automatic reasoning effort escalation (default + escalated levels)”  
   *4 comments* — Users want smarter, dynamic adjustment of `reasoningEffort` based on task complexity, rather than a fixed configuration.

3. **[Issue #4418](https://github.com/HKUDS/nanobot/issues/4418)** (closed) — “Heartbeat tasks should deliver results to the channel where the task was added”  
   *4 comments* — The current behavior delivers heartbeat results to the most recent chat channel, which is confusing; the community expects channel‑aware routing.

4. **[Issue #4513](https://github.com/HKUDS/nanobot/issues/4513)** (closed) — “Windows system service restart bug with nssm”  
   *2 comments* — Reports of port conflicts and inconsistent service state after `/restart`.

5. **[Issue #1023](https://github.com/HKUDS/nanobot/issues/1023)** (closed) — “Provider login tokens not persisted + config refresh drops unknown providers”  
   *2 👍* — Critical for OAuth‑based providers; tokens lost on config refresh.

---

## 5. Bugs & Stability
Bugs reported or updated today, ranked by severity:

| Severity | Bug / Issue | Status | Description | Fix PR exists? |
|----------|------------|--------|-------------|----------------|
| **Critical** | [Issue #4595](https://github.com/HKUDS/nanobot/issues/4595) (closed) | `StreamingFileEditTracker.apply_final_call_ids()` overwrites `tool_call.id` for all tools, causing session poisoning. | Fixed? Issue closed, but no linked PR in data. Possibly fixed by #4603 (refactor) which is still open. |
| **High** | [Issue #4592](https://github.com/HKUDS/nanobot/issues/4592) (open) | `ExecTool` path extraction misses absolute paths after `=`, potentially breaking `restrictToWorkspace` guard. | No fix PR yet. |
| **High** | [Issue #4603](https://github.com/HKUDS/nanobot/issues/4603) (open) | Mutation of `tool_call.id` for WebUI progress correlation violates provider protocol. | No fix PR; proposed as a refactor. |
| **Medium** | [Issue #4513](https://github.com/HKUDS/nanobot/issues/4513) (closed) | nssm service restart issues on Windows (port conflict, inconsistent state). | No fix PR visible; closed possibly by workaround. |
| **Low** | [Issue #4599](https://github.com/HKUDS/nanobot/issues/4599) (closed) | Linux install script crashes immediately in TUI. | No fix PR; closed quickly. |

**Note**: The session‑poisoning bug (#4595) is the most serious. Although the issue is closed, the root‑cause PR (#4603) is still open, so the fix may not be merged yet.

---

## 6. Feature Requests & Roadmap Signals
Notable user‑requested features from the last 24 hours:

- **[Issue #4605](https://github.com/HKUDS/nanobot/issues/4605)** (open) — Trigger an agent action from an **external script** (e.g., after Gmail email classification). This reflects demand for programmatic invocation beyond chat interfaces.
- **[Issue #4604](https://github.com/HKUDS/nanobot/issues/4604)** (open) — **Anthropic OAuth** support. A concrete request to enable OAuth‑based login for Anthropic providers.
- **[Issue #4419](https://github.com/HKUDS/nanobot/issues/4419)** (open) — **Automatic reasoning effort escalation**. Likely to be addressed soon, as `reasoningEffort` already exists in config and the community consensus is strong.
- **[Issue #4418](https://github.com/HKUDS/nanobot/issues/4418)** (closed) — Heartbeat result delivery to original channel. Possibly already implemented or planned; closed without a PR link.
- **[Issue #4220](https://github.com/HKUDS/nanobot/issues/4220)** (closed) — **GitHub Copilot for Business / Enterprise** support. Closed but not merged? Likely deferred or solved differently.
- **[Issue #4580](https://github.com/HKUDS/nanobot/issues/4580)** (closed) — Use **conda environment** for subprocesses. Closed as enhancement; likely merged into `feat(exec)` changes (PR #4545 still open).

**Predictions for next version**: The Automation dashboard (#4330) and script triggers (#4382) are already merged — these will be in the next release. Provider‑scoped proxy (#4578) and idle auto‑compact defaults (#4370) are also expected. The reasoning effort escalation (#4419) and external trigger (#4605) are strong candidates for the following minor release.

---

## 7. User Feedback Summary
**Pain points**:
- The “ultra‑lightweight” claim is questioned when Node.js is a hard dependency (#660). Users would like a clarification or a Node‑free Docker image.
- Windows users face service management issues (#4513) and installation script crashes (#4599).
- OAuth token persistence failures (#1023) erode trust in provider authentication.
- The session poisoning bug (#4595) may impact users who rely on file‑edit tools heavily.

**Satisfaction**:
- A new user (Issue #4605) praised the “lightweight codebase” and ease of reading source, contrasting favorably with OpenClaw.
- Many issues are closed promptly (e.g., #660, #1023, #4513 closed within hours), indicating a responsive maintainer team.

**Use cases**:
- Gmail integration (#4605) shows real‑world adoption.
- Enterprise users need GitHub Enterprise support (#4220) and OAuth for Anthropic (#4604).
- Multi‑channel workflows (heartbeat routing #4418) are important for users with many chat rooms.

Overall, the community is **engaged and constructive**, providing detailed bug reports and feature motivations. The high closure rate suggests active triage.

---

## 8. Backlog Watch
Based on the data provided (updated in the last 24 hours), there are **no critical long‑unanswered issues**; most open items have recent maintainer activity. However, a few deserve attention:

- **[Issue #4592](https://github.com/HKUDS/nanobot/issues/4592)** (open, since June 29) — ExecTool path extraction bug. No maintainer comment or linked PR yet. Given the security implications (path‑guard bypass), this should be prioritised.
- **[Issue #4603](https://github.com/HKUDS/nanobot/issues/4603)** (open, June 30) — Refactor request to stop mutating `tool_call.id`. It is directly related to the critical bug #4595, so it is likely being worked on.
- **[Issue #4605](https://github.com/HKUDS/nanobot/issues/4605)** (open, June 30) — External script trigger. While very new, it could become a high‑demand feature if no maintainer responds.

**No old (30+ day) unresolved issues** appear in this snapshot, indicating a healthy backlog turnover.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest – 2026-06-30

## 1. Today’s Overview
The project experienced very high activity on June 30, with **50 pull requests updated** (32 merged or closed) and **11 issues updated** (7 open, 4 closed). No new releases were published. The majority of contributions were bug fixes and security hardening, particularly around the MoA (Mixture-of-Agents) pipeline, gateway session management, credential redaction, and WhatsApp integration. Despite the large volume of fixes, several critical session-state and UI bugs remain open, indicating ongoing stability work.

## 2. Releases
No new releases.

## 3. Project Progress (Merged/Closed PRs)
Today **32 pull requests were merged or closed**, advancing several areas:

**Platform stability & security**
- [PR #27123](https://github.com/NousResearch/hermes-agent/pull/27123) – Added missing `role` field to `systemInstruction` for Gemma/Google compatibility (P1).
- [PR #15093](https://github.com/NousResearch/hermes-agent/pull/15093) – Added session staleness guard to `GatewayStreamConsumer` (P1).
- [PR #27426](https://github.com/NousResearch/hermes-agent/pull/27426) – Two‑layer defense against hallucinated `acp_command` crashing the gateway (P1).
- [PR #29450](https://github.com/NousResearch/hermes-agent/pull/29450) – Fixed ZIP bomb via malicious central directory metadata in `ClawHubSource._download_zip` (P1).
- [PR #55583](https://github.com/NousResearch/hermes-agent/pull/55583) – Redacted Fireworks AI API keys in logs and stack traces (P2).
- [PR #55584](https://github.com/NousResearch/hermes-agent/pull/55584) – Gated WhatsApp owner‑typed forwards on `WHATSAPP_ALLOWED_USERS` allowlist (P2).

**MoA / Agent pipeline**
- [PR #54021](https://github.com/NousResearch/hermes-agent/pull/54021) – Resolved MoA preset to aggregator for auxiliary tasks (P3).
- [PR #54384](https://github.com/NousResearch/hermes-agent/pull/54384) – Propagated `api_mode` from MoA slot runtime to `call_llm`, fixing Copilot GPT‑5.x 400 errors (P2).

**CLI & Kanban**
- [PR #55585](https://github.com/NousResearch/hermes-agent/pull/55585) – Unknown skills now warn instead of crashing the Kanban worker (P3, fixes #27136).

**Gateway & Cron**
- [PR #55559](https://github.com/NousResearch/hermes-agent/pull/55559) – Salvaged session staleness guard for the stream consumer (P2, closes #15093).
- [PR #54349](https://github.com/NousResearch/hermes-agent/pull/54349) – Added audit log for missed cron jobs (feature, still open).

## 4. Community Hot Topics
The most active issues (by comment count) were:

- **[Issue #52445](https://github.com/NousResearch/hermes-agent/issues/52445) (3 comments, P3, open)** – WhatsApp missing from “Deliver to” dropdown in cron job form even when `WHATSAPP_ENABLED=true`. The underlying need is for the GUI dropdown to dynamically reflect enabled integrations rather than a hardcoded list. This is a common UX gap flagged by multiple users.
- **[Issue #27136](https://github.com/NousResearch/hermes-agent/issues/27136) (3 comments, P3, closed)** – Kanban worker crashed on unknown skills, blocking the entire task queue. The fix (#55585) was merged today, addressing a high‑visibility reliability concern.
- **[Issue #55572](https://github.com/NousResearch/hermes-agent/issues/55572) (1 comment, P2, open)** – Tail‑protection token estimation ignores `codex_reasoning_items`, causing Codex/Responses sessions to compact too late. This may lead to out‑of‑context errors in long sessions.

No PRs received comments today though many were opened. The high number of open PRs (18) suggests active collaboration.

## 5. Bugs & Stability
**Critical open bugs (P2):**
- **[#55572](https://github.com/NousResearch/hermes-agent/issues/55572)** – Codex/Responses tail‑protection misses `codex_reasoning_items`; no fix PR yet.
- **[#55589](https://github.com/NousResearch/hermes-agent/issues/55589)** – WhatsApp replies to cron deliveries lose task context (`sweeper:risk-session-state`).
- **[#55588](https://github.com/NousResearch/hermes-agent/issues/55588)** – Compression‑chain sessions disappear from sidebar when root is archived (`sweeper:risk-session-state`).
- **[#55594](https://github.com/NousResearch/hermes-agent/issues/55594)** – TUI long assistant responses become unreachable after scrolling out of view (virtual scroll clamp + sticky scroll interaction). Reported as “#1 UX frustration from daily use”.

**P3 bugs:**
- **[#52445](https://github.com/NousResearch/hermes-agent/issues/52445)** – WhatsApp missing from cron dropdown (duplicate).
- **[#55578](https://github.com/NousResearch/hermes-agent/issues/55578)** – Desktop async delegation completion can revive old session while follow‑up creates a new session (`sweeper:risk-session-state`).

**Fixed today:** P1 bugs #27123, #15093, #27426, #29450; P2 bugs #55582, #55583, #55584, #55559, #55585; P3 bug #55585 (closes #27136). The large number of merged fixes indicates strong maintainer responsiveness.

## 6. Feature Requests & Roadmap Signals
- **[#55573](https://github.com/NousResearch/hermes-agent/issues/55573) (feature, P3, open)** – Support multi‑model configuration/selection for custom OpenAI‑compatible providers (e.g., local Lemonade Server). Currently the interface forces a single model per provider config, limiting flexibility for users running multiple local models.
- **[PR #54349](https://github.com/NousResearch/hermes-agent/pull/54349) (feature, open)** – Cron audit log to surface missed jobs – a clear operational need for production deployments.

**Prediction for next release:** Multi‑model custom providers and the cron audit log are likely candidates. Also expected are the remaining MoA fixes (e.g., full propagation of `api_mode` and proper virtual address handling) and the session‑state bug fixes for compression chains.

## 7. User Feedback Summary
Real pain points expressed through issues:
- **WhatsApp integration gaps:** Missing delivery option in cron form, loss of task context on replies, lack of multi‑model selection for custom providers.
- **Session management frustrations:** Compression chains disappearing, async delegation creating session splits, TUI responses scrolling out of reach.
- **Reliability:** MoA preset errors (400 on “not a valid model ID”), Kanban worker crashes (now fixed), and late session compaction causing out‑of‑context errors.
- **Security concerns:** Raw API keys in logs (Fireworks, Anthropic) – addressed today with multiple redaction PRs.

Overall, users are actively deploying Hermes Agent with WhatsApp, cron, TUI, and MoA features. The project’s rapid bug‑fix cycle indicates high maintainer engagement and is likely improving user satisfaction.

## 8. Backlog Watch
All issues and PRs shown were updated within the last 24 hours, so no long‑unanswered items are visible in this window. However, the closed issue #27136 was originally opened on 2026‑05‑16 and took over a month to be resolved – a reminder that older P3‑P4 bugs may linger. Maintainers should review any remaining unaddressed issues with labels like `sweeper:risk-session-state` or `sweeper:risk-compatibility` that were not touched today.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest – 2026-06-30

---

## 1. Today’s Overview

The project closed two pull requests and published an automated nightly build, while five open bugs and four open feature PRs signal sustained community engagement. A total of six issues were updated (five remain open, one closed) and seven pull requests received activity (two merged, five open). The nightly `v0.3.1-nightly.20260630` release continues the iterative path toward a stable v0.3.1. Overall, the project remains active with a healthy mix of bug fixes, security hardening, and feature work, though several older feature requests and PRs are marked as stale and may need maintainer attention.

---

## 2. Releases

- **nightly (v0.3.1-nightly.20260630.52320f48)**  
  Automated unstable build. No breaking changes or migration notes are documented in the release announcement. Full changelog:  
  [https://github.com/sipeed/picoclaw/compare/v0.3.1...main](https://github.com/sipeed/picoclaw/compare/v0.3.1...main)  

---

## 3. Project Progress

Two pull requests were merged/closed today:

- **fix(providers): surface friendly auth error messages** ([PR #3198](https://github.com/sipeed/picoclaw/pull/3198)) – Improves provider authentication error handling, giving users clearer guidance when API keys or tokens fail.  
- **fix(web): block private IPv4 embeds in ISATAP literals** ([PR #3143](https://github.com/sipeed/picoclaw/pull/3143)) – Fixes an SSRF guard bypass (issue #3074) by teaching the IP classifier to recognize ISATAP IPv6 literals that embed private/loopback IPv4 addresses.

Both merges advance security posture and user experience.

---

## 4. Community Hot Topics

The most active discussions and reactions over the past 24 hours involve:

- **Feature request: SimpleX / Tox gateway** ([Issue #3093](https://github.com/sipeed/picoclaw/issues/3093)) – 4 comments, 1 👍. The author asks for gateway support for SimpleX, Wire, or Tox. This issue has been open since June 10 and remains a popular ask among community members seeking decentralised messaging integrations.  
- **Closed bug: Panel broken on Safari iOS < 16.4** ([Issue #3090](https://github.com/sipeed/picoclaw/issues/3090)) – 3 comments, now closed. The bug affected iOS users, and the closure suggests a fix has been applied or is planned.  
- **Bug: Volcengine Doubao Seed tool calls leak raw text** ([Issue #3153](https://github.com/sipeed/picoclaw/issues/3153)) – 2 comments. Users report that tool calls occasionally appear as `<seed:tool_call>` text instead of being executed, impacting workflows that rely on the Volcengine provider.

Other notable PRs with ongoing community interest:  
- **feat: add deltachat gateway** ([PR #3063](https://github.com/sipeed/picoclaw/pull/3063)) – Open since June 8, still awaiting review.  
- **feat(bedrock): leverage Converse prompt caching** ([PR #3163](https://github.com/sipeed/picoclaw/pull/3163)) – Proposes cache points for AWS Bedrock to reduce costs and latency.

---

## 5. Bugs & Stability

Three new open bugs were reported today, all with zero comments so far:

| Issue | Summary | Severity | Existing Fix PR? |
|-------|---------|----------|------------------|
| [#3195](https://github.com/sipeed/picoclaw/issues/3195) | OpenAI GPT does not work on NanoKVM with default config (v0.2.9) | **High** – Blocks usage on a popular KVM-over-IP device (NanoKVM 2.4.0) | None yet |
| [#3197](https://github.com/sipeed/picoclaw/issues/3197) | Codex and antygravity OAuth login not working (v0.2.9) | **High** – Prevents users of those providers from authenticating | None yet |
| [#3196](https://github.com/sipeed/picoclaw/issues/3196) | Same report as #3197 (duplicate by same author) | **High** | None |

Existing open bug: [#3153](https://github.com/sipeed/picoclaw/issues/3153) (Doubao tool call leak, open since June 22) – **Medium** severity; no fix PR yet.

Today’s merged PR [#3198](https://github.com/sipeed/picoclaw/pull/3198) improves auth error messaging, which may help users diagnose OAuth failures like #3196/#3197, but does not directly fix the underlying issue.

---

## 6. Feature Requests & Roadmap Signals

Active user-requested features (by vote or discussion activity):

- **SimpleX / Tox gateway** ([#3093](https://github.com/sipeed/picoclaw/issues/3093)) – Strong community interest. Likely to be considered for the next stable release if a contributor steps up.  
- **DeltaChat gateway** ([PR #3063](https://github.com/sipeed/picoclaw/pull/3063)) – Ready for review; could land in v0.3.1 if merged.  
- **Remote Pico WebSocket mode** ([PR #3118](https://github.com/sipeed/picoclaw/pull/3118)) – Extends the `picoclaw agent` command to support remote connections. Still open for three weeks.  
- **Bedrock prompt caching** ([PR #3163](https://github.com/sipeed/picoclaw/pull/3163)) – Targets cost efficiency for AWS users.  
- **Per-turn LLM token usage** ([PR #3156](https://github.com/sipeed/picoclaw/pull/3156)) – Improves observability for downstream consumers.

Based on the nightly build tag (`v0.3.1-nightly`), the next stable release will likely be **v0.3.1**. The merges today (auth error messages, SSRF fix) are typical candidates for a point release. Feature PRs that gain maintainer approval this week could also be included.

---

## 7. User Feedback Summary

- **Pain points:**  
  - Authentication failures (OAuth with Codex/antygravity) prevent users from accessing their preferred providers.  
  - OpenAI compatibility on NanoKVM is broken out of the box, frustrating users on that specific hardware.  
  - Tool call leaks (Volcengine) degrade reliability for agent workflows.  
  - Safari iOS incompatibility (now closed) was a usability blocker for mobile users.

- **Use cases voiced:**  
  - Users want to integrate PicoClaw with decentralised messaging networks (SimpleX, Tox, DeltaChat) for private communication channels.  
  - Remote agent mode (WebSocket) indicates demand for distributed or cloud-based deployments.  
  - AWS Bedrock users seek cost optimization via prompt caching.

- **Satisfaction:** The swift closure of the SSRF bypass (PR #3143) and the improved auth error messages (PR #3198) demonstrate maintainer responsiveness, which likely satisfies security-conscious and new users alike.

---

## 8. Backlog Watch

The following issues and PRs have been open for an extended period (≥20 days) with limited maintainer response, and may benefit from a status update or review:

| Item | Type | Opened | Notes |
|------|------|--------|-------|
| [#3093](https://github.com/sipeed/picoclaw/issues/3093) – SimpleX/Tox gateway | Feature Request | 2026-06-10 (20 days) | Marked stale; no PR yet |
| [#3063](https://github.com/sipeed/picoclaw/pull/3063) – DeltaChat gateway | PR | 2026-06-08 (22 days) | Ready for review, no comments from maintainers |
| [#3115](https://github.com/sipeed/picoclaw/pull/3115) – Fix inline data URL media extraction | PR | 2026-06-12 (18 days) | Still open, touches session-history corruption |
| [#3118](https://github.com/sipeed/picoclaw/pull/3118) – Remote Pico WebSocket mode | PR | 2026-06-12 (18 days) | No comments from maintainers |

All of these represent community effort that could deliver value to PicoClaw users; a response or merge decision would help maintain contributor morale and project momentum.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-30

## 1. Today's Overview
The project shows moderate-to-high activity with **1 open issue** and **6 open pull requests** updated in the last 24 hours. None of the PRs have been merged today, indicating that all recent changes are still under review or in progress. The single issue (#2888) reports a functional bug in the Discord adapter around attachment handling, while the open PRs span multiple feature additions (daily news agent, WeChat adapter, template system, Discord adapter) and several bug/security fixes. No new releases were published. Overall, the project is actively developing new capabilities and addressing stability concerns, though community engagement remains low (no reactions on issues/PRs).

---

## 2. Releases
**None.** No new versions were published in the observed period.

---

## 3. Project Progress
No PRs were merged or closed today. However, the following pull requests represent ongoing work that advanced:

- **#2889** – *feat: daily-news-agent for Andy group + WeChat channel* – Adds a scheduled news agent and a new WeChat channel adapter. Includes 33 test cases (TDD-driven).  
  [PR #2889](https://github.com/nanocoai/nanoclaw/pull/2889)

- **#2890** – *feat(templates): agent template loader, setup flow, and docs* – Introduces a reusable template system for agent groups, loadable from public library, local path, or git repo.  
  [PR #2890](https://github.com/nanocoai/nanoclaw/pull/2890)

- **#2884** – *feat(discord): add Discord channel adapter + fix Gateway approval-button routing* – Implements a Discord adapter via the Chat SDK bridge and fixes approval-card button routing.  
  [PR #2884](https://github.com/nanocoai/nanoclaw/pull/2884)

- **#2886** – *fix: channel-registered new agents inherit the install's provider* – Prevents 401 errors on single-provider installs by ensuring new agents use the configured provider instead of the hardcoded default.  
  [PR #2886](https://github.com/nanocoai/nanoclaw/pull/2886)

- **#2880** – *fix(security): contain inbox symlink escapes in attachment writes (#2828)* – Closes a CWE-59 vulnerability that could allow a compromised agent to write arbitrary files on the host via symlink attacks.  
  [PR #2880](https://github.com/nanocoai/nanoclaw/pull/2880)

- **#2885** – *fix(setup): offer Slack Socket Mode in the guided setup flow* – Adds the missing Slack Socket Mode option to the guided setup (previously only available on a feature branch).  
  [PR #2885](https://github.com/nanocoai/nanoclaw/pull/2885)

---

## 4. Community Hot Topics
The only issue with any comment activity is:

- **#2888** – *Discord (and likely other url-only chat-sdk adapters) drop image/file attachments — agent only sees filename*  
  [Issue #2888](https://github.com/nanocoai/nanoclaw/issues/2888)  
  The reporter describes that when a user sends an image, screenshot, or file in Discord, the agent receives only attachment metadata (type, name, mimeType, size) but never the actual content. Telegram works correctly. The root cause is traced to `messageToInbound` in `src/channels/chat-sdk-bridge.ts`, which skips downloading attachment bytes. This issue has 1 comment but no reactions.

**Underlying need:** Users expect cross-channel consistency for file/attachment handling. The fix likely requires modifying the chat-sdk bridge to fetch attachment bytes where the adapter does not provide them inline, or adding a configurable fallback mechanism.

No PRs have comments, so community discussion is minimal. The lack of reactions may indicate a small user base or that the issue has not yet gained visibility.

---

## 5. Bugs & Stability
Reported bugs and fixes (ranked by severity):

| Severity | Issue/PR | Description | Fix PR exists? |
|----------|----------|-------------|----------------|
| **Critical** (Security) | PR #2880 | Symlink-follow vulnerability (CWE-59) allows arbitrary host file writes via agent session dirs. | Yes, PR #2880 is the fix. |
| **High** (Functional) | Issue #2888 | Discord attachments not delivered to agent; only filename metadata is passed. Breaks file-based workflows on Discord. | No dedicated fix PR yet, but PR #2884 adds the Discord adapter and may need to be complemented with a fix. |
| **Medium** (Auth) | PR #2886 | New agents created via channel approval use hardcoded provider (Claude), causing 401 errors on single-provider installs. | Yes, PR #2886. |
| **Low** (Setup) | PR #2885 | Slack Socket Mode not offered in guided setup flow (webhook-only). | Yes, PR #2885. |

All bug/security fixes are currently open but not yet merged.

---

## 6. Feature Requests & Roadmap Signals
The following open PRs represent new features being actively developed. They are strong candidates for the next release:

- **WeChat channel adapter** (PR #2889) – Enables NanoClaw to connect to WeChat, expanding the supported chat platforms beyond Discord, Slack, Telegram, etc.
- **Discord channel adapter** (PR #2884) – Fills a major platform gap, though the attachment bug (issue #2888) will need to be resolved for full functionality.
- **Agent template system** (PR #2890) – Provides reusable agent configurations, lowering the barrier for new users and enabling community sharing of agent bundles.
- **Daily news agent** (PR #2889) – Demonstrates a practical use case (HN + RSS fetch, LLM digest, scheduled task) that could become a built-in template or reference agent.

These features address two user needs: 1) broader platform support (WeChat, Discord) and 2) easier setup and reuse of agent configurations. Future releases are likely to include at least the Discord adapter and the template system, with the news agent as an optional example.

---

## 7. User Feedback Summary
The only direct user feedback in the observed data is from **issue #2888**:

> *“When a user sends an image/screenshot/file in Discord, the agent receives only attachment metadata — never the content. Telegram works fine.”*

- **Pain point:** Cross-channel inconsistency in file attachment handling. Users relying on Discord for image-based workflows find the agent effectively blind to visual content.
- **Use case:** Sharing screenshots, diagrams, or documents via Discord and expecting the agent to analyze them (e.g., vision models).
- **Satisfaction:** The issue is reported as a bug, signaling dissatisfaction with the current behavior. The existence of a working Telegram implementation highlights a regression or omission in the Discord adapter.

No other user sentiment is captured (no 👍/👎 on issues, no comments on PRs).

---

## 8. Backlog Watch
No long-unanswered issues or PRs were identified. All open items are recent (created/updated within the last 2 days). No issues older than 48 hours appear in the data. The project maintainers appear to be actively triaging new submissions.

**Note:** If any issues or PRs from prior weeks are missing from the provided data, they may still exist in the repository but were not updated in the last 24 hours. The digest reflects only items with recent activity.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-06-30

## Today’s Overview
Activity was moderate with 4 pull requests updated in the last 24 hours (1 merged/closed, 3 open) and 1 new bug report. No new releases were published. The merged PR #960 improves CLI usability by adding arrow‑key support in the agent REPL. A single user‑reported bug (#972) highlights a potential stability issue with the Telegram channel after idle periods, though the backend continues to function. A large, long‑standing cron feature PR (#783) saw its latest update today, indicating ongoing development despite being open since April.

## Releases
No new releases were created on this date. The project currently has no published release entries.

## Project Progress
- **PR #960 (merged/closed)** — *fix(cli): handle arrow keys in agent REPL*  
  Adds a compact line editor and POSIX raw‑mode input for the interactive agent REPL, enabling cursor movement, history navigation, backspace, Home/End, and common word‑level shortcuts. This directly addresses a community pain point reported in earlier discussions.  
  [https://github.com/nullclaw/nullclaw/pull/960](https://github.com/nullclaw/nullclaw/pull/960)

- **PR #783 (updated, still open)** — *feat(cron): cron subagent, run history, JSON output, security hardening*  
  This large feature PR was updated today after two months of inactivity. It introduces a database‑backed cron scheduler with job types (skill, agent, shell), timezone support, delivery routing, operator alerts, and enhanced CLI output with JSON formatting. No merge timeline is evident.  
  [https://github.com/nullclaw/nullclaw/pull/783](https://github.com/nullclaw/nullclaw/pull/783)

## Community Hot Topics
No issues or PRs attracted comments or reactions today. The new bug report (#972) is the only community interaction. The most active discussion items remain the open feature PRs (#783, #971, #970), none of which have public commentary. The lack of engagement may indicate a small or passive user base, but also suggests no contentious debates are underway.

## Bugs & Stability
- **Issue #972 (new, open)** — *[bug] telegram channel stop respond after some idle time*  
  **Severity: Medium** — The Telegram channel becomes unresponsive after an idle period (e.g., overnight), while the backend `nullclaw agent` process continues to work. Root cause is unknown; the report shows the agent responds to “ping” commands. No fix PR exists yet. Potential causes include stale WebSocket connections, token expiry, or idle timeout in the Telegram API.  
  [https://github.com/nullclaw/nullclaw/issues/972](https://github.com/nullclaw/nullclaw/issues/972)

- **Stability improvement (merged)** — PR #960’s arrow‑key fix addresses a common CLI crash or misbehavior when users interact with the REPL, improving session robustness.

## Feature Requests & Roadmap Signals
No explicit feature requests were filed as issues. The roadmap is primarily visible through open pull requests:

- **Native tool calls during streaming** (PR #971) — Decouples tool‑call support from the streaming path, enabling providers that support native tools during SSE streaming. This is likely to be merged next as it has been open only since June 29.
- **Cron subagent** (PR #783) — A major new subsystem for scheduled tasks. Its long open duration suggests it may be in a late review cycle or awaiting maintainer bandwidth.
- **REPL improvements** (PR #970) — An enhancement over the merged fix, adding more cursor/editing capabilities; it may follow quickly.

Given the absence of user‑submitted feature requests, the next version will likely focus on integrating the streaming tool‑call and cron features once reviewed.

## User Feedback Summary
- **Pain point (Telegram idle)** — User reports that Telegram channel stops responding after a period of inactivity. The backend is confirmed working, implying a channel‑specific reconnection issue. This is a functional regression for users relying on Telegram as a primary interface.
- **UX dissatisfaction (REPL keyboard handling)** — The merged fix for arrow keys (PR #960) directly addresses user reports of broken cursor navigation in the interactive session. This satisfaction is now delivered.
- **No other feedback** captured today. The project’s community appears to use the software without submitting many compliments or complaints.

## Backlog Watch
- **PR #783 (cron subagent)** — Open since April 7, updated today but still without maintainer review comments. This is the largest pending feature addition and risks becoming stale if not prioritized.  
  [https://github.com/nullclaw/nullclaw/pull/783](https://github.com/nullclaw/nullclaw/pull/783)

- **Issue #972** — New, but no assignee or response yet. Should be triaged quickly to assess if it is a widespread regression or an isolated configuration issue.

- **PR #971 and #970** — Both opened June 29; they have not received any maintainer attention as of today. If left unattended, they may drift into the backlog.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest – 2026-06-30

## 1. Today's Overview
The IronClaw project (Reborn branch) remains in a period of **high-intensity development**, with 50 pull requests updated in the past 24 hours (20 merged/closed, 30 open) and 9 issues updated (8 open, 1 closed). No new releases were cut today. The core team is driving multiple parallel efforts: porting legacy WebUI browser coverage to the Reborn stack, hardening the QA/infrastructure layer, and fixing regressions in routine delivery, web search, and chat UX. Nightly E2E tests continue to fail intermittently (issue #4108), and a new daily failure taxonomy (#5437) suggests the team is methodically tracking CI instability. Overall project health is **active but under pressure** from both feature velocity and stability debt.

## 2. Releases
No new versions were released today.

## 3. Project Progress
Five pull requests were merged/closed today, advancing concrete improvements:

- **[PR #5395]** ([closed] nearai/ironclaw PR #5395) – Fixed Web Access Exa content fetch by routing `get_content` through `web_fetch_exa` while preserving cached `response_id` lookup. Updated schemas to make cached/live modes explicit.
- **[PR #5424]** ([closed] nearai/ironclaw PR #5424) – Linked failed Reborn QA cases to debug artifact paths in Slack reports, improving observability without altering passing-case output.
- **[PR #5436]** ([closed] nearai/ironclaw PR #5436) – Made Reborn WebUI v2 live QA success-wording checks more flexible to match actual canary output variants.
- **[PR #5374]** ([closed] nearai/ironclaw PR #5374) – Ported legacy browser coverage for extension registry, install/remove, channel setup, OAuth setup, configure modal, and extension tabs to the Reborn WebUI v2 test suite.
- **[PR #5373]** ([closed] nearai/ironclaw PR #5373) – Ported channel pairing flows (generic proof-code pairing alongside Slack) and added stable selectors with retry behavior.

Additionally, issue **#5412** ([closed] nearai/ironclaw Issue #5412) – “webui v2: log entry text is not selectable / copyable” – was resolved.

## 4. Community Hot Topics
While most issues have 0–1 comments, the following items attracted attention:

- **[Issue #5420]** ([open] nearai/ironclaw Issue #5420) – The most-discussed bug: routine delivery target is accidentally global instead of per-routine, causing Slack rerouting to affect all automations. The sole comment suggests this is a high-priority regression for users with multiple routines.
- **[Issue #5437]** ([open] nearai/ironclaw Issue #5437) – “Daily ironclaw failure taxonomy – 2026-06-30” documents a full suite of 146 non-pass tests failing identically with HTTP 400 due to a model endpoint issue (deepseek/…). Signals a systemic integration failure that the team is actively monitoring.
- **[PR #5441]** ([open] nearai/ironclaw PR #5441) – Implements header notification bell for automation runs, directly addressing user request #5443. This is a high-visibility UX enhancement likely to be welcomed by the user community.
- **[PR #4841]** ([open] nearai/ironclaw PR #4841) – Open since June 13, aims to eliminate “run-borking” terminal failures by providing recovery/explanation for run-level errors. Its long lifespan indicates the complexity of the refactoring.

The underlying pattern is a **push toward stability and user-facing polish** after a series of feature ports. The delivery-target bug (#5420) especially shows that core abstractions need stronger per-entity scoping.

## 5. Bugs & Stability
Reported bugs ranked by severity:

| Severity | Issue | Description | Fix PR Exists? |
|----------|-------|-------------|----------------|
| **Critical** | [#5420](nearai/ironclaw Issue #5420) | Routine delivery target is global per‑user, not per‑routine – setting Slack for one routine reroutes all routines. | No |
| **High** | [#5429](nearai/ironclaw Issue #5429) | Web Search always requires a NEAR AI Cloud API token, breaking zero‑config expectation. | No (see also #5421) |
| **High** | [#5426](nearai/ironclaw Issue #5426) | Cannot create a routine because “system drive is not available” on hosted‑staging. | No |
| **Medium** | [#5421](nearai/ironclaw Issue #5421) | Web search under `ironclaw-reborn serve` not zero‑config; bundled Exa MCP remains inactive. | No |
| **Medium** | [#4108](nearai/ironclaw Issue #4108) | Nightly E2E scheduled run failing repeatedly since May 27. | No |
| **Low** | [#5412](nearai/ironclaw Issue #5412) | Log entry text not selectable/copyable in WebUI v2 (closed). | Fixed |
| **Low** | [#5428](nearai/ironclaw Issue #5428) | Three pre‑existing defects in mock‑MCP test scaffolding (test‑only). | Work tracked |
| **Info** | [#5437](nearai/ironclaw Issue #5437) | Daily failure taxonomy reveals bulk HTTP 400 failures from model endpoint. | Not a code bug, but operational |

**Stability alerts:** The Nightly E2E runner (#4108) has been failing for over a month without a public fix. The daily failure taxonomy (#5437) is a new practice that suggests the team is actively cataloging – but not yet resolving – systemic CI breaks. No fix PRs exist yet for the critical/high-severity bugs above.

## 6. Feature Requests & Roadmap Signals
User-visible feature requests surfaced today:

- **[Issue #5443]** (nearai/ironclaw Issue #5443) – “Add header notifications for newly triggered automation tasks.” This is already implemented in **[PR #5441]** (open), which adds a notification bell with unread‑dot support. Likely to be merged in the next release.

- **[PR #5435]** (nearai/ironclaw PR #5435) – “feat(reborn): env-overridable compaction context budget (cut multi-turn re-send cost).” Though stacked on #5149 (progressive tool disclosure), this directly addresses long‑running tool-heavy tasks by reducing the 108k token blow‑up before compaction. Signals a focus on cost optimization and long‑session stability.

- **[PR #5444]** (nearai/ironclaw PR #5444) – “Add provider latency stress mode.” A developer‑facing tool, but indicates the team is preparing to benchmark and improve provider integration performance.

- **[PR #5313]** (nearai/ironclaw PR #5313) – “add storage stress harness” for filesystem‑backed resource governor stress testing. Suggests upcoming storage resiliency work.

**Predictions for next version:** The header notification feature (#5441) and the chat composer clearing fix (#5404) are small, self‑contained UX improvements likely to land soon. The compaction budget PR (#5435) may also be merged if the stacked tool‑disclosure work stabilizes. Larger items like run‑borking elimination (#4841) and the zero‑config web search fix (#5421, #5429) may be pushed to a subsequent milestone.

## 7. User Feedback Summary
User pain points expressed through issues (likely from QA or community):

- **Routine misrouting is frustrating:** A user reports that setting one routine to Slack overrides all other routines’ delivery targets. This is a **clear usability regression** for power users running multiple automations (e.g., email summary + Slack alerts).
- **Web search friction:** Users are surprised that chat works without authentication but web search demands a NEAR AI API key. The promise of zero‑config is broken, especially when the bundled Exa MCP exists but is inactive.
- **Routine creation blocked:** On hosted‑staging, users cannot create any routine because the system claims “system drive is not available.” This blocks a core workflow entirely.
- **Goodwill signals:** The ability to select/copy log text (#5412, now fixed) and the upcoming header notifications (#5443) are directly responsive to user feedback.

No explicit positive feedback was captured in the data, but the rapid closure of #5412 and the implementation of #5443 indicate the team is listening.

## 8. Backlog Watch
The following items require maintainer attention due to age or importance:

- **[Issue #4108](nearai/ironclaw Issue #4108)** – Nightly E2E failure (scheduled run failing since May 27, 2026). No fix PR linked. This blocks reliable CI signal and should be prioritized.
- **[PR #4841](nearai/ironclaw PR #4841)** – “reborn: no run-borking failures” (open since June 13, 2026). A fundamental reliability improvement that remains unmerged after 17 days. Its size (XL) may warrant breaking into smaller pieces.
- **[PR #5247](nearai/ironclaw PR #5247)** – “fix(webui): link approval card to global auto-approve settings” (open since June 25, 2026). A small UX fix (S/M) that has received no updates in five days.
- **[PR #5338](nearai/ironclaw PR #5338)** – “fix(reborn): surface real failure detail instead of generic ‘invalid_input’” (open since June 26, 2026). Addresses user‑visible error opacity, but is still open with no recent commits.
- **[Issue #5426](nearai/ironclaw Issue #5426)** – “Cannot create a routine: system drive is not available” (open since June 29, 2026). Critical for hosted‑staging users, but no PR yet.

The team appears to be prioritizing feature ports and QA infrastructure over long‑standing bugs and PRs. Without intervention, #4108 and #4841 risk becoming stale blockers.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest – 2026-06-30

---

## 1. Today's Overview

The project saw **high activity** over the last 24 hours: **16 PRs were updated**, with **13 merged or closed** (81% closure rate), and **1 new release** was published. **12 issues** were active, 2 of which were closed. The community continues to report both regressions and feature requests, while maintainers focused on stabilising `cowork` navigation rails and `openclaw` integration. The release **2026.6.29** was promoted to `main` today, consolidating several critical fixes. Overall project health is good, with a clear signal of rapid iteration.

---

## 2. Releases

**LobsterAI 2026.6.29** (released 2026-06-29)  
[Release tag](https://github.com/netease-youdao/LobsterAI/releases/tag/2026.6.29)

Notable changes (from changelog snippet and PR #2228 merge notes):
- **OpenClaw stability improvements**: plugin approvals now respect permissions; agent bootstrap workspace is kept separate from task working directory; cron run follow-up history is preserved; user turn cache stability improved.
- **Cowork fixes**: navigation rail tooltip previews are cleaned (filter thinking messages, plan tags, media artifacts); accidental revert of rail changes undone.
- **No breaking changes or migration notes** reported.

---

## 3. Project Progress

Today **13 PRs were merged/closed**, spanning four key areas:

| Area | Key Merges |
|------|------------|
| **Cowork (UI/UX)** | #2226 – reapply conversation rail fixes; #2223, #2222 – clean and align rail tooltips; #2218 – clean navigation rail previews; #2233 – remove prompt intent fields from analytics |
| **OpenClaw (agent runtime)** | #2232 – fallback catalog max token limits; #2227 – keep bootstrap workspace separate from task cwd |
| **Scheduled tasks** | #2231 – restore gateway-backed run history |
| **Logging & diagnostics** | #2229 – add diagnostic logging for Cowork and OpenClaw flows |
| **Notifications** | #1428 – feature: system notification on session complete/error (window unfocused) |

Three PRs remain **open**:
- #2234 – fix cron yield descendant finalization (critical for multi-agent workflows)
- #1372 – fix multiple file selection in cowoork input (stale since April)
- #1277 – Dependabot dependencies bump (stale)

---

## 4. Community Hot Topics

Most active discussions (by comments and novelty) include:

- **Issue #2120** (2 comments) – [Feature suggestions](https://github.com/netease-youdao/LobsterAI/issues/2120)  
  User requests task pre-input queue, longer single-task duration (to avoid `terminated`), and a 3-column skill grid. Underlying need: improved workflow continuity and UI density.

- **Issue #2121** (2 comments) – [Possible bug: repeated output consumes excessive tokens](https://github.com/netease-youdao/LobsterAI/issues/2121)  
  User observes duplicate text output and worries about token waste. Suggests a bug in Claw’s generation loop.

- **Issue #2079** (2 comments) – [Scrolling to top freezes execution result window](https://github.com/netease-youdao/LobsterAI/issues/2079)  
  Reproducible freeze in version 2026.5.27; no fix PR yet.

- **Issue #2230** (0 comments, created today) – [Same model runs 10x slower than CodeBuddy](https://github.com/netease-youdao/LobsterAI/issues/2230)  
  Serious performance regression claim: 2m24s vs 25 minutes, 60M vs 67k tokens. Likely to attract high attention.

- **Issue #2131** (2 comments) – [Request: support for Hermes agent](https://github.com/netease-youdao/LobsterAI/issues/2131)  
  Community interest in expanding agent backend options.

---

## 5. Bugs & Stability

Bugs reported or active today, ranked by severity:

| Severity | Issue | Description | Fix Exists? |
|----------|-------|-------------|-------------|
| **High** | [#2230](https://github.com/netease-youdao/LobsterAI/issues/2230) | Same model 10x slower than CodeBuddy (token count also disproportionate) | None yet |
| **High** | [#2079](https://github.com/netease-youdao/LobsterAI/issues/2079) | Execution result window freezes on scroll to top (reproduces) | None |
| **Medium** | [#2121](https://github.com/netease-youdao/LobsterAI/issues/2121) | Repeated output text causing token waste | None |
| **Medium** | [#1384](https://github.com/netease-youdao/LobsterAI/issues/1384) | Multiple file selection shows only last file (stale) | PR #1372 open but stale |
| **Medium** | [#1383](https://github.com/netease-youdao/LobsterAI/issues/1383) | WeChat bot: duplicate text sent only syncs once | None |
| **Medium** | [#1385](https://github.com/netease-youdao/LobsterAI/issues/1385) | WeChat bot: deleting session does not clear history for new questions | None |
| **Low** | [#1426](https://github.com/netease-youdao/LobsterAI/issues/1426) | No success feedback after skills upload; list not refreshed (closed) | Closed (fixed?) |
| **Low** | [#1427](https://github.com/netease-youdao/LobsterAI/issues/1427) | Duplicate skill addition allowed (closed) | Closed (fixed?) |

Notably, **no regression-related issues** were tied to the new 2026.6.29 release, suggesting the release was stable.

---

## 6. Feature Requests & Roadmap Signals

Top feature requests from active issues:

- **#2131** – [Hermes agent support](https://github.com/netease-youdao/LobsterAI/issues/2131) – high interest from power users exploring multi-agent frameworks.
- **#2120** – Pre-input tasks, extended run duration, 3-column skill grid – points to improving long-running script monitoring and UI polish.
- **#1381** – Option to show cron task results in the same conversation instead of new sessions – addresses token/session clutter.
- **#1382** – Change red-colored export log warnings (red usually implies errors) – minor UX request.

**Prediction for next release (2026.7.x):** Given the current maintenance focus on `openclaw` and `cowork`, the next version will likely prioritise fixing the **performance regression (#2230)** and the **scrolling freeze (#2079)**, while the Hermes agent request may be deferred pending architectural discussion.

---

## 7. User Feedback Summary

Real pain points expressed:

- **Performance gap**: A user reports LobsterAI being **10x slower** than a comparable tool (CodeBuddy) using the same model and prompt (#2230). Token consumption difference (60M vs 67k) suggests a possible loop or misconfiguration.
- **Token waste**: Duplicate output (#2121) frustrates users who pay for tokens.
- **Workflow friction**: Lack of task pre-input (#2120), new sessions for every cron job (#1381), and frozen scroll (#2079) disrupt productivity.
- **WeChat bot instability**: Messages not synced properly, history not cleared, multiple file uploads broken (#1383, #1384, #1385) – these are older but still unresolved.
- **Positive signal**: The addition of **system notifications** (PR #1428, merged today) shows the team listens to feedback about background sessions – a sentiment echoed in #2120’s request for longer run durations.

---

## 8. Backlog Watch

Issues and PRs requiring maintainer attention due to age or importance:

| Item | Age | Reason for Watch |
|------|-----|------------------|
| [#1381](https://github.com/netease-youdao/LobsterAI/issues/1381) – cron sessions reuse | 88 days | Stale but still open; affects daily cron users |
| [#1384](https://github.com/netease-youdao/LobsterAI/issues/1384) – multiple file upload bug | 88 days | Open; fix PR #1372 is also stale (90 days) |
| [#2079](https://github.com/netease-youdao/LobsterAI/issues/2079) – scrolling freeze | 31 days | Reproducible, no fix yet |
| [#1372](https://github.com/netease-youdao/LobsterAI/pull/1372) – fix multiple file selection (PR) | 89 days | Author no longer active? Needs review or reassignment |
| [#1277](https://github.com/netease-youdao/LobsterAI/pull/1277) – Dependabot (Electron bump) | 89 days | Stale; security updates may be blocked |
| [#2230](https://github.com/netease-youdao/LobsterAI/issues/2230) – extreme performance regression | <1 day | New but critical; needs immediate investigation |

---

*Data sourced from GitHub: [github.com/netease-youdao/LobsterAI](https://github.com/netease-youdao/LobsterAI). All times refer to UTC-12 to UTC+12 activity window from 2026-06-29 to 2026-06-30.*

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

# CoPaw Project Digest – 2026-06-30

## 1. Today's Overview

The CoPaw project saw very high activity on **2026-06-30**, with **15 issues updated** (11 open, 4 closed) and **50 pull requests updated** (26 open, 24 merged/closed). No new releases were published. The community is actively participating – several first-time contributors submitted PRs, and the **v2.0.0 pre-release bug tracker** (#5273) continues to collect feedback. The majority of activity focuses on **bug fixes for DeepSeek V4 compatibility**, **DingTalk channel enhancements**, and **frontend stability improvements**. The maintainers are also advancing major features such as **Windows native sandboxing**, **desktop computer-use tools**, and **plugin market version filtering**.

---

## 2. Releases

**No new releases** were published today. The latest available version remains **v1.1.12.post2** (as referenced in multiple bug reports).

---

## 3. Project Progress

**Merged/closed PRs today (24 total)** – key highlights:

| PR | Title | Summary |
|----|-------|---------|
| [#5664](https://github.com/agentscope-ai/QwenPaw/pull/5664) | feat(chat): add non-owner tab info banner | Shows an Alert banner when the current tab is not the owner, preventing confusion in multi-tab sessions. |
| [#5662](https://github.com/agentscope-ai/QwenPaw/pull/5662) | fix(ci): modify channel name in PR template | CI housekeeping – minimal code change. |
| [#5590](https://github.com/agentscope-ai/QwenPaw/pull/5590) | feat(channels): support dingtalk mentions in proactive sends | Adds `@mention` support for DingTalk proactive messages (CLI, API, cron). Closes #5564. |
| [#5655](https://github.com/agentscope-ai/QwenPaw/pull/5655) | docs(readme): update and refine readme | Documentation improvement. |
| [#5656](https://github.com/agentscope-ai/QwenPaw/pull/5656) | fix(layouts): isolate sidebar session list scrolling to simple mode | UI bug fix for sidebar scrolling. |
| [#5639](https://github.com/agentscope-ai/QwenPaw/pull/5639) | feat(skill): Add skill auto sync | Skills from the pool now automatically synchronise to selected agents; sync history recorded in inbox. |
| [#4041](https://github.com/agentscope-ai/QwenPaw/pull/4041) | feat(desktop): add Tauri tray behavior | Native system tray integration for Windows/macOS (merged after earlier rework). |

**Also closed today (issues)**:
- [#5401](https://github.com/agentscope-ai/QwenPaw/issues/5401) (frontend crash on large tool-use history) – fixed and closed.
- [#5573](https://github.com/agentscope-ai/QwenPaw/issues/5573) (DeepSeek V4 400 errors) – closed with fix accepted.
- [#5564](https://github.com/agentscope-ai/QwenPaw/issues/5564) (DingTalk @mention) – resolved by PR #5590.

**Open PRs advancing features** (no maintainer action today but updated):
- [#5525](https://github.com/agentscope-ai/QwenPaw/pull/5525) – **Windows native sandbox** (first-time contributor).
- [#5187](https://github.com/agentscope-ai/QwenPaw/pull/5187) – **Windows desktop computer-use (UIA + Tauri)**.
- [#5069](https://github.com/agentscope-ai/QwenPaw/pull/5069) – **Visual model fallback** for text-only primary LLMs.
- [#5652](https://github.com/agentscope-ai/QwenPaw/pull/5652) – **Per-request model override in cron jobs**.
- [#5659](https://github.com/agentscope-ai/QwenPaw/pull/5659) – **Allow sending attachments without text**.
- [#5651](https://github.com/agentscope-ai/QwenPaw/pull/5651) – **Custom Telegram BaseURL**.

---

## 4. Community Hot Topics

**Most active issues** (by comment count, last 24h):

| Issue | Comments | Topic |
|-------|----------|-------|
| [#5401](https://github.com/agentscope-ai/QwenPaw/issues/5401) | 6 | Frontend crash on large tool-use history (closed today) |
| [#5403](https://github.com/agentscope-ai/QwenPaw/issues/5403) | 5 | Browser autofill hijacks search input in Model Configuration page |
| [#5588](https://github.com/agentscope-ai/QwenPaw/issues/5588) | 4 | Feature: two-stage memory search with dedicated reranker |
| [#5573](https://github.com/agentscope-ai/QwenPaw/issues/5573) | 4 | DeepSeek V4 400 errors with streaming reasoning_content and schema null (closed) |
| [#5561](https://github.com/agentscope-ai/QwenPaw/issues/5561) | 4 | Feishu bot fails to deliver long messages; forced to send as file |
| [#5564](https://github.com/agentscope-ai/QwenPaw/issues/5564) | 3 | DingTalk @mention support (closed, merged in #5590) |
| [#5630](https://github.com/agentscope-ai/QwenPaw/issues/5630) | 2 | Custom BaseURL for Telegram channel |

**Underlying needs**:  
- **Frontend performance** with large conversation history (especially tool-use traces) remains a pain point, now fixed in #5401.  
- **Channel integration completeness**: users need DingTalk `@mention`, Telegram custom endpoints, Feishu long-message handling, and DingTalk streaming speed improvements.  
- **Memory search precision**: single-stage embedding retrieval is insufficient as memory grows; a reranker stage is requested.  
- **Browser UX friction**: autofill interference on configuration pages.

---

## 5. Bugs & Stability

**Bugs reported or updated today**, ranked by severity:

| Severity | Issue | Description | Fix PR Exists? |
|----------|-------|-------------|----------------|
| 🔴 High | [#5658](https://github.com/agentscope-ai/QwenPaw/issues/5658) | **Cannot connect to 9router for QwenPaw models** – persistent `[400]` error since earlier versions. | No |
| 🔴 High | [#5587](https://github.com/agentscope-ai/QwenPaw/issues/5587) | **Qwen-Image Tool install error** – installation fails. | No |
| 🟡 Medium | [#5561](https://github.com/agentscope-ai/QwenPaw/issues/5561) | **Feishu bot cannot send long replies** – forced to send as file, poor UX. | No |
| 🟡 Medium | [#5657](https://github.com/agentscope-ai/QwenPaw/issues/5657) | **Agent loops with Qwen3.6 models** – loop detection mechanism requested (enhancement but labeled bug). | No |
| 🟡 Medium | [#5603](https://github.com/agentscope-ai/QwenPaw/issues/5603) | **DingTalk card streaming output very slow** – character-by-character rendering. | No |
| 🟢 Low | [#5403](https://github.com/agentscope-ai/QwenPaw/issues/5403) | **Browser autofill interference** on configuration search inputs. | No |
| 🟢 Low | [#5566](https://github.com/agentscope-ai/QwenPaw/issues/5566) | **Cron tasks can't be silent** & `channels send` unreachable from scripts (partially addressed by #5654). | #5654 (open, under review) |
| 🟢 Low | [#5663](https://github.com/agentscope-ai/QwenPaw/issues/5663) | **Question**: add toggle to bypass debounce when sending media-only messages. | #5659 (open, first-time contributor) |

**Key fix PRs under review**:  
- [#5654](https://github.com/agentscope-ai/QwenPaw/pull/5654) – surfaces DingTalk delivery failures and skips empty (silent) cron messages.  
- [#5660](https://github.com/agentscope-ai/QwenPaw/pull/5660) – restores `spawn_subagent` for Runtime 2.0 (fixes four regressions).

---

## 6. Feature Requests & Roadmap Signals

**User-requested features** (with strong community interest):

| Issue/PR | Request | Likely in Next Version? |
|----------|---------|-------------------------|
| [#5588](https://github.com/agentscope-ai/QwenPaw/issues/5588) | Two-stage memory search with reranker | Medium – popular enhancement |
| [#5630](https://github.com/agentscope-ai/QwenPaw/issues/5630) | Custom Telegram BaseURL | High – PR #5651 submitted |
| [#5603](https://github.com/agentscope-ai/QwenPaw/issues/5603) | Faster DingTalk card streaming | Medium – user frustration |
| [#5657](https://github.com/agentscope-ai/QwenPaw/issues/5657) | Loop detection mechanism | Medium – linked to model quality |
| [#5663](https://github.com/agentscope-ai/QwenPaw/issues/5663) | Debounce bypass for media-only messages | High – PR #5659 submitted |
| [#5566](https://github.com/agentscope-ai/QwenPaw/issues/5566) | Silent cron execution | Medium – #5654 addresses part |
| [#5525](https://github.com/agentscope-ai/QwenPaw/pull/5525) | Windows native sandbox | High – active PR, likely in v2.0 |
| [#5187](https://github.com/agentscope-ai/QwenPaw/pull/5187) | Windows desktop computer-use (UIA + Tauri) | High – large feature in development |
| [#5069](https://github.com/agentscope-ai/QwenPaw/pull/5069) | Visual model fallback for text-only LLMs | Medium – now open for 20 days |
| [#5639](https://github.com/agentscope-ai/QwenPaw/pull/5639) | Skill auto sync | **Already merged today** |
| [#5661](https://github.com/agentscope-ai/QwenPaw/pull/5661) | Plugin market version compatibility filtering | Under review – important for plugin ecosystem |

**Prediction**: The next minor release (v1.2.x or v2.0.0 alpha) will likely include: Telegram custom BaseURL, media-only send, DingTalk @mention (already merged), cron model override (PR #5652), skill auto sync, and possibly Windows sandbox.

---

## 7. User Feedback Summary

**Pain points expressed by real users**:
- **Frontend crashes** on large tool-use histories (#5401) – fixed but shows need for better rendering scalability.
- **DeepSeek V4 compatibility** (#5573) – users rely on third-party/compatibility endpoints for DeepSeek V4; issues with reasoning content and null schemas.
- **Feishu long message truncation** (#5561) – users forced to send as file; no workaround.
- **9router connectivity failures** (#5658) – persistent bug across versions, no response from maintainers yet.
- **DingTalk streaming speed** (#5603) – “typing animation” effect is too slow, impacting daily use.
- **Browser autofill** (#5403) – small annoyance but reported by multiple users.
- **Cron silent execution** (#5566) – users need background monitoring without notification spam.

**Positive signals**:  
- First-time contributors are submitting quality PRs (e.g., Windows sandbox, Telegram custom URL, attachment-only send).  
- The v2.0.0 pre-release tracker (#5273) has only 1 👍 but is being actively updated – indicates the community is engaged in testing.  
- The **DingTalk @mention** feature was requested, discussed, and merged within 4 days – responsive development.

**Satisfaction**: Mixed. Users appreciate rapid bug fixes (e.g., #5401 closed today) but are frustrated by lingering issues (Feishu, 9router, streaming speed). No explicit praise or dissatisfaction comments in the sample.

---

## 8. Backlog Watch

**Issues/PRs that need maintainer attention** – unanswered, stagnant, or lacking triage:

| Item | Created | Status | Concern |
|------|---------|--------|---------|
| [#5616](https://github.com/agentscope-ai/QwenPaw/issues/5616) | 2026-06-29 | Open, no reply | User reports automated tasks terminating spontaneously with no intervention – serious stability issue, no diagnosis. |
| [#5658](https://github.com/agentscope-ai/QwenPaw/issues/5658) | 2026-06-30 | Open, no reply | 9router connectivity failure – user has experienced this across versions, no maintainer comment. |
| [#5657](https://github.com/agentscope-ai/QwenPaw/issues/5657) | 2026-06-30 | Open, no reply | Loop detection request – labelled enhancement but tied to model bugs; no triage. |
| [#5587](https://github.com/agentscope-ai/QwenPaw/issues/5587) | 2026-06-28 | Open, no reply | Qwen-Image Tool install error – possibly a packaging issue. |
| [#5561](https://github.com/agentscope-ai/QwenPaw/issues/5561) | 2026-06-26 | Open, no reply | Feishu long-message failure – no response or workaround. |
| [#5525](https://github.com/agentscope-ai/QwenPaw/pull/5525) | 2026-06-25 | Open, no review | Windows native sandbox – first-time contributor PR with no maintainer feedback. May need code review. |
| [#5187](https://github.com/agentscope-ai/QwenPaw/pull/5187) | 2026-06-14 | Open, no review | Windows computer-use feature – large, important PR, waiting for review. |
| [#5069](https://github.com/agentscope-ai/QwenPaw/pull/5069) | 2026-06-10 | Open, no review | Visual model fallback – open for 20 days, no maintainer comments. |
| [#5151](https://github.com/agentscope-ai/QwenPaw/pull/5151) | 2026-06-12 | Open, no review | GitPanel CSS prefix fix – minor but ready. |

**Maintainer responsiveness note**: While many bugs were fixed today, several open issues (especially #5616, #5658, #5561) have no maintainer reply despite being reported days ago. The PR backlog for features like Windows sandbox and computer-use is growing without review. This may slow down the v2.0.0 release timeline.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

## ZeroClaw Project Digest – 2026-06-30

### 1. Today's Overview
ZeroClaw remains in an intense development sprint toward **v0.8.3**, with **50 pull requests updated in the last 24 hours** (41 still open) and **8 issues touched**. Activity is dominated by large, risk‑high trackers that systematically break down runtime, provider, channel, and config work. Nine PRs were merged or closed today, including the fix for a long‑standing bug where the `session_end` hook was never fired. No new releases were cut, but the sheer volume of in‑flight PRs signals that a v0.8.3 release is being assembled.

### 2. Releases
**None** – no new versions were published in the last 24 hours.

### 3. Project Progress (Merged/Closed PRs Today)
- **#8003** [fix(runtime): fire session_end hook on session termination] – Wires the previously dead `on_session_end` hook into session lifecycle paths, closing issue #7889.  
  *Link:* [PR #8003](https://github.com/zeroclaw-labs/zeroclaw/pull/8003)
- **#8148** [fix(anthropic): propagate serialization error in streaming request builder] – Replaces an `expect()` that could panic with proper error propagation.  
  *Link:* [PR #8148](https://github.com/zeroclaw-labs/zeroclaw/pull/8148)
- **#8458** [test(eval): cover RecordingObserver tool recording and token accumulation] – Adds test coverage for the eval observer, no production changes.  
  *Link:* [PR #8458](https://github.com/zeroclaw-labs/zeroclaw/pull/8458)

### 4. Community Hot Topics
None of today’s updated issues or PRs show more than **2 comments**, indicating low‑volume discussion. The most active item is:

- **#8251** [Feature: Surface relationship memory as user-facing workflows] (2 comments) – This enhancement request proposes exposing ZeroClaw’s knowledge‑graph relationship tools as documented operator workflows, building on the restored typed‑relationship functionality from #8182.  
  *Link:* [Issue #8251](https://github.com/zeroclaw-labs/zeroclaw/issues/8251)

The trackers (#8071, #8360, #8362, #8073, #8070, #8363) are all created by the same maintainer (Audacity88) and serve as internal roadmaps; no external community discussion is visible.

### 5. Bugs & Stability
Only one bug issue was updated today, and it was **closed**:

- **#7889** [Bug: session_end hook is defined but never fired] (severity S3 – minor) – Fixed by PR #8003 (merged).  
  *Link:* [Issue #7889](https://github.com/zeroclaw-labs/zeroclaw/issues/7889)

Several open bugs have active fix PRs in the pipeline:

| Bug | Risk | Fix PR | Status |
|-----|------|--------|--------|
| **#8465** [fix(cron): thread CancellationToken for explicit shutdown] | high | #8465 (open) | PR submitted |
| **#8496** [fix(tools/mcp): centralize deferred-MCP access policy] | high | #8496 (open, needs-author-action) | PR submitted |
| **#8477** [fix(zerocode): let active sessions switch agents] | medium | #8477 (open) | PR submitted |
| **#8535** [fix(runtime): gate Unix-only shell test helper] | medium | #8535 (open) | PR submitted |
| **#7637** [fix(zerocode): auto-normalise agent alias input] | low | #7637 (open, stale-candidate) | PR submitted |
| **#7535** [feat(whatsapp-web): implement add_reaction/remove_reaction] | medium | #7535 (open, stale-candidate) | PR submitted |

**No crashes or regressions were reported today.**

### 6. Feature Requests & Roadmap Signals
Today’s issues are almost entirely **internal trackers** (v0.8.3 scope) rather than user feature requests. One explicit user‑facing enhancement is **#8251** (relationship memory workflows).  

New features landing via open PRs point strongly to the **v0.8.3** feature set:

- **#8384** – Native Inkbox channel (email + SMS + voice + iMessage).  
- **#8443** – Matrix single‑message streaming drafts.  
- **#8504** – Git forge (GitHub App) channel with SOP ingress.  
- **#7535** – WhatsApp reaction parity.  
- **#8427** – WhatsApp location pin support.  
- **#8139** – Channel session TTL cleanup.  
- **#8440** – Telegram per‑channel inbound debounce.  

**Prediction for next version:** v0.8.3 will likely include all the above channels, the MCP access‑policy centralization (#8496), cron shutdown fixes, and the session_end hook wiring.

### 7. User Feedback Summary
No direct user feedback (e.g., 👍 counts, user comments) is available in the dataset beyond the two comments on issue #8251. That issue explicitly asks for **documented operator workflows** around relationship memory, indicating that users want to leverage ZeroClaw’s knowledge‑graph capabilities in a guided, reproducible way.

The flurry of bug‑fix PRs suggests users may have encountered:

- Dead session hooks (fixed).
- Panics in Anthropic streaming (fixed).
- MCP tool access misconfiguration (ongoing fix).
- Cron scheduler not shutting down cleanly (ongoing fix).
- Zerocode quickstart alias input quirks (ongoing fix).

Overall, the project is responding quickly to issues, with most bugs having a corresponding PR within days.

### 8. Backlog Watch
Two PRs remain **stale candidates with “needs‑author‑action”** – they have not been updated in over two weeks and may require maintainer ping:

- **#7637** (zerocode alias normalisation) – last updated June 14, risk low  
  *Link:* [PR #7637](https://github.com/zeroclaw-labs/zeroclaw/pull/7637)
- **#7535** (WhatsApp reactions) – last updated June 12, risk medium  
  *Link:* [PR #7535](https://github.com/zeroclaw-labs/zeroclaw/pull/7535)

A third PR, **#8496** (MCP access policy), also carries the `needs-author-action` label but was opened only yesterday and is not yet stale.

No old, unanswered issues without any maintainer response were identified in today’s update set.

---

*Generated from ZeroClaw GitHub data for 2026-06-30.*

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*