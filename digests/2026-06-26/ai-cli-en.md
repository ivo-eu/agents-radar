# AI CLI Tools Community Digest 2026-06-26

> Generated: 2026-06-26 10:38 UTC | Tools covered: 9

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

# AI CLI Tools Landscape: Cross-Tool Comparison Report

**Date:** 2026-06-26  
**Prepared by:** Senior Technical Analyst, AI Developer Tools Ecosystem

---

## 1. Ecosystem Overview

The AI CLI tools ecosystem has entered a phase of **intense platform maturation** where feature velocity is colliding with reliability expectations. Today's community data reveals three systemic pressures: (1) enterprise adoption is driving demands for secrets management, persistent permission systems, and centralized configuration; (2) the cost of model inference is emerging as a first-class UX concern, with rate-limit anomalies (Codex) and runaway token consumption (Claude Code, OpenCode) eroding user trust; and (3) subagent architectures—once a differentiator—are becoming a shared pain point, with all five tools employing agent delegation reporting hangs, misreported success states, or silent failures. The landscape is bifurcating: mature tools (Claude Code, Codex, OpenCode) are fighting infrastructure fires while rapidly iterating tools (Gemini CLI, Qwen Code, CodeWhale) push ambitious new collaboration primitives. A notable undercurrent is the **spread of MCP (Model Context Protocol)** integration—every tool now has MCP surface area—signaling a convergence on tool-extensibility standards even as each vendor differentiates on core interaction models.

---

## 2. Activity Comparison (2026-06-26)

