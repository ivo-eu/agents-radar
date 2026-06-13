# OpenClaw Ecosystem Digest 2026-06-13

> Issues: 147 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-13 10:15 UTC

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

# OpenClaw Project Digest — 2026-06-13

## 1. Today’s Overview
OpenClaw shows **very high activity** with 147 issues and 500 PRs updated in the last 24 hours. Of those, 81 issues remain open/active, 66 were closed; 343 PRs are open, and 157 were merged or closed. Two new releases (v2026.6.7‑beta.1 and v2026.6.6) were published, both focusing on **security hardening** and **channel delivery improvements**. The community is engaged around **cost regressions** (e.g., DeepSeek prompt cache), **RISC‑V compatibility**, and **session‑state bugs**. The project is actively merging fixes for long‑standing issues, with several critical PRs posted today.

## 2. Releases
**v2026.6.7‑beta.1** (2026‑06‑07)  
- **Highlights**: Tighter channel delivery across Slack, Telegram, outbound media, silent replies, progress drafts, and paged action results.  
- Changes include: same‑channel Slack finals persist in transcripts; top‑level `image` message‑tool sends attached media; expandable Telegram blockquotes and spool.  

**v2026.6.6** (2026‑06‑06)  
- **Highlights**: Substantially tightened security boundaries across transcripts, sandbox binds, host environment inheritance, MCP stdio, Codex HTTP access, native search policy, elevated sender checks, deleted‑agent ACP bypasses, loopback tools, Discord moderation, and Teams group actions.  
- Both releases are **beta/stable** – users on older versions are encouraged to upgrade, especially for the security fixes.  
- No explicit **breaking changes** or **migration notes** were published; however, the security changes may affect custom plugins or sandbox configurations.

## 3. Project Progress (Merged/Closed PRs Today)
Among the ~157 merged/closed PRs today, several notable ones landed:
- **#92651** – Fix: require admin for HTTP session kills (hardens auth).  
- **#92652** – Test: stabilize plugin auth marker fixtures (improves test reliability).  
- **#89438** – Fix(slack): warn when channels map keyed by name instead of ID (fixes #81665).  
- **#90660** – Fix(discord): filter startup slots by runtime‑enabled predicate (fixes #77429).  
- **#89835** – Feat(usage): native templated `/usage full` footer renderer.  
- **#90686** – Fix(gateway): honor profile auth for SecretRef model entries (fixes #90685).  
- **#74914** – Allow native google‑vertex model provider configs.  

These fixes improve **Slack routing**, **Discord startup stability**, **usage reporting**, **auth credential handling**, and **test infrastructure**.

## 4. Community Hot Topics

