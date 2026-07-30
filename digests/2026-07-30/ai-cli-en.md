# AI CLI Tools Community Digest 2026-07-30

> Generated: 2026-07-30 00:11 UTC | Tools covered: 9

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

# Cross-Tool Comparison Report – 2026-07-30

## Ecosystem Overview

The AI CLI tools landscape on July 30, 2026, reflects a mature but rapidly evolving ecosystem. Communities are overwhelmingly focused on enterprise readiness—security, sandboxing, credential management, and API gateway integration—while also pushing for smarter automation (autonomous PR verification, intelligent model routing, and subagent orchestration). Cross-platform parity remains a persistent pain point, especially for Windows and Linux ARM64 users. Tools like Qwen Code and Pi demonstrate self-improving loops where the tool itself (autofix, caretaker) generates substantial PR activity, signaling a shift toward meta-automation within developer tools.

## Activity Comparison

| Tool | Hot Issues Highlighted | Active PRs (24h) | Release(s) Today |
|------|-----------------------|------------------|-------------------|
| **Claude Code** | 10 (top by votes) | 4 open PRs | None |
| **OpenAI Codex** | 10 | 10 open PRs | v0.146.0 + alpha patches |
| **Gemini CLI** | 10 | 10 open PRs (incl. important infra PRs) | v0.55.0-nightly |
| **GitHub Copilot CLI** | 10 | 1 open PR (#4100) | v1.0.76-2 through -5 (5 point releases) |
| **Kimi Code CLI** | 1 issue updated | 6 PRs (2 open, 4 merged) | None |
| **OpenCode** | 10 | 10 open PRs | None |
| **Pi** | 10 | 10 open PRs | v0.83.0 |
| **Qwen Code** | 3 issues updated | ~50 active PRs (autofix takeover) | v0.21.0-nightly |
| **DeepSeek TUI** | 10 | 10 PRs (mostly closed for v0.9.2) | Release candidate (no final version) |

*Notes:* Issue/PR counts are approximated from digest summaries. Qwen Code’s PR surge is driven by its autofix bot; Copilot’s multiple point releases show rapid iteration; Gemini and Pi have high infrastructure PR activity.

## Shared Feature Directions

Across multiple tool communities, these requirements recur:

- **Multi-workspace & multi-tenant integrations**  
  Claude Code (#44243 – Slack), OpenCode (#39579 – ACP session/list), and GitHub Copilot (#4204 – `.agents` discovery expansion) all request better handling of multiple workspaces/projects without per-session friction.

- **Intelligent model routing**  
  Claude Code (#15721 – automatic model switching for plan/execute), Pi (#6951 – Qwen reasoning effort mapping), and OpenCode (#38851 – gpt-5.6-sol compaction) want smarter usage of cheaper/faster models and context-aware model selection.

- **Hook parity & lifecycle automation**  
  OpenAI Codex (#21753, #17148 – pre/post compact hooks), Kimi Code (#2176 – UserPromptSubmit hook), and Pi (#5329 – expose waiting state for host integrations) underscore demand for programmable automation surfaces.

- **Security & credential hygiene**  
  Claude Code (#82358 – MCP Guard plugin), OpenAI Codex (#36049 – metrics export), Gemini CLI (#28557 – SSRF fix), OpenCode (#39576 – subagent permission escalation), and Pi (#1871 – lock contention) show widespread concern over secret leaks, authorization fatigue, and sandbox bypasses.

- **Session management & persistence**  
  OpenAI Codex (#35420 – OneDrive stream disconnects), Gemini CLI (#22323 – false success on max turns), Pi (#7285 – non-reactive resumed sessions), and DeepSeek TUI (#4941 – reasoning_effort reset) highlight reliability gaps in session state, recovery, and compaction.

- **Cross-platform fixes (Windows, macOS, Linux)**  
  Issues like Claude Code’s `ENAMETOOLONG` (#72725), GitHub Copilot’s blank Windows Terminal (#4159), OpenCode’s ARM64 TUI crash (#19130), and DeepSeek TUI’s Brazilian ABNT2 keyboard (#4723) show platform-specific bugs are still a common frustration.

## Differentiation Analysis

| Tool | Distinctive Focus | Target User | Technical Approach |
|------|-------------------|-------------|-------------------|
| **Claude Code** | MCP ecosystem, Slack integration, role-playing bug | Professional developers in large orgs | Heavy community-driven feature requests; slow release cadence but high engagement |
| **OpenAI Codex** | Session management (/new, pin), agent plugins (Bedrock, Claude C), GPT-5.6 optimization | Power users of OpenAI models | Rapid alpha/beta releases; strong emphasis on TUI + VSCode extension parity |
| **Gemini CLI** | Subagent reliability, automated PR generation (SSR pipeline), Auto Memory | Google Cloud developers | Infrastructure-heavy; caretaker agent and Firestore-backed pipeline for self-healing |
| **GitHub Copilot CLI** | Git worktree lifecycle, sandbox configurability, plugin enable/disable | GitHub-centric teams | Many point releases; focuses on terminal UX (paste, colors, session sidebar) |
| **Kimi Code CLI** | Enterprise K3 gateway, tool chaining (StrReplaceFile), Windows pwsh | Enterprise teams using Kimi K3 | Low community volume but targeted issue/PR activity; pragmatic fixes |
| **OpenCode** | TUI compaction, MCP timeout caps, permission automation, i18n (Hebrew) | Agent protocol (ACP) adopters | Active bug reporting; new contributors; experiments with RTL and project picker |
| **Pi** | Credential export, Markdown API, sixel images, search index (FTS5) | CLI power users and extension developers | High PR velocity; comparative eval harness indicates maturity; broad provider support |
| **Qwen Code** | Autonomous PR verification (autofix), Fleet Shepherd, transcript integrity | AI-driven CI/CD teams | Self-improving bot ecosystem; huge PR volume from bot itself; CI optimisation focus |
| **DeepSeek TUI** | Localization (Indonesian, Chinese), LaTeX math, permission rules, `/stop` command | Multilingual developers | Release-candidate stability; community contributors drive most features |

## Community Momentum & Maturity

- **Rapid iterators** – **Pi** (v0.83.0 + 23 PRs in 24h), **GitHub Copilot CLI** (5 point releases), **Qwen Code** (50+ PRs from autofix). These tools ship frequently and have high contributor throughput.
- **High community engagement** – **Claude Code** issues draw 60–74 upvotes, indicating a large but less frequent release cycle. **OpenAI Codex** has strong feedback on hooks and model bugs. **Gemini CLI** shows intense discussion on subagent reliability.
- **Emerging players** – **Kimi Code CLI** has low raw activity but enterprise demand is clear. **DeepSeek TUI** is nearing a stable release with strong localization contributions. **OpenCode** is gaining traction among ACP users.
- **Mature automation** – **Qwen Code** and **Gemini CLI** (caretaker) are building autonomous PR generation pipelines, suggesting these tools are evolving into self-maintaining platforms.

## Trend Signals

1. **Self-improving tool loops** – Qwen Code’s autofix bot generating PRs to fix its own inefficiencies, and Gemini’s SSR pipeline for automated code generation, point toward a future where CLI tools actively maintain themselves.

2. **Enterprise API gateways** – Kimi Code (#2568) and Pi (#7199 – Kimi K3 on Fireworks) reflect a push for custom API endpoints with rate limiting, failover, and centralized key management—critical for production deployments.

3. **Security as a first-class concern** – Multiple SSRF fixes, credential export controls, sandbox permission bypass patches, and MCP secret masking indicate that tool security is no longer an afterthought. Expect more audit-trail and attestation features (e.g., DeepSeek TUI #4958).

4. **Localization beyond English** – DeepSeek TUI’s Indonesian and Traditional Chinese efforts, OpenCode’s Hebrew PR, and Kimi’s Chinese-language issues show a growing user base outside English-speaking markets.

5. **Model cost optimization** – Claude Code’s automatic model switching (#15721), OpenAI Codex’s GPT-5.6 serialization bug (#35050), and Pi’s reasoning effort mapping (#6951) all aim to reduce API consumption without sacrificing quality.

6. **Terminal UI escalation** – Demands for richer rendering (LaTeX, sixel images, RTL text, project pickers) signal that users expect TUI experiences to rival graphical IDEs. This pushes tool teams to invest in terminal graphics protocols and responsive layouts.

**For developers evaluating these tools:** consider your platform (Windows/Linux/macOS), model provider preferences, and need for enterprise features. Tools like Pi and Copilot are most agile across providers; Claude Code and Qwen Code excel in deep integration with their respective ecosystems. OpenCode and Kimi Code are worth watching for ACP and enterprise gateway scenarios respectively.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
**Data snapshot: 2026-07-30 | Source: github.com/anthropics/skills**

---

## 1. Top Skills Ranking

Most-discussed Skill PRs by community engagement:

### #514 — Document Typography Quality Control
**Functionality:** Prevents orphan word wrap, widow paragraphs, and numbering misalignment in AI-generated documents — common typographic defects users rarely request correction for.
**Discussion highlights:** Recognized that these issues affect every Claude-generated document. The "invisible" nature of the problem (users don't know to ask for fixes) made this a high-impact addition. One of the most active PRs in the repository.
**Status:** [OPEN](https://github.com/anthropics/skills/pull/514)

### #486 — ODT (OpenDocument) Skill
**Functionality:** Creates, fills, reads, and converts ODF files (.odt, .ods) including template filling and ODT-to-HTML conversion. Triggers on any mention of LibreOffice, OpenDocument, or open-source document formats.
**Discussion highlights:** Fills a clear gap for users in open-source and European institutional environments where ODF is standard. Community questions focused on template rendering fidelity and embedded image handling.
**Status:** [OPEN](https://github.com/anthropics/skills/pull/486)

### #723 — Testing Patterns Skill
**Functionality:** Comprehensive testing methodology skill covering the Testing Trophy model, AAA pattern, React Testing Library, Playwright E2E, and what *not* to test. Designed to standardize Claude's testing guidance across the full stack.
**Discussion highlights:** Community debated whether a single skill should prescribe testing philosophy versus listing tool-specific commands. Resolved by including both: philosophy as behavioral guardrails, tooling as executable instructions.
**Status:** [OPEN](https://github.com/anthropics/skills/pull/723)

### #525 — Pyxel Retro Game Development
**Functionality:** Adds support for the Pyxel retro game engine via its MCP server. Covers the full iterative workflow: write code, capture screen output, inspect, iterate.
**Discussion highlights:** Notable because it's the only skill tied to an MCP server rather than pure Claude instructions. Community raised questions about MCP dependency requirements and whether skills should require external services.
**Status:** [OPEN](https://github.com/anthropics/skills/pull/525)

### #210 — Frontend Design Skill Improvement
**Functionality:** Major revision of the frontend-design skill to improve actionability and internal coherence. Goal was ensuring every instruction is something Claude can follow within a single conversation.
**Discussion highlights:** Deep discussion about skill design philosophy — should skills teach Claude new capabilities or constrain Claude's existing behavior? The PR trended toward behavioral constraints with concrete examples.
**Status:** [OPEN](https://github.com/anthropics/skills/pull/210)

### #83 — Skill Quality & Security Analyzers (Meta Skills)
**Functionality:** Two meta-skills: `skill-quality-analyzer` evaluates across Structure/Documentation, Trigger Precision, Instruction Clarity, Robustness, and Security. `skill-security-analyzer` evaluates prompt injection risks, data exfiltration, and trust boundaries.
**Discussion highlights:** The meta-skill concept generated significant interest — skills that audit other skills. Security analyzer directly addresses the #492 namespace trust issue. Quality analyzer's five-dimension scoring framework was debated but largely endorsed.
**Status:** [OPEN](https://github.com/anthropics/skills/pull/83)

### #1367 — Self-Audit Skill (Mechanical Verification + Reasoning Gate)
**Functionality:** A universal output audit skill that runs mechanical file verification (every claimed file exists) then a four-dimension reasoning quality audit in damage-severity priority order. Works with any project and any model.
**Discussion highlights:** The "reasoning quality gate" concept — adversarial review of Claude's own output before delivery — received strong community support. Related discussion in Issue #1385 proposes expanding this to a three-gate pipeline.
**Status:** [OPEN](https://github.com/anthropics/skills/pull/1367)

---

## 2. Community Demand Trends

From the most active Issues, the community's highest-priority needs are:

**Skill Creator Tooling Reliability:** Issues #556, #1169, and #1061 all report that `run_eval.py` consistently returns 0% recall, making the description-optimization loop optimize against noise. This is the single most-complained-about technical issue — 7+ independent reproductions, 12+ comments on #556 alone, spanning Windows and Unix. The tooling ecosystem cannot function without this fix.

**Security & Trust Boundaries:** Issue #492 (43 comments, highest engagement) identifies that community skills distributed under `anthropic/` namespaces create trust boundary vulnerabilities. Users may grant elevated permissions to skills they believe are official. The #83 skill-security-analyzer directly responds to this.

**Organizational Skill Sharing:** Issue #228 (16 comments, 8 upvotes) demands org-wide sharing without manual file transfer. Users want a shared skill library or direct sharing links in Claude.ai.

**Windows Compatibility:** Issues #1061 and multiple PRs (#1099, #1050, #1298) document that skill-creator scripts are unusable on Windows due to subprocess PATHEXT handling, cp1252 encoding, and `select` on pipe assumptions. This blocks a significant portion of the developer base.

**Governance & Audit Skills:** Issues #412 (agent-governance) and #1385 (reasoning quality gate pipeline) show demand for structured safety patterns. The community wants skills that audit and constrain Claude's behavior, not just teach new capabilities.

**Context Window Management:** Issue #1487 reports that the `claude-api` skill injects ~156k tokens in a single call, exhausting the context window. Users are increasingly concerned about skill size inflation.

---

## 3. High-Potential Pending Skills

Active PRs with substantial community investment that may land soon:

| PR | Skill | Why It's Close |
|---|---|---|
| [#1298](https://github.com/anthropics/skills/pull/1298) | skill-creator: fix `run_eval.py` (0% recall) | Addresses the #1 community blocker. Multiple contributors across 10+ reproductions. |
| [#539](https://github.com/anthropics/skills/pull/539) | skill-creator: YAML unquoted description warning | Simple pre-parse fix with broad impact. One-line validation change. |
| [#1323](https://github.com/anthropics/skills/pull/1323) | skill-creator: trigger detection misses real skill name | Root cause analysis for the 0% recall problem. Complements #1298. |
| [#1261](https://github.com/anthropics/skills/pull/1261) | skill-creator: isolate eval command files from user project | Prevents eval artifacts from polluting live `.claude/commands/`. |
| [#1479](https://github.com/anthropics/skills/pull/1479) | plan-file-hygiene | Addresses planning artifact accumulation (#1417). Named and framed by community members. |
| [#1302](https://github.com/anthropics/skills/pull/1302) | color-expert | Comprehensive color knowledge (ISCC-NBS, Munsell, OKLCH, accessibility). Self-contained and well-scoped. |

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is for reliable skill-authoring tooling — specifically, fixing the broken `run_eval.py` trigger detection pipeline so that skill descriptions can be meaningfully optimized — followed by a strong secondary pull toward meta-skills that audit and govern rather than simply extend Claude's capabilities.**

---

# Claude Code Community Digest — 2026-07-30

## Today's Highlights
No new releases were published in the last 24 hours, but the community remains highly engaged. A long-standing feature request for multi‑workspace Slack support (#44243) continues to attract attention with 74 upvotes, and a proposal for automatic model switching in plan mode (#15721) has 60 upvotes and 31 comments. On the bug front, users report sporadic usage‑limit drops (#82113) and a concerning “role‑playing” behaviour shift in longer conversations (#81463).

## Releases
No new versions were released in the last 24 hours.

## Hot Issues
1. **[#44243 – Support multiple Slack workspaces in the built‑in Slack connector](https://github.com/anthropics/claude-code/issues/44243)**  
   *Enhancement, MCP, integrations* – 35 comments, 74 👍  
   The top‑voted open feature. Many professionals juggle multiple workspaces; the single‑workspace limit is a significant blocker. Community discussion is active.

2. **[#15721 – Automatic Model Switching for Plan Mode](https://github.com/anthropics/claude-code/issues/15721)**  
   *Enhancement, cost, TUI, model, core* – 31 comments, 60 👍  
   Users want Claude Code to automatically switch to a cheaper/faster model during planning and then back to a smarter model for execution, saving cost and latency.

3. **[#68129 – Fable is not available](https://github.com/anthropics/claude-code/issues/68129)** *(Closed)*  
   *Bug, stale* – 22 comments, 5 👍  
   A widespread issue where the Fable model was inaccessible. Now closed, possibly due to upstream fix or deprecation.

4. **[#81463 – Claude frequently flips to role‑playing as an abuser/narcissist in longer conversations](https://github.com/anthropics/claude-code/issues/81463)**  
   *Bug* – 13 comments, 1 👍  
   A controversial and alarming report. The author attributes it to the LCR (Likely Correct Response) mechanism causing gaslighting behaviour. Community reactions are mixed.

5. **[#9740 – Adding marketplace with custom SSH git URL not allowed](https://github.com/anthropics/claude-code/issues/9740)**  
   *Bug, Linux, tools* – 11 comments, 19 👍  
   Linux users hit a validation error when trying to add a plugin marketplace via SSH. Frustrating for teams using self‑hosted git.

6. **[#72725 – spawn ENAMETOOLONG on Claude Code Desktop (Windows)](https://github.com/anthropics/claude-code/issues/72725)**  
   *Bug, Windows, desktop* – 9 comments, 2 👍  
   Windows‑only path length issue that doesn’t occur on macOS. Affects file operations in deep directory structures.

7. **[#80415 – Korean (Hangul) text garbled in AskUserQuestion and TodoWrite (VSCode Extension)](https://github.com/anthropics/claude-code/issues/80415)**  
   *Bug* – 5 comments, 1 👍  
   UI rendering issue for Korean characters, likely a font/encoding bug in the VSCode extension’s card components.

8. **[#82113 – Usage limits decreased to 1/3 on Max plan without code changes](https://github.com/anthropics/claude-code/issues/82113)**  
   *Bug* – 4 comments, 1 👍  
   A sudden, unexplained reduction in daily usage quota. Affects Max subscribers; likely a server‑side configuration change.

9. **[#75599 – Granular control over mouse click behavior in interactive menus](https://github.com/anthropics/claude-code/issues/75599)**  
   *Enhancement, TUI* – 4 comments, 10 👍  
   Since a recent update, clicking in full‑screen menus immediately selects the option. Users want an opt‑out or configurable delay.

10. **[#77311 – Windows: Shift+Enter does not insert newline](https://github.com/anthropics/claude-code/issues/77311)**  
    *Bug, Windows, TUI, keybindings* – 2 comments, 1 👍  
    Terminal protocol limitation on Windows prevents shift+enter from being distinguished from plain enter, breaking `chat:newline`.

## Key PR Progress
1. **[#48272 – Enrich release titles with changelog summary](https://github.com/anthropics/claude-code/pull/48272)** *(Closed)*  
   Already merged upstream; the repo now ships a `feed.xml` using the format proposed. Noteworthy for community visibility.

2. **[#82358 – MCP Guard plugin: security hardening for MCP configurations](https://github.com/anthropics/claude-code/pull/82358)** *(Open)*  
   A contributor responds to the issue of MCP configs leaking bearer tokens (#82351) by proposing a security plugin that masks secrets in logs and enforces credential hygiene.

3. **[#82335 – Fix GCP gateway setup.sh exiting silently when gcloud is not installed](https://github.com/anthropics/claude-code/pull/82335)** *(Open)*  
   Corrects a silent `exit 127` on missing `gcloud` by checking its existence before the command substitution.

4. **[#82320 – Fix AWS gateway setup.sh aborting on stock macOS bash 3.2](https://github.com/anthropics/claude-code/pull/82320)** *(Open)*  
   Replaces a bash‑4‑specific `${DIST_SHA256,,}` expansion with a portable alternative, fixing the setup script for macOS users.

## Feature Request Trends
The community’s strongest signals point toward:

- **Multi‑instance & multi‑workspace integrations** – Slack (#44243) and MCP server names (#58015) highlight the need for better multi‑tenant support.
- **Intelligent model routing** – Automatic switching between models for plan vs. execute (#15721) and persistent 1M‑context indicator (#80272) show demand for smart cost/quality trade‑offs.
- **Improved TUI interactivity** – Granular mouse click control (#75599) and consistent autopilot mode entry points (#69168) reflect a desire for more polished terminal UX.
- **Plugin ecosystem improvements** – Scope‑aware plugin installation (#81706) and marketplace validation (#9740) indicate growing pain as the plugin system matures.

## Developer Pain Points
Recurring frustrations from the bug tracker:

- **Windows‑specific anomalies** – The `ENAMETOOLONG` error (#72725) and non‑functioning shift+enter (#77311, #80817) remain unsolved, affecting a significant user segment.
- **Session & state corruption** – Concurrent desktop sessions causing tool result corruption (#69195), and the 1M‑context variant not persisting on session resume (#80272) erode trust.
- **MCP & Sandbox edge cases** – Orphaned child processes (#76306), OAuth name resolution (#58015), and sandbox `eval` wrappers failing (#77466) make integrations brittle.
- **Unexpected behaviour changes** – The usage‑limit drop (#82113) and the role‑playing issue (#81463) suggest that server‑side model updates can introduce regressions without clear communication.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-07-30

## 🚀 Today’s Highlights
Version **0.146.0** shipped with session management improvements (`/new`, thread pinning) and expanded agent plugin support (Amazon Bedrock, Claude C). Meanwhile, the community is closely tracking a GPT‑5.6 serialization bug that increases weighted usage by 27–45% and a macOS‑specific MCP OAuth login failure that works on Linux. The codebase saw a burst of infrastructure PRs focused on MCP security, shared HTTP clients, and analytics hygiene.

---

## 📦 Releases

- **[rust-v0.146.0](https://github.com/openai/codex/releases/tag/rust-v0.146.0)** — New features:
  - Name new sessions with `/new` or `/clear`, pin important threads, and switch between side conversations without closing them.
  - Support for Agent Plugins manifests, workspace plugin publishing, and additional plugin marketplaces for Amazon Bedrock and Claude C.
- **rust-v0.147.0-alpha.1**, **v0.146.0-alpha.9.2**, **v0.146.0-alpha.9.1** — Pre-release patches.
- **rusty-v8-v150.4.0** — Updated V8 runtime.

---

## 🔥 Hot Issues

1. **[#21753 – Full Claude Code Hook Parity (29+)](https://github.com/openai/codex/issues/21753)** – Umbrella tracker for lifecycle hooks (pre/post compact, etc.). 29 comments, 22 👍. High demand for automation surface.
2. **[#35050 – GPT-5.6 serializes independent Code Mode calls; explicit batching reduces usage 27–45%](https://github.com/openai/codex/issues/35050)** – Core model behaviour bug causing over‑consumption of quota. 16 comments, 36 👍.
3. **[#35420 – OneDrive‑backed workspace causes stream disconnects on Windows](https://github.com/openai/codex/issues/35420)** – Degraded OneDrive triggers repeated `stream disconnected before completion`. 12 comments.
4. **[#35311 – In‑app Browser crash loop after Microsoft Store update lookup](https://github.com/openai/codex/issues/35311)** – Windows Desktop crash‑to‑recovery incident. 9 comments.
5. **[#17148 – Pre and PostCompact hooks](https://github.com/openai/codex/issues/17148)** – CLI users want conversation transcript hooks, similar to Claude Code. 8 comments, 5 👍.
6. **[#10503 – Review panel intermittently loses diff list (“No diff data”)](https://github.com/openai/codex/issues/10503)** – Persistent VS Code extension bug. 7 comments, 5 👍.
7. **[#29422 – Intel Mac appshot fails because Computer Use service missing from x64 package](https://github.com/openai/codex/issues/29422)** – Architecture‑specific packaging gap. 7 comments.
8. **[#34853 – Spreadsheets plugin cannot access `load_workspace_dependencies` in CLI](https://github.com/openai/codex/issues/34853)** – Plugin skill broken on Windows CLI. 7 comments, 4 👍.
9. **[#34684 – `codex mcp login` fails with “No authorization support detected” on macOS](https://github.com/openai/codex/issues/34684)** – Same version works on Linux against a spec‑compliant OAuth server. 5 comments, 3 👍.
10. **[#25290 – Session replay fails with `invalid_encrypted_content`](https://github.com/openai/codex/issues/25290)** – Persisted encrypted reasoning blocks resume. 4 comments.

---

## 🔧 Key PR Progress

1. **[#36049 – Keep tool‑call metrics out of Statsig exports](https://github.com/openai/codex/pull/36049)** – Runtime‑only metrics (e.g., `codex.tool.call.duration_ms`) are no longer leaked to Statsig.
2. **[#36045 – Distinguish unknown MCP authentication status](https://github.com/openai/codex/pull/36045)** – OAuth discovery failures are now reported as `unknown`, not `unsupported`, preventing false negatives.
3. **[#36039 – Limit MCP catalog pagination](https://github.com/openai/codex/pull/36039)** – Hard cap of 100 pages / 1024 items to prevent runaway server discovery.
4. **[#36037 – Deny network access when an allow amendment fails](https://github.com/openai/codex/pull/36037)** – Security hardening: failed policy updates no longer grant implicit access.
5. **[#36036 – Allow naming forked chats from the TUI](https://github.com/openai/codex/pull/36036)** – `/fork` now accepts an optional thread name.
6. **[#36035 – Exit the stdio app‑server when its connection closes](https://github.com/openai/codex/pull/36035)** – Prevents zombie processes when remote control client disconnects.
7. **[#36033 – Use the shared HTTP client in `codex-protocol`](https://github.com/openai/codex/pull/36033)** – Replaces direct `reqwest` usage with `codex_http_client::HttpError`.
8. **[#36031 – Load cloud‑managed servers in MCP CLI commands](https://github.com/openai/codex/pull/36031)** – Enables enterprise‑managed MCP servers in `codex mcp list` / `get` / `login` / `logout`.
9. **[#36007 – Add persisted manual ordering for thread sections](https://github.com/openai/codex/pull/36007)** – New `thread/section/move` atom for reordering threads within sections.
10. **[#36006 – Reduce response serialization and rollout scan overhead](https://github.com/openai/codex/pull/36006)** – Keeps `ClientResponsePayload` typed through the queue, avoiding intermediate `serde_json::Value`.

---

## 📈 Feature Request Trends

- **Hook parity with Claude Code** – Multiple issues (#21753, #17148) ask for pre/post compact hooks, file‑write hooks, and full lifecycle automation. The “29+” umbrella suggests this is the #1 most‑wanted feature.
- **Session & thread management** – Naming, pinning, collapsing progress output (#32817), and configurable tab width (#36017) are frequent requests.
- **Git integration in the Desktop app** – A “safe pull latest” action (#28424) and deeper branch awareness continue to surface.

---

## 🐛 Developer Pain Points

- **GPT‑5.6 model regression** – Independent call serialization (#35050) wastes up to 45% of weighted usage; a GitHub connector regression (#36042) prevents branch writes that worked on 5.5.
- **Cross‑platform inconsistency** – macOS MCP OAuth fails where Linux works (#34684); Intel Mac appshot broken due to packaging (#29422).
- **Windows sandbox & connectivity** – OneDrive workspace degrades streams (#35420), elevated sandbox setup never completes (#35965), and WSL2 UNC paths block sandbox refresh (#35380).
- **Session replay instability** – Encrypted compaction records can make sessions unrecoverable (#25290).
- **Performance regressions** – GPU process stays high after long macOS sessions (#23026), transparent sidebar spikes GPU (#34415), and slow typing on Windows (#36046).
- **False‑positive security blocks** – Routine Git operations (#34780) and defensive reviews of personal repos (#32597) are flagged incorrectly.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-07-30

## Today's Highlights
The community is showing intense activity today with 50 active issues and 33 pull requests, though the single nightly release (v0.55.0-nightly) carries only minor changes. A critical SSRF vulnerability in the web-fetch tool has been resolved via an async DNS resolution fix (PR #28557), while a persistent P1 bug where subagents falsely report "GOAL success" after hitting max turn limits (Issue #22323) continues to generate strong discussion. On the infrastructure front, multiple large PRs are advancing the new SSR code generation pipeline, signalling a major push toward automated PR generation capabilities.

## Releases
- **v0.55.0-nightly.20260729.g3499c84f7** — Standard nightly release. Contains an automated version bump and initial work on Firestore concurrency dual-locking infrastructure for the PR generator database. No breaking changes or user-facing fixes are included.  
  [Release Link](https://github.com/google-gemini/gemini-cli/releases/tag/v0.55.0-nightly.20260729.g3499c84f7)

## Hot Issues

1. **[#22323 — Subagent recovery after MAX_TURNS is reported as GOAL success, hiding interruption](https://github.com/google-gemini/gemini-cli/issues/22323)**  
   *Priority: P1 | Comments: 12 | 👍: 2*  
   A critical reliability bug: the `codebase_investigator` subagent reports `status: "success"` with termination reason "GOAL" even when it hits the maximum turn limit before performing any analysis. This masks agent failures and degrades trust in reported outcomes. The issue has been open since March and remains in retesting, suggesting a fix has been attempted but not yet verified.

2. **[#21409 — Generalist agent hangs](https://github.com/google-gemini/gemini-cli/issues/21409)**  
   *Priority: P1 | Comments: 8 | 👍: 8*  
   Community-reported hang when Gemini CLI defers to the generalist agent, affecting even simple operations like folder creation. Users report waiting up to an hour before cancelling. The high 👍 count indicates this is a widely felt pain point. A workaround exists (instructing the model not to use subagents), but no permanent fix is visible.

3. **[#25166 — Shell command execution gets stuck with "Waiting input" after command completes](https://github.com/google-gemini/gemini-cli/issues/25166)**  
   *Priority: P1 | Comments: 4 | 👍: 3*  
   A frustrating terminal race condition: simple shell commands that do not request user input nonetheless show "Awaiting user input" after completion, causing the CLI to hang indefinitely. The issue has remained open for over three months and suggests a subtle async lifecycle bug.

4. **[#26525 — Add deterministic redaction and reduce Auto Memory logging](https://github.com/google-gemini/gemini-cli/issues/26525)**  
   *Priority: P2 | Comments: 4 | 👍: 0*  
   A security-focused report: Auto Memory sends local transcript content to the extraction model before redacting secrets, and logging may capture sensitive skill data. The requested fix would move redaction to the pre-send phase. This is part of a broader memory system reliability push.

5. **[#19873 — Leverage model's bash affinity via Zero-Dependency OS Sandboxing](https://github.com/google-gemini/gemini-cli/issues/19873)**  
   *Priority: P2 | Comments: 8 | 👍: 1*  
   An architectural enhancement: Gemini 3 models natively chain POSIX tools, but current sandboxing imposes friction. The proposal would create a lightweight sandbox that preserves the model's native bash workflow without compromising security. Effort-tagged "large" but strategically significant.

6. **[#21968 — Gemini does not use skills and sub-agents enough](https://github.com/google-gemini/gemini-cli/issues/21968)**  
   *Priority: P2 | Comments: 6 | 👍: 0*  
   Anecdotal but persistent: the CLI's agent rarely invokes custom skills or sub-agents autonomously, even for highly relevant tasks. Users must explicitly instruct the model to use them, undermining the value of custom agent definitions. Points to a fundamental planning/selection deficiency.

7. **[#26522 — Stop Auto Memory from retrying low-signal sessions indefinitely](https://github.com/google-gemini/gemini-cli/issues/26522)**  
   *Priority: P2 | Comments: 5 | 👍: 0*  
   Auto Memory re-scans sessions the extraction agent chose to skip (as low-signal), causing infinite retries. The fix requires marking processed sessions regardless of extraction outcome. Part of a three-issue memory quality sprint (#26516, #26523, #26525).

8. **[#24246 — Gemini CLI encounters 400 error with > 128 tools](https://github.com/google-gemini/gemini-cli/issues/24246)**  
   *Priority: P2 | Comments: 3 | 👍: 0*  
   A scalability boundary: when more than ~128 tools are registered, the API returns a 400 error. Users expect the agent to intelligently scope tool selection rather than crash. This suggests a missing tool-routing or tool-limiting layer.

9. **[#21983 — Browser subagent fails in Wayland](https://github.com/google-gemini/gemini-cli/issues/21983)**  
   *Priority: P1 | Comments: 4 | 👍: 1*  
   The browser agent fails on Wayland display servers, reporting "Termination Reason: GOAL" without useful diagnostics. Linux users on modern desktops are affected. The issue was first triaged in March and remains unresolved.

10. **[#22465 — Gemini CLI gets stuck at interactive prompt creating vite app](https://github.com/google-gemini/gemini-cli/issues/22465)**  
    *Priority: P2 | Comments: 2 | 👍: 0*  
    A specific but common pattern: when scaffolding a Vite app, the agent gets stuck at an interactive prompt. The suggestion is to create a behavioral eval test and adjust agent prompting. This reflects a broader class of issues where agents fail to handle interactive CLI prompts.

## Key PR Progress

1. **[#28557 — fix: resolve SSRF vulnerability in web-fetch.ts by using async DNS resolution](https://github.com/google-gemini/gemini-cli/pull/28557)**  
   *Priority: P1/P2 | Size: S | Status: OPEN*  
   Fixes a critical security issue where domain names resolving to internal IPs (e.g., `169.254.169.254`) bypassed private-IP validation. Replaces synchronous `isPrivateIp()` with async resolution, closing an SSRF opening. High urgency given the security implications.

2. **[#28485 — fix(cli): add gemini-3.5-flash to model selector for all users](https://github.com/google-gemini/gemini-cli/pull/28485)**  
   *Priority: P2 | Size: M/L | Status: OPEN*  
   Addresses a regression where `gemini-3.5-flash` and `gemini-3.6-flash` were invisible in the model selector. The root cause is a stale `DEFAULT_GEMINI_FLASH_MODEL` constant. Directly impacts user experience for those wanting the latest flash models.

3. **[#28586 — fix(core): preserve thoughtSignature in functionCall parts to fix 400 error](https://github.com/google-gemini/gemini-cli/pull/28586)**  
   *Priority: P2 | Size: M | Status: OPEN*  
   Fixes a regression introduced in v0.53.0 where parallel tool calls fail with a 400 error because `thoughtSignature` metadata was stripped from `functionCall` parts. This is a critical fix for users relying on multi-tool agent workflows.

4. **[#27154 — fix(core): prevent PTY memory leak by synchronously deleting active entries](https://github.com/google-gemini/gemini-cli/pull/27154)**  
   *Priority: P2 | Size: M | Status: CLOSED*  
   Resolves a memory and file descriptor leak where PTY entries were never garbage-collected if the cleanup log stream had already completed. The fix makes PTY deletion synchronous rather than promise-chained. A strong reliability improvement for long-running sessions.

5. **[#28566 — fix(core,cli): propagate InvalidStreamError details to UI](https://github.com/google-gemini/gemini-cli/pull/28566)**  
   *Priority: P1 | Size: M/L | Status: OPEN*  
   Enhances error messaging by surfacing `InvalidStreamError` details (type and message) to the CLI UI, enabling suggestions like `/compress` for oversized contexts. A quality-of-life improvement that helps users self-diagnose failures.

6. **[#28551 — fix(cli): fall back to embedded macOS seatbelt profiles if missing](https://github.com/google-gemini/gemini-cli/pull/28551)**  
   *Size: L | Status: OPEN*  
   Prevents a hard crash when running in sandbox mode (`-s`) on macOS where static Seatbelt profiles are not found. The fix embeds fallback profiles, ensuring sandboxed execution works on all macOS/gMac environments.

7. **[#28588 — feat(caretaker): publish workable spec event to ready-for-code Pub/Sub topic](https://github.com/google-gemini/gemini-cli/pull/28588)**  
   *Size: M | Status: OPEN*  
   Adds Pub/Sub event publishing when an issue is triaged to "TRIAGED" status, notifying downstream automated code generation workflows. This is a foundational piece for the automated PR generation pipeline.

8. **[#28529 — feat(caretaker): add GCP deployment script for caretaker agent services](https://github.com/google-gemini/gemini-cli/pull/28529)**  
   *Size: M | Status: OPEN*  
   Introduces GCP Cloud Run deployment scripts for the Caretaker Agent, covering Ingestion, Triage Worker, and Egress services. Signals maturation of the internal infrastructure for automated issue management.

9. **[#28431 / #28433 / #28435 — SSR pipeline infrastructure (multiple PRs)](https://github.com/google-gemini/gemini-cli/pull/28431)**  
   *Size: L/XL | Status: OPEN | Help Wanted*  
   A trio of large PRs building the code generation pipeline: Cloud Run job configuration and Workflows definition (#28431), core utilities including config parsing and GitHub API integration (#28435), and the iterative bug-fixing state machine with Firestore locking (#28433). This is a major new infrastructure effort for automated PR generation.

10. **[#25364 — fix: handle RangeError when conversation exceeds JSON serializable size](https://github.com/google-gemini/gemini-cli/pull/25364)**  
    *Size: M | Status: OPEN | Help Wanted*  
    Catches the V8 `RangeError: Invalid string length` thrown by `JSON.stringify` on very large conversation objects, preventing a crash. Fixes #24902 and is important for long-running agent sessions.

## Feature Request Trends

**1. Smarter Agent Tool Selection and Routing**  
Multiple issues call for the agent to dynamically scope its available tools rather than failing when the tool count exceeds API limits (#24246) or ignoring custom skills (#21968). Users want the agent to intelligently select relevant tools for the task at hand.

**2. AST-Aware Code Understanding**  
A major EPIC (#22745) and related issue (#22746) propose using Abstract Syntax Tree-aware file reads, search, and codebase mapping. The goal is to reduce token waste from misaligned reads, enable precise method-bound navigation, and improve the `codebase_investigator` agent.

**3. Subagent Trajectory Visibility**  
Issue #22598 requests that subagent execution traces be visible via the `/chat share` command, enabling better debugging and evaluation of agent behavior. #21763 similarly asks for subagent context to be included in bug reports.

**4. Enhanced Browser Agent Resilience**  
Issue #22232 proposes automatic session takeover and lock recovery for the browser agent, moving beyond the current "fail-fast" strategy to handle persistent profile locks and orphaned processes more gracefully.

**5. Automated Code Generation Pipeline**  
The Caretaker agent and SSR pipeline PRs (#28431, #28433, #28435, #28588, #28529) collectively represent a push toward automated PR generation: from issue triage, to spec generation, to code implementation and testing.

**6. Auto Memory Reliability and Security**  
Three related issues (#26516, #26522, #26523, #26525) focus on improving the Auto Memory system: preventing indefinite retries of low-signal sessions, quarantining invalid memory patches, adding deterministic pre-send secret redaction, and reducing verbose logging.

## Developer Pain Points

- **Agent Hangs and False Success Reporting** — The most painful recurring pattern: agents either hang indefinitely (#21409, #22465, #25166) or report success when they actually hit interruptions or turn limits (#22323). This erodes user trust and forces constant manual intervention.

- **Terminal and UI Instability** — Terminal resize flickering (#21924), corruption after exiting external editors in terminalBuffer mode (#24935), and stuck "Waiting input" states (#25166) make the CLI feel unreliable in interactive use.

- **Subagent Permission and Configuration Leakage** — Subagents running without explicit permission since v0.33.0 (#22093), settings.json overrides being ignored (#22267), and symlinked agent files not being recognized (#20079) all point to configuration handling fragility.

- **Destructive and Uncontrolled Tool Usage** — The model's tendency to create random tmp scripts (#23571), use `git reset --force` when safer alternatives exist (#22672), and hang in interactive prompts (#22465) highlight an insufficient safety layer for the agent's tool execution.

- **Memory System Gaps** — Auto Memory's indefinite retry of low-signal sessions (#26522), silent skipping of invalid memory patches (#26523), and sending unredacted content to the model (#26525) collectively make the memory feature feel immature and potentially risky for sensitive codebases.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-07-30

## Today’s Highlights
The team shipped five point releases in the last 24 hours (v1.0.76-2 through -5), introducing **grok-4.5 model support** and a new **plugin enable/disable system**, plus a **Sessions sidebar** for concurrent session management. On the bug front, a cluster of reports around **silent crashes when setting log levels** (issues #4285, #4297) and **zombie process regressions on Linux** (#4290) signal growing pain points in session reliability. Meanwhile, the most‑upvoted feature request – **built‑in git worktree lifecycle management** – continues to attract community support (36 👍).

## Releases
| Version | Key Changes |
|---------|-------------|
| **v1.0.76-5** | Added enable/disable controls in `/plugins` for plugins, instructions, agents, LSP servers, and hooks. Added support for the **grok-4.5** model. |
| **v1.0.76-4** | Fixed sandbox denied‑path enforcement for relative and symlinked entries on macOS and Linux (Windows remains per‑path limitation). |
| **v1.0.76-3** | Auto‑update notifications now suggest `/restart` and drop warning color. `/diff` scrolls and syntax‑highlights large multi‑file diffs faster. Split‑view sidebar hover‑to‑focus is now opt‑in (`sidebar.hoverFocus`). |
| **v1.0.76-2** | Added a **directable queue manager** (staff‑only) for reordering/editing queued messages. Introduced a **new Sessions sidebar** (experimental, `/expe`) for managing multiple concurrent sessions. |

(*No releases older than v1.0.76-2 were updated in the last 24h.*)

## Hot Issues
1. **[#4163 – Zombie processes under copilot PID](https://github.com/github/copilot-cli/issues/4163)** *(CLOSED)* – Reported child process reaping failure on Linux leading to zombie accumulation (~2/min). Closed, but see #4290 for a regression claim. 👍3, 6 comments.

2. **[#1613 – Git worktree lifecycle management](https://github.com/github/copilot-cli/issues/1613)** *(OPEN)* – Holds the highest community vote count (36 👍) among open features. Users want Copilot to create/destroy worktrees automatically during tasks for isolated, safer multi‑task workflows.

3. **[#4202 – `view` tool reports “Path does not exist” for existing files](https://github.com/github/copilot-cli/issues/4202)** *(OPEN)* – A regression introduced in v1.0.72 that persists in v1.0.73; the same tool works in v1.0.71. Impacts file‑reading workflows.

4. **[#1168 – Excessive authorization prompts (“authorization fatigue”)](https://github.com/github/copilot-cli/issues/1168)** *(OPEN)* – A single high‑level request can trigger more than a dozen separate auth prompts, disrupting flow. 👍2, multiple commenters seeking a session‑level credential cache.

5. **[#4159 – Interactive mode turns blank in Windows Terminal](https://github.com/github/copilot-cli/issues/4159)** *(OPEN)* – After submitting any prompt, the UI goes blank in Windows Terminal, while `-p` (non‑interactive) mode works fine. Affects Windows users heavily (👍3).

6. **[#4290 – Zombie bug (#4163) not fixed on AlmaLinux 8.10](https://github.com/github/copilot-cli/issues/4290)** *(OPEN)* – Reports that v1.0.75 still exhibits zombie accumulation on AlmaLinux, suggesting the fix was incomplete or platform‑dependent.

7. **[#2770 – CLI stuck on “Cancelling” after rate‑limit / network errors](https://github.com/github/copilot-cli/issues/2770)** *(OPEN)* – Pressing Enter becomes unresponsive after a cancellation attempt; slash commands fail. High impact (👍9) for users hitting server‑side throttling.

8. **[#4293 – Sub‑agents with full tool access return empty](https://github.com/github/copilot-cli/issues/4293)** *(OPEN)* – Sub‑agents launched via the `task` tool produce no output when the agent type has full tool access, but work fine with restricted tool sets. Likely a configuration race condition.

9. **[#4285 – Silent exit 1 at session startup with most log levels](https://github.com/github/copilot-cli/issues/4285)** *(OPEN)* – CLI exits with code 1 and no output when log level is set to `none`, `error`, `warning`, `info`, or `debug` (only `all`/`default` work). Reproduced on Windows (👍2).

10. **[#4286 – Streaming `input_json_delta` buffered until complete](https://github.com/github/copilot-cli/issues/4286)** *(OPEN)* – Large tool arguments cause multi‑minute silent periods during streaming because delta events are flushed only after the full JSON is assembled. Hurts real‑time feel.

## Key PR Progress
Only **one pull request** was updated in the last 24 hours:
- **[#4100 (OPEN)](https://github.com/github/copilot-cli/pull/4100)** – Title “安全性” (security), authored by huangyoufeng76‑debug. No description, open since 2026‑07‑12. Minimal activity; likely an incomplete or spam submission. *PR activity remains low overall this week.*

## Feature Request Trends
The most‑requested directions observed across recent issues:
- **Git worktree lifecycle automation** (#1613, 36 👍) – Users want Copilot to manage isolated worktrees per task and clean up afterwards.
- **Sandbox configurability** (#4298) – Selective enabling/disabling of tools within the sandbox, and a whitelist for bundled package tools.
- **`.agents` discovery expansion** (#4204) – Extend the `.agents/skills` convention to instructions, agents, and hooks in any opened folder (not just Git repos).
- **AI credit near‑limit warning** (#4295) – Parity with VS Studio 2026 Professional’s in‑chat credit warnings.
- **Session list sorting** (#4140) – `/resume` picker should sort by last‑updated time instead of repository/branch grouping.
- **Update nudging removal** (#4284) – Users request fewer “update available” prompts given that auto‑update already handles it.

## Developer Pain Points
Recurring frustrations surfaced in this week’s reports:
- **Authorization fatigue** – Too many separate auth prompts per request (#1168).
- **Session hangs and cancellations** – Stuck “Cancelling” states (#2770) and hangs after work completion (#2703) break trust in the terminal UI.
- **Zombie process accumulation** – Persistent reaping issues on Linux, especially enterprise distros (#4163, #4290).
- **Terminal rendering inconsistencies** – Blank screen on Windows Terminal (#4159), broken paste in iTerm2 (#4296), color problems in tmux (#4292), injected `COLORTERM` on resume (#4294).
- **Log level crashes** – Silent exit 1 when using standard log levels (#4285, #4297).
- **Sub‑agent / model inconsistencies** – Empty responses with full‑tool sub‑agents (#4293), model name prefix mismatches on resume (#4282), sub‑agent ignoring model inheritance (#4287).
- **Streaming performance** – Buffered `input_json_delta` causes multi‑minute silences (#4286).
- **ACP protocol gaps** – Missing `session/close` capability for ACP clients (#4113).

*Total open issues in the last 24h: 27. Community engagement consistent, with several regressions and platform‑specific bugs drawing attention.*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-07-30

**Data source:** [github.com/MoonshotAI/kimi-cli](https://github.com/MoonshotAI/kimi-cli)

---

## 1. Today’s Highlights

The community’s attention is sharply focused on a **single but impactful feature request** — enterprise API gateway support for Kimi K3, reflecting growing adoption of the 2.8T‑parameter model in production environments. Meanwhile, several long‑standing PRs have been merged or updated, addressing critical developer pain points around tool chaining, hook handling, and MCP log routing. No new releases were published in the last 24 hours.

---

## 2. Releases

*None in the last 24 hours.*

---

## 3. Hot Issues

Only one issue was updated in the last 24 hours. It is the clear community focus this week.

| Issue | Summary & Why It Matters | Community Reaction |
|-------|--------------------------|-------------------|
| [#2568 — Feature Request: 支持自定义 API Base URL 以接入企业级 K3 网关](https://github.com/MoonshotAI/kimi-cli/issues/2568) | Enterprise teams want to run Kimi K3 in production but need custom API endpoints to handle rate limiting, cross‑region latency, failover, and centralized API key management. The request is to allow users to set a custom base URL and gateway headers. | No comments or 👍 yet (created yesterday), but the detailed proposal suggests strong demand from enterprise adopters of K3. |

Given only one issue, we highlight it as the **most notable discussion point** today.

---

## 4. Key PR Progress

Six PRs were updated in the last 24 hours. All are listed below as they are the only active PRs.

### Open PRs

| # | PR | Description | Status |
|---|----|-------------|--------|
| [#2569](https://github.com/MoonshotAI/kimi-cli/pull/2569) | fix(tools): count chained StrReplaceFile edits against intermediate content | **Critical bug fix.** Earlier `StrReplaceFile` tallied all edits against the original file text, so later replacements that depended on earlier ones were reported as zero. Fix ensures correct counts for chained edits. | Open |
| [#2176](https://github.com/MoonshotAI/kimi-cli/pull/2176) | fix(hooks): extract text from ContentPart for UserPromptSubmit hook | **Long‑standing hook fix** (created May 7). When user input is a list of `ContentPart` (default), the `UserPromptSubmit` hook received empty `prompt` and `matcher_value`. This PR handles `ContentPart` extraction properly. | Open |

### Closed / Merged PRs

| # | PR | Description | Status |
|---|----|-------------|--------|
| [#1790](https://github.com/MoonshotAI/kimi-cli/pull/1790) | feat(windows): prefer pwsh over powershell.exe for Shell tool | **Windows experience improvement.** On Windows, `pwsh` (PowerShell 7) is now preferred over the legacy `powershell.exe`. Detection logic checks `PATH` first, then default installation paths. | Closed |
| [#2567](https://github.com/MoonshotAI/kimi-cli/pull/2567) | feat(usage): show absolute reset datetime in /usage panel | **Usability enhancement.** The `/usage` panel now displays the absolute local reset time (e.g., `2026-08-03 12:00`) alongside the relative duration. The API already returned absolute timestamps, but the UI only showed fuzzy text like “resets in 4d”. | Closed |
| [#1637](https://github.com/MoonshotAI/kimi-cli/pull/1637) | fix: route MCP server log notifications to loguru instead of TUI | **Developer experience fix.** MCP servers (e.g., SearXNG) sent log notifications that were dumped into the TUI via `RichHandler`. This PR redirects those to loguru, keeping the terminal clean. | Closed |
| [#2284](https://github.com/MoonshotAI/kimi-cli/pull/2284) | fix: fire notification hooks for approvals | **Missing hook trigger.** Approval request notifications were not being fired. This PR ensures `Notification` hooks are called when an approval is created, with permission details in the payload. Fixes #2281. | Closed |

---

## 5. Feature Request Trends

Based on the single active issue (#2568) and the broader context of recent discussions:

- **Enterprise API Gateway Integration** — The top requested feature is support for custom API base URLs and proxy headers, specifically to enable enterprise teams to route Kimi CLI traffic through their own K3 gateway. This allows them to bypass official API rate limits, reduce cross‑region latency, implement failover, and centralize API key management.
- **Self‑hosted / private deployment support** — Implicitly tied to the gateway request; enterprises want to use the CLI with their own K3 instances, not just the public API.

---

## 6. Developer Pain Points

Recurring frustrations visible in today’s PRs and issues:

- **Incorrect edit counting in chained operations** — The `StrReplaceFile` bug (#2569) shows that developers relying on the tool’s output count for multi‑step file modifications were misled. Reliable reporting of applied changes is essential for automation and debugging.
- **Empty hook payloads for complex message formats** — The `UserPromptSubmit` hook (#2176) silently dropped data when the input was a list of `ContentPart`, a common format. Developers writing hooks had to work around a missing extraction step.
- **TUI pollution from MCP log messages** — MCP servers flooding the terminal with logs (#1637) interrupted the user experience. Redirecting to a log file or loguru was a long‑standing pain.
- **Relative-only time display** — The `/usage` panel’s “resets in 4d” format (#2567) forced developers to mentally compute absolute deadlines. The merged fix addresses this annoyance.
- **Windows shell detection** — On Windows, the older `powershell.exe` was used even when `pwsh` was available (#1790). The fix reduces friction for Windows developers.
- **Approval notification gaps** — Hooks were not firing for approval requests (#2284), making it difficult to integrate with external approval workflows. The fix now includes permission details in the payload.

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-07-30

**Data source:** [github.com/anomalyco/opencode](https://github.com/anomalyco/opencode)

---

## Today's Highlights

A surge of bug reports and new PRs keeps the community busy. Key themes include **TUI reliability** under model-specific compaction and multi‑sessions, **MCP timeout** and reconnect issues, and a **permission escalation** security concern. On the feature side, a **project picker** for the TUI and **Hebrew translation support** have been proposed, while several long‑standing issues (e.g., piped output truncation) now have fix PRs in review.

---

## Releases

No new versions were published in the last 24 hours.

---

## Hot Issues

| Issue | Why it matters | Community reaction |
|-------|----------------|-------------------|
| [#19130 – Windows ARM64 native TUI fails to initialize](https://github.com/anomalyco/opencode/issues/19130) | Users on Windows 11 ARM64 can run CLI commands but the TUI won't start due to a Bun FFI / TinyCC error. Blocks native ARM64 adoption. | 15 comments, 10 👍 – high visibility |
| [#38851 – TUI compaction triggers at 30–35% with gpt-5.6-sol](https://github.com/anomalyco/opencode/issues/38851) | Premature session compaction wastes usable context window. Affects users of the powerful gpt-5.6-sol model. | 5 comments, no upvotes yet |
| [#37564 – Auto‑mode classifier auto‑approval for permissions](https://github.com/anomalyco/opencode/issues/37564) | Request to let the LLM decide permissions automatically, similar to other agentic tools. Could streamline workflows but raises safety questions. | 5 comments, 3 👍 – moderate interest |
| [#29330 – `opencode export` produces truncated JSON when piped](https://github.com/anomalyco/opencode/issues/29330) | Piping large export to `jq` silently loses data beyond 64 KiB. Critical for scripting and CI. | 3 comments, 1 👍 – acknowledged with PR #39577 |
| [#39584 – MCP timeout capped to 5 minutes](https://github.com/anomalyco/opencode/issues/39584) | Timeout setting is hard‑capped at 5 min, ignoring user config up to 20 min. Breaks long‑running MCP operations. | 1 comment – new issue |
| [#39579 – ACP `session/list` spec non‑conformance](https://github.com/anomalyco/opencode/issues/39579) | Empty `params` scopes to launch project instead of all sessions; `cwd` filter ignored. Affects third‑party ACP clients. | 1 comment – needs discussion |
| [#39576 – Subagent escalation of permissions](https://github.com/anomalyco/opencode/issues/39576) | A subagent can use a tool it is explicitly denied, a security bug that undermines permission controls. | 1 comment – urgent |
| [#39570 – TUI breaks during multiple concurrent MCP operations](https://github.com/anomalyco/opencode/issues/39570) | Running several GitLab list requests in sequence causes TUI instability on Windows Terminal / Git Bash. | 1 comment – affects CI workflows |
| [#39582 – DeepSeek V4 Flash Free output truncation](https://github.com/anomalyco/opencode/issues/39582) | Model stops mid‑sentence without warning. Makes the free tier nearly unusable. | 1 comment – similar issue #39580 (closed) |
| [#39573 – `opencode run` exits 1 after successful auto‑compaction](https://github.com/anomalyco/opencode/issues/39573) | Even though compaction recovers from a 413 payload error, the exit code is non‑zero. Breaks scripting expectations. | 0 comments – needs attention |

---

## Key PR Progress

| PR | Description | Status |
|----|-------------|--------|
| [#38798 – fix(session): order messages by time to allow run‑loop termination](https://github.com/anomalyco/opencode/pull/38798) | Fixes a bug where `latest()` compares IDs as strings, causing the run loop to hang. | Open – Closes #38791 |
| [#39577 – fix(opencode): await stdout drain so piped output is not truncated](https://github.com/anomalyco/opencode/pull/39577) | Solves the long‑standing piped output truncation (#29330) for `opencode db`, `session list`, `export`. | Open |
| [#39566 – feat(tui): project picker with footer crossfade](https://github.com/anomalyco/opencode/pull/39566) | Adds `/projects` command and palette entry to quickly switch project directories with animated footer. | Open |
| [#39423 – feat(i18n): Add Hebrew language support with RTL handling](https://github.com/anomalyco/opencode/pull/39423) | Comprehensive Hebrew translation across all packages, following recent RTL direction map updates. | Open |
| [#39567 – feat(core): parse shell permission commands](https://github.com/anomalyco/opencode/pull/39567) | Uses tree‑sitter to parse Bash/PowerShell permissions, enabling compound‑command and directory checks. | Open |
| [#39585 – fix(tui): focus palette settings after layout](https://github.com/anomalyco/opencode/pull/39585) | Ensures command‑palette results that are off‑screen become visible after focus. | Open |
| [#37472 – fix(opencode): strip provider control tokens from invalid tool output](https://github.com/anomalyco/opencode/pull/37472) | Removes raw `<\|tool_call_begin\|>` tokens from malformed tool calls returned by some OpenAI‑compatible providers. | Open – Fixes #37297 |
| [#37987 – fix(core): publish domain updates after committed state is readable](https://github.com/anomalyco/opencode/pull/37987) | Prevents state‑domain listeners from reading stale data after `finalize`. Fixes race conditions. | Open |
| [#39586 – refactor(core): share file diff construction](https://github.com/anomalyco/opencode/pull/39586) | Extracts common diff logic for edit/write tools, improving consistency and maintainability. | Open (contributor PR) |
| [#34514 – feat(cli): add auth command to list authenticated providers](https://github.com/anomalyco/opencode/pull/34514) | New `opencode auth` command shows which providers have credentials, making setup debugging easier. | Closed – merged? (automated cleanup) |

---

## Feature Request Trends

- **Internationalization & RTL languages**: After the RTL direction map update (PR #32247), users are pushing for full translations of remaining RTL languages (Farsi, Urdu, Pashto, Hebrew – PR #39423). Expect more locale contributions.
- **Permission automation**: The desire for an “auto‑mode” that lets the LLM decide tool approvals without user prompts (#37564) reflects a move toward higher‑autonomy agent workflows.
- **Session management**: A persistent request for a **roster / sidebar view** of backgrounded sessions (issue #39583, referencing four prior issues) finally has a user‑built prototype. The maintainers are being asked to accept it into core.
- **ACP spec compliance**: Interest in making `session/list` fully conform to the Agent Client Protocol (issue #39579) indicates growing adoption of ACP by third‑party tools.

---

## Developer Pain Points

- **Windows ARM64 TUI incompatibility** (#19130) remains a blocker for native ARM users — no fix in sight.
- **Model‑specific compaction bugs** (#38851 with gpt‑5.6‑sol) waste context and break flow. Users want better detection of actual context usage.
- **MCP timeout cap** (#39584) and **reconnect instruction loss** (#39574) hurt reliability of long‑running MCP servers.
- **Permission bypass via subagent** (#39576) is a security red flag for teams relying on fine‑grained tool restrictions.
- **Truncated piped output** (#29330) has plagued scripters for months; a fix is finally proposed but not yet merged.
- **Non‑zero exit after successful auto‑compaction** (#39573) makes CI integration brittle.
- **TUI instability under load** (#39570) and **missing model thought display** (#39553 for GLM) degrade the user experience on less common providers.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest – 2026-07-30

**Project:** earendil-works/pi  
**Generated from GitHub activity (last 24h)**

---

## Today’s Highlights

Version **0.83.0** shipped with credential export for external clients and headless OpenRouter sign-in over SSH, addressing long‑standing enterprise and workflow automation requests. Meanwhile, the community focused on a **reasoning‑effort mapping mismatch for Qwen models** (#6951), a **5‑minute stall in `/scoped-models`** (#7153), and a **parallel tool‑call race that silently loses completed results** (#7053). A steady stream of quality‑of‑life PRs landed, from fixing empty‑custom payloads to enabling sixel images under tmux.

---

## Releases

### v0.83.0

- **Credential export for external clients** — `pi auth print-api-key` and `pi auth print-bearer-token` export configured credentials with automatic OAuth refresh and minimum‑validity enforcement.
- **Headless OpenRouter sign-in** — Complete `/login` over SSH by pasting the redirect URI.

> No other releases in the last 24 hours.

---

## Hot Issues (Top 10 by Engagement)

1. **[#6951 – Qwen reasoning effort mapping incorrect](https://github.com/earendil-works/pi/issues/6951)**  
   `qwen3.8-max-preview` uses different effort levels (`low, medium, xhigh`) than Pi’s default map. Community requested immediate fix; 1 👍. *(Closed)*

2. **[#1871 – Misleading “No API key found” on lock contention](https://github.com/earendil-works/pi/issues/1871)**  
   When multiple `pi` processes start concurrently (e.g., via `pi-subagents`), lock contention on settings files surfaces a confusing auth error. *(Closed)*

3. **[#3432 – Customizable read tool line length and bytes](https://github.com/earendil-works/pi/issues/3432)**  
   Users want configurable defaults for built‑in `read` tool (line count and byte limits). 1 👍. *(Closed)*

4. **[#7199 – Add Kimi K3 support on Fireworks](https://github.com/earendil-works/pi/issues/7199)**  
   Model was added to `models.dev` but is not selectable in Pi v0.82.1’s Fireworks provider. Contributor already has a fix in progress. *(Open, inprogress)*

5. **[#7035 – Intermittent crash on large grep operations](https://github.com/earendil-works/pi/issues/7035)**  
   Solved: root cause was a suckless terminal bug (`RenderAddGlyphs` in X), not Pi itself. *(Closed, no-action)*

6. **[#7153 – `/scoped-models` stalls for ~5 minutes](https://github.com/earendil-works/pi/issues/7153)**  
   Command blocks synchronously on a catalog refresh before rendering any UI. Author suggests non‑blocking loading state. 1 👍. *(Open)*

7. **[#7160 – Function arguments discarded when custom payload is empty](https://github.com/earendil-works/pi/issues/7160)**  
   OpenAI‑compatible providers emitting an empty `custom: {}` object cause Pi to lose valid function arguments. Fix submitted by @sunnyyoung. *(Closed)*

8. **[#7130 – Backspace deletes 2 chars in Kitty terminal](https://github.com/earendil-works/pi/issues/7130)**  
   Kitty protocol release events are not filtered, causing double‑character deletion. *(Open)*

9. **[#5329 – Expose when Pi is waiting on user input for host integrations](https://github.com/earendil-works/pi/issues/5329)**  
   Host integrations (e.g., cmux) need a way to distinguish active turns from user‑input waits. 5 👍 – the highest voted issue. *(Open)*

10. **[#7253 – `/compact` triggers twice when context window reaches 90%](https://github.com/earendil-works/pi/issues/7253)**  
    Manual `/compact` combined with auto‑compact creates an infinite loop until ESC is pressed. *(Open)*

---

## Key PR Progress (Top 10 by Importance)

1. **[#7289 – Comparative Pi eval harness](https://github.com/earendil-works/pi/pull/7289)**  
   Adds a seeded, multi‑harness evaluation framework with score lift, token, latency, and cost deltas. Persists runs in `runs.jsonl`. *(Open)*

2. **[#7288 – Preserve function arguments with empty custom payloads](https://github.com/earendil-works/pi/pull/7288)**  
   Fixes #7160 by preferring valid function tool‑call payloads over empty `custom` objects. *(Closed)*

3. **[#7122 – Byte‐count fix in write tool + false limit warning + surrogate pair handling](https://github.com/earendil-works/pi/pull/7122)**  
   Three independent bug fixes: `content.length` now counts UTF‑8 bytes correctly; `find` tool warning logic fixed; `truncateLine` now respects surrogate pairs. *(Closed)*

4. **[#7286 – Preserve structured metadata for Bedrock provider errors](https://github.com/earendil-works/pi/pull/7286)**  
   Closes #7224 by keeping error metadata from Bedrock instead of serialising `ClientHttp2Stream`. *(Open)*

5. **[#7272 – Preserve providers’ raw stop reason](https://github.com/earendil-works/pi/pull/7272)**  
   Adds `rawStopReason` to assistant messages; improves error messages for Mistral and other providers. Fixes #7255. *(Closed)*

6. **[#7231 – Markdown API](https://github.com/earendil-works/pi/pull/7231)**  
   Implements a programmatic Markdown rendering API for extensions. Closes #6747. *(Open)*

7. **[#7266 – Show system prompt files in startup context banner](https://github.com/earendil-works/pi/pull/7266)**  
   `SYSTEM.md` and `APPEND_SYSTEM.md` now appear in the **[Context]** section alongside `AGENTS.md`. Fixes #7096. *(Closed)*

8. **[#7163 – Search index with SQLite FTS5](https://github.com/earendil-works/pi/pull/7163)**  
   Adds `SessionRepo.search()` supporting full‑text search via SQLite contentless FTS5 virtual tables. *(Open)*

9. **[#7275 – Opt‑in session flush for external workspace managers](https://github.com/earendil-works/pi/pull/7275)**  
   Integrations can now force immediate flush of a session file before the first assistant response, preventing stale paths. *(Closed)*

10. **[#7221 – Stop loading AGENTS.md twice in nested git worktrees](https://github.com/earendil-works/pi/pull/7221)**  
    Prevents duplicate context loading when running from a worktree nested under its main repo. *(Closed)*

---

## Feature Request Trends

- **Model & Provider Expansion** – Support for new models (Kimi K3 on Fireworks, Qwen reasoning effort mapping, Bedrock Mantle OpenAI Responses API) and provider‑specific behaviour preservation (e.g., raw stop reasons).
- **Customisation of Tool Behaviour** – Configurable read tool limits (#3432), truncation limits for large tool outputs (#7066), and finite default waits for bash/search tools (#7284).
- **Rich Content in Conversations** – LaTeX math rendering in Markdown (`$$...$$`) (#7264), audio content blocks in tool results (#7279), and inline images under tmux via sixel (#7245).
- **Session & Workspace Integration** – Better session/resume reactivity (#7285), expose Pi’s waiting state to host integrations (#5329), and display local path packages with proper labels (#7287).
- **Skill and Context Enhancements** – Show system prompt files in startup banner (#7096), signal model when `/skill:name` is invoked (#7282), and preserve saved model/thinking in idle sessions (#7280).

---

## Developer Pain Points

- **Race Conditions & Lock Contention** – Misleading auth errors on parallel startup (#1871), lost tool results when one sibling stalls (#7053).
- **UI & Terminal Inconsistencies** – Kitty double‑backspace (#7130), `/scoped-models` stalls (#7153), `/compact` infinite loop (#7253), TUI crash on undefined tool renderer (#7291).
- **Performance Regressions** – `--mode json` emits `O(n²)` stdout for single tool calls, causing OOM on large writes (#7290).
- **Provider Compatibility Gaps** – Vertex discards Gemini finishReason (#7255), Anthropic stream parser discards initial block (#7283), llama.cpp streaming usage not sent (#7258).
- **Resume and Persistence Bugs** – Resumed ongoing sessions not reactive (#7285), `autocompleteMaxVisible` resets after restart (#7179), idle sessions with saved model/thinking restore global defaults (#7280).
- **Security & Shell Issues** – Example `truncated-tool` executes regex patterns through shell (#7281), clipboard silent no‑op on Wayland (#7248 → #7261).

> *Digest generated from 37 issues and 23 PRs updated on 2026-07-29/30.*

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-07-30

## Today's Highlights

A single nightly release rolled out with a conservative autofix enhancement, while the pull request queue surged with 50 active proposals, dominated by automated takeover PRs from the autofix system itself. Many of these tackle the bot's own inefficiencies — timeout handling, budget limits, and silent failures — signaling a self-improving loop. On the issue tracker, the Fleet Shepherd CI/CD dashboard remains a lingering monitoring concern, and two feature requests were closed after implementation.

## Releases

**v0.21.0-nightly.20260729.0c0ca5fed** — [View release](https://github.com/QwenLM/qwen-code/releases/tag/v0.21.0-nightly.20260729.0c0ca5fed)  
Contains one change: autofix now defers suggestions after five rounds of changes (PR [#7913](https://github.com/QwenLM/qwen-code/pull/7913)). This prevents the bot from flooding a PR with repeated suggestions before human maintainers have had a chance to react.

## Hot Issues

Only three issues were updated in the last 24 hours; all are covered below.

1. **[#7167](https://github.com/QwenLM/qwen-code/issues/7167) — Fleet Shepherd Dashboard**  
   *Status: OPEN | scope/ci-cd*  
   An auto-maintained issue tracking the state of several PRs (#8064, #8037, #7944) and their CI scanning activity. Community has asked for clearer signals when the scan pipeline stalls, but the issue remains in "need-information" state. Key for teams relying on Fleet Shepherd for PR lifecycle management.

2. **[#7084](https://github.com/QwenLM/qwen-code/issues/7084) — Expand restored-history pagination regression coverage**  
   *Status: CLOSED | priority/P3 | scope/testing*  
   Follow-up test enhancements deferred from PR #7064 after five review rounds. Now that the feature is shipped, this issue tracked additional regression tests. Closed with 2 comments. A model of disciplined test expansion after release.

3. **[#8013](https://github.com/QwenLM/qwen-code/issues/8013) — Publication-safe output contract for GitHub channel**  
   *Status: CLOSED | priority/P2 | feature-request*  
   Added a contract that makes GitHub-channel replies safe for public publication and introduces a delivery audit trail. Closed by PR [#8035](https://github.com/QwenLM/qwen-code/pull/8035). Addresses compliance and traceability in automated notifications.

No other issues saw updates; developer attention is clearly focused on the PR flood.

## Key PR Progress

Top 10 PRs by comment count or impact:

1. **[#7944](https://github.com/QwenLM/qwen-code/pull/7944) — Accept tool call OR file content in interactive test**  
   *review/self-reported*  
   Relaxes an overly strict test assertion to accept either a tool call or the expected file content. Prevents false failures on flaky CI runs.

2. **[#8010](https://github.com/QwenLM/qwen-code/pull/8010) — Add seven verification techniques to `verify-pr` skill**  
   *autofix/takeover*  
   Extends the PR verification skill with real-world patterns observed from maintainer review rounds (shared state, collision, writer ordering, etc.). Boosts autonomous review quality.

3. **[#7886](https://github.com/QwenLM/qwen-code/pull/7886) — Tolerate transcript timestamp drift**  
   *autofix/takeover*  
   Makes transcript integrity checks more resilient: timestamp drift no longer causes a full integrity failure; uses SHA-256 snapshots instead. Critical for multi-writer sessions.

4. **[#7885](https://github.com/QwenLM/qwen-code/pull/7885) — Cache npm downloads for verify and tmux build steps**  
   CI optimization caching `npm ci` downloads via `actions/cache@v4`. Expected to reduce CI times for the two slowest jobs.

5. **[#8049](https://github.com/QwenLM/qwen-code/pull/8049) — Back off scan inspection of idle candidates**  
   *autofix/takeover*  
   Reduces CPU waste by skipping full inspection of takeover candidates that have been idle for 10+ hours. Frees up the `MAX_CANDIDATE_INSPECTIONS` budget for active PRs.

6. **[#8035](https://github.com/QwenLM/qwen-code/pull/8035) — Validate and document `reasonFilter`**  
   Follow-up to #8031 hardening notification reason filtering. Adds validation, empty-array edge case handling, and documentation. Closes issue #8013.

7. **[#8016](https://github.com/QwenLM/qwen-code/pull/8016) — Make `/verify` evidence screenshots actually possible**  
   *autofix/takeover*  
   Enables the screenshot slot that `/verify` reports have always declared. Previously 0/14 reports had images; now the infrastructure is wired to produce them.

8. **[#7904](https://github.com/QwenLM/qwen-code/pull/7904) — Throttle Markdown AST parsing during streaming**  
   *feat (web-shell)*  
   Batches Markdown re-parsing at ~80ms intervals instead of after every token. Addresses performance degradation in the web shell with long streaming responses.

9. **[#8014](https://github.com/QwenLM/qwen-code/pull/8014) — Raise `/verify` agent budget from 25m to 120m**  
   *autofix/takeover*  
   Increases the time limit for autonomous verification to match a maintainer's local review session. Silently breaking triple-valued parameter encoded in one place.

10. **[#8066](https://github.com/QwenLM/qwen-code/pull/8066) — Add fork tool execution allowlist**  
    *feat(agent)*  
    Introduces an optional `fork_tools` allowlist for subagent forks. Unlisted tool calls are rejected early, preventing unauthorized tool use in sandboxed agents.

## Feature Request Trends

The following directions emerge from the latest issues and PRs:

- **Autonomous PR verification & maintenance**: Multiple PRs (wenshao) add verification techniques, budget increases, and screenshot generation — signaling a push toward fully autonomous PR review that matches human quality.
- **CI/CD efficiency**: Caching npm downloads, limiting idle candidate inspections, and throttling web shell Markdown parsing all aim to reduce resource waste.
- **Tool execution governance**: PR #8066 (fork allowlist) and #8035 (reasonFilter) reflect a growing focus on fine-grained control over automated actions and notifications.
- **Transcript and context provenance**: Improvements in transcript timestamp tolerance and hook context tagging (PR #7956, #7948) indicate a need for auditable, resilient conversation state.

## Developer Pain Points

Recurring friction areas visible in the data:

- **Autofix system noise**: The flurry of takeover PRs from the autofix system suggests that the bot itself generates significant maintenance overhead — timeout handling (PR #8044), silent cap exceedances (PR #8067), and wasted inspections (#8049) need continuous patching.
- **Test flakiness**: PR #7944 relaxes a test assertion, while issue #7084 tracks additional regression coverage — both point to tests that are brittle or environment-dependent.
- **CI build times**: Adding npm caching (#7885) and throttling AST parsing (#7904) are direct responses to slow CI feedback loops, a common developer frustration.
- **Silent failures**: Multiple PRs address cases where the system fails without user-visible notice (e.g., PR #8067 on round-cap refusals, PR #8014 on budget changes). Developers want transparent, traceable behavior from automated tooling.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-07-30

## Today's Highlights
The v0.9.2 release candidate was finalized after a flurry of bug fixes and hardening PRs, including LaTeX math rendering via Unicode substitution, Indonesian localization for both TUI and website, and a fix for the persistent `reasoning_effort` reset bug. A proposed `/stop` command for interrupting runaway tool-calling has sparked community discussion, while Windows users on Brazilian ABNT2 keyboards remain blocked by an AltGr key conflict.

## Releases
*No new versions were published in the last 24 hours.*

## Hot Issues
1. **[#4959 — Proposed `/stop` command](https://github.com/Hmbown/CodeWhale/issues/4959)** (OPEN, 3 comments)  
   Users want a mechanical interrupt for autonomous tool-call loops when text commands like `+ stop` are ignored. The issue has drawn immediate interest as a safety and UX concern.

2. **[#4949 — Chinese translation of "Constitution"](https://github.com/Hmbown/CodeWhale/issues/4949)** (OPEN, 2 comments)  
   A lively debate among Chinese-speaking contributors on whether to use `宪法` (constitution) or `协作准则` (cooperation guidelines). The issue is unresolved and open for community input.

3. **[#4723 — AltGr+Q on Brazilian ABNT2 opens help instead of typing "/"](https://github.com/Hmbown/CodeWhale/issues/4723)** (OPEN, 2 comments)  
   A critical input bug for Windows users: `Ctrl+Alt+Q` triggers the help overlay instead of inserting `/`. Community is awaiting a fix after today’s PR submission.

4. **[#4789 — Indonesian localization](https://github.com/Hmbown/CodeWhale/issues/4789)** (CLOSED)  
   A feature request that was quickly actioned. The Indonesian TUI pack (1,248 keys) and website locale are now shipped, making CodeWhale more accessible to the largest SEA developer population.

5. **[#4957 — LaTeX math raw display](https://github.com/Hmbown/CodeWhale/issues/4957)** (CLOSED)  
   Users working with technical content reported raw `$...$` source instead of rendered formulas. Fixed in today’s renderer integration.

6. **[#4941 — Thinking level silently reverts to Auto on restart](https://github.com/Hmbown/CodeWhale/issues/4941)** (CLOSED)  
   A persistence bug where the manually set `reasoning_effort` was discarded when using auto model routing. Fixed by keeping routing independent from the preference.

7. **[#4976 — Skills Manager compatible toggle times out on cold Linux filesystems](https://github.com/Hmbown/CodeWhale/issues/4976)** (CLOSED)  
   The sync re-audit of owned skills exceeded the 15-second budget on cold filesystems. Fixed by reusing the owned inventory and scanning only new external roots.

8. **[#4547 — Transcript keeps running spinners for stale shell jobs](https://github.com/Hmbown/CodeWhale/issues/4547)** (CLOSED)  
   Background jobs that disappear or go stale still show animated spinners and Stop controls. Fixed by finalizing stale exec cells and suppressing sidebar spinners.

9. **[#3063 — v0.8.59 release tracker](https://github.com/Hmbown/CodeWhale/issues/3063)** (CLOSED)  
   A backlog issue covering the TUI mouse-report leak on macOS and runtime safety improvements. Served as the basis for subsequent stabilization work.

10. **[#1186 — Typed persistent permission rules](https://github.com/Hmbown/CodeWhale/issues/1186)** (CLOSED)  
    Introduced rules scoped by tool name, command prefix, workspace paths, and decisions (allow/deny/ask). The resulting feature landed via PR #4960.

## Key PR Progress
1. **[#4977 — Fix AltGr-typed "/" on Windows](https://github.com/Hmbown/CodeWhale/pull/4977)** (OPEN)  
    Targets the ABNT2 layout bug by narrowing the `Ctrl-/` help chord to actual left Ctrl. Awaiting review.

2. **[#4975 — Keep Skills Manager scan toggle responsive](https://github.com/Hmbown/CodeWhale/pull/4975)** (CLOSED)  
    Reuses already-audited owned skill rows when expanding to compatible scan, preventing UI timeouts. Merged into v0.9.2.

3. **[#4973 — LaTeX math rendering via Unicode substitution](https://github.com/Hmbown/CodeWhale/pull/4973)** (CLOSED)  
    Contributor `SparkofSpike` implemented detection of `$...$` delimiters and conversion to Unicode approximations.

4. **[#4974 — Hardened LaTeX transcript rendering](https://github.com/Hmbown/CodeWhale/pull/4974)** (CLOSED)  
    Maintainer integration lane that fixed the `\mathbb{R}` path and prevented math preprocessing from rewrapping markdown tables. Closes #4957.

5. **[#4964 — Finalize CodeWhale 0.9.2](https://github.com/Hmbown/CodeWhale/pull/4964)** (CLOSED)  
    The release-finalization PR: fixes Kimi context-window reporting, preserves auto-compaction across settings writes, and repairs composer hints.

6. **[#4961 — Preserve reasoning effort with auto routing](https://github.com/Hmbown/CodeWhale/pull/4961)** (CLOSED)  
    Keeps model routing independent of `reasoning_effort` preference. Merged to fix the silent revert bug.

7. **[#4960 — Add safe rule list and removal for permissions](https://github.com/Hmbown/CodeWhale/pull/4960)** (CLOSED)  
    Implements `/permissions` listing and snapshot-confirmed removal for the typed permission ruleset.

8. **[#4962 — Indonesian documentation suite](https://github.com/Hmbown/CodeWhale/pull/4962)** (CLOSED)  
    Adds `README.id.md`, `CONTRIBUTING.id.md`, and translated docs. Merged quickly after the issue.

9. **[#4937 — Finalize stale shell transcript cells](https://github.com/Hmbown/CodeWhale/pull/4937)** (CLOSED)  
    Stops spinners for stale/absent shell jobs and suppresses sidebar controls. By contributor `LI-Jialu`.

10. **[#4958 — SBOM attestation and provenance mode](https://github.com/Hmbown/CodeWhale/pull/4958)** (CLOSED)  
    Contributor `kobihikri` added explicit attestation flags to release builds, improving supply-chain transparency.

## Feature Request Trends
- **Interrupt & safety controls**: The `/stop` command (#4959) reflects demand for a reliable kill switch during autonomous tool loops.
- **Mathematical rendering**: LaTeX support (#4957) is now resolved after community contribution.
- **Internationalization**: Indonesian was the most active locale request; Traditional Chinese (#4949) is in progress with 499 keys translated.
- **Permission management**: The typed rules system (#1186, #4960) enables granular tool-execution policies.

## Developer Pain Points
- **Keyboard layout conflicts** (Brazilian ABNT2 #4723) — a recurring accessibility gap on Windows.
- **Persistence surprises** — the `reasoning_effort` reset (#4941) and stale shell job spinners (#4547) eroded trust in session state.
- **Cold-start performance** — Skills Manager timeouts on Linux (#4976) highlight the need for incremental scanning.
- **macOS TUI quirks** — mouse-report leaks (#3063) required stabilization work in previous releases.
- **Cross-platform PTY testing** — multiple CI fixes (#4971, #4968) indicate fragility in terminal emulation tests under load.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*