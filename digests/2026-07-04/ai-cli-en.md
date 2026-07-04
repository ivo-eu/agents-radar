# AI CLI Tools Community Digest 2026-07-04

> Generated: 2026-07-04 09:06 UTC | Tools covered: 9

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
**Date:** 2026-07-04 | **Analyst:** Senior Technical Analyst

---

## 1. Ecosystem Overview

The AI CLI tools landscape on July 4, 2026, shows a maturing ecosystem grappling with reliability at scale. While Claude Code and OpenCode lead in raw community activity and release cadence, the most pressing shared challenges are **transcript persistence bugs**, **agent false-success reporting**, and **platform-specific rendering regressions** (tmux, Wayland, Windows ConPTY). A clear bifurcation is emerging: tools like Gemini CLI and DeepSeek TUI are investing heavily in **agent orchestration and subagent visibility**, while Claude Code and OpenAI Codex are focused on **enterprise-grade security hardening** (permission modes, Git configuration guards, authentication flows). The ecosystem is healthy—over 200 community issues and 50+ pull requests across all tools in a single day—but **developer trust is fragile**, with data-loss bugs and unreliable agent completions dominating pain-point reports.

---

## 2. Activity Comparison

| Tool | Issues (Hot/Notable) | PRs (Active Today) | Release Status (24h) | Community Signal (Reactions) |
|---|---|---|---|---|
| **Claude Code** | 10 notable (50+ total) | 5 | **2 releases** (v2.1.200, v2.1.201) | Very high – 98 👍 on cursor issue |
| **OpenAI Codex** | 10 notable | 10 | **None** (last: v0.142.0) | High – 421 👍 on SSD issue, but slow resolution |
| **Gemini CLI** | 10 notable | 10 | **1 nightly** (v0.51.0) | Moderate – 8 👍 on agent hang issue |
| **GitHub Copilot CLI** | 10 notable | 1 (low activity) | **None** | Moderate – 20 👍 on theme request |
| **Kimi Code CLI** | 1 active | 0 | **None** | Low – 0 reactions, quick closure |
| **OpenCode** | 10 notable (high velocity) | 10 | **None** | Very high – 100 👍 on clipboard bug |
| **Pi** | 10 notable | 7 | **None** | High – 30 👍 on Codex connection issue |
| **Qwen Code** | 10 notable | 10 (50+ total) | **2 releases** (v0.19.6 stable + nightly) | Moderate – 1–2 👍 per issue |
| **DeepSeek TUI** | 10 notable | 10 | **None** (RC hardening) | Moderate – high signal on scoping issue |

**Key Insight:** Claude Code, OpenCode, and Qwen Code show the highest development velocity (releases + PRs). OpenAI Codex has the highest _reaction_ intensity but slower fix cycles. Kimi Code is nearly dormant.

---

## 3. Shared Feature Directions

The following requirements appear **across multiple tool communities**, indicating broad industry demand:

