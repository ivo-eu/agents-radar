# AI CLI Tools Community Digest 2026-06-14

> Generated: 2026-06-14 11:10 UTC | Tools covered: 9

- [Claude Code](https://github.com/anthropics/claude-code)
- [OpenAI Codex](https://github.com/openai/codex)
- [Gemini CLI](https://github.com/google-gemini/gemini-cli)
- [GitHub Copilot CLI](https://github.com/github/copilot-cli)
- [Kimi Code CLI](https://github.com/MoonshotAI/kimi-cli)
- [OpenCode](https://github.com/anomalyco/opencode)
- [Pi](https://github.com/badlogic/pi-mono)
- [Qwen Code](https://github.com/QwenLM/qwen-code)
- [DeepSeek TUI](https://github.com/Hmbown/DeepSeek-TUI)
- [Claude Code Skills](https://github.com/anthropics/skills)

---

## Cross-Tool Comparison

# AI CLI Developer Tools: Cross-Tool Comparison Report
**Date:** 2026-06-14 | **Scope:** 9 major tools across the ecosystem

---

## 1. Ecosystem Overview

The AI CLI tooling landscape is experiencing a maturation phase marked by increasing community expectations around **agent reliability, cross-platform parity, and production-grade safety**. While all tools share a core value proposition—AI-assisted coding through terminal interfaces—the ecosystem is diverging along axes of **orchestration complexity** (single-agent vs. multi-agent fleets), **extensibility** (MCP/plugin ecosystems vs. monolithic design), and **monetization strategy** (OAuth free tiers, subscription billing, API consumption models). A significant undercurrent is the tension between rapid feature iteration and the accumulation of unresolved technical debt—particularly around Windows support, session lifecycle management, and account provisioning. Developers are increasingly vocal about expecting CLI tools to match the reliability of traditional developer tooling, with silent failures and opaque state transitions becoming the most cited sources of frustration.

---

## 2. Activity Comparison

| Tool | Notable Issues (last 24h) | PRs (last 24h) | Release Today |
|------|---------------------------|----------------|---------------|
| **Claude Code** | 10 (2 open critical) | 4 (1 open) | None |
| **OpenAI Codex CLI** | 10 (1 closed showstopper) | 10 (8 merged) | 2 alpha (`rust-v0.140.0-a.18, .19`) |
| **Gemini CLI** | 10 (4 P1 open) | 10 (2 merged, 8 open) | None |
| **GitHub Copilot CLI** | 6 (3 triage, 2 spurious) | 0 | **v1.0.62** |
| **Kimi Code CLI** | 1 (long-standing bug) | 1 (closed after 4 months) | None |
| **OpenCode** | 10 (4 critical/reliability) | 10 (3 merged, 7 open) | **v1.17.5, v1.17.6** |
| **Pi** | 10 (4 open, critical) | 8 (all merged) | None |
| **Qwen Code** | 10 (policy debate + 4 critical) | 10 (all merged) | None |
| **CodeWhale** (ex-DeepSeek TUI) | 10 (3 design epics) | 10 (9 merged) | **v0.8.60** (rebrand) |

**Key observations:**
- **High-throughput repos:** OpenAI Codex, Pi, Qwen Code, and CodeWhale merged 8–10 PRs today, indicating sustained engineering velocity.
- **Lowest activity:** Kimi Code CLI (1 issue, 1 PR) and Copilot CLI (0 PRs) show significantly less visible development.
- **Release cadence:** OpenCode, Copilot CLI, and CodeWhale shipped user-facing releases today; Codex shipped incremental alpha builds.

---

## 3. Shared Feature Directions

### Requirements appearing across ≥3 tool communities

| Requirement | Tools Requesting | Context |
|-------------|------------------|---------|
| **MCP tool preloading / lazy-load fix** | Copilot CLI (#3787), OpenCode (#32286 retry loops), Gemini CLI (OAuth refresh, tool schema), Pi (#5687 hangs) | MCP reliability is a cross-cutting pain point—tools not discovered, sessions stale, retry logic unbounded |
| **Multi-agent orchestration & fleets** | CodeWhale (#3154, #3096, #3167), OpenCode (#32166), Qwen Code (#4721 dynamic workflows), Claude Code (sub-agent patterns) | Growing demand for persistent, manager-worker agent architectures beyond single-turn interactions |
| **Agent safety & destructive action guards** | Gemini CLI (#22672), Claude Code (#67722 paid scripts), Qwen Code (#5102 probe side-effects), OpenCode (#8463 YOLO mode) | Community asking for explicit safety layers—permission contracts, rollback capabilities, destructive-action warnings |
| **Cross-platform / Windows parity** | Claude Code (Unicode, `/desktop`, process reaping), Codex (WSL path mapping, CLI binary), Kimi (#839 Windows shell), OpenCode (CJK, tab titles) | Windows and WSL remain the most fragile platforms; tool-specific regressions compound |
| **Persistent / long-running workflows** | CodeWhale (#891 `/goal`), OpenCode (#29059 dynamic workflows), Claude Code (Routines notifications), Copilot CLI (skill execution) | Users want agents that work toward multi-step objectives without continuous human oversight |
| **Billing / usage transparency** | Claude Code (weekly limit reset, API fallback billing), Qwen Code (#3203 free tier slash), OpenCode (#32278 payment stuck), Pi (#5703 cache cost inflation) | Usage tracking and cost attribution remain opaque, eroding trust |

### Secondary shared themes
- **Voice input integration:** Codex (#3000, 172 👍) and CodeWhale (#3051, already merged) both invest in speech-to-text
- **Theme/customization:** Claude Code (project-theme PR #68239), Gemini (#27887 border fix), OpenCode (#32295 cursor config)
- **OAuth/free tier churn:** Qwen Code (#3203, 134 comments), Codex (#20161 phone verification), Pi (#5714 Grok OAuth) all dealing with authentication flow maturity

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | OpenCode | Qwen Code | CodeWhale |
|-----------|-------------|--------------|------------|----------|-----------|-----------|
| **Primary focus** | Plugin ecosystem, billing systems | Rust alpha client, turn queuing | Agent reliability, safety, AIQ evals | Dynamic workflows, skill hot-reload | Multi-provider support, policy debate | Agent fleet architecture, persistent goals |
| **Target user** | Max-plan subscribers, professional devs | Early adopters, Rust/TS devs | Cloud/Google ecosystem devs | Plug-n-play, community-driven | Chinese market, OAuth free tier users | Power users, agent-orchestration enthusiasts |
| **Technical approach** | Hook/plugin laden, 3rd-party extensions | Incremental alpha→stable migration | Sub-agent orchestration, behavioral evals | MCP-first, provider-agnostic | Multi-provider/region, skill system | Headless worker runtime, durable ledgers |
| **Community engagement style** | Bug-report driven, stale issue backlog | High engagement on authentication | Active P1 triage, internal eval investment | Rapid feature PRs, YOLO feature requests | Policy debate (134 comments), Pro plan demand | Design epics, architecture discussions |
| **Windows support quality** | Poor (Unicode, process reaping, npm AVX) | Mixed (WSL path mapping broken) | Unknown (not flagged in digest) | Mixed (CJK rendering, tab reset) | Not flagged | Poor (MSBuild FileTracker fails) |
| **Maturity signal** | Large backlog of closed/stale issues | Alpha releases, rapid protocol iteration | PR-driven, robust eval infrastructure | Two patch releases in one day | Free tier policy adjustment (mature product phase) | Rebranding signals deliberate product identity |

---

## 5. Community Momentum & Maturity

### High-velocity, high-engagement
- **OpenCode:** Most active release cadence (2 patches today). Community actively filing detailed feature proposals (5-level nested agents). 70-👍 YOLO mode signals strong power-user demand.
- **CodeWhale:** Rebranding and architectural shift toward agent fleets. 10 PRs merged, 3 active epics. Deliberate "gate" on `/swarm` until runtime is production-grade—shows product discipline.
- **Pi:** All 8 PRs merged today. Strong extension ecosystem (#5714 Grok OAuth, #5385 theme detection). Architectural debt recognized (#5653 shrinkwrap).

### Rapid iteration, early-stage
- **OpenAI Codex:** Two alpha releases today. Durable turn queuing nearing completion (5 related PRs). High community engagement on authentication (#20161, 125 👍).
- **Qwen Code:** Mature product facing strategic inflection (free tier reduction). High PR throughput (10 merged). Community polarized on pricing.

### Stable but reactive
- **Claude Code:** Largest issue corpus but many stale. 0 releases today. Focused on bugfix PRs (#67722 paid scripts). Billing confusion remains top pain point.
- **GitHub Copilot CLI:** Lowest activity of the major tools. 1 release (v1.0.62, minor UX). Community small but exhibiting quality signal (session poisoning, skill path bugs).

### Stagnation signals
- **Kimi Code CLI:** 1 issue and 1 PR updated today; the bug (#640) has been open 5 months. Minimal maintainer engagement.

---

## 6. Trend Signals

### Industry trends from community feedback

1. **Agent orchestration is the next frontier.** Every major community (CodeWhale, OpenCode, Qwen, Claude Code) has active discussions about multi-agent fleets, persistent goals, and manager-worker patterns. The single-turn chat paradigm is being abandoned in favor of autonomous, checkpointed, multi-step workflows.

2. **MCP is essential but immature.** Every tool supporting MCP reports reliability issues—session expiry, lazy discovery, duplicate tool calls, stale connections. The protocol's extensibility is valued, but production readiness is inconsistent. "MCP tool preloading" is a shared cross-cutting pain point.

3. **Platform parity is failing Windows.** Claude Code (Unicode, process reaping, npm), Codex (WSL path mapping), CodeWhale (MSBuild), and OpenCode (CJK, tab reset) all show Windows as the most fragile platform. This creates a real ecosystem blindspot given the enterprise developer base.

4. **Safety is becoming table-stakes.** Silent side-effect execution (Qwen #5102), command injection fixes (Pi #3140, Gemini #27575), and destructive action guards (Gemini #22672) are no longer "nice to have." Communities now expect explicit permission contracts and rollback capabilities.

5. **Usage transparency is broken.** Billing confusion crosses all monetization models: subscription (Claude Code weekly limit reset), API (Qwen daily quota), and OAuth (Codex phone verification). Users want real-time, auditable, predictable cost attribution.

6. **Voice input is emerging as a differentiator.** Codex (172 👍 on #3000) and CodeWhale (already shipped) are investing in speech-to-text. This may become a distinguishing feature for tools targeting mobile/headless workflows.

### Reference value for developers

- **If you build on Windows:** Expect friction. Prefer OpenCode or Copilot CLI (moderate support) over Claude Code or CodeWhale (poor).
- **If you need persistent agent workflows:** CodeWhale (Agent Fleet, `/goal`), OpenCode (dynamic workflows), and Qwen Code (multi-agent) are the most advanced tracks.
- **If billing predictability matters:** Monitor Claude Code (risk of unexpected API billing) and Qwen Code (free tier reduction). OpenCode's transparent "YOLO mode" discussion suggests more honesty about cost.
- **If you value stability over novelty:** Copilot CLI (lowest churn, least activity) and Pi (steady PRs, architectural focus) are conservative choices.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
**Data snapshot:** github.com/anthropics/skills | Date: 2026-06-14

---

## 1. Top Skills Ranking

The following PRs received the most discussion activity (sorted by comment volume). All remain **open** as pending contributions.

### 1.1. document-typography ([#514](https://github.com/anthropics/skills/pull/514))
A skill that prevents common typographic defects in AI‑generated documents: orphan word wrap, widow paragraphs, and numbering misalignment. The discussion centered on whether this is a core formatting fix versus a situational preference. **Status:** Open, community feedback supported the approach.

### 1.2. ODT skill ([#486](https://github.com/anthropics/skills/pull/486))
Adds full support for OpenDocument Format files (.odt, .ods) – creation, template filling, and conversion to HTML. The thread debated whether ODT support should be merged into the existing DOCX skill or kept separate. **Status:** Open, awaiting maintainer decision on scoping.

### 1.3. frontend-design clarity & actionability ([#210](https://github.com/anthropics/skills/pull/210))
A revision of the frontend-design skill to make instructions more specific and executable within a single Claude conversation. Reviewers highlighted the need for concrete examples and conflict resolution with other design-oriented skills. **Status:** Open, iterative feedback from multiple contributors.

### 1.4. skill-quality-analyzer / skill-security-analyzer ([#83](https://github.com/anthropics/skills/pull/83))
Meta‑skills that assess other skills across five quality dimensions and security patterns. The discussion revealed strong interest in tooling that validates skills before submission, but concerns were raised about circular dependencies. **Status:** Open, design refinements ongoing.

### 1.5. agent-creator skill + multi-tool evaluation fix ([#1140](https://github.com/anthropics/skills/pull/1140))
A meta‑skill for creating task-specific agent sets, with fixes for parallel tool evaluation and Windows path support. This PR attracted comments around the interaction between meta‑skills and the core skill‑creation pipeline. **Status:** Open, recently updated.

### 1.6. testing-patterns skill ([#723](https://github.com/anthropics/skills/pull/723))
A comprehensive testing skill covering the Testing Trophy model, unit tests, React component tests, and integration patterns. The conversation focused on skill length constraints and whether to split into multiple specialized skills. **Status:** Open, waiting for modularization proposal.

### 1.7. AURELION skill suite ([#444](https://github.com/anthropics/skills/pull/444))
Four skills implementing a structured cognitive framework (kernel, advisor, agent, memory). Reviewers praised the depth but questioned the maintainability of a multi‑skill suite in a single PR. **Status:** Open, maintainers requested a split.

### 1.8. shodh-memory skill ([#154](https://github.com/anthropics/skills/pull/154))
A persistent memory system that maintains context across conversations via `proactive_context` calls. The thread highlighted overlap with existing memory skills and the need for a unified memory abstraction. **Status:** Open, discussion ongoing.

---

## 2. Community Demand Trends

Analysis of the most‑commented **Issues** reveals clear unmet needs:

| # | Issue | Comments | 👍 | Core Demand |
|---|-------|----------|----|-------------|
| 228 | [Org‑wide skill sharing](https://github.com/anthropics/skills/issues/228) | 14 | 7 | Enterprise distribution: direct sharing links / shared skill library |
| 556 | [`run_eval.py` 0% recall bug](https://github.com/anthropics/skills/issues/556) | 12 | 7 | **Skill‑creator tooling reliability** – optimizer loop is broken for all users |
| 62 | [Skills disappearing after rename](https://github.com/anthropics/skills/issues/62) | 10 | 2 | Skill persistence & file management UX |
| 202 | [skill-creator best practice update](https://github.com/anthropics/skills/issues/202) | 8 | 1 | **Tooling quality** – current skill‑creator reads like docs, not an actionable skill |
| 492 | [Namespace security vulnerability](https://github.com/anthropics/skills/issues/492) | 7 | 2 | **Security / trust boundary** – community skills under `anthropic/` impersonation |
| 412 | [Agent governance skill proposal](https://github.com/anthropics/skills/issues/412) | 6 | 0 | Governance & safety patterns for multi‑agent systems |
| 189 | [Duplicate skills from plugin overlap](https://github.com/anthropics/skills/issues/189) | 6 | 8 | **Plugin dependency management** – skills installed twice |
| 184 | [agentskills.io redirect error](https://github.com/anthropics/skills/issues/184) | 3 | 4 | Documentation infrastructure reliability |
| 1061 | [Windows compatibility blocker](https://github.com/anthropics/skills/issues/1061) | 3 | 0 | **Cross‑platform parity** for skill‑creator scripts |
| 1169 | [`recall=0%` in optimisation loop](https://github.com/anthropics/skills/issues/1169) | 3 | 1 | Repeat of #556 – underscores criticality of tooling fix |

**Dominant themes:**
- **Skill‑creator tooling reliability** is the #1 pain point – three issues (#556, #1169, #1061) all report the same `recall=0%` bug, blocking skill optimization.
- **Enterprise features** – org‑wide sharing (#228) and security boundaries (#492) signal growing enterprise adoption.
- **Cross‑platform support** – multiple Windows‑specific issues (#1061, #62) indicate a significant Windows user base facing fundamental incompatibility.
- **Governance & safety** – the agent‑governance proposal (#412) shows demand for structured safety patterns as agent usage scales.

---

## 3. High-Potential Pending Skills

These PRs have active discussion and are likely to merge soon, either as‑is or after revision:

| PR | Skill | Key Blocker | Latest Activity |
|----|-------|-------------|-----------------|
| [#1140](https://github.com/anthropics/skills/pull/1140) | agent‑creator meta‑skill | Multi‑tool evaluation fix required | Updated June 2 |
| [#539](https://github.com/anthropics/skills/pull/539) | YAML special character detection | None – straightforward fix | Updated April 16 |
| [#541](https://github.com/anthropics/skills/pull/541) | DOCX tracked change ID collision | None – clear root cause | Updated April 16 |
| [#509](https://github.com/anthropics/skills/pull/509) | CONTRIBUTING.md | Community health metric gap | Updated March 19 |
| [#1050](https://github.com/anthropics/skills/pull/1050) | Windows subprocess + encoding fix | Critical for Windows users | Updated May 24 |

**Notable:** Multiple PRs target the same `run_eval.py`/`run_loop.py` Windows breakage (#1099, #1298, #1050). This cluster suggests an imminent resolution – the first cleanly‑reviewed Windows fix will likely merge rapidly.

---

## 4. Skills Ecosystem Insight

**The community’s most concentrated demand is for *reliable, cross‑platform skill‑authoring tooling* – the skill‑creator scripts must stop returning false‑negative trigger rates and must work on Windows – before any new skill submissions can be meaningfully evaluated.**

More broadly, the ecosystem shows a healthy split: half the PRs introduce **new domain skills** (typography, ODT, testing patterns, governance), while the other half improve **infrastructure** (skill‑creator fixes, validation, Windows support). The latter currently attracts more discussion because without functioning tooling, new skills cannot be tested or optimized.

---

# Claude Code Community Digest
**2026-06-14**

## Today's Highlights
This week’s digest surfaces two critical open bugs—remote-control slash commands being routed as plain text and unreaped process trees on Windows—as well as a wave of closed issues confirming long-standing billing and platform regression patterns. No new releases were published today.

## Releases
No new versions were released in the last 24 hours.

## Hot Issues (10 Noteworthy)

1. **[BUG] Remote Control: slash commands route to model as plain text** `#68252` *(OPEN)*  
   ⭐ 4 upvotes | 🔥 New  
   When driving a local session via `/rc` from mobile/desktop, commands like `/clear` or `/compact` are intermittently sent to the model as plain text instead of being executed. This breaks a core remote-control workflow.  
   [→ Issue](https://github.com/anthropics/claude-code/issues/68252)

2. **[BUG] Windows local-agent-mode: session processes + MCP fleets not reaped** `#68394` *(OPEN)*  
   ⭐ 0 upvotes | 🔥 New  
   Each launched session leaves `claude.exe` and its MCP server fleet running, accumulating processes across launches. Memory pressure and zombie processes reported on Windows.  
   [→ Issue](https://github.com/anthropics/claude-code/issues/68394)

3. **[BUG] Weekly limits counter resets on 24-hour cycle** `#52921` *(CLOSED, stale)*  
   📌 9 comments | ❌ Closed  
   Max-plan users reported their usage counter resets daily instead of every 7 days, causing premature throttling. A long-standing billing grievance.  
   [→ Issue](https://github.com/anthropics/claude-code/issues/52921)

4. **[BUG] Claude Code bills API instead of Max subscription** `#58625` *(CLOSED, stale)*  
   📌 9 comments | ❌ Closed  
   OAuth response missing `organizationType` leads to fallback API billing for Max subscribers. Linked to broader auth/account handling issues.  
   [→ Issue](https://github.com/anthropics/claude-code/issues/58625)

5. **[BUG] Edit Tool Unicode-to-ASCII corruption on Windows (.md files)** `#29699` *(CLOSED, stale)*  
   ⭐ 3 upvotes | ❌ Closed  
   The `Edit` tool mangles Unicode characters to ASCII on Windows when editing Markdown files. Severe for non-English projects.  
   [→ Issue](https://github.com/anthropics/claude-code/issues/29699)

6. **[BUG] SessionStart hook fires before TUI is ready** `#22936` *(CLOSED, stale)*  
   📌 7 comments | ❌ Closed  
   Hooks executing prematurely cause lost configuration. Workaround exists but root cause never fixed.  
   [→ Issue](https://github.com/anthropics/claude-code/issues/22936)

7. **[BUG] Skill file "Save and Replace" silently fails in Cowork** `#46844` *(CLOSED, stale)*  
   ⭐ 1 upvote | ❌ Closed  
   Skills packaged from host-mounted paths (with old files) instead of VM working copies, causing cascading failures. Critical for pipeline reliability.  
   [→ Issue](https://github.com/anthropics/claude-code/issues/46844)

8. **[BUG] `claude -p` rejected with "out of extra usage" while Max shows 35% remaining** `#57096` *(CLOSED, stale)*  
   📌 5 comments | ❌ Closed  
   Duplicate of #45203; Max users blocked from `-p` despite visible quota. Likely a stale cache or rate-limiting logic issue.  
   [→ Issue](https://github.com/anthropics/claude-code/issues/57096)

9. **[BUG] `head -n` macOS command destroyed production DB manifests** `#59664` *(CLOSED, stale)*  
   📌 2 comments | ❌ Closed  
   Incompatible `head -n -10` on macOS (Linux syntax) caused irreversible data loss in LanceDB. Highlights lack of platform-aware command generation.  
   [→ Issue](https://github.com/anthropics/claude-code/issues/59664)

10. **[BUG] npm install broken on non-AVX Linux x64 CPUs** `#51976` *(CLOSED, stale)*  
    ⭐ 2 upvotes | ❌ Closed  
    Since v2.1.113, Node.js fallback removed, breaking installation on older CPUs. Regression flagged by community.  
    [→ Issue](https://github.com/anthropics/claude-code/issues/51976)

## Key PR Progress (4 Total)

1. **[#67722] Claude autonomously ran background scripts calling paid external** *(CLOSED)*  
   Pull request to fix an issue where Claude Code ran scripts that triggered paid external services without user consent. Described as a bug fix.  
   [→ PR](https://github.com/anthropics/claude-code/pull/67722)

2. **[#1] Create SECURITY.md** *(CLOSED)*  
   Repository foundation document, no functional change.  
   [→ PR](https://github.com/anthropics/claude-code/pull/1)

3. **[#68239] feat: add project-theme plugin for per-project theme settings** *(OPEN)*  
   Adds a `SessionStart` hook that reads theme/color from `.claude/settings.json`. Addresses long-standing request #43216. Community-positive addition.  
   [→ PR](https://github.com/anthropics/claude-code/pull/68239)

4. **[#58673] s** *(OPEN)*  
   Minimal PR with description "s". Likely a test or placeholder.  
   [→ PR](https://github.com/anthropics/claude-code/pull/58673)

## Feature Request Trends

- **Plugin & Hooks Ecosystem** – Demand for richer plugin APIs includes configurable namespace aliases (`#59069`), hook events for recap text (`#59429`), and the ability to render animated/interactive UI in tool output (`#52235`).  
- **Routine Notifications** – “Notify me” from Claude Routines (`#54994`) to push alerts when background tasks complete.  
- **Per-Project Persistence** – Project-level theme/color settings (`#43216`, now addressed in PR #68239) and user configurable plugin aliases show a desire for more flexible customization.  
- **IDE Integration** – Option to disable Up/Down arrow history recall in VS Code (`#51202`, 7 upvotes) indicates friction between terminal muscle memory and editor controls.

## Developer Pain Points

- **Billing & Subscription Confusion** – Weekly limits resetting on a 24-hour cycle and API billing for Max subscribers are the most-reported frustrations, eroding trust in usage tracking.  
- **Platform-Specific Regressions** – Windows users face Unicode corruption and missing `/desktop` commands; macOS users suffer incompatible system commands (`head -n`) and memory leaks; Linux users see broken installs on non-AVX CPUs.  
- **Session Lifecycle Bugs** – SessionStart hooks firing before TUI readiness, Cowork VMs packaging stale files, and unreaped processes on Windows all point to incomplete teardown and state management.  
- **Silent Failures** – “Reinstall” button doing nothing, “Save and Replace” silently failing, and managed settings being silently ignored are recurring pain points that erode confidence.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# Codex Community Digest – 2026-06-14

## Today’s Highlights
Two more `rust-v0.140.0` alpha releases landed, continuing the iteration toward a stable Rust client. On the issue tracker, an **authentication showstopper** (#20161, 200 comments, 125 👍) was closed today after weeks of community outcry about broken phone‑number verification. Merged PRs signal that **durable turn queuing** (app‑server + TUI) is nearing completion, while new open issues highlight **cross‑platform path‑mapping regressions** on Windows/WSL and **macOS syspolicyd resource exhaustion**.

## Releases
- **`rust-v0.140.0-alpha.18`** (2026-06-14) – alpha release  
  [`openai/codex` Releases](https://github.com/openai/codex/releases)
- **`rust-v0.140.0-alpha.19`** (2026-06-14) – alpha release  
  [`openai/codex` Releases](https://github.com/openai/codex/releases)

No detailed changelogs were published; both are incremental internal builds.

## Hot Issues (10 selected)

| # | Issue | Why It Matters | Community Reaction |
|---|-------|----------------|-------------------|
| [#20161](https://github.com/openai/codex/issues/20161) | [CLOSED] Phone number verification failure | Blocked SSO login on new devices; a two‑month‑old blocker finally resolved. | 200 comments, 125 👍 – highest engagement in the repo. |
| [#20552](https://github.com/openai/codex/issues/20552) | View > Toggle File Tree unreliable on macOS | Core UX frustration; users cannot reliably show/hide the file panel. | 44 comments, 14 👍 |
| [#3000](https://github.com/openai/codex/issues/3000) | Voice transcription for IDE extension | Most‑upvoted feature request (172 👍); push‑to‑talk would greatly speed up prompting. | 31 comments, 172 👍 – strong demand. |
| [#25243](https://github.com/openai/codex/issues/25243) | macOS relaunch loop exhausts `syspolicyd` FDs | Blocks app launches entirely; affects Pro users. | 26 comments, 2 👍 |
| [#22423](https://github.com/openai/codex/issues/22423) | CLI binary not found; `CODEX_CLI_PATH` required | Breaks Electron desktop app on WSL setups; unclear resolution path. | 25 comments |
| [#23122](https://github.com/openai/codex/issues/23122) | [CLOSED] Android QR code opens unhandled link | Mobile setup stuck in a loop because deep links aren’t routed to the ChatGPT Android app. | 20 comments, 16 👍 |
| [#27979](https://github.com/openai/codex/issues/27979) | Windows app no longer opens after June 12 update | Full breakage on Win 11; “About” dialog unreachable. | 19 comments, 3 👍 |
| [#19504](https://github.com/openai/codex/issues/19504) | Add full RTL support for Arabic & Hebrew | Accessibility and usability gap for right‑to‑left language users. | 14 comments, 12 👍 |
| [#19365](https://github.com/openai/codex/issues/19365) | Windows: Browser Use unavailable because Node REPL tool not exposed | Feature parity gap – Windows users cannot use the in‑app browser tool due to missing runtime component. | 10 comments, 14 👍 |
| [#27694](https://github.com/openai/codex/issues/27694) | macOS crash: `CodexDockTilePlugin` recursion | Dock icon customisation causes infinite recursion; crashes on macOS 26.5. | 8 comments, 4 👍 |

## Key PR Progress (10 selected)

| # | PR | Status | Description |
|---|-----|--------|-------------|
| [#28151](https://github.com/openai/codex/pull/28151) | pipeline Windows targets separately | OPEN | Splits x64 and ARM64 packaging into independent pipelines to avoid idle wait times. |
| [#23620](https://github.com/openai/codex/pull/23620) | feat: dispatch queued turns from app-server | CLOSED | Merges the queue dispatching logic – queued follow‑ups now run serially after active turn completes. |
| [#24987](https://github.com/openai/codex/pull/24987) | feat(mcp): load pending tools through lazy search | CLOSED | Defers optional MCP server initialization until `tool_search` is actually invoked, reducing startup latency. |
| [#25276](https://github.com/openai/codex/pull/25276) | refactor: separate turn submissions from thread setting overrides | CLOSED | Clean protocol separation so queued turns don’t copy session‑level overrides. |
| [#25258](https://github.com/openai/codex/pull/25258) | feat: queue TUI follow‑ups through app-server | CLOSED | Proof‑of‑concept for TUI queuing behind `app_server_queue` flag. |
| [#25340](https://github.com/openai/codex/pull/25340) | feat(tui): animate startup header mascot | CLOSED | Adds a Unicode mascot (14×6) to the TUI banner – lightweight, no image protocols. |
| [#23619](https://github.com/openai/codex/pull/23619) | feat: expose queued turn app-server RPCs | CLOSED | New API endpoints for add/list/delete/reorder queued turns; introduces `thread/queue/change` notifications. |
| [#23618](https://github.com/openai/codex/pull/23618) | feat: add durable queued turn store | CLOSED | FIFO database layer for follow‑up turns; foundational for the queue feature. |
| [#24124](https://github.com/openai/codex/pull/24124) | feat(tui): add usage report command | CLOSED | `/usage` command shows token consumption per feature with daily/weekly views. |
| [#27955](https://github.com/openai/codex/pull/27955) | [codex] retain resolved turn environments in session state | OPEN | Avoids re‑resolving environment IDs on every turn; reduces MCP re‑initialization overhead. |

## Feature Request Trends
- **Voice input** – Issue #3000 (172 👍) and related #418 push for Push‑to‑Talk in both the IDE extension and CLI.
- **RTL / localization** – #19504 (Arabic/Hebrew RTL) and #24827 / #22846 (Chinese language selector not persisting) show strong demand for non‑English UX.
- **Multi‑tab in‑app browser** – #23314 (6 👍) requests visible tabs instead of a single‑view browser.
- **Sidebar project management** – #26026 asks for ability to delete stale project entries from the sidebar.
- **Environment overlays** – #27336 proposes named local environment profiles that can be switched per turn.

## Developer Pain Points
- **Cross‑platform path mapping** – Multiple Windows/WSL issues (#28094, #28172, #28174, #19365) where WSL paths like `/home/…` are rewritten to `C:\home\…`, breaking project associations and tool‑call execution.
- **Authentication and setup** – Phone verification (#20161), SSO loops (#23122), and CLI binary path discovery (#22423) continue to plague first‑time and cross‑device usage.
- **macOS stability** – `syspolicyd` FD exhaustion (#25243), DockTile crash (#27694), and Capturing appshot failures (#25269) indicate ongoing macOS integration issues.
- **Language setting persistence** – #24827 and #22846 report that switching to Chinese (or other locales) fails to stick after restart on both Windows and macOS.
- **File mutation deadlocks** – #24407 documents a non‑deterministic deadlock in `apply_patch`/`file_change` that leaves turn completions missing, affecting CLI reliability.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

Here is the **Gemini CLI Community Digest** for **2026-06-14**.

---

## Gemini CLI Community Digest – 2026-06-14

### Today’s Highlights
The community is focused on **agent reliability and safety**. A critical P1 bug where the generalist agent hangs indefinitely remains the most upvoted open issue. The team is actively patching infrastructure, landing fixes for the MCP OAuth refresh flow, command injection vulnerabilities, and a stack overflow caused by regex backtracking. The AI evaluation framework continues to be a major area of internal investment, with a new epic for component-level evaluations.

---

### Releases
No new releases in the last 24 hours.

---

### Hot Issues

1.  **[Generalist agent hangs indefinitely](https://github.com/google-gemini/gemini-cli/issues/21409)**
    - **Why it matters:** The most upvoted open bug (8 👍). Users report the CLI hangs when deferring to the generalist agent for simple tasks like folder creation. Working around it by disabling sub-agents indicates a core agent orchestration failure.
    - **Status:** P1, Needs Retesting.

2.  **[Subagent recovery after MAX_TURNS is falsely reported as success](https://github.com/google-gemini/gemini-cli/issues/22323)**
    - **Why it matters:** A misleading feedback loop. The `codebase_investigator` subagent reports `status: "success"` even when it hit its turn limit before doing any analysis. This degrades trust in the agent’s self-reporting.
    - **Status:** P1.

3.  **[Shell command execution gets stuck waiting for input](https://github.com/google-gemini/gemini-cli/issues/25166)**
    - **Why it matters:** A common and frustrating UX bug. After a command completes, the CLI remains in a "Waiting input" state. This happens for simple commands that don't require input.
    - **Status:** P1, Effort/Medium.

4.  **[Browser agent fails on Wayland](https://github.com/google-gemini/gemini-cli/issues/21983)**
    - **Why it matters:** Linux users on Wayland (a growing demographic) are unable to use the browser subagent. The error is cryptic, showing a "GOAL" termination reason despite failing.
    - **Status:** P1, Needs Retesting.

5.  **[Robust component level evaluations (Epic)](https://github.com/google-gemini/gemini-cli/issues/24353)**
    - **Why it matters:** A P1 epic following up on the introduction of behavioral evals. The team has generated 76 tests across 6 Gemini models and now needs to make them robust and systematic.
    - **Status:** P1, AIQ/Eval Infra.

6.  **[Model creates tmp scripts in random spots](https://github.com/google-gemini/gemini-cli/issues/23571)**
    - **Why it matters:** Users report significant workspace pollution. The model writes multiple temporary scripts across the project tree, creating a messy state that is hard to clean for a clean commit.
    - **Status:** P2.

7.  **[Auto Memory retries low-signal sessions indefinitely](https://github.com/google-gemini/gemini-cli/issues/26522)**
    - **Why it matters:** An efficiency and privacy concern. The Auto Memory system keeps retrying sessions the model deemed "low-signal," wasting tokens and potentially re-exposing user data.
    - **Status:** P2.

8.  **[Gemini doesn't use skills and sub-agents enough](https://github.com/google-gemini/gemini-cli/issues/21968)**
    - **Why it matters:** An anecdotal but critical P2 behavioral issue. Users report the model ignores custom skills and sub-agents unless explicitly instructed, defeating the purpose of user-created automation.
    - **Status:** P2, Needs Retesting.

9.  **[400 error when > 128 tools are available](https://github.com/google-gemini/gemini-cli/issues/24246)**
    - **Why it matters:** A scaling limitation. Users with many MCP servers or skills hit an API error. The agent needs smarter tool selection instead of passing all tools to the model.
    - **Status:** P2.

10. **[Agent should stop/discourage destructive behavior](https://github.com/google-gemini/gemini-cli/issues/22672)**
    - **Why it matters:** A safety concern. The model can use `git reset --force` or other destructive commands when safer alternatives exist. This is a request for a safety layer that understands the risks of modifying production resources.
    - **Status:** P2, Customer Issue.

---

### Key PR Progress

1.  **[fix(security): prevent command injection in findCommand](https://github.com/google-gemini/gemini-cli/pull/27575) [CLOSED]**
    - **Impact:** Replaces shell-interpolated `execSync` with safe `spawnSync`. This directly prevents command injection via shell metacharacters. A critical security improvement.
    - **Status:** Closed (Merged).

2.  **[fix(at-command): prevent stack overflow from regex backtracking](https://github.com/google-gemini/gemini-cli/pull/27580) [CLOSED]**
    - **Impact:** Fixes a crash caused by catastrophic backtracking on large pasted inputs. Addresses issue #27539, improving stability for power users.
    - **Status:** Closed (Merged).

3.  **[fix(core): cap pending tool responses](https://github.com/google-gemini/gemini-cli/pull/27870) [OPEN]**
    - **Impact:** Fixes issue #27738 where a very large tool result could hang or crash the CLI by capping the pending `functionResponse` payload. A P1 fix for agent reliability.
    - **Status:** Open, P1.

4.  **[fix(core): refresh MCP OAuth with stored client ID](https://github.com/google-gemini/gemini-cli/pull/27889) [OPEN]**
    - **Impact:** Fixes the MCP OAuth refresh flow when an auto-discovered server has a dynamic client ID. Without this, token refresh silently fails.
    - **Status:** Open, P1.

5.  **[fix(core): normalize MCP tool schemas to root type object](https://github.com/google-gemini/gemini-cli/pull/27888) [OPEN]**
    - **Impact:** Some MCP servers lack a root `type: "object"` in their tool schemas. This fix ensures compatibility with strict downstream APIs (like Vertex AI), preventing 400 errors.
    - **Status:** Open, P2.

6.  **[fix(core): respect .gitignore and .geminiignore in session_context](https://github.com/google-gemini/gemini-cli/pull/27886) [OPEN]**
    - **Impact:** Fixes a bug where ignored files (e.g., `node_modules`, secrets) could leak into the session context directory. This improves privacy and prevents unwanted file uploads.
    - **Status:** Open, P2.

7.  **[fix(cli): honor custom theme border.default](https://github.com/google-gemini/gemini-cli/pull/27887) [OPEN]**
    - **Impact:** Fixes a cosmetic bug in theme customization. The documented `border.default` setting was being ignored on some terminals. Improves the developer experience.
    - **Status:** Open, P2.

8.  **[fix(vscode-ide-companion): register all activate() disposables](https://github.com/google-gemini/gemini-cli/pull/27885) [OPEN]**
    - **Impact:** Fixes a resource leak in the VS Code extension where two activation disposables were never properly unregistered. Essential for long-running IDE sessions.
    - **Status:** Open, P2.

9.  **[fix(core): add image-grounding hint](https://github.com/google-gemini/gemini-cli/pull/27711) [OPEN]**
    - **Impact:** Addresses an issue where image attachments were not properly grounded in function responses. Improves the reliability of multi-modal interactions.
    - **Status:** Open, Size/M.

10. **[fix(trust): disclose hooks declared in canonical nested shape](https://github.com/google-gemini/gemini-cli/pull/27903) [OPEN]**
    - **Impact:** Fixes a security dialog bug. Previously, the folder-trust prompt could hide hook commands defined in a nested (canonical) JSON shape, meaning users might trust a folder without seeing all running code.
    - **Status:** Open, Size/S.

---

### Feature Request Trends

- **Robust Evaluation Infrastructure (AIQ):** The highest priority "AIQ" work is formalizing the evaluation framework. Issues like #24353 and #23166 show a push to stabilize behavioral evals, implement component-level testing, and make results automatically actionable.
- **AST-Aware Code Understanding:** Multiple linked issues (#22745, #22746, #22747) are investigating Abstract Syntax Tree (AST) aware tools for file reads, search, and codebase mapping. The goal is to reduce token usage and improve the model's ability to find function boundaries.
- **Agent Safety and Guardrails:** A strong feature request theme is preventing destructive / irreversible actions. Issue #22672 is a request for explicit safety layers around git operations and database modifications.

---

### Developer Pain Points

- **Agent Hangs and Recovery:** The most painful recurring theme. Issues #21409 and #25166 describe scenarios where the agent hangs on simple operations, and recovery is slow or non-existent. The community is frustrated by the "falsely reported success" pattern (#22323).
- **Browser Agent Instability:** The browser subagent is a persistent pain point on Linux (Wayland) and ignores user configuration like `maxTurns` (#22267, #21983). Lock recovery (#22232) is also requested.
- **Memory System Overhead:** Developers are concerned about the Auto Memory system’s efficiency. Issues #26522 (indefinite retries), #26525 (secrets in logs), and #26523 (invalid patches) indicate a system that is generating significant overhead and potential privacy risks.
- **Sub-agent Management:** Users feel they lack control. Sub-agents are either running without permission (#22093) or not being used by the model when needed (#21968). The scaling limit of >128 tools (#24246) also creates a friction point for power users.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-06-14

## Today’s Highlights
Copilot CLI v1.0.62 shipped with improved dialog scrolling and blank-line handling in reasoning summaries. A critical new bug (#3791) shows that a malformed encrypted Excel attachment can permanently poison a CLI session, requiring restart. The community also active on long-standing skill execution path issues (#956) and two feature requests for bring-your-own-model and MCP tool preloading.

## Releases
**v1.0.62** (2026-06-13)
- Ask/elicitation dialogs now scroll together with the timeline instead of taking over the screen — tall dialogs no longer hide agent output.
- Blank lines are preserved between reasoning summary sections.
- User-typing indicator improved (full note cut off in source).

[Release tag](https://github.com/github/copilot-cli/releases/tag/v1.0.62)

## Hot Issues (6 total, all relevant)
All open issues updated in the last 24h are covered due to low volume.

### #956 — [area:agents] Agent skills scripts executed in wrong folder
**Author:** msundman78 | **Created:** 2026-01-13 | **Updated:** 2026-06-14 | **Comments:** 6 | 👍: 2  
**Why matters:** Breaks the official Agent Skills specification; a `scripts/myscript.sh` reference in SKILLS.md runs from an incorrect working directory. This has been open for 5 months and still unresolved, causing frustration for skill authors.  
[Issue #956](https://github.com/github/copilot-cli/issues/956)

### #3788 — [CLOSED, invalid] Spam/bug report with no details
**Author:** twinfire55002020infoorg-sudo | **Created/Closed:** 2026-06-13 | **Comments:** 1 | 👍: 0  
**Note:** Immediately closed as invalid. No actionable content.  
[Issue #3788](https://github.com/github/copilot-cli/issues/3788)

### #3791 — [triage] Malformed attachment poisons session; all subsequent turns fail with 400
**Author:** jay-tau | **Created:** 2026-06-14 | **Comments:** 0 | 👍: 0  
**Why matters:** A password-protected `.xlsx` causes a CAPI 400, and **every later turn in the same session also fails with 400** even when no attachment is sent. This is a severe session-poisoning bug that forces users to restart their terminal session. No official response yet.  
[Issue #3791](https://github.com/github/copilot-cli/issues/3791)

### #3793 — [triage] Gibberish title (hex strings), no bug description
**Author:** ja552588 | **Created:** 2026-06-14 | **Comments:** 0 | 👍: 0  
**Why matters:** Likely spam or a corrupted submission. No useful information.  
[Issue #3793](https://github.com/github/copilot-cli/issues/3793)

### #3789 — [triage] Request: Ollama API Key return to Bring Your Own Model
**Author:** Oncorporation | **Created/Updated:** 2026-06-13 | **Comments:** 0 | 👍: 0  
**Why matters:** Users want to use remote Ollama servers with Copilot CLI, but the current “Bring Your Own Model” menu lacks an API Key field for setting the host header. A workaround requires a forward proxy.  
[Issue #3789](https://github.com/github/copilot-cli/issues/3789)

### #3787 — [triage] Preload MCP server tools into the initial agent function list
**Author:** tamirdresher | **Created/Updated:** 2026-06-13 | **Comments:** 0 | 👍: 0  
**Why matters:** MCP tools registered via `.mcp.json` or `--additional-mcp-config` are lazy-loaded and not advertised in the agent’s initial `<available_tools>`. Agents that don’t actively probe for them will never see these tools, limiting extensibility.  
[Issue #3787](https://github.com/github/copilot-cli/issues/3787)

## Key PR Progress
No pull requests were updated in the last 24 hours.

## Feature Request Trends
Two distinct feature directions emerged from this batch of issues:
1. **Bring Your Own Model flexibility** — Users want to connect to arbitrary Ollama servers (local or remote) by providing an API key, not just the limited set of built-in models.  
2. **MCP tool preloading** — Developers building agents on top of Copilot CLI want MCP tools to be available from the start of a session, not lazily discovered. This would improve reliability for agent workflows that rely on custom tools.

## Developer Pain Points
- **Session poisoning by attachments** (#3791): A single malformed file can render an entire session unusable. No clear recovery path other than restarting the CLI.  
- **Agent skill execution path mismatch** (#956): Long-standing bug (6+ months) where referenced scripts run from the wrong working directory, breaking the official Agent Skills spec. Community engagement is low (only 2 upvotes), but the severity is high for skill authors.  
- **Lazy MCP tool discovery** (#3787): Non-obvious design choice that silently breaks agents expecting tools to be pre-announced.  
- **Spam/low-quality submissions** (#3788, #3793): Indicate that better issue templates or validation may be needed.

---

*Data source: [github.com/github/copilot-cli](https://github.com/github/copilot-cli) — Digest generated 2026-06-14.*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

## Kimi Code CLI Community Digest — 2026-06-14

### Today’s Highlights
Activity on the Kimi CLI repository was light today, with only one issue and one pull request updated in the last 24 hours. A long-standing bug causing the CLI to loop repeatedly over the same file (#640) saw renewed discussion after months of inactivity, while a PR adding configurable Windows shell support (#839) was finally closed after four months. No new releases were published.

### Releases
No new releases were published in the last 24 hours. The latest available version remains **0.76**.

### Hot Issues
*Only one issue was updated in the last 24 hours.*

1. **[#640] [bug] Kimi CLI stuck in reading one file again and again and stuck in a loop**  
   *Author: isbafatima90-arch | Created: 2026-01-19 | Updated: 2026-06-13 | Comments: 13*  
   The user reports that Kimi CLI (v0.76) enters an infinite loop repeatedly reading the same file when using a custom Anthropic endpoint (`mimo-v2-flash` model) on Arch Linux. The bug has been open since January and has accumulated 13 comments, suggesting it affects users relying on non‑standard model configurations. The lack of a fix after five months may indicate a deeper design issue or low maintainer prioritization.  
   [GitHub Issue #640](https://github.com/MoonshotAI/kimi-cli/issues/640)

### Key PR Progress
*Only one PR was updated in the last 24h.*

1. **[#839] feat(shell): add configurable shell support for Windows (CLOSED)**  
   *Author: HamzaETTH | Created: 2026-02-02 | Updated: 2026-06-14 | Closed: 2026-06-14*  
   This PR adds the ability to configure which shell Kimi CLI uses on Windows, likely addressing compatibility gaps with PowerShell, CMD, or WSL. The PR was closed today after four months of review, implying it has been merged or rejected (the summary does not specify the merge status). Cross‑platform shell configurability is a recurring need for teams using the CLI in mixed environments.  
   [GitHub PR #839](https://github.com/MoonshotAI/kimi-cli/pull/839)

### Feature Request Trends
Based on the limited activity, the most prominent theme is **cross‑platform shell configurability** (PR #839). Users want the flexibility to choose their default shell on Windows, which currently may be hard‑coded. No new feature requests were filed in the last 24 hours.

### Developer Pain Points
- **Infinite file‑reading loops** (Issue #640): A critical bug causing the CLI to become unusable when using custom endpoints or non‑default models. The prolonged open time (5 months) suggests either a tricky root cause or insufficient maintainer bandwidth.
- **Lack of Windows parity**: The closing of PR #839 indicates that Windows shell support was a gap; developers on Windows have likely been limited to a single shell until now.

*Note: With only two items updated today, this digest reflects an otherwise quiet day in the Kimi CLI community.*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-14

## Today’s Highlights
Two patch releases landed: **v1.17.6** improves MCP server compatibility by declaring client capabilities, and **v1.17.5** adds Snowflake Cortex OAuth, improves project copy management, and fixes MCP session recovery. The most upvoted issue remains the 70-👍 “YOLO mode” (`--dangerously-skip-permissions`), while multiple feature requests and PRs target **dynamic workflows** (inspired by Claude Code) and **hot-reloadable skills** — both of which are shaping the near-term roadmap.

---

## Releases

### [v1.17.6](https://github.com/anomalyco/opencode/releases/tag/v1.17.6)
- **Bugfix:** Declared OpenCode’s supported MCP client capabilities to improve compatibility with MCP servers.

### [v1.17.5](https://github.com/anomalyco/opencode/releases/tag/v1.17.5)
- **Improvements:**  
  - External browser OAuth for Snowflake Cortex provider (thanks @santigc6)  
  - Better project copy management and move-session flows in v2 layout  
- **Bugfixes:**  
  - Recovered expired MCP sessions instead of leaving MCP tools disconnected  
  - Cleared closed MCP clients to prevent stale connections  

---

## Hot Issues (10 selected)

1. **[#8463 – `--dangerously-skip-permissions` (YOLO mode)](https://github.com/anomalyco/opencode/issues/8463)**  
   *70 upvotes, 20 comments* — The most requested feature. Users want a flag to bypass permission prompts in trusted/automated environments. Community is actively discussing safety guardrails.

2. **[#29059 – Dynamic workflows for repeatable multi-step automation](https://github.com/anomalyco/opencode/issues/29059)**  
   *15 upvotes, 11 comments* — Adapting the popular “dynamic workflows” pattern from Claude Code. Several duplicates exist (#30308), showing strong demand for project-local automation.

3. **[#31755 – MiniMax direct API caching broken by new thinking toggle](https://github.com/anomalyco/opencode/issues/31755)**  
   *8 comments* — Regression where MiniMax M3 caching no longer works while OpenRouter BYOK still caches correctly. Users rely on caching to reduce usage limits.

4. **[#23595 – `<system-reminder>` keeps moving, breaking prompt cache in llama.cpp](https://github.com/anomalyco/opencode/issues/23595)**  
   *8 upvotes, 3 comments* — A subtle performance issue: repositioning the system reminder invalidates the prompt cache, causing excessive reprocessing in local models.

5. **[#29871 – Same remote URL projects share session ID](https://github.com/anomalyco/opencode/issues/29871)**  
   *2 upvotes, 3 comments* — Multiple clones of the same repository get the same session, making independent work impossible. Comes up often for monorepo users.

6. **[#32172 – Add GLM-5.2 model for Z.AI provider](https://github.com/anomalyco/opencode/issues/32172)**  
   *6 comments* — Z.AI released a new reasoning model; users want it available in OpenCode’s provider list. Quick community engagement.

7. **[#32166 – Nested sub-agent spawning (up to 5 levels) + multi-agent orchestration](https://github.com/anomalyco/opencode/issues/32166)**  
   *2 comments* — A detailed proposal backed by a working fork. Requests design review for deep agent trees and orchestration layers.

8. **[#32278 – Paid OpenCode Go ($10) stuck on “Insufficient balance”](https://github.com/anomalyco/opencode/issues/32278)**  
   *2 comments* — Payment processing issue blocking usage. Affects new subscribers immediately.

9. **[#32294 – Stuck after one command](https://github.com/anomalyco/opencode/issues/32294)**  
   *1 comment* — User must restart after every single command. Likely a race condition in session handling.

10. **[#32286 – OpenCode repeatedly retries after mid-tool provider disconnect](https://github.com/anomalyco/opencode/issues/32286)**  
    *1 comment* — 588 failed provider requests in one session due to unbounded retry logic. Urgent reliability concern.

---

## Key PR Progress (10 selected)

1. **[#32295 – feat(tui): add cursor style configuration](https://github.com/anomalyco/opencode/pull/32295)**  
   New PR that closes #11738. Adds configurable cursor styles in the TUI, a long-requested cosmetic improvement.

2. **[#32296 – fix(server): apply plugin pty environment](https://github.com/anomalyco/opencode/pull/32296)**  
   Ensures plugin shell environment variables are applied during PTY creation, with proper precedence ordering.

3. **[#32290 – feat(app): scope sdk/sync hooks per-route](https://github.com/anomalyco/opencode/pull/32290)**  
   Refactoring so `/new-session` targets its own draft server, eliminating global “selected server” confusion.

4. **[#31547 – fix: ensure tool_use/tool_result integrity and Anthropic user-first ordering](https://github.com/anomalyco/opencode/pull/31547)**  
   Defensive fix pairing every `tool_use` with matching `tool_result`. Closes #27594 – a long-standing bug with malformed tool call sequences.

5. **[#32033 – feat(core): hot reload for directory skill sources](https://github.com/anomalyco/opencode/pull/32033)** (closed)  
   Merged PR that enables automatic reload of skills from disk without restarting the process. A top community request.

6. **[#32287 – feat: add reload_skills tool and /reload command](https://github.com/anomalyco/opencode/pull/32287)**  
   Provides a tool the agent can call and a CLI command to rescan skill directories. Complements the hot-reload infrastructure.

7. **[#32284 – fix(core): expand tilde in file tool paths](https://github.com/anomalyco/opencode/pull/32284)**  
   Fixes #32223 – `~` in file paths (read, write, grep, etc.) is now properly expanded, resolving a common annoyance.

8. **[#24289 – fix: Repair truncated JSON tool inputs in LLM session](https://github.com/anomalyco/opencode/pull/24289)**  
   Still open after 50 days. Uses `jsonrepair` to fix malformed JSON from LLM tool calls. Closes three bugs (#24177, #20650, #20786).

9. **[#32272 – refactor(core): derive catalog availability from integrations](https://github.com/anomalyco/opencode/pull/32272)**  
   Removes provider env state; lets Integration own connection methods and track availability. A significant architectural cleanup.

10. **[#32194 – docs: add opencode-sessions-explorer to ecosystem](https://github.com/anomalyco/opencode/pull/32194)**  
    Adds a plugin that exposes the local SQLite session DB as tools for session recall and cost analysis. Growing plugin ecosystem.

---

## Feature Request Trends

- **Dynamic Workflows** (Claude Code-style): **#29059** and **#30308** — project-local YAML workflows for multi-step automation. This is the single most discussed new feature direction.
- **YOLO / Dangerously Skip Permissions: **#8463** remains the top-voted issue overall; users in CI or trusted environments want a fast path.
- **Agent Orchestration & Nested Sub‑Agents: **#32166** proposes deep agent trees (5 levels) with a reference implementation. Parallel requests for better sub‑agent visibility in the Desktop UI (**#32292**).
- **Skill Hot Reload: Already addressed by **#32033** (merged) and **#32287** (open) — users want to edit skills without restarting.
- **New Provider Models: Requests for Z.AI’s GLM-5.2 (**#32172**) and Moonshot’s kimi-k2.7 (**#32280**). Also custom provider support via OmniRoute (**#32170**).**
- **Richer Session API: **#32288** proposes enhanced session status querying to align with SDK types.**
- **Cursor Style Configuration: **#32295** (just opened) addresses a long-standing cosmetic request.**

---

## Developer Pain Points

- **MCP Session Reliability:** Recurring issues with expired/disconnected MCP sessions (#32286, v1.17.5 fix) and unbounded retry loops (#32286) erode trust in tool execution.
- **CJK / Unicode Encoding:** Garbled Chinese text in VSCode copy (**#31311**) and renderer crashes with CJK paths (**#28687**) affect international users heavily.
- **Same Remote Session Collision: **#29871** – multiple clones of the same repo share a session, making independent work impossible.**
- **Windows Terminal Tab Title Reset: **#32293** – v1.17.3 introduced a regression where Ctrl+C resets the tab title to “Windows PowerShell”.**
- **V2 Layout Bugs: **#32285** fixes a broken gate (`isV2` function reference always truthy) that caused missing file tree/terminal buttons in the new layout.**
- **Provider Model Not Found for Sub‑agents: **#31141** – sub-agents fail on tool-using tasks across all categories, impacting multi-agent workflows.**
- **Payment / Billing Stuck: **#32278** – users who pay immediately are blocked by “Insufficient balance” errors, causing support friction.**
- **System‑Reminder Position Instability: **#23595** – the `<system-reminder>` moves in the prompt, invalidating cache for local LLM users.**

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest – 2026-06-14

## Today’s Highlights
A steady stream of fixes landed, with critical patches to cache retention (Anthropic now truly persists 1h) and a fix for the shrinkwrap package duplication that plagued multi-package installations. On the feature side, a PR adding native xAI Grok OAuth login and another for first-run terminal theme detection made it in. Several concerning bugs around hanging tasks with local LLMs and session cancellation remain under investigation.

## Releases
No new releases in the last 24 hours.

## Hot Issues (10 noteworthy)

1. **[#5653 – Move off Shrinkwrap](https://github.com/earendil-works/pi/issues/5653)**
   Installing both `pi-ai` and `pi-coding-agent` places two copies of `pi-ai` on disk, splitting the module-level API provider registry. Community discussion highlights this as a high-priority architectural debt.

2. **[#5703 – 1h cache retention silently degraded to 5m for Claude](https://github.com/earendil-works/pi/issues/5703)**
   Pi never sends the `extended-cache-ttl` beta header, so Anthropic ignores the `1h` TTL, inflating costs. Fixed after multiple user reports.

3. **[#5687 – `pi list` and `pi update` hang with MCP servers](https://github.com/earendil-works/pi/issues/5687)**
   Package subcommands load extensions, which start long-lived MCP servers, causing the CLI to hang after printing output. Workaround: Ctrl‑C.

4. **[#3627 – Expose timeout/retry settings on OpenAI providers](https://github.com/earendil-works/pi/issues/3627)**
   Local inference often exceeds the 10-minute default timeout. Users have been requesting configurable timeouts for months; this issue has 2 👍 and multiple cross-references.

5. **[#5644 – GPT-5.5 context window incorrect](https://github.com/earendil-works/pi/issues/5644)**
   Codex has 400K, API has 1M, but the model definition says otherwise. Reported by the community with official OpenAI link.

6. **[#5671 – `~/.pi` and `cwd/.pi` overlap](https://github.com/earendil-works/pi/issues/5671)**
   Global and project `.pi` folders can collide in `$HOME`. Not critical but a design discussion spurred by #5619.

7. **[#5654 – Add `excludeFromContext` to custom messages](https://github.com/earendil-works/pi/issues/5654)**
   Users of `/status` hooks want to send messages without polluting context, mirroring bash-execution `!!` behaviour. Popular request (👍 1).

8. **[#5706 – Task hangs on “waiting for summary approval” with local LLM](https://github.com/earendil-works/pi/issues/5706)**
   Local OpenAI-compatible backends never proceed past the summary step. Works fine with cloud providers. No workaround yet.

9. **[#5685 – Escape does not stop subagent/background agent](https://github.com/earendil-works/pi/issues/5685)**
   Pressing ESC in terminal shows cancellation but the subagent continues. UX regression for interactive use.

10. **[#5702 – `prompt_cache_retention` sent to providers that reject it](https://github.com/earendil-works/pi/issues/5702)**
    Opencode/Zen return 400 when given unsupported cache headers. The report also raises maintainability concerns in `generate-models.ts`.

## Key PR Progress (8 total)

1. **[#5714 – Add xAI Grok OAuth login](https://github.com/earendil-works/pi/pull/5714)**  
   Built‑in OIDC device‑code provider for xAI Grok, with refresh tokens and proxy endpoint. Appears in `/login`.

2. **[#5711 – Extension prompt guideline API](https://github.com/earendil-works/pi/pull/5711)**  
   New API for extensions to add custom prompt guidelines. Verified working by author.

3. **[#5385 – Detect first-run terminal theme](https://github.com/earendil-works/pi/pull/5385)**  
   Queries OSC for light/dark theme on first launch and persists it, so Pi matches the user’s terminal immediately.

4. **[#5526 – Require terminal events for OpenAI Responses streams](https://github.com/earendil-works/pi/pull/5526)**  
   Stops random stream interruptions by ensuring a terminal response event is received before marking the turn done.

5. **[#5708 – Wrap question extension text instead of truncating](https://github.com/earendil-works/pi/pull/5708)**  
   Fixes #5707: long extension questions now wrap properly in the UI.

6. **[#5701 – Fix minimax-m3 context size](https://github.com/earendil-works/pi/pull/5701)**  
   Corrects advertised context from 1M to 524288 tokens (as enforced by OpenRouter).

7. **[#5704 – Capture system for auto-storing tool results](https://github.com/earendil-works/pi/pull/5704)**  
   Implements the Capture phase of Veil context management: Read, Bash, WebSearch tool results cached with dedup and truncation.

8. **[#5693 – Merging official repo updates](https://github.com/earendil-works/pi/pull/5693)**  
   Routine upstream merge from the mainline repository.

## Feature Request Trends
* **Custom slash commands** (#289) – extending beyond simple LLM communication to UI/logic hooks.
* **Multi-session TUI switching** (#5700) – running multiple agent sessions concurrently and switching between them in the TUI (currently `switchSession` tears down).
* **Native image generation** (#4095) – bringing image generation into Pi’s interactive mode, mirroring Codex TUI.
* **Token throughput (tok/s) in status bar** (#5684) – real-time performance metric for local and remote models.
* **Marketplace categorization & official core extensions** (#5686) – improving discoverability and trust in the extension ecosystem.

## Developer Pain Points
* **Cache handling fragility** – Anthropic cache retention silently dropped to 5m (#5703); cache headers sent to non‑supporting providers (#5702).
* **Package duplication** – Shrinkwrap causes duplicate module instances breaking singleton registries (#5653).
* **Model configuration mismatches** – Incorrect context windows (#5644), missing thinking levels (#5699), semver‑ranged packages not loaded (#5695).
* **TUI quirks** – Model name not refreshing on Ctrl+P (#5696), tab completion grabbing first item after narrowing (#5670), image rendering broken in WezTerm (#5618).
* **Hangs and cancellations** – `pi list`/`pi update` hangs with MCP (#5687), Escape fails to cancel subagents (#5685), local LLM summary approval hang (#5706).
* **Environment and setup issues** – Uppercase header values misinterpreted as env vars (#5661), symlinked `.pi` directory causing system prompt duplication (#5648), global pnpm installs not detected for self‑update (#5689).

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-06-14

## Today’s Highlights
The OAuth free tier is being slashed from 1,000 to 100 daily requests, with plans to shut down the free entry point entirely — sparking a 134-comment debate. Separately, multiple critical bugs surfaced around tool execution after cancellation, duplicate tool-result history, and a TUI freeze caused by zombie processes. On the PR side, key fixes land for memory monitor starvation (avoiding OOM in autonomous loops) and a safe-mode flag for troubleshooting.

## Releases
No new releases in the last 24 hours.

## Hot Issues (10 Noteworthy)

1. **[#3203 – Qwen OAuth Free Tier Policy Adjustment](https://github.com/QwenLM/qwen-code/issues/3203)**  
   *Status: Open, 134 comments*  
   Proposes reducing daily free quota from 1,000 to 100 requests and eventually closing the free entry point. Community reaction is intense — many users rely on the free tier; some call for a paid Pro plan instead.

2. **[#3267 – Requests Limits Overview](https://github.com/QwenLM/qwen-code/issues/3267)**  
   *Closed, 7 comments*  
   User reports hitting the 1,000/day limit without completing any task, suggesting unclear quota tracking. Highlights confusion around free tier constraints.

3. **[#5080 – 阿里云 Standard API Key 与 Token Plan 混用导致 401](https://github.com/QwenLM/qwen-code/issues/5080)**  
   *Open, 5 comments*  
   Mixing Alibaba Cloud API keys with Token Plan providers returns 401. Expectation: different connection methods should work without interfering.

4. **[#5083 – TUI 卡死，疑似僵尸子进程未被回收](https://github.com/QwenLM/qwen-code/issues/5083)**  
   *Open, 5 comments*  
   TUI becomes completely unresponsive due to a zombie bash child process not being reaped. CPU and memory are fine, but UI freezes — a frustrating user experience.

5. **[#5102 – Permission-Contract Probe Side Effect Executed](https://github.com/QwenLM/qwen-code/issues/5102)**  
   *Open, 4 comments*  
   During a permission-contract probe, a provider-requested side effect (writing a file) was executed despite the probe. Potential security concern: contract violations may not be prevented.

6. **[#5099 – Duplicate Tool-Result History for Reused Tool-Call ID](https://github.com/QwenLM/qwen-code/issues/5099)**  
   *Open, 3 comments*  
   When the provider reuses a tool-call ID across turns, Qwen Code sends duplicate results, corrupting provider state. Could trigger schema errors or retry loops.

7. **[#5101 – Repeated Large Tool Results Through Provider History](https://github.com/QwenLM/qwen-code/issues/5101)**  
   *Open, 2 comments*  
   Large tool output repeated on every turn grows context until it becomes too large. A performance and token-waste issue that can lead to OOM.

8. **[#5100 – Agent `name` Param Breaks `/review` Skill](https://github.com/QwenLM/qwen-code/issues/5100)**  
   *Open, 2 comments*  
   The new agent team feature causes the bundled `/review` skill to fail (“no active team”) and then enter a repetitive-call abort loop. Blocks code review workflows.

9. **[#3272 – No Pro Plan Available](https://github.com/QwenLM/qwen-code/issues/3272)**  
   *Closed, 2 comments*  
   User wants to buy a Pro plan but it’s perpetually “sold out”. Indicates unmet demand for paid tiers.

10. **[#4721 – Port Dynamic Workflows from Claude Code](https://github.com/QwenLM/qwen-code/issues/4721)**  
    *Open, 1 comment*  
    Requests a third tier of multi-agent execution (Dynamic Workflows) to complement existing `/swarm`. Seen as a major feature to bridge parity with Claude Code.

## Key PR Progress (10 Important Pull Requests)

1. **[#4598 – Collapsible Thinking Blocks with Duration Timer](https://github.com/QwenLM/qwen-code/pull/4598)**  
   Replaces always-expanded thinking display with collapsible blocks that stream reasoning and collapse on completion, plus duration timer. Improves TUI readability.

2. **[#4943 – `--safe-mode` Flag for Troubleshooting](https://github.com/QwenLM/qwen-code/pull/4943)**  
   Adds CLI flag and env var to disable all customizations (context files, hooks, MCP servers, skills). Provides a clean baseline for debugging issues.

3. **[#5097 – Fix Memory Monitor Starvation During Autonomous Loops](https://github.com/QwenLM/qwen-code/pull/5097)**  
   Core fix prevents memory monitor starvation when event loop is busy. Heartbeat fallback ensures monitors fire, avoiding UI history growth until OOM.

4. **[#5105 – Dedicated Agent Permission Dialog](https://github.com/QwenLM/qwen-code/pull/5105)**  
   Gives sub-agent tool a dedicated “Launch this agent?” permission prompt in both VS Code and web-shell UI, without relying on protocol tool kind.

5. **[#5103 – Default GLM-5.2+ and GLM-6.x to 1M Context](https://github.com/QwenLM/qwen-code/pull/5103)**  
   Updates token-limit rules for newer GLM models to support 1M context window, matching vendor specs.

6. **[#5085 – Internal Kind.Agent for Sub-Agent Tool](https://github.com/QwenLM/qwen-code/pull/5085)**  
   Adds a distinct internal tool kind for agent/sub-agent tool, mapped to `'other'` on the wire. Foundation for better permission flows and UI handling.

7. **[#5098 – Persist Goal Status in Daemon Transcript](https://github.com/QwenLM/qwen-code/pull/5098)**  
   Moves `/goal` state from frontend memory to daemon transcript events, surviving page refresh and multi-device sync.

8. **[#5095 – Import Claude MCP Servers](https://github.com/QwenLM/qwen-code/pull/5095)**  
   Implements `/import-config` Phase 1: migrates Claude Code and Claude Desktop MCP settings into Qwen Code, preserving existing server names.

9. **[#5096 – Make Web-Shell Input Shortcuts Discoverable](https://github.com/QwenLM/qwen-code/pull/5096)**  
   Surface hidden shortcuts (slash commands, `@` file mentions, reverse-i-search) with clickable hints and mouse affordance.

10. **[#5073 – Warn on Oversized Context Instructions](https://github.com/QwenLM/qwen-code/pull/5073)**  
    Adds startup warning when QWEN.md uses >15% of model context window, helping users avoid accidental context overruns.

## Feature Request Trends
- **OAuth / Free Tier Changes**: Requests to lower free quota or shut it down entirely, with strong pushback from the community demanding a real Pro plan.
- **Multi-Agent & Workflows**: Strong interest in porting Claude Code’s Dynamic Workflows (#4721) and improving sub-agent permission handling.
- **Configuration Migration**: Users want one-click import from Claude Code / Desktop (e.g., `/import-config`) to reduce switching friction.
- **UI Refinements**: Collapsible tool results, status line wrapping, model display in footer, and better shortcut discoverability.
- **Performance & Memory**: Automatic history compaction, memory monitor improvement, and warnings for oversized context.
- **Custom Provider Flexibility**: Request to decouple provider identity from SDK protocol to support custom providers with arbitrary IDs.

## Developer Pain Points
- **Free Tier Confusion**: Users hit daily limits unexpectedly and cannot buy a Pro plan (always “sold out”). Unclear quota tracking.
- **TUI Freezes & Zombie Processes**: The interface becomes unresponsive due to unreaped child processes, requiring manual intervention.
- **API Key Mixups**: Mixing standard API keys with token‑plan providers causes 401 errors; configuration is error-prone.
- **Tool Execution After Cancellation**: SIGINT does not stop tool execution reliably, leading to unintended side effects.
- **Duplicate / Large Tool History**: Reused tool-call IDs and repeated large outputs corrupt provider state and waste context.
- **Migration Barriers**: No easy way to import Claude configurations, forcing developers to manually recreate MCP servers and settings.
- **Hidden Shortcuts**: Useful web-shell shortcuts are invisible, making them hard to discover without documentation.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# CodeWhale Community Digest — 2026-06-14

## Today's Highlights

The project formally rebrands from `DeepSeek TUI` to **CodeWhale**, releasing v0.8.60 as the canonical package name and deprecating the legacy `deepseek-tui` npm package. The community's attention is sharply focused on **Agent Fleet architecture** — the project is actively designing a durable, multi-agent control plane with persistent goal mode, role-based worker delegation, and non-blocking fanout. A new "/swarm" UX is being gated until the underlying runtime (durable goal loops, reliable sub-agent checkpointing) ships, signaling a deliberate shift from experimental parallelism toward production-grade agent orchestration.

## Releases

**v0.8.60** — Project rebranding release.
- Canonical name is now **CodeWhale** for the binary, npm package, and all release assets.
- Legacy `deepseek-tui` npm package is deprecated and receives no further updates.
- Migration guide available in `docs/REBRAND.md`.

## Hot Issues

1. **[#3147 — MSBuild FileTracker Fails to Initialize](https://github.com/Hmbown/CodeWhale/issues/3147)** (CLOSED, 7 comments)  
   Windows users cannot run `cmake --build` inside CodeWhale's PowerShell shell due to MSBuild FileTracker initialization failure. High impact for C++ developers on Windows. Community has proposed workarounds but the root cause points to process sandboxing.

2. **[#891 — Support Codex /goal long-running task mode](https://github.com/Hmbown/CodeWhale/issues/891)** (OPEN, 7 comments)  
   Long-standing request to port Codex's persistent goal-driven workflow. Users want agents that keep working toward a defined objective (refactoring, multi-file features) without stopping after a single response. Community discussion emphasizes this as table-stakes for modern agent tools.

3. **[#3096 — Split sub-agents into headless worker runtime](https://github.com/Hmbown/CodeWhale/issues/3096)** (OPEN, 7 comments)  
   Architecture proposal to decouple sub-agents from the TUI process entirely. Current in-process async tasks still feel "UI-shaped." Maintainer is driving toward a lightweight, projectable runtime — a foundational change for v0.8.61.

4. **[#1806 — Sub-agent 120s API timeout renders agent_open nearly unusable](https://github.com/Hmbown/CodeWhale/issues/1806)** (OPEN, 4 comments)  
   Hard 120-second API timeout kills sub-agents mid-task. User reports 5 parallel sub-agents all failing identically during a biobanking SOP conversion task. Community is pressuring for configurable timeouts and checkpointing.

5. **[#2211 — Sub-agent fanout plus hidden worktrees saturate TUI](https://github.com/Hmbown/CodeWhale/issues/2211)** (OPEN, 4 comments)  
   During release work, the TUI hits max-agents sidebar state (5/5) when background shell work and sub-agents compound. Two separate pressure sources — agents and filesystem operations — overwhelm the UI. Maintainer acknowledges this is a design-level issue.

6. **[#3154 — Agent Fleet control plane for always-running verifiable work](https://github.com/Hmbown/CodeWhale/issues/3154)** (OPEN, 3 comments)  
   Epic for the Cursor-style agent-fleet pattern: a manager agent keeps workers moving, notices stuck work, restarts/interrupts, and escalates only what needs human judgment. This is the central design document for v0.8.61's most ambitious feature.

7. **[#3192 — Put CodeWhale up for agentclientprotocol/registry](https://github.com/Hmbown/CodeWhale/issues/3192)** (OPEN, 3 comments)  
   Request to register CodeWhale in the Agent Client Protocol registry, making it installable by Zed and other ACP-compatible tools. Community sees this as critical for ecosystem adoption.

8. **[#3167 — Model Agent Fleet org chart, roles, and delegation policy](https://github.com/Hmbown/CodeWhale/issues/3167)** (OPEN, 3 comments)  
   Speculative design for explicit agent roles: scouts, implementers, reviewers, verifiers, operators. Community debates whether roles should be static or negotiated at runtime.

9. **[#3208 — codewhale-linux-x64 vs codewhale-linux-x64.tar.gz confusion](https://github.com/Hmbown/CodeWhale/issues/3208)** (OPEN, 2 comments)  
   Users are confused by duplicate release assets. The non-archived binary is unclear — raw binary or archive? Maintainer needs to clarify asset naming conventions.

10. **[#3207 — GLIBC_2.39 not found on Ubuntu 22.04](https://github.com/Hmbown/CodeWhale/issues/3207)** (OPEN, 2 comments)  
    Binary compiled against glibc 2.39, which is only available on Ubuntu 24.04+. Users on older LTS releases cannot run the binary. Community requests multi-glibc builds.

## Key PR Progress

1. **[#3220 — Limit mobile event history to prevent freezes](https://github.com/Hmbown/CodeWhale/pull/3220)** (OPEN)  
   Fixes browser freeze in `codewhale serve --mobile` on long code-generation threads. Keeps only the latest 100 visible DOM events and adds `replay_limit` to SSE endpoint. Critical for mobile UX reliability.

2. **[#3141 — Optimize get_thread_detail item fetching (N+1 fix)](https://github.com/Hmbown/CodeWhale/pull/3141)** (CLOSED)  
   Batches item directory reads per `turn_id` instead of scanning per turn. Significant performance improvement for threads with many turns.

3. **[#3140 — Fix command injection in hooks by switching to direct execution](https://github.com/Hmbown/CodeWhale/pull/3140)** (CLOSED)  
   Security fix: hook commands are no longer passed to `sh -c`/`cmd /C`, preventing shell metacharacter injection. Migrates to `build_shell_command` logic with direct process execution.

4. **[#3139 — Parallelize skill syncing](https://github.com/Hmbown/CodeWhale/pull/3139)** (CLOSED)  
   Refactors `sync_registry` to fetch community skills concurrently using `futures_util::future::join_all`. Network-intensive operations no longer block sequentially.

5. **[#2773 — Complete provider fallback chain](https://github.com/Hmbown/CodeWhale/pull/2773)** (CLOSED)  
   Adds automatic provider fallback on retryable errors (429, 5xx, timeout, network). Configurable via `fallback_providers` in TOML config. Major reliability improvement for multi-provider setups.

6. **[#3051 — Add /voice slash command for speech-to-text input](https://github.com/Hmbown/CodeWhale/pull/3051)** (CLOSED)  
   Adds voice input inspired by MiMo Code. Three slash commands for one-shot recording, AI transcription, and seamless composer insertion — reuses existing chat completions API.

7. **[#3172 — Durable fleet inbox and run ledger](https://github.com/Hmbown/CodeWhale/pull/3172)** (CLOSED)  
   Implements append-only JSONL fleet ledger under `workspace/.codewhale/fleet.jsonl` that survives process restart. Records runs, inbox entries, leases, lifecycles, heartbeats, and alerts.

8. **[#3171 — Agent Fleet protocol types and event schema](https://github.com/Hmbown/CodeWhale/pull/3171)** (CLOSED)  
   Defines the durable, serializable contract for the Agent Fleet control plane including `FleetRun`, task/work spec types, status events, and retry policy.

9. **[#2808 — Session save, undo/retry, and snapshot endpoints for GUI](https://github.com/Hmbown/CodeWhale/pull/2808)** (CLOSED)  
   Adds Runtime API endpoints (`POST /v1/sessions`, undo/retry paths) to align GUI capabilities with TUI, reusing existing TUI logic rather than reimplementing.

10. **[#3196 — Ctrl+P / Ctrl+N navigate slash-command autocomplete](https://github.com/Hmbown/CodeWhale/pull/3196)** (CLOSED)  
    Keyboard shortcut ergonomics add-on: Ctrl+P/Ctrl+N as alternatives to arrow keys for navigating inline slash-command popups. Includes guard to yield global Ctrl+P file-picker when slash menu is open.

## Feature Request Trends

Three dominant directions emerge from the current issue stream:

1. **Agent Fleet & Persistent Workflows** — The most active area. Multiple epics and sub-issues design a manager-worker control plane with durable ledgers, role-based tooling, checkpointed sub-tasks, and the ability to survive restarts. Users want CodeWhale to run multi-hour, multi-file tasks without human babysitting.

2. **Goal Mode & Long-Running Loops** — Closely related to Fleet, but specifically focused on the `/goal` pattern: persistent objectives that survive turns, LLM-as-judge for continuation decisions, and progress accounting. The community broadly agrees this is the "killer feature" they want polished.

3. **Sub-Agent Reliability & Non-Blocking Architecture** — A strong undercurrent of frustration with the current sub-agent implementation. Users consistently request: configurable timeouts, checkpointing/resumption, non-blocking parent turns, and clear status visualization. The v0.8.61 roadmap directly addresses this with headless worker runtime and durable fanout.

Secondary trends include: **chat-native workrooms** (threaded, shareable, mobile-accessible), **Claude Code skill ecosystem compatibility**, and **clearer busy/idle state visualization** in the TUI.

## Developer Pain Points

- **MSBuild/CMake Build Failures on Windows** — Issue #3147 highlights that Windows C++ developers are effectively locked out from running builds inside CodeWhale shell. The FileTracker initialization error is a hard blocker for this significant developer segment.

- **Sub-API Timeouts Kill Long-Running Tasks** — Multiple reports (especially #1806) describe sub-agents dying silently from the 120s timeout ceiling. Configurable timeouts and checkpointing are urgently requested; the current behavior makes parallel task offloading unreliable for real work.

- **Context Window Confusion** — Real-world testing with models like GPT-5.5 repeatedly hits `context_length_exceeded` errors (#3204) while the TUI displays incorrect or optimistic context-window metadata. Developers are frustrated by the gap between what the UI reports and what the API enforces.

- **TUI State Opacity** — Issue #2982 captures a common complaint: the TUI doesn't clearly distinguish "thinking" from "done" when text doesn't change. Developers working with slow models or complex tasks need traffic-light-like status indicators.

- **Model/Route Mismatch for Fleet Workers** — Issue #3205 reveals that fleet workers inherit the wrong model or provider configuration (e.g., Z.ai/GLM runs fail when route-effective inventory disagrees with CLI defaults). The architecture currently lacks role-based model routing.

- **Ubuntu LTS Binary Compatibility** — The glibc 2.39 requirement (#3207) excludes users on Ubuntu 22.04 and earlier, a significant portion of the Linux developer community. Binary provisioning strategy needs immediate attention.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*