# OpenClaw Ecosystem Digest 2026-06-21

> Issues: 124 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-21 11:26 UTC

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

# OpenClaw Project Digest — 2026-06-21

## 1. Today's Overview
OpenClaw sees extremely high activity with **124 issues** and **500 pull requests** updated in the last 24 hours, of which **4 issues were closed** and **14 PRs merged/closed**. Two beta/stable releases dropped today, targeting Telegram delivery quality and agent turn reliability. However, the sheer volume of open items (120 open issues, 486 open PRs) signals that the project is struggling to keep pace with incoming reports—many carry `clawsweeper:needs-maintainer-review` or wait on product decisions. A **P1 regression** (#95495) silently relocating the memory store without migration was reported today, underscoring upgrade risk.

## 2. Releases
Two releases published today:

- **v2026.6.10-beta.1** — Focuses on agent turn and session state stability: pending subagent completion announcements are now preserved, chat history transcripts stay non-empty, media index alignment is maintained, dormant follow-up drains restart properly, and compaction model aliases are resolved consistently.
- **v2026.6.9** — Delivers richer Telegram delivery: HTML, markdown, sticker paths, progress drafts, command output, normalized HTML tables, and correct mention/spooled handler routing. No breaking changes or migration notes were provided for either release.

## 3. Project Progress
14 PRs were merged or closed today. Key merges include:

- [PR #95413](https://github.com/openclaw/openclaw/pull/95413) (closed) — Fix Telegram persistent rich message line breaks (fixes #95409)
- [PR #95037](https://github.com/openclaw/openclaw/pull/95037) (closed) — Convert newlines to `<br>` in Telegram rich HTML
- [PR #95532](https://github.com/openclaw/openclaw/pull/95532) (closed) — Materialize rich message line breaks as `<br>` in Telegram
- [PR #68936](https://github.com/openclaw/openclaw/pull/68936) (closed) — Add PR review autofix pipeline + Windows daemon

These signal steady progress on Telegram formatting and automation infrastructure, but the volume of pending PRs suggests a growing maintenance backlog.

## 4. Community Hot Topics
Most active issues by comment count (all updated today):

| Issue | Comments | Summary |
|-------|----------|---------|
| [#86215](https://github.com/openclaw/openclaw/issues/86215) (P1) | 9 | Codex OAuth refresh failures can wedge an agent for hours without alerting |
| [#76935](https://github.com/openclaw/openclaw/issues/76935) (P2) | 7 | QQ Bot sends verbose/repetitive replies after upgrading to 2026.5.2 |
| [#89315](https://github.com/openclaw/openclaw/issues/89315) (P1) | 6 | Gateway heap grows unbounded, gets killed by cgroup OOM on long-running Linux |
| [#95495](https://github.com/openclaw/openclaw/issues/95495) (P1) | 5 | 2026.6.9 silently relocates memory store with no migration, forcing full re-embed |
| [#80176](https://github.com/openclaw/openclaw/issues/80176) (P3) | 5 | JSONL session-replay harness for testing and diffing trajectories |

**Underlying needs:** Users are demanding **reliable OAuth handling**, **no silent data migrations**, **better memory management** for long-running deployments, and **testing infrastructure** to prevent regressions. The QQ Bot regression (#76935) is a recurring pain point for channels with high user volume.

## 5. Bugs & Stability
Today's newly reported bugs (creation date 2026-06-21):

| Issue | Severity | Summary | Fix PR Exists? |
|-------|----------|---------|----------------|
| [#95495](https://github.com/openclaw/openclaw/issues/95495) | **P1 / Diamond Lobster** | 2026.6.9 silently moves memory store from `~/.openclaw/memory/main.sqlite` to `~/.openclaw/agents/main/agent/openclaw-agent.sqlite` with no migration, forcing a full 1499-file re-embed on upgrade. | **No** |
| (Referenced) [#95409](https://github.com/openclaw/openclaw/issues/95409) | P2 | Telegram persistent rich message line breaks collapse (fixed in #95413) | Yes (merged) |
| (Referenced) [#95519](https://github.com/openclaw/openclaw/issues/95519) | P2 | Fallback not triggered on provider upstream_error | Pending in #95524 & #95542 |
| (Referenced) [#95474](https://github.com/openclaw/openclaw/issues/95474) | P2 | missing_tool_result classified unclassified, triggers cross-provider fallback incorrectly | Pending in #95543 |

Additionally, several high-severity preexisting issues were updated today:
- **P1** [#86215](https://github.com/openclaw/openclaw/issues/86215) (OAuth wedging), **P1** [#89315](https://github.com/openclaw/openclaw/issues/89315) (OOM heap growth), **P1** [#88562](https://github.com/openclaw/openclaw/issues/88562) (models.json writes apiKey as plain string), **P1** [#89473](https://github.com/openclaw/openclaw/issues/89473) (reasoning tokens leak to chat channels).

**Takeaway:** The memory store relocation bug (#95495) is especially concerning—it was introduced in the stable 2026.6.9 release and lacks an existing fix. Users should pause upgrades until a patch is available.

## 6. Feature Requests & Roadmap Signals
Notable feature requests active today:

- [#89876](https://github.com/openclaw/openclaw/issues/89876) — Telegram decision packet cards for bounded agent input (P2)
- [#89274](https://github.com/openclaw/openclaw/issues/89274) — Optionally show model and thinking level after `/new` or `/reset` (P3)
- [#89339](https://github.com/openclaw/openclaw/issues/89339) — New `openclaw sessions diagnose` subcommand with failure summary and fix suggestions (P3)
- [#87023](https://github.com/openclaw/openclaw/issues/87023) — First-class relay of MCP ImageContent to channel extensions (iMessage, Telegram, Signal) (P2)
- [#80176](https://github.com/openclaw/openclaw/issues/80176) — JSONL session-replay harness for testing (P3)

**Roadmap signal:** The project is investing in **Telegram UX** (decision packets), **operator tooling** (sessions diagnose), and **multimodal channel relay** (MCP images). The session-replay harness (#80176) is a Phase 5 task of a larger parity initiative and could land within 1–2 releases. The model display tweak (#89274) is small and likely to be included in an upcoming patch.

## 7. User Feedback Summary
**Pain points expressed today:**
- **Silent data migration (#95495):** “After upgrading from 2026.6.8 to 2026.6.9, openclaw memory status –deep reported the memory vector store relocated … with zero upgrade-time warning.” — 5 comments, 1 👍
- **OAuth unreliability (#86215):** “Codex OAuth refresh can keep retrying inside the same provider/auth lane for hours without surfacing a clear operator-visible incident.” — 9 comments, 1 👍
- **QQ Bot regression (#76935):** “Same LLM model (deepseek-v4-flash) produced concise, single-message replies on 2026.4.29, so this is not a model change.” — 7 comments, 1 👍
- **Heap OOM (#89315):** “Gateway heap grows unbounded over time, gets killed by cgroup OOM on long-running Linux systemd –user deployments.” — 6 comments, 3 👍 (most upvoted this period)
- **Subagent silent drops (#88856):** “Agent / sessions_spawn tool_use can emit with no matching tool_result and no parent-visible signal — silent subagent drop (~3.8% historical rate).” — 4 comments, 1 👍

**Satisfaction signals:** The Telegram release (v2026.6.9) addresses long-standing formatting issues; users are actively testing and reporting remaining gaps. The beta release with better agent turn reliability is welcomed but the memory regression undermines confidence.

## 8. Backlog Watch
Long‑unanswered issues and PRs needing maintainer attention:

| Item | Date Created | Last Maintainer Action |
|------|--------------|------------------------|
| [#76935](https://github.com/openclaw/openclaw/issues/76935) — QQ Bot verbose replies (P2) | 2026-05-03 | `clawsweeper:needs-maintainer-review` since May 3 |
| [#80176](https://github.com/openclaw/openclaw/issues/80176) — JSONL session-replay harness (P3) | 2026-05-10 | No maintainer comment |
| [#86215](https://github.com/openclaw/openclaw/issues/86215) — Codex OAuth wedging (P1) | 2026-05-24 | Needs product decision & maintainer review |
| [#87689](https://github.com/openclaw/openclaw/issues/87689) — Dreaming guard for QMD migrations (P2) | 2026-05-28 | Needs maintainer review & product decision |
| [PR #49488](https://github.com/openclaw/openclaw/pull/49488) — Android minSdk 26 (P2) | 2026-03-18 | Awaiting maintainer review for 3 months |
| [PR #74564](https://github.com/openclaw/openclaw/pull/74564) — Use agent auth scope in /models (P2) | 2026-04-29 | `status: needs maintainer proof decision` |
| [PR #75148](https://github.com/openclaw/openclaw/pull/75148) — Fallback path user-visible bug (P1) | 2026-04-30 | `status: needs maintainer proof decision` |

The sheer number of P1/P2 issues and PRs awaiting maintainer review or product decisions indicates a bottleneck. The ***diamond lobster***‑rated issues (e.g., #86215, #88562, #89100, #89473) represent critical risks to production deployments and merit immediate prioritization. The project would benefit from dedicated maintainer time or a community triage rotation to clear the queue.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: AI Agent Open-Source Ecosystem
**Date:** 2026-06-21

---

## 1. Ecosystem Overview

The open-source personal AI assistant landscape remains vibrant but fragmented, with at least a dozen active projects competing for developer mindshare. Activity levels vary widely—from high-velocity, well-resourced projects (NanoBot, IronClaw, ZeroClaw) to near-dormant repositories (NullClaw, TinyClaw, ZeptoClaw). A common theme is the tension between feature velocity and stability: projects with high PR throughput (OpenClaw, Hermes Agent, CoPaw) often accumulate significant maintenance backlogs and regressions, while more focused projects (NanoBot, NanoClaw) enjoy faster fix cycles. The ecosystem is converging on a shared set of pain points—reliable OAuth, memory management, mobile UX, and multi-channel support—but each project takes a distinct architectural path to address them.

---

## 2. Activity Comparison (24h ending 2026-06-21)

| Project | Issues Updated | PRs Updated | Merged/Closed PRs | Releases Today | Health Score |
|---------|---------------|-------------|-------------------|----------------|--------------|
| **OpenClaw** | 124 | 500 | 14 | 2 (beta + stable) | Medium |
| **NanoBot** | 2 | 12 | 12 | 0 | High |
| **Hermes Agent** | 9 | 50 | 12 | 0 | Medium-High |
| **PicoClaw** | 2 | 1 | 0 | 1 (nightly) | Low |
| **NanoClaw** | 1 | 4 | 0 | 0 | Low-Medium |
| **NullClaw** | 0 | 0 | 0 | 0 | Inactive |
| **IronClaw** | 2 | 20 | 9 | 0 | High |
| **LobsterAI** | 14 (all stale) | 0 | 0 | 0 | Low |
| **TinyClaw** | 0 | 0 | 0 | 0 | Inactive |
| **Moltis** | 0 | 2 | 1 | 0 | Low |
| **CoPaw** | 13 | 18 | 0 | 0 | Medium |
| **ZeptoClaw** | 0 | 0 | 0 | 0 | Inactive |
| **ZeroClaw** | 6 | 50 | 13 | 0 | High |

**Health Score Criteria:**
- **High:** Rapid fix cycles, low open-to-closed ratio, responsive maintainers.
- **Medium:** Active but with notable backlog or regressions.
- **Low:** Minimal activity, stale issues, or maintenance-only.
- **Inactive:** Zero activity in 24h.

---

## 3. OpenClaw’s Position in the Ecosystem

### Advantages
- **Scale & Community:** Largest raw activity (124 issues, 500 PRs) – 10× more than any peer. Two releases in one day (beta + stable) demonstrate continuous delivery.
- **Feature Breadth:** Telegram rich formatting, agent turn reliability, memory compaction – covers more channels and deployment scenarios than any other project.
- **Reference Implementation:** Serves as the core downstream base for several forks (PicoClaw, NanoClaw).

### Weaknesses
- **Maintenance Bottleneck:** 120 open issues + 486 open PRs → the team is overwhelmed. Several P1 **diamond lobster** issues (OAuth wedging, heap OOM, silent memory store migration) remain unpatched.
- **Regression Risk:** The v2026.6.9 release introduced a silent memory store relocation (#95495) with no migration path, undermining trust in stable releases.
- **Slower Bug Fixes:** While NanoBot fixed a concurrency bug within 24 hours, OpenClaw’s P1 OAuth bug (#86215) has been waiting for product decision since May 24.

### Technical Approach Differences
- **Architecture:** Monolithic core reference – tries to do everything in one codebase (gateway, CLI, memory, multiple channels). Peers like ZeroClaw split into hardware features, test harnesses, and separate RFC processes.
- **Community Size:** Likely the largest user base (judging by issue volume and comment counts), but also the most fragmented feedback signals.

---

## 4. Shared Technical Focus Areas

The following requirements appear across multiple projects, indicating ecosystem-wide priorities:

| Need | Projects Affected | Specific Pain Points |
|------|-------------------|----------------------|
| **Robust OAuth & Token Management** | OpenClaw (#86215), IronClaw (#5071), ZeroClaw (Codex mismatch in #8030), Hermes (OAuth tests) | Silent refresh failures, forced re-authentication, provider credential mismatches |
| **Memory & Data Migration Safety** | OpenClaw (#95495), NanoBot (#4256, cursor monotonic), ZeroClaw (#7942, embedding API key decoupling), CoPaw (#5342, unpruned tool results) | Silent data loss, missing migration paths, context window pollution |
| **Mobile & Responsive UI** | Hermes (#49834 Android draft), CoPaw (#5329, #5355 mobile dropdown), NanoBot (iOS Safari fix #4427) | Agent switching on mobile, dropdowns obscured, high CPU/heat on desktop |
| **Channel Diversity & Reliability** | OpenClaw (Telegram, QQ), NanoBot (iMessage, WhatsApp), Hermes (DingTalk, Telegram), CoPaw (message queue cross-talk #5354) | Formatting issues, reconnect loops, cross-conversation message leaks |
| **Cost & Performance Optimization** | NanoBot (#4431 heartbeat override, #4420 prompt caching), NanoClaw (#2768 prompt caching), Hermes (#50107 STREAM_BATCH_MS) | Expensive re-encoding, unbounded token consumption, high GPU/CPU usage |
| **Testing & Diagnostics Infrastructure** | OpenClaw (#80176 session-replay), ZeroClaw (#7916 regression tests, #8029 doctor), Hermes (#50106 silent message loss) | Missing regression coverage, silent failures, poor error messages |

---

## 5. Differentiation Analysis

| Project | Primary Focus | Target Users | Technical Architecture |
|---------|---------------|--------------|------------------------|
| **OpenClaw** | Omnichannel personal assistant (Telegram, QQ, etc.) with rich features | Power users, multi-platform deployers | Monolithic core with plugin channels; memory store SQLite/vector; heavy CLI |
| **NanoBot** | Lightweight, performant SDK for embedded agents | Developers building custom assistants | Modular Python SDK; fast iteration; community-driven channels (iMessage, WhatsApp) |
| **Hermes Agent** | Enterprise-grade agent with DingTalk, MCP, and security | Enterprise/org deployments, Chinese market | Gateway+CLI+plugin architecture; P1 security PR (#6335 unauthenticated webhooks) |
| **PicoClaw** | Embedded/edge agent (Sipeed hardware) | Hardware hackers, IoT enthusiasts | Lightweight, minimal; nightly builds; vision pipeline |
| **NanoClaw** | Minimalistic reference for learning/forks | Students, tinkerers | Clean, small codebase; few PRs; clear maintenance |
| **IronClaw** | Declarative manifest-driven channels + workspace sharing | Collaborative teams, multi-agent setups | Rust/WebAssembly; concurrent turn execution; manifest-projected ingress |
| **LobsterAI** | Chinese IM bot integration (WeChat, QQ, DingTalk) | Chinese enterprise users | Bot-focused; many stale bugs closed without fixing; stable but neglected |
| **CoPaw** | Community-driven, mobile-first, multi-provider | General developers, mobile-first users | High community contribution rate; bottlenecks in review; first-timer friendly |
| **ZeroClaw** | Robust, well-typed agent runtime with strong diagnostics | Developers needing reliable production agents | Rust; RFC-driven governance; hardware dep feature-gated; active bug triage |

---

## 6. Community Momentum & Maturity

### Tier 1 – Rapidly Iterating (High Velocity + Responsive Maintainers)
- **NanoBot, IronClaw, ZeroClaw** – All three merge multiple PRs daily, address bugs within hours, and maintain low open/closed ratios. IronClaw stands out for fixing a Google OAuth bug and consolidating manifest channels in one day. ZeroClaw closed an S1 bug in under 3 days.
- **OpenClaw** – Despite highest volume, the sheer backlog and silent regressions put it at risk of losing community trust. Momentum is high but chaotic.

### Tier 2 – Moderate Activity (Feature Work + Some Backlog)
- **Hermes Agent** – Heavy PR activity but notable security PR (#6335) open for over 2 months indicates governance gaps. Many P2 bugs remain unpatched.
- **CoPaw** – Thriving community (many first-time contributors) but zero PRs merged today; review bottleneck could stifle momentum.

### Tier 3 – Stabilizing / Maintenance Mode
- **LobsterAI, Moltis** – Both saw only automated or dependency updates. LobsterAI closed 14 open bugs without fixing them – a sign of triage rather than development.
- **PicoClaw, NanoClaw** – Very low activity; maintainer attention appears minimal. Risk of becoming inert.

### Tier 4 – Inactive
- **NullClaw, TinyClaw, ZeptoClaw** – No activity for 24h. Likely abandoned or waiting for maintainers.

### Maturity Observations
- **Security Response:** Only Hermes Agent and IronClaw explicitly tag security PRs (P1). OpenClaw’s `clawsweeper:needs-maintainer-review` label suggests a triage system, but long wait times reduce effectiveness.
- **Documentation:** ZeroClaw’s PR #6870 (feature support matrix) and NanoBot’s README updates indicate efforts to improve discoverability. Most projects lack clear roadmaps.
- **Governance:** ZeroClaw’s RFC process for retiring crates is the most structured. CoPaw and IronClaw show signs of community-driven decision-making.

---

## 7. Trend Signals

The following industry trends emerge from community feedback across projects, with actionable insights for AI agent developers:

### 1. **The "Last Mile" of Mobile & Multimodal UX**
Users are actively pushing desktop-only agents onto mobile browsers (CoPaw, Hermes). Combined with demand for iMessage (NanoBot), Telegram decision cards (OpenClaw), and MCP ImageContent relay (OpenClaw #87023), the ecosystem is moving toward **pervasive, multi-device interaction**. Developers should prioritize responsive web UIs and lightweight sidecar patterns.

### 2. **OAuth Fatigue & Credential Hygiene**
Four projects report OAuth-related bugs (OpenClaw, IronClaw, ZeroClaw, Hermes). The pattern: silent refresh failures, stale tokens, credential mismatches. **Recommendation:** Invest in token refresh lifecycle monitoring, proactive health checks (like IronClaw’s two-tier TTL), and user-visible diagnostics (ZeroClaw’s `doctor` command).

### 3. **Memory Is the New Battleground**
From silent migration (OpenClaw) to cursor conflicts (NanoBot) to missing embeddings (ZeroClaw), memory management is the most common source of regressions. **Takeaway:** Any memory store change must include backward-compatible migration paths, deterministic tests, and clear warnings. Projects that treat memory as a first-class upgradable component (IronClaw’s Reborn learning system, NanoBot’s caching) will win user trust.

### 4. **Cost Consciousness Drives Feature Requests**
Users explicitly ask for heartbeat model overrides (NanoBot), prompt caching (NanoClaw), and configurable streaming batch ms (Hermes). **Action:** Expose API cost controls as user-configurable parameters; prioritize caching and token-efficient defaults.

### 5. **Channel Proliferation Demands Abstraction**
Each project adds channels independently, leading to duplication (Telegram, iMessage, DingTalk, WhatsApp). **Insight:** CoPaw’s message queue cross-talk bug (#5354) and OpenClaw’s QQ Bot regression (#76935) suggest that shared channel abstraction layers (like IronClaw’s manifest-driven ingress) are needed to reduce per-channel maintenance burden.

### 6. **Quality Gates Are Emerging**
ZeroClaw’s RFC for a local pre-submission gate (`zerocode`) and OpenClaw’s session-replay harness (#80176) signal a shift toward **pre-merge quality enforcement**. Developers should consider adding replay-based regression testing to catch silent regressions before release.

---

*Report generated from community digest summaries dated 2026-06-21. Data reflects snapshots taken at end-of-day; actual project status may change rapidly.*

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-06-21

## Today's Overview
The project saw very heavy development activity, with **12 pull requests merged/closed** in the last 24 hours, far exceeding typical daily volumes. Two issues were resolved (a concurrency bug and a performance optimization), while two new feature requests were opened. No new release was cut, but the rapid pace of commits suggests a release is likely near. The community contributed multiple channels (iMessage, WhatsApp enhancements) and several maintainer-led performance/stability fixes landed.

---

## Releases
*None.* No new versions were published today.

---

## Project Progress
**12 pull requests were merged or closed today**, spanning SDK, tooling, channels, and performance:

- **Concurrency & Stability**  
  - `#4425` (fix: isolate per-run hooks for concurrent `run()` calls) – resolves the critical bug [#4408](https://github.com/HKUDS/nanobot/issues/4408).  
  - `#4343` (fix: reject unknown builtin parameters) – prevents silent failures from malformed tool call arguments.  
  - `#4256` (fix: keep memory cursor monotonic) – prevents cursor conflicts after history compaction.

- **Performance**  
  - `#4421` (cache tool-definition JSON in `estimate_prompt_tokens`) – reduces redundant serialization on every agent turn.  
  - `#4428` (cache tool schema estimates) – adds bounded caching for token estimation, complementary to `#4421`.

- **SDK**  
  - `#4296` (expand Python SDK runtime controls) – enriches `RunResult`, adds stable helper clients, and maintains backward compatibility.

- **Channels**  
  - `#4426` (add iMessage channel via Photon Spectrum) – brings iMessage support with a Node sidecar pattern.  
  - `#4407` (seed WhatsApp LID→phone mappings on startup) – resolves first-message resolution failures.

- **Web UI & Fetch**  
  - `#4430` (feat: configure web_fetch provider) – adds explicit provider selection (`auto`, `tavily`, `jina`, `readability`). *(Still open – see below)*  
  - `#4427` (fix: prevent iOS Safari auto‑zoom on textarea focus) – a mobile UX fix.  
  - `#4423` (fix: narrow rich capability error detection in Telegram) – avoids false positives on transient `BadRequest` errors.

- **Dependencies & Docs**  
  - `#4405` (allow Keenable search without an API key) – uses a token‑less public endpoint.  
  - `#4432` (docs: update README news through 2026-06-20) – backfills changelog for the past days.

---

## Community Hot Topics
The most active items (by comments) were resolved issues from earlier days:

- **Issue #4408** [concurrency bug](https://github.com/HKUDS/nanobot/issues/4408) (2 comments) – raised by @waelantar, quickly fixed in PR #4425 today. Though closed, the discussion highlighted a latent design risk in SDK hook management.
- **Issue #4420** [performance optimization](https://github.com/HKUDS/nanobot/issues/4420) (1 comment) – filed by @codeLong1024, who identified and benchmarked redundant tiktoken encoding. The fix was merged as two PRs (#4421, #4428) within 24 hours, demonstrating attentive maintainer response.
- **PR #4296** [SDK expansion](https://github.com/HKUDS/nanobot/pull/4296) – though merged 10 days ago, it was closed today and represents a foundational change that will affect all downstream SDK users.

Underlying needs: Performance efficiency for high‑turn agents, concurrency safety for multi‑user SDK usage, and channel diversity (iMessage, WhatsApp robustness).

---

## Bugs & Stability
| Severity | Issue / PR | Description | Status |
|----------|------------|-------------|--------|
| **Critical** | [#4408](https://github.com/HKUDS/nanobot/issues/4408) | Per‑run hooks clobbered under concurrent `run()` calls | Fixed in [#4425](https://github.com/HKUDS/nanobot/pull/4425) (merged today) |
| **Medium** | [#4343](https://github.com/HKUDS/nanobot/pull/4343) | Unknown builtin parameters silently accepted | Fixed (merged today) |
| **Medium** | [#4256](https://github.com/HKUDS/nanobot/pull/4256) | Memory cursor becomes non‑monotonic after compaction | Fixed (merged today) |
| **Low** | [#4423](https://github.com/HKUDS/nanobot/pull/4423) | Telegram rich‑capability error detection too broad | Fixed (merged today) |
| **Low** | [#4427](https://github.com/HKUDS/nanobot/pull/4427) | iOS Safari auto‑zoom on textarea focus | Fixed (merged today) |

No new bugs were opened today.

---

## Feature Requests & Roadmap Signals
The following open items represent community demands and likely near‑term additions:

- **Thinking / Reasoning mode** ([#4429](https://github.com/HKUDS/nanobot/issues/4429)) – Request to allow custom provider to configure non‑standard thinking parameters (e.g., VolcEngine’s `thinking.type`). Likely to land soon given close coupling with provider abstraction.
- **Heartbeat model override** ([#4431](https://github.com/HKUDS/nanobot/issues/4431)) – Allow `HeartbeatService` to use a cheaper model. Straightforward config change; high demand for cost control.
- **Web fetch provider config** ([#4430](https://github.com/HKUDS/nanobot/pull/4430)) – Already a PR; explicit selection of web‑fetch backends. Expected in next release.
- **Subagent aggregated result mode** ([#4414](https://github.com/HKUDS/nanobot/pull/4414)) – Buffered subagent output instead of real‑time streaming. Useful for summarized multi‑step tasks.
- **Onboard wizard improvements** ([#4395](https://github.com/HKUDS/nanobot/pull/4395)) – TTY detection, JetBrains‑inspired palette, draft retention. Improves first‑time setup experience.

Prediction: Features #4430 and #4429 (thinking style) have strong chances for the next minor release. Heartbeat override (#4431) is small and may be cherry‑picked.

---

## User Feedback Summary
- **Performance pain point** (@codeLong1024, #4420): The `estimate_prompt_tokens` function caused noticeable slowness in their `nanobee` project, especially in high‑turn sessions. The maintainers responded with two optimizations within 24h.
- **Feature gaps** (@gkd2323c, #4429): Inability to enable reasoning on custom providers like VolcEngine/Doubao. This signals a need for extensible thinking parameter definitions.
- **Cost sensitivity** (@steeveroucaute10-epping, #4431): Wanting a separate model for heartbeat to reduce expenses. Indicates real‑world deployment concerns.
- **Channel contributions** (@morandot, #4426 iMessage; @franciscomaestre, #4407 WhatsApp LID): The community is actively building channel integrations, suggesting high demand for multi‑platform deployment.

Overall satisfaction appears high, evidenced by rapid bug fixes and responsive feature additions.

---

## Backlog Watch
Items requiring maintainer attention (oldest open PRs/issues without recent updates):

- **PR #4346** ([fix: mark stripped images as unviewable](https://github.com/HKUDS/nanobot/pull/4346), opened June 15) – Addresses a bug where stripped image paths are leaked instead of replaced. No comments since open; still a draft. The bug affects all providers that strip images. Needs review.
- **PR #4395** ([onboard wizard improvements](https://github.com/HKUDS/nanobot/pull/4395), opened June 18) – No reviewer feedback yet. This is a user‑facing feature and may need design alignment.
- **PR #4409** ([alternative hooks approach](https://github.com/HKUDS/nanobot/pull/4409), opened June 18) – Draft PR for a different fix of #4408; now superseded by #4425 which was merged. Should be closed to avoid confusion.

All three are safe to close or move forward.

---

*Generated from GitHub data (HKUDS/nanobot) as of 2026-06-21.*

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest – 2026-06-21

## 1. Today's Overview

The project saw very high activity today with **50 pull requests updated** (38 open, 12 merged/closed) and **9 issues updated** (8 open, 1 closed). No new release was published. Development momentum is strong, with a heavy focus on bug fixes across the gateway, CLI, cron, and tool components, alongside several feature PRs advancing mobile, plugin, and memory capabilities. The number of open issues (especially P2 bugs) suggests ongoing stability work, but the dense PR pipeline indicates an actively responsive maintainer team.

## 2. Releases

**No new releases today.** The latest release information is unavailable.

## 3. Project Progress

**Merged/closed pull requests today (12 total):**

- **#50115** `fix(cli): warn when in-session model switch will preflight-compress` – Merged. Adds a warning to the user when switching to a lower-context model mid-session, preventing silent context loss.
- **#49539** `fix(cli): warn when in-session model switch will preflight-compress` – Closed (superseded by #50115). Partial fix for the same guardrail.
- Other 10 closed/merged PRs not shown in the top-20 list, indicating additional fixes merged across components.

**Notable open PRs advancing features or fixing critical bugs:**

- **#50121** `fix(gateway): harden HERMES_TELEGRAM_FOLLOWUP_GRACE_SECONDS env parse` – Addresses a crash bug (#50120) on the message-handling hot path.
- **#50119** `feat: jcode feature adoption (semantic recall, worktree conflicts, benchmark)` – Introduces opt-in semantic memory backends and a recall store.
- **#49834** `feat(mobile): Hermes Agent Android app – Capacitor thin client` – Concept draft for Android support.
- **#50014** `feat(dingtalk-platform): audio ASR, file messages, and Robot OpenAPI send` – Extends the DingTalk plugin with missing capabilities.
- **#50109** `fix(tools): preserve core tools when a platform bundle is disabled` – Stops platform bundles from accidentally disabling shared toolsets.

## 4. Community Hot Topics

The most active discussion today was relatively subdued:

- **Issue #41782** [Bug]: Desktop conversation interface jumps unexpectedly. (3 comments)  
  Link: https://github.com/NousResearch/hermes-agent/issues/41782  
  *Underlying need*: A reliable, non-interfering UI for reviewing conversation details. Users expect the viewport to stay fixed during interaction.

- **Issue #50081** [Bug]: TTS hard timeout truncates long-running command TTS providers. (2 comments)  
  Link: https://github.com/NousResearch/hermes-agent/issues/50081  
  *Underlying need*: Flexible timeout handling for TTS subprocesses, especially for providers that need longer processing.

Most other issues and PRs have zero comments, indicating limited community debate or that maintainers are handling items quickly.

## 5. Bugs & Stability

**Bugs reported today (21 June 2026), ranked by severity:**

| Severity | Issue | Component | Summary | Fix PR exists? |
|----------|-------|-----------|---------|----------------|
| **P2** | [#50081](https://github.com/NousResearch/hermes-agent/issues/50081) | tools/tts | `_run_command_tts` hard timeout truncates long-running TTS | No |
| **P2** | [#50120](https://github.com/NousResearch/hermes-agent/issues/50120) | gateway | Malformed env var crashes message handler on hot path | Yes (#50121) |
| **P2** | [#50106](https://github.com/NousResearch/hermes-agent/issues/50106) | cli | `_write_status_dot` causes silent message loss when status is a directory | No |
| **P2** | [#50105](https://github.com/NousResearch/hermes-agent/issues/50105) | cli/config | `validate_config_structure()` misses model sub-keys → silent fallback & unexpected billing | No |
| **P2** | [#50028](https://github.com/NousResearch/hermes-agent/issues/50028) | tools/mcp | Ping keepalive on unsupported server causes reconnect loop | Closed (likely fixed) |
| **P3** | [#41782](https://github.com/NousResearch/hermes-agent/issues/41782) | desktop | Conversation interface jumps to another location | No |
| **P3** | [#50101](https://github.com/NousResearch/hermes-agent/issues/50101) | tool/memory | `mnemosyne_diagnose` reports `fastembed: MISSING` despite successful install | No |

**Notable:** Five P2 bugs reported today, with one fix immediately opened (#50121 for #50120). The MCP ping reconnect loop (#50028) was closed, likely resolved. Config validation (#50105) represents a silent billing risk that needs attention.

## 6. Feature Requests & Roadmap Signals

**New feature requests today:**

- **Issue #50107** `[Feature Request]: Add configurable STREAM_BATCH_MS to reduce Desktop App CPU/heat` – User on M2 MacBook reports 30–40% CPU idle / 60–100%+ during streaming.  
  Link: https://github.com/NousResearch/hermes-agent/issues/50107  
  *Prediction*: Likely to land in the next minor release given the clear performance pain point.

- **Issue #50091** `[Feature Request]: Skill Precipitator — Auto-Discover Reusable Workflows from Session History` – Suggests automating skill creation from patterns in conversation history.  
  Link: https://github.com/NousResearch/hermes-agent/issues/50091  
  *Prediction*: More experimental, may appear as an optional plugin or config flag in a later release.

**Roadmap signals from PRs today:**

- Mobile client: PR #49834 (Android thin client) – draft, but signals interest in mobile support.
- DingTalk platform: PR #50014 – adding audio/file capabilities, showing platform plugin maturation.
- Semantic recall adoption: PR #50119 – opt-in memory backends from jcode, a long-requested memory enhancement.
- Telegram anti-flood hardening: PR #50114 – token bucket and circuit breaker, improving gateway reliability.
- Gateway startup notification: PR #24271 (still open since May) – opt-in online notification.

These collectively point toward **mobile compatibility**, **richer platform integrations**, **improved memory**, and **gateway resilience** as near-term priorities.

## 7. User Feedback Summary

**Real pain points expressed by users (derived from issues):**

- **Desktop UI instability** (#41782): Conversation view jumps unpredictably – a core UX frustration.
- **TTS timeout rigidity** (#50081): Long-running TTS tasks get forcefully cut, losing output.
- **Silent message loss in CLI** (#50106): `_write_status_dot` causing messages to disappear without warning.
- **Unexpected billing from misconfiguration** (#50105): Incomplete `model:` config leads to silent fallback and higher costs.
- **MCP keepalive bugs** (#50028): Reconnect loops waste resources and degrade reliability.
- **False memory diagnostic** (#50101): `mnemosyne_diagnose` incorrectly reports missing dependencies, undermining trust.
- **High CPU usage on Desktop** (#50107): Users seeking configurable streaming batches to reduce heat/power consumption.

**Overall satisfaction** appears moderate – users are actively filing detailed bug reports and feature requests, indicating engagement, but the volume of P2 bugs signals that stability is a top concern. The quick fix timeline (e.g., #50121 within hours of #50120) suggests responsive maintainers.

## 8. Backlog Watch

**Important issues and PRs that have been open without resolution for extended periods:**

| Item | Component | Priority | Created | Summary | Status |
|------|-----------|----------|---------|---------|--------|
| [#6335](https://github.com/NousResearch/hermes-agent/pull/6335) | gateway/security | **P1** | 2026-04-08 | Restrict unauthenticated webhook binds to loopback (fixes RCE surface) | Open |
| [#24271](https://github.com/NousResearch/hermes-agent/pull/24271) | gateway | P3 | 2026-05-12 | Add opt-in startup online notification | Open |
| [#39166](https://github.com/NousResearch/hermes-agent/pull/39166) | cli | P3 | 2026-06-04 | Restart managed dashboard service after update | Open |
| [#44165](https://github.com/NousResearch/hermes-agent/pull/44165) | cli | P3 | 2026-06-11 | Tokenized dashboard cmdline matcher – detect global flags before subcommand | Open |

**Most critical:** PR #6335 addresses a **P1 security vulnerability** (unauthenticated RCE on webhooks). It has been open for over two months. Given the high severity, maintainer attention is strongly advised.

**Other long-standing items** (PRs #24271, #39166, #44165) are lower priority but are small UX/operational improvements that have been ready for merge for weeks. Their accumulation suggests either maintainer bandwidth constraints or a need for review delegation.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest — 2026-06-21

## 1. Today's Overview

Project activity remains low, with two open issues and one open pull request updated in the past 24 hours, and no items closed or merged. A new nightly build (`v0.3.0-nightly`) was published, but it is explicitly marked as unstable. Both tracked issues are labelled `stale`, suggesting that while attention is being paid (updates within the last 1–2 days), core maintainer involvement appears limited. The single open PR (#2964) has been awaiting review for over three weeks without comment from maintainers. Overall, the project shows a modest maintenance cadence, with no clear signals of active development or rapid iteration.

## 2. Releases

A **nightly build** was released today:  
**v0.3.0-nightly.20260621.287853ab**  

- **Type:** Automated nightly build (unstable)  
- **Changes:** Includes changes between the last stable tag `v0.3.0` and the current `main` branch.  
- **Breaking changes / migration notes:** None documented; the release note advises caution due to potential instability.  
- **Full Changelog:** [v0.3.0...main](https://github.com/sipeed/picoclaw/compare/v0.3.0...main)

No official stable release was issued today.

## 3. Project Progress

**No pull requests were merged or closed today.** The only PR updated in the last 24 hours is still open:

- **#2964** *(open since May 28)* — Adds configurable inbound image compression for the vision pipeline. The PR remains uncommented by maintainers, and its status suggests no progress toward integration.

No bug fixes or feature advancements were delivered in the current day.

## 4. Community Hot Topics

Two issues attracted the most community interaction:

- **#3012** — *[BUG] Continuous consumption of tokens every minute when evolution is enabled*  
  [Link](https://github.com/sipeed/picoclaw/issues/3012)  
  - 5 comments, 0 reactions  
  - Reported by xpader, running PicoClaw v0.2.9 on FreeBSD.  
  - **Underlying need:** Users enabling Evolution (Draft mode with Code Path Trigger) experience a token leak, where the agent repeatedly consumes tokens even without user input. The issue points to a potential logic bug in the evolution loop, causing unnecessary API costs and breaking normal usage.

- **#2984** — *[Feature][Protocol] Add explicit turn completion signal for Pico WebSocket clients*  
  [Link](https://github.com/sipeed/picoclaw/issues/2984)  
  - 3 comments, 2 👍 reactions  
  - Authored by Brook-sys.  
  - **Underlying need:** External Pico Protocol WebSocket clients lack a deterministic end-of-turn indicator. Clients currently rely on low-level event sequences (`typing.stop`, etc.) which are ambiguous. A clear `turn.end` signal would improve client reliability and enable better synchronisation in multi-turn conversations.

Both issues remain open and have not received a response from the maintainer team.

## 5. Bugs & Stability

Only one bug report was updated today:

- **#3012** — *Continuous token consumption with evolution enabled*  
  - **Severity:** Moderate to High (causes unbounded token usage, impacting cost and usability).  
  - **Status:** Open, stale, no fix PR identified.  
  - **Environment:** v0.2.9, FreeBSD, MiniMax model.  
  - **Potential regression:** Not clear if this is a new behaviour; reporter notes it happens "every minute" when evolution is enabled.

No crashes, regressions, or security issues were reported today. The lack of a fix PR or maintainer comment on this bug is a concern.

## 6. Feature Requests & Roadmap Signals

- **#2984** — *Explicit turn completion signal*  
  Demand for a deterministic signalling mechanism in the Pico Protocol. Likely to be prioritised if the project aims to support third-party WebSocket clients. Could appear in a future minor release if merged.

- **#2964** — *Image input compression*  
  This PR adds configurable compression for inbound images, addressing a gap where only a blanket `max_media_size` existed. If merged, it would be a useful addition for users sending larger or lower-quality images through vision channels. The feature seems well-scoped, but its stalled review (no maintainer feedback) suggests it may not be a current priority.

Both items signal user interest in improved performance, cost control, and protocol clarity. No roadmap or milestone tags are present.

## 7. User Feedback Summary

- **Pain points:**  
  - Token waste due to evolution loop bug (#3012) — directly impacts cost and trust in the Evolution feature.  
  - Lack of deterministic turn end for WebSocket clients (#2984) — forces clients to implement fragile heuristics.

- **Use cases:**  
  - Users relying on Evolution for draft/code generation pipelines (xpader).  
  - Developers integrating PicoClaw as a backend agent via custom WebSocket clients (Brook-sys).

- **Satisfaction / dissatisfaction:**  
  - No overtly positive feedback observed. Users express frustration through bug reports and feature requests that remain unanswered.  
  - The stale labels on both issues may indicate perceived neglect by the community.

## 8. Backlog Watch

Two open items are particularly important and have been waiting for maintainer attention:

| Item | Age (days since creation) | Last Maintainer Response | Risk |
|------|---------------------------|--------------------------|------|
| [#3012](https://github.com/sipeed/picoclaw/issues/3012) – Token leak bug | 16 | None | High – could drive users away from Evolution feature |
| [#2984](https://github.com/sipeed/picoclaw/issues/2984) – Turn completion signal | 19 | None | Medium – blocks client ecosystem growth |
| [#2964](https://github.com/sipeed/picoclaw/pull/2964) – Image compression PR | 24 | None | Medium – valuable PR languishes unreviewed |

These items represent the most critical gaps in responsiveness. Without maintainer engagement, the project risks losing community momentum and trust.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-21

## 1. Today’s Overview

NanoClaw shows moderate activity with **1 open issue** and **4 open pull requests** updated in the last 24 hours, but **no releases** and **no merged/closed PRs** today. The single open issue (#2768) seeks a performance improvement for the Claude provider, while the four PRs focus on cleanup and reliability fixes—two from the same contributor (CutSnake01) removing dead code and stale documentation, one fixing a setup socket timing issue, and another correcting the main seed prompt. The absence of merged work and new releases suggests a maintenance- or review-oriented day rather than forward momentum. Overall project health appears stable, with active community contributions addressing minor technical debt and user-facing issues.

## 2. Releases

**None.** No new versions were published in the last 24 hours or in the project’s latest release list.

## 3. Project Progress

**No pull requests were merged or closed today.** The four open PRs represent work in progress:

- **#2825** (amit-shafnir) – *fix(setup): wait for the host socket before failing the first chat*  
  Addresses a race condition where setup’s “first chat” step fails because the host socket is not yet bound. A user-facing reliability fix.

- **#2824** (CutSnake01) – *fix: drop stale “Global Memory” instruction from main seed prompt*  
  Removes outdated references from the system prompt, which could confuse agent behavior.

- **#2823** (CutSnake01) – *fix: remove groups/global/CLAUDE.md (host deletes it on every startup)*  
  Eliminates a file that is automatically overwritten by the host, preventing confusion.

- **#2822** (CutSnake01) – *refactor(container-runner): drop dead /workspace/global mount*  
  Cleans up an unused Docker volume mount, reducing configuration clutter.

## 4. Community Hot Topics

The most discussed item is **Issue #2768** – *Enable prompt caching by default in Claude provider*  
- **Author:** galmorduku  
- **Created:** 2026-06-14 | **Updated:** 2026-06-20 | **Comments:** 1  
- **Links:** [Issue #2768](https://github.com/nanocoai/nanoclaw/issues/2768)

This issue points out that the Claude provider does not set `enablePromptCaching` when calling the Anthropic Agent SDK, causing the full system prompt to be re-sent on every turn. The single comment indicates community interest in reducing latency and cost for agents with rich prompts. The underlying need is for better performance and more efficient API usage—a common pain point for Claude-based agents.

All four open PRs have **zero comments** and **zero reactions**, suggesting they are new or under review without broader community discussion yet.

## 5. Bugs & Stability

**No new bugs or crashes were reported in the last 24 hours.**  
The only stability-related work is **PR #2825** (fixing a socket timing race condition during setup), which is still open and not yet merged. This is a low-severity issue that only affects the first-run experience. No other regressions or critical bugs are recorded.

## 6. Feature Requests & Roadmap Signals

The clearest feature signal is **Issue #2768** – enabling prompt caching in the Claude provider. Given its specificity (a single configuration flag) and the performance benefits for power users, this is likely to be implemented in the next minor release or patch. The community contributor has already identified the exact code location (`container/agent-runner/src/providers/claude.ts`) and the fix is trivial—adding `enablePromptCaching: true` to the `sdkQuery()` call. No other user-requested features appeared in the day’s data.

## 7. User Feedback Summary

While direct user feedback is sparse, the nature of the open PRs reveals implicit pain points:

- **Frustration with setup reliability** – PR #2825 stems from a user (or developer) who experienced the “first chat” step failing because the host socket wasn’t ready. This suggests the onboarding experience can be flaky.
- **Confusion from stale documentation** – PR #2824 removes a “Global Memory” instruction from the seed prompt, implying users or agents may have been misled by outdated system instructions.
- **Wasted configuration complexity** – PRs #2823 and #2822 clean up files and mounts that are either auto-deleted or unused, indicating that the current configuration structure can confuse users trying to customize or understand the system.

Overall satisfaction is not measurable from today’s data, but the proactive cleanup and the prompt caching request suggest a community that values correctness, performance, and simplicity.

## 8. Backlog Watch

**No critically stale items identified.**  
The only open issue (#2768) is 7 days old and has received one comment, which is a normal pace for a feature request. The four open PRs are all very recent (last 1–2 days). No issues or PRs older than two weeks are present in the last-24h dataset. However, it is worth monitoring **#2768** if it remains unaddressed for another week, as straightforward performance improvements can lose community momentum if left unattended.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-06-21

## 1. Today’s Overview

IronClaw shows high activity with 20 pull requests updated in the last 24 hours (9 merged/closed, 11 open) and 2 issues updated (1 closed, 1 still open). Project momentum is strong: the core team is consolidating several manifest-driven channel features (ingress, auth, credential coherence) and landing workspace sharing infrastructure. A high-risk OAuth refresh bug was fixed and merged. One long-standing nightly E2E failure (Issue #4108) remains open without a fix, indicating a possible stability blind spot.

## 2. Releases

No new releases were published today. No release notes or version bumps are present in the data.

## 3. Project Progress (Merged/Closed PRs Today)

Nine PRs were closed or merged in the last 24 hours, spanning bug fixes, CI improvements, and feature consolidation:

| PR | Summary | Contributor |
|----|---------|-------------|
| [#5087](https://github.com/nearai/ironclaw/pull/5087) | **Fix:** Proactively refresh Google OAuth tokens before expiry (closes #5071) — two‑tier TTL refresh logic for access/refresh tokens. | henrypark133 |
| [#5106](https://github.com/nearai/ironclaw/pull/5106) | **Refactor:** Collapse per‑channel host‑ingress mount sprawl into a single generic `serve.rs` plan (Move 4). | serrrfirat |
| [#5102](https://github.com/nearai/ironclaw/pull/5102) | **Feat:** Add cross‑contract credential coherence invariant to v2 extension‑manifest projection (Move 3). | serrrfirat |
| [#5103](https://github.com/nearai/ironclaw/pull/5103) | **Feat:** Manifest‑projected ingress policy + typed auth + transport discriminator (keystone PR). | serrrfirat |
| [#4777](https://github.com/nearai/ironclaw/pull/4777) | **Fix:** Persist Slack connected state in WebUI to eliminate reconnect loop. | serrrfirat |
| [#4829](https://github.com/nearai/ironclaw/pull/4829) | **CI:** Remove dormant `reborn-integration` workflow; add Reborn suites to nightly deep CI. | serrrfirat |
| [#5105](https://github.com/nearai/ironclaw/pull/5105) | **Test:** Fix three stale provider/OAuth guard tests broken on `main`. | serrrfirat |
| [#5104](https://github.com/nearai/ironclaw/pull/5104) | **Feat:** Typed auth verifier + transport discriminator for host‑ingress (Move 2, −54 lines). | serrrfirat |
| [#2548](https://github.com/nearai/ironclaw/pull/2548) | **Feat:** Workspace entities with membership and cross‑workspace sharing (DB migration). | standardtoaster |

These merges advance the “manifest‑driven channels” initiative (now consolidated from four stacked PRs into one) and bring workspace sharing to production.

## 4. Community Hot Topics

No issues or PRs have accumulated comments or reactions in the observed window. However, two PRs involve **new contributors** (tagged `contributor: new`):

- [#5109](https://github.com/nearai/ironclaw/pull/5109) — Adds Composio connector routes for the Workbench (read‑only + gated‑write). Draft, awaiting review.
- [#4002](https://github.com/nearai/ironclaw/pull/4002) — Ongoing dependabot bump of 16 GitHub Actions (open since May 24, no merge yet).

The **nightly E2E failure** ([#4108](https://github.com/nearai/ironclaw/issues/4108), open since May 27) remains the most prominent unresolved community concern, though it has no comments or reactions. It likely affects users relying on nightly builds.

## 5. Bugs & Stability

**High‑severity**, fixed today:
- **#5071** (closed) – **Google OAuth token expiry bug** (risk: high). Users were forced to reauthenticate hourly. Fixed by PR [#5087](https://github.com/nearai/ironclaw/pull/5087), which implements proactive refresh.

**Medium‑severity**, fixed today:
- **#5105** – Three stale security‑relevant guard tests (provider/OAuth) that were failing on `main` due to outdated assertions. Fixed by PR [#5105](https://github.com/nearai/ironclaw/pull/5105).

**Medium‑severity**, ongoing:
- **#4108** – **Nightly E2E scheduled run fails** consistently since May 27. No fix PR in sight. Potential root cause in “Full E2E / E2E (extensions)” job – may affect extension stability.

**Low‑severity**, in progress:
- **#5108** (open) – Fixes three “reborn‑closure tail failures” surfaced by CI, including an over‑exposed GitHub manifest visibility bug.

## 6. Feature Requests & Roadmap Signals

Several open PRs signal upcoming features likely to land in the next release:

| PR | Feature | Likely next‑version? |
|----|---------|---------------------|
| [#5085](https://github.com/nearai/ironclaw/pull/5085) | **Concurrent turn execution** via `TurnRunScheduler` with per‑user/per‑type caps. Reduces LLM inference serialization. | High |
| [#5065](https://github.com/nearai/ironclaw/pull/5065) | **One‑shot scheduled triggers** (`TriggerSchedule::Once`), complementing existing Cron. | High |
| [#5107](https://github.com/nearai/ironclaw/pull/5107) | **Manifest‑driven channel ingress contract** – consolidates auth, transport, secrets, and onboarding into manifest definitions. | High |
| [#4937](https://github.com/nearai/ironclaw/pull/4937) | **Reborn learning system** (WS‑1) – memory document with confidence, categories; A/B gated. First step of “learn from mistakes”. | Medium (depends on landing WS‑1) |
| [#5081](https://github.com/nearai/ironclaw/pull/5081) | **Hosted single‑tenant Postgres profile** – durable state for Reborn preview. | High |
| [#4765](https://github.com/nearai/ironclaw/pull/4765) | **Lift subagent inline prompt body budget** – introduces `LoopInlineMessageBody` >512 bytes. | Medium |

The overall roadmap is focused on **concurrency, declarative channel configuration, durable state, and agent learning**.

## 7. User Feedback Summary

No explicit user feedback (comments, reactions) was captured in the data. However, downstream fixes address these likely pain points:

- **Google OAuth reauthentication** (#5071) – users were forced to reconnect every hour; now fixed.
- **Slack reconnect loop** (#4777) – WebUI now correctly reflects Slack delivery connection state.
- **Stale test failures** (#5105) – three security tests that were failing without causing alerts; now fixed, improving CI reliability.

No new dissatisfaction signals were recorded.

## 8. Backlog Watch

The following open items have remained unattended for extended periods:

- **Issue #4108** (created May 27, 2026) – **Nightly E2E failure** with no comments, no assignee, and no linked fix PR. The failure is in “Full E2E / E2E (extensions)” and has been failing for 25 days. Requires maintainer investigation.
- **PR #4002** (dependabot, open since May 24) – 16 GitHub Actions updates blocked. No merge or review activity. Could be waiting for CI green light or manual override.
- **PR #5109** (new contributor, draft) – Composio connector route for Workbench. No reviewer comments yet; may need core team prioritisation.

All other open PRs have recent activity (updated within 2 days) and appear actively managed.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-06-21

## Today's Overview

Today was a maintenance-only day for LobsterAI. No new pull requests were merged or opened, and no releases were published. The only activity was the closure of 14 stale issues by the automatic stale bot, all originally filed in early April 2026. This suggests the core team may be focused on higher-priority work or consolidating existing code before introducing new changes. The project remains stable with no reports of regressions or new bugs.

## Releases

No new releases today.

## Project Progress

No pull requests were updated or merged today. Feature development and bug fixes appear to be paused; the only change was the automated closure of old issues.

## Community Hot Topics

The most active issue by number of comments was **#1509** (3 comments), filed by jimmy-xz, regarding blocking problems during skills file generation with no user feedback, and inconsistent model understanding compared to Openclaw. This issue was closed today as stale but generated discussion about user experience gaps in the skills pipeline.

Other issues had only 2 comments each (mostly bot + author) and were closed automatically. None attracted external discussion beyond the original reporter and maintainers.

## Bugs & Stability

All 14 issues updated today were bugs or minor UI misalignments, all closed as stale. Severity ranges from critical UX blocking to cosmetic. Key bugs include:

- **Critical (blocking functionality):**
  - [#1500](https://github.com/netease-youdao/LobsterAI/issues/1500) – Disabled skills still appear in `activeSkillIds`, causing them to be injected into prompts.
  - [#1502](https://github.com/netease-youdao/LobsterAI/issues/1502) – Active skill list not updated in current session after saving Agent settings; requires switching agents.
  - [#1504](https://github.com/netease-youdao/LobsterAI/issues/1504) – IM robot AES Key validation missing; empty keys can be saved.
  - [#1506](https://github.com/netease-youdao/LobsterAI/issues/1506) – Scheduled IM notifications silently fail when no session is selected.
  - [#1512](https://github.com/netease-youdao/LobsterAI/issues/1512) – QQ Bot whitelist UI lacks input field for adding group IDs; whitelist mode unusable.
  - [#1516](https://github.com/netease-youdao/LobsterAI/issues/1516) – GitHub Copilot OAuth polling continues after closing Settings panel; authentication token lost on success.

- **Moderate (usability/missing feedback):**
  - [#1509](https://github.com/netease-youdao/LobsterAI/issues/1509) – No intermediate status during skills file generation; user cannot tell if the agent is working.
  - [#1513](https://github.com/netease-youdao/LobsterAI/issues/1513) – Inconsistent formatting (duplicate numbering, missing parentheses) in legal statement pages.

- **Minor (infrastructure):**
  - [#1518](https://github.com/netease-youdao/LobsterAI/issues/1518) – CI labeler workflow failing due to insufficient permissions; fix was already documented.

None of these issues have open fix PRs; they were closed by the stale bot, meaning they were not resolved in code. The team may revisit them when they re-prioritize.

## Feature Requests & Roadmap Signals

Multiple feature requests were filed in early April and closed today. These represent strong user demand for enhanced conversation management:

- **Color tags** ([#1525](https://github.com/netease-youdao/LobsterAI/issues/1525)) — Users want to visually distinguish different conversation types.
- **Batch export** ([#1528](https://github.com/netease-youdao/LobsterAI/issues/1528)) — Ability to export multiple conversations at once in JSON.
- **Local usage statistics** ([#1532](https://github.com/netease-youdao/LobsterAI/issues/1532)) — SQLite-based statistics panel (total sessions, messages, daily activity).
- **Message bookmark/favorite** ([#1537](https://github.com/netease-youdao/LobsterAI/issues/1537)) — Flag important AI replies in long conversations.
- **Conversation tags & filtering** ([#1541](https://github.com/netease-youdao/LobsterAI/issues/1541)) — Custom labels and sidebar filters to organize sessions.

These features, if accepted, could significantly improve LobsterAI’s appeal as a long-term productivity tool. No roadmap commitment has been given.

## User Feedback Summary

Today’s closed issues paint a picture of mixed user satisfaction. On the positive side, no critical regressions were reported; the platform is stable. However, users are frustrated by:

- Lack of visibility into agent operations (skills generation, scheduling) — “can’t tell if it’s working”
- Inconsistent model behavior across different frontends (LobsterAI vs. Openclaw)
- Missing validation and error feedback for configuration forms (IM bots, whitelists)
- UI/UX gaps that reduce efficiency in large-scale conversation management (color, tags, bookmarks, batch export)
- Persistent synchronization bugs between UI and internal state (skill activation, agent settings)

Overall, the feedback indicates that while core AI functionality works, the surrounding tooling (settings, conversation management, user feedback) needs significant refinement.

## Backlog Watch

No important open issues or pull requests remain unaddressed today — all updated items were closed. However, the high volume of stale-closed bugs (especially those with critical severity like #1500, #1502, #1506, #1512, #1516) suggests that the team may have a backlog of unresolved technical debt. Maintainers should prioritize re-opening and fixing these issues in the next sprint, as they directly impact daily usage reliability.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest – 2026-06-21

## Today’s Overview
Moltis saw minimal activity over the past 24 hours, with no issues updated and only two dependency‑bump pull requests processed. One PR was merged and closed, while one remains open. No new releases were published, and the project appears to be in a quiet maintenance phase. The focus today was entirely on updating JavaScript dependencies across the documentation and website directories.

## Releases
No new releases were cut today. The latest release version remains unchanged. Users can check [Moltis Releases](https://github.com/moltis-org/moltis/releases) for any future updates.

## Project Progress
- **Merged/closed PR:** [#1133 [CLOSED]](https://github.com/moltis-org/moltis/pull/1133) – Bumped `astro` from 6.3.3 to 6.4.8 in `/docs` as part of the `npm_and_yarn` group. No functional changes to Moltis core.
- **Open PR:** [#1134 [OPEN]](https://github.com/moltis-org/moltis/pull/1134) – A broader bump across two directories (`/docs` and `/website`) updating both `astro` (6.3.3 → 6.4.8) and `undici` (latest version). This is still awaiting review/merge.

No feature or bug‑fix PRs were merged today.

## Community Hot Topics
No issues or pull requests with significant comments or reactions were recorded in the last 24 hours. The only two PRs are automated dependency updates by `dependabot[bot]` with zero comments and zero reactions. Community discussion remains dormant.

## Bugs & Stability
No bug reports, crashes, or regressions were filed or updated today. The project’s stability appears unchanged, with no new issues to assess. The dependency bumps are routine maintenance and do not address any known defects.

## Feature Requests & Roadmap Signals
No feature requests were submitted or updated in the last 24 hours. There is no signal of upcoming functionality changes. The roadmap remains static; the next version may simply incorporate the already‑merged dependency updates.

## User Feedback Summary
No user feedback (comments, pain points, or satisfaction indicators) was recorded in the observed data. The lack of issues or discussion suggests either stable user satisfaction or low engagement. No actionable pain points can be derived from today’s activity.

## Backlog Watch
There are no long‑unanswered issues or PRs requiring maintainer attention. The only open PR (#1134) is recent (created yesterday) and is a routine dependency bump. No backlog items were identified. Maintainers should monitor the open PR for merge readiness and look for any new community contributions that may arise.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

Based on the provided GitHub data for the **CoPaw** repository (`agentscope-ai/CoPaw`) for the date **2026-06-21**, here is the project digest.

---

## CoPaw Project Digest — 2026-06-21

### 1. Today's Overview

The project is in a **high-activity, community-building phase** with a steady pulse of contributions but **zero new releases**. There are **13 open issues** and **18 open pull requests** updated in the last 24 hours, all of which remain unmerged or unclosed, indicating a bottleneck in review/merge velocity. The community is very active, with a significant number of "first-time contributor" PRs (4 out of 18), suggesting successful lowering of contribution barriers. The hottest topics revolve around **mobile UI responsiveness**, **custom provider compatibility** (specifically function calling), and **agent session stability** during long or interrupted thinking processes. While the feature pipeline is full, the project faces moderate stability challenges related to message queues, tool result size management, and model connection errors.

### 2. Releases

**None** — No new releases were recorded in the last 24 hours.

### 3. Project Progress

No PRs were merged or closed today. However, several critical fixes and significant features have been submitted and are awaiting review:

- **UI & Mobile Responsiveness:** Multiple PRs directly address user-reported mobile pain points. PR #5355 fixes the model selector dropdown being obscured on mobile, and PR #5334 adds an agent-switching button in the collapsed sidebar mode (directly addressing Issue #5329).
- **Core Stability Fixes:** PR #5335 provides a fix for the agent "freezing" issue (Issue #5333) by ensuring an error event is sent back to the UI when a model execution fails. PR #5340 resolves a bug where interrupting an agent leaves an empty message that causes the following turn to fail.
- **Feature Development:** Work is progressing on several requested features, including real-time SSE push notifications with voice beep (PR #5331 for Issue #5322), recency-aware memory ranking (PR #5325 for Issue #5316), and a native task progress panel for multi-step agent plans (PR #5323).
- **Security & Integrity:** PR #5341 constrains built-in file tools to the agent's workspace, preventing potential path traversal vulnerabilities.

### 4. Community Hot Topics

The most active discussions and underlying needs are concentrated on **mobile UI usability** and **custom model/provider compatibility**.

- **#5329 - Agent Switching on Mobile (4 comments):** [Link](agentscope-ai/QwenPaw Issue #5329)
    - **Analysis:** This is the most commented-on issue. The user is hacking QwenPaw to work on mobile but finds the UI unusable. The need is for a **responsive, mobile-first UI** that doesn't require complex navigation for core tasks like switching agents. This is clearly a significant pain point for a demographic of users who want to interact on the go.

- **#5345 - Custom Provider (OMLX) Function Calling Broken:**
    - **Analysis:** This issue highlights a critical **interoperability gap**. While QwenPaw supports custom OpenAI-compatible providers, the function calling pathway is broken for many. This limits the agent's utility, effectively making it a simple chatbot for a sizable group of users who rely on non-Ollama infrastructure.

- **#5328 - DeepSeek Agent Freezing (2 comments):**
    - **Analysis:** This is a high-frustration bug affecting a major model provider (DeepSeek). The agent gets "stuck" during a "thinking" phase, requiring manual intervention. This points to a **race condition or timeout issue** in the streaming/SSE handling for models that have an extended reasoning phase.

- **#5354 - Message Queue "Cross-Talk":**
    - **Analysis:** A newly reported but potentially high-impact bug. The new message queue feature, while improving efficiency, introduces a critical **state-management bug** where messages are sent to the wrong agent or conversations become inaccessible. This undermines a core user experience of having multiple, distinct conversations.

### 5. Bugs & Stability

Bugs reported today are ranked by potential impact on core functionality.

- **[HIGH] #5354 - Message Queue Cross-Talk & Conversation Loss:** [Link](agentscope-ai/QwenPaw Issue #5354)
    - **Analysis:** This is a **critical regression** in a recently introduced feature (message queue). The "cross-talk" breaks the fundamental assumption of isolated conversations, and the "conversation lockout" bug is a data loss risk. A fix is needed urgently. *No associated fix PR found.*

- **[HIGH] #5328 - Agent Freezing on DeepSeek (Thinking Step):** [Link](agentscope-ai/QwenPaw Issue #5328)
    - **Analysis:** A **high-frustration issue** affecting a popular model. The agent becomes unresponsive, requiring a manual click to proceed. This is a negative UX event. *PR #5335 provides a partial fix for a related "stuck" scenario by handling exceptions.*

- **[MEDIUM] #5345 - Custom OpenAI-Compatible Provider Function Calling Failure:** [Link](agentscope-ai/QwenPaw Issue #5345)
    - **Analysis:** A **core feature broken for a niche but growing user base**. This makes QwenPaw's key "Agent" feature non-functional for users of alternative providers like OMLX. *No associated fix PR found.*

- **[MEDIUM] #5344 - API Returns 200 but Silently Drops Messages:** [Link](agentscope-ai/QwenPaw Issue #5344)
    - **Analysis:** A **silent failure** in the API. Messages sent while the agent is "busy" are accepted an 200 but discarded. This is a serious reliability issue for programmatic users (e.g., automations, integrations). *No associated fix PR found.*

- **[MEDIUM] #5330 - Zhipu Provider Connectivity Paradox:** [Link](agentscope-ai/QwenPaw Issue #5330)
    - **Analysis:** An **inconsistency between provider-level and model-level connection tests**. This is confusing for users and blocks the setup of an officially listed provider. *PR #5339 directly addresses this bug.*

- **[LOW] #5342 - Unpruned Tool Results on LLM Failure:** [Link](agentscope-ai/QwenPaw Issue #5342)
    - **Analysis:** A **defense-in-depth issue**. While not a crash bug, the context window can be silently blown by error cascades, leading to long-term degradation of agent performance. *No associated fix PR found.*

### 6. Feature Requests & Roadmap Signals

Several user-requested features are already in the PR pipeline, suggesting they are likely candidates for the next minor release.

- **Mobile UI Revamp:** The high demand for mobile usage is being answered by PRs #5355, #5350, and #5334, which address dropdowns, header responsiveness, and agent switching in compact mode. **Prediction: Likely included in v1.1.13.**
- **Agent Office Interaction:** Issue #5327 requests an "Interact" button directly from the agent office dashboard. This is a well-defined UX improvement for multi-agent management. **Prediction: Strong candidate for a future release.**
- **Real-time Notifications & Voice:** PR #5331 is already a working implementation of the feature requested in Issue #5322. **Prediction: Likely included in v1.1.13.**
- **Memory Improvements:** PR #5325 implements the recency-aware ranking requested in Issue #5316, and a WIP PR (#5349) suggests a major memory backend migration (to ReMe4) is being explored.

### 7. User Feedback Summary

Real user feedback today highlights both strong engagement and notable pain points:

- **Pain Point #1 - Mobile Client Experience:** Users are actively trying to use QwenPaw on mobile (via browser) and find the UI fundamentally broken for that use case. The inability to switch agents or access history without a desktop screen is a significant barrier. *("...but the UI interface is problematic, I have no way to switch the agent...")*
- **Pain Point #2 - Model Compatibility Gaps:** Users are frustrated that models and providers officially listed (e.g., OMLX, Zhipu, DeepSeek) have critical issues like broken function calling or freezing behavior. This erodes trust in the platform's core "multi-provider" value proposition. *("...the model only returns text and will not call tools..."), ("...the agent often gets stuck while thinking...")*
- **Pain Point #3 - Core Workflow Instability:** The new message queue feature, while a "great improvement" to efficiency, introduced frustrating bugs like sending messages to the wrong conversation. This directly harms the user's organization and sense of control. *("...the message will be sent to agent B, causing it to do something inexplicable...")*
- **Satisfaction Signal:** Despite the bugs, users are actively pushing the boundaries of the project, integrating it with mobile browsers and APIs, suggesting a high degree of initial satisfaction and desire for the product to succeed.

### 8. Backlog Watch

There are no issues or PRs that are significantly old or have been left unanswered within the provided data window. All items were created or updated within the last 2-3 days. However, there are a few notable items to monitor:

- **[Critical Regression] #5354 - Message Queue Bugs:** This issue has only one comment, but it describes two severe bugs that could erode user confidence in the new queue feature. Maintainer response is critical to confirm the bugs and prioritize a fix.
- **[WIP] #5349 - Major Memory Backend Migration:** This PR is a large-scale, long-term refactoring (switching to "ReMe4"). It is work-in-progress and tagged with `[WIP]`. It should be watched as it could introduce breaking changes.
- **[First-Time Contributors] Various PRs:** PRs #5355, #5352, #5346, #5348, #5340, and #5341 are from first-time contributors. These require timely review and feedback to nurture new community members. A lack of response here could discourage future contributions.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-06-21

## 1. Today's Overview
Development activity remains high with **50 pull requests updated** in the last 24 hours (37 open, 13 merged/closed) and **6 issues updated** (4 closed, 2 open). No new releases were published. The momentum is concentrated on bug fixes for channel and runtime stability (e.g., preserving tool results on history trim), new features like a chat-based onboarding assistant (`zeroclaw onboard`), and infrastructure RFCs (retiring the standalone `aardvark-sys` crate, introducing a local pre-submission gate). Overall project health is strong, with critical S1 bugs from earlier this week already closed and merged.

## 2. Releases
**None.** No new tags or releases were published in the last 24 hours.

## 3. Project Progress
Thirteen pull requests were merged or closed today. Notable advances:

- **Memory & Testing** — [PR #7916](https://github.com/zeroclaw-labs/zeroclaw/pull/7916) (merged) adds deterministic regression tests for storage-reader timestamp/ordering edge cases (closes [Issue #7694](https://github.com/zeroclaw-labs/zeroclaw/issues/7694)).
- **Memory Stability** — [PR #7942](https://github.com/zeroclaw-labs/zeroclaw/pull/7942) (merged) decouples memory embedding API key from the chat provider, surviving embed failures gracefully.
- **Gateway Fix** — [PR #7940](https://github.com/zeroclaw-labs/zeroclaw/pull/7940) (merged) fixes agent rename ordering to persist config before moving owned state (closes S1 bug [Issue #7907](https://github.com/zeroclaw-labs/zeroclaw/issues/7907)).
- **Cron Pause/Resume** — [PR #7666](https://github.com/zeroclaw-labs/zeroclaw/pull/7666) (merged) adds `enabled` field to `CronPatchBody` and scopes agent checks to shell-command patches.
- **Observability** — [PR #8067](https://github.com/zeroclaw-labs/zeroclaw/pull/8067) (merged) categorizes and verb-tags agent-loop log events for better filtering.
- **Channel History** — [PR #8050](https://github.com/zeroclaw-labs/zeroclaw/pull/8050) (merged) preserves tool-result content when proactively trimming channel history (fixes [Issue #8049](https://github.com/zeroclaw-labs/zeroclaw/issues/8049)).
- **Doctor & Providers** — [PR #8010](https://github.com/zeroclaw-labs/zeroclaw/pull/8010) (merged) and [PR #8077](https://github.com/zeroclaw-labs/zeroclaw/pull/8077) (merged) improve custom provider validation and remove unused `rumqttc` dependency.
- **Hardware Deps** — [PR #8028](https://github.com/zeroclaw-labs/zeroclaw/pull/8028) (merged) gates `aardvark-sys` behind the `hardware` feature, preparing for an RFC to retire the standalone crate.

## 4. Community Hot Topics

**Most commented issues:**
- **[Issue #7694](https://github.com/zeroclaw-labs/zeroclaw/issues/7694)** (4 comments) — Enhancement for memory storage-reader timestamp edge-case coverage. Already closed via merged PR #7916. Demonstrates strong community engagement on test infrastructure.
- **[Issue #8043](https://github.com/zeroclaw-labs/zeroclaw/issues/8043)** (2 comments) — RFC to retire the standalone `aardvark-sys` crate and fold it into `zeroclaw-hardware`. Sponsored by @JordanTheJet, builds on PR #8028. Signals a desire to simplify the dependency tree.
- **[Issue #8078](https://github.com/zeroclaw-labs/zeroclaw/issues/8078)** (0 comments, new) — RFC for a local pre-submission gate (`zerocode`) that enforces contributor bar before code leaves the machine. Generated by @perlowja; likely to attract attention soon.

**Active large PRs:**
- **[PR #8033](https://github.com/zeroclaw-labs/zeroclaw/pull/8033)** (XL size, high risk) — Chat-based conversational setup assistant as default `zeroclaw onboard`. This is a major user experience improvement for first-time users.
- **[PR #8030](https://github.com/zeroclaw-labs/zeroclaw/pull/8030)** (M size, high risk) — Warns in `zeroclaw doctor` on OpenAI Codex profile/slot wiring mismatches. Addresses silent failures during model configuration.

**Underlying needs:** Users are asking for smoother onboarding (conversational setup), better error diagnostics (Codex credential mismatch, custom provider validation), and infrastructure simplification (dependency consolidation, pre-submit quality gates).

## 5. Bugs & Stability
No new bugs were reported today as open issues. All recent severe bugs have been fixed and merged:

- **S1 (workflow blocked)** — [Issue #7907](https://github.com/zeroclaw-labs/zeroclaw/issues/7907) (agent rename could move owned state before config persistence) → fixed by PR #7940.
- **S3 (minor)** — [Issue #7888](https://github.com/zeroclaw-labs/zeroclaw/issues/7888) (unnecessary `rumqttc` dependency) → fixed by PR #8077.
- **S3 (minor)** — [Issue #8049](https://github.com/zeroclaw-labs/zeroclaw/issues/7964) (channel proactive trim dropping tool results) → fixed by PR #8050.

Additionally, several bug-fix PRs remain open:
- **[PR #8009](https://github.com/zeroclaw-labs/zeroclaw/pull/8009)** — Wires HMAC tool receipts through agent turn paths (medium risk, missing feature integration).
- **[PR #8048](https://github.com/zeroclaw-labs/zeroclaw/pull/8048)** — Fixes runtime to honor `history_pruning` config and keep tool results under context pressure (medium risk).
- **[PR #7864](https://github.com/zeroclaw-labs/zeroclaw/pull/7864)** — Omits `tool_choice` when tool list is empty (previously caused errors with vLLM). Open since June 17.

No regressions were reported today.

## 6. Feature Requests & Roadmap Signals

**Strong candidates for the next release:**
- **Conversational onboarding** ([PR #8033](https://github.com/zeroclaw-labs/zeroclaw/pull/8033)) — Revives `zeroclaw onboard` as a chat-based setup. High impact for new users.
- **Auto-clean temporary files** ([PR #7923](https://github.com/zeroclaw-labs/zeroclaw/pull/7923)) — Adds a `[files_cleanup]` config block for automated temporary file management. Open since June 18.
- **ZeroCode pre-submission gate** ([Issue #8078](https://github.com/zeroclaw-labs/zeroclaw/issues/8078)) — RFC proposes enforcing contributor bar locally before PR submission. If accepted, it will become a workflow feature.
- **Provider credential diagnostics** ([PR #8029](https://github.com/zeroclaw-labs/zeroclaw/pull/8029)) — Distinguishes missing vs expired OpenAI Codex credentials, improving error messages.

**RFCs under discussion:**
- **Retire `aardvark-sys`** ([Issue #8043](https://github.com/zeroclaw-labs/zeroclaw/issues/8043)) — If ratified, the hardware crate will simplify its dependency footprint.

## 7. User Feedback Summary
While the dataset lacks direct user quotes, the bug reports and feature requests reveal key pain points:

- **Agent rename peril** (S1) — Users or operators renaming agents could lose state before config persistence. The fix emerged within three days, showing responsiveness.
- **Silent context loss** — Channel history trimming discarded tool results, breaking long sessions. Fix accepted quickly.
- **Provider configuration friction** — Multiple PRs target credential mismatches (Codex), custom provider validation, and empty tool choice serialization. Indicates users are struggling with provider setup.
- **Onboarding friction** — The revival of `zeroclaw onboard` as a conversational flow suggests existing CLI setup is not user-friendly for newcomers.
- **Documentation gaps** — [PR #6870](https://github.com/zeroclaw-labs/zeroclaw/pull/6870) (feature support matrix) has been open since May 23, indicating users want clearer capability comparisons.

Overall satisfaction appears positive given the rapid bug triage and feature velocity.

## 8. Backlog Watch

**Long-unanswered items needing maintainer attention:**

| Item | Days Open | Notes |
|------|-----------|-------|
| [PR #6870](https://github.com/zeroclaw-labs/zeroclaw/pull/6870) | 29 | Docs: feature support matrix. Low risk, no maintainer review yet. |
| [PR #7864](https://github.com/zeroclaw-labs/zeroclaw/pull/7864) | 4 | Fix providers: omit `tool_choice` when empty. Needs review; targets multiple providers. |
| [PR #7923](https://github.com/zeroclaw-labs/zeroclaw/pull/7923) | 3 | Feat: auto-clean temporary files. Config change; waiting for maintainer sign-off. |
| [Issue #8043](https://github.com/zeroclaw-labs/zeroclaw/issues/8043) | 1 | RFC: retire aardvark-sys. Needs ratification vote per process. |
| [PR #8009](https://github.com/zeroclaw-labs/zeroclaw/pull/8009) | 2 | Fix: wire HMAC receipts through agent paths. Medium risk, no comments yet. |
| [PR #8030](https://github.com/zeroclaw-labs/zeroclaw/pull/8030) | 2 | Feat: doctor Codex mismatch warning. High risk, no maintainer review. |

The oldest open PR (#6870) is purely documentation and low priority. The RFCs (#8043, #8078) require formal ratification process. No issue has been abandoned for more than a week without updates.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*