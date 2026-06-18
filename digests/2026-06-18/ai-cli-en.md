# AI CLI Tools Community Digest 2026-06-18

> Generated: 2026-06-18 03:18 UTC | Tools covered: 9

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

## 1. Ecosystem Overview

The AI CLI tool landscape in mid-2026 is marked by rapid iteration and maturing community expectations. Developers across nine major tools are grappling with common pain points—authentication friction, session reliability, subagent orchestration, and cross-platform consistency—while tool makers race to ship niche differentiators like remote control, multi-agent workspaces, and enterprise integrations. Activity is highest in the Claude Code and OpenAI Codex communities, but several smaller tools (Qwen, OpenCode, DeepSeek) demonstrate accelerating release cadences. The ecosystem is converging on shared feature themes—message queuing, token transparency, MCP ecosystem maturation—while diverging in target user profiles: Anthropic and OpenAI target power developers with deep reasoning, Google focuses on agentic autonomy, GitHub on enterprise compliance, and niche tools optimize for specific workflows (CLI-only, multimodal, or lightweight TUI). Supply chain security and platform-specific regressions are emerging as critical quality-of-life concerns across the board.

---

## 2. Activity Comparison

| Tool | Hot Issues Updated Today | PRs Updated Today | Releases Today |
|------|--------------------------|-------------------|----------------|
| Claude Code (Anthropic) | 10 | 7 | v2.1.181 |
| OpenAI Codex | 10 | 10 | 3 alpha (rust-v0.141.0-alpha.5–7) |
| Gemini CLI (Google) | 10 | 10 | v0.47.0, v0.48.0-preview.0 |
| GitHub Copilot CLI | 10 | 0 | None |
| Kimi Code CLI (MoonshotAI) | 2 | 0 | None |
| OpenCode (anomalyco) | 10 | 10 | v1.17.8 |
| Pi (earendil-works) | 10 | 10 | None |
| Qwen Code (QwenLM) | 6 | 10 | v0.18.3 + nightlies |
| DeepSeek TUI (CodeWhale) | 10 | 10 | None |

**Observations:**  
- OpenAI Codex and OpenCode show the highest PR throughput; Codex shipped three alpha builds in 24 hours.  
- Claude Code maintains the highest community engagement volume (1,475 comments on a single issue).  
- Kimi Code CLI is the least active, with only two issues and zero PRs.  
- GitHub Copilot CLI had zero PR activity despite several high-severity issues, suggesting a slower development cycle.

---

## 3. Shared Feature Directions

Several requirements recur across multiple tool communities, indicating broad developer demand:

