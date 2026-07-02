# AI CLI Tools Community Digest 2026-07-02

> Generated: 2026-07-02 10:17 UTC | Tools covered: 9

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

# Cross-Tool AI CLI Comparison Report — 2026-07-02

## 1. Ecosystem Overview

The AI CLI tool ecosystem has entered a phase of **intense platform competition** where each major vendor is racing to differentiate on security, automation, and extensibility. Today's data reveals three dominant themes: **reliability crises** (rate-limit bugs, silent timeouts, and infinite loops across Claude Code, OpenCode, and Gemini CLI), **security hardening** (symbolic-link escapes in Gemini CLI, sandbox filter blocking in Codex, and permission gating in Qwen Code), and **a systemic shift toward background automation**—multiple tools (Qwen Code, Pi, DeepSeek TUI) are adding daemonized scheduling, cron jobs, and autonomous subagent workflows. While Claude Code and Codex remain the most feature-rich, they are increasingly **challenged by open-source and smaller tools** like OpenCode and Pi that ship faster and address platform parity more aggressively. The community is clearly demanding **reliable, secure, and autonomous agents** that work cross-platform without configuration debt.

---

## 2. Activity Comparison

| Tool | Hot Issues (today) | PRs (last 24h) | Release (last 24h) | Community Momentum |
|---|---|---|---|---|
| **Claude Code** | 10 | 2 (1 trivial) | v2.1.198 ✅ | Moderate; low PR activity suggests internal pipelines dominate |
| **OpenAI Codex** | 10 | 10 (all substantive) | v0.143.0-alpha.33 ✅ | High; strong security hardening, active feature development |
| **Gemini CLI** | 10 | 10 (all substantive) | v0.51.0-nightly ✅ | High; rapid iteration, good mix of bugfixes and features |
| **GitHub Copilot CLI** | 10 | 0 | v1.0.69-0 ✅ | Moderate; feature releases are small, bug backlog growing |
| **Kimi Code CLI** | 3 | 1 | None ❌ | Low; quietest of the cohort, possible stabilization phase |
| **OpenCode** | 10 | 10 (all substantive) | v1.17.13 ✅ | High; service incident drove major activity, rapid response |
| **Pi** | 10 | 10 (all merged) | None ❌ | High; healthy community, strong Linux/WSL focus |
| **Qwen Code** | 10 | 10 (all substantive) | v0.19.4 + nightly ✅ | High; aggressive feature cadence, background automation push |
| **DeepSeek TUI (CodeWhale)** | 10 | 10 (all substantive) | None ❌ | High; massive setup wizard cycle, dynamic MCP capability |

**Key observations:**
- **Codex, Gemini, OpenCode, Pi, Qwen Code, and DeepSeek TUI** each had **10+ substantial PRs** in 24 hours—indicating very high development velocity.
- **Claude Code** and **Copilot CLI** have lower visible PR activity, likely due to internal/pipeline-driven workflows.
- **Kimi Code CLI** stands out as the quietest, with only 3 issues and 1 PR—may indicate a slower development cycle or resource constraints.
- **OpenCode** had the most reactive community today due to a Go-tier service outage.

---

## 3. Shared Feature Directions

### Cross-cutting requirements emerging across **4+ tools**:

| Requirement | Affected Tools | Specific Need |
|---|---|---|
| **Linux desktop builds** | Claude Code (495 👍), Gemini CLI, Pi (XDG compliance) | Native Linux support for desktop/TUI interfaces |
| **Custom model endpoints** | Copilot CLI (#4003), OpenCode (#34889), Pi (local models) | Ability to use private, local, or non-standard model endpoints |
| **Background/scheduled automation** | Qwen Code (cron daemon), Pi (AOT compilation), DeepSeek TUI (dynamic MCP), Claude Code (background agents) | Persistent agents that run autonomously without user session |
| **Granular permission/policy control** | Claude Code (project guardrails), Copilot CLI (deny-rules), Qwen Code (Tool(param:value)), DeepSeek TUI (constitution) | Fine-grained security policies beyond simple allow/block |
| **Context management improvements** | Claude Code (compaction thrashing), Codex (custom compaction hooks), Gemini CLI (auto-memory retries), Pi (SQLite storage) | Better token budgeting, user-controlled compaction, context visibility |
| **Cross-platform parity (Windows)** | Claude Code (bug #72149), Codex (5+ Windows issues), Copilot CLI (#4001), Pi (Windows drive root) | Missing commands, auth hooks, sandbox setup on Windows |

### Notable 3-tool convergences:
- **MCP server spawning** – DeepSeek TUI (dynamic MCP from LLM), Qwen Code (MCP in YOLO mode), Codex (prototype HTTP MCP)
- **Usage/quota transparency** – Codex (#29963), OpenCode (#34884), Copilot CLI (#3994) – users can't trust billing displays
- **Subagent lifecycle visibility** – Gemini CLI (#22598), DeepSeek TUI (#3837), Qwen Code (#6180) – subagent state not reflected in UI

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | Copilot CLI | OpenCode | Pi | Qwen Code | DeepSeek TUI |
|---|---|---|---|---|---|---|---|---|
| **Primary target** | Enterprise devs | Pro/paid users | Google ecosystem | GitHub ecosystem | Price-sensitive / open-source | Hackers / extensibility | Chinese enterprise + automation | Power users / security-first |
| **Key differentiator** | Chrome integration, `/dataviz` | Security sandbox hardening | AST-aware tools, subagent model | Tight IDE integration, plugins | Multi-provider flexibility | TypeScript extension API, local-first | Background daemon + scheduling | Constitution-driven security posture |
| **Security posture** | Mixed (silent timeout bug) | Strong (filter blocking, sandbox escapes) | Good (symlink fix, thought leakage fix) | Moderate (plugin MCP holes) | Moderate (rate-limit bugs) | Good (AOT compilation, session integrity) | Emerging (npm audit, scanner flags) | Strong (constitution, dynamic MCP gating) |
| **Platform maturity** | Mac-focused, Linux gaps | Mac-focused, Windows gaps | Strong on Linux | Windows gaps | Cross-platform | Linux/WSL focus | Cross-platform (mobile included) | Cross-platform (Windows bugs) |
| **Extensibility model** | MCP + skills | MCP + plugins | Skills + subagents | Plugin marketplace | Custom providers | TypeScript extensions + MCP | Channel adapters + MCP | Dynamic MCP + constitution |
| **Release cadence** | Weekly | Weekly+alphas | Nightly | Twice weekly | Weekly | Weekly | Nightly + stable | No daily release |

**Strategic takeaways:**
- **Claude Code** leads in **ecosystem integration** (Chrome, desktop, `/dataviz`) but lags on security fundamentals (the 60s timeout bug is a serious liability).
- **OpenAI Codex** is the **security leader**—aggressively patching sandbox escapes, filter execution paths, and supply chain vectors. This positions it for enterprise compliance.
- **Gemini CLI** differentiates on **agent orchestration** (subagents, AST tools, memory system) but suffers from reliability issues (false success reports, hangs).
- **Qwen Code and DeepSeek TUI** are racing toward **autonomous background agents**—a direction no other tool has fully committed to yet.
- **Pi** occupies a unique niche: **extensible, local-first, and community-driven**—attracting developers who want control over their toolchain.
- **Kimi Code CLI** is **dangerously quiet**—risk of losing developer mindshare if velocity doesn't pick up.

---

## 5. Community Momentum & Maturity

### High-velocity (10+ substantive PRs/day, rapid issue response):
- **OpenAI Codex** – Mature project, enterprise-grade engineering, strong security focus.
- **Gemini CLI** – Very active, good balance of bugfixes and features, growing community (nightly releases).
- **OpenCode** – Highly reactive to incidents (today's Go outage), strong community engagement.
- **Pi** – Healthy, focused community; strong Linux advocacy, good maintainer responsiveness.
- **Qwen Code** – Extremely active, ambitious feature roadmap (background automation, channel adapters).
- **DeepSeek TUI (CodeWhale)** – Maintainer-led with strong vision (constitution, dynamic MCP), active contributor base.

### Moderate velocity (fewer visible PRs, but meaningful releases):
- **Claude Code** – Low PR visibility suggests internal development; large community (495 👍 for Linux) indicates high interest but slower external contribution loop.
- **GitHub Copilot CLI** – Small feature releases, growing bug backlog; community is vocal but PRs are absent.

### Low velocity:
- **Kimi Code CLI** – Zero releases, 1 PR, minimal issue activity. This is a **watch-out signal**—the project may be under-resourced or paused.

### Community maturity indicators:
- **Most vocal communities** (by issue engagement): Claude Code (495 👍 on one issue), Codex (415 👍), OpenCode (22 👍 on incident).
- **Most developer frustration** (by recurring pain points): OpenCode (rate limit), Gemini CLI (false subagent successes), Claude Code (security timeout).
- **Most collaborative communities** (by PRs from contributors): Pi (many 3rd-party PRs merged), DeepSeek TUI, Qwen Code.

---

## 6. Trend Signals

### 🔴 Critical (Urgent for all tool maintainers):

1. **Security is now table stakes** – The Gemini symlink escape, Codex sandbox bypass, and Claude Code silent timeout prove that security bugs are being weaponized. Tools without robust permission systems (OpenCode's rate-limit bugs, Copilot CLI's plugin MCP holes) will lose enterprise trust.

2. **Reliability trumps features** – The silent timeout (#73125) and infinite loops (#3158, #640) are the most dangerous bugs because they destroy trust in automation. Developers are vocal: they want predictable, not flashy.

3. **Windows remains an afterthought** – Every tool except Pi has significant Windows bugs (missing commands, auth failures, sandbox setup). This is a massive opportunity for any tool that gets Windows right.

### 🟡 Emerging (Shaping the next 3-6 months):

4. **Background agents are the new frontier** – Qwen Code's `/schedule` daemon, DeepSeek TUI's dynamic MCP, and Claude Code's background agent notifications all point toward tools becoming **persistent infrastructure**, not just interactive assistants. This will require new reliability patterns (heartbeat, crash recovery, queue management).

5. **Context management is broken everywhere** – Compaction thrashing (Claude Code), token waste (Qwen Code), and unbounded memory (Codex) are universal pain points. The tool that ships user-controlled, transparent, and efficient context management will have a significant advantage.

6. **Billing/rate-limit transparency is failing** – OpenCode, Codex, and Copilot CLI all face user anger over opaque quota consumption. Trust in AI tools requires **predictable cost accounting**.

7. **Constitution/guardrails as first-class concepts** – DeepSeek TUI's "constitution-first" approach and Claude Code's "project-level guardrails" (#72043) signal a shift from simple YOLO/auto-approve to **structured, user-defined security policies** that travel with the codebase.

### 🟢 Long-term signals (12+ months):

8. **The MCP ecosystem is fragmenting** – Every tool defines MCP differently (dynamic spawning in DeepSeek, prototype HTTP MCP in Codex, gated MCP in Qwen Code). Standardization is needed, or the ecosystem will become Balkanized.

9. **Local/offline models are a growing demand** – Pi's local model support, OpenCode's custom endpoint requests, and Kimi Code's custom Anthropic endpoint indicate a user base that wants to **decouple from vendor APIs** for privacy, cost, and latency.

10. **The "agent orchestration" category is emerging** – Subagent lifecycle management (Gemini CLI, DeepSeek TUI, Qwen Code) is becoming a distinct problem domain, demanding dedicated UI (subagent trajectories, sidebar state, cancellation workflows).

---

## Recommendations for Technical Decision-Makers

- **For enterprise deployments**: **OpenAI Codex** currently has the strongest security posture, but **Claude Code** has the richest ecosystem integration. Evaluate your compliance requirements carefully—especially the silent timeout bug in Claude Code.
- **For cross-platform teams**: **Pi** and **OpenCode** have the best cross-platform support today. **Windows-heavy teams** should avoid Copilot CLI and Codex until their Windows bugs are addressed.
- **For automation-first workflows**: **Qwen Code** and **DeepSeek TUI (CodeWhale)** are worth watching—their daemon/scheduling capabilities are ahead of the market. However, both have reliability gaps (infinite loops, mode confusion).
- **For security-conscious users**: Implement **constitution-based policy files** (DeepSeek TUI pattern) or **project-level guardrails** (Claude Code pattern) as a mitigation against silent timeouts and over-autonomous agents.
- **For cost-constrained teams**: **OpenCode** offers the best multi-provider flexibility, but monitor the Go-tier rate-limit bugs closely. **Pi** is the strongest open-source option with an active extension ecosystem.

**Final assessment**: The AI CLI tool market is **pre-mature**—no single tool has solved security, reliability, cross-platform support, and extensibility simultaneously. The next 12 months will likely see consolidation as one or two tools break away. Currently, **OpenAI Codex** has the strongest engineering discipline, **Qwen Code** has the most ambitious vision, and **Claude Code** has the largest community. The winner will be the one that closes its reliability and security gaps first.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

## Claude Code Skills Community Highlights Report

### 1. Top Skills Ranking (Most-Discussed Pull Requests)

| # | PR | Skill / Fix | Status | Key Discussion Points |
|---|----|-------------|--------|----------------------|
| 1 | [#1298](https://github.com/anthropics/skills/pull/1298) | **Fix: run_eval.py reports 0% recall** – installs eval artifact as real skill, fixes Windows streams, trigger detection, and parallel workers | Open | Multiple independent reproductions of the recall bug (#556); description-optimisation loop has been optimizing against noise; heavy interest from skill-creator users |
| 2 | [#514](https://github.com/anthropics/skills/pull/514) | **Add document-typography skill** – corrects orphan words, widow paragraphs, numbering misalignment in AI-generated documents | Open | Addresses a universal pain point in Claude output; discussion about integrating with existing document skills; strong community support |
| 3 | [#538](https://github.com/anthropics/skills/pull/538) | **Fix: case-sensitive file references in PDF skill** – corrects 8 filename casing mismatches | Open | Simple but critical fix for case-sensitive filesystems (Linux/macOS); shows maintenance community |
| 4 | [#486](https://github.com/anthropics/skills/pull/486) | **Add ODT skill** – OpenDocument text creation, template filling, ODT-to-HTML conversion | Open | Fills gap for LibreOffice/open-source document workflows; discussion on LibreOffice integration vs native parsing |
| 5 | [#210](https://github.com/anthropics/skills/pull/210) | **Improve frontend-design skill** – clearer, more actionable instructions for Claude | Open | Aims to make the skill actually steer Claude within a single conversation; revision of existing skill |
| 6 | [#83](https://github.com/anthropics/skills/pull/83) | **Add skill-quality-analyzer & skill-security-analyzer** – meta-skills for evaluating other skills | Open | Community meta-tooling; good reception as it helps maintain quality across the ecosystem |
| 7 | [#541](https://github.com/anthropics/skills/pull/541) | **Fix: DOCX tracked change w:id collision with existing bookmarks** – prevents document corruption | Open | Root-caused OOXML shared ID space bug; hardcoded low IDs clashed with bookmarks |
| 8 | [#1367](https://github.com/anthropics/skills/pull/1367) | **Add self-audit skill v1.3.0** – mechanical file verification + four-dimension reasoning audit | Open | New universal audit skill; high potential for reducing hallucinations in code generation |

### 2. Community Demand Trends (from Issues)

The most active issue discussions reveal three concentrated demand areas:

- **Security & Trust** – Issue [#492](https://github.com/anthropics/skills/issues/492) (34 comments) warns that community skills distributed under `anthropic/` namespace impersonate official skills, enabling trust boundary abuse. Community is demanding namespace controls or verification badges.

- **Organizational Sharing** – Issue [#228](https://github.com/anthropics/skills/issues/228) (14 comments, 7 👍) requests org-wide skill sharing in Claude.ai, avoiding manual `.skill` file distribution via Slack.

- **Skill-Creator Tooling Reliability** – Multiple issues ([#556](https://github.com/anthropics/skills/issues/556), [#1169](https://github.com/anthropics/skills/issues/1169), [#1061](https://github.com/anthropics/skills/issues/1061), [#62](https://github.com/anthropics/skills/issues/62)) document hard crashes, 0% recall in evaluation, Windows incompatibilities, and skill disappearance. Community strongly wants a robust skill development pipeline.

Emerging new skill directions proposed in Issues include **agent-governance** ([#412](https://github.com/anthropics/skills/issues/412)), **compact-memory** (symbolic notation for agent state, [#1329](https://github.com/anthropics/skills/issues/1329)), and **MCP exposure of skills** ([#16](https://github.com/anthropics/skills/issues/16)).

### 3. High-Potential Pending Skills (Likely to Land Soon)

These PRs have active discussion, clear maintainability, and address well-validated pain points:

- **run_eval.py fixes** – Multiple PRs ([#1298](https://github.com/anthropics/skills/pull/1298), [#1099](https://github.com/anthropics/skills/pull/1099), [#1050](https://github.com/anthropics/skills/pull/1050), [#1323](https://github.com/anthropics/skills/pull/1323), [#362](https://github.com/anthropics/skills/pull/362), [#361](https://github.com/anthropics/skills/pull/361)) are converging. The core bug is confirmed and solutions exist — a merge is imminent.

- **Add CONTRIBUTING.md** ([#509](https://github.com/anthropics/skills/pull/509)) – addresses a community health gap (25% GitHub score); likely to merge as it requires no code change and closes a popular issue.

- **Add testing-patterns skill** ([#723](https://github.com/anthropics/skills/pull/723)) – covers testing philosophy, unit testing, React testing, E2E, mock strategy. Well-structured and generic.

- **Add sensory skill (macOS automation via AppleScript)** ([#806](https://github.com/anthropics/skills/pull/806)) – two-tier permission system; alternative to screenshot-based computer use; strong macOS user demand.

- **Add color-expert skill** ([#1302](https://github.com/anthropics/skills/pull/1302)) – comprehensive color knowledge (ISCC-NBS, Munsell, color spaces); self-contained and applicable to many domains.

### 4. Skills Ecosystem Insight

The community’s most concentrated demand is for **reliable tooling to author and evaluate skills** (fixes to `run_eval.py`, Windows compatibility, YAML validation), followed by **security and governance patterns** (trust boundary protection, agent-governance, organizational sharing), and **document formatting skills** (typography, ODT, DOCX, PDF fixes) to polish AI-generated output.

---

# Claude Code Community Digest — 2026-07-02

---

## Today's Highlights

Claude Code **v2.1.198** ships with the GA release of Claude in Chrome, background agent notifications, and a new `/dataviz` skill for chart/dashboard guidance. The community remains focused on two long-standing requests — **official Linux desktop support** (495 👍, still closed as resolved elsewhere) and a **critical security bug** where the cowork egress allowlist blocks custom domains (48 👍, open since March). A fresh **“EXTREME DANGER”** bug (#73125) warns that the tool can silently proceed without user confirmation after a 60-second timeout, drawing 8 comments in its first day.

---

## Releases

**v2.1.198** (latest)

- **Claude in Chrome** – Now generally available.
- **Background agent notifications** – `claude agents` sessions now fire OS `Notification` hooks (`agent_needs_input` / `agent_completed`).
- **`/dataviz` skill** – Provides chart and dashboard design guidance directly in chats.

[View release](https://github.com/anthropics/claude-code/releases/tag/v2.1.198)

---

## Hot Issues (10 noteworthy)

1. **[BUG] Cowork network egress allowlist not working** — Custom domains blocked with 403. 52 comments, 48 👍, open since March. Affects enterprise users relying on allowlist for secure egress.
   [#30112](https://github.com/anthropics/claude-code/issues/30112)

2. **[FEATURE] Official Claude Desktop build for Linux** — Closed after 50 comments and 495 👍. Community eager for native Linux desktop; closed status suggests it may be queued or addressed elsewhere.
   [#65697](https://github.com/anthropics/claude-code/issues/65697)

3. **[BUG] EXTREME DANGER: No response after 60s — continued without an answer** — The CLI’s `AskUserQuestion` silently accepts “no answer” and proceeds. Raised today with 8 comments and 25 👍; security-critical.
   [#73125](https://github.com/anthropics/claude-code/issues/73125)

4. **[BUG] Opus 4.8 missing from /model picker** — Model not listed in CLI or VS Code extension on Windows. 4 comments, regression reported.
   [#72918](https://github.com/anthropics/claude-code/issues/72918)

5. **[BUG] Autocompact thrashing** — Context fills within 2–4 turns after compact, reproducible with single word. Persists across sessions. 2 comments, core stability issue.
   [#72857](https://github.com/anthropics/claude-code/issues/72857)

6. **[BUG] Orphaned claude processes with --replay-user-messages** — Processes accumulate unbounded, consuming 2GB+ RAM. macOS, closed as duplicate but highlights resource leak.
   [#72109](https://github.com/anthropics/claude-code/issues/72109)

7. **[BUG] Cross-session message isolation failure** — An unrelated user message leaked into active conversation. Privacy/security concern, closed as duplicate.
   [#72051](https://github.com/anthropics/claude-code/issues/72051)

8. **[Bug] /remote-control command missing on Windows** — CLI command absent in PowerShell environment. 3 comments.
   [#72149](https://github.com/anthropics/claude-code/issues/72149)

9. **[BUG] Desktop app chat panel autoscrolls uncontrollably** — Screen reader users cannot read because view jumps; `autoScrollEnabled` setting ignored. Accessibility issue, closed as duplicate.
   [#72050](https://github.com/anthropics/claude-code/issues/72050)

10. **[FEATURE] Message queue management** — Request to rearrange/pause/cancel queued messages. 2 comments, reflects need for better async workflow control.
    [#72039](https://github.com/anthropics/claude-code/issues/72039)

---

## Key PR Progress

Only **2 pull requests** were updated in the last 24 hours:

1. **docs: fix Github → GitHub typo in README** – Small documentation fix, no code changes.
   [#72866](https://github.com/anthropics/claude-code/pull/72866)

2. **Create Cha** – Empty/incomplete PR without summary, likely a test or spam.
   [#72543](https://github.com/anthropics/claude-code/pull/72543)

> Note: PR activity is very low today. Most development work appears to be merged via internal pipelines, with community contributions being rare.

---

## Feature Request Trends

The most-requested feature directions distilled from recent issues:

- **Official Linux desktop build** – By far the highest-requested feature (495 👍). Despite closure, community still hopes for native Linux support.
- **Project-level guardrails** – A file automatically appended to every message for consistent policy enforcement. [#72043](https://github.com/anthropics/claude-code/issues/72043)
- **Cross-platform context sharing** – Seamless context continuity between Claude Desktop Projects and Claude Code. [#72068](https://github.com/anthropics/claude-code/issues/72068)
- **Message queue management** – Ability to rearrange, pause, or cancel pending messages. [#72039](https://github.com/anthropics/claude-code/issues/72039)
- **Resume from rewound state** – After accidental rewind, allow forward navigation without losing progress. [#72113](https://github.com/anthropics/claude-code/issues/72113)

---

## Developer Pain Points

Recurring frustrations reported by the community:

- **Safety filter false positives** – Multiple reports from user `sworrl` describing session-halting blocks on legitimate reverse-engineering and firmware analysis work, particularly for open-source tooling. Issues include AUP and “cyber” misclassifications.
- **Context compaction thrashing** – Frequent auto-compact cycles that degrade session quality, reported across macOS and unrelated projects.
- **Orphaned process leaks** – Background processes not reaped after session end, consuming significant RAM (2GB+).
- **Windows-specific gaps** – Missing `/remote-control` command, missing model in picker, and HTTP MCP server tool registration failures.
- **Silent user-input bypass** – The 60-second timeout behavior (#73125) poses a serious risk: critical decision points are skipped without user response.
- **Accessibility regressions** – Chat panel autoscrolls making screen reader use impossible; settings have no effect.

---

*Data sourced from [github.com/anthropics/claude-code](https://github.com/anthropics/claude-code). Digest prepared for 2026-07-02.*

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-07-02

## Today’s Highlights
A critical SSD‑endurance bug (Issue #28224) that could write ~640 TB/year has been fixed after three merged PRs, reducing logs by 85%. Meanwhile, a new pattern of GPT‑5.5 reasoning‑token clustering at 516/1034/1552 is under investigation (Issue #30364), and the team landed a series of security‑hardening PRs that block repository‑selected Git merge drivers and filters during patch operations. A prototype turning Codex apps into virtual HTTP MCP servers (PR #30000) also reached review.

## Releases
**rust‑v0.143.0‑alpha.33** — A minor alpha bump with no detailed changelog provided. No other versions were released in the last 24 hours.

## Hot Issues (10 noteworthy)
1. **[#28224 – Codex SQLite feedback logs can write ~640 TB/year](https://github.com/openai/codex/issues/28224)**  
   *121 comments, 415 👍* — Reported as a critical SSD‑endurance issue. Three PRs have been merged (released in v0.142.0), cutting log volume by ~85%. Now closed.  
   *Why it matters*: An extreme I/O problem resolved through community reporting and rapid fixes.

2. **[#3355 – Error after MacBook sleeps](https://github.com/openai/codex/issues/3355)**  
   *39 comments, 20 👍* — Requests to `chatgpt.com/backend-api/codex/responses` fail after lid close. Long‑standing (opened Sep 2025) with no resolution yet.  
   *Why it matters*: Disrupts mobile/portable workflows on macOS.

3. **[#29072 – Windows `apply_patch` fails due to sandbox setup EXE path](https://github.com/openai/codex/issues/29072)**  
   *35 comments, 23 👍* — Every `apply_patch` attempt on Windows fails because `codex-windows-sandbox-setup.exe` cannot be launched from the package path.  
   *Why it matters*: Blocks core functionality for Windows Desktop users.

4. **[#30364 – GPT‑5.5 reasoning‑token clustering at 516/1034/1552](https://github.com/openai/codex/issues/30364)**  
   *32 comments, 46 👍* — Aggregated metadata shows responses land disproportionately at exact token counts, coinciding with lower reasoning quality.  
   *Why it matters*: Suggests model‑side throttling or quantization that may degrade complex task performance.

5. **[#30440 – Codex uses bundled pnpm instead of host toolchain](https://github.com/openai/codex/issues/30440)**  
   *14 comments, 17 👍* — Sandbox ignores the system’s pnpm, causing build scripts to fail when they rely on local configuration.  
   *Why it matters*: Breaks reproducibility for Node.js projects.

6. **[#16099 – High GPU usage when app is on screen (Mac)](https://github.com/openai/codex/issues/16099)**  
   *10 comments, 13 👍* — 50–90% GPU utilisation even when idle, impacting battery life on MacBooks.  
   *Why it matters*: A performance/UX regression for portable users.

7. **[#29963 – Codex quota consumption has a serious bug](https://github.com/openai/codex/issues/29963)**  
   *10 comments, 9 👍* — Usage allowance is consumed far faster than expected; users report losing 20x Pro allocation prematurely.  
   *Why it matters*: Financial impact for paying subscribers.

8. **[#11912 – Request: custom compaction hook](https://github.com/openai/codex/issues/11912)**  
   *9 comments, 9 👍* — Developers want a hook to inject custom context‑compaction logic when token thresholds are reached.  
   *Why it matters*: High demand for fine‑grained control over model context.

9. **[#20538 – Desktop Preferences stuck on “unable to save”](https://github.com/openai/codex/issues/20538)**  
   *9 comments, 17 👍* — `configVersionConflict` errors prevent saving settings; persists across restarts.  
   *Why it matters*: Basic configuration is broken on Windows.

10. **[#30212 – 5‑hour allowance consumed in ~1 hour](https://github.com/openai/codex/issues/30212)**  
    *8 comments, 9 👍* — Pro 20x users see abnormal depletion, possibly related to the quota bug (#29963).  
    *Why it matters*: Another layer of the same quota issue, affecting trust in billing.

## Key PR Progress (10 important)
1. **[#30854 – Block selected merge drivers before three‑way patch application](https://github.com/openai/codex/pull/30854)**  
   Prevents `git apply --3way` from executing repository‑chosen low‑level merge drivers, closing a potential security vector.

2. **[#30848 – Block selected executable Git filters before patch application](https://github.com/openai/codex/pull/30848)**  
   Disallows clean/smudge filters from running during patch apply, revert, and preflight – another sandbox hardening step.

3. **[#30850 – Block selected Git filters before staging patch paths](https://github.com/openai/codex/pull/30850)**  
   Extends filter blocking to `git add` operations, preventing TOCTOU races where a file becomes a directory.

4. **[#30911 – Telemetry: narrow structured tool call timing](https://github.com/openai/codex/pull/30911)**  
   Restores existing `codex.tool_result` behaviour and emits `codex.tool_call` events only for direct tool calls, with safe snapshot markers.

5. **[#30770 – fix(websockets): ignore metadata for incremental requests](https://github.com/openai/codex/pull/30770)**  
   Prevents the Responses API from failing incremental requests when metadata is absent, improving success rate.

6. **[#30801 – sanitize exec config summary values](https://github.com/openai/codex/pull/30801)**  
   Normalises control characters in repository‑controlled values before displaying them in the exec config summary.

7. **[#30000 – Prototype Codex apps as virtual HTTP MCP servers](https://github.com/openai/codex/pull/30000)**  
   Architecture prototype exploring MCP as a standard interface for Codex apps. Reviews focus on ownership boundaries.

8. **[#28626 – Reuse directory entry metadata in skill scans](https://github.com/openai/codex/pull/28626)**  
   Avoids separate metadata requests per directory entry, significantly reducing overhead on remote exec servers.

9. **[#27190 – Add streaming file APIs](https://github.com/openai/codex/pull/27190)**  
   Introduces pull‑based `fs/readFile` and `fs/writeFile` for large files with bounded memory, backpressure, and cancellation.

10. **[#27091 – Eagerly compact Guardian threads between reviews](https://github.com/openai/codex/pull/27091)**  
    Runs compaction work immediately after a completed review, preventing expensive pre‑turn compaction delays on the next request.

## Feature Request Trends
- **Custom compaction hooks** (#11912) – Developers want to intercept token‑limit thresholds to inject their own context‑shrinking logic.
- **Inline LaTeX math rendering** (#14985) – Block equations work, but inline `$...$` is not rendered in the Codex App.
- **Inject command output into active session** (#22003) – Background completions should be able to pipe results into the current Codex thread.
- **Usage dashboard tier clarity** (#30783) – Users on Pro 5x/20x cannot tell which tier’s remaining percentage is shown.
- **User‑configurable feature gates** (implied by #28481) – MCP tools like `mcp__node_repl__js` are gated client‑side without any toggle for advanced users.

## Developer Pain Points
- **Windows‑platform friction** – Repeated issues with patch application (#29072, #16229), UTF‑8 BOM (#15967), menu bars (#14450), sandbox setup (#29072), and config‑save errors (#20538) make Windows a second‑class experience.
- **Quota / rate‑limit bugs** – Multiple reports (#29963, #30212, #30783) of usage being consumed far faster than expected, with unclear tier identification.
- **MacBook sleep disconnect** (#3355) – After lid close, all API requests fail until the app is manually restarted.
- **Bundled toolchain conflicts** (#30440) – The sandbox ignores the host environment’s Node/Bun/pnpm, breaking builds that rely on local configuration.
- **GPU overhead on macOS** (#16099) – Idle GPU usage drains battery and heats up devices.
- **Subagent wake‑up failures** (#15723) – Background subprocesses do not notify the calling agent when they finish, stalling multi‑step workflows.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-07-02

## Today’s Highlights

A **critical security patch** lands in today’s nightly release, fixing a symbolic-link directory escape in the memory import processor. Meanwhile, the community is buzzing about a long-standing issue where subagents falsely report “GOAL” success after hitting `MAX_TURNS`, and a PR to natively support `AGENTS.md` out of the box has been opened. The pull request queue also shows strong activity around fixing thought leakage, encoding handling, and CI supply-chain vulnerabilities.

## Releases

**v0.51.0-nightly.20260702.gff00dacd9**  
_One change_  
- **Fix**: Resolve symbolic link directory escape in memory import processor ([#28233](https://github.com/google-gemini/gemini-cli/pull/28233)) – prevents an attacker from crafting a malicious repository that could read outside the workspace via symlinks.

[Full changelog](https://github.com/google-gemini/gemini-cli/compare/v0.51.0-nightly.20260701.g7f00c5fe5...v0.51.0-nightl)

## Hot Issues (10 noteworthy)

1. **[#22323 – Subagent recovery after MAX_TURNS is reported as GOAL success](https://github.com/google-gemini/gemini-cli/issues/22323)**  
   A subagent (`codebase_investigator`) claims `status: "success"` even though its own output says it hit the maximum turn limit without doing any analysis. This masks real failures and misleads users. 9 comments, 2 👍.

2. **[#21409 – Generalist agent hangs](https://github.com/google-gemini/gemini-cli/issues/21409)**  
   Since v0.33.0, many users report the CLI hangs indefinitely when delegating to the generalist agent. Workaround: tell the model not to use subagents. 7 comments, 8 👍 – high community impact.

3. **[#21968 – Gemini does not use skills and sub-agents enough](https://github.com/google-gemini/gemini-cli/issues/21968)**  
   Custom skills and sub-agents are almost never invoked unless explicitly requested, even for closely related tasks. 6 comments.

4. **[#25166 – Shell command execution gets stuck with "Waiting input"](https://github.com/google-gemini/gemini-cli/issues/25166)**  
   After a simple CLI command finishes, Gemini sometimes hangs while the shell is still shown as “awaiting user input.” 4 comments, 3 👍.

5. **[#26525 – Add deterministic redaction and reduce Auto Memory logging](https://github.com/google-gemini/gemini-cli/issues/26525)**  
   Auto Memory reads transcripts and sends content to an extraction model *before* redaction – a privacy concern. Also logs may contain secrets. 5 comments.

6. **[#26522 – Stop Auto Memory from retrying low-signal sessions indefinitely](https://github.com/google-gemini/gemini-cli/issues/26522)**  
   Sessions judged low-signal are never marked as processed, so they keep reappearing. 5 comments.

7. **[#21983 – Browser subagent fails in Wayland](https://github.com/google-gemini/gemini-cli/issues/21983)**  
   The browser agent terminates with “GOAL” immediately on Wayland desktops. 4 comments, 1 👍.

8. **[#24246 – Gemini CLI encounters 400 error with > 128 tools](https://github.com/google-gemini/gemini-cli/issues/24246)**  
   More than 128 tools (custom skills + built-ins) causes a 400 error – the agent should scope tools better. 3 comments.

9. **[#22672 – Agent should stop/discourage destructive behavior](https://github.com/google-gemini/gemini-cli/issues/22672)**  
   The model occasionally uses `git reset`, `--force`, or dangerous DB commands when safer alternatives exist. 3 comments, 1 👍.

10. **[#22598 – Subagent trajectory should be visible via `/chat share`](https://github.com/google-gemini/gemini-cli/issues/22598)**  
   Subagent logs are saved but not accessible in shared chats, making debugging difficult. 2 comments, 1 👍.

## Key PR Progress (10 important)

1. **[#28240 – Fix #28227: add support for AGENTS.md out of the box](https://github.com/google-gemini/gemini-cli/pull/28240)**  
   Makes `AGENTS.md` a default context file alongside `GEMINI.md`, so users don’t need to manually configure it. (Open)

2. **[#27747 – fix(cli): prevent infinite loop in ghost text wrapping](https://github.com/google-gemini/gemini-cli/pull/27747)**  
   Resolves a freeze when an emoji or wide character is displayed in a very narrow terminal – fixes #19985. (Closed)

3. **[#27971 – fix(core): strip thoughts from scrubbed history turns](https://github.com/google-gemini/gemini-cli/pull/27971)**  
   Stops internal reasoning (“thoughts”) from leaking into history, which caused the model to imitate scratchpad monologues in subsequent turns. (Closed)

4. **[#27979 – Wrap read_mcp_resource output with wrapUntrusted()](https://github.com/google-gemini/gemini-cli/pull/27979)**  
   Security alignment – MCP resource text is now treated as untrusted, consistent with the `mcp_tool` implementation. (Closed)

5. **[#27996 – fix(core): decode response body using charset from Content-Type header](https://github.com/google-gemini/gemini-cli/pull/27996)**  
   `web-fetch` now respects `charset` in HTTP headers, fixing garbled text on non-UTF-8 sites (e.g., Chinese/Japanese). (Closed)

6. **[#27994 – fix(core): insert skill/agent content literally in system prompt substitutions](https://github.com/google-gemini/gemini-cli/pull/27994)**  
   Prevents `String.replace` from interpreting special characters (`$`, `&`, etc.) in skill/sub-agent definitions, which could corrupt prompts. (Closed)

7. **[#27986 – feat(acp): report cached and thought tokens in PromptResponse.usage](https://github.com/google-gemini/gemini-cli/pull/27986)**  
   ACP mode now exposes cached input tokens and reasoning tokens – important for accurate cost estimation. (Closed)

8. **[#28126 – fix(core-tools): show ellipsis on multi-line edit snippets](https://github.com/google-gemini/gemini-cli/pull/28126)**  
   Prevents multi-line edits with a short first line from appearing as a single-line change in the UI. (Open)

9. **[#28233 – fix(core): resolve symbolic link directory escape in memory import processor](https://github.com/google-gemini/gemini-cli/pull/28233)**  
   High-severity security fix – incorporated in today’s nightly release. (Closed)

10. **[#28223 – fix(core-tools): bypass LLM correction for JSON and IPYNB files](https://github.com/google-gemini/gemini-cli/pull/28223)**  
    The `write_file` and `replace` tools now skip LLM-based “correction” for `.json` and `.ipynb` files, preventing corruption. (Open)

## Feature Request Trends

- **AST-aware file operations** – Several EPICs (e.g., #22745, #22746) propose using AST-aware tools for file reads, search, and codebase mapping to reduce tokens and improve precision.
- **Component-level evaluations** – Issue #24353 tracks the expansion of behavioral eval tests beyond simple prompts to cover subagents, tools, and memory.
- **Better subagent utilisation** – Users want the generalist agent to more aggressively delegate to custom skills/sub-agents (#21968) and to auto-discover `AGENTS.md` (#28240).
- **Agent self-awareness** – Requests for the CLI to accurately describe its own flags, hotkeys, and subagent capabilities (#21432).
- **Subagent trajectory sharing** – Making subagent logs visible in `/chat share` (#22598) for debugging and evaluation.
- **Destructive operation safeguards** – The community wants the agent to avoid dangerous commands (e.g., `git reset --hard`) without confirmation (#22672).

## Developer Pain Points

- **Subagent false successes** – Subagents frequently report “GOAL”/“success” even when they hit a turn limit or fail (#22323, #21983).
- **Stuck execution** – Shell commands and subagents hang or require manual cancellation (#21409, #25166, #22465).
- **Memory system bugs** – Auto Memory retries low-signal sessions, logs secrets before redaction, and silently skips invalid patches (#26522, #26525, #26523).
- **Terminal/UI issues** – Ghost text wrapping freezes on narrow terminals (#27747), corruption after exiting external editors (#24935), and flicker on resize (#21924).
- **Tool overload** – Having >128 tools causes a 400 error; the agent does not limit its tool selection (#24246).
- **Symlink and encoding quirks** – Symlinked agent files aren’t recognised (#20079), and web fetching fails on non-UTF-8 pages (#27996).
- **JSON/IPYNB file corruption** – LLM correction logic breaks structured file writes (#28223).
- **Thought leakage** – Internal reasoning from previous turns contaminates future history (#27971).

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-07-02

## Today’s Highlights
Two releases landed in the last 24 hours: v1.0.68 introduced support for the **kimi-k2.7-codex** model and smarter IDE disconnect handling, while v1.0.69-0 improved the `/sandbox` experience with file/folder completion. On the issue tracker, authentication bugs resurface (Issue #3596) and several new triage items highlight pain points in plugin registration, Windows hooks, and model billing accuracy.

## Releases
Two new versions published:

- **[v1.0.69-0](https://github.com/github/copilot-cli/releases/tag/v1.0.69-0)** *(2026-07-02)*  
  **Added** – File and folder completion for `/sandbox` path entries.  
  **Fixed** – Session branch label now updates when working directory changes in the Sessions split view. Unnecessary MCP reloads avoided when returning to an already-loaded session. Prevented `tgrep` indexer from running unnecessarily.

- **[v1.0.68](https://github.com/github/copilot-cli/releases/tag/v1.0.68)** *(2026-07-01)*  
  **Added** – Support for the kimi-k2.7-codex model. Focused field in `/mcp` config form now marked with `❯` chevron (not color alone).  
  **Fixed** – IDE tools remain available during transient disconnects, returning a clear error and auto-recovering.

## Hot Issues
*10 noteworthy issues from the last 24 hours – covering authentication, model availability, sandboxing, Windows‑specific bugs, and feature gaps.*

1. **[#3596 – Error loading model list: Not authenticated](https://github.com/github/copilot-cli/issues/3596)**  
   *area: authentication / sessions / models*  
   **Why it matters:** Resuming a session breaks model listing (`/model`), forcing users to start a new session. Received 11 👍 and 8 comments. Not yet fixed in the latest releases.

2. **[#1504 – Add custom theme support](https://github.com/github/copilot-cli/issues/1504)**  
   *area: theming/accessibility*  
   **Why it matters:** Community’s most-upvoted open feature (20 👍). Users want JSON‑based custom themes and a `/theme` command to load/share them. 6 comments.

3. **[#3997 – Model “gpt-5.3-codex” is not available](https://github.com/github/copilot-cli/issues/3997)**  
   *triage*  
   **Why it matters:** Session creation fails entirely with this error. Affects agent mode; opened yesterday, already has 5 comments.

4. **[#3948 – Any web_fetch: TypeError: fetch failed](https://github.com/github/copilot-cli/issues/3948)**  
   *area: networking / tools*  
   **Why it matters:** Every web_fetch tool call fails even though authentication works. No proxy or environment issue identified. 4 comments.

5. **[#3158 – Plan→Compact→Re-Plan infinite loop (217 cycles, zero execution)](https://github.com/github/copilot-cli/issues/3158)**  
   *area: agents / context-memory*  
   **Why it matters:** High‑severity bug where the agent gets stuck in a compaction loop and never executes. Reported from internal agency tool; 3 comments.

6. **[#3331 – Auto-update plugins on CLI startup via marketplace flag](https://github.com/github/copilot-cli/issues/3331)**  
   *area: plugins*  
   **Why it matters:** Teams deploying plugin updates have no guarantee users are on the latest version. 4 👍, 3 comments.

7. **[#4004 – plugin install does not register plugin MCP servers into mcp-config.json](https://github.com/github/copilot-cli/issues/4004)**  
   *triage*  
   **Why it matters:** Newly opened (2026-07-02). Plugin MCP `.mcp.json` is copied but never registered, breaking MCP‑based plugin functionality.

8. **[#4003 – Support custom model endpoint in Copilot CLI](https://github.com/github/copilot-cli/issues/4003)**  
   *triage*  
   **Why it matters:** Users want the same custom‑endpoint capability VS Code offers, enabling local/private model development. Zero comments yet but highly demanded.

9. **[#3978 – Copilot CLI incorrectly switches back to previous model after switching to BYOK](https://github.com/github/copilot-cli/issues/3978)**  
   *area: sessions / models*  
   **Why it matters:** After using BYOK (bring‑your‑own‑key), the CLI reverts to the previous model, causing billing surprises. No comments yet.

10. **[#3995 – Support persistent command deny-rules in permissions-config.json](https://github.com/github/copilot-cli/issues/3995)**  
    *area: permissions / tools*  
    **Why it matters:** Only allow‑rules exist today. Users need persistent deny‑rules for command families (e.g., `rm -rf`). 1 👍, opened 2026-07-01.

## Key PR Progress
No pull requests were updated or created in the last 24 hours.

## Feature Request Trends
*The most‑requested feature directions emerging from recent issues:*

- **Custom model endpoints** – Mirroring VS Code’s ability to use local/private models (Issue #4003).
- **Custom theme / accessibility** – JSON‑based theme system (#1504) and screen‑reader echo of typed characters (#3993) indicate growing demand for non‑visual customization.
- **Plugin lifecycle automation** – Auto‑update on startup (#3331) and live terminal panels for plugin UIs (#3979).
- **Extended permissions model** – Persistent deny‑rules (#3995) to complement the current allow‑only system.
- **Platform parity** – Windows hooks (`$CLAUDE_PROJECT_DIR` not set, #4001) and sandbox support on Linux (#3653) show that cross‑platform consistency remains a high priority.

## Developer Pain Points
*Recurring frustrations reported this week:*

- **Authentication session leaks** – Resuming a session loses authentication status (#3596) and model state (#3978).
- **Model availability errors** – Frequent “model not available” (#3997) and billing mismatches (#3998) degrade trust.
- **Plugin MCP registration gaps** – Plugin MCP servers are copied but never wired into the CLI configuration (#4004).
- **Windows‑specific failures** – Hooks run under PowerShell instead of bash (#4001), plugin update uses cache instead of fresh download (#3627), and `/ide` fails with IntelliJ (#2891).
- **Data loss on session switch** – `/new` discards in‑memory usage stats without writing `session.shutdown` to `events.jsonl` (#3994).
- **Infinite agent loops** – The plan‑compact‑re‑plan cycle (#3158) stalls development completely.
- **Terminal rendering regressions** – Flicker on Windows (#3984) and inability to copy output in browser‑based VS Code Server (#3996) hamper daily use.

---

*Data as of 2026-07-02 23:59 UTC. For full details visit [github.com/github/copilot-cli](https://github.com/github/copilot-cli).*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

Here is the **Kimi Code CLI Community Digest** for **2026-07-02**, based on the provided GitHub data.

---

## Kimi Code CLI Community Digest — 2026-07-02

### 1. Today's Highlights
The community remains quiet on the release front (no new versions), but two strategic discussions dominate the agenda: a critical **branding inconsistency** (#2483) that fragments the ecosystem’s naming conventions, and a **highly-requested feature** (#2482) to automatically persist long `/goal` prompts to a file. Meanwhile, the long-standing **infinite read-loop bug** (#640) resurfaces with a recent comment indicating the issue persists on version `0.76` with custom Anthropic endpoints. The "Kimi CLI" → "Kimi Code" migration is effectively half-done.

### 2. Releases
**No new releases in the last 24 hours.** The latest published version remains **v0.76**. This signals a potential stabilization period or a lull in the CI pipeline. Developers relying on `mimo-v2-flash` via custom endpoints should note the unresolved loop bug below.

### 3. Hot Issues (Noteworthy Picks)

**#2483 — [branding] "Kimi CLI" → "Kimi Code" migration is half-done**
- *Why it matters:* This issue exposes a systemic naming fracture. The repo description, README, Zed/VS Code extensions, SDK, binary path, and PyPI package name are all using **at least four different names**. This causes confusion for new adopters writing tutorials or CI scripts.
- *Community reaction:* Zero comments so far, but high strategic importance. The issue author has methodically cataloged all inconsistencies, making it actionable for maintainers.
- *Link:* [Issue #2483](https://github.com/MoonshotAI/kimi-cli/issues/2483)

**#2482 — Feature Request: Super long goals auto-save to goal.md with CLI edit/pause**
- *Why it matters:* The current 4,000-byte limit for `/goal` is a bottleneck for complex, long-running tasks. Users want a workflow similar to Codex/Claude Code where long instructions are persisted to disk and re-read on resume.
- *Community reaction:* Positive. No explicit +1s yet, but the request is clearly scoped and references existing industry patterns.
- *Link:* [Issue #2482](https://github.com/MoonshotAI/kimi-cli/issues/2482)

**#640 — [bug] Kimi CLI stuck in reading one file again and again (infinite loop)**
- *Why it matters:* This is a **blocking reliability bug**. The user reports the CLI enters an infinite read-loop on `Linux 6.18.3-arch1` using a custom Anthropic endpoint (`mimo-v2-flash`). 16 comments indicate significant community troubleshooting.
- *Community reaction:* High frustration. The issue was opened in January 2026 but updated today—meaning it is **not yet resolved**. One 👍 suggests moderate, not viral, impact.
- *Link:* [Issue #640](https://github.com/MoonshotAI/kimi-cli/issues/640)

### 4. Key PR Progress

**#2369 — [CLOSED] feat(subagent): add API key pool for parallel subagent execution**
- *Description:* Introduces a **round-robin API key allocator** (`llm_key_pool.py`) to support parallel subagent execution. Also includes fixes for transaction loop issues and subagent error handling.
- *Why it matters:* This is a major infrastructure improvement for users running concurrent subagents. It closes a long-standing performance bottleneck.
- *Status:* Merged/Closed. This is a significant win for the project’s concurrency model.
- *Link:* [PR #2369](https://github.com/MoonshotAI/kimi-cli/pull/2369)

### 5. Feature Request Trends
Based on recent issues, the community is converging on three major directions:

- **Branding & Ecosystem Consistency**: Requests for a unified "Kimi Code" naming across all downstream projects (extension, SDK, docs, binary name). This is not a functional feature but a **developer experience hygiene** request.
- **Long Goal Persistence**: The desire for automatic file-based goal management (`.goal.md`) with inline CLI editing and pause/resume support. Users want to move beyond the 4KB slash command limit.
- **Parallelism & API Key Management**: High enthusiasm for the recently merged API key pool feature suggests the community wants more robust support for **concurrent subagents and multi-key configurations**.

### 6. Developer Pain Points
The most frequently recurring frustrations include:

- **Reliability Bugs (Looping)**: Issue #640 highlights a critical reliability gap where the CLI enters an infinite read-loop on custom endpoints. Developers using non-standard models face the highest risk.
- **Configuration Friction**: Users running custom Anthropic endpoints via `config.toml` are more likely to hit edge cases. The absence of validation or fallback logic for such setups is a persistent pain point.
- **Naming Confusion**: The "Kimi CLI" vs "Kimi Code" branding split causes confusion in documentation, error messages, and package managers. Developers integrating the tool into CI/CD or editor plugins report inconsistent paths and names.
- **Goal Length Limitations**: Power users working on complex, multi-step goals find the 4KB slash command limit too restrictive, forcing them to manually split prompts or use external files.

---

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest – 2026-07-02

## Today's Highlights
A major service disruption hit OpenCode Go users today, with widespread **provider rate limit errors** and **inference unavailability** on DeepSeek V4 Flash and Mimo models, sparking 28+ comments on a single issue. The team shipped **v1.17.13** to fix reasoning model compat and stale Copilot response IDs, while pull requests addressing tool serialization bugs and model output limits are now under review.

## Releases
**v1.17.13** – [Release](https://github.com/anomalyco/opencode/releases/tag/v1.17.13)
- **Core**: Force reasoning mode for OpenAI-compatible reasoning models; stop replaying stale GitHub Copilot response item IDs (fixes follow-up request failures).
- **Desktop**: Allow question prompts to be minimized.

## Hot Issues (10 selected)

1. **Inference temporarily unavailable** [#34893](https://github.com/anomalyco/opencode/issues/34893)  
   *28 comments, 22 👍* – DeepSeek V4 Flash (OpenCode Go) went down for ~5 minutes on Ubuntu; users report a 503 error. Highly active, likely backend incident.

2. **Go returns "Provider rate limit exceeded" despite 0% rolling usage** [#34884](https://github.com/anomalyco/opencode/issues/34884)  
   *13 comments, 6 👍* – Free Zen models work, but Go tier returns rate-limit errors even with no usage. Impacts paying users; retry loop never recovers.

3. **"Eror from provider (Console Go): Provider rate limit exceeded"** [#34886](https://github.com/anomalyco/opencode/issues/34886)  
   *8 comments, 2 👍* – Similar pattern; user notes monthly usage increased without Code plan activating. Suggests a billing/rate-limit sync bug.

4. **Keep getting rate limited on Zen** [#13318](https://github.com/anomalyco/opencode/issues/13318)  
   *10 comments* – Long-standing issue (Feb 2026) with Kimi-K2.5 on Zen; paying users still hit limits. Resurfaced today with new comments.

5. **Qwen 3.7 Plus/Max via OpenRouter – invalid tool calls** [#33618](https://github.com/anomalyco/opencode/issues/33618)  
   *8 comments* – Tool calls fail with empty names, causing retries and aborted sessions. Affects users of newer Qwen models.

6. **opencode go无法使用 (Go unavailable)** [#34898](https://github.com/anomalyco/opencode/issues/34898)  
   *6 comments* – Chinese user reports HTTP 429 for DeepSeek v4-flash and Mimo v2.5. Regional unconfirmed.

7. **New Layout and Designs – cannot switch Plan/Build** [#31972](https://github.com/anomalyco/opencode/issues/31972)  
   *5 comments, 7 👍* – Enabling the new layout breaks the Plan/Build mode toggle (UI and shortcut). Desktop regression.

8. **TUI corrupted by auto-fetch of models.dev error log leaks** [#34730](https://github.com/anomalyco/opencode/issues/34730)  
   *5 comments, 4 👍* – Network-restricted environments get periodic ERROR logs from `models.dev/api.json` fetch failures, corrupting TUI rendering.

9. **DeepSeek V4 Flash – "Provider rate limit exceeded" with Go** [#34885](https://github.com/anomalyco/opencode/issues/34885)  
   *4 comments, 3 👍* – Retry loop never succeeds; model became unavailable suddenly. Part of today’s Go outage wave.

10. **Service is too busy – 10 min retries** [#34899](https://github.com/anomalyco/opencode/issues/34899)  
    *3 comments* – “Service is too busy” with 10-minute retry cooldown; user cannot switch to alternative LLM. Backend capacity issue.

## Key PR Progress (10 selected)

1. **feat(uk): comprehensive translation improvements** [#32566](https://github.com/anomalyco/opencode/pull/32566)  
   Improves Ukrainian, Polish, Russian localization across core, web, and docs. Long-running, resolves type-check blocker.

2. **fix(provider): respect model limit.output instead of capping at 32k** [#34901](https://github.com/anomalyco/opencode/pull/34901)  
   Closes several long-standing issues about output token limits being overridden. Important fix for large context models.

3. **fix(app): show readable plugin labels** [#34897](https://github.com/anomalyco/opencode/pull/34897)  
   Displays human-readable names for local plugins in status popover; preserves registry semantics.

4. **fix(shell): add PowerShell UTF-8 command wrapper on Windows** [#31985](https://github.com/anomalyco/opencode/pull/31985)  
   Closes 5 issues by ensuring PowerShell commands handle UTF-8 correctly. Critical for Windows adoption.

5. **fix(tui): scope non-git sessions by directory, not hierarchical path** [#31210](https://github.com/anomalyco/opencode/pull/31210)  
   Fixes 6 issues where non-git projects shared session data; sessions now properly isolated by directory.

6. **fix: zen toOaCompatibleRequest reading tool.function.name** [#34887](https://github.com/anomalyco/opencode/pull/34887) (closes [#34892](https://github.com/anomalyco/opencode/issues/34892))  
   Tool serialization bug: `toOaCompatibleRequest()` read `tool.name` instead of `tool.function.name`. Quick fix merged in progress.

7. **fix(app): reconcile session_status in bootstrap so stale busy clears** [#32128](https://github.com/anomalyco/opencode/pull/32128)  
   Fixes stale "busy" states after restarts. Important for Desktop reliability.

8. **feat(tui): add global session picker toggle** [#33450](https://github.com/anomalyco/opencode/pull/33450)  
   Adds ability to discover/resume sessions from other projects, addressing a long-requested workflow gap.

9. **feat(app): terminal improvements** [#34747](https://github.com/anomalyco/opencode/pull/34747)  
   Introduces `dnd-kit` tabs for terminal panel; fixes terminal layout. Improves multi-terminal UX.

10. **fix(tui): collapse fragmented reasoning parts and strip thinking echo** [#32152](https://github.com/anomalyco/opencode/pull/32152)  
    Deduplicates reasoning blocks and removes echoed thinking from chat output. Addresses UI clutter.

## Feature Request Trends
- **Context usage display** ([#34900](https://github.com/anomalyco/opencode/issues/34900)) – Users want to see token/context usage in Desktop, especially for local OpenAI-compatible models.
- **New model additions** ([#34889](https://github.com/anomalyco/opencode/issues/34889), [#34883](https://github.com/anomalyco/opencode/issues/34883)) – Requests for GLM-5.2 and other models in the Alibaba provider list.
- **Inline artifact rendering** ([#25076](https://github.com/anomalyco/opencode/issues/25076)) – Preview SVG/HTML generated by models directly in chat, not as raw code.
- **Custom tool/hook capabilities** ([#34890](https://github.com/anomalyco/opencode/issues/34890)) – Hooks similar to Claude CLI for creating custom commands without LLM.
- **Structured multi-choice prompt primitive** ([#34871](https://github.com/anomalyco/opencode/issues/34871)) – Mid-turn N-option decisions for agents/plugins (AskUserQuestion equivalent).

## Developer Pain Points
- **Go/Zen rate limit errors** dominate today: many users report "Provider rate limit exceeded" despite low or zero usage. The retry mechanism (15s–10min) is ineffective and aggravates.
- **Service unavailability of Go models** (DeepSeek V4 Flash, Mimo) – sudden 503/429 errors, impacting paying subscribers.
- **Desktop UI regressions** – New layout breaks Plan/Build switching; prompt minimization works but other toggles fail.
- **Terminal/input issues** – Backspace key unresponsive in some terminal multiplexers ([#34878](https://github.com/anomalyco/opencode/issues/34878)); ConPTY compatibility needs attention.
- **Session isolation bugs** – Duplicate git directories merge sessions on Desktop ([#31632](https://github.com/anomalyco/opencode/issues/31632)); fixed in PR but not yet released.
- **Network-dependent fetches** – `models.dev` HTTP errors corrupt TUI; users with restricted networks hit persistent error logs.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest – 2026-07-02

## Today's Highlights
The community pushed hard on Linux compliance and WSL stability, with the long‑running XDG Base Directory request (#2870) finally closed after 17 comments and 34 👍. Several critical bugs were fixed around session storage integrity and TUI copy‑paste, while model updates for GitHub Copilot (Sonnet 5) and new provider support for Amazon Bedrock Mantle and SQLite session storage moved forward. The pace of PRs remains high, with 19 merged or open in the last 24 hours.

## Releases
No new releases in the last 24 hours. The latest available version remains 0.80.3, though an upgrade regression (#6215) was reported and closed.

## Hot Issues

1. **#2870 – Follow XDG Base Directory**  
   *Closed, 17 comments, 34 👍*  
   The application now respects `$XDG_CONFIG_HOME` for Linux configuration, ending years of home‑directory clutter. Strong community consensus drove this change.  
   [GitHub](https://github.com/earendil-works/pi/issues/2870)

2. **#6187 – Pi login hangs in WSL after GitHub Copilot device auth**  
   *Closed, 9 comments*  
   On WSL, browser‑based Copilot auth completes but the terminal never detects it, leaving the client stuck. The fix ensures WSL listeners pick up the OAuth callback.  
   [GitHub](https://github.com/earendil-works/pi/issues/6187)

3. **#6234 – Escape leaves Pi stuck in “Working” when extension hook never settles**  
   *Closed, 6 comments*  
   Pressing Escape during an active run could leave the TUI frozen if an extension context hook never completes. The double‑Escape abort path is now more resilient.  
   [GitHub](https://github.com/earendil-works/pi/issues/6234)

4. **#6231 – Auth blocking local models**  
   *Closed, 4 comments*  
   Local models (e.g., DeepSeek via Dwarf Star Four) incorrectly required OAuth/login. The provider check was fixed to skip authentication for local endpoints.  
   [GitHub](https://github.com/earendil-works/pi/issues/6231)

5. **#6189 – `question` example hangs when multiple calls are batched**  
   *Closed, 4 comments*  
   The example UI struggled when the model emitted several `question` tool calls simultaneously. Adding `executionMode: sequential` resolved the batch handling.  
   [GitHub](https://github.com/earendil-works/pi/issues/6189)

6. **#6191 – Add `PI_SKILL_PATH` environment variable**  
   *Closed, 4 comments*  
   Users with skills in multiple repos can now set a per‑project env var instead of remembering CLI flags. Environment‑file integration was highly requested.  
   [GitHub](https://github.com/earendil-works/pi/issues/6191)

7. **#5570 – Support `--no-skills` / `--skill` behavior in project settings**  
   *Open, 4 comments*  
   The ability to configure skill inclusion/exclusion in `.pi/settings.json` is still open. A partial PR (#6236) landed today, but full settings support remains.  
   [GitHub](https://github.com/earendil-works/pi/issues/5570)

8. **#6208 – Add Claude Sonnet 5 to the GitHub Copilot provider**  
   *Closed, 4 comments*  
   After Sonnet 5 went GA on Copilot, the community quickly requested support. The model catalog was updated to include it under the GitHub Copilot provider.  
   [GitHub](https://github.com/earendil-works/pi/issues/6208)

9. **#6215 – `pi update` fails on 0.80.3 due to missing `@smithy/node-http-handler@^4.9.1`**  
   *Closed, 4 comments*  
   An unresolved dependency blocked upgrades. The maintainers patched the version range to resolve the PNPM error.  
   [GitHub](https://github.com/earendil-works/pi/issues/6215)

10. **#6206 – Clamping to context window prevents artificial context limits**  
    *Open, 3 comments*  
    A recent fix clamps `max_tokens` to the model’s context window, but this breaks users who intentionally set smaller limits (e.g., for cost or latency). Discussion continues.  
    [GitHub](https://github.com/earendil-works/pi/issues/6206)

## Key PR Progress

1. **#6252 – Fix find paths from Windows drive root**  
   *Merged*  
   Resolves path‑relative formatting for Windows drive roots (e.g., `C:\`), preventing doubled slashes in find results. Windows regression tests added.  
   [GitHub](https://github.com/earendil-works/pi/pull/6252)

2. **#6248 – Stop padding TUI lines with trailing spaces**  
   *Merged*  
   Eliminates trailing space characters that break copy/paste in xterm.js terminals (VS Code integrated terminal). Each line is now written without padding.  
   [GitHub](https://github.com/earendil-works/pi/pull/6248)

3. **#6244 – Keep interactive input and footer sticky**  
   *Merged*  
   Introduces a sticky‑bottom boundary API so input, widgets, and footer remain visible during long scrollback. Includes tests for the new “sticky” rendering.  
   [GitHub](https://github.com/earendil-works/pi/pull/6244)

4. **#6243 – Fix UUID collisions and race conditions in session storage**  
   *Merged*  
   Three critical bugs in `JsonlSessionStorage` and `InMemorySessionStorage`: UUID truncation, race in `setLeafId`, and a missing rollback on write failure. Session corruption is now prevented.  
   [GitHub](https://github.com/earendil-works/pi/pull/6243)

5. **#6241 – Fix TUI offscreen redraws for stable‑height updates**  
   *Open*  
   When the rendered line count doesn’t change, the renderer now clamps repainting to visible rows instead of replaying the entire scrollback. Performance improvement.  
   [GitHub](https://github.com/earendil-works/pi/pull/6241)

6. **#6236 – Allow project‑level skills setting**  
   *Merged*  
   Partially addresses #5570 by introducing a `skill` key in `.pi/settings.json`. Users can now list skill paths per project without CLI flags.  
   [GitHub](https://github.com/earendil-works/pi/pull/6236)

7. **#5678 – Add `excludeFromContext` for custom messages**  
   *Merged*  
   Allows extensions to mark messages that should be persisted and rendered but excluded from the LLM context. Compaction, summarization, and branching now respect this flag.  
   [GitHub](https://github.com/earendil-works/pi/pull/5678)

8. **#6227 – Add SQLite session storage**  
   *Open*  
   Under the `PI_SQLITE_SESSION_STORAGE` flag, session transcripts are written to both JSONL (default) and SQLite in parallel. New API for querying sessions via SQL.  
   [GitHub](https://github.com/earendil-works/pi/pull/6227)

9. **#6225 – Infer toolUse when provider omits finish_reason**  
   *Merged*  
   Some OpenAI‑compatible providers (e.g., NVIDIA NIM) send tool‑call chunks without a `finish_reason` marker. The stream parser now infers tool calls from the presence of tool‑call delta fields.  
   [GitHub](https://github.com/earendil-works/pi/pull/6225)

10. **#6213 – Implement AOT compilation for TypeScript extensions**  
    *Merged*  
    Uses esbuild to pre‑compile extensions during `pi install`/`pi update`, cutting startup time from seconds to milliseconds. Reduces dependency on JIT compilation.  
    [GitHub](https://github.com/earendil-works/pi/pull/6213)

## Feature Request Trends

- **Configuration flexibility**: Multiple requests for environment‑variable and project‑level settings (skill paths, skills on/off, theme overrides). The community wants declarative, per‑repo configuration rather than repeated CLI flags.
- **Model/provider updates**: Rapid response to new models (Sonnet 5, GLM‑5.2, MiniMax‑M3) and new providers (Amazon Bedrock Mantle, Anthropic Vertex). Users expect the catalog to stay current with cloud releases.
- **Context and diagnostics**: Calls for `/context` breakdown commands, scoped model management from the `/model` selector, and better visibility into token usage by category.
- **Extension API enhancements**: Desire for extensions to call tools by name (Code Mode patterns), support for argument hints in prompt templates, and hooks for every provider‑bound payload.
- **Local‑first improvements**: XDG compliance, bash default timeout, and removal of auth requirements for local models show a push toward smoother offline/self‑hosted workflows.

## Developer Pain Points

- **WSL / cross‑platform issues**: Login hangs, file path mismatches on Windows drive roots, and Ctrl+V paste failures on Linux/X11 remain recurring friction points.
- **UI and terminal quirks**: Trailing spaces breaking copy, HOME/END key inaction, and the Escape “stuck in Working” state frustrate daily use. Sticky footer and off‑screen redraw fixes are welcome but incomplete.
- **Update and dependency breakage**: `pi update` failures (missing `@smithy/node-http-handler`), together with deprecated models (Together.ai), create unexpected upgrade blocks.
- **Provider compatibility**: Ghost models (mimo‑v2‑omni on Xiaomi), missing finish_reason for tool calls, and ignored `supportsDeveloperRole` flags cause silent failures. Users need better error messages and fallback logic.
- **Session data integrity**: UUID collisions, race conditions, and lost conversation history (now fixed in #6243) highlight fragility in session storage. The SQLite storage PR (#6227) aims to provide a more robust alternative.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-07-02

## Today's Highlights

The primary stable release **v0.19.4** shipped today, bringing a refreshed daemon documentation set and a configurable auto-compact threshold for the core engine. On the bug front, two high-priority issues have active PRs in progress: the mobile session-switching jank in Web Shell (P1) and the YOLO mode MCP invocation freeze (P2). The community also saw a cluster of feature requests around background automation—channel loops, cron job scheduling, and scheduled daemon capabilities—indicating a clear shift toward making Qwen Code a persistent agent infrastructure rather than just an interactive coding tool.

## Releases

**v0.19.4 (stable)**
- [Release v0.19.4](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.4)
  - **docs(daemon):** Refresh of daemon documentation covering recent PRs (wave 2) — aids users who rely on the daemon mode for background agents.
  - **feat(core):** Added a configurable auto-compact threshold and Stop — reduces memory bloat for long-running sessions without manual intervention.

**v0.19.4-nightly.20260702.46814e4f1**
- [Nightly build](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.4-nightly.20260702.46814e4f1)
  - **feat(cli):** Hardened daemon-managed channel worker — improves reliability for persistent channel adapters.
  - **fix(web-shell):** Deferred session creation until needed — reduces unnecessary resource allocation on startup.

---

## Hot Issues (10 selected)

### #6144 — [P2/Bug] Incorrect context window calculation
[Issue #6144](https://github.com/QwenLM/qwen-code/issues/6144)  
**What:** User configured a Qwen3-Coder instance with `ctx-size = 65536` and Q8_0/KV turbo caching, but the system misreports the usable context window.  
**Why it matters:** Affects any user relying on long-context models—incorrect accounting can cause unexpected truncation or OOM mid-session. 4 comments, 1 👍. Community is requesting a fix for the KV cache size accounting logic.

### #5979 — [P2/Bug] /auth config not persisted for new sessions
[Issue #5979](https://github.com/QwenLM/qwen-code/issues/5979)  
**What:** Changing model vendor config via `/auth` works for the current session, but new sessions still return 401 errors using old keys.  
**Why it matters:** Blocks multi-session workflows—users must re-authenticate after every restart. 4 comments, Chinese-language discussion. Status: in-review.

### #6181 — [P1/Bug] Mobile session switching jank
[Issue #6181](https://github.com/QwenLM/qwen-code/issues/6181)  
**What:** Web Shell on mobile freezes for seconds when switching sessions due to four overlapping performance issues: unrestricted sidebar polling, full-transcript sync render, uncompressed full-history load, and O(transcript) cost per render frame.  
**Why it matters:** P1 because it degrades the primary mobile UX. Detailed root-cause analysis provided by submitter. Status: ready-for-agent, autofix in-progress.

### #6175 — [P2/Bug] "Thought for 0s" and non-streaming thinking output
[Issue #6175](https://github.com/QwenLM/qwen-code/issues/6175)  
**What:** OpenAI-compatible models that emit `reasoning_content` show `Thought for 0s` always, and the thinking output no longer streams—appears all at once.  
**Why it matters:** Breaks the real-time reasoning display that users rely on for transparency. Status: needs-triage, welcome-pr.

### #6131 — [P2/Bug] YOLO mode cannot invoke MCP
[Issue #6131](https://github.com/QwenLM/qwen-code/issues/6131)  
**What:** In YOLO mode, MCP servers configured in `settings.json` cause the CLI to freeze—the permission system requires manual confirmation that YOLO mode should bypass.  
**Why it matters:** YOLO + MCP is a core power user workflow. The bug blocks headless automation entirely. Status: closed (PR merged).

### #6163 — [P2/Bug] npm package flagged as malicious by scanners
[Issue #6163](https://github.com/QwenLM/qwen-code/issues/6163)  
**What:** Published package flagged for a bundled IOC domain literal in auto-mode classifier prompt and a generated `postinstall.js` that applies patches.  
**Why it matters:** Security scanner false positives erode trust and can block CI/CD pipelines. The team will need to sanitize bundled content. Status: closed.

### #6062 — [P2/Enhancement] Patch runtime npm audit findings
[Issue #6062](https://github.com/QwenLM/qwen-code/issues/6062)  
**What:** Current dependency install reports high and critical npm audit findings in packages used by CLI and core runtime (shell parsing, git/worktree, archive extraction, HTTP/proxy).  
**Why it matters:** Supply chain risk—these are runtime dependencies, not just dev. 2 comments. Status: closed.

### #6097 — [P2/Enhancement] System prompt fixed overhead reaches ~22k tokens for minimal input
[Issue #6097](https://github.com/QwenLM/qwen-code/issues/6097)  
**What:** A single "hello" results in 22,300 tokens sent to the model due to system prompt + QWEN.md overhead. Signal-to-noise ratio: 0.2%.  
**Why it matters:** Token waste directly impacts cost and latency for narrow-domain agents. 2 comments. Status: closed.

### #6168 — [P2/Feature] Add DingTalk proactive send support for channel loops
[Issue #6168](https://github.com/QwenLM/qwen-code/issues/6168)  
**What:** Channel loops (#6073) require proactive-send support from adapters. DingTalk currently only supports reply via `sessionWebhook`, which expires—making scheduled loop delivery impossible.  
**Why it matters:** Blocks scheduled automation in the largest Chinese enterprise IM platform. Status: closed.

### #6167 — [P2/Feature] Make recurring cron/loop job expiration configurable
[Issue #6167](https://github.com/QwenLM/qwen-code/issues/6167)  
**What:** `RECURRING_MAX_AGE_MS` is hardcoded to 7 days in `cronScheduler.ts`. User wants configurable expiration or disable.  
**Why it matters:** Inflexible for long-running workflows like weekly reports or monitoring daemons. Status: open, feature-request.

---

## Key PR Progress (10 selected)

### #6183 — fix(web-shell): cut mobile session-switch jank
[PR #6183](https://github.com/QwenLM/qwen-code/pull/6183)  
**What:** Two P0 fixes for the laggy session switch on mobile: `MessageList` memoized, session-timeline signature computation only runs when visible; replay-first dispatch eliminates per-frame transcript concatenation.  
**Why it matters:** Directly addresses the most impactful mobile UX bug (Issue #6181). Author: qqqys.

### #6177 — fix(cli): skip MCP approval dialogs in YOLO mode
[PR #6177](https://github.com/QwenLM/qwen-code/pull/6177)  
**What:** MCP gating system now respects YOLO bypass in `permissionFlow.ts:needsConfirmation()`. The fix bridges the gap between the two independent permission systems.  
**Why it matters:** Unblocks headless MCP workflows. Merged today. Author: qwen-code-dev-bot.

### #6125 — feat(schedule): local always-on /schedule daemon
[PR #6125](https://github.com/QwenLM/qwen-code/pull/6125)  
**What:** Complete `/schedule` daemon with SKILL.md-backed task store, foreground daemon, CRUD commands, and cron-based execution without an open session. Covers four rollout phases.  
**Why it matters:** Major feature—turns Qwen Code into a persistent agent that runs routines autonomously. Status: open, need-discussion.

### #6079 — feat(cli): VP mode — inline thought expand on click + auto-hiding scrollbar
[PR #6079](https://github.com/QwenLM/qwen-code/pull/6079)  
**What:** Two VP-mode UX improvements: clicking a thought expands it inline (replaces full-screen modal), and a hidden scrollbar reduces visual noise.  
**Why it matters:** Competes with other TUI agents that already offer inline reasoning. Status: open.

### #6169 — fix(serve): keep skill slash commands available when ACP child is unavailable
[PR #6169](https://github.com/QwenLM/qwen-code/pull/6169)  
**What:** Web Shell caches last skills status and replays it when ACP child can't list skills. Adds daemon-side fallback via `sessionStore` and a `skillsCache` KV store.  
**Why it matters:** Prevents `/review` and similar commands from disappearing after transient failures. Author: wenshao.

### #6176 — fix: enforce plan mode over allowed tools
[PR #6176](https://github.com/QwenLM/qwen-code/pull/6176)  
**What:** Fixes permission priority so `allowedTools` cannot auto-execute write tools while plan mode is active. Read-only tools still bypass prompting; write-capable ones require confirmation.  
**Why it matters:** Prevents accidental destructive actions during planning. Author: yiliang114.

### #6106 — feat(core): add Tool(param:value) permission syntax
[PR #6106](https://github.com/QwenLM/qwen-code/pull/6106)  
**What:** New syntax `Tool(param:value)` for permission rules—e.g., `Agent(model:opus)` denies subagent launches requesting Opus model.  
**Why it matters:** Granular access control enables fine-grained cost governance and security policies. Status: open.

### #6173 — feat(scheduler): make recurring cron/loop job expiration configurable
[PR #6173](https://github.com/QwenLM/qwen-code/pull/6173)  
**What:** Adds `experimental.cronRecurringMaxAgeDays` setting (default 7, can be set to 0 for no expiration).  
**Why it matters:** Implements Issue #6167—flexibility for long-running schedules. Author: TianYuan1024.

### #6180 — fix(core): prevent subagent crash when ${hook_context} has no hook configured
[PR #6180](https://github.com/QwenLM/qwen-code/pull/6180)  
**What:** If subagent uses `${hook_context}` but no `SubagentStart` hook exists, the template engine previously crashed. Now returns empty string gracefully.  
**Why it matters:** Fixes a silent crash that could abort multi-agent workflows. Author: DennisYu07.

### #6171 — fix(web-shell): polish session timeline rail
[PR #6171](https://github.com/QwenLM/qwen-code/pull/6171)  
**What:** Tightened tick sizing, closer to message column, hidden in wide chat mode, removed browser-native hover tooltip. Acts as compact navigation aid instead of content competitor.  
**Why it matters:** UI polish that makes long-session navigation feel native. Author: callmeYe.

---

## Feature Request Trends

1. **Background Automation & Scheduled Agents** — The dominant theme. Multiple requests (DingTalk proactive send #6168, cron expiration config #6167, `/schedule` daemon #6125, channel loops #6073) all aim to make Qwen Code run autonomously without an active user session. The roadmap labels `roadmap/background-automation` and `roadmap/subagents-tools` confirm this is an intentional product direction.

2. **Security Hardening & Supply Chain Trust** — npm audit gates (#6062), scanner flag false positives (#6163), and fork PR precheck gating (#6178) reflect growing community concern about publishing trustworthy artifacts. Users want CI gates that prevent vulnerable or flagged packages from reaching release.

3. **Performance at Scale** — Context window miscalculation (#6144), system prompt token overhead (#6097), mobile session switch jank (#6181), and glob traversal optimizations (#6123) all point to scaling pain points. The community is pushing for better memory management and faster startups as session lengths and agent complexity grow.

4. **Tool & MCP Extensibility** — Parameter-level access control (#6106), MCP retry on network errors (#6048), YOLO + MCP compatibility (#6131), and MCP auto-discovery endpoints (#6048) show that power users want fine-grained control over tool execution without sacrificing automation.

---

## Developer Pain Points

- **Configuration Persistence Bugs:** The `/auth` 401 issue (#5979) is a recurring theme—users tired of re-authenticating after every restart or new session. The fix requires decoupling config storage from session lifecycle.

- **Mobile UX Degradation:** Multiple P1/P2 bugs for Web Shell on mobile (#6181, #6175, #6142) indicate the mobile experience lags significantly behind desktop. Session switching, thought streaming, and safe-area handling are all broken.

- **Security Scanner False Positives:** The npm package flagging (#6163) is a high-frustration issue because it's not a real vulnerability but blocks adoption in security-conscious organizations. Developers are calling for pre-release scanning and sanitization of bundled content.

- **UX Blocking Automation:** YOLO mode failing with MCP (#6131) and planned tool execution blocking write tools (#6176) show that permission gating systems have inconsistencies that frustrate headless/CI usage. Users want a single consistent bypass mechanism.

- **Unconfigurable Defaults:** Hardcoded constants (7-day cron expiration, 22k token system prompt overhead, fixed context window calculation) limit use cases. The community consistently asks for configurability or dynamic optimization rather than static limits.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-07-02

**Data source:** [Hmbown/CodeWhale](https://github.com/Hmbown/CodeWhale) *(Note: the tool formerly branded as DeepSeek TUI is now CodeWhale; community references still mix both names.)*

## Today’s Highlights

The v0.8.67 release cycle is in full swing with a massive constitution-first setup wizard landing in PR [#3861](https://github.com/Hmbown/CodeWhale/pull/3861) — the biggest single feature of the cycle. A long-standing community frustration around **project-scope instructions being overly restricted** is finally addressed by PR [#3892](https://github.com/Hmbown/CodeWhale/pull/3892), which auto-discovers `.codewhale/rules/` and `.claude/rules/` directories. Meanwhile, the dynamic MCP server infrastructure (PRs [#3869](https://github.com/Hmbown/CodeWhale/pull/3869), [#3866](https://github.com/Hmbown/CodeWhale/pull/3866)) opens a new capability for LLMs to spawn tools on the fly.

## Releases

None in the last 24 hours. The /web dependencies saw minor bumps (vitest 4.1.9, wrangler 4.106.0, mermaid 11.16.0) via automated PRs.

## Hot Issues

1. **[#3275](https://github.com/Hmbown/CodeWhale/issues/3275) – CodeWhale over-extends, self-dialogues, ignores user intent**  
   A regression from a previous fix. The agent enters a self-driven loop of proposing and executing without waiting for confirmation. 15 comments, heated community discussion.

2. **[#3406](https://github.com/Hmbown/CodeWhale/issues/3406) – Runtime posture card with constitution boundary**  
   Maintainer Hmbown proposes a visible security posture selector (ask-first / normal agent / high-trust local) that constitution creators cannot silently override. 14 comments.

3. **[#3867](https://github.com/Hmbown/CodeWhale/issues/3867) – Project-scope instructions overly denied**  
   Since v0.8.8, the `instructions` config key has been hard-blocked at project scope with no glob, rules-directory auto-discovery, or trust awareness. 7 comments, directly driving PR [#3892](https://github.com/Hmbown/CodeWhale/pull/3892).

4. **[#3793](https://github.com/Hmbown/CodeWhale/issues/3793) – Guided localized constitution creator**  
   The maintainer envisions a language-first, guided-plus-open-canvas flow for creating personal constitutions, with strict separation from runtime security settings. 11 comments.

5. **[#3883](https://github.com/Hmbown/CodeWhale/issues/3883) – YOLO mode still prompts through auto-review gates**  
   In workspace v0.8.66, YOLO mode continues to surface approval takeovers for ordinary orchestration, making it feel broken. Closed quickly, but highlights lingering UX confusion.

6. **[#3830](https://github.com/Hmbown/CodeWhale/issues/3830) – Provider & model route manager needed**  
   `/provider` and `/model` currently feel like flat config editors. This issue asks for a dedicated route manager for multiple provider accounts and model selection. 6 comments.

7. **[#3880](https://github.com/Hmbown/CodeWhale/issues/3880) – DSML interrupt task not merged into release**  
   A Windows user reports that the DSML interrupt fix is missing from the latest v0.8.66 binary. 3 comments, potential release blunder.

8. **[#3859](https://github.com/Hmbown/CodeWhale/issues/3859) – “Ctrl+B backgrounds this command” is misleading**  
   The TUI hint implies true backgrounding, but in practice bash commands cannot be backgrounded for LLM continuation. Closed with documentation fix.

9. **[#3837](https://github.com/Hmbown/CodeWhale/issues/3837) – Agents sidebar does not reconcile sub-agent completion**  
   Sub-agent lifecycle state goes stale in the sidebar, causing confusion during cancellation workflows. 3 comments.

10. **[#3864](https://github.com/Hmbown/CodeWhale/issues/3864) – Sub-agent state persists to `.deepseek/` instead of `.codewhale/`**  
    A lingering rebranding path bug — worker state JSON still uses the old directory. Easy fix but a reminder of incomplete migration.

## Key PR Progress

1. **[#3861](https://github.com/Hmbown/CodeWhale/pull/3861) – v0.8.67 constitution-first setup (Hmbown)**  
    The marquee PR of the cycle: model-assisted onboarding, runtime posture card, and cleanup. Not yet tagged but represents the entire setup wizard spine.

2. **[#3892](https://github.com/Hmbown/CodeWhale/pull/3892) – Auto-discover `.codewhale/rules/` and `.claude/rules/` (yekern)**  
    Directly closes [#3867](https://github.com/Hmbown/CodeWhale/issues/3867). Adds rules-directory scanning to `load_project_context()` — a major win for multi-project workflows.

3. **[#3869](https://github.com/Hmbown/CodeWhale/pull/3869) – Dynamic MCP server infrastructure in McpPool (bistack)**  
    Foundation for runtime-started MCP servers. Adds `dynamic_servers` field with `parking_lot::RwLock`, enabling future `start_mcp_server` tool.

4. **[#3866](https://github.com/Hmbown/CodeWhale/pull/3866) – LLM can start MCP servers from chat context (bistack)**  
    Companion PR that introduces the actual tool. Supports both stdio and HTTP transports. Combined with [#3869](https://github.com/Hmbown/CodeWhale/pull/3869), this is a powerful new capability.

5. **[#3643](https://github.com/Hmbown/CodeWhale/pull/3643) – Setup summary wizard step for MCP/skills/plugins (cy2311)**  
    Implements the first step of the v0.8.67 setup wizard — scrollable TUI modal with status info. 5 unit tests, part of [#3407](https://github.com/Hmbown/CodeWhale/issues/3407).

6. **[#3578](https://github.com/Hmbown/CodeWhale/pull/3578) – Preserve Windows SDK env roots for shell (nightt5879)**  
    Fixes [#3572](https://github.com/Hmbown/CodeWhale/issues/3572) by recovering Windows SDK path variables from registry before shell launch. Important for Windows developer workflows.

7. **[#3574](https://github.com/Hmbown/CodeWhale/pull/3574) – Fix context window edge cases (nightt5879)**  
    Follow-up to [#3573](https://github.com/Hmbown/CodeWhale/pull/3573) ensuring `max_tokens` stays positive when `ContextBudget` clamps output to zero, and provider-switch rollback restores override.

8. **[#3881](https://github.com/Hmbown/CodeWhale/pull/3881) – Prune localization QA metadata (nightt5879)**  
    Cleanup removing unused localization QA structs while keeping runtime helpers. Fixes [#3857](https://github.com/Hmbown/CodeWhale/issues/3857).

9. **[#3873](https://github.com/Hmbown/CodeWhale/pull/3873) – Remove unused execpolicy amend module (cyq1017)**  
    Drops dead code and the direct fd-lock dependency from `codewhale-tui`. Part of ongoing cleanup.

10. **[#3870](https://github.com/Hmbown/CodeWhale/pull/3870) – Refactor McpTool storage to `Arc<McpTool>` (bistack)**  
    Prerequisite for dynamic MCP server support. Changes `Vec<McpTool>` to `Vec<Arc<McpTool>>` in `McpConnection`, enabling shared ownership across `Arc<RwLock<>>`.

## Feature Request Trends

- **Constitution-first setup & security posture** — The majority of v0.8.67 issues revolve around making first-run onboarding a guided ritual with a user-defined constitution and a visible runtime security posture (ask-first / normal / high-trust). Language localization and structured constitution creation are key sub-themes.

- **Provider & model route management** — Users want a dedicated UI surface for managing multiple API provider accounts and selecting models, rather than editing config files.

- **Project-scope rules & instruction discovery** — The hard block on project-scope `instructions` config sparked strong demand for glob patterns and auto-discovery of `.codewhale/rules/` and `.claude/rules/` directories (now being addressed).

- **Agent orchestration improvements** — Real-time sub-agent lifecycle updates, tool catalog parity between parent and child agents, and reliable backgrounding of shell commands are recurring asks.

- **Dynamic MCP server spawning** – Allowing the LLM itself to start MCP servers from conversation context (already in progress via bistack's PRs).

## Developer Pain Points

- **Agent over-autonomy** – Issue [#3275](https://github.com/Hmbown/CodeWhale/issues/3275) reflects a deep frustration: CodeWhale ignores user intent, enters self-driven loops, and modifies code without confirmation. This is the single most upvoted complaint.

- **Confusing mode/approval policies** – YOLO mode, approval_mode, trust_mode, auto_approve — the four overlapping knobs create a UX nightmare. Issue [#3736](https://github.com/Hmbown/CodeWhale/issues/3736) (closed) captured the structural problem, and [#3883](https://github.com/Hmbown/CodeWhale/issues/3883) shows it persists in practice.

- **Rebranding migration incomplete** – Sub-agent state writes to `.deepseek/` instead of `.codewhale/` ([#3864](https://github.com/Hmbown/CodeWhale/issues/3864)). Minor but a tripping hazard for anyone migrating directories.

- **Windows-specific bugs** – Copy/paste context menu overrides the entire GUI ([#3868](https://github.com/Hmbown/CodeWhale/issues/3868)), DSML interrupt not merged in release binary ([#3880](https://github.com/Hmbown/CodeWhale/issues/3880)), and SDK environment recovery ([#3578](https://github.com/Hmbown/CodeWhale/pull/3578)) — Windows users face a disproportionate share of pain.

- **Stale UI state** – The Agents sidebar not updating sub-agent completion ([#3837](https://github.com/Hmbown/CodeWhale/issues/3837)) and the misleading "background this command" hint ([#3859](https://github.com/Hmbown/CodeWhale/issues/3859)) erode trust in real-time feedback.

---

*Digest generated from the [Hmbown/CodeWhale](https://github.com/Hmbown/CodeWhale) project. Not affiliated with DeepSeek or Hmbown.*

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*