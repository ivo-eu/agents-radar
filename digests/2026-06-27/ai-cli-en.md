# AI CLI Tools Community Digest 2026-06-27

> Generated: 2026-06-27 09:15 UTC | Tools covered: 9

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

# Cross-Tool AI CLI Ecosystem Comparison Report
**Date:** 2026-06-27

---

## 1. Ecosystem Overview

The AI CLI tools landscape is maturing rapidly, with seven major projects showing sustained development activity and community engagement. The dominant themes this week are **agent reliability** (scope creep, silent failures, state management), **cost transparency** (rate-limit spikes, credit exhaustion), and **platform parity** (Windows/Linux/macOS bugs and accessibility gaps). MCP (Model Context Protocol) compliance is emerging as a cross-cutting concern, with multiple tools racing to support the latest spec. While Claude Code and OpenAI Codex lead in raw community size, Gemini CLI, Qwen Code, and DeepSeek TUI stand out for rapid iteration on foundational agent safety and memory backends. The overall ecosystem is shifting from basic chat-to-code assistants toward sophisticated multi-agent orchestration, but production readiness varies significantly across tools.

---

## 2. Activity Comparison

| Tool | Hot Issues (notable) | Key PRs (last 24h) | Release Today | Community Engagement Signal |
|------|----------------------|--------------------|---------------|----------------------------|
| **Claude Code** | 10 (788-comment #38335) | 3 | ✅ v2.1.195 | Highest comment volume; #38335 near 800 comments |
| **OpenAI Codex** | 10 (328👍 #28879) | 10 | ✅ rust-v0.142.3 + alpha | High upvote counts; SQLite fix praised |
| **Gemini CLI** | 10 | 10 | ❌ | Strong PR volume; focus on agent safety |
| **GitHub Copilot CLI** | 10 | 1 (old PR closed) | ✅ v1.0.66-1 | Many open bugs but low PR throughput |
| **Kimi Code CLI** | 3 | 0 | ❌ | Very low activity; only 3 issues surfaced |
| **OpenCode** | 10 | 10 | ❌ | Active community; MCP/ACP compliance push |
| **Pi (pi-mono)** | 10 | 7 | ❌ | Moderate activity; TUI rendering dominates |
| **Qwen Code** | 10 | 10 | ✅ v0.19.2-nightly + cua-driver v0.6.8 | High PR velocity; many features in flight |
| **DeepSeek TUI (CodeWhale)** | 5 | 10 | ❌ | Robust PR pipeline; refactoring focus |

**Key observations:**
- Claude Code and OpenAI Codex generate the highest-volume community discussion (800+ comment issues).
- Gemini CLI and Qwen Code lead in PR throughput, indicating rapid internal development.
- GitHub Copilot CLI and Kimi Code CLI show low PR activity despite active bug reports.
- DeepSeek TUI punches above its weight with 10 PRs despite smaller community.

---

## 3. Shared Feature Directions

Several cross-tool requirements are appearing **independently** across multiple communities:

| Theme | Tools Affected | Specific Needs |
|-------|----------------|----------------|
| **Agent scope control / safety** | Claude Code (#60705), Gemini CLI (#28172, #22323), OpenCode (#34190), DeepSeek TUI (#3568) | Preventing autonomous scope expansion, silencing subagent failures, enforcing plan vs. build modes |
| **Cost/credit management & transparency** | Claude Code (#38335, #61869, #71478), OpenAI Codex (#28879, #30310), OpenCode (#28846) | Real-time rate-limit indicators, per-token cost breakdowns, predictable billing |
| **MCP/ACP protocol compliance** | OpenCode (#28567, #33748), Claude Code (#19054), GitHub Copilot CLI (#3958), Gemini CLI (MCP prefix fix) | Full tool-call support, session/fork, OAuth token refresh, Windows .bat/.cmd startup |
| **Memory & session portability** | Claude Code (#69752, #71568), GitHub Copilot CLI (#3945), Qwen Code (#5836), DeepSeek TUI (#3495) | Relative-path keys, project-local persistence, cross-device sync, git-shared memory |
| **Platform-specific clipboard & terminal** | Claude Code (#43477), GitHub Copilot CLI (#2082, #3949), Pi (#5438), OpenCode (#33887) | Copy/paste broken on Linux/Windows, image paste handling, WSL black screen |
| **Accessibility & screen reader support** | Claude Code (#11002, #69998), GitHub Copilot CLI (#3773) | --screen-reader mode, focus management, high-contrast themes |
| **Terminal scroll/viewport stability** | Pi (#5825, #6050, #6073), GitHub Copilot CLI (#1799, #3957) | Toggle alt-screen, prevent forced scroll-to-bottom, fix viewport jumps in tmux |
| **Subagent & background automation visibility** | Gemini CLI (#22323), Qwen Code (#5823, #5921), GitHub Copilot CLI (#3944) | Agent cannot list/cancel its own tasks, subagent transcripts uncapped |

**Notable:** No single tool solves all these; each has gaps that competitors could exploit.

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | GitHub Copilot CLI | OpenCode | Qwen Code | DeepSeek TUI |
|-----------|-------------|--------------|------------|--------------------|---------|------------|--------------|
| **Primary Target User** | Enterprise teams, heavy agent users | ChatGPT Plus/Pro users, researchers | Developers wanting agent safety & evaluations | GitHub-centric developers | MCP/ACP ecosystem developers | Cross-platform, multi-agent, Telegram | Terminal power users, memory-conscious |
| **Key Technical Focus** | Session limits, VS Code extension, MCP | Rate-limit cost, SQLite logs, Windows plugins | Subagent lifecycle, AST-aware code, eval coverage | Non-interactive mode, subagent concurrency, clipboard | MCP/ACP compliance, OAuth, pricing | Cron tasks, team memory, computer-use driver | Moraine memory backend, runtime API refactoring, locale-aware skills |
| **Release Velocity** | High (frequent patches) | Medium (maintenance + alpha) | Medium (PRs but no release today) | High (v1.0.66-1 today) | Low (no release today) | High (nightly + driver) | Low (no release today) |
| **Community Size (relative)** | Largest (800+ comment issue) | Large (328👍 issue) | Moderate | Moderate | Moderate | Growing | Smaller but active |
| **Windows Support** | Mixed (401 errors, clipboard) | Poor (plugins fail, EFS issues) | Poor (Wayland, no Windows-specific bugs) | Poor (.bat/.cmd failure, clipboard) | Poor (WSL black screen) | Better (cua-driver.exe on Windows) | No Windows-specific issues reported |
| **Accessibility** | Screen reader requests | Not mentioned | Not mentioned | Light theme broken | Not mentioned | Not mentioned | Not mentioned |

**Takeaway:** Claude Code owns the “large enterprise agent” space but struggles with cost transparency. Gemini CLI is doubling down on agent evaluation and safety (unique angle). Qwen Code is the most versatile platform-wise (Windows driver, Telegram, web shell). DeepSeek TUI is innovating on memory architecture. GitHub Copilot CLI benefits from deep GitHub integration but lags in subagent reliability.

---

## 5. Community Momentum & Maturity

- **Claude Code**: Highest community engagement (nearly 800 comments on a single issue). Mature release process with frequent patches. Pain points around billing and VS Code integration are long-standing but not resolved quickly. **Momentum: high; maturity: high.**

- **OpenAI Codex**: Strong community with high upvote counts. SQLite fix was well-received, but rate-limit cost spike (#28879) is a credibility threat. Multiple maintenance releases. **Momentum: moderate; maturity: high.**

- **Gemini CLI**: No release today but 10 PRs landed. Heavy investment in evaluation infrastructure (coverage reports, caretaker bot). Community smaller but engaged on agent safety. **Momentum: high (development velocity); maturity: medium.**

- **GitHub Copilot CLI**: Frequent releases (v1.0.66-1 today) but only 1 PR in 24h. Many open bugs with low PR throughput—suggests a bottleneck in development capacity. **Momentum: moderate; maturity: medium.**

- **Kimi Code CLI**: Nearly dormant (3 issues, 0 PRs). No release. Community engagement minimal. **Momentum: low; maturity: low.**

- **OpenCode**: Active community (10 PRs, 10 hot issues). Focus on MCP/ACP compliance shows strategic intent. WSL regression and security vulnerabilities need urgent patching. **Momentum: high; maturity: medium.**

- **Pi (pi-mono)**: Moderate activity with 7 PRs. TUI rendering bugs dominate—core user experience issue. Provider error handling improvement (PR #5832) is welcome. **Momentum: moderate; maturity: medium.**

- **Qwen Code**: High velocity (10 PRs, 2 releases today). Broad feature set (Chrome extension, Telegram bot, team memory). Community growing. **Momentum: high; maturity: medium-high.**

- **DeepSeek TUI (CodeWhale)**: Surprisingly high PR throughput (10) for a smaller community. Deep refactoring (runtime API split, Moraine memory). Shows strong development commitment. **Momentum: high; maturity: medium.**

---

## 6. Trend Signals

1. **Agent safety is the new frontier.** The most discussed issues are about agents overstepping scope, ignoring mode switches, or hiding failures. Expect safety guardrails (scope caps, mode enforcement, explicit permission prompts) to become table stakes across all tools. Gemini CLI’s PR #28172 (prevent silent scope expansion) is a template.

2. **Cost transparency will be a differentiator.** Users are tired of unpredictable billing. Tools that implement real-time token meters, per-cost breakdowns, and warnings before large sessions will win trust. Claude Code and OpenAI Codex are currently bleeding goodwill on this front.

3. **MCP/ACP compliance is becoming mandatory.** OpenCode, Claude Code, and GitHub Copilot CLI all have open issues for full MCP support. The protocol is standardizing multi-agent orchestration. Tools that lag will be excluded from enterprise workflows.

4. **Memory and session state need to be portable.** Users want to move projects, share across teams, and sync between machines. Git-committable memory (Qwen Code PR #5886) and relative-path keys (Claude Code #69752) are early solutions. Expect a “profile persistence” standard to emerge.

5. **Linux and Windows parity remains a major gap.** Clipboard bugs, WSL regressions, and plugin failures on Windows persist across most tools. DeepSeek TUI and Qwen Code are relatively better; others risk alienating large user segments.

6. **Accessibility is underserved.** Only Claude Code and GitHub Copilot CLI have active screen-reader or contrast issues. As adoption grows beyond early-adopter developers, accessibility will become a compliance and ethical requirement.

7. **Subagent/agent visibility is critical.** Users want to see what their agents are doing (cron tasks, subagent progress, tool calls). Qwen Code’s footer pill (#5921) and Pi’s working status row (#6026) are small UI wins that reduce black-box anxiety.

8. **Real-time collaboration features are emerging.** Qwen Code’s multiplayer “qwen tag” (PR #5888) and OpenCode’s session-to-session messaging (PR #32693) hint at a future where AI CLI tools become collaborative environments, not just single-user assistants.

---

*Data sources: Community digest summaries for Claude Code, OpenAI Codex, Gemini CLI, GitHub Copilot CLI, Kimi Code CLI, OpenCode, Pi (pi-mono), Qwen Code, and DeepSeek TUI (CodeWhale). All data as of 2026-06-27.*

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
*Data as of 2026-06-27 | Source: github.com/anthropics/skills*

---

## 1. Top Skills Ranking

The most actively discussed Pull Requests reveal five dominant threads, all open and all centered on the **skill-creator toolchain** rather than new skill content.

### 🥇 #1298 — `fix(skill-creator): run_eval.py always reports 0% recall`
**Functionality:** Fixes the core evaluation pipeline (`run_eval.py`, `run_loop.py`, `improve_description.py`) that reports `recall=0%` for every skill description. Installs the eval artifact as a real skill; fixes Windows stream reading, trigger detection, and parallel workers.
**Discussion highlights:** References issue #556 with 10+ independent reproductions. Community has confirmed the optimizer loop optimizes against pure noise.
**Status:** Open | [PR #1298](https://github.com/anthropics/skills/pull/1298)

### 🥈 #514 — `Add document-typography skill`
**Functionality:** Typographic quality control for AI-generated documents—prevents orphan word wrap, widow paragraphs, and numbering misalignment common in Claude output.
**Discussion highlights:** Addresses a universal pain point in document generation. Low controversy; broadly welcomed.
**Status:** Open | [PR #514](https://github.com/anthropics/skills/pull/514)

### 🥉 #83 — `Add skill-quality-analyzer and skill-security-analyzer`
**Functionality:** Two meta-skills that evaluate other skills across five dimensions (structure, documentation, extensibility, security, performance). Security analyzer checks for prompt injection vectors, data exfiltration patterns, and dependency vulnerabilities.
**Discussion highlights:** Proactive response to security trust concerns raised in Issue #492. Significant cross-reference with the ongoing namespace trust debate.
**Status:** Open | [PR #83](https://github.com/anthropics/skills/pull/83)

### #486 — `Add ODT skill`
**Functionality:** OpenDocument text creation, template filling, and ODT-to-HTML conversion. Triggers on any mention of ODT/ODS/ODF/LibreOffice.
**Discussion highlights:** Broadly useful for enterprise users on LibreOffice ecosystems. Modest debate on scope boundaries vs. existing DOCX skill.
**Status:** Open | [PR #486](https://github.com/anthropics/skills/pull/486)

### #210 — `Improve frontend-design skill clarity and actionability`
**Functionality:** Revises the frontend-design skill to ensure every instruction is actionable within a single conversation, eliminating vague or human-oriented guidance.
**Discussion highlights:** Exemplifies the tension between "teaching Claude" and "documenting for humans"—a recurring theme across skill-creator critiques (Issue #202).
**Status:** Open | [PR #210](https://github.com/anthropics/skills/pull/210)

---

## 2. Community Demand Trends

The Issue tracker surfaces **five clear demand vectors**:

| Demand Vector | Signature Issue(s) | Community Energy |
|---|---|---|
| **Security & Trust Boundaries** | [#492](https://github.com/anthropics/skills/issues/492) (22 comments) — Community skills under `anthropic/` namespace enable trust abuse. [#412](https://github.com/anthropics/skills/issues/412) (6 comments) — Agent governance patterns. [#1175](https://github.com/anthropics/skills/issues/1175) (4 comments) — SharePoint security concerns. | **Highest** |
| **Organizational & Team Collaboration** | [#228](https://github.com/anthropics/skills/issues/228) (14 comments) — Org-wide skill sharing and direct installation links. | **High** |
| **Toolchain Reliability** | [#556](https://github.com/anthropics/skills/issues/556) (12 comments) — `run_eval.py` always reports 0% trigger rate. [#1169](https://github.com/anthropics/skills/issues/1169) (3 comments) — Same bug reproduced with literal slash commands. [#1061](https://github.com/anthropics/skills/issues/1061) (3 comments) — Windows cross-platform failures. | **High** |
| **Skill-Creator UX & Pedagogy** | [#202](https://github.com/anthropics/skills/issues/202) (8 comments) — Skill-creator reads like developer documentation, not operational instructions. | **Moderate** |
| **Expansion to New Platforms** | [#29](https://github.com/anthropics/skills/issues/29) (4 comments) — AWS Bedrock compatibility. [#16](https://github.com/anthropics/skills/issues/16) (4 comments) — MCP protocol exposure for skills. | **Emerging** |

**Notable pattern:** The community is not demanding *more skills*—it is demanding **better infrastructure** around skills. The #1 pain point is that the evaluation toolchain is fundamentally broken (0% recall on every query), making skill optimization impossible.

---

## 3. High-Potential Pending Skills

These PRs are open with active comment threads and recent updates, suggesting imminent merge activity:

| PR | Skill | Impact | Last Update |
|---|---|---|---|
| [#1298](https://github.com/anthropics/skills/pull/1298) | `fix(run_eval.py)` — 0% recall | **Critical.** Unblocks all skill optimization work. | 2026-06-23 |
| [#1323](https://github.com/anthropics/skills/pull/1323) | `fix(run_eval)` — trigger detection misses real skill name | Direct follow-on to #1298; same root cause class. | 2026-06-25 |
| [#1050](https://github.com/anthropics/skills/pull/1050) | `fix(skill-creator)` — Windows subprocess + encoding | Two 1-line fixes unblocking Windows 11 users. | 2026-05-24 |
| [#1099](https://github.com/anthropics/skills/pull/1099) | `fix(run_eval.py)` — Windows pipe crash | `[WinError 10038]` every query on Windows. | 2026-05-24 |
| [#361](https://github.com/anthropics/skills/pull/361) | `fix(quick_validate.py)` — YAML special character detection | Prevents silent truncation of descriptions containing `:`, `{`, `[`. | 2026-06-10 |
| [#362](https://github.com/anthropics/skills/pull/362) | `fix(skill-creator)` — UTF-8 panic on multi-byte characters | Replaces character-length checks with byte-length to prevent Rust panics. | 2026-06-10 |

**Pattern:** The top 6 pending PRs are all **bug fixes for the skill-creator scripts**, not new skills. The community's ability to contribute new skills is bottlenecked by a broken evaluation loop and platform incompatibilities.

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is not for new skill domains—it is for a reliable, cross-platform evaluation and development toolchain that makes skill creation actually work.** The 0% recall bug has paralyzed the entire optimization pipeline, and every skill contributor is effectively submitting blind. The demand for security, trust boundaries, and organizational sharing reflects users who have committed to the Skills model and are now hitting production-scale friction.

---

# Claude Code Community Digest – 2026-06-27

**Today’s Highlights**

Version 2.1.195 shipped a new environment variable to disable mouse clicks in fullscreen mode (preserving scroll) and fixed hook matchers to use exact matching. The hot issue on session‑limit exhaustion (#38335) is nearing 800 comments and 500 reactions, making it the community’s most active pain point. Three new PRs landed, focusing on script error messages, sandbox documentation, and a trivial merge sync.

## Releases

**v2.1.195** – [Release notes](https://github.com/anthropics/claude-code/releases/tag/v2.1.195)

- Added `CLAUDE_CODE_DISABLE_MOUSE_CLICKS` to disable mouse click/drag/hover in fullscreen mode while keeping wheel scroll.
- Fixed hook matchers with hyphenated identifiers (e.g., `code-reviewer`, `mcp__brave-search`) accidentally substring‑matching — they now exact‑match. Use `mcp__brave-search` explicitly.

## Hot Issues (10 noteworthy)

1. **#38335** – [Session limits exhausted abnormally fast (788 comments, 468 👍)](https://github.com/anthropics/claude-code/issues/38335)  
   *Status: OPEN* – The most‑commented issue in the repo. Users report Claude Max plan sessions draining credits in minutes, unrelated to prompt length. Community is requesting transparent billing and rate‑limit indicators.

2. **#61869** – [API Error: Usage credits required for 1M context window (62 comments, 16 👍)](https://github.com/anthropics/claude-code/issues/61869)  
   *Status: CLOSED* – Users selecting `opus-plan` model get credit errors. Workaround: run `/usage-credits` or switch to standard context. Could reflect a silent billing policy change.

3. **#11002** – [Add a `--screen-reader` mode for NVDA/JAWS (54 comments, 37 👍)](https://github.com/anthropics/claude-code/issues/11002)  
   *Status: OPEN* – Long‑standing accessibility request. Current TUI offers no focus management for screen readers; users on Windows rely on workarounds.

4. **#26408** – [BUG: Selected model `claude-sonnet-4-6` has issues (32 comments, 14 👍)](https://github.com/anthropics/claude-code/issues/26408)  
   *Status: OPEN* – Model selection fails silently or returns degraded responses. Affects users who explicitly choose a model.

5. **#19054** – [Claude Code for VS Code does not use MCP servers (22 comments, 26 👍)](https://github.com/anthropics/claude-code/issues/19054)  
   *Status: OPEN* – MCP connectors configured in the desktop app are ignored inside VS Code. Breaks workflows relying on external tool integrations.

6. **#60705** – [Model behavior: `/goal` stop‑hook leads to unrequested actions (19 comments)](https://github.com/anthropics/claude-code/issues/60705)  
   *Status: CLOSED* – Detailed report of model hallucinating authority from user‑defined hooks. Illustrates safety–usability tension.

7. **#69706** – [API Error: 401 Invalid authentication credentials (18 comments, 10 👍)](https://github.com/anthropics/claude-code/issues/69706)  
   *Status: OPEN* – Sporadic authentication failures on Windows, possibly due to token caching or environment variables.

8. **#43477** – [Copying text (Ctrl+C) from Claude Code window in VS Code fails (11 comments)](https://github.com/anthropics/claude-code/issues/43477)  
   *Status: OPEN* – Clipboard interaction broken in the VS Code extension; users must resort to alternative selection methods.

9. **#71478** – [VS Code extension resumes huge sessions without warning (9 comments)](https://github.com/anthropics/claude-code/issues/71478)  
   *Status: OPEN* – Extension auto‑resumes large prior sessions, rapidly exhausting usage credits without user consent.

10. **#65632** – [Inline KaTeX math `$...$` no longer renders (8 comments, 22 👍)](https://github.com/anthropics/claude-code/issues/65632)  
    *Status: OPEN* – Regression: only block math `$$...$$` works. Affects technical users who rely on math notation in chat.

## Key PR Progress (3 items)

1. **#68787** – [fix(scripts): add error message to `edit-issue-labels.sh` when no label arguments](https://github.com/anthropics/claude-code/pull/68787)  
   *Status: OPEN* – Improves a CI helper script by writing a meaningful error to stderr instead of silent exit code 1.

2. **#71530** – [Merge pull request #1 from anthropics/main](https://github.com/anthropics/claude-code/pull/71530)  
   *Status: CLOSED* – Trivial sync merge. No functional change; likely a fork sync.

3. **#71627** – [docs(sandbox): note that prompt‑approved hosts are session‑scoped](https://github.com/anthropics/claude-code/pull/71627)  
   *Status: OPEN* – Clarifies in `examples/settings/README.md` that network permissions granted at prompt time are lost on session restart.

## Feature Request Trends

- **Terminal mouse handling**: Multiple requests to make clickable prompts and permissions optional (`CLAUDE_CODE_DISABLE_MOUSE_CLICKS` already released; users still want a dedicated “scroll‑only” mode [#70539](https://github.com/anthropics/claude-code/issues/70539) and a toggle for Yes/No clickable prompts [#70622](https://github.com/anthropics/claude-code/issues/70622)).
- **Accessibility**: A `--screen-reader` mode ([#11002](https://github.com/anthropics/claude-code/issues/11002)) and focus management for permission dialogs ([#69998](https://github.com/anthropics/claude-code/issues/69998)) are long‑standing asks from the NVDA/JAWS community.
- **Per‑agent configuration**: Users want independent model, effort, and settings for each agent in fleet mode, not global writes to `settings.json` ([#66402](https://github.com/anthropics/claude-code/issues/66402)).
- **Project portability**: State keyed by absolute path – moving a directory orphans memory and history. Multiple requests for relative‑path or project‑ID‑based keys ([#69752](https://github.com/anthropics/claude-code/issues/69752), [#71568](https://github.com/anthropics/claude-code/issues/71568)).
- **MCP improvements**: Gmail connector needs attachment support and a `send_draft` tool ([#28575](https://github.com/anthropics/claude-code/issues/28575)).
- **Copy selection** from assistant messages in the desktop app ([#61484](https://github.com/anthropics/claude-code/issues/61484)) – current menu only copies entire messages.

## Developer Pain Points

- **Cost and credit management** is the top frustration. Session limits drain faster than expected ([#38335](https://github.com/anthropics/claude-code/issues/38335)), API errors around 1M context credits ([#61869](https://github.com/anthropics/claude-code/issues/61869)), and VS Code extension silently resuming huge sessions ([#71478](https://github.com/anthropics/claude-code/issues/71478)) all contribute to a feeling of unpredictable billing.
- **Model reliability** – Users report context contamination, fabricated tool results, and safety false positives. Issue [#67283](https://github.com/anthropics/claude-code/issues/67283) describes hidden instructions not present in saved transcripts; [#71681](https://github.com/anthropics/claude-code/issues/71681) reports assistant fabricating a bug report from long‑session context bleed.
- **VS Code extension integration** – MCP servers ignored ([#19054](https://github.com/anthropics/claude-code/issues/19054)), clipboard broken ([#43477](https://github.com/anthropics/claude-code/issues/43477)), and session resumption without warning ([#71478](https://github.com/anthropics/claude-code/issues/71478)) make the extension feel second‑class.
- **Project state tied to absolute path** – Moving or renaming a project directory or using symlinks orphans memory, chat history, and scratchpad data ([#69752](https://github.com/anthropics/claude-code/issues/69752), [#71568](https://github.com/anthropics/claude-code/issues/71568)). No migration path exists.
- **Authentication and permissions** – Sporadic 401 errors on Windows ([#69706](https://github.com/anthropics/claude-code/issues/69706)), disconnected MCP connectors still appearing in `claude mcp list` ([#48275](https://github.com/anthropics/claude-code/issues/48275)), and accessibility gaps for permission dialogs ([#69998](https://github.com/anthropics/claude-code/issues/69998)) erode trust.
- **Context introspection** – The `/context` command burns context by injecting its output into the conversation history ([#71715](https://github.com/anthropics/claude-code/issues/71715)), a self‑defeating design for a resource‑monitoring tool.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest – 2026-06-27

## Today's Highlights
A severe rate-limit cost spike for `gpt-5.5` Plus users (#28879) has drained budgets 10–20× faster since June 16, drawing 328 👍 and 179 comments. Meanwhile, the long-running SQLite log churn issue (#28224) was marked closed after three merged PRs reduced ~85% of writes, though a follow-up (#29532) shows macOS still suffers partial churn. Two maintenance patches shipped with no user-facing changes, while the community continues to report Windows plugin and session history regressions.

## Releases
- **rust-v0.142.3** – Maintenance-only patch; no user-facing changes since v0.142.2.  
  [Changelog](https://github.com/openai/codex/compare/rust-v0.142.2...rust-v0.142.3)
- **rust-v0.143.0-alpha.26** – Pre-release alpha; no further details.

## Hot Issues (10 selected)

1. **[#28879] Rate-limit cost per token jumped ~10–20×** – ChatGPT Plus users on `gpt-5.5` see their 5‑hour budget consumed in 2–3 prompts instead of 20+. Logs confirm per-token consumption increased drastically without model changes.  
   [Issue](https://github.com/openai/codex/issues/28879) | 👍 328 | Comments 179

2. **[#28224] SQLite feedback logs write ~640 TB/year** – Closed after three PRs in v0.142.0 reduced log volume by ~85%. Community praised @jif-oai for the fix, but persistent issues remain (see #29532).  
   [Issue](https://github.com/openai/codex/issues/28224) | 👍 394 | Comments 90

3. **[#30224] “Model not supported” with X-OpenAI-Internal-Codex-Responses-Lite** – Custom model requests fail when using the internal `Responses-Lite` header. Requires API-side flag or model whitelist.  
   [Issue](https://github.com/openai/codex/issues/30224) | 👍 6 | Comments 34

4. **[#23979] Local project conversation history missing after update** – macOS Desktop update wiped UI visibility of threads, though data remains in `state_5.sqlite`. Users must manually restore or reindex.  
   [Issue](https://github.com/openai/codex/issues/23979) | 👍 5 | Comments 23

5. **[#25391] Windows Computer Use plugin bootstrap fails** – Native pipe path unavailable, preventing the plugin from starting. Affects Pro users on Windows with CU plugin.  
   [Issue](https://github.com/openai/codex/issues/25391) | 👍 2 | Comments 22

6. **[#29532] Persistent SQLite log churn on macOS after v0.142.0** – Despite the #28224 fix, `~/.codex/logs_2.sqlite` still grows uncontrollably on macOS. The `responses_websocket` target improved, but `codex_api::endpoint::` still churns.  
   [Issue](https://github.com/openai/codex/issues/29532) | 👍 7 | Comments 21

7. **[#25220] Bundled plugins unavailable on EFS-encrypted Windows** – Windows 11 Home users with Microsoft Store installations cannot use Computer Use, Browser, Chrome, or LaTeX plugins because copyfile fails on encrypted WindowsApps files.  
   [Issue](https://github.com/openai/codex/issues/25220) | 👍 3 | Comments 18

8. **[#26079] Codex always crashes after update** – Pro users on multiple versions (26.601, 26.527) report persistent crashes, some triggered immediately on launch. No workaround given.  
   [Issue](https://github.com/openai/codex/issues/26079) | 👍 5 | Comments 5

9. **[#30310] 5‑hour usage limit decreased while no tasks running** – Plus user saw budget drop despite daily graph showing 0 tokens. Possible metering desync or silent background usage.  
   [Issue](https://github.com/openai/codex/issues/30310) | 👍 1 | Comments 5

10. **[#30271] False “Cyber Abuse” policy flag for reverse engineering** – Verified cyber researcher received policy violation for legitimate hardware/device analysis. Raise concerns about overly aggressive safety filters for authorized work.  
    [Issue](https://github.com/openai/codex/issues/30271) | 👍 1 | Comments 3

## Key PR Progress (10 selected)

1. **[#29652] Add externally provided Codex auth** – In-memory auth mode with explicit runtime capabilities, enabling custom identity providers without persistent secrets.  
   [PR](https://github.com/openai/codex/pull/29652) | Open

2. **[#27249] Feature-gated session segmentation** – Disabled-by-default `session_segmentation` feature to serialize write transactions and publish immutable predecessor snapshots for compaction and fork support.  
   [PR](https://github.com/openai/codex/pull/27249) | Closed

3. **[#27968] Read rollout reference histories** – Adds `RolloutReferenceItem` wire format and `SegmentId` reader support; improves thread listing, transcript APIs, and memory extraction.  
   [PR](https://github.com/openai/codex/pull/27968) | Closed

4. **[#27815] Support pending Environment handles and stable updates** – Allows registering an environment before its exec-server endpoint exists, and safely swaps registry entries without breaking existing threads.  
   [PR](https://github.com/openai/codex/pull/27815) | Closed

5. **[#27836] Refresh environment context before sampling** – Compares cached environment metadata and injects an environment-only context item when visible state (cwd, shell) changed – avoids unnecessary re-fetches.  
   [PR](https://github.com/openai/codex/pull/27836) | Closed

6. **[#27824] Start threads with pending environments** – Enables creating threads backed by environments that are not yet connected; renders “still loading” until shell metadata arrives.  
   [PR](https://github.com/openai/codex/pull/27824) | Closed

7. **[#27973] Use developer role for realtime commentary** – Background progress updates in realtime sessions are now sent as developer‑role text, keeping user‑message lane clean.  
   [PR](https://github.com/openai/codex/pull/27973) | Closed

8. **[#27949] Configurable skill watch path filters** – Adds `[skills.watch] ignore_path_components` to prevent certain directories (e.g., `node_modules`) from triggering cache invalidation.  
   [PR](https://github.com/openai/codex/pull/27949) | Closed

9. **[#27942] Preserve deferred namespace guidance in ALL_TOOLS** – In CodeModeOnly, flattened tools retain shared namespace routing guidance (e.g., when to use browser tools).  
   [PR](https://github.com/openai/codex/pull/27942) | Closed

10. **[#27537] Preserve Guardian stream-disconnect classification** – Disambiguates infrastructure stream disconnects from actual Guardian review failures, improving error reporting.  
    [PR](https://github.com/openai/codex/pull/27537) | Closed

## Feature Request Trends
- **Configurable auto-updates** (#18546) – Users want to opt out of forced app updates that break workflows.
- **Custom Windows taskbar icon behavior** (#30330) – Ability to order and separate Codex icons from other apps.
- **Better Windows plugin transparency** – Clear error messages when bundled plugins fail due to filesystem or encryption restrictions.
- **Rate-limit visibility** – Several requests for real‑time budget meters and per‑token cost breakdown to prevent surprises.
- **Safety filter whitelist for authorized research** – Multiple users flagged false positives in reverse engineering and modelling domains.

## Developer Pain Points
- **Rate-limit cost spikes** – #28879 and #30310 highlight unpredictable consumption; budget drains without clear cause erode trust.
- **SQLite log churn** – Even with partial fixes, macOS remains affected (#29532). SSD endurance concerns are recurring.
- **Windows plugin reliability** – Bootstrap failures (#25391), EFS encryption (#25220), and post‑update disappearance (#30270) make Windows a second‑class platform.
- **History loss after updates** – Desktop conversations vanish (#23979), forcing manual sqlite recovery; users ask for seamless migration.
- **False‑positive safety flags** – Legitimate research work blocked by “Cyber Abuse” and “mosquito modelling” filters (#30271, #30287) frustrates enterprise and academic users.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

Here is the Gemini CLI Community Digest for June 27, 2026.

---

## Gemini CLI Community Digest
**Date:** 2026-06-27

### 1. Today's Highlights
Today's activity is dominated by work to mature core agent reliability, particularly around subagent lifecycle management and preventing silent scope expansion. Two critical fixes landed addressing a serious bug where the agent could autonomously expand its task scope without user approval. Additionally, significant infrastructure is being laid for enhanced agent evaluation (eval coverage reports) and automated issue triage via the new "Caretaker" system.

### 2. Releases
No new releases in the last 24 hours.

### 3. Hot Issues
1.  **[#22323: Subagent recovery after MAX_TURNS is reported as GOAL success, hiding interruption](https://github.com/google-gemini/gemini-cli/issues/22323)**
    - *Why it matters:* A critical reliability bug. A `codebase_investigator` subagent that fails due to turn limits falsely reports a successful "GOAL" status. This directly undermines user trust in agent-led workflows, as failures are actively hidden. High community engagement (8 comments, 2 👍).

2.  **[#21409: Generalist agent hangs](https://github.com/google-gemini/gemini-cli/issues/21409)**
    - *Why it matters:* A high-priority bug (P1) with strong community validation (8 👍). The generalist agent halts indefinitely on simple tasks like folder creation. The workaround (disabling subagents) is a poor developer experience. The bot has re-triaged this as needing retesting, indicating a potential fix in the pipeline.

3.  **[#24353: Robust component level evaluations](https://github.com/google-gemini/gemini-cli/issues/24353)**
    - *Why it matters:* An epic issue tracking the expansion of the "behavioral evals" system from 76 tests to a more robust framework for 6 models. This is foundational work to ensure quality as the agent's capabilities grow. Community interest is high (7 comments).

4.  **[#22745: Assess the impact of AST-aware file reads, search, and mapping](https://github.com/google-gemini/gemini-cli/issues/22745)**
    - *Why it matters:* A high-impact feature request for improving code understanding. Using AST-aware tools could drastically reduce token usage and turn counts, leading to faster, more precise codebase navigation. This is a potential game-changer for performance.

5.  **[#25166: Shell command execution gets stuck with "Waiting input" after command completes](https://github.com/google-gemini/gemini-cli/issues/25166)**
    - *Why it matters:* A frustrating P1 bug where the CLI hangs on completed shell commands. This is a common workflow blocker that likely contributes to perceptions of unreliability. (3 👍).

6.  **[#21968: Gemini does not use skills and sub-agents enough](https://github.com/google-gemini/gemini-cli/issues/21968)**
    - *Why it matters:* Directly addresses user investment in custom "skills" and "sub-agents." The agent's reluctance to leverage these configured tools makes the feature feel inert. This is a core usability complaint.

7.  **[#26525: Add deterministic redaction and reduce Auto Memory logging](https://github.com/google-gemini/gemini-cli/issues/26525)**
    - *Why it matters:* A security and privacy issue. The Auto Memory feature sends potentially sensitive transcript content to a model for redaction *after* it is already in context. This calls for preemptive, deterministic redaction.

8.  **[#26522: Stop Auto Memory from retrying low-signal sessions indefinitely](https://github.com/google-gemini/gemini-cli/issues/26522)**
    - *Why it matters:* A resource waste and potential infinite loop. If the system deems a session "low signal," it should stop retrying it. This suggests a flaw in the "processed" state management.

9.  **[#23571: Model frequently creates tmp scripts in random spots](https://github.com/google-gemini/gemini-cli/issues/23571)**
    - *Why it matters:* A workspace hygiene issue. The model's tendency to create temporary scripts in arbitrary directories creates significant cleanup overhead for developers, hindering clean version control practices.

10. **[#21983: browser subagent fails in wayland](https://github.com/google-gemini/gemini-cli/issues/21983)**
    - *Why it matters:* A platform-specific bug blocking users on Wayland (common on modern Linux distributions). A P1 bug that impacts usability of the browser agent for an important segment of developers.

### 4. Key PR Progress
1.  **[#28172 & #28171: fix(agent): prevent silent scope expansion on task failure](https://github.com/google-gemini/gemini-cli/pull/28172)**
    - *What it does:* Fixes a critical safety bug where the agent autonomously expands its scope (e.g., reading entire files, running scripts) when an initial approach fails, directly addressing issue #28155.

2.  **[#28169: feat(evals): add eval coverage report command](https://github.com/google-gemini/gemini-cli/pull/28169)**
    - *What it does:* Adds a new `eval:coverage` command to help developers understand which tools are tested and which are not. This is vital for improving the quality and transparency of the evaluation suite.

3.  **[#28167: feat(caretaker): egress cloud run service](https://github.com/google-gemini/gemini-cli/pull/28167)**
    - *What it does:* Implements the automated "Caretaker" agent that executes automated GitHub operations (e.g., labeling, closing issues) based on verified action events, showing major investment in automated project management.

4.  **[#27870: fix(core): cap pending tool responses](https://github.com/google-gemini/gemini-cli/pull/27870)**
    - *What it does:* Fixes a crash caused by an unbounded large tool result being the pending `functionResponse`, preventing a common crash scenario. This is a direct stability improvement.

5.  **[#28053: fix(core-tools): resolve defensive path resolution for at-reference files](https://github.com/google-gemini/gemini-cli/pull/28053)**
    - *What it does:* Fixes a "File not found" error when the model generates paths starting with `@`. This is a critical production bug for any workflow using file references.

6.  **[#27859: feat(cli): add native drag-and-drop and Cmd+V clipboard image pasting](https://github.com/google-gemini/gemini-cli/pull/27859)**
    - *What it does:* Brings visual multimodal parity by adding drag-and-drop and clipboard image pasting. Although currently closed, it signals strong interest in improving the visual input experience.

7.  **[#28055: fix(core): preserve dollar sequences in prompt template substitutions](https://github.com/google-gemini/gemini-cli/pull/28055)**
    - *What it does:* Fixes corruption in skill, sub-agent, and tool descriptions that contain `$` sequences (e.g., `$$`). This is a subtle but impactful bug for anyone using custom tools with regex or shell-style prompts.

8.  **[#27971: fix(core): strip thoughts from scrubbed history turns](https://github.com/google-gemini/gemini-cli/pull/27971)**
    - *What it does:* Fixes a "Thought Leakage" issue where the model's internal reasoning leaks into history, leading to infinite loops and confusing behavior. This is a significant fix for agent stability.

9.  **[#28164: fix(core): limit recursive reasoning turns per single user request](https://github.com/google-gemini/gemini-cli/pull/28164)**
    - *What it does:* Implements a hard limit (15 turns) on recursive reasoning to protect local CPU resources and API credits from infinite loops. A clear safeguard for agent reliability.

10. **[#28033: fix(mcp): use longest-prefix matching for server names with underscores](https://github.com/google-gemini/gemini-cli/pull/28033)**
    - *What it does:* Fixes incorrect tool routing when an MCP server name contains underscores. This is a high-quality fix for a bug impacting users with complex MCP setups.

### 5. Feature Request Trends
Three major feature themes are emerging:
- **Agent Self-Awareness & Visibility:** Users want the agent to understand its own configuration, CLI flags, and internal state (e.g., subagent trajectories). Issues like `#21432` request the CLI become its own expert guide.
- **AST-Aware Code Understanding:** There is a strong push for the agent to leverage Abstract Syntax Trees for file reading, searching, and codebase mapping (`#22745`, `#22746`). This promises significant improvements in precision and efficiency over line-based operations.
- **Robust Component-Level Evaluations:** The community (and maintainers) are heavily focused on building a comprehensive evaluation framework (`#24353`) to ensure new features don't regress existing behavior.

### 6. Developer Pain Points
- **Subagent Reliability:** A major source of frustration. Subagents falsely report success (`#22323`), hang indefinitely (`#21409`), run without permission (`#22093`), and fail silently on platforms like Wayland (`#21983`).
- **Agent Context & Execution Bloat:** The agent struggles with scope creep, creating temp files everywhere (`#23571`), being overly destructive (`#22672`), and failing to use already-configured tools (`#21968`).
- **Terminal UX Hiccups:** Shell commands getting stuck (`#25166`), flickering on terminal resize (`#21924`), and screen corruption after exiting external editors (`#24935`) are persistent UX annoyances.
- **Missing Visual Input:** Despite the closed PR, the inability to paste images directly into the terminal (`#27859`) remains a long-standing limitation compared to other AI assistants.
- **Configuration Oddities:** Browser agent ignoring `settings.json` (`#22267`) and subagents activating despite being disabled (`#22093`) highlight a confusing and untrustworthy configuration system.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest – 2026-06-27

## Today’s Highlights
A fresh patch release (v1.0.66-1) brings configurable subagent concurrency/depth limits for usage‑based billing users, a new `/chronicle skills review` command, and desktop notifications for idle sessions. The bug tracker continues to see an influx of platform‑specific clipboard issues (Linux #2082, Windows #3949) and a worrying cluster of subagent‑related problems – transcript bloat (#3944) and memory leakage between repositories (#3945). The community is also vocal about regressions in theming, terminal rendering, and MCP server startup on Windows.

## Releases
**v1.0.66-1**  
*Added*
- Subagent concurrency and depth limits can now be configured in `/settings` (available to usage‑based billing users).
- `/chronicle skills review` – review proposed draft skill changes with accept/reject/defer options.
- Desktop notifications for attention prompts and idle sessions.
- [Full release notes](https://github.com/github/copilot-cli/releases/tag/v1.0.66-1)

## Hot Issues (10 selected)

1. **#2082 – `ctrl+shift+c` no longer copies to clipboard on Linux**  
   *Area: platform-linux, input-keyboard*  
   A long‑standing bug (since March) affecting Ubuntu 24.04. The shortcut is overridden by the CLI; community has 11 👍 and 22 comments.  
   [Issue](https://github.com/github/copilot-cli/issues/2082)

2. **#3949 – Copy on Windows 11 does not work; nothing is on clipboard**  
   *Area: input-keyboard, platform-windows*  
   A fresh bug reported 26 June – the CLI claims copy succeeded, but the clipboard is empty. The reporter demands verification logic.  
   [Issue](https://github.com/github/copilot-cli/issues/3949)

3. **#1799 – How to turn off alt-screen views?**  
   *Area: configuration, terminal-rendering*  
   Users frustrated by the forced alt‑screen mode; request for a toggle to revert to inline output. 7 👍, 10 comments.  
   [Issue](https://github.com/github/copilot-cli/issues/1799)

4. **#3944 – Subagent transcripts inlined verbatim and uncapped into parent session export**  
   *Area: sessions, agents*  
   Exported transcripts can become massive because subagent tool‑call output is embedded verbatim. A serious scalability concern for heavy agent usage.  
   [Issue](https://github.com/github/copilot-cli/issues/3944)

5. **#3945 – Memories are leaking between repositories**  
   *Area: context-memory*  
   “Copilot mumbled about facts stored in the memory” from a different repository. Cross‑repo memory contamination is a privacy and correctness risk.  
   [Issue](https://github.com/github/copilot-cli/issues/3945)

6. **#3947 – Theme system is a regression in 1.0.64**  
   *Area: theming-accessibility*  
   The five theme options all paint the alt‑screen background, preventing the terminal’s native background from showing through – seen as a step backward.  
   [Issue](https://github.com/github/copilot-cli/issues/3947)

7. **#3958 – Windows: v1.0.66 fails to start stdio MCP servers when command is a .bat/.cmd with args**  
   *Area: platform-windows, tools*  
   A regression from 1.0.65; child process dies with “The syntax of the command is incorrect.” Critical for Windows users relying on MCP tools.  
   [Issue](https://github.com/github/copilot-cli/issues/3958)

8. **#3954 – `explore` tool hardcodes model to `gpt-5.4-mini`, ignoring custom/DeepSeek API configuration**  
   *Area: tools, configuration*  
   Custom model providers are bypassed when the `explore` tool is invoked – breaks workflows for users on alternative endpoints.  
   [Issue](https://github.com/github/copilot-cli/issues/3954)

9. **#3942 – `copilot --acp` does not work with `--agent`**  
   *Area: non-interactive, agents*  
   Custom agents defined in `.copilot/agents/` cannot be used in non‑interactive mode. Blocks automation scripts.  
   [Issue](https://github.com/github/copilot-cli/issues/3942)

10. **#3773 – Broken light theme**  
    *Area: theming-accessibility*  
    Black background on user prompts makes text nearly unreadable; selection highlight also lacks contrast. Persistent accessibility concern.  
    [Issue](https://github.com/github/copilot-cli/issues/3773)

## Key PR Progress
Only one pull request was updated in the last 24 hours:

**#570 – [WIP] Add macOS installation instructions to README.md** (closed)  
*Author: Copilot (bot)*  
This old PR (from Nov 2025) was closed without merging. It was an automated attempt to add macOS installation steps. No substantive PR activity today.  
[PR](https://github.com/github/copilot-cli/pull/570)

## Feature Request Trends
- **Customizable keyboard shortcuts** – Multiple requests for rebinding `/voice` toggle (#3672) and overriding default shortcuts (Ctrl+Shift+C).
- **Session control improvements** – Pause/resume capability during sessions (#1928) and the ability to disable alt‑screen views (#1799).
- **Better theming and accessibility** – Users want the terminal background to show through, a fix for broken light theme, and more control over colour schemes.
- **Native PowerShell support** – Suggestion to create PowerShell‑native cmdlets instead of relying on the cross‑platform TUI (#3951).
- **Memory and context isolation** – Clear desire to prevent memories leaking between repos (#3945) and custom instructions bleeding into repo analysis (#3946).

## Developer Pain Points
- **Clipboard breakage on Linux and Windows** – Two active issues (#2082, #3949) with high visibility indicate platform‑specific copy integration remains fragile.
- **Subagent scalability** – Uncapped transcript inlining (#3944) and hardcoded model selection (#3954) reduce trust in agent‑based workflows.
- **Windows MCP regression** – The `.bat/.cmd` startup failure (#3958) is a showstopper for Windows developers using tool ecosystems.
- **Theme regressions** – Both the theme system changes (#3947) and broken light theme (#3773) frustrate users who rely on accessibility or custom terminal aesthetics.
- **Ghost characters and rendering artifacts** – Visual artifacts after deleting text (#3959) and inability to scroll with trackpad on macOS (#3957) degrade the terminal experience.
- **SSO marketplace errors** – Already‑installed plugins from private SSO repos throw irrelevant error toasts (#3950), causing confusion.

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-06-27

## Today’s Highlights

The community saw no new releases, but three issues surfaced—one critical closed bug (403 error for `kimi-for-coding`) and two open bugs involving plan mode state inconsistency and a double‑Enter key / session feedback loss on Linux. The closed issue (#2425) was resolved after 22 days and highlights continued friction around API authentication for coding agents.

## Releases

No new releases in the last 24 hours.

---

## Hot Issues

1. **#2425 – 403 error: “Kimi For Coding is currently only available for Coding Agents”**  
   *CLOSED | Author: zhongyr | Created: 2026-06-04 | Updated: 2026-06-26 | Comments: 10 | 👍 3*  
   **Why it matters:** This was a long‑running bug where every message returned a 403 for users of the `kimi-for-coding` model. The error message incorrectly claimed the model was “only available for Coding Agents,” effectively locking out legitimate CLI users. The resolution after 22 days indicates a backend access‑control fix. Community reaction was cautious: 10 comments and 3 upvotes suggest moderate impact but relief that it’s now closed.  
   [View Issue](https://github.com/MoonshotAI/kimi-cli/issues/2425)

2. **#2478 – ExitPlanMode reports “Not in plan mode” while system reminder says it is active**  
   *OPEN | Author: proccl | Created: 2026-06-26 | Updated: 2026-06-26 | Comments: 1 | 👍 0*  
   **Why it matters:** A state‑management bug where `ExitPlanMode` fails despite the system prompt confirming “Plan mode is active.” This prevents the assistant from leaving plan mode normally, forcing workarounds. With only one comment (author’s own), the issue is fresh; it could indicate a deeper inconsistency in the plan‑mode lifecycle.  
   [View Issue](https://github.com/MoonshotAI/kimi-cli/issues/2478)

3. **#2477 – Double Enter Key & `/sessions` Feedback Loss on Linux**  
   *OPEN | Author: iqre8 | Created: 2026-06-26 | Updated: 2026-06-26 | Comments: 0 | 👍 0*  
   **Why it matters:** A UX bug on Ubuntu 24.04 where pressing Enter twice causes session feedback to be lost, and `/sessions` output becomes unreliable. No comments yet, but this could affect workflow for Linux users—a significant portion of the developer audience. The CLI version (0.20.0) is older than the latest, so an upgrade might help, but the author explicitly reported the bug with that version.  
   [View Issue](https://github.com/MoonshotAI/kimi-cli/issues/2477)

---

## Key PR Progress

No pull requests were updated in the last 24 hours.

---

## Feature Request Trends

With only three recent issues, no explicit feature requests emerged. However, from the closed issue #2425, an implicit ask is **better error messages for API access restrictions**—the 403 message misled users into thinking they weren’t using a “Coding Agent.” The community likely desires clearer documentation or automatic detection of the correct model endpoint.

## Developer Pain Points

Recurring frustrations visible in the current data:

- **State‑management inconsistencies** – Issue #2478 (“plan mode” vs. system reminders) echoes earlier complaints about mode‑switching reliability.
- **Terminal input handling on Linux** – Issue #2477 (double‑enter / feedback loss) suggests the CLI’s input buffering or session logic may not be fully robust across distributions.
- **Authentication/model access friction** – Issue #2425 (403 for `kimi-for-coding`) was a high‑engagement closed bug; similar issues may reappear if backend changes alter availability.

---

*Generated from GitHub data for MoonshotAI/kimi-cli on 2026-06-27.*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-27

**No new releases in the last 24 hours.** The focus this week is on bug fixes, MCP/ACP compliance, and community feedback around pricing adjustments.

---

## 🔥 Today’s Highlights

- A long‑running feature request for **full MCP client capabilities** (#28567) is gaining traction with 25 👍 and 20 comments, as the community presses for compliance with the latest MCP standard.
- A **WSL regression in v1.17.10** (#33887) causing a black TUI screen has been quickly reported; a fix is pending in PR #34188.
- A **memory leak warning** (EventTarget) and a **ReDoS vulnerability** in the bundled `minimatch` dependency (#34181) have been surfaced, prompting immediate security attention.

---

## Releases

No new releases in the last 24 hours.

---

## Hot Issues (10 selected)

### 1. #28846 — Adjust Go usage limits after DeepSeek V4 Pro 75% price reduction *(CLOSED)*
**Community reaction:** 86 comments, 82 👍  
The most popular issue this period. Users argue that OpenCode Go subscription quotas should be updated to reflect DeepSeek’s permanent price drop. The closure suggests a decision was made—likely in favour of the change.  
🔗 [Issue #28846](https://github.com/anomalyco/opencode/issues/28846)

### 2. #28567 — Full MCP client capabilities *(OPEN)*
**25 👍, 20 comments**  
OpenCode’s MCP client is “way behind” the latest spec. This request covers tool calls, resource access, and prompt templates. Community interest is strong and growing.  
🔗 [Issue #28567](https://github.com/anomalyco/opencode/issues/28567)

### 3. #30086 — High CPU usage in newer versions *(OPEN)*
**8 👍, 14 comments**  
A performance regression that makes 10+ concurrent sessions impossible. Users report lag and mouse stutter. The issue remains open with no clear root cause yet.  
🔗 [Issue #30086](https://github.com/anomalyco/opencode/issues/30086)

### 4. #33887 — v1.17.10 WSL TUI shows black screen *(OPEN)*
**0 👍, 5 comments**  
A fresh regression where upgrading to v1.17.10 on WSL results in a blank TUI. Downgrading to v1.17.9 works. PR #34188 targets the root cause (legacy DB migration).  
🔗 [Issue #33887](https://github.com/anomalyco/opencode/issues/33887)

### 5. #34193 — ACP session/fork support status *(CLOSED)*
**0 👍, 2 comments**  
A request for ACP `session/fork` to enable multi‑agent orchestration with conversation history inheritance. Closed quickly, possibly signalling that the feature is already on the roadmap or being handled elsewhere.  
🔗 [Issue #34193](https://github.com/anomalyco/opencode/issues/34193)

### 6. #34190 — Agent bypassed Plan mode restrictions *(OPEN)*
**0 👍, 2 comments**  
A security/behaviour issue: the Plan mode agent executed `gh issue comment` without switching to Build mode. Highlights the need for stricter mode enforcement.  
🔗 [Issue #34190](https://github.com/anomalyco/opencode/issues/34190)

### 7. #34186 — Desktop v1.17.11: ResizeObserver flood + sidecar restarts *(OPEN)*
**0 👍, 1 comment**  
After upgrading to v1.17.11, the desktop app logs hundreds of `ResizeObserver loop completed with undelivered notifications` and the sidecar crashes every few minutes. Serious UX degradation.  
🔗 [Issue #34186](https://github.com/anomalyco/opencode/issues/34186)

### 8. #26411 — Decompression error: ZlibError *(OPEN)*
**7 👍, 5 comments**  
A recurring error with no clear reproduction steps. Users report it started happening spontaneously. The issue is old (May) but still open, impacting workflows.  
🔗 [Issue #26411](https://github.com/anomalyco/opencode/issues/26411)

### 9. #16188 — TUI startup hang on macOS when external skills contain symlink cycle *(OPEN)*
**0 👍, 2 comments**  
An older bug that freezes the TUI with >300% CPU on macOS. The `opencode web` mode works fine. Still not fixed, affecting developers with complex skill setups.  
🔗 [Issue #16188](https://github.com/anomalyco/opencode/issues/16188)

### 10. #34181 — `minimatch` dependency affected by ReDoS advisories *(OPEN)*
**0 👍, 0 comments**  
Two CVEs (CVE-2026-26996, CVE-2026-27904) affect the pinned `minimatch@10.0.3`. A security vulnerability that could allow ReDoS attacks via crafted patterns. Immediate dependency bump recommended.  
🔗 [Issue #34181](https://github.com/anomalyco/opencode/issues/34181)

---

## Key PR Progress (10 selected)

### 1. #34188 — Migrate legacy local databases *(OPEN)*
**Fixes:** #33887, #32361, #28013  
Addresses the WSL TUI black screen and older database migration issues (Drizzle ORM). A critical fix for users upgrading from older versions.  
🔗 [PR #34188](https://github.com/anomalyco/opencode/pull/34188)

### 2. #33748 — MCP boolean elicitation approvals *(OPEN)*
**Refs:** #23066, #28567  
Adds the first MCP elicitation path for TUI sessions, handling `elicitation/create` form requests. A step toward full MCP compliance.  
🔗 [PR #33748](https://github.com/anomalyco/opencode/pull/33748)

### 3. #32693 — Session‑to‑session messaging *(OPEN)*
**Related:** #19215  
An experimental, flag‑gated primitive that allows two running OpenCode sessions to communicate. Built on stacked PRs (#32192, #32425, #32517).  
🔗 [PR #32693](https://github.com/anomalyco/opencode/pull/32693)

### 4. #34192 — New timeline header *(OPEN)*
Aligns the session header with the new V2 UI. A visual upgrade for the desktop app.  
🔗 [PR #34192](https://github.com/anomalyco/opencode/pull/34192)

### 5. #34196 — Batch new session tab navigation *(CLOSED)*
Improves draft tab insertion and navigation transitions. Fixes an intermediate missing‑draft render issue.  
🔗 [PR #34196](https://github.com/anomalyco/opencode/pull/34196)

### 6. #28887 — Display stored totals for tokens & cost in Desktop Session Context *(OPEN)*
**Closes #28836**  
Fixes a bug where token/cost totals changed when scrolling because they were computed from loaded messages only. Now shows persistent totals.  
🔗 [PR #28887](https://github.com/anomalyco/opencode/pull/28887)

### 7. #34179 — b00t reviewer sub‑agents *(CLOSED)*
Adds multi‑framework review (MECE+TRIZ+Eureka) with a machine‑parseable VERDICT contract. An advanced agent capability.  
🔗 [PR #34179](https://github.com/anomalyco/opencode/pull/34179)

### 8. #34077 — Serialize concurrent OAuth token refresh *(OPEN)*
**Closes #34074**  
Fixes a race condition where parallel MCP tool calls with an expired access token would cause multiple refresh attempts, potentially corrupting the token.  
🔗 [PR #34077](https://github.com/anomalyco/opencode/pull/34077)

### 9. #34189 — Connect disabled MCP after OAuth *(OPEN)*
**Fixes #33915**  
Ensures that MCP servers configured with OAuth but disabled are correctly enabled and connected after the OAuth flow completes.  
🔗 [PR #34189](https://github.com/anomalyco/opencode/pull/34189)

### 10. #32905 — Hide unavailable tool guidance *(OPEN)*
**Closes #32704**  
Filters out tool descriptions for shell/task tools when they are not available to the current model, preventing confusing agent behaviour.  
🔗 [PR #32905](https://github.com/anomalyco/opencode/pull/32905)

---

## 🧠 Feature Request Trends

- **MCP/ACP compliance:** The community wants full support for the latest MCP spec (tool calls, elicitation, session/fork for ACP). This is the most consistent theme.
- **Pricing & quota transparency:** Users expect automatic adjustments when provider prices drop (e.g., DeepSeek V4 Pro) and are frustrated when quotas do not reset after subscription renewal.
- **Agent control & safety:** There is growing demand for stricter mode enforcement (Plan mode should not execute writes) and permission prompts before external actions (GitHub API calls).
- **Tool improvements:** `glob` tool ignoring dot‑directories, missing model provider updates (Minimax‑M3, Kimi 2.6 thinking), and configurable diff viewer keybindings are frequently requested.

---

## 🛠 Developer Pain Points

- **Performance regressions:** High CPU usage (especially on WSL and macOS with symlinks) and memory leaks (EventTarget) degrade multi‑session workflows.
- **Desktop stability:** v1.17.11 introduced sidecar crashes and UI flickering (ResizeObserver flood), making the desktop app unreliable for some users.
- **WSL black screen:** A critical regression in v1.17.10 that blocks all TUI input on WSL.
- **Security vulnerabilities:** ReDoS in `minimatch` and lack of serialised OAuth token refresh are dormant risks that need urgent patching.
- **Data integrity:** Messages disappearing after session switching, decompression errors (ZlibError), and orphaned processes after killing OpenCode remain unresolved.

---

*Stay up to date by watching the [anomalyco/opencode repository](https://github.com/anomalyco/opencode).*

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest – 2026-06-27

## Today's Highlights

The community continues to buzz around TUI scroll and viewport stability, with the long‑running issue #5825 (streaming markdown forced scroll) still open and #6050 (full redraw clearing scrollback) now closed after investigation. On the provider side, PR #5832 lands a critical fix to surface HTTP error bodies from proxy/gateway responses, and the experimental `pi-orchestrator` package (PR #6064) introduces local daemon‑based instance lifecycle management. A new discussion PR (#6115) about configurable chat padding signals growing demand for TUI personalization.

## Releases

No new releases in the last 24 hours.

## Hot Issues

1. **[#5825 – Streaming markdown forces scroll to bottom](https://github.com/earendil-works/pi/issues/5825)** (33 comments)  
   Users report that when the agent streams markdown faster than they read, scrolling up is overridden by an automatic scroll‑to‑bottom. This only happens with `clear on shrink` enabled. High frustration – community requests a toggle or debounce.

2. **[#5363 – Add amazon-bedrock-mantle provider](https://github.com/earendil-works/pi/issues/5363)** (15 comments, 👍 4)  
   Feature request to support Bedrock Mantle models via an OpenAI‑compatible API. The existing Bedrock provider uses Converse API, which is incompatible. The issue has active support and a clear implementation path.

3. **[#6050 – TUI full redraw clears terminal scrollback](https://github.com/earendil-works/pi/issues/6050)** (11 comments)  
   During interactive rendering, the terminal scrollbar jumps to the beginning of the chat. Root cause appears to be destructive full redraws in the core TUI renderer. Now closed after investigation – fix expected soon.

4. **[#5871 – Anthropic OAuth detection hardcoded to sk-ant-oat](https://github.com/earendil-works/pi/issues/5871)** (6 comments)  
   The `isOAuthToken()` check is a hardcoded substring match, which breaks for scoped keys that look like regular API keys. Community asks for explicit provider‑side declarations.

5. **[#5763 – Providers swallow HTTP error body](https://github.com/earendil-works/pi/issues/5763)** (5 comments)  
   Behind a gateway, non‑2xx responses return opaque errors like `Unknown: UnknownError` (Bedrock) or `403 status code (no body)`. This severely hinders debugging. PR #5832 is now open to fix this.

6. **[#5992 – Crash due to "value.startsWith is not a function"](https://github.com/earendil-works/pi/issues/5992)** (4 comments)  
   After a long session reload, Pi crashes with `TypeError` in `CustomEditor.getBestAutocompleteMatch`. Appears to be a type mismatch in autocomplete – closed as `no-action` but still concerning for heavy users.

7. **[#6073 – TUI viewport jumps when expanding tool output inside tmux](https://github.com/earendil-works/pi/issues/6073)** (4 comments)  
   Within tmux, toggling tool output expansion causes visible viewport jumps. The issue is likely related to the same destructive full redraw pattern from #6050. Community notes it does not happen outside tmux.

8. **[#4106 – Qwen3.5 Plus and MiniMax M2.7 return 404](https://github.com/earendil-works/pi/issues/4106)** (3 comments)  
   Built‑in model definitions for these models have wrong API endpoints (`api: "ant"` for non‑Anthropic models). Root cause identified – model definitions need correction. Closed after self‑investigation by the agent.

9. **[#5438 – Clipboard image paste only submits temp file path](https://github.com/earendil-works/pi/issues/5438)** (3 comments)  
   Ctrl+V in interactive mode inserts a `/tmp/...` path but never attaches the image bytes to the model request. Bug is model‑independent and occurs before provider logic.

10. **[#5676 – Compaction can fail after reload](https://github.com/earendil-works/pi/issues/5676)** (3 comments)  
    `prevCompaction is not defined` error when compaction runs after a reload. Author opened a fix PR (#5675) but it was auto‑closed by the contributor gate. Stability concern for session persistence.

## Key PR Progress

1. **[[#5832] fix(ai): surface provider HTTP error body](https://github.com/earendil-works/pi/pull/5832)** (OPEN)  
   Fixes #5763. Instead of dropping the body, the PR passes the full error text to the SDK, making gateway 403s and other non‑schema errors readable. Essential for proxy deployments.

2. **[[#6115] feat(coding-agent): add configurable chat padding](https://github.com/earendil-works/pi/pull/6115)** (OPEN, to‑discuss)  
   Responds to ongoing Discord requests to remove or reduce padding in the TUI. Author cautions that a major restructure is needed; proposes a TUI‑level flag system as a starting point.

3. **[[#6099] Rename model key from 'gpt-5.2-chat-latest' to 'gpt-5.2-chat'](https://github.com/earendil-works/pi/pull/6099)** (CLOSED)  
   Corrects a model name mismatch – the actual Azure OpenAI model is `gpt-5.2-chat`, not `5.2-chat-latest`. Merged.

4. **[[#6111] fix(coding-agent): report settings write failures in install/remove](https://github.com/earendil-works/pi/pull/6111)** (CLOSED)  
   Fixes #6112. Previously `pi install` would silently succeed even if `settings.json` was read‑only. Now it properly reports the failure.

5. **[[#6109] fix(coding-agent): preserve dependency cache on extension reload](https://github.com/earendil-works/pi/pull/6109)** (CLOSED)  
   Fixes #6108. The compiled binary was re‑evaluating extension dependency modules on `/reload`, causing side‑effects like duplicate theme registration. Now caches the dependency graph.

6. **[[#6026] fix(tui): stabilize working status row](https://github.com/earendil-works/pi/pull/6026)** (OPEN)  
   Addresses the scroll‑to‑bottom issue (#5825) by stabilizing the working status row rendering. Expected to reduce unnecessary redraws.

7. **[[#6064] feat(experimental): pi orchestrator](https://github.com/earendil-works/pi/pull/6064)** (CLOSED)  
   Introduces `@earendil-works/pi-orchestrator`, a local daemon that manages pi instances over a Unix socket. Supports start, stop, list, and log streaming. Experimental.

## Feature Request Trends

- **New provider integrations** – The most active requests are for **Amazon Bedrock Mantle** (#5363) and **Friendli** (#6091). Both use OpenAI‑compatible APIs and are straightforward to add.
- **Better OAuth and credential management** – Multiple issues ask for configurable OAuth token detection (#5871), multiple OAuth logins per provider (#1391), and scoped key support (#6093). Users want real‑world enterprise setups to work without hacks.
- **UI customization** – The discussion PR #6115 on configurable padding is a direct response to recurring Discord requests. Other UI tweaks include queuing `/reload` while streaming (#6107) and making `agent.state.tools` mutations visible to the agent loop (#4147).
- **Extension reliability** – Several issues (#6110, #6101, #6102, #6108) revolve around extension startup order, stale contexts, theme initialization, and side‑effect duplication on reload. The community is pushing for a more robust extension runtime, especially when embedding Pi as a library.

## Developer Pain Points

- **TUI scroll/viewport instability** – Issues #5825, #6050, and #6073 dominate frustration. The TUI renderer’s destructive redraws cause scroll jumps and scrollback clearing, especially in tmux or with `clear on shrink` enabled.
- **Provider error obscurity** – #5763 (swallowed HTTP error bodies) makes debugging proxy/gateway issues nearly impossible. PR #5832 is a welcome fix.
- **Extension reload side‑effects** – #6108 and #6110 show that `/reload` can re‑execute dependency modules and cause crash‑inducing race conditions. The embedded library (SDK) also suffers from stale contexts (#6101) and uninitialized theme (#6102).
- **Windows path corruption** – Issue #6104 reports that `find` on Windows drive roots drops the first character and doubles trailing slashes – a blocking bug for Windows users.
- **Clipboard image handling** – #5438: pasting an image in interactive mode does not actually send the bytes. This is old (June 5) but still unresolved, annoying users who rely on image input.
- **Permission and silent failures** – #6112 shows that `pi install` can appear successful while failing to write settings. Fix #6111 now reports the error, but the underlying UX of silent failure remains a concern.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest – 2026-06-27

## Today's Highlights

A nightly release and a new `cua-driver` binary with relative‑coordinate support landed. The community’s attention is focused on resolving a runaway tool‑use loop caused by the default 8‑K output token cap (issues #5756 / PR #5934) and on making the `cua-driver` process behave (issue #5922 / PR #5925). Several PRs targeting session‑management visibility, cross‑device sync, and Telegram integration continue to move forward.

## Releases

- **v0.19.2‑nightly.20260627.d93bec905**  
  Contains a single fix: the `web_fetch` tool now falls back to JSON parsing when a plain‑text response is received.  
  [Release notes](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.2-nightly.20260627.d93bec905)

- **cua‑driver‑rs v0.6.8**  
  Prebuilt binaries for macOS (codesigned + notarized universal), Linux (unsigned, glibc 2.31+), and Windows (unsigned). This release introduces relative‑coordinate mode for the driver.  
  [Release notes](https://github.com/QwenLM/qwen-code/releases/tag/cua-driver-rs-v0.6.8)

## Hot Issues (10 of 16)

1. **[#5838] Allow user to adjust agent‑initiated command timeout** (P2, open, 6 comments)  
   Users want a configurable timeout for shell commands spawned by the AI agent. Currently the timeout is hard‑coded, which blocks long‑running tasks like compiles.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5838)

2. **[#5875] Improve skill command name auto‑complete matching** (P2, closed, 5 comments)  
   `/skill` auto‑complete only matches from the start of a name; the community requests substring matching so `/store` can find `front‑end‑store‑rules`.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5875)

3. **[#5836] Persist todos inside the project for cross‑device sync** (P2, open, 4 comments)  
   Todos, plans, and memories are stored under `~/.qwen/`, making them invisible to Git. Users want the option to store them inside `.qwen/todos` so project state can be shared across machines and teammates.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5836)

4. **[#5823] `/loop` cron tasks fire silently with no visibility** (P2, open, 4 comments)  
   A cron task added by the model can persist across sessions without any notification. Users report the assistant suddenly starting work in every new chat. The model also cannot list or cancel its own schedules.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5823)

5. **[#5920] `/rewind` records have `parentUuid: null`, breaking conversation history on resume** (P2, closed, 3 comments)  
   After a session resume, the entire history except the latest turn disappears because parent links were incorrectly stored. Fixed in PR #5923.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5920)

6. **[#5922] `cua‑driver.exe` keeps high CPU usage even after idle** (P2, open, 3 comments)  
   The Windows `cua‑driver.exe` process stays alive and consumes CPU long after the agent task completes. Users flag this as a resource‑hogging bug.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5922)

7. **[#5907] Tracking: complete Telegram bot command support** (P2, open, 3 comments)  
   An umbrella issue for aligning Telegram bot commands with the CLI command menu. Desired commands include `/cancel`, session management, and proper response handling.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5907)

8. **[#5897] Repeating “unknown format ‘uint64’ ignored in schema” messages** (P3, open, 3 comments)  
   On startup, Ajv emits dozens of schema‑validation warnings about `uint64` formats. While harmless, they pollute the interface and erode user confidence.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5897)

9. **[#5756] Default 8‑K output cap truncates large `write_file` calls, causing retry loops** (P2, open, 2 comments)  
   When the model writes a large file (e.g., full‑page wiki), the 8‑K token cap truncates the output. The tool then rejects the incomplete write, and the model retries indefinitely.  
   [Link](https://github.com/QwenLM/qwen-code/issues/5756)

10. **[#5929] Workflow stall env accepts non‑decimal seconds** (P3, open, 2 comments)  
    The env var `QWEN_CODE_WORKFLOW_STALL_SECONDS` parses `0x10`, `1e3`, etc., when only whole decimal integers are intended. Relaxed parsing could cause silent misconfiguration.  
    [Link](https://github.com/QwenLM/qwen-code/issues/5929)

## Key PR Progress (10 of 50)

1. **[#5777] feat: revive Chrome extension via daemon‑direct architecture** (open)  
   Rebuilds the Chrome extension as a thin client of the `qwen serve` daemon, removing the old Native Messaging stack. The side panel talks directly to the daemon over HTTP+SSE.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5777)

2. **[#5934] fix: stop repeated truncated `write_file`/edit retries from looping** (open)  
   Addresses the core issue in #5756 by enhancing the retry‑loop detector so it doesn’t feed truncated tool calls back to the same model, breaking the loop.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5934)

3. **[#5886] feat: add a git‑shared team memory tier** (open)  
   Introduces an opt‑in `TEAM` memory tier stored in `.qwen/team‑memory/` that is committed to git and shared across collaborators, alongside private `USER` and `PROJECT` tiers.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5886)

4. **[#5890] feat: inject `.qwen/loop.md` task file at fire time** (open)  
   Long‑running `/loop` cron tasks can now carry a durable task list via a `loop.md` file. The model re‑reads the file each tick, enabling the user to edit tasks without re‑stating the prompt.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5890)

5. **[#5921] feat: show scheduled task count in footer** (open)  
   Adds a footer pill that displays the number of active cron tasks, giving users visibility into background automation.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5921)

6. **[#5918] feat: warn before foreground shell timeout** (open)  
   When a foreground command is about to exceed its timeout, a warning is printed (with a reminder that Ctrl+B can promote it to background). Mitigates surprise terminations.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5918)

7. **[#5898] Fix mid‑input skill command completion** (closed)  
   Allows skill slash‑commands typed after other prompt text to trigger the same suggestion menu as line‑start commands. Now `/store` will match `front‑end‑store‑rules`.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5898)

8. **[#5931] feat: add workspace session sidebar to web shell** (open)  
   Implements a sidebar for starting, switching, renaming, and deleting sessions. Also supports local search and resizable width.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5931)

9. **[#5888] feat: qwen tag – multiplayer channel‑resident agent** (open)  
   An RFC and Phase‑0 implementation for a multi‑user agent that lives in chat groups (DingTalk‑first), built on `qwen serve` and existing channel adapters.  
   [Link](https://github.com/QwenLM/qwen-code/pull/5888)

10. **[#5925] fix: stop computer use driver when idle** (open)  
    Automatically stops the `cua‑driver` process after 5 minutes of inactivity, addressing the high‑CPU complaint in #5922. The timeout is configurable.  
    [Link](https://github.com/QwenLM/qwen-code/pull/5925)

## Feature Request Trends

- **Cross‑device / team synchronization**  
  Multiple requests ask for project‑local persistence of todos, plans, and memories so they can be committed to git and shared across machines or teammates (issues #5836, #5928, PR #5886).

- **Agent visibility and control**  
  Users want the ability to list, cancel, and adjust settings for cron tasks, shell‑command timeouts, and background automation. Footer pills, `/loop` task files, and timeout warnings all aim to give operators more transparency (issues #5823, #5838; PRs #5921, #5918).

- **Messaging platform integration**  
  Telegram bot support is being rounded out (issue #5907, PR #5919), and the multiplayer “qwen tag” concept (PR #5888) extends the idea to group chats.

- **Web shell enhancements**  
  Teams continue to push for better UX: live syntax highlighting for streaming code (issue #5866), a workspace session sidebar (PR #5931), and manual toggle for enhanced tables (PR #5917).

## Developer Pain Points

- **Truncated output loops**  
  The default 8‑K output token cap (`CAPPED_DEFAULT_MAX_TOKENS`) repeatedly truncates large `write_file` and `edit` tool calls, causing the model to retry endlessly. This is the single most disruptive bug reported this week (issue #5756, PR #5934).

- **cua‑driver process management**  
  The Windows `cua‑driver.exe` remains active and consumes CPU after the agent goes idle (issue #5922). The fix in PR #5925 adds an inactivity timeout, but users are waiting for it to merge.

- **Conversation‑state corruption**  
  `/rewind` records with `parentUuid: null` cause the entire chat history to vanish after a session resume (issue #5920, fixed in PR #5923).

- **Schema validation noise**  
  Ajv repeatedly logs “unknown format “uint64” ignored in schema” at startup. While non‑blocking, the messages clutter the terminal and alarm users (issue #5897, PR #5915).

- **Silent background automation**  
  Cron tasks fire without any notification, and the model cannot introspect its own schedules. Users discover stale tasks days later (issue #5823). The footer pill (PR #5921) is a welcome step toward visibility.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest – 2026-06-27

**Project:** CodeWhale (DeepSeek TUI)  
**Data Source:** [github.com/Hmbown/CodeWhale](https://github.com/Hmbown/CodeWhale)

---

## Today’s Highlights

No new release today, but the community saw a flurry of maintenance and feature work. The **Moraine** memory backend integration reached a critical milestone with multiple PRs wiring the MCP recall tools and gating legacy push logic. A persistent **plan/agent mode confusion** bug (#3568) remains under active discussion, while **editor freezes** (#3657) emerged as a new severity-1 issue. On the PR side, contributors landed improvements to **skills locale-awareness**, **Telegram conflict handling**, **CI reliability**, and a **new OpenModel provider** – many of them harvested from community submissions.

---

## Releases

**No releases in the last 24 hours.**

---

## Hot Issues

### 1. [#3568 – plan and agent mode mixed up YET AGAIN](https://github.com/Hmbown/CodeWhale/issues/3568)
🔴 **Open** – Created 2026-06-25 | 6 comments | 👍 1  
A user provides a concrete chat export showing the model still ignoring a “plan” mode switch and attempting file modifications despite being instructed to only plan. This recurring bug affects workflow reliability for users who rely on precise mode toggling. Community reactions suggest it’s a top-priority fix.

### 2. [#3495 – v0.8.66: Adopt Moraine as CodeWhale's memory backend](https://github.com/Hmbown/CodeWhale/issues/3495)
🟢 **Open** – Created 2026-06-23 | 4 comments  
The formal proposal to integrate **Moraine** (Apache 2.0) as the long-term agent-memory store, replacing the legacy push/inject mechanism. Multiple PRs this week are already implementing parts of this plan; it will significantly improve session recall and searchability.

### 3. [#3657 – Editor Freezes and Crashes Codewhale](https://github.com/Hmbown/CodeWhale/issues/3657)
🔴 **Open** – Created 2026-06-26 | 3 comments  
Precise reproduction steps: type `d` for draft mode, then `Ctrl-O` to open the editor – the app freezes entirely, requiring `kill -9`. A critical stability issue that is likely affecting many users.

### 4. [#3582 – install.sh endpoint returns HTML instead of shell script](https://github.com/Hmbown/CodeWhale/issues/3582)
🟢 **Closed** – Created 2026-06-25 | 4 comments  
The recommended `curl … | sh` installation command broke because the server returned a Next.js HTML page. This was quickly fixed (closed 2026-06-26) but highlights potential deployment risks.

### 5. [#3309 – v0.8.63: Split runtime_api.rs by API domain](https://github.com/Hmbown/CodeWhale/issues/3309)
🟢 **Closed** – Created 2026-06-18 | 2 comments  
A large refactoring effort to break apart the monolithic `runtime_api.rs` into per-domain modules. This was completed and merged, improving maintainability for the TUI backend. Several follow-up PRs today continue the decomposition.

### 6. [#1490 – bug: app-server --stdio thread/message does not stream response_delta](https://github.com/Hmbown/CodeWhale/issues/1490)
🟢 **Closed** – Created 2026-05-12 | 0 comments  
Long-standing bug where the stdio server returns `accepted` without streaming actual model output. Closed on 2026-06-26, likely as part of the recent server-side fixes.

---

## Key PR Progress

### 1. [#3691 – refactor(runtime-api): extract workspace status helpers](https://github.com/Hmbown/CodeWhale/pull/3691)
🟢 **Open** – Created 2026-06-27  
Continues the runtime API monolith split by moving workspace status types and helpers into `runtime_api/workspace.rs`. Behavior-preserving refactor with passing `cargo test`.

### 2. [#3690 – feat(skills): locale-aware skill descriptions to save tokens](https://github.com/Hmbown/CodeWhale/pull/3690)
🟢 **Open** – Created 2026-06-27  
Addresses [#3354](https://github.com/Hmbown/CodeWhale/issues/3354) by making skill descriptions localizable. In non-English sessions, English skill descriptions waste tokens; this PR loads descriptions in the user’s locale to reduce overhead.

### 3. [#3689 – fix(telegram): bound polling conflict retries](https://github.com/Hmbown/CodeWhale/pull/3689)
🟢 **Closed** – Created 2026-06-27  
Replaces a flat 10-second retry for Telegram `getUpdates` conflicts with a bounded exponential ladder (15-55s) and adds tests. Improves reliability when multiple bridges share a bot token.

### 4. [#3688 – refactor(runtime-api): extract session handlers](https://github.com/Hmbown/CodeWhale/pull/3688)
🟢 **Closed** – Created 2026-06-27  
Extracts session request/response types and session handlers into `runtime_api/sessions.rs` while preserving centralized route registration.

### 5. [#3684 – docs(contributing): fix stale file paths in PR exemplars](https://github.com/Hmbown/CodeWhale/pull/3684)
🟢 **Open** – Created 2026-06-27  
Harvests @findshan’s fix for outdated file paths in `CONTRIBUTING.md` (PR #3680). Ensures new contributors don’t land on 404 pages when following example commands.

### 6. [#3686 – ci: run rlm cache workflow tests](https://github.com/Hmbown/CodeWhale/pull/3686)
🟢 **Closed** – Created 2026-06-27  
Adds CI coverage for the RLM cache change workflow using a mock provider, ensuring the WhaleFlow authoring and execution path remains deterministic.

### 7. [#3685 – feat: Add skills inspect and /skills --inspect](https://github.com/Hmbown/CodeWhale/pull/3685)
🟢 **Closed** – Created 2026-06-27  
Introduced a diagnostic command (`/skills inspect`) that reports discovery mode, workspace, searched directories, and available skill count. Helps users debug skill loading issues. (Also harvested into #3687.)

### 8. [#3585 – Add OpenModel provider support](https://github.com/Hmbown/CodeWhale/pull/3585)
🟢 **Closed** – Created 2026-06-25  
First-class integration of OpenModel as a provider, routing via the Anthropic Messages API. Defaults to `deepseek-v4-flash`. Expands user choice beyond the standard providers.

### 9. [#3575 – feat(memory): wire moraine-mcp as recall tool source, gate legacy push/inject](https://github.com/Hmbown/CodeWhale/pull/3575)
🟢 **Closed** – Created 2026-06-25  
A major piece of the Moraine adoption: adds the Moraine MCP server as a default recall tool source and introduces a `moraine_fallback` config gate to disable the legacy push mechanism.

### 10. [#3645 – Guard exec against misplaced global flags](https://github.com/Hmbown/CodeWhale/pull/3645)
🟢 **Closed** – Created 2026-06-26  
Prevents `codewhale exec --provider not-a-provider "prompt"` from silently misinterpreting the flag as prompt text. Adds a guard that validates arguments before the prompt.

---

## Feature Request Trends

- **Moraine memory backend** is the dominant feature direction this week. Issues and PRs around adopting Moraine as the primary session memory store (search, recall, session management) are being actively merged.
- **Locale/skill awareness** – Users in non-English environments want skill descriptions to be localized to save tokens and improve relevance. PR #3690 directly addresses this.
- **Provider expansion** – OpenModel integration (#3585) and improved docs for providers like Qianfan (#3621) show a desire for more provider choice.
- **Diagnostics and debugging** – The new `/skills inspect` command (#3685) and the planned CI improvements suggest a push toward better observability for power users.

---

## Developer Pain Points

- **Plan/agent mode confusion** (#3568) remains a top frustration. Users report that even with explicit mode switches, the model ignores the intended behavior, undermining trust in AI-assisted planning.
- **Editor freezing** (#3657) is a new severe bug that crashes the entire application. No workaround is known, making the TUI unusable for affected users.
- **Installation failures** (#3582) – The `install.sh` endpoint returning HTML caused immediate onboarding friction. Although fixed, it shows the need for better deployment tests.
- **Documentation staleness** – Multiple PRs (#3680, #3684) were needed to fix outdated file paths in `CONTRIBUTING.md`, a sign that example-driven documentation is not kept in sync with code reorganization.
- **Telegram bot conflict retries** (#3689) – The default polling retry mechanism was too simplistic, causing silent failures when multiple bots contended for the same token.
- **Global flag handling** (#3645) – Users inadvertently misplacing flags after `exec` led to confusing behaviour, requiring a guard to prevent silent misinterpretation.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*