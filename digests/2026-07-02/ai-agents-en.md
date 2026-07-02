# OpenClaw Ecosystem Digest 2026-07-02

> Issues: 58 | PRs: 500 | Projects covered: 13 | Generated: 2026-07-02 10:17 UTC

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

# OpenClaw Project Digest — 2026-07-02

## Today's Overview

OpenClaw saw **very high activity** on 2026-07-02, with 58 issues updated (36 open, 22 closed) and 500 PRs updated (416 open, 84 merged/closed) in the last 24 hours. A new beta release **v2026.7.1-beta.1** landed, adding GPT-5.6 support and an external harness attachment command. The project remains intensely maintained, with dozens of bug reports and PRs moving forward, though the sheer volume of open PRs (416) and long-lived critical issues indicates a growing backlog that may require prioritization. Community engagement is strong, especially around session-state and message-loss bugs.

## Releases

**v2026.7.1-beta.1** (https://github.com/openclaw/openclaw/releases/tag/v2026.7.1-beta.1)

- **New features:**
  - **OpenAI GPT-5.6 support** — OpenClaw now recognizes the GPT-5.6 model family across catalog, capability, and runtime selection paths. (#98333, thanks @steipete-oai)
  - **External harness attachment** — `openclaw attach` launches an external harness against an existing Gateway session.
- **Breaking changes:** None documented in the highlights.
- **Migration notes:** No explicit migration steps provided; likely a minor upgrade.

## Project Progress

Today **84 PRs were merged or closed**, reflecting active feature work and bug fixing. Notable merged/closed items visible from the provided data:

- **Fixed critical rendering bugs:**
  - [#98874](https://github.com/openclaw/openclaw/issues/98874) — Tool text results (exec, read, browser) rendering as image attachments fixed (closed).
  - [#99022](https://github.com/openclaw/openclaw/issues/99022) — Chinese bug report: exec/read/memory_search tools returned image placeholders instead of text output (closed).
  - [#98872](https://github.com/openclaw/openclaw/issues/98872) — Terminal table display incorrectly rewrote home directory prefixed paths (closed).
- **Channel-specific fixes:**
  - [#98871](https://github.com/openclaw/openclaw/issues/98871) — Mattermost peer directory omitted users beyond the first 200 team members (closed).
  - [#98971](https://github.com/openclaw/openclaw/issues/98971) — Added delete button for individual chat messages in WebChat (closed).
  - [#98929](https://github.com/openclaw/openclaw/issues/98929) — Improved native iOS chat hierarchy and composer (closed).
- **Infrastructure and security:**
  - [#98550](https://github.com/openclaw/openclaw/issues/98550) — CVE assigned for GHSA-cf2p-f286-mphf (identity-bearing HTTP callers reaching admin-scoped tools, closed).
  - [#98973](https://github.com/openclaw/openclaw/issues/98973) — Gateway restart on Windows from inside the process fixed (closed).
- **PRs merged** (examples from top 30):
  - [#95195](https://github.com/openclaw/openclaw/pull/95195) — Fix(fix) bound Discord gateway metadata reads to prevent OOM.
  - [#93516](https://github.com/openclaw/openclaw/pull/93516) — Fix Matrix reverse-proxy path support for homeserver URLs behind proxies.
  - [#98640](https://github.com/openclaw/openclaw/pull/98640) — Fix Anthropic Opus 4.8 compatibility (add `additionalProperties:false` to tool inputSchema).

Many of these fixes address long-standing regressions and user complaints.

## Community Hot Topics

The most active issues and PRs (by comment count) highlight persistent pain points:

| Issue | Comments | Summary |
|-------|----------|---------|
| [#25592](https://github.com/openclaw/openclaw/issues/25592) — [OPEN, P1] | 33 | Text between tool calls leaks to messaging channels (Slack, iMessage, etc.). Highly upvoted (👍1), classified as diamond lobster. |
| [#38327](https://github.com/openclaw/openclaw/issues/38327) — [OPEN, P1] | 10 | Regression: "Cannot convert undefined or null to object" with google-vertex/gemini-3.1-pro-preview in 2026.3.2. 3 👍. |
| [#92433](https://github.com/openclaw/openclaw/issues/92433) — [OPEN, P1] | 7 | Subagent completion silently dropped when announce steers into requester run ending early. |
| [#70024](https://github.com/openclaw/openclaw/issues/70024) — [OPEN, P1] | 4 | Channel stop timeout leaves channel permanently dead with stale store entries. |
| [#32530](https://github.com/openclaw/openclaw/issues/32530) — [OPEN, P2] | 3 | Feature request: Auto-discovery of agent configurations from external workspaces. |

**Analysis:** The top issues all revolve around **session state integrity and message delivery** — text leaking, completions dropped, channels becoming dead, and regressions with specific providers. Users are demanding reliable, predictable behavior in production multi-channel deployments.

## Bugs & Stability

Today saw a **wave of new bug reports**, many of high severity (P1). Below are the most critical, ranked by severity and likely impact, along with any existing fix PRs:

### Critical (P1, impact: session-state / message-loss / crash-loop)
- **[#98956](https://github.com/openclaw/openclaw/issues/98956)** — Embedded abort settle timed out causes gateway deadlock after provider chain timeout; requires manual restart. No fix PR yet.
- **[#98938](https://github.com/openclaw/openclaw/issues/98938)** — Gradual JS-heap object-graph retention (~4.2 MB/min) leading to daily OOM on long-lived multi-account Matrix gateways. No fix PR.
- **[#98976](https://github.com/openclaw/openclaw/issues/98976)** — Provider refusals (Anthropic refusal / OpenAI content_filter) never trigger model fallback chain; turn dies with generic error. No fix PR.
- **[#99021](https://github.com/openclaw/openclaw/issues/99021)** — Discord reply with >10MB attachment fails with 413 and is silently lost (text + attachment). No fix PR.
- **[#99031](https://github.com/openclaw/openclaw/issues/99031)** — iOS discovered gateway fails with missing TLS / no trusted fingerprint, blocking connection. No fix PR.
- **[#98925](https://github.com/openclaw/openclaw/issues/98925)** — `fetchWithSsrFGuard` strict mode resolves DNS locally before managed proxy, causing failures in proxy-only sandboxes. No fix PR.

### High (P1/P2, impact: data-loss, auth-provider)
- **[#98945](https://github.com/openclaw/openclaw/issues/98945)** — Gemini CLI credential staging races in-flight runs and writes token files non-atomically (security, data-loss). No fix PR.
- **[#98964](https://github.com/openclaw/openclaw/issues/98964)** — `dispatch-wrapper`: flock `--` separator skipped and script command misdetected (security). No fix PR.
- **[#98982](https://github.com/openclaw/openclaw/issues/98982)** — Compaction dead-end: recently compacted session with overflow blocks permanently (session-state). No fix PR.
- **[#98614](https://github.com/openclaw/openclaw/issues/98614)** — Regression: `sessions_spawn` missing `scope: operator.write` between v2026.6.1 and v2026.6.11. No fix PR.
- **[#99022](https://github.com/openclaw/openclaw/issues/99022)** (closed) — exec/read/memory_search tools returning image placeholders in Windows 11. **Fix was merged** (closed).

### Moderate (P2/P3)
- [#98978](https://github.com/openclaw/openclaw/issues/98978) — `openclaw --help` descriptions contradict subcommand help (cosmetic).
- [#98970](https://github.com/openclaw/openclaw/issues/98970) — Default English audio prompt biases non-English STT.
- [#98960](https://github.com/openclaw/openclaw/issues/98960) — Usage-bar template file cache grows unbounded with live fs.watch watchers.
- [#98959](https://github.com/openclaw/openclaw/issues/98959) — proxy-capture: deletePathBasedSessions runs two DELETEs without transaction.

**Summary:** The project experienced a significant bug influx today, especially around network reliability, provider fallback, and memory management. Most of these lack corresponding fix PRs, suggesting maintainers are still triaging.

## Feature Requests & Roadmap Signals

New feature requests submitted today indicate strong user demand for:

- **Better voice/STT support:** [#98927](https://github.com/openclaw/openclaw/issues/98927) — Doubao/Volcengine realtime Talk mode easier to use (P3). Users in mainland China want clearer path.
- **Diagnostics for plugin LLM calls:** [#98968](https://github.com/openclaw/openclaw/issues/98968) — Emit `model.usage` diagnostics for plugin `runtime.llm.complete` calls (P2). Suggests growing plugin ecosystem.
- **iOS UI refinements:** [#98995](https://github.com/openclaw/openclaw/issues/98995), [#98943](https://github.com/openclaw/openclaw/issues/98943) — Appearance selector placement, About page copy improvement (both P3). Reflects investment in mobile client.
- **Transcript store migration to SQLite:** [#98986](https://github.com/openclaw/openclaw/issues/98986) — For meeting capture transcripts (P2). Aligns with project’s “SQLite only” direction. A draft PR [#99006](https://github.com/openclaw/openclaw/pull/99006) already exists.
- **Per-agent filesystem roots:** PR [#52951](https://github.com/openclaw/openclaw/pull/52951) (still open) — `tools.fs.roots` with access modes. Longstanding feature (March) for multi-tenant deployments.
- **Durable session task runtime:** PR [#98718](https://github.com/openclaw/openclaw/pull/98718) (open) — Foundation for long-running sessions surviving process restarts. Likely a candidate for next minor release.
- **User-specific memory isolation:** PR [#47277](https://github.com/openclaw/openclaw/pull/47277) (open) — Per-user memory files. Essential for multi-user gateways.

**Prediction:** The next release (v2026.7.x) may include SQLite transcript migration, GPT-5.6 support (already in beta), and potentially the durable session runtime if merged. iOS improvements and per-agent filesystem roots are likely further out.

## User Feedback Summary

**Pain points voiced today:**

- **Reliability:** Users report silent failures (Discord 413 lost messages #99021, provider refusals not falling back #98976, gateway deadlock #98956). This erodes trust.
- **Regressions:** Several users mention "worked before, now fails" (#38327, #98614, #99022). Frequent regressions harm adoption.
- **Mobile/iOS:** iOS TLS fingerprint issue (#99031) blocks discovery; Android chat.send not reaching gateway (#91872) remains open since June 10. Mobile experience needs attention.
- **CLI/help confusion:** [#98978] shows documentation drift, confusing new users.

**Satisfaction signals:**

- Positive reception of GPT-5.6 support and external harness attachment in the new beta.
- Quick closure of some bugs (e.g., #98874, #99022) shows maintainers are responsive.
- Feature requests are well received (👍1-3 per request) indicating engaged community.

## Backlog Watch

The following important issues and PRs have been **open for weeks or months** without closure, requiring maintainer attention:

### Long-open Critical Issues
- [#25592](https://github.com/openclaw/openclaw/issues/25592) — [P1] Text between tool calls leaks (33 comments, since Feb 24). **No fix PR.**
- [#38327](https://github.com/openclaw/openclaw/issues/38327) — [P1] Regression with gemini-3.1-pro-preview (10 comments, since Mar 6). **No fix PR.**
- [#70024](https://github.com/openclaw/openclaw/issues/70024) — [P1] Channel stop timeout dead channel (4 comments, since Apr 22). **No fix PR.**
- [#92433](https://github.com/openclaw/openclaw/issues/92433) — [P1] Subagent completion dropped (7 comments, since Jun 12). **No fix PR.**
- [#91872](https://github.com/openclaw/openclaw/issues/91872) — [P1] Android chat.send never reaches gateway (since Jun 10). **No fix PR.**

### Long-open High-Value PRs
- [#53467](https://github.com/openclaw/openclaw/pull/53467) — [P2] Slack `ignoreOtherMentions` config (since Mar 24, ready for maintainer look).
- [#52951](https://github.com/openclaw/openclaw/pull/52951) — [P2] `tools.fs.roots` per-agent filesystem roots (since Mar 23, ready for maintainer look).
- [#51762](https://github.com/openclaw/openclaw/pull/51762) — [P2] Configurable default agent ID (since Mar 21, waiting on author).
- [#51822](https://github.com/openclaw/openclaw/pull/51822) — [P2] Reject cron webhook URLs with embedded credentials (since Mar 21, needs proof).
- [#47604](https://github.com/openclaw/openclaw/pull/47604) — [P2] Wear OS app MVP (since Mar 15, needs proof).
- [#47277](https://github.com/openclaw/openclaw/pull/47277) — [P1] User-specific memory isolation (since Mar 15, needs proof).
- [#48940](https://github.com/openclaw/openclaw/pull/48940) — [P1] ACP gateway-owned node-backed runtime (since Mar 17, needs proof).

These items represent both **critical bugs** and **high-demand features** that have stalled. The maintainer team should prioritize reviewing and merging/responding to these to reduce backlogs and improve user trust.

---

## Cross-Ecosystem Comparison

Based on the community digest summaries for 2026-07-02 across 13 tracked projects (7 with significant activity), here is the cross-project comparison report.

---

## 1. Ecosystem Overview

The open-source personal AI assistant ecosystem is in a phase of rapid, concurrent development across multiple projects, each targeting different niches but converging on common challenges. Activity levels are high, with hundreds of pull requests and dozens of issues updated daily. Projects are prioritizing multi-channel reliability (Slack, Discord, Telegram, Feishu, Matrix), expanding model provider support (GPT‑5.6, Gemini, Anthropic, Claude, open‑source models), and hardening security (CVE fixes, credential leakage prevention). Community contributions are strong, with many first‑time PRs being merged. However, persistent gaps in cross‑platform support (especially Windows) and a growing backlog of critical bugs in larger projects signal a need for disciplined triage and stabilization.

## 2. Activity Comparison

| Project | Issues Updated (24h) | PRs Updated (24h) | Merged/Closed PRs (24h) | New Release (24h) | Health Score |
|---------|----------------------|--------------------|-------------------------|--------------------|-------------|
| **OpenClaw** | 58 (36 open, 22 closed) | 500 (416 open, 84 merged) | 84 | v2026.7.1-beta.1 | B- (large backlog; critical bugs unfixed) |
| **NanoBot** | 3 closed, 21 open | 66 (28 open, 38 merged) | 38 | None | B (rapid fix cadence; some stale issues) |
| **Hermes Agent** | 6 updated | 50 (25 open, 25 merged) | 25 | v0.18.0 / v2026.7.1 | A- (clean merge velocity; minor open P2s) |
| **PicoClaw** | 3 updated | 15 (13 open, 2 merged) | 2 | Nightly v0.3.1-nightly | C (critical Android & Matrix bugs unattended) |
| **NanoClaw** | 1 new issue | 12 (6 open, 6 merged) | 6 | None | B+ (healthy feature delivery; small backlog) |
| **IronClaw** | 5 new issues | 50 (24 open, 26 merged) | 26 | None | B (good merge rate; nightly E2E failing for 36d) |
| **LobsterAI** | 5 stale issues updated | 8 (1 open, 7 merged) | 7 | None | C (3‑month old critical BSOD bug unanswered) |
| **CoPaw** | 10 updated | 50 (18 open, 32 merged) | 32 | v2.0.0-beta.2 | B (good progress; memory leak & streaming lag open) |
| **ZeroClaw** | 12 updated | 50 (40 open, 10 merged) | 10 | None | C+ (p1 bugs unaddressed; high contributor activity) |
| **NullClaw / TinyClaw / Moltis / ZeptoClaw** | 0 | 0 | 0 | None | D (no activity – dormant) |

*Health Score: A = strong merge/issue response, low critical backlog; B = good momentum but some concerns; C = critical bugs or stale items; D = inactive.*

## 3. OpenClaw’s Position

**Advantages:**
- **Scale & maturity**: With 500 daily PR updates and a core reference release, OpenClaw has the largest contributor base and longest history.
- **Broadest channel support**: Integrates Slack, Discord, Mattermost, Matrix, WebChat, iMessage, iOS, Android, and more.
- **Latest LLM support**: Already ships GPT‑5.6 support in the new beta.
- **Rich tool ecosystem**: Includes exec, browser, memory_search, and a plugin system.

**Technical approach:**
- Monolithic Python-based architecture targeting full multi‑channel enterprise deployments.
- Extensive modularity through “harness” and “gateway” abstractions, but the volume leads to high complexity and many inter‑component bugs.

**Community size comparison:**
- OpenClaw’s numbers dwarf others (58 issues, 500 PRs vs. typical 5‑50 for peers). It is the de‑facto “core” project, while projects like NanoBot and Hermes are leaner, more focused alternatives.
- **Weakness**: The backlog of 416 open PRs and multiple long‑standing P1 bugs (some since February) erodes contributor confidence. Users report frequent regressions.

## 4. Shared Technical Focus Areas

Emerging requirements that appear across multiple projects:

| Focus Area | Projects Involved | Specific Need |
|------------|-------------------|---------------|
| **Multi‑channel reliability** | OpenClaw, NanoBot, Hermes, PicoClaw, CoPaw, ZeroClaw | Silent message drops, broken reconnection, lost thread context (Slack, Discord, Feishu, Matrix) |
| **Per‑channel model routing** | OpenClaw (requested), Hermes (merged), NanoBot (cron-level model), ZeroClaw (per‑chat model switching) | Users want different model/provider per channel per conversation |
| **Security hardening** | OpenClaw (CVE GHSA‑cf2p‑f286‑mphf), NanoBot (unauthenticated API), ZeroClaw (config self‑modification, zip‑bomb), CoPaw (secret sanitization) | Provider credential leak prevention, auth enforcement, sandbox enforcement |
| **Windows cross‑platform** | OpenClaw (reported regressions), NanoBot (shell inconsistency), CoPaw (native sandbox PR), ZeroClaw (74 test failures) | Parity for Windows users; current Linux‑only CI gaps |
| **Memory & session state** | OpenClaw (#25592 text leaking, #70024 dead channels), NanoBot (context trimming), IronClaw (idempotency mismatch), CoPaw (vector index explosion, memory leak) | Reliable long‑running sessions, compaction, and consistent turn history |
| **Voice / STT integration** | OpenClaw (English bias), NanoClaw (whisper PR), ZeroClaw (OpenAI STT env fallback), PicoClaw (QQ streaming) | Privacy‑preserving local transcription, streaming voice output |
| **OpenAI‑compatible API endpoint** | ZeroClaw (#8550 RFC), OpenClaw (already has it) | Allow third‑party UIs (Open WebUI) to connect – a major gateway feature |
| **Agent template / repeatable deployment** | NanoClaw (template PRs merged), OpenClaw (per‑agent filesystem roots PR), Hermes (per‑channel config) | Operators want pre‑configured agent groups for fast provisioning |

## 5. Differentiation Analysis

| Project | Primary Focus | Target User | Architecture | Key Differentiator |
|---------|---------------|-------------|--------------|---------------------|
| **OpenClaw** | Full‑featured multi‑channel AI assistant | Enterprise / power users | Python monolith with gateway & harness | “Reference” implementation; broadest LLM & channel coverage |
| **NanoBot** | Lightweight, modular agent core | Developers / tinkerers | Python, modular plugins | Extremely fast iteration (38 PRs merged/day), simplicity |
| **Hermes Agent** | Per‑channel model routing & security | Multi‑channel teams | Python with per‑channel provider config | “The Judgment” release – fine‑grained model selection per channel |
| **PicoClaw** | Embedded / mobile footprint | Raspberry Pi / Termux users | Rust‑based minimal build | ARMv7 target, low resource usage |
| **NanoClaw** | Agent group orchestration & templates | SRE / platform teams | Python with CLI tools | Template‑based agent provisioning, Codex support |
| **IronClaw** | Reborn runtime with WASM sandbox | Advanced developers | Rust + WASM | “Reborn” deterministic runtime, web‑based QA, circuit breaker |
| **LobsterAI** | Chinese‑first agent with Pageant helper | End‑users in China | Python desktop app | Focus on Baidu/NetEase ecosystem, Pageant (Baidu ID) integration |
| **CoPaw** | Qwen‑optimized agent with Feishu focus | Alibaba Cloud / Chinese users | Python with QwenPaw | Built for Alibaba’s Qwen models; strong Feishu (Lark) integration |
| **ZeroClaw** | Rust‑based agent with SOP engine | DevOps / security‑conscious | Rust with SOP scripting | High performance, SOP authoring, OpenTelemetry instrumentation |

## 6. Community Momentum & Maturity

**Tier 1 – Rapidly iterating (≥25 PRs merged/day, new releases or near‑release):**
- **Hermes Agent**: High merge velocity, clean release (v0.18.0), well‑triaged bugs. Most mature release process.
- **CoPaw**: High merge velocity, beta release v2.0.0‑beta.2, active community features. Needs to resolve memory leak before stable.
- **ZeroClaw**: High PR volume, strong contributor base, but two critical p1 bugs linger. Momentum is high but risk is high.
- **NanoBot**: Excellent fix cadence (38 merged/day), batch security fixes. Small team but effective.
- **IronClaw**: Good balance of new features (reborn runtime) and stability fixes; only one major E2E backlog issue.

**Tier 2 – Stabilizing / moderate pace:**
- **OpenClaw**: Very high absolute activity but plagued by backlog and regressions. The beta release indicates stabilization efforts, but 416 open PRs is unsustainable.
- **NanoClaw**: Steady feature delivery; template system is maturing. Low bug volume, but few community comments.

**Tier 3 – Stalled or low engagement:**
- **LobsterAI**: Critical BSOD bug untouched for 3 months; no maintainer response. Low contributor confidence.
- **PicoClaw**: Two critical bugs (Android crash, Matrix reconnection) open for 9+ days without resolution. Nightly build is the only release.
- **NullClaw, TinyClaw, Moltis, ZeptoClaw**: No activity – effectively dormant.

## 7. Trend Signals

The following industry trends emerge from cross‑project community feedback and bug reports, offering actionable insights for AI agent developers:

- **Reliability over novelty**: Users consistently prioritize predictable session-state behavior (no dropped messages, no dead channels) over new model support. Projects that ignore this (e.g., OpenClaw’s persistent “text leaking” bug) lose trust.
- **Cross‑platform is a table‑stakes requirement**: Windows failures (ZeroClaw, NanoBot, OpenClaw) and mobile breakage (OpenClaw iOS, PicoClaw Termux) are top complaints. AI agent developers targeting SaaS or enterprise must invest in CI for all major OSes.
- **Per‑channel intelligence is the next UX frontier**: Hermes already ships it; other projects are building it. Expect “which model for what conversation” to become a standard config knob in 2026.
- **Security is a rising compliance concern**: Multiple CVEs and security advisories (OpenClaw, ZeroClaw, NanoBot) indicate that agent platforms are being audited. Developers should treat credential handling, tool sandboxing, and config validation as first‑class features, not afterthoughts.
- **Observability is underinvested**: Tools that add latency tracing (IronClaw), model‑usage diagnostics (OpenClaw request), and audit logs (ZeroClaw) are being requested. Users want to see *why* an agent behaved a certain way.
- **Agent templates are the new “Dockerfile”**: NanoClaw’s template system (pre‑configured agent groups) mirrors the shift from manual server setup to declarative deployment. Look for similar patterns in other projects as multi‑agent orchestration grows.
- **Local & private voice is a growing niche**: Whisper‑based transcription (NanoClaw) and streaming output (QQ, Telegram) reflect user demand for voice interfaces that don’t rely on cloud APIs. Expect more investment in on‑device STT/TTS.

These signals point to a maturing ecosystem where the differentiators are **reliability, security, and operational simplicity** — not just raw model capability. Developers should prioritize these foundations when selecting or contributing to an open‑source agent platform.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-07-02

## 1. Today’s Overview
NanoBot saw very high development activity today with **66 pull requests updated** in the last 24 hours, of which **38 were merged or closed** — indicating a strong release sprint. **3 issues were closed**, while **21 issues remain open/active**. No new version was cut, but the volume of merged PRs (fixes, tests, refactors) points to a stabilization phase following recent feature additions. The community is actively discussing Windows compatibility, Anthropic OAuth support, and several security/stability bugs that are being addressed in parallel.

---

## 2. Releases
No new releases today.

---

## 3. Project Progress (Merged/Closed PRs)
The 38 merged/closed PRs today include several important fixes and test harnesses:

- **Fix: WebUI subagent backfill payloads** – [#4641](https://github.com/HKUDS/nanobot/pull/4641) hides internal subagent completion payloads from chat refresh (fixes #4640).
- **Test: Runner blocked tool‑call finish reasons** – [#4630](https://github.com/HKUDS/nanobot/pull/4630) adds coverage for `refusal`, `content_filter`, and `error` responses.
- **Test: Cron stale instance mutation consistency** – [#4633](https://github.com/HKUDS/nanobot/pull/4633) regresses against #1033.
- **Fix: Exec early return on session command exit** – [#4643](https://github.com/HKUDS/nanobot/pull/4643) improves performance for short‑lived commands.
- **Fix: Cap local trigger audit records** – [#4642](https://github.com/HKUDS/nanobot/pull/4642) truncates large audit fields and handles filesystem `EINVAL` on directory fsync.
- **Fix: Repeated tool‑result hints** – [#4645](https://github.com/HKUDS/nanobot/pull/4645) adds a model‑facing hint after the third identical text tool output.
- **Refactor: Extract turn‑history recovery** – [#4650](https://github.com/HKUDS/nanobot/pull/4650) moves persistence logic into `nanobot.session.turn_history`.
- **Fix: Hide exec compatibility aliases** – [#4639](https://github.com/HKUDS/nanobot/pull/4639) removes `cmd`, `workdir`, `max_output_tokens` from the tool schema while keeping runtime compatibility.

These PRs collectively improve stability, test coverage, and code maintainability.

---

## 4. Community Hot Topics
- **[#4604 – Feature request: Anthropic OAuth](https://github.com/HKUDS/nanobot/issues/4604)** (4 comments) — The most discussed item today. Users want to use Claude Code tokens generated via `claude setup-token` without an Anthropic Console API key. A corresponding PR [#4632](https://github.com/HKUDS/nanobot/pull/4632) has been opened, suggesting this will land soon.
- **[#4544 – Windows exec: inconsistent shell semantics](https://github.com/HKUDS/nanobot/issues/4544)** (2 comments) — Single‑line commands use `cmd.exe`, multi‑line commands use `powershell`, causing confusion for cross‑platform agent authors. This is a long‑standing paper cut that may be resolved in a future `exec` rewrite.
- **[#4064 – Pending mid‑turn messages lose runtime context](https://github.com/HKUDS/nanobot/issues/4064)** (1 👍) — While only one reaction, the issue touches core agent reliability; a fix is included in the batch PR [#4648](https://github.com/HKUDS/nanobot/pull/4648).

---

## 5. Bugs & Stability
Several bugs updated or reported today, many with fix PRs already in progress:

| Severity | Issue | Summary | Fix Status |
|----------|-------|---------|------------|
| **High** | [#4078](https://github.com/HKUDS/nanobot/issues/4078) | OpenAI‑compatible API accepts unauthenticated requests | Fix in [#4648](https://github.com/HKUDS/nanobot/pull/4648) |
| **High** | [#4076](https://github.com/HKUDS/nanobot/issues/4076) | `message` tool lacks outbound recipient authorization | Fix in [#4648](https://github.com/HKUDS/nanobot/pull/4648) |
| **High** | [#4072](https://github.com/HKUDS/nanobot/issues/4072) | ExecTool symlink bypasses restricted workspace | Fix in [#4629](https://github.com/HKUDS/nanobot/pull/4629) |
| **High** | [#4075](https://github.com/HKUDS/nanobot/issues/4075) | Dream can overwrite user‑created skills without ownership check | Fix in [#4648](https://github.com/HKUDS/nanobot/pull/4648) |
| **Medium** | [#4061](https://github.com/HKUDS/nanobot/issues/4061) | OpenAI‑compatible text‑format tool calls not parsed | Fix in [#4648](https://github.com/HKUDS/nanobot/pull/4648) |
| **Medium** | [#4068](https://github.com/HKUDS/nanobot/issues/4068) | Matrix stream buffer keyed only by chat_id (corruption) | Fix in [#4648](https://github.com/HKUDS/nanobot/pull/4648) |
| **Medium** | [#4055](https://github.com/HKUDS/nanobot/issues/4055) | Dream compaction deletes unprocessed history entries | Fix in [#4648](https://github.com/HKUDS/nanobot/pull/4648) |
| **Low** | [#4637](https://github.com/HKUDS/nanobot/issues/4637) | Telegram long message split — prior trunks cannot render | No fix PR yet |
| **Low** | [#4619](https://github.com/HKUDS/nanobot/issues/4619) | Feishu new‑session message lacks visual separator | No fix PR yet |

The batch fix PR [#4648](https://github.com/HKUDS/nanobot/pull/4648) addresses many of the older bugs (from May), indicating a strong push to close the security and reliability gap.

---

## 6. Feature Requests & Roadmap Signals
- **Anthropic OAuth** ([#4604](https://github.com/HKUDS/nanobot/issues/4604)) — The most likely next feature; PR [#4632](https://github.com/HKUDS/nanobot/pull/4632) is already open. Expect it in the next release.
- **Trigger agent from external script** ([#4605](https://github.com/HKUDS/nanobot/issues/4605) — closed) — A user asked for a programmatic hook to launch agent actions (e.g., from Gmail filtering). The issue was closed without action, but the discussion signals demand for an external API or webhook endpoint.
- **Cron‑level model/preset** ([#4378](https://github.com/HKUDS/nanobot/issues/4378)) — Users want per‑cron‑job model overrides. Currently no PR, but a workaround (cron job calling a model‑switching script) was suggested. May be deferred.
- **Optional feature enablement** — PR [#4396](https://github.com/HKUDS/nanobot/pull/4396) (open) adds a plugin‑style system for enabling optional capabilities (Bedrock extra, WebUI features). This architectural change could appear in the next minor release.

---

## 7. User Feedback Summary
- **Positive**: A new user in [#4605](https://github.com/HKUDS/nanobot/issues/4605) praised the lightweight codebase compared to OpenClaw and expressed enjoyment learning the source. They were able to set up a Gmail skill successfully.
- **Pain points**:
  - Windows users continue to face inconsistent shell behavior ([#4544](https://github.com/HKUDS/nanobot/issues/4544)).
  - Telegram users report poor rendering of long markdown messages ([#4637](https://github.com/HKUDS/nanobot/issues/4637)).
  - Feishu (Lark) users want a cleaner visual separator for new sessions ([#4619](https://github.com/HKUDS/nanobot/issues/4619)).
- **Satisfaction**: The team’s rapid response to reported bugs (e.g., the fsync crash [#4615](https://github.com/HKUDS/nanobot/issues/4615) fixed in one day) is reassuring to the community.

---

## 8. Backlog Watch
Several long‑standing issues (from late May) have remained open with minimal maintainer attention but now have fix PRs in progress:

- [#4056](https://github.com/HKUDS/nanobot/issues/4056) – Context trimming can drop the assistant question before user reply (no fix PR yet).
- [#4058](https://github.com/HKUDS/nanobot/issues/4058) – Tool‑result protocol repair allows missing/duplicate `tool_call_id` (fix in [#4648](https://github.com/HKUDS/nanobot/pull/4648)).
- [#4062](https://github.com/HKUDS/nanobot/issues/4062) – WebSocket drops proactive messages with no subscribers (fix in [#4648](https://github.com/HKUDS/nanobot/pull/4648)).
- [#4067](https://github.com/HKUDS/nanobot/issues/4067) – Invalid config silently falls back to defaults (explicitly excluded from [#4648](https://github.com/HKUDS/nanobot/pull/4648); still open).
- [#4082](https://github.com/HKUDS/nanobot/issues/4082) – Cron jobs reuse fixed session context across runs (explicitly excluded from [#4648](https://github.com/HKUDS/nanobot/pull/4648); still open).
- [#4136](https://github.com/HKUDS/nanobot/issues/4136) – Refactor session retention result API (fix in [#4648](https://github.com/HKUDS/nanobot/pull/4648)).

Issues [#4067](https://github.com/HKUDS/nanobot/issues/4067) and [#4082](https://github.com/HKUDS/nanobot/issues/4082) remain unaddressed and could cause silent misbehavior in production. Maintainers should prioritize them alongside the incoming batch fix.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-07-02

## Today’s Overview
The project saw heavy activity: **50 pull requests** and **6 issues** were updated in the last 24 hours, alongside the release of **v0.18.0** (“The Judgment Release”). Half of the PRs were merged or closed, reflecting rapid iteration. Community engagement remains high with **370+ contributors** to the v0.18.0 milestone. Major themes include per-channel model routing, security hardening for gateways, and desktop integration improvements. A significant batch of older PRs from March were finally merged, indicating a backlog unblocked.

## Releases
**v2026.7.1 / v0.18.0 (Hermes Agent — The Judgment Release)**  
- **Scope**: ~1,720 commits, 998 merged PRs, 2,215 files changed (~251,000 additions, ~41,000 deletions), 949 issues closed.  
- **No explicit breaking changes or migration notes** are provided in the release snippet. Given the volume of changes, users upgrading from v0.17.0 should consult the full changelog for any config or API deprecations.  
- Noteworthy: the release title “Judgment” aligns with the many feature and security improvements merged in this cycle.

## Project Progress
**25 PRs were merged or closed today** (out of 50 updated). Key advances include:

- ✅ **Per-channel model & system prompt overrides** — #1991 (fixed #1955) and its salvage #56967 are now closed, giving Discord/Telegram/Slack channels individual model/provider/instruction configs.
- ✅ **Separate chat model for interactive messages** — #2840 merges the ability to use a different model for real‑time conversations vs. cron/background tasks.
- ✅ **Least‑privilege toolset for SMS** — #2860 restricts SMS gateway to a safer tool subset.
- ✅ **Email adapter hardening** — #2794 and its salvage #56956 fix IMAP malformed‑response crashes.
- ✅ **Teams security fix** — #56968 rejects unauthorized senders before attachment processing.
- ✅ **QQBot reconnect fix** — #56970 adds missing `is_reconnect` kwarg.
- ✅ **Kasia gateway integration** — #2707 brings encrypted peer‑to‑peer messaging support.
- ✅ **Env flag consistency** — #2863 accepts `on` as truthy for boolean environment variables.

Several of these PRs were originally created in March and only now merged, demonstrating a successful “salvage” effort to rebase and land valuable work.

## Community Hot Topics
- **#1955 / #1991 — Per‑channel model overrides**  
  *Issue* (10 comments, 7 👍) / *PR* — This was the most discussed feature request. Users needed different models for #daily (cheap) vs. #dev (capable) channels. The closure of both issue and PR today satisfies a long‑standing pain point.  
  [Issue #1955](https://github.com/NousResearch/hermes-agent/issues/1955) | [PR #1991](https://github.com/NousResearch/hermes-agent/pull/1991)

- **#54748 — Model routing tiers** (open, 0 comments)  
  A config‑driven routing layer for fallback chains. Although not heavily commented, it touches the same model selection space and could reduce complexity for multi‑provider setups.  
  [PR #54748](https://github.com/NousResearch/hermes-agent/pull/54748)

- **#46466 — Desktop plugin loader** (open, 0 comments)  
  Brings web dashboard plugins (kanban, achievements) into the native Electron app. Signals community desire for a richer desktop experience.  
  [PR #46466](https://github.com/NousResearch/hermes-agent/pull/46466)

## Bugs & Stability
| Severity | Bug | Status | Fix PR? |
|----------|-----|--------|---------|
| **P2** | **Desktop model picker shows implicitly discovered providers** (#56974) | Open | #56966 (open) |
| **P2** | **Docker exec runs as root → fleet 401s** (#56942) | Closed (fixed) | (fix in image) |
| **P2** | **Reasoning‑model responses broken in `_extract_text`** (#56948) | Open | #56948 (open) |
| **P2** | **MCP approval tools no‑op** (#56971) | Open | #56971 (open) |
| **P2** | **Web extract final URL re‑check missing** (#56972) | Open | #56972 (open) |
| **P3** | **Email adapter malformed IMAP crash** (#2794) | Closed (fixed) | #2794 / #56956 |
| **P3** | **Teams attachment bypass** (#56968) | Closed (fixed) | #56968 |
| **P3** | **QQBot reconnect failure** (#56970) | Closed (fixed) | #56970 |

The most impactful P2 bug (Docker root) was fixed. Remaining open P2 issues have associated PRs in review or draft.

## Feature Requests & Roadmap Signals
- **Cron output with model/provider metadata** (#56791, P3, open) — users want to attribute cron runs to specific models. Likely to land soon given PR #56973 is open.
- **Pre‑execution hook for URL routing** (#56969, P3, open) — would allow tool dispatch to be aware of content before execution. Signals deeper hook infrastructure.
- **Native image/video generation** (#56965, P3, open) — users with multimodal endpoints want to bypass third‑party image APIs. Suggests growing demand for “model native” generation.
- **Model routing tiers** (#54748, P3, open) — if merged, the next release could include fallback chains per task tier.

These align with the release’s focus on configurable model selection and security hardening.

## User Feedback Summary
- **Satisfaction**: The closure of the long‑standing per‑channel override request (#1955) has high reaction score (7 👍). Users also praise the multiple salvage PRs that unblocked old contributions.
- **Pain points**:
  - **Desktop model picker confusing** (#56974) — users see providers they never configured.
  - **Plugin developers** hit reasoning‑model incompatibility (#56948) and missing pre‑execution hooks (#56969).
  - **Security‑conscious users** welcome the Teams, SMS, and web‑extract fixes but demand more consistent URL validation (#56972).
- **Use cases** continue to center on multi‑channel, multi‑model deployments and desktop‑first management.

## Backlog Watch
Several older PRs (March 2026) were finally merged today, clearing a significant queue. Still open and needing attention:

- **#54748** (model routing tiers, June 29) — no comments from maintainers yet; potential for inclusion in v0.19.
- **#46466** (desktop plugin loader, June 15) — no maintainer response; may need rebasing.
- **#56948** (reasoning‑model fix) — has an open PR but no review; could block plugin LLM adoption.
- **#56971** (MCP approval handshake) — open PR touching security boundary; should be prioritized.

No issue older than 3 weeks remains unattended, indicating healthy maintainer engagement.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

**PicoClaw Project Digest – 2026-07-02**

**1. Today’s Overview**  
Activity is moderate, with 3 issues and 15 pull requests updated in the last 24 hours. Two PRs were merged/closed, the nightly release pipeline produced a new unstable build, and two new bug reports surfaced around mobile compatibility and Matrix reliability. The project is clearly under active development, with a healthy mix of bug fixes, platform expansions, and dependency upkeep, though the number of stale items (7+ days without update) signals a need for maintainer review.

**2. Releases**  
- **Nightly Build – v0.3.1-nightly.20260702.2cf030d2**  
  *Full changelog:* `v0.3.1 → main`  
  This is an automated, potentially unstable build with no explicit feature or breaking-change notes. Users should treat it as a development snapshot only. No migration steps required.

**3. Project Progress**  
Two PRs were merged/closed today, both originating from the community:  
- **PR #3116** (closed; fix) – `fix(pico): complete turn.done lifecycle signaling`  
  Preserves `request_id` for queued steering and follow-up messages, ensuring every inbound Pico request receives a proper lifecycle signal. Closes issue #2984.  
- **PR #2975** (closed; feature) – `feat(telegram): treat reply to bot message as mention in group chats`  
  When `mention_only: true` is configured, Telegram group chats now interpret replying to a bot as equivalent to @-mentioning, reducing friction for users.

Additionally, two new PRs opened today are progressing toward merge:  
- **PR #3205** – Adds Linux ARMv7 build target and fixes 9router response parsing for Raspberry Pi users.  
- **PR #3204** – Restores Azure SDK dependency freeze baseline to align with downstream supply-chain checks.

**4. Community Hot Topics**  
- **Issue #3164** – `[BUG] Process hooks crash gateway on Android/Termux` (opened 2026-06-23, 1 comment, still open).  
  A self-identified user reports that even a minimal “hello world” JSON-RPC hook crashes the gateway within 2 seconds on Android. This has been open for 9 days without resolution, indicating a significant blocker for mobile users.  
  *Link:* https://github.com/sipeed/picoclaw/issues/3164

- **Issue #3203** – `[BUG] Matrix sync loop has no reconnection logic` (opened today, 0 comments).  
  Reports silent death after network/server disruption. No workaround exists. This is a reliability-critical issue for Matrix users.  
  *Link:* https://github.com/sipeed/picoclaw/issues/3203

- **Issue #3201** – `[Feature] Support streaming output for QQ channel` (opened 2026-07-01, 0 comments).  
  User requests token-by-token streaming to match Telegram and WebSocket channels. This feature gap is highlighted as a quality-of-life improvement.  
  *Link:* https://github.com/sipeed/picoclaw/issues/3201

**5. Bugs & Stability**  
| Severity | Bug Description | Issue/PR | Fix Available? |
|----------|----------------|----------|----------------|
| **Critical** | Process hooks crash gateway on Android/Termux (#3164) – no recoverable workaround. | Issue #3164 | No open PR |
| **High** | Matrix sync loop dies silently after network disruption; systemd doesn’t restart (#3203). | Issue #3203 | No open PR |
| **Medium** | CLI tool calls with invalid JSON arguments cause batch drop (#3180). | PR #3180 (open, proposed fix) | PR #3180 |
| **Medium** | 9router gateway responses not parsed on Raspberry Pi (#3205). | PR #3205 (open, proposed fix) | PR #3205 |
| **Low** | Leading/trailing underscores not stripped in ID normalization (#3202). | PR #3202 (open) | PR #3202 |

No regressions were reported today.

**6. Feature Requests & Roadmap Signals**  
- **QQ channel streaming (#3201)** – Following Telegram and WebSocket, adding streaming to QQ is a clear community demand. Likely target: v0.4.0.  
- **DeltaChat gateway (#3063)** – Open PR adding a new chat platform integration. Signals interest in federated / email-based messengers.  
- **Bedrock prompt caching (#3163)** – Open PR leveraging AWS Converse API cache points. Shows enterprise/cloud optimisation push.  
- **9router gateway support (#3205)** – Not a feature request per se, but the fix indicates users are experimenting with alternative OpenAI-compatible backends.

Predictions: The next minor release (v0.3.2 or v0.4.0) will likely include QQ streaming (if merged) and the Linux ARMv7 build target.

**7. User Feedback Summary**  
- **Pain points:**  
  - Android/Termux users cannot run even basic hooks (#3164) – major dissatisfaction.  
  - Matrix channel is unreliable in production (#3203) – can cause silent outages.  
  - QQ users frustrated by waiting for full response instead of streaming (#3201).  
- **Positive signals:** Two PRs merged today directly address user-reported issues (lifecycle completeness and Telegram group usability), showing that maintainers respond to community needs.

**8. Backlog Watch**  
Several older, high-value items remain unattended:  
- **Issue #3164** (Android crash) – 9 days stale, no fix PR. Needs urgent triage.  
- **PR #3165** (fix openai_compat Seed XML tool calls) – Open since 2026-06-24, no recent activity. Important for Volcengine users.  
- **PR #3161** (fix exec deny patterns) – Stale for 9 days; security-relevant.  
- **PR #3160** (fix auth cross-site requests) – Also stale and security-critical.  
- **dependabot PRs** (#3104, #3103, #3100) – All stale for >20 days; should be merged or closed to reduce risk.

These backlog items, if left unresolved, may slow down the project’s stability and user trust. Maintainers should prioritise the Android crash and the security fixes.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest – 2026-07-02

## 1. Today’s Overview

NanoClaw saw a burst of development activity in the past 24 hours, with 12 pull requests updated (6 closed/merged, 6 still open) and a single new issue filed. The project is clearly in the midst of a major feature rollout: the **agent template system** (parts 1 and 2) is nearing completion, with two stacked PRs (#2890, #2909) driving the local template loader, group stamping, and setup-wizard integration. A supporting PR (#2908) makes templates work end-to-end under the Codex provider. Meanwhile, critical stability fixes for WhatsApp reconnection and Slack thread handling were merged or are in review. No new releases were cut on this date. Overall, the project is healthy and actively shipping both new capabilities and quality-of-life improvements.

## 2. Releases

_No new releases were published on or near this date._

## 3. Project Progress

Six PRs were **merged or closed** in the last 24 hours, covering fixes and new skills:

- **[#2905 – fix(whatsapp): end the old socket on reconnect (host memory leak)](https://github.com/nanocoai/nanoclaw/pull/2905)**  
  _Merged._ Solves a memory leak caused by orphaned Baileys socket graphs on frequent WhatsApp reconnects. Each reconnect now properly closes the previous socket, preventing resource exhaustion.

- **[#2677 – fix(scheduling): retry pre-task script once on failure with diagnostics](https://github.com/nanocoai/nanoclaw/pull/2677)**  
  _Merged._ Adds a single retry with diagnostic output when a pre-task script fails, improving scheduling reliability.

- **[#1716 – feat: add /check-contribution operational skill for PR pre-flight checks](https://github.com/nanocoai/nanoclaw/pull/1716)**  
  _Merged._ A new operational skill that automates SKILL.md validation, type consistency, code quality, and PR readiness checks for contributors.

- **[#1257 – feat: support custom API endpoints (e.g., z.ai)](https://github.com/nanocoai/nanoclaw/pull/1257)**  
  _Merged._ Allows use of Anthropic-compatible APIs mounted at a sub-path (e.g., z.ai), expanding provider flexibility.

- **[#1693 – feat: add /add-backup utility skill for automated state backup](https://github.com/nanocoai/nanoclaw/pull/1693)**  
  _Merged._ Introduces automated, channel-agnostic backups of NanoClaw state (messages, groups, sessions, IPC) to a local git repo with optional remote push.

- **[#1597 – feat: add QMD skill for semantic conversation search](https://github.com/nanocoai/nanoclaw/pull/1597)**  
  _Merged._ A new feature skill enabling semantic search across conversations using QMD (Query-Model-Document) methods.

These merges represent significant progress in stability, developer tooling, provider support, and data management.

## 4. Community Hot Topics

No issue or PR attracted more than a single reaction or comment in the data set. However, the most **structurally active** items are the three template-related PRs that are stacked together:

- **[#2890 – feat(templates): local template loader, ncl --template, and docs](https://github.com/nanocoai/nanoclaw/pull/2890)** (OPEN) – Part 1 of 2, adds the local template loader and the `ncl groups create --template` command.
- **[#2909 – feat(setup): template setup flow in the wizard and first-agent stamping](https://github.com/nanocoai/nanoclaw/pull/2909)** (OPEN, draft) – Part 2, builds the setup-wizard flow for selecting and stamping templates.
- **[#2908 – feat(codex): persona prepend + git-independent skill discovery for template agents](https://github.com/nanocoai/nanoclaw/pull/2908)** (OPEN) – Complements the template system by enabling Codex provider agents to use templates through persona prepending and a provider-agnostic skills mirror.

These three PRs form a coherent feature push and are likely the most anticipated by the community. The underlying need is **operational agility**: operators want to spin up agent groups from predefined templates without manual configuration.

## 5. Bugs & Stability

- **HIGH – WhatsApp reconnect memory leak** – Fixed in [#2905](https://github.com/nanocoai/nanoclaw/pull/2905) (merged). Previously, frequent reconnects (triggered by `reason: 408`) accumulated orphaned socket graphs and timers, leading to host memory exhaustion. This was a critical issue for production deployments using WhatsApp.

- **MEDIUM – Slack thread history not reloaded on @mention** – Addressed in open PR [#2904](https://github.com/nanocoai/nanoclaw/pull/2904). In `engage_mode: 'mention'`, re-tagging the bot deep in an existing thread only delivered the single tagged message, making previous human replies invisible. The fix reloads the full thread from the platform.

- **LOW – Scheduling pre-task script failures** – Improved with a retry and diagnostic logging in [#2677](https://github.com/nanocoai/nanoclaw/pull/2677) (merged). Not a bug per se, but enhances robustness.

No new bugs were reported in the single open issue (#2907), which appears to be a feature request or meta issue (title "_ape_claw_cli_").

## 6. Feature Requests & Roadmap Signals

Several new features are actively in development or were recently merged, signalling the project’s roadmap priorities:

- **Agent templates** (PRs #2890, #2909, #2908) – A major new capability to stamp agent groups from templates. The setup wizard will guide first-time stamping, and the Codex provider gains full support. This points toward **simplified onboarding** and **repeatable agent configurations**.

- **Instance-wide default agent provider** ([#2906 – OPEN](https://github.com/nanocoai/nanoclaw/pull/2906)) – Lets operators set `DEFAULT_AGENT_PROVIDER` in a `.env` file, so new groups automatically use the configured provider (e.g., `claude`). Reduces per-group manual setup. Likely to be merged soon.

- **Free local voice transcription** ([#2317 – OPEN](https://github.com/nanocoai/nanoclaw/pull/2317)) – A skill to wire local voice transcription using `openai-whisper` or `whisper.cpp`. Pending since May, but recently updated. Suggests community interest in privacy-preserving voice features.

- **Backup automation** ([#1693 – merged](https://github.com/nanocoai/nanoclaw/pull/1693)) and **semantic search** ([#1597 – merged](https://github.com/nanocoai/nanoclaw/pull/1597)) are now part of the core skill library.

- **Custom API endpoints** ([#1257 – merged](https://github.com/nanocoai/nanoclaw/pull/1257)) enables compatibility with third-party Anthropic API proxies, a common need for self-hosted users.

**Prediction**: The next version (likely v0.x) will include the full template system, the default provider setting, and the local transcription skill. The template feature is clearly the most complex and is being delivered in two parts.

## 7. User Feedback Summary

Direct user feedback is sparse in this data set. The only open issue is #2907 titled "_ape_claw_cli_" with no description, thumbs, or comments – likely a placeholder or user request for a CLI wrapper. No complaints or satisfaction signals are visible.

However, the PR activity reveals several implicit user pain points that the maintainers are addressing:

- **Memory leaks in WhatsApp integrations** – The fix in #2905 indicates that real-world deployments suffer from repeated reconnections, a problem that affected stability for users relying on WhatsApp as a chat channel.
- **Slack thread context loss** – The @mention fix (#2904) shows that users engaging the bot in existing threads expected the bot to see the full conversation history, not just the latest message.
- **Difficulty managing provider configuration** – The default provider PR (#2906) responds to the friction of having to set the provider for every new agent group.
- **Desire for repeatable setup** – The template system directly addresses operators who want to avoid manual per-group configuration.

Overall, the community appears satisfied with the pace of development, as evidenced by the rapid merge of several community-contributed skills (backup, semantic search, check-contribution).

## 8. Backlog Watch

- **[#2317 – feat(skills): add /add-voice-transcription-free-whisper](https://github.com/nanocoai/nanoclaw/pull/2317)** – Open since **2026-05-07** (nearly 2 months). Last updated 2026-07-01. This is a substantial skill addition that has been awaiting integration. It may be blocked by review capacity or feature completeness. Maintainers should prioritize merging it to avoid stagnation.

- **[#2904 – fix(slack): reload thread history from platform on @mention](https://github.com/nanocoai/nanoclaw/pull/2904)** – Open since 2026-07-01, updated same day. No comments. This fix addresses a real user-facing bug and should be reviewed promptly.

- **[#2906 – feat: instance-wide default agent provider for new groups](https://github.com/nanocoai/nanoclaw/pull/2906)** – Open since 2026-07-01, also no comments. While not old, it is a small, clearly scoped feature that could be merged quickly after review.

No old issues or PRs appear to have been abandoned; the project’s maintainers are actively handling incoming work. The single open issue (#2907) is brand new and uncommented.

**Action item**: Consider prioritizing review of #2317 (voice transcription) to avoid a multi-month backlog, and provide feedback or merge #2904 and #2906 as they are low-risk improvements.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-07-02

## 1. Today’s Overview

High activity in the IronClaw repository: 50 pull requests were updated in the last 24 hours, with 26 merged or closed and 24 still open. Five new issues were filed, all active and none closed. No new releases were cut. The bulk of merged work focuses on the “reborn” runtime: fixes for skill-learned bubble delivery after SSE resume, auth-resume approval surfacing, OAuth refresh guidance, and a circuit breaker for fast‑failing degraded LLM providers. Several design‑doc‑only PRs (`#5529`, `#5249`) lay architectural groundwork for future crate decomposition and lease durability. A cluster of bugs uncovered by production‑faithful tests (e.g., idempotency write/read mismatch, unreachable skill auto‑activation) points to ongoing reliability hardening.

## 2. Releases

No releases were published in the last 24 hours. The last release PR (`#5311`) remains open.

## 3. Project Progress (Merged/Closed PRs Today)

**26 PRs were merged or closed today.** Key advances include:

- **Reborn runtime improvements** (multiple PRs by `serrrfirat`):
  - `#5407` — Fix skill-learned bubble dropped on SSE resume.
  - `#4998` — Surface approval gate after auth resume, preventing terminal authorization failures.
  - `#5054` — Guide Google OAuth refresh setup for offline access.
  - `#5203` — Enable LLM circuit breaker by default to fast‑fail degraded providers.
  - `#5228` — Decouple runner heartbeat timeout from write timeout to avoid operator‑forced slowdowns.
  - `#5238` — Suppress Reborn debug log noise from `cranelift_codegen` / `wasmtime_internal_cranelift`.
  - `#4911` — Keep prior approval metadata atomic when flowing into auth resume.
  - `#4145` — Anchor turn objective in loop context for prompt transcript handling.

- **Secrets management**:
  - `#2754` — Add self‑service user secrets UI and durable binding approvals (large feature, scoped PR merged).

- **Scoped lifecycle storage**:
  - `#4544` — Foundation for admin‑installed skill/tool availability dimension.

- **Test & CI**:
  - `#5531` — Add LLM semantic judge fallback for Reborn WebUI v2 live QA.
  - `#5532` — Wire codebase knowledge graph + OpenWiki for agent code discovery.

- **Documentation & design**:
  - `#5529` — Final crate/module refactor design for reborn stack (docs only).
  - `#5249` — Design doc for write‑behind lease durability and side‑effect gate.

## 4. Community Hot Topics

The most discussed items by comment count and immediate user impact:

- **Issue #5459 — Configurable skills and tools** (2 comments)  
  Asks for admin‑ vs user‑scoped install of WASM tools and skills. This reflects a clear user need for access control granularity in multi‑tenant deployments.  
  [GitHub Issue #5459](https://github.com/nearai/ironclaw/issues/5459)

- **Issue #5522 — [QA] Reborn routine fails reading Slack DMs — missing capability + retry loop**  
  A reported production blocking scenario: the agent repeatedly retries `capability_info` when Slack DM read is not enabled. High user dissatisfaction.  
  [GitHub Issue #5522](https://github.com/nearai/ironclaw/issues/5522)

- **PR #5534 — Instrument first‑party tool latency** (open, large)  
  Adds latency tracing around runtime dispatch and egress calls. Likely motivated by live trace gaps and performance monitoring needs.  
  [GitHub PR #5534](https://github.com/nearai/ironclaw/pull/5534)

- **PR #4841 — Reborn: no run‑borking failures** (open, large)  
  Aims to eliminate terminal “run‑borking” failures by converting every run‑terminal error into a recoverable or explained state. Community interest in stability is high.  
  [GitHub PR #4841](https://github.com/nearai/ironclaw/pull/4841)

## 5. Bugs & Stability

Four new bugs reported today, ranked by severity:

| Severity | Issue | Description | Fix PR? |
|----------|-------|-------------|---------|
| **High** | #5522 | Reborn routine fails with infinite retry when reading Slack DMs; missing Slack read capability. | No fix PR yet; requires capability policy update. |
| **High** | #5527 | `FilesystemSessionThreadService` idempotency write (owner scope) vs read (system scope) never coincide — early replay‑before‑policy dead in production. | Found via PR #5526; fix expected in that line. |
| **Medium** | #5530 | Skill criteria‑based auto‑activation unreachable from modern submit path (TurnCoordinator). | Dead code; needs re‑routing to `SelectableSkillContextSource`. |
| **Medium** | #4108 | Nightly E2E scheduled run has been failing since May 27 (updated again today). | No root cause posted; workflow run link available in issue. |

Additionally, the recent merge of `#5407` addresses a skill‑bubble delivery bug that was previously causing test failures on `main`.

## 6. Feature Requests & Roadmap Signals

- **Configurable skills & tools** (`#5459`) — Admin vs user scoped installation of WASM tools and skills. This is a high‑demand feature for multi‑tenant and enterprise deployments.
- **Codebase knowledge graph + OpenWiki** (`#5532`, merged) — Adds auto‑maintained code discovery for agents; likely to reduce agent confusion during development.
- **Tool latency instrumentation** (`#5534`, open) — Signals a push toward observability and performance diagnostics.
- **Run‑borking elimination** (`#4841`, open) — If accepted, would markedly improve user trust in long‑running agent sessions.

**Prediction**: The next minor release (`0.29.x` or `0.30.0`) will likely include the configurable skills foundation from `#5459` (once resolved) and the circuit‑breaker, heartbeat decoupling, and auth‑resume fixes that have already been merged.

## 7. User Feedback Summary

Real pain points surfaced in today’s issues:

- **Slack DM integration broken** (`#5522`): Agent fails when task requires reading Slack DMs, then enters a retry loop. The user expects a clear error or automatic capability grant, not a stuck routine.
- **Session token staleness** (addressed by PR `#5511`, open): WebUI SSO tokens became stale without proper relogin; the fix allows OAuth callback to replace old tokens.
- **Unclear failure mode on model outage** (PR `#5203`, merged): Users previously experienced frozen instances during NEAR AI outages; the circuit breaker now surfaces “model unavailable” quickly.
- **Lost skill‑learned notifications** (PR `#5407`, merged): After SSE resume, the “learned a new skill” bubble was dropped – a confusing UX gap.

Satisfaction is likely improving with the rapid fix cadence, but the new bug reports indicate production gaps remain.

## 8. Backlog Watch

| Issue | Age | Status | Notes |
|-------|-----|--------|-------|
| **#4108** — Nightly E2E failure | Opened 2026‑05‑27 (36 days) | Open, updated today | No maintainer comment; a persistent stability blocker. |
| **#5459** — Configurable skills/tools | Opened 2026‑06‑30 (2 days) | Open, 2 comments | Fresh but already the most‑commented issue today. |
| **#5530, #5527, #5522** | All opened today | Open | Early‑stage; should be triaged. |

The most concerning backlog item is **#4108** — a nightly E2E suite that has been failing for over a month without a documented root cause or assigned owner. It may indicate a deeper systemic issue that will affect release quality.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest – 2026-07-02

## 1. Today’s Overview
Project activity is **moderate** with 8 PRs updated in the last 24 hours (7 merged/closed) and 5 open issues receiving maintenance touches. No new releases were published. The team focused on squashing two renderer‑side bugs (scheduled‑task notification persistence and a white‑screen crash on custom‑model deletion) and merged a feature adding a subagent artifact panel. Meanwhile, five stale issues from early April remain open with no maintainer response, indicating a possible backlog in user‑reported bug triage.

## 2. Releases
*No new releases were published today.*

## 3. Project Progress
**Merged/closed PRs today** (7 items):

| PR | Description |
|----|-------------|
| [#2252](https://github.com/netease-youdao/LobsterAI/pull/2252) | **Fix (settings):** Prevent white screen when deleting active custom provider. |
| [#2255](https://github.com/netease-youdao/LobsterAI/pull/2255) | **Fix (scheduled tasks):** “不通知” notification channel now correctly takes effect after save. |
| [#2251](https://github.com/netease-youdao/LobsterAI/pull/2251) | **Fix (deployment):** Use isolated Node environment for executing deploy commands; improved error messages for missing tools. |
| [#2249](https://github.com/netease-youdao/LobsterAI/pull/2249) | **Feature (cowork):** Add subagent artifact panel with list/detail views; open subagent details in right‑side panel. |
| [#2254](https://github.com/netease-youdao/LobsterAI/pull/2254), [#2253](https://github.com/netease-youdao/LobsterAI/pull/2253), [#2250](https://github.com/netease-youdao/LobsterAI/pull/2250) | **Documentation:** Update main page image and README files. |

**Still open:** [#2256](https://github.com/netease-youdao/LobsterAI/pull/2256) combines the same two fixes (scheduled‑task notification + settings model‑delete) as a single squashed PR.

## 4. Community Hot Topics
Activity on issues is low—all five updated today are **stale** (created April 2) and carry only 1–2 comments with zero reactions. The most discussed issue is:

- **[#1354](https://github.com/netease-youdao/LobsterAI/issues/1354)** – “Blue screen after launching Pageant” (2 comments). The user reports a crash with attached logs; no maintainer reply yet.

Other issues with one comment each:
- [#1357](https://github.com/netease-youdao/LobsterAI/issues/1357) – Pageant not actually started despite affirmative response.
- [#1358](https://github.com/netease-youdao/LobsterAI/issues/1358) – Timed tasks have no visual feedback on activation.
- [#1359](https://github.com/netease-youdao/LobsterAI/issues/1359) – Deleted tasks reappear after restart.
- [#1360](https://github.com/netease-youdao/LobsterAI/issues/1360) – No duplicate name validation when creating custom agents.

**Underlying need:** Users are experiencing reliability and UX gaps in agent lifecycle (startup, deletion, feedback) and authentication helpers (Pageant). The lack of maintainer responses may erode trust.

## 5. Bugs & Stability
**Bugs reported today (all stale, but updated today):**

| Severity | Issue | Summary |
|----------|-------|---------|
| **Critical** | [#1354](https://github.com/netease-youdao/LobsterAI/issues/1354) | Blue screen of death (BSOD) after using Lobster to launch Pageant. |
| **High** | [#1357](https://github.com/netease-youdao/LobsterAI/issues/1357) | “帮我开启pageant” responds as started but Pageant is not actually running. |
| **Medium** | [#1359](https://github.com/netease-youdao/LobsterAI/issues/1359) | Deleted tasks reappear every restart. |
| **Low** | [#1360](https://github.com/netease-youdao/LobsterAI/issues/1360) | Duplicate agent names can be created without validation. |
| **Low** | [#1358](https://github.com/netease-youdao/LobsterAI/issues/1358) | No interaction feedback when clicking a timed task. |

**Fixes in today’s PRs:** Two bugs were resolved:
- White‑screen crash on deleting active custom model ([#2252](https://github.com/netease-youdao/LobsterAI/pull/2252))
- Scheduled task “不通知” channel not persisting ([#2255](https://github.com/netease-youdao/LobsterAI/pull/2255))

Only one PR (#2256) is still open and incorporates both fixes.

## 6. Feature Requests & Roadmap Signals
No explicit feature requests were filed today. However, the merged PR **[#2249](https://github.com/netease-youdao/LobsterAI/pull/2249)** adds a **subagent artifact panel** in the cowork area—a clear roadmap signal that the team is investing in multi‑agent collaboration UX.

The stale issue [#1360](https://github.com/netease-youdao/LobsterAI/issues/1360) (duplicate agent names) reflects a demand for **input validation** in agent creation forms. Such a feature is low‑effort and could be expected in the next patch release.

## 7. User Feedback Summary
User pain points visible in today’s issues:

- **Stability**: A user reports a BSOD (blue screen) when using Pageant integration—this is the most severe complaint.
- **Misleading feedback**: “帮我开启pageant” says it started but didn’t. Users want accurate status confirmation.
- **Persistence bugs**: Deleted tasks resurrecting on restart is confusing and undermines trust.
- **Missing validation**: Duplicate agent names cause ambiguity.
- **Affordance**: Timed tasks give no visual indication of being active.

Satisfaction is low for these users—they are filing bugs with logs and screenshots but receiving no public replies.

## 8. Backlog Watch
Five stale issues from **April 2, 2026** remain open with no maintainer intervention:

- [#1354](https://github.com/netease-youdao/LobsterAI/issues/1354) – BSOD (critical, 2 comments)
- [#1357](https://github.com/netease-youdao/LobsterAI/issues/1357) – Pageant false start (high)
- [#1358](https://github.com/netease-youdao/LobsterAI/issues/1358) – No timed‑task feedback (low)
- [#1359](https://github.com/netease-youdao/LobsterAI/issues/1359) – Deleted tasks reappear (medium)
- [#1360](https://github.com/netease-youdao/LobsterAI/issues/1360) – Duplicate agent name (low)

These have been untouched for **3 months**. The team should prioritize triaging them, especially the BSOD report, before the next release to avoid accumulation of unresolved crashes.

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

# CoPaw Project Digest — 2026-07-02

## 1. Today's Overview
CoPaw saw high activity on July 2, with 50 pull requests and 10 issues updated in the last 24 hours. A new beta release **v2.0.0-beta.2** was published as part of the ongoing QwenPaw 2.0 development cycle, signaling active preparation for a major version. The team merged 32 PRs — a strong mix of bug fixes, feature additions, and infrastructure improvements. The project continues to evolve rapidly, though the beta release carries stability warnings for production use.

## 2. Releases
**v2.0.0-beta.2** (early beta, 2026-07-02)  
[Release page](https://github.com/agentscope-ai/CoPaw/releases/tag/v2.0.0-beta.2)  
⚠️ This is an early beta for QwenPaw 2.0.0, currently under active development. It may contain **breaking changes and instability** and is intended for developers/early adopters only.  

**What's Changed**  
- `feat(cli): add cron up` – preliminary CLI cron support  

No migration notes or breaking changes explicitly documented for this release; developers should test carefully before upgrading.

## 3. Project Progress
**Merged/closed PRs today: 32**  
Notable merged PRs that advanced features or fixed bugs:

- **Goal Mode Gate Fix** ([#5727](https://github.com/agentscope-ai/CoPaw/pull/5727)) – Fixed goal mode gate architecture and scope filtering bug that prevented `/goal` mode from working correctly.  
- **ADBPG Memory REST-only** ([#5296](https://github.com/agentscope-ai/CoPaw/pull/5296)) – Removed SQL/psycopg2 path for ADBPG long-term memory, now purely REST-based with auto-search wiring.  
- **File-only Messages Fix** ([#5693](https://github.com/agentscope-ai/CoPaw/pull/5693)) – File-only messages (e.g., Excel via WeChat) were silently dropped by no-text debounce; now bypass for audio and file-only cases.  
- **Chat Session UI Improvements** ([#5728](https://github.com/agentscope-ai/CoPaw/pull/5728)) – Added date grouping, search, and VariableSizeList to the ChatSessionDrawer.  
- **Builtin Model Updates** ([#5730](https://github.com/agentscope-ai/CoPaw/pull/5730)) – Updated built-in model definitions.  
- **Governance Policy Pattern** ([#5546](https://github.com/agentscope-ai/CoPaw/pull/5546)) – Generalized governance policy pattern for agent configuration.  
- **Feishu Bot Filter Fix** ([#5724](https://github.com/agentscope-ai/CoPaw/pull/5724)) – Replaced blanket bot-message discard with self-bot-only filter, allowing inter-agent Feishu @mentions.  
- **Website Default Language** ([#5729](https://github.com/agentscope-ai/CoPaw/pull/5729)) – Switched website default to English and updated meta description.

## 4. Community Hot Topics
Most active issues (by comment count in the last 24h):

- **[#5403](https://github.com/agentscope-ai/CoPaw/issues/5403)** – Browser autofill hijacks search input on Model Configuration page (6 comments)  
  *Need:* User experience and input field type detection: browser misidentifies a search box as a credential field.  
- **[#5705](https://github.com/agentscope-ai/CoPaw/issues/5705)** – Secret sanitization and secure storage (env var fallback + log redaction missing) (5 comments)  
  *Need:* Stronger security for API keys, especially in agent.json and logs.  
- **[#5720](https://github.com/agentscope-ai/CoPaw/issues/5720)** – Memory leak in v1.1.12.post2 (4 comments)  
  *Need:* Root cause analysis: async task leakage + unreleased HTTP sessions caused 580MB memory growth in ~64 minutes.  
- **[#5725](https://github.com/agentscope-ai/CoPaw/issues/5725)** – Browser lag during streaming output (3 comments)  
  *Need:* UI performance: streaming output causes browser stutter, which does not occur with DeepSeek web.  
- **[#5708](https://github.com/agentscope-ai/CoPaw/issues/5708)** – Feishu interactive card messages not parsed (2 comments)  
  *Need:* Missing support for interactive card content in Feishu channel.  
- **[#5721](https://github.com/agentscope-ai/CoPaw/issues/5721)** – Feishu group sessions lose per-message sender identity (1 comment)  
  *Need:* When `share_session_in_group=True`, the real speaker is not stored, causing context confusion.

Community concerns center on **Feishu integration quality**, **performance (memory leaks / streaming lag)**, and **security hardening**.

## 5. Bugs & Stability
**High Severity**
- **Memory leak** ([#5720](https://github.com/agentscope-ai/CoPaw/issues/5720)) – `v1.1.12.post2`: ~5.5MB/min growth, process killed after 64 minutes, data corruption on restart. **No fix PR yet.**
- **Vector index explosion** ([#4795](https://github.com/agentscope-ai/CoPaw/issues/4795) – **closed** after fix: ChromaDB index swelled to 37GB over 3 months, causing `memory_search` crashes. The user resolved by removing the directory; likely patched in a later release.

**Medium Severity**
- **Browser autofill on Model Configuration** ([#5403](https://github.com/agentscope-ai/CoPaw/issues/5403) – **open**) – UX issue, not data corruption.
- **Streaming lag** ([#5725](https://github.com/agentscope-ai/CoPaw/issues/5725) – **open**) – Browser frame drops during LLM streaming.
- **Feishu card parsing** ([#5708](https://github.com/agentscope-ai/CoPaw/issues/5708) – **open**) – Card messages ignored; no fix PR yet.
- **Feishu sender identity lost** ([#5721](https://github.com/agentscope-ai/CoPaw/issues/5721) – **open**) – Shared sessions lose per-user attribution.

**Low Severity (fixed today)**
- Feishu bot message filtering ([#5709](https://github.com/agentscope-ai/CoPaw/issues/5709) – **closed**) – Fixed by PR [#5724](https://github.com/agentscope-ai/CoPaw/pull/5724).
- File-only messages being dropped ([#5693](https://github.com/agentscope-ai/CoPaw/issues/5693) – **closed**) – Fixed by PR [#5693](https://github.com/agentscope-ai/CoPaw/pull/5693).

## 6. Feature Requests & Roadmap Signals
**Major feature requests raised today:**
- **[#5737](https://github.com/agentscope-ai/CoPaw/issues/5737)** – **Enhanced CLI** for non-graphical scenarios (e.g., pre-installing skills at startup). Suggests growing need from B2B integrators.
- **[#5705](https://github.com/agentscope-ai/CoPaw/issues/5705)** – **Secret sanitization** (env var fallback expansion + log redaction). Security is a recurring theme.

**In-progress features visible in open PRs (likely for v2.0 or next minor):**
- **Reranker support for memory search** – Two PRs: backend ([#5692](https://github.com/agentscope-ai/CoPaw/pull/5692)) and UI ([#5691](https://github.com/agentscope-ai/CoPaw/pull/5691)) for reme0.4 with `qwen3-rerank` and generic reranker API.
- **Vision fallback** ([#5726](https://github.com/agentscope-ai/CoPaw/pull/5726)) – Automatically use a vision model when the active model is text-only and images are uploaded.
- **Windows native sandbox** ([#5525](https://github.com/agentscope-ai/CoPaw/pull/5525)) – First-time contributor PR for Windows sandbox.
- **`none` memory backend** ([#5732](https://github.com/agentscope-ai/CoPaw/pull/5732)) – Allows completely disabling memory system via console.
- **Multi-dimensional rate limiting** ([#5738](https://github.com/agentscope-ai/CoPaw/pull/5738)) – Account + IP based protection.
- **GitHub Models provider update** ([#5735](https://github.com/agentscope-ai/CoPaw/pull/5735)) – Migrate to new inference endpoint, support fine-grained PAT.

**Prediction:** v2.0.0-beta.2 likely paves the way for a stable 2.0 with reranker, vision fallback, and enhanced security. The CLI enhancement request may accelerate into an early 2.0.x release if B2B demand is high.

## 7. User Feedback Summary
**Pain points expressed today:**
- **Memory leak** (NICKMXAK): Process grows to 580MB in ~1 hour, killed externally, config corruption leaves a poor impression of stability.
- **Streaming performance** (593199118): QwenPaw browser UI lags while DeepSeek does not – points to inefficient frontend rendering during streaming.
- **Feishu inter-agent communication** (ZhaoX666, hellozhouuu): Bot messages and interactive cards are not working; agents cannot @-mention each other or read forms.
- **Vector database bloat** (liudao008, historical but still relevant): 37GB ChromaDB growth alarmed users, though closed as fixed.
- **Browser autofill** (xiaofengtt): Small annoyance interfering with configuration workflow.

**Satisfaction signals:**
- The community actively contributes features (e.g., multiple first-time-contributor PRs today: [#5525](https://github.com/agentscope-ai/CoPaw/pull/5525), [#5693](https://github.com/agentscope-ai/CoPaw/pull/5693), [#5731](https://github.com/agentscope-ai/CoPaw/pull/5731), [#5669](https://github.com/agentscope-ai/CoPaw/pull/5669)).
- The swift fix for Feishu bot filtering (reported July 1, fixed July 2) demonstrates responsive issue handling.

## 8. Backlog Watch
Issues and PRs that appear to require maintainer attention due to age or lack of response:

- **[#5187](https://github.com/agentscope-ai/CoPaw/pull/5187)** – **Windows desktop GUI automation (computer_use)** – Open since June 14, no comments or reviews. A significant feature (UIA + Tauri control mode) that has not been merged or discussed.
- **[#5403](https://github.com/agentscope-ai/CoPaw/issues/5403)** – Browser autofill bug – Open since June 23, 6 comments, no assignee or PR. The user workaround is trivial, but the fix is UI-level (adding `autocomplete="off"` or similar).
- **[#5725](https://github.com/agentscope-ai/CoPaw/issues/5725)** – Streaming lag – Opened today, but root cause (frontend rendering vs network) needs triage. One comment suggests debouncing or DOM optimization.
- **[#5708](https://github.com/agentscope-ai/CoPaw/issues/5708)** – Feishu interactive card parsing – Open since July 1, no assignee. The user mentions a previous report from June 13 that was not tracked in GitHub, indicating a potential gap in issue tracking.

**Recommendation:** Prioritize the memory leak ([#5720](https://github.com/agentscope-ai/CoPaw/issues/5720)) for a hotfix, and assign reviewers to the large open PRs ([#5187](https://github.com/agentscope-ai/CoPaw/pull/5187), [#5525](https://github.com/agentscope-ai/CoPaw/pull/5525)) to prevent stagnation.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-07-02

## 1. Today's Overview

The ZeroClaw project is experiencing a phase of **intense development**, with **50 pull requests updated in the last 24 hours** (10 merged/closed) and **12 active issues** — all open. No new releases were cut today. The activity signals a major push to stabilize cross‑platform support, integrate new channel providers, and address long‑standing provider compatibility bugs. However, **two high‑priority (p1) bugs** — 74 Windows test failures and a Gemini 400 error — remain unresolved and pose significant risk to reliability. The community is actively contributing features (OpenAI‑compatible endpoint, SOP improvements, per‑chat model switching) and security hardening (zip‑bomb protection, signing‑key leak fix). Release velocity appears to be temporarily paused while maintainers process a large backlog of PRs.

## 2. Releases

No releases were published in the last 24 hours. The latest release remains undetermined from this data set.

## 3. Project Progress

**Merged/closed PRs today (10 total)** — two of the largest are captured in the top‑20 list:

- **#8504** ([feat(channels): add Git forge channel with SOP ingress](zeroclaw-labs/zeroclaw PR #8504)) — A massive XL‑sized PR adding a provider‑backed `channel-git` implementation (GitHub App provider, polling, issue/PR comments, lifecycle events). This PR was closed, meaning it has been merged or superseded.
- **#7361** ([feat(rfc-6969): per-turn output routing via send_via + voice delivery fixes](zeroclaw-labs/zeroclaw PR #7361)) — An XL‑sized feature implementing RFC‑6969, including a double‑send bug fix for Telegram and voice‑only peer delivery consolidation.

The remaining 8 closed PRs (not shown in the top 20) likely include smaller fixes and dependency bumps. The high merge rate indicates **active maintainer throughput** despite the heavy workload.

## 4. Community Hot Topics

The most active discussions (by comment count) highlight critical user pain points:

- **#7462** ([Bug]: 74 test failures on Windows](zeroclaw-labs/zeroclaw Issue #7462) — 6 comments. The root cause is Unix‑only test commands and path semantics. Users are frustrated by the lack of Windows CI.
- **#6302** ([Bug]: Gemini 400 — assistant tool_call emitted as first non-system turn](zeroclaw-labs/zeroclaw Issue #6302) — 5 comments. A provider‑side invariant violation that blocks Gemini users entirely.
- **#7952** ([Feature]: publish full-channel prebuilt assets alongside default prebuilts](zeroclaw-labs/zeroclaw Issue #7952) — 4 comments. Users want optional bundles to avoid missing channels.
- **#7108** ([feat(ci): improve cached Rust builds and CI critical path](zeroclaw-labs/zeroclaw Issue #7108) — 3 comments. Developers want faster CI feedback loops.
- **#8550** ([Feature]: Add OpenAI-compatible chat completions endpoint](zeroclaw-labs/zeroclaw Issue #8550) — 2 comments. A gateway feature that would unlock many third‑party integrations.

**Underlying needs:** Cross‑platform parity (Windows), provider compliance (Gemini), build tooling speed, and API interoperability. The OpenAI endpoint request (#8550) is accompanied by a dedicated RFC PR (#8603), indicating strong community demand.

## 5. Bugs & Stability

**High‑priority (p1) bugs reported or updated today:**

- **#7462** — Windows test suite 74 failures (risk:high). No fix PR yet; CI remains Linux‑only.
- **#6302** — Gemini 400 error (risk:high). No fix PR yet; remains a blocker for Gemini provider users.
- **#8605** ([Bug]: runtime-config self-modification guard misses the real config.toml](zeroclaw-labs/zeroclaw Issue #8605) — **New today**, filed by Nillth. The security policy guard that prevents agents from modifying their own config is misaligned with the per‑agent directory layout, leaving the guard ineffective.

**Medium‑priority (p2) bugs:**

- **#8615** ([Bug]: compatible provider silently deletes content via unconditional `<think>` tag stripping](zeroclaw-labs/zeroclaw Issue #8615) — New today. Content loss without user awareness.
- **#8598** ([Bug]: skills install cannot install owner-qualified ClawHub skill URLs](zeroclaw-labs/zeroclaw Issue #8598) — Reported yesterday, no comments.

**Fixes in progress (open PRs addressing bugs):**

- **#8576** ([fix(channels): add env-var fallback for OpenAI STT credentials](zeroclaw-labs/zeroclaw PR #8576)) — Fix for missing API key handling.
- **#8574** ([fix(skills): harden extract_zip_secure against zip-bomb inflation](zeroclaw-labs/zeroclaw PR #8574)) — Addresses #8554.
- **#8591** ([fix(audit): prevent signing-key leak via VarError formatting](zeroclaw-labs/zeroclaw PR #8591)) — Security fix.
- **#8547** ([fix(audit): remove rag-pdf feature to clear RUSTSEC-2026-0192](zeroclaw-labs/zeroclaw PR #8547)) — Removes a vulnerable dependency.
- **#8599** ([fix(agent): align Agent::from_config tool dispatcher and prompt with active provider per turn](zeroclaw-labs/zeroclaw PR #8599)) — Addresses #8054.
- **#8616** ([fix(skills): restore always: true frontmatter flag for compact prompt mode](zeroclaw-labs/zeroclaw PR #8616)) — Fixes #7904.

**Security advisory:** PR #8547 eliminates `RUSTSEC-2026-0192` (ttf‑parser vulnerability) by dropping the `rag-pdf` feature. This is a positive step, but users relying on PDF tooling will lose that capability.

## 6. Feature Requests & Roadmap Signals

Several notable feature requests surfaced or gained activity today:

- **#8600** ([Feature]: easy per-chat model switching for multi-model providers](zeroclaw-labs/zeroclaw Issue #8600) — A user migrating from Moltis wants the ability to switch between any model in a provider’s catalog without restarting. High‑risk, likely requires runtime provider flexibility.
- **#8602** ([Feature]: Enhance file_read — default line cap, charset detection, paged PDF, notebook awareness, chunked binary](zeroclaw-labs/zeroclaw Issue #8602) — New today, with detailed spec inspired by Claude Code's `Read` tool.
- **#8550** and **#8603** ([RFC: OpenAI Chat Completions compatibility adapter](zeroclaw-labs/zeroclaw Issue #8603)) — Two closely related issues proposing an OpenAI‑compatible REST endpoint. This would be a **major gateway feature** and likely a candidate for the next minor release.
- **#8587** ([Docs]: adding more SOPs examples to syntax](zeroclaw-labs/zeroclaw Issue #8587) — Documentation gap; SOP engine is praised but lacks usage depth.
- **#7952** (publish full‑channel prebuilts) — Blocked, awaiting maintainer review.

**Prediction for next release:** The presence of RFC #8603 and a dedicated adapter PR suggest the **OpenAI Chat Completions endpoint** may land soon. The **file_read enhancements** (#8602) and **per‑chat model switching** (#8600) are also strong candidates, but they are new and need design validation. The **visual SOP authoring** PR (#8590) is XL‑sized and may be too large for a quick merge.

## 7. User Feedback Summary

**Pain points (explicit or implicit from issues/PRs):**

- **Windows is a second‑class citizen** — #7462: 74 failures, zero CI coverage. User "NiuBlibing" (multiple Windows‑related reports) is likely a dedicated Windows tester.
- **Gemini integration is broken** — #6302: no workaround; users of Gemini via LiteLLM are stuck.
- **Skill installation is restrictive** — #8598: owner‑qualified ClawHub URLs rejected. User "Audacity88" reports degraded behavior.
- **Config guard bypasses security** — #8605: agents can modify their own config, a serious self‑modification loophole.
- **Silent content loss** — #8615: `<think>` tag stripping discards parts of responses without warning.
- **Desire for model flexibility** — #8600: user "vvuk" explicitly compares to Moltis, indicating a gap in provider model selection UX.
- **Third‑party integration barrier** — #8550 / #8603: users of Open WebUI, LobeChat cannot connect.

**Signs of satisfaction:**

- SOP engine is described as a “great concept” (#8587).
- Many PRs are authored by community contributors (Nillth, Audacity88, wangmiao0668000666, etc.), indicating an engaged developer base.
- PRs like #7361 (per‑turn routing) and #8504 (Git forge channel) are large, complex features that were successfully merged/closed, showing active maintainer review capacity.

## 8. Backlog Watch

Issues that have been open for a significant time without maintainer response or progress:

- **#6302** (Gemini 400, p1) — Open since **May 3, 2026** (60 days). Last updated today but no assignee or fix PR. This is a critical blocker for Gemini users.
- **#7462** (Windows test failures, p1) — Open since **June 10, 2026** (22 days). Still no Windows CI; risk remains high.
- **#7108** (CI improvement, p2) — Open since **June 2, 2026** (30 days). Accepted but no PR. Developers rely on CI speed.
- **#7952** (full‑channel prebuilts, p2) — Open since **June 19, 2026** (13 days). Blocked, needs maintainer review.
- **#8550** (OpenAI endpoint, p2) — Open since **June 30, 2026** (2 days). Already has an accompanying RFC, but both lack maintainer assignment.

**PRs needing attention:** While not “unchanged”, several high‑risk PRs remain open and could become stale:

- **#8590** (visual SOP authoring, XL size) — No comments from maintainers yet.
- **#8567** (OTel content policy, L size) — No comments.
- **#8488** (fix channel prompt tool‑availability, S size) — Open since June 29.
- **#8576** (STT env‑var fallback, S size) — Simple fix, good candidate for quick merge.

Maintainers should prioritize the p1 bugs (#6302, #7462, #8605) and review the lightweight fixes (#8576, #8591, #8616) to improve project health and contributor morale.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*