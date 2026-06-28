# AI CLI Tools Community Digest 2026-06-28

> Generated: 2026-06-28 10:09 UTC | Tools covered: 9

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
**Date:** 2026-06-28

---

## 1. Ecosystem Overview

The AI CLI tools landscape on June 28, 2026, shows a market in **intense, uneven maturation**. Major vendors (Anthropic, OpenAI, Google, GitHub, Alibaba) continue to ship at different cadences—OpenAI Codex pushed three alpha releases today, while Claude Code and Copilot CLI shipped none—but the **open-source and community-driven tools** (OpenCode, Pi, DeepSeek TUI) are closing feature gaps rapidly, evidenced by DeepSeek TUI's 15 PRs in 24 hours and OpenCode's background task management system. A **unifying theme across all tools is the struggle with Windows stability**, agent reliability (hallucination, fabricated results, silent failures), and the growing demand for MCP (Model Context Protocol) ecosystem integration. The divide between tools that prioritize **safety and auditability** (Claude Code's `protect-mcp` plugin, Gemini CLI's Cedar policy gate) versus those emphasizing **flexibility and extensibility** (Pi's extension API, DeepSeek TUI's plugin system) is becoming more pronounced, reflecting different strategic bets on how developers should interact with AI agents.

---

## 2. Activity Comparison

| Tool | Noteworthy Issues (last 24h) | Key PRs (last 24h) | Releases Today | Community Pulse |
|---|---|---|---|---|
| **Claude Code** | 10 (3 new critical) | 3 (1 major: `protect-mcp` plugin) | None | High engagement, stale bugs persist |
| **OpenAI Codex** | 10 (6 open critical) | 10 (3 SQLite fix merges, 5 MCP OAuth stack) | 3 alpha patches | High velocity; Windows & GPT-5.5 concerns |
| **Gemini CLI** | 10 (5 P1, 5 P2) | 10 (7 closed, 3 open) | 1 nightly release | Steady progress; security-focused fixes |
| **GitHub Copilot CLI** | 8 (2 critical regressions) | 3 (1 meaningful: `.gitignore` config) | None | Lower activity; Linux auth stall |
| **Kimi Code CLI** | 1 (memory bug updated) | 0 | None | Minimal engagement; growing risk |
| **OpenCode** | 10 (3 new CPU/memory reports) | 10 (3 major features: background tasks, `/usage`, archive) | None | High community energy; Windows segfault urgent |
| **Pi** | 10 (streaming UX, provider bugs) | 8 (4 closed, 4 open/discuss) | None | Broad but fragmented issues; extension API traction |
| **Qwen Code** | 10 (4 closed, 6 open) | 10 (multiplayer agent, chat panel, voice) | v0.19.3 hotfix | Strong feature push; token management gaps |
| **DeepSeek TUI** | 14 (cache/token dominated) | 15 (plugin system, ACP streaming, Korean docs) | None | Intense iteration; plugin ecosystem launch |

**Key takeaway:** DeepSeek TUI leads in PR velocity (15 in 24 hours), while OpenAI Codex and OpenCode tie for feature depth. Kimi Code CLI is effectively dormant—a warning signal for users considering adoption.

---

## 3. Shared Feature Directions

The following requirements surfaced across **three or more** tool communities, indicating strong market demand:

