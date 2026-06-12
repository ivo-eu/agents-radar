# AI CLI Tools Community Digest 2026-06-12

> Generated: 2026-06-12 06:22 UTC | Tools covered: 9

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
**Date:** 2026-06-12 | **Analyst:** Senior Technical Analyst, AI Developer Tools Ecosystem

---

## 1. Ecosystem Overview

The AI CLI tool landscape on June 12, 2026, reveals a bifurcated ecosystem: mature tools (Claude Code, Copilot CLI) grapple with regressions and community frustration, while mid-tier tools (Qwen Code, CodeWhale, OpenAI Codex) push aggressively into new architecture—standalone binaries, declarative agents, and durable workflows. Cross-cutting concerns dominate: **context management instability**, **Windows/i18n support gaps**, and **MCP (Model Context Protocol) OAuth reliability** are universal pain points across all major tools. The most striking signal is the **breakdown of trust in agent autonomy**—three separate tools report bugs where agents execute unwanted actions (paid scripts, destructive loops, irreversible file changes), indicating the industry's safety layer is lagging behind capability growth. Meanwhile, niche tools (Kimi, Pi) focus on provider pluralism and terminal theming, suggesting a maturing market where differentiation shifts from core LLM capabilities to UX, extensibility, and operational reliability.

---

## 2. Activity Comparison (2026-06-12)