| Requirement | Involved Tools | Specific Needs |
|-------------|----------------|----------------|
| **Message / Task Queuing** | Claude Code (#50246), OpenCode (multi-agent orchestration #17994) | Queue prompts without interrupting active tasks; multi-agent coordination. |
| **Token & Context Transparency** | Claude Code (#16157 usage limits), Codex (#23794, #8190), OpenCode (#6096 TPS), Gemini (subagent turn limits), Copilot (#3355 configurable context) | Real-time token-per-second, visible context gauge, configurable context window, usage alerts. |
| **Authentication Friction Reduction** | Codex (#25749, #25737), Copilot (#254), Kimi (#2458 SSL), Gemini (Auto Memory secret redaction) | SMS OTP bypass, passkey support, SSL certificate override, persistent login state. |
| **Multi-Agent / Subagent Reliability** | Claude Code (subagent visibility #67485), Gemini (#21409 agent hangs, #22323 false success), Copilot (#3812 subagent MCP access), OpenCode (#17994 orchestration), DeepSeek (#3275 self-looping) | Consistent tool access across subagents, accurate success/failure reporting, guardrails against loops. |
| **MCP / Plugin Ecosystem Maturity** | Claude Code (#26094 MCP parameter serialization), Copilot (#3292 skill-declared MCP, #3787 lazy loading), OpenCode (plugin hooks #32758), Pi (#5654 excludeFromContext) | Proper JSON parameter handling, preloading tools, plugin output array support. |
| **Cross-Platform Consistency** | Claude Code (Windows paste #23146, VS Code drag-drop #25128), Codex (Windows Computer Use #25178, macOS syspolicyd #25719), Gemini (Wayland #21983), OpenCode (ANSI corruption on Windows), Pi (XDG compliance #534), Qwen (TUI unresponsive on SSH #5281) | Parity across Windows, macOS, Linux; terminal emulator compatibility. |
| **Configuration & Permissions** | Claude Code (/config inline syntax), Copilot (#1973 tool whitelist, #3840 instructions opt-out), DeepSeek (#3282 config comment preservation, #3295 honor permission rules), Pi (#5654 excludeFromContext) | Flexible runtime settings, permission granularity, config file integrity. |
| **Session Management** | Claude Code (message queue), Codex (#21211 thread loading), OpenCode (session picker #32752), Pi (tree navigator #5830, session entries RPC #5810), DeepSeek (Workrooms #3209, --session-id) | Resume with spaces, tree views, programmatic session access, persistent chat workspaces. |

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | GitHub Copilot | Kimi Code | OpenCode | Pi | Qwen Code | DeepSeek TUI |
|-----------|-------------|--------------|------------|----------------|-----------|----------|-----|-----------|--------------|
| **Primary differentiator** | Cowork VMs, Remote Control, rich plugin system | Computer Use (screen capture), realtime voice | Subagent autonomy, Auto Memory, AST tools | Enterprise managed models, plugin hooks, MCP | SSL bypass, Agent/Cluster mode | Multi-provider, TPS metrics | Shrinkwrap removal, streaming UX, Azure provider | Channel adapters (QQ, WeChat), provider disambiguation | Workrooms, Plan/Agent mode toggle, permission files |
| **Target user** | Power developers, remote teams | Pro users, desktop-first | Agent orchestration enthusiasts | Enterprise devs, GitHub ecosystem | Corporate/regional users | Multi-provider power users | Advanced CLI users, multimodal | Asian market, multi-endpoint users | Developers wanting lean TUI |
| **Technical approach** | Sandboxed macOS, VS Code extension, Terminal CLI | Desktop app + TUI, Rust alpha | Terminal CLI, subagent orchestration | Terminal CLI integrated with GitHub | Minimal CLI | TUI + Web, plugin system | TUI + RPC, Nix packaging | TUI + channel bots | Rust TUI, toml config |
| **Current focus** | Pricing/rupee, queue mode, MCP fixes | Auth overhaul, context indicator, Crashpad leak | Agent hang, file corruption, subagent reliability | Auth stability, tool whitelist, custom models | SSL bypass, mode switching | Multi-agent, fuzzy edit, DB migration | Shrinkwrap removal, scrolling regression, provider errors | Provider disambiguation, OOM fixes, circuit breaker | Workrooms, permission granularity, static builds |

---

## 5. Community Momentum & Maturity

**Highest community engagement:**  
- **Claude Code** – Single usage-limits issue (#16157) has 1,475 comments and 691 reactions, dwarfing all other communities.  
- **OpenAI Codex** – Token indicator issue (#23794) reached 170 comments and 168 upvotes in a closed time frame; auth issues continue to generate heat.

**Most rapid iteration (24-hour releases):**  
- **OpenAI Codex** – Three alpha releases of the Rust rewrite indicate aggressive development, though stability remains a concern (5+ GB/day Crashpad dumps, post-update DB corruption).  
- **Qwen Code** – v0.18.3 + multiple nightlies, plus 10 PRs, shows disciplined delivery cadence.  
- **OpenCode** – v1.17.8 with critical fixes alongside 10 active PRs suggests healthy maintenance.  
- **Claude Code** and **Gemini CLI** also shipped releases, but with fewer structural changes.

**Maturing but slower:**
- **GitHub Copilot CLI** – No PRs in 24 hours; stagnant auth and transient error issues (e.g., #254 open 8 months) indicate slower resolution.  
- **Kimi Code CLI** – Minimal activity; lacks the scale to generate strong feedback loops.  
- **Pi** – Consistent PR volume but no release; Shrinkwrap refactor is long-running.  
- **DeepSeek TUI** – Busy PR pipeline (25 total) with no release; feature depth (Workrooms, permissions) promising but not yet shipped.

**Emerging threats:**  
- **Qwen Code** and **OpenCode** are gaining traction in specific niches (Asian markets, multi-provider power users) and could challenge incumbents if they improve community support.

---

## 6. Trend Signals

1. **Authentication as a blocker** – The most frequent and severe pain point across tools. Users with passkeys, Google OAuth, or enterprise SSO face hard lockouts (Codex, Copilot, Kimi). CLI tools must invest in headless, browser-optional auth flows that respect modern security primitives.

2. **Multi-agent orchestration is the next frontier** – Claude Code’s subagent visibility request, Gemini’s agent loops, Copilot’s MCP tool access for subagents, OpenCode’s orchestration feature, and DeepSeek’s Workrooms all point to a community demand for coordinated agent teams with clear permissions and lineage.

3. **Token/context weariness** – Developers want to see, manage, and budget their token consumption in real time. The “context window ran out” error without recovery path (Codex #8190) is unacceptable; configurable context windows (Copilot #3355) and TPS indicators (OpenCode) are becoming table stakes.

4. **MCP ecosystem fragility** – Multiple tools report serialization bugs, lazy loading issues, and plugin hook failures. The community expects MCP to “just work” for common integrations (Notion, file systems, design tools). Tool makers must prioritize MCP core stability over new features.

5. **Enterprise compliance vs. developer freedom** – Corporate proxies, SSL inspection, content-exclusion policies, and managed custom models create friction. Tools that offer escape hatches (e.g., `--insecure` flags, cert pinning, configurable timeouts) while maintaining default security will win enterprise adoption.

6. **Silent resource leaks erode trust** – Unbounded Crashpad dumps (Codex +5GB/day), infinite Snapchat-like retries (Gemini Auto Memory), and unchecked snapshot disk usage (DeepSeek) show that background process hygiene is a growing concern. Developers expect tools to be good citizens.

7. **Cross-platform quality gap** – Windows Ctrl+V conflicts, macOS CPU runaway, Linux `SIGTSTP` suspend, Wayland browser failures – terminal emulator variations remain a major source of unreported bugs. A proactive testing matrix across Windows Terminal, iTerm2, GNOME Terminal, and others is overdue.

8. **Pricing regionalisation** – Claude Code and Gemini have strong Indian user demand for INR pricing; OpenAI and Google already offer it. Token-based pricing models that don’t lock quality behind “Max” subscriptions will be a competitive lever.

*Report generated from community digests dated 2026-06-18. Data reflects last 24 hours of activity on each tool’s public GitHub repository.*

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
*Data as of 2026-06-18 | Source: github.com/anthropics/skills*

---

## 1. Top Skills Ranking

The following pull requests represent the most-discussed Skill submissions by community engagement. All remain **open** as of the data snapshot.

### #514 — Document Typography Skill
- **Functionality:** Prevents orphan word wrap, widow paragraphs, and numbering misalignment in AI-generated documents—addressing pervasive typographic issues across all Claude document output.
- **Discussion Highlights:** The PR explicitly frames typographic quality as a universal pain point ("These issues affect every document Claude generates"), generating broad community resonance.
- **Status:** Open
- [GitHub: PR #514](https://github.com/anthropics/skills/pull/514)

### #486 — ODT (OpenDocument) Skill
- **Functionality:** Enables creation, filling, reading, and conversion of OpenDocument Format files (.odt, .ods), including template filling and ODT-to-HTML parsing. Triggers on mentions of ODF/LibreOffice terminology.
- **Discussion Highlights:** Addresses a clear interoperability gap for users in open-source and European public-sector environments where ODF is mandatory.
- **Status:** Open
- [GitHub: PR #486](https://github.com/anthropics/skills/pull/486)

### #210 — Frontend-Design Skill Improvement
- **Functionality:** Revisions to the existing frontend-design skill to improve clarity, actionability, and internal coherence—ensuring every instruction is executable within a single conversation.
- **Discussion Highlights:** Focus on making skill guidance *specific enough to steer behavior without being overly prescriptive*—a recurring theme in skill quality discussions.
- **Status:** Open
- [GitHub: PR #210](https://github.com/anthropics/skills/pull/210)

### #83 — Skill Quality & Security Analyzers
- **Functionality:** Two meta-skills for evaluating other Skills across five dimensions: Structure & Documentation (20%), plus security analysis coverage.
- **Discussion Highlights:** Represents the community's growing interest in **skill governance and quality assurance**—tools to self-regulate the ecosystem.
- **Status:** Open
- [GitHub: PR #83](https://github.com/anthropics/skills/pull/83)

### #181 — SAP-RPT-1-OSS Predictor Skill
- **Functionality:** Wraps SAP's open-source tabular foundation model (Apache 2.0, released at TechEd 2025) for predictive analytics on SAP business data.
- **Discussion Highlights:** Enterprise AI use case with significant downstream demand from SAP ecosystem developers.
- **Status:** Open
- [GitHub: PR #181](https://github.com/anthropics/skills/pull/181)

### #1298 — Skill-Creator Eval Fix (Critical Bug)
- **Functionality:** Fixes `run_eval.py` reporting `recall=0%` for every skill description—the description-optimization loop was optimizing against noise. Includes Windows stream reading, trigger detection, and parallel worker fixes.
- **Discussion Highlights:** Directly addresses the most reported bug in the repository (see Issue #556). High urgency given it blocks skill description optimization entirely.
- **Status:** Open
- [GitHub: PR #1298](https://github.com/anthropics/skills/pull/1298)

### #723 — Testing-Patterns Skill
- **Functionality:** Comprehensive testing stack coverage: Testing Trophy philosophy, unit testing (AAA pattern), React component testing with Testing Library, and guidance on what *not* to test.
- **Discussion Highlights:** Fills a noticeable gap—no existing skill covered testing methodology holistically.
- **Status:** Open
- [GitHub: PR #723](https://github.com/anthropics/skills/pull/723)

### #444 — AURELION Skill Suite
- **Functionality:** Four-skills ecosystem (kernel, advisor, agent, memory) implementing a structured 5-floor cognitive framework for professional knowledge management and AI collaboration.
- **Discussion Highlights:** Ambitious scope raises questions about skill granularity vs. monolithic design; represents the "cognitive framework" sub-community.
- **Status:** Open
- [GitHub: PR #444](https://github.com/anthropics/skills/pull/444)

---

## 2. Community Demand Trends

Analysis of the most commented Issues reveals four concentrated demand signals:

| Trend | Evidence (Issue) | Signal Strength |
|---|---|---|
| **Org-wide skill sharing & distribution** | #228 (14 comments, 7 👍): Users want direct sharing links or libraries—not manual `.skill` file transfer via Slack | **Highest** |
| **Skill creation tooling reliability** | #556 (12 comments, 7 👍): `run_eval.py` zero-trigger bug blocks all automated skill optimization. #1169 (3 comments): Same bug, independent replication. #1061 (3 comments): Windows compatibility blockers. | **High** |
| **Security & trust boundaries** | #492 (7 comments, 2 👍): Community skills under `anthropic/` namespace creates impersonation risk. #1175 (4 comments): Access control concerns in SharePoint document handling. | **Medium** |
| **Deduplication & ecosystem hygiene** | #189 (6 comments, 9 👍): `document-skills` and `example-skills` plugins install identical content, wasting context window. | **Medium** |

**Emerging directional demand:** Agent governance/safety patterns (Issue #412, closed but 6 comments) and MCP integration for skills (Issue #16) represent early-stage but repeated requests.

---

## 3. High-Potential Pending Skills

These PRs combine active discussion, clear utility, and maintainer engagement signals:

- **#538 — PDF case-sensitive file references fix** – Corrects 8 case-mismatches breaking skills on Linux/macOS. Low complexity, high correctness impact. [GitHub](https://github.com/anthropics/skills/pull/538)

- **#539/#361 — YAML special character validation** – Two PRs solving the same problem (unquoted `:` in descriptions causing silent parsing failures). Community energy is concentrated; one will likely merge. [PR #539](https://github.com/anthropics/skills/pull/539) | [PR #361](https://github.com/anthropics/skills/pull/361)

- **#541 — DOCX tracked change ID collision fix** – Prevents document corruption when bookmarks share `w:id` space with tracked changes. Niche but breaking for all DOCX skill users. [GitHub](https://github.com/anthropics/skills/pull/541)

- **#568 — ServiceNow platform skill** – Broad coverage (ITSM, ITOM, SecOps, ITAM, FSM, CSDM, IntegrationHub). Enterprise demand; recent updates suggest active iteration. [GitHub](https://github.com/anthropics/skills/pull/568)

- **#1050/#1099 — Windows compatibility fixes** – Two independent PRs fixing `subprocess` `PATHEXT` issues and pipe reading crashes. High potential to merge given multiple reporters on Issue #1061. [PR #1050](https://github.com/anthropics/skills/pull/1050) | [PR #1099](https://github.com/anthropics/skills/pull/1099)

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is for *skill-authoring infrastructure reliability*—fixing the broken evaluation loop, YAML parsing, and Windows compatibility—before the ecosystem can sustainably scale new Skill submissions.**

Document-format Skills (typography, ODT, DOCX fixes) form the second-strongest cluster, reflecting Claude Code's heavy use as a document generation engine. Enterprise platform Skills (ServiceNow, SAP) represent a smaller but vocal cohort with distinct ecosystem needs around security and namespace trust.

---

# Claude Code Community Digest – June 18, 2026

## Today's Highlights

A new release (v2.1.181) introduces inline `/config key=value` syntax for settings and a macOS sandbox opt-in for Apple Events. On the community side, the long‑running usage‑limits issue (#16157) now has over 1,475 comments and 691 reactions, making it the most urgent pain point. New feature requests for message queuing, India‑specific pricing, and better subagent visibility continue to dominate the enhancement queue.

## Releases

**v2.1.181** – released today

- `/config key=value` syntax – set any setting directly from the prompt (works in interactive, `-p`, and Remote Control modes)
- `sandbox.allowAppleEvents` – opt‑in setting that enables sandboxed commands to send Apple Events on macOS
- `CLAUDE_CLIENT_P`… (truncated in changelog – likely `CLAUDE_CLIENT_PATH` or similar environment variable support)

[Release details](https://github.com/anthropics/claude-code/releases/tag/v2.1.181)

---

## Hot Issues (10 Noteworthy)

1. **#16157 – Instantly hitting usage limits with Max subscription**  
   *1,475 comments, 691 👍*  
   The most discussed issue. Users on the Max tier report being capped immediately after purchase. Anthropic has labeled it `oncall` and `area:cost`, but resolution is still pending.  
   [Issue #16157](https://github.com/anthropics/claude-code/issues/16157)

2. **#17432 – India‑specific pricing plans (INR)**  
   *198 comments, 444 👍*  
   Strong demand for local currency pricing, similar to OpenAI and Google. Affects both Claude Pro and Claude Code subscriptions.  
   [Issue #17432](https://github.com/anthropics/claude-code/issues/17432)

3. **#34255 – Remote Control: automatic reconnection fails – connection drops silently**  
   *50 comments, 90 👍*  
   macOS/iOS users report that the connection drops without recovery. No built‑in retry logic, forcing manual restart.  
   [Issue #34255](https://github.com/anthropics/claude-code/issues/34255)

4. **#50246 – Message queue mode – queue prompts instead of interrupting active tasks**  
   *32 comments, 99 👍*  
   Highly requested: allow users to queue follow‑ups while Claude is busy, rather than forcing an interrupt.  
   [Issue #50246](https://github.com/anthropics/claude-code/issues/50246)

5. **#39636 – Cowork VM guest kernel never boots on Snapdragon X Plus (ARM64)**  
   *29 comments, 9 👍*  
   Windows on ARM users cannot start Cowork VMs. Connection timeout on every attempt.  
   [Issue #39636](https://github.com/anthropics/claude-code/issues/39636)

6. **#25128 – Drag and drop not working in VS Code extension chat panel**  
   *20 comments, 40 👍*  
   Regression since v2.1.6. Drag‑and‑drop is fully functional in terminal CLI but completely broken in the VS Code extension panel.  
   [Issue #25128](https://github.com/anthropics/claude-code/issues/25128)

7. **#5277 – Image paste in SSH / remote sessions**  
   *17 comments, 31 👍*  
   Users on remote servers cannot paste images via the CLI. Workarounds exist but native handling is desired.  
   [Issue #5277](https://github.com/anthropics/claude-code/issues/5277)

8. **#63870 – Bash tool calls emitted as raw `<invoke>` text instead of executing**  
   *16 comments, 20 👍*  
   A major reliability bug: some sessions produce 23+ malformed calls that are printed as text rather than executed. Multiple duplicates exist.  
   [Issue #63870](https://github.com/anthropics/claude-code/issues/63870)

9. **#26094 – MCP object parameters serialized as strings (Claude Desktop Cowork)**  
   *13 comments, 19 👍*  
   Breaks Notion MCP writes because object params are serialised as raw strings instead of proper JSON. Affects MCP integration reliability.  
   [Issue #26094](https://github.com/anthropics/claude-code/issues/26094)

10. **#23146 – Ctrl+V in `/resume` conflicts with Windows paste**  
    *8 comments, 6 👍 (CLOSED)*  
    Windows users cannot rebind the default paste shortcut (`Ctrl+V`) when entering the `/resume` command. Marked as closed, but the conflict remains unresolved for many.  
    [Issue #23146](https://github.com/anthropics/claude-code/issues/23146)

---

## Key PR Progress (7 total)

1. **#41611 – "add the missing source to claude code"**  
   *OPEN* – A user submission adding source files; likely a community attempt at open‑sourcing parts of the CLI. Not merged.  
   [PR #41611](https://github.com/anthropics/claude-code/pull/41611)

2. **#41447 – "feat: open source claude code"**  
   *OPEN* – Ambitious PR claiming to close multiple open‑source issues. No maintainer activity; likely a symbolic request.  
   [PR #41447](https://github.com/anthropics/claude-code/pull/41447)

3. **#69226 – Update frontend-design skill**  
   *CLOSED (merged)* – Improvements to the frontend-design skill plugin. Bumps version to 1.1.0 so existing installations auto‑update.  
   [PR #69226](https://github.com/anthropics/claude-code/pull/69226)

4. **#19867 – Fix code‑review: allow re‑reviews when new commits are pushed**  
   *OPEN* – Adds smarter skip logic that checks for new commits since Claude’s last comment. Documents `--force` flag to bypass the check.  
   [PR #19867](https://github.com/anthropics/claude-code/pull/19867)

5. **#33443 – Fix Dockerfile to use native installer**  
   *OPEN* – Updates `.devcontainer/Dockerfile` from deprecated `npm install` to the native installer for Node 24.14.  
   [PR #33443](https://github.com/anthropics/claude-code/pull/33443)

6. **#60427 – Docs: use standard GitHub capitalization in README**  
   *CLOSED (merged)* – Minor fix for product name casing in the README.  
   [PR #60427](https://github.com/anthropics/claude-code/pull/60427)

7. **#60732 – Docs: polish plugins README wording**  
   *CLOSED (merged)* – Tiny wording improvement for the plugin ecosystem description.  
   [PR #60732](https://github.com/anthropics/claude-code/pull/60732)

---

## Feature Request Trends

The community’s most‑voiced feature directions fall into five categories:

- **Pricing & regionalisation** – India‑specific INR plans (#17432) and a general call for more flexible cost structures.
- **Workflow ergonomics** – A message queue mode (#50246, #68998) to avoid interrupting active tasks; better visibility into background subagent activity (#67485); and queuing multiple prompts.
- **IDE & platform integration** – Auto‑accept edits in JetBrains (#69241), restoring drag‑and‑drop in VS Code (#25128), and native Claude Design handoff (#69239, #69246).
- **Remote & headless usage** – Image paste over SSH (#5277), stable Remote Control reconnection (#34255), and OAuth support for MCP servers on remote machines (#69205).
- **MCP & design‑system interaction** – Querying external design systems via MCP (#60327) and fixing MCP parameter serialisation (#26094).

---

## Developer Pain Points

- **Usage limits on Max subscription** (#16157) – the most severe and widespread complaint, with nearly 1,500 comments. Users feel locked out despite paying top tier.
- **Cost transparency** – CLI embedded price maps are sometimes incorrect (#64701), and users want real‑time cost visibility.
- **Connection reliability** – Remote Control silently dropping (#34255) and Cowork VMs failing to boot on ARM64 (#39636) disrupt headless workflows.
- **Cross‑platform inconsistencies** – Windows Ctrl+V conflicts (#23146), VS Code extension bugs (#25128), and WSL/terminal paste issues (#61043) show uneven parity.
- **Core execution bugs** – Bash tool calls rendered as raw text (#63870) and idle sessions spinning CPU at 100% on macOS (#68931) degrade basic usability.
- **Claude Design handoff** – The “send to coding agent” flow is broken because the generated prompt references a connector the CLI does not have (#69246).

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-06-18

## Today's Highlights
Three new Rust alpha releases (v0.141.0-alpha.5 through .7) shipped within the last 24 hours, signaling active iteration on the CLI/TUI backend. Community attention remains sharply focused on persistent authentication friction — particularly SMS OTP forced step-up during CLI login and inaccessible legacy phone number recovery — alongside mounting frustration over unbounded Crashpad dump accumulation and database corruption after updates. A notable closed issue (#23794) regarding the disappearance of the visible token usage indicator in Codex Desktop gathered 170 comments and 168 upvotes before resolution, reflecting strong user demand for transparent quota and context-limit feedback.

---

## Releases
Three Rust alpha releases were published (no changelogs beyond version bumps):
- **rust-v0.141.0-alpha.5** — [Release](https://github.com/openai/codex/releases/tag/rust-v0.141.0-alpha.5)
- **rust-v0.141.0-alpha.6** — [Release](https://github.com/openai/codex/releases/tag/rust-v0.141.0-alpha.6)
- **rust-v0.141.0-alpha.7** — [Release](https://github.com/openai/codex/releases/tag/rust-v0.141.0-alpha.7)

---

## Hot Issues (Top 10 Noteworthy)

1. **#23794** — [CLOSED] *Codex Desktop no longer shows visible context/token usage indicator*  
   Author: HushHuang | Comments: 170 | 👍: 168  
   This was the most active issue by a wide margin. Users strongly rely on the token/context gauge for session management; its removal caused confusion and frustration. Now closed, likely implying a fix or restoration is incoming.  
   [Issue #23794](https://github.com/openai/codex/issues/23794)

2. **#25749** — [OPEN] *Codex requires verification of inaccessible legacy phone number*  
   Author: Walker1o | Comments: 49 | 👍: 30  
   A major auth UX gap: users with Google OAuth + MFA can access ChatGPT but cannot authenticate in Codex, as the CLI/desktop forces SMS OTP step-up to a legacy number. No recovery path exists, creating account lockout.  
   [Issue #25749](https://github.com/openai/codex/issues/25749)

3. **#25719** — [OPEN] *macOS syspolicyd / trustd CPU and memory runaway*  
   Author: energissimo-mg | Comments: 31 | 👍: 39  
   Codex Desktop on macOS causes persistent high CPU/memory consumption from `syspolicyd` and `trustd` processes. This is a platform-level performance issue affecting all macOS users with the desktop app.  
   [Issue #25719](https://github.com/openai/codex/issues/25719)

4. **#17827** — [OPEN] *Customizable status line (TUI)*  
   Author: pkondaurov | Comments: 16 | 👍: 71  
   A long-standing feature request (since April 2026) with very high community support. Users want a customizable terminal status line (token usage, model name, git branch) similar to Claude Code's implementation.  
   [Issue #17827](https://github.com/openai/codex/issues/17827)

5. **#21211** — [OPEN] *Thread navigation/loading slows from unbounded metadata*  
   Author: clairernovotny | Comments: 12 | 👍: 2  
   Performance regression in thread loading caused by unbounded metadata bloat in SQLite. Supersedes an earlier issue; the core team is aware but the problem persists across versions.  
   [Issue #21211](https://github.com/openai/codex/issues/21211)

6. **#24006** — [OPEN] *macOS: Codex cannot access its local database after update*  
   Author: garrettXu | Comments: 11 | 👍: 9  
   A recurring post-update crash on macOS where the app fails to open due to database access errors. Users on Pro subscriptions are affected, with no workaround besides full reinstall.  
   [Issue #24006](https://github.com/openai/codex/issues/24006)

7. **#25737** — [OPEN] *CLI login forces SMS OTP despite security-key-only account*  
   Author: Jesssullivan | Comments: 11 | 👍: 6  
   FIDO2/CTAP2 hardware security key authentication succeeds in the browser flow, but the CLI redirect then forces SMS OTP verification. This breaks the security model for passkey-first users.  
   [Issue #25737](https://github.com/openai/codex/issues/25737)

8. **#25178** — [OPEN] *Windows Computer Use screenshot fails on Windows 10 22H2*  
   Author: Define1165250535 | Comments: 11 | 👍: 4  
   A Windows-specific bug where `get_window_state` calls fail with `SetIsBorderRequired` error (0x80004002). Computer Use can list apps and send keyboard input but cannot capture screenshots.  
   [Issue #25178](https://github.com/openai/codex/issues/25178)

9. **#25921** — [OPEN] *Crashpad pending dumps grow without limit (+5GB/day)*  
   Author: Jolg42 | Comments: 9 | 👍: 2  
   Codex Desktop continuously writes `.dmp` and `_sidecar.json` files under `Crashpad/pending`, consuming 5+GB per day. A silent resource leak that can fill user disks.  
   [Issue #25921](https://github.com/openai/codex/issues/25921)

10. **#8190** — [OPEN] *Context window exhausted — no recovery path*  
    Author: robertotcestari | Comments: 8 | 👍: 2  
    Users hitting `Context window ran out of room` receive no actionable guidance for recovery. A long-running issue (since Dec 2025) affecting Pro users on GPT-5.2.  
    [Issue #8190](https://github.com/openai/codex/issues/8190)

---

## Key PR Progress (Top 10 Important PRs)

1. **#28838** — [OPEN] *Support Codex home instructions directory*  
   Author: charlesgong-openai  
   Adds support for loading `*.md` files from `~/.codex/instructions/` as global instructions, preserving existing override precedence. A significant usability improvement for power users.  
   [PR #28838](https://github.com/openai/codex/pull/28838)

2. **#19049** — [OPEN] *Opt ChatGPT auth into agent identity*  
   Author: adrian-openai  
   Part of a multi-PR stack for Agent Identity JWT authentication. This is foundational work for enabling agent-run tasks with proper identity assertion.  
   [PR #19049](https://github.com/openai/codex/pull/19049)

3. **#28836** — [OPEN] *Support assistant realtime append text*  
   Author: guinness-oai  
   Enables replay of previous-session overlap as actual conversation items in realtime voice sessions. Fixes a role-enum incompatibility in the Rust websocket layer.  
   [PR #28836](https://github.com/openai/codex/pull/28836)

4. **#28835** — [OPEN] *Add app-server current-time provider*  
   Author: rka-oai  
   Implements a server-backed current-time clock and reminder system. This enables time-aware agent behavior and is part of the "varlatency" feature series.  
   [PR #28835](https://github.com/openai/codex/pull/28835)

5. **#28813** — [OPEN] *Pause active goals before Esc interrupts*  
   Author: etraut-openai  
   Fixes #28104: pressing Esc during an active `/goal` now pauses the goal state properly, aligning behavior with Ctrl+C. Addresses a TUI workflow inconsistency.  
   [PR #28813](https://github.com/openai/codex/pull/28813)

6. **#28814** — [OPEN] *Assign response item IDs when recording history*  
   Author: pakrym-oai  
   Ensures client-created response items receive IDs at the history-recording boundary, preserving identity across persistence and resume. Critical for reliable thread reconstruction.  
   [PR #28814](https://github.com/openai/codex/pull/28814)

7. **#28824** — [OPEN] *Current time reminders for system clock*  
   Author: rka-oai  
   Adds a host-injectable current-time provider with a built-in system clock implementation. Records UTC developer reminders in history immediately before model requests.  
   [PR #28824](https://github.com/openai/codex/pull/28824)

8. **#27986** — [CLOSED] *Control automatic realtime handoff delivery*  
   Author: jiayuhuang-openai  
   Adds optional `codexResponseHandoffPrefix` to realtime handoff configuration, giving control over automatic V1 commentary vs. final answers. Merged and closed.  
   [PR #27986](https://github.com/openai/codex/pull/27986)

9. **#27132** — [OPEN] *Emit Trusted MCP App Identity on Tool-Call Items*  
   Author: martinauyeung-oai  
   Adds `appContext` metadata to MCP tool-call items, including trusted `connectorId`, `linkId`, and `mcpAppResourceUri`. Preserves identity across tool-call events and reconnects.  
   [PR #27132](https://github.com/openai/codex/pull/27132)

10. **#28784** — [CLOSED] *Fix: support older awk checksum parsing*  
    Author: fcoury-oai  
    Fixes the standalone installer failing on systems with older `mawk` (common on Debian). The awk interval expression was incompatible; now accepts valid 64-character digests across all architectures.  
    [PR #28784](https://github.com/openai/codex/pull/28784)

---

## Feature Request Trends

The most-requested feature directions across all issues:

- **Customizable TUI status line** (#17827, 71 👍): Persistent demand for a configurable terminal status bar showing token usage, model name, rate limits, and git branch — mirroring Claude Code's functionality.
- **Usage/reset management in CLI** (#28827, #28823): Users want a `/reset` or usage-meter feature directly in the CLI/Codex Cloud, not just in the desktop app.
- **Context window transparency** (#23794, #8190): Strong desire for real-time visibility into context/token consumption to proactively manage session limits.
- **Better recovery from context exhaustion** (#8190): When the model's context window fills, users want actionable recovery paths (e.g., automatic compaction, smarter history pruning).
- **Remote plugin sharing and discovery** (#26703, #26704): Multiple stacked PRs indicate active work on plugin sharing, workspace-wide plugin catalogs, and remote catalog rendering.

---

## Developer Pain Points

Recurring frustrations emerging from high-activity issues:

- **Authentication friction**: The most severe pain point. Users with modern auth setups (passkeys, security keys, Google OAuth) are forced into SMS OTP flows (#25749, #25737). For CLI and desktop users, this creates hard lockouts, especially for Business/Enterprise accounts.
- **Post-update database corruption**: Multiple reports (#24006, #24030, #28666) of SQLite state databases becoming malformed after Codex Desktop updates. No recovery path other than full reinstall, which loses session history.
- **Platform-specific performance regressions**: macOS `syspolicyd` CPU runaway (#25719), Windows crash on Korean-character usernames (#28262), and unbounded Crashpad dumps (#25921) suggest inadequate cross-platform testing.
- **Computer Use instability**: Cross-platform screenshot and tool-call failures (#25178 on Windows, #24207 on macOS Intel) indicate the Computer Use feature is still brittle outside ideal environments.
- **Thread performance degradation**: Unbounded metadata bloat (#21211) and slow navigation in long-running sessions remain unresolved, affecting power users with many threads.
- **Silent resource leaks**: The Crashpad dump accumulation (#25921) and unbounded SQLite bloat (#21211) point to missing storage-management guardrails in the desktop app.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest – 2026-06-18

## Today's Highlights
Two patch releases landed (v0.47.0 and v0.48.0-preview.0) while the community continues to report agent hangs and subagent reliability issues. A critical PR fixing Jupyter Notebook and JSON corruption in `write_file` was opened, and several security‑focused PRs aim to harden CI against fork‑based attacks.

## Releases
- **[v0.47.0](https://github.com/google-gemini/gemini-cli/releases/tag/v0.47.0)** – Routine bump with changelog updates and a fix for “Respect backend def”.  
- **[v0.48.0-preview.0](https://github.com/google-gemini/gemini-cli/releases/tag/v0.48.0-preview.0)** – Preview release; includes a dependabot cooldown for npm packages and internal refactoring.

## Hot Issues (Top 10)
1. **[#24353 – Robust component level evaluations](https://github.com/google-gemini/gemini-cli/issues/24353)**  
   Epic for scaling behavioral evals (76 tests, 6 models). No community reaction yet but critical for quality assurance.

2. **[#22745 – AST‑aware file reads, search, and mapping](https://github.com/google-gemini/gemini-cli/issues/22745)**  
   Investigation into using AST tools (tilth, glyph) for precise code understanding. 1 👍, 7 comments – strong interest from power users.

3. **[#21409 – Generalist agent hangs](https://github.com/google-gemini/gemini-cli/issues/21409)**  
   **8 👍** – top pain point. CLI hangs indefinitely when deferring to generalist agent. Users work around by disabling subagents.

4. **[#22323 – Subagent recovery after MAX_TURNS reported as GOAL success](https://github.com/google-gemini/gemini-cli/issues/22323)**  
   `codebase_investigator` lies about success when hitting turn limits. Misleading status hides real failures. 2 👍.

5. **[#21968 – Gemini does not use skills and sub‑agents enough](https://github.com/google-gemini/gemini-cli/issues/21968)**  
   Agent ignores custom skills unless explicitly instructed. Community wants better self‑orchestration.

6. **[#26525 – Add deterministic redaction and reduce Auto Memory logging](https://github.com/google-gemini/gemini-cli/issues/26525)**  
   Security concern: secrets sent to model before redaction. Plus noisy logging.

7. **[#26522 – Stop Auto Memory from retrying low‑signal sessions indefinitely](https://github.com/google-gemini/gemini-cli/issues/26522)**  
   Auto Memory never marks low‑signal sessions as processed, causing infinite retries.

8. **[#25166 – Shell command execution gets stuck with “Waiting input” after command completes](https://github.com/google-gemini/gemini-cli/issues/25166)**  
   **3 👍** – simple commands hang, showing “Awaiting user input” despite already finishing.

9. **[#21983 – Browser subagent fails on Wayland](https://github.com/google-gemini/gemini-cli/issues/21983)**  
   Browser agent reports GOAL but actually fails. 1 👍; affects Linux users.

10. **[#20079 – Symlinked agent files not recognized](https://github.com/google-gemini/gemini-cli/issues/20079)**  
    `~/.gemini/agents/filename.md` ignored if a symlink. Low effort fix, blocks dotfiles management.

## Key PR Progress (Top 10)
1. **[#28000 – fix(core-tools): resolve Jupyter Notebook and JSON corruption in write_file](https://github.com/google-gemini/gemini-cli/pull/28000)**  
   Critical bug fix – `write_file` silently corrupts `.ipynb` and `.json` files. New PR.

2. **[#28002 – Changelog for v0.47.0](https://github.com/google-gemini/gemini-cli/pull/28002)**  
   Auto‑generated changelog for the latest stable release.

3. **[#27648 – feat(core): support list format in trustedFolders.json](https://github.com/google-gemini/gemini-cli/pull/27648)**  
   Adds easier manual editing of trusted folders. Already closed.

4. **[#27649 – fix(docs): fix broken telemetry page](https://github.com/google-gemini/gemini-cli/pull/27649)**  
   Documentation fix – Traces section incorrectly scoped under Metrics.

5. **[#27788 – test(core): add subfolder ignore test for getFolderStructure](https://github.com/google-gemini/gemini-cli/pull/27788)**  
   Ensures `.gitignore` rules apply to subdirectories.

6. **[#27780 – security: gate chained E2E on same‑repository checkout](https://github.com/google-gemini/gemini-cli/pull/27780)**  
   Prevents fork‑supplied repository from leaking `GEMINI_API_KEY`. Security hardening.

7. **[#27948 – chore(deps): pin dependencies and enforce 14‑day update cooldown](https://github.com/google-gemini/gemini-cli/pull/27948)**  
   Strict dependency pinning + Dependabot cooldown to reduce supply chain risk.

8. **[#27996 – fix(core): decode response body using charset from Content‑Type header](https://github.com/google-gemini/gemini-cli/pull/27996)**  
   `web-fetch` now respects non‑UTF‑8 charsets (GBK, ISO‑8859‑1). Important for i18n.

9. **[#27994 – fix(core): insert skill/agent content literally in system prompt substitutions](https://github.com/google-gemini/gemini-cli/pull/27994)**  
   Prevents `$` and backreferences in skill descriptions from being mangled.

10. **[#27859 – feat(cli): add native drag‑and‑drop and Cmd+V clipboard image pasting](https://github.com/google-gemini/gemini-cli/pull/27859)**  
    Enables visual multimodal input in the terminal. Long‑requested UX improvement.

## Feature Request Trends
- **AST‑aware code understanding** (#22745, #22746) – users want precise method/class navigation and codebase mapping using abstract syntax trees.
- **Better agent self‑awareness** (#21432) – ability to explain its own CLI flags, hotkeys, and self‑execution workflows.
- **Destructive behavior guardrails** (#22672) – model should avoid `git reset --force` and other risky commands.
- **Browser agent resilience** (#22232) – automatic session takeover and lock recovery, especially for persistent sessions.
- **Local subagent improvements** (#20195, #21000) – native file tools for task tracker, sprint 1 deliverables.
- **Remote agents** (#20303) – advanced auth and background operations for enterprise use.

## Developer Pain Points
- **Agent hangs and freezes** ( #21409, #25166, #22465) – especially with generalist agent or interactive prompts (vite). High frustration.
- **Subagent reliability** ( #22323, #21968, #22093) – false success reports, ignores skills, runs without permission.
- **File corruption** ( #28000) – `write_file` silently corrupts Jupyter/JSON files; blocks data science workflows.
- **Configuration overrides ignored** ( #22267) – Browser agent does not respect `settings.json` for `maxTurns`.
- **Symlink and path issues** ( #20079, #27990) – symlinked agents not recognized; macOS `/var`→`/private/var` mismatches.
- **Auto Memory inefficiency** ( #26522, #26523) – infinite retries on low‑signal sessions, silent skip of invalid patches.
- **Security concerns** ( #26525) – secrets leak before redaction in Auto Memory.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest – 2026-06-18

## Today’s Highlights
The aftermath of the June 16 Copilot outage continues to affect users, with multiple issues reporting that all models appear “Blocked / Disabled” or that transient API errors stall workflows. A long-standing request for a tool whitelist in interactive mode (Issue #1973) remains the most upvoted feature request, while the plugin hooks community is pushing for silent command rewrites (Issue #2643). Several new triage items highlight friction with enterprise content-exclusion policies and custom model routing.

## Releases
No new releases were published in the last 24 hours.

---

## Hot Issues (10 Noteworthy)

1. **#2643 – Silent command rewrite via `updatedInput` still shows confirmation dialog**  
   [Link](https://github.com/github/copilot-cli/issues/2643)  
   Users want `preToolUse` hooks to silently rewrite commands when `permissionDecision: allow` is set, but the CLI always shows an interactive confirmation. 10 comments, 1 👍.  
   *Why it matters*: Blocks fully automated plugin workflows.

2. **#1973 – Tool whitelist for Interactive Mode**  
   [Link](https://github.com/github/copilot-cli/issues/1973)  
   A heavily upvoted feature request (20 👍) to allow safe read-only tools (grep, cat, git log) without manual approval, without resorting to `/allow-all`. 10 comments.  
   *Why it matters*: Reduces friction in everyday CLI usage; community demand is clear.

3. **#254 – Copilot CLI keeps asking to log in again**  
   [Link](https://github.com/github/copilot-cli/issues/254)  
   Persistent login loop bug reported over 8 months ago, still open. 9 comments, 4 👍.  
   *Why it matters*: A core authentication issue that frustrates long-running sessions.

4. **#3560 – Duplicate item ID error (`CAPIError: 400`)**  
   [Link](https://github.com/github/copilot-cli/issues/3560)  
   Sudden failures after tool calls with duplicate `fc_call_*` IDs. Plain chat works; only agentic turns break. 5 comments.  
   *Why it matters*: Indicates a backend race condition or caching bug.

5. **#3832 – All models show “Blocked/Disabled” after June 16 outage**  
   [Link](https://github.com/github/copilot-cli/issues/3832)  
   Closed quickly after the outage but caused widespread session lockout. 5 comments, 13 👍.  
   *Why it matters*: Highlights fragility of model selection after service disruptions.

6. **#3831 – “Request failed due to a transient API error” retry loop**  
   [Link](https://github.com/github/copilot-cli/issues/3831)  
   Users stuck in an infinite retry loop mid-workflow. 4 comments.  
   *Why it matters*: Blocks all productivity until session restart; no clear resolution path.

7. **#3355 – Claude Opus 4.6 capped at 200K tokens (native 1M)**  
   [Link](https://github.com/github/copilot-cli/issues/3355)  
   Request to allow configurable context window so deep technical sessions don’t force compaction. 3 comments, 4 👍.  
   *Why it matters*: Power users need larger context for long agentic sessions.

8. **#3730 – Support Enterprise-Managed Custom Models in Copilot CLI**  
   [Link](https://github.com/github/copilot-cli/issues/3730)  
   Custom models configured via GitHub Enterprise Admin dashboard are not visible in CLI. 2 comments, 4 👍.  
   *Why it matters*: Enterprise customers expect parity with VS Code Copilot.

9. **#3754 – `copilot --resume "Name With Spaces"` fails silently**  
   [Link](https://github.com/github/copilot-cli/issues/3754)  
   Session names containing spaces cause exit code 1 with no output. 2 comments.  
   *Why it matters*: A basic usability bug that contradicts documented behavior.

10. **#3812 – Subagents can no longer access MCP tools**  
    [Link](https://github.com/github/copilot-cli/issues/3812)  
    Custom subagents lost visibility of MCP tools; top-level agent still works. Possibly caused by deferred loading. 2 comments.  
    *Why it matters*: Breaks multi-agent orchestration patterns.

---

## Key PR Progress
No pull requests were updated in the last 24 hours.

---

## Feature Request Trends

- **Refined Permission & Control**  
  Persistent `/instructions` opt-out per repo (#3840), tool whitelist for interactive mode (#1973), and custom aliases (#3844) all aim to reduce repetitive approvals.

- **Model Flexibility**  
  Configurable context window (#3355), an `/effort` command for reasoning effort (#3074), and support for enterprise-managed custom models (#3730) reflect demand for more granular model tuning.

- **Plugin & MCP Ecosystem Maturity**  
  Batch plugin update (#3830), skill-declared MCP servers (#3292), and preloading MCP tools (#3787) show community interest in a richer extension framework.

- **Session & Attachment Robustness**  
  Resume with spaces (#3754), session origin display (#3837), and graceful handling of malformed attachments (#3791) point to a need for more resilient session management.

---

## Developer Pain Points

- **Authentication instability** – Issue #254 remains open for months; users report forced re-logins multiple times per day.
- **Transient API errors** – Duplicate item IDs (#3560) and infinite retry loops (#3831) halt workflows unpredictably.
- **Opaque content exclusions** – Organization policies incorrectly applied to CLI (#3841) despite docs stating otherwise.
- **Plugin installation failures** – `core.fsmonitor` causes `copilot plugin install` to crash (#3842).
- **Sub-agent model mismatch** – Spawned sub-agents silently use a different model than the configured session model (#3824).
- **Lazy MCP tool loading** – Tools invisible to agents that don’t probe for them (#3787, #3812).
- **Session name constraints** – Spaces in session names cause silent failure (#3754).

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest – 2026-06-18

## Today's Highlights
No new releases or pull requests landed today, but the community submitted two focused feature requests. One asks for **runtime switching between Agent and Cluster execution modes** during a session, reflecting a desire for more flexible workflows. The other seeks an **option to ignore SSL certificate validation**, a common pain point for developers behind corporate proxies or antivirus software that intercepts TLS traffic. Both issues are fresh and have drawn initial attention, signaling potential areas for upcoming improvements.

## Releases
None in the last 24 hours.

## Hot Issues
Only two issues were updated in the past 24 hours. We highlight both below:

1. **#2459 – [Feature Request] Supports switching execution mode during session (Agent ↔ Cluster)**  
   - **Author:** PresentXoX  
   - **Created/Updated:** 2026-06-17  
   - **Summary:** Request to allow users to toggle between Agent and Cluster modes without restarting a session. This would enable ad‑hoc scaling of compute or switching between single‑agent and distributed execution mid‑workflow.  
   - **Why it matters:** Many advanced users run hybrid pipelines that start with local reasoning and later need to offload heavy tasks to a cluster. No comments yet, but the request taps into a growing need for dynamic resource management.  
   - **Link:** [Issue #2459](https://github.com/MoonshotAI/kimi-cli/issues/2459)

2. **#2458 – [enhancement] Add option to ignore SSL certificate**  
   - **Author:** dmorsin  
   - **Created/Updated:** 2026-06-17  
   - **Summary:** The author’s corporate antivirus uses man‑in‑the‑middle (MiTM) to inspect all SSL connections, replacing the server certificate with its own. This causes login failures because the CLI verifies the certificate chain. The request is for a flag to disable certificate validation (e.g., `--insecure` or an environment variable).  
   - **Why it matters:** While security purists may object, real‑world environments (corporate proxies, legacy systems, test labs) frequently require this escape hatch. The single 👍 on a first‑day issue suggests others share the pain.  
   - **Link:** [Issue #2458](https://github.com/MoonshotAI/kimi-cli/issues/2458)

## Key PR Progress
None in the last 24 hours.

## Feature Request Trends
Based on the last day’s activity, two distinct trend lines emerge:

- **Session flexibility**: Users want to change operational modes (Agent vs. Cluster) without tearing down and restarting a session. This implies a demand for **stateful session management** where resource allocations can be adjusted on‑the‑fly.
- **Environment compatibility**: The request to bypass SSL validation highlights the tension between security defaults and real‑world deployment constraints. Expect more requests for **certificate pinning, custom CA bundles, and `--insecure` flags** as corporate usage grows.

## Developer Pain Points
The SSL certificate issue (#2458) crystallizes a recurring frustration: **platform‑enforced traffic inspection** (via antivirus, proxies, or endpoint protection) breaks CLI tools that strictly validate certificates. Developers in managed environments often lack the ability to disable or override these controls on their machines, forcing them to either patch the tool locally or seek a built‑in workaround.

The execution‑mode switching request (#2459) points to another pain point: **lack of runtime configurability**. Users who start a session in one mode and later decide they need the other are forced to restart, losing context and intermediate state. A solution would reduce friction in multi‑stage workflows.

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-18

## Today's Highlights
**v1.17.8** ships with faster session timelines and critical provider compatibility fixes. The community is most active around **token-per-second metrics** (#6096, 55 👍) and **multi-agent orchestration** (#17994, 21 comments), while several quality-of-life PRs (fuzzy matching, plugin hook fixes, session picker) are gaining traction.

## Releases
**v1.17.8** ([changelog](https://github.com/anomalyco/opencode/releases/tag/v1.17.8))  
- **Core improvements**: Session timelines load much faster, eliminating flicker and scroll jumps.  
- **Bug fixes**:  
  - OpenAI-compatible providers now accept MCP tool schemas that previously failed validation. (@jquense)  
  - Cloudflare AI Gateway correctly receives the configured API key. (@keefetang)

## Hot Issues (10 noteworthy)
1. **#6096 – Tokens per second display** – 55 👍, 18 comments.  
   Community strongly wants a real-time TPS indicator per message.  
   [Issue](https://github.com/anomalyco/opencode/issues/6096)

2. **#17994 – Multi-agent orchestration in isolated workspaces** – 21 comments, 2 👍.  
   Request for built-in “team of agents” support; long-running discussion.  
   [Issue](https://github.com/anomalyco/opencode/issues/17994)

3. **#23566 – LSP docs suggest it’s enabled by default but it’s not** – 10 comments, 20 👍.  
   Confusion between docs and actual default; many users affected.  
   [Issue](https://github.com/anomalyco/opencode/issues/23566)

4. **#31204 – `session_message.seq` NOT NULL constraint crash on agent switch** – 7 comments, 3 👍.  
   Database migration regression breaks sessions that switch agents.  
   [Issue](https://github.com/anomalyco/opencode/issues/31204)

5. **#24817 – Ctrl+Z suspends OpenCode on Linux instead of undoing text** – 5 comments, 2 👍.  
   SIGTSTP conflict with expected undo behavior.  
   [Issue](https://github.com/anomalyco/opencode/issues/24817)

6. **#7928 – Runtime permission mode toggle (like Claude Code’s Shift+Tab)** – 5 comments, 17 👍.  
   Many users want on-the-fly switching between auto-edit and confirm-every-action.  
   [Issue](https://github.com/anomalyco/opencode/issues/7928)

7. **#31119 – Error: no such column: name** – 4 comments, 5 👍.  
   Update regression preventing app usage; high severity for affected users.  
   [Issue](https://github.com/anomalyco/opencode/issues/31119)

8. **#32745 – “Authorization in progress…” forever when connecting OpenRouter** – 4 comments.  
   Desktop client OAuth flow hangs; likely provider-specific.  
   [Issue](https://github.com/anomalyco/opencode/issues/32745)

9. **#32704 – Bash tool description references Edit/Write tools even when unavailable** – 3 comments.  
   Misleading tool prompt leaks information about unavailable capabilities.  
   [Issue](https://github.com/anomalyco/opencode/issues/32704)

10. **#32749 – Explore agent wastes tokens by always spawning subagents** – 2 comments.  
    Users call out unnecessary exploration overhead for simple tasks.  
    [Issue](https://github.com/anomalyco/opencode/issues/32749)

## Key PR Progress (10 important)
1. **#32767 – Restore ESC interrupt for delegated subagent sessions** – Fixes regression; ESC can now stop subagents again.  
   [PR](https://github.com/anomalyco/opencode/pull/32767)

2. **#32766 – Accept explicit storage in public API layer** – Enables tests and embeddings to inject disposable DB storage.  
   [PR](https://github.com/anomalyco/opencode/pull/32766)

3. **#32765 – Code cleanup, formatter consolidation, and perf improvements** – Removes dead code, merges Ruff/uv formatter, optimizes message normalization.  
   [PR](https://github.com/anomalyco/opencode/pull/32765)

4. **#32762 – Prevent recursive sub-skill discovery** – Single-level glob stops nested skills from being loaded as independent skills.  
   [PR](https://github.com/anomalyco/opencode/pull/32762)

5. **#32761 – Port V1 fuzzy edit matching to V2 core edit tool** – 9 fuzzy strategies + Levenshtein distance improve edit reliability for LLM outputs.  
   [PR](https://github.com/anomalyco/opencode/pull/32761)

6. **#32758 – Fix plugin trigger output to support array replacement** – Plugins modifying `output.messages` are no longer silently ignored.  
   [PR](https://github.com/anomalyco/opencode/pull/32758)

7. **#32753 – Add clipboard fallback for non-HTTPS contexts** – Fixes copy button in OpenCode Web on localhost.  
   [PR](https://github.com/anomalyco/opencode/pull/32753)

8. **#32752 – Interactive `session select` picker** – New CLI command to list and choose sessions via `@clack/prompts`.  
   [PR](https://github.com/anomalyco/opencode/pull/32752)

9. **#30879 – Improve display and replay of shell commands** – Real-time output streaming and better command titles in ACP mode.  
   [PR](https://github.com/anomalyco/opencode/pull/30879)

10. **#32731 – Auto-discover models from OpenAI-compatible providers** – Calls `GET /models` to populate model list automatically.  
    [PR](https://github.com/anomalyco/opencode/pull/32731)

## Feature Request Trends
- **Multi-agent orchestration** (#17994) and **isolated workspaces** remain the highest-voted long-term feature.  
- **Token-per-second metrics** (#6096) and **runtime permission toggles** (#7928) gain frequent upvotes.  
- **Local model support** (Ollama, #32756) and **new model integration** (GLM-5.2 for Z.AI, #32172 and Ollama, #32620) are recurring themes.  
- Session management improvements (global/local scope toggle, interactive picker) appear in both issues and PRs.

## Developer Pain Points
- **Windows ANSI escape code corruption** (#21277, #16675, #32754) persists across versions and terminals.  
- **Database migration breakage** (#31204, #31119) causes crashes after updates – high urgency.  
- **Keybinding conflicts** (Ctrl+Z suspends on Linux, Ctrl+X vs Ctrl+C close danger) frustrate daily use.  
- **OAuth authorization stalls** (OpenRouter, #32745) leave users stuck without feedback.  
- **Edit tool reliability** – LLM-produced `oldString` mismatches require fuzzy matching (#32760), which the V2 port (#32761) addresses.  
- **Explore agent overhead** (#32749) wastes tokens for trivial tasks, calling for smarter task sizing.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest – 2026-06-18

## Today’s Highlights
A notable streaming UX regression (#5825) is being actively fixed: the TUI now forces scroll-to-bottom when “clear on shrink” is enabled, making it impossible to read ahead. Meanwhile, critical infrastructure work continues to move off Shrinkwrap (#5653) to eliminate duplicate dependency copies, and a batch of compaction and provider error-handling PRs landed. The week saw strong community demand for XDG compliance on Linux and for extending Pi’s subscription model to Anthropic OAuth.

## Releases
No new versions were published in the last 24 hours.

## Hot Issues (10 items)

1. **[#5825 – Streaming markdown forces scroll to bottom](https://github.com/earendil-works/pi/issues/5825)**  
   *Open / in-progress* · 12 comments  
   A major UX regression: when “clear on shrink” is enabled, every new token forces the terminal to scroll to the bottom, making reading while the agent talks impossible. PR #5846 marks this as fixed in the next release.

2. **[#5653 – Move off Shrinkwrap](https://github.com/earendil-works/pi/issues/5653)**  
   *Open / in-progress* · 11 comments  
   Installing both `pi-ai` and `pi-coding-agent` as direct dependencies puts two identical copies of `pi-ai` on disk, causing singleton registries to break. The community considers this a blocker for complex setups.

3. **[#3715 – local-llm streams terminate at 5 min from undici default `bodyTimeout`](https://github.com/earendil-works/pi/issues/3715)**  
   *Closed* · 11 comments · 4 👍  
   Long-running `Write` tool calls against local backends (e.g., vLLM) are killed after 5 minutes because `retry.provider.timeoutMs` cannot override undici’s internal timeout. Still affecting users with large write operations.

4. **[#534 – Config folder is out of place on Linux](https://github.com/earendil-works/pi/issues/534)**  
   *Closed* · 9 comments · 20 👍  
   Pi stores configuration directly in `~/.pi/` instead of following the XDG Base Directory spec. The high number of upvotes shows this is a long-standing pain point for Linux users.

5. **[#5654 – Add `excludeFromContext` to custom messages](https://github.com/earendil-works/pi/issues/5654)**  
   *Open* · 7 comments  
   Developers building extensions want to send custom display messages (e.g., status reports) that are never included in the LLM context. Mirroring the existing `!!` bash execution flag would enable this cleanly.

6. **[#5821 – Support Anthropic OAuth Subscription Usage](https://github.com/earendil-works/pi/issues/5821)**  
   *Closed* · 7 comments  
   Anthropic confirmed Agent SDK apps can still use direct subscriptions. The community is eager to avoid a separate credit system for Pi integration.

7. **[#5830 – Tree navigator truncates long entries](https://github.com/earendil-works/pi/issues/5830)**  
   *Closed* · 4 comments  
   The `/tree` view truncates entries at terminal width with no way to see the full content – a bad UX when sessions contain deeply nested structures.

8. **[#5810 – RPC: expose session entries and tree](https://github.com/earendil-works/pi/issues/5810)**  
   *Open* · 3 comments  
   A request for two read-only RPC commands (`get_entries`, `get_tree`) so that external tools can drive Pi sessions programmatically.

9. **[#4973 – Regression: prompt templates collapse multi-line arguments](https://github.com/earendil-works/pi/issues/4973)**  
   *Closed* · 3 comments · 2 👍  
   Prompt templates using `$@` or `$ARGUMENTS` now flatten newlines into spaces, breaking multi-line input pass-through. A clear regression that affects template authors.

10. **[#3200 – Support video/audio content in prompt command](https://github.com/earendil-works/pi/issues/3200)**  
    *Open* · 3 comments  
    Extends the RPC `prompt` command to send video/audio alongside images, enabling multimodal models (Gemma 4, GPT-4o) to process richer inputs.

## Key PR Progress (10 items)

1. **[#5846 – fix(tui): stabilize streaming code fence rendering](https://github.com/earendil-works/pi/pull/5846)**  
   *Open* · Directly closes #5825. Prevents the forced scroll-to-bottom bug by properly handling code fences during streaming.

2. **[#5859 – fix(ai): send responses prompts as instructions](https://github.com/earendil-works/pi/pull/5859)**  
   *Open* · Fixes OpenAI Responses API compliance by moving system prompts into the `instructions` top-level field instead of replaying them as `input` messages.

3. **[#5849 – feat(ai): add Azure AI Foundry provider for Anthropic Claude](https://github.com/earendil-works/pi/pull/5849)**  
   *Closed* · Introduces a new `azure-foundry` provider with full URL shape, headers, and Entra ID auth – matching Python SDK parity.

4. **[#5832 – fix(ai): surface provider HTTP error body](https://github.com/earendil-works/pi/pull/5832)**  
   *Open* · Behind proxies, opaque SDK error messages (e.g., “UnknownError”) now show the raw HTTP body, making debugging much easier.

5. **[#5829 – feat: add "max" thinking level for adaptive reasoning models](https://github.com/earendil-works/pi/pull/5829)**  
   *Closed* · Extends `ThinkingLevel` with `max` for Anthropic models like Claude Opus 4.8 that support a fifth reasoning effort tier.

6. **[#5833 – Compaction-related fixes](https://github.com/earendil-works/pi/pull/5833)**  
   *Closed* · Three fixes for the compaction mechanism: reordering sub-conversations, handling zero-length last line, and respecting `maxTokens` more accurately.

7. **[#5801 – Nixify pi](https://github.com/earendil-works/pi/pull/5801)**  
   *Closed* · Adds Nix flake packaging so Pi can be built and run with `nix run` – great for NixOS users.

8. **[#5812 – fix(tui): protect pipe characters inside inline code in markdown tables](https://github.com/earendil-works/pi/pull/5812)**  
   *Closed* · Prevents markdown table renderer from splitting on `|` inside backticks, preserving cell content.

9. **[#5850 – chore(deps): bump vitest to 3.2.6](https://github.com/earendil-works/pi/pull/5850)**  
   *Closed* · Mechanical dependency bump that resolves 5 of 6 high‑severity `npm audit` advisories in dev dependencies.

10. **[#631 – fix(ai): Google thinking detection](https://github.com/earendil-works/pi/pull/631)**  
    *Closed* · Correctly identifies Google´s `thoughtSignature` as thinking content and removes unsupported `id` fields that caused schema errors.

## Feature Request Trends

- **Extended context window management** – Several issues and PRs (#5848, #5768, #5692) demand percentage‑based compaction triggers, 1M token support for Copilot/GLM models, and configurable auto‑compaction thresholds.
- **Multimodal expansion** – Issue #3200 and the ongoing `prompt` RPC enhancements show growing interest in sending video/audio alongside images to multimodal LLMs.
- **Subscription and provider diversity** – Requests for Anthropic OAuth subscription support (#5821), Azure AI Foundry (#5849, #5851), and Codex subscription parity (#5862) reflect a desire to use Pi with existing paid plans.
- **Extension API maturation** – Community wants finer‑grained control: `excludeFromContext` on custom messages (#5654), executable tool objects (#5781), and command execution events (#5608).
- **New model onboarding** – Frequent requests for GLM‑5.2 effort levels (#5770), Copilot model thinking modes (#5768), and Opus 4.8 `max` reasoning (#5829) indicate Pi is treated as a primary client for frontier models.

## Developer Pain Points

- **Package duplication from Shrinkwrap** – #5653 remains a core complaint; two direct dependencies cause duplicate copies of `pi-ai`, breaking module‑level singletons.
- **Terminal UX friction** – Forced scroll-to-bottom (#5825), tree navigator truncation (#5830), and pipe characters breaking markdown tables (#5812) degrade the interactive experience.
- **Provider timeout limitations** – Undici’s hard‑coded 5‑minute body timeout (#3715) and MCP HTTP server retry backoff on 401/403 (#5857) block long‑running operations and cause startup hangs.
- **Error message opacity** – Opaque SDK errors (#5832, #5828) hide real HTTP error bodies, making debugging behind proxies or gateways painful.
- **Configuration and platform compliance** – The Linux config folder not following XDG (#534) and missing CLI value diagnostics (#5856) frustrate power users and script integrators.
- **Compaction inefficiencies** – Unexplained reordering, zero‑length file handling, and inaccurate `maxTokens` accounting (#5833, #5845) reduce reliability for large sessions.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest – 2026-06-18

## Today's Highlights

The Qwen Code team shipped **v0.18.3** and several nightly builds today, fixing a critical `sed` edit tracking bug and improving CLI cancellation handling. A major **model provider disambiguation fix** (PR #5241) merged to resolve persistent issues when multiple providers share the same model ID. Meanwhile, a new **circuit breaker for runaway tool‑call loops** (PR #5279) was introduced to prevent infinite recursion scenarios, and a **QQ Bot channel adapter** (PR #5202) is under review to expand the platform’s integration reach.

---

## Releases

| Version | Key Changes |
|---------|-------------|
| **v0.18.3** (latest stable) | [`chore(release): v0.18.2`](https://github.com/QwenLM/qwen-code/pull/5243) + `fix(cli): Stop after cancelled ask_user_question` by @doudouOUC |
| **v0.18.3-nightly** | Includes `fix(core): Track supported sed edits in file history` by @doud |
| **v0.18.3-preview.0** | Same CLI fix as v0.18.3, tagged for early adopters |
| **v0.18.2** | `fix: warn on oversized context instructions` by @he-yufeng + docs fixes for stale defaults and CLI syntax |
| **v0.18.1-preview.1** | Backport of the oversized‑context warning and documentation updates |

---

## Hot Issues

*(6 issues updated in the last 24h – all listed due to limited volume)*

| Issue | Status | Summary | Why It Matters |
|-------|--------|---------|----------------|
| [#5173](https://github.com/QwenLM/qwen-code/issues/5173) | **CLOSED** | Model provider disambiguation fails when multiple providers (e.g. Token Plan, IdeaLab, BFF) share the same model id (`qwen3.7-max`). Selection doesn’t persist across sessions. | Affects users with multi‑endpoint setups; fixed in PR #5241. |
| [#5280](https://github.com/QwenLM/qwen-code/issues/5280) | OPEN | Re‑enable long command search suggestion coverage (test skipped). Community member @tt-a1i offers to take it. | Restoring test coverage improves CLI reliability; welcome PR. |
| [#5277](https://github.com/QwenLM/qwen-code/issues/5277) | OPEN | Re‑enable TableRenderer foreground reset test (skipped). Functional assertions pass; old snapshot needs update. | Low‑risk test maintenance to prevent regression in terminal rendering. |
| [#5275](https://github.com/QwenLM/qwen-code/issues/5275) | OPEN | Re‑enable BaseSelectionList scroll‑up test (skipped). React `act` warnings need wrapping. | Ensures scrolling behaviour in TUI lists remains correct. |
| [#2845](https://github.com/QwenLM/qwen-code/issues/2845) | **CLOSED** | Request: treat `.dat` files with text content as text instead of binary. User wants file‑type‑agnostic content detection. | Resolved by PR #5256 – now `.dat` files are inspected by content, not extension. |
| [#5281](https://github.com/QwenLM/qwen-code/issues/5281) | OPEN | TUI becomes unresponsive when `login1.inhibit-block-sleep` triggers an authentication prompt. Affects SSH sessions on Linux without DE login. | Critical for headless/remote usage; introduced by sleep‑prevention feature. Requires triage. |

---

## Key PR Progress

*(10 most impactful pull requests among the 50 updated in the last 24h)*

| PR | Status | Description | Significance |
|----|--------|-------------|--------------|
| [#5202](https://github.com/QwenLM/qwen-code/pull/5202) | OPEN | Adds `@qwen-code/channel-qqbot` – a QQ Bot adapter (WebSocket gateway, heartbeat, reconnect). | Expands platform integrations beyond Telegram/WeChat/Feishu; strong community interest. |
| [#5181](https://github.com/QwenLM/qwen-code/pull/5181) | **in‑review** | Fixes OOM crash on `/quit` due to `buildTranscriptMessages()` processing full history with regex. | Prevents `FATAL ERROR: Reached heap limit` during auto‑memory extraction. |
| [#5279](https://github.com/QwenLM/qwen-code/pull/5279) | OPEN | Adds always‑on tool‑call circuit breaker to stop runaway loops (re‑scoped from #5242). | Safety net for infinite tool‑calling; lightweight and focused. |
| [#5241](https://github.com/QwenLM/qwen-code/pull/5241) | **CLOSED** | Disambiguates providers sharing a model ID by persisting `baseUrl` selection across sessions. | Directly solves #5173 – restores multi‑provider setups. |
| [#5231](https://github.com/QwenLM/qwen-code/pull/5231) | OPEN | Workflow tool token budget + per‑run UI surfacing (banner, background dialog phase tree). | Improves transparency of token consumption in workflow tasks. |
| [#5145](https://github.com/QwenLM/qwen-code/pull/5145) | OPEN | Shows follow‑up suggestion in input placeholder after model response (generated by fast model). | Enhances user flow – immediate next‑step suggestion without extra UI element. |
| [#5256](https://github.com/QwenLM/qwen-code/pull/5256) | **CLOSED** | Detects `.dat` files by content instead of extension (fixes #2845). | Removes false‑positive binary classification for text‑based `.dat` files. |
| [#5258](https://github.com/QwenLM/qwen-code/pull/5258) | OPEN | Stops current turn on any permission cancellation, not just `ask_user_question`. | Consistent cancellation behaviour across all ACP tools. |
| [#5284](https://github.com/QwenLM/qwen-code/pull/5284) | OPEN | Compiles macOS 26+ Liquid Glass `Assets.car` automatically in `brand-create`. | Future‑proofs macOS desktop builds for upcoming OS version. |
| [#5283](https://github.com/QwenLM/qwen-code/pull/5283) | OPEN | Re‑enables long command search suggestion coverage (fixes #5280) with updated snapshots. | Clears long‑standing skipped test for input prompt. |

---

## Feature Request Trends

- **Multi‑provider model selection** – Users increasingly run the same model ID (e.g. `qwen3.7-max`) across different providers; persistent provider selection by `baseUrl` is now resolved (#5173, #5241).
- **File‑type‑agnostic content detection** – The community wants binary/text detection to ignore file extensions (#2845, #5256), now implemented.
- **Workflow token budgeting** – Demand for visibility and limits on token consumption during workflow runs (#5231).
- **Enhanced TUI input** – Follow‑up suggestions in placeholders (#5145) and Ctrl‑P/N navigation in completions (#5259) show appetite for smoother interactive editing.
- **Channel expansion** – The QQ Bot adapter (#5202) signals a trend toward covering Asian‑market messaging platforms (WeChat already supported).
- **MacOS forward‑compatibility** – Proactive compilation of Liquid Glass assets (#5284) indicates a desire to stay ahead of OS changes.

---

## Developer Pain Points

1. **TUI unresponsiveness due to system sleep authentication** (#5281) – The sleep‑prevention feature triggers polkit/dbus password prompts that hijack the TUI input stream, making the tool unusable in SSH‑only Linux sessions. Workaround: disable the setting in `/settings`.
2. **Test coverage gaps** – Several skipped UI tests (command search, TableRenderer, BaseSelectionList) have been lingering for weeks (#5280, #5277, #5275). While the underlying behaviour works, the lack of automated regression coverage is a maintenance risk.
3. **OOM on `/quit` with large histories** (#5181) – Auto‑memory extraction caused heap exhaustion during conversation shutdown, particularly for long‑running sessions. Fixed in PR #5181.
4. **Oversized context warnings** (#5073) – Users hitting token limits needed clearer feedback; now a warning is shown, reducing silent failures.
5. **CLI termination after permission cancellation** – Previously only `ask_user_question` cancellation stopped the turn; other permission cancellations left the tool hanging (#5258). Now unified.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest | 2026-06-18

## Today's Highlights
The community is heavily focused on stabilizing the **v0.9.0** roadmap, with major PRs landing for the chat-native "Workrooms" architecture and a slew of fixes for Plan/Agent mode toggles, config file preservation, and self-questioning agent loops. Developers are also pushing for better **static Linux builds** (musl), **snapshot config compliance**, and **UI performance** improvements. A recurring theme is the tension between agent autonomy and user control, prompting multiple issues and PRs around permission enforcement.

## Releases
No new releases were published in the last 24 hours.

## Hot Issues (10 of 11)

1. **[#2870 – EPIC: staged command-boundary refactor for #2791](https://github.com/Hmbown/CodeWhale/issues/2870)**  
   *Open, documentation/tui* – Tracks the incremental mergeable pieces of a large command-boundary refactor targeted for v0.9.0. Five comments, low engagement, but foundational for the upcoming release.

2. **[#3275 – CodeWhale overly involved in self-questioning, deviating from user intent](https://github.com/Hmbown/CodeWhale/issues/3275)**  
   *Open, bug* – A regression from #3061. The agent enters a self-sustaining loop of proposing, answering, and executing without waiting for approval. **4 comments**, strong community concern about agent autonomy.

3. **[#3279 – Plan/Agent Mode Toggle Inconsistency & Tool Permission Chaos](https://github.com/Hmbown/CodeWhale/issues/3279)**  
   *Open, bug/ux* – Switching from Plan to Agent leaves `write_file`/`exec_shell` denied, and after recovery the AI auto-executes plans without consent. **3 comments**, high frustration over broken mode logic.

4. **[#3289 – v0.8.61 UI froze after auto-spawning several agents](https://github.com/Hmbown/CodeWhale/issues/3289)**  
   *Open, bug* – Reproducible freeze when the tool spawns multiple agents automatically. **2 comments**, likely blocking users with complex workflows.

5. **[#3281 – [moonshot] v0.8.61 fix incomplete — parameters still missing `type:object` for root schemas](https://github.com/Hmbown/CodeWhale/issues/3281)**  
   *Open, bug* – The hotfix for Kimi/Moonshot schema rejection only covers a narrow set of conditions; `$ref`/`anyOf`/`allOf` root schemas remain broken. **2 comments**, impacts users of those providers.

6. **[#3292 – `pre_tool_snapshot` did not respect config `snapshots.enabled=false`](https://github.com/Hmbown/CodeWhale/issues/3292)**  
   *Open, bug* – Snapshotting copies the entire repo even when disabled, consuming GBs of disk space. **1 comment**, reported as a silent resource hog.

7. **[#1481 – Support OpenCode Go/Zen (DeepSeek-V4)](https://github.com/Hmbown/CodeWhale/issues/1481)**  
   *Open, enhancement* – Request to add a cheap DeepSeek-V4 provider. **2 comments, 1 👍** – moderate interest, competitive pricing.

8. **[#3209 – v0.9.0 EPIC: Chat-native CodeWhale workrooms](https://github.com/Hmbown/CodeWhale/issues/3209)**  
   *Open, documentation/enhancement* – Proposes threads, channels, shareable links, multi-agent collaboration. **2 comments** – the headline feature for the next major release.

9. **[#3282 – config.toml file content improvement](https://github.com/Hmbown/CodeWhale/issues/3282)**  
   *Open, enhancement* – TUI edits erase all commented-out lines and notes. **1 comment** – a quality-of-life request that resonates with config-heavy users.

10. **[#2917 – Cargo distribution: error after rename from `deepseek-tui` to `codewhale`](https://github.com/Hmbown/CodeWhale/issues/2917)**  
    *Closed, bug* – Upgrade path broke because PATH points to old binary name. **4 comments** – already resolved but caused confusion during transition.

## Key PR Progress (10 of 25)

1. **[#3295 – feat(tui): honor ask permission rules at runtime](https://github.com/Hmbown/CodeWhale/pull/3295)**  
   Open by greyfreedom – Wires `permissions.toml` ask-only rules into the TUI runtime approval path, giving users fine-grained control over tool execution.

2. **[#3277 – feat: implement Workrooms Phase 1](https://github.com/Hmbown/CodeWhale/pull/3277)**  
   Open by idling11 – Introduces the data model, RFC documentation, and endpoints for the v0.9.0 chat-native workroom abstraction. A 242-line design RFC included.

3. **[#3283 – Fix: Plan/Agent Mode Toggle — approval_mode restore + auto-execution guard](https://github.com/Hmbown/CodeWhale/pull/3283)**  
   Open by idling11 – Directly addresses #3279 by ensuring `approval_mode` is restored after mode switch and preventing auto-execution of plans without user confirmation.

4. **[#3284 – perf(tui): debounce thinking-stream re-renders (#1620)](https://github.com/Hmbown/CodeWhale/pull/3284)**  
   Open by LeoLin990405 – Fixes painfully slow reasoning block rendering by debouncing cell revision bumps, vastly improving responsiveness for fast-thinking models.

5. **[#3285 – fix(tui): persist session before stall/cancel recovery](https://github.com/Hmbown/CodeWhale/pull/3285)**  
   Open by LeoLin990405 – Prevents data loss when a long turn is stalled or cancelled; ensures `--continue` loads the full conversation history.

6. **[#3290 – fix(prompts): add scope_discipline rules to prevent self-questioning agent loops](https://github.com/Hmbown/CodeWhale/pull/3290)**  
   Open by yekern – Adds 47 lines to the constitution prompt to curb the self-questioning-and-self-answering behaviour reported in #3273.

7. **[#3291 – Fix/preserve comments in config files](https://github.com/Hmbown/CodeWhale/pull/3291)**  
   Open by zlh124 – Uses `toml_edit` to merge user comments back into `config.toml`, `settings.toml`, and `tui.toml` after any write operation. Solves #3282.

8. **[#3293 – fix(tui): respect snapshots.enabled for per-tool snapshots (#3292)](https://github.com/Hmbown/CodeWhale/pull/3293)**  
   Open by wuisabel-gif – Adds the missing `snapshots.enabled` guard to the per-tool snapshot commit site, preventing unwanted disk usage.

9. **[#3294 – fix(tui): write composer history under `.codewhale`, not legacy `.deepseek`](https://github.com/Hmbown/CodeWhale/pull/3294)**  
   Open by wuisabel-gif – Stops recreating the legacy `~/.deepseek/` directory on fresh installs by using the modern `.codewhale` path.

10. **[#3274 – feat(release): build static Linux x64 binaries with musl](https://github.com/Hmbown/CodeWhale/pull/3274)**  
    Open by wavezhang – Switches release builds from dynamic glibc to static musl, improving portability across Linux distributions.

## Feature Request Trends

- **Multi-provider expansions** – Ongoing demand for cheap alternatives like OpenCode Go/Zen (DeepSeek-V4) and Atlas Cloud, alongside better model auto-routing fallback (PR #3280).
- **Session continuity and workrooms** – Non-interactive `--resume` / `--session-id` mode (#1530) and the ambitious chat-native Workrooms (#3209) signal a desire for persistent, multi-turn agent collaboration.
- **Configuration UX improvements** – Preserving comments in TUI-edited config files (#3282, PR #3291) and surfacing experimental feature flags in the config view (#3177) are repeated requests.
- **Agent control and permission granularity** – Users want both strict mode toggles (Plan vs Agent) and configurable permission files (#3295), plus guardrails against self-looping (#3290).

## Developer Pain Points

- **Config erosion** – TUI edits silently deleting comments frustrates users who rely on annotations and commented-out stanzas (#3282).
- **Agent autonomy gone rogue** – The self-questioning loop (#3275) and Plan/Agent permission chaos (#3279) undermine trust in the tool’s behaviour.
- **Snapshot overhead** – `snapshots.enabled=false` being ignored (#3292) wastes disk space and surprises users.
- **UI freezes and performance** – Auto-spawning multiple agents causes freezes (#3289); thinking-stream rendering is painfully slow (#3284).
- **Incomplete provider patches** – The Kimi/Moonshot `type:object` fix (#3281) missed common schema patterns, indicating rushed, narrow hotfixes.
- **Migration pain** – The rename from `deepseek-tui` to `codewhale` left broken PATH references (#2917), though now closed.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*