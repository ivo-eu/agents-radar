# AI CLI Tools Community Digest 2026-06-18

> Generated: 2026-06-18 03:43 UTC | Tools covered: 9

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

# Cross-Tool Comparison Report — 2026-06-18

## 1. Ecosystem Overview

The AI CLI tool ecosystem remains in a phase of rapid maturation, with six of seven tools showing active development cycles and meaningful community engagement this week. The dominant themes across the landscape are **pricing transparency and regional equity** (Claude Code, Codex), **agent reliability and control** (Gemini CLI, Copilot CLI, CodeWhale), and **permission model granularity** (Copilot CLI, OpenCode). While the major players (Claude Code, Codex, Gemini, Copilot) continue to define the baseline feature set, smaller tools like DeepSeek TUI/CodeWhale and Pi are innovating rapidly on specific niches—particularly static builds, XDG compliance, and provider extensibility. **Platform portability** (ARM64, Wayland, SSH environments) and **enterprise authentication** remain consistent pain points across the board.

---

## 2. Activity Comparison

| Tool | Hot Issues (last 24h) | PRs Updated (last 24h) | Release Status | Notable Volume |
|---|---|---|---|---|
| **Claude Code** | 10 active | 7 | ✅ v2.1.181 released | 1,475 comments on #16157 (Max plan bug) |
| **OpenAI Codex** | 10 active | 10 | ✅ 3 nightly Rust alphas | 49 comments on #25749 (phone lockout) |
| **Gemini CLI** | 10 active | 10 | ✅ v0.47.0 + v0.48.0-preview.0 | P0 nightly release failure |
| **GitHub Copilot CLI** | 10 active | 0 | ❌ None | Post-outage fallout dominates |
| **Kimi Code CLI** | 2 active | 0 | ❌ None | Lightest activity day |
| **OpenCode** | 10 active | 10 | ✅ v1.17.8 released | 55 👍 on TPS request #6096 |
| **Pi** | 10 active | 10 | ❌ None | High engagement on scroll bug #5825 |
| **Qwen Code** | 6 active | 10 | ✅ v0.18.3 stable + nightly | Multiple release tags today |
| **CodeWhale (DeepSeek TUI)** | 10 active | 10 | ❌ None | Intense bug-fix cycle pre-v0.9.0 |

- **Most releases:** Qwen Code (3 versions), Claude Code (1), Gemini CLI (2), OpenCode (1), Codex (3 nightlies)
- **Highest community engagement:** Claude Code (1,400+ comments on single issue), Codex (49 comments on lockout issue)
- **Least active:** Kimi Code CLI (only 2 issues, no PRs)

---

## 3. Shared Feature Directions

Several feature requirements appear across two or more tool communities, suggesting broad unmet needs:

