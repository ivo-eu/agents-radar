# AI CLI Tools Community Digest 2026-06-30

> Generated: 2026-06-30 10:45 UTC | Tools covered: 9

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
**Analysis Date:** 2026-06-30

---

## 1. Ecosystem Overview

The AI CLI tools landscape in mid-2026 reflects a maturing ecosystem where reliability, cross-platform parity, and enterprise governance are displacing raw feature velocity as the primary community concerns. All six major tools—Claude Code, OpenAI Codex, Gemini CLI, GitHub Copilot CLI, OpenCode, and Pi—show signs of convergent evolution toward shared architectural patterns: MCP (Model Context Protocol) integration, daemon-backed session management, and agent resilience improvements. However, the community dynamics differ sharply: Claude Code and OpenCode exhibit the highest engagement volume, while Kimi Code and Qwen Code show quieter but deliberate progress on UX fundamentals. A notable bifurcation is emerging between "generalist" coding agents (Claude, Codex, Gemini) and "task-specific" or "opinionated" tools (Copilot CLI, Pi, DeepSeek TUI), with the latter groups emphasizing terminal ergonomics and permission models over raw model capability.

---

## 2. Activity Comparison (2026-06-30)

| Tool | Issues (Hot/Notable) | Key PRs (Last 24h) | Release Status | Community Engagement (Top Issue 👍) |
|---|---|---|---|---|
| **Claude Code** | 10 hot issues | 10 PRs (AZERDSQ131 series + infrastructure) | v2.1.196 (org defaults, readable sessions) | 401 👍 (multi-account switching) |
| **OpenAI Codex** | 10 hot issues | 10 PRs (queued messages, MCP file transfer) | rust-v0.143.0-alpha.31 | 71 👍 (markdown export) |
| **Gemini CLI** | 10 hot issues | 10 PRs (MCP elicitation, recursive reasoning fix) | v0.51.0-nightly | 8 👍 (generalist agent hangs) |
| **GitHub Copilot CLI** | 10 hot issues | 1 PR (issue classification automation) | v1.0.66-2 (multi-plugin coexistence) | 17 👍 (project-scoped plugins) |
| **Kimi Code CLI** | 0 new issues | 2 PRs (bright-blue highlight, `--prompt-interactive`) | None | N/A (quiet day) |
| **OpenCode** | 10 hot issues | 10 PRs (temperature fix, Bedrock thinking, OOM fix) | None | 90 👍 (model fallback) |
| **Pi (pi-mono)** | 10 hot issues | 10 PRs (streaming stability, error transparency) | None | 42 comments (scroll-jump bug) |
| **Qwen Code** | 5 issues (low volume) | 10 PRs (daemon archive, OOM fix, inline model override) | v0.19.3-nightly | Low engagement (quiet day) |
| **DeepSeek TUI** | 10 hot issues | 10 PRs (v0.8.66 blocker series) | None (pre-release fixes) | ~24 comments (cache hit rate) |

**Key Observations:**
- **Claude Code** and **OpenCode** dominate community engagement with high-vote features (multi-account, model fallback).
- **Gemini CLI** and **DeepSeek TUI** show intense PR activity despite lower issue upvote counts—indicating rapid internal iteration.
- **Kimi Code** and **Qwen Code** had the quietest day, suggesting either release cycle lulls or smaller user bases.
- **Copilot CLI** has the most mature release version (v1.0.x), while **Gemini CLI** and **Qwen Code** remain in nightly/alpha territory.

---

## 3. Shared Feature Directions

Several requirements recur across tool communities, independent of vendor allegiance:

