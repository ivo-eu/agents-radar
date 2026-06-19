# AI CLI Tools Community Digest 2026-06-19

> Generated: 2026-06-19 12:58 UTC | Tools covered: 9

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

# Cross-Tool Comparison Report – 2026-06-19

## 1. Ecosystem Overview

The AI CLI tooling landscape continues to mature rapidly, with eight major projects showing high activity across repositories. Today’s digests reveal a common struggle between feature expansion and stability—every tool reports regressions in performance, security, or platform compatibility following recent updates. Enterprise concerns (OAuth, sandboxing, proxy support) are now front and center, while community demand for sub-agent orchestration control and memory efficiency reaches a peak. The ecosystem is bifurcating along two axes: tools that prioritize safety/guardrails (Claude Code, Copilot CLI) versus those racing to support more providers and extensibility (Pi, Qwen Code). Kimi Code remains the smallest player, but its proxy fix signals effort to serve restricted environments.

---

## 2. Activity Comparison (Last 24 Hours)

| Tool | Hot Issues Shown | PRs Updated | New Releases | Notable Community Signal |
|------|----------------|-------------|--------------|--------------------------|
| Claude Code | 10 | 3 | ✅ v2.1.183 | 143 👍 on hanging bug (#26224) |
| OpenAI Codex | 10 | 10 | ✅ 3 alpha releases | 604 👍 on Linux desktop (#11023) |
| Gemini CLI | 10 | 10 | ❌ None | 8 👍 on agent hang (#21409) |
| GitHub Copilot CLI | 10 | 0 | ❌ None | High-severity WSL2 CPU spin (#3700) |
| Kimi Code | 2 | 1 | ❌ None | Proxy env fix in review (#2461) |
| OpenCode | 10 | 10 | ❌ None | 24 👍 on MCP support (#28567) |
| Pi (pi-mono) | 10 | 10 | ✅ v0.79.7 & v0.79.8 | Fuzzy edit data-loss bug (#5899) |
| Qwen Code | 10 | 10 | ❌ None | 5 comments on model config UX (#4814) |
| DeepSeek TUI | 10 | 10 | ❌ None | TUI freeze on Windows (#1812) |

*Note: Issues and PR counts are those highlighted in each digest, not absolute totals.*

---

## 3. Shared Feature Directions

### Sub‑agent Orchestration & Control
- **Claude Code**: Subagent infinite recursion (#68619) – critical token burn.
- **DeepSeek TUI**: First-class sub-agent toggle PR (#3327); freeze after spawning sub-agents (#3289).
- **Gemini CLI**: Generalist agent hangs (#21409); subagent recovery falsely reports success (#22323).
- **Copilot CLI**: “Subconscious” sidekick keeps spawning (#3859).

### Proxy & Network Environment Support
- **Kimi Code**: `FetchURL` ignores `HTTP_PROXY` (#2455), fix in PR #2461.
- **DeepSeek TUI**: Proxy env not passed to JS execution (#3273), fixed in PR #3331.
- **Gemini CLI**: OAuth token atomic writes (PR #27664) – security for proxied setups.
- **Copilot CLI**: MCP OAuth not attached after reauth (#3838).

### Windows & Cross‑Platform Compatibility
- **DeepSeek TUI**: TUI freeze on Windows 11 (#1812); glibc mismatch on Ubuntu (#3238).
- **Kimi Code**: VS Code extension fails on Windows+Git Bash (#2462).
- **Qwen Code**: `SANDBOX_MOUNTS` misparses drive paths (#5386); grep drops colons (#5370).
- **OpenCode**: Editor context syncing breaks in WSL (#29570).
- **Copilot CLI**: WSL2 CPU spin regression (#3700).
- **Claude Code**: Chat history loss on mapped drives (#14088).

### Memory & Context Efficiency
- **Claude Code**: `autoMemoryEnabled=false` still loads 11–16k token preamble (#63903).
- **Pi**: Compaction inefficiencies (#5845); coding agent forgets turns (#5882).
- **Qwen Code**: Token count accuracy concerns (#4951); microcompaction cache fixes (#5407).
- **Gemini CLI**: Auto memory retries low-signal sessions indefinitely (#26522).
- **Copilot CLI**: Memory sidekick runs even when disabled (#3859).
- **OpenCode**: Snapshot on home directory hangs (#32981).

### Plugin / Extension Ecosystems
- **Copilot CLI**: Project‑scoped plugins (#1665) – top feature request.
- **OpenCode**: Full MCP client capabilities (#28567); AXI integration in sidebar (#32994).
- **Qwen Code**: Archive‑based extension installs (#4909); web‑shell extension management (#5398).
- **Pi**: Extension API to expose active tools (#5781).
- **DeepSeek TUI**: Hook layer proposal (#1917) for all actions.

### Real‑time Metrics & Observability
- **Qwen Code**: Optional tokens‑per‑second display (merged #5401).
- **Pi**: OSC 9998/9999 for web UI status (#5900).
- **OpenCode**: Manual performance diagnostics suite (#32937).
- **Copilot CLI**: Stale model selectors require restart (#2408, discussed in Pi but similar pattern).

### Safety & Guardrails
- **Claude Code**: Destructive Git commands blocked (#183 release); email leakage in commit trailers (#66079).
- **Codex**: MITM CA key protection (PR #29013); network approval scoping (PR #28899).
- **Gemini CLI**: Agent should discourage destructive commands (#22672); MCP OAuth atomicity (PR #27664).
- **Copilot CLI**: `preToolUse` hooks bypassed under parallel calls (#2893) – critical security gap.
- **Pi**: Fuzzy edit silently rewrites entire file (#5899) – data‑loss risk.
- **DeepSeek TUI**: Over‑involvement loop bypasses user intent (#3275).

---

## 4. Differentiation Analysis

| Tool | Core Differentiator | Target User | Technical Approach |
|------|---------------------|-------------|-------------------|
| **Claude Code** | Mature safety model, deep Git integration | Enterprise developers with strict compliance | Node.js, heavy focus on destructive command blocking and privacy |
| **OpenAI Codex** | Desktop app + CLI, MCP/agent orchestration | Multi‑agent workflows, Pro subscribers | Rust, desktop Tauri/Electron, sandboxed execution |
| **Gemini CLI** | Google‑native, strong sub‑agent / evaluation system | Google Cloud users, research‑oriented | TypeScript, hierarchical agent structure, evaluated from day one |
| **Copilot CLI** | Plugin hooks, GitHub authentication | GitHub‑centric teams | Go, hooks pipeline, project‑scoped config |
| **Kimi Code** | Lightweight, proxy‑friendly | Developers behind firewalls (Asia, enterprises) | Python (aiohttp), minimal external deps |
| **OpenCode** | TUI + Electron, MCP/AXI ecosystem | Power users who want terminal + desktop | TypeScript, Tauri→Electron, plugin agent model |
| **Pi (pi-mono)** | Maximum provider support, extensible SDK | Developers who want to bring their own models | TypeScript, modular base entrypoints, fuzzy edit tools |
| **Qwen Code** | Alibaba ecosystem, rapid bug fixing, Windows focus | Chinese‑market developers, pragmatic users | TypeScript, web shell, extension archives, token counters |
| **DeepSeek TUI** | Rust‑based performance, workflow engine (WhaleFlow) | Performance‑sensitive, refactoring‑focused | Rust (crossterm), active monolith→modular transition |

**Key differentiators:**
- *Safety depth*: Claude Code > Copilot CLI > others.
- *Provider diversity*: Pi leads with many backends; Qwen Code and DeepSeek TUI are catching up.
- *Desktop integration*: Codex, OpenCode, and Pi (OSC bridge) offer richer UI.
- *Enterprise readiness*: Codex and Claude Code have OAuth/SSO issues; Copilot CLI and Pine lack enterprise auth stories.
- *Windows support*: Qwen Code is most proactive on Windows fixes; DeepSeek TUI and Copilot CLI have known regressions.

---

## 5. Community Momentum & Maturity

### High Momentum / Rapid Iteration
- **Qwen Code**: 10 PRs in 24h, many closed same day, strong contributor (tt-a1i). Bug–fix cycle is hours.
- **Pi**: Two releases today, 10 PRs, fast responses on data‑loss bug (#5899 → PR #5898). Active maintainer.
- **OpenCode**: 50 PRs updated in 24h, huge issue count, active community engagement (33 issues updated). Ecosystem growth (AXI integration).
- **DeepSeek TUI**: 10 PRs, structural refactoring, clear v0.9 roadmap. High engagement on sub‑agent features.

### Mature but Slower Release Cadence
- **Claude Code**: One release today, but only 3 PRs. Long‑standing hanging bug (#26224) unresolved, suggesting focus on safety rather than iteration speed.
- **Codex**: Three alpha releases but no changelog; many PRs but enterprise issues persist (phone auth, business tokens). Moderate community frustration.
- **Gemini CLI**: 10 PRs, but no release today. Top blocker (agent hang) remains open. Epics indicate deep planning.
- **Copilot CLI**: Zero PRs today, multiple high‑severity bugs (WSL2 CPU, hook bypass). Community seems less active or maintainers are stretched.

### Smallest / Nichest
- **Kimi Code**: Only 2 issues and 1 PR today. Community is tiny but maintainers are responsive (proxy fix in review). Likely early‑stage.

---

## 6. Trend Signals (Actionable Insights for Developers)

1. **Sub‑agent governance is the next frontier.** Every tool with multi‑agent capabilities is seeing runaway recursion, false success reporting, or wasted tokens. Expect built‑in budget regulators, turn limits, and user‑confirmable sub‑agent spawning soon.

2. **Proxy & network flexibility is no longer optional.** With enterprise adoption and regional firewalls, tools that ignore `HTTP_PROXY` (Kimi) or fail on corporate networks will lose users. This is a baseline requirement.

3. **Windows parity remains fragile.** Despite many tools being cross‑platform, Windows regressions surface daily—path parsing, sandbox execution, keyboard shortcuts, TUI deadlocks. Teams targeting developers on Windows must invest in CI testing on WSL2 and Git Bash.

4. **Token transparency is demanded.** Users no longer accept opaque context costs. Qwen Code’s token‑per‑second meter, Pi’s compaction issues, and Claude Code’s memory preamble complaints all point to a need for clear, real‑time cost visibility.

5. **Plugin/extension ecosystems are shifting from “nice to have” to critical.** Copilot CLI’s per‑project plugins, OpenCode’s MCP alignment, Qwen’s archive installs—extensibility is becoming a differentiator. Tools without a plugin story risk being locked out of enterprise workflows.

6. **Data integrity incidents erode trust quickly.** Pi’s fuzzy edit data loss (#5899) and Claude Code’s email leakage (#66079) are alarm bells. Any tool that silently rewrites files or leaks PII will face severe backlash.

7. **Desktop + TUI convergence continues.** Codex’s desktop app, OpenCode’s Electron shift, and Pi’s OSC bridge to web UI show that a pure CLI is insufficient. Users want visual context (sidebars, thread management, permission previews) without leaving the terminal.

8. **Chinese‑origin tools are growing fast.** Qwen Code and DeepSeek TUI are both iterating aggressively, fixing bugs in hours, and responding to community feedback. They are closing the gap with Western tools on quality while offering unique features (Qwen’s web shell, DeepSeek’s performance).

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
*Data: github.com/anthropics/skills | Snapshot: 2026-06-19*

---

## 1. Top Skills Ranking

The following Pull Requests represent the most-discussed new or improved skill submissions, ranked by community engagement (comments and cross-references). All are currently **open**.

### #514 – document-typography
- **What it does:** Prevents orphan word wrap, widow paragraphs, and numbering misalignment in AI-generated documents — common typographic issues that affect readability.
- **Discussion highlights:** Addresses a universal pain point for any user generating documents with Claude. The skill fills a gap between raw content generation and professional formatting.
- **Status:** Open  
  [PR #514](https://github.com/anthropics/skills/pull/514)

### #486 – ODT (OpenDocument) creation, template filling, and conversion
- **What it does:** Enables Claude to create, fill, read, and convert `.odt` and `.ods` files (LibreOffice/OpenOffice formats). Supports ISO-standard open-source document workflows.
- **Discussion highlights:** Community interest in interoperability beyond DOCX. Multiple related issues about ODF support exist.
- **Status:** Open  
  [PR #486](https://github.com/anthropics/skills/pull/486)

### #83 – skill-quality-analyzer & skill-security-analyzer (meta skills)
- **What it does:** Two meta-skills for evaluating other skills across dimensions (structure, documentation, security, test coverage). Aims to raise the bar for skill submissions.
- **Discussion highlights:** Reflects growing community concern about skill quality and trustworthiness. Aligns with Issue #492 (namespace security) and #202 (skill-creator best practices).
- **Status:** Open  
  [PR #83](https://github.com/anthropics/skills/pull/83)

### #181 – SAP-RPT-1-OSS predictor
- **What it does:** Integrates SAP's open-source tabular foundation model for predictive analytics on enterprise business data (e.g., supply chain, finance).
- **Discussion highlights:** First enterprise-focused predictive skill; signals demand for domain-specific AI in large organizations.
- **Status:** Open  
  [PR #181](https://github.com/anthropics/skills/pull/181)

### #723 – testing-patterns
- **What it does:** Comprehensive skill covering unit testing (AAA pattern), React Testing Library, integration/E2E tests, and testing philosophy (Testing Trophy model).
- **Discussion highlights:** Responds to repeated requests for developer tooling. Complements existing code generation skills.
- **Status:** Open  
  [PR #723](https://github.com/anthropics/skills/pull/723)

### #335 – masonry-generate-image-and-videos
- **What it does:** Wraps Masonry CLI for AI-powered image (Imagen 3.0) and video (Veo 3.1) generation, including job management.
- **Discussion highlights:** First dedicated media generation skill; intersects with multimodal AI demand.
- **Status:** Open  
  [PR #335](https://github.com/anthropics/skills/pull/335)

### #154 – shodh-memory (persistent context for AI agents)
- **What it does:** Provides a structured memory system so Claude can retain and surface relevant context across conversations via a `proactive_context` mechanism.
- **Discussion highlights:** Pushes toward long-running agent scenarios. Related to Issue #1329 (compact-memory proposal) and broader interest in agent state management.
- **Status:** Open  
  [PR #154](https://github.com/anthropics/skills/pull/154)

### #568 – ServiceNow platform skill
- **What it does:** Broad ServiceNow assistant covering ITSM, ITOM, ITAM, SecOps, HR/CSM, and IntegrationHub — not just scripting.
- **Discussion highlights:** Largest single-platform enterprise skill submitted. Demonstrates demand for comprehensive enterprise domain coverage.
- **Status:** Open  
  [PR #568](https://github.com/anthropics/skills/pull/568)

---

## 2. Community Demand Trends

From the most-commented Issues, the community’s pressing concerns and anticipated directions are:

| Trend | Key Issues | What the community wants |
|-------|------------|--------------------------|
| **Skill sharing & org management** | [#228](https://github.com/anthropics/skills/issues/228) (14 comments, 👍7) | Direct skill sharing within organizations instead of manual file transfer. Indicates enterprise adoption is bottlenecked by distribution. |
| **Evaluation tooling reliability** | [#556](https://github.com/anthropics/skills/issues/556) (12 comments, 👍7), [#1169](https://github.com/anthropics/skills/issues/1169) (3 comments) | `run_eval.py` consistently reports 0% trigger rate (recall=0%) — the optimization loop is broken. Multiple reproductions and fix PRs exist (#1298, #1099, #1050). |
| **Security & trust boundaries** | [#492](https://github.com/anthropics/skills/issues/492) (7 comments, 👍2) | Community skills distributed under `anthropic/` namespace enable trust abuse. Users want clear provenance and permission boundaries. |
| **Cross-platform (Windows) support** | [#1061](https://github.com/anthropics/skills/issues/1061) (3 comments), related PRs #1050, #1099 | `PATHEXT`, code page 1252, and `select()` on pipes break skill tooling on Windows. Linux-only assumptions are a barrier for many developers. |
| **MCP integration** | [#16](https://github.com/anthropics/skills/issues/16) (4 comments) | Expose skills as MCP tools for standardised API access. Long-standing request. |
| **Agent governance & safety** | [#412](https://github.com/anthropics/skills/issues/412) (6 comments) | Proposes a dedicated skill for agent safety patterns (policy enforcement, threat detection, audit). Aligns with PR #83 (security meta-skill). |
| **Deduplication / plugin structure** | [#189](https://github.com/anthropics/skills/issues/189) (6 comments, 👍9) | `document-skills` and `example-skills` plugins install identical content. Community wants clear separation and no context-window waste. |

**Most-anticipated new skill directions:** Agent governance, long-running memory systems, enterprise platform coverage (ServiceNow, SAP), and media generation. There is also strong demand for **meta-tooling** — quality analyzers, security analyzers, and reliable evaluation loops — before the ecosystem can scale safely.

---

## 3. High-Potential Pending Skills

These PRs have active discussion or represent critical fixes. They are likely to land soon.

- **#1298** – `fix(skill-creator): run_eval.py always reports 0% recall` (June 10–11) – Addresses the most-reported bug in the ecosystem. Includes Windows stream reading, trigger detection, and parallel worker fixes.  
  [PR #1298](https://github.com/anthropics/skills/pull/1298)

- **#538, #539, #541** – Series of fix PRs by Lubrsy706: case-sensitive file references in PDF skill, YAML quoting warnings in skill-creator, and DOCX tracked-change ID collision. Together improve core document skills and tooling validation.  
  [PR #538](https://github.com/anthropics/skills/pull/538) | [PR #539](https://github.com/anthropics/skills/pull/539) | [PR #541](https://github.com/anthropics/skills/pull/541)

- **#1050 & #1099** – Two independent Windows compatibility fixes for skill-creator scripts. Both are single-line changes and have been validated by users.  
  [PR #1050](https://github.com/anthropics/skills/pull/1050) | [PR #1099](https://github.com/anthropics/skills/pull/1099)

- **#361 & #362** – YAML detection and UTF-8 panic fixes for `quick_validate.py`. These prevent silent misparsing and Rust panics when descriptions contain special characters or multi-byte content.  
  [PR #361](https://github.com/anthropics/skills/pull/361) | [PR #362](https://github.com/anthropics/skills/pull/362)

- **#210** – `Improve frontend-design skill clarity and actionability` – Revision of an existing skill to make instructions specific and executable within a single conversation. Indicates a broader movement toward skill quality refinement.  
  [PR #210](https://github.com/anthropics/skills/pull/210)

---

## 4. Skills Ecosystem Insight

**The community’s most concentrated demand is on ecosystem reliability** — fixing the evaluation loop (0% recall bug), ensuring cross-platform support, securing trust boundaries, and improving skill quality tooling — **before scaling new skill domains** such as enterprise platforms (ServiceNow, SAP, ODT), agent memory/governance, and media generation.

---

# Claude Code Community Digest – 2026-06-19

## Today’s Highlights
Anthropic shipped **v2.1.183** with critical safety improvements for destructive Git operations, addressing a long-standing community pain point. Meanwhile, the community is grappling with a resurgence of **“No response from API” errors** (issues #69238, #69358) and a severe **subagent infinite‑recursion bug** (#68619) that can burn through tokens catastrophically. The most‑requested feature remains a **message queue mode** (#50246, 101 👍) to avoid interrupting active tasks.

## Releases
**v2.1.183** – [Release details](https://github.com/anthropics/claude-code/releases/tag/v2.1.183)
- **Improved auto‑mode safety:** Destructive Git commands (`git reset --hard`, `git checkout -- .`, `git clean -fd`, `git stash drop`) are now blocked when the user did not explicitly ask to discard local work.
- `git commit --amend` is blocked when the commit wasn’t created by the agent in this session.

## Hot Issues (10 of 30, updated in last 24h)

1. **[BUG] Hanging / freezing for 5–20 minutes** – [#26224](https://github.com/anthropics/claude-code/issues/26224) (122 comments, 143 👍)  
   Long‑standing severe performance issue; still no fix. ❗

2. **[Feature] Message queue mode** – [#50246](https://github.com/anthropics/claude-code/issues/50246) (33 comments, 101 👍)  
   Queue follow‑up messages instead of interrupting active work. Very high community demand.

3. **[BUG] Subagent infinite recursion + token burn** – [#68619](https://github.com/anthropics/claude-code/issues/68619) (8 comments, 2 👍)  
   Subagents spawn 50+ levels deep, ignoring `CLAUDE_CODE_FORK_SUBAGENT=0`. Permissions denials trigger more agents. **Critical cost risk.**

4. **[BUG] “No response from API” on Linux (v2.1.181)** – [#69358](https://github.com/anthropics/claude-code/issues/69358) (5 comments, 24 👍)  
   Constant API errors after latest update; users retrying for hours.

5. **[BUG] “No response from API” when Advisor triggered** – [#69238](https://github.com/anthropics/claude-code/issues/69238) (11 comments, 11 👍)  
   Advisor (Opus 4.8) fails; base model Sonnet works fine. Network‑check loop.

6. **[Bug] Ugrep wrapper causes V8‑heap OOM on WSL2** – [#54394](https://github.com/anthropics/claude-code/issues/54394) (6 comments)  
   Regex backtracking from embedded `ugrep` balloons memory to 8 GB, freezing the host.

7. **[Bug] Co-authored‑by trailer leaks user email** – [#66079](https://github.com/anthropics/claude-code/issues/66079) (2 comments)  
   Since v2.1.165, Git commit trailers reveal real email even when `git author.email` is a noreply. **Privacy concern.**

8. **[Bug] Chat history not persisting on mapped drives / OneDrive** – [#14088](https://github.com/anthropics/claude-code/issues/14088) (34 comments, 12 👍)  
   Windows users lose conversations; also affects VSCode extension (#63527). Long‑running.

9. **[Bug] autoMemoryEnabled=false still loads memory preamble** – [#63903](https://github.com/anthropics/claude-code/issues/63903) (15 comments)  
   Setting doesn’t suppress the 11–16k token memory preamble; inflates prompt costs.

10. **[Bug] OAuth token exchange fails with Entra ID** – [#69547](https://github.com/anthropics/claude-code/issues/69547) (3 comments)  
    Missing `scope` parameter crashes MCP server authentication on Windows.

## Key PR Progress (only 3 PRs updated in last 24h)

- **[CLOSED] Fix lock‑closed‑issues workflow** – [#69470](https://github.com/anthropics/claude-code/pull/69470)  
  Switches from offset pagination to search API, fixing a 53‑day streak of workflow failures.

- **[OPEN] resolve duplicate IPs** – [#45553](https://github.com/anthropics/claude-code/pull/45553) (stale, last updated 2026-06-18)  
  Unclear scope; no recent activity.

- **[OPEN] fix(scripts): break pagination when page not full** – [#68673](https://github.com/anthropics/claude-code/pull/68673) (last updated 2026-06-19)  
  Corrects an early‑exit pagination logic issue.

*PR activity is low today; the team may be focused on the hanging and API error regressions.*

## Feature Request Trends
- **Task pre‑emption control:** Message queue mode (#50246, 101 👍) and “stash” input (#38227, closed) both aim to let users compose ideas without interrupting the agent.
- **Model intelligence:** Automatic model switching for plan vs. execution (#15721) and exemption of `CLAUDE.md` from compaction (#68636) reflect desire for smarter context management.
- **Accessibility & UX:** Built‑in text‑to‑speech (#58429) and clipboard copy for `/btw` output (#39992) show demand for more inclusive tooling.
- **Plugin/skill granularity:** Per‑project OAuth tokens (#39952) and granular skill enable/disable (#47747) would reduce startup overhead and workspace conflicts.

## Developer Pain Points
- **Performance & reliability:** Hanging/freezing (#26224), “No response from API” regressions (#69238, #69358), and subagent recursion (#68619) are the top blockers – costing time and money.
- **Token waste:** Memory preamble ignores `autoMemoryEnabled=false` (#63903), repeated failed edits burn tokens (#60893), and Git `--no-stat` failures consume context (#13071).
- **Windows & storage:** Chat history loss on mapped drives (#14088, #63527) and OAuth failures (#69547) plague many enterprise users.
- **Privacy & safety:** Email leakage in commit trailers (#66079) and blocked destructive Git commands (now fixed in v2.1.183) show ongoing tension between automation and control.

*Data source: [github.com/anthropics/claude-code](https://github.com/anthropics/claude-code) – snapshot taken 2026-06-19.*

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest – 2026-06-19

## Today’s Highlights

Three new CLI alpha releases (0.142.0-alpha.2 through .4) landed today, but the community’s attention is on a wave of Windows sandbox regressions and a long-running Linux desktop app feature request that has topped 600 👍. Several security‑focused PRs around MITM proxy hardening and permission profile exposure also signal active infrastructure work.

## Releases

- **[rust-v0.142.0-alpha.2](https://github.com/openai/codex/releases/tag/rust-v0.142.0-alpha.2)**, **[.3](https://github.com/openai/codex/releases/tag/rust-v0.142.0-alpha.3)**, **[.4](https://github.com/openai/codex/releases/tag/rust-v0.142.0-alpha.4)** – All tagged within the last 24 hours. No changelog details provided; likely iterative bug fixes and internal refactoring.

## Hot Issues

1. **[#11023 – Codex desktop app for Linux](https://github.com/openai/codex/issues/11023)** (124 comments, 604 👍)  
   Still the most upvoted open feature request. Users report the macOS app is unusable due to an existing bug, and demand a native Linux build to avoid power‑consumption issues on MacBook Pros.

2. **[#25749 – Inaccessible legacy phone number blocks login](https://github.com/openai/codex/issues/25749)** (54 comments)  
   Users authenticated via Google OAuth cannot use Codex because it insists on verifying a stale phone number with no recovery path. High frustration for business users.

3. **[#18960 – Frequent WebSocket reconnect loop on macOS](https://github.com/openai/codex/issues/18960)** (46 comments)  
   Streaming failures force repeated reconnection. Still unresolved after two months; affects Pro subscribers heavily.

4. **[#13018 – Allow deletion of threads in Codex app](https://github.com/openai/codex/issues/13018)** (26 comments, 104 👍)  
   Closed but still highly requested. Users want native thread deletion instead of manual file removal from `~/.codex/archived_sessions/`.

5. **[#28988 – “Full Access” permission loop in latest macOS app](https://github.com/openai/codex/issues/28988)** (15 comments)  
   After updating to 26.614.11602, the app repeatedly asks for permission on every operation, breaking workflows.

6. **[#25246 – Business access tokens broken (401)](https://github.com/openai/codex/issues/25246)** (14 comments)  
   Enterprise customers can no longer generate access tokens via the documented endpoint. Affects CLI and CI/CD pipelines.

7. **[#28422 – Image generation regression: valid image not saved](https://github.com/openai/codex/issues/28422)** (14 comments)  
   In CLI 0.140.0, the status remains “generating” even after a valid image is produced; the file is never written.

8. **[#28978 – MCP invalid request after desktop app update](https://github.com/openai/codex/issues/28978)** (10 comments)  
   New conversations fail with “missing field `inputSchema`”. The CLI with the same config works, pointing to a desktop‑only regression.

9. **[#25667 – macOS app leaves ~965MB of `code_sign_clone` directories per launch](https://github.com/openai/codex/issues/25667)** (4 comments, 13 👍)  
   Significant disk waste; directories are not cleaned on quit.

10. **[#29084 – Git status spam causes CPU storm on macOS](https://github.com/openai/codex/issues/29084)** (2 comments)  
    Source‑control watcher in the desktop app spawns thousands of `git status` calls per second on workspaces with nested repos, maxing out fans.

## Key PR Progress

1. **[#29071 – Update Guardian policy wording](https://github.com/openai/codex/pull/29071)**  
   Refines security policy to reduce ambiguity around user‑authorized destructive operations and credential access.

2. **[#28942 – Add config toggles for orchestrator skills and MCP](https://github.com/openai/codex/pull/28942)** (closed)  
   Gives hosts the ability to independently disable orchestrator‑provided skills and MCP tools without affecting local skills.

3. **[#28489 – Add indexed web search mode](https://github.com/openai/codex/pull/28489)** (closed)  
   Introduces a `web_search = "indexed"` option that gates external web access behind a pre‑built index, reducing latency and cost.

4. **[#29086 – Document raw response item compatibility](https://github.com/openai/codex/pull/29086)** (closed)  
   Adds a note in `AGENTS.md` to treat raw response item events as compatibility‑sensitive during future changes.

5. **[#29013 – Protect managed MITM CA private keys from sandboxed commands](https://github.com/openai/codex/pull/29013)**  
   Ensures the CA private key under `$CODEX_HOME/proxy` is not readable by sandboxed processes, closing a security gap.

6. **[#29014 – Honor startup custom CA bundles with managed MITM](https://github.com/openai/codex/pull/29014)**  
   Makes the managed proxy preserve custom CA overrides (e.g., `SSL_CERT_FILE`) when rewriting child CA variables.

7. **[#28899 – Scope network approvals by environment](https://github.com/openai/codex/pull/28899)**  
   Network host approvals are now scoped to the execution environment, preventing unintended cross‑environment leakage.

8. **[#29075 – Batch skill discovery filesystem reads](https://github.com/openai/codex/pull/29075)**  
   Reduces expensive remote filesystem probes by batching directory scans and metadata reads during skill discovery.

9. **[#28859 – Expose resolved permission presets through app server](https://github.com/openai/codex/pull/28859)**  
   Experimental endpoint `permissionPreset/list` lets clients consume a resolved permission catalog instead of reconstructing policy client‑side.

10. **[#29067 – Namespace multi‑agent v2 tools under “collaboration”](https://github.com/openai/codex/pull/29067)**  
    Standardises multi‑agent v2 tools to `functions.collaboration.*`, aligning model‑visible hints with the actual tool surface.

## Feature Request Trends

- **Linux desktop app** (#11023, 604 👍) – remains the single most demanded feature.  
- **Thread lifecycle management** – delete threads (#13018), restore archived chats in the main UI (#27207, #20732), and explicit deletion of cloud sessions (#24610).  
- **Agent management UI** – a dedicated “Agent View” in CLI/TUI to monitor and control multiple agents (#22321).  
- **Customisable providers** – configurable `base_url` for Amazon Bedrock (#28902) and other non‑OpenAI backends.  
- **Web search modes** – indexed search (#28489) as a middle ground between cached and live.

## Developer Pain Points

- **Windows sandbox errors** – multiple reports of `codex-windows-sandbox-setup.exe` failing with “module not found” (#28982, #29089, #29072).  
- **Authentication/recovery dead ends** – legacy phone number verification (#25749) and broken business access tokens (#25246) block legitimate users.  
- **Performance regressions** – orphaned subagents freezing sessions (#19197), 965MB temporary directories (#25667), and git‑status‑induced CPU storms (#29084).  
- **Desktop‑specific regressions** – MCP input schema error (#28978), permission loops (#28988), and PowerShell encoding corruption (#13755, #29085).  
- **Context compaction issues** – VS Code extension reports 950k context but compacts around 260k (#29080).  
- **UI/UX annoyances** – invisible white tray icon on Windows light taskbar (#29037) and non‑focusing composer input (#28872).

*All links use the repository base `https://github.com/openai/codex`.*

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest – 2026-06-19

## Today's Highlights
No new releases landed in the last 24 hours, but the community remains active with 50 open issues and 29 pull requests updated. Critical agent-hang bugs (e.g., #21409) and subagent recovery misreporting (#22323) continue to draw attention, while security-focused PRs like atomic MCP OAuth token writes (#27664) and image MIME sniffing (#27850) signal growing maturity in the codebase.

## Releases
*No new releases in the last 24 hours.*

## Hot Issues (10 noteworthy picks)

1. **Generalist agent hangs forever** [#21409](https://github.com/google-gemini/gemini-cli/issues/21409) – P1, 7 comments, 8 👍. Users report that delegating to the generalist agent leads to indefinite hangs. Workaround (instructing the model not to use sub‑agents) points to a core orchestration problem. Community reaction is strong; this is a top candidate for immediate fix.

2. **Subagent recovery after MAX_TURNS falsely reports GOAL success** [#22323](https://github.com/google-gemini/gemini-cli/issues/22323) – P1, 6 comments. The `codebase_investigator` subagent returns `status: "success"` even when it hits the turn limit, hiding real interruptions. Misleading diagnostics make debugging harder.

3. **Gemini does not use skills and sub‑agents enough** [#21968](https://github.com/google-gemini/gemini-cli/issues/21968) – P2, 6 comments. Users with custom skills (e.g., gradle/git) find the model rarely invokes them autonomously. Suggests the orchestration prompt needs richer contextual hints.

4. **Shell command execution gets stuck after completion** [#25166](https://github.com/google-gemini/gemini-cli/issues/25166) – P1, 4 comments, 3 👍. Even trivial commands sometimes leave the TTY in “Awaiting user input” state. Impacts developer productivity significantly.

5. **Robust component‑level evaluations** [#24353](https://github.com/google-gemini/gemini-cli/issues/24353) – P1 (epic), 7 comments. Follow‑up to behavioral eval introduction; 76 tests already exist. Focus on expanding coverage for agent sub‑components.

6. **AST‑aware file reads, search, and mapping** [#22745](https://github.com/google-gemini/gemini-cli/issues/22745) – P2, 7 comments, 1 👍. Epic tracking whether AST‑aware tools can reduce turns, token noise, and improve codebase understanding. Promising direction for agent quality.

7. **Add deterministic redaction and reduce Auto Memory logging** [#26525](https://github.com/google-gemini/gemini-cli/issues/26525) – P2, 5 comments. Auto Memory sends transcripts to a model for extraction before redacting secrets – a privacy gap. Also logs skill contents unnecessarily.

8. **Stop Auto Memory retrying low‑signal sessions indefinitely** [#26522](https://github.com/google-gemini/gemini-cli/issues/26522) – P2, 5 comments. Sessions that don’t contain enough signal are never marked processed, causing repeated extraction attempts. Wastes quota and slows the pipeline.

9. **Browser subagent fails on Wayland** [#21983](https://github.com/google-gemini/gemini-cli/issues/21983) – P1, 4 comments, 1 👍. `browser_agent` reports GOAL but produces empty results on Wayland systems. Linux users face a hard block.

10. **Agent should discourage destructive behavior** [#22672](https://github.com/google-gemini/gemini-cli/issues/22672) – P2, 3 comments. The model occasionally uses `git reset --hard` or `--force` when safer alternatives exist. Community asks for guardrails around destructive commands.

## Key PR Progress (10 important pull requests)

- **fix(ci): append trailing slash to registry URL in npmrc** [#28038](https://github.com/google-gemini/gemini-cli/pull/28038) *(open)* – Fixes nightly release credential mapping by ensuring npm can match the registry URL.

- **fix(mcp): use longest‑prefix matching in parseMcpToolName for server names with underscores** [#28033](https://github.com/google-gemini/gemini-cli/pull/28033) *(open)* – Resolves incorrect tool routing when MCP server names contain underscores (fixes #27981).

- **fix(cli): throw FatalConfigError instead of process.exit in parseArguments** [#27987](https://github.com/google-gemini/gemini-cli/pull/27987) *(closed)* – Refactors argument parsing to avoid hard exit, enabling cleaner test and error handling.

- **fix(core): write MCP OAuth tokens atomically** [#27664](https://github.com/google-gemini/gemini-cli/pull/27664) *(closed)* – Uses temp-file + atomic rename to prevent token corruption on concurrent writes (fixes #27663, security‑critical).

- **fix(core): hide ignored folders from session context** [#27678](https://github.com/google-gemini/gemini-cli/pull/27678) *(closed)* – Respects `.gitignore`/`.geminiignore` in the directory tree provided to the model, improving privacy and focus.

- **feat(cli): add 'models' command to list available Gemini models** [#27848](https://github.com/google-gemini/gemini-cli/pull/27848) *(open)* – New `gemini models` subcommand outputting model names, context windows, and tiers; supports JSON output.

- **fix(core): sniff MCP image MIME types** [#27850](https://github.com/google-gemini/gemini-cli/pull/27850) *(open)* – Corrects misdeclared MIME types (e.g., WebP reported as PNG) by reading image headers, fixing model image‑processing issues (fixes #27731).

- **fix(cli): prompt for folder trust before auth** [#27845](https://github.com/google-gemini/gemini-cli/pull/27845) *(open)* – Adds early trust prompt to ensure workspace settings (including local MCP servers) are loaded under correct trust level (fixes #27844).

- **fix(core-tools): resolve Jupyter Notebook and JSON corruption in write_file** [#28000](https://github.com/google-gemini/gemini-cli/pull/28000) *(open)* – Prevents silent corruption of `.ipynb` and `.json` files; ensures round‑trip integrity for structured data.

- **chore(deps): pin dependencies and enforce 14‑day update cooldown** [#27948](https://github.com/google-gemini/gemini-cli/pull/27948) *(closed)* – Stricter dependency management: exact versions and a cooldown period to reduce supply‑chain risk and upgrade fatigue.

## Feature Request Trends
Several cross‑cutting themes emerge from the issues:

- **AST‑aware tooling** – Requests to integrate Abstract Syntax Tree awareness for file reads, search, and codebase mapping (e.g., #22745, #22746, #22747) to reduce token usage and improve agent precision.
- **Agent self‑awareness** – Users want the CLI to accurately describe its own mechanics, flags, and limitations (#21432) and to use its own agents/skills more proactively (#21968).
- **Memory system improvements** – Multiple issues (e.g., #26522, #26523, #26516) call for smarter session processing, better patch validation, and transparent error handling.
- **Destructive action guardrails** – #22672 and related requests ask for the model to avoid risky commands (`--force`, `git reset`) without explicit user approval.
- **Browser agent resilience** – Forks like #22232 and #21983 demand better session lock recovery, Wayland support, and respect for user configuration.

## Developer Pain Points
- **Widespread hanging** – The generalist agent (#21409) and shell command stuck (#25166) are top‑voted bugs, causing hours of lost productivity.
- **Subagent behaviour misreporting** – False “success” after turn limit (#22323) and silent permission violations (#22093) erode trust in multi‑agent workflows.
- **Memory system noise** – Infinite retries on low‑signal sessions (#26522), invalid patches silently skipped (#26523), and secret leakage before redaction (#26525) are recurring complaints.
- **Platform compatibility gaps** – The browser agent failure on Wayland (#21983) and terminal corruption on resize (#21924) affect Linux and terminal‑centric users.
- **Unwanted tmp scripts** – The model frequently writes temporary scripts in random directories (#23571), complicating workspace cleanup.
- **Configuration ignored** – Browser agent ignoring `settings.json` overrides (#22267) and tool registration failures with more than 128 tools (#24246) hinder advanced setups.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest – 2026-06-19

## Today’s Highlights
A mix of high-severity regressions and nuanced security concerns dominated the day. A WSL2 CPU-spin bug (Issue #3700) renders the TUI unusable on every fresh session, while a subtle `preToolUse` hook bypass under parallel tool calls (Issue #2893) raises alarm for enterprise safety policies. On the feature front, the community continues to push for project-scoped plugins (Issue #1665) and automatic model switching (Issue #2896). No new releases or pull requests were published in the last 24 hours.

## Releases
No new releases in the last 24 hours.

## Hot Issues (10 noteworthy)

1. **[#3700 – High‑severity WSL2 regression: 215% CPU idle, TUI frozen](https://github.com/github/copilot-cli/issues/3700)**  
   *Open* – After upgrading to 1.0.60, the CLI spins at ~215% CPU on WSL2 (most likely a regression of #2208). Live output never paints until restart. Immediate impact for WSL2 users.

2. **[#2893 – preToolUse hooks silently bypassed under parallel tool calls](https://github.com/github/copilot-cli/issues/2893)**  
   *Open* – When a `preToolUse` hook times out, the subprocess is not terminated and the CLI proceeds with an *implicit allow* fallback, then dispatches serial requests. This defeats the entire hook security model.

3. **[#3838 – Drive MCP OAuth not attached after successful reauth](https://github.com/github/copilot-cli/issues/3838)**  
   *Open* – OAuth flow succeeds, cache files are written, but subsequent tool calls lack auth credentials. Affects Homebrew-installed 1.0.63.

4. **[#1665 – Support Copilot CLI plugins scoped to project/repository](https://github.com/github/copilot-cli/issues/1665)**  
   *Closed* – Highly upvoted (👍17): users want to install plugins per repository rather than globally. A clear signal that enterprise teams need isolation.

5. **[#731 – Incompatibility with Zsh + direnv: Invalid session ID](https://github.com/github/copilot-cli/issues/731)**  
   *Closed* – A long-running bug where direnv (and nix-direnv) with Z shell causes “Invalid session ID” errors. Resolved after 6 months, but highlights environmental fragility.

6. **[#3859 – “Subconscious” sidekick keeps spawning even with memory disabled](https://github.com/github/copilot-cli/issues/3859)**  
   *Open* – The per-prompt memory “voting” agent fires on every prompt despite `/memory off` or `"memory": false`. Wasteful and confusing for users who opted out.

7. **[#3860 – Content-exclusion over-blocks entire working tree](https://github.com/github/copilot-cli/issues/3860)**  
   *Closed* – A severe block-state where even `/dev/null` and binaries were denied. Sticky across a session. Quickly fixed (closed same day), but demonstrates fragility of inclusion/exclusion rules.

8. **[#3855 – Scrolling no longer works in terminal](https://github.com/github/copilot-cli/issues/3855)**  
   *Closed* – Users cannot scroll back through conversation history, likely introduced with the full‑screen scrollbar in 1.0.61. Affects both tmux and non‑tmux terminals.

9. **[#3858 – Ctrl+Backspace doesn’t work on Windows](https://github.com/github/copilot-cli/issues/3858)**  
   *Open* – Standard Windows shortcut for “delete previous word” does nothing; only Alt+Backspace works (a Unix/macOS convention). Small but high‑frequency pain for Windows users.

10. **[#3864 – Plugin cache_path stored as absolute path breaks in Docker/multi‑HOME](https://github.com/github/copilot-cli/issues/3864)**  
    *Open* – Hardcoded `$HOME` paths prevent plugin hooks from firing when `~/.copilot` is volume‑mounted in Docker with a different HOME. Blocks container‑based development flows.

## Key PR Progress
No pull requests were updated in the last 24 hours.

## Feature Request Trends

- **Project/repo‑scoped plugins** (Issue #1665) – The most upvoted open feature. Teams want per‑project isolation without globally installed plugins.
- **Automatic model switching** (#2896) – A desire for the AI or system to switch models based on task complexity, reducing manual `/model` commands.
- **Per‑session directory access** (#3857) – Users want an intermediate “Yes, only for this session” option instead of permanent “Allow” or “Deny”.
- **LLM‑callable directory change tool** (#3865) – Let Copilot itself run `/cd` and update the status bar, making multi‑worktree workflows smoother.
- **Better documentation alignment** (#3861) – Users report that sandbox features (per‑host filtering, cross‑platform isolation) are documented as working but don’t function, indicating a documentation‑reality gap.

## Developer Pain Points

- **Performance regressions on specific platforms** – WSL2 CPU spinning (#3700) and Ubuntu 20.04 glibc incompatibility (#3296) show that platform coverage remains fragile.
- **Security/permission blips** – Hook bypasses (#2893), over‑blocking (#3860), and missing background‑agent hooks (#3013) erode trust in the safety model.
- **Input and rendering regressions** – Scrolling broken (#3855), `@` file expansion not working (#3854 / #3834), and Windows keyboard shortcuts (#3858) degrade the daily experience.
- **Session state poisoning** – Malformed attachments (#3791) and `/update` in resumed sessions (#3821) force restart or drop context.
- **Configuration inconsistency** – Disabled MCP servers still loaded (#3582), memory disabled but sidekick still runs (#3859), absolute paths in Docker (#3864) – configuration purity is a recurring friction point.

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-06-19

## Today's Highlights
The last 24 hours saw no new releases, but two open issues and one pull request dominated the conversation. A critical bug surfaced where `FetchURL` fails to respect system proxy settings, blocking users behind firewalls — a fix is already in review as PR #2461. Meanwhile, a Windows + Git Bash compatibility problem (#2462) highlights ongoing friction for VS Code extension users on non-Linux platforms.

**Full repo:** [MoonshotAI/kimi-cli](https://github.com/MoonshotAI/kimi-cli)

---

## Releases
None (no new releases in the last 24 hours).

---

## Hot Issues
*(Picked the 2 most noteworthy; only 2 updated in last 24h)*

1. **[#2455] [Bug] FetchURL 未读取系统代理，在被墙环境下无法访问外网**
   - **Author:** KuangYin-Z | **Created:** 2026-06-15 | **Updated:** 2026-06-18 | **Comments:** 2 | 👍: 0
   - **Why it matters:** This issue exposes a fundamental networking gap — `FetchURL` and `WebSearch` ignore system proxy environment variables (e.g., `HTTP_PROXY`, `HTTPS_PROXY`), making the CLI unusable behind corporate or regional firewalls where `curl` works fine. This is a blocker for developers in restricted environments.
   - **Community reaction:** Low engagement so far (0 upvotes), but the PR #2461 fix is already up, suggesting the maintainers are treating it seriously.
   - **Link:** [Issue #2455](https://github.com/MoonshotAI/kimi-cli/issues/2455)

2. **[#2462] [Bug] Windows + Git Bash: VS Code extension fails to extract bundled CLI because tar cannot handle zip**
   - **Author:** yplgame | **Created:** 2026-06-18 | **Updated:** 2026-06-18 | **Comments:** 0 | 👍: 0
   - **Why it matters:** VS Code extension bundling uses `tar` to extract the CLI binary, but on Windows + Git Bash (MSYS2), `tar` cannot handle `.zip` archives. This breaks the extension install process for a significant portion of Windows users who prefer Git Bash.
   - **Community reaction:** No comments yet, but this is a clear platform-specific regression.
   - **Link:** [Issue #2462](https://github.com/MoonshotAI/kimi-cli/issues/2462)

---

## Key PR Progress
*(Only 1 PR updated in last 24h)*

1. **[#2461] fix(net): honour system proxy env vars in aiohttp sessions**
   - **Author:** logicwu0 | **Created:** 2026-06-18 | **Updated:** 2026-06-18 | **Comments:** 0 | 👍: 0
   - **Summary:** This PR directly addresses Issue #2455. The root cause is that all outbound HTTP paths use `aiohttp.ClientSession` without integration with `proxy_env` from the standard library. The fix applies `proxy_env` configuration to `aiohttp` sessions, making `FetchURL` and `WebSearch` respect `HTTP_PROXY`/`HTTPS_PROXY` system variables.
   - **Why it matters:** This is the exact blocker for users behind proxies. If merged, it restores full functionality for corporate, educational, and VPN-connected environments.
   - **Link:** [PR #2461](https://github.com/MoonshotAI/kimi-cli/pull/2461)

---

## Feature Request Trends
*(Distilled from the 2 open issues)*

- **System proxy/network environment compatibility:** The most distinct request is for `FetchURL` and `WebSearch` to respect `HTTP_PROXY` / `HTTPS_PROXY` environment variables. This is critical for users behind firewalls, corporate proxies, and VPNs — a common scenario for developers in Asia and enterprises.
- **Cross-platform VS Code extension installation:** The Windows + Git Bash problem (#2462) suggests a need for better platform detection and archive handling (e.g., detecting whether to use `.tar.gz` vs `.zip` based on the shell environment).

---

## Developer Pain Points

1. **Network environment detection gaps:** The most acute frustration is the CLI's inability to read system proxy variables — a feature that should be automatic. Users expect `kimi-cli` to behave like `curl`, but it silently fails in proxied environments.
2. **Windows + Git Bash VS Code extension flakiness:** The extension's assumption that `tar` can extract `.zip` fails on Git Bash (MSYS2), causing a silent install failure. This suggests a lack of testing against common Windows development setups.

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-19

## Today's Highlights
The community remains highly active with 33 issues and 50 PRs updated in the last 24 hours, though no new releases were published. A cluster of five environment-specific bugs from a single user uncovered significant TUI gaps in GNU Screen (no truecolor, broken copy/paste), while the long‑standing plugin‑agent disappearance after desktop updates continues to generate discussion. On the PR side, a handful of targeted fixes are moving toward merging, including better MCP/AXI sidebar integration and a platform‑aware binary naming fix for npm users.

## Releases
No new releases in the last 24 hours. The latest published version remains **v1.17.8** (as referenced in several reports).

## Hot Issues (Top 10 by Impact & Community Engagement)

1. **[#28567 – Full MCP client capabilities](https://github.com/anomalyco/opencode/issues/28567)** *(24 👍, 17 comments)*  
   A high‑demand feature request pointing out that OpenCode’s MCP support lags behind the latest protocol spec. Users want parity with the standard, including resource subscription and improved tool discovery.

2. **[#10221 – Black screen on installation](https://github.com/anomalyco/opencode/issues/10221)** *(16 👍, 30 comments, closed)*  
   A long‑standing bug where fresh installs show only a black screen. Closed after a community workaround was documented, but the resolution remains unclear for many.

3. **[#30855 – Plugin‑registered agents disappear after desktop update (v1.15.13 → v1.16.0)](https://github.com/anomalyco/opencode/issues/30855)** *(0 👍, 7 comments, closed)*  
   Triggered by the Tauri‑to‑Electron architecture shift. Users report plugins register but only ‘Plan/Build’ agents appear. A similar report (#30903) echoes the frustration.

4. **[#31119 – Error: no such column: name](https://github.com/anomalyco/opencode/issues/31119)** *(5 👍, 5 comments)*  
   A database schema migration error affecting users upgrading from older versions. The SQLite column `name` is missing, causing the app to fail on launch.

5. **[#32986 – Failed to fetch models.dev ERROR on every startup](https://github.com/anomalyco/opencode/issues/32986)** *(0 👍, 2 comments)*  
   Every session logs an error trying to fetch `models.dev`. The domain resolves to `0.0.0.0`, pointing to a possible TLS or network configuration issue.

6. **[#32981 – Snapshot on home directory hangs opencode indefinitely](https://github.com/anomalyco/opencode/issues/32981)** *(0 👍, 2 comments)*  
   Running `opencode` from a 223 GB home directory with 20k+ git repos freezes for minutes. A companion PR (#32991) aims to skip huge untracked directories.

7. **[#32985 – opencode works poorly inside GNU Screen](https://github.com/anomalyco/opencode/issues/32985)** *(0 👍, 2 comments)*  
   Colors are wrong, copy/paste broken, and mouse support missing in `screen`. The reporter submitted five related issues (see #32984, #32987, #32982, #32983).

8. **[#32988 – Add interactive /settings command](https://github.com/anomalyco/opencode/issues/32988)** *(0 👍, 2 comments)*  
   A feature request for a TUI‑based settings editor to avoid manually editing `opencode.json`. Low community reaction but aligns with UX quality.

9. **[#32970 – Zen plan Claude models return “Error from provider (Anthropic)”](https://github.com/anomalyco/opencode/issues/32970)** *(0 👍, 2 comments, closed)*  
   All Claude models on the Zen plan fail with provider errors. The user had a positive balance, suggesting an upstream API or authentication issue.

10. **[#32972 – Thinking‑enabled model support](https://github.com/anomalyco/opencode/issues/32972)** *(0 👍, 1 comment)*  
    A feature request for better handling of reasoning models like Claude or minimax‑m3, where the thinking output is not surfaced properly in the UI.

## Key PR Progress (10 Important Pull Requests)

1. **[#32994 – feat(tui): surface AXI CLI tools alongside MCP servers in sidebar](https://github.com/anomalyco/opencode/pull/32994)**  
   Adds gh‑axi, npm‑axi, chrome‑devtools‑axi, and cluster‑ops‑axi tools to the TUI sidebar. Renames the section to “MCP/AXI” when tools are present.

2. **[#32937 – test(app): add manual performance diagnostics](https://github.com/anomalyco/opencode/pull/32937)**  
   Introduces an opt‑in Playwright suite for measuring session tab switching, repaint, and streaming latency. Excluded from normal CI to avoid noise.

3. **[#28116 – fix(snapshot): treat all candidates as ignored when git check-ignore …](https://github.com/anomalyco/opencode/pull/28116)** *(closed)*  
   Fixes snapshot performance by correctly ignoring files returned from `git check-ignore --stdin`. Addresses issue #28033.

4. **[#32962 – fix(tui): unify todo panel label to “Todos”](https://github.com/anomalyco/opencode/pull/32962)** *(closed)*  
   Resolves a UI inconsistency where the sidebar said “Todo” but the main view used “Todos”.

5. **[#28766 – fix(core): use platform‑aware binary naming in postinstall and publish](https://github.com/anomalyco/opencode/pull/28766)**  
   Prevents the npm package from shipping `bin/opencode.exe` on all platforms. Critical for non‑Windows users who hit “no such file” errors.

6. **[#28540 – fix(opencode): keep TUI interactive with piped stdin](https://github.com/anomalyco/opencode/pull/28540)**  
   Ensures `opencode` remains usable when stdin is piped (e.g., `echo "hello" | opencode`). Fixes a long‑standing usability regression.

7. **[#32991 – Don't git snapshot huge untracked directories](https://github.com/anomalyco/opencode/pull/32991)**  
   Prevents the almost‑hang when the current directory contains massive untracked folders. Directly addresses the hang reported in #32981.

8. **[#30729 – fix(opencode): point F# tree‑sitter queries to upstream ionide repo](https://github.com/anomalyco/opencode/pull/30729)**  
   Corrects the F# syntax‑highlighting queries from the stale `nvim-treesitter` fork to the official `ionide` repository. Fixes #28965.

9. **[#32952 – fix(plugin): fall back from unauthorized live client](https://github.com/anomalyco/opencode/pull/32952)** *(closed)*  
   When a plugin server returns 401/403, the client now falls back to an in‑process fetch instead of failing silently. Closes #31237.

10. **[#30636 – fix(storage): add session(time_updated) index for global session list](https://github.com/anomalyco/opencode/pull/30636)**  
    Adds a database index to match the `ORDER BY time_updated DESC` query pattern, improving performance of the global session list for users with many sessions.

## Feature Request Trends
- **MCP Standard Alignment**: The most‑upvoted request (#28567) calls for full compliance with the MCP specification, including resource subscriptions and tool annotations.
- **Interactive Configuration**: A `/settings` TUI command (#32988) would let users edit `opencode.json` without leaving the terminal.
- **Thinking / Reasoning Model Support**: Multiple users want proper rendering of chain‑of‑thought output from models like Claude and minimax (#32972).
- **AXI Tool Ecosystem**: The new PR #32994 suggests official integration of AXI CLI tools, indicating a growing interest in bundled agent‑specific utilities.

## Developer Pain Points
- **Terminal Compatibility**: GNU Screen lacks truecolor, mouse, and copy/paste support (#32985, #32984). GNU Screen users are a niche but vocal group.
- **Plugin Instability After Updates**: The Tauri‑to‑Electron migration broke plugin‑registered agents (##30855, #30903). Users are frustrated by repeated regressions.
- **Model Provider Errors**: Zen‑plan Claude models fail with “no provider available” or Anthropic errors (#30192, #32970). Users suspect rate limiting or credential issues.
- **WSL Integration Gaps**: Editor context syncing (VS Code + WSL) intermittently stops (#29570), and WSL credentials editing is still pending in PR #31256.
- **Silent Failures**: When no model is configured, prompts vanish without clear feedback (#32983). The `--print-logs` flag also omits runtime events (#32987), making debugging harder.
- **Startup Hangs & DB Migrations**: Large home‑directory snapshots (#32981) and missing SQLite columns (#31119) prevent first use for many developers.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest – 2026-06-19

## Today’s Highlights
Two patch releases landed in the last 24 hours (`v0.79.7` and `v0.79.8`), bringing automatic theme mode and selective provider base entry points to reduce bundle size. A data‑loss bug in the fuzzy `edit` tool sparked an immediate fix (PR #5898), and the community flagged several provider compatibility issues, including unavailable Copilot models and Moonshot/Kimi schema rejections.

---

## Releases

**v0.79.8**  
*Selective provider base entry points* – SDK users can now import `@earendil‑works/pi‑ai/base` and `@earendil‑works/pi‑agent‑core/base` and register only the providers they need, keeping unused transport code out of bundled applications. [Release notes](https://github.com/earendil-works/pi/releases/tag/v0.79.8)

**v0.79.7**  
*Automatic theme mode* – The `/settings` panel lets you choose separate light and dark themes, and Pi will follow your terminal’s colour‑scheme changes. *Self‑only updates* – The `/update` command (from #2729) now works when run inside the agent. [Release notes](https://github.com/earendil-works/pi/releases/tag/v0.79.7)

---

## Hot Issues (10 Noteworthy)

1. **#5897 – Unavailable models offered in Copilot integration**  
   When logging in with a Copilot subscription, Pi shows models (e.g., various Opus versions, GPT‑nano) that are not actually available to the user, leading to silent failures.  
   [Issue](https://github.com/earendil-works/pi/issues/5897) | 9 comments | closed

2. **#5899 – Fuzzy edit tool silently rewrites entire file, causing data loss**  
   If an `edit` matches only fuzzily (trailing whitespace, smart quotes, NFKC folding), the whole file is rewritten in normalized form, stripping trailing whitespace and folding characters – a critical data‑loss bug.  
   [Issue](https://github.com/earendil-works/pi/issues/5899) | 1 comment | open

3. **#5882 – Coding agent forgets previous turns after provider’s per‑request size runs out**  
   With Ollama, large file reads cause the agent to lose context of earlier conversation turns, making multi‑step tasks unreliable.  
   [Issue](https://github.com/earendil-works/pi/issues/5882) | 1 comment | closed

4. **#5871 – Anthropic OAuth‑token detection is hardcoded to `sk-ant-oat`**  
   Currently Pi only recognizes Anthropic OAuth tokens by a hardcoded substring, preventing users with custom‑prefixed credentials from using OAuth.  
   [Issue](https://github.com/earendil-works/pi/issues/5871) | 1 comment | open

5. **#5862 – Codex subscription error: “You exceeded your current quota” but Codex CLI works**  
   Pi’s integration with OpenAI Codex reports quota errors even when the official CLI works fine, indicating a ticket‑handling mismatch.  
   [Issue](https://github.com/earendil-works/pi/issues/5862) | 2 comments | closed

6. **#5845 – Three compaction‑related fixes**  
   A contributor reports inefficiencies in the compaction process (used to keep context under limits). The PR addresses three small but frequent issues, especially noticeable with local LLMs.  
   [Issue](https://github.com/earendil-works/pi/issues/5845) | 2 comments | closed

7. **#5842 – `npm-shrinkwrap.json` pins vulnerable `protobufjs` and `ws`**  
   The published shrinkwrap locks transitive dependencies with known DoS advisories (≤7.6.2 and <8.21.0). Downstream consumers cannot patch these.  
   [Issue](https://github.com/earendil-works/pi/issues/5842) | 1 comment | closed

8. **#5815 – Setting to suppress tmux extended‑keys warning**  
   Users want a configurable setting to silence the `Warning: tmux extended-keys is off` startup message, especially if they deliberately disable it.  
   [Issue](https://github.com/earendil-works/pi/issues/5815) | 2 comments | closed

9. **#5792 – Add built‑in ZAI China provider**  
   A request to add a `zai‑cn` provider targeting `open.bigmodel.cn`, sitting alongside the existing ZAI provider. Community shows strong interest in supporting Chinese‑based models.  
   [Issue](https://github.com/earendil-works/pi/issues/5792) | 2 comments | closed

10. **#5781 – Extension API: expose active tools as executable objects**  
    Extension authors need `getActiveExecutableTools(): AgentTool[]` (not just names) to fork the live tool set for observer-style extensions that run alongside the agent.  
    [Issue](https://github.com/earendil-works/pi/issues/5781) | 2 comments | closed

---

## Key PR Progress (10 Important Pull Requests)

1. **#5900 – `feat(coding-agent): emit OSC 9998/9999 for freecode-web adapter`**  
   Adds a WebBridge that translates AgentSession events into OSC frames so pi sessions show accurate status, cost, and context in the freecode‑web web UI.  
   [PR](https://github.com/earendil-works/pi/pull/5900) | closed

2. **#5898 – `fix(coding-agent): preserve untouched content in fuzzy edit matches`**  
   Directly addresses the data‑loss bug (#5899): when an edit matches only fuzzily, only the changed lines are rewritten; unmodified lines keep their original formatting.  
   [PR](https://github.com/earendil-works/pi/pull/5898) | closed

3. **#5509 – `feat: Add Amazon Bedrock Mantle OpenAI Responses provider`**  
   New provider for Bedrock Mantle (OpenAI Responses API), adding support for GPT‑5.5 and GPT‑5.4 models. Modelled after the existing Azure OpenAI provider.  
   [PR](https://github.com/earendil-works/pi/pull/5509) | open

4. **#5866 – `feat(ai): add OpenRouter Fusion alias`**  
   Adds `openrouter/fusion` as a synthetic router alias, matching the existing `openrouter/auto` pattern, so users can explicitly target Fusion routes.  
   [PR](https://github.com/earendil-works/pi/pull/5866) | closed

5. **#5348 – `Add selective pi-ai base entrypoints`**  
   Merges the base entry points (`@earendil-works/pi-ai/base` and `@earendil-works/pi-agent-core/base`) that allow side‑effect‑free imports and explicit provider registration – delivered in v0.79.8.  
   [PR](https://github.com/earendil-works/pi/pull/5348) | closed

6. **#5874 – `feat(coding-agent): add automatic theme mode`**  
   Implements separate light/dark theme support that follows the terminal colour scheme – delivered in v0.79.7.  
   [PR](https://github.com/earendil-works/pi/pull/5874) | closed

7. **#2408 – `fix(coding-agent): /model shows stale scoped models after models.json edit`**  
   The scoped model selector now updates correctly when `models.json` is modified without requiring a restart.  
   [PR](https://github.com/earendil-works/pi/pull/2408) | closed

8. **#5796 – `chore: bump TS target and lib to ES2024, use Promise.withResolvers()`**  
   Upgrades TypeScript target to ES2024 and replaces hand‑rolled `Promise.withResolvers()` implementations with the built‑in API. A clean infrastructure improvement.  
   [PR](https://github.com/earendil-works/pi/pull/5796) | closed

9. **#5812 – `fix(tui): protect pipe characters inside inline code in markdown tables`**  
   Fixes table rendering where a `|` inside backticks was incorrectly treated as a column separator, discarding content after it.  
   [PR](https://github.com/earendil-works/pi/pull/5812) | closed

10. **#5884 – `fix(ai): handle orphaned tool result messages to prevent Moonshot 400 errors`**  
    Adds guards to strip orphaned `tool`–role messages (tool results without a preceding `tool_calls` assistant message) – fixes Moonshot/Kimi API 400 errors.  
    [PR](https://github.com/earendil-works/pi/pull/5884) | closed

---

## Feature Request Trends

Across the last 24 hours, the most‑requested feature directions cluster around:

- **Provider coverage**: native support for Azure AI Foundry (Anthropic) (#5851), ZAI China (#5792), and Moonshot/Kimi compatibility fixes (#5822, #5870, #5884).  
- **Configurable UX**: separate light/dark themes (shipped), suppressible warnings (#5815), CLI flag validation (#5856), and configurable OAuth token detection (#5871).  
- **Memory & context management**: percentage‑based compaction triggers (#5848), max thinking levels for supported models (#5831), and exposing AgentHarness readiness for extensions (#5855).  
- **Extension API expansion**: executable tool objects (#5781), edit‑diff exposure (#5756), and a read‑only agent phase API (#5855).  
- **Optimisation & security**: prompt caching for Mistral (#5854), OpenRouter Fusion alias (#5866), and fixing pinned vulnerable dependencies (#5842).

---

## Developer Pain Points

Recurring frustrations visible in the issue tracker:

- **Provider compatibility breaks** – Non‑OpenAI providers (Moonshot, DeepSeek, Mistral) often reject Pi’s tool schemas or message formatting, forcing manual workarounds (#5822, #5811, #5884). The “unavailable models” bug in Copilot (#5897) wastes user time.
- **Data integrity risks** – The fuzzy‑edit rewrite bug (#5899) can silently corrupt files; compaction logic (#5845) occasionally drops context; and the “forgets turns” problem (#5882) breaks multi‑step coding workflows.
- **Terminal UI quirks** – Scrolling back to top during long output (#5839), counter‑intuitive tree navigation (#5225), missing tmux extended‑keys warning overrides (#5815), and poor JetBrains terminal capability detection.
- **CLI ergonomics** – Flags that silently accept missing values (#5856) and stale model selectors (#2408) degrade the command‑line experience.
- **Security and supply chain** – The published `npm-shrinkwrap.json` (#5842) blocks downstream patching of known‑vulnerable dependencies, forcing users to fork or wait for releases.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

Here is the **Qwen Code Community Digest** for June 19, 2026, based on the latest activity across the repository.

---

## Qwen Code Community Digest – 2026-06-19

### 1. Today's Highlights

Today’s activity is dominated by a wave of high-quality bug fixes addressing edge cases in parsing and caching, largely contributed by user **tt-a1i**, who filed several priority-2 (P2) issues covering Windows path handling, grep output parsing, and cache logic. On the feature front, extension management is maturing with a new web-shell UI and support for archive-based installs. The community also continues to push for granular UI customization, with a new optional tokens-per-second rate display now merged.

### 2. Releases

No new releases were published in the last 24 hours.

---

### 3. Hot Issues (10 selected)

1. **[#4814 – Feature Request: UI should make it easier for Custom Provider users to add new models (OPEN)](https://github.com/QwenLM/qwen-code/issues/4814)**
   - **Why it matters:** This is a high-visibility UX pain point. Onboarding wizard flows for custom providers are currently cumbersome, requiring manual configuration for each new model.
   - **Community reaction:** 5 comments, with users detailing the friction between first-launch wizard steps and ongoing model management.

2. **[#3361 – Agent misinterprets shell output as empty despite successful execution (OPEN)](https://github.com/QwenLM/qwen-code/issues/3361)**
   - **Why it matters:** A critical bug that breaks agentic workflows. The agent ignores visible shell output (e.g., from `git` commands), leading to incorrect downstream decisions.
   - **Community reaction:** Still open after 2 months; users have provided clear repro steps and screenshots.

3. **[#4951 – Are the token counts in the status bar accurate? (OPEN)](https://github.com/QwenLM/qwen-code/issues/4951)**
   - **Why it matters:** Trust in telemetry is foundational for cost-aware users. The reported inflation (hundreds of K tokens after a few exchanges) raises suspicion of double-counting.
   - **Community reaction:** 4 comments; users are seeking clarification on whether the metric is cumulative or per-turn.

4. **[#4616 – Model 'qwen3.7-max' not available for auth type 'openai' (CLOSED)](https://github.com/QwenLM/qwen-code/issues/4616)**
   - **Why it matters:** A configuration mismatch issue where model availability is tied to auth type. Highlights the need for better validation and model list discoverability.
   - **Community reaction:** Closed after 3 comments; resolved by clarifying model provider settings.

5. **[#5366 – Feature Request: Optional estimated response time (CLOSED)](https://github.com/QwenLM/qwen-code/issues/5366)**
   - **Why it matters:** Users want real-time feedback on generation progress. This is a follow-up to a related feature (#4899) that was recently implemented, indicating the dev team is responsive to UX feedback.
   - **Community reaction:** Closed quickly (3 comments, 1 day) after the feature was implemented as `ui.showResponseTokensPerSecond`.

6. **[#5390 – `web_fetch` rejects uppercase HTTP URL schemes (CLOSED)](https://github.com/QwenLM/qwen-code/issues/5390)**
   - **Why it matters:** A subtle but blocking bug for users with case-insensitive URL standards. It prevents fetching from URLs like `HTTPS://example.com`.
   - **Community reaction:** 3 comments; marked `welcome-pr` and fixed same day by the reporter.

7. **[#5386 – `SANDBOX_MOUNTS` misparses Windows drive paths (CLOSED)](https://github.com/QwenLM/qwen-code/issues/5386)**
   - **Why it matters:** A Windows-specific parity gap that breaks sandbox mount configurations for Windows developers.
   - **Community reaction:** 2 comments; quickly identified and fixed.

8. **[#5370 – Grep drops matches from file paths containing colons (CLOSED)](https://github.com/QwenLM/qwen-code/issues/5370)**
   - **Why it matters:** A fundamental parsing error in the grep tool that silently drops valid results when file paths contain colons (e.g., `C:\Users`).
   - **Community reaction:** 2 comments; fixed same day by the reporter (tt-a1i).

9. **[#5387 – Grep tools accept non-positive limits (CLOSED)](https://github.com/QwenLM/qwen-code/issues/5387)**
   - **Why it matters:** A validation gap that allows `limit=0` or `-1`, which can truncate or suppress all output unexpectedly.
   - **Community reaction:** 2 comments; fixed promptly.

10. **[#5363 – File search cache should not reuse prefix results for glob patterns (CLOSED)](https://github.com/QwenLM/qwen-code/issues/5363)**
    - **Why it matters:** A caching correctness bug. The `ResultCache` assumed simple prefix semantics, which is invalid for glob patterns like `*.js`.
    - **Community reaction:** 2 comments; identified and fixed swiftly.

---

### 4. Key PR Progress (10 selected)

1. **[#5407 – fix(core): target microcompaction cache disarms (OPEN)](https://github.com/QwenLM/qwen-code/pull/5407)**
   - **Overview:** Addresses token-efficiency refinements from a prior review (#4259). Improves cache eviction logic to avoid redundant file re-reads.
   - **Author:** tt-a1i

2. **[#5398 – feat(web-shell): add extension management (OPEN)](https://github.com/QwenLM/qwen-code/pull/5398)**
   - **Overview:** Brings `/extensions` management into the web-shell UI, including install, enable/disable, update checks, and a management UI.
   - **Author:** ytahdn

3. **[#5404 – fix(auth): preserve custom provider models on install (OPEN)](https://github.com/QwenLM/qwen-code/pull/5404)**
   - **Overview:** Prevents custom provider models from being overwritten during provider re-installation, with regression tests.
   - **Author:** tt-a1i

4. **[#5230 – fix(cli): gate cron scheduler startup on config initialization (OPEN)](https://github.com/QwenLM/qwen-code/pull/5230)**
   - **Overview:** Fixes a race condition causing "Chat not initialized" errors when durable cron tasks fire at startup.
   - **Author:** qwen-code-ci-bot

5. **[#4909 – feat(extensions): support archive install sources (OPEN)](https://github.com/QwenLM/qwen-code/pull/4909)**
   - **Overview:** Allows installing extensions from `.zip` and `.tar.gz` archives, including remote URLs, reusing existing validation and extraction pipelines.
   - **Author:** kkhomej33-netizen

6. **[#5126 – feat(vision-bridge): transcribe images to text for text-only models (OPEN)](https://github.com/QwenLM/qwen-code/pull/5126)**
   - **Overview:** An opt-in bridge that routes image inputs to a multimodal model, transcribes them to text, and passes text to the primary text-only model.
   - **Author:** yiliang114

7. **[#4971 – fix(cli): reduce retained interactive tool output memory (CLOSED)](https://github.com/QwenLM/qwen-code/pull/4971)**
   - **Overview:** Compacts large tool-output display metadata after rendering, reducing memory pressure for long sessions with heavy tool use.
   - **Author:** kkhomej33-netizen

8. **[#5392 – feat(cli): serve the Web Shell UI from `qwen serve` (CLOSED)](https://github.com/QwenLM/qwen-code/pull/5392)**
   - **Overview:** Bundles the web shell SPA into the `qwen serve` command, enabling a single-process, single-port browser UI out of the box.
   - **Author:** wenshao

9. **[#5401 – feat(cli): show optional response token rate (CLOSED)](https://github.com/QwenLM/qwen-code/pull/5401)**
   - **Overview:** Adds an opt-in `ui.showResponseTokensPerSecond` setting to display real-time tokens per second alongside the cumulative counter.
   - **Author:** tt-a1i

10. **[#5396 – fix(ui): reduce UI flicker — throttle + compact transition + batch STREAM_TEXT (OPEN)](https://github.com/QwenLM/qwen-code/pull/5396)**
    - **Overview:** A triple-threat fix for UI flicker: throttling stream events, using `startTransition` for compact mode, and batching stream text updates.
    - **Author:** aspnmy

---

### 5. Feature Request Trends

Across all open issues, the following feature directions are most requested:

- **Provider & Model Flexibility:** Users repeatedly ask for easier model addition for custom/third-party providers, better model list discoverability, and presets for new models (e.g., `glm-5.2`).
- **UI/UX Customization:** Demand is strong for configurable status lines (token counts, response rates, estimated time remaining) and reduced visual clutter (compact mode improvements).
- **Token Tracking & Efficiency:** Users want accurate, configurable, and real-time token counters to manage costs. Issues like #4951 and #5401 reflect this trend.
- **Extension Ecosystem:** Growing interest in installable extensions from archives, remote URLs, and a richer management UI (tabs for Discover, Installed, Sources).
- **Diagnostic & Traceability Tools:** Features like response rate display and better error messages for model availability suggest a need for more runtime observability.

---

### 6. Developer Pain Points

Recurring frustrations and high-frequency requests from today’s data:

- **Model availability & provider configuration confusion:** #4616 and #4814 highlight that finding and setting up models—especially for custom or non-standard providers—is still confusing.
- **Token count accuracy & transparency:** Users are distrustful of status line token metrics (#4951), indicating a need for clearer documentation or a per-turn/total toggle.
- **Windows compatibility gaps:** Issues #5386 (sandbox mounts), #5370 (grep path parsing) show that Windows users still encounter platform-specific parsing bugs.
- **String parsing edge cases:** A wave of reports by a single contributor (tt-a1i) around colon and equals-sign parsing (URL schemes, file paths, environment variables) indicates a systematic weakness in string handling in tools.
- **Environment variable handling in CLI:** The `mcp add` issue (#5374) shows a common developer pain point: environment values containing `=` are truncated, which breaks secrets or complex config values.
- **Cache and state management correctness:** Bugs in cache prefix matching (#5363) and startup race conditions (#5230) point to lingering reliability issues in the core session and cache layers.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-06-19

## Today's Highlights
The community saw a surge of structural refactoring issues and PRs aimed at breaking down monolith Rust files (mcp.rs, config.rs, app.rs, ui.rs) into focused modules. Concurrently, three top-priority bugs — TUI freezes on Windows, spurious “turn stalled” errors, and sub-agent UI hangs — remain open with active discussion. A first-class sub‑agent toggle PR (#3327) was opened after prior closed attempts, signaling strong demand for runtime sub‑agent control.

## Releases
No new releases in the last 24 hours.

## Hot Issues (Top 10)

1. **[#2487] Frequent error: Turn stalled – no completion signal received**  
   *Bug/enhancement, 16 comments*  
   Yolo‑mode operations freeze with “Turn stalled” messages; `continue` does not resume. The community is debugging signal‑timeout logic in the runtime dispatcher.  
   [Issue #2487](https://github.com/Hmbown/CodeWhale/issues/2487)

2. **[#1812] TUI‑freeze‑Windows‑crossterm‑poll**  
   *Bug, 7 comments*  
   v0.8.39 freezes entirely on Windows 11 (UI unresponsive, process alive). Logs and thread dumps point to crossterm polling deadlock. Two confirmed events recorded.  
   [Issue #1812](https://github.com/Hmbown/CodeWhale/issues/1812)

3. **[#2870] EPIC: staged command-boundary refactor**  
   *Documentation/cleanup, 6 comments*  
   Tracks incremental PRs for the `#2791` command‑boundary refactor, targeting v0.9. Already has a proof‑of‑concept PR (#2851).  
   [Issue #2870](https://github.com/Hmbown/CodeWhale/issues/2870)

4. **[#3275] CodeWhale over‑involvement – self‑questioning and self‑answering**  
   *Bug/security, 5 comments*  
   Regression from #3061: the agent enters a self‑driven loop of proposing, answering, and executing without user confirmation. Discussion on provenance and approval redirection.  
   [Issue #3275](https://github.com/Hmbown/CodeWhale/issues/3275)

5. **[#3238] Does not work on Ubuntu 22.04 LTS (glibc mismatch)**  
   *Bug/documentation, 4 comments*  
   `npm install -g codewhale` fails due to glibc version incompatibility. Community requests a pre‑compiled binary or static linking.  
   [Issue #3238](https://github.com/Hmbown/CodeWhale/issues/3238)

6. **[#3289] v0.8.61 UI freezes after auto‑spawning several sub‑agents**  
   *Bug, 4 comments*  
   Plan mode spawns multiple sub‑agents; the interface becomes unresponsive. Already logged in #2487 but with different trigger.  
   [Issue #3289](https://github.com/Hmbown/CodeWhale/issues/3289)

7. **[#1917] Proposal: universal PreToolUse/PostToolUse hook layer**  
   *Enhancement, 4 comments*  
   Architectural proposal to add Cancel/Pause/Resume hooks for all actions, unifying bug workarounds found across #1886–#1900.  
   [Issue #1917](https://github.com/Hmbown/CodeWhale/issues/1917)

8. **[#3240] Legacy .deepseek configuration directory still created**  
   *Bug, 3 comments*  
   Despite rename to CodeWhale, the program creates both `.codewhale` and `.deepseek` directories on Windows. Users want cleanup.  
   [Issue #3240](https://github.com/Hmbown/CodeWhale/issues/3240)

9. **[#2973] WhaleFlow: real async executor – replace MockWorkflowExecutor**  
   *Enhancement, 2 comments*  
   Milestone for v0.9: implement production‑grade bounded concurrent branches, cooperative cancellation, token budgets, and control‑flow nodes.  
   [Issue #2973](https://github.com/Hmbown/CodeWhale/issues/2973)

10. **[#3328] 0.8.62 sidebar missing**  
    *Question, 1 comment*  
    After upgrading, the sidebar disappears; `/sidebar` responds but no visual. Likely a TUI rendering regression.  
    [Issue #3328](https://github.com/Hmbown/CodeWhale/issues/3328)

## Key PR Progress (Top 10)

1. **[#3333] refactor(tui): split MCP header helpers**  
   Extracts HTTP header framing from `mcp.rs` into a dedicated `mcp::headers` module, making the upcoming MCP transport split (#3310) incremental.  
   [PR #3333](https://github.com/Hmbown/CodeWhale/pull/3333)

2. **[#3330] Layer 4: replay FEAT-005 command extraction on Hunter**  
   Semantic replay of command extraction against the trait‑backed registry, targeting `hunter/0.8.62-glm-subagents` as part of the command‑boundary refactor EPIC.  
   [PR #3330](https://github.com/Hmbown/CodeWhale/pull/3330)

3. **[#3331] fix(tui): enable proxy env for js execution**  
   Fixes #3273: passes proxy environment variables (`HTTP_PROXY`, `ALL_PROXY`, etc.) into the Node child process so `js_execution` respects system proxy settings on Windows.  
   [PR #3331](https://github.com/Hmbown/CodeWhale/pull/3331)

4. **[#3332] fix(app-server): require auth for non-loopback binds**  
   Fixes #3258: rejects `app-server` binds on `0.0.0.0` unless an explicit `--auth-token` is provided, preventing accidental public exposure.  
   [PR #3332](https://github.com/Hmbown/CodeWhale/pull/3332)

5. **[#3329] fix(config): restore huggingface env precedence**  
   Restores Hugging Face API key environment precedence in TUI config to pass the CI `check-provider-registry.py` lint gate.  
   [PR #3329](https://github.com/Hmbown/CodeWhale/pull/3329)

6. **[#3300] feat(tui): preserve thinking/tool blocks when seeding thread from session**  
   Replaces text‑only seeding with block‑type‑aware implementation, preserving `Thinking`, `ToolUse`, and `ToolResult` variants across thread resumes.  
   [PR #3300](https://github.com/Hmbown/CodeWhale/pull/3300)

7. **[#3321] fix(workflow): add token budget regulator for high fan-out agent runs**  
   Closes the enforcement gap between `BudgetSpec` and runtime execution: adds token‑based budget regulation for workflows and sub‑agent orchestration.  
   [PR #3321](https://github.com/Hmbown/CodeWhale/pull/3321)

8. **[#3327] v0.8.63: Add first-class sub-agent toggle**  
   Introduces `/config subagents on|off|status` and `/config features.subagents` commands; session‑only changes are wired through `AppAction::UpdateFeatures`.  
   [PR #3327](https://github.com/Hmbown/CodeWhale/pull/3327)

9. **[#3316] v0.8.63: Add source wiki and Agents/Workflows terminology**  
   Adds generated source wiki at `wiki/` and a terminology documentation (`ORCHESTRATION_TERMINOLOGY.md`) to align public naming around “Agents” and “Workflows”.  
   [PR #3316](https://github.com/Hmbown/CodeWhale/pull/3316)

10. **[#3301] feat(tui): save ask permission rules from approvals**  
    Adds an “ask‑only” action in the approval modal that persists a shell‑approval rule as a `permissions.toml` entry, including a TOML preview.  
    [PR #3301](https://github.com/Hmbown/CodeWhale/pull/3301)

## Feature Request Trends

- **Sub‑agent runtime control** – multiple issues request a simple on/off switch, editable recursion/concurrency limits, and status indicators in the TUI (#3304, #3305).  
- **Architectural refactoring** – a strong push to split large Rust files (config.rs, mcp.rs, app.rs, ui.rs, runtime_threads.rs) into domain‑owned modules for maintainability (#3306–#3314).  
- **WhaleFlow async executor** – the community expects production‑grade bounded concurrency with token budgets and cancellation support (#2973, #3230).  
- **Better provider management** – removing the ballooning config/enum pattern in favor of a registry (#2608) and adding Aliyun Bailian integration (#3320).  
- **Proxy/environment support** – ensuring all execution paths (js_execution, app‑server) respect system proxy variables (#3273, #3331).  

## Developer Pain Points

- **Frequent stalls and freezes** – TUI turns stale (#2487, #3289) and Windows‑specific polling deadlocks (#1812) are the most frustrating bugs, with no fix yet.  
- **glibc compatibility** – Ubuntu 22.04 users cannot install via npm due to glibc version mismatch (#3238); static builds are requested.  
- **Over‑eager LLM behavior** – the agent’s self‑driven loops bypassing user intent (#3275) raises trust and scope concerns, requiring provenance enforcement.  
- **Legacy configuration mess** – leftover `.deepseek` directories and dual config folders on Windows create confusion (#3240).  
- **Monolith contribution barriers** – large files (config.rs >9k lines) make adding new providers or UI elements error‑prone; the refactoring effort is welcomed but adds short‑term merge friction.  
- **Missing UI parity** – sidebar disappearance in v0.8.62 (#3328) and missing editable sub‑agent controls (#3304) show ongoing TUI polish gaps.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*