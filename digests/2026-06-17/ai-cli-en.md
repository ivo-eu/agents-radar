# AI CLI Tools Community Digest 2026-06-17

> Generated: 2026-06-17 03:58 UTC | Tools covered: 9

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

# AI CLI Tools Ecosystem Cross-Tool Comparison Report — June 17, 2026

## 1. Ecosystem Overview

The AI CLI developer tools landscape continues to mature rapidly, with six major open-source tools showing active community engagement and frequent iteration. All tools are grappling with **model compatibility fractures**—Claude Code’s Opus 4.8 regressions, Gemini and Pi’s DeepSeek V4 parameter conflicts, and inconsistent tool-call handling across providers reveal that the model layer remains the weakest link. **MCP (Model Context Protocol) integration** has become the de facto standard for extensibility, but every tool faces leaks, resource exhaustion, and configuration management issues. **Windows support** remains a universal pain point, with all tools reporting process management, path encoding, and sandbox execution bugs specific to the platform. Meanwhile, **agentic loops, runaway quota consumption, and sub-agent deadlocks** are the top reliability concerns across communities, suggesting the industry has not yet solved agent orchestration at scale.

---

## 2. Activity Comparison (Last 24 Hours)

| Tool | Hot Issues (Noteworthy) | Pull Requests (Active) | Releases (Today) | Notable Signals |
|---|---|---|---|---|
| **Claude Code** | 10 (50+ total today) | 10 (18 from one contributor) | 1 (v2.1.179 patch) | Largest community, Windows/system prompt explosion, Opus 4.8 regression |
| **OpenAI Codex** | 10 | 10 (50-PR queue) | 2 (rust alphas) | Account lockout (#25749, 30👍), macOS relaunch loop, MCP OAuth coordination |
| **Gemini CLI** | 10 | 10 | 0 | CI nightly failure (P0), agent hangs (#21409), security hardening (case-insensitive blocklist) |
| **GitHub Copilot CLI** | 10 | 0 | 1 (v1.0.64-0) | Post-outage all models "Blocked" (#3832), Windows ARM64 crash, `/diagnose` command |
| **Kimi Code CLI** | 4 | 1 | 0 | Critical bugs (MCP auto-discover after delete, fresh install fails without login) |
| **OpenCode** | 10 | 10 | 0 | Random hangs (#2940, 18👍), clipboard failure (#7048), CPU-bound processing, image regression |
| **Pi** | 10 | 10 | 2 (v0.79.5, v0.79.6) | DeepSeek V4 fixes, multi-session demand, Nix flake, provider-scoped envs |
| **Qwen Code** | 10 | 10 | 0 (3 release failures) | OOM on auto-memory (#5147), vision bridge PR, CI pipeline fragility |
| **DeepSeek TUI (CodeWhale)** | 10 | 9 | 1 (v0.8.61) | Rebrand to CodeWhale, task stalls (#2487), chat-native workrooms EPIC |

---

## 3. Shared Feature Directions

### MCP Ecosystem Maturation (Present in: Claude Code, OpenAI Codex, Gemini CLI, GitHub Copilot CLI, Kimi Code CLI, OpenCode, Pi)

- **Auto-discovery conflicts**: Multiple tools struggle with stale MCP server configs (Kimi Code CLI’s unfixable 400 errors; Claude Code’s system prompt ballooning from injected MCP configs).
- **Resource leaks**: Headless `claude -p` calls booting full MCP stacks (Claude Code); Pi’s `pi list` hanging on long-lived MCP servers; OpenCode’s MCP tool schema sanitization for OpenAI compatibility.
- **OAuth token management**: Gemini CLI and OpenAI Codex both adding atomic MCP OAuth writes and coordinated refresh across clients.
- **Registry/installation**: GitHub Copilot CLI adding `/mcp` registry commands; OpenCode’s LAN provider discovery.

### Agent Reliability & Orchestration (Present in: All tools)

- **Sub-agent deadlocks and misrouting**: Claude Code’s background agent notification collisions; DeepSeek TUI’s `agent_eval` TUI freeze with multiple sub-agents; Gemini CLI’s generalist agent hang; OpenCode’s subagent-to-subagent delegation PR (#7756).
- **Plan mode enforcement failures**: Claude Code’s Plan Mode bypass, Qwen Code’s stuck-in-ExitPlanMode, DeepSeek TUI’s overly proactive agents.
- **Runaway loops and quota waste**: Claude Code sees 30–40% quota burn on cache rehydration; Gemini CLI’s Auto Memory retries low-signal sessions indefinitely; OpenCode’s infinite clarification loop on empty git repo.

### Cross-Platform Pain (Universal, especially Windows)

- **Process management**: Orphaned processes preventing relaunch (Claude Code Desktop), BEX64 crash on Windows ARM64 (GitHub Copilot CLI), `CreateProcessAsUserW` errors (OpenAI Codex).
- **Encoding**: UTF-8/CP932 mojibake (Claude Code), PowerShell UTF-8 wrapper (OpenCode PR #31985), normalization of path separators (Claude Code PR #68694).
- **Sandbox execution failures**: OpenAI Codex’s Computer Use bootstrap fails on Windows; Pi’s proxy not honoured in JS execution on Windows.

### Maturity of Configuration & Observability

- **Provider-scoped settings**: Pi adds `env` overrides per provider (v0.79.5); Qwen Code’s vision bridge for text-only models; GitHub Copilot CLI’s `/diagnose` command.
- **Token and latency metrics**: Pi adds `durationMs` and tokens/sec to footer; Claude Code community requests streaming toggle; OpenCode’s CPU-bound internal processing flagged.
- **Internationalization (i18n)**: Qwen Code localizes tool display names (PR #5220); DeepSeek TUI’s rebrand to CodeWhale needs cleanup.

### Multi-Agent Sessions & Background Workers

- Claude Code background agent notifications; Pi’s #5700 request for TUI switching between live agent sessions; DeepSeek TUI’s whole v0.9.0 EPIC around chat-native workrooms; OpenCode’s subagent delegation with budgets.

### Security Hardening

- Symlink escape blocking (Claude Code PR #68689); case-insensitive sensitive path blocklist (Gemini CLI PR #27966); atomic OAuth token writes (Gemini CLI, OpenAI Codex); pinning dependencies with cooldown (Gemini CLI); credential broker behind proxy (OpenAI Codex PR #28034).

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | GitHub Copilot CLI | Kimi Code | OpenCode | Pi | Qwen Code | DeepSeek TUI (CodeWhale) |
|---|---|---|---|---|---|---|---|---|---|
| **Core Model** | Anthropic Opus (4.8) | OpenAI (GPT-4o, o-series) | Gemini (3.1) | OpenAI models via Copilot | Kimi models (K2) | Multi-provider (OpenAI, Gemini, Ollama, etc.) | Multi-provider (DeepSeek, OpenAI, Claude, etc.) | Qwen (3.7-max) | DeepSeek (V4) |
| **Target User** | Professional developers, agent-heavy workflows | Enterprise, security-conscious teams | GCP/Google ecosystem | GitHub-centric teams | MoonshotAI users, China market | Open-source enthusiasts, power users | Tinkerers, multi-provider users | Chinese developers, Qwen ecosystem | DeepSeek power users, rebranding to generalist |
| **Top Differentiator** | Deep agentic skill system, plan mode | Managed config enforcement, workload identity | AST-aware code understanding, Auto Memory | Tight GitHub integration, `/security-review` | Simplicity, minimal configuration | Most flexible provider support, subagent delegation | DeepSeek V4 first-class, Nix, provider-scoped envs | Vision bridge, i18n, QQ bot channel | Chat-native workrooms v0.9.0 EPIC, hippocampal memory v2 |
| **Community Maturity** | Most mature, largest issue tracker | Very mature, enterprise-focused | Mature, strong security culture | Growing, tied to Copilot ecosystem | Smallest, early stage | Active open-source, 18👍 top bug | Growing, responsive maintainers | Moderate, Chinese-language heavy | Moderate, rebranding momentum |
| **Windows Support** | Poor (orphaned locks, system prompt explosion) | Poor (sandbox, Computer Use bootstrap) | Not highlighted | Poor (ARM64 crash) | Windows 10 reported crashes | Poor (path cache, clipboard) | Not highlighted | Not highlighted | Not highlighted |

---

## 5. Community Momentum & Maturity

- **Claude Code** remains the most active community by raw volume (50+ issues per day, 18 PRs in 24h from a single contributor). However, **Opus 4.8 regressions** and **Windows pain** are eroding trust; the community is vocal but frustrated.
- **OpenAI Codex** has the largest PR queue (50 open) and strong enterprise focus. The account recovery blocker (#25749 with 30👍) indicates a critical UX gap that could drive users away.
- **Gemini CLI** shows disciplined security practices (case-insensitive blocklists, dependency pinning) but is held back by a **blocked CI nightly** and persistent agent hang bugs.
- **GitHub Copilot CLI** is gaining momentum with the new `/diagnose` and `/mcp` commands, but yesterday’s outage trauma (all models blocked) and Windows ARM64 issues need rapid resolution.
- **Kimi Code CLI** has the **smallest community** but two critical bugs (auto-discovery, login guidance) that could alienate new users.
- **OpenCode** has strong open-source appeal with multi-provider flexibility, but the #2940 hang bug (18👍) and clipboard failure are severe UX blockers.
- **Pi** is **iterating fastest** with two releases in 24h, responsive to DeepSeek V4 issues, and gaining traction with Nix users. Multi-session demand (#5700) signals growing sophistication.
- **Qwen Code** is advancing i18n and platform expansion (QQ bot) but is **pipeline-fragile**—three release failures today indicate process immaturity.
- **DeepSeek TUI (CodeWhale)** is in a **transitional phase** (rebrand, v0.9.0 EPIC), with core stability issues (task stalls) still unresolved despite momentum.

**Velocity ranking** (releases + PRs): Claude Code / Pi > OpenAI Codex / Qwen Code / DeepSeek TUI > Gemini CLI / OpenCode > GitHub Copilot CLI > Kimi Code CLI.

---

## 6. Trend Signals

### For Developers Evaluating Tools

1. **MCP is the emerging standard, but immature.** Every tool implements MCP differently—resource leaks, OAuth collisions, schema incompatibilities. Expect standardization efforts in H2 2026, but for now, MCP-heavy workflows will encounter friction.

2. **Model compatibility is the #1 reliability risk.** Opus 4.8, DeepSeek V4, and other rapid model releases introduce regressions (malformed tool calls, parameter conflicts, thinking-mode issues). Tools with explicit fallback chains (OpenCode PR #27939) and model-agnostic routing will win trust.

3. **Windows remains a second-class platform.** If your team uses Windows or WSL2, expect process management, encoding, and sandbox execution bugs across all tools. Claude Code and GitHub Copilot CLI have the most reported issues; Pi and Qwen Code have fewer but that likely reflects smaller userbases.

4. **Agent orchestration is not production-ready.** Runaway loops, sub-agent deadlocks, and notification collisions are universal. Plan-mode enforcement is bypassed; quota waste is rampant. For production use, limit agent autonomy and monitor token consumption.

5. **Security hardening is accelerating.** Atomic OAuth token writes, credential brokers, dependency pinning, and path traversal protection are now baseline. Enterprise adoption is driving these changes across all tools.

6. **i18n and multi-platform support are differentiators.** Qwen Code’s i18n push and QQ bot channel, DeepSeek TUI’s workrooms, and Pi’s Nix flake suggest that tools are targeting specific ecosystems rather than competing on raw intelligence.

7. **Community health signals are mixed.** Claude Code’s large but frustrated community contrasts with Pi’s smaller but highly responsive maintainers. For long-term support, tools with transparent release pipelines (e.g., Pi’s frequent releases) and active security PRs (Gemini CLI) may be more reliable than those with fragile CI (Qwen Code).

### For Tool Maintainers

- **Model compatibility testing must be mandatory in CI.** The Qwen Code pipeline failures and Gemini CLI nightly failures show that integration tests are too often deferred. E2E tests on every PR (as Gemini CLI is attempting) should be the standard.
- **Windows is not optional.** The consistent stream of Windows-specific bugs across all tools indicates that Electron-based architectures and cross-platform path handling need fundamental rework. Consider platform-specific champions.
- **User-perceptible metrics** (token latency, quota remaining, time-to-first-token) are becoming expected features. Pi’s footer metrics and Claude Code’s streaming toggle request signal a demand for observability.
- **Clarify model selection and fallback behavior.** Silent downgrades (Copilot CLI) and opaque model switching (Gemini CLI’s sub-agents) erode trust. Explicit chains and transparent fallback models are the path forward.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report  
*Data as of 2026-06-17 | Source: github.com/anthropics/skills*

---

## 1. Top Skills Ranking  

Based on discussion volume, cross-referencing with related issues, and overall attention, the eight most-discussed Skill submissions are:

| Rank | PR # | Skill / Title | Functionality | Discussion Highlights | Status |
|------|------|---------------|---------------|-----------------------|--------|
| 1 | [#514](https://github.com/anthropics/skills/pull/514) | **document-typography** | Prevents orphan words, widowed headers, and numbering misalignment in AI-generated documents. | Broad agreement that typographic quality is a universal pain point. Received support from document-format authors. | **OPEN** |
| 2 | [#486](https://github.com/anthropics/skills/pull/486) | **ODT (OpenDocument) skill** | Creates, fills, reads, and converts ODT/ODS files; also parses ODT to HTML. | Strong interest from LibreOffice/enterprise users. Discussion focused on template-filling robustness and ISO-standard compliance. | **OPEN** |
| 3 | [#210](https://github.com/anthropics/skills/pull/210) | **frontend-design (revision)** | Rewrites the existing frontend-design skill to be more actionable: every instruction must be executable within a single conversation. | Detailed debate about token efficiency vs. completeness; reviewers praised the clarity push. | **OPEN** |
| 4 | [#83](https://github.com/anthropics/skills/pull/83) | **skill-quality-analyzer + skill-security-analyzer** | Two meta-skills: one evaluates Skill quality across five dimensions (structure, documentation, examples, etc.), the other checks for security issues. | High interest in meta-level tooling for skill creators. Discussion raised concerns about false positives. | **OPEN** |
| 5 | [#538](https://github.com/anthropics/skills/pull/538) | **PDF: case-sensitive file references** | Fixes 8 uppercase/lowercase mismatches in `skills/pdf/SKILL.md` that break on case-sensitive filesystems. | Small but critical fix; uncovered broader documentation consistency problems. | **OPEN** |
| 6 | [#541](https://github.com/anthropics/skills/pull/541) | **DOCX: tracked change w:id collision** | Prevents document corruption when tracked changes are added to DOCX files that already contain bookmarks with conflicting `w:id` values. | Root cause analysis (OOXML shared ID space) was widely appreciated; several users confirmed the bug. | **OPEN** |
| 7 | [#1298](https://github.com/anthropics/skills/pull/1298) | **skill-creator: run_eval.py recall=0% fix** | Addresses the systematic 0% recall bug in the evaluation/optimization loop (related to issues #556, #1169). Also fixes Windows stream reading, trigger detection, and parallel workers. | Most commented PR in recent weeks; multiple independent reproductions. Considered a blocker for skill development. | **OPEN** |
| 8 | [#154](https://github.com/anthropics/skills/pull/154) | **shodh-memory** | Persistent memory system for AI agents: teaches Claude to call `proactive_context` on every user message, structure memories with rich content, and maintain context across conversations. | Discussion focused on privacy and memory eviction strategies. Authors provided detailed usage examples. | **OPEN** |

---

## 2. Community Demand Trends  

From the 13 most-commented Issues, five clear demand directions emerge:

- **🔧 Skill-Developer Tooling** – Issues #556, #1169, #1061, #202 all point to a pressing need for a stable, cross-platform skill creation pipeline. Windows compatibility (`PATHEXT`, `cp1252`, `select()` on pipes) and the `run_eval.py` recall bug dominate. The community wants a hardened `skill-creator` that actually works out of the box.
- **🛡️ Security & Trust Boundaries** – Issue #492 highlights impersonation risks under the `anthropic/` namespace. Issue #1175 raises concerns about embedding access-control logic directly in `SKILL.md`. Demand for security-audit meta-skills (like PR #83) is growing.
- **🏢 Organizational Sharing & Governance** – Issue #228 (14 comments, 👍7) calls for org-wide skill libraries/direct sharing links. Issue #412 proposes an *agent-governance* skill for policy enforcement and audit trails.
- **📄 Document-Format Completeness** – Beyond the existing PDF/DOCX skills, users regularly request ODT (#486), improved typography (#514), and better handling of tracked changes (#541). The `agentskills.io` redirect error (#184) also frustrated users seeking documentation.
- **🔗 MCP & Platform Integration** – Issue #16 advocates exposing Skills as MCP tools. Issue #29 asks about AWS Bedrock compatibility. The community wants Skills to work beyond Claude Code’s native environment.

---

## 3. High-Potential Pending Skills  

These open PRs have active discussion, clear value propositions, and are likely to be merged soon:

| PR | Skill | Why It Stands Out |
|----|-------|-------------------|
| [#1298](https://github.com/anthropics/skills/pull/1298) | `run_eval.py` fix bundle | Blocking all description-optimization workflows. Multiple contributors have verified the fix. Merge expected within days. |
| [#723](https://github.com/anthropics/skills/pull/723) | **testing-patterns** | Covers the full testing stack (unit, React, integration, E2E) with clear philosophy. Fills a gap in the collection. |
| [#568](https://github.com/anthropics/skills/pull/568) | **ServiceNow platform** | Broad enterprise assistant covering ITSM, ITOM, SecOps, etc. High upvote count and detailed documentation. |
| [#444](https://github.com/anthropics/skills/pull/444) | **AURELION suite** (kernel, advisor, agent, memory) | Structured cognitive framework with 5-floor thinking model + memory. Well-received by knowledge-management users. |
| [#335](https://github.com/anthropics/skills/pull/335) | **Masonry image/video generation** | Integrates Imagen 3.0 and Veo 3.1 via CLI. Direct demand from creative workflows. |

All are **OPEN**; none have been merged as of the data snapshot.

---

## 4. Skills Ecosystem Insight  

**The community’s most concentrated demand is for robust skill-developer tooling (especially Windows compatibility and evaluation-loop accuracy) and for high-quality document-format skills (typography, ODT, PDF, DOCX) that address the persistent quality gaps in AI-generated documents.**

---

# Claude Code Community Digest — June 17, 2026

**Edition v2026-06-17** – A daily briefing for developers building with Claude Code.

---

## Today's Highlights

An urgent patch (v2.1.179) lands to fix mid-stream connection drops and a WSL2 scroll regression. Meanwhile, the community coalesces around two Opus 4.8 regressions—malformed tool calls and severe performance degradation—while a new wave of Windows-specific bugs, MCP memory leaks, and agent-loop quota consumption dominate the issue tracker. A flurry of 18 PRs, mostly from a single contributor, tighten security and fix scripting edge cases across platforms.

---

## Releases

### [v2.1.179](https://github.com/anthropics/claude-code/releases/tag/v2.1.179) – Connection Drop & WSL2 Fix

- **Fixed:** Mid-stream connection drops no longer show raw errors; partial responses are preserved and the spinner no longer hangs on "running tool".
- **Fixed:** Mouse-wheel scrolling in WSL2 under Windows Terminal and VS Code (regression introduced in v2.1.172).
- **Fixed:** Sandbox `denyR` (likely a deny-read permission) now applies correctly.

*No other releases in the last 24 hours.*

---

## Hot Issues (10 Noteworthy)

Highlights from 50+ issues updated today, selected by community engagement and impact.

1. **[#42776 – Claude Code Desktop fails to Relaunch on Windows due to orphaned process file lock](https://github.com/anthropics/claude-code/issues/42776)**  
   *87 comments, 31 👍* – A long-standing bug (April) that prevents relaunch after exit. Community workarounds involve manually killing orphaned processes. Critical for Windows users relying on the desktop app.

2. **[#65514 – Usage credits required for 1M context despite Pro plan with 17% usage](https://github.com/anthropics/claude-code/issues/65514)**  
   *17 comments, 2 👍* – Pro users hitting a wall when trying to use the 1M context window, even with plenty of quota left. Highly visible as context size grows in importance.

3. **[#63604 – Opus 4.8 repeatedly emits malformed `tool_use` blocks, entire response discarded](https://github.com/anthropics/claude-code/issues/63604)**  
   *10 comments, 12 👍* – A clear regression from 4.7. Opus 4.8 responses are sometimes thrown away, wasting tokens and breaking workflows. Multiple duplicates flagged.

4. **[#65429 – System prompt consumes ~9.3M tokens every session after installing Claude Desktop on Windows (WSL)](https://github.com/anthropics/claude-code/issues/65429)**  
   *9 comments* – After installing the Windows desktop app, every WSL session gets a system prompt ballooning to ~9.3M tokens, likely due to MCP server configs being injected. Massively inflates costs.

5. **[#68933 – skill-creator eval/optimizer leaks MCP child processes via headless `claude -p`](https://github.com/anthropics/claude-code/issues/68933)**  
   *4 comments* – A memory-exhaustion bug: each headless `claude -p` call boots MCP servers, and the skill-creator plugin spawns dozens, leading to forced hard reboots. Affects macOS and heavy MCP users.

6. **[#68973 – High quota consumption (~30–40%) on first request after limit reset due to expired prompt cache](https://github.com/anthropics/claude-code/issues/68973)**  
   *2 comments* – After hitting the Pro rate limit and waiting, the first request re-hydrates all caches from scratch, burning significant quota. Users report wasted tokens.

7. **[#68961 – Excessive agentic loop iterations consuming API usage quota](https://github.com/anthropics/claude-code/issues/68961)**  
   *2 comments* – "Claude keeps sending dozens/hundreds of agents using up all of my usage" – a sharp complaint about runaway agent loops. Echoes earlier Plan Mode enforcement issues.

8. **[#68969 – Workflow tool: `args` arrives JSON-stringified, not as object; named workflow edits don't hot-reload](https://github.com/anthropics/claude-code/issues/68969)**  
   *2 comments* – Custom saved workflows (`~/.claude/workflows/devloop.js`) break because `args` is passed as a string, not a structured object. Reproducible.

9. **[#68065 – Background agent notifications route through wrong agent ID when launched sequentially](https://github.com/anthropics/claude-code/issues/68065)**  
   *2 comments, 2 👍* – Sequential background agent notifications collide; the second agent's completion arrives on the first agent's task ID. Causes confusion in execution trees.

10. **[#68982 – Cloud session: message silently dropped – UI shows running, no tokens consumed, message disappears after refresh](https://github.com/anthropics/claude-code/issues/68982)**  
    *2 comments* – In cloud Claude Code sessions, messages can enter a "running" state and then vanish server-side, never persisted. Hard to reproduce but alarming for remote users.

---

## Key PR Progress (10 Important Pull Requests)

PRs updated in the last 24 hours, selected for impact and relevance.

1. **[#46351 – Enable PowerShell tool on macOS and Linux when pwsh is available](https://github.com/anthropics/claude-code/pull/46351)**  
   ✅ *Closed.* Unlocks PowerShell tool (`CLAUDE_CODE_USE_POWERSHELL_TOOL=1`) on non-Windows platforms. Big win for macOS/Linux users with PowerShell installed.

2. **[#68786 – Avoid shell injection in test-hook.sh via stdin redirection](https://github.com/anthropics/claude-code/pull/68786)**  
   Fixes insecure `bash -c` string interpolation in `test-hook.sh` – a potential injection vector in plugin-dev tooling.

3. **[#68785 – Hook JSON to stdout, tighten `su*` glob, fix CI detection and JSON injection in examples](https://github.com/anthropics/claude-code/pull/68785)**  
   Three critical fixes in hook script examples: decision JSON incorrectly sent to stderr, incorrect `su*` glob, and JSON injection in example hooks. Essential for plugin developers relying on reference implementations.

4. **[#68689 – Block symlink escape in extensibility config reads](https://github.com/anthropics/claude-code/pull/68689)**  
   Security hardening: prevents malicious symlinks from reading arbitrary files when Claude Code reads plugin/config directories.

5. **[#68694 – Normalize `CLAUDE_PLUGIN_ROOT` path separators on Windows](https://github.com/anthropics/claude-code/pull/68694)**  
   Fixes backslash/forward-slash inconsistencies on Windows, which caused plugin loading failures. Important for Windows users of MCP and skills.

6. **[#68699 – Add Python wrapper and normalize plugin root paths on Windows for hookify](https://github.com/anthropics/claude-code/pull/68699)**  
   Further Windows path normalization in hookify, plus a Python wrapper to handle cross-platform shell differences.

7. **[#68702 – Guard `PROMPT_PARTS` expansion against `set -u` on bash 3.x (macOS)](https://github.com/anthropics/claude-code/pull/68702)**  
   Fixes a crash on macOS default bash (3.x) when `set -u` is active – relevant for shell-script based plugins and hooks.

8. **[#68707 – Add `/bug` command to file GitHub issues from the terminal](https://github.com/anthropics/claude-code/pull/68707)**  
   A new slash command (`/bug`) to create GitHub issues directly from the Claude Code TUI – streamlining bug reporting.

9. **[#68678 – Don't mark Claude Desktop issues as invalid in triage](https://github.com/anthropics/claude-code/pull/68678)**  
   Process fix: prevents the triage bot from auto-closing issues reported against the desktop app when labelled "invalid".

10. **[#68680 – Safe JSON construction and correct event name in `log-issue-events`](https://github.com/anthropics/claude-code/pull/68680)**  
    Fixes malformed JSON in event logging and a wrong event name – improves reliability of issue-tracking automation.

---

## Feature Request Trends

Distilled from open issues and enhancement requests, the community is pushing for:

- **MCP ecosystem maturation** – Expanded MCP server support (Microsoft 365 attachments, better stdio troubleshooting docs, lower system prompt overhead, and MCP visibility in subprocesses).
- **Streaming output control** – A toggle to disable token-by-token streaming (👍15) to reduce distraction and help parallel work.
- **Image inline preview in TUI** – Pasted image tags should become clickable links or open OS-level previews, especially in windowed terminals.
- **Plan Mode enforcement** – Multiple users report that Plan Mode is bypassed; requests that permission bypass be blocked until the user explicitly exits Plan Mode.
- **Cross-platform agent execution** – Background agent routing, consistent IDE references in agents mode, and hot-reloading of custom workflows.

---

## Developer Pain Points

Recurring themes from issues and community reaction:

- **Windows is hurting.** Orphaned process locks, UTF-8/CP932 mojibake, TUI line overlapping, and WSL2-specific system prompt explosions make Claude Code on Windows a second-class experience.
- **Opus 4.8 is unstable.** Malformed `tool_use` blocks, severe performance degradation ("became a potato"), and long first-token latency erode trust in the premium model.
- **Quota waste is rampant.** Expired prompt caches consuming 30–40% of quota, runaway agent loops, and malformed responses being thrown away all burn credits.
- **MCP server resource leaks.** Headless `claude -p` calls boot full MCP stacks, accumulating child processes until memory exhaustion. Skill-creator workflows are particularly affected.
- **Agent routing bugs.** Background agent notifications misdirected, console messages lost in multi-agent sessions, and Plan Mode not respected undermine the agentic experience.
- **Lack of usable slash commands.** The new `/bug` PR is promising, but overall the community wants more native commands (disable streaming, clear context, etc.) to avoid context bloat.

---

*Next community digest: June 18, 2026. Data sourced from [github.com/anthropics/claude-code](https://github.com/anthropics/claude-code).*

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-06-17

## Today's Highlights
Two Rust alpha releases (0.141.0-alpha.3 and .4) landed today alongside an active 50-PR queue. The community is most vocal about a stubborn account‑recovery blocker (Issue #25749) and recurring macOS stability problems, but the internal team is making visible progress on tool namespacing, managed config enforcement, and MCP OAuth coordination.

---

## Releases (last 24h)
- **[rust-v0.141.0-alpha.3](https://github.com/openai/codex/releases/tag/rust-v0.141.0-alpha.3)** — Release 0.141.0-alpha.3 (no changelog published)
- **[rust-v0.141.0-alpha.4](https://github.com/openai/codex/releases/tag/rust-v0.141.0-alpha.4)** — Release 0.141.0-alpha.4 (no changelog published)

No desktop or CLI releases accompanied these Rust bumps. The alpha train continues without visible user‑facing changes.

---

## Hot Issues (10 selected)

1. **[#25749 – Legacy phone verification blocks account access; no recovery path](https://github.com/openai/codex/issues/25749)**  
   *46 comments, 30 👍* — User can log in via Google OAuth but Codex demands verification of an inaccessible legacy phone number with no alternative recovery route. This is the most upvoted open bug today, indicating a critical auth‑UX gap.

2. **[#25243 – macOS relaunch loop exhausts syspolicyd file descriptors](https://github.com/openai/codex/issues/25243)**  
   *33 comments, 3 👍* — Codex relaunches itself repeatedly on macOS, burning syspolicyd FDs until other apps cannot launch. Community reports it persists across several versions; a systemic Electron‑process lifecycle issue.

3. **[#28095 – Archived chat Delete button does nothing](https://github.com/openai/codex/issues/28095)**  
   *12 comments, 4 👍* — UI shows a Delete button for archived chats but the action silently fails. Frustrating for users who rely on archiving as an intermediate step before cleanup.

4. **[#19913 – Request: default parent folder for “Start from scratch” projects](https://github.com/openai/codex/issues/19913)**  
   *6 comments, 26 👍* — Highly upvoted enhancement asking for a configurable default project directory so users aren’t forced to re‑navigate to the same folder every session.

5. **[#25865 – App freezes when pasting JSON stack traces with escaped backslashes](https://github.com/openai/codex/issues/25865)**  
   *9 comments, 7 👍* — Enterprise user reports a full freeze when pasting medium‑sized JSON into the composer. Points to a parser/rendering bottleneck that affects debugging workflows.

6. **[#27287 – Computer Use bootstrap fails on Windows: @oai/sky subpath not exported](https://github.com/openai/codex/issues/27287)**  
   *9 comments, 9 👍* — Windows users cannot initialize Computer Use due to a packaging misconfiguration. Related issues #28121 and #28275 confirm this is a widespread Windows regression.

7. **[#27536 – macOS: code_sign_clone grows unbounded (62 GB+)](https://github.com/openai/codex/issues/27536)**  
   *5 comments, 0 👍 (CLOSED)* — Closed as fixed, but the original report of a temp‑folder leak consuming 62+ GB of disk space highlights a recurring Electron‑update cleanup deficiency.

8. **[#25436 – Windows: local runner fails with CreateProcessAsUserW error 5](https://github.com/openai/codex/issues/25436)**  
   *8 comments, 2 👍* — Pro user cannot start the local sandbox runner on Windows. Permission/identity elevation issue that stalls all tool‑execution workflows.

9. **[#20567 – Windows: 1000+ git processes per minute](https://github.com/openai/codex/issues/20567)**  
   *9 comments, 1 👍* — Enterprise‑grade bug: Codex spawns ~1000 git commands per minute on Windows, thrashing the process table. Most likely a chokepoint in the repo‑indexing layer.

10. **[#21509 – Request: first‑class SSH remote workspace support in Desktop](https://github.com/openai/codex/issues/21509)**  
    *4 comments, 0 👍* — Wants Codex Desktop to natively manage SSH‑remote projects (instead of requiring manual CLI). Remains a highly requested workflow enabler.

---

## Key PR Progress (10 selected)

1. **[#28409 – Enforce exact managed config values](https://github.com/openai/codex/pull/28409)**  
   Extends `requirements.toml` to pin config keys (`sqlite_home`, `model_catalog_json`, etc.) with startup warnings when violated. Strengthens enterprise managed‑device compliance.

2. **[#27713 – Prototype workload identity federation for CLI auth](https://github.com/openai/codex/pull/27713)**  
   **Do not merge.** Experiment integrating OAuth workload identity into the CLI through existing infrastructure. Single review surface while hosted contract is proven.

3. **[#28632 – Tell Codex to avoid changing rollout format](https://github.com/openai/codex/pull/28632)**  
   A prompt‑level nudge added to the path‑types skill so Codex stops mutating rollout type definitions. Small but prevents subtle data‑corruption bugs.

4. **[#28651 – exec-server: expose environment registry payloads](https://github.com/openai/codex/pull/28651)**  
   Makes internal Noise‑registration and harness‑key types public so proxying services can deserialize them without duplicating code. Enables cleaner architecture for managed runners.

5. **[#26704/26705 – TUI Plugin Sharing: remote catalog polish & coverage](https://github.com/openai/codex/pull/26704)**  
   Two PRs in the sharing‑stack: (#26705) refines admin/default‑install labels, (#26704) adds tests for remote catalog browsing, install/uninstall, dedupe, and share‑management flows.

6. **[#28219 / #28189 – Canonicalize tool namespaces](https://github.com/openai/codex/pull/28219)** / **[Namespace client tool search identity](https://github.com/openai/codex/pull/28189)**  
   Two connected PRs that standardise how default tools are namespaced and how the client resolves tool‑search identity. Prepares for cleaner multi‑tool‑provider support.

7. **[#27946 – Use input items for Responses Lite tools](https://github.com/openai/codex/pull/27946)**  
   Migrates Responses Lite to use `additional_tools` and developer items instead of top‑level `tools` / `instructions`. Keeps the interface 1‑to‑1 with the API. Namespacing follow‑up to come.

8. **[#28418 – Remove deprecated AskForApproval::OnFailure](https://github.com/openai/codex/pull/28418)**  
   Deletes the `OnFailure` variant that has been deprecated since PR #11631. Housekeeping that reduces API surface area.

9. **[#28647 – Coordinate MCP OAuth refresh across clients](https://github.com/openai/codex/pull/28647)**  
   Fixes a race condition where multiple Codex clients persist the same MCP OAuth refresh token and collide during provider refresh. Adds coordination to avoid token‑invalidation cascades.

10. **[#28034 – Experimental local credential broker](https://github.com/openai/codex/pull/28034)**  
    First slice of moving injectable credentials behind the managed network proxy so child processes cannot exfiltrate plaintext values. Protects long‑standing credential‑leak vulnerability.

---

## Feature Request Trends

- **Configurable project folder** — Multiple requests (e.g., #19913, #28648) for a default parent‑folder setting to avoid repeated navigation.
- **SSH remote‑workspace support** — #21509 and related discussions want native SSH remote management in Desktop, not just CLI.
- **In‑place Markdown editing** — #28644 asks for editable Markdown preview (specs, PRDs, docs) in the side viewer.
- **Thread automation tools** — #28650 requests exposure of a `automation_update` tool for recurring reminders/monitors within sessions.
- **File viewer improvements** — #28648 requests single‑click file opening (instead of double‑click) and better sidebar interactions.

---

## Developer Pain Points

- **Account lockout / no recovery** — #25749 (30 👍) captures the top frustration: users with valid Google Auth cannot use Codex because an old phone‑number verification step has no fallback.
- **macOS stability regressions** — #25243 (relaunch loop), #27536 (disk leak), #26415 (Computer Use CPU spin) together paint a picture of an Electron app that still leaks processes and disk space on macOS.
- **Windows sandbox friction** — #25436 (runner start failure), #27287/#28121 (Computer Use bootstrap failures), #20567 (git‑process storm) show that Windows is still a second‑class citizen for sandboxed execution.
- **Context‑window exhaustion** — #18052 continues to receive comments; users hitting the model’s context limit without clear guidance or automatic truncation.
- **Stale UI behaviors** — #28095 (Delete button that does nothing), #28643 (unreliable line‑link jumping) degrade trust in the UI’s feedback loop.
- **Performance on paste / input** — #25865 (freeze on JSON paste) and #28652 (30+ second dispatch delay) indicate input‑processing bottlenecks that interrupt flow.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-06-17

## Today’s Highlights

A critical **nightly release failure** (v0.48.0-nightly) is blocking the CI pipeline, while the team continues to prioritise **agent reliability** and **security hardening**. Several long-standing bugs—such as the generalist agent hang and subagent recovery misreporting—remain open with active community engagement, and multiple PRs are advancing OAuth atomicity, MCP header encoding, and cross-server URI isolation.

---

## Releases (last 24h)

No new releases in the past 24 hours.

---

## Hot Issues (Top 10)

1. **[#27973] Nightly Release Failed for v0.48.0-nightly**  
   Priority P0 release-failure. The nightly CI workflow crashed, blocking all nightly builds. No user discussion yet.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/27973)

2. **[#21409] Generalist agent hangs**  
   Priority P1 (bug). The agent hangs indefinitely when deferring to sub-agents; users report waiting up to an hour. Instructing the model not to use sub-agents works around it. High community reaction (8 👍).  
   [Link](https://github.com/google-gemini/gemini-cli/issues/21409)

3. **[#22323] Subagent recovery after MAX_TURNS reported as GOAL success**  
   Priority P1 (bug). The `codebase_investigator` subagent claims success even when it hit the max turn limit before any useful analysis. Misleading status hides true interruption.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/22323)

4. **[#25166] Shell command execution stuck with “Waiting input” after command completes**  
   Priority P1 (bug). After simple CLI commands, Gemini shows the shell as active and awaiting input even though the command finished. Happens with trivial commands (e.g., `ls`).  
   [Link](https://github.com/google-gemini/gemini-cli/issues/25166)

5. **[#24353] Robust component level evaluations**  
   Priority P1 (epic). Follow-up to the behavioral-evals initiative; 76 tests exist for 6 Gemini models. Aims to make component-level evals systematic and reliable.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/24353)

6. **[#22745] Assess impact of AST-aware file reads, search, and mapping**  
   Priority P2 (epic). Investigating whether AST-aware tools (e.g., tilth, glyph) can reduce token usage, misaligned reads, and improve codebase mapping. High developer interest (7 comments).  
   [Link](https://github.com/google-gemini/gemini-cli/issues/22745)

7. **[#26525] Add deterministic redaction and reduce Auto Memory logging**  
   Priority P2 (bug). Auto Memory sends transcript content to the model before redacting secrets; needs deterministic pre-redaction to prevent secret leaks.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/26525)

8. **[#26522] Stop Auto Memory from retrying low-signal sessions indefinitely**  
   Priority P2 (bug). Sessions that the extraction agent decides are low-signal remain unprocessed and are repeatedly surfaced, causing infinite retries.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/26522)

9. **[#21968] Gemini does not use skills and sub-agents enough**  
   Priority P2 (bug). Anecdotal reports that custom skills and sub-agents are rarely invoked unless explicitly instructed, even for clearly related tasks.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/21968)

10. **[#21983] Browser subagent fails in Wayland**  
    Priority P1 (bug). The browser agent crashes or fails to initialise under Wayland environments. Termination reason reported as “GOAL” despite failure.  
    [Link](https://github.com/google-gemini/gemini-cli/issues/21983)

---

## Key PR Progress (Top 10)

1. **[#27966] Enforce case-insensitive sensitive path blocklist and vscode HITL**  
   Security fix: implements a 100% robust blocklist for `.git`, `.env`, `node_modules`, and ensures VSCode human-in-the-loop prompts are enforced for modifications.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27966)

2. **[#27753] Validate workflow_run origin before consuming E2E artifact**  
   CI security: prevents fork artifact poisoning by checking `repo.full_name` and `sha` before running E2E tests with repository secrets.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27753)

3. **[#27971] Strip thoughts from scrubbed history turns to resolve thought leakage**  
   Bug fix: prevents the model’s internal monologue from leaking into plain-text history, which previously caused infinite loop monologues.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27971)

4. **[#27771] Fix MCP header encoding for non-ASCII values**  
   Bug fix: MCP HTTP transport now correctly encodes Unicode header values (e.g., `mąka`) as ByteString, preventing discovery failure.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27771)

5. **[#27664] Write MCP OAuth tokens atomically**  
   Security fix: uses temp file + atomic rename to avoid corruption when writing MCP OAuth token files. Also covers deletion path.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27664)

6. **[#27889] Refresh MCP OAuth with stored client ID**  
   Bug fix: when auto-discovered MCP servers have no explicit `oauth.clientId`, the CLI now uses the persisted client ID from token metadata to refresh.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27889)

7. **[#27643] Fix parallel workspace compilation race condition**  
   Build fix: splits the monorepo build into sequential topological stages (core → libraries → applications) to avoid race conditions in parallel builds.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27643)

8. **[#27948] Pin dependencies and enforce 14-day update cooldown**  
   Dependency management: strips `^`/`~` ranges and enforces a 14-day cooldown for automated dependency updates to reduce supply-chain risk.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27948)

9. **[#27763] Document read_file 20MB file size limit**  
   Documentation: adds explicit documentation for the 20MB limit enforced by `read_file`, addressing user confusion when hitting the error.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27763)

10. **[#27718] Keep auto model visible without preview access**  
    Bug fix: marks the top-level `auto` alias as non-preview so it remains visible in `/model` for users without preview access, while still hiding preview-only aliases.  
    [Link](https://github.com/google-gemini/gemini-cli/pull/27718)

---

## Feature Request Trends

- **AST-Aware Code Understanding**: Multiple issues (e.g., #22745, #22746, #22747) propose using AST tools (tilth, glyph, AST grep) for file reads, search, and codebase mapping to reduce token waste and improve tool-call precision.
- **Auto Memory Improvements**: Epics (#26525, #26522, #26523) focus on deterministic redaction, infinite retry prevention, and quarantining invalid memory patches.
- **Agent Self-Awareness & Skill Usage**: Users want the agent to autonomously leverage custom skills and sub-agents (#21968, #21432) and to understand its own CLI flags, hotkeys, and execution behaviour.
- **Remote Agents & Background Operations**: Epics #20303 and #20195 drive sprint goals for remote sub-agents with advanced auth, 1P agent support, and background processing.
- **Destructive Action Prevention**: Demand for the agent to discourage or prevent destructive commands (e.g., `git reset --force`) with safer fallbacks (#22672).
- **Browser Agent Resilience**: Requests for automatic session takeover, lock recovery (#22232), and Wayland compatibility (#21983).

---

## Developer Pain Points

1. **Agent Hangs and Crashes** — The generalist agent hangs indefinitely (#21409) and the “get-shit-done” output hook causes crashes (#22186). Shell commands remain stuck in “Waiting input” (#25166).
2. **Subagent Misbehaviour** — Subagents misreport success on turn-limit exhaustion (#22323), ignore `settings.json` overrides (#22267), and run without permission after updates (#22093).
3. **Secret Leakage & Indefinite Retries** — Auto Memory sends transcript content to the model before redaction (#26525) and retries low-signal sessions forever (#26522).
4. **Configuration Ignored** — Browser agent and subagents bypass `settings.json` overrides (e.g., `maxTurns`) (#22267).
5. **Platform Incompatibility** — Browser agent fails under Wayland (#21983); terminal resize causes flicker (#21924); external editor exits corrupt the terminal buffer (#24935).
6. **Token Efficiency & Noise** — Model frequently writes tmp scripts in random spots (#23571), uses too many tools (400+ error, #24246), and suffers from thought leakage into history (#27971).
7. **Cross-Server MCP Issues** — MCP resource URI collisions can lead to shadowing of trusted resources (#27964); non-ASCII header values break discovery (#27771).

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-06-17

## Today’s Highlights

A new minor release (v1.0.64-0) lands with several quality-of-life improvements, most notably making `/security-review` available to all users and adding a `/diagnose` command for session log analysis. Meanwhile, community reports are dominated by fallout from yesterday’s June 16 GitHub Copilot outage, with a critical bug showing all models as “Blocked/Disabled” (#3832). Stability issues on Windows ARM64 (#3687) and a regression where sub‑agents lose access to MCP tools (#3812) are also drawing attention.

## Releases

**v1.0.64-0** was published in the last 24 hours with the following additions:

- `/diagnose` command to analyze session logs  
- `/mcp` registry installation for browsing and installing MCP servers  
- `/security-review` is no longer experimental — available to all users  
- Automatic discovery of MCP servers provided by installed plugins  
- CSV output support for MCP tools  

[Full release details](https://github.com/github/copilot-cli/releases/tag/v1.0.64-0)

## Hot Issues

1. **#3832 – All models show as “Blocked/Disabled” after June 16 outage**  
   *Triage* – Users cannot select any model or start new sessions after the GitHub Copilot outage on June 16, 17:45–18:15 UTC.  
   [Issue](https://github.com/github/copilot-cli/issues/3832) | 👍 0

2. **#3687 – `copilot.exe` fatal-aborts under load on Windows ARM64**  
   *area:sessions, area:platform-windows* – Hard abort (BEX64 / 0xc0000409) reproducible across versions 1.0.57–1.0.60, especially when multiple sessions start simultaneously under memory pressure.  
   [Issue](https://github.com/github/copilot-cli/issues/3687) | 👍 1 | 5 comments

3. **#1168 – Excessive authorization prompts during a single request (“authorization fatigue”)**  
   *area:permissions* – A single high‑level prompt can trigger over a dozen consent popups, severely disrupting workflow.  
   [Issue](https://github.com/github/copilot-cli/issues/1168) | 👍 2 | 2 comments

4. **#3828 – `ContentExclusionFilter.isExcluded` crash**  
   *area:non-interactive, area:tools* – The `rg` tool crashes with `TypeError: Cannot read properties of undefined` from the exclusion filter class.  
   [Issue](https://github.com/github/copilot-cli/issues/3828) | 👍 0 | 1 comment

5. **#3821 – Running `/update` from a resumed session leaves conflicting flags**  
   *area:sessions, area:installation* – After `copilot -r`, `/update` restarts with both `--session-id` and `-r/--resume`, preventing normal resumption.  
   [Issue](https://github.com/github/copilot-cli/issues/3821) | 👍 0 | 1 comment

6. **#3812 – Sub‑agents can no longer access MCP tools**  
   *area:agents, area:mcp* – Custom sub‑agents cannot see or use MCP tools, a regression from a previous working state. The issue likely relates to deferred loading of MCP tools.  
   [Issue](https://github.com/github/copilot-cli/issues/3812) | 👍 0 | 1 comment

7. **#3825 – `--allow-all` read permissions leak to UI dispatcher, wedging the TUI**  
   *area:permissions* – Using `--allow-all` with `-i` or `--resume` causes read permission requests to bypass the intended UI, leaving the terminal stuck with no input box.  
   [Issue](https://github.com/github/copilot-cli/issues/3825) | 👍 0 | 0 comments

8. **#3824 – Sub‑agents run a different model than the configured session model**  
   *area:agents, area:models* – Spawned sub‑agents default to a different model because of agent‑type defaults and experiment overrides, with no indication to the user.  
   [Issue](https://github.com/github/copilot-cli/issues/3824) | 👍 0 | 0 comments

9. **#3823 – Reasoning effort `xhigh` silently downgraded to `medium`**  
   *area:models* – When the active model does not support `xhigh`, the CLI falls back to `medium` instead of the nearest clamp (e.g., `max`), causing silent behaviour degradation.  
   [Issue](https://github.com/github/copilot-cli/issues/3823) | 👍 0 | 0 comments

10. **#3830 – Feature request: single command to update all installed plugins**  
    *area:plugins* – Users want a bulk update command to avoid updating plugins one at a time. This request has immediate traction.  
    [Issue](https://github.com/github/copilot-cli/issues/3830) | 👍 0 | 0 comments

## Key PR Progress

No pull requests were merged or updated in the last 24 hours.

## Feature Request Trends

Several recurring feature directions emerge from the latest issue feed:

- **Plugin & MCP management** – Bulk plugin update (#3830), async slash commands for `/mcp show` and `/plugin list` (#3829), and supporting `skillDirectories` at the repository level (#3822) are all highly requested to reduce friction.
- **Enterprise custom models** – Users with GitHub Copilot Enterprise want the same admin‑configured models available in Copilot CLI (#3730).
- **Session lifecycle** – The ability to unarchive/restore archived sessions (#3518) and better error messaging around rate limits (including timezone info, #3819) are common requests.
- **Asynchronous operations** – Making read‑only slash commands non‑blocking (#3829) would align with the already async `/tasks` and improve responsiveness.

## Developer Pain Points

The community is frustrated by several recurring issues that impact productivity:

- **Authorization fatigue** (#1168) – Repeated permission prompts during a single task chain remain a major workflow breaker.
- **Windows stability** (#3687) – The BEX64 crash on ARM64 under load is a hard blocker for some users, with no fix yet.
- **Transient API errors** (#3831) – Retry loops that kill long‑running workflows with no explanation.
- **Silent configuration drops** – Reasoning effort downgrades (#3823) and model switching in sub‑agents (#3824) happen without any warning, eroding trust.
- **TUI wedging** – The `--allow-all` permission leak (#3825) and the cancellation message re‑injection (#3826) both leave the terminal in an unresponsive state.
- **Rate‑limit UX** – The “wait for limit to reset” message lacks a human‑readable timestamp/timezone (#3819), making planning difficult.

These pain points highlight a need for better error handling, user‑facing transparency, and more robust input/window management.

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest – 2026-06-17

**Data source:** [github.com/MoonshotAI/kimi-cli](https://github.com/MoonshotAI/kimi-cli)

---

## 1. Today's Highlights

No new releases were published in the last 24 hours. The community’s attention is focused on two critical bugs: a persisted 400 error caused by the CLI auto‑discovering a deleted MCP server, and a poor onboarding experience where a fresh install fails without any hint to run `kimi login`. Meanwhile, a long‑standing request to raise the default maximum steps per turn remains open, and a closed feature request to hide thinking content shows continued interest in output customization.

---

## 2. Releases

_No new versions were released in the last 24 hours._

---

## 3. Hot Issues

### #1327 – [enhancement] More Steps per turn By Default
- **Author:** sssxks · **Created:** 2026-03-03 · **Updated:** 2026-06-16  
- **Comments:** 3 · 👍: 0  
- **Link:** [Issue #1327](https://github.com/MoonshotAI/kimi-cli/issues/1327)

**Why it matters:** The default limit of 100 steps is frequently hit even when context usage is low (e.g., 34.5%). Users can manually raise the limit, but the low default causes unnecessary interruptions for long‑running tasks. The community has not yet reached a strong consensus (no thumbs‑up), but the issue has been open for over three months, indicating many developers hit this wall.

---

### #1632 – [CLOSED] Feature Request: Option to hide thinking content while using thinking models
- **Author:** yuantianyu177 · **Created:** 2026-03-29 · **Updated:** 2026-06-16  
- **Comments:** 2 · 👍: 3  
- **Link:** [Issue #1632](https://github.com/MoonshotAI/kimi-cli/issues/1632)

**Why it matters:** Users of thinking models (e.g., `kimi-k2-thinking-turbo`) want the reasoning visible during development but prefer a clean output for production or presentation. The request was closed, but the 3 thumbs‑up show moderate interest. The community may benefit from a simple configuration toggle.

---

### #2457 – [bug] Kimi Code CLI auto‑discovers MCP server after user deleted it, causing unfixable 400 errors
- **Author:** xavier2sy8827-cmyk · **Created:** 2026-06-16 · **Updated:** 2026-06-16  
- **Comments:** 0 · 👍: 0  
- **Link:** [Issue #2457](https://github.com/MoonshotAI/kimi-cli/issues/2457)

**Why it matters:** This is a **critical bug** reported on Windows 10 with the `kimi-code` platform. Deleting an MCP server should remove it from the CLI’s configuration, but the tool continues to auto‑discover it, leading to persistent 400 errors that cannot be cleared by the user. No workaround is currently suggested. This issue was filed just yesterday and already has high visibility.

---

### #2456 – Bug: Fresh install reports "LLM not set" with no guidance to run login
- **Author:** lming112 · **Created:** 2026-06-16 · **Updated:** 2026-06-16  
- **Comments:** 0 · 👍: 0  
- **Link:** [Issue #2456](https://github.com/MoonshotAI/kimi-cli/issues/2456)

**Why it matters:** A fresh Homebrew install (`brew install kimi-cli`) fails on any command (e.g., `kimi --print`) with the cryptic error `LLM not set`. No hint that `kimi login` is required. This creates a poor first‑impression for new users and wastes debugging time. The issue is brand new and has not yet received official response.

---

## 4. Key PR Progress

### #1771 – [fix] always stringify tool message content in Chat Completions provider
- **Author:** he‑yufeng · **Created:** 2026-04-06 · **Updated:** 2026-06-16  
- **Comments:** 0 · 👍: 0  
- **Link:** [PR #1771](https://github.com/MoonshotAI/kimi-cli/pull/1771)

**What it fixes:** Resolves Issue #1762 where the OpenAI Chat Completions API rejected tool results containing multiple `ContentPart` objects. The fix forces `content` to be a plain string for `role: "tool"` messages, eliminating the 400 “Failed to parse tool message” error. This is a small but important correctness improvement for providers that rely on the standard OpenAI API shape.

---

## 5. Feature Request Trends

Based on the latest issues, two distinct feature directions are emerging:

1. **Higher default step limit** – Users want the CLI to allow more processing steps before hitting the “Max number of steps reached” wall, especially when context is not saturated.
2. **Configurable thinking output** – Developers want the ability to hide or show the reasoning trace of thinking models, switching between a clean terminal view and a detailed debug view.

Both trends point to a need for **more user‑controllable defaults** and **output verbosity toggles**.

---

## 6. Developer Pain Points

- **Auto‑discovery of deleted resources** – The CLI caches MCP servers aggressively and continues to probe them even after the user has explicitly removed them, leading to unrecoverable error states.
- **Poor first‑run experience** – No auto‑detection or suggestion to run `kimi login` when no LLM is configured, causing immediate friction for new users.
- **Low default step limit** – The 100‑step default is too conservative for real‑world use, forcing developers to dig into configuration files to adjust it.

These pain points affect both new adopters and experienced users and are likely to increase friction in the developer workflow until addressed.

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-17

## Today’s Highlights
No new releases landed in the last 24 hours, but community attention is concentrated on a long-standing hang bug (#2940), a clipboard copy failure (#7048), and a sudden regression in image reading (#25832). On the pull request side, several critical fixes are moving forward: PowerShell UTF-8 support on Windows, improved desktop file‑watcher safety, and a workaround for MiniMax’s rejection of tool‑heavy sessions.

---

## Releases
No releases in the last 24 hours.

---

## Hot Issues

### 1. [BUG] OpenCode hangs randomly after receiving instructions  
**#2940** — [anomalyco/opencode#2940](https://github.com/anomalyco/opencode/issues/2940)  
*39 comments · 18 👍*  
The most‑voted open bug. Users report that OpenCode freezes regardless of model, especially in Laravel projects. `/compact` only sometimes helps, and `/exit` is often unresponsive. The long thread suggests a deep scheduler or process‑management issue.

### 2. Copy Text "Copied to clipboard" does never work  
**#7048** — [anomalyco/opencode#7048](https://github.com/anomalyco/opencode/issues/7048)  
*23 comments · 13 👍*  
On Ubuntu Desktop (GhostTTY), right‑clicking to copy text in the output/input window shows a confirmation but the clipboard is never updated. High‑severity UX bug reported six months ago, still open.

### 3. opencode cannot read images anymore  
**#25832** — [anomalyco/opencode#25832](https://github.com/anomalyco/opencode/issues/25832)  
*13 comments · 4 👍*  
After April 29, 2026, image input via PNG/JPG stopped working — the tool returns `err: Bad ...`. Users relied on this for HTML modifications. Likely an API change in the image‑processing pipeline.

### 4. OpenCode is heavily CPU‑bound  
**#21470** — [anomalyco/opencode#21470](https://github.com/anomalyco/opencode/issues/21470)  
*11 comments · 10 👍*  
With Gemini‑3.1, over 1.5M tokens are consumed inside OpenCode itself rather than by the model API. Performance profiling shows the tool is spending most time on internal processing, not waiting on external calls.

### 5. Skills don’t show up in TUI autocomplete (closed with fix)  
**#22129** — [anomalyco/opencode#22129](https://github.com/anomalyco/opencode/issues/22129)  
*Closed · 10 comments · 12 👍*  
Skills appear correctly in the web app but are absent from TUI slash‑commands. Root cause located in `prompt/autocomplete.tsx:363`. A high‑interest fix has been identified.

### 6. IDE (VSCode): Context Awareness function didn't take effect  
**#22235** — [anomalyco/opencode#22235](https://github.com/anomalyco/opencode/issues/22235)  
*7 comments · 4 👍*  
The “Context Awareness” feature (similar to Claude Code’s auto‑attach) never activates in VSCode. Users ask whether prerequisite settings are needed.

### 7. Move project folder to path B and delete old path A ... OpenCode still navigates to old path  
**#30697** — [anomalyco/opencode#30697](https://github.com/anomalyco/opencode/issues/30697)  
*4 comments*  
On Windows, after moving a project folder, OpenCode Desktop persists the old path and fails to open the new location. Path‑cache invalidation bug.

### 8. DeepSeek model: edit tool frequently fails to invoke  
**#31849** — [anomalyco/opencode#31849](https://github.com/anomalyco/opencode/issues/31849)  
*4 comments · 1 👍*  
On OpenCode 1.17.0, the edit tool for code modifications fails intermittently when using DeepSeek. Provider‑specific tool routing issue.

### 9. Infinite clarification/compaction loop on empty git repo  
**#32615** — [anomalyco/opencode#32615](https://github.com/anomalyco/opencode/issues/32615)  
*3 comments*  
OpenCode can enter a token‑burning loop when opened on a directory with only `.git/`. Both a correctness and cost‑control bug.

### 10. GET /session/:id/message returns 400 on inline JSON schema  
**#26929** — [anomalyco/opencode#26929](https://github.com/anomalyco/opencode/issues/26929)  
*3 comments · 5 👍*  
Messages containing `format` with a JSON schema cause a 400 when fetched. Affects structured output workflows.

---

## Key PR Progress

### 1. fix(shell): add PowerShell UTF-8 command wrapper on Windows  
**#31985** — [anomalyco/opencode#31985](https://github.com/anomalyco/opencode/pull/31985)  
Closes five related issues, supersedes an earlier attempt. Ensures commands executed via PowerShell handle UTF‑8 correctly — a long‑standing Windows pain point.

### 2. fix: OpenAI-compatible provider improvements  
**#23501** — [anomalyco/opencode#23501](https://github.com/anomalyco/opencode/pull/23501)  
Fixes system message handling, image support, and stream interruption for Ollama/local models. Consolidates three earlier PRs into one.

### 3. feat(task): Add subagent-to-subagent delegation with budgets, persistent sessions  
**#7756** — [anomalyco/opencode#7756](https://github.com/anomalyco/opencode/pull/7756)  
Merged (closed). A massive feature enabling hierarchical agent workflows with token budgets and persistent sub‑sessions. Closes three related issues.

### 4. fix(desktop): skip file watcher on $HOME and filesystem root  
**#32610** — [anomalyco/opencode#32610](https://github.com/anomalyco/opencode/pull/32610)  
Desktop was watching entire home directories, causing inotify timeouts and CPU spikes. Now skips broad roots and sanitizes persisted project lists.

### 5. fix(opencode): sanitize OpenAI MCP tool schemas  
**#32489** — [anomalyco/opencode#32489](https://github.com/anomalyco/opencode/pull/32489)  
MCP servers sometimes expose tuple‑style `items` that OpenAI rejects. PR sanitizes schemas to a compatible subset.

### 6. fix(session): preserve reasoning part type on model switch  
**#32604** — [anomalyco/opencode#32604](https://github.com/anomalyco/opencode/pull/32604)  
Switching models invalidates prefix cache, forcing long re‑processing. This PR preserves the reasoning part type to reduce delays.

### 7. fix(provider): stub orphan MiniMax tool results  
**#32609** — [anomalyco/opencode#32609](https://github.com/anomalyco/opencode/pull/32609)  
MiniMax rejects sessions with prior tool‑call history. This PR replaces missing tool results with stubs to keep the message sequence valid.

### 8. feat(opencode): local LAN provider discovery + auto‑discover models  
**#27554** — [anomalyco/opencode#27554](https://github.com/anomalyco/opencode/pull/27554)  
Adds mDNS‑based discovery of local OpenAI‑compatible servers and auto‑populates model lists. A long‑requested feature for team environments.

### 9. fix(opencode): send system context as structured messages on OpenAI OAuth path  
**#32592** — [anomalyco/opencode#32592](https://github.com/anomalyco/opencode/pull/32592)  
Closes #32505 — the OAuth path was flattening system context into a single instructions string instead of structured messages, causing compatibility issues.

### 10. feat(session): add configurable fallback model chain  
**#27939** — [anomalyco/opencode#27939](https://github.com/anomalyco/opencode/pull/27939)  
Allows users to define an ordered list of fallback models when the primary model fails (e.g., rate‑limited). Closes #7602.

---

## Feature Request Trends

- **Skills improvements** are the most requested: recursive discovery, multi‑skill selection in TUI (#21495), and better metadata persistence (#31616).  
- **Provider/model expansion**: users want native support for `glm‑5.2:cloud` in Ollama (#32620), token‑plan integrations (#32623), and model‑specific tool routing (#32626).  
- **Token and cost monitoring**: several requests for a desktop‑side token monitor (#32619) and better visibility into per‑model spending.  
- **Plugin extensibility**: support for specifying which model executes a plugin‑registered tool (#32626) and deferred MCP tool exposure (#32621) to reduce token overhead.

---

## Developer Pain Points

- **Random hangs and freezes** — #2940 remains the most upvoted open bug, with users forced to Ctrl‑C often.  
- **Clipboard/copy issues** — #7048 shows that basic text selection and copy are broken in the TUI, severely hindering daily use.  
- **Performance bloat** — #21470 highlights that OpenCode itself consumes more CPU and tokens than external model calls, especially with Gemini providers.  
- **Image input regression** — #25832 broke a key workflow for visual HTML editing, with no clear timeline for fix.  
- **Model‑specific failures** — MiniMax (#32608, #32611, #32614, #32617), DeepSeek (#31849), and Gemini (#32625, #11286) each require custom workarounds, increasing maintenance friction.  
- **Persistent path/config bugs** — stale project paths (#30697), disabled_providers not working (#32528), and permissions lost on page refresh (#20998) erode trust.  
- **RTL/Arabic rendering** — #32602 shows the TUI still lacks proper bidirectional text support.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

Here is the Pi community digest for 2026-06-17.

---

## Pi Community Digest — 2026-06-17

### Today's Highlights
The team shipped two fast-follow releases (v0.79.5 and v0.79.6) that resolve critical DeepSeek V4 thinking-parameter conflicts and preserve user-level `fetch` overrides, a direct response to community reports. A wave of DeepSeek V4 compatibility issues dominated the day, with multiple providers hitting 400 errors on tool calls and thinking/effort parameter conflicts. On the operations side, a major Nix flake PR landed, and the team is steadily closing issues around hang/freeze bugs and opaque HTTP error bodies.

### Releases
Two releases rolled out in the last 24 hours:
- **v0.79.6**: Fixes two regressions—HTTP dispatcher no longer reinstalls the global `undici` fetch over a caller’s deliberate override, and inherited OpenCode Go DeepSeek V4 thinking-off requests now properly send `thinking: { type: "disabled" }`.
- **v0.79.5**: Introduces **provider-scoped API key environments**, allowing `auth.json` entries to carry `env` overrides for Cloudflare, Azure OpenAI, Google Vertex, Amazon Bedrock, cache retention, and proxy settings per provider.

### Hot Issues (10 Notable)
1. **#4877 – Session folder collision** (Open, 19 comments)  
   Distinct paths like `/a/b/c/d` and `/a-b/c-d` can map to the same session folder. Community notes this is a design limiter for long-lived or complex projects.
   - [GitHub Issue #4877](https://github.com/earendil-works/pi/issues/4877)

2. **#5696 – Model name does not refresh in TUI on Ctrl+P** (Closed, 9 comments)  
   Switching models via keyboard shortcut sometimes lags by two positions. User mixed frustration with reproducible steps—a high-priority UX fix.
   - [GitHub Issue #5696](https://github.com/earendil-works/pi/issues/5696)

3. **#5687 – `pi list` and `pi update` hang forever when an extension runs an MCP server** (Closed, 8 comments)  
   Package subcommands never exit if an installed extension keeps a long-lived MCP server alive. Affects CI/CD pipelines.
   - [GitHub Issue #5687](https://github.com/earendil-works/pi/issues/5687)

4. **#5816 – Tool `search` not found after update** (Closed, 7 comments)  
   Reported on v0.79.4; the agent tries to use a `search` tool that no longer exists or fails to load. Multiple users chimed in with same experience.
   - [GitHub Issue #5816](https://github.com/earendil-works/pi/issues/5816)

5. **#5790 – Support `httpProxy` in settings.json** (Closed, 7 comments)  
   Want to route Pi through a proxy without environment variables. Community showed strong demand for declarative proxy configuration.
   - [GitHub Issue #5790](https://github.com/earendil-works/pi/issues/5790)

6. **#5571 – `pi -p` hangs indefinitely when stdin is a non-TTY pipe** (Closed, 7 comments)  
   Agent freezes (3+ minutes) on fresh installs when the default provider has no credentials instead of failing fast.
   - [GitHub Issue #5571](https://github.com/earendil-works/pi/issues/5571)

7. **#5728 – Provider-specific config in auth.json** (Closed, 7 comments)  
   Users need more than an API key for providers like Cloudflare AI Gateway (accountId, gatewayId). This issue drove the v0.79.5 feature.
   - [GitHub Issue #5728](https://github.com/earendil-works/pi/issues/5728)

8. **#5700 – Support multiple live agent sessions with TUI switching** (Open, 5 comments)  
   Cannot run a background agent while working on another. Community sees this as the next big productivity leap.
   - [GitHub Issue #5700](https://github.com/earendil-works/pi/issues/5700)

9. **#5763 – Providers swallow HTTP error bodies** (Open, 4 comments)  
   Behind gateways, error responses (403, etc.) show opaque messages like "Unknown: UnknownError". Debugging is nearly impossible.
   - [GitHub Issue #5763](https://github.com/earendil-works/pi/issues/5763)

10. **#5825 – Streaming markdown forces scroll to bottom** (Open, 3 comments)  
    When `clear on shrink` is enabled, reading ahead is broken by forced scroll resets—a persistent annoyance in the TUI.
    - [GitHub Issue #5825](https://github.com/earendil-works/pi/issues/5825)

### Key PR Progress (10 Important)
1. **#5812 – fix(tui): protect pipe characters inside inline code in markdown tables** (Closed)  
   Prevents markdown renderer from splitting table cells on pipes inside backticks. Low-risk, high-impact rendering fix.
   - [GitHub PR #5812](https://github.com/earendil-works/pi/pull/5812)

2. **#5820 – fix: Preserve raw HTTP error status and bodies for non-schema errors** (Closed)  
   Implements a shared error helper to surface raw HTTP status and body across all providers. Addresses #5763.
   - [GitHub PR #5820](https://github.com/earendil-works/pi/pull/5820)

3. **#5807 – feat: add provider-scoped environment overrides** (Closed)  
   Auth.json `env` object now overrides process env per provider. Core infrastructure for the v0.79.5 release.
   - [GitHub PR #5807](https://github.com/earendil-works/pi/pull/5807)

4. **#5809 – feat(ai): add durationMs and timeToFirstTokenMs to Usage, display tokens/sec in footer** (Closed)  
   Adds latency and throughput metrics, enabling footer extensions to show tokens/sec.
   - [GitHub PR #5809](https://github.com/earendil-works/pi/pull/5809)

5. **#5789 – fix(tui): restore cursorUp line-start jump before history browsing** (Closed)  
   Fixes a regression where pressing Up on the empty input would incorrectly enter history navigation.
   - [GitHub PR #5789](https://github.com/earendil-works/pi/pull/5789)

6. **#5803 – fix(ai): reject malformed OpenAI tool calls** (Closed)  
   Streamed tool calls without `id` or function name are now rejected before persisting to session history.
   - [GitHub PR #5803](https://github.com/earendil-works/pi/pull/5803)

7. **#5801 – Nixify pi** (Closed)  
   Adds Nix flake packaging for Pi. Useful for dev-ops workflows and reproducible environments.
   - [GitHub PR #5801](https://github.com/earendil-works/pi/pull/5801)

8. **#5798 – feat(coding-agent): add Vercel AI Gateway attribution** (Closed)  
   Adds `http-referer` and `x-title` headers for Vercel AI Gateway identification.
   - [GitHub PR #5798](https://github.com/earendil-works/pi/pull/5798)

9. **#5796 – chore: bump TS target and lib to ES2024, use Promise.withResolvers()** (Open)  
   Replaces hand-rolled `Promise.withResolvers()` implementations with native ES2024 API. Clean tech-debt reduction.
   - [GitHub PR #5796](https://github.com/earendil-works/pi/pull/5796)

10. **#5818 – Deepseek 4 over opencode: thinking and reasoning_effort conflict** (Closed)  
    Hotfix preventing both `thinking` and `reasoning_effort` parameters from being sent together. Directly tied to v0.79.6 release.
    - [GitHub Issue #5818](https://github.com/earendil-works/pi/issues/5818) (linked from PR)

### Feature Request Trends
- **Multi-session & switching**: The #1 most-discussed feature request is ability to run multiple concurrent agent sessions with TUI toggling (#5700).
- **RPC & programmatic access**: Growing demand for `get_entries` / `get_tree` RPC endpoints to drive Pi from external tools (#5810).
- **Proxy & gateway configuration**: Users want declarative proxy (`httpProxy` in settings.json) and custom OAuth callback pages (#5372).
- **DeepSeek V4 compatibility**: Multiple issues (thinking/reasoning_effort, tool call serialization) indicate users need stable, documented DeepSeek V4 support.

### Developer Pain Points
- **DeepSeek V4 parameter conflicts**: High-frequency frustration with contradictory parameters (`thinking` vs `reasoning_effort`) and tool chain serialization errors. The team fixed this in v0.79.6, but users hit it repeatedly.
- **Hangs and freezes**: Four distinct issues (#5687, #5571, #5778, #5805) involve Pi hanging on non-TTY pipes, long-lived MCP extensions, or Windows updates. A systemic reliability concern.
- **Terminal/UI quirks**: Double key events in Kitty, scroll jumps during streaming, and broken URL rendering in Warp point to incomplete terminal emulator support.
- **Error message opacity**: Swallowed HTTP error bodies (#5763) make debugging proxy/gateway failures nearly impossible—partially fixed in PR #5820.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

Here is the **Qwen Code Community Digest** for **2026-06-17**, based on the provided GitHub data.

---

## Qwen Code Community Digest — 2026-06-17

### Today's Highlights
Three release workflows failed today (v0.18.1-preview.1 & nightly), blocking the build pipeline and highlighting a fragile CI process. Meanwhile, the community is closely tracking a stubborn OOM bug in auto-memory (#5147) and a new feature request for automatic model switching (#5225). On the PR front, a critical fix for ACP daemon sessions (#5216) and the introduction of a "vision bridge" feature (#5126) stand out as major developments.

### Releases
No new releases were published in the last 24 hours. However, three separate release workflows failed: two for `v0.18.1-preview.1` (#5222, #5215) and one for a nightly build (#5214). This is a significant blocker for the team.

### Hot Issues
1.  **OOM after `/quit` with auto-memory** [#5147](https://github.com/QwenLM/qwen-code/issues/5147) — [OPEN/P2] A critical memory leak persists in managed auto-memory even after a short session. Prior fixes for full-history OOM are confirmed present, but the issue now points to the `GeminiClient.runManagedAuto...` background task. High scrutiny and active debugging.

2.  **Stale marker blocks worktree cleanup** [#5208](https://github.com/QwenLM/qwen-code/issues/5208) — [OPEN/P2] A `.qwen-session` marker file from a previous session causes the `exit_worktree` tool to refuse cleanup, breaking normal user workflows. The community is concerned about session isolation.

3.  **Auto-update broken on old glibc (CentOS 7)** [#5206](https://github.com/QwenLM/qwen-code/issues/5206) — [CLOSED] This was a high-impact bug for Linux users. The auto-updater silently migrated npm installations to a standalone installer with a newer Node.js, breaking on systems with glibc 2.17. It was resolved in PR #5207, but the risk remains for other edge cases.

4.  **Stuck in ExitPlanMode** [#5210](https://github.com/QwenLM/qwen-code/issues/5210) — [OPEN/P2] A user reports being stuck in "ExitPlanMode" for over 7 hours when using the qwen3.7-max model. This is a severe UX regression for users heavy on plan-then-yolo workflows.

5.  **Terminal stuck in SGR mouse mode after exit** [#5212](https://github.com/QwenLM/qwen-code/issues/5212) — [CLOSED] A significant annoyance: crashes or normal exits could leave the terminal in an unusable state. The root cause was identified in `useMouseEvents.ts`. Community welcome for PR contributors.

6.  **`/model` lists discontinued OAuth coder-model** [#5160](https://github.com/QwenLM/qwen-code/issues/5160) — [CLOSED] A small but confusing UX bug for users not using OAuth. The CLI was showing an inaccessible, discontinued model. Quick fix and merge.

7.  **Release pipeline failures** [#5222](https://github.com/QwenLM/qwen-code/issues/5222), [#5215](https://github.com/QwenLM/qwen-code/issues/5215), [#5214](https://github.com/QwenLM/qwen-code/issues/5214) — [ALL OPEN] Three automated release pipeline failures today. This is a major process-level issue blocking all new version distribution. The community is watching for updates.

8.  **Integration tests not running on PRs** [#5219](https://github.com/QwenLM/qwen-code/issues/5219) — [OPEN/P2] A critical CI blind spot. E2e tests only run on nightly releases, meaning regressions are hidden until they block a release. This is directly responsible for the current release failures.

9.  **`exit_plan_mode` fails with empty plan** [#5177](https://github.com/QwenLM/qwen-code/issues/5177) — [CLOSED] The model occasionally calls `exit_plan_mode` with an empty plan, causing a tool error and consuming a retry turn. A minor but costly bug for users.

10. **Localize hardcoded English UI strings** [#5186](https://github.com/QwenLM/qwen-code/issues/5186) — [CLOSED] A community-driven improvement exposed several undocumented English strings in `web-shell`. This was quickly resolved with PR #5189. Positive community engagement.

### Key PR Progress
1.  **feat(vision-bridge): Transcribe images to text** [#5126](https://github.com/QwenLM/qwen-code/pull/5126) — [OPEN] This is a major feature introducing an opt-in "vision bridge." It allows text-only models to handle images by sending them to a multimodal model for transcription. A game-changer for users who prefer text-only models.

2.  **fix(acp): Load extension commands in daemon sessions** [#5216](https://github.com/QwenLM/qwen-code/pull/5216) — [OPEN] Critical fix for the Agent Communication Protocol (ACP). Daemon sessions were loading with empty extensions, nullifying all extension functionality. This PR restores the normal loading behavior.

3.  **ci: Run CLI integration tests in the merge queue** [#5224](https://github.com/QwenLM/qwen-code/pull/5224) — [OPEN] Addresses the core CI problem (Issue #5219) by moving e2e tests into the merge queue. This is the most important process fix in the queue today.

4.  **fix(cli): Stop after cancelled ask_user_question** [#5218](https://github.com/QwenLM/qwen-code/pull/5218) — [OPEN] Improves tool execution reliability when a user cancels an ACP tool. Stops subsequent tools from running incorrectly and preserves necessary history for replay.

5.  **feat(i18n): Localize tool display names in TUI and web-shell** [#5220](https://github.com/QwenLM/qwen-code/pull/5220) — [OPEN] A significant UX improvement for non-English users. Tool-call badges like "TodoWrite" or "Shell" will now be translated. Builds on recent i18n efforts.

6.  **feat(channel): Add QQ Bot channel adapter** [#5202](https://github.com/QwenLM/qwen-code/pull/5202) — [OPEN] A new platform integration adding support for the QQ Bot (QQ机器人) ecosystem, joining Telegram, WeChat, DingTalk, and Feishu. Broadens the platform reach significantly.

7.  **fix(cli): Keep sudo-required npm installs on npm** [#5207](https://github.com/QwenLM/qwen-code/pull/5207) — [CLOSED] The solution to Issue #5206. Prevents auto-update from silently migrating users to a potentially incompatible standalone installer. A crucial fix for Linux users.

8.  **fix(e2e): Add daemon_status to serve capabilities baseline** [#5211](https://github.com/QwenLM/qwen-code/pull/5211) — [CLOSED] A quick fix to align integration test expectations with actual features. Intended to unblock the release pipeline, but the pipeline ran into other issues.

9.  **fix(core): Fall back to encrypted-file storage for extension secrets** [#5221](https://github.com/QwenLM/qwen-code/pull/5221) — [OPEN] A robustness improvement for extension secret management. Falls back to AES-256-GCM encrypted file storage when the OS keychain is unavailable. Good for security and portability.

10. **feat(loop): Wire prompt-only /loop to self-paced wakeups** [#5197](https://github.com/QwenLM/qwen-code/pull/5197) — [OPEN] Part 2 of aligning `/loop` with Claude Code. Allows users to run a prompt at most once on a schedule, rather than a fixed recurring cron. Requires a new wakeup engine (#5182).

### Feature Request Trends
- **Automatic Model Selection**: A strong desire for an automated mechanism to switch between "pro" and "flash" model tiers based on task complexity to save costs (#5225). This is already a feature in competing agent software.
- **Internationalization (i18n)**: The community is demanding a fully localized experience. Recent work on hardcoded strings (#5186) and tool display names (#5220) shows this is a top priority.
- **Platform & Channel Expansion**: The addition of a QQ Bot adapter (#5202) is the latest in a trend of broadening platform support beyond core terminals and typical chat apps.

### Developer Pain Points
1.  **Fragile CI/CD Pipeline**: The core of today's issues. The release pipeline is unstable, and integration tests are not run early enough (only on nightly cron). This leads to regressions and blocked releases, causing significant friction for maintainers and users waiting for fixes.
2.  **High-Severity, Low-Frequency Bugs**: Bugs like the auto-update glibc issue and the sticky SGR mouse mode are hard to test for but cause major user disruption. The community is sensitive to these "corner case" failures.
3.  **Worktree & Session Management**: Issues like the stale `.qwen-session` marker highlight pain points in session isolation and cleanup. Users expect to be able to start fresh in their projects without manual cleanup.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-06-17

## 1. Today’s Highlights

The project continues its rebrand to **CodeWhale** (npm package `codewhale`, canonical name `CodeWhale`), with v0.8.61 as the latest release and legacy `deepseek-tui` deprecated. The community is actively reporting stability issues in production use, particularly around task stalls and sub‑agent deadlocks, while a major v0.9.0 EPIC for chat‑native workrooms (PR #3277) landed its foundation layer. Multiple quality‑of‑life PRs and documentation fixes also shipped today, including static musl Linux builds and improved paste handling.

## 2. Releases

- **v0.8.61** (latest, last 24h)  
  Official release under the `CodeWhale` name. The legacy npm package `deepseek-tui` is deprecated and receives no further updates. Migration guide available at `docs/REBRAND.md`.  
  [Release details](https://github.com/Hmbown/CodeWhale/releases/tag/v0.8.61)

No other new releases in the period.

## 3. Hot Issues (10 noteworthy)

1. **#2487 – Frequent “Turn stalled – no completion signal received”**  
   *Author: yahayao* | 14 comments | 1 👍  
   Core stability problem: `yolo` mode freezes and becomes unresponsive. Despite sending `continue`, the operation never resumes. High impact for heavy users.  
   [Issue #2487](https://github.com/Hmbown/CodeWhale/issues/2487)

2. **#2739 – Task execution still freezes (Chinese report)**  
   *Author: zoomtint* | 4 comments  
   Long‑running tasks hang indefinitely; `Esc` shows connection timeout. Past fix (v0.8.52) didn’t fully resolve. User says “unbearable” and considered abandoning the tool.  
   [Issue #2739](https://github.com/Hmbown/CodeWhale/issues/2739)

3. **#3209 – v0.9.0 EPIC: Chat‑native CodeWhale workrooms**  
   *Author: Hmbown* | 2 comments  
   Ambitious feature request for threaded, shareable agent work – channels, mentions, mobile access. Core of the next major version.  
   [Issue #3209](https://github.com/Hmbown/CodeWhale/issues/3209)

4. **#3275 – CodeWhale overly self‑questioning and deviating from user intent**  
   *Author: yekern* | 1 comment  
   Regression from #3061: agents extend scope without confirmation, entering self‑driven loops. Directly impacts trust.  
   [Issue #3275](https://github.com/Hmbown/CodeWhale/issues/3275)

5. **#3264 – Restrict skill scanning to `~/.codewhale/skills/`**  
   *Author: Yunqing2022* | 3 comments  
   Proposal to limit skill discovery to a single directory for better control and security.  
   [Issue #3264](https://github.com/Hmbown/CodeWhale/issues/3264)

6. **#3240 – Legacy deepseek configuration directory still created**  
   *Author: Final527* | 2 comments  
   Despite rebrand, the program still creates `.deepseek` folder alongside `.codewhale` on Windows – a cleanup gap.  
   [Issue #3240](https://github.com/Hmbown/CodeWhale/issues/3240)

7. **#3238 – Fails on Ubuntu 22.04 due to glibc version mismatch**  
   *Author: thahmidul-islam-nafi* | 2 comments  
   `npm install -g codewhale` fails because pre‑built binary requires newer glibc. Affects many LTS users.  
   [Issue #3238](https://github.com/Hmbown/CodeWhale/issues/3238)

8. **#3273 – `js_execution` Node fetch doesn’t honour proxy config on Windows**  
   *Author: lordwedggie* | 1 comment  
   Shell tools work behind proxy, but built‑in JS execution tool times out. Environment variable issue.  
   [Issue #3273](https://github.com/Hmbown/CodeWhale/issues/3273)

9. **#3266 – `agent_eval` with `block=True` causes TUI freeze with multiple sub‑agents**  
   *Author: giovanni-paolilla* | 2 comments  
   Parent session deadlocks when ≥2 sub‑agents run. Force‑kill required. High severity for multi‑agent workflows.  
   [Issue #3266](https://github.com/Hmbown/CodeWhale/issues/3266)

10. **#2870 – EPIC: staged command‑boundary refactor for v0.9.0**  
    *Author: aboimpinto* | 4 comments  
    Tracks the architecture overhaul to support smaller mergeable PRs. Reference PR #2851.  
    [Issue #2870](https://github.com/Hmbown/CodeWhale/issues/2870)

## 4. Key PR Progress (9 PRs in last 24h)

1. **#3277 – Workrooms Phase 1 (data model, endpoints, docs, tool)**  
   *Author: idling11* | Foundation for v0.9.0 chat‑native workrooms. RFC, data model, WorkroomLink URL format, runtime mapping.  
   [PR #3277](https://github.com/Hmbown/CodeWhale/pull/3277)

2. **#3271 – Docs: add Ponytail personality to project instructions**  
   *Author: ousamabenyounes* | Blocked on upstream dependency; adds CodeWhale as supported agent.  
   [PR #3271](https://github.com/Hmbown/CodeWhale/pull/3271)

3. **#3269 – feat(tui): expose slash commands as hotbar actions**  
   *Author: reidliu41* | Refs #2067. Allows binding `slash.mode`, `slash.task`, etc. as hotbar actions.  
   [PR #3269](https://github.com/Hmbown/CodeWhale/pull/3269)

4. **#2998 – chore(deps-dev): bump tailwindcss from 3.4.19 to 4.3.1**  
   *Author: dependabot[bot]* | Major version bump for the web marketing site.  
   [PR #2998](https://github.com/Hmbown/CodeWhale/pull/2998)

5. **#3270 – docs: add Linux build‑time deps to cargo install guides**  
   *Author: zlh124* | Fixes `cargo install` failure on bare Ubuntu 24.04 by documenting `libdbus-1-dev` and `pkg-config`.  
   [PR #3270](https://github.com/Hmbown/CodeWhale/pull/3270)

6. **#3274 – feat(release): build static Linux x64 binaries with musl**  
   *Author: wavezhang* | Switches release builds to musl for broader glibc compatibility. Counterpart to #2903 (CNB pipeline).  
   [PR #3274](https://github.com/Hmbown/CodeWhale/pull/3274)

7. **#3236 – Add DeepInfra provider support**  
   *Author: nightt5879* | Fixes #3231. Includes runtime/TUI/CLI/TOML wiring and docs.  
   [PR #3236](https://github.com/Hmbown/CodeWhale/pull/3236)

8. **#3267 – feat(tui): keep oversized paste inline with truncation & auto‑expand**  
   *Author: idling11* | Addresses #3263: no longer replaces large paste with `@file` mention – keeps text editable.  
   [PR #3267](https://github.com/Hmbown/CodeWhale/pull/3267)

9. **#2933 – feat(hippocampal): v2 memory system (glossary, namespaces, rollback, daemon)**  
   *Author: cy2311* | Upgrades hippocampal memory to a full cross‑session layer. Schema migration, auto‑inject, background daemon. Needs human review.  
   [PR #2933](https://github.com/Hmbown/CodeWhale/pull/2933)

## 5. Feature Request Trends

- **Chat‑native workrooms** (#3209, PR #3277) – A major architectural push toward threaded, shareable agent sessions with channels, mentions, and mobile access. Likely dominates v0.9.0 roadmap.
- **Model metadata registry** (#3071, #3072, #3073) – In flight: replace hard‑coded model lists with a hydrated, cacheable registry from provider APIs.
- **Clarification questions as first‑class UI** (#3102) – Agents should ask for input through modal or highlighted prompts, not just chat messages.
- **Skill discovery scoping** (#3264) – Request to limit scanning to `~/.codewhale/skills/` for security and predictability.
- **Memory v2** (PR #2933) – Cross‑session memory with namespaces, rollback, and background daemon – a significant enhancement for long‑running agent collaboration.

## 6. Developer Pain Points

- **Task stalls / freezes** (#2487, #2739, #3266) – Production blockers: yolo mode, long tasks, and multi‑sub‑agent orchestration all hang or deadlock. Users frustrated despite previous fixes.
- **Rebrand cleanup gaps** (#3240) – Legacy configuration directory still created, confusing users on Windows.
- **Glibc incompatibility** (#3238) – Pre‑built binary requires newer glibc than Ubuntu 22.04 LTS provides; workaround via musl builds now landing (PR #3274).
- **Proxy not honoured in JS execution** (#3273) – Windows users behind VPN/proxy cannot use `js_execution` tool.
- **Overly proactive agents** (#3275) – Agents self‑initiate tasks without user confirmation, causing loss of control. This regression is a trust issue.
- **Paste handling regression** (#3263, fixed in PR #3267) – Large pastes were replaced with file references, making them uneditable; now fixed with inline truncation and auto‑expand.
- **Digit key hijack in composer** (#3243) – Bare digits 1‑8 triggered hotbar slots instead of text entry (fixed in subsequent release).

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*