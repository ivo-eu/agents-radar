# AI CLI Tools Community Digest 2026-06-18

> Generated: 2026-06-18 12:31 UTC | Tools covered: 9

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
**Date:** 2026-06-18 | **Data Sources:** Community digests for Claude Code, OpenAI Codex, Gemini CLI, GitHub Copilot CLI, Kimi Code CLI, OpenCode, Pi, Qwen Code, and DeepSeek TUI (CodeWhale)

---

## 1. Ecosystem Overview

The AI CLI tools ecosystem is in a phase of rapid maturation and fragmentation. Nine actively maintained tools now compete to define developer workflows for agentic coding, with **MCP integration**, **agent orchestration reliability**, and **cross-platform consistency** emerging as universal friction points. Two parallel market forces are shaping the landscape: enterprise tools (Claude Code, Codex, Copilot) are investing in security scoping, remote execution, and custom model support, while open‑source alternatives (OpenCode, Pi, Qwen Code) race to match feature parity while building distinct extensibility layers. Community engagement remains intense — top issues routinely gather triple‑digit reactions — but tool maturity varies widely, from Kimi’s sparse activity to OpenCode’s blistering issue/PR throughput.

---

## 2. Activity Comparison

| Tool | Top Issues Counted | PRs Updated (24h) | New Release Today | Release Version(s) |
|---|---|---|---|---|
| Claude Code | 10 | 5 | ✅ v2.1.181 | v2.1.181 |
| OpenAI Codex | 10 | 10 | ✅ 4 releases | rust‑v0.141.0, alpha builds |
| Gemini CLI | 10 | 10 | ✅ 2 releases | v0.47.0, v0.48.0-preview.0 |
| GitHub Copilot CLI | 10 | 1 | ❌ | – |
| Kimi Code CLI | 2 | 0 | ❌ | – |
| OpenCode | 10 | 10 | ✅ v1.17.8 | v1.17.8 |
| Pi | 10 | 10 | ❌ (PRs merged) | – |
| Qwen Code | 10 | 10 | ✅ 3 releases | v0.18.3, nightly, preview |
| DeepSeek TUI (CodeWhale) | 10 | 10 | ✅ v0.8.62 | v0.8.62 |

**Key observations:**
- **Codex** and **Qwen Code** are the most prolific release engines, with multiple versions per day.
- **Copilot CLI** and **Kimi** show minimal PR activity, suggesting slower development cycles or internal pipeline focus.
- **OpenCode** and **Pi** maintain high PR throughput even without formal releases, indicating active merging of community contributions.

---

## 3. Shared Feature Directions

### 3.1 MCP Integration & Tool Reliability
*Appears in:* **All tools** (Claude Code, Codex, Copilot, Kimi, OpenCode, Pi, Qwen Code, DeepSeek)

