# AI CLI Tools Community Digest 2026-07-06

> Generated: 2026-07-06 13:05 UTC | Tools covered: 9

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

# Cross-Tool Comparison Report: AI CLI Ecosystem — 2026-07-06

## 1. Ecosystem Overview

The AI CLI tool landscape is rapidly maturing, with nine actively maintained projects serving distinct developer workflows—from Claude Code’s deep agentic sessions to OpenAI Codex’s enterprise-grade desktop app and Gemini CLI’s research-oriented platform. A clear bifurcation is emerging: tools optimizing for autonomous agent behavior (Claude Code, Gemini, DeepSeek TUI, OpenCode) versus those prioritizing developer productivity within existing IDE/Git workflows (Copilot CLI, Qwen Code, Kimi Code). Across all tools, memory management, safety filter reliability, and cross-platform consistency remain the most persistent pain points. The community is increasingly vocal about cost visibility, session lifecycle control, and the need for programmatic APIs (ACP/MCP) to enable custom integrations.

---

## 2. Activity Comparison

| Tool | Hot Issues (Noteworthy) | Active PRs (Last 24h) | Release Status (Last 24h) |
|---|---|---|---|
| **Claude Code** | 10 | 1 | No new release |
| **OpenAI Codex** | 10 | 10 | No new release |
| **Gemini CLI** | 10 | 10 | v0.51.0-nightly (automated) |
| **GitHub Copilot CLI** | 10 | 0 | v1.0.69-2 (patch) |
| **Kimi Code CLI** | 2 | 0 | No new release |
| **OpenCode** | 10 | 10 | No new release |
| **Pi** | 10 | 10 | No new release |
| **Qwen Code** | 10 | 10 | v0.19.6-nightly (PR triage fix) |
| **DeepSeek TUI** | 10 | 10 | No new release |

**Key observations:**
- **Copilot CLI** is the only tool with a stable patch release today (v1.0.69-2); others rely on nightly/automated builds.
- **Kimi Code CLI** shows near-zero community activity—only 2 issues and 0 PRs—indicating either very low adoption or a closed development model.
- **OpenAI Codex** and **OpenCode** have the highest PR throughput (10 each), reflecting active engineering teams.
- **DeepSeek TUI** and **Gemini CLI** show strong issue volume with heavy focus on pre-release (v0.8.x) and structural refactoring.

---

## 3. Shared Feature Directions

The following requirements appear across **multiple** tool communities, indicating ecosystem-wide developer needs:

