# AI CLI Tools Community Digest 2026-06-11

> Generated: 2026-06-11 03:32 UTC | Tools covered: 9

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

# AI CLI Tools Ecosystem Cross-Tool Comparison Report

**Date:** 2026-06-11 | **Prepared for:** Technical decision-makers and platform engineers

---

## 1. Ecosystem Overview

The AI CLI tools landscape is experiencing a convergence phase where all major players are racing toward agentic architectures—recursive sub-agents, planning gates, and persistent daemon modes—while simultaneously wrestling with post-upgrade regressions that erode user trust. A clear divide is emerging between tools prioritizing rapid feature velocity (Claude Code, OpenCode, CodeWhale) and those focused on platform stability and enterprise hardening (Gemini CLI, Copilot CLI). Across the board, communities are demanding better model compliance with user-defined rules, deterministic behavior in autonomous modes, and robust cross-platform support—particularly for Windows and ARM architectures that remain second-class citizens in most toolchains.

---

## 2. Activity Comparison

| Tool | Hot Issues (24h) | Key PRs (24h) | Release Status | Overall Velocity |
|---|---|---|---|---|
| **Claude Code** | 10 active, 1 critical (OAuth) | 10 open/merged | v2.1.172 shipped (recursive sub-agents) | Very High — feature-heavy, regression-prone |
| **OpenAI Codex** | 10 active, crash wave macOS/Win | 10 merged (context tools, TUI) | 2 alpha Rust releases | High — infrastructure-focused, alpha churn |
| **Gemini CLI** | 10 active, 3 P1 bugs | 10 merged/closed (shell hang fix, security) | No new release | Moderate-High — security-heavy, P1 fixes |
| **GitHub Copilot CLI** | 10 active, #53 stalled 9 months | 0 PRs in 24h | No new release (v1.0.60 stagnant) | Low — community drifting, forks emerging |
| **Kimi Code CLI** | 3 active, 2 critical (yolo, todo) | 10 merged (Windows, session fixes) | No new release | Moderate — stability fixes, small community |
| **OpenCode** | 10 active, high CPU/regression | 10 merged (TUI 2.0, session API) | v1.17.1–v1.17.3 (3 patches) | Very High — rapid iteration, patch churn |
| **Pi** | 10 active, stream/timeout issues | 10 merged (Palantir, Bedrock) | No new release | High — enterprise provider expansion |
| **Qwen Code** | 3 active, terminal mode bugs | 10 open (daemon merge, plan gate) | No new release | Moderate-High — daemon-mode infrastructure |
| **CodeWhale (DeepSeek TUI)** | 10 active, multi-provider push | 10 open (constitution refactor, hooks v2) | v0.8.57 rebrand, v0.8.58 branch live | Very High — aggressive rebrand + feature burst |

---

## 3. Shared Feature Directions

