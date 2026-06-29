# AI CLI Tools Community Digest 2026-06-29

> Generated: 2026-06-29 14:39 UTC | Tools covered: 9

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
**Date:** 2026-06-29 | **Scope:** 8 major AI CLI tools

---

## 1. Ecosystem Overview

The AI CLI tools landscape on June 29, 2026 shows a market in active maturation, with all eight tools handling high-volume issue tracking and delivering regular fixes. The dominant themes are **session reliability**, **cost management**, and **agent autonomy**—developers are demanding tools that don't silently fail, waste tokens, or lose context mid-session. A notable rate-limit cost shock (10–20× increase at OpenAI Codex) has sparked industry-wide concern about unpredictable API pricing, while session-recovery fixes across Gemini CLI and CodeWhale signal that **persistence and crash resilience** are becoming table stakes. MCP (Model Context Protocol) integration is a common friction point, with authentication, startup blocking, and permission persistence issues appearing across nearly every tool. The ecosystem is converging on shared features (deferred MCP loading, inline model switching, cost-aware compaction) while differentiating on IDE integration depth, enterprise configuration models, and provider ecosystem breadth.

---

## 2. Activity Comparison

| Tool | Hot Issues (24h) | PRs Updated (24h) | Release Status | Community Engagement Signal |
|------|-----------------|-------------------|----------------|----------------------------|
| **Claude Code** | 10 | 3 | No release | High (114👍 top request; 47 comments on connection bug) |
| **OpenAI Codex** | 10 | 10 | **rust-v0.142.4** | Very High (338👍 rate-limit issue; 197 comments) |
| **Gemini CLI** | 10 | 10 | v0.51.0-nightly | Moderate (8👍 top issue; active session-recovery PRs) |
| **Copilot CLI** | 10 | 1 | **v1.0.66-2** | Low (2👍 top issue; 1 PR trivial) |
| **Kimi Code** | 2 | 0 | No release | Very Low (1👍 top issue; 15 comments on 6-month bug) |
| **OpenCode** | 10 | 10 | No release | High (9👍 CPU bug; 10 active PRs) |
| **Pi** | 10 | 8 | No release | High (30👍 connection reliability; 72 comments) |
| **Qwen Code** | 10 | 10 | No release | Moderate (7 comments top issue; rapid PR throughput) |
| **CodeWhale (DeepSeek TUI)** | 10 | 10 | Pre-v0.8.66 | Moderate (release-blocker fixes dominate) |

**Key observation:** OpenAI Codex leads in community raw engagement (rate-limit crisis), while Claude Code, OpenCode, and Pi show the highest sustained issue depth. Kimi Code is effectively stagnant. Copilot CLI has the most stable release cadence but least community interaction.

---

## 3. Shared Feature Directions