| Specific Need | Tools Affected |
|---|---|
| OAuth token auto‑refresh | Codex (#17265), Copilot (#3838) |
| Sub‑agent MCP tool access | Copilot (#3812), Gemini (#21409) |
| Schema validation ($ref, object serialization) | OpenCode (#28472, #32829), Qwen Code (#5322) |
| Session recovery on non‑404 errors | OpenCode (#32809) |
| Image/inline UI rendering | Codex (#21019), OpenCode (#32832) |
| Simplified onboarding (wiring MCP+plugins) | Kimi (#2460) |

### 3.2 Agent Reliability & Lifecycle Controls
*Appears in:* **All tools** except Kimi

| Specific Need | Tools Affected |
|---|---|
| Prevent recursive self‑spawning | Claude Code (#69332), Gemini (#22323) |
| Session continuity after stall/cancel | DeepSeek (#2739), Claude Code (#46561) |
| Context‑window health monitoring | Claude Code (#66807), Codex (#9046) |
| Programmatic effort/model switching | Claude Code (#59502), Copilot (#3074) |
| Background agent usage budgets | Codex (PR #28707), Claude Code (#69332) |

### 3.3 Security & Sandboxing
*Appears in:* **Claude Code, Codex, OpenCode, Qwen Code, Gemini**

| Specific Need | Tools Affected |
|---|---|
| Directory/file access restrictions | OpenCode (#2242), Qwen Code (#5316) |
| Network approval scoping by environment | Codex (PR #28899) |
| Secret redaction before model context | Gemini (#26525), Qwen Code (#5326) |
| Windows sandbox compatibility | Codex (#18918), Copilot (#3849) |
| Configurable shell permission persistence | DeepSeek (PR #3301) |

### 3.4 Cross‑Platform Consistency
*Appears in:* **All tools** (especially Linux/Windows)

| Specific Need | Tools Affected |
|---|---|
| XDG Base Directory compliance | Claude Code (#1455, #2350) – 383 👍 |
| Windows git operations in sandbox | Codex (#18918), Copilot (#3849) |
| Terminal‑specific bugs (Kitty, Wayland, Warp) | Gemini (#21983), Pi (#5407, #5827) |
| OneDrive/mapped‑drive persistence | Claude Code (#14088) |

### 3.5 Custom Model & Provider Flexibility
*Appears in:* **Copilot, OpenCode, Pi, Qwen Code, DeepSeek**

| Specific Need | Tools Affected |
|---|---|
| Enterprise‑managed custom models | Copilot (#3730), Qwen Code (#5303) |
| Ollama/BYOK compatibility | Copilot (#3839), OpenCode (#32829) |
| Multi‑provider config (same ID, different keys) | Qwen Code (#5303), Pi (#5866) |
| GLM‑5.2 / Azure Foundry support | Pi (#5873, #5849), OpenCode (#32833) |

### 3.6 Developer Experience & Observability
*Appears in:* **Codex, Claude Code, OpenCode, Qwen Code**

| Specific Need | Tools Affected |
|---|---|
| Rate‑limit reset times & plan info in TUI | Codex (#24080, #28908) |
| Structured timestamps for model reasoning | Claude Code (#44763, #49084) |
| Hot‑reload configs (agents, skills, commands) | OpenCode (#8751) – 77 👍 |
| Input validation with clear error messages | Qwen Code (5+ issues), OpenCode (#28472) |
| Comment preservation in config files | DeepSeek (#3282), OpenCode (#32824) |

---

## 4. Differentiation Analysis

| Tool | Primary Differentiator | Target User | Technical Approach |
|---|---|---|---|
| **Claude Code** | Agentic workflows (Cowork, scheduled tasks, context health) | Power developers, long‑running agent operators | Node.js, TUI‑first, heavy prompt engineering |
| **OpenAI Codex** | Remote execution with encrypted relays, security scoping | Enterprise, multi‑platform, VS Code users | Rust, Noise protocol, VSCode extension |
| **Gemini CLI** | Evaluation framework (component‑level evals), Auto Memory | Google Cloud developers, test engineers | TypeScript/Node, behavioral eval pipelines |
| **GitHub Copilot CLI** | GitHub ecosystem integration, BYOK, session management | GitHub enterprise users, custom model adopters | Node.js, SDK + server mode |
| **Kimi Code CLI** | Simplicity & fast onboarding | New users, quick‑start scenarios | Minimalist CLI, low configuration surface |
| **OpenCode** | Extensibility (agents, skills, commands), MCP debugging | Terminal enthusiasts, power users, multi‑provider | TypeScript, plugin architecture |
| **Pi** | Terminal compatibility, RPC for external tooling | Extension builders, multi‑terminal developers | TypeScript, TUI‑native, provider‑agnostic |
| **Qwen Code** | Chinese market features (QQ bot, WeChat), sandboxing | Asia‑based developers, security‑conscious | TypeScript, channel adapter architecture |
| **DeepSeek TUI (CodeWhale)** | Policy engine (permissions, scope discipline), plan/agent modes | DeepSeek users, strict agent‑control workflows | Rust, TUI‑first, constitution‑based governance |

**Release cadence (24h):** Codex (4) > Qwen Code (3) > Gemini, Claude Code, OpenCode, DeepSeek (1–2) > Copilot, Pi, Kimi (0)

**Community engagement intensity:** Claude Code (383 👍 on single issue) > OpenCode (77 👍 on hot‑reload) > Codex (51 👍 on history regression) > Gemini (8 👍 on hang bug) > others.

---

## 5. Community Momentum & Maturity

**High momentum / rapid iteration:**
- **OpenCode** – 10 PRs + 1 release today, top request (hot‑reload) at 77 👍, 73‑comment sandbox issue. MCP bugs dominate but community patches are fast.
- **OpenAI Codex** – 10 PRs + 4 releases, strong security and observability features (rollout budgets, time reminders). VSCode regression (#18993, 51 👍) is a maturity gap.
- **Qwen Code** – 10 PRs + 3 releases, aggressive input validation fixes. Multi‑agent stability (#5180) is the biggest risk.
- **Pi** – 10 PRs, responsive maintainers (fast closure of critical bugs), building extension ecosystem.

**Mature with occasional regressions:**
- **Claude Code** – Largest community (383 👍 on XDG). Slow to fix long‑standing Linux complaints, but releases are steady. Agent runaway bugs erode trust.
- **Gemini CLI** – Stable releases but agent hangs (#21409) and false goal success (#22323) are top blockers. Evaluation framework is a unique strength.
- **Copilot CLI** – Enterprise trust but single PR in 24h, login loop (#254) unresolved for 4 months. MCP integration still fragile.

**Early‑stage / lower activity:**
- **Kimi Code CLI** – Only 2 issues today, no PRs. Rebranded/quiet period. Low community engagement.
- **DeepSeek TUI (CodeWhale)** – Rebranding complete, flurry of fixes (10 PRs). Still battling core stall errors and mode‑switching bugs. Promising foundation for v0.9.0.

---

## 6. Trend Signals

1. **MCP is becoming the universal integration protocol** – Every tool now supports it, but implementation quality varies wildly. OAuth handling, schema validation, and sub‑agent propagation are the biggest gaps. Expect a stabilization wave across all tools in Q3.

2. **Agent reliability is the #1 functional demand** – Self‑spawning loops, false success reports, context overruns, and session continuity failures dominate high‑reaction issues. Tools that solve this (e.g., through policy engines, cron‑based wakeups, or budget enforcement) will gain trust.

3. **Security sandboxing is shifting from nice‑to‑have to must‑have** – With agents gaining file system and network access, enterprises are demanding scoped execution environments (network approval per environment, directory jailing, secret redaction). This is emerging as a key buying criterion.

4. **Cross‑platform consistency remains a deferred debt** – Linux XDG compliance (Claude Code, 383 👍) and Windows sandbox/git issues (Codex, Copilot) have been open for months or years. Developers increasingly expect CLI tools to “just work” on any platform.

5. **Custom model/provider flexibility is growing** – BYOK, Ollama, Azure Foundry, and private endpoints are being requested across enterprise and open‑source tools. The era of single‑model dependency is ending.

6. **Observability is undervalued today but will be critical** – Rate‑limit visibility, structured timestamps, context‑window health, and cost confirmation are recurring feature requests. Tools that surface this data in the TUI will reduce cost surprises and improve debugging.

7. **Terminal and TUI customization is a competitive edge** – From cursor blink control to theme auto‑detection to image protocol support (Warp, Kitty), developers are increasingly picky about their terminal experience. Tools that ignore platform‑specific terminal quirks lose users to alternatives.

8. **Input validation gaps are a silent trust killer** – Qwen Code alone had 5+ issues where loose parsing (e.g., `parseInt` on `"1.5"`, regex on remote hosts) led to security holes or data corruption. Expect stricter validation and schema enforcement to become a hygiene standard.

---

*Report generated from community digest data as of 2026-06-18. All references link to the respective GitHub issues and PRs cited in the source digests.*

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
_Data as of 2026-06-18 · Source: github.com/anthropics/skills_

---

## 1. Top Skills Ranking

The following pull requests have drawn the most community discussion and attention. All remain **open** as of the data snapshot.

| Rank | PR | Skill / Change | Discussion Highlights |
|------|----|----------------|----------------------|
| 1 | [#514](https://github.com/anthropics/skills/pull/514) | **document-typography** – quality control for AI-generated documents (orphan word wrap, widow paragraphs, numbering alignment). | Strong agreement that typographic issues are pervasive in Claude output. Community wants this as a core skill. |
| 2 | [#486](https://github.com/anthropics/skills/pull/486) | **ODT skill** – create, fill, read, and convert OpenDocument files (.odt, .ods); includes template filling and ODT→HTML parsing. | High demand for open‑standard document formats. Discussion focuses on completeness of ISO format support and LibreOffice compatibility. |
| 3 | [#210](https://github.com/anthropics/skills/pull/210) | **Improve frontend-design skill** – rewrite for clarity, actionability, and single‑conversation feasibility. | Debate over balancing brevity with specificity; several alternate rewrites proposed in comments. |
| 4 | [#83](https://github.com/anthropics/skills/pull/83) | **skill-quality-analyzer + skill-security-analyzer** – meta‑skills that evaluate other skills across structure, documentation, security, and more. | First attempt at quality gates for the ecosystem. Community divided on whether meta‑skills should be part of the official collection or live in a separate registry. |
| 5 | [#538](https://github.com/anthropics/skills/pull/538) | **fix(pdf): case‑sensitive file references** – corrects 8 mismatches between SKILL.md and actual lowercase filenames. | Exposes wider issue: the repo’s PDF skill is known to break on Linux and macOS due to case sensitivity. PR is uncontroversial but delayed by need for maintainer review. |
| 6 | [#539](https://github.com/anthropics/skills/pull/539) | **fix(skill-creator): warn on unquoted YAML description** – adds pre‑parse validation for unquoted `:` in description fields. | Linked to #361. Community points out that silent YAML truncation has burned many contributors; this fix is seen as a low‑risk, high‑impact improvement. |
| 7 | [#541](https://github.com/anthropics/skills/pull/541) | **fix(docx): prevent w:id collision with existing bookmarks** – stops document corruption when tracked changes are added to documents that already use bookmarks. | Root‑cause analysis well received. Several users confirmed the corruption bug in their own DOCX workflows. |
| 8 | [#1298](https://github.com/anthropics/skills/pull/1298) | **fix(skill-creator): run_eval.py always reports 0% recall** – overhauls the evaluation pipeline to install the artifact as a real skill, fix Windows stream reading, and improve trigger detection. | Directly addresses the blocking issue #556 (12 comments) and #1169. The most active PR this month with parallel comment threads on Windows and parallel‑worker logic. |

---

## 2. Community Demand Trends

From the most‑commented issues (13 total, sorted by activity), three clear demand vectors emerge:

1. **Infrastructure reliability** – Issues [#556](https://github.com/anthropics/skills/issues/556), [#1169](https://github.com/anthropics/skills/issues/1169), and [#1061](https://github.com/anthropics/skills/issues/1061) all report that `run_eval.py` (and consequently the skill‑creator optimisation loop) returns 0% recall on every query. This blocks anyone from using the official skill development tools. Windows compatibility is a recurring sub‑theme ([#1061](https://github.com/anthropics/skills/issues/1061), [#1099](https://github.com/anthropics/skills/pull/1099)).  
2. **Skill sharing & governance** – [#228](https://github.com/anthropics/skills/issues/228) (14 comments, 7 👍) requests org‑wide skill sharing without manual file transfer. [#492](https://github.com/anthropics/skills/issues/492) (7 comments) raises security concerns about the `anthropic/` namespace being used for community skills, a trust‑boundary vulnerability.  
3. **New skill directions** – The most‑requested new skill area is **agent governance** ([#412](https://github.com/anthropics/skills/issues/412): policy enforcement, threat detection, audit trails). **Document processing** (ODT, typography) and **enterprise platform integration** (SAP, ServiceNow) also see strong demand from both PRs and issues. Duplicate skill conflicts ([#189](https://github.com/anthropics/skills/issues/189)) hint at the need for a skill registry or deduplication tool.

---

## 3. High‑Potential Pending Skills

These PRs are actively discussed, recently updated, and likely to land soon:

- **#1298 – fix(skill-creator) recall 0%** – Last updated 2026‑06‑11. The maintainer has requested a second reviewer; several users have tested the fix on Linux and Windows. Resolving the evaluation bottleneck unblocks all other skill‑creator improvements.
- **#361 – Detect unquoted YAML special characters** – Last updated 2026‑06‑10. Broader than #539 (it also catches `#`, `{`, `[`). Already has partial consensus; waiting on the maintainer to confirm a preferred integration approach.
- **#1099 – fix run_eval Windows crash** – Last updated 2026‑05‑24. A targeted fix for the `[WinError 10038]` pipe issue. Community feedback is positive; merge may be fast‑tracked once #1298 is resolved.
- **#514 – document-typography skill** – Last updated 2026‑03‑13 (stale by date but still high discussion). The concept is universally welcomed; the PR needs a commit to incorporate community suggestions on edge cases (e.g., CJK text, justified alignment).

---

## 4. Skills Ecosystem Insight

**The community’s most concentrated demand is for a reliable, Windows‑compatible skill‑development pipeline that produces high‑quality document‑processing skills (typography, ODT, PDF, DOCX) and enterprise‑grade agent memory/governance capabilities, all while addressing the trust and sharing gaps in the current ecosystem.**

---

*All data sourced from github.com/anthropics/skills. PRs and issues referenced by number link directly to GitHub.*

---

# Claude Code Community Digest — June 18, 2026

## Today's Highlights

A modest release (v2.1.181) brings prompt‑based `/config` syntax—a quality‑of‑life win for power users—and an opt‑in Apple Events sandbox setting. The community remains vocal about XDG Base Directory compliance (the top‑voted open issue at 383 👍) while a fresh batch of bugs around agent self‑spawning and malformed tool calls on Opus 4.7 signals growing pains as agentic workloads expand.

## Releases

**v2.1.181** — [Release notes](https://github.com/anthropics/claude-code/releases/tag/v2.1.181) *(partial changelog)*  
- `/config key=value` syntax to set any setting from the prompt (works in interactive, `-p`, and Remote Control).  
- `sandbox.allowAppleEvents` opt‑in setting for sandboxed commands on macOS.  
- Added `CLAUDE_CLIENT_P` (detail truncated in source; likely a client‑side prefix or path variable).  

No other releases in the last 24 hours.

## Hot Issues

1. **[#1455 – XDG Base Directory spec ignored](https://github.com/anthropics/claude-code/issues/1455)** — Linux users demand Claude Code stop writing configs/cache to `~/.claude`. 383 👍, 61 comments, open since May 2025. The highest‑engagement open bug/enhancement.

2. **[#44763 – Timestamps on conversation messages](https://github.com/anthropics/claude-code/issues/44763)** — Needed for monitoring long‑running agents. 46 👍, 33 comments.

3. **[#14088 – Chat history lost on mapped drives / OneDrive](https://github.com/anthropics/claude-code/issues/14088)** — Windows‑specific persistence bug, 12 👍, 30 comments.

4. **[#2350 – Follow XDG properly (move cache out of config)](https://github.com/anthropics/claude-code/issues/2350)** — Duplicate of #1455 with concrete spec citations, 88 👍, 20 comments.

5. **[#12962 – Settings.json parent directory traversal for monorepos](https://github.com/anthropics/claude-code/issues/12962)** — Shared settings at repo root not discovered from subdirectories. 62 👍, 19 comments.

6. **[#8504 – Disable user input background highlighting](https://github.com/anthropics/claude-code/issues/8504)** — Accessibility request to customize TUI appearance. 20 👍, 13 comments.

7. **[#54859 – Cowork: configure scheduled tasks storage location](https://github.com/anthropics/claude-code/issues/54859)** — Currently hardcoded to `~/Documents/Claude/Scheduled/`. 3 👍, 13 comments.

8. **[#49084 – Expose timestamps as structured data](https://github.com/anthropics/claude-code/issues/49084)** — Time‑aware reasoning for agents. 3 👍, 12 comments.

9. **[#66807 – Context‑health monitoring](https://github.com/anthropics/claude-code/issues/66807)** — New feature request (June 10) for TUI context window health; 8 comments.

10. **[#69332 – Sub‑agent recursive self‑spawn → exponential fan‑out](https://github.com/anthropics/claude-code/issues/69332)** — Critical bug reported today: background agents exhaust usage limits via runaway spawning. 2 comments. **High severity.**

## Key PR Progress

Only five PRs were updated in the last 24 hours. Notable among them:

1. **[#41611 – Add the missing source to Claude Code](https://github.com/anthropics/claude-code/pull/41611)** — Open, by tornikeo. Formatting/metadata fix, updated June 18.

2. **[#41447 – Open source Claude Code ✨](https://github.com/anthropics/claude-code/pull/41447)** — Open, by gameroman. Claims to close #59, #456, #2846, #22002, #41434. No comments, last updated June 18. Ambitious statement.

3. **[#69226 – Update frontend‑design skill](https://github.com/anthropics/claude-code/pull/69226)** — Merged/closed June 17–18. Bumps plugin version to 1.1.0 with improvements.

4. **[#19867 – Fix code‑review: re‑review on new commits](https://github.com/anthropics/claude-code/pull/19867)** — Open since January 2026. Adds smarter skip logic and `--force` flag. Reviewed June 17.

5. **[#33443 – Update Dockerfile to use native installer](https://github.com/anthropics/claude-code/pull/33443)** — Open since March. Replaces deprecated npm install with Node 24.14 native installer. Updated June 17.

## Feature Request Trends

- **XDG Base Directory compliance** — Dominant theme on Linux (#1455, #2350). Users expect cache and config separation per freedesktop.org spec.
- **Monorepo/workspace awareness** — Parent directory traversal for settings (#12962) and shared skill inheritance.
- **Time awareness** — Both visual timestamps (#44763) and structured timestamps for model reasoning (#49084).
- **Context health UI** — Monitoring token usage, message age, and context window saturation (#66807).
- **Agent lifecycle controls** — Inherit context for sub‑agents (#69283), programmatic `/effort`/`/model` switching (#59502).
- **Storage flexibility** — Configurable paths for Cowork scheduled tasks (#54859) and general state/cache (#1455, #2350).

## Developer Pain Points

- **Cross‑platform inconsistency** — XDG violation (Linux) and OneDrive/mapped‑drive persistence failures (Windows) show platform testing gaps.
- **Agent runaway behavior** — Recursive sub‑agent spawning (#69332) and stale loop resurrections after `/compact` (#46561) erode trust.
- **Cost surprises** — `/ultrareview` triggered by default Return with no cost confirmation (#60635) and background agent usage limit exhaustion (#69332) frustrate users.
- **Tool‑call reliability** — Opus 4.7 malformed tool calls above 150k context (#68472) and file truncation on skill uploads >99 KB (#51435) hurt advanced workflows.
- **OAuth pain** — MCP OAuth tokens not honored after successful dance (#60260), blocking remote MCP adoption.
- **Session management friction** — Cannot preview old sessions before committing to resume (#60484); chat history lost on non‑standard paths (#14088).

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest – 2026-06-18

## Today’s Highlights
Three major themes dominate today’s digest: **security-scoped network approvals** and **rollout token budgets** landed in core PRs, while the community remains vocal about **VSCode extension regressions**, **MCP OAuth token management**, and **Windows sandbox compatibility**. A new rust-v0.141.0 stable release introduces authenticated, end-to-end encrypted relay channels for remote executors, alongside cross-platform working‑directory preservation.

## Releases
Four new versions shipped in the last 24 hours:

- **[rust-v0.141.0](https://github.com/openai/codex/releases/tag/rust-v0.141.0)** – Stable release. Remote executors now use authenticated, end-to-end encrypted Noise relay channels. Cross‑platform remote execution preserves executor‑native working directories and shells, including filesystem permission paths across app‑server and exec‑server boundaries.
- **[rust-v0.142.0-alpha.1](https://github.com/openai/codex/releases/tag/rust-v0.142.0-alpha.1)** – Alpha with further improvements.
- **[rust-v0.141.0-alpha.6](https://github.com/openai/codex/releases/tag/rust-v0.141.0-alpha.6)** / **[rust-v0.141.0-alpha.7](https://github.com/openai/codex/releases/tag/rust-v0.141.0-alpha.7)** – Additional alpha iterations.

## Hot Issues
Selected 10 issues by community engagement (comments, reactions) and impact:

1. **[Issue #7291](https://github.com/openai/codex/issues/7291) – VSCode extension failed to revert changes**  
   *45 comments, 13 👍*. Users report that the VS Code extension (v0.4.46) cannot revert changes on macOS. High likelihood of affecting daily editing workflows.

2. **[Issue #18993](https://github.com/openai/codex/issues/18993) – Unable to open past conversation history in VS Code**  
   *30 comments, 51 👍*. Regression in IDE extension (v1.117.0) – past conversations are inaccessible. Heavy community demand for a fix.

3. **[Issue #9046](https://github.com/openai/codex/issues/9046) – Context window runs out immediately**  
   *30 comments, 1 👍*. Even fresh threads quickly hit the context window limit. Points to either aggressive prompt accumulation or a need for smarter context management.

4. **[Issue #17265](https://github.com/openai/codex/issues/17265) – MCP OAuth tokens not auto-refreshed**  
   *15 comments, 29 👍*. Codex persists refresh tokens but does not use them automatically, causing all MCP tool calls to fail after the access token expires. A critical blocker for MCP integrations.

5. **[Issue #14601](https://github.com/openai/codex/issues/14601) – Configuration pollution from `trusted_level`**  
   *14 comments, 43 👍*. Proposes separating `projects.xxxx.trusted_level` from `config.toml` to prevent unchecked propagation of trust decisions across projects. Strong community support.

6. **[Issue #20683](https://github.com/openai/codex/issues/20683) – Computer Use crashes when inspecting Outlook on macOS**  
   *12 comments, 2 👍*. Calling `get_app_state` on Outlook triggers a crash in `SkyComputerUseService`. Affects macOS users of Computer‑Use automation.

7. **[Issue #20567](https://github.com/openai/codex/issues/20567) – Windows app spawns ~1000 git commands per minute**  
   *10 comments, 1 👍*. Enterprise user reports infinite git process spawning, causing severe CPU/IO load. Likely a polling or file‑watcher bug.

8. **[Issue #18918](https://github.com/openai/codex/issues/18918) – Windows sandbox DENY ACLs on `.git` directories**  
   *9 comments, 5 👍* (Pro subscription). The sandbox applies deny permissions to `.git` folders inside writable roots, breaking `git commit`. Blocks development on Windows.

9. **[Issue #24080](https://github.com/openai/codex/issues/24080) – Expose rate‑limit reset times, balance, plan in TUI**  
   *8 comments, 0 👍*. Feature request to surface richer rate‑limit data (`resetsAt`, credits, plan type) in the CLI status line. Very low thumbs but two similar follow‑ups (#28908, #28896) highlight a broader need.

10. **[Issue #21019](https://github.com/openai/codex/issues/21019) – Codex Desktop does not render MCP Apps inline UI**  
    *8 comments, 13 👍*. MCP tools return valid URIs but the desktop app never calls `read‑mcp‑resource`, so inline iframe UIs are invisible. Regression or missing feature.

## Key PR Progress
Ten pull requests shaping the next releases:

1. **[PR #28899](https://github.com/openai/codex/pull/28899) – Scope network approvals by environment**  
   Prevents host‑level network approvals from leaking across execution environments (e.g., sandbox vs. no sandbox). Enhances security granularity.

2. **[PR #28707](https://github.com/openai/codex/pull/28707) – Interrupt sessions when rollout budgets expire**  
   Stops a multi‑thread session when the shared rollout token budget is exhausted. Part of the `varlength` stack that caps total usage per rollout.

3. **[PR #28494](https://github.com/openai/codex/pull/28494) – Rollout budget implementation**  
   Implements shared budget accounting with model‑visible reminders. Work is split across three PRs (#28746, #28494, #28707).

4. **[PR #28895](https://github.com/openai/codex/pull/28895) – Recover exec process stdin writes**  
   Fixes a crash when a remote stdio MCP server’s websocket drops during `process/write`. Previously the session recovery would fail; now it retries gracefully.

5. **[PR #28835](https://github.com/openai/codex/pull/28835) – Add app‑server current‑time implementation**  
   New `currentTime/read` RPC method to let the server query the client’s current time. First step of the `varlatency` feature for time‑based reminders.

6. **[PR #28824](https://github.com/openai/codex/pull/28824) – Current time reminders for system clock**  
   Adds a host‑injectable time provider and records time reminders in history before model requests. Configurable cadence per session.

7. **[PR #28822](https://github.com/openai/codex/pull/28822) – Config for time reminders**  
   Defines `[features.current_time_reminder]` with options: `enabled`, `reminder_interval_model_requests`, `clock_source`. Foundation for all `varlatency` work.

8. **[PR #28843](https://github.com/openai/codex/pull/28843) – Persist fsmonitor status refreshes**  
   Fixes a subtle Git performance bug: when Git’s built‑in fsmonitor is active, Codex no longer passes `--no‑optional‑locks`, allowing the daemon to persist its state.

9. **[PR #28489](https://github.com/openai/codex/pull/28489) – Index‑gated web search mode**  
   Introduces `web_search = "index_gated"` as a new mode that allows web access only when an index is present. Useful for controlled retrieval.

10. **[PR #28720](https://github.com/openai/codex/pull/28720) – Remove the core skills runtime**  
    After moving skill catalog rendering and explicit invocation into `SkillsExtension`, this PR removes the old core‑level skills runtime, simplifying the architecture.

## Feature Request Trends
- **Rate‑limit visibility**: Multiple issues (#24080, #28908, #28896) ask for raw reset times, balance, and plan info in both TUI and desktop UI.
- **MCP App inline rendering**: Users consistently want MCP apps to render their inline UIs (iframe, tool results) – Issue #21019 and #28900 (regression).
- **Multi‑agent mode**: PRs #28884 and #28792 expose per‑thread multi‑agent selection. Community expects this to become a first‑class setting.
- **Mobile/headless remote**: Issue #23200 (17👍) requests direct Linux SSH support on mobile without requiring the desktop app.
- **Chrome tab group management**: Issue #28907 asks Codex to expose Chrome tab group operations (rename, regroup) via the Chrome extension.
- **Project‑local override files**: PR #23547 adds `config.override.toml` – a direct response to developers needing private settings without modifying repo‑checked `config.toml`.

## Developer Pain Points
- **VSCode extension instability** – Issues #7291 and #18993 show persistent regressions in revert and history access, with high community anger (51👍 on #18993).
- **MCP OAuth token handling** – No auto‑refresh (#17265) and lack of support for pre‑registered clients (#19154) frustrate users running private MCP servers.
- **Windows sandbox / git incompatibilities** – DENY ACLs on `.git` (#18918) and infinite git spawning (#20567) make Codex nearly unusable on Windows for git‑heavy workflows.
- **Computer‑Use crashes** – macOS Outlook inspection (#20683) and Windows `@oai/sky` package export errors (#28713, #28717) block automation on both platforms.
- **Context window limits** – Issue #9046 highlights that even trivial conversations run out of context, suggesting aggressive prompt growth or insufficient compression.
- **Usage limit decreases without activity** – Issues #28908 and #28896 report unexplained drops in rate‑limit percentages, causing confusion and potential service interruptions.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-06-18

## Today's Highlights
Two new releases landed: **v0.47.0** (stable) with backend definition respect, and **v0.48.0-preview.0** (preview) adding a dependabot cooldown for npm. The community is most vocal about agent reliability issues — a critical bug causing the generalist agent to hang indefinitely remains under retesting (👍8), and a new PR fixes `write_file` corruption of Jupyter/JSON files. Security hardening is also in focus, with several CI workflow gates now patched.

## Releases
- **[v0.47.0](https://github.com/google-gemini/gemini-cli/releases/tag/v0.47.0)** — Stable release. Notable change: “Respect backend def”. Other changes are release automation and version bumps.
- **[v0.48.0-preview.0](https://github.com/google-gemini/gemini-cli/releases/tag/v0.48.0-preview.0)** — Preview release. Includes a cooldown period for npm dependency updates (Dependabot) and refactoring.

## Hot Issues
*10 selected by priority, comment count, and community reaction.*

1. **[#24353](https://github.com/google-gemini/gemini-cli/issues/24353) — Robust component level evaluations (p1)**  
   Epic following up on behavioral evals. 7 comments. High priority for ensuring test quality across agent components.

2. **[#22745](https://github.com/google-gemini/gemini-cli/issues/22745) — Assess AST-aware file reads, search, and mapping (p2)**  
   Investigates potential token savings and precision gains from AST-aware tools. 7 comments, 1 👍.

3. **[#21409](https://github.com/google-gemini/gemini-cli/issues/21409) — Generalist agent hangs (p1)**  
   Agent hangs indefinitely when deferring to sub-agents. 7 comments, 8 👍 (most upvoted open bug). A top pain point.

4. **[#22323](https://github.com/gemini-cli/cli/issues/22323) — Subagent recovery after MAX_TURNS reported as GOAL success (p1)**  
   `codebase_investigator` falsely reports success when hitting turn limit. 6 comments, 2 👍. Misleading diagnostics.

5. **[#21968](https://github.com/gemini-cli/cli/issues/21968) — Gemini does not use skills and sub-agents enough (p2)**  
   Anecdotal report that custom skills are rarely invoked automatically. 6 comments. Requesting better agent orchestration.

6. **[#26525](https://github.com/google-gemini/gemini-cli/issues/26525) — Add deterministic redaction and reduce Auto Memory logging (p2)**  
   Security concern: secrets sent to model before redaction. 5 comments. Auto Memory privacy improvement.

7. **[#26522](https://github.com/google-gemini/gemini-cli/issues/26522) — Stop Auto Memory from retrying low-signal sessions indefinitely (p2)**  
   Background extraction agent can keep re-processing low-value sessions. 5 comments.

8. **[#25166](https://github.com/google-gemini/gemini-cli/issues/25166) — Shell command execution gets stuck with "Waiting input" after command completes (p1)**  
   Common frustration: simple shell commands hang after finishing. 4 comments, 3 👍.

9. **[#21983](https://github.com/google-gemini/gemini-cli/issues/21983) — Browser subagent fails in Wayland (p1)**  
   Browser agent fails silently on Wayland systems. 4 comments, 1 👍. Cross-platform stability issue.

10. **[#22672](https://github.com/google-gemini/gemini-cli/issues/22672) — Agent should stop/discourage destructive behavior (p2)**  
    Model occasionally uses `git reset --force` or unsafe DB commands when safer alternatives exist. 3 comments, 1 👍.

## Key PR Progress
*10 important PRs updated in the last 24h.*

1. **[#28000](https://github.com/google-gemini/gemini-cli/pull/28000) — fix(core-tools): resolve Jupyter Notebook and JSON corruption in write_file**  
   Critical fix: `write_file` was silently corrupting `.ipynb` and JSON files. Fresh PR from author `amelidev`.

2. **[#28009](https://github.com/google-gemini/gemini-cli/pull/28009) — feat: add eval:inventory CLI command**  
   New `npm run eval:inventory` to list eval cases grouped by policy. Size/l. Useful for test management.

3. **[#27996](https://github.com/google-gemini/gemini-cli/pull/27996) — fix(core): decode response body using charset from Content-Type header**  
   Fixes garbled text for non-UTF-8 pages (e.g., Chinese, Japanese sites). Important for `web-fetch` tool.

4. **[#27987](https://github.com/google-gemini/gemini-cli/pull/27987) — fix(cli): throw FatalConfigError instead of process.exit in parseArguments**  
   Improves error handling by using proper exceptions instead of `process.exit(1)`. Enables clean E2E tests.

5. **[#27994](https://github.com/google-gemini/gemini-cli/pull/27994) — fix(core): insert skill/agent content literally in system prompt substitutions**  
   Fixes a bug where special characters in skill content could break prompt templates (using `String.replace` with string form instead of regex).

6. **[#27995](https://github.com/google-gemini/gemini-cli/pull/27995) — fix(core): skip user agents dir when it resolves to the project agents dir**  
   Eliminates false “Duplicate agent name” warnings when running from home directory.

7. **[#27648](https://github.com/google-gemini/gemini-cli/pull/27648) — feat(core): support list format in trustedFolders.json**  
   Adds JSON array format support for simpler manual editing of trusted directories.

8. **[#27788](https://github.com/google-gemini/gemini-cli/pull/27788) — test(core): add subfolder ignore test for getFolderStructure**  
   Adds test to verify `.gitignore` rules applied correctly in subdirectories.

9. **[#27948](https://github.com/google-gemini/gemini-cli/pull/27948) — chore(deps): pin dependencies and enforce 14-day update cooldown**  
   Pins all direct dependencies exactly and enforces a cooldown to reduce churn. Size/xl.

10. **[#27780](https://github.com/google-gemini/gemini-cli/pull/27780) — security: gate chained E2E on same-repository checkout for workflow_run**  
    Security improvement to prevent fork PRs from accessing `GEMINI_API_KEY` during E2E tests.

## Feature Request Trends
The community is pushing for:
- **AST-aware code understanding** – smarter file reads and codebase mapping to reduce token waste and improve tool accuracy (see #22745, #22746).
- **Better agent orchestration** – sub-agents and skills should be used automatically, not just when explicitly instructed (#21968). Also, agents should self-limit tool usage when >128 tools are available (#24246).
- **Auto Memory improvements** – deterministic redaction of secrets (#26525), avoiding infinite retry loops (#26522), surfacing invalid patches (#26523).
- **Robust evaluations** – component-level tests (#24353) and stable internal project evaluations (#23166).
- **Remote agent support** – advanced auth and background operations (#20303).
- **Self-awareness** – Gemini CLI should know its own flags, hotkeys, and how to run itself (#21432).

## Developer Pain Points
Recurring frustrations from open issues:
- **Agent hangs and false success** – The generalist agent hangs indefinitely (#21409); sub-agents report “GOAL” success after hitting turn limits (#22323). Often worked around by disabling sub-agents.
- **Shell execution stuck** – Simple commands complete but Gemini shows “Waiting for input” permanently (#25166).
- **Browser agent fragility** – Fails silently on Wayland (#21983); ignores `settings.json` overrides (#22267); no automatic lock recovery (#22232).
- **Memory system bugs** – Low-signal sessions retried forever (#26522); invalid patches silently skipped (#26523); secrets not redacted before model context (#26525).
- **Configuration quirks** – Symlinked agent files not recognized (#20079); browser agent ignores `maxTurns` override (#22267); `settings.json` merge issues.
- **Destructive behavior** – Model uses `--force` or unsafe commands when safer options exist (#22672); creates temporary scripts in random directories (#23571).
- **Terminal issues** – High performance flicker on resize (#21924); corruption after exiting external editor (#24935).

---
*Generated from GitHub data for google-gemini/gemini-cli, last updated 2026-06-18.*

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-06-18

## Today's Highlights
The past 24 hours saw a surge in **MCP- and agent-related issues**, including a closed fix for subagents losing access to MCP tools and a new bug report for Drive MCP OAuth credentials not attaching. Enterprise users continue to push for custom model support, while a single pull request aims to improve compatibility with strict OpenAI-compatible backends. No new releases were published.

## Releases
*No new releases in the last 24 hours.*

## Hot Issues (Top 10 by Impact)

1. **[#254 – copilot-cli keeps asking to login again](https://github.com/github/copilot-cli/issues/254)**  
   *9 comments, 4 👍*  
   Persistent re-authentication cycle for GitHub Business accounts. The community is frustrated; workaround is to re-login each session. The `more-info-needed` label suggests the team is still investigating.

2. **[#1974 – After upgrading to 1.0.3, generated Markdown links are not clickable](https://github.com/github/copilot-cli/issues/1974)**  
   *5 comments, 1 👍*  
   Regression affecting documentation output. Users working with rendered Markdown find this disruptive.

3. **[#3560 – Execution failed: CAPIError: 400 “Duplicate item found” on tool calls](https://github.com/github/copilot-cli/issues/3560)**  
   *5 comments, 1 👍*  
   Sudden websocket error after tool/function calls. Plain chat still works, indicating a server-side or state management issue.

4. **[#3838 – Drive MCP OAuth not attached: tools fail after successful reauth](https://github.com/github/copilot-cli/issues/3838)**  
   *2 comments, 0 👍*  
   OAuth flow completes, cache files created, but requests are sent without auth credentials. A clear MCP integration gap.

5. **[#3812 – Subagents can no longer access MCP tools](https://github.com/github/copilot-cli/issues/3812)**  
   *2 comments, 0 👍 – **CLOSED***  
   Downgrading didn’t help; deferred loading of MCP tools was identified as the root cause. Closed, so a fix is likely in the pipeline.

6. **[#3730 – Support Enterprise-Managed Custom Models in CLI](https://github.com/github/copilot-cli/issues/3730)**  
   *2 comments, 4 👍 – **CLOSED***  
   Feature request for centrally configured custom models to appear in CLI. The closure suggests it may have been accepted; expect a future release.

7. **[#3839 – Ollama Cloud does not support custom_tool_call payload](https://github.com/github/copilot-cli/issues/3839)**  
   *1 comment, 7 👍*  
   High community reaction. BYOK users with Ollama get a `400 BadRequest` due to unsupported `custom_tool_call` format. Fleet mode is blocked.

8. **[#3754 – `copilot --resume "Name With Spaces"` fails silently with exit 1](https://github.com/github/copilot-cli/issues/3754)**  
   *2 comments, 1 👍*  
   Session names with spaces break resume, contradicting documented behaviour. A simple but painful UX regression.

9. **[#3344 – Messages submitted while background subagents are running get stranded in Queued UI](https://github.com/github/copilot-cli/issues/3344)**  
   *1 comment, 0 👍*  
   Messages queued during agent idle state never drain automatically. The workaround (submit another message) doesn’t help.

10. **[#3850 – SDK/server mode: `session.create` drops `mcpServers`](https://github.com/github/copilot-cli/issues/3850)**  
    *1 comment, 0 👍 – **CLOSED***  
    Programmatically provided MCP servers via the SDK are never started. Closed quickly, indicating a rapid fix.

## Key PR Progress

Only one pull request was opened in the last 24 hours:

- **[#3847 – Plan review: add compatibility fallback design + test vectors](https://github.com/github/copilot-cli/pull/3847)**  
  Addresses [Issue #3846](https://github.com/github/copilot-cli/issues/3846) (closed) – plan review menus failing on strict OpenAI-compatible backends. Proposes a JSON-first parsing strategy with YAML and list heuristics as fallbacks. This is a design/document PR; no code changes yet.

## Feature Request Trends

From recent issues, the community is requesting:

- **Better BYOK/hosted model compatibility** – Enterprise-managed custom models and Ollama Cloud support (#3730, #3839).
- **MCP evolution** – Skill files declaring additional MCP servers (#3292), stable plugin installation with lock files (#3136), and MCP schema alignment with VSCode (#3835).
- **Session management improvements** – Ability to unarchive sessions (#3518), persistent `--resume` with spaces (#3754), and showing the spawning directory on resume (#3837).
- **Customizability** – User-defined aliases and `/commands` (#3844), persistent `/instructions` opt-out per repo (#3840).
- **Effort switches** – A dedicated `/effort` command to change reasoning effort quickly (#3074).

## Developer Pain Points

Recurring frustrations highlighted in the data:

- **Authentication instability** – The login loop (#254) remains unresolved, affecting enterprise users disproportionately.
- **Regressions after upgrades** – Broken Markdown links (#1974) and effort mismatch between CLI and VS Code (#3851) erode trust in release quality.
- **MCP integration friction** – OAuth not attached (#3838), schema incompatibility (#3835), and subagent tool access (#3812) create a fragmented experience.
- **Cross-session and backgrounding quirks** – Messages stranded in queue (#3344), `Ctrl+X B` only backgrounds shell commands, not subagents (#3845), and session name parsing issues (#3754) hinder advanced workflows.
- **Platform-specific gaps** – Windows sandbox fails with `E_NOTIMPL` (#3849) and `/voice` always transcribes in English (#3848) exclude non-English and Windows users.
- **Content exclusion policy confusion** – CLI incorrectly enforcing a policy meant for GitHub code review (#3841) blocks local file access.

---

*Generated from GitHub Copilot CLI repository activity up to 2026-06-18 23:59 UTC.*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-06-18

A quiet day on the Kimi Code CLI repository with no new releases or pull requests. The community’s focus sharpened on two key issues: one addressing onboarding pain points for MCP servers and plugins, and another requesting runtime execution mode switching. Below is a breakdown of the activity.

---

## Today's Highlights

- **No new releases** in the last 24 hours; the current version remains stable.  
- **Two new issues** surfaced, both reflecting growing demands around configuration ergonomics and execution flexibility.  
- The community highlighted that while Kimi Code works well once set up, the initial wiring of MCP servers, plugins, and sub-skills remains a friction point.

---

## Releases

No new releases in the last 24 hours. The latest stable version from previous days continues to be the active build.

---

## Hot Issues (2 items)

Since only two issues were updated in the last 24 hours, we cover both in detail.

### [#2460 – [CLOSED] Feedback: onboarding and configuring MCP servers, plugins, and sub-skills is harder than it needs to be](https://github.com/MoonshotAI/kimi-cli/issues/2460)
- **Author:** PowerBeef  
- **Created/Updated:** 2026-06-18  
- **Comments:** 0 | 👍: 0  
- **Why it matters:** This issue captures a recurring friction point for new users: the multi-step configuration of core components (MCP servers, plugins, sub-skills). The author praises the core functionality but calls for streamlined onboarding – a clear signal that the developer experience (DX) for setup needs improvement. Although closed, the feedback may inform upcoming CLI improvements.

### [#2459 – [Feature Request] 支持会话运行中切换执行模式（Agent ↔ 集群） || Supports switching execution mode during session running (Agent ↔ Cluster)](https://github.com/MoonshotAI/kimi-cli/issues/2459)
- **Author:** PresentXoX  
- **Created:** 2026-06-17 | **Updated:** 2026-06-17  
- **Comments:** 0 | 👍: 0  
- **Why it matters:** This bilingual feature request asks for the ability to switch between Agent and Cluster execution modes **mid-session**. Currently users must restart or reconfigure; this would enable greater flexibility for long-running or iterative tasks. The request signals demand for runtime dynamism in Kimi Code’s workflow.

---

## Key PR Progress

No pull requests were updated in the last 24 hours. All existing PRs remain under review or pending updates from authors.

---

## Feature Request Trends

From the limited data, two clear feature directions emerge:

1. **Simplified configuration wizard** – Users want a guided or auto-setup process for MCP servers, plugins, and sub-skills instead of manual editing of config files.
2. **Runtime execution mode switching** – The ability to hot-swap between Agent (single-task assistant) and Cluster (multi-agent/parallel) modes during an active session without restarting.

These reflect a broader desire for **adaptability and lower barriers to entry** in the Kimi Code CLI.

---

## Developer Pain Points

**Onboarding complexity** is the dominant pain point: setting up the environment for MCP servers, plugins, and sub-skills requires multiple disconnected steps. Feedback indicates that once everything is wired, the tool is excellent, but the initial configuration is frustrating and error-prone. No other major frustrations were noted in the last 24 hours.

---

*All data sourced from [github.com/MoonshotAI/kimi-cli](https://github.com/MoonshotAI/kimi-cli). Digest generated 2026-06-18.*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest – 2026-06-18

## Today’s Highlights
OpenCode v1.17.8 shipped with session timeline performance improvements and two important bugfixes for MCP tool schema validation and Cloudflare AI Gateway. The community is heavily discussing agent sandboxing (#2242, 73 comments) and a new hot-reload feature request (#8751, 77 👍). Two PRs landed to fix long-standing cursor blink (#32295) and mobile newline input (#32830), while config directory resolution (#32824) and MCP structured argument normalization (#32812) are moving forward.

## Releases
**v1.17.8** – Session timelines now load faster without flicker or scroll jumps. Bugfixes for OpenAI-compatible providers (MCP tool schemas no longer fail validation, thanks @jquense) and Cloudflare AI Gateway (configured API key correctly, thanks @keefetang).

## Hot Issues (Top 10)
1. **[#2242 – Sandbox the agent](https://github.com/anomalco/opencode/issues/2242)**  
   73 comments, 55 👍. Requests terminal command restrictions to prevent file access outside the current directory. No equivalent of macOS seatbelt exists yet – high demand.

2. **[#1905 – Terminal cursor blinks despite user setting](https://github.com/anomalco/opencode/issues/1905)**  
   18 comments, 35 👍. Since v0.4.27, cursor ignores Gnome Terminal’s non-blinking preference. A long‑standing annoyance finally addressed by PR #32295.

3. **[#8751 – Hot-reload agents, skills, and commands](https://github.com/anomalco/opencode/issues/8751)**  
   18 comments, **77 👍**. Top‑voted feature request: allow config changes to take effect without restarting OpenCode. Strong consensus.

4. **[#16610 – Hang at startup when inotify user instances exhausted](https://github.com/anomalco/opencode/issues/16610)**  
   10 comments, 7 👍. If `.git` is present and `fs.inotify.max_user_instances` is low, OpenCode hangs. Needs graceful fallback or warning.

5. **[#32749 – Explore agent wastes tokens](https://github.com/anomalco/opencode/issues/32749)**  
   4 comments, new today. The agent spawns a full explore subagent even for trivial tasks instead of using grep. Main agent then re‑reads all touched files.

6. **[#28472 – MCP tool parameters of type “object” serialized as strings](https://github.com/anomalco/opencode/issues/28472)**  
   4 comments. Top‑level `body` is passed as a JSON string instead of a native object, causing MCP validation errors. PR #32812 has been submitted to fix.

7. **[#32832 – MCP tool can no longer return image attachments](https://github.com/anomalco/opencode/issues/32832)**  
   New regression. MCP results with images broken in v1.17.5+. Works in v1.17.4 – high priority bug.

8. **[#32809 – MCP session recovery only triggers on HTTP 404, not 400](https://github.com/anomalco/opencode/issues/32809)**  
   Session recovery was added for 404, but some MCP servers return 400 for expired sessions. Recovery never fires, forcing restarts.

9. **[#32829 – DeepSeek provider: $ref/$defs in MCP tool schemas causes AttributeError](https://github.com/anomalco/opencode/issues/32829)**  
   Breaks Asana and Notion MCP servers. A schema resolution issue specific to DeepSeek.

10. **[#32825 – OPENCODE_CONFIG_DIR replaces global config in v2 services](https://github.com/anomalco/opencode/issues/32825)**  
    Config directory env var is resolved differently between old and v2 config loaders, causing missing global config. PR #32824 aims to fix.

## Key PR Progress (Top 10)
1. **[#23108 – Add cache_point_ttl option for Bedrock provider](https://github.com/anomalco/opencode/pull/23108)**  
   Adds `cache_point_ttl` to Bedrock config, injecting a cachePoint block after the system prompt to reduce latency.

2. **[#32833 – Complete Chinese translations + GLM 5.2 support](https://github.com/anomalco/opencode/pull/32833)**  
   Provides zh‑CN and zh‑TW translations and adds GLM 5.2 model support – a community contribution addressing multiple locale gaps.

3. **[#32675 – Managed background shell execution](https://github.com/anomalco/opencode/pull/32675)**  
   Introduces managed background mode for bash and shell tools across core and TUI, improving isolation and lifecycle.

4. **[#32824 – Support additive config directories](https://github.com/anomalco/opencode/pull/32824)**  
   Fixes config resolution regression from v1/v2 split. Old loader treated `OPENCODE_CONFIG_DIR` as extra; v2 replaced global. Now additive again.

5. **[#32830 – Fix mobile Enter key: insert newline instead of submit](https://github.com/anomalco/opencode/pull/32830)**  
   Fixes #20965 (and older duplicates). Mobile prompt input now respects Enter for newline, matching desktop behavior.

6. **[#32000 – Cancel active run before revert](https://github.com/anomalco/opencode/pull/32000)**  
   Prevents race conditions by canceling any running session operation before applying revert/undo state.

7. **[#32009 – Wait for revert before restoring prompt](https://github.com/anomalco/opencode/pull/32009)**  
   Revert actions no longer prematurely restore the reverted message text into the input field.

8. **[#31322 – Poll active sessions for external-writer changes](https://github.com/anomalco/opencode/pull/31322)**  
   Syncs session updates from another OpenCode client, remote agent, or `opencode compose` – fixes #31073.

9. **[#32295 – Add cursor style configuration](https://github.com/anomalco/opencode/pull/32295)**  
   Closes #1905 and #11738. Users can now configure cursor blink/block/underline in the TUI, respecting terminal preferences.

10. **[#32812 – Normalize MCP structured arguments](https://github.com/anomalco/opencode/pull/32812)**  
    Fixes #28472 by passing top‑level object parameters (e.g., `body`) as native objects instead of JSON strings.

## Feature Request Trends
- **Sandboxing & Security** (#2242, #32801) – High demand for restricting agent file/command access to current directory.
- **Hot‑Reload Configs** (#8751) – Users want to change agents, skills, and commands without restart.
- **Per‑Agent MCP Tool Filtering** (#28662) – Control tool allowances per agent to stay within model tool limits and improve trust boundaries.
- **Local Project Detection** (#15192) – Treat directories with `.git` or `opencode.json` as distinct projects, not global.
- **Broader Provider Support** (#32818, #32833) – Requests for GLM 5.2, MinMax M3, Kimi K 2.7‑code, Persian language, and German documentation.
- **Performance Optimizations** (#32819) – Reduce O(N) plugin dispatch overhead by pre‑indexing hooks.

## Developer Pain Points
- **MCP Tool Handling** – Repeated bugs with schema validation ($ref/$defs), object serialization, session recovery (404 vs 400), and image attachments. MCP integration remains fragile.
- **Terminal / TUI Quirks** – Cursor blink ignoring system settings, Shift+Enter submitting instead of newline, mobile Enter key.
- **Startup & Resource Issues** – Hang on inotify limit exhaustion, 100% CPU crash in LXC containers (#31668).
- **Config Confusion** – `OPENCODE_CONFIG_DIR` behavior difference between old and v2 config loaders, lack of documentation for SQLite schema (#32828).
- **Agent Efficiency** – Explore agent over‑subagents for simple tasks (#32749), wrong model selection or early stop (#30061).
- **Multi‑Workspace / Branch Handling** – OpenCode opening wrong directory when multiple branches checked out (#32801).

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-06-18

## Today’s Highlights
Today saw a flurry of activity around provider support and terminal compatibility. Key fixes landed for Warp terminal detection, streaming markdown rendering, and several provider-specific bugs (GLM-5.2, Azure Foundry, Copilot). The community is also pushing hard on extension API ergonomics and RPC capabilities, with multiple feature requests and PRs converging on better observability and config exposure.

## Releases
No new releases in the last 24 hours.

## Hot Issues
1. **#5825** – *Streaming markdown forces scroll to bottom*  
   A long-running bug (18 comments) where Pi auto-scrolls the user to the bottom while the agent streams, preventing reading. Only occurs with `clear on shrink`. Community frustration is evident, but a fix is already in progress via PR #5846.

2. **#5653** – *Move off Shrinkwrap*  
   Duplicate dependency copies caused by hoisting/nesting issues. This is a structural packaging problem affecting anyone using both `pi-ai` and `pi-coding-agent` directly. High impact for extension authors.

3. **#5654** – *Add `excludeFromContext` to custom messages*  
   A feature request to allow custom messages (e.g., `/status`) to be displayed but not sent to the LLM. Strong positive reaction (+1 👍) and clear use-case. Touches core message handling.

4. **#5407** – *Double backspace and double enter on Kitty terminal*  
   New user report about Pi registering keypresses twice on Kitty, but not on COSMIC Terminal. Already closed after investigation, but highlights an ongoing platform compatibility concern.

5. **#5428** – *Refining a plan leads to error in plan mode*  
   Plan mode refinement triggers “Agent is already processing” error. Related to a known race condition (#5062). Closed after fix, but signals complexity in concurrent message handling.

6. **#5868** – *pi-RPC bug: unknown command response missing id field*  
   Critical: unknown RPC commands cause client to hang for 30s because error response lacks the `id` field. Instantly fixed and closed. High severity for anyone building on RPC.

7. **#5827** – *Warp terminal not detected for Kitty image protocol*  
   Warp Terminal users couldn’t see inline images. Community patch landed quickly in PR #5841. Good example of responsive maintenance.

8. **#5865** – *TypeError: this.message.content.filter is not a function*  
   Crashing bug when a custom message has malformed content (e.g., `null`). Uncaught exception kills the whole process. Closed same day with fix.

9. **#5810** – *RPC: expose session entries and tree*  
   Feature request to add read-only RPC commands `get_entries` and `get_tree`. Needed for external tooling driving Pi from editors or scripts. Low controversy, likely to be implemented.

10. **#5770** – *Added support for GLM-5.2 effort level configuration (High & Max)*  
   Community-driven addition for Zhipu GLM-5.2 effort levels. Closed with PR #5873. Shows eagerness for cutting-edge model support.

## Key PR Progress
1. **#5874** – *feat(coding-agent): add automatic theme mode*  
   Open PR by mitsuhiko for light/dark theme auto-detection. Requires terminal events. Promises better UX for “euromaxxing in the summer”.

2. **#5846** – *fix(tui): stabilize streaming code fence rendering*  
   Directly addresses issue #5825. Stabilizes scroll behavior during markdown streaming. High priority merge expected soon.

3. **#5873** – *Feat/fireworks glm 5p2*  
   Merged PR adding GLM-5.2 support to the Fireworks provider. Closes #5872. Community quickly rallied to add model support across providers.

4. **#5869** – *Export config dirname*  
   Merged. Exposes `CONFIG_DIR_NAME` to extensions, fixing hardcoded paths. Responds to issue #5867.

5. **#5841** – *feat(tui): detect Warp terminal and enable Kitty image protocol*  
   Merged. Detects Warp Terminal via environment variables, enabling inline images. Closes #5827.

6. **#5866** – *feat(ai): add OpenRouter Fusion alias*  
   Open PR adding a synthetic `openrouter/fusion` alias for OpenRouter’s router model. Improves discoverability.

7. **#5701** – *fix(ai/model): adjust minimax-m3 context size*  
   Merged. Corrects the context size from 1M to 524288 based on OpenRouter error feedback. Simple but essential fix.

8. **#5859** – *fix(ai): send responses prompts as instructions*  
   Open PR fixing OpenAI Responses API: system prompts must go in `instructions`, not in `input`. Prevents silent errors for OpenAI, Azure, and Codex providers.

9. **#5849** – *feat(ai): add Azure AI Foundry provider for Anthropic Claude*  
   Merged. New `azure-foundry` provider with Entra ID auth. Major addition for enterprise users on Azure.

10. **#5832** – *fix(ai): surface provider HTTP error body instead of opaque SDK message*  
   Open PR that exposes the actual HTTP error body from proxy/gateway failures, rather than `Unknown: UnknownError`. Great for debugging.

## Feature Request Trends
- **Multi-provider expansion**: GLM-5.2, Azure Foundry, OpenRouter aliases, Mistral prompt caching, and Copilot 1M context are all hot requests. The community wants Pi to support every major model endpoint.
- **Extension API empowerment**: Exposing `CONFIG_DIR_NAME`, active tools as executable objects, read-only harness state, and persistent display messages not sent to LLM are recurring themes.
- **RPC enhancements**: Standardised RPC commands for session data, error handling, and request/response matching are needed for external tool integrations.
- **Terminal compatibility**: Beyond Kitty and Warp, users want Pi to detect and adapt to various terminal features (image protocols, color schemes, key handling).
- **Command and UX polish**: `/exit`, template newline preservation, and better CLI flag validation show demand for more polished user-facing behaviour.

## Developer Pain Points
- **Dependency duplication** (#5653) – Using multiple Pi packages installs duplicate copies of `pi-ai`, leading to separate module-level state. The community is frustrated by shrinkwrap and wants a solution.
- **Terminal detection gaps** (#5827, #5407, #5428) – Terminal-specific bugs (Kitty, Warp, COSMIC) require per-environment workarounds. Users expect Pi to “just work” on any modern terminal.
- **Schema and provider validation errors** (#5870, #5876, #5864) – Loosely typed tool schemas break on Kimi, Cloudflare’s `detectCompat` forces wrong field names, and Copilot intermittently rejects system prompt placement. These are low-level but blocking.
- **Silent failures** (#5856, #5868) – CLI flags swallow missing values, and RPC errors drop the `id` – both cause mysterious hangs or silent misbehaviour. Developers want clear, actionable error messages.
- **Context control** (#5654, #5861) – The inability to tag custom messages as “display only” and the lack of persistent visible content not sent to the LLM are recurring friction points for extension builders.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest – 2026-06-18

## Today's Highlights
Qwen Code released **v0.18.3** with a fix for cancelled CLI prompts and a nightly build adding core improvements to sed edit tracking. The community reported a cluster of input‑validation bugs and one security‑relevant remote‑host check issue, many of which have correlating PRs already in review. The multi‑agent crash report (#5180) continues to attract attention as the team works on background automation infrastructure.

## Releases
Three versions were published in the last 24 hours:

- **[v0.18.3-nightly.20260618](https://github.com/QwenLM/qwen-code/releases/tag/v0.18.3-nightly.20260618.bc3e0b405)** – Chore release; fix(core): track supported sed edits in file history.
- **[v0.18.3](https://github.com/QwenLM/qwen-code/releases/tag/v0.18.3)** – Official patch; fix(cli): stop CLI processing after `ask_user_question` is cancelled.
- **[v0.18.3-preview.0](https://github.com/QwenLM/qwen-code/releases/tag/v0.18.3-preview.0)** – Pre‑release with the same CLI fix.

All releases include the `.github/release.yml` configuration update. Developers should upgrade to **v0.18.3** for the CLI stability fix.

## Hot Issues (10 selected)

1. **[#5180 – Multi‑agent session crashes mid‑execution](https://github.com/QwenLM/qwen-code/issues/5180)**  
   *Priority P2, type/bug, core, performance, memory, multi‑agent, welcome‑pr*  
   A 12‑hour session using a manager/subagent pattern crashes unpredictably. Community discussion (4 comments) suggests memory/context overflow during task delegation. Tagged with `roadmap/multi-agent`, this is a top‑priority stability concern.

2. **[#5329 – readStdin counts JS characters, not UTF‑8 bytes](https://github.com/QwenLM/qwen-code/issues/5329)**  
   *Priority P2, CLI, non‑interactive*  
   The 8 MB truncation limit uses `string.length`, so multi‑byte characters (e.g., `é` repeated 4 million times) can exceed the byte limit. A critical input‑handling bug for scripted pipelines.

3. **[#5326 – GitHub remote check accepts lookalike hosts](https://github.com/QwenLM/qwen-code/issues/5326)**  
   *Priority P2, CLI, security*  
   `isGitHubRepository()` uses a bare `/github\.com/` regex, allowing `https://github.com.evil/...` to pass. A security issue that could enable repo‑spoofing attacks. A fix PR (#5327) is already open.

4. **[#5324 – Sandbox image parser mishandles registry ports](https://github.com/QwenLM/qwen-code/issues/5324)**  
   *Priority P3, CLI, sandbox, welcome‑pr*  
   `image.split(":")` incorrectly treats `localhost:5000/team/qwen-code:dev` as tag `5000/team/qwen-code:dev`, losing the basename. Blocks private‑registry usage. PR #5325 addresses it.

5. **[#5322 – MCP prompt parser drops empty named arguments](https://github.com/QwenLM/qwen-code/issues/5322)**  
   *Priority P3, CLI, MCP*  
   `--trail=""` is treated as missing, causing optional args to be lost and required args to be reported missing. Affects MCP tool invocation reliability.

6. **[#5320 – TOML command migration emits invalid Markdown descriptions](https://github.com/QwenLM/qwen-code/issues/5320)**  
   *Priority P2, core, commands*  
   `description = "false"` becomes YAML boolean `false`, causing the migrated command to be skipped. Breaks command migration from TOML to Markdown format.

7. **[#5316 – Shell semantics miss attached stdout fd redirects](https://github.com/QwenLM/qwen-code/issues/5316)**  
   *Priority P2, core, shell*  
   Commands like `echo hi 1>.qwen/settings.json` are not recognised as write operations, bypassing shell permission checks. Affects security analysis of shell commands.

8. **[#5313 – ACP settings accept malformed timeout strings](https://github.com/QwenLM/qwen-code/issues/5313)**  
   *Priority P3, CLI, settings*  
   `normalizeOptionalNumber` uses `parseInt`, so `"10ms"` is silently converted to `10`. A validation hole that could lead to unexpected timeout behaviour.

9. **[#5310 – OpenAI schema converter truncates fractional length constraints](https://github.com/QwenLM/qwen-code/issues/5310)**  
   *Priority P2, core, welcome‑pr*  
   `parseInt("1.5")` turns `minLength: "1.5"` into `1`, silently changing schema semantics. Affects Gemini‑to‑OpenAI tool parameter conversion.

10. **[#5308 – Invalid cron task entries can be permanently deleted](https://github.com/QwenLM/qwen-code/issues/5308)**  
    *Priority P2, core, session‑management, background‑automation*  
    `readCronTasks()` filters out malformed entries, then `updateCronTasks()` writes back the filtered list, losing data. A potential data‑loss bug for scheduled tasks.

## Key PR Progress (10 selected)

1. **[#5330 – SSE stream idle watchdog](https://github.com/QwenLM/qwen-code/pull/5330)**  
   Adds a configurable timeout that aborts hung SSE connections, preventing frozen UI spinners. Directly addresses #4177 (not listed today but a known pain point).

2. **[#5319 – Rename TodoWrite display name to TodoList](https://github.com/QwenLM/qwen-code/pull/5319)**  
   User‑facing renaming across terminal UI, web‑shell, and webui while keeping the internal tool name `todo_write`. Improves consistency and model alignment.

3. **[#5328 – Keep qwen3.6‑flash and kimi‑k2.6 presets text‑only](https://github.com/QwenLM/qwen-code/pull/5328)**  
   Removes stray `modalities: { image: true, video: true }` from specs that are not actually multimodal. Prevents token‑plan mismatches.

4. **[#5182 – Second‑resolution session wakeup engine](https://github.com/QwenLM/qwen-code/pull/5182)**  
   Step 1 of aligning `/loop` with Claude Code’s `ScheduleWakeup`. Adds a wakeup primitive inside `CronScheduler` for self‑paced loops. Part of the background‑automation roadmap.

5. **[#5327 – Validate GitHub remote hosts](https://github.com/QwenLM/qwen-code/pull/5327)**  
   Fixes #5326 by parsing remote URLs properly. Includes regex tests for `git@github.com:owner/repo.git` and rejects lookalike domains.

6. **[#5318 – Add --no-ask‑password to systemd‑inhibit](https://github.com/QwenLM/qwen-code/pull/5318)**  
   Fixes #5281 by passing `--no-ask-password` to prevent TTY interaction during sleep inhibition. Simple but important for headless environments.

7. **[#4909 – Support archive install sources for extensions](https://github.com/QwenLM/qwen-code/pull/4909)**  
   Enables installing extensions from local `.zip`/`.tar.gz` and remote archive URLs. Reuses existing pipeline – a long‑requested feature for offline/air‑gapped setups.

8. **[#5325 – Parse sandbox image registry ports](https://github.com/QwenLM/qwen-code/pull/5325)**  
   Fixes #5324 with a tested helper that correctly splits registry ports from image tags. Keeps container‑name prefixes based on basename and tag.

9. **[#5181 – Prevent OOM in auto‑memory extraction during /quit](https://github.com/QwenLM/qwen-code/pull/5181)**  
   Fixes #5147 by chunking the conversation history in `buildTranscriptMessages()`. The `heap limit` crash after `/quit` is a high‑severity memory issue now addressed.

10. **[#5202 – QQ Bot channel adapter](https://github.com/QwenLM/qwen-code/pull/5202)**  
    Adds a new `@qwen-code/channel-qqbot` adapter using the official WebSocket gateway API. Joins Telegram, WeChat, DingTalk, Feishu – expands Qwen Code’s reach in Chinese markets.

## Feature Request Trends
- **Multi‑agent orchestration stability** – Issue #5180 and roadmap items for background automation show strong demand for reliable agent delegation and session recovery.
- **Model configuration flexibility** – #5303 asks how to configure the same model ID with different endpoints/keys, reflecting a need for multi‑provider setups.
- **Background/scheduled task automation** – #5184, #5156, and #5308 point toward a unified `/loop` command with second‑resolution wakeups and robust cron task management.
- **Extension distribution** – PR #4909 (archive install) indicates interest in offline extension deployment and simplified sharing.

## Developer Pain Points
- **Input validation gaps** – Multiple issues (#5329, #5313, #5324, #5322, #5310, #5316) highlight parsing edge cases that silently accept malformed or oversize inputs, leading to security holes or data corruption.
- **Memory management** – OOM crashes (#5181), unbounded streaming buffers (#5314), and retained tool output (#4971) are recurring memory‑related frustrations.
- **Security surface** – The GitHub remote‑host check (#5326) and shell redirect semantics (#5316) show that even basic validation can be bypassed, raising trust concerns.
- **Schema conversion fragility** – #5320 and #5310 demonstrate that migrating between formats (TOML→Markdown, Gemini→OpenAI) introduces silent data loss, eroding developer confidence.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI / CodeWhale Community Digest — 2026-06-18

## Today's Highlights
The project officially rebranded from `deepseek-tui` to **CodeWhale** with v0.8.62, and the legacy npm package is fully deprecated. On the bug-fix front, maintainers merged critical fixes for the persistent "turn stalled" error and the thinking-stream rendering bottleneck — two of the most upvoted community pain points. A new Workrooms Phase 1 foundation PR signals the direction for v0.9.0.

## Releases
**v0.8.62** — This release completes the **CodeWhale** rebranding. The canonical project name, command, npm package, and release assets now all use "CodeWhale". Users still on the legacy `deepseek` / `deepseek-tui` names should follow `docs/REBRAND.md` for migration instructions. The old npm package receives no further updates.

## Hot Issues

1. **[#2487 — Frequent "Turn stalled" error in yolo mode](https://github.com/Hmbown/CodeWhale/issues/2487)**  
   *Open, 15 comments*  
   A top community blocker: `yolo` mode freezes during long operations with "Turn stalled — no completion signal received". The `continue` command fails to resume, and sessions appear lost. This is the most-commented open issue and directly impacts daily workflows.

2. **[#3275 — Agent over-extends scope with self-questioning loops](https://github.com/Hmbown/CodeWhale/issues/3275)**  
   *Open, 4 comments*  
   A regression from a previous fix: the agent enters a self-sustaining loop of proposing, answering, and executing without waiting for user confirmation, completely disregarding the requested scope. Contributors flagged this as a critical UX regression.

3. **[#3279 — Plan/Agent mode toggle inconsistency & tool permission chaos](https://github.com/Hmbown/CodeWhale/issues/3279)**  
   *Closed, 3 comments*  
   Swapping from Plan to Agent mode would leave `write_file` and `exec_shell` permanently denied by user, despite the UI showing Agent mode. Upon recovery, the AI would then auto-execute plans without approval. A high-severity mode UX bug.

4. **[#1620 — Thinking process painfully slow — "one character takes ages"](https://github.com/Hmbown/CodeWhale/issues/1620)**  
   *Closed, 5 comments*  
   Reasoning blocks stream characters one at a time on fast models, making the thinking display nearly unusable. Community described it as "unbearably slow". This was resolved via a debounce fix (see PR #3284).

5. **[#3292 — `pre_tool_snapshot` ignored `snapshots.enabled=false` config](https://github.com/Hmbown/CodeWhale/issues/3292)**  
   *Closed, 1 comment*  
   Even after disabling snapshots in config, the tool copied entire git repos to `~/.deepseek/snapshots/`, consuming gigabytes of disk. A configuration-respecting bug that surprised users expecting clean behavior.

6. **[#3282 — Config.toml comment lines erased on TUI edits](https://github.com/Hmbown/CodeWhale/issues/3282)**  
   *Closed, 1 comment*  
   Every modification via TUI (e.g., trust zone fields) silently strips all comments from `config.toml`. Users rely on comments for notes and toggling configs, so this was a persistent annoyance.

7. **[#1481 — Support OpenCode Go/Zen as provider](https://github.com/Hmbown/CodeWhale/issues/1481)**  
   *Open, 2 comments, 👍1*  
   Request to add OpenCode Go/Zen as a cheap DeepSeek-V4 provider. This has been open for over a month with minimal engagement, but represents a common demand for alternative, cost-effective backends.

8. **[#3281 — `sanitize_for_kimi_parameters` fix incomplete for root schemas](https://github.com/Hmbown/CodeWhale/issues/3281)**  
   *Closed, 2 comments*  
   A v0.8.61 fix only added `"type": "object"` under narrow conditions, missing `$ref`, `anyOf`, and `allOf` root schemas. Community tester identified the edge case, and a tighter fix was accepted.

9. **[#2739 — Task execution freezes, impossible to resume after cancel](https://github.com/Hmbown/CodeWhale/issues/2739)**  
   *Closed, 4 comments*  
   Long-running tasks freeze with infinite waits; pressing Esc shows timeout. Using `--continue` loses all in-progress session context. User reported this across multiple versions (0.8.51 → 0.8.61), calling it "unbearable".

10. **[#3264 — Restrict skill scanning to `~/.codewhale/skills/` only](https://github.com/Hmbown/CodeWhale/issues/3264)**  
    *Closed, 3 comments*  
    Users want an option to limit skill discovery to CodeWhale-specific directories to avoid scanning unrelated paths. A clear config-gating request for power users with complex environments.

## Key PR Progress

1. **[PR #3301 — Save ask permission rules from approvals](https://github.com/Hmbown/CodeWhale/pull/3301)**  
   *Open*  
   Introduces a TUI approval modal that persists `exec_shell` rules into `permissions.toml` as ask-only policies. Adds a `s` key binding to save rules. This directly addresses permission chaos reported in #3279.

2. **[PR #3300 — Preserve thinking/tool blocks when seeding thread from session](https://github.com/Hmbown/CodeWhale/pull/3300)**  
   *Open*  
   Replaces the text-only seed approach with block-type-aware preservation, keeping Thinking, ToolUse, and ToolResult blocks intact across sessions. Critical for session continuity after stalls.

3. **[PR #3283 — Fix Plan/Agent mode toggle — `approval_mode` restore + auto‑execution guard](https://github.com/Hmbown/CodeWhale/pull/3283)**  
   *Closed*  
   Two targeted fixes for #3279: restores `approval_mode` correctly on Plan→Agent switch, and adds a guard to prevent auto-execution after mode recovery.

4. **[PR #3290 — Add `scope_discipline` rules to prevent self-questioning loops](https://github.com/Hmbown/CodeWhale/pull/3290)**  
   *Closed*  
   Adds prompt-level discipline rules to `constitution.md` that constrain the agent from entering self-sustaining question-answer loops. Direct fix for #3275.

5. **[PR #3284 — Debounce thinking-stream re-renders](https://github.com/Hmbown/CodeWhale/pull/3284)**  
   *Closed*  
   Fixes #1620 by debouncing `bump_active_cell_revision()` calls during reasoning deltas. Dramatically improves thinking-block rendering speed on fast models.

6. **[PR #3285 — Persist session before stall/cancel recovery](https://github.com/Hmbown/CodeWhale/pull/3285)**  
   *Closed*  
   Fixes #2739 data-loss aspect: saves session state before clearing turn bookkeeping during stall/cancel, so `--continue` recovers the in-progress turn instead of reverting to the previous session.

7. **[PR #3291 — Preserve comments in config files](https://github.com/Hmbown/CodeWhale/pull/3291)**  
   *Closed*  
   All TUI/CLI config file writes now use `toml_edit` to merge serialized output with original comments. Resolves #3282 by keeping user annotations and commented-out keys intact.

8. **[PR #3293 — Respect `snapshots.enabled` for per-tool snapshots](https://github.com/Hmbown/CodeWhale/pull/3293)**  
   *Closed*  
   Fixes #3292 by adding the `snapshots.enabled` guard to the per-tool snapshot call site in `turn_loop.rs`, ensuring zero snaps when the config flag is false.

9. **[PR #3274 — Build static Linux x64 binaries with musl](https://github.com/Hmbown/CodeWhale/pull/3274)**  
   *Closed*  
   Switches Linux x64 release builds from dynamic glibc to static musl for better portability across Linux distros.

10. **[PR #3277 — Implement Workrooms Phase 1 — data model, endpoints, docs, and tool](https://github.com/Hmbown/CodeWhale/pull/3277)**  
    *Closed*  
    Foundation for the v0.9.0 Workroom abstraction: a chat-native, durable, addressable container for threaded agent conversations. Includes a 242-line RFC, data model, endpoints, and migration tool.

## Feature Request Trends

- **Agent mode guards and scope discipline** — Multiple issues request better control over when the agent should stop, wait for confirmation, or avoid self-questioning loops. Expect more permission-rule and policy-engine features.
- **Config persistency and comment preservation** — Users want TUI config edits to respect human-added comments and structure without silent erasure. The `toml_edit` merge approach is now merged but broader config hygiene is desired.
- **Skill scanning scoping** — The community wants fine-grained control over which directories CodeWhale scans for skills, with `scan_codewhale_only` being the first such option.
- **Alternative provider support** — Beyond DeepSeek-native backends, requests for OpenCode, Atlas Cloud, and other cheap or geo-distributed providers continue to grow.

## Developer Pain Points

- **Stalls and unresponsiveness** — The most acute pain: `yolo` mode stalls, long-task freezes, and unrecoverable sessions. The "Turn stalled" error and `--continue` data loss have dominated recent community discussion.
- **Mode switching instability** — Plan/Agent/YOLO mode transistions frequently break permission states, auto-execute without approval, or fail to restore saved settings. This undermines trust in the mode system.
- **Thinking-stream rendering performance** — Slow reasoning block display on fast models was a long-standing UX issue that only recently received a fix (debounce).
- **Configuration surprises** — Silent comment erasure, ignored `snapshots.enabled` flags, and missing documentation for app-server flags make configuration feel fragile.
- **Session continuity after failures** — The inability to recover stalled sessions without losing context was a top complaint, now partially addressed by PR #3285.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*