| **Tool** | **Hot Issues Tracked** | **Key PRs Active** | **Releases Today** | **Community Pulse** |
|---|---|---|---|---|
| **Claude Code** | 10 (1 critical: auto-compact) | 10 (5 open, 3 merged) | **2** (v2.1.174, v2.1.175) | High engagement, safety concerns rising |
| **OpenAI Codex** | 10 (1 hot: #14860, 103 comments) | 10 (all active) | **5 alpha** (0.140.0-a.9–14) | Very high, compaction bug dominating |
| **Gemini CLI** | 10 (P1 hangs, P1 subagent bugs) | 10 (7 closed, 3 open) | **0** | Moderate, focused on agent reliability |
| **Copilot CLI** | 10 (1 long-standing: #53) | 1 (stub PR) | **0** | High frustration, low maintainer activity |
| **Kimi Code** | 2 (quota calculation, WS crash) | 2 (1 closed, 1 open) | **0** | Low visible activity, niche concerns |
| **OpenCode** | 10 (Windows regressions dominate) | 10 (mixed open/closed) | **1** (v1.17.4) | Moderate, strong Windows fix pipeline |
| **Pi** | 10 (SSE timeout, compaction) | 10 (6 closed, 4 open) | **0** | Moderate, provider/protocol focus |
| **Qwen Code** | 10 (duplicate tool calls top) | 10 (5 open, 5 closed) | **1** (v0.18.0-preview.2) | Active, rapid feature iteration |
| **CodeWhale** *(fka DeepSeek TUI)* | 10 (PDF panic, TUI hangs) | 10 (1 merged, rest active) | **1** (v0.8.58 rebrand) | Growing, test coverage catch-up |

**Key observation:** OpenAI Codex shows the highest community engagement on a single issue (#14860), while Copilot CLI shows the lowest PR activity despite high frustration—a dangerous divergence for project health.

---

## 3. Shared Feature Directions

Despite different codebases, **four universal requirements** emerge across tool communities:

| **Requirement** | **Tools Mentioning** | **Specific Needs** |
|---|---|---|
| **Context compaction reliability** | Claude Code (#63015), OpenAI Codex (#14860), OpenCode (#27315), Pi (#4046) | Auto-compact not triggering, compaction corrupting data, manual control missing |
| **MCP/Provider OAuth stability** | Claude Code (#52871, #65036), Gemini CLI (#27664), OpenCode (#31998), Pi (#4945, #5427) | Trailing-slash bugs, token refresh failures, SSE timeouts, atomic writes |
| **Windows & i18n encoding support** | OpenAI Codex (#27506), OpenCode (#30068, #31985), Pi (#5640), CodeWhale (#3147) | Non-ASCII paths crash apps; CJK mojibake in copy/paste; PowerShell quoting |
| **Custom model / provider flexibility** | Claude Code (enforceAvailableModels), OpenAI Codex (#10867, #26472), Pi (#5363, #5509), Qwen Code (model switching) | Per-session model selection, third-party provider compatibility, config persistence control |
| **Sub-agent safety & transparency** | Claude Code (#67665, #67722), Gemini CLI (#22323, #26525), Qwen Code (#4999, #4955), CodeWhale (#3143, #3146) | Unwanted autonomous actions, misleading success signals, run ledgers, approval gates |

**Notable:** No tool community is *satisfied* with context compaction behavior. This is the single largest cross-cutting pain point—every team using AI CLI tools for sustained sessions is burning cycles on workarounds.

---

## 4. Differentiation Analysis

| **Dimension** | **Claude Code** | **OpenAI Codex** | **Gemini CLI** | **Copilot CLI** | **Qwen Code** | **OpenCode** | **CodeWhale** |
|---|---|---|---|---|---|---|---|
| **Core philosophy** | Tool orchestration + MCP | Rust-native performance | Sub-agent orchestration | GitHub ecosystem lock-in | Declarative agents + workflows | Multi-provider desktop app | Lightweight, fast, terminal-first |
| **Target user** | Power developers, MCP users | Enterprise/Pro, Rust shops | Google Cloud ecosystem | GitHub enterprise orgs | Open-source / cost-sensitive | Cross-platform desktop devs | TUI enthusiasts, early adopters |
| **Key differentiator** | MCP plugin ecosystem | Guardian security layer | Auto Memory system | Git integration parity | Declarative agent frontmatter | Desktop app + CLI hybrid | Rebranding to CodeWhale |
| **Recent architectural move** | Managed settings, preview tool | Standalone code-mode binary | AST-aware tools | None (stagnant) | Durable cron jobs, Safe Mode | Workspace-relative MCP | Sub-agent fanout |
| **Open source health** | Active, high community | Active alpha cycle | Steady but slower | Near-dormant | Fast iteration, test gaps | Strong Windows bugfix | Test coverage catch-up |

**Emerging patterns:**
- **Qwen Code** is the most innovative in agent definition (declarative frontmatter, durable cron, plan approval gates)—pushing toward a "agent-as-configuration" paradigm.
- **OpenAI Codex** is the most technically ambitious (standalone binary, Guardian decoupling, async hooks) but paying for it with alpha churn.
- **Copilot CLI** is the clear laggard: one stub PR in 24 hours, unaddressed top issue (#53) for 6 months, stream corruption regressions. The community is actively forking.
- **Claude Code** maintains the largest plugin ecosystem (flappy-claude game, hooks examples) but suffers from the broadest set of reliability regressions—suggesting feature velocity outpacing quality.

---

## 5. Community Momentum & Maturity

| **Metric** | **High Momentum** | **Mature but Stalled** | **Growing Rapidly** |
|---|---|---|---|
| **Issue engagement** | OpenAI Codex (#14860: 103 comments) | Copilot CLI (#53: 75👍, 6mo silence) | CodeWhale (#3098: new roadmap) |
| **PR velocity** | Qwen Code (10 PRs, 5 merged) | Copilot CLI (1 PR, 0 merged) | CodeWhale (10 PRs, test-focused) |
| **Release cadence** | OpenAI Codex (5 alphas/day) | Copilot CLI (0 releases) | OpenCode (v1.17.4 today) |
| **Community trust** | **Declining** for Claude Code, Copilot CLI | **Stable** for Gemini CLI | **Building** for Qwen Code, CodeWhale |

**Verbatim community sentiment:**
- **Claude Code:** "Auto-compact never triggers" (🔴 critical regression)
- **OpenAI Codex:** "Error running remote compact task" (103 comments, highest engagement)
- **Copilot CLI:** "Bring back the CLI commands to not break workflows" (75👍, zero official response in 6 months)
- **Kimi Code:** "2 hours for only 2 queries is ridiculous" (billing transparency)
- **CodeWhale:** "Sub-agent fanout can leave TUI stuck" (new tool, scaling pains)

**Assessment:** The ecosystem is entering a **"trust winter"** for established tools (Claude Code, Copilot CLI) while newer tools (Qwen Code, CodeWhale) have a window to capture disaffected users—if they can ship reliable context compaction and safety boundaries.

---

## 6. Trend Signals

1. **Context compaction is the #1 unsolved UX problem.** Every tool's community is burning cycles on auto-compact that doesn't trigger, corrupts data, or leaves sessions at 100%. This is the single largest user experience gap—and an opportunity for the first tool to solve it reliably.

2. **Agent autonomy without safety layers is backfiring.** Three independent tool communities report unwanted autonomous actions (paid scripts, destructive loops, irreversible changes). The industry is likely approaching a "kill switch" standard—expect mandatory approval gates and sandboxing to become table stakes within 6 months.

3. **Windows support is a persistent second-class citizen.** OpenAI Codex crashes on non-ASCII paths, OpenCode has 5+ Windows encoding bugs, CodeWhale's MSBuild integration fails, Pi only now adding image paste. The dominance of macOS in early AI tooling is creating a structural gap—and a market opportunity for any tool with first-class Windows support.

4. **Declarative agent configuration is the next paradigm.** Qwen Code and Claude Code are converging on YAML/frontmatter-based agent definitions (mcpServers, hooks, approval mode). This pattern reduces user friction and enables composable, shareable agent templates—think "Dockerfile for AI workflows."

5. **MCP is protocol-v1, but OAuth is failing enterprises.** Trailing-slash bugs, missing token refresh, and SSE timeouts are blocking enterprise adoption. MCP's OAuth flow needs a stability pass before enterprise orgs can depend on it.

6. **The "pet" phenomenon is real but niche.** OpenAI Codex pets and flappy-claude plugins show user appetite for gamification, but these are not decision-driving features. They signal community health more than product direction.

7. **Bounty-driven bug fixing is emerging as a pattern.** Claude Code's background-script bug has three competing bounty PRs ($29–$200). This market-based approach may accelerate fixes but risks fragmentation—expect standardization on bounty platforms (like Gitcoin) if the trend continues.

---

**Bottom line for decision-makers:**
- **For reliability-sensitive teams:** Wait on Claude Code v2.1.176+ for auto-compact fix. OpenAI Codex alphas are too unstable for production. Qwen Code or Gemini CLI are safer bets today.
- **For Windows-heavy orgs:** OpenCode is actively fixing encoding crashes. Copilot CLI and CodeWhale are riskier.
- **For enterprise security compliance:** No tool is ready. MCP OAuth failures and sub-agent autonomy bugs are universal. Budget for sandboxing layers.
- **For evaluating new tools:** Qwen Code's declarative agent model and CodeWhale's rapid iteration are worth watching. Copilot CLI's stagnation is a red flag.
- **For contributing:** Qwen Code and CodeWhale have the highest PR acceptance rate for non-trivial changes. Claude Code's plugin ecosystem is the most doc-ready.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report  
**Data snapshot:** 2026-06-12 | Source: github.com/anthropics/skills  

---

## 1. Top Skills Ranking (by Community Discussion Activity)

### #1 – Multi-skill definitions pack  
**PR #1046** (ALMMECHANICAL) — Adds three new skill definitions: *frontend-design*, *ai-experience-consultant*, and *automation-workflows-builder*.  
**Status:** Open | **Discussion:** Not yet detailed, likely a bulk addition covering design, consulting, and workflow automation.  
[View PR](https://github.com/anthropics/skills/pull/1046)

### #2 – Document-typography skill  
**PR #514** (PGTBoos) — Prevents orphan words, widow paragraphs, and numbering misalignment in AI-generated documents. Identifies a universal pain point in Claude’s output.  
**Status:** Open | **Discussion:** Strong rationale – typographic issues affect every generated document.  
[View PR](https://github.com/anthropics/skills/pull/514)

### #3 – ODT (OpenDocument) skill  
**PR #486** (GitHubNewbie0) — Create, fill, parse, and convert `.odt`/`.ods` files. Covers LibreOffice/ISO standard documents.  
**Status:** Open | **Discussion:** High demand for open-source document formats.  
[View PR](https://github.com/anthropics/skills/pull/486)

### #4 – Frontend-design clarity overhaul  
**PR #210** (justinwetch) – Rewrites the existing *frontend-design* skill to make instructions actionable and internally coherent.  
**Status:** Open | **Discussion:** Focuses on real-world usability – “every instruction is something Claude can actually follow.”  
[View PR](https://github.com/anthropics/skills/pull/210)

### #5 – Skill-quality-analyzer & skill-security-analyzer (meta-skills)  
**PR #83** (eovidiu) – Two new meta-skills for evaluating skill quality (structure, documentation, examples) and security (vulnerability scanning).  
**Status:** Open | **Discussion:** Early proposal (Nov 2025); addresses quality assurance for the ecosystem.  
[View PR](https://github.com/anthropics/skills/pull/83)

### #6 – SAP-RPT-1-OSS predictor skill  
**PR #181** (amitlals) – Brings SAP’s open-source tabular foundation model for predictive analytics on SAP business data.  
**Status:** Open | **Discussion:** Enterprise-focused; leverages Claude’s ability to work with structured SAP data.  
[View PR](https://github.com/anthropics/skills/pull/181)

### #7 – Testing-patterns skill  
**PR #723** (4444J99) – Comprehensive testing skill covering philosophy (Testing Trophy), unit tests, React components, integration tests.  
**Status:** Open | **Discussion:** Well-structured and broad; fills a gap in the skills collection for software testing.  
[View PR](https://github.com/anthropics/skills/pull/723)

### #8 – Color-expert skill  
**PR #1302** (meodai) – Self-contained color expertise: naming systems (ISCC-NBS, Munsell, XKCD, RAL), color spaces and “what to use when” tables.  
**Status:** Open | **Discussion:** Very recent (June 2026); covers a domain with clear practical applications (design, data viz, branding).  
[View PR](https://github.com/anthropics/skills/pull/1302)

---

## 2. Community Demand Trends (from Issues)

The top issues reveal four major demand clusters:

| Trend | Key Issues | Signal |
|-------|------------|--------|
| **Sharing & collaboration** | #228 (org-wide skill sharing), #184 (agentskills.io redirects) | Strong desire to distribute and reuse skills across teams; current manual download/upload flow is a bottleneck. |
| **Evaluation & optimization tooling** | #556 (run_eval.py 0% trigger rate), #202 (skill-creator best practice), #1061 (Windows compatibility) | The skill-creation toolchain has critical bugs (especially on Windows) that undermine the development loop. The community wants reliable evaluation and iteration. |
| **Security & trust boundaries** | #492 (namespace impersonation), #1175 (SharePoint security/context window) | Growing concern about community skills being distributed under the `anthropic/` namespace and about handling sensitive data (e.g., SharePoint). |
| **Integration & extensibility** | #29 (AWS Bedrock), #16 (MCP exposure), #412 (agent governance) | Users want skills to work outside Claude Code (Bedrock, MCP protocol) and need governance patterns for multi-agent setups. |
| **Windows compatibility** | #1061, also reflected in multiple PRs | A recurring pain point – many skill-creator scripts fail on Windows due to Unix assumptions (PATHEXT, cp1252 encoding, select on pipes). |

**Call to action:** The community urgently demands a reliable, cross-platform evaluation pipeline (Issue #556), transparent skill provenance (Issue #492), and easier sharing mechanisms (Issue #228).

---

## 3. High-Potential Pending Skills (Active PRs Likely to Merge Soon)

| Skill | PR | Why It Matters |
|-------|----|----------------|
| **n8n-builder & n8n-debugger** | [#190](https://github.com/anthropics/skills/pull/190) | Production-tested workflow automation skills for n8n; strong community interest in no-code automation. |
| **skill-creator fixes (Windows + UTF-8)** | [#1099](https://github.com/anthropics/skills/pull/1099), [#1050](https://github.com/anthropics/skills/pull/1050), [#1298](https://github.com/anthropics/skills/pull/1298), [#362](https://github.com/anthropics/skills/pull/362) | Multiple contributors independently patching the same pain points – merging likely soon. |
| **agent-creator meta-skill** | [#1140](https://github.com/anthropics/skills/pull/1140) | Adds Windows support and fixes multi-tool evaluation; addresses issue #1120. |
| **CONTRIBUTING.md** | [#509](https://github.com/anthropics/skills/pull/509) | Community health gap closure – low-hanging fruit for onboarding new contributors. |
| **SAP predictor, testing-patterns, color-expert** | Already listed in Top Skills Ranking – all actively discussed and well-scoped. | |

---

## 4. Skills Ecosystem Insight

**The community’s most concentrated demand is for a robust, cross-platform skill-creation and evaluation toolchain** — with fixes for Windows compatibility, accurate trigger detection, and byte-safe UTF-8 handling — enabling contributors to reliably develop, test, and share skills without hitting infrastructure bugs.

---

# Claude Code Community Digest — 2026-06-12

## Today's Highlights
Two maintenance releases rolled out: v2.1.175 introduces a managed `enforceAvailableModels` setting, while v2.1.174 brings mouse‑wheel acceleration control and fixes the `/model` picker. The community is buzzing about a spate of resource‑leak bugs (PTY handles, memory accumulation) and a critical auto‑compact regression that leaves sessions at 100% context without triggering compaction.

## Releases
- **[v2.1.175](https://github.com/anthropics/claude-code/releases/tag/v2.1.175)** — Added `enforceAvailableModels` managed setting. When enabled, the `availableModels` allowlist also constrains the Default model; user/project settings can no longer widen the list.
- **[v2.1.174](https://github.com/anthropics/claude-code/releases/tag/v2.1.174)** — Added `wheelScrollAccelerationEnabled` setting to disable scroll acceleration in fullscreen mode. Fixed the `/model` picker so Opus, Sonnet, etc. appear as separate rows depending on plan.

## Hot Issues
1. **[Auto‑compact never triggers (#63015)](https://github.com/anthropics/claude-code/issues/63015)** — Regression on v2.1.153: statusline reports “100% context used” but no compaction fires. 26 comments, 18 👍. Critical for long‑running sessions.
2. **[Assistant text between tool calls not rendered (#67071)](https://github.com/anthropics/claude-code/issues/67071)** — Text blocks between tool calls are invisible in both GUI and CLI, though persisted in session JSONL. Data‑loss regression from #41814.
3. **[Cowork locks up after ~5 min (#67780)](https://github.com/anthropics/claude-code/issues/67780)** — EventEmitter `MaxListenersExceededWarning` accumulates, freezing the Electron renderer. New today, 2 comments.
4. **[PTY leak exhausts macOS pool (#67781)](https://github.com/anthropics/claude-code/issues/67781)** — Desktop app leaks `/dev/ptmx` master fds, reaching the system limit after ~1.5 days. Duplicate reports (#67786) confirm widespread impact.
5. **[Socket connection closed unexpectedly (#67766)](https://github.com/anthropics/claude-code/issues/67766)** — Intermittent FIN‑mid‑stream kills active turns ~8–18 times/day. Packet captures provided.
6. **[MCP OAuth trailing slash breaks Entra ID (#52871)](https://github.com/anthropics/claude-code/issues/52871)** — Trailing slash on `resource` parameter causes AADSTS9010010. 22 comments, 16 👍. Enterprise blocker.
7. **[MCP OAuth token refresh never happens (#65036)](https://github.com/anthropics/claude-code/issues/65036)** — Daily “Connection expired” despite valid refresh token. 1 comment but 7 👍, high impact for persistent connections.
8. **[View/edit pasted text blocks (#3412)](https://github.com/anthropics/claude-code/issues/3412)** — Dictation/pasted content appears as collapsed non‑editable blocks. 75 comments, 266 👍. Long‑standing accessibility/enhancement request.
9. **[Eternal loop with exit hooks (#67665)](https://github.com/anthropics/claude-code/issues/67665)** — Claude gets stuck in a self‑created loop and may take irreversible actions without asking the user. Safety concern.
10. **[Preview tool feedback (#67789)](https://github.com/anthropics/claude-code/issues/67789)** — Users request device/viewport presets, better stability, and UX polish for the built‑in previewer.

## Key PR Progress
1. **[fix(ralph‑wiggum): case‑insensitive completion matching (#67753)](https://github.com/anthropics/claude-code/pull/67753)** — Uses `tr` for portable case‑insensitive comparison, preventing false negatives when Claude outputs different casing.
2. **[fix: correct state file path in ralph‑wiggum help.md (#61956)](https://github.com/anthropics/claude-code/pull/61956)** — Merged. Fixes path mismatch that caused plugin misconfiguration.
3. **[feat(plugins): add flappy‑claude terminal game (#50301)](https://github.com/anthropics/claude-code/pull/50301)** — Merged. Pure Python + curses game, playable via `/flappy‑claude` slash command.
4. **[Proposal: inline image rendering in TUI (#54551)](https://github.com/anthropics/claude-code/pull/54551)** — Closed. Adds README proposal to track the only first‑party client lacking inline image display.
5. **[PermissionDenied hook example with retry (#41694)](https://github.com/anthropics/claude-code/pull/41694)** — Closed. Documents the `PermissionDenied` hook introduced in v2.1.88.
6. **[PermissionDenied hook example (second PR) (#41695)](https://github.com/anthropics/claude-code/pull/41695)** — Closed. Duplicate of above, same content.
7. **[Fix for Claude autonomously running background scripts (#67722)](https://github.com/anthropics/claude-code/pull/67722)** — Open. Addresses a safety issue where Claude executed paid external scripts without user consent.
8. **[Bounty fix for same background‑script bug (#67699)](https://github.com/anthropics/claude-code/pull/67699)** — Open. Automated implementation via NVIDIA AI, bountied $29.
9. **[Alternative bounty fix for background‑script bug (#67697)](https://github.com/anthropics/claude-code/pull/67697)** — Open. Another automated attempt for the same issue.
10. **[Fix for account downgrade due to billing error (#67409)](https://github.com/anthropics/claude-code/pull/67409)** — Open. Bountied $200, aims to prevent accidental downgrades.

## Feature Request Trends
- **Context management visibility** — Users want reliable auto‑compact triggers (or manual control) and real‑time tokens‑per‑second (TPS) display during streaming (#58248).
- **MCP OAuth reliability** — Persistent complaints about trailing‑slash bugs and lack of automatic token refresh (#52871, #65036).
- **Accessibility & input flexibility** — Editing pasted/dictated content inline (#3412) is the highest‑voted open enhancement.
- **UI polish** — Requests for inline image rendering (#54551), better preview tool presets (#67789), and improved model picker consistency (addressed in v2.1.174).
- **Plugin ecosystem** — Continued interest in novel plugins (games, hooks examples) and better documentation for hooks.

## Developer Pain Points
- **Resource leaks** — PTY master fd leaks (#67781, #67786), unbounded memory accumulation from scheduled tasks (#67803), and Electron listeners piling up (#67780) are eroding trust in long‑running sessions.
- **Unpredictable agent behavior** — Claude autonomously executing paid scripts (#67699 family), getting into destructive loops (#67665), and modifying code without requests (#67793) raise serious safety flags.
- **Authentication friction** — MCP OAuth issues (trailing slash, missing refresh) plague enterprise users; false‑positive usage policy blocks (#67751) waste time.
- **Session management** — Auto‑compact not firing (#63015), hidden assistant text (#67071), and socket closures (#67766) degrade the interactive experience.
- **Cross‑platform gaps** — Windows users face MSIX/Smart App Control blocks (#67800), LSP tool finding issues (#59114), and poor localization (#58473). JetBrains plugin on WSL misses environment variables (#67502).

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-06-12

## Today's Highlights

A flurry of Rust alpha releases (0.140.0-alpha.9 through .14) landed, signaling active iteration on the CLI's core engine. The community is buzzing around two major poles: a long-running context compaction bug (Issue #14860) has amassed 103 comments and 73 reactions, while Windows users are experiencing a wave of startup crashes linked to non-ASCII user profile paths. On the PR side, the team is aggressively refactoring Guardian extension boundaries, introducing asynchronous hooks, and shipping a new standalone code-mode binary.

## Releases

Five Rust alpha releases were published in the last 24 hours—all `rust-v0.140.0-alpha.x` (versions 9, 10, 11, 13, and 14). No changelogs or commit-level summaries were provided, but the rapid versioning suggests active hotfixes or feature experimentation ahead of a stable 0.140.0 release.

## Hot Issues

1. **[#14860 – Error running remote compact task](https://github.com/openai/codex/issues/14860)** (103 comments, 73 👍)  
   The highest-engagement open bug. Users on Codex CLI 0.114.0 with gpt-5.4 report that remote compact tasks fail without clear recovery. Community frustration is palpable; many have attempted workarounds without success. This is likely impacting Pro/Enterprise users reliant on remote context management.

2. **[#10867 – Support custom model providers in app](https://github.com/openai/codex/issues/10867)** (15 comments, 34 👍)  
   A long-standing enhancement request (since February) to bring CLI's `/model` switching capability into the desktop app. The app currently ignores custom provider configurations. High demand from self-hosted or cost-sensitive teams.

3. **[#27296 – Fn global dictation hotkey stops working after update](https://github.com/openai/codex/issues/27296)** (9 comments, 14 👍)  
   A regression in the macOS app (v26.608.12217) that breaks the Fn key dictation shortcut across all apps. Users on 26.602.71036 report it worked fine. Signals a possible Electron/input-layer issue.

4. **[#27506 – Windows app crashes with non-ASCII user profile path](https://github.com/openai/codex/issues/27506)** (7 comments, 6 👍)  
   Korean and other non-ASCII characters in %USERPROFILE% cause `windows-updater.node` to throw "Illegal byte sequence" on launch. Multiple duplicates (e.g., #27722, #27780) suggest a widespread Windows Store/MSIX regression.

5. **[#27350 – Subagent transcript pane renders blank/loading](https://github.com/openai/codex/issues/27350)** (6 comments, 7 👍)  
   Closed quickly but still notable. The app's subagent UI panel stopped rendering transcripts after update 26.608.12217. The fix appears to have been rolled; caution for those still on the broken build.

6. **[#27335 – Support pets in VS Code integrated terminal with Sixel](https://github.com/openai/codex/issues/27335)** (6 comments, 0 👍)  
   A quirky enhancement: Codex CLI pets (ASCII pet companions) don't render in VS Code's integrated terminal even when Sixel support is enabled. The community has identified an environment-variable workaround, indicating a bug in terminal detection.

7. **[#26472 – Model selection should not be persisted to config](https://github.com/openai/codex/issues/26472)** (5 comments, 13 👍)  
   Enterprise users object that `/model` changes are written to `config`, causing accidental model switches on shared machines. The ask: keep model selection ephemeral (session-only).

8. **[#23842 – Stream disconnected before completion](https://github.com/openai/codex/issues/23842)** (4 comments, 2 👍)  
   A network-level bug causing `error decoding response body` mid-stream. Users on macOS with Cursor terminal hit this with plus-tier models. Might be linked to proxy or WebSocket timeout handling.

9. **[#27583 – Chat input unresponsive after opening from pet icon](https://github.com/openai/codex/issues/27583)** (2 comments, 0 👍)  
   A niche but annoying UI bug: after clicking the desktop pet to summon Codex, the chat input sometimes refuses keyboard focus. Reproducible on current app build 26.608.12217.

10. **[#27789 – macOS code_sign_clone directories leak ~1GB per launch](https://github.com/openai/codex/issues/27789)** (2 comments, 0 👍)  
    Closed today but highly relevant for macOS storage-constrained users. Every app launch creates a ~1GB clone under `/private/var/folders/.../code_sign_clone.*/` that is never cleaned up. The fix should ship soon.

## Key PR Progress

1. **[#27091 – Eagerly compact Guardian threads between reviews](https://github.com/openai/codex/pull/27091)**  
   Reduces latency spikes during approval reviews by proactively compacting Guardian threads before the next request. A performance win for multi-turn agent workflows.

2. **[#27696 – Load AGENTS.md from all bound environments](https://github.com/openai/codex/pull/27696)**  
   Expands multi-environment support so that `AGENTS.md` files from all environments are visible to the model, not just the primary one. Critical for users running separate dev/staging settings.

3. **[#27778 – Translate non-English issues](https://github.com/openai/codex/pull/27778)**  
   A new GitHub Actions workflow that auto-translates issues written in non-English languages into English, using Codex itself. Improves triage efficiency for the global community.

4. **[#27794 – Remove terminal resize reflow flag gates](https://github.com/openai/codex/pull/27794)**  
   Stabilizes terminal resize reflow as always-on. Removes deprecated config paths, simplifying the CLI surface for terminal emulator users.

5. **[#27793 – Fix child cwd inheritance for review sessions](https://github.com/openai/codex/pull/27793)**  
   Fixes a bug where review and guardian child sessions inherited a stale `config.cwd` after environment overrides, causing file-path confusion.

6. **[#27783 – Persist update dismissal without cache](https://github.com/openai/codex/pull/27783)**  
   Addresses Issue #27147: the "Don't remind me" button on update popups silently fails when `version.json` is missing. Now writes minimal version info directly.

7. **[#27791 – Reject transcript backtrack in side conversations](https://github.com/openai/codex/pull/27791)**  
   Fixes Issue #27735: attempting to edit previous messages in `/side` conversations now shows a graceful rejection instead of a cryptic thread rollback error.

8. **[#27790 – Remove the Guardian extension core dependency](https://github.com/openai/codex/pull/27790)**  
   Decouples the Guardian security extension from the core crate, reducing compilation overhead and enabling independent versioning of the extension.

9. **[#27727 – code-mode standalone: Cutover to new process](https://github.com/openai/codex/pull/27727)**  
   Phase 3 of a 4-phase stack migrating code mode to a new IPC-based standalone binary. This PR performs the actual cutover, replacing the old in-process implementation.

10. **[#27771 – Add a bounded runtime for async hooks](https://github.com/openai/codex/pull/27771)**  
    Introduces session-scoped, resource-limited execution for async hooks. The foundation for allowing hooks to run independently and deliver output to subsequent model requests.

## Feature Request Trends

- **Custom model support proliferation** – Two issues (#10867, #26472) highlight demand for per-session model selection and third-party provider compatibility in the desktop app. Community wants CI/CD pipelines to not accidentally swap models on shared configs.
- **Pets everywhere** – Issues like #27335 show users want ASCII pets to work across all terminal environments, including VS Code integrated terminals and Windows Terminal. A lighthearted but persistent request.
- **Usage window clarity** – #27435 and #27786 ask for unambiguous date+time in rate-limit panels, and clarification on whether merely opening the app triggers the 5-hour timer.
- **Async hooks and extensibility** – The PR stack (#27771, #27452, #27772) indicates the team is investing in a first-class extension system. Community feedback in related issues (not shown in top 30) likely expects easier plugin development.
- **Configuration management** – Enterprise users want managed (`requirements.toml`) settings for hooks, environment variables, and model selection (PR #27666). Standardization of enforced vs. preferred configs is a rising theme.

## Developer Pain Points

- **Windows crashes with non-ASCII paths** – At least three distinct issues (#27506, #27722, #27780) describe the desktop app crashing at launch if the user profile contains accented characters, CJK, or other non-ASCII symbols. The root cause appears in `windows-updater.node` and `cua_node` when calling `__char_to_wide`. Several duplicates suggest this is a critical hotfix priority.
- **Stream stability** – Issue #23842 and others report mid-stream disconnections with "error decoding response body." Users on macOS with Cursor and other third-party terminals are affected; the fix may involve adjusting WebSocket or HTTP/2 framing.
- **TUI input interference** – Approval prompts stealing keyboard focus (#27374) and unresponsive chat after pet icon launch (#27583) degrade the TUI experience. These race conditions are hard to reproduce but disruptive.
- **macOS disk bloat** – The `code_sign_clone` leak (#27789) wastes 1GB per launch. Already closed, but users on previous builds are cleaning up manually.
- **VS Code sidebar spinner** – Multiple Windows + VS Code extension issues (#11046, #26131, #19954) show the sidebar getting stuck on an infinite loading spinner, especially after updates or when using Dev Containers. High frequency suggests a fragile startup sequence in the extension's webview.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-06-12

## Today's Highlights
The repository saw no new releases but 30+ issues and 20+ PRs were updated, reflecting sustained focus on agent reliability and security. A cluster of PTY-related crash fixes landed, while the community continues to push for better sub-agent orchestration and memory system hardening. Notably, the 24-hour window saw no new releases or high-activity PR merges, suggesting an imminent batch of fixes may be queued.

---

## Releases
*No new versions were published in the last 24 hours.*

---

## Hot Issues
*10 noteworthy issues, ranked by impact and community engagement.*

1. **[#24353 – Robust component level evaluations (EPIC)](https://github.com/google-gemini/gemini-cli/issues/24353)**
   *p1, area/agent, customer-issue* – Tracks the expansion of behavioral eval tests (76+ generated) across 6 supported Gemini models. A key quality gate for ensuring agent changes don’t regress core behaviours.

2. **[#21409 – Generalist agent hangs](https://github.com/google-gemini/gemini-cli/issues/21409)**
   *p1, bug, 8 👍* – Simple tasks like folder creation cause indefinite hangs when the agent defers to the generalist sub-agent. Users can work around by explicitly disabling sub-agents, but the root cause remains under investigation.

3. **[#22745 – Assess impact of AST-aware file reads, search, and mapping](https://github.com/google-gemini/gemini-cli/issues/22745)**
   *p2, area/agent, EPIC* – Explores whether AST-aware tools (e.g., AST grep, tilth) can reduce token waste, improve method-bound accuracy, and speed up codebase navigation. High community interest (1 👍).

4. **[#22323 – Subagent recovery after MAX_TURNS reported as GOAL success](https://github.com/google-gemini/gemini-cli/issues/22323)**
   *p1, bug* – The `codebase_investigator` sub-agent signals `status: "success"` even when it hit maximum turns without performing analysis, hiding actual failures from users.

5. **[#21968 – Gemini does not use skills and sub-agents enough](https://github.com/google-gemini/gemini-cli/issues/21968)**
   *p2, bug* – Even with well-described custom skills (e.g., Gradle, Git), the model rarely invokes them autonomously. Users must explicitly instruct the agent, defeating the purpose of skill definitions.

6. **[#26525 – Add deterministic redaction and reduce Auto Memory logging](https://github.com/google-gemini/gemini-cli/issues/26525)**
   *p2, area/security* – Auto Memory sends transcripts to the model *before* redacting secrets, and logs skill content. A fundamental security concern raised by maintainers.

7. **[#26522 – Stop Auto Memory from retrying low-signal sessions indefinitely](https://github.com/google-gemini/gemini-cli/issues/26522)**
   *p2, bug* – If the extraction agent skips a session due to low signal, that session remains unprocessed and may be repeatedly re-surfaced, causing infinite retry loops.

8. **[#25166 – Shell command execution gets stuck with "Waiting input"](https://github.com/google-gemini/gemini-cli/issues/25166)**
   *p1, bug, 3 👍* – Simple CLI commands (no interactive input) hang after completion, showing "Awaiting user input". A high-friction issue for daily use.

9. **[#20079 – Symlink not recognized as agent in ~/.gemini/agents/](https://github.com/google-gemini/gemini-cli/issues/20079)**
   *p2, bug* – Symlinks inside the agents directory are silently ignored. Users who version-control agent definitions via symlinks cannot use them.

10. **[#21983 – Browser subagent fails in Wayland](https://github.com/google-gemini/gemini-cli/issues/21983)**
    *p1, bug* – The browser agent terminates with `GOAL` status but provides no usable output on Wayland sessions. Limits Linux users with modern display servers.

---

## Key PR Progress
*10 important pull requests updated in the last 24 hours.*

1. **[#27505 – Prevent extra spaces on width-0 CJK continuation cells](https://github.com/google-gemini/gemini-cli/pull/27505)**
   *p2, area/core, closed* – Fixes rendering glitches that injected unwanted whitespace between CJK characters in shell output, improving cross-platform copy-paste behaviour.

2. **[#27529 – Handle errors safely in shellExecutionService](https://github.com/google-gemini/gemini-cli/pull/27529)**
   *p2, area/core, closed* – Crashes from `EBADF` errors during PTY resizing are now gracefully caught, preventing the entire CLI from terminating.

3. **[#27526 – Harden uncaughtException handler for PTY resize errors](https://github.com/google-gemini/gemini-cli/pull/27526)**
   *p1, closed* – Additional guarding against unhandled exceptions in the PTY resize path, working in conjunction with #27529 for defence-in-depth.

4. **[#27527 – Guard isFunctionCall/isFunctionResponse against empty parts](https://github.com/google-gemini/gemini-cli/pull/27527)**
   *p1, closed* – Prevents crashes when the model returns empty or malformed function-call parts, a common edge case in agent tool-use loops.

5. **[#27524 – Read bootstrap settings from correct path when GEMINI_CLI_HOME is set](https://github.com/google-gemini/gemini-cli/pull/27524)**
   *p1, closed* – Ensures the custom home directory (`GEMINI_CLI_HOME`) is respected for bootstrap configuration, fixing broken custom installs.

6. **[#27678 – Hide ignored folders from session context](https://github.com/google-gemini/gemini-cli/pull/27678)**
   *p2, area/agent, open* – Directories matching `.gitignore`/`.geminiignore` are now omitted from the initial session context tree, reducing token waste and improving model focus.

7. **[#27664 – Write MCP OAuth tokens atomically](https://github.com/google-gemini/gemini-cli/pull/27664)**
   *p1, area/security, open* – Uses a temp-file + atomic rename pattern for OAuth token persistence to prevent corruption on crash. Fixes [#27663](https://github.com/google-gemini/gemini-cli/issues/27663).

8. **[#27572 – Handle tmux false positive background detection](https://github.com/google-gemini/gemini-cli/pull/27572)**
   *size/m, open* – Corrects a regression where white backgrounds were falsely detected inside tmux/mosh, causing improper theme switching and warnings.

9. **[#27705 – Promote Gemini 3.1 Flash Lite to GA and support Gemini 3.5 Flash](https://github.com/google-gemini/gemini-cli/pull/27705)**
   *size/xl, open* – Replaces retired preview models with the stable `gemini-3.1-flash-lite` model and adds support for the upcoming `gemini-3.5-flash`. Significant for users on the latest model track.

10. **[#27698 – Ensure zero-quota limits fail fast to prevent retry loop hang](https://github.com/google-gemini/gemini-cli/pull/27698)**
    *size/s-m, open* – Fixes a critical bug where accounts with a hard quota of 0 (e.g., unbilled free tier) would spin in a 10-attempt retry loop, effectively hanging the CLI.

---

## Feature Request Trends
From the past 24-hour issue activity, three major feature directions emerge:

- **AST-aware code understanding** – Multiple EPICs (e.g., #22745, #22746, #22747) propose using AST grep, AST-based file mapping, and method-bound detection to improve agent precision, reduce token usage, and speed up codebase investigations.
- **Memory system hardening** – Issues #26525, #26522, and #26523 call for deterministic redaction, infinite-retry prevention, and patch quarantine in the Auto Memory subsystem, driven by security and efficiency concerns.
- **Agent self-awareness and configuration** – Requests for the agent to understand its own CLI flags, hotkeys, and settings (e.g., #21432) and to respect `settings.json` overrides (e.g., #22267) indicate a desire for more transparent and configurable agents.

---

## Developer Pain Points
Recurring frustrations highlighted in today’s issue updates:

- **Agent hangs and stuck states** – The generalist agent (#21409) and shell command execution (#25166) both cause the CLI to hang indefinitely, severely impacting usability.
- **Misleading success signals** – Sub-agents reporting `GOAL` success when they actually hit execution limits (#22323) erodes trust in the tool’s feedback.
- **Sub-agent adoption failures** – The model rarely uses custom skills autonomously (#21968), undermining the investment in skill definitions.
- **Browser agent platform issues** – Wayland users cannot use the browser sub-agent (#21983), limiting Linux adoption.
- **Configuration friction** – Symlinks not recognised as agents (#20079) and `settings.json` entries ignored by the browser agent (#22267) add unnecessary setup complexity.
- **Tool limit errors** – With >128 tools enabled, the CLI returns a 400 error (#24246), forcing users to choose between disabling tools or facing crashes.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-06-12

## Today's Highlights
The community remains deeply frustrated with the **unaddressed removal of inline CLI commands (issue #53)**, the most-upvoted open issue, which has led to community forks. Meanwhile, **streaming corruption bugs (#3755, #3749)** are dominating recent reports, with multiple users reporting duplicated and truncated terminal output. A critical **regression in clipboard support on WSL2 (ARM64)** (#3534) and a **spate of input-keybinding regressions in v1.0.61** (#3760, #3768, #3770) suggest the latest release introduced several quality regressions.

## Releases
No new releases in the last 24 hours.

## Hot Issues
1. **[#53 – Bring back the GitHub Copilot in the CLI commands to not break workflows](https://github.com/github/copilot-cli/issues/53)** — The most reacted open issue (👍75, 37 comments). After six months with no official response, the community has started building alternatives like `shell-ai`. This remains the top pain point for long-time users.

2. **[#223 – "Copilot Requests" permission for fine-grained tokens should be visible for org-owned tokens](https://github.com/github/copilot-cli/issues/223)** — Enterprise users can’t create org-owned fine-grained tokens with the required permission, forcing use of personal PATs. 👍76, 30 comments. High enterprise demand.

3. **[#1481 – SHIFT+ENTER should spawn a line break, but executes the prompt instead](https://github.com/github/copilot-cli/issues/1481)** — Closed but still generating discussion (25 comments). The current `Ctrl+Enter` for line breaks is non-standard and confusing. Recently reopened in spirit by multiple new issues.

4. **[#892 – Add sandbox mode to restrict Copilot CLI file access to a specified working directory](https://github.com/github/copilot-cli/issues/892)** — Strong community support (👍49, 12 comments). Users want filesystem confinement to prevent accidental writes outside project roots.

5. **[#2306 – "You are not authorized to use this Copilot feature" – intermittent enterprise auth error](https://github.com/github/copilot-cli/issues/2306)** — Intermittent authorization failures (2–3 times per week) that disappear on retry. No official explanation yet.

6. **[#2486 – MCP server was blocked by policy – personal Pro+ account](https://github.com/github/copilot-cli/issues/2486)** — MCP functionality suddenly blocked for a paying user, with no clear cause. Workaround exists but feels hacky. Needs GitHub response.

7. **[#3755 – Reasoning/thinking display garbles streamed text with duplicated overlapping chunks](https://github.com/github/copilot-cli/issues/3755)** — Newly reported (2026-06-10) but already 3 comments. The `showReasoning: true` option corrupts output with repeated fragments like "numbnumber". Likely a rendering buffer issue.

8. **[#3534 – WSL2 (ARM64): `/copy` fails with `clip.exe exited with code 1` due to cmd.exe quoting in v1.0.55](https://github.com/github/copilot-cli/issues/3534)** — Clipboard copy is completely broken on ARM64 WSL2. A quoting bug in the Windows wrapper.

9. **[#2056 – Feature request: Scheduled/recurring prompts](https://github.com/github/copilot-cli/issues/2056)** — Users want agentic workflows that run autonomously on a timer (e.g., hourly health checks). Complementary requests in #2129 (loop/scheduled commands). Growing interest.

10. **[#2243 – Worktrees are a nightmare, should be disabled by default](https://github.com/github/copilot-cli/issues/2243)** — The agent’s automatic use of git worktrees has caused data loss and confusion. Users demand an opt-in model. 👍8.

## Key PR Progress
Only one pull request was updated in the last 24 hours:

- **[#3771 – Initial project setup](https://github.com/github/copilot-cli/pull/3771)** — Opened by `limenpchuolto112-creator` (seems to be a test/stub PR). No meaningful changes. No other notable PR activity this week.

With no significant contributions in flight, the project appears to be in a maintenance lull — possibly because the community is diverting energy to forks or waiting for the next official release.

## Feature Request Trends
The community is clearly demanding more **autonomous and scheduled execution** – multiple issues (#2056, #2129) ask for cron-like `/after` and `/loop` capabilities. **Filesystem sandboxing** (#892) is broadly supported. Enterprises are pushing for **fine-grained token support** (#223) and **authenticated MCP registry reads** (#3772) to avoid exposing internal registries. Users also want **per-repository plugin control** (#3761) and the ability to **disable worktrees globally** (#2243). Finally, a consistent request for **configurable context tiers** (#3762) indicates confusion around model context limits.

## Developer Pain Points
- **Terminal rendering corruption** is the hottest bug cluster: issues #3755, #3749, #3765, #3769 all describe streamed output being garbled, duplicated, or leaking tool calls as plain text (the "course" prefix bug).
- **Input keybinding regressions** in v1.0.61 have broken `Shift+Enter` for multiline (#3768), swapped `Ctrl+Enter`/`Ctrl+Q` behavior (#3760), and disabled Windows Voice Typing (`Win+H`) (#3770). These are basic UX failures.
- **Session token expiry** (#3763) silently kills in-progress workflows, though a `/resume` workaround exists.
- The **content exclusion service** (#3757) can fail-closed after token refresh, blocking all shell commands — a severe security-mitigation bug.
- **Environment pollution** (#3602): the Copilot SDK unconditionally injects `safe.bareRepository=explicit` into `process.env`, affecting all child processes. Surprising for a developer tool.
- **Oversized attachments** (#3767) permanently wedge a conversation with no recovery path — a UX dead end.
- **Multiple directory access prompts** (#3764) with no explanation of why re-approval is needed creates friction in agentic workflows.

The community’s patience is wearing thin, especially around the unresolved #53 (deprecated CLI commands) and the influx of v1.0.61 regressions. Stable, well-tested releases are urgently needed.

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-06-12

**GitHub Repository:** [MoonshotAI/kimi-cli](https://github.com/MoonshotAI/kimi-cli)

---

## Today’s Highlights

Community attention this week focuses on two long-standing pain points: a **flawed usage quota calculation** (#1994) that burns through subscription minutes much faster than expected, and a **WebSocket crash** in the Work tab (#2435) causing an infinite reload loop. Two pull requests show progress on forward-looking improvements—Python 3.13 compatibility guard (#1597) and a user-customizable YAML color skin system (#2170). No new releases were published in the last 24 hours.

---

## Releases

*No new releases in the last 24 hours.*  
Latest stable version remains **v1.41.0** (released prior to June 12).

---

## Hot Issues

1. **#1994 – kimiCode usage calculation is broken**  
   *Author: wanghonghust* | [Issue](https://github.com/MoonshotAI/kimi-cli/issues/1994)  
   **Why it matters:** Users report that two simple tasks consume the entire 2-hour quota, while the official documentation promises 300–1200 API requests per 5 hours. The core complaint is that Kimi’s billing counts **tokens** (with K2.6’s extremely long thought chains) rather than **API calls**, making the subscription feel drastically overpriced. The issue has **7 👍** and 5 comments, sparking heated debate about transparency. Community sentiment: “2 hours for only 2 queries is ridiculous.”

2. **#2435 – Work tab: “Daimon control WS not ready” + infinite reload at 99%**  
   *Author: JoseLuisMartinezMeza* | [Issue](https://github.com/MoonshotAI/kimi-cli/issues/2435)  
   **Why it matters:** A critical UI bug on Windows 10/11 makes the **Work** tab completely unusable. The WebSocket daemon fails to initialize, trapping users in a 99% loading loop. Only 1 comment so far, but the bug affects a core feature of the CLI for Windows users. Versions since 1.41.0 are impacted.

---

## Key PR Progress

1. **#1597 – fix: guard trafilatura import to prevent cascading tool load failure on Python 3.13**  
   *Author: he-yufeng* | [PR](https://github.com/MoonshotAI/kimi-cli/pull/1597)  
   **Description:** On Python 3.13, the `charset-normalizer` package ships incompatible `.so` binaries compiled with mypyc, causing `trafilatura` to crash at import time. This PR wraps the `import trafilatura` with a try/except guard, preventing the entire `web` tool module from failing. Essential for forward compatibility as Python 3.13 adoption grows.

2. **#2170 – feat: add user-customizable color skins via YAML (CLOSED)**  
   *Author: VrtxOmega* | [PR](https://github.com/MoonshotAI/kimi-cli/pull/2170)  
   **Description:** Introduces a `/skin` slash command and YAML-based skin loader (`~/.kimi/skins/<name>.yaml`) that lets users define complete color palettes in a Hermes-compatible format. Closed on June 11 after merging. This is a major quality-of-life improvement for developers who want personalized terminal aesthetics without modifying theme source code.

---

## Feature Request Trends

Based on recently active issues (including those not detailed above), the most‑requested feature directions are:

- **Transparent & fair usage billing** – Users want clear documentation on whether quotas are counted by tokens or API calls, and demand a “resets after X requests” model instead of token-based exhaustion.
- **Customizable UI theming** – Beyond YAML skins (#2170), requests for dark-only mode, high‑contrast themes, and per‑workspace color schemes continue to appear.
- **WebSocket reliability** – The Work tab crash (#2435) is only the latest symptom; developers want robust reconnection logic and clear error messages instead of infinite loading spinners.
- **Python 3.13 native support** – Several dependency incompatibilities beyond `trafilatura` have been reported. A comprehensive compatibility pass is desired.

---

## Developer Pain Points

Recurring frustrations voiced by the community in recent issues/PRs:

- **Quota calculation is counter‑intuitive** – Even paid subscribers are caught off guard when a single complex query (e.g., with K2.6 chain‑of‑thought) consumes disproportionately large token budgets. The lack of real‑time token/usage counters in the CLI makes debugging harder.
- **Platform instability on Windows** – WebSocket daemon failure (#2435) is not isolated; older issues mention similar WS‑related hangs on the `kimi web` command. Windows users feel left behind.
- **Dependency breakage on Python 3.13** – With PyPI now shipping mypyc‑compiled packages, CLI tools that rely on `trafilatura`, `charset-normalizer`, and similar libs break silently. The community wants either a pinned Python version baseline or a more aggressive dependency refresh cycle.
- **Lack of usage dashboard** – There is no built‑in command like `kimi usage --stats` to preview remaining quota or token cost per session, forcing users to guess.

---

*Generated for 2026-06-12. Data sourced from the MoonshotAI/kimi-cli GitHub repository.*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-12

## Today's Highlights
OpenCode **v1.17.4** shipped with workspace-relative local MCP servers, connector-based authentication, and new v2 session API endpoints. A **Windows encoding crisis** is being tackled by multiple PRs (PowerShell EncodedCommand, lazy code page detection). Meanwhile, a **range of i18n and desktop regression bugs** continue to draw community attention — the most-upvoted issue this week is the missing console terminal in the Desktop app (👍12).

---

## Releases
**v1.17.4** — Core improvements:
- `cwd` support for local MCP servers (workspace-relative startup).
- Connector-based authentication flows and stored provider credentials.
- v2 API endpoints to create / fetch sessions and list sessions.

---

## Hot Issues (10 noteworthy)

1. **#29829** — Desktop missing console terminal & "Open in Explorer" since v1.15.6 (👍12, 3 comments)  
   *Why it matters:* Major regression for Windows desktop users who rely on built-in terminal.
   [GitHub](https://github.com/anomalyco/opencode/issues/29829)

2. **#30068** — Copying Japanese text from chat produces mojibake (UTF-8 misinterpreted as Latin1) (👍3, 8 comments)  
   *Why it matters:* Breaks clipboard for CJK users; widespread impact.
   [GitHub](https://github.com/anomalyco/opencode/issues/30068)

3. **#27315** — Empty sessions auto-trigger “Input exceeds context window” and compaction (👍2, 3 comments)  
   *Why it matters:* Makes the product unusable in minimal configurations; core bug.
   [GitHub](https://github.com/anomalyco/opencode/issues/27315)

4. **#31996** — Invalid JSON schema due to unsupported regex lookaround in `fileKey` pattern on GPT 5.5 (👍2, 3 comments)  
   *Why it matters:* Blocks users of OpenAI-compatible providers with strict schema validators.
   [GitHub](https://github.com/anomalyco/opencode/issues/31996)

5. **#30381** — Cloudflare Workers AI: `AiError: Bad input` – message content format mismatch (0👍, 4 comments)  
   *Why it matters:* Provider-specific breakage; no workaround.
   [GitHub](https://github.com/anomalyco/opencode/issues/30381)

6. **#29875** — Desktop file explorer no longer shows “Open folder in file system” on Windows (0👍, 6 comments)  
   *Why it matters:* Another desktop regression affecting navigation for Windows users.
   [GitHub](https://github.com/anomalyco/opencode/issues/29875)

7. **#29859** — `@` symbol fails to reference files in any directory on Windows (0👍, 3 comments)  
   *Why it matters:* File referencing is a core workflow; completely broken on Windows.
   [GitHub](https://github.com/anomalyco/opencode/issues/29859)

8. **#31990** — SQLite UPSERT crash during `step-finish` event projection (0👍, 2 comments)  
   *Why it matters:* Hard crash during normal usage; data loss risk.
   [GitHub](https://github.com/anomalyco/opencode/issues/31990)

9. **#31987** — `/models` command cannot switch back to previous model after switching (0👍, 2 comments)  
   *Why it matters:* Breaks model switching workflow; easy to reproduce.
   [GitHub](https://github.com/anomalyco/opencode/issues/31987)

10. **#31991** — Version downgrades from 1.17.4 to 1.2.26 after Ctrl+C exit on Linux (0👍, 1 comment)  
    *Why it matters:* Bizarre and scary behaviour; suggests built-in upgrade logic corruption.
    [GitHub](https://github.com/anomalyco/opencode/issues/31991)

---

## Key PR Progress (10 important PRs)

1. **#31993** (`fix(app): restore desktop open menu`) — Fixes #29875 and #29951 (missing “Open in” control).  
   [GitHub](https://github.com/anomalyco/opencode/pull/31993)

2. **#31985** (`fix(shell): use PowerShell EncodedCommand for reliable UTF-8 output on Windows`) — Closes 5 issues related to Windows encoding (Shift-JIS, GBK, etc.).  
   [GitHub](https://github.com/anomalyco/opencode/pull/31985)

3. **#31980** (`fix(bash): lazy Windows code page detection with periodic refresh`) — Another Windows encoding fix for non-UTF-8 locales.  
   [GitHub](https://github.com/anomalyco/opencode/pull/31980)

4. **#31995** (`fix(opencode): preserve reasoning_content for Moonshot/Kimi tool-call assistant messages`) — Fixes #29619 (Kimi K2.6 tool call failure).  
   [GitHub](https://github.com/anomalyco/opencode/pull/31995)

5. **#31998** (`fix(opencode): gate empty MCP resource discovery loops`) — Prevents infinite polling storm on MCP servers that return empty resources.  
   [GitHub](https://github.com/anomalyco/opencode/pull/31998)

6. **#31940** (`feat(opencode): support MCP resource content`) — Resolves MCP resource mentions, delivers image blobs, adds `list_mcp_resources` / `read_mcp_resource` tools.  
   [GitHub](https://github.com/anomalyco/opencode/pull/31940)

7. **#31638** (`fix(session): avoid full history hydration after compaction`) — Addresses performance regression (#22208) by streaming only needed data.  
   [GitHub](https://github.com/anomalyco/opencode/pull/31638)

8. **#29684** (`feat(app): add markdown preview with enhanced markdown support`) — Adds Code/Preview toggle and Mermaid diagram rendering to the file viewer.  
   [GitHub](https://github.com/anomalyco/opencode/pull/29684)

9. **#31201** (`fix(app): preserve prompt drafts across session switches`) — Prevents draft loss when switching sessions in Desktop.  
   [GitHub](https://github.com/anomalyco/opencode/pull/31201)

10. **#31864** (`fix(ui): reduce settings panel backdrop opacity, add dialog animations, fix clipped icons`) — Polishes V2 settings UI visual bugs.  
    [GitHub](https://github.com/anomalyco/opencode/pull/31864)

---

## Feature Request Trends

- **Quick command/prompt insertion** – Several users want "preset phrases" (#31988) or shortcut commands for repetitive instructions (coding conventions, project context).
- **Desktop & IDE integration** – Feature requests for better VS Code extension UX (#31997), persistent sidebar entry, one-click launch, and deeper editor integration.
- **File picker refinements** –  Requests to respect `.ignore` negation patterns (#31801) and for the `glob` tool to ignore `.gitignore` only when appropriate (#31994).
- **Progress indicators** – Visual feedback for `opencode upgrade` (#31623) and other long-running operations.
- **Manual file editing** – The ability to edit files directly inside OpenCode (#26970) remains a long-standing request.

---

## Developer Pain Points

- **Windows regressions** dominate the bug tracker: missing desktop terminal, broken “Open in Explorer”, `@` file referencing failure, encoding issues (Shift-JIS, GBK, UTF-8). Three PRs are actively fixing this.
- **Context window / compaction instability** – Empty sessions triggering compaction, large MCP tool sets causing erratic token jumps (#27315, #27631).
- **i18n / CJK issues** – Japanese text mojibake (#30068), Chinese input method incompatibility (#28619).
- **Provider-specific failures** – Cloudflare Workers AI schema mismatch (#30381), Kimi reasoning_content missing (#29690), GPT 5.5 regex lookaround rejection (#31996).
- **MCP server interaction** – Unbounded polling loops, empty resource discovery, and crashes when MCP servers don’t implement certain methods.
- **Crash bugs** – SQLite UPSERT failure (#31990) and the baffling version downgrade after Ctrl+C on Linux (#31991).

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest – 2026-06-12

## Today's Highlights
The Pi ecosystem saw a flurry of bug fixes and feature contributions, with significant attention on OpenAI Codex transport stability, Amazon Bedrock provider improvements, and Windows/WSL image paste support. Multiple patches landed to prevent CLI hangs on non‑TTY environments and to resolve npm duplication of `@earendil-works/pi-ai`. New provider integrations for Bedrock Mantle and Anthropic Vertex continue to progress.

## Releases
No new releases in the last 24 hours. (Latest stable: v0.79.1)

## Hot Issues
1. **#4945 – openai-codex hangs on `Working…` with zero-usage aborted turns**  
   *Open, 54 comments, 30 👍* – A critical UX bug where `openai-codex` leaves the TUI stuck. Community tempers are high as workarounds require pressing Escape.  
   [GitHub](https://github.com/earendil-works/pi/issues/4945)

2. **#5363 – Add `amazon-bedrock-mantle` provider for OpenAI‑compatible models**  
   *Open, 9 comments* – Feature request to support Bedrock Mantle’s OpenAI‑compatible API. A corresponding PR (#5509) is already in draft.  
   [GitHub](https://github.com/earendil-works/pi/issues/5363)

3. **#4046 – Compaction just deletes everything**  
   *Closed, 9 comments* – A concerning data‑loss bug during compaction, now fixed but serving as a reminder of the importance of safe state management.  
   [GitHub](https://github.com/earendil-works/pi/issues/4046)

4. **#5427 – OpenAI Codex transport issues (SSE timeout)**  
   *Closed, 5 comments, 4 👍* – Users reported `Codex SSE response headers timed out after 10000ms` after updating to 0.78.1. Fixed in later patches.  
   [GitHub](https://github.com/earendil-works/pi/issues/5427)

5. **#5652 – npm-shrinkwrap.json missing integrity, causing duplicate `@earendil-works/pi-ai` install**  
   *Closed, 3 comments* – A packaging regression that splits the API provider registry into two isolated copies. Urgent fix shipped.  
   [GitHub](https://github.com/earendil-works/pi/issues/5652)

6. **#5633 – Kimi 2.6 error: `reasoning_content` missing in tool call message**  
   *Closed, 2 comments* – Model‑specific bug when resuming sessions with out‑of‑cache continuation. Quickly resolved.  
   [GitHub](https://github.com/earendil-works/pi/issues/5633)

7. **#5584 – Bedrock provider ignores `models.json` `apiKey`**  
   *Closed, 2 comments* – Gateway users cannot pass a bearer token via `apiKey`. A PR (#5586) later allowed it.  
   [GitHub](https://github.com/earendil-works/pi/issues/5584)

8. **#5501 – Tolerate extra keys on edit tool `edits[]` items**  
   *Closed, 2 comments* – Models sometimes append stray keys (e.g. `newText_strip`) causing validation failures. Relaxed schema.  
   [GitHub](https://github.com/earendil-works/pi/issues/5501)

9. **#5558 – Streamed model calls can hang indefinitely**  
   *Closed, 2 comments* – Headless usage against `opencode-go` provider could hang on transient upstream stalls – no inactivity/deadline handling.  
   [GitHub](https://github.com/earendil-works/pi/issues/5558)

10. **#5658 – Codex SSE headers timed out after 10000ms**  
    *Closed, 1 comment* – New session works, but old sessions show the timeout permanently. A follow‑up to #5427.  
    [GitHub](https://github.com/earendil-works/pi/issues/5658)

## Key PR Progress
1. **#5586 – fix(ai/bedrock): use resolved apiKey as bearer‑token fallback**  
   *Closed* – Fixes #5584, enabling explicit `apiKey` for gateway‑fronted Bedrock endpoints.  
   [GitHub](https://github.com/earendil-works/pi/pull/5586)

2. **#5650 – fix(ai): remove stale OpenRouter Kimi free model assertion**  
   *Open* – Unblocks CI by updating model definitions after OpenRouter changed its API response.  
   [GitHub](https://github.com/earendil-works/pi/pull/5650)

3. **#5385 – feat: detect first‑run terminal theme**  
   *Open* – Queries terminal for light/dark theme and persists it to settings. Improves initial user experience.  
   [GitHub](https://github.com/earendil-works/pi/pull/5385)

4. **#5627 – fix(coding-agent): skip first‑time setup for forks**  
   *Closed* – Prevents re‑running onboarding for fresh fork branches from the same repo.  
   [GitHub](https://github.com/earendil-works/pi/pull/5627)

5. **#5647 – fix(coding-agent): canonicalize path when loading context files**  
   *Closed* – Fixes duplicate `AGENTS.md` content when the config directory is a symlink.  
   [GitHub](https://github.com/earendil-works/pi/pull/5647)

6. **#5641 – fix(coding-agent): exit package commands from CLI entrypoint**  
   *Closed* – Resolves `pi update` and other package commands hanging after completion on all platforms.  
   [GitHub](https://github.com/earendil-works/pi/pull/5641)

7. **#5640 – feat(coding-agent): paste clipboard images via Ctrl+V on Windows Terminal**  
   *Closed* – Adds WSL‑aware image paste support, fixing #5632 where Ctrl+V was swallowed.  
   [GitHub](https://github.com/earendil-works/pi/pull/5640)

8. **#5637 – feat: add PI_GIT_TOKEN / GITHUB_TOKEN support for private HTTPS git installs**  
   *Closed* – Enables installing private repos via `pi install git:…` using token auth.  
   [GitHub](https://github.com/earendil-works/pi/pull/5637)

9. **#5509 – feat: Add Amazon Bedrock Mantle OpenAI Responses provider**  
   *Open* – New provider for Bedrock Mantle’s API, validating credentials via AWS STS. Adds GPT‑5.5/5.4 models.  
   [GitHub](https://github.com/earendil-works/pi/pull/5509)

10. **#5262 – feat(ai): add Anthropic Vertex provider**  
    *Open* – A thin adapter for Claude on Google Cloud Vertex AI, reusing existing Anthropic streaming code.  
    [GitHub](https://github.com/earendil-works/pi/pull/5262)

## Feature Request Trends
- **Provider proliferation** – Requests for new backend integrations continue (Bedrock Mantle, Anthropic Vertex, additional OpenAI‑compatible gateways).
- **Configurable timeouts** – The hard‑coded 10‑second SSE header timeout on Codex is a recurring pain point; users want an env variable or setting.
- **Windows/WSL parity** – Image paste, terminal rendering, and CLI process exit issues are being actively addressed.
- **Private repository support** – Token‑based authentication for git installs is now merged, but further improvements (SSH key preference) are anticipated.
- **Better error resilience** – Graceful handling of upstream stalls, model‑specific quirks (e.g. Kimi `reasoning_content`), and stream interruptions.

## Developer Pain Points
- **OpenAI Codex transport fragility** – Intermittent SSE timeouts and hang‑on-`Working…` scenarios frustrate users despite multiple fixes.
- **Provider credential confusion** – Bedrock’s multiple authentication paths (`apiKey`, `bearerToken`, env vars) are poorly documented and easily misconfigured.
- **npm packaging regression** – The duplicate `@earendil-works/pi-ai` installation broke custom provider registrations; root cause was a missing integrity field in `npm-shrinkwrap.json`.
- **CLI hangs on non‑TTY** – `pi -p` and `pi update` fail to exit when stdout is piped or on Windows, requiring manual kills.
- **TUI rendering glitches** – Single `+` rendered as `-`, `Box.render` crashes on non‑Component children, and symlink‑induced prompt duplication show immature UI testing.
- **OAuth port leak** – The fixed‑port (53692) OAuth callback server is not cleaned up on abort, potentially blocking future logins.

---
*Generated from GitHub repository [earendil-works/pi](https://github.com/earendil-works/pi).*

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest – 2026-06-12

## Today's Highlights
A new preview release v0.18.0-preview.2 is out, though it is a chore-only bump. The community is actively debating two critical bug clusters: **repeated/duplicate tool call execution** (issues #5014, #5015, #5016) and **long-context task degradation** (issues #5018, #5019). Meanwhile, several PRs land important features, including a Tool Fallback rule in the system prompt, a `--safe-mode` troubleshooting flag, and durable `/loop` cron jobs.

## Releases
- **v0.18.0-preview.2** – [Release notes](https://github.com/QwenLM/qwen-code/releases/tag/v0.18.0-preview.2) – No significant functional changes; primarily a version bump from v0.17.1 and a CLI fix that skips "thought" parts in copy output.

## Hot Issues (10 noteworthy)
1. **[#4821](https://github.com/QwenLM/qwen-code/issues/4821) – Declarative agent definitions via frontmatter** (CLOSED)  
   Request to support YAML frontmatter in Markdown files to define custom agents, mirroring Claude Code’s pattern. Generated 6 comments and was closed – likely superseded by PR #4996.

2. **[#4891](https://github.com/QwenLM/qwen-code/issues/4891) – Terminal resize fragments scrollback** (CLOSED)  
   Resizing the terminal during streaming leaves content rendered at mixed widths. Closed after PR #4919 debounced repaints and cleared stale scrollback.

3. **[#5018](https://github.com/QwenLM/qwen-code/issues/5018) – Long-context attention loss** (OPEN)  
   User reports severe forgetfulness and inattention in long tasks (Chinese description). High priority for model/long-context category; 3 comments. Community is watching.

4. **[#5019](https://github.com/QwenLM/qwen-code/issues/5019) – Long-context repetitive tool calls** (OPEN)  
   Related to #5018: identical tool calls repeated across rounds, causing API errors. Marked as duplicate but underscores a systematic issue with long sessions.

5. **[#4999](https://github.com/QwenLM/qwen-code/issues/4999) – /goal iteration counter resets on resume** (CLOSED)  
   `MAX_GOAL_ITERATIONS` safety cap is re-granted every session resume. Closed – likely addressed by a related fix.

6. **[#5015](https://github.com/QwenLM/qwen-code/issues/5015) – Repeated identical tool-call execution** (OPEN)  
   With deterministic providers, Qwen Code executes duplicate tool calls. Priority P1 – a critical reliability bug.

7. **[#5014](https://github.com/QwenLM/qwen-code/issues/5014) – Duplicate tool calls from provider stream** (OPEN)  
   Shell tool calls repeated due to identical IDs in stream. Closely related to #5015; 1 comment.

8. **[#5016](https://github.com/QwenLM/qwen-code/issues/5016) – Tool execution after SIGINT cancellation** (OPEN)  
   Tool work still runs even after user cancels streaming. Priority P1 – dangerous for destructive operations.

9. **[#4879](https://github.com/QwenLM/qwen-code/issues/4879) – Add `/cd` command** (CLOSED)  
   Feature request to change working directory without restart. Closed after community discussion.

10. **[#4712](https://github.com/QwenLM/qwen-code/issues/4712) – `/bug`/`/docs` crash on headless Linux** (CLOSED)  
    `xdg-open` missing causes crash. Fixed by PR #4716 which replaced direct `open` calls with secure browser launcher.

## Key PR Progress (10 important)
1. **[#4931](https://github.com/QwenLM/qwen-code/pull/4931) – Tool Fallback rule in system prompt** (CLOSED)  
   Instructs the model to try alternative tools before giving up on empty/unhelpful results. Simple but impactful improvement.

2. **[#4914](https://github.com/QwenLM/qwen-code/pull/4914) – OOM prevention and idempotent compaction** (OPEN)  
   Adds regression tests for compactOldItems, explicit GC, and debug log defaults. Targets issue #4815.

3. **[#4909](https://github.com/QwenLM/qwen-code/pull/4909) – Archive install sources for extensions** (OPEN)  
   Enables installing extensions from local or remote `.zip`/`.tar.gz` archives. Reuses existing validation and metadata paths.

4. **[#5017](https://github.com/QwenLM/qwen-code/pull/5017) – Support `.toml` command files in extensions** (OPEN)  
   Extension command discovery now includes `.toml` files, enabling tools like [caveman](https://github.com/JuliusBrussee/caveman) to work.

5. **[#4971](https://github.com/QwenLM/qwen-code/pull/4971) – Reduce interactive tool output memory** (OPEN)  
   Compacts oversized display metadata before storing in scheduler state, UI history, and chat recordings.

6. **[#4947](https://github.com/QwenLM/qwen-code/pull/4947) – Workflow P2: `parallel()` and `pipeline()`** (CLOSED)  
   Implements concurrent fan-out with sliding window (max 16 agents) and sequential pipeline on top of P1’s `agent()`.

7. **[#4996](https://github.com/QwenLM/qwen-code/pull/4996) – Declarative-agent mcpServers + hooks parity** (CLOSED)  
   Follows up on #4842 and #4870 to reach Claude Code 2.1.168 compat: frontmatter fields `mcpServers` and `hooks` now parse, round-trip, and fire at runtime.

8. **[#4853](https://github.com/QwenLM/qwen-code/pull/4853) – Add `enter_plan_mode` and Plan Approval Gate** (OPEN)  
   Model can proactively lower into plan mode; exit runs a singleturn approval step in AUTO/YOLO modes. Large community interest.

9. **[#4955](https://github.com/QwenLM/qwen-code/pull/4955) – Permission bubbling for background subagents** (OPEN)  
   Subagents can set `approvalMode: bubble` to surface permission prompts in the parent session’s Background Agents view.

10. **[#5004](https://github.com/QwenLM/qwen-code/pull/5004) – Durable cron jobs (survive restarts)** (OPEN)  
    `/loop` tasks can be persisted to `.qwen/scheduled_tasks.json` and auto-resume on relaunch. Default stays session-only.

## Feature Request Trends
- **Declarative agent configuration** – The most requested direction: define agents via Markdown frontmatter (YAML), supporting MCP servers, hooks, and approval modes. Almost all subagent/tools roadmap issues point here.
- **Persistent / session-crossing workflows** – Durable cron jobs and stateful `/goal` loops that survive restarts are heavily discussed.
- **Interactive slash commands** – Requests for `/cd`, improved `/extensions` manager (multi‑tab with discover/install), and plan approval gates suggest a desire for richer interactive control.
- **Safe-mode troubleshooting** – A `--safe-mode` flag to disable customizations (context, hooks, extensions, MCP) for clean baseline sessions is gaining traction.

## Developer Pain Points
- **Repeated / duplicate tool executions** – The #1 reliability complaint this week. Three P1 issues (#5014, #5015, #5016) highlight bugs where the same tool call is executed multiple times or after cancellation. The community expects immediate fixes to prevent side effects.
- **Long-context task degradation** – Users report severe tool repetition, forgetting, and API errors (400) on long sessions. Issue #5018 and #5019 are high‑priority but lack a clear solution.
- **Terminal resize artifacts** – Fragmented scrollback during streaming is now fixed (#4919), but the underlying rendering pipeline still needs hardening.
- **Headless/CI environment crashes** – While `/bug`/`/docs` crash is closed (#4712), similar “spawn xdg-open” errors appear in other paths (e.g., `/docs` for MCP).
- **Desktop input lock after Escape** – Fixed in PR #4788, but the pattern of UI state getting stuck after cancel/escape is a recurring theme.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest – 2026-06-12

The project has officially rebranded to **CodeWhale**. Today’s activity centers on the v0.8.59 roadmap, with a major focus on sub‑agent reliability, UX improvements inspired by Cursor, and a wave of test and cleanup PRs. A critical panic on PDFs with non-standard fonts was reported, alongside a handful of breaking shell and TUI bugs.

---

## Today’s Highlights

- v0.8.58 ships under the new **CodeWhale** name; the legacy `deepseek-tui` package is deprecated.
- A dense v0.8.59 execution roadmap (Issue #3098) bundles provider correctness, sub‑agent architecture, WhaleFlow workflows, localization, and diagnostics.
- Two high‑impact bugs surfaced: a PDF reading panic and a TUI hang during sub‑agent fanout.

---

## Releases

- **v0.8.58** – Renamed to **CodeWhale** (npm package, command, release assets). All future releases will use the new name. Migration guide: `docs/REBRAND.md`.  
  [GitHub Release](https://github.com/Hmbown/CodeWhale/releases/tag/v0.8.58)

---

## Hot Issues (10 most noteworthy)

1. **#3098 – v0.8.59 execution roadmap** – Consolidates work originally slated for v0.8.60 (provider model correctness, sub‑agents, WhaleFlow workflows, localization, docs). Without a roadmap the backlog becomes unmanageable.  
   [Issue](https://github.com/Hmbown/CodeWhale/issues/3098)

2. **#3095 – Sub‑agent fanout can leave TUI stuck** – When the parent model launches several sub‑agents the UI shows “waiting for model” with no feedback on whether they are alive or failed. No backpressure or recovery.  
   [Issue](https://github.com/Hmbown/CodeWhale/issues/3095)

3. **#3149 – `read_file` panics on PDFs with non-Identity-H CMap fonts** – An assertion failure (`name == "Identity-H"`) kills the entire turn without a user‑facing error.  
   [Issue](https://github.com/Hmbown/CodeWhale/issues/3149)

4. **#3143 – Prompt source map and context‑usage report** – Inspired by Cursor’s explainable prompt context: show which rules, tools, memory, and skills contributed to the prompt.  
   [Issue](https://github.com/Hmbown/CodeWhale/issues/3143)

5. **#3144 – Natural‑language auto‑review policy and pre‑push gate** – A middle ground between manual approval and unchecked autonomous execution, modelled on Cursor’s `bugbot` review.  
   [Issue](https://github.com/Hmbown/CodeWhale/issues/3144)

6. **#3147 – MSBuild FileTracker fails in CodeWhale shell** – `cmake --build` is unusable on Windows when Visual Studio’s FileTracker cannot initialize inside the managed PowerShell.  
   [Issue](https://github.com/Hmbown/CodeWhale/issues/3147)

7. **#3145 – Visual inspection artifacts for browser/UI tasks** – Providing the agent with richer evidence (element selection, layout, screenshots) during UI interactions.  
   [Issue](https://github.com/Hmbown/CodeWhale/issues/3145)

8. **#3142 – Agent run ledger with follow‑up, takeover, artifact receipts** – Treat background/cloud work as an operational run rather than a hidden chat branch.  
   [Issue](https://github.com/Hmbown/CodeWhale/issues/3142)

9. **#3102 – First‑class clarification questions for agents** – Give agents a harness‑native way to ask the user for clarification instead of emitting a plain message that can be missed.  
   [Issue](https://github.com/Hmbown/CodeWhale/issues/3102)

10. **#3146 – Render tool work as expandable activity metadata rows** – Compress noisy tool calls (file reads, command output) into collapsible rows, keeping the main conversation readable.  
   [Issue](https://github.com/Hmbown/CodeWhale/issues/3146)

---

## Key PR Progress (10 important changes)

1. **#3148 – fix(exec): resolve `--model auto` via env** – Fixes a clap argument parsing bug where `--model auto` was consumed as prompt text, causing fallback to `deepseek-reasoner`.  
   [PR](https://github.com/Hmbown/CodeWhale/pull/3148)

2. **#2951 – fix(prompts): explain `visibility="internal"` in Runtime Policy Reference** – Adds documentation for the `visibility="internal"` attribute on `<runtime_prompt>` tags.  
   [PR](https://github.com/Hmbown/CodeWhale/pull/2951)

3. **#3141 – ⚡ Optimize `get_thread_detail` (N+1 fix)** – Batches directory scans per turn, significantly reducing item‑fetch overhead.  
   [PR](https://github.com/Hmbown/CodeWhale/pull/3141)

4. **#3140 – 🔒 Fix command injection in hooks** – Switches from `sh -c`/`cmd /C` to direct execution, preventing shell metacharacter expansion.  
   [PR](https://github.com/Hmbown/CodeWhale/pull/3140)

5. **#3139 – ⚡ Parallelize skill syncing** – Uses `join_all` to fetch and sync community skills concurrently instead of sequentially.  
   [PR](https://github.com/Hmbown/CodeWhale/pull/3139)

6. **#3138 – 🧪 test: add `ToolError::path_escape` test** – Covers the constructor and its `Display` output.  
   [PR](https://github.com/Hmbown/CodeWhale/pull/3138)

7. **#3137 – 🧪 Add tests for `release_base_url_from_env`** – Uses `serial_test` to safely test environment‑dependent logic.  
   [PR](https://github.com/Hmbown/CodeWhale/pull/3137)

8. **#3135 – 🧹 Remove unused `prompt_persist` module** – Entire file annotated `#[allow(dead_code)]` and unused; removed to reduce clutter.  
   [PR](https://github.com/Hmbown/CodeWhale/pull/3135)

9. **#3134 – 🧪 Add unit tests for `is_beta_tag`** – Covers casing variations and embedded beta tags.  
   [PR](https://github.com/Hmbown/CodeWhale/pull/3134)

10. **#3131 – 🧪 Add unit tests for `resolve_release_query`** – Tests environment‑sensitive release‑query resolution.  
   [PR](https://github.com/Hmbown/CodeWhale/pull/3131)

---

## Feature Request Trends

- **Explainable context** – Users and developers want to see which rules, tools, memory, and skills are contributing to the prompt (inspired by Cursor).  
- **Sub‑agent transparency** – Requests for run ledgers, artifact receipts, and follow‑up/takeover controls so that background work is visible as an operational run.  
- **Rich UI feedback** – Expandable tool‑activity rows, visual inspection artifacts (screenshots, element selection), and clarification modals for agents.  
- **Safety gates** – Natural‑language auto‑review policies and pre‑push gates to balance autonomy with user oversight.  
- **Localization & documentation** – README/site localization and better docs are bundled into the v0.8.59 roadmap.

---

## Developer Pain Points

- **Unexpected panics** – `read_file` crashes on PDFs with non-Identity-H CMap fonts, killing the whole turn.  
- **TUI hangs** – Sub‑agent fanout leaves the UI stuck in “waiting for model” with no progress indication or recovery.  
- **Shell incompatibilities** – MSBuild FileTracker fails inside the CodeWhale managed shell on Windows, breaking `cmake --build`.  
- **CLI argument quirks** – `--model auto` is consumed as prompt text due to `trailing_var_arg` on `ExecArgs`.  
- **Security in hooks** – Command injection was possible via `sh -c`/`cmd /C` (now fixed in #3140).  
- **Missing test coverage** – A large batch of PRs today addresses untested constructors and environment‑sensitive functions, indicating a codebase still hardening.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*