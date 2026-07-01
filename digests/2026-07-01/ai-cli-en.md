# AI CLI Tools Community Digest 2026-07-01

> Generated: 2026-07-01 11:36 UTC | Tools covered: 9

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

Here is the cross-tool comparison report based on the provided community digests.

---

## Cross-Tool Comparison Report: AI CLI Developer Tools
**Date:** 2026-07-01

### 1. Ecosystem Overview

The AI CLI tools ecosystem is characterized by rapid iteration and a fierce battle for developer workflow dominance. While **Claude Code** and **OpenAI Codex** remain the leading contenders by community size and feature breadth, a second tier of tools—**Gemini CLI**, **Copilot CLI**, and **OpenCode**—are aggressively closing gaps on specific features (e.g., hooks, MCP support) while differentiating on enterprise controls and model flexibility. The community across all platforms is laser-focused on three core themes: **safety and cost control** (preventing runaway agent loops and data loss), **cross-platform stability** (the Windows experience remains a universal pain point), and **workflow non-blocking** (message queuing, background task management). Newer entrants like **Pi** and **CodeWhale** are innovating on developer experience (e.g., AOT compilation, rebranding, constitution-first setup) but face maturity challenges with dependency management and platform-specific bugs.

### 2. Activity Comparison

| Tool | Open Issues (Active) | Closed Issues (24h) | Open PRs (Active) | Merged PRs (Recent) | Releases Today |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **Claude Code** | 4 | 6 | 1 | 1 | v2.1.197 |
| **OpenAI Codex** | 9 | 1 | 1 | 9 (recent days) | rust-v0.142.5, rust-v0.143.0-alpha.32 |
| **Gemini CLI** | 10 | 0 | 9 (merged) | 10 (recent) | v0.51.0-nightly.20260701 |
| **GitHub Copilot CLI** | 8 | 1 | 0 | 0 | v1.0.66, v1.0.67 |
| **Kimi Code CLI** | 1 | 1 | 1 | 0 | None |
| **OpenCode** | 7 | 4 | 5 (open) | 3 | v1.17.12 |
| **Pi** | 4 | 6 | 1 (open) | 9 | v0.80.3 |
| **Qwen Code** | 2 | 8 | 9 (open) | 0 | v0.19.3-nightly (failed preview) |
| **CodeWhale** | 9 | 1 | 8 (open) | 2 | v0.8.66 |

**Key Takeaway:** **OpenAI Codex** and **Pi** showed the highest PR merge velocity, indicating significant feature development. **Gemini CLI** and **Qwen Code** have large numbers of open PRs, suggesting a backlog of pending features. **Claude Code** remains the dominant player by sheer issue volume and community engagement, but this also indicates a higher number of user-reported defects.

### 3. Shared Feature Directions (Convergent Needs)

Several feature requests and pain points are appearing across multiple tool communities, signaling a strong developer demand for a more mature and integrated tooling layer.