### Session & Lifecycle Management
- **Session continuation after limit** (Claude Code #13354, OpenAI Codex #30918/#31125, Qwen Code #6383, OpenCode #28695)
- **Agent pause/resume** (Claude Code #74689, Gemini CLI #21409, OpenCode #32767)
- **Background task visibility** (Claude Code, OpenCode, Qwen Code #6383, DeepSeek TUI #4039)

### Safety Filter Reliability
- **False-positive blocks on legitimate tasks** (Claude Code issues #74743/#74615, multiple Fable 5 complaints)
- **Constitution compliance concerns** (DeepSeek TUI #4032 — ignoring user-provided scripts)
- **Destructive behavior prevention** (Gemini CLI #22672 — `git reset --force`, Qwen Code #6396)

### Programmatic APIs & Integration
- **ACP/MCP protocol support for custom tooling** (Kimi Code #2486, OpenCode #35550, Copilot CLI #3028, OpenAI Codex #31163)
- **Plugin/hook systems for lifecycle events** (OpenCode #16626/#28695, Pi #6350, Copilot CLI #1665)
- **Usage/billing data exposure via API** (Kimi Code #2486, OpenAI Codex #30918, Qwen Code #5964)

### Cross-Platform Parity
- **Windows shell/line-ending issues** (Claude Code #2805/#14828, OpenAI Codex #30009/#30882, Gemini CLI #21983, OpenCode #35536/#35545, Qwen Code #6298)
- **Linux desktop app demand** (OpenAI Codex #11023 — 692👍, Claude Code #4953 — memory leak on Linux)
- **Wayland/WSL compatibility** (Gemini CLI #21983, Pi #6187)

### Cost & Token Visibility
- **Real-time token consumption metrics** (OpenCode #5374, Pi #6352/#6353, Qwen Code #6264)
- **Rate-limit transparency** (OpenAI Codex #30918/#31125, Kimi Code #2486)
- **Context window calculation fixes** (Qwen Code #6384, Gemini CLI #8132)

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | Copilot CLI | OpenCode | Pi | Qwen Code | DeepSeek TUI | Kimi Code |
|---|---|---|---|---|---|---|---|---|---|
| **Primary Target** | Power user agents | Enterprise desktop | Research platform | Git workflow | Plugin ecosystem | Provider flexibility | Code review/PR | Agent orchestration | IDE integration |
| **Platform Focus** | Linux/macOS | Windows/Mac | Cross-platform | All major | All major | All major | All major | All major | Windows |
| **Model Alignment** | Anthropic (Sonnet/Opus) | OpenAI (GPT-5.5) | Gemini models | GitHub/Copilot | Multi-provider | Multi-provider (30+) | QwenLM | Multi-provider | MoonshotAI (Kimi) |
| **Agent Style** | Deep sessions, sub-agents | Sandboxed, deterministic | Sub-agent delegation | Specialized agents | Plugin-driven | Edit-tool focused | PR triage pipeline | Conductor agents | Lightweight CLI |
| **Safety Approach** | Fable 5 filters (aggressive) | Permission prompts | Sub-agent isolation | Hook-based | Configurable | Strict tool schemas | CI-aware gates | Constitution enforcement | Minimal |
| **Key Differentiator** | Agentic depth, but memory leak | Enterprise auth, but Linux gap | Research eval pipeline | GitHub integration | Plugin seams (VCS, hooks) | Provider diversity | Review intelligence | Workflow orchestration | ACP integration |

### Notable Strategic Differences:

- **Agent Autonomy vs. Determinism**: Claude Code and DeepSeek TUI push for highly autonomous agents with sub-agent orchestration, while OpenAI Codex and Copilot CLI emphasize sandboxed, deterministic operations suitable for enterprise compliance.

- **Provider Strategy**: Pi (30+ providers) and OpenCode differentiate on provider-agnostic flexibility. Claude Code and OpenAI Codex are tied to their respective model families. Gemini CLI and Qwen Code each prioritize their own models but add third-party support.

- **Safety Philosophy**: Claude Code’s Fable 5 is generating the most community backlash for false positives. Gemini CLI takes a permission-based sub-agent approach. Pi is moving toward strict tool schemas (PR #6341) to prevent hallucinated tool arguments. DeepSeek TUI uses "constitution" enforcement.

- **Development Stage**: DeepSeek TUI (v0.8.x) and Gemini CLI (v0.51.x) are pre-1.0, with active architectural refactoring. Claude Code and OpenAI Codex are mature but dealing with regression debt. Kimi Code appears earliest-stage based on community activity.

---

## 5. Community Momentum & Maturity

| Tool | Community Sentiment | Iteration Speed | Risk Level |
|---|---|---|---|
| **Claude Code** | Volatile: high engagement but high frustration (memory leak, Fable 5) | Slow: 1 PR/day, core bugs unresolved for months | High: trust erosion from safety overreach |
| **OpenAI Codex** | Active: 10 PRs/day, strong enterprise interest | Moderate: steady fixes, new features | Moderate: auth/integration gaps |
| **Gemini CLI** | Engaged: pre-1.0 enthusiasm, but agent behavior concerns | High: nightly builds, active refactoring | Moderate: hangs and sub-agent issues |
| **Copilot CLI** | Stable: patch release today, focused issues | Moderate: one new release, low PR activity | Low: mature, incremental improvements |
| **OpenCode** | Growing: active PR pipeline, community-driven features | High: 10 PRs/day, strong plugin focus | Low: well-architected, modular |
| **Pi** | Active: steady issue resolution, provider expansion | High: 10 PRs/day, multi-provider focus | Moderate: null-content crashes, WSL issues |
| **Qwen Code** | Committed: nightly builds, active triage | High: 10 PRs/day, review pipeline focus | Low: focused on specific use case |
| **DeepSeek TUI** | Enthusiastic: pre-1.0, feature-rich roadmap | High: 10 PRs/day, architectural changes | Moderate: pre-release instability |
| **Kimi Code** | Nascent: minimal community activity | Low: stalled | High: unclear roadmap, low adoption |

### Maturity Clusters:

- **Mature (1.0+)**: Claude Code, OpenAI Codex, Copilot CLI — established but accumulating technical debt
- **Growing (0.5–1.0)**: OpenCode, Pi, Qwen Code — active development with clear roadmaps
- **Early (pre-0.8)**: Gemini CLI, DeepSeek TUI, Kimi Code — still defining architecture

---

## 6. Trend Signals

### 1. The Rise of Agent Orchestration
Multiple tools (DeepSeek TUI #4010, Gemini CLI #21409, Claude Code #74689) are investing in **conductor agents** that manage sub-agent ensembles. This reflects a shift from single-turn assistants to multi-step, parallelized workflows. DeepSeek’s "Workflow" (née WhaleFlow) and Qwen Code’s multi-workspace daemon (#6378) are the most ambitious examples. **For developers**: expect agent orchestration to be a key differentiator in 2027.

### 2. Safety Fatigue
The backlash against Claude Code’s Fable 5 (multiple issues blocking legitimate work) mirrors concerns across Gemini CLI (#22672) and DeepSeek TUI (#4032). Aggressive safety filters are generating **trust erosion** rather than protection. The industry trend appears to be moving toward **configurable safety levels** (Copilot CLI’s permission hooks, Pi’s strict tool schemas) rather than blanket filters. **For developers**: prioritize tools that offer graduated safety controls.

### 3. Cost Transparency Becomes Non-Negotiable
Across OpenAI Codex (#30918), Qwen Code (#6264/#5964), OpenCode (#5374), and Pi (#6352/#6353), users demand real-time token consumption, rate-limit visibility, and zombie-session detection. The "invisible billing" problem—where users discover unexpected charges hours later—is a universal pain point. **For developers**: tools with built-in cost dashboards will have a competitive advantage.

### 4. Cross-Platform Parity as a Barrier to Adoption
Windows path escaping (OpenCode #35536, Claude Code #2805), Wayland browser subagent failures (Gemini CLI #21983), and Linux desktop app absence (OpenAI Codex #11023) are blocking significant user segments. The market is fragmenting: Windows-centric (Kimi Code), Linux-optimized (Claude Code), multi-platform (OpenCode, Pi). **For developers**: if you work in heterogeneous environments (WSL, remote SSH), prioritize tools with explicit platform compatibility testing.

### 5. The API-First Tooling Shift
Kimi Code (#2486), OpenCode (#35550), and Copilot CLI (#3028) are all pushing for **public APIs to expose tool internals** (usage data, session state, tool catalogs). This mirrors the broader industry move toward composable AI tooling—developers want to build custom dashboards, CI integrations, and IDE plugins on top of CLI tools. **For developers**: evaluate a tool’s ACP/MCP surface area, not just its CLI features.

### 6. Memory and Context Management Are Under-Engineered
Claude Code’s 120GB memory leak (#4953), Gemini CLI’s retrying low-signal sessions (#26522), and OpenCode’s CPU-burning idle waits (#19466) all point to fundamental gaps in how these tools manage long-running state. The community is demanding **session lifecycle hooks** (OpenCode #28695), **context budget trimming** (DeepSeek TUI #4015), and **AST-aware file reads** (Gemini CLI #22745) to reduce token waste. **For developers**: tools that solve the context-window efficiency problem intelligently (vs. just throwing more RAM at it) will win long sessions.

### 7. Model Lock-In vs. Flexibility
The contrast between Pi (30+ providers, strict schema enforcement) and Claude Code (Anthropic-only) represents a strategic fork. While provider diversity gives users flexibility, it introduces consistency challenges (Pi’s null-content crashes with GLM-5.2). **For developers**: make an explicit decision—do you want deep integration with one model family (better reliability, fewer edge cases) or multi-model flexibility (more options, more bugs to manage)?

### Reference Value for Developers

| If you need... | Consider... |
|---|---|
| Deep agentic sessions with sub-agent orchestration | Claude Code (but watch memory), DeepSeek TUI (pre-1.0 but ambitious) |
| Enterprise-grade deterministic operations | OpenAI Codex (sandboxed), Copilot CLI (Git integration) |
| Multi-provider flexibility | Pi (30+ providers), OpenCode (plugin ecosystem) |
| PR/review automation | Qwen Code (triage pipeline), Copilot CLI (GH integration) |
| Lightweight IDE integration | Kimi Code (ACP focus), Copilot CLI (VS Code synergy) |
| Research/experimental agent behavior | Gemini CLI (eval pipeline), DeepSeek TUI (conductor agents) |
| Maximum control and extensibility | OpenCode (plugin seams, lifecycle hooks), Pi (strict schemas) |

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
*Data as of 2026-07-06 | Source: github.com/anthropics/skills*

---

## 1. Top Skills Ranking

The following Skills (Pull Requests) have attracted the most community discussion and attention:

### 1.1 `skill-creator` Fixes: run_eval Trigger Detection
**PR [#1298](https://github.com/anthropics/skills/pull/1298)** — *Open* | Author: MartinCajiao | Created: 2026-06-10  
**Functionality:** Fixes the core evaluation pipeline (`run_eval.py` → `run_loop.py` → `improve_description.py`) which has been reporting `recall=0%` for every skill description under test. The PR addresses: installing the eval artifact as a real skill (so Claude actually loads it), fixing Windows stream reading and subprocess pipe handling, correcting trigger detection logic that ignores non-Skill tools, and parallel worker reliability.  
**Discussion Highlights:** This is the convergence point for multiple community pain points (Issues #556, #1169, #1061). The `recall=0%` bug has been independently reproduced 10+ times. Discussion centers on whether the root cause is the `-p` flag not invoking skills, or the detection algorithm comparing against the wrong target.  
**Status:** Open, active revisions through June 23.

### 1.2 Document-Typography
**PR [#514](https://github.com/anthropics/skills/pull/514)** — *Open* | Author: PGTBoos | Created: 2026-03-04  
**Functionality:** Prevents orphan word wrap (1–6 words on a new line), widow paragraphs (headers stranded at page bottom), and numbering misalignment in AI-generated documents. Addresses a class of typographic issues that affect virtually every document Claude generates.  
**Discussion Highlights:** Community strongly supports this as filling a universal quality gap. No controversy—the value proposition is clear.  
**Status:** Open, no recent updates since March 13.

### 1.3 PDF Case-Sensitivity Fix
**PR [#538](https://github.com/anthropics/skills/pull/538)** — *Open* | Author: Lubrsy706 | Created: 2026-03-06  
**Functionality:** Corrects 8 file-reference mismatches in `skills/pdf/SKILL.md` where uppercase filenames (`REFERENCE.md`, `FORMS.md`) are referenced but the actual files are lowercase. This breaks on case-sensitive filesystems (Linux, macOS with case-sensitive volumes).  
**Discussion Highlights:** Small fix but high-impact—catches a systemic issue in cross-platform skill distribution.  
**Status:** Open, no recent updates.

### 1.4 ODT Skill (OpenDocument Format)
**PR [#486](https://github.com/anthropics/skills/pull/486)** — *Open* | Author: GitHubNewbie0 | Created: 2026-03-01  
**Functionality:** Enables Claude to create, fill, read, and convert OpenDocument Format files (.odt, .ods). Triggers on mentions of "ODT", "ODS", "ODF", "OpenDocument", "LibreOffice". Fills a gap for the open-source document ecosystem.  
**Discussion Highlights:** High interest from LibreOffice-heavy workflows and EU public sector users. Discussion covers template-filling API design and ODT-to-HTML conversion quality.  
**Status:** Open, last updated April 14.

### 1.5 Self-Audit Skill (v1.3.0)
**PR [#1367](https://github.com/anthropics/skills/pull/1367)** — *Open* | Author: YuhaoLin2005 | Created: 2026-06-28  
**Functionality:** A universal output-auditing skill that performs mechanical file verification (do all claimed output files exist?) followed by a four-dimension reasoning quality audit ordered by damage severity. Works across any project, tech stack, or model.  
**Discussion Highlights:** Very recent (June 28). Community interested in the "universal" claim and how it interacts with existing model self-correction mechanisms. Some debate on whether audit dimensions are comprehensive enough.  
**Status:** Open, actively updated through July 2.

### 1.6 Testing-Patterns
**PR [#723](https://github.com/anthropics/skills/pull/723)** — *Open* | Author: 4444J99 | Created: 2026-03-22  
**Functionality:** Comprehensive testing stack coverage: testing philosophy (Trophy model), unit testing (AAA pattern, naming conventions, edge cases), React component testing (Testing Library), integration testing, E2E testing, property-based testing, and accessibility testing.  
**Discussion Highlights:** Well-received as filling a major gap in the collection. Discussion around breadth vs. depth—some argue it should be split into multiple specialized skills.  
**Status:** Open, last updated April 21.

### 1.7 Sensory Skill (macOS AppleScript Automation)
**PR [#806](https://github.com/anthropics/skills/pull/806)** — *Open* | Author: AdelElo13 | Created: 2026-03-29  
**Functionality:** Teaches Claude to use `osascript` (AppleScript) for native macOS automation instead of screenshot-based computer use. Two-tier permission system: Tier 1 works out of the box; Tier 2 requires Accessibility permissions.  
**Discussion Highlights:** Significant interest from Mac power users. Security discussion around permission tiers and whether the skill adequately warns users about accessibility risks.  
**Status:** Open, last updated April 2.

---

## 2. Community Demand Trends

From the most-discussed Issues, the community is requesting:

### 2.1 Security & Trust Boundaries (Issue [#492](https://github.com/anthropics/skills/issues/492) — 34 comments, 👍2)
**Demand:** Community skills distributed under the `anthropic/` namespace enable trust boundary abuse. Users may grant elevated permissions to skills they believe are official. This is the #1 most-commented issue.  
**Direction:** Demand for a community skills namespace, security review gates, or explicit provenance labeling. Some contributors are proposing a `skill-security-analyzer` (PR #83) as a meta-skill to address this.

### 2.2 Organizational Skill Sharing (Issue [#228](https://github.com/anthropics/skills/issues/228) — 14 comments, 👍7)
**Demand:** org-wide skill sharing within Claude.ai. Currently users must download `.skill` files and share via Slack/Teams. Requested: shared skill libraries or direct sharing links.  
**Direction:** Enterprise deployment features, centralized skill management, team collaboration workflows.

### 2.3 Skill Creator Reliability (Issue [#556](https://github.com/anthropics/skills/issues/556) — 12 comments, 👍7)
**Demand:** `run_eval.py` produces 0% trigger rate across all queries, making the description-optimization loop useless. Multiple independent reproductions.  
**Direction:** This is the community's most urgent **bug**. Fixes are actively being attempted across multiple PRs. The conversation focuses on whether the `claude -p` flag can invoke commands at all, and whether the evaluation artifact installation method is correct.

### 2.4 Skill Persistence & Stability (Issue [#62](https://github.com/anthropics/skills/issues/62) — 10 comments, 👍2)
**Demand:** Skills disappearing after file renames; error states not recoverable by users.  
**Direction:** Demand for skill state management, more robust file handling, and user-facing error recovery instructions.

### 2.5 Compact Agent Memory (Issue [#1329](https://github.com/anthropics/skills/issues/1329) — 9 comments, 👍0)
**Demand:** A `compact-memory` skill that uses symbolic notation for compact agent state, reducing context consumed by long-running agent's notes and persistent memory.  
**Direction:** Growing interest in skill-level support for agent context optimization and memory management.

### 2.6 Agent Governance (Issue [#412](https://github.com/anthropics/skills/issues/412) — 6 comments, 👍0, *Closed*)
**Demand:** Safety patterns for AI agent systems—policy enforcement, threat detection, trust scoring, audit trails.  
**Direction:** Even though closed, the discussion signals desire for governance skills alongside the technical and creative ones already in the collection.

### 2.7 Windows Compatibility (Issue [#1061](https://github.com/anthropics/skills/issues/1061) — 3 comments, 👍1)
**Demand:** `skill-creator` scripts fail on native Windows (Python 3) due to Unix-first assumptions: subprocess `PATHEXT` handling, `cp1252` encoding, `select()` on pipes.  
**Direction:** A persistent pain point for Windows users. Multiple PRs address sub-problems (#1050, #1099, #1298). The community wants a comprehensive Windows compatibility PR rather than piecemeal fixes.

### Summary of Demand Trends
| Theme | Signal Strength | Priority |
|---|---|---|
| Security namespace/trust | Very High (34 comments) | Critical |
| Skill reliability/testing | High (12 comments + multiple PRs) | Critical |
| Org sharing | High (14 comments, 7 upvotes) | High |
| Context optimization | Medium (9 comments) | Medium |
| Windows compatibility | Medium (3 comments, expanding) | Medium |

---

## 3. High-Potential Pending Skills

These open PRs have active discussion and are likely to merge soon:

| Skill | PR | Created | Last Updated | Why It's Likely to Land |
|---|---|---|---|---|
| `skill-creator` trigger detection fix | [#1298](https://github.com/anthropics/skills/pull/1298) | 2026-06-10 | 2026-06-23 | Fixes the #1 community bug (recall=0%). Multiple contributors converging. |
| `self-audit` (v1.3.0) | [#1367](https://github.com/anthropics/skills/pull/1367) | 2026-06-28 | 2026-07-02 | Very recent, active updates. Addresses growing quality-gate demand. |
| `color-expert` | [#1302](https://github.com/anthropics/skills/pull/1302) | 2026-06-10 | 2026-06-12 | Self-contained, well-specified, no controversy. Author is experienced contributor (@meodai). |
| `sensory` (macOS AppleScript) | [#806](https://github.com/anthropics/skills/pull/806) | 2026-03-29 | 2026-04-02 | High interest from Mac users. Two-tier permission model addresses security concerns. |
| `testing-patterns` | [#723](https://github.com/anthropics/skills/pull/723) | 2026-03-22 | 2026-04-21 | Fills major gap in collection. May need splitting but the core is solid. |
| `document-typography` | [#514](https://github.com/anthropics/skills/pull/514) | 2026-03-04 | 2026-03-13 | Low controversy, high universal value. May need maintainer re-engagement. |
| `self-audit` mechanical verification | [#1367](https://github.com/anthropics/skills/pull/1367) | 2026-06-28 | 2026-07-02 | Already listed above—re-emphasizing as dual-track (dev + reasoning audit). |

**Close Watch:** The `ODT` skill (PR #486) and `frontend-design` revision (PR #210) have strong foundations but need maintainer attention to advance.

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is for a reliable, cross-platform skill-creation and evaluation infrastructure, with secondary demand surging for universal output quality assurance (audit, testing, typography) and security governance.**

---

**Claude Code Community Digest — 2026-07-06**

*Data source: [github.com/anthropics/claude-code](https://github.com/anthropics/claude-code)*

---

### Today’s Highlights
A severe memory leak causing OOM kills at 120+ GB RAM continues to dominate community attention (#4953, 96 comments, 72 👍). Meanwhile, multiple users report that safety filters (Fable 5) are blocking legitimate code review and system administration tasks, with several users unable to use the tool at all for days. The session limit enforcement also remains a top frustration, with a long-standing feature request to allow continuation after the limit (#13354, 67 comments, 157 👍) still open.

---

### Releases
No new releases in the last 24 hours.

---

### Hot Issues (10 Noteworthy)

1. **[#4953 – Memory Leak – Process Grows to 120+ GB RAM and Gets OOM Killed](https://github.com/anthropics/claude-code/issues/4953)**  
   *🟢 Open, Bug, 96 comments, 72 👍*  
   The top-voted issue affecting Linux users in extended sessions. Despite being reported nearly a year ago, the leak remains unresolved. Community members have shared repro steps and JSONL logs.

2. **[#13354 – Feature Request: Continue When Session Limit Reached](https://github.com/anthropics/claude-code/issues/13354)**  
   *🟢 Open, Enhancement, 67 comments, 157 👍*  
   Highest 👍 count on any issue. Users hit daily session limits mid-task and want a seamless “continue” option rather than starting fresh or paying for a new subscription tier.

3. **[#69238 – No Response from API Error When Advisor Is Triggered](https://github.com/anthropics/claude-code/issues/69238)**  
   *🟢 Open, Bug, 41 comments, 66 👍*  
   Using Sonnet as base model, triggering the Advisor (Opus 4.8) causes repeated “No response from API” retry loops. Suggests a deeper issue with model switching or API routing.

4. **[#2805 – Claude Code Creates Files with Windows Line Endings on Linux](https://github.com/anthropics/claude-code/issues/2805)**  
   *🟢 Open, Bug, 39 comments, 32 👍*  
   Despite explicit `.claude.md` instructions, CRLF line endings are written on Linux. Breaks shell scripts and causes “No such file or directory” errors. A long-standing quality-of-life issue.

5. **[#14828 – Windows: Console Window Flashing When Executing Tools](https://github.com/anthropics/claude-code/issues/14828)**  
   *🟢 Open, Bug, 39 comments, 33 👍*  
   A visual annoyance that disrupts workflow on Windows. Each tool execution spawns and hides a console window.

6. **[#67606 – Opus 4.8 Confabulates User Messages and Fake “Prompt Injection” Narratives in Long Sessions](https://github.com/anthropics/claude-code/issues/67606)**  
   *🟢 Open, Bug, 10 comments, 0 👍*  
   Two independently verified sessions where Opus 4.8 fabricated user messages and described a fake attack. Raises concerns about model reliability in extended contexts.

7. **[#67071 – Assistant Text Between Tool Calls Not Rendered in GUI or CLI](https://github.com/anthropics/claude-code/issues/67071)**  
   *🟢 Open, Bug, 8 comments, 4 👍*  
   A regression amplified by Fable 5: text between tool calls is invisible in the UI, though persisted in session JSONL. Hampers debugging and trust.

8. **[#58904 – Heredoc Pipe Bypass Still Present in v2.1.141](https://github.com/anthropics/claude-code/issues/58904)**  
   *🔴 Closed, Bug, 7 comments, 1 👍*  
   A permission bypass allowed by heredoc markers. The fix for the CPU-loop symptom did not fix the underlying security issue. Closed as a known regression.

9. **[#74743 – Artifact HTML Creation Skill Incorrectly Blocks Local Repository Analysis](https://github.com/anthropics/claude-code/issues/74743)**  
   *🔴 Closed, Bug, 1 comment*  
   A safety filter (Fable) falsely flagged legitimate codebase exploration. Quickly closed, but users report feeling blocked from basic tasks.

10. **[#74761 – `claude -p` Exits 0 Mid-Task Without Completing Agent Loop](https://github.com/anthropics/claude-code/issues/74761)**  
    *🟢 Open, Bug, 1 comment*  
    The non-interactive `-p` mode exits success (exit code 0) while the agent is still mid-operation – tool results are delivered but no follow-up API call is made. Breaks CI/CD pipelines that rely on `-p` for deterministic output.

---

### Key PR Progress

Only one PR was updated in the last 24 hours:

- **[#74722 – feat(commit-commands): Support Conventional Branch naming in /commit-push-pr](https://github.com/anthropics/claude-code/pull/74722)**  
  *🟢 Open, Enhancement*  
  Adds an optional `conventional` argument to automatically name branches per the [Conventional Branch 1.0.0](https://conventionalbranch.org/) spec (e.g., `feature/<description>`, `bugfix/<description>`). Type is inferred from the diff. Useful for teams enforcing branch naming conventions.

No other pull requests were submitted or updated. Community PR activity remains low relative to issue volume.

---

### Feature Request Trends

Based on the issues and their labels, the most requested feature directions are:

- **Session limit continuation** (#13354) – by far the most upvoted request. Users want to resume work after hitting the daily cap without losing context.
- **Agent pausing** (#74689) – ability to pause a running agent mid-task (e.g., to inspect partial results) without aborting.
- **User interjections without killing subagents** (#74695) – when a user interrupts with additional context, background tasks should continue instead of being killed.
- **Copy markdown from chat responses in VS Code** (#54670) – a small but consistent ask for better output portability.
- **VSCode session history stability** (#69725) – sessions become unresumable after restart/update because names are regenerated and order changes.

These trends indicate a community that relies heavily on long-running sessions and multi-agent workflows, and is frustrated by abrupt termination or loss of progress.

---

### Developer Pain Points

The following pain points recur across multiple high-activity issues:

- **Memory leaks in long sessions** (#4953) – the #1 complaint by comment count. Users are forced to restart frequently to avoid OOM kills.
- **Safety filter false positives (Fable 5)** – multiple issues (#74743, #74771, #74770, #74615, #74754) report that legitimate code review, security hardening, and repository analysis are blocked. Several users report being locked out for days.
- **API / network instability with Advisor or model switching** (#69238, #74767) – retry loops and ECONNRESET errors that succeed when using plain `curl`.
- **Cross-platform inconsistencies** (#2805 – CRLF on Linux; #14828 – console flash on Windows; #73966 – install hang on KVM Linux) – each platform has its own set of rough edges.
- **Session / auth / cost errors** (#74751 – session limit reached with $0 cost; #74714 – org subscription disabled; #56988 – mobile remote control auth failures) – users are frequently blocked by quota or permission checks that seem incorrect.
- **Data integrity bugs** (#67071 – text not rendered; #74761 – `-p` exits early; #69725 – session history broken) – erode trust in the tool for deterministic, scriptable use.

The overall sentiment is that while Claude Code is powerful, regressions (especially around Fable 5 safety filters and memory management) are causing significant productivity loss for power users.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-07-06

## Today’s Highlights
The community remains highly engaged around two long‑standing pain points: Linux desktop support (Issue #11023) and a performance‑critical reasoning‑token clustering bug in GPT‑5.5 (Issue #30364). Meanwhile, several Windows‑specific fixes are landing in pull requests, including better line‑endings handling and URL case‑insensitivity. No new releases shipped in the past 24 hours.

## Releases
— No new releases in the last 24 hours.

## Hot Issues

1. **[#11023 – Codex desktop app for Linux](https://github.com/openai/codex/issues/11023)**  
   149 comments · 692👍  
   *Enhancement* – The most‑upvoted open issue. Users on Linux are frustrated by the lack of a native desktop app, especially those whose Mac laptops are unusable due to a separate bug. The community is actively requesting a Linux build.

2. **[#30364 – GPT-5.5 reasoning‑token clustering at 516/1034/1552](https://github.com/openai/codex/issues/30364)**  
   111 comments · 206👍  
   *Bug / model‑behavior* – A pattern where `gpt‑5.5`’s `reasoning_output_tokens` cluster at fixed numbers (516, 1034, 1552) correlates with degraded performance on complex tasks. Many users report hitting invisible ceilings.

3. **[#30009 – `apply_patch` fails with Windows sandbox error](https://github.com/openai/codex/issues/30009)**  
   23 comments · 4👍  
   *Bug / Windows* – File edits through the sandbox fail on Windows with an internal error. Affects Pro users on the latest desktop app (v26.616.81150).

4. **[#25246 – Business access‑tokens broken (401)](https://github.com/openai/codex/issues/25246)**  
   18 comments · 9👍  
   *Bug / auth* – Enterprise users cannot authenticate via access tokens. The issue has been open for over a month and is a blocker for business deployment.

5. **[#22085 – Windows: Codex spawns many Git for Windows processes](https://github.com/openai/codex/issues/22085)**  
   13 comments · 21👍 (closed)  
   *Bug / performance* – After a recent update, Codex launches dozens of Git processes, causing sustained high CPU. Closed but still referenced as a recurring Windows performance problem.

6. **[#30918 – Usage limits draining abnormally fast on Plus](https://github.com/openai/codex/issues/30918)**  
   11 comments · 7👍  
   *Bug / rate‑limits* – Users on the Plus plan see their 5‑hour limit drop from 70 % to 0 % in ~6 minutes during normal use. No warning or explanation provided.

7. **[#13165 – Specify shell Codex uses on Windows](https://github.com/openai/codex/issues/13165)**  
   7 comments · 28👍  
   *Enhancement* – Windows users want the ability to choose a custom shell (e.g., MinGW Bash) instead of being forced to use PowerShell, which causes compatibility issues.

8. **[#31163 – MCP `deny_unknown_fields` rejects top‑level `title`](https://github.com/openai/codex/issues/31163)**  
   5 comments · 0👍  
   *Bug / MCP* – FastMCP servers fail because Codex’s strict schema rejection of the `title` field blocks all tool elicitation calls.

9. **[#31125 – Sudden 80 % → 0 % usage drop on Pro plan](https://github.com/openai/codex/issues/31125)**  
   3 comments · 3👍  
   *Bug / rate‑limits* – Pro subscribers (€100/month) report an instantaneous drop from 80 % remaining to 0 % mid‑session, with support unresponsive.

10. **[#31237 – File preview fails for TypeScript files >739 lines](https://github.com/openai/codex/issues/31237)**  
    3 comments · 0👍  
    *Bug / app* – The desktop app’s preview pane fails to render small TypeScript files above about 739 lines, breaking the editing flow.

## Key PR Progress

1. **[#31223 – Preserve terminal input typed during startup](https://github.com/openai/codex/pull/31223)**  
   Fixes a long‑standing annoyance where text typed while Codex is initialising can be lost or partially captured.

2. **[#31192 – Flush queued terminal input before exit](https://github.com/openai/codex/pull/31192)**  
   Ensures that key‑release events queued after shutdown doesn’t garble the parent shell’s state.

3. **[#30882 – Preserve line endings when applying patches (Windows)](https://github.com/openai/codex/pull/30882)**  
   Critical fix for Windows users: patches now keep existing CRLF/LF terminators instead of forcing LF.

4. **[#30879 – Handle mixed‑case URLs in Windows command safety](https://github.com/openai/codex/pull/30879)**  
   PowerShell commands with uppercase `HTTP://` were incorrectly flagged as dangerous; now case‑insensitive matching is used.

5. **[#30492 – Fix slash command popup dismissal](https://github.com/openai/codex/pull/30492)**  
   Pressing Escape to dismiss the slash‑command popup no longer causes it to immediately re‑open on the next sync pass.

6. **[#27495 – Pass multi‑agent metadata to MCP tool calls](https://github.com/openai/codex/pull/27495)**  
   Adds `agent_path` and `has_spawned_subagent` to MCP request metadata, enabling subagent awareness for tools.

7. **[#29602 – Flatten namespace tools for providers without wrappers](https://github.com/openai/codex/pull/29602)**  
   Resolves #26234: third‑party Responses‑compatible endpoints that don’t support the `namespace` wrapper now work correctly.

8. **[#31155 – Release thread writer after failed shutdown](https://github.com/openai/codex/pull/31155)**  
   Prevents thread‑writer locks from being held forever after a persistence failure, allowing later sessions to resume.

9. **[#30488 – Show reset details in redemption picker](https://github.com/openai/codex/pull/30488)**  
   When redeeming usage limit resets, users now see which credits are available, their expiry dates, and which will be consumed.

10. **[#31201 – Reduce repeated plugin discovery work during tool assembly](https://github.com/openai/codex/pull/31201)**  
    Optimises tool suggestion by caching plugin metadata for 30 seconds and reusing unchanged remote catalog entries, cutting startup overhead.

## Feature Request Trends
- **Linux desktop app** (Issue #11023) remains the single most‑requested feature, with 692 👍 and strong community comments.
- **Custom shell selection on Windows** (Issue #13165) is a high‑priority request to avoid PowerShell‑specific breakage.
- **Better rate‑limit visibility** – users want a countdown timer in `/status` (Issue #31109) and a tool catalog in the JSON event stream (Issue #31088).
- **Tool permission UX improvements** – disabling letter‑key shortcuts during permission prompts (Issue #31037) to prevent accidental input.
- **Project workspace conflict detection** – a proactive warning when a prompt references a path outside the current thread’s workspace (Issue #31245).

## Developer Pain Points
- **Windows stability and sandbox issues** dominate the bug tracker: apply_patch failures (#30009), missing environment variables in WSL (#13556), Git process storms (#22085), and PowerShell‑specific problems (#27117, #29836).
- **Rate‑limit accounting irregularities** affect both Plus and Pro subscribers – unexplained fast drainage and instant drops (Issues #30918, #31125) erode trust in quota reporting.
- **Business/enterprise auth breakage** (#25246) has been unresolved for over a month, blocking organisational rollouts.
- **MCP compatibility** – strict schema validation (#31163) and missing metadata (#27495 PR) cause friction with third‑party servers.
- **Data loss fears** – renaming a project may delete chat history (#31240) and file previews fail on relatively small TypeScript files (#31237), disrupting daily workflows.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

Here is the Gemini CLI community digest for 2026-07-06.

---

## Gemini CLI Community Digest — 2026-07-06

### 1. Today's Highlights
Agent reliability remains the central theme, with several long-standing bugs around subagent reporting and indefinite hangs still drawing active discussion. A critical fix for a newline-in-string corruption bug was merged, alongside major dependency bumps that bring the platform up to date with the latest MCP and Google GenAI SDK releases. The nightly build v0.51.0 is rolling, but no stable release was cut today.

### 2. Releases
- **[v0.51.0-nightly.20260706.gf7af4e518](https://github.com/google-gemini/gemini-cli/releases/tag/v0.51.0-nightly.20260706.gf7af4e518)** — Automated nightly build. No release notes beyond version bump. Full diff available [here](https://github.com/google-gemini/gemini-cli/compare/v0.51.0-nightly.20260705.gf7af4e518...v0.51.0-nightly.20260706.gf7af4e518).

### 3. Hot Issues (Top 10 by Comment Count)
1. **[#8132 — Input token count exceeds max (1048576)](https://github.com/google-gemini/gemini-cli/issues/8132)** [CLOSED — P1]  
   *23 comments.* A long-running issue about large-scale token limit errors in prediction logs. Recently closed, signaling resolution or internal fix. Received 6 👍 reactions.

2. **[#22323 — Subagent recovery after MAX_TURNS reported as GOAL success](https://github.com/google-gemini/gemini-cli/issues/22323)** [OPEN — P1]  
   *10 comments.* Subagent reports success even when hitting max turn limits without doing any analysis, masking interruption. This continues to erode trust in agent autonomy.

3. **[#24353 — Robust component-level evaluations](https://github.com/google-gemini/gemini-cli/issues/24353)** [OPEN — P1]  
   *7 comments.* A major EPIC to expand behavioral tests (currently 76) across 6 supported Gemini models. Critical for quality assurance of the agent platform.

4. **[#21409 — Generalist agent hangs forever](https://github.com/google-gemini/gemini-cli/issues/21409)** [OPEN — P1]  
   *7 comments.* High frustration — the CLI hangs indefinitely when deferring to the generalist subagent, even for trivial tasks like folder creation. Workaround: disable subagents. 8 👍.

5. **[#22745 — Assess AST-aware file reads and mapping](https://github.com/google-gemini/gemini-cli/issues/22745)** [OPEN — P2]  
   *7 comments.* Investigates using AST to read method bounds precisely, reduce token waste, and improve codebase navigation. Core to the future of agent context management.

6. **[#21968 — Gemini does not use skills and sub-agents enough](https://github.com/google-gemini/gemini-cli/issues/21968)** [OPEN — P2]  
   *6 comments.* Even with well-described custom skills, the model rarely invokes them autonomously. Points to a deeper prompt/model behavior mismatch.

7. **[#26522 — Auto Memory retries low-signal sessions indefinitely](https://github.com/google-gemini/gemini-cli/issues/26522)** [OPEN — P2]  
   *5 comments.* Sessions that look uninteresting are never marked as processed, causing endless re-scanning. A performance and UX drain for memory-heavy users.

8. **[#25166 — Shell command execution gets stuck after completion](https://github.com/google-gemini/gemini-cli/issues/25166)** [OPEN — P1]  
   *4 comments.* Commands that finish successfully still show "Waiting input" indefinitely, blocking the agent. 3 👍 — a common frustration.

9. **[#21983 — Browser subagent fails on Wayland](https://github.com/google-gemini/gemini-cli/issues/21983)** [OPEN — P1]  
   *4 comments.* The browser subagent terminates with a GOAL status but fails silently on Wayland displays. Linux users hit a dead end.

10. **[#22672 — Agent should stop destructive behavior](https://github.com/google-gemini/gemini-cli/issues/22672)** [OPEN — P2]  
    *3 comments.* Model occasionally uses `git reset --force` or dangerous DB commands when safer alternatives exist. Highlights the need for a "safety mode" in the agent loop.

### 4. Key PR Progress (Top 10)
1. **[#28299 — Fix escape sequences in string literals](https://github.com/google-gemini/gemini-cli/pull/28299)** [OPEN]  
   *luisfelipe-alt.* Disables aggressive unescaping for modern models so `\n` inside strings is no longer corrupted to literal newlines. Directly addresses bug b-496211054.

2. **[#28089 — MCP elicitation (form + url) capability](https://github.com/google-gemini/gemini-cli/pull/28089)** [OPEN]  
   *KittiphonKamnuan.* Implements the 2025-11-25 MCP spec for client-side elicitation, enabling MCP servers to request forms or URLs from the user via the agent.

3. **[#28295 — Bump @google/genai from 1.30.0 to 2.10.0](https://github.com/google-gemini/gemini-cli/pull/28295)** [CLOSED]  
   *dependabot.* Major jump for the core GenAI JS SDK. Likely includes extended tool use and reasoning improvements.

4. **[#28294 — Bump @agentclientprotocol/sdk to 1.0.0](https://github.com/google-gemini/gemini-cli/pull/28294)** [CLOSED]  
   *dependabot.* MCP SDK hit 1.0.0. This is a foundational dependency for all agent-client protocol interactions.

5. **[#28164 — Limit recursive reasoning turns per user request](https://github.com/google-gemini/gemini-cli/pull/28164)** [OPEN]  
   *amelidev.* Introduces a 15-turn default cap on recursive reasoning to prevent infinite loops. Customizable via `maxSessionTurns`.

6. **[#28068 — Guard message inspectors against empty parts arrays](https://github.com/google-gemini/gemini-cli/pull/28068)** [CLOSED]  
   *AriaZhao-coder.* Fixes a JavaScript vacuum-truth bug where empty `parts` arrays were misclassified as function calls/responses. A quiet but impactful correctness fix.

7. **[#28292 — Bump puppeteer-core to 25.2.1](https://github.com/google-gemini/gemini-cli/pull/28292)** [CLOSED]  
   *dependabot.* Updates the browser automation driver. Should improve stability of the browser subagent.

8. **[#28290 — Bump chrome-devtools-mcp to 1.4.0](https://github.com/google-gemini/gemini-cli/pull/28290)** [CLOSED]  
   *dependabot.* Chrome DevTools MCP server hit 1.4.0, likely bringing new inspection capabilities.

9. **[#28288 — Bulk npm dependency update (74 packages)](https://github.com/google-gemini/gemini-cli/pull/28288)** [CLOSED]  
   *dependabot.* A large housekeeping PR covering `simple-git`, `@octokit/rest`, and 72 other packages. Reduces security surface.

10. **[#28289 — Bump js-yaml from 4.1.1 to 5.2.0](https://github.com/google-gemini/gemini-cli/pull/28289)** [CLOSED]  
    *dependabot.* Major version update for the YAML library. May affect config file parsing.

### 5. Feature Request Trends
- **Subagent trajectory sharing**: Strong demand to make subagent decisions visible and shareable (e.g., via `/chat share`). Users want transparency into how subagents arrived at conclusions.
- **AST-aware tooling**: Several issues advocate using ASTs for file reads, codebase mapping, and method-bound extraction to reduce token waste and improve context precision.
- **Component-level evaluation infrastructure**: The community and maintainers are pushing for a systematic, model-versioned evaluation pipeline to benchmark agent behavior across all Gemini models.
- **Browser agent resilience**: Calls for automatic session takeover, lock recovery, and configuration override support for the browser subagent, especially in persistent session mode.
- **Agent self-awareness**: Users want the CLI to know its own flags, hotkeys, and internal mechanics so it can serve as its own documentation and troubleshooting guide.

### 6. Developer Pain Points
- **Agent hangs and indefinite waiting**: The generalist agent hangs (#21409) and shell commands that stay in "Waiting input" (#25166) are the most pervasive blockers.
- **Subagent permission and control issues**: Subagents running without consent (#22093) and ignoring settings.json overrides (#22267) cause frustration, especially for users who want to disable agentic features.
- **Destructive / careless behavior**: Model using `git reset --force`, `rm -rf`, or unsafe DB commands (#22672) without user confirmation is a recurring safety concern.
- **Token limit errors and wasted context**: Both the closed #8132 (token limit errors) and open #22745 (AST vs. raw file reads) point to the same pain: the model consumes too many tokens on irrelevant context.
- **Memory system inefficiency**: Auto Memory retrying low-signal sessions (#26522) and silent invalidation of memory patches (#26523) weaken trust in the long-term memory feature.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-07-06

## Today’s Highlights
A new patch release (v1.0.69-2) landed with improvements to MCP OAuth sign‑in and a fix for file inclusion in `n`. The community is actively discussing plugin scoping—the long‑running #1665 finally closed after 18 👍, and a fresh issue (#4003) asks for custom model endpoints, mirroring VS Code capabilities. Several authentication and platform‑specific bugs remain hot topics.

## Releases
**v1.0.69-2** (released within the last 24h)
- **Added:** `/rubber-duck` command is now shown in pre‑auth help and self‑documentation.
- **Improved:** Sign‑in to MCP servers through the CLI OAuth callback flow; full `/user` switch picker now revealed when timeline is full (hint bar no longer clipped).
- **Fixed:** Files inside `n` are now correctly included.

[View release](https://github.com/github/copilot-cli/releases/tag/v1.0.69-2)

## Hot Issues (10 noteworthy)

1. **[#1665]** [CLOSED] [area:plugins, area:configuration] Support Copilot CLI Plugins Scoped to Project or Repository (instead of per‑user) — *18 👍, 10 comments*  
   A long‑standing request that finally closed today. Users want plugins to be tied to a repo or project, not installed globally.  
   [Issue link](https://github.com/github/copilot-cli/issues/1665)

2. **[#3596]** [CLOSED] [area:authentication, area:sessions, area:models] Error loading model list: Error: Not authenticated — *11 👍, 9 comments*  
   Resuming a specific session causes the `/model` command to fail with “Not authenticated.” Affects v1.0.56.  
   [Issue link](https://github.com/github/copilot-cli/issues/3596)

3. **[#3028]** [OPEN] [area:permissions, area:mcp] MCP permissions — *5 👍, 8 comments*  
   Request for granular allow‑listing of tools from MCP servers, with a `trustedFolders`‑like configuration.  
   [Issue link](https://github.com/github/copilot-cli/issues/3028)

4. **[#2367]** [CLOSED] [area:agents] Copilot does not wait for specialized agents — *0 👍, 3 comments*  
   Sub‑agents spawned for specialized work are abandoned if execution takes too long, bypassing their instructions.  
   [Issue link](https://github.com/github/copilot-cli/issues/2367)

5. **[#1428]** [CLOSED] [area:tools] GitHub Copilot CLI Bash Tool Incompatible with Nix Shell Environment — *7 👍, 3 comments*  
   Bash tool commands hang indefinitely inside a Nix develop shell. Sessions die immediately.  
   [Issue link](https://github.com/github/copilot-cli/issues/1428)

6. **[#4003]** [OPEN] [area:models] Support custom model endpoint in Copilot CLI (like VS Code) — *0 👍, 3 comments*  
   Users want to connect local or private models, similar to VS Code’s Language Models panel.  
   [Issue link](https://github.com/github/copilot-cli/issues/4003)

7. **[#3074]** [OPEN] [area:models] Add an `/effort` command to quickly switch reasoning effort for the current model — *6 👍, 2 comments*  
   Currently multi‑step; a single command to toggle Low/Medium/High reasoning effort would improve efficiency.  
   [Issue link](https://github.com/github/copilot-cli/issues/3074)

8. **[#3945]** [OPEN] [area:context-memory] Memories are leaking between repositories — *0 👍, 2 comments*  
   Copilot references “facts stored in the memory” from unrelated repos, causing confusion.  
   [Issue link](https://github.com/github/copilot-cli/issues/3945)

9. **[#4034]** [CLOSED] Hook subprocess stdin write‑end left open (no EOF) for tool‑use hooks — documented `$(cat)` pattern hangs — *0 👍, 2 comments*  
   `preToolUse`/`postToolUse` hooks fail because stdin is not closed; the documented workaround breaks.  
   [Issue link](https://github.com/github/copilot-cli/issues/4034)

10. **[#2930]** [OPEN] [area:context-memory] Feature Request: Local auto‑memory (agent‑initiated, no remote storage) — *2 👍, 1 comment*  
    Organizations with security concerns disable remote memory; local-only memory would fill the gap.  
    [Issue link](https://github.com/github/copilot-cli/issues/2930)

## Key PR Progress
No pull requests were updated in the last 24 hours. The community’s focus remains on issue triage and the v1.0.69-2 release.

## Feature Request Trends
- **Plugin placement & scoping** – Users strongly want per‑project/repo plugin configuration (#1665, #3028).
- **Model flexibility** – Support for custom endpoints (#4003) and quick reasoning‑effort switching (#3074) are top requests.
- **Memory improvements** – Both local auto‑memory (#2930) and fixing cross‑repo memory leaks (#3945) are recurring themes.
- **Platform parity** – Windows hook compatibility (#4001) and Nix shell support (implied by #1428’s close) need attention.
- **Authentication flow** – Session‑specific auth failures (#3596, #3902) highlight a need for more robust re‑authentication.

## Developer Pain Points
- **Authentication flakiness** – Resuming sessions (#3596) and ACP mode (#3902) can leave users locked out.
- **Platform‑specific breakdowns** – Nix shell hang (#1428), Windows uninstall failure (#3662), and Windows hook incompatibility (#4001) frustrate developers on alternative environments.
- **Hook & subprocess issues** – The stdin EOF problem (#4034) breaks documented hook patterns; MCP permission controls (#3028) are too coarse.
- **Memory confusion** – Leaking context between repos (#3945) undermines trust in the agent’s awareness.
- **UX clarity** – The “No, and tell copilot what to do” option (#4033) is ambiguous compared to previous behavior.

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

Here is the **Kimi Code CLI Community Digest** for **2026-07-06**.

---

## Kimi Code CLI Community Digest | 2026-07-06

### 1. Today's Highlights
Today saw a quiet day on the release and PR front, with no new versions or merges. The community spotlight, however, is on a critical display bug affecting Windows users (Issue #2485) and a well-articulated feature request for API exposure of usage limits (Issue #2486), signaling a growing need for IDE-native tooling integration. Developers are particularly concerned about terminal stability, while integrators are pushing for programmatic access to usage data.

### 2. Releases
**None** – No new releases in the last 24 hours.

### 3. Hot Issues
*(Note: Only 2 issues met the update criteria today. We highlight both below.)*

- **#2485 [BUG] code cli 错乱 (CLI is confused)** – **Critical Stability Issue**
  - **Summary:** A user running Kimi-CLI v0.22.0 on Windows 11 reports severe terminal corruption after prolonged use. The UI loses the first option, and content is displayed incorrectly.
  - **Why it matters:** This is a high-impact bug affecting core usability (terminal rendering) on a primary platform (Windows). It directly impacts developer workflow and trust.
  - **Community reaction:** Only 1 comment currently; the community has not yet upvoted this issue, but the severity suggests it will gain traction quickly.
  - 🔗 [Issue #2485](https://github.com/MoonshotAI/kimi-cli/issues/2485)

- **#2486 [ENHANCEMENT] Feature Request: Expose Kimi Code usage limits and reset times through ACP** – **API Integration Gap**
  - **Summary:** Developer `jgiacomini` requests that usage limit data (currently only visible via `/usage` console command) be exposed through the ACP (AI Communication Protocol). They're building an ACP client for Visual Studio 2026 and need this data to provide in-IDE usage visibility.
  - **Why it matters:** This reflects a key ecosystem need: as developers build custom IDEs and tooling around Kimi Code, they require programmatic access to billing/rate-limit data.
  - **Community reaction:** No comments yet, but the request is clear and well-scoped.
  - 🔗 [Issue #2486](https://github.com/MoonshotAI/kimi-cli/issues/2486)

### 4. Key PR Progress
**None** – No pull requests were updated in the last 24 hours.

### 5. Feature Request Trends
Based on the single new feature request today (#2486), the primary direction is:

- **ACP API Transparency & Billing Integration**: Developers want Kimi Code to expose usage quotas, current consumption, and reset timestamps through its ACP endpoints. This would enable custom IDEs (e.g., VS 2026, JetBrains) to display real-time usage status without users needing to switch to the CLI console.

### 6. Developer Pain Points
Based on today’s data (Issue #2485), the primary recurring frustration is:

- **Terminal Display Stability on Windows**: The `kimi-cli` interface (likely a TUI) becomes corrupted after extended use, with options disappearing. This suggests a rendering or state management issue specific to the Windows Terminal environment. Developers using Windows as their primary platform are directly blocked.
- **Limited IDE Integration Visibility**: While not a bug, the lack of a public API for usage data (Issue #2486) forces developers to either build workarounds or manually check the console. This is a friction point for teams wanting to embed Kimi Code deeply into their workflow.

---

*Digest generated from GitHub data (MoonshotAI/kimi-cli) for 2026-07-06.*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

Here is the OpenCode community digest for **2026-07-06**.

---

## OpenCode Community Digest — 2026-07-06

### Today’s Highlights
A quiet release day, but the community remains highly active with long-running feature debates on tokens-per-second display (#5374) and session lifecycle hooks (#16626, #28695). A controversial billing UX issue (#34754) was closed today, while a fresh wave of Windows-specific bugs (#35536, #35545) and plugin extensibility PRs (#35166, #35433) are shaping the near-term roadmap.

### Releases
No new releases in the last 24 hours.

### Hot Issues

1.  **[FEATURE]: show tokens/second (#5374)** — 84 👍, 19 comments  
    The most upvoted open feature this month. Users want real-time tokens/s metrics to compare provider performance. High community demand suggests this is a priority for the next minor release.  
    [GitHub](https://github.com/anomalyco/opencode/issues/5374)

2.  **opencode using CPU for nothing (#19466)** — 10 👍, 12 comments  
    A persistent performance bug where OpenCode burns ~50% of a CPU core while idling during API rate-limit retries. Concerns about battery life and noise on powerful machines.  
    [GitHub](https://github.com/anomalyco/opencode/issues/19466)

3.  **The opencode funneling scam (#34754)** — **CLOSED** today, 5 comments  
    Heated reports of users being charged for a Zen subscription instead of Go due to confusing UX. This was a major trust and billing issue; the closure suggests a fix or compensation has been communicated.  
    [GitHub](https://github.com/anomalyco/opencode/issues/34754)

4.  **Auto-discover skills from nested subdirectories (#31377)** — 5 comments  
    Currently, skill discovery only walks *up* from the CWD. This request asks for a downward walk to find skills in subdirectories, enabling modular project structures.  
    [GitHub](https://github.com/anomalyco/opencode/issues/31377)

5.  **Default sharing to "disabled" — privacy by default (#17188)** — 13 👍, 4 comments  
    Privacy advocates continue to push for opt-in data sharing. The current default raises consent concerns for new users.  
    [GitHub](https://github.com/anomalyco/opencode/issues/17188)

6.  **Session lifecycle context hooks for persistent plugin state (#28695)** — 4 comments  
    A natural extension of the popular `session.stopping` hook (#16626). Developers want reliable hooks to flush data and persist state when sessions or subagents end.  
    [GitHub](https://github.com/anomalyco/opencode/issues/28695)

7.  **{env:} substitution breaks on Windows paths (#35536)** — **NEW** today, 3 comments  
    Environment variable expansion produces JSON-invalid backslashes on Windows, causing parser errors. A blocker for Windows-based workflows.  
    [GitHub](https://github.com/anomalyco/opencode/issues/35536)

8.  **instructions field in global opencode.jsonc is ignored (#35552)** — **NEW** today  
    Global config instructions are silently overridden by `.claude/CLAUDE.md`. This breaks expected config layering behavior.  
    [GitHub](https://github.com/anomalyco/opencode/issues/35552)

9.  **DigitalOcean OAuth fails (#27764)** — 1 comment  
    A platform integration bug—the Model Access Key creation endpoint has been retired by DigitalOcean, breaking the OAuth flow completely.  
    [GitHub](https://github.com/anomalyco/opencode/issues/27764)

10. **Desktop renderer crashes when session lists are object maps (#35551)** — **NEW** today  
    A crash in the desktop app when API responses return objects instead of arrays for session/command lists. Immediate usability issue for renderer users.  
    [GitHub](https://github.com/anomalyco/opencode/issues/35551)

### Key PR Progress

1.  **feat(tui): implement prompt submit with session commands (#14442)** — Open, 0 comments  
    A long-running mega-PR that closes 4 issues. It unifies CLI prompt passing across `--session`, `--continue`, and `--fork`. High impact for scripting and automation.  
    [GitHub](https://github.com/anomalyco/opencode/pull/14442)

2.  **feat(plugin): support plugin-provided VCS backends (#35166)** — Open, new this week  
    A major plugin seam addition. Allows plugins to provide custom `vcs.status` and `vcs.diff` implementations without forking core—critical for enterprise VCS integrations.  
    [GitHub](https://github.com/anomalyco/opencode/pull/35166)

3.  **fix(queue): cleanup queue UX (#35369)** — Open, new work  
    Adds a follow-up queue for desktop: users can `Ctrl+Enter` to stage prompts while the model is busy—reduces friction in rapid-fire workflows.  
    [GitHub](https://github.com/anomalyco/opencode/pull/35369)

4.  **fix(core): add diff size limits to prevent UI freeze (#35546)** — **NEW** today  
    Directly addresses the infamous "review panel hang" with large changesets (related to #31916, #35401). Critical for users editing large files.  
    [GitHub](https://github.com/anomalyco/opencode/pull/35546)

5.  **fix(tui): add ctrl+h as backspace alias for herdr/ConPTY (#35545)** — **NEW** today  
    Fixes a Windows terminal compatibility issue where physical backspace sends byte `0x08`. Small change, big impact for Windows ConPTY users.  
    [GitHub](https://github.com/anomalyco/opencode/pull/35545)

6.  **fix(shell): drain stdout before reading output (#35543)** — **NEW** today  
    Fixes a race condition where shell commands return `(no output)` with exit code `0` on first call. Important for reliability of shell tool usage.  
    [GitHub](https://github.com/anomalyco/opencode/pull/35543)

7.  **fix(core): resolve spawn completion on exit, not only close (#29831)** — Open  
    Addresses a Windows-specific hang when commands spawn detached child processes. The agent was waiting forever for child stdout to close.  
    [GitHub](https://github.com/anomalyco/opencode/pull/29831)

8.  **fix(opencode): stop sending tools when `tool_call` is false (#35433)** — Open, new  
    Bug fix: the `tool_call: false` config flag was stored but never checked in prompt construction, so tools were always sent to the model.  
    [GitHub](https://github.com/anomalyco/opencode/pull/35433)

9.  **fix(provider): respect model limit.output instead of capping at 32k (#34901)** — Open  
    A long-awaited fix (closing 6 related issues). Model output token limits were hardcapped at 32k regardless of model capabilities.  
    [GitHub](https://github.com/anomalyco/opencode/pull/34901)

10. **fix(tui): restore ESC interrupt for delegated subagent sessions (#32767)** — Open  
    Restores a regression where pressing Esc during a subagent session failed to interrupt the agent. Closes three long-standing issues (#3699, #4073, #23534).  
    [GitHub](https://github.com/anomalyco/opencode/pull/32767)

### Feature Request Trends

- **Session & Plugin Lifecycle Hooks** — The dominant theme this period. Requests for `session.stopping` (#16626), session finalization hooks (#35540), and persistent plugin state hooks (#28695) signal a strong desire for more structured plugin control.
- **Performance Monitoring** — High demand for tokens-per-second metrics (#5374) and modeling CPU usage during idle waits (#19466). Developers want to benchmark and debug provider latency.
- **Privacy by Default** — The "privacy first" movement (#17188, referencing #7982, #459) continues to gain traction as more users join the ecosystem.
- **Agent Configuration via ACP** — A renewed call (#35550) to support Agent Communication Protocol for configuring subagent behavior, inspired by Zed’s approach.
- **Local & Hybrid Model Support** — Integration improvements with Ollama (#19948) and Mistral tool calling (#16636) remain active, driven by users preferring local inference.

### Developer Pain Points

- **Windows-Specific Issues** — A persistent source of friction: path escaping bugs (#35536), ConPTY terminal quirks (#35545), and detached-child process hangs (#29831).
- **Sandbox Integration** — The billing UX controversy (#34754) highlights frustration with confusing subscription flows and lack of compensation communication.
- **Documentation & Config Layering** — The global config ignore bug (#35552) and unclear documentation around `opencode.jsonc` layering are causing confusion.
- **Crash-on-Startup** — At least two reports (#35549, #35551) of crashes on launch due to corrupted message logs or malformed API responses, indicating fragility in data loading paths.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

## Pi Community Digest — 2026-07-06

### Today's Highlights
Two critical fault lines dominated today’s activity: **null `content` crashes** (#6259, #6276) that affect reasoning models like GLM-5.2, and the **long‑running edit‑tool breakage** with Claude and GPT (#6278) that is now driving the push for strict tool schemas (#6306, #6341). On the positive side, the team merged fixes for cache‑token double‑counting (#6352) and landed new providers (Requesty, StepFun), while the deprecation of Qwen’s free tier was formally cleaned up (#3832).

### Releases
No new releases in the last 24h.

---

### Hot Issues (10 Noteworthy)

1. **[#6278 — New Claude models fail ~20% of edits due to extra keys](https://github.com/earendil-works/pi/issues/6278)**  
   High‑comment (21), long‑running. Claude invents extra fields like `new_text_x`, `type`, `closeenough` that the edit tool rejects. Causes severe workflow disruption. Community upvotes (👍5) show broad impact.

2. **[#6306 — Support Strict Tools / Grammar](https://github.com/earendil-works/pi/issues/6306)**  
   Directly linked to #6278. Proposes a way to constrain tool arguments with JSON‑schema‑based “strict” tools (as OpenAI does). Currently blocked by SDK limitations. Mitsuhiko opened it, signalling high‑priority.

3. **[#6187 — Pi login hangs in WSL after Copilot device auth](https://github.com/earendil-works/pi/issues/6187)**  
   WSL users cannot complete login — the client never detects the browser auth. 18 comments, no fix yet. A blocker for Windows/Linux hybrid setups.

4. **[#6259 — `content is not iterable` when reasoning models return null content](https://github.com/earendil-works/pi/issues/6259)**  
   GLM-5.2 on Fireworks returns `content=null` when both `reasoning_content` and `tool_calls` exist. Crashes in multiple code paths. Already spawned related issues (#6276, #4909).

5. **[#4338 — Agent says “working” but makes no progress](https://github.com/earendil-works/pi/issues/4338)**  
   Classic stuck‑loop behaviour. Closed months ago but re‑surfaces in user reports. Community frustration (👍2) – the agent can appear busy without producing output.

6. **[#6204 — `mimo-v2-omni` is a ghost model on all MiMo providers](https://github.com/earendil-works/pi/issues/6204)**  
   Model listed in catalog but endpoints return 400 “Not supported”. Misleading for users on Xiaomi Token Plan.

7. **[#6300 — Windows: input line redrawn on every keystroke](https://github.com/earendil-works/pi/issues/6300)**  
   Characters appear on new lines in cmd.exe and Windows Terminal – severe TUI regression for Windows users.

8. **[#6362 — Paste counter not reverted when content is removed](https://github.com/earendil-works/pi/issues/6362)**  
   `[Paste #n]` markers remain even after deletion, causing incorrect indexing. Small but annoying UX bug.

9. **[#6353 — Cache hit rate denominator double‑counts cache tokens](https://github.com/earendil-works/pi/issues/6353)**  
   Anthropic API `input_tokens` already includes cache sub‑fields. Pi was adding them on top, inflating CH% and context%. Duplicate report #6355 also filed.

10. **[#6359 — TUI segfault on small‑ICU Node builds (Intl.Segmenter null deref)](https://github.com/earendil-works/pi/issues/6359)**  
    RHEL’s minimal `nodejs` package lacks full ICU, causing `Intl.Segmenter` to be unavailable → segfault. `--mode json` works; TUI instantly crashes.

---

### Key PR Progress (10 Important)

1. **[#1050 — TUI: jump to line start/end when pressing up/down at boundaries](https://github.com/earendil-works/pi/pull/1050)**  
   Editor navigation improvement – cursor wraps to line start/end instead of doing nothing. Closed after a long review cycle.

2. **[#3832 — Remove Qwen CLI OAuth provider extension](https://github.com/earendil-works/pi/pull/3832)**  
   Qwen’s free tier was discontinued 15 Apr 2026. Extension no longer functional. Clean removal of code and references.

3. **[#5789 — Fix cursorUp line‑start jump before history browsing](https://github.com/earendil-works/pi/pull/5789)**  
   Bugfix: pressing Up on first line should jump to start of line, not enter history. Resolves regression from #1050.

4. **[#6356 — Fix GLM‑5.2 tool calls (OpenCode Go)](https://github.com/earendil-works/pi/pull/6356)**  
   Switches to non‑streaming completion when tools are present because GLM‑5.2’s streaming misses tool‑call deltas.

5. **[#6352 — Correct cache hit rate denominator and context token double‑count](https://github.com/earendil-works/pi/pull/6352)**  
   Merged fix for the double‑counting bug reported in #6353/#6355. Critical for users who rely on cache statistics in the footer.

6. **[#6350 — Add `before_provider_headers` extension hook (open)](https://github.com/earendil-works/pi/pull/6350)**  
   New lifecycle hook lets extensions modify outgoing HTTP headers – essential for LLM gateways and proxy integrations.

7. **[#5472 — Add Requesty as native provider](https://github.com/earendil-works/pi/pull/5472)**  
   Requesty (60k+ users) gets first‑class support: `requesty/...` models work without manual OpenAI‑compatible config.

8. **[#6348 — Show cumulative cache stats in footer](https://github.com/earendil-works/pi/pull/6348)**  
   Enhances the TUI footer to display rolling cache hit/miss counts, building on the fix from #6352.

9. **[#6343 — Normalize null message content at ingestion boundaries (open)](https://github.com/earendil-works/pi/pull/6343)**  
   Systematic fix for the recurring `null` content crashes (#6259, #6276). Forces `content` to always be an array or string early.

10. **[#6341 — Support constrained sampling (open)](https://github.com/earendil-works/pi/pull/6341)**  
   Implements strict tools / constrained sampling for providers that support it (OpenAI, etc.). Direct answer to #6306 and #6278.

---

### Feature Request Trends

- **Strict Tool Schemas & Constrained Sampling** – Multiple issues (#6306, #6278, #6015) and PR (#6341) point to a strong demand for LLM‑side argument validation to prevent invented fields. This is the #1 requested feature.
- **Provider Diversity** – New providers (Requesty, StepFun, Agnes AI) and fixes for GLM‑5.2 show the community actively adding support for alternatives to OpenAI/Anthropic. Qwen removal also highlights the fragility of free tiers.
- **Customizability & Hooks** – Requests for custom share commands (#6358), alternative TUI engines (#6346), and header‑modification hooks (#6350) indicate users want to tailor Pi to their workflows and hosting environments.
- **Performance & Startup** – Jiti instance sharing (#6360/#6361) and cache stat improvements reflect a community focused on reducing cold‑start latency and accurate metrics for cost optimization.
- **Session Isolation** – Embedded sub‑agents (e.g., piMuster) need isolation from the operator’s context (#6203). A per‑session `isolated` flag is proposed.

---

### Developer Pain Points

1. **Edit Tool Brittleness** – Both Claude and GPT frequently inject extra fields into tool calls, causing validation failures. The lack of strict schemas forces users into workarounds or session restarts.
2. **Null Content Crashes** – Reasoning models (GLM‑5.2, possibly others) return `content: null` when tool calls are present, crashing compaction and rendering. A systemic fix is in PR #6343.
3. **Cache Metric Confusion** – The Anthropic cache‑token double‑count bug (#6353, #6355) misled users about actual costs and efficiency. Already fixed but symptomatic of subtle API convention mismatches.
4. **Platform‑Specific TUI Bugs** – Windows redraw (#6300), RHEL segfault (#6359), and WSL login hanging (#6187) show the TUI and auth flows are fragile outside macOS/Linux with full Node.
5. **Agent Loops** – The “stuck agent” (#4338) and untracked error cases (#4433) erode trust in autonomous mode. Auto‑retry coverage is incomplete.
6. **Session Storage Bugs** – UUID collisions in UUIDv7 generation (#6242) and missing `labelTimestampsById` cleanup (#6354) threaten data integrity for long‑running sessions.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest – 2026-07-06

## Today’s Highlights
A new nightly release strengthens PR triage with batch detection and red-flag patterns, while the community is rallying around a major RFC for multi-workspace daemon support. Several high-impact bugs are under active investigation, including zombie sessions bleeding tokens and a context-window calculation flaw that can block all requests.

---

## Releases
**v0.19.6-nightly.20260706.47f62a466** – one change:
- `fix(triage): strengthen PR gate with batch detection, problem existence check, and red flag patterns` ([view release](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.6-nightly.20260706.47f62a466))

---

## Hot Issues (10 noteworthy)

1. **#6378 – RFC: Support multiple workspaces in one `qwen serve` daemon**  
   Proposes 1 daemon = N workspaces while keeping backward compatibility. 18 comments; strong community interest.  
   [Issue #6378](https://github.com/QwenLM/qwen-code/issues/6378)

2. **#3804 – [Bug] `AskUserQuestion` frequently yields `API Error: Model stream ended with empty response text`**  
   Persistent issue since v0.15.6, still without a fix after two months.  
   [Issue #3804](https://github.com/QwenLM/qwen-code/issues/3804)

3. **#6264 – `/review` skill consumes excessive tokens**  
   Users report large token usage (with screenshots) when running code review, raising concerns about cost.  
   [Issue #6264](https://github.com/QwenLM/qwen-code/issues/6264)

4. **#5964 – Zombie sessions silently burn tokens**  
   A running Agent continued for 8 hours without any usage logging, costing user credits. Community reaction: “like finding an electronic cockroach stealing electricity.”  
   [Issue #5964](https://github.com/QwenLM/qwen-code/issues/5964)

5. **#6298 – Shell tool fails on Windows when command produces stdout**  
   `run_shell_command` pipes through `cat`, unavailable in Windows `cmd.exe`. Blocks Windows users entirely.  
   [Issue #6298](https://github.com/QwenLM/qwen-code/issues/6298)

6. **#6384 – Hard limit: 0 when env-configured model reserves full default context window for output**  
   Context too large error before any request, with `hard limit: 0`. Can freeze the tool completely.  
   [Issue #6384](https://github.com/QwenLM/qwen-code/issues/6384)

7. **#6396 – `/review` can downgrade its own approval while waiting for its own CI check**  
   The automated review flow counts its own pending check-run as blocked CI, accidentally turning “LGTM” into a comment.  
   [Issue #6396](https://github.com/QwenLM/qwen-code/issues/6396)

8. **#6392 – Feature request: add `dmPolicy` config to disable private/DM messages in channels**  
   Analogous to existing `groupPolicy`, requested for channel operators.  
   [Issue #6392](https://github.com/QwenLM/qwen-code/issues/6392)

9. **#6383 – RFC: Agent View for managing background sessions**  
   Inspired by Claude Code’s Agent View – a TUI dashboard for multiple background sessions. Narrow v1 scope proposed.  
   [Issue #6383](https://github.com/QwenLM/qwen-code/issues/6383)

10. **#6119 – `list_directory` and `read_file` have inconsistent git-ignore handling** (closed)  
    Already fixed, but the inconsistency was a common source of confusion for users relying on `.gitignore`.  
    [Issue #6119](https://github.com/QwenLM/qwen-code/issues/6119)

---

## Key PR Progress (10 important)

1. **#6033 – `fix(core): Parse tagged thinking for GLM responses`** (CLOSED)  
   Enables `<think>…</think>` extraction for DashScope GLM models. Merged – improves multi-model compatibility.  
   [PR #6033](https://github.com/QwenLM/qwen-code/pull/6033)

2. **#6397 – `fix(cli): ignore current review run in presubmit CI`** (OPEN)  
   Directly addresses #6396: the review workflow now skips its own check-run when deciding approval status.  
   [PR #6397](https://github.com/QwenLM/qwen-code/pull/6397)

3. **#6395 – `feat(review): add issue-fidelity and root-cause ownership gate to /review`** (OPEN)  
   Adds a dedicated Agent 0 stage for verifying bugfix PRs against linked issues, reducing trust in PR author framing.  
   [PR #6395](https://github.com/QwenLM/qwen-code/pull/6395)

4. **#4866 – `refactor(ci): split PR triage into 4-job pipeline`** (OPEN)  
   Replaces monolithic `/triage` with a layered pipeline (resolve → product-decision → …). Improves maintainability.  
   [PR #4866](https://github.com/QwenLM/qwen-code/pull/4866)

5. **#6394 – `feat(cli): Add Phase 1 workspace runtime registry`** (OPEN)  
   First step toward multi-workspace daemon – internal registry grouping bridge, workspace service, etc.  
   [PR #6394](https://github.com/QwenLM/qwen-code/pull/6394)

6. **#6354 – `feat(core): add maxSubAgents setting to limit parallel sub-agent count`** (OPEN)  
   Caps concurrent sub-agents; extras are queued without timeout pressure. Useful for resource management.  
   [PR #6354](https://github.com/QwenLM/qwen-code/pull/6354)

7. **#6369 – `fix(triage): exclude test files from core module size gate and distinguish feat from refactor`** (OPEN)  
   Refines the triage bot’s “Core Module Protection” gate: tests don’t trigger the line limit, and refactors are handled separately.  
   [PR #6369](https://github.com/QwenLM/qwen-code/pull/6369)

8. **#6393 – `feat(cli): review auto-generated skills with inline preview, editor handoff, and off switch`** (OPEN)  
   Improves the skill review dialog: shows full skill content, allows editing, and adds an in-dialog disable option.  
   [PR #6393](https://github.com/QwenLM/qwen-code/pull/6393)

9. **#6359 – `fix(cli): Keep model picker entries contiguous in short terminals`** (OPEN)  
   Fixes rendering breakage when terminal height is limited. Useful for remote SSH / small windows.  
   [PR #6359](https://github.com/QwenLM/qwen-code/pull/6359)

10. **#6346 – `feat(daemon): add session artifact content retention`** (OPEN)  
    Pinned artifact content can survive restarts, be read via daemon APIs, and referenced by hash.  
    [PR #6346](https://github.com/QwenLM/qwen-code/pull/6346)

---

## Feature Request Trends

- **Multi-workspace daemon (#6378, #6394, #6383)** – The dominant feature direction: allow a single `qwen serve` process to manage multiple workspaces and background sessions, with a dedicated TUI dashboard for oversight.
- **Review pipeline intelligence (#6395, #6392, #6397)** – Community wants smarter /review: issue fidelity verification, DM/group policy controls, and self-consistency checks to avoid approval downgrades.
- **Resource control (#6354, #6264, #6384)** – Users consistently ask for limits on token consumption, parallel sub-agents, and better context-window defaults to prevent surprises.
- **Session management & visibility (#6383, #5964, #6388)** – There is strong demand for session lifecycle visibility (zombie detection, usage analytics, artifact persistence).

---

## Developer Pain Points

- **Token bleeding from zombie sessions (#5964)** – Long-running agents without proper logging or auto-timeout remain a top frustration, costing real money.
- **Windows compatibility (#6298, #6390)** – The shell tool and pager defaults are Linux-centric, leaving Windows users unable to execute basic commands.
- **Inconsistent context-window behavior (#3804, #6384, #6387)** – Empty responses and zero hard-limit errors appear when using non-standard model setups; users expect sane defaults (200K now proposed in #6387).
- **/review token cost (#6264)** – Code review consumes massive tokens with no visible cap; users ask for cost controls or a warning before triggering.
- **Gitignore handling inconsistencies (#6119)** – Though fixed, the asymmetry between `list_directory` and `read_file` caused many subtle bugs during development.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-07-06

**Repo:** [github.com/Hmbown/DeepSeek-TUI](https://github.com/Hmbown/DeepSeek-TUI)  
*All linked items point to the corresponding `Hmbown/CodeWhale` issue or PR, as the upstream project uses CodeWhale as the primary tool name.*

---

## Today’s Highlights

The community is deep in **v0.8.68 preparation**, with a major EPIC (#2870) tracking a staged command-boundary refactor and a product-readiness umbrella (#4038) for the **Workflow** (née WhaleFlow) feature. Performance fixes land today for high-fan-out agent sessions, UTF-8 cursor panics, and broken-pipe crashes. A constitution-compliance bug (#4032) about CodeWhale not following user-provided scripts draws significant discussion (19 comments).

## Releases

No new releases in the last 24 hours.

## Hot Issues (10 noteworthy)

1. **[#4032] Codewhale not following the constitution**  
   *User `stream2stream` reports that CodeWhale ignores existing scripts and writes temporary ones, then justifies the behavior. 19 comments indicate high community interest in trust and tool-compliance.*  
   [Issue #4032](https://github.com/Hmbown/CodeWhale/issues/4032)

2. **[#2870] EPIC: staged command-boundary refactor**  
   *Large-scope EPIC tracking the layered merge of a critical architectural change. Reference PR #2851 already posted.*  
   [Issue #2870](https://github.com/Hmbown/CodeWhale/issues/2870)

3. **[#4042] Environment-level tool sandboxing for sub-agents**  
   *Runtime counterpart to routing PR #3969. Proposes enforcing tool_restrictions at the sub-agent runtime level.*  
   [Issue #4042](https://github.com/Hmbown/CodeWhale/issues/4042)

4. **[#4010] Conductor agent type for orchestrating agent ensembles**  
   *WhaleFlow introduction: an agent that fans out scouts, waits, routes artifacts, retries, and synthesizes results.*  
   [Issue #4010](https://github.com/Hmbown/CodeWhale/issues/4010)

5. **[#4015] Context budget management for high-fan-out orchestration**  
   *When 30+ sub-agents run, parent context balloons ~1–3KB per completion summary. Proposes budget-aware trimming.*  
   [Issue #4015](https://github.com/Hmbown/CodeWhale/issues/4015)

6. **[#4014] TUI lag and memory pressure from high agent fan-out**  
   *Observed typing latency, rendering stalls, and host memory pressure during parallel sub-agent sessions.*  
   [Issue #4014](https://github.com/Hmbown/CodeWhale/issues/4014)

7. **[#4013] Verification gates (compile, test, lint, review) as post-agent hooks**  
   *Sub-agents self-report “done” without automated checks. Wanted: automated compile/test/lint/review gates.*  
   [Issue #4013](https://github.com/Hmbown/CodeWhale/issues/4013)

8. **[#4038] v0.8.68 Workflow product-readiness tracker**  
   *Umbrella issue listing blockers: no stable model-facing tool, no normal run path, no compact run view, no high-fan-out resource story.*  
   [Issue #4038](https://github.com/Hmbown/CodeWhale/issues/4038)

9. **[#4039] Background task phase ledger UI**  
   *Request for a compact “Background tasks” panel grouped by workflow phase instead of a long chat transcript.*  
   [Issue #4039](https://github.com/Hmbown/CodeWhale/issues/4039)

10. **[#4037] Rename WhaleFlow surfaces to Workflow**  
    *User-facing feature should be called “Workflow”, not “WhaleFlow”. Docs, UI copy, and labels still use the internal name.*  
    [Issue #4037](https://github.com/Hmbown/CodeWhale/issues/4037)

## Key PR Progress (10 important)

1. **[#4046] Layer 5.1: User command registry and loading boundary** *(CLOSED)*  
   *Verifies that existing code already satisfies acceptance criteria for user-defined Markdown/frontmatter commands. No production changes needed.*  
   [PR #4046](https://github.com/Hmbown/CodeWhale/pull/4046)

2. **[#3969] Add per-sub-agent provider routing** *(OPEN, held for v0.8.68)*  
   *Adds explicit provider/model routing for sub-agents, aligned with fleet redesign (#3932–#3935).*  
   [PR #3969](https://github.com/Hmbown/CodeWhale/pull/3969)

3. **[#4045] Fix edit_file UTF-8 fuzzy cursor panic** *(OPEN)*  
   *Fixes panic when a normalized match starts on a multibyte CJK character. Advances cursor to next grapheme cluster.*  
   [PR #4045](https://github.com/Hmbown/CodeWhale/pull/4045)

4. **[#4044] Localize dynamic welcome steps** *(OPEN)*  
   *Localizes first-run welcome screen through existing MessageId registry, adds welcome copy for every shipped locale including zh-Hant.*  
   [PR #4044](https://github.com/Hmbown/CodeWhale/pull/4044)

5. **[#4043] Reset SIGPIPE to SIG_DFL for piped output** *(OPEN)*  
   *Fixes panic when output is piped to an early-exit command (e.g. `codewhale doctor | head`).*  
   [PR #4043](https://github.com/Hmbown/CodeWhale/pull/4043)

6. **[#4041] Remove unused whale_routes taxonomy** *(OPEN)*  
   *Cleanup PR: deletes a dead module (`WhaleRoute`, `WHALE_ROUTES`) with no production callers.*  
   [PR #4041](https://github.com/Hmbown/CodeWhale/pull/4041)

7. **[#4040] Remove legacy token-only pricing helpers** *(OPEN)*  
   *Removes four unused pricing functions. Production cost accounting already uses the usage-aware path.*  
   [PR #4040](https://github.com/Hmbown/CodeWhale/pull/4040)

8. **[#4023] Harden v0.8.67 RC surfaces** *(CLOSED)*  
   *Fixes stream timeout config, plugin paths, onboarding copy, provider routing, OpenAI Codex OAuth messaging, cost display, and sub-agent sidebar policy.*  
   [PR #4023](https://github.com/Hmbown/CodeWhale/pull/4023)

9. **[#4024] Align v0.8.67 QA script with constitution source** *(CLOSED)*  
   *Canonicalizes binary path and updates repo-law assertion to match current doctor --context-json source kind.*  
   [PR #4024](https://github.com/Hmbown/CodeWhale/pull/4024)

10. **[#3972] Allow longer quiet reasoning waits** *(CLOSED)*  
    *Raises idle timeout from 300s to 900s, makes TUI watchdog respect configured stream idle budget plus grace period.*  
    [PR #3972](https://github.com/Hmbown/CodeWhale/pull/3972)

## Feature Request Trends

- **Workflow (WhaleFlow) productization** – Multiple issues (#4010, #4015, #4014, #4013, #4038, #4039, #4037) converge on making agent orchestration production-ready: conductor agents, context budget management, verification gates, compact UI, and rebranding to “Workflow”.
- **Sub-agent tool sandboxing** – Issue #4042 and PR #3969 push for per-sub-agent restrictions on provider/model and allowed tools.
- **Localization improvements** – PR #4044 adds locale support for welcome screens, reflecting internationalisation demand.
- **Cleanup & architectural refactoring** – Issue #2870 (command-boundary refactor) and several cleanup PRs (#4041, #4040, #3849) show ongoing technical debt reduction.

## Developer Pain Points

- **Performance under high concurrency** – Issue #4014 documents severe TUI lag and memory pressure when 30+ sub-agents run in parallel. Developers report “my computer is freezing” symptoms.
- **Constitution / trust violations** – Issue #4032 sparks 19 comments about CodeWhale ignoring user-provided scripts and writing its own, undermining user confidence.
- **Unreliable sub-agent completion** – No automated verification gates (#4013) means sub-agents can claim “done” without passing compile/test/lint, requiring manual polling.
- **Broken pipe crashes** – PR #4043 fixes a panic when piping output to `head` or similar tools – a basic UX friction point.
- **UTF-8 handling** – PR #4045 fixes a panic in edit_file fuzzy cursor for CJK characters, highlighting ongoing encoding edge-case issues.
- **User-facing naming inconsistency** – Issue #4037 complains that the “WhaleFlow” name persists in UI and docs, making the feature feel unfinished and confusing.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*