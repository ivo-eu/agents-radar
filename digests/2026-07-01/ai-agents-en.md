# OpenClaw Ecosystem Digest 2026-07-01

> Issues: 65 | PRs: 500 | Projects covered: 13 | Generated: 2026-07-01 11:36 UTC

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

# OpenClaw Project Digest — 2026-07-01

## 1. Today's Overview
OpenClaw shows very high activity: **65 issues** updated in the last 24 hours (48 open, 17 closed) and **500 PRs** updated (415 open, 85 merged/closed). A new release **v2026.6.11** landed, targeting reliability fixes for misplaced replies, stuck sends, reconnects, and model setup failures. However, the release also introduced several **P1 regressions** (e.g., tool output returning empty after first call, session initialization conflicts) that are being actively triaged. The community is strongly engaged around memory trust and session continuity features, with three long-standing enhancement issues accumulating the highest comment counts. Overall project health appears robust but under post-release patch load.

## 2. Releases
### v2026.6.11 (2026-06-11)
**Focus:** Polish and dependability fixes for rough edges reported by users.

**Highlights:**
- Fixed misplaced replies and stuck sends
- Improved reconnection logic
- Resolved model setup failures
- Safer admin defaults (reduced implicit exposure)
- [Full release notes](https://docs.openclaw.ai/releases/2026.6.11)

**Known regressions in this release (reported today):**
- Tool output (exec, web_fetch, web_search) returns empty after first call per turn ([#98528](https://github.com/openclaw/openclaw/issues/98528) – P1)
- New session initialization immediately errors with “reply session initialization conflicted” ([#98562](https://github.com/openclaw/openclaw/issues/98562) – P1)
- Container image upgrades skip configured‑plugin convergence before gateway ready ([#98565](https://github.com/openclaw/openclaw/issues/98565) – P1)
- Mattermost repair does not enable external plugin after installation ([#98564](https://github.com/openclaw/openclaw/issues/98564) – P1)

**Migration notes:** No breaking changes are documented, but users upgrading to v2026.6.11 should verify tool calls and session startup behave correctly. A patch release is expected shortly.

## 3. Project Progress
**85 PRs** were merged or closed today. Notable advancements:

| PR | Focus | Status |
|----|-------|--------|
| [#98596](https://github.com/openclaw/openclaw/pull/98596) | Preserve profile env for macOS launch at login | Merged |
| [#98592](https://github.com/openclaw/openclaw/pull/98592) | Align macOS config path resolution | Merged |
| [#94964](https://github.com/openclaw/openclaw/pull/94964) | Cancel deferred channel reload on in‑process restart (fixes #79487) | Merged |
| [#90389](https://github.com/openclaw/openclaw/pull/90389) | Fix Mattermost slash commands returning permanent 503 | Merged |
| [#98568](https://github.com/openclaw/openclaw/pull/98568) | Add unit tests for Gmail watcher error classification | Merged |
| [#96344](https://github.com/openclaw/openclaw/pull/96344) | Fix Telegram channel crash loop during deferred reload | Merged |
| [#68936](https://github.com/openclaw/openclaw/pull/68936) | Autofix pipeline for PR review + Windows daemon | Merged |
| [#98104](https://github.com/openclaw/openclaw/pull/98104) | Log chat‑abort terminal persistence failures | Merged |

These fixes address core stability (channel reloading, macOS integration, authentication edge cases) and improve observability (test coverage, error logging).

## 4. Community Hot Topics
Most active issues by comment count (all **open**, high community interest):

| Issue | Comments | 👍 | Topic |
|-------|----------|----|-------|
| [#7707](https://github.com/openclaw/openclaw/issues/7707) | 13 | 0 | Memory Trust Tagging by source – prevent memory poisoning from untrusted content |
| [#45608](https://github.com/openclaw/openclaw/issues/45608) | 11 | 4 | Pre‑reset agentic memory flush – `/new` and daily reset should flush memory like compaction |
| [#40418](https://github.com/openclaw/openclaw/issues/40418) | 8 | 1 | Automated session memory preservation & synthesis across sessions |
| [#90414](https://github.com/openclaw/openclaw/issues/90414) | 6 | 2 | `agentmemory__memory_search` returns “index metadata is missing” – persistent memory‑core bug |
| [#20935](https://github.com/openclaw/openclaw/issues/20935) | 6 | 0 | Audit log for agent memory changes – detect tampering |

**Underlying needs:** Users demand trustworthy memory systems (tagging, audit, safe reset), session continuity, and reliable memory search. The high engagement on memory‑related issues signals this is the top community priority.

Most commented PRs (all with low comment count, but significant size or maintainer attention):

| PR | Size | Topic |
|----|------|-------|
| [#98587](https://github.com/openclaw/openclaw/pull/98587) | S | Fix Slack relay WebSocket JSON.parse crash |
| [#70990](https://github.com/openclaw/openclaw/pull/70990) | XL | Add model failover and terminal failure hooks (waiting on author) |
| [#98572](https://github.com/openclaw/openclaw/pull/98572) | S | Fix memory sync‑on‑miss blocking tool calls (fixes #98520) |
| [#63015](https://github.com/openclaw/openclaw/pull/63015) | S | Honor filePath/path/media fallbacks in outbound reply normalization (needs proof) |

## 5. Bugs & Stability
Today’s bug reports, ranked by severity. Many are regressions from the latest release.

### P1 – Critical / Blockers
| Issue | Summary | Fix PR exists? |
|-------|---------|----------------|
| [#98528](https://github.com/openclaw/openclaw/issues/98528) | **Regression** Tool output empty after first call per turn (v2026.6.11) | No |
| [#98562](https://github.com/openclaw/openclaw/issues/98562) | **Regression** New session error “reply session initialization conflicted” | No |
| [#98565](https://github.com/openclaw/openclaw/issues/98565) | **Regression** Container image upgrades skip plugin convergence | No |
| [#98564](https://github.com/openclaw/openclaw/issues/98564) | **Regression** Mattermost not enabled after repair install | No |
| [#98573](https://github.com/openclaw/openclaw/issues/98573) | Windows – spawn gemini ENOENT (missing .cmd extension) | No |
| [#98556](https://github.com/openclaw/openclaw/issues/98556) | Heartbeat polls zombie main sessions, causing silent token burn until spend cap | No |
| [#98532](https://github.com/openclaw/openclaw/issues/98532) | `replay_invalid` thinking block error not retried – session broken permanently | No |
| [#98522](https://github.com/openclaw/openclaw/issues/98522) | Stuck session recovery aborts isolated cron jobs after ~6 min | No |
| [#98550](https://github.com/openclaw/openclaw/issues/98550) | Request CVE for authorization bypass (GHSA‑cf2p‑f286‑mphf) | No |

### P2 – High Impact
| Issue | Summary | Fix PR exists? |
|-------|---------|----------------|
| [#98540](https://github.com/openclaw/openclaw/issues/98540) | Composer shows idle state while agent actively executing tools | No |
| [#98560](https://github.com/openclaw/openclaw/issues/98560) | Plugin `node.invoke` policy approvals cannot use turn‑source routes | No |
| [#98534](https://github.com/openclaw/openclaw/issues/98534) | Plugin approvals can store blank titles | No |
| [#98529](https://github.com/openclaw/openclaw/issues/98529) | Old approval timers can remove newer same‑id prompts | No |
| [#98523](https://github.com/openclaw/openclaw/issues/98523) | Control UI reconnect does not reload pending approvals | No |
| [#98345](https://github.com/openclaw/openclaw/issues/98345) | **P0** memory‑wiki: transient read failure during re‑ingest silently wipes user Notes block | Closed with fix? Actually closed, likely fixed. |
| [#98557](https://github.com/openclaw/openclaw/issues/98557) | Telegram rich output exposes raw `<parameter name>` wrappers | No |
| [#98547](https://github.com/openclaw/openclaw/issues/98547) | macOS remote URL probe says `ws://` is localhost‑only | No |
| [#98537](https://github.com/openclaw/openclaw/issues/98537) | macOS Remote mode can leave local gateway launchd running | No |
| [#98486](https://github.com/openclaw/openclaw/issues/98486) | macOS Dashboard injects native device token as browser shared auth (security) | No |
| [#98466](https://github.com/openclaw/openclaw/issues/98466) | usage‑bar `in` operator matches Object.prototype inherited properties | No |

### Fixed Today (closed bugs)
- [#91167](https://github.com/openclaw/openclaw/issues/91167) – memory: gateway cannot self‑heal missing index identity (closed)
- [#94432](https://github.com/openclaw/openclaw/issues/94432) – OpenAI/Codex OAuth fails with Cloudflare HTML 403 (closed)
- [#79487](https://github.com/openclaw/openclaw/issues/79487) – Channel‑reload race causing EADDRINUSE crash loop (closed, fixed via #94964)
- [#66957](https://github.com/openclaw/openclaw/issues/66957) – `models.mode="replace"` still triggers implicit provider discovery (closed)
- [#97985](https://github.com/openclaw/openclaw/issues/97985) – Post‑update plugin sync false‑fails on bundle plugins (closed)
- [#98270](https://github.com/openclaw/openclaw/issues/98270) – `config.patch` rejects empty baseUrl for built‑in providers (closed)
- [#98520](https://github.com/openclaw/openclaw/issues/98520) – `memory_search` blocks foreground tool calls on sync (closed, fix #98572 open)
- [#98345](https://github.com/openclaw/openclaw/issues/98345) – memory‑wiki silently wipes Notes block (closed, likely fixed)

## 6. Feature Requests & Roadmap Signals
Top feature requests from today and recent weeks that are likely candidates for the next release:

| Issue | Feature | Community Traction |
|-------|---------|-------------------|
| [#7707](https://github.com/openclaw/openclaw/issues/7707) | **Memory Trust Tagging by Source** – tag entries by origin to prevent poisoning | 13 comments, diamond lobster rating |
| [#45608](https://github.com/openclaw/openclaw/issues/45608) | **Pre‑reset memory flush** – flush memory before `/new` and daily reset | 11 comments, 4 👍 |
| [#40418](https://github.com/openclaw/openclaw/issues/40418) | **Automated Session Memory Preservation & Synthesis** – continuous learning across sessions | 8 comments |
| [#20935](https://github.com/openclaw/openclaw/issues/20935) | **Audit Log for Memory Changes** – append‑only log for tamper detection | 6 comments |
| [#13615](https://github.com/openclaw/openclaw/issues/13615) | **Rate limiting and throttling** for external API calls | 3 comments, 2 👍 |
| [#98549](https://github.com/openclaw/openclaw/issues/98549) | **Scoped brokered credentials** for plugin operations (SecretRef) | New, maintainer‑tagged |
| [#98542](https://github.com/openclaw/openclaw/issues/98542) | **Conversation identity modes** for personal/team/shared audiences | New, maintainer‑tagged |
| [#98521](https://github.com/openclaw/openclaw/issues/98521) | **Metadata‑only audit records** for agent runs and tool actions | New, maintainer‑tagged |
| [#98498](https://github.com/openclaw/openclaw/issues/98498) | **Telegram silent messages** config option (disable_notification) | New, P3 |

**Prediction:** The next minor release (v2026.7.x) will likely focus on **memory security & auditability** (trust tagging, audit log, flush on reset) and **credential brokering** for plugins. The maintainer‑tagged features (#98549, #98542, #98521) indicate active design work.

## 7. User Feedback Summary
Real user pain points and use cases reflected in today’s data:

**Pain points:**
- **Memory/data loss:** Users report silent wiping of hand‑written notes during re‑ingest ([#98345](https://github.com/openclaw/openclaw/issues/98345)) and loss of session context after reset ([#40418](https://github.com/openclaw/openclaw/issues/40418)).
- **Regressions after update:** Multiple users downgrading or reverting after v2026.6.11 broke tool calls ([#98528](https://github.com/openclaw/openclaw/issues/98528)), session init ([#98562](https://github.com/openclaw/openclaw/issues/98562)), and Mattermost channels ([#98564](https://github.com/openclaw/openclaw/issues/98564)).
- **Approval UX:** Approval window disappears after resolving one of multiple pending approvals ([#98585](https://github.com/openclaw/openclaw/issues/98585)), and plugin tool approvals leave stale cards ([#98576](https://github.com/openclaw/openclaw/issues/98576)).
- **Plugin installation:** `@openclaw/acpx` publishes `workspace:*` devDep, breaking npm installations ([#98583](https://github.com/openclaw/openclaw/issues/98583)).
- **Performance:** Per‑request auth (5.5s) and tool bundling (8.9s) dominate gateway TTFT ([#80131](https://github.com/openclaw/openclaw/issues/80131)).

**Use cases driving feature requests:**
- **Team digital employees** need auditable, tamper‑resistant memory and credential scoping ([#98549](https://github.com/openclaw/openclaw/issues/98549), [#98542](https://github.com/openclaw/openclaw/issues/98542)).
- **Cron job operators** hit stuck‑session recovery prematurely ([#98522](https://github.com/openclaw/openclaw/issues/98522)).
- **Reverse‑proxy deployments** need `--context-path` support for WebSocket connections ([#97678](https://github.com/openclaw/openclaw/issues/97678)).
- **macOS remote mode** users face password auth unsupported ([#98552](https://github.com/openclaw/openclaw/issues/98552)) and leftover launchd processes ([#98537](https://github.com/openclaw/openclaw/issues/98537)).
- **Windows users** hit ENOENT when spawning native CLIs ([#98573](https://github.com/openclaw/openclaw/issues/98573)).

**Satisfaction:** Mixed. The v2026.6.11 release addressed long‑standing reliability issues, but the fresh regressions have dampened enthusiasm. Community feedback is detailed and constructive, indicating high engagement.

## 8. Backlog Watch
Long‑unanswered issues and PRs that require maintainer attention (no update for weeks or months, still open with activity needed):

| Issue/PR | Created | Latest Update | Comments | Status |
|----------|---------|---------------|----------|--------|
| [#7707](https://github.com/openclaw/openclaw/issues/7707) Memory Trust Tagging | 2026-02-03 | 2026-07-01 | 13 | Needs product decision & security review |
| [#45608](https://github.com/openclaw/openclaw/issues/45608) Pre‑reset memory flush | 2026-03-14 | 2026-07-01 | 11 | Needs product decision & security review |
| [#40418](https://github.com/openclaw/openclaw/issues/40418) Automated session memory preservation | 2026-03-09 | 2026-07-01 | 8 | Needs product decision & security review |
| [#20935](https://github.com/openclaw/openclaw/issues/20935) Audit log for memory | 2026-02-19 | 2026-07-01 | 6 | Needs product decision & security review |
| [#13615](https://github.com/openclaw/openclaw/issues/13615) Rate limiting for external API calls | 2026-02-10 | 2026-07-01 | 3 | Needs product decision & security review |
| [#90414](https://github.com/openclaw/openclaw/issues/90414) memory_search “index metadata missing” | 2026-06-04 | 2026-07-01 | 6 | Needs live reproduction & maintainer review |
| [#70990](https://github.com/openclaw/openclaw/pull/70990) Model failover hooks | 2026-04-24 | 2026-07-01 | 0 | Waiting on author |
| [#73338](https://github.com/openclaw/openclaw/pull/73338) TUI follow active gateway port | 2026-04-28 | 2026-07-01 | 0 | Waiting on author |
| [#72610](https://github.com/openclaw/openclaw/pull/72610) Hermes import path | 2026-04-27 | 2026-07-01 | 0 | Automerge armed |

**Call to action:** The five top‑voted memory‑related issues have been awaiting maintainer product decisions for 4‑5 months. Given the community’s clear interest, a roadmap commitment on memory trust and audit would reduce uncertainty. The stalled PRs (#70990, #73338, #72610) have been waiting on author responses for weeks.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: AI Agent & Personal Assistant Ecosystem
**Date**: 2026-07-01 | **Scope**: 12 open-source projects

---

## 1. Ecosystem Overview

The personal AI assistant open-source ecosystem is experiencing a **rapid maturation phase**, characterized by intense development velocity across multiple projects and a strong convergence on memory reliability, security hardening, and production-grade infrastructure. The ecosystem is bifurcating: core framework projects (OpenClaw, IronClaw) focus on runtime stability and enterprise features, while specialized agents (NanoBot, PicoClaw, NullClaw) prioritize platform integration and developer UX. A notable pattern is the **post-release regression cycle**—several projects shipped major updates only to face critical P1 bugs, signaling that the community is prioritizing shipping speed over exhaustive testing. The landscape remains dominated by OpenClaw-derivative forks, but differentiation is emerging through unique architectural choices (MoA in ZeroClaw, Hermes' salvage workflow, CoPaw's v2.0 pre-release blitz).

---

## 2. Activity Comparison

| Project | Issues Updated (Open/Closed) | PRs Updated (Merged/Closed) | Release This Period | Health Score | Notes |
|---------|------------------------------|----------------------------|---------------------|--------------|-------|
| **OpenClaw** | 65 (48/17) | 500 (85 merged) | ✅ v2026.6.11 | 🟡 Good | Post-release regressions dampen otherwise robust health |
| **IronClaw** | 9 (7/2) | 50 (30 merged) | ❌ | 🟡 Good | High merge velocity, focused on stability fixes |
| **CoPaw** | 12 (9/3) | 50 (25 merged) | ❌ (v2.0.0b2 pre-release) | 🟡 Good | Intense sprint mode but regressions in beta |
| **NanoBot** | 5 (2/3) | 41 (9 merged) | ❌ | 🟢 Very Good | Strong security response, no critical regressions |
| **NanoClaw** | 8 (7/1) | 17 (10 merged) | ❌ | 🟡 Fair | High bug-to-fix ratio; critical setup bug unresolved |
| **ZeroClaw** | 17 (16/1) | 50 (5 merged) | ❌ | 🟡 Fair | S1 blockers without fixes undermine velocity |
| **Hermes Agent** | 3 (1/2) | 50 (19 merged) | ❌ | 🟢 Very Good | Efficient salvage process; rapid bug closure |
| **LobsterAI** | 4 (3/1) | 26 (23 merged) | ✅ v2026.6.30 | 🟢 Very Good | Exceptional merge throughput; backlog clearing |
| **PicoClaw** | 4 (2/2) | 5 (2 merged) | ✅ Nightly | 🟡 Fair | Moderate; high-impact bug (#3153) unresolved |
| **NullClaw** | 1 (1/0) | 4 (4 merged) | ❌ | 🟢 Good | Clean merge activity; single issue stale |
| **Moltis** | 0 (0/0) | 3 (2 merged) | ❌ | 🟢 Dormant | Only Dependabot activity; no community engagement |
| **TinyClaw** | 0 (0/0) | 0 (0 merged) | ❌ | 🟢 Dormant | No activity in 24h |

**Health Scoring**: Very Good (no critical bugs, high merge velocity) → Good (active but some issues) → Fair (critical bugs unresolved or setup broken) → Dormant (no human-driven activity)

---

## 3. OpenClaw's Position

**Advantages vs Peers**:
- **Scale**: OpenClaw dwarfs all competitors in raw activity—500 PRs and 65 issues in 24h vs. the next highest (IronClaw/CoPaw at 50 PRs each). It is the **de facto reference implementation** with the largest contributor base.
- **Release cadence**: v2026.6.11 is the only GA release shipped today among core frameworks, demonstrating mature release engineering.
- **Community gravity**: The highest-comment issues (13, 11, 8 comments) exceed any other project's engagement, signaling a deeply invested user base.

**Technical Approach Differences**:
- **Memory architecture**: OpenClaw is pioneering memory trust tagging (#7707) and audit logs (#20935)—no other project has comparable proposals for provenance-aware memory. Others (NanoBot, PicoClaw) treat memory as a simpler key-value store.
- **Session continuity**: Automated session memory preservation (#40418) and pre-reset flush (#45608) are uniquely OpenClaw concerns, reflecting its role as the "operating system" for agents.
- **Plugin ecosystem**: OpenClaw's plugin system is the most mature, with active work on credential brokering (#98549) and scoped identities—features absent from NanoBot's WASM plugin and PicoClaw's simpler tool registry.

**Comparison Metrics**:
| Metric | OpenClaw | Next Best (IronClaw) | Gap |
|--------|----------|----------------------|-----|
| PRs/24h | 500 | 50 | 10x |
| Contributors (implied) | Highest | Moderate | Significant |
| Critical bugs without fix | 9 P1 | 1 P1 (#5456) | 9x |
| Memory trust features | 4 high-engagement issues | 0 | Unique |

**Weakness**: The post-release P1 regression pileup (#98528, #98562, #98565, #98564) is the highest severity bug count of any project. While community scale allows rapid triage, it undermines confidence in the v2026.6.11 release.

---

## 4. Shared Technical Focus Areas

The following requirements are emerging **across multiple projects**, indicating ecosystem-wide priorities:

### Memory & Context Reliability
- **Projects**: OpenClaw, Hermes, NanoBot, CoPaw, PicoClaw, ZeroClaw
- **Specific needs**:
  - **Memory persistence on reset**: OpenClaw (#45608), Hermes (#56341–#56343), NanoBot (#4402) all address session context loss.
  - **Memory poisoning/drift**: OpenClaw (#7707, #20935), CoPaw (#5676—skills not listed) worry about trust and auditability.
  - **Context window management**: IronClaw (#5456—runner lease), NanoBot (#4608—tool result overflow), CoPaw (#5510—tool response capping) all addressing context budget limits.

### Security Hardening
- **Projects**: NanoBot, IronClaw, ZeroClaw, Hermes, OpenClaw
- **Specific needs**:
  - **Shell command bypass**: NanoBot (#4562—`&&` chaining), Hermes (#56353/#56352—abbreviated flags) fixing approval bypasses.
  - **DNS/SSRF attacks**: NanoBot (#4611—DNS rebinding) is a unique concern but echoes ZeroClaw's tool secrets (#8553) and OpenClaw's CVE request (#98550).
  - **Secret management**: OpenClaw (#98549—scoped credentials), ZeroClaw (#8553—env var secrets), IronClaw (#5483—secret injection testing).

### Session & Run Reliability
- **Projects**: OpenClaw, IronClaw, Hermes, NanoBot, CoPaw, ZeroClaw
- **Specific needs**:
  - **Graceful recovery**: OpenClaw (#98522—stuck sessions), IronClaw (#5456—runner lease), CoPaw (#5696—QQ websocket reconnect), ZeroClaw (#8560—hanging tools).
  - **WebSocket/connection stability**: OpenClaw (#98587), CoPaw (#5696), ZeroClaw (#8521) all patching reconnect paths.

### Channel & Platform Expansion
- **Projects**: PicoClaw, NanoClaw, CoPaw, ZeroClaw, OpenClaw
- **Specific needs**:
  - **QQ/WeChat streaming**: PicoClaw (#3201), CoPaw (#5699—typing fix).
  - **Slack/Telegram/Discord reliability**: OpenClaw (#98589), NanoClaw (#2884—Discord), CoPaw (#5699—Telegram).
  - **Matrix E2EE**: NanoClaw (#2844—Rust crypto), ZeroClaw (#8541—thread history).

### WebUI/UX Improvements
- **Projects**: IronClaw, NanoClaw, ZeroClaw, LobsterAI, CoPaw
- **Specific needs**:
  - **Approval workflow**: IronClaw (#5441), OpenClaw (#98523—pending approvals), CoPaw (#5703—toggle persistence).
  - **Debugging visibility**: IronClaw (#5457—empty logs page), ZeroClaw (#8556—secret state indicators).
  - **Localization**: ZeroClaw (#8584—Fluent-based), LobsterAI (#1361—Chinese delete button).

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | IronClaw | NanoBot | CoPaw | ZeroClaw | PicoClaw |
|-----------|----------|----------|---------|-------|----------|----------|
| **Primary Focus** | Memory trust, session continuity, plugin ecosystem | Reborn runtime stability, multi-user turn management | Security, CLI UX, cron automation | v2.0 major refactor, channel integration, LoRA engineering | SOP orchestration, CI hardening, MoA virtual models | Platform expansion (QQ, Android), model fallback |
| **Target Users** | Enterprises/teams needing auditable AI assistants | Large-scale multi-user deployments | Individual developers, automation enthusiasts | Chinese-market users, WeChat/QQ bot operators | DevOps/sysadmin teams, CI pipeline integrations | Hobbyists, multilingual users |
| **Technical Architecture** | Monolithic core with plugin sandbox | One-runtime group harness, runner leases | Lightweight CLI-first, WASM plugins | Channel-focused with pre-release refactor | SOP-first orchestration, gateway-centric | Minimalist, easy self-hosting |
| **Memory Model** | Provenance-tagged, auditable, session-synthesizing | In-memory turn-state authority | Opt-in eager consolidation | ADBPG REST-only memory; reranker incoming | Not a focus—SOP-driven | Simple KV store |
| **Notable Unique Feature** | Memory trust tagging (#7707) | Kubernetes sandbox runtime (#2979) | DNS rebinding protection (#4611) | Loop Engineering (#5665) | MoA virtual provider (#8568) | QQ channel streaming (#3201) |
| **Deployment Model** | Container-native, macOS/Windows | Kubernetes, Docker | CLI binary, single host | Docker, cloud pre-release | Self-hosted, Docker | Lightweight, ARM64 support |
| **Community Language** | International (English-primary) | International | International | Chinese-primary | International | International + Chinese |

**Key Differentiators**:
- **Hermes Agent** is unique for its "salvage" workflow—cherry-picking fixes from older branches into main. This creates a fast path for critical bug closure but may fragment the commit history.
- **NullClaw** focuses on cron scheduling with DB-backed subagent engines—niche but well-executed for automation use cases.
- **LobsterAI** differentiates through artifact management (subagent panels, auto-open preview, agent import/export), targeting creative workflow users rather than developer tools.
- **Moltis** and **TinyClaw** are effectively dormant, providing no competitive signal.

---

## 6. Community Momentum & Maturity

### Tier 1: Rapid Iteration (Shipping code daily, high contributor churn)
- **OpenClaw** (500 PRs/24h) — The ecosystem engine. Mature processes but post-release regression strain.
- **IronClaw** (50 PRs/24h, 30 merged) — Core team driving stability; high CI discipline.
- **CoPaw** (50 PRs/24h, 25 merged) — Pre-release blitz; high risk of regressions but clear momentum.

### Tier 2: Active Development (Shipping weekly, dedicated maintainers)
- **NanoBot** (41 PRs/24h, 9 merged) — Security-first; responsive to bug reports but slower feature velocity.
- **ZeroClaw** (50 PRs/24h, 5 merged) — High open issue count (16) suggests maintenance debt; feature work outpacing bug fixes.
- **Hermes Agent** (50 PRs/24h, 19 merged) — Efficient salvage process; good balance of feature and fix work.
- **NanoClaw** (17 PRs/24h, 10 merged) — Good merge rate but critical setup bug unresolved; risk of user churn.

### Tier 3: Steady Maintenance (Sustained but lower volume)
- **PicoClaw** (5 PRs/24h, 2 merged) — Steady but slow; high-impact bug (#3153) needs attention.
- **NullClaw** (4 PRs/24h, 4 merged) — Clean execution but niche audience.
- **LobsterAI** (26 PRs/24h, 23 merged) — Surprisingly high throughput for a less-visible project; strong backlog clearing.

### Tier 4: Inactive/Dormant
- **Moltis** — Dependabot-only activity.
- **TinyClaw** — Zero activity.

**Maturity Signals**:
- **OpenClaw** and **IronClaw** are the only projects with formal release notes and migration guides, indicating production maturity.
- **CoPaw**'s v2.0.0 pre-release tracker (#5273) with active triage demonstrates disciplined beta management.
- **NanoBot**'s same-day fix for #4615 (crash) and #4434 (security bypass) shows exceptional responsiveness.

---

## 7. Trend Signals

### 1. Trustworthy Memory is the Dominant Concern
Across OpenClaw (4 high-engagement issues), Hermes (session persistence fixes), and CoPaw (skills not listed in system prompt), users are demanding:
- **Provenance tracking** (who wrote what to memory)
- **Tamper detection** (audit logs)
- **Graceful context compression** (no silent data loss)
- **Agent developers should invest in memory trust frameworks**—this will become a top selection criterion for enterprise buyers.

### 2. Post-Release Regression Fatigue
OpenClaw's 4 P1 regressions in one release and CoPaw's multiple v2.0 beta bugs signal that **shipping velocity is outpacing test coverage**. The ecosystem is adopting a "test-in-production" model, relying on community bug reports. IronClaw's investment in T0-ERRPATHS and T0-SECRET-INJECT tests is the counter-signal—**expect automated test suites to become a competitive differentiator**.

### 3. Security is Shifting from Perimeter to Workflow
Projects are moving beyond auth and CORS to **workflow-level security**: approval bypass patterns (abbreviated flags, chained commands), SSRF with DNS rebinding, and memory poisoning. The trend is toward **zero-trust agent security**—assuming the model can be compromised and hardening every data path.

### 4. Multi-Platform is Becoming a Requirement
QQ, WeChat, Telegram, Discord, Slack, Matrix—users expect agents to work across all channels. Projects like PicoClaw and CoPaw (Chinese market) and NanoClaw (Discord/WhatsApp) are racing to support every major platform. **Cross-platform channel adapters (like NanoClaw's Chat SDK bridge) will become table stakes**.

### 5. Cost Optimization Through Context Management
Headroom compression (#5063 in CoPaw), input token reduction (#4581 in NanoBot), and tool response capping (#5510 in CoPaw) all point to a growing **focus on API cost reduction**. With LLM inference still expensive, agent frameworks that minimize token waste will win budget-conscious users. This is a greenfield competitive opportunity—no project has a complete solution yet.

### 6. Multi-Agent Orchestration is Emerging
MoA (Mixture-of-Agents in ZeroClaw), Reborn one-runtime (IronClaw), and subagent artifact panels (LobsterAI) indicate the ecosystem is moving toward **coordinated multi-model, multi-agent workflows**. Single-agent chatbots are becoming commoditized; the next frontier is agent-to-agent collaboration.

### Value for AI Agent Developers
- **Build memory trust features now**—the community has validated demand before major projects ship production solutions.
- **Invest in automated regression testing** before feature expansion—the post-release bug pattern will hurt your credibility.
- **Prioritize security at the workflow/approval level**—perimeter security alone is insufficient for agent deployments.
- **Design for multi-platform from day one**—channel adapters are not afterthoughts.
- **Context cost optimization** is a rapidly growing competitive moat—solve this and you'll capture the cost-sensitive developer market.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest – 2026-07-01

## 1. Today's Overview
NanoBot saw **very high activity** over the past 24 hours: **41 pull requests** were updated (9 merged/closed), and **5 issues** were touched (2 closed). The project continues to mature with a strong focus on **security hardening, reliability fixes, and feature expansion**. Notably, a critical startup crash bug (#4615) was identified and a fix PR (#4617) opened the same day. Several high-priority security issues (MCP policy bypass, API auth parity) were resolved. The community submitted multiple feature requests, including support for OpenAI’s Response API and multiline CLI input.

## 2. Releases
No new releases were published today.

## 3. Project Progress
**9 PRs were merged or closed** in the last 24 hours. Key advancements:

- **Security & Authentication**:  
  - `feat(api): require api_key when binding to all interfaces (parity with WS gateway)` (#4548, closed) – Fixes #4490, bringing API server auth to parity with the WebSocket gateway.  
  - `fix(tools): gate MCP resource and prompt registration behind enabledTools` (#4436, closed) – Addresses #4434, a security bypass where `enabledTools: []` leaked resources and prompts to the model.

- **Refactoring & Code Quality**:  
  - `refactor(tools): use structured tool error results` (#4610, closed) – Introduces `ToolResult` error marker, replacing fragile `"Error"` string matching.  
  - `refactor(webui): derive provider model catalog kind` (#4613, closed) – Moves classification logic to `ProviderSpec` metadata.

- **CLI Usability**:  
  - `feat(cli): support multiline input via Shift+Enter / Alt+Enter` (#4614, open) – Adds multiline composition support to the interactive prompt.

- **Cron Stability**:  
  - `fix(cron): tolerate unsupported directory fsync` (#4617, open) – Fixes the startup crash reported in #4615 by ignoring `EINVAL` on directory `fsync()`.

## 4. Community Hot Topics
The most engaging issues/PRs (by comments/reactions) this week:

- **Issue #4615** ([link](https://github.com/HKUDS/nanobot/issues/4615)) – **Gateway startup crash** due to `CronService` calling `fsync()` on a parent directory. 2 comments. Critical, with a fix PR (#4617) opened the same day.

- **Issue #4434** ([link](https://github.com/HKUDS/nanobot/issues/4434)) – **MCP `enabledTools` deny-all bypass** exposing resources and prompts. Now closed with fix PR #4436 merged. Reflected strong community concern over security boundaries.

- **Issue #4611** ([link](https://github.com/HKUDS/nanobot/issues/4611)) – **DNS rebinding TOCTOU in SSRF validation** (`validate_url_target` does not pin resolved IP). 1 thumbs-up, no comments yet. Represents a high-severity security bug awaiting maintainer response.

## 5. Bugs & Stability
Bug reports and crashes are ranked by severity:

| Severity | Issue/PR | Description | Status |
|----------|----------|-------------|--------|
| **P0** | #4434 (closed) | MCP deny-all bypass – resources & prompts leaked to model | Fixed in #4436 |
| **P1** | #4615 | Gateway crash on startup when `CronService` does directory `fsync()` | Fix PR #4617 |
| **P1** | #4611 | DNS rebinding TOCTOU in SSRF validation | No fix yet |
| **P1** | #4521 (reported) / PR #4562 | Shell allowlist bypass via chained commands (`&&`) | Fix PR open (#4562) |
| **P1** | PR #4608 | Agent context overflow from multiple tool results in same turn | Fix PR open (#4608) |
| **P1** | PR #4545 | Windows command execution divergence (`cmd.exe` vs PowerShell) | Fix PR open (#4545) |

**Missing fix**: #4611 (DNS rebinding) has a clear problem description but no maintainer reaction or PR as of today. Should be prioritized.

## 6. Feature Requests & Roadmap Signals
Several user-driven features may shape the next release:

- **OpenAI Response API support** (#4612, open) – Request to connect to GPT models using the native `/v1/responses` endpoint. Presence of a contributor suggests it could land soon.
- **Multiline CLI input** (#4614, open) – Already implemented in a PR; likely to be merged quickly.
- **Session-bound local triggers** (#4591, open) – Adds `/trigger <name>` and CLI command `nanobot trigger`. Enhances workflow automation.
- **Dollar skill shortcuts** (#4284, open) – WebUI shortcut `$<skill>` for skill invocation. Long-standing; perhaps stalled on review.
- **Performance optimization (context usage)** (#4581, open) – Reduces input tokens per turn to lower costs. High community interest.
- **Heartbeat trigger command** (#4437, open) – Phase 1 decision-making and execution for cron-like heartbeat triggers.
- **Eager memory consolidation** (#4402, open) – Opt-in archiving of conversation slices to `memory/history.jsonl`.

**Prediction**: The next minor version will likely include the OpenAI Response API, multiline CLI, and basic heartbeat/trigger support.

## 7. User Feedback Summary
Real pain points expressed by the community:

- **Frustration**: “I can only use OPENAI response api way (but not compatible way) to connect to chat gpt model” – Issue #4612. Lack of native endpoint support forces users to use workarounds.
- **Stability**: “`nanobot gateway` crashes during startup” – Issue #4615. Users on filesystems that don’t support directory `fsync` (e.g., vboxsf) are blocked.
- **Security concerns**: “No authentication option on API server” – #4490 (now fixed). Users were wary of exposing agent to the network. Also, the MCP bypass (#4434) was alarming but quickly resolved.
- **Feature gaps**: “No way to compose multi-line message before submitting” – PR #4614. The single-line buffer was a known UX limitation.
- **Context management**: “Accumulated tool results exceed context window budget” – PR #4608. Users hitting context limits during multi-tool turns.

Overall, satisfaction with the project’s **rapid response to security and stability bugs** (same-day fixes for #4436, #4548, #4617) contrasts with some **delays on feature-rich PRs** like #4284 (dollar shortcuts, open since June 10).

## 8. Backlog Watch
Older, unanswered issues/PRs that need maintainer attention:

- **PR #4284** ([link](https://github.com/HKUDS/nanobot/pull/4284)) – feat(webui): add dollar skill shortcuts. Open since June 10, no reviewer comments. A small UX improvement that has gathered no traction.
- **PR #4402** ([link](https://github.com/HKUDS/nanobot/pull/4402)) – feat(memory): add opt-in eager consolidation. Open since June 18. May need feedback on configurability.
- **PR #4373** ([link](https://github.com/HKUDS/nanobot/pull/4373)) – fix(memory): preserve delivery context during consolidation. Open since June 16. Could conflict with #4402.
- **PR #4416** ([link](https://github.com/HKUDS/nanobot/pull/4416)) – feat(cron): support job model presets. Open since June 19. Addresses #4378; likely waiting for review.
- **PR #4437** ([link](https://github.com/HKUDS/nanobot/pull/4437)) – add heartbeat trigger command. Open since June 21. Large feature; no maintainer interaction.
- **Issue #4611** ([link](https://github.com/HKUDS/nanobot/issues/4611)) – Security: DNS rebinding TOCTOU in SSRF validation. No comments from maintainers since opened on June 30. Needs triage.

These items represent potential bottlenecks in community contribution onboarding. Prioritizing review of the older PRs would improve contributor retention and project velocity.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest – 2026-07-01

## 1. Today's Overview

Hermes Agent saw extremely high activity in the past 24 hours, with **50 pull requests updated** (31 open, 19 merged/closed) and **3 issues updated** (1 open, 2 closed). A large batch of "salvage" PRs—cherry-picked fixes from earlier branches—dominated the merge queue, indicating a focused effort to consolidate critical bug fixes and security patches into the main branch. No new releases were cut, but the project tackled several session-state corruption bugs, gateway‑specific regressions, and approval‑gate bypasses. Overall project health appears strong, with maintainers actively responding to community contributions and driving quality improvements.

## 2. Releases

No new releases are available for this digest period. The latest publicly released version remains **hermes-agent@0.16.0** (referenced in issue #48765).

## 3. Project Progress

**Merged/closed PRs (selected notable items):**

- **#50405** – `fix(acp): stop _persist from deleting compression-archived history` – A critical fix for ACP session persistence that was deleting pre‑compaction transcripts.  
- **#56334** – `fix(auxiliary_client): dedup resolve_provider_client fall-through warnings` – Reduces log noise on misconfigured providers.  
- **#56335** – `fix(macos): retry launchd reload until registered on transient bootstrap failure` – Prevents macOS launchd outages under load.  

**Merged/closed issues:**

- **#48765** – `[Bug]: Hermes Agent repeats completed shell tool results` – Closed after reproducing and applying a fix.  
- **#56323** – `[Bug]: Telegram compact /reasoning level commands are rejected in groups` – Closed with a fix in the Telegram adapter.

**Key feature advancement:** PR **#54535** (Slack read‑only history tool) and PR **#55614** (mem0 self‑hosted dashboard support) remain open but received updates, suggesting ongoing work on platform integration and memory management.

## 4. Community Hot Topics

Most discussions remain inside pull request review threads. The **highest‑volume activity** revolves around a series of **salvage PRs authored by @kshitijk4poor** that cherry‑pick and rework fixes from earlier contributions (e.g., #56352, #56353, #56340–#56344). These PRs together address session persistence, gateway recovery, and security hardening – reflecting a community‑wide need for **reliability and security in production deployments**.

Notable open PR with cross‑cutting impact:  
**#33505** – `ISSUE-003: Add Origin header validation to WebSocket handle_ws()` – a security enhancement that has been open since May 27 without a merge decision.

## 5. Bugs & Stability

**New bugs reported today (2026-07-01):**

- **#56337** (OPEN, P2) – Telegram strips `/command@BotName args` into an unknown command. An immediate fix PR **#56338** has been submitted.  
- **#56323** (CLOSED, P2) – Telegram compact `/reasoningmedium` commands rejected in groups (fix already merged).  

**Stability fixes merged or submitted today (by severity):**

| Severity | Issue / PR | Description |
|----------|------------|-------------|
| P1 | **#56340** (PR, open) | Gateway/Telegram sessions stall on truncated responses – recovery attempt added. |
| P1 | **#56345** (PR, open) | Matrix gateway inbound messages not dispatched – `wait_sync=True` fix. |
| P1 | **#56341** (PR, open) | CLI `/resume` and `/branch` lose un‑persisted messages – flush fix. |
| P2 | **#56351** (PR, open) | Same class of bug as #56341, applied to `/branch` handlers. |
| P2 | **#56342** (PR, open) | ACP `_persist` destructive replacement of compressed history – now preserves. |
| P2 | **#56343** (PR, open) | Codex sessions not persisted to session DB – FTS/search now sees them. |
| P2 | **#56353** (PR, open) | Windows destructive shell commands not detected by approval gate. |
| P2 | **#56352** (PR, open) | Abbreviated flag bypasses in `git`/`sudo` approval patterns. |
| P2 | **#56344** (PR, open) | MoA reference fan‑out ignores user interrupt – now checks on each poll cycle. |

One regressed bug (#48765, duplicate tool results) was closed today with a reproducer already provided.

## 6. Feature Requests & Roadmap Signals

**Feature PRs in progress:**

- **#54535** – *Slack read‑only channel history tool* – Allows the agent to inspect recent Slack channel history during a session. Expected to land in next minor release.  
- **#55614** – *mem0 self‑hosted dashboard support* – Third connection mode for the memory plugin, enabling local dashboarding without Mem0 cloud.  
- **#56333** – *Claude Code history scanner (M1 of sidebar feature)* – Desktop app script to read Claude Code session metadata. Signals a long‑term plan to integrate with Claude Code history.  
- **#56336** – *Fix MoA profile switching by moving into TOOL_CATEGORIES* – Addresses a configuration gap that prevented Mixture‑of‑Agents from working with multi‑env setups.

**Likely roadmap inclusion:** The slew of session persistence fixes (#56341–#56343) suggests that **context‑loss prevention and session reliability** are a top priority for the next release.

## 7. User Feedback Summary

Real pain points surfaced in today’s issues:

- **Telegram group users** report frustration with slash‑command parsing – both compact commands (`/reasoningmedium`) and bot‑suffix stripping (`/command@BotName args`) cause the agent to misinterpret or reject valid inputs. Fixes are already in flight.  
- **Deterministic test providers** (via #48765) revealed a subtle bug where Hermes could duplicate tool results – caught by a community member with a standalone reproducer.  
- **Session data loss** during `/resume` and `/branch` operations was a recurring source of dissatisfaction; the three salvage fixes aim to close that gap permanently.

Overall sentiment appears constructive, with contributors providing clear reproduction steps and maintainers responding rapidly with patches.

## 8. Backlog Watch

| Item | Created | Status | Notes |
|------|---------|--------|-------|
| **#33505** | 2026-05-27 | OPEN | `Origin header validation for WebSocket` – security‑boundary PR, no maintainer comment in over a month. Risk of CORS‑bypass attacks if left unaddressed. |
| **#50405** | 2026-06-21 | **CLOSED today** | ACP persist bug – was open for 10 days before being salvaged. Demonstrates that salvage process is working, but initial delay suggests review bottleneck. |
| **#54535** | 2026-06-29 | OPEN | Slack history tool – no maintainer review yet. Feature may be waiting for capacity. |

**Recommendation:** The community would benefit from a decision on #33505, either merging with modifications or clarifying requirements, to avoid lingering security exposure.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest – 2026-07-01

## 1. Today's Overview

Activity remains moderate with **4 issues** and **5 pull requests** updated in the last 24 hours. One new **nightly release** (v0.3.1-nightly) was published, bringing early access to changes on the `main` branch. The project merged **two bug-fix PRs** (better auth error messages and safer type assertions) and saw the opening of a significant feature PR introducing a configurable model fallback chain. A fresh feature request for QQ channel streaming indicates ongoing expansion of platform support. Overall, the project is steadily addressing both user-reported bugs and community-driven enhancements.

## 2. Releases

| Version | Type | Notes |
|---------|------|-------|
| `v0.3.1-nightly.20260701.2cf030d2` | **Nightly** | Automated build from `main` – **unstable**. Use for testing only. |

**Changes**: This nightly includes all code merged up to this date, notably the `fix(providers)` and `fix(registry)` PRs (see below).  
**Migration notes**: No stable release; no breaking changes expected. Users on v0.3.1 stable should wait for the next stable tag.

## 3. Project Progress

Two pull requests were merged/closed today, both addressing stability and user experience:

- **PR #3198 – `fix(providers): surface friendly auth error messages`**  
  Closed/Merged. Improves error handling when API keys, tokens, or provider permissions fail. Users now receive clearer, actionable feedback instead of generic errors.  
  [GitHub](https://github.com/sipeed/picoclaw/pull/3198)

- **PR #3131 – `fix(registry): add ok checks for tool schema type assertions`**  
  Closed/Merged. Adds defensive checks for three type assertions in `pkg/tools/registry.go`, preventing potential panics when tool schemas contain unexpected value types.  
  [GitHub](https://github.com/sipeed/picoclaw/pull/3131)

Both fixes strengthen core tool-calling and provider-connection reliability.

## 4. Community Hot Topics

The most active discussion centres on a tool-call leak bug:

- **Issue #3153 – `[BUG] Volcengine Doubao Seed tool calls occasionally leak as <seed:tool_call> text`**  
  **2 comments** (most-discussed recent issue). Users report that with `doubao-seed-2.0-pro`, tool calls sometimes appear as raw XML instead of being executed. This is a functional bug affecting a major model provider. No fix PR has been linked yet.  
  [GitHub](https://github.com/sipeed/picoclaw/issues/3153)

- **Issue #3159 – `[BUG] 经常重复任务` (Repeated tasks)**  
  1 comment. User reports that when querying different news topics (e.g., US news then French news), the AI re-executes the first task before answering the second. Indicates a workflow or history-merging issue.  
  [GitHub](https://github.com/sipeed/picoclaw/issues/3159)

- **PR #3200 – `feat(models): add configurable default fallback chain`** (new, open)  
  A highly anticipated feature for setting default model + fallback ordering from the web UI. This PR gathered immediate attention, though it has no comments yet.  
  [GitHub](https://github.com/sipeed/picoclaw/pull/3200)

## 5. Bugs & Stability

| Issue | Severity | Description | Status |
|-------|----------|-------------|--------|
| #3153 | **High** | Volcengine Doubao tool-call leak – raw `<seed:tool_call>` text shown to user. Functional bug for a widely used model. | Open |
| #3159 | **Medium** | Repeated task execution when multiple queries are issued in succession. User experience degradation. | Open |
| #3199 | **High** | Custom model provider cannot connect to `http://127.0.0.1` OpenAI-compatible endpoint. **Closed**; the issue was reported and resolved (likely fixed without a visible PR). | Closed |

Additionally, the two merged PRs (#3198, #3131) address stability improvements for auth error reporting and type safety. No regressions are reported in the current nightly.

## 6. Feature Requests & Roadmap Signals

- **Feature Request #3201 – `Support streaming output for QQ channel`**  
  New, opened today. Users want token-by-token streaming for QQ, similar to Telegram and Pico WebSocket. Likely to be picked up soon as it aligns with the project’s goal of parity across chat gateways.  
  [GitHub](https://github.com/sipeed/picoclaw/issues/3201)

- **PR #3200 – Configurable default fallback chain**  
  Open. Adds a dedicated UI and backend API for setting the default model and reordering fallback models. This is a strong candidate for the next stable release.

- **PR #3157 – `feat: add Android ADB remote operations tool`** (open since June 22, stale label)  
  Experimental tool for controlling Android devices via ADB. If merged, it would expand PicoClaw’s tool ecosystem significantly.

- **PR #3063 – `feat: add deltachat gateway`** (open since June 8)  
  Adds DeltaChat as a new chat channel – another signal of broadening platform support.

**Prediction for next stable version**: The fallback chain (#3200) and QQ streaming (#3201) are the most likely additions, along with the two merged bug fixes.

## 7. User Feedback Summary

**Pain points**:
- Tool call leaking (exposed raw XML) – affects reliability and user trust.  
- Repeated task execution – frustrating for users who need sequential, non-redundant responses.  
- Unclear authentication errors – mitigated by PR #3198.  

**Use cases**:
- Multilingual support (Chinese issue #3159) indicates growing non-English user base.  
- Local model endpoints (custom providers) are actively tested; bug #3199 shows demand for local LLM hosting.  

**Satisfaction indicators**: The community continues to submit feature requests (QQ streaming, fallback chain) and reports bugs constructively. The project’s responsiveness – merging two fixes in 24 hours – is a positive signal.

## 8. Backlog Watch

The following long-standing items lack maintainer response or resolution:

| Item | Created | Days Open | Notes |
|------|---------|-----------|-------|
| **PR #3063** – DeltaChat gateway | 2026-06-08 | 23 days | No maintainer review; last update June 30 (stale label). Could benefit from feedback. |
| **PR #3157** – Android ADB tool | 2026-06-22 | 9 days | Similar situation – stale label, no comments from maintainers. |
| **Issue #3153** – Tool-call leak | 2026-06-22 | 9 days | High-impact bug with no linked fix PR. Needs prioritisation. |
| **Issue #3159** – Repeated tasks | 2026-06-23 | 8 days | Medium impact; anecdotal but no diagnostic steps provided by maintainer. |

These items share the `stale` label or lack recent official replies. The maintainer team may need to allocate time for review or request additional information to move them forward.

---

*Generated on 2026-07-01 from GitHub data (github.com/sipeed/picoclaw).*

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-07-01

## 1. Today's Overview
July 1 saw heavy development activity with **17 pull requests** updated (7 open, 10 merged/closed) and **8 issues** filed (7 open, 1 closed). No new releases were published. The day’s work was dominated by fixing critical infrastructure bugs (webhook bind crash, silent message dropping, WhatsApp media recovery) and merging several long-running feature PRs, including a Discord channel adapter, document rendering in ephemeral containers, and a daily‑news agent with WeChat support. Despite the high throughput, a cluster of fresh integration bugs—especially a broken OneCLI default setup (Issue #2903)—suggests the out‑of‑box experience remains fragile and may require maintenance attention before the next milestone.

## 2. Releases
*No new releases in the last 24 hours.*

## 3. Project Progress
**Merged/closed PRs today (10)**:
- **Security** – [PR #2880](https://github.com/nanocoai/nanoclaw/pull/2880) (fixes #2828) closes a symlink‑escape vulnerability (CWE‑59) on inbound file writes; a critical hardening for all multi‑agent deployments.
- **Channel adapters** – [PR #2884](https://github.com/nanocoai/nanoclaw/pull/2884) adds a full Discord adapter via the Chat SDK bridge, fixing approval‑button routing in DMs. [PR #2889](https://github.com/nanocoai/nanoclaw/pull/2889) merges a WeChat channel adapter plus a daily‑news agent with 33 test cases.
- **Media/attachment fixes** – [PR #2895](https://github.com/nanocoai/nanoclaw/pull/2895) recovers inbound WhatsApp media (images, video, audio) by passing `reuploadRequest` context; [PR #2896](https://github.com/nanocoai/nanoclaw/pull/2896) fixes a regression on approval‑answer messages.
- **Infrastructure** – [PR #2893](https://github.com/nanocoai/nanoclaw/pull/2893) adds a `render_document` MCP tool using ephemeral network‑isolated containers (Quarto/LaTeX/Chromium). [PR #2891](https://github.com/nanocoai/nanoclaw/pull/2891) adds `resolveChannelName` to the adapter interface, unblocking Slack and Telegram builds.
- **Setup/operations** – [PR #2885](https://github.com/nanocoai/nanoclaw/pull/2885) offers Slack Socket Mode in the guided setup flow. [PR #2874](https://github.com/nanocoai/nanoclaw/pull/2874) makes the Signal adapter survive boot flaps instead of crash‑looping. [PR #2875](https://github.com/nanocoai/nanoclaw/pull/2875) adds a Coolify deployment skill. [PR #2018](https://github.com/nanocoai/nanoclaw/pull/2018) fixes DM‑context approval routing for Discord clicker identification.

**Open PRs with notable progress**:
- [PR #2890](https://github.com/nanocoai/nanoclaw/pull/2890) – agent template loader and setup flow (updated today).
- [PR #2317](https://github.com/nanocoai/nanoclaw/pull/2317) – free voice transcription skill (updated after a long pause).
- [PR #2844](https://github.com/nanocoai/nanoclaw/pull/2844) – native Matrix E2EE adapter (persistent review).

## 4. Community Hot Topics
- **Security fix for symlink escape** – Issue [#2828](https://github.com/nanocoai/nanoclaw/issues/2828) (closed, 2 👍) was the only issue with reactions. The prompt‑injected agent attack vector drew attention; PR #2880 was merged swiftly. This signals strong community concern for sandbox integrity.
- **Discord approval‑button fix** – PR [#2884](https://github.com/nanocoai/nanoclaw/pull/2884) (merged) and its follow‑up PR [#2899](https://github.com/nanocoai/nanoclaw/pull/2899) (open) address a frustrating UX bug where every DM button tap led to rejection. The root cause was a newline delimiter in `custom_id`, affecting all Discord Gateway users.
- **WhatsApp media reliability** – Issue [#2894](https://github.com/nanocoai/nanoclaw/issues/2894) and merged PR [#2895](https://github.com/nanocoai/nanoclaw/pull/2895) address silent media drops. The community (echarrod) surfaced a concrete failure mode (CDN fetch failures) which is now recovered.

## 5. Bugs & Stability
| Severity | Issue | Description | Fix PR exists? |
|----------|-------|-------------|----------------|
| **Critical** | [#2903](https://github.com/nanocoai/nanoclaw/issues/2903) | OneCLI default setup binds gateway to `127.0.0.1` while clients target Docker bridge `10.0.0.1` → agents never respond out of the box. | No |
| **High** | [#2902](https://github.com/nanocoai/nanoclaw/issues/2902) | Messages accepted on a channel (e.g. Telegram) are silently swallowed if the agent container fails to spawn; failure only logged, user never notified. | No |
| **High** | [#2900](https://github.com/nanocoai/nanoclaw/issues/2900) | Webhook server `EADDRINUSE` crash kills the entire host daemon, even though webhooks are optional. | No |
| **Medium** | [#2901](https://github.com/nanocoai/nanoclaw/issues/2901) | `WEBHOOK_PORT` set in `.env` is silently ignored (only real process env works). | No |
| **Medium** | [#2894](https://github.com/nanocoai/nanoclaw/issues/2894) | WhatsApp inbound media silently dropped when direct CDN fetch fails. | Yes (PR #2895 merged) |
| **Low** | [#2898](https://github.com/nanocoai/nanoclaw/issues/2898), [#2897](https://github.com/nanocoai/nanoclaw/issues/2897) | Smoke test issues from bot; safe to close. | N/A |

**Outcome**: Three critical/high‑severity bugs (OneCLI, silent swallowing, webhook crash) were reported today without any committed fix. The WhatsApp media fix is already merged. The webhook crash bug (#2900) is particularly concerning because it forces the host into a crash‑loop for optional infrastructure.

## 6. Feature Requests & Roadmap Signals
- **Agent templates** – PR [#2890](https://github.com/nanocoai/nanoclaw/pull/2890) proposes a reusable template system (folder, instructions, MCP tools, skills) loadable from a public library, local path, or git repo. This could land in the next minor release.
- **Voice transcription** – PR [#2317](https://github.com/nanocoai/nanoclaw/pull/2317) (open since May 7) adds a free Whisper‑based transcription skill with two backends (Python/GPU, whisper.cpp/CPU). Updated today, suggesting renewed interest.
- **Native Matrix E2EE adapter** – PR [#2844](https://github.com/nanocoai/nanoclaw/pull/2844) replaces the Chat SDK bridge with a Rust‑based crypto binding for persistent end‑to‑end encryption. High value for privacy‑conscious deployments.
- **Telegram thread support** – PR [#2892](https://github.com/nanocoai/nanoclaw/pull/2892) simply flips `supportsThreads: true`. Likely to merge quickly.
- **Container ergonomics** – PR [#2771](https://github.com/nanocoai/nanoclaw/pull/2771) adds configurable `--shm-size` (default 1g) and `--init` for agent containers to support headless Chromium. Essential for browser‑based agents.

**Prediction**: Agent templates, Telegram threads, and container shm‑size are strong candidates for the next version. The Matrix adapter and voice transcription may slip if review bandwidth is tight.

## 7. User Feedback Summary
- **Pain points**:
  - *Out‑of‑box unusability* (Issue #2903) – fresh installs cannot reach the gateway.
  - *Silent failures* (Issue #2902) – users send messages that vanish.
  - *Configuration confusion* (Issue #2901) – `.env` is the documented config location but `WEBHOOK_PORT` is ignored.
  - *Crash on optional services* (Issue #2900) – webhook bind failures shouldn’t kill the daemon.
- **Use cases**:
  - Daily news aggregation via agents (PR #2889) shows interest in agent workflows with scheduling and multi‑channel output.
  - Document rendering (PR #2893) indicates demand for agent‑assisted Quarto/LaTeX report generation.
  - WhatsApp media reliability (Issue #2894) highlights real‑world media sharing expectations.

Overall satisfaction appears moderate—the project is fixing bugs aggressively (10 PRs merged), but the volume of fresh critical bugs in a single day suggests the current release is not yet stable for production.

## 8. Backlog Watch
| Item | Age | Status | Notes |
|------|-----|--------|-------|
| **PR #2317** – Voice transcription skill | Open since **2026‑05‑07** (55 days) | Updated today (July 1) | Needs maintainer review; could be a major UX win for voice channels. |
| **PR #2771** – Configurable `--shm-size` + `--init` | Open since **2026‑06‑15** (16 days) | Updated today | Low risk, high impact for browser agents; should be fast‑tracked. |
| **PR #2844** – Native Matrix E2EE adapter | Open since **2026‑06‑24** (7 days) | Updated June 30 | Requires careful crypto review; no signs of activity today. |

All other open issues are very fresh (< 24h). No long‑abandoned issues are visible in the current window. The security issue #2828 has been closed, so that backlog item is resolved.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

Here is the NullClaw project digest for 2026-07-01, based on the provided GitHub data.

---

## NullClaw Project Digest — 2026-07-01

### 1. Today's Overview

Project activity is moderate, with no new releases today. The main action is on the Pull Request front, where four long-standing PRs have been closed or merged after a period of inactivity. The single open issue, a build failure on Android/Termux, has received recent attention but remains unresolved. Overall, the project shows a healthy cleanup of pending features and fixes, though responsiveness to user-reported platform-specific bugs could be improved.

### 2. Releases

No new releases were created or updated in the last 24 hours.

### 3. Project Progress

Four Pull Requests were merged or closed today, representing significant progress on both cron infrastructure and provider compatibility:

- **[PR #783](https://github.com/nullclaw/nullclaw/pull/783) (feat)**: A major update to the cron system, introducing a DB-backed subagent engine with run history, atomic scheduling, per-job timezone offsets, JSON CLI output, and security hardening. This represents a substantial backend improvement for scheduled tasks.
- **[PR #641](https://github.com/nullclaw/nullclaw/pull/641) (fix)**: Fixes a critical issue with GLM/ZhipuAI providers where thinking mode was always-on, causing response loops. Also corrects native tool_calls behavior for this provider.
- **[PR #645](https://github.com/nullclaw/nullclaw/pull/645) (fix)**: Adds a missing `--account` flag to the `cron add-agent` CLI command, enabling users to specify delivery routing (e.g., which Telegram bot) directly from the command line.
- **[PR #643](https://github.com/nullclaw/nullclaw/pull/643) (fix)**: Makes the `command` field optional for agent jobs in `cron.json`, fixing a bug where agent cron jobs were silently deleted on gateway restart.

### 4. Community Hot Topics

The most active discussion remains centered on platform compatibility:

- **[Issue #868](https://github.com/nullclaw/nullclaw/issues/868) (6 comments)**: A user reports that `zig build` fails on aarch64 Android (Termux) due to an `AccessDenied` error related to `linkat` during the linking phase. The issue has generated interest from other users potentially facing the same barrier. The underlying need is for reliable cross-platform build support, particularly for mobile/ARM64 environments which are becoming more common for running local AI agents.

No other issue or PR has garnered significant community engagement in this window.

### 5. Bugs & Stability

One active bug is being tracked:

- **[Issue #868](https://github.com/nullclaw/nullclaw/issues/868) — Severity: Medium/High**: The build process fails entirely on Android/Termux (aarch64) due to a filesystem permission error (`linkat`). This is a platform-specific blocker for users trying to compile and run NullClaw on recent Android devices. The error appears related to the Zig linker, not the application code itself. No fix PR is currently linked.

### 6. Feature Requests & Roadmap Signals

No new feature requests were filed today. However, the merged **[PR #783](https://github.com/nullclaw/nullclaw/pull/783)** (cron subagent engine) signals a strong roadmap focus on robust, production-grade scheduling. The addition of run history and JSON output suggests a push toward making NullClaw more suitable for automated operations and observability. Given the volume of recent cron-related fixes ([#643](https://github.com/nullclaw/nullclaw/pull/643), [#645](https://github.com/nullclaw/nullclaw/pull/645)), we can expect continued refinement of the cron system in the next release, likely with better documentation and CLI defaults.

### 7. User Feedback Summary

Available feedback is limited but telling:

- **Pain Point (Platform Support)**: The reporter of [#868](https://github.com/nullclaw/nullclaw/pull/868) is an active power user trying to run NullClaw in a non-standard environment (Termux on Android). The specific error (linkat access) indicates friction at the system call level, likely a Zig toolchain limitation on that platform rather than a code bug. This user is frustrated that a standard build command fails without a clear workaround.
- **Use Case (Scheduling & Delivery)**: The merged cron PRs ([#645](https://github.com/nullclaw/nullclaw/pull/645), [#783](https://github.com/nullclaw/nullclaw/pull/783)) address clear user demand for configuring agent jobs and routing their output to specific delivery accounts (e.g., Telegram bots) entirely from the CLI, without manual JSON editing. This suggests users are deploying NullClaw as a background automation service.
- **Satisfaction**: The merging of PR #641 will likely satisfy GLM/ZhipuAI users who experienced looping behavior due to forced thinking mode.

### 8. Backlog Watch

The following item requires maintainer attention:

- **[Issue #868](https://github.com/nullclaw/nullclaw/issues/868) — Created 2026-04-23, Unresolved**: This build failure has been open for over two months with no official response or fix. While the environment is niche (LineageOS/Termux), the fact that it’s the only active issue means it represents a 100% open bug rate for recent user reports. A maintainer acknowledgment or suggested workaround (e.g., using a different Zig version or a static build) would improve project health perception even without an immediate code fix.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

## IronClaw Project Digest — 2026-07-01

### 1. Today's Overview
The project had a very active day with **50 PRs updated** (30 merged or closed) and **9 issues updated** (7 still open). The pace of development remains high, with the core team focusing on **Reborn runtime stability**, **WebUI v2 usability fixes**, and **CI infrastructure improvements**. A major recurring theme is the **runner lease expiration problem** (#5456, P1), which is being addressed by a combination of in-memory turn-state authority (#5486), compact recovery snapshots (#5493), and better lease logging (#5494). Multiple WebUI polish items (composer clearing, duplicate header, skill activation noise) were closed today, reflecting an ongoing effort to improve the user experience.

### 2. Releases
**No new releases** were published today. The last release activity is tracked in PR #5311 (still open), which proposes version bumps for `ironclaw_common`, `ironclaw_safety`, `ironclaw_skills`, and `ironclaw` itself.

### 3. Project Progress
**30 PRs were merged or closed today**, including several significant items:

- **WebUI v2 fixes:**
  - Hide skill activation system messages from chat transcript (#5489, [PR](https://github.com/nearai/ironclaw/pull/5489))
  - Remove duplicate Logs page header (#5491, [PR](https://github.com/nearai/ironclaw/pull/5491))
  - Chat composer now clears immediately on send (#5404, [PR](https://github.com/nearai/ironclaw/pull/5404))
  - Fix Slack connect canary to use real WebUI v2 surface (#5485, [PR](https://github.com/nearai/ironclaw/pull/5485))

- **Reborn runtime stability:**
  - **In-memory turn-state authority** (#5486, [PR](https://github.com/nearai/ironclaw/pull/5486)) — addresses the “runtime wedge” caused by CAS contention on per-user `/turns/state.json`
  - Enable `inmemory-turn-state` feature in the deploy build (#5492, [PR](https://github.com/nearai/ironclaw/pull/5492))
  - Add inner latency spans for the agent loop executor (#5487, [PR](https://github.com/nearai/ironclaw/pull/5487))
  - Label turn state filesystem write traces (#5490, [PR](https://github.com/nearai/ironclaw/pull/5490))

- **Testing and CI:**
  - Unblock main-only checks by removing generated WebUI bundle from source control (#5448, [PR](https://github.com/nearai/ironclaw/pull/5448))
  - Error/deny-path coverage for HTTP/shell/MCP tools (T0-ERRPATHS) (#5484, [PR](https://github.com/nearai/ironclaw/pull/5484))
  - Prove credential injection reaches the wire (T0-SECRET-INJECT) (#5483, [PR](https://github.com/nearai/ironclaw/pull/5483))

- **Large feature merge:**
  - Kubernetes sandbox runtime (PR #2979, [PR](https://github.com/nearai/ironclaw/pull/2979)) was closed, providing an alternative to Docker for job execution.

### 4. Community Hot Topics
Activity is primarily driven by the core team and QA testers. The most notable issues and PRs:

- **#5456** ([Issue](https://github.com/nearai/ironclaw/issues/5456)) — Runner lease expiration P1 bug. This is the single most commented issue (1 comment) and is the dominant failure pattern from 6/30 testing.
- **#5476** ([Issue](https://github.com/nearai/ironclaw/issues/5476)) — QA report of Reborn failures under turn-state CAS contention + model latency. Directly related to #5456.
- **#5479** ([Issue](https://github.com/nearai/ironclaw/issues/5479)) — Blocking E-MULTIUSER/C-MULTIUSER due to driver_unavailable error with one-runtime group harness.
- **#5279** ([PR](https://github.com/nearai/ironclaw/pull/5279)) — Fix Reborn queued message steering (open, XL size, high activity).
- **#5441** ([PR](https://github.com/nearai/ironclaw/pull/5441)) — Header notifications for automation approvals (open, XL size, human-verified).

The underlying need is clear: **multi-user/multi-thread turn management** and **debugging visibility** are the top community concerns.

### 5. Bugs & Stability
**High severity:**
- **#5456** ([Issue](https://github.com/nearai/ironclaw/issues/5456)) — Runner lease expiration (P1). Fix PRs: #5486 (in-memory turn-state), #5493 (compact recovery snapshot), #5494 (lease logging).
- **#5476** ([Issue](https://github.com/nearai/ironclaw/issues/5476)) — Reborn runs fail under turn-state CAS contention + model latency. Fixes in flight with #5486.
- **#5479** ([Issue](https://github.com/nearai/ironclaw/issues/5479)) — Second thread with distinct actor fails in one-runtime group harness. Blocks E-MULTIUSER/C-MULTIUSER.

**Medium severity:**
- **#5457** ([Issue](https://github.com/nearai/ironclaw/issues/5457)) — Logs page remains empty and never loads (blocks debugging). No fix PR yet.
- **#5458** ([Issue](https://github.com/nearai/ironclaw/issues/5458)) — Double header on Logs page (P3). Fix merged in #5491.

**Low severity:**
- **#4108** ([Issue](https://github.com/nearai/ironclaw/issues/4108)) — Nightly E2E still failing (since May 27). No resolution visible.

Also merged today: #5486 (in-memory turn-state), #5492 (enable feature), #5490 (latency traces), #5493 (recovery snapshot). These directly target the top bugs.

### 6. Feature Requests & Roadmap Signals
Today’s activity highlights several roadmap directions:

- **Reborn one-runtime architecture** — The harness issue (#5479) and PRs #5486/#5493/#5494 show the team is pushing toward a unified runtime model.
- **WebUI approval UX** — PR #5441 (header notifications for automation approvals, still open) is a high-priority feature.
- **Testing coverage** — The new T0-ERRPATHS and T0-SECRET-INJECT tests (#5484, #5483) indicate a systematic push to increase integration test robustness.
- **Kubernetes sandbox** — The merge of #2979 (closed today) is a major new feature for deployment flexibility.

**Prediction for next release:** Likely includes the Reborn runtime wedge fix, improved logging for lease failures, WebUI v2 polish (composer, header), and the Kubernetes sandbox.

### 7. User Feedback Summary
Pain points from QA testers (joe-rlo, thisisjoshford, henrypark133) and automated reports:

- **Runner lease expiration** is the #1 failure, causing routine runs to fail consistently. Developers are frustrated by the 90-second inactivity threshold.
- **Logs page unavailable** (#5457) blocks debugging of failed runs, compounding the pain.
- **Multi-user scenarios** are broken (#5479), blocking key collaboration features.
- **Chat experience** had minor annoyances (composer delay, skill activation noise) which were quickly fixed today.
- **Approval workflow** improvements are still needed (PR #5441, #5247); users want easier access to global auto-approve settings.

Overall satisfaction seems moderate—the team is reactive to bug reports and merging fixes rapidly, but the stability of new Reborn features remains a concern.

### 8. Backlog Watch
- **#4108** ([Issue](https://github.com/nearai/ironclaw/issues/4108)) — Nightly E2E failure has been open since May 27 with no fix PR. This is a high-priority item that should not linger.
- **#5279** ([PR](https://github.com/nearai/ironclaw/pull/5279)) — Fix Reborn queued message steering (open 5 days, XL, core contributor). Large change waiting for review.
- **#5338** ([PR](https://github.com/nearai/ironclaw/pull/5338)) — Surface real failure details instead of generic error tokens (open 5 days, human-verified). Important for usability.
- **#5247** ([PR](https://github.com/nearai/ironclaw/pull/5247)) — Link approval card to global auto-approve settings (open 6 days). Small UX improvement that would reduce friction.

These items need maintainer attention to avoid blocking progress. The nightly E2E failure (#4108) is especially concerning as it undermines confidence in CI.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-07-01

## Today's Overview
The project experienced a burst of activity today, with **26 pull requests updated** in the last 24 hours—23 of which were merged or closed, including a batch of older (stale) PRs that had been open since early April. One new release (`2026.6.30`) was published. On the issue tracker, **4 issues were updated**: 1 closed (a stale shortcut validation bug) and 3 opened today, including a performance-critical report about the skills file watcher and a strategic feature suggestion. The overall project health appears strong, with a high merge throughput and active maintenance of both new and legacy contributions.

## Releases
**LobsterAI 2026.6.30** was released on June 30, 2026. Key changes include:
- `feat(logging)`: Added diagnostics for Cowork and OpenClaw flows.
- `fix(openclaw)`: Fallback catalog max token limits.
- Schedule-related improvements (details truncated in changelog).

No breaking changes or migration notes were mentioned.

## Project Progress
The following features, fixes, and improvements were merged/closed today (selected highlights):

- **New Features & Enhancements**  
  - **Subagent artifact panel** (PR [#2249](https://github.com/netease-youdao/LobsterAI/pull/2249)) – Added subagent-specific tab in artifact panel with list/detail views.  
  - **Auto-open preview cards** (PR [#2248](https://github.com/netease-youdao/LobsterAI/pull/2248)) – Automatically opens the most suitable preview tab after a new artifact is generated.  
  - **Qichacha MCP integration & grouped server management** (PR [#2244](https://github.com/netease-youdao/LobsterAI/pull/2244)) – Added Qichacha account authorization and improved multi-server MCP presentation.  
  - **Agent import/export** (PR [#1366](https://github.com/netease-youdao/LobsterAI/pull/1366) – merged today) – Backup and share agent configurations as JSON.  
  - **Scheduled task import/export** (PR [#1291](https://github.com/netease-youdao/LobsterAI/pull/1291) – merged today) – Package tasks as `.lobstertasks` files.  
  - **Sidebar collapse with icon bar** (PR [#1253](https://github.com/netease-youdao/LobsterAI/pull/1253) – merged today) – Retains icon navigation when sidebar is minimized.  
  - **Agent task statistics & collapsible list** (PR [#1171](https://github.com/netease-youdao/LobsterAI/pull/1171) – merged today) – Shows running/total badges and collapsible sections.  
  - **Streaming activity timer** (PR [#1548](https://github.com/netease-youdao/LobsterAI/pull/1548) – merged today) – Displays elapsed time for tool calls and streaming.

- **Bug Fixes**  
  - **Mac fullscreen black screen** (PR [#2246](https://github.com/netease-youdao/LobsterAI/pull/2246)) – Exit native fullscreen before hiding window.  
  - **Cowork plan recovery locking** (PR [#2247](https://github.com/netease-youdao/LobsterAI/pull/2247)) – Delay plan recovery until aborted OpenClaw run finishes.  
  - **Share deployment with isolated Node environment** (PR [#2251](https://github.com/netease-youdao/LobsterAI/pull/2251)) – Fixes missing tools and dependency issues.  
  - **Analytics reporting edge cases** (PR [#2245](https://github.com/netease-youdao/LobsterAI/pull/2245)) – Corrects usage events for skills, IM settings, sidebar, etc.  
  - **Windows drag & drop for .pptx/.docx** (PR [#1355](https://github.com/netease-youdao/LobsterAI/pull/1355) – merged today) – Handles virtual file descriptors from Windows Explorer.  
  - **Scheduled task stop IPC returning false success** (PR [#1424](https://github.com/netease-youdao/LobsterAI/pull/1424) – merged today) – Now propagates errors to UI instead of silently returning `{success: true}`.  
  - **Duplicate shortcut validation** (Issue [#1425](https://github.com/netease-youdao/LobsterAI/issues/1425) – closed) – Fixed lack of validation when saving duplicate key bindings.  

Additional minor fixes: compact prompt toolbar (PR [#2242](https://github.com/netease-youdao/LobsterAI/pull/2242)), one-click clear attachments & input (PR [#1242](https://github.com/netease-youdao/LobsterAI/pull/1242) – merged today), permission modal ESC support (PR [#1362](https://github.com/netease-youdao/LobsterAI/pull/1362) – still open), etc.

## Community Hot Topics
Activity levels are relatively low in terms of comments and reactions. The **most discussed issue** is:

- **[#1425 – 快捷键重复无校验](https://github.com/netease-youdao/LobsterAI/issues/1425)** (2 comments, now closed). The user reported that duplicate keyboard shortcuts could be saved without validation. The issue was filed in April and was finally closed today, likely due to the associated fix being merged.

- **[#2239 – 趋势判断：编程工具的“OpenClaw 化”和OpenClaw 类工具的编程工具化](https://github.com/netease-youdao/LobsterAI/issues/2239)** (opened today, 0 comments). While lacking discussion, this strategic proposal is likely to draw attention as it outlines a vision for LobsterAI to integrate with programming tools via MCP and deeper agent orchestration.

No PRs have visible comment counts (all reported as `undefined`), indicating either minimal discussion or a data capture limitation.

## Bugs & Stability
| Severity | Issue | Description | Status |
|----------|-------|-------------|--------|
| **High** | [#2243](https://github.com/netease-youdao/LobsterAI/issues/2243) | `skills.load.watch` causes performance bottleneck (scanning 174+ skills on every file change), persistent bug, and no UI toggle to disable it. | Open – filed today, no fix PR yet. |
| **Medium** | [#1361](https://github.com/netease-youdao/LobsterAI/issues/1361) | Custom agent detail page shows “delete” button in English instead of Chinese. | Open (stale since April). |
| **Low** | [#2246 – fix](https://github.com/netease-youdao/LobsterAI/pull/2246) | Black screen when closing macOS fullscreen app – fixed today. | Merged. |
| **Low** | [#2247 – fix](https://github.com/netease-youdao/LobsterAI/pull/2247) | Lock collision during Cowork plan recovery after abort – fixed today. | Merged. |
| **Low** | [#2242 – fix](https://github.com/netease-youdao/LobsterAI/pull/2242) | Prompt toolbar layout broken at narrow widths – fixed today. | Merged. |

No new crash or regression reports were filed today beyond #2243, which is a performance/persistence issue rather than a crash.

## Feature Requests & Roadmap Signals
Two notable user-requested features surfaced:

1. **[#2239 – OpenClaw-ization of LobsterAI](https://github.com/netease-youdao/LobsterAI/issues/2239)** – The user proposes deeper integration with programming tools (e.g., OpenCode, CodeBuddy CN) via MCP protocol, enabling full workflow automation. This aligns with today’s merged PR [#2244](https://github.com/netease-youdao/LobsterAI/pull/2244) (Qichacha MCP integration) and suggests future support for more agent-to-tool bridging.

2. **[#2243 – skills.load.watch UI toggle](https://github.com/netease-youdao/LobsterAI/issues/2243)** – The reporter requests a manual toggle with UI control to disable automatic file watching for skills, citing performance degradation with large skill libraries. This is a concrete usability improvement that may be fast-tracked given its direct impact on daily workflows.

Among today’s merged PRs, the **subagent artifact panel** (PR [#2249](https://github.com/netease-youdao/LobsterAI/pull/2249)) and **auto-open preview cards** (PR [#2248](https://github.com/netease-youdao/LobsterAI/pull/2248)) indicate a focus on making the artifact and cowork experience more fluid. These are likely to be part of the next release.

## User Feedback Summary
- **Pain Points**  
  - Performance: Skills file watcher causes significant slowdown for users with many skills (Issue [#2243](https://github.com/netease-youdao/LobsterAI/issues/2243)).  
  - Localization: Delete button remains English in Chinese UI for custom agent detail page (Issue [#1361](https://github.com/netease-youdao/LobsterAI/issues/1361)).  
  - Usability: No validation for duplicate keyboard shortcuts (Issue [#1425](https://github.com/netease-youdao/LobsterAI/issues/1425) – now fixed).  

- **Desired Features**  
  - Manual toggle for skills file watching.  
  - Model selector in the input toolbar (open PR [#1364](https://github.com/netease-youdao/LobsterAI/pull/1364)).  
  - ESC key to close permission dialogs (open PR [#1362](https://github.com/netease-youdao/LobsterAI/pull/1362)).  
  - Integration with programming tools (Issue [#2239](https://github.com/netease-youdao/LobsterAI/issues/2239)).  

- **Satisfaction Indicators**  
  - High merge velocity (23 PRs merged today) suggests active development and responsiveness.  
  - Many older feature PRs (Agent export/import, scheduled task import/export, sidebar improvements) were finally merged, indicating maintainers value backlog clearing.

## Backlog Watch
The following items have been open for an extended period (since early April 2026) and may require maintainer attention:

| Item | Type | Since | Last Update | Status |
|------|------|-------|-------------|--------|
| [#1361 – Delete button English](https://github.com/netease-youdao/LobsterAI/issues/1361) | Issue | 2026-04-02 | 2026-07-01 (updated? or just stale bump) | Open, no assignee |
| [#1362 – ESC close for permission modal](https://github.com/netease-youdao/LobsterAI/pull/1362) | PR | 2026-04-02 | 2026-07-01 | Open, no comments |
| [#1364 – Model selector in input toolbar](https://github.com/netease-youdao/LobsterAI/pull/1364) | PR | 2026-04-02 | 2026-07-01 | Open, no comments |
| [#1367 – Validate duplicate task names](https://github.com/netease-youdao/LobsterAI/pull/1367) | PR | 2026-04-02 | 2026-07-01 | Open, no comments |

*Note: All three open PRs have zero comments and no visible review activity. They were updated today (likely because of automated notification when the repo was active), but have not been merged. The issue #1361 also appears untouched by maintainers.*

**Additionally**, the new performance bug [#2243](https://github.com/netease-youdao/LobsterAI/issues/2243) is currently unaddressed – no fix PR or assignee yet. Given its severity, it should be prioritized.

---

*Data source: GitHub data from [netease-youdao/LobsterAI](https://github.com/netease-youdao/LobsterAI), snapshot taken 2026-07-01.*

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

## Moltis Project Digest – 2026-07-01

### 1. Today's Overview
Moltis showed minimal human-driven activity over the past 24 hours. No new issues were created or updated, and the only PRs touched are automated dependency bumps by Dependabot (one still open, two merged). No new releases were published. The project is in a maintenance phase with no visible feature work or bug reports, indicating a low-activity period. Community engagement appears dormant.

### 2. Releases
*No new releases.* The latest release remains unknown from this data snapshot.

### 3. Project Progress
Two Dependabot PRs were merged/closed in the last 24 hours:
- **#1134** – chore(deps): bump npm_and_yarn group across 2 directories with 2 updates (astro 6.3.3→6.4.8, undici). [Closed 2026-06-30](https://github.com/moltis-org/moltis/pull/1134)
- **#1121** – chore(deps-dev): bump esbuild from 0.25.12 to 0.28.1 in `/crates/web/ui`. [Closed 2026-06-30](https://github.com/moltis-org/moltis/pull/1121)

These updates improve dependency security and compatibility but do not introduce new features or fix user-reported bugs.

### 4. Community Hot Topics
No issues or PRs received any comments or reactions in the last 24 hours. The only open PR is **#1141** (dependabot bumping esbuild and vite). It has no discussions. The lack of community interaction suggests either low usage or that all recent needs are satisfied.

### 5. Bugs & Stability
No new bugs, crashes, or regressions were reported in the last 24 hours. No fix PRs exist. The project is stable from a code quality perspective, though this may reflect low usage rather than true robustness.

### 6. Feature Requests & Roadmap Signals
No feature requests were filed or discussed. There are no signals pointing toward upcoming features. The roadmap remains unclear; next versions are likely to focus on maintaining dependency freshness and security.

### 7. User Feedback Summary
No user feedback (issues, comments, reactions) was observed in the last 24 hours. It is not possible to gauge satisfaction or pain points from this data.

### 8. Backlog Watch
No issues or PRs are awaiting maintainer attention. The single open PR **#1141** is a trivial Dependabot update that will likely auto-merge or be handled quickly. There are no long-unanswered items in the backlog.

**Project Health Assessment:** Low activity. Moltis is in a quiet period with only automated dependency chores. Monitor for signs of revived development or community engagement.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest – 2026-07-01

*Data source: GitHub (agentscope-ai/CoPaw) — reflects activity as of 2026-07-01 23:59 UTC.*

---

## 1. Today's Overview

The CoPaw project (also referred to as **QwenPaw** in code and issues) is in an **intense development sprint**, with **50 PRs updated** and **25 merged/closed** in the last 24 hours — indicating a fast-paced release pipeline. The bug tracker is also active: **12 issues** were updated, 9 of which remain open, covering a mix of regressions in the v2.0.0 pre-release and new feature requests. However, **no new release** was tagged today; the latest stable is still v1.1.12.post2, while the pre-release v2.0.0b2 (beta) is being actively patched. Maintenance focus appears to be on **channel integration fixes**, **memory/reranker enhancements**, and **plugin system reliability**.

---

## 2. Releases

**No new releases today.** The last published version remains **v1.1.12.post2** (stable) and **v2.0.0b2** (pre-release). A new release is likely imminent given the volume of merged fixes.

---

## 3. Project Progress

Today **25 PRs were merged or closed**. Notable advances among the top-20 list:

| PR | Title | Impact |
|----|-------|--------|
| [#5574](https://github.com/agentscope-ai/QwenPaw/pull/5574) | fix: refresh chat on channel session updates | First-time contributor; Web UI now auto-refreshes when channel messages arrive. |
| [#5562](https://github.com/agentscope-ai/QwenPaw/pull/5562) | fix: re-enqueue in-flight batch on CancelledError | Prevents message loss during hot-reload. |
| [#5690](https://github.com/agentscope-ai/QwenPaw/pull/5690) | fix: add `audio` to `_FORMATTER_SKIPPED_TYPES` | Fixes reasoning content alignment when audio blocks are present. |
| [#5699](https://github.com/agentscope-ai/QwenPaw/pull/5699) | fix(telegram): move typing indicator start | Prevents "typing…" from being stuck for 180s on file-only messages. |
| [#5510](https://github.com/agentscope-ai/QwenPaw/pull/5510) | fix(tool-calls): cap tool responses before context insertion | Defense-in-depth against context explosion when LLM call fails. |
| [#5068](https://github.com/agentscope-ai/QwenPaw/pull/5068) | fix(e2e): fix token usage test empty state detection | Resolves nightly E2E test failure in clean environments. |

Additionally, several **new feature PRs remain open** (see Section 6), including Loop Engineering, memory reranker, and computer-use automation.

---

## 4. Community Hot Topics

Most active discussions (by comments and reactions):

- **[Issue #5630](https://github.com/agentscope-ai/QwenPaw/issues/5630) – Support custom BaseURL for Telegram channel**  
  *8 comments, 0 👍*  
  User requests ability to set a custom Telegram API base URL (e.g., for proxy/censorship circumvention). No maintainer response yet. Underlying need: **enterprise/proxy environments** require configurable endpoints.

- **[Issue #5063](https://github.com/agentscope-ai/QwenPaw/issues/5063) – Integrate Headroom context compression (closed)**  
  *8 comments, 0 👍*  
  Feature request seeking 60–95% token reduction via a reversible compression layer. Although closed, the discussion suggests strong community interest in **cost-saving context management**.

- **[Issue #5273](https://github.com/agentscope-ai/QwenPaw/issues/5273) – v2.0.0 Pre-release Bug Tracker**  
  *2 comments, 1 👍*  
  Central tracking issue for all v2.0.0 regressions. The only issue with a thumbs-up reaction — signals community engagement.

- **[Issue #5701](https://github.com/agentscope-ai/QwenPaw/issues/5701) – Concurrent access freeze**  
  *2 comments*  
  Multiple agent sessions cause the system to hang. A critical stability concern (see Bugs section).

- **[Issue #5676](https://github.com/agentscope-ai/QwenPaw/issues/5676) – Skills not listed in system prompt**  
  *2 comments*  
  Discoverability of available skills is broken in v2.0.0b2, contrary to official guidelines.

---

## 5. Bugs & Stability

### High Severity

| Issue | Description | Status |
|-------|-------------|--------|
| [#5701](https://github.com/agentscope-ai/QwenPaw/issues/5701) | Concurrent access to the same agent UI freezes the system (v1.1.10). | Open, no fix PR. Likely a server/connection pool issue. |
| [#5689](https://github.com/agentscope-ai/QwenPaw/issues/5689) | Remote SSH plugin removal leaves stale imports, breaking all conversations (v1.1.12.post2). | Open, but related PR [#5695](https://github.com/agentscope-ai/QwenPaw/pull/5695) (still open) addresses plugin version retention. |

### Medium Severity

| Issue | Description | Status |
|-------|-------------|--------|
| [#5696](https://github.com/agentscope-ai/QwenPaw/issues/5696) | QQ Channel websocket reconnect causes `_http` to become None, throwing AttributeError. | Open, no fix PR. Affects all QQ bot users. |
| [#5703](https://github.com/agentscope-ai/QwenPaw/issues/5703) | "Disable all tool approval" setting is not persisted; approval windows still appear. | Open, no fix PR. Regression from v1.1.x to v2.0.0b1. |
| [#5676](https://github.com/agentscope-ai/QwenPaw/issues/5676) | Available skills metadata missing from system prompt in v2.0.0b2. | Open, no fix PR. |

### Low Severity

| Issue | Description | Status |
|-------|-------------|--------|
| [#5688](https://github.com/agentscope-ai/QwenPaw/issues/5688) | CSS selector prefix mismatch (`ant-` vs configured `qwenpaw-`) | Open, no fix PR. Visual inconsistency. |

**Note:** Several bug-fix PRs were merged today (see Section 3), including fixes for message loss during hot-reload, Telegram typing indicator, and tool-response truncation — indicating responsive maintainers.

---

## 6. Feature Requests & Roadmap Signals

### High-priority community requests

| Issue | Feature | Likely inclusion |
|-------|---------|------------------|
| [#5630](https://github.com/agentscope-ai/QwenPaw/issues/5630) | Custom BaseURL for Telegram | High — simple config change; similar patterns exist for other channels. |
| [#5063](https://github.com/agentscope-ai/QwenPaw/issues/5063) | Headroom context compression (closed, but not yet implemented) | Medium — may be part of v2.0.0 roadmap. |
| [#5703](https://github.com/agentscope-ai/QwenPaw/issues/5703) | Reliable tool approval toggle | High — regression fix, likely in next patch. |

### Signals from open PRs (features in progress)

The following PRs are open and likely to land in the **next minor release**:

- **[#5665](https://github.com/agentscope-ai/QwenPaw/pull/5665)** – Loop Engineering: composable gate architecture with frontend settings.
- **[#5691](https://github.com/agentscope-ai/QwenPaw/pull/5691) & [#5692](https://github.com/agentscope-ai/QwenPaw/pull/5692)** – Reranker configuration UI and backend for reme0.4 memory search.
- **[#5187](https://github.com/agentscope-ai/QwenPaw/pull/5187)** – Windows desktop GUI automation (computer use with UIA + Tauri).
- **[#5697](https://github.com/agentscope-ai/QwenPaw/pull/5697)** – New blog and docs section on the website.
- **[#5296](https://github.com/agentscope-ai/QwenPaw/pull/5296)** – ADBPG memory: remove SQL mode, keep REST-only + auto-search.

The **v2.0.0 pre-release** is clearly targeting a major upgrade to memory, tool orchestration, and user interface, but several regressions still need to be resolved before GA.

---

## 7. User Feedback Summary

### Pain points

- **Stability regressions in v2.0.0 pre-release:** Users report that tool approval settings are ignored ([#5703](https://github.com/agentscope-ai/QwenPaw/issues/5703)) and skills are not discoverable ([#5676](https://github.com/agentscope-ai/QwenPaw/issues/5676)).
- **Concurrency issues:** Multi-tab or multi-user access to the same agent causes freezes ([#5701](https://github.com/agentscope-ai/QwenPaw/issues/5701)) — a deal-breaker for team deployments.
- **Channel integration quirks:** Telegram typing indicator stuck for 180s on file-only messages (now patched in #5699), QQ bot crashes on reconnect ([#5696](https://github.com/agentscope-ai/QwenPaw/issues/5696)), and custom BaseURL is not supported ([#5630](https://github.com/agentscope-ai/QwenPaw/issues/5630)).
- **Plugin removal leaves broken state:** Deleting the Remote SSH plugin breaks all conversations until manual cleanup ([#5689](https://github.com/agentscope-ai/QwenPaw/issues/5689)).

### Satisfaction signals

- First-time contributors are actively submitting fixes ([#5574](https://github.com/agentscope-ai/QwenPaw/pull/5574), [#5562](https://github.com/agentscope-ai/QwenPaw/pull/5562), [#5690](https://github.com/agentscope-ai/QwenPaw/pull/5690))—a sign of healthy community engagement.
- The **v2.0.0 pre-release tracker** (#5273) has positive sentiment (1 👍), indicating users are willing to help test.
- Large features like **Loop Engineering** and **memory reranker** are being developed transparently, which users appreciate.

### Use cases

- **Enterprise/censored environments:** Need customizable endpoints (Telegram BaseURL).
- **Multi-user deployments:** Concurrent access and approval control are critical.
- **Channel-heavy workflows:** WeChat/QQ/Telegram bots are core use cases; reliability is essential.
- **Cost-sensitive users:** Context compression (Headroom) remains a desired feature.

---

## 8. Backlog Watch

The following issues/PRs have not received maintainer attention in an unusually long time (relative to project tempo) or have high impact:

| Item | Last Update | Days Stale | Priority |
|------|-------------|------------|----------|
| [#5630](https://github.com/agentscope-ai/QwenPaw/issues/5630) – Telegram custom BaseURL | 2026-07-01 (today) | <1 day | **Low** – but no assignee or maintainer comment. Should be triaged. |
| [#5063](https://github.com/agentscope-ai/QwenPaw/issues/5063) – Headroom integration (closed but not implemented) | 2026-07-01 (closed) | – | **Medium** – community expected this to be implemented. Consider roadmap communication. |
| [#5523](https://github.com/agentscope-ai/QwenPaw/issues/5523) – `spawn_subagent` missing from Runtime 2.0 tool registry (closed) | 2026-07-01 (closed) | 6 days | **Low** – fixed via PRs? Check if merged. |
| [#5676](https://github.com/agentscope-ai/QwenPaw/issues/5676) – Skills not listed in system prompt | 2026-07-01 | <1 day | **High** – breaks agent skill discovery. No response yet. |
| [#5696](https://github.com/agentscope-ai/QwenPaw/issues/5696) – QQ channel websocket reconnect crash | 2026-07-01 | <1 day | **High** – affects all QQ bot users. No fix PR active. |

**Recommendation:** The maintainer team should prioritize triaging the high-severity open bugs ([#5701](https://github.com/agentscope-ai/QwenPaw/issues/5701), [#5676](https://github.com/agentscope-ai/QwenPaw/issues/5676), [#5696](https://github.com/agentscope-ai/QwenPaw/issues/5696), [#5703](https://github.com/agentscope-ai/QwenPaw/issues/5703)) before adding new features, especially given the velocity of the v2.0.0 line. The first-time contributor momentum is positive and should be encouraged with quick feedback cycles.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-07-01

## 1. Today's Overview
The project is in a period of high activity, with **17 issues** and **50 pull requests** updated in the last 24 hours. **16 issues remain open**, reflecting active bug reporting and feature requests, while **5 PRs were merged/closed**, advancing critical infrastructure components. Activity is concentrated on security and CI hardening, web dashboard stability, and new interoperability features (OpenAI‑compatible API, Mixture‑of‑Agents). Many high‑severity bugs (S1 workflow‑blocked) are still awaiting fix PRs, indicating that reliability work is ongoing. The community appears engaged but is hitting real blockers in production use.

## 2. Releases
No new releases were published today.

## 3. Project Progress (Merged/Closed PRs Today)
Five pull requests were merged or closed in the last 24 hours (three captured in the top‑20 list):

- **AMQP SOP fan‑in dispatch path** – [PR #8521](https://github.com/zeroclaw-labs/zeroclaw/pull/8521) (closed)  
  Adds SOP‑driven dispatch for AMQP deliveries, supports `sop` and `sop_and_agent_loop` modes, and fixes AMQP credential secrets.

- **Slack thread history scope** – [PR #8579](https://github.com/zeroclaw-labs/zeroclaw/pull/8579) (closed)  
  Introduces `channels.slack.<alias>.thread_history_scope` with granular sender/thread/channel options; default remains per‑sender.

- **Gateway A2A discovery port fix** – [PR #8549](https://github.com/zeroclaw-labs/zeroclaw/pull/8549) (closed)  
  Advertises the actual listener port in A2A discovery cards when `--port` override or `--port 0` is used.

- **Two additional PRs** (not shown in top‑20) were also closed – likely minor fixes or dependency bumps.

Merged work strengthens channel stability (AMQP, Slack) and corrects a gateway configuration defect.

## 4. Community Hot Topics

### Most‑discussed Issue
- **[#8057](https://github.com/zeroclaw-labs/zeroclaw/issues/8057) – CI: scheduled/manual security jobs**  
  (3 comments) – Split from a larger enforcement‑level discussion, this issue proposes adding CodeQL, Trivy, SBOM, and other heavy checks that run off the critical PR path. It reflects growing concern about supply‑chain security and long‑term maintainability.

### High‑interest Feature Proposals
- **OpenAI‑compatible chat completions** – [#8550](https://github.com/zeroclaw-labs/zeroclaw/issues/8550) and companion [PR #8486](https://github.com/zeroclaw-labs/zeroclaw/pull/8486)  
  Wanted for integrating with tools like Open WebUI and Continue.dev. The feature is critical for broadening ZeroClaw’s ecosystem reach.

- **Mixture‑of‑Agents (MoA) virtual model provider** – [#8568](https://github.com/zeroclaw-labs/zeroclaw/issues/8568)  
  A proposal to allow multiple models to contribute analysis to a single aggregator/judge model. This signals interest in multi‑model reasoning without custom orchestration.

### Large, Pending PRs with Broad Impact
- **[Git forge channel (GitHub App) + SOP ingress](https://github.com/zeroclaw-labs/zeroclaw/pull/8504)** (XL size) – Adds polling, issue/PR comment events, and SOP support for GitHub, indicating a move toward DevOps and developer workflow integration.

- **[Plugin channel host bindings](https://github.com/zeroclaw-labs/zeroclaw/pull/8551)** (L size) – WASM plugin infrastructure for network channels, enabling third‑party channel development without core binary changes.

- **[OpenAI endpoint PR #8486](https://github.com/zeroclaw-labs/zeroclaw/pull/8486)** (XL size, needs‑author‑action) – Blocked on author updates but remains one of the most awaited integrations.

## 5. Bugs & Stability

### Severity S1 — Workflow Blocked
These bugs are blocking users from completing essential tasks. None currently have an associated fix PR in the data, making them urgent:

- **[#8559](https://github.com/zeroclaw-labs/zeroclaw/issues/8559) – Agents stop when exiting chat window**  
  Agent loops are interrupted on web dashboard chat close. **No fix PR yet.**

- **[#8563](https://github.com/zeroclaw-labs/zeroclaw/issues/8563) – SOPs not available in web dashboard chat**  
  Shared SOPs are ignored by the agent runtime when started from the web UI. **No fix PR yet.**

- **[#8553](https://github.com/zeroclaw-labs/zeroclaw/issues/8553) – Environment variables cannot be used as HTTP request secrets**  
  No path exists for agents to use env‑var secrets with the `http_request` tool. **No fix PR yet.**

- **[#8560](https://github.com/zeroclaw-labs/zeroclaw/issues/8560) – `browser_open` hangs agent turn on headless hosts**  
  Subprocess wait is unbounded; also affects TTS and ffmpeg in channels. **No fix PR yet.**

### Severity S2 — Degraded Behavior
- **[#8554](https://github.com/zeroclaw-labs/zeroclaw/issues/8554) – Zip‑bomb hardening needed in skill extractor** (status: in‑progress)  
  Missing decompressed‑size and entry‑count limits. No fix PR yet, but marked as accepted/in‑progress.

### Severity S3 — Minor Issue
- **[#8578](https://github.com/zeroclaw-labs/zeroclaw/issues/8578) – `zerocode` daemon does not terminate on startup failure**  
  A fix is already in review: [PR #8582](https://github.com/zeroclaw-labs/zeroclaw/pull/8582) introduces proper ephemeral daemon cleanup.

## 6. Feature Requests & Roadmap Signals
The set of open RFCs and enhancements points toward several likely directions for the next release:

- **Interoperability first**: OpenAI‑compatible endpoint ([#8550](https://github.com/zeroclaw-labs/zeroclaw/issues/8550)) and MoA virtual provider ([#8568](https://github.com/zeroclaw-labs/zeroclaw/issues/8568)) are both tagged with `priority:p2` but have significant community momentum. The OpenAI endpoint already has a companion PR (#8486) and may merge soon.

- **Channel and SOP refinements**: Centralized SOP ingress adapters ([#8581](https://github.com/zeroclaw-labs/zeroclaw/issues/8581)), webhook dispatch refactor ([#8586](https://github.com/zeroclaw-labs/zeroclaw/issues/8586)), and Matrix thread‑scoped history ([#8541](https://github.com/zeroclaw-labs/zeroclaw/issues/8541)) show a pattern of improving channel architecture for reusability.

- **UX and localization**: Web dashboard localization via Fluent ([#8584](https://github.com/zeroclaw-labs/zeroclaw/issues/8584)) and secret‑field state indicators ([#8556](https://github.com/zeroclaw-labs/zeroclaw/issues/8556)) are small but user‑visible improvements.

- **Documentation gaps**: [#8587](https://github.com/zeroclaw-labs/zeroclaw/issues/8587) requests more SOP syntax examples, indicating that the SOP engine is gaining adoption but lacks clear tutorials.

## 7. User Feedback Summary
Users are encountering real production blockers:

- **Workflow interruption** – exiting a web chat kills a running agent, making background tasks impossible (#8559).
- **Missing SOP integration** – shared SOPs configured on disk do not load in web‑session contexts (#8563).
- **Secret management friction** – environment‑variable secrets cannot be used in HTTP tools (#8553), and the UI does not show whether a secret is already set (#8556).
- **Hanging tools** – `browser_open` on headless hosts hangs indefinitely (#8560).
- **Language support** – Spanish translations are incomplete and locale detection is broken (PR #8589, still open).

Positive signals: the community is actively contributing translation fixes (PR #8589) and security documentation corrections (PR #8588). Interest in advanced features like MoA and Git forge channels suggests a maturing user base that wants to scale ZeroClaw.

## 8. Backlog Watch
Several important items are waiting for maintainer attention or author updates:

- **OpenAI‑compatible endpoint** – [#8550](https://github.com/zeroclaw-labs/zeroclaw/issues/8550) and its implementation [PR #8486](https://github.com/zeroclaw-labs/zeroclaw/pull/8486) are tagged `needs‑maintainer‑review` and `needs‑author‑action`, respectively. This is the most‑requested feature and has stalled.

- **Environment variables as secrets** – [#8553](https://github.com/zeroclaw-labs/zeroclaw/issues/8553) (S1 blocker) has no associated PR and is still `needs‑maintainer‑review`.

- **Matrix thread history** – [#8541](https://github.com/zeroclaw-labs/zeroclaw/issues/8541) also `needs‑maintainer‑review`.

- **CI security jobs** – [#8057](https://github.com/zeroclaw-labs/zeroclaw/issues/8057) is an accepted `priority:p2` with `risk:high` and 3 comments, but no PR yet. The split from #7675 still needs implementation.

- **Cron shutdown fix** – [PR #8465](https://github.com/zeroclaw-labs/zeroclaw/pull/8465) adds a CancellationToken to the cron loop but is still open with no recent comments.

- **Unbounded input fix** – [PR #8463](https://github.com/zeroclaw-labs/zeroclaw/pull/8463) caps interactive CLI stdin to 1 MiB but is marked `needs‑author‑action`.

These items represent both community pain points and architectural improvements that, if unaddressed, may slow adoption.

---
*Data from GitHub (zeroclaw-labs/zeroclaw) – issues and PRs updated between 2026‑06‑30 and 2026‑07‑01.*

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*