| Issue/PR | Comments | Reactions | Summary |
|----------|----------|-----------|---------|
| [Issue #54253](https://github.com/openclaw/openclaw/issues/54253) | 14 | 4 👍 | *Bug*: OpenClaw returns “run Error: LLM Request Failed” on RISC‑V64 systems. Author is excited about OpenClaw but blocked on RISC‑V. |
| [Issue #39604](https://github.com/openclaw/openclaw/issues/39604) | 13 | 9 👍 | *Feature*: Add `tools.web.fetch.allowPrivateNetwork` to enable internal network access. High demand (9 👍). |
| [Issue #39476](https://github.com/openclaw/openclaw/issues/39476) | 10 | 0 | *Bug*: A2A `sessions_send` causing duplicate messages when target agent calls back. |
| [Issue #7707](https://github.com/openclaw/openclaw/issues/7707) | 9 | 0 | *Feature*: Memory trust tagging by source to prevent poisoning. |
| [Issue #41165](https://github.com/openclaw/openclaw/issues/41165) | 8 | 2 👍 | *Bug*: Telegram DMs still route into `agent:main:main` instead of isolated session. |
| [Issue #41201](https://github.com/openclaw/openclaw/issues/41201) | 8 | 1 👍 | *Bug*: Control UI avatar not displaying (broken image). |
| [Issue #91778](https://github.com/openclaw/openclaw/issues/91778) | 8 | 3 👍 | *Bug* (closed P0): `memory_search` index metadata missing since v2026.6.1. |
| [Issue #91018](https://github.com/openclaw/openclaw/issues/91018) | 7 | 3 👍 | *Bug*: DeepSeek prompt cache broken after upgrade – $6 burned in one hour. |

**Underlying needs**: Users are demanding **security conscious** features (private network access, memory trust, sender validation) and **stability** across platforms (RISC‑V, Telegram, Discord). The DeepSeek cost issue highlights a **critical regression** for users paying per‑token.

## 5. Bugs & Stability
**Critical / P0–P1 bugs reported or updated today:**

- **#91778 (CLOSED P0)** – `memory_search` index metadata missing since v2026.6.1 – was reproduced on v2026.6.5. Closed today.  
- **#91018 (OPEN P1)** – DeepSeek prompt cache broken after upgrading to v2026.6.1, causing **~$6/hour burn**. Fix PR not yet linked; issue remains open with 7 comments.  
- **#54253 (OPEN P2)** – LLM request failed on **RISC‑V64** systems. No fix PR visible.  
- **#41165 (OPEN P1)** – Telegram DMs still pollute main session after #40519.  
- **#66443 (OPEN P1)** – Overflow recovery duplicates `role=user` messages, amplifying transcript growth.  
- **#39847 (OPEN P1)** – Echo contamination: stripInboundMetadata not called in outbound delivery pipeline (security).  
- **#92479 (OPEN P1)** – Zen provider ships no model catalog; every model must be hand‑registered.  
- **#92094 (OPEN P1)** – `message` tool returns “unsupported channel: telegram”.  

**Regressions**: #41201 (avatar broken), #91778 (memory search), #91018 (DeepSeek cache), #91920 (cron jobs failing).  

**Existing fix PRs** for some of these:  
- #92639 (fix memory_search transient mode) – addresses #91778? not directly linked but related.  
- #92642 (fix subagent output file) – addresses #86872.  
- #92650 (split OpenAI 431 embedding batches) – addresses #92465.  
- #90861 (fix `sessions_yield` on MCP path) – addresses #77426.  

Many bugs remain **unfixed**, especially those needing product/security review.

## 6. Feature Requests & Roadmap Signals

| Feature | Issue | Votes | Likely Next‑Version Candidate |
|---------|-------|-------|-------------------------------|
| Private network access for `web_fetch` | [#39604](https://github.com/openclaw/openclaw/issues/39604) | 9 👍 | **Strong candidate** – high demand, clear config option. |
| Memory trust tagging by source | [#7707](https://github.com/openclaw/openclaw/issues/7707) | 0 | Open since Feb – may need product decision. |
| Path‑scoped RWX permissions | [#39979](https://github.com/openclaw/openclaw/issues/39979) | 0 | **Security‑focused** – aligns with v2026.6.6 trends. |
| Automated session memory preservation on `/new` | [#40418](https://github.com/openclaw/openclaw/issues/40418) | 1 👍 | Useful for continuous learning. |
| Cumulative context usage in `/usage` footer | [#40215](https://github.com/openclaw/openclaw/issues/40215) | 3 👍 | PR #89835 merged – **already in v2026.6.7?** |
| Per‑agent dreaming configuration | [#67413](https://github.com/openclaw/openclaw/issues/67413) | 4 👍 | OOM protection – high value. |
| Image batching / media group buffering | [#39343](https://github.com/openclaw/openclaw/issues/39343) | 1 👍 | UX quality improvement. |
| Agent teleconference primitive | [#65403](https://github.com/openclaw/openclaw/issues/65403) | 1 👍 | Experimental RFC. |
| WhatsApp create group tool | [#74261](https://github.com/openclaw/openclaw/issues/74261) | 1 👍 | Niche but clear use case. |
| Cron calendar timeline view | [#40644](https://github.com/openclaw/openclaw/issues/40644) | 1 👍 | PR #41892 open – may land soon. |

**Prediction**: The next minor release (v2026.6.8?) is likely to include **private network access** (#39604) and **per‑agent dreaming** (#67413) – both have strong community backing and align with the security/dreaming themes of recent releases.

## 7. User Feedback Summary
- **Pain points**:
  - **RISC‑V incompatibility** – one user excited but blocked.
  - **Cost surprises** – DeepSeek prompt cache loss burned real money.
  - **Session state confusion** – duplicate messages, wrong routing, lost context.
  - **Security loopholes** – metadata leaks, weak sandboxing, missing permission controls.
  - **Broken UI elements** – avatar, cron jobs, media attachments.
- **Use cases**:
  - Personal assistants on **Mac/Android**.
  - **Multi‑platform** agents (Discord, Telegram, Slack, Feishu, WhatsApp).
  - **Agent‑to‑agent** collaboration (A2A).
  - **Low‑power hardware** (RISC‑V) – growing interest.
- **Satisfaction**:
  - Users praise the project’s ambition and feature set (e.g., “great experiences on Mac & x86”).
  - However, bugs and regressions cause frustration, especially after upgrades.

**Overall sentiment**: Mixed. The project is innovating fast but **stability regressions** (especially around memory and provider caching) are eroding trust. The security‑focused releases are appreciated.

## 8. Backlog Watch (Long‑Unanswered Important Issues)

| Issue | Created | Status | Needs |
|-------|---------|--------|-------|
| [#7707](https://github.com/openclaw/openclaw/issues/7707) – Memory trust tagging | 2026‑02‑03 | Open, P2 | Maintainer review, product decision, security review |
| [#38721](https://github.com/openclaw/openclaw/issues/38721) – Gateway shutdown timeout | 2026‑03‑07 | Open, P1 | Live repro, maintainer review |
| [#39604](https://github.com/openclaw/openclaw/issues/39604) – Private network access | 2026‑03‑08 | Open, P2 | All reviews (maintainer, product, security) |
| [#39476](https://github.com/openclaw/openclaw/issues/39476) – A2A duplicate messages | 2026‑03‑08 | Open, P1 | Product decision, linked PR open |
| [#39979](https://github.com/openclaw/openclaw/issues/39979) – Path‑scoped RWX | 2026‑03‑08 | Open, P2 | All reviews |
| [#40418](https://github.com/openclaw/openclaw/issues/40418) – Session memory preservation | 2026‑03‑09 | Open, P2 | Maintainer review, product decision |
| [#41346](https://github.com/openclaw/openclaw/issues/41346) – Cron jobs auto‑enable without confirmation | 2026‑03‑09 | Open, P1 | Security review, product decision |
| [#39496](https://github.com/openclaw/openclaw/pull/39496) – Feishu comprehensive plugin PR | 2026‑03‑08 | Open | Needs real‑behavior proof – massive PR stalled |
| [#40311](https://github.com/openclaw/openclaw/pull/40311) – Brave Goggles support | 2026‑03‑08 | Open | Needs proof |
| [#41375](https://github.com/openclaw/openclaw/pull/41375) – Fix hooks delivery | 2026‑03‑09 | Open | Waiting on author |

These issues and PRs have been open for **3–4 months** and carry high‑impact tags (P1, security, session‑state). They represent **unresolved architectural concerns** that could block future stability. Community members and maintainers should prioritise them.

---

## Cross-Ecosystem Comparison

# Cross‑Project Comparison Report — 2026‑06‑13

## 1. Ecosystem Overview

The open‑source personal AI agent landscape is experiencing a surge in both feature velocity and stability pressure. Multiple projects—OpenClaw, NanoBot, Hermes Agent, IronClaw, ZeroClaw—show daily PR counts exceeding 50, reflecting a highly active contributor base. Core infrastructure (memory, tool‑call reliability, session state) remains the dominant technical concern, while security hardening and cost control (especially around LLM provider caching) have emerged as cross‑project priorities. The ecosystem is fragmenting into reference implementations (OpenClaw), lightweight alternatives (NanoBot, PicoClaw), and enterprise‑focused forks (IronClaw, ZeroClaw), each targeting distinct deployment profiles.

## 2. Activity Comparison

| Project | Issues (24h) | PRs (24h) | Releases (24h) | Health Score | Notes |
|---------|-------------|-----------|----------------|--------------|-------|
| **OpenClaw** | 147 updated (81 open, 66 closed) | 500 updated (343 open, 157 merged/closed) | 2 (v2026.6.7‑beta.1, v2026.6.6) | ⚠️ High activity but critical regressions (DeepSeek cache, memory_search) | Reference implementation; security‑focused releases |
| **NanoBot** | 6 (3 closed, 3 open) | 35 (10 merged, 25 open) | 0 | ✅ Stable growth; few new regressions | Active feature work (TUI, TTS, audit) |
| **Hermes Agent** | 12 (all open) | 50 (5 merged, 45 open) | 0 | ⚠️ Very high PR throughput, but critical bugs (self‑termination, E2EE gap) | Rapid iteration, high community contribution |
| **PicoClaw** | 2 new open | 10 (3 merged, 7 open) | 1 nightly | ✅ Solid; small scope, focused fixes | Turn‑lifecycle and WebSocket improvements |
| **NanoClaw** | 5 updated | 15 (10 merged, 5 open) | 0 | ✅ Positive; many reliability fixes merged | Signal support, backup/restore, API retry |
| **NullClaw** | 2 updated | 3 open (0 merged) | 0 | 🔴 Low activity; critical cron bug unfixed | Stalled PRs, no maintainer response |
| **IronClaw** | 7 (4 open, 3 closed) | 50 (16 merged, 34 open) | 0 (release PR pending) | ✅ Strong; DeferredBusy drain & attachment tracks merging | Enterprise‑scale reliability improvements |
| **LobsterAI** | 4 (all open, stale) | 8 (5 merged, 3 open) | 0 | ⚠️ Consolidation phase; open issues >70 days | Computer Use MVP merged, but skill lifecycle bugs |
| **Moltis** | 2 updated | 0 | 0 | 🔴 Very low activity; no PRs merged | Only bug report and feature request |
| **CoPaw** | 12 (11 open) | 4 (1 merged) | 0 | ⚠️ Moderate, but high‑severity bugs (infinite loop, timer failure) | Timer‑scheduling and long‑context freezes |
| **ZeroClaw** | 5 (3 closed) | 50 (26 merged, 24 open) | 0 | ✅ Heavy consolidation; turn engine refactor | High‑risk architectural changes |

*Health Score: ✅ Positive (active merging, few regressions), ⚠️ Mixed (high activity but unresolved critical issues), 🔴 Stagnant (low activity or stalled fixes)*

## 3. OpenClaw’s Position

**Advantages:** OpenClaw remains the ecosystem’s most comprehensive reference implementation, with the largest issue/PR volume (147 issues, 500 PRs) and two security‑focused releases in one week. It supports the widest range of channels (Slack, Telegram, Discord, WhatsApp) and providers (OpenAI, DeepSeek, Google Vertex). Security hardening is a clear differentiator—v2026.6.6 introduced tightened boundaries for transcripts, sandbox binds, and MCP stdio.

**Weaknesses:** The same scale creates stability regressions. Critical P0/P1 bugs (e.g., DeepSeek cache burn costing $6/h, RISC‑V incompatibility) erode user trust. Community sentiment is mixed: “great experiences on Mac & x86” but frustration with “stability regressions after upgrades.” The project’s ambition sometimes outpaces quality assurance.

**Technical approach differences:** OpenClaw uses a modular multi‑engine runtime (A2A, MCP, native tools) while lighter projects like NanoBot or PicoClaw favour streamlined single‑engine designs. OpenClaw’s security model is the most sophisticated (sandbox binds, host environment inheritance) but also the most complex to configure.

**Community size:** OpenClaw’s contributor base is by far the largest (~157 merged PRs/day). Comparative daily merges: ZeroClaw 26, IronClaw 16, NanoBot 10, NanoClaw 10, Hermes Agent 5, LobsterAI 5, PicoClaw 3, CoPaw 1, NullClaw 0, Moltis 0.

## 4. Shared Technical Focus Areas

The following requirements appear across **three or more** projects, indicating ecosystem‑wide pain points:

- **Session state consistency & deduplication** – **OpenClaw** (A2A duplicates, Telegram routing), **NanoBot** (orphan tool results, idle‑compact loss), **NanoClaw** (silent response drops), **Hermes Agent** (Missed thread‑initial messages), **CoPaw** (timer scheduling). All projects struggle with reliable state management across turns and channels.

- **Tool‑call reliability / timeout handling** – **NanoBot** (`session_key` crash, malformed history), **NanoClaw** (hung MCP tool for 30+ min), **Hermes Agent** (`delegate_task` blocks parent), **OpenClaw** (DeepSeek cache regression). The ecosystem lacks a standardised per‑tool timeout or circuit‑breaker.

- **Security gating & credential handling** – **OpenClaw** (private network access, memory trust tagging), **Hermes Agent** (Matrix E2EE gap), **NanoClaw** (ungated `create_agent` tool), **IronClaw** (approval gate order). There is growing demand for fine‑grained access controls, especially around LLM‑generated tool invocations.

- **Memory / context management** – **OpenClaw** (`memory_search` metadata missing), **NanoBot** (history corruption in dream reads), **ZeroClaw** (dream mode consolidation), **CoPaw** (long‑context freeze). Projects are adopting periodic compaction or summarisation, but no consistent approach exists.

- **Cost governance** – **OpenClaw** (DeepSeek cache break caused $6/h burn), **NanoBot** (zero token usage in API), **PicoClaw** (evolution mode token drain). As AI agent usage grows, users demand transparent cost monitoring and prevention mechanisms.

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | PicoClaw | IronClaw | ZeroClaw |
|-----------|----------|---------|----------|----------|----------|
| **Target user** | Power users, multi‑channel deployments | API‑first developers, lightweight agents | Embedded/edge devices (RISC‑V) | Enterprise, Slack‑centric | CLI/skill‑driven users |
| **Core architecture** | Multi‑engine (A2A, MCP, native) | Single engine, OpenAI‑compatible API | Minimal Go runtime, WebSocket protocol | Rust‑based, DeferredBusy draining | Rust, turn engine consolidation |
| **Channel strength** | 10+ (Slack, Telegram, Discord, WhatsApp, Feishu) | 5 (WebUI, CLI, Telegram, Discord) | 4 (Telegram, DeltaChat, WebSocket) | Slack, Teams, MCP | CLI, ZeroCode Chat |
| **Security posture** | High (sandbox binds, stdio isolation) | Low (no sandbox, limited permission model) | Medium (JSON marshal checks) | Medium (approval gates, credential checks) | Medium (MCP auto‑reconnect) |
| **Release cadence** | Weekly betas + stable | Irregular | Nightly | Monthly | Rolling (v0.8.1 pending) |
| **Differentiator** | Security & channel breadth | Simplicity & TUI/TTS features | WebSocket protocol & edge support | Reliability engineering (DeferredBusy) | Skills background review & dream mode |

## 6. Community Momentum & Maturity

**Tier 1 – Rapidly iterating (high throughput, new features merging):**  
**OpenClaw**, **ZeroClaw**, **IronClaw** – maintainers are merging dozens of PRs daily, including architectural refactors. All three have defined next‑release roadmaps (security hardening, turn engine consolidation, attachment support). Risk of regressions is high but matched by quick fix cycles.

**Tier 2 – Stable growth (consistent fixes, occasional features):**  
**NanoBot**, **NanoClaw**, **PicoClaw**, **LobsterAI** – these projects merge 3–10 PRs daily, focusing on bug fixes and incremental features. No critical outages reported. They benefit from smaller scope and less complex merge conflicts.

**Tier 3 – Stabilising / low activity:**  
**Hermes Agent** (high PR volume but critical self‑termination bug), **CoPaw** (timer and context freeze bugs), **NullClaw** (no merges, key cron bug unfixed), **Moltis** (zero PRs). These projects need maintainer intervention or community support to regain momentum.

**Tier 4 – Dormant:**  
**TinyClaw**, **ZeptoClaw** – no activity in 24 hours. Likely inactive or awaiting maintainer contributions.

## 7. Trend Signals

1. **Security‑first development is becoming the norm** – OpenClaw’s two security releases (v2026.6.6, v2026.6.7‑beta) and IronClaw’s approval gate fixes signal a shift from “move fast” to “move safely.” Expect more projects to adopt sandboxing, credential validation, and private network controls.

2. **Edge and embedded deployment is growing** – PicoClaw’s WebSocket protocol and RISC‑V support, alongside OpenClaw’s RISC‑V compatibility issue, indicate that users want agents on low‑power hardware (IoT, routers, edge servers). This will drive demand for lightweight runtimes and async protocol design.

3. **Cost observability is a major gap** – DeepSeek cache regressions burning $6/h, evolution‑mode token drain, and zero‑usage API responses show that no project offers reliable real‑time cost metering. This is a clear opportunity for tooling (e.g., `usage full` footer in OpenClaw) but needs standardisation.

4. **Agent‑to‑agent (A2A) collaboration is emerging** – Issues across OpenClaw (A2A duplicate messages), NanoClaw (agent‑to‑agent reply routing), and Hermes Agent (`delegate_task` blocking) show that inter‑agent trust and communication are still immature. Standard protocols (e.g., MCP‑based A2A) are likely to become a differentiator.

5. **User‑interface diversity is increasing** – TUI clients (NanoBot), ZeroCode Chat (ZeroClaw), Desktop voice mode (Hermes Agent), and WebUI parity (NanoBot) reflect a shift beyond traditional chat interfaces. Developers should invest in multi‑modal inputs (voice, images, file attachments) and platform‑aware rendering (macOS vs. Linux keyboard shortcuts).

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest – 2026-06-13

## Today's Overview

Over the past 24 hours, NanoBot maintained a high development cadence with 35 pull requests updated (10 merged/closed, 25 open) and 6 issues updated (3 closed, 3 open). No new releases were published. The majority of activity focused on fixing critical bugs in conversation history handling, async crash prevention, and improving WebUI configuration parity. Several feature PRs were merged—notably an audit tool for agent observability and enhanced WebUI settings—while new capabilities such as a TUI client, text-to-speech support, and reverse‑proxy compatibility entered review. Community participation remained strong, with contributions from 10+ distinct authors.

## Releases

None in this period.

## Project Progress

Ten pull requests were merged or closed today, advancing both stability and feature breadth:

- **WebUI / config parity** – PR #4313 (merged) closed the gap between WebUI settings panels and `config.json`, adding write endpoints for temperature, tool limits, dream, channels, and memory fields.
- **Audit tool (agent observability)** – Three PRs (#4318, #4319, #4320) were merged, introducing a configurable `tools.audit` module with four transport backends (loguru, HTTP webhook, JSONL file, callback). Only the final version (#4320) remains open for further review.
- **Bug fixes** – Merged PRs addressed: env‑var template resolution in transcription config (#4323), WebUI settings read/write paths (#4324, #4325), and media attachment validation (#4312). The `idleCompact` fix for full session tail summarization (#4326) and a fix for dream cursor advancement (#4321) are also open but derived from today’s merged work.
- **Documentation / schema** – PR #4314 refactored tool config schemas to break an import cycle, preserving the self‑describing tool pattern.

## Community Hot Topics

The most active discussion threads this period revolve around **conversation history integrity**:

- **Issue #4203** ([link](https://github.com/HKUDS/nanobot/issues/4203), 3 comments) – A closed bug report describing a logical error in `find_legal_message_start` that caused all messages to be discarded when a user message is followed by an orphaned tool result. The fix was merged in an earlier PR.
- **Issue #4006** ([link](https://github.com/HKUDS/nanobot/issues/4006), 2 comments) – Reports lingering orphan tool results in conversation history after the tool_call_id substitution fix. Underlying need: strict API compatibility (OpenAI/Anthropic) and avoiding trace renderer errors.
- **Issue #4264** ([link](https://github.com/HKUDS/nanobot/issues/4264), 1 comment) – Requests that `idleCompact` use the complete session history rather than dropping the last 8 messages. A fix PR (#4326) is already open.

These issues signal that **tool‑call/result pairing** and **session memory compaction** are the community’s top pain points, directly affecting reliability for production deployments.

## Bugs & Stability

Reported bugs ranked by severity:

| Severity | Issue | Description | Fix PR status |
|----------|-------|-------------|---------------|
| **Critical** | [#4322](https://github.com/HKUDS/nanobot/issues/4322) | `NameError: 'session_key' not defined` in `context.py` after merging `fix/prompt-caching` branch → crashes agent on startup. | No fix PR yet; newly reported. |
| **High** | [#4309](https://github.com/HKUDS/nanobot/issues/4309) | `/v1/chat/completions` always returns zero token usage despite real tracking inside the agent loop. Undermines billing/monitoring for API users. | No fix PR yet. |
| **Medium** | [#4264](https://github.com/HKUDS/nanobot/issues/4264) | `idleCompact` omits the last 8 messages, potentially losing the correct final answer and leaving erroneous history. | Fix PR #4326 open. |
| **Medium** | [#4203](https://github.com/HKUDS/nanobot/issues/4203) (closed) | History drop when orphan tool results follow a user message. | Already fixed; regression test added. |
| **Low** | [#4006](https://github.com/HKUDS/nanobot/issues/4006) (closed) | Orphan tool results still present after earlier fix. | Fix merged, but issue closed only today. |

Additionally, several open PRs target crash bugs: #4303 (MCP generator crash on reconnect), #4315 (malformed history entries in Dream reads), #4311 (non‑positive pagination limits).

## Feature Requests & Roadmap Signals

New features under active development or discussion:

- **Terminal UI (TUI)** – PR #4329 introduces an inline interactive TUI for `nanobot agent`, with markdown rendering, persistent input loop, slash commands, and multi‑modal input. Likely to be merged in next minor release.
- **Text‑to‑Speech (TTS)** – PR #4316 adds a multi‑provider TTS system (OpenAI, Groq/Orpheus, ElevenLabs) with WebUI settings. Indicates growing demand for voice output.
- **Multiple custom providers** – Issue #4305 (closed) requested a “template” parameter to re-use built‑in provider schemas. Although closed quickly, the idea may reappear as a more structured config enhancement.
- **File tool toggle** – PR #4138 adds `tools.file.enable` to disable built‑in filesystem tools, mirroring `exec` and `web` toggles. Useful for sandboxed deployments.
- **Audit tool** – Merged and already in review, this will become a core observability feature for agent actions.
- **Reverse‑proxy / sub‑path support** – PR #4328 fixes WebUI asset/API paths when served under a prefix; important for enterprise deployments.

Predictions for next version (likely 0.5.x): TUI as default interaction mode, TTS configuration, audit tool stable release, and the idleCompact fix.

## User Feedback Summary

- **Pain points**: Users continue to struggle with conversation history corruption (orphan tool results, lost corrections) and zero token reporting in the OpenAI‑compatible API. Env‑var template resolution not happening in WebUI settings caused silent failures (addressed by #4323–#4325).
- **Use cases driving development**: Multi‑provider setups (custom providers, TTS, transcription), reverse‑proxy hosting, observability for compliance (audit), and interactive CLI usage without full‑screen TUI.
- **Satisfaction**: High engagement on PRs and quick maintainer response (many issues closed within days). Contributors actively shape the project, as seen by the WebUI parity and audit tool additions.
- **Dissatisfaction**: The `session_key` crash (#4322) after merging a feature branch highlights test‑gap concerns for complex merges. No official releases since the previous version may cause some frustration for users wanting stable binaries.

## Backlog Watch

Items requiring maintainer attention:

- **PR #4138** (open since 2026-06-01) – `tools.file.enable` toggle. Blocked? No comments from maintainers; important for deployments that disable filesystem access.
- **PR #4303** (open since 2026-06-11) – MCP generator crash fix. Needs review; the crash affects streaming MCP server sessions.
- **Issue #4264** (open, 1 comment) – idleCompact history loss; fix PR #4326 exists but hasn’t been merged. This backlog item affects daily users.
- **PR #4316** (open) – TTS system. Large feature with three providers; should be reviewed soon to avoid bitrot.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-06-13

## 1. Today's Overview
The Hermes Agent project saw very high activity over the past 24 hours, with 12 issues updated (all still open) and 50 pull requests updated (45 open, 5 merged/closed). No new releases were published. The community is reporting a wide range of bugs — from critical startup failures to encryption gaps and UI issues — while also contributing a large volume of repair PRs and a few new feature implementations. The rapid PR throughput (50 updates in one day) indicates an actively maintained and responsive codebase, though the sheer number of open issues suggests ongoing stability challenges.

## 2. Releases
No new releases were published today. The latest available release remains the v0.16.0 mentioned in issues (e.g., #45508).

## 3. Project Progress
Five pull requests were merged or closed today, although their specific content is not detailed in the available data. Among the 45 open PRs updated today, the majority are bug fixes (e.g., TTS profile routing, Matrix encryption, gateway streaming, desktop environment variables) that are likely to be merged soon. Key fix areas include:
- Desktop voice-mode TTS config (#45522)
- E2EE encryption for Matrix text messages (#45518)
- Gateway stream final‑send suppression (#45517)
- CLI fallback providers config nesting (#45516)
- Environment variable parsing robustness (#45521, #45514, #45511)
- Telegram gateway socket leak (#45507)

Feature PRs also advanced, notably adding the **Z.AI GLM‑5.2 model** (1M context window) (#45483) and an **auth credential pool switch command** (#45513). These reflect active progress on both stability and new capabilities.

## 4. Community Hot Topics
The most active issue by comment count is **[#18646 – send_message ignores WhatsApp group target](https://github.com/NousResearch/hermes-agent/issues/18646)** (6 comments, created 2026‑05‑02, still open). Users report that messages intended for WhatsApp groups are incorrectly routed to the user’s home channel. The issue has been open for over a month and is tagged `P2`, suggesting it is a known pain point for WhatsApp users.

Another notable issue is **[#45508 – Agent self‑diagnosed as 'killing herself' and cannot automatically restart](https://github.com/NousResearch/hermes-agent/issues/45508)** (1 comment, created today). The user describes a scenario where the agent terminates itself and fails to recover after a v0.16.0 update. While it has few comments, its dramatic title and critical nature have likely drawn attention.

No PRs in the top‑20 list had comment counts recorded, so the most active discussions appear to be on issues rather than PRs.

## 5. Bugs & Stability
Many stability and correctness bugs were reported today. They are ranked by severity below:

| Severity | Issue | Description | Fix PR Exists? |
|----------|-------|-------------|----------------|
| 🔴 **Critical** | [#45508](https://github.com/NousResearch/hermes-agent/issues/45508) | Agent self‑diagnoses as “killing herself” and cannot restart after v0.16.0 update. | No specific PR yet |
| 🔴 **High** | [#45500](https://github.com/NousResearch/hermes-agent/issues/45500) | Matrix text messages bypass E2EE encryption (security/encryption gap). | [#45518](https://github.com/NousResearch/hermes-agent/pull/45518) |
| 🔴 **High** | [#45496](https://github.com/NousResearch/hermes-agent/issues/45496) | `delegate_task` blocks parent session silently when child agent hangs on tool calls (no progress output). | Not yet |
| 🔴 **High** | [#37414](https://github.com/NousResearch/hermes-agent/issues/37414) | Codex app‑server shell output not truncated, saturating context and causing mid‑thread context loss. | Not yet |
| 🟠 **Medium** | [#45505](https://github.com/NousResearch/hermes-agent/issues/45505) | OpenWebUI Responses API sessions cannot approve guarded tool/code execution. | Not yet |
| 🟠 **Medium** | [#45506](https://github.com/NousResearch/hermes-agent/issues/45506) | Desktop voice mode uses default profile TTS config for non‑default profiles. | [#45522](https://github.com/NousResearch/hermes-agent/pull/45522) |
| 🟠 **Medium** | [#45499](https://github.com/NousResearch/hermes-agent/issues/45499) | API/gateway turn setup seeds NULL `system_prompt` before prompt restore, causing misleading warnings. | Not yet |
| 🟠 **Medium** | [#45520](https://github.com/NousResearch/hermes-agent/issues/45520) | Dashboard `/chat` fails on Linux VPS with software rendering (WebGL unavailable). | Not yet |
| 🟠 **Medium** | [#45519](https://github.com/NousResearch/hermes-agent/issues/45519) | GLM‑5.2 context window misdetected as 202,752 (instead of 1M), leading to premature compression. | [#45483](https://github.com/NousResearch/hermes-agent/pull/45483) addresses model metadata |
| 🟡 **Low** | [#45493](https://github.com/NousResearch/hermes-agent/issues/45493) | Matrix: agent’s own thread‑initial message lost from next‑turn context. | Not yet |
| 🟡 **Low** | [#18646](https://github.com/NousResearch/hermes-agent/issues/18646) | WhatsApp group routing ignored (older bug, still open). | Not yet |

Multiple PRs were opened today to address several of these issues (noted in the rightmost column). The Matrix encryption gap (#45500) and Desktop TTS (#45506) each already have dedicated fixes proposed.

## 6. Feature Requests & Roadmap Signals
The following feature‑oriented PRs and issues indicate likely directions for the next release:

- **Model Support**: [#45483](https://github.com/NousResearch/hermes-agent/pull/45483) adds **Z.AI GLM‑5.2** (1M context window) to the provider registry. This suggests the team is actively adding support for high‑context models.
- **Authentication**: [#45513](https://github.com/NousResearch/hermes-agent/pull/45513) introduces a **credential pool switch command** (`hermes auth switch`), allowing users to manually select which credential in a same‑provider pool is tried first — a quality‑of‑life feature for multi‑credential setups.
- **Integration**: Two long‑standing PRs from April 17 continue to be updated: [#11460](https://github.com/NousResearch/hermes-agent/pull/11460) (bidirectional Nextcloud file sync) and [#11459](https://github.com/NousResearch/hermes-agent/pull/11459) (Nextcloud notifications/activity background service). Both are tagged `P3` but have seen activity today, possibly moving closer to merge.
- **Documentation**: [#15150](https://github.com/NousResearch/hermes-agent/issues/15150) (Thai translation for tools‑reference, created April 24) is still open, showing community interest in i18n.

If these PRs merge, v0.17.0 could include GLM‑5.2 support, Nextcloud file sync, and credential management improvements.

## 7. User Feedback Summary
Users are reporting a mixed experience:

- **Satisfaction**: The community is highly engaged — 50 PRs updated in one day indicates strong contributor trust and willingness to fix bugs. The availability of new features like Z.AI GLM‑5.2 and credential switching shows active development.
- **Pain points**: Critical bugs dominate today’s feedback:
  - The agent self‑terminating (#45508) is alarming and blocks basic usage.
  - Matrix E2EE being bypassed for text messages (#45500) is a security concern.
  - `delegate_task` hanging (#45496) can silently block automated workflows.
  - Desktop voice profile misconfiguration (#45506) frustrates multi‑profile users.
  - OpenWebUI integration (#45505) is broken for approval workflows.
  - Linux VPS users are unable to use the dashboard (#45520), limiting accessibility.

The overall tone is pragmatic: users are filing detailed bug reports and, in many cases, contributing fixes themselves (e.g., #45522, #45518, #45516). The high ratio of fix PRs to new issues suggests the community is actively stabilizing the agent.

## 8. Backlog Watch
Several important issues and PRs have been open for weeks or months without visible maintainer activity:

- **#18646** (WhatsApp group routing, created May 2, last updated today) – still no fix merged.
- **#15150** (Thai translation, created April 24) – stalled documentation task.
- **#37414** (Codex context saturation, created June 2) – high‑impact bug with no fix PR yet.
- **#9833** (Patch tool sensitive‑path check bypass, created April 14) – security‑related PR still open, tagged `P1`, risks being forgotten.
- **#11460** and **#11459** (Nextcloud features, created April 17) – large feature PRs that have not been merged, though they were updated today. They may need maintainer review to finalize.

These items should be prioritized for maintainer attention to prevent technical debt and community frustration.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

## PicoClaw Project Digest — 2026-06-13

### 1. Today’s Overview
The project saw **high activity** over the past 24 hours, with **2 new open issues**, **10 pull requests updated** (7 open, 3 merged/closed), and **1 new nightly release**. Key development themes include fixing agent evolution token consumption, completing turn‑lifecycle signaling for WebSocket clients, and adding new providers (NEAR AI Cloud) and channels (DeltaChat). The community is actively contributing bug fixes and protocol enhancements, while maintainers are closing long‑standing refactoring PRs. Overall project health remains solid, though several stale items remain in the backlog.

### 2. Releases
- **`v0.2.9-nightly.20260613.c362114c`** (Nightly Build) — automated build, potentially unstable.  
  **Full Changelog**: [compare/v0.2.9...main](https://github.com/sipeed/picoclaw/compare/v0.2.9...main)  
  *No breaking changes or migration notes provided; use with caution in non‑production environments.*

### 3. Project Progress
Three pull requests were **merged/closed** in the last 24 hours:
- **#2551 (closed)** — `refactor: standardize channel identification and decouple name from provider type`  
  *Decouples channel names from provider IDs, enabling multiple instances of the same provider. Adds `ChannelType` to `InboundContext`.*  
  [PR #2551](https://github.com/sipeed/picoclaw/pull/2551)
- **#3113 (closed)** — `fix(channels): check json marshal/unmarshal errors in toChannelHashes`  
  *Catches three silently discarded serialization errors in channel manager.*  
  [PR #3113](https://github.com/sipeed/picoclaw/pull/3113)
- **#3112 (closed)** — `fix(tools): handle json.Marshal error in toolloop tool call arguments`  
  *Prevents tool call data loss when `Arguments` marshaling fails.*  
  [PR #3112](https://github.com/sipeed/picoclaw/pull/3112)

Open PRs advancing significant features include:  
| PR | Description |
|----|-------------|
| [#3117](https://github.com/sipeed/picoclaw/pull/3117) | Route media turns to image models (fixes #3108) |
| [#3118](https://github.com/sipeed/picoclaw/pull/3118) | Add remote Pico WebSocket mode to `picoclaw agent` |
| [#3116](https://github.com/sipeed/picoclaw/pull/3116) | Complete `turn.done` lifecycle signaling for #2984 |
| [#3063](https://github.com/sipeed/picoclaw/pull/3063) | Add DeltaChat gateway |
| [#2917](https://github.com/sipeed/picoclaw/pull/2917) | Add NEAR AI Cloud provider |
| [#2964](https://github.com/sipeed/picoclaw/pull/2964) | Configurable inbound image compression |
| [#3115](https://github.com/sipeed/picoclaw/pull/3115) | Fix inline data URL media extraction for generic tool output |

### 4. Community Hot Topics
The two most active issues:
- **#2984** — `[Feature][Protocol] Add explicit turn completion signal for Pico WebSocket clients`  
  *2 comments, 2 👍. Users want a deterministic way to know when an agent has finished processing.*  
  [Issue #2984](https://github.com/sipeed/picoclaw/issues/2984)  
  *Underlying need:* External protocol clients require a clear `turn.done` event to avoid race conditions and unreliable polling.
- **#3012** — `[BUG] Continuous consumption of tokens every minutes when evolution is enabled`  
  *3 comments, 0 👍. Token usage spikes even when idle.*  
  [Issue #3012](https://github.com/sipeed/picoclaw/issues/3012)  
  *Underlying need:* Cost‑sensitive users expect evolution mode to not burn tokens without user interaction.

### 5. Bugs & Stability
| Severity | Bug | Fix PR Exists? |
|----------|-----|----------------|
| **High** | **#3012**: Continuous token consumption every minute when evolution is enabled (MiniMax, FreeBSD). *No fix PR yet.* | No |
| **Medium** | **#3108** (referenced by #3117): Media turns not routed to image models (text‑only model retry). *Fixed in PR #3117.* | Yes |
| **Medium** | **#3115**: Session history corruption from inline `data:image/...` strings in tool output (e.g., `read_file`, `exec`). *Fixed in PR #3115.* | Yes |
| **Low** | **#3113 / #3112**: Discarded JSON marshal errors (channels, tools). *Already merged.* | Yes (closed) |

### 6. Feature Requests & Roadmap Signals
- **#2984** — Turn completion signal for WebSocket clients (community‑requested). Already implemented in PR #3116; likely lands in next release.
- **#3118** — Remote agent mode via WebSocket. Enables headless or distributed usage.
- **#3063** — DeltaChat gateway integration.
- **#2917** — NEAR AI Cloud provider (OpenAI‑compatible, TEE‑capable).
- **#2964** — Image input compression with configurable policy.
- **#2551** (merged) — Multi‑instance channel support.

These indicate user demand for **richer protocol integration**, **cost‑efficient media handling**, and **expanded provider/channel ecosystems**.

### 7. User Feedback Summary
- **Pain points**: Excessive token usage in evolution mode (#3012) and lack of explicit turn‑completion signal (#2984) frustrate both casual and power users.
- **Use cases**: External protocol clients (bots, dashboards) require reliable signaling; users with image‑heavy workflows desire compression; multi‑instance channels for admins managing multiple bot accounts.
- **Satisfaction**: Contributors are actively submitting high‑quality fixes (JSON error handling, media routing), indicating a healthy contributor base.

### 8. Backlog Watch
Items needing maintainer attention (stale or long‑unanswered):
- **#3012** — `[stale]` Bug (token consumption). Last updated 2026-06-13 but marked stale. No fix PR yet; risk of user churn if not addressed soon.
- **#2964** — `[stale]` PR: image compression. Last updated 2026-06-12, but no reviewer activity since 2026-05-28.
- **#2917** — `[stale]` PR: NEAR AI Cloud provider. Last updated 2026-06-12, still awaiting maintainer review since May 21.
- **#2551** (already closed) — Was marked stale before merge; similar refactoring PRs may appear.

*Maintainers should prioritize reviewing #3012 and the two stalled PRs to keep momentum.*

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-13

## Today’s Overview
Activity remains steady, with **5 issues** updated and **15 pull requests** updated in the last 24 hours. The project merged/closed **10 PRs** (mostly bug fixes and feature work from maintainer `ddaniels`) while **5 new PRs** remain open for review. No releases were published today. The overall health is good, though several high-severity bugs (e.g., silent response drops, hung tool calls) are still unresolved and attracting community discussion.

## Releases
*None.* No new versions were published in the reporting period.

## Project Progress
**Merged/closed PRs today** (10 items) reflect continued work on Signal support, disaster recovery, and core agent-runner reliability:

- [#2203](https://github.com/nanocoai/nanoclaw/pull/2203) — **Signal reactions**: inbound + outbound via `add_reaction` MCP tool and chat-sdk-bridge integration.
- [#2072](https://github.com/nanocoai/nanoclaw/pull/2072) — **Ollama multimodal**: `ollama_generate` now accepts workspace-relative image paths (base64-encoded).
- [#2071](https://github.com/nanocoai/nanoclaw/pull/2071) — **Signal attachments**: all non-audio attachments are routed through the inbox path.
- [#2070](https://github.com/nanocoai/nanoclaw/pull/2070) — **Inbox attachment support**: `extractAttachmentFiles()` now handles host-file paths (Signal, future native adapters).
- [#2084](https://github.com/nanocoai/nanoclaw/pull/2084) — **Backup & restore**: daily project snapshots with local/S3 backends and per-agent CLI restore.
- [#2670](https://github.com/nanocoai/nanoclaw/pull/2670) — **Poisoned-resume crash loop fix**: session no longer hangs on corrupt transcripts (fixes #2669).
- [#2277](https://github.com/nanocoai/nanoclaw/pull/2277) — **Routing refresh on follow-up**: poll loop now re-evaluates routing context when follow-up messages arrive mid-query.
- [#2267](https://github.com/nanocoai/nanoclaw/pull/2267) — **Agent-to-agent reply routing**: a2a replies now go back to the originating session (fixes split-brain with multiple sessions per agent group).
- [#2692](https://github.com/nanocoai/nanoclaw/pull/2692) — **Poll-loop retry on 5xx API errors**: transient errors (e.g., `529 Overloaded`) are retried; user is notified on exhaustion.
- [#2040](https://github.com/nanocoai/nanoclaw/pull/2040) — **Signal outbound attachments**: `deliver()` now sends files via signal-cli’s `attachments` parameter.

## Community Hot Topics
The most active conversation centers on **silent message drops** and **session timeout issues**:

- [#2506](https://github.com/nanocoai/nanoclaw/issues/2506) (3 comments) — *Bug: send_message dedup silently drops responses when turns complete within 60 seconds of each other*. The underlying need is a robust dedup mechanism that doesn’t discard legitimate parallel responses. Users report client timeouts and lost replies.
- [#2632](https://github.com/nanocoai/nanoclaw/issues/2632) (1 comment) — *Clarify status of Telegram agent-swarm / multi-bot identity in v2*. A community member planning a v1→v2 migration seeks clarity on the future of the `add-telegram-swarm` feature.
- [#2711](https://github.com/nanocoai/nanoclaw/issues/2711) (1 comment) — *create_agent MCP tool is ungated despite “admin-only” comment*. A security-conscious user flags that any container can create agent groups, contradicting documentation.
- [#2668](https://github.com/nanocoai/nanoclaw/issues/2668) (1 comment) — *No per-tool timeout: hung MCP tool blocks session up to 30 min*. A single long tool call can stall an entire agent session, a critical reliability concern for production workloads.

## Bugs & Stability
Several open bugs affect core reliability. The most critical are:

| Severity | Issue / PR | Description | Fix available? |
|----------|------------|-------------|----------------|
| **High** | [#2506](https://github.com/nanocoai/nanoclaw/issues/2506) | Silent response drops when turns close together or follow-up arrives during streaming. | No fix PR yet. |
| **High** | [#2668](https://github.com/nanocoai/nanoclaw/issues/2668) | No per-tool timeout—hung MCP tool blocks session for 30+ minutes. | No fix PR yet. |
| **High** | [#2751](https://github.com/nanocoai/nanoclaw/issues/2751) (closed) | Budget-exhausted LLM turns silently dropped (synthetic 200 via gateway). | Closed, but fix not yet visible in a merged PR. |
| **Medium** | [#2711](https://github.com/nanocoai/nanoclaw/issues/2711) | Security: `create_agent` MCP tool is ungated for all containers. | No fix PR yet. |
| **Low-Medium** | [#2750](https://github.com/nanocoai/nanoclaw/pull/2750) (open) | Stale `outbound.db` journals after container kills; hot-journal poll races. | Fix PR open, attaches to #2516/#2640. |
| **Low** | [#2752](https://github.com/nanocoai/nanoclaw/pull/2752) (open) | Discord inbound attachments not reaching agent (URL-based only). | Fix PR open. |

## Feature Requests & Roadmap Signals
The community is asking for:

- **Proper admin gating for MCP tools** ([#2711](https://github.com/nanocoai/nanoclaw/issues/2711)) — likely to be addressed in an upcoming security-focused release.
- **Telegram swarm/multi-bot identity** ([#2632](https://github.com/nanocoai/nanoclaw/issues/2632)) — unclear if/when v2 will support this; may appear as a skill or native adapter extension.
- **Per-tool timeouts** ([#2668](https://github.com/nanocoai/nanoclaw/issues/2668)) — a strong candidate for the next minor release given the severity and clear design space.

Already-merged features (Signal reactions, backup/restore, Ollama multimodal, agent-to-agent routing) are now available on `main` and will likely ship in the next tagged release.

## User Feedback Summary
- **Pain points**: Silent response drops (urgent), long session blocks from hung tools, and unsecured MCP tools. Users migrating from v1 express anxiety about feature parity (Telegram swarm).
- **Satisfaction**: The recent batch of merged PRs addresses long-standing reliability issues (poisoned-resume loops, stale journals, API retry exhaustion). The new backup/restore feature is a direct response to operational risk.
- **Sentiment**: Constructive and security-aware. Reports include detailed reproduction steps and even proposed fix shapes (e.g., #2750, #2752).

## Backlog Watch
The following issues have been open for more than a week without a merged fix and may need maintainer attention:

- [#2506](https://github.com/nanocoai/nanoclaw/issues/2506) — *send_message dedup bug* (open since May 16, 3 comments). No PR has been linked.
- [#2632](https://github.com/nanocoai/nanoclaw/issues/2632) — *Telegram swarm status* (open since May 28). A clear roadmap answer would help community planning.
- [#2668](https://github.com/nanocoai/nanoclaw/issues/2668) — *No per-tool timeout* (open since June 1, 1 comment). Given the severity, a fix or workaround would be valuable.
- [#2711](https://github.com/nanocoai/nanoclaw/issues/2711) — *create_agent gating* (open since June 7). Security-impacting; attention is warranted.

No old PRs appear neglected; all open PRs were created within the last 48 hours.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-06-13

## 1. Today's Overview
The project shows moderate activity with 2 issues and 3 pull requests updated in the last 24 hours, though no PRs were merged or closed. A critical bug involving agent-type cron jobs failing to spawn subprocesses (Issue #941) remains open with no immediate fix. Three open PRs address configuration flexibility, error output suppression, and Discord gateway stability, indicating ongoing refinement. Overall, the project is in an active development phase but lacks recent releases or merged improvements.

## 2. Releases
*None.* No new versions were published in the reporting period.

## 3. Project Progress
No pull requests were merged or closed today. The following open PRs represent work in progress:
- **#949** – `fix: make queue_mode configurable from config.json` (vernonstinebaker)  
  Adds a `agent.default_queue_mode` config field and refactors the `QueueMode` enum.
- **#951** – `fix(agent_runner): suppress stderr initialization logs on agent failure` (vernonstinebaker)  
  Prevents internal initialization logs from being posted to channels when an agent process fails.
- **#953** – `fix(discord): recover closed gateway sockets` (vernonstinebaker)  
  Improves Discord gateway reconnection logic to handle stalled pre-HELLO scenarios and adds test coverage.

## 4. Community Hot Topics
- **Issue #941** – *Agent-type cron jobs don't spawn a subprocess — Telegram delivery never happens*  
  [GitHub](https://github.com/nullclaw/nullclaw/issues/941)  
  2 comments; created May 31, updated Jun 13.  
  *Analysis:* This is the most active issue by comment count. The underlying need is reliable scheduled agent execution. Users expect scheduled tasks with `job_type: "agent"` to actually run the agent and deliver results via Telegram, but the job completes without any subprocess spawning.

- **Issue #914** – *[enhancement] Create JIRA access tool*  
  [GitHub](https://github.com/nullclaw/nullclaw/issues/914)  
  1 comment; created May 13, updated Jun 13.  
  *Analysis:* A feature request for JIRA integration. The user wants agents and workflows to read, create, update, and comment on JIRA issues. This suggests demand for project management automation within the platform.

Pull requests have no comments and thus are less active in community discussion.

## 5. Bugs & Stability
| Severity | Bug Description | Status | Related Fix PR |
|----------|----------------|--------|----------------|
| **Critical** | Agent-type cron jobs do not spawn a subprocess; Telegram delivery never occurs (Issue #941). | Open, no fix PR yet | None |
| **Medium** | When an agent process fails, internal initialization logs (memory plan, MCP server registration) are incorrectly posted to channels as agent output (PR #951). | Addressed by PR #951 (open) | #951 |
| **Medium** | Discord gateway sockets remain closed after disconnect, preventing reconnection; stalled pre-HELLO reconnects are not detected as unhealthy (PR #953). | Addressed by PR #953 (open) | #953 |

The most impactful unresolved bug is #941, which breaks a core feature (scheduled agent execution). No fix has been proposed yet.

## 6. Feature Requests & Roadmap Signals
- **JIRA integration (Issue #914)** – Request for a tool to authenticate and perform common JIRA operations (reading issues, creating/updating tickets, adding comments, retrieving sprints).  
  *Prediction:* If accepted, this could appear in a future release as a new built‑in tool, likely requiring OAuth handling and a tool registry update.

- **Configurable queue mode for agents (PR #949)** – Although submitted as a fix, this introduces a new configuration option (`agent.default_queue_mode`). It signals that users want more control over how agent sessions are queued (e.g., "latest" mode). This is likely to be merged soon.

No other roadmap signals are present in the current data.

## 7. User Feedback Summary
- **Pain point:** User `weissfl` reports that scheduled agent jobs (Issue #941) are completely non‑functional for Telegram delivery. This represents a significant dissatisfaction with the reliability of the `schedule` feature.
- **Use case:** User `sayjeyhi` explicitly requests JIRA integration, indicating a need for enterprise project management connectivity.
- **Satisfaction:** No positive feedback or satisfaction indicators were observed in the current data.

## 8. Backlog Watch
- **Issue #914** – *Create JIRA access tool*  
  [GitHub](https://github.com/nullclaw/nullclaw/issues/914)  
  Opened May 13, 2026; last updated Jun 13, 2026. No maintainer response.  
  *Status:* Unanswered enhancement request >1 month old. Needs prioritization or at least a comment acknowledging it.

- **Issue #941** – *Agent-type cron jobs don't spawn a subprocess*  
  [GitHub](https://github.com/nullclaw/nullclaw/issues/941)  
  Opened May 31, 2026; updated Jun 13, 2026 with one user comment. No maintainer response.  
  *Status:* Critical bug with no maintainer acknowledgment or fix PR. Should be triaged immediately.

- **Pull Requests #949, #951, #953** – All open since Jun 10–12, 2026. No maintainer comments or reviews. While not extremely old, they represent pending contributions that risk stagnation.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-06-13

## Today's Overview

Activity remains high with 7 issues updated in the last 24 hours (4 open, 3 closed) and 50 pull requests touched (34 open, 16 merged/closed). The team is pushing forward on two major tracks: **attachment support** (Track 6 of #4644) and **DeferredBusy drain follow-ups** while also fixing several user-facing regressions like the Slack re-approval loop. No new releases were published today, but a release PR (#3708) continues to be updated. Project health is strong – most PRs are sized XL but carry low risk, indicating careful incremental work.

## Releases

*No new releases today.*  
The open release PR (#3708) includes breaking changes in `ironclaw_common` (v0.4.2 → 0.5.0) and `ironclaw_skills` (v0.3.0 → 0.4.0), but has not been merged yet.

## Project Progress

**16 pull requests were merged or closed today.** Based on the provided data, the following key changes landed:

- **DeferredBusy Drain Optimisation**  
  - [#4831](https://github.com/nearai/ironclaw/issues/4831) (closed) – Route drain resubmission through a product workflow replay entry point.  
  - [#4832](https://github.com/nearai/ironclaw/issues/4832) (closed) – Batch drained messages into a single run (depends on #4831).  
  - [#4833](https://github.com/nearai/ironclaw/issues/4833) (closed) – Filesystem backend: per-thread DeferredBusy index to avoid full transcript scans.  
  - [#4812](https://github.com/nearai/ironclaw/pull/4812) (closed) – Core drain logic: messages blocked on a gate now drain automatically when the blocking run terminates.

- **Attachment Foundation (#4644)**  
  - [#4654](https://github.com/nearai/ironclaw/pull/4654) (closed) – Extensible attachment format registry in `ironclaw_common`.  
  - [#4655](https://github.com/nearai/ironclaw/pull/4655) (closed) – Carry attachment refs through the Reborn transcript contract.

These advances bring the byte-storage and transcript integration tracks closer to completion, with the remaining PRs (e.g., #4668, #4670, #4672, #4675, #4676, #4677, #4680, #4738) still open and actively updated.

## Community Hot Topics

- **Issue #4817** ([link](https://github.com/nearai/ironclaw/issues/4817)) – *DeferredBusy drain follow-ups*: 3 comments, open. Discusses three structural decisions (trusted-resubmit seam, stale-intent policy, startup sweep) that were deferred from the main drain PR. The community (and authors) are tracking these as non-blocking but important design refinements.

- **Issue #4825** ([link](https://github.com/nearai/ironclaw/issues/4825)) – *Reborn: persist "always allow" approvals across threads*: 3 comments, open. A clear user pain point – approving a capability with “always allow” in one thread still re-prompts in new threads. This is a high-impact UX issue that likely needs a design decision on approval scope vs thread isolation.

- **PR #4841** ([link](https://github.com/nearai/ironclaw/pull/4841)) – *No run-borking failures*: XL, low risk, core contributor. This PR addresses a fundamental reliability need – every run-terminal error should be either recovered or explained. It aligns with the growing emphasis on non-borking agent loops.

- **PR #4839** ([link](https://github.com/nearai/ironclaw/pull/4839)) – *Fix Slack re-approval loop*: Direct fix for a repeat QA failure. High user impact – four consecutive approval prompts for one logical call.

Underlying these discussions is a clear need for **reliability and seamless UX** – users expect persistent approvals, non-repeating gates, and clear error recovery without lost work.

## Bugs & Stability

| Severity | Bug | Status | Fix PR |
|----------|-----|--------|--------|
| **High** | Nightly E2E scheduled run failing repeatedly ([#4108](https://github.com/nearai/ironclaw/issues/4108)) | Open since May 27, updated today with new failure (commit 2a4e017fd2) | None yet |
| **Medium** | Slack re-approval loop – invocation identity lost across auth‑gate re-dispatch ([#4839](https://github.com/nearai/ironclaw/pull/4839)) | Fix PR open | #4839 |
| **Medium** | Capability gate order: approval before credential check leads to burned approvals ([#4840](https://github.com/nearai/ironclaw/pull/4840)) | Fix PR open | #4840 |
| **Low** | Missing runtime context causing Slack Account Connection failures ([#4828](https://github.com/nearai/ironclaw/issues/4828) → PR #4836) | Fix PR open | #4836 |

The Nightly E2E failure remains the highest-priority open stability issue, with no fix PR yet. The other bugs have active fixes in review.

## Feature Requests & Roadmap Signals

- **Attachment end-to-end support** (#4644) – The largest feature in flight. With the registry and transcript refs merged, the remaining PRs (ingress upload, web UX, extraction, context folding) are nearing completion. Likely to be part of the next release.

- **Runtime context slice for models** – PR #4836 (open) implements #4828, giving the model visibility into connected channels, delivery state, and run origin. This directly addresses tester failures where Slack connections were not sensed.

- **Gated final-answer nudge** – PR #4837 adds a tool-free model call when the agent loop would produce an empty or canned ending. A UX improvement to avoid “I stopped because…” messages without real answers.

- **Explicit gate-open feedback** – PR #4838 replaces the DeferredBusy drain with a simple rejection notice when a thread is busy. Lightweight, but shifts retry responsibility to the user.

- **Non‑borking failure recovery** – PR #4841 aims to eliminate opaque run-terminal errors with explanation and retry capability. This is a significant reliability signal.

These features point toward a **more robust, user‑transparent agent** – less silent swallowing, more intelligent error handling, and richer model context.

## User Feedback Summary

- **Pain point: Approval gate re-prompting across threads** (Issue #4825) – Users who choose “always allow” expect durability. Re-prompting in new threads breaks trust and increases friction.
- **Pain point: Missing runtime context** (Issue #4828) – Slack Account Connection completion was not reflected to the model, causing repeat failures. Users could not proceed after connecting.
- **Pain point: Burned approvals due to credential ordering** (PR #4840) – Approving an action only to discover a missing credential wastes user effort.
- **Use case: Attachment uploads** – Multiple PRs (#4672, #4738) are adding web UX and model–visible document text. This is a heavily requested capability for daily workflows.

Overall user satisfaction appears mixed – core functionality is improving rapidly, but integration bugs (especially around Slack and persistent state) cause repeated friction.

## Backlog Watch

- **Nightly E2E Failure** ([#4108](https://github.com/nearai/ironclaw/issues/4108)) – Open since May 27, updated daily with failure logs. No assignment or fix PR. This is a blocker for release confidence.

- **Release PR** ([#3708](https://github.com/nearai/ironclaw/pull/3708)) – Open since May 16 with breaking changes. The PR is still being updated but the release itself has not been cut. Delays may affect downstream consumers.

- **Attachments Track 6 (byte landing)** – #4668, #4670, #4672, #4675, #4676, #4677, #4680, #4738 are all open since June 10–11, though actively updated. They represent the bulk of the remaining attachment work and will need maintainer attention to merge before the next release.

- **Gateway routine create endpoint** ([#4264](https://github.com/nearai/ironclaw/pull/4264)) – Open since May 31 from a new contributor (`wcc945`). Requires review to unblock external contributions.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-06-13

Generated from GitHub data (netease-youdao/LobsterAI). Activity snapshot based on items updated in the last 24 hours.

---

## 1. Today’s Overview

The project shows moderate activity with 8 PRs updated and 4 issues touched, although no new releases were published today. Five PRs were merged or closed, including the integration of the `release/2026.6.11` branch into `main` (PR #2158), which brings several major features to the codebase. All four open issues are marked as stale (last activity April 3) but were updated today, indicating maintainers may be reviewing them. The open PR count is three, focusing on skills management, artifacts preview, and UI layout improvements. Overall, the project appears to be in a consolidation phase following a feature release, with ongoing attention to bug fixes and community feedback.

---

## 2. Releases

**No new releases today.**  
The latest release was merged into `main` via PR #2158 (release/2026.6.11) on June 12. That release includes Computer Use MVP, real-time ASR voice input for cowork prompts, HTML artifact public sharing mode, and image/SVG artifact sharing support.

---

## 3. Project Progress

Five PRs were merged or closed in the last 24 hours:

- **#2158** – `chore(release): merge release/2026.6.11 into main`  
  Integrated the official June 11 release branch. Key highlights: Computer Use MVP, ASR voice input, artifact sharing modes.  
  [View](https://github.com/netease-youdao/LobsterAI/pull/2158)

- **#2156** – `fix(computer-use): bump runtime to 1.0.7`  
  Updated the managed Computer Use runtime with new CDN package and SHA-256 metadata. Includes UIA breadcrumbs for diagnosing helper exits.  
  [View](https://github.com/netease-youdao/LobsterAI/pull/2156)

- **#2157** – `fix(media): 修正文生图保存图片的扩展名`  
  Fixed image extension detection when saving generated images – now uses real file format instead of server-provided suffix, preventing incorrect `.jpg`/`.webp` extensions.  
  [View](https://github.com/netease-youdao/LobsterAI/pull/2157)

- **#1466** – `fix(mcp): modal close button unreachable when content grows tall`  
  Resolved scroll behaviour on MCP server form modal that caused Cancel button to go off-screen.  
  [View](https://github.com/netease-youdao/LobsterAI/pull/1466)

- **#1467** – `fix(shortcuts): display Cmd (⌘) instead of Ctrl on macOS`  
  Corrected keyboard shortcut labels to show platform‑appropriate modifier keys in Settings > Shortcuts.  
  [View](https://github.com/netease-youdao/LobsterAI/pull/1467)

The three currently open PRs (#1440, #1441, #1445) are still under review – see *Backlog Watch*.

---

## 4. Community Hot Topics

All four open issues were created on April 3 and received their latest update today. While each has only 1–2 comments, the topics reflect real user pain points:

- **#1443** – “有计划支持新版本的openclaw吗？” (Plan to support new openclaw version?)  
  User upgraded to openclaw v2026.3.24 and hit breaking changes. Asks for official support.  
  [Issue](https://github.com/netease-youdao/LobsterAI/issues/1443)

- **#1437** – “创建定时任务时，计划选择不重复，清空日历，点击【创建任务】按钮没反应”  
  Creating a one‑time scheduled task with empty calendar causes a silent failure – no error message.  
  [Issue](https://github.com/netease-youdao/LobsterAI/issues/1437)

- **#1439** – “上传技能已停用，对话中仍然可以调用”  
  Stopped skills still get invoked in conversation, indicating the disable action is not fully respected.  
  [Issue](https://github.com/netease-youdao/LobsterAI/issues/1439)

- **#1442** – “Agent添加技能，对话后引用的技能不展示，重新切换agent会重新展示”  
  After adding skills to an Agent, they disappear from display after the first conversation turn; reappear only after switching sessions. User questions the intended behaviour.  
  [Issue](https://github.com/netease-youdao/LobsterAI/issues/1442)

**Underlying needs:** Users want a reliable skill lifecycle – consistent display, proper disable semantics, and compatibility with external dependencies (openclaw). The silent failure on task creation suggests missing validation feedback.

---

## 5. Bugs & Stability

| Severity | Issue | Description | Fix PR? |
|----------|-------|-------------|---------|
| **High** | #1439 | Stopped skills still triggered in conversation – bypasses intended disabling | No |
| **High** | #1442 | Agent skills vanish after first conversation turn, reappear only after session switch | No |
| **Medium** | #1437 | “Create Task” button silently does nothing when calendar is cleared | No |
| **Medium** | #1443 | New openclaw version causes crash/startup failure | No |

Additionally, PR #1445 (still open) addresses two skill‑related bugs:
- Zip imports result in random directory names instead of skill names
- Duplicate skills are silently installed via any import channel (zip, folder, GitHub)  

[PR #1445](https://github.com/netease-youdao/LobsterAI/pull/1445)

No crash or regression was reported as *new* today; existing issues are being revisited.

---

## 6. Feature Requests & Roadmap Signals

- **#1443** – Support for openclaw v2026.3.24 (breaking changes).  
  Likely to be addressed if the maintainers plan to stay current with openclaw upstream.

- **PR #1441** – Extensible preview pipeline for HTML, React, and Mermaid artifacts (resurrected from older PR #1011). This was opened on April 3 and may be merged if community demand for richer artifact previews aligns with the June release’s artifact sharing features.

- **Released features from PR #2158** (now in `main`):  
  - **Computer Use MVP** (agent can control desktop GUI)  
  - **Realtime ASR voice input** for cowork prompts  
  - **HTML artifact public sharing mode** and image/SVG sharing support  

These indicate the roadmap is focusing on agent autonomy (Computer Use), multimodal input, and artifact dissemination.

**Prediction for next version:**  
- Openclaw compatibility update (if #1443 gains traction)  
- Skill lifecycle improvements (fixes for #1439, #1442, and #1445)  
- Further Computer Use refinements (runtime bumps, stability)

---

## 7. User Feedback Summary

**Real pain points from issues:**

- *Skill disabling is broken* – users expect stopped skills to be blocked from invocation, but they are still being routed (#1439).
- *Skill display is confusing* – after conversation, selected skills “disappear” from UI, leading to doubt about whether the agent is actually using them (#1442).
- *Task creation lacks validation* – clearing a date field yields a completely silent failure, leaving users frustrated without any error feedback (#1437).
- *Upgrade friction* – new openclaw versions break LobsterAI, and users want a clear migration path or version pinning (#1443).

**Satisfaction signals:**  
- The merged release PR includes Computer Use and voice input – these are likely welcomed features.
- PR #1440 (open) improves skill badge placement in the input bar, addressing a UI clutter complaint from users selecting many skills.

**Overall tone:** Users are experiencing functional + UX regressions, but remain engaged (they file issues and follow up). The project would benefit from clearer feedback on whether these known problems are being prioritized.

---

## 8. Backlog Watch

The following items have been open for over **70 days** (since April 3) and are marked as stale. Maintainers have not yet assigned or commented on them despite activity today:

- **#1443** – Openclaw version support (2 comments, user eagerly awaiting a response)  
  [Issue](https://github.com/netease-youdao/LobsterAI/issues/1443)

- **#1437** – Silent failure on scheduled task creation (1 comment)  
  [Issue](https://github.com/netease-youdao/LobsterAI/issues/1437)

- **#1439** – Disabled skills still invocable (1 comment)  
  [Issue](https://github.com/netease-youdao/LobsterAI/issues/1439)

- **#1442** – Agent skill display disappears after conversation (1 comment)  
  [Issue](https://github.com/netease-youdao/LobsterAI/issues/1442)

- **#1440** – Open PR: skill badge layout improvement (no comments from maintainers)  
  [PR](https://github.com/netease-youdao/LobsterAI/pull/1440)

- **#1441** – Open PR: artifacts preview pipeline (no comments)  
  [PR](https://github.com/netease-youdao/LobsterAI/pull/1441)

- **#1445** – Open PR: skill import validation (no comments)  
  [PR](https://github.com/netease-youdao/LobsterAI/pull/1445)

**Recommendation:** The above issues and PRs represent the most vocal user concerns and unsolved improvements. Assigning maintainers and providing status updates would improve project health and community trust.

---

*This digest was automatically generated from GitHub data. All links are to netease-youdao/LobsterAI.*

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest – 2026-06-13

## 1. Today's Overview
Project activity in the last 24 hours has been very low with **2 open issues updated** and **no pull requests or releases**. Both issues are fresh (created June 11–12) and have attracted initial community discussion. The lack of merged PRs or closed items suggests the team is either in a quiet maintenance phase or preparing work that hasn’t surfaced yet. The open issues cover a user-reported bug (Fastmail MCP authorisation) and a substantial feature request (Kubernetes-native sandbox), indicating that the community is actively shaping Moltis’s roadmap.

## 2. Releases
No new releases in the last 24 hours. No prior releases are listed in the provided data.

## 3. Project Progress
**Merged/closed PRs today:** 0  
No pull requests were updated or merged. No visible feature advancements or fixes were completed in this period.

## 4. Community Hot Topics
Two issues generated the only comments in the last 24 hours:

- **Issue #1115 – [Bug]: Fastmail MCP Authorisation**  
  _Author: kmath313 | Updated: 12 Jun | Comments: 2_  
  [Link](https://github.com/moltis-org/moltis/issues/1115)  
  This bug report describes an authorisation failure with the Fastmail MCP integration. The user has filed a preflight checklist but has not yet included full session context. The two comments suggest the maintainers may be requesting additional logs or steps to reproduce. Underlying need: reliability of MCP (Model Context Protocol) integrations with external email providers.

- **Issue #1118 – [Feature]: Add Kubernetes-native sandbox backend with runtimeClassName support**  
  _Author: AzgadAGZ | Updated: 12 Jun | Comments: 1_  
  [Link](https://github.com/moltis-org/moltis/issues/1118)  
  A detailed feature request proposing ephemeral Kubernetes pods for agent command execution, with support for VM‑level isolation (Kata Containers, gVisor). The single comment likely acknowledges the proposal. Underlying need: stronger isolation for untrusted LLM-generated commands in production deployments.

Both topics reflect the community’s dual focus: fixing integration reliability and expanding sandboxing options for enterprise/cloud-native environments.

## 5. Bugs & Stability
**One bug reported today:**

- **#1115 – Fastmail MCP Authorisation** – **Medium severity**  
  The issue prevents users from authenticating with Fastmail via Moltis’s MCP layer. No fix PR exists yet. The user has not provided full session context, making reproduction difficult. Severity is medium because it blocks a specific integration, not core functionality. No crashes or regressions were reported.

No other bugs or stability concerns were recorded.

## 6. Feature Requests & Roadmap Signals
The notable feature request is **#1118 – Kubernetes-native sandbox backend**. This aligns with the trend of running AI agents in secure, ephemeral environments. Given the detailed proposal and community interest (1 comment), it is a strong candidate for inclusion in the **next minor or major version**. Other signals are absent; the project appears to be prioritising infrastructure improvements over new user-facing features.

## 7. User Feedback Summary
From the available data, user sentiment is mixed:
- **Pain point:** A user encountered an authentication bug with Fastmail, indicating friction in integrating third‑party email services. The lack of immediate resolution may cause dissatisfaction.
- **Use case:** A user is requesting Kubernetes-native sandboxing, implying a growing adoption of Moltis in Kubernetes clusters where security and isolation are paramount. This suggests users are deploying Moltis in production or CI/CD pipelines.
- **Satisfaction:** No positive feedback is visible; the low activity volume may indicate either stability or stagnation.

## 8. Backlog Watch
No long‑standing or unanswered issues or PRs were identified in the last 24 hours. The two reported items are recent and have received timely attention (both have comments from the maintainer or community). No backlog items require escalation at this time.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest – 2026-06-13

## Today's Overview
The project shows moderate activity with 12 issues and 4 pull requests updated in the last 24 hours. One bug (attachment download in v1.1.11.post2) was resolved and closed, and one feature PR (skill tag batch download) was merged. However, 11 open issues remain, including several high-severity bugs and a growing list of feature requests. The community is actively engaging around a persistent timer-scheduling failure (11 comments) and a critical long-context freeze. No new releases were published today.

## Releases
*None* – No releases were created in the last 24 hours.

## Project Progress
**Merged / Closed PRs (1):**  
- **#4969** – `feat(skill): Add skill tag batch download` (by Leirunlin)  
  Added tag-filtering support for skill batch download to workspace. Closes issue #2961.  
  [Link](agentscope-ai/QwenPaw PR #4969)

**Closed Issue (1):**  
- **#5140** – Bug: Attachment download for non-text files (docx/pdf) returned 404 in v1.1.11.post2. **Fixed and closed.**  
  [Link](agentscope-ai/QwenPaw Issue #5140)

Other PRs remain open:  
- #5088 – Initial governance & sandbox interface discussion (under review)  
- #5069 – Visual model fallback for text-only primary models (under review)  
- #4622 – DataPaw plugin with 12 BI skills (first-time contributor, under review since May 22)

## Community Hot Topics
**Most Active Issue**  
- **#5064** – `[Bug]: 由Agent生成的定时任务, 无法正常触发` (11 comments)  
  Timer tasks created by the agent are not triggered at the scheduled time and cannot be manually edited. This indicates a core scheduling or persistence issue.  
  [Link](agentscope-ai/QwenPaw Issue #5064)

**Other Active Discussions**  
- **#5047** – `[Question]: Windows Tauri启动特别慢` (3 comments) – Users report startup times increasing from 1–2 minutes to over 10 minutes after the switch to Tauri.  
  [Link](agentscope-ai/QwenPaw Issue #5047)  
- **#5161** – `[Question]: 长对话后QwenPaw无响应` (2 comments) – Conversation stalls after many turns.  
  [Link](agentscope-ai/QwenPaw Issue #5161)  
- **#5156** – `[Feature]: 建议支持 kimi-for-coding / 加入 uv 白名单` (3 comments) – Users want to use Kimi Coding subscription directly.  
  [Link](agentscope-ai/QwenPaw Issue #5156)

## Bugs & Stability
Bugs reported in the last 24 hours, ranked by severity:

| Severity | Issue | Description | Status |
|----------|-------|-------------|--------|
| **Critical** | #5162 | Conversation thinking logic enters infinite loop (dead loop) – renders the agent unusable. | Open, no fix PR |
| **Critical** | #5161 | Long conversation causes QwenPaw to stop responding entirely. | Open, no fix PR |
| **High** | #5165 | Packaged install (Tauri + PyInstaller) results in white screen on launch – two missing modules (`qwenpaw.app.api`, `qwenpaw.app.middleware`). | Open, no fix PR |
| **High** | #5064 | Agent-generated timer tasks never trigger; no manual editing possible. | Open, 11 comments, no fix PR |
| **Medium** | #5163 | Regression: Gemini tool calling broken between v1.1.10 and v1.1.11.post2. | Open, no fix PR |
| **Medium** | #5166 | Plugin (TeamChat) fails to install on Python 3.13 due to missing `imghdr` module. | Open, no fix PR |

One high-severity bug (attachment download #5140) was closed today – fix verified.

## Feature Requests & Roadmap Signals
Several enhancement requests were opened today, suggesting community priorities for the next release:

- **#5156** – Add `kimi-for-coding` to API allowlist and support `uv` as a package manager.  
- **#5164** – Improve desktop system tray, auto-start, background service, and service management capabilities.  
- **#5167** – Optimize Feishu CardKit streaming card performance for long replies (current "word-by-word" display is too slow).  
- **#5168** – Add official Zalo Bot channel support (popular messaging platform in Vietnam).  

The **DataPaw** plugin (#4622) and **visual model fallback** (#5069) are already under review and likely to land in a future release.

## User Feedback Summary
**Pain Points (most frequently mentioned):**  
- Timer/agent-triggered tasks don't work (#5064) – *core reliability issue.*  
- Extreme startup slowness on Windows Tauri (#5047) – *migration regression.*  
- Long conversations freeze the UI (#5161) – *memory/context management issue.*  
- Gemini tool calling regression (#5163) – *API compatibility break.*  
- Installation white screen after packaging (#5165) – *build/deployment failure.*  

**Satisfaction Signals:**  
- Attachment download for plain text files is working correctly after the v1.1.11 fix (#5140).  
- The Feishu CardKit streaming feature is appreciated, though optimizations are requested (#5167).  
- The community is actively contributing new plugins (DataPaw, governance sandbox) and proposing integrations (Zalo, Kimi), indicating healthy engagement.

## Backlog Watch
The following issues and PRs have been open for several days without maintainer response or progress:

- **#5047** – *Windows Tauri slow startup* (opened June 9, 3 comments, no maintainer reply).  
- **#5064** – *Timer task not triggering* (opened June 10, 11 comments, no fix or official acknowledgment).  
- **#4622** – *DataPaw plugin PR* (opened May 22, no updates in 22 days, still under review).  
- **#5088** – *Governance & sandbox interface discussion* (opened June 10, no recent activity).  

These items would benefit from maintainer attention to avoid community frustration and to unblock contributions.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-06-13

*Generated from GitHub activity (zeroclaw‑labs/zeroclaw)*

---

## Today’s Overview

The project saw **heavy PR activity** with **50 pull requests updated** in the last 24 hours, of which **26 were merged or closed**. Issue management was lighter: **5 issues updated**, 2 remaining open, 3 closed. No new releases today. The focus was on consolidating runtime architecture (three turn engines → one unified loop), fixing installation bugs, and advancing the v0.8.1 integration queue tracked in issue #6970. Overall, the project remains highly active with contributors addressing both high‑risk refactors and user‑reported regressions.

---

## Releases

*None today.*

---

## Project Progress

**26 PRs were merged/closed** today. Notable items:

- **[#7554] [CLOSED]** – fix(install): install web dashboard on source installs and point users at the installer  
  *Author: singlerider* – Fixes issue #7553 where source builds never copied the web dashboard into the data directory.

- **[#7555] [CLOSED]** – feat(config): declarative section grouping for the Config menu  
  *Author: Nillth* – Introduced a declarative `SectionGroup` enum and `SECTION_GROUPS` table, replacing ad‑hoc grouping.

- **[#7553] [CLOSED]** – *Issue:* source install never copies web/dist → daemon serves no dashboard (fixed by #7554).

- Several other closed PRs (not individually listed) contributed to channel fixes, CI centralization, and tooling improvements.

- Among the **24 open PRs**, the most impactful ongoing work includes:
  - [#7540] – refactor(runtime): consolidate three agent turn engines onto `run_tool_call_loop` (implements #7415)
  - [#6667] – feat(skills): background review fork + `skill_manage` tool (skill improvement wiring)
  - [#6693] – feat(memory): dream mode for periodic memory consolidation
  - [#7546] – fix(runtime): unify SopEngine construction (single `Arc<Mutex<SopEngine>>` per daemon)

---

## Community Hot Topics

The most discussed **issues** (by comment count):

- **[#5470] [CLOSED] – [Bug]: Multiple issues when running safely**  
  *Comments: 5* – Reports of telegram memory duplication, bad user‑mention parsing, and other runtime anomalies. Scored `risk: high`, but now closed.  
  [zeroclaw-labs/zeroclaw Issue #5470](https://github.com/zeroclaw-labs/zeroclaw/issues/5470)

- **[#5570] [CLOSED] – [Enhancement]: Faster SQLite memory vector search (ANN)**  
  *Comments: 5* – Proposal to replace O(n) brute‑force scan with in‑process ANN index. Closed without merge, but signals user desire for performance improvements.  
  [zeroclaw-labs/zeroclaw Issue #5570](https://github.com/zeroclaw-labs/zeroclaw/issues/5570)

- **[#7557] [OPEN] – [Bug]: In ZeroCode Chat Shift+Enter not works**  
  *New issue (0 comments)* – A minor UX regression that may attract further discussion.  
  [zeroclaw-labs/zeroclaw Issue #7557](https://github.com/zeroclaw-labs/zeroclaw/issues/7557)

**Pull requests** with high activity (estimated from size/labels rather than comment counts, as comment counts were not provided):

- [#7540] – *refactor(runtime): consolidate turn engines* – Large refactor (XL size, high risk) touching agent, channel, gateway, runtime. Likely to generate significant review.
- [#7546] – *fix(runtime): unify SopEngine* – Critical to state consistency between agent tools and MQTT listener.

---

## Bugs & Stability

| Issue | Summary | Severity | Status | Fix PR |
|-------|---------|----------|--------|--------|
| [#7557] | Shift+Enter in ZeroCode Chat acts as Enter (can't add new line) | **S3** – minor | Open | None yet |
| [#7553] | Source install never copies web dashboard → daemon serves no dashboard | **S2** – degraded | Closed | [#7554] (merged) |
| [#5470] | Multiple runtime bugs (telegram memory, mention parsing, etc.) | **S2** – degraded | Closed | Not specified |

Additionally, several open **bug PRs** address stability concerns:

- [#6719] – *fix(runtime,channels): persist model_switch across turns in all paths* (p1, stale candidate)
- [#7399] – *fix(skills): sanitize skill tool names to satisfy provider name regex* (high risk)
- [#7351] – *feat(mcp): auto-reconnect on stale session or dropped stream* (high risk, MCP resilience)
- [#7207] – *fix(gateway): point paircode recovery hint at loopback for non-loopback binds* (high risk, gateway security)

No new regressions were reported today beyond the Shift+Enter minor issue.

---

## Feature Requests & Roadmap Signals

- **[#5570]** – *Faster SQLite memory vector search (ANN)* – Closed but reflects user demand for memory performance. May be revisited.
- **[#6970]** – *v0.8.1 integration/channel/provider/tool PR queue* – Operational tracker showing the maintainers are actively curating the next release’s feature set.
- **[#6667]** – *feat(skills): background review fork + skill_manage tool* – A major enhancement for skill self‑improvement, likely targeted for v0.8.1.
- **[#6693]** – *feat(memory): dream mode for periodic memory consolidation* – Another high‑profile enhancement under review.
- **[#7540]** – *refactor(runtime): consolidate turn engines* – Architectural improvement that simplifies future development.

**Prediction:** The next minor release (v0.8.1) will likely include the skills background review (#6667), dream mode (#6693), and the unified turn engine (#7540), alongside the installation fix already merged (#7554).

---

## User Feedback Summary

**Pain points reported today:**

- **Installation friction:** Source builds silently miss the web dashboard ([#7553]), leading to a broken UX. The fix was quick, but the root cause (missing copy step) suggests install scripts need better testing.
- **Telegram memory duplication:** User `nxtkofi` reported that every message is saved multiple times in memory, and Telegram user‑mention parsing fails ([#5470]).
- **Chat input inconsistency:** `technologicalsingularity` cannot use Shift+Enter to insert newlines in ZeroCode Chat ([#7557]) – a minor but irritating UX bug.
- **Skill timeouts:** A long‑running skill could never complete due to hardcoded 60s timeout – fixed by [#7552] (open PR).

**Satisfaction signals:** Contributors are actively engaging with bug reports (e.g., #7553 closed within hours). The high number of merged PRs (26) indicates a healthy contribution and review pipeline.

---

## Backlog Watch

The following **high‑risk / important items** have remained open for over a month without maintainer closure or author response:

| Item | Since | Label | Status |
|------|-------|-------|--------|
| [#6719] – persist model_switch across turns | 2026-05-16 | `priority:p1`, `needs-author-action`, `stale-candidate` | Open (PR) |
| [#6684] – distinct error for skill_manage patch disabled | 2026-05-15 | `needs-author-action` | Open (PR) |
| [#6667] – skills background review fork (XL) | 2026-05-14 | `needs-author-action` | Open (PR) |
| [#6693] – dream mode memory consolidation (XL) | 2026-05-16 | `needs-author-action` | Open (PR) |
| [#7351] – MCP auto‑reconnect (high risk) | 2026-06-07 | `needs-author-action` | Open (PR) |

These items risk becoming stale. The maintainers should consider merging or requesting updates to avoid losing momentum on critical features and fixes.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*