# OpenClaw Ecosystem Digest 2026-07-28

> Issues: 250 | PRs: 500 | Projects covered: 13 | Generated: 2026-07-28 00:11 UTC

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

# OpenClaw Project Digest — 2026-07-28

## 1. Today's Overview
OpenClaw saw very high activity over the past 24 hours, with 250 issues and 500 pull requests updated. Of the PRs, 227 were merged or closed, reflecting strong momentum in bug fixes and feature work. No new releases were published. Critical stability issues — including a P0 memory leak and several P1 regressions — remain under active investigation, while community feature requests like Linux/Windows desktop apps and memory trust tagging continue to gather support. The maintainer team is processing a large backlog of review- and decision‑pending items, many of which have been open for months.

## 2. Releases
None in the past 24 hours. The most recent tagged release appears to be `2026.7.2‑beta.4` (referenced in issue #113434). No release notes or migration guides are available for today.

## 3. Project Progress
**Merged/closed PRs today (selected from top comment‑count PRs):**

- **[PR #114803](https://github.com/openclaw/openclaw/pull/114803)** (closed) — `fix(cli): prevent update dry runs from mutating state`. Fixes an issue where `--dry-run` previews could inadvertently delete handoff directories or persist restart sentinels.
- **[PR #114801](https://github.com/openclaw/openclaw/pull/114801)** (closed) — `fix(heartbeat): prevent duplicated cron wakes and stale handler races`. Resolves concurrent cron execution and blocking.
- **[PR #114780](https://github.com/openclaw/openclaw/pull/114780)** (closed) — `fix(diagnostics): exclude active steering from session backlog`. Stops counting messages consumed by an active run as queued backlog.
- **[PR #94459](https://github.com/openclaw/openclaw/pull/94459)** (closed) — `fix: add OTel reply latency decomposition`. Adds bounded diagnostics spans and metrics for reply phases.

Other notable open PRs that advanced today include a refactor of agent run termination and fallback ([#114810](https://github.com/openclaw/openclaw/pull/114810)), a large DM‑policy contract suite cleanup ([#114776](https://github.com/openclaw/openclaw/pull/114776)), and ongoing work on Feishu channel attachment resolution ([#113987](https://github.com/openclaw/openclaw/pull/113987)).

## 4. Community Hot Topics
The most active discussions (by comment count and reactions) reveal strong demand for platform expansion and security features:

- **[Issue #75](https://github.com/openclaw/openclaw/issues/75)** (115 comments, 80 👍) — *Linux/Windows Clawdbot Apps*. The community is pushing hard for desktop agent clients on Linux and Windows to match the existing macOS, iOS, and Android offerings. This has been open since January and remains the highest‑reaction issue.
- **[Issue #91588](https://github.com/openclaw/openclaw/issues/91588)** (21 comments, 1 👍) — *Critical: Gateway Memory Leak — RSS grows from 350MB to 15.5GB over days*. A P0 stability issue causing OOM crashes; users are closely tracking progress.
- **[Issue #10659](https://github.com/openclaw/openclaw/issues/10659)** (15 comments, 4 👍) — *Feature Request: Masked Secrets*. A security enhancement to prevent agents from reading raw API keys.
- **[Issue #6615](https://github.com/openclaw/openclaw/issues/6615)** (10 comments, 8 👍) — *Feature: Add denylist support for exec‑approvals*. Users want “allow everything except X” policies.
- **[Issue #7722](https://github.com/openclaw/openclaw/issues/7722)** (10 comments, 4 👍) — *Filesystem Sandboxing Config*. A request for configurable file access restrictions.
- **[Issue #84569](https://github.com/openclaw/openclaw/issues/84569)** (11 comments, 3 👍) — *WhatsApp session stalls on long model_call* (closed). Users highlighted multi‑channel reliability as a key concern.

The underlying needs centre on broader platform support, better security controls, and improved stability for high‑volume channels.

## 5. Bugs & Stability
**Critical & High‑Severity Bugs (P0/P1):**

- **[Issue #91588](https://github.com/openclaw/openclaw/issues/91588)** (P0, open) — *Critical memory leak in gateway*. RSS grows to 15.5 GB, causing OOM kills and restart cycles. No fix PR yet; the issue has been open since June.
- **[Issue #109867](https://github.com/openclaw/openclaw/issues/109867)** (P0, closed) — *State migration creates index before adding column, blocking gateway startup*. Fixed in today’s PR activity (not explicitly linked).
- **[Issue #86519](https://github.com/openclaw/openclaw/issues/86519)** (P1, open) — *Agent repeats identical replies 2–10× on Telegram after 5.20 update*. Regression, partially improved in 5.22 but not fully resolved.
- **[Issue #113434](https://github.com/openclaw/openclaw/issues/113434)** (P1, open) — *Codex sessions.reset reuses retired session ID; catalog/file scans can exhaust Gateway RAM*. Confirmed on `2026.7.2‑beta.4`.
- **[Issue #113306](https://github.com/openclaw/openclaw/issues/113306)** (P1, open) — *SQLite snapshot restore lacks end‑to‑end crash and identity guarantees*.
- **[Issue #113323](https://github.com/openclaw/openclaw/issues/113323)** (P1, open) — *LLM idle timeout aborts agent runs during reasoning‑token streaming* on local reasoning models.
- **[Issue #114211](https://github.com/openclaw/openclaw/issues/114211)** (P1, open) — *Matrix room agents can loop on visible no‑reply output, restart recovery, and stale session replay*.

**Other notable bugs:**
- **[Issue #113754](https://github.com/openclaw/openclaw/issues/113754)** (P2, open) — Runtime context markers leak verbatim to delivery channels.
- **[Issue #102020](https://github.com/openclaw/openclaw/issues/102020)** (P1, closed) — “Second message in a session fails” – fixed.
- **[Issue #49603](https://github.com/openclaw/openclaw/issues/49603)** (P1, closed) — Orphaned lock files not cleared on gateway restart.

Most critical bugs have no linked fix PRs yet, but the high PR activity indicates that several fixes are in progress or awaiting review.

## 6. Feature Requests & Roadmap Signals
The following enhancements have strong community support and are likely candidates for upcoming releases:

| Issue | Title | Reactions | Status |
|-------|-------|-----------|--------|
| [#75](https://github.com/openclaw/openclaw/issues/75) | Linux/Windows Clawdbot Apps | 80 👍 | Open since Jan, high demand |
| [#6615](https://github.com/openclaw/openclaw/issues/6615) | Denylist support for exec‑approvals | 8 👍 | Open, linked PR open? (not in top PR list) |
| [#10659](https://github.com/openclaw/openclaw/issues/10659) | Masked Secrets | 4 👍 | Open, security‑focused |
| [#7707](https://github.com/openclaw/openclaw/issues/7707) | Memory Trust Tagging by Source | 0 👍 (but 22 comments) | Open, long‑standing |
| [#7722](https://github.com/openclaw/openclaw/issues/7722) | Filesystem Sandboxing Config | 4 👍 | Open |
| [#10687](https://github.com/openclaw/openclaw/issues/10687) | Dynamic model discovery (OpenRouter) | 3 👍 | Open |
| [#12219](https://github.com/openclaw/openclaw/issues/12219) | Skill Permission Manifest Standard | 0 👍 | Open, security |
| [#9986](https://github.com/openclaw/openclaw/issues/9986) | Trigger fallback on context length exceeded | 0 👍 | Open |

**Predictions for the next minor release:** The Masked Secrets feature (#10659) and exec‑approval denylist (#6615) align with ongoing security work and could ship in the next beta. The gateway memory leak (#91588) is the top priority and may see a hotfix release. Linux/Windows desktop apps (#75) are a larger undertaking but continue to gain traction.

## 7. User Feedback Summary
Based on today’s issues, users are experiencing several pain points:

- **Stability:** Memory leaks, OOM crashes, and duplicate message delivery on Telegram are causing significant frustration.
- **Missing features:** The lack of Linux/Windows desktop apps, filesystem sandboxing, and dynamic model discovery limits deployment scenarios.
- **Security:** Users are requesting more granular permission controls (masked secrets, denylist, skill manifests) to prevent accidental leaks and injection attacks.
- **Channel reliability:** WhatsApp stalls, Matrix looping, and Google Chat OAuth limitations are hampering multi‑channel workflows.
- **Configuration friction:** Session context bloat, hot‑reload model registry loss, and misleading error messages reduce trust in the system.

Overall satisfaction appears mixed: the community is engaged and vocal about improvements, reflecting both strong adoption and high expectations for reliability.

## 8. Backlog Watch
Several important issues have been open for months without a fix PR and remain tagged `clawsweeper:needs-maintainer-review` or `clawsweeper:needs-product-decision`:

- **[Issue #75](https://github.com/openclaw/openclaw/issues/75)** – Linux/Windows apps (since Jan 2026). Needs product decision.
- **[Issue #7707](https://github.com/openclaw/openclaw/issues/7707)** – Memory Trust Tagging (Feb 2026). Needs security review.
- **[Issue #10659](https://github.com/openclaw/openclaw/issues/10659)** – Masked Secrets (Feb 2026). Needs security review and product decision.
- **[Issue #10687](https://github.com/openclaw/openclaw/issues/10687)** – Dynamic model discovery (Feb 2026). Needs live repro.
- **[Issue #6615](https://github.com/openclaw/openclaw/issues/6615)** – Exec‑approval denylist (Feb 2026). Needs security review and product decision.
- **[Issue #7722](https://github.com/openclaw/openclaw/issues/7722)** – Filesystem sandboxing (Feb 2026). Needs product decision and live repro.
- **[Issue #8299](https://github.com/openclaw/openclaw/issues/8299)** – Suppress sub‑agent announce (Feb 2026). Needs product decision.
- **[Issue #67419](https://github.com/openclaw/openclaw/issues/67419)** – Session context bloat (Apr 2026). Needs product decision and live repro.
- **[Issue #113306](https://github.com/openclaw/openclaw/issues/113306)** – SQLite snapshot guarantees (Jul 24). Needs maintainer review.

These items, if unaddressed, risk accumulating technical debt and user dissatisfaction. The high volume of open PRs suggests maintainer capacity is stretched, but prioritising these long‑standing requests would improve community sentiment.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report — 2026-07-28

## 1. Ecosystem Overview

The personal AI assistant open-source ecosystem continues to mature rapidly, with **IronClaw's v1.0.0 ground-up rebuild** marking a significant architectural milestone this week. Across projects, the community is converging on **security hardening, cross-platform desktop support, and memory reliability** as top priorities, while individual projects differentiate through channel breadth, extensibility models, and target deployment scenarios. **OpenClaw maintains its position as the most active project by raw volume** (250 issues, 500 PRs daily), though its critical memory leak (P0) remains unresolved for weeks, creating an opening for competitors. The ecosystem is bifurcating between **comprehensive multi-channel platforms** (OpenClaw, ZeroClaw, Hermes) and **focused, lightweight alternatives** (NanoBot, Moltis), with IronClaw attempting to bridge both with a modernized architecture.

---

## 2. Activity Comparison

| Project | Issues Updated (24h) | Issues Closed | PRs Updated (24h) | PRs Merged/Closed | New Release | Health Score |
|---------|---------------------|---------------|-------------------|-------------------|-------------|--------------|
| **OpenClaw** | 250 | — | 500 | 227 merged | No | ⚠️ **Good** (very high activity, but P0 leak unaddressed, backlog growing) |
| **ZeroClaw** | 24 | 2 | 50 | 8 merged | No | ✅ **Good** (high activity, active bug-fixing, structured roadmap) |
| **IronClaw** | 19 | 3 | 50 | 19 merged | v1.0.0 (Jul 27) | ✅ **Excellent** (post-release stabilization, clear epics, responsive) |
| **CoPaw (QwenPaw)** | 19 | 5 | 49 | 15 merged | No | ✅ **Good** (high contributor engagement, active model-provider work) |
| **NanoBot** | 64 | 63 closed | 37 | 24 merged | No | ✅ **Excellent** (exceptional triage rate, responsive maintainers) |
| **Hermes Agent** | 12 | 2 | 50 | 10 merged | No | ✅ **Good** (rapid fix PRs for new bugs, NeMo metrics suite in progress) |
| **LobsterAI** | 7 | — | 9 | 5 merged | No | ⚠️ **Moderate** (critical `\f` corruption bug, 4-month stale backlog) |
| **PicoClaw** | 5 | 0 | 4 | 0 merged | No | ⚠️ **Low** (no merges today, MCP hang bug critical, low maintainer throughput) |
| **NanoClaw** | 0 | 0 | 8 | 0 merged | No | ✅ **Steady** (active PR refinement, no new user-reported bugs) |
| **Moltis** | 0 | 0 | 5 | 0 merged | No | ✅ **Steady** (feature PRs pending review, clean issue tracker) |
| **NullClaw** | 0 | 0 | 1 | 0 merged | No | ❌ **Stalling** (single Dependabot PR open 6 weeks, zero engagement) |
| **TinyClaw** | 0 | 0 | 0 | 0 | No | ❌ **Dormant** |
| **ZeptoClaw** | 0 | 0 | 0 | 0 | No | ❌ **Dormant** |

---

## 3. OpenClaw's Position

**Advantages:**
- **Raw community engagement is unmatched** — 80 👍 on the Linux/Windows desktop request (#75) demonstrates passionate demand
- **Channel breadth** supports Telegram, WhatsApp, Matrix, Feishu, Google Chat, WeCom — the widest multi-channel support in the ecosystem
- **Core reference implementation** status attracts constant PR velocity (500/day), ensuring bugs are found quickly
- **Open PR pipeline is healthy** — 227 merged/closed today indicates maintainers are working through backlog

**Disadvantages:**
- **P0 memory leak (#91588)** remains unpatched after 6+ weeks, causing OOM crashes and eroding trust
- **Maintainer bottleneck** — 250 issues daily with many `needs-maintainer-review` tags sitting for months (Jan 2026)
- **Desktop/mobile fragmentation** — macOS/iOS/Android apps exist but Linux/Windows are absent, creating a platform gap that Hermes Agent's native desktop app fills
- **Security lagging behind** — requests for masked secrets, denylist support, and filesystem sandboxing remain unimplemented while ZeroClaw and Moltis ship operator controls

**Technical Approach Differences:**
- OpenClaw uses a **Go-based gateway** with plugin skills model, while IronClaw rebuilt entirely in **Rust** for safety guarantees
- NanoBot and Moltis favor **Python/PyO3** with lightweight deployment; OpenClaw's architecture is heavier and more complex
- OpenClaw's **multi-channel-first** design contrasts with Hermes Agent's **desktop-first** and IronClaw's **sandbox-first** philosophy

**Community Size Comparison:**
- OpenClaw has the **largest absolute user base** (115 comments on Issue #75 alone), but ZeroClaw and IronClaw show **higher maintainer responsiveness per capita**
- OpenClaw's reaction counts (80👍, 22 comments) dwarf other projects' metrics, indicating a much larger installed base

---

## 4. Shared Technical Focus Areas

| Focus Area | Projects Affected | Specific Needs |
|------------|------------------|----------------|
| **Security & Access Control** | OpenClaw, PicoClaw, Moltis, ZeroClaw, LobsterAI | Masked secrets (#10659 OpenClaw), exec denylist (#6615 OpenClaw), operator whitelists (#1170 Moltis), sandbox credential firewalls (#6723 IronClaw), WhatsApp token leak (#9417 ZeroClaw), approval-level inheritance (#6506 CoPaw) |
| **Memory & Session Reliability** | OpenClaw, Hermes Agent, NanoBot, CoPaw, LobsterAI | Gateway memory leak (#91588 OpenClaw), session lock/file corruption (#113434 OpenClaw), Dream input integrity (#5114 NanoBot), `\f` byte corruption (#2393 LobsterAI), unbounded sub-sessions (#6505 CoPaw) |
| **Cross-Platform Desktop** | OpenClaw, Hermes Agent, PicoClaw, Moltis, LobsterAI | Linux/Windows desktop apps (#75 OpenClaw), macOS Tahoe plist (#42376 Hermes), spawn-helper permissions (#61396 Hermes), Windows PowerShell encoding (#2390 LobsterAI), PWA notifications (#1173 Moltis) |
| **Local Model / Provider Compatibility** | NanoBot, CoPaw, OpenClaw | Ollama/LM Studio errors (#2570 NanoBot), MiniMax-M3 truncation (#6324 CoPaw), OpenAI `max_tokens` not respected (#6258 CoPaw), DeepSeek serialization (#4548 IronClaw) |
| **Stability & Observability** | IronClaw, Hermes Agent, OpenClaw, ZeroClaw | Error-recoverability epic (#6284 IronClaw), NeMo Relay metrics (#67607 Hermes), OTel reply latency (#94459 OpenClaw), CI flakiness (#9357 ZeroClaw) |
| **Multi-Provider Flexibility** | NanoBot, CoPaw, Moltis | Custom model providers (#1991 NanoBot), dynamic model discovery (#10687 OpenClaw), custom MCP servers (#6727 IronClaw), Zvec memory backend (#1158 Moltis) |

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | IronClaw | Hermes Agent | NanoBot | CoPaw | ZeroClaw |
|-----------|----------|----------|--------------|---------|-------|----------|
| **Target User** | Power users, multi-channel teams | Enterprise, security-conscious | Desktop-first professionals | Lightweight, hobbyist, model-switchers | Chinese/Asian market, model-provider integrators | Operators, workflow automation |
| **Architecture** | Go gateway + plugin skills | Rust Reborn rewrite, sandbox-first | Electron desktop + gateway | Python/PyO3, WebUI-focused | Mission-mode, sub-agents, multi-model | Rust, SOP control plane, PostgreSQL |
| **Channel Support** | Broadest (8+ channels) | Moderate (Telegram, Discord, Signal) | Desktop-focused + gateway | WeChat/Feishu/Discord | WeChat/Feishu + phone | Discord, WhatsApp, Telegram |
| **Extensibility Model** | Skill plugins + OTel | Manifest V3 extensions + MCP | Skills + hooks + tools | Community skills.sh marketplace | Sub-agents + workspace checkpoints | WASM plugins + unified catalog |
| **State Management** | SQLite, gateway-based | Reborn process journal + migration | Desktop local + cloud | Git storage, memory consolidation | Shadow Git checkpoints | PostgreSQL, daemon-owned SOP |
| **Mobile Support** | iOS, Android | None | None | None | None | None |
| **Critical Risk** | P0 memory leak unpatched 6 weeks | Migration friction from legacy 0.29.x | macOS Tahoe compatibility | None immediate | Unbounded sub-session loops (critical) | SOP cancellation missing, auth migration broken |

**Key Differentiators:**

- **IronClaw** is the **most architecturally ambitious** — the Reborn rewrite, sandbox credential firewall, and error-recoverability epic set it apart for enterprise security use cases
- **Hermes Agent** owns the **desktop UX niche** — no other project has native macOS/Windows desktop apps with integrated terminal, pet overlay, and session sidebar
- **NanoBot** excels at **triaging user issues** — 63 issues closed in 24h is the best rate in the ecosystem, with a clean backlog
- **CoPaw** leads on **model-provider breadth** — integrating Volcano Engine, Xiaomi MiMo, Atlas Cloud alongside standard APIs, serving the Chinese AI market
- **ZeroClaw** is strongest on **infrastructure and testing** — PostgreSQL backend, CI hardening, and formal RFC governance

---

## 6. Community Momentum & Maturity

### Tier 1: Rapidly Iterating (High Activity + Clear Roadmap)
| Project | Momentum Signal | Maturity |
|---------|----------------|----------|
| **IronClaw** | v1.0.0 release, 10 new epics in 24h, 19 PRs merged | Post-major-release stabilization; architecture is modern but migration is painful |
| **ZeroClaw** | 50 PRs, structured RFC process, PostgreSQL backend merged | Maturing rapidly; v0.8.5 release tracker active; CI gaps being closed |
| **OpenClaw** | 500 PRs, 250 issues daily | **Highest raw activity** but bottlenecked maintainers; P0 bug erodes confidence |

### Tier 2: Stabilizing (High Activity, Fewer Critical Issues)
| Project | Momentum Signal | Maturity |
|---------|----------------|----------|
| **Hermes Agent** | 50 PRs, fix PRs opened same day as bug reports | Healthy fix pipeline; NeMo metrics suite indicates next-gen observability |
| **NanoBot** | 64 issues closed, 24 PRs merged | **Most responsive maintainers**; minimal backlog; WebUI enhancements |
| **CoPaw** | 49 PRs, 15 merged, strong contributor growth | New model-provider influx drives feature work; unbounded loop bug is concerning |

### Tier 3: Steady State (Moderate Activity)
| Project | Momentum Signal | Maturity |
|---------|----------------|----------|
| **NanoClaw** | 8 PRs, no new bugs | Feature additions ongoing; Signal channel fixes show attention to detail |
| **Moltis** | 5 PRs, clean tracker | Small but focused; security fix, Zvec backend, ACP agent role signal long-term vision |
| **PicoClaw** | 4 PRs, no merges | **Lowest throughput** among active projects; MCP hang and UI lag bugs unaddressed |

### Tier 4: Low Activity / Dormant
| Project | Status |
|---------|--------|
| **LobsterAI** | Critical `\f` corruption bug + 4-month stale PRs; engaged community but slow maintainers |
| **NullClaw** | Single Dependabot PR open 6 weeks; zero community engagement |
| **TinyClaw** | Dormant |
| **ZeptoClaw** | Dormant |

---

## 7. Trend Signals

1. **Security hardening is the dominant requirement across all active projects.** Masked secrets, operator whitelists, sandbox credentials, and exec approvals are being implemented in parallel — the ecosystem recognizes that open-source AI agents are high-value targets for credential theft and prompt injection.

2. **Memory and session reliability remain unsolved at scale.** Every major project reports memory leaks (OpenClaw P0, CoPaw sub-session explosions, LobsterAI `\f` corruption) or session-state corruption (ID reuse, snapshot guarantees). This is the #1 barrier to enterprise adoption — agents that lose state or run away are not production-ready.

3. **Desktop + mobile cross-platform support is the most requested feature gap** (OpenClaw #75 with 80👍, Hermes macOS/Windows issues). Developers want their agent to work on their development machine, not just in chat apps. The trend is toward **hybrid architectures** (local desktop agent + cloud gateway).

4. **Local-first AI is gaining traction.** Multiple projects are fixing local model compatibility (Ollama, LM Studio, MLX) and optimizing for CPU/on-device inference. IronClaw's sandbox-first design and NanoBot's Git-backed memory signal a move toward **offline-capable, privacy-preserving agents**.

5. **Extensibility is fragmenting.** IronClaw proposes Manifest V3 extensions, ZeroClaw uses WASM plugins, NanoBot has skills.sh marketplace, OpenClaw has skill plugins, and Moltis introduces ACP agent role. The lack of a **unified extension standard** creates ecosystem fragmentation. IronClaw's unified catalog RFC (#9346) may address this.

6. **Observability is becoming a competitive differentiator.** IronClaw's error-recoverability epic (#6284), Hermes Agent's NeMo Relay metrics suite, and OpenClaw's OTel reply latency decomposition all focus on making agent behavior transparent. Users are tired of silent failures and hallucinated actions.

7. **Multi-provider failover is table stakes.** Users across NanoBot, CoPaw, and OpenClaw report model-specific bugs (DeepSeek serialization, MiniMax-M3 truncation, OpenAI max_tokens). The community expects seamless fallback between providers — dynamic model discovery (#10687 OpenClaw) and configurable fallback chains (#3200 PicoClaw) are emerging solutions.

8. **The market is splitting into "kitchen sink" vs. "opinionated" platforms.** OpenClaw, ZeroClaw, and IronClaw aim to be comprehensive. NanoBot, Moltis, and PicoClaw are carving niches by staying focused (lightweight, security-first, language-specific). The opinionated projects are closing issues faster and have higher user satisfaction.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-07-28

## 1. Today’s Overview
The NanoBot project saw heavy activity in the last 24 hours, with **64 issues updated** (63 closed, 1 open) and **37 pull requests updated** (13 open, 24 merged/closed). The exceptionally high closure rate suggests a focused triage or bulk closure of older tickets, while the volume of merged PRs indicates active development across WebUI, memory/git storage, documentation, and channel support. No new releases were published, but the project appears to be in a rapid iteration phase ahead of a future version.

## 2. Releases
No new releases were published today. The latest available version remains the previous release (v0.1.4.post3, as referenced in issue #1672). Users should track the `main` branch for the latest fixes and features.

## 3. Project Progress
**24 pull requests were merged or closed today**, reflecting substantial progress. Notable changes include:

- **WebUI improvements**:  
  - [PR #5121](https://github.com/HKUDS/nanobot/pull/5121) – Fixed composer resize scroll jitter.  
  - [PR #5119](https://github.com/HKUDS/nanobot/pull/5119) – Softened model selector emphasis.  
  - [PR #5113](https://github.com/HKUDS/nanobot/pull/5113) – Stabilized repeated model preset rows.  
  - [PR #5077](https://github.com/HKUDS/nanobot/pull/5077) – Added ability to switch model presets directly from the composer.  
  - [PR #5076](https://github.com/HKUDS/nanobot/pull/5076) – Honored custom gateway port with Vite.  
  - [PR #5080](https://github.com/HKUDS/nanobot/pull/5080) – Migrated README and WebUI assets to SVG.

- **Memory & git storage fixes**:  
  - [PR #5124](https://github.com/HKUDS/nanobot/pull/5124) – Fixed double-encoding of git object IDs in `GitStore` (hex-of-hex bug).  
  - [PR #5114](https://github.com/HKUDS/nanobot/pull/5114) – Preserved Dream input integrity in memory consolidation.

- **Documentation**:  
  - [PR #5123](https://github.com/HKUDS/nanobot/pull/5123) – Improved README landing page with clearer use cases and contribution paths.

- **Other**:  
  - [PR #1683](https://github.com/HKUDS/nanobot/pull/1683) – Added `LLM_LOGGING` env var for request/response debug logging (merged after long review).

## 4. Community Hot Topics
The following issues and PRs attracted the most discussion (by comment count):

- **[Issue #1991](https://github.com/HKUDS/nanobot/issues/1991)** (9 comments, closed) – Request for support of multiple custom model providers. User Wcowin wanted the ability to add `custom2`, `custom3`, etc., and switch freely. This reflects a growing need for multi-model flexibility.

- **[Issue #3123](https://github.com/HKUDS/nanobot/issues/3123)** (8 comments, closed) – Problem with cron/scheduled tasks: messages sent via cron session cannot be later queried or corrected. Users want scheduled messages to be part of a regular session.

- **[Issue #2570](https://github.com/HKUDS/nanobot/issues/2570)** (7 comments, closed) – Local Ollama configuration difficulties (404, port not listening). A recurring pain point for users running local models.

- **[Issue #2329](https://github.com/HKUDS/nanobot/issues/2329)** (6 comments, closed) – Custom model provider works on CLI but fails on channels (e.g., Feishu). Highlights incompatibility between provider configuration and channel integration.

- **[Issue #2373](https://github.com/HKUDS/nanobot/issues/2373)** (5 comments, closed) – MiniMax API error with invalid function arguments. Users hitting provider-specific issues when using tool calls.

## 5. Bugs & Stability
Several bugs of varying severity were reported or fixed today:

| Severity | Bug | Status | Fix PR |
|----------|-----|--------|--------|
| **High** | [Issue #4792](https://github.com/HKUDS/nanobot/issues/4792): `/stop` silently discards pending queue messages – permanent message loss. | Closed | Not tracked; issue closed without linked PR? (likely addressed elsewhere) |
| **High** | [PR #5126](https://github.com/HKUDS/nanobot/pull/5126) (open) + [PR #5124](https://github.com/HKUDS/nanobot/pull/5124) (closed): Git object IDs doubly hex-encoded, corrupting memory versioning. | Fix merged | #5124 |
| **Medium** | [Issue #4805](https://github.com/HKUDS/nanobot/issues/4805): `suppress(Exception)` in `prepare_call` swallows tool validation errors. | Closed | Not noted |
| **Medium** | [Issue #1174](https://github.com/HKUDS/nanobot/issues/1174): Memory consolidation can take long or fail with local models. | Closed | Not noted |
| **Low** | [Issue #1487](https://github.com/HKUDS/nanobot/issues/1487): Qwen model returns JSON format error for function arguments. | Closed | Not noted |
| **Low** | [Issue #1033](https://github.com/HKUDS/nanobot/issues/1033): Inter-instance cache staleness for cron jobs. | Closed | Not noted |

A new open PR [ #5117](https://github.com/HKUDS/nanobot/pull/5117) fixes a bug where invalid idle-compaction timestamps could cause session listing failures.

## 6. Feature Requests & Roadmap Signals
User-requested features that gained traction include:

- **Multiple custom model providers** ([#1991](https://github.com/HKUDS/nanobot/issues/1991)) – Likely to be implemented in a future config redesign.
- **Optional tools and memory** ([#1881](https://github.com/HKUDS/nanobot/issues/1881)) – A toggle to disable memory updates and tools for weaker models, plus plugin support similar to openclaw.
- **Customize or disable 🐈 emoji** ([#2747](https://github.com/HKUDS/nanobot/issues/2747)) – Simple quality-of-life config request.
- **WebSocket for proactive message delivery** ([#3559](https://github.com/HKUDS/nanobot/issues/3559)) – Important for multi-tenant setups; maintainer noted official support.
- **Message push across channels** ([#3074](https://github.com/HKUDS/nanobot/issues/3074)) – Ability to send messages from API to another channel.

New feature PRs today signal the next version’s direction:

- **[PR #5112](https://github.com/HKUDS/nanobot/pull/5112)** – Expose Dream runs as read-only sessions in WebUI (in-progress).
- **[PR #5116](https://github.com/HKUDS/nanobot/pull/5116)** – Add skills.sh marketplace and skill management in WebUI.
- **[PR #5098](https://github.com/HKUDS/nanobot/pull/5098)** – Introduce a unified extension platform for code-level capabilities.
- **[PR #5115](https://github.com/HKUDS/nanobot/pull/5115)** – Add LINE Messaging API channel (popular in Asia).
- **[PR #5110](https://github.com/HKUDS/nanobot/pull/5110)** – Make `nanobot status` actionable for agent readiness checks.

These strongly suggest the upcoming release will emphasize **WebUI enhancements, skill ecosystem (marketplace), extension system, and new channel support**.

## 7. User Feedback Summary
Real user pain points observed from today’s data:

- **Configuration complexity**: Users struggle with custom model provider setup, especially with local models (Ollama, LM Studio) and API key requirements (#2570, #1947, #1478, #1590).
- **Channel inconsistencies**: Custom providers work on CLI but break on channels like Feishu (#2329); cron messages cannot be followed up (#3123); Discord slash commands conflict (#1315).
- **Memory and session reliability**: Memory consolidation fails with local models (#1174); session switch does not stop old cron jobs (#2358); cache staleness across instances (#1033).
- **Feature gaps**: No native support for multiple custom models (#1991), no progress notifications on Feishu (#3166), no way to disable memory/tools (#1881).

Satisfaction signals: the rapid closure of issues (63 in one day) and merging of multiple high-value PRs indicate an active, responsive maintainer team.

## 8. Backlog Watch
Despite the high closure rate, a few items remain open and warrant attention:

- **[PR #4667](https://github.com/HKUDS/nanobot/pull/4667)** (opened July 2, 2026) – Fix to protect user skills from Dream writes. Still open after 26 days; requires review and merge to prevent potential data loss.
- **[Issue #3559](https://github.com/HKUDS/nanobot/issues/3559)** (closed, but underlying feature is not yet implemented) – WebSocket channel for proactive delivery. While closed, the feature is still pending in code; maintainer may need to track implementation.
- **[Issue #1033](https://github.com/HKUDS/nanobot/issues/1033)** (closed) – Inter-instance cache staleness remains unfixed? The issue was closed, but the root cause (no re-read from disk) may still affect users running multiple instances.

The single open issue (not shown in top 30) is a positive sign – the project is well-maintained with minimal backlog. However, contributors and maintainers should prioritize the Dream write guard fix (#4667) as it directly affects user skills integrity.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-07-28

## 1. Today's Overview

The project shows **high activity** with 50 pull requests updated in the last 24 hours (10 merged/closed) and 12 issues updated (2 closed). No new releases were published. The PR surge is driven largely by a multi-part **NeMo Relay observability stack** (7 stacked drafts) and several user-facing desktop fixes. Bug reports are concentrated around session state, interrupt reliability, and macOS/Windows platform quirks. The maintainers appear responsive: two P1 bugs attracted companion fix PRs within hours of reporting.

## 2. Releases

**No new releases** in the last 24 hours. The latest published version remains v0.16.0 (2026.6.5) as referenced in issue #42376.

## 3. Project Progress

- **Merged/closed PRs** visible in the top tier:
  - [`#72985`](https://github.com/NousResearch/hermes-agent/pull/72985) — fix(desktop): delay Tip hover-open by 200ms (UI polish)
  - [`#72929`](https://github.com/NousResearch/hermes-agent/pull/72929) — fix(agent): never replay chain-of-thought in the active-turn redirect checkpoint (critical Anthropic session brick fix)
- **Closed issues** that signal progress:
  - [`#72016`](https://github.com/NousResearch/hermes-agent/issues/72016) — Gateway sessions lack activity watchdog (P1, closed after fix)
  - [`#72970`](https://github.com/NousResearch/hermes-agent/issues/72970) — Windows startup slow: skill provenance backfill (closed as duplicate, likely patched upstream)

The **NeMo Relay observability** PR series ([#67607](https://github.com/NousResearch/hermes-agent/pull/67607), [#69437](https://github.com/NousResearch/hermes-agent/pull/69437), [#69416](https://github.com/NousResearch/hermes-agent/pull/69416), [#68978](https://github.com/NousResearch/hermes-agent/pull/68978), [#68883](https://github.com/NousResearch/hermes-agent/pull/68883), [#68882](https://github.com/NousResearch/hermes-agent/pull/68882), [#68881](https://github.com/NousResearch/hermes-agent/pull/68881)) remains open but is actively iterating; it will deliver shared-metrics, skill and tool lifecycle tracking, and client-resource reporting.

## 4. Community Hot Topics

- **`#67600`** [OPEN] [*Bug*: Desktop session sidebar empty for `default` profile only](https://github.com/NousResearch/hermes-agent/issues/67600) — 13 comments. Users report that the sidebar goes blank for the **default profile** while named profiles work fine and the backend confirms data is served. Discussion suggests a frontend state or query filter bug; still unmerged.

- **`#61396`** [OPEN] [*Bug*: fix(desktop): preserve node-pty spawn-helper execute bit on macOS](https://github.com/NousResearch/hermes-agent/issues/61396) — 4 comments. Integrated terminal fails on macOS arm64 due to missing execute permission on `node-pty`’s `spawn-helper`. A core terminal reliability concern.

- **`#42376`** [OPEN] [*Bug*: `hermes gateway restart` generates broken plist on macOS Tahoe](https://github.com/NousResearch/hermes-agent/issues/42376) — 3 comments. `LimitLoadToSessionType` prevents launchctl bootstrap on macOS 26.5.1. Affects all users on the latest macOS version.

- **`#72971`** [OPEN] [*Bug*: Desktop GUI sends `prompt.submit` to wrong session after switching while model is slow](https://github.com/NousResearch/hermes-agent/issues/72971) — 2 comments. Multi-session users experience message misrouting.

**Underlying needs:** Clear appetite for robust multi-profile/multi-session experiences, reliable terminal integration on macOS, and macOS Tahoe compatibility. The session sidebar bug (#67600) is especially painful because it disables the primary UI for the default profile.

## 5. Bugs & Stability

| Bug | Severity | Summary | Fix PR |
|-----|----------|---------|--------|
| [`#72975`](https://github.com/NousResearch/hermes-agent/issues/72975) | **P1** | Interrupt/abort silently no-ops when `force_close_tcp_sockets()` finds 0 sockets, leaving request alive for minutes | [`#72982`](https://github.com/NousResearch/hermes-agent/pull/72982) (fixes socket shutdown on httpx mount pools) |
| [`#72989`](https://github.com/NousResearch/hermes-agent/issues/72989) | **P2** | Quoted/ambiguous text treated as permanent model ban, persists in memory | [`#72988`](https://github.com/NousResearch/hermes-agent/pull/72988) (stages unconfirmed policy claims instead of persisting) |
| [`#72971`](https://github.com/NousResearch/hermes-agent/issues/72971) | **P2** | Prompt submitted to wrong session after session switch | No fix PR yet |
| [`#72969`](https://github.com/NousResearch/hermes-agent/issues/72969) | **P2** | Windows cua-driver version mismatch (installer 0.12.6 vs runtime 0.8.3) | No fix PR yet |
| [`#72981`](https://github.com/NousResearch/hermes-agent/issues/72981) | **P3** | Honcho dependency install fails with permission denied on Managed Cloud v0.19.0 | No fix PR yet |
| [`#64115`](https://github.com/NousResearch/hermes-agent/issues/64115) | **P2** | Windows cua-driver returns 0x0 screenshot (Chinese username path) | No fix PR yet |
| [`#64681`](https://github.com/NousResearch/hermes-agent/issues/64681) | **P2** | Same root cause as #72989 (duplicate) | Covered by #72988 |

The **P1 interrupt hang** (#72975) is the most critical stability issue reported today; a fix PR (#72982) was opened simultaneously. The **model ban bug** (#72989) is equally concerning from a data-integrity perspective and already has a companion fix.

## 6. Feature Requests & Roadmap Signals

- **NeMo Relay observability suite** (PRs #67607, #69437, #69416, #68978, #68883, #68882, #68881) — a large coordinated effort to add runtime, client, skill, tool, and model metrics. Likely targeted for the next minor release (v0.17.0).
- **Desktop pet overlay enhancements** ([`#72986`](https://github.com/NousResearch/hermes-agent/pull/72986)) — drag animation and floor roaming. A playful UX addition.
- **Post-review hook** ([`#72984`](https://github.com/NousResearch/hermes-agent/pull/72984)) — allow custom pipelines after each agent turn. Signals community demand for extensibility.
- **Git status visibility** ([`#72793`](https://github.com/NousResearch/hermes-agent/pull/72793)) — add config to show/hide composer Git status.
- **Right-click menus mirrored to kebabs** ([`#72987`](https://github.com/NousResearch/hermes-agent/pull/72987)) — desktop UX parity.
- **Tool call grouping in Desktop** ([`#72893`](https://github.com/NousResearch/hermes-agent/pull/72893)) — reduce visual noise from many tool executions.

**Prediction**: The NeMo Relay metrics stack will be the headline feature of the next release, alongside the Anthropic chain-of-thought fix and the interrupt abort fix.

## 7. User Feedback Summary

**Pain points expressed:**
- *"Desktop session sidebar is completely empty for the default profile"* — #67600: a showstopper for users relying on the default profile.
- *"Interrupt logs success while the request stays alive for minutes"* — #72975: leads to frustration and work lost.
- *"Messages appearing in the wrong conversation"* — #72971: serious UX regression for multi-session users.
- *"Quoted text permanently banned a model"* — #72989/#64681: concern about irreversible agent behavior changes.
- *"Windows computer_use screenshot returns 0x0"* — #64115: blocks desktop automation on Windows with non-ASCII usernames.
- *"Terminal failed to start: posix_spawnp failed"* — #61396: macOS terminal users left without a workaround.
- *"Gateway restart breaks launchctl bootstrap"* — #42376: macOS Tahoe users cannot restart their gateway.

**Use cases:**
- Multi-profile power users who toggle between "default", "local", "minim", etc.
- macOS users on the latest Tahoe release.
- Windows users relying on computer_use for automation (WeChat, etc.).
- Developers using integrated terminal and session history.

**Satisfaction:** The rapid availability of fix PRs (#72982, #72988) for newly reported P1/P2 bugs suggests a healthy fix pipeline, but the sheer number of open bugs indicates overall stability is still maturing.

## 8. Backlog Watch

| Issue | Age | Notes |
|-------|-----|-------|
| [`#24790`](https://github.com/NousResearch/hermes-agent/issues/24790) | 2026-05-13 (2.5 months) | WeCom heartbeat reconnect + retry — stale, no recent maintainer comment. P2, but affects WeCom users. |
| [`#42376`](https://github.com/NousResearch/hermes-agent/issues/42376) | 2026-06-08 (1.5 months) | macOS Tahoe plist issue — still open, no fix PR. Updated today but no resolution. |
| [`#61396`](https://github.com/NousResearch/hermes-agent/issues/61396) | 2026-07-09 (19 days) | macOS terminal spawn-helper — no fix PR. Affects all macOS arm64 users. |
| [`#64115`](https://github.com/NousResearch/hermes-agent/issues/64115) | 2026-07-14 (14 days) | Windows cua-driver 0x0 screenshot — no fix PR. Chinese username path issue unreproduced by maintainers. |
| [`#64681`](https://github.com/NousResearch/hermes-agent/issues/64681) | 2026-07-15 (13 days) | Model ban from ambiguous text — duplicate of #72989 which now has a fix PR, but #64681 lacks maintainer triage. |
| [`#67600`](https://github.com/NousResearch/hermes-agent/issues/67600) | 2026-07-19 (9 days) | Empty desktop sidebar for default profile — high comment activity but no clear path to fix yet. |

**Maintainers should prioritise** #24790 (longest unresolved) and #42376 (macOS Tahoe compatibility) as they block core functionality for specific platforms. #67600 remains the most active community pain point without a solution in sight.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest — 2026-07-28

## 1. Today's Overview
Activity on the PicoClaw repository has been moderate over the past 24 hours, with **5 open issues** and **4 open pull requests** receiving updates, but no merges or new releases. The project continues to see community contributions in the form of localisation and provider updates, while several bugs and feature requests remain open, indicating a healthy mix of development and user feedback. One notable concern is the lack of merged PRs today, which may point to a bottleneck in review/merge throughput.

## 2. Releases
None. No new releases were published.

## 3. Project Progress
**No pull requests were merged or closed today.** All four PRs updated in the last 24 hours remain open:

- [#3273 [OPEN] feat(webui): add Japanese (ja) localization](https://github.com/sipeed/picoclaw/pull/3273) – submitted by honbou, implements the feature requested in issue #3272.
- [#3271 [OPEN] chore(providers): update default model names to 2026-07 latest](https://github.com/sipeed/picoclaw/pull/3271) – refreshes model lists for nine AI providers.
- [#3270 [OPEN] feat: add DashScope TTS provider and WeChat audio file sending](https://github.com/sipeed/picoclaw/pull/3270) – adds Alibaba Cloud TTS and WeChat integration.
- [#3200 [OPEN] feat(models): add configurable default fallback chain](https://github.com/sipeed/picoclaw/pull/3200) – a larger feature from July 1 that has not seen recent activity, but was updated (likely rebased or re-reviewed).

No completed work to report for the day, but the presence of these open PRs shows ongoing development efforts.

## 4. Community Hot Topics
All updated issues and PRs have exactly **1 comment each** with **no 👍 reactions**, so no single discussion dominates. However, the following items address recurring user needs:

- **Issue #3276** – [Launcher: support/detect an externally-managed gateway (systemd)](https://github.com/sipeed/picoclaw/issues/3276). Users deploying PicoClaw as a headless systemd service report friction with the launcher assuming control of the gateway lifecycle. This reflects a broader pain point for server‑side deployments.
- **Issue #3272** – [Feature: Add Japanese localization to PicoClaw WebUI and Launcher](https://github.com/sipeed/picoclaw/issues/3272). Has an accompanying open PR (#3273) that provides the full translation, suggesting this will be resolved soon.
- **PR #3200** – [feat(models): add configurable default fallback chain](https://github.com/sipeed/picoclaw/pull/3200). This feature has been open for nearly a month and, although updated today, still lacks a merge. It would allow users to define model fallback chains via the WebUI – a frequently requested capability for multi‑provider reliability.

## 5. Bugs & Stability
Three bugs were reported/updated in the last 24 hours, all open and without visible fix PRs:

1. **Critical** – [#3269 [BUG] If the MCP server connection fails, the agent loop will hang](https://github.com/sipeed/picoclaw/issues/3269). The agent loop stops responding to user chat when an MCP server is unreachable. This is a severe stability issue that can break the core chat experience. No corresponding fix PR exists.
2. **High** – [#3281 [BUG] Web UI chat input is very laggy when history has a little bit long](https://github.com/sipeed/picoclaw/issues/3281). Poor UI performance with long session history degrades daily usability. Likely related to frontend rendering of large chat contexts.
3. **Medium** – [#3268 [Bug]: exec tool action parameter should default to "run" instead of being required](https://github.com/sipeed/picoclaw/issues/3268). While not a crash bug, it causes unpredictable failures when LLM tool calls omit the `action` field. A default would improve reliability.

**No fixes for these bugs have been proposed in the open PRs.** Maintainers should prioritize the MCP hang issue (#3269) as it directly blocks usage.

## 6. Feature Requests & Roadmap Signals
Several feature requests were active:

- **Japanese localization** (#3272) – an open PR (#3273) already implements it; this is likely to be merged soon.
- **Externally‑managed gateway / systemd support** (#3276) – addresses a clear deployment use case for server environments. May be considered for the next minor version.
- **DashScope TTS and WeChat audio sending** (#3270) – an open PR adds this functionality, indicating community demand for voice and messaging integrations.
- **Configurable default fallback chain** (#3200) – a mature feature that has been languishing; if merged, it would be a significant UX improvement for multi‑model setups.

**Prediction for next version:** Japanese localisation and the TTS/WeChat features (both with PRs ready) are the most likely inclusions. The fallback chain PR (#3200) may also land if it receives maintainer review.

## 7. User Feedback Summary
Real pain points voiced by the community:

- **Headless deployment friction** – PicoClaw launcher assumes full lifecycle control, making systemd integration clumsy.
- **Missing localisation** – Japanese‑speaking users want UI translation, already addressed by a community PR.
- **Tool defaults** – The `exec` tool’s required `action` parameter forces LLMs to guess the correct value.
- **MCP reliability** – A failed MCP server hangs the entire agent loop; users need graceful degradation.
- **UI performance** – Long chat histories cause input lag, hampering productivity.

Overall sentiment appears constructive: users are actively contributing PRs and issue reports, but there is some frustration around unaddressed bugs that break core functionality.

## 8. Backlog Watch
Several items require maintainer attention:

- **Issue #3269** – MCP hang (critical, open 8 days, no fix in sight).
- **Issue #3281** – Web UI lag (high severity, open 7 days, no assignee).
- **PR #3200** – Fallback chain feature (open since July 1, updated today but no merge). This is a well‑scoped feature that would benefit a wide user base; its long review cycle raises questions about prioritisation.
- **Issue #3268** – `exec` tool default (open 9 days, minor but easy to fix).

All items are labelled `stale` except #3281, which may indicate that maintainers have not yet triaged them. Prompt engagement on #3269 and #3281 would improve project health.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-07-28

## 1. Today's Overview

Project activity remains steady with zero new issues or releases, but eight open pull requests received updates in the last 24 hours, indicating an active development cadence. No merged or closed PRs were recorded today, suggesting a focus on refining in-flight work. The PRs span bug fixes, new skills (utility and channel integrations), documentation improvements, and core wiring changes. The lack of new issues may reflect a mature codebase or a lull in user-reported problems, but the breadth of ongoing PRs shows the team is addressing both stability and feature gaps.

## 2. Releases

No new releases were published today. The latest release date is not available from the provided data; users should check the [releases page](https://github.com/nanocoai/nanoclaw/releases) for the most recent version.

## 3. Project Progress

No pull requests were merged or closed today. Several open PRs advanced their state, the most notable being:

- **Fix: Preserve resolved approval card content** – PR [#3143](https://github.com/nanocoai/nanoclaw/pull/3143) (Koshkoshinsk) retains title and request details after resolution, improving UX for terminal cards.  
- **Fix: Forward image/file attachments via mounted inbox** – PR [#3142](https://github.com/nanocoai/nanoclaw/pull/3142) (ira-at-work) corrects a dead path bug that prevented agents from reading non-image attachments in the Signal channel.  
- **Fix: Respect container.json skill selection for CLAUDE.md fragments** – PR [#3141](https://github.com/nanocoai/nanoclaw/pull/3141) (ERMOKHINNA) ensures skill selection is honored during container composition.  
- **Feature: Add Dial to channel picker + wizard/skills** – PR [#3050](https://github.com/nanocoai/nanoclaw/pull/3050) (OmriBenShoham) introduces a new communication channel.  
- **Feature: Add ncc utility skill for host operational and health CLI** – PR [#2971](https://github.com/nanocoai/nanoclaw/pull/2971) (zivisaiah) provides a standalone tool for server health checks.  
- **Core fix: Engagement consistency and expose self-serve wiring controls** – PR [#3137](https://github.com/nanocoai/nanoclaw/pull/3137) (Koshkoshinsk) allows group-scoped agents to inspect and request policy updates, and prevents invalid engagement regexes.  

## 4. Community Hot Topics

No issues received comments or reactions in the last 24 hours. Among pull requests, the most active by update frequency and scope are:

- **PR [#3137](https://github.com/nanocoai/nanoclaw/pull/3137)** – *Fix engagement consistency and expose self-serve wiring controls*: This core-team PR touches on agent autonomy and policy management. The underlying need is for multi-agent environments to allow self-service adjustments without maintainer intervention.  
- **PR [#2685](https://github.com/nanocoai/nanoclaw/pull/2685)** – *docs(signal): group typing, outbound reactions, quote-reply fix* (ira-at-work): Although opened a month ago, it was recently updated. It addresses documentation gaps for Signal group features, indicating a user demand for clearer integration docs.  
- **PR [#3050](https://github.com/nanocoai/nanoclaw/pull/3050)** – *Add Dial to channel picker*: A new channel integration suggests community interest in expanding supported messaging platforms beyond Signal.  

## 5. Bugs & Stability

Three fix-oriented PRs were updated today, all ranked as medium to high severity:

1. **Critical** – **PR [#3142](https://github.com/nanocoai/nanoclaw/pull/3142)** (`fix(signal): forward image/file attachments through the mounted inbox`): Attachments (PDFs, documents) were unreachable because a dead path was spliced into message text. This rendered the Read tool unable to open any non-image attachment. The fix mounts the correct path. No workaround exists without this change.  
2. **High** – **PR [#3143](https://github.com/nanocoai/nanoclaw/pull/3143)** (`fix: preserve resolved approval card content`): Resolved approval cards lost their details after button removal. The fix persists the original body and shows a muted status. This improves reliability for approval workflows.  
3. **Medium** – **PR [#3141](https://github.com/nanocoai/nanoclaw/pull/3141)** (`fix(compose): respect container.json skill selection for CLAUDE.md fragments`): When skills were selected via `container.json`, the container composition sometimes ignored the selection, leading to missing fragments. The fix aligns composition logic with the manifest.  

No new bugs were reported via Issues today. An older open fix (PR [#2346](https://github.com/nanocoai/nanoclaw/pull/2346)) addresses unknown slash commands being misinterpreted as Claude Code commands, causing silent response drops – a subtle but significant usability bug.

## 6. Feature Requests & Roadmap Signals

Based on recent PRs, the following features are likely to land in the next release:

- **Dial channel integration** (PR [#3050](https://github.com/nanocoai/nanoclaw/pull/3050)) – Adds a new messaging channel, expanding agent reach.  
- **NCC utility skill** (PR [#2971](https://github.com/nanocoai/nanoclaw/pull/2971)) – A CLI tool for host operational and health queries, enabling agents to perform system diagnostics.  
- **Engagement self-serve controls** (PR [#3137](https://github.com/nanocoai/nanoclaw/pull/3137)) – Allows group agents to request policy updates, moving toward more autonomous multi-agent setups.  

No user-requested features were submitted as Issues today, but the presence of a new channel skill suggests the community values extensibility.

## 7. User Feedback Summary

No direct user comments were captured in the last 24 hours. However, the fixes in progress reveal real pain points:

- **Attachment handling in Signal** (PR [#3142]) – Users could not view non-image attachments (PDFs, text files, documents) because the path was unmounted. The fix directly addresses a broken user workflow.  
- **Approval card resolution** (PR [#3143]) – Users likely lost context after approving or rejecting cards; the fix ensures information persists.  
- **Slash command ambiguity** (PR [#2346]) – Users typing unknown commands like `/joke` would get silent failures. The fix will treat them as normal chat, improving overall reliability.  

These indicate a user base that relies heavily on Signal for file sharing and on approval workflows, and that expects the tool to handle edge cases gracefully.

## 8. Backlog Watch

Two open pull requests have been waiting for an extended period and may require maintainer attention:

- **PR [#2685](https://github.com/nanocoai/nanoclaw/pull/2685)** – *docs(signal): group typing, outbound reactions, quote-reply fix* (opened 2026-06-04, last updated 2026-07-27). Despite its age, it received updates recently. The PR adds documentation for Signal group features and should be reviewed to align with the codebase.  
- **PR [#2346](https://github.com/nanocoai/nanoclaw/pull/2346)** – *fix(formatter): treat unknown slash commands as normal chat* (opened 2026-05-08, last updated 2026-07-27). This fix addresses a silent failure mode that could affect any user typing random slash commands. Low complexity but high impact; prioritising it would improve stability.  

No new Issues are languishing. The open PRs are being updated, so the backlog is actively managed, but merging these older items would reduce technical debt.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-07-28

## 1. Today's Overview
NullClaw saw no new issues, releases, or merged pull requests in the last 24 hours. The only activity was an update to an existing open pull request (#956) that bumps the Alpine base image from 3.23 to 3.24. This dependency update, submitted by Dependabot, has remained open for over six weeks with no maintainer feedback. Overall project activity remains low, indicating a stable but quiet phase with minimal community engagement.

## 2. Releases
No new releases were published in the reporting period.

## 3. Project Progress
No pull requests were merged or closed today. No new features or fixes were advanced.

## 4. Community Hot Topics
The sole open pull request is [#956](https://github.com/nullclaw/nullclaw/pull/956):  
- **Title:** `ci(deps): bump alpine from 3.23 to 3.24 in the docker-images group`  
- **Author:** dependabot[bot]  
- **Created:** 2026-06-15 | **Updated:** 2026-07-27  
- **Comments:** 0 | **Reactions:** 0  

This PR has generated no discussion or reactions, suggesting limited maintainer priority or user awareness. The underlying need is routine dependency hygiene — keeping the Docker base image up to date for security and compatibility.

## 5. Bugs & Stability
No bugs, crashes, or regressions were reported in the last 24 hours. No open issues tagged as bugs exist.

## 6. Feature Requests & Roadmap Signals
No feature requests were submitted or updated today. The data provides no signals for upcoming roadmap items.

## 7. User Feedback Summary
No user feedback (comments, reactions, or support issues) was recorded in the past 24 hours. Without active issues or PR discussions, it is not possible to assess user satisfaction or pain points.

## 8. Backlog Watch
The only item requiring attention is [PR #956](https://github.com/nullclaw/nullclaw/pull/956), which has been open since June 15, 2026 — more than 40 days. While it is a low-risk dependency update from Dependabot, the lack of maintainer review or merge leaves the project running on Alpine 3.23. A decision (merge, close, or request changes) would help keep the CI and Docker images current. No other long-unanswered issues or PRs are present.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

**Date:** 2026-07-28  
**Analysis Window:** 2026-07-27 – 2026-07-28 (24 hours)

---

## 1. Today's Overview

The IronClaw project is in a high-activity phase following the **v1.0.0 stable release** on 2026-07-27. In the last 24 hours, 50 pull requests and 19 issues were updated, with 19 PRs merged/closed and 3 issues closed. The primary thrust is **stabilization, migration tooling, and extending the new "Reborn" architecture** — 10 new epic issues were created covering extension platform unification, memory providers, messaging channels, and a migration path from legacy. The CI/e2e testing infrastructure is being hardened (epic #6524, PRs #6738, #6728), and several critical bugs (DeepSeek serialization, delivery target leaks, systemd onboarding) were closed. Overall project health is robust with a clear roadmap for Q3 2026.

---

## 2. Releases

**`ironclaw-v1.0.0`** – 2026-07-27  
[Release](https://github.com/nearai/ironclaw/releases/tag/ironclaw-v1.0.0) | [Issue #6725](https://github.com/nearai/ironclaw/issues/6725) (migration tracking)

- **This is a ground-up rebuild** of the agent runtime, storage, extension host, and Web UI — not an increment on the 0.29.x line.  
- **Breaking change:** The main binary is now `ironclaw`. The legacy 0.29.x monolith is available as `ironclaw-legacy`.  
- **Migration path:** A dedicated epic (#6725) has been opened to track the transition from pre-Reborn (legacy `src/`) to v1. All existing users must re-onboard and reconfigure extensions; no automated upgrade path exists yet.  
- **Documentation restructured:** PR #6692 ([merged](https://github.com/nearai/ironclaw/pull/6692)) fixed 33 internal doc paths that were accidentally served publicly on the live docs site.

---

## 3. Project Progress (Merged/Closed PRs Today)

19 PRs were merged or closed. Major items:

| PR | Summary | Impact |
|----|---------|--------|
| [#6684](https://github.com/nearai/ironclaw/pull/6684) | Collapse five failure-kind enums into one `FailureKind` (36 variants) + fix 6 wrongful-terminal/retry bugs | Core stability; part of error‑recoverability epic #6284 |
| [#6723](https://github.com/nearai/ironclaw/pull/6723) | Add unwired sandbox credential firewall primitives (CA + obligation staging) | Foundation for persistent per‑user sandbox containers |
| [#6692](https://github.com/nearai/ironclaw/pull/6692) | Restructure docs site around shipped 1.0 binary; fix internal doc exposure | Documentation quality & security |
| [#6687](https://github.com/nearai/ironclaw/pull/6687) | Dependencies: bump 33 packages across the workspace | Dependency hygiene |
| [#3847](https://github.com/nearai/ironclaw/pull/3847) | Add filesystem-backed Reborn skill bundle source | Enables local skill development (from May, merged today) |

Also closed:  
- [#4548](https://github.com/nearai/ironclaw/issues/4548) Bug: DeepSeek duplicate `model` field (fix merged earlier, confirmed closed)  
- [#6060](https://github.com/nearai/ironclaw/issues/6060) Bug: routine delivery target leak (fixed)  
- [#6575](https://github.com/nearai/ironclaw/issues/6575) Bug: systemd error after `ironclaw onboard` (fixed)

**Open high‑impact PRs progressing:**  
- [#6696](https://github.com/nearai/ironclaw/pull/6696) – Collapse lifecycle state into process journal (DB migration, XL)  
- [#6724](https://github.com/nearai/ironclaw/pull/6724) – Rebuild memory provider contract around declared capabilities  
- [#6740](https://github.com/nearai/ironclaw/pull/6740) – TLS termination seam for sandbox egress proxy  
- [#6737](https://github.com/nearai/ironclaw/pull/6737) – Fix extension behaviors silently reverted in PR #6616  

---

## 4. Community Hot Topics

Most active discussions (by comment count on issues):

- **[#6284 – Error‑Recoverability Endgame](https://github.com/nearai/ironclaw/issues/6284)** – 14 comments  
  Epic demanding that every mid‑run error be survivable, visible to the model, and actionable. This is the top‑priority user‑facing reliability initiative. Five sub‑PRs are already in flight (e.g., #6684, #6697).

- **[#6524 – Hermetic Testing Platform](https://github.com/nearai/ironclaw/issues/6524)** – 3 comments  
  Epic to guarantee deterministic, meaningful coverage for every capability and journey. Workstreams are actively producing PRs for isolation proofs (#6738, #6728) and fault containment.

- **[#6481 – Unified Manifest‑Driven Extension Platform](https://github.com/nearai/ironclaw/issues/6481)** – New today, but already spawning implementation PRs (#6655, #6729). High interest from extension developers.

**Underlying needs:**  
The community (both contributors and users) is demanding **reliability, testability, and extensibility** as the project transitions to v1. The high volume of epic issues suggests a structured approach to hardening before further feature growth.

---

## 5. Bugs & Stability

| Bug | Severity | Status | Notes |
|-----|----------|--------|-------|
| DeepSeek serialization duplicate `model` field (#4548) | **High** (API 400) | Closed – fix merged | Reported by community member, fixed and verified |
| Routine delivery target leaks across routines (#6060) | **High** (data leak/confusion) | Closed – fix merged | Per‑routine delivery now properly isolated |
| `systemd` service error after `ironclaw onboard` (#6575) | **Medium** (onboarding broken) | Closed – fix merged | Ubuntu users affected; root cause addressed |
| `register_generic_channel_outbound_targets` is no‑op (#6726) | **Low** (dead code) | Open – suggests testing gap | No functional impact but indicates coverage blind spot |
| Memory provider lifecycle tools incorrectly registered (#6730) | **Medium** (agent‑visible tools) | Open – fix PR #6724 in progress | Part of pluggable memory epic #6482 |

**No critical unhandled regressions today.** The three high‑severity bugs from previous days were all closed, reflecting strong momentum on stability.

---

## 6. Feature Requests & Roadmap Signals

New epics filed today (all created 2026-07-27):

| Issue | Feature | Likely Version |
|-------|---------|----------------|
| [#6725](https://github.com/nearai/ironclaw/issues/6725) | Migration path: legacy → v1 | **v1.0.x** (urgent) |
| [#6734](https://github.com/nearai/ironclaw/issues/6734) | Agent self‑documentation (tool/channel config) | v1.1 |
| [#6731](https://github.com/nearai/ironclaw/issues/6731) | IronHub integration (marketplace for tools/skills) | v1.2+ |
| [#6727](https://github.com/nearai/ironclaw/issues/6727) | Custom/arbitrary MCP server support | v1.1 |
| [#6733](https://github.com/nearai/ironclaw/issues/6733) | `/model` and `/status` commands across all channels | v1.1 |
| [#6732](https://github.com/nearai/ironclaw/issues/6732) | Unified outbound delivery and automation targeting | v1.1 |
| [#6730](https://github.com/nearai/ironclaw/issues/6730) | Memory provider lifecycle corrections | v1.0.1 (hotfix) |
| [#6729](https://github.com/nearai/ironclaw/issues/6729) | Normalize extension installation persistence | v1.1 |
| [#6641](https://github.com/nearai/ironclaw/issues/6641) | Skill self‑creation (design doc) | v1.2+ |
| [#6481-6484](https://github.com/nearai/ironclaw/issues/6481) | Unified extension platform, Telegram hardening, pluggable memory, shared messaging | v1.1 |

**Prediction:** The next point release (v1.0.1) will focus on bug fixes and memory lifecycle corrections. The v1.1 release will likely ship **custom MCP server support**, **channel‑level commands** (`/model`, `/status`), and the **unified extension manifest** (Manifest V3) – all actively being built now.

---

## 7. User Feedback Summary

**Pain points addressed today:**
- **DeepSeek API 400 errors** (duplicate model field) – fixed, user `darren2013` reported it, now resolved.
- **Routine delivery misconfiguration** (leaks to wrong channels) – fixed, user `sergeiest` reported, now per‑routine.
- **Systemd onboarding failure** on Ubuntu – fixed, user `fadeevab` reported, root cause addressed.

**Emerging user needs (from new epics):**
- Users want the agent to be able to explain its own configuration (self‑docs #6734) – indicating current output is confidently wrong.
- Users want a marketplace for tools (IronHub #6731) – showing demand for extensibility beyond bundled tools.
- Users want custom MCP servers (#6727) – limiting factor for integrations.

**Satisfaction signals:** The v1.0.0 release is a major milestone; the team is actively responding to bugs and feature requests. No negative sentiment is evident beyond the normal growing pains of a large rewrite.

---

## 8. Backlog Watch

No issues or PRs appear to have been neglected for more than a few days. All open items have been updated within the last 48 hours. A few items to monitor:

- **[#6284 – Error‑Recoverability Endgame](https://github.com/nearai/ironclaw/issues/6284)** – 14 comments, created July 19. While active, this is a large epic that could stall if sub‑tasks are not resourced. Three blocking PRs (#6684, #6697, #6738) are progressing well.

- **[#6524 – Hermetic Testing Platform](https://github.com/nearai/ironclaw/issues/6524)** – Created July 22, spawned multiple PRs. Good momentum; no concern.

- **[#6726 – No‑op function passes all tests](https://github.com/nearai/ironclaw/issues/6726)** – Created yesterday, no assignee yet. Could indicate a test quality gap that should be addressed before it propagates.

- **[#5598 – Release PR](https://github.com/nearai/ironclaw/pull/5598)** – Open since July 3, appears to be a recurring automation for crate releases. Not a blocker but could be simplified.

**Overall backlog health:** The team is keeping pace with incoming issues and PRs. No critical neglect.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-07-28

Generated from GitHub data for the LobsterAI project (netease-youdao/LobsterAI) as of 2026-07-28.

---

## 1. Today’s Overview

The LobsterAI project experienced a high-activity day on July 27, with 7 issues and 9 pull requests updated, reflecting strong community engagement and a focused burst of code contributions. No new releases were published. The development team merged or closed 5 pull requests, including a critical fix for email attachment path traversal and a new Artifact preview toolbar with share/deploy capabilities. However, several severe bugs were reported — notably a data-corruption issue in the accelerator that silently mutates `\f` byte pairs — and three stale issues remain unresolved for months, indicating areas where maintainer attention is needed.

---

## 2. Releases

No new releases were created in the last 24 hours.

---

## 3. Project Progress

Five pull requests were merged or closed today, representing significant forward motion:

- **[#2389] fix(email): prevent attachment path traversal** (CLOSED)  
  Sanitizes attachment filenames and enforces download directory boundaries. Includes cross-platform security tests and bumps the bundled email skill version.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/pull/2389)

- **[#2388] feat(artifacts): 新增预览工具栏分享与部署入口** (CLOSED)  
  Adds share button to Artifact file preview toolbar and a deploy entry in the browser toolbar. Refines the toolbar publish-target strategy with unit tests and adds instrumentation for share/deploy entry points.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/pull/2388)

- **[#2387] Feat/2026.7.20 sites** (CLOSED)  
  Implements the “sites” feature, enhancing the project’s site management capabilities. (Fixes linked issue, not specified in PR description.)  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/pull/2387)

- **[#2386] fix(agentEngine): terminate no-progress tool loops before token budget exhaustion** (CLOSED)  
  Prevents runaway tool-calling loops by stopping them before reaching the token budget, improving reliability for agents that get stuck.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/pull/2386)

- **[#1323] fix(cowork): narrow input-too-long error classification** (CLOSED)  
  Corrects overly broad error classification where unrelated messages containing `max_tokens` could trigger a misleading “input too long” UI.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/pull/1323)

**Open PRs still pending (4):**  
- [#1277] chore(deps-dev): bump the electron group (stale, 4 months open)  
- [#2394] Fix/windows install manual overwrite blocked (new, needs review)  
- [#1239] feat(main): flash taskbar/Dock icon on task completion (stale, 4 months)  
- [#1241] feat(settings): close confirmation for unsaved changes (stale, directly addresses Issue #1237)

---

## 4. Community Hot Topics

Three issues and one PR attracted the most discussion or reactions today:

- **Issue #1237** [OPEN, stale] — *Settings 关闭无确认，API Key 等配置静默丢失*  
  Highlights a UX gap: users can lose API key configuration changes by accidentally closing the settings dialog. A fix PR (#1241) exists but is also stale. Underlying need: better data-integrity guardrails for configuration changes.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/issues/1237)

- **Issue #1240** [OPEN, stale] — *现有大模型受限后无法切换到其他大模型，所有对话框任务都会受限*  
  Reports a critical failure mode when one API provider is rate-limited: the entire application becomes unusable even for other providers. This suggests a global-state issue that blocks model switching.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/issues/1240)

- **Issue #2062** [OPEN, stale] — *任务超过最大时长*  
  Involves a task timeout when attempting 24-hour continuous runs, with unclear user feedback about whether the task continues in the background. Points to a need for better long-running task UX and timeout handling.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/issues/2062)

- **Issue #2393** [OPEN, new] — *LobsterAI 加速器在字符串改写时把 \f 字节对替换为 \x0C (form feed)*  
  Already a hot topic due to its **severity**: the accelerator tool silently corrupts data by transforming literal `\f` byte sequences into a form-feed control character. Community reaction is likely high given the data-loss nature.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/issues/2393)

---

## 5. Bugs & Stability

Three new bugs were reported today, ranked by severity:

🟥 **CRITICAL — Issue #2393: Accelerator corrupts `\f` byte pairs**  
  The accelerator tool replaces the literal bytes `5C 66` (`\f`) with the control character `\x0C` (form feed). This silently corrupts any file written by the tool that contains tokens like `\foo`, `\filename`, or `\firecrawl`. Reproducible 100% on the same machine. This is a **data integrity bug** that requires urgent patching. No fix PR yet.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/issues/2393)

🟧 **HIGH — Issue #2390: exec tool defaults to Windows PowerShell 5.1 and breaks on Chinese usernames**  
  The `exec` tool hard-codes `powershell.exe` instead of using the system’s installed PowerShell 7 (`pwsh.exe`). When Windows usernames contain Chinese characters, encoding breaks. The user’s environment has username `M幸福`. No fix PR yet.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/issues/2390)

🟨 **MEDIUM — Issue #2392: Scheduled tasks cannot select agent or skill**  
  A functional limitation: the scheduled-task UI does not allow choosing which agent or which skill to use, reducing the utility of automation. No fix PR yet.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/issues/2392)

**Stale bugs still open (no movement in 3+ months):**  
- #1237 (Settings unsaved changes lost — fix PR #1241 exists but stale)  
- #1240 (Global model-switching lock when one provider limited)  
- #2062 (Tasks exceeding maximum duration)

---

## 6. Feature Requests & Roadmap Signals

Two feature requests emerged today that could influence the next release:

- **Issue #2392: Scheduled tasks should support agent and skill selection**  
  User wants the ability to pick which agent and skill a scheduled task uses. This expands the automation feature and would likely be welcomed by power users. Given the existing scheduled-task infrastructure, this seems like a moderate-scope addition that could appear in the next minor release.

- **Issue #2391: Skill renaming**  
  A simple UX request: allow users to rename skills after creation. Low implementation cost, high user-impact gain. Likely to be picked up quickly, possibly in the 2026.7.x patch series.

**Roadmap signal from merged PRs:**  
The merger of #2388 (Artifact share/deploy toolbar) and #2387 (sites feature) suggests the project is investing in artifact delivery and site management — likely pushing toward a “build and share” workflow. The fix for no-progress tool loops (#2386) indicates a focus on agent stability and token efficiency.

---

## 7. User Feedback Summary

Real user pain points expressed in today’s activity:

- **Data corruption anxiety** (#2393): Users discovered that the accelerator silently damages file contents, leading to “bytes abnormal” when saving memory or configuration files. This is a serious trust-breaking issue.
- **Single point of failure** (#1240, stale): When one API key is rate-limited, the entire app blocks all models — users cannot switch to a working provider, rendering the tool unusable for hours.
- **Config loss without warning** (#1237, stale): API key configuration can be accidentally discarded. Users want a simple “unsaved changes” confirmation dialog.
- **Task automation limitations** (#2392, #2062): Inability to configure scheduled tasks with specific agents/skills, and unclear timeout behavior for long-running tasks (users report not knowing if background tasks continue).
- **Cross-platform desktop polish** (#2390): Windows users with non-ASCII usernames face encoding failures in the exec tool, which breaks workflows.

Overall sentiment: **High dissatisfaction** with data-integrity issues and stale bugs that block core workflows. But **appreciation** for the rate of new features (Artifact sharing, sites, agent loop fixes) that show the team is investing in the platform.

---

## 8. Backlog Watch

Several important issues and PRs have gone months without maintainer attention:

- **Issue #1237** (created Apr 1, 2026, stale) — *Settings loss*  
  Affects every user who configures API keys. A fix PR (#1241) has been open for 4 months.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/issues/1237)

- **Issue #1240** (created Apr 1, 2026, stale) — *Global model lock after rate limit*  
  A critical stability bug that can disable the entire application. No fix PR exists.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/issues/1240)

- **Issue #2062** (created May 27, 2026, stale) — *Task timeout for long runs*  
  User trying to build 24-hour continuous tasks hits an opaque timeout with unclear feedback.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/issues/2062)

- **PR #1239** (created Apr 1, 2026, stale) — *Flash taskbar on task completion*  
  A UI polish feature PR that has not been merged or reviewed for 4 months.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/pull/1239)

- **PR #1241** (created Apr 1, 2026, stale) — *Setting close confirmation*  
  The fix PR for Issue #1237, also untouched since April.  
  [GitHub Link](https://github.com/netease-youdao/LobsterAI/pull/1241)

**Recommendation:** The maintainers should prioritize reviewing and merging #1241 (settings safety) and investigating #1240 (model lock) to resolve two of the oldest and most impactful items in the backlog. The new critical bug #2393 also demands an immediate hotfix.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest — 2026-07-28

## 1. Today's Overview

Moltis saw no issue activity in the last 24 hours, while five pull requests received updates — all remaining open with none merged or closed. No new releases were published. The project's pulse is dominated by feature additions and infrastructure improvements, particularly around memory backends, ACP agent exposure, channel security, instrumentation, and PWA reliability. While no bugs or regressions were reported, several PRs have been open for over a week, suggesting a backlog of review work for maintainers.

## 2. Releases

*No new releases were published in the last 24 hours.*

## 3. Project Progress

**Merged/Closed PRs today:** 0  
No pull requests were merged or closed. All five updated PRs remain open, indicating that feature work is ongoing but has not yet reached a conclusion.

## 4. Community Hot Topics

*No issues or PRs have accumulated comments or reactions in the last 24 hours.* The five active PRs are the only visible activity:

- **[#1158 – feat(memory): add zvec vector database memory backend](https://github.com/moltis-org/moltis/pull/1158)** (demyanrogozhin)  
  A new, optional backend for memory using Zvec and Redb, gated behind a `zvec` feature. Reflects interest in flexible, self-hosted memory solutions.

- **[#1169 – feat(acp): expose Moltis as an ACP agent over stdio](https://github.com/moltis-org/moltis/pull/1169)** (penso)  
  Makes Moltis usable *as* an ACP agent (not just a client), enabling integration with ACP harnesses like Zed and buzz-acp.

- **[#1170 – fix(channels): gate /sh and privileged tools behind a per-account operators list](https://github.com/moltis-org/moltis/pull/1170)** (penso)  
  Closes a security gap where `/sh` was accessible to any channel member. Adds an operator whitelist per account.

- **[#1173 – feat(pwa): make push notifications reliable and non-disruptive](https://github.com/moltis-org/moltis/pull/1173)** (penso)  
  Fixes silent notification replacement and adds `renotify` flag; improves PWA notification UX.

- **[#1174 – Add instrumentation and feedback collection infrastructure](https://github.com/moltis-org/moltis/pull/1174)** (penso)  
  Introduces an `ObservationSink` fanout for agent runtime instrumentation, enabling pluggable backends and end-user feedback.

## 5. Bugs & Stability

No bugs, crashes, or regressions were reported in the last 24 hours. However, the **security fix in PR #1170** addresses a potential arbitrary command execution vulnerability in `/sh` on shared Discord guilds, which was previously accessible to all authorized channel members. No associated bug report is open; the fix is preemptive.

## 6. Feature Requests & Roadmap Signals

The open PRs themselves signal upcoming features likely to land in the next release:

- **Zvec memory backend** (#1158) — a lightweight, embeddable vector database alternative.
- **ACP agent role** (#1169) — positions Moltis as a drop-in agent for third-party ACP hosts.
- **Operator controls** (#1170) — adds per-account operator lists, a requested security hardening for multi-user deployments.
- **PWA reliability** (#1173) — addresses notification usability pain points.
- **Instrumentation & feedback** (#1174) — lays groundwork for telemetry and user-driven improvements.

These additions suggest the next release will focus on **expandability** (new backends and protocols), **security**, and **observability**.

## 7. User Feedback Summary

No direct user feedback (issues, comments) appeared today. Inferred from PRs:

- **Pain point (security):** Channel-level `/sh` access is too broad for multi-user Discord servers. The operator list in #1170 addresses this.
- **Pain point (UX):** PWA push notifications silently replaced messages without alerting users. #1173 fixes this with `renotify`.
- **Use case (flexibility):** Users want to run Moltis with alternative vector databases (Zvec/Redb) beyond the default, especially with custom embedding services.
- **Use case (integration):** Developers want to use Moltis as an ACP agent inside tools like Zed, not just as a client orchestrating other agents.

## 8. Backlog Watch

**No open issues exist,** but several pull requests have been open without merge for an extended period:

- **PR #1158** (zvec memory backend) — opened **July 17**, updated **July 28** (11 days). Waits for review and potential merge.
- **PRs #1169, #1170, #1173, #1174** — opened **July 26–27**, still pending maintainer attention.

These PRs represent significant feature work and security hardening. Prompt review would reduce the risk of merge conflicts and keep the development cycle moving.

---

*Digest generated from GitHub data (moltis-org/moltis) as of 2026-07-28.*

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest – 2026-07-28

**Data snapshot**: 19 issues updated in the last 24h (14 open, 5 closed) | 49 PRs updated (34 open, 15 merged/closed) | 0 new releases

---

## 1. Today's Overview

The QwenPaw project (part of CoPaw) continues to show **high activity** with a substantial volume of both bug reports and feature contributions. In the past 24 hours, the community closed 5 issues and merged or closed 15 PRs, while 14 issues and 34 PRs remain open, indicating a healthy development pace. The project is actively addressing **model provider compatibility** (multiple OpenAI-compatible endpoint issues), **mission‑mode session management**, and **cross‑platform sandbox support**. No new releases were cut today, but several major feature PRs (computer‑use, browser extension, third‑party agent integration) are under review, suggesting a forthcoming feature‑rich release in the coming weeks.

---

## 2. Releases

**None** – no new versions were published in the reporting window.

---

## 3. Project Progress

- **5 issues closed** (resolved or answered):  
  - [#6467] – Server setup question for proxy nodes (answered)  
  - [#6239] – Windows PATH semicolon drop bug (fix likely in progress)  
  - [#6464] – Model connection failure on AgentScope Platform  
  - [#6177] – Cron mode `final` delivery not working  
  - [#6406] – Windows `execute_shell_command` collapses multiline

- **15 PRs merged/closed** (full list not shown, but notable among the updated PRs are):  
  - [#6462] – Docs: clarified native Windows sandbox support (AppContainer/restricted-token)  
  - [#6502] – Fix: `dev` install instructions now include the `test` extra for contributors (first‑time contributor PR)  
  - [#6489] – Test: added Driver unit tests with `fail_under=50` coverage gate (improves CI quality)  
  - [#6068] – Fix: session ID preservation during Scroll history migration  

- **Major feature PRs under active review** (open):  
  - [#6424] – Native desktop GUI automation (computer‑use) for Windows and macOS  
  - [#6157] – Chrome extension plugin with native messaging bridge  
  - [#6397] – Third‑party agent integration (Codex, Qoder, Skills, MCP)  
  - [#6269] – Workspace checkpoint management via shadow Git  
  - [#6284] – QwenPaw Creator app (storyboard‑to‑video workflow)

---

## 4. Community Hot Topics

| Issue / PR | Comments | Topic |
|------------|----------|-------|
| [#6258] (open) | 4 | OpenAI `max_tokens` setting not respected |
| [#6324] (open) | 3 | Response truncation with MiniMax‑M3 model |
| [#6457] (open) | 3 | Mission mode creates excessive sub‑sessions in history |
| [#6460] (open) | 3 | High CPU in Edge+Wayland when viewing large sessions |
| [#6239] (closed) | 3 | Windows PATH concatenation bug (affects npm/resolve) |

**Underlying needs**:  
- Users expect reliable token‑limit enforcement across all providers (#6258).  
- The truncation bug (#6324) indicates incomplete support for non‑OpenAI models.  
- The session clutter (#6457, [#6507]) reflects a UX gap in mission‑mode history management.  
- Performance issue (#6460) suggests inefficiencies in rendering large conversation trees or WebSocket event processing.

---

## 5. Bugs & Stability

Reported bugs ranked by severity:

| Severity | Issue | Description | Fix PR exists? |
|----------|-------|-------------|----------------|
| **Critical** | [#6505] | Mission mode spawns unbounded sub‑sessions – no server‑side iteration cap, only stops when LLM balance runs out | No |
| **Critical** | [#6506] | `spawn_subagent` child sessions do not inherit parent’s `approval_level=OFF`, causing unwanted approval prompts | Yes – [#6508] (under review) |
| **High** | [#6258] | OpenAI `max_tokens` not applied – model generates beyond limit | No |
| **High** | [#6324] | Responses truncated when using MiniMax‑M3 | No |
| **High** | [#6386] | Tool calls repeated (e.g., file sent multiple times) | No |
| **High** | [#6358] | Context injection as `role='system'` fails on GLM/OpenAI when system message appears mid‑list | No |
| **High** | [#6496] | Legacy plugins silently disabled on QwenPaw 2.0+ due to implicit `max_version` derivation | No |
| **Medium** | [#6460] | High CPU on Edge+Wayland when browsing large sessions | No |
| **Medium** | [#6406] | Windows `execute_shell_command` collapses multiline PowerShell | (issue closed) |
| **Medium** | [#6177] | Cron `final` mode still forwards all events | (issue closed) |
| **Low** | [#6501] | Development install missing `test` extra, breaking `pytest` | Yes – [#6502] merged |

---

## 6. Feature Requests & Roadmap Signals

- **Multi‑model agent** ([#6455]): Allow one agent to run the same query on multiple models simultaneously and aggregate results. High demand for “ask all and merge” workflows.
- **Sub‑agent session grouping** ([#6507]): Filter or collapse sub‑sessions in chat history to reduce clutter.
- **New model providers** ([#6490], [#6498]): Volcano Engine Agent Plan, Xiaomi MiMo, Atlas Cloud – all submitted this week, indicating growing ecosystem demand.
- **Visual context compression** ([#6456]): “Visual Compact” for long agent histories using PawFocus – under review.
- **On‑demand SDK installation** ([#6387]): Allow installing missing channel SDKs from the Console – under review.
- **Computer‑use** ([#6424]): Native desktop automation for Windows/macOS.
- **Browser extension** ([#6157]): Chrome plugin for web automation.
- **Third‑party agents** ([#6397]): Integration with Codex, Qoder, Skills, MCP.

**Likely next‑version candidates**: Multi‑model agent (high user interest), provider additions, approval‑level inheritance fix, and one or both of the browser/computer‑use features.

---

## 7. User Feedback Summary

**Pain points** (real user reports):  
- “History is cluttered with sub‑sessions after mission mode” (#6457, #6507).  
- “Setting approval to OFF doesn’t apply to spawned workers” (#6506).  
- “Connection test fails for all models after deploying” (#6464).  
- “Max output token does not work with OpenAI models” (#6258).  
- “Response gets truncated when using MiniMax‑M3” (#6324).  
- “CPU goes high when viewing a session with many images/tool calls” (#6460).  

**Satisfaction signals**:  
- Active bug reporting and feature suggestions indicate an engaged user base.  
- Several contributors submitted PRs for providers and documentation – community trust is growing.  
- The quick closure of issues like #6177 and #6406 shows maintainers are responsive.

---

## 8. Backlog Watch

The following issues or PRs have been open for an extended period and may need maintainer attention:

| Item | Created | Days Open | Description |
|------|---------|-----------|-------------|
| [#5490] (PR) | 2026-06-24 | 34 | Console: show tool‑card images inline and gallery navigation – **no review activity in weeks** |
| [#6068] (PR) | 2026-07-13 | 15 | Fix session ID preservation in Scroll migration – **ready for merge?** |
| [#6258] (issue) | 2026-07-19 | 9 | OpenAI `max_tokens` bug – **no response from maintainers** |
| [#6358] (issue) | 2026-07-22 | 6 | Context injection causes `ValueError` on GLM/OpenAI – **no fix PR yet** |
| [#6324] (issue) | 2026-07-22 | 6 | MiniMax‑M3 response truncation – **no fix PR yet** |

**Notable**: PR #5490 has been open for over a month with no updates. It addresses a core console UX improvement – inline image display – and may benefit from a maintainer review or reassignment.

[#6258]: https://github.com/agentscope-ai/QwenPaw/issues/6258
[#6324]: https://github.com/agentscope-ai/QwenPaw/issues/6324
[#6457]: https://github.com/agentscope-ai/QwenPaw/issues/6457
[#6460]: https://github.com/agentscope-ai/QwenPaw/issues/6460
[#6467]: https://github.com/agentscope-ai/QwenPaw/issues/6467
[#6239]: https://github.com/agentscope-ai/QwenPaw/issues/6239
[#6464]: https://github.com/agentscope-ai/QwenPaw/issues/6464
[#6177]: https://github.com/agentscope-ai/QwenPaw/issues/6177
[#6406]: https://github.com/agentscope-ai/QwenPaw/issues/6406
[#6505]: https://github.com/agentscope-ai/QwenPaw/issues/6505
[#6506]: https://github.com/agentscope-ai/QwenPaw/issues/6506
[#6507]: https://github.com/agentscope-ai/QwenPaw/issues/6507
[#6501]: https://github.com/agentscope-ai/QwenPaw/issues/6501
[#6496]: https://github.com/agentscope-ai/QwenPaw/issues/6496
[#6358]: https://github.com/agentscope-ai/QwenPaw/issues/6358
[#6386]: https://github.com/agentscope-ai/QwenPaw/issues/6386
[#6455]: https://github.com/agentscope-ai/QwenPaw/issues/6455
[#6490]: https://github.com/agentscope-ai/QwenPaw/issues/6490
[#6498]: https://github.com/agentscope-ai/QwenPaw/issues/6498
[#6456]: https://github.com/agentscope-ai/QwenPaw/pull/6456
[#6424]: https://github.com/agentscope-ai/QwenPaw/pull/6424
[#6157]: https://github.com/agentscope-ai/QwenPaw/pull/6157
[#6397]: https://github.com/agentscope-ai/QwenPaw/pull/6397
[#6269]: https://github.com/agentscope-ai/QwenPaw/pull/6269
[#6284]: https://github.com/agentscope-ai/QwenPaw/pull/6284
[#6387]: https://github.com/agentscope-ai/QwenPaw/pull/6387
[#6068]: https://github.com/agentscope-ai/QwenPaw/pull/6068
[#6489]: https://github.com/agentscope-ai/QwenPaw/pull/6489
[#6502]: https://github.com/agentscope-ai/QwenPaw/pull/6502
[#6508]: https://github.com/agentscope-ai/QwenPaw/pull/6508
[#6462]: https://github.com/agentscope-ai/QwenPaw/pull/6462
[#5490]: https://github.com/agentscope-ai/QwenPaw/pull/5490

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-07-28

## Today's Overview
ZeroClaw saw very high activity over the last 24 hours, with **24 issues** and **50 pull requests** updated. The project remains in an intensive bug-fixing and feature‑integration cycle, with **2 issues closed** and **8 PRs merged/closed**. No new releases were published. Several **workflow‑blocking (S1) bugs** were reported, notably a missing cancellation path for running SOP jobs and a broken auth‑profile store after a field rename. The daemon‑owned SOP control plane tracker (#8288) and the v0.8.5 weekly release tracker (#9459) indicate ongoing roadmap work. Many of the day’s updates involve test‑infrastructure improvements (flaky tests, CI gaps, Windows compatibility) and security/telemetry hardening.

## Releases
*None in the reporting period.*

## Project Progress (Merged/Closed PRs & Issues)
The following PRs were merged or closed today (all targeting `master`):

- **#9298** [bug, tests] – Classify integration tests by path component in config‑save gate (closed).  
  *Fixed a Windows‑specific test‑safety net.*  
  [PR #9298](https://github.com/zeroclaw-labs/zeroclaw/pull/9298)

- **#9388** [docs, governance] – Retire `CONTRIBUTORS.md` and ground maintainer roles in FND‑003 (closed).  
  [PR #9388](https://github.com/zeroclaw-labs/zeroclaw/pull/9388)

- **#9251** [enhancement, infrastructure] – PostgreSQL as the first supported session backend (closed).  
  *A large PR that reduces a planned five‑backend series to foundation + PostgreSQL, per maintainer guidance.*  
  [PR #9251](https://github.com/zeroclaw-labs/zeroclaw/pull/9251)

- **#9442** [tests] – Stop wall‑clock guards in channels tests that asserted scheduling behaviour (closed).  
  *Addresses flakiness on slow CI runners.*  
  [PR #9442](https://github.com/zeroclaw-labs/zeroclaw/pull/9442)

- **#9434** [dependencies] – Bump `rust-all` group with 44 updates (closed).  
  [PR #9434](https://github.com/zeroclaw-labs/zeroclaw/pull/9434)

Closed issues today:

- **#9429** [bug, CI] – Channels tests use fixed wall‑clock timeouts – flaky on slow runners (closed).  
  [Issue #9429](https://github.com/zeroclaw-labs/zeroclaw/issues/9429)

- **#9238** [bug, tests] – `config_save_isolation` skips all test files on Windows (closed).  
  [Issue #9238](https://github.com/zeroclaw-labs/zeroclaw/issues/9238)

## Community Hot Topics
The most discussed items in the past 24 hours (by comment count) reveal three underlying concerns:

1. **Test flakiness and CI reliability**  
   - **#9357** (5 comments) – `cargo test -p zeroclaw-runtime --lib` fails 19/20 runs on master; a flaky assertion poisons a global mutex.  
     [Issue #9357](https://github.com/zeroclaw-labs/zeroclaw/issues/9357)

2. **Localization gaps**  
   - **#9363** (3 comments) – Config metadata remains English in localized ZeroCode and web surfaces.  
     [Issue #9363](https://github.com/zeroclaw-labs/zeroclaw/issues/9363)

3. **Security token leaks and missing cancellation paths**  
   - **#9417** (2 comments) – WhatsApp Cloud `request_approval` leaks a live approval token on send failure or cancellation.  
     [Issue #9417](https://github.com/zeroclaw-labs/zeroclaw/issues/9417)

Other frequently updated items include the **RFC for AI‑assisted PR pre‑review** (#9330, 2 comments) and the **SOP cancellation path bug** (#9425, 1 comment but high urgency).

## Bugs & Stability
Bugs reported or updated today, ranked by severity:

### S1 – Workflow Blocked
- **#9425** – Running SOP jobs have no operator cancellation path on the web dashboard.  
  [Issue #9425](https://github.com/zeroclaw-labs/zeroclaw/issues/9425)  
  *No fix PR yet.*
- **#9474** – Auth profile store fails to load because `model_provider` field is required with no migration from pre‑rename stores.  
  [Issue #9474](https://github.com/zeroclaw-labs/zeroclaw/issues/9474)  
  *No fix PR yet.*

### S2 – Degraded Behaviour
- **#9357** – Runtime lib test flaky (19/20 failures); global mutex poisoning.  
  [Issue #9357](https://github.com/zeroclaw-labs/zeroclaw/issues/9357)  
  *Fix PRs likely in progress (see #9424, #9447).*
- **#9363** – Config metadata not localized.  
  [Issue #9363](https://github.com/zeroclaw-labs/zeroclaw/issues/9363)
- **#9417** – WhatsApp approval token leak.  
  [Issue #9417](https://github.com/zeroclaw-labs/zeroclaw/issues/9417)
- **#9340** – CLI‑created cron jobs have delivery hardcoded to `None`.  
  [Issue #9340](https://github.com/zeroclaw-labs/zeroclaw/issues/9340)
- **#9463** – WASM memory plugins are compiled and tested in isolation but never wired into runtime backend selection.  
  [Issue #9463](https://github.com/zeroclaw-labs/zeroclaw/issues/9463)
- **#9465** – Inbound channel messages declined by precheck produce only a reaction – no text delivered to sender.  
  [Issue #9465](https://github.com/zeroclaw-labs/zeroclaw/issues/9465)

### S3 – Minor
- **#9462** – `zeroclaw-plugins` lib unit tests behind `plugins-wasmtime` feature never run in CI.  
  [Issue #9462](https://github.com/zeroclaw-labs/zeroclaw/issues/9462)
- **#9408** – Telegram bot‑command menu descriptions not localized (English‑only literals).  
  [Issue #9408](https://github.com/zeroclaw-labs/zeroclaw/issues/9408)

Many of these bugs have associated fix PRs that are open (e.g., #9424, #9447, #9448, #9449, #9472). The project is actively addressing the stability concerns.

## Feature Requests & Roadmap Signals
New or updated feature requests and RFCs in the past 24 hours:

- **#9464** – [RFC] Anthropic stored‑profile OAuth alias contract.  
  [Issue #9464](https://github.com/zeroclaw-labs/zeroclaw/issues/9464)
- **#9336** – Render `TodoWrite` plan events in the web dashboard.  
  [Issue #9336](https://github.com/zeroclaw-labs/zeroclaw/issues/9336)
- **#9460** – Harden Windows key‑file ACLs at creation.  
  [Issue #9460](https://github.com/zeroclaw-labs/zeroclaw/issues/9460)
- **#8858** – Tracker: audit existing drift surfaces across codebase.  
  [Issue #8858](https://github.com/zeroclaw-labs/zeroclaw/issues/8858)
- **#8288** – Tracker: SOP milestone – daemon‑owned SOP control plane to 5/5.  
  [Issue #8288](https://github.com/zeroclaw-labs/zeroclaw/issues/8288)
- **#9459** – Tracker: v0.8.5 weekly non‑breaking release.  
  [Issue #9459](https://github.com/zeroclaw-labs/zeroclaw/issues/9459)
- **#9346** – [RFC] Unified package/capability/config/runtime‑state catalog contract.  
  [Issue #9346](https://github.com/zeroclaw-labs/zeroclaw/issues/9346)
- **#9473, #9471, #9470** – Cleanup tasks: recover disabled tests, retire dormant test modules, correct telemetry attribution.  
  [Issue #9473](https://github.com/zeroclaw-labs/zeroclaw/issues/9473) | [#9471](https://github.com/zeroclaw-labs/zeroclaw/issues/9471) | [#9470](https://github.com/zeroclaw-labs/zeroclaw/issues/9470)

**Predictions for v0.8.5:** The release tracker (#9459) will likely include the PostgreSQL session backend (#9251, already merged), the skills compact‑injection default (#8313, still open but high activity), and several of the S2 bug fixes currently in PR. The unified catalog RFC (#9346) and AI‑assisted review RFC (#9330) are longer‑term.

## User Feedback Summary
Real pain points visible in the last 24‑hour activity:

- **Test flakiness** (#9357, #9429) wastes developer time and undermines CI confidence. Users like `AngryPacifist` and `MannXo` are reporting systemic issues.
- **Localization incompleteness** (#9363, #9408) frustrates non‑English users who see a partially translated interface.
- **Missing cancellation** for SOP jobs (#9425) blocks operators from stopping runaway workflows.
- **Silent failures** – cron output discarded without notice (#9340), precheck‑declined messages give only a reaction (#9465) – make the system appear broken.
- **Security/credential management** – WhatsApp token leak (#9417) and auth‑profile migration failure (#9474) are critical for enterprise adoption.
- **Windows CI gaps** have been partially addressed (#9238, #9298), but users still encounter platform‑specific issues.

Satisfaction: The steady stream of merged PRs (e.g., PostgreSQL backend, test fixes) shows active maintenance and responsiveness.

## Backlog Watch
Issues and PRs that have been open for an extended period and still need maintainer attention:

- **#8692** – Maintainer decision queue for RFCs and design issues (created Jul 4, last updated Jul 27, 1 comment).  
  [Issue #8692](https://github.com/zeroclaw-labs/zeroclaw/issues/8692)

- **#8288** – SOP milestone tracker (created Jun 24, last updated Jul 27, 1 comment).  
  [Issue #8288](https://github.com/zeroclaw-labs/zeroclaw/issues/8288)

- **#8858** – Drift surface audit tracker (created Jul 8, last updated Jul 27).  
  [Issue #8858](https://github.com/zeroclaw-labs/zeroclaw/issues/8858)

- **#9346** – RFC for unified catalog contract (created Jul 24, 0 comments).  
  [Issue #9346](https://github.com/zeroclaw-labs/zeroclaw/issues/9346)

- **#9330** – RFC for AI‑assisted PR review (created Jul 24, 2 comments).  
  [Issue #9330](https://github.com/zeroclaw-labs/zeroclaw/issues/9330)

- **PR #8313** – feat(skills): default to compact injection (created Jun 25, last updated Jul 27, labelled `needs-author-action`? actually not labelled, but no recent author activity).  
  [PR #8313](https://github.com/zeroclaw-labs/zeroclaw/pull/8313)

- **PR #8443** – feat(matrix): single‑message progress drafts (created Jun 28, last updated Jul 27, labelled `needs-author-action`).  
  [PR #8443](https://github.com/zeroclaw-labs/zeroclaw/pull/8443)

These items represent architectural decisions and large features that have been in review for weeks. Maintainer action (acceptance, rejection, or deferral) would help unblock dependent work.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*