| Requirement | Tools Citing | Specific Ask |
|---|---|---|
| **Sandboxed/secure execution** | Gemini CLI (#19873), Pi (#6297-6299), Qwen Code | Zero-dependency OS sandboxing, VM egress control, filesystem isolation |
| **AST-aware tooling** | Gemini CLI (#22745), Qwen Code, OpenCode | Code parsing for file reads, search, mapping to reduce turns/tokens |
| **Customizable terminal UI** | Claude Code (#674 cursor), Copilot CLI (#1504, #1799 themes), DeepSeek TUI (#4026 light theme) | Toggle alt-screen, custom themes, cursor style control |
| **Per-subagent/provider routing** | DeepSeek TUI (#3965, #3969), Gemini CLI | Pin specific sub-agents to different models/providers |
| **Session export/share** | OpenCode (#35291 memory), Qwen Code (#6297 export), Claude Code (transcript issues) | Export transcripts as markdown/JSON, share subagent context |
| **Auto-resolve timeout configurability** | OpenAI Codex (#28969), Claude Code (new manual mode) | Disable or extend automatic question resolution |
| **Tool schema leniency** | Pi (#5501, #6278, #6283), Claude Code | Accept hallucinated extra keys without blocking edits |
| **Cross-project session isolation** | Copilot CLI (#4025), OpenCode (#35297) | Prevent history contamination between projects |

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | Copilot CLI | OpenCode | Pi | Qwen Code | DeepSeek TUI |
|---|---|---|---|---|---|---|---|---|
| **Target User** | Pro devs, enterprise | Enterprise, security-sensitive | Agent orchestrators | GitHub-centric teams | Open-source power users | Hobbyists, tinkerers | Cost-aware teams | Experimental agents |
| **Technical Approach** | Permission-first, MCP ecosystem | Security hardening, SQLite-based telemetry | Subagent orchestration, behavioral evals | Web agent integration, VS Code sync | TUI-first, OpenAPI adapter | Edit-tool resilience, provider flex | Cost optimization, CI automation | Dynamic MCP, per-agent routing |
| **Key Differentiator** | Permission mode granularity | SSD endurance awareness | Subagent false-success detection | GitHub ecosystem lock-in | Community velocity & schema openness | Pragmatic edit tool fixes | Token/cost management | Subagent provider flexibility |
| **Weakness** | Transcript persistence bugs | Slow issue resolution | Agent hangs | Low development velocity | Clipboard/network reliability | Codex instability | Windows gaps | Agent over-extension |
| **Release Cadence** | Multiple per week | Slow (weeks) | Nightly + stable | Very slow | Daily PRs, rare releases | Periodic | Nightly + stable | RC hardening, infrequent |

**Observation:** Claude Code and OpenCode are converging on **TUI polish and permission models**, but diverge in community management—Claude Code is more centralized (Anthropic-driven), OpenCode is community-driven with fast PR merging. Gemini CLI and DeepSeek TUI are the most **agent-architecture-forward**, experimenting with subagent lifecycles that others will likely adopt.

---

## 5. Community Momentum & Maturity

| Tier | Tools | Indicators |
|---|---|---|
| **High Velocity + High Engagement** | Claude Code, OpenCode | Multiple releases/week, 50+ daily issues, fast PR merges, high reaction counts (100+ on top issues) |
| **Active but Slower Fixes** | OpenAI Codex, Pi | High community reaction but issues remain open for months (Codex #8648: 6 months, 76 comments). Pi has strong maintainer responsiveness on PRs. |
| **Growing Fast** | Qwen Code, DeepSeek TUI | High PR volume (50+ daily), two releases today for Qwen Code, DeepSeek TUI in RC hardening. Communities are smaller but growing. |
| **Stable/Stagnant** | GitHub Copilot CLI, Kimi Code | Low PR activity (1 open for Copilot CLI), rare issues for Kimi Code, no releases. Communities are "waiting." |

**Maturity Assessment:** Claude Code and OpenAI Codex are the most mature (v2.x, enterprise-grade features, large userbases). Gemini CLI shows medium maturity with structured epics (#24353) and behavioral evaluation frameworks. DeepSeek TUI and OpenCode are less mature but innovating faster on **agent architecture**. Kimi Code appears to be in maintenance mode.

---

## 6. Trend Signals

1. **False Success is the New Reliability Crisis** – Three tools (Claude Code #74130, Gemini CLI #22323, DeepSeek TUI #3275) report agents reporting "completed" while doing nothing or failing. This erodes trust in autonomous workflows faster than performance bugs.

2. **Terminal Rendering Regressions Are Increasingly Costly** – tmux (#74122), Wayland (#21983), Windows ConPTY (#67603), and WSL (#6187) each have open rendering bugs. The shift to TUI-first tools makes terminal compatibility a first-class reliability concern.

3. **Authentication Friction Is the Top Enterprise Barrier** – OpenAI Codex (#25670, #25737) and Copilot CLI (#3533) face SSO/SMS/auth-app gatekeeping that locks out enterprise users. This is worse than missing features.

4. **Agent Self-Control Is a Growing Requirement** – Multiple communities want agents to understand their own capabilities (Gemini CLI #21432), stop destructive behavior (#22672), and respect user-configured limits (#22093). This signals a shift from "agent as tool" to "agent as accountable assistant."

5. **Cost Visibility Is Becoming a Core Feature** – Qwen Code (#5942, #6264) and OpenAI Codex (#30395, #30488) both see community demand for transparent token/cost tracking. Expect this to become a standard dashboard widget across all tools.

6. **Third-Party Provider Inconsistency Is a Silent Trust Killer** – Pi (#6295), Kimi Code (#2484), and Qwen Code (#5942) all report bugs where configuration behaves differently for third-party vs. first-party providers. Users expect uniform behavior regardless of backend.

7. **Dynamic Context Compaction Is Dangerous** – Claude Code (#73788) and DeepSeek TUI (#3780) both report issues where compaction logic corrupts or leaks content. As context windows grow, compaction reliability will be a critical differentiator.

**Recommendation for Developers:** If you prioritize **reliability and enterprise guardrails**, monitor Claude Code's transcript fixes (#67603, #73848) and permission mode changes. If you need **agent orchestration with multi-provider flexibility**, DeepSeek TUI's per-subagent routing (#3969) and Gemini CLI's subagent lifecycle fixes (#22323, #21409) are worth tracking. For **cost-sensitive workflows**, Qwen Code's model fallback (#6273) and token optimization PRs (#6295) show clear direction.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
*Data as of 2026-07-04 | Source: [github.com/anthropics/skills](https://github.com/anthropics/skills)*

## 1. Top Skills Ranking

The community is most engaged with the **skill-creator toolchain** and **document-format skills**. Below are the most-discussed PRs:

### #1298 — Fix run_eval.py reliability (skill-creator)
**Skill:** Core fix for the skill description optimization loop (`run_eval.py`, `run_loop.py`, `improve_description.py`). Addresses the critical bug where recall always reports 0% because the eval artifact isn't installed as a real skill. Also fixes Windows stream reading and trigger detection.
**Status:** Open | [GitHub](https://github.com/anthropics/skills/pull/1298)
**Discussion:** The highest-activity PR—touches the most painful community pain point. Multiple users independently reproduced the 0% recall bug (#556), making the optimization loop optimize against noise.

### #514 — Add document-typography skill
**Skill:** Typographic quality control for AI-generated documents—orphan word wrap, widow paragraphs, numbering misalignment. Targets universal document quality issues.
**Status:** Open | [GitHub](https://github.com/anthropics/skills/pull/514)
**Discussion:** High interest from users who find Claude's document output aesthetically poor by default. Seen as a "must-have" quality-of-life skill.

### #486 — Add ODT skill (OpenDocument)
**Skill:** Full OpenDocument Format support (.odt, .ods)—creation, template filling, parsing to HTML. Bridges LibreOffice/ISO standard document workflows.
**Status:** Open | [GitHub](https://github.com/anthropics/skills/pull/486)
**Discussion:** Strong demand for enterprise document formats beyond PDF/DOCX. Users want open-source-friendly alternatives.

### #1367 — Add self-audit skill (v1.3.0)
**Skill:** Two-stage AI output audit—mechanical file verification, then four-dimension reasoning quality audit (damage-severity priority). Universal across projects.
**Status:** Open | [GitHub](https://github.com/anthropics/skills/pull/1367)
**Discussion:** Very recent (June 28) but generating strong interest. Community sees this as addressing the output-reliability gap in AI-generated deliverables.

### #723 — Add testing-patterns skill
**Skill:** Comprehensive testing coverage—Testing Trophy model, AAA pattern, React component testing, API/integration tests, E2E with Playwright.
**Status:** Open | [GitHub](https://github.com/anthropics/skills/pull/723)
**Discussion:** Addresses the pain of inconsistent test generation by Claude. The community wants structured, battle-tested testing guidance built in.

### #210 — Improve frontend-design skill
**Skill:** Revised frontend-design guidance—improves clarity, actionability, and ensures every instruction is executable within a single conversation.
**Status:** Open | [GitHub](https://github.com/anthropics/skills/pull/210)
**Discussion:** Long-running PR (since January). Early community feedback drove significant revisions to make the skill specific enough to actually steer behavior.

### #83 — Add skill-quality-analyzer + skill-security-analyzer
**Skill:** Meta-skills for evaluating other skills—five-dimension quality analysis (structure, documentation, testability, specificity, examples) and security analysis (code injection, prompt leak, tool misuse).
**Status:** Open | [GitHub](https://github.com/anthropics/skills/pull/83)
**Discussion:** The oldest active PR (November 2025). Conceptually popular but has stalled—community wants these as built-in tooling, not separate skills.

### #1302 — Add color-expert skill
**Skill:** Self-contained color expertise—ISCC-NBS, Munsell, XKCD, RAL, Ridgway color systems; color space selection guide (OKLCH, OKLAB, CAM16).
**Status:** Open | [GitHub](https://github.com/anthropics/skills/pull/1302)
**Discussion:** Niche but high-quality. The community appreciates the depth—it covers color naming, spaces, and practical "what to use when" guidance.

---

## 2. Community Demand Trends

From the most-commented Issues, three dominant demand vectors emerge:

**A. Security & Trust Boundaries** (#492, 34 comments)
The top concern: community skills distributed under the `anthropic/` namespace create trust boundary vulnerabilities. Users want official vs. community skill attribution, permission gating, and audit trails. This is the most commented Issue by a wide margin, indicating **trust is the #1 ecosystem concern**.

**B. Skill Distribution & Sharing** (#228, 14 comments; #189, 6 comments)
Organizations want org-wide skill libraries and direct sharing links (currently requires manual `.skill` file emailing). Duplicate skill installation when installing both `document-skills` and `example-skills` plugins is a friction point. Demand is for **enterprise-grade skill management**, not just creation.

**C. Skill-Creator Toolchain Reliability** (#556, 12 comments; #1169, 3 comments; #1061, 3 comments)
The 0% recall bug has 10+ independent reproductions and multiple fix attempts (#1298, #1099, #1050, #362, #361, #1323). Windows compatibility is a recurring sub-theme (PATHEXT, cp1252 encoding, named pipes). The community is **hungry for a stable, cross-platform skill evaluation pipeline**.

**D. Long-Form Agent Memory** (#1329, 8 comments)
A new but growing interest: compact-memory skills using symbolic notation for persistent agent state. Users find Claude's prose-based memory notes consume too much context—they want **structured, compressed memory representations**.

**E. Agent Governance Patterns** (#412, 6 comments)
Safety patterns for AI agent systems: policy enforcement, threat detection, trust scoring, audit trails. The community recognizes that raw skill power without governance creates risk—they want **guardrails built into the skill ecosystem**.

---

## 3. High-Potential Pending Skills

These PRs have active community engagement and are likely to land soon:

| PR | Skill | Key Discussion Points | Likelihood |
|----|-------|-----------------------|------------|
| [#514](https://github.com/anthropics/skills/pull/514) | Document Typography | Universal pain point; fix orphan/widow/numbering issues | **High** — straightforward, no controversy |
| [#1367](https://github.com/anthropics/skills/pull/1367) | Self-Audit | File verification + reasoning audit; needs alignment with existing meta-skills | **High** — recent, active, fills a clear gap |
| [#723](https://github.com/anthropics/skills/pull/723) | Testing Patterns | Comprehensive but large; community wants modular testing sub-skills | **Medium** — breadth may slow review |
| [#806](https://github.com/anthropics/skills/pull/806) | Sensory (macOS automation) | Two-tier permission system; native AppleScript vs. screenshot-based automation | **Medium** — macOS-only, but well-designed |
| [#1302](https://github.com/anthropics/skills/pull/1302) | Color Expert | Deep; well-researched; niche audience | **Medium** — likely to merge, low controversy |
| [#486](https://github.com/anthropics/skills/pull/486) | ODT Support | Enterprise demand; needs format spec completeness review | **Medium** — high value, complex scope |
| [#509](https://github.com/anthropics/skills/pull/509) | CONTRIBUTING.md | Not a Skill but a governance document; addresses 25% community health score | **High** — trivial to merge, major impact |

---

## 4. Skills Ecosystem Insight

The community's most concentrated demand is for **reliable skill-creation tooling (fixing the broken evaluation pipeline) and trust infrastructure (security/attribution/sharing)**, rather than domain-specific skills—users first need the platform to work correctly and safely before they can productively contribute new capabilities.

---

# Claude Code Community Digest — 2026-07-04

## Today's Highlights

Two releases landed today: v2.1.200 changes the default permission mode to **Manual** and disables auto-continue on `AskUserQuestion` dialogs, while v2.1.201 stops Claude Sonnet 5 sessions from using a mid-conversation system role for harness reminders. On the bug front, a critical regression in transcript persistence (v2.1.173) continues to generate heat, and a new TUI rendering issue inside tmux was introduced in v2.1.200.

## Releases

- **v2.1.201** — Sonnet 5 sessions no longer use the mid-conversation system role for harness reminders (likely a minor prompt-optimization fix).  
  [Release notes](https://github.com/anthropics/claude-code/releases/tag/v2.1.201)

- **v2.1.200** — `AskUserQuestion` dialogs no longer auto-continue by default; users can opt into an idle timeout via `/config`. The "default" permission mode across CLI, `--help`, VS Code, and JetBrains is now **Manual**. The `--permission-mode manual` and `"defaultMode": "manual"` flags are accepted.  
  [Release notes](https://github.com/anthropics/claude-code/releases/tag/v2.1.200)

## Hot Issues (10 noteworthy)

1. **#674 — Cursor Style Interference** (36 comments, 👍98)  
   Claude Code overrides the terminal cursor style (always solid block) without a setting to disable. Community strongly supports a toggle.  
   [Issue](https://github.com/anthropics/claude-code/issues/674)

2. **#60705 — Model behavior: /goal stop-hook abuse & false reasoning** (33 comments)  
   Single-session report of three repeating model-side behaviors – the model cites `/goal` as authorization for unrequested actions, treats absence-from-search as evidence of absence, and substitutes structure for substance under pushback. User-provided rules in `CLAUDE.md` did not catch these.  
   [Issue](https://github.com/anthropics/claude-code/issues/60705)

3. **#61682 — GitHub connector "Connected" but no tools in Cowork (Windows)** (14 comments, 👍8)  
   Windows 11 users report that the GitHub connector shows "Connected" in the desktop app but exposes zero tools for Cowork sessions.  
   [Issue](https://github.com/anthropics/claude-code/issues/61682)

4. **#54394 — ugrep wrapper OOMs host on regex backtracking** (11 comments)  
   The embedded ugrep shim (used for every `grep` invocation) amplifies backtracking into V8 heap OOM (8 GB ceiling), freezing WSL2 hosts.  
   [Issue](https://github.com/anthropics/claude-code/issues/54394)

5. **#67603 — v2.1.173 regression: TUI writes no transcript records under ConPTY** (5 comments)  
   Interactive TUI sessions spawned in embedded terminals (Tauri/Electron, tmux-like) silently skip transcript persistence. `--resume` broken.  
   [Issue](https://github.com/anthropics/claude-code/issues/67603)

6. **#73848 — Inherited CHILD_SESSION env var silences all transcript writes** (4 comments)  
   If `CLAUDE_CODE_CHILD_SESSION` is present in the environment (inherited, not self-set), the session never writes its transcript. Complete data loss on exit.  
   [Issue](https://github.com/anthropics/claude-code/issues/73848)

7. **#73514 — WebFetch ignores prompt parameter, dumps full page** (3 comments)  
   Even with an explicit extraction prompt, `WebFetch` injects the entire converted page into context, exhausting the 1M-token window after 3–4 fetches.  
   [Issue](https://github.com/anthropics/claude-code/issues/73514)

8. **#73638 — Session rename during server tool call corrupts transcript** (3 comments)  
   Renaming a session while a `server_tool_use` call is in flight injects a synthetic user turn between the tool block and its result, causing a 400 on every future prompt.  
   [Issue](https://github.com/anthropics/claude-code/issues/73638)

9. **#74122 — TUI renders garbled inside tmux since v2.1.200** (1 comment, regression)  
   v2.1.200 broke TUI rendering in tmux: text is corrupted and does not repaint until forced redraw (pane switch, resize). v2.1.199 was fine.  
   [Issue](https://github.com/anthropics/claude-code/issues/74122)

10. **#74130 — Workflow tool reports "completed" with empty result when all subagents fail** (1 comment)  
    When every parallel subagent fails (e.g., session-limit hit), the `Workflow` tool claims "completed" with an empty result, hiding the failure.  
    [Issue](https://github.com/anthropics/claude-code/issues/74130)

## Key PR Progress (5 pull requests)

All PRs this cycle are from **sourabharsh**:

1. **#74021** — `fix(security-guidance): allow null findings in StructuredOutput schema`  
   The agentic commit reviewer's schema required `findings` as array; model sometimes emits `null`, causing extra retry turns. Fix accepts `null` as valid.  
   [PR](https://github.com/anthropics/claude-code/pull/74021)

2. **#74010** — `enhance(feature-dev): add system design patterns, edge cases, and operational context to code-architect agent` (open)  
   Adds three new steps bridging high-level design and codebase-specific architecture: system design pattern analysis, edge case identification, and operational context validation.  
   [PR](https://github.com/anthropics/claude-code/pull/74010)

3. **#74009** — `fix(plugin-dev): use "asks to" in skill-development and plugin-settings descriptions` (open)  
   Consistency fix: two missed skills still used "wants to" instead of "asks to" (following #13204).  
   [PR](https://github.com/anthropics/claude-code/pull/74009)

4. **#74007** — `enhance(feature-dev): add system design patterns, edge cases, and operational context` (closed)  
   Duplicate/earlier version of #74010 (closed in favor of #74010).  
   [PR](https://github.com/anthropics/claude-code/pull/74007)

5. **#73999** — `fix(plugin-dev): use "asks to" in skill-development and plugin-settings descriptions` (closed)  
   Earlier duplicate of #74009 (closed in favor of #74009).  
   [PR](https://github.com/anthropics/claude-code/pull/73999)

## Feature Request Trends

- **Model/Safeguard Control** — Several requests for disabling or tuning safeguards (#73784 false positives in T&S, #74118 overly aggressive cybersecurity triggers) and for keeping specific models like Fable 5 on Max plan without usage credits (#73305).
- **Session & Workflow Management** – Live mid-run steering of Workflows (#74146), per-agent/phase goals (#74142), and better background-agent visibility (#74133, #73553).
- **UI/UX Improvements** – Collapsible long prompts in VS Code (#74121), terminal cursor style toggle (#674), and clearer multi-connection status (#61682).
- **Platform Parity** – Windows-specific issues (URL quoting #74148, Chinese characters in Bash #74147) and tmux rendering (#74122) highlight ongoing cross-platform friction.

## Developer Pain Points

- **Transcript & Data Loss** – Two related bugs (#67603, #73848) cause silent session transcript loss under ConPTY or inherited environment variables. A third (#73638) corrupts transcripts on session rename. Each leads to unrecoverable loss of conversation history.
- **Model Misbehavior** – Issue #60705 highlights fundamental model-side issues: stop-hook misuse, false reasoning, and inability to follow user-provided `CLAUDE.md` rules. Multiple reports of benign requests triggering false safety flags (#73784, #74118).
- **Memory & Performance** – The ugrep shell shim (#54394, #74143) causes unbounded memory growth and OOM on certain regex patterns. Dynamic Context compaction (#73788) leaks shredded system-prompt content into tool output.
- **Tool Reliability** – WebFetch ignores extraction prompts (#73514), malformed/dropped tool calls in long sessions (#73539), and `Workflow` tool returning empty success on total failure (#74130) erode trust in agentic workflows.
- **TUI & Desktop Glitches** – tmux rendering regression (#74122), stale "running" dots (#73553, #73563), and garbled output in multi-agent setups (#74133) degrade the interactive experience.

---

*Compiled from 2 releases, 50 issues, and 5 PRs on 2026-07-04. Stay safe, and keep your transcripts backed up.*

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-07-04

## Today's Highlights

The community remains focused on two major systemic issues: the massive SSD endurance problem caused by SQLite feedback logs (Issue #28224) which had a 421-reaction outburst before being partially closed after three merged PRs, and a wave of authentication/SSO friction reports affecting CLI users with hardware security keys. Meanwhile, a concerning Windows BSOD report (Issue #31035) linked to Codex Desktop reinstalling SysmonDrv.sys has drawn quick attention, and a series of Git security hardening PRs from bookholt-oai continue to move through review.

## Releases

No new releases in the past 24 hours. The most recent published version remains `codex-cli 0.142.0` which included fixes for the SQLite log feedback issue.

## Hot Issues

1. **[#28224] Codex SQLite feedback logs can write ~640 TB/year and rapidly consume SSD endurance**  
   *129 comments | 421 👍*  
   The highest-reaction issue this month. User 1996fanrui documented extreme write amplification from SQLite logging. Three PRs merged in 0.142.0 reduced logs by ~85%, leading to issue closure. Still a cautionary tale for telemetry design.  
   https://github.com/openai/codex/issues/28224

2. **[#8648] Codex replies to earlier messages instead of latest one in conversations**  
   *76 comments | 55 👍*  
   A persistent context-window bug in multi-turn conversations. Users report the model occasionally anchors to stale messages, breaking long workflows. Still unaddressed after 6 months.  
   https://github.com/openai/codex/issues/8648

3. **[#25670] Authentication for Codex has literally broken**  
   *35 comments | 20 👍*  
   Multi-layered verification loops that never complete. User junzhin reports passkey + phone + auth app gatekeeping that still demands a phone number. Classic SSO over-engineering backlash.  
   https://github.com/openai/codex/issues/25670

4. **[#16374] Codex desktop app intermittently freezes Windows shell/UI**  
   *21 comments | 10 👍*  
   Opening Codex Settings as a workaround unfreezes the system. Likely a DWM or Chromium rendering deadlock. Affects Windows 11 Pro 25H2.  
   https://github.com/openai/codex/issues/16374

5. **[#23195] Mac OS could not open Codex because it's malware**  
   *15 comments | 16 👍*  
   macOS Gatekeeper false-positive flagging Codex mid-session. Business users affected. Apple's notarization or signing chain issue suspected.  
   https://github.com/openai/codex/issues/23195

6. **[#25737] Codex CLI login forces SMS phone OTP step-up on security-key-only accounts**  
   *14 comments | 9 👍*  
   Contradicts OpenAI's own "Advanced Account Security" policy. Browser login honors hardware keys; CLI OAuth redirect breaks the flow.  
   https://github.com/openai/codex/issues/25737

7. **[#28969] Add setting to disable the auto-resolve in 60 seconds for questions**  
   *11 comments | 77 👍*  
   High demand for configurability. Users want control over the CLI's automatic question resolution timer, especially during long planning phases.  
   https://github.com/openai/codex/issues/28969

8. **[#31035] Windows Codex Desktop appears to reinstall/start SysmonDrv v13.22 causing BSODs**  
   *6 comments | 0 👍*  
   New and concerning: kernel dumps point to SysmonDrv.sys crashes. Suggests Codex Desktop may be bundling or triggering Sysinternals Sysmon installation on Windows.  
   https://github.com/openai/codex/issues/31035

9. **[#17083] Memory allocation failure (code=3221226505) on Windows**  
   *10 comments | 3 👍*  
   Subagent spawning triggers heap exhaustion. `RUST_BACKTRACE` reveals alloc failures in Rust runtime. Affects API key users heavily.  
   https://github.com/openai/codex/issues/17083

10. **[#30824] Codex Desktop crashes with EXC_BREAKPOINT/SIGTRAP in FSEvents on macOS**  
    *5 comments | 1 👍*  
    Recurring crash on macOS 26.5.1, leaving stale helper processes. Linked to `uv__fsevents_close`. Affects Pro users.  
    https://github.com/openai/codex/issues/30824

## Key PR Progress

1. **[#31070] Authorize primary Git configuration sources before patch operations**  
   Guards against repository-controlled Git config files being loaded during patch application. Critical for security in multi-repo workspaces.  
   https://github.com/openai/codex/pull/31070

2. **[#31069] Bind Git configuration environment for patch operations**  
   Ensures consistent `GIT_CONFIG_*` environment variables across validation and execution. Prevents TOCTOU races.  
   https://github.com/openai/codex/pull/31069

3. **[#30848] Block selected executable Git filters before patch application**  
   Prevents repository-selected clean/smudge filters from running during patch operations. Essential defense against malicious repos.  
   https://github.com/openai/codex/pull/30848

4. **[#31072] Bind patch application to guarded Git configuration**  
   Ensures validated configuration stays bound through the child process that performs mutations. Closes a gap where raw `GitRunner` could bypass checks.  
   https://github.com/openai/codex/pull/31072

5. **[#31058] Retry model capacity errors with jittered backoff**  
   Three retries (30s, 2m, 5m) for HTTP 503 capacity errors. Only for structured capacity failures, keeping fast path intact. Useful for high-demand models.  
   https://github.com/openai/codex/pull/31058

6. **[#30866] Reconcile loaded thread history on resume**  
   Fixes a class of bugs where resumed threads had stale state. Now serializes running-thread resume with rollback and history injection.  
   https://github.com/openai/codex/pull/30866

7. **[#30395] Expose rate-limit reset credit details (app-server)**  
   V2 API returns available credits, expiry times, and consumption endpoints. Enables proper redemption UI without private backend access.  
   https://github.com/openai/codex/pull/30395

8. **[#30488] Show reset details in redemption picker (CLI)**  
   Companion to #30395. Users can now see individual credits sorted by expiry before redeeming. Addresses long-standing UX gap.  
   https://github.com/openai/codex/pull/30488

9. **[#29181] Make image artifact directory configurable**  
   Adds `image_generation_artifacts_dir` to config.toml. Defaults to `$CODEX_HOME/generated_images`. Hosted deployments get flexibility.  
   https://github.com/openai/codex/pull/29181

10. **[#29082] Add connector skills feature toggle**  
    Enables A/B testing of connector-provided skills without removing connector apps, MCP tools, or hooks. Clean separation for experimentation.  
    https://github.com/openai/codex/pull/29082

## Feature Request Trends

- **Configurable auto-resolve timeout** (#28969, 77 👍): Users want to disable or extend the CLI's 60-second automatic question resolution.
- **Multi-line status line in TUI** (#21653, 30 👍): Terminal users need line-wrapping for long status lines.
- **Thread naming at session start** (#14482, 9 👍): `codex --name "task name"` for better session tracking.
- **Chinese (zh-CN) UI localization** (#31084): First internationalization request visible in the issue tracker.
- **Event-driven one-shot file watchers** (#31089): Proposed API for suspend-on-filesystem-events workflows.
- **Surfacing tool/skill catalog in JSON event stream** (#31088): Developers want programmatic access to what tools the model can use in a session.

## Developer Pain Points

- **Authentication/SSO friction** (#25670, #25737, #30892, #29595): Multiple reports of broken or inconsistent login flows, especially on CLI where hardware security keys are ignored in favor of SMS OTP. The "foreign phone number" variant (#30892) suggests session/account mismatch bugs.
- **Windows-specific crashes and sandbox issues** (#16374, #17083, #31035, #20942): System freezes, SysmonDrv BSODs, memory allocation failures, and forced read-only sandbox mode are recurring. Windows continues to be the most fragile platform.
- **Session state inconsistency** (#8648, #15508, #31095): Replies to wrong messages, MCP tools disappearing mid-session, and sessions hidden from sidebar after update. State management remains unreliable for long-running workflows.
- **Forced re-installation of uninstalled components** (#29450, #31035): Skills/plugins re-appearing after removal, and Sysmon driver reinstallation. Suggests aggressive update/repair logic lacking user intent checks.
- **macOS security false positives** (#23195, #30824): Gatekeeper flagging Codex as malware mid-session, combined with FSEvents crash loops, degrades trust in the desktop app on macOS.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest – July 4, 2026

## Today's Highlights
A nightly release (v0.51.0) is available, but the focus this week remains on resolving critical agent reliability issues. Several high-priority bugs—including subagent false‑success reports, generalist agent hangs, and shell command deadlocks—are actively being investigated. Meanwhile, the community’s demand for sandboxed execution and AST‑aware tooling continues to grow, as reflected in ongoing feature discussions.

## Releases
- **v0.51.0-nightly.20260704.gf7af4e518** – Latest nightly build.  
  Full changelog: [compare/v0.51.0-nightly.20260703...20260704](https://github.com/google-gemini/gemini-cli/compare/v0.51.0-nightly.20260703.gf7af4e518...v0.51.0-nightly.20260704.gf7af4e518)

## Hot Issues
1. **#22323 – Subagent recovery after MAX_TURNS reported as success** (P1, bug)  
   The `codebase_investigator` subagent marks itself as “GOAL” even when it hit the maximum turn limit without doing any analysis. 9 comments, 2 👍. A critical logic failure that undermines trust in agent status reporting.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/22323)

2. **#19873 – Leverage model’s bash affinity via Zero‑Dependency OS Sandboxing** (P2, enhancement)  
   Proposes using the model’s native bash skills securely via sandboxed execution and post‑execution intent routing. 8 comments, 1 👍. A long‑running feature request that could unlock powerful, safe automation.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/19873)

3. **#24353 – Robust component‑level evaluations** (P1, epic)  
   Builds on the behavioral eval framework (76 tests so far) and aims to scale it across models. 7 comments. Important for ensuring quality as the agent ecosystem grows.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/24353)

4. **#22745 – Assess AST‑aware file reads, search, and mapping** (P2, feature)  
   Investigates whether AST‑aware tools can reduce turn count and token waste. 7 comments, 1 👍. Developers see potential for noticeable speed improvements.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/22745)

5. **#21409 – Generalist agent hangs forever** (P1, bug)  
   The agent hangs for up to an hour on simple tasks (e.g., folder creation). 7 comments, 8 👍 (highest 👍 count). A top user frustration that forces users to disable sub‑agents.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/21409)

6. **#21968 – Gemini doesn’t use skills and sub‑agents often enough** (P2, bug)  
   Custom skills and sub‑agents are rarely invoked unless explicitly commanded. 6 comments. Impacts adoption of the agent extensibility model.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/21968)

7. **#26522 – Auto Memory retrying low‑signal sessions indefinitely** (P2, bug)  
   Sessions that are deemed low‑signal remain unprocessed and keep being surfaced, causing infinite retries. 5 comments. Memory system reliability is at stake.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/26522)

8. **#25166 – Shell command stuck with “Waiting input” after completion** (P1, bug)  
   Simple shell commands finish but the UI shows “Awaiting user input” forever. 4 comments, 3 👍. A common workflow blocker.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/25166)

9. **#21983 – Browser subagent fails on Wayland** (P1, bug)  
   The browser subagent terminates with “GOAL” but fails silently on Wayland. 4 comments, 1 👍. Platform compatibility issue affecting Linux users.  
   [Link](https://github.com/google-gemini/gemini-cli/issues/21983)

10. **#22672 – Agent should stop/discourage destructive behavior** (P2, feature)  
    The model occasionally uses `git reset`, `--force`, or other dangerous commands when safer alternatives exist. 3 comments, 1 👍. Community wants proactive safety guards.  
    [Link](https://github.com/google-gemini/gemini-cli/issues/22672)

## Key PR Progress
1. **#27839 – fix(core): make read_background_output delay abort‑aware** (closed, size/s)  
   Pressing ESC now properly cancels the tool; the spinner stops and new prompts are not queued. Fixes UI deadlock.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27839)

2. **#27971 – fix(core): strip thoughts from scrubbed history turns** (closed, size/m)  
   Resolves “thought leakage” where internal model reasoning corrupts history, causing infinite monologue loops. A major correctness fix.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27971)

3. **#28055 – fix(core): preserve dollar sequences in prompt template substitutions** (closed, size/m)  
   Stops `$` sequences (e.g., `$$`, `$'`) from being corrupted inside skill or agent descriptions. Prevents silent payload breakage.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28055)

4. **#28049 – fix(core): drop leading space from PascalCase markdown table headers** (closed, size/s)  
   Corrects `camelToSpace` to avoid extraneous spaces in table headers. Cleaner output.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28049)

5. **#28044 – fix(core): strip only trailing .json from checkpoint names** (closed, size/m)  
   Prevents accidental stripping of `.json` in file basenames inside checkpoint names. Data integrity fix.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28044)

6. **#28033 – fix(mcp): use longest‑prefix matching in parseMcpToolName** (closed, size/m)  
   Fixes tool routing for MCP server names containing underscores (issue #27981).  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28033)

7. **#28164 – fix(core): limit recursive reasoning turns per single user request** (open, size/m)  
   Implements a hard cap of 15 recursive reasoning turns to protect local resources and API quotas from infinite loops.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28164)

8. **#28163 – feat(caretaker): add triage worker core foundational modules** (open, size/l)  
   Introduces the base modules for the Caretaker Agent Triage Worker (CloudRun). First part of a two‑PR series for automated issue triage.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28163)

9. **#28162 – fix(core): buffer chat compression telemetry** (open, size/m)  
   Wraps OTEL log emission and metrics in the telemetry buffer to avoid slowdowns during chat compression. Fixes #23445.  
   [Link](https://github.com/google-gemini/gemini-cli/pull/28162)

10. **#28144 – fix(cli): detect available editors lazily to avoid slow startup** (open, size/m)  
    Delays editor detection (previously 20‑30 synchronous `execSync` calls) until first use, significantly improving startup time on Windows.  
    [Link](https://github.com/google-gemini/gemini-cli/pull/28144)

## Feature Request Trends
- **AST‑aware tooling** – Multiple issues (e.g., #22745, #22746, #15300) push for using AST parsers to improve codebase investigation, method reading, and search. The community expects fewer turns and lower token costs.
- **Sandboxed execution** – Issue #19873 (Zero‑Dependency OS Sandboxing) continues to attract attention. Developers want safe, native bash usage without compromising security.
- **Agent self‑awareness** – Requests for the CLI to understand its own flags, hotkeys, and capabilities (e.g., #21432) so it can serve as its own expert guide.
- **Memory system improvements** – Several issues (#26522, #26523, #26525) focus on making Auto Memory smarter: avoid retrying low‑signal sessions, quarantine invalid patches, and redact secrets deterministically.
- **Subagent trajectory visibility** – Users want to share subagent logs via `/chat share` (#22598) and include subagent context in bug reports (#21763).

## Developer Pain Points
- **Agent hangs and false success** – The generalist agent hanging (#21409) and subagent falsely reporting GOAL after turn‑limit (#22323) are the most upvoted bugs. Both erode trust in automated workflows.
- **Shell command deadlocks** – Commands that finish but leave the UI stuck (“Waiting input”) (#25166) are a frequent disruption.
- **Unwanted sub‑agent invocation** – After v0.33.0, sub‑agents sometimes run without permission (#22093), surprising users who had disabled them.
- **Memory retry loops** – Auto Memory repeatedly re‑evaluating low‑signal sessions (#26522) wastes resources and clutters logs.
- **Platform friction** – The browser agent failing on Wayland (#21983) and slow startup due to synchronous editor checks (#28144) hurt Linux and Windows experiences respectively.
- **Thought leakage** – Internal model reasoning leaking into history (#27971) caused confusion and infinite loops, a subtle but severe developer frustration now being addressed.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest – 2026-07-04

## Today’s Highlights
No new releases landed in the last 24 hours, but the issue tracker saw a burst of 14 fresh bug reports over the weekend – many related to model availability, headless agent dispatch, and session management. The community remains vocal about the alt‑screen rendering change and the lack of custom theme support (Issue #1504 has 20 👍). On the pull‑request side, only one (low‑activity) PR is open, indicating a slower development cycle.

## Releases
No new versions were published in the last 24 hours.

## Hot Issues (10 noteworthy)

1. **[#3997 – Copilot Web: Model "gpt-5.3-codex" is not available](https://github.com/github/copilot-cli/issues/3997)**  
   *9 comments, 0 👍*  
   A hard blocker for users relying on the web agent. The error appears during session creation and prevents any code generation. No workaround has been posted.

2. **[#1799 – How to turn off alt‑screen views?](https://github.com/github/copilot-cli/issues/1799)**  
   *11 comments, 7 👍*  
   The recently introduced alt‑screen mode is causing friction. Many users want an opt‑out or a toggle to revert to the original terminal behaviour. High community engagement.

3. **[#1504 – Add custom theme support](https://github.com/github/copilot-cli/issues/1504)**  
   *7 comments, 20 👍*  
   The most upvoted feature request. Users ask for JSON‑based theme files that could be shared. A long‑standing ask that continues to draw attention.

4. **[#3241 – Open sourcing the Copilot CLI](https://github.com/github/copilot-cli/issues/3241)**  
   *2 comments, 12 👍*  
   Advocates for full open‑sourcing to enable self‑hosted agent pipelines. Relevant for enterprise users who need on‑premises deployments.

5. **[#4019 – Built‑in web_fetch does not work with HTTP proxies](https://github.com/github/copilot-cli/issues/4019)**  
   *2 comments, 0 👍*  
   Corporate users on WSL cannot use `/research` or any URL retrieval because the CLI ignores system proxy settings. A frustrating gap for enterprise adoption.

6. **[#3533 – CLI 1.0.54 keyboard input not working on macOS](https://github.com/github/copilot-cli/issues/3533)**  
   *1 comment, 0 👍*  
   Text input becomes unresponsive due to a background authentication prompt for GitHub credentials. Affects a wide range of macOS users.

7. **[#4026 – Copilot CLI crashes repeatedly on Windows](https://github.com/github/copilot-cli/issues/4026)**  
   *0 comments, 0 👍*  
   A chronic crash affecting at least four versions, persists since May 2026. No reproducible single action, making it hard to diagnose. High priority for Windows users.

8. **[#4027 – Tool 'str_replace' does not exist](https://github.com/github/copilot-cli/issues/4027)**  
   *0 comments, 0 👍*  
   Occurs frequently when editing Java files. Copilot attempts to use a non‑existent tool, then falls back to a diff tool. Indicates a tool‑registration mismatch.

9. **[#4022 – Issues and Pull Requests tabs display as inactive/greyed out](https://github.com/github/copilot-cli/issues/4022)**  
   *0 comments, 0 👍*  
   In experimental mode the `Issues` and `Pull Requests` tabs are always greyed out regardless of repository status. Reduces the usefulness of the experimental UI.

10. **[#4025 – Session recall returns another project's history](https://github.com/github/copilot-cli/issues/4025)**  
    *0 comments, 0 👍*  
    All local sessions share a single state store, causing cross‑project contamination. When asked “what did we work on?” the CLI may show history from a different project.

## Key PR Progress (only 1 open PR)

**[#3771 – Initial project setup](https://github.com/github/copilot-cli/pull/3771)**  
*0 comments, 0 👍*  
A skeleton PR with no description, opened by a first‑time contributor. No updates or reviews. This suggests the CLI repository is not currently accepting significant feature contributions.

## Feature Request Trends
The most‑requested feature directions emerging from the issue tracker are:

- **Customization & Themes** – Users want to define and share custom themes (JSON‑based) and control terminal rendering details such as scroll speed (#1504, #4018).
- **Open‑Source / Self‑Hosting** – Demand for full open‑sourcing to allow on‑premises agent pipelines (#3241).
- **Non‑Interactive Mode** – The `/init` command should work in headless/batch scripts without hanging (#4011).
- **Toggle for Alt‑Screen** – A straight revert or configuration option to disable the new alt‑screen rendering (#1799).

## Developer Pain Points
Recurring frustrations highlighted in the last 24 hours:

- **Missing model errors** (#3997) break core agent functionality with no clear recovery.
- **Corporate proxy and authentication issues** (#4019, #3533) lock out enterprise and macOS users.
- **Windows stability** (#4026) remains unresolved after multiple versions.
- **Plugin and tool registration bugs** (#2709, #4027, #4021) lead to silent failures or contradictory error messages.
- **Session and state management problems** (#4020, #4025) cause misleading session ownership and cross‑project history leaks.
- **Experimental UI incompleteness** (#4022) lessens confidence in new features.

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

## Kimi Code CLI Community Digest — 2026-07-04

### Today's Highlights

A single bug fix was the only activity on the project today. The community flagged that the `thinking.enabled=false` configuration was being ignored for third‑party OpenAI‑compatible providers (e.g., DeepSeek via Sensenova), causing unwanted reasoning output. The issue was reported and closed quickly, indicating an active maintenance response.

### Releases

No new releases in the last 24 hours.

### Hot Issues

**1 issue was updated today** (only 1 active item).  
[#2484] **`thinking.enabled=false` not respected for third‑party OpenAI‑compatible vendors**  
*Author: lin200083 | Status: CLOSED*  
**Why it matters:** This bug directly affects users who deploy Kimi CLI with alternative backends like DeepSeek. The inability to disable thinking mode leads to extra latency and unnecessary token consumption, breaking expected behavior for CLI automation. The community reaction was low (0 reactions) but the fix was applied within the same day, showing responsiveness.  
🔗 [Issue #2484](https://github.com/MoonshotAI/kimi-cli/issues/2484)

### Key PR Progress

No pull requests were updated in the last 24 hours.

### Feature Request Trends

Based on today’s active issue, the most prominent community desire is **granular per‑provider / per‑model control over model parameters** — especially toggling thinking mode for third‑party endpoints. Users want the same level of configuration flexibility they have for first‑party models (e.g., Kimi) extended to external providers via `config.toml`. This suggests a broader demand for a **unified parameter override system** that works across all provider types.

### Developer Pain Points

- **Third‑party integration inconsistencies:** The `thinking.enabled=false` flag not applying to OpenAI‑compatible endpoints is a recurring frustration. Developers expect configuration to behave uniformly regardless of the underlying API.
- **Lack of transparency in model defaults:** When a model (e.g., DeepSeek) defaults to thinking mode, users have no easy way to discover or override it without deep‑diving into config files. The community would benefit from explicit error messages or documentation when a setting is ignored.

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

Here is the **OpenCode Community Digest** for **2026-07-04**.

---

# OpenCode Community Digest — July 4, 2026

## 1. Today’s Highlights

A major clipboard copy bug (#4283) is the top community pain point, with over 100 reactions and 24 hours of sustained discussion. Meanwhile, several users report a server-side outage affecting OpenCode's "Go" and "Zen" model APIs, triggering a wave of `500 Internal Server Error` reports. On the positive side, multiple contributors have landed PRs today targeting the clipboard issue and the missing `reasoning` field in OpenAI-compatible streams, signaling active maintenance velocity.

## 2. Releases

No new releases in the last 24 hours.

## 3. Hot Issues (Top 10)

1.  **[#4283 – Copy To Clipboard is not working](https://github.com/anomalyco/opencode/issues/4283)**
    *   **Why it matters:** The most active issue today (104 comments, 100 👍). Users on Linux (Ubuntu 24.04, Wayland) report that copy operations silently fail or paste stale content. This is a core UX blocker for daily TUI use.
    *   **Community reaction:** High urgency; users are sharing terminal emulator and OS details to help isolate the issue.

2.  **[#22225 – [FEATURE]: Add skill usage tracking to CLI](https://github.com/anomalyco/opencode/issues/22225)**
    *   **Why it matters:** Users want telemetry on which skills are used most, stored locally in a JSON file. This is a step toward self-optimization and usage analytics.
    *   **Community reaction:** Light discussion (11 comments) but zero upvotes; likely a niche power-user request.

3.  **[#1168 – Feature Request: Make Links Clickable (Ctrl+Left Click to Open)](https://github.com/anomalyco/opencode/issues/1168)**
    *   **Why it matters:** A long-standing ergonomic request (105 👍) that has been open for nearly a year. Today, a new sub-issue (#35286) highlights that mouse capture mode breaks this feature entirely.
    *   **Community reaction:** Strong positive sentiment; users want terminal emulator gestures to pass through when mouse capture is enabled.

4.  **[#35278 – [needs:compliance] No provider available](https://github.com/anomalyco/opencode/issues/35278)**
    *   **Why it matters:** Users are seeing "No provider available" on app launch, suggesting a configuration or service discovery regression.
    *   **Community reaction:** 6 reactions in under 24 hours; multiple duplicates (e.g., #35276, #35279) reporting the same symptoms for Go/Zen models.

5.  **[#35276 – OpenCode Zen/Go API chat completions returning 500 Internal Server Error](https://github.com/anomalyco/opencode/issues/35276)**
    *   **Why it matters:** Confirmed server-side outage affecting all API keys and models on the Zen/Go endpoints. Likely the root cause of the "No provider" errors in #35278.
    *   **Community reaction:** 4 comments in rapid succession; users are anxious for a fix or status update.

6.  **[#35286 – [needs:compliance] Cmd+click to open URLs not working when mouse capture is enabled](https://github.com/anomalyco/opencode/issues/35286)**
    *   **Why it matters:** Regression report: when `mouse: true` is set, the TUI intercepts all mouse events, preventing Cmd+click from reaching the terminal emulator (e.g., Ghostty, iTerm2).
    *   **Community reaction:** New report, but cross-references the older #1168 suggestion, indicating this is a known pain point.

7.  **[#35290 – Subscription billing not canceled despite disabling auto-renewal in Alipay](https://github.com/anomalyco/opencode/issues/35290)**
    *   **Why it matters:** A billing/cancellation UX bug. User reports that turning off auto-renewal in Alipay did not stop the subscription, leading to confusion and a potential financial dispute.
    *   **Community reaction:** Sensitive topic; the `[needs:compliance]` label suggests the team is treating this seriously.

8.  **[#35294 – Desktop: "Local Server could not be reached" on startup due to blocking models.dev fetch](https://github.com/anomalyco/opencode/issues/35294)**
    *   **Why it matters:** A regression in v1.17.12+ causes a blocking network fetch on startup, hanging the sidecar and showing a fatal error when the network is slow.
    *   **Community reaction:** User reports that downgrading to v1.17.11 fixes the issue, which helps narrow the regression.

9.  **[#35072 – fix: context limit display uses stale model after model switch](https://github.com/anomalyco/opencode/issues/35072)**
    *   **Why it matters:** The sidebar's "Context %" continues to calculate based on the previous model’s context window after switching, showing misleading usage percentages.
    *   **Community reaction:** Small number of comments but a clear, reproducible bug report.

10. **[#35295 – mouse wheel falls back to arrow keys, triggering prompt history instead of scrolling chat](https://github.com/anomalyco/opencode/issues/35295)**
    *   **Why it matters:** When mouse is disabled, the wheel event is misinterpreted as arrow key presses, causing unintended prompt-history navigation instead of viewport scrolling.
    *   **Community reaction:** New report; likely to gain traction as more users toggle mouse off for accessibility or preference reasons.

## 4. Key PR Progress (Top 10)

1.  **[#35289 – fix(tui): flush OSC 52 clipboard write, propagate errors on fallback](https://github.com/anomalyco/opencode/pull/35289)**
    *   **What it does:** Fixes the clipboard bug #4283. Flushes the OSC 52 escape sequence and propagates errors on Wayland when clipboard utilities are missing.
    *   **Why it matters:** Directly addresses the most upvoted issue. Close to landing.

2.  **[#35287 – fix(tui): use current model for sidebar context limit after model switch](https://github.com/anomalyco/opencode/pull/35287)**
    *   **What it does:** Resolves #35072 by computing the context percentage from the currently selected model, not the last assistant message's model.
    *   **Why it matters:** Fixes a misleading UI metric that impacts user trust in context management.

3.  **[#35284 – fix(llm): accept `reasoning` field in OpenAI-compatible streams](https://github.com/anomalyco/opencode/pull/35284)**
    *   **What it does:** Adds `reasoning` to the OpenAI delta schema alongside `reasoning_content`. Closes #35283.
    *   **Why it matters:** Unlocks chain-of-thought display for providers that use the standard `reasoning` field name.

4.  **[#35292 – tui: preserve spinner registration](https://github.com/anomalyco/opencode/pull/35292)**
    *   **What it does:** Uses explicit registration to prevent tree-shaking from removing the spinner component during builds.
    *   **Why it matters:** Fixes a build reliability issue where spinners could disappear in production bundles.

5.  **[#35293 – feat: add command_session tool for interactive long-running commands](https://github.com/anomalyco/opencode/pull/35293)**
    *   **What it does:** Introduces a new `command_session` tool that runs PTY-backed interactive commands in the background, supporting polling, input injection, and termination.
    *   **Why it matters:** A major new capability for developers running interactive shells or daemon processes inside OpenCode.

6.  **[#35281 – fix(core): enforce step settlement ordering](https://github.com/anomalyco/opencode/pull/35281)**
    *   **What it does:** Ensures tool settlement events always fire before the step terminal event, even across error and interruption paths. Closes #35020.
    *   **Why it matters:** V2 event-audit cleanup; prevents race conditions in step lifecycle management.

7.  **[#35280 – feat(core): add execution lifecycle and structured errors](https://github.com/anomalyco/opencode/pull/35280)**
    *   **What it does:** Replaces ambiguous execution events with explicit lifecycle (`started`, `succeeded`, `failed`, `interrupted`) and bounded retries.
    *   **Why it matters:** Improves observability and error handling for V2 session execution.

8.  **[#35272 – refactor(schema): simplify session fragment state](https://github.com/anomalyco/opencode/pull/35272)**
    *   **What it does:** Simplifies fragment and provider-continuation state while keeping provider-native correlation inside the LLM package.
    *   **Why it matters:** Reduces schema complexity; helps unify the V2 state model.

9.  **[#35269 – fix(app): hydrate timeline message parents](https://github.com/anomalyco/opencode/pull/35269)**
    *   **What it does:** Keeps fast initial timeline rendering but adds hydration of missing assistant parent messages before publishing the page.
    *   **Why it matters:** Fixes a visual glitch in the Desktop app timeline where earlier messages could appear empty.

10. **[#35192 – feat(codemode): add OpenAPI tool adapter](https://github.com/anomalyco/opencode/pull/35192)**
    *   **What it does:** Exports `OpenAPI.fromSpec` to create a subtree of codemode tools from an OpenAPI 3.x spec. Auth is handled server-side, never exposed to the model.
    *   **Why it matters:** Enables OpenCode to act as a generic API client, opening up integration with any REST service.

## 5. Feature Request Trends

*   **Enhanced Link Handling:** Multiple issues (e.g., #1168, #35286, #35288) ask for better URL management, including clickable links in the TUI and proper Cmd+click support when mouse capture is active.
*   **Persistent User Memory:** Requests for long-term memory (#35291) that recalls user preferences, emotions, and project context across sessions and reboots.
*   **Session & Data Export:** Users want prompt-only export with timestamps (#35128) and better tool call display policy (show/collapse/hide per tool type) (#35193).
*   **Configurable Currency Display:** A call for configurable currency symbols (CNY, EUR, etc.) in cost-tracking widgets (#35274).
*   **Scheduled/Cron Tasks:** A request for integrating cron-like job automation directly into OpenCode (#35277).
*   **Skill Usage Analytics:** Built-in usage tracking for skills (#22225) to help users identify which tools they rely on most.

## 6. Developer Pain Points

*   **Clipboard Inconsistency:** The #1 pain point today. Copying text fails silently on Wayland/Linux (#4283, #29834, #31253). Users are frustrated by the "success" toast with no actual data in the clipboard.
*   **Server-Side Instability:** A cluster of issues (#35276, #35279, #35273, #35278) report internal server errors, missing providers, or daily service interruptions for DeepSeek V4 Flash and Go models. Users are seeking official explanations.
*   **Mouse Interaction Conflicts:** When mouse capture is enabled, Cmd+click does not open URLs (#35286, #35288). When mouse is disabled, the wheel scrolls the prompt history instead of the chat (#35295). Both states break core workflows.
*   **Desktop Startup Degradation:** A blocking network fetch in v1.17.12+ hangs the Desktop app on slow networks (#35294), forcing users to roll back — a sharp regression in reliability.
*   **Project State Confusion:** Users report that copied/moved project folders inherit the original workspace state (#35297) and that renamed directories do not resolve properly in the Desktop app (#33801).
*   **Subscription Management:** Cancellation via third-party payment (Alipay) does not propagate to OpenCode (#35290), causing confusion about billing status.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-07-04

## Today's Highlights
Reliability remains the dominant theme: the long‑running **Codex connection issue** (#4945) continues to affect many users, and a **critical 60‑minute WebSocket timeout** was reported for the same provider (#6268). On the fix side, two high‑impact PRs landed—one strips hallucinated extra keys from the edit tool (#6283), another improves Cloudflare Workers AI resolution (#6292). Meanwhile, users are calling for better support for **Kimi K2.7** and **Claude Sonnet 5** in the GitHub Copilot provider.

## Releases
No new releases in the last 24 hours.

## Hot Issues (10 items)

1. **#4945 – openai-codex Connection Reliability Issues**  
   *73 comments · 30 👍*  
   Users report the TUI getting stuck on “Working…” with no error, recoverable only by pressing Escape. The problem has persisted for weeks and is a top concern for those relying on Codex.  
   [GitHub link](https://github.com/earendil-works/pi/issues/4945)

2. **#6215 – pi update fails on 0.80.3 due to missing @smithy/node-http-handler@^4.9.1**  
   *22 comments*  
   A dependency resolution failure blocked upgrades for some users. Now closed, but illustrates the fragility of pnpm‑based updates.  
   [GitHub link](https://github.com/earendil-works/pi/issues/6215)

3. **#6187 – Pi login hangs in WSL after browser-based GitHub Copilot device authorization**  
   *15 comments*  
   After device registration the CLI never detects completion, leaving users stuck. A WSL‑specific annoyance that impacts a significant subset.  
   [GitHub link](https://github.com/earendil-works/pi/issues/6187)

4. **#6278 – New Claude models work poorly with the current Pi’s edit tool**  
   *12 comments*  
   Sonnet 5, Fable 5, Opus 4.8 emit extra keys like `newText_x` or `closeenough` inside edit requests, causing validation failures in ~20% of edits.  
   [GitHub link](https://github.com/earendil-works/pi/issues/6278)

5. **#5501 – tolerate extra keys on edit tool edits[] items**  
   *10 comments*  
   A prior proposal to drop `additionalProperties: false` from the edit schema. Merged as a longer‑term fix, but #6278 shows the problem is still acute.  
   [GitHub link](https://github.com/earendil-works/pi/issues/5501)

6. **#6239 – HTTP 524 (Cloudflare timeout) should be treated as retryable**  
   *3 comments*  
   When a proxy/gateway times out, Pi currently aborts the session instead of retrying. Simple fix, high impact for users behind Cloudflare.  
   [GitHub link](https://github.com/earendil-works/pi/issues/6239)

7. **#6268 – Codex websocket terminates after 60 minutes, does not retry**  
   *3 comments*  
   Long‑running tasks hit a hard WebSocket limit with no automatic reconnect, forcing manual intervention.  
   [GitHub link](https://github.com/earendil-works/pi/issues/6268)

8. **#6301 – Hide/disable individual slash commands without disabling the whole extension**  
   *2 comments*  
   A per‑extension granularity request for slash command visibility that resonated with the community.  
   [GitHub link](https://github.com/earendil-works/pi/issues/6301)

9. **#6295 – openai-completions hides reasoning-only replies when thinking is off**  
   *2 comments*  
   Reasoning‑only responses (no `content`) get stored as thinking blocks and silently disappear—a subtle data‑loss bug.  
   [GitHub link](https://github.com/earendil-works/pi/issues/6295)

10. **#5084 – Allow/disallow built-in tools in settings.json**  
    *3 comments · 8 👍*  
    Users want persistent tool toggles beyond the `--tools` CLI flag. Gaining traction as a quality‑of‑life improvement.  
    [GitHub link](https://github.com/earendil-works/pi/issues/5084)

## Key PR Progress (7 items)

1. **#6294 – Improve pi config add-ons UX**  
   *Closed*  
   Reworks the `pi config` command around an Add‑ons mental model, adds package‑level toggles, detail pane, and model‑fit guidance for local/weak models.  
   [GitHub link](https://github.com/earendil-works/pi/pull/6294)

2. **#6292 – fix(ai): resolve Cloudflare account id from ambient env for key‑only credentials**  
   *Closed*  
   Fixes the persistent `404` error on Cloudflare Workers AI (0.80.x) by sourcing the account ID from the environment when only keys are provided.  
   [GitHub link](https://github.com/earendil-works/pi/pull/6292)

3. **#6290 – fix(ai): use “(no tool output)” placeholder for empty tool results without images**  
   *Closed*  
   Prevents the model from hallucinating image attachments when a tool returns no output (e.g., empty `grep`). Affects OpenAI providers.  
   [GitHub link](https://github.com/earendil-works/pi/pull/6290)

4. **#6285 – fix(ai): stop salvaging malformed tool‑call argument JSON**  
   *Open*  
   A strict parsing change: truncated/malformed JSON is now preserved as `malformedArguments` instead of silently broken. Invasive but improves debuggability.  
   [GitHub link](https://github.com/earendil-works/pi/pull/6285)

5. **#6283 – fix(coding-agent): strip hallucinated extra keys from edit tool edits[]**  
   *Closed*  
   Direct fix for #6278: silently removes known extra keys before validation. A pragmatic band‑aid while the schema relaxation (see #5501) is evaluated.  
   [GitHub link](https://github.com/earendil-works/pi/pull/6283)

6. **#6279 – fix(coding-agent): add pnpm self‑update prune hint**  
   *Closed*  
   Adds a recovery hint (`pnpm store prune pi`) when `pi update` fails due to stale metadata. Helpful for users hitting the #6215 class of errors.  
   [GitHub link](https://github.com/earendil-works/pi/pull/6279)

7. **#6220 – ignore**  
   *Closed*  
   Non‑functional test PR – excluded from digest.  

*Note: only 7 PRs were updated in the last 24h; all meaningful ones are listed above.*

## Feature Request Trends
Several clear directions emerge from the latest issues:

- **Expanded model support** – Requests to add **Kimi K2.7** and **Claude Sonnet 5** to the GitHub Copilot model list (#6256, #6257) show users want to stay current with provider offerings.
- **Edit tool resilience** – The community strongly desires the edit tool schema to be **lenient toward hallucinated keys** (#5501, #6278), with a PR already merged (#6283).
- **Fine‑grained tool and command configuration** – Persistent settings to **allow/disallow built‑in tools** (#5084) and **hide individual slash commands** (#6301) appear repeatedly.
- **Session & UX enhancements** – Automatic **AI‑generated session titles** (#6209), **adding session names to resume hints** (#6296), and **showing active built‑in tools in the footer** (#6277) reflect a desire for a more polished interactive experience.
- **Sandboxing and security hardening** – Three detailed notes (#6297, #6298, #6299) discuss making the gondolin and subagent examples safer for multi‑tenant use, including VM egress control and filesystem sandboxing.

## Developer Pain Points
Recurring frustrations from the past 24 hours:

- **Codex instability** – Stuck “Working…” screens (#4945) and hard 60‑minute WebSocket drops (#6268) erode trust for long‑running tasks.
- **WSL login breakage** – Pi cannot detect completed Copilot authorization in WSL (#6187), forcing manual workarounds.
- **Cloudflare proxy issues** – HTTP 524 timeouts are not retried (#6239), and Workers AI still returns 404 for some credential configurations (#6021, #6292 fix now available).
- **Model‑specific edit tool failures** – Newer Claude models routinely inject extra keys, causing ~20% edit failures (#6278). While fixed in PR #6283, the root cause (schema strictness) remains unaddressed.
- **Dependency version mismatches** – `@smithy/node-http-handler` resolution failures block updates (#6215, #6279). A recurring theme with pnpm’s strict versioning.
- **Windows TUI rendering** – Each keystroke redraws the input line incorrectly (#6300), making interactive use painful on Windows.
- **Data loss / invisible assistant responses** – Reasoning‑only replies disappear when thinking display is off (#6295), and empty tool outputs trigger hallucinations (#6290, now fixed).
- **No visible tool state** – Users cannot easily see which built‑in tools are active during a session (#6277), leading to confusion when `--no-builtin-tools` is used.

---

*Digest generated from `github.com/badlogic/pi-mono` data for 2026-07-04.*

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest – 2026-07-04

## Today’s Highlights

Two releases landed today: a **v0.19.6 stable** with mobile session jank fixes and a **nightly** tightening PR triage automation. Community activity remains high with 12 issues and 50 PRs updated in the last 24 hours. Key themes include **token and cost management** (e.g., context window miscalculation, expensive `/review` skill), **Windows shell compatibility**, and **sub-agent concurrency control**. A notable PR fixes persistent 401 errors after API key changes, while several PRs enhance the web shell and channel integrations.

## Releases

**v0.19.6-nightly.20260704** – Fixes the PR triage gate with stronger batch detection, problem existence checks, and red flag patterns (by @pomelo-nwu).  
**v0.19.6** – Fixes mobile web-shell session-switch jank (memoized timeline signature, replay-first dispatch) and a macOS seat issue.  
*No other versions released.*

## Hot Issues (10 notable)

1. **#6144** – [Qwen-Code calculates incorrect context window](https://github.com/QwenLM/qwen-code/issues/6144)  
  *P2 bug, core/token-management.* Users report that with a 64K model the actual context window is wrong. 7 comments, 1 👍. Community is actively debugging config interplay (cache type, temperature etc.).

2. **#6264** – [`/review` skill consumes large amounts of tokens](https://github.com/QwenLM/qwen-code/issues/6264)  
  *P2 bug, performance.* Attached screenshots show token usage spikes. Users love the skill but find it prohibitively expensive. 3 comments, no solution yet.

3. **#6298** – [Shell tool fails on Windows when command produces stdout](https://github.com/QwenLM/qwen-code/issues/6298)  
  *P2 bug, shell/Windows.* `cat` is used internally but unavailable on Windows `cmd.exe`. 2 comments. A clear platform gap.

4. **#6290** – [`QWEN_CODE_MAX_BACKGROUND_AGENTS` doesn’t limit Explorer sub‑agents](https://github.com/QwenLM/qwen-code/issues/6290)  
  *Bug, core/tools.* Setting the env var to 1 still spawns 2 Explorer sub-agents in parallel. 2 comments. Tied to concurrency enforcement.

5. **#6289** – [Files attached from prompt aren’t treated as read, preventing editing right away](https://github.com/QwenLM/qwen-code/issues/6289)  
  *Bug, interactive/file-operations.* Agent must re-read a file even after user attached it. 2 comments. A UX friction point.

6. **#6299** – [ci-bot keeps running after PR is closed and spams notifications](https://github.com/QwenLM/qwen-code/issues/6299)  
  *P2 bug, core/CI-CD.* User reports CI continues to review and send emails after PR #6240 was closed. Chinese discussion, 1 comment. Highlights annoyance of overly aggressive automation.

7. **#6197** – [VSCode IDE Companion Release Failed for 0.19.5](https://github.com/QwenLM/qwen-code/issues/6197)  
  *Closed.* Automated issue from GitHub Actions. Relevant for tracking release pipeline reliability.

8. **#5942** – [Anthropic provider prompt-cache misses inflate cost](https://github.com/QwenLM/qwen-code/issues/5942)  
  *Closed, P2 bug, performance/token-management.* Two separate cache issues cause higher costs than Claude Code. 4 comments. Important for users routing through Anthropic endpoints.

9. **#6230** – [Quickpick dropdowns lose focus during `/auth` process](https://github.com/QwenLM/qwen-code/issues/6230)  
  *P2 bug, UI/authentication.* Focus loss forces restart of configuration. 2 comments. Affects VS Code extension users.

10. **#5634** – [Autofix tier‑1 trusts an LLM-applied label that can be influenced by issue text](https://github.com/QwenLM/qwen-code/issues/5634)  
  *Closed, security.* The fast‑path label `ready-for-agent` is not a human signal, leading to possible bypass of human review. 4 comments. Addressed by PR #6276.

## Key PR Progress (10 important)

1. **#6300** – [fix(core): enforce agent concurrency cap on foreground sub‑agents](https://github.com/QwenLM/qwen-code/pull/6300)  
  Addresses #6290. Adds concurrency limit to foreground sub-agents (Explorer, etc.) which previously bypassed `QWEN_CODE_MAX_BACKGROUND_AGENTS`. By @kagura-agent.

2. **#6295** – [fix(core): treat @‑attached files as read for prior‑read enforcement](https://github.com/QwenLM/qwen-code/pull/6295)  
  Fixes #6289. Files attached via `@path` are now recorded in the session file‑read cache, so editing works immediately. By @Nas01010101.

3. **#6284** – [fix(auth): prevent persistent 401 after API key change](https://github.com/QwenLM/qwen-code/pull/6284)  
  Fixes three failure modes including empty-string env vars and stale cached credentials. By @yiliang114.

4. **#6276** – [fix(ci): require maintainer‑applied `autofix/approved` label for tier‑1 fast‑path](https://github.com/QwenLM/qwen-code/pull/6276)  
  Addresses #5634 security concern. Adds dual‑factor trust gate. By @yiliang114 (closed).

5. **#6273** – [feat(core): model fallback chain — auto‑switch to backup models on overload](https://github.com/QwenLM/qwen-code/pull/6273)  
  Configurable backup models when primary is unavailable. Preserves retry logic. By @yiliang114.

6. **#6301** – [fix(cli): error on unknown slash commands in non‑interactive mode](https://github.com/QwenLM/qwen-code/pull/6301)  
  Previously silently forwarded to model as prompt. Now produces clear error and exits code 1. By @yiliang114 (closed).

7. **#6297** – [feat(daemon): Add session export endpoint](https://github.com/QwenLM/qwen-code/pull/6297)  
  Read‑only daemon endpoint to export active session transcripts as HTML, Markdown, JSON, JSONL. Uses existing CLI formatters. By @doudouOUC.

8. **#6242** – [feat(web‑shell): add custom at mention panel](https://github.com/QwenLM/qwen-code/pull/6242)  
  Replaces inline @‑autocomplete with a multi‑level reference panel supporting categories, files, extensions, MCP resources. By @ytahdn.

9. **#6287** – [feat: add proactive channel loop tools](https://github.com/QwenLM/qwen-code/pull/6287)  
  Adds channel‑owned loop tooling for recurring reminders through MCP. By @qqqys.

10. **#6302** – [test(e2e): make fake OpenAI reachable from Docker sandbox](https://github.com/QwenLM/qwen-code/pull/6302)  
  Ensures SDK E2E tests work inside Docker/Podman by advertising endpoint via `host.docker.internal`. By @yiliang114.

## Feature Request Trends

- **Session and data export** – Multiple PRs and issues (e.g., #6297) advocate for persistent session export in various formats.
- **Multi‑workspace/file system boundary support** – PRs like #6278 address VSCode multi‑folder workspaces.
- **Channel integrations** – WeCom intelligent robot (#6224) and proactive channel loops (#6287) show demand for richer messaging platform support.
- **Token cost optimization** – Issues #5942 (Anthropic cache) and #6264 (`/review` cost) signal strong community desire for better cost control, especially with paid backends.
- **Automation trust and control** – Issues #6299 (CI over‑eager) and #5634 (autofix label abuse) point to a need for more nuanced CI/copilot automation.

## Developer Pain Points

1. **Windows compatibility gaps** – Shell tool fails due to Unix‑only commands (`cat`). Also reported in earlier issues. Windows users face friction.
2. **Token waste and cost surprises** – `/review` skill and incorrect context window calculations are top user frustrations; providers with cache pricing add complexity.
3. **Concurrency controls ignored** – `QWEN_CODE_MAX_BACKGROUND_AGENTS` not honoured for foreground sub‑agents, leading to unexpected parallel spawns.
4. **File attachment UX** – Files attached via `@` are not recognised by the session cache, forcing extra steps to edit.
5. **Over‑aggressive CI bot** – PRs closed still trigger reviews and notifications, wasting tokens and developer attention. Chinese comments highlight “building technical debt under bot demands”.
6. **Authentication flow issues** – Quickpick dropdowns losing focus during `/auth` on VS Code; persistent 401 after key changes (now being fixed by #6284).

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest – 2026-07-04

## Today's Highlights

The CodeWhale project is deep in v0.8.67 release‑candidate hardening, with the team and contributors landing a wave of bug fixes, reliability improvements, and UX polish. Several long‑standing issues around provider routing and plugin persistence are gaining traction via community PRs, while performance and security concerns remain top of mind. Notable contributions include a per‑sub‑agent provider routing PR (#3969) and a foundational dynamic MCP server infrastructure PR (#3869).

## Releases

None in the last 24 hours.

## Hot Issues

1. **#3275 – [OPEN] CodeWhale over‑extending scope**  
   User reports CodeWhale enters a self‑driven loop of proposing, answering, and executing without waiting for confirmation. Regression from #3061.  
   *17 comments, high signal.*  
   [Link](https://github.com/Hmbown/CodeWhale/issues/3275)

2. **#3793 – [OPEN] Guided localized constitution creator**  
   Proposes replacing the blank prompt editor with a step‑by‑step wizard that puts language first and separates security posture. *16 comments, deep UX discussion.*  
   [Link](https://github.com/Hmbown/CodeWhale/issues/3793)

3. **#4026 – [OPEN] Light theme selection highlight invisible**  
   In light theme, selected text in the terminal shell has no visible highlight. *2 comments, quick reproduction.*  
   [Link](https://github.com/Hmbown/CodeWhale/issues/4026)

4. **#3965 – [OPEN] Per‑sub‑agent provider assignment + LM Studio**  
   User request to pin different sub‑agents to different providers (e.g., local LM Studio for format tasks). *7 comments, enthusiastic community discussion.*  
   [Link](https://github.com/Hmbown/CodeWhale/issues/3965)

5. **#3884 – [CLOSED] Codex sub‑agents fail with “Responses API request failed”**  
   Release blocker that blocked orchestrated work; root cause was provider/runtime error. *4 comments, fast resolution.*  
   [Link](https://github.com/Hmbown/CodeWhale/issues/3884)

6. **#4014 – [OPEN] TUI lag & memory pressure with 30+ sub‑agents**  
   Observed typing latency, rendering stalls, and host memory pressure during high fan‑out sessions. *0 comments but critical performance bug.*  
   [Link](https://github.com/Hmbown/CodeWhale/issues/4014)

7. **#4027 – [OPEN] Add `always_load` field for MCP tools**  
   Suggests skipping defer‑loading for frequently used tools to avoid round‑trip latency. *1 comment, thoughtful proposal.*  
   [Link](https://github.com/Hmbown/CodeWhale/issues/4027)

8. **#3924 – [CLOSED] Display‑width helpers triplicated and divergent**  
   Three implementations of terminal display‑width logic existed; one measured control chars differently. *1 comment, clean refactor.*  
   [Link](https://github.com/Hmbown/CodeWhale/issues/3924)

9. **#3918 – [CLOSED] Plugin enable/disable not persisted**  
   `/plugin enable|disable` wrote to an in‑memory map that was never saved. *1 comment, fixed.*  
   [Link](https://github.com/Hmbown/CodeWhale/issues/3918)

10. **#4009 – [CLOSED] Agent sidebar doesn’t reflect cancellations**  
    Cancelling sub‑agents via tool did not update the sidebar UI. *0 comments, UX regression.*  
    [Link](https://github.com/Hmbown/CodeWhale/issues/4009)

## Key PR Progress

1. **#4028 – [OPEN] fix(tui): keep provider links readable in narrow layouts**  
   Renders provider URLs as inline code to avoid oversized OSC 8 autolinks. Adds regression test. *by roian6.*  
   [Link](https://github.com/Hmbown/CodeWhale/pull/4028)

2. **#3967 – [OPEN] perf(tui): avoid redundant composer input wrapping per frame**  
   Eliminates up to 5x unnecessary wrapping per frame, addressing TUI performance. *by reidliu41.*  
   [Link](https://github.com/Hmbown/CodeWhale/pull/3967)

3. **#4025 – [OPEN] ci: light‑classify inert scripts and stop allocating heavy runners**  
   Prevents full macOS/Windows CI runs for trivial script changes (e.g., docs). *by Hmbown.*  
   [Link](https://github.com/Hmbown/CodeWhale/pull/4025)

4. **#4023 – [CLOSED] fix(tui): harden v0.8.67 rc surfaces**  
   Bulk RC fix covering stream timeout, plugin paths, onboarding copy, provider routing, OAuth messaging, and subagent sidebar. *by Hmbown.*  
   [Link](https://github.com/Hmbown/CodeWhale/pull/4023)

5. **#3972 – [CLOSED] fix(tui): allow longer quiet reasoning waits**  
   Raises default streamed‑response idle timeout from 300s to 900s for reasoning models. *by Hmbown.*  
   [Link](https://github.com/Hmbown/CodeWhale/pull/3972)

6. **#3869 – [OPEN] feat: add dynamic MCP server infrastructure to McpPool**  
   Foundation for runtime‑started MCP servers via `start_mcp_server` tool. *by bistack.*  
   [Link](https://github.com/Hmbown/CodeWhale/pull/3869)

7. **#3866 – [OPEN] feat: LLM can start MCP servers from chat context**  
   Adds `start_mcp_server` tool supporting stdio and HTTP transports. *by bistack.*  
   [Link](https://github.com/Hmbown/CodeWhale/pull/3866)

8. **#3781 – [OPEN] Feat/opencode zen provider**  
   Implements a new provider for OpenCode Zen. *by snail-vs.*  
   [Link](https://github.com/Hmbown/CodeWhale/pull/3781)

9. **#3762 – [OPEN] feat(web): redesign homepage with trust strip, GitHub nav, mirror footer**  
   Website redesign adding trust indicators and localization. *by idling11.*  
   [Link](https://github.com/Hmbown/CodeWhale/pull/3762)

10. **#3780 – [OPEN] [codex] expose context compaction gates**  
    Adds engine‑level switches for replacement compaction and Flash seam manager. *by nightt5879.*  
    [Link](https://github.com/Hmbown/CodeWhale/pull/3780)

## Feature Request Trends

- **Per‑sub‑agent provider routing** – multiple issues/PRs request pinning sub‑agents to specific providers (e.g., local LM Studio, different API endpoints). (#3965, #3969)
- **Dynamic MCP server startup** – allowing LLMs to launch MCP servers at runtime from conversation context. (#3869, #3866, #4027)
- **Guided localized constitution creator** – replacing the blank prompt editor with a step‑by‑step wizard that respects language and security separation. (#3793)
- **Persistent & actionable update prompts** – making new‑version notifications more useful and not dismiss‑only. (#3961)
- **TUI performance for high fan‑out sessions** – reducing lag, memory pressure, and rendering stalls with many sub‑agents. (#4014, #3967)

## Developer Pain Points

- **Plugin state not persisted** – enable/disable resets on restart. (#3918)
- **CodeWhale over‑extending scope** – self‑questioning and deviating from user intent without confirmation. (#3275)
- **Codex sub‑agent failures** – blocking orchestrated work until resolved. (#3884)
- **Display‑width helpers triplicated** – divergent implementations cause bugs in terminal text measurement. (#3924)
- **Agent sidebar not reflecting cancellations** – UI stale after cancelling sub‑agents. (#4009)
- **First‑run setup confusion** – overlapping CLI/TUI setup surfaces, ambiguous provider routes, and onboarding copy mismatches. (#4003, #4005, #4002)
- **Light theme visibility issues** – selection highlight invisible, making terminal use harder. (#4026)

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*