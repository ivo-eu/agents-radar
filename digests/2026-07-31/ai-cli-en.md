# AI CLI Tools Community Digest 2026-07-31

> Generated: 2026-07-31 00:15 UTC | Tools covered: 9

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

# Cross-Tool Comparison Report: AI CLI Developer Tools
**Date:** 2026-07-31

## Ecosystem Overview

The AI CLI tool ecosystem is experiencing rapid iteration with a clear bifurcation between stability-focused releases (Copilot CLI v1.0.77, OpenCode v1.18.10) and ambitious architectural refactoring (DeepSeek TUI's v0.9.3 rewrite, PI's remote session protocol). Across all tools, **agent reliability and cost governance** have emerged as the dominant pain points—users are encountering runaway billing, silent failures, and subagent control defects at scale. The landscape is also witnessing a convergence toward platform features: persistent memory, desktop GUI experiences, and multi-provider OAuth integration are now baseline expectations rather than differentiators. Notably, **Windows stability** remains a universal weak spot, with every tool reporting platform-specific crashes, sandbox failures, or UI bugs.

## Activity Comparison

| Tool | Hot Issues (Today) | Key PRs (Today) | Release (24h) | Overall Tempo |
|---|---|---|---|---|
| **Claude Code** | 10 | 1 (spam) | ❌ | Bug triage focus; low development velocity |
| **OpenAI Codex** | 10 | 10 | ✅ Alpha (rust-v0.147.0-alpha.2) | High activity; enterprise automation push |
| **Gemini CLI** | 10 | 10 | ✅ Nightly | High feature velocity; sandbox & memory focus |
| **GitHub Copilot CLI** | 10 | 0 | ✅ v1.0.77 | Stable release cycle; OAuth improvements |
| **Kimi Code CLI** | 3 | 1 | ❌ | Low activity; 2 critical blockers |
| **OpenCode** | 10 | 10 | ✅ v1.18.10 | High activity; Web UI & Desktop fixes |
| **Pi** | 10 | 10 | ❌ | Strong infrastructure PRs; protocol & lifecycle |
| **Qwen Code** | 10 | 10 | ✅ Nightly | Very high activity; Anthropic converter fixes |
| **DeepSeek TUI (CodeWhale)** | 10 | 10 | ✅ v0.9.2 | Refactoring phase; runtime consolidation |

## Shared Feature Directions

### 1. Agent Control & Cost Governance *(Present in: Claude Code, Gemini CLI, Copilot CLI, Qwen, Pi)*
- **Subagent lifecycle management:** Users across tools report that `TaskStop`/kill commands don't propagate to child agents, causing token waste. Specific manifestations:
  - Claude Code: #82104 (750k tokens billed after kill)
  - Gemini CLI: #22323 (subagents falsely report success)
  - Copilot CLI: #4293 (subagents with full tool access return empty)
- **Live token/billing visibility:** An implicit demand across all tools, formalized in Claude Code's missing feature request for per-session caps.