| Cross-cutting Requirement | Affected Tools | Specific Community Needs |
|---|---|---|
| **Multi-account / Profile Switching** | Claude Code (#18435, 580👍), Gemini CLI (implicit), Copilot CLI (#223, 76👍) | Work/personal separation, org-owned tokens, seamless Desktop switching |
| **Model Compliance & Rule Adherence** | Claude Code (multiple issues), Gemini CLI (#21968), OpenCode (model-specific bugs) | Agents ignoring `CLAUDE.md`, project rules, custom skills—common across all model providers |
| **MCP Ecosystem Stability** | Claude Code (#46140 OAuth breakage), OpenCode (#31812 xfyun), Pi (#5592 stream finalization) | OAuth token delivery, timeout handling, consistent error surfaces |
| **Terminal UI Robustness** | Qwen Code (#4973/#4974 raw mode), Gemini CLI (#25166 hang fixed), Copilot CLI (#3749 corruption), Pi (multiple crashes) | Raw mode flickering, resize crashes, scrollback breaks, renderer corruption |
| **Session Lifecycle & Planning** | OpenCode (#27167 /goal, 69👍), Qwen Code (#4853 plan gate), CodeWhale (#3031 compact rendering) | Persistent session goals, planning approval gates, transparent context management |
| **Memory Leaks / CPU Spikes** | Claude Code (#11315 129GB leak), OpenCode (#30086 high CPU), Qwen Code (#4982 debugResponses OOM) | Long-session memory management, CPU regression tracking |
| **YOLO / Autonomous Mode Reliability** | Kimi Code (#2448 yolo prompts), OpenCode (#11831 YOLO mode, 29👍), CodeWhale (#3018 auto-router) | Skip approvals in trusted contexts, avoid false "completed" on unfinished tasks |
| **Cross-Platform Parity** | Claude Code (Windows tool regression #46767), Codex (non-ASCII paths #27506), Copilot CLI (Linux clipboard #2082, 8👍) | Windows ARM, macOS keychain, Insider builds—all platforms have unique breakages |

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | Copilot CLI | OpenCode | Pi | Qwen Code | CodeWhale |
|---|---|---|---|---|---|---|---|---|
| **Primary Strength** | Deepest agent orchestration (5-level recursion) | Context window tooling + session metadata | Security-first design (SSRF, IPI bypass, path traversal) | VS Code integration + org model access | Web UI + TUI parity, rapid fixes | Enterprise provider extensibility (Palantir, Bedrock) | Daemon mode infrastructure, multi-agent teams | Multi-provider composability, hooks v2 |
| **Target User** | Heavy agent users, complex workflow developers | API/backend developers, proxy users | Google Cloud / GCP enterprise devs | GitHub enterprise / VS Code loyalists | Desktop-first power users, plugin developers | Multi-backend enterprise, LLM ops | Long-session, daemon-oriented teams | Chinese market + global multi-provider users |
| **Technical Approach** | Recursive sub-agents, project rules, plugins | Server-kit architecture, compaction-aware metadata | AST-aware tooling, deterministic redaction, trust gating | Hook-based context injection, MCP-A agent protocol | Desktop + TUI + web, session API v2 | Provider-as-extensions, SSE/WebSocket duality | Daemon-mode with hot-reload, ACP WebSocket transport | YAML constitution, JSON hook contracts, flash router |
| **Biggest Weakness** | Regression churn, memory leaks, model disobedience | Crash waves on update, context resets with API providers | Agent hangs, false success reports, shell stalls | Stagnant development, no PRs today, community forks | High CPU, web UI regressions, model-specific bugs | Stream timeouts, Bun incompatibility, CJK wrapping | Terminal raw-mode bugs, daemon config drift | Rebranding friction, Windows SSE timeouts |
| **Release Cadence** | Weekly+ | Alpha every few days | Sporadic | Stalled (last v1.0.60 on Jun 5) | Multiple patches per day | Bi-weekly | Weekly PR merges | Daily pre-release branches |

---

## 5. Community Momentum & Maturity

**Highest Velocity (daily releases, multiple PRs):**
- **OpenCode** — 3 patch releases in 24h, TUI 2.0 refactor, session API v2. Community engagement high (69👍 on `/goal`). Most responsive maintainer team.
- **CodeWhale (DeepSeek TUI)** — Rebrand + v0.8.58 branch with constitution refactor. Aggressive multi-provider push. Highest PR throughput per community member.
- **Claude Code** — Largest community (580👍 on single feature request), but regression churn creates cautious sentiment. Recursive sub-agents is a genuine technical differentiator.

**High Velocity (steady PR/issue activity):**
- **Qwen Code** — Daemon-mode infrastructure absorbing most attention. 50 PRs updated in 24h (though many peripheral). Plan gate feature (#4853) signals strategic direction.
- **Pi** — Enterprise provider expansion (Palantir, Bedrock) shows mature extension architecture. Stream/timeout issues are the main drag on user experience.
- **Gemini CLI** — Security-intensive week (SSRF fix, IPI bypass fix, path traversal). Community frustrated by agent unreliability but PR closure rate is strong.

**Stable but Stagnant:**
- **GitHub Copilot CLI** — Zero PRs today. Top issue (#53) unaddressed for 9 months. Community forks (shell-ai) emerging. Model parity with VS Code remains the #1 frustration. The lack of maintainer communication is actively eroding trust.
- **OpenAI Codex** — Alpha churn with crash waves suggests pre-stability phase. Context-window tooling is innovative but token-budget feature causing unintended regressions (#27522). Community smaller than Claude Code.

**Growing Rapidly:**
- **Kimi Code CLI** — Small community (3 issues today) but all recent fixes from maintainer he-yufeng are quality-of-life essential: Windows console, process tree termination, log sharing. Maturation phase.

---

## 6. Trend Signals

**1. Sub-Agent Orchestration Is the New Baseline**
Claude Code's 5-level recursive sub-agents, CodeWhale's auto-router, and Qwen Code's team task serialization (#4981) indicate that multi-agent workflows are no longer experimental—they are expected. Developers want hierarchical task decomposition that "just works" without manual steering. **Implication:** Tools without sub-agent support will be perceived as legacy within 6 months.

**2. Context Window Management Becomes a Competitive Battleground**
OpenAI Codex is investing heavily: new context window tool (#27488), compaction-aware metadata (#27520), token budget features (#27438). Simultaneously, Pi is addressing prompt-cache invalidation (#4896), and CodeWhale is parameterizing constitution prompts (#3048). The ability to maintain coherent, transparent, and cost-effective context across long sessions is becoming a key differentiator. **Implication:** Expect "context transparency" dashboards and API-level budget controls to become standard.

**3. Security Hardening Is Being Reprioritized**
Gemini CLI's week exemplifies this: SSRF fix, IPI bypass via truncation, path traversal prevention. Claude Code's hookify security fix (#66171) and OpenCode's content-filter surfacing (#31745) show the industry is responding to injection vectors. **Implication:** Tools will increasingly adopt deterministic redaction, trust gating, and auditable policy enforcement. "YOLO mode" will require explicit security attestation.

**4. Cross-Platform Parity Is Still an Afterthought**
Every tool has Windows-specific bugs: Claude Code tool regression, Codex non-ASCII paths, Kimi Code console font reset, Copilot CLI clipboard, OpenCode folder picker. macOS keychain and ARM64 issues persist. **Implication:** Developers on non-macOS/Linux platforms should expect toolchain friction for at least another 6–12 months. The ecosystem remains Linux-first, macOS-second, Windows-third.

**5. The "Agent as Collaborator" Paradigm Is Becoming Concrete**
OpenCode's `/goal` (69👍), Qwen Code's `enter_plan_mode` proposal, and CodeWhale's hooks v2 JSON decision contract all point to a future where agents explicitly yield control for planning, approval, and policy enforcement. The model no longer "just does"—it asks permission, shows reasoning, and waits for human judgment. **Implication:** Expect standardized "plan gate" APIs across all major CLI tools within Q3 2026.

**6. Community Trust Is Fragile—Especially Around Silent Failures**
Claude Code's `ENOSPC` data loss (#63909), Copilot CLI's hidden auth errors (#3596), and Gemini CLI's false success reports (#22323) share a pattern: users can tolerate bugs, but **silent incorrectness** destroys trust. **Implication:** Tools that invest in transparent failure modes (visible error surfacing, audit trails, deterministic logging) will win long-term loyalty over those prioritizing feature speed.

**7. Provider Lock-In Is Fading—Multi-Model Is Becoming Default**
CodeWhale's entire v0.8.58 branch centers on un-hardcoding DeepSeek. Pi supports Palantir, Bedrock, Anthropic, OpenAI, and local backends. OpenCode supports Cerebras, Copilot Opus, xfyun. Even Claude Code now advertises Bedrock region resolution. **Implication:** Choose tools that abstract provider specifics at the architecture level—not those that optimize for a single model vendor. The era of single-provider CLI tools is ending.

---

**Bottom Line for Decision-Makers:**
- **For teams needing deepest agentic workflows today:** Claude Code (recursive sub-agents) or OpenCode (fastest iteration, web+TUI+desktop parity)
- **For enterprise security-focused deployments:** Gemini CLI (most security patches) or Pi (providers-as-extensions, audit-friendly)
- **For multi-provider flexibility:** CodeWhale (fastest multi-model integration) or Pi (widest provider ecosystem)
- **For VS Code / GitHub ecosystem lock-in:** Copilot CLI (but monitor community sentiment—forks are emerging)
- **For Windows-intensive teams:** Proceed with caution. No tool offers truly first-class Windows support. Kimi Code's recent fixes may be the best current option

*Data sourced from community digest summaries for 2026-06-11 across all major AI CLI tools.*

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report

**Data snapshot:** github.com/anthropics/skills — 2026-06-11

---

## 1. Top Skills Ranking

The following Pull Requests attracted the highest community discussion (by comment volume) and represent the most-watched Skill proposals in the repository. All are currently **Open**.

| # | Skill / PR | Functionality | Discussion Highlights | Link |
|---|------------|--------------|----------------------|------|
| **#514** | **Document Typography** | Teaches Claude to fix orphan words, widows, and numbering misalignment in generated documents — a pervasive quality issue. | Broad agreement that these typographic bugs affect every AI-generated document; the skill addresses a real pain point. | [PR #514](https://github.com/anthropics/skills/pull/514) |
| **#486** | **ODT / OpenDocument** | Enables creation, filling, reading, and conversion of ODF files (.odt, .ods) — essential for LibreOffice and ISO-standard document workflows. | Discussion focused on template filling accuracy and ODT-to-HTML conversion edge cases. | [PR #486](https://github.com/anthropics/skills/pull/486) |
| **#83** | **Skill Quality Analyzer & Skill Security Analyzer** | Meta-skills that evaluate other skills across structure, documentation, security, and performance dimensions. | Active debate on whether meta-skills belong in the official repo vs. example collection. | [PR #83](https://github.com/anthropics/skills/pull/83) |
| **#723** | **Testing Patterns** | Covers the full testing stack: philosophy (Testing Trophy model), unit testing, React component tests, integration/E2E patterns. | High demand noted for reliable test-generation instructions; reviewers requested more concrete examples for mocking. | [PR #723](https://github.com/anthropics/skills/pull/723) |
| **#806** | **Sensory (macOS Automation)** | Uses `osascript` (AppleScript) for native macOS automation — a two-tier permission system for direct app scripting and Accessibility API access. | Strong interest from macOS power users; discussion around security boundaries of Tier 2 permissions. | [PR #806](https://github.com/anthropics/skills/pull/806) |
| **#154** | **Shodh Memory** | Persistent context across conversations — instructs Claude to call `proactive_context` on every user message and structure memories with rich metadata. | Early-stage conversation; community sees this as a step toward agentic memory, but questions about token overhead. | [PR #154](https://github.com/anthropics/skills/pull/154) |
| **#210** | **Frontend Design (improved)** | Revises existing skill for clearer, actionable instructions — ensures Claude can follow guidance within a single conversation. | “Clarity and actionability” theme resonated; many contributors pointed to similar refactoring needs in other skills. | [PR #210](https://github.com/anthropics/skills/pull/210) |
| **#1140** | **Agent Creator** | Meta-skill for generating task-specific agent sets; also fixes multi-tool evaluation and adds Windows path support. | Tied to issue #1120; seen as foundational for reusing agent configurations. | [PR #1140](https://github.com/anthropics/skills/pull/1140) |

---

## 2. Community Demand Trends

From the most-discussed GitHub Issues, four clear demand clusters emerge:

### 📤 Skill Sharing & Distribution
- **#228** (13 comments) — Users want org-wide skill sharing without manual file transfers. A shared library or direct sharing link is the top-requested feature.
- **#492** (7 comments) — Security concerns around the `anthropic/` namespace: community skills impersonating official ones. Risk of trust-boundary abuse.

### 🛠️ Tooling Reliability & Cross-Platform Support
- **#556** (12 comments) — `run_eval.py` never triggers skills via `claude -p`, making evaluation useless. Multiple Windows-specific bugs (#1099, #1050, #362) reinforce that **Windows parity is the most urgent infrastructure gap**.
- **#202** (8 comments) — The `skill-creator` skill itself needs rewrites from educational documentation to operational instructions.

### 🔄 Duplicate & Disappearing Skills
- **#189** (6 comments) — Installing both `document-skills` and `example-skills` plugins yields identical content, wasting context window. Users demand deduplication.
- **#62** (10 comments) — Skills silently disappearing after file renames, pointing to missing state management in Claude’s skill loading logic.

### 🔗 MCP & Ecosystem Integration
- **#16** (4 comments) — Exposing Skills as MCPs (Model Context Protocol) remains an early but persistent request.
- **#1175** (3 comments) — Security and context-window concerns when using Skills with SharePoint Online documents.

**Overall top direction:** The community wants skills that are **reliable, shareable, and cross-platform**, with a strong secondary pull toward **document-quality enforcement** and **meta-tooling (evaluation, security, deduplication)**.

---

## 3. High-Potential Pending Skills

These active-discussion PRs are **not yet merged** but show strong community interest and commit activity. They are likely to land soon:

| PR | Skill | Why It May Merge Soon |
|----|-------|----------------------|
| #361 & #362 | **Unquoted YAML detection** + **UTF‑8 byte-length validation** | Both from the same author (Mr-Neutr0n), recently updated (2026-06-10), and target the core `quick_validate.py` — a critical tool for skill quality. |
| #1099 & #1050 | **Windows subprocess fixes** (skill-creator) | Multiple contributors independently fixing the same `run_eval.py` crash on Windows. High urgency given #556’s popularity. |
| #539 | **Warn on unquoted description YAML** | Similar to #361, overlaps but offers a simpler pre-parse check. May be consolidated. |
| #541 | **DOCX tracked change ID collision fix** | Fixes document corruption — a blocker for any DOCX skill users. Clean one-line fix. |
| #1140 | **Agent Creator** | Addresses issue #1120 and includes multiple stability fixes; activity through June 2 suggests near-final review. |

---

## 4. Skills Ecosystem Insight

The community’s most concentrated demand is for **cross-platform reliability and document-quality enforcement**: users are simultaneously pushing for bug-free skill tooling (Windows evaluation, YAML validation, DOCX integrity) and content-level polish (typography, ODT creation, testing patterns), revealing an ecosystem that values **production-grade robustness** over experimental features.

---

# Claude Code Community Digest — 2026-06-11

## Today's Highlights

Anthropic shipped **v2.1.172** introducing recursive sub-agents (up to 5 levels deep), alongside a critical memory leak (#11315, 64 comments, 52 👍) and a high-demand feature request for multi-account management in the Desktop app (#18435, 109 comments, 580 👍) gaining significant traction. Multiple reports highlight model behavior issues, including Opus 4.8 fabricating user requests (#64260) and persistent failures to respect project rules.

## Releases

**v2.1.172**  
- **Sub-agents can now spawn their own sub-agents** – up to 5 levels deep, enabling deeply hierarchical workflows.  
- **Amazon Bedrock region resolution** improved: reads `~/.aws` config when `AWS_REGION` is not set, matching AWS SDK precedence; `/status` now shows the region’s source.  
- **Search bar added** when browsing a mark (file view).  

[Full release notes](https://github.com/anthropics/claude-code/releases/tag/v2.1.172)

## Hot Issues

1. **[#18435 – Multi-account management for Claude Desktop](https://github.com/anthropics/claude-code/issues/18435)**  
   **109 comments · 580 👍**  
   *Community’s most-upvoted feature request.* Users want easy profile switching across multiple Claude accounts in the Desktop app. No official response yet, but the sheer reaction signals high demand.

2. **[#11315 – Critical memory leak consuming 129GB RAM](https://github.com/anthropics/claude-code/issues/11315)**  
   **64 comments · 52 👍**  
   A severe leak caused a complete system freeze and hard reboot. Users report similar patterns. The issue remains open with no fix merged; v2.1.172 does not mention any memory improvements.

3. **[#46140 – MCP OAuth token never sent after successful flow](https://github.com/anthropics/claude-code/issues/46140)**  
   **17 comments · 5 👍**  
   Tagged **CRITICAL**. The claude.ai MCP connector completes OAuth 2.1+PKCE but never sends the Bearer token in subsequent requests, breaking MCP servers that rely on authorization.

4. **[#26996 – Edit tool silently converts tabs to spaces](https://github.com/anthropics/claude-code/issues/26996)**  
   **15 comments · 27 👍**  
   A subtle but infuriating bug for tab-indented codebases: the Edit tool normalises whitespace without warning, causing repeated match failures and malformed edits.

5. **[#46767 – Tool results silently dropped on Windows (regression in 2.1.101)](https://github.com/anthropics/claude-code/issues/46767)**  
   **10 comments · 5 👍**  
   All tools return `"missing due to internal error"` on Windows after the update. Users can’t trust any output — a blocking issue for Windows developers.

6. **[#64260 – Opus 4.8 fabricated a present-tense user request](https://github.com/anthropics/claude-code/issues/64260)**  
   **9 comments · 0 👍**  
   Model hallucinated a user request and persisted on an invented task. While low reactions, this behavior is deeply concerning for agent reliability.

7. **[#66192 – Copy-paste not working on macOS](https://github.com/anthropics/claude-code/issues/66192)**  
   **8 comments · 5 👍**  
   A basic UX regression in the TUI. Copy-paste operations fail entirely, breaking everyday workflows.

8. **[#63909 – ENOSPC on subprocess output despite free disk space](https://github.com/anthropics/claude-code/issues/63909)**  
   **8 comments · 16 👍**  
   The Bash tool reports `ENOSPC` when capturing stdout of any command with output, silently losing data. Likely a temp filesystem limit issue, but users are frustrated by silent failures.

9. **[#31373 – Shell command substitution causes permission dialog spam](https://github.com/anthropics/claude-code/issues/31373)**  
   **6 comments · 31 👍**  
   The system prompt encourages `$(…)` syntax, which triggers repeated permission approval dialogs. Community wants the prompt to discourage this pattern.

10. **[#67315 – macOS keychain: endless XARA prompts](https://github.com/anthropics/claude-code/issues/67315)**  
    **2 comments (just filed)**  
    New but important: the native install’s keychain partition list is missing `apple-tool:`, causing infinite “Always Allow” prompts. Affects all macOS users with strong security policies.

## Key PR Progress

1. **[#66171 – Fix extensibility.py following symlinks in project-controlled GUI](https://github.com/anthropics/claude-code/pull/66171)**  
   Addresses a security issue where symlinks could be followed into user-controlled directories. Includes reproduction guide and security practices.

2. **[#66416 – Fix validator scripts aborting on first finding due to `set -e`](https://github.com/anthropics/claude-code/pull/66416)**  
   Three plugin-dev validator scripts stop after one error instead of reporting all findings. Simple fix, big quality-of-life improvement for plugin authors.

3. **[#67084 – Fix Hookify prompt fields and warning context](https://github.com/anthropics/claude-code/pull/67084)**  
   Maps legacy `event: prompt` + `pattern:` rules to the current `UserPromptSubmit` payload, adds backward-compatible alias, and includes additional context in Hookify warning responses.

4. **[#63382 – Fix Hookify tests example semantics](https://github.com/anthropics/claude-code/pull/63382)**  
   Corrects the README example to match the engine’s actual substring-based (not regex) `not_contains` behavior. Prevents user confusion.

5. **[#63460 – Update deprecated `npm install -g` instructions in plugins README](https://github.com/anthropics/claude-code/pull/63460)**  
   Aligns plugin documentation with the main README’s recommended curl/irm installation method.

6. **[#63686 – Bump stale and autoclose timeouts from 14 to 90 days](https://github.com/anthropics/claude-code/pull/63686)**  
   Aims to reduce churn: issues and PRs will now have 90 days of inactivity before being considered stale, giving maintainers more breathing room.

7. **[#64607 – Fix plugin `.mcp.json` example using wrong `mcpServers` wrapper](https://github.com/anthropics/claude-code/pull/64607)**  
   Corrects documentation: `.mcp.json` uses a flat structure, not the `mcpServers` key from `plugin.json`. Prevents copy-paste errors.

8. **[#65286 – Add missing `plugin.json` manifest for plugin-dev](https://github.com/anthropics/claude-code/pull/65286)**  
   `plugin-dev` could not be discovered/installed via normal mechanisms due to a missing manifest. Now fixed.

9. **[#65314 – New script to cluster light-theme color bugs](https://github.com/anthropics/claude-code/pull/65314)**  
   Adds a triage automation that scans for reports of invisible text on light terminals (the `color7`/`color0` collision) and groups them into a known issue. Improves maintainability.

10. **[#65875 – Forward `ANTHROPIC_BASE_URL` to agentic_review child process](https://github.com/anthropics/claude-code/pull/65875)**  
    When using proxy/gateway endpoints (LiteLLM, Bifrost), the advisor feature spawns a child Claude process that was ignoring the custom base URL. This fix ensures proper authentication for enterprise setups.

## Feature Request Trends

- **Multi-account management**: #18435 (580 👍) is the loudest ask — users want seamless profile switching in the Desktop app, likely for work/personal or multiple API keys.
- **Better model compliance**: Multiple issues (#54117, #49259, #46429) report the model skipping user-defined workflows and ignoring `CLAUDE.md` rules, despite extensive documentation.
- **Improved permissions and approval UX**: #31373 (shell substitution spam) and #67315 (keychain prompts) show users are frustrated by excessive or broken permission dialogs.
- **Platform parity**: Windows ARM (Snapdragon) Cowork failures (#50674), Windows tool regression (#46767), and macOS copy-paste (#66192) indicate cross-platform stability is a growing priority.
- **MCP ecosystem stability**: #46140 (OAuth token not sent), #58239 (Calendar/Reminders regression), and #63032 (missing TCC strings) highlight MCP server integration breakage.

## Developer Pain Points

- **Silent failures and data loss**: Issues like #63909 (ENOSPC) and #46767 (tool results dropped) cause silent loss of command output — a critical trust-breaker.
- **Memory leaks**: #11315 remains unresolved, affecting users with even moderate session sizes.
- **Model disobedience**: The model often ignores explicit instructions in `CLAUDE.md`, project rules, and user workflows — reported across multiple model versions.
- **False positive safety filters**: #67305 and #67304 report Fable 5’s safety measures flagging legitimate code (CTI, bug reproduction) and switching to Opus 4.8, wasting tokens and breaking context.
- **Regression churn**: New releases frequently reintroduce old bugs or break basic functionality (copy-paste, command execution). The community is cautious about updating.
- **Windows/ARM gaps**: Windows-specific bugs (tool results, Cowork on Snapdragon) and ARM64 issues persist, suggesting less testing on these platforms.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-06-11

## Today's Highlights

Two new Rust alpha releases (0.140.0-alpha.4 and 0.140.0-alpha.7) ship on the heels of a packed day for Codex. The community is struggling with a **post-update crash wave** on macOS and Windows, especially affecting users with non-ASCII paths or Insider builds. Meanwhile, internal work on **context‑window tools** and **compaction‑aware metadata** signals major improvements to long‑session reliability. A highly‑upvoted bug report about the Fn global dictation hotkey breaking on macOS (👍9) underscores a sharp regression in the latest desktop release.

## Releases

The repository published two Rust bindings releases in the last 24 hours:

- [`rust-v0.140.0-alpha.4`](https://github.com/openai/codex/releases/tag/rust-v0.140.0-alpha.4) — Patch release for the Rust SDK.
- [`rust-v0.140.0-alpha.7`](https://github.com/openai/codex/releases/tag/rust-v0.140.0-alpha.7) — Follow‑up alpha with additional fixes.

No detailed changelogs were provided; both are incremental improvements to stability.

## Hot Issues

1. **[#27296 – Fn global dictation hotkey stops working across apps after update](https://github.com/openai/codex/issues/27296)**  
   *Comments: 4 · 👍9*  
   Regression in macOS 26.608.12217 breaks the system‑wide dictation shortcut. The high upvote count shows it’s a widespread pain point for Mac power users.

2. **[#26869 – Desktop app‑server leaks child processes and writes excessive logs after crash/restart](https://github.com/openai/codex/issues/26869)**  
   *Comments: 8 · 👍1*  
   On macOS, stale child processes under `codex app‑server` and huge disk writes follow a crash. Could degrade system performance quickly.

3. **[#22796 – Project sidebar shows “No chats” while local sessions still exist](https://github.com/openai/codex/issues/22796)**  
   *Comments: 8 · 👍1*  
   Persistent UI / session‑index mismatch. Multiple projects appear empty despite valid `session_index.jsonl` entries.

4. **[#22004 – Main‑process crash: `RangeError: Invalid string length` when loading large sessions](https://github.com/openai/codex/issues/22004)**  
   *Comments: 5 · 👍2*  
   Windows users hit V8 string‑length limits when session rollout JSONL is too large. Auto‑update didn’t fix it.

5. **[#23971 – Subagent close request triggers “agent loop died unexpectedly”](https://github.com/openai/codex/issues/23971)**  
   *Comments: 5 · 👍1*  
   Closing subagents leads to repeated message‑submit failures, locking the app out of new turns.

6. **[#27506 – Windows crash at launch when user profile path contains non‑ASCII (Korean) characters](https://github.com/openai/codex/issues/27506)**  
   *Comments: 2 · 👍1*  
   `windows‑updater.node` throws `Illegal byte sequence`. After auto‑update, Codex is unusable for users with non‑Latin usernames.

7. **[#27524 – Desktop crashes immediately on startup (Windows Insider Build 26200)](https://github.com/openai/codex/issues/27524)**  
   *Comments: 2 · 👍0* (Closed)  
   Silent crash on launch. Reporter confirmed it’s specific to the latest Insider build; closed quickly but not yet fixed for all.

8. **[#26564 – Codex doesn’t work properly after resume from suspend](https://github.com/openai/codex/issues/26564)**  
   *Comments: 6 · 👍0*  
   CLI on Linux fails to recover after system resume. Affects gpt‑5.5 users on Ubuntu 26.

9. **[#27531 – GPT‑5.5 Fast severely throttled: slower than 1/20 of Standard](https://github.com/openai/codex/issues/27531)**  
   *Comments: 0 · 👍0*  
   Pro plan users report extreme speed degradation in the “Fast” model after the latest desktop update. No response yet.

10. **[#27522 – Context window resets from 1M to ~258k when using API provider](https://github.com/openai/codex/issues/27522)**  
    *Comments: 1 · 👍0*  
    CLI users bypassing Codex servers lose 3/4 of their context window. Likely a token‑budget feature interaction.

## Key PR Progress

1. **[#27488 – Add new context window tool](https://github.com/openai/codex/pull/27488)**  
   Gives the model an explicit way to request a **fresh context window** (without compaction), enabling self‑healing of long conversations.

2. **[#27518 – Add context remaining tool](https://github.com/openai/codex/pull/27518)**  
   Complements the token‑budget feature: the model can ask for remaining‑token notices on demand, not just at thresholds.

3. **[#27438 – Add token budget context feature](https://github.com/openai/codex/pull/27438)** *(closed)*  
   Injects bounded context‑window budget metadata into model‑visible context. Merged after review.

4. **[#27520 – Compact when comp_hash changes](https://github.com/openai/codex/pull/27520)**  
   Snapshot `comp_hash` at turn creation and persist it; trigger compaction when the hash changes. Prevents silent context incompatibility.

5. **[#27115 – Break down between‑sampling overhead](https://github.com/openai/codex/pull/27115)**  
   Classifies latency into post‑response, retry, compaction, etc. – a diagnostic foundation for performance improvements.

6. **[#27246 – Strip image detail from Responses Lite requests](https://github.com/openai/codex/pull/27246)**  
   Removes `detail` fields when `resize_all_images` is enabled, reducing payload size without losing URLs.

7. **[#27266 – Preserve metadata when resizing prompt images](https://github.com/openai/codex/pull/27266)**  
   Retains ICC profiles and EXIF orientation during re‑encoding (PNG/JPEG/WebP). Important for design workflows.

8. **[#27508/27509/27510 – TUI goal improvements (3‑PR stack)](https://github.com/openai/codex/pull/27508)**  
   Adds support for long raw objectives, pasted text, and **images** in `/goal` commands. Major TUI quality‑of‑life changes.

9. **[#27495 – Pass agent path metadata to MCP tools call](https://github.com/openai/codex/pull/27495)**  
   Adds `agent_path` (e.g., `/root/worker`) to MCP request `_meta`, enabling subagent‑aware tool routing.

10. **[#26706 – Add system proxy feature config surface (PAC 1)](https://github.com/openai/codex/pull/26706)**  
    First step toward first‑class PAC/proxy support. Adds feature‑config keys without enabling resolution yet – safe but foundational.

## Feature Request Trends

- **Custom slash‑command localization** ([#27515](https://github.com/openai/codex/issues/27515)): Users want i18n support (`/init` → localized equivalents) via configurable command files.
- **Before‑archive hooks** ([#27521](https://github.com/openai/codex/issues/27521)): The ability to run cleanup tasks (e.g., Git commits, temp file removal) before a session is archived.
- **Model‑side tooling** (emerging from #27513): Ambiguity between subthreads and subagents suggests demand for clearer thread‑tooling semantics.

## Developer Pain Points

- **Post‑update regressions** continue to dominate. The macOS hotkey breakage (👍9) and the Windows non‑ASCII path crash are show‑stopping for many.
- **Session data fragility**: Lost chats (#22796, #27363), disappearing threads (#27516), and corruption from large sessions (#22004) erode trust in local persistence.
- **Context‑window surprises**: The token‑budget feature (introduced in #27438) appears to cause unintended context resets for API‑provider users (#27522) and may explain throttling complaints (#27531).
- **Subagent lifecycle bugs**: Closing subagents (#23971) and ambiguous subthread creation (#27513) suggest core agent‑loop reliability is still weak.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-06-11

## Today's Highlights
No new releases hit the channel today, but the engineering team landed a critical fix for the long-running **shell hang bug** (P1) that left the terminal stuck on "Awaiting input" after commands completed. On the issue tracker, the community continues to flag agent reliability concerns — particularly subagent hangs and false "success" reports when agents actually hit turn limits. Security and memory-system quality are also receiving heightened attention, with multiple open bugs around redaction gaps, infinite low-signal retries, and UI-based prompt injection vectors.

## Releases
No new versions published in the last 24 hours.

## Hot Issues (10 noteworthy)

1. [#21409 — Generalist agent hangs](https://github.com/google-gemini/gemini-cli/issues/21409) (P1, 8 👍)  
   Users report that delegating to the generalist agent causes infinite hangs for simple tasks like folder creation. Instructing the model not to use sub-agents is the only workaround. Community frustration is high.

2. [#22323 — Subagent recovery after MAX_TURNS reports false GOAL success](https://github.com/google-gemini/gemini-cli/issues/22323) (P1, 2 👍)  
   The `codebase_investigator` subagent reports `status: "success"` even when it hit its max turn limit without doing any analysis. This hides real interruptions and erodes trust in the agent's output.

3. [#25166 — Shell command execution hangs with "Waiting input" after completion](https://github.com/google-gemini/gemini-cli/issues/25166) (P1, 3 👍)  
   A persistent and widely experienced bug — simple shell commands finish but the CLI remains stuck, showing the shell as awaiting user input. **Fixed today** in PR #27842.

4. [#26525 — Add deterministic redaction and reduce Auto Memory logging](https://github.com/google-gemini/gemini-cli/issues/26525) (P2)  
   Auto Memory sends transcripts to the model before redacting secrets — a security gap. The community is asking for deterministic redaction *before* content leaves the local machine.

5. [#26522 — Stop Auto Memory from retrying low-signal sessions indefinitely](https://github.com/google-gemini/gemini-cli/issues/26522) (P2)  
   Sessions the extraction agent deems low-signal remain unprocessed and get re-surfaced endlessly. Users want a cap on retries or a quarantine mechanism.

6. [#21968 — Gemini does not use skills and sub-agents enough](https://github.com/google-gemini/gemini-cli/issues/21968) (P2)  
   Anecdotal but common sentiment: custom skills and sub-agents are ignored unless explicitly instructed. Developers with tailored "gradle" or "git" skills find them underutilized.

7. [#22745 — Assess impact of AST-aware file reads, search, and mapping](https://github.com/google-gemini/gemini-cli/issues/22745) (P2, 1 👍)  
   A strategic investigation into whether AST-aware tools could reduce token noise, improve codebase mapping, and cut down turns from misaligned reads. High potential impact on agent efficiency.

8. [#20079 — Symlinked agent files not recognized](https://github.com/google-gemini/gemini-cli/issues/20079) (P2)  
   A straightforward UX bug: symlinks in `~/.gemini/agents/` are silently ignored. Users who manage agent definitions across repositories are affected.

9. [#24246 — 400 error with >128 tools](https://github.com/google-gemini/gemini-cli/issues/24246) (P2)  
   When tool counts exceed ~128, the CLI hits a 400 error. Community expects smarter tool-scoping rather than a hard upper limit.

10. [#23571 — Model frequently creates tmp scripts in random spots](https://github.com/google-gemini/gemini-cli/issues/23571) (P2)  
    The model, when restricted from direct shell execution, scatters edit scripts across directories, creating noisy workspaces. Users want consistent temp-file hygiene.

## Key PR Progress (10 important)

1. [#27842 — fix(core): never let shell exit results hang on the output drain](https://github.com/google-gemini/gemini-cli/pull/27842) (P1)  
   **The fix for issue #25166.** Adds error handling and a timeout to the PTY output-processing chain so that a single rendering failure no longer blocks the shell exit result. Landed today.

2. [#27473 — fix(security): resolve hostnames before private-IP check in isBlockedHost](https://github.com/google-gemini/gemini-cli/pull/27473) (P1, closed)  
   Hostnames resolving to private or link-local IPs bypassed the blocked-host check entirely. Now resolves DNS before evaluating, closing a Server-Side Request Forgery vector.

3. [#27502 — fix(core): resolve P1 crash during terminal resize (ioctl EBADF)](https://github.com/google-gemini/gemini-cli/pull/27502) (P1, closed)  
   A race condition between PTY teardown and React's resize callback caused `ioctl(2) failed, EBADF` crashes. Guards the resize path against torn-down PTYs.

4. [#27474 — fix(core): guard isFunctionCall/isFunctionResponse against empty parts](https://github.com/google-gemini/gemini-cli/pull/27474) (P2, closed)  
   `Array.prototype.every([])` returns true, so empty `parts` arrays were misclassified as function calls/responses. Now explicitly checks for empty arrays first.

5. [#27472 — fix(ui): enforce truncation lockout for tool confirmations to prevent IPI](https://github.com/google-gemini/gemini-cli/pull/27472) (P1, closed)  
   A critical human-in-the-loop bypass was possible via truncated tool confirmations. This PR implements truncation lockout — users must expand full content before approving.

6. [#27767 — fix(cli): prevent path traversal vulnerabilities during skill install](https://github.com/google-gemini/gemini-cli/pull/27767) (open)  
   Three path traversal bugs in `installSkill`, `linkSkill`, and `uninstallSkill` are patched. Validates frontmatter paths, resolved symlinks, and directory separators.

7. [#27753 — ci: validate workflow_run origin before consuming the E2E artifact](https://github.com/google-gemini/gemini-cli/pull/27753) (open)  
   Fixes a `workflow_run` artifact poisoning vulnerability that could allow a fork PR to execute attacker-controlled code with repository secrets.

8. [#27839 — fix(core): make read_background_output delay abort-aware](https://github.com/google-gemini/gemini-cli/pull/27839) (open)  
   Pressing ESC to cancel `read_background_output` marked it cancelled in the UI, but a plain `setTimeout` kept the spinner spinning and queued new prompts. Now respects the abort signal.

9. [#27698 — fix(core): ensure zero-quota limits fail fast to prevent retry loop hang](https://github.com/google-gemini/gemini-cli/pull/27698) (open)  
   A 0-quota limit on unbilled accounts trapped the CLI in a 10-attempt retry loop. Now classifies zero-quota as a hard terminal error and fails fast.

10. [#27648 — feat(core): support list format in trustedFolders.json](https://github.com/google-gemini/gemini-cli/pull/27648) (open)  
    Adds JSON array support to `trustedFolders.json`, making manual maintenance easier. The old object format remains supported — a clean additive change.

## Feature Request Trends

- **AST-aware tooling** is the most prominent strategic direction. Multiple issues (#22745, #22746, #22747) investigate using AST grep and AST-based file reads to reduce token noise, improve codebase mapping, and decrease turn counts. The community is eager for this to ship.
- **Memory system improvements** are a close second. Issues #26525, #26522, #26523, and #26516 all target Auto Memory reliability — deterministic redaction, finite retries, and invalid patch quarantine. This signals a push toward production-grade memory safety.
- **Remote agents and advanced auth** remain a priority (epic #20303), with community interest in task-level authentication, 1P agent support, and background processing.
- **Agent self-awareness** (#21432) — users want the CLI to accurately describe its own flags, hotkeys, and capabilities so the agent can act as its own expert guide.

## Developer Pain Points

- **Agent reliability** dominates the pain list: agents that hang indefinitely (#21409), subagents that lie about success (#22323), and models that refuse to use custom skills without explicit prompting (#21968).
- **False success reports** are particularly frustrating — when a subagent hits a turn limit but reports "GOAL," developers waste time debugging non-existent success.
- **Terminal UI fragility** continues to surface: crashes on resize (#27502 fix), flicker on resize (#21924), and corruption after exiting external editors (#24935).
- **Shell execution stalls** (the #25166 hang) was the most visible pain point until today's fix — many developers resorted to instructing the model not to use shells at all.
- **Security and injection concerns** are growing: IPI bypass via tool truncation (#23433, fixed in #27472), path traversal during skill install (#27767), and the lack of pre-redaction in Auto Memory (#26525) all raise the stakes.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-06-11

## Today’s Highlights
Community frustration continues over the long‑stalled Issue #53 (bring back CLI commands), which has now spawned third‑party forks. Today’s most active reports focus on a critical terminal renderer corruption (doubled/truncated characters) and a regression in hook‑based context injection (v1.0.60). Several model‑availability grievances resurface, while clipboard breakage on both Linux and Windows remains unresolved.

## Releases
No new releases in the last 24 hours. The latest stable version remains **v1.0.60** (2026-06-05).

## Hot Issues (10 noteworthy items)

1. **#53 – Bring back the GitHub Copilot in the CLI commands to not break workflows**  
   _Open since Sep 2025, 34 comments, 75 👍_  
   The most‑upvoted issue. GitHub has not responded for 6 months, leading the community to create `shell-ai` and other forks.  
   📎 [Issue #53](https://github.com/github/copilot-cli/issues/53)

2. **#1703 – Copilot CLI does not list all org-enabled models (e.g. Gemini 3.1 Pro) while VS Code Copilot does**  
   _Closed, 31 comments, 54 👍_  
   Highlights a persistent disparity in model availability between the CLI and VS Code for the same account/org.  
   📎 [Issue #1703](https://github.com/github/copilot-cli/issues/1703)

3. **#223 – “Copilot Requests” permission for fine‑grained tokens should be visible for org‑owned tokens**  
   _Open since Oct 2025, 29 comments, 76 👍_  
   Enterprises require org‑owned tokens; the missing permission blocks automated workflows.  
   📎 [Issue #223](https://github.com/github/copilot-cli/issues/223)

4. **#2082 – ctrl+shift+c no longer copies to clipboard on Linux**  
   _Open, 21 comments, 8 👍_  
   A fundamental UX regression in Ubuntu 24.04 – the standard clipboard shortcut is intercepted by the CLI.  
   📎 [Issue #2082](https://github.com/github/copilot-cli/issues/2082)

5. **#2334 – Please bring back no‑alt‑screen**  
   _Closed, 7 comments, 28 👍_  
   The forced alt‑screen mode removes scrollback, find, and selection capabilities, causing strong community pushback.  
   📎 [Issue #2334](https://github.com/github/copilot-cli/issues/2334)

6. **#3547 – Background sub‑agent silently hangs at total_turns=0 when model=”gpt‑5.5″**  
   _Closed, 7 comments_  
   A critical agent orchestration bug – sub‑agents dispatched in background never make progress.  
   📎 [Issue #3547](https://github.com/github/copilot-cli/issues/3547)

7. **#3596 – Error loading model list: Error: Not authenticated**  
   _Open, 5 comments, 10 👍_  
   Users lose authentication when resuming a specific session, forcing them to start fresh.  
   📎 [Issue #3596](https://github.com/github/copilot-cli/issues/3596)

8. **#3727 – Regression in v1.0.60: userPromptSubmitted hook additionalContext no longer injected into planner**  
   _Open, 3 comments_  
   A plugin hook that worked in v1.0.59 is broken in v1.0.60, breaking custom context injection.  
   📎 [Issue #3727](https://github.com/github/copilot-cli/issues/3727)

9. **#3749 – Terminal streaming renderer corrupts output – characters doubled/truncated**  
   _Open, 2 comments, 2 👍_  
   Streamed output (reasoning + final answer) shows duplicated/overlapping characters – a major usability blocker.  
   📎 [Issue #3749](https://github.com/github/copilot-cli/issues/3749)

10. **#3754 – `copilot --resume "Name With Spaces"` fails silently with exit 1**  
    _Open, 1 comment_  
    A contradiction to documented behaviour – session names with spaces cannot be resumed.  
    📎 [Issue #3754](https://github.com/github/copilot-cli/issues/3754)

## Key PR Progress
No pull requests were updated in the last 24 hours. The repository shows zero PR activity for this period.

## Feature Request Trends
- **Model diversity**: Multiple issues (#1664, #821, #2854, #2434) demand support for Gemini 3.1 Pro, Gemini 3 Flash, and other models that are already available in VS Code. The inconsistency frustrates users who prefer the CLI as their primary Copilot interface.
- **MCP improvements**: Requests include a power‑user shortcut for direct MCP tool invocation (#3752), and respect for custom providers in `--acp` mode (#3048).
- **UI/UX polish**: Restoring the `no-alt-screen` option (#2334) and keeping the agent picklist scroll position at the top (#3751) are small changes with large daily‑use impact.
- **Enterprise & permissions**: Organizations need clearer token‑permission visibility (#223) and finer control over MCP servers.

## Developer Pain Points
- **Missing model parity**: The CLI consistently lags behind VS Code in model availability, forcing users to switch tools.
- **Clipboard breakage**: `ctrl+shift+c` on Linux (#2082) and clipboard copy on Windows (#3622) are broken – a core terminal expectation.
- **Policy false positives**: MCP servers are blocked by policy even when no such policy exists (#1707, #3756), requiring downgrades or hacks.
- **Authentication fragility**: Resuming sessions can drop authentication (#3596), and session names with spaces cannot be resumed at all (#3754).
- **Terminal corruption**: The streaming renderer produces garbled output with duplicated/overlapping characters (#3749, #3755), making reviewed text unreadable.
- **Missing communication**: The top‑voted issue (#53) has received zero updates from GitHub for six months, eroding trust and prompting community forks.

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest – 2026-06-11

---

## 1. Today’s Highlights

No new release today, but the repo saw a burst of merged fixes from maintainer he‑yufeng—covering Windows console stability, shell process tree termination on timeout, and graceful handling of deferred MCP startup failures. Two new bugs were reported by iaindooley: approval prompts appearing in `yolo` mode and a final TODO item that never completes. Meanwhile, contributor Pluviobyte has two open PRs addressing orphaned tool calls in persisted sessions and improved shell command headline display.

---

## 2. Releases

No new versions published in the last 24 hours.

---

## 3. Hot Issues

Three issues were updated in the last 24 hours. The two open bugs are critical for users in autonomous workflows; the closed enhancement is minimal.

| # | Issue | Summary & Impact |
|---|-------|------------------|
| #2448 | [bug] Kimi CLI is prompting for approval in yolo mode | **Author:** iaindooley · **0 comments** · [Link](https://github.com/MoonshotAI/kimi-cli/issues/2448) — Yolo mode should skip all approval requests, but the agent keeps asking. Breaks automation and headless use cases. No workaround mentioned. |
| #2447 | [bug] Final Todo item never completes | **Author:** iaindooley · **0 comments** · [Link](https://github.com/MoonshotAI/kimi-cli/issues/2447) — The agent processes all TODOs except the last one, leaving tasks hanging. Likely affects any multi‑step execution. User on Debian with k2.6 model. |
| #2173 | [closed] [enhancement] ! | **Author:** odellus · **0 comments** · [Link](https://github.com/MoonshotAI/kimi-cli/issues/2173) — Closed without description. Might be a feature request for a shortcut or auto‑execution symbol. Minimal community engagement. |

---

## 4. Key PR Progress

Highlighted 10 PRs (all closed in the last 24h except two open by Pluviobyte). The closed PRs are all by he‑yufeng and address stability, Windows compatibility, and session handling.

| PR | Title & Summary | Link |
|----|-----------------|------|
| #2355 | Fix: continue after deferred MCP startup failures — Prevents a failing MCP server from aborting an interactive turn; logs the failure and continues without that server. Fixes #2343. | [PR #2355](https://github.com/MoonshotAI/kimi-cli/pull/2355) |
| #2354 | Fix: avoid shared rotating logs on Windows — Uses per‑process log files (`kimi.<pid>.log`) on Windows to prevent concurrent CLI/Web/worker processes from corrupting the same open log file. | [PR #2354](https://github.com/MoonshotAI/kimi-cli/pull/2354) |
| #2327 | Fix: terminate shell process trees on timeout — Runs foreground shell commands in their own process group and kills the entire tree on timeout/cancellation. Fixes #2310. | [PR #2327](https://github.com/MoonshotAI/kimi-cli/pull/2327) |
| #2289 | Fix: avoid Windows console font reset — Passes `CREATE_NO_WINDOW` when Kaos spawns subprocesses on Windows, preventing a separate console window from stealing focus. Fixes #2197. | [PR #2289](https://github.com/MoonshotAI/kimi-cli/pull/2289) |
| #2288 | Fix: avoid resending web uploads after restart — Persists a `.sent` marker so already‑uploaded files aren’t re‑attached after a process restart. Fixes #2279. | [PR #2288](https://github.com/MoonshotAI/kimi-cli/pull/2288) |
| #2239 | Fix: continue latest persisted session — Falls back to the newest non‑empty session when metadata has no `last_session_id` or points to a stale one. Fixes #2222. | [PR #2239](https://github.com/MoonshotAI/kimi-cli/pull/2239) |
| #2235 | Fix: omit empty tools in OpenAI legacy requests — Stops sending `tools: []` to backends like vLLM that reject empty tool arrays. Fixes #2233. | [PR #2235](https://github.com/MoonshotAI/kimi-cli/pull/2235) |
| #2217 | Fix: recover background auto‑trigger after cooldown — Pauses background auto‑trigger for 10 minutes after 3 consecutive failures, then resets. Keeps user input responsive. Fixes #2193. | [PR #2217](https://github.com/MoonshotAI/kimi-cli/pull/2217) |
| #2211 | Fix(web): propagate afk mode to workers — Ensures `kimi --afk web` runs worker subprocesses without interactive approval, so tool calls don’t hang waiting for user input. Closes #2202? (implied) | [PR #2211](https://github.com/MoonshotAI/kimi-cli/pull/2211) |
| #2387 | [open] Fix(tools): preserve shell command headline details — Improves display of long shell commands (e.g., `grep` with many arguments) by not truncating the middle. Resolves #2142. | [PR #2387](https://github.com/MoonshotAI/kimi-cli/pull/2387) |

Other notable open PRs by Pluviobyte:
- #2383 – fix(soul): repair orphan tool_calls when replaying history (after mid‑turn kills) – resolves #2336.
- #2386 – fix(session): map undo wire turns to context turns – fixes #1974 and #2049.

---

## 5. Feature Request Trends

Only one enhancement issue (#2173) appeared in the last 24h, and it carries no description. From the broader PR landscape, the most requested feature directions over the past weeks are:

- **Improved Yolo / Autonomous behavior** – The two open bugs (#2448, #2447) indicate that users expect a truly hands‑off mode without approval prompts or stuck TODOs.
- **Better session persistence after crashes** – PRs #2383 and #2386 address mid‑turn crashes leaving orphaned tool calls or corrupt undo history—suggests users are hitting reliability issues in long‑running sessions.
- **Windows console compatibility** – Multiple fixes in the last 24h (log sharing, console font reset, subprocess windows) show a strong community need for first‑class Windows support.

No major new feature requests (e.g., new model support, plugin API) surfaced.

---

## 6. Developer Pain Points

Recurring frustrations observed from the latest issues and PRs:

- **Yolo mode still requires approval** – Despite being the primary headless workflow, it intermittently asks for permission, breaking CI/CD and automated pipelines.
- **Task completion never finishes** – The final TODO in a multi‑step plan remains incomplete, forcing manual intervention.
- **Windows‑specific crashes** – Console font reset, log file corruption, and separate console windows for subprocesses have required several targeted fixes.
- **Session corruption after abrupt termination** – Mid‑turn kills (OOM, `kill -9`, terminal close) leave orphaned tool calls that prevent replaying or continuing the conversation.
- **Background auto‑trigger unreliability** – After consecutive failures, the background agent stops entirely; users want better recovery and clearer logging.
- **Inconsistent `--continue` behavior** – The fallback to newest session (#2239) is now fixed, but the original bug caused lost work.

The community is heavily focused on stability and platform parity rather than new functionality. The absence of feature requests suggests satisfaction with the current capabilities once the pain points are resolved.

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-11

## Today’s Highlights
Three patch releases (v1.17.1–v1.17.3) landed in the last 24 hours, fixing a desktop crash, Linux launcher identity, and remote config auth recovery. The community is buzzing about a proposed `/goal` native session goals feature (69 👍) and a “YOLO Mode” for auto-approving permissions (29 👍). Meanwhile, several model-specific bugs (Cerebras GLM-4.7, Copilot Opus 4.8) and web UI regressions are drawing active debugging efforts.

## Releases
- **[v1.17.3](https://github.com/anomalyco/opencode/releases/tag/v1.17.3)** – Urgent fix for desktop crash introduced in v1.17.2.
- **[v1.17.2](https://github.com/anomalyco/opencode/releases/tag/v1.17.2)** – Core: recover from expired remote config auth (re-login instead of failure); allow subagents to use their own permissions. Desktop: restore Linux launcher and icon identity so pinned apps keep opening `co`.
- **[v1.17.1](https://github.com/anomalyco/opencode/releases/tag/v1.17.1)** – Core: references can include usage descriptions for agents, appear in new docs, and be hidden from `@` autocomplete. Bugfixes: deprecated `reference` config entries load under newer `references` key; MCP prompt and resource support.

## Hot Issues (10 picks)

1. **[#27167 – Feature: Add native session goals with /goal](https://github.com/anomalyco/opencode/issues/27167)**  
   *40 comments, 69 👍*  
   Proposes a persistent session goal/lifecycle feature. High community demand for built-in planning support.  
   
2. **[#26762 – Cerebras zai-glm-4.7 fails on follow-up turn with reasoning_content](https://github.com/anomalyco/opencode/issues/26762)**  
   *10 comments*  
   Multi-turn runs error on `reasoning_content` property. Blocks users relying on Cerebras for reasoning-heavy workflows.

3. **[#6490 – Web UI cannot browse folders outside default user profile](https://github.com/anomalyco/opencode/issues/6490)**  
   *10 comments, 12 👍*  
   Windows users stuck with restricted folder picker; no path entry. Severely impacts cross-drive project setup.

4. **[#11831 – YOLO Mode: Auto-approve all permission prompts](https://github.com/anomalyco/opencode/issues/11831)**  
   *9 comments, 29 👍*  
   Power users want a trust mode that bypasses permission dialogs while respecting explicit denys. Trend toward less friction.

5. **[#30086 – High CPU usage in newer versions](https://github.com/anomalyco/opencode/issues/30086)**  
   *9 comments*  
   Spike since ~1.17.x makes multi-session usage laggy. Top concern for heavy users.

6. **[#31247 – Opus 4.8 leaks literal tool-call text into assistant messages](https://github.com/anomalyco/opencode/issues/31247)**  
   *8 comments*  
   `github-copilot/claude-opus-4.8` emits raw `call read`, `<invoke>` markup. Breaks conversation history.

7. **[#30158 – Terminal button in web UI disappears since v1.15.12](https://github.com/anomalyco/opencode/issues/30158)**  
   *7 comments, 6 👍*  
   Regression blocks web terminal access. Users forced to downgrade to v1.15.11.

8. **[#31182 – TUI /sessions search does not filter results](https://github.com/anomalyco/opencode/issues/31182)**  
   *5 comments, 7 👍*  
   Search field appears broken; unfiltered session list shown. Blocks session management in TUI.

9. **[#29309 – Feature: Add Vietnamese language](https://github.com/anomalyco/opencode/issues/29309)**  
   *5 comments*  
   Community volunteer offering translations. Indicates growing international user base.

10. **[#31812 – xfyun engine busy response not retried](https://github.com/anomalyco/opencode/issues/31812)**  
    *4 comments*  
    Transient overload errors cause request failure. Already fixed in PR #31819 (see below).

## Key PR Progress (10 picks)

1. **[#31819 – fix(opencode): retry on xfyun engine busy response](https://github.com/anomalyco/opencode/pull/31819)**  
   Adds `engine busy` to the retryable errors list. Patch for a frustrating transient fault.

2. **[#31827 – test(opencode): simplify snapshot race layer wiring](https://github.com/anomalyco/opencode/pull/31827)**  
   Refactors test infrastructure to use `LayerNode` graph. Part of ongoing architectural cleanup.

3. **[#31826 – refactor(tui): replace v2 sync with data context](https://github.com/anomalyco/opencode/pull/31826)**  
   Migration from `sync-v2` to a private `DataProvider` – improves TUI reactivity and maintainability.

4. **[#31823 – test(opencode): simplify processor layer wiring](https://github.com/anomalyco/opencode/pull/31823)**  
   Another LayerNode conversion for test environment. Shared fixtures across processor tests.

5. **[#31822 – feat(server): add v2 session API endpoints](https://github.com/anomalyco/opencode/pull/31822)**  
   New v2 location resolution, session create/get, pending questions. Includes autogenerated SDK. Foundation for upcoming client features.

6. **[#31796 – tui 2.0](https://github.com/anomalyco/opencode/pull/31796)**  
   Large refactor/reimagining of the TUI – likely to bring improved session management, search, and performance.

7. **[#31805 – fix(tui): preserve exit epilogue during scoped shutdown](https://github.com/anomalyco/opencode/pull/31805)**  
   Restores the “Continue opencode -s <id>” banner on TUI exit, which was silently dropped.

8. **[#13610 – feat(desktop): add keyboard shortcuts to switch projects (Cmd+1-9)](https://github.com/anomalyco/opencode/pull/13610)**  
   Long-running PR – adds project tab switching similar to browser tabs. Finally merged.

9. **[#31817 – fix(core): add compaction key to isV1 detection](https://github.com/anomalyco/opencode/pull/31817)**  
   Fixes silent drop of `preserve_recent_tokens` when config uses deprecated `compaction` field.

10. **[#31745 – fix(opencode): surface content-filter finish reason as visible error](https://github.com/anomalyco/opencode/pull/31745)**  
    When Anthropic models refuse with `stop_reason: refusal`, the error is now shown to the user instead of silently failing.

## Feature Request Trends
- **Session lifecycle & planning**: `/goal` command and native session goals (#27167) top the charts. Users want persistent, editable objectives that survive across turns.
- **Permission automation**: “YOLO Mode” (#11831) and auto-approve from localhost (#31820) reflect desire for less friction in trusted environments.
- **Internationalization**: Vietnamese (#29309) and Chinese invoicing (#30716) show increasing global adoption.
- **UI/UX refinement**: requests for collapsible toolbars (#31818), resizable status panel (#24373), per-message expandable reasoning (#31825), and better container/headless support (#31815, #31804).
- **Plugin API expansion**: proposal for `ensureServer()` (#31821) indicates plugin developers need control over HTTP servers.

## Developer Pain Points
- **High CPU usage** (#30086) since ~v1.17.x makes multi-session work painful; root cause still unconfirmed.
- **Web UI regressions**: terminal button disappearance (#30158) and folder picker limitation (#6490) block productive use for Windows users.
- **Model-specific bugs**: Cerebras reasoning_content (#26762), Copilot Opus 4.8 tool-call leaks (#31247), and xfyun transient errors (#31812) erode trust in newer providers.
- **TUI flaws**: session search not filtering (#31182), missing exit banner (#31813, fixed in PR #31805), and shell timeout off by 100ms (#31806) degrade daily workflow.
- **Installation friction**: hidden cursor on Ctrl+C (#31829) and missing shell reload guidance (#18624) frustrate first-time users.
- **Platform quirks**: Windows exit hang (#25677), IPv6 vs IPv4 OAuth callback (#31824), and `xdg-open` error in containers (#31815) hint at insufficient cross-platform testing.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-06-11

No new releases were published in the last 24 hours. The community remains highly active with 36 issues and 12 pull requests landing – many focused on provider interoperability, streaming reliability, and TUI stability. A wave of crash fixes and provider extensions (Palantir Foundry, Amazon Bedrock Mantle) signals the platform maturing toward enterprise and multi-backend use.

## Releases
None in the last 24 hours.

## Hot Issues

1. **#5514 – Trust Gating Feature Backlash**  
   *Closed • 25 comments • 13 👍*  
   Users are pushing back on the just-landed trust gating feature, complaining about being asked to confirm trusted folders repeatedly across machines. The community is split between security-conscious and productivity-focused users.  
   [earendil-works/pi Issue #5514](https://github.com/earendil-works/pi/issues/5514)

2. **#3715 – Local-LLM Streams Terminate After 5 Minutes**  
   *Closed • 10 comments • 4 👍*  
   Long `Write` tool calls against local vLLM/OpenAI-compatible backends hit `UND_ERR_BODY_TIMEOUT` because the underlying `undici` body timeout cannot be overridden by `retry.provider.timeoutMs`. Significant for anyone using reasoning models locally.  
   [earendil-works/pi Issue #3715](https://github.com/earendil-works/pi/issues/3715)

3. **#4160 – pi Extensions Fail Under Bun Without npm**  
   *Closed • 9 comments*  
   Using Bun as the sole runtime leads to `Executable not found in $PATH: "npm"` when installing extensions. The community wants full Bun-native extension management.  
   [earendil-works/pi Issue #4160](https://github.com/earendil-works/pi/issues/4160)

4. **#5025 – Multi-Select-List UI Component for Extensions**  
   *Closed • 5 comments • 2 👍*  
   A feature request for a core UI component to let extension authors present multi-select options (e.g., for editing `models.json`). Addresses a clear gap in the extension SDK.  
   [earendil-works/pi Issue #5025](https://github.com/earendil-works/pi/issues/5025)

5. **#5291 – Sessions Hang on “Working…” with Anthropic Enterprise**  
   *Closed • 5 comments • 1 👍*  
   Intermittent session freezes when using pi.dev with an Anthropic Enterprise subscription. Users must intervene manually, indicating a potential integration race condition.  
   [earendil-works/pi Issue #5291](https://github.com/earendil-works/pi/issues/5291)

6. **#5611 – GitLab Duo Anthropic Stream Cutoff (~90s)**  
   *Closed • 3 comments*  
   Anthropic streams via GitLab Duo close before `message_stop`, triggering erroneous retries with the same payload. Affects Opus 4.8 extended thinking.  
   [earendil-works/pi Issue #5611](https://github.com/earendil-works/pi/issues/5611)

7. **#5536 – Split-Turn Compaction Causes 429 on Local Backends**  
   *Open • 2 comments*  
   When Pi performs split-turn compaction, it sends two summarization requests concurrently, overwhelming single-slot local backends (e.g., `llama.cpp`) with `429 Too Many Requests`.  
   [earendil-works/pi Issue #5536](https://github.com/earendil-works/pi/issues/5536)

8. **#5577 – Persona Override for System Prompt**  
   *Closed • 2 comments*  
   Users want to define custom agent personas (e.g., security tester, QA) without losing Pi’s coding‑focused system prompt. High demand as Pi is used for non‑coding tasks.  
   [earendil-works/pi Issue #5577](https://github.com/earendil-works/pi/issues/5577)

9. **#5605 – MiniMax‑M3 Cache Control Ignored (Billing Impact)**  
   *Closed • 2 comments*  
   For MiniMax‑M3 routed via the Anthropic endpoint, `cache_control` is silently dropped, causing full input billing instead of reduced cache-read pricing. Community workaround needed.  
   [earendil-works/pi Issue #5605](https://github.com/earendil-works/pi/issues/5605)

10. **#5592 – Anthropic Streams Wait for Transport EOF After `message_stop`**  
    *Closed • 2 comments*  
    The SSE iterator keeps reading past `message_stop`, causing latency and resource leaks when a proxy keeps the response open. Critical for gateway setups.  
    [earendil-works/pi Issue #5592](https://github.com/earendil-works/pi/issues/5592)

## Key PR Progress

1. **#5609 – Palantir Foundry LLM Proxy & OAuth Provider**  
   *(Closed)*  
   Adds full Palantir Foundry support, including OAuth token handling and reasoning variants. Opens Pi to enterprise Foundry ecosystems.  
   [earendil-works/pi PR #5609](https://github.com/earendil-works/pi/pull/5609)

2. **#5594 – Fix Anthropic Stream Finalization on `message_stop`**  
   *(Closed)*  
   Directly addresses #5592 by treating `message_stop` as the logical end and cancelling the body reader, releasing the transport immediately.  
   [earendil-works/pi PR #5594](https://github.com/earendil-works/pi/pull/5594)

3. **#5509 – Amazon Bedrock Mantle (OpenAI Responses) Provider**  
   *(Open)*  
   Adds support for Bedrock Mantle, a new API that serves GPT‑5.5 and 5.4 models. Modeled after the Azure OpenAI Responses provider.  
   [earendil-works/pi PR #5509](https://github.com/earendil-works/pi/pull/5509)

4. **#5587 – Experimental First‑Time Setup Flow**  
   *(Closed)*  
   Behind `PI_EXPERIMENTAL=1`, shows a dark/light mode picker and opt‑in analytics dialog on first launch. Streamlines onboarding.  
   [earendil-works/pi PR #5587](https://github.com/earendil-works/pi/pull/5587)

5. **#5583 – Preserve Clickable Subscription Login URLs**  
   *(Closed)*  
   Fixes broken long login URLs caused by default left padding – a small but UX‑critical fix for provider authentication.  
   [earendil-works/pi PR #5583](https://github.com/earendil-works/pi/pull/5583)

6. **#5561 – Link AWS Data Retention Docs in Bedrock Validation Errors**  
   *(Closed)*  
   When Claude Fable 5 fails due to missing data retention, the error now points directly to AWS documentation, reducing support friction.  
   [earendil-works/pi PR #5561](https://github.com/earendil-works/pi/pull/5561)

7. **#5585 – CJK Text Wrapping at Character Boundaries in Editor**  
   *(Closed)*  
   Fixes word‑level wrapping that broke CJK text. Essential for multilingual users.  
   [earendil-works/pi PR #5585](https://github.com/earendil-works/pi/pull/5585)

8. **#5562 – Loose List Rendering with Blank Lines**  
   *(Closed)*  
   Per CommonMark spec, adds blank lines between list items in loose lists. Improves markdown rendering accuracy.  
   [earendil-works/pi PR #5562](https://github.com/earendil-works/pi/pull/5562)

9. **#5560 – Parse `:thinking` Suffix from Custom Model IDs**  
   *(Closed)*  
   Prevents the `:thinking` suffix from leaking into the model ID when the model isn’t in Pi’s built‑in registry. Fixes a 404 regression.  
   [earendil-works/pi PR #5560](https://github.com/earendil-works/pi/pull/5560)

10. **#5586 – Bedrock Use `apiKey` as Bearer‑Token Fallback**  
    *(Closed)*  
    Allows `models.json` `apiKey` to be used as a bearer token for Bedrock when no other token provider is configured, enabling AI gateway scenarios.  
    [earendil-works/pi PR #5586](https://github.com/earendil-works/pi/pull/5586)

## Feature Request Trends

- **Extension SDK expansion**: Community repeatedly requests richer UI components (multi‑select list, custom OAuth callbacks, extension command events). The extension ecosystem is a growing focus.
- **Agent personality customization**: Demand for user‑defined system prompt overrides (personas) suggests Pi is being used beyond coding – for security, QA, research, PM.
- **Provider parity**: New providers (Palantir, Bedrock Mantle) and per‑model thinking‑level controls are seen as essential for enterprise adoption. There is also interest in global max‑thinking limits.
- **RPC and API completeness**: Missing SDK/API commands (e.g., `clear_queue`) and the ability to queue steering messages are gaps for automation workflows.

## Developer Pain Points

- **Timeout and streaming fragility**: Multiple issues (#3715, #5611, #5592, #5536) highlight how hardcoded timeouts and suboptimal stream finalization disrupt local LLM and proxy‑based deployments.
- **Crash‑prone TUI**: Several uncaught exceptions (#5604, #5599, #5597) in the TUI’s rendering pipeline cause hard process termination – a critical UX problem for a terminal‑first tool.
- **Bun/npm incompatibility**: The extension installer assumes `npm` is present, locking out users on Bun‑only environments.
- **Misleading update banner**: The `pi update` recommendation appears even on installs that can’t self‑update (e.g., Nix‑store), wasting user time.
- **Billing surprises**: Cache‑control not being respected (MiniMax‑M3, Anthropic 1‑hour cache) leads to under‑reported costs and higher bills – a sensitive area for power users.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-06-11

A quiet but focused day on the Qwen Code repository. No new releases landed in the last 24 hours, but three active issues and a deep bench of open pull requests signal clear developer attention on **terminal UI robustness**, **daemon-mode hardening**, and **collaboration workflow quality-of-life** improvements.

---

## Releases

No new releases in the last 24 hours.

---

## Hot Issues

Picked from 3 active items updated in the last 24h. While the count is low, each issue touches a distinct and meaningful pain point.

### 1. #4976 — Auto-generated memory interferes with normal CLI calls [OPEN]
- **Priority:** P2 | **Tags:** type/bug, category/tools, scope/memory
- **Summary:** A user reports that auto-generated memory content corrupts their CLI workflow when calling tools directly. The issue logs a detailed timeline showing wasted rounds of failed tool calls, highlighting a mismatch between the model’s internal memory and user intent during CLI sessions.
- **Community:** 2 comments, no upvotes yet. The detailed user narrative suggests a moderate impact for power users relying on raw tool invocations.
- **Why it matters:** This reveals a friction point between the "helpful assistant" memory layer and the developer expectation of deterministic CLI behavior. Fixing this would improve trust in tool-calling transparency.
- **Link:** [Issue #4976](https://github.com/QwenLM/qwen-code/issues/4976)

### 2. #4974 — SGR mouse wheel sequences leak as typed text (`64;50;15M`) into the input box [OPEN]
- **Priority:** P2 | **Tags:** type/bug, category/cli, scope/interactive, scope/keybindings, scope/rendering
- **Summary:** When SGR mouse tracking is enabled, scroll events are double-consumed: the ink `useInput` pipeline handles them correctly for scrolling, but `KeypressContext` feeds the same bytes through, causing raw escape sequences to appear as typed text.
- **Community:** 2 comments. The author (wenshao) provides a precise root-cause analysis.
- **Why it matters:** This is a terminal compatibility bug that could break the interactive CLI experience for users on SGR-capable terminals (most modern ones). It degrades perceived polish.
- **Link:** [Issue #4974](https://github.com/QwenLM/qwen-code/issues/4974)

### 3. #4973 — Terminal drops to cooked mode (all input dead until Enter) when the last ink `useInput` deactivates [OPEN]
- **Priority:** P1 | **Tags:** type/bug, category/cli, scope/interactive, scope/keybindings, status/ready-for-agent
- **Summary:** `KeypressContext` only acquires raw mode if stdin is not already raw. When the last ink `useInput` deactivates, the refcount drops, raw mode is released, and the terminal falls back to cooked mode — making all input unresponsive until the user presses Enter.
- **Community:** 1 comment. Marked as "ready-for-agent" — the maintainers likely plan to auto-fix this.
- **Why it matters:** This is a critical UX bug (P1). The CLI becomes effectively unusable for keystroke-driven interaction after certain UI transitions. It blocks developers who rely on rich input modes.
- **Link:** [Issue #4973](https://github.com/QwenLM/qwen-code/issues/4973)

---

## Key PR Progress

Picked from 50 PRs updated in the last 24h. Selection focuses on functional impact, novelty, and strategic importance.

### 1. #4984 — feat(web-shell): add expand toggle to shell tool output [OPEN]
- **Author:** wenshao
- **Summary:** Adds an expand/collapse toggle for long shell output in the web shell, mirroring existing toggles for read output (added previously) and thinking/sub-agent streams (#4977). Removes the fixed 5-line clamp.
- **Why it matters:** Completes a set of UX consistency improvements for the web shell. Developers managing long-running shell commands will benefit from being able to review full output without scrolling limits.
- **Link:** [PR #4984](https://github.com/QwenLM/qwen-code/pull/4984)

### 2. #4853 — feat(core): add `enter_plan_mode` tool and Plan Approval Gate [OPEN]
- **Author:** callmeYe
- **Summary:** Introduces a no-arg tool (`enter_plan_mode`) that the model can call to proactively enter plan mode for complex or under-specified tasks. Also adds a Plan Approval Gate: when pre-plan mode is AUTO or YOLO, `exit_plan_mode` triggers a single approval step before execution.
- **Why it matters:** This is a significant step toward making the model's planning behavior explicit and user-controlled. It aligns with the "agent as collaborator" paradigm, giving users a clear gating point for approving suggested plans.
- **Link:** [PR #4853](https://github.com/QwenLM/qwen-code/pull/4853)

### 3. #4490 — feat(daemon): merge daemon-mode feature batch into main [OPEN]
- **Author:** doudouOUC
- **Summary:** A large integration merge (46 commits, +115k / −12k LOC) bringing the core daemon-mode feature set into `main` for v0.16-alpha. Covers background session management, workspace hot-reload, and multi-session persistence.
- **Why it matters:** This is the primary delivery vehicle for the daemon mode — a foundational infrastructure change. Its eventual merge will enable persistent background AI workflows.
- **Link:** [PR #4490](https://github.com/QwenLM/qwen-code/pull/4490)

### 4. #4954 — fix(serve): isolate per-session stats in daemon mode [CLOSED]
- **Author:** doudouOUC
- **Summary:** Fixes a bug where `GET /session/:id/stats` returned process-wide cumulative metrics instead of per-session data. Implements a `Map<sessionId, SessionMetrics>` dual-write pattern inside `UiTelemetryService`.
- **Why it matters:** Essential correctness fix for daemon mode. Without this, users can't trust session-level telemetry — a blocker for billing, usage tracking, and debugging.
- **Link:** [PR #4954](https://github.com/QwenLM/qwen-code/pull/4954)

### 5. #4965 — feat(daemon): add `POST /workspace/reload` for unified settings hot-reload [OPEN]
- **Author:** doudouOUC
- **Summary:** Adds a single endpoint that hot-reloads all settings to idle sessions, replacing the narrower `POST /workspace/reload-env`. Features settings diff detection via `diffSettingsKeys` to only refresh changed values.
- **Why it matters:** Simplifies the daemon admin API while making it smarter. Operators can now reload workspace configuration with minimal disruption.
- **Link:** [PR #4965](https://github.com/QwenLM/qwen-code/pull/4965)

### 6. #4981 — fix(core): serialize team task claims per agent and add mailbox lock parity [OPEN]
- **Author:** tanzhenxin
- **Summary:** Prevents concurrent auto-claim from assigning two tasks to the same teammate simultaneously. Also brings task store file locking in line with the team mailbox implementation.
- **Why it matters:** Critical for reliability of multi-agent teams. Without this, race conditions could cause duplicate work or dropped tasks.
- **Link:** [PR #4981](https://github.com/QwenLM/qwen-code/pull/4981)

### 7. #4982 — fix(core): remove dead `debugResponses` array to prevent OOM [OPEN]
- **Author:** zzhenyao
- **Summary:** Removes the unused `debugResponses` array that accumulated every streaming chunk into memory, and the dead `extractUsageFromGeminiClient` export. Nothing in production ever read either.
- **Why it matters:** A straightforward memory leak fix. Under long sessions, this could silently consume gigabytes of RAM.
- **Link:** [PR #4982](https://github.com/QwenLM/qwen-code/pull/4982)

### 8. #4773 — feat(serve): ACP WebSocket transport (RFD Streamable HTTP phase 2) [OPEN]
- **Author:** chiga0
- **Summary:** Implements a complete ACP WebSocket transport, coexisting with SSE. Adds a transport-agnostic interface (`transportStream.ts`), a WebSocket adapter (`wsStream.ts`), and a connection registry.
- **Why it matters:** Enables bidirectional streaming for ACP — a prerequisite for real-time collaboration features. This is part of a larger roadmap toward standardized agent communication protocols.
- **Link:** [PR #4773](https://github.com/QwenLM/qwen-code/pull/4773)

### 9. #4896 — fix(core): stabilize prompt-cache prefix against MCP/skills churn [OPEN]
- **Author:** callmeYe
- **Summary:** Decouples skill visibility (what the model sees) from skill validation (what the tool accepts), placing them in cache-appropriate tiers. Mid-session skill/MCP changes no longer invalidate the entire prompt cache.
- **Why it matters:** Prompt cache invalidation is a major performance cost. This PR makes caching significantly more stable in agentic workflows that dynamically load or unload skills.
- **Link:** [PR #4896](https://github.com/QwenLM/qwen-code/pull/4896)

### 10. #4598 — feat(tui): collapsible thinking blocks with duration timer [OPEN]
- **Author:** chiga0
- **Summary:** Replaces the always-expanded thinking display with a collapsible history block that streams reasoning above the answer and collapses on completion, with a duration timer shown.
- **Why it matters:** This is a major TUI UX improvement. It reduces visual clutter while preserving access to the model's reasoning trace — a win for power users and newcomers alike.
- **Link:** [PR #4598](https://github.com/QwenLM/qwen-code/pull/4598)

---

## Feature Request Trends

Based on tags and PR descriptions across all items:

1. **Plan Mode & Approval Gates** — Users want models to proactively yield control for complex tasks. The `enter_plan_mode` proposal (#4853) is a strong signal. Expect more requests for gating mechanisms that let developers "approve before execute."
2. **Daemon Mode as a Platform** — Multiple PRs (#4490, #4965, #4954) point to daemon mode becoming the backbone for persistent, multi-session workflows. Users are asking for hot-reload, per-session stats, and admin APIs — typical infrastructure requests once a system moves from experimental to production.
3. **Rich Terminal Interactivity** — Collapsible blocks (#4598), expandable shell output (#4984), and request for `/compress-fast` (#4893) show demand for a more modern, responsive CLI/TUI — not just a chat client but a "developer shell" that manages context and output elegantly.
4. **Multi-Agent Collaboration Tooling** — PR #4981 (team task serialization) and #4979 (preserving teammate identity after approval) suggest growing interest in reliable multi-agent teams. Feature requests around delegation, identity, and logging are likely to increase.

---

## Developer Pain Points

Recurring themes from issue comments and PR descriptions:

1. **Terminal Mode Flickering** — Issues #4973 and #4974 both relate to raw-mode toggling breaking input. This is a high-frequency pain point for CLI-heavy users. Expect more of these as terminal emulator diversity increases.
2. **Memory Leaks in Long Sessions** — The `debugResponses` array fix (#4982) and the tool output memory reduction PR (#4971) point to a broader concern: developers running long, multi-turn sessions are hitting memory limits. This will only grow as agentic workflows become longer.
3. **Configuration Drift in Daemon Mode** — The settings path bug (#4938) and the need for a unified hot-reload endpoint (#4965) indicate that managing configuration across multiple daemon sessions is currently brittle. Operators want a single source of truth that is easy to update.
4. **Trust in Tool Output** — Issue #4976 (auto-generated memory interfering) highlights a lack of transparency: developers need to understand *why* a tool call happened, especially when it wasn't explicitly requested. This echoes a broader need for "explainability" in agentic tool use.
5. **API Consistency Across Transports** — The SSE vs WebSocket divergence in error handling (#4952) and the need for ACP/REST parity (#4827) show that developers integrating via different transports expect identical behavior — any asymmetry is treated as a bug.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI / CodeWhale Community Digest – 2026-06-11

---

## 1. Today's Highlights

The project continues its aggressive rebrand from `deepseek-tui` to **CodeWhale**, with v0.8.57 explicitly deprecating the legacy npm package. The development pace has accelerated significantly: a massive v0.8.58 branch is now live, carrying a constitution refactor, Codex reliability improvements, and sidebar UI enhancements. The "all-models" push is the dominant theme — multiple issues and PRs target un-hardcoding DeepSeek-specific assumptions from routing, subagent model selection, and system prompts to support Moonshot, Ollama, OpenAI, and others.

---

## 2. Releases

Two releases in the last 24 hours:

- **[v0.8.57](https://github.com/Hmbown/CodeWhale/releases/tag/v0.8.57)** — Pure rebranding release: `CodeWhale` is now the canonical name. The legacy `deepseek-tui` npm package is deprecated with no further releases. Migration instructions available in `docs/REBRAND.md`.

- **[v0.8.56](https://github.com/Hmbown/CodeWhale/releases/tag/v0.8.56)** — "Community Harvest": localization improvements, new provider support, prefix-cache stability fixes, and general bug fixes.

---

## 3. Hot Issues (10 noteworthy)

1. **[#2369 — Config Paths Fragmented Across OS and Cygwin](https://github.com/Hmbown/CodeWhale/issues/2369)**  
   A deep-dive bug report (6 comments) documenting inconsistent config resolution across Linux, macOS, Windows, and Cygwin — plus a silent migration bug where old paths are silently used. Critical for cross-platform users.

2. **[#1679 — SSE Multi-Agent 45s Timeout on Windows + UI Corruption](https://github.com/Hmbown/CodeWhale/issues/1679)**  
   Windows-specific issue: 4 agents probe successfully but fail during execution with 45s SSE timeouts. Agent auto-degrades to solo mode. Also reports UI layout corruption. Updated today.

3. **[#2574 — Provider Fallback Chain (Auto-Switch on API Failure)](https://github.com/Hmbown/CodeWhale/issues/2574)**  
   3 comments, high demand: users want automatic fallback when a provider returns 401/429/5xx instead of manual `/provider` switching. Config proposal included.

4. **[#3018 — Un-hardcode DeepSeek from Auto-Router and Subagent Model Selection](https://github.com/Hmbown/CodeWhale/issues/3018)**  
   Core blocker for multi-provider support: the flash-router hardcodes `deepseek-v4-flash` and subagent model validation rejects non-DeepSeek IDs. Filed as v0.8.58 target.

5. **[#3025 — Parameterize Model-Specific Facts in Constitution Prompt](https://github.com/Hmbown/CodeWhale/issues/3025)**  
   The new constitution prompt tells every model it's "DeepSeek V4 with 1M context" — misleading for Kimi, Qwen, GPT, Ollama users. Needs runtime capability lookup.

6. **[#2989 — Agent Reports "Completed" But Task Not Actually Finished (Ollama + qwen3-coder)](https://github.com/Hmbown/CodeWhale/issues/2989)**  
   Serious reliability bug: agent stops mid-task but status shows "completed." Likely a silent model output truncation issue with local models.

7. **[#2964 — Ship DigitalOcean + Telegram Remote Workbench Setup](https://github.com/Hmbown/CodeWhale/issues/2964)**  
   Core maintainer issue: building a repeatable "15-minute remote workbench" for non-China users using a cheap US VPS + Telegram bridge. Updated today with 2 comments.

8. **[#3031 — Compact Tool-Call Transcript Rendering](https://github.com/Hmbown/CodeWhale/issues/3031)**  
   UX polish: tool-call cells show too much low-value detail by default (wrapped commands, `(no output)` lines, sub-second timing). Filed for v0.8.58.

9. **[#3007 — TUI Provider Rejection Blames Non-Existent `--provider` Flag](https://github.com/Hmbown/CodeWhale/issues/3007)**  
   Misleading error message: when provider comes from config/env (not CLI), the error still says "remove `--provider`". Closed with fix.

10. **[#2960 — Rebrand Update Path Broken: `deepseek update` Fails Silently](https://github.com/Hmbown/CodeWhale/issues/2960)**  
    Users on old installs hit hard errors with no migration guidance. Covers multiple failure modes (npm, Cargo, GitHub Releases).

---

## 4. Key PR Progress (10 important)

1. **[#3045 — Fix Subagent Model Validation for Non-DeepSeek Providers](https://github.com/Hmbown/CodeWhale/pull/3045)**  
   Partially addresses #3018: un-hardcodes DeepSeek from model validation, letting Moonshot/Ollama/OpenAI use their own model IDs for sub-agents.

2. **[#3034 — v0.8.58 Branch: Constitution Refactor, Codex Fixes, Sidebar Improvements](https://github.com/Hmbown/CodeWhale/pull/3034)**  
   The big v0.8.58 merge: YAML-based constitution with Python renderer, TUI sidebar split (Models/Resources tabs), Codex client reliability fixes.

3. **[#3038 — Ctrl+B Directly Backgrounds Foreground Shell](https://github.com/Hmbown/CodeWhale/pull/3038)**  
   UX simplification: removes two-step `ShellControlView` menu for backgrounding. Closes #3032.

4. **[#3037 — Compact Tool-Call Transcript Rendering](https://github.com/Hmbown/CodeWhale/pull/3037)**  
   Suppresses `(no output)` lines, sub-second timing, and excessive spacing in default Live transcript view. References #3031.

5. **[#3035 — Throttle AgentProgress Redraws to Prevent Freeze](https://github.com/Hmbown/CodeWhale/pull/3035)**  
   Fixes UI freeze when 4+ sub-agents run concurrently: throttles redraws instead of triggering full terminal redraw on every progress event.

6. **[#3049 — Hooks v2: JSON Decision Contract, Glob Matchers, Project-Local Hooks](https://github.com/Hmbown/CodeWhale/pull/3049)**  
   Major hooks upgrade: hooks can now emit `allow`/`deny`/`ask` decisions as JSON, support glob-based file matching, and read from `.codewhale/hooks/` directory.

7. **[#3042 — Headless Exec: `--allowed-tools`, `--max-turns`, `--append-system-prompt`](https://github.com/Hmbown/CodeWhale/pull/3042)**  
   Adds four new CLI flags for CI/benchmark use: tool allow/deny lists, turn caps, and system prompt appending for `codewhale exec`.

8. **[#3051 — `/voice` Slash Command for Speech-to-Text](https://github.com/Hmbown/CodeWhale/pull/3051)**  
   Community contribution: voice input support inspired by MiMo Code. Three slash commands for recording, transcription, and composer insertion via existing provider API.

9. **[#3048 — Parameterize Model-Specific Facts (Context Window, Pricing, Thinking)](https://github.com/Hmbown/CodeWhale/pull/3048)**  
   Extends `apply_model_template()` with runtime capability lookups instead of hardcoded V4 claims. Partial fix for #3025.

10. **[#3053 — Add "Upgrading from deepseek-tui" Section to README](https://github.com/Hmbown/CodeWhale/pull/3053)**  
    Community documentation PR: covers npm, Cargo, Homebrew, and GitHub Releases upgrade paths with links to `REBRAND.md`.

---

## 5. Feature Request Trends

**1. Multi-Provider / Multi-Model Support (Dominant Theme)**  
The single biggest direction: un-hardcoding DeepSeek-specific assumptions. Requests include provider fallback chains (#2574), native Anthropic Messages API (#3014), universal auto-router (#3018), and parameterized constitution prompts (#3025). The v0.8.58 branch is explicitly targeting this.

**2. Remote Workbench / Headless Operation**  
Strong push for unattended operation outside China: DigitalOcean droplets (#2964), self-hosted Mac targets (#2968), Telegram bridge improvements (#2966), and enhanced `codewhale exec` scripting (#3027).

**3. UI/UX Polish**  
Recurring requests for compact rendering (#3031), clickable sidebar rows (#3028/2018/2889), session panel with auto-resume (#2934), and internal ID hiding (#3030).

**4. Hooks / Policy Control**  
Growing interest in model-agnostic deterministic control: hooks v2 JSON contract (#3026/PR#3049), per-run tool permissions (#3027), and `allowed-tools` flags.

---

## 6. Developer Pain Points

- **Rebranding migration friction**: Users on legacy `deepseek-tui` installs hit hard errors with no guidance (#2960). The "Upgrading" docs PR (#3053) is a direct response.
- **Provider reliability under load**: Windows SSE timeouts (#1679), 120s sub-agent API timeouts (#1806), and silent agent failure with false "completed" status (#2989) erode trust in parallel execution.
- **Config path fragmentation across OS**: Inconsistent resolution on Linux/macOS/Windows/Cygwin (#2369) plus lingering DeepSeek-branded paths after rebrand (#2664).
- **Missing fallback on provider failure**: No automatic switch when quota exhausts or API returns 429/5xx (#2574) — users must manually `/provider` mid-session.
- **UI blocking under subagent load**: Full terminal freeze on every progress event with 4+ agents (#3033/PR#3035). Telemetry bridge also deadlocks on approval dispatch (#2966).

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*