### 3.1 Session & State Management
| Need | Affected Tools | Specific Ask |
|---|---|---|
| Session archive/restore (without deletion) | Qwen Code (#6057), Claude Code (#72461), Pi (#4877) | Lightweight offloading of old sessions |
| Session resume after interruption | Claude Code (#16607), Gemini CLI (implicit crashes), DeepSeek TUI (#3821) | Continue agent work after crash/timeout |
| Global session search | Claude Code (#71611), OpenCode (implicit via `opencode stats`) | Cross-project conversation retrieval |

### 3.2 MCP & Plugin Ecosystem Maturity
| Need | Affected Tools | Specific Ask |
|---|---|---|
| OAuth token refresh for MCP | OpenCode (#34582), DeepSeek TUI (#3819), Copilot CLI (#3973) | Short-lived tokens break enterprise workflows |
| Duplicate MCP server detection | DeepSeek TUI (#3461), Copilot CLI (#3893) | Name collision and multiple instances |
| Plugin sandboxing & security | Copilot CLI (#3874, preToolUse denials ignored), Gemini CLI (#28215, file-write scope hardening) | Preventing sandbox escape |

### 3.3 Model Flexibility & Fallback
| Need | Affected Tools | Specific Ask |
|---|---|---|
| Model failover when at capacity | OpenCode (#7602), Claude Code (rate-limit complaints), Codex (#29760) | Automatic retry with alternative model |
| Custom provider support fix | OpenCode (#5674), Copilot CLI (#3954), Pi (#5763) | Options silently dropped for non-standard APIs |
| Service tier control | Codex (#2916), Claude Code (Ultra plan disable request #71613) | Explicit cost/latency trade-off |

### 3.4 Agent Reliability & Resilience
| Need | Affected Tools | Specific Ask |
|---|---|---|
| Agent stuck in loops (false "working") | Gemini CLI (#22323, #21409), Pi (#4338), Claude Code (#64108) | Subagents claiming success after turn limit |
| Tool call serialization failures | Claude Code (#64108), OpenCode (MCP tool mislabeling) | Raw XML leaks instead of execution |
| Think/chain-of-thought display | Claude Code (#30958), OpenCode (#33630 Bedrock thinking), Gemini CLI (#27971) | Extended reasoning audit trails |

### 5. Terminal & IDE Parity
| Need | Affected Tools | Specific Ask |
|---|---|---|
| `/btw` (quick side-question) in IDE | Claude Code (#37323), OpenCode (implicit) | Feature parity between terminal and VS Code |
| Markdown export | Codex (#2880), OpenCode (#31453), Pi (implicit) | Native conversation documentation |
| Disable alt-screen views | Copilot CLI (#1799), Claude Code (opt-out culture) | Users want traditional terminal output |

---

## 4. Differentiation Analysis

### 4.1 Focus Areas by Tool

| Tool | Primary Focus | Target User | Technical Signature |
|---|---|---|---|
| **Claude Code** | Agentic coding with thinking display | Individual developers, power users | Extended thinking (Opus), session management UX |
| **OpenAI Codex** | Multi-model agent platform | Pro/Enterprise | Rust tooling, queued message persistence, sandboxing |
| **Gemini CLI** | Sub-agent orchestration | Developers using Gemini models | MCP elicitation, recursive reasoning limits, Caretaker Agent |
| **GitHub Copilot CLI** | Seamless GitHub integration | GitHub ecosystem users | Multi-plugin coexistence, `tgrep` indexer, v1.0 maturity |
| **Kimi Code CLI** | Minimalist shell discussion | Rapid prototyping | `--prompt-interactive`, bright-blue highlights (UX polish) |
| **OpenCode** | Open-source extensibility | Self-hosters, tinkerers | Provider fallback, YOLO mode, MCP resource autocomplete |
| **Pi** | Stable streaming & error resilience | Reliability-focused users | Provider error body transparency, redo support, auto-compaction fix |
| **Qwen Code** | Daemon-as-platform | Multi-client (CLI + QQ bot) | Sessionless workspace remember, channel workers, emoji cleanup |
| **DeepSeek TUI** | High fanout & permission granularity | Power users, MMO-style interaction | Wildcard tool disallow, Hotbar, YOLO authority persistence |

### 4.2 Technical Approach Differences

- **Binary vs. Plugin Architecture**: Copilot CLI (v1.0) and Pi offer stable binary releases with plugin hooks, while Claude Code and OpenCode rely more on community plugins (OpenCode's `@ai-sdk` ecosystem) and MCP. DeepSeek TUI leans into a semi-MMO interface with persistent permission rules—unique in the space.

- **Model Governance**: Claude Code now supports org-default models (enterprise admin control), while OpenCode enables per-agent provider fallback chains. Codex lacks this but offers `responses-lite` headers for custom pipelines. Qwen Code's inline model override (`/model <id> <prompt>`) is the most lightweight approach.

- **Session Architecture**: Qwen Code is pioneering daemon-managed sessions with channel workers and sessionless memory tasks—positioning its daemon as an extensible backend beyond CLI. Pi focuses on session compaction and error recovery. Claude Code and Codex treat sessions as first-class conversation objects with resume/archive.

- **Terminal UI Philosophy**: Copilot CLI and Kimi Code favor minimal disruption (alt-screen controversy), while DeepSeek TUI embraces maximal information density (Hotbar, expansion views). Pi and Qwen Code are converging on right-sizing (padding control, emoji removal).

### 4.3 Enterprise Readiness

| Dimension | Claude Code | Codex | Gemini CLI | Copilot CLI | OpenCode | Pi | DeepSeek TUI | Qwen Code |
|---|---|---|---|---|---|---|---|---|
| Org defaults | ✅ (v2.1.196) | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ | ❌ |
| Plugin sandboxing | Partial | Sandbox.exe issues | File-write hardening | preToolUse broken | MCP OAuth issues | Error transparency | Wildcard disallow | Allow/deny MCP |
| Rate limit handling | ❌ (429 reports) | ❌ (capacity errors) | ❌ | ❌ | #7602 (requested) | Retryable error fix | ❌ | ❌ |
| Compliance (audit trails) | Thinking summaries broken | ❌ | Thought leak fix | ❌ | ❌ | ❌ | ❌ | ❌ |

**Enterprise gap**: No tool has comprehensive rate-limit handling, and only Claude Code is actively building admin controls. MCP OAuth token refresh is a cross-cutting failure.

---

## 5. Community Momentum & Maturity

### 5.1 High-Velocity Communities (Rapid Iteration)

- **DeepSeek TUI**: Most aggressive PR volume for a single release cycle (v0.8.66 blocker series: 6+ fixes in 24h). Community reports high severity bugs (session corruption, MCP OAuth) but team responds quickly. Still pre-v1.0, but shipping fast.

- **OpenCode**: Consistently high issue upvote counts (model fallback at 90👍, YOLO mode at 89👍) with corresponding PR merges (Bedrock thinking fix, temperature fix). Strong open-source contributor ecosystem.

- **Claude Code**: Largest absolute community (401👍 on top feature). However, duplicate bug reports (SSL cert expiry, empty thinking summaries) suggest regression management challenges. Release cadence is healthy (v2.1.196).

### 5.2 Stable, Maturing Platforms

- **GitHub Copilot CLI**: Highest version number (v1.0.66-2). Slower PR cadence (1 PR today) but consistent release stream. Community demands are more about UX polish (trackpad, alt-screen) than core features.

- **Pi**: Balanced approach—10 PRs merged with targeted fixes (streaming stability, error handling). Community engagement is moderate (42 comments max), suggesting a satisfied but smaller user base.

### 5.3 Emerging/Quiet Tools

- **Kimi Code**: Very quiet day (0 issues, 2 PRs). May reflect a smaller user community or a period of internal refactoring. The two PRs (shell highlighting, `--prompt-interactive`) are quality-of-life but not transformative.

- **Qwen Code**: Low issue volume but substantial PR output (10) around daemon infrastructure. This suggests a **developer-first** strategy—building platform capabilities before scaling community.

- **Gemini CLI**: Moderate issue activity (10 hot) but heavy PR output (10). Mix of critical bugs (CLI crash, subagent hangs) and forward-looking features (MCP elicitation, Caretaker Agent). Still in nightly release cycle (v0.51.0).

### 5.4 Community Health Signals

| Signal | Claude Code | Codex | Gemini CLI | Copilot CLI | Kimi Code | OpenCode | Pi | DeepSeek TUI | Qwen Code |
|---|---|---|---|---|---|---|---|---|---|
| Top issue 👍 count | **401** | 71 | 8 | 17 | N/A | **90** | N/A (42 comments) | ~24 comments | Low |
| Issues updated today | 10 | 10 | 10 | 10 | 0 | 10 | 10 | 10 | 5 |
| PRs merged today | 10 | 10 | 10 | 1 | 2 | 10 | 10 | 10 | 10 |
| Release last 24h | ✅ | ✅ | ✅ | ✅ | ❌ | ❌ | ❌ | ❌ | ✅ |

**Interpretation**: Claude Code and Codex have the largest user bases (high upvote counts). DeepSeek TUI, Gemini CLI, and Qwen Code are iterating fastest per developer. Kimi Code appears dormant but may be building.

---

## 6. Trend Signals

### 6.1 Session Lifecycle Management Becomes a Product Category

Every tool is investing in session-related features: Claude Code (deep compaction #72461, readable names), Codex (queued message persistence), Gemini CLI (thread settings synchronization), OpenCode (session persistence), Pi (collision checks), Qwen Code (archive/restore, sessionless memory). This signals that **session management is no longer an afterthought but a core UX differentiator**. Expect dedicated session browsers, cross-device sync, and session-based compliance auditing in 2027.

### 6.2 MCP Convergence, but Standardization Lag

MCP is the dominant plugin protocol across all tools, but **no single implementation handles OAuth token refresh, duplicate detection, or sandboxing correctly**. The community is filing near-identical bugs across tools (#34582 OpenCode, #3819 DeepSeek, #3874 Copilot CLI). This is a gap that opens the door for a cross-tool tooling abstraction layer—or a faster protocol evolution from MCP maintainers.

### 6.3 Agent Reliability is the #1 Developer Pain Point

"Agent stuck working," "false success on turn limit," "hangs on deferral," "crashes on confirmation menu"—these are not one-off bugs but **systemic failures in agent state machines**. The frequency of these reports (Gemini CLI #21409, Pi #4338, DeepSeek #3800, Claude Code #16607) suggests the current agent-as-REPL pattern is hitting architectural limits. **We may see a shift from prompt-based agents to state-machine or workflow-based agents** in 2027, where task completion is explicit and interruptible.

### 6.4 Terminal UX is Undergoing a Quiet Revolution

Kimi Code's bright-blue user input highlights, Pi's padding controls, Copilot CLI's alt-screen controversy, DeepSeek's Hotbar, Qwen Code's emoji removal—all point to **a generation of developers who want terminal AI tools to feel as refined as modern IDEs, not like experimental demos**. The scroll-jump bug in Pi (#5825, 42 comments) shows that even small UX regressions are dealbreakers. Expect **customizable TUI themes, configurable input/output separation, and accessibility improvements** (color blindness, screen reader support) to become table stakes.

### 6.5 Model Flexibility Becomes a Feature, Not a Given

Users are increasingly unhappy with hardcoded model selections (Copilot CLI #3954, OpenCode #34570, DeepSeek #743) and opaque fallback behavior. **Model failover, service tier control, and custom provider support are now "must-have" features**, not nice-to-haves. OpenCode's #7602 (model fallback at 90👍) and Codex's #2916 (service tier at 50👍) both signal that users want to **optimize for cost, latency, and availability independently**.

### 6.6 Security Posture is Improving, But Unevenly

- **Good**: Gemini CLI's file-write scope hardening (#28215), DeepSeek's wildcard tool disallow (#3824), OpenCode's MCP resource sandboxing.
- **Bad**: Copilot CLI's `preToolUse` hook denial is ignored (#3874), Claude Code's SSL cert regression (#71663), DeepSeek's duplicate MCP server processes (#3461).
- **Missing**: No tool offers **granular permission rules with rollback audit logs**—the closest is DeepSeek's #1186 (typed persistent permissions, still open).

Enterprise users should watch this space; a unified security model across tools could become a buying decision.

### 6.7 Signals for Developers

| Signal | Implication | Action |
|---|---|---|
| "Always-on" archiving trend | Sessions are becoming assets, not ephemeral | Plan for session storage, search, and export |
| MCP OAuth token refresh broken everywhere | Enterprise MCP rollout will hit auth walls | Prefer server-side OAuth with refresh tokens |
| Agent loop bugs persist across 5+ tools | Agent frameworks need better termination guarantees | Budget for agent timeout & interrupt recovery |
| Model capacity errors dominate Codex & Claude | Multi-model fallback is table stakes now | Design for "model routing" in your toolchain |
| Terminal UX maturation (padding, scroll, themes) | CLI tools are competing with IDE polish | Invest in TUI design as product differentiator |

---

## 7. Summary

The AI CLI tools ecosystem in mid-2026 is **converging on session management, MCP standardization, and agent reliability** while **diverging on terminal UX philosophy and enterprise governance**. Claude Code and OpenCode lead in community engagement and feature breadth, while DeepSeek TUI and Gemini CLI iterate fastest. GitHub Copilot CLI is the most mature release, but its slower PR cadence suggests a plateau. Kimi Code and Qwen Code are quieter but making deliberate architectural bets (daemon platform, minimal shell UX). **The tools that succeed will be those that balance agent autonomy with user control, and model flexibility with reliable execution**—the technical debt in current agent state machines is the ecosystem's most critical risk.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
*Data as of June 30, 2026 | Source: github.com/anthropics/skills*

---

## 1. Top Skills Ranking

The most-discussed pull requests reveal a community intensely focused on **reliability, toolchain quality, and cross-platform compatibility** rather than flashy new features.

### #1: **Skill-Creator Eval Fix** (PR #1298)
- **Skill:** Fixes `run_eval.py` which always reports 0% recall, breaking the description-optimization loop
- **Discussion Highlights:** Ties together 10+ independent reproductions of the same bug (#556). Addresses Windows stream reading, trigger detection, and parallel worker issues
- **Status:** Open (created June 10, last updated June 23)
- **Link:** [PR #1298](https://github.com/anthropics/skills/pull/1298)

### #2: **Document Typography Skill** (PR #514)
- **Skill:** Prevents orphan words/widows, numbering misalignment in AI-generated documents
- **Discussion Highlights:** Addresses universal document quality issues affecting every Claude-generated output. Low complexity, high impact
- **Status:** Open (created March 4, last updated March 13)
- **Link:** [PR #514](https://github.com/anthropics/skills/pull/514)

### #3: **ODT Skill — OpenDocument Support** (PR #486)
- **Skill:** Create, fill, read, and convert OpenDocument Format files (.odt, .ods)
- **Discussion Highlights:** Enables LibreOffice/ISO standard document workflows. Triggers on "ODT/ODS/ODF/OpenDocument" mentions
- **Status:** Open (created March 1, last updated April 14)
- **Link:** [PR #486](https://github.com/anthropics/skills/pull/486)

### #4: **PDF Case-Sensitivity Fix** (PR #538)
- **Skill:** Corrects 8 case-sensitive file references in `skills/pdf/SKILL.md` (REFERENCE.md → reference.md, FORMS.md → forms.md)
- **Discussion Highlights:** Breaks on case-sensitive filesystems (Linux, macOS). Simple but critical fix for cross-platform PDF skill reliability
- **Status:** Open (created March 6, last updated April 29)
- **Link:** [PR #538](https://github.com/anthropics/skills/pull/538)

### #5: **DOCX Tracked Changes Fix** (PR #541)
- **Skill:** Prevents document corruption when adding tracked changes to documents with existing bookmarks
- **Discussion Highlights:** Root cause was `w:id` collision in OOXML shared ID space. Hardcoded low IDs caused real-world corruption
- **Status:** Open (created March 6, last updated April 16)
- **Link:** [PR #541](https://github.com/anthropics/skills/pull/541)

### #6: **Self-Audit Skill** (PR #1367)
- **Skill:** Four-dimension reasoning quality gate before delivery — completeness, consistency, grounding, helpfulness
- **Discussion Highlights:** Universal skill (any project, any stack, any model). Replaces #1361. Very recent (June 28), rapidly gaining attention
- **Status:** Open (last updated June 29)
- **Link:** [PR #1367](https://github.com/anthropics/skills/pull/1367)

### #7: **Codebase Inventory Audit** (PR #147)
- **Skill:** Systematic 10-step workflow for identifying orphaned code, unused files, documentation gaps, and infrastructure bloat
- **Discussion Highlights:** Produces a CODEBASE-STATUS.md single source of truth. Strong interest in codebase hygiene automation
- **Status:** Open (created December 16, last updated February 4)
- **Link:** [PR #147](https://github.com/anthropics/skills/pull/147)

### #8: **Testing Patterns Skill** (PR #723)
- **Skill:** Comprehensive coverage of testing philosophy (Trophy model), unit testing, React component testing, E2E, accessibility testing
- **Discussion Highlights:** Covers "what to test vs. what NOT to test" — addresses a core developer pain point
- **Status:** Open (created March 22, last updated April 21)
- **Link:** [PR #723](https://github.com/anthropics/skills/pull/723)

---

## 2. Community Demand Trends

From the most-commented Issues, five clear demand vectors emerge:

### 🛡️ **Security & Trust Boundaries (Issue #492 — 32 comments)**
The top issue by a wide margin. Community members are concerned that community skills distributed under the `anthropic/` namespace enable **trust boundary abuse** — users grant elevated permissions thinking skills are official Anthropic releases. This is the single most-discussed concern across the entire repository.

📎 [Issue #492](https://github.com/anthropics/skills/issues/492)

### 🏢 **Org-Wide Skill Sharing (Issue #228 — 14 comments)**
Strong demand for **enterprise skill distribution**. Current workflow requires manual .skill file sharing via Slack/Teams and manual upload. Community wants direct sharing links or a shared skill library.

📎 [Issue #228](https://github.com/anthropics/skills/issues/228)

### 🐛 **Eval Tooling Reliability (Issue #556 — 12 comments)**
The `run_eval.py` 0% recall bug is a **blocker for skill creators**. Multiple independent reproductions confirm the description-optimization loop optimizes against noise. This is the most active debugging thread.

📎 [Issue #556](https://github.com/anthropics/skills/issues/556)

### 📦 **Skill Persistence & Stability (Issue #62 — 10 comments)**
Users report **skills disappearing** after file renames or Claude updates. Community wants robust skill persistence guarantees.

📎 [Issue #62](https://github.com/anthropics/skills/issues/62)

### 🧰 **Skill-Creator Best Practices (Issue #202 — 8 comments)**
The skill-creator skill itself needs rework — currently reads as **developer documentation** rather than operational instructions. Verbose, educational tone undermines token efficiency. Name also violates guidelines.

📎 [Issue #202](https://github.com/anthropics/skills/issues/202)

---

## 3. High-Potential Pending Skills

These actively-discussed PRs are likely to land soon:

| Skill | PR | Last Activity | Key Driver |
|-------|----|---------------|------------|
| **Skill-Creator Eval Fix** | [#1298](https://github.com/anthropics/skills/pull/1298) | June 23 | Fixes entire eval loop — unblocks all skill optimization |
| **Self-Audit Skill** | [#1367](https://github.com/anthropics/skills/pull/1367) | June 29 | Universal quality gate, very recent |
| **Run Eval Trigger Detection Fix** | [#1323](https://github.com/anthropics/skills/pull/1323) | June 25 | Misses real skill names, bails on non-Skill tools |
| **Testing Patterns** | [#723](https://github.com/anthropics/skills/pull/723) | April 21 | Comprehensive testing methodology |
| **Compact Memory** | [#1329](https://github.com/anthropics/skills/issues/1329) | June 26 | Symbolic notation for compact agent state |

The **meta-skill theme** is dominant: three of the top five pending items are fixes to the skill-creation toolchain itself, not new domain skills.

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is for foundational reliability and trust infrastructure — fixing the skill-creation toolchain, establishing namespace security boundaries, and enabling enterprise-grade sharing — over adding new functional skills.**

---

# Claude Code Community Digest — 2026-06-30

## Today's Highlights

Anthropic shipped **v2.1.196** with org-default model support and more readable session names, addressing two long-standing enterprise and UX requests. Meanwhile, the community saw a surge of duplicate bug reports on June 26, hinting at a broader regression wave (SSL cert expiry, empty thinking summaries, rate limiting). The top-voted feature request—multi-account switching in the mobile app—continues to dominate discussion with 401 👍 and 111 comments.

## Releases

**v2.1.196** ([changelog](https://github.com/anthropics/claude-code/releases/tag/v2.1.196))
- Admins can now set **organization default models** via the org console; users see “Org default” (or “Role default”) in `/model` until they pick a personal override.
- Session names now have **readable default names** at creation, making them easier to identify in the resume picker and sidebar.

No version bumps for the VS Code extension or Desktop app were reported in the last 24h.

---

## Hot Issues

Selected for community impact, signal-to-noise, and developer relevance.

### 1. Multi-account switching in Claude Mobile ([#36151](https://github.com/anthropics/claude-code/issues/36151))
401 👍, 111 comments, open since March. The most-wanted feature: switching between personal and work accounts on mobile without shared emails. Still unassigned.

### 2. `/btw` command missing from VS Code extension ([#37323](https://github.com/anthropics/claude-code/issues/37323))
119 👍, 29 comments. Users want IDE parity for quick side-questions that don’t persist in the main conversation. High demand for `area:ide`.

### 3. Tool calls emitted as literal text instead of executing ([#64108](https://github.com/anthropics/claude-code/issues/64108))
19 👍, 12 comments, open. A serious agent reliability bug: on Opus/CLI with large context, tool calls appear as raw `<invoke>` XML in the transcript instead of executing. Stray `court` token observed.

### 4. Empty thinking summaries in transcripts/TUI ([#30958](https://github.com/anthropics/claude-code/issues/30958))
15 👍, 12 comments, regression from v2.1.69. Undocumented behavior change that strips thinking summaries; affects users relying on extended thinking for reasoning audit trails.

### 5. Setting to disable copy-on-selection in agents view ([#60755](https://github.com/anthropics/claude-code/issues/60755))
36 👍, 8 comments. macOS-native terminal users want an explicit copy model (Cmd+C) instead of auto-copy on selection, which breaks clipboard workflows.

### 6. Agent resume after interruption ([#16607](https://github.com/anthropics/claude-code/issues/16607))
8 👍, 8 comments. Long-running sessions that hit context limits or are manually interrupted should let the agent pick up where it left off, not start over.

### 7. “Buy credits” button permanently disabled for free-tier users ([#62644](https://github.com/anthropics/claude-code/issues/62644))
0 👍, 8 comments. Free-tier accounts show a $500 limit but the billing page returns HTTP 429, making it impossible to upgrade. Billing infrastructure issue.

### 8. Opaque/deep compaction for long sessions ([#72461](https://github.com/anthropics/claude-code/issues/72461))
0 👍, 2 comments, filed today. Requests a compaction strategy beyond simple summarization to preserve semantic context in long-running coding sessions.

### 9. SSL certificate expired on macOS since v2.1.190+ ([#71663](https://github.com/anthropics/claude-code/issues/71663))
0 👍, 3 comments, regression. SSL verification failures on macOS likely due to bundled certificate store being out of sync. Affects all network calls.

### 10. `/code-review` skill misbehaves on long-lived branches ([#63048](https://github.com/anthropics/claude-code/issues/63048))
6 👍, 3 comments. Falls back to `main...HEAD` silently, reviewing other people’s commits. JSON-only output also hard to read.

---

## Key PR Progress

### Windows & cross-platform fixes (AZERDSQ131 series)
These 7 PRs address long-standing plugin breakage on Windows and older macOS:

- **[#68699](https://github.com/anthropics/claude-code/pull/68699)** – `fix(hookify)`: Python wrapper + platform root normalisation for Windows; Microsoft Store `python3` stub fix.
- **[#68701](https://github.com/anthropics/claude-code/pull/68701)** – `fix(security-guidance)`: Strip CRLF from Python version probe on Windows so version checks don’t fail.
- **[#68702](https://github.com/anthropics/claude-code/pull/68702)** – `fix(ralph-wiggum)`: Guard empty array expansion with `:-` for bash 3.x on macOS.
- **[#68686](https://github.com/anthropics/claude-code/pull/68686)** – `fix(hookify)`: Rename shadowed `field` variable and fix inline dict comma parsing in config loader.
- **[#68694](https://github.com/anthropics/claude-code/pull/68694)** – `fix(security-guidance)`: Normalise backslash separators in `CLAUDE_PLUGIN_ROOT` for hook commands on Windows.
- **[#68689](https://github.com/anthropics/claude-code/pull/68689)** – `fix(security-guidance)`: Block symlink escape in config reads to prevent local file disclosure.
- **[#68690](https://github.com/anthropics/claude-code/pull/68690)** – `fix(ralph-wiggum)`: Correct state file path in help.md (leading dot removed).
- **[#68693](https://github.com/anthropics/claude-code/pull/68693)** – `fix(scripts)`: Make duplicate label additive so existing platform/area labels aren’t erased on closure.

### Infrastructure & example improvements
- **[#72451](https://github.com/anthropics/claude-code/pull/72451)** – `fix: remove statsig.anthropic.com from init-firewall.sh`. The hostname no longer resolves, breaking devcontainer startup. Cleanup PR.
- **[#72361](https://github.com/anthropics/claude-code/pull/72361)** – `Add Claude Gateway on GCP example deployment assets`. Ready-to-use Terraform and scripts for running Gateway on Google Cloud.

### Feature plugin
- **[#68707](https://github.com/anthropics/claude-code/pull/68707)** – `feat(bug-reporter)`: Adds `/bug` slash command to file GitHub issues directly from the terminal, with auto-collected environment info and error logs.

### Documentation
- **[#72264](https://github.com/anthropics/claude-code/pull/72264)** – `docs(examples/hooks)`: Document additional Bash tool fields (`run_in_background`, `description`, `timeout`) exposed in pre-tool-use hooks.

---

## Feature Request Trends

1. **Account & session management**  
   Multi-account switching dominates (401 👍). Global session search across projects ([#71611](https://github.com/anthropics/claude-code/issues/71611)) and deep compaction for long sessions ([#72461](https://github.com/anthropics/claude-code/issues/72461)) are rising themes.

2. **IDE parity**  
   `/btw` in VS Code (119 👍) leads a broader push for feature equivalence between terminal and IDE, including hook event parity and status line updates.

3. **Configurability**  
   Users want knobs for: disabling copy-on-selection ([#60755](https://github.com/anthropics/claude-code/issues/60755)), disabling Ultra plan ([#71613](https://github.com/anthropics/claude-code/issues/71613)), and customisable session status indicators ([#71587](https://github.com/anthropics/claude-code/issues/71587)).

4. **Agent resilience**  
   Resuming after interruption ([#16607](https://github.com/anthropics/claude-code/issues/16607)), hook events on user interrupt (Esc) ([#71652](https://github.com/anthropics/claude-code/issues/71652)), and background agent status reflecting real activity ([#71587](https://github.com/anthropics/claude-code/issues/71587), [#71598](https://github.com/anthropics/claude-code/issues/71598)).

5. **Plugins & extensibility**  
   The `/bug` command PR is a community-led plugin example; more hook-based integrations and MCP server support in background sessions are requested.

---

## Developer Pain Points

- **Rate limiting (HTTP 429)** – Multiple reports ([#71635](https://github.com/anthropics/claude-code/issues/71635), [#71666](https://github.com/anthropics/claude-code/issues/71666), [#71673](https://github.com/anthropics/claude-code/issues/71673)) of server-side rate limiting on valid requests, especially on Windows CLI and macOS.
- **Tool call serialisation failures** – Raw XML leaks instead of execution ([#64108](https://github.com/anthropics/claude-code/issues/64108)), deeply concerning for agentic workflows.
- **Thinking display breakage** – Empty summaries on CLI ([#30958](https://github.com/anthropics/claude-code/issues/30958)) and stripped `thinking.display` on Vertex/Bedrock ([#71599](https://github.com/anthropics/claude-code/issues/71599)).
- **VS Code extension errors** – “Unsupported content type” for server/advisor tool results ([#40816](https://github.com/anthropics/claude-code/issues/40816)) still occurring.
- **SSL cert expiry on macOS** – Regression in v2.1.190+ ([#71663](https://github.com/anthropics/claude-code/issues/71663)).
- **Background agent pitfalls** – MCP servers not loaded ([#71597](https://github.com/anthropics/claude-code/issues/71597)), ESC key kills all agents ([#71623](https://github.com/anthropics/claude-code/issues/71623)), timer continues after turn completes ([#71598](https://github.com/anthropics/claude-code/issues/71598)).
- **Terminal keyboard conflicts** – Ctrl-L triggers `/clear` instead of screen repaint ([#71656](https://github.com/anthropics/claude-code/issues/71656)).
- **Stale UI state** – Status bar shows old working directory/git branch after `cd` ([#71661](https://github.com/anthropics/claude-code/issues/71661)).

---

*Digest generated from GitHub data for `anthropics/claude-code` as of 2026-06-30. All links are to the respective issues/PRs.*

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-06-30

## Today’s Highlights
Capacity and rate-limit issues continue to dominate community discussions, with multiple reports of “Selected model is at capacity” across Windows, macOS, and various subscription tiers. A user-discovered reasoning token clustering pattern in GPT-5.5 responses has sparked investigation into potential performance degradation on complex tasks. On the engineering side, a major multi-PR push landed queued user message infrastructure, enabling follow-up messages to survive client restarts.

---

## Releases
- **rust-v0.143.0-alpha.31** — No detailed changelog was provided. This alpha continues the Rust-based tooling evolution for Codex.

---

## Hot Issues (10 notable)

1. **[#30224](https://github.com/openai/codex/issues/30224)** — `[bug, custom-model, app, config] This model is not supported when using X-OpenAI-Internal-Codex-Responses-Lite`  
   61 comments, 20 👍. A blocking error for users leveraging the internal responses lite header. High engagement suggests many rely on this custom pipeline.

2. **[#29760](https://github.com/openai/codex/issues/29760)** — `[bug, rate-limits, CLI] Selected model is at capacity. Please try a different model.`  
   29 comments. Affects Pro and Plus users on macOS and Windows. Capacity errors are the most frequently upvoted pain point this week.

3. **[#29072](https://github.com/openai/codex/issues/29072)** — `[bug, windows-os, sandbox, tool-calls, app] Windows Codex App: apply_patch fails because codex-windows-sandbox-setup.exe cannot launch from package path`  
   27 comments, 19 👍. Windows sandboxing remains a fragile spot; users cannot apply patches at all on certain builds.

4. **[#2880](https://github.com/openai/codex/issues/2880)** — `[enhancement, TUI] Copy/Export Message as Markdown`  
   24 comments, 71 👍. The most upvoted issue in the last 24h. Developers want a native way to export conversation for documentation and issue tracking.

5. **[#30364](https://github.com/openai/codex/issues/30364)** — `[bug, model-behavior, rate-limits] GPT-5.5 Codex reasoning-token clustering at 516/1034/1552 may be leading to degraded performance on complex tasks`  
   22 comments, 32 👍. A data-driven analysis of token count spikes that may correlate with lower reasoning quality. OpenAI engineers have acknowledged the pattern.

6. **[#2916](https://github.com/openai/codex/issues/2916)** — `[enhancement, config] OpenAI service tier support`  
   18 comments, 50 👍. Users want explicit control over API service tiers to manage cost/latency trade-offs, especially for batch jobs.

7. **[#29418](https://github.com/openai/codex/issues/29418)** — (Closed) `[bug, windows-os, sandbox, app] [Windows] codex-windows-sandbox-setup.exe fails with "The specified module could not be found"`  
   Similar to #29072 but closed – may indicate a fix in progress or duplicate consolidation.

8. **[#29492](https://github.com/openai/codex/issues/29492)** — `[bug, windows-os, sandbox, app] Windows Codex desktop app creates an empty .git folder, then spawns git process repetitively`  
   A resource-wasting bug that fills disk I/O with infinite git spawns when a project isn’t under version control.

9. **[#30440](https://github.com/openai/codex/issues/30440)** — `[bug, sandbox, app] Codex uses bundled pnpm instead of host toolchain`  
   Build scripts that depend on global pnpm are broken because Codex invokes its own stale bundled version. 9 👍.

10. **[#30639](https://github.com/openai/codex/issues/30639)** — `[bug, rate-limits, app, config, computer-use, memory] Codex Desktop (macOS): Chronicle runs continuous background screen-recording summaries every 10 min, draining plan limits`  
    Most recent report (same day). Background “Chronicle” screen capture silently consumes quota, angering users who disable it without immediate effect.

---

## Key PR Progress (10 important changes)

1. **[#28307](https://github.com/openai/codex/pull/28307)** — `feat: queue TUI follow-ups through app-server`  
   Moves queued follow-up messages from client memory to the app-server, making them persist across TUI process restarts.

2. **[#28267](https://github.com/openai/codex/pull/28267)** — `feat: dispatch queued user messages through core idle extensions`  
   Integrates queued messages into the core idle extension pipeline, preventing races with tool results.

3. **[#28265](https://github.com/openai/codex/pull/28265)** — `feat: accept user submissions at idle turn boundaries`  
   Allows queued user messages to enter the response loop with full context, telemetry, and Plan mode support.

4. **[#28266](https://github.com/openai/codex/pull/28266)** — `feat: add durable user message queue storage`  
   Introduces a dedicated `queue_1.sqlite` database with transactional claim semantics for multi-process safety.

5. **[#25283](https://github.com/openai/codex/pull/25283)** — `feat: synchronize runtime workspace roots in thread settings`  
   Ensures queued turns and direct turns see consistent workspace roots by storing them in thread settings.

6. **[#28425](https://github.com/openai/codex/pull/28425)** — `[codex] Carry fork lineage in initial history`  
   Cleanly tracks fork ancestry in `InitialHistory` instead of splitting metadata across session configuration and rollouts.

7. **[#27945](https://github.com/openai/codex/pull/27945)** — `Seed session pickers from the state DB`  
   Drastically reduces startup latency for fork/resume pickers by serving indexed data from the state DB while the filesystem scan runs in the background.

8. **[#25629](https://github.com/openai/codex/pull/25629)** — `[plugins] Add remote plugin search tool`  
   A model-visible tool to query the ChatGPT plugin discovery endpoint, allowing Codex models to search and suggest plugins during conversations.

9. **[#28378](https://github.com/openai/codex/pull/28378)** — `[codex] Retry transient models fetch failures`  
   Adds retry logic to the Rust release preparation workflow, fixing intermittent failures when fetching the models catalog.

10. **[#28337](https://github.com/openai/codex/pull/28337)** — `[code-reviewed] Add gated MCP SEP-2631 file transfer`  
   Implements disabled-by-default support for the draft MCP file transfer protocol, laying groundwork for cross-client file sharing.

---

## Feature Request Trends

- **Markdown export** (#2880) remains the single most requested UX enhancement, with developers wanting to capture Codex sessions for documentation and bug reports.
- **Service tier control** (#2916) is the second most upvoted feature; users want to choose between low-cost API tiers and high-performance inference per task.
- **Quota reset transparency** (#30686) has surfaced again—users are confused by inconsistent reset limits across subscription plans and request a predictable policy.
- **Plugin ownership metadata** (see PR #26340) – clients need canonical labels to distinguish plugin-owned skills from user-created ones, improving skill management.

---

## Developer Pain Points

- **Model capacity errors** dominate the bug tracker: issues #29760, #28507, #30575, #30579 all report the same “model at capacity” message across platforms. Pro and Plus users are equally affected, especially during peak hours.
- **Windows sandboxing** is a persistent weak spot: #29072, #29418, #29492, and #30486 all expose failures in the Windows sandbox setup, patch application, or tool exposure.
- **Background resource drain**: #30639 (Chronicle screen recording) and #30525 (unexpected idle quota consumption) show that desktop app background processes can silently exhaust user quotas without clear UI indication.
- **Performance regressions**: #26992 (CLI startup slowdown in 0.137.0) and #30696 (extreme slowness after recent update) indicate that newer builds occasionally ship with degraded responsiveness.
- **MCP tool inconsistency**: #30486 and #30699 highlight that MCP servers may be registered but their tools not actually exposed to the model, breaking plugin workflows.
- **Context/reasoning bugs**: #30364 (GPT-5.5 token clustering) and #24069 (subagent regression with Ollama) point to model-level issues that can silently reduce output quality or break custom providers.

---

*Digest generated from GitHub data (openai/codex) as of 2026-06-30. All links point to the respective issues and pull requests.*

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest – 2026-06-30

## Today's Highlights
A new nightly release (v0.51.0) brings another day of incremental improvements, while the community continues to flag critical agent reliability issues—most notably crashes on shell command confirmation prompts and subagents hanging indefinitely. On the PR side, a major MCP elicitation feature and a fix for recursive reasoning turns are progressing, alongside security hardening for file-write scoping.

## Releases
- **v0.51.0-nightly.20260630.gae0a3aa7b**: Nightly release with full changelog available.  
  [Changelog](https://github.com/google-gemini/gemini-cli/compare/v0.51.0-nightly.20260629.gae0a3aa7b...v0.51.0-nightly.20260630.gae0a3aa7b)

## Hot Issues
1. **#14561 – CLI crashes on `run_shell_command` confirmation prompt**  
   *Closed / Need‑triage* – A hard crash when the terminal tries to display the user confirmation menu for shell commands. 8 comments, 3 👍.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/14561)

2. **#22323 – Subagent reports GOAL success after MAX_TURNS hit**  
   *P1, Bug* – A subagent (`codebase_investigator`) returns “success” and “Termination Reason: GOAL” even though it was interrupted by the turn limit. Misleads users into thinking work was completed. 8 comments.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/22323)

3. **#21409 – Generalist agent hangs forever**  
   *P1, Bug* – Deferring to the generalist agent causes indefinite hangs even for simple folder creation. Workaround: instruct the model not to use sub‑agents. 7 comments, 8 👍.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/21409)

4. **#25166 – Shell command gets stuck with “Waiting input” after completion**  
   *P1, Bug* – After executing a trivial command, the CLI shows the command as still active and awaiting input, even though the process has finished. 4 comments, 3 👍.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/25166)

5. **#21983 – Browser agent fails in Wayland**  
   *P1, Bug* – The browser subagent terminates with “GOAL” but fails to operate properly under Wayland. 4 comments, 1 👍.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/21983)

6. **#22186 – “get-shit-done” output hook causes crash**  
   *P1, Bug* – Repeated crashes when the `get-shit-done` mode prints its final summary. 3 comments.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/22186)

7. **#21968 – Gemini does not skill/sub‑agent usage**  
   *P2, Bug* – Anecdotal evidence that the agent rarely picks custom skills and sub‑agents unless explicitly told, even when the task is highly relevant. 6 comments.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/21968)

8. **#22267 – Browser agent ignores `settings.json` overrides (e.g., maxTurns)**  
   *P2, Bug* – Configuration values from global or project `settings.json` are completely ignored by the Browser Agent. 3 comments.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/22267)

9. **#22672 – Agent should discourage destructive behavior**  
   *P2, Feature Request* – Agents often resort to `git reset --force` or risky database commands when safer options exist. 3 comments, 1 👍.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/22672)

10. **#24353 – Robust component‑level evaluations (EPIC)**  
    *P1, Evaluation* – Follow‑up to the initial behavioral eval system. Aims to expand from 76 tests to a comprehensive evaluation suite across all supported Gemini models. 7 comments.  
    [Link](https://github.com/google-gemini/gemini-cli/issues/24353)

## Key PR Progress
1. **#28089 – MCP elicitation (form + url) capability**  
   *Area: extensions, size: L* – Implements the MCP specification for elicitation modes, allowing clients to advertise and accept form-based and URL-based interactions.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28089)

2. **#28164 – Limit recursive reasoning turns per request**  
   *Size: M* – Caps recursive reasoning at 15 turns per user request (configurable via `maxSessionTurns`), protecting local CPU and API quotas from infinite loops.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28164)

3. **#27971 – Strip thoughts from scrubbed history turns**  
   *Size: M/L* – Fixes “Thought Leakage” where the model’s internal monologue appears in plain‑text history, causing confusion and monologue loops in subsequent turns.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27971)

4. **#28053 – Defensive path resolution for `@`‑prefixed files**  
   *Size: XL* – Resolves a production bug where tools (`read_file`, `replace`, `write_file`) fail with “File not found” when the model passes paths like `@policies/new-policies.txt`. Also fixes macOS tests.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28053)

5. **#28215 – Harden file‑write scope: stop sandboxed writes to `.gemini` and `.gitconfig`**  
   *Size: M* – Prevents prompt‑injection‑based sandbox escapes by blocking auto‑accepted writes to the `.gemini/` folder and `.gitconfig` files.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28215)

6. **#28163 – Caretaker agent triage worker core (part 1)**  
   *Size: L* – Foundation for a Cloud Run–based triage worker that ingests and processes GitHub issues. Part 1 introduces core modules and data structures.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28163)

7. **#28015 – Cloud Run webhook ingestion service**  
   *Size: L/XL* – Implements the full webhook ingestion pipeline for the Caretaker Agent, including signature verification, Firestore storage, and Pub/Sub publishing.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28015)

8. **#28219 – Parse commented `settings.json` in memory bootstrap**  
   *Size: S* – Fixes the lightweight parent process’s inability to read JSON settings files that contain comments, preventing incorrect defaults for memory auto‑configuration.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28219)

9. **#28096 – Drop late tool calls after SIGINT cancellation**  
   *Area: core, size: M* – Prevents delayed provider tool‑call chunks from being executed after the user has already cancelled, avoiding unintended side effects.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28096)

10. **#28099 – Show descriptive sandbox label in footer**  
    *Area: core, size: S* – Replaces the hardcoded “current process” string in the footer’s sandbox indicator with the actual sandbox profile name (e.g., `sandbox-exec`). Closes #26697.  
    [Link](https://github.com/google-gemini/gemini-cli/pull/28099)

## Feature Request Trends
- **AST‑aware tooling**: Several issues (#22745, #22746) request AST‑based file reads, search, and codebase mapping to reduce token waste and improve navigation accuracy.
- **Enhanced evaluation infrastructure**: #24353 and #15300 signal a push for comprehensive component‑level behavioral evals, including sub‑agent trajectories and shareable chat records (#22598).
- **Agent self‑awareness**: #21432 asks the agent to accurately know its own CLI flags, hotkeys, and capabilities so it can guide users effectively.
- **Memory system improvements**: #26516 and its sub‑issues call for better automatic memory extraction, redaction, and handling of low‑signal sessions—reducing noise and preventing infinite retries.
- **Browser agent resilience**: #22232 proposes automatic session takeover and lock recovery for persistent browser profiles under Wayland and other environments.

## Developer Pain Points
- **Unexplained crashes & hangs**: The CLI crashes when showing shell confirmation prompts (#14561) and hangs on deferring to sub‑agents (#21409) or after shell commands finish (#25166). These destabilize everyday workflows.
- **Misleading agent behavior**: Sub‑agents report success after hitting turn limits (#22323), and the generalist agent sometimes runs without permission (#22093) or ignores user‑configured skills (#21968).
- **Settings & configuration ignored**: The browser agent skips `settings.json` overrides (#22267) and sub‑agent settings are not respected.
- **Destructive operations**: Agents frequently use risky git commands and `--force` flags (#22672), causing workspace cleanup headaches.
- **Browser & Wayland issues**: The browser sub‑agent fails on Wayland (#21983) and lacks proper lock recovery (#22232).
- **Tool count limits**: Encountering a 400 error when more than 128 tools are available (#24246) forces manual tool pruning.
- **External editor & terminal corruption**: Exiting an external editor in terminal buffer mode leaves terminal corruption (#24935).

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-06-30

## Today's Highlights
A new release (v1.0.66-2) introduces multi-plugin skill coexistence and LSP server logs, alongside a fix to prompt for `gh` CLI installation when missing. Community attention is focused on a critical memory‑consumption bug (`tgrep` OOM on large monorepos) and a regression that breaks trackpad scrolling on macOS. The only PR merged today adds AI‑powered issue classification, expected to improve triage velocity.

## Releases
**v1.0.66-2** — [View on GitHub](https://github.com/github/copilot-cli/releases/tag/v1.0.66-2)  
**Added**  
- Skills with identical names from different plugins can now coexist.  
- Integrations can read and write CLI user settings.  
- LSP server logs are exposed via `/lsp logs` and `read_agent`.  
- The CLI now prompts to install `gh` when it is missing in GitHub repositories.  
- GitHub attachment variants are supported in prompt rendering.

## Hot Issues *(10 of 23 updated today)*

1. **#1665 – Project‑scoped plugins (closed)**  
   *High demand (17 👍, 10 comments).* Proposed allowing plugins to be scoped to a project/repository instead of globally per user. Closed—likely resolved or superseded.  
   [Issue #1665](https://github.com/github/copilot-cli/issues/1665)

2. **#1799 – Turn off alt‑screen views**  
   *7 👍, 10 comments.* Users are frustrated by the forced alt‑screen rendering; many want an option to revert to the original terminal output.  
   [Issue #1799](https://github.com/github/copilot-cli/issues/1799)

3. **#3874 – `preToolUse` agent hook denial does not work**  
   *Security/plugin bug.* A hook intended to deny all commands is ignored, potentially allowing unauthorized tool execution. No community traction yet but critical for safety.  
   [Issue #3874](https://github.com/github/copilot-cli/issues/3874)

4. **#3948 – `web_fetch` always fails with `TypeError: fetch failed`**  
   *Networking bug.* Every `web_fetch` tool call fails, unrelated to proxy or authentication. Impacts any user relying on web content retrieval.  
   [Issue #3948](https://github.com/github/copilot-cli/issues/3948)

5. **#3957 – Trackpad scrolling regression on macOS**  
   *5 👍, regression.* Scrolling with a trackpad selects prompts instead of scrolling the window. Affected version: 1.0.65 on Ghostty 1.3.1.  
   [Issue #3957](https://github.com/github/copilot-cli/issues/3957)

6. **#3976 – Native `tgrep` indexer OOM‑kills the host on large monorepos**  
   *Critical new bug.* The experimental `tgrep` daemon has no memory cap, causing out‑of‑memory kills on sizable repositories. No comments yet, but urgent.  
   [Issue #3976](https://github.com/github/copilot-cli/issues/3976)

7. **#3954 – `explore` tool hardcodes model to `gpt-5.4-mini` ignoring custom config**  
   *Model configuration bug.* The `explore` tool ignores user‑configured models (e.g., DeepSeek) and always tries `gpt-5.4-mini`. Blocks custom model users.  
   [Issue #3954](https://github.com/github/copilot-cli/issues/3954)

8. **#3955 – Drag‑and‑drop file attachment broken on macOS (regression)**  
   *UX regression.* Dropping files from Finder into the Copilot app no longer attaches them. Regression from a previous version.  
   [Issue #3955](https://github.com/github/copilot-cli/issues/3955)

9. **#3972 – UI displays mouse movement characters instead of proper interface**  
   *Rendering bug.* On first load, a stream of characters representing mouse movements is shown, breaking the UI entirely.  
   [Issue #3972](https://github.com/github/copilot-cli/issues/3972)

10. **#3973 – MCP OAuth re‑auth fails on Windows due to cached loopback port exclusion**  
    *Platform‑specific bug.* Repeated OAuth re‑auth attempts fail on Windows when the cached redirect port falls within an excluded TCP range. Manual workaround required.  
    [Issue #3973](https://github.com/github/copilot-cli/issues/3973)

## Key PR Progress
Only one pull request was updated in the last 24 h:

- **#2587 – Add automated issue classification with GitHub Agentic Workflows** *(closed)*  
  Introduces an AI‑powered workflow (`gh-aw`) that automatically applies `area:` labels and the `triage` label when issues are opened or reopened. This should streamline issue triage and categorisation. No community comments.  
  [PR #2587](https://github.com/github/copilot-cli/pull/2587)

## Feature Request Trends
The most‑requested directions from recent issues include:

- **Per‑project / repository‑scoped plugins** (#1665) – users want to avoid global plugin installation.
- **Parity with Claude Code** – features like `Ctrl+G` to expand paste tokens in `$EDITOR` (#3936) and the ability to disable alt‑screen views (#1799).
- **Rich session management** – a full file‑tree browser for repository‑backed sessions (#3971) and desktop notifications when the CLI awaits input (#2941).
- **MCP enhancements** – ability to save HTTP headers in the GUI (#2849), warning on MCP server name conflicts (#3893), and more robust OAuth handling.
- **Better plugin installation** – support for Windows git symlinks (#2286).

## Developer Pain Points
Recurring frustrations and high‑frequency requests:

- **Forced alt‑screen rendering** – no opt‑out, causing workflow disruption (#1799).
- **Regressions** – trackpad scrolling (#3957) and drag‑and‑drop (#3955) worked in earlier versions but broke recently.
- **Instability on large repos** – `tgrep` consumes unbounded memory, leading to OOM kills (#3976).
- **Model hardcoding** – the `explore` tool ignores custom model configurations (#3954) – a blocker for users of alternative providers.
- **Network tool failures** – `web_fetch` consistently fails (#3948) with no clear workaround.
- **Windows‑specific pain** – MCP OAuth re‑auth loopback port conflicts (#3973) and git symlink issues for plugin install (#2286).
- **Security gaps** – `preToolUse` hook denials are not enforced (#3874), undermining plugin safety mechanisms.

---

*Generated from [github.com/github/copilot-cli](https://github.com/github/copilot-cli) – data as of 2026‑06‑30.*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest – 2026-06-30

**Data source:** [github.com/MoonshotAI/kimi-cli](https://github.com/MoonshotAI/kimi-cli)

---

## Today’s Highlights

No new releases or issues were created in the last 24 hours, but two long‑standing pull requests received updates, signaling ongoing refinement of the shell interface. The community’s focus remains on improving user visibility and interactive workflow flexibility, with a bright‑blue highlight for user input and a new `--prompt-interactive` option nearing conclusion.

---

## Releases

No new releases in the last 24 hours.

---

## Hot Issues

**No new issues were updated in the last 24 hours.** The issue tracker currently shows zero open issues with recent activity. This may indicate a period of low bug‑reporting or a lull after a previous burst of fixes. We will monitor for any emerging threads in the next digest.

---

## Key PR Progress

Two pull requests saw activity in the last 24 hours. While only a handful are available, each is significant for shell UX:

### 1. **PR #1600 – [OPEN] feat(shell): highlight user input with bright blue and separator**
   - **Author:** liuchong  
   - **Created:** 2026‑03‑27 | **Updated:** 2026‑06‑30  
   - **Summary:** Applies a bright blue color (`#007AFF`) to user echo text in the shell and adds a full‑width separator line below each user input to improve visual distinction.  
   - **Why it matters:** This addresses a common UI readability pain point, especially in long chat sessions. The update has been open for three months – its recent activity may signal an approaching merge.  
   - **[View PR #1600](https://github.com/MoonshotAI/kimi-cli/pull/1600)**

### 2. **PR #2246 – [CLOSED] feat(shell): add `--prompt-interactive` option**
   - **Author:** shuizhongyueming  
   - **Created:** 2026‑05‑12 | **Updated:** 2026‑06‑30  
   - **Summary:** Implements a new CLI flag `--prompt-interactive` (short `-P`) that lets users supply an initial prompt when starting the shell UI, then keeps the session open for follow‑up questions.  
   - **Why it matters:** This feature was requested in issue #2240 and resolves the need for scriptable one‑shot interactions without losing the conversational context. The PR was closed – likely merged – making this a notable enhancement for power users and automation workflows.  
   - **[View PR #2246](https://github.com/MoonshotAI/kimi-cli/pull/2246)**

---

## Feature Request Trends

Despite the quiet day, the two active PRs reveal the community’s top feature directions:

- **Shell UI readability enhancements** – Users want clearer visual separation of roles (user vs. AI) and higher contrast for accessibility.
- **Flexible interactive prompts** – The desire to inject an initial query when launching the shell, while preserving the ability to continue the conversation, points to a demand for both automation and fluid interactivity.

No new feature requests were filed in the last 24 hours, so these remain the dominant themes.

---

## Developer Pain Points

- **Visual clutter in long sessions** – The bright‑blue highlight PR (#1600) directly addresses the difficulty of distinguishing user inputs from AI responses in scrolling chat logs.
- **Lack of scriptable initial prompts** – The closed PR #2246 solved the pain point of having to manually type the first query when using the CLI in scripts or pipelines. This was a long‑standing gap (issue #2240) that the community had been vocal about.

No additional pain points surfaced today, suggesting that these two items are the most pressing for current users.

---

*Digest generated from GitHub activity between 2026‑06‑29 00:00 UTC and 2026‑06‑30 23:59 UTC.*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-30

## Today’s Highlights
The most active area today is model configuration and integration—three new bug reports confirm issues with custom OpenAI-compatible providers, Bedrock extended thinking, and MiniMax-M3’s hardcoded thinking type. On the PR side, the team merged fixes for agent temperature precedence, Bedrock DeepSeek model ID preservation, and restored chunk type safety in the Copilot provider. Community demand for model fallback/failover and a “YOLO” permission bypass mode continues to dominate the feature request pipeline.

## Releases
No new releases in the last 24 hours.

## Hot Issues (10 picks)

**[#7602 – Native Model Fallback / Failover Support](https://github.com/anomalyco/opencode/issues/7602)**  
**90 👍 | 24 comments**  
Request for automatic retry between different models (e.g., “if model A rate-limits, fall back to model B”). Currently OpenCode only supports provider fallback when the model ID is identical. This is the single most upvoted open feature request, reflecting a clear pain point for reliability.

**[#8463 – `--dangerously-skip-permissions` (YOLO mode)](https://github.com/anomalyco/opencode/issues/8463)**  
**89 👍 | 22 comments**  
A heavily requested flag to bypass permission prompts in trusted environments. Community discussion focuses on automated workflows and CI/CD where interactive prompts break scripts.

**[#5674 – Custom OpenAI-compatible provider options not passed](https://github.com/anomalyco/opencode/issues/5674)**  
**13 👍 | 25 comments**  
Long-standing bug: `baseURL` and `apiKey` from `opencode.json` are silently dropped when using `@ai-sdk/openai-compatible`. High engagement, with users providing workarounds but no fix yet.

**[#25239 – Expose GitHub Copilot “Auto” model selector](https://github.com/anomalyco/opencode/issues/25239)**  
**14 👍 | 13 comments**  
Request to let users choose Copilot’s auto-routed model (the `/models/session` endpoint) directly in OpenCode, mimicking VS Code’s experience.

**[#20235 – GitHub Copilot auto model routing API access + plugin hook](https://github.com/anomalyco/opencode/issues/20235)**  
**25 👍 | 8 comments**  
Companion to #25239: asks for a plugin hook so the Copilot auto-routing can be extended or logged. Strong community interest in deeper Copilot integration.

**[#15907 – Clipboard copy not working over SSH + tmux in Ghostty](https://github.com/anomalyco/opencode/issues/15907)**  
**10 👍 | 10 comments**  
Users rely on remote development; the false positive “copied to clipboard” notification is misleading. Affects productivity for remote workers.

**[#34250 – OpenAI GPT models (ChatGPT Plus browser auth) return no response](https://github.com/anomalyco/opencode/issues/34250)**  
**0 👍 | 4 comments**  
Burning issue for ChatGPT Plus users: prompts are accepted but no response is returned. No error in TUI—hard to debug.

**[#33630 – Bedrock Converse cannot enable extended thinking](https://github.com/anomalyco/opencode/issues/33630)**  
**0 👍 | 3 comments**  
Critical for Claude users on Bedrock: `thinking` config is silently ignored, so no reasoning tokens are ever returned.

**[#34582 – MCP OAuth: access token not refreshed despite refresh token present](https://github.com/anomalyco/opencode/issues/34582)**  
**0 👍 | 2 comments**  
Token refresh broken for OAuth-protected MCP servers. Blocks use of external MCP servers with short-lived tokens.

**[#34570 – MiniMax-M3 ignores user-configured thinking.type](https://github.com/anomalyco/opencode/issues/34570)**  
**0 👍 | 0 comments**  
Fresh report: two code paths hardcode `"adaptive"` over the user’s `"enabled"` setting. Indicates a broader pattern of model-specific overrides.

## Key PR Progress (10 picks)

**[#34583 – fix: respect agent temperature config when capabilities.temperature is false](https://github.com/anomalyco/opencode/pull/34583)**  
Fixes a regression for custom OpenAI-compatible providers. Agent-level temperature now takes precedence over model capability checks.

**[#34441 – fix: preserve Bedrock DeepSeek model ids](https://github.com/anomalyco/opencode/pull/34441)**  
Prevents cross-region model ID mangling (e.g., `deepseek.v3.2` was incorrectly transformed). Directly addresses #34412.

**[#34590 – fix(llm): forward bedrock.thinking provider option to Bedrock Converse](https://github.com/anomalyco/opencode/pull/34590)**  
Adds `reasoning_config` via `additionalModelRequestFields`, enabling extended thinking on Bedrock-hosted Claude models. Closes #33630.

**[#33636 – chore(github-copilot): restore stream chunk type safety](https://github.com/anomalyco/opencode/pull/33636)**  
Fixes a MUST FIX TODO: restores TypeScript type safety on chunk stream transforms in the Copilot provider.

**[#34597 – feat(app): autocomplete mcp resources](https://github.com/anomalyco/opencode/pull/34597)**  
Adds MCP resource autocomplete in the prompt `@` menu, pulling from the app directory sync. Improves discoverability of available resources.

**[#34175 – fix(app): restore prompt cursor on focus](https://github.com/anomalyco/opencode/pull/34175)**  
Fixes web prompt contenteditable cursor position restoration—usability win for the desktop/web UI.

**[#34595 – fix(desktop): persist last active url](https://github.com/anomalyco/opencode/pull/34595)**  
Wraps MemoryRouter with localStorage persistence, so the desktop app remembers the last route across sessions.

**[#34307 – feat(tui): insert absolute path for unsupported pastes and toast on failure](https://github.com/anomalyco/opencode/pull/34307)**  
Two Windows-specific TUI paste fixes: inserts absolute path when clipboard read fails, and shows a toast on failure.

**[#34591 – feat(app): hide separators around active tabs](https://github.com/anomalyco/opencode/pull/34591)**  
UI polish: removes separator lines on both sides of the active titlebar tab, reducing visual noise.

**[#30719 – feat(i18n): add Italian (it) locale](https://github.com/anomalyco/opencode/pull/30719)**  
Community contribution adding full Italian UI translation. Closes #28641.

## Feature Request Trends
- **Model fallback & failover** (#7602) – most upvoted feature, users want retry chains with different models.
- **YOLO/skip permissions mode** (#8463) – second most upvoted; strong demand for headless/automated operation.
- **GitHub Copilot deep integration** (#25239, #20235) – requests for auto model routing, plugin hooks, and model selector exposure.
- **JetBrains ACP integration** (#34551) – reasoning effort/level selector missing in JetBrains AI Assistant.
- **Built-in `/insights`** (#12981) – want it to be a first-class command, not just a prompt wrapper around `opencode stats`.
- **Desktop app `/export`** (#31453) – parity with TUI’s Markdown export.

## Developer Pain Points
- **Custom provider options silently dropped** (#5674) – ongoing frustration for self-hosted/enterprise setups.
- **Bedrock extended thinking not working** (#33630) – blocks reasoning features for Claude on AWS.
- **MCP OAuth token refresh broken** (#34582, #34592) – missing `resource` parameter causes 401 with Atlassian and other OAuth servers.
- **ChatGPT Plus browser auth silent failure** (#34250) – no feedback, no error.
- **MiniMax-M3 thinking type hardcoded** (#34570) – user config ignored, bypass intended behavior.
- **Remote clipboard copy not working** (#15907) – misleading success notification.
- **Paid subscription not activating** (#32420) – although closed, multiple identical reports signal a support gap.
- **Zero responses from OSS LLMs** (#29605) – on Ubuntu/Codium/Zen, models don’t return any output.
- **ARM64 installer missing executable** (#34581) – Windows on ARM users left without a working install.

*Data sourced from [anomalyco/opencode](https://github.com/anomalyco/opencode) Issues and Pull Requests updated on 2026-06-30.*

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-06-30

## Today’s Highlights
The past 24 hours saw substantial progress on two long‑standing pain points: the streaming markdown scroll‑jump bug (#5825) is being addressed by PRs that stabilise the TUI status row and introduce optional padding control. Meanwhile, provider‑side error transparency improved significantly with the merge of PR #5832, which surfaces HTTP error bodies instead of opaque SDK messages. Developer attention also shifted toward reliability: a new fix for hung streams on Bedrock (#6051) and a crash‑causing `ECONNRESET` during streaming (#6133) have both been resolved.

## Releases
No new versions were published in the last 24 hours.

## Hot Issues
*(10 noteworthy issues, selected for community impact, severity, or active discussion)*

1. **#5825 – Streaming markdown forces scroll to bottom**  
   *Comments: 42*  
   A highly‑reported UX bug: when `clear on shrink` is enabled, Pi scrolls to the bottom while the user is reading earlier content. Two related PRs (#6169, #6026) are now in flight.  
   [Issue #5825](https://github.com/earendil-works/pi/issues/5825)

2. **#4877 – Session folder collision**  
   *Comments: 20*  
   Sessions at distinct paths (e.g., `/a/b/c/d` vs `/a-b/c-d`) can map to the same folder due to simple path‑to‑name conversion. Though cosmetic, it risks confusing users when projects accidentally share state.  
   [Issue #4877](https://github.com/earendil-works/pi/issues/4877)

3. **#4338 – Agent says “working” but makes no progress**  
   *Comments: 6*  
   A classic loop scenario where the agent claims to be working but produces no output. The issue was closed as a “weekend/refactor” candidate, indicating a complex root cause that needs deeper architectural changes.  
   [Issue #4338](https://github.com/earendil-works/pi/issues/4338)

4. **#6019 – OpenAI Responses mid‑stream retryable error is not retried**  
   *Comments: 5*  
   OpenAI explicitly marks some errors as retryable, yet Pi finalises the message with `stopReason: "error"`. This breaks reliability for users relying on the OpenAI Responses API.  
   [Issue #6019](https://github.com/earendil-works/pi/issues/6019)

5. **#5763 – Providers swallow the HTTP error body**  
   *Comments: 5*  
   A debugging nightmare: behind a proxy, a 403 from Bedrock returns `Unknown: UnknownError` while OpenAI gives `403 status code (no body)`. Fixed by PR #5832, which now surfaces the original error body.  
   [Issue #5763](https://github.com/earendil-works/pi/issues/5763)

6. **#5901 – Durable HITL tool‑call interrupts**  
   *Comments: 4*  
   Feature request for human‑in‑the‑loop approval of tool calls, similar to LangGraph’s middleware. The author has a concrete proposal and discusses implementation. Open for discussion.  
   [Issue #5901](https://github.com/earendil-works/pi/issues/5901)

7. **#6138 – Incorrect pricing for Xiaomi MiMo native provider models**  
   *Comments: 4*  
   Hardcoded pricing in `xiaomi.models.js` does not match official Xiaomi pay‑as‑you‑go rates. A quick‑fix candidate that affects cost tracking for users on MiMo models.  
   [Issue #6138](https://github.com/earendil-works/pi/issues/6138)

8. **#6103 – OpenAI Responses mislabels empty tool results as “(see attached image)”**  
   *Comments: 3*  
   A latent bug exposed by the `pi-hashline-edit-pro` extension: empty tool results are incorrectly displayed as image attachments. The root cause lies in core response handling.  
   [Issue #6103](https://github.com/earendil-works/pi/issues/6103)

9. **#5463 – Auto‑compaction after final turn throws error**  
   *Comments: 3*  
   Auto‑compaction triggers after the last assistant message, draining both queues and throwing `Error("Cannot continue from message role: assistant")`. Affects coding agent sessions.  
   [Issue #5463](https://github.com/earendil-works/pi/issues/5463)

10. **#6133 – pi crashes with uncaughtException: ECONNRESET during streaming**  
    *Comments: 2*  
    When the upstream provider resets the TCP connection mid‑stream, undici’s `TypeError: terminated` escapes the try/catch and crashes the process. A high‑severity reliability bug.  
    [Issue #6133](https://github.com/earendil-works/pi/issues/6133)

## Key PR Progress
*(10 PRs merged or open in the last 24 hours, chosen for impact or design significance)*

1. **#6169 – Disable padding for assistant messages**  
   Closes #6168. Introduces a `--no-padding` flag (or config) to remove the TUI padding that many users have requested on Discord.  
   [PR #6169](https://github.com/earendil-works/pi/pull/6169)

2. **#6182 – Add redo support to editors**  
   Implements redo operations in the TUI, completing the undo/redo feature requested in #959.  
   [PR #6182](https://github.com/earendil-works/pi/pull/6182)

3. **#6175 – Emit session name changes to extensions**  
   Fixes #6174: extensions now receive events when a session is renamed, enabling consistent state in custom UIs.  
   [PR #6175](https://github.com/earendil-works/pi/pull/6175)

4. **#6115 – Configurable chat padding**  
   A discussion‑oriented PR that proposes flag‑based control over chat padding. Not yet decided whether this is the best approach, but it demonstrates growing demand for layout customisation.  
   [PR #6115](https://github.com/earendil-works/pi/pull/6115)

5. **#6178 – Guard against undefined content in tool result messages**  
   Fixes a crash when extension tools (e.g., `get_kline`) return empty content. Adds defensive checks before rendering.  
   [PR #6178](https://github.com/earendil-works/pi/pull/6178)

6. **#5832 – Surface provider HTTP error body instead of opaque message**  
   Merges the fix for #5763. Provider error handling now preserves the HTTP body, allowing developers to read gateway/proxy error messages directly.  
   [PR #5832](https://github.com/earendil-works/pi/pull/5832)

7. **#6176 – Apply extension tool changes before next provider request**  
   *(open PR)*  
   Fixes #6162: when an extension tool calls `pi.setActiveTools()`, the updated tool list is now immediately used in the next provider request within the same run.  
   [PR #6176](https://github.com/earendil-works/pi/pull/6176)

8. **#6170 – Avoid replaying historical inline images**  
   Reduces context bloat: historical session replay now replaces inline image escape payloads with lightweight `[Image: ...]` labels, while live tool results continue to render images.  
   [PR #6170](https://github.com/earendil-works/pi/pull/6170)

9. **#6051 – Recover from hung streams and retry unmodeled Bedrock errors**  
   Adds `streamIdleTimeoutMs` and `connectTimeoutMs` to detect half‑open sockets; throws retryable timeouts instead of blocking indefinitely. Also retries 5xx errors from Bedrock.  
   [PR #6051](https://github.com/earendil-works/pi/pull/6051)

10. **#6026 – Stabilise working status row**  
    Improves the TUI “working” indicator to avoid flickering and inaccurate scroll positions, directly related to the #5825 scroll‑jump bug.  
    [PR #6026](https://github.com/earendil-works/pi/pull/6026)

## Feature Request Trends
The most requested directions from the last day’s issues are:

- **Session & state management** – Users want `--profile` isolation (#3966), shorter session‑naming commands (#6046), and the ability to create named sessions inline.
- **Extension API expansion** – Requests for exposing `navigateTree()` to all extension contexts (#5932), durable HITL intercepts for tool calls (#5901), and allowing steering messages to opt out of waking the agent (#5895) indicate a push toward more programmable agents.
- **Provider diversity** – Multiple new provider integrations (Scaleway #6165, fixing Xiaomi pricing #6138, updating MiniMax context window #6171) and support for `image_url` content type without base64 conversion (#6151) show appetite for broader platform coverage.
- **Editor UX** – Redo support (#6183) and configurable padding (#6115) reflect a desire for a more polished terminal experience.
- **Enterprise features** – Admin‑level settings enforced via `/etc` or `%ProgramData%` (#6159) and the ability to have multiple skill invocations in a single message (#6180) point to growing adoption in team environments.

## Developer Pain Points
Recurring frustrations and high‑frequency issues observed in the last day:

- **Agent loop / stuck “working”** – Issues #4338 and #6158 describe the agent repeating the same tool calls without progress. This remains a top reliability concern.
- **Streaming reliability** – Crashes from `ECONNRESET` (#6133), unsupported OpenAI retries (#6019), and misplaced “working” status (#5825) erode trust in long‑running sessions.
- **Provider error opacity** – Despite the fix in #5832, many older integrations still swallow meaningful error details; developers struggle to debug gateway misconfigurations.
- **Pricing mismatches** – Hardcoded model pricing (Xiaomi, MiniMax) leads to inaccurate cost tracking and user confusion.
- **Tool call bugs** – Empty tool results mislabelled as images (#6103), tool list not refreshed between turns (#6162), and base64 corruption on certain providers (Kimi Coding #6164) create integration headaches.
- **Missing undo/redo** – The lack of redo was highlighted again (#6183); the PR merging today addresses this gap.
- **Multi‑client orchestration** – The `rpc_stream` UI handler being silently overwritten (#6177) indicates scaling issues for headless or multi‑window setups.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-06-30

## Today's Highlights

The daemon session archive feature (Issue #6057) and sessionless workspace remember (PR #5884) signal a major push toward making the daemon a first-class session management platform—users can now offload old chats without deletion and enqueue memory tasks without creating a visible session. On the CLI side, the new `classifyAllShell` setting (PR #6040) and inline one-shot model override in `/model` (PR #6022) give power users finer-grained control over shell execution and model routing. Meanwhile, a critical OOM fix (PR #6018) addresses two long-standing paths where full-history clones could exhaust memory on large sessions.

---

## Releases

**v0.19.3-nightly.20260630.e00fe6a27** — Nightly release with continued documentation refresh for the daemon (wave 2). The change log was truncated but includes docs updates from PRs #5954 and an unfinished `feat(core)` entry for configurable auto-* behavior.  
[View release](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.3-nightly.20260630.e00fe6a27)

---

## Hot Issues

*Picking the most impactful issues from the 5 updated in the last 24h (insufficient volume for 10; below are all 5).*

1. **#6049 — `timeout: 0` causes immediate request timeout**  
   *[Bug, P2]* Setting `generationConfig.timeout` to 0 results in an immediate timeout error, which is surprising given that `0` often means "no timeout" in other tools. 5 comments, no upvotes yet—community is likely affected silently.  
   [Issue](https://github.com/QwenLM/qwen-code/issues/6049)

2. **#6057 — Add daemon session archive support**  
   *[Feature, P2, roadmap/session-management]* Requests a lightweight way to hide/retire sessions without deleting history. Archived sessions cannot be loaded or resumed but are restorable. This fills a long-standing gap for users with hundreds of old chats. 1 comment.  
   [Issue](https://github.com/QwenLM/qwen-code/issues/6057)

3. **#5759 — `ui.history.collapsePreviewCount` to show last N messages on resume**  
   *[Feature, closed]* When `collapseOnResume` hides all history, users lose context. This feature would show the last N messages as a preview. Accepted and closed after 3 comments—community eagerly awaited this.  
   [Issue](https://github.com/QwenLM/qwen-code/issues/5759)

4. **#6024 — Exclude line numbers from clipboard when copying code blocks**  
   *[Feature, P3, welcome-pr]* Copying code blocks in the TUI includes gutter line numbers, which breaks pasting into editors. Simple fix; 1 comment, marked as `welcome-pr` for new contributors.  
   [Issue](https://github.com/QwenLM/qwen-code/issues/6024)

5. **#4940 — Add `deniedMcpServers` policy**  
   *[Feature, P2, closed]* Complements the existing allowlist by adding a deny-list for MCP servers. Useful for admins who want a default-allow setup but need to block specific servers. 1 comment, closed after 20 days.  
   [Issue](https://github.com/QwenLM/qwen-code/issues/4940)

---

## Key PR Progress

1. **PR #6018 — Avoid full-history clones in OOM-prone paths**  
   Replaces deep cloning of full chat history in API error reporting and forked-agent cache snapshots with compact summaries and boundary guards. Critical for users with large sessions.  
   [PR](https://github.com/QwenLM/qwen-code/pull/6018)

2. **PR #6021 — Handle ACP `read_file` for managed local paths**  
   Preserves local read behavior for skill instructions, temp outputs, subagent transcripts, and other managed roots when the workspace boundary rejects server-side reads.  
   [PR](https://github.com/QwenLM/qwen-code/pull/6021)

3. **PR #5902 — QQ bot streaming improvements**  
   Replaces BlockStreamer coalescing with 2-second idle flush, removes a 2000-character limit, adds 5-minute TTL to passive reply tracking, and fixes markdown table detection.  
   [PR](https://github.com/QwenLM/qwen-code/pull/5902)

4. **PR #5884 — Sessionless workspace remember**  
   Adds a daemon API to enqueue hidden managed-memory tasks without creating a user-visible session. Advertises workspace memory capability and returns in-memory task IDs.  
   [PR](https://github.com/QwenLM/qwen-code/pull/5884)

5. **PR #6027 — Sanitize subagent result tags**  
   Removes internal `<analysis>` blocks from subagent results before feeding them back to the parent agent. Raw transcripts still preserve model output.  
   [PR](https://github.com/QwenLM/qwen-code/pull/6027)

6. **PR #6022 — Inline one-shot model override in `/model`**  
   `/model <model-id> <prompt>` runs the trailing prompt on a different model for one turn, then reverts automatically. Includes tool-call continuations.  
   [PR](https://github.com/QwenLM/qwen-code/pull/6022)

7. **PR #6031 — Daemon-managed channel worker for `--channel`**  
   Implements the PR2 channel worker path: `qwen serve --channel <name>` starts an out-of-process worker connected to the current server instance.  
   [PR](https://github.com/QwenLM/qwen-code/pull/6031)

8. **PR #5999 — Replace all emoji with Unicode text symbols in TUI**  
   Completes the emoji cleanup across all TUI rendering paths, replacing width-2 emoji with width-1 Unicode glyphs (e.g., `💡`→`∴`, `✅`→`✓`). Follows #5787/#5788.  
   [PR](https://github.com/QwenLM/qwen-code/pull/5999)

9. **PR #6005 — Queue prompts while turns are running in web shell**  
   Adds daemon-backed FIFO prompt queuing so messages submitted during an active turn are accepted client-side and executed in order. Includes controls to reorder/cancel queued prompts.  
   [PR](https://github.com/QwenLM/qwen-code/pull/6005)

10. **PR #6044 — Tabbed Settings dialog with Status and Stats tabs**  
    Reworks `/settings` into a tabbed view (Settings / Status / Stats) with a search box on the Settings tab and keyboard navigation across three regions.  
    [PR](https://github.com/QwenLM/qwen-code/pull/6044)

---

## Feature Request Trends

- **Session lifecycle management** dominates: archive/restore (Issue #6057), collapse preview (Issue #5759), and sessionless operations (PR #5884) all point to a need for more sophisticated session handling beyond simple create/delete.
- **MCP server policy** continues to evolve: after `allowedMcpServers`, the community now wants `deniedMcpServers` (Issue #4940) for mixed allow/deny setups—common in enterprise deployments with third-party MCP integrations.
- **Daemon as a platform**: Several PRs (sessionless remember, channel workers, session archive) are positioning the daemon as an extensible backend, not just a CLI companion. Expect more daemon-first features.

---

## Developer Pain Points

- **Surprising default behavior**: Issue #6049 (`timeout: 0` → immediate timeout) violates the principle of least surprise. Similar issues have occurred with other `generationConfig` defaults—community would benefit from explicit documentation or a validation warning.
- **Clipboard pollution**: Issue #6024 (line numbers in copied code) is a small but frequent annoyance affecting daily TUI usage. Marked as `welcome-pr`, likely simple enough for new contributors.
- **OOM on large sessions**: PR #6018 addresses two OOM-prone paths caused by full-history deep clones. This suggests the codebase has scaling issues with long conversations—affects power users with complex multi-turn agent workflows.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-06-30

## Today’s Highlights
The v0.8.66 release cycle is in full swing, with a flurry of last‑minute blocker fixes merged today to handle multi‑sub‑agent fanout stalls, TUI layout overflows, and blocking lock contention. Meanwhile, two high‑severity bugs were opened — one describing permanent session corruption after long tool output, another detailing MCP OAuth authentication UX failures. The team also landed features for environment variable expansion in MCP configs and wildcard tool disallow patterns, along with documentation syncs for the Xiaomi MiMo UltraSpeed model.

## Releases
No new versions were published in the last 24 hours.

## Hot Issues (Top 10 by community attention)

1. **[#1177 – Input cache hit rate too low](https://github.com/Hmbown/CodeWhale/issues/1177)**  
   *trytsomile* reports that compared to DeepSeek‑Reasonix (95%+), CodeWhale’s input cache hit rate is abysmal. 24 comments, no upvotes. Community consensus: this is a critical performance regression needing immediate investigation.

2. **[#1120 – Cache hit problems persist](https://github.com/Hmbown/CodeWhale/issues/1120)**  
   *pmsleepcheck* continues the cache discussion, suspecting the `input_cache_miss` bug may not be fully resolved. 21 comments. Suggests deeper root‑cause analysis.

3. **[#743 – Token consumption increased dramatically](https://github.com/Hmbown/CodeWhale/issues/743)**  
   *YaYII* reports burning 400M tokens in half a day. 13 comments. Users are alarmed by the spike in request density and call for optimised conversation context.

4. **[#3461 – MCP duplicate server instance lifecycle](https://github.com/Hmbown/CodeWhale/issues/3461)**  
   *stream2stream* finds that CodeWhale spawns **two MCP server processes** from a single `mcp.json` entry, wasting RAM and causing orphan processes. 9 comments. Marked as closed after today’s triage, but the issue highlights lingering MCP reliability concerns.

5. **[#3800 – Release gate: multi sub‑agent fanout freeze](https://github.com/Hmbown/CodeWhale/issues/3800)**  
   *Hmbown* opened this as a v0.8.66 blocker: launching ~20 agents freezes the TUI. 2 comments. The parent issue for six sub‑fixes that all shipped today. Critical for release stability.

6. **[#3821 – Session permanently damaged after long tool output/approval timeout](https://github.com/Hmbown/CodeWhale/issues/3821)**  
   *bevis-wong* reports a high‑severity bug where an approval dialog timeout locks the agent permanently, forcing a hard restart. 1 comment (today). No workaround known.

7. **[#3819 – MCP OAuth authentication UX issues](https://github.com/Hmbown/CodeWhale/issues/3819)**  
   *bevis-wong* describes stale tokens, silent errors, and foreground login timeouts when using OAuth‑protected MCP servers. 1 comment. Community: “This makes MCP setup painful for enterprise use.”

8. **[#3807 – Ship Hotbar hidden by default](https://github.com/Hmbown/CodeWhale/issues/3807)**  
   *Hmbown* closes this with a product decision: fresh installs will not show Hotbar until explicit opt‑in. 1 comment. Helps keep v0.8.66 focused on reliability.

9. **[#3799 – Fix TUI modal and text overflow systemically](https://github.com/Hmbown/CodeWhale/issues/3799)**  
   *Hmbown* identifies layout failures across modals – text clipping, bottom action rows hidden. 1 comment. Merged PR #3814 addresses the approval prompt case.

10. **[#1186 – Add typed persistent permission rules](https://github.com/Hmbown/CodeWhale/issues/1186)**  
    *greyfreedom* proposes scoped permission rules (tool name, command prefix, path pattern) with allow/deny/ask decisions. 10 comments. Long‑requested feature for granular security.

## Key PR Progress (Top 10)

1. **[#3825 – Expand ${VAR} env placeholders in MCP stdio config](https://github.com/Hmbown/CodeWhale/pull/3825)**  
   *h3c-hexin* allows secret injection via environment variables in MCP server configurations, avoiding hard‑coded API keys. Important for secure MCP setups.

2. **[#3824 – Support wildcard disallowed tool prefixes](https://github.com/Hmbown/CodeWhale/pull/3824)**  
   *h3c-hexin* adds `disallowed_tools` matching by wildcard prefixes (e.g., `nordic-*`), enabling blanket hiding of entire MCP server toolsets.

3. **[#3823 – Suppress background console windows on Windows](https://github.com/Hmbown/CodeWhale/pull/3823)**  
   *h3c-hexin* fixes flickering console windows for every child process spawn, improving TUI UX on Windows.

4. **[#3822 – Prefer exact binary release assets](https://github.com/Hmbown/CodeWhale/pull/3822)**  
   *LI-Jialu* updates the self‑update mechanism to prefer platform‑specific binary assets before falling back to archives.

5. **[#3818 – Expand active tool run summaries](https://github.com/Hmbown/CodeWhale/pull/3818)**  
   *cyq1017* ensures in‑flight tool runs appear in dense expansion views, fixing a transcript edge case.

6. **[#3820 – Sync Xiaomi MiMo Token Plan docs](https://github.com/Hmbown/CodeWhale/pull/3820)**  
   *nightt5879* documents the `mimo-v2.5-pro-ultraspeed` model and Token Plan key names, aligning docs with code.

7. **[#3817 – Preserve standing YOLO authority for runtime continuations](https://github.com/Hmbown/CodeWhale/pull/3817)**  
   *Hmbown* fixes YOLO mode not being respected during agent handoffs or continuations, preventing unnecessary approval prompts.

8. **[#3816 – Persist sub‑agent state off the manager write‑lock hot path](https://github.com/Hmbown/CodeWhale/pull/3816)**  
   *Hmbown* splits JSON serialisation out of the write lock, reducing contention under high fanout. Closes #3805.

9. **[#3814 – Keep approval controls visible inline](https://github.com/Hmbown/CodeWhale/pull/3814)**  
   *Hmbown* re‑renders approval prompts as a bottom‑docked bar, fixing clipping on short terminals. Closes #3799.

10. **[#3813 – Use nonblocking send for ListSubAgents refresh](https://github.com/Hmbown/CodeWhale/pull/3813)**  
    *Hmbown* converts blocking `.send().await` to `try_send`, preventing TUI event‑loop stalls during sub‑agent status storms. Closes #3802.

## Feature Request Trends
- **Hotbar as an optional, customizable surface** – many issues (e.g., #2061, #3389, #3731) focus on making the MMO‑style quick‑action bar hidden by default, with setup wizard and localization coming in v0.8.68.
- **Persistent permission rules** – requests for typed, scoped allow/deny/ask decisions (#1186) indicate a need for fine‑grained execution policy.
- **Setup wizard and user constitution** – v0.8.67 will introduce a first‑run wizard and `/constitution` for personalisation (#3402, #3412). Community wants guided onboarding.
- **Cache optimisation** – the two cache‑hit issues (#1177, #1120) show strong demand for parity with DeepSeek‑Reasonix’s 95%+ hit rate.
- **Model‑specific context policies** – #2693 proposes per‑provider harness strategies, reflecting a desire to tune system prompts per model (e.g., DeepSeek V4 vs. Xiaomi MiMo).

## Developer Pain Points
- **MCP reliability** – duplicate server processes (#3461), OAuth UX failures (#3819), and blocked approval dialogs (#3821) make MCP integration brittle.
- **Session crashes due to long operations** – large file processing (#1425) and approval timeouts (#3821) cause unrecoverable freezes, requiring restarts.
- **Blocking I/O in async paths** – the v0.8.66 blocker series (#3800–#3805) highlights systemic lock contention in sub‑agent management, shell manager, and event channels.
- **Token consumption spikes** – issue #743 complains of 400M tokens in half a day, suggesting excessive context or looped requests.
- **Windows console window flicker** – PR #3823 addresses a long‑standing annoyance where every child process spawns a console window.
- **YOLO mode not respected** – #3790 showed that YOLO still prompts for push/PR approvals, undermining the trust‑mode promise.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*