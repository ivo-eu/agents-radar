# AI CLI Tools Community Digest 2026-06-25

> Generated: 2026-06-25 10:25 UTC | Tools covered: 9

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
**Analysis Date: 2026-06-25**

---

## 1. Ecosystem Overview

The AI CLI tools landscape on June 25, 2026, reveals a maturing yet volatile ecosystem where rapid feature iteration is colliding with production stability demands. The dominant theme across tools is the tension between aggressive safety/security classifiers and legitimate developer workflows—Claude Code alone saw 20+ false-positive reports in 24 hours, while Codex users face systemic quota accounting failures. Multi-agent orchestration reliability remains the single biggest technical challenge, affecting Claude Code, Gemini CLI, and Copilot CLI alike, with stuck agent pointers, zombie processes, and misleading success reports eroding trust. Windows support emerges as a critical second-class-citizen problem: OpenCode's v1.17.10 release effectively broke the platform, DeepSeek TUI fixed environment inheritance bugs, and Codex users report sandbox execution errors. Meanwhile, MCP (Model Context Protocol) integration is converging as the de facto extensibility standard, with all major tools either shipping MCP improvements or having them as top community requests.

---

## 2. Activity Comparison

| Tool | Hot Issues (Today) | Active PRs | Release Status | Community Signal |
|-----|-------------------|------------|----------------|-----------------|
| **Claude Code** | 10 (top issues) | 3 open PRs | **v2.1.191** (today) + v2.1.190 | Crisis mode: 20+ false positive reports in 24h |
| **OpenAI Codex** | 10 (hot issues) | 10 open PRs | **rust-v0.142.2** (today) + alphas | Systemic: rate-limit/accounting bugs dominate |
| **Gemini CLI** | 10 (top by conversation) | 10 open PRs | **v0.49.0-nightly** (today) | Steady: agent reliability + security focus |
| **Copilot CLI** | 10 (top impact) | 2 PRs updated | **v1.0.65** (yesterday) | Moderate: resumed session auth bug trending |
| **Kimi Code** | 4 open/closed | 2 PRs closed | No new release today | Quiet: low activity, long-standing bugs unresolved |
| **OpenCode** | 10 (top issues) | 10 open PRs | **v1.17.10** (today) | **Crisis**: Windows Bun segfault outbreak |
| **Pi** | 10 noteworthy | 9 important PRs | No new release today | Active: stream reliability + Bedrock Mantle |
| **Qwen Code** | 8 total (all updated) | 10 open PRs | **v0.19.2** (today) | Steady: voice dictation + session APIs |
| **DeepSeek TUI** | 10 (noteworthy) | 10 open/closed PRs | No new release today | Stable: Fleet loadout automation in progress |

**Key Observation:** Claude Code, OpenCode, and Codex are in crisis/systems mode with scalability or regression failures. Gemini CLI, Qwen Code, and Pi show steady measured progress. Kimi Code and DeepSeek TUI have lower community engagement.

---

## 3. Shared Feature Directions

### A. Safety/Classifier Tuning & False Positive Reduction
- **Claude Code** (#70801, #70792, #70773, #70821, #70796, #70790, #70808 — sworrl reporting 20+ issues)
- **Gemini CLI** (#27966 — case-insensitive blocklist, #26525 — secret redaction)
- **Copilot CLI** (indirect: preToolUse rewrite confirmation #2643)
- **Cross-cutting demand:** Context-aware safety filters that distinguish legitimate sysadmin/security work (IAM hardening, malware forensics, autostart tools) from malicious behavior. Current classifiers use pattern-matching on keywords ("cloud-iam") causing severe workflow disruption.

### B. Multi-Agent Orchestration Reliability
- **Claude Code** (#64550 — stuck agent pointer, #69033 — memory-aware throttling)
- **Gemini CLI** (#21409 — agent hangs, #22323 — false success reports after MAX_TURNS)
- **Kimi Code** (#1942 — MCP propagation to subagents)
- **OpenCode** (#32767 — ESC interrupt for subagent sessions)
- **Pi** (#6054 — parallel agent tasks)
- **Core challenges:** Subagent lifecycle (stuck pointers, resurrection), concurrency management (memory-aware vs count-based), recovery from turn limits without misleading success signals.

### C. MCP Integration & Plugin Ecosystem
- **Claude Code** (#42138 — Telegram MCP notifications dropped)
- **OpenAI Codex** (#28529 — OAuth for HTTP MCP servers, #28407 — non-blocking MCP startup)
- **Gemini CLI** (#27979 — MCP resource output security)
- **Kimi Code** (#1942 — MCP configs to subagents, #2469 — MCP server path bug)
- **OpenCode** (MCP server instructions, resource templates, config validation #33845, #33853, #33861)
- **Common need:** Predictable MCP lifecycle, OAuth support, config validation, subagent propagation.

### D. Session State & Persistence
- **Claude Code** (`/rewind` after `/clear`)
- **Copilot CLI** (#3680, #3596 — model selection breaks on resume)
- **OpenCode** (#33328 — session state across project switches)
- **Qwen Code** (#5855, #5863 — richer session status APIs)
- **Pi** (#6002 — session truncation bug)
- **Trend:** Better resume semantics, model selection persistence, RPC for external consumers.

### E. Token Efficiency & Context Management
- **Claude Code** (pending attention)
- **Gemini CLI** (#22745 — AST-aware file reads)
- **Kimi Code** (#2472 — compaction wastes ~20k tokens)
- **Qwen Code** (#401 — streaming timeout, #5861 — non-streaming compression)
- **Copilot CLI** (indirect via model selection)
- **User demand:** Avoid redundant token consumption during compaction, streaming defaults for large contexts, AST-based file access to reduce turns.

### F. Windows Platform Support
- **OpenCode** (#33742 — Bun segfault, #33767 — Ctrl+C unresponsive, #33871 — IME crash, #33863 — zombie git processes)
- **DeepSeek TUI** (#3572 — environment variable inheritance)
- **Codex** (#29782 — COM+ registry error, #30024 — CreateProcessAsUserW failure)
- **Copilot CLI** (#3925 — Linux AppImage, but Windows proxy issues)
- **Cross-cutting:** Windows remains a second-class platform. Segfaults, environment inheritance, IME crashes, and process management bugs plague all tools.

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | Copilot CLI | OpenCode | Pi | Qwen Code |
|-----------|-------------|--------------|------------|-------------|----------|-----|-----------|
| **Primary Focus** | Safety-conscious orchestration | Extensible hooks & MCP | Subagent orchestration | Enterprise integration | Experimental rapid iteration | Lightweight streaming | Voice + desktop integration |
| **Target User** | Power developers/security teams | Enterprise teams | Cloud/backend devs | GitHub ecosystem | Bleeding-edge early adopters | TUI purists | Multilingual/voice users |
| **Technical Approach** | Agent Teams, ClAudit classifier | Hook parity, deferred executors | ACP, Auto Memory, AST tools | GitHub integration, skills system | Bun/WASM, OpenTUI | Modular extensions, Bedrock | Server-side daemon, LSP |
| **Differentiator** | Aggressive safety (too aggressive per community) | Most complete hook system | Deep subagent introspection | Strongest enterprise proxy/GitHub integration | Fastest iteration cycle | Best streaming reliability | Unique voice dictation |
| **Biggest Pain Point** | False positive safety blocks | Quota accounting bugs | Agent hangs | Resumed session auth | Windows stability | Stream hangs | Streaming timeouts |

### Key Strategic Differences:

1. **Claude Code** is taking the most aggressive safety stance, which backfired spectacularly today. The ClAudit auto-mode classifier is blocking legitimate work at scale. This is a trust crisis.

2. **OpenAI Codex** has the most sophisticated hook/executor architecture but is being undermined by backend quota/rate-limit issues that are entirely server-side. The community is frustrated with systemic infrastructure problems, not feature gaps.

3. **Gemini CLI** has the strongest focus on subagent introspection and memory systems but struggles with basic agent reliability (hangs, false success reports). Its evals framework (#24353) suggests long-term quality investment.

4. **OpenCode** is iterating fastest but at the cost of Windows stability. The Bun 1.4 migration (#33822) is a bet on newer runtime stability. The `--mini` mode and MCP improvements show strategic depth.

5. **Pi** differentiates on streaming reliability and provider diversity (Bedrock Mantle). Its extension system is modular but less comprehensive than Codex hooks.

6. **Qwen Code** uniquely targets voice dictation and desktop integration, carving a niche underserved by others.

---

## 5. Community Momentum & Maturity

### High Momentum (Rapid Iteration, High Engagement)
- **OpenCode** — 26 upvotes on #33742 in 1 day; 10 PRs in flight. Community is vocal and patches are landing fast. However, the Windows crash crisis signals fragility at scale.
- **Claude Code** — Extremely active but in crisis mode. The 20+ false positive reports indicate either classifier degradation or a community stress-testing boundaries. Two patch releases today suggest the team is responsive.
- **Pi** — 70 comments on #4945 (Codex reliability); 30 upvotes. Active feature development with orchestrator daemon and Bedrock Mantle support.

### Steady Mid-Maturity
- **Gemini CLI** — Measured pace with security-focused PRs (thought leakage, blocklist). Community engagement is moderate but constructive. The eval framework (#24353) indicates mature development practices.
- **OpenAI Codex** — High issue volume but many are infrastructure complaints (quota, rate limits). The PR pipeline is healthy but many features are server-side gated. Community trust is strained.
- **Copilot CLI** — Low PR activity (2 today) but stable releases. Community is vocal about specific bugs (resumed session auth) but not in crisis mode.

### Lower Momentum / Niche
- **Kimi Code** — 2 PRs closed after months; only 4 issues updated today. A 5-month-old looping bug (#640) unresolved. Community appears small or dormant.
- **DeepSeek TUI** — Steady but quiet. The Fleet loadout feature (#3205) is the main driver. Rebranding to CodeWhale suggests organizational focus.
- **Qwen Code** — 8 issues updated, 10 PRs open. Low upvote counts suggest smaller community but steady development.

### Maturity Assessment
| Tool | Maturity Rating | Risk Level | Response Quality |
|------|----------------|------------|------------------|
| Claude Code | **Medium** | High (safety backlash) | Fast patches, but root causes persist |
| OpenAI Codex | **Medium-High** | High (infrastructure) | Server-side dependencies delay fixes |
| Gemini CLI | **Medium-High** | Low | Careful, security-first releases |
| Copilot CLI | **High** | Low | Stable, but slow on feature requests |
| OpenCode | **Low-Medium** | Very High (Windows crash) | Fast response, but broke platform |
| Pi | **Medium** | Medium | Good trajectory, growing |
| Qwen Code | **Medium** | Low | Quiet but consistent |
| DeepSeek TUI | **Medium** | Low | Niche but stable |

---

## 6. Trend Signals

### For AI Developer Tool Builders:

1. **False Positive Safety Crisis is Reaching Inflection Point.** Claude Code's 20+ reports in one day signal that overly aggressive classifiers are becoming a systemic industry problem. The "sworrl pattern" (single user filing 20 issues) suggests either a coordinated test or a user whose entire workflow (cybersecurity, sysadmin) is being blocked. **Recommendation:** Build context-aware safety filters that use user intent signals (project type, file paths, user roles) rather than keyword pattern-matching.

2. **Multi-Agent Orchestration is the Next Frontier — and It's Hard.** Every tool with subagent support (Claude Code, Gemini CLI, Kimi Code, OpenCode, Pi) has bugs in lifecycle management, concurrency, and error reporting. The shift from single-turn to multi-agent workflows is exposing fundamental reliability gaps. **Recommendation:** Invest in telemetry and observability for subagent traces before adding more features.

3. **MCP is Becoming the USB-C of AI Tool Extensibility.** All major tools are shipping MCP improvements simultaneously. The convergence on MCP as the standard plugin protocol means tool interoperability is achievable. OAuth for MCP (#28529 Codex), notifications (#42138 Claude), and resource templates (#33861 OpenCode) are early signs of protocol maturation. **Recommendation:** If building an AI CLI tool, adopt MCP as the primary extension mechanism.

4. **Token Efficiency is the Silent Product Killer.** Sessions wasting 20k tokens on compaction reloads (Kimi Code), requests timing out at 6 seconds (Qwen Code), and non-streaming compression causing gateway timeouts (#5861) all point to developers hitting real cost ceilings. Users are becoming token-aware and demanding configurability. **Recommendation:** Default to streaming, expose token budgets, and implement memory-aware throttling.

5. **Windows is the Unresolved Platform Gap.** Every tool has Windows-specific bugs that erode trust: segfaults, environment inheritance failures, IME crashes, sandbox execution errors. The Windows developer population is vocal and underserved. **Recommendation:** Allocate dedicated platform QA for Windows; Bun-native tools should monitor the Bun 1.3.14→1.4 migration closely.

6. **Server-Side Infrastructure Bottlenecks are Becoming Community Issues.** Codex's quota accounting bugs, Claude's rate limiting, and Gemini's auth regressions show that backend reliability directly impacts CLI tool reputation. The community does not distinguish between client bugs and API failures — they blame the tool. **Recommendation:** Implement client-side rate-limit resilience (retry headers, exponential backoff) and expose quota status transparently in the TUI.

### For Developers Choosing a Tool:

- **Stability:** Copilot CLI + Qwen Code have the fewest "cannot work" bugs today.
- **Features:** Claude Code (risky safety) or Gemini CLI (agent introspection) for power users; Codex for enterprise hook parity.
- **Windows Safety:** Avoid OpenCode v1.17.10; use Copilot CLI or Gemini CLI.
- **Multi-Agent:** Gemini CLI has the most sophisticated subagent system but watch for hangs.
- **Voice/Desktop:** Qwen Code is the only tool with dedicated voice dictation.
- **Token Conscious:** Kimi Code's compaction bug and Qwen Code's streaming timeout are warning signs for high-volume users.

*Report generated from public GitHub issue/PR data as of 2026-06-25 23:59 UTC. Sample sizes vary by tool. Data reflects only community-reported issues and may not represent all active development.*

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
**Data snapshot:** 2026-06-25 | Source: github.com/anthropics/skills

---

## 1. Top Skills Ranking

The following pull requests have drawn the most community attention (by comment activity and cross-referencing in Issues).

### 1. `skill-creator` evaluation pipeline overhaul — PR [#1298](https://github.com/anthropics/skills/pull/1298)
**Status:** Open | **Author:** MartinCajiao
A critical fix for `run_eval.py` that has been reporting **0% recall** for every skill description regardless of content. The root cause involves the eval artifact not being installed as a real skill, plus Windows stream-reading failures and broken trigger detection. This PR directly addresses Issue [#556](https://github.com/anthropics/skills/issues/556) (12 comments, 7 👍) and Issue [#1169](https://github.com/anthropics/skills/issues/1169), making it the most-impacted infrastructure fix in the pipeline.

### 2. `document-typography` — PR [#514](https://github.com/anthropics/skills/pull/514)
**Status:** Open | **Author:** PGTBoos
A quality-of-life skill preventing orphan word wrap, widow paragraphs, and numbering misalignment in generated documents. The discussion highlights that these issues are universal across Claude-generated content and that users rarely request typographic fixes explicitly, making a proactive skill highly valuable. Currently unmerged with sustained interest since March.

### 3. `odt` (OpenDocument) support — PR [#486](https://github.com/anthropics/skills/pull/486)
**Status:** Open | **Author:** GitHubNewbie0
Adds full ODT/ODS creation, template filling, and ODT-to-HTML conversion. The community discussion centers on LibreOffice compatibility and the gap between proprietary format support (PDF, DOCX) and open-standard alternatives. The skill has been actively discussed April through June.

### 4. `frontend-design` clarity overhaul — PR [#210](https://github.com/anthropics/skills/pull/210)
**Status:** Open | **Author:** justinwetch
A comprehensive revision to make the frontend-design skill actionable and internally coherent, ensuring every instruction is executable within a single conversation. Discussion critiques the original skill's vagueness and emphasizes the need for concrete design principles Claude can follow deterministically.

### 5. `skill-quality-analyzer` / `skill-security-analyzer` — PR [#83](https://github.com/anthropics/skills/pull/83)
**Status:** Open | **Author:** eovidiu
Two meta-skills for evaluating other Skills across structure, documentation, and security dimensions. Community interest reflects demand for quality assurance tooling in the ecosystem — a "skill for skills." Discussion (spanning Nov 2025–Jan 2026) centered on evaluation methodology and false-positive reduction.

### 6. `testing-patterns` — PR [#723](https://github.com/anthropics/skills/pull/723)
**Status:** Open | **Author:** 4444J99
Covers the full testing stack: philosophy (Testing Trophy model), unit testing (AAA pattern), React component testing (Testing Library), and edge-case guidance. The discussion highlights the need for testing best practices to be encoded as a skill rather than left to ad-hoc prompting.

### 7. `codebase-inventory-audit` — PR [#147](https://github.com/anthropics/skills/pull/147)
**Status:** Open | **Author:** p19dixon
A systematic 10-step workflow for identifying orphaned code, unused files, and documentation gaps, producing a `CODEBASE-STATUS.md` artifact. Discussion focuses on integration with existing CI/CD pipelines and the skill's token budget for large monorepos.

### 8. `appdeploy` deployment skill — PR [#360](https://github.com/anthropics/skills/pull/360)
**Status:** Open | **Author:** avimak
Enables Claude to deploy and manage full-stack web apps via AppDeploy.ai, including lifecycle management. Community interest centers on whether deployment skills should be external tool integrations or core Claude capabilities.

---

## 2. Community Demand Trends

Analysis of the 14 most-commented Issues reveals five concentrated demand directions:

**🔒 Security & Trust Boundaries** — Issue [#492](https://github.com/anthropics/skills/issues/492) (19 comments, highest across all Issues) exposes that community skills distributed under the `anthropic/` namespace create trust-boundary vulnerabilities. Users are wary of granting elevated permissions to skills they may mistake as official. This is the #1 community concern by signal strength.

**🏢 Enterprise & Org Workflows** — Issue [#228](https://github.com/anthropics/skills/issues/228) (14 comments, 7 👍) demands org-wide skill sharing without manual `.skill` file distribution. Issue [#1175](https://github.com/anthropics/skills/issues/1175) raises SharePoint Online permission handling and context window management. Combined, these signal strong enterprise adoption friction.

**🪟 Windows Compatibility Blockers** — Issues [#1061](https://github.com/anthropics/skills/issues/1061), [#556](https://github.com/anthropics/skills/issues/556), and [#1169](https://github.com/anthropics/skills/issues/1169) (collectively ~18 comments) report that `skill-creator` scripts are effectively broken on Windows due to `PATHEXT`, `cp1252` encoding, and pipe-select issues. Multiple independent PRs (#1298, #1099, #1050) attempt to fix overlapping subsets, indicating a fragmented fix landscape.

**🔄 Agent Governance & Safety Patterns** — Issue [#412](https://github.com/anthropics/skills/issues/412) proposes an `agent-governance` skill for policy enforcement, threat detection, and audit trails. This reflects growing concern about agent safety as Claude Code is used in production.

**🧠 Compact Memory / Symbolic Notation** — Issue [#1329](https://github.com/anthropics/skills/issues/1329) proposes a `compact-memory` skill for long-running agents, using symbolic notation to reduce context window consumption. This is a novel direction addressing the practical limits of persistent context.

---

## 3. High-Potential Pending Skills

These PRs have active community engagement and are likely to land in the near term:

| PR | Skill | Last Updated | Why It May Land Soon |
|---|---|---|---|
| [#1298](https://github.com/anthropics/skills/pull/1298) | `skill-creator` eval fix | 2026-06-23 | Addresses 3 linked Issues (#556, #1169, #1061); most active PR this month |
| [#1323](https://github.com/anthropics/skills/pull/1323) | `skill-creator` trigger fix | 2026-06-23 | Follow-up to #1298; targets the same 0% recall bug from a different angle |
| [#514](https://github.com/anthropics/skills/pull/514) | `document-typography` | 2026-03-13 | Self-contained, low-risk, high-value quality-of-life skill |
| [#486](https://github.com/anthropics/skills/pull/486) | `odt` | 2026-04-14 | Fills a clear open-format gap; discussion has converged |
| [#723](https://github.com/anthropics/skills/pull/723) | `testing-patterns` | 2026-04-21 | Comprehensive and well-structured; demand validated by community |
| [#360](https://github.com/anthropics/skills/pull/360) | `appdeploy` | 2026-05-04 | External integration pattern; may serve as template for future deployment skills |

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is fixing the broken `skill-creator` evaluation pipeline — the foundational tool for developing all other Skills — while simultaneously demanding trust-boundary protections against namespace impersonation, reflecting a mature ecosystem that recognizes infrastructure reliability and security as prerequisites for growth.**

---

# Claude Code Community Digest — 2026-06-25

---

## Today’s Highlights

Two patch releases landed today: **v2.1.191** brings long-awaited `/rewind` support after `/clear`, fixes scroll-jumping during streaming, and prevents background agents from resurrecting after being stopped. Meanwhile, the community is seeing an **unprecedented wave of false-positive safety blocks** — a single reporter (sworrl) filed over 20 issues in 24 hours detailing ClAudit auto-mode classifier denials and cybersecurity/AUP filter false positives on legitimate admin work. Three pull requests are open, addressing server rate limiting and a critical SSRF vulnerability in the security-guidance plugin.

---

## Releases

### v2.1.191
- **New**: `/rewind` now works after `/clear` — resume a conversation from before the clear point.
- **Fixed**: Scroll position no longer jumps to bottom while reading earlier output during a streaming response.
- **Fixed**: Background agents no longer resurrect after being stopped (stopping an agent from the tasks panel now actually kills it).

### v2.1.190
- Bug fixes and reliability improvements (no further details).

---

## Hot Issues

1. **#42138 – Telegram plugin: inbound MCP notifications not injected**  
   *Author: Pampidu*  
   This long-standing bug (opened Apr 1) affects anyone using the Telegram MCP plugin — `notifications/claude/channel` events are silently dropped. High comment count (7) suggests many users rely on this integration.  
   [GitHub](https://github.com/anthropics/claude-code/issues/42138)

2. **#70801 – ClAudit auto-mode classifier denies legitimate autostart launcher**  
   *Author: sworrl*  
   One of many false-positive reports today: the harness blocked a persistent autostart launcher for a legitimate watch tool, incorrectly classifying it as autonomous behavior.  
   [GitHub](https://github.com/anthropics/claude-code/issues/70801)

3. **#64550 – Agent Teams: team-lead pointer sticks on teammate after long session**  
   *Author: suhang56*  
   Critical Windows bug: after a compacted session the lead agent’s “active agent” pointer becomes stuck, causing all spawns to fail with “Teammates cannot spawn other teammates.” Blocks complex multi-agent workflows.  
   [GitHub](https://github.com/anthropics/claude-code/issues/64550)

4. **#70792 – Another ClAudit false positive: autostart tool for bulk file operations**  
   *Author: sworrl*  
   Similar to #70801 — the classifier blocked a legitimate script that auto-files bug reports, mistaking it for unauthorized automation.  
   [GitHub](https://github.com/anthropics/claude-code/issues/70792)

5. **#70773 – ClAudit blocks long-running watcher as “mass false positive”**  
   *Author: sworrl*  
   A file-watcher used for development was blocked by the auto-mode classifier. The user notes the same pattern is reported across multiple issues today.  
   [GitHub](https://github.com/anthropics/claude-code/issues/70773)

6. **#69033 – Workflow harness: memory-aware throttling for subagent fan-out**  
   *Author: spitfire94*  
   A large deep-research workflow (84–92 subagents) caused OOM crash because concurrency is count-based (`min(16, cores-2)`) rather than memory-aware. Community needs this for reliable large-scale agent orchestration.  
   [GitHub](https://github.com/anthropics/claude-code/issues/69033)

7. **#70821 – Cybersecurity false positive blocks malware forensics**  
   *Author: sworrl*  
   The safety filter blocked analysis of a captured malware sample for incident forensics — a standard defensive cybersecurity task. Filed as a false positive with request ID.  
   [GitHub](https://github.com/anthropics/claude-code/issues/70821)

8. **#70796 – Cybersecurity false positive on cloud IAM work**  
   *Author: sworrl*  
   The classifier pattern-matched on cybersecurity terminology (e.g., “cloud-iam”) and blocked benign administrative work.  
   [GitHub](https://github.com/anthropics/claude-code/issues/70796)

9. **#70790 – AUP false positive on cloud IAM task**  
   *Author: sworrl*  
   Usage-policy block triggered on a routine, authorized cloud-identity management operation.  
   [GitHub](https://github.com/anthropics/claude-code/issues/70790)

10. **#70808 – AUP false positive on IAM role/policy hardening**  
    *Author: sworrl*  
    Reviewing and hardening IAM roles on own tenant flagged as policy violation. The user explicitly notes the task is on systems they own and are authorized to manage.  
    [GitHub](https://github.com/anthropics/claude-code/issues/70808)

---

## Key PR Progress

Three pull requests are open as of today’s digest:

1. **#70634 – fix: handle server rate limiting during normal usage**  
   *Author: Siliconlive*  
   Implements client-side handling for server rate limits (likely Anthropic API). Closes #70631.  
   [GitHub](https://github.com/anthropics/claude-code/pull/70634)

2. **#70633 – fix: Handle rate limiting headers for Anthropic API**  
   *Author: Siliconlive*  
   Parses and respects `Retry-After` and related headers from the API, improving resilience during high-throughput usage.  
   [GitHub](https://github.com/anthropics/claude-code/pull/70633)

3. **#70582 – fix: application accepts user-controlled URLs in llm.py (CRITICAL)**  
   *Author: orbisai0security*  
   Fixes a critical SSRF-style vulnerability in `plugins/security-guidance/hooks/llm.py` where user-controlled URLs could be accepted without validation. Marked as critical severity by a multi-agent AI scanner.  
   [GitHub](https://github.com/anthropics/claude-code/pull/70582)

---

## Feature Request Trends

- **Safety classifier tune-down**: A dominant theme from today’s issue flood is the demand for more accurate, context-aware safety filters — especially for cybersecurity, AUP, and auto-mode classifier denials. Users performing legitimate security/administration work are being blocked repeatedly, hurting trust and productivity.
- **Memory-aware subagent throttling**: #69033 and similar comments indicate a strong need for dynamic concurrency control in workflow harnesses, replacing simple CPU-core-based caps with memory-aware limits.
- **Agent Teams stability**: Bugs like #64550 (stuck agent pointer) and reports of agent resurrection highlight that multi-agent orchestration still needs robustness improvements.
- **Plugin/Integration reliability**: The Telegram MCP issue (#42138) points to a broader desire for first-class MCP notification support and plugin lifecycle reliability.
- **Rate limiting resilience**: The two open PRs (#70633, #70634) reflect community demand for graceful API rate-limit handling.

---

## Developer Pain Points

- **False-positive safety blocks**: The #1 pain point today. Over 20 issues filed in a single day by one user (sworrl), highlighting that the ClAudit auto-mode classifier and the cybersecurity/AUP filters are far too aggressive. Blocks occur during routine IAM management, defensive hardening, even placeholder keystrokes (#70812, #70813). This actively interferes with developer workflows.
- **Agent lifecycle bugs**: The “active agent pointer” stickiness (#64550) and background agents resurrecting after stop (fixed in v2.1.191) show that agent management is still brittle under load or long sessions.
- **Memory crashes with large agent fan-outs**: The count-based concurrency cap leads to OOM crashes (#69033), forcing developers to guess safe subagent counts.
- **Scroll position jumping**: A frustrating UI bug during streaming (fixed in v2.1.191) that previously made reading long outputs painful.
- **Lack of rate-limit handling**: Users hitting Anthropic API rate limits without proper retry/backoff logic (addressed by pending PRs #70633, #70634).
- **Security vulnerability exposure**: The critical SSRF issue in the security-guidance plugin (#70582) raises concerns about code quality in official plugins — though a fix is already proposed.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-06-25

## Today’s Highlights

A flurry of rate-limit and quota accounting bugs dominated the issue tracker today, with multiple Pro and Business users reporting inexplicably drained limits, stalled resets, and “model at capacity” errors on both CLI and Desktop. On the development side, the team shipped **rust-v0.142.2** with MCP tool search by default and system proxy support for macOS, while several PRs laid groundwork for deferred executors, environment inheritance, and OAuth-based MCP servers.

## Releases

Two new stable releases and several alpha tags landed in the last 24 hours:

- **rust-v0.142.2**
  - MCP tools now use tool search by default when supported, improving tool discovery while preserving compatibility with older models and providers. ([#29486](https://github.com/openai/codex/issues/29486))
  - macOS authentication clients can honor system proxy, PAC, and WPAD settings when `respect_system_proxy` is enabled. ([#26709](https://github.com/openai/codex/issues/26709))

- **rust-v0.142.1**
  - Added opt-in Windows system proxy support for authentication, including PAC, WPAD, static proxies, and bypass rules. ([#26708](https://github.com/openai/codex/issues/26708))
  - Full changelog: [v0.142.0…v0.142.1](https://github.com/openai/codex/compare/rust-v0.142.0...rust-v0.142.1)

- **rust-v0.143.0-alpha.15**, **alpha.16**, **alpha.21** — no release notes beyond version bumps.

## Hot Issues

| # | Title & Link | Why It Matters | Community Reaction |
|---|-------------|----------------|-------------------|
| #21753 | [Full Claude Code Hook Parity (29+)](https://github.com/openai/codex/issues/21753) | Umbrella for bringing Codex hooks to full Claude Code parity – a major automation surface gap. | 20 comments, 18 👍; highly upvoted, one of the longest-running enhancement requests. |
| #29760 | [“Selected model is at capacity” on CLI](https://github.com/openai/codex/issues/29760) | GPT Pro users hitting capacity errors with `gpt-5.4 high`; blocks productivity. | 20 comments, 1 👍; many users reporting similar. |
| #30008 | [Same model capacity error on App + CLI](https://github.com/openai/codex/issues/30008) | Duplicate report but on App side; suggests a systemic backend issue. | 16 comments, 0 👍; frustrated users. |
| #29968 | [Pro20x subscription behaves like Plus](https://github.com/openai/codex/issues/29968) | Billing/entitlement bug causing misapplied rate limits. | 16 comments, 10 👍; high visibility, implies potential charging error. |
| #29955 | [100 credits gone after 1 message](https://github.com/openai/codex/issues/29955) | Quota drained instantly; severe UX failure for Pro*5 users. | 15 comments, 0 👍; clear reproduction of bug. |
| #30002 | [Server-side quota accounting over-reports consumption after 5h reset](https://github.com/openai/codex/issues/30002) | Detailed analysis showing 1.35M vs 156M token discrepancy. Alerts to a fundamental accounting bug. | 12 comments, 3 👍; technical depth appreciated. |
| #27189 | [5h usage limit stuck at 99% for 48+ hours](https://github.com/openai/codex/issues/27189) | Reset not happening; user locked out indefinitely. | 8 comments, 2 👍; plus users affected. |
| #15477 | [Codex Cloud auto review silent fail + quota mismatch](https://github.com/openai/codex/issues/15477) | Three bugs in one: stale GitHub token, contradictory quota display, silent failures. Pro users blocked from automated review. | 7 comments, 4 👍; long-standing issue since March. |
| #29963 | [Quota consumption serious bug on Pro20x](https://github.com/openai/codex/issues/29963) | Another entitlement bug affecting high-tier users; includes screenshot evidence. | 5 comments, 5 👍; strong signal from paying customers. |
| #27598 | [“Workspace out of credits” despite remaining limits](https://github.com/openai/codex/issues/27598) | Business users blocked across VS Code and Desktop; false positive lockout. | 4 comments, 0 👍; Business tier critical for teams. |

## Key PR Progress

| # | Title & Status | Description |
|---|----------------|-------------|
| #29923 | [Support external clock sleeps](https://github.com/openai/codex/pull/29923) — Open | Adds experimental `currentTime/sleep` notification and `wake` request for external clocks, with precise pending-sleep management. |
| #29845 | [Plumb explicit application paths through Windows launchers](https://github.com/openai/codex/pull/29845) — Open | Foundation for Windows unified-exec resolution; introduces `WindowsProcessLaunch` value without changing command resolution yet. |
| #30017 | [Refresh turn diff roots from step context](https://github.com/openai/codex/pull/30017) — Open | Fixes path formatting when deferred environments attach mid-turn; `TurnDiffTracker` now caches display roots from actual step context. |
| #28529 | [Support OAuth for HTTP MCP servers from selected executor plugins](https://github.com/openai/codex/pull/28529) — Closed | Complete OAuth discovery and token exchange through executor network boundary; includes PKCE verification and mock token endpoint. |
| #30016 | [Inherit current step environments in subagents](https://github.com/openai/codex/pull/30016) — Open | Subagents now get the environment snapshot from the sampling request, not the stale `TurnContext`, enabling deferred executors to work correctly. |
| #29917 | [Handle post-init requests concurrently in exec-server](https://github.com/openai/codex/pull/29917) — Open | Allows independent RPCs to make progress while another long-polls, without breaking `initialize`/`initialized` ordering. |
| #29935 | [Attribute app-server analytics by thread originator](https://github.com/openai/codex/pull/29935) — Open | Desktop Work threads now send analytics with their own `product_client_id` instead of inheriting the connection-level ID. |
| #29990 | [Parallelize environment skill loading](https://github.com/openai/codex/pull/29990) — Closed | Avoids request waterfall by polling skill parse futures concurrently, capped at 64 in-flight. |
| #28406 | [Reuse tool router within a turn](https://github.com/openai/codex/pull/28406) — Closed | Lazy-builds `ToolRouter` once per turn and reuses it across sampling requests; tool schema stays immutable mid-turn. |
| #28407 | [Avoid blocking on optional MCP startup](https://github.com/openai/codex/pull/28407) — Closed | Lists tools from MCP servers concurrently; optional pending servers with a cached snapshot are exposed without blocking. |

## Feature Request Trends

- **Hooks parity with Claude Code** (#21753) remains the most-requested automation feature, with demand for complete lifecycle hooks (start, stop, systemMessage severities).
- **Internationalization/l10n** (#30025) – Chinese (Simplified) and other languages requested for the Electron desktop app.
- **Local-first plugin architecture** (#30023) – Users want plugins installed locally per machine, separate from account-level connectors, and OAuth to be machine-scoped.
- **Improved terminal Markdown rendering** (#30007) – Better support for wide math output, tables, and mixed-language blocks in the CLI TUI.
- **Home directory expansion in Unix socket paths** (#29884) – Enterprise users need `~` expansion for sandbox and authentication sockets.
- **Stop hook systemMessage severity levels** (#29906) – Request for `info`/`notice` severity instead of forced `warning:` prefix.

## Developer Pain Points

- **Rate-limit and quota accounting bugs** – Today’s top frustration. Reports of “Selected model at capacity,” instant credit drains, stuck 99% limits, and mismatched consumption between Pro and Plus tiers suggest backend accounting is fundamentally broken for many users. Several issues (#29760, #30008, #29968, #29955, #30002, #27189) involve silent failures or false positives that block all work.
- **Windows sandbox and execution failures** – Multiple open bugs: COM+ registry errors on `apply_patch` (#29782), Bubblewrap loopback errors on Ubuntu (#29908), `CreateProcessAsUserW` failed (#30024), blocked per-user executables (#27171), and empty `.git` folders spawned repeatedly (#29492). Windows users face a fragmented experience.
- **Context window exhaustion and recovery** – Codex reporting “ran out of room” mid-thread, leaving threads unrecoverable (#28920). This is a workflow killer for long sessions.
- **Slow thread switching on Windows** (#29187) – Desktop app performance disparity between platforms.
- **Plugins catalog hiding local plugins** (#29633) – Regression in 0.142.0 where `/plugins` remote catalog obscures installed openai-curated plugins in TUI.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-06-25

## Today’s Highlights
A new nightly release (`v0.49.0-nightly.20260625`) ships a critical path‑traversal fix for skill installation and improvements to pending tools handling. Community attention remains fixed on agent reliability: the top‑voted issue (8 👍) reports a generalist agent hang, and several bugs highlight subagent lifecycle and tool‑selection problems. Meanwhile, PRs tackling thought leakage, security blocklists, and `@`‑prefixed file paths are moving through review, signalling a strong focus on core safety and correctness.

## Releases
**[v0.49.0-nightly.20260625.gd845bc5d4](https://github.com/google-gemini/gemini-cli/releases/tag/v0.49.0-nightly.20260625.gd845bc5d4)**  
- `fix(cli)`: prevent path traversal vulnerabilities during skill install ([#27767](https://github.com/google-gemini/gemini-cli/pull/27767))  
- `fix/pending tools and trust overrides` ([#27854](https://github.com/google-gemini/gemini-cli/pull/27854))

## Hot Issues (Top 10 by Conversation)

| Issue | Title | Comments | 👍 | Why it matters |
|-------|-------|----------|----|----------------|
| [#22323](https://github.com/google-gemini/gemini-cli/issues/22323) | Subagent recovery after MAX_TURNS is reported as GOAL success, hiding interruption | 8 | 2 | Misleading success reports mask subagent failures, harming debugging and trust. |
| [#24353](https://github.com/google-gemini/gemini-cli/issues/24353) | Robust component level evaluations | 7 | 0 | Epic for a proper eval framework; 76 behavioral tests already exist but need structure. |
| [#22745](https://github.com/google-gemini/gemini-cli/issues/22745) | Assess impact of AST‑aware file reads, search, and mapping | 7 | 1 | Could drastically reduce token waste and turn count by understanding code structure. |
| [#21409](https://github.com/google-gemini/gemini-cli/issues/21409) | Generalist agent hangs | 7 | 8 | **Most upvoted** – agent deadlocks when users don’t explicitly disable sub‑agents. |
| [#21968](https://github.com/google-gemini/gemini-cli/issues/21968) | Gemini does not use skills and sub‑agents enough | 6 | 0 | Custom skills ignored unless explicitly requested; undermines extensibility. |
| [#26525](https://github.com/google-gemini/gemini-cli/issues/26525) | Add deterministic redaction and reduce Auto Memory logging | 5 | 0 | Security: secrets may leak through Auto Memory before redaction. |
| [#26522](https://github.com/google-gemini/gemini-cli/issues/26522) | Stop Auto Memory from retrying low‑signal sessions indefinitely | 5 | 0 | Inefficient processing wastes API quota on empty transcripts. |
| [#28092](https://github.com/google-gemini/gemini-cli/issues/28092) | Google account not authorized for gemini | 4 | 0 | Authentication regression blocks new users. |
| [#25166](https://github.com/google-gemini/gemini-cli/issues/25166) | Shell command execution gets stuck with “Waiting input” after command completes | 4 | 3 | Frequent hang in core shell integration, disrupts workflow. |
| [#21983](https://github.com/google-gemini/gemini-cli/issues/21983) | Browser subagent fails in Wayland | 4 | 1 | Linux desktop users cannot rely on browser automation. |

## Key PR Progress

| PR | Title | Size/Tags | Summary |
|----|-------|-----------|---------|
| [#27971](https://github.com/google-gemini/gemini-cli/pull/27971) | `fix(core)`: strip thoughts from scrubbed history turns and resolve thought leakage | size/m, need‑issue | Prevents the model’s internal monologue from contaminating future turns, a major source of looping. |
| [#28053](https://github.com/google-gemini/gemini-cli/pull/28053) | `fix(core-tools)`: resolve defensive path resolution for at‑reference files and fix macOS tests | size/xl, need‑issue | Fixes “File not found” when model uses `@`‑prefixed paths; adds comprehensive macOS test coverage. |
| [#27966](https://github.com/google-gemini/gemini-cli/pull/27966) | `fix(security)`: enforce case‑insensitive sensitive path blocklist and vscode hitl | size/m, need‑issue | Closes prompt‑injection bypass via case variations (`.GIT`, `.Env`). |
| [#27986](https://github.com/google-gemini/gemini-cli/pull/27986) | `feat(acp)`: report cached and thought tokens in PromptResponse.usage | size/m, non‑interactive | ACP clients can now accurately track cost; cached input was previously invisible. |
| [#27994](https://github.com/google-gemini/gemini-cli/pull/27994) | `fix(core)`: insert skill/agent content literally in system prompt substitutions | size/s, agent | Avoids regex‑special‑character injection when skills contain `$` or backreferences. |
| [#27979](https://github.com/google-gemini/gemini-cli/pull/27979) | Wrap `read_mcp_resource` output with `wrapUntrusted()` for consistency | size/s, security | MCP resource output no longer bypasses untrusted‑content handling. |
| [#28094](https://github.com/google-gemini/gemini-cli/pull/28094) | `fix(a2a-server)`: deep‑merge user and workspace settings | size/m, core | Nested config keys (e.g., `tools.*`) now correctly merge instead of being overwritten. |
| [#27636](https://github.com/google-gemini/gemini-cli/pull/27636) | `perf`: optimize VirtualizedList and fix click handling | size/xl, p1 | Major terminal rendering performance improvement for large histories. |
| [#27975](https://github.com/google-gemini/gemini-cli/pull/27975) | `docs`: add Linux native dependency FAQ | size/s, docs | Helps Linux users resolve build‑tool and D‑Bus header errors. |
| [#28131](https://github.com/google-gemini/gemini-cli/pull/28131) | Fix `no_proxy` test | size/xs, need‑issue | Prevents test failures in environments pre‑configured with `NO_PROXY`. |

## Feature Request Trends
- **AST‑aware code tools** (#22745, #22746) – multiple investigations propose using abstract syntax trees for more precise file reads, codebase mapping, and tool calls to cut turns and token waste.
- **Agent introspection & self‑awareness** (#21432, #22598) – users want the CLI to know its own flags, hotkeys, and subagent trajectories so it can serve as its own doc and share subagent traces for debugging.
- **Memory system polish** (#26522, #26523, #26516) – Auto Memory needs deterministic redaction, smarter retry logic, and better error surfacing for invalid patches.
- **Robust evaluation framework** (#24353) – expansion of behavioral evals into a structured component‑level suite to prevent regressions.
- **Browser agent resilience** (#22232, #22267) – demands for session‑takeover recovery, configuration overrides, and Wayland support.

## Developer Pain Points
- **Agent hangs and misreported success** (#21409, #22323) – subagents either hang indefinitely or report “success” after hitting turn limits, eroding user trust.
- **Tool under‑selection** (#21968, #24246) – Gemini rarely uses custom skills/sub‑agents and hits 400 errors when too many tools are available, forcing manual instructions.
- **Security gaps** (#26525, #27966 context) – secrets can leak via Auto Memory logging and case‑insensitive blocklist bypasses remain a concern.
- **Terminal integration fragility** (#25166, #24935) – shells remain stuck after completion, and external editors cause screen corruption.
- **Environment‑specific failures** (#21983 Wayland, #28092 auth, #28131 test flakiness) – inconsistent behaviour across platforms undermines reliability for Linux and new users.
- **Noise and overhead** (#23571 temp scripts, #22465 interactive prompts) – the model creates scattered tmp files and gets stuck at interactive prompts, requiring cleanup workarounds.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-06-25

## Today's Highlights

- **Release v1.0.65** landed yesterday, fixing a spurious filesystem permission prompt for commands with slash‑prefixed string arguments and improving session resume by persisting the working directory.
- The **model‑picker issue on resumed sessions** (#3680, #3596, #3913) continues to be the most discussed bug, with users unable to list or select models after resuming a session — a closed duplicate (#3913) was merged, but the root cause remains open.
- A new **argument‑hint validation error** (#3929) is blocking skill loading for users who use YAML list syntax, signaling a breaking change in the skills specification.

---

## Releases

### v1.0.65 (2026‑06‑24)
[Release Page](https://github.com/github/copilot-cli/releases/tag/v1.0.65)

- `/cd` now persists the working directory across session resumes and discovers custom agents in the new directory.
- Commands with slash‑prefixed string arguments (e.g. `--body "/azp run"`) no longer trigger spurious filesystem permission prompts.
- Fullscreen timeline stability improvements (details truncated in changelog).

---

## Hot Issues (Top 10 by Community Impact)

1. **[#700 – Provide a way to list all models](https://github.com/github/copilot-cli/issues/700)**  
   *Open since Dec 2025, 14 comments, 👍4*  
   Users want a `copilot --list-models` command to see all supported models and their multiplier info. Highly requested but not yet implemented; discussion suggests it would simplify model switching workflows.

2. **[#2643 – preToolUse silent command rewrite shows confirmation dialog](https://github.com/github/copilot-cli/issues/2643)**  
   *12 comments, 👍2*  
   Even when a plugin hook sets `permissionDecision: allow` with `updatedInput`, the CLI still prompts the user for every rewritten command. Plugin authors cannot achieve silent command execution, limiting automation.

3. **[#1632 – Support subfolders for skills](https://github.com/github/copilot-cli/issues/1632)**  
   *9 comments, 👍21*  
   Users with many skills are forced into a flat folder structure. The ability to organise skills into subfolders is the most upvoted feature request in the issue tracker.

4. **[#3596 – Error loading model list on resumed session](https://github.com/github/copilot-cli/issues/3596)**  
   *7 comments, 👍11*  
   Resuming a specific session triggers `✗ Error loading model list: Error: Not authenticated` while a fresh session works. Affects v1.0.56+. Strong community reaction due to workflow disruption.

5. **[#3832 – All models show as blocked after June 16 outage](https://github.com/github/copilot-cli/issues/3832)**  
   *6 comments, 👍13 – **Closed***  
   Post‑outage bug where model selection showed every model as "Blocked/Disabled". Fixed quickly but generated significant frustration and visibility.

6. **[#3881 – Incorrect quota subtraction (5% instead of 2%)](https://github.com/github/copilot-cli/issues/3881)**  
   *3 comments, 👍0*  
   A user reports that using a 6× multiplier model consumed 3% extra quota. Financial impact on paid accounts makes this a high‑priority issue for the team.

7. **[#3913 – Model selection empty when resuming a session](https://github.com/github/copilot-cli/issues/3913)**  
   *3 comments, 👍1 – **Closed***  
   Duplicate of #3596, closed as such. Highlights that multiple users are hitting the same resumed‑session authentication bug.

8. **[#3692 – Escape cancels queued prompt instead of focusing it](https://github.com/github/copilot-cli/issues/3692)**  
   *2 comments, 👍1*  
   Keyboard interaction friction: pressing Escape during a task discards the typed follow‑up rather than cancelling the current task and promoting the queued prompt.

9. **[#2419 – Configurable key bindings for fast model switching](https://github.com/github/copilot-cli/issues/2419)**  
   *2 comments, 👍5*  
   Users want F‑key hotkeys for `/model` commands to switch models faster without memorising IDs. Related to #1729 (general keybind config).

10. **[#3929 – argument-hint format validation issue](https://github.com/github/copilot-cli/issues/3929)**  
    *1 comment – **New***  
    Skills with `argument-hint: [regression directory]` (YAML list) now fail validation with `must be a string`. Breaking change in skill parsing – affects all users who followed VS Code docs example.

---

## Key PR Progress

Only two PRs were updated in the last 24 hours:

- **[#3928 – Add .gitignore and settings configuration](https://github.com/github/copilot-cli/pull/3928)**  
  *Open, created 2026-06-25*  
  A straightforward PR adding configuration files for development environment setup. No comments yet; likely a contributor’s first patch.

- **[#2587 – Add automated issue classification with GitHub Agentic Workflows](https://github.com/github/copilot-cli/pull/2587)**  
  *Closed, merged 2026-06-24*  
  Merged PR that introduces an AI‑powered workflow to automatically apply `area:` labels and add `triage` label when issues are opened. Expected to improve issue triage speed and consistency.

---

## Feature Request Trends

The following feature directions are most requested from recent issues:

- **Model management** – A dedicated list‑models command (#700), configurable key bindings for fast model switching (#2419), and seamless model change while editing a prompt (#3138).
- **Plugin / Skills improvements** – Subfolder support (#1632), silent command rewrite via hooks (#2643), and better validation feedback for skill definitions (#3929).
- **Input and keyboard UX** – Configurable key binds (#1729), cancel/focus behaviour for queued prompts (#3692), and consistent Ctrl+Enter vs Ctrl+Q enqueue semantics (#3760).
- **Enterprise & proxy** – Kerberos proxy support (#523), server‑managed env settings (#3909), and headless SDK proxy fixes (#2978).
- **Mobile app integration** – Support for `!` shell commands, file upload, and `/` slash commands in remote sessions (#3922, #3923, #3924).
- **Theming** – Fine‑grained per‑element theme customisation beyond dark/light (#2123).

---

## Developer Pain Points

Recurring frustrations and high‑frequency complaints:

- **Resumed session authentication** – Multiple reports (#3596, #3680, #3913) that model selection breaks after resuming a session, showing "Not authenticated". Fresh sessions work fine, indicating a state‑management bug.
- **Auto‑update ignored** (#2615) – The `autoUpdate: false` setting in config.json has been broken since v1.0.4, forcing users to manually block updates.
- **Keyboard inconsistencies** – Documentation says Ctrl+Enter enqueues, but on Windows it enters a newline; users must use Ctrl+Q instead (#3760). Escape discards pending input instead of promoting it (#3692).
- **Linux AppImage environment leak** (#3925) – Bundled `LD_LIBRARY_PATH` breaks spawned git HTTPS operations, blocking session creation on some Linux distros.
- **Markdown rendering bugs** – Two em‑dashes trigger strikethrough (#3920); multi‑line answer UI cuts off characters (#3921). These small UX issues erode trust in terminal output.
- **Quota calculation errors** (#3881) – Users concerned about overcharging for premium requests, especially when using high‑multiplier models.

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest – 2026-06-25

## Today’s Highlights
The repository saw no new releases in the past 24 hours, but two long‑running pull requests finally closed, bringing fixes for MCP propagation to subagents and a vim‑style keyboard navigation feature. A newly filed enhancement request about context compaction wasting ~20k tokens has drawn attention, while a months‑old looping bug remains open and unresolved.

---

## Releases
No new versions were published in the last 24 hours. The latest stable release (0.76) and the current web‑based CLI (0.19.2) remain unchanged.

---

## Hot Issues

1. **[#640 – [bug] Kimi CLI stuck in reading one file again and again and stuck in a loop](https://github.com/MoonshotAI/kimi-cli/issues/640)**  
   *Status: Open since January 2026 – last updated yesterday*  
   A user running v0.76 on Arch Linux with a custom Anthropic endpoint reports the CLI entering an infinite file‑reading loop. The thread has 14 comments and only 1 👍, suggesting a niche configuration issue, but the lack of a fix after five months is concerning for power users.

2. **[#2473 – [CLOSED] [bug] web bug](https://github.com/MoonshotAI/kimi-cli/issues/2473)**  
   *Created & closed yesterday*  
   User DCY501 reported that the `/web` command throws errors in the web‑based CLI. The issue was closed without comments, possibly a duplicate or user error. No resolution details are available.

3. **[#2469 – [CLOSED] [bug] `kimi web` starts MCP servers from CLI installation directory](https://github.com/MoonshotAI/kimi-cli/issues/2469)**  
   *Closed yesterday after two days open*  
   This bug broke workspace‑relative MCP tools because servers were launched from the global install path instead of the project folder. A fix was likely included in PR #1942 (see below). Community reaction: silent close, no comments.

4. **[#2472 – [OPEN] [enhancement] Context compaction reloads system prompt and project instructions, wasting ~20k tokens](https://github.com/MoonshotAI/kimi-cli/issues/2472)**  
   *Filed yesterday*  
   A well‑documented issue: after context compaction, the system prompt, `AGENTS.md`, skills, and environment context are reloaded, costing ~20k tokens each time. This is a significant inefficiency for long sessions. No comments yet, but likely to gain traction.

---

## Key PR Progress

1. **[#1942 – fix(mcp): propagate MCP configs to subagents and resume immediately](https://github.com/MoonshotAI/kimi-cli/pull/1942)**  
   *Closed yesterday after two months of development*  
   This PR addresses two issues: subagents (e.g., `explore`, `coder`, `plan`) previously received empty `mcp_configs`, making MCP tools unusable in sub‑tasks. It also fixes resume sessions that previously lost MCP configuration. A high‑impact fix for teams using MCP server extensions.

2. **[#1377 – feat: add vim-style j/k keyboard navigation for approval and question interactions](https://github.com/MoonshotAI/kimi-cli/pull/1377)**  
   *Closed yesterday after three months*  
   Adds `j`/`k` keyboard shortcuts for navigating through approval prompts and question lists. Aimed at terminal‑first users who prefer modal editing. Community discussion was minimal (0 comments), but the feature aligns with developer ergonomics.

---

## Feature Request Trends
Based on recent open issues, the community is asking for:

- **Token efficiency** – Avoid redundant token consumption during context compaction (Issue #2472).
- **Better MCP integration** – Users want predictable MCP server paths and subagent support (addressed in PR #1942).
- **Keyboard‑driven workflows** – The merged vim‑style navigation (#1377) signals demand for non‑mouse interactions.

No new feature requests appeared in the last 24 hours beyond the compaction token waste report.

---

## Developer Pain Points

- **Unresolved long‑standing bugs** – Issue #640 (file‑reading loop) has been open for five months with no fix, eroding trust in stability for custom model setups.
- **Configuration fragility** – MCP server path issues (#2469) and context compaction reloads (#2472) highlight that the CLI’s state management still has rough edges, particularly for power users.
- **Silent issue closures** – Bugs #2473 and #2469 were closed without public explanation, leaving the community guessing about root causes or workarounds.
- **Token cost awareness** – With models costing per token, the ~20k token waste in #2472 is a practical pain point for heavy users, though the issue is only a day old.

---

*Digest generated from data up to 2026-06-25 23:59 UTC. For the latest, visit [github.com/MoonshotAI/kimi-cli](https://github.com/MoonshotAI/kimi-cli).*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-25

## 1. Today's Highlights

OpenCode v1.17.10 dropped today with significant MCP improvements and a new `--mini` CLI mode—but the release has triggered a **major Windows stability crisis**. Dozens of users are reporting Bun v1.3.14 segmentation faults on Windows, with a critical escalation thread (#33742) drawing 26 upvotes in a single day. The community is rapidly submitting patches: Bun 1.4 canary for CI, Kitty keyboard disabling for IME crashes, and several TUI fixes are all in flight.

---

## 2. Releases

**v1.17.10** released today. Core improvements include:

- MCP server instructions added to session context (@Arcadi4)
- New Opencode-managed provider integration support
- MCP resource template listing and resource read tools
- `--mini` CLI mode for lightweight usage
- Bugfix: hid MCP resource template tools when conditions are unmet (details truncated)

**Note:** The release notes are incomplete—a bugfix line cuts off mid-sentence. This is likely the same build that introduced the Windows Bun segfault described in multiple issues.

---

## 3. Hot Issues (Top 10)

1. **[#33742 — OpenCode v1.17.10 crashes with Bun segmentation fault on Windows](https://github.com/anomalyco/opencode/issues/33742)** — *19 comments, 26 👍*  
   The most upvoted issue today. A clear regression: v1.17.9 was stable, v1.17.10 crashes on Windows with Bun v1.3.14 segfault. Core team is likely prioritizing this.

2. **[#8785 — Closed: Bun has crashed](https://github.com/anomalyco/opencode/issues/8785)** — *40 comments, 7 👍*  
   A historical issue from January being re-referenced—showing this crash pattern is not entirely new, though the scale is unprecedented.

3. **[#19081 — reasoning_content stripped from assistant messages on replay](https://github.com/anomalyco/opencode/issues/19081)** — *15 comments, 21 👍*  
   Persistent bug: thinking/reasoning tokens are silently removed from conversation history on subsequent turns, causing KV cache invalidation for local inference. High impact for local model users.

4. **[#33773 — v1.17.10 Bun segfault + downgrade breaks SQLite schema](https://github.com/anomalyco/opencode/issues/33773)** — *7 comments, 15 👍*  
   Double whammy: upgrading crashes, but downgrading to v1.17.9 corrupts the database due to a missing `replacement_seq` column. Users are trapped between two broken states.

5. **[#33767 — Segfault on Windows 11 in terminal](https://github.com/anomalyco/opencode/issues/33767)** — *12 comments, 9 👍*  
   User reports the crash doesn't respond to Ctrl+C—requires closing the entire terminal. Indicates a deeper process-level hang.

6. **[#33752 — Bun crashed when answering to agent](https://github.com/anomalyco/opencode/issues/33752)** — *12 comments, 0 👍*  
   Crash triggered specifically when an agent presents options. Suggests a UI/rendering interaction with Bun's runtime.

7. **[#33865 — Repeated crashes on Windows x64 with Bun v1.3.14 segfault](https://github.com/anomalyco/opencode/issues/33865)** — *2 comments, 4 👍*  
   Another Windows crash report, but specifically notes this happens in both PowerShell and Windows Terminal.

8. **[#33743 — Segfault with Excel files / stack overflow](https://github.com/anomalyco/opencode/issues/33743)** — *3 comments, 0 👍*  
   Crash triggered by working with `.xlsm`/`.xlsx` files. Stack trace shows infinite recursion—likely a Bun bug with large binary reads.

9. **[#33863 — OpenCode spawns dozens of zombie git.exe processes](https://github.com/anomalyco/opencode/issues/33863)** — *1 comment, 0 👍*  
   Resource leak: 7 OpenCode instances spawn hundreds of `git.exe` + `conhost.exe` processes, consuming 53% RAM and 50% disk. Persists even with one instance.

10. **[#33866 — Skills config uses old syntax (array vs object)](https://github.com/anomalyco/opencode/issues/33866)** — *2 comments, 0 👍*  
    The `skills` field changed from array to object syntax, but AI assistants still emit the old format. Config compatibility regression.

---

## 4. Key PR Progress (Top 10)

1. **[#33871 — fix: disable Kitty keyboard on Windows to fix IME crash](https://github.com/anomalyco/opencode/pull/33871)** — *New*  
   Critical fix for Chinese/Japanese IME crashes. Kitty keyboard intercepts IME keystrokes as CSI sequences; combined with non-UTF-8 code pages, this causes TUI panic. High priority for i18n users.

2. **[#33822 — fix(ci): use Bun 1.4 canary for Windows](https://github.com/anomalyco/opencode/pull/33822)** — *Open*  
   "Bun 1.3.14 segfaults lots on Windows. The Rust rewrite looks more stable now lol." — Direct response to today's crash crisis. Moving to Bun 1.4 canary for CI.

3. **[#33842 — Revert "deps: update OpenTUI to 0.4.2"](https://github.com/anomalyco/opencode/pull/33842)** — *Closed*  
   Reverting an OpenTUI dependency update. Likely related to TUI instability on Windows.

4. **[#33870 — fix(tui): ignore empty slash commands](https://github.com/anomalyco/opencode/pull/33870)** — *New*  
   Standalone `/` caused TUI to panic if command list contained an empty entry. Edge case but could crash on certain configurations.

5. **[#33854 — fix(config): error on missing {env:VAR} in config templates](https://github.com/anomalyco/opencode/pull/33854)** — *Open*  
   Currently, undefined `{env:VAR}` silently resolves to `""`, allowing misconfigured MCP servers to start with empty auth headers. Missing env vars will now throw `ConfigInvalidError`.

6. **[#33868 — fix(tui): search files under reference aliases](https://github.com/anomalyco/opencode/pull/33868)** — *Closed*  
   `@docs` autocomplete worked, but `@docs/...` file search failed because the alias match wasn't used for nested paths. Fixed autocomplete for referenced directories.

7. **[#33861 — fix(mcp): render prompt commands with real args](https://github.com/anomalyco/opencode/pull/33861)** — *Closed*  
   MCP prompt commands displayed placeholder `$1`, `$2` when listing. Servers with Linux-required args could fail. Now renders actual argument values.

8. **[#33807 — fix(opencode): re-download skills when declared version changes](https://github.com/anomalyco/opencode/pull/33807)** — *Closed*  
   Skills were not re-downloaded when their declared version changed in config, causing stale skill usage. Now detects version mismatch and refreshes.

9. **[#33850 — fix(app): use project server for prompt drafts](https://github.com/anomalyco/opencode/pull/33850)** — *Closed*  
   Draft prompts used the active server instead of the selected project's server. Fixed to associate drafts with the correct provider.

10. **[#32767 — fix(tui): restore ESC interrupt for delegated subagent sessions](https://github.com/anomalyco/opencode/pull/32767)** — *Open*  
    Regression: ESC key no longer interrupted delegated subagent sessions. Restores a key UX capability for multi-agent workflows.

---

## 5. Feature Request Trends

- **MCP Ecosystem Growth**: Multiple PRs and issues focus on MCP server configuration, resource templates, and validation. The community is pushing for better MCP tooling, including config validation for `type` fields (#33845) and variable substitution (#33853).
- **Better Windows Support**: The overwhelming theme. Users want Windows to be a first-class platform without segfaults, zombie processes, or IME crashes. The Bun 1.4 migration (#33822) is a direct response.
- **Session Persistence & Multi-Instance**: Requests for session state preservation across project switches (#33328), draft model selection persistence (#33858), and handling of local-only sessions in Desktop sidebar (#27539).
- **Skill/Plugin Lifecycle**: Users want version-aware skill downloads (#33807), proper config syntax compatibility (#33866), and automatic detection of newly installed skills (opencode.jsonc auto-update).
- **Performance & Resource Management**: Reports of excessive git zombie processes (#33863) and high GPU/CPU consumption (#33872) indicate a need for better resource governance.

---

## 6. Developer Pain Points

- **Windows Instability (Critical)**: The v1.17.10 release has effectively broken OpenCode on Windows. Multiple segfaults, IME crashes, zombie processes, and a no-win upgrade/downgrade scenario (#33773). This is the #1 developer frustration today.
- **Silent Config Failures**: `{env:VAR}` resolving to `""` (#33853) and missing config validation for MCP types (#33845) cause hard-to-debug runtime failures. Developers want early, clear error messages.
- **Reasoning Content Stripping**: Think/Reasoing tokens being silently removed from conversation history (#19081) breaks local inference workflows. Requires full replay with no output, leading to confused debugging.
- **Segfault Recovery UX**: Multiple users report that Ctrl+C doesn't kill the crashed process (#33767). The application leaves terminal sessions in an unrecoverable state.
- **Skills Config Migration Pain**: The `skills` field changed from array to object, but AI assistants still emit old syntax (#33866). The configuration format evolution lacks backward compatibility tooling.
- **Desktop vs CLI Inconsistency**: Sessions created from CLI with `version='local'` are invisible in the Desktop sidebar (#27539). Developers using both interfaces lose session continuity.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-06-25

## Today’s Highlights
The community is seeing a surge in reliability‑focused work: multiple fixes for streaming hangs and provider‑level retry logic landed today, alongside a new experimental orchestrator daemon. A long‑standing connection reliability issue with OpenAI Codex continues to attract heavy discussion (70 comments), while the team shipped important infrastructure like extension loading benchmarks and a fix for scrollback corruption in the TUI.

## Releases
*No new releases in the last 24 hours.*

## Hot Issues (10 Noteworthy)

1. **[#4945 – openai-codex Connection Reliability Issues](https://github.com/earendil-works/pi/issues/4945)**  
   *Open, 70 comments, 30 👍*  
   Major ongoing discussion about GPT‑5.5 leaving the TUI stuck on `Working…` with no error. The only workaround is Escape. High community impact.

2. **[#5363 – Add amazon-bedrock-mantle provider for OpenAI‑compatible models](https://github.com/earendil-works/pi/issues/5363)**  
   *Open, 14 comments, 4 👍*  
   Request to add a new provider for Bedrock Mantle’s OpenAI‑compatible API, which is incompatible with the existing Converse‑based provider.

3. **[#5291 – Sessions hang on “working” when used with Anthropic subscription](https://github.com/earendil-works/pi/issues/5291)**  
   *Closed, 7 comments, 2 👍*  
   Reproducible hang pattern with Anthropic Enterprise accounts; intermittent resume helped only sometimes.

4. **[#5671 – ~/.pi and cwd/.pi overlap](https://github.com/earendil-works/pi/issues/5671)**  
   *Closed, 6 comments, 5 👍*  
   Configuration UX problem where global and project `.pi` directories can conflict in `$HOME`. Sparked discussion about renaming or better isolation.

5. **[#5595 – openai-completions maxTokens not passing through](https://github.com/earendil-works/pi/issues/5595)**  
   *Open, 5 comments, 2 👍*  
   Reasoning models (e.g., DeepSeek v4pro) exhaust output tokens regardless of user settings; token limit not forwarded correctly.

6. **[#5810 – RPC: expose session entries and tree](https://github.com/earendil-works/pi/issues/5810)**  
   *Closed, 5 comments*  
   Read‑only RPC commands `get_entries` and `get_tree` to help headless integrations drive pi.

7. **[#6050 – TUI full redraw clears terminal scrollback during active rendering](https://github.com/earendil-works/pi/issues/6050)**  
   *Closed, 5 comments*  
   Frequent full redraws cause the terminal scrollbar to jump to beginning; root cause in core renderer.

8. **[#6019 – OpenAI Responses mid‑stream retryable error is not retried](https://github.com/earendil-works/pi/issues/6019)**  
   *Closed, 4 comments*  
   Provider error after stream start is not retried even when OpenAI explicitly says it is safe to retry.

9. **[#6057 – Add reasoning token counts to Usage](https://github.com/earendil-works/pi/issues/6057)**  
   *Closed, 2 comments*  
   Request to surface `usage.output_tokens_details.reasoning_tokens` from OpenAI, Anthropic, etc. – currently dropped.

10. **[#6002 – SessionManager.open() silently truncates non‑session files](https://github.com/earendil-works/pi/issues/6002)**  
    *Open, 2 comments*  
    Critical data‑loss bug: pointing `--session` at an NDJSON log file truncates it to 133 bytes with no warning.

## Key PR Progress (9 Important)

1. **[#5832 – fix(ai): surface provider HTTP error body instead of opaque SDK message](https://github.com/earendil-works/pi/pull/5832)**  
   *Open* – Makes proxy/gateway errors visible (e.g., actual 403 body) instead of generic `UnknownError`.

2. **[#6063 – Extension stats](https://github.com/earendil-works/pi/pull/6063)**  
   *Closed* – Adds per‑extension load timing (closes #6062) and fixes OSC garbage printed after benchmarking.

3. **[#6064 – feat(experimental): pi orchestrator](https://github.com/earendil-works/pi/pull/6064)**  
   *Open* – New `@earendil-works/pi-orchestrator` package that runs a local daemon for lifecycle management of pi instances over UNIX socket IPC.

4. **[#6056 – feat(subagent): simplify agent configs, add default agent, use minimax model](https://github.com/earendil-works/pi/pull/6056)**  
   *Closed* – Switches example agents to MiniMax‑M2.7, simplifies output formats, adds a default general‑purpose agent config.

5. **[#6055 – feat(subagent): simplify agent configs… (duplicate of #6056)](https://github.com/earendil-works/pi/pull/6055)**  
   *Closed* – Same content as #6056; likely a stale branch.

6. **[#6054 – feat(agent,coding-agent): add runParallelAgentTasks + parallel batching guideline](https://github.com/earendil-works/pi/pull/6054)**  
   *Closed* – New utility for running independent agent loops concurrently; adds system prompt guideline to batch tool calls.

7. **[#5509 – feat: Add Amazon Bedrock Mantle OpenAI Responses provider](https://github.com/earendil-works/pi/pull/5509)**  
   *Open* – Adds support for Bedrock Mantle models (GPT‑5.5/5.4) via OpenAI‑compatible API, modelled after Azure.

8. **[#6051 – fix(ai): recover from hung streams and retry unmodeled Bedrock errors](https://github.com/earendil-works/pi/pull/6051)**  
   *Closed* – Implements idle timeout (240s) and connect timeout; retries silent Bedrock errors to prevent infinite hangs.

9. **[#6048 – fix(coding-agent): show resources before messages when resuming session](https://github.com/earendil-works/pi/pull/6048)**  
   *Closed* – Moves loaded resources (Context, Skills, etc.) to a container above chat history so they appear at top.

## Feature Request Trends
- **Provider diversity**: Bedrock Mantle is the most demanded new provider; also requests for custom fetch support in OpenAI adapter and more granular token tracking (reasoning tokens).
- **TUI / UX enhancements**: Inline skill selector on `/`, shortcut for named sessions, dedicated terminal scrollback stability, and single‑executable binary packaging.
- **Agent infrastructure**: Concurrent agent sub‑task loops (`runParallelAgentTasks`), durable human‑in‑the‑loop for tool calls, and exposing session data via RPC for headless use.
- **Extensibility**: Extension loading timers, BMP file support in the `read` tool, and custom editor component ordering fixes.

## Developer Pain Points
- **Stream reliability dominates**: OpenAI Codex and Anthropic subscriptions both cause sessions to hang on `Working…` without errors; mid‑stream errors are not retried.
- **Silent data loss**: `SessionManager.open()` truncates non‑session files with no backup; `maxTokens` settings are ignored for reasoning models.
- **TUI instability**: Full redraws clear scrollback, overly long lines crash the terminal, and input history is lost after `/resume` when custom editors are used.
- **Configuration friction**: Overlap between global and project `.pi` directories, and reasoning level settings not respected by local models (e.g., Gemma 4B).

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest – 2026-06-25

---

## Today’s Highlights

Three releases were cut today, including the stable **v0.19.2** which ships a new remote LSP status route and various fixes. The community is actively shaping the voice dictation experience (configurable keyterms, desktop integration) and pushing for better session management APIs. A streaming timeout bug (#401) remains a high-priority concern for users with large inputs.

---

## Releases

- **v0.19.2-nightly.20260625** (Build `b2f11b735`) – Includes a core fix for web_fetch JSON fallback.  
  [Release notes](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.2-nightly.20260625.b2f11b735)  

- **v0.19.2** – Stable release containing the remote LSP status route (`feat(serve): Add remote LSP status route`).  
  [Release notes](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.2)  

- **v0.19.2-preview.0** – Preview of v0.19.2 features, including the same LSP route.  
  [Release notes](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.2-preview.0)

---

## Hot Issues

*Note: only 8 issues were updated in the last 24h; all are listed below.*

| # | Title | Status | Priority | Why It Matters |
|---|-------|--------|----------|----------------|
| [#401](https://github.com/QwenLM/qwen-code/issues/401) | Streaming setup timeout after 6s | OPEN | P1 | Core usability: users with long inputs hit hard timeout; 8 comments requesting workaround or fix. |
| [#5789](https://github.com/QwenLM/qwen-code/issues/5789) | Enable built-in status line by default for new users | CLOSED | P3 | Improves onboarding – context (model, git, directory) shown immediately. |
| [#5863](https://github.com/QwenLM/qwen-code/issues/5863) | Enrich GET /session/:id/status with live turn-phase, active tools, pending permissions | OPEN | P2 | Follow-up to session status work; would give external consumers rich telemetry. |
| [#5861](https://github.com/QwenLM/qwen-code/issues/5861) | Context compression request should use stream=true to avoid gateway timeout | OPEN | P1 | Critical bug for large sessions; non-streaming compression blocks server. |
| [#5782](https://github.com/QwenLM/qwen-code/issues/5782) | WebFetch should reject URLs containing userinfo | CLOSED | P3 | Security: prevents accidental credential leakage via URLs. |
| [#5816](https://github.com/QwenLM/qwen-code/issues/5816) | Voice dictation: support user-configurable keyterms file | CLOSED | P2 | Users need to extend hardcoded ASR bias list for project-specific terms. |
| [#5855](https://github.com/QwenLM/qwen-code/issues/5855) | feat(serve): query a single session's status by id | OPEN | – | Adds missing fundamental API endpoint for session monitoring. |
| [#5742](https://github.com/QwenLM/qwen-code/issues/5742) | Improve voice package distribution for mirror registries | CLOSED | P2 | Install from private registries currently missing native audio capture dependency. |

---

## Key PR Progress

| # | Title | Status | Impact |
|---|-------|--------|--------|
| [#5847](https://github.com/QwenLM/qwen-code/pull/5847) | Add runtime context injection for per-turn system-reminders | OPEN | Enables dynamic operator identity & rules injected into every turn – valuable for team deployments. |
| [#5738](https://github.com/QwenLM/qwen-code/pull/5738) | Default to virtualized terminal history | OPEN | New users automatically get in-app scrollable history; opt-out available. |
| [#5778](https://github.com/QwenLM/qwen-code/pull/5778) | Add /model --vision for vision model fallback | OPEN | Allows text-only models to delegate image vision tasks to a configured vision model. |
| [#5809](https://github.com/QwenLM/qwen-code/pull/5809) | Split serve server routes | OPEN | Refactor to reduce monolithic daemon server – improves maintainability. |
| [#5828](https://github.com/QwenLM/qwen-code/pull/5828) | Add bundled extension creator skill | OPEN | Guides users through scaffolding new extensions – lowers barrier for contributing. |
| [#5864](https://github.com/QwenLM/qwen-code/pull/5864) | Show duration on finished thinking summary | OPEN | Polishes web shell: keeps elapsed time visible after thinking completes. |
| [#5844](https://github.com/QwenLM/qwen-code/pull/5844) | Make self-paced /loop lean on monitor/background-task notifications | OPEN | Smarter loop scheduling – avoids polling when background tasks are active. |
| [#5856](https://github.com/QwenLM/qwen-code/pull/5856) | Voice dictation in the desktop app | OPEN | Brings /voice to the desktop UI, matching CLI and web shell. |
| [#5030](https://github.com/QwenLM/qwen-code/pull/5030) | Resume interrupted turn without synthetic "continue" | OPEN | Long-awaited: true turn continuation across crashes without injecting fake messages. |
| [#5661](https://github.com/QwenLM/qwen-code/pull/5661) | Partition tool display by type (collapse read/search, show mutations individually) | OPEN | Improves debugging TUI readability – read tools collapsed, mutation tools expanded. |

---

## Feature Request Trends

- **Voice dictation expansion** – Three issues/PRs (#5816, #5742, #5856) focus on enhancing voice input: configurable keyterms, better package distribution, and desktop app support.  
- **Session management APIs** – #5855, #5863, #5857 (PR) all target richer session query endpoints for the daemon, indicating a push toward external integration and monitoring.  
- **Improved UX defaults** – #5789 (status line), #5738 (virtualized history), and #5848 (collapse preview count) show a focus on reducing cognitive load for new users.  
- **Security hardening** – #5782 (reject userinfo in URLs), #5550 (secret disclosure mandate), and #5829 (unsafe source slugs) signal growing attention to credential safety.

---

## Developer Pain Points

1. **Streaming timeouts** – Issue #401 (“Streaming setup timeout after 6s”) and #5861 (context compression non-streaming) both block heavy workloads. Users are seeking better configurability and streaming defaults.  
2. **Voice installation friction** – #5742 highlights that users on mirror/private registries miss native voice capture packages; a silent failure that frustrates adopters.  
3. **Session state visibility** – Developers integrating with the daemon need richer session status (active tools, pending permissions) – single boolean `hasActivePrompt` is insufficient (#5863).  
4. **Long-running turn resumption** – PR #5030 tackles the pain of crashes mid-turn forcing synthetic “continue” messages – a long-standing gap for reliable agent workflows.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

Here is the **DeepSeek TUI Community Digest** for **2026-06-25**, based on the provided GitHub data for the `Hmbown/CodeWhale` project (which serves as the current codebase for the DeepSeek TUI).

---

### 1. Today's Highlights

Today’s digest centers on the v0.8.65 release stabilization and several significant regression fixes. Key updates include a critical fix for Windows environment variable inheritance, a resolution to the problematic MCP server duplicate lifecycle, and progress on the **Fleet** loadout automation. The community is actively testing the new **ACP (Agent Communication Protocol) adapter** and requesting deeper provider configuration controls, particularly for context window sizing and custom model endpoints.

### 2. Releases

**No new releases in the last 24 hours.**

---

### 3. Hot Issues

**1. v0.8.65: Fleet model classes, loadout auto, and semantic route roles**
- **#3205** (Open)
- **Why it matters:** This is the core feature request for v0.8.65, proposing a unified model/loadout selector for TUI, CLI, exec, and Fleet workers. The goal is to introduce a **Fleet loadout auto** mode that dynamically resolves complete compute loads for a given role, rather than just picking a model string.
- **Community Reaction:** High engagement with 10 comments and 7 👍. The author (Hmbown) is driving the implementation, indicating this is a top-priority feature.
- **Link:** [Issue #3205](https://github.com/Hmbown/CodeWhale/issues/3205)

**2. v0.8.65: MCP duplicate server instance lifecycle and doctor coverage**
- **#3461** (Closed)
- **Why it matters:** This resolves a serious bug where CodeWhale spawned **two MCP server processes** from a single `mcp.json` entry, wasting ~4MB RAM and creating a shared stdio pipe that could cause silent failures.
- **Community Reaction:** Quick turnaround from creation (Jun 23) to closure (Jun 24) with 8 comments. The fix is critical for reliability in multi-server workflows.
- **Link:** [Issue #3461](https://github.com/Hmbown/CodeWhale/issues/3461)

**3. v0.8.65: Multi-model compatibility, provider docs, and automatic Fleet loadout selection**
- **#2300** (Open)
- **Why it matters:** This serves as the user-facing acceptance fixture for the multi-model support and Fleet routing features in v0.8.65. It highlights confusion between `provider = vllm` and `provider = openai` for local endpoints.
- **Community Reaction:** Moderate traction with 7 comments, suggesting the documentation gap is still a pain point for users.
- **Link:** [Issue #2300](https://github.com/Hmbown/CodeWhale/issues/2300)

**4. Extend ACP support to expose provider and model selection**
- **#3546** (Closed)
- **Why it matters:** This enhancement request (and its corresponding PR) allows external ACP clients (like Paseo) to discover and switch providers/models. This unlocks deeper integration capabilities for the CodeWhale TUI.
- **Community Reaction:** Positive feedback, closed quickly with 3 comments after the PR was merged.
- **Link:** [Issue #3546](https://github.com/Hmbown/CodeWhale/issues/3546)

**5. Bug: Windows user environment variables not inherited by codewhale shell**
- **#3572** (Closed)
- **Why it matters:** A critical OS-specific bug where `exec_shell` commands did not inherit user-level environment variables set in Windows System Properties, breaking build tools and SDK detection.
- **Community Reaction:** Rapid closure (created and closed same day) with a fix in PR #3578. The 1 comment indicates a direct, resolved issue.
- **Link:** [Issue #3572](https://github.com/Hmbown/CodeWhale/issues/3572)

**6. [Enhancement] 希望在 providers 的配置里可以自定义上下文的大小**
- **#3545** (Closed)
- **Why it matters:** A user (Chinese community) requests the ability to configure custom context window sizes for providers. They note that Qwen models support 1M context, but the default is locked to 128k.
- **Community Reaction:** The request was quickly addressed by the context window override feature in PRs #3573 and #3574.
- **Link:** [Issue #3545](https://github.com/Hmbown/CodeWhale/issues/3545)

**7. [Rust Version] change rust-toolchain to stable**
- **#3570** (Open)
- **Why it matters:** A proposal to change the Rust toolchain from a pinned version (`1.88`) to `stable`. This would allow developers to use the latest Rust features but could break compatibility for users on older toolchains.
- **Community Reaction:** Low engagement (1 comment), but it touches on a fundamental build infrastructure decision.
- **Link:** [Issue #3570](https://github.com/Hmbown/CodeWhale/issues/3570)

**8. [bug] plan and agent mode mixed up YET AGAIN**
- **#3568** (Open)
- **Why it matters:** A frustrating regression where the AI fails to respect the switch between `plan` and `agent` modes. The user provides a chat export showing the AI attempting multiple file modification methods while in `plan` mode.
- **Community Reaction:** The user is clearly frustrated ("YET AGAIN"), and the issue has only 1 comment, suggesting it's a known pain point that has resurfaced.
- **Link:** [Issue #3568](https://github.com/Hmbown/CodeWhale/issues/3568)

**9. v0.8.65: Split config modules around provider/model/catalog boundaries**
- **#3311** (Closed)
- **Why it matters:** A significant refactoring effort to split large config modules into cleaner provider/model/offering boundaries. This is foundational for all the provider routing work in v0.8.65.
- **Community Reaction:** 5 comments and closed, indicating a successful internal refactor.
- **Link:** [Issue #3311](https://github.com/Hmbown/CodeWhale/issues/3311)

**10. v0.8.65: Custom provider endpoints, models, and auth within provider-scoped routing**
- **#1519** (Closed)
- **Why it matters:** This high-level feature request covers support for custom endpoints, model IDs, and authentication within the new provider routing architecture. It acts as a master ticket for several sub-issues.
- **Community Reaction:** 4 comments, closed as the underlying work (issues #2608, #3311, #3084, #3384) has landed.
- **Link:** [Issue #1519](https://github.com/Hmbown/CodeWhale/issues/1519)

---

### 4. Key PR Progress

**1. Feat(memory): wire moraine-mcp as recall tool source**
- **#3575** (Open)
- **What it does:** Connects CodeWhale to a new MCP server (`moraine mcp`) for memory recall tools (search sessions, file attention). It also introduces a `moraine_fallback` config gate to deprecate the old push/inject system.
- **Why it matters:** Lays the foundation for a more robust, externalized memory system.
- **Link:** [PR #3575](https://github.com/Hmbown/CodeWhale/pull/3575)

**2. Feat: expose provider and model selection over the ACP stdio adapter**
- **#3576** (Closed)
- **What it does:** Implements the feature requested in Issue #3546. It allows external ACP clients to read and switch the active provider/model.
- **Why it matters:** Enables deeper integrations with other AI tools like Paseo.
- **Link:** [PR #3576](https://github.com/Hmbown/CodeWhale/pull/3576)

**3. [codex] Address context window review feedback**
- **#3574** (Closed)
- **What it does:** A follow-up fix for the context window override feature, ensuring `max_tokens` stays positive when a small context window is configured, and restoring the active context override on provider switch.
- **Why it matters:** Prevents subtle bugs during provider switching and edge cases in token budget calculations.
- **Link:** [PR #3574](https://github.com/Hmbown/CodeWhale/pull/3574)

**4. [codex] Add provider context window overrides**
- **#3573** (Closed)
- **What it does:** Adds a `context_window` override to provider configurations, threading it through TUI startup, model switches, CLI exec, compaction, and request output caps.
- **Why it matters:** Directly addresses the user request for customizable context sizes (Issue #3545) and is a major configuration feature.
- **Link:** [PR #3573](https://github.com/Hmbown/CodeWhale/pull/3573)

**5. Fix(tools): make js_execution fetch honor proxy env vars**
- **#3577** (Closed)
- **What it does:** Fixes a bug where `js_execution`'s `fetch()` function ignored `HTTP_PROXY`/`HTTPS_PROXY` environment variables, causing timeouts behind a proxy/VPN.
- **Why it matters:** Critical for developers working in corporate or proxy-restricted network environments.
- **Link:** [PR #3577](https://github.com/Hmbown/CodeWhale/pull/3577)

**6. [codex] preserve Windows SDK env roots for shell**
- **#3578** (Closed)
- **What it does:** Fixes Issue #3572 by ensuring Windows SDK/toolchain path variables are properly inherited from the User/Machine environment registry before shell launch.
- **Why it matters:** Resolves a major pain point for Windows developers using CodeWhale for build or SDK-dependent tasks.
- **Link:** [PR #3578](https://github.com/Hmbown/CodeWhale/pull/3578)

**7. Feat(tui): show configured ask rules**
- **#3569** (Closed)
- **What it does:** Adds a read-only `/config ask-rules` surface in the TUI displaying configured ask-only permission rules, the path to `permissions.toml`, and rule count.
- **Why it matters:** Improves discoverability and transparency of permission settings.
- **Link:** [PR #3569](https://github.com/Hmbown/CodeWhale/pull/3569)

**8. Clarify condensed tool transcript rows**
- **#3566** (Closed)
- **What it does:** Improves the readability of compact tool transcript rows, keeping tool identity visible and suppressing non-essential defaults.
- **Why it matters:** Enhances the AI tool use logging and debugging experience.
- **Link:** [PR #3566](https://github.com/Hmbown/CodeWhale/pull/3566)

**9. [codex] Rename DeepSeek blue consumers to whale accent**
- **#3197** (Closed)
- **What it does:** A branding and code consistency PR that renames `DEEPSEEK_BLUE` to `WHALE_ACCENT_PRIMARY` across the TUI, keeping backwards-compatible aliases.
- **Why it matters:** Reflects the ongoing rebranding from DeepSeek to CodeWhale.
- **Link:** [PR #3197](https://github.com/Hmbown/CodeWhale/pull/3197)

**10. [codex] enforce main-backed release tags**
- **#3526** (Closed)
- **What it does:** A security/reliability PR that prevents release artifacts from shipping from commits that have not been merged to `main`.
- **Why it matters:** Ensures release provenance and prevents accidental releases from feature branches.
- **Link:** [PR #3526](https://github.com/Hmbown/CodeWhale/pull/3526)

---

### 5. Feature Request Trends

Based on the data, the most-requested feature directions are:

1.  **Customizable Configuration & Overrides:** A strong push for user-controlled settings, specifically **context window sizes** (Issue #3545) and **custom provider endpoints/models/auth** (Issue #1519). The community wants granular control over provider behavior.
2.  **External Integration & Protocol Expansion:** Enabling deeper connections with external tools is a key theme. This includes exposing provider/model selection via **ACP** (Issue #3546) and wiring in external memory systems like **Moraine MCP** (PR #3575).
3.  **Cross-Platform & Environment Parity:** Ensuring feature parity and bug-free operation across different OS environments (Windows vs. Linux/Mac) is a clear trend, highlighted by the Windows environment variable fix (Issue #3572).
4.  **Infrastructure & Developer Tooling:** A push for keeping the Rust toolchain on `stable` (Issue #3570) and improving documentation clarity around provider routing (Issue #2300) indicates a community focus on reducing maintenance overhead and lowering the barrier for new contributors.

### 6. Developer Pain Points

- **Regressions in Core Modes:** The recurring issue of the **Plan/Agent mode confusion** (Issue #3568) is a significant source of frustration, as it undermines trust in the AI's ability to follow user intent. This appears to be a persistent problem that has resurfaced despite previous fixes.
- **Complex Configuration & Documentation:** The high number of open issues around provider routing and multi-model support (Issues #2300, #3205) suggests the current configuration paradigm is daunting for end-users. The request for clearer documentation on `vllm` vs `openai` providers is a concrete example.
- **Windows Support Gaps:** The bug regarding environment variable inheritance (Issue #3572) highlights ongoing challenges with Windows as a first-class platform. Developers on Windows face specific environment management headaches that are less common on Unix-like systems.
- **Version Tracking & Release Management:** The need to bump stale version references in documentation (PR #3579) and enforce `main`-backed release tags (PR #3526) suggests some developer friction with the current release and versioning workflow.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*