| Tool | Hot Issues (24h) | Key PRs (24h) | Release Today? | Notable |
|------|-----------------|---------------|----------------|---------|
| **Claude Code** | 10 (3 open) | 2 | ✅ v2.1.193 | Mostly duplicate closures; `classifyAllShell` shipped |
| **OpenAI Codex** | 10 (10 open) | 10+ | ✅ 3 alpha releases | Rate-limit crisis (#28879, 315👍); MCP OAuth stack |
| **Gemini CLI** | 10 (10 open) | 10 | ✅ 3 releases | Security PRs on trust dialog + OAuth socket reuse |
| **GitHub Copilot CLI** | 10 (10 open) | 1 | ✅ v1.0.66-0 | MCP toggle shipped; enterprise config gaps |
| **Kimi Code CLI** | 2 (2 open) | 2 | ❌ No release | Low activity; MCP/scaling bug on Windows |
| **OpenCode** | 10 (8 open) | 10 | ✅ v1.17.11 | CPU usage dominates (3 high-traffic issues); session snapshots |
| **Pi** | 10 (10 open) | 10 | ❌ No release | Orchestrator daemon (experimental); connection hangs with Codex |
| **Qwen Code** | 10 (10 open) | 10 | ❌ No release | Multiplayer agent (#5887); CI contamination (#5882) |
| **CodeWhale** | 10 (6 open) | 10 | ✅ v0.8.65 (rebrand) | 20+ PRs today; IME fix for CJK; permission rules |

**Key takeaway:** Activity is uniformly high across the top 7 tools. Kimi Code CLI is the outlier at low velocity. CodeWhale’s 20+ PR burst and Gemini CLI’s three releases in 24 hours signal rapid iteration phases.

---

## 3. Shared Feature Directions

The following requirements appear across **three or more** tool communities, indicating industry-level consensus:

| Requirement | Tools Involved | Specific Pain Points |
|-------------|---------------|---------------------|
| **Secrets & credential management** | Claude Code (#32733, 112👍), Gemini CLI (#26525), OpenCode (#34030, #34048) | API keys/enterprise tokens injected into web sessions leak to logs/model context; GitHub Copilot model auth fails with missing `Integration-Id` header |
| **Persistent permission rules** | Claude Code (#70051), CodeWhale (#1186, #3650), OpenCode (approval systems) | Piped commands bypass allow-rules; permission decisions not saved across sessions; no `deny`/`allow`/`ask` scoping by tool or path |
| **Subagent reliability & observability** | Gemini CLI (#22323, #21409), OpenCode (#33878, #34043), Claude Code (#70536) | Subagents report "success" after MAX_TURNS; silent completion; model name mangling in fallback chains; loops with no structured tool calls |
| **Configurable polling / background activity** | Claude Code (#70186), OpenCode (#21470, #30086), Pi (idle CPU) | Git polling every ~20s causes lock contention; CPU at 50% even during API backoff; animated progress indicators burn cycles |
| **Windows platform parity** | Codex (#29072, #30009), Claude Code (#70039, #71561), Copilot CLI (#3501, #3534), Kimi Code (#2475), Qwen Code (#5892) | `apply_patch` sandbox fails (error 1920); updater loops; scroll bar rendering; git-bash detection; process-tree leaks in PTY |
| **MCP server management & reliability** | Copilot CLI (#2956, #3564), Codex (OAuth PR stack #29017-21), Pi (#6087), OpenCode (OAuth URL fix), Kimi Code (#2475) | RPC timeouts too short; OAuth token refresh crashes; large tool sets (212) not handled; MCP servers blocked by policy |
| **Plan/agent mode fidelity** | CodeWhale (#3568), Qwen Code (#5881), Claude Code (autoMode) | Model executes edits while in "plan" mode; approval gates don't cover all exits; auto-mode classifier bypassed by shell patterns |
| **HTTPS-only transport options** | Codex (#27381, #19821), Pi (networking issues), Gemini CLI (behind proxies) | WebSocket retries stall for minutes; proxy/VPN users blocked; no flag to force HTTP fallback |

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | Copilot CLI | OpenCode | Pi (mono) | Qwen Code | CodeWhale |
|-----------|-------------|-------------|------------|-------------|----------|------------|-----------|-----------|
| **Core philosophy** | Secure agent with permission gates | Multi-agent orchestration | Subagent delegation + Auto Memory | Enterprise Copilot integration | Open-source modular TUI | Minimalist provider-agnostic TUI | Collaborative persistent agent | Rust-native CLI with full tool ecosystem |
| **Target user** | Pro developers, solo to team | Power users, heavy API consumers | Google ecosystem, team workflows | GitHub enterprise users | Open-source community, early adopters | Tinkerers, provider-switchers | Multiplayer teams, DingTalk users | CLI purists, Rust ecosystem |
| **Differentiator** | Secrets injection (#1 ask); `classifyAllShell` | Rate-limit attack surface; MCP OAuth maturity | Security PR velocity; AST-aware code requests | Managed settings for orgs; MCP toggle shipped | Session snapshots + revert; subagent interrupt | Orchestrator daemon; Friendli provider added | Multiplayer channel agents (tag); team memory tier | Rebrand to CodeWhale; permission `.toml`; CJK IME fix |
| **Key weakness today** | Model loop regression (#70536) | Rate-limit crisis; Windows sandbox broken | Agent hangs; subagent misreported success | Theme override; autopilot persistence bug | CPU-bound idle polling; Copilot integration fragile | Connection hangs with Codex; TUI scroll regressions | CI contamination; TypeScript LSP broken | Plan/agent mode confusion; prompt bloat |
| **Technical stack** | TypeScript, MCP | Rust, multi-agent v2 | TypeScript, tool registry DI | Go, MCP, OpenTelemetry | TypeScript, TUI, session API | Go, experimental daemon | TypeScript, channels, ACP | Rust, TUI, MCP, i18n |
| **Release cadence** | Biweekly stable + patches | Multiple alphas weekly | Nightly + preview + stable | ~Weekly stable | ~Weekly stable | Irregular (no release today) | Irregular (no release today) | Weekly stable + critical patches |

**Key differentiations:**
- **Security posture:** Gemini CLI shipped more security PRs in 24h (#27915 trust dialog, #28103 OAuth socket, #27845 trust-before-auth) than any other tool this week. Claude Code has the most community-validated security request (#32733). Codex is investing in MCP OAuth concurrency.
- **Multiplayer:** Qwen Code is uniquely pursuing persistent channel-resident agents (#5887/#5888) and git-shared team memory (#5886). No other tool has this focus.
- **Infrastructure reliability:** Codex is in crisis mode on rate-limits. OpenCode has a CPU crisis. Others (Copilot, Gemini) are comparatively stable but face feature-regression issues.
- **Platform support breadth:** Pi is adding provider diversity (Friendli, NVIDIA NIM, Ant Ling). CodeWhale is optimizing for Rust-native minimalism. Claude Code and Codex treat Windows as secondary.

---

## 5. Community Momentum & Maturity

| Tool | Community Size Indicator | Velocity Signal | Maturity Assessment |
|------|-------------------------|-----------------|---------------------|
| **Claude Code** | Highest-voted open issue (112👍); large duplicate volume | Stable releases; 90-day stale timeout PR signals patient triage | **Most mature**—enterprise-grade security focus, but model reliability regressions |
| **OpenAI Codex** | #14593 has 623 comments (longest thread across all tools) | 3 alpha releases today; 10 PRs; high issue volume | **High maturity, under stress**—rate-limit crisis eroding trust despite strong engineering output |
| **Gemini CLI** | 8👍 highest vote (#21409); 10 PRs | 3 releases in 24h; security fix velocity is highest | **Rapidly maturing**—aggressive security posture; subagent reliability is the weak point |
| **Copilot CLI** | 9👍 highest vote (#3501); low PR volume | 1 release; 1 PR today | **Stable enterprise tool**—lowest community drama; feature requests focused on configuration gaps |
| **OpenCode** | #21470 (14 comments, 13👍); high comment density | Session snapshots shipped; 10 PRs; active bug triage | **High community engagement, mid-maturity**—CPU issues signal architecture debt; fastest subagent feature iteration |
| **Pi** | #4945 (71 comments, 30👍) | Orchestrator daemon (experimental); 10 PRs | **Growing quickly**—connection reliability is #1 blocker; provider diversity is a strength |
| **Qwen Code** | Low upvote counts but high proposal density | 10 PRs; multiplayer features; CI issues (contamination) | **Early but ambitious**—unique collaboration features; infrastructure reliability needs work |
| **CodeWhale** | Low upvote counts; 20+ PR burst | Rebrand + massive PR day; permission rules shipped | **Rapid iteration phase**—Rust-native advantage; plan/agent confusion is the critical weakness |

**Momentum ranking (today):** CodeWhale > Gemini CLI > OpenCode > Qwen Code > Pi > Codex > Claude Code > Copilot CLI > Kimi Code

*(Ranking based on PR volume + issue resolution + new feature surface in the last 24h, weighted by release cadence.)*

---

## 6. Trend Signals for Developers

### 1. The Model is the New Attack Surface
The most serious security issue this week—**injected pseudo-instructions in tool results urging `git push --force`** (Claude Code #71564)—demonstrates that model context poisoning is now a real attack vector. Combined with Gemini CLI’s inverse trust dialog (#27915) and Codex’s missing `Copilot-Integration-Id` header (#34048), the community is waking up to the fact that AI CLI tools introduce a new class of supply-chain risk. **Takeaway:** Vet your tool’s permission model; prefer tools that surface denial reasons and support persistent, scoped rules.

### 2. Subagent Architectures Are Still Immature—Expect Hangs
Every tool with subagents (Gemini, OpenCode, Codex, Claude Code) has reports of hangs, misreported success, or infinite retry loops. This is the single most consistent pain point across the ecosystem. **Takeaway:** If your workflow depends on background subagents, test fallback behavior rigorously; budget for manual interrupt handling.

### 3. Rate-Limit Anomalies Signal Backend Accounting Fragility
Codex’s crisis (10–20× faster token burn, #28879) and OpenCode’s runaway CPU suggest the **backend cost management layer is under-engineered** relative to model capability growth. As models become more capable and expensive, accurate accounting is non-negotiable. **Takeaway:** Monitor your actual token consumption versus reported usage; consider tools that expose programmatic cost data (Claude Code #70107).

### 4. Windows Remains a Second-Class Citizen
Despite the enterprise push, Windows users across 6 tools report broken installers, sandbox failures, clipboard bugs, process leaks, and rendering glitches. **Takeaway:** If your team uses Windows, budget for workaround time or choose tools with demonstrated Windows investment (Qwen Code’s tree-kill fix #5892 is a rare positive signal).

### 5. Enterprise Configuration Is the Next Battleground
Copilot CLI’s managed settings (#3909), Claude Code’s secrets injection (#32733), and Gemini CLI’s team memory tier (#5886) all point to a shared recognition: **individual developer tooling must integrate with org-level policy, credential management, and observability pipelines.** **Takeaway:** Prefer tools that support OAuth, centralized config, and OpenTelemetry export (Copilot CLI shipped this today).

### 6. The "Smart Default" Era Is Ending
Users across tools are asking for **configurable polling intervals, HTTPS-only transport flags, overridable system prompts, and custom auto-compact thresholds.** The one-size-fits-all approach is failing as use cases diversify. **Takeaway:** Evaluate a tool’s configuration surface area early; the tools with the most knobs today (Gemini CLI’s tool registry DI, OpenCode’s session API) are likely to age better.

---

*Report generated from community digest data for 2026-06-26, covering Claude Code, OpenAI Codex, Gemini CLI, GitHub Copilot CLI, Kimi Code CLI, OpenCode, Pi, Qwen Code, and CodeWhale.*

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report (Data as of 2026-06-26)

## Top Skills Ranking

The following Pull Requests have attracted the most community discussion (sorted by comment count). All are currently **open**.

1. **#1298 – Fix `run_eval.py` always reporting 0% recall**  
   *Functionality*: Patches the skill‑creator’s evaluation pipeline – installs the eval artifact as a real skill, fixes Windows stream reading, parallel workers, and trigger detection.  
   *Discussion*: Addresses the root cause of the widely reported “recall=0%” bug (see Issue #556).  
   👉 [PR #1298](https://github.com/anthropics/skills/pull/1298)

2. **#514 – Add `document-typography` skill**  
   *Functionality*: Prevents orphan word wrap, widow paragraphs, and numbering misalignment in AI‑generated documents.  
   *Discussion*: Users appreciate the practical quality‑of‑life improvement; some debate whether this overlaps with existing document skills.  
   👉 [PR #514](https://github.com/anthropics/skills/pull/514)

3. **#538 – Fix case‑sensitive file references in `pdf` skill**  
   *Functionality*: Corrects 8 mismatched file names (e.g. `REFERENCE.md` → `reference.md`) that break on case‑sensitive filesystems.  
   *Discussion*: Minimal controversy – widely seen as a necessary robustness fix for cross‑platform use.  
   👉 [PR #538](https://github.com/anthropics/skills/pull/538)

4. **#486 – Add ODT skill (OpenDocument text)**  
   *Functionality*: Creates, fills, reads, and converts `.odt`/`.ods` files, including template filling and ODT‑to‑HTML conversion.  
   *Discussion*: Strong demand from LibreOffice and open‑source document users; some questions about `.ods` spreadsheet support scope.  
   👉 [PR #486](https://github.com/anthropics/skills/pull/486)

5. **#210 – Improve `frontend-design` skill clarity**  
   *Functionality*: Rewrites the frontend‑design skill to make every instruction actionable within a single conversation.  
   *Discussion*: Positive reception for better structure; a few requests to add specific CSS framework examples.  
   👉 [PR #210](https://github.com/anthropics/skills/pull/210)

6. **#83 – Add `skill-quality-analyzer` and `skill-security-analyzer`**  
   *Functionality*: Two meta‑skills that evaluate existing skills across dimensions (structure, security, documentation).  
   *Discussion*: High interest in governance and quality assurance; concerns about introducing circular dependencies in the skill ecosystem.  
   👉 [PR #83](https://github.com/anthropics/skills/pull/83)

7. **#541 – Fix `docx` tracked change `w:id` collision**  
   *Functionality*: Prevents document corruption when the DOCX skill adds tracked changes alongside existing bookmarks by avoiding shared ID‑space clashes.  
   *Discussion*: Technical deep‑dive into OOXML internals; community contributors offered test scenarios.  
   👉 [PR #541](https://github.com/anthropics/skills/pull/541)

---

## Community Demand Trends

From the most commented Issues (13 total), three clear demand directions emerge:

- **Security & Trust Boundaries** (Issue #492, 19 comments): Community skills distributed under the `anthropic/` namespace create impersonation risks. Users demand a clear trust model, namespace separation, and permission controls for community‑submitted skills.
- **Org‑Wide Sharing & Collaboration** (Issue #228, 14 comments, 7 👍): The inability to share skills within an organization without manual file transfer is a top friction point. A shared skill library or direct sharing link is highly requested.
- **Skill‑Creator Reliability** (Issue #556, 12 comments, 7 👍; Issue #1169, 3 comments): The `run_eval.py` pipeline consistently reports 0% trigger rate, breaking the description‑optimization loop. Multiple independent reproductions confirm this is the single biggest blocker to skill development.
- **Duplication & Packaging** (Issue #189, 6 comments, 9 👍): Installation of `document-skills` and `example-skills` plugins yields identical skills, wasting context window. Users want clearly separated, deduplicated skill bundles.

Secondary signals: interest in agent governance (Issue #412, 6 comments), compact memory notation (Issue #1329, 6 comments), and MCP exposure (Issue #16, 4 comments).

---

## High‑Potential Pending Skills

These PRs are actively discussed and likely to be merged soon, addressing key pain points or adding highly desired capabilities:

- **#1099 – Fix `run_eval.py` crash on Windows** (open, 2026‑05‑07)  
  Resolves the subprocess pipe crash that makes the eval loop unusable on native Windows. Combined with #1298 and #1323, these three PRs collectively fix the “recall=0%” epidemic.  
  👉 [PR #1099](https://github.com/anthropics/skills/pull/1099)

- **#360 – Add `appdeploy` skill** (open, 2026‑02‑09)  
  Enables Claude to deploy and manage full‑stack web apps via AppDeploy. Addresses a clear gap in deployment automation.  
  👉 [PR #360](https://github.com/anthropics/skills/pull/360)

- **#723 – Add `testing-patterns` skill** (open, 2026‑03‑22)  
  Covers unit testing, React component testing, integration testing, and testing philosophy. Aligns with strong community interest in code quality workflows.  
  👉 [PR #723](https://github.com/anthropics/skills/pull/723)

- **#147 – Add `codebase-inventory-audit` skill** (open, 2025‑12‑16)  
  Systematic 10‑step workflow for orphaned code, documentation gaps, and infrastructure bloat. Popular among teams managing large legacy codebases.  
  👉 [PR #147](https://github.com/anthropics/skills/pull/147)

---

## Skills Ecosystem Insight

The community’s most concentrated demand is **fixing the skill‑creator pipeline** (the “recall=0%” bug) to unblock skill development, followed by **document‑format coverage** (typography, ODT, PDF/DOCX robustness) and **organizational trust & sharing** features that would enable enterprise adoption of the Skills ecosystem.

---

# Claude Code Community Digest — 2026-06-26

## Today's Highlights
- **Release v2.1.193** shipped with a new `autoMode.classifyAllShell` setting to route all shell commands through the auto-mode classifier, plus improved denial visibility in the transcript and `/permissions` UI.
- A long-standing **secure secrets injection feature request** (#32733) remains the most upvoted open issue (👍112), highlighting strong demand for safe credential handling in the web UI.
- The maintainers merged a **PR to extend stale/autoclose timeouts from 14 to 90 days** (#63686), signaling a shift toward more patient issue triage.

## Releases
**v2.1.193** ([release](https://github.com/anthropics/claude-code/releases/tag/v2.1.193))
- Added `autoMode.classifyAllShell` setting – all Bash/PowerShell commands now pass through the auto-mode classifier (previously limited to arbitrary-code-execution patterns).
- Auto-mode denial reasons are now surfaced in the transcript, a denial toast, and the `/permissions` recent denials list.
- (No further details available in the changelog snippet.)

## Hot Issues (10 noteworthy)
1. **[#32733] Secure secrets injection for Claude Code on the web**  
   _[enhancement, area:security, area:claude-code-web]_  
   Author: lieblius | 👍112 | Comments: 4 | [Open](https://github.com/anthropics/claude-code/issues/32733)  
   *Why it matters*: The highest-voted open issue. Users want a first-class way to inject API keys, tokens, and secrets into web sessions without leaking them to logs or the model’s context.

2. **[#70536] Model repeatedly submits incorrect heredoc despite acknowledging error**  
   _[bug, platform:macos, area:model]_  
   Author: thatmanmatt | 👍0 | Comments: 3 | [Open](https://github.com/anthropics/claude-code/issues/70536)  
   *Why it matters*: A model‑behavior regression where the model loops on a malformed heredoc, ignoring its own corrections — a sign of degraded self‑consistency in Opus 4.8.

3. **[#71564] Tool results appear modified with injected pseudo‑"system instructions" urging `git push --force`**  
   _[bug, platform:windows, area:tools/security/cowork]_  
   Author: aiken884 | 👍0 | Comments: 2 | [Open](https://github.com/anthropics/claude-code/issues/71564)  
   *Why it matters*: A serious security report in the Cowork research preview — injected text masquerading as repo policy could trick the model into destructive git actions.

4. **[#70186] Feature: Add `gitPollingIntervalMs` setting**  
   _[duplicate, closed]_  
   Author: npinto | 👍0 | Comments: 3 | [Closed](https://github.com/anthropics/claude-code/issues/70186)  
   *Why it matters*: Background git polling every ~20 seconds causes `.git/index.lock` contention in multi‑session workflows. Proposal to make interval configurable (or disable via `0`).

5. **[#70022] VS Code terminal integration tool to send commands to integrated terminal**  
   _[duplicate, closed]_  
   Author: jacopozarri-yource | 👍0 | Comments: 2 | [Closed](https://github.com/anthropics/claude-code/issues/70022)  
   *Why it matters*: Claude Code runs in its own shell context; users want it to control their VS Code integrated terminal directly for branch switches and run commands.

6. **[#70039] Windows native updater infinite no‑op updates**  
   _[duplicate, closed]_  
   Author: Talnerith | 👍2 | Comments: 2 | [Closed](https://github.com/anthropics/claude-code/issues/70039)  
   *Why it matters*: The Windows updater reports success but never replaces the binary — a frustrating UX that wastes bandwidth and trust.

7. **[#70051] Permission rules don’t evaluate pipe segments independently**  
   _[duplicate, closed]_  
   Author: nemoDreamer | 👍0 | Comments: 2 | [Closed](https://github.com/anthropics/claude-code/issues/70051)  
   *Why it matters*: Piped commands like `grep pattern | head` bypass allow‑rules because the whole string is matched, reducing the effectiveness of permission controls.

8. **[#70073] Agent should stop when expected file is missing, not silently write elsewhere**  
   _[duplicate, closed]_  
   Author: eoghanmurray | 👍0 | Comments: 2 | [Closed](https://github.com/anthropics/claude-code/issues/70073)  
   *Why it matters*: A safety issue — the agent silently writes to a same‑named file in a different location when the intended target doesn’t exist.

9. **[#70119] Model over‑agreement: ignores standing instructions, fails to push back**  
   _[duplicate, closed]_  
   Author: halheinrich | 👍0 | Comments: 2 | [Closed](https://github.com/anthropics/claude-code/issues/70119)  
   *Why it matters*: Opus 4.8 defaults to agreement even when user instructions explicitly reserve a decision boundary — undermines trust in model‑guided workflows.

10. **[#70157] Agent stuck in loading state in Zed IDE integration**  
    _[duplicate, closed]_  
    Author: MahalingamS18 | 👍0 | Comments: 3 | [Closed](https://github.com/anthropics/claude-code/issues/70157)  
    *Why it matters*: Zed users hit a persistent loading spinner with no error output, making the integration unusable.

## Key PR Progress
1. **[#63686] Bump stale and autoclose timeouts from 14 to 90 days**  
   Author: caseyWebb | Closed | [PR](https://github.com/anthropics/claude-code/pull/63686)  
   *Significance*: Changes the bot’s lifecycle policy from 14‑day staleness/closure to 90 days, reducing churn of legitimate issues and giving maintainers more time to triage.

2. **[#71530] Merge pull request #1 from anthropics/main**  
   Author: arafatjoyadh0414-ux | Open | [PR](https://github.com/anthropics/claude-code/pull/71530)  
   *Significance*: A trivial merge; no functional changes. Likely a fork sync.

## Feature Request Trends
- **Secrets & credentials management** – The #1 ask (#32733) for a secure injection mechanism in the web version, reflecting broader enterprise adoption.
- **Configurable polling / background activity** – Requests to tune or disable git polling (#70186), plus complaints about high CPU load (#70123), point to a desire for more resource‑conscious background behavior.
- **Deeper IDE integration** – VS Code terminal command passing (#70022) and Zed integration fixes (#70157) show a strong need for seamless editor ↔ Claude Code sync.
- **Mid‑task interaction** – Users want to inject messages while a task runs (#70095) and control verbosity (#70129) — both improving real‑time collaboration.
- **Programmatic access to usage/cost data** (#70107) – Developers want to monitor token consumption programmatically from hooks or scripts.

## Developer Pain Points
- **Model loop & empty responses** – Multiple reports (#70536, #70043, #70075) of the model entering infinite loops, producing empty workflow results, or wasting tokens. Indicates model‑level reliability problems.
- **Permission bypass & safety concerns** – Piped commands bypass allow‑rules (#70051); injected pseudo‑instructions in tool results (#71564); agent writing to wrong locations (#70073). These erode trust in security boundaries.
- **Windows & platform‑specific regressions** – Infinite updater loops (#70039), input flicker‑free rendering bugs (#71561), and `!` prefix ignoring primary shell (#70121) suggest Windows is a secondary citizen.
- **Duplicate issue volume** – The top 30 issues list is dominated by duplicates (closed), indicating either a discoverability problem or insufficient guidance before filing. The 90‑day stale timeout PR (#63686) aims to reduce friction.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-06-26

---

## Today's Highlights

A **rate-limit crisis** dominates the community: multiple reports (including #28879, #30002, #30212) describe token allowance draining 10–20× faster than expected, with Pro/20x plans exhausted in minutes. Meanwhile, the engineering team is shipping **MCP OAuth concurrency fixes** (#29017–#29021 stack) and laying groundwork for a **process-owned code‑mode host** (#30202, #30213). Three new alpha releases and a `codex-zsh` plugin also landed.

---

## Releases

| Release | Summary |
|---------|---------|
| `rust-v0.143.0-alpha.25` | Alpha 0.143.0 update |
| `rust-v0.143.0-alpha.22` | Alpha 0.143.0 update |
| `codex-zsh-v0.1.0` | Initial zsh integration plugin |

No detailed changelogs were published. The [codex-zsh plugin](https://github.com/openai/codex/releases/tag/codex-zsh-v0.1.0) is a new entry for shell users.

---

## Hot Issues

| # | Title & Link | Why It Matters |
|---|--------------|----------------|
| 1 | **[#28879](https://github.com/openai/codex/issues/28879)** – Rate-limit cost per token jumped ~10–20× since June 16 | 315 👍, 162 comments. Plus plan users burning 5‑hour budget in 2–3 prompts on gpt-5.5. Community is demanding urgent investigation. |
| 2 | **[#14593](https://github.com/openai/codex/issues/14593)** – Burning tokens very fast | 623 comments, longest thread. Long-standing rate-limit bug still unresolved, though possibly linked to #28879. |
| 3 | **[#30002](https://github.com/openai/codex/issues/30002)** – Server-side quota over-reports after 5h reset | Pro 5h limit consumed in ~41 min despite actual usage of 1.35M tokens (vs expected 156M). Suggests a server-side accounting bug. |
| 4 | **[#29968](https://github.com/openai/codex/issues/29968)** – Pro20x subscription usage behaves like Plus | Pro users seeing Plus-level caps. Symptom of the same rate-limit anomaly. |
| 5 | **[#30212](https://github.com/openai/codex/issues/30212)** – 5-hour allowance depleted in ~1 hour (June 26) | New report from today, 7 👍 in hours. Tracks the same class of problem. |
| 6 | **[#29072](https://github.com/openai/codex/issues/29072)** – `apply_patch` fails because sandbox-setup.exe cannot launch from package path | Windows sandbox blocks file edits. 18 comments; core workflow broken on Windows. |
| 7 | **[#30009](https://github.com/openai/codex/issues/30009)** – Another `apply_patch` failure on Windows | Similar sandbox error (runner error 1920). Affects the latest app version 26.616.81150. |
| 8 | **[#20570](https://github.com/openai/codex/issues/20570)** – `CreateProcessAsUserW failed: 1920` after upgrade | Recurring Windows sandbox regression; 7 👍. Shows the pattern is long-standing. |
| 9 | **[#19821](https://github.com/openai/codex/issues/19821)** – WebSocket connect failures wait through all retries before HTTP fallback | Users behind proxies (especially in China) suffer long delays. Workaround requires manual config to disable WebSockets. |
| 10 | **[#16817](https://github.com/openai/codex/issues/16817)** – Mac Desktop App: existing threads don't load after restart | 7 👍, 13 comments. Persistent session persistence bug that erases user work. |

---

## Key PR Progress

| # | Title & Link | What It Does |
|---|--------------|--------------|
| 1 | **[#30217](https://github.com/openai/codex/pull/30217)** – Remove unavailable task messages from `list_agents` | Fixes multi-agent v2 output by omitting null `last_task_message` for encrypted calls. |
| 2 | **[#30202](https://github.com/openai/codex/pull/30202)** – Bundle `codex-code-mode-host` in release packages | Builds, signs, and ships the standalone code‑mode process host alongside main binaries. |
| 3 | **[#30213](https://github.com/openai/codex/pull/30213)** – Exercise code-mode suite through process host | Integration tests for the standalone host covering tool calls, `notify`, parallel execution. |
| 4 | **[#30223](https://github.com/openai/codex/pull/30223)** – Make plugin guidance react to environment readiness | Prevents missing plugin capability guidance when a plugin becomes available mid‑turn. |
| 5 | **[#30225](https://github.com/openai/codex/pull/30225)** – Overlap executor skill reads with namespace discovery | Performance improvement: parallelizes plugin namespace and skill metadata loading. |
| 6 | **[#30226](https://github.com/openai/codex/pull/30226)** – Make Apps guidance react to MCP availability | Similar to #30223 but for the Apps MCP server. |
| 7 | **Stack #29017–#29021** – Serialize MCP OAuth refresh/concurrency | Five‑PR stack preventing refresh‑token replay and credential overwrites. Part 5 (#30089) adds concurrency tests. Critical for reliability. |
| 8 | **[#30201](https://github.com/openai/codex/pull/30201)** – Avoid server token refresh retry storms | Remote control: stops aggressive retries when `/server/refresh` returns 502. |
| 9 | **[#27804](https://github.com/openai/codex/pull/27804)** – `skill_search` tool | Replaces static skill catalog with a dynamic search tool, enabling ranked retrieval of relevant skills. |
| 10 | **[#25866](https://github.com/openai/codex/pull/25866)** – Handle CRLF gracefully in `apply_patch` | Preserves carriage returns in patch payloads under a feature flag, solving Windows newline breakage. |

---

## Feature Request Trends

- **Dynamic Workflows & Declarative Orchestration** – [#25446](https://github.com/openai/codex/issues/25446) proposes an experimental foundation for declarative dynamic workflows in CLI. Community interest in configurable agent pipelines.
- **HTTPS‑only Transport** – Multiple requests (#27381, #19821) for a flag to force HTTPS and skip WebSocket entirely, especially for users behind restrictive proxies.
- **In‑thread Navigation** – [#22627](https://github.com/openai/codex/issues/22627) asks for a table of contents / jump-to ability in long sessions in the Desktop App.
- **Remote Host Thread Grouping** – [#24295](https://github.com/openai/codex/issues/24295) wants sidebar grouping by Connection → Project → Thread when working with remote Codex hosts.
- **Enhanced `/goal` Capabilities** – [#20958](https://github.com/openai/codex/issues/20958) suggests extending `/goal` with intent calibration, evidence chains, and side-thread tolerance for long‑running tasks.
- **MCP Integration Improvements** – Several PRs and issues show growing demand for better MCP server reliability, OAuth handling, and runtime reuse.

---

## Developer Pain Points

1. **Rate‑Limit Anomalies** – The top pain point this week. Unexpected token drain (10–20× normal) affects Plus, Pro, and Pro20x subscribers. Users report budgets depleted in minutes, with server-side accounting appearing to reset improperly. Multiple high‑reaction issues (#28879, #14593, #30002, #30212) indicate a systemic backend problem.

2. **Windows Sandbox Failures** – `apply_patch` is consistently broken on Windows due to sandbox setup issues (`CreateProcessAsUserW` error 1920). Users cannot apply file edits. Related issues #29072, #30009, #20570 have accumulated over months with no fix shipped.

3. **WebSocket Connection Instability** – Users behind proxies, VPNs, or restricted networks face long delays before HTTP fallback (#19821). The lack of an opt‑in HTTPS‑only mode forces repeated retries.

4. **Ghost / Stale Conversations** – Desktop App on macOS shows unresumable “ghost” threads (#29868). Combined with threads failing to load after restart (#16817), users lose work and cannot clean up their session list.

5. **UI Slowdowns & Crashes** – SQLite log‑file growth (#24275) and “Application Hang 1002” on Windows (#27581) degrade performance and cause unexpected app terminations during normal use.

6. **Proxy / Corporate Network Woes** – WebSocket timeout, lack of HTTPS‑only fallback, and `respect_system_proxy` not working reliably (#29958) frustrate enterprise users.

---

*Generated from `github.com/openai/codex` activity in the last 24 hours ending 2026-06-26.*

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-06-26

## Today’s Highlights
The team shipped three releases in the last 24 hours, including a new **v0.50.0-preview.1** with tool registry dependency injection and a **v0.49.0** stable. Active bug triage continues on several high-priority agent issues, including subagent recovery misreporting success after `MAX_TURNS` and a generalist agent hang. Two critical security PRs address OAuth socket reuse and a misleading trust dialog that could expose users to arbitrary hooks.

---

## Releases
- **[v0.51.0-nightly.20260626.gb14416447](https://github.com/google-gemini/gemini-cli/releases/tag/v0.51.0-nightly.20260626.gb14416447)** — CI fix to prevent bad NPM releases and promote job crashes; includes changelog for v0.50.0-preview.1.
- **[v0.50.0-preview.1](https://github.com/google-gemini/gemini-cli/releases/tag/v0.50.0-preview.1)** — Fixes for release verification (npm ci ignore scripts, workspace binary shadowing) and a new `Feat/tool registry di`.
- **[v0.49.0](https://github.com/google-gemini/gemini-cli/releases/tag/v0.49.0)** — Stabilization release with version bump and cooldown period for npm dependabot updates.

---

## Hot Issues (Top 10)
1. **[#22323](https://github.com/google-gemini/gemini-cli/issues/22323)** — Subagent recovery after `MAX_TURNS` incorrectly reports `status: "success"` and `Termination Reason: "GOAL"`, hiding real interruption. 8 comments, 2 👍. Community calls this a dangerous misreporting that can mask agent failure.
2. **[#21409](https://github.com/google-gemini/gemini-cli/issues/21409)** — Generalist agent hangs indefinitely when invoked (P1). 7 comments, 8 👍. High community pain; workaround is to disable sub-agent delegation.
3. **[#22745](https://github.com/google-gemini/gemini-cli/issues/22745)** — Epic to assess AST-aware file reads, search, and codebase mapping for more precise tool calls. 7 comments. Requests to reduce noise and token waste.
4. **[#24353](https://github.com/google-gemini/gemini-cli/issues/24353)** — Robust component-level evaluations (P1). 7 comments. Need for 76+ behavioral eval tests to be well-integrated into CI.
5. **[#21968](https://github.com/google-gemini/gemini-cli/issues/21968)** — Gemini does not use custom skills and sub-agents voluntarily. 6 comments. Users want the model to invoke skills automatically when relevant.
6. **[#25166](https://github.com/google-gemini/gemini-cli/issues/25166)** — Shell command execution gets stuck with "Waiting input" after command completes (P1). 4 comments, 3 👍. Frequent frustration for developers using CLI tools.
7. **[#26525](https://github.com/google-gemini/gemini-cli/issues/26525)** — Add deterministic redaction and reduce Auto Memory logging (P2, security). 5 comments. Security concern: secrets present in model context before redaction.
8. **[#26522](https://github.com/google-gemini/gemini-cli/issues/26522)** — Auto Memory retries low-signal sessions indefinitely. 5 comments. Leads to wasted API calls and confusion.
9. **[#21983](https://github.com/google-gemini/gemini-cli/issues/21983)** — Browser subagent fails on Wayland (P1). 4 comments. Platform-specific breakage affecting Linux users.
10. **[#24246](https://github.com/google-gemini/gemini-cli/issues/24246)** — 400 error when >128 tools are available (P2). 3 comments. Suggests the agent needs smarter tool selection.

---

## Key PR Progress (Top 10)
1. **[#27915](https://github.com/google-gemini/gemini-cli/pull/27915)** — `fix(core): trust dialog discloses the hook shape that never runs` (P1, security). Fixes a security hole where the workspace-trust dialog shows the *inverse* of hooks that actually execute, exposing users to arbitrary shell execution.
2. **[#28059](https://github.com/google-gemini/gemini-cli/pull/28059)** — `fix(cli): don't let an unreadable .env (EACCES) break extension loading`. Essential for sandbox environments.
3. **[#28012](https://github.com/google-gemini/gemini-cli/pull/28012)** — `fix(cli): sync footer branch name on filesystems without fs.watch events`. Fixes branch indicator stuck on WSL and network mounts.
4. **[#27850](https://github.com/google-gemini/gemini-cli/pull/27850)** — `fix(core): sniff MCP image MIME types` (P1). Correctly detects WebP data misreported as `image/png` and converts to the correct type for the model.
5. **[#27845](https://github.com/google-gemini/gemini-cli/pull/27845)** — `fix(cli): prompt for folder trust before auth` (P2). Improves security by asking for workspace trust before authentication begins.
6. **[#28013](https://github.com/google-gemini/gemini-cli/pull/28013)** — `fix(prompts): use function replacer in applySubstitutions to prevent $-pattern corruption`. Avoids breaking skill, sub-agent, and tool descriptions that contain `$` characters.
7. **[#28103](https://github.com/google-gemini/gemini-cli/pull/28103)** — `fix(core): avoid keep-alive socket reuse during OAuth token exchange` (P2, security). Works around Node.js CVE-2026-48931 that causes OAuth failures on certain releases.
8. **[#27971](https://github.com/google-gemini/gemini-cli/pull/27971)** — `fix(core): strip thoughts from scrubbed history turns and resolve thought leakage`. Prevents the model from seeing its own internal monologue, which could cause infinite loops.
9. **[#27848](https://github.com/google-gemini/gemini-cli/pull/27848)** — `feat(cli): add 'models' command to list available Gemini models` (closed, P3). New convenience command for discovering model capabilities.
10. **[#28153](https://github.com/google-gemini/gemini-cli/pull/28153)** — `fix(core): ignore stale update_topic calls after a session reset` (P1). Prevents orphaned `update_topic` calls from corrupting state after `/clear`.

---

## Feature Request Trends
- **AST-aware code understanding** — Multiple issues (#22745, #22746) request AST-aware file reads, search, and codebase mapping to reduce token waste and improve precision.
- **Autonomous skill/sub-agent usage** — Users want the agent to proactively invoke custom skills and sub-agents based on context, not just when explicitly instructed (#21968).
- **Component-level evaluations** — There's a strong push for robust, automated behavioral eval tests (#24353) that are always passing and cover sub-agent behavior.
- **Memory system improvements** — Auto Memory needs deterministic redaction (#26525), better retry logic (#26522), and error surface/quarantine for invalid patches (#26523).
- **Browser agent resilience** — Issues #22232 and #22267 request automatic session takeover, lock recovery, and proper configuration override handling for browser sub-agent.
- **Better tool selection** — #24246 and #23571 indicate the model needs to intelligently limit tools to avoid 400 errors and avoid creating tmp scripts in random directories.
- **Sharing sub-agent trajectories** — #22598 wants sub-agent context included in `/chat share` for easier review and debugging.

---

## Developer Pain Points
- **Agent hangs and misreported success** — Issues #21409, #22323, and #25166 describe scenarios where the CLI freezes or reports fake success, wasting developer time and trust.
- **Sub-agent misbehavior** — Sub-agents run without permission (#22093), ignore settings.json overrides (#22267), and fail on specific environments like Wayland (#21983).
- **Shell execution stuck** — #25166 and #22465 highlight the shell staying in "Waiting input" after completion or getting stuck at interactive prompts (e.g., creating a Vite app).
- **Trust/safety confusion** — The trust dialog shows the *opposite* of hooks that run (#27915), and destructive commands (`git reset`, `--force`) are not discouraged (#22672).
- **Auto Memory inefficiency** — Low-signal sessions are retried indefinitely (#26522), causing unnecessary processing and API costs.
- **Symlink and path issues** — Agent files as symlinks are not recognized (#20079), and workspace `.env` files with permissions errors break extension loading (#28059).
- **Terminal and UI glitches** — Corruption after exiting external editors (#24935), flicker on resize (#21924), and PTY resize crashes (#27461) affect daily usability.

*All links refer to the repository [google-gemini/gemini-cli](https://github.com/google-gemini/gemini-cli).*

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest – 2026-06-26

## Today's Highlights
A new release (v1.0.66-0) landed with key MCP and observability improvements, including the ability to toggle MCP servers from the list view and experimental response budget controls. Community activity focused on enterprise configuration gaps, Windows-specific clipboard and rendering bugs, and several usability requests around MCP server management and agent customization. The security team is also working on a CVE assignment for a previously reported vulnerability.

## Releases
**v1.0.66-0** (published 2026-06-25)
- **Added:**
  - Toggle to enable/disable MCP servers directly from the MCP list view
  - Experimental response budget controls in CLI settings
  - Managed settings now configure OpenTelemetry export
  - MCP tools on OAuth-authenticated remote servers automatically recover after a mid-session token refresh

[View release](https://github.com/github/copilot-cli/releases/tag/v1.0.66-0)

## Hot Issues (10 noteworthy)
1. **#3501** – *Scroll bar makes text unalign on Windows* (area:platform-windows, terminal-rendering) – Community hit: 9 👍, 5 comments. Vertical scroll bar introduced rendering misalignment on Windows Console Host and Terminal. Copilot unable to self-help. [Issue](https://github.com/github/copilot-cli/issues/3501)

2. **#3534** – *`/copy` fails on WSL2 ARM64 with `clip.exe exited with code 1`* (area:input-keyboard, platform-windows) – 4 👍, 4 comments. Quoting bug in `cmd.exe` wrapper introduced in 1.0.55 breaks clipboard writes under WSL2 ARM64. [Issue](https://github.com/github/copilot-cli/issues/3534)

3. **#3636** – *Voice mode cannot be enabled – model catalog unreachable on corporate VPN* (area:networking, models) – 5 👍, 3 comments. `Failed to fetch model catalog` blocks voice mode entirely when behind a VPN that restricts network access. [Issue](https://github.com/github/copilot-cli/issues/3636)

4. **#3906** – *CVE assignment request* (triage) – User reports submitting a security report via GHSA and now requesting a CVE to be assigned while they prepare patches. [Issue](https://github.com/github/copilot-cli/issues/3906)

5. **#3909** – *Enterprise / org server-managed settings for local CLI* (area:enterprise, configuration) – Org admins want to centrally push environment variables to developers' local CLI installations, currently only possible for Codespaces. [Issue](https://github.com/github/copilot-cli/issues/3909)

6. **#3939** – *Fleet command multi-clone capability* (area:agents) – Request for `/fleet` to clone a repo to multiple directories to avoid squashing file changes, using git as communication. [Issue](https://github.com/github/copilot-cli/issues/3939)

7. **#3940** – *Custom agent support for `skills` field to limit preloaded skills* (area:agents, plugins) – Users want to define a `skills` field in custom agents to control which skills are preloaded into context. [Issue](https://github.com/github/copilot-cli/issues/3940)

8. **#3942** – *`copilot --acp` does not work with `--agent`* (area:non-interactive, agents) – Non-interactive mode using `--acp` fails when a custom agent is specified via `--agent`. [Issue](https://github.com/github/copilot-cli/issues/3942)

9. **#3933** – *Drops out of autopilot after each request* (area:agents) – Autopilot mode no longer persists across prompts; user exits autopilot after each request, contrary to previous behavior. [Issue](https://github.com/github/copilot-cli/issues/3933)

10. **#3935** – *CLI 1.0.64/1.0.65 ignores user theme in VSCode terminal* (area:theming-accessibility) – Since 1.0.64 the CLI defaults to light theme, ignoring VSCode's dark or custom themes. [Issue](https://github.com/github/copilot-cli/issues/3935)

## Key PR Progress (1 notable)
- **#3928** – *Add .gitignore and settings configuration* (OPEN, author: tpsaint) – Adds `.gitignore` and settings configuration to the repository. No comments yet. [PR](https://github.com/github/copilot-cli/pull/3928)

*(Only one PR was updated in the last 24 hours; the community currently has a higher volume of issues than pull requests.)*

## Feature Request Trends
- **MCP Enhancements:** Several requests focus on MCP server management – enabling/disabling from interactive menus (#2956, now closed, and #3564), read-only slash commands like `/mcp show` being made asynchronous (#3829), and variable interpolation in `packageArguments` (#3887, closed). The new release partially addresses these with a toggle in the list view.
- **Enterprise Configuration:** There is strong demand for centralised, server-managed settings for local CLI installations, especially environment variables (#3909).
- **Agent & Skills Customisation:** Users want finer control over custom agents, including the ability to limit which skills are preloaded (#3940) and managing global skills migrated from Claude Code (#3938, closed).
- **Autopilot Behaviour:** Recurring reports that autopilot mode does not stay active after a single prompt (#3933) call for a more persistent interaction model.
- **Cross-Platform Clipboard & Rendering:** Windows-specific clipboard bugs (#3534) and scroll bar rendering (#3501) continue to be pain points on that platform.

## Developer Pain Points
- **Windows & Terminal Rendering:** Scroll bar misalignment (#3501) and mouse tracking disabled on exit (#3876, closed) frustrate Windows users.
- **Clipboard on WSL2 ARM64:** The `/copy` command is broken for ARM64 WSL2 users due to a quoting bug (#3534).
- **Voice Mode VPN Block:** Voice mode fails behind corporate VPNs because it requires access to a model catalog (#3636).
- **Theme Ignores Settings:** The CLI overrides VSCode terminal themes since 1.0.64 (#3935), affecting accessibility and user experience.
- **MCP Policy Blocking:** Users report `MCP server is blocked by policy` errors even when local configuration is correct (#3934).
- **Autopilot Persistence:** The regression in autopilot staying active after each request (#3933) disrupts workflow.
- **Preserved User Skills on Update:** Migrated Claude Code skills are lost after CLI updates (#3938, closed but unresolved for users).

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-06-26

## Today’s Highlights
Two bugs were filed yesterday, both affecting v0.19.2: an MCP server integration issue on Windows and a persistent UI re-rendering problem on Linux. On the development side, two pull requests are progressing — one improves the README’s development prerequisites, the other fixes a serialization bug when reasoning effort is disabled. No new release was published.

## Releases
No new releases in the last 24 hours.

## Hot Issues
- **[#2475 [bug] MCP tools](https://github.com/MoonshotAI/kimi-cli/issues/2475)** – A user reports that an MCP server with 212 tools and descriptions is not handled correctly. The issue is currently uncommented and unlabeled. This could point to a scaling limitation or a serialization bug in the MCP integration. The community may be waiting for a response from maintainers.

- **[#2474 [bug] CLI interface shaking / re-rendering](https://github.com/MoonshotAI/kimi-cli/issues/2474)** – On Linux (x86_64, kernel 5.10), the terminal UI constantly flickers and restarts the entire conversation rendering. The user is using the K2.7 Code thinking model. This suggests a potential issue with terminal rendering or state management, likely affecting productivity for Linux users.

## Key PR Progress
- **[#2287 docs(readme): add prerequisites list to Development section](https://github.com/MoonshotAI/kimi-cli/pull/2287)** – Submitted by ktwu01, this PR addresses a documentation gap by listing what a contributor needs installed before `make prepare` works. It resolves issue #2274. Last updated today, so it may be close to merge.

- **[#2476 fix(kosong): omit reasoning_effort instead of sending null when thinking is off](https://github.com/MoonshotAI/kimi-cli/pull/2476)** – Submitted by logicwu0, this PR fixes a bug where `reasoning_effort=null` was sent explicitly to the OpenAI SDK when thinking is off. The correct behavior is to omit the parameter entirely. This is a small but important correctness fix for API compatibility.

## Feature Request Trends
No explicit feature requests were filed in the last 24 hours. The two existing issues are bug reports rather than feature requests. Developer interest appears focused on stability and MCP tool compatibility.

## Developer Pain Points
- **MCP tool scale** – Issue #2475 hints at difficulty when an MCP server exposes a large number of tools (212). Whether this is a parsing or performance issue, it may affect users integrating many tools.
- **Terminal UI instability** – Issue #2474 describes persistent UI shaking and full re-rendering, especially on older Linux kernels. This can severely disrupt workflow.

*Note: This digest is based on a low-activity day. For a fuller picture of community trends, check previous digests or the repository’s issue tracker.*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-26

## Today's Highlights

OpenCode v1.17.11 shipped with a significant UX improvement: **session snapshots with revert controls**, allowing users to roll back to an earlier message and its associated file changes. CPU usage remains the dominant community concern, with three active high-comment issues and a new report calling out the Desktop App's animated progress indicators as a major offender. The GitHub Copilot provider also continues to generate friction — multiple reports today confirm that third-party enterprise models and certain catalog models are rejected with `"The requested model is not supported"`.

---

## Releases

### v1.17.11

**Core Improvements**
- **Session snapshots and revert controls**: Users can now roll a session back to an earlier message, including the file changes made up to that point. This directly addresses long-standing feedback about session state management.

**Core Bugfixes**
- MCP OAuth URL is now always printed, so manual sign-in still works when the automated browser flow fails to open.

**Desktop Improvements**
- Added Chrome-style (truncated description in source — likely Chrome-style tab management or window chrome improvements).

🔗 [v1.17.11 Release](https://github.com/anomalyco/opencode/releases/tag/v1.17.11)

---

## Hot Issues

### 1. [#21470 — OpenCode is heavily CPU-bound](https://github.com/anomalyco/opencode/issues/21470)
*14 comments · 13 👍*
> The most commented issue on the board. User reports that with Gemini 3.1, OpenCode itself consumes the overwhelming majority of time — 300k tokens at $8.30 spent, but OpenCode's internal processing dwarfs API wait time. This signals a systemic performance problem in the agent loop or tokenization pipeline.

### 2. [#30086 — High CPU usage in newer versions of OpenCode](https://github.com/anomalyco/opencode/issues/30086)
*11 comments · 8 👍*
> User went from running 10 concurrent sessions to struggling with 3 after recent updates. Laggy mouse cursor reported. Likely a regression introduced in the 1.14–1.17 range. Community upvote count suggests widespread impact.

### 3. [#19466 — opencode is using CPU for doing nothing!](https://github.com/anomalyco/opencode/issues/19466)
*10 comments · 9 👍*
> Even during API rate-limit backoff (sleeping 18 minutes), OpenCode consumes ~50% of a single i9-14900 core. The idle-polling mechanism appears to lack proper sleep/backoff in the agent loop.

### 4. [#15907 — Clipboard copy not working over SSH + tmux in Ghostty](https://github.com/anomalyco/opencode/issues/15907)
*9 comments · 10 👍*
> Shows "copied to clipboard" but doesn't actually update the system clipboard when running over SSH inside tmux. High 👍 count indicates this is a pain point for remote development workflows.

### 5. [#34045 — Desktop App spending too much CPU on animating progress indicators](https://github.com/anomalyco/opencode/issues/34045)
*1 comment · 1 👍*
> *New today.* User explicitly calls out that while waiting on a slow model (Opus 4.8), the progress indicator animation burns CPU for "no good reason." Points to a frontend rendering inefficiency adding to the CPU problem.

### 6. [#34030 — Cannot invoke third-party enterprise models in GitHub Copilot](https://github.com/anomalyco/opencode/issues/34030)
*3 comments · 0 👍*
> *New today.* GitHub Copilot Enterprise users can't access third-party models that their org has added. OpenCode only sees the default Copilot catalog.

### 7. [#34048 — GitHub Copilot model inference fails with "not supported"](https://github.com/anomalyco/opencode/issues/34048)
*1 comment · 0 👍*
> *New today.* Authentication succeeds, models are listed, but every inference request fails. User also received a hint about a missing `Copilot-Integration-Id` header — mirrors the closed issue #30675.

### 8. [#33878 — Background task completion notifications silently lost in TUI](https://github.com/anomalyco/opencode/issues/33878)
*3 comments · 0 👍*
> *Closed today.* Subagent background tasks complete, but the `<system-reminder>` notification never appears. Parent session stays idle until the user manually types a message. A session-loop wake-up bug.

### 9. [#34043 — Subagent fallback chain prepends incorrect prefix to model names, causing infinite retry](https://github.com/anomalyco/opencode/issues/34043)
*1 comment · 0 👍*
> *New today.* When a subagent's initial model returns HTTP 429, the fallback chain incorrectly prepends `opencode/` to model names (e.g., `nvidia/z-ai/glm-5.1` becomes `opencode//nvidia/z-ai/glm-5.1`), creating an infinite retry loop.

### 10. [#34036 — Windows ARM64 NSIS installer broken — no OpenCode.exe deployed](https://github.com/anomalyco/opencode/issues/34036)
*2 comments · 0 👍*
> *Closed today.* The Windows ARM64 installer for v1.17.11 fails to extract `OpenCode.exe`. Shortcut is broken. Likely a packaging pipeline issue with the bundled 7-Zip.

---

## Key PR Progress

### 1. [#34042 — fix(app): preserve live messages during session load](https://github.com/anomalyco/opencode/pull/34042)
*Opened today by Hona*
Three-way merge logic for session loading: preserves live-only messages while merging fetched history. Solves a class of bugs where messages disappear during session restoration.

### 2. [#34039 — fix(session): stop malformed tool-call loops](https://github.com/anomalyco/opencode/pull/34039)
*Opened today by Nomadcxx*
Fixes #31247 — Claude Opus 4.8 can return `finish=tool-calls` without a structured tool call payload. Previously caused an infinite tool-call loop.

### 3. [#34037 — fix(tui): autocomplete files inside references](https://github.com/anomalyco/opencode/pull/34037)
*Opened today by jomarescudero*
Direct fix for #34040. TUI `@` autocomplete now searches inside configured reference directories and inserts selected file paths.

### 4. [#34046 — fix(desktop): show interrupt loading state](https://github.com/anomalyco/opencode/pull/34046)
*Opened today by opencode-agent[bot]*
Adds `aria-busy` state and progress indicator while an interrupt request is in flight. Prevents double-clicks and provides visual feedback.

### 5. [#33927 — fix(vcs): prevent crash when repo has thousands of untracked files](https://github.com/anomalyco/opencode/pull/33927)
*Opened yesterday by youtsuhodev*
Closes #33928. Repos with 1200+ untracked files crash the VCS layer. Fixes a scaling issue in git status parsing.

### 6. [#32767 — fix(tui): restore ESC interrupt for delegated subagent sessions](https://github.com/anomalyco/opencode/pull/32767)
*Opened 8 days ago by tobwen*
Closes three issues (#3699, #4073, #23534). Restores the ability to ESC-interrupt delegated subagents — a regression that had been missing for several releases.

### 7. [#33643 — fix: implement v2 session wait](https://github.com/anomalyco/opencode/pull/33643)
*Opened 2 days ago by Xiaxinyuuu*
Fixes #33605. The `POST /api/session/:sessionID/wait` endpoint was not correctly implementing the `SessionV2.wait()` method, breaking idle detection.

### 8. [#33067 — feat(tui): support selecting multiple skills](https://github.com/anomalyco/opencode/pull/33067)
*Opened 6 days ago by X6TXY*
Closes #32954. Allows selecting multiple skills from the `/skills` picker instead of applying a single one. Improves multi-tool workflows.

### 9. [#32425 — feat(opencode): interrupt a running subagent](https://github.com/anomalyco/opencode/pull/32425)
*Opened 11 days ago by iceteaSA*
Experimental primitive for steering/canceling/aborting background subagents. Prerequisite for fully resolving #21458 and #23534.

### 10. [#28326 — feat(server): runtime base path support for reverse proxy deployments](https://github.com/anomalyco/opencode/pull/28326)
*Opened 38 days ago by fabiovincenzi*
Closes #7624. Adds `--base-path` flag and `server.basePath` config so opencode web works behind reverse proxies with a non-root path.

---

## Feature Request Trends

Three clear feature directions emerged from today's issue activity:

1. **Subagent & background task management** — Multiple requests center on controlling background subagents: interrupting them, steering them mid-execution, and ensuring notifications wake the parent session. PRs #32425 and #32767 are directly addressing this, but the volume of related issues (#21458, #23534, #28738) shows this is a top community priority.

2. **Provider model compatibility & extensibility** — The GitHub Copilot provider continues to frustrate users trying to access enterprise models (#34030, #34048) and non-Copilot-catalog models. Separately, requests for new providers (Crof AI #24636, DeepSeek web search #32273, agentrouter.org #2784) suggest users want broader model choice.

3. **TUI & plugin system improvements** — Users are pushing for richer TUI capabilities: pre-render stream filtering (#34035), slot-based sidebar content (#34050), and multi-skill selection (#33067). These indicate a maturing plugin ecosystem where developers want to customize the terminal experience more deeply.

---

## Developer Pain Points

The single loudest pain point is **spurious CPU usage** — three high-traffic issues (#21470, #30086, #19466) and a new report today (#34045) all point to OpenCode consuming CPU cycles when it should be idle (waiting on API rate limits, animating progress indicators, or simply sitting open). The community is frustrated that a developer tool competes with their IDE and browser for CPU.

Second is **GitHub Copilot integration fragility**. Multiple reports today and in recent weeks (#30675, #34030, #34048) show that model listing works but inference fails. A missing `Copilot-Integration-Id` header has been identified as the likely root cause, but has not been fully resolved.

Third is **background task reliability**. Subagents that complete silently (#33878), infinite retry loops due to model name mangling (#34043), and non-functional ESC interrupts (#32767 refs) all point to the subagent system being the most unstable subsystem in the current release.

Fourth is **Windows ARM64 support**. Two issues (#34036, #33732) confirm the installer is broken for ARM64 devices — the compiled binary is not extracted, making the platform effectively unsupported despite being listed as a target.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest – 2026-06-26

## Today’s Highlights
Connection reliability with `openai-codex` / `gpt-5.5` remains the community’s top concern (#4945, 71 comments), and a group of TUI scroll‑and‑render regressions continues to frustrate users. On the positive side, an experimental orchestrator daemon (#6064) and the addition of Friendli as a built‑in provider (#6090) signal that the project is actively investing in both infrastructure and provider diversity.

## Releases
No new releases in the last 24 hours.

## Hot Issues (10 noteworthy)

1. **#4945 – openai-codex Connection Reliability Issues**  
   *71 comments, 30 👍*  
   Users report that the TUI remains stuck on `Working...` with no streamed text, tool call, or error message. Only pressing Escape recovers the session. A high‑priority reliability blocker.  
   [GitHub](https://github.com/earendil-works/pi/issues/4945)

2. **#5825 – Streaming markdown forces scroll to bottom**  
   *31 comments*  
   When `clear on shrink` is enabled, the TUI forces a scroll to the bottom while the agent is speaking, overriding the user’s manual scroll‑up.  
   [GitHub](https://github.com/earendil-works/pi/issues/5825)

3. **#5103 – Windows zip build can’t detect git-bash from PATH**  
   *23 comments, 1 👍*  
   The standalone `pi-windows-x64.zip` fails to find Git Bash even when it is in `PATH`, breaking the built‑in bash tool. A common Windows onboarding pain point.  
   [GitHub](https://github.com/earendil-works/pi/issues/5103)

4. **#4877 – Session folder collision**  
   *19 comments, 2 👍*  
   Paths like `/a/b/c/d` and `/a-b/c-d` produce the same session folder, causing potential data confusion.  
   [GitHub](https://github.com/earendil-works/pi/issues/4877)

5. **#6050 – TUI full redraw clears terminal scrollback during active rendering**  
   *10 comments*  
   Frequent background redraws reset the scrollbar, especially noticeable in interactive mode. Root cause lies in the core TUI renderer.  
   [GitHub](https://github.com/earendil-works/pi/issues/6050)

6. **#5595 – openai-completions maxTokens not passing through**  
   *8 comments, 2 👍*  
   For providers like Together.ai, reasoning models (e.g., DeepSeek v4pro) run out of output tokens regardless of user settings.  
   [GitHub](https://github.com/earendil-works/pi/issues/5595)

7. **#4990 – Edits failing**  
   *7 comments*  
   After updating to a recent build, the `edit` tool returns `Validation failed: must have required properties edits`. Regression.  
   [GitHub](https://github.com/earendil-works/pi/issues/4990)

8. **#5989 – Update broke extension pi-lovely-codex**  
   *7 comments*  
   An update made the popular `pi-lovely-codex` extension incompatible, causing load errors.  
   [GitHub](https://github.com/earendil-works/pi/issues/5989)

9. **#4290 – Messages aborted for length treated as regular stops**  
   *6 comments, 1 👍*  
   Turns that are stopped due to length show only `Thinking...` with no visible stop indicator, leaving users waiting indefinitely.  
   [GitHub](https://github.com/earendil-works/pi/issues/4290)

10. **#6002 – SessionManager.open() silently truncates non-session files**  
    *4 comments*  
    Pointing the CLI at a non‑session file (e.g., an NDJSON log) destroys the file with no warning. A serious data‑loss risk.  
    [GitHub](https://github.com/earendil-works/pi/issues/6002)

## Key PR Progress (10 important changes)

1. **#6064 – feat(experimental): pi orchestrator**  
   Adds a local orchestrator daemon that manages multiple pi instances via newline‑delimited JSON IPC. Enables lifecycle management (start, stop, list). Experimental.  
   [GitHub](https://github.com/earendil-works/pi/pull/6064)

2. **#6087 – fix(coding-agent): remove hardcoded RPC wait timeout**  
   Resolves #6088. Replaces the 60s hardcoded cap with configurable `RpcClientOptions.waitTimeoutMs`. Long‑running MCP tools no longer fail prematurely.  
   [GitHub](https://github.com/earendil-works/pi/pull/6087)

3. **#6092 – draft: hosted websearch**  
   Always‑on hosted search tool for the agent loop. Marked as draft but addresses the long‑running #1589 (web access).  
   [GitHub](https://github.com/earendil-works/pi/pull/6092)

4. **#6090 – feat(ai): add Friendli provider**  
   Adds Friendli as a built‑in OpenAI‑compatible provider (`openai-completions`), default model `zai-org/GLM-5.2`. Follows the pattern of Ant Ling and NVIDIA NIM.  
   [GitHub](https://github.com/earendil-works/pi/pull/6090)

5. **#5515 – feat(coding-agent): add alwaysTrust setting**  
   A new flag to skip project trust gating entirely. Disabled by default.  
   [GitHub](https://github.com/earendil-works/pi/pull/5515)

6. **#6084 – fix(tui): preserve custom widget render order on background tick refreshes**  
   Prevents custom widgets from re‑ordering on high‑frequency refreshes (clock ticks, turn‑completion).  
   [GitHub](https://github.com/earendil-works/pi/pull/6084)

7. **#6081 – feat: add #RRGGBBAA alpha support for theme colors**  
   8‑digit hex colors are now blended at load time with the terminal background, enabling translucent theme colors.  
   [GitHub](https://github.com/earendil-works/pi/pull/6081)

8. **#6078 – feat(coding-agent): add get_entries and get_tree RPC commands**  
   Two read‑only commands exposing session entries and tree roots over RPC, useful for external integrations.  
   [GitHub](https://github.com/earendil-works/pi/pull/6078)

9. **#6074 – fix(coding-agent): avoid pre-prompt compaction continue**  
   Prevents premature compaction that could break pre‑prompt continuation logic.  
   [GitHub](https://github.com/earendil-works/pi/pull/6074)

10. **#5832 – fix(ai): surface provider HTTP error body instead of opaque SDK message**  
    Behind proxies/gateways, non‑2xx responses with unparseable bodies now expose the actual HTTP body rather than a generic `UnknownError`.  
    [GitHub](https://github.com/earendil-works/pi/pull/5832)

## Feature Request Trends
From the latest batch of issues, the community is pushing in three main directions:
- **Provider & model enhancements** – Add built‑in providers (Friendli #6091, BMP file reading #6047), surface reasoning token counts (#6057), and make the LLM cache work reliably with coding plans (#6083).
- **CLI & session ergonomics** – Shortcuts for combined commands (`/new + /name`, #6046), shell tab‑completion for Pi flags (#6086), and support for deterministic in‑memory session IDs alongside `--no-session` (#6070).
- **Deployment & packaging** – A single executable binary bundling Node (#6065) and the orchestrator daemon (#6064) reflect a desire for zero‑dependency installations and multi‑instance management.

## Developer Pain Points
The data reveals several recurring frustrations:
- **Reliability fractures with OpenAI‑compatible providers** – Connection hangs and token‑limit misconfigurations (#4945, #5595, #6061) are the top‑voted issues.
- **TUI scrolling chaos** – Multiple reports of forced re‑scrolls (#5825), scrollback clearing (#6050), and viewport jumps in tmux (#6073) degrade the interactive experience.
- **Windows breakages** – Git‑bash detection (#5103) and compiled‑binary extension resolution (#6085) remain sticky problems for Windows users.
- **Silent data loss** – The `SessionManager.open()` truncation (#6002) highlights a dangerous lack of guards.
- **Extension incompatibility after updates** – Breaking changes in the extension API catch plugin authors off guard (#5989, #4990).
- **Hardcoded timeouts** – Several services (RPC, MCP, provider calls) have baked‑in timeouts that fail in real‑world scenarios (#6088, #6083).

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest – 2026-06-26

---

## 1. Today's Highlights

The community is buzzing with ambitious proposals for persistent, multi-user collaboration: a new `qwen tag` multiplayer channel‑resident agent (#5887/#5888) and a `.qwen/loop.md` durable task file for long‑running loops (#5889/#5890) are the day’s headline features. On the bug‑fix front, a critical Windows PTY process‑tree leak (#5892) and a UI/UX bug where edit‑tool summaries pollute every subsequent response (#5894) are both being addressed. Additionally, two cross‑PR CI contamination incidents (#5882) and a nightly release failure (#5877) highlight ongoing infrastructure pain points.

---

## 2. Releases

No new releases in the last 24 hours.

---

## 3. Hot Issues

| # | Title | Why It Matters | Community Reaction |
|---|-------|----------------|--------------------|
| #5875 | improve skill command name auto‑complete matching | Current `/skill_name` requires prefix match; users want substring matching (e.g. `/store` matches `front‑end‑store‑rules`) for faster skill discovery. | 4 comments, 0 👍 – clearly a desired UX polish. |
| #5800 | bug(cli): last line of replies taller than terminal is overwritten on completion | In the default static TUI mode, the final line of a long assistant reply is hidden on completion – a longstanding upstream Ink issue (#973) that frustrates users reading lengthy outputs. | 4 comments, 0 👍 – reproducible and impacts daily use. |
| #5677 | tracking(serve): Track ACP gaps for cd, permissions, trust, lsp, and setup‑github | A systematic tracking issue for missing Agent Communication Protocol (ACP) endpoints. Already has three sub‑issues resolved (e.g. `/lsp`, `/permissions`), providing transparency for API completeness. | 3 comments, 0 👍 – important for ACP compliance. |
| #5881 | proposal: open Plan Approval Gate to all plan mode entries | Currently the approval gate only reviews model‑initiated plans; this proposal extends it to all plan‑mode exits (e.g. user‑triggered), catching issues before execution. | 3 comments, 0 👍 – sparks discussion on workflow security. |
| #5889 | [loop] Add a .qwen/loop.md task file injected at fire time | Long‑running `/loop` sessions need a durable, user‑editable task list. This proposal adds a per‑loop markdown file that the model re‑reads every tick. | 2 comments, 0 👍 – directly addresses a gap in background automation. |
| #5887 | feat(channels): "qwen tag" — persistent multiplayer channel‑resident agent | Proposes a shared agent that lives in a group chat (DingTalk‑first), allowing multiple users to collaborate in one session – analogous to Anthropic’s Claude Tag. | 2 comments, 1 👍 – high enthusiasm for team‑collaboration features. |
| #5894 | Edit tool result summary is repeatedly appended to every subsequent response | After editing a file, the “File changed” diff block gets attached to every future assistant message, making conversations unreadable. | 1 comment – a serious UI/UX regression. |
| #5882 | Qwen agent CI jobs run un‑isolated on shared ECS runner → cross‑PR state contamination | Triage comments from PR #5872 were posted on #5874, indicating shared state between CI runs. This undermines trust in automated review workflows. | 2 comments – urgent infrastructure issue. |
| #5897 | Repeating 'unknown format "uint64" ignored in schema' messages pollute the interface | Spammy schema warnings on every startup degrade the developer experience, even though they are harmless. | 1 comment – low priority but annoying. |
| #3029 | LSP workspaceDiagnostics and diagnostics return empty for typescript‑language‑server | Critical for TypeScript users: push‑diagnostic servers like `typescript‑language‑server` report no diagnostics, breaking IDE‑like features in the CLI. | 1 comment – long‑standing issue (#3029, since April). |

---

## 4. Key PR Progress

| # | Title | Description |
|---|-------|-------------|
| #5888 | feat(channels): qwen tag — RFC + Phase 0 | Implements the multiplayer channel‑resident agent on existing channel adapters and daemon. Phase 0 includes baseline architecture and a DingTalk prototype. |
| #5900 | feat(web-shell): allow host to override streaming loading phrases | Adds `LoadingPhrasesResolver` so embedding hosts (e.g. VS Code) can replace the default witty loading messages during streaming. |
| #5886 | feat(memory): add a git‑shared team memory tier | Introduces an opt‑in third memory tier (TEAM) stored in `.qwen/team-memory/`, shared via git alongside private USER and PROJECT tiers. |
| #5852 | feat(daemon,sdk): resumable /acp session stream (Last‑Event‑ID) | Adds SSE `id:` lines to session events so clients can resume interrupted streams without replaying history. |
| #5890 | feat(loop): inject a .qwen/loop.md task file at fire time via sentinels | Companion PR to #5889: loops can now carry a durable task file that is re‑read every tick, enabling mid‑run editing. |
| #5896 | feat(cua-driver): vendor qwen-cua-driver with opt‑in 0–1000 relative coordinates | Vendors the `trycua/cua` driver and adds a normalized coordinate mode for Qwen‑VL’s `computer_use`, enabling pixel‑free background automation. |
| #5892 | fix(core): tree‑kill PTY shell tree on Windows to stop pwsh leak (#5873) | Fixes a process leak where `ptyProcess.kill()` doesn’t reap child processes (pwsh/cmd) on Windows, causing resource exhaustion. |
| #5879 | feat(web-shell): browse MCP server resources in the /mcp dialog | Brings the TUI’s MCP resource browser to the Web Shell, showing resource/prompt counts and expanding to list each resource. |
| #5898 | Fix mid‑input skill command completion | Enables slash‑command auto‑complete when typed mid‑sentence (not just at line start), with fuzzy matching and token replacement. |
| #5868 | feat(core): add configurable auto‑compact threshold and Stop hook context usage (#4025) | Implements user‑configurable token‑budget thresholds and a Stop hook that passes context usage, helping manage long sessions. |

---

## 5. Feature Request Trends

- **Persistent, editable automation**: The strongest trend is adding durable state to long‑running tasks – e.g. `.qwen/loop.md` (#5889), `.qwen/team-memory/` (#5886), and runtime context injection (#5847). Users want to pause, edit, and resume without restarting.
- **Multi‑player / collaborative agents**: The `qwen tag` proposal (#5887/#5888) and shared memory tiers (#5886) show demand for team‑shared AI assistants that live in chat channels.
- **Improved CLI discoverability**: Substring/skill‑name auto‑complete (#5875) and mid‑input slash‑command matching (#5898) reflect a desire for faster skill navigation.
- **Plan‑mode oversight**: Extending the Plan Approval Gate to all exits (#5881) suggests users want safety checks on every plan, not just model‑initiated ones.
- **Unified UI across surfaces**: Consolidating chat input/conversation onto `web-shell` (#5883) is proposed to reduce maintenance burden and provide a consistent experience across CLI, VS Code, and desktop.

---

## 6. Developer Pain Points

- **Terminal rendering glitches**: The static‑mode overflow bug (#5800) is a daily frustration for users reading long replies. The upstream Ink issue (#973) remains unresolved.
- **UI pollution after edits**: The edit‑tool summary being appended to every response (#5894) makes conversations nearly unreadable – a high‑severity UX regression.
- **CI reliability**: Cross‑PR state contamination on shared runners (#5882) and a nightly release failure (#5877) erode trust in automated workflows. The need for isolated CI runners is urgent.
- **Schema warnings**: Spammy `unknown format "uint64"` messages on startup (#5897) clutter the interface, even if functionally harmless.
- **LSP integration gaps**: TypeScript‑language‑server returning empty diagnostics (#3029) blocks advanced IDE features. This issue has been open since April with no fix.
- **Memory bloat in long sessions**: Although #2036 is closed, it flagged 4–8 GB memory usage during long tasks; the community continues to watch for memory improvements (addressed in #5868 with auto‑compact thresholds).

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI / CodeWhale Community Digest — 2026-06-26

## Today’s Highlights
The project officially rebranded from `deepseek-tui` to **CodeWhale** with the v0.8.65 release; the legacy npm package is now deprecated. A burst of 20+ PRs landed today, including a critical fix for IME composition ghosting (CJK users), a new approval-card save preview for persistent permissions, and the first layer of the v0.9 command-boundary refactor. Meanwhile, a persistent plan/agent mode confusion bug (#3568) and the long‑running thinking‑collapse investigation (#861) continue to draw community attention.

## Releases
- **v0.8.65 (CodeWhale)** — Canonical project, command, npm package, and release‑asset name. Legacy `deepseek-tui` npm package is deprecated; users on `deepseek`/`deepseek-tui` v0.8.x should migrate per `docs/REBRAND.md`.  
  [GitHub Release](https://github.com/Hmbown/CodeWhale/releases/tag/v0.8.65)

## Hot Issues (10 selected)

1. **#3063 [CLOSED] v0.8.59 release tracker — TUI mouse‑report leak, runtime safety**  
   Release‑blocking tracker that shipped the macOS mouse‑report input leak fix. Closed after 11 comments across 15 days.  
   [Issue #3063](https://github.com/Hmbown/CodeWhale/issues/3063)

2. **#1186 [OPEN] feat(execpolicy): add typed persistent permission rules**  
   Would add `allow`/`deny`/`ask` rules scoped by tool name, command prefix, and path pattern. A top community feature request with 10 comments.  
   [Issue #1186](https://github.com/Hmbown/CodeWhale/issues/1186)

3. **#861 [CLOSED] bug: thinking collapse – multiple root causes**  
   The #1 verified reasoning bug: thinking blocks freeze, truncate, or drop `reasoning_content`, causing HTTP 400 on subsequent turns. Closed with 8 comments after a thorough audit.  
   [Issue #861](https://github.com/Hmbown/CodeWhale/issues/861)

4. **#2870 [OPEN] EPIC: staged command‑boundary refactor**  
   Tracks the multi‑layer refactor that migrates all subcommands into group‑owned files. Needed for v0.9. 7 comments.  
   [Issue #2870](https://github.com/Hmbown/CodeWhale/issues/2870)

5. **#3568 [OPEN] plan and agent mode mixed up YET AGAIN**  
   New report with a chat export showing the model executing file‑edit tools while in plan mode. 5 comments, serious UX regression.  
   [Issue #3568](https://github.com/Hmbown/CodeWhale/issues/3568)

6. **#3407 [OPEN] v0.8.67 setup: tools, MCP, skills, and plugins step**  
   Proposes a unified setup wizard to discover and enable power features. Part of a planned release cycle. 2 comments.  
   [Issue #3407](https://github.com/Hmbown/CodeWhale/issues/3407)

7. **#2953 [OPEN] Slim the default prompt path toward Codex‑parity input tokens**  
   Systematic effort to reduce CodeWhale’s prompt footprint (currently ~100k+ tokens larger than Codex in some tasks). 3 comments.  
   [Issue #2953](https://github.com/Hmbown/CodeWhale/issues/2953)

8. **#3537 [CLOSED] Replace hard‑coded localization with a dedicated i18n library**  
   The 5,000+ line `localization.rs` is a maintainability bottleneck. Closed with a plan to adopt an i18n library. 3 comments.  
   [Issue #3537](https://github.com/Hmbown/CodeWhale/issues/3537)

9. **#3638 [OPEN] Exposed main prompt for broader use cases**  
   User requests making the system prompt overridable via config (instead of hard‑loaded) for non‑software‑engineering uses (e.g. creative writing). 1 comment, new today.  
   [Issue #3638](https://github.com/Hmbown/CodeWhale/issues/3638)

10. **#2612 [CLOSED] Composer: hide placeholder during IME composition (CJK input)**  
    The composer ghost placeholder overlays terminal IME preedit text, making pinyin/hiragana/hangul input unreadable. Fixed in today’s PR landing.  
    [Issue #2612](https://github.com/Hmbown/CodeWhale/issues/2612)

## Key PR Progress (10 selected)

1. **#3651 [CLOSED] feat(tui): show ask‑rule save details in approval cards**  
    Adds a read‑only preview of the persistent ask rules that the `S` shortcut saves. Improves transparency for permission decisions.  
    [PR #3651](https://github.com/Hmbown/CodeWhale/pull/3651)

2. **#3652 [OPEN] Layer 4.1: Project, memory, skills, utility extraction (FEAT‑007)**  
    Continues the command‑boundary refactor, migrating remaining subcommands into group‑owned files. Follow‑up to PR #3330.  
    [PR #3652](https://github.com/Hmbown/CodeWhale/pull/3652)

3. **#3650 [OPEN] Permission control: deny, allow, and ask actions in permissions.toml**  
    Implements the core feature from #1186 – adds an `action` field to permission rules, similar to Claude Code’s model.  
    [PR #3650](https://github.com/Hmbown/CodeWhale/pull/3650)

4. **#3648 [CLOSED] fix(bridge): keep Telegram chunks readable**  
    Prevents word‑split and broken code blocks when bridging long replies to Telegram by preferring whitespace split points.  
    [PR #3648](https://github.com/Hmbown/CodeWhale/pull/3648)

5. **#3649 [OPEN] ci: add DCO (Developer Certificate of Origin) check**  
    Adds a DCO workflow to enforce `Signed-off-by` lines on commits, improving contribution compliance.  
    [PR #3649](https://github.com/Hmbown/CodeWhale/pull/3649)

6. **#3641 [CLOSED] fix(tui): bound agent card text width**  
    Fixes CJK/ASCII text overflow in delegate/fanout agent cards by reusing the display‑width truncation helper.  
    [PR #3641](https://github.com/Hmbown/CodeWhale/pull/3641)

7. **#3639 [CLOSED] fix(tui): expand failed tool output live**  
    Shows the full input summary and output of failed tools in the live transcript instead of hiding behind the omission marker.  
    [PR #3639](https://github.com/Hmbown/CodeWhale/pull/3639)

8. **#3647 [CLOSED] fix(tui): emit approval.decided after auto‑approve for external clients**  
    Ensures SSE events are fired for auto‑approved tools so that VS Code extension (and other clients) can update their UI.  
    [PR #3647](https://github.com/Hmbown/CodeWhale/pull/3647)

9. **#3634 / #3635 [CLOSED] fix(tui): keep empty composer hint off IME cursor row**  
    Solves #2612 by rendering the ghost placeholder on the row below the cursor, giving IME preedit text a clean cursor row. Two PRs merged (harvested from community contributor).  
    [PR #3634](https://github.com/Hmbown/CodeWhale/pull/3634) · [PR #3635](https://github.com/Hmbown/CodeWhale/pull/3635)

10. **#3646 [CLOSED] feat: surface weekly community digest in sitemap and navigation**  
    Makes the bilingual weekly digest accessible via `/digest`, sitemap, and navigation — now published as part of the website.  
    [PR #3646](https://github.com/Hmbown/CodeWhale/pull/3646)

## Feature Request Trends

- **Persistent permission rules** (#1186, #3650, #3651) – users want `deny`/`allow`/`ask` with tool, command, and path scoping, saved to `permissions.toml`. This is the most actively developed new feature right now.
- **Prompt customization** (#3638, #2953) – requests to make the system prompt overridable via config for non‑engineering use cases, and to slim the default prompt to Codex‑parity token counts.
- **Unified setup wizard** (#3407) – a single interactive step to discover and enable MCP servers, skills, tools, and plugins.
- **I18n library adoption** (#3537) – replace the 5k‑line `localization.rs` with a proper i18n framework to ease translation.
- **Community digest automation** (#3420, #3646) – the digest is now generated by an agent; users want it surfaced and published regularly.
- **Bridge chat improvements** (#3648, #3637) – better Telegram/WeCom message splitting and natural‑language approval responses.

## Developer Pain Points

- **Plan/agent mode confusion** (#3568) – the model frequently executes file modifications while in plan mode, breaking the intended workflow. This is a recurring, hard‑to‑fix issue.
- **Thinking collapse** (#861) – reasoning blocks freezing or silently dropping `reasoning_content` is the most severe reliability bug, though it was closed today after the audit fix.
- **IME composition ghosting** (#2612) – CJK users could not see their pinyin/hiragana input because the placeholder overlay prevented clear rendering. Now fixed.
- **Release process fragility** (#2642, #2643) – manual steps for Cargo/npm publishing and asset matching led to tag mismatches and recovery work; the community has asked for a hardened checklist.
- **REPL / approval visibility** (#3380) – the approval modal’s key hints are low‑contrast and easily missed, especially in destructive mode. Ongoing UX improvement.
- **Benchmark output discipline** (#2957, #2954) – CodeWhale generates significantly more output tokens than Codex in benchmark runs, making fair comparisons difficult; multiple issues track reducing both input and output token waste.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*