### Session Reliability & Persistence
- **Gemini CLI, Claude Code, OpenCode, Copilot CLI, CodeWhale** – All have open issues about lost sessions, corrupted metadata, or resume failures.
- **Common ask:** Deterministic session recovery, graceful corruption handling, and visible resume status.
- **Solution patterns:** Gemini's closed PRs (#27904, #27912) on JSONL metadata recovery; CodeWhale's legacy session fallback (#3724).

### Cost Management & Token Transparency
- **OpenAI Codex (#28879):** 10–20× rate-limit cost jump sparking demand for per-turn cost dashboards.
- **Claude Code (#63903):** Memory preamble cannot be disabled, wasting 11–16k tokens of context.
- **Qwen Code (#5942, #5956):** Prompt-cache misses inflate costs; dedicated compaction model requested.
- **Pi (#6083):** Cache misses with third-party providers burning session limits.
- **Trend:** Users want configurable cheap models for compaction (Qwen #6019), visible token consumption, and controllable preamble suppression.

### TUI/UX Flexibility & Input Handling
- **Claude Code (#23134):** 114👍 for disabling paste-text collapse.
- **OpenCode (#34198):** TUI paste corruption on Windows.
- **Pi (#5825):** Streaming markdown forces scroll-to-bottom (36 comments).
- **Copilot CLI (#3972):** Raw mouse characters display on first load.
- **Qwen Code (#6011):** Mouse support being added.
- **Common pattern:** Users need terminal UIs that behave like modern editors—no unwanted scrolling, paste previews, scrollback management, and consistent keyboard/mouse interaction.

### Agent Autonomy & Sub-agent Reliability
- **Claude Code (#47180):** Cowork scheduled tasks ignore "Always allow" permissions.
- **Gemini CLI (#22323):** Subagents falsely report GOAL success on turn-limit hits.
- **OpenCode (#2945):** Automatic session compaction destroys working context.
- **Qwen Code (#6023):** Subagent tags leak into parent output.
- **Pi (#6158):** Repeated tool calls loop without interruption.
- **Shared problem:** Agents cannot be trusted to self-regulate—they ignore configuration, report false success, and waste resources.

### MCP Tool Ecosystem & Authentication
- **Claude Code (#65036):** MCP OAuth tokens not auto-refreshing.
- **OpenAI Codex (#30509, #30500):** PRs to prevent MCP startup from blocking reviews.
- **Copilot CLI (#3958):** Windows MCP startup regression with `.bat` commands.
- **Pi (#5871, #6093):** Hardcoded Anthropic OAuth token detection breaks scoped keys.
- **Qwen Code (#6004):** MCP installation crashes with OOM.
- **Insight:** MCP is a critical integration point, but every tool has unresolved pain in authentication, startup latency, or platform compatibility.

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | Copilot CLI | OpenCode | Pi | Qwen Code | CodeWhale |
|-----------|-------------|--------------|------------|-------------|----------|-----|-----------|-----------|
| **Primary Focus** | Deep IDE/TUI integration | Rate-limit & cost transparency | Session recovery & agent orchestration | Enterprise/GitHub integration | V2 API migration & performance | Provider diversity & TUI polish | Cost-aware architecture & feature velocity | Release stabilization & localization |
| **Target User** | Power developers, long sessions | Plus/Pro plan users | Google Cloud/GCP developers | GitHub Enterprise orgs | Open-source ecosystem | Multi-provider tinkerers | Cost-sensitive teams | Localization-first users |
| **Differentiator** | Memory system (controversial) | Highest community engagement | AST-aware code navigation | Org-level settings (server-managed) | Deferred MCP loading | Profile-based state isolation | Configurable compaction model | Multi-locale support (9 tracked) |
| **Top Pain Point** | Connection errors on Windows | Rate-limit cost spikes | Agent hangs & false success | Silent failures (web_fetch, session_store) | CPU waste during retry | TUI scroll-to-bottom | OOM during MCP install | Release-gate CI reliability |
| **Release Velocity** | Low (no release today) | Moderate (patch release) | High (nightly builds) | High (patch release) | Moderate (no release, active PRs) | Low (no release today) | High (10 PRs/day) | Pre-release (blocker fixes) |

**Key differentiators:**
- **OpenAI Codex** has the strongest community voice (338👍 on a single issue) but is also the most disrupted by API provider changes.
- **Gemini CLI** is the most aggressive on session recovery—multiple PRs today on corruption, metadata, and signal forwarding.
- **Copilot CLI** is the most enterprise-focused (org settings, plugin coexistence) but has the weakest community engagement.
- **Qwen Code** has the fastest PR throughput and most feature-dense roadmap, but at the cost of OOM crashes and regression risk.
- **CodeWhale (DeepSeek TUI)** is the only tool prioritizing multi-locale support and has the most structured release process (explicit release-blocker tracking).

---

## 5. Community Momentum & Maturity

### High Momentum (Rapid iteration, growing engagement)
- **Qwen Code** – 10 PRs in 24h, addressing both bugs and features. High feature velocity with cost-aware design. Best positioned for price-sensitive developers.
- **OpenCode** – Accelerating V2 API migration with 10 active PRs. Performance regressions are being addressed proactively (off-thread diff computation). Clear roadmap.
- **CodeWhale (DeepSeek TUI)** – Pre-release intensity with 10 bug-fix PRs. Structured release management (release-blocker labels, CI gate fixes). Localization matrix signals maturity.

### Established but Turbulent
- **OpenAI Codex** – Highest raw engagement but in crisis mode over rate-limit costs. The 338👍 issue indicates a trust problem. Community is vocal but tool stability is suffering.
- **Claude Code** – Strong TUI/UX community but no release today and 3 PRs only. Long-standing bugs (Windows connection, memory preamble) unresolved for months. Mature but slow.

### Maturing Steadily
- **Gemini CLI** – Heavy session-recovery investment today. Agent reliability issues persist (false GOAL success, hangs). Good PR throughput but many bugs remain open.
- **Copilot CLI** – Stable release cadence (v1.0.66-2) but minimal community PR activity (1 trivial PR). Enterprise focus may limit open-source engagement.
- **Pi** – Moderate activity with 8 PRs. Connection reliability (72 comments) and TUI UX (36 comments) are the main drags. No release today suggests a slower pace.

### Low Activity/Stagnant
- **Kimi Code** – Only 2 issues updated, 0 PRs. Critical 6-month bug (#640) unresolved. Minimal community engagement (1👍). Risk of abandonment.

---

## 6. Trend Signals for Developers

### 1. Cost Transparency is Becoming a Hard Requirement
- OpenAI Codex's 10–20× rate-limit jump is a warning: **API pricing volatility is the top risk for AI CLI tools**. Developers are demanding per-turn cost dashboards, configurable cheap models for compaction (Qwen #5956), and deterministic rate-limit resets (Codex #9508).
- **Action:** Evaluate tools with explicit cost management features (Qwen's compaction model flag, Claude's memory preamble toggle—even if buggy).

### 2. Session Persistence is the New Reliability Baseline
- Losing a session mid-build (OpenCode #2945, Gemini #27368, Claude #72270) is the #1 workflow killer. Tools that handle crashes, corruption, and resume gracefully will win trust.
- **Signal:** Gemini CLI's 4 PRs today on session recovery (missing projectHash, corrupt metadata lines, ENOSPC handling) set a standard others need to match.

### 3. MCP Ecosystem Maturation is Painful but Inevitable
- Every tool has MCP auth, startup, or platform issues. The solution pattern is emerging: **deferred loading** (OpenCode #34368, Codex #30509), **background initialization** (Codex #30509), and **configurable allow/exclude patterns** (Qwen #6012).
- **Practical impact:** If you use MCP extensively, expect daily friction with OAuth token refresh (Claude #65036) or startup blocking (Codex).

### 4. TUI UX Convergence Toward Editor Standards
- Users universally dislike forced scrolling (Pi #5825), paste corruption (OpenCode #34198), collapsed input (Claude #23134), and non-standard keybindings (Kimi #2479).
- **Winners:** Qwen Code (adding mouse support #6011) and CodeWhale (localization + hotbar customization) are investing in UX. Losers: tools that treat TUI as an afterthought.

### 5. Agent Trust Deficit is Growing
- Subagents reporting false success (Gemini #22323), ignoring configuration (Gemini #21968), and leaking internal tags (Qwen #6023) erode trust in autonomous modes.
- **Recommendation:** Prefer tools with transparent agent logging (Codex's new agent communication logging PR #30516) and explicit "safe mode" flags (Qwen #4883).

### 6. Localization is an Emerging Differentiator
- CodeWhale's localization matrix (9 locales), Pi's Devanagari corruption (#6124), and Qwen's multilingual daemon channels (#6010) signal that non-English support is becoming a competitive edge—especially for Asian and European developer markets.

### 7. Provider Diversity is Both Strength and Fragility
- Pi and Qwen Code support the widest range of providers (Anthropic, Bedrock, Xiaomi MiMo, OpenAI, etc.), but each integration introduces auth bugs, pricing mismatches, and cache inefficiencies.
- **Advice:** If you use a non-major provider, expect more friction. Qwen Code and Pi are the best bets for multi-provider workflows.

---

## Summary Recommendations

| Decision Context | Recommended Tool | Why |
|-----------------|-----------------|-----|
| **Cost-sensitive teams** | **Qwen Code** | Configurable compaction model, cost-aware architecture, fastest bug-fix PRs |
| **Enterprise/GitHub orgs** | **Copilot CLI** | Server-managed settings, plugin coexistence, stable release cadence |
| **Multi-provider power users** | **Pi or Qwen Code** | Widest provider support; Pi for TUI polish, Qwen for feature velocity |
| **Localization/Non-English** | **CodeWhale (DeepSeek TUI)** | Only tool with explicit localization matrix and multi-locale PRs in flight |
| **Session-heavy long workflows** | **Gemini CLI** | Aggressive session recovery investment (4 PRs today), most transparent about corruption fixes |
| **Deep IDE integration** | **Claude Code** | Best TUI feature requests (paste preview, toast notifications) despite slow release cadence |
| **Avoid if:** | **Kimi Code** | 6-month critical bug unresolved, 0 PR activity, minimal engagement |

**Bottom line:** No tool is fully mature yet. The ecosystem is in a "feature race with reliability debt" phase—choose based on which pain points you can tolerate, and monitor the rate-limit cost trajectory as your primary risk factor.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
*Data snapshot: 2026-06-29 | Source: github.com/anthropics/skills*

---

## 1. Top Skills Ranking

### 1. fix(skill-creator): run_eval.py always reports 0% recall
**PR #1298** — *Author: MartinCajiao | Status: Open*
[GitHub Link](https://github.com/anthropics/skills/pull/1298)

Fixes the evaluation pipeline (`run_eval.py`, `run_loop.py`, `improve_description.py`) where the description-optimization loop was optimizing against noise. Root cause: the eval artifact wasn't installed as a real skill. Also addresses Windows stream reading, trigger detection, and parallel worker bugs. This is the community's **most-discussed PR** because it directly fixes the broken loop that prevents skill description improvements. Community reproduction reports in issues #556 and #1169 confirm 10+ independent reproductions of 0% recall.

### 2. Add document-typography skill
**PR #514** — *Author: PGTBoos | Status: Open*
[GitHub Link](https://github.com/anthropics/skills/pull/514)

Provides typographic quality control for AI-generated documents: orphan word wrap, widow paragraphs, and numbering misalignment. Broadly applicable since these issues affect every document Claude generates. The PR highlights a universal pain point users rarely articulate but encounter frequently.

### 3. fix(pdf): correct case-sensitive file references in SKILL.md
**PR #538** — *Author: Lubrsy706 | Status: Open*
[GitHub Link](https://github.com/anthropics/skills/pull/538)

Fixes 8 case-sensitivity mismatches in `skills/pdf/SKILL.md` where `REFERENCE.md` and `FORMS.md` were referenced in uppercase but the actual files are lowercase. This breaks on case-sensitive filesystems (Linux/macOS). A small but critical correctness fix.

### 4. Add ODT skill — OpenDocument text creation and template filling
**PR #486** — *Author: GitHubNewbie0 | Status: Open*
[GitHub Link](https://github.com/anthropics/skills/pull/486)

Comprehensive skill for creating, filling, reading, and converting OpenDocument Format files (.odt, .ods). Covers LibreOffice document production and HTML conversion. Addresses the open-source office format gap in the skills collection.

### 5. Improve frontend-design skill clarity and actionability
**PR #210** — *Author: justinwetch | Status: Open*
[GitHub Link](https://github.com/anthropics/skills/pull/210)

Revises the frontend-design skill to ensure every instruction is executable within a single conversation. Focuses on specificity and behavioral steering rather than vague guidelines. A quality-of-life improvement for an existing skill.

### 6. Add skill-quality-analyzer and skill-security-analyzer to marketplace
**PR #83** — *Author: eovidiu | Status: Open*
[GitHub Link](https://github.com/anthropics/skills/pull/83)

Two meta-skills: **skill-quality-analyzer** evaluates skills across five dimensions (Structure & Documentation, etc.), and **skill-security-analyzer** provides security auditing. These are meta-tools for skill authors, not end-user skills.

### 7. fix(docx): prevent tracked change w:id collision with existing bookmarks
**PR #541** — *Author: Lubrsy706 | Status: Open*
[GitHub Link](https://github.com/anthropics/skills/pull/541)

Fixes document corruption when the DOCX skill adds tracked changes to documents that already contain bookmarks. The root cause is a shared `w:id` ID space in OOXML — hardcoded low IDs collide with existing bookmark IDs.

### 8. fix(skill-creator): warn on unquoted description with YAML special characters
**PR #539** — *Author: Lubrsy706 | Status: Open*
[GitHub Link](https://github.com/anthropics/skills/pull/539)

Pre-parse validation in `quick_validate.py` to detect unquoted `description` fields containing `:`. Prevents silent YAML parsing failures where descriptions are truncated or split into multiple keys. Duplicate of PR #361 by Mr-Neutr0n, indicating strong community demand for this fix.

---

## 2. Community Demand Trends

### 🏢 Enterprise & Security Concerns
**Issue #492** (31 comments, 2 👍) — Security: Community skills under `anthropic/` namespace enables trust boundary abuse. **Most commented issue in the repo.** Raises concerns about impersonation of official Anthropic skills, leading to potential permission escalation. The 2-month ongoing discussion signals deep community concern.

**Issue #228** (14 comments, 7 👍) — Enable org-wide skill sharing in Claude.ai. Users want a shared skill library instead of manual `.skill` file distribution via Slack/Teams. **Highest 👍 count** among open issues.

### 🔧 Developer Tooling Pain Points
**Issue #556** (12 comments, 7 👍) — `run_eval.py`: claude -p never triggers skills/commands (0% trigger rate). Matches PR #1298. **The #1 blocking bug** for skill creators.

**Issue #202** (8 comments, 1 👍) — skill-creator should be updated to best practice. Critiques the current skill-creator as "developer documentation" rather than an operational skill.

### 🧠 Memory & Context Management
**Issue #1329** (6 comments) — Proposed compact-memory skill using symbolic notation for compact agent state. Addresses the long-running agent context overflow problem.

### 🛡️ Agent Governance
**Issue #412** (6 comments) — Proposed agent-governance skill: safety patterns for AI agent systems (policy enforcement, threat detection, trust scoring, audit trails). No equivalent skill exists in the collection.

### 📦 Packaging & Distribution
**Issue #189** (6 comments, 9 👍) — document-skills and example-skills plugins install identical content, causing duplicate skills. **Highest 👍 count** among issues overall.

---

## 3. High-Potential Pending Skills

| PR | Skill | Author | Last Updated | Status |
|----|-------|--------|-------------|--------|
| [#1298](https://github.com/anthropics/skills/pull/1298) | fix(skill-creator): run_eval.py recall | MartinCajiao | 2026-06-23 | **Most active** |
| [#1302](https://github.com/anthropics/skills/pull/1302) | Add color-expert skill | meodai | 2026-06-12 | Recently created |
| [#1367](https://github.com/anthropics/skills/pull/1367) | self-audit: four-dimension reasoning gate | YuhaoLin2005 | 2026-06-29 | **Most recent** |
| [#723](https://github.com/anthropics/skills/pull/723) | Add testing-patterns skill | 4444J99 | 2026-04-21 | Comprehensive |
| [#147](https://github.com/anthropics/skills/pull/147) | Add codebase-inventory-audit skill | p19dixon | 2026-02-04 | Orphan detection |
| [#181](https://github.com/anthropics/skills/pull/181) | Add SAP-RPT-1-OSS predictor | amitlals | 2026-03-16 | Enterprise ML |
| [#154](https://github.com/anthropics/skills/pull/154) | Add shodh-memory skill (persistent context) | varun29ankuS | 2026-03-03 | Memory system |

**Notable**: PR #1367 (self-audit, created just 1 day ago) already has active discussion. PR #1302 (color-expert) is a well-structured, self-contained skill covering 10+ color naming systems. PR #1298 remains the highest-priority fix due to blocking all skill description optimization.

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is for fixing the skill-creator evaluation loop (0% recall bug)** — three independent PRs (#1298, #1099, #1050) and two issues (#556, #1169) all tackle the same root cause, making reliable skill description optimization the single largest blocker in the ecosystem today, followed closely by a growing demand for **security safeguards, enterprise sharing infrastructure, and memory/context management skills**.

---

# Claude Code Community Digest — 2026-06-29

## Today's Highlights

No new releases were published in the last 24 hours, but the community remains highly active with several critical issues gaining traction. Top attention goes to a long-standing API connection error on Windows (#4297, 47 comments, 25 👍) and a CPU-exhaustion bug where scheduled-task runners leak idle sub-agents (#72270, filed today). The most upvoted feature request (114 👍) asks for an option to disable paste-text collapse in the TUI, reflecting a desire for more flexible input review.

---

## Releases

*None in the last 24 hours.*

---

## Hot Issues

1. **#4297 – [BUG] API Error (Connection error.)** — Windows + WARP terminal; connection failures with automatic retries. 47 comments, 25 👍.  
   https://github.com/anthropics/claude-code/issues/4297

2. **#23134 – Feature Request: Option to disable paste text collapse in input field** — Users want to see full pasted content before submitting. 45 comments, 114 👍 (highest upvoted).  
   https://github.com/anthropics/claude-code/issues/23134

3. **#47180 – [BUG] Cowork scheduled tasks ignore "Always allow" folder/tool permissions** — Permissions prompts re-appear every run on macOS, breaking autonomy. 23 comments, 29 👍.  
   https://github.com/anthropics/claude-code/issues/47180

4. **#63903 – [BUG] `autoMemoryEnabled=false` does not suppress memory preamble** — System prompt still loads ~11k–16k of memory-related text, wasting context. 20 comments.  
   https://github.com/anthropics/claude-code/issues/63903

5. **#69415 – [BUG] API Error: Connection closed mid-response** — Makes Claude Code unusable on VSCode/WSL. 8 comments, 33 👍.  
   https://github.com/anthropics/claude-code/issues/69415

6. **#65036 – [BUG] MCP OAuth: Claude doesn't auto-refresh access tokens** — Daily "Connection expired" despite valid refresh token. 4 comments, 15 👍.  
   https://github.com/anthropics/claude-code/issues/65036

7. **#57230 – [VSCode Extension] Add native system/toast notifications** — Users miss permission requests and task completions when not watching the terminal. 5 comments, 15 👍.  
   https://github.com/anthropics/claude-code/issues/57230

8. **#29580 – [BUG] `DISABLE_TELEMETRY=1` breaks remote-control** — Misleading "not yet enabled" error instead of graceful fallback. 14 comments, 15 👍.  
   https://github.com/anthropics/claude-code/issues/29580

9. **#7387 – [BUG] Shell script fails due to crazy escaping** — Intermittent errors on macOS/Bedrock when tool commands contain special characters. 19 comments, 10 👍.  
   https://github.com/anthropics/claude-code/issues/7387

10. **#72270 – [BUG] Scheduled-task runner leaks stream-json sessions → CPU exhaustion** — 136 leaked agent processes, 85% CPU, filed today. 1 comment.  
    https://github.com/anthropics/claude-code/issues/72270

---

## Key PR Progress

1. **#72264 – docs(examples/hooks): note Bash `tool_input` also exposes `run_in_background`/`description`/`timeout`** (Open) — Small documentation improvement to clarify hook payload fields. Merged? No, still open.  
   https://github.com/anthropics/claude-code/pull/72264

2. **#62315 – Fix hookify event filtering in pre/post hooks** (Closed) — Merged fix for event-filtering logic; not yet in a release.  
   https://github.com/anthropics/claude-code/pull/62315

3. **#41447 – feat: open source claude code ✨** (Open, not merged) — A community PR jokingly referencing several closed issues. No substantive changes; not expected to be merged.  
   https://github.com/anthropics/claude-code/pull/41447

*Note: Only 3 PRs were updated in the last 24 hours; no significant feature PRs are in progress.*

---

## Feature Request Trends

- **TUI & Input Flexibility** – Top request: disable paste-text collapse. Others ask for week-start configuration in heatmap (#55915), click-to-select only when window is focused (#72273), and warning when paste obscures content.
- **IDE Integration** – Strong demand for a real JetBrains plugin (#47166) and native VSCode toast notifications (#57230) to catch attention without constant terminal watching.
- **Graceful Usage Limits** – Several requests for hooks when tokens run out (#55945) and better mid-session handling of 429 errors (#56978) instead of session lock.
- **Permissions & Autonomy** – Users want persistent "always allow" settings honored in Cowork tasks, and permission prompts that never use numbers for "No" to avoid accidental denials (#55854).

---

## Developer Pain Points

- **Connection & API Stability** – Recurring connection-closed errors on WSL and Windows, plus SSL certificate expiration (#71663) and request timeouts (#70128). Developers report these make the tool unusable for sustained tasks.
- **Memory & Cost Overhead** – The memory preamble cannot be fully disabled (#63903), burning context and cost. Sub-agents sometimes enter infinite loops (#72080) consuming excessive tokens.
- **Cowork & MCP Friction** – Cowork permissions don't persist across runs (#47180). MCP OAuth tokens fail to auto-refresh (#65036), requiring daily manual reconnection. GitHub connector regressions (#71542) block all repository access.
- **Platform Quirks** – macOS sleep/wake triggers 401 auth (#60104). German Neo 2 keyboard layout causes Cmd+V to start new session (#67522). Fullscreen TUI breaks login URL copy (#70857). Telemetry-disable flag breaks remote control (#29580).

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-06-29

## Today’s Highlights
A severe rate-limit cost jump (10–20× since June 16) is draining Plus and Pro budgets in 2–3 prompts, sparking 197 comments and 338 upvotes on the leading issue. Concurrently, a wave of “model at capacity” errors is blocking access to GPT-5.5 and GPT-5.4 on multiple platforms. On the positive side, a major SQLite log write amplification bug (estimated 640 TB/year) was fixed in three merged PRs, and new pull requests aim to prevent MCP startup from blocking review workflows.

## Releases
- **rust-v0.142.4** — Chores only; no user-facing changes.  
  Full Changelog: [openai/codex compare/rust-v0.142.3...rust-v0.142.4](https://github.com/openai/codex/compare/rust-v0.142.3...rust-v0.142.4)

## Hot Issues
1. **[#28879 – Rate-limit cost per token jumped ~10–20x since June 16](https://github.com/openai/codex/issues/28879)**  
   *197 comments, 338 👍* — Budget drains in 2–3 prompts on `gpt-5.5` with Plus plan. Session logs confirm limit-% consumed per token increased drastically. Community strongly impacted; still open.

2. **[#28224 – SQLite feedback logs can write ~640 TB/year](https://github.com/openai/codex/issues/28224)**  
   *104 comments, 406 👍* — Three PRs merged (released 0.142.0) reduced log volume by ~85%. Issue now closed by author. High visibility and relief.

3. **[#9508 – Make weekly limit reset deterministic](https://github.com/openai/codex/issues/9508)**  
   *43 comments, 31 👍* — Users on Pro plan report unpredictable weekly reset timing. Long-standing enhancement request (opened Jan 2026).

4. **[#28507 – “Selected model is at capacity” on GPT-5.5](https://github.com/openai/codex/issues/28507)**  
   *21 comments, 13 👍* — Persistent capacity errors on Pro 5x subscription, Windows. Multiple duplicate issues appearing today (#30546, #30547, #30575, #30577, #30579).

5. **[#30364 – GPT-5.5 reasoning token clustering at 516/1034/1552](https://github.com/openai/codex/issues/30364)**  
   *18 comments, 22 👍* — Disproportionate counts at fixed token boundaries. Users suspect degraded performance on complex tasks.

6. **[#2909 – Support for multi-root workspaces](https://github.com/openai/codex/issues/2909)**  
   *16 comments, 137 👍* — VS Code extension feature request for Multi-root Workspaces. High upvote count, actively discussed since Aug 2025.

7. **[#26951 – Codex IDE extension stuck over VS Code Remote-SSH](https://github.com/openai/codex/issues/26951)**  
   *11 comments* — Extension loads indefinitely while CLI works. Affects Windows local + Ubuntu remote.

8. **[#30009 – apply_patch fails with Windows sandbox error](https://github.com/openai/codex/issues/30009)**  
   *10 comments* — File edits through `apply_patch` fail on Windows sandbox. Pro plan.

9. **[#26896 – Windows sandbox CreateProcessAsUserW failed: 5](https://github.com/openai/codex/issues/26896)**  
   *8 comments, 3 👍* — Unfixable on Windows 11 Enterprise (Dell). Sandbox cannot start.

10. **[#23999 – Codex Desktop sidebar chat history disappears](https://github.com/openai/codex/issues/23999)**  
    *8 comments, 2 👍* — Latest updates do not restore hidden chats. Affects Pro plan, macOS.

## Key PR Progress
1. **[#30509 – Allow review while MCP startup runs in background](https://github.com/openai/codex/pull/30509)**  
   *Open* — Prevents background MCP initialization from blocking `/review` input. Directly addresses common startup frustration.

2. **[#30500 – Run reviews without unfinished MCP servers](https://github.com/openai/codex/pull/30500)**  
   *Open* — Review sessions skip MCP servers still starting, removing wait on OAuth discovery.

3. **[#30467 – Treat “max” as a first-class reasoning effort](https://github.com/openai/codex/pull/30467)**  
   *Open* – Normalizes `max` effort display for GPT-5.6 Bedrock catalog; fixes rendering as lowercase.

4. **[#30493 – Add configurable multi-agent mode hint text](https://github.com/openai/codex/pull/30493)**  
   *Open* — Allows deployments to override built-in delegation policies with a stable hint.

5. **[#30516 – Add explicit agent communication logging](https://github.com/openai/codex/pull/30516)**  
   *Open* — Introduces uniform start/end logging for agent communication events, aiding debugging.

6. **[#30488 – Show reset details in redemption picker](https://github.com/openai/codex/pull/30488)**  
   *Open* — Displays available reset credits, expiration dates, and which credit will be consumed.

7. **[#30504 – Replace rollback with session forks](https://github.com/openai/codex/pull/30504)**  
   *Open* — Moves TUI transcript time travel from deprecated `thread/rollback` to non-destructive session forks.

8. **[#27723 – Preserve user goal evidence in approval review](https://github.com/openai/codex/pull/27723)**  
   *Closed* — Separates user-provided goals from continuation metadata in Guardian review, improving audit clarity.

9. **[#28131 – Refresh SSH agent for app-server proxy](https://github.com/openai/codex/pull/28131)**  
   *Closed* — Adds `--forward-ssh-agent` flag to keep live forwarded agent after original SSH session exits.

10. **[#30487 – Fall back from unsupported reasoning effort](https://github.com/openai/codex/pull/30487)**  
    *Closed* — Prevents cross-thread messages from breaking inference when effort exceeds model capabilities.

## Feature Request Trends
- **Deterministic and transparent rate limits** — Users want predictable weekly reset times and clear visibility into token consumption (#9508, #28879).
- **Workspace flexibility** — Multi-root workspace support in VS Code extension (#2909) and per-selection follow-up comments in the desktop app (#22677).
- **Session UX enhancements** — Community asking for a `/recap` command and `/btw` alias similar to Claude Code (#18884), plus session forks to avoid destructive rollback (#30504).
- **Automation model selection** — The `model` field in `automation.toml` is being ignored, forcing `gpt-5.5` regardless of configuration (#30439).
- **Non-blocking MCP startup** — MCP servers should not block core workflows like thread creation or `/review` (#29321, #29376); addressed by PRs #30509 and #30500.

## Developer Pain Points
- **Rate-limit unpredictability** — Sudden 10–20× cost increase and recurring “model at capacity” errors dominate the bug tracker, affecting Plus and Pro users across platforms.
- **Windows sandbox instability** — Several unresolved issues: `CreateProcessAsUserW` fails on Enterprise, `apply_patch` errors, blank tray icon, and RTL text rendering problems.
- **MCP startup blocking** — Failing or slow MCP servers freeze new conversations and reviews, causing timeouts.
- **Desktop app UX regressions** — Sidebar chat history disappearing, spellcheck showing no suggestions on Windows, and freeze due to stale local rollout.
- **Bundled toolchain conflicts** — Codex uses its own `pnpm` instead of the host toolchain (#30440), breaking build scripts.
- **Reasoning token clipping** — Fixed token boundaries (516/1034/1552) may indicate truncation or degraded performance on complex tasks (#30364).

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest – 2026-06-29

## 1. Today's Highlights

A significant batch of session-recovery fixes landed today, addressing corruption, missing metadata, and race conditions that caused lost conversations and broken resume flows. On the security front, two PRs closed a critical bug where the workspace-trust dialog could hide the actual hooks that would execute. Meanwhile, the community continues to report persistent agent reliability issues—hangs, false success on turn limits, and subagents ignoring configuration—though many of these remain open.

## 2. Releases

- **[v0.51.0-nightly.20260629.gae0a3aa7b](https://github.com/google-gemini/gemini-cli/releases/tag/v0.51.0-nightly.20260629.gae0a3aa7b)** – Standard nightly build. No user-facing release notes beyond the automated changelog diff from yesterday's nightly.

## 3. Hot Issues

1. **#22323 – Subagent recovery after MAX_TURNS reported as GOAL success**  
   *8 comments, 2 👍*  
   A `codebase_investigator` subagent falsely reports `status: "success"` and `Termination Reason: "GOAL"` even though it hit the turn limit before doing any analysis. This masks real failures and undermines trust in agent reporting.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/22323)

2. **#24353 – Robust component level evaluations**  
   *7 comments*  
   An epic tracking growth from 76 to more behavioral eval tests. The team runs these across 6 Gemini models, but the infrastructure needs hardening for consistent, reproducible results.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/24353)

3. **#22745 – Assess impact of AST-aware file reads, search, and mapping**  
   *7 comments, 1 👍*  
   Investigates whether AST-aware tools can reduce token noise, improve method-boundary precision, and speed up codebase navigation. A potential step-change for large-repo performance.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/22745)

4. **#21409 – Generalist agent hangs forever**  
   *7 comments, 8 👍*  
   Simple operations like folder creation hang indefinitely when the generalist agent is invoked. Users report waiting up to an hour. Workaround: explicitly disable subagents. High community impact.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/21409)

5. **#27368 – Latest chat session permanently lost after `gemini --resume`**  
   *6 comments, 1 👍*  
   Using `--resume` causes the most recent session to vanish from `/chat` list. Data appears lost or the session index gets corrupted. Priority P1 bug affecting daily workflows.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/27368)

6. **#21968 – Gemini does not use skills and sub-agents enough**  
   *6 comments*  
   Custom skills with clear descriptions (e.g., "gradle", "git") are ignored unless the user explicitly instructs the model. The agent fails to self-select relevant tooling.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/21968)

7. **#26525 – Add deterministic redaction and reduce Auto Memory logging**  
   *5 comments*  
   Auto Memory sends transcript content to an extraction model before redacting secrets, and can log existing skill content. A security concern that needs pre-redaction and better logging hygiene.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/26525)

8. **#26522 – Stop Auto Memory from retrying low-signal sessions indefinitely**  
   *5 comments*  
   Sessions remain in an "unprocessed" state if the extraction agent decides they're low-signal, causing them to be re-surfaced repeatedly. Resource waste and user confusion.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/26522)

9. **#28036 – Resumed sessions repeatedly stop after partial execution**  
   *4 comments*  
   Reproducible on v0.47.0 and v0.49.0. Tasks start normally after `--resume` but stop before completion with no quota or context-limit errors. Users must manually re-prompt.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/28036)

10. **#25166 – Shell command execution gets stuck with "Waiting input"**  
    *4 comments, 3 👍*  
    Simple CLI commands hang after finishing, still showing "Awaiting user input" and preventing progress. Frequent enough to be a P1 with high user frustration.  
    [Issue](https://github.com/google-gemini/gemini-cli/issues/25166)

## 4. Key PR Progress

1. **#28053 – [OPEN] Defensive path resolution for @-reference files**  
   Fixes a critical production bug where `read_file`/`replace`/`write_file` fail with "File not found" when the model passes paths like `@policies/new-policies.txt`. Also resolves macOS test failures.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28053)

2. **#28200 – [CLOSED] Sanitize trailing periods from URLs in auth error messages**  
   Auth error messages containing URLs like `https://goo.gle/...` could have trailing sentence periods appended, breaking terminal hyperlink detection.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28200)

3. **#28202 – [CLOSED] Forward SIGINT/SIGTERM/SIGQUIT to child process during relaunch**  
   Pressing Ctrl+C during an update/relaunch killed the parent but left the child orphaned. Now signals are properly forwarded.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28202)

4. **#28201 – [CLOSED] Remove double-wrapping of VS Code disposables**  
   Extra parentheses in the VS Code IDE Companion extension caused a subscription leak by double-wrapping disposables. Small but important for long-running editor sessions.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28201)

5. **#27915 – [CLOSED] Trust dialog discloses the hook shape that never runs**  
   Security fix: the workspace-trust dialog showed the **inverse** of the hooks that actually execute. A project could hide a dangerous `SessionStart` hook from the user.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/27915)

6. **#27914 – [CLOSED] Don't offer to resume a session that wasn't saved**  
   When a write hits `ENOSPC`, the chat recorder disables itself but the exit summary still printed a "resume" command. Now properly omitted.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/27914)

7. **#27916 – [CLOSED] Validate GCP project ID format and prevent alias extraction in memory**  
   Prevents Auto Memory from storing GCP display names/aliases, which caused 403 and CONSUMER_INVALID errors in later sessions.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/27916)

8. **#27910 – [CLOSED] Bound web search tool latency with 120s timeout**  
   Adds a local 120s timeout around the `google_web_search` call. Aborts the underlying `generateContent` request and returns a clear tool error instead of hanging.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/27910)

9. **#27905 – [CLOSED] Keep recreated session files loadable after deletion**  
   Fixes a bug where manually deleting a session file while the process is running would cause `appendFileSync` to recreate it with incomplete state. Now properly re-loadable.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/27905)

10. **#27904 & #27912 – [CLOSED] Load JSONL sessions when projectHash is missing; recover corrupt metadata lines**  
    A two-PR fix stack: sessions with missing `projectHash` now fall back correctly instead of corrupting, and sessions with a corrupt first metadata line can be recovered from the remaining JSONL data.  
    [#27904](https://github.com/google-gemini/gemini-cli/pull/27904) · [#27912](https://github.com/google-gemini/gemini-cli/pull/27912)

## 5. Feature Request Trends

The most-requested capability directions this week center on **agent self-awareness and tool selection**. Users want the CLI to autonomously decide which skills, sub-agents, and file-reading strategies to deploy without explicit instruction, including AST-aware code navigation for more precise method-boundary reads. There is also strong demand for **agent trajectory transparency**—making subagent internal steps visible via `/chat share` and in bug reports. On the evaluation side, the community is pushing for **component-level behavioral evals** that are more granular than end-to-end tests, and for **deterministic session recovery** that handles corruption gracefully. Finally, **browser agent resilience** continues to be a theme, with requests for automatic session takeover, lock recovery, and respect for `settings.json` overrides.

## 6. Developer Pain Points

- **Agent hangs and false progress** – The generalist agent and browser subagent both hang indefinitely on simple tasks. More concerning: subagent termination reasons (e.g., `GOAL`) are reported as success even when the agent hit turn limits or failed internally, hiding real errors.
- **Session and data corruption** – `--resume` frequently loses entire chat sessions, lists sessions missing, and corrupt metadata blocks rendering sessions unloadable. Today's batch of recovery PRs addresses some of these, but root-cause fixes are still in progress.
- **Subagents ignoring configuration** – Users report that `settings.json` overrides for agents (e.g., `maxTurns`, `disabled`) are silently ignored, and that subagents started launching even when explicitly disabled since v0.33.0.
- **Auto Memory resource waste** – The background extraction agent retries low-signal sessions indefinitely, and invalid patches are silently skipped rather than quarantined. Users report confusion and wasted tokens.
- **Model destructive behavior** – The model occasionally uses destructive commands (`git reset`, `--force`) when safer alternatives exist. Users want the agent to understand the consequences of modifying databases and version control state.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-06-29

## Today's Highlights
A new patch release (v1.0.66-2) focuses on plugin coexistence, local CLI user settings access, and better LSP debugging. Two critical session bugs (#2364, #3600) were closed, but a fresh regression on Windows MCP startup (#3958) and a persistent `web_fetch` failure (#3948) are drawing community attention. The single PR updated is a trivial file rename.

---

## Releases

**v1.0.66-2** – [View release](https://github.com/github/copilot-cli/releases/tag/v1.0.66-2)
- Added: Skills with the same name from different plugins can now coexist.
- Added: Integrations can read and write CLI user settings.
- Added: View LSP server logs via `/lsp logs` and `read_agent`.
- Added: Prompt to install `gh` CLI when missing in GitHub repositories.
- Added: GitHub attachment variants to prompt rendering.

No other releases in the last 24h.

---

## Hot Issues (10 Noteworthy)

1. **[#2364 – Copilot Agent session runs indefinitely, cannot stop (CLOSED, Critical)**](https://github.com/github/copilot-cli/issues/2364)  
   *4 comments, 2 👍*  
   A critical bug where Coding Agent sessions in org repos get stuck after the initial plan, producing no commits or fixes. Community frustration due to lack of a dedicated place to report. Resolved in the latest release? Closed as of today.

2. **[#3909 – Enterprise/org server-managed settings for local CLI](https://github.com/github/copilot-cli/issues/3909)**  
   *3 comments*  
   Org admins cannot centrally push env variables to local Copilot CLI installs. Currently only cloud environments (Codespaces) support this. Request is gaining traction.

3. **[#3600 – Orphaned sessions running for two months (CLOSED, Critical)](https://github.com/github/copilot-cli/issues/3600)**  
   *3 comments*  
   Users unable to remove stale sessions. Closed today, likely fixed in v1.0.66-2.

4. **[#2654 – `session_store_sql` silently returns empty when sync is local](https://github.com/github/copilot-cli/issues/2654)**  
   *2 comments, 1 👍*  
   The tool remains callable but returns 0 rows with no indication to agents. Causes silent failures in workflows relying on session history.

5. **[#3948 – `web_fetch` always fails with TypeError: fetch failed](https://github.com/github/copilot-cli/issues/3948)**  
   *1 comment*  
   Every `web_fetch` call fails, even though auth and proxy settings seem correct. Affects all users, no workaround reported.

6. **[#3958 – Windows: stdio MCP servers with .bat/.cmd commands fail (regression)](https://github.com/github/copilot-cli/issues/3958)**  
   *1 comment*  
   v1.0.66 breaks MCP server startup on Windows when the command is a batch file with args. Immediate regression from v1.0.65.

7. **[#3904 – CloudQueryError blocks `/chronicle standup` despite local fallback](https://github.com/github/copilot-cli/issues/3904)**  
   *1 comment*  
   Session store cloud errors prevent a core feature even when local data is available. DuckDB-style timestamp predicates cause full failure.

8. **[#3972 – UI displays stream of mouse movement characters](https://github.com/github/copilot-cli/issues/3972)**  
   *0 comments*  
   On first load, the terminal UI shows raw mouse movement characters instead of the interface. Newly filed; no response yet.

9. **[#3962 – v1.0.65 not working after update](https://github.com/github/copilot-cli/issues/3962)**  
   *1 comment*  
   Copilot CLI stuck at "Working" after any command input. User likely encountering a rendering or session init issue.

10. **[#2619 – Billed $2.9 during trial period](https://github.com/github/copilot-cli/issues/2619)**  
    *2 comments*  
    Support ticket submitted but unanswered. Community notes similar billing confusion on free/ trial plans.

---

## Key PR Progress

Only **one PR** was updated in the last 24h:

- **[#3968 – Rename changelog.md to changelog.md (CLOSED)](https://github.com/github/copilot-cli/pull/3968)**  
  *Author: creepyalissa-coder*  
  A trivial no-op rename. No other active PRs were updated.

PR activity is low, suggesting the team is focused on bug fixes and release stabilization.

---

## Feature Request Trends

1. **Session management enhancements**  
   Multiple requests for session retention/expiration display (#3963), plan status indicators (#3969), user‑defined tags (#3970), and a full file‑tree browser for repository‑backed sessions (#3971).

2. **Enterprise / org‑level configuration**  
   Central management of env variables for local CLI (#3909) is a recurring ask from admins.

3. **MCP plugin coexistence warnings**  
   When MCP servers share names across plugins, the last‑installed wins silently. Users want a warning or priority display (#3893).

4. **Billing / quota transparency**  
   Free‑plan quota not resetting (#2340) and unexpected trial charges (#2619) suggest a need for clearer usage dashboards or notifications.

---

## Developer Pain Points

- **Silent failures**: `session_store_sql` returns empty with no agent awareness (#2654) and `web_fetch` fails without actionable diagnostics (#3948).
- **Platform regressions**: Windows MCP startup broken in the latest patch (#3958) and disappearing CLI installations on Linux (#3967).
- **Terminal rendering glitches**: Stream of mouse characters (#3972), ghost characters after deletion (#3959), and inability to scroll with trackpad on macOS (#3957).
- **Session permanence issues**: Orphaned sessions that cannot be removed (#3600) and cloud query errors blocking features even when local fallback exists (#3904).

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

**Kimi Code CLI Community Digest – 2026-06-29**

## Today's Highlights

Activity remains low today with no new releases or pull requests. A long-standing bug (Issue #640) continues to attract community attention—users report the CLI entering an infinite loop while reading a single file on custom Anthropic endpoints. A new enhancement request (Issue #2479) highlights poor mobile usability and inconsistent Enter/Shift+Enter behavior across desktop and phone.

## Releases

No new releases in the last 24 hours. The latest stable version remains **0.76**.

## Hot Issues

Only two issues were updated in the last 24 hours. Both are highlighted below.

[#640 – **Kimi CLI stuck in reading one file again and again and stuck in a loop**](https://github.com/MoonshotAI/kimi-cli/issues/640)  
*Author: isbafatima90-arch | Created 2026-01-19 | Updated 2026-06-28 | Comments: 15 | 👍: 1*

**Why it matters:** This is a persistent, critical bug affecting users who configure custom Anthropic endpoints (e.g., model `mimo-v2-flash`) via `config.toml`. The CLI loops indefinitely, reading the same file repeatedly, making the tool unusable. Despite being open for six months and having 15 comments, the issue remains unresolved—a significant pain point for advanced users who rely on custom models. Community reaction shows frustration, as the single thumbs-up indicates limited engagement but the high comment count suggests active discussion.

[#2479 – **[Enhancement] Bad usage of return and enter for desktop and mobile**](https://github.com/MoonshotAI/kimi-cli/issues/2479)  
*Author: Dealazer | Created 2026-06-29 | Updated 2026-06-29 | Comments: 0 | 👍: 0*

**Why it matters:** Filed today, this request points to a core UX flaw: on mobile, pressing Enter sends the prompt immediately, making multi-line input nearly impossible. On desktop, creating a new line requires Shift+Enter, which is non-standard for most terminal-based tools. The issue has no comments yet, but it reflects a growing need for cross-platform input consistency—especially as developers use Kimi on phones for quick tasks.

## Key PR Progress

No pull requests were updated in the last 24 hours. There are no active PRs to report.

## Feature Request Trends

Based on the two active issues, the most requested feature directions are:

1. **Custom model stability** – Users integrating Kimi with non-standard API endpoints (Anthropic, custom providers) need robust error handling and prevention of infinite loops.
2. **Mobile-first input handling** – A strong desire to differentiate between “send message” and “new line” across devices, including proper keyboard mapping for mobile browsers and terminals.

## Developer Pain Points

- **Infinite loop on custom endpoints** – Issue #640 remains the top frustration. Users cannot safely use Kimi CLI with their own model deployments, and the lack of a fix for over six months erodes trust.
- **Inconsistent keyboard behavior** – The Enter vs. Shift+Enter confusion (Issue #2479) makes the CLI difficult to use in mobile browsers and for multi-line prompts on desktop. This is a basic UX expectation that should be addressed early.

*Generated from GitHub data for MoonshotAI/kimi-cli on 2026-06-29.*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-29

## Today's Highlights
The V2 API migration is accelerating with a flurry of new feature requests and tracking issues for porting session, shell, and MCP surfaces to the new `@opencode-ai/client`. User-facing stability remains a hot topic: three critical bugs—desktop freezes on large diffs, TUI paste corruption, and plan mode crashes—have active fix PRs in review. No new releases landed today, but the steady stream of issues and PRs signals a release may be imminent.

## Releases
No new versions published in the last 24 hours.

## Hot Issues
1. **[#19466 – "opencode is using CPU for doing nothing!"](https://github.com/anomalyco/opencode/issues/19466)**  
   Awaiting API rate limits uses ~50% of a single i9-14900 core. 11 comments, 9 👍. Community frustration over wasted resources during retry delays.

2. **[#2945 – Session automatically compacted, destroying working context](https://github.com/anomalyco/opencode/issues/2945)**  
   Agent forgets all session details mid-build. 6 comments, 4 👍. A long-standing pain point that breaks long workflows.

3. **[#17173 – OpenCode Go Performance is abysmal](https://github.com/anomalyco/opencode/issues/17173)**  
   Agent startup and tool calls extremely slow with glm-5. 5 comments, 4 👍. High demand for performance improvements.

4. **[#34222 – GitHub Copilot MAI-Code-1-Flash not accessible](https://github.com/anomalyco/opencode/issues/34222)**  
   Model returns “not accessible via /chat/completions”. 4 comments, 2 👍. Blocked for Copilot Enterprise users.

5. **[#34442 – Windows Desktop installer broken offline](https://github.com/anomalyco/opencode/issues/34442)**  
   Ripgrep not bundled; core tools (`grep`, `glob`, `skill`) fail without internet. 1 comment, 0 👍. Critical for air-gapped environments.

6. **[#34437 – Desktop renderer freezes on large file diffs](https://github.com/anomalyco/opencode/issues/34437)**  
   `execEditLength` runs synchronously on UI thread, causing seconds-long freezes on C++ projects. 1 comment. A fix PR is already open.

7. **[#34198 – TUI rendering corruption on paste (Windows)](https://github.com/anomalyco/opencode/issues/34198)**  
   Since v1.16.2, pasting content visually corrupts the terminal window. Layout duplicates, data intact. 1 comment.

8. **[#34427 – Plan mode breaks CLI](https://github.com/anomalyco/opencode/issues/34427)**  
   Switching to plan mode crashes the CLI with a visual error. 1 comment. Blocks a core workflow.

9. **[#34426 – OpenCode Desktop v1.17.11 GUI broken in menus](https://github.com/anomalyco/opencode/issues/34426)**  
   Sidebar toggle missing, app unusable after update. 2 comments. Significant UI regression.

10. **[#34438 – GitHub Copilot Claude effort levels in /models](https://github.com/anomalyco/opencode/issues/34438)**  
    OpenCode should trust the model endpoint instead of hardcoding Claude effort filters. 2 comments. Important for pay-per-token parity.

## Key PR Progress
1. **[#34415 – fix(ui): prepare diffs off the render thread](https://github.com/anomalyco/opencode/pull/34415)**  
   Moves diff computation to a Web Worker, eliminating UI freezes on large files (fixes #34437).

2. **[#34414 – fix(app): avoid O(n²) dedup hang on large diff summaries](https://github.com/anomalyco/opencode/pull/34414)**  
   Replaces inefficient `result.some()` + `reduceRight` with `Set‑based` lookup, fixing renderer hangs reported in #28844.

3. **[#34440 – feat(llm): enforce request precedence](https://github.com/anomalyco/opencode/pull/34440)**  
   Resolves option merging order: route defaults → model defaults → call options. Adds tests for Anthropic max_tokens.

4. **[#34368 – feat(opencode): defer large MCP tool catalogs](https://github.com/anomalyco/opencode/pull/34368)**  
   Adds experimental deferred MCP tool search/call bridge to avoid loading all tools upfront. New session-level tool assembly.

5. **[#34416 – fix(opencode): preserve shell command variant](https://github.com/anomalyco/opencode/pull/34416)**  
   Threads the active model variant through `!cmd` shell requests, preserving thinking effort selection.

6. **[#33925 – feat(core): load native provider packages](https://github.com/anomalyco/opencode/pull/33925)**  
   Migrates config, catalog, and providers to flat schema; encodes legacy AI SDK providers as `aisdk:<package>`.

7. **[#34419 – feat(desktop): add setting to swap panel layout side](https://github.com/anomalyco/opencode/pull/34419)**  
   Toggle in Settings → Appearance to move the side panel (chat vs editor) to the opposite side (closes #16349).

8. **[#34441 – fix: preserve Bedrock DeepSeek model IDs](https://github.com/anomalyco/opencode/pull/34441)**  
   Prevents `deepseek.v3.2` from being mis‑classified as a generic cross‑region model ID.

9. **[#34431 – fix: exempt org issues from compliance close](https://github.com/anomalyco/opencode/pull/34431)**  
   Automated compliance closer now skips issues/PRs authored by OWNER/MEMBER roles, reducing noise.

10. **[#34385 – refactor(core): finish test layer node conversion](https://github.com/anomalyco/opencode/pull/34385)**  
    Converts remaining core service tests from manual `defaultLayer` composition to `LayerNode` graphs, improving testability.

## Feature Request Trends
- **V2 API migration**: The majority of new feature requests target exposing V2 protocol surfaces—session fork, diff, shell, command, MCP lifecycle, and @‑tagged file references. This is the clear roadmap focus.
- **Offline/air‑gapped support**: Requests for bundling `ripgrep` (#31734) and fixing the Windows Desktop installer (#34442) show strong demand for fully offline environments.
- **Desktop customization**: Users want adjustable font size/line height (#27684), searchable prompt history (Ctrl+R, #34406), and panel layout swaps (#34419).
- **Provider parity**: Issues around Claude effort levels (#34438) and Copilot endpoint compatibility (#34222) indicate pressure to treat all API providers equally.
- **Session context improvements**: Proposals for session‑scoped keyed context (#34380) and progressive `AGENTS.md` loading (#34341) aim to make agent memory more reliable.

## Developer Pain Points
- **Performance regressions** dominate: idle CPU usage (#19466), general Go slowness (#17173), desktop freezes on large diffs (#34437), and O(n²) renderer hangs (#34414).
- **Session reliability** is fragile: automatic compaction destroys context (#2945), plan mode crashes the CLI (#34427), and paste corrupts the TUI (#34198).
- **Offline/air‑gapped support** is broken on Windows: core tools fail without internet (#34442, #31734).
- **Desktop stability** suffers from frequent GUI regressions (sidebar missing, menu corruption #34426).
- **Provider integration** still has rough edges: Copilot Enterprise models blocked (#34222) and Bedrock model ID mangling (#34441).

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-06-29

## Today's Highlights

- **Connection reliability & streaming UX** dominate discussions: `openai-codex` TUI hangs (#4945) and forced scroll-to-bottom (#5825) are the most active bugs.
- **Anthropic and Bedrock auth issues** hit critical mass — multiple issues/PRs address hardcoded OAuth detection and token handling (#5871, #6093, #6161, #6148).
- **No release in the last 24 hours.** The community remains focused on provider reliability, TUI polish, and extension API refinements.

## Releases

*No new releases in the last 24 hours.*

---

## Hot Issues *(10 noteworthy)*

1. **[#4945 – openai-codex Connection Reliability Issues](https://earendil-works/pi/issues/4945)**  
   *72 comments, 30 👍*  
   `gpt-5.5` leaves the TUI stuck on `Working...` with no error; only Escape recovers. A long-standing pain point that has generated significant community discussion.  

2. **[#5825 – Streaming markdown forces scroll to bottom](https://earendil-works/pi/issues/5825)**  
   *36 comments*  
   When `clear on shrink` is enabled, fast streaming scrolls the user back down. A major UX annoyance for heavy readers.

3. **[#6083 – LLM cache not working with z.ai GLM coding plan](https://earendil-works/pi/issues/6083)**  
   *8 comments, 9 👍*  
   Cache misses burn session limits rapidly during multi‑tool tasks. Closed but highlights a broader concern about third‑party provider efficiency.

4. **[#5871 – Anthropic OAuth‑token detection hardcoded to `sk-ant-oat`](https://earendil-works/pi/issues/5871)**  
   *6 comments*  
   Scoped Anthropic keys don’t match the hardcoded prefix, breaking authentication. Community requests configurable detection.

5. **[#3454 – Zero padding still in‑scope?](https://earendil-works/pi/issues/3454)**  
   *5 comments, 2 👍*  
   A long‑standing feature request for configurable TUI padding. Recently re‑opened after a PR (#6115) was proposed but not merged.

6. **[#6093 – Scoped Anthropic API keys need necessary request params](https://earendil-work/pi/issues/6093)**  
   *5 comments*  
   Scoped keys (`sk-ant-api03-...`) are misidentified; related to #5871. Closed but root cause not fully resolved.

7. **[#6158 – Repeated tool calls can loop without interruption](https://earendil-works/pi/issues/6158)**  
   *3 comments*  
   Agent gets stuck in a loop (e.g., `ls` 6 times) instead of progressing. Closed as `no-action` but indicates need for better loop detection.

8. **[#6138 – Incorrect pricing for Xiaomi MiMo native models](https://earendil-works/pi/issues/6138)**  
   *3 comments*  
   Hardcoded prices don’t match official Xiaomi rates; potentially causing inaccurate cost tracking.

9. **[#6124 – Devanagari breaking the Pi harness](https://earendil-works/pi/issues/6124)**  
   *3 comments*  
   Non‑Latin scripts (e.g., `नेटवर्क`) corrupt the TUI display. Points to broader Unicode rendering gaps.

10. **[#6103 – OpenAI Responses API mislabels empty tool results as “(see attached image)”](https://earendil-works/pi/issues/6103)**  
    *2 comments*  
    Empty tool outputs are incorrectly labeled as image attachments. Fixed by PR #6156 shortly after reporting.

---

## Key PR Progress *(8 total – all listed)*

1. **[#6161 – fix(ai): map Bedrock apiKey auth to bearer token env](https://earendil-works/pulls/6161)**  
   *Closed* – Maps `apiKey` into `AWS_BEARER_TOKEN_BEDROCK` for Bedrock Converse API. Resolves a major auth mismatch.

2. **[#5832 – fix(ai): surface provider HTTP error body](https://earendil-works/pulls/5832)**  
   *Open, in progress* – Fixes opaque SDK error messages when behind a proxy. Wide impact.

3. **[#6026 – fix(tui): stabilize working status row](https://earendil-works/pulls/6026)**  
   *Open, in progress* – Targets the scroll‑to‑bottom behavior (#5825). Under active development.

4. **[#6156 – fix(ai): return empty string for empty tool results](https://earendil-works/pulls/6156)**  
   *Closed* – Direct fix for #6103 (OpenAI Responses API).

5. **[#6148 – fix(ai): support Anthropic bearer token env](https://earendil-works/pulls/6148)**  
   *Open, to discuss* – Attempts to resolve #5871 but raises interface design concerns.

6. **[#6074 – fix(coding-agent): avoid pre-prompt compaction continue](https://earendil-works/pulls/6074)**  
   *Closed* – Prevents unnecessary compaction on continuation.

7. **[#6078 – feat(coding-agent): add get_entries and get_tree RPC commands](https://earendil-works/pulls/6078)**  
   *Closed* – Exposes session management data for extensions.

8. **[#6115 – feat(coding-agent): add configurable chat padding](https://earendil-works/pulls/6115)**  
   *Open, to discuss* – Responds to long‑standing request (#3454) but notes TUI structural challenges.

---

## Feature Request Trends

- **Profile‑based state isolation** – Multiple proposals for `--profile` to keep work/personal/local setups separate (#3966, plus similar comments in other threads).
- **Extension API expansion** – Community wants `ctx.navigateTree()` exposed to all extension contexts (#5932), and better tool renderer type safety (#6098).
- **Non‑English support** – Sessions in languages like Hindi (#6124) and requests to generate compaction summaries in the conversation language (#6157).
- **Admin/enterprise settings** – A new call for system‑wide config files (`/etc`, `%ProgramData%`) for managed environments (#6159).
- **Provider diversity** – Requests to add Charm Hyper (#6042) and to better support Xiaomi MiMo, OpenCode Go, and others.

---

## Developer Pain Points

- **Connection reliability & streaming** – The TUI freezing (#4945) and uncontrolled scrolling (#5825) are the top‑voted pain points; users spend significant effort recovering from “working…” hangs.
- **Authentication confusion** – Hardcoded OAuth token patterns (#5871, #6093) and lack of configurable bearer‑token mapping for Bedrock (#6161) cause repeated provider failures.
- **Unicode/Localization gaps** – Devanagari breaking the TUI (#6124) and requests for non‑English summaries (#6157) show that i18n is an underserved area.
- **Tool loop detection** – Agent sometimes repeats the same tool call indefinitely (#6158), leading to wasted time and token burn.
- **Pricing inaccuracy** – Hardcoded model prices (e.g., Xiaomi MiMo, #6138) mislead users and cause frustration when billing differs from actual provider rates.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-06-29

## Today's Highlights

The past 24 hours saw a wave of critical memory and prompt‑cache fixes land against a backdrop of feature expansion. An MCP installation crash (#6004) triggered an investigation into OOM‑prone code paths, prompting a proactive avoidance of full‑history clones (#6018). At the same time, upstream prompt‑cache efficiency for Anthropic backends (#5942) and a dedicated compaction model setting (#5956 / #6019) signal growing community attention on cost awareness. On the UI side, the removal of file‑name display in tool‑use messages (#6014) sparked immediate negative feedback, while several open PRs aim to modernise the TUI interaction model (mouse support, prompt queuing, frozen transcript).

---

## Releases

**No new versions were published in the last 24 hours.**

---

## Hot Issues

| # | Issue & Link | Why it matters | Community reaction |
|---|--------------|----------------|-------------------|
| 1 | [#6004 – MCP installation crash](https://github.com/QwenLM/qwen-code/issues/6004) | CLI flash‑crashes during MCP server install with a heap‑out‑of‑memory GC trace. Blocks the MCP onboarding experience. | 7 comments; labelled `welcome‑pr` – immediate triage and PR #6018 directly addresses the cause. |
| 2 | [#5942 – Anthropic prompt‑cache misses](https://github.com/QwenLM/qwen-code/issues/5942) | Two independent request‑side cache problems inflate `cache_read` token cost compared to Claude Code on the same backend. | 4 comments; detailed analysis with expected cost delta. |
| 3 | [#5956 – Configurable compaction model](https://github.com/QwenLM/qwen-code/issues/5956) | Auto‑compact uses the active (often expensive) model to summarise context. Request to set a dedicated cheap model for compaction. | 3 comments; PR #6019 already implements a `--compaction` flag. |
| 4 | [#5967 – Inline model switching](https://github.com/QwenLM/qwen-code/issues/5967) | `/model <id> <prompt>` as a single‑turn override without a two‑step workflow. | 2 comments; PR #6022 implements it. |
| 5 | [#6014 – File‑name display removed](https://github.com/QwenLM/qwen-code/issues/6014) | New TUI version silently hides which files the agent read; regresses to generic “read 1 file”. | 2 comments; strong negative sentiment (“downgrade”). |
| 6 | [#6010 – Hot‑reloadable daemon channels](https://github.com/QwenLM/qwen-code/issues/6010) | `qwen serve` should allow live add/update/remove of DingTalk, WeChat, Telegram etc channels. | 2 comments; part of the broader `daemon/background-automation` roadmap. |
| 7 | [#6007 – GLM‑5.2 leaks thinking text](https://github.com/QwenLM/qwen-code/issues/6007) | Internal reasoning leaks to normal output when `max_tokens` is at the default 131072. | 2 comments; affects users behind OpenAI‑compatible providers. |
| 8 | [#6023 – Subagent tags leak into parent](https://github.com/QwenLM/qwen-code/issues/6023) | `<analysis>` / `<summary>` tags from subagent final results break daemon UI markdown rendering. | 1 comment; noted as “first seen in long daemon sessions”. |
| 9 | [#6020 – `read_file` returns `[object Object]`](https://github.com/QwenLM/qwen-code/issues/6020) | Surface error on ACP skill reads outside workspace boundary is useless. | 1 comment; PR #6021 already submitted to fix the root cause. |
| 10 | [#4883 – `--safe-mode` flag](https://github.com/QwenLM/qwen-code/issues/4883) | Request for a troubleshooting flag disabling all customisations. | 1 comment; PR #4943 was merged earlier this month. |

---

## Key PR Progress

| # | PR & Link | What it does |
|---|-----------|--------------|
| 1 | [#6005 – Web‑shell prompt queue](https://github.com/QwenLM/qwen-code/pull/6005) | Server‑side FIFO queue for web shell; queued prompts are visible and draggable before execution. |
| 2 | [#6018 – Avoid full‑history clones](https://github.com/QwenLM/qwen-code/pull/6018) | Prevents OOM in API error reporting and forked‑agent cache snapshots by compacting payloads instead. |
| 3 | [#6022 – Inline model override](https://github.com/QwenLM/qwen-code/pull/6022) | Implements `/model <id> <prompt>` one‑shot switch. Reverts to session model after one turn. |
| 4 | [#6019 – `/model --compaction`](https://github.com/QwenLM/qwen-code/pull/6019) | Adds a dedicated compaction model flag, solving #5956. |
| 5 | [#6011 – TUI mouse clicks & hover](https://github.com/QwenLM/qwen-code/pull/6011) | Enable mouse in alternate‑screen mode: select menus, permission dialogs, scroll. |
| 6 | [#6006 – Browser MCP tools by default](https://github.com/QwenLM/qwen-code/pull/6006) | Auto‑registers browser MCP path in `qwen serve` sessions unless disabled. |
| 7 | [#6021 – ACP `read_file` local roots](https://github.com/QwenLM/qwen-code/pull/6021) | Preserves local‑root reads (skill instructions, memory) when workspace boundary rejects – fixes #6020. |
| 8 | [#6012 – Glob patterns for MCP allow/exclude](https://github.com/QwenLM/qwen-code/pull/6012) | Replaces exact‑match `Array.includes` with `*` / `?` patterns. |
| 9 | [#6013 – Serve health responsiveness](https://github.com/QwenLM/qwen-code/pull/6013) | Defers runtime graph until after first `/health` probe so startup stays fast. |
| 10 | [#5991 – Autonomous `/loop`](https://github.com/QwenLM/qwen-code/pull/5991) | Bare `/loop` arms a self‑paced autonomous loop without requiring an explicit prompt. |

---

## Feature Request Trends

1. **Cost‑aware architecture** – Configurable compaction models (#5956), prompt‑cache efficiency (#5942), and the related suggestion of per‑turn cost dashboards reflect a growing need to manage cloud‑API spend at scale.

2. **Inline & multi‑turn commands** – `/model <id> <prompt>` (#5967), inline model override in #6022, and the autonomous `/loop` (#5991) point to demand for concise, single‑line interaction flows without entering/exiting sub‑modes.

3. **Daemon & extension‑channel expansion** – Hot‑reloadable channels (#6010), sessionless workspace memory (#5884), and resumable SSE streams (#5852) show the community pushing `qwen serve` toward a flexible, long‑running service.

4. **UI/UX modernisation** – Mouse support (#6011), frozen transcript (#5666), collapse preview counts (#5848), and the negative reaction to #6014 (hidden file names) all indicate that the terminal UI is under active ergonomic scrutiny.

5. **Safety & diagnostics** – The `--safe-mode` flag (#4883) and SSL‑ignore requests (#3535) suggest a need for troubleshooting‑friendly launch options.

---

## Developer Pain Points

| Pain Point | Signals |
|------------|---------|
| **Out‑of‑memory crashes** | #6004 (MCP install flash crash → GC failure), #6018 (preventative clone avoidance) |
| **Cost inflation from cache misses** | #5942 (Anthropic prompt‑cache misses), #5956 (expensive models used for compaction) |
| **Regressions in TUI information density** | #6014 (file names removed), #6023 (internal tags leaking into visible output) |
| **Poor error diagnostics** | #6020 (`[object Object]` surface error), #6007 (silent thinking‑text leak) |
| **Configuration inflexibility** | #5956 (no dedicated compaction model), #6012 (exact‑match MCP allow/exclude), #3535 (no SSL ignore flag) |

The overall sentiment is pragmatic: the community welcomes rapid feature delivery but is equally quick to flag regressions that reduce transparency or increase operating cost.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest – 2026-06-29

## Today's Highlights

The project (now branded as **CodeWhale**) is in the final push for the **v0.8.66 release**, with multiple release-blocker bugs being closed in the last 24 hours. Several CI/CD reliability issues were uncovered—most critically, release gates like `check:facts` and `check:docs` could pass even when stale, and the npm installer path ignored the `codewhaleBinaryVersion` field. At the same time, the rebrand migration from `~/.deepseek/` to `~/.codewhale/` gained crucial UX improvements (silent migration, session fallback, and `codewhale doctor` diagnostics). Community contributors stepped up with PRs for localization, startup speed, and a new provider integration.

## Releases

*No new releases in the last 24 hours.*

---

## Hot Issues

*(10 most noteworthy issues updated in the last 24h)*

1. **#3766 – Approval UI copy misleads users about session-scoped approval**  
   *[bug/release-blocker]* The approval modal says “always” but only grants approval for the current TUI session. A trust-boundary copy bug that must be fixed before v0.8.66 ships.  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/issues/3766)

2. **#3730 – Auto mode flags read-only commands as DESTRUCTIVE**  
   *[bug]* Running `codewhale --version` in Auto mode forces an approval prompt. Also the policy copy references “YOLO” instead of “Auto”. Auto mode is broken at its core.  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/issues/3730)

3. **#3768 – Accidental deletion of 0.8.52 changelog section**  
   *[bug/release-blocker]* The local 0.8.66 candidate diff removed the full `0.8.52` entry from `crates/tui/CHANGELOG.md`. Must be restored before release.  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/issues/3768)

4. **#3767 – Guard install-doc version snippets in release bump checks**  
   *[bug/release-blocker]* `prepare-release.sh` and `check-versions.sh` can pass while public install snippets still point at the previous release.  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/issues/3767)

5. **#3772 – `check:facts` doesn’t fail on unmapped ApiProvider variants**  
   *[bug/release-blocker]* New provider variants can be added to the Rust enum without being mapped on the website, and the CI gate won’t catch it.  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/issues/3772)

6. **#3771 – Web facts drift check regenerates before checking**  
   *[bug/release-blocker]* CI regenerates `facts.generated.ts` before comparing, so a stale committed file can pass the drift check.  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/issues/3771)

7. **#3770 – `check-docs` prints FAIL but exits PASS**  
   *[bug/release-blocker]* Stale install snippets cause a FAIL message, but the script exits 0, letting CI through.  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/issues/3770)

8. **#3769 – npm installer ignores `codewhaleBinaryVersion`**  
   *[bug/release-blocker]* The installer path resolves assets without consulting the dedicated version field, causing install-time drift.  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/issues/3769)

9. **#3724 – Sessions appear lost after upgrade (legacy fallback missing)**  
   *[bug/migration]* Old `~/.deepseek/sessions/` are not surfaced by CodeWhale. The read-path doesn’t use the centralized fallback logic.  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/issues/3724)

10. **#3757 – Launch startup is slow; needs profiling**  
    *[bug/v0.8.67]* Startup feels noticeably sluggish in local testing, especially after rebuild/install cycles.  
    📎 [GitHub](https://github.com/Hmbown/CodeWhale/issues/3757)

---

## Key PR Progress

*(10 important pull requests merged or opened in the last 24h)*

1. **#3785 – Localize Hotbar setup wizard**  
   Localizes wizard chrome, slot labels, help text, and validation errors. Enables non-English setup experience. *(by nightt5879)*  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/pull/3785)

2. **#3780 – Expose context compaction gates via config.toml**  
   Adds `[compaction].enabled` and `[seam_manager].enabled` switches, closing #3765. *(by nightt5879)*  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/pull/3780)

3. **#3777 – Fail `check:facts` on unmapped ApiProvider variants**  
   Maps `Openmodel` and `Sakana` providers; ensures CI catches provider drift. *(by Hmbown)*  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/pull/3777)

4. **#3779 – Guard public install/version snippets in release checks**  
   Extends `prepare-release.sh` to protect `--version` and `INSTALL.md` pointers. *(by Hmbown)*  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/pull/3779)

5. **#3778 – Fix `sync-changelog` accidentally dropping releases**  
   Prevents counting `[Unreleased]` as a section, closing #3768. *(by Hmbown)*  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/pull/3778)

6. **#3776 – Run web facts drift check before regenerating tracked facts**  
   Reorders CI steps so stale `facts.generated.ts` is caught. *(by Hmbown)*  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/pull/3776)

7. **#3775 – Make `check-docs` exit non-zero on stale install snippets**  
   Now calls `process.exit(1)` when install snippets are stale. *(by Hmbown)*  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/pull/3775)

8. **#3774 – Honor `codewhaleBinaryVersion` in npm installer**  
   Fixes asset resolution in `install.js` to respect the dedicated version field. *(by Hmbown)*  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/pull/3774)

9. **#3773 – Label session-scoped approval honestly**  
   Changes “always” to “approve for this session” in the UI, closing #3766. *(by Hmbown)*  
   📎 [GitHub](https://github.com/Hmbown/CodeWhale/pull/3773)

10. **#3763 – Define localization matrix with locale registry and drift checks**  
    Introduces `LOCALIZATION.md` and a locale registry (`ALL_LOCALES`) for 9 tracked locales. *(by idling11)*  
    📎 [GitHub](https://github.com/Hmbown/CodeWhale/pull/3763)

---

## Feature Request Trends

- **Multi-locale website & README parity** – Issues #3091, #3759, and PR #3763 show strong demand for non-English locales, especially Japanese, Vietnamese, and Russian. The project is actively building a localization matrix.
- **Hotbar customization & QA** – Several issues (#3731, #3758, #3759) target the Hotbar activation UX, terminal testing of `Alt-1..8` chords, and localized setup wizards.
- **New AI providers** – #3751 requests Neuralwatt (GLM 5.2, non-token pricing). PR #3781 adds an “OpenCode Zen” provider. The provider ecosystem is expanding.
- **Config persistence for GUI** – #3784 (PR) adds nested-table config persistence, enabling a GUI config panel.
- **Startup performance** – Multiple reports (#3757) request launch time profiling and optimization. PR #3761 defers non-critical cleanup to background threads.

---

## Developer Pain Points

- **Release-gate false positives** – `check:facts`, `check:docs`, and `check-versions` each had bugs allowing CI to pass when public data was stale. This is a reliability headache for maintainers.
- **Rebrand migration headaches** – Silent state relocation (#3726), lost sessions (#3724), and no fallback read-path have frustrated upgraders. The community needs clear migration notices and diagnostic commands (now addressed by #3727 and PR #3763).
- **Auto mode is broken** – It behaves identically to Agent mode, flags safe commands as destructive, and has misleading copy (#3730, #3733). The feature was pulled from 0.8.66 for redesign.
- **Approval UX confusion** – The “always” vs “session-scoped” discrepancy (#3766) damages user trust; it’s a critical fix that landed just in time for the release candidate.
- **WSL2 install issues** – Missing `pkg-config`/`libdbus-1-dev` dependencies on Windows Subsystem for Linux are a recurring pain for new users (#1816, PR #3755).

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*