### 2. Cross-Session Memory & Context Persistence *(Present in: Claude Code, Kimi Code, OpenCode, DeepSeek TUI, Qwen)*
- **Persistent context across sessions:** Kimi Code (#1283), OpenCode (#39772), and PI (#7163) all have open requests. Qwen Code implemented `memory.agentMaxTurns` (PR #8171).
- **Auto-compression/archival:** Gemini CLI (#28488) merged chat compression; Claude Code reports session auto-archiving as a bug rather than a feature (#71616).

### 3. OAuth & Multi-Provider Authentication *(Present in: Claude Code, Codex, Gemini CLI, Copilot CLI, Qwen, Pi)*
- **MCP OAuth token refresh failures:** Reported by Claude Code (#59854), Codex (#13200), Gemini CLI (#28481), Qwen Code (#8170), and Pi (#7161). The pattern is systemic—OAuth credentials are deleted or corrupted on refresh failure.
- **Multi-account switching:** Claude Code's #36151 (530 👍) is the most-upvoted request, but Copilot CLI and Pi also have similar demands for workspace portability.

### 4. Desktop/GUI Experience *(Present in: Claude Code, Codex, OpenCode, DeepSeek TUI, Qwen)*
- **Beyond CLI:** DeepSeek TUI (#4986) explicitly requests a desktop app. Qwen Code is packaging Web Shell as a desktop app (PR #8132). OpenCode's Desktop app v1.18.10 includes UI refinements. Claude Code's mobile app also demands GUI improvements.

### 5. Session Reliability & Data Persistence *(Present in: Claude Code, Codex, Copilot CLI, Kimi Code, OpenCode, Pi)*
- **Session recovery:** Auto-updates wiping data (Claude Code #43719), sessions auto-archiving on iOS (#71616), and session list empty in Web UI (OpenCode #27837) all point to broken persistence guarantees.
- **Attachment size limits:** Copilot CLI #3767 (5 MB limit) and PI #7337 (tool-result image bloat) show underlying infrastructure constraints.

## Differentiation Analysis

| Dimension | Claude Code | Codex | Gemini CLI | Copilot CLI | DeepSeek TUI | Pi | Qwen Code |
|---|---|---|---|---|---|---|---|
| **Primary Focus** | Agent orchestration & Cowork | Sandbox enforcement & enterprise | Memory & subagent intelligence | OAuth UX & credit transparency | Runtime consolidation & desktop | Extensible protocol & provider parity | Multi-provider conversion & Desktop |
| **Target Users** | Max subscribers (heavy agent users) | Enterprise & MCP ecosystem | Google ecosystem developers | GitHub & VS Code users | Rust/CLI power users | Plugin/extension developers | Multi-model teams (China + global) |
| **Technical Approach** | Monolithic CLI + mobile | Electron desktop + V8 sandbox | Monorepo with nightly releases | GitHub-integrated CLI | Rust monolith (refactoring) | Monorepo + CBOR wire protocol | Dockerized sandbox + TUI |
| **Open Source Maturity** | ✅ Active | ✅ Active | ✅ Active | ✅ Active | ✅ Active | ✅ Active | ✅ Active |
| **Windows Support** | ❌ Poor (heap corruption, OAuth) | ❌ Poor (freezes, crashes) | ❌ Limited (Wayland issues) | ⚠️ Mixed | ⚠️ Mixed (Cygwin paths) | ❌ Poor | ❌ Poor (installer, LMStudio) |
| **Unique Differentiator** | Cowork local-agent mode | Enterprise automation plans | Memory system robustness | OAuth login flow innovation | LaTeX rendering + v0.9.3 rewrite | Remote session protocol | Anthropic content converter |

## Community Momentum & Maturity

**Highest Velocity:** *OpenAI Codex* and *Qwen Code* show the strongest feature throughput—both landed 10+ meaningful PRs in 24 hours, with Codex focusing on sandbox hardening and Qwen tackling Anthropic converter bugs at scale. *Pi* is building foundational infrastructure (remote session protocol, wire encoding, harness lifecycle) that positions it as a platform rather than a tool.

**Highest Community Engagement:** *Claude Code* has the most active bug reporting (530 👍 on multi-account switching), but the conversation is dominated by critical issues (data loss, runaway billing) rather than features. *OpenCode* shows healthy contributor diversity (Kit Langton, OpeOginni, rekram1-node all active today).

**Rapid Iteration Phase:** *Gemini CLI* and *DeepSeek TUI* are both in major refactoring cycles—Gemini with memory/compression features, DeepSeek TUI with runtime consolidation and desktop ambitions. Their releases are frequent but carry regression risks (e.g., DeepSeek TUI's v0.9.2 finalized after handoff fixes).

**Mature & Stable:** *GitHub Copilot CLI* is the most polished—v1.0.77 landed with OAuth and editor integration improvements, and the issue tracker shows fewer critical bugs. The tradeoff is slower feature velocity.

**At Risk:** *Kimi Code CLI* has the lowest activity (3 issues, 1 PR) and two critical blockers (provider overload, CLI freeze). Without rapid response, user trust will erode.

## Trend Signals

1. **Agent trust is the #1 industry problem.** Across all tools, subagents running silently, failing without reporting failure, and burning credits without user consent represents a systemic trust deficit. Tools that solve "agent observability" (live token counters, cancellable subagent trees, clear success/failure reporting) will win developer trust.

2. **OAuth for MCP is a universal bottleneck.** Every tool with MCP integration reports OAuth credential corruption on refresh, DCR unsupported errors, or token leakage. The industry needs a standardized MCP OAuth client implementation—tool teams are reinventing the wheel independently.

3. **"Desktop app" is the new CLI.** DeepSeek TUI, Qwen Code, and OpenCode are all pivoting toward desktop experiences. Pure CLI tools are losing ground to GUI-first approaches (Codex, Copilot). Terminal-based tools must innovate on visualization (LaTeX, rich session management) to compete.

4. **Memory systems are table stakes, not differentiators.** Everyone is building persistent memory—within 12 months, any CLI tool without cross-session context will be considered incomplete. The differentiation will shift to redaction/deterministic privacy (Gemini CLI #26525) and user-controlled memory budgets (Qwen Code #8171).

5. **Windows is the neglected majority.** Every tool reports platform-specific bugs on Windows (heap corruption, freeze after long sessions, OneDrive sync issues, installer failures). A tool that prioritizes Windows parity as a first-class feature could capture significant market share from frustrated users.

6. **Cost transparency is the new UX frontier.** Users are demanding live token counters, credit near-limit warnings (Copilot CLI #4295), per-session caps, and post-session billing reconciliation. The community is moving from "does this tool work?" to "can I afford to run this tool at scale?"

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

Here is the Claude Code Skills community highlights report, based on activity in the `anthropics/skills` repository as of July 31, 2026.

---

## 1. Top Skills Ranking

The most actively discussed pull requests reveal the community’s focus on both novel skill capabilities and critical fixes to the developer toolchain.

1.  **Skill-Creator Fixes (Multiple PRs)** — `#1298`, `#1099`, `#1050`, `#1323`, `#1261`
    - **Functionality:** These PRs address a critical bug in `run_eval.py` where the skill-description optimization loop always reports `recall=0%`, making it impossible to improve skill descriptions automatically. Issues span Windows subprocess handling (missing `PATHEXT`, cp1252 encoding, `select` on pipes), YAML parsing failures, and command-file isolation.
    - **Discussion Highlights:** The most concentrated debugging effort in the repo. Multiple authors independently reproduced the same silent failure. Discussions confirm the loop was "optimizing against noise" for months.
    - **Status:** All open. These are the highest-priority technical debt items for the skill-creator ecosystem.
    - [PR #1298](https://github.com/anthropics/skills/pull/1298) | [PR #1099](https://github.com/anthropics/skills/pull/1099) | [PR #1050](https://github.com/anthropics/skills/pull/1050) | [PR #1323](https://github.com/anthropics/skills/pull/1323) | [PR #1261](https://github.com/anthropics/skills/pull/1261)

2.  **document-typography (`#514`)** — Typographic quality control for generated documents.
    - **Functionality:** Prevents orphan word wrap, widow paragraphs, and numbering misalignment in AI-generated documents—common output-quality issues users rarely request fixes for.
    - **Discussion Highlights:** Strong support as a "background quality" skill. The author argues these issues affect every document Claude generates, making the skill a universal default.
    - **Status:** Open.
    - [PR #514](https://github.com/anthropics/skills/pull/514)

3.  **self-audit (`#1367`)** — Mechanical file verification plus a four-dimension reasoning quality gate.
    - **Functionality:** A meta-skill that audits AI output before delivery: verifies claimed output files exist, then runs a reasoning audit in damage-severity priority order. Universal across models and tech stacks.
    - **Discussion Highlights:** Introduces a structured "pre-delivery QA" paradigm. Follow-up proposal in Issue `#1385` extends this into a three-gate pipeline.
    - **Status:** Open (v1.3.0).
    - [PR #1367](https://github.com/anthropics/skills/pull/1367)

4.  **testing-patterns (`#723`)** — Comprehensive testing stack skill.
    - **Functionality:** Covers the Testing Trophy model, AAA pattern, React Testing Library, Cypress, Playwright, and property-based testing.
    - **Discussion Highlights:** A rare example of a genuinely cross-stack skill that spans philosophy, unit, component, and end-to-end testing in a single cohesive instruction set.
    - **Status:** Open.
    - [PR #723](https://github.com/anthropics/skills/pull/723)

5.  **color-expert (`#1302`)** — Self-contained color expertise for any task.
    - **Functionality:** Covers ISCC-NBS, Munsell, XKCD, RAL, Ridgway 1912 naming systems plus a "what to use when" table for color spaces (OKLCH for scales, OKLAB for gradients, CAM16 for appearance).
    - **Discussion Highlights:** Niche but highly authoritative. The author is a known color-library maintainer, lending domain credibility.
    - **Status:** Open.
    - [PR #1302](https://github.com/anthropics/skills/pull/1302)

6.  **pyxel (`#525`)** — Retro game development with the Pyxel engine.
    - **Functionality:** Integrates the `pyxel-mcp` server to support a write → run_and_capture → inspect → iterate workflow for 8-bit/pixel-art games.
    - **Discussion Highlights:** One of the first MCP-integrated skills. Demonstrates how Skills can bridge Claude to specialized game engines via tool servers.
    - **Status:** Open.
    - [PR #525](https://github.com/anthropics/skills/pull/525)

7.  **ODT (`#486`)** — OpenDocument text creation, template filling, and ODT-to-HTML conversion.
    - **Functionality:** Covers `.odt`, `.ods`, and `.odf` formats with triggers for "LibreOffice document" and "OpenDocument" requests.
    - **Discussion Highlights:** Addresses a gap for open-source document formats. Discussion notes the need to handle LibreOffice-specific rendering quirks.
    - **Status:** Open.
    - [PR #486](https://github.com/anthropics/skills/pull/486)

8.  **plan-file-hygiene (`#1479`)** — Lifecycle management for planning artifacts.
    - **Functionality:** Automatically cleans up stale planning files (`PLAN.md`, `ARCHITECTURE.md`, etc.) that accumulate across long sessions.
    - **Discussion Highlights:** Originated from a community naming and framing effort (Issue `#1417`). The author explicitly credits prior contributors for the problem framing.
    - **Status:** Open.
    - [PR #1479](https://github.com/anthropics/skills/pull/1479)

---

## 2. Community Demand Trends

The Issues reveal concentrated, unsolved pain points and desired new skill directions:

- **Skill-Creator Reliability (Issue `#556`, `#1169`, `#1061`):** The single most commented-on problem. `run_eval.py` reports 0% trigger rate on all queries, making the entire description-optimization pipeline non-functional—especially on Windows. This blocks any community member from iterating on skill quality.

- **Security and Trust Boundary (Issue `#492` — 43 comments):** Community skills distributed under the `anthropic/` namespace create a impersonation vulnerability. Users may grant elevated permissions to skills they believe are official. This has the highest comment count of any issue and is a structural governance concern.

- **Org-Wide Skill Sharing (Issue `#228` — 16 comments):** Demand for a native sharing mechanism. Current workflow (download `.skill` → Slack → manual upload) is friction-heavy. Requested features include shared skill libraries and direct sharing links.

- **Skill Proposal: compact-memory (Issue `#1329`):** A symbolic notation for compact agent state, designed to reduce context overhead from long-running agents' prose notes. This represents demand for **context efficiency tooling** rather than a domain-specific skill.

- **Skill Proposal: agent-governance (Issue `#412`):** Safety patterns including policy enforcement, threat detection, trust scoring, and audit trails. Indicates growing enterprise interest in agent guardrails.

- **Context Window Exhaustion (Issue `#1487`):** The `claude-api` skill was found to inject ~156k tokens in a single tool call, eliminating the entire usable context. This flags a missing **cost-awareness** concern for skill design standards.

- **Duplicate Skill Installs (Issue `#189` — 6 comments, 9 upvotes):** The `document-skills` and `example-skills` plugins install identical content. Users want a clean separation between document-format skills and general-purpose example skills.

---

## 3. High-Potential Pending Skills

These open PRs have active discussion and appear likely to land in the near future:

| PR | Skill | Likelihood |
|---|---|---|
| `#514` | **document-typography** | High — solves a universal output-quality problem with broad appeal. |
| `#1367` | **self-audit** | High — strong rationale, v1.3.0 iteration, follow-up proposal already filed. |
| `#723` | **testing-patterns** | Medium-High — comprehensive and well-structured, though broad scope may require review cycles. |
| `#1302` | **color-expert** | Medium — niche but authoritative author; depends on maintainer bandwidth. |
| `#1479` | **plan-file-hygiene** | Medium — community-driven naming, addresses a clear lifecycle gap. |
| `#525` | **pyxel** | Medium — MCP integration interest is rising, but specialized use case. |
| `#1298` | **skill-creator fixes** | High — community blocking dependency; maintainers likely to prioritize. |

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is to make the skill-development toolchain (especially `run_eval.py`) reliable and cross-platform—because without that, no skill author can confidently improve description quality, and the entire ecosystem's iterative optimization loop is broken.**

---

# Claude Code Community Digest — 2026-07-31

## Today’s Highlights
No new releases landed in the last 24 hours, but the community is actively reporting a wave of **data‑loss** and **runaway‑billing** bugs in background agents, Cowork sessions, and scheduled one‑shots. A long‑standing feature request for **multi‑account switching** on mobile (#36151) continues to dominate engagement with 148 comments and 530 👍, indicating strong unmet demand for workspace portability. Several critical Cowork regressions—including disk corruption on Windows and OAuth failures on macOS—remain unaddressed across multiple builds.

## Releases
*No new versions were published in the last 24 hours.*

## Hot Issues (10 Noteworthy)

1. **#36151** – [FEATURE] Multi‑account switching in Claude Mobile app without shared email  
   *Author: CorneAussems · Comments: 148 · 👍: 530*  
   The most‑upvoted open request. Users want a first‑class account picker in the mobile app to switch between personal and work workspaces without linking a shared email address.  
   [GitHub](https://github.com/anthropics/claude-code/issues/36151)

2. **#77730** – [BUG] Background agent and task IDs stop resolving across a session‑identity boundary (transcripts and outputs remain on disk)  
   *Author: simplysdm · Comments: 7*  
   Background transcripts become unreachable, forcing full‑context respawns and burning tokens. A Max subscriber reports consistent reproducibility.  
   [GitHub](https://github.com/anthropics/claude-code/issues/77730)

3. **#43719** – [BUG] Auto‑update wiped my Cowork session disk — need my projects restored  
   *Author: brandonup · Comments: 5 · 👍: 2*  
   A suspected regression where the auto‑updater deleted local project data mid‑session, causing complete Cowork session loss. Labeled `data-loss` and `regression`.  
   [GitHub](https://github.com/anthropics/claude-code/issues/43719)

4. **#59854** – [BUG] Cowork — GitHub connector unusable: OAuth DCR unsupported, UI shows misleading state, Disconnect button dead  
   *Author: nathanpancakelegion · Comments: 5 · 👍: 12*  
   The GitHub connector in Cowork is broken on macOS: OAuth fails, the UI shows a connected state when it’s actually disconnected, and the Disconnect button is unresponsive.  
   [GitHub](https://github.com/anthropics/claude-code/issues/59854)

5. **#71616** – [BUG] All newly‑created Code sessions auto‑archive on iOS and become inaccessible from mobile app  
   *Author: jackknoebber · Comments: 4*  
   On iOS, any fresh session is immediately archived, making it invisible in the mobile app. A Cowork‑related regression affecting mobile users.  
   [GitHub](https://github.com/anthropics/claude-code/issues/71616)

6. **#82728** – Scheduled one‑shots: 6 of 6 failed — 3 never dispatched, 3 killed mid‑tool‑call  
   *Author: wshallwshall · Comments: 3*  
   All six background one‑shots on a single machine failed. Three never fired (remained “armed”), three were terminated mid‑execution but reported as successful. A merged report of a systematic reliability issue.  
   [GitHub](https://github.com/anthropics/claude-code/issues/82728)

7. **#82104** – TaskStop does not stop subagent children: 750k tokens billed after kill, with no live usage visibility and no cap  
   *Author: simplysdm · Comments: 2*  
   A critical agent‑control defect: killing a parent agent does not propagate to subagents, which continue running and billing. This incident burned 750k tokens with no way to observe or cap usage.  
   [GitHub](https://github.com/anthropics/claude-code/issues/82104)

8. **#80973** – [BUG] Cannot update payment method — “connection to link account has been closed”; support reports Free tier while plan is Max 5x  
   *Author: itsrayforreal · Comments: 2*  
   Billing state inconsistency between the client and backend prevents users from updating payment details, with support seeing a different plan tier.  
   [GitHub](https://github.com/anthropics/claude-code/issues/80973)

9. **#77549** – [BUG] `AskUserQuestion` breaks after a web/async session resume — either aborts instantly or must be answered 2–7×  
   *Author: janpio · Comments: 1 · 👍: 1*  
   Tool permission prompts fail after resuming a web/async session, forcing repeated answers or outright aborting, making async workflows fragile.  
   [GitHub](https://github.com/anthropics/claude-code/issues/77549)

10. **#79575** – [BUG] `/fork` blocked in sessions launched with `--dangerously-skip-permissions` — rationale inverted  
    *Author: Butanium · Comments: 1*  
    The fork guard incorrectly claims the child would run “with fewer restrictions”, even though the parent already skipped all permissions. The check logic is inverted.  
    [GitHub](https://github.com/anthropics/claude-code/issues/79575)

## Key PR Progress

Only **one** pull request was updated in the last 24 hours:

- **#82555** – `Claude/youtube instagram mcp yn2u6s` (Author: batuhunca-del · Closed)  
  This PR appears to be a spam or test submission—no meaningful diff. It was closed without merge. No impactful PR activity to report today.  
  [GitHub](https://github.com/anthropics/claude-code/pull/82555)

*Note: The repository saw no significant code‑change proposals in this window. The focus remains squarely on bug triage and community feedback.*

## Feature Request Trends

1. **Account & Workspace Management** – The overwhelming demand is for **multi‑account switching** (#36151) across mobile, desktop, and CLI. Users need to separate personal/professional workspaces without sharing emails or logging out.

2. **Security & Data Privacy** – Request for **in‑memory storage for background task outputs** (#82734) to prevent sensitive data from being written to disk. Reflects growing use of Claude Code in compliance‑sensitive environments.

3. **Agent Team Control** – Ask for a **blocking/priority field in agent frontmatter** (#69391), parallel to the existing skill blocking mechanism, to prevent unwanted agent auto‑activation.

4. **Better Observability** – Implicit in many bug reports: users want **live token usage visibility** and **capping per session** (#82104, #77730). Not yet a formal feature request, but clearly the next frontier.

## Developer Pain Points

- **Data Loss & Session Reliability** – Auto‑updates wiping Cowork session disks (#43719), iOS sessions auto‑archiving (#71616), and background transcripts becoming unresumable (#77730) erode trust in session persistence across platforms.

- **Token Waste & Uncontrolled Billing** – Subagent run‑on after `TaskStop` (#82104) and unresolvable background IDs leading to forced respawns (#77730) burn tokens—often hundreds of thousands—without user visibility or abort capabilities.

- **Cowork Fragility** – The “local‑agent‑mode” feature suffers from multiple regressions: broken GitHub OAuth (#59854), kernel‑mode heap corruption on Windows (#72377), disabled plugins being synced (#68020), and permission prompts that can’t be approved from mobile (#69371). The feature feels unfinished.

- **Permission & Workflow Inconsistencies** – The `/fork` guard with `--dangerously-skip-permissions` (#79575) and `AskUserQuestion` failures after resume (#77549) show that the permission model has edge cases that break legitimate workflows, forcing users to find workarounds or restart sessions.

- **Scheduled Task Reliability** – 6/6 scheduled one‑shots failing (#82728) signals that background execution paths are not production‑ready, with no retry logic or user notification.

*Digest generated from GitHub data for `github.com/anthropics/claude-code` on 2026-07-31.*

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# Codex Community Digest – 2026-07-31

## Today’s Highlights
A fresh alpha release (`rust-v0.147.0-alpha.2`) landed, though no changelog is attached. Windows stability continues to dominate community attention: the long‑running UI freeze bug (#20214, 83 comments) remains unresolved, while a new crash in `CrBrowserMain` (#32683) and a OneDrive‑backed workspace disconnect issue (#35420) add to the platform’s pain points. On the PR side, the team is busy hardening sandbox enforcement, improving streaming performance, and rolling out enterprise automation support.

## Releases
- **rust-v0.147.0-alpha.2**  
  [openai/codex/releases](https://github.com/openai/codex/releases) – no description provided; likely a routine nightly.

## Hot Issues (10 noteworthy)

1. **#20214 – Codex App freezes/stutters on Windows 11 Pro**  
   [openai/codex/issues/20214](https://github.com/openai/codex/issues/20214)  
   *83 comments, 77 👍* – The most upvoted open bug. Users report stutters despite ample hardware (Ryzen 5, 32 GB). No fix yet; community suspects render‑thread contention.

2. **#31573 – OAuth authentication fails at issuer validation**  
   [openai/codex/issues/31573](https://github.com/openai/codex/issues/31573)  
   *31 comments, 66 👍* – Affects CLI users on Free plan. Breaks MCP login flows. High visibility due to impact on multi‑provider setups.

3. **#32683 – App crashes in CrBrowserMain (0xC0000005) when Browser Use opens a page**  
   [openai/codex/issues/32683](https://github.com/openai/codex/issues/32683)  
   *29 comments, 8 👍* – Windows‑only crash in Chromium embedded browser component. Crashing on page load severely limits Codex’s browser‑automation skill.

4. **#26234 – Flatten MCP namespace tools for non‑OpenAI providers**  
   [openai/codex/issues/26234](https://github.com/openai/codex/issues/26234)  
   *27 comments, 40 👍* – Models like Ollama, LM Studio, and AWS Bedrock cannot call MCP tools because Codex serializes them in a proprietary `namespace` format. Blocks local/custom model users.

5. **#26478 – Windows spellcheck shows “No Guesses Found”**  
   [openai/codex/issues/26478](https://github.com/openai/codex/issues/26478)  
   *18 comments, 25 👍* – Spellcheck UI dead‑ends even though Windows native spellcheck works. Minor but persistent annoyance for many Windows users.

6. **#35420 – Stream disconnects when workspace is OneDrive‑backed**  
   [openai/codex/issues/35420](https://github.com/openai/codex/issues/35420)  
   *15 comments* – Newly reported. Degraded OneDrive sync causes `stream disconnected before completion` errors. Relates to file system latency.

7. **#13200 – `codex mcp login` fails for Slack with “Dynamic client registration not supported”**  
   [openai/codex/issues/13200](https://github.com/openai/codex/issues/13200)  
   *10 comments, 58 👍* – Enterprise users blocked from using official Slack MCP. Long‑standing (March 2026) with no fix in sight.

8. **#35362 – VS Code diff crashes; inline diff works**  
   [openai/codex/issues/35362](https://github.com/openai/codex/issues/35362)  
   *10 comments, 13 👍* – Full Review/diff page renders blank or crashes. Inline diff is fine – suggests a rendering issue in the diff viewer.

9. **#18620 – Sandboxed shell commands fail with `CreateProcessWithLogonW` errors**  
   [openai/codex/issues/18620](https://github.com/openai/codex/issues/18620)  
   *9 comments, 5 👍* – Windows sandbox workspace breaks shell execution. Error codes 1326 / 1909 point to credential or session issues.

10. **#35097 – MultiAgent V1 model rejected by V2 `spawn_agent`**  
    [openai/codex/issues/35097](https://github.com/openai/codex/issues/35097)  
    *6 comments, 13 👍* – `gpt-5.6-luna` marked as MultiAgent V1, so V2 sub‑agent spawning fails. Model compatibility regression affecting advanced agent pipelines.

## Key PR Progress (10 notable)

1. **#36239 – Refresh precomputed app‑server protocol exports**  
   [openai/codex/pull/36239](https://github.com/openai/codex/pull/36239)  
   Adds `ExternalAgentConfigDetectResponse` with connector candidates and session counts. Enterprise automation plan added.

2. **#36237 – Ignore symbolic `/tmp` permissions on Windows**  
   [openai/codex/pull/36237](https://github.com/openai/codex/pull/36237)  
   Fixes sandbox policy decisions that incorrectly considered Unix `/tmp` on Windows. Directly addresses workspace sandbox failures.

3. **#36228 – Support Enterprise automation account plans**  
   [openai/codex/pull/36228](https://github.com/openai/codex/pull/36228)  
   Recognizes `enterprise_cbp_automation` plan throughout auth and rate‑limit APIs.

4. **#36217 – Run code mode exclusively through standalone host**  
   [openai/codex/pull/36217](https://github.com/openai/codex/pull/36217)  
   Moves V8 runtime out of main process into dedicated `codex-code-mode-host`. Reduces memory contention and isolates crashes.

5. **#36194 – Avoid shifting bytes in streaming output buffers**  
   [openai/codex/pull/36194](https://github.com/openai/codex/pull/36194)  
   Replaces `Vec::drain` with a ring‑buffer approach. Expected to improve latency for long‑running `exec` streams on Windows/Linux.

6. **#36207 – Record normalized sandbox violation events**  
   [openai/codex/pull/36207](https://github.com/openai/codex/pull/36207)  
   Introduces a common event shape for filesystem denials and network blocks. Aims to make debugging sandbox issues easier for developers.

7. **#31458 – exec‑server: route remote network policy decisions**  
   [openai/codex/pull/31458](https://github.com/openai/codex/pull/31458)  
   Architecturally important: separates proxy policy evaluation from the exec‑server process, adding fail‑closed behavior on disconnect.

8. **#31922 – core: add tool‑free thread mode**  
   [openai/codex/pull/31922](https://github.com/openai/codex/pull/31922)  
   Opt‑in lightweight threads (e.g., title generation) that skip MCP/skill startup. Reduces overhead for helper tasks.

9. **#31591 – Enable parallel tool calls for Codex Apps**  
   [openai/codex/pull/31591](https://github.com/openai/codex/pull/31591)  
   Disabled by default feature flag to allow parallel tool invocation for host‑owned MCP servers. Boosts throughput for multi‑tool workflows.

10. **#36184 – Coalesce concurrent remote metadata requests**  
    [openai/codex/pull/36184](https://github.com/openai/codex/pull/36184)  
    Deduplicates parallel `fs/getMetadata` RPCs for the same path. Reduces network chatter and sandbox overhead.

## Feature Request Trends
- **VS Code notification integration** (#26555, #31519): Users want native completion and approval notifications inside the editor.
- **LaTeX rendering in TUI/CLI** (#36233): Strong visual evidence suggests a demand for math typesetting in terminal‑based Codex.
- **Rate limit fairness for Plus users** (#36213): With GPT‑5.6 Sol, Plus users feel throttled; requests for secondary rate limits on small models.
- **Re‑open closed side chats** (#27716): Side chat history is unrecoverable after closing, a common workflow blocker on Desktop.

## Developer Pain Points
- **Windows stability** remains the top frustration: freezes (#20214), crashes (#32683, #35681), sandbox failures (#18620, #35803), and spellcheck bugs (#26478) create a fractured experience.
- **MCP / multi‑provider friction**: OAuth validation (#31573), tool namespace flattening (#26234), and Slack registration (#13200) block users who rely on custom or third‑party model endpoints.
- **Backward‑compatibility breaks**: MultiAgent version mismatches (#35097) and model regression (#36229) erode trust in new releases.
- **Sandbox diagnostics mislead**: #35803 highlights that `codex doctor` reports “healthy” even when the workspace is corrupted and a reinstall doesn’t fix it.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest – 2026-07-31

## Today’s Highlights
The team merged a critical fix for `MODEL_CAPACITY_EXHAUSTED` errors that were causing client‑side hangs, and landed a security‑critical MCP OAuth token refresh patch. On the issue tracker, a high‑visibility P1 bug (#22323) shows subagents falsely reporting success after hitting turn limits, while the long‑standing generalist agent hang (#21409) continues to draw community attention. Several PRs point to ongoing work on sandbox stability, model selector improvements, and automatic chat compression.

## Releases
No stable release today. The latest nightly (`v0.55.0-nightly.20260730.gdc859e8e4`) primarily bumps the version and includes changelogs for `v0.54.0-preview.0` and `v0.53.0`.

## Hot Issues (10 noteworthy)

1. **#22323** – *Subagent recovery after MAX_TURNS is reported as GOAL success*  
   `codebase_investigator` marks itself `success: "GOAL"` even when it hit the turn limit without doing any analysis, misleading users. 12 comments, 2 👍.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/22323)

2. **#21409** – *Generalist agent hangs forever*  
   The agent freezes when deferring to subagents for simple tasks (e.g., folder creation). Users can only work around it by instructing the model not to use subagents. 8 comments, 8 👍.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/21409)

3. **#25166** – *Shell command execution gets stuck with “Waiting input” after command completes*  
   After running trivial CLI commands, Gemini CLI hangs while showing the command as active and awaiting input. 4 comments, 3 👍.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/25166)

4. **#21983** – *Browser subagent fails in Wayland*  
   The browser agent terminates with `GOAL` but fails silently on Wayland sessions. 4 comments, 1 👍.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/21983)

5. **#22093** – *(Sub)agents running without permission since v0.33.0*  
   Users who disabled subagents find them executing anyway after an automatic update. 3 comments.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/22093)

6. **#24246** – *Gemini CLI encounters 400 error with >128 tools*  
   When the available tool count exceeds 128, the API returns a 400 error. The agent should cap tools in scope. 3 comments.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/24246)

7. **#26522** – *Stop Auto Memory from retrying low-signal sessions indefinitely*  
   The memory extraction agent keeps re‑reading low‑value transcripts because the session is never marked processed. 5 comments.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/26522)

8. **#26525** – *Add deterministic redaction and reduce Auto Memory logging*  
   Secrets can leak into model context because redaction is only prompt‑based; logs may also contain skill code. 4 comments.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/26525)

9. **#23571** – *Model frequently creates tmp scripts in random spots*  
   Gemini scatters temporary scripts across directories, creating cleanup overhead. 3 comments.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/23571)

10. **#22465** – *Gemini CLI gets stuck at interactive prompt creating vite app*  
    When asked to scaffold a Vite app, the agent freezes on an interactive prompt. 2 comments.  
    [Link](https://github.com/google-gemini/gemini-cli/issues/22465)

## Key PR Progress (10 important)

1. **#28599** – `fix(core): classify capacity exhaustion as terminal to prevent retry hangs`  
   **CLOSED.** Treats `MODEL_CAPACITY_EXHAUSTED` (429) as a terminal error, stopping infinite retry loops for preview models.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28599)

2. **#28481** – `fix(core): refresh MCP OAuth tokens with the stored client ID`  
   **OPEN (P1/security).** Prevents token refresh failures that previously deleted stored credentials, forcing re‑auth on every session.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28481)

3. **#28485** – `fix(cli): add gemini-3.5-flash to model selector for all users`  
   **OPEN (P2).** Resolves the issue where users on v0.51.0 couldn’t select Gemini 3.5/3.6 Flash models from the UI.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28485)

4. **#28488** – `feat(cli): auto-compress chat history on context window overflow`  
   **OPEN.** Adds a setting to automatically compress history when the context window is about to overflow, instead of aborting with a warning.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28488)

5. **#28603** – `fix(docker): upgrade sandbox Dockerfile to Node 22`  
   **OPEN (P1/security).** Moves from EOL Node 20 to Node 22 in the sandbox runtime, addressing supply‑chain risk for model‑directed commands.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28603)

6. **#28581** – `fix(cli): skip diff hunk markers during @ processing`  
   **OPEN (P2).** Prevents `@file` globbing from matching diff hunk markers, avoiding recursive workspace scans that cause heap growth.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28581)

7. **#28551** – `fix(cli): fall back to embedded macOS seatbelt profiles if missing`  
   **OPEN.** Fixes startup crash on macOS when static `.sb` seatbelt files are not found in runfiles, critical for sandbox mode.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28551)

8. **#28566** – `fix(core,cli): propagate InvalidStreamError details to UI for specific empty response guidance`  
   **OPEN (P1).** Surfaces actionable error messages (e.g., recommend `/compress`) when the API returns an empty/invalid stream.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28566)

9. **#28410** – `fix(core): shorten MCP tools/list discovery timeout so it fails fast`  
   **CLOSED (P1).** Prevents a 10‑minute startup freeze when an MCP server doesn’t respond to `tools/list`.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28410)

10. **#28596** – `feat(cli): add --list-all-sessions option to list sessions across all workspaces`  
    **OPEN (P3).** New CLI flag to view and manage sessions grouped by workspace, addressing a frequent user request for session discovery.  
    [Link](https://github.com/google-gemini/gemini-cli/pull/28596)

## Feature Request Trends
The most‑requested feature directions fall into four themes:

- **Agent intelligence and autonomy**: AST‑aware file reads and codebase mapping (#22745), improved subagent self‑awareness (#21432), and better tool‑usage heuristics.
- **Memory system robustness**: Deterministic redaction and minimal logging (#26525), automatic session cleanup (#26522), and quarantining invalid memory patches (#26523).
- **Sandboxing and security**: Zero‑dependency OS sandboxing that leverages the model’s bash affinity (#19873) and secure handling of secrets in skill code.
- **Evaluation infrastructure**: Component‑level behavioral evaluations (#24353) and shareable subagent trajectories via `/chat share` (#22598).

## Developer Pain Points
Recurring frustrations include:

- **Agent hangs and false successes**: Subagents silently fail or report success after hitting limits (#22323, #21409), and the CLI freezes on simple commands (#25166, #22465).
- **Permission and configuration bypass**: Subagents run even when disabled (#22093) and ignore settings overrides (#22267).
- **Environment‑specific failures**: Browser agent breaks on Wayland (#21983), and sandbox mode crashes on macOS without seatbelt profiles.
- **Terminal and UI quirks**: Flicker on resize (#21924), corruption after external editors (#24935), and high CPU from recursive path globbing (#28581).
- **Memory and MCP fragility**: Auto Memory loops on low‑signal sessions (#26522), OAuth tokens get deleted on refresh failure (#28481), and 128‑tool limit triggers 400 errors (#24246).

*Generated from GitHub data for `google-gemini/gemini-cli` on 2026-07-31.*

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-07-31

## Today’s Highlights
The Copilot CLI team shipped **v1.0.77** yesterday with two key quality-of-life improvements: a new browser-based OAuth login flow (now the default for local terminals) and the ability to edit `ask_user` freeform answers directly in your editor via `Ctrl+G`. On the bug side, a critical session‑wedging issue caused by oversized attachments has been resolved, but fresh reports around credit consumption after task completion and sub‑agent failures with full tool access are drawing community attention.

## Releases
Two versions landed in the last 24 hours:

- **[v1.0.77](https://github.com/github/copilot-cli/releases/tag/v1.0.77)** (2026-07-30)
  - Unconditional autopilot approval now disables sandbox for the current session when bypass is allowed.
  - `Ctrl+G` opens your editor to edit `ask_user` freeform answers without closing the prompt.
  - Added a browser-based (web) OAuth login flow; it is now the default for `copilot login` on local interactive terminals. Device code flow remains the default on remote/headless terminals. Use `--web-flow` or `--device-code` to force a mode.

- **[v1.0.77-0](https://github.com/github/copilot-cli/releases/tag/v1.0.77-0)** (2026-07-30)
  - Same OAuth improvements (added precise `--web-flow`/`--device-code` flags and interactive `/login` command).
  - "Support enfor…" (description truncated in the raw data; see release notes for full details).

Note: v1.0.76 from 2026-07-29 added plugin enable/disable controls, Grok‑4.5 model support, and sandbox denied‑path enforcement for relative/symlink entries on macOS/Linux.

## Hot Issues
10 noteworthy issues updated in the last 24 hours, selected for community impact and severity:

1. **[Issue #3767](https://github.com/github/copilot-cli/issues/3767) – Oversized attachment wedges session permanently**  
   *Closed* | 13 comments | 1 👍  
   A 9.1 MB attachment hits CAPI’s 5 MB native limit and leaves the session unrecoverable. The fix was included in a recent release; critical for users who paste large images or documents.

2. **[Issue #4295](https://github.com/github/copilot-cli/issues/4295) – AI Credits Near-Limit Warning**  
   *Open* | 8 comments  
   Feature parity request: VS 2026 Professional already warns users when they’re nearing their AI credit cap. The CLI lacks this, forcing users to monitor manually.

3. **[Issue #1381](https://github.com/github/copilot-cli/issues/1381) – “Rewind is not available because you're not in a git repository.”**  
   *Open* | 4 comments | **10 👍**  
   Long‑standing pain point: users of other VCS (e.g., jj) cannot use the Rewind feature. The VS Code extension supports it without git, so the CLI is lagging.

4. **[Issue #4258](https://github.com/github/copilot-cli/issues/4258) – Interactive `-i` startup prompt ignored with custom/BYOK provider**  
   *Closed* | 3 comments  
   A bug that caused the `-i` flag to be silently dropped when using a bring‑your‑own‑key provider in TTY sessions. Fixed in a recent release.

5. **[Issue #4266](https://github.com/github/copilot-cli/issues/4266) – Generic Exit Command Bug (no exit screen)**  
   *Closed* | 2 comments  
   Race condition in shutdown caused the session‑ID exit screen to be missing on normal exit (Ctrl+C/D, `/exit`). Workaround: `/exit print`. Now fixed.

6. **[Issue #4293](https://github.com/github/copilot-cli/issues/4293) – Sub-agents with full tool access return empty; restricted‑tool agents work**  
   *Open* | 2 comments  
   A puzzling bug: sub‑agents launched via the `task` tool produce zero output when they have access to the full tool set, but work fine with a restricted set. No error or logs. Community suspects a tool‑routing issue.

7. **[Issue #4310](https://github.com/github/copilot-cli/issues/4310) – Bad default: engine falls back to 128K token budget for model**  
   *Open* | 0 comments  
   When a model reports a zero context window, the agent engine defaults to a hardcoded 128K token budget, causing unnecessary context compaction for large‑context models (e.g., 1M‑token Anthropic). Needs model‑aware limit detection.

8. **[Issue #4308/4309](https://github.com/github/copilot-cli/issues/4308) – Session continues consuming AI credits after all visible tasks complete**  
   *Open* | 0 comments each (two reports)  
   Users report ~97.8% credit usage even though no visible work was happening. Could indicate background agent activity or a billing‑logic bug. High importance for enterprise users with limited credits.

9. **[Issue #4305](https://github.com/github/copilot-cli/issues/4305) – “Failed to convert JavaScript value 'Undefined' into rust type 'String'”**  
   *Closed* | 0 comments  
   After upgrading to v1.0.76, some users hit a Rust‑JS interop crash on any `/model` command. Likely a configuration parsing error; likely fixed in the latest patch.

10. **[Issue #4299](https://github.com/github/copilot-cli/issues/4299) – Increasing typing latency over long copilot sessions**  
    *Open* | 0 comments | 1 👍  
    Typing becomes unusably slow after hours of work, especially with background agents running. Echoes similar reports from other CLI‑based AI tools.

## Key PR Progress
No notable pull requests were merged or updated in the last 24 hours. The development focus appears to be on bug fixes and the OAuth feature.

## Feature Request Trends
The most‑requested directions from recent issues include:

- **Credit and usage transparency** – Users want near‑limit warnings, accurate post‑session credit accounting, and better visibility into background credit consumption.
- **Authentication flexibility** – Requests for bearer‑token auth (e.g., for corporate compliance) and support for custom auth brokers.
- **Sandbox configurability** – Selective tool whitelisting inside the sandbox to enforce least‑privilege policies.
- **Terminal and input parity** – Clipboard paste (Cmd+V), mouse scroll, and `Ctrl+G` editor editing are all being fine‑tuned based on feedback.

## Developer Pain Points
Recurring frustrations that emerged from this week’s issues:

- **VCS‑agnostic features** – Rewind and other session‑control features that depend on git exclude users of Mercurial, jj, etc.
- **Attachment size limits** – The hard 5 MB CAPI limit with no graceful recovery or compression hints.
- **Sub‑agent reliability** – Sub‑agents failing silently with full tool sets, freezing, or returning empty responses wastes developer time.
- **Long‑session degradation** – Typing latency and memory growth make hour‑long sessions painful.
- **MCP tool argument handling** – Union schemas with `anyOf` (array‑or‑string) are incorrectly stringified, breaking tool integrations.
- **Log level crashes** – Setting anything other than “all” or “default” causes a startup crash; a trivial but irritating regression.

---

*Generated automatically from data retrieved on 2026-07-31. All links point to the official [github/copilot-cli](https://github.com/github/copilot-cli) repository.*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest – 2026-07-31

## Today's Highlights

The past 24 hours brought no new releases but saw three open issues and one pull request. A critical connectivity issue (#2571) and a CLI freeze bug (#2570) are the main blockers for users, while a long-standing feature request for persistent memory (#1283) continues to gather attention. A PR fixing weak-reference hook cleanup (#2565) moves toward merge.

## Releases

No new releases in the last 24 hours.

## Hot Issues

- **#2571 – [bug] LLM Overloaded! Can't use Kimi at all**  
  *Author: andrew-sz* | [Link](https://github.com/MoonshotAI/kimi-cli/issues/2571)  
  Newly filed today (2026-07-30). A user on Moderato subscription with Kimi K3 receives HTTP 429 (rate limit / overload error). This is a show-stopper for anyone relying on the tool right now. Community will likely escalate if the provider-side throttling isn’t mitigated soon.

- **#2570 – [bug] CLI intermittently freezes with spinning moon; correlated with browser tab state**  
  *Author: XbackMK* | [Link](https://github.com/MoonshotAI/kimi-cli/issues/2570)  
  Windows 11, KIMI Login Subscription, version 0.29.2. The CLI becomes unresponsive, and the user reports a correlation with browser tab state. Possibly a local resource or authentication token renewal issue. No comments yet, but may indicate a deeper stability problem.

- **#1283 – [enhancement] Feature Request: Memory System – Persistent context across sessions**  
  *Author: CatKang* | [Link](https://github.com/MoonshotAI/kimi-cli/issues/1283)  
  Opened February 2026 but still open. Proposes both automatic (AI-managed) and manual (user-defined) memory. 7 comments, no thumbs-up yet. This is a high-value feature for developers who use Kimi Code CLI iteratively on long-running projects.

## Key PR Progress

- **#2565 – fix(hooks): keep a strong reference to fire-and-forget hook triggers**  
  *Author: LHMQ878* | [Link](https://github.com/MoonshotAI/kimi-cli/pull/2565)  
  Fixes issue #2564 where `asyncio` tasks stored in a `WeakSet` could be garbage-collected prematurely. The PR ensures that fire-and-forget hook triggers retain a strong reference, preventing silent task loss. This is a critical stability fix for the hook system.

## Feature Request Trends

The single enhancement request (#1283) for a persistent memory system dominates the conversation. Users want context retention (project patterns, preferences, AI-managed notes) across CLI sessions. No other feature requests appeared in the last 24 hours, but the age of #1283 suggests that memory remains the top desired capability.

## Developer Pain Points

1. **Provider overload / rate limiting** (#2571) – users on paid subscriptions are completely blocked by 429 errors. This is an urgent infrastructure concern.  
2. **CLI freeze & unresponsiveness** (#2570) – intermittent freezes correlated with browser state on Windows, indicating possible token refresh or UI thread contention.  
3. **Weak reference / task cancellation bugs** (#2565 fix) – underlying stability issue with async hooks, already addressed by the pending PR.

No additional pain points emerged in the last 24 hours, but the two new bugs (#2571, #2570) should be prioritised alongside the long-standing memory feature request.

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-07-31

## Today's Highlights
OpenCode **v1.18.10** ships with automatic model discovery for Modal providers and several Desktop UI refinements. The community is reporting a **GPT-5.6 Sol model overload** issue (10 👍, 16 comments), while multiple **compliance-oriented issues** from openchat-ai highlight gaps in timeout handling, blocking operations, and debugging loops. On the PR front, Kit Langton contributed session title improvements, and OpeOginni fixed a stale-read crash in the Desktop app.

## Releases
**v1.18.10** — released within the last 24 hours.
- **Core:** Modal models are now automatically discovered.
- **Desktop:** Prevents duplicate attachment uploads; always shows the new session button; improved toast stacking, dismissal, and mobile layout; refined tab hover/active states.

## Hot Issues (10 noteworthy)
1. **[#39653 – GPT-5.6 Sol, server overloaded errors](https://github.com/anomalyco/opencode/issues/39653)**  
   High activity (16 comments, 10 👍). Users hitting persistent “server overloaded” errors with the Sol model; Pi and Codex work fine. Critical for teams relying on Sol for production coding.

2. **[#39491 – Plan mode can write and edit files via bash](https://github.com/anomalyco/opencode/issues/39491)**  
   Claude Sonnet 4.6 escaped plan mode and used `cat >` via bash to create files. A security and workflow concern – plan mode must be strictly read-only.

3. **[#27837 – Web UI: session list empty on left panel when using web server mode](https://github.com/anomalyco/opencode/issues/27837)**  
   Long-standing bug (since May). Despite `/api/session` returning data, the SSE-driven frontend fails to populate the session list. Affects all `opencode --web` users.

4. **[#39655 – OpenCode Web shows “No folders found” although projects are returned by the backend](https://github.com/anomalyco/opencode/issues/39655)**  
   Similar Web UI gap – home page and “Open Project” dialog appear empty. Points to frontend–backend sync issues.

5. **[#39773 – Empty git repos share the same “global” project ID, causing sessions to leak](https://github.com/anomalyco/opencode/issues/39773)**  
   All repos without commits or remotes fall back to `project_id = "global"`, making sessions from unrelated directories appear together. A data isolation bug.

6. **[#39772 – Debugging loops at same layer + no cross-session memory wastes time](https://github.com/anomalyco/opencode/issues/39772)**  
   The assistant circles the same root-cause layer; no cross-session memory leads to repeated guesswork. Tagged `needs:compliance`. High relevance for CI-driven workflows.

7. **[#39771 – No fast failure on network errors, verbose output wastes tokens](https://github.com/anomalyco/opencode/issues/39771)**  
   No short timeout for network ops, no fallback. Especially painful in regions with flaky connectivity (e.g., China). Requests wasting tokens on repeated retries.

8. **[#39769 – Long-running shell commands block the entire conversation](https://github.com/anomalyco/opencode/issues/39769)**  
   Commands like `gh run watch` or polling loops freeze the UI entirely. No cancel, no parallel work possible – a fundamental UX blocker.

9. **[#39766 – Desktop crashes with stale Show read when opening a cross-project session](https://github.com/anomalyco/opencode/issues/39766)**  
   Crashes with `Error: Stale read from <Show>.` since v1.18.10. Reproducible when opening a session from a different project than the active one.

10. **[#39763 – TUI: mouse wheel scrolls by message blocks – should scroll line-by-line](https://github.com/anomalyco/opencode/issues/39763)**  
    On Windows Terminal the mouse wheel jumps 3–4 lines per tick instead of smooth line-by-line scrolling. Degraded reading experience.

## Key PR Progress (10 important)
1. **[#27554 – feat: local LAN provider discovery + auto-discover models](https://github.com/anomalyco/opencode/pull/27554)**  
   Adds mDNS/SSDP discovery for local OpenAI-compatible servers, plus automatic model listing. A long-requested feature (closes #6231, #27553). Still open – community eager to see it merged.

2. **[#39748 – fix(session): retry failed title generation](https://github.com/anomalyco/opencode/pull/39748)**  
   Automatic title generation now retries after a failed first step and respects the original user prompt. Closes #39529. Contributor: Kit Langton.

3. **[#39747 – feat(session): make generated titles optional](https://github.com/anomalyco/opencode/pull/39747)**  
   Sessions stay genuinely `NULL` until generation succeeds or a user renames them. All surfaces (TUI, CLI, search, export, etc.) handle missing titles. Kit Langton.

4. **[#39767 – fix(app): prevent stale session tab reads](https://github.com/anomalyco/opencode/pull/39767)**  
   Fixes the Desktop crash (#39766 / #39704) by protecting titlebar content during Solid transitions. Contributor: OpeOginni.

5. **[#39770 – fix(app): prevent file tree tab clipping](https://github.com/anomalyco/opencode/pull/39770)**  
   Ensures the File Changes tab has a minimum width so it never gets clipped when the file tree is resized small. OpeOginni.

6. **[#39752 – feat(tui): add open menu for sessions and projects](https://github.com/anomalyco/opencode/pull/39752)**  
   `ctrl+o` in v2 TUI opens a unified dialog to jump to sessions or switch projects. Also fixes the all-projects toggle state persistence. Kit Langton.

7. **[#39768 – fix(tui): name deleted session in toast](https://github.com/anomalyco/opencode/pull/39768)**  
   Toast now says `Session "OpenCode Drive" was deleted` instead of generic “The current session was deleted”. Kit Langton.

8. **[#39764 – feat(plugin): add session request hook](https://github.com/anomalyco/opencode/pull/39764)**  
   Exposes `session.request` on plugin boundaries, allowing plugins to mutate HTTP headers and request bodies after native serialization. Enables custom middleware. rekram1-node.

9. **[#38360 – fix(core): configure Figma MCP OAuth client](https://github.com/anomalyco/opencode/pull/38360)**  
   Adds a built-in V2 plugin that provides OpenCode’s registered OAuth client ID for `mcp.figma.com` servers. Preserves explicit `client_id` overrides. rekram1-node.

10. **[#39585 – fix(tui): focus palette settings after layout](https://github.com/anomalyco/opencode/pull/39585)**  
    Settings opened from the command palette now scroll into view and are immediately visible/active. Fixes “Sounds” search, for example. Kit Langton.

## Feature Request Trends
- **Better Web UI stability:** Issues #27837 and #39655 both show core Web UI features (session list, project folders) failing due to frontend/backend sync problems.
- **Enhanced provider integration:** #29935 (LiteLLM proxy) was closed but remains a requested direction; #27554 (LAN discovery) is still open and highly anticipated. Users want easier access to local and proxy-based models.
- **Compliance & developer experience:** Three issues from openchat-ai (#39772, #39771, #39769) all tagged `needs:compliance` highlight demands for timeouts, non-blocking commands, cancellation, and cross-session memory. These resonate with teams running automated CI debugging.
- **Documentation clarity:** #39256 asks for explicit naming conventions (camelCase vs snake_case) in model variant documentation.

## Developer Pain Points
- **Server overload with Sol model** (#39653) – impacting heavy users; no clear workaround yet.
- **Plan mode bypass** (#39491) – undermines trust in read-only mode for planning.
- **Web UI session list emptiness** (#27837, #39655) – core navigation broken for web-server users.
- **Cross-project session leakage** (#39773) – data isolation failure in empty repos.
- **Desktop crash** (#39766) – stale read error makes cross-project navigation unstable.
- **TUI scrolling** (#39763) – Windows users suffer from chunky scroll behavior.
- **Blocking operations** (#39769) – long-running commands freeze the entire interface.
- **No network fallback** (#39771) – wasted tokens and time in flaky environments.
- **Debugging loops** (#39772) – lack of cross-session memory and deeper root-cause analysis frustrates iterative development.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest – 2026-07-31

## Today’s Highlights
Today’s commits focus on strengthening the core platform: a **remote session wire protocol** and **harness shutdown lifecycle** landed in the monorepo, enabling safe, stateful agent backends. On the user-facing side, the long‑standing **Wayland clipboard issue** (#7248) has been fixed with `wl‑paste` support, and a new **Markdown mutation API** (#6747/PR #7231) gives extensions control over agent message rendering. Several closed‑source‑provider bugs (Anthropic missing `x‑client‑request‑id`, Gemini 3 tool‑call IDs stripped) remain open and are drawing active community discussion.

## Releases
No new releases were published in the last 24 hours.

## Hot Issues
1. [#6747 [CLOSED] API for enhancing agent message markdown](https://github.com/earendil-works/pi/issues/6747) — *12 comments*  
   Wanted: allow extensions to mutate agent‑message markdown without changing the LLM payload. Essential for formula renderers and custom formatting. Now closed via PR #7231.

2. [#7153 [OPEN] `/scoped‑models` appears to do nothing for ~5 minutes](https://github.com/earendil-works/pi/issues/7153) — *6 comments*  
   Running `/scoped‑models` hangs the REPL while the model catalog refreshes. No loading indicator, no timeout – a clear UX regression for multi‑model users.

3. [#7161 [OPEN] anthropic‑messages never sends `x‑client‑request‑id`](https://github.com/earendil-works/pi/issues/7161) — *6 comments*  
   Proxies that use this header for session affinity cannot group Anthropic conversations. The reporter runs a round‑robin proxy with two Claude accounts; without the header, conversations land on random backends.

4. [#7047 [OPEN] Gemini 3.x tool‑call IDs stripped from function calls](https://github.com/earendil-works/pi/issues/7047) — *5 comments*  
   In multi‑turn tool conversations, the `id` field is dropped from both `functionCall` and `functionResponse` parts. Gemini 3 requires echoed IDs, so sessions break after one tool call.

5. [#7187 [CLOSED] Silent crash caused by inconsistent error handling](https://github.com/earendil-works/pi/issues/7187) — *4 comments*  
   A typo in a third‑party package manifest permanently killed all chat and scheduled sessions for a user. The crash happens before extensions run, so `pi -ne` doesn’t help.

6. [#7248 [OPEN] Ctrl+V text paste silently fails on Wayland](https://github.com/earendil-works/pi/issues/7248) — *4 comments*  
   `readClipboardText()` relies on X11‑only clipboard‑rs. On Wayland (KDE Plasma 6), copying from a Wayland app gives empty paste. A fix via CLI tools (`wl‑paste`) was merged in PR #7261.

7. [#7027 [OPEN] API‑key login can hang after saving credential when model catalog refresh stalls](https://github.com/earendil-works/pi/issues/7027) — *3 comments, 4 👍*  
   `/login openrouter` can stall even after the key is saved. The catalog refresh blocks the UI indefinitely, leaving users stuck with no escape. High community upvote count signals wide impact.

8. [#7332 [CLOSED] Streaming output becomes extremely slow as conversation context grows](https://github.com/earendil-works/pi/issues/7332) — *2 comments*  
   TUI streaming performance degrades over long sessions. No user action is needed – the terminal simply stops printing at a reasonable rate. Potential bottleneck in context‑size handling.

9. [#7283 [OPEN] Anthropic stream parser discards initial block](https://github.com/earendil-works/pi/issues/7283) — *2 comments*  
   The stream parser assumes `content_block.start` events are empty; when they contain text (e.g., thinking blocks), that content is lost. Affects all Anthropic‑compatible providers.

10. [#7337 [CLOSED] Tool‑result images are never resized](https://github.com/earendil-works/pi/issues/7337) — *1 comment*  
    Only `read` and CLI file‑processor calls `processImage`. Any other tool (extension tools, MCP servers, screenshots) stores full‑resolution images, which can exceed provider limits and brick sessions. Fixed by PR #7330.

## Key PR Progress
1. [#7348 [OPEN] feat(client): add runtime‑neutral session client](https://github.com/earendil-works/pi/pull/7348)  
   Introduces `@earendil-works/pi-client` – a transport‑agnostic client for remote sessions. Exposes typed requests, observable listener failures, and multi‑session handles.

2. [#7344 [CLOSED] feat(protocol): add remote session wire protocol](https://github.com/earendil-works/pi/pull/7344)  
   Defines validated remote‑session commands, events, snapshots, and errors with CBOR encoding and length‑prefixed framing. Waits for #7346 to merge schemas.

3. [#7343 [CLOSED] feat(agent): add harness shutdown lifecycle](https://github.com/earendil-works/pi/pull/7343)  
   Idempotent `AgentHarness.shutdown()` – rejects new work, aborts active turns, and prevents provider startup after shutdown. Critical for clean resource cleanup.

4. [#7261 [CLOSED] fix(coding-agent): read clipboard via wl‑paste on Wayland, xclip/xsel on X11](https://github.com/earendil-works/pi/pull/7261)  
   Fixes #7248. On Linux, prefers `wl‑paste` (Wayland) or `xclip`/`xsel` (X11), falling back to clipboard‑rs. Platform‑specific clipboard support is now robust.

5. [#7231 [CLOSED] Markdown api](https://github.com/earendil-works/pi/pull/7231)  
   Closes #6747. Adds an extension API to mutate agent message markdown representation without touching LLM content.

6. [#6987 [CLOSED] fix(tui): align grapheme widths with terminal cells](https://github.com/earendil-works/pi/pull/6987)  
   Major fix for Unicode rendering – addresses Devnagri (#6124) and other non‑Latin scripts that previously broke the harness UI.

7. [#7163 [OPEN] feat: search index sqlite](https://github.com/earendil-works/pi/pull/7163)  
   Adds `SessionRepo.search()` with FTS5 full‑text search for SQLite. For JSONL/memory it still does naive scanning; future work will optimize.

8. [#7286 [CLOSED] feat(ai): preserve structured metadata for Bedrock provider errors](https://github.com/earendil-works/pi/pull/7286)  
   Ensures AWS Bedrock V3‑protocol errors are properly parsed and surfaced, instead of showing raw stream objects.

9. [#7330 [CLOSED] fix(coding-agent): resize images returned by tools](https://github.com/earendil-works/pi/pull/7330)  
   Fixes #7337. Applies `processImage` (resize to 2000×2000, format conversion) to tool‑produced images – extension tools, MCP bridges, and browser screenshots now resized.

10. [#7061 [CLOSED] fix(openai‑completions): handle array content and missing finish_reason](https://github.com/earendil-works/pi/pull/7061)  
    Two bugs: Databricks models returning `choice.delta.content` as typed arrays (producing `[object Object]`), and missing `finish_reason`. Now correctly extracts text blocks and defaults complete.

## Feature Request Trends
- **Extension UI APIs** – Multiple requests for better control over agent message rendering (#6747), custom editor forwarding (#7333), user‑message submission outcomes (#7345), and OAuth page templates (#6930). The Markdown API (#7231) is a first step.
- **Provider‑specific enhancements** – OpenAI **background mode** (#7339), Anthropic `x‑client‑request‑id` (#7161), and Gemini 3 tool‑call ID preservation (#7047). Users want parity with first‑party clients.
- **Session & performance** – Context window selection (#5064), rounded work duration in session list (#7349), and full‑text session search (#7163). Moving toward a more browseable, efficient conversation history.
- **Internationalization** – Devnagri support (#6124) and Unicode grapheme width fixes (#6987) point to growing global adoption.
- **Package security** – A report flagged `@mjasnikovs/pi-task` for silently compressing thinking blocks (#7347), raising concerns about extension trust.

## Developer Pain Points
1. **Silent hangs & deadlocks** – `/scoped‑models` stalls for minutes (#7153), API‑key login hangs on catalog refresh (#7027), and concurrent `custom()` prompts deadlock (#7007). Lack of user‑visible progress or error messages frustrates developers.
2. **Inconsistent provider behaviour** – Anthropic missing request headers (#7161), Gemini tool‑call IDs stripped (#7047), Fireworks instant timeouts (#7315). Debugging across providers is time‑consuming.
3. **Platform‑specific breakage** – Wayland clipboard (#7248) and iTerm2 display corruption (#6784) show that terminal compatibility is patchy. The Wayland fix (#7261) is welcome but other platforms (macOS iTerm) still suffer.
4. **Reliability in production** – Silent crashes from package manifest typos (#7187) and JSON‑parse failures in RPC stdout (#7309) break entire installations. Error handling and validation need hardening.
5. **Tool‑image bloat** – Unresized tool images (#7337) can exceed provider limits and brick sessions – a subtle, hard‑to‑diagnose failure that affects any extension using images.
6. **Streaming performance degradation** – Over long conversations, TUI streaming becomes intolerably slow (#7332), suggesting a missing buffer or pagination strategy for large contexts.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest – 2026-07-31

## Today's Highlights
A nightly release brings CI and web-shell fixes. The community is deeply engaged with the Anthropic content converter, with multiple bugs reported and several PRs in flight to fix thinking‑signature pruning, tool_use ID sanitization, and message ordering. Major features advancing include integration of Goal v3 into the interactive TUI, packaging the Web Shell as a desktop app, and a new OpenAI Responses API content generator.

## Releases
- **[v0.21.1-nightly.20260730.1643a6c9a](https://github.com/QwenLM/qwen-code/releases/tag/v0.21.1-nightly.20260730.1643a6c9a)**  
  CI fix: default bash shell in container jobs (PR #7838). Web‑shell fix (details truncated in source).

## Hot Issues
1. **[Startup banner missing top lines on first paint (intermittent) – #8124](https://github.com/QwenLM/qwen-code/issues/8124)**  
   P2 bug affecting the TUI’s interactive experience; 9 comments. Likely a rendering race condition with pending provider updates.
2. **[Anthropic converter: stale thinking signatures not pruned after tool_use removal – #8162](https://github.com/QwenLM/qwen-code/issues/8162)**  
   Core converter bug that can cause malformed assistant turns after history trimming. 4 comments; being actively addressed in PR #8166.
3. **[Worktree settings.json writes to project root instead of worktree’s .qwen – #8138](https://github.com/QwenLM/qwen-code/issues/8138)**  
   Regression for git worktree users; settings changes affect the wrong workspace. 4 comments. PR #8152 attempts to fix.
4. **[Desktop app fails to connect to LMStudio on Windows – #8146](https://github.com/QwenLM/qwen-code/issues/8146)**  
   P2 integration bug blocking Windows users. Shows “doing something” but no API calls. 4 comments.
5. **[Windows standalone installer fails when Get-FileHash unavailable – #7118](https://github.com/QwenLM/qwen-code/issues/7118)**  
   Long‑standing P2 installer issue (2 👍). Falls back to npm method, but should handle missing PowerShell cmdlets gracefully.
6. **[Anthropic converter: tool_result blocks not guaranteed first in mixed‑content user messages – #8161](https://github.com/QwenLM/qwen-code/issues/8161)**  
   Causes incorrect message order when Gemini functionResponse is accompanied by text. 3 comments. PR #8165 fixes.
7. **[Anthropic converter: tool_use.id not sanitized to Anthropic’s required charset – #8160](https://github.com/QwenLM/qwen-code/issues/8160)**  
   IDs from other providers pass through unsanitized, potentially triggering API rejections. 3 comments.
8. **[Anthropic converter: cleanOrphanedToolCalls strips a trailing tool_use that has no next message – #8159](https://github.com/QwenLM/qwen-code/issues/8159)**  
   Loses the model’s last tool call if no user turn follows. 3 comments. Urgent for multi‑turn flows.
9. **[Provider warning sanitizer truncates messages with ports and leaks passwords containing `@` – #8136](https://github.com/QwenLM/qwen-code/issues/8136)**  
   Security bug in CLI status payloads; credentials can leak. 3 comments. P2.
10. **[Desktop client cannot reference correct files using `@` symbol (Chinese user) – #8123](https://github.com/QwenLM/qwen-code/issues/8123)**  
    File indexing bug in Windows desktop client. 3 comments. Suggests unmet need for robust file search.

## Key PR Progress
1. **[feat(cli): adopt Goal v3 in interactive TUI – #8005](https://github.com/QwenLM/qwen-code/pull/8005)**  
   Adds `/goal` commands, lifecycle cards, resume/branch recovery, and two‑lane input queue. Major UX enhancement for goal‑driven workflows.
2. **[feat(desktop): package Web Shell as a release‑ready desktop app – #8132](https://github.com/QwenLM/qwen-code/pull/8132)**  
   Turns Tauri proof‑of‑concept into a proper desktop shell around the existing Web Shell. Native lifecycle, workspace picker, update UX.
3. **[feat(core): add current PR Autofix watcher – #8121](https://github.com/QwenLM/qwen-code/pull/8121)**  
   Opt‑in `/autofix` watcher that reports PR CI/review state and can create fix sessions. Bridges development and code review.
4. **[fix(serve): isolate managed memory by selected workspace – #8056](https://github.com/QwenLM/qwen-code/pull/8056)**  
   Workspace‑qualified remember/forget/dream operations; opt‑in exact‑workspace storage. Prevents memory leaks across projects.
5. **[feat(memory): configure background agent turn limits – #8171](https://github.com/QwenLM/qwen-code/pull/8171)**  
   Adds `memory.agentMaxTurns` setting for dream and auto‑skill review agents (default 8). Addresses feature request #8168.
6. **[fix: make the test suite portable on Windows – #8050](https://github.com/QwenLM/qwen-code/pull/8050)**  
   Reuses self‑hosted Windows CI with stable locale and temp directories. Critical for Windows developer confidence.
7. **[feat(core): add OpenAI Responses API content generator – #8169](https://github.com/QwenLM/qwen-code/pull/8169)**  
   New content generator for the OpenAI Responses API endpoint. Expands provider compatibility beyond chat completions.
8. **[fix(mcp): prevent race condition in OAuth token refresh – #8170](https://github.com/QwenLM/qwen-code/pull/8170)**  
   High‑severity TOCTOU race in OAuth flow. Ensures only one refresh per batch of concurrent requests; avoids corrupted tokens.
9. **[fix(anthropic): prune stale thinking signatures after a sibling tool_use is removed – #8166](https://github.com/QwenLM/qwen-code/pull/8166)**  
   Two‑patch fix for #8162: cascading cleanup and multi‑turn pruning. Downstream‑tested fork port.
10. **[fix(anthropic): move tool_result blocks first in mixed‑content user messages – #8165](https://github.com/QwenLM/qwen-code/pull/8165)**  
    Guarantees `tool_result` blocks precede other content in user‑role messages. Fixes #8161.

## Feature Request Trends
- **Memory agent configurability** – request for `memory.agentMaxTurns` setting (now implemented in PR #8171). Users want fine‑grained control over background agent budget.
- **Broader provider integration** – issues with LMStudio and interest in OpenAI Responses API (PR #8169) signal demand for more plug‑and‑play model backends.
- **Improved file referencing & search** – the desktop `@` symbol not finding files (#8123) highlights need for smarter workspace indexing, especially on Windows.
- **Agent team async messaging** – #8172 reveals that teammate messages are not delivered until idle, which blocks fluent multi‑agent collaboration.

## Developer Pain Points
- **Windows compatibility** – Installer fails (Get‑FileHash, #7118), desktop app broken with LMStudio (#8146), and test suite portability (#8050) remain recurring frustrations.
- **Anthropic converter instability** – Four separate bugs surfaced in one day (stale thinking, ID sanitization, orphan tool_use, message ordering). The converter needs a holistic review.
- **Git worktree isolation** – Settings and context file resolution still target the project root instead of the worktree directory (#8138, PR #8152).
- **UI rendering flakiness** – Startup banner missing top lines (#8124) and statusline unselectable in virtualized history (#8131) degrade the interactive experience.
- **Credential sanitization** – Provider warning sanitizer both truncates legitimate port‑containing messages and leaks passwords containing `@` (#8136), exposing security gaps.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-07-31

> **Data source:** github.com/Hmbown/DeepSeek-TUI  
> *Note: The repository has been renamed to `CodeWhale`; the legacy `deepseek-tui` npm package is deprecated as of v0.9.2.*

---

## Today's Highlights

The community is deep into **v0.9.3 refactoring**, with several umbrella issues consolidating runtime ownership, tool system reconciliation, and context diet work. A new **compilation-time discussion** (#4991) has ignited conversation about the TUI crate monolith, while the finalized **v0.9.2 release** brings permission, sub-agent, and credential UX fixes. A **desktop app** feature request (#4986) signals growing demand beyond the terminal.

---

## Releases

### v0.9.2 (finalized 2026-07-30)
- Shipped as the final release of the `codewhale` npm package; legacy `deepseek-tui` v0.8.x is deprecated and receives no further updates.
- Handoff fixes across permission truth, Fleet setup/persistence, reasoning inspection, compaction errors, sub‑agent supervision/steering, sandbox truth, provider credential UX, and ambient silhouettes.
- See [PR #4982](https://github.com/Hmbown/CodeWhale/pull/4982) for full changelog.

---

## Hot Issues (10 Noteworthy)

1. **[#4991] Compilation times and the TUI crate monolith**  
   *[discussion]* Community member @aboimpinto reports excessive wait times while working on command refactoring. Highlights that 87% of Rust code lives in `codewhale-tui`. Sparks discussion on crate splitting.  
   [Open Issue](https://github.com/Hmbown/CodeWhale/issues/4991)

2. **[#2870] EPIC: staged command-boundary refactor**  
   Tracks layered mergeable work for Hmbown/CodeWhale#2791. Reference PR #2851. Critical for clean user‑command dispatch precedence.  
   [Open Issue](https://github.com/Hmbown/CodeWhale/issues/2870)

3. **[#2369] CodeWhale Config Paths Fragmented Across OS and Cygwin**  
   Legacy migration bug leads to divergent secrets paths on Windows/Cygwin. 7 comments – still unresolved in v0.9.2.  
   [Open Issue](https://github.com/Hmbown/CodeWhale/issues/2369)

4. **[#4022] CLI/TUI parity for subagent and runtime control surfaces**  
   Ensures subagent status, expand/collapse, and cancellation aren’t trapped in TUI only. Needed for future cloud/remote workbench.  
   [Open Issue](https://github.com/Hmbown/CodeWhale/issues/4022)

5. **[#3306] v0.9.3 Refactor: converge runtime ownership, delete duplication**  
   Umbrella for consolidating parallel runtime, tool, config, session, hook, and control paths into a single executable.  
   [Open Issue](https://github.com/Hmbown/CodeWhale/issues/3306)

6. **[#4949] Discussion: Chinese translation of “Constitution”**  
   Community debate between 宪法 vs 协作准则, with political sensitivity concerns. Author @SparkofSpike invites native speakers to vote.  
   [Open Issue](https://github.com/Hmbown/CodeWhale/issues/4949)

7. **[#4906] Show, don't tell: record a real Codewhale session**  
   README and website lack a demo GIF. Fundamental for first-time user perception. Author @Hmbown calls it a “motion-heavy product” needing visual proof.  
   [Open Issue](https://github.com/Hmbown/CodeWhale/issues/4906)

8. **[#4986] feat(desktop): first-class desktop app**  
   User @JoeKerF requests a complete desktop‑product experience (similar to Codex Desktop) instead of managing terminal, working directories, and processes manually.  
   [Open Issue](https://github.com/Hmbown/CodeWhale/issues/4986)

9. **[#4978] Anthropic API error: `'type' must be in ["enabled", "disabled", "auto"]`**  
   Recurring HTTP 400 when using `providers.openmodel` compatible with Anthropic Messages API. Affects reliability for multi‑provider users.  
   [Open Issue](https://github.com/Hmbown/CodeWhale/issues/4978)

10. **[#4704] Context diet: minimize every model-facing prompt, schema, and payload**  
    Part of v0.9.3 performance push. Aims to reduce token waste without sacrificing task quality. Parent to three sub‑issues.  
    [Open Issue](https://github.com/Hmbown/CodeWhale/issues/4704)

---

## Key PR Progress (10 Important)

1. **[#4992] Layer 5.2: User command dispatch precedence, shadowing, and error semantics**  
   Adds Gherkin acceptance tests for built‑in vs. user‑command shadowing. @aboimpinto – important for command‑boundary refactor.  
   [Open PR](https://github.com/Hmbown/CodeWhale/pull/4992)

2. **[#4979] fix(tui): detach foreground shell before steering**  
   Allows user to type Enter while `sleep`/`cargo build` is blocking – moves foreground process to `/jobs` before steering. @nightt5879.  
   [Closed PR](https://github.com/Hmbown/CodeWhale/pull/4979)

3. **[#4981] feat(tui): LaTeX environments, text, and command support**  
   Extends math rendering with environment blocks, inline commands, accents. @SparkofSpike – improves academic/technical output readability.  
   [Open PR](https://github.com/Hmbown/CodeWhale/pull/4981)

4. **[#4984] fix runtime config persistence and workspace task scoping**  
   Rebase GUI‑facing TUI runtime API; adds `workspace` filter to `GET /v1/tasks`. @gaord – en route to desktop integration.  
   [Closed PR](https://github.com/Hmbown/CodeWhale/pull/4984)

5. **[#4985] feat(runtime-api): scope task listing by workspace**  
   Follow‑up to #4984; includes regression tests. @gaord – enables GUI clients to filter tasks per project.  
   [Open PR](https://github.com/Hmbown/CodeWhale/pull/4985)

6. **[#4982] release: finalize Codewhale v0.9.2**  
   Bundles all handoff fixes; closes the v0.9.2 cycle. @Hmbown.  
   [Closed PR](https://github.com/Hmbown/CodeWhale/pull/4982)

7. **[#4977] fix(tui): let AltGr-typed "/" reach the composer**  
   Fixes #4723 – on Brazilian ABNT2 layout, `AltGr+Q` (which is `/`) opened help overlay. @yyyCode.  
   [Open PR](https://github.com/Hmbown/CodeWhale/pull/4977)

8. **[#4990] fix(devcontainer): support Windows development**  
   Uses dedicated image with Rust toolchain and named volumes to avoid invalid Windows HOME expansion. @pingg02 – lowers onboarding friction.  
   [Open PR](https://github.com/Hmbown/CodeWhale/pull/4990)

9. **[#4942] fix(tools): preserve CRLF edits**  
   Ensures `edit_file` handles Windows line endings correctly; maps spans back to original CRLF bytes. @nightt5879.  
   [Closed PR](https://github.com/Hmbown/CodeWhale/pull/4942)

10. **[#4680] fix(tui): register debt compatibility aliases**  
    Adds `/slop` and `/canzha` as `/debt` aliases, removing pre‑registry special case. @nightt5879.  
    [Closed PR](https://github.com/Hmbown/CodeWhale/pull/4680)

---

## Feature Request Trends

- **Desktop / GUI application** – Issue #4986 is the most direct request, but the runtime‑API PRs (#4984, #4985) and the CLI/TUI parity work (#4022) all point toward a richer desktop experience beyond the terminal.
- **Visual onboarding & demos** – Issue #4906 calls for a README GIF and website screencast; the project now prioritises “show, don’t tell”.
- **LaTeX rendering** – PR #4981 extends math support, appealing to academic and scientific users.
- **i18n and localization** – Debate over Chinese translation for “Constitution” (#4949) and recent locale fixes (PR #4856) show growing international community.
- **Subagent steering** – Issue #4989 (closed) and #4022 envision explicit, durable sub‑agent control – a core feature for multi‑agent workflows.

---

## Developer Pain Points

- **Compilation times** – Issue #4991: the monolith `codewhale-tui` crate (~771k Rust lines) causes slow iteration. Several refactor epics (#3306, #3948, #3957, #3950) directly address this by splitting runtime, tool, and main‑file responsibilities.
- **Tool system duplication** – Two separate tool registries (`codewhale-tools::ToolRegistry` vs `codewhale-tui::ToolRegistry`) create maintenance overhead (#4174).
- **Config fragmentation** – Paths diverge across OS and Cygwin, plus a silent migration bug (#2369), causing credential loss and confusion.
- **Performance of skill discovery** – Recursive full‑root rescans on every prompt construction (#3921) – rarely‑changing data wasted repeatedly.
- **Context bloat** – Repeated system‑prompt layers, duplicate environment/skill facts, and oversized payloads (#4704, #4709, #4710) – both token‑waste and model‑fragility concerns.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*