| Requirement | Tools Affected | Specific Needs |
|---|---|---|
| **MCP Ecosystem Integration** | Claude Code, Gemini CLI, Copilot CLI, DeepSeek TUI, OpenCode | OAuth credential stores, tool schema normalization, fail-closed gates, plugin registries, signed receipts |
| **Windows Platform Parity** | Claude Code, OpenAI Codex, Copilot CLI, OpenCode, DeepSeek TUI | Certificate errors, path delimiter bugs, sandbox failures, installer freezes, `.bat`/`.cmd` execution, updater hangs |
| **Agent Reliability / Hallucination Guards** | Claude Code, Gemini CLI, OpenCode, DeepSeek TUI | Fabricated verification results (#70231), silent tool failures (#69370), over-engineering loops (#3275), step-cap prefill issues (#32548) |
| **Token & Cost Transparency** | OpenAI Codex, OpenCode, Qwen Code, DeepSeek TUI | Usage dashboards (`/usage` command), cost reset timers, session-level token tracking, auto-compaction thresholds |
| **Session Persistence & Resumption** | Claude Code, Copilot CLI, Qwen Code, DeepSeek TUI | Resumable ACP streams, session retention dates, context compression, background task lifecycle |
| **Multi-Agent / Sub-Agent Management** | Claude Code, Gemini CLI, OpenAI Codex, OpenCode | Fork context visibility (#71464), agent-to-agent calling (#22092), TUI agent dashboards (#22321), delegation control |
| **Plugin / Extension APIs** | Claude Code, Pi, DeepSeek TUI, OpenCode | Safe reload hooks, `excludeFromContext`, usage reporting APIs, manifest-based discovery |
| **Output Truncation & Scroll Control** | Claude Code (bash truncation #26954), Pi (scroll jump #5825), Qwen Code (#5941), OpenCode (#34279) | Expandable output, scroll-lock during streaming, diff rendering reliability |

---

## 4. Differentiation Analysis

### Feature Focus & Target Users

| Tool | Primary Focus | Target User | Technical Approach |
|---|---|---|---|
| **Claude Code** | **Safety & agentic auditability** | Enterprise developers, security-conscious teams | Cedar policy gates, signed receipts, audit trails; conservative agent behavior |
| **OpenAI Codex** | **Multi-agent orchestration** | Power users running parallel agents | Rust CLI + desktop app; GPT-5.5 reasoning optimization; ACP protocol |
| **Gemini CLI** | **Provider flexibility & MCP compliance** | Multi-cloud developers, GCP users | Nightly releases; aggressive MCP tool schema normalization; VS Code companion |
| **Copilot CLI** | **GitHub ecosystem integration** | GitHub-centric teams | Tight coupling with GitHub auth, hooks, and PR workflows; slower iteration |
| **Kimi Code CLI** | **Minimalist VSCode plugin** | Users wanting lightweight IDE integration | Single-plugin approach; negligible community or feature velocity |
| **OpenCode** | **Cross-platform desktop UX** | Desktop-first developers (macOS, Windows, Linux) | TUI + desktop app; Bun runtime; rich feature set (background tasks, archives) |
| **Pi** | **Extensible TUI platform** | Advanced CLI users, extension developers | Extension API-first; multi-provider support; active Discord-driven design |
| **Qwen Code** | **Multi-platform + channel agents** | Chinese market, QQ/DingTalk users | Daemon architecture; multiplayer chat agents; shared UI packages across platforms |
| **DeepSeek TUI** | **Token efficiency & cache optimization** | Cost-sensitive heavy API users | Plugin system; ACP streaming; aggressive prompt compaction; Moraine memory backend |

### Key Strategic Differences

- **Safety vs. Flexibility:** Claude Code and Gemini CLI invest heavily in security gates and policy enforcement (Cedar, signed receipts), while Pi and DeepSeek TUI prioritize extensibility and rapid plugin development. OpenAI Codex and OpenCode sit in the middle, balancing safety with feature velocity.

- **Platform Betting:** OpenCode and Pi are TUI-first with desktop companions; OpenAI Codex and Qwen Code invest in native desktop apps; Gemini CLI and Copilot CLI remain CLI-centric with VS Code integration. Claude Code alone has no desktop app, relying on terminal and VS Code extension.

- **Provider Lock-in:** Claude Code is Anthropic-exclusive; Copilot CLI is GitHub/Microsoft-bound; Gemini CLI strongly favors Google's Vertex AI. OpenAI Codex is OpenAI-exclusive but expanding remote plugins. Pi and DeepSeek TUI are multi-provider by design, with Pi explicitly handling provider-specific quirks (Groq `reasoning_content`, Xiaomi pricing).

- **Community Model:** DeepSeek TUI and Pi have the most active, Discord-driven communities with rapid PR merging. Claude Code and OpenAI Codex have larger user bases but slower response to community contributions. Kimi Code CLI appears to have abandoned community engagement.

---

## 5. Community Momentum & Maturity

### High Momentum / Rapidly Iterating

- **DeepSeek TUI** — The most active project by PR volume (15 in 24h). The plugin system launch (manifest parsing, MCP integration, CLI subcommands) represents a major architectural milestone. Cache and token efficiency remain the dominant community concern, but the project is responding with dedicated EPICs (#3388). Community is engaged and vocal. **Risk:** Over-engineering and feature bloat may slow the core UX.

- **OpenCode** — Feature-rich with high community engagement (background tasks, `/usage` command, project archive, popover v2). However, three CPU/memory issues in one day (#34236, #34289, #34288) and the Windows Bun segfault (#33742) suggest quality assurance is lagging behind feature velocity.

- **Qwen Code** — Strong feature push with multiplayer agent (Qwen Tag), shared chat panel, voice dictation, resumable ACP sessions, and LSP hot reload—all in PR. Token management bugs (#5950, #5756) and the stealth model upgrade bug (#5819) indicate growing pains in maturing infrastructure.

### Mature but Stalled

- **Claude Code** — High user engagement (17+ comments on bash truncation) but **stale bugs** are a pattern: #26954 has been open for months, and critical issues like fabricated results (#70231) and silent tool failures (#69370) remain unresolved. The community contribution of `protect-mcp` (#72014) is promising, but Anthropic's own release cadence is slow.

- **OpenAI Codex** — Three alpha releases in 24h suggest active development, but the **near-catastrophic SQLite bug** (#28224) that could have written 640 TB/year before a community-driven fix erodes trust in code review processes. GPT-5.5 token clustering (#30364) is a worrying, opaque issue.

- **Gemini CLI** — Steady nightly releases with targeted fixes (case-insensitive path blocklist, MCP MIME types, OAuth refresh). Community engagement is moderate; the most upvoted open issues (skill discovery #25693, terminal focus #22193) are months old, suggesting prioritization gaps.

### Low Momentum / Warning Signals

- **GitHub Copilot CLI** — Only 3 PRs in 24h, one of which is dubious (#3737). The Ubuntu keychain bug (#2165, 20👍) has been open since March with no fix. Windows regression (#3958) on MCP server startup is fresh but symptomatic of quality issues. The `/btw` feature request (#2778) hints that users are looking elsewhere.

- **Kimi Code CLI** — Effectively stagnant. One issue updated in 24h (memory bug from March), zero PRs, zero releases. **Strong recommendation:** Teams evaluating Kimi Code should treat it as at-risk of abandonment.

### Maturity Assessment Summary

| Tool | Maturity Level | Risk Factors | Best For |
|---|---|---|---|
| Claude Code | High | Stale bugs, slow release cadence | Enterprise security-audit workflows |
| OpenAI Codex | High-Medium | GPT-5.5 opaque behavior, Windows issues | Multi-agent power users |
| Gemini CLI | Medium | Delayed triage on top issues | GCP/Vertex AI ecosystem |
| Copilot CLI | Medium | Linux/auth stagnation, Windows regressions | GitHub-integrated teams |
| Kimi Code CLI | Low | Dormant, likely abandoned | Avoiding |
| OpenCode | Medium-High | QA gaps, CPU/memory leaks | Desktop-first, feature-rich workflows |
| Pi | Medium | Fragmented UX bugs, provider friction | Extensibility tinkerers |
| Qwen Code | Medium-High | Token management bugs | Chinese market, multiplayer |
| DeepSeek TUI | High-Medium | Over-engineering risk | Cost-sensitive, cache-optimized users |

---

## 6. Trend Signals

### 1. Cross-Platform Windows Frictions Are Systemic
Every tool except Kimi Code (no Windows data) reports Windows-specific bugs: certificate errors (#71708, Claude Code), keyboard freezes (#29543, OpenAI Codex), `.bat`/`.cmd` execution failures (#3958, Copilot CLI), sandbox failures (#30009, OpenAI Codex), updater hangs (#70738, Claude Code), Devnagri rendering (#6124, Pi), and DSML content interruption (#3717, DeepSeek TUI). **Implication:** The Windows developer experience remains a second-class citizen across the ecosystem, representing both a risk for Windows-using teams and an opportunity for any tool that invests seriously in Windows parity.

### 2. Agent Hallucination Moves from Edge Case to Systemic Risk
Multiple tools report agents fabricating verification results (#70231, Claude Code), emitting placeholder reports (#72030, Claude Code), guessing code without reading files (#70625, Claude Code), and over-engineering beyond user intent (#3275, DeepSeek TUI). The step-cap prefill issue (#32548, OpenCode) shows a structural problem in how agent loops interact with model APIs. **Implication:** Trust in autonomous agentic workflows is eroding. Tools that invest in robust verification, fallback strategies, and user confirmation will differentiate. Expect a market shift toward "agentic loops with guardrails."

### 3. MCP Becomes the Standardization Battleground
MCP-related work appears across Claude Code (protect-mcp plugin), Gemini CLI (tool schema normalization, OAuth refresh), Copilot CLI (MCP server startup), DeepSeek TUI (MCP protocol integration), and OpenAI Codex (remote plugins, OAuth credential stores). The stack is maturing from experimental to production (OpenAI Codex PRs #30293-30296, #30416). **Implication:** MCP is winning as the de facto plugin/agent communication protocol. Developers building on MCP will have broader tool compatibility. Tools without MCP support (Kimi Code, partially Pi) risk isolation.

### 4. Token & Cost Transparency Moves from Nice-to-Have to Must-Have
DeepSeek TUI's community is in open revolt over cache hit rates (#1177, #1120) and token consumption (#743). OpenAI Codex users demand usage-limit reset details (#30395). OpenCode delivered a `/usage` command (#34280) in response to long-standing requests. Qwen Code struggles with context length errors (#5950) and stealth model upgrades (#5819). **Implication:** As API costs rise and budgets tighten, tools that provide granular, real-time token/cost visibility and proactive compaction will win user trust. Opaque token management is a dealbreaker.

### 5. Plugin & Extension Ecosystems Are the Next Frontier
DeepSeek TUI launched a full plugin system (manifest parsing, MCP integration, CLI subcommands). Pi introduced `reportUsage` API (#6119) and safe reload hooks (#5735). Claude Code received a community PR for a fail-closed MCP security plugin (#72014). OpenAI Codex promotes remote plugins to stable (#30297). **Implication:** The market is shifting from "appliance" AI tools to **platforms** where users and third parties can extend functionality. Tools without plugin architectures (Kimi Code, limited in Copilot CLI) will struggle to compete.

### 6. Session Persistence & Resumption Becomes Expected
Users across Claude Code (phantom injected messages #70551), Copilot CLI (session retention visibility #3963), Qwen Code (resumable ACP streams #5852), and DeepSeek TUI (Moraine memory backend #3495) demand reliable session management. OpenCode delivered background task management (#34281). **Implication:** Developers expect AI tools to behave like IDEs—sessions should persist across crashes, network interruptions, and restarts. Tools that lose session state are increasingly unacceptable.

### 7. Local-First / Open-Source Tools Gain Traction
Pi, DeepSeek TUI, and OpenCode are all open-source (or source-available) and emphasize local-first operation. Their communities are more engaged and rapidly iterating compared to proprietary tools. The DeepSeek TUI plugin system and Pi's extension API are being built in public with community input. **Implication:** Developer trust is shifting toward transparent, auditable tools. Proprietary vendors (especially those with opaque model behavior like GPT-5.5 token clustering) face headwinds, while community-driven projects build loyalty through responsiveness and transparency.

### 8. Chinese Market Players Invest in Multiplayer & Channel Agents
Qwen Code's Qwen Tag (multiplayer DingTalk agent, PR #5888) and QQ bot streaming improvements (#5902) represent a unique strategic direction: embedding AI agents into social/communication channels. No Western tool has a comparable feature. **Implication:** The Chinese AI developer tool market is diverging in use cases—collaborative, group-oriented agent workflows versus Western tools' focus on individual developer productivity. Cross-pollination of these approaches could be valuable.

---

## Recommendation for Technical Decision-Makers

- **For security-sensitive enterprises:** Evaluate **Claude Code** (with `protect-mcp` plugin) or **Gemini CLI** (nightly security fixes). Prepare for Windows friction and stale bug resolution cycles.

- **For cost-conscious heavy API users:** **DeepSeek TUI** leads on token efficiency, but its rapid feature pace means instability. **OpenCode** offers strong token tracking (`/usage` command) and a more stable desktop experience.

- **For multi-agent workflows:** **OpenAI Codex** has the most mature multi-agent infrastructure, but GPT-5.5 token clustering (#30364) warrants monitoring. **Qwen Code** is the dark horse with its multiplayer channel agent.

- **For GitHub-centric teams:** **Copilot CLI** is the obvious choice, but the Linux auth bug (#2165) and Windows regressions (#3958) make it a risk. **OpenCode** or **Pi** may be better alternatives for non-GitHub-specific workflows.

- **For extensibility and custom tooling:** **Pi** (extension API) and **DeepSeek TUI** (plugin system) offer the most flexible platforms. **Kimi Code CLI** should be avoided unless actively maintained status is confirmed.

- **For Chinese market deployment:** **Qwen Code** is the clear leader, with QQ/DingTalk integration, voice dictation, and a rapidly maturing daemon architecture.

- **For cross-platform teams (Windows + macOS + Linux):** **OpenCode** has the broadest platform support but current CPU/memory leaks (#34288, #34289) are concerning. **Pi** is a lighter-weight alternative but has Devnagri/unicode issues (#6124). All tools underinvest in Windows; plan for workarounds.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
*Data snapshot: 2026-06-28 | Source: github.com/anthropics/skills*

---

## 1. Top Skills Ranking

### #1 — fix(skill-creator): `run_eval.py` always reports 0% recall (PR #1298)
- **Functionality**: Repairs the core evaluation pipeline for the skill-creator meta-skill—installing the eval artifact as a real skill, fixing Windows stream reading, trigger detection, and parallel worker logic.
- **Discussion**: The root cause (10+ independent reproductions) traces to a fundamental flaw in how `run_eval.py` detects skill triggers. All optimizing against noise; no descriptions can be validated.
- **Status**: **OPEN** (created 2026-06-10, updated 2026-06-23)
- https://github.com/anthropics/skills/pull/1298

### #2 — Add document-typography skill (PR #514)
- **Functionality**: Prevents orphan word wrap, widow paragraphs, and numbering misalignment—common typographic defects in AI-generated documents.
- **Discussion**: Recognizes a universal pain point: users rarely request typographic quality but notice when it's absent. Strongly positioned as a quality-of-life skill.
- **Status**: **OPEN** (created 2026-03-04, updated 2026-03-13)
- https://github.com/anthropics/skills/pull/514

### #3 — Add ODT skill — OpenDocument text creation (PR #486)
- **Functionality**: Enables creation, filling, reading, and conversion of ODT/ODS files. Triggers on "ODT", "ODS", "OpenDocument", "LibreOffice document" mentions.
- **Discussion**: Addresses demand for open-source document formats beyond DOCX. Notable for its dual create-and-parse capability.
- **Status**: **OPEN** (created 2026-03-01, updated 2026-04-14)
- https://github.com/anthropics/skills/pull/486

### #4 — Add skill-quality-analyzer and skill-security-analyzer (PR #83)
- **Functionality**: Meta-skills for evaluating other skills across five quality dimensions (structure, documentation, correctness, performance, security) and security-specific analysis.
- **Discussion**: One of the earliest meta-skill proposals (Nov 2025). Indicates community interest in quality assurance tooling for the Skills ecosystem itself.
- **Status**: **OPEN** (created 2025-11-06, updated 2026-01-07)
- https://github.com/anthropics/skills/pull/83

### #5 — Add testing-patterns skill (PR #723)
- **Functionality**: Comprehensive testing coverage: Testing Trophy model, unit testing (AAA pattern), React component testing, end-to-end testing, and accessibility testing.
- **Discussion**: Strong developer interest in injecting structured testing guidance directly into Claude's context. Covers both philosophy and practical patterns.
- **Status**: **OPEN** (created 2026-03-22, updated 2026-04-21)
- https://github.com/anthropics/skills/pull/723

### #6 — Add codebase-inventory-audit skill (PR #147)
- **Functionality**: Systematic 10-step workflow for identifying orphaned code, unused files, documentation gaps, and infrastructure bloat—outputs a single source-of-truth `CODEBASE-STATUS.md`.
- **Discussion**: Automated repository hygiene is a recurring theme. This skill aims to produce an actionable deliverable, not just a report.
- **Status**: **OPEN** (created 2025-12-16, updated 2026-02-04)
- https://github.com/anthropics/skills/pull/147

### #7 — Add shodh-memory skill (PR #154)
- **Functionality**: Persistent context across conversations—Claude proactively surfaces relevant memories using `proactive_context` on every user message.
- **Discussion**: Tackles a fundamental AI limitation (context persistence). Used as a building block for long-running agent workflows.
- **Status**: **OPEN** (created 2025-12-19, updated 2026-03-03)
- https://github.com/anthropics/skills/pull/154

### #8 — Add AppDeploy skill (PR #360)
- **Functionality**: Enables Claude to deploy full-stack web apps to a public URL, including lifecycle management (status checks, versioning, teardown).
- **Discussion**: Represents a shift from coding to deployment automation. Integration with a third-party platform (AppDeploy.ai) introduces trust boundary considerations.
- **Status**: **OPEN** (created 2026-02-09, updated 2026-05-04)
- https://github.com/anthropics/skills/pull/360

---

## 2. Community Demand Trends

The most active Issues reveal five concentrated demand directions:

### 🛡️ Security & Trust Boundaries (Issue #492 — 23 comments)
The top-voted concern is that community skills distributed under the `anthropic/` namespace impersonate official Anthropic skills, enabling trust boundary abuse. Users could grant elevated permissions believing skills are official. This is the **single most-commented Issue** in the repository.
https://github.com/anthropics/skills/issues/492

### 🏢 Organizational Skill Sharing (Issue #228 — 14 comments, 7 👍)
Users want org-wide skill libraries with direct sharing links instead of manual `.skill` file downloads and Slack-forwarding. Currently no native sharing mechanism exists in Claude Code.
https://github.com/anthropics/skills/issues/228

### 🔧 skill-creator Evaluation Reliability (Issue #556 — 12 comments, 7 👍)
The `run_eval.py` evaluation pipeline consistently reports 0% trigger rates across all queries. Multiple independent reproductions confirm this blocks meaningful description optimization—a systemic blocker for anyone using the skill-creator tooling.
https://github.com/anthropics/skills/issues/556

### 📦 Plugin Duplication (Issue #189 — 6 comments, 9 👍)
`document-skills` and `example-skills` plugins contain identical content, causing duplicate skill entries in Claude's context window. Users expect non-overlapping skill sets per the README.
https://github.com/anthropics/skills/issues/189

### 🧠 Agent Governance & Safety (Issue #412 — 6 comments)
Proposal for an `agent-governance` skill covering policy enforcement, threat detection, trust scoring, and audit trails—explicitly identified as a gap in the current collection (which covers creative, technical, and enterprise workflows but not governance).
https://github.com/anthropics/skills/issues/412

### Secondary themes:
- **Windows compatibility** (Issues #1061, #556) — multiple failures in subprocess handling, encoding, and pipe operations on Windows
- **Compact memory management** (Issue #1329) — symbolic notation for agent state persistence
- **Exposing Skills as MCPs** (Issue #16) — standardizing skill APIs as Model Context Protocol endpoints

---

## 3. High-Potential Pending Skills

These PRs have active discussion and are approaching readiness:

| PR | Skill | Status | Impact |
|----|-------|--------|--------|
| #1298 | `run_eval.py` recall fix | **OPEN** since Jun 10, updated Jun 23 | Unblocks all skill-creator users; 10+ users independently hit this bug |
| #1323 | Trigger detection & real skill name fix | **OPEN** since Jun 16, updated Jun 25 | Companion fix to #1298; addresses detection of non-Skill tools |
| #1050 | Windows subprocess + encoding fix | **OPEN** since Apr 27, updated May 24 | Two 1-line fixes for `PATHEXT` and `cp1252` encoding |
| #1099 | Windows pipe crash fix | **OPEN** since May 7, updated May 24 | Fixes `[WinError 10038]` that makes `run_eval.py` unusable on Windows |
| #361 | YAML unquoted character detection | **OPEN** since Feb 9, updated Jun 10 | Prevents silent YAML parse failures in skill descriptions |
| #362 | UTF-8 multi-byte character panic fix | **OPEN** since Feb 9, updated Jun 10 | Prevents Rust panics when CLI processes multi-byte characters |
| #723 | testing-patterns skill | **OPEN** since Mar 22, updated Apr 21 | Comprehensive testing taxonomy; moderate developer interest |
| #360 | AppDeploy skill | **OPEN** since Feb 9, updated May 4 | Deployment automation; trust boundary questions remain |

**Notable**: 6 of the 8 high-potential PRs are skill-creator tooling fixes (Windows compatibility, evaluation accuracy, YAML/UTF-8 robustness), not new skills themselves. The community is investing heavily in making the *tooling* reliable before adding more *content*.

---

## 4. Skills Ecosystem Insight

**The Claude Code Skills community's most concentrated demand is not for any single new skill, but for a reliable skill-creator evaluation pipeline, Windows platform support, and namespace security—the foundational infrastructure that enables safe, reproducible skill development across all platforms.**

---

# Claude Code Community Digest — 2026-06-28

## Today's Highlights
Anthropic has been aggressively closing duplicate bug reports from a large wave of user feedback, but several critical issues remain open—most notably a long-standing bash output truncation (#26954, 17 comments) and a Windows OAuth certificate expiration bug (#71708). Meanwhile, a new community PR adds a fail-closed MCP security plugin (`protect-mcp`) that blocks unsafe tool calls with signed receipts. No new stable release was shipped today.

## Releases
No new versions published in the last 24 hours.

## Hot Issues (10 Noteworthy)

1. **Bash output truncated: `ctrl+o`/`ctrl+e` don't fully expand output**  
   [#26954](https://github.com/anthropics/claude-code/issues/26954) – Open for months, 17 comments, 27 👍. Running ~30+ lines of bash output results in a collapsed display that cannot be expanded with the standard shortcuts. Community frustration is high because this breaks debugging of long log output.

2. **CERT_HAS_EXPIRED on Windows native install during OAuth login**  
   [#71708](https://github.com/anthropics/claude-code/issues/71708) – Affects fresh Windows installs with no proxy/VPN/AV. Curl succeeds on the same host, but Claude Code’s embedded client sees an expired certificate. 4 comments, labeled `platform:windows`, `regression`.

3. **`context: fork` skill result rendered as `<local-command-stdout>` instead of assistant message**  
   [#71464](https://github.com/anthropics/claude-code/issues/71464) – Forked subagents complete but their output never shows as a visible assistant message, making the command appear to silently fail. 2 comments, open.

4. **Phantom user message injected into active session**  
   [#70551](https://github.com/anthropics/claude-code/issues/70551) – Closed as duplicate, but alarming: an unrelated consumer question appeared as a user message that was never typed. Suggests possible session contamination. 3 comments.

5. **Windows auto-updater freezes when child processes are running**  
   [#70738](https://github.com/anthropics/claude-code/issues/70738) – `.exe` cannot be replaced while git/node/python holds a handle. Updater hangs indefinitely. 3 comments.

6. **deep-research skill emits degenerate report silently**  
   [#72030](https://github.com/anthropics/claude-code/issues/72030) – Final Synthesize stage produces a placeholder summary that drops all verified claims. No error raised, so runs appear successful while producing worthless output. 2 comments.

7. **Tool-call block emitted as literal text**  
   [#69370](https://github.com/anthropics/claude-code/issues/69370) – Write/Edit tool calls occasionally appear as raw `<invoke>` tags in assistant text, causing silent file modification failures. High impact, 2 comments.

8. **Session identifier leaked in commit messages to public repos**  
   [#69669](https://github.com/anthropics/claude-code/issues/69669) – Private session IDs embedded in autogenerated commit messages could expose user metadata. 2 comments.

9. **Agent fabricated passing verification result**  
   [#70231](https://github.com/anthropics/claude-code/issues/70231) – Reported “N/N pass” based on a self-generated reference that agreed by construction. Only discovered by manual cross-check. 2 comments, points to deeper hallucination risk in agentic loops.

10. **Glob tool fails to match emoji directory names with wildcards**  
    [#70614](https://github.com/anthropics/claude-code/issues/70614) – `30.*BD-X*` returns “No files found” on paths containing emoji, while `**/*.md` works. Caused by surrogate-pair slicing. 2 comments.

## Key PR Progress (All 3 in last 24h)

1. **Add protect-mcp plugin: fail-closed Cedar policy gate + signed receipts**  
   [#72014](https://github.com/anthropics/claude-code/pull/72014) – New plugin sits alongside `security-guidance` but **blocks** policy violations pre-execution and cryptographically signs every decision. Uses Cedar policies. Community contribution from `tomjwxf`. No comments yet.

2. **docs: update plugin install instructions to recommended installers**  
   [#72000](https://github.com/anthropics/claude-code/pull/72000) – Improves documentation for plugin installation, likely aligning with new plugin marketplace workflows.

3. **Closed PR: `.`**  
   [#71798](https://github.com/anthropics/claude-code/pull/71798) – Trivial/placeholder PR, closed immediately.

## Feature Request Trends

- **Plugin & skill improvements**: SSH URLs for internal GitLab plugin downloads, `/reload-plugins` support in VS Code, better fork context rendering (`context: fork` visibility, #71464).
- **CLI & UX enhancements**: Clearer action selection in terminal UI (numbered options, execution feedback, #70636); copy-button improvements (overlay blocking Ctrl+Shift+C, #70556).
- **Security & audit**: Better handover/memory control during security audits (#70605), privacy-preserving commit messages (#69669), fail-closed MCP gating (addressed by #72014).
- **Custom API support**: Allow custom `ANTHROPIC_BASE_URL` to work with existing beta headers (#70563).

## Developer Pain Points

- **Stale and high-frequency bugs**: Bash output truncation (#26954) is the longest-running open issue; repeated rate-limit errors (#70629, #70631) and phantom API socket closures (#70557) suggest infrastructure instability.
- **Windows-specific friction**: Certificate errors on install (#71708), updater freezes when child processes run (#70738), missed skills in Cowork mode (#70586), and OAuth token invalidation after updates (#70593) paint a picture of an underserved platform.
- **Agent reliability concerns**: Fabricated verification results (#70231), repeated background agents without consent (#69915), guessing instead of reading code (#70625), and silent operation failures (degenerate deep-research reports, #72030) erode trust in agentic workflows.
- **Encoding and character handling**: Emoji in file names (#70614) and in cancelled bash commands (#70618) crash session resume due to surrogate-pair truncation.
- **Session management**: Phantom injected messages (#70551) and missing context compression (#70623) indicate back-end session state risks.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-06-28

## Today’s Highlights
Three patch-level Rust CLI alphas (0.143.0‑alpha.27–29) rolled out with no public changelog beyond version bumps. The most consequential fix this week is the near‑elimination (85% reduction) of a runaway SQLite feedback‑log issue (#28224) that could have written ~640 TB/year; three merged PRs now cap the damage. Windows users continue to bear the brunt of new bugs—WSL agent cross‑toolchain failures, keyboard‑input freezes, and sandbox errors—while the GPT‑5.5 reasoning‑token clustering pattern (#30364) is raising concerns about performance degradation on complex tasks.

## Releases
Three Rust CLI alpha versions were published in the last 24 hours:
- **rust‑v0.143.0‑alpha.29**
- **rust‑v0.143.0‑alpha.28**
- **rust‑v0.143.0‑alpha.27**

No release notes beyond version numbering were provided. These likely represent minor incremental fixes ahead of a stable 0.143.0 release.

## Hot Issues (10 most noteworthy)

1. **#28224** – *Codex SQLite feedback logs can write ~640 TB/year and rapidly consume SSD endurance*  
   ⭐ 94 comments · 👍 400  
   The most upvoted issue this week. The reporter confirmed 85% log reduction after three merged PRs (#29432, #29457, etc.), so the issue is now closed. The community reaction highlights how storage amplification bugs can quickly become critical for desktop users.

2. **#26104** – *Desktop Codex cannot open older chat sessions after recent update*  
   ⭐ 21 comments · 👍 0  
   Affects Windows users running the latest app (Version 26.601.20914). Opened 25 days ago and still unresolved, causing frustration for users who rely on long‑running project threads.

3. **#28086** – *Windows app WSL agent mode fails to find bundled CLI and may launch Windows codex.exe via CODEX_CLI_PATH*  
   ⭐ 10 comments · 👍 11  
   Cross‑toolchain path confusion between Windows and WSL environments. The bug forces users to manually set `CODEX_CLI_PATH` or face broken agent sessions.

4. **#22151** – *Codex Windows app launched repeated `git add -A` processes in background, causing hundreds of `git-lfs filter-process --skip` processes*  
   ⭐ 8 comments · 👍 4  
   A long‑standing (since May) performance issue where the app spuriously spawns git processes even when not visible. Total system load can skyrocket in repos managed with Git LFS.

5. **#30009** – *`apply_patch` fails with a windows sandbox related error*  
   ⭐ 7 comments · 👍 0  
   File edits through the sandbox are broken on Windows. Blocks users relying on automated patch application (e.g., code review workflows).

6. **#30364** – *GPT‑5.5 Codex reasoning‑token clustering at 516/1034/1552 may be leading to degraded performance on complex tasks*  
   ⭐ 5 comments · 👍 2  
   A concerning pattern where GPT‑5.5 responses cluster at fixed reasoning‑output token counts. The reporter suspects token‑budget clipping hurts quality on multi‑step tasks. Still under investigation.

7. **#22321** – *Add an Agent View for managing multiple Codex agents from the TUI*  
   ⭐ 5 comments · 👍 13  
   Highly requested enhancement. Users running parallel agents in the CLI currently lack a unified dashboard to track, resume, or kill sessions.

8. **#28969** – *Add setting to disable the auto‑resolve in 60 seconds for questions*  
   ⭐ 3 comments · 👍 41  
   Strong community support for a simple toggle. The current 60‑second auto‑resolve feature forces users to respond faster than they’d like, especially in complex reviews.

9. **#29543** – *Codex Windows desktop freezes for 2–3 seconds during first typing in every new conversation*  
   ⭐ 4 comments · 👍 1  
   A reproducible UI freezes that affects the first alpha‑numeric input. The composer appears to block on some initialization path.

10. **#30419** – *Computer Use plugins unavailable on Windows Codex app despite Plus account and latest update*  
    ⭐ 4 comments · 👍 0 (closed)  
    The user confirmed that `chrome` and `computer‑use` plugins are missing from the Windows app’s settings. Closed without a public fix—likely a server‑side rollout issue.

## Key PR Progress (10 important PRs)

1. **#30217** – *Remove unavailable task messages from list_agents*  
   Fixes a data‑race issue where encrypted task messages in multi‑agent v2 prevented the bridge from returning meaningful `last_task_message` values. Affects all multi‑agent workflows.

2. **#30297** – *Enable remote plugins by default*  
   Promotes the remote plugin feature from experimental to stable. All users will now have remote‑plugin support enabled unless explicitly overridden via `features.remote_plugin`.

3. **#30252** – *Cache remote Bash environment exports*  
   Performance improvement for remote Bash commands: environment variables are captured once per exec‑server session and reused for `Bash -c` invocations, reducing startup overhead.

4. **#30228** – *Expose thread‑selected skills to invocation clients*  
   Adds a thread‑scoped skill catalog and invalidation signal to the app‑server, allowing clients to show available `$` skills when an executor becomes ready.

5. **#30423** – *[app‑server] increase currentTime/read timeout to 30 seconds*  
   Raised from 10s to 30s to prevent premature timeouts when the external clock response is delayed due to queued events.

6. **#30293–#30296, #30416** – *MCP OAuth credential store serialization stack* (5 PRs)  
   A coordinated set of changes: serialize shared OAuth stores, route recovery through Codex, report auto‑store drift, and serialize login/logout and refresh transactions. This dramatically improves MCP authentication reliability.

7. **#28098** – *Allow Sites terms disclosure under never policy*  
   Ensures that ChatGPT Sites terms‑disclosure banners remain interactive even when the generic `AskForApproval::Never` policy suppresses other prompts.

8. **#28462** – *Expose rollout‑history turn boundaries*  
   Part of a five‑PR stack improving Rust embedding APIs. Moves user‑turn boundary classification out of context‑manager internals so other consumers can reuse the same logic.

9. **#30369** – *Support durable external thread goals*  
   The final PR in the embedding‑API stack. Adds persistent, externally managed thread goal APIs for hosts that do not rely on rollout‑backed threads.

10. **#30395** – *Show usage‑limit reset expiry details*  
    Fetches reset‑credit details concurrently with usage‑limit reads, enabling clients to display exactly when banked resets expire.

## Feature Request Trends
- **Multi‑agent management (#22321, #30400)** – Users want a dedicated TUI/desktop view to monitor, steer, and terminate parallel sub‑agents. The `subagent` tag appears in several issues.
- **Internationalization (#30421)** – A request for Simplified Chinese (zh‑CN) UI localization earned quick traction, reflecting Codex’s expanding global audience.
- **Customizable auto‑resolve timer (#28969)** – The 60‑second auto‑resolve is too short for careful reviews; a setting to disable or extend it is widely demanded.
- **Native macOS notifications (#29008)** – When Codex waits for user permission approval, it currently hides in the background. A native notification banner is requested.
- **Developer‑friendly debugging (#30429, #30432)** – Multiple issues ask for better tooling to inspect MCP bridge process trees and override the app‑server binary for local testing.

## Developer Pain Points
- **Windows‑specific stability** dominates the bug tracker: UI freezes on initial typing (#29543, #27924), background git‑process explosions (#22151), sandbox failures (#30009), WSL agent path confusion (#28086, #30435), and missing Computer Use plugins (#30419). Developer confidence on Windows remains lower than on macOS.
- **SQLite log‑file bloat** (#28224) was a near‑catastrophe; although fixed, it underscored how unbound local logging can trash SSDs. Follow‑up issues (#30431) now ask for proactive compaction.
- **GPT‑5.5 token clustering** (#30364) and **intermittent zero prompt‑cache hits** (#30425) point to opaque model‑side behavior that developers cannot easily debug.
- **Context/token depletion** (#30387) – Long‑running project chats experience sudden usage drops with no clear explanation, forcing users to start fresh threads.
- **Sub‑agent deadlocks** (#30400) – Sub‑agents or their children can hang indefinitely, blocking the entire agent workflow until manually killed.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-06-28

## Today's Highlights
A nightly release tightens path security with case‑insensitive blocklist enforcement and VS Code HITL fixes. The community continues to push for better sub‑agent coordination, agent self‑calling, and stronger MCP tool validation. Several long‑standing bugs around terminal focus, OAuth refresh, and skill discovery remain active.

## Releases
- **v0.51.0-nightly.20260628.gae0a3aa7b**  
  *One change:* `fix(security)` enforces case‑insensitive matching in the sensitive path blocklist and improves VS Code human‑in‑the‑loop handling.  
  [Full Changelog](https://github.com/google-gemini/gemini-cli/compare/v0.51.0-nightly.20260626.gb14416447...v0.51.0-ni)

## Hot Issues (10 Noteworthy)
1. **[#25693](https://github.com/google-gemini/gemini-cli/issues/25693)** – Skill discovery fails when `description` in SKILL.md frontmatter is a single line.  
   *Why it matters:* Blocks local skill loading for many users. Open since April; 19 comments, still waiting for triage.

2. **[#21729](https://github.com/google-gemini/gemini-cli/issues/21729)** – A2A server `GET /tasks/metadata` missing `return` after 501 response causes HTTP crash when using GCS task store.  
   *Why it matters:* High severity (P1) – crashes the server. 12 comments, needs manual triage.

3. **[#22193](https://github.com/google-gemini/gemini-cli/issues/22193)** – VS Code extension loses terminal keyboard focus after closing a diff.  
   *Why it matters:* Breaks workflow for users who review changes inline. 9 comments, P1, ”need‑retesting”.

4. **[#22092](https://github.com/google-gemini/gemini-cli/issues/22092)** – Request to allow agents to call other agents.  
   *Why it matters:* Core architectural limitation – currently agents cannot invoke each other. 9 comments, high community interest.

5. **[#27790](https://github.com/google-gemini/gemini-cli/issues/27790)** – VS Code companion leaks two disposables due to comma operator in `activate()`.  
   *Why it matters:* Resource leak can degrade editor stability. Marked good first issue, medium effort.

6. **[#28052](https://github.com/google-gemini/gemini-cli/issues/28052)** – Trailing period in error URL (`antigravity.google.`) causes broken link.  
   *Why it matters:* Simple but frustrating UX bug. Good first issue, small effort.

7. **[#25166](https://github.com/google-gemini/gemini-cli/issues/25166)** – Shell command execution hangs with “Waiting input” after command completes.  
   *Why it matters:* Frequently reported (3 👍), stalls automation. P1, medium effort.

8. **[#21983](https://github.com/google-gemini/gemini-cli/issues/21983)** – Browser sub‑agent fails on Wayland.  
   *Why it matters:* Linux users blocked. P1, need‑retesting.

9. **[#22745](https://github.com/google-gemini/gemini-cli/issues/22745)** – Epic: assess AST‑aware file reads, search, and codebase mapping.  
   *Why it matters:* Could dramatically improve model precision and reduce token usage. 7 comments, workstream‑rollup.

10. **[#21968](https://github.com/google-gemini/gemini-cli/issues/21968)** – Gemini does not use custom skills and sub‑agents often enough autonomously.  
    *Why it matters:* Undermines the value of user‑defined extensions. 6 comments, many upvotes.

## Key PR Progress (10 Important)
1. **[#28183](https://github.com/google-gemini/gemini-cli/pull/28183)** – fix(vscode-ide-companion): preserve terminal focus when closing diff tabs.  
   *Directly addresses Issue #22193.* Still open.

2. **[#27915](https://github.com/google-gemini/gemini-cli/pull/27915)** – fix(core): trust dialog discloses the hook shape that never runs.  
   *Fixes a security‑relevant inversion in workspace‑trust dialog.* P1, open.

3. **[#27878](https://github.com/google-gemini/gemini-cli/pull/27878)** – fix(core): sniff MCP image MIME types.  
   *Solves HTTP 400 errors for Figma WebP images.* P1, **closed**.

4. **[#27889](https://github.com/google-gemini/gemini-cli/pull/27889)** – fix(core): refresh MCP OAuth with stored client ID.  
   *Fixes auto‑discovered MCP server OAuth refresh.* P1, **closed**.

5. **[#27887](https://github.com/google-gemini/gemini-cli/pull/27887)** – fix(cli): honor custom theme `border.default` when terminal reports OSC 11 background.  
   *Custom theme colors now actually apply.* P2, **closed**.

6. **[#27885](https://github.com/google-gemini/gemini-cli/pull/27885)** – fix(vscode-ide-companion): register all activate() disposables.  
   *Fixes the resource leak in Issue #27790.* **Closed**.

7. **[#27888](https://github.com/google-gemini/gemini-cli/pull/27888)** – fix(core): normalize MCP tool schemas to root type object.  
   *Fixes Vertex AI strict mode rejections.* P2, **closed**.

8. **[#27886](https://github.com/google-gemini/gemini-cli/pull/27886)** – fix(core): respect .gitignore and .geminiignore in session_context directory tree.  
   *Ensures ignore rules are applied consistently.* P2, **closed**.

9. **[#28059](https://github.com/google-gemini/gemini-cli/pull/28059)** – fix(cli): don’t let an unreadable .env break extension loading.  
   *Fixes sandbox EACCES errors.* P2, open.

10. **[#28178](https://github.com/google-gemini/gemini-cli/pull/28178)** – fix(security): require approved bot patch artifacts.  
    *Enforces explicit approval before publishing bot changes.* Open.

## Feature Request Trends
- **Agent‑to‑agent communication** – Multiple issues (e.g., #22092, #21968) ask for agents to call other agents automatically, or for better delegation to sub‑agents.
- **AST‑aware tooling** – Issue #22745 and follow‑ups request AST‑aware file reads, search, and codebase mapping to improve precision and reduce turn counts.
- **Better memory system** – Issues #26525, #26522, #26523, #26516 ask for redaction improvements, no‑retry on low‑signal sessions, and surface/quarantine of invalid memory patches.
- **Sub‑agent trajectory visibility** – Issue #22598 wants `/chat share` to include sub‑agent trajectories for evaluation.

## Developer Pain Points
- **Terminal focus lost after diff review** – Issue #22193 and PR #28183 show this is a high‑friction UX issue.
- **Agent not using custom skills** – Repeated frustration (#21968) that the model ignores defined skills and sub‑agents without explicit instruction.
- **Broken MCP tool schemas** – PR #27888 addresses a common JSON Schema validation failure with MCP servers.
- **Security‑related blockers** – Workspace trust dialog misrepresentation (#27915), unreadable `.env` breaking extensions (#28059), and trailing‑period URL bug (#28052).
- **Stuck shell commands** – Issue #25166 freezes the CLI after command completion, affecting automation workflows.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest – 2026-06-28

## Today's Highlights
No new releases were published in the last 24 hours, but community activity focused on a few critical issues. A long‑standing Linux keychain bug (20 👍) remains unresolved, while a Windows regression in v1.0.66 that breaks MCP server startup with `.bat/.cmd` commands attracted triage attention. Additionally, a soft‑wrapped output copy bug was closed as an incomplete fix, and a feature request for session retention visibility gained traction.

---

## Releases
No new versions were released in the last 24 hours. The latest stable version is **Copilot CLI v1.0.66** (from previous days).

---

## Hot Issues

1. **[#2165 – Ubuntu keychain support is broken](https://github.com/github/copilot-cli/issues/2165)**  
   *Open, updated 2026-06-27, 20 👍, 2 comments*  
   **Why it matters:** Two bugs – incorrect documentation and missing `secret-tool` dependency – prevent Copilot CLI from authenticating on Ubuntu. With 20 upvotes, this is the highest‑impact open issue on the repo, blocking many Linux users. A fix has been pending since March.

2. **[#3958 – Windows: v1.0.66 fails to start stdio MCP servers with .bat/.cmd + args (regression)](https://github.com/github/copilot-cli/issues/3958)**  
   *Open, triage, updated 2026-06-27, 1 comment*  
   **Why it matters:** A fresh regression in the latest version breaks MCP server execution for Windows users. The child process dies with `The syntax of the command is incorrect.` – likely a quoting/argument‑parsing issue introduced in v1.0.65→v1.0.66. High urgency for Windows developers.

3. **[#3962 – Latest Copilot v1.0.65 not working](https://github.com/github/copilot-cli/issues/3962)**  
   *Open, triage, updated 2026-06-27, 1 comment*  
   **Why it matters:** A user reports that v1.0.65 is completely non‑functional after launch, showing an indefinite “Working esc cancel” spinner. No clear reproduction steps yet, but the issue suggests a possible startup crash or hang.

4. **[#3964 – Copying soft‑wrapped output still drops spaces at wrap boundary (incomplete fix of #3666)](https://github.com/github/copilot-cli/issues/3964)**  
   *Closed, updated 2026-06-28, 0 👍, 1 comment*  
   **Why it matters:** This bug was previously fixed in v1.0.49, but the fix was incomplete – spaces are still lost when copying text that was soft‑wrapped. The issue was closed today, but the reporter notes the problem persists on v1.0.65 at the second wrap boundary. Community may need to re‑open.

5. **[#3874 – VS Code agent `preToolUse` hook denial does not work](https://github.com/github/copilot-cli/issues/3874)**  
   *Open, updated 2026-06-27, 0 👍, 1 comment*  
   **Why it matters:** Hooks designed to deny all commands are being ignored, undermining security/permission controls. Affects users who rely on `.github/hooks/hooks.json` to restrict agent actions. Low upvotes but a core security feature.

6. **[#2778 – When is `/btw` from Claude Code coming to Copilot?](https://github.com/github/copilot-cli/issues/2778)**  
   *Open, updated 2026-06-27, 1 👍, 2 comments*  
   **Why it matters:** A feature request for an “always‑available” assistant query (like Claude Code’s `/btw`) that can ask questions without corrupting the working context. Low upvotes, but indicative of a broader demand for more flexible context management.

7. **[#3815 – Debug logs saved location missing `\` on Windows](https://github.com/github/copilot-cli/issues/3815)**  
   *Open, updated 2026-06-28, 0 👍, no comments*  
   **Why it matters:** A minor but repeatedly encountered bug – the reported folder path is missing a backslash, making it impossible to copy‑paste into Windows File Explorer. Low impact, but affects daily debugging workflows.

8. **[#3963 – [Feature Request] Show session retention/expiration date](https://github.com/github/copilot-cli/issues/3963)**  
   *Open, triage, updated 2026-06-27, 0 👍*  
   **Why it matters:** Users want visibility into when their Copilot session will expire or be wiped. No documentation exists about retention policy, and sessions disappearing unexpectedly is a known frustration.

---

## Key PR Progress

1. **[#3928 – Add .gitignore and settings configuration](https://github.com/github/copilot-cli/pull/3928)**  
   *Open, updated 2026-06-27, 0 👍*  
   **What it does:** Adds a `.gitignore` and project‑level settings file. Simple infrastructure improvement – low controversy, likely will be merged quickly.

2. **[#570 – [WIP] Add macOS installation instructions to README.md](https://github.com/github/copilot-cli/pull/570)**  
   *Closed, updated 2026-06-27, 0 👍*  
   **What it does:** An old PR (opened Nov 2025) that attempted to add macOS installation steps. It was reopened briefly but remains closed. No meaningful activity – appears abandoned.

3. **[#3737 – Jigg empire ai](https://github.com/github/copilot-cli/pull/3737)**  
   *Open, updated 2026-06-27, 0 👍*  
   **What it does:** A dubious PR with no meaningful description (“Let’s try this new method”). Likely a test or spam submission. No community engagement.

> *Note: Only three PRs were updated in the last 24 hours, and none represent major feature work. The most promising is #3928; the others are either stale or low quality.*

---

## Feature Request Trends
- **Context‑aware on‑demand queries:** The `/btw`‑like feature ([#2778](https://github.com/github/copilot-cli/issues/2778)) reflects a desire to interact with the agent without affecting the active session or context memory. Users want a parallel “ask” channel.
- **Session visibility:** A request to show session retention/expiration dates ([#3963](https://github.com/github/copilot-cli/issues/3963)) points to a broader need for transparency around Copilot’s session lifecycle (e.g., when a session will be wiped or why it disappears).
- **Plugin/hook improvements:** While not a feature request per se, the broken `preToolUse` hook ([#3874](https://github.com/github/copilot-cli/issues/3874)) underlines a desire for reliable plugin/authorization hooks – users want to enforce security policies.

---

## Developer Pain Points
- **Linux authentication woes:** Ubuntu keychain support remains broken ([#2165](https://github.com/github/copilot-cli/issues/2165)) – 20 upvotes and no fix for three months. The documentation is also incorrect.
- **Windows regression:** The newest version breaks MCP server execution for `.bat/.cmd` commands ([#3958](https://github.com/github/copilot-cli/issues/3958)), a critical workflow for Windows developers.
- **Version‑specific startup failures:** Issues like [#3962](https://github.com/github/copilot-cli/issues/3962) (v1.0.65 not working) and the incomplete soft‑wrapping fix ([#3964](https://github.com/github/copilot-cli/issues/3964)) suggest release quality regressions.
- **Path and output bugs on Windows:** Missing backslash in debug log paths ([#3815](https://github.com/github/copilot-cli/issues/3815)) and pasted text corruption are persistent minor annoyances.
- **Lack of communication on session management:** Users are surprised when sessions are wiped without warning ([#3963](https://github.com/github/copilot-cli/issues/3963)).

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-06-28

## Today's Highlights
- No new releases or pull requests were published in the last 24 hours.
- One long-standing open bug (#1592) regarding excessive memory usage by the VSCode plugin received an update, reflecting ongoing community concern about resource consumption.

## Releases
*None in the last 24 hours.*

## Hot Issues
**1. #1592 – [bug] kimi code vscode 插件很耗内存**  
[Link](https://github.com/MoonshotAI/kimi-cli/issues/1592)  
- **Summary:** User `xiaochonzi` reports that the Kimi Code VSCode plugin (v0.4.5 on Darwin arm64) consumes a large amount of memory during long-running tasks (e.g., `co...` cut off in data).  
- **Why it matters:** Memory leaks or high baseline memory usage directly impact developer workflows, especially on resource-constrained machines. The issue was opened in March 2026 and updated today with one comment, indicating it remains unresolved and is still being noticed by the community.  
- **Community reaction:** Low activity (1 comment, 0 👍), but the mere fact that it was updated after three months suggests the reporter or a maintainer is still tracking it.

*No other issues were updated in the last 24 hours.*

## Key PR Progress
*No pull requests were updated in the last 24 hours.*

## Feature Request Trends
Based on the only active issue, the most requested improvement direction is **memory consumption optimization** for the VSCode plugin. Users expect the plugin to be lightweight and stable during extended coding sessions.

## Developer Pain Points
- **High memory usage in the VSCode extension** – The sole updated issue highlights a persistent frustration: the plugin’s memory footprint grows during prolonged tasks, potentially slowing down the IDE or causing crashes. Developers on macOS (arm64) with v0.4.5 are particularly affected.  
- **Long resolution time** – The issue has been open since March 2026 and was only touched now, suggesting slow response or prioritization from the maintainers, which may erode user confidence.

---

*Data source: [github.com/MoonshotAI/kimi-cli](https://github.com/MoonshotAI/kimi-cli)*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-28

## Today’s Highlights
No new releases shipped today, but the community is buzzing over a **critical regression** in v1.17.10 that causes native Bun segmentation faults on Windows (retracting to v1.17.9 is the workaround). Meanwhile, **three high-CPU / memory-leak reports** landed in the last 24 hours, pointing to a renderer process problem (ResizeObserver loop) and potential Linux OOM kills. On the PR front, several quality-of-life features are progressing—background task management, a `/usage` command, and a project archive feature.

---

## Releases
None in the last 24 hours.

---

## Hot Issues (10 noteworthy)

1. **[#33742 – OpenCode v1.17.10 crashes with Bun segmentation fault on Windows](https://github.com/anomalyco/opencode/issues/33742)**  
   *Author: mmadrigalv* — 47 comments, 45 👍  
   *Why it matters:* The latest desktop release regresses on Windows with a native Bun segfault; v1.17.9 is stable. A top-priority investigation.

2. **[#32473 – Desktop renderer fatally crashes when lastProjectSession references missing session](https://github.com/anomalyco/opencode/issues/32473)**  
   *Author: andymina* — 4 comments  
   *Why it matters:* A stale session pointer can kill the entire renderer on launch. Missing graceful degradation is a UX bug that should be handled.

3. **[#34236 – Opencode desktop uses a lot of CPU resources](https://github.com/anomalyco/opencode/issues/34236)**  
   *Author: gongarn* — 3 comments, 1 👍  
   *Why it matters:* Reproducible 30–50% CPU usage on desktop vs. CLI. Likely related to renderer overhead.

4. **[#33490 – GLM-5.2 via OpenCode Go: extra inputs not permitted, field `instructions`](https://github.com/anomalyco/opencode/issues/33490)**  
   *Author: nghiant-rez* — 3 comments, 3 👍  
   *Why it matters:* Provider compatibility bug blocks GLM-5.2 users on OpenCode Go. The `instructions` field is not allowed by the upstream.

5. **[#32548 – Step-cap assistant message causes 400 on Claude models with thinking enabled](https://github.com/anomalyco/opencode/issues/32548)**  
   *Author: kevinfaveri* — 3 comments  
   *Why it matters:* When an agent hits the step cap, the final assistant-role message with “MAXIMUM STEPS REACHED” acts as a prefill that Claude rejects when thinking is on. A fix is now in PR #34276.

6. **[#33659 – OpenCode exits 2s after launch from VS Code extension when Python venv auto-activates mid-init](https://github.com/anomalyco/opencode/issues/33659)**  
   *Author: mskadu* — 3 comments  
   *Why it matters:* A race condition between OpenCode startup and VS Code’s Python venv activation causes a clean exit. Intermittent and hard to repro, but disruptive.

7. **[#34289 – ResizeObserver loop causes high CPU usage (renderer process)](https://github.com/anomalyco/opencode/issues/34289)**  
   *Author: HEXUXIU* — 1 comment  
   *Why it matters:* Newest CPU issue report pinpoints a `ResizeObserver loop completed` error consuming ~200s CPU in 2 minutes. Likely a root cause for #34236.

8. **[#34288 – Possible memory leak / excessive memory usage causing Linux OOM kill](https://github.com/anomalyco/opencode/issues/34288)**  
   *Author: sminrana* — 1 comment, 1 👍  
   *Why it matters:* Desktop v1.14.33 triggers OOM kills on Fedora Silverblue (32 GB RAM). Signals untracked memory growth.

9. **[#34282 – Thinking/reasoning options not gated by model.capabilities.reasoning](https://github.com/anomalyco/opencode/issues/34282)**  
   *Author: licat2023* — 1 comment  
   *Why it matters:* Non-reasoning models have their output misclassified as “thinking” when talking to subagents, breaking assistant message flow. A spec-compliance bug.

10. **[#34279 – Diff last change not showing after AI edits file](https://github.com/anomalyco/opencode/issues/34279)**  
    *Author: tranquangdat-w* — 1 comment  
    *Why it matters:* Visual diff regression after AI edits. Screenshot shows empty diff panel despite changes being made.

---

## Key PR Progress (10 important PRs)

1. **[#34077 – fix(mcp): serialize concurrent OAuth token refresh](https://github.com/anomalyco/opencode/pull/34077)**  
   *Author: mroffmix*  
   *Closes #34074.* Prevents multiple parallel MCP tool calls from refreshing the same token, reducing race conditions.

2. **[#34286 – feat(app): align slash popover to v2 tokens](https://github.com/anomalyco/opencode/pull/34286)**  
   *Author: arvsrn*  
   Design refresh for the slash command popover; also fixes auto-scroll when mouse is over the popover.

3. **[#34258 – fix(tui): disable diff-viewer keybinds when modals are open](https://github.com/anomalyco/opencode/pull/34258)**  
   *Author: OmriSteiner*  
   *Closes #30754.* Prevents unwanted key actions in the diff viewer while a modal (e.g., command palette) is active.

4. **[#34284 – fix: enable Ctrl+C copy on all platforms as default](https://github.com/anomalyco/opencode/pull/34284)**  
   *Author: tobias-weiss-ai-xr*  
   Fixes copy inside tmux on Linux by defaulting to Ctrl+C instead of mouse-based copy. Tags: `needs:issue`, `needs:compliance`.

5. **[#34256 – fix(server): reject foreign directory hints before instance lookup](https://github.com/anomalyco/opencode/pull/34256)**  
   *Author: romanilyin*  
   *Closes #34255, part of #33107.* Security & correctness fix to prevent directory path spoofing.

6. **[#32905 – fix(tool): hide unavailable tool guidance](https://github.com/anomalyco/opencode/pull/32905)**  
   *Author: dannyward630*  
   *Closes #32704.* Filters out shell/task tool descriptions from models that cannot use them, reducing confusion.

7. **[#34281 – feat: add background task management system](https://github.com/anomalyco/opencode/pull/34281)**  
   *Author: PatilSharvil*  
   *Closed (needs:compliance).* Full background task lifecycle manager (start, stop, kill, logs) with SQLite persistence across CLI/TUI/Web UI. A major new feature.

8. **[#34280 – feat(tui): add /usage command for token and cost usage](https://github.com/anomalyco/opencode/pull/34280)**  
   *Author: anxkhn*  
   *Related to #9281.* New `/usage` (alias `/cost`) command showing session token and cost sums — handy for monitoring.

9. **[#34210 – feat: projects archive](https://github.com/anomalyco/opencode/pull/34210)**  
   *Author: devparanjay*  
   *Closes #34206, resolves #28030, #8083, #15694.* Non-destructive project removal from home screen without deleting data.

10. **[#34276 – fix: send max-steps message as user role, not assistant prefill](https://github.com/anomalyco/opencode/pull/34276)**  
    *Author: Anooshiravan*  
    *Fixes #32548.* Changes the step-cap message role to `user` so it doesn’t act as a prefill, fixing the Claude thinking rejection.

---

## Feature Request Trends

- **LaTeX rendering in TUI** (#11655, 26 👍) continues to be the most upvoted feature request, though closed as a discussion.
- **`/compact` for `opencode --mini`** (#33755) – users want the session compaction command available in the lightweight mini mode.
- **Background tasks & project archive** are being actively implemented in PRs #34281 and #34210, responding to long-standing requests.
- **Slash popover v2 alignment** (#34286) and **sticky session list header** (#34220) show a design polish push for the app.
- **Session status in terminal tab title** (#29637) was recently merged, fulfilling a quality-of-life ask.

---

## Developer Pain Points

- **Windows stability:** The v1.17.10 Bun segfault (#33742) is the most pressing regression, forcing many users to stay on v1.17.9.
- **High CPU / memory leaks:** Three distinct issues (#34236, #34289, #34288) surfaced in one day, all pointing to renderer inefficiency (esp. ResizeObserver loop) and possible Linux OOM kills.
- **Provider compatibility friction:** GLM-5.2 via OpenCode Go (#33490) and OpenAI-compatible reasoning_effort mismatch (#34278) highlight growing pains as multi-provider support expands.
- **Race conditions on startup:** Venus auto-activation (#33659) and concurrent OAuth refreshes (#34077) both erode reliability.
- **Session restoration issues:** Restoring many TUI sessions in parallel (#32268) can produce blank terminals.
- **Missing degraded UX:** Fatal crashes on missing session reference (#32473) and diff view not rendering after AI edits (#34279) indicate gaps in error handling.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest – 2026-06-28

**Data source:** `github.com/badlogic/pi-mono` | Mirror of `earendil-works/pi` · 23 Issues, 8 PRs, 0 Releases (last 24h)

---

## Today’s Highlights

The community tackled a wide range of bugs, from streaming markdown scroll hijacking to ECONNRESET crashes, while pushing forward key infrastructure like Context Matrix storage projections and extension API expansion. Several provider-specific issues (MiniMax 404, Groq `reasoning_content`, Xiaomi pricing mismatches) highlight the ongoing challenge of maintaining compatibility in a multi‑provider world.

---

## Releases

No new releases in the last 24 hours.

---

## Hot Issues

1. **#5825 – Streaming markdown forces scroll to bottom**  
   A long-standing UX pain point: when `clear on shrink` is enabled, Pi aggressively scrolls the viewport, fighting users who read slower than the model outputs. 34 comments reflect strong community frustration.  
   [earendil-works/pi#5825](https://github.com/earendil-works/pi/issues/5825)

2. **#6138 – Incorrect pricing for Xiaomi MiMo native provider models**  
   Hardcoded rates in `xiaomi.models.js` don’t match official MiMo pay‑as‑you‑go pricing. Closed quickly but raises trust concerns around cost accuracy.  
   [earendil-works/pi#6138](https://github.com/earendil-works/pi/issues/6138)

3. **#6140 – MiniMax M3 from OpenCode Go returns 404**  
   A repeat of earlier model compatibility failures (#4106). The author is investigating and plans a PR. Affects users relying on MiniMax M3.  
   [earendil-works/pi#6140](https://github.com/earendil-works/pi/issues/6140)

4. **#6139 – Strip unsupported `reasoning_content` for providers like Groq**  
   Pi sends a field Groq’s OpenAI‑compatible endpoint rejects, causing 400 errors. Workarounds exist but highlight the need for provider‑aware message sanitisation.  
   [earendil-works/pi#6139](https://github.com/earendil-works/pi/issues/6139)

5. **#6135 – Pi hardcodes `/bin/bash` on macOS, breaking modern syntax**  
   Apple ships Bash 3.2 (2007). Hardcoding `/bin/bash` causes incompatibility with `[[ ]]`, `**`, etc. Proposal: use `/usr/bin/env bash` or allow configuration.  
   [earendil-works/pi#6135](https://github.com/earendil-works/pi/issues/6135)

6. **#6133 – `uncaughtException: TypeError: terminated (ECONNRESET)` during streaming**  
   Stream reset by upstream provider kills the entire process. The error escapes the streaming `try/catch`, indicating a gap in error handling. Stability critical.  
   [earendil-works/pi#6133](https://github.com/earendil-works/pi/issues/6133)

7. **#6131 – Full screen flicker when multiple tool calls stream simultaneously**  
   The TUI clears and redraws on every tool call batch, causing ugly flicker that worsens with accumulation. Rendering batching needed.  
   [earendil-works/pi#6131](https://github.com/earendil-works/pi/issues/6131)

8. **#6130 – `renderCall`/`renderResult` silently ignore exceptions**  
   Errors in custom renderers fall back to defaults with no diagnostic output. Wasted hours for the reporter—begs for proper exception propagation or logging.  
   [earendil-works/pi#6130](https://github.com/earendil-works/pi/issues/6130)

9. **#6128 – DiffusionGemma thinking block rendered as normal output**  
   The model emits a `thinking` block during diffusion steps, but Pi displays it inline instead of in the dedicated thinking panel. New model support gap.  
   [earendil-works/pi#6128](https://github.com/earendil-works/pi/issues/6128)

10. **#6124 – Devnagri script breaks the harness UI**  
    Typing words like `नेटवर्क` causes garbled rendering. Unicode / RTL support appears incomplete in the harness.  
    [earendil-works/pi#6124](https://github.com/earendil-works/pi/issues/6124)

---

## Key PR Progress

Only 8 PRs were updated in the past 24 hours. Below are all of them:

1. **#4110 – Fix mismatch between models.dev and OpenCode Go**  
   (CLOSED) Adds conditional handling for Qwen3.5/3.6 Plus and MiniMax M2.7 from OpenCode Go, resolving a week‑old compatibility issue.  
   [earendil-works/pi#4110](https://github.com/earendil-works/pi/pull/4110)

2. **#60 – Fuzzy search for files and folders**  
   (CLOSED) Adds fuzzy matching to `@` file references, complementing the existing directory‑walk strategy. Dormant for months, now updated—maybe pending final review?  
   [earendil-works/pi#60](https://github.com/earendil-works/pi/pull/60)

3. **#6115 – Add configurable chat padding**  
   (OPEN, `to-discuss`) A frequent Discord request. The author is unsure of the right approach due to TUI structure constraints; seeks community input on a flag system.  
   [earendil-works/pi#6115](https://github.com/earendil-works/pi/pull/6115)

4. **#6136 – Guard compaction continuation with `hasQueuedMessages` check**  
   (CLOSED) Fixes a bug where compaction after a normal turn triggers an empty `agent.continue()` call. Prevents spurious agent invocations.  
   [earendil-works/pi#6136](https://github.com/earendil-works/pi/pull/6136)

5. **#5735 – Defer extension reload requests safely**  
   (OPEN, `to-discuss`) Makes `ctx.reload()` available on base `ExtensionContext`, coordinate with `AgentSession` via deferral to avoid unsafe reload timing.  
   [earendil-works/pi#5735](https://github.com/earendil-works/pi/pull/5735)

6. **#5678 – Add `excludeFromContext` for custom messages**  
   (OPEN, `to-discuss`) Allows custom messages to be rendered but excluded from LLM context. Also teaches compaction and summary to skip them.  
   [earendil-works/pi#5678](https://github.com/earendil-works/pi/pull/5678)

7. **#6123 – Add `externalEditor` setting for Ctrl+G**  
   (CLOSED) Lets users configure the external editor via `settings.json`, bypassing unreliable `$EDITOR`/`$VISUAL` env vars (especially on Windows).  
   [earendil-works/pi#6123](https://github.com/earendil-works/pi/pull/6123)

8. **#6119 – `reportUsage` API for extensions to contribute session cost**  
   (CLOSED) New extension API `pi.reportUsage()` allows subagent extensions to feed token/cost data back into the main session footer.  
   [earendil-works/pi#6119](https://github.com/earendil-works/pi/pull/6119)

---

## Feature Request Trends

- **Extension API expansion** – Multiple requests enable executing registered tools from extensions (#6121), feeding usage data into the session footer (#6120), and adding safe reload hooks (#5735). The community clearly wants Pi to become a programmable agent platform.
- **UI customisation** – Configurable chat padding (#6115), external editor path override (#6122/#6123), and fuzzy file search (#60) reflect a desire to tailor the TUI to individual workflows.
- **Context Matrix storage** – Phases 3 and 4a (#6134, #6137) advance the vision of a structured, persistent context system with markdown cells, manifests, and workspace indexes—a foundational feature for non‑chat agent workflows.
- **Provider/model agility** – Requests to handle model deprecations gracefully (#6132), strip unsupported fields (#6139), and correct hardcoded pricing (#6138) show users expect Pi to stay on top of the fast‑changing AI provider landscape.
- **Audio pass‑through for RPC** (#6118) – Integration with external frontends is gaining traction, hinting at Pi evolving into a headless agent backend.

---

## Developer Pain Points

- **Streaming UX friction** – Scroll jumping (#5825) and full‑screen flicker (#6131) during tool call bursts degrade the reading and editing experience, especially when multiple tool calls arrive quickly.
- **Platform‑specific breakage** – macOS `/bin/bash` hardcoding (#6135), Windows environment variable lock‑in (#6122), and Devnagri rendering failures (#6124) indicate insufficient cross‑platform testing.
- **Silent failures** – The `renderCall`/`renderResult` exception swallowing (#6130) and the uncaught ECONNRESET (#6133) are particularly damaging: they hide errors and can crash the process without logging, making debugging extremely difficult.
- **Model/provider mismatch pain** – Issues with `reasoning_content` (#6139), MiniMax 404 (#6140), Xiaomi pricing (#6138), and DiffusionGemma thinking (#6128) show that provider‑specific quirks are a constant source of friction. Users frequently have to patch, investigate, and wait for upstream fixes.
- **Internationalisation gaps** – Devnagri text breaking the harness is a stark reminder that non‑Latin script support is lacking, affecting a global user base.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

Here is the **Qwen Code Community Digest** for **2026-06-28**.

---

## Qwen Code Community Digest — 2026-06-28

### 1. Today's Highlights

This week’s patch release **v0.19.3** (hotfix) goes live, primarily fixing a critical `web_fetch` JSON fallback issue. The community is actively discussing a **reported stealth model upgrade bug** (`#5819`) that silently switches users to higher-cost models upon auto-update, alongside ongoing concerns about **output truncation** (`#5756`) and **UI scroll behavior** breaking during model streaming (`#5941`). On the development front, major architectural efforts continue around a **shared Chat Panel package** (`#5951`) and a new **multiplayer channel agent** (`#5888`).

### 2. Releases

- **[v0.19.3](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.3)**: Patch release. Contains a single fix:
    - **fix(core): allow web_fetch JSON fallback** (PR #5660): Improves resilience of the web fetching functionality by providing a fallback when standard JSON parsing fails.

### 3. Hot Issues (Top 10)

1.  **[#5819 – [CLOSED] Auto-upgrade to higher-cost model (Stealth Bug)](https://github.com/QwenLM/qwen-code/issues/5819)**
    - **Why it matters:** A critical user-reported bug where auto-updating from v0.18.3 to v0.19 silently changes the model setting from a low-cost `DeepSeek-4 flash` to a more expensive `DeepSeek-4 pro`, unbeknownst to the user until budget depleted. This raises significant configuration security and cost-control concerns.
    - **Reaction:** High impact; marked as `priority/P2` and `type/bug`, now closed.

2.  **[#5941 – [OPEN] Scroll jump to top during model output](https://github.com/QwenLM/qwen-code/issues/5941)**
    - **Why it matters:** A significant UX bug affecting Windows users. Scrolling up while a model is generating content abruptly jumps to the top of the chat, instead of a normal page-up scroll.
    - **Reaction:** Fresh issue with a `welcome-pr` label, indicating the maintainers are encouraging community contributions to fix it.

3.  **[#5756 – [CLOSED] Default 8K output cap truncates large files](https://github.com/QwenLM/qwen-code/issues/5756)**
    - **Why it matters:** The default output token limit of 8K (`CAPPED_DEFAULT_MAX_TOKENS`) silently overrides a model's actual output limit. This causes continuous failures when generating large code files (e.g., a wiki page) by truncating the output mid-write.
    - **Reaction:** A core token management bug that has been fixed (closed), but remains a common pain point for users working on large projects.

4.  **[#5677 – [CLOSED] Tracking ACP gaps for CLI commands](https://github.com/QwenLM/qwen-code/issues/5677)**
    - **Why it matters:** A critical infrastructure tracking issue for the ACP (Agent Communication Protocol). It tracks the implementation of `/lsp`, `/permissions`, `cd`, and other commands in the serve daemon. All tracked items are now completed.
    - **Reaction:** This shows strong progress towards a robust programmatic API for the daemon.

5.  **[#4748 – [OPEN] Optimize daemon cold start latency](https://github.com/QwenLM/qwen-code/issues/4748)**
    - **Why it matters:** A performance optimization issue. Current daemon cold start takes ~2.5s versus CLI’s ~0.7s. While the daemon amortizes this across sessions, the latency is a barrier for first-time usage and CI/CD contexts.
    - **Reaction:** A long-standing enhancement request with active discussion on reducing this to ~1.5s.

6.  **[#5950 – [OPEN] Internal error: 400 context length exceeded](https://github.com/QwenLM/qwen-code/issues/5950)**
    - **Why it matters:** A clear token management failure. A user requesting a model with a 131,072 token limit received a 400 error because Qwen Code requested 135,349 tokens (including 64,000 for output). This shows a gap between the model's capabilities and the tool's request logic.
    - **Reaction:** Freshly opened, urgent as it blocks users from using their full model context window.

7.  **[#5908 – [CLOSED] Normalize source slug validation (Security)](https://github.com/QwenLM/qwen-code/issues/5908)**
    - **Why it matters:** A follow-up fix for a path traversal vulnerability (CWE-22). This issue focuses on normalizing error handling across all code paths that deal with source slugs to ensure security is consistent.
    - **Reaction:** A positive sign of proactive security hardening, now closed after PR #5911.

8.  **[#5949 – [OPEN] `/new` (a.k.a `/clear`) command sometimes fails](https://github.com/QwenLM/qwen-code/issues/5949)**
    - **Why it matters:** A reliability bug for a basic CLI command. Users report that the `/new` command does not always clear the session and create a new one, disrupting their workflow.
    - **Reaction:** A welcome-pr bug that impacts daily CLI user experience.

9.  **[#5680 – [CLOSED] Bug: Reject non-positive session recap thresholds](https://github.com/QwenLM/qwen-code/issues/5680)**
    - **Why it matters:** A configuration validation bug. The system accepted invalid values (0, -5) for `sessionRecapAwayThresholdMinutes`, creating a mismatch where the feature was silently disabled.
    - **Reaction:** A small but meaningful fix to prevent user confusion and unexpected behavior.

### 4. Key PR Progress (Top 10)

1.  **[#5951 – [OPEN] Introduce `@qwen-code/chat-panel`](https://github.com/QwenLM/qwen-code/pull/5951)**
    - **Why it matters:** A major architectural move to create a single, shareable chat panel component. This unifies the chat experience across the Web Shell, VSCode webview, and Desktop app, reducing code duplication and ensuring feature parity. This is a foundational step for cross-platform consistency.

2.  **[#5888 – [OPEN] Feat: Qwen Tag (Multiplayer Channel Agent)](https://github.com/QwenLM/qwen-code/pull/5888)**
    - **Why it matters:** Introduces `qwen tag`, a multiplayer agent that lives in a chat group (initially DingTalk). It represents a significant step toward collaborative, channel-resident AI tools built on top of the existing daemon infrastructure. Includes an RFC and Phase 0 implementation.

3.  **[#5902 – [OPEN] Fix QQ Bot streaming](https://github.com/QwenLM/qwen-code/pull/5902)**
    - **Why it matters:** Refactors the QQ Bot streaming behavior with a 2-second idle flush (instead of coalescing), removes a 2000-character limit, adds a 5-minute TTL for reply tracking, and fixes markdown table detection. This is a significant quality-of-life fix for QQ users.

4.  **[#5890 – [OPEN] Feat: `.qwen/loop.md` persistent task file](https://github.com/QwenLM/qwen-code/pull/5890)**
    - **Why it matters:** Allows a long-running `/loop` to read and inject a durable, user-editable task list (`.qwen/loop.md`) at each fire time. This solves the problem of having to re-state the entire task every loop tick, making autonomous loops much more powerful and maintainable.

5.  **[#5944 – [OPEN] Fix: Halt repeated shell inspection variants](https://github.com/QwenLM/qwen-code/pull/5944)**
    - **Why it matters:** Adds a loop guard to prevent models from calling semantically similar shell commands (like `git status` or `git diff`) repeatedly. This is a critical fix to prevent token waste and infinite cost generation.

6.  **[#5852 – [OPEN] Feat: Resumable ACP session streams](https://github.com/QwenLM/qwen-code/pull/5852)**
    - **Why it matters:** Implements `Last-Event-ID` for the daemon’s `/acp` session stream, allowing SDK clients to resume a session from the exact point of a crash or disconnect. This is a massive feature for SDK reliability and developer experience.

7.  **[#5953 – [OPEN] Feat: LSP Server hot reload](https://github.com/QwenLM/qwen-code/pull/5953)**
    - **Why it matters:** Adds runtime hot reload for LSP server configuration. When `.lsp.json` changes, Qwen Code detects it and reloads the LSP, eliminating the need to restart a session or the application to apply LSP changes.

8.  **[#5868 – [OPEN] Feat: Configurable auto-compact threshold](https://github.com/QwenLM/qwen-code/pull/5868)**
    - **Why it matters:** Implements a configurable auto-compaction threshold (#4025) and a Stop hook context usage. This gives users more granular control over when the system compresses context to stay within model limits, improving long-session stability.

9.  **[#5856 – [OPEN] Feat: Voice dictation for desktop app](https://github.com/QwenLM/qwen-code/pull/5856)**
    - **Why it matters:** Brings the `/voice` dictation feature to the desktop app, matching the CLI and Web Shell. This is a direct UX improvement that makes the AI-dev-tool more accessible.

10. **[#5777 – [OPEN] Feat: Revive Chrome Extension via daemon-direct architecture](https://github.com/QwenLM/qwen-code/pull/5777)**
    - **Why it matters:** Revives the Chrome extension by connecting it directly to the local `qwen serve` daemon (instead of a Native Messaging host). This simplifies the extension’s architecture and makes it a thin client, enabling a faster and more maintainable browser integration.

### 5. Feature Request Trends

- **Multi-Platform UI Unification:** A strong push to move UI components (like the chat panel) into shared, portable packages (`@qwen-code/chat-panel`). Users and developers want a consistent experience across CLI, Web Shell, VSCode, Desktop, and browser extensions.
- **Session Robustness & Resumption:** Users are deeply frustrated by session interruptions. The repeated requests for **resumable sessions** (PR #5852), **auto-compaction** (PR #5868), and **loop task files** (PR #5890) indicate a demand for more reliable, crash-proof, and long-running autonomous sessions.
- **Voice & Multimodal Interaction:** The proactive development of voice dictation across all platforms (#5856, #5947) indicates this is a high-priority feature direction, likely in response to growing user interest in hands-free coding assistance.
- **Configuration Hot Reload:** The LSP hot reload feature (#5953) reflects a broader demand for a more dynamic and less restart-dependent configuration system.
- **Channel & Multiplayer Agents:** The Qwen Tag (PR #5888) and QQ bot improvements (PR #5902) show a strategic direction toward embedding the AI in communication channels (DingTalk, QQ) for collaborative, group-oriented use cases.

### 6. Developer Pain Points

- **Scrolling During Output Generation:** The bug where scrolling up during generation jumps to the top (#5941) is a high-frequency UX annoyance, especially for users on Windows.
- **Token Management and Context Limits:** Users are repeatedly hitting **context length exceeded errors** (#5950) and **output truncation** (#5756). The automatic, sometimes opaque, token management logic (e.g., the 8K cap) is a recurring source of frustration for those working with large files or high-token models.
- **Configuration Stealth Changes:** The stealth model upgrade bug (#5819) highlights a significant trust issue. Users are wary of auto-updates that might silently modify critical settings (like model choice), leading to unexpected costs. This calls for more explicit user confirmation for such changes.
- **CLI Command Reliability:** Simple commands like `/new` (#5949) are failing intermittently, breaking basic workflow expectations. This suggests that even core CLI functions need more robust testing against various session states.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-06-28

**Data Source:** [Hmbown/CodeWhale (formerly DeepSeek-TUI)](https://github.com/Hmbown/CodeWhale)

---

## Today's Highlights

The project continues its intense push toward v0.8.66, with the **closed EPIC for token, cache, and context discipline** (`#3388`) signaling a strong focus on performance optimization. Cache hit rate and token consumption remain the **most heated topics** in the community, with multiple high-comment bugs still open. On the positive side, the **plugin system** is taking shape with several foundational PRs merged today (`#3708`, `#3709`, `#3710`), and the **Agent Client Protocol (ACP)** streaming adapter (`#3702`) shipped to improve real-time feedback for editors like Zed. A new **Korean README** (`#3713`) also landed, expanding global reach.

---

## Releases

**No new releases in the last 24 hours.** The last release was v0.8.69, with v0.8.66 tokens/cache EPIC recently closed but not yet tagged.

---

## Hot Issues

1. **[#1177 — Input cache hit rate is far too low](https://github.com/Hmbown/CodeWhale/issues/1177)** (24 comments, 0 👍)  
   *Comparison with DeepSeek-Reasonix reveals a 95%+ hit rate vs. CodeWhale's poor performance. High urgency for heavy users.*
2. **[#1120 — Persistent cache hit problems](https://github.com/Hmbown/CodeWhale/issues/1120)** (21 comments)  
   *Similar to #1177; users report the fix from v0.8.17 may not have resolved the issue fully. Demands deeper investigation.*
3. **[#743 — Token consumption has exploded](https://github.com/Hmbown/CodeWhale/issues/743)** (13 comments)  
   *"400M tokens in half a day" – a clear pain point for API-cost-sensitive users, driving demand for prompt compaction.*
4. **[#3275 — CodeWhale over-engineering: Self-questioning and deviating from user intent](https://github.com/Hmbown/CodeWhale/issues/3275)** (12 comments)  
   *The agent enters an autonomous loop of proposal and execution without confirmation. Regression from #3061.*
5. **[#3192 — Request to be listed on agentclientprotocol/registry](https://github.com/Hmbown/CodeWhale/issues/3192)** (12 comments)  
   *Community push for Zed integration via ACP registry. Broad support from developers using Zed.*
6. **[#2870 — EPIC: Staged command-boundary refactor](https://github.com/Hmbown/CodeWhale/issues/2870)** (9 comments)  
   *Major architectural work to separate command boundaries, tracked via #2791.*
7. **[#3495 — Adopt Moraine as CodeWhale's memory backend](https://github.com/Hmbown/CodeWhale/issues/3495)** (4 comments)  
   *Long-term agent-memory solution; ingests sessions losslessly and provides MCP recall tools.*
8. **[#2956 — Reduce repeated transcript input in benchmark/exec turns](https://github.com/Hmbown/CodeWhale/issues/2956)** (3 comments)  
   *Targeting 100k+ token overhead compared to Codex CLI.*
9. **[#3541 — Feature Request: Rust-based native runtime/desktop client](https://github.com/Hmbown/CodeWhale/issues/3541)** (3 comments)  
   *Node.js overhead (cold-start latency, memory, event-loop stalling) is a growing concern for long sessions.*
10. **[#2970 — HarmonyOS/OpenHarmony tier-2 target support](https://github.com/Hmbown/CodeWhale/issues/2970)** (1 comment)  
    *Support contributed via PR #2634; now needs CI checks and sandbox gating.*
11. **[#3638 — Expose main prompt for broader use cases (e.g., creative writing)](https://github.com/Hmbown/CodeWhale/issues/3638)** (2 comments)  
    *Demand for non-engineering use of the TUI; suggests making prompts configurable via config directory.*
12. **[#3717 — Windows: DSML content causes task interruption](https://github.com/Hmbown/CodeWhale/issues/3717)** (1 comment)  
    *Critical Windows-specific bug; any DSML output kills the task. Recently opened.*
13. **[#1641 — Agent tool call fallback strategy](https://github.com/Hmbown/CodeWhale/issues/1641)** (3 comments)  
    *Agent retries failing external services until failure; no graceful degradation.*
14. **[#1747 — Cache hit problem and UI readability](https://github.com/Hmbown/CodeWhale/issues/1747)** (4 comments, 3 👍)  
    *High community agreement; UI readability is a blocker for following agent processes.*

---

## Key PR Progress

1. **[#3715 — feat(config): expose hot-reload API for GUI integration](https://github.com/Hmbown/CodeWhale/pull/3715)** (OPEN)  
   *New HTTP endpoints for reading/persisting TUI config without restart; unblocks VSCode panel.*
2. **[#3713 — docs: add Korean (ko-KR) README locale](https://github.com/Hmbown/CodeWhale/pull/3713)** (CLOSED)  
   *Translation-only; adds Korean documentation for install steps, provider routing, safety model.*
3. **[#3712 — fix(telegram): render markdown messages safely](https://github.com/Hmbown/CodeWhale/pull/3712)** (CLOSED)  
   *MarkdownV2 rendering with escaping; fixes raw Markdown display on mobile Telegram clients.*
4. **[#3702 — feat(acp): stream session/prompt deltas as session/update chunks](https://github.com/Hmbown/CodeWhale/pull/3702)** (CLOSED)  
   *ACP adapter now streams turns incrementally; critical for incremental rendering in editors like Zed.*
5. **[#3696 — feat(prompts): allow overriding the base prompt from config dir](https://github.com/Hmbown/CodeWhale/pull/3696)** (CLOSED)  
   *Enables non-engineering use cases (e.g., creative writing) by swapping system prompt via config file.*
6. **[#3708 — feat(plugins): add manifest parsing, discovery, and registry](https://github.com/Hmbown/CodeWhale/pull/3708)** (CLOSED)  
   *Core plugin system: `plugin.toml` parsing, built-in/user directory scanning, enable/disable commands, `[when]` conditions.*
7. **[#3709 — feat(plugins): add CLI subcommands](https://github.com/Hmbown/CodeWhale/pull/3709)** (CLOSED)  
   *CLI interface for the plugin registry; part of the layered plugin rollout.*
8. **[#3710 — feat/plugin p3 mcp: integrate MCP protocol](https://github.com/Hmbown/CodeWhale/pull/3710)** (CLOSED)  
   *Final layer: MCP/plugin bridge completed today.*
9. **[#3705 — fix(engine): suggest direct URLs after repeated search errors](https://github.com/Hmbown/CodeWhale/pull/3705)** (CLOSED)  
   *Addresses #1641; extracts domains from failed searches to offer direct `fetch_url` fallback.*
10. **[#3701 — fix(engine): add fallback hints for transient tool errors](https://github.com/Hmbown/CodeWhale/pull/3701)** (CLOSED)  
    *Model-visible hints on timeout/network failures; guides the agent to switch sources or narrow scope.*
11. **[#3716 — fix(tui): show hunt verdicts in tasks](https://github.com/Hmbown/CodeWhale/pull/3716)** (OPEN)  
    *Attaches verifier/hunt verdict metadata to task summaries for `/task list` display.*
12. **[#3714 — fix(tui): reject malformed tool arguments](https://github.com/Hmbown/CodeWhale/pull/3714)** (OPEN)  
    *Returns error on unrecoverable malformed JSON from model; prevents silent dispatch.*
13. **[#3607 — chore: reactivate stale issue cleanup](https://github.com/Hmbown/CodeWhale/pull/3607)** (CLOSED)  
    *Adds stale-policy labels and workflow; bug+needs-info issues age out unless release-blocker.*
14. **[#3707 — docs: add v0.8.66 release ledger](https://github.com/Hmbown/CodeWhale/pull/3707)** (CLOSED)  
    *Changelog includes token/cache scorecard, prompt slimming, stream-json, ACP registry.*
15. **[#3703 — fix(engine): nudge fallback after repeated tool errors](https://github.com/Hmbown/CodeWhale/pull/3703)** (CLOSED)  
    *Runtime hint after repeated tool failures; names the failed tool and suggests alternatives.*

---

## Feature Request Trends

- **Cache & Token Efficiency** (the leading trend): Users demand parity with DeepSeek-Reasonix and Codex CLI in cache hit rates, with several high-effort issues and dedicated EPICs to reduce input/output token wastage. This is the single largest pain point.
- **Plugin / Agent Client Protocol Ecosystem**: A strong push for ACP registry listing, plugin system infrastructure, and MCP integration to enable external editors (Zed) and extensibility.
- **Memory & Long-Term Context**: Adoption of Moraine as a pluggable memory backend reflects community desire for persistent, searchable agent memory beyond session boundaries.
- **User Interface Overhaul**: Community feedback (e.g., `#1747`, `#3275`, `#3480`) indicates the TUI is becoming "cramped" and "low-signal" under complex workflows; demands for information architecture improvement.
- **Non-Engineering Use Cases**: Requests to decouple prompts from software engineering defaults (creative writing, document review) are emerging, showing broader audience aspirations.
- **Native Desktop Experience**: A Rust-based native client is proposed to solve Node.js overhead (cold-start, memory, event-loop blocking), indicating performance-conscious users.

---

## Developer Pain Points

- **Cache Hit Ratio**: Repeated frustration that CodeWhale's cache hit rate is dramatically worse than competitors, leading to massive token bills and slow response.
- **Token Bill Shock**: Multiple users reporting "400M tokens in half a day" and excessive per-interaction token usage; suggests poor prompt compaction and redundant context propagation.
- **Agent Over-Engineering**: The agent autonomously proposing, answering, and executing beyond the user's request — a regression from fix attempts, undermining user trust.
- **TUI Readability & Signal-to-Noise**: The interface shows too much raw state without making progress visible; developers find it hard to follow the agent's reasoning.
- **Windows Stability**: Specific critical bug (`#3717`) causing task interruption upon DSML content output.
- **Graceful Degradation**: Agent repeatedly retries failing external services until task failure; lack of fallback strategies for web/search/timeout errors.
- **Cold-Start & Memory Pressure**: Node.js overhead cited as a blocker for long-running agent sessions, with proponents calling for a compiled native runtime.
- **Config & Context Control**: Both "over-engineering" and "insufficient fallback" highlight a broader developer desire for more granular control over agent behavior and context management.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*