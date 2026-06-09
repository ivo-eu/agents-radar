# AI CLI Tools Community Digest 2026-06-09

> Generated: 2026-06-09 04:18 UTC | Tools covered: 9

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

Here is a cross-tool comparison report based on the provided community digest data.

## Cross-Tool Comparison Report: AI CLI Developer Tools
**Date:** 2026-06-09
**Analyst:** Senior Technical Analyst, AI Developer Tools Ecosystem

---

### 1. Ecosystem Overview

The AI CLI tools ecosystem is marked by high activity and a clear convergence on critical features, but is also grappling with significant maturity challenges around data integrity and model reliability. Three dominant themes emerge: **data loss** (Claude Code) and **agent hangs/false completions** (Gemini CLI, GitHub Copilot CLI) are eroding user trust, while platform-specific reliability issues, particularly on **Windows**, remain a persistent pain point across multiple tools. The community is concurrently driving a strong push toward more sophisticated agent orchestration, with **sub-agent management, session lifecycle controls, and MCP ecosystem enhancements** being the most requested cross-tool features.

### 2. Activity Comparison

The following table provides a high-level activity snapshot for the seven tools based on the June 9, 2026 digest data.

| Tool | Hot Issues (Last 24h) | Key PRs (Last 24h) | Release Status |
| :--- | :--- | :--- | :--- |
| **Claude Code** | 10 (High Severity) | 6 | v2.1.169 |
| **OpenAI Codex** | 10 | 10 | rust-v0.138.0 |
| **Gemini CLI** | 10 | 10 | v0.47.0-nightly |
| **GitHub Copilot CLI** | 10 | 1 | None |
| **Kimi Code CLI** | 3 | 0 | None (v1.47.0) |
| **OpenCode** | 10 | 10 | None |
| **Pi** | 10 | 10 | v0.79.0 |
| **Qwen Code** | 10 | 10 | v0.18.0-preview.0 |
| **DeepSeek TUI (CodeWhale)** | 10 | 10 | v0.8.54 |

**Analysis:** Most tools are highly active with a steady flow of issues and PRs. **OpenAI Codex**, **OpenCode**, **Pi**, **Gemini CLI**, and **Qwen Code** show the highest development velocity with 10 PRs each. **GitHub Copilot CLI** and **Kimi Code CLI** show significantly lower PR activity, suggesting a slower development cadence.

---

### 3. Shared Feature Directions

Several feature requirements appear across multiple tool communities, indicating strong ecosystem-wide demand.

- **Session & State Management:**
    - **Claude Code** (Centralized session view, conversation branching), **GitHub Copilot CLI** (session pause/control, multi-session management), **OpenCode** (session lifecycle management, durable event export), **Pi** (file-based undo/rewind), **Gemini CLI** (memory system maturity).
    - *Need:* Users want persistent, controllable, and exportable session state across all tools.

- **MCP Ecosystem & Agent Tooling:**
    - **Claude Code** (sub-agent MCP tool inheritance), **OpenAI Codex** (MCP routing confusion), **Gemini CLI** (atomic MCP tool discovery, configurable timeouts), **GitHub Copilot CLI** (custom MCP registry URL), **OpenCode** (MCP abort signals, pagination).
    - *Need:* MCP (Model Context Protocol) support is becoming a core infrastructure component, requiring robust handling, discovery, and cancellation.

- **Model Flexibility & Provider Support:**
    - **GitHub Copilot CLI** (mid-session model switching, BYOK), **Pi** (new provider integrations), **Qwen Code** (model-provider disambiguation), **DeepSeek TUI** (expansive model catalog).
    - *Need:* Users demand the ability to switch between models (e.g., gpt-5.5, Claude Opus, local) freely within a single session or workflow.

- **Agent Reliability & Safety:**
    - **Claude Code** (AUP false positives, silent data loss), **Gemini CLI** (agent hangs, false success, destructive behavior), **GitHub Copilot CLI** (inconsistent tool execution), **Pi** (project trust system).
    - *Need:* A pervasive demand for more trustworthy, predictable, and safe autonomous agents that respect user intent and system boundaries.

- **Terminal UI/UX Polish:**
    - **GitHub Copilot CLI** (visual delimiters, stash commands), **Pi** (terminal compatibility, CJK rendering), **OpenCode** (clickable file:line references), **DeepSeek TUI** (multi-tab system).
    - *Need:* A frictionless terminal experience is a baseline expectation, with demand for modern IDE-like features (tabs, autocomplete, export).

- **Security & Permissions:**
    - **Claude Code** (`--safe-mode`), **Gemini CLI** (SSRF protection, private IP blocking, secret redaction), **Pi** (Project Trust system), **Qwen Code** (workspace approval gating), **OpenCode** (immutable releases).
    - *Need:* As agents gain more capabilities, security gating and permission models are a growing priority.

- **Windows Platform Support:**
    - **OpenAI Codex** (sandbox fails, crash handling), **GitHub Copilot CLI** (WSL delays, uninstallation issues, ReFS problems), **OpenCode** (PowerShell UTF-8), **Pi** (terminal popups).
    - *Need:* Windows is clearly the most problematic platform, with recurring issues in sandbox, terminal, and installation.

---

### 4. Differentiation Analysis

The tools are diverging in their fundamental design philosophy and target user base.

- **Claude Code (Anthropic):** The most feature-rich and mature, acting as the pace-setter. It is a high-authority, custom-driven agent environment.
    - *Differentiation:* Balances maximum autonomy (sub-agents, dynamic workflows) with deep customization (hooks, plugins). Its primary weakness is **instability** (data loss, AUP glitches) which risks its trust position.
- **OpenAI Codex:** A deeply integrated part of the OpenAI/ChatGPT ecosystem, offering a seamless desktop-to-CLI handoff.
    - *Differentiation:* Focus on smooth model lifecycle, platform parity, and desktop integration. Its core tension is between **fast feature delivery** (pushing many PRs) and **model reliability** (GPT-5.5 issues).
- **Gemini CLI (Google):** A multi-agent orchestration platform with a strong emphasis on modularity and security.
    - *Differentiation:* A clear focus on robust sub-agent delegation and security (SSRF prevention, OAuth). Its main challenge is **core workflow stability**, with agent hangs undermining confidence in its architecture.
- **GitHub Copilot CLI:** The most constrained tool, existing within the GitHub Copilot ecosystem and relying on a single backend.
    - *Differentiation:* Deep integration with GitHub’s platform (hooks, agents). Its slow development pace and persistent platform bugs suggest a **mature but lower-priority** product.
- **Pi, OpenCode, Qwen Code:** These are open-source, community-driven alternatives focused on reproducing and improving upon the features of commercial tools (Claude Code).
    - *Differentiation:* High innovation velocity, particularly in **security gating** (Pi’s Project Trust) and **multi-tab/task** management (OpenCode, Qwen Code). They are more experimental and responsive to community feedback.

---

### 5. Community Momentum & Maturity

- **Rapid Iteration (High Momentum):** **OpenAI Codex**, **OpenCode**, **Pi**, **Qwen Code**, and **Gemini CLI** demonstrate the highest development velocity. They are actively shipping new features and fixes, with large numbers of PRs being merged daily.
- **High Community Engagement (Vocal & Active):** **Claude Code** has the most vocal user base, evidenced by the high engagement on critical bugs (e.g., #63060 with 79 comments). **OpenAI Codex** also shows strong community interaction on model behavior issues.
- **Mature but Slower:** **GitHub Copilot CLI** and **Kimi Code CLI** show less activity. The former appears to be in a maintenance-like phase, while the latter is transitioning its codebase (deprecating Python CLI for TypeScript rewrite), leading to user frustration.
- **Underlying Volatility (Maturity Risk):** **Claude Code**’s data loss bugs and **Gemini CLI**'s agent hanging issues are signs that despite high feature velocity, fundamental reliability for core workflows is not yet mature. Users are acting as **QA testers** for core features.

---

### 6. Trend Signals

These community signals are valuable for any developer building or investing in AI developer tools.

- **Reliability is the New Feature:** Users will tolerate missing features far less than they will tolerate **unreliable core workflows** (data loss, agent hangs, model crashes). The tools with the most serious reliability issues (Claude Code, Gemini CLI) face the highest level of user frustration.
- **Security is a First-Class Feature:** The independent introduction of **trust gating**, **project approval**, and **SSRF protection** across Pi, Qwen Code, and Gemini CLI indicates that security is evolving from an afterthought to a core architectural requirement for agentic tools.
- **The "Claude Code Tax":** Claude Code has set the market's feature expectations. Every other tool is now explicitly or implicitly trying to achieve "Claude Code parity" (e.g., Qwen Code’s PRs on Dynamic Workflows, `safe-mode`, and workspace approval). This is the most powerful "trend signal" shaping the entire ecosystem.
- **Platform Fragmentation is a Cost Barrier:** The persistent and severe issues on **Windows** (and to a lesser extent, FreeBSD/Linux-Wayland) represent a significant barrier to enterprise adoption. Tools that can offer a truly cross-platform, reliable experience will have a major competitive advantage.
- **Agent Observability is Critical:** The demand for **structured error logs, durable event exports, and session lifecycle management** shows that users are moving beyond just using the tool. They need to **instrument**, **monitor**, and **retrace** agent behavior, which is essential for adoption in CI/CD and professional development environments.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report

**Data snapshot:** 2026-06-09 | Source: github.com/anthropics/skills

---

## 1. Top Skills Ranking

The following pull requests represent the most-discussed Skill proposals and improvements in the community. All remain **open** as of this report.

| # | Skill / PR | Functionality | Discussion Highlights | Status |
|---|-----------|---------------|----------------------|--------|
| **#514** | [document-typography](https://github.com/anthropics/skills/pull/514) | Prevents orphan words, widow paragraphs, and numbering misalignment in generated documents. Addresses a universal pain point across all Claude-generated content. | Strong community resonance — this is a cross-cutting quality issue that affects every document Claude creates. | OPEN |
| **#486** | [ODT skill](https://github.com/anthropics/skills/pull/486) | OpenDocument (.odt, .ods) creation, template filling, and conversion to HTML. Targets LibreOffice and ISO standard document workflows. | Broad applicability in enterprise and open-source office environments. Discussion focuses on format fidelity. | OPEN |
| **#210** | [frontend-design](https://github.com/anthropics/skills/pull/210) | Revises the existing frontend design skill for clarity, actionability, and internal coherence — ensuring every instruction is executable in a single conversation. | Highly engaged debate on what constitutes "actionable" guidance vs. generic advice for an LLM. | OPEN |
| **#83** | [skill-quality-analyzer & skill-security-analyzer](https://github.com/anthropics/skills/pull/83) | Two meta-skills: one evaluates Skill quality across structure/documentation/examples; the other audits for security vulnerabilities. | Meta-skills are a new category — community interested in self-improving the ecosystem. Discussion around evaluation criteria. | OPEN |
| **#181** | [SAP-RPT-1-OSS predictor](https://github.com/anthropics/skills/pull/181) | Wraps SAP’s open-source tabular foundation model for predictive analytics on SAP business data. | Brings enterprise ML directly into Claude workflows. Interest from SAP ecosystem developers. | OPEN |
| **#1140** | [agent-creator (meta-skill)](https://github.com/anthropics/skills/pull/1140) | Creates task-specific agent sets. Also fixes multi-tool evaluation stability and Windows support. | High activity as it addresses a fundamental gap: dynamically assembling agents. Also ties into broader Windows compatibility discussions. | OPEN |
| **#723** | [testing-patterns](https://github.com/anthropics/skills/pull/723) | Comprehensive testing guide covering philosophy (Testing Trophy), unit testing, React component testing, integration testing, and mocking. | Large surface area — community split on whether a single skill can cover the full stack effectively vs. focusing on one layer. | OPEN |
| **#335** | [masonry-generate-image-and-videos](https://github.com/anthropics/skills/pull/335) | CLI-based skill for AI image (Imagen 3.0) and video (Veo 3.1) generation, including job management. | Early traction for multimodal generation skills. Discussion around CLI integration vs. direct API usage. | OPEN |

---

## 2. Community Demand Trends

From the most-active **Issues**, three clear demand directions emerge:

1. **Enterprise integration & governance**  
   - [#228 — Org-wide skill sharing](https://github.com/anthropics/skills/issues/228): Users want shared skill libraries and direct sharing links, not manual `.skill` file transfers.  
   - [#492 — Security of community skills under anthropic/ namespace](https://github.com/anthropics/skills/issues/492): Trust boundary concerns; demand for namespace verification and permission scoping.  
   - [#412 — Agent governance skill proposal](https://github.com/anthropics/skills/issues/412): Policy enforcement, threat detection, and audit trails for agent systems.

2. **Infrastructure and tooling**  
   - [#202 — Skill-creator best practices](https://github.com/anthropics/skills/issues/202): Long-running discussion on operational vs. educational tone in skill instructions.  
   - [#16 — Skills as MCPs](https://github.com/anthropics/skills/issues/16): Desire to expose Skills via Model Context Protocol for programmable API-like interfaces.  
   - [#1220 — Multi-file preload / inline bundling](https://github.com/anthropics/skills/issues/1220): Skills split across multiple reference files need better delivery into agent context.

3. **Cross-platform compatibility**  
   - [#556 — run_eval.py 0% trigger rate](https://github.com/anthropics/skills/issues/556) and multiple Windows-related PRs (e.g., #1099, #1050) signal strong demand for reliable evaluation tooling, especially on Windows.

---

## 3. High-Potential Pending Skills

These PRs are actively discussed but not yet merged — likely to land in the near term based on community momentum:

- **[#190 – n8n-builder & n8n-debugger](https://github.com/anthropics/skills/pull/190)** — Four production-tested skills including n8n workflow automation and a FAF-formatted context system. Has been open since late 2025 with sustained activity.
- **[#154 – shodh-memory](https://github.com/anthropics/skills/pull/154)** — Persistent memory system for AI agents, maintaining context across conversations. A foundational primitive for long-running agents.
- **[#568 – ServiceNow platform skill](https://github.com/anthropics/skills/pull/568)** — Broad coverage of ITSM, ITOM, SecOps, and IntegrationHub. Enterprise demand is high.
- **[#444 – AURELION skill suite](https://github.com/anthropics/skills/pull/444)** — Four skills (kernel, advisor, agent, memory) forming a structured cognitive framework. Represents a growing interest in "skill suites" rather than single-topic skills.

---

## 4. Skills Ecosystem Insight

The community’s most concentrated demand is for **skills that improve the quality and trustworthiness of Claude’s output** (typography, frontend design, testing patterns) paired with **enterprise integration capabilities** (ODT, SAP, ServiceNow, n8n), while simultaneously building **meta-skills** (quality analyzers, agent creators) to sustain and scale the ecosystem itself.

---

# Claude Code Community Digest — 2026-06-09

## Today’s Highlights

The latest release v2.1.169 introduces a `--safe-mode` flag for troubleshooting and a new `/cd` command for session directory switching. The community is grappling with a wave of **silent data loss** bugs — multiple reports confirm that the cleanup feature deletes session transcripts despite user settings, and several users have lost thousands of sessions after updates. Meanwhile, **Usage Policy false positives** continue to block legitimate coding work, with a persistent regression spanning recent versions.

## Releases

**v2.1.169** (latest)
- Added `--safe-mode` flag (and `CLAUDE_CODE_SAFE_MODE`) to start Claude Code with all customizations (CLAUDE.md, plugins, skills, hooks, MCP servers) disabled for troubleshooting.
- Added `/cd` command to move a session to a new working directory without breaking the prompt cache.

[View release](https://github.com/anthropics/claude-code/releases)

## Hot Issues (10 noteworthy)

1. **#63060 – API Error: Usage credits required for 1M context**  
   Users with active subscriptions are blocked when attempting to use extended context. 79 comments, 30 👍.  
   [Issue](https://github.com/anthropics/claude-code/issues/63060)

2. **#63875 – Model tool call parsing error interrupts sessions**  
   Frequent “tool call could not be parsed (retry also failed)” errors abort in-progress actions. 54 comments, 83 👍 — high community impact.  
   [Issue](https://github.com/anthropics/claude-code/issues/63875)

3. **#62272 – Chat JSONLs deleted despite `cleanupPeriodDays` set high**  
   Deletion triggered by updates/restarts; user published a recovery script for macOS Time Machine. 16 comments.  
   [Issue](https://github.com/anthropics/claude-code/issues/62272)

4. **#41458 – `cleanupPeriodDays: 99999` ignored — 490 sessions deleted**  
   Explicit setting overridden silently. 15 comments, 1 👍 (but data loss severity).  
   [Issue](https://github.com/anthropics/claude-code/issues/41458)

5. **#65866 – AUP classifier false-positive (session poisoning) persists through 2.1.167**  
   Benign git commands blocked entirely; regression spanning 2.1.165–2.1.167. 7 comments.  
   [Issue](https://github.com/anthropics/claude-code/issues/65866)

6. **#33045 – Agent tool isolation: `worktree` has no effect for team agents**  
   Agents run in the main repo instead of isolated worktree. 19 comments, 9 👍.  
   [Issue](https://github.com/anthropics/claude-code/issues/33045)

7. **#30280 – Sub-agents don’t reliably inherit MCP tools**  
   Documentation says they should, but practice shows inconsistency. 10 comments, 12 👍.  
   [Issue](https://github.com/anthropics/claude-code/issues/30280)

8. **#62016 – Claude passes `rg -rn` parsed as `--replace=n`, corrupting output**  
   Ripgrep flag collision leads to silent search corruption. 5 comments, 9 👍 — clever bug.  
   [Issue](https://github.com/anthropics/claude-code/issues/62016)

9. **#62071 – Usage policy classifier blocks legitimate kernel security work**  
   “Hi” after a kernel debugging session triggers violation. 9 comments.  
   [Issue](https://github.com/anthropics/claude-code/issues/62071)

10. **#59248 – Silent retention cleanup deletes all session transcripts with no warning**  
    User lost ability to resume any recent session. 14 comments, 5 👍.  
    [Issue](https://github.com/anthropics/claude-code/issues/59248)

## Key PR Progress (6 recent PRs; only 6 merged/opened in last 24h)

1. **#66416 – fix(plugin-dev): validator scripts abort on first finding due to `set -e`**  
   Prevents early exit during validation. [PR](https://github.com/anthropics/claude-code/pull/66416)

2. **#65286 – fix(plugins): add missing plugin.json manifest for plugin-dev**  
   Enables proper discovery/installation of plugin-dev. [PR](https://github.com/anthropics/claude-code/pull/65286)

3. **#65619 (CLOSED) – fix(plugins): align frontend-design author with marketplace entry**  
   Fixes malformed author/email fields in plugin.json. [PR](https://github.com/anthropics/claude-code/pull/65619)

4. **#66372 – fix(devcontainer): detect Docker daemon failures via `$LASTEXITCODE`**  
   Corrects false-positive Docker daemon checks on PowerShell. [PR](https://github.com/anthropics/claude-code/pull/66372)

5. **#26914 (CLOSED) – docs: add rules frontmatter `paths:` syntax examples and validation hook**  
   Adds examples and a hook to catch broken `paths:` syntax — long-standing silent failure. [PR](https://github.com/anthropics/claude-code/pull/26914)

6. **#66171 – fix(extensibility): follow symlinks in project-controlled gui**  
   Addresses symlink traversal in extensibility.py (referencing #64582). [PR](https://github.com/anthropics/claude-code/pull/66171)

## Feature Request Trends

- **Conversation branching** (#32631, 30👍): Advanced fork/merge/tree navigation for sessions — one of the most-upvoted enhancements.
- **Centralized session view** (#58039): A unified “all sessions” view across CLI, VS Code, JetBrains, desktop app, and claude.ai/code.
- **Local export of conversations** (#39587): Markdown/JSON export for offline backup and sharing.
- **Secret scrubbing in session logs** (#50014): Automated handling of credentials captured in `.jsonl` files.
- **Multi-account switching** (#36151, 320👍): Feature request for mobile app but reflects broader desire for account flexibility.

## Developer Pain Points

- **Silent data loss is the #1 concern** — multiple issues (#62272, #41458, #59248, #46621, #56904, #62041, #64403, #62476) report that Claude Code deletes conversation transcripts without warning, ignoring user settings. This affects all platforms and has led to loss of hundreds to thousands of sessions.
- **Usage Policy (AUP) false positives** (#65866, #62071, #65779) block legitimate development work (git commands, kernel debugging) and can lock entire sessions, with no easy recovery.
- **Model tool call parsing errors** (#63875) and **ripgrep flag misinterpretation** (#62016) show fragility in the model’s tool execution, causing mid-session failures.
- **Cost and credit surprises** (#52819, #63060): Free ultrareview credits consumed on crashes; usage credit errors even on valid subscriptions.
- **MCP tool inheritance inconsistency** (#30280) and **agent isolation failures** (#33045) undermine trust in team/agent workflows.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest – 2026-06-09

---

## Today's Highlights

GPT‑5.5 availability and reliability dominate the conversation today, with a **high‑severity bug (#26892)** reporting that `gpt-5.5` is listed as available locally but fails with a 404 on actual requests, and a separate thread (#26876) documenting gradual model degradation during complex workflows. On the development side, maintainers are pushing a **four‑PR series to expose dedicated Python goal operations** (#27110–#27113), and a new release `rust-v0.138.0` ships Desktop handoff on macOS/Windows and local image attachments.

---

## Releases

**`rust-v0.138.0`** (full release)
- **`/app` command** can now hand off the current CLI thread into Codex Desktop on macOS and native Windows; Windows workspace launches can open directly into Desktop without a manual prompt. (#25638, #26500)
- Local image attachments and standalone image generation support.

Additionally, several alpha releases were cut today: `rust-v0.139.0-alpha.1`, `rust-v0.139.0-alpha.2`, `rust-v0.138.0-alpha.7`, and `rust-v0.138.0-alpha.8`.

---

## Hot Issues

1. **#26892 – `gpt-5.5` listed as available but returns 404**  
   *Labels: bug, windows-os, CLI, app* | 👍 28 | 💬 77  
   The highest‑activity issue today. Users on both Desktop and CLI see `gpt-5.5` in local model metadata, but requests to the Codex responses endpoint fail with `Model not found`. Workaround: use `gpt-5.4`. Community suspects a stale metadata cache or a server‑side rollout mismatch.  
   [GitHub](https://github.com/openai/codex/issues/26892)

2. **#24391 – Windows sandbox spawn/setup refresh fails**  
   *Labels: bug, windows-os, sandbox, CLI* | 👍 25 | 💬 35  
   Shell commands fail after updating to `codex-cli 0.133.0` on Windows. A long‑running frustration that persists across several patch releases.  
   [GitHub](https://github.com/openai/codex/issues/24391)

3. **#2880 – Copy/Export Message as Markdown**  
   *Labels: enhancement, TUI* | 👍 67 | 💬 22  
   One of the highest‑voted feature requests. Users want to copy conversation messages as Markdown for documentation and issue filing. Currently only plaintext or manual extraction is available.  
   [GitHub](https://github.com/openai/codex/issues/2880)

4. **#21128 – Desktop silently hides old project conversations**  
   *Labels: bug, app, session* | 👍 16 | 💬 21  
   Conversations fall outside the global “recent‑50” window and disappear from the UI, making the Desktop app unreliable as working memory for long‑running projects.  
   [GitHub](https://github.com/openai/codex/issues/21128)

5. **#26876 – GPT‑5.5 degradation over time**  
   *Labels: bug, model-behavior, CLI* | 💬 5  
   Since the April 24 rollout of GPT‑5.5, users report material quality degradation during complex engineering workflows. Suggests a runtime regression or context‑management issue.  
   [GitHub](https://github.com/openai/codex/issues/26876)

6. **#27124 – Codex tries `resources/read` on a tool‑only MCP server**  
   *Labels: bug, MCP* | 💬 0 (filed today)  
   Codex over‑routes URL fetching through MCP resources even when only a `fetch` tool is available. Leads to empty results and a poor user experience with tool‑only MCP servers.  
   [GitHub](https://github.com/openai/codex/issues/27124)

7. **#27121 – Manual Codex memory refresh + shared memory with ChatGPT**  
   *Labels: enhancement, memory* | 💬 0 (filed today)  
   Users want explicit control over Codex memory (ask it to remember/update), similar to ChatGPT. Currently Codex memory feels opaque and auto‑trigger only.  
   [GitHub](https://github.com/openai/codex/issues/27121)

8. **#25921 – Crashpad pending dumps grow to ~5 GB**  
   *Labels: bug, app, performance* | 💬 3  
   Codex Desktop continuously generates `.dmp` and `_sidecar.json` files in `~/Library/Application Support/.../Crashpad/pending`. One user saw 54,504 files in a day.  
   [GitHub](https://github.com/openai/codex/issues/25921)

9. **#27120 – Desktop exits after `git add -A` runs against home directory**  
   *Labels: bug, windows-os, code-review, app* | 💬 1 (filed today)  
   Windows Desktop repeatedly exits without a crash dump when `git add -A` is executed in the background against the user’s home directory. Suspected AppX container destruction.  
   [GitHub](https://github.com/openai/codex/issues/27120)

10. **#27027 – Weekly quota not reset on unified reset (June 4)**  
    *Labels: bug, codex-web, rate-limits* | 💬 4  
    A Plus account did not receive the June 4 quota reset. The user’s reset date remains June 14 while other accounts moved to June 11. Metering inconsistency.  
    [GitHub](https://github.com/openai/codex/issues/27027)

---

## Key PR Progress

1. **#27110–#27113 – Python SDK goal operations (4‑PR series)**  
   Authored by `aibrahim-oai`. Adds `run_goal(objective)` / `start_goal(objective)` to the Python SDK, a private routing foundation, and end‑to‑end integration coverage.  
   [PR #27110](https://github.com/openai/codex/pull/27110) | [PR #27111](https://github.com/openai/codex/pull/27111) | [PR #27112](https://github.com/openai/codex/pull/27112) | [PR #27113](https://github.com/openai/codex/pull/27113)

2. **#27105 – Refresh account plan from `/usage`**  
   Keeps the latest plan in memory, scoped to the current ChatGPT account. Resolves plan from usage before falling back to token claims.  
   [GitHub](https://github.com/openai/codex/pull/27105)

3. **#26880 – Preserve `fsmonitor` for worktree Git reads**  
   Stops Codex from forcing `core.fsmonitor=false`, which previously disabled Git’s built‑in daemon in large repositories.  
   [GitHub](https://github.com/openai/codex/pull/26880)

4. **#27115 – Break down between‑sampling overhead**  
   Splits `between_sampling_overhead_ms` into post‑response, retry, compaction, follow‑up, request‑preparation, and residual phases for better latency tracing.  
   [GitHub](https://github.com/openai/codex/pull/27115)

5. **#27039 – Add detached async command hooks**  
   Supports `async: true` for command hooks, allowing non‑blocking, informational results on later user turns.  
   [GitHub](https://github.com/openai/codex/pull/27039)

6. **#25704 – Normalize Codex images for Responses strict mode**  
   Feature‑flagged. Converts local/data URL images into prepared data URL format, enabling strict‑mode image support in the Responses API.  
   [GitHub](https://github.com/openai/codex/pull/25704)

7. **#27102 – Centralize plugin telemetry metadata construction**  
   Pure refactor, extracting metadata logic from a larger combined PR (#26281) for independent review.  
   [GitHub](https://github.com/openai/codex/pull/27102)

8. **#27100 / #27099 – Plugin analytics smoke workflows**  
   Two PRs adding non‑mutating and account‑mutating end‑to‑end validation for remote plugin analytics.  
   [PR #27100](https://github.com/openai/codex/pull/27100) | [PR #27099](https://github.com/openai/codex/pull/27099)

9. **#27116 – Stop mirroring Codex user input into realtime**  
   Prevents raw typed messages, steers, and worker reports from being duplicated to the realtime frontend model. Keeps orchestrator output as the single source of truth.  
   [GitHub](https://github.com/openai/codex/pull/27116)

10. **#27114 – Prompt realtime turns for voice‑ready final responses**  
    Ensures every realtime turn (including typed or steered turns) produces a concise, standalone final message for the voice frontend.  
    [GitHub](https://github.com/openai/codex/pull/27114)

---

## Feature Request Trends

- **Improved copy/export capabilities** – Users consistently request the ability to copy messages as Markdown ( #2880, 👍67) and add token‑usage breakdowns to the status line ( #13222, 👍24).
- **Session management & storage** – Requests for delta/DAG‑based storage to avoid duplication when forking sessions ( #22593) and for versioned, stable session action exports for external tooling ( #18469).
- **Memory & control** – Manual memory refresh and optional shared memory with ChatGPT ( #27121, filed today) reflect a desire for explicit, predictable memory behavior.
- **Cross‑platform parity** – Repeated requests for Windows‑specific fixes (sandbox, localization, crash handling) suggest an uneven Desktop experience between platforms.

---

## Developer Pain Points

- **Windows‑specific reliability** – Multiple open bugs around sandbox setup failures ( #24391), broken text rendering ( #27123), update/PSModulePath inheritance ( #27117), and AppX container destruction after background `git` operations ( #27120).
- **GPT‑5.5 instability** – “Ghost” availability ( #26892, 404 errors on real requests), quality degradation during long sessions ( #26876), and mid‑task stoppages through Amazon Bedrock ( #26860) all point to a systemic rollout issue.
- **Performance & storage bloat** – Large JSONL rollout files causing UI freezes ( #22991) and Crashpad dumps growing to ~5 GB ( #25921) impact long‑running sessions.
- **MCP routing confusion** – Codex incorrectly attempts `resources/read` on tool‑only servers ( #27124) and fails to install MCP servers for plugins on Windows ( #26693).
- **Rate‑limit & usage metering** – Users report payment cycles not resetting usage ( #25821) and inconsistent quota resets across accounts ( #27027), eroding trust in billing transparency.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-06-09

## Today's Highlights
A new nightly release (v0.47.0) lands with an Antigravity banner limit fix and the "experimental" tag removed from browser agent docs. Community attention remains focused on agent reliability, with three high-comment issues still unresolved around subagent hangs, false success reporting, and underutilization of skills. On the PR side, a flurry of community-contributed fixes for credential preservation, Docker sandbox detection, and PTY resizing are moving through review.

## Releases
- **[v0.47.0-nightly.20260609.g0567b25a2](https://github.com/google-gemini/gemini-cli/releases/tag/v0.47.0-nightly.20260609.g0567b25a2)**  
  Two changes:
  1. Limits the max display count of the Antigravity transition banner (`@DavidAPierce`).
  2. Removes "experimental" label from browser agent documentation (`@gsquared94`).

## Hot Issues (10 noteworthy)
1. **[#21409 — Generalist agent hangs](https://github.com/google-gemini/gemini-cli/issues/21409) (P1, 7 comments, 👍8)**  
   The most-upvoted open bug. CLI hangs indefinitely when deferring to the generalist agent for simple tasks (e.g., folder creation). Workaround: explicitly forbid subagent delegation. *Community sentiment: frustrated — this is a core workflow blocker.*

2. **[#22323 — Subagent false success after MAX_TURNS](https://github.com/google-gemini/gemini-cli/issues/22323) (P1, 6 comments, 👍2)**  
   `codebase_investigator` reports `status: "success"` even when hitting its turn limit without doing any analysis. *Impact: hides real failures — users get misleading GOAL signals.*

3. **[#21968 — Gemini doesn’t use skills/sub-agents enough](https://github.com/google-gemini/gemini-cli/issues/21968) (P2, 6 comments)**  
   Even with well-described custom skills, the model rarely invokes them autonomously. *Community asks for better tool selection heuristics.*

4. **[#22745 — AST-aware file reads/search/mapping](https://github.com/google-gemini/gemini-cli/issues/22745) (P2, 7 comments, 👍1)**  
   Epic tracking whether AST-aware tools can reduce turn count, token noise, and improve codebase understanding. *Linked sub-issues on AST grep and mapping.*

5. **[#24353 — Robust component level evaluations](https://github.com/google-gemini/gemini-cli/issues/24353) (P1, 7 comments)**  
   Follow-up to behavioral evals initiative. Already 76 test cases; seeks better coverage, automation, and component isolation. *Foundation work for quality gates.*

6. **[#26525 — Auto Memory deterministic redaction & logging](https://github.com/google-gemini/gemini-cli/issues/26525) (P2, 5 comments)**  
   Auto Memory sends transcripts to model context before secrets redaction occurs. *Security concern — raises data-leak risk.*

7. **[#25166 — Shell command stuck "Waiting input" after completion](https://github.com/google-gemini/gemini-cli/issues/25166) (P1, 4 comments, 👍3)**  
   CLI hangs post-execution even for trivial commands. *Commonly reported, hard to reproduce — high frustration.*

8. **[#21983 — Browser subagent fails on Wayland](https://github.com/google-gemini/gemini-cli/issues/21983) (P1, 4 comments, 👍1)**  
   Linux/Wayland users hit `GOAL` termination with no output. *Platform fragmentation issue.*

9. **[#22186 — get-shit-done output hook crashes](https://github.com/google-gemini/gemini-cli/issues/22186) (P1, 3 comments)**  
   Repeated crash when GSD agent prints summary. *Blocks the core "get-shit-done" workflow entirely.*

10. **[#22672 — Agent should discourage destructive behavior](https://github.com/google-gemini/gemini-cli/issues/22672) (P2, 2 comments, 👍1)**  
    Model uses `git reset --force`, dangerous DB operations without warning. *Safety concern — users want guardrails.*

## Key PR Progress (10 important)
1. **[#27749 — Vertex AI model mapping fix](https://github.com/google-gemini/gemini-cli/pull/27749) (OPEN)**  
   Routes `gemini-3.5-flash` correctly for LOGIN_WITH_GOOGLE and COMPUTE_ADC auth types. *Fixes server-side model ID mismatch.*

2. **[#27428 — Docker inspect exit code for sandbox detection](https://github.com/google-gemini/gemini-cli/pull/27428) (CLOSED, P1)**  
   Replaces fragile `docker images -q` stdout parsing with `docker inspect` exit code. *Fixes false-negative image existence checks under BuildKit.*

3. **[#27463 — Preserve refresh_token in file-based cacheCredentials](https://github.com/google-gemini/gemini-cli/pull/27463) (OPEN, P1)**  
   Resolves `refresh_token` overwrite for default file storage (not just encrypted). *Community-contributed fix for #21691.*

4. **[#27438 — Configurable per-tool-call timeout](https://github.com/google-gemini/gemini-cli/pull/27438) (CLOSED)**  
   Adds `tools.callTimeout` config and enforces it in `ToolExecutor`. *Gives users escape valve from hung tools.*

5. **[#27619 — Atomic MCP tool discovery](https://github.com/google-gemini/gemini-cli/pull/27619) (OPEN)**  
   Retains last-known MCP tools during transient network failures instead of wiping the registry. *Addresses "tool not found" errors.*

6. **[#27603 — Platform-aware shell guidance](https://github.com/google-gemini/gemini-cli/pull/27603) (OPEN)**  
   Adds Windows-specific command examples in the operational prompt. *Improves cross-platform usability.*

7. **[#27626 — Block private OAuth metadata URLs](https://github.com/google-gemini/gemini-cli/pull/27626) (OPEN, P2, area/security)**  
   SSRF protection: prevents MCP OAuth discovery from calling internal IPs. *Security hardening.*

8. **[#27505 — Fix CJK rendering: prevent extra spaces on width-0 continuation cells](https://github.com/google-gemini/gemini-cli/pull/27505) (OPEN, P2)**  
   Corrects terminal serialization for wide characters. *Fixes copy-paste errors for international users.*

9. **[#27698 — Zero-quota fail-fast to prevent retry loop](https://github.com/google-gemini/gemini-cli/pull/27698) (OPEN)**  
   Detects hard `0` quota limit and stops retrying after 10 attempts. *Unbilled free-tier users stuck in hang.*

10. **[#27744 — DNS resolution before SSRF guard](https://github.com/google-gemini/gemini-cli/pull/27744) (OPEN, size/m)**  
    Resolves hostname to IP before checking private ranges — blocks `127.0.0.1.nip.io` style bypasses. *Security improvement.*

## Feature Request Trends
- **AST-aware tooling**: Multiple issues ([#22745](https://github.com/google-gemini/gemini-cli/issues/22745), [#22746](https://github.com/google-gemini/gemini-cli/issues/22746), [#22747](https://github.com/google-gemini/gemini-cli/issues/22747)) propose replacing naive file reads/searches with AST-based grep, tilth, or glyph. Expected benefits: fewer turns, less token noise.
- **Agent self-awareness & tool delegation**: Users want the agent to know its own capabilities (flags, hotkeys) ([#21432](https://github.com/google-gemini/gemini-cli/issues/21432)) and autonomously invoke skills/sub-agents ([#21968](https://github.com/google-gemini/gemini-cli/issues/21968)).
- **Remote & background agents**: Epic-level push for remote agents with advanced auth ([#20303](https://github.com/google-gemini/gemini-cli/issues/20303)) and backgroundable local subagents ([#22741](https://github.com/google-gemini/gemini-cli/issues/22741)).
- **Memory system maturity**: Three issues from `SandyTao520` ([#26516](https://github.com/google-gemini/gemini-cli/issues/26516), [#26522](https://github.com/google-gemini/gemini-cli/issues/26522), [#26523](https://github.com/google-gemini/gemini-cli/issues/26523)) target Auto Memory reliability — deduplication, invalid patch quarantine, and background retry loops.
- **Security & redaction**: Requests for deterministic secret redaction before model context ([#26525](https://github.com/google-gemini/gemini-cli/issues/26525)) and guardrails against destructive operations ([#22672](https://github.com/google-gemini/gemini-cli/issues/22672)).

## Developer Pain Points
1. **Agent hangs & false completions** — The most vocal pain: agents either hang forever ([#21409](https://github.com/google-gemini/gemini-cli/issues/21409), [#25166](https://github.com/google-gemini/gemini-cli/issues/25166)) or report success when failing ([#22323](https://github.com/google-gemini/gemini-cli/issues/22323)). Erases trust in the tool.
2. **Subagent permission bypass** — Since v0.33.0, subagents run even when disabled in settings ([#22093](https://github.com/google-gemini/gemini-cli/issues/22093)). *Configuration violation.*
3. **Browser agent fragility on Linux** — Wayland failures ([#21983](https://github.com/google-gemini/gemini-cli/issues/21983)) and ignored `settings.json` overrides ([#22267](https://github.com/google-gemini/gemini-cli/issues/22267)) make browser automation unreliable outside macOS.
4. **Terminal rendering bugs** — Flicker on resize ([#21924](https://github.com/google-gemini/gemini-cli/issues/21924)), corruption after external editor exit ([#24935](https://github.com/google-gemini/gemini-cli/issues/24935)), and wrong CJK spacing ([#27505](https://github.com/google-gemini/gemini-cli/pull/27505)) plague daily CLI use.
5. **Destructive behavior** — Model uses `git reset --force`, deploys unsafe DB commands without prompting ([#22672](https://github.com/google-gemini/gemini-cli/issues/22672)). *Needs safety layer.*

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-06-09

## Today’s Highlights
No new releases were published in the last 24 hours, but several long-standing bugs and feature requests saw renewed discussion. A critical background sub‑agent hang with `model="gpt-5.5"` (#3547) and a wrong URL construction for custom MCP registries (#3436) are top concerns. Windows users continue to report startup delays in WSL (#3652) and uninstallation difficulties (#3662). On the PR side, a lone contribution improved the install script to accept an authenticated `GITHUB_TOKEN` (#1960).

## Releases
*None in the last 24 hours.*

---

## Hot Issues (Top 10 by impact & community reaction)

1. **[#1928 – Allow to pause copilot work](https://github.com/github/copilot-cli/issues/1928)**  
   *9 comments · 2 👍*  
   **Why it matters:** Users want to pause and provide mid‑session corrections without losing context. The current workaround (type text and submit) is fragile. High community interest in a native “pause” flow.

2. **[#3547 – Background sub-agent silently hangs at `total_turns=0` with `model="gpt-5.5"`](https://github.com/github/copilot-cli/issues/3547)**  
   *6 comments*  
   **Why it matters:** A reproducible hang that breaks multi‑agent workflows. The sub‑agent reports success but never processes a turn – critical for anyone relying on background agents.

3. **[#3436 – `/mcp search` constructs wrong URL for custom MCP registries](https://github.com/github/copilot-cli/issues/3436)**  
   *5 comments · 1 👍*  
   **Why it matters:** Self‑hosted MCP registries with a `/v0.1/` path segment receive a 404 because the CLI omits that segment. Affects enterprise users with private registries.

4. **[#2867 – "model not supported" after waiting for quota reset (Claude Opus 4.6)](https://github.com/github/copilot-cli/issues/2867)**  
   *5 comments · 1 👍*  
   **Why it matters:** Users who respect quota limits are then locked out permanently – a frustrating UX that undermines trust in automated quota guidance.

5. **[#2540 – Plugin-defined `preToolUse` hooks never fire](https://github.com/github/copilot-cli/issues/2540)**  
   *4 comments · 3 👍*  
   **Why it matters:** Plugin extensibility is broken: hooks are silently ignored both in the main session and in sub‑agents. The most‑upvoted issue in this batch.

6. **[#2201 – `sessionStart` hook doesn’t print to terminal](https://github.com/github/copilot-cli/issues/2201)**  
   *3 comments · 2 👍*  
   **Why it matters:** Tutorials show hooks printing a banner, but nothing appears at CLI startup. Affects onboarding and audit‑log setups.

7. **[#3652 – WSL startup delays (40‑80s) due to `CopilotCLIChatSessionContentProvider.listSessions`](https://github.com/github/copilot-cli/issues/3652)**  
   *3 comments*  
   **Why it matters:** Windows + WSL users face severe latency that makes the tool feel unusable. Educational‑quota users are especially affected.

8. **[#3701 – Runaway MCP server spawning on Windows (IDE lock‑file watcher loop)](https://github.com/github/copilot-cli/issues/3701)**  
   *2 comments*  
   **Why it matters:** A critical bug in v1.0.60: multiple VS Code workspaces cause the `playwright` MCP server to be spawned repeatedly, consuming resources.

9. **[#2966 – Built‑in tooling for managing multiple concurrent CLI sessions](https://github.com/github/copilot-cli/issues/2966)**  
   *2 comments*  
   **Why it matters:** Power users running several `--yolo --autopilot` sessions across repos have no CLI‑level session management – a major productivity gap.

10. **[#3716 – Regression: Function call fails with invalid JSON schema error (v1.0.60)](https://github.com/github/copilot-cli/issues/3716)**  
    *1 comment*  
    **Why it matters:** A newly introduced regression breaks tool‑calling with `$ref` in parameters. Affects all models that rely on function calling.

---

## Key PR Progress

Only **one pull request** was updated in the last 24 hours:

- **[#1960 – install: use GITHUB_TOKEN for authenticated GitHub requests](https://github.com/github/copilot-cli/pull/1960)** (CLOSED)  
  **Summary:** The installation script now passes a `GITHUB_TOKEN` environment variable as an `Authorization` header to avoid rate limits and support installation from private repositories.  
  **Impact:** A small but welcome quality‑of‑life improvement for CI/CD and enterprise setups.

> *Note: PR activity is very low today. The community may be waiting for maintainer responses on the many open issues.*

---

## Feature Request Trends

Based on the latest issues, the most‑requested feature directions are:

- **Session Pause & Control** – #1928, #2966: Users want to pause mid‑session, inject instructions, and manage multiple concurrent sessions (e.g., across branches) without restarting.
- **Model Flexibility** – #3709, #3707, #3705: Switch models mid‑session (including BYOK/local), access lower‑cost open‑weight models, and expand free‑tier model selection beyond Haiku.
- **CLI Hooks & Plugin Extensions** – #3713, #2540, #2201: Hooks that can modify user prompts, print to terminal reliably, and fire in sub‑agents.
- **Non‑interactive / Cron‑like Scheduling** – #3714: Run scheduled tasks via Copilot CLI (a direct ask from users of Claude Code).
- **Terminal UX Polish** – #3718, #3720, #3722: Visual delimiters between agentic iterations, stash half‑typed commands with `ESC ESC`, and fix disappearing multi‑line user input.

---

## Developer Pain Points

Recurring frustrations gathered from the last 24h of issue activity:

- **Windows‑specific woes:**  
  - Uninstallation fails via Control Panel (#3662).  
  - WSL startup delays make the tool nearly unusable (#3652).  
  - Copy‑on‑select in Windows Terminal is broken (#3724).  
  - Home‑directory path handling backslash/forward‑slash issues (#3719).  
  - Local sandbox fails on ReFS / Dev Drive (#3712).  
  - Registry version not updated after `/update` (#3711).

- **Inconsistent agent & tool behavior:**  
  - `preToolUse` hooks never fire (#2540).  
  - Agent‑level tool whitelists in `.agent.md` frontmatter are ignored (#2638).  
  - `--available-tools` / `--excluded-tools` flags silently ignored in ACP mode (#2948).  
  - Repository‑level custom agents vs. skills use different base directories (#3688).  
  - `/model` picker inconsistency (arrow‑keys vs. numeric input) (#3715).

- **BYOK / custom model friction:**  
  - No way to disable streaming even when the provider requires it (#3717).  
  - `/model` selector cannot list models from local BYOK providers (#3709).  
  - Function call regression with `$ref` in JSON schema (#3716).

- **Installation & platform detection:**  
  - Install script misidentifies FreeBSD as Windows (#3710).  
  - `GITHUB_TOKEN` support only just added (PR #1960) – previously rate‑limited installs were common.

---

*Generated from github.com/github/copilot-cli data as of 2026-06-09. Digest compiled by a technical analyst focused on AI developer tools.*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest – 2026-06-09

## 1. Today's Highlights
Only three issues saw updates in the past 24 hours. A critical installation regression (#2436) and a silent removal of API key authentication (#2442) are the main developer pain points. On the positive side, the team closed an enhancement that adds a deprecation banner to the Python-based docs, nudging users toward the TypeScript rewrite.

## 2. Releases
No new releases were published in the last 24 hours. The latest stable version remains **1.47.0**.

## 3. Hot Issues (3 items – all updated in the last 24h)
- **#2436 [bug] Installation failed – “Kimi can’t seem to make up her mind”**  
  Author: pleabargain | Created: 2026-06-07 | Updated: 2026-06-08 | Comments: 1  
  Running `kimi-cli -V` shows version 1.47.0 but installation itself fails with a confusing message. Community reaction: minimal engagement (0 👍), but the issue suggests a possible packaging or dependency resolution problem.  
  [GitHub Issue](https://github.com/MoonshotAI/kimi-cli/issues/2436)

- **#2442 [bug] Broken Workflow – API key authentication silently removed**  
  Author: andrew-sz | Created: 2026-06-08 | Updated: 2026-06-08 | Comments: 0  
  A major regression: under “Kimi Code” model 2.6 on macOS, API key authentication was removed without notice, breaking all CI/CD pipelines that rely on environment variable authentication. Zero comments yet, but the severity is high.  
  [GitHub Issue](https://github.com/MoonshotAI/kimi-cli/issues/2442)

- **#2376 [CLOSED] [enhancement] Add deprecation banner on GitHub Pages**  
  Author: MengyangGao | Created: 2026-05-27 | Updated: 2026-06-08 | Comments: 0  
  The Python-based `kimi-cli` documentation now shows a deprecation notice redirecting users to the TypeScript rewrite (`kimi-code`). This acknowledges the ongoing migration and reduces user confusion.  
  [GitHub Issue](https://github.com/MoonshotAI/kimi-cli/issues/2376)

## 4. Key PR Progress
No pull requests were updated in the last 24 hours. (0 items)

## 5. Feature Request Trends
The only trending direction visible from recent issues is **documentation / migration guidance** – specifically deprecating the Python CLI and directing users to the TypeScript rewrite. The closed enhancement #2376 is the clearest signal; expect more such cleanup tasks in the coming weeks.

## 6. Developer Pain Points
- **Installation reliability** (#2436): Users installing version 1.47.0 encounter a broken setup process, with an ambiguous error message that undermines trust.
- **Silent authentication changes** (#2442): The most concerning regression – API key authentication was removed from the “Kimi Code” model without any deprecation warning or changelog entry. This breaks existing automated workflows and forces urgent manual intervention.
- **General silence from maintainers**: Both open bugs have had zero maintainer responses in the last 24 hours, which may frustrate early adopters.

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-09

## Today’s Highlights
No new releases today, but the community saw a flurry of activity around long-standing bugs and feature requests. Three PRs were merged that improve MCP abort signals, paginate MCP catalogs, and add draft tab support to the desktop app. Meanwhile, a dozen issues saw updates, including a decade-old session view crash (#16494) and a high‑demand request for crypto payments (#23153).

## Releases
*None*

## Hot Issues (10 noteworthy)

1. **[#16494 – TypeError: Cannot read properties of undefined (reading 'length') in session view](https://github.com/anomalyco/opencode/issues/16494)**  
   *Open for 3 months, updated today* – A runtime crash in the session UI that blocks users. 10 comments, only 1 👍 suggests it may be niche but critical. Community suspects a race condition in rendering.

2. **[#23153 – [FEATURE] Pay Go with crypto](https://github.com/anomalyco/opencode/issues/23153)**  
   *15 👍, 7 comments* – Strong demand for cryptocurrency payment support for OpenCode Go. No official response yet; likely a business decision pending.

3. **[#31247 – Opus 4.8 via GitHub Copilot leaks repeated literal tool‑call text](https://github.com/anomalyco/opencode/issues/31247)**  
   *Created 2 days ago, 6 comments* – A bizarre bug where the assistant emits raw tool markup (e.g., `<invoke>`) into normal responses. High urgency as it breaks model trust.

4. **[#5558 – [FEATURE] Plugin‑Provided Autocomplete](https://github.com/anomalyco/opencode/issues/5558)**  
   *Closed but highly discussed, 6 👍* – The TUI’s autocomplete only supports built‑in items; the community wants plugins to register completions. This was closed without merge, leaving demand unaddressed.

5. **[#31441 – Folder and other nav buttons gone](https://github.com/anomalyco/opencode/issues/31441)**  
   *Closed today, 5 comments* – Regression in v1.14 where top‑menu navigation buttons disappeared. Quickly acknowledged and fixed; users report it was a UI refactor oversight.

6. **[#13430 – [FEATURE] Clickable file:line references in web UI](https://github.com/anomalyco/opencode/issues/13430)**  
   *5 comments, updated today* – Users want assistant‑provided file references to be clickable and jump to the exact line. A small UX improvement that would save significant manual scrolling.

7. **[#19988 – [FEATURE] Add `reasoning` as interleaved field option](https://github.com/anomalyco/opencode/issues/19988)**  
   *Closed, 4 comments* – vLLM renamed its reasoning field; this issue was resolved by PR #30477. Demonstrates the project’s responsiveness to upstream API changes.

8. **[#19255 – [FEATURE] Make GitHub releases immutable](https://github.com/anomalyco/opencode/issues/19255)**  
   *Closed, 4 comments* – A request to thwart supply‑chain attacks by disabling release edits. Gained 5 👍; aligns with industry best practices.

9. **[#19567 – [FEATURE] Durable event export / replay surface](https://github.com/anomalyco/opencode/issues/19567)**  
   *4 comments, updated today* – External integrations need a stable event stream beyond the current ephemeral sync event architecture. Still open, suggesting ongoing design work.

10. **[#16101 – [FEATURE] Session Lifecycle Management](https://github.com/anomalyco/opencode/issues/16101)**  
    *8 👍, 2 comments* – Users want auto‑archival, TTL, and storage caps for unbound session growth. A recurring pain point as sessions bloat in long‑running projects.

## Key PR Progress (10 important)

1. **[#30477 – feat: add "reasoning" as interleaved field option for vLLM providers](https://github.com/anomalyco/opencode/pull/30477)**  
   *Merged* – Closes #19988. Adds support for vLLM’s renamed `reasoning` field, ensuring seamless integration with the latest vLLM API.

2. **[#31454 – feat(app): tabs help button](https://github.com/anomalyco/opencode/pull/31454)**  
   *Merged* – Adds a contextual help button for the new tab system, currently shown only in dev builds. Prepares for user‑facing documentation.

3. **[#31452 – feat(app): draft prompt state](https://github.com/anomalyco/opencode/pull/31452)**  
   *Merged* – Part of decomposing the larger draft tab feature (#31034). Lays groundwork for unsaved prompt drafts in the desktop UI.

4. **[#31455 – fix(opencode): pass abort signal to MCP tool calls](https://github.com/anomalyco/opencode/pull/31455)**  
   *Merged* – Closes #24312. Forwards the AI‑SDK tool abort signal to MCP servers, allowing proper cancellation notifications when users stop a turn.

5. **[#31442 – fix(opencode): paginate MCP catalogs](https://github.com/anomalyco/opencode/pull/31442)**  
   *Merged* – Adds cursor‑based pagination for MCP tools, prompts, and resources. Ensures large catalogs (up to 1000 pages) are fully traversed without hang.

6. **[#28326 – feat(server): runtime base path support for reverse proxy deployments](https://github.com/anomalyco/opencode/pull/28326)**  
   *Open* – Adds `--base-path` flag and config key for running OpenCode behind a reverse proxy. Highly requested by container/self‑hosting users.

7. **[#30943 – fix(opencode): sanitize exported session path](https://github.com/anomalyco/opencode/pull/30943)**  
   *Open* – Closes #30709. Redacts the `info.path` field when using `--sanitize` export, improving privacy for shared logs.

8. **[#30658 – [contributor] feat(acp): emit plan session updates from todowrite tool calls](https://github.com/anomalyco/opencode/pull/30658)**  
   *Open* – Enables the Agent Communication Protocol (ACP) to render plans in OpenCode, matching Claude’s behavior. Bridges a feature gap for ACP users.

9. **[#31343 – feat(app): add draft tab support to tabs store](https://github.com/anomalyco/opencode/pull/31343)**  
   *Merged* – Part 1 of the draft tabs feature, introducing store support for draft states. Expected to improve multi‑session workflow.

10. **[#31297 – fix(shell): force UTF‑8 encoding for PowerShell output](https://github.com/anomalyco/opencode/pull/31297)**  
    *Open* – Closes #23636, #31187, #30205. Ensures non‑ASCII characters display correctly in PowerShell. A long‑standing pain point for Windows users.

## Feature Request Trends
The most requested feature directions observed across all issues today:
- **Payment flexibility** (#23153 – crypto payments) and **improved session export** (#19567, #30511, #31453 – subagent recursion, desktop export, durable event streams).
- **Desktop‑TUI parity** – users want `/export`, message search (`Cmd+F`), and navigational buttons ported to the desktop app (#19143, #31453, #31441).
- **MCP ecosystem** – better tool cancellation signals (#24312), pagination (#31442), and compatibility with Vercel MCP (#27761).
- **Storage management** – cache cleanup settings (#29037), session lifecycle TTL (#16101), and immutable releases (#19255).
- **Reasoning model support** – flexible reasoning field definitions (#19988) and OpenRouter variant discovery (#30216).

## Developer Pain Points
Recurring frustrations from the past 24 hours:
- **Patch failures caused by auto‑formatting** (#31422) – formatters that run after model edits corrupt the patch context, leading to `apply_patch` failures. Developers want formatter coordination.
- **JSON export truncation when piped** (#29330) – large sessions produce invalid JSON under pipe, forcing workarounds like writing to file first.
- **UI regressions on update** – users report missing navigation buttons (#31441), missing “Open project in IDE” (#31458), and composer dock rendering issues (#31448) after upgrading to v1.14/1.16.
- **Model cost display misleading** (#29971) – providers that omit the `cost` field cause OpenCode to show `$0`, tricking users into thinking a model is free.
- **Subagent hangs on bad files** (#22252) – researchers and task subagents deadlock on empty/corrupted PDFs, blocking batch pipelines.
- **Config directory loss after auto‑update** (#31447) – when the config directory is deleted, OpenCode crashes on startup due to missing `.gitignore` write.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-06-09

## Today’s Highlights
The `v0.79.0` release lands the much‑debated **Project Trust** system, which gates local inputs behind user prompts. Community feedback is mixed: a highly‑upvoted issue (#5514) calls the gating “annoying”. On the fix side, a cluster of PRs addresses several regressions (Azure OpenAI stateful mode, missing export assets) and a critical performance hotspot (quadratic session traversal). A new `alwaysTrust` setting and configurable clipboard image paths partially address user‑experience friction.

## Releases
- **[v0.79.0](https://github.com/earendil-works/pi/releases/tag/v0.79.0)**  
  Introduces **Project Trust** – Pi now asks before loading project‑local settings, resources, instructions, and packages. Saved decisions can be overridden with `--approve` / `--no-approve` for non‑interactive use.

## Hot Issues (10 noteworthy)

1. [#5514 – Project Trust Feature Feedback](https://github.com/earendil-works/pi/issues/5514)  
   **14 comments, 6 👍** | Open  
   The flagship feature of v0.79.0 is already controversial – the author finds the repeated trust prompts disruptive across machines. A follow‑up PR (#5515) adds an `alwaysTrust` opt‑out.

2. [#4180 – Links not clickable anymore](https://github.com/earendil-works/pi/issues/4180)  
   **10 comments** | Closed  
   After an update hyperlinks in the TUI became non‑interactive. Root cause suspected in the switch to alternate term mode.

3. [#5464 – Local models: 3‑5 minute “Working” status latency](https://github.com/earendil-works/pi/issues/5464)  
   **6 comments** | Closed  
   Mid‑session simple messages stall for minutes when using Ollama local models – a deal‑breaker for offline workflows.

4. [#5363 – Add amazon‑bedrock‑mantle provider](https://github.com/earendil-works/pi/issues/5363)  
   **6 comments, 3 👍** | Open  
   AWS Bedrock Mantle models use an OpenAI‑compatible API; the existing Bedrock provider (Converse) is incompatible. Community demand for native support.

5. [#5286 – Missing pricing info for GitHub Copilot models](https://github.com/earendil-works/pi/issues/5286)  
   **6 comments** | Closed  
   GitHub Copilot’s new per‑token pricing isn’t reflected in usage display – users see `$0.000 (sub)` instead of actual costs.

6. [#5407 – Double backspace and double enter on Kitty](https://github.com/earendil-works/pi/issues/5407)  
   **3 comments** | Closed  
   Key repeat issue specific to the Kitty terminal, absent in COSMIC Terminal. Points to underlying TTY handling differences.

7. [#5419 – Unable to change directory](https://github.com/earendil-works/pi/issues/5419)  
   **3 comments** | Closed  
   Working directory reset on every bash invocation prevents agents from building projects in subdirectories.

8. [#5531 – kimi.com: Thinking enabled despite `thinking off`](https://github.com/earendil-works/pi/issues/5531)  
   **3 comments** | Closed  
   The model still spends tokens thinking even when the user toggles thinking off – a logic bug in the provider.

9. [#5492 – High CPU in TUI on large sessions](https://github.com/earendil-works/pi/issues/5492)  
   **3 comments** | Closed  
   Quadratic branch traversal (`getBranch` called from `render`) caused 100% CPU on sessions with ~62k branches. Fixed in PR #5493.

10. [#5530 – `azure‑openai‑responses` missing `store: false`](https://github.com/earendil-works/pi/issues/5530)  
    **2 comments** | Open  
    Unlike the OpenAI‑native provider, Azure OpenAI runs in stateful mode, causing server‑side reasoning object drops. A three‑line fix is already submitted (#5524).

## Key PR Progress (10 important)

1. [#5539 – fix(minimax): default thinking on for MiniMax‑M3](https://github.com/earendil-works/pi/pull/5539)  
   Closed – Forces thinking for MiniMax‑M3’s Anthropic‑compatible endpoint; omitting the parameter yields empty responses.

2. [#5537 – feat(agent): add beforeModel hook and reactive compaction](https://github.com/earendil-works/pi/pull/5537)  
   Closed – Adds two new `AgentLoopConfig` callbacks: `beforeModel` for provider‑specific adjustments, and reactive compaction that stops the tool loop mid‑turn when the context window is exceeded.

3. [#5533 – fix(coding-agent): add missing template.{css,js} to copy‑binary‑assets](https://github.com/earendil-works/pi/pull/5533)  
   Closed – Resolves `pi --export` failures when running from the dist folder.

4. [#5524 – fix(ai): send store: false on Azure OpenAI Responses](https://github.com/earendil-works/pi/pull/5524)  
   Closed – Three‑line fix that switches Azure to stateless mode, preventing reasoning object corruption addressed in #5530.

5. [#5527 – fix(amazon-bedrock): extract region from inference profile ARNs](https://github.com/earendil-works/pi/pull/5527)  
   Open – Parses region directly from the ARN instead of relying on `AWS_REGION`, fixing cross‑region deployment scenarios.

6. [#5521 – feat(coding-agent): restore files on rewind (checkpoints)](https://github.com/earendil-works/pi/pull/5521)  
   Closed – Adds file‑level undo to conversation rewind: `Esc Esc` now prompts to restore edited files to the checkpoint state.

7. [#5515 – feat(coding-agent): add alwaysTrust setting](https://github.com/earendil-works/pi/pull/5515)  
   Closed – A direct response to #5514: a setting to permanently disable the trust gating pop‑up (disabled by default).

8. [#5518 – feat(coding-agent): configurable clipboard image storage path](https://github.com/earendil-works/pi/pull/5518)  
   Closed – Lets users set `images.storagePath` in `settings.json` instead of relying on the ephemeral `os.tmpdir()`.

9. [#5513 – Enforce context window mid‑turn via shouldStopAfterTurn](https://github.com/earendil-works/pi/pull/5513)  
   Closed – Exposes the `shouldStopAfterTurn` hook and wires it to auto‑compaction, stopping long tool loops before exceeding `contextWindow`.

10. [#5493 – Avoid quadratic session branch traversal](https://github.com/earendil-works/pi/pull/5493)  
    Closed – Fixes the high‑CPU issue #5492 by optimising `getBranch` in the session manager.

## Feature Request Trends
- **New model providers:** Multiple requests for Amazon Bedrock Mantle (#5363), Wafer (#5517), and Claude OAuth session usage (#5519) show a drive for broader, non‑API‑key integrations.
- **Configurable storage:** Users want control over where Pi stores clipboard images (#5520) and a clear separation between Pi‑managed data and user config files (#5508).
- **Trust model refinement:** The Project Trust feature spawned requests to expose trust decisions to extensions (#5523) and to provide an `alwaysTrust` opt‑out (already merged in #5515).
- **File‑based undo:** The ability to restore files on rewind (#5522) is a top‑requested quality‑of‑life improvement – now implemented in #5521.
- **Terminal theme detection:** PR #5385 (open) proposes auto‑detecting the terminal’s light/dark theme on first run.

## Developer Pain Points
- **Performance regressions:** Quadratic session traversal (#5492) and split‑turn compaction causing 429s on local backends (#5536) highlight ongoing scalability issues in large sessions.
- **Terminal compatibility:** Double keystrokes on Kitty (#5407) and Windows terminal popups from child processes (#5529) indicate fragile TTY integration.
- **Context window management:** Auto‑compaction lacks a mid‑turn guard (#5512) and fails with “context shift is disabled” errors (#5511) – the new `shouldStopAfterTurn` hook (#5513) aims to address this.
- **Model‑specific quirks:** MiniMax‑M2.5 removed from serverless (#5506), Gemini parallel tool calls cause 400 errors (#5528), and kimi.com ignores `thinking off` (#5531) – each requiring per‑provider workarounds.
- **Broken release notes & exports:** The v0.79.0 release notes contain dead links (#5516), and binary exports fail due to missing template assets (#5534) – a frustrating start for new users.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-06-09

## Today’s Highlights
A new preview release **v0.18.0-preview.0** landed with a critical fix for CLI copy output. The community is buzzing over the **implementation of Dynamic Workflows (P1)**, a **multi-tab extension manager**, and a highly requested **Plan Approval Gate**. On the bug front, subagent image-reading failures and model-provider disambiguation are top of mind.

## Releases
**v0.18.0-preview.0**  
- Fix: `cli` – skip “thought” parts in copy output ([#4742])  
- Release notes auto-generated from `.github/release.yml`  

[View release](https://github.com/QwenLM/qwen-code/releases/tag/v0.18.0-preview.0)

---

## Hot Issues

1. **#4747 [CLOSED] – Global user-level auto-memory**  
   Suggests `~/.qwen/memories/` to persist user preferences across projects (like Claude’s user memory). 4 comments, no upvotes. *Closed early, but the idea is still referenced in other proposals.*  
   [Issue](https://github.com/QwenLM/qwen-code/issues/4747)

2. **#4876 [OPEN] – Subagent reading images returns unrelated content**  
   `MachineXu` reports that subagents using `read_file` for images fail to describe the image, while the main agent works. Priority P2, 3 comments. **Key pain point for multi-agent workflows.**  
   [Issue](https://github.com/QwenLM/qwen-code/issues/4876)

3. **#2014 [CLOSED] – Structured error log output for external monitoring**  
   Feature request for standalone error logs (JSON/structured) for integration with observability tools. Closed after implementation? 3 comments.  
   [Issue](https://github.com/QwenLM/qwen-code/issues/2014)

4. **#4882 [OPEN] – `terminalSequence` field on hooks**  
   Ports Claude Code’s hook JSON extension to emit terminal-side effects (notifications, bells). Priority P3, 2 comments.  
   [Issue](https://github.com/QwenLM/qwen-code/issues/4882)

5. **#4877 [OPEN] – OpenWork cannot distinguish same model from different providers**  
   When multiple providers offer the same model ID (e.g., `glm-5`), the UI shows only one entry. Priority P2, 2 comments.  
   [Issue](https://github.com/QwenLM/qwen-code/issues/4877)

6. **#4883 [OPEN] – `--safe-mode` flag for troubleshooting**  
   Disables all user customizations to help diagnose issues. Priority P2, 1 comment.  
   [Issue](https://github.com/QwenLM/qwen-code/issues/4883)

7. **#4884 [OPEN] – Preserve CLI flags when resuming background agent sessions**  
   Currently only `resolvedApprovalMode` is persisted; other flags are lost on resume. Priority P2, 1 comment.  
   [Issue](https://github.com/QwenLM/qwen-code/issues/4884)

8. **#4879 [OPEN] – `/cd` command to change session working directory**  
   Interactive slash command to move to a new directory without restarting. Priority P2, 1 comment.  
   [Issue](https://github.com/QwenLM/qwen-code/issues/4879)

9. **#4872 [OPEN] – Automated CHANGELOG**  
   Add a `CHANGELOG.md` synced with releases (like Claude Code). Priority P3, 1 comment.  
   [Issue](https://github.com/QwenLM/qwen-code/issues/4872)

10. **#4721 [OPEN] – Port Dynamic Workflows / Ultracode from Claude Code 2.1.160**  
    Multi-agent third tier complementing `/swarm`. Priority unset, 1 comment. *Now being implemented in PR #4732.*  
    [Issue](https://github.com/QwenLM/qwen-code/issues/4721)

---

## Key PR Progress

1. **#4866 – Refactor CI: split PR triage into 4-job pipeline**  
   `yiliang114` replaces a monolithic skill with four parallel jobs (`product-decision`, `approval-decision`, etc.).  
   [PR](https://github.com/QwenLM/qwen-code/pull/4866)

2. **#4850 – Multi-tab `/extensions` dialog (Discover / Installed / Marketplaces)**  
   `BZ-D` upgrades the extension manager to a tabbed UI aligned with Claude Code’s `/plugin` command.  
   [PR](https://github.com/QwenLM/qwen-code/pull/4850)

3. **#4844 – Agent Team experimental feature**  
   `tanzhenxin` adds parallel sub-agent coordination with named teams, messaging, and task lists.  
   [PR](https://github.com/QwenLM/qwen-code/pull/4844)

4. **#4713 – Project `.mcp.json` + workspace approval gating**  
   `qqqys` introduces trust-based gating for checked-in MCP sources, aligning with Claude Code’s precedence model.  
   [PR](https://github.com/QwenLM/qwen-code/pull/4713)

5. **#4880 – Layered tool-output truncation (per-message budget, per-tool limits)**  
   `LaZzyMan` implements three-layer truncation mirroring Claude Code, with overflow to temp files.  
   [PR](https://github.com/QwenLM/qwen-code/pull/4880)

6. **#4874 – Web-shell mode indicator now mouse-selectable**  
   `wenshao` makes the approval-mode button clickable, opening the inline picker.  
   [PR](https://github.com/QwenLM/qwen-code/pull/4874)

7. **#4853 – `enter_plan_mode` tool and Plan Approval Gate**  
   `callmeYe` adds a tool for the model to self-lower into plan mode, plus a gate that runs plan summarization before execution.  
   [PR](https://github.com/QwenLM/qwen-code/pull/4853)

8. **#4881 – Auto-generated CHANGELOG.md from releases**  
   `LaZzyMan` delivers a CI-generated changelog (closes #4872).  
   [PR](https://github.com/QwenLM/qwen-code/pull/4881)

9. **#4732 – Workflow tool P1: `node:vm` sandbox + sequential `agent()`**  
   `LaZzyMan` starts the Dynamic Workflows port with a minimal JavaScript sandbox and sequential sub-agent calls (refs #4721).  
   [PR](https://github.com/QwenLM/qwen-code/pull/4732)

10. **#4842 – Declarative agent frontmatter v1 (Claude Code 2.1.168 parity)**  
    `LaZzyMan` bridges `permissionMode`, `maxTurns`, and color allowlist to qwen’s `approvalMode`.  
    [PR](https://github.com/QwenLM/qwen-code/pull/4842)

---

## Feature Request Trends
- **Cross-session & cross-project state**: Global user memory (`~/.qwen/memories/`), `--safe-mode` for clean troubleshooting, preserving CLI flags on resume, and `/cd` for working directory changes.
- **Claude Code parity**: `terminalSequence` on hooks, Dynamic Workflows / Ultracode, declarative agent frontmatter, and layered tool output truncation.
- **Developer tooling**: Automated CHANGELOG generation, structured error output for monitoring, and better model-provider disambiguation in OpenWork.

---

## Developer Pain Points
- **Subagent image handling broken** – Subagents cannot correctly read and describe images (#4876).  
- **Duplicate model entries cannot be distinguished** – Same model ID from different providers collapses into one UI entry (#4877).  
- **Background agent flags lost on resume** – Only approval mode is preserved, requiring reconfiguration (#4884).  
- **No safe mode for debugging** – Users want a flag to disable all customizations (#4883).  
- **Lack of interactive directory switching** – Exiting and restarting for a new working directory is cumbersome (#4879).  

These pain points indicate the community is pushing for more robust multi-agent, multi-session, and multi-provider support.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

Based on the provided GitHub data, here is the DeepSeek TUI (CodeWhale) community digest for 2026-06-09.

---

## 🐋 DeepSeek TUI (CodeWhale) Community Digest — 2026-06-09

This digest covers the open-source AI terminal interface now operating under the **CodeWhale** brand (formerly `deepseek-tui`). The primary repository is [Hmbown/CodeWhale](https://github.com/Hmbown/CodeWhale).

### 1. Today's Highlights

The project underwent a full brand migration to **CodeWhale**, with the v0.8.54 release solidifying the new name across commands, packages, and assets. The team is aggressively expanding its model ecosystem, with notable new additions planned for **Together AI**, **OpenAI Codex**, **NVIDIA Nemotron 3 Ultra**, and several others like **Qwen 3.7 Max** and **MiniMax 2.7**. While feature velocity is high, the community is facing significant pain points from the rebranding, including broken update paths (`npm` and `cargo`) and several configuration and provider compatibility bugs.

### 2. Releases

- **[v0.8.54](https://github.com/Hmbown/CodeWhale/releases/tag/v0.8.54)**: The "CodeWhale" brand is now the canonical project, command, npm package, and release-asset name. The legacy `deepseek-tui` npm package is officially deprecated and will receive no further updates. Installation now requires `cargo install codewhale-cli codewhale-tui --locked`.

### 3. Hot Issues (10 picks)

- **[#2924](https://github.com/Hmbown/CodeWhale/issues/2924) - [bug] I can't update to the new version using npm.** (Newly Created): A user reports a failed update via npm, a critical pain point caused by the recent rebranding. This is a high-priority regression for users who haven't migrated to the new `codewhale` packages.
- **[#2917](https://github.com/Hmbown/CodeWhale/issues/2917) - [bug] Cargo distribution: error: failed to spawn `codewhale' after changed from deepseek-tui**: Another instance of the brand migration causing user friction. The auto-update mechanism from `deepseek-tui` left users with a broken state, unable to find the new `codewhale` binary on their PATH.
- **[#2641](https://github.com/Hmbown/CodeWhale/issues/2641) - [bug] issue_pdf_bug**: A critical bug where the `read_file` tool for PDFs causes a channel close (and hangs) when no `pages` parameter is provided. The workaround is to always specify pages, but this represents a major functionality gap for the agent.
- **[#1327](https://github.com/Hmbown/CodeWhale/issues/1327) - [bug] FreeBSD x86_64: "Turn dispatch timed out; the engine may have stopped" on every prompt**: A persistent and long-standing issue for FreeBSD users. Every prompt operation fails, making the tool completely unusable on that platform.
- **[#2922](https://github.com/Hmbown/CodeWhale/issues/2922) - [question] agent会在执行操作前反复强调是YOLO模式，这个正常吗**: A user questions a potential UI/UX flaw where the AI agent repeats a "YOLO mode" confirmation before every single atomic action (e.g., editing a CSS file). This indicates a need to streamline agent verbosity.
- **[#2904](https://github.com/Hmbown/CodeWhale/issues/2904) - Feature request: Persistent agent state and signed compressed KV cache capsules**: A sophisticated feature request for long-running coding tasks. The user proposes persistent agent state and "KV cache capsules" to reduce cost and latency, showing a desire for more advanced, stateful agentic workflows.
- **[#2490](https://github.com/Hmbown/CodeWhale/issues/2490) - [bug] 不能编译UE工程 (Cannot compile Unreal Engine project)**: A user reports that the tool cannot compile an Unreal Engine project. This points to a more general issue with the tool's execution environment or shell tool integration for large, complex build systems.
- **[#2596](https://github.com/Hmbown/CodeWhale/issues/2596) - [bug, enhancement] TUI /model picker does not show custom models from other providers**: Users configuring custom models via `config.toml` for a non-primary provider find that the TUI's `/model` picker only shows the currently active provider's models, creating a confusing configuration experience.
- **[#2900](https://github.com/Hmbown/CodeWhale/issues/2900) - [bug] DSML调用错误 (DSML call error)**: A bug where the AI model sporadically outputs raw `dsml` (likely a task description language) text instead of executing it. This can fill the context window or cause erratic streaming behavior, severely disrupting agentic workflows.
- **[#2914](https://github.com/Hmbown/CodeWhale/issues/2914) - [bug, documentation, enhancement, migration, release-blocker] TUI: fix large-paste .deepseek writes and long status readability**: A release-blocker issue identified by the maintainer for v0.8.55. It addresses remaining UI roughness, including writing large pasted files to the legacy `.deepseek/` directory and poor readability of long status messages.

### 4. Key PR Progress (10 picks)

- **[#2916](https://github.com/Hmbown/CodeWhale/pull/2916) - v0.8.55 — Together AI provider + experimental OpenAI Codex (ChatGPT) provider**: The upcoming release candidate adding two major new provider integrations. The inclusion of the OpenAI Codex (ChatGPT) OAuth path is particularly significant for expanding user base.
- **[#2925](https://github.com/Hmbown/CodeWhale/pull/2925) - feat(provider): add dedicated Together AI support (#2906)**: A critical infrastructure PR that establishes Together AI as a first-class provider with its own config, auth, and model registry, instead of requiring users to hijack the generic OpenAI-compatible config.
- **[#2927](https://github.com/Hmbown/CodeWhale/pull/2927) - feat(model): add Qwen 3.7 Max to OpenRouter model catalog**: Shows the project's rapid response to new model releases. This PR adds the latest Qwen model with tool-call and reasoning support to the catalog resolver.
- **[#2930](https://github.com/Hmbown/CodeWhale/pull/2930) - feat(model): complete Qwen 3.6 Plus support with dedicated tests**: Complements the above PR by solidifying support for a previously added model with dedicated provider-hinted resolution tests.
- **[#2903](https://github.com/Hmbown/CodeWhale/pull/2903) - feat: build static linux x64 binaries with musl**: A significant DevOps improvement that builds fully static Linux binaries, eliminating runtime dependencies on glibc and dbus. This will dramatically improve portability and reduce link errors.
- **[#2753](https://github.com/Hmbown/CodeWhale/pull/2753) - feat(tui): multi-tab system with cross-tab collaboration**: A substantial feature PR introducing a tab manager with persistent chat history, session contexts, and a cross-tab `TaskDelegator`. This moves the TUI toward a more IDE-like workspace.
- **[#2882](https://github.com/Hmbown/CodeWhale/pull/2882) - fix: security bugs in execution policy** & **[#2884](https://github.com/Hmbown/CodeWhale/pull/2884) - fix: client response handling and desktop tray icon safety**: Two closely related PRs addressing 10 security and stability bugs, including a critical execution policy bypass via whitespace injection.
- **[#2920](https://github.com/Hmbown/CodeWhale/pull/2920) - fix(tui): write oversized paste files to .codewhale/pastes/**: A direct fix for the migration headache identified in issue #2914. This PR corrects the default paste location from the legacy `.deepseek/` to the new `.codewhale/` directory.
- **[#2482](https://github.com/Hmbown/CodeWhale/pull/2482) - [v0.9.0, whaleflow, workflow-runtime] feat: add WhaleFlow — declarative multi-agent workflow orchestration**: A merged high-level feature for a future release (v0.9.0) that introduces a JSON-configurable, multi-agent swarm orchestration system.

### 5. Feature Request Trends

- **Expansive Model Catalog (Top Trend)**: There is an overwhelming push for support of new models. Specific requests include [Qwen 3.7 Max](https://github.com/Hmbown/CodeWhale/issues/2907), [MiniMax 2.7](https://github.com/Hmbown/CodeWhale/issues/2910), [NVIDIA Nemotron 3 Ultra](https://github.com/Hmbown/CodeWhale/issues/2912), [Kimi K2.6](https://github.com/Hmbown/CodeWhale/issues/2909), and [Qwen 3.6 Plus](https://github.com/Hmbown/CodeWhale/issues/2908). The team is clearly acting on this, as many of these have dedicated PRs.
- **First-Class Provider Integrations**: Users want more than just OpenAI-compatible endpoints. The requests for dedicated providers like [Together AI](https://github.com/Hmbown/CodeWhale/issues/2906) and [OpenAI Codex/ChatGPT](https://github.com/Hmbown/CodeWhale/issues/2915) are a direct response to the limitations of generic configuration.
- **Advanced Agentic Persistence**: Feature requests are becoming more sophisticated. The [proposal for persistent agent state and KV cache capsules](https://github.com/Hmbown/CodeWhale/issues/2904) indicates users are pushing the tool beyond simple chat into long-running, complex development tasks, seeking cost efficiency and continuity.

### 6. Developer Pain Points

- **Brand Migration Confusion**: The most immediate and widespread pain point is the "DeepSeek TUI" to "CodeWhale" migration. Users are encountering [broken `cargo update` paths](https://github.com/Hmbown/CodeWhale/issues/2917) and [failed `npm` updates](https://github.com/Hmbown/CodeWhale/issues/2924), indicating a need for more robust migration documentation or tooling.
- **Platform Compatibility Stalls**: Long-standing issues like the [FreeBSD timeout bug](https://github.com/Hmbown/CodeWhale/issues/1327) remain open with no clear resolution in sight, causing frustration for developers on alternative OS platforms.
- **Agent Tooling Flakiness**: Several bugs highlight issues with the tool-calling internals. The [PDF reader hanging](https://github.com/Hmbown/CodeWhale/issues/2641), [DSML output issues](https://github.com/Hmbown/CodeWhale/issues/2900), and problems with [UE4 compilation](https://github.com/Hmbown/CodeWhale/issues/2490) suggest that the agent's file system and shell tool usage can be brittle.
- **Configuration Fragility**: Customizing the tool, especially with multiple providers, is proving difficult. The issue with the [/model picker not showing cross-provider custom models](https://github.com/Hmbown/CodeWhale/issues/2596) is a clear example of configuration logic not scaling with the tool's growing feature set.

</details>