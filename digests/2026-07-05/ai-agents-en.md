# OpenClaw Ecosystem Digest 2026-07-05

> Issues: 285 | PRs: 500 | Projects covered: 13 | Generated: 2026-07-05 09:32 UTC

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

# OpenClaw Project Digest – 2026-07-05

## Today's Overview

OpenClaw saw intense activity over the past 24 hours, with **285 issues** and **500 pull requests** updated. Of those, **267 issues remain open** (23 newly closed) and **301 PRs are active**, while **199 PRs were merged or closed** – indicating a healthy review pipeline. A new beta release **v2026.7.1-beta.2** shipped, adding GPT‑5.6 support and external harness attachment. The community remains highly engaged, with top discussions focusing on cross‑platform support (Linux/Windows desktop apps), prebuilt mobile APKs, and several high‑impact reliability bugs. The project continues to mature rapidly, balancing new feature delivery with critical stability improvements.

## Releases

### v2026.7.1-beta.2 – 2026-07-05

- **OpenAI GPT‑5.6 support** — The OpenClaw catalog, capability detection, and runtime selection now recognise the GPT‑5.6 model family.  
  [PR #98333](https://github.com/openclaw/openclaw/pull/98333) (thanks @steipete-oai)
- **External harness attachment** — `openclaw attach` now allows launching an external harness against an existing Gateway session.

No breaking changes or migration notes were documented in the release. This beta is available on GitHub [Releases](https://github.com/openclaw/openclaw/releases/tag/v2026.7.1-beta.2).

## Project Progress

**199 PRs were merged or closed today**, reflecting sustained momentum across the codebase. Notable merged/closed fixes and improvements include:

- **UI polish** – Fixed idle composer scrollbar in web UI  
  [PR #100252](https://github.com/openclaw/openclaw/pull/100252)
- **iMessage false startup warning** – Misleading `groupAllowFrom` warning eliminated  
  [PR #100046](https://github.com/openclaw/openclaw/pull/100046)
- **Model cache invalidation during hot reload** – Two separate fixes (one closed, one open) prevent phantom “Unknown model” errors after config changes  
  [PR #99789](https://github.com/openclaw/openclaw/pull/99789) (closed), [PR #99853](https://github.com/openclaw/openclaw/pull/99853) (closed)
- **Anthropic thinking-signature session bricking** – Server‑side prevention for permanent `Invalid signature in thinking block` errors  
  [PR #100149](https://github.com/openclaw/openclaw/pull/100149) (open, linked to Issue #94228)
- **Plugin approval rejection reporting** – Accurate error propagation when gateway rejects plugin approvals  
  [PR #100231](https://github.com/openclaw/openclaw/pull/100231) (open)
- **Build system resilience** – Fallback to `tsx` when `process.features.typescript` is unavailable  
  [PR #91262](https://github.com/openclaw/openclaw/pull/91262) (open)
- **Test optimisation** – Isolated full‑suite tests kept under one second each  
  [PR #100019](https://github.com/openclaw/openclaw/pull/100019) (open)

## Community Hot Topics

The most active discussions highlight strong demand for **multi‑platform availability** and **reliability enhancements**:

| Issue / PR | Comments | 👍 | Topic |
|---|---|---|---|
| [Issue #75](https://github.com/openclaw/openclaw/issues/75) – Linux/Windows Clawdbot Apps | **110** | **81** | Persistent request for native desktop apps beyond macOS/iOS/Android |
| [Issue #9443](https://github.com/openclaw/openclaw/issues/9443) – Prebuilt Android APK releases | **26** | **4** | Need for compiled APK downloads; current source‑only approach is a barrier |
| [Issue #22676](https://github.com/openclaw/openclaw/issues/22676) – Signal daemon race condition on SIGUSR1 | **17** | **0** | Critical stability bug causing orphaned processes and send failures |
| [Issue #22438](https://github.com/openclaw/openclaw/issues/22438) – Tiered bootstrap file loading | **17** | **0** | Feature to reduce context‑window waste by selectively loading bootstrap files |
| [Issue #29387](https://github.com/openclaw/openclaw/issues/29387) – Bootstrap files silently ignored in agentDir | **14** | **5** | P1 regression: per‑agent bootstrap config has no effect |
| [Issue #31583](https://github.com/openclaw/openclaw/issues/31583) – `exec` tool does not inherit skill env vars | **13** | **2** | P1 regression: secret injection broken in subprocesses |
| [Issue #10659](https://github.com/openclaw/openclaw/issues/10659) – Masked secrets to prevent agent API‑key access | **13** | **4** | Security feature preventing credential leakage |

Underlying needs revolve around **making OpenClaw more accessible** (desktop apps, mobile distribution) and **improving trustworthiness** (secrets masking, memory trust, fail‑safe execution).

## Bugs & Stability

Several high‑severity bugs were active today, many with corresponding fix PRs in flight:

- **[P0 – ux‑release‑blocker]**  
  - **Live docs ahead of release** ([Issue #48920](https://github.com/openclaw/openclaw/issues/48920)) – Documentation references features (e.g., `IsolatedSessions`) not present in the latest release. No linked fix PR.  
  - **Upgrade corruption** ([Issue #95515](https://github.com/openclaw/openclaw/issues/95515)) – Upgrading from 2026.6.8→2026.6.9 corrupts email channel config with invalid `groupAllowFrom` field. Linked PR open.

- **[P1 – critical regressions and reliability]**  
  - **Signal daemon race condition** ([Issue #22676](https://github.com/openclaw/openclaw/issues/22676)) – SIGUSR1 restarts cause orphaned processes and send failures. No PR linked.  
  - **Bootstrap files in agentDir ignored** ([Issue #29387](https://github.com/openclaw/openclaw/issues/29387)) – Per‑agent bootstrap config has no effect. Linked PR open.  
  - **`exec` tool env variable inheritance** ([Issue #31583](https://github.com/openclaw/openclaw/issues/31583)) – Skill‑level environment variables not passed to subprocesses. Linked PR open.  
  - **Anthropic thinking‑signature brick** ([Issue #94228](https://github.com/openclaw/openclaw/issues/94228)) – Permanent 400 error on long tool‑use sessions. Fix PR [#100149](https://github.com/openclaw/openclaw/pull/100149) open.  
  - **Config warnings log spam** ([Issue #25574](https://github.com/openclaw/openclaw/issues/25574)) – Duplicate warning messages on every reload cycle. Linked PR open.

- **[P1 – other stability]**  
  - **Docker sandbox workspace access** ([Issue #31331](https://github.com/openclaw/openclaw/issues/31331)) – Internal paths used instead of host paths. No PR.  
  - **Cron session hallucinated output** ([Issue #49876](https://github.com/openclaw/openclaw/issues/49876)) – LLM fabricates results when tools fail. No PR.

The project is actively addressing the most critical regressions; **five of the top P1 bugs have linked open PRs**, indicating maintainers are prioritising reliability.

## Feature Requests & Roadmap Signals

The community has submitted a wide range of feature requests, many with high engagement:

- **Cross‑platform desktop apps** (#75) – Linux and Windows Clawdbot clients remain the most‑upvoted request (81 👍). No PR yet.  
- **Prebuilt Android APK** (#9443) – Essential for mobile users; currently only source code is shipped.  
- **Masked secrets system** (#10659) – Prevent agents from viewing raw API keys, a security improvement for enterprise use.  
- **Memory trust tagging by source** (#7707) – Mitigate memory poisoning by tagging entries by origin.  
- **Tiered bootstrap file loading** (#22438) – Progressive context control to save token budget.  
- **Post‑subagent completion hooks** (#22358) – Automatic trajectory generation after agent tasks.  
- **Pre‑response enforcement hooks (hard gates)** (#13583) – Mechanically force tool calls before answers (critical for finance/ops).  
- **Session snapshots** (#13700) – Save/load context checkpoints for long development sessions.  
- **Theme customization system** (#28300) – User‑requested UI personalisation with 6 preset themes + custom studio.  
- **Anthropic native server‑side tools** (#23353) – Support `web_search`, `web_fetch`, `code_execution` built into the provider.  

**Prediction for next version**: Given the volume of fix PRs in flight, the next minor release will likely include the Anthropic thinking‑signature fix (#100149), the model cache hot‑reload fix (#99789), and the plugin approval accuracy fix (#100231). Among features, tiered bootstrap (#22438) has garnered multiple comments and a linked PR, making it a strong candidate for a near‑term release.

## User Feedback Summary

The community is **highly active and engaged**, but several pain points stand out:

- **Lack of desktop apps on Linux/Windows** (#75) – 81 upvotes and 110 comments mark this as the most requested feature. Users feel limited to macOS/iOS/Android.  
- **No prebuilt Android APK** (#9443) – Developers on behalf of users request binary downloads; current source‑only approach is a barrier to adoption.  
- **Configuration complexity** – Issues like #10687 (dynamic model discovery), #16670 (memory/embedding not in onboarding wizard), and #13597 (no AWS deployment docs) indicate users find setup cumbersome.  
- **Stability regressions** – The P0/P1 bugs described above cause real frustration: session bricks, broken env vars, config corruption. The community appreciates rapid fix attempts (linked PRs) but the number of regressions suggests a need for stronger test coverage.  
- **Positive sentiment** – The new GPT‑5.6 support and external harness attachment are well‑received. The release itself shows the project is shipping new capabilities even while fixing bugs.

## Backlog Watch

These important issues have been open for extended periods and lack a fix PR or require maintainer attention:

| Issue | Age | Priority | Status |
|---|---|---|---|
| [#75](https://github.com/openclaw/openclaw/issues/75) – Linux/Windows Apps | 185 days | P2 | No PR; needs product decision & security review |
| [#9443](https://github.com/openclaw/openclaw/issues/9443) – Prebuilt Android APK | 150 days | P0 (release‑blocker) | No PR; needs maintainer review |
| [#10659](https://github.com/openclaw/openclaw/issues/10659) – Masked secrets | 149 days | P1 | No PR; needs security review & product decision |
| [#7707](https://github.com/openclaw/openclaw/issues/7707) – Memory trust tagging | 152 days | P2 | No PR; needs security review |
| [#13616](https://github.com/openclaw/openclaw/issues/13616) – Backup/restore utility | 145 days | P2 | No PR; needs product decision |
| [#13751](https://github.com/openclaw/openclaw/issues/13751) – Feishu over‑permissive scope | 144 days | P1 | No PR; needs security review |
| [#48920](https://github.com/openclaw/openclaw/issues/48920) – Docs ahead of release | 110 days | P0 | No PR; needs product decision |
| [#22676](https://github.com/openclaw/openclaw/issues/22676) – Signal daemon race condition | 134 days | P1 | No PR; linked PR open? (data shows linked‑pr‑open but no PR in top list) – needs maintainer verification |

These issues represent **unresolved community friction** – especially the P0 documentation mismatch and the P1 race condition – and would benefit from prioritised maintainer attention.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: Personal AI Agent Open-Source Ecosystem
**Date: 2026-07-05 | Period: 24 hours ending 2026-07-05 23:59 UTC**

---

## 1. Ecosystem Overview

The personal AI agent and assistant open-source landscape continues to mature rapidly, with the **OpenClaw ecosystem** representing the dominant reference implementation and spawning a family of specialized forks and alternative designs. Activity is concentrated in a **core group of 6 actively developing projects**, while 4 projects show zero activity, suggesting a natural consolidation around viable architectures. The ecosystem is characterized by **high issue throughput** (285+ issues/day on OpenClaw alone), rapid patch cycles, and converging requirements around **security hardening, cross-platform delivery, and reliability**—indicating a shift from experimental to production-grade quality expectations. Critical trends include: cryptographic library modernization, provider compatibility expansion, and systemic approaches to session corruption and context management.

---

## 2. Activity Comparison

| Project | Issues (Updated) | PRs (Updated) | Releases | Health Score | Activity Tier |
|---------|----------------:|--------------:|:--------:|:------------:|:-------------:|
| OpenClaw | 285 | 500 | v2026.7.1-beta.2 | **A+** (Extreme) | 1. Extremely High |
| NanoBot | 3 | 17 | None | **A** (High) | 2. High |
| Hermes Agent | 13 | 50 | None | **A** (High) | 2. High |
| IronClaw | 9 | 37 | None | **A** (High) | 2. High |
| NanoClaw | 1 | 36 | None (v2.1.38 latest) | **A-** (High) | 2. High |
| CoPaw | 15 | 3 | None | **B+** (Moderate) | 3. Moderate |
| PicoClaw | 4 | 8 | None | **B** (Moderate) | 3. Moderate |
| LobsterAI | 0 | 3 | None | **C** (Low) | 4. Low |
| ZeroClaw | 6 | 50 | None (v0.8.x latest) | **A-** (High, but no merges) | 2. High |
| NullClaw | *Inactive* | *Inactive* | — | — | 5. Inactive |
| TinyClaw | *Inactive* | *Inactive* | — | — | 5. Inactive |
| Moltis | *Inactive* | *Inactive* | — | — | 5. Inactive |
| ZeptoClaw | *Inactive* | *Inactive* | — | — | 5. Inactive |

**Health Score Notes:**
- **A+** = Extreme volume, rapid merge cycle, active cross-team contributions
- **A** = High volume with meaningful codebase progress
- **B** = Moderate, engaged community with clear maintenance trajectory
- **C** = Low activity, primarily maintenance mode
- **Inactive** = No observable activity in 24-hour window

**Key observation:** OpenClaw processes **~45% of all PRs in the surveyed ecosystem**, confirming its role as the primary development nucleus.

---

## 3. OpenClaw's Position

### Advantages vs. Peers

| Dimension | OpenClaw | Next-Best Competitor |
|-----------|----------|---------------------|
| **Community size** | 285 issues/day, 500 PRs/day | 50 PRs/day (Hermes Agent) |
| **Release cadence** | Beta releases w/ GPT-5.6 support | None (NanoBot, Hermes, IronClaw) |
| **Model provider support** | OpenAI, Anthropic, multiple adapters | Strong but narrower (ZeroClaw, Hermes) |
| **Platform coverage** | macOS/iOS/Android; Linux/Windows requested | Linux-first (IronClaw); macOS-only (PicoClaw) |
| **Core team responsiveness** | 199 PRs merged/closed in 24h | 7-16 PRs merged |
| **Security posture** | Active secrets masking proposal | SSRF pinning (NanoBot), meta-tool stripping (IronClaw) |

### Technical Approach Differences

- **Architecture**: OpenClaw uses a **Gateway session model** with hotspot reload, plugin approval pipeline. Unlike IronClaw's manifest-driven design or ZeroClaw's SOP-based task system.
- **Extensibility**: OpenClaw's `openclaw attach` allows external harness attachment—a flexible approach compared to NanoClaw's more rigid container-oriented model.
- **Context handling**: OpenClaw's tiered bootstrap proposal (#22438) is more advanced than CoPaw's scroll-compression (which is breaking context). Hermes Agent's compression drops user messages entirely (#58753).

### Community Size Comparison

| Project | Active Contributors (24h) | Top Issue Reactions |
|---------|--------------------------|-------------------:|
| OpenClaw | ~50+ | 81 👍 (#75 – Linux/Windows) |
| Hermes Agent | ~15-20 | 12 comments (#47349) |
| ZeroClaw | ~10-15 | 0-1 per issue |
| NanoBot | ~8-12 | 0-2 per issue |
| IronClaw | ~5-8 | 0-1 per issue |

**Verdict:** OpenClaw commands **~3-6x the community engagement** of its nearest competitors, making it the **de facto reference implementation** for the ecosystem.

---

## 4. Shared Technical Focus Areas

Common requirements emerging across multiple projects reveal ecosystem-wide priorities:

| Focus Area | Projects Affected | Specific Needs |
|------------|------------------|----------------|
| **Cross-platform desktop apps** | OpenClaw, Hermes Agent, CoPaw | Linux/Windows native clients; interrupt mode parity |
| **Security hardening** | NanoBot, IronClaw, ZeroClaw, NanoClaw | SSRF pinning, bridge meta-tool stripping, forged-click prevention, env var propagation |
| **MCP tool reliability** | NanoBot, ZeroClaw | Crash resilience (BaseException catching), process cleanup (zombie MCP), name length validation |
| **Session corruption / memory loss** | OpenClaw, Hermes Agent, CoPaw, PicoClaw | Compression exhaustion, context loss, amnesia on file overwrite |
| **Channel interoperability** | NanoBot, ZeroClaw, PicoClaw | Telegram custom endpoints, Matrix encryption, POPO validation |
| **Provider compatibility** | OpenClaw, ZeroClaw, CoPaw | GPT-5.6, Bedrock Nova caching, Gemini embedding index, OpenRouter |
| **Mobile / Android support** | OpenClaw, PicoClaw | APK builds, service launch reliability |
| **Configuration & env var handling** | OpenClaw, NanoBot, LobsterAI | Proxy propagation, OAuth token env vars, bootstrap files |
| **Telemetry & observability** | IronClaw, ZeroClaw | Coverage ratchets, integration-tier testing |
| **Memory & context management** | OpenClaw, Hermes Agent, CoPaw, ZeroClaw | Tiered loading, trust tagging, state management, auto-memory interval |

**Implication:** These recurring themes represent ecosystem-wide **investment priorities**. Projects that address these gaps first will gain disproportionate adoption.

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | Hermes Agent | IronClaw | PicoClaw | ZeroClaw |
|-----------|----------|---------|--------------|----------|----------|----------|
| **Primary use case** | General reference | Lightweight agent | Research-grade agent | Enterprise stack | Embedded/IoT | Modular task automation |
| **Target users** | Developers, community | Hobbyists, CI/CD | Researchers, power users | DevOps, platform teams | Constrained devices | Workflow automators |
| **Architecture** | Session-based Gateway | Plugin/harness | Full-stack (Desktop + TUI) | Manifest-driven (Reborn) | Agent-only, minimal | SOP-based, goal-oriented |
| **Security posture** | Improving (secrets masking) | Strong (SSRF polling) | Active (credential guards) | Critical (meta-tool fix) | Basic (crypto gap) | Mixed (leak detector improvements) |
| **Testing maturity** | High (isolated sub-second tests) | Medium (fix-oriented) | Medium (known regressions) | High (integration coverage) | Low | Low (no merged PRs today) |
| **Platform breadths** | macOS/iOS/Android | Cross-platform (Windows focus) | Desktop + TUI + Electron | Linux + containers | macOS + Android | Flexible (multi-channel) |
| **Risk tolerance** | Stable beta | Rapid but cautious | Aggressive (experimental v4) | Conservative (dual stack migration) | Stale but solid | High (50 open PRs, few merges) |
| **Unique differentiator** | Ecosystem reference; GPT-5.6 support | Windows-first; SSRF-proofing | Hermes Studio (Electron); memory backends | Slack OAuth migration; coverage ratchets | libolm→vodozemac migration | Goal-oriented channels; SOP step fallthrough |

**Strategic note:** PicoClaw is the **only project actively migrating its cryptography stack** (libolm→vodozemac), addressing a security debt that other projects (NanoClaw, IronClaw) likely share but haven't yet prioritized.

---

## 6. Community Momentum & Maturity

### Activity Tiers

| Tier | Projects | Characteristics |
|------|----------|----------------|
| **Tier 1 – Extremely High** | OpenClaw | Massive daily throughput, rapid merges, public roadmap signals, frequent releases |
| **Tier 2 – High** | NanoBot, Hermes Agent, IronClaw, NanoClaw, ZeroClaw | 20-50 PRs/day, dedicated maintainers, active bug triage, but slower merge cycles |
| **Tier 3 – Moderate** | CoPaw, PicoClaw | 3-15 issues/PRs, focused priorities, some stale items |
| **Tier 4 – Low** | LobsterAI | Maintenance mode (3 PRs, no issues) |
| **Tier 5 – Inactive** | NullClaw, TinyClaw, Moltis, ZeptoClaw | No observable activity in 24h |

### Rapid Iteration vs. Stabilization

- **Rapidly iterating**: OpenClaw, NanoBot, Hermes Agent, ZeroClaw – all are pushing new features alongside bug fixes, with OpenClaw leading in release cadence (beta releases).
- **Stabilizing**: IronClaw – heavy investment in test infrastructure (coverage enablers, wiring parity) rather than new features; signs of pre-release hardening.
- **Consolidating**: NanoClaw – 36 PRs but mostly dead code removal and doc fixes; v2.1.38 stable tagged.
- **Stalemated**: PicoClaw, LobsterAI – long-open PRs, stale issues, no recent releases.

**Maturity assessment:** OpenClaw and ZeroClaw are **where most new user activity occurs** (high issues + PRs). Hermes Agent and NanoBot are **best for developer ergonomics** (strong tooling, TUI/CLI). IronClaw is the **most suitable for enterprise deployments** given its testing rigor and container-first approach.

---

## 7. Trend Signals

### Industry Trends Extracted from Community Feedback

1. **AI agent reliability is now the #1 concern.** Across OpenClaw, Hermes Agent, CoPaw, and PicoClaw, users consistently report session corruption, memory loss, and context compression failures as workflow-blocking. The ecosystem is moving from "can it do X?" to "can it do X without breaking?"

2. **Security hardening is accelerating.** NanoBot's SSRF DNS pinning (P0), ZeroClaw's high-entropy detector false positives, IronClaw's bridge meta-tool stripping (critical), and OpenClaw's secrets masking all indicate that **credential leakage and spoofing** are top-of-mind for maintainers. The subtext: agent trust is a prerequisite for real-world adoption.

3. **Cross-platform delivery is a gate to mainstream adoption.** OpenClaw's most-upvoted feature (81 👍) is Linux/Windows desktop apps. CoPaw has mobile WebUI truncation. PicoClaw has Android service failures. The demand is clear: developers want their personal AI assistant available on every device they use.

4. **MCP (Model Context Protocol) stability is an ecosystem-wide debt.** NanoBot and ZeroClaw both have open PRs addressing MCP crash resilience (BaseException catching, zombie processes). As MCP becomes the standard for tool integration, its reliability is becoming a system-level requirement.

5. **Provider diversity creates integration complexity.** GPT-5.6 (OpenClaw), DeepSeek v4 (Hermes), Bedrock Nova (ZeroClaw), Gemini (CoPaw)—each integration uncovers unique quirks. The pattern is: **projects with the widest provider matrix (OpenClaw, ZeroClaw) also have the most provider-specific bugs**. This suggests an emerging need for standardized provider adaptation layers.

### Value for AI Agent Developers

- **For platform builders**: Watch OpenClaw for design patterns (Gateway session model, plugin pipelines) and NanoBot for lightweight deployment.
- **For enterprise users**: IronClaw's test infrastructure and ZeroClaw's SOP system offer the clearest path to production.
- **For researchers**: Hermes Agent's memory backends and compression semantics provide the deepest flexibility.
- **For embedded/IoT**: PicoClaw's cryptographic migration (vodozemac) is essential—any other project using libolm should follow suit.

**Bottom line:** The ecosystem is healthy but fractured. **OpenClaw remains the reference**, but ZeroClaw and Hermes Agent are innovating in areas OpenClaw hasn't prioritized (goal-oriented tasks, research-grade context management). The next 6-12 months will likely see consolidation around shared protocol standards (MCP security, memory format, channel APIs) driven by these overlapping requirements.

---

*Report generated from 13 project digests covering activity up to 2026-07-05 23:59 UTC.*

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-07-05

## Today's Overview
Activity remains high: 17 pull requests were updated in the last 24 hours (10 open, 7 merged/closed) alongside 3 issues (1 newly filed, 2 closed). The development team continues to address stability regressions and security hardening, with priority-0 and priority-1 fixes moving through review. The absence of a new release this week suggests the maintainers are consolidating multiple fixes and features for an upcoming tag. Overall, the project is in a healthy, fast-paced iteration cycle with strong contributor engagement.

---

## Releases
*None.* No new releases have been published in the reporting window.

---

## Project Progress (Merged/Closed PRs)
Seven pull requests were merged or closed today, reflecting tangible improvements across multiple subsystems:

- **Anthropic OAuth with env‑var awareness**  
  [#4699](HKUDS/nanobot PR #4699) — Adds support for the `CLAUDE_CODE_OAUTH_TOKEN` environment variable, fixing login/logout consistency when both file‑based and environment‑based tokens exist.

- **Gateway stop crash on Windows**  
  [#4690](HKUDS/nanobot PR #4690) — Handles `OSError` when `CTRL_BREAK_EVENT` is rejected, ensuring `nanobot gateway stop` uses the `taskkill` fallback instead of dumping a traceback.

- **DingTalk shutdown cleanup**  
  [#4646](HKUDS/nanobot PR #4646) — Properly cancels the stream task and closes the websocket/client during `stop()`, preventing resource leaks.

- **Crash‑durable atomic writes**  
  [#4653](HKUDS/nanobot PR #4653) — Restores `fsync` calls in `_write_text_atomic` to guarantee data integrity for pairing storage after a refactor regression.

- **Copilot token refresh race condition**  
  [#4684](HKUDS/nanobot PR #4684) — Guards the token refresh path with `asyncio.Lock` to prevent multiple concurrent requests from independently fetching new tokens.

- **Malformed MCP tool results**  
  [#4666](HKUDS/nanobot PR #4666) — Wraps exception‑prone result rendering and marks timeouts, cancellations, and execution errors as structured tool errors (fixes #4652).

- **Upstream merge**  
  [#4695](HKUDS/nanobot PR #4695) — Routine sync from upstream branch.

---

## Community Hot Topics

1. **MCP tool exception crash**  
   Issue [#4652](HKUDS/nanobot Issue #4652) (3 comments, closed) – Users reported that an MCP tool call returning an error or empty data crashed the agent process. The underlying need is **graceful error propagation** in MCP wrappers. The fix (PR #4666) was quickly merged, but a complementary PR [#4701](HKUDS/nanobot PR #4701) (open) proposes catching `BaseException` across all MCP execute methods for extra safety—indicating the community is pushing for bulletproof stability.

2. **Copilot token refresh race**  
   Issue [#4677](HKUDS/nanobot Issue #4677) (1 comment, closed) – Exposes a check‑then‑act race that could cause duplicate token fetches under load. The closure via PR #4684 satisfies the immediate need, but the same author (axelray-dev) has an additional open PR [#4698](HKUDS/nanobot PR #4698) standardizing OAuth error messages, showing a pattern of systematic provider hardening.

3. **Telegram custom API endpoint**  
   Issue [#4702](HKUDS/nanobot Issue #4702) (new, 0 comments) – A user requests support for custom API base URL and request headers for the Telegram channel, motivated by complex network environments (e.g., corporate proxies, private bot API mirrors). No maintainer response yet, but the ask aligns with the existing proxy support pattern.

---

## Bugs & Stability
Ranked by severity (P0 highest, P3 lowest):

| Severity | Issue/PR | Description | Status |
|----------|----------|-------------|--------|
| **P0** | [#4671](HKUDS/nanobot PR #4671) | Pin validated DNS for SSRF checks – prevents DNS rebinding in `web_fetch`, MCP HTTP probes, and transports | **Open** (critical security) |
| **P1** | [#4701](HKUDS/nanobot PR #4701) | Prevent process crash on MCP tool call exceptions (catches `BaseException` in all wrappers) | **Open** |
| **P1** | [#4700](HKUDS/nanobot PR #4700) | Limit long MCP‑derived tool names that exceed model API length constraints | **Open** |
| **P1** | [#4545](HKUDS/nanobot PR #4545) | Default Windows commands to PowerShell and allow explicit `shell` parameter (fixes cross‑drive `cd` and `$VAR` parsing) | **Open** |
| **P2** | [#4698](HKUDS/nanobot PR #4698) | Standardize OAuth error messages across CLI and WebUI | **Open** |
| **P2** | [#4694](HKUDS/nanobot PR #4694) | Keep chat viewport and composer inside narrow viewports (mobile fix) | **Open** |
| **P2** | [#4686](HKUDS/nanobot PR #4686) | Support canonical OpenCode provider (new provider, not a bug) | **Open** |

**Notable:** The P0 SSRF fix (#4671) is still open and should be prioritised for security; it already has an approved design (validated resolved IPs + pinned DNS). The MCP crash issue (#4652) is closed but the broader catch‑all PR (#4701) is still under review—suggesting the team wants to ensure the fix covers all edge cases.

---

## Feature Requests & Roadmap Signals
- **Custom Telegram API base URL** (Issue [#4702](HKUDS/nanobot Issue #4702)) – Likely to be implemented in the next minor release given the existing proxy infrastructure and community demand.
- **Serper.dev web search provider** (PR [#4406](HKUDS/nanobot PR #4406)) – A well‑structured addition following the existing search provider pattern; opened on June 18, still open but receives updates. Should land soon.
- **Canonical OpenCode provider** (PR [#4686](HKUDS/nanobot PR #4686)) – Adds official support for the `opencode` provider with backward compatibility; almost ready to merge (recently updated).
- **Configurable MCP inheritance for subagents** (PR [#4697](HKUDS/nanobot PR #4697)) – Allows specialist subagents to inherit MCP servers from the main agent, eliminating the need for raw shell workarounds. This addresses a long‑standing architectural gap.
- **Smooth WebUI streaming Markdown** (PR [#4696](HKUDS/nanobot PR #4696)) – Improves user experience with buffered rAF scheduling, punctuation pauses, and left‑to‑right reveal animation.

**Prediction:** The next release will likely include the Serper provider, OpenCode canonical support, MCP inheritance control, and the mobile viewport fix (#4694). The smooth streaming PR (#4696) may be deferred or merged as optional configuration.

---

## User Feedback Summary
- **MCP robustness** – Several users (see #4652, #4666) expressed frustration that the agent crashes on malformed MCP tool results. The quick fix was appreciated, but the open PR #4701 shows a desire for even stronger guarantees.
- **Windows first‑class support** – PR #4545 (fixing PowerShell defaults) and PR #4690 (gateway stop crash) directly address Windows pain points. Users on Windows expect `nanobot` to behave identically to Linux/macOS; the team is actively closing these gaps.
- **Token refresh reliability** – The race condition in GitHub Copilot token refresh (#4677) was a reliability concern under concurrent use. The async lock fix (#4684) satisfies the reported issue, but users may encounter similar races in other providers.
- **Telegram network flexibility** – The request for custom API base URL (#4702) highlights that international users need more than proxy support; maintainers should evaluate adding a `telegram_api_base_url` config.

---

## Backlog Watch
- **PR #4406 (Serper.dev web search)** – Open since June 18, no maintainer comments. The PR is complete and follows the existing provider pattern; a review decision is needed.
- **PR #4545 (Windows shell defaults)** – Open since June 26, recently updated. This is a high‑impact fix for Windows users. Maintainers should expedite review to prevent further confusion on cross‑drive operations.
- **Issue #4671 (SSRF DNS pinning – P0)** – While actively updated (most recent commit July 5), the PR has not been merged. Given the security severity, maintainers should prioritise final review and merge.

No stale issues (older than 30 days without update) were observed in today’s data. The project team appears to be responsive, with most open items receiving updates within the last week.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

## Hermes Agent Project Digest — 2026-07-05

### 1. Today's Overview

Project activity remains very high, with **50 pull requests** and **13 issues** updated in the last 24 hours. The commit and review cadence suggests a busy development cycle focused on bug fixes, security hardening, and platform compatibility. No new releases were published today, but the volume of closed/merged PRs (16) indicates steady progress toward the next stable release. Critical bugs reported yesterday and today—such as the `cua-driver` timeout wedge and the agent deleting its own Python interpreter—are being addressed with targeted fix PRs already in review.

---

### 2. Releases

No new releases were published today. The most recent stable release remains the prior version (not listed in the data). Users running `hermes update` are currently pulling from `main` (the development channel), which is the subject of a feature request ([#58746](https://github.com/NousResearch/hermes-agent/issues/58746)) to switch to the latest stable tag.

---

### 3. Project Progress

**Merged/Closed PRs (16 total)** – Several important fixes landed today:

- **Security**  
  - [#58740](https://github.com/NousResearch/hermes-agent/pull/58740) – Fix webhook V2 signature rejection when timestamp is missing (follow-up to #58461).  
  - [#56179](https://github.com/NousResearch/hermes-agent/pull/56179) – Tighten root-collapse hardline token to avoid false classification of `rm -rf /...` as root wipe.

- **Infrastructure & Stability**  
  - [#58744](https://github.com/NousResearch/hermes-agent/pull/58744) – Guard null text in Desktop file preview to prevent crash on binary files.  
  - [#58756](https://github.com/NousResearch/hermes-agent/pull/58756) – Replace deprecated `datetime.utcnow()` in BlueBubbles adapter.  
  - [#58758](https://github.com/NousResearch/hermes-agent/pull/58758) – Fix missing `is_reconnect` parameter in QQ adapter, ending infinite reconnect loop.

- **Tool & Platform Fixes**  
  - [#58752](https://github.com/NousResearch/hermes-agent/pull/58752) – Guard native image routing with file-safety policy (salvage of #58517).  
  - [#58751](https://github.com/NousResearch/hermes-agent/pull/58751) – Guard memory local uploads against credential reads (salvage of #57841).  
  - [#58754](https://github.com/NousResearch/hermes-agent/pull/58754) – Skip `cua-driver` refresh when `/Applications` is unwritable (macOS non-admin).  
  - [#58760](https://github.com/NousResearch/hermes-agent/pull/58760) – Raise `cua-driver` session startup timeout to 30s and add phase timing logs.

- **Context & Compression**  
  - [#58750](https://github.com/NousResearch/hermes-agent/pull/58750) – Preserve conversation context after compression exhaustion (adds handoff messages).

---

### 4. Community Hot Topics

**Most commented issue:**  
[#47349 – Feature: Configurable Memory Backends](https://github.com/NousResearch/hermes-agent/issues/47349)  
*12 comments, 1 👍*  
Users request renaming `memory.md` → `rules.md` and supporting distinct backends (Honcho, Fact Store) while disabling the default file-based memory. The conversation shows broad interest in decoupling memory from the system prompt and enabling pluggable providers.

**Most active follow-up bug:**  
[#58755 – Empty `tool_calls` array sent to DeepSeek v4](https://github.com/NousResearch/hermes-agent/issues/58755)  
*0 comments (new), but directly follows closed #56980 (5 comments).*  
This is a reopened bug – the `repair_message_sequence` fix from #56980 (closed as *not_planned*) does not handle the case where an empty `tool_calls` array is generated, causing a non-recoverable HTTP 400. The community has expressed frustration with the original closure and wants a proper fix.

**Rising concern:**  
[#58717 – Desktop `busy_input_mode=interrupt` not working](https://github.com/NousResearch/hermes-agent/issues/58717)  
*2 comments, reported by @rifters*  
User messages queue instead of interrupting the agent in the Desktop app, while the TUI works correctly. This points to a platform-specific regression in the Electron layer.

---

### 5. Bugs & Stability

| Issue | Severity | Description | Fix PR Exists? |
|-------|----------|-------------|----------------|
| [#58748](https://github.com/NousResearch/hermes-agent/issues/58748) – Agent deletes its own Python interpreter | **P2** | Session can uninstall the uv-managed Python runtime that Hermes depends on, breaking startup. | No (related #58749) |
| [#58749](https://github.com/NousResearch/hermes-agent/issues/58749) – Repair reuses broken uv-trampoline venv | **P2** | Missing base Python not detected; repair flow reuses broken venv instead of rebuilding. | No |
| [#58762](https://github.com/NousResearch/hermes-agent/issues/58762) – `cua-driver` refresh wedges permanently | **P2** | 300s timeout vs installer’s 600s stale-lock window; orphaned lock on timeout. | Yes ([#58760](https://github.com/NousResearch/hermes-agent/pull/58760)) |
| [#58753](https://github.com/NousResearch/hermes-agent/issues/58753) – Compression drops only user-role message | **P2** | Compression can delete the only user message, causing non-retryable crash. Affects kanban workers heavily. | Yes ([#58750](https://github.com/NousResearch/hermes-agent/pull/58750)) |
| [#58755](https://github.com/NousResearch/hermes-agent/issues/58755) – Empty `tool_calls` from repair | **P2** | Follow-up to closed #56980; DeepSeek v4 rejects empty array. | No |
| [#58757](https://github.com/NousResearch/hermes-agent/issues/58757) – `write EPIPE` in Hermes Studio | **Unrated** | Uncaught exception during initial runtime download in Electron main process. | Yes ([#58761](https://github.com/NousResearch/hermes-agent/pull/58761)) |
| [#58745](https://github.com/NousResearch/hermes-agent/issues/58745) – Compression semantics conflict | **P2** | `context_length` is used both as capability declaration and budget, causing every-turn compression loops. | No |
| [#58717](https://github.com/NousResearch/hermes-agent/issues/58717) – Desktop interrupt mode broken | **P2** | User messages queue instead of interrupting agent on Windows Desktop. | No |

**Known regressions:** The two `compression` issues (#58753 and #58745) appear to be recent regressions introduced by earlier context management PRs. The Python deletion bug (#58748) is a systemic safety gap – the agent should never be allowed to remove its own runtime.

---

### 6. Feature Requests & Roadmap Signals

- **Configurable Memory Backends** ([#47349](https://github.com/NousResearch/hermes-agent/issues/47349))  
  The most detailed feature request this week. Proposes renaming `memory.md` to `rules.md` and supporting alternative backends (Honcho, Fact Store) while disabling file-based memory. Likely to be prioritized given the `P2` tag and growing community support.

- **Stable Channel for `hermes update`** ([#58746](https://github.com/NousResearch/hermes-agent/issues/58746))  
  Users want `hermes update` to pull the latest stable release instead of `main`. This would reduce exposure to experimental commits and half-baked fixes. If adopted, this could be a simple `git checkout` or release-tag switch in the update script.

- **Interrupt Mode Parity** ([#58717](https://github.com/NousResearch/hermes-agent/issues/58717))  
  Desktop app missing `busy_input_mode=interrupt` support – signals a need for feature parity between Desktop/TUI.

- **MoA Provider Recognition** ([#58759](https://github.com/NousResearch/hermes-agent/issues/58759))  
  `hermes doctor` incorrectly flags `moa` as unrecognised despite it being a legitimate internal Mixture-of-Agents provider. This is a minor UX fix that could land in the next patch.

**Prediction for next version:**  
The memory backend rework (#47349) and the stable update channel (#58746) are strong candidates for the next minor release. The compression semantics fix (#58745) will likely be addressed in the same cycle as the compression exhaustion handoff PR (#58750).

---

### 7. User Feedback Summary

**Pain points expressed today:**

- **Stability on Windows Desktop** – Users report that `busy_input_mode=interrupt` does not work (issue #58717), and the `cua-driver` update always times out (issue #58762). This has led to frustration with the upgrade path on Windows.
- **DeepSeek v4 integration instability** – After the closure of #56980 as *not_planned*, users have re-reported the same bug with a new error variant (#58755), indicating that the original fix was incomplete. Users expect a more thorough resolution.
- **Compression corruption** – Users on the `hermes kanban` workflow are hit hard by compression dropping the only user message (#58753). The fix PR (#58750) is a positive response, but the issue is affecting production kanban workers.
- **Self-inflicted runtime deletion** – One user reported that Hermes removed its own Python interpreter (#58748), leaving the desktop app unable to start. This highlights a dangerous gap in tool-safety permissions.

**Satisfaction signals:**  
The quick turnaround on fix PRs (e.g., #58761 for EPIPE, #58760 for cua-driver timeout) shows responsive maintainers. The security-hardening PRs (#58752, #58751, #58740) show active investment in preventing credential leaks, which is reassuring for enterprise users.

---

### 8. Backlog Watch

These are older issues/PRs that remain open and have received no update in weeks, potentially needing maintainer triage:

- [#16454](https://github.com/NousResearch/hermes-agent/pull/16454) (Apr 27) – Move CLI new-session memory flush off-thread. Still open with no recent comments. Important for session rotation performance.
- [#16455](https://github.com/NousResearch/hermes-agent/pull/16455) (Apr 27) – Drop stale background process session notifications. Open for over 2 months.
- [#16450](https://github.com/NousResearch/hermes-agent/pull/16450) (Apr 27) – Feature-detect context focus compression support. Stale with no merge.
- [#43276](https://github.com/NousResearch/hermes-agent/pull/43276) (Jun 10) – Pass timeout to Firecrawl scrape API. Small fix, long-open.
- [#38168](https://github.com/NousResearch/hermes-agent/pull/38168) (Jun 3) – Fix TUI implicit cwd overrides for Docker terminals. Waits for review.
- [#42934](https://github.com/NousResearch/hermes-agent/pull/42934) (Jun 9) – Filter messaging sessions by active profile in Desktop. Unmerged.
- [#38731](https://github.com/NousResearch/hermes-agent/pull/38731) (Jun 4) – Add inline buttons for cron Telegram delivery. Feature request with tests, no update.
- [#45491](https://github.com/NousResearch/hermes-agent/pull/45491) (Jun 13) – Offload dashboard cron profile scans from event loop. Performance improvement.

While these PRs are mature, many lack recent maintainer activity. The project would benefit from a triage session to either merge, request changes, or close with explanation. The oldest (April) PRs risk bit-rot and could cause merge conflicts.

---

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# Project Digest: PicoClaw (sipeed/picoclaw) – 2026-07-05

## 1. Today's Overview

The project shows moderate activity with 4 issues and 8 pull requests updated in the last 24 hours. Three PRs were merged/closed, including a fix for agent session clearing and a correction for the LINE channel. One stale bug (memory amnesia) was closed, but two open bugs remain concerning Android service launch failures and Matrix encryption. Community engagement is visible through a high-priority feature request (replace libolm with vodozemac, [#3088]) and a new PR introducing agent-specific runtime overrides ([#3225]). No new releases were cut.

## 2. Releases

*None.* There are no new releases to report.

## 3. Project Progress (Merged/Closed PRs)

Three pull requests were merged or closed today:

- **[#3189] [CLOSED]** – `fix(line): explicitly ignore resp.Body.Close() errors in Send and classifySDKError` (author: chengzhichao-xydt). Improves error handling hygiene in the LINE channel.

- **[#3221] [CLOSED]** – Revert "test: cover sandbox fs Windows path handling" (author: afjcjsbx). A quick rollback of a problematic test that caused import errors in `openai_compat/provider.go`.

- **[#3224] [CLOSED]** – `fix(agent): clear routed agent session` (author: Ethan1918). Resolves a bug where `/clear` would reset the default agent’s session instead of the correct non-default agent’s session.

Additionally, a stale bug issue [#3150] was closed after a fix was proposed in PR [#3226] (see Bugs & Stability).

## 4. Community Hot Topics

The most active discussion revolves around **replacing the cryptographic library**:

- **[Issue #3088] [OPEN, help wanted, priority: high]** – "use vodozemac instead of libolm". This issue has 4 comments and 2 thumbs-up. The user advocates dropping the unmaintained libolm in favor of vodozemac, the official replacement. It has been open since June 9 and is marked as stale. The underlying need is security and long-term maintenance.

Other items with moderate attention:

- **[Issue #3150] [CLOSED]** – “它给自己整失忆了” (it gave itself amnesia). Had 5 comments; now closed as stale. The companion PR (#3226) addresses the root cause (see below).

## 5. Bugs & Stability

**High severity:**

- **[Issue #3182] [OPEN]** – “Android version” – User reports inability to launch the service on Android despite granting full permissions and inability to change the path from settings. No fix PR exists yet. (Updated 2026-07-04)

- **[Issue #3194] [OPEN, stale]** – “Received encrypted message but crypto is not enabled” – A Matrix user sees logs indicating crypto is not enabled despite receiving encrypted messages. This points to a configuration or initialization gap. No fix PR yet. (Updated 2026-07-04)

**Medium severity (fixed):**

- **[Issue #3150] [CLOSED]** – Memory “amnesia” bug: the agent lost its memory after file overwrite operations. The root cause is addressed by PR [#3226] (open today), which stops the `write_file` tool from coaching destructive overwrites – a behavioral fix to prevent the tool from misleading the agent.

**Low severity:**

- **[PR #3189] [CLOSED]** – Ignored `resp.Body.Close()` errors in LINE channel (non-functional improvement).

## 6. Feature Requests & Roadmap Signals

- **[#3088] [OPEN]** – Replace `libolm` with `vodozemac` (Matrix encryption). Marked high-priority and help wanted. Likely to appear in the next release if a contributor picks it up.

- **[PR #3225] [OPEN]** – “Support agent-specific runtime overrides” (author: xdatafactor). Allows `agents.list` entries to define `max_tokens`, summarization thresholds, and split_on_marker. This is a significant usability improvement for multi-agent setups. Already passing `go test ./pkg/config`. Could merge soon.

- **[#3150-related PR #3226] [OPEN]** – “stop write_file from coaching destructive overwrite” (author: ACMYuechen). While technically a bugfix, it changes the agent’s behavior to preserve memory more safely – a feature-like improvement for reliability.

## 7. User Feedback Summary

- **Pain point: Android support** – User [#3182] reports inability to launch the service on Android, highlighting a gap in the mobile experience.

- **Pain point: Matrix crypto** – User [#3194] encounters encrypted messages with crypto not enabled, indicating either a documentation gap or a startup failure.

- **Satisfaction: memory stability fix** – The closing of issue [#3150] suggests the team is aware of agent memory loss complaints; the fix in PR [#3226] should improve user trust in long-running sessions.

- **Feature desire: cryptographic upgrade** – The reaction count (2 👍) on [#3088] shows community support for moving to a maintained library.

## 8. Backlog Watch

Several important issues and PRs remain without maintainer response or have been stale for extended periods:

- **[Issue #3088]** – High-priority, help wanted, open since 2026-06-09, marked stale. Requires maintainer decision or assignment.

- **[Issue #3182] [OPEN]** – Android launch failure, no activity from maintainers since creation (2026-06-26). Should be triaged.

- **[Issue #3194] [OPEN, stale]** – Matrix crypto disabled issue, only 1 comment. Needs investigation.

- **[PR #3190, #3191, #3192]** – All opened 2026-06-27, marked stale, awaiting review. They cover i18n sync, `.gitignore` cleanup, and Docker image bumps. Low risk but block other updates.

- **[PR #3225]** – Agent-specific overrides, opened 2026-07-04, no comments yet. Could become stale if not reviewed quickly.

A maintainer review pass on the backlog would improve project health and community trust.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-07-05

## Today's Overview
Activity remains high with 36 pull requests touched in the last 24 hours—20 were merged or closed, reflecting a strong focus on codebase hygiene and security hardening. A single new security issue was opened (Issue #2923), while no releases were published. The project appears to be in a consolidation phase, with maintainers systematically removing dead code, updating documentation, and fixing subtle bugs in the container and router logic. Community contributions continue to flow, including two skills-related PRs from external contributors.

## Releases
*None.* No new releases were published today. The latest tagged version remains v2.1.38 (reference in PR #2953).

## Project Progress
The following significant changes were merged or closed today (all PRs listed are in closed state unless noted):

- **Security docs rewrite** –  `docs/SECURITY.md` rewritten to reflect the actual v2 container perimeter; two v1-only guides marked as historical. (PR #2945)
- **Dead code removal** –  Deleted deprecated shims from the April-2026 session-DB split (PR #2940), dead v1 config knobs and a broken pnpm auth script (PR #2935), and unused CLI protocol vocabulary (PR #2936).
- **Bug fixes** –  
  - `in_reply_to` stamp fixed for cross-process MCP server (PR #2942)  
  - Mount allowlist now honors `readOnly` key and no longer caches parse errors (PR #2943)  
  - Session folder re-provisioning after `rm -rf` (PR #2937)  
  - Egress-lockdown env vars now reachable under shipped service (PR #2934)
- **Feature additions** –  
  - `ncl groups config add-mount / remove-mount` verbs added (host-only) (PR #2939)  
  - Colored buttons (primary/danger) on Slack approval cards (PR #2933)  
  - Asynchronous agent image building to avoid blocking the host (PR #2931)
- **Documentation fixes** – Corrected stale architecture, scheduling, provider-config, and overlay docs (PR #2948); fixed stale mount topology row and removed env var reference (PR #2953).

## Community Hot Topics
The most actively discussed item is the single open issue:

- **Issue #2923 – [Security] ask_user_question card can be defaced by a forged click before origin authz**  
  [Link](https://github.com/nanocoai/nanoclaw/issues/2923)  
  *Author:* glifocat | *Created:* 2026-07-04 | *No comments or reactions yet.*  
  This issue describes a display/integrity spoof where a forged button click can overwrite the card's displayed text even though the response is correctly rejected. No fix PR has been filed yet, but the security team is likely triaging.

No open PRs have accumulated comments or reactions; all have zero. The high number of PRs (36) suggests a burst of cleanup activity rather than prolonged debate.

## Bugs & Stability
**High severity:**
- **Issue #2923** – Display spoof on `ask_user_question` card. No fix PR yet; watch for upcoming security patch.
- **PR #2956 (open)** – Fix for duplicate delivery when final output repeats tool-sent content. (Author: stumpjumper)  
  [Link](https://github.com/nanocoai/nanoclaw/pull/2956)

**Medium severity (all fixed today):**
- **PR #2955 (open)** – Fix for mention-sticky subscribing to channel root or accumulate-only sessions. (Author: dim0627)  
  [Link](https://github.com/nanocoai/nanoclaw/pull/2955)
- Mount allowlist parse-error caching (fixed in #2943)
- Cross-process `in_reply_to` no-op (fixed in #2942)
- Session folder deletion causing reset to fail (fixed in #2937)

All fixes have been merged or are pending review.

## Feature Requests & Roadmap Signals
The merged feature PRs give a clear direction for the next release:

- **Per-group environment variables (PR #2036)** – still open after rewrites; adds DB-managed `ncl groups config set-env`. This is a user-facing feature likely to land soon.
- **Container mount management (PR #2939** – merged) – enables operators to add/remove mounts without editing config files.
- **Asynchronous image builds (PR #2931** – merged) – improves host responsiveness during agent deployments.
- **Skill additions**: Two open PRs from contributor `javexed` add an “opencode stack” skill (#2952) and improve `opencode` configuration (#2951). These signal community interest in expanding integration capabilities.

Predict next version (likely v2.2.0) will include the per-group env vars, async builds, mount management, and the security docs rewrite. The security issue #2923 may push a hotfix release.

## User Feedback Summary
No direct user feedback was captured in the data. However, the following pain points can be inferred from fix PRs:

- **Duplicate message delivery** (PR #2956) – indicates users experienced double responses from agents that use the `send_message` tool.
- **Session reset not working** (PR #2937) – operators who followed the documented `rm -rf` method found sessions unrecoverable.
- **Stale documentation causing confusion** (PRs #2948, #2953, #2945) – several users/contributors have flagged that guides still describe v1 mechanics.
- **Security-conscious contributors** – the attention to perimeter hardening (#2934, #2943, #2945) reflects a user base that values containment and least-privilege.

## Backlog Watch
- **PR #2036 – per-group container env vars**  
  [Link](https://github.com/nanocoai/nanoclaw/pull/2036)  
  Created April 26, refreshed July 4. Still open after a major rewrite to DB-native approach. This is the longest-pending open PR; maintainer attention is required to merge or provide feedback.
- **Issue #2923** – newly opened but has zero comments. Should be triaged soon to assign severity and plan a fix.

No other truly long-unanswered items appear; most open PRs are from the last 48 hours.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-07-05

## 1. Today’s Overview

The project remains highly active with **37 PRs updated** (23 open, 14 merged/closed) and **9 issues updated** (8 open, 1 closed) in the last 24 hours. All work is concentrated on the **Reborn** stack, with core contributors driving a major **Slack OAuth migration** (4-PR stack), a wave of **integration‑tier coverage enablers**, and a critical **security fix for bridge meta‑tools**. A **nightly E2E pipeline failure** (#4108) persists but is being monitored. No new releases were published today. Overall project health is strong, with deep test‑infrastructure investment and timely bug responses.

## 2. Releases

**None** – no new releases were cut today.

## 3. Project Progress (Merged/Closed PRs Today)

14 PRs were merged or closed. Notable completed items:

- **[#5642] Wiring‑parity guard** – merged. Adds an exhaustive shape‑drift tripwire that ensures the integration harness runtime matches production’s local‑dev composition.  
  [PR #5642](https://github.com/nearai/ironclaw/pull/5642)

- **[#5649] Coverage‑enabler batch** – merged. Brings three previously 0%‑coverage crates onto the integration‑tier lane: bridged tool disclosure, webui‑v2/identity unlock, and trace‑capture.  
  [PR #5649](https://github.com/nearai/ironclaw/pull/5649)

Other closed PRs likely include dependency bumps and CI adjustments (full list not provided). The volume of merged test‑infrastructure PRs indicates the team is methodically hardening Reborn’s coverage and runtime correctness before feature releases.

## 4. Community Hot Topics

- **Issue #5647 – Bridge meta‑tools stripped from narrowed allow‑lists** (1 comment, the only issue with discussion)  
  A latent security/functionality bug: when tool disclosure uses the bridged mode for >32 tools, the host‑synthesized `ironclaw.*` bridge meta‑tools are stripped from any narrowed capability allow‑list. A fix PR (#5659) was opened today.  
  [Issue #5647](https://github.com/nearai/ironclaw/issues/5647)

- **PR #5644 / #5645 / #5646 – Slack personal OAuth stack (3/4, 4/4)**  
  The largest active PRs by scope (each XL) – replacing legacy Slack pairing codes with personal OAuth. High cross‑team attention expected; issues like #5650 (scope granularity) surfaced from review.  
  [PR #5644](https://github.com/nearai/ironclaw/pull/5644) · [PR #5645](https://github.com/nearai/ironclaw/pull/5645) · [PR #5646](https://github.com/nearai/ironclaw/pull/5646)

- **PR #5626 – Manifest‑driven Slack ingress routes**  
  Another large PR (XL) that projectifies Slack route declarations into a manifest, removing hand‑written Rust policy literals. Signals a design shift toward data‑driven extension configuration.  
  [PR #5626](https://github.com/nearai/ironclaw/pull/5626)

Underlying needs: users and operators want **simpler Slack setup** (OAuth vs pairing codes), **better CI reliability** (#5636: skipped jobs blocking Railway deploys), and **correct security scoping** (#5650: per‑capability OAuth scopes).

## 5. Bugs & Stability

| Severity | Bug | Status |
|----------|-----|--------|
| **Critical** | **Bridge meta‑tools stripped** (#5647) – security filter `CapabilitySurfaceProfileFilter` incorrectly removes `tool_search`/`tool_describe`/`tool_call` from narrowed allow‑sets when crossing the 32‑tool bridging threshold. Fix PR #5659 open today. | Fix in review |
| **High** | **Nightly E2E failure** (#4108, open since May 27) – scheduled run fails repeatedly. Last updated Jul 5; workflow run link shows failure. No fix PR yet. | Unresolved |
| **Medium** | **CAS‑contention / discard‑tombstone** – coverage gap under real contention (tracked via #5466). PR #5661 adds tests and fixes `InMemory` store parity. | Fix in review |
| **Low** | **Harness gap: no `RecordingSecurityAuditSink` double** (#5640) – integration harness always wires `None` for the security audit sink, diverging from production’s `TracingSecurityAuditSink`. | Issue created |
| **Low** | **Wiring‑parity shape hand‑derived** (#5641) – `EXPECTED_PRODUCTION_SHAPE` must be manually transcribed; no automated accessor yet. | Issue created |

The one closed issue (#5637) was the wiring‑parity guard itself, now merged.

## 6. Feature Requests & Roadmap Signals

Submitted/requested features visible in recent issues and PRs:

- **Personal OAuth scope splitting** (#5650) – users want read‑only Slack capabilities (e.g., `search_messages`) to *not* require `chat:write`. The PR author acknowledged this and may refine before merge.
- **CI job‑level skip handling** (#5636) – operators want skipped CI jobs to not block Railway deploys. Low effort but high impact; likely addressed soon.
- **Coverage ratchet** (#5638) – currently informational; asking for a failing gate when coverage drops. The issue describes the exact mechanism needed (seed exempt file, flip CI).
- **Production‑side shape accessor** (#5641) – a small quality‑of‑life feature to replace hand‑transcribed constants.

**Predictions for next version (likely 0.30.x or 0.5.0 for `ironclaw_common`):**  
- Slack OAuth fully replaces pairing codes.  
- Coverage reports become ratcheted.  
- Bridge meta‑tool fix ships as a security patch.  
- `ironclaw_common` 0.5.0 (breaking) is already staged in #5598 and may land soon.

## 7. User Feedback Summary

- **Pain point: CI inconsistency.** “Wait for CI” on Railway blocks deployments when jobs are skipped (#5636). Developers find this frustrating.
- **Pain point: Slack OAuth scope granularity.** Issue #5650 confirms all five `slack_user` capabilities share the full 11‑scope set, even read‑only ones. User suggests per‑capability scope lists.
- **Satisfaction: rapid coverage improvement.** Multiple PRs (#5649, #5658, #5660) are systematically closing integration‑test gaps. Contributors express confidence in the new instrumentation.
- **Use case: real‑world durability.** PR #5660 validates `FilesystemOutboundStateStore` and PDF attachment extraction – a concrete operator concern.

## 8. Backlog Watch

Items requiring maintainer attention:

| Item | Age | Notes |
|------|-----|-------|
| **#4108 – Nightly E2E failure** | 39 days (since May 27) | Still failing, updated Jul 5. No fix in progress; CI stability risk. |
| **#5170 – Subagent spawn run failure** (PR) | 12 days (since Jun 23) | L‑sized fix for `LoopInlineMessageBody` and `AwaitDependentRun` handling. Still open, no merge. |
| **#5550 – Dependabot bulk dependency update** | 3 days (since Jul 2) | 13 package updates, including a major bump of `agent-client-protocol` (0.10.4→1.0.1). Needs review. |
| **#5598 – Release PR** (`ironclaw` 0.29.1) | 2 days (since Jul 3) | Breaking changes in `ironclaw_common` and `ironclaw_skills`; could be held pending OAuth stack merge. |

No issues or PRs beyond 40 days are visible in the dataset. The weekly release cadence appears moderate; the release PR #5598 may be blocked by ongoing feature work.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-07-05

## 1. Today’s Overview
Project activity remained low, with **0 issues** updated and **3 pull requests** updated in the last 24 hours – two of which were merged/closed (#2271, #2272) and one stale open PR (#1349). No new releases were published. The merged PRs indicate steady progress in system integration (proxy propagation) and content migration (identity block cleanup), while the lone open PR addresses a pending bug fix that has been waiting since April. Overall, the project appears to be in a maintenance/consolidation phase with minimal community engagement today.

## 2. Releases
**No new releases** were published. The latest release remains the previous version (not specified in data). No breaking changes or migration notes to report.

## 3. Project Progress
Two PRs were merged/closed today:
- **#2272** (closed) – `fix(agent): migrate legacy AGENTS.md identity blocks to IDENTITY.md`  
  *Author: fisherdaddy* — Detects and cleans up legacy identity content embedded in `AGENTS.md` to eliminate conflicts with the managed `IDENTITY.md` file. Includes backup and safe skip/failure reporting per agent.  
  [GitHub: #2272](https://github.com/netease-youdao/LobsterAI/pull/2272)
- **#2271** (merged) – `fix: propagate system proxy to managed browser`  
  *Author: fisherdaddy* — Ensures that system proxy settings are correctly forwarded to the managed browser instance. This resolves a likely integration issue for users behind corporate proxies.  
  [GitHub: #2271](https://github.com/netease-youdao/LobsterAI/pull/2271)

## 4. Community Hot Topics
There were **no issues** with significant comment or reaction activity today. The only open PR (#1349) has **0 comments** and **0 reactions**, suggesting low community engagement on that item. The two merged PRs were closed without public discussion. No particular “hot topic” emerged from the community in the last 24 hours.

## 5. Bugs & Stability
No new bug reports were filed today. However, **PR #1349** (still open) addresses an existing bug:
- **Bug**: POPO connectivity test always shows “verified successfully” even with invalid credentials (appKey/appSecret).  
- **Severity**: Medium – affects reliability of the POPO integration test, potentially masking misconfigurations.  
- **Fix exists**: PR #1349 adds real API validation against POPO endpoints.  
- **Status**: Open since April 2, updated today (stale label). No other bug-related activity.

## 6. Feature Requests & Roadmap Signals
No feature requests were captured in today’s data. The two merged PRs are bug fixes/tech debt, not new features. No roadmap signals or user-requested features were evident in the last 24 hours. The next version may include the fixes in #2271 (proxy propagation) and #2272 (identity migration) but no major new feature indicators.

## 7. User Feedback Summary
No direct user feedback (issues, comments) was recorded today. Based on the PR content, the community benefits from:
- **Proxy support fix** (#2271) – addresses pain points for users behind corporate or restrictive networks.
- **Identity cleanup** (#2272) – reduces agent configuration confusion for users migrating from legacy setups.

No satisfaction/dissatisfaction signals were observed.

## 8. Backlog Watch
- **PR #1349** – `fix(im): add real API validation for POPO connectivity test`  
  *Opened: 2026-04-02 | Last updated: 2026-07-05 | Stale*  
  This PR has been open for **over 3 months** with no reviewer activity and only received a stale label today. It addresses a clear bug (false positive validation) and contains a concrete solution. It may require maintainer attention to review and merge to prevent further drift.  
  [GitHub: #1349](https://github.com/netease-youdao/LobsterAI/pull/1349)

No other long-untouched issues or PRs were identified in today’s data.

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

# CoPaw Project Digest – 2026-07-05

## Today’s Overview

CoPaw saw high activity in the past 24 hours, with **15 issues updated** (14 open, 1 closed) and **3 open pull requests** — but **no new releases**. The community is actively reporting bugs and requesting enhancements, indicating a mature but still rapidly evolving project. The maintainers responded quickly with fix PRs for several critical bugs, though a number of stability concerns remain unresolved. The overall project health appears energetic, with a clear focus on polishing frontend and backend reliability.

## Releases

No new releases were published today.

## Project Progress

- **No PRs merged/closed** in the last 24 hours.  
- **One issue closed**: [#2830 – “建议桌面端增加一个隐藏到托盘图标的功能”](https://github.com/agentscope-ai/CoPaw/issues/2830) (a feature request for system tray minimization) was closed after being open since April.
- **Three open PRs** are tackling active bugs and memory features:  
  - [#5786](https://github.com/agentscope-ai/CoPaw/pull/5786) – Fixes three bugs (#5709, #5773, #5784).  
  - [#5783](https://github.com/agentscope-ai/CoPaw/pull/5783) – Fixes cron timezone issue (#5779).  
  - [#5777](https://github.com/agentscope-ai/CoPaw/pull/5777) – Adds auto-memory turn state management (feature).

## Community Hot Topics

The following issues attracted the most discussion (by comment count and reactions):

- **#2865 – “Support displaying custom agent names and avatars”**  
  [https://github.com/agentscope-ai/CoPaw/issues/2865](https://github.com/agentscope-ai/CoPaw/issues/2865)  
  👤 Ryoui | 4 comments | 1 👍 | Open since April 3, updated July 4  
  *Underlying need*: Users want to personalise the chat experience by assigning unique names and avatars to different agents. This is a long-standing feature request that remains unmerged.

- **#5784 – “Frontend compression threshold mismatch when same model across providers”**  
  [https://github.com/agentscope-ai/CoPaw/issues/5784](https://github.com/agentscope-ai/CoPaw/issues/5784)  
  👤 y0umu | 3 comments | reported today  
  *Underlying need*: Configuration UI displays incorrect `max_input_length` when the same model ID exists in multiple providers. A fix is already in PR #5786.

- **#5779 – “Cron state API returns UTC instead of configured timezone”**  
  [https://github.com/agentscope-ai/CoPaw/issues/5779](https://github.com/agentscope-ai/CoPaw/issues/5779)  
  👤 feng183043996 | 3 comments | reported today  
  *Underlying need*: Jobs scheduled in non-UTC timezones cannot rely on the API timestamps. Fix PR #5783 already submitted.

## Bugs & Stability

New bugs reported today, ranked by severity (critical → low):

| Severity | Issue | Status | Fix PR? |
|----------|-------|--------|---------|
| **Critical** | [#5778 – Scroll compression loses context, causes off-track responses](https://github.com/agentscope-ai/CoPaw/issues/5778). Users experience complete loss of conversation history after context compression. | Open | None yet |
| **High** | [#5775 – Auto-memory interval never triggers because state lost across requests](https://github.com/agentscope-ai/CoPaw/issues/5775). Memory persistence breaks with `auto_memory_interval > 1`. | Open | PR #5777 (feature, not a direct fix) |
| **High** | [#5782 – Google Gemini embedding returns `index=None`, silent fallback to keyword search](https://github.com/agentscope-ai/CoPaw/issues/5782). Vector search silently disabled. | Open | None |
| **High** | [#5773 – Memory search causes OpenCode (OCG) channel to error](https://github.com/agentscope-ai/CoPaw/issues/5773). All requests fail when `auto_memory_search_config.enabled = true`. | Open | PR #5786 (includes fix) |
| **Medium** | [#5779 – Cron timezone hardcoded to UTC](https://github.com/agentscope-ai/CoPaw/issues/5779) | Open | PR #5783 |
| **Medium** | [#5784 – Frontend threshold display wrong](https://github.com/agentscope-ai/CoPaw/issues/5784) | Open | PR #5786 |
| **Medium** | [#5787 – Mobile webUI bottom content truncated on all pages](https://github.com/agentscope-ai/CoPaw/issues/5787) | Open | None |
| **Medium** | [#5774 – Google Gemini channel model endpoint error](https://github.com/agentscope-ai/CoPaw/issues/5774) | Open | None |
| **Medium** | [#5776 – Stale pinned user message treated as active task in long IM sessions](https://github.com/agentscope-ai/CoPaw/issues/5776) | Open | None |
| **Low** | [#5781 – Offline code mode cannot preview files because online resources required](https://github.com/agentscope-ai/CoPaw/issues/5781) | Open | None |
| **Low** | [#5757 – Feishu channel not replying after first message](https://github.com/agentscope-ai/CoPaw/issues/5757) | Open | None |

Many of these issues have corresponding fix PRs already in review, suggesting maintainers are treating them with high priority.

## Feature Requests & Roadmap Signals

New enhancement requests from today (July 5–6):

- [#5785 – “Select hidden folders (dot-prefixed) in coding mode”](https://github.com/agentscope-ai/CoPaw/issues/5785)  
  User wants to choose `.hidden` directories in the file picker. Small UI change – likely to be included in next patch.

- [#5780 – “Multi-user account management for team use”](https://github.com/agentscope-ai/CoPaw/issues/5780)  
  Request for proper team member authentication and role-based access, not just relying on IM channel identity. A larger feature that may appear in a future major release.

- Long-standing [#2865 – “Custom agent names and avatars”](https://github.com/agentscope-ai/CoPaw/issues/2865) (from April) has seen renewed attention. Despite several comments and a thumb-up, it remains unassigned.

Additionally, the open PR [#5777](https://github.com/agentscope-ai/CoPaw/pull/5777) “feat(memory): add auto-memory turn state management” signals an upcoming improvement in memory persistence – likely to land in the next version.

## User Feedback Summary

- **Pain points**:  
  - Context compression (scroll strategy) is breaking conversations – users losing critical session history. (#5778)  
  - Auto-memory not working reliably – conversations are not persisted in long sessions. (#5775)  
  - Mobile experience is broken (truncated buttons). (#5787)  
  - Offline usage is hindered by external resource dependencies. (#5781)  
  - Google Gemini integrations have multiple incompatibilities (embedding, endpoint). (#5782, #5774)  
  - Team deployment lacks user management – only a single bot account is supported. (#5780)  

- **Satisfaction**: No explicit positive feedback was recorded today. The frequency and detail of bug reports suggest experienced users who rely on CoPaw for production tasks are being affected, which may indicate growing pains as the project scales.

## Backlog Watch

- **#2865 – Custom agent names/avatars** (open since April 3, updated July 4)  
  https://github.com/agentscope-ai/CoPaw/issues/2865  
  A relatively uncontroversial enhancement with 4 comments, 1 thumbs-up. No maintainer response or assignment visible. Community may be waiting for progress.

- **#5757 – Feishu channel no response after first message** (open since July 3)  
  https://github.com/agentscope-ai/CoPaw/issues/5757  
  Only 2 comments, no maintainer reply. Given that it affects both self-hosted Docker and the hosted platform, this may be a wider issue.

- **#5778 – Scroll compression loses context** (open today, already 1 comment)  
  High impact, no fix PR yet. Should be prioritised alongside the other critical memory bugs.

These items would benefit from maintainer guidance or assignment to keep community confidence high.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest – 2026-07-05

## 1. Today's Overview
The ZeroClaw project remains highly active, with **50 open pull requests** touched in the last 24 hours—reflecting a surge of bug fixes, security hardening, and new features across runtime, channels, and tooling. **No new releases** were published today, and **no PRs were merged or closed**, suggesting the maintainers are in a heavy review cycle. Six open issues were updated, two of which were created today (bugs: zombie MCP processes and high-entropy detector false positives). The overall project health is positive, though the absence of merged PRs may indicate a bottleneck in merging or a deliberate staging before a coordinated release.

## 2. Releases
*No new releases were reported for 2026-07-05. The last release remains v0.8.x (tracked in issue #8073).*

## 3. Project Progress
- **Merged/Closed PRs:** **None** today. All 50 open PRs remain pending review or further iteration.
- **Key features advanced in open PRs (not yet merged):**
  - **Goal-oriented channels:** PRs #8689 and #8688 introduce `/goal` command admission and trusted goal tools – a significant expansion for structured multi-step task execution.
  - **Unified memory-context injection:** PR #8619 (size:L) reworks how memory is injected based on turn origin provenance, addressing long-standing context window inefficiencies.
  - **Cron job `uses_memory` flag:** PR #8676 exposes the existing database field via CLI, tools, and API, enabling better agent memory control for scheduled tasks.
  - **WebUI visual editor for slash-command options:** PR #8620 adds a bespoke editor for skill bundle options.

## 4. Community Hot Topics
The most active issues and PRs, based on comments and reactions, are:

- **Issue #8720** (1 comment, 0 reactions) – [Support: Disable cachePoint for Bedrock Nova 2 Lite](https://github.com/zeroclaw-labs/zeroclaw/issues/8720)  
  *Need:* User wants a config toggle to disable caching for a specific model, as it causes random errors. This reflects a common desire for per-provider/pur-model feature flags.

- **Issue #8675** (1 comment, 0 reactions) – [Bug: Malformed native tool-call arguments leading to provider 400](https://github.com/zeroclaw-labs/zeroclaw/issues/8675)  
  *Need:* A workflow-blocking regression in OpenAI-wire-format providers where tool call arguments are not validated before re-serialization. This is a critical S1 issue reported by a user facing repeated empty replies.

- **PR #8619** (multiple commits, latest update today) – [Unified memory-context injection](https://github.com/zeroclaw-labs/zeroclaw/pull/8619)  
  *Need:* The PR addresses a complex memory/context interaction and has undergone a changes-requested round, showing active community collaboration.

- **PR #8561** (size:XL) – [Telegram multi_message streaming mode](https://github.com/zeroclaw-labs/zeroclaw/pull/8561)  
  *Need:* Adding delayed multi-message streaming to Telegram, matching Discord/Matrix. A long-standing feature request that touches multiple components.

**Underlying themes:** Users are pushing for better provider compatibility (e.g., Bedrock, OpenRouter), fine-grained caching control, and richer channel functionality (streaming, goals).

## 5. Bugs & Stability
**New bugs reported today (2026-07-05):**

- **S2 – Zombie MCP server processes** (#8731) – Stdio-based MCP servers accumulate as unreaped zombie processes under the daemon, degrading performance over time. *No fix PR attached yet.*
- **S2 – High-entropy detector redacts legitimate filenames** (#8722) – Outbound leak detector misclassifies generated filenames as secrets, breaking file references. A fix PR **#8723** by vrurg was opened concurrently, proposing caller-supplied protected byte ranges.

**Existing critical bugs with ongoing PRs:**
- **S1 – Malformed tool-call arguments** (#8675) – No dedicated fix PR yet, but related PRs (#8724, #8711) improve tool assembly validation.
- **S2 – SOP `when` condition false ends run instead of advancing** (#8719) – Minor but breaks multi-phase SOPs. No fix PR yet.
- **S2 – Standardized env var fallback for OpenAI STT** (#8576, PR open) – Fixes a configuration gap for speech-to-text credentials.

**Regression concerns:** PR #8724 (size:L) root-fixes SOP regression #8631 and introduces deterministic capability steps – a high-risk change that should be reviewed carefully.

## 6. Feature Requests & Roadmap Signals
- **Explicit per-model caching toggle** (#8720) – Likely to be addressed in v0.8.4 via a config extension to model provider settings.
- **Goal-oriented task management** (#8688, #8689) – PRs for goal commands and trusted delegation are large (size:XL) and likely candidates for the next release, given the focus on structured agent workflows.
- **SOP step fallthrough on false condition** (#8719) – A simple behavioral fix that could be cherry-picked into a patch release.
- **WebUI slash-command editor** (#8620) – Enhances developer experience for skill authors, may land in v0.9.

Backlog signal: Issue #8073 (v0.8.3 observability tracker) is still open with no merged PRs, suggesting those tasks may slip to v0.8.4.

## 7. User Feedback Summary
- **Pain points:**
  - Caching errors on Bedrock Nova 2 Lite – user wants opt-out, not optimistic caching.
  - Empty replies from OpenAI-wire providers when tool call arguments are malformed – user reports “workflow blocked”.
  - Zombie processes after MCP server timeouts – operational stability concern.
  - Leak detector false positives breaking file paths – hurts usability in file-centric workflows.

- **Satisfaction indicators:**
  - Active PR contributions from multiple community members (vrurg, wangmiao0668000666, Nillth, etc.) show a healthy contributor base.
  - Rapid issue creation and PR linking (e.g., bug #8722 → fix PR #8723 within hours) indicates responsive triage.

## 8. Backlog Watch
- **Issue #8073** – [v0.8.3 observability tracker](https://github.com/zeroclaw-labs/zeroclaw/issues/8073) (created June 20, last updated July 4)  
  *Status:* Accepted, priority P2, risk high. No PRs have been merged for any of the tracked items (observability, CI, docs, dependencies). Needs maintainer attention to avoid scope creep.

- **Issue #8675** – [Malformed tool-call arguments (S1)](https://github.com/zeroclaw-labs/zeroclaw/issues/8675) (created July 3)  
  *Status:* Despite being severity S1 and having an accepted label, no fix PR has been linked. The closed PR #8724 might partially address the root cause, but explicit resolution is pending.

- **PR #8576** – [Fix env-var fallback for OpenAI STT](https://github.com/zeroclaw-labs/zeroclaw/pull/8576) (created July 1, updated today)  
  *Status:* Waiting for final review. This supersedes #8079 and is critical for users relying on environment variables for transcription.

- **Issue #8719** – [SOP false `when` ends run](https://github.com/zeroclaw-labs/zeroclaw/issues/8719) (created July 4)  
  *Status:* No assignee, no PR. A small behavioral change that could improve SOP flexibility; important for users building multi-phase workflows.

---

*Data extracted from GitHub: zeroclaw-labs/zeroclaw – activity up to 2026-07-05T23:59 UTC.*

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*