- **Message Queue & Non-Blocking Interaction (Claude Code #50246, OpenCode #4821, OpenAI Codex #28268):** Users across tools want to queue prompts, compose messages, or cancel pending tasks without interrupting the currently executing agent turn. This is a fundamental UX requirement for practical, long-running agent workflows.
- **Configurable Tool Permissions & Safety Controls (Claude Code #72732/72733, Copilot CLI #179/#3028, Gemini CLI #22672, Qwen Code #6106):** The demand for granular, user-defined allow/deny lists for tools (shell, file write, web fetch) and sub-agents is urgent. Recent incidents of destructive `rm -rf` operations and runaway cost spirals ($600 sessions) have made this a top priority.
- **Multi-Account/Instance Management (Claude Code #27302, Copilot CLI #1665, Kimi Code #1938):** Developers want to manage multiple accounts (personal vs. work) or project-specific contexts without global configuration changes. This is essential for team environments.
- **Enhanced MCP and Plugin Ecosystem (Copilot CLI #3028, OpenAI Codex #26234, Pi #6211, CodeWhale #3866):** The community is pushing for deeper integration with the Model Context Protocol (MCP), including dynamic server spawning, namespace flattening for non-OpenAI providers, and richer plugin APIs that allow hooks at every lifecycle stage.
- **AST-Aware and Context-Efficient Operations (Gemini CLI #22745, Claude Code - community inference):** Reducing token usage by reading specific method boundaries or code segments (instead of entire files) is a recurring idea to improve cost efficiency and response speed.
- **Notification & Result Delivery (Kimi Code #1938, OpenCode - inferred):** There is a clear desire for push notifications for long-running tasks (e.g., web search results, code generation) so developers can remain productive without polling the terminal.

### 4. Differentiation Analysis

While the tools are converging on core features, their strategic focus and target users create clear differentiation:

- **Claude Code** is positioned as the **most powerful and feature-rich, but riskier**, tool. It leads in community features (hooks, sub-agents) but also suffers from the most severe safety and cost-control incidents due to its high agent autonomy. Its strength is its depth; its weakness is the potential for runaway behavior.
- **OpenAI Codex** is the **API-first infrastructure play**. It focuses on durable state, internal agent-to-agent communication (User Message Queue APIs), and a stable, scalable foundation. It is less visible in front-end UX but stronger on backend architecture. The "hook parity" demand shows it is catching up to Claude Code's agent lifecycle features.
- **GitHub Copilot CLI** is the **safety-and-integration-focused enterprise tool**. It emphasizes sandboxing (`/bypass`), session limits, and project-scoped configuration. Its community is smaller but highly focused on stability and predictable behavior, making it a strong choice for regulated environments.
- **Gemini CLI** is an **academic/research-oriented tool** experimenting with advanced concepts like AST-aware tooling and component-level evaluations. Its agent is powerful but unreliable (hangs, false success reporting), reflecting a "move fast and break things" development phase. It is the least production-ready of the major tools.
- **OpenCode** is the **open-source, community-driven alternative** with a strong focus on UI/UX polish (question tool fixes, session status). It is rapidly adopting features from Claude Code (hooks, MCP) while building a distinct identity in crypto payments and terminal customization.
- **Pi** and **CodeWhale** are the **agile, opinionated newcomers**. Pi excels at rapid feature delivery (AOT compilation, multiple model providers), while CodeWhale is innovating on "constitution-first" safety and dynamic tool spawning. Both suffer from dependency and platform stability issues typical of young projects.

### 5. Community Momentum & Maturity

- **Most Active / Highest Maturity:** **Claude Code** remains the bellwether. Its community is the largest and most vocal, driving feature requests with hundreds of upvotes. However, its "closed" issue count and frequent bug reports suggest a more "power-user, accept-rough-edges" posture.
- **High Momentum / Rapid Iteration:** **OpenAI Codex** and **Pi** are shipping features at the highest velocity. Codex is stabilizing after its massive SQLite log fix, while Pi is aggressively adding model providers and extension improvements. **Gemini CLI** is also iterating rapidly but with more frequent regressions.
- **Niche but Growing:** **Copilot CLI** and **Qwen Code** have smaller, more specialized communities. Copilot's focus on enterprise reliability is a strength, while Qwen Code's struggles with local model compatibility and release pipeline failures suggest it is further behind on operational maturity.
- **Fragile / Early Stage:** **CodeWhale** (formerly DeepSeek TUI) is in a transitional rebranding phase. While innovative, it has a high number of open Windows-specific bugs and a fragmented feature request landscape, indicating a community still finding its footing.

### 6. Trend Signals for Decision-Makers & Developers

- **The "Runaway Agent" Problem is the #1 Risk:** All major tools are grappling with uncontrolled cost spirals and destructive actions. **Developers should set strict budget limits and use sandboxed environments (e.g., Docker) when using high-autonomy agents.** Demand for `limit.output` and tool-level permissions is not a feature request; it is a safety requirement.
- **"Compose-First" UX is the Next Battleground:** The universal demand for message queuing and non-blocking interaction signals a shift from a synchronous REPL model to an asynchronous "IDE-for-agents" model. Tools that solve this (OpenAI Codex's Message Queue API is a leading candidate) will gain a significant workflow advantage.
- **Windows Parity is a Competitive Moat:** The aggregated data shows that Windows users consistently face a lower-quality experience (crashes, sandbox failures, terminal bugs). Any tool that invests heavily in a robust, first-class Windows experience will capture a large, underserved user base.
- **MCP is Winning, but Standardization is Needed:** MCP support is now table stakes across all tools. The next challenge is interoperability and provider-agnostic tool discovery (flattening namespaces), as highlighted by the OpenAI Codex community.
- **Self-Hosting & Local Models are a Growing but Painful Niche:** Issues with Ollama, LM Studio, and other local providers (common in Qwen Code, Gemini CLI) show a strong desire for offline/private AI coding. However, the tool support is still inconsistent and requires significant user effort to debug.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
*Data as of 2026-07-01 | Source: github.com/anthropics/skills*

---

## 1. Top Skills Ranking

The following PRs represent the most-discussed new Skills by community engagement. All are currently **open**; none have been merged as of the data snapshot.

**1. document-typography (#514)**
- *Functionality:* Prevents orphan words, widow paragraphs, and numbering misalignment in AI-generated documents. Targets the pervasive typographic flaws Claude introduces.
- *Discussion highlights:* Community consensus that these issues affect every document; users rarely request fixes explicitly but benefit from automatic correction. Low friction, high value.
- *Status:* Open.  [GitHub: PR #514](https://github.com/anthropics/skills/pull/514)

**2. ODT (#486)**
- *Functionality:* Creates, fills, reads, and converts OpenDocument Format files (.odt, .ods). Includes template filling and ODT-to-HTML conversion. Triggers on "ODT", "ODS", "OpenDocument", "LibreOffice".
- *Discussion highlights:* Addresses the ISO-standard office format gap in the skills collection. Users noted interoperability with LibreOffice and Microsoft Word via conversion.
- *Status:* Open.  [GitHub: PR #486](https://github.com/anthropics/skills/pull/486)

**3. frontend-design (#210)**
- *Functionality:* Revised skill to improve clarity and actionability. Ensures every instruction is executable within a single conversation and steers behavior without over-constraining creativity.
- *Discussion highlights:* Strong focus on token efficiency and removing verbose explanations that wasted context. Many users reporting the original skill underperformed.
- *Status:* Open.  [GitHub: PR #210](https://github.com/anthropics/skills/pull/210)

**4. skill-quality-analyzer & skill-security-analyzer (#83)**
- *Functionality:* Meta-skills for evaluating other Skills across five quality dimensions (structure, documentation, example coverage, edge case handling, cross-platform) and two security dimensions (trust boundary, permission scope).
- *Discussion highlights:* Spawned the security/namespace debate (see Issue #492). Community split on whether meta-skills should live alongside regular skills or in a separate audit toolchain.
- *Status:* Open.  [GitHub: PR #83](https://github.com/anthropics/skills/pull/83)

**5. SAP-RPT-1-OSS (#181)**
- *Functionality:* Wraps SAP's open source tabular foundation model for predictive analytics on SAP business data. Handles data ingestion, feature engineering, and model inference.
- *Discussion highlights:* Enterprise-heavy interest. Users asked about model size constraints and whether the skill works with on-premise SAP instances. No maintainer response yet.
- *Status:* Open.  [GitHub: PR #181](https://github.com/anthropics/skills/pull/181)

**6. self-audit (#1367)**
- *Functionality:* A universal reasoning quality gate that checks AI output across four dimensions (completeness, consistency, grounding, actionability) before delivery. Model-agnostic.
- *Discussion highlights:* Very recent (submitted 2026-06-28). Early commenters see it as a lightweight alternative to full agent-governance systems. Potential for integration with the skill-creator pipeline.
- *Status:* Open.  [GitHub: PR #1367](https://github.com/anthropics/skills/pull/1367)

**7. testing-patterns (#723)**
- *Functionality:* Covers the full testing stack: Testing Trophy philosophy, AAA pattern, React Testing Library, integration testing, end-to-end flows, and CI integration.
- *Discussion highlights:* Community requested a "what NOT to test" section, which was added. Some concern about token length (skill is ~400 lines). Widely viewed as addressing a critical workflow gap.
- *Status:* Open.  [GitHub: PR #723](https://github.com/anthropics/skills/pull/723)

**8. sensory (#806)**
- *Functionality:* Native macOS automation via `osascript` (AppleScript). Two-tier permission system: Tier 1 works out of the box (direct app scripting), Tier 2 requires Accessibility permissions (System Events UI control).
- *Discussion highlights:* Strong platform-specific demand. macOS users see it as an alternative to screenshot-based computer use. Privacy/permission model praised.
- *Status:* Open.  [GitHub: PR #806](https://github.com/anthropics/skills/pull/806)

---

## 2. Community Demand Trends

From the top 13 Issues by comment volume, four concentrated demand directions emerge:

**a) Trust and Security Boundaries (Issue #492, 32 comments)**
The most commented issue. Community skills distributed under the `anthropic/` namespace create a trust abuse vulnerability. Users may grant elevated permissions to skills they believe are official. No resolution path yet; the community is split between namespace separation (e.g., `community/` prefix) and mandatory security audit gates.

**b) Organizational Skill Sharing (Issue #228, 14 comments)**
Seven thumbs-up (highest in the dataset). Users want org-wide skill libraries with direct sharing links rather than file-based Slack distribution. This is a platform feature request, not a skill content request—but it dominates discussion because the lack of sharing infrastructure limits adoption.

**c) Skill-Creation Pipeline Reliability (Issues #556, #1169, #1061, 18 combined comments)**
The `run_eval.py` tool reports 0% recall on every query due to subprocess handling bugs, particularly on Windows. Three separate Issues confirm the same failure mode. The community is spending significant effort debugging the evaluation framework rather than creating new skills.

**d) New Skill Types: Agent Governance & Compact Memory (Issues #412, #1329, 13 combined comments)**
Users propose two distinct directions: (1) agent-governance—policy enforcement, threat detection, audit trails for AI agent systems; (2) compact-memory—symbolic notation for reducing context consumed by agent state prose. Both are "meta-skills" that modify how Claude operates, not what it produces.

---

## 3. High-Potential Pending Skills

These PRs have active discussion threads and appear close to landing. All are open as of data collection.

- **#514 – document-typography:** Minimal merge controversy; low-risk, high-value fix. Pending final maintainer review.
  [PR #514](https://github.com/anthropics/skills/pull/514)

- **#486 – ODT:** Only blocker is whether to include strict conformance tests for LibreOffice vs. Microsoft ODT rendering differences.
  [PR #486](https://github.com/anthropics/skills/pull/486)

- **#1367 – self-audit:** Very fresh (submitted June 28) but already generating traction as a lightweight alternative to full governance frameworks.
  [PR #1367](https://github.com/anthropics/skills/pull/1367)

- **#723 – testing-patterns:** Token length was the main concern; a trimmed version may merge soon.
  [PR #723](https://github.com/anthropics/skills/pull/723)

- **#806 – sensory:** Apple-specific but cleanly architected. Permission model was refined after community feedback. No maintainer objection flagged.
  [PR #806](https://github.com/anthropics/skills/pull/806)

- **#147 – codebase-inventory-audit:** 10-step orphan code detection workflow. Community requested better integration with existing linter output. Author responsive.
  [PR #147](https://github.com/anthropics/skills/pull/147)

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is for a reliable, cross-platform skill-creation evaluation pipeline that actually measures skill triggering correctly—followed by trust-safe distribution mechanisms that prevent namespace impersonation and enable organizational sharing.**

---

# Claude Code Community Digest – 2026-07-01

## Today's Highlights

Claude Code v2.1.197 ships with **Claude Sonnet 5** as the new default model, boasting a native 1M-token context window and promotional pricing ($2/$10 per Mtok) through August 31. The community remains highly engaged on two major feature requests – multi-connector account support (296 👍) and a message queue mode (119 👍) – while a small wave of critical bug reports around recursive agent cost spirals and tool execution data loss signals growing pains as users push agents to their limits.

## Releases

- **v2.1.197** – [Release notes](https://github.com/anthropics/claude-code/releases/tag/v2.1.197)  
  Introduces Claude Sonnet 5 as the default model, with a native 1M-token context window and promotional pricing through August 31. Update to this version to access the new model.

## Hot Issues

1. **\[OPEN\] [FEATURE] Support multiple Connector accounts** – [Issue #27302](https://github.com/anthropics/claude-code/issues/27302)  
   Highly requested (296 👍, 204 comments) feature to allow multiple accounts per connector in both Claude AI and Claude Code. A top community priority for over four months.

2. **\[OPEN\] Feature Request: Message queue mode** – [Issue #50246](https://github.com/anthropics/claude-code/issues/50246)  
   Proposes queuing messages while Claude is busy instead of interrupting active tasks. 119 👍, 35 comments – indicates strong desire for non-blocking interaction.

3. **\[CLOSED\] Windows 11 Segmentation Fault (v2.1.112+)** – [Issue #50640](https://github.com/anthropics/claude-code/issues/50640)  
   Crash on startup affecting Windows users. Fixed in later releases; closed as stale after fix verification.

4. **\[CLOSED\] CoworkVMService DACL blocks installer upgrades** – [Issue #57035](https://github.com/anthropics/claude-code/issues/57035)  
   Root cause for multiple Windows installation failures (e.g., #25136, #48367, #49655). Resolved.

5. **\[CLOSED] Chrome extension CSP blocks MCP bridge** – [Issue #62002](https://github.com/anthropics/claude-code/issues/62002)  
   Inline script blocked by Content Security Policy prevented side panel and MCP browser bridge from working. Fixed.

6. **\[CLOSED] Destructive `rm -rf` on explicitly excluded directory** – [Issue #62402](https://github.com/anthropics/claude-code/issues/62402)  
   Sonnet 4.6 running a cleanup session removed files in a user-excluded directory. Data-loss bug that raised safety concerns. Closed as stale.

7. **\[CLOSED] Opus 4.8: Tool call parsing fails with extended thinking** – [Issue #63481](https://github.com/anthropics/claude-code/issues/63481)  
   Multiple reports (8 👍) that Opus 4.8 fails with parse errors when extended thinking is triggered. Closed as duplicate but no public fix noted.

8. **\[CLOSED] Regression: /remote-control missing in v2.1.196** – [Issue #72424](https://github.com/anthropics/claude-code/issues/72424)  
   The `/remote-control` slash command was removed or broken in the previous latest release. Quick response likely needed.

9. **\[OPEN] Controllable recursive agent spawning spirals costs** – [Issue #72732](https://github.com/anthropics/claude-code/issues/72732)  
   User reports uncontrolled subagent spawning burned $600 in a session. Currently open with 0 comments – likely to become a major concern.

10. **\[OPEN] Tool execution deletes files despite "replace only" instruction** – [Issue #72733](https://github.com/anthropics/claude-code/issues/72733)  
    Another data-loss report where Claude Code deleted all work. Author mentions legal action; needs immediate attention.

## Key PR Progress

Only two pull requests were updated in the last 24 hours, reflecting a quieter period for collaborative development:

- **\[OPEN] Create Cha** – [PR #72543](https://github.com/anthropics/claude-code/pull/72543)  
  Title is ambiguous; likely a new feature or placeholder. No details available.

- **\[CLOSED] docs(plugin-dev): add plugin cache sync guidance** – [PR #46903](https://github.com/anthropics/claude-code/pull/46903)  
  Documents that local plugin source changes are not automatically synced to `~/.claude/plugins/cache/`. Essential for plugin developers who edit locally.

## Feature Request Trends

- **Multi-account / multi-connector support** (#27302) remains the single most-upvoted feature, with no sign of waning demand.
- **Message queue / compose-while-working** (#50246, #62856) is a recurring theme: users want to queue follow-ups or compose messages without interrupting an active task.
- **Desktop recents enumeration** (#54911) – a minor but practical request to list recent projects from `~/.claude/projects/` directly in the desktop app.
- **Improved local plugin development workflow** (#46903 PR) – users want hot-reload or explicit sync for locally developed plugins.

## Developer Pain Points

- **Thinking block mutation errors**: A cluster of closed bugs (#63258, #63278, #63247, #63448, #63508) all point to Opus 4.7/4.8 and interleaved thinking causing 400 "cannot modify thinking blocks" errors, especially with subagents or background tasks. A systemic issue that appears partially addressed but surfaces across multiple configurations.
- **Windows platform fragility**: Multiple Windows bugs (crashes, no bash output, VS Code path issues, CoworkVMService DACL) indicate the Windows experience lags behind macOS/Linux in stability.
- **Uncontrolled cost and data loss**: Two brand-new open issues (#72732, #72733) highlight dangerous agent behavior – runaway subagent spawning costing hundreds of dollars, and file deletion despite explicit instructions. These signal that safety and cost-control guardrails need strengthening.
- **VS Code extension integration bugs**: Issues around thinking display, terminal setup paths, and invalid API messages suggest the VS Code extension has specific regression points that differ from the CLI experience.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-07-01

## Today’s Highlights
Two patch releases landed today, with `rust-v0.142.5` fixing a sensitive trace-logging issue and `rust-v0.143.0-alpha.32` moving the alpha channel forward. Community discussion remains focused on persistent SQLite log churn (even after the 0.142.0 fix), an alarming number of Windows‑exclusive bugs, and a growing demand for full hook parity with Claude Code. The automations feature stack (nine PRs) was merged this week, signalling a major new capability in the pipeline.

## Releases
- **[rust-v0.142.5](https://github.com/openai/codex/releases/tag/rust-v0.142.5)** — Bug fix: prevents full Responses WebSocket request payloads from being written to trace logs (PR [#30771](https://github.com/openai/codex/pull/30771)).  
- **[rust-v0.143.0-alpha.32](https://github.com/openai/codex/releases/tag/rust-v0.143.0-alpha.32)** — Alpha release; no detailed changelog provided.

## Hot Issues
1. **[#28224 – SQLite feedback logs can write ~640 TB/year](https://github.com/openai/codex/issues/28224)** (119 comments, 👍410)  
   *Closed after three merged PRs reduced logs by ~85%.* Still the most upvoted issue ever – a community win.

2. **[#29532 – macOS: Persistent SQLite TRACE log churn remains after 0.142.0](https://github.com/openai/codex/issues/29532)** (31 comments, 👍7)  
   *Partial fix: `responses_websocket` improved, but `log_2.sqlite` still grows. Users want full eradication.*

3. **[#29320 – Codex app only displays “Something went wrong…” on Windows](https://github.com/openai/codex/issues/29320)** (27 comments, 👍2)  
   *Frequent crash on Windows 11 after update. Critical for Windows users.*

4. **[#21753 – Full Claude Code Hook Parity (29+)](https://github.com/openai/codex/issues/21753)** (25 comments, 👍19)  
   *Umbrella tracker for hook coverage – every major lifecycle event. High community desire for automation.*

5. **[#26234 – Flatten MCP namespace tools for non-OpenAI providers](https://github.com/openai/codex/issues/26234)** (23 comments, 👍35)  
   *Ollama, LM Studio, OpenRouter users cannot call MCP tools because of proprietary wrapper. Active PR #29602.*

6. **[#28823 – 5-hour usage meter consumes much faster than historical](https://github.com/openai/codex/issues/28823)** (19 comments, 👍2)  
   *Regression in quota accounting – multiple reports of meter draining faster than actual usage.*

7. **[#30212 – Usage limit depleted abnormally in ~1 hour](https://github.com/openai/codex/issues/30212)** (7 comments, 👍9)  
   *Another rate‑limit anomaly; even Pro subscribers seeing 20× allowance burned too quickly.*

8. **[#25271 – Computer Use cannot determine Chrome URL on Windows](https://github.com/openai/codex/issues/25271)** (13 comments, 👍6)  
   *Blocks browser automation on Windows – a key feature for many.*

9. **[#20570 – Windows sandbox: CreateProcessAsUserW failed 1920](https://github.com/openai/codex/issues/20570)** (9 comments, 👍9)  
   *Recurring sandbox runner error after upgrades – forces manual workaround each time.*

10. **[#30689 – Unexpectedly high usage consumption after context compaction](https://github.com/openai/codex/issues/30689)** (5 comments, 👍0)  
    *New report linking context compaction to abnormal rate limit depletion.*

## Key PR Progress
1. **[#29602 – Flatten namespace tools for providers without wrappers](https://github.com/openai/codex/pull/29602)** (OPEN)  
   *Fixes #26234 by allowing non‑OpenAI providers to discover/register namespaced MCP tools.*

2. **[#28409 – Enforce exact managed config values](https://github.com/openai/codex/pull/28409)** (CLOSED)  
   *Adds exact‑value enforcement for SQLite home, log dir, model catalog, etc. Warns on startup.*

3. **[#28602 – Force offline standalone web search after connector use](https://github.com/openai/codex/pull/28602)** (CLOSED)  
   *Sticky per‑thread policy to switch Codex to offline web search once a connector (MCP) is used.*

4. **[#28268 – Expose the User Message Queue app-server API](https://github.com/openai/codex/pull/28268)** (CLOSED)  
   *New experimental API for durable thread‑scoped message queuing.*

5. **[#22722 – Persist thread artifacts in app-server runtime](https://github.com/openai/codex/pull/22722)** (CLOSED)  
   *Generic artifact storage with SQLite caching and notifications – foundational for stateful threads.*

6. **[#28594 – Fail thread forks on malformed source rollouts](https://github.com/openai/codex/pull/28594)** (CLOSED)  
   *Prevents silent data loss when forking threads with corrupt JSONL history.*

7. **[#26259 – Add advisory Interrupt hooks for interrupted turns](https://github.com/openai/codex/pull/26259)** (CLOSED)  
   *New `on_interrupt` hook – advisory only, cannot block, but gives scripts a cleanup opportunity.*

8. **[#28573 – Allow nonblocking environment resolution checks](https://github.com/openai/codex/pull/28573)** (CLOSED)  
   *Unblocks turn startup by allowing inspection of environment readiness without blocking.*

9. **[#28456 – Reduce resume and fork orchestration overhead](https://github.com/openai/codex/pull/28456)** (CLOSED)  
   *Optimizes startup by reusing cached thread metadata and avoiding redundant history loads.*

10. **[#28455 – Repair stale and custom rollout paths](https://github.com/openai/codex/pull/28455)** (CLOSED)  
    *Fixes thread recovery when rollout paths are moved or archived – reduces “thread not found” errors.*

## Feature Request Trends
- **Hook parity with Claude Code** (#21753, #26259): Community wants hook support for every major lifecycle event – interrupt, tool call, turn start/stop – to enable robust automation.
- **Plugin management from CLI** (#17431): Users request a `codex plugins` subcommand to list, install, and upgrade plugins without the TUI.
- **MCP namespace flattening** (#26234): Non‑OpenAI provider users are blocked by the proprietary `type: "namespace"` wrapper. PR #29602 is in progress.
- **Thread artifact persistence** (#22722, #14251): Several requests for durable thread state (artifact caching, session recovery) are being addressed in the app‑server runtime.
- **Rate‑limit transparency** (#28823, #30212, #30689): Users want clear telemetry and a fix for apparently inflated consumption after context compaction.

## Developer Pain Points
- **Windows‑exclusive bugs dominate**: Over half of the top issues affect Windows – sandbox failures, kernel‑pool memory leaks, Chrome plugin registration problems, duplicate MCP pools, and “Something went wrong” crashes. Each update seems to bring fresh regressions.
- **SQLite log churn despite fixes**: Issues #28224 and #29532 show that even after reducing 85% of logs, macOS users still see persistent `log_2.sqlite` growth, consuming SSD endurance.
- **Rate‑limit accounting inconsistencies**: Multiple reports (#28823, #30212, #30689) describe 5‑hour meters draining 2–5× faster than historical usage, with users unable to distinguish genuine use from a quota bug.
- **Stuck threads and session recovery**: Several issues (#14251, #30748, #19558) document threads that get stuck in “generating” or fail to recover after interruption, forcing manual log editing or thread abandonment. PRs #28455 and #28594 aim to address these cases.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-07-01

## Today's Highlights
The team shipped a nightly release (v0.51.0-nightly.20260701) with defensive path resolution for `@`-prefixed file references and a new Cloud Run webhook ingestion service for the Caretaker agent. Several long-standing agent reliability issues remain under active investigation, particularly around subagent turn limits and shell command hangs, while the community continues to push for AST-aware tooling and better agent self-awareness.

## Releases
**v0.51.0-nightly.20260701.g7f00c5fe5**  
- Fix: Defensive path resolution for `@`-referenced files in core-tools (PR [#28053](https://github.com/google-gemini/gemini-cli/pull/28053)).
- Feature: Caretaker Cloud Run webhook ingestion service (PR [#28015](https://github.com/google-gemini/gemini-cli/pull/28015) – merged in previous cycle).

## Hot Issues (Top 10 by Community Activity)

1. **#22323** — Subagent recovery after `MAX_TURNS` falsely reports success as "GOAL"  
   *Why it matters:* Misleading termination reason hides true interruptions; users cannot distinguish between successful completion and turn-limit halts.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/22323) | 9 comments | 👍 2

2. **#24353** — Epic: Robust component-level evaluations  
   *Why it matters:* Follow-on from behavioral evals; aims to run 76+ tests across 6 Gemini models—critical for regressions.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/24353) | 7 comments

3. **#22745** — Assess impact of AST-aware file reads, search, and mapping  
   *Why it matters:* Could drastically reduce token usage and turn count by reading precise method bounds instead of whole files.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/22745) | 7 comments | 👍 1

4. **#21409** — Generalist agent hangs forever when deferring to subagents  
   *Why it matters:* Most upvoted bug (👍8). Users work around by disabling subagents entirely—core reliability issue.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/21409) | 7 comments

5. **#21968** — Gemini does not use custom skills and sub-agents unless explicitly instructed  
   *Why it matters:* Undermines the value of custom skills; agent ignores descriptions when tasks are closely related.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/21968) | 6 comments

6. **#26525** — Add deterministic redaction and reduce Auto Memory logging  
   *Why it matters:* Secrets sent to model before redaction; Auto Memory logging may expose skill content—security concern.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/26525) | 5 comments

7. **#26522** — Stop Auto Memory from retrying low-signal sessions indefinitely  
   *Why it matters:* Wastes API quota; sessions without useful content remain unprocessed and keep being re-evaluated.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/26522) | 5 comments

8. **#25166** — Shell command execution gets stuck “Waiting input” after command completes  
   *Why it matters:* Blocks workflow for simple commands (e.g., `ls`, `git status`); high frustration (👍3).  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/25166) | 4 comments

9. **#21983** — Browser subagent fails on Wayland  
   *Why it matters:* Linux Wayland users cannot use browser automation; termination reason wrongly reports “GOAL”.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/21983) | 4 comments

10. **#22672** — Agent should stop/discourage destructive behavior (e.g., `git reset --force`)  
    *Why it matters:* Safety concern; model sometimes uses dangerous commands when safer alternatives exist.  
    [Issue](https://github.com/google-gemini/gemini-cli/issues/22672) | 3 comments

## Key PR Progress (10 Important Pull Requests)

1. **#28053** — Defensive path resolution for `@`-referenced files + macOS test fixes (merged)  
   *Fix:* Prevents “File not found” errors when model passes paths like `@policies/new-policies.txt`.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28053)

2. **#27971** — Strip thoughts from scrubbed history turns to prevent thought leakage (merged)  
   *Fix:* Eliminates cycles where model mimics its own scratchpad after reading internal thoughts in history.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/27971)

3. **#28094** — Deep-merge user and workspace settings in A2A server  
   *Fix:* Shallow spread previously lost nested configs (tools, telemetry, etc.) from workspace settings.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28094)

4. **#28103** — Avoid keep-alive socket reuse during OAuth token exchange  
   *Fix:* Resolves “Premature close” error on Node.js versions patched for CVE-2026-48931.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28103)

5. **#28112** — Add SSRF protection to MCP OAuth metadata discovery  
   *Fix:* Prevents SSRF attacks when fetching URLs from MCP server responses (gap vs. existing `web-fetch.ts`).  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28112)

6. **#28163** — Core foundation for Caretaker triage worker (Part 1/2)  
   *Feature:* Introduces Cloud Run triage worker modules; second PR to follow with full logic.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28163)

7. **#28223** — Bypass LLM correction for JSON and IPYNB files in write_file/replace  
   *Fix:* Corrupted Jupyter Notebooks and `.json` files when model attempted to “correct” them—now skips rewriting.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28223)

8. **#28126** — Show ellipsis on multi-line edit snippets  
   *Fix:* `EditToolInvocation.getDescription()` now appends `...` for multi-line or long first-line edits, improving UI clarity.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28126)

9. **#28224** — Avoid splitting emoji when truncating display strings  
   *Fix:* `sanitizeForDisplay` now uses proper Unicode-aware truncation to prevent broken renderings.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28224)

10. **#28164** — Limit recursive reasoning turns per single user request (15-turn cap)  
    *Fix:* Protects CPU/API credits from infinite loops; customizable via `maxSessionTurns`.  
    [PR](https://github.com/google-gemini/gemini-cli/pull/28164)

## Feature Request Trends

- **AST-aware tooling** (Issues [#22745](https://github.com/google-gemini/gemini-cli/issues/22745), [#22746](https://github.com/google-gemini/gemini-cli/issues/22746)): Users and maintainers alike want the CLI to read method bounds and map codebases using abstract syntax trees, reducing token waste and improving edit accuracy.
- **Component-level evaluations** (Issue [#24353](https://github.com/google-gemini/gemini-cli/issues/24353)): Scaling from 76 to hundreds of behavioral tests across models is a high priority for quality assurance.
- **Agent self-awareness** (Issues [#21432](https://github.com/google-gemini/gemini-cli/issues/21432), [#22598](https://github.com/google-gemini/gemini-cli/issues/22598)): The community wants the agent to know its own flags, hotkeys, and subagent trajectories and to share subagent context via `/chat share`.
- **Resilient browser agent** (Issues [#22232](https://github.com/google-gemini/gemini-cli/issues/22232), [#21983](https://github.com/google-gemini/gemini-cli/issues/21983)): Better lock recovery and Wayland support are frequently requested.
- **Native file tools for task tracking** (Issue [#21000](https://github.com/google-gemini/gemini-cli/issues/21000)): Experiment with using file-based tools instead of LLM-driven approaches for maintaining task lists.

## Developer Pain Points

- **Agent hangs & misleading terminations** (Issues [#21409](https://github.com/google-gemini/gemini-cli/issues/21409), [#22323](https://github.com/google-gemini/gemini-cli/issues/22323)): Subagents either hang indefinitely or falsely report success after hitting limits, breaking trust.
- **Subagent underutilization** (Issue [#21968](https://github.com/google-gemini/gemini-cli/issues/21968)): Custom skills and subagents are ignored unless explicitly prompted, negating their purpose.
- **Shell command hangs** (Issue [#25166](https://github.com/google-gemini/gemini-cli/issues/25166)): Simple commands like `git status` trigger “Waiting input” even after completion.
- **Memory system inefficiencies** (Issues [#26522](https://github.com/google-gemini/gemini-cli/issues/26522), [#26523](https://github.com/google-gemini/gemini-cli/issues/26523)): Auto Memory retries low-signal sessions endlessly and silently skips invalid patches, wasting API calls.
- **Security & false positives** (Issues [#26525](https://github.com/google-gemini/gemini-cli/issues/26525), [#28230](https://github.com/google-gemini/gemini-cli/issues/28230)): Secrets leak before redaction; some JS files trigger antivirus alarms (Kaspersky).
- **Destructive command execution** (Issue [#22672](https://github.com/google-gemini/gemini-cli/issues/22672)): The model occasionally runs `git reset --force` or other dangerous operations without warning.
- **Configuration overrides ignored** (Issue [#22267](https://github.com/google-gemini/gemini-cli/issues/22267)): Browser agent fails to respect `settings.json` overrides for `maxTurns` and other parameters.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-07-01

## Today's Highlights
Two patch releases landed yesterday (v1.0.66 and v1.0.67), addressing sandbox UX, cursor behavior, and host-agent error handling. The community continues to push for project-scoped plugin configuration (#1665) and globally configurable tool permissions (#179), while a persistent authentication bug (#2684) remains the most active open issue after two months.

## Releases
**v1.0.67** (2026-06-30) — [Release Link](https://github.com/github/copilot-cli/releases/tag/v1.0.67)
- Sandbox bypass via `/bypass` now takes effect immediately for shell/search commands mid-turn.
- Subagent sessions inherit parent agent tool restrictions.
- Warnings and errors are surfaced when host custom agents fail to load.
- Session limit enforcement added.

**v1.0.66** (2026-06-30) — [Release Link](https://github.com/github/copilot-cli/releases/tag/v1.0.66)
- Block cursor during interactive sessions, restoring terminal default on exit.
- Added support for **Claude Opus 4.8 Fast**; deprecated Claude Opus 4.6 Fast.
- MCP add/edit form now accepts HTTP-style `Key: value` headers.
- Fixed duplicate LSP server startup.

## Hot Issues
1. **#2684** — `[area:authentication, area:mcp]` Persistent "Authorization error, you may need to run /login" even after successful login. 13 comments, 0 reactions. [Issue](https://github.com/github/copilot-cli/issues/2684)
2. **#1665** — `[area:plugins, area:configuration]` Request for project/repo-scoped plugins instead of per-user global loading. 18 👍, 10 comments. [Issue](https://github.com/github/copilot-cli/issues/1665)
3. **#2334** (CLOSED) — `[area:configuration, area:terminal-rendering]` Calls to restore no-alt-screen mode, citing usability issues with scrollbar and text selection. 29 👍. [Issue](https://github.com/github/copilot-cli/issues/2334)
4. **#179** — `[area:permissions, area:configuration]` Global tool allowlist in config.json (similar to Claude Code). 41 👍, strong community interest. [Issue](https://github.com/github/copilot-cli/issues/179)
5. **#3727** — `[area:context-memory, area:plugins]` Regression in v1.0.60: `userPromptSubmitted` hook `additionalContext` no longer injected into planner. 6 comments. [Issue](https://github.com/github/copilot-cli/issues/3727)
6. **#3028** — `[area:permissions, area:mcp]` Request for configurable tool allow/deny rules per MCP server. 5 👍. [Issue](https://github.com/github/copilot-cli/issues/3028)
7. **#3948** — `[area:networking, area:tools]` `web_fetch` always fails with `TypeError: fetch failed` despite correct proxy/env. 3 comments. [Issue](https://github.com/github/copilot-cli/issues/3948)
8. **#3982** — `[area:authentication, area:mcp]` Copilot CLI ignores `grant_types_supported` and attempts authorization_code flow for client_credentials-only MCP servers. [Issue](https://github.com/github/copilot-cli/issues/3982)
9. **#3282** — `[area:models, area:configuration]` Request for multiple BYOK model support in CLI (currently limited to single env var). 11 👍. [Issue](https://github.com/github/copilot-cli/issues/3282)
10. **#1504** — `[area:theming-accessibility]` Custom theme support via JSON files. 20 👍. [Issue](https://github.com/github/copilot-cli/issues/1504)

## Key PR Progress
No significant pull requests were updated in the last 24 hours. The two open PRs (#3873, #3880) appear to be test or spam submissions without substantive changes. Activity on the PR front remains minimal.

## Feature Request Trends
- **Project-scoped plugins** (#1665, 18 👍) — Users want per-repo `.copilot/plugins` loading rather than global-only installation.
- **Globally configurable tool permissions** (#179, 41 👍) — A persistent request for a settings.json-based allowlist (mirroring Claude Code's approach).
- **Multi-model BYOK support** (#3282, 11 👍) — Users want the ability to register multiple bring-your-own-key models and switch between them without restarting.
- **Custom themes** (#1504, 20 👍) — Demand for user-defined, shareable theme JSON beyond the built-in options.
- **Per-mode default models** (#2958, 14 👍) — Allow separate default model configurations for plan mode vs. autopilot mode.

## Developer Pain Points
- **Authentication flakiness** — Issue #2684 remains the most active open bug, with users repeatedly re-authenticating but still hitting authorization errors. No resolution in sight after 2.5 months.
- **Plugin hook regression** — The v1.0.60 regression (#3727) broke the `userPromptSubmitted` hook's ability to inject context into the planner, disrupting plugin workflows that worked in v1.0.59.
- **Cursor and clipboard UX** — Multiple issues (#2507, #3981, #3993) report terminal cursor not respecting system defaults (Windows), clipboard copy breaking while CLI is running (Windows), and screen-reader non-echo of typed input.
- **MCP OAuth compatibility** — Issue #3982 highlights that the CLI ignores OAuth server capabilities, forcing interactive auth flows on servers that only support `client_credentials`.
- **web_fetch failure** — Issue #3948 points to a systemic `fetch failed` error when using the built-in web fetch tool, with no proxy workaround despite correct environment setup.

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

好的，这是基于您提供的 GitHub 数据生成的 Kimi Code CLI 社区摘要。

---

### Kimi Code CLI 社区摘要 | 2026-07-01

**1. Today's Highlights**

今日社区主要关注一个持续半年的棘手 Bug，该问题导致 CLI 在读取特定文件时陷入死循环，且至今仍未解决，引发社区对第三方模型及配置兼容性的讨论。同时，一个关于为 Web 界面增加推送通知功能的增强请求被关闭，虽然最终未被采纳，但反映出用户对异步任务通知功能的强烈渴望。在代码贡献方面，一项针对 Windows 终端粘贴媒体文件失败的修复 PR 已被提交，旨在解决 Windows 和 VS Code 集成终端中的关键体验问题。

**2. Releases**

无

**3. Hot Issues**

*   **[#640] [Bug] Kimi CLI stuck in reading one file again and again and stuck in a loop**
    *   **重要性**: 这是一个从2026年1月持续至今的严重 Bug，直接导致 CLI 无法正常使用。用户反馈程序陷入无限读取单个文件的死循环，使用体验极差。
    *   **社区反应**: 该问题持续被关注，已有15条评论，表明多位用户都曾遭遇类似问题。社区可能正在等待核心团队的修复方案。
    *   **链接**: [MoonshotAI/kimi-cli Issue #640](https://github.com/MoonshotAI/kimi-cli/issues/640)

*   **[#1938] [Enhancement] 是否可以为Kimi-CLI-Web增加推送功能?**
    *   **重要性**: 该 Issue 请求为 Kimi CLI 的 Web 版本和终端版本增加任务完成通知功能。这直接关系到用户在异步任务场景（如长文本生成）下的工作效率，是提升用户体验的关键功能。
    *   **社区反应**: 该 Issue 已被关闭，但获得了一次讨论。关闭原因可能是优先级低、架构不支持或已纳入其他计划，但其需求的提出表明用户对“被动接收结果”的强需求。
    *   **链接**: [MoonshotAI/kimi-cli Issue #1938](https://github.com/MoonshotAI/kimi-cli/issues/1938)

*(由于数据限制，此处仅提取到2个符合条件的活跃 Issue。实际每日摘要中通常会列出10个，其余将基于社区反馈的常见主题进行推断。)*

**4. Key PR Progress**

*   **[#2481] fix(shell): read clipboard media on BracketedPaste for Windows terminals**
    *   **描述**: 此 PR 修复了 Windows Terminal 和 VS Code 集成终端中，使用 `Ctrl+V` 粘贴图片等二进制内容失败的问题。它通过让 `_handle_bracketed_paste()` 函数在接到粘贴事件时主动尝试读取剪贴板内容，绕过了终端无法通过文本传递二进制数据的限制。
    *   **影响**: 直接解决了 Windows 用户在 CLI 中通过粘贴方式上传图片等媒体文件的关键体验问题，使功能行为与 macOS / Linux 更一致。
    *   **链接**: [MoonshotAI/kimi-cli PR #2481](https://github.com/MoonshotAI/kimi-cli/pull/2481)

*(由于数据限制，此处仅提取到1个符合条件的活跃 PR。实际每日摘要中通常会列出10个，其余将基于不同模块的修复和新功能进行推断。)*

**5. Feature Request Trends**

基于今日及过往 Issue 分析，用户最期待的功能方向主要集中在：

1.  **通知与推送机制**: 用户在 [Issue #1938] 中的需求表明，社区强烈希望在 CLI 终端和 Web 界面中增加异步任务完成的通知能力，无论用户当前处于哪个界面（如手机、桌面），都能即时获得反馈。
2.  **文件处理增强**: 尽管 [Issue #640] 是 Bug，但其“在读取文件时卡死”的现象也间接反映了用户对复杂、长文档或特定格式文件处理的稳定性和鲁棒性有更高要求。
3.  **跨平台体验统一**: [PR #2481] 的提交突显了社区对 Windows 平台体验优化的关注，用户期望所有主流操作系统（macOS, Linux, Windows）的核心功能体验保持一致。

**6. Developer Pain Points**

*   **无限循环与程序挂起**: 如 [Issue #640] 所示，开发者在使用 CLI 处理文件时遭遇了程序无限循环的严重 Bug，导致工作完全中断且无法自行恢复。
*   **第三方模型/配置兼容性**: 该 Bug 报告发生在用户使用自定义的 Anthropic 端点（即第三方模型）时，暴露了 CLI 在非标准 API 配置下的潜在兼容性和稳定性问题，增加了使用非官方模型的开发者风险。
*   **缺失的关键通知机制**: 由于 CLI 缺少任务完成推送，开发者在使用如 Kimi CLI Web 等界面进行长时间推理时，必须手动轮询结果，无法复用其他任务，这成为了一个明确的生产力瓶颈。

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-07-01

## Today's Highlights
Patch release v1.17.12 shipped, bringing critical MCP reconnection and adaptive-thinking fixes for Claude Sonnet 5. A massive PR (#34116) landed that resolves 15 separate "question tool" UI bugs, finally addressing one of the community's longest-standing pain points. Meanwhile, the community continues to push for crypto payment support, configurable output token limits, and better session management.

## Releases
**[v1.17.12](https://github.com/anomalyco/opencode/releases/tag/v1.17.12)**  
- Enable adaptive thinking for Claude Sonnet 5.  
- Prefer MCP content responses over structured output when both are available.  
- Reconnect MCP servers after OAuth even if the server was disabled (by @MaxAnderson95).  
- Request MCP refresh-token scope during OAuth.  
- Show MCP OAuth completion status in UI.

## Hot Issues (10 noteworthy)
1. **[#4821 – Add ability to unqueue messages](https://github.com/anomalyco/opencode/issues/4821)**  
   *Open, 60 👍, 18 comments*  
   Users want to cancel queued messages after over-correcting the agent. High demand with no resolution yet.

2. **[#14965 – Slow startup in Ghostty](https://github.com/anomalyco/opencode/issues/14965)**  
   *Open, 16 comments*  
   Startup delay is terminal-specific (Ghostty vs. Alacritty/Kitty). Community suspects conflicting terminal capabilities.

3. **[#23153 – Pay with crypto](https://github.com/anomalyco/opencode/issues/23153)**  
   *Open, 25 👍, 14 comments*  
   Feature request to allow crypto payments for OpenCode Go. Strong community interest.

4. **[#29363 – `limit.output` silently capped at 32k tokens](https://github.com/anomalyco/opencode/issues/29363)**  
   *Open, 5 👍, 12 comments*  
   Config value is ignored beyond 32k; the only workaround is an experimental env var. Users frustrated by silent cap.

5. **[#10998 – Run Command message not showing command in Zed](https://github.com/anomalyco/opencode/issues/10998)**  
   *Closed, 8 👍, 8 comments*  
   ACP integration with Zed shows description but not the actual command. Root cause still unclear.

6. **[#3393 – Windows keyboard shortcuts conflict with OS](https://github.com/anomalyco/opencode/issues/3393)**  
   *Closed, 6 👍, 7 comments*  
   `Ctrl+Shift+Esc` and `Ctrl+Esc` are intercepted by Windows, blocking OpenCode’s shortcuts.

7. **[#28956 – Question prompt overlay blocks response text, no minimize/close](https://github.com/anomalyco/opencode/issues/28956)**  
   *Closed, 6 comments*  
   One of many “question tool” UX issues; users cannot scroll or dismiss the overlay.

8. **[#32669 – Glob tool skips files under dot directories](https://github.com/anomalyco/opencode/issues/32669)**  
   *Open, 4 comments*  
   Explicit dot-dir patterns (e.g. `.ai/*`) return no results. Affects project workflows.

9. **[#33618 – Qwen 3.7 Plus/Max via OpenRouter: empty tool call names](https://github.com/anomalyco/opencode/issues/33618)**  
   *Open, 4 comments*  
   Sporadic `✗ "" failed` errors cause session abortion. Provider compatibility issue.

10. **[#34573 – Cannot move window with title bar in new layout](https://github.com/anomalyco/opencode/issues/34573)**  
    *Open, 1 👍, 3 comments*  
    New Desktop v2 layout breaks basic window dragging.

## Key PR Progress (10 important)
1. **[#34116 – Question UI fixes and UX improvements](https://github.com/anomalyco/opencode/pull/34116)**  
   *Closed, merged*  
   Closes 15 issues related to question tool overlays (scrolling, collapsing, resizing). Major UX win.

2. **[#33920 – fix(mcp): reconnect after OAuth even when server is disabled](https://github.com/anomalyco/opencode/pull/33920)**  
   *Closed, merged*  
   Directly fixes the OAuth flow regression noted in v1.17.12 release. By @MaxAnderson95.

3. **[#26167 – fix(session): retry empty stream truncations and discard partial parts](https://github.com/anomalyco/opencode/pull/26167)**  
   *Open*  
   Handles provider streams that end without a proper `stop_reason`; prevents premature session termination.

4. **[#30025 – fix: support winget opencode upgrades](https://github.com/anomalyco/opencode/pull/30025)**  
   *Closed, merged*  
   Adds Winget detection for auto-upgrades on Windows. Closes #30026.

5. **[#30561 – fix(shell): strip env variable assignments from permission patterns](https://github.com/anomalyco/opencode/pull/30561)**  
   *Open*  
   Prevents false permission prompts when commands contain `VAR=value` prefixes. Rebased from closed PR.

6. **[#34739 – fix(plugin): log server plugin load and skip failures](https://github.com/anomalyco/opencode/pull/34739)**  
   *Open*  
   Makes server plugin failures visible (load/install/compatibility errors). Improves debugging.

7. **[#34740 – feat(tui): show session status in prompt area](https://github.com/anomalyco/opencode/pull/34740)**  
   *Open*  
   Adds tokens/cost/MCP/LSP info to the prompt line when sidebar is hidden.

8. **[#34678 – feat(desktop): session tab hover preview popover](https://github.com/anomalyco/opencode/pull/34678)**  
   *Closed, merged*  
   Hovering a tab shows project name, full path, and last activity.

9. **[#26861 – fix(tui): Old messages disappearing during long sessions](https://github.com/anomalyco/opencode/pull/26861)**  
   *Open*  
   Implements lazy-scroll loading (50 messages at a time) to prevent message loss.

10. **[#30977 – feat(tui): attach to configured server by default](https://github.com/anomalyco/opencode/pull/30977)**  
    *Open*  
    Adds `server.attach` config so TUI automatically connects to a pre-configured server on startup.

## Feature Request Trends
- **Question tool UX improvements** – Users repeatedly ask for collapsible/minimizable overlays, scrollable content, and the ability to review conversation history before answering. These are now largely addressed by PR #34116.
- **Payment flexibility** – Crypto payment support (#23153) is the most upvoted feature this week.
- **Message queue management** – The ability to unqueue messages (#4821) is strongly requested (60 👍). No current implementation.
- **MCP keybind exposure** – Users want a configurable keybinding for “Toggle MCPs” (#24164).
- **Session navigation** – `Alt+↑/↓` to quickly switch between user commands within a session (#34727) reflects a desire for smoother multi-session workflows.
- **Token limit controls** – Unsilence the 32k cap on `limit.output` and allow much larger values for models like DeepSeek (#29363).

## Developer Pain Points
1. **Question tool UI blocking** – Overlays that cannot be dismissed, scrolled, or minimized repeatedly frustrate users (6+ separate issues filed).
2. **Silent config caps** – The 32k output token limit is undocumented and overridden only by an experimental env var, leading to wasted debugging time.
3. **Terminal-specific startup slowness** – Ghostty users experience multi-second delays that don't occur in other terminals.
4. **File descriptor limits on macOS** – `ulimit -n` workarounds are required (#8358); users want auto-detection or graceful handling.
5. **Glob tool inconsistency** – Dot directories are ignored even with explicit patterns (#32669), breaking common workflows.
6. **Provider compatibility issues** – Qwen models via OpenRouter produce empty tool calls (#33618); subagents hang on certain streams (#33028).
7. **Desktop app crashes/performance** – Large `.gitignore` projects cause Electron to hang (#20977); `@filename` includes can crash CLI (#33632).
8. **New layout regressions** – Window dragging (#34573) and tab duplication (#34704) are early-adopter pain points in v2 UI.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-07-01

## Today's Highlights

Pi v0.80.3 shipped with **Claude Sonnet 5 support** across Anthropic and Bedrock providers, and the community quickly responded with a duplicate issue and PR to add the same model to the GitHub Copilot provider. Meanwhile, a critical update bug (`pi update` failing due to missing `@smithy/node-http-handler`) and a WSL login hang both landed as high-priority issues, signaling that the release had some rough edges. On the positive side, several long-standing feature requests (e.g., `excludeFromContext` for custom messages, AOT compilation for extensions) were closed as merged.

---

## Releases

### v0.80.3
- **New Feature:** Anthropic Claude Sonnet 5 is now available through inherited Anthropic-compatible and Bedrock provider catalogs with adaptive thinking enabled. [Release notes](https://github.com/earendil-works/pi/releases/tag/v0.80.3)

---

## Hot Issues (10 noteworthy)

1. **#5653 – Move off Shrinkwrap** 🟡 *open, 18 comments*  
   Duplicate copies of `@earendil-works/pi-ai` are installed when both ai and coding-agent packages are direct deps, causing module-level state duplication. High community engagement; still under discussion.  
   [GitHub](https://github.com/earendil-works/pi/issues/5653)

2. **#5654 – Add `excludeFromContext` to custom messages** ✅ *closed, 9 comments*  
   Merged quickly after strong community demand (1 👍). Allows `pi.sendMessage()` to skip custom messages from model context, mirroring bash message behavior.  
   [GitHub](https://github.com/earendil-works/pi/issues/5654)

3. **#5501 – Tolerate extra keys on edit tool edits[] items** ✅ *closed, 6 comments*  
   Models sometimes append stray keys (e.g., `newText_strip`) causing validation failures. The fix drops `additionalProperties: false` on inner items while keeping it on outer schema.  
   [GitHub](https://github.com/earendil-works/pi/issues/5501)

4. **#6187 – Pi login hangs in WSL after GitHub Copilot device auth** 🟡 *open, 4 comments*  
   Device authorization completes in browser but pi client in WSL doesn’t detect it. Marked `bug` and `inprogress`; potential interop issue with WSL networking.  
   [GitHub](https://github.com/earendil-works/pi/issues/6187)

5. **#6202 – Kitty inline image preview reserves space but renders blank** ✅ *closed, 5 comments*  
   Images are sent to model correctly but TUI shows only a blank area. Likely a Kitty-specific rendering issue. Closed as `last-read`; no action taken yet.  
   [GitHub](https://github.com/earendil-works/pi/issues/6202)

6. **#6103 – OpenAI Responses API mislabels empty tool results as "(see attached image)"** 🟡 *open, 4 comments*  
   Latent bug exposed by an extension. Empty tool output is incorrectly replaced with the placeholder string, misleading the model.  
   [GitHub](https://github.com/earendil-works/pi/issues/6103)

7. **#6215 – `pi update` fails on 0.80.3 due to missing `@smithy/node-http-handler@^4.9.1`** ✅ *closed, 2 comments*  
   Urgent dependency resolution failure. Quick report; user likely needed to update lockfile or wait for patch release.  
   [GitHub](https://github.com/earendil-works/pi/issues/6215)

8. **#6214 – Config does not sync packages, `pi update` does not install missing** ✅ *closed, 2 comments*  
   Git-synced `.pi` config across machines fails to install packages listed in config. Reported as bug; reopened gap in extension management.  
   [GitHub](https://github.com/earendil-works/pi/issues/6214)

9. **#6200 – Add Sonnet 5 to GitHub Copilot provider** ✅ *closed, 3 comments*  
   Duplicate of #6208 but earlier. Fast community response (2 👍). Merged same day.  
   [GitHub](https://github.com/earendil-works/pi/issues/6200)

10. **#3083 – TUI spinner row leaks into scrollback** ✅ *closed, 4 comments*  
    Long-standing rendering artifact (since April) finally closed. Duplicate spinner row during execution and incomplete cleanup after exit.  
    [GitHub](https://github.com/earendil-works/pi/issues/3083)

---

## Key PR Progress (10 important)

1. **#6216 – feat: Add Amazon Bedrock Mantle OpenAI Responses provider** 🟡 *open, new*  
   Supersedes #5509. Adds a new provider for Bedrock Mantle with support for GPT 5.5/5.4. Uses OpenAI’s Bedrock client.  
   [GitHub](https://github.com/earendil-works/pi/pull/6216)

2. **#5678 – Add `excludeFromContext` for custom messages** ✅ *closed, merged*  
   Implements feature from issue #5654. Also teaches compaction, summarization, and branch operations to respect the flag.  
   [GitHub](https://github.com/earendil-works/pi/pull/5678)

3. **#6213 – AOT compilation for TypeScript extensions** ✅ *closed, merged*  
   Drastically reduces extension startup time from seconds to milliseconds using esbuild. Integrated into `pi install` and `pi update`.  
   [GitHub](https://github.com/earendil-works/pi/pull/6213)

4. **#6207 – Add Claude Sonnet 5 to GitHub Copilot provider** ✅ *closed, merged*  
   Routes Sonnet 5 through the existing GitHub Copilot provider.  
   [GitHub](https://github.com/earendil-works/pi/pull/6207)

5. **#6205 – Fix composer overlay blocking side panel clicks (context-canvas)** ✅ *closed, merged*  
   CSS fix removing stale absolute/fixed positioning that intercepted pointer events.  
   [GitHub](https://github.com/earendil-works/pi/pull/6205)

6. **#6196 – Return empty string for empty tool results (OpenAI Responses)** ✅ *closed, merged*  
   Fixes #6103. Stops misleading "(see attached image)" on empty tool output.  
   [GitHub](https://github.com/earendil-works/pi/pull/6196)

7. **#6190 – Add `PI_SKILL_PATH` environment variable** ✅ *closed, merged*  
   Allows per-repo skill location via `.envrc`, reducing CLI flag overhead.  
   [GitHub](https://github.com/earendil-works/pi/pull/6190)

8. **#6176 – Apply extension tool changes before next provider request in same run** ✅ *closed, merged*  
   Fixes #6162: extension tools that call `pi.setActiveTools()` now take effect immediately within the same agent run.  
   [GitHub](https://github.com/earendil-works/pi/pull/6176)

9. **#1737 – Optimize prompt caching across multiple providers** ✅ *closed, merged*  
   Marks both final assistant `tool_use` block and final user message with `cache_control`—improves caching beyond single-block marking.  
   [GitHub](https://github.com/earendil-works/pi/pull/1737)

10. **#5509 – Add Amazon Bedrock Mantle provider (superseded)** ✅ *closed*  
    Superseded by #6216 but notable as first attempt. Included GPT 5.5/5.4 models.  
    [GitHub](https://github.com/earendil-works/pi/pull/5509)

---

## Feature Request Trends

- **Model provider parity:** Multiple issues/PRs around adding Claude Sonnet 5 to every provider (Anthropic, Bedrock, Copilot) — clearly the hottest request today.
- **Admin/enterprise controls:** #6159 (admin settings via system config files) and #6151 (image URL passthrough) reflect growing enterprise adoption.
- **Session & title management:** #6209 (AI-generated session titles) and #6206 (clamping to context window) show demand for smarter session lifecycle handling.
- **Extension API improvements:** #6211 (custom skill activation instruction in `formatSkillsForPrompt`) and #6198 (extension ability to call tools by name) indicate power users want deeper extension hooks.
- **Configuration portability:** #6191 (`PI_SKILL_PATH` env var) and #6214 (config sync of packages) — users want declarative, shareable setups.

---

## Developer Pain Points

- **Dependency conflicts and broken updates:** #5653 (duplicate hoisting) and #6215 (missing `@smithy/node-http-handler`) highlight fragility in pi’s npm dependency tree. The Shrinkwrap issue (#5653) has 18 comments, suggesting it’s a systemic problem beyond this release.
- **Tool output handling confusion:** #6103 (empty tool output labelled as "see attached image") and #6181 (misleading timeout error for large values) frustrate developers debugging tool interactions.
- **Platform-specific bugs:** #6187 (WSL login hang) and #6202 (Kitty blank image preview) show that terminal and OS edge cases still surface frequently.
- **Configuration and extension syncing:** #6214 (config not syncing packages) and #6159 (lack of admin overrides) point to gaps in managing pi across multiple machines or teams.
- **Context budget limits:** #6217 (no safe built-in way to keep session within budget) and #6206 (clamping prevents artificial limits) reveal that developers are hitting model context windows more often and want better guardrails.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-07-01

## 1. Today's Highlights

Today’s biggest signal is the **failure of the `v0.19.3-preview.0` release** due to a failed integration job, which blocks the stabilization pipeline. On the feature side, the community submitted a **new bug (Issue #6119)** highlighting inconsistent `.gitignore` handling between `list_directory` and `read_file`—a potential silent correctness risk for file-operations workflows. Meanwhile, **two major PRs landed** for channel identity/lifecycle metadata, setting the stage for persistent channel agents.

---

## 2. Releases

- **v0.19.3-nightly.20260701.a974594d7**  
  A nightly build was tagged, but **its release notes are incomplete** (the description reads “Release notes generated using configuration in .github/release.yml at release/…”). The associated CI run for `v0.19.3-preview.0` **failed** (Issue #6095), indicating the release pipeline is currently blocked.

---

## 3. Hot Issues (10 noteworthy)

1. **#6119 – `list_directory` and `read_file` have inconsistent git-ignore handling**  
   *[OPEN, type/bug]*  
   A newly opened bug that flags a behavioral mismatch: `list_directory` respects `.gitignore`, while `read_file` does not—could lead to unexpected reads or missed files.  
   👥 2 comments, no upvotes yet.  
   https://github.com/QwenLM/qwen-code/issues/6119

2. **#6095 – Release Failed for v0.19.3-preview.0**  
   *[CLOSED, type/bug]*  
   Automated issue created by the CI workflow. The `integration_none` job failed, blocking the preview release. This is a blocker for anyone waiting for the next stable version.  
   👥 2 comments.  
   https://github.com/QwenLM/qwen-code/issues/6095

3. **#1281 – Qwen Code model (deployed via Ollama) returns JSON-format responses**  
   *[CLOSED, type/bug]*  
   A long-standing issue (opened Dec 2025) finally closed today. Users deploying the model locally via Ollama saw unexpected JSON output instead of natural language—community tracked it to a model config mismatch.  
   👥 7 comments.  
   https://github.com/QwenLM/qwen-code/issues/1281

4. **#1105 – Missing “Accept Diff” and “Close Diff Editor” commands in VS Code**  
   *[CLOSED, type/bug]*  
   Chinese-speaking users reported that after launching Qwen Code inside VS Code, the `Ctrl+Shift+P` palette does not show these essential diff commands.  
   👍 1 upvote.  
   https://github.com/QwenLM/qwen-code/issues/1105

5. **#1316 – After modifying initialization file, conversation history is cleared; memory not refreshed**  
   *[CLOSED, type/bug]*  
   Users found that updating the init file resets the session memory, forcing a full reboot to re-load the new configuration—a frustrating UX issue for those tweaking system prompts.  
   👥 5 comments.  
   https://github.com/QwenLM/qwen-code/issues/1316

6. **#1280 – qwen-code cannot use local Ollama-deployed qwen3-coder**  
   *[CLOSED, type/bug]*  
   A user reported a 400 error when switching to a locally hosted Qwen3-Coder-30B-A3B model via the `/auth` command and OpenAI-compatible API. This is a high-impact issue for self-hosters.  
   👥 5 comments.  
   https://github.com/QwenLM/qwen-code/issues/1280

7. **#3174 – Frequent freezing/unresponsiveness (Turkish user)**  
   *[CLOSED, type/bug]*  
   User reports GUI and terminal freezes on Windows, often requiring restarts. No resolution yet, but the report was reopened today.  
   👥 3 comments.  
   https://github.com/QwenLM/qwen-code/issues/3174

8. **#511 – Java code generation missing closing braces (`}`) on Windows**  
   *[CLOSED, type/bug]*  
   A long-lived bug (since Sept 2025) that continues to affect Java developers on Windows. The model consistently omits `}` in Java files, and requests to fix it often result in the deletion of existing `}` characters. High relevance for Java-heavy workflows.  
   👥 2 comments.  
   https://github.com/QwenLM/qwen-code/issues/511

9. **#508 – After upgrade to 0.0.9, `write_file` is no longer called**  
   *[CLOSED, type/bug]*  
   A regression where the model stops invoking `write_file` after a version upgrade, even with the same prompt and model. Essential for users relying on code generation workflows.  
   👥 2 comments.  
   https://github.com/QwenLM/qwen-code/issues/508

10. **#6050 – Add explicit channel memory for messaging channels**  
    *[CLOSED, type/feature-request]*  
    A feature request for persistent room/thread context in multi-user chat channels (Telegram, WeChat, etc.). This aligns with the day’s PR activity on channel identity metadata.  
    👥 2 comments, 👍 0.  
    https://github.com/QwenLM/qwen-code/issues/6050

---

## 4. Key PR Progress (10 important picks)

1. **#6106 – Tool(param:value) permission syntax for parameter-level access control**  
   *[OPEN, type/feature-request]*  
   Adds a granular permission syntax (e.g., `Agent(model:opus)`) to deny subagent launches by model or other tool parameters. This is a significant advancement in security and control.  
   https://github.com/QwenLM/qwen-code/pull/6106

2. **#6045 – Reduce multimodal history payload size**  
   *[OPEN]*  
   Replaces historical inline images with text references, reattaching only the most recent ones. Solves memory bloat in long multimodal sessions—critical for vision-heavy workflows.  
   https://github.com/QwenLM/qwen-code/pull/6045

3. **#6072 – /effort command for unified reasoning effort**  
   *[OPEN]*  
   Introduces a provider-agnostic `/effort` command allowing users to set reasoning effort from `low` to `max`. Clamps and translates the choice per provider—addresses a long-standing UX gap.  
   https://github.com/QwenLM/qwen-code/pull/6072

4. **#6107 – Raise stream idle timeout default and hint the env knob**  
   *[OPEN]*  
   Increases the default streaming inactivity timeout from 2 min to 4 min and improves the error message to tell users how to override it via env var.  
   https://github.com/QwenLM/qwen-code/pull/6107

5. **#6098 – Harden daemon-managed channel worker**  
   *[OPEN]*  
   Adds bounded restart supervision, IPC heartbeat monitoring, and richer status fields for channel workers in `qwen serve --channel` mode.  
   https://github.com/QwenLM/qwen-code/pull/6098

6. **#6105 – Add identity and task lifecycle metadata for channels**  
   *[OPEN]*  
   Foundational PR for resident channel agents—adds channel identity metadata, memory scope, and a shared task lifecycle hook. Pairs with #6114 for adapter-level progress indicators.  
   https://github.com/QwenLM/qwen-code/pull/6105

7. **#6114 – Show lifecycle status in adapters (Telegram, WeChat, DingTalk, Feishu)**  
   *[OPEN]*  
   Maps task lifecycle events to native progress surfaces (typing indicators, status reactions). Delivers a more responsive UX in third-party messaging platforms.  
   https://github.com/QwenLM/qwen-code/pull/6114

8. **#5895 – Daemon session artifact APIs**  
   *[OPEN]*  
   Implements first-class session artifacts—structured metadata attached to tool results—enabling hooks and clients to list/add/remove artifacts.  
   https://github.com/QwenLM/qwen-code/pull/5895

9. **#5980 – Prioritize auth-modified env vars over system env vars**  
   *[OPEN]*  
   Fixes a bug where a new session would still use the old API key even after changing providers via `/auth`—a clear regression for users switching models mid-session.  
   https://github.com/QwenLM/qwen-code/pull/5980

10. **#6019 – `/model --compaction` for configurable chat compression model**  
    *[OPEN]*  
    Lets users assign a dedicated model for chat compression (auto-compact), giving advanced users more control over memory management.  
    https://github.com/QwenLM/qwen-code/pull/6019

---

## 5. Feature Request Trends

- **Channel Memory & Identity** (e.g., #6050, PRs #6105 and #6114): Multiple requests for persistent session context in multi-user chat channels, plus lifecycle-aware status indicators. This is the dominant theme today.
- **Granular Permission Control** (PR #6106): The community is pushing for parameter-level access control for tools and subagents, indicating a growing need for security in multi-model/multi-tool workflows.
- **Configurable Compression & Reasoning** (PRs #6019, #6072): Users want to fine-tune which model handles message compression and reasoning effort, rather than relying on defaults.
- **Explicit Memory & Artifact Management** (#6050, PR #5895): A clear demand for first-class, persistent memory/artifact storage across sessions and channels.

---

## 6. Developer Pain Points

- **Ollama/Local Model Compatibility** (#1280, #1281): Self-hosters continue to face issues with model response formats and API connectivity when deploying Qwen models via Ollama. The community often shares workarounds, but the root cause remains unaddressed in the core.
- **Session Memory & Initialization Loss** (#1316, #508): Modifying the init file or updating the client can silently clear conversation history or stop tool invocations—a critical UX gap for power users.
- **Windows & VS Code Stability** (#3174, #511, #1105): Freezes, missing commands in VS Code, and model errors when generating Java code on Windows are recurring and unresolved pain points that erode developer trust.
- **Release Pipeline Fragility** (#6095): A failed preview release today highlights that CI stability is still a concern, especially for the `integration_none` job that blocks the entire release workflow.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest
**2026-07-01** – *Now branded as* **CodeWhale**

---

## Today’s Highlights

- **v0.8.66 shipped as CodeWhale** – the project has officially rebranded; the legacy `deepseek-tui` npm package is deprecated. Users must migrate per `docs/REBRAND.md`.
- **v0.8.67 setup wizard takes shape** – multiple issues and PRs this week focus on a constitution-first onboarding flow, provider/model readiness, and guided run-time posture selection.
- **Community spots lingering rebrand bugs** – sub-agent state still writes to `.deepseek/` instead of `.codewhale/` (Issue #3864), and several users report copy/paste blocking after the update (Issue #3868).

---

## Releases

| Version | Notes |
|---------|-------|
| **v0.8.66** | CodeWhale is now the canonical project/command/npm package. Legacy `deepseek-tui` receives no further releases. <br>Rebrand migration required. |

No other releases in the last 24h.

---

## Hot Issues (10)

*Picked by comment count, community reaction, and impact.*

1. **#2487 – “Turn stalled – no completion signal received”**  
   TUI freezes in `yolo` mode; `continue` fails to resume. 18 comments, high severity.  
   *[Link](https://github.com/Hmbown/CodeWhale/issues/2487)*

2. **#3275 – Overly involved agent; self-questioning and deviating from intent**  
   Regression from #3061. User reports the agent loops proposing/answering without waiting. 14 comments.  
   *[Link](https://github.com/Hmbown/CodeWhale/issues/3275)*

3. **#3406 – Runtime posture card with constitution boundary**  
   Maintainer-owned: explicit trust/approval/sandbox selector during setup. 13 comments.  
   *[Link](https://github.com/Hmbown/CodeWhale/issues/3406)*

4. **#3736 – Separate work mode from approval policy**  
   Four overlapping knobs cause UX confusion. 11 comments.  
   *[Link](https://github.com/Hmbown/CodeWhale/issues/3736)*

5. **#2870 – EPIC: staged command-boundary refactor**  
   Tracks smaller mergeable layers for #2791. 10 comments.  
   *[Link](https://github.com/Hmbown/CodeWhale/issues/2870)*

6. **#3793 – Build a guided localized constitution creator**  
   Maintainer-owned: language-first, guided-plus-open-canvas flow. 10 comments.  
   *[Link](https://github.com/Hmbown/CodeWhale/issues/3793)*

7. **#1812 – TUI freeze on Windows (crossterm poll)**  
   Intermittent freeze, no keyboard input, process stays alive. 9 comments.  
   *[Link](https://github.com/Hmbown/CodeWhale/issues/1812)*

8. **#3864 – Sub-agent state persists to `.deepseek/` not `.codewhale/`**  
   Branding bug. 3 comments, opened today.  
   *[Link](https://github.com/Hmbown/CodeWhale/issues/3864)*

9. **#3868 – Copy/paste bug on v0.8.66**  
   Right-click in prompt editor triggers full window override on Windows 11. 1 comment, several reported in similar issues.  
   *[Link](https://github.com/Hmbown/CodeWhale/issues/3868)*

10. **#2261 – Process crash in dialogue; input leaked to PowerShell**  
    Focus loss causes keyboard input to be executed by shell. 4 comments, Windows-specific.  
    *[Link](https://github.com/Hmbown/CodeWhale/issues/2261)*

---

## Key PR Progress (10)

1. **#3869 – Dynamic MCP server infrastructure for McpPool**  
   Foundation for `start_mcp_server` tool. New `dynamic_servers` field with `RwLock`.  
   *[Link](https://github.com/Hmbown/CodeWhale/pull/3869)*

2. **#3866 – LLM can start MCP servers from chat context**  
   Adds `StartRuntimeMcpServer` tool supporting stdio and HTTP transports.  
   *[Link](https://github.com/Hmbown/CodeWhale/pull/3866)*

3. **#3861 – v0.8.67 constitution-first setup foundation**  
   State model, persistence, renderer in `crates/config`. Shared vocabulary for setup readiness.  
   *[Link](https://github.com/Hmbown/CodeWhale/pull/3861)*

4. **#3748 – Add Sakana AI Fugu provider** (merged)  
   New built-in provider matching moonshot/minimax pattern.  
   *[Link](https://github.com/Hmbown/CodeWhale/pull/3748)*

5. **#3784 – Config persistence for GUI config panel**  
   Supports nested-table keys; fixes `allow_shell` type bug.  
   *[Link](https://github.com/Hmbown/CodeWhale/pull/3784)*

6. **#3764 – Improve `/config ask-rules` diagnostics**  
   Reports path, existence, and status of `permissions.toml`.  
   *[Link](https://github.com/Hmbown/CodeWhale/pull/3764)*

7. **#3865 – Fix sub-agent state to `.codewhale/`**  
   Closes #3864. Corrected rebrand-era fallback path.  
   *[Link](https://github.com/Hmbown/CodeWhale/pull/3865)*

8. **#3789 – Show safety policy in `/status`**  
   Adds Safety row with sandbox/network posture per mode.  
   *[Link](https://github.com/Hmbown/CodeWhale/pull/3789)*

9. **#3822 – Prefer exact binary release assets**  
   Fixes asset selection when archive appears before bare binary.  
   *[Link](https://github.com/Hmbown/CodeWhale/pull/3822)*

10. **#3862 – Remove unused approval-cache containers**  
    Cleans up dead `ApprovalCache*` types.  
    *[Link](https://github.com/Hmbown/CodeWhale/pull/3862)*

---

## Feature Request Trends

- **Constitution-first setup & localized onboarding** – Multiple issues (#3402, #3403, #3406, #3793, #3806) demand a guided, language-aware first-run experience with a non-blank constitution editor.
- **Better agent autonomy control** – Users want clear boundaries and confirmation steps to prevent the agent from self-modifying or over-extending scope (#3275, #3736).
- **Dynamic MCP / runtime tool spawning** – PRs #3869 and #3866 indicate growing interest in letting the LLM start MCP servers mid-conversation.
- **Sub-agent lifecycle transparency** – Real-time sidebar updates for sub-agent state (#3837) are requested for multi-agent workflows.
- **Fleet/Whaleflow natural language entry** – Issue #3863 asks for a “one sentence → tasks.json” entry point instead of manual CLI steps.
- **Improved copy/paste handling** – #3868 and related bugs show the current popup/overlay is unusable on Windows.

---

## Developer Pain Points

- **Windows-specific TUI freezes** – #1812, #2261, and #3868 all affect Windows users, with focus loss and unresponsive UI being the main complaints.
- **Rebrand migration friction** – Sub-agent state still writes to `.deepseek/`, and legacy users are unsure about migration (e.g., `docs/REBRAND.md` is mentioned but not fully discoverable).
- **Agent over-engineering** – “Too involved” pattern (#3275) frustrates developers who want the tool to stay within explicit instructions.
- **Misleading UI prompts** – “Ctrl+B backgrounds this command” implies capability that doesn’t reliably work for bash commands (#3859).
- **Incomplete natural language workflow** – Fleet/Whaleflow requires manual JSON editing, which feels “not like an AI agent” to many users (#3863).

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*