### 3.1 Granular Permission & Approval Controls
- **Copilot CLI** (#1973): Tool whitelist for interactive mode (read-only tools without approval)
- **OpenCode** (#7928): Runtime permission mode toggle (auto-edit ↔ permission-request)
- **CodeWhale** (#3279, #3295): Plan/Agent mode toggle consistency + runtime permission enforcement
- **Claude Code** (#69241): Auto-accept edits in JetBrains (skip diff dialog)

**Signal:** Users want *graduated* autonomy—not all-or-nothing—and the ability to switch modes at runtime without restarting sessions.

### 3.2 Session Continuity & Non-Blocking Interaction
- **Claude Code** (#50246): Message queue mode to enqueue prompts without interrupting active tasks
- **CodeWhale** (#1530): `exec --resume` / `--session-id` for multi-turn CLI workflows
- **Copilot CLI** (#3844, #3837): Custom aliases, show working directory on resume
- **Codex** (#28835, #28824): Current-time reminders for latency-aware tooling

**Signal:** The ambient, always-on assistant model creates friction when users need to queue work or recover from interruptions. Session persistence and non-blocking interfaces are becoming table stakes.

### 3.3 Pricing Transparency & Regional Equity
- **Claude Code** (#17432): India-specific INR pricing (444 👍)
- **Codex** (#25749): Permanent lockout due to inaccessible phone verification
- **Copilot CLI** (#3355): Configurable context window (200K vs 1M token cap)

**Signal:** Users are increasingly aware of cost asymmetries and demand billing parity, region-specific pricing, and visibility into how their tokens are consumed.

### 3.4 Provider Extensibility & BYOK
- **Pi** (#5849): Azure AI Foundry provider for Claude
- **Copilot CLI** (#3839): Ollama Cloud BYOK compatibility (400 errors)
- **OpenCode** (#32172): GLM-5.2 model support
- **CodeWhale** (#1481): OpenCode Go/Zen (DeepSeek V4) provider

**Signal:** Users expect CLI tools to be universal frontends, not locked to a single model provider. MCP and plugin ecosystems are accelerating this demand.

### 3.5 Platform-Specific Fixes (ARM64, Wayland, SSH, Windows)
- **Claude Code** (#39636): Cowork VM fails on Snapdragon X Plus (ARM64 Windows)
- **Gemini CLI** (#21983): Browser subagent fails on Wayland
- **Pi** (#5841): Warp terminal detection for Kitty image protocol
- **OpenCode** (#21277): ANSI escape codes after crash on PowerShell
- **Qwen Code** (#5281): TUI unresponsive on SSH with PolKit auth hijacking

**Signal:** The "works on my machine" era is ending. Users on non-macOS/Linux-standard environments demand first-class support, and these gaps are becoming deal-breakers.

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | Copilot CLI | OpenCode | Pi | Qwen Code | CodeWhale |
|---|---|---|---|---|---|---|---|---|
| **Primary focus** | Agentic coding assistant | Universal AI CLI | Agent orchestration | GitHub-native assistant | Developer workflow tool | Extensible universal frontend | Multilingual coding agent | Lightweight TUI agent |
| **Target user** | Professional developers | Enterprise/power users | Google ecosystem devs | GitHub ecosystem devs | Cross-platform devs | Terminal power users | Chinese & global devs | Rust/performance-focused |
| **Plugin ecosystem** | MCP-based, skill files | Plugin agent roles (TOML) | Skill/agent system | Pre/post tool hooks | Plugin trigger hooks | Provider abstraction | Channel adapters (QQ, etc.) | Config-based skills |
| **Language** | TypeScript/Node | Rust (CLI) + TS (app) | TypeScript/Node | TypeScript/Node | Rust | TypeScript/Node | TypeScript/Node | Rust |
| **Unique strength** | `/remote-control` mode, Cowork VM | Computer Use, unified-exec | Jupyter/Colab integration | GitHub enterprise auth | V2 fuzzy edit matching | Provider diversity, TUX | Multi-channel (QQ/Telegram) | Static musl builds, config preservation |
| **Key pain point** | Max plan billing bugs | Phone verification lockout | Agent hangs/false success | Post-outage instability | Windows terminal corruption | Shrinkwrap dependency issues | Chinese i18n completeness | Renaming migration gaps |

**Key differentiation:**
- **Claude Code** leads on remote/mobile workflows but struggles with billing trust.
- **Codex** is building the richest plugin/agent role system (TOML manifests, MCP namespace exposure) but has severe authentication friction.
- **Gemini CLI** excels at agent orchestration (subagents, browser agents) but reliability issues erode confidence.
- **Copilot CLI** owns the GitHub enterprise integration but lags in plugin flexibility.
- **OpenCode** stands out for edit reliability (fuzzy matching, V2 tooling) and terminal UX polish.
- **Pi** differentiates on provider diversity (Azure Foundry, Anthropic OAuth, local LLMs) and rapid bug-fixing.
- **Qwen Code** is the most multi-platform (QQ, Telegram, WeChat channels) and release-rapid (3 releases today).
- **CodeWhale** (DeepSeek TUI) innovates on Rust-native performance, config preservation, and static builds, but the rename caused fragmentation.

---

## 5. Community Momentum & Maturity

### High Maturity (Stable, enterprise adoption)
- **Claude Code**: Largest community (1,400+ comments on single issue), highest feature request volume, established release cadence. The Max plan bug (#16157) is a **revenue-critical issue** that Anthropic must resolve to maintain enterprise trust.
- **OpenAI Codex**: Strong PR velocity (10/day), active plugin ecosystem development. The authentication lockout issue (#25749) is a **trust-breaker** for security-conscious users.
- **GitHub Copilot CLI**: Post-outage instability is the dominant narrative, but the tool benefits from GitHub's enterprise distribution. Slower PR cadence (0 today) suggests a maintenance phase.

### Rapid Iteration (High release frequency, fast feature development)
- **Qwen Code**: Three releases today (v0.18.2 → v0.18.3), the fastest cadence in the ecosystem. Pushes i18n, OOM fixes, and circuit breakers simultaneously. Strong momentum.
- **OpenCode**: v1.17.8 shipping yesterday with active TUI/performance work. Fuzzy edit matching port signals a focus on **core tooling quality**.
- **Pi**: No release today but 10 PRs updated. Azure Foundry provider, Anthropic cache pricing fix, and TUI streaming fixes show rapid iteration.

### Niche but Growing
- **CodeWhale (DeepSeek TUI)**: Intense bug-fix cycle (10 PRs, 26 updated) with ambitious v0.9.0 Workroom feature. The rename from deepseek-tui caused short-term confusion but the community remains engaged.
- **Kimi Code CLI**: Least active today (2 issues, 0 PRs). The SSL certificate ignore request (#2458) suggests a small but committed enterprise user base.

---

## 6. Trend Signals

### 6.1 Agent Autonomy Is Hitting Trust Boundaries
Multiple communities report agents that **over-execute, fail silently, or refuse to respect user intent**: Gemini CLI's false success reports (#22323), CodeWhale's self-looping (#3275), Copilot CLI's silent command rewrites (#2643). Users want **verifiable, interruptible, and permission-enforced** agent execution. Expect **circuit breakers, execution budgets, and runtime permission toggles** to become standard features within 6 months.

### 6.2 "MCP Everywhere" Is Becoming Reality
Claude Code (MCP tools), Codex (MCP namespace exposure #28825), Copilot CLI (MCP server plugins), and OpenCode (MCP schema fixes) are all standardizing on Model Context Protocol. **Tools that don't support MCP will lose developer mindshare** in 2027.

### 6.3 Context Management Is the Next Frontier
Claude Opus 4.6's 1M token window (#3355), Pi's percentage-based compaction (#5848), and Codex's context-exhaustion errors (#8190) point to a shared challenge: **tooling hasn't caught up with model context sizes**. Expect smart compaction, token budgeting, and progressive summarization to become key differentiators.

### 6.4 Cost Visibility Is a Competitive Advantage
Claude Code's Max plan bug (#16157), Codex's write amplification (#28224), and OpenCode's TPS request (#6096) all indicate that **users want to know where their money/tokens go**. Tools that surface per-session cost, token budgets, and real-time usage metrics will win enterprise adoption over those that bury billing in settings.

### 6.5 Enterprise Authentication Is Fragmented
Codex's FIDO2/SMS OTP conflict (#25737), Copilot CLI's repeated login prompts (#254), and Qwen Code's PolKit hijacking (#5281) show that **OAuth, passkeys, and enterprise SSO are not yet seamless** in CLI tools. Users expect browser-grade authentication flows in the terminal.

### 6.6 Static Binaries & Portable Distribution Are Gaining Traction
CodeWhale's musl static build (#3274) and Pi's XDG compliance (#534) signal a trend toward **self-contained, platform-conformant binaries**. As these tools are used in CI/CD pipelines, Docker containers, and locked-down enterprise environments, dynamic dependency chains become unacceptable.

---

**Bottom line for decision-makers:** Claude Code remains the feature leader but faces a trust crisis around billing. Codex is building the most extensible plugin architecture but must solve authentication. Gemini CLI and Copilot CLI offer strong ecosystem integration but need reliability improvements. For teams prioritizing provider flexibility, Pi and OpenCode are solid bets. Qwen Code and CodeWhale are fast-moving alternatives worth watching—especially for Rust-focused or Chinese-language workflows.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
**Data as of 2026-06-18** | Source: github.com/anthropics/skills

---

## 1. Top Skills Ranking

Below are the most-discussed Skill submissions (PRs), ranked by community engagement.

**#1 — Document Typography Skill** ([PR #514](https://github.com/anthropics/skills/pull/514))
*Functionality:* Prevents orphan word wrap, widow paragraphs, and numbering misalignment in AI-generated documents. Targets a pervasive quality gap in Claude's document output.
*Discussion highlights:* The community broadly agrees typographic issues affect essentially every document Claude generates. No substantive objections — the discussion centers on implementation scope.
*Status:* **Open** (created 2026-03-04)

**#2 — ODT Skill** ([PR #486](https://github.com/anthropics/skills/pull/486))
*Functionality:* Create, fill, read, and convert OpenDocument Format files (.odt, .ods). Bridges Claude into the LibreOffice/open-source document ecosystem.
*Discussion highlights:* The author explicitly names the trigger set ("ODT", "ODS", "ODF", "OpenDocument"), indicating attention to precise invocation design. Minimal controversy.
*Status:* **Open** (created 2026-03-01, updated 2026-04-14)

**#3 — Frontend-Design Skill Improvement** ([PR #210](https://github.com/anthropics/skills/pull/210))
*Functionality:* Rewrites the existing frontend-design skill for clarity, actionability, and internal coherence. Every instruction must be executable within a single conversation.
*Discussion highlights:* Represents a "skill improvement" category rather than net-new skill. Community values precision and executability over breadth.
*Status:* **Open** (created 2026-01-05)

**#4 — Skill Quality Analyzer + Skill Security Analyzer** ([PR #83](https://github.com/anthropics/skills/pull/83))
*Functionality:* Two meta-skills that evaluate other Skills across structure/documentation (20%), security, and quality dimensions. A self-audit layer for the ecosystem.
*Discussion highlights:* Opens the question of governance — who audits the auditors? Significant interest in quality tooling.
*Status:* **Open** (created 2025-11-06)

**#5 — SAP-RPT-1-OSS Predictor** ([PR #181](https://github.com/anthropics/skills/pull/181))
*Functionality:* Wrapper for SAP's open-source tabular foundation model for predictive analytics on SAP business data (released at SAP TechEd 2025).
*Discussion highlights:* Enterprise/ERP-adjacent skill. Taps demand for Claude interacting with specialized ML models rather than generating content.
*Status:* **Open** (created 2025-12-28)

**#6 — Skill-Creator Evaluation Fixes** (multiple PRs: [#1298](https://github.com/anthropics/skills/pull/1298), [#1099](https://github.com/anthropics/skills/pull/1099), [#1050](https://github.com/anthropics/skills/pull/1050))
*Functionality:* Addresses the `run_eval.py` recall=0% bug — the evaluation harness reports 0% trigger rate for all skill descriptions, breaking the description-optimization loop. Includes Windows subprocess fixes (PATHEXT, cp1252 encoding, select-on-pipes).
*Discussion highlights:* **Highest-urgency infrastructure fix.** Multiple independent reproductions. These PRs collectively represent the community's biggest debugging effort.
*Status:* **Open** (PR #1298 most recent, 2026-06-10)

**#7 — ServiceNow Platform Skill** ([PR #568](https://github.com/anthropics/skills/pull/568))
*Functionality:* Broad ServiceNow assistant covering ITSM, ITOM, ITAM/SAM Pro, FSM, HRSD/CSM, SPM/PPM, Vulnerability Response, SecOps, CSDM, and IntegrationHub.
*Discussion highlights:* Extremely ambitious scope. Enterprise-adjacent. The skill aims to be a "platform assistant" rather than a narrow scripting helper.
*Status:* **Open** (created 2026-03-08, updated 2026-04-23)

**#8 — Testing Patterns Skill** ([PR #723](https://github.com/anthropics/skills/pull/723))
*Functionality:* Covers testing philosophy (Testing Trophy model), unit testing (AAA pattern, naming, edge cases), React component testing (Testing Library), E2E (Playwright), and accessibility testing.
*Discussion highlights:* Broad community interest in quality assurance automation. Fills a gap in the skills collection.
*Status:* **Open** (created 2026-03-22)

---

## 2. Community Demand Trends

From Issues (top 13, sorted by discussion activity), the community's most-anticipated needs cluster in four directions:

**Organizational skill sharing** — Issue [#228](https://github.com/anthropics/skills/issues/228) (14 comments, 7 👍) requests org-wide skill sharing without manual .skill file exchange. Currently teams must download, Slack, and upload — users want a shared library or direct sharing link.

**Evaluation harness reliability** — Issues [#556](https://github.com/anthropics/skills/issues/556) (12 comments, 7 👍) and [#1169](https://github.com/anthropics/skills/issues/1169) document that `run_eval.py` scores 0% recall on every query, making the description-optimization loop useless. This is the community's most immediate blocker.

**Windows compatibility** — Issue [#1061](https://github.com/anthropics/skills/issues/1061) lists three Unix-first assumptions blocking native Windows Python 3.14 usage. Multiple PRs (#1050, #1099) address this, indicating sustained cross-platform demand.

**Security & governance** — Issue [#492](https://github.com/anthropics/skills/issues/492) warns that community skills under the `anthropic/` namespace create trust boundary abuse vulnerabilities. Issue [#412](https://github.com/anthropics/skills/issues/412) proposes an agent-governance skill for safety patterns (policy enforcement, threat detection, audit trails). Security tooling is the emerging governance demand.

**Duplicate skill management** — Issue [#189](https://github.com/anthropics/skills/issues/189) (6 comments, 9 👍) reports that `document-skills` and `example-skills` plugin installations contain identical content. Users want deduplication logic in the installation pipeline.

---

## 3. High-Potential Pending Skills

Active PRs not yet merged with strong community discussion momentum:

| Skill | PR | Last Updated | Why Watch |
|-------|-----|--------------|-----------|
| Document Typography Control | [#514](https://github.com/anthropics/skills/pull/514) | 2026-03-13 | Addresses universal pain point; likely to merge soon |
| ODT / OpenDocument Format | [#486](https://github.com/anthropics/skills/pull/486) | 2026-04-14 | Broad ISO-standard document format; fills LibreOffice gap |
| PDF case-sensitive fix | [#538](https://github.com/anthropics/skills/pull/538) | 2026-04-29 | Simple fix; low-risk merge candidate |
| YAML unquoted description validation | [#539](https://github.com/anthropics/skills/pull/539) | 2026-04-16 | Prevents silent parsing failures; quality-of-life improvement |
| DOCX tracked-change w:id collision | [#541](https://github.com/anthropics/skills/pull/541) | 2026-04-16 | Fixes document corruption with existing bookmarks |
| ServiceNow platform skill | [#568](https://github.com/anthropics/skills/pull/568) | 2026-04-23 | Enterprise demand; broad scope may slow approval |
| Testing patterns skill | [#723](https://github.com/anthropics/skills/pull/723) | 2026-04-21 | High community utility; fills QA gap |
| AURELION skill suite | [#444](https://github.com/anthropics/skills/pull/444) | 2026-05-06 | Structured cognitive framework; novel paradigm |

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is for three intertwined capabilities: producing high-quality formatted documents (typography, ODT, PDF, DOCX), building evaluation infrastructure that actually works (run_eval recall, Windows support, YAML validation), and securing the skill-supply-chain against trust-boundary abuse (namespace impersonation, governance patterns).**

---

# Claude Code Community Digest — 2026-06-18

## Today’s Highlights
A new patch release (v2.1.181) lands with two quality-of-life improvements: inline `/config` syntax for settings and an Apple Events sandbox opt-in. The community continues to voice strong demand for region‑specific pricing (India) and a message‑queue mode to avoid interrupting active tasks. A long‑running usage‑limits bug on Max subscriptions (#16157) still dominates discussion with over 1,400 comments.

## Releases
**v2.1.181** (latest)
- Added `/config key=value` syntax to change any setting from the prompt (works in interactive, `-p`, and Remote Control modes).
- Added `sandbox.allowAppleEvents` opt‑in to let sandboxed commands send Apple Events on macOS.
- Included `CLAUDE_CLIENT_P…` (truncated in source) – likely a new environment variable for client configuration.

---

## Hot Issues (10 noteworthy)

1. **[#16157] Instantly hitting usage limits with Max subscription**  
   *Open, 1,475 comments, 691 👍*  
   The highest‑engagement bug; users on the Max plan report hitting token caps within minutes of starting a task. The volume of reactions and comments suggests this is a critical revenue/trust issue for Anthropic.  
   https://github.com/anthropics/claude-code/issues/16157

2. **[#17432] Feature Request: India‑Specific Pricing Plans (INR)**  
   *Open, 198 comments, 444 👍*  
   High demand for localised pricing comparable to ChatGPT/Gemini. The sheer number of upvotes makes it the top feature request by community support.  
   https://github.com/anthropics/claude-code/issues/17432

3. **[#34255] Remote Control: automatic reconnection doesn’t work**  
   *Open, 50 comments, 90 👍*  
   Connections drop silently with no recovery, affecting macOS/iOS users. Persists for months – a reliability pain point for mobile workflows.  
   https://github.com/anthropics/claude-code/issues/34255

4. **[#50246] Feature Request: Message queue mode**  
   *Open, 32 comments, 99 👍*  
   Users want to enqueue follow‑up prompts instead of interrupting an active task. Popularity indicates a need for non‑blocking interaction.  
   https://github.com/anthropics/claude-code/issues/50246

5. **[#39636] Cowork VM guest kernel never boots on Snapdragon X Plus (ARM64)**  
   *Open, 29 comments*  
   ARM64 Windows users of Cowork are stuck with connection timeouts. A niche but impactful platform gap.  
   https://github.com/anthropics/claude-code/issues/39636

6. **[#25128] Drag‑and‑drop not working in VS Code extension chat panel**  
   *Open, 20 comments, 40 👍*  
   Regression introduced in v2.1.6, still unresolved. Breaks a basic UX expectation in the IDE integration.  
   https://github.com/anthropics/claude-code/issues/25128

7. **[#63870] Bash tool calls emitted as raw `<invoke>` text instead of executing**  
   *Open, 17 comments, 20 👍*  
   Malformed tool invocations; similar reports (#61122, etc.) suggest a recurring model‑side bug.  
   https://github.com/anthropics/claude-code/issues/63870

8. **[#28379] Slash commands not supported in `/remote-control` UI**  
   *Open, 10 comments, 48 👍*  
   Parity gap: `/clear`, `/compact`, etc. sent as plain text when typed from the web or mobile remote interface.  
   https://github.com/anthropics/claude-code/issues/28379

9. **[#69241] [FEATURE] Add plugin setting to auto‑accept edits in JetBrains IDE**  
   *Opened today, 4 comments*  
   A small but vocal request to skip the diff dialog in JetBrains. Highlights a desire for faster, less obstructive editing.  
   https://github.com/anthropics/claude-code/issues/69241

10. **[#69227] VS Code extension pollutes process environment with `NoDefaultCurrentDirectoryInExePath=1`**  
    *Opened yesterday, 3 comments*  
    A side‑effect that breaks other extensions. Shows the importance of environment isolation in IDE plugins.  
    https://github.com/anthropics/claude-code/issues/69227

---

## Key PR Progress (7 items; all updated in last 24h)

1. **[#41611] Add the missing source to claude code** (Open)  
   Minor fix adding a missing source identifier.  
   https://github.com/anthropics/claude-code/pull/41611

2. **[#41447] feat: open source claude code ✨** (Open)  
   A community PR claiming to close multiple open‑source‑related issues (#59, #456, #2846, #22002, #41434). Unclear if official or satire, but it sparks discussion.  
   https://github.com/anthropics/claude-code/pull/41447

3. **[#69226] Update frontend‑design skill** (Merged)  
   Bumps the frontend‑design plugin to v1.1.0 with improvements.  
   https://github.com/anthropics/claude-code/pull/69226

4. **[#19867] fix(code‑review): allow re‑reviews when new commits are pushed** (Open)  
   Adds smarter skip detection and a `--force` flag for the code‑review plugin. Addresses a long‑standing workflow gap.  
   https://github.com/anthropics/claude-code/pull/19867

5. **[#33443] fix: Update Dockerfile to use native installer** (Open)  
   Switches `.devcontainer/Dockerfile` from deprecated npm to the native installer (Node 24.14).  
   https://github.com/anthropics/claude-code/pull/33443

6. **[#60427] docs: use standard GitHub capitalization in README** (Merged)  
   Tiny documentation polish.  
   https://github.com/anthropics/claude-code/pull/60427

7. **[#60732] docs: polish plugins README wording** (Merged)  
   Minor wording improvement for the plugin ecosystem description.  
   https://github.com/anthropics/claude-code/pull/60732

---

## Feature Request Trends
The most requested directions from open issues include:

- **Localised pricing** (India INR plans) – #17432 is the loudest signal, suggesting Anthropic may need to expand billing regions.
- **Non‑blocking interaction** – Message‑queue mode (#50246) and queue commands (#68998) are repeatedly requested to avoid interrupting active tasks.
- **Remote control parity** – Slash commands (#28379) and per‑message copy buttons (#69254) are missing from the web interface.
- **IDE integration depth** – Auto‑accept edits in JetBrains (#69241) and drag‑and‑drop in VS Code (#25128) show users want fewer friction points inside their editors.
- **Design system integration** – Querying Claude Design from Claude Code (#60327) and native resolution of `claude.ai/design` links (#69239) point to a desire for tighter cross‑product workflows.

---

## Developer Pain Points
Recurring themes among bugs and high‑frequency requests:

- **Cost and usage surprises** – The Max‑plan limit bug (#16153) and excessive token consumption reports (#69253) indicate that billing transparency and predictable usage are major frustrations.
- **Remote Control reliability** – Silent connection drops (#34255) and missing reconnection break trust in mobile/web workflows.
- **Platform gaps** – Cowork fails on ARM64 Windows (#39636); image paste broken on Windows long sessions (#69234); CPU busy‑spin on macOS idle (#68931). Developers on non‑Intel Macs or Snapdragon Windows feel second‑class.
- **Model‑side tool invocation bugs** – Raw `<invoke>` text instead of execution (#63870) and nested subagent routing failures (#69212, #69249) degrade agentic workflows.
- **Environment side‑effects** – VS Code extension polluting other extensions’ environment (#69227) and MCP OAuth issues on SSH (#69205) show that integration hygiene needs attention.
- **UI/UX regressions** – Flickering in resumed sessions (#69255) and hidden scrollbar in dark themes (#69250) are minor but frequent annoyances that erode overall polish.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-06-18

## Today's Highlights

A cluster of critical authentication and database integrity issues is dominating the community's attention, with multiple reports of unrecoverable lockouts due to legacy phone verification and corrupted local SQLite state after updates. On the development side, a significant wave of pull requests lands plugin agent roles, MCP namespace exposure, and a new time-reminder infrastructure, signaling a push toward richer plugin ecosystems and latency-aware tooling. Three Rust CLI alpha releases (v0.141.0-alpha.5 through .7) were published in the last 24 hours but lack visible changelogs.

---

## Releases

Three nightly Rust CLI releases were published without detailed changelogs:
- **rust-v0.141.0-alpha.5**, **rust-v0.141.0-alpha.6**, **rust-v0.141.0-alpha.7** — no changelog content provided. These appear to be incremental nightly builds.

---

## Hot Issues

1. **[#25749 — Codex requires verification of an inaccessible legacy phone number and provides no replacement or recovery path](https://github.com/openai/codex/issues/25749)**  
   *49 comments, 30 👍*  
   Users authenticated via Google OAuth and protected by MFA are blocked from Codex when the system demands SMS OTP to a phone number they no longer have access to. No account recovery or phone-change path exists within Codex, creating a permanent lockout. This is the single most commented issue today.

2. **[#17827 — Customizable status line](https://github.com/openai/codex/issues/17827)**  
   *16 comments, 71 👍*  
   A long-running feature request (since April) for a TUI status bar showing token usage, model, rate limits, and git branch — inspired by Claude Code. The high upvote count signals strong demand for real-time telemetry in the terminal UI.

3. **[#25178 — Windows Computer Use screenshot fails on Windows 10 22H2: `SetIsBorderRequired` error](https://github.com/openai/codex/issues/25178)**  
   *12 comments, 4 👍*  
   Computer Use on Windows can list windows and send input but fails on any screenshot-capable call due to a missing interface error (`0x80004002`). Affects a core Computer Use workflow on a still-common OS version.

4. **[#21211 — Thread navigation/loading slows from unbounded metadata and eager large-history hydration](https://github.com/openai/codex/issues/21211)**  
   *12 comments, 2 👍*  
   A performance regression where thread metadata bloat and eager hydration cause severe slowdowns in navigation and loading. Filed by a well-known .NET community member.

5. **[#24006 — Codex cannot access its local database after update on macOS](https://github.com/openai/codex/issues/24006)**  
   *11 comments, 9 👍*  
   The app refuses to launch post-update with `database disk image is malformed`. No recovery path is offered. High community impact given the number of affected users.

6. **[#25737 — Codex CLI login forces SMS OTP step-up on security-key-only (passkey) accounts](https://github.com/openai/codex/issues/25737)**  
   *11 comments, 6 👍*  
   Even after successful FIDO2/WebAuthn authentication, the CLI redirects to a phone-OTP channel. The browser flow respects the Advanced Account Security setting; the CLI does not.

7. **[#28422 — Image generation regression: valid image not saved when status remains `generating`](https://github.com/openai/codex/issues/28422)**  
   *10 comments, 2 👍*  
   A 0.140.0 regression where successfully generated images are discarded if the status field never flips to `completed`. Users lose output without warning.

8. **[#8190 — "Codex ran out of room in the model's context window" error](https://github.com/openai/codex/issues/8190)**  
   *8 comments, 2 👍*  
   A recurring context-exhaustion error during remote compact tasks with no recovery guidance. First filed in December 2025, still open — a long-standing pain point.

9. **[#24182 — Feature request: show usage information persistently in Codex App](https://github.com/openai/codex/issues/24182)**  
   *6 comments, 2 👍*  
   Users want in-view rate-limit and usage monitoring instead of burying it in Settings. Mirrors the TUI status-line request for the GUI app.

10. **[#28224 — Codex SQLite feedback logs write ~640 TB/year, rapidly consuming SSD endurance](https://github.com/openai/codex/issues/28224)**  
    *6 comments, 1 👍*  
    A diagnostics-heavy write amplification bug: excessive logging to `logs_2.sqlite` can theoretically consume 640 TB of writes per year, which on consumer SSDs means hardware wear-out in months. Filed yesterday and already drawing attention.

---

## Key PR Progress

1. **[#28825 — Expose selected MCP namespaces as direct model tools](https://github.com/openai/codex/pull/28825)**  
   Adds `features.code_mode.direct_only_tool_namespaces` so tools like history and notes remain available during code mode without being exposed to `exec`. Critical for MCP deferral correctness.

2. **[#28845 — Support plugin agent roles](https://github.com/openai/codex/pull/28845)**  
   Introduces TOML-based agent role manifests for plugins, enabling namespaced `spawn_agent` calls like `sample:researcher`. A foundational piece for the plugin ecosystem.

3. **[#28824 — Current time reminders impl for system clock (varlatency 2/n)](https://github.com/openai/codex/pull/28824)**  
   Injects current-time data into history before model requests, with session cadence tracking. No client-side clock integration yet — see next PR.

4. **[#28844 — Reuse parsed plugin skills during session startup](https://github.com/openai/codex/pull/28844)**  
   Caches parsed plugin skill-root snapshots to avoid reparsing on every session start. A performance optimization for multi-plugin users.

5. **[#28780 — unified-exec: retain PathUri in command events](https://github.com/openai/codex/pull/28780)**  
   Ensures app-server can report foreign-platform paths without breaking existing client path-string formats. Enables cross-platform remote execution reporting.

6. **[#28835 — Add app-server current-time implementation (varlatency 3/n)](https://github.com/openai/codex/pull/28835)**  
   Defines a `currentTime/read` JSON-RPC method so the server can request the client's current time. Paired with #28824.

7. **[#28843 — Persist fsmonitor status refreshes](https://github.com/openai/codex/pull/28843)**  
   Fixes a subtle git performance bug where `GIT_OPTIONAL_LOCKS=0` prevents fsmonitor token refresh, causing unnecessary full worktree scans on every background status.

8. **[#28836 — Support assistant realtime append text](https://github.com/openai/codex/pull/28836)**  
   Enables replay of assistant text in realtime voice continuity across sessions. Unblocks frontend voice-handoff UX.

9. **[#28790 — Support plugin manifest path lists](https://github.com/openai/codex/pull/28790)**  
   Allows `plugin.json` `skills` field to accept an array of paths instead of a single string. Enables multi-directory skill packages.

10. **[#28838 — Support Codex home instructions directory](https://github.com/openai/codex/pull/28838)**  
    Loads all non-empty `.md` files from `~/.codex/instructions/` as global instructions, sorted deterministically. Extends the existing `AGENTS.md` pattern with a richer directory-based system.

---

## Feature Request Trends

1. **Persistent usage/rate-limit display** — Multiple requests (#17827, #24182) ask for real-time token/rate-limit info visible without leaving the workspace, both in the TUI status line and the GUI app toolbar.

2. **Customizable quota refresh anchor** — Issue #28839 requests that users control when their quota resets rather than inheriting a fixed rolling schedule. Indicates power users hitting artificial throttling due to misaligned refresh windows.

3. **User-visible telemetry at rest** — The "status line" request (#17827) and "persistent usage" request (#24182) both reflect a broader desire for ambient awareness of system state without disruptive navigation.

---

## Developer Pain Points

1. **Authentication friction with security keys** — Issues #25749 and #25737 both describe scenarios where FIDO2/passkey authentication is complete but Codex still demands SMS OTP, with no recovery path for lost phone numbers. This creates permanent lockouts for security-conscious users.

2. **Local database instability after updates** — Multiple macOS and Windows reports (#24006, #24030, #28666, #28841) describe the app becoming completely non-functional after an update due to corrupted or locked SQLite databases, with no built-in recovery.

3. **SSD endurance risk from logging** — Issue #28224 reveals extreme write amplification in the feedback logging system. This is a silent hardware-wear issue that most users won't detect until their SSD fails.

4. **Update/launch failures on non-English Windows** — Issues #28262 and #28842 show that Windows paths containing non-ASCII characters or specific system driver states (Null driver) prevent the app from launching at all, suggesting insufficient i18n and edge-case testing in the Windows packaging.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-06-18

## Today’s Highlights
Two minor releases landed today: **v0.47.0** (stable) and **v0.48.0-preview.0** (preview). A **nightly release failure (P0)** triggered immediate attention. The community saw critical bug fixes for Jupyter Notebook/JSON corruption and shell command hangs, alongside several security-focused CI hardening PRs. The most active discussion threads continue to revolve around agent reliability and subagent false successes.

## Releases

- **[v0.47.0](https://github.com/google-gemini/gemini-cli/releases/tag/v0.47.0)** — Changelog preparation and a fix to respect backend definitions.  
- **[v0.48.0-preview.0](https://github.com/google-gemini/gemini-cli/releases/tag/v0.48.0-preview.0)** — Dependency cooldown for npm packages and internal refactoring.

Both releases are primarily housekeeping; no major feature flags were introduced.

---

## Hot Issues (Top 10 by Impact & Community Interest)

1. **[#21409 – Generalist agent hangs](https://github.com/google-gemini/gemini-cli/issues/21409)**  
   `P1, bug, 8 👍`  
   The most upvoted open bug. Users report the generalist agent hangs indefinitely on simple tasks (e.g., folder creation). Workaround: instruct the model not to use subagents. Still open after three months.

2. **[#22323 – Subagent recovery after MAX_TURNS falsely reports success](https://github.com/google-gemini/gemini-cli/issues/22323)**  
   `P1, bug`  
   `codebase_investigator` subagent returns `status: "success"` even after hitting the turn limit, hiding the interruption. Misleading termination reason undermines trust in agent logging.

3. **[#25166 – Shell command execution gets stuck "Waiting input"](https://github.com/google-gemini/gemini-cli/issues/25166)**  
   `P1, bug, 3 👍`  
   After a simple CLI command finishes, Gemini CLI remains stuck showing the command as active. Intermittent but severe – disrupts interactive workflows.

4. **[#21983 – Browser subagent fails on Wayland](https://github.com/google-gemini/gemini-cli/issues/21983)**  
   `P1, bug, 1 👍`  
   The browser subagent crashes with a `GOAL` termination on Wayland-based Linux environments. Still awaiting retesting after initial fix attempts.

5. **[#24353 – Robust component level evaluations (Epic)](https://github.com/google-gemini/gemini-cli/issues/24353)**  
   `P1, customer-issue, area/agent`  
   Epic tracking the expansion of behavioral eval tests (76 so far) across 6 supported Gemini models. Critical for preventing regressions in agent behavior.

6. **[#22745 – Assess AST-aware file reads, search, and mapping](https://github.com/google-gemini/gemini-cli/issues/22745)**  
   `P2, customer-issue, 1 👍`  
   Investigation into whether Abstract Syntax Tree (AST) aware tools can improve codebase understanding, reduce token waste, and improve navigation accuracy.

7. **[#26525 – Add deterministic redaction and reduce Auto Memory logging](https://github.com/google-gemini/gemini-cli/issues/26525)**  
   `P2, security`  
   Auto Memory sends transcripts to model context before redaction. This issue demands deterministic secret redaction earlier in the pipeline and reduction of skill logging.

8. **[#26522 – Stop Auto Memory from retrying low-signal sessions indefinitely](https://github.com/google-gemini/gemini-cli/issues/26522)**  
   `P2, bug`  
   Auto Memory only marks a session processed when the extraction agent successfully reads it; low-signal sessions are re-surfaced repeatedly, wasting LLM calls and memory.

9. **[#22672 – Agent should stop/discourage destructive behavior](https://github.com/google-gemini/gemini-cli/issues/22672)**  
   `P2, customer-issue, 1 👍`  
   Model occasionally uses `git reset --force` or similar destructive commands when safer alternatives exist. Highlights a need for safety guardrails in tool execution.

10. **[#28001 – Nightly Release Failed for v0.48.0-nightly](https://github.com/google-gemini/gemini-cli/issues/28001)**  
    `P0, release-failure`  
    The nightly build workflow failed on 2026-06-18. P0 priority – blocks all downstream testing and preview consumption.

---

## Key PR Progress

1. **[#28000 – fix(core-tools): resolve Jupyter Notebook and JSON corruption in write_file](https://github.com/google-gemini/gemini-cli/pull/28000)**  
   `size/m, open`  
   Critical fix: `write_file` silently corrupts `.ipynb` and `.json` files, causing environments like Colab to discard changes. Root cause isolated and patched.

2. **[#27987 – fix(cli): throw FatalConfigError instead of process.exit](https://github.com/google-gemini/gemini-cli/pull/27987)**  
   `P1, size/xs, open`  
   Refactors argument parsing to throw errors instead of calling `process.exit()`, resolving E2E test hangs for `--help`/`--version`.

3. **[#27996 – fix(core): decode response body using charset from Content-Type header](https://github.com/google-gemini/gemini-cli/pull/27996)**  
   `area/agent, size/m, open`  
   `web-fetch` now respects charset parameters (e.g., GBK, ISO-8859-1), fixing garbled text on Chinese/Japanese and legacy sites.

4. **[#27994 – fix(core): insert skill/agent content literally in system prompt substitutions](https://github.com/google-gemini/gemini-cli/pull/27994)**  
   `area/agent, size/s, open`  
   `String.prototype.replace` was interpreting `$` characters in skill content as backreferences, causing replacement corruption. Now uses literal insertion.

5. **[#27948 – chore(deps): pin dependencies and enforce 14-day update cooldown](https://github.com/google-gemini/gemini-cli/pull/27948)**  
   `size/xl, open`  
   Locks all direct dependencies to exact versions and introduces a cooldown period for automated updates to reduce breakage risk.

6. **[#27997 – docs: remove references to deprecated consumer and free tiers](https://github.com/google-gemini/gemini-cli/pull/27997)**  
   `size/l, open`  
   Cleans up documentation after Google’s announcement that consumer/free tiers (Gemini Code Assist for individuals, Free Tier) will stop serving requests on June 14.

7. **[#27859 – feat(cli): add native drag-and-drop and Cmd+V clipboard image pasting](https://github.com/google-gemini/gemini-cli/pull/27859)**  
   `size/m, open`  
   Adds first-class terminal multimodal support – drag-and-drop files and paste images via clipboard in standard terminal emulators.

8. **[#27780 – security: gate chained E2E on same-repository checkout for workflow_run](https://github.com/google-gemini/gemini-cli/pull/27780)**  
   `size/xs, open`  
   Prevents fork PRs from injecting attacker-controlled `repository` + `sha` in CI jobs that use `GEMINI_API_KEY`. Guards against supply-chain attacks.

9. **[#27783 – security: gate PRT label workflows on same-repository pull_request_target](https://github.com/google-gemini/gemini-cli/pull/27783)**  
   `size/xs, open`  
   Ensures automated PR size/rate-limit labeling only runs for PRs from the same repository, not external forks.

10. **[#27648 – feat(core): support list format in trustedFolders.json](https://github.com/google-gemini/gemini-cli/pull/27648)**  
    `size/m, closed`  
    Adds support for a simple JSON array format in `trustedFolders.json` alongside the existing object format, making it easier for users to maintain trusted directories manually.

---

## Feature Request Trends

- **AST-aware tooling** (Issues #22745, #22747) – Users and maintainers strongly desire integration of AST-based code search, file reads, and codebase mapping to reduce token overhead and improve precision.
- **Agent self-awareness & configuration** (#21432, #22672) – Multiple requests for the agent to accurately describe its own flags, hotkeys, and capabilities, and to avoid destructive commands without confirmation.
- **Browser agent resilience** (#22232, #22267) – Calls for automatic session takeover, lock recovery, and proper merging of user settings (e.g., `maxTurns`).
- **Memory system improvements** (#26522, #26523, #26525) – Users want deterministic redaction, better handling of low-signal sessions, and detection of invalid memory patches.
- **Remote agents & background operations** (#20303) – Ongoing epic for advanced authentication and background processing in remote agent mode.

---

## Developer Pain Points

1. **Agent hangs and false success reporting** – The generalist agent hangs indefinitely (#21409) and subagents falsely report `GOAL` after turn limits (#22323). Both erode user confidence in autonomous execution.
2. **Shell command stuck states** (#25166) – Simple commands that finish remain marked as “awaiting input”, forcing manual cancellation. Intermittent but highly disruptive.
3. **Platform-specific failures** – Browser agent fails on Wayland (#21983); macOS symlink path mismatches in tests (#27990). Users on non-standard platforms face workarounds.
4. **Tool limit errors** (#24246) – Users encounter 400 errors when >128 tools are enabled, with no built-in tool filtering.
5. **Memory system waste** – Auto Memory retries low-signal sessions forever (#26522) and sends transcripts to model before redacting secrets (#26525), leading to token waste and privacy risks.
6. **Configuration ignored** – Browser agent ignores `settings.json` overrides (#22267); duplicate agent name warnings in home directory (#27995). Users expect consistent behavior.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-06-18

## Today’s Highlights
This week’s activity is dominated by the fallout from the June 16 Copilot service outage, with users reporting models appearing “blocked/disabled” and transient API errors stalling workflows. Meanwhile, a long-standing plugin hook issue (#2643) preventing silent command rewrites continues to frustrate developers, and the community is strongly rallying behind a request for a tool whitelist in interactive mode (#1973). No new releases or pull requests landed in the last 24 hours.

---

## Releases
*None in the last 24 hours.*

---

## Hot Issues

### 1. [area:plugins] `preToolUse` silent command rewrite still shows confirmation dialog (#2643)
**Author:** jeziellopes | **Comments:** 10 | **👍:** 1  
Even when a plugin hook sets `permissionDecision: allow`, the CLI pops an interactive confirmation on every rewritten command. Plugin developers see this as a blocker for building truly autonomous hooks.  
[🔗 Issue #2643](https://github.com/github/copilot-cli/issues/2643)

### 2. [area:permissions] Feature Request: Tool whitelist for Interactive Mode (#1973)
**Author:** Dicer-J | **Comments:** 10 | **👍:** 20  
Users want to allow safe read-only tools (grep, cat, git log) without the manual approval overhead or falling back to the risky `/allow-all` command. This is the most upvoted open feature request.  
[🔗 Issue #1973](https://github.com/github/copilot-cli/issues/1973)

### 3. [more-info-needed] Repeated login prompts (#254)
**Author:** yurivict | **Comments:** 9 | **👍:** 4  
Despite being logged in, the CLI constantly re-asks for authentication after Ctrl-C or between sessions. Affects GitHub Business users. Open for months with no resolution.  
[🔗 Issue #254](https://github.com/github/copilot-cli/issues/254)

### 4. [area:models, area:tools] Duplicate `fc_call` ID error (#3560)
**Author:** lgphp | **Comments:** 5 | **👍:** 1  
Sudden `400` error with “Duplicate item found with id” only after tool calls. Plain chat works fine. Points to a server-side state issue or CAPI regression.  
[🔗 Issue #3560](https://github.com/github/copilot-cli/issues/3560)

### 5. [Bug] All models show as “Blocked/Disabled” after June 16 outage (#3832)
**Author:** yzeng58 | **Comments:** 5 | **👍:** 13  
Following the 30-minute Copilot outage on June 16, the model selection UI displays everything as disabled. Closed by maintainers, but no root cause or permanent fix was shared.  
[🔗 Issue #3832](https://github.com/github/copilot-cli/issues/3832)

### 6. [Bug] Transient API error retry loop (#3831)
**Author:** chanman4821 | **Comments:** 4 | **👍:** 2  
The CLI entered an infinite “Request failed due to a transient API error. Retrying…” loop mid-workflow. Closed as a transient outage side-effect, but users are concerned about resilience.  
[🔗 Issue #3831](https://github.com/github/copilot-cli/issues/3831)

### 7. [area:context-memory] Configurable context window for Claude Opus 4.6 (#3355)
**Author:** aksingh | **Comments:** 3 | **👍:** 4  
Copilot CLI caps Opus 4.6 at 200K tokens despite the model supporting 1M. Heavy users report frequent forced compaction during deep troubleshooting sessions.  
[🔗 Issue #3355](https://github.com/github/copilot-cli/issues/3355)

### 8. [area:enterprise, area:models] Support Enterprise custom models in CLI (#3730)
**Author:** sebdanielsson | **Comments:** 2 | **👍:** 4  
Enterprise admins can configure custom models for Copilot in VS Code and other clients, but these models do not appear in the CLI. This is a gap for organisations using private endpoints.  
[🔗 Issue #3730](https://github.com/github/copilot-cli/issues/3730)

### 9. [triage] Ollama Cloud incompatible with Copilot CLI’s `custom_tool_call` payload (#3839)
**Author:** weweaaa | **Comments:** 1 | **👍:** 7  
When using BYOK models through Ollama Cloud in Fleet Mode, requests fail with a `400` error because Ollama doesn’t understand the custom tool call format. High community interest in BYOK compatibility.  
[🔗 Issue #3839](https://github.com/github/copilot-cli/issues/3839)

### 10. [triage] Content-exclusion policy incorrectly enforced on CLI (#3841)
**Author:** quoma | **Comments:** 0 | **👍:** 0  
The CLI is blocking local file reads due to organization content‑exclusion policies, even though GitHub’s docs state these exclusions only apply to code review. A policy bug that breaks local agent workflows.  
[🔗 Issue #3841](https://github.com/github/copilot-cli/issues/3841)

---

## Key PR Progress
*No pull requests were updated in the last 24 hours.*

---

## Feature Request Trends
The community is pushing heavily for **granular permission control** in interactive mode (#1973) rather than the binary all-or-nothing allow/deny. **Plugin hook flexibility** is another strong theme: developers want silent rewrites (#2643), proper matchers for `postToolUse` (#3820), and the ability for skill files to declare additional MCP servers (#3292). On the model side, users are asking for **more configurability** – an `/effort` command to quickly adjust reasoning effort (#3074), the full context window for Claude Opus 4.6 (#3355), and Enterprise custom model support (#3730). **Session management** improvements include persistent `/instructions` opt-out (#3840), custom aliases for frequent model/prompt combos (#3844), and showing the working directory when resuming sessions (#3837).

---

## Developer Pain Points
**Post-outage instability** is the dominant pain point this week: models appearing disabled (#3832), infinite retry loops (#3831), and lingering API errors (#3560). The **repeated login prompt** bug (#254) continues to degrade the experience, especially for business users. **Plugin/MCP friction** is increasing – the `plugin install` command fails when Git’s fsmonitor is enabled (#3842), sub‑agents lose access to MCP tools (#3812), and MCP tools are only lazy-loaded, causing some agents never to see them (#3787). **Session corruption** is a concern: malformed attachments (e.g., password-protected Excel files) can poison the entire session (#3791), and resuming sessions with spaces in the name silently fails (#3754). Lastly, **confusing policy enforcement** (#3841) undermines trust in local agent workflows when organisation content policies are incorrectly applied.

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-06-18

## Today's Highlights
Today's activity is light, with two new feature requests but no releases or pull requests merged in the last 24 hours. The community is asking for runtime execution mode switching (Agent ↔ Cluster) and an SSL certificate ignore option, reflecting ongoing needs around enterprise network configurations and flexible workflow control.

## Releases
No new releases in the last 24 hours.

## Hot Issues
*(Only 2 issues were updated in the last 24h; listed below)*

1. **[#2459] [Feature Request] 支持会话运行中切换执行模式（Agent ↔ 集群）**  
   - **Why it matters:** Enables users to switch between single-agent and cluster modes mid-session without restart, critical for dynamic scaling and debugging workflows.  
   - **Community reaction:** 0 comments, 0 👍 – likely still early-stage discussion.  
   - [GitHub Link](https://github.com/MoonshotAI/kimi-cli/issues/2459)

2. **[#2458] [enhancement] Add option to ignore ssl certificate**  
   - **Why it matters:** User reports that organization-managed antivirus software injects its own SSL certificate (MITM), breaking authentication. A flag to bypass certificate validation would unblock users in restricted corporate environments.  
   - **Community reaction:** 0 comments, 0 👍 – but the issue description clearly articulates a real enterprise pain point.  
   - [GitHub Link](https://github.com/MoonshotAI/kimi-cli/issues/2458)

## Key PR Progress
No pull requests were updated in the last 24 hours.

## Feature Request Trends
Based on the latest issues, two feature directions stand out:

- **Runtime execution mode flexibility:** The request for mid-session switching between Agent and Cluster modes (#2459) suggests a desire for more adaptive resource management without restarting workflows.
- **Enterprise security compatibility:** The SSL certificate bypass request (#2458) highlights a growing need to support corporate-managed environments with custom certificate authorities or MITM proxies.

## Developer Pain Points
- **Corporate network restrictions:** The SSL certificate issue (#2458) is a classic pain point for developers behind corporate firewalls or antivirus software that rewrites TLS traffic. The lack of a `--insecure` or `--ssl-no-verify` flag blocks initial login and API access.
- **Session inflexibility:** The inability to change execution mode mid-session (Agent ↔ Cluster) forces developers to kill and restart sessions, disrupting long-running tasks or iterative debugging. This is a high-frequency request in multi-node/parallel execution tools.

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-18

## Today’s Highlights
OpenCode **v1.17.8** shipped yesterday with faster session timeline rendering and bugfixes for OpenAI-compatible MCP schemas and Cloudflare AI Gateway. Community attention is centered on a long-standing feature request for **tokens-per-second display** (#6096, 55 👍), and several **Windows terminal corruption** issues continue to generate discussion. A new wave of contributions focuses on **fuzzy edit matching** in the V2 core tool and **auto‑discovery of models** from OpenAI‑compatible providers.

---

## Releases
### [v1.17.8](https://github.com/anomalyco/opencode/releases/tag/v1.17.8)
- **Session timelines** now load faster and avoid flicker/scroll jumps.
- **Bugfix** – OpenAI‑compatible providers accept MCP tool schemas that previously failed validation (thanks @jquense).  
- **Bugfix** – Cloudflare AI Gateway receives the configured API key correctly (thanks @keefetang).

---

## Hot Issues (Top 10)

1. **[#6096 – Feature: Adding Experimental Calculation and Display of Tokens per second](https://github.com/anomalyco/opencode/issues/6096)**  
   *55 👍, 18 comments*  
   Users strongly desire a Tokens‑Per‑Second (TPS) metric for each message response. The high upvote count signals that performance observability is a top community priority.

2. **[#23566 – Docs suggest LSP is enabled by default](https://github.com/anomalyco/opencode/issues/23566)**  
   *20 👍, 10 comments*  
   Documentation claims LSP auto‑installs for Kotlin, but LSP is actually disabled by default. The mismatch has caused confusion and wasted setup time.

3. **[#32172 – Feature: Add GLM‑5.2 model support for Z.AI provider](https://github.com/anomalyco/opencode/issues/32172)**  
   *0 👍, 10 comments (closed)*  
   Request to add Z.AI’s latest reasoning model. Closed quickly, likely added or rejected; valuable for users on Z.AI plans.

4. **[#24817 – Ctrl+Z closes/suspends OpenCode instead of undoing text input (Linux)](https://github.com/anomalyco/opencode/issues/24817)**  
   *2 👍, 5 comments*  
   Classic keybinding conflict – Ctrl+Z sends SIGTSTP instead of text undo. Linux users find this disruptive.

5. **[#7928 – Runtime Permission Mode Toggle (like Claude Code’s Shift+Tab)](https://github.com/anomalyco/opencode/issues/7928)**  
   *17 👍, 5 comments*  
   High demand for an interactive toggle between auto‑edit and permission‑request modes, avoiding static config edits.

6. **[#31204 – SQLite NOT NULL constraint failed on agent‑switched sessions](https://github.com/anomalyco/opencode/issues/31204)**  
   *3 👍, 7 comments (closed)*  
   A critical database migration bug from June 3‑5 breaks sessions that trigger agent switches. Closed with a fix, but caused data loss for some users.

7. **[#31119 – Error: no such column: name](https://github.com/anomalyco/opencode/issues/31119)**  
   *5 👍, 4 comments*  
   Users returning after an update face a blocking SQL schema issue. Suggests a missing migration for the `name` column.

8. **[#32745 – Forever stuck in "Authorization in progress…" (OpenRouter desktop)](https://github.com/anomalyco/opencode/issues/32745)**  
   *0 👍, 4 comments (closed)*  
   OAuth flow hangs when connecting OpenRouter on the desktop app. Likely a redirect URI mismatch.

9. **[#21277 – Terminal left with raw ANSI escape codes after crash (PowerShell)](https://github.com/anomalyco/opencode/issues/21277)**  
   *0 👍, 3 comments*  
   After OpenCode crashes on Windows, the terminal displays literal ANSI sequences and the cursor disappears, requiring a manual reset.

10. **[#32749 – Explore agent is a huge waste of tokens](https://github.com/anomalyco/opencode/issues/32749)**  
    *0 👍, 2 comments*  
    The Explore subagent spawns unconditionally, even for trivial tasks where a simple grep would suffice. Users report massive token waste and duplicated file reads.

---

## Key PR Progress (Top 10)

1. **[#32774 – fix(opencode): preserve execution metadata on tool completion](https://github.com/anomalyco/opencode/pull/32774)**  
   Fixes subagent task entries becoming unclickable when result metadata drops `sessionId`. Closes #32773.

2. **[#32771 – feat(tui): show assistant completion time](https://github.com/anomalyco/opencode/pull/32771)**  
   Adds turn completion timestamps to native run summaries and legacy TUI, improving observability of response latency.

3. **[#32767 – fix(tui): restore ESC interrupt for delegated subagent sessions](https://github.com/anomalyco/opencode/pull/32767)**  
   Regression fix – ESC now correctly interrupts subagent sessions. Closes three related issues (#3699, #4073, #23534).

4. **[#23688 – feat(app): add markdown preview with mermaid diagram support](https://github.com/anomalyco/opencode/pull/23688)**  
   Long‑awaited feature: preview markdown files including Mermaid diagrams directly in the desktop app.

5. **[#32766 – feat(core): accept explicit storage in public API layer](https://github.com/anomalyco/opencode/pull/32766)**  
   Extracts the database layer out of session composition, enabling injectable storage for tests and embeddings. Closes #32764.

6. **[#32765 – chore(opencode): code cleanup, formatter consolidation, and perf improvements](https://github.com/anomalyco/opencode/pull/32765)**  
   Removes dead types, merges Ruff/uv into a shared formatter, optimizes message normalization. Closes #32763.

7. **[#32762 – fix(skill): prevent recursive sub‑skill discovery using single‑level glob](https://github.com/anomalyco/opencode/pull/32762)**  
   Stops nested skill directories from being loaded as independent skills. Closes #28485.

8. **[#32761 – feat(core): port V1 fuzzy edit matching to V2 core edit tool](https://github.com/anomalyco/opencode/pull/32761)**  
   Ports 9 fuzzy replacer strategies and Levenshtein matching to the V2 tool, making LLM edits more robust to whitespace/indentation differences. Closes #32760.

9. **[#32758 – fix(opencode): re-read plugin.trigger output to support array replacement](https://github.com/anomalyco/opencode/pull/32758)**  
   Plugin hooks that reassign `output.messages` now take effect. Root cause: call sites discarded the modified output object.

10. **[#32753 – fix(web): add clipboard fallback for non‑HTTPS contexts](https://github.com/anomalyco/opencode/pull/32753)**  
    Copy‑to‑clipboard now works on localhost without HTTPS using a textarea‑based fallback. Closes #32664.

---

## Feature Request Trends
- **New model/provider support** – Requests for GLM‑5.2 (Z.AI & Ollama), auto‑discovery of models from OpenAI‑compatible endpoints, and better local Ollama integration.  
- **Performance observability** – Tokens‑per‑second display (#6096) and completion timestamps (#32771) reflect demand for runtime metrics.  
- **Permission and safety controls** – A runtime permission mode toggle (#7928) and interactive confirmation for file edits remain highly upvoted.  
- **Windows terminal compatibility** – Continued friction with ANSI escape codes, Ctrl+Z behavior, and terminal reset after crashes.  
- **Editing reliability** – Fuzzy matching, optional tool arguments, and better edit tool fallbacks are recurring themes.  
- **Session management** – Global session list scope, interactive picker, and session metadata preservation are being actively built.

---

## Developer Pain Points
- **Windows terminal stability** – Raw ANSI escape codes after crashes (#21277), broken Ctrl+Z (#24817), and `Get-Command` errors after re‑entering the terminal (#32757) plague PowerShell users.  
- **Database migration fragility** – `NOT NULL constraint failed` (#31204) and `no such column: name` (#31119) indicate insufficient migration coverage across version upgrades.  
- **Subagent resource waste** – The Explore agent (#32749) and recursive skill discovery (#32762) cause unnecessary token consumption and confusion.  
- **MCP integration friction** – Optional arguments not omitted (#32772) and schema validation failures (fixed in v1.17.8) require ongoing attention.  
- **Authentication flows** – OpenRouter OAuth hangs (#32745) and “Authorization in progress” loops are recurrent complaints.  
- **Keybinding conflicts** – `Ctrl+X` near `Ctrl+C` (quit) (#23322) and `Ctrl+Z` suspending instead of undo (#24817) are consistent frustration signals.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest – 2026-06-18

## Today’s Highlights
The community continues to push Pi’s provider ecosystem with new Azure AI Foundry and Copilot support, while the core team addresses long-standing compaction inefficiencies and terminal UX bugs. A high-impact fix for Anthropic 1‑hour cache pricing lands, and a Nix flake packaging PR opens the door to declarative installs.

## Releases
No new releases in the last 24 hours.

## Hot Issues

1. **[#5825 – Streaming markdown forces scroll to bottom](https://github.com/earendil-works/pi/issues/5825)** (12 comments, inprogress)  
   A persistent annoyance: when `clear on shrink` is enabled, the TUI aggressively scrolls to the bottom, preventing the user from reading earlier content. Community strongly sympathises – fixes are already in sight via PR #5846.

2. **[#5653 – Move off Shrinkwrap](https://github.com/earendil-works/pi/issues/5653)** (11 comments, inprogress)  
   Duplicate `pi-ai` copies caused by npm shrinkwrap, leading to separate module-level registries. A blocker for multi‑package dependency trees. High engagement points to wide impact.

3. **[#3715 – local-llm streams terminate at 5 min](https://github.com/earendil-works/pi/issues/3715)** (11 comments, closed)  
   Long `Write` tool calls against local LLMs (vLLM, Qwen3) die after 5 min due to undici’s internal `bodyTimeout` which cannot be overridden by provider retry settings. A known pain for users running heavy local models.

4. **[#534 – Config folder out of place on Linux](https://github.com/earendil-works/pi/issues/534)** (9 comments, closed, 👍20)  
   XDG violation: config stored in `$HOME/.pi` instead of `$XDG_CONFIG_HOME`. Long-running, highly upvoted – finally closed (likely via a recent refactor?). Symbolic of community care about platform conventions.

5. **[#5654 – Add `excludeFromContext` to custom messages](https://github.com/earendil-works/pi/issues/5654)** (7 comments, open)  
   Request to allow custom messages (e.g. `/status`) to be displayed without being sent to the LLM. A clean UX separation that mirrors existing `!!` bash execution messages.

6. **[#5821 – Support Anthropic OAuth Subscription Usage in Agent SDK](https://github.com/earendil-works/pi/issues/5821)** (7 comments, closed)  
   Clarification that Anthropic Agent SDK works with existing subscriptions. Users want seamless integration, but the short thread suggests it’s mostly documentation / configuration.

7. **[#5830 – Tree navigator truncates long entries](https://github.com/earendil-works/pi/issues/5830)** (4 comments, closed)  
   In the `/tree` view (double‑Escape), long entries are silently truncated. Bad UX for anyone with deeply nested or verbose session histories. Quickly fixed (PR likely merged).

8. **[#5810 – RPC: expose session entries and tree](https://github.com/earendil-works/pi/issues/5810)** (3 comments, open)  
   Adding `get_entries` and `get_tree` RPC commands – enables driving Pi from external scripts (e.g. Neovim). Low comments but high strategic value for automation.

9. **[#5770 – GLM-5.2 effort level configuration](https://github.com/earendil-works/pi/issues/5770)** (3 comments, closed)  
   Zhipu’s GLM-5.2 supports `high` / `max` effort levels, but Pi only had “on/off”. Quickly merged, shows responsiveness to new model capabilities.

10. **[#5848 – Compaction: percentage-based trigger](https://github.com/earendil-works/pi/issues/5848)** (1 comment, closed)  
    Allow auto-compaction to trigger at a percentage of the model’s context window instead of an absolute token reserve. Essential for large‑context models where one size doesn’t fit all.

## Key PR Progress

1. **[#5846 – fix(tui): stabilize streaming code fence rendering](https://github.com/earendil-works/pi/pull/5846)** (open, closes #5825)  
   Direct fix for the scroll‑to‑bottom bug. Still open, likely awaiting review.

2. **[#5738 – fix(ai): price anthropic 1h cache writes at 2x input](https://github.com/earendil-works/pi/pull/5738)** (closed, merged)  
   Prior to this, all cache writes were priced at the 5‑minute rate, undercounting 1‑hour writes by 1.6×. Now reads `ephemeral_1h_input_tokens` and charges correctly.

3. **[#631 – fix(ai): Google thinking detection + remove unsupported id fields](https://github.com/earendil-works/pi/pull/631)** (closed, merged)  
   `isThinkingPart()` was over‑matching on `thoughtSignature`. Now correctly excludes `thought` parts from thinking detection. Also removes fields that Google API rejects.

4. **[#5859 – fix(ai): send responses prompts as instructions](https://github.com/earendil-works/pi/pull/5859)** (open)  
   OpenAI Responses APIs require system prompts at top‑level `instructions`, not as replayed `input` messages. Fixes provider compatibility for OpenAI / Azure / Codex.

5. **[#5850 – chore(deps): bump vitest to 3.2.6 and override esbuild](https://github.com/earendil-works/pi/pull/5850)** (closed, merged)  
   Mechanical bump that resolves 5 of 6 `npm audit` high advisories. Zero runtime change – but keeps CI green.

6. **[#5849 – feat(ai): add Azure AI Foundry provider for Anthropic Claude](https://github.com/earendil-works/pi/pull/5849)** (closed, merged)  
   New `azure-foundry` provider with full Python SDK parity (base URL, headers, Entra ID auth). Brings Pi to Azure enterprise users.

7. **[#5847 – Comath/research exploration mode](https://github.com/earendil-works/pi/pull/5847)** (closed, merged)  
   Draft checkpoint for a “co‑math research” prototype: research paths, async workstreams, source‑backed data intake. Possibly a new feature area – watch this space.

8. **[#5812 – fix(tui): protect pipe characters inside inline code in markdown tables](https://github.com/earendil-works/pi/pull/5812)** (closed, merged)  
   Pipe `|` inside backticks within table cells was being split as a column delimiter. Custom tokeniser now escapes them – a classic markdown parser edge case.

9. **[#5841 – feat(tui): detect Warp terminal and enable Kitty image protocol](https://github.com/earendil-works/pi/pull/5841)** (open)  
   Detects Warp via environment variables and enables Kitty graphics + OSC 8 hyperlinks without workarounds. Warp users rejoice.

10. **[#5832 – fix(ai): surface provider HTTP error body](https://github.com/earendil-works/pi/pull/5832)** (open, fixes #5763)  
    Behind a proxy/gateway, non‑2xx responses with unparseable bodies were swallowed. Now raw HTTP body is surfaced instead of opaque SDK messages – a huge debugging win.

## Feature Request Trends
- **Provider expansion**: Constant requests for new model support – GLM‑5.2 effort levels, Azure Foundry, Copilot 1M‑token window, Opencode Go subscriptions, Mistral prompt caching. Indicates Pi is seen as a universal frontend for many backends.
- **RPC & automation**: Several issues (e.g. #5810, #5654, #5861) push for richer APIs to drive Pi programmatically – session tree access, custom persistence, extension events.
- **Compaction intelligence**: A cluster of issues (#5848, #5845, #5844) shows users want smarter compaction triggers (percentage‑based, less aggressive under local LLMs) and fewer edge‑case crashes.

## Developer Pain Points
- **Shrinkwrap / duplicate dependencies** (#5653) remains a thorny issue for anyone using Pi as a library (two copies of `pi-ai` break module‑level registries).
- **Provider timeout tuning** (#3715): The undici `bodyTimeout` is invisible and uncapped by Pi’s own retry settings, frustrating local‑model users.
- **Terminal UX regressions**: Streaming scroll force (#5825) and truncated tree entries (#5830) show that even mature features can break with new TUI changes.
- **Opaque error messages** (#5832): Provider SDKs mask real HTTP errors; the community welcomes PRs that surface raw bodies.
- **XDG compliance** (#534): Linux users expect config files in `~/.config` – the long delay in fixing this (months) eroded trust in platform awareness.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-06-18

## Today’s Highlights
The team shipped three releases today, including stable `v0.18.3` with core fixes for file history tracking and CLI cancellation. A critical OOM bug during `/quit` auto-memory extraction was addressed in PR #5181, and a community‑reported model‑provider disambiguation issue (#5173) was resolved via PR #5179. Meanwhile, new feature work continues on a QQ Bot channel adapter (#5202) and macOS Liquid Glass icon support (#5284).

---

## Releases
Three new versions were published in the past 24 hours:

- **[v0.18.3-nightly](https://github.com/QwenLM/qwen-code/releases/tag/v0.18.3-nightly.20260618.bc3e0b405)** – Contains `fix(core): Track supported sed edits in file history`.
- **[v0.18.3](https://github.com/QwenLM/qwen-code/releases/tag/v0.18.3)** – Includes `fix(cli): Stop after cancelled ask_user_question` and the core fix from the nightly.
- **[v0.18.3-preview.0](https://github.com/QwenLM/qwen-code/releases/tag/v0.18.3-preview.0)** – Same changes as v0.18.3 but as a preview tag.
- **[v0.18.2](https://github.com/QwenLM/qwen-code/releases/tag/v0.18.2)** – Warns on oversized context instructions and updates documentation for CLI syntax and tool naming.

No changes were listed for v0.18.1-preview.1.

---

## Hot Issues
*(All 6 issues updated in the last 24h are listed below.)*

1. **[#5173 — Model provider disambiguation fails when multiple providers share the same model id](https://github.com/QwenLM/qwen-code/issues/5173)**  
   *Priority P2, Bug, Closed.* A long‑standing frustration: selecting a provider (e.g., IdeaLab’s `qwen3.7-max`) did not persist across sessions. Community reported the issue is now fixed in PR #5241. 3 comments, low engagement.

2. **[#5280 — `test(cli): re-enable long command search suggestion coverage`](https://github.com/QwenLM/qwen-code/issues/5280)**  
   *Priority P3, Enhancement, Open.* A test for command‑search expand/collapse is still skipped; the behavior works locally. The author offers to take it. This reflects a push to improve test reliability.

3. **[#5277 — Re-enable TableRenderer foreground reset coverage](https://github.com/QwenLM/qwen-code/issues/5277)**  
   *Priority P3, Enhancement, Open.* Skipped test for `\x1b[39m` reset. Functional assertions pass but snapshots are stale. Minor but signals ongoing test cleanup.

4. **[#5275 — Re-enable BaseSelectionList scroll-up coverage](https://github.com/QwenLM/qwen-code/issues/5275)**  
   *Priority P3, Enhancement, Open.* Another skipped test due to React `act` warnings. Community member is willing to fix with `act`-wrapped rendering. Good for contributor onboarding.

5. **[#2845 — Please perform normal file content recognition](https://github.com/QwenLM/qwen-code/issues/2845)**  
   *Feature request, Closed.* Request to treat `.dat` files with text content as text, not binary. Closed after PR #5256 was merged today. Shows the project responds to community pain points around file handling.

6. **[#5281 — TUI becomes unresponsive when `login1.inhibit-block-sleep` triggers authentication](https://github.com/QwenLM/qwen-code/issues/5281)**  
   *Priority P2, Bug, Open.* New issue: on Linux SSH sessions without a desktop environment, the sleep‑inhibit feature prompts for PolKit auth which hijacks the TUI input stream. One comment, likely will be escalated. Affects users running Qwen Code remotely.

---

## Key PR Progress
*(10 important PRs updated in the last 24h)*

1. **[#5284 — `feat(desktop): compile macOS 26+ Liquid Glass Assets.car in brand-create`](https://github.com/QwenLM/qwen-code/pull/5284)**  
   Automates bundling of macOS 26 Liquid Glass icons into custom brands. Polishes the desktop packaging experience for Mac users.

2. **[#5202 — `feat(channel): add QQ Bot channel adapter`](https://github.com/QwenLM/qwen-code/pull/5202)**  
   Adds a new `@qwen-code/channel-qqbot` adapter for the QQ messaging platform, joining Telegram/WeChat/Feishu. Expands cross‑platform reach (WebSocket gateway, heartbeats, resume logic).

3. **[#5030 — `feat(core,cli,sdk): resume an interrupted turn without a synthetic "continue" message`](https://github.com/QwenLM/qwen-code/pull/5030)**  
   A sophisticated feature to continue a session after crash or interruption without injecting a phony `"continue"` user message. Classifies the continuation state into three shapes. Improves conversational UX.

4. **[#5145 — `feat(cli): show follow-up suggestion in input placeholder`](https://github.com/QwenLM/qwen-code/pull/5145)**  
   Displays the model’s suggested next prompt directly in the input placeholder, using the fast model for generation. Reduces cognitive load for users.

5. **[#5231 — `feat(core,cli): Workflow tool token budget + per-run UI surfacing`](https://github.com/QwenLM/qwen-code/pull/5231)**  
   Adds a per‑run output‑token budget for the Workflow tool, plumbed through orchestrator dispatch and UI dialogs. Helps control costs in long‑running workflows.

6. **[#5279 — `fix(core): add always-on tool-call circuit breaker for runaway loops`](https://github.com/QwenLM/qwen-code/pull/5279)**  
   Focused re‑scope of an earlier PR. Adds a circuit breaker to prevent infinite tool‑call loops. Addresses a safety issue noted in #5234.

7. **[#5181 — `fix(core): prevent OOM in auto-memory extraction during /quit`](https://github.com/QwenLM/qwen-code/pull/5181)**  
   Fixes a heap‑limit crash when `/quit` triggers memory extraction on large conversation histories. Uses chunked processing to avoid `buildTranscriptMessages()` overhead.

8. **[#2915 — `feat(cli): add /clear --all to fully reset session including IDE/editor context`](https://github.com/QwenLM/qwen-code/pull/2915)**  
   Long‑standing PR (stale for months but active today). Adds a `--all` flag to `/clear` to wipe IDE/editor context as well. Interactive confirmation guard. Useful for troubleshooting.

9. **[#5220 — `feat(i18n): localize tool display names in TUI and web-shell badges`](https://github.com/QwenLM/qwen-code/pull/5220)**  
   Routes tool‑call badges (e.g., `TodoWrite`, `Shell`) through i18n, so they show translated names when UI language is switched to Chinese. Improves internationalization consistency.

10. **[#5256 — `fix(core): detect dat files by content`](https://github.com/QwenLM/qwen-code/pull/5256)**  
    Closes issue #2845. `.dat` files are no longer blindly treated as binary; text‑content `.dat` files are read as text, while binary ones are still handled correctly. A simple but impactful fix.

---

## Feature Request Trends
From the Issues and PRs, the following user‑driven themes emerge:

- **Multi‑provider model disambiguation** – Users with several API endpoints sharing the same model ID need persistent selection (#5173, #5241). Now resolved.
- **Content‑based file handling** – The request to treat ambiguous extensions (`.dat`) by content is a recurring theme (#2845, #5256). Expect more requests for similar handling of other extensions.
- **TUI stability on Linux/SSH** – The sleep‑inhibit authentication issue (#5281) highlights usability problems when Qwen Code runs headless or over SSH. Likely to attract more attention.
- **Chinese i18n completeness** – PR #5220 and #2993 show strong demand for full localization of tool names and commands. The community is actively contributing translations.
- **Test infrastructure improvements** – Multiple issues (#5280, #5277, #5275) focus on re‑enabling skipped tests. While not user features, they indicate a project hardening for regression coverage.

---

## Developer Pain Points
- **OOM crashes on large conversations** – The `/quit` memory extraction bug (#5181) caused heap exhaustion, a clear pain point for users with long sessions.
- **TUI hijacking by system auth prompts** – The sleep‑inhibit feature (#5281) broke SSH sessions by capturing input streams. Workaround: disable the setting. Expect a fix soon.
- **Stale snapshots blocking test reactivation** – Several skipped tests (#5277, #5275) fail only due to outdated snapshot data, not logic errors. This slows down test maintenance.
- **Command search test coverage gaps** – The long‑suggestion expand/collapse test (#5280) was skipped for an unknown period; behavior now works but coverage lags.
- **Configuration documentation drift** – Issues like `docs: fix stale defaults, CLI syntax, and tool naming drift` (in v0.18.2) suggest users were confused by outdated docs. The project is actively addressing this.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-06-18

> **Project Note:** The tool is now named **CodeWhale** (on `github.com/Hmbown/CodeWhale`). The digest refers to it as CodeWhale (formerly DeepSeek TUI).

---

## 1. Today’s Highlights

No releases landed today, but the project saw intense bug-fixing and feature work. Key developments include: the **EPIC staged command-boundary refactor** (tracking #2791) moving forward with a replay PR on the `hunter` branch; two critical bugs—**agent self‑looping** (#3275) and **Plan/Agent mode toggle inconsistency** (#3279)—attracted detailed reports and immediate PR fixes; and a batch of performance and reliability PRs addressed snapshot‑disabled‑yet‑running, config‑comment erasure, and slow thinking‑stream re‑renders.

---

## 2. Releases

No new releases in the last 24 hours.

---

## 3. Hot Issues

All 10 issues updated in the last 24h (5 open, 2 closed, 3 closed with comments). The most noteworthy:

- [#2870 – EPIC: staged command-boundary refactor for #2791](https://github.com/Hmbown/CodeWhale/issues/2870)  
  **Open, documentation/cleanup/TUI.** Tracks the smaller mergeable chunks of a large refactor. Reference PR #2851. Community has 5 comments; this is the organizing issue for v0.9.0 command rework.

- [#3275 – CodeWhale overly involved in self‑questioning and self‑answering](https://github.com/Hmbown/CodeWhale/issues/3275)  
  **Open, bug/question.** Users report the agent loops without waiting for confirmation—a regression from #3061. High importance as it degrades user control.

- [#2917 – Cargo distribution: failed to spawn `codewhale` after rename from deepseek-tui](https://github.com/Hmbown/CodeWhale/issues/2917)  
  **Closed, bug.** Rename caused PATH issues. Resolution likely involved migration docs; 4 comments.

- [#3279 – Plan/Agent Mode Toggle Inconsistency & Tool Permission Chaos](https://github.com/Hmbown/CodeWhale/issues/3279)  
  **Open, bug/UX.** Switching from Plan to Agent mode fails to restore permissions; AI then auto‑executes plans. Detailed reproduction steps (Chinese). Three comments.

- [#1481 – Support OpenCode Go/Zen (DeepSeek-V4 provider)](https://github.com/Hmbown/CodeWhale/issues/1481)  
  **Open, enhancement.** Old request for cheap DeepSeek-V4 access. Still only 2 comments and 1 👍, but remains open – potential community desire.

- [#3289 – v0.8.61 UI freezes after auto‑spawn of several agents](https://github.com/Hmbown/CodeWhale/issues/3289)  
  **Open, bug.** Plan mode triggers multiple agents and causes freeze. Only 2 comments, likely resource contention.

- [#3281 – Kimi/Moonshot parameters fix incomplete: type:object still missing for $ref/allOf schemas](https://github.com/Hmbown/CodeWhale/issues/3281)  
  **Open, bug.** The v0.8.61 fix for #3265 only covered narrow schema patterns. Community identified remaining gaps.

- [#1530 – Feature: Support session continuity in non‑interactive mode (exec --resume / --session-id)](https://github.com/Hmbown/CodeWhale/issues/1530)  
  **Closed, enhancement.** Request for multi‑turn CLI workflows. Closed with resolution; likely merged or superseded.

- [#3292 – pre_tool_snapshot did not respect config snapshots.enabled=false](https://github.com/Hmbown/CodeWhale/issues/3292)  
  **Open, bug.** User disabled snapshots but tool‑level snapshots still copied whole repo to `~/.deepseek/snapshots`, consuming GBs. Clear config violation.

- [#3282 – config.toml file content improvement: comment lines erased by TUI edits](https://github.com/Hmbown/CodeWhale/issues/3282)  
  **Open, enhancement.** When modifying config via TUI, all commented‑out lines are lost. User wants preservation for notes and temporary disabling.

---

## 4. Key PR Progress

10 important pull requests from the 26 updated in the last 24h:

- [#3295 – feat(tui): honor ask permission rules at runtime](https://github.com/Hmbown/CodeWhale/pull/3295)  
  Wires `permissions.toml` into TUI approval path. Loads as `ExecPolicyEngine` and applies to `exec_shell` tool calls. (greyfreedom)

- [#3296 – [codex] Gate cross-tool skill discovery](https://github.com/Hmbown/CodeWhale/pull/3296)  
  Adds `[skills].scan_codewhale_only` config to scope skill scanning to CodeWhale roots, preserving broad compatibility. (nightt5879)

- [#3277 – feat: implement Workrooms Phase 1](https://github.com/Hmbown/CodeWhale/pull/3277)  
  Foundation for v0.9.0 Workroom abstraction: RFC, data model, endpoints, and tooling for threaded agent conversations. (idling11)

- [#3280 – fix(auto): allow heuristic‑only auto routing when flash router unavailable](https://github.com/Hmbown/CodeWhale/pull/3280)  
  Fixes `--model auto` failing without DeepSeek API key by falling back to heuristics. (hongchen1993)

- [#3274 – feat(release): build static Linux x64 binaries with musl](https://github.com/Hmbown/CodeWhale/pull/3274)  
  Switches GH Actions release from dynamic glibc to static musl for better portability. (wavezhang)

- [#3283 – Fix: Plan/Agent Mode Toggle — approval_mode restore + auto‑execution guard](https://github.com/Hmbown/CodeWhale/pull/3283)  
  Two fixes for #3279: restores `approval_mode` and adds guard against auto-execution after mode switch. (idling11)

- [#3284 – perf(tui): debounce thinking‑stream re‑renders](https://github.com/Hmbown/CodeWhale/pull/3284)  
  Fixes #1620 — slow reasoning streaming. Debounces `bump_active_cell_revision()` to avoid per‑delta invalidation. (LeoLin990405)

- [#3285 – fix(tui): persist session before stall/cancel recovery so --continue keeps history](https://github.com/Hmbown/CodeWhale/pull/3285)  
  Partial fix for #2739: saves turn before stall/cancel to prevent data loss when using `--continue`. (LeoLin990405)

- [#3290 – fix(prompts): add scope_discipline rules to prevent self‑questioning agent loops](https://github.com/Hmbown/CodeWhale/pull/3290)  
  Targets #3273/#3275 by adding constitution rules to stop self‑sustaining loops. (yekern)

- [#3291 – Fix/preserve comments in config files](https://github.com/Hmbown/CodeWhale/pull/3291)  
  All config write paths now merge comments via `toml_edit`, addressing #3282. (zlh124)

---

## 5. Feature Request Trends

From the issues, the most‑requested feature directions are:

- **Provider expansion** – support for cheaper/alternative backends like OpenCode Go/Zen (DeepSeek V4) (#1481).
- **Better configuration UX** – preserving comments in `config.toml` through TUI edits (#3282), and respecting config toggles like `snapshots.enabled` (#3292).
- **Non‑interactive session continuity** – `exec --resume` / `--session-id` for multi‑turn CLI workflows (#1530, already closed but indicative of ongoing need).
- **Mode & permission consistency** – users want clear, reliable Plan/Agent mode transitions with proper tool permission scoping (#3279).

---

## 6. Developer Pain Points

Recurring frustrations visible in today’s issues:

- **Agent over‑engineering** – the model frequently exceeds requested scope, entering self‑questioning/answering loops (#3275, #3290). Users feel a loss of control.
- **Config not honoured** – settings like `snapshots.enabled = false` are silently ignored by certain code paths (#3292). Also, auto‑erasure of comments (#3282) adds friction.
- **UI freezes and stalls** – spawn of multiple agents can freeze the TUI (#3289); reasoning streaming is painfully slow (#3284).
- **API schema issues** – provider‑specific parameter shapes (Kimi/Moonshot) continue to cause 400 errors after incomplete fixes (#3281).
- **Renaming migration pains** – the rename from `deepseek-tui` to `codewhale` caused broken PATH and installation errors (#2917), now resolved but a cautionary tale.

---

*Digest generated from GitHub data for 2026‑06‑18. All links point to `github.com/Hmbown/CodeWhale`.*

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*