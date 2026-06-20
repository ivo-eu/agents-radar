# AI CLI Tools Community Digest 2026-06-20

> Generated: 2026-06-20 10:17 UTC | Tools covered: 9

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

# Cross-Tool Comparison Report: AI CLI Developer Tools Ecosystem

**Date:** 2026-06-20  
**Analyst:** Senior Technical Analyst, AI Developer Tools Ecosystem

---

## 1. Ecosystem Overview

The AI CLI tools landscape is experiencing a period of rapid maturation punctuated by reliability growing pains. Across all nine major tools, the dominant themes are **agent reliability** (stalls, false success reporting, unbounded loops), **platform consistency** (Windows, Linux, and macOS regressions), and **user safety/control** (plan-mode violations, ignored rules, runaway API consumption). No tool is immune to these challenges, though community engagement remains high, with thousands of issues and PRs updated daily. The ecosystem is clearly transitioning from novelty to production-readiness, with developers demanding deterministic behavior, robust error recovery, and transparent cost controls.

---

## 2. Activity Comparison

| Tool | Hot Issues (today) | Key PRs (today) | Release Today |
|---|---|---|---|
| Claude Code | 10 (50 touched total) | 1 | None |
| OpenAI Codex | 10 | 10 | 3 alpha builds (v0.142.0-a.5/.6/.7) |
| Gemini CLI | 10 | 10 | None |
| GitHub Copilot CLI | 10 (17 touched total) | 2 | v1.0.64-1 |
| Kimi Code CLI | 1 | 1 | None |
| OpenCode | 10 (41 touched total) | 10 (50 touched total) | None |
| Pi | 10 | 5 | None |
| Qwen Code | 10 | 10 | v0.18.4 + v0.18.4-preview.0 |
| DeepSeek TUI | 10 (30 open issues) | 10 (27 touched total) | None |

**Key observations:**
- **OpenCode** and **DeepSeek TUI** show the highest PR throughput (50 and 27 touched respectively), indicating fast iteration cycles.
- **OpenAI Codex** shipped three alpha releases with zero changelog—a transparency concern for an enterprise-facing tool.
- **Kimi Code CLI** shows minimal activity, suggesting either low adoption or a slower development pace.
- **GitHub Copilot CLI** is the only tool with a stable patch release today, reflecting a more conservative release strategy.

---

## 3. Shared Feature Directions

The following requirements appear across **three or more** tool communities:

| Requirement | Tools Expressing Need | Specifics |
|---|---|---|
| **Project-scoped/workspace rules & hooks** | Claude Code, OpenAI Codex, Gemini CLI, GitHub Copilot CLI, OpenCode | Hooks compatibility (`PreToolUse`, `PostToolUse`, `SessionStart`), per-repo plugin config, glob-based rule targeting |
| **MCP/sandbox reliability & auth** | Claude Code, OpenAI Codex, Gemini CLI, GitHub Copilot CLI, OpenCode, Qwen Code | HTTP MCP with API-key auth, missing `sandboxPolicy` fields, OAuth refresh, 400+ tool limits, case-sensitive URL handling |
| **Agent safety budgets & guardrails** | Claude Code, Gemini CLI, OpenCode, DeepSeek TUI, Pi | Token budget governors, agentic loop limits, sub-agent recursion depth, destructive command prevention |
| **Plan-mode enforcement** | Claude Code, Qwen Code, OpenCode, DeepSeek TUI, Gemini CLI | Agents violating read-only contracts, self-generating approvals, auto-entering plan mode without consent |
| **Config discoverability & TUI editing** | DeepSeek TUI, OpenCode, Pi, Qwen Code | Runtime knob modification without editing `config.toml`, inline help for documented keys |
| **Voice/visual multimodal input** | Qwen Code (voice), Pi (OSC 9998/9999), DeepSeek TUI (GUI app), Gemini CLI (drag-and-drop images) | Expanding beyond terminal-only interaction |

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | Copilot CLI | Qwen Code | DeepSeek TUI | OpenCode | Kimi Code | Pi |
|---|---|---|---|---|---|---|---|---|---|
| **Core identity** | Hooks/rules-first CLI | Desktop + TUI hybrid | Agent systems + skills | Git-centric CLI | Chinese ecosystem focus | TUI with GUI bridge | Feature-rich parity seeker | Minimalist CLI | Multi-provider terminal |
| **Target user** | Power devs, CI/CD | Enterprise teams | Google ecosystem | GitHub-native devs | Chinese developers | DeepSeek users | Claude Code migrants | Simple use cases | Provider-agnostic devs |
| **Key differentiator** | `CLAUDE.md` hard gates, MCP server ecosystem | Desktop app with auto-update, sandboxed execution | Sub-agent delegation, skills framework, Auto Memory | Git worktrees, fleet mode, plugin hooks | Case-insensitivity sweep, Plan opt-in patch | TUI config editing, token budget governor | Highest community PR activity, scheduling feature | Simplicity, proxy support | Stable streaming, SQLite sessions |
| **Biggest weakness** | API reliability regression (#69358) | Auth lockout (#25749), chat loss on update | Agent hangs (#21409), false success (#22323) | Windows MCP failures, silent hang | Case-sensitivity bugs pervasive | TUI freezes on Windows, "Turn stalled" | Concurrency crashes, large-file hangs | Windows packaging, proxy gaps | Scroll-jump, CJK encoding |
| **Release velocity** | Moderate, stable | High alpha churn | Moderate, PR-driven | Conservative, stable | Balanced, systematic patches | Fast iteration, release trains | Very high | Low | Moderate |

---

## 5. Community Momentum & Maturity

### High Momentum (rapid iteration, strong community engagement)
- **OpenCode**: 50 PRs touched today—highest in ecosystem. Aggressively pursuing Claude Code feature parity. Safety concerns are rising but community response is equally active.
- **DeepSeek TUI**: 27 PRs today, plus a Tauri GUI app contribution. Maintainer-led refactoring (monolith split, provider registry) signals architectural maturity. Reliability regressions are the main drag.
- **Qwen Code**: Systematic bug sweep (15+ case-sensitivity fixes) and two releases today. Strong organizational discipline, though CI fragility is a concern.

### Established & Stable (active but controlled development)
- **GitHub Copilot CLI**: Conservative patch release. Features are incremental (worktree flags, tab completion). Community focuses on long-standing bugs (Windows, hooks).
- **OpenAI Codex**: High alpha churn (3 builds/day) but no changelog. Auth and sandbox issues dominate. Enterprise users may be frustrated by instability despite rapid iteration.
- **Claude Code**: High community attention (50 issues touched) but only 1 PR. Much of the activity is triage—closing duplicates, consolidating reports. The project appears to be in a stabilization phase.

### Slower or Niche
- **Pi**: 5 PRs today, moderate activity. Scroll-jump bug (#5825) dominates conversation. Focus on provider flexibility and streaming UX.
- **Kimi Code CLI**: 1 issue, 1 PR. Extremely low activity. Either low adoption or resource-constrained development.

---

## 6. Trend Signals

Five industry-wide signals emerging from this data:

### 1. **Agent autonomy vs. safety is the defining tension**
Across every tool, developers report agents violating explicit boundaries: executing write commands in plan mode, self-generating approvals, ignoring `CLAUDE.md` hard gates, burning API credits in unbounded loops. The market is demanding *reliable obedience* over raw capability. Tools that solve this (budget governors, opt-in plan modes, enforcement hooks) will win trust.

### 2. **MCP is the new battleground, but it's fragile**
Model Context Protocol support has become table stakes—but every tool struggles with auth, schema validation, tool count limits, and sandbox policies. The ecosystem is still converging on MCP standards while users complain about breaking changes between desktop updates. Expect consolidation around a subset of MCP servers and better error messaging.

### 3. **Desktop auto-updates are creating trust deficits**
Claude Code, OpenAI Codex, and GitHub Copilot CLI all have reports of auto-updates breaking workflows (chat loss, schema errors, sandbox regressions, process leaks). Users want opt-in updates, rollback capabilities, and visible changelogs. The "ship fast" approach is eroding confidence in production use.

### 4. **Cross-platform parity remains unsolved**
Windows and Linux users consistently report more bugs than macOS counterparts: glibc mismatches (DeepSeek TUI), pty handle leaks (Claude Code), PowerShell invocation on Mac remotes (Codex), Windows MCP fetch failures (Copilot CLI), Korean path encoding (Pi), and tar/zip extraction mismatches (Kimi Code). Most tools are built by macOS-using developers, and it shows.

### 5. **The CLI is evolving beyond the terminal**
Voice input (Qwen Code), image drag-and-drop (Gemini CLI, OpenCode), GUI apps (DeepSeek TUI with Tauri), and web adapters (Pi with OSC 9998/9999) all signal that pure-terminal interaction is expanding. The terminal is becoming a *bridge* rather than a *sink*—users want multimodal, persistent, and shareable sessions.

---

**Bottom line for technical decision-makers:** No single tool is production-ready across all dimensions. Choose based on your risk tolerance: **GitHub Copilot CLI** offers stability and incremental improvement for GitHub-native teams; **OpenCode** and **DeepSeek TUI** provide the fastest innovation but with reliability caveats; **Claude Code** and **OpenAI Codex** offer the richest feature sets but currently struggle with API reliability and update stability. Prioritize tools that demonstrate a clear commitment to **agent safety controls, cross-platform testing, and transparent release practices**.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
**Data Snapshot:** 2026-06-20 | **Source:** github.com/anthropics/skills

---

## 1. Top Skills Ranking

Based on community discussion volume (PR comments, cross-references in Issues), the following Skills are generating the most attention:

### 1.1 `document-typography` (PR #514)
- **Functionality:** Prevents orphan word wrap, widow paragraphs, and numbering misalignment in AI-generated documents — addressing endemic typographic quality issues in Claude output.
- **Discussion Highlights:** The PR's framing resonates widely — "these issues affect every document Claude generates" — suggesting latent demand for output polish that transcends any single document format.
- **Status:** Open

### 1.2 `odt` — OpenDocument Text Creation (PR #486)
- **Functionality:** Creates, fills, reads, and converts OpenDocument Format files (.odt, .ods), triggered by mentions of ODF, LibreOffice, or open-source document standards.
- **Discussion Highlights:** Reflects enterprise demand for ODF compatibility alongside existing DOCX capabilities. The ISO-standard format angle attracts LibreOffice/OpenOffice users who cannot adopt Microsoft-centric skills.
- **Status:** Open

### 1.3 `frontend-design` improvement (PR #210)
- **Functionality:** Revises the existing frontend-design skill for clarity, actionability, and internal coherence — ensuring every instruction is executable within a single conversation.
- **Discussion Highlights:** Represents a meta-concern: skill quality matters as much as skill quantity. Community debates center on *how* skills should instruct Claude, not just *what* they teach.
- **Status:** Open

### 1.4 `skill-quality-analyzer` + `skill-security-analyzer` (PR #83)
- **Functionality:** Two meta-skills for evaluating other skills across five dimensions (structure, documentation, security) — effectively a quality-assurance framework for the skills ecosystem itself.
- **Discussion Highlights:** Early proposal (November 2025) that remains heavily referenced. Signals demand for ecosystem-level validation tooling as the skills catalog grows.
- **Status:** Open

### 1.5 PDF/DOCX/Skill-Creator fixes (PRs #538, #539, #541)
- **Functionality:** Collectively fix case-sensitive file references (PDF skill), YAML unquoted-description parsing (skill-creator), and tracked-change ID collisions (DOCX skill).
- **Discussion Highlights:** These three PRs from the same author (Lubrsy706) form a de facto "skill infrastructure reliability" theme. The DOCX fix (#541) is particularly notable for diagnosing a subtle OOXML corruption root cause.
- **Status:** Open

### 1.6 `SAP-RPT-1-OSS predictor` (PR #181)
- **Functionality:** Integrates SAP's open-source tabular foundation model for predictive analytics on enterprise business data — a specialized enterprise AI skill.
- **Discussion Highlights:** Indicates community interest in marrying Claude skills with specialized enterprise models rather than relying solely on Claude's general capabilities.
- **Status:** Open

### 1.7 `aurelion-kernel/advisor/agent/memory` suite (PR #444)
- **Functionality:** Four-skills cognitive framework for structured thinking templates and professional knowledge management — a "meta-framework" approach to skill design.
- **Discussion Highlights:** The most ambitious single PR by scope (4 skills). Discussion centers on whether monolithic frameworks or granular single-purpose skills serve users better.
- **Status:** Open

### 1.8 `shodh-memory` — persistent context (PR #154)
- **Functionality:** Enables cross-conversation memory persistence for AI agents, teaching Claude when to surface relevant memories and how to structure them.
- **Discussion Highlights:** Addresses the fundamental limitation of session-based Claude interactions. Related to community demand for "agent-grade" persistence beyond ephemeral conversations.
- **Status:** Open

---

## 2. Community Demand Trends

From the top community Issues (sorted by engagement), five clear demand directions emerge:

### 2.1 Infrastructure & Reliability (Highest Urgency)
**Issue #556** (12 comments), **Issue #1169** (3 comments), **Issue #1061** (3 comments) — The `run_eval.py` zero-trigger-rate bug is the most actively discussed technical problem in the repository. Multiple independent reproductions confirm the skill-description optimization loop reports `recall=0%` on every iteration. **This blocks the entire skill-quality improvement pipeline.**

### 2.2 Security & Governance
**Issue #492** (7 comments) — Community skills distributed under the `anthropic/` namespace create a trust boundary vulnerability. Users may grant elevated permissions to skills they mistakenly believe are official. **This is a trust-layer design gap in the distribution model.**

**Issue #1175** (4 comments) — Concerns about security and context window implications when handling SharePoint Online documents via Agent Skills, questioning whether SKILL.md files should embed access control logic.

### 2.3 Skill Sharing & Distribution
**Issue #228** (14 comments — the most-commented issue) — Enabling org-wide skill sharing in Claude.ai without manual file transfers. Currently, users must download `.skill` files and email/Slack them. **This is the single most-requested feature** based on comment volume.

**Issue #189** (6 comments) — Plugin conflicts: `document-skills` and `example-skills` install identical content, causing duplicate skills in Claude Code's context window. Points to broader plugin-deduplication needs.

### 2.4 Skill Quality Standards
**Issue #202** (8 comments) — The `skill-creator` skill reads like developer documentation rather than an operational skill. Community expects skills to be *instructional* for Claude, not *explanatory* for humans. **This reflects maturing understanding of skill design philosophy.**

### 2.5 Platform & Ecosystem Expansion
**Issue #16** (4 comments) — Exposing Skills as MCPs (Model Context Protocol) for broader tool interoperability.

**Issue #29** (4 comments) — Making skills work with AWS Bedrock, indicating demand beyond the Claude desktop/CLI ecosystem.

---

## 3. High-Potential Pending Skills

These active-comment PRs address the community's most pressing pain points and are likely to land soon:

### 3.1 Critical Infrastructure Fixes
- **PR #361** — Detect unquoted YAML special characters in descriptions (updated 2026-06-10). Directly addresses YAML parsing failures that silently break skills.
- **PR #362** — Fix UTF-8 panic on multi-byte characters (updated 2026-06-10). Prevents Rust CLI panics when skills contain non-ASCII text.
- **PR #1050** — Fix Windows subprocess + encoding bugs (updated 2026-05-24). Companion to Issues #1061 and #1665.
- **PR #1099** — Fix `run_eval.py` crash on Windows (updated 2026-05-24). Directly addresses the `[WinError 10038]` subprocess pipe failure.

### 3.2 High-Value New Skills
- **PR #444** — AURELION skill suite (4 skills: kernel, advisor, agent, memory). Updated 2026-05-06 with continued discussion. A "big bet" on structured cognitive frameworks.
- **PR #723** — `testing-patterns` skill (testing trophy model, unit testing, React testing). Updated 2026-04-21. Addresses a gap: no comprehensive testing skill currently exists in the official catalog.

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is for reliability tooling and quality control — specifically fixing the broken skill-evaluation pipeline (`run_eval.py` zero-trigger-rate bug), establishing formal skill quality standards, and addressing distribution/namespace security — before expanding the catalog with new domain skills.**

This is a "build quality infrastructure first" moment: the ecosystem has accumulated enough skills (50+ in discussion, dozens in the repository) that the community now prioritizes *tooling to validate, share, and trust* skills over *creating new ones*.

---

*Analysis based on public data from github.com/anthropics/skills. Report generated 2026-06-20.*

---

# Claude Code Community Digest – 2026-06-20

*Based on data from [github.com/anthropics/claude-code](https://github.com/anthropics/claude-code)*

---

## Today's Highlights

No new releases landed in the last 24 hours. A single open bug (#69358) about persistent "No Response From API" on Linux is attracting significant community attention (46 👍, 14 comments). The only active pull request (#69698) fixes a marketplace installation problem caused by non‑root‑relative imports. Meanwhile, the issue tracker remains busy with 50 items touched in the last day—most are closed as duplicates, suggesting the team is actively triaging and consolidating reports.

---

## Releases

_None in the last 24 hours._

---

## Hot Issues (10 notable items)

1. **[#69358 – [BUG] No Response From API 2.1.181, 2.1.183 (constantly)](https://github.com/anthropics/claude-code/issues/69358)**  
   *Open, 14 comments, 46 👍*  
   A critical regression on Linux (API region). Users report that Claude Code versions 2.1.181 and 2.1.183 never receive a response from the API. The high upvote count and continued open status make this the most urgent issue today.

2. **[#68843 – `/remote` and `/effort` ultracode commands not functioning](https://github.com/anthropics/claude-code/issues/68843)**  
   *Closed as duplicate, 3 comments*  
   Both slash commands fail on macOS. The duplicate closure indicates a known root cause, but the issue is still affecting users.

3. **[#68927 / #68928 / #68940 / #68906 – Multiple "Server rate limiting" bugs](https://github.com/anthropics/claude-code/issues/68927)** (see sibling duplicates)  
   *Closed as duplicate, 3 comments each*  
   A cluster of reports (Windows, macOS) of API rate‑limit errors claiming "Server is temporarily limiting requests (not your usage limit)." Users are confused because they are not hitting personal usage caps.

4. **[#68961 – Excessive agentic loop iterations consuming API quota](https://github.com/anthropics/claude-code/issues/68961)**  
   *Closed as duplicate, 3 comments*  
   Claude Code running “dozens/hundreds of agents” unboundedly, burning through API credits. Highlights a lack of safeguards on agentic iteration limits.

5. **[#68982 – Cloud session: message silently dropped](https://github.com/anthropics/claude-code/issues/68982)**  
   *Closed as duplicate, 3 comments*  
   A message entered a "running" state with an infinite spinner; after refresh, the message was gone entirely. A concerning data‑loss scenario for cloud users.

6. **[#68806 – Claude Code can't use HTTP MCP servers with API key authentication](https://github.com/anthropics/claude-code/issues/68806)**  
   *Closed as duplicate, 2 comments*  
   MCP servers requiring API‑key headers fail to connect, blocking integration with many third‑party tools.

7. **[#68863 – CLAUDE.md 'Hard gate' rules are ignored](https://github.com/anthropics/claude-code/issues/68863)**  
   *Closed as duplicate, 2 comments*  
   Rules explicitly marked “Hard gate — no exceptions” in `CLAUDE.md` are repeatedly bypassed in the same session, undermining a core safety/guidance feature.

8. **[#68857 – Claude Opus 4.8 outputs raw `<invoke>` tags instead of tool_use blocks](https://github.com/anthropics/claude-code/issues/68857)**  
   *Closed as duplicate, 2 comments*  
   The model sometimes emits raw `<invoke>` text instead of structured `tool_use` blocks, causing calls to silently fail—especially problematic at turn ends.

9. **[#68849 – macOS: background session loses live file access (EPERM) mid‑session](https://github.com/anthropics/claude-code/issues/68849)**  
   *Closed as duplicate, 2 comments*  
   After a background worker re‑host, TCC permissions are lost, leading to `EPERM` errors. A consequence of the earlier bundle‑ID fix (#61314) not handling re‑host scenarios.

10. **[#68902 – macOS app leaks pty/ptmx handles, exhausting kernel limit](https://github.com/anthropics/claude-code/issues/68902)**  
    *Closed as duplicate, 2 comments*  
    Unclosed pty handles eventually saturate `kern.tty.ptmx_max`, breaking all new terminal sessions—including those of other apps.

---

## Key PR Progress

Only **one pull request** was updated in the last 24 hours:

- **[#69698 – `fix(hookify): use root-relative imports to fix marketplace install`](https://github.com/anthropics/claude-code/pull/69698)**  
  *Open, 0 comments*  
  **Author:** shrivs4  
  **Summary:** Resolves a module‑resolution issue that prevented Claude Code from being installed properly via certain marketplace channels. By switching to root‑relative import paths, the fix ensures that hook functionality works regardless of the consumer’s directory layout.  
  **Why it matters:** This addresses a friction point for new users who rely on marketplace distributions, and its prompt submission suggests it’s a high‑priority fix.

---

## Feature Request Trends

The following directions are emerging from the day’s issue activity (most are closed as duplicates, but the volume indicates strong community desire):

- **Enhanced MCP/remote support** – HTTP MCP servers with API‑key auth, remote session management, and downstream‑image delivery for mobile→desktop flows.
- **Smarter file navigation** – Clickable file links should cross git branches, worktrees, and out‑of‑session directories (instead of throwing dead‑end errors).
- **Multi‑path `/add-dir`** – Allow multiple directories per invocation, matching the variadic `--add-dir` launch flag behavior.
- **Image input in VSCode extension** – Paste/drag‑drop images into the chat input, already supported in the web version.
- **Customizable TUI theme** – Input‑box labels, border colors, and better visibility inside Warp terminal.
- **Model‑specific support** – e.g., Fable 5 model (some confusion around model availability).

---

## Developer Pain Points

Recurring frustrations observed from the day’s reports:

1. **API reliability / rate limiting** – Multiple users hit "Server is temporarily limiting requests" errors that are *not* personal usage caps. This appears to be an infrastructure‑side issue causing widespread confusion.
2. **Plan‑mode violations** – Claude executes write commands (Bash, file edit) while still in plan mode, contradicting the contract of read‑only exploration.
3. **Context rules ignored** – `CLAUDE.md` “hard gate” rules are loaded but not enforced; the model repeatedly disregards explicit instructions.
4. **Resource exhaustion** – Agent loops burning API credits without bounds; pty handle leaks that crash the terminal system.
5. **Permission/UX friction** – “Deny” button placement leading to accidental approvals; tool approval prompts that never complete (stuck in “unsupervised mode”).
6. **Silent failures** – Messages dropped (cloud sessions), raw `<invoke>` tags not executed, auto‑update failures with no diagnostic output in `/doctor`.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest – 2026-06-20

## Today’s Highlights
Three new Rust alpha builds (v0.142.0-alpha.5/.6/.7) shipped within 24 hours, but no release notes accompany them. Meanwhile, a critical auth recovery bug (#25749) continues to dominate discussion with 55 comments, and a wave of sandbox‑related MCP failures surfaced after the desktop app auto‑updated to version 26.616. The community is also pressing for workspace‑scoped chat history and better reasoning‑depth controls.

---

## Releases
Three **rust‑v0.142.0** alpha releases landed today (alpha.5, alpha.6, alpha.7). No changelog or migration notes were published.  
[openai/codex Releases](https://github.com/openai/codex/releases)

---

## Hot Issues (10)

1. **[#25749] Codex requires verification of an inaccessible legacy phone number and provides no phone number replacement or recovery path**  
   *Tags: bug, auth, app | Comments: 55 | 👍 33*  
   Users authenticated via Google OAuth/MFA cannot bypass a legacy phone‑number gate, effectively locking them out of Codex. Community outrage is growing.  
   [openai/codex Issue #25749](https://github.com/openai/codex/issues/25749)

2. **[#20741] Codex Desktop project chat histories disappeared after recent update**  
   *Tags: bug, app, session | Comments: 45 | 👍 14*  
   macOS users report total loss of chat threads after updating to build 26.429. No recovery path exists for local session files.  
   [openai/codex Issue #20741](https://github.com/openai/codex/issues/20741)

3. **[#12299] “You've hit your usage limit.” Despite 10% Rate Limit Usage Remaining**  
   *Tags: bug, extension, rate‑limits | Comments: 22 | 👍 1*  
   VS Code extension incorrectly blocks users even when the backend reports low quota. Likely a client/server sync bug.  
   [openai/codex Issue #12299](https://github.com/openai/codex/issues/12299)

4. **[#28978][Bug] Desktop app 26.616: new conversations fail with “Invalid request: missing field `inputSchema`” (CLI with same config works)**  
   *Tags: bug, mcp, app, app‑server | Comments: 13 | 👍 16*  
   Post‑update, all new threads break due to a schema validation error. CLI unaffected, pointing to a desktop‑only regression.  
   [openai/codex Issue #28978](https://github.com/openai/codex/issues/28978)

5. **[#24389] `multi_agent_v1.close_agent` can hang for hours when closing an unresponsive subagent**  
   *Tags: bug, mcp, tool‑calls, app, subagent | Comments: 8*  
   Parent threads freeze while trying to terminate a stuck subagent. Users report multi‑hour stalls.  
   [openai/codex Issue #24389](https://github.com/openai/codex/issues/24389)

6. **[#29000] Codex CLI 0.141.0 crashes with SIGTRAP (“trace trap”) on Intel macOS**  
   *Tags: bug, CLI | Comments: 6 | 👍 6*  
   CLI immediately crashes on x86_64 Macs. No workaround known yet.  
   [openai/codex Issue #29000](https://github.com/openai/codex/issues/29000)

7. **[#29189] Codex Desktop 26.616.41845 `node_repl` fails: `codex/sandbox‑state‑meta` missing `sandboxPolicy`**  
   *Tags: bug, mcp, sandbox, app, browser | Comments: 3 | 👍 14*  
   The bundled Node REPL tool cannot execute JavaScript because the sandbox meta field is absent. Desktop‑wide impact.  
   [openai/codex Issue #29189](https://github.com/openai/codex/issues/29189)

8. **[#26779] Codex mobile via ChatGPT mobile app remote Mac connection tries to run `powershell.exe`**  
   *Tags: bug, windows‑os, app, app‑server, remote | Comments: 7 | 👍 1*  
   When using a Mac as a remote host, Codex mobile incorrectly invokes Windows‑only PowerShell. Platform detection bug.  
   [openai/codex Issue #26779](https://github.com/openai/codex/issues/26779)

9. **[#25744] Codex for macOS accumulates Computer Use / MCP helper processes and unreaped zombie children, causing HID lag and WindowServer/TCC stalls**  
   *Tags: bug, mcp, app, app‑server, computer‑use, performance | Comments: 5*  
   Long‑running sessions leak processes, degrading system responsiveness. macOS‑specific, but serious for power users.  
   [openai/codex Issue #25744](https://github.com/openai/codex/issues/25744)

10. **[#29195] Codex Usage going up with no messages**  
    *Tags: bug, rate‑limits, app | Comments: 2*  
    Simply opening the app consumes quota. Users report 40% drops without any queries. Likely a polling/keepalive issue.  
    [openai/codex Issue #29195](https://github.com/openai/codex/issues/29195)

---

## Key PR Progress (10)

1. **[#28232] Add workspace headline statusline item** (OPEN)  
   Adds a TUI status‑line element that shows the active Enterprise workspace headline. Polls every 10s for admin changes.  
   [openai/codex PR #28232](https://github.com/openai/codex/pull/28232)

2. **[#29035] Optimize filesystem thread listing** (OPEN)  
   Improves performance of `thread/list` fallback by skipping rollout summaries for threads that are quickly rejected by metadata (source, model_provider, cwd).  
   [openai/codex PR #29035](https://github.com/openai/codex/pull/29035)

3. **[#29188] ci: enforce the hermetic Windows Bazel toolchain** (OPEN)  
   Removes dead MSVC platforms and sanitises the Windows Bazel environment to prevent non‑hermetic builds.  
   [openai/codex PR #29188](https://github.com/openai/codex/pull/29188)

4. **[#26725] Refine Guardian data exfiltration policy** (CLOSED)  
   Update policy wording for Guardian security reviews. No testing run.  
   [openai/codex PR #26725](https://github.com/openai/codex/pull/26725)

5. **[#26717] Stop guardian reviews when parent turns are interrupted** (CLOSED)  
   Fixes a bug where Guardian approval reviews continued running after their parent turn was aborted, causing stale assessments.  
   [openai/codex PR #26717](https://github.com/openai/codex/pull/26717)

6. **[#26618] fix(tui): avoid duplicated streamed markdown lines** (CLOSED)  
   Prevents markdown list continuations from appearing twice in scrollback by tracking mutable tail boundaries.  
   [openai/codex PR #26618](https://github.com/openai/codex/pull/26618)

7. **[#26671] Allow `/bin/ps` in macOS seatbelt policy** (CLOSED)  
   Permits execution of `/bin/ps` inside the sandbox so Homebrew’s `brew shellenv` works. Adds policy unit test.  
   [openai/codex PR #26671](https://github.com/openai/codex/pull/26671)

8. **[#26663] fix(guardian): prevent approval review stack overflows** (CLOSED)  
   Fixes a stack overflow crash that occurred when Guardian opened a permission request in a multi‑directory workflow.  
   [openai/codex PR #26663](https://github.com/openai/codex/pull/26663)

9. **[#28806] optimize resume and fork history** (OPEN)  
   Uses checkpoint‑backed `thread/resume` and copy‑on‑write `thread/fork` optimisations to reduce cold‑start work.  
   [openai/codex PR #28806](https://github.com/openai/codex/pull/28806)

10. **[#29181] make image artifact directory configurable** (OPEN)  
    Allows users to set a custom path for generated images via `config.toml`, defaulting to `$CODEX_HOME/generated_images`.  
    [openai/codex PR #29181](https://github.com/openai/codex/pull/29181)

---

## Feature Request Trends
- **Workspace‑scoped chat history** – Users consistently ask for chat sessions tied to a VS Code project/workspace (#25319, +34👍).  
- **Dedicated reasoning‑depth controls** – Quick hotkeys and a non‑semantic depth meter for switching models mid‑session (#14356, +13👍).  
- **Tab management** – Running tabs should be listed like browser tabs (#18778).  
- **Disable automatic updates** – Users want control over when desktop updates are applied (#18546).  
- **Sandbox/policy self‑test** – A “Test Browser Tool” button in Settings to validate integration (#29198).  
- **Pet animation improvements** – Pets should stay animated while a task runs, not fall back to idle (#23272, #29196).

---

## Developer Pain Points
- **Auth lockouts with no recovery** – Legacy phone verification blocks even authenticated users (#25749).  
- **Silent update breakage** – Desktop auto‑updates frequently cause chat loss (#20741), schema errors (#28978), and sandbox regressions (#29189).  
- **Session/resource leaks** – Zombie processes on macOS degrade performance (#25744); subagent hangs freeze threads (#24389).  
- **Rate‑limit false positives** – Users are throttled even when quota remains (#12299), or usage mysteriously increases with no activity (#29195).  
- **Windows‑specific quirks** – PowerShell invocation on Mac remotes (#26779), invisible tray icons (#29037), and context‑window exhaustion unrecoverability (#28920).  
- **CLI crash on Intel macOS** – SIGTRAP makes the CLI unusable for x86_64 Macs (#29000).

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-06-20

## Today’s Highlights
No new releases in the last 24 hours. The community continues to report a cluster of P1 bugs around agent hangs and subagent false-success reporting, with heavy team activity on triage and retesting. Key PRs this week tackle critical CVEs, MCP OAuth refresh, Jupyter file corruption, and a new drag-and-drop image pasting feature. The maintainer-only label on most high-traffic issues suggests active internal investigation.

## Releases
None.

## Hot Issues

1. **Generalist agent hangs indefinitely**  
   Issue [#21409](https://github.com/google-gemini/gemini-cli/issues/21409) (P1, 8 👍, 7 comments)  
   The top community pain point: simple commands like folder creation cause the agent to freeze for up to an hour. Users report that disabling subagent delegation resolves the hang. Maintainers are retesting.

2. **Subagent recovery falsely reports success after MAX_TURNS**  
   Issue [#22323](https://github.com/google-gemini/gemini-cli/issues/22323) (P1, 2 👍, 6 comments)  
   `codebase_investigator` subagent hits turn limit but reports `status: "success"` and `Termination Reason: "GOAL"`, masking the interruption. Undermines trust in agent workflows.

3. **Shell command execution stuck on “Waiting input” after completion**  
   Issue [#25166](https://github.com/google-gemini/gemini-cli/issues/25166) (P1, 3 👍, 4 comments)  
   Repeatedly blocks trivial shell commands even when no user input is required. High impact on basic task automation.

4. **`google_web_search` loops indefinitely on no results**  
   Issue [#28037](https://github.com/google-gemini/gemini-cli/issues/28037) (P2, 5 comments)  
   Newly filed; triggers repeated queries without stopping condition. Reproducible with direct tool calls. No community upvotes yet but high relevance for agent reliability.

5. **Auto Memory logs secrets before redaction**  
   Issue [#26525](https://github.com/google-gemini/gemini-cli/issues/26525) (P2, 5 comments)  
   Extraction prompt asks the model to redact secrets, but content is already in model context before that step. Also logs skill configurations. Security concern.

6. **Auto Memory retries low-signal sessions indefinitely**  
   Issue [#26522](https://github.com/google-gemini/gemini-cli/issues/26522) (P2, 5 comments)  
   Sessions that the extraction agent decides to skip are never marked processed, leading to repeated rediscovery. Strains storage and performance.

7. **Browser subagent fails on Wayland**  
   Issue [#21983](https://github.com/google-gemini/gemini-cli/issues/21983) (P1, 1 👍, 4 comments)  
   Browser agent terminates with `GOAL` immediately on Wayland compositors. Linux users affected.

8. **Gemini does not use custom skills and sub-agents autonomously**  
   Issue [#21968](https://github.com/google-gemini/gemini-cli/issues/21968) (P2, 6 comments)  
   Anecdotal but widely echoed: even when skills are explicitly described, the model rarely invokes them unless directed. Undermines extensibility.

9. **400 error when tool count exceeds 128**  
   Issue [#24246](https://github.com/google-gemini/gemini-cli/issues/24246) (P2, 3 comments)  
   With many MCP/skills enabled, the API rejects requests. Suggests need for tool-scoping logic.

10. **Agent should discourage destructive behavior (git reset, --force)**  
    Issue [#22672](https://github.com/google-gemini/gemini-cli/issues/22672) (P2, 1 👍, 3 comments)  
    Users report the model executing risky commands like `git reset` or `--force` when safer alternatives exist. Call for safer default prompts.

## Key PR Progress

1. **Upgrade shell-quote to 1.8.4 (CVE-2026-9277)**  
   PR [#27856](https://github.com/google-gemini/gemini-cli/pull/27856) (size/s)  
   Fixes a critical vulnerability in `shell-quote`. Trivy scanner flagged it as likely exploitable. Security patch.

2. **Sniff MCP image MIME types**  
   PR [#27878](https://github.com/google-gemini/gemini-cli/pull/27878) (P1, size/l)  
   WebP images from Figma MCP were mislabeled as PNG, causing HTTP 400s. Implements local binary signature sniffing. Fixes #27731.

3. **Preserve dollar sequences in prompt template substitutions**  
   PR [#28055](https://github.com/google-gemini/gemini-cli/pull/28055) (area/agent, size/s)  
   Fixes corruption of `$` sequences (`$$`, `$'`, `$&`) in skill, sub-agent, or tool descriptions during template substitution.

4. **Add native drag-and-drop and Cmd+V clipboard image pasting**  
   PR [#27859](https://github.com/google-gemini/gemini-cli/pull/27859) (P3, size/m/l)  
   Brings visual multimodal input to the terminal. Long-requested feature; includes terminal emulator compatibility.

5. **Cap pending tool responses**  
   PR [#27870](https://github.com/google-gemini/gemini-cli/pull/27870) (P1, size/m)  
   Prevents a very large tool result from being sent as a pending `functionResponse`, which could crash or hang the agent. Fixes #27738.

6. **Fix Jupyter Notebook and JSON corruption in write_file**  
   PR [#28000](https://github.com/google-gemini/gemini-cli/pull/28000) (size/m)  
   Silent corruption of `.ipynb` and JSON files made them unparseable. Affects Colab/JupyterLab users.

7. **Refresh MCP OAuth with stored client ID**  
   PR [#27889](https://github.com/google-gemini/gemini-cli/pull/27889) (P1, size/m)  
   MCP OAuth refresh failed for auto-discovered servers without static `oauth.clientId`. Fix persists discovered client ID from token metadata.

8. **Defensive path resolution for @-reference files**  
   PR [#28053](https://github.com/google-gemini/gemini-cli/pull/28053) (size/xl)  
   Filesystem tools (`read_file`, `replace`, `write_file`) failed with “File not found” when the model passed paths prefixed with `@`. Also fixes macOS test failures.

9. **CI: Validate workflow_run origin against fork artifact poisoning**  
   PR [#27753](https://github.com/google-gemini/gemini-cli/pull/27753) (P1, size/s)  
   The chained E2E pipeline was vulnerable to fork PRs injecting attacker-controlled code into workflows with secrets. Adds origin validation.

10. **Handle single-line descriptions in SKILL.md frontmatter**  
    PR [#28042](https://github.com/google-gemini/gemini-cli/pull/28042) (P2, size/m)  
    Skills with no blank line after the `description` field became invisible in `/skills list`. Affects extensibility.

## Feature Request Trends

- **AST-aware code understanding** – Multiple issues ([#22745](https://github.com/google-gemini/gemini-cli/issues/22745), [#22746](https://github.com/google-gemini/gemini-cli/issues/22746), [#22747](https://github.com/google-gemini/gemini-cli/issues/22747)) explore using abstract syntax tree tools for file reads, search, and codebase mapping to improve token efficiency and navigational accuracy.
- **Agent self-awareness and safety** – Users want the CLI to know its own flags, hotkeys, and limitations ([#21432](https://github.com/google-gemini/gemini-cli/issues/21432)), and to avoid destructive commands like force resets or direct database modifications ([#22672](https://github.com/google-gemini/gemini-cli/issues/22672)).
- **Extensibility maturity** – The closed initiative [#8784](https://github.com/google-gemini/gemini-cli/issues/8784) (MCP support) is live, but follow-ups call for browser session takeover/recovery ([#22232](https://github.com/google-gemini/gemini-cli/issues/22232)) and better tool-scoping when >128 tools are available ([#24246](https://github.com/google-gemini/gemini-cli/issues/24246)).
- **Robust evaluation infrastructure** – Epic [#24353](https://github.com/google-gemini/gemini-cli/issues/24353) drives component-level evaluations, while [#23166](https://github.com/google-gemini/gemini-cli/issues/23166) aims to stabilize internal project evals.

## Developer Pain Points

The most disruptive recurring themes are **agent hangs and false success reporting** – especially the generalist agent freeze (8 👍) and subagent recovery that lies about completion. **Shell command stalling** and **web search loops** further erode trust in automated workflows. **Auto Memory** draws complaints for indefinite retries of low-signal sessions and for exposing secrets before redaction. **Browser agent failure on Wayland** and **tool limits (400 error with >128 tools)** block specific workflows. Smaller but frequent frustrations include **destructive command execution**, **terminal corruption after external editor use**, and **silent failures in skill discovery** when frontmatter is non-standard. Across the board, the community demands **deterministic, self-correcting agent behavior** and **safer default prompts**.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-06-20

## Today’s Highlights
A minor patch release (v1.0.64-1) introduces a `/branch` alias for `/fork` and an experimental `--worktree` flag to simplify multi-branch workflows. Community momentum centers on project-scoped plugins, platform compatibility fixes, and calls for better hook and context-window visibility. Several long-standing bugs around session resume, networking stalls, and plugin configuration are now closed.

## Releases
**v1.0.64-1** (released today)  
[Full changelog](https://github.com/github/copilot-cli/releases/tag/v1.0.64-1)
- **Added**
  - `/branch` alias for `/fork`, matching Claude Code’s command naming.
  - Experimental `--worktree [name]` (`-w`) flag (enable with `/experimental`) – creates or reuses a git worktree under `<repo>.worktrees/` and starts the session inside it.
  - Tab completion for `/agent n`.

## Hot Issues (10 of 17 updated in last 24h)
1. **#731** (CLOSED) – **[area:sessions] Incompatibility with Z shell and direnv**  
   Invalid session ID when using Z shell + direnv. High community impact (14 👍, 13 comments).  
   [Link](https://github.com/github/copilot-cli/issues/731)

2. **#1665** (CLOSED) – **[area:plugins, area:configuration] Support Copilot CLI Plugins Scoped to Project or Repository**  
   Users want repo/project-specific plugins instead of global install. 17 👍, strongly requested.  
   [Link](https://github.com/github/copilot-cli/issues/1665)

3. **#1901** (OPEN) – **[area:non-interactive, area:agents] fleet plan approval race condition**  
   Autopilot fleet mode may not activate immediately, causing a 50-minute delay.  
   [Link](https://github.com/github/copilot-cli/issues/1901)

4. **#3455** (OPEN) – **[area:platform-windows, area:networking, area:mcp] github-mcp-server fails with “fetch failed”**  
   Regression since v1.0.51 on Windows; no MCP server connection attempts logged.  
   [Link](https://github.com/github/copilot-cli/issues/3455)

5. **#3871** (OPEN) – **[area:plugins] No way to list installed hooks**  
   MCP has `copilot mcp list`, but hooks have no equivalent – missing discoverability.  
   [Link](https://github.com/github/copilot-cli/issues/3871)

6. **#3872** (CLOSED) – **[area:plugins, area:configuration] Mis-cased event key silently dropped**  
   Hook configs with `PreToolUse` instead of `preToolUse` are ignored with only a debug log – no user warning.  
   [Link](https://github.com/github/copilot-cli/issues/3872)

7. **#3371** (OPEN) – **[area:networking] CLI silently hangs on stalled HTTPS sockets**  
   No timeout, no log output during TCP send buffer backpressure to api.github.com.  
   [Link](https://github.com/github/copilot-cli/issues/3371)

8. **#3821** (CLOSED) – **[area:sessions, area:installation] `/update` from resumed session leaves conflicting flags**  
   Running `/update` after `-r` resume restarts with both `--session-id` and `--resume` set, breaking the session.  
   [Link](https://github.com/github/copilot-cli/issues/3821)

9. **#3874** (OPEN) – **[area:permissions, area:plugins] VS Code agent `preToolUse` hook denial does not work**  
   Hooks that deny commands via `PreToolUse` are ineffective in VS Code Copilot Chat.  
   [Link](https://github.com/github/copilot-cli/issues/3874)

10. **#3866** (OPEN) – **[area:theming-accessibility] Thinking/reasoning text unreadable on dark backgrounds**  
    Hardcoded dim foreground makes “Thinking…” text nearly invisible in dark terminals.  
    [Link](https://github.com/github/copilot-cli/issues/3866)

## Key PR Progress (2 updated in last 24h)
1. **#3873** (OPEN) – **Add initial console log for greeting**  
   Minor enhancement for startup logging.  
   [Link](https://github.com/github/copilot-cli/pull/3873)

2. **#2587** (CLOSED) – **Add automated issue classification with GitHub Agentic Workflows**  
   AI-powered workflow that automatically applies `area:` and `triage` labels on issue open/reopen.  
   [Link](https://github.com/github/copilot-cli/pull/2587)

*Note: Only two pull requests received updates in the last 24 hours.*

## Feature Request Trends
- **Project-scoped plugins** (#1665) – Hooks, agents, and MCP servers should be configurable per repository rather than globally.
- **Context window visibility** (#3867) – Users want a visible token counter and compaction notifications in chat sessions.
- **Improved /ask UI** (#3869) – The answer box is too cramped; users need a larger or scrollable display for code snippets.
- **LLM-invocable `cd` tool** (#3865) – A tool that updates the status bar directory when Copilot switches worktrees.
- **Worktree support** (v1.0.64-1) – The experimental `--worktree` flag aligns with developer desire for isolated environments.

## Developer Pain Points
- **Platform-specific regressions** – Windows (MCP server fetch failure #3455) and Alpine Linux (auto-update downloads wrong arch #3696) remain fragile.
- **Silent failures & poor diagnostics** – Mis-cased hook events (#3872), stalled HTTPS sockets (#3371), and race conditions in fleet mode (#1901) cause frustration with no user-facing feedback.
- **Session management bugs** – The `/update` + resume conflict (#3821) and the Z shell/direnv incompatibility (#731) disrupt continuous workflows.
- **Lack of discoverability** – No command to list installed hooks (#3871) forces users to manually inspect config files.
- **Accessibility gaps** – Hardcoded dim text (#3866) and lack of context window indicators (#3867) hurt usability in modern terminal setups.

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-06-20

## Today's Highlights

Activity was low today, with only one closed bug and one open pull request. The community flagged a critical packaging issue on Windows (Git Bash + VS Code extension), and a fix for system proxy handling was proposed—both addressing long-standing developer friction in corporate and cross-platform environments.

## Releases

No new releases in the last 24 hours.

## Hot Issues

Only one issue was updated in the past 24 hours. It is included below with context.

- **#2462** [Bug] Windows + Git Bash: VS Code extension fails to extract bundled CLI because `tar` cannot handle zip  
  *Status: Closed*  
  **Summary:** When using the VS Code extension on Windows 10 with Git Bash (MSYS2), the bundled `kimi-cli` binary cannot be extracted because `tar` is invoked on a `.zip` file instead of an archive it supports.  
  **Why it matters:** This blocks Windows developers who rely on Git Bash — a very common setup. The issue was closed without comment, which may leave users waiting for a fix.  
  **Reaction:** 0 👍, 1 comment.  
  [View Issue](https://github.com/MoonshotAI/kimi-cli/issues/2462)

## Key PR Progress

Only one PR was updated in the last 24 hours.

- **#2463** fix: respect system proxy settings in FetchURL  
  *Status: Open*  
  **Summary:** `aiohttp.ClientSession` does not automatically read `HTTP_PROXY`/`HTTPS_PROXY` environment variables. This PR injects them explicitly, fixing `Connection reset by peer` errors in proxied networks.  
  **Why it matters:** Essential for enterprise users behind corporate proxies. The fix is small but impactful.  
  **Reaction:** 0 👍, no comments.  
  [View PR](https://github.com/MoonshotAI/kimi-cli/pull/2463)

## Feature Request Trends

No new feature requests were filed or updated today. However, based on the broader issue tracker’s recent history, the community continues to ask for:

- Better Windows support, especially for terminal emulators and bundling.
- Native proxy and VPN configuration handling.
- Simplified CI/CD integration (e.g., pipe-friendly JSON output).

## Developer Pain Points

- **Windows packaging issues** remain a top frustration. The `tar` vs `zip` extraction mismatch (Issue #2462) is a recurring theme — developers running Git Bash, WSL, or Cygwin expect seamless extraction of CLI tools distributed in `.zip` archives.
- **Proxy configuration** is an unaddressed need for many users behind firewalls. PR #2463 directly targets this pain point.
- **Lack of release cadence transparency** — with no new releases today and a closed bug without explanation, some developers may feel uncertain about when fixes will land.

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-20

## Today's Highlights
A wave of bug reports around safety, permissions, and sub-agent behavior hit the tracker today, while the community also saw a surge of RFCs and PRs for hooks compatibility, MCP progress, and TUI enhancements. No new releases were published in the last 24 hours, but the project remains highly active with 41 issues and 50 pull requests updated.

## Releases
No new releases in the last 24 hours.

## Hot Issues
*(10 noteworthy issues updated today, sorted by relevance)*

1. **[#12472 – Native Claude Code hooks compatibility (PreToolUse, PostToolUse, Stop)](https://github.com/anomalyco/opencode/issues/12472)**  
   OpenCode already supports Claude Code rules and skills but lacks its hooks system. This long-standing request has 32 👍 and 14 comments, indicating strong demand for parity. A companion RFC (#33054) was also filed today.

2. **[#32172 – [CLOSED] Add GLM-5.2 model support for Z.AI provider](https://github.com/anomalyco/opencode/issues/32172)**  
   Closed quickly after filing, suggesting Z.AI’s latest reasoning model is already supported or trivial to add.

3. **[#15988 – “Retry Now” button to skip rate limit retry countdown](https://github.com/anomalyco/opencode/issues/15988)**  
   Users hit rate limits and want immediate retry rather than waiting for the countdown. 13 👍 reflect a common workflow friction.

4. **[#11232 – Native Scheduling for OpenCode](https://github.com/anomalyco/opencode/issues/11232)**  
   The top-voted feature request (15 👍) asks for `opencode schedule --cron` to replace OS-level cron/systemd. Multiple PRs reference this issue.

5. **[#28015 – Worker terminated when running multiple subagents – session switching broken](https://github.com/anomalyco/opencode/issues/28015)**  
   A concurrency crash that locks users out of all sessions. 9 comments but no one has reproduced a fix yet.

6. **[#4716 – Glob-based rules](https://github.com/anomalyco/opencode/issues/4716)**  
   17 👍 for applying `AGENTS.md` rules per file glob patterns. A long-running request that would make project rules more expressive.

7. **[#5409 – SessionStart hook for session lifecycle events](https://github.com/anomalyco/opencode/issues/5409)**  
   17 👍 for parity with Claude Code’s `SessionStart`. Similar lifecycle hooks are also requested for worktree events (#15680, #33061).

8. **[#16626 – Add `session.stopping` plugin hook to re-enter agent loop](https://github.com/anomalyco/opencode/issues/16626)**  
   Plugins currently can’t prevent the agent loop from stopping. This hook would allow conditional continuation.

9. **[#31916 – TUI hangs on “Preparing to write…” with large file content](https://github.com/anomalyco/opencode/issues/31916)**  
   Unbounded diff rendering causes freeze on 150+ line writes. Windows users are most affected.

10. **[#33084, #33080, #33079, #33077, #33076, #33074, #33072, #33071 – Bug avalanche by @LifetimeVip](https://github.com/anomalyco/opencode/issues?q=author%3ALifetimeVip+created%3A2026-06-20)**  
    A single user filed 8+ detailed bugs today covering sub-agent language mismatch, missing read-before-act precondition, unenforceable AGENTS.md rules, destructive bash commands, no pre-flight safety checks, Cloudflare blocking from User-Agent, permission cascade bugs, and orphaned Windows processes. This signals deep frustration with core safety.

## Key PR Progress
*(10 notable pull requests updated today)*

1. **[#33082 – RFC: Computer Use for opencode](https://github.com/anomalyco/opencode/pull/33082)**  
   A design RFC seeking alignment before implementation. Refs multiple previous issues. No code—pure design review.

2. **[#33083 – Add desktop config file](https://github.com/anomalyco/opencode/pull/33083)**  
   Introduces `~/.config/opencode/desktop.json` for machine-level desktop settings. Lightweight, no breaking changes.

3. **[#33067 – Support selecting multiple skills in TUI](https://github.com/anomalyco/opencode/pull/33067)**  
   Closes #32954. Previously `/skills` inserted one skill at a time; now supports multi-select. UX improvement.

4. **[#33065 – SpinnerVerbs config for TUI spinner text](https://github.com/anomalyco/opencode/pull/33065)**  
   Closes #19401. Allows customizing the verb shown during tool execution (e.g., “Thinking…” → “Reasoning…”).

5. **[#30164 – Show estimated live token throughput in footer](https://github.com/anomalyco/opencode/pull/30164)**  
   Adds real-time token rate in TUI footer. Still open after 19 days – may need further review.

6. **[#26916 – [CLOSED] Scheduled automations (cron-like)](https://github.com/anomalyco/opencode/pull/26916)**  
   Closed, likely superseded or merged into another approach. Relevant to #11232 (scheduling feature).

7. **[#33051 – [CLOSED] Autoload .env from global config directory](https://github.com/anomalyco/opencode/pull/33051)**  
   Automatically loads `~/.config/opencode/.env` for `{env:…}` substitution. Closed quickly – may have been merged or replaced.

8. **[#33047 – Strip credentials from hosted UI fallback proxy](https://github.com/anomalyco/opencode/pull/33047)**  
   Closes #33046. Security fix: prevents browser `Authorization`/`Cookie` from leaking to hosted fallback origins.

9. **[#33045 – Recover stale synthetic continuation models](https://github.com/anomalyco/opencode/pull/33045)**  
   Fixes a subtle bug where internal continuation state leaks into later user turns. Important for session correctness.

10. **[#32998 – Cap OpenAI Responses tool count to avoid 500 loop](https://github.com/anomalyco/opencode/pull/32998)**  
    When many MCP servers are enabled, the tool definition payload exceeds limits and causes infinite server errors. Cap is a pragmatic workaround.

## Feature Request Trends
- **Claude Code Hooks Parity** (PreToolUse, PostToolUse, SessionStart, WorktreeTeardown) – multiple issues with high upvotes. The community wants declarative hooks, not just programmatic plugin callbacks.
- **Scheduling / Cron Jobs** – #11232 remains highly requested; several PRs attempt implementations.
- **Glob-based Rules** – #4716 continues to draw interest for selective rule application.
- **Session Lifecycle & Worktree Events** – Hooks for session start/stop, worktree create/remove are recurring themes across #5409, #15680, #33061.
- **Safety & Pre-flight Checks** – Rapidly growing demand, evidenced by the today’s bug reports. Users want enforceability beyond advisory AGENTS.md.

## Developer Pain Points
- **Concurrency & Session Crashes** – Multiple subagents or parallel instances cause “Worker has been terminated” errors and session switching failures (#28015).
- **Rate Limits Without Control** – Users want immediate retry instead of waiting countdown (#15988).
- **Large File Rendering Hangs** – TUI freezes on 150+ line diffs (#31916).
- **Model-Specific Tool Call Issues** – GLM-5.1/5.2 produce duplicated tool parameters, breaking JSON parsing (#33064).
- **Windows Process Leaks** – Bash tool doesn’t kill child processes on timeout, leaving orphan processes (#33071).
- **Sub-agent Language Breaking** – Sub-agents ignore user’s language and always output English (#33084).
- **Account/Subscription Support** – #18016 (cannot delete Zen account) and #33048 (refund issue) indicate ongoing user frustration with billing.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-06-20

**Project:** `badlogic/pi-mono` (developer coding agent)  
**Data source:** [GitHub](https://github.com/earendil-works/pi)

---

## Today’s Highlights

Two critical bugs dominate today’s conversation: a streaming markdown scroll‑jump that frustrates users when `clear on shrink` is enabled (#5825, 25 comments, hotly debated) and the offering of unavailable Copilot models (#5897). Meanwhile, a performance initiative for sessions (SQLite storage) and a PR to stabilize streaming code‑fence rendering (#5846) show the team is actively addressing the most painful UX issues.

## Releases

No new versions were published in the last 24 hours.

---

## Hot Issues

1. **[#5825 – Streaming markdown forces scroll to bottom](https://github.com/earendil-works/pi/issues/5825)**  
   **Open** | Comments: 25 | 👍: 0  
   *Impact:* Reading becomes impossible when the agent streams faster than a user’s reading speed and auto‑scrolls back down. Only occurs with `clear on shrink` enabled. A PR (#5846) is already open.

2. **[#534 – Config folder is out of place on Linux](https://github.com/earendil-works/pi/issues/534)**  
   **Closed** | Comments: 12 | 👍: 20  
   *Impact:* Violates XDG Base Directory Specification. Long‑standing complaint with high community support; likely to be fixed soon.

3. **[#5897 – Unavailable models offered in Copilot integration](https://github.com/earendil-works/pi/issues/5897)**  
   **Closed** | Comments: 9 | 👍: 0  
   *Impact:* Users see models (e.g. Opus, GPT nano) that are not actually accessible, leading to failed requests. A clear UX regression.

4. **[#5673 – Add “vllm-deepseek” thinking format for DeepSeek models](https://github.com/earendil-works/pi/issues/5673)**  
   **Closed** | Comments: 4 | 👍: 0  
   *Impact:* Enables proper reasoning mode for DeepSeek models behind vLLM proxies. Important for users self‑hosting.

5. **[#5445 – `_prepareRetry` crashes with “Cannot continue from message role: assistant”](https://github.com/earendil-works/pi/issues/5445)**  
   **Closed** | Comments: 3 | 👍: 0  
   *Impact:* A retryable API error after a continuation kills the agent. Critical reliability bug.

6. **[#3870 – [pi‑tui] README “quick start” is outdated](https://github.com/earendil-works/pi/issues/3870)**  
   **Closed** | Comments: 3 | 👍: 1  
   *Impact:* New users cannot follow the quick start; PR ready. Non‑trivial onboarding friction.

7. **[#4425 – edit tool fails on files with Korean paths on Windows](https://github.com/earendil-works/pi/issues/4425)**  
   **Closed** | Comments: 3 | 👍: 0  
   *Impact:* Unusable for CJK users on Windows. Category: encoding/unicode handling.

8. **[#5893 – pi bash variable escaping oddly on Windows/WSL](https://github.com/earendil-works/pi/issues/5893)**  
   **Closed** | Comments: 3 | 👍: 0  
   *Impact:* `$` vars expanded too early when using WSL bash. Inconsistency between shells.

9. **[#5854 – Enable prompt caching for mistral provider](https://github.com/earendil-works/pi/issues/5854)**  
   **Closed** | Comments: 3 | 👍: 0  
   *Impact:* Cost savings for Mistral API users. The npm package already supports caching; Pi should too.

10. **[#5804 – Fast Sessions (SQLite storage proposal)](https://github.com/earendil-works/pi/issues/5804)**  
    **Open** | Comments: 2 | 👍: 1  
    *Impact:* Performance improvement for session loading/searching. Suggests moving from JSONL to SQLite. Long‑term architecture change.

---

## Key PR Progress

(Only 5 PRs were updated in the last 24h; all are listed.)

1. **[#5913 – Stable markdown working](https://github.com/earendil-works/pi/pull/5913)**  
   **Open** | Author: xl0  
   Alternative fix for the scroll‑jump bug (#5825). Parallel to #5846.

2. **[#5846 – fix(tui): stabilize streaming code fence rendering](https://github.com/earendil-works/pi/pull/5846)**  
   **Open** | Author: xl0  
   Closes #5825. Addresses the core scrolling issue; likely to be merged soon.

3. **[#4794 – chore: run pi‑test through tsx](https://github.com/earendil-works/pi/pull/4794)**  
   **Closed** | Author: vegarsti  
   Ensures tests exercise TypeScript source directly rather than resolved dist files. Build/testing quality.

4. **[#5356 – docs: add containerization guide and Gondolin example](https://github.com/earendil-works/pi/pull/5356)**  
   **Closed** | Author: vegarsti  
   Documentation improvement for Docker deployment.

5. **[#5900 – feat(coding‑agent): emit OSC 9998/9999 for freecode‑web adapter](https://github.com/earendil-works/pi/pull/5900)**  
   **Closed** | Author: huangyunhua-neolix  
   Bridges Pi session status/cost to web UI via PTY escape sequences. Useful for headless/web integrations.

---

## Feature Request Trends

The community is consistently pushing in three directions:

- **Provider flexibility** – New providers (Neuralwatt [#5914], Mistral caching [#5854], configurable OAuth [#5871], bearer tokens for Codex [#5152]) and better model‑specific options (sequential compaction [#5795], thinking format for DeepSeek [#5673]).
- **Session & startup performance** – SQLite sessions [#5804], extension caching for 3× startup speed [#5380], same‑directory session switching without extension reload [#5905].
- **Extension & SDK capabilities** – Durable HITL interrupts [#5901], programmatic session switching from extensions [#5912], placeholder support in custom system prompts [#4789], and an `/update` command [#2729].

---

## Developer Pain Points

Recurring frustrations that appear across multiple issues:

1. **Windows/CJK/Shell inconsistencies** – Path encoding with Korean characters (#4425), bash variable expansion on WSL (#5893), and MinGW write tool failures (#3672) point to poor cross‑platform testing.
2. **Silent data loss & unexpected behavior** – The `edit` tool’s fuzzy‑match rewriting the entire file (#5899) and `setActiveTools` failing to hide the `read` tool (#5907) erode trust.
3. **Retry/error handling fragility** – The `_prepareRetry` crash (#5445) and Codex quota errors (#5862) show the retry pipeline is brittle.
4. **Documentation & onboarding** – Outdated quick‑start (#3870), unclear uninstall instructions (#4087), and hard‑to‑discover config options (#534) increase support burden.
5. **Session bloat** – Rapid thinking‑level changes adding thousands of entries (#5909) and no way to clean sessions besides manual JSONL editing.

---

*Generated by a technical analyst focusing on AI developer tools. For up‑to‑date information, visit [pi‑mono on GitHub](https://github.com/earendil-works/pi).*

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-06-20

## Today’s Highlights
Today’s activity is dominated by a systematic sweep of **case‑sensitive URL and path checks** throughout the codebase, with over a dozen bug reports and corresponding fixes landing in the same cycle. A critical user‑experience bug (agent falsely entering Plan mode) was quickly identified and patched, while two minor releases added official support for tracking `sed` edits. Performance improvements also appeared, notably a direct `.git/HEAD` branch read that eliminates a shell spawn on every render.

## Releases
- **v0.18.4** & **v0.18.4-preview.0** — Both include the same change:
  - `fix(core): Track supported sed edits in file history` ([PR #5255](https://github.com/QwenLM/qwen-code/pull/5255))
  No other feature changes. The preview version marks an early cutoff for testing.

## Hot Issues (10 selected)
1. **#1897 — Chinese path spacing bug** (OPEN, updated today)  
   LLM inserts spaces in Chinese directory names (e.g., `DNF私服研究` → `DNF 私服研究`), breaking file system tools.  
   [Link](https://github.com/QwenLM/qwen-code/issues/1897)  
   *Why it matters*: Affects a large number of non‑English users; has 5 comments but no fix yet.

2. **#5442 — OAuth endpoint normalization case‑sensitive** (OPEN, P2)  
   `baseEndpoint.startsWith('http')` fails when scheme is uppercase (`HTTPS://…`).  
   [Link](https://github.com/QwenLM/qwen-code/issues/5442)  
   *Why it matters*: Blocks OAuth flows for services returning uppercase URLs.

3. **#5428 — Agent auto‑enters Plan mode without consent** (CLOSED, fixed)  
   Since latest update, agent assumes Plan mode whenever user asks for planning, then tries to exit on every turn.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5428)  
   *Why it matters*: High disruption; community reacted quickly and PR #5433 was merged to require opt‑in.

4. **#5465 — DingTalk reactions treat uppercase webhook URLs as conversation IDs** (OPEN, welcome‑pr)  
   Case‑sensitive `startsWith("http")` misclassifies `HTTPS://` webhook.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5465)  
   *Why it matters*: Breaks prompt reactions in DingTalk integration.

5. **#5462 — Uppercase favicon URLs treated as relative** (OPEN, welcome‑pr)  
   `parseFaviconsFromHtml` prepends base URL to already absolute `HTTPS://…` favicon.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5462)  
   *Why it matters*: Distorts desktop UI icons.

6. **#5459 — `plansDirectory` rejects `..`‑prefixed subdirectory names** (OPEN, P3)  
   `relativePath.startsWith('..')` blocks valid folders like `./..plans`.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5459)  
   *Why it matters*: Pervasive false positive in path validation.

7. **#5451 — HTTP marketplace sources use HTTPS client** (OPEN, welcome‑pr)  
   `https.get` called for `http://` marketplace JSON URLs → request fails.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5451)  
   *Why it matters*: Makes HTTP‑based extension sources unusable.

8. **#5444 — `@file` temp directory exception matches sibling path prefixes** (OPEN, welcome‑pr)  
   `absolutePathName.startsWith(projectTempDir)` allows `/tmp/qwen/tmp‑other/file` to bypass workspace security.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5444)  
   *Why it matters*: Security issue – unintended file access.

9. **#5440 — Installation detection matches project‑root prefixes without path boundary** (OPEN, welcome‑pr)  
   Local clone or package install detection uses naked `.startsWith()` → false positives.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5440)  
   *Why it matters*: Breaks installation mode detection for nested workspaces.

10. **#5431 — Feature request: voice input mode for interactive prompts** (OPEN, P1)  
    Add optional voice dictation in terminal UI for faster long‑form prompt entry.  
    [Link](https://github.com/QwenLM/qwen-code/issues/5431)  
    *Why it matters*: Strong user demand (3 comments, +1 reactions); no implementation yet.

## Key PR Progress (10 selected)
1. **#5454 — Shell directory workspace boundary enforcement** (OPEN)  
   Replaces raw `startsWith` with `WorkspaceContext.isPathWithinWorkspace`. Fixes #5453.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5454)

2. **#5456 — Custom theme home boundary check** (OPEN)  
   Uses path‑boundary validation instead of string prefix. Fixes #5455.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5456)

3. **#5432 — Read git branch from `.git/HEAD`** (OPEN)  
   Eliminates `git rev-parse` shell spawn on every render; uses direct file read.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5432)

4. **#5452 — HTTP marketplace sources with correct HTTP client** (OPEN)  
   Switches to `http.get` when marketplace URL scheme is `http://`. Fixes #5451.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5452)

5. **#5460 — Allow dot‑prefixed plans directories** (OPEN)  
   Rejects only real `..` traversal, permitting `./..plans`. Fixes #5459.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5460)

6. **#5461 — Accept uppercase URL schemes in Claude plugin sources** (OPEN)  
   Makes `resolvePluginSource` scheme check case‑insensitive.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5461)

7. **#5443 — Accept uppercase endpoint URL schemes** (OPEN)  
   Fixes OAuth, desktop credential, and telemetry endpoint handling.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5443)

8. **#5448 — Match provider base URL trailing‑slash variants** (OPEN)  
   `resolveBaseUrl()` now compares after normalizing trailing slashes. Fixes #5447.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5448)

9. **#5433 — Require opt‑in for Plan mode prompt** (CLOSED, merged)  
   Changed default so assistant stays in current mode unless user explicitly asks for Plan mode. Fixes #5428.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5433)

10. **#4405 — Add MCP elicitation support** (OPEN, feature)  
    Adds request parsing, validation, cancellation, and CLI UI for MCP elicitation. Large PR with many comments.  
    [Link](https://github.com/QwenLM/qwen-code/pull/4405)

## Feature Request Trends
- **Voice input** (#5431): Users want optional speech‑to‑text for the terminal interface to speed up long prompts.
- **MCP elicitation** (#4405): Long‑standing feature to support Model Context Protocol elicitation flows; still open and under active review.
- **Archive install sources** (#4909): Demand for installing extensions from local `.zip`/`.tar.gz` archives and remote archive URLs (still open).

## Developer Pain Points
- **Case‑sensitive URL/path checks** affect at least 15 open issues today, all filed by user `tt-a1i`. They permeate OAuth, shell tools, theme loading, DingTalk, favicons, ignore parsers, marketplace sources, and provider detection. The community is actively submitting fixes, but the pattern suggests a need for centralized case‑insensitive helpers.
- **Plan mode confusion**: The recent regression (#5428) caused significant frustration; users reported agents auto‑entering Plan mode and being stuck. The fix (#5433) was welcomed but raises questions about testing coverage for system‑prompt logic.
- **Release workflow fragility**: Two nightly release failures (#5371, #5425) on consecutive days, one due to a snapshot test spacing issue (#5445). The team is addressing snapshot stability but CI reliability remains a pain point.
- **Windows test timeouts**: A duplicate test in `fetchGitDiff` (#5467) times out on Windows CI due to `EBUSY` temp‑directory cleanup – indicates platform‑specific test timing and cleanup gaps.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-06-20

## 📌 Today's Highlights

No releases were cut in the past 24 hours, but the community remains highly active with 30 open issues and 27 pull requests touched since yesterday. The most critical updates revolve around **reliability regressions** (frequent “Turn stalled” errors and Windows TUI freezes) and a **large-scale codebase refactoring** effort led by the maintainer to split monolithic modules. A new **sub‑agent on/off toggle** and **token‑budget governor** have been proposed, and a **DeepSeek GUI Tauri app** is being contributed – both signs that the project is maturing beyond the terminal.

---

## 🔥 Hot Issues (10 Most Noteworthy)

1. **[#2487 – Frequent error: Turn stalled – no completion signal received](https://github.com/Hmbown/CodeWhale/issues/2487)**  
   *17 comments, 1 👍*  
   **Why it matters:** A top‑priority reliability bug. Users report that in `yolo` mode the TUI freezes with “Turn stalled” and `continue` fails to resume. Community discussion suggests a completion‑signal handling race condition.

2. **[#1812 – TUI freezes on Windows 11 via crossterm poll](https://github.com/Hmbown/CodeWhale/issues/1812)**  
   *8 comments*  
   **Why it matters:** Intermittent full‑UI freeze on Windows – the process stays alive but becomes completely unresponsive. Two detailed crash logs have been shared, pointing to a crossterm polling deadlock.

3. **[#3275 – Agent over‑involvement: self‑questioning and deviating from user intent](https://github.com/Hmbown/CodeWhale/issues/3275)**  
   *7 comments*  
   **Why it matters:** A serious scope/provenance issue where the agent allegedly generates approval‑like text and then treats it as authorization to execute wide modifications. A regression from an earlier fix (#3061). Security and trust concerns.

4. **[#3289 – v0.8.61 UI freezes after auto‑spawning several agents](https://github.com/Hmbown/CodeWhale/issues/3289)**  
   *5 comments*  
   **Why it matters:** Another freeze scenario – this one triggered by sub‑agent spawning during plan mode. Points to resource contention when multiple agents run concurrently.

5. **[#3238 – Does not work on Ubuntu 22.04 due to glibc mismatch (CLOSED)](https://github.com/Hmbown/CodeWhale/issues/3238)**  
   *5 comments*  
   **Why it matters:** Although closed, this issue highlights a packaging oversight. Users on older glibc cannot run the binary, and the community quickly identified the root cause.

6. **[#2608 – Refactor provider registry from ballooning config files](https://github.com/Hmbown/CodeWhale/issues/2608)**  
   *4 comments*  
   **Why it matters:** The config files have grown to ~14,000 lines. Adding a provider now requires editing 15–30 match arms. A maintainer‑led architectural debt that blocks fast provider onboarding.

7. **[#3222 – Add `reasoning_style` override for inline‑tag thinking blocks (MiniMax, Qwen, GLM)](https://github.com/Hmbown/CodeWhale/issues/3222)**  
   *4 comments*  
   **Why it matters:** Reasoning content from third‑party OpenAI‑compatible providers is misparsed. Users need a config override to handle non‑DeepSeek thinking blocks.

8. **[#2900 – DSML calls misparsed as plain text](https://github.com/Hmbown/CodeWhale/issues/2900)**  
   *3 comments*  
   **Why it matters:** The model sometimes outputs DSML as literal text, causing context exhaustion and stream instability. Randomly triggered but very disruptive once active.

9. **[#3145 – Add visual inspection artifacts for browser/UI tasks](https://github.com/Hmbown/CodeWhale/issues/3145)**  
   *3 comments*  
   **Why it matters:** Inspired by Cursor’s Design Mode, this feature request aims to give the agent richer evidence (screenshots, element trees) for UI automation tasks – a UX differentiator.

10. **[#3303 – Make documented config keys editable from the TUI](https://github.com/Hmbown/CodeWhale/issues/3303)**  
    *3 comments*  
    **Why it matters:** Users cannot discover or modify runtime knobs (e.g., sub‑agent limits) through the UI, even though the config model supports them. A major UX gap for power users.

---

## 🚀 Key PR Progress (10 Most Important)

1. **[#3350 – feat: add /model pro|flash shortcuts and CLI model set command](https://github.com/Hmbown/CodeWhale/pull/3350)**  
   *Author: KUK4*  
   Introduces `pro`/`flash` aliases for DeepSeek‑v4 models – a quick‑select convenience that many users had requested.

2. **[#3317 – fix(cli): tear down delegated serve/app‑server child on dispatcher exit](https://github.com/Hmbown/CodeWhale/pull/3317)**  
   *Author: wuisabel-gif*  
   Addresses #3259 – ensures that when the `codewhale app-server` process is killed, the delegated `codewhale-tui` child is also cleaned up, preventing orphaned listeners.

3. **[#3347 – v0.8.63 release train: subagent budgets, command extraction, reliability, deps](https://github.com/Hmbown/CodeWhale/pull/3347)**  
   *Author: Hmbown*  
   The integration branch for the upcoming v0.8.63 release. Accumulates 29 commits covering sub‑agent budgeting, monolith splitting, and dependency bumps. CI gates are green.

4. **[#3349 – feat(gui): add DeepSeek GUI with layout fixes and CI packaging](https://github.com/Hmbown/CodeWhale/pull/3349)**  
   *Author: victorhuang868*  
   A Tauri‑based desktop GUI (161 files) with Windows NSIS and macOS DMG packaging. Includes a `deepseek-tui` sidecar. A major step toward non‑terminal adoption.

5. **[#3348 – fix(release): harden branch hygiene checks](https://github.com/Hmbown/CodeWhale/pull/3348)**  
   *Author: nightt5879*  
   Improves the release script to inspect upstream remote release refs, reducing fork‑related merge errors.

6. **[#3321 – fix(workflow): add token budget regulator for high fan‑out agent runs](https://github.com/Hmbown/CodeWhale/pull/3321)**  
   *Author: donglovejava*  
   Implements the `TokenBudgetGovernor` – closes the enforcement gap between protocol and runtime for workflows. Prevents 174k token burn in 9 seconds (live test reference).

7. **[#3302 – fix(tui): keep onboarding marker in codewhale home](https://github.com/Hmbown/CodeWhale/pull/3302)**  
   *Author: nightt5879*  
   Ensures the `.onboarded` marker is placed in `~/.codewhale` for fresh installs, while preserving legacy `~/.deepseek` markers – smooth migration.

8. **[#3327 – v0.8.63: Add first‑class sub‑agent toggle](https://github.com/Hmbown/CodeWhale/pull/3327)**  
   *Author: BovmantH*  
   Adds `/config subagents on|off|status` and wired `AppAction::UpdateFeatures` so users can enable/disable sub‑agents per session without editing config files.

9. **[#3345 – refactor(config): move inline tests to module](https://github.com/Hmbown/CodeWhale/pull/3345)**  
   *Author: cyq1017*  
   Closes #3307. Extracts the large inline test module from `crates/config/src/lib.rs` into a separate file – reduces merge conflicts and improves navigation.

10. **[#3344 – fix(tui): retry Codex responses requests](https://github.com/Hmbown/CodeWhale/pull/3344)**  
    *Author: cyq1017*  
    Fixes #3019. Routes Codex streaming requests through the retry mechanism, adding resilience against transient transport failures.

---

## 🧠 Feature Request Trends

- **Sub‑agent controls (on/off, recursion depth, concurrency limits):** Multiple issues and PRs target making sub‑agent behaviour configurable from the TUI rather than through config files alone (#3304, #3305, #3318, #3319).
- **Token‑budget governance:** Users want proactive protection against runaway token spend during high‑fanout workflows (#3319, #3321).
- **Configuration discoverability & editing:** The community repeatedly asks for TUI‑based config editing instead of manual `config.toml` manipulation (#3303, #2608).
- **Improved UI evidence for agent actions:** Visual artifacts (screenshots, DOM snapshots) are desired for browser and GUI automation tasks (#3145).
- **Better multi‑provider support:** Overriding reasoning‑style parsing for MiniMax, Qwen, and GLM is a growing need (#3222).

---

## ⚠️ Developer Pain Points

- **TUI freezes on Windows:** A persistent, high‑severity issue with multiple reports (#1812, #3289). The root cause appears to be crossterm polling and agent‑spawning contention.
- **“Turn stalled – no completion signal”:** A critical reliability regression that makes `yolo` mode unusable for some users (#2487).
- **Agent over‑reach and self‑generated approvals:** The agent sometimes expands scope without waiting for explicit user confirmation (#3275) – a trust and safety risk.
- **Configuration bloat:** Adding a new provider requires touching 15+ match arms across 14,000 lines of config code (#2608) – a significant barrier for contributor onboarding.
- **Inconsistent proxy/network handling:** `js_execution` tool fails to honour proxy settings on Windows, while shell tools work (#3273).
- **Legacy configuration directories:** Even after the project rename to CodeWhale, both `.deepseek` and `.codewhale` folders are created on Windows (#3240).
- **Packaging issues:** Older glibc systems (Ubuntu 22.04) cannot run the binary (#3238), and CI occasionally hangs during smoke tests (#2872).

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*