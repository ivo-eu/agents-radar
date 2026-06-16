# AI CLI Tools Community Digest 2026-06-16

> Generated: 2026-06-16 05:20 UTC | Tools covered: 9

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

# AI CLI Developer Tools Ecosystem — Cross-Tool Comparison Report
**Date:** 2026-06-16

---

## 1. Ecosystem Overview

The AI CLI developer tools landscape continues to mature rapidly, with all seven major tools shipping updates within the past 72 hours and cumulative community engagement exceeding 1,200+ issue comments across tracked repositories. The dominant themes are **platform parity** (especially Windows/WSL support), **agent reliability** (stalls, false success reports, and memory leaks), and **security hardening** (SSRF protection, symlink traversal prevention, and credential management). A notable shift is the convergence toward MCP (Model Context Protocol) as the standard extension interface, with four tools actively working on compliance. Meanwhile, regional pricing demands—particularly India-specific billing—have emerged as the single highest-signal feature request across the entire ecosystem (Claude Code #17432 at +443 upvotes). Developer pain points cluster around **silent failures** (consumed credits without results), **TUI rendering regressions**, and **session continuity** bugs.

---

## 2. Activity Comparison

| Tool | Issues (Notable) | New PRs (24h) | Release Status | Community Signal |
|------|-----------------|---------------|----------------|-----------------|
| **Claude Code** | 10 hot issues, top at +443 👍 | 22 merged/active | v2.1.178 shipped | Highest upvote density; most active bug tracker |
| **OpenAI Codex** | 10 hot issues, top at +55 👍 | 10 significant PRs | v0.140.0 stable + alphas | Steady, moderate engagement |
| **Gemini CLI** | 10 hot issues, top at +8 👍 | 10 notable PRs | No release today | Quiet but security-critical work |
| **GitHub Copilot CLI** | 10 hot issues, top at +8 👍 | 1 PR (draft) | v1.0.63 + v1.0.63-0 | Low PR velocity; issues-driven |
| **Kimi Code CLI** | 4 issues (all updated) | 2 PRs open | No release today | Smallest community, early-stage |
| **OpenCode** | 10 hot issues, top at +84 👍 | 10 PRs open/closed | No release today | Strong feature-request engagement |
| **Pi** | 10 hot issues, top at +30 👍 | 10 PRs merged/open | v0.79.4 shipped | Active, broad provider expansion |
| **Qwen Code** | 10 hot issues, top at P1 security | 10 PRs merged/closed | v0.18.1 + desktop-v0.0.4 | High PR velocity; structured roadmap |
| **DeepSeek TUI** | 10 hot issues, top at 13 comments | 10 PRs open/closed | No release today | Growing quickly; sub-agent focus |

**Key takeaway:** Claude Code dominates community volume; Qwen Code and Pi have the highest PR throughput per day. OpenAI Codex and Copilot CLI show stability but lower iteration pace.

---

## 3. Shared Feature Directions

### Requirements appearing across 3+ tool communities:

| Theme | Tools | Specific Needs |
|-------|-------|----------------|
| **Persistent session goals** | Claude Code, OpenCode, DeepSeek TUI, Qwen Code | Long-lived objectives (`/goal`, `/loop`) that survive command switches and agent delegation |
| **Background execution** | Claude Code, OpenCode, Qwen Code | Non-blocking shell tasks (servers, watchers) while maintaining agent interactivity |
| **MCP ecosystem maturity** | Claude Code, OpenCode, Copilot CLI, Qwen Code | Full MCP spec compliance: resource/prompt/tool schemas, server instructions, credential brokering |
| **Agent reliability & stalls** | All seven tools | "Turn stalled", false success on max turns, infinite retry loops, sub-agent crashes |
| **Security hardening** | Claude Code, Gemini CLI, Copilot CLI, Pi | SSRF protection, symlink traversal prevention, credential brokering, provenance attestation |
| **Platform parity** | Claude Code, OpenAI Codex, Copilot CLI, Kimi Code | WSL2 native support, Windows encoding fixes, Wayland compatibility |
| **Observability & cost visibility** | OpenAI Codex, OpenCode, Qwen Code, Copilot CLI | Token/usage dashboards, token-per-second metrics, rate-limit transparency |
| **Permission & trust systems** | Claude Code, DeepSeek TUI, Qwen Code | Scoped persistent permission rules, opt-in tool access, enterprise policy controls |

### Pattern: The ecosystem is converging on **agent lifecycle management** as the next battleground—how to make autonomous agents reliable, auditable, and safe at scale.

---

## 4. Differentiation Analysis

| Tool | Core Differentiator | Target User | Technical Approach |
|------|--------------------|-------------|-------------------|
| **Claude Code** | Deep IDE integration + advanced plugin system | Professional developers on Windows/Mac/Linux | Fullscreen TUI with rich permission rules, MCP hooks, subagent orchestration |
| **OpenAI Codex** | apply_patch stability + code review workflows | Developers needing reliable code modification | Rust-based core with sandbox exec; `/goal` and `/usage` dashboards |
| **Gemini CLI** | AST-aware tools + SSRF-first security | Security-conscious teams, enterprise | Python-based with strict private-IP guards, Auto Memory, subagent health monitoring |
| **GitHub Copilot CLI** | GitHub ecosystem lock-in + MCP server hub | Enterprise GitHub users | Node/TypeScript; deeply integrated with GitHub Actions, OAuth, and enterprise policies |
| **Kimi Code CLI** | Lightweight, China-friendly | Developers in restricted networks | Minimal dependencies; proxy-aware design (still lacking); hook-based extensibility |
| **OpenCode** | Open-source parity with Claude Code | Open-source community, cost-sensitive | TypeScript/Node; strong MCP support; `/goal` as top feature request |
| **Pi** | Provider-agnostic + extension ecosystem | Power users, multi-provider setups | Modular provider registry; vim-like modal editor extension; first-run theme detection |
| **Qwen Code** | Self-paced background automation | Developers needing long-running agents | `/loop` wakeup primitives, dynamic workflows, multi-tab extension manager |
| **DeepSeek TUI** | Sub-agent architecture + global provider support | DeepSeek users, open-source enthusiasts | Rust-based with headless worker runtime; WeChat bridge; provider fallback chains |

**Key insight:** The tools are diverging along **three axes**: (a) security posture (Gemini CLI vs. feature-first Claude Code), (b) provider loyalty (OpenAI Codex vs. provider-agnostic Pi), and (c) automation depth (Qwen Code's `/loop` vs. DeepSeek TUI's sub-agent orchestration).

---

## 5. Community Momentum & Maturity

| Tool | Community Size | Iteration Pace | Maturity Stage |
|------|---------------|----------------|----------------|
| **Claude Code** | Very large (443 👍 on top issue) | Fast (daily releases) | Mature but regressing on platform parity |
| **OpenAI Codex** | Large (55 👍 on top issue) | Moderate (weekly releases) | Stable with persistent friction points |
| **Gemini CLI** | Medium (8 👍 on top issue) | Moderate (weekly PRs) | Security-critical, stable core |
| **GitHub Copilot CLI** | Medium (8 👍 on top issue) | Slow (3 PRs/week) | Mature but low innovation velocity |
| **Kimi Code CLI** | Small (2 comments on top issue) | Slow (2 PRs/day) | Early stage, gaps in fundamentals |
| **OpenCode** | Medium-large (84 👍 on top feature) | Moderate (10 PRs/day) | Growing fast, feature-rich |
| **Pi** | Medium (30 👍 on top issue) | Fast (10 PRs/day) | Mature, expanding provider support |
| **Qwen Code** | Medium (P1 security issue) | Very fast (10+ PRs/day) | Rapid iteration with structured roadmap |
| **DeepSeek TUI** | Medium (13 comments on stall bug) | Moderate (10 PRs/day) | Growing, sub-agent architecture evolving |

**Ranking by community engagement:** Claude Code > OpenCode > Pi > Qwen Code > OpenAI Codex > DeepSeek TUI > Gemini CLI > GitHub Copilot CLI > Kimi Code CLI.

**Ranking by iteration velocity:** Qwen Code (fastest) > Claude Code > Pi > DeepSeek TUI > OpenCode > OpenAI Codex > Gemini CLI > GitHub Copilot CLI > Kimi Code CLI.

---

## 6. Trend Signals

### 🟢 Positive Signals (Industry Opportunities)
- **Agent reliability is the #1 pain point**—tools that solve stalls, false success, and infinite retries will win developer trust.
- **Regional pricing pressure** (India INR, China proxy support) signals global market expansion—tools ignoring this risk market share loss.
- **MCP as universal standard** is accelerating—the four tools investing in full spec compliance (Claude Code, OpenCode, Copilot CLI, Qwen Code) will shape the extension ecosystem.
- **Background automation** (self-paced loops, cron-like agents) is the next frontier beyond single-session chat.

### 🟡 Cautionary Signals (Developer Frustration)
- **Silent failures and credit consumption** without results (Claude Code #62466, Copilot CLI #3814, Qwen Code #5180) erode trust—this is a recurring pattern across 5+ tools.
- **Platform fragmentation is worsening**—Windows/WSL users encounter unique regressions in every tool, suggesting insufficient cross-platform testing investment.
- **Multi-agent orchestration complexity** is underestimated—sub-agent crashes, deadlocks, and false success reports appear in DeepSeek TUI, Qwen Code, and Gemini CLI simultaneously.

### 🔴 Risk Signals (Strategic)
- **Supply-chain security demands are rising**—Pi (#5739) and Claude Code (#68689) both address provenance and symlink attacks; tools without security posture may face enterprise adoption barriers.
- **Observability gap**—only OpenAI Codex and Qwen Code are shipping usage dashboards; developers are demanding cost/performance metrics before committing to a tool.
- **The "apply_patch stability" problem** spans OpenAI Codex, Qwen Code, and Claude Code—this is the core value proposition of AI CLI tools, and it remains fragile across the board.

### For Technical Decision-Makers
- **If security is paramount:** Start with Gemini CLI (SSRF protection, AST-aware safety) or Claude Code (symlink defense, permission rules)
- **If provider flexibility is critical:** Pi or OpenCode offer the widest provider support
- **If deep GitHub integration is required:** Copilot CLI is the natural choice despite lower iteration velocity
- **If long-running autonomous agents are needed:** Qwen Code's `/loop` and DeepSeek TUI's sub-agent architecture lead the way
- **If you're in a restricted network:** Kimi Code CLI is the only early-stage option—but lacks proxy support, a critical gap to watch

**Bottom line:** The ecosystem is healthy and competitive, but no single tool has solved the reliability-safety-flexibility triangle. Expect consolidation around MCP compliance, session lifecycle management, and platform parity over the next 6–12 months.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
*Data snapshot: 2026-06-16 | Source: github.com/anthropics/skills*

---

## 1. Top Skills Ranking

The following PRs have attracted the most community discussion and attention:

**#514 – Document Typography Skill** ([PR #514](https://github.com/anthropics/skills/pull/514))
*Functionality:* Prevents orphan word wrap, widow paragraphs, and numbering misalignment in AI-generated documents — issues that affect virtually every document Claude produces.
*Status:* Open. Discussion centers on whether typographic rules should be a single skill or split into separate formatting concerns. High enthusiasm due to universal applicability.

**#486 – ODT Skill (OpenDocument Text)** ([PR #486](https://github.com/anthropics/skills/pull/486))
*Functionality:* Create, fill, read, and convert OpenDocument Format files (.odt, .ods), including template filling and ODT-to-HTML parsing. Triggered by any mention of LibreOffice or OpenDocument.
*Status:* Open. Discussion highlights the need for broader office-format support beyond DOCX.

**#83 – Skill-Quality-Analyzer & Skill-Security-Analyzer** ([PR #83](https://github.com/anthropics/skills/pull/83))
*Functionality:* Meta-skills that evaluate other skills across five quality dimensions (structure, documentation, security) — essentially a linter and security scanner for the Skills ecosystem itself.
*Status:* Open. Significant interest from contributors building skill tooling. Debate on whether this should be integrated into the skill-creator pipeline.

**#210 – Improve Frontend-Design Skill Clarity** ([PR #210](https://github.com/anthropics/skills/pull/210))
*Functionality:* Major revision of the frontend-design skill to ensure every instruction is actionable within a single conversation, with specific behavioral guidance.
*Status:* Open. Represents a template for how existing skills should be audited and improved.

**#1140 – Agent-Creator Meta-Skill** ([PR #1140](https://github.com/anthropics/skills/pull/1140))
*Functionality:* Enables task-specific agent set creation, with critical fixes to multi-tool evaluation and Windows path support. Addresses Issue #1120.
*Status:* Open. High-priority because it unlocks composable agent workflows.

**#538, #539, #541, #362, #1298 – Skill-Creator & Bug Fix Cluster**
*Functionality:* A series of PRs fixing case-sensitivity on Linux (PR #538), YAML parsing failures with unquoted descriptions (PR #539, #361), DOCX bookmark ID collisions (PR #541), UTF-8 multi-byte character panics (PR #362), and the critical `run_eval.py` 0% recall bug (PR #1298).
*Status:* Open. These form a continuous quality-improvement wave for the skill-creator tooling itself.

---

## 2. Community Demand Trends

From the 13 most-commented Issues, clear demand patterns emerge:

**Organizational & Sharing Infrastructure (Issue #228, 14 comments, 👍7):** The #1 requested feature is org-wide skill sharing. Users currently must download `.skill` files and redistribute manually via Slack/Teams. A shared skill library or direct sharing link is the most-anticipated capability.

**Tooling Reliability (Issue #556, 12 comments, 👍7; Issue #1169, 3 comments):** The `run_eval.py` 0% recall bug is the most impactful technical blocker — the skill-description optimization loop is effectively optimizing against noise. Multiple independent reproductions confirm this is a systemic issue, not configuration error.

**Security & Trust Boundaries (Issue #492, 7 comments, 👍2):** Concern that community skills distributed under the `anthropic/` namespace create trust boundary abuse vulnerabilities. Users may grant elevated permissions to skills they believe are official. A namespace governance solution is needed.

**Windows Compatibility (Issue #1061, 3 comments):** Three distinct blocking issues for native Windows users: `PATHEXT` handling for `claude.cmd`, `cp1252` encoding, and `select` on pipes. Multiple PRs (#1099, #1050, #1298) attempt to address this.

**Framework Integration (Issue #16, Issue #29):** Persistent demand for exposing Skills as MCPs (Model Context Protocol) and compatibility with AWS Bedrock. The MCP proposal suggests Skills could become callable APIs with structured parameters.

**Duplicate Skill Proliferation (Issue #189, 6 comments, 👍9):** When installing both `document-skills` and `example-skills` plugins, identical content creates duplicate skills in Claude's context window. Community wants clear separation of concerns.

---

## 3. High-Potential Pending Skills

These PRs have active discussion and may land in the near term:

**#1298 – Fix run_eval.py 0% recall** ([PR #1298](https://github.com/anthropics/skills/pull/1298))
Root-cause fix for the evaluation pipeline's silent failure. Installs the eval artifact as a real skill, fixes Windows stream reading, trigger detection, and parallel workers. Directly unblocks all description-optimization workflows.

**#723 – Testing Patterns Skill** ([PR #723](https://github.com/anthropics/skills/pull/723))
Comprehensive testing stack coverage: Testing Trophy model, AAA pattern, React component testing, mocking strategies, and E2E testing. Addresses a clear gap in the existing collection.

**#444 – AURELION Skill Suite** ([PR #444](https://github.com/anthropics/skills/pull/444))
Four skills (kernel, advisor, agent, memory) forming a structured cognitive framework for professional knowledge management. Significant interest from users building layered AI collaboration systems.

**#154 – Shodh-Memory Skill** ([PR #154](https://github.com/anthropics/skills/pull/154))
Persistent memory system for AI agents across conversations. Teaches Claude when to call proactive context and how to structure rich memory entries. Addresses a fundamental limitation of stateless conversation models.

**#147 – Codebase Inventory Audit Skill** ([PR #147](https://github.com/anthropics/skills/pull/147))
Systematic 10-step workflow for identifying orphaned code, unused files, and documentation gaps. Produces a single-source-of-truth CODEBASE-STATUS.md. Appeals to teams managing legacy code.

**#181 – SAP-RPT-1-OSS Predictor** ([PR #181](https://github.com/anthropics/skills/pull/181))
Skills wrapper for SAP's open-source tabular foundation model for predictive analytics on business data. Niche but strategically valuable for enterprise users.

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is for reliable skill-creator tooling infrastructure (fixing evaluation pipelines, YAML parsing, and Windows compatibility) rather than new skill functionality — the ecosystem's growth is currently bottlenecked by its own tooling reliability.**

---

Here is the **Claude Code Community Digest** for **2026-06-16**.

---

## Claude Code Community Digest — 2026-06-16

### Today's Highlights

The community is buzzing about **India-specific pricing** (the single most-upvoted issue on the repo at +443) and **multi-session collaboration features**. In the last 24 hours, the team shipped **v2.1.178** with powerful new permission-rule syntax (`Tool(param:value)`) that lets users block specific models or tool parameters with wildcards. Meanwhile, a firehose of 22 pull requests—mostly workflow and plugin fixes from a prolific contributor—suggests Anthropic is actively hardening the Windows and macOS TUI after a wave of regressions.

### Releases

**v2.1.178** (latest)
- Added `Tool(param:value)` syntax for permission rules to match a tool’s input parameters, including `*` wildcards. Example: `Agent(model:opus)` blocks Opus subagents.
- Skills placed in nested `.claude/skills` directories now load correctly when the working directory is inside that subtree. On name clashes, the nested skill wins.

### Hot Issues

1. **[#17432 — India-Specific Pricing Plans (INR)](https://github.com/anthropics/claude-code/issues/17432)**
   195 comments, +443 👍. The community is loudly requesting local-currency billing, noting both OpenAI and Google already offer INR pricing. This is the most-requested feature in the entire tracker.

2. **[#62466 — Repeated “Image couldn’t be processed” API errors](https://github.com/anthropics/claude-code/issues/62466)**
   26 comments, +20 👍. Users report that failed image processing consumes their usage limit without delivering results. A costly and frustrating UX bug.

3. **[#38788 — Claude Code broken on WSL1 since v2.1.83](https://github.com/anthropics/claude-code/issues/38788)**
   23 comments, +4 👍. A long-standing regression for WSL1 users that has now been open for nearly three months. The "duplicate" label suggests it’s a known but unresolved issue.

4. **[#24285 — Can’t see Claude’s thinking anymore](https://github.com/anthropics/claude-code/issues/24285)**
   13 comments, +39 👍. The model’s chain-of-thought reasoning is no longer displayed in the TUI. Community sentiment is negative: users feel they’ve lost visibility into the agent’s decision process.

5. **[#49933 — Native WSL Remote Integration for Claude Desktop](https://github.com/anthropics/claude-code/issues/49933)**
   12 comments, +61 👍. Windows users want a seamless way to run Claude Code Desktop against WSL Linux projects without manual shell configuration.

6. **[#29355 — Allow programmatic session renaming](https://github.com/anthropics/claude-code/issues/29355)**
   9 comments, +65 👍. Users want hooks or tools to auto-rename sessions based on ticket IDs (e.g., `TICKET-123`). This is a strong signal for deeper workflow integration.

7. **[#63197 — Compaction fails with “context window limit” at 20% usage](https://github.com/anthropics/claude-code/issues/63197)**
   Closed today. A regression in v2.1.153 that was fixed in one of the latest patches—relevant for any developer hitting false-positive context limits.

8. **[#65585 — Auto-compact stopped working for third-party API providers](https://github.com/anthropics/claude-code/issues/65585)**
   5 comments, +3 👍. Users who route through Bedrock or Vertex report that auto-compaction silently broke since v2.1.161, reducing effective context window for non-Anthropic API users.

9. **[#68742 — Stale spinner rows + statusline desync under subagent load](https://github.com/anthropics/claude-code/issues/68742)**
   Filed today. TUI rendering corruption under parallel subagent workloads on WSL2—a sign that the fullscreen TUI has edge cases with concurrency.

10. **[#68584 — Desktop extension installs fail silently on macOS Tahoe](https://github.com/anthropics/claude-code/issues/68484)**
    Reported as `invalid`, but the lack of error feedback for .mcpb file installations is a developer experience issue for macOS 26.5 users.

### Key PR Progress

1. **[#68678 — Don’t mark Claude Desktop issues as invalid](https://github.com/anthropics/claude-code/pull/68678)**
   Fixes the triage bot that was incorrectly labeling cross-product bugs. A small but important process improvement.

2. **[#68707 — Add `/bug` command to file GitHub issues from terminal](https://github.com/anthropics/claude-code/pull/68707)**
   A new slash command that collects environment snapshots and submits bug reports entirely inside the TUI—no more browser tab switching.

3. **[#68672 — Load only `event:all` rules for unknown tools in Hookify](https://github.com/anthropics/claude-code/pull/68672)**
   Fixes Hookify’s rule engine so that unrecognized tool names don’t cause silent evaluation failures.

4. **[#68671 — PostToolUse hooks can now return `permissionDecision: deny`](https://github.com/anthropics/claude-code/pull/68671)**
   Previously, post-use deny was ignored. This fix allows security policies to revoke permissions retroactively.

5. **[#68681 — Correct pagination break condition in workflow scripts](https://github.com/anthropics/claude-code/pull/68681)**
   Fixes infinite loops in GitHub Actions pagination: the break condition now checks for `< 100` items instead of `=== 0`.

6. **[#68699 — Add Python wrapper for Hookify on Windows](https://github.com/anthropics/claude-code/pull/68699)**
   Resolves `python3` stub resolution and backslash path mangling—a major Windows UX fix for hook developers.

7. **[#68689 — Block symlink escape in security-guidance config reads](https://github.com/anthropics/claude-code/pull/68689)**
   Security fix that uses `realpath` + `startswith` to prevent symlink-based path traversal when reading `.claude/` configs.

8. **[#68700 — Normalize plugin root path for Windows in learning-output-style](https://github.com/anthropics/claude-code/pull/68700)**
   Ensures `SessionStart` hooks work on Windows by prefixing `bash` and converting backslashes.

9. **[#68693 — Add `duplicate` label additively, don’t replace existing labels](https://github.com/anthropics/claude-code/pull/68693)**
   Prevents the close-as-duplicate workflow from stripping existing labels like `bug` or `regression`.

10. **[#60427 — Use standard GitHub capitalization in README](https://github.com/anthropics/claude-code/pull/60427)**
    A cosmetic docs fix from the community, keeping the README consistent with GitHub’s brand guidelines.

### Feature Request Trends

Three high-signal themes:

- **Global Pricing & Accessibility** — Issue #17432 (+443) dwarfs all other feature requests. The Indian developer community is vocal about needing INR billing.
- **Collaborative Memory & Sessions** — Issues #38536 (Shared Team Memory) and #60082 (Multi-user real-time collaboration) show demand for team workflows beyond single-user chat.
- **Deeper IDE & Platform Integration** — Requests for WSL Remote support (#49933, +61), multi-selection in VSCode chat (#33058), and a configurable `workingDirectory` (#68738) indicate users want Claude Code to feel like a native IDE component, not a standalone CLI.

### Developer Pain Points

- **Platform-specific TUI regressions** — Windows/WSL users continue to hit UTF-8 corruption (#65394), stale spinners (#68742), and complete session non-responsiveness (#68739). macOS users face timezone heatmap bugs (#67625).
- **Broken auto-compaction with third-party APIs** — Reports (#65585) suggest that fixes for Anthropic’s API have inadvertently broken compaction for Bedrock/Vertex users, limiting the effective context window.
- **Silent failures** — Tool results being dropped (#68741), .mcpb installs dead-ending (#68484), and sessions disappearing from the resume list (#68717) erode trust in the tool’s reliability.
- **Rate-limit transparency** — The `/usage` command lacks detailed rate-limit breakdowns (#68744), making it hard for power users to diagnose throttling.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-06-16

## Today's Highlights
The team shipped **Codex v0.140.0** with new `/usage` token-activity dashboards and better handling of oversized content in `/goal`. On the bug front, the long-standing issue of Codex replying to earlier messages in conversations (Issue #8648) continues to generate community heat with 55 upvotes. Performance-focused PRs are landing to reduce resume overhead and cache plugin capabilities, while multiple `apply_patch` failures remain a top developer friction point.

---

## Releases

**rust-v0.140.0** – Stable release with notable features:
- `/usage` views for daily, weekly, and cumulative account token activity.
- `/goal` now preserves oversized text, large pasted blocks, and image attachments in remote sessions.
- Permanent session deletion support.

Several alpha releases (v0.141.0-alpha.1/.2, v0.140.0-alpha.20–22) were also tagged, indicating continued work on the next minor version.  
[Release v0.140.0](https://github.com/openai/codex/releases/tag/rust-v0.140.0)

---

## Hot Issues (10 Notable)

1. **Context ordering bug**  
   [Issue #8648](https://github.com/openai/codex/issues/8648) – Codex replies to earlier messages instead of the latest one in multi-message conversations. 61 comments, 55 👍. A widely reported and frustrating UX bug.

2. **`apply_patch` invoked as shell command**  
   [Issue #2235](https://github.com/openai/codex/issues/2235) – Models sometimes emit `apply_patch:` as a shell command leading to “command not found”. 17 comments, still unresolved after many months.

3. **Windows elevated spawns medium-integrity shell**  
   [Issue #25296](https://github.com/openai/codex/issues/25296) – Even with `--do-not-de-elevate`, the actual agent shell runs with reduced privileges. 8 comments, 2 👍.

4. **`apply_patch` fails with custom local models**  
   [Issue #17899](https://github.com/openai/codex/issues/17899) – Using a local OSS model via LM Studio breaks `apply_patch` tool calls. 7 comments, 3 👍.

5. **Auto-review rejects `codex exec` as untrusted**  
   [Issue #23211](https://github.com/openai/codex/issues/23211) – The auto-review feature considers `codex exec` an external service risk. 6 comments, 2 👍.

6. **Generic `apply_patch` bug**  
   [Issue #17517](https://github.com/openai/codex/issues/17517) – Another report of `apply_patch` tool malfunction, on Ubuntu with GPT-5.3. 6 comments, 4 👍.

7. **Archived chats: Delete button does nothing**  
   [Issue #28095](https://github.com/openai/codex/issues/28095) – Clicking “Delete” in archived chats has no effect. 5 comments, reported on macOS.

8. **Playwright browser tests fail in macOS sandbox**  
   [Issue #14368](https://github.com/openai/codex/issues/14368) – Sandbox prevents Chromium startup due to Mach port and network restrictions. 2 comments but 7 👍, affecting test automation.

9. **504 timeout on long code completions**  
   [Issue #28447](https://github.com/openai/codex/issues/28447) – On Windows, generating completions over ~2000 tokens causes a 10-second hang and 504 error. New, 2 comments.

10. **Windows desktop app crashes on launch**  
    [Issue #28442](https://github.com/openai/codex/issues/28442) – Version 26.609.9530.0 fails to open any window on Windows 10 Pro. New, 2 comments.

---

## Key PR Progress (10 Important)

1. **Reduce resume and fork orchestration overhead**  
   [PR #28456](https://github.com/openai/codex/pull/28456) – Isolates app-server and TUI changes to avoid reloading history on resume/fork. Performance-critical for long sessions.

2. **Repair stale and custom rollout paths**  
   [PR #28455](https://github.com/openai/codex/pull/28455) – Fixes SQLite metadata corruption for rollout paths that become stale, improving stability for branched threads.

3. **Add local credential broker**  
   [PR #28034](https://github.com/openai/codex/pull/28034) – Introduces a credential broker that keeps GitHub/OpenAI secrets inside a managed proxy, injecting them only on MITM-matched requests. Security enhancement.

4. **Speed up rollout metadata and history reads**  
   [PR #28452](https://github.com/openai/codex/pull/28452) – Uses buffered reads, direct deserialization, and bounded allocation to reduce latency when loading session metadata.

5. **Record external agent import results**  
   [PR #28396](https://github.com/openai/codex/pull/28396) – Persists import success/failure details for external agent configs, enabling robust rollback and progress tracking.

6. **Render remote environment cwd natively**  
   [PR #28152](https://github.com/openai/codex/pull/28152) – Fixes cross-platform path rendering (e.g., `/C:/windows`) when app-server and exec environment use different OSes.

7. **Cache discoverable plugin capabilities**  
   [PR #27812](https://github.com/openai/codex/pull/27812) – Reuses pre-computed `PluginCatalogSnapshot` to avoid re-reading plugin capability files on every sampling request. Large performance gain.

8. **Add interruptible sleep tool**  
   [PR #28429](https://github.com/openai/codex/pull/28429) – Adds a built-in `sleep` tool behind a feature flag, enabling model pauses without tying up a shell process.

9. **Queue TUI follow-ups through app-server**  
   [PR #28307](https://github.com/openai/codex/pull/28307) – Enables plain follow-ups to survive TUI process restarts by persisting them in the app-server’s idle message queue.

10. **Expose Bedrock credential source in account/read**  
    [PR #27751](https://github.com/openai/codex/pull/27751) – Lets the UI distinguish between Codex-managed and user-supplied AWS Bedrock credentials, improving account state rendering.

---

## Feature Request Trends

- **Observability & Metrics** – Multiple requests for AI LOC (Lines of Code) metrics exported via OpenTelemetry ([#28449](https://github.com/openai/codex/issues/28449)), and surfacing non-content stream status (e.g., capacity waits) in the TUI ([#28445](https://github.com/openai/codex/issues/28445)).
- **Native IDE Integration** – Strong interest in deeper VS Code integration (chat, in-editor completions) to replace GitHub Copilot workflows ([#26906](https://github.com/openai/codex/issues/26906)).
- **MCP / Long-Horizon State** – The “OpenTTT” proposal for cryptographic event DAGs to prevent context compression amnesia in long agent workflows ([#25210](https://github.com/openai/codex/issues/25210)) signals demand for better persistent memory.
- **Credential & Plugin Management** – Users want secure local credential brokering (PR #28034), per-plugin enable/disable controls that actually work ([#28443](https://github.com/openai/codex/issues/28443)), and managed cloud provider login flows (PR #28148).

---

## Developer Pain Points

- **`apply_patch` instability** – At least four distinct issues ([#2235](https://github.com/openai/codex/issues/2235), [#17899](https://github.com/openai/codex/issues/17899), [#17517](https://github.com/openai/codex/issues/17517), [#17969](https://github.com/openai/codex/issues/17969)) describe the tool failing when invoked directly, as a shell command, inside sandboxes, or with custom models. This is the single biggest source of user frustration.
- **Windows integrity & elevation** – A cluster of bugs ([#25296](https://github.com/openai/codex/issues/25296), [#28107](https://github.com/openai/codex/issues/28107), [#28442](https://github.com/openai/codex/issues/28442)) reveal inconsistent behavior when launching Codex Desktop elevated: agent shells run at medium integrity, app-server de-elevates, and the latest update crashes on launch.
- **Sandbox incompatibilities** – Playwright tests fail on macOS ([#14368](https://github.com/openai/codex/issues/14368)), `apply_patch` fails inside bubblewrap on Ubuntu ([#17969](https://github.com/openai/codex/issues/17969)), and auto-review rejects `codex exec` as untrusted ([#23211](https://github.com/openai/codex/issues/23211)).
- **Timeout & reliability** – 504 errors on long completions ([#28447](https://github.com/openai/codex/issues/28447), [#28446](https://github.com/openai/codex/issues/28446)) and remote SSH session timeouts ([#23312](https://github.com/openai/codex/issues/23312)) indicate network resilience gaps.
- **Configuration failures** – Winget alias not created ([#28321](https://github.com/openai/codex/issues/28321)), cron automations not triggering ([#28444](https://github.com/openai/codex/issues/28444)), and archived delete button broken ([#28095](https://github.com/openai/codex/issues/28095)) show edge cases in fundamental features.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-06-16

## Today's Highlights
A quiet day on the release front, but significant progress on S**SRF protection** and **execution stability** across several open PRs. Two high-priority bugs remain hot: the **generalist agent hang** (Issue #21409, 8 👍) and a **subagent recovery false success** (Issue #22323), both actively discussed by the community with 7+ comments each. The **Auto Memory** and **AST-aware tools** workstreams continue to attract developer attention.

---

## Releases
No new versions published in the last 24 hours.

---

## Hot Issues (10 noteworthy)

1. **#21409 – Generalist agent hangs**  
   `[priority/p1, area/agent, kind/bug]`  
   The agent hangs forever when deferring to the generalist subagent for simple tasks (e.g., folder creation). Community workaround: instruct the model not to use sub-agents. High impact with 8 👍.  
   🔗 https://github.com/google-gemini/gemini-cli/issues/21409

2. **#22323 – Subagent recovery after MAX_TURNS reported as GOAL success**  
   `[priority/p1, area/agent, kind/bug]`  
   `codebase_investigator` subagent reports `status: "success"` even after hitting the max turn limit with zero analysis done. Masks real failure from users.  
   🔗 https://github.com/google-gemini/gemini-cli/issues/22323

3. **#22745 – Assess impact of AST-aware file reads, search, and mapping**  
   `[priority/p2, area/agent, kind/feature]`  
   Epic tracking whether AST-aware tools can reduce token use, improve read precision, and reduce agent turns. 1 👍 from community.  
   🔗 https://github.com/google-gemini/gemini-cli/issues/22745

4. **#25166 – Shell command execution stuck with "Waiting input" after command completes**  
   `[priority/p1, area/core, kind/bug, effort/medium]`  
   After executing simple CLI commands (no prompts), the shell remains marked as active/awaiting input. 3 👍 — a clear reliability regression.  
   🔗 https://github.com/google-gemini/gemini-cli/issues/25166

5. **#24353 – Robust component level evaluations**  
   `[priority/p1, area/agent, kind/customer-issue]`  
   Follow-up EPIC for behavioral evals. Now has 76 tests across 6 Gemini models. Critical for ensuring eval reliability.  
   🔗 https://github.com/google-gemini/gemini-cli/issues/24353

6. **#26525 – Add deterministic redaction and reduce Auto Memory logging**  
   `[priority/p2, area/security, kind/bug]`  
   Auto Memory sends transcripts to model context before redacting secrets. Also logs skill content. Privacy/security concern flagged by community.  
   🔗 https://github.com/google-gemini/gemini-cli/issues/26525

7. **#26522 – Stop Auto Memory from retrying low-signal sessions indefinitely**  
   `[priority/p2, area/agent, kind/bug]`  
   Sessions that look low-signal remain unprocessed and are infinitely re-surfaced. Blocks progress on memory quality.  
   🔗 https://github.com/google-gemini/gemini-cli/issues/26522

8. **#21968 – Gemini does not use skills and sub-agents enough**  
   `[priority/p2, area/agent, kind/bug]`  
   Custom skills (e.g., "gradle", "git") are ignored unless the user explicitly instructs the model to use them. Impacts productivity for power users.  
   🔗 https://github.com/google-gemini/gemini-cli/issues/21968

9. **#22672 – Agent should stop/discourage destructive behavior**  
   `[priority/p2, area/agent, kind/customer-issue]`  
   The model occasionally uses unsafe git operations (`git reset`, `--force`) when safer alternatives exist. 1 👍.  
   🔗 https://github.com/google-gemini/gemini-cli/issues/22672

10. **#21983 – Browser subagent fails in Wayland**  
    `[priority/p1, area/agent, kind/bug, agent/browser]`  
    The browser subagent fails silently on Wayland with `Termination Reason: GOAL` but no output. Blocks Linux users with modern display servers.  
    🔗 https://github.com/google-gemini/gemini-cli/issues/21983

---

## Key PR Progress (10 important)

1. **#27956 (Open) – Support GDC air-gapped Service Identity**  
   `[area/security, size/m]`  
   Enables token exchange for Google Distributed Cloud Hosted air-gapped environments after the auth library update to v10.7.0.  
   🔗 https://github.com/google-gemini/gemini-cli/pull/27956

2. **#27744 (Open) – Resolve DNS before SSRF guard to block hostname-to-private-IP bypass**  
   `[size/m, status/pr-nudge-sent]`  
   Fixes a gap where hostnames like `127.0.0.1.nip.io` bypass `isPrivateIp()` because it only checks IP literals.  
   🔗 https://github.com/google-gemini/gemini-cli/pull/27744

3. **#27739 (Open) – Prevent SSRF via DNS hostnames and redirects**  
   `[size/m, status/need-issue]`  
   Companion to #27744; adds redirect-following inspection to block SSRF attacks through DNS rebinding and redirect chains.  
   🔗 https://github.com/google-gemini/gemini-cli/pull/27739

4. **#27626 (Closed) – Block private OAuth metadata URLs**  
   `[priority/p2, area/security, size/m]`  
   Adds SSRF protection to MCP OAuth metadata discovery – prevents fetching metadata from private network endpoints. Merged.  
   🔗 https://github.com/google-gemini/gemini-cli/pull/27626

5. **#27572 (Closed) – Fix tmux false positive background detection**  
   `[size/m, status/need-issue]`  
   Fixes regression where Gemini incorrectly detects light background (`#ffffff`) inside tmux+mosh, causing inappropriate theme switching.  
   🔗 https://github.com/google-gemini/gemini-cli/pull/27572

6. **#27854 (Closed) – Fix/pending tools and trust overrides**  
   `[size/m, status/need-issue]`  
   Prevents premature state progression when waiting for user tool approval. Forces file writes to execute sequentially to avoid race conditions.  
   🔗 https://github.com/google-gemini/gemini-cli/pull/27854

7. **#27943 (Open) – Resolve defensive path resolution for at-reference files**  
   `[size/m, status/need-issue]`  
   Fixes `File not found` errors when tools (`read_file`, `replace`, `write_file`) encounter files referenced via `@` mention syntax.  
   🔗 https://github.com/google-gemini/gemini-cli/pull/27943

8. **#27948 (Open) – Pin dependencies and enforce 14-day update cooldown**  
   `[size/xl, status/need-issue]`  
   Strictly pins all direct dependencies to exact versions and enforces a 14-day cooldown for automated updates. Reduces supply-chain risk.  
   🔗 https://github.com/google-gemini/gemini-cli/pull/27948

9. **#27603 (Closed) – Add platform-aware shell guidance**  
   `[priority/p3, area/agent, size/m]`  
   Updates the operational prompt to show Windows-specific inspection commands instead of Unix-only examples on `win32`. Merged.  
   🔗 https://github.com/google-gemini/gemini-cli/pull/27603

10. **#27947 (Open) – Migrate coreTools setting to tools.core**  
    `[size/m, status/need-issue]`  
    Moves from deprecated `coreTools` property to the nested `tools.core` schema across workflows and A2A server configuration.  
    🔗 https://github.com/google-gemini/gemini-cli/pull/27947

---

## Feature Request Trends

- **AST-aware tooling** – Multiple issues (#22745, #22746, #22747) request AST-based file reads, search, and codebase mapping to reduce token waste and improve precision.  
- **Agent self-awareness** – Community wants the agent to understand its own CLI flags, hotkeys, and available sub-agents so it can guide users without external docs (#21432).  
- **Browser agent resilience** – Improvements to session takeover, lock recovery, and persistent profile handling are frequently requested (#22232).  
- **Remote & background agents** – Epic-level work continues on remote agents with advanced auth and background operations (#20303).  
- **Memory system maturity** – Auto Memory needs deterministic secret redaction (#26525), infinite retry protection (#26522), and invalid patch quarantining (#26523).

---

## Developer Pain Points

- **Agent hangs & false success** – The generalist agent hangs on simple tasks (#21409), while subagents sometimes report success after hitting turn limits (#22323). Erodes trust in autonomous mode.  
- **Shell execution stalling** – Simple commands leave the shell in "Waiting input" state indefinitely (#25166). Very disruptive to daily workflow.  
- **Subagent underutilization** – Custom skills and sub-agents are ignored unless explicitly requested (#21968). Undermines the extensibility promise.  
- **Configuration & permission leaks** – Subagents run without permission after v0.33.0 (#22093); browser agent ignores `settings.json` overrides (#22267).  
- **Destructive behavior** – The model occasionally uses unsafe git commands (`--force`, `git reset`) without user awareness (#22672).  
- **Wayland & terminal compatibility** – Browser subagent fails on Wayland (#21983); terminal resize causes high-performance flicker (#21924).

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest – 2026-06-16

## Today’s Highlights
Two patch releases (v1.0.63 and v1.0.63-0) landed yesterday, improving image-attachment error messages, sorting `--help` options alphabetically, and adding a `deferTools` MCP option. Meanwhile, the community is grappling with several regressions: the `userPromptSubmitted` hook broke in v1.0.60, MCP server respawning loops appeared in v1.0.61, and session‑wedging bugs persist. A notable new feature request asks for prompt caching support for Claude Sonnet to reduce latency and costs.

## Releases
- **[v1.0.63](https://github.com/github/copilot-cli/releases/tag/v1.0.63)** (2026-06-15)  
  - Blocked image attachments now show a clear explanation and actionable steps (enable vision, switch model, or use a different image) instead of a confusing error.  
  - Options in `--help` output are now sorted alphabetically.

- **[v1.0.63-0](https://github.com/github/copilot-cli/releases/tag/v1.0.63-0)** (2026-06-15)  
  - **Added:** Press `w` in `/diff` to hide whitespace‑only changes.  
  - **Added:** `deferTools` option in MCP server config to keep a server’s tools always available even when tool search is enabled.  
  - **Improved:** Reliability of OpenAI, Anthropic, and Azure OpenAI requests.  
  - **Experimental:** `/rewind` improvements (details not fully listed).

## Hot Issues (10 noteworthy)
1. **[#953](https://github.com/github/copilot-cli/issues/953) – Over excessive permissions request** (area:authentication, enterprise)  
   Users want finer‑grained control over which repos/areas the AI can access; current OAuth scope is too broad. 7 comments, 3 👍.

2. **[#3727](https://github.com/github/copilot-cli/issues/3727) – Regression: `userPromptSubmitted` hook no longer injects context** (area:plugins, v1.0.60)  
   Plugin authors report a major regression – the hook worked in v1.0.59 but is broken in v1.0.60+, breaking custom workflows. 4 comments.

3. **[#3282](https://github.com/github/copilot-cli/issues/3282) – Add multiple BYOK model capability** (area:models, configuration)  
   Users want to configure multiple bring‑your‑own‑key models and switch between them in‑session without restarting. 3 comments, 8 👍 – highest 👍 count.

4. **[#3781](https://github.com/github/copilot-cli/issues/3781) – Session enters unrecoverable error when pasting image with non‑multimodal model** (CLOSED)  
   Once an image attachment is added, every prompt fails with HTTP 400; the only workaround is manual editing of events.jsonl. 3 comments.

5. **[#3756](https://github.com/github/copilot-cli/issues/3756) – Third‑party MCP servers disabled by organization policy** (area:enterprise, MCP)  
   Enterprise users are blocked from using external MCP servers without clear policy controls. 3 comments.

6. **[#2966](https://github.com/github/copilot-cli/issues/2966) – Built‑in tooling for managing multiple concurrent CLI sessions** (area:sessions)  
   Power users running many sessions concurrently want first‑class session management. 3 comments.

7. **[#3769](https://github.com/github/copilot-cli/issues/3769) – Output mangled during streaming in Agency mode** (CLOSED, terminal‑rendering)  
   Threading issues cause garbled output until response is complete. 2 comments, 3 👍.

8. **[#3813](https://github.com/github/copilot-cli/issues/3813) – Copy/paste garbled text (Japanese) from VS Code terminal** (area:input‑keyboard, terminal‑rendering)  
   Similar to [#3776](https://github.com/github/copilot-cli/issues/3776) but specific to VS Code terminal and non‑ASCII text. 1 comment.

9. **[#3814](https://github.com/github/copilot-cli/issues/3814) – Requests failing but AIC consumption keeps increasing** (area:agents, models)  
   User experienced repeated transient errors while still being billed for failed requests. 2 👍.

10. **[#3812](https://github.com/github/copilot-cli/issues/3812) – Subagents cannot access MCP tools** (area:agents, MCP)  
   Custom subagents lost access to MCP tools; likely related to deferred loading. 0 comments but high visibility.

## Key PR Progress
Only one PR was updated in the last 24h:

- **[#3817](https://github.com/github/copilot-cli/pull/3817) – kCreate "#"** (OPEN, by edge500)  
   Summary: “aquellos” – appears to be a draft/test PR with no meaningful description. No further activity.

Given the low PR volume, the community focus remains on bug fixes and feature discussions via Issues.

## Feature Request Trends
- **Multi‑model BYOK management** – [#3282](https://github.com/github/copilot-cli/issues/3282), [#3399](https://github.com/github/copilot-cli/issues/3399): Allow multiple BYOK models with custom headers, and in‑session switching.
- **Concurrent session management** – [#2966](https://github.com/github/copilot-cli/issues/2966): Built‑in UI/tooling for running many sessions across repos/branches.
- **Prompt caching** – [#3808](https://github.com/github/copilot-cli/issues/3808): Leverage Anthropic’s prompt caching to reduce latency and cost for repeat context.
- **Unified history with VS Code** – [#3816](https://github.com/github/copilot-cli/issues/3816): Ingest VS Code Copilot Chat history into `/chronicle`.
- **Session content search** – [#3807](https://github.com/github/copilot-cli/issues/3807): Make `--resume` search inside session messages, not just names/IDs.

## Developer Pain Points
- **Regressions in plugin hooks** – The `userPromptSubmitted` hook regression (v1.0.60) has broken many custom plugins; no fix yet in the latest releases.
- **Session‑wedging bugs** – Unrecoverable errors from oversized attachments ([#3767](https://github.com/github/copilot-cli/issues/3767)), non‑multimodal images ([#3781](https://github.com/github/copilot-cli/issues/3781)), and failed requests still consuming AIC credits ([#3814](https://github.com/github/copilot-cli/issues/3814)) are causing frustration.
- **MCP server instability** – Unbounded respawn loops ([#3782](https://github.com/github/copilot-cli/issues/3782)) and enterprise policy blocks ([#3756](https://github.com/github/copilot-cli/issues/3756)) hinder adoption of MCP extensions.
- **Cross‑platform copy/paste issues** – UTF‑8 mojibake when copying from WSL/Ubuntu/VS Code to Windows persists ([#3776](https://github.com/github/copilot-cli/issues/3776), [#3813](https://github.com/github/copilot-cli/issues/3813)).
- **Platform‑specific failures** – Linux ARM64 Tokio reactor panic ([#3784](https://github.com/github/copilot-cli/issues/3784)) and Windows standalone executable extraction error ([#3810](https://github.com/github/copilot-cli/issues/3810)) affect users on non‑standard environments.

---

*Digest generated from `github/copilot-cli` data on 2026-06-16.*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-06-16

## Today's Highlights

Two important bug‑fix pull requests landed today, addressing long‑standing issues with the `UserPromptSubmit` hook receiving empty prompts and the `kimi --continue` command failing to find sessions. Meanwhile, a fresh issue reports that `FetchURL` ignores system proxy settings, a critical gap for users in restricted network environments.

## Releases

No new releases in the last 24 hours.

## Hot Issues

*(All 4 issues updated in the last 24h are listed below.)*

1. **[#2402] Error: [compaction.failed] APIStatusError: 400 – request rejected as high risk**  
   *Author:* thoughtworld · *Updated:* 2026‑06‑16 · *Comments:* 2  
   Running version 0.6.0 on Windows 10 x64 with model Kimi‑k2.6. Compaction fails with a 400 error, likely due to content moderation triggers. The issue has been open for over two weeks; the team hasn't responded yet, causing concern among users about sensitive content handling.  
   [GitHub](https://github.com/MoonshotAI/kimi-cli/issues/2402)

2. **[#2303] UserPromptSubmit hook receives empty prompt when input comes from shell UI**  
   *Author:* AkaCoder404 · *Updated:* 2026‑06‑15 · *Comments:* 1  
   On macOS (Sonoma 14.7, Apple Silicon), the `UserPromptSubmit` hook always gets an empty string, making regex‑based prompt matching impossible. This blocks many advanced workflows that rely on hook automation. A fix has been proposed in PR #2454.  
   [GitHub](https://github.com/MoonshotAI/kimi-cli/issues/2303)

3. **[#2222] `kimi --continue` fails with "No previous session found"**  
   *Author:* LiPingFeel · *Updated:* 2026‑06‑15 · *Comments:* 1  
   On Windows, the `-C` flag incorrectly claims no session exists even though running `kimi` without `--continue` shows the conversation history. The root cause is a missing `last_session_id` in the session metadata. PR #2453 provides a fix.  
   [GitHub](https://github.com/MoonshotAI/kimi-cli/issues/2222)

4. **[#2455] FetchURL ignores system proxy – cannot access external network in blocked environments**  
   *Author:* KuangYin‑Z · *Updated:* 2026‑06‑15 · *Comments:* 0  
   On Linux WSL2 with version 1.43.0, `FetchURL` does not respect system proxy settings, while `shell`/`curl` work fine. This prevents users in firewalled or censored networks from fetching external resources. No comments yet but the issue is fresh and likely to gather support.  
   [GitHub](https://github.com/MoonshotAI/kimi-cli/issues/2455)

## Key PR Progress

*(Both PRs updated in the last 24h are listed below.)*

1. **[#2454] fix(hooks): pass prompt text to UserPromptSubmit from structured input**  
   *Author:* logicwu0 · *Updated:* 2026‑06‑15 · *Status:* Open  
   Fixes issue #2303. The bug occurred because `KimiSoul._turn` derived the prompt from structured input incorrectly. The PR changes the text derivation to properly extract the user‑typed plain text, enabling regex‑based hook matching.  
   [GitHub](https://github.com/MoonshotAI/kimi-cli/pull/2454)

2. **[#2453] fix(session): resume latest session when last_session_id is missing**  
   *Author:* logicwu0 · *Updated:* 2026‑06‑15 · *Status:* Open  
   Fixes issue #2222. The `--continue` command relied solely on `work_dir` to find sessions, but when `last_session_id` wasn’t stored, it returned early with an error. This PR falls back to the most recent session in the directory when the explicit ID is absent.  
   [GitHub](https://github.com/MoonshotAI/kimi-cli/pull/2453)

## Feature Request Trends

Based on all recent issues, the community is requesting:

- **System proxy awareness** – `FetchURL` and probably other network‑bound components should respect OS proxy settings automatically. This is a strong demand from users in China and corporate networks.  
- **Better session continuity** – The `--continue` flag should robustly resume the last session, not only when `last_session_id` is present. The PR #2453 addresses this, but users would like a more deterministic session‑switching mechanism.  
- **Hook reliability** – Hooks (`UserPromptSubmit`) must work with all input methods, including the interactive shell UI, to allow advanced automation and custom workflows.

## Developer Pain Points

Recurring frustrations visible in the issue tracker include:

- **High‑risk request errors** – The “compaction.failed” 400 error (issue #2402) suggests content moderation may be blocking legitimate use, and the lack of a clear error message or retry strategy frustrates developers.  
- **Proxy ignorance** – Not reading system proxy settings forces users to use workarounds, breaking CI/CD pipelines and remote development environments.  
- **Session inconsistency** – The difference between plain `kimi` and `kimi --continue` behaviour is confusing; users expect a single‑source‑of‑truth for session history.  
- **Cross‑platform gaps** – Issues affect all three major OSes (Windows, macOS, Linux) in different ways, indicating insufficient testing in diverse environments.

*Digest generated from MoonshotAI/kimi-cli data as of 2026‑06‑16 23:59 UTC.*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-16

## Today’s Highlights
Memory instability remains the top community concern (97 comments), but the team is actively collecting heap snapshots to diagnose. Two high-impact fixes landed today: a Copilot Claude assistant-message prefill bug (#32508) and a proper UTF-8 encoding solution for Windows shell output (#31985). The long-running `/goal` session feature request continues to gain traction with 84 upvotes.

## Releases
No new releases in the last 24 hours.

## Hot Issues

1. **[#20695 – Memory Megathread](https://github.com/anomalyco/opencode/issues/20695)**  
   *97 comments, 65 👍*  
   Central issue for all memory-related reports. Team asks for heap snapshots (no LLM-generated suggestions). Highest-engagement thread this month.

2. **[#27167 – [FEATURE]: Add native session goals with /goal](https://github.com/anomalyco/opencode/issues/27167)**  
   *49 comments, 84 👍*  
   Users want persistent session lifecycle (goals that outlive individual prompts). Top-upvoted feature request.

3. **[#1970 – Feature Request: Add Background Bash Execution](https://github.com/anomalyco/opencode/issues/1970)**  
   *18 comments, 30 👍*  
   Parity with Claude Code’s `Ctrl+b` – running long tasks without blocking the agent.

4. **[#27906 – v1.15.1+ Breaks Bun Installs](https://github.com/anomalyco/opencode/issues/27906)**  
   *18 comments, 13 👍*  
   Recent release requires postinstall scripts; Bun blocks them by default. Affects global package installs.

5. **[#5374 – [FEATURE]: show tokens / second](https://github.com/anomalyco/opencode/issues/5374)**  
   *17 comments, 81 👍*  
   Display real-time and average throughput – critical for comparing providers and models.

6. **[#28567 – [FEATURE]: Full MCP client capabilities](https://github.com/anomalyco/opencode/issues/28567)**  
   *14 comments, 22 👍*  
   OpenCode’s MCP support lags behind the latest spec; users request resources, prompts, and proper tool schema.

7. **[#31247 – Copilot Claude Opus 4.8 emits pseudo tool-call text](https://github.com/anomalyco/opencode/issues/31247)**  
   *5 comments*  
   Critical bug: Copilot returns malformed assistant text instead of structured tool calls, breaking agent execution.

8. **[#30869 – bash.ts: hardcoded UTF-8 decoding garbles output on non-UTF-8 systems](https://github.com/anomalyco/opencode/issues/30869)**  
   *5 comments, 1 👍*  
   Windows users with CJK locales (e.g., GBK) get corrupted compiler errors.

9. **[#32493 – [FEATURE]: Moonshot provider missing kimi-k2.7-code-highspeed](https://github.com/anomalyco/opencode/issues/32493)**  
   *4 comments (new)*  
   Provider model list is not kept in sync; users request manual override or auto-discovery.

10. **[#32452 – Desktop: renderer unresponsive on startup](https://github.com/anomalyco/opencode/issues/32452)**  
    *2 comments*  
    Windows desktop 1.17.7 freezes within 60 seconds due to synchronous Markdown AST traversal – high severity for desktop users.

## Key PR Progress

1. **[#32508 – fix: handle Copilot Claude assistant prefill and tool text leaks](https://github.com/anomalyco/opencode/pull/32508)**  
   *Closed today*  
   Fixes 400 error when a conversation ends with an assistant message after interrupted tool calls. Closes #31807.

2. **[#31985 – fix(shell): use PowerShell EncodedCommand for reliable UTF-8 output on Windows](https://github.com/anomalyco/opencode/pull/31985)**  
   *Open*  
   Major fix for Windows shell encoding – replaces hardcoded UTF-8 with PowerShell base64 encoding. Closes 5 issues.

3. **[#31794 – fix(opencode): handle malformed percent-encoding in decodeDataUrl](https://github.com/anomalyco/opencode/pull/31794)**  
   *Open*  
   Prevents `URIError` crash when `data:` URLs contain unescaped characters.

4. **[#32499 – fix(opencode): allow clearing session archive time](https://github.com/anomalyco/opencode/pull/32499)**  
   *Open*  
   UX improvement: users can now un-archive a session instead of losing the ability to reorder.

5. **[#29150 – fix(opencode): break auto-compact loop when compaction makes no progress](https://github.com/anomalyco/opencode/pull/29150)**  
   *Open*  
   Prevents infinite compaction loop when model context limit is misconfigured. Closes #28543.

6. **[#32494 – fix(opencode): include pr identity in github context](https://github.com/anomalyco/opencode/pull/32494)**  
   *Open*  
   Adds PR number and URL to `opencode github run` context, enabling authoritative PR-comment operations.

7. **[#31645 – feat(cli): add progress feedback to upgrade command](https://github.com/anomalyco/opencode/pull/31645)**  
   *Open*  
   Real-time download progress instead of silent hang during `opencode upgrade`.

8. **[#32490 – feat(mcp): append server instructions to context](https://github.com/anomalyco/opencode/pull/32490)**  
   *Open*  
   First step toward full MCP compliance: sends `InitializeResult.instructions` to the agent. Refs #28567.

9. **[#31644 – fix(acp): register compact and summarize commands for visibility](https://github.com/anomalyco/opencode/pull/31644)**  
   *Open*  
   `/compact` and `/summarize` were missing from autocomplete and `/help` – now exposed.

10. **[#32489 – fix(opencode): sanitize OpenAI MCP tool schemas](https://github.com/anomalyco/opencode/pull/32489)**  
    *Open*  
    Strips incompatible JSON Schema keywords from MCP tool inputs to avoid OpenAI API errors.

## Feature Request Trends
- **Persistent session goals** (#27167) – top-voted feature; users want `/goal` to set long-lived objectives that survive command switches.
- **Background shell execution** (#1970) – demand for non-blocking `Ctrl+b` style tasks (servers, watchers).
- **Performance metrics** (#5374) – tokens/second display is the second most-upvoted feature.
- **MCP spec compliance** (#28567) – multiple PRs (#32490, #32489) show active work to catch up with standard resources, prompts, and tool schema handling.
- **Token budget optimization** (#21345) – moving verbose git instructions out of bash tool description to save ~1.7K tokens per request.
- **Localization & currency** (#32485, #15053) – custom currency display and KaTeX math rendering hints.
- **Agent-scoped skills** (#19344) – only load skills relevant to the active agent to reduce context bloat.

## Developer Pain Points
- **Memory instability** – unresolved, requires heap snapshots from the community (#20695).
- **Windows encoding issues** – repeated reports of garbled output on non-UTF-8 systems (#30869); PR #31985 targets root cause.
- **Copilot Claude tool-call breaks** – malformed assistant messages (#31247) and assistant prefill rejections (#31807) degrade the Copilot experience.
- **Package manager regressions** – Bun installs broken since v1.15.1 (#27906); npm postinstall lifecycle requirement.
- **Desktop crashes** – UI thread blocked by synchronous Markdown parsing on Windows (#32452).
- **MCP tool schema mismatches** – both JetBrains IDE and OpenAI proxy report schema validation errors (#32491, #32489).
- **Auto-compaction loops** – infinite context trimming when provider caps are smaller than ideal (#29150).
- **CLI startup speed** – lazy-loading PR (#27800) addresses slow `--help`/`--version` for users with many plugins.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest – 2026-06-16

**Repository:** [earendil-works/pi](https://github.com/earendil-works/pi)  
**Data source:** github.com/badlogic/pi-mono  
**Digest type:** Technical developer update

---

## Today’s Highlights

A new release (v0.79.4) brings automatic first-run theme selection, detecting terminal background to default to dark or light mode. Community attention is drawn to persistent reliability issues with the `openai-codex` provider (#4945, 58 comments), and a critical bug causing `pi` CLI commands to hang after completion (#5687) has been fixed. The project continues to see strong activity around provider expansion, session management fixes, and extension API refinements.

---

## Releases

**v0.79.4** – [Release link](https://github.com/earendil-works/pi/releases/tag/v0.79.4)  
- **Automatic first-run theme selection** – pi now detects the terminal background on first run and defaults to `dark` or `light` theme. See [Selecting a Theme](https://github.com/earendil-works/pi/blob/v0.79.4/packages/coding-agent/docs/themes.md#selecting-a-theme).  
- **Standalone** – (description truncated in source; likely further binary improvements).

---

## Hot Issues

1. **[#4945] openai-codex Connection Reliability Issues**  
   *Author: liushuaiiu, Comments: 58, 👍 30*  
   The `openai-codex` / `gpt-5.5` provider leaves the TUI stuck on `Working…` with no error. Only recovery is pressing Escape. High community engagement due to frequent occurrence.  
   [GitHub](https://github.com/earendil-works/pi/issues/4945)

2. **[#5103] pi-windows-x64.zip can't detect git-bash from PATH**  
   *Author: CXwudi, Comments: 21*  
   Windows build fails to locate Git Bash even when in PATH, blocking shell tool usage. Important for Windows users.  
   [GitHub](https://github.com/earendil-works/pi/issues/5103)

3. **[#4877] Session folder collision**  
   *Author: olivierverdier, Comments: 15, 👍 2*  
   Sessions from different paths may map to the same folder (e.g., `/a/b/c/d` vs `/a-b/c-d`). Low severity but could surprise users.  
   [GitHub](https://github.com/earendil-works/pi/issues/4877)

4. **[#5653] Move off Shrinkwrap**  
   *Author: yoyofield, Comments: 10*  
   Installing both `pi-ai` and `pi-coding-agent` as direct deps causes duplicate module copies, leading to a broken provider registry. Proposal to remove shrinkwrap.  
   [GitHub](https://github.com/earendil-works/pi/issues/5653)

5. **[#5687] pi list and pi update never exit when an extension runs an MCP server**  
   *Author: unship, Comments: 7*  
   Package subcommands hang after completing output. Caused by long-lived MCP server handles not being cleaned up.  
   [GitHub](https://github.com/earendil-works/pi/issues/5687)

6. **[#5728] Support provider-specific config in auth.json**  
   *Author: michaeldwan, Comments: 6*  
   Providers like `cloudflare-ai-gateway` need more than an API key (accountId, gatewayId). Request to store such config in `auth.json`.  
   [GitHub](https://github.com/earendil-works/pi/issues/5728)

7. **[#5739] Add SHA256SUMS and provenance attestations to binary releases**  
   *Author: shaanmajid, Comments: 5*  
   npm packages have provenance, but binary releases lack integrity checks. Community pushing for supply chain security.  
   [GitHub](https://github.com/earendil-works/pi/issues/5739)

8. **[#5778] pi-agent-core hangs indefinitely on unresponsive streams or tool deadlocks**  
   *Author: Paramveersingh-S, Comments: 3*  
   Critical vulnerability: agent loop wedges if LLM stream drops or tool `execute()` never resolves. Fix merged quickly.  
   [GitHub](https://github.com/earendil-works/pi/issues/5778)

9. **[#5774] pi update with `npmCommand ["bun"]` creates package.json/bun.lock in caller cwd**  
   *Author: unitdhda, Comments: 3*  
   Running `pi update` from home directory with bun writes unwanted files to cwd. Affects users with non-npm package managers.  
   [GitHub](https://github.com/earendil-works/pi/issues/5774)

10. **[#5008] Working spinner comes back after response ends**  
    *Author: cuba6112, Comments: 3*  
    Spinner persists after `/new`, `/reload`, or session switch until restart. Annoying visual bug with community workarounds.  
    [GitHub](https://github.com/earendil-works/pi/issues/5008)

---

## Key PR Progress

1. **[#5789] fix(tui): restore cursorUp line-start jump before history browsing**  
   *Author: 4h9fbZ*  
   Fixes editor boundary behavior where pressing Up on first line jumps to line start instead of entering history.  
   [PR](https://github.com/earendil-works/pi/pull/5789)

2. **[#5675] fix: stabilize compaction after reload**  
   *Author: SeanThomasWilliams*  
   Prevents compaction failures after reload or during queued message delivery. Ensures token boundaries are preserved.  
   [PR](https://github.com/earendil-works/pi/pull/5675)

3. **[#5784] fix(coding-agent): sort threaded sessions by latest activity in the subtree**  
   *Author: Perlence*  
   Improves session list ordering for forked sessions, sorting by most recent child activity instead of root modification date.  
   [PR](https://github.com/earendil-works/pi/pull/5784)

4. **[#5779] feat(coding-agent): XML-structure /review prompt responses**  
   *Author: hansjm10*  
   Converts `/review` to XML-structured instructions and task envelopes, with coverage-aware workflow and updated JSON parsing.  
   [PR](https://github.com/earendil-works/pi/pull/5779)

5. **[#5776] Fix Agent Wedge on Unresponsive Streams & Tool Executions**  
   *Author: Paramveersingh-S*  
   Adds timeout/error handling to prevent indefinite hangs in agent loop. Closes #2381.  
   [PR](https://github.com/earendil-works/pi/pull/5776)

6. **[#5758] feat(coding-agent): diagnose when a child holds stdio open past exit**  
   *Author: Mearman*  
   Follow-up to #5753: improves handling of detached descendants that keep writing after shell exits.  
   [PR](https://github.com/earendil-works/pi/pull/5758)

7. **[#5587] feat(coding-agent): add experimental first-time setup flow**  
   *Author: vegarsti*  
   Behind `PI_EXPERIMENTAL=1`, shows first-run dialog with theme selection (dark/light live preview) and analytics opt-in.  
   [PR](https://github.com/earendil-works/pi/pull/5587)

8. **[#2331] feat(extensions): add vim-like modal editor extension**  
   *Author: Nokodoko*  
   Full modal editor with Normal, Insert, Visual, Command modes, motions, operators, counts, and ex commands. Status bar included.  
   [PR](https://github.com/earendil-works/pi/pull/2331)

9. **[#5769] fix(render-utils): TUI crash when tool returns result without content array**  
   *Author: TomSpoct*  
   Fixes crash caused by `getTextOutput()` assuming content array always present.  
   [PR](https://github.com/earendil-works/pi/pull/5769)

10. **[#5509] feat: Add Amazon Bedrock Mantle OpenAI Responses provider**  
    *Author: unexge*  
    New provider for Amazon Bedrock Mantle’s OpenAI-compatible API, supporting GPT 5.5/5.4 models.  
    [PR](https://github.com/earendil-works/pi/pull/5509)

---

## Feature Request Trends

- **Provider expansion** – Multiple requests for new LLM providers: ZhipuAI (#2345), ZAI China (#5792), Amazon Bedrock Mantle (#5509), Gemini 3.5 Flash for Vertex (#5761), and Moonshot AI K2.7 (#5760).  
- **Provider-specific configuration** – Users want to store per-provider settings (e.g., accountId, gatewayId) in `auth.json` (#5728) instead of relying only on environment variables.  
- **Security and provenance** – Requests for integrity checks (SHA256SUM, provenance attestations) for binary releases (#5739) and pinning AWS SDK dependencies (#5782).  
- **Extension API improvements** – Making `pi.sendUserMessage()` and `pi.sendMessage()` return a proper Promise (#5751), exposing edit-diff for extensions (#5756), and providing a prompt guideline API (#5711).  
- **Truncation configuration** – Users want environment variables to override hardcoded max lines/bytes (#5759).  
- **Theme and setup flow** – First-run theme detection (already in v0.79.4) and experimental setup dialog (#5587) address onboarding friction.

---

## Developer Pain Points

- **Connection reliability** – Frequent issues with `openai-codex` provider stalling the TUI without errors (#4945, #5778, #5740).  
- **CLI hang problems** – Commands like `pi list` and `pi update` not exiting when MCP servers run (#5687, #5645).  
- **Windows compatibility gaps** – Git Bash detection fails (#5103); session folder collisions (#4877) affect cross-platform usage.  
- **Spinner and rendering bugs** – Stuck spinner after response (#5008, #5740), TUI crashes on long status lines (#5773), broken Markdown rendering (#5766).  
- **Session management headaches** – Auto-compaction errors (#5463), threaded session sorting (#5784), and non-descriptive session folder names (#4877).  
- **Package manager integration** – `bun` creates unwanted files in caller cwd (#5774); shrinkwrap duplication (#5653) wastes disk space.  
- **Security concerns** – `--min-release-age=0` in npm updates undermines supply-chain defenses (#5785).  
- **Missing terraform/enterprise features** – Enterprise users face Nexus IQ failures due to floating AWS SDK dependencies (#5782).

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-06-16

## Today's Highlights

The team shipped **v0.18.1** with a security-sensitive opt‑in for direct shell access, plus a **desktop‑v0.0.4** release that persists MCP server removals. The community is actively discussing a false‑positive Trojan detection in the VSIX package (#5055) and a crash‑prone multi‑agent workflow (#5180). On the development side, the `/loop` self‑paced wakeup primitive (#5182) and a fix for plan‑mode infinite retries (#5185) landed, signalling continued investment in background automation.

---

## Releases

| Version | Key Changes |
|---------|-------------|
| **v0.18.1** ([changelog](https://github.com/QwenLM/qwen-code/releases/tag/v0.18.1)) | `feat(daemon): gate direct session shell behind explicit opt‑in` – users must now opt in before the daemon can present a raw shell. |
| **v0.18.1‑preview.0** / **v0.18.1‑nightly.20260616** | Fix: warn on oversized context instructions; docs: stale defaults, CLI syntax, tool naming drift. |
| **desktop‑v0.0.4** ([changelog](https://github.com/QwenLM/qwen-code/releases/tag/desktop-v0.0.4)) | `fix(cli): persist MCP server removals` and `fix(models): refresh raw model‑derived defaults`. |

---

## Hot Issues (10 of 17)

1. **[#5055 – Trojan:JS/ShaiWorm.DBA!MTB false positive in VSIX](https://github.com/QwenLM/qwen-code/issues/5055)**  
   *Priority P1, Security*  
   A Windows user reports Windows Defender flagging the `vscode-ide-companion-0.18.0` VSIX. Team has 5 comments but no resolution yet – high community concern.

2. **[#5124 – Track `/loop` alignment work (parent issue)](https://github.com/QwenLM/qwen-code/issues/5124)**  
   *Feature request, CLI, roadmap*  
   A staged plan for self‑paced `/loop` iterations. Each child issue is small, landed independently. Already spawning #5156, #5184, #5182.

3. **[#5180 – Multi‑agent task crashes after hours of execution](https://github.com/QwenLM/qwen-code/issues/5180)**  
   *P2, Bug, Core/Tools, welcome‑pr*  
   A 12‑hour session where a main agent delegates to sub‑agents, but sub‑agents crash half‑way. Community suspects memory pressure or tool‑call imbalance.

4. **[#5173 – Model provider disambiguation fails when multiple providers share same model ID](https://github.com/QwenLM/qwen-code/issues/5173)**  
   *P2, Bug, Configuration*  
   Selecting `qwen3.7-max` from IdeaLab’s provider does not persist across sessions – model picker reverts. Quickly followed by PR #5179.

5. **[#4966 – SchemaValidator missing numeric string coercion causes MCP tool failures](https://github.com/QwenLM/qwen-code/issues/4966)**  
   *P2, Bug, Core/Tools, MCP*  
   LLMs emit `"depth": "3"` but strict MCP servers expect integer. Closed after adding coercion.

6. **[#5184 – Wire prompt‑only `/loop` to self‑paced wakeups](https://github.com/QwenLM/qwen-code/issues/5184)**  
   *P3, Enhancement, blocked*  
   Child of #5124. Makes `/loop <prompt>` start a self‑paced loop instead of a fixed schedule – model schedules one future continuation.

7. **[#5177 – `exit_plan_mode` fails with empty plan parameter](https://github.com/QwenLM/qwen-code/issues/5177)**  
   *P3, Bug, Core/Tools*  
   In plan mode, the model sometimes calls `exit_plan_mode` without a plan, causing wasted retry turns. Has 1 👍.

8. **[#4939 – Let grep/egrep/fgrep satisfy read‑before‑edit check](https://github.com/QwenLM/qwen-code/issues/4939)**  
   *P2, Feature request, File operations*  
   Avoiding an extra `Read` call when the model already grepped the file. Closed after implementation.

9. **[#5154 – Discussion: does `--expose-gc` wrapper earn the extra process?](https://github.com/QwenLM/qwen-code/issues/5154)**  
   *P3, Performance, Memory‑usage*  
   Design question from #4914: is a separate Node process worth the memory overhead to make `global.gc()` available? Not urgent.

10. **[#5176 – Allow sub‑agent max parallel count setting with queue](https://github.com/QwenLM/qwen-code/issues/5176)**  
    *Feature request*  
    For local LLMs: limit parallel sub‑agents, put extras in a pending queue. No comments yet.

---

## Key PR Progress (10 of 50)

1. **[#5182 – Add session wakeup primitive (`loop_wakeup` tool)](https://github.com/QwenLM/qwen-code/pull/5182)**  
   Foundational slice for self‑paced `/loop`. One‑shot wakeup, not yet wired into the skill. Closes #5156.

2. **[#5185 – Fix plan‑gate infinite retry loop](https://github.com/QwenLM/qwen-code/pull/5185)**  
   Isolates gate agent’s `AbortSignal` from parent signal chain – prevents stuck retries in AUTO/YOLO pre‑plan mode.

3. **[#5181 – Prevent OOM in auto‑memory extraction during `/quit`](https://github.com/QwenLM/qwen-code/pull/5181)**  
   Fixes `FATAL ERROR: Reached heap limit` – `buildTranscriptMessages()` now processes memory‑efficiently instead of attaching full conversation.

4. **[#4971 – Reduce retained interactive tool output memory](https://github.com/QwenLM/qwen-code/pull/4971)**  
   Compacts oversized tool‑output display metadata before storing in scheduler state – reduces memory pressure in CLI.

5. **[#4943 – Add `--safe-mode` flag for troubleshooting](https://github.com/QwenLM/qwen-code/pull/4943)**  
   Disables all customizations (context files, hooks, extensions, MCP, subagents) – clean baseline for diagnosing user issues.

6. **[#5179 – Remember selected provider when multiple share a model ID](https://github.com/QwenLM/qwen-code/pull/5179)**  
   Persists `baseUrl` alongside model name – resolves #5173’s non‑persistent provider selection.

7. **[#4564 – Expose token usage stats for cost visibility](https://github.com/QwenLM/qwen-code/pull/4564)**  
   Persists daily/monthly token usage; `/stats` broken down by model/auth‑type; export as CSV/JSON.

8. **[#4850 – Interactive multi‑tab `/extensions` manager](https://github.com/QwenLM/qwen-code/pull/4850)**  
   Three tabs – Installed, Discover, Sources – covering full lifecycle of extensions and MCP servers.

9. **[#5094 – Workflow P4 – meta extraction + `/workflows` + phase-tree](https://github.com/QwenLM/qwen-code/pull/5094)**  
   Completes Phase 4 of Dynamic Workflows port: meta extraction, workflow loading, phase‑tree display.

10. **[#5174 – daemon status API (closed)](https://github.com/QwenLM/qwen-code/pull/5174)**  
    Adds read‑only `GET /daemon/status` for `qwen serve` – reports session counts, permission pressure, transport counts.

---

## Feature Request Trends

- **Self‑paced background automation** – the `/loop` alignment work (#5124, #5184, #5156) is the clearest signal. Users want the model to schedule its own future wakeups instead of fixed recurring schedules.
- **Sub‑agent concurrency control** – #5176 requests a configurable max parallel sub‑agents with a pending queue, especially for local‑LLM users with limited resources.
- **Richer hook/terminal integration** – #4882 asks for a `terminalSequence` field on hooks (desktop notifications, bell, window‑title updates), mirroring Claude Code.
- **Reduced friction for common workflows** – #4939 (grep satisfying read‑before‑edit) and #4966 (numeric string coercion) show demand for fewer redundant tool calls.

---

## Developer Pain Points

- **False positive antivirus detection** (#5055) – the VSIX flagged as Trojan creates trust issues; team needs to provide signing or exclusion guidance.
- **Memory and OOM crashes** in long sessions – #5180 (multi‑agent crash) and #5181 (OOM after `/quit`) are recurring themes. The memory‑reduction PR #4971 is welcome, but more systematic memory management is needed.
- **Model provider disambiguation** (#5173) – multiple providers with the same model ID remain confusing; PR #5179 is a fix but the UX of the model picker may need rethinking.
- **Plan‑mode infinite retries** (#5177, #5185) – `exit_plan_mode` failures waste user time; the abort‑signal fix (#5185) should help but more robust handling is desired.
- **Windows shell compatibility** (#4562) – users struggle with `cmd.exe` vs PowerShell and `!` commands; ongoing need for better Windows defaults.
- **Session resumption unclear** (#3099) – old issue still open: resume list doesn’t show conversation titles, making it impossible to distinguish sessions.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-06-16

## Today's Highlights

CodeWhale continues to stabilize after the v0.8.60–v0.8.61 cycle, with substantial activity around sub-agent reliability, provider compatibility, and TUI responsiveness. The most urgent item is the persistent **“Turn stalled” bug** (Issue #2487), which has drawn the highest community engagement in recent days. Several key PRs landed today addressing paste truncation, new provider support (DeepInfra), and update reliability, while two new bugs were filed around sub-agent deadlocks and Kimi API compatibility.

## Releases

*No new releases in the last 24 hours.*

## Hot Issues

1. **[#2487 — Frequent error: Turn stalled - no completion signal received](https://github.com/Hmbown/CodeWhale/issues/2487)**  
   *13 comments* — The top community pain point. Users in `yolo` mode experience freezes with no ability to resume via `continue`. The stalled state persists across sessions, and session content may be lost. This is a critical reliability concern affecting v0.8.61.

2. **[#3063 — v0.8.59 release tracker: TUI mouse-report leak, runtime safety, and issue/PR queue](https://github.com/Hmbown/CodeWhale/issues/3063)**  
   *11 comments* — CLOSED. The maintainer’s stabilization tracker covering the TUI input leak on macOS and triage of pending PRs. Serves as the canonical reference for the v0.8.59 patch scope.

3. **[#1186 — feat(execpolicy): add typed persistent permission rules](https://github.com/Hmbown/CodeWhale/issues/1186)**  
   *9 comments* — A long-standing enhancement request for scoped permission rules (by tool, command prefix, workspace path). Community has been waiting since early May; this is tagged for v0.9.0.

4. **[#3096 — Split sub-agents into headless worker runtime with lightweight TUI projections](https://github.com/Hmbown/CodeWhale/issues/3096)**  
   *8 comments* — CLOSED. A major architectural proposal to decouple sub-agent execution from TUI-heavy scaffolding. The community discussion reflects interest in leaner multi-agent workflows.

5. **[#3192 — Add CodeWhale to AgentClientProtocol/registry](https://github.com/Hmbown/CodeWhale/issues/3192)**  
   *6 comments* — A straightforward request to be listed in the ACP registry, which would enable seamless installation from Zed and similar editors. Community reaction is positive; simple lift for maintainers.

6. **[#1812 — TUI freeze on Windows (crossterm poll)](https://github.com/Hmbown/CodeWhale/issues/1812)**  
   *6 comments* — Windows 11 users report intermittent UI freezes with logs and thread-state analysis. Process stays alive but no input is accepted. Two separate events captured, pointing to a crossterm polling issue.

7. **[#3102 — v0.8.62: Add first-class clarification question requests for agents](https://github.com/Hmbown/CodeWhale/issues/3102)**  
   *4 comments* — Proposes a formal mechanism for agents to ask clarifying questions via modals instead of hoping users notice chat messages. Tagged for v0.8.62.

8. **[#2629 — Cannot use SiliconFlow / TokenHub due to 401 error](https://github.com/Hmbown/CodeWhale/issues/2629)**  
   *4 comments* — Users on Windows 11 with v0.8.50 cannot use OpenAI-compatible providers SiliconFlow and TokenHub. The 401 persists despite correct configuration. Indicates a provider compatibility gap.

9. **[#3264 — Restrict skill scanning to ~/.codewhale/skills/ only](https://github.com/Hmbown/CodeWhale/issues/3264)**  
   *3 comments* — NEW today. User requests an option to limit skill discovery to a single directory, avoiding conflicts from global installations.

10. **[#3266 — Sub-agent agent_eval with block=True causes TUI freeze/deadlock](https://github.com/Hmbown/CodeWhale/issues/3266)**  
    *1 comment* — NEW today. A critical bug where calling `agent_eval` with `block=True` on multiple sub-agents causes indefinite TUI deadlock. No completion events are received, forcing terminal kill.

## Key PR Progress

1. **[#3267 — Keep oversized paste inline with truncation and auto-expand](https://github.com/Hmbown/CodeWhale/pull/3267)**  
   OPEN. Fixes a UX regression where pastes exceeding 16K chars were silently replaced with a file reference. Now preserves the text inline with truncation, allowing users to select/copy/edit.

2. **[#3235 — Add DeepInfra provider support](https://github.com/Hmbown/CodeWhale/pull/3235)**  
   CLOSED. Adds full support for DeepInfra’s inference API (100+ models including DeepSeek V4). Config accepts aliases `deepinfra`, `deep-infra`, `deep_infra`.

3. **[#3244 — Retry release lookups and downloads](https://github.com/Hmbown/CodeWhale/pull/3244)**  
   CLOSED. Improves update robustness by retrying transient GitHub API failures on release metadata and asset downloads, with fallback URL construction.

4. **[#3241 — Accept dollar skill aliases (e.g., `$skill-name`)](https://github.com/Hmbown/CodeWhale/pull/3241)**  
   CLOSED. Adds composer shorthand for activating skills (e.g., `$lint`). Preserves backward compatibility with `/skill` and `/<skill>` flows.

5. **[#3233 — Persist ask-only permission rules atomically](https://github.com/Hmbown/CodeWhale/pull/3233)**  
   CLOSED. Foundation PR for persistent permissions: adds `ConfigStore::append_ask_rules` and refactors config error enums. No UI changes yet — paves the way for #1186.

6. **[#3206 — WeChat bridge leveraging Feishu and Tencent OpenClaw](https://github.com/Hmbown/CodeWhale/pull/3206)**  
   CLOSED. Community contribution that adds a WeChat integration using the existing Feishu bridge infrastructure. Extends CodeWhale’s reach to China’s dominant messaging platform.

7. **[#3257 — Make app-server the canonical runtime API entrypoint](https://github.com/Hmbown/CodeWhale/pull/3257)**  
   CLOSED. Refactors `codewhale app-server --http`/`--mobile` to delegate to the `serve` runtime path. Adds release smoke coverage for the control surface.

8. **[#3242 — Add workspace_follow_symlinks setting](https://github.com/Hmbown/CodeWhale/pull/3242)**  
   OPEN. Enables symlink-aware directory traversal for walk-based tools — important for users with linked project directories.

9. **[#3239 — Add Atlas Cloud as OpenAI-compatible LLM backend](https://github.com/Hmbown/CodeWhale/pull/3239)**  
   OPEN. Documentation-only PR adding Atlas Cloud (59 models) as a provider option. No code changes, low risk.

10. **[#3005 — Refactor provider metadata into data-driven registry](https://github.com/Hmbown/CodeWhale/pull/3005)**  
    CLOSED. Major refactor replacing ~100 hand-maintained match arms with a static `PROVIDER_REGISTRY`. Introduces `aliases()` per provider trait. Reduces maintenance overhead significantly.

## Feature Request Trends

- **Provider fallback chains** (#2574) — Users want automatic failover when an API provider returns 401/429/5xx errors, instead of manual `/provider` switching. This is the most-requested new feature.
- **Goal/long-running task mode** (#891, #2058, #1976) — Persistent interest in Codex-style objective-driven workflows with LLM-as-judge continuation loops. Multiple issues track this for v0.9.0.
- **Dynamic/scripted API keys** (#3004) — Users want to fetch API keys via shell scripts at runtime rather than storing them in plaintext `.env` or `config.toml`. Security-conscious developers are driving this.
- **Persistent typed permission rules** (#1186, #3233) — The community wants scoped `allow`/`deny`/`ask` rules that persist across sessions, scoped by tool, command, or path.
- **AgentClientProtocol registry listing** (#3192) — A low-effort request with high impact for editor integration, particularly Zed.

## Developer Pain Points

- **TUI freezes and stalls** (#2487, #1812, #2739, #3266) — The single biggest reliability concern. Users experience frozen TUI across Windows and Linux, often losing session context. The “Turn stalled” error in yolo mode (#2487) is the most commented issue in the dataset.
- **Sub-agent reliability** (#1679, #2652, #3266) — Multi-agent workflows remain fragile: SSE timeouts on Windows, clipped evaluation output that misleads the model, and deadlocks with `block=True`. Users are downgrading to solo mode.
- **Provider compatibility gaps** (#2629, #3265) — OpenAI-compatible providers like SiliconFlow and Kimi silently fail (401 or 400 errors) despite correct configuration. The Kimi bug (#3265) shows a strict `parameters.type` requirement not handled by CodeWhale’s tool schema.
- **Session loss on errors** (#2739) — When tasks stall or time out, users report that exiting and re-entering with `--continue` loses the previous session entirely. This trust-breaking behavior has driven some users to abandon the tool.
- **glibc version requirements** (#1067) — Static builds requiring glibc 2.38/2.39 exclude many enterprise server environments still on 2.35 (Ubuntu 22.04 LTS).

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*