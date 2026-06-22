# AI CLI Tools Community Digest 2026-06-23

> Generated: 2026-06-22 17:18 UTC | Tools covered: 9

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

# AI CLI Tools Ecosystem: Cross-Tool Comparison Report
**Date:** 2026-06-23

---

## 1. Ecosystem Overview

The AI CLI tools landscape is entering a **stabilization phase** characterized by intense focus on reliability, security, and enterprise readiness rather than raw feature expansion. Across all seven tools surveyed, three themes dominate: **permission granularity and safety**, **MCP (Model Context Protocol) integration maturity**, and **cross-platform stability**—particularly Windows support. The ecosystem is bifurcating between tools that emphasize autonomous agent behavior (Claude Code, Gemini CLI, CodeWhale) and those optimizing for developer workflow integration (OpenCode, Qwen Code). A notable signal is the emergence of **cost and usage transparency** as a first-class concern, with rate-limit regressions and billing bugs causing significant community backlash. No major releases landed today, but all projects show active PR throughput, suggesting a "fix and polish" cycle before the next feature wave.

---

## 2. Activity Comparison (Last 24 Hours)

| Tool | Hot Issues | Key PRs | Releases (24h) | Community Engagement Signal |
|------|-----------|---------|----------------|---------------------------|
| **Claude Code** | 10 (168 👍 top) | 3 | 0 | Highest single-issue engagement; compound permission parsing trending |
| **OpenAI Codex** | 10 (231 👍 top) | 10 | 2 (alpha) | Most urgent bug (#28879 rate-limit regression); highest overall reaction count |
| **Gemini CLI** | 10 (8 👍 top) | 10 | 0 | Broadest scope of fixes; 48 open issues + 38 PRs updated; highest total activity volume |
| **GitHub Copilot CLI** | 10 (11 👍 top) | 1 (spam) | 0 | Lowest engagement; only 1 actionable PR (non-functional); community concerns about stagnation |
| **Kimi Code CLI** | 6 | 2 | **1 (v1.48.0)** | Smallest but focused community; MCP and provider bugs dominate |
| **OpenCode** | 10 | 10 | 0 | Security-focused burst; 3 closed issues by same reporter; strong PR velocity |
| **Qwen Code** | 10 | 10 | 1 (nightly) | Systematic validation sweep; highest PR-to-issue correlation |
| **CodeWhale** | 10 | 10 | **1 (v0.8.64)** | Security hardening release; rebrand signals strategic shift |

**Key insight:** Gemini CLI leads in raw activity (38 PRs+48 issues), but OpenAI Codex has the most urgent community pain (#28879). Copilot CLI shows stagnation risk.

---

## 3. Shared Feature Directions

### 3.1 Permission & Safety Systems
*(Across all tools except Copilot CLI)*

- **Compound command parsing**: Claude Code (#16561) and OpenCode (#33077) both need granular permission matching for `&&`, `|`, and `;` chains
- **Common theme**: Users want *incremental approval* (per-command) not *bulk approval* (per-pipeline)
- **Pre-execution guardrails**: CodeWhale (#3364) and Gemini CLI (#22672) propose read-before-edit and destructive-command prevention

### 3.2 MCP Integration Maturity
*(Most active in Kimi, Claude Code, OpenCode)*

| Requirement | Tools Affected |
|------------|---------------|
| `envFile` for secret loading before env expansion | Claude Code (#28942), OpenCode |
| OAuth DCR persistence (avoid client_id orphaned) | Claude Code (#59460) |
| Null/malformed parameters handling | Kimi (#2465), OpenCode (#28472, #33160) |
| Auto-discovery of deleted/duplicate MCP servers | Kimi (#2457), Claude Code |
| Streaming HTTP Accept header compliance | OpenCode (#25650) |

### 3.3 Windows Platform Parity
*(All non-Microsoft tools affected)*

- Drive-letter case sensitivity in path comparison: Claude Code (#62288)
- Auth credential storage failures: Claude Code (#69706)
- ARM64 crash under load: Copilot CLI (#3687)
- Process tree TOCTOU race: OpenCode (#33071)
- Desktop app freezes: OpenAI Codex (#20214)

### 3.4 Cross-Session Memory & Context
*(Dominant feature request in multiple communities)*

- **Kimi Code** (#1283): Persistent AI-managed memory system
- **Claude Code**: Org-level skills linked to Git repos (#28729)
- **Gemini CLI**: Auto Memory retry loops cause noisy feedback (#26522)
- **CodeWhale**: Automatic compaction with carried-forward summaries (#3363)
- **Common pain**: Current implementations either lose context or produce infinite loops

### 3.5 Cost Transparency & Budget Control
*(Explosive concern in OpenAI Codex, emerging elsewhere)*

- **Codex** (#28879): 10-20x rate-limit cost jump on June 16; Plus plan users hit 5-hour budget in 2-3 prompts
- **Copilot CLI** (#3886, #3881): AI credits consumed on restart (174 per restart); 5% vs 2% quota discrepancy
- **Claude Code** (#68773): Auto-recharge loop costing $661; phantom "extra usage" on Max plan (#45390)
- **Pi**: OpenRouter cost display accuracy (#5950)
- **Common theme**: Users demand *predictable* billing; silent regressions erode trust faster than feature gaps

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | Copilot CLI | Kimi Code | OpenCode | Qwen Code | CodeWhale |
|-----------|-------------|-------------|------------|-------------|-----------|----------|-----------|-----------|
| **Primary user** | Power CLI developers | ChatGPT Plus users | GCP/enterprise teams | GitHub ecosystem | Chinese market devs | Multi-provider OSS | Self-hosted/AI-CI | Open source agents |
| **Architecture** | Agent-first (Opus) | Rust-based lightweight | Sub-agent orchestration | .NET/WPF desktop | Node.js + soul engine | SolidJS TUI | SDK-decoupled | Rust + WebAssembly |
| **Differentiator** | Permission granularity leader | Rate-limit cost battleground | Agent reliability breadth | Slowing; legacy | Fast shipping (daily releases) | Security/TOCTOU fixes | Env validation rigor | Rebrand + security |
| **Community growth rate** | Medium | High (but frustrated) | High | Low (stagnating) | Small but loyal | Stable OSS | Growing | Rebrand boosting |

### Key Differentiation Highlights

- **Claude Code** owns the **permission granularity** conversation (168 👍 on #16561) and Windows portability pain
- **OpenAI Codex** is the **cost/limitation battleground**—its #28879 with 231 reactions signals a trust crisis that other tools should monitor
- **Gemini CLI** has the **broadest engineering velocity** (38 PRs + 48 issues updated) but lower per-issue engagement; their subagent architecture complexity generates more bugs per feature
- **Copilot CLI** is showing **stagnation risk**: 1 PR (non-functional spam) in 24h, multiple unaddressed bugs (#3854 regression since "a few versions ago")
- **Kimi Code** ships fastest (daily releases) but has **smallest community surface area**
- **OpenCode** is the **security/TOCTOU leader** with systematic vulnerability reporting by a single prolific reporter
- **Qwen Code** distinguishes via **validation rigor**: their env-input validation sweep (6 PRs in 24h) is unmatched
- **CodeWhale** is **rebranding for privacy focus**—security hardening tracker (#3368) and WeCom integration suggest China-first strategy

---

## 5. Community Momentum & Maturity

### High-Maturity Communities (Stable, predictable cadence)

| Tool | Maturity Indicators |
|------|-------------------|
| **Claude Code** | Well-defined issue templates; maintainers respond consistently; 168 reactions on top issue shows engaged user base |
| **OpenCode** | Systematic security reporting pipeline; closed issues per structured process; PRs linked to issues |

### High-Momentum Communities (Rapid iteration, high activity)

| Tool | Momentum Indicators |
|------|-------------------|
| **Gemini CLI** | 48 open issues + 38 PRs updated in 24h; maintainers actively merging; broadest scope of concurrent work |
| **Qwen Code** | 6 PRs targeting same category (env validation) shows coordinated sweep; release failure triggered immediate hotfix (#5676) |
| **CodeWhale (rebranded)** | Rebrand + security release + WeCom integration = strategic expansion; highest PR velocity relative to issue count |

### Stagnation Risk

| Tool | Warning Signs |
|------|--------------|
| **Copilot CLI** | 1 PR (non-functional) in 24h; regression #3854 unaddressed; no maintainer activity visible in digest; community sentiment flat |

### Controlled Growth

| Tool | Indicators |
|------|-----------|
| **Kimi Code** | Small issue count (6); daily releases maintain momentum; community focuses on narrow set of MCP bugs |
| **Pi** | 43 open issues but low engagement per issue (top issue 64 comments); PR throughput low (7 PRs) |

---

## 6. Trend Signals for Developers

### Signal 1: The "Rate-Limit Trust Crisis"
**Evidence**: OpenAI Codex #28879 (231 reactions)  
**Implication**: Users are hypersensitive to hidden cost changes. Any tool that prices per-token or per-request must provide *transparent, real-time* usage feedback. The 2026 developer expects **budget controls, not budget surprises**. Tools that proactively expose cost accumulation (Pi PR #5950, Qwen Code PR #29423) will win trust.

### Signal 2: Windows as the "Compatibility Litmus Test"
**Evidence**: 6+ Windows-specific bugs across Claude Code, Copilot CLI, OpenCode, OpenAI Codex  
**Implication**: Cross-platform parity is no longer optional. The Linux-first era is ending—developers on Windows ARM64, Intel macOS, and Wayland Linux expect equal experience. Tools ignoring Windows (or treating it as second-class) will hit adoption ceilings in enterprise and education markets.

### Signal 3: MCP Standardization Squeeze
**Evidence**: MCP-related issues across 5 tools (Claude Code, Kimi, OpenCode, Gemini CLI, Qwen Code)  
**Implication**: MCP is becoming the universal protocol, but implementations are fragmenting on edge cases (null params, Accept headers, OAuth DCR). The *first tool* that achieves seamless MCP compliance will capture the plugin/extension ecosystem. Current state: everyone has bugs.

### Signal 4: "Developer Agent Safety" Becomes a Product Differentiator
**Evidence**: CodeWhale #3364 (read-before-edit), Gemini #22672 (discourage destructive behavior), Claude Code #16561 (granular permission), OpenCode #33077 (destructive commands protection)  
**Implication**: The market is moving from "can the agent do it?" to "can I trust the agent to do it safely?" Tools that implement **provenance tracking, pre-execution guardrails, and transparent approval flows** will differentiate. The era of black-box agent autonomy is ending.

### Signal 5: Cross-Session Memory Is Table Stakes
**Evidence**: Kimi Code #1283 (36 👍), Claude Code #28729, Gemini Code #26522, CodeWhale #3363  
**Implication**: Every major tool is receiving requests for persistent, AI-managed memory. But current implementations cause infinite retry loops, duplicate messages, and corruption. The *killer feature* for 2026-2027 is not more memory—it's *reliable* memory that doesn't leak, loop, or lie.

### Signal 6: Chinese AI Tools Are Scaling Differently
**Evidence**: Kimi (daily releases, narrow scope), CodeWhale (rebrand + WeCom integration)  
**Implication**: Chinese-market tools prioritize **speed to iteration** and **enterprise-legacy integration** (WeCom). They are more willing to rebrand and pivot. Non-Chinese tool developers should watch their UX patterns (especially around permission and compaction) as they scale.

---

## Summary for Technical Decision-Makers

| If you prioritize... | Consider... | But watch for... |
|---------------------|-------------|-----------------|
| Permission granularity | Claude Code | Windows compatibility |
| Agent autonomy breadth | Gemini CLI | Subagent reliability issues |
| Cost stability | Qwen Code | Smaller community size |
| Security validation | OpenCode | CPU spikes (#33399) |
| Rapid iteration | Kimi Code | Narrow provider support |
| Enterprise integration | Copilot CLI | Stagnation risk |
| Privacy-first | CodeWhale | Rebrand transition friction |
| Multi-provider flexibility | Pi | Connection reliability (#4945) |

**Bottom line**: The AI CLI tools ecosystem is healthy but fractured. No single tool dominates across all dimensions. The most critical unaddressed need across all tools is **reliable, predictable billing transparency**—and the tool that solves it first (without breaking existing features) will capture significant mindshare.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
**Data snapshot:** github.com/anthropics/skills | 2026-06-23

---

## 1. Top Skills Ranking

The most-discussed Skills (by PR activity) reveal a community focused on *infrastructure quality* and *enterprise integration*:

**#1 – `skill-creator` eval fix** ([PR #1298](https://github.com/anthropics/skills/pull/1298))
- **What it does:** Repairs the `run_eval.py` pipeline so the description-optimization loop reports accurate recall, rather than always 0%. Fixes Windows stream reading, trigger detection, and parallel workers.
- **Discussion:** The topmost PR by position; tied directly to Issue #556 (12 comments, 7 👍). Multiple independent reproductions confirm the bug. The fix is broad — it installs the eval artifact as a real skill and rewrites Windows subprocess handling.
- **Status:** Open (created 2026-06-10, updated 2026-06-22). High activity suggests imminent merge.

**#2 – Document Typography** ([PR #514](https://github.com/anthropics/skills/pull/514))
- **What it does:** Prevents orphan word wrap, widow paragraphs, and numbering misalignment in AI-generated documents. Addresses a universal pain point.
- **Discussion:** Strong support; typographic polish is a low-cost, high-visibility quality improvement.
- **Status:** Open (created 2026-03-04).

**#3 – ODT (OpenDocument) Skill** ([PR #486](https://github.com/anthropics/skills/pull/486))
- **What it does:** Creates, fills, reads, and converts `.odt`/`.ods` files. Covers LibreOffice and ISO-standard document workflows.
- **Discussion:** The community is actively pushing for open-source document format support — a clear gap in the ecosystem.
- **Status:** Open (created 2026-03-01, updated 2026-04-14).

**#4 – Frontend Design Skill** ([PR #210](https://github.com/anthropics/skills/pull/210))
- **What it does:** Revises the frontend-design skill for clarity and actionability. Ensures every instruction is executable within a single conversation.
- **Discussion:** Focuses on *skill quality* itself — making instructions specific enough to steer behavior without ambiguity.
- **Status:** Open (created 2026-01-05).

**#5 – Testing Patterns Skill** ([PR #723](https://github.com/anthropics/skills/pull/723))
- **What it does:** Covers the full testing stack: philosophy (Testing Trophy model), unit testing (AAA pattern), React component testing (Testing Library), E2E (Playwright), and property-based testing.
- **Discussion:** Broad scope points to demand for structured testing guidance inside Claude Code.
- **Status:** Open (created 2026-03-22).

**#6 – ServiceNow Platform Skill** ([PR #568](https://github.com/anthropics/skills/pull/568))
- **What it does:** Comprehensive ServiceNow assistant covering ITSM, ITOM, ITAM, SAM Pro, FSM, HRSD, CSM, SPM, Vulnerability Response, and IntegrationHub.
- **Discussion:** Enterprise users are investing heavily in ServiceNow automation via Claude.
- **Status:** Open (created 2026-03-08).

**#7 – AURELION Skill Suite** ([PR #444](https://github.com/anthropics/skills/pull/444))
- **What it does:** Four skills — kernel (structured thinking templates), advisor (cognitive workflows), agent (autonomous execution), memory (persistent context). Built on a 5-floor cognitive framework.
- **Discussion:** Represents the most ambitious structured-knowledge contribution; combines meta-cognition with professional knowledge management.
- **Status:** Open (created 2026-02-21).

**#8 – SAP-RPT-1-OSS Predictor** ([PR #181](https://github.com/anthropics/skills/pull/181))
- **What it does:** Wraps SAP's open-source tabular foundation model for predictive analytics on SAP business data.
- **Discussion:** Early indication of demand for industry-specific ML model integration via Skills.
- **Status:** Open (created 2025-12-28).

---

## 2. Community Demand Trends

The Issues data reveals three concentrated demand vectors:

**a) Organizational Skill Sharing & Governance** (Issue #228: 14 comments, 7 👍)
Users want to share `.skill` files across teams without manual download/upload workflows. A shared skill library or direct sharing link is the most-upvoted feature request. Closely related is the **trust boundary concern** (Issue #492: 9 comments) — community skills distributed under the `anthropic/` namespace could impersonate official skills, raising security risks.

**b) Cross-Platform Compatibility** (Issues #556, #1061, #1169)
Windows users face persistent subprocess and encoding bugs in `skill-creator` scripts. Three separate issues (and multiple PRs) address PATHEXT handling, `cp1252` encoding, and `select()` on pipes. This is the community's most *reproduced* pain point.

**c) Agent Governance & Safety Patterns** (Issue #412: 6 comments)
A formal proposal for governance skills — policy enforcement, threat detection, trust scoring, audit trails. No equivalent Skill currently exists in the collection, making this a clear greenfield opportunity.

**Secondary trends:**
- **MCP protocol integration** (Issue #16) — exposing Skills as Model Context Protocol endpoints
- **AWS Bedrock compatibility** (Issue #29) — Skills currently are Claude Code-specific
- **Duplicate skill conflicts** (Issue #189) — when installing both `document-skills` and `example-skills` plugins

---

## 3. High-Potential Pending Skills

These open PRs receive active comment threads and are likely to land soon:

| PR | Skill | Why It May Merge Soon |
|---|---|---|
| [#514](https://github.com/anthropics/skills/pull/514) | Document Typography | Low complexity, universal benefit, no breaking changes |
| [#723](https://github.com/anthropics/skills/pull/723) | Testing Patterns | Broad community value, well-structured scope |
| [#568](https://github.com/anthropics/skills/pull/568) | ServiceNow Platform | Strong enterprise signal, single-vendor focus |
| [#444](https://github.com/anthropics/skills/pull/444) | AURELION Suite | Four skills in one PR, but modular design |
| [#1099](https://github.com/anthropics/skills/pull/1099) | Windows fix (run_eval) | Directly fixes a blocking bug for Windows users |
| [#154](https://github.com/anthropics/skills/pull/154) | shodh-memory | Persistent memory across conversations — high novelty |

The **Windows compatibility fixes** (PRs #1099, #1050, #1298) are the most likely to merge *fastest* given the volume of related issues.

---

## 4. Skills Ecosystem Insight

The community's most concentrated demand is **reliable cross-platform Skill infrastructure** — fixing the `skill-creator` evaluation pipeline for Windows users — combined with a clear appetite for **enterprise-grade Skills** (ServiceNow, SAP, governance, testing patterns) that extend Claude Code beyond developer tooling into organizational workflows.

---

# Claude Code Community Digest — 2026-06-23

## Today's Highlights
The community remains focused on **permission granularity and Windows stability** — the most-upvoted issue asks for compound Bash command parsing to avoid unnecessary approval prompts, while a growing number of Windows-specific bugs (auth failures, session hiding, RPC errors) signal rough edges for that platform. A billing auto-recharge loop costing $661 and persistent HTTP 429 on auto-mode classifier are causing significant user frustration. No new releases landed in the last 24 hours.

## Releases
*No new versions published in the last 24 hours.*

## Hot Issues
1. **[Feature: Parse compound Bash commands and match each component against permissions](https://github.com/anthropics/claude-code/issues/16561)** (44 comments, 168 👍) — The highest-voted open issue: compound operators (`&&`, `|`, `;`, `||`) cause the entire string to be evaluated as a single unit, blocking commands whose parts would individually be allowed. Community strongly supports granular permission matching.

2. **[1M context incorrectly requires extra usage on Max plan](https://github.com/anthropics/claude-code/issues/45390)** (35 comments, 27 👍) — Max plan users report that selecting Opus 4.6 with 1M context triggers an “Extra usage required” error despite the plan covering it. The issue is still open after two months.

3. **[Link a source control repo as the source for organization skills](https://github.com/anthropics/claude-code/issues/28729)** (28 comments, 54 👍) — Admins want to point org-level skills at a Git repo for version-controlled, team-synced definitions. Currently only manual definition is possible.

4. **[Persistent HTTP 429 on auto-mode classifier (xml_s1), account-side config](https://github.com/anthropics/claude-code/issues/60438)** (15 comments, 2 👍) — A long-running rate-limit bug that persists across versions; the auto-mode classifier apparently triggers excessive API calls. Closed today but unresolved.

5. **[API Error: 401 Invalid authentication credentials on Windows](https://github.com/anthropics/claude-code/issues/69706)** (15 comments, 9 👍) — Recent Windows users hit persistent 401 errors. Not yet confirmed whether it’s a token storage or credential manager issue.

6. **[Windows: VS Code extension hides sessions when drive-letter case in cwd differs](https://github.com/anthropics/claude-code/issues/62288)** (7 comments, 2 👍) — Sessions recorded with `C:\...` are invisible when opened from `c:\...` due to case-sensitive comparison. A classic Windows portability bug.

7. **[Support `envFile` in `.mcp.json` for loading secrets before env var expansion](https://github.com/anthropics/claude-code/issues/28942)** (6 comments, 16 👍) — Teams using HTTP-based MCP servers need a declarative way to load secrets (e.g., API keys) without plumbing them into the shell environment. Highly requested.

8. **[MCP OAuth: Claude Code re-runs DCR on every authenticate, orphaning client_id](https://github.com/anthropics/claude-code/issues/59460)** (4 comments, 4 👍) — Each `authenticate` flow spawns a fresh Dynamic Client Registration, leaking client IDs and invalidating existing refresh tokens. Breaks persistent auth for self-hosted MCP servers.

9. **[Opening a chat re-appends duplicate mode/custom-title record, bumping mtime and reordering it](https://github.com/anthropics/claude-code/issues/69939)** (3 comments) — Every time you open a session, identical JSONL bookkeeping lines are appended, pushing the session to the top of “Recent chats” even without changes. Causes confusion in chat history.

10. **[Session resume pollutes conversation files with duplicate system messages](https://github.com/anthropics/claude-code/issues/69013)** (3 comments) — Using `claude --continue` or `--resume` adds duplicate `mode` and `permission-mod` records, making the conversation appear empty. Users lose visibility into past context.

## Key PR Progress
1. **[docs: fix stale marketplace name in plugin-dev README](https://github.com/anthropics/claude-code/pull/70074)** — Corrects `claude-code-marketplace` → `claude-code-plugins` in three locations, aligning documentation with the actual marketplace config.

2. **[docs(plugin-dev): update marketplace install docs](https://github.com/anthropics/claude-code/pull/70066)** — Updates install instructions to use the official marketplace name, replaces `cc --plugin-dir` with `claude --plugin-dir`, and clarifies contribution guidelines.

3. **[fix: print error message before silent exit in edit-issue-labels.sh](https://github.com/anthropics/claude-code/pull/69916)** — Fixes a silent failure when the script is invoked without `--add-label` or `--remove-label`. Affects the Claude Issue Triage workflow.

## Feature Request Trends
The most-requested feature directions this cycle include:
- **Permission granularity** — compound Bash command parsing (#16561), pre-tool hook `defer` handling (#64389)
- **Org/team control** — linking source control repos to org skills (#28729), session-scoped compact hints (#70083)
- **MCP improvements** — `envFile` support for secret loading (#28942), OAuth DCR persistence (#59460)
- **Cross-model flexibility** — ability to use cheaper third-party models alongside Claude for low-level tasks (#55163)
- **User experience** — idle-queue batch-flush for multi-question responses (#55152), token counter persistence after input focus (#55127)

## Developer Pain Points
Recurring frustrations evident from the issue tracker:
- **Windows friction** — multiple authentication failures (#69706), drive-letter case sensitivity in VS Code (#62288), RPC message size errors in Cowork (#70099)
- **Session/data corruption** — duplicate JSONL records on open/resume (#69939, #69013), subagents reading `node_modules` instead of source (#55146), `.env` file corruption during tool execution (#55158)
- **Billing & limits** — auto-recharge loops (#68773), phantom “extra usage” on Max plan (#45390), silent output token caps (#55147)
- **MCP integration headaches** — OAuth DCR re-registration (#59460), env var expansion ordering (#28942)
- **CLI & editor quirks** — `/issue` command returning "Unknown command" (#55155), effort level overrides on model first-run (#55141), Monitor tool spawning infinite loops (#55151)

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-06-23

## Today’s Highlights
A critical rate-limit cost regression (#28879) is draining Plus budgets 10–20x faster than expected, drawing 231 reactions and 116 comments. On the performance front, a newly reported SQLite log write amplification issue (#28224) could wear out SSDs within months. Meanwhile, community demand for IDE-integrated diff/approval (#2998) and inline image previews in the TUI (#29451) continues to grow, with several related PRs now in flight.

## Releases
Two alpha versions of the Rust client were published in the last 24 hours:
- **rust-v0.142.0-alpha.10** – Release 0.142.0-alpha.10
- **rust-v0.142.0-alpha.11** – Release 0.142.0-alpha.11

No detailed changelog was provided; these are incremental alpha releases ahead of the stable 0.142.0.

## Hot Issues (10 of 26)
1. **[#28879] Rate-limit cost per token jumped ~10-20x since June 16**  
   *Critical bug* – ChatGPT Plus users on `gpt-5.5` are exhausting their 5‑hour budget in 2–3 prompts instead of 20+. Logs show `limit-% consumed per token` spiked. 231 👍, 116 comments.  
   [GitHub](https://github.com/openai/codex/issues/28879)

2. **[#28224] Codex SQLite feedback logs can write ~640 TB/year and rapidly consume SSD endurance**  
   *Performance bug* – Continuous high-volume writes to `~/.codex/logs_2.sqlite` and related WAL files. 214 👍, 25 comments.  
   [GitHub](https://github.com/openai/codex/issues/28224)

3. **[#2998] IDE-integrated diff / approval**  
   *Enhancement* – Users want the CLI’s red/green diff and approval flow natively inside VS Code and other IDEs. 198 👍, 62 comments.  
   [GitHub](https://github.com/openai/codex/issues/2998)

4. **[#29000] CLI 0.141.0 crashes with SIGTRAP on Intel macOS**  
   *Bug* – `codex-cli` immediately crashes on `x86_64` macOS under `gpt-5.5`. 11 👍, 9 comments.  
   [GitHub](https://github.com/openai/codex/issues/29000)

5. **[#20214] Codex App frequently freezes/stutters on Windows 11 Pro**  
   *Bug* – Despite sufficient RAM/CPU, the desktop app becomes unresponsive. 31 👍, 18 comments.  
   [GitHub](https://github.com/openai/codex/issues/20214)

6. **[#14559] Error running remote compact task: stream disconnected before completion**  
   *Bug* – Remote task fails with a disconnection error under `gpt-5.4`. 14 👍, 23 comments.  
   [GitHub](https://github.com/openai/codex/issues/14559)

7. **[#4242] Use proxy environment variables across Codex HTTP clients**  
   *Enhancement* – Enterprise users need consistent `HTTPS_PROXY` / `HTTP_PROXY` / `ALL_PROXY` support in all outgoing requests. 80 👍, 5 comments.  
   [GitHub](https://github.com/openai/codex/issues/4242)

8. **[#29451] Show image artifacts inline in the Codex TUI**  
   *Enhancement* – Request for terminal-native previews of images returned by `view_image` and `imagen`. 29 👍, 0 comments (new).  
   [GitHub](https://github.com/openai/codex/issues/29451)

9. **[#2932] Add support to view diffs in the VSCode diff viewer (CLOSED)**  
   *Enhancement* – Closed but heavily upvoted (28 👍); the demand for richer IDE diff experiences remains high.  
   [GitHub](https://github.com/openai/codex/issues/2932)

10. **[#28500] Unified mentions menu cannot navigate or select items with arrow keys**  
    *Bug* – The new `@` mentions menu in CLI v0.140.0 becomes unusable. 4 👍, 2 comments.  
    [GitHub](https://github.com/openai/codex/issues/28500)

## Key PR Progress (10 of 50)
1. **[#28918] Make selected plugin roots URI-native**  
   Changes plugin root paths to `file://` URIs for consistent cross-platform behavior, fixing issues on Windows and Linux.  
   [GitHub](https://github.com/openai/codex/pull/28918)

2. **[#29155] Expose service tier and reasoning effort in OTEL**  
   Adds `service_tier` and `model_reasoning_effort` to OpenTelemetry logs, per an NVIDIA integration request.  
   [GitHub](https://github.com/openai/codex/pull/29155)

3. **[#28991] Allow ChatGPT accounts without email**  
   Fixes PAT login for service accounts that lack an email field in their metadata, unblocking headless/CI usage.  
   [GitHub](https://github.com/openai/codex/pull/28991)

4. **[#26703] TUI Plugin Sharing 3 – render remote plugin catalog sections**  
   Third PR in the chain bringing remote plugin discovery and browsing to the terminal UI.  
   [GitHub](https://github.com/openai/codex/pull/26703)

5. **[#26707] PAC 2 – Add shared auth system proxy contract**  
   Moves Codex’s authentication and startup HTTP clients through a common proxy-aware boundary, paving the way for system proxy resolution.  
   [GitHub](https://github.com/openai/codex/pull/26707)

6. **[#29423] Configure rollout budget reminder thresholds**  
   Replaces the fixed `reminder_interval_tokens` with a user-configurable list of remaining-token thresholds (e.g., `[65536, 32768, ...]`).  
   [GitHub](https://github.com/openai/codex/pull/29423)

7. **[#29424] Report remote sandbox denials semantically**  
   After sandbox enforcement moved to the exec server, this PR returns meaningful error messages (e.g., “network access denied”) instead of opaque failures.  
   [GitHub](https://github.com/openai/codex/pull/29424)

8. **[#27946] Use input items for Responses Lite tools**  
   Migrates Responses Lite calls to `additional_tools` and a developer system item, improving API consistency and tool namespace isolation.  
   [GitHub](https://github.com/openai/codex/pull/27946)

9. **[#29457] Filter noisy targets from persistent logs**  
   Reduces the rate of SQLite log writes by filtering high-volume TRACE entries and duplicated OpenTelemetry events – directly addresses the log bloat seen in #28224.  
   [GitHub](https://github.com/openai/codex/pull/29457)

10. **[#29003] Cache codex_apps tools in memory**  
    Moves tool reads from disk to an in-memory snapshot, reducing latency and I/O when `list_all_tools()` is called frequently.  
    [GitHub](https://github.com/openai/codex/pull/29003)

## Feature Request Trends
- **IDE / Editor Integration**: The most-requested direction continues to be native diff viewing and approval flows inside VS Code, JetBrains, and code-server. Issues #2998 and #2932 (closed) show strong community desire for an experience equivalent to the CLI’s approval flow.
- **Terminal UI Enhancements**: Users want richer inline previews (images, diffs, scrollable mentions) without leaving the terminal. #29451 (image previews) and #28500 (mentions navigation) are recent examples.
- **Enterprise Networking**: Consistent proxy support (#4242) and managed network sandbox context (#29456) are high priorities for corporate deployments.
- **Budget & Usage Transparency**: Configurable token-reminder thresholds (#29423) and better OTEL reporting (#29155) reflect a need for finer-grained cost control and observability.
- **Remote Plugin Discovery**: The multi‑PR chain (#26703) indicates ongoing work to let users browse and install community plugins directly from the TUI.

## Developer Pain Points
- **Rate-Limit Instability**: The sudden cost-per-token jump (#28879) has eroded trust in budget predictions; users report that their Plus plan is effectively unusable.
- **Excessive Disk I/O**: SQLite logging (#28224) and session log bloat (#24948) cause SSD wear and slowdowns. Multiple PRs (#29457, #29432) are now tackling the problem.
- **Cross-Platform Crashes**: macOS Intel crashes (SIGTRAP – #29000) and Windows freezes (#20214) remain open, affecting a wide user base.
- **Remote / Sandbox Reliability**: Stream disconnections (#14559) and sandbox timeouts (#29253) disrupt long-running tasks, especially on remote SSH setups.
- **UI Responsiveness**: Desktop app flickering (#22860), silent processing stops (#29453), and launch crashes due to corrupted drafts (#29299) degrade the user experience.

---

*Digest generated from `github.com/openai/codex` data as of 2026-06-22 23:59 UTC.*

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-06-23

## Today's Highlights

A major burst of activity across agent reliability, memory system hygiene, and core tooling fixes dominated the past 24 hours. While no new releases landed, the project saw 48 open issues and 38 pull requests updated, with maintainer focus on subagent recovery bugs, file corruption bugs in `write_file`, and foundational infrastructure for MCP elicitation. The most visible community pain point remains the generalist agent hang (#21409), which has accumulated 8 👍 reactions and remains open since March.

## Releases

No new releases in the last 24 hours.

## Hot Issues

1. **#21409 – Generalist agent hangs** [🔗](https://github.com/google-gemini/gemini-cli/issues/21409)
   *The most upvoted open bug (8 👍)*. Users report the CLI hangs "forever" when deferring to the generalist agent for simple tasks like folder creation. Workaround exists (disable sub-agent delegation), but the community is frustrated by the lack of resolution since March 6.

2. **#22323 – Subagent MAX_TURNS reported as GOAL success** [🔗](https://github.com/google-gemini/gemini-cli/issues/22323)
   A critical misreporting bug: the `codebase_investigator` subagent returns `status: "success"` and `Termination Reason: "GOAL"` even when it hit the maximum turn limit before doing any analysis. This masks real failures and degrades trust in agent output.

3. **#25166 – Shell command gets stuck on "Waiting input"** [🔗](https://github.com/google-gemini/gemini-cli/issues/25166)
   +3 👍. After executing trivial CLI commands, Gemini hangs with the shell appearing active despite completion. This disrupts workflow and is flagged as `priority/p1, effort/medium`.

4. **#26525 – Add deterministic redaction for Auto Memory** [🔗](https://github.com/google-gemini/gemini-cli/issues/26525)
   A security concern: Auto Memory sends transcripts to the extraction model *before* redaction occurs. Secrets are only redacted after content is already in model context. The issue calls for deterministic redaction before transmission.

5. **#26522 – Auto Memory retries low-signal sessions indefinitely** [🔗](https://github.com/google-gemini/gemini-cli/issues/26522)
   Auto Memory only marks sessions as processed after a successful `read_file`. Low-signal sessions are never marked as processed and keep being surfaced, creating a noisy feedback loop.

6. **#28092 – Google account not authorized for Gemini** [🔗](https://github.com/google-gemini/gemini-cli/issues/28092)
   A fresh issue from yesterday. Users report being blocked by an authorization wall when running `/auth` with their Google account. Could be a configuration or entitlement issue.

7. **#21983 – Browser subagent fails on Wayland** [🔗](https://github.com/google-gemini/gemini-cli/issues/21983)
   Linux Wayland users see browser agent crashes with a `GOAL` termination reason, but the actual behavior suggests failure. A persistent platform compatibility issue.

8. **#22672 – Agent should discourage destructive behavior** [🔗](https://github.com/google-gemini/gemini-cli/issues/22672)
   Community request (+1 👍). Users report the model uses `git reset`, `--force`, or destructive DB commands when safer alternatives exist. The ask is for the agent to understand consequences before executing.

9. **#22267 – Browser Agent ignores settings.json overrides** [🔗](https://github.com/google-gemini/gemini-cli/issues/22267)
   Configuration overrides (e.g., `maxTurns`) provided in `settings.json` are completely ignored by the Browser Agent. The `AgentRegistry` reads them correctly, but they never propagate to the agent's runtime.

10. **#24246 – 400 error with > 128 tools** [🔗](https://github.com/google-gemini/gemini-cli/issues/24246)
    When the number of available tools exceeds 128, the CLI encounters a 400 error. The community expects smarter tool scoping rather than a hard limit.

## Key PR Progress

1. **#28000 – Fix Jupyter Notebook and JSON corruption in write_file** [🔗](https://github.com/google-gemini/gemini-cli/pull/28000)
   Resolves a critical bug where `write_file` silently corrupts `.ipynb` and JSON files. This fix prevents environments (Colab, JupyterLab) from discarding changes. Size M, awaiting issue linking.

2. **#28053 – Defensive path resolution for `@`-prefixed file paths** [🔗](https://github.com/google-gemini/gemini-cli/pull/28053)
   A comprehensive fix for a production bug where tools fail with "File not found" when the model passes paths like `@policies/new-policies.txt`. Also fixes macOS tests. Size XL.

3. **#28096 – Drop late tool calls after SIGINT cancellation** [🔗](https://github.com/google-gemini/gemini-cli/pull/28096)
   Closes #28091. Prevents delayed provider tool-call chunks from executing locally after a user cancels with SIGINT. An important UX reliability fix.

4. **#28094 – Deep-merge user and workspace settings in A2A server** [🔗](https://github.com/google-gemini/gemini-cli/pull/28094)
   Fixes a shallow merge bug where nested config sections (`tools`, `telemetry`) in workspace settings were overwritten by user settings. Priority P2, size M.

5. **#28093 – Buffer chat compression telemetry until SDK initialization** [🔗](https://github.com/google-gemini/gemini-cli/pull/28093)
   Fixes telemetry emission that bypassed the buffering wrapper, potentially sending records before the OpenTelemetry SDK was ready. Priority P2, enterprise area.

6. **#27915 – Fix trust dialog disclosing wrong hook shape** [🔗](https://github.com/google-gemini/gemini-cli/pull/27915)
   Closes #27901. The workspace-trust dialog showed the *inverse* of the hooks that actually run, meaning a `SessionStart` hook could execute on a single "Trust folder" click without being displayed. Security-critical fix.

7. **#28068 – Guard message inspectors against empty parts arrays** [🔗](https://github.com/google-gemini/gemini-cli/pull/28068)
   `isFunctionCall()` and `isFunctionResponse()` returned `true` for messages with empty `parts` arrays due to JavaScript's vacuous `[].every()`. Could cause misclassification of model messages. Size M.

8. **#28089 – Implement MCP elicitation (form + url) capability** [🔗](https://github.com/google-gemini/gemini-cli/pull/28089)
   Implements the 2025-11-25 MCP specification for elicitation, supporting `form` and `url` modes. This is a foundational feature for MCP integration. Size L, addresses #28074.

9. **#27910 – Bound web search tool latency with 120s timeout** [🔗](https://github.com/google-gemini/gemini-cli/pull/27910)
   Fixes #27890. Adds a 120-second local timeout around `google_web_search` calls, preventing indefinite hangs and returning a clear tool error for agent recovery. Priority P1.

10. **#27914 – Don't offer to resume a session that wasn't saved** [🔗](https://github.com/google-gemini/gemini-cli/pull/27914)
    Fixes #27277. When disk space runs out (`ENOSPC`), the chat recorder disables itself but the exit summary still printed a resume prompt. This fix suppresses the misleading message. Priority P2.

## Feature Request Trends

1. **AST-aware file operations** – Multiple issues (#22745, #22746, #22747) call for integrating AST-aware CLI tools (tilth, glyph, AST grep) for smarter codebase mapping, file reads, and search. The goal is reduced token usage and fewer misaligned reads.

2. **Subagent trajectory visibility** – Issue #22598 requests that subagent trajectories be viewable via `/chat share`. The community wants better debuggability and evaluation of subagent behavior.

3. **Agent self-awareness and configuration** – Issue #21432 asks the CLI to understand its own mechanics (hotkeys, CLI flags) well enough to serve as its own expert guide. Related to #22267 (settings.json override propagation).

4. **Browser agent resilience** – Multiple requests (#22232, #21983) want automatic session takeover, lock recovery, and Wayland support for the browser subagent. The current "fail-fast" strategy on locked profiles is too restrictive.

5. **Safe agent behavior** – Issue #22672 (discourage destructive git/DB operations) and #23571 (avoid random temp scripts) reflect a desire for more cautious agent operations, especially on production-like resources.

## Developer Pain Points

- **Agent hangs and false success**: The most upvoted issues (#21409, #22323) involve agents either hanging indefinitely or reporting success despite hitting limits. This erodes trust in autonomous operation.
- **Subagent permission and control**: Users report subagents running without permission (#22093) and browser agents ignoring configured settings (#22267). The configuration system feels unreliable.
- **File corruption and data loss**: The `write_file` corruption of JSON/IPYNB (#28000) and the session file deletion race conditions (#27905, #27912) represent concrete data-loss risks.
- **Authentication and authorization**: The new #28092 (Google account not authorized) and existing GCP project ID validation issues (#27916) suggest ongoing friction in the auth flow.
- **Path resolution quirks**: The `@`-prefixed path bug (#28053) and the `\n` escape behavior issue (#22466) highlight fragility in how the model and CLI communicate about file system paths.
- **Terminal and rendering problems**: Issues like high performance on terminal resize (#21924) and corruption after exiting external editors (#24935) affect the day-to-day UX for CLI power users.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

**GitHub Copilot CLI Community Digest – 2026-06-23**

---

## Today’s Highlights

No new releases landed in the last 24 hours, but the community is buzzing about several critical bugs and feature requests. A regression in the `@` file-reference syntax (#3854) and an authentication error when resuming sessions (#3596) are top frustrations, while calls for better timer/elapsed-time display (#3278, #3055, #3111) and i18n support (#3883) signal growing demand for UX improvements.

---

## Releases

None in the last 24 hours.

---

## Hot Issues

1. **#3854 – `@` syntax for file reference broken**  
   *Area: input-keyboard* – Users can no longer autocomplete file paths with `@`. The feature stopped working a few versions ago, affecting prompt efficiency. Silent regression with no prior announcement.

2. **#3596 – Error loading model list: Not authenticated**  
   *Area: authentication, sessions, models* – Resuming a specific session triggers an authentication failure when listing models via `/model`. Requires starting a new session. 11 👍, 6 comments – a clear pain point for session-heavy workflows.

3. **#3687 – copilot.exe fatal-aborts under load (BEX64) on Windows ARM64**  
   *Area: sessions, platform-windows* – Hard crashes (0xc0000409) when multiple sessions start simultaneously under memory pressure. Affects versions 1.0.57–1.0.60. Urgent stability concern for Windows ARM64 users.

4. **#3861 – Sandbox docs present capabilities that don’t work**  
   *Area: permissions, networking* – Per-host filtering (`allowedHosts`/`blockedHosts`) and cross-platform isolation are documented but non-functional. Enterprise users can’t rely on promised security controls.

5. **#3886 – Restarting Copilot uses AI credits (174 per restart)**  
   *Area: sessions, models* – `/restart`, `/resume`, and `/update` consume a fixed 174 AI credits, contradicting documentation that only actual requests should count. Financial impact for users on limited plans.

6. **#3885 – Long text not scrolling inside input**  
   *Area: input-keyboard, terminal-rendering* – When editing multi-line prompts, the mouse wheel scrolls the outer view instead of the textarea. Workaround needed for longer prompts.

7. **#3884 – No docs for enterprise policy enforcement of local sandbox**  
   *Area: permissions, enterprise* – While Intune/MDM management is mentioned, there is zero detail on how enterprise admins can configure and enforce sandbox policies.

8. **#3883 – Feature request: i18n support for top 10 languages**  
   *Area: theming-accessibility* – Users want menus, prompts, and errors displayed in their preferred language. First milestone: 10 most-spoken languages. 1 👍, requests already emerging.

9. **#3881 – Wrong AI credit subtraction (5% instead of 2%)**  
   *Area: models* – A single request with a 6× multiplier consumed 5% of quota instead of the expected 2%. User reports a 3% discrepancy and requests refund.

10. **#1632 – Support subfolders for skills**  
    *Area: plugins* – Users with 10+ skills cannot organize them into subfolders. The flat folder structure is cumbersome. 20 👍, 8 comments – high demand.

---

## Key PR Progress

Only one pull request was updated in the last 24 hours:

- **#3880 – [OPEN] “beyond the streets of america”**  
  *Author: 4tha5* – A test/spam PR that imports UI components with no functional relevance to the CLI. No discussion or reviews. Not actionable.

No significant PR merges or reviews occurred today.

---

## Feature Request Trends

The most-requested feature directions from recent issues:

- **Elapsed-time displays** – Multiple issues (#3278, #3055, #3111) ask for per-response, per-shell-command, and agent-thinking timers.
- **Skills folder hierarchy** – Organization via subfolders (#1632) is highly upvoted.
- **Internationalization (i18n)** – #3883 raises support for non-English locales, starting with top-10 languages.
- **MCP improvements** – Better handling of MCP server instructions (#1579) and cross-editor integration (VS Code) (#3638).
- **Plugin install optimization** – Sparse checkout (#2399) to avoid downloading full repo when only assets are needed.
- **Enterprise sandbox documentation** – Clear policy enforcement guidance for administrators (#3884, #3861).

---

## Developer Pain Points

- **Session/auth flakiness** – Resuming sessions can lose authentication (#3596) and restarting consumes unexplained credits (#3886).
- **Windows stability** – ARM64 crash under load (#3687) undermines reliability for high-concurrency terminal users.
- **Regression in file references** – The broken `@` autocomplete (#3854) disrupts a core workflow with no workaround.
- **Misaligned documentation** – Sandbox network filtering is advertised but non-functional (#3861), eroding trust.
- **Input ergonomics** – Long input scroll bug (#3885) hinders editing.
- **Quota miscalculation** – Credit usage does not match documentation (#3881), causing unexpected billing concerns.
- **Lack of timer visibility** – Users cannot estimate how long operations take, especially in autopilot mode (#3278).

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest – 2026-06-23

## Today's Highlights
Version **1.48.0** shipped with an escalated repeated-tool-call handling mechanism and a fix for empty reasoning content in the `kosong` provider. Several new bugs surfaced around MCP server auto-discovery, workspace-relative path breaking, and a hang after detached child-process calls, while the long-standing Memory System feature request (#1283) continues to gather community interest.

## Releases
**v1.48.0** – [Release](https://github.com/MoonshotAI/kimi-cli/releases/tag/1.48.0)  
- `fix(kosong)`: round-trip empty reasoning content (PR #2446, by @RealKai42)  
- `feat(soul)`: escalate repeated-tool-call reminders with three tiers and force-stop on dead-end streak (PR #2466, by @jackfish212)  
- `chore(release)`: bump kimi-cli to 1.48.0 and kosong to 0.54.0 (PR #2467, by @sailist)

## Hot Issues
*(All 6 open issues updated in the last 24h are listed.)*

- **[#1283] Memory System – Persistent context across sessions**  
  *[enhancement]* Long-standing feature request for AI-managed and user-defined memory. 6 comments, no reactions yet.  
  [Issue #1283](https://github.com/MoonshotAI/kimi-cli/issues/1283)

- **[#2457] Auto-discovers deleted MCP server, causing unfixable 400 errors**  
  *[bug]* After a user deletes an MCP server, the CLI re-discovers it and produces permanent 400 errors. Affects `kimi-code` subscription.  
  [Issue #2457](https://github.com/MoonshotAI/kimi-cli/issues/2457)

- **[#2469] `kimi web` starts MCP servers from CLI installation directory**  
  *[bug]* Workspace-relative MCP tools break because the working directory is the CLI install path. Reported on v0.18.0.  
  [Issue #2469](https://github.com/MoonshotAI/kimi-cli/issues/2469)

- **[#2468] CLI hangs after detached child-process tool call**  
  *[bug]* Using a local mock provider, the CLI hangs indefinitely after a detached child-process tool call. Linux x86_64.  
  [Issue #2468](https://github.com/MoonshotAI/kimi-cli/issues/2468)

- **[#2465] `kosong` OpenAILegacy emits `reasoning_effort: null` for thinking off**  
  *[bug]* Invalid JSON schema – `null` is not allowed; the field should be absent or a valid enum. Does not actually disable reasoning.  
  [Issue #2465](https://github.com/MoonshotAI/kimi-cli/issues/2465)

- **[#2464] `kimi acp` does not load MCP servers**  
  *[bug]* The `--mcp-config-file` flag is inert in ACP mode, while interactive mode loads MCP tools correctly. macOS arm64.  
  [Issue #2464](https://github.com/MoonshotAI/kimi-cli/issues/2464)

## Key PR Progress
*(Only 2 PRs were updated in the last 24h.)*

- **[#2467] chore(release): bump kimi-cli to 1.48.0 and kosong to 0.54.0**  
  *Closed* – Internal version bump, no changelog entries. Updated root `kosong[contrib]` pin and `kimi-code` wrapper.  
  [PR #2467](https://github.com/MoonshotAI/kimi-cli/pull/2467)

- **[#2466] feat(soul): escalate repeated-tool-call reminders and force-stop on dead-end streak**  
  *Closed* – Ported repeated-tool-call handling from `kimi-code`. Injects escalating reminders (r1/r2/r3) after 3+ consecutive identical calls, and force-stops the conversation turn on a dead-end streak.  
  [PR #2466](https://github.com/MoonshotAI/kimi-cli/pull/2466)

## Feature Request Trends
The dominant feature request remains **a comprehensive Memory System** (#1283) for cross-session context persistence, including automatic AI-managed notes and user-defined instructions. No other feature requests were updated in the last 24h, suggesting this is the single most desired direction.

## Developer Pain Points
- **MCP server configuration fragility**: Multiple bugs report auto-discovery of deleted servers (#2457), workspace-relative path errors (#2469), and MCP tools being absent in ACP mode (#2464). These indicate poor isolation and configuration handling.
- **API compatibility and provider issues**: The `kosong` provider’s invalid `reasoning_effort: null` (#2465) causes strict API failures and fails to disable reasoning. This affects any user relying on OpenAI-compatible local or strict providers.
- **Hangs and deadlocks**: The detached child-process hang (#2468) and the repeated-tool-call dead-end streak (addressed by PR #2466) show ongoing pain points in tool-call lifecycle management.

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest – 2026-06-23

## Today’s Highlights
- No new releases landed, but the community saw a wave of **security‑focused bug reports** (directory traversal, destructive commands, TOCTOU race) filed by the prolific reporter `LifetimeVip`, all closed quickly.
- A long‑standing issue with **Codex model listing** (`gpt-5.5-pro` shown but rejected) was fixed by two merged PRs, and a **child session event forwarding** feature for `opencode run --format json` is progressing through review.
- Several **MCP‑related fixes** (null parameters, Accept header, high‑precision numbers) are in the PR pipeline, reflecting continued polish of the MCP integration.

## Releases
No new versions in the last 24 hours.

## Hot Issues (10 noteworthy)

1. **[#28472 – MCP tool parameters of type "object" serialized as strings](https://github.com/anomalyco/opencode/issues/28472)**  
   *5 comments, 👍1* – Top‑level `body` parameters are passed as JSON strings instead of objects, causing MCP input validation errors. Closed May 20, but the root cause is still affecting users relying on object‑typed tool arguments.

2. **[#25785 – AI should respect repo templates when using GitHub account](https://github.com/anomalyco/opencode/issues/25785)**  
   *5 comments* – Models like DeepSeek ignore `CONTRIBUTING.md`, PR templates, and language conventions. A related fix PR [#31989](https://github.com/anomalyco/opencode/pull/31989) was merged to auto‑discover these files.

3. **[#32435 – Codex/Codex‑Pro models listed but rejected for ChatGPT OAuth users](https://github.com/anomalyco/opencode/issues/32435)**  
   *4 comments, 👍1* – `gpt-5.5-pro` appears as selectable but fails with “model not supported”. Closed by PRs [#33400](https://github.com/anomalyco/opencode/pull/33400) and [#32612](https://github.com/anomalyco/opencode/pull/32612).

4. **[#33063 – Todo dock UI doesn’t refresh after `todowrite` tool update](https://github.com/anomalyco/opencode/issues/33063)**  
   *4 comments* – Stale `session_todo` data in SolidJS reactivity. Closed with a fix that properly triggers re‑renders.

5. **[#33395 – DeepSeek V4 Pro (Max) shows empty response with official provider](https://github.com/anomalyco/opencode/issues/33395)**  
   *3 comments* – No error, just blank output until a full reinstall. Works via OpenRouter. Still **open**, with minimal debugging info.

6. **[#25650 – MCP Streamable HTTP missing Accept header](https://github.com/anomalyco/opencode/issues/25650)**  
   *3 comments* – Connections to remote MCP servers fail with HTTP 400 because the `Accept` header does not include `text/event-stream`. Closed with a fix.

7. **[#25351 – `/models` command shows presets instead of actual LM Studio models](https://github.com/anomalyco/opencode/issues/25351)**  
   *3 comments, 👍2* – Hardcoded list shown instead of querying the local API. Users must manually override, causing confusion.

8. **[#25816 – Factory reset, session management, and cache cleanup tools](https://github.com/anomalyco/opencode/issues/25816)**  
   *3 comments, 👍2* – Request for abilities to reset all local state, manage sessions, and clear caches. High community interest.

9. **[#33071 – Bash tool TOCTOU race in Windows process tree cleanup](https://github.com/anomalyco/opencode/issues/33071)**  
   *3 comments* – `taskkill /T /F` on Windows leaves orphan processes due to race condition. Security concern, quickly closed.

10. **[#33077 – Bash tool lacks protection against destructive commands](https://github.com/anomalyco/opencode/issues/33077)**  
    *3 comments* – No safety checks for `git reset --hard`, `rm -rf`, etc. Users can accidentally destroy data. Closed with recommendation to use `tool.execute.before` hooks.

## Key PR Progress (10 important)

1. **[#33404 – refactor(core): drop legacy compaction event](https://github.com/anomalyco/opencode/pull/33404)**  
   *Open* – Removes old V2 compaction event while preserving canonical V1 session history. A core infrastructure cleanup.

2. **[#27231 – feat: add edit button for connected providers](https://github.com/anomalyco/opencode/pull/27231)**  
   *Open* – Allows users to edit provider configurations (API keys, endpoints) directly from the UI. Closes #20598.

3. **[#33374 – fix(prompt): guard against deleting backups/credentials on cleanup tasks](https://github.com/anomalyco/opencode/pull/33374)**  
   *Closed* – Prevents broad “clean up old files” instructions from removing user credentials or backups. Merged quickly.

4. **[#33403 – feat(run): forward child session events to NDJSON stream](https://github.com/anomalyco/opencode/pull/33403)**  
   *Open* – Closes #33397. Adds `subtask_event`/`subtask_delta` to `opencode run --format json`, enabling real‑time visibility into sub‑agent progress.

5. **[#33227 – fix(opencode): improvements to workspace feature on TUI](https://github.com/anomalyco/opencode/pull/33227)**  
   *Closed* – Enables switching workspaces while sessions run, plus subpath filtering. Addresses #25398.

6. **[#33400 – fix: don’t show gpt-5.5-pro when using codex oauth](https://github.com/anomalyco/opencode/pull/33400)**  
   *Closed* – Quick fix for the mis‑listed models. Closes #32435.

7. **[#33160 – fix(mcp): prevent null parameters in MCP tool calls](https://github.com/anomalyco/opencode/pull/33160)**  
   *Open* – Stops MiniMax and other OpenAI‑compatible providers from receiving `null` values when parameters lack an explicit `type` field. Closes #21080.

8. **[#33082 – docs(rfc): Computer Use for opencode](https://github.com/anomalyco/opencode/pull/33082)**  
   *Open* – RFC for mouse/keyboard automation. No code changes; seeks design alignment first.

9. **[#33393 – fix(core): distinguish WebFetch URL errors from network errors](https://github.com/anomalyco/opencode/pull/33393)**  
   *Open* – Prevents malformed URLs and network failures from being conflated in error messages. Closes #33073.

10. **[#33392 – feat(llm): pass `strict` through tool definitions for Codex parity](https://github.com/anomalyco/opencode/pull/33392)**  
    *Open* – Mirrors Codex CLI’s `strict: false` on function tools to avoid rejection of dynamic schemas. Aims for better compatibility.

## Feature Request Trends
- **Session & Data Management** (#25816, #26078, #33063): Users want factory reset, archived session visibility, and a re‑sync UI for cached data.
- **UI/UX Polish** (#25872, #25871, #25923): Edit/recall/delete of messages, quick‑command selector, and one‑click translation for foreign‑language AI output.
- **Repository Convention Awareness** (#25785, #33374): AI should auto‑discover and respect `CONTRIBUTING.md`, PR templates, and language preferences.
- **Guided Configuration** (#25815, #25817): Wizards for `AGENTS.md` generation and preset instruction templates to avoid repeating prompts.
- **Docker/Container Safety** (#25610, #33077): Protection against destructive file operations and explicit safe‑zone marking in the sidebar.

## Developer Pain Points
- **MCP Integration Gaps** – Object parameters serialized as strings (#28472), missing Accept headers (#25650), high‑precision number loss (#28124), and null parameter handling (#33160) frustrate tool interoperability.
- **Model Compatibility** – DeepSeek V4 Pro fails silently with official provider (#33395), Codex‑Pro models mis‑listed (#32435), and LM Studio model list is hardcoded (#25351).
- **Security & Reliability** – Directory traversal in Glob/Grep (#33070), destructive command execution (#33077), WebFetch User‑Agent blocker (#33074), and Windows process‑tree TOCTOU race (#33071) raise concerns about inadvertent data loss or bypasses.
- **High CPU / Unresponsiveness** – Random 99‑100% CPU spikes (#33399) make the CLI unusable; root cause unknown.
- **UI Usability** – Scaling issues hiding confirm buttons (#25583), stale todo dock (#33063), and inability to edit/delete past messages (#25872) reduce everyday productivity.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-06-23

## Today's Highlights
A minor release (v0.79.10) enriches compaction events with `reason` and `willRetry`, closing a long-standing gap between RPC and extension APIs. Meanwhile, the community is buzzing about persistent connection drops with `openai-codex`/GPT-5.5 (64 comments) and a longstanding request for a first-class local LLM provider extension (36 👍). On the PR side, an Anthropic Vertex provider lands, and critical fixes address agent loop hangs and secret disclosure in broad file operations.

## Releases
**v0.79.10** — [View on GitHub](https://github.com/earendil-works/pi/releases/tag/v0.79.10)  
- **Extension compaction event context**: `session_before_compact` and `session_compact` events now include `reason` (`"manual"`, `"threshold"`, `"overflow"`) and `willRetry`. Extensions can finally distinguish manual `/compact` from auto-compaction or overflow retry flows. This closes issue [#5217](https://github.com/earendil-works/pi/issues/5217).

## Hot Issues (10 of 43)

1. **[#4945 – openai-codex Connection Reliability Issues](https://github.com/earendil-works/pi/issues/4945)**  
   *Status: Open, inprogress* | Comments: 64 | 👍: 30  
   **Summary**: The TUI gets stuck on "Working..." with no streaming text, tool calls, or errors when using `openai-codex`/`gpt-5.5`. Only pressing Esc recovers. Frequent over the last few days.  
   **Why it matters**: Blocks interactive use for a major provider; high community impact.

2. **[#3357 – Official local LLM provider extension](https://github.com/earendil-works/pi/issues/3357)**  
   *Status: Open* | Comments: 26 | 👍: 36  
   **Summary**: Request to fetch model list from `{baseUrl}/models` dynamically, enabling seamless integration with llama.cpp, Ollama, LM Studio, etc.  
   **Why it matters**: Most-requested feature by 👍; would unlock self-hosted and offline workflows.

3. **[#5653 – Move off Shrinkwrap](https://github.com/earendil-works/pi/issues/5653)**  
   *Status: Open, inprogress, to-discuss* | Comments: 15  
   **Summary**: Installing both `pi-ai` and `pi-coding-agent` as direct deps duplicates `pi-ai` on disk, causing separate module-level `Map` instances for the provider registry.  
   **Why it matters**: Breaking plugin isolation; a core packaging issue for extension authors.

4. **[#5916 – Support provider extensions with model aliases and improve search](https://github.com/earendil-works/pi/issues/5916)**  
   *Status: Open, inprogress* | Comments: 11  
   **Summary**: No UI to configure OpenRouter providers; models.json overrides work but are confusing. Needs model alias support and search improvements.  
   **Why it matters**: Power users want to mix providers; lack of UI frustrates adoption.

5. **[#5571 – pi -p hangs indefinitely when stdin is a non-TTY pipe](https://github.com/earendil-works/pi/issues/5571)**  
   *Status: Closed* | Comments: 10  
   **Summary**: `pi -p` with a non-TTY pipe and no credentials freezes instead of failing fast.  
   **Why it matters**: A regression that breaks CI/CD and scripting pipelines.

6. **[#5778 – pi-agent-core hangs indefinitely on unresponsive streams](https://github.com/earendil-works/pi/issues/5778)**  
   *Status: Closed* | Comments: 8  
   **Summary**: Agent loop wedges when LLM provider stream drops connection or tool `execute()` never resolves.  
   **Why it matters**: Critical vulnerability; tool execution deadlocks can freeze the whole agent.

7. **[#5939 – Make auto-compaction opt-in and safe](https://github.com/earendil-works/pi/issues/5939)**  
   *Status: Closed, no-action* | Comments: 7  
   **Summary**: Proposes auto-compaction run only after tool-use turn finishes and before next provider request, defaulting to off.  
   **Why it matters**: Safety vs performance tradeoff; community wants controlled behaviour.

8. **[#4748 – pi-tui getKeybindings() singleton breaks extensions](https://github.com/earendil-works/pi/issues/4748)**  
   *Status: Open* | Comments: 5 | 👍: 2  
   **Summary**: `globalKeybindings` module-scope singleton is duplicated across extension `node_modules`, breaking key binding detection.  
   **Why it matters**: Blocks extensions that rely on TUI key labels; architectural flaw.

9. **[#5263 – Make in-session model and thinking-level changes ephemeral](https://github.com/earendil-works/pi/issues/5263)**  
   *Status: Open* | Comments: 4 | 👍: 4  
   **Summary**: Changes should only affect the active session; add a "Default model" entry in `/settings` for global defaults.  
   **Why it matters**: Improves UX for users who switch models mid-session without persisting unintended defaults.

10. **[#5871 – Anthropic OAuth-token detection is hardcoded](https://github.com/earendil-works/pi/issues/5871)**  
    *Status: Open, inprogress* | Comments: 4  
    **Summary**: The `isOAuthToken()` check only looks for `sk-ant-oat` prefix; providers/models should be able to declare OAuth explicitly.  
    **Why it matters**: Breaks custom Anthropic-compatible providers with different key formats.

## Key PR Progress (7 PRs)

1. **[#5262 – feat(ai): add Anthropic Vertex provider](https://github.com/earendil-works/pi/pull/5262)** (Open)  
   Adds a built-in `anthropic-vertex` provider for Claude on Google Cloud Vertex AI. Thin adapter using `AnthropicVertex` SDK, reuses existing streaming/tool path.  
   **Why it matters**: Opens Pi to GCP users who need Claude; simplifies onboarding.

2. **[#5970 – feat: add auto-router extension for DeepSeek V4 Pro/Flash cost optimization](https://github.com/earendil-works/pi/pull/5970)** (Closed)  
   New extension `auto-router.ts` analyses prompt complexity and routes between DeepSeek V4 Flash (simple) and Pro (complex), claiming 60–70% cost savings.  
   **Why it matters**: Demonstrates extension API power for cost-aware routing; community can adapt for other providers.

3. **[#5962 – feat(coding-agent): add compaction reason and willRetry to extension compaction events](https://github.com/earendil-works/pi/pull/5962)** (Closed)  
   Closes #5217. `session_before_compact` and `session_compact` now carry `reason` and `willRetry`. Matches RPC protocol.  
   **Why it matters**: Unifies extension and RPC signals; enables smarter compaction logic in extensions.

4. **[#5963 – fix(ai): reject malformed final tool call arguments](https://github.com/earendil-works/pi/pull/5963)** (Closed)  
   Validates final streamed tool-call argument JSON; fails with `stopReason: "error"` before exposing malformed arguments.  
   **Why it matters**: Prevents tool execution crashes from bad model output; improves robustness.

5. **[#5941 – fix(coding-agent): add required reason and willRetry to compaction ex…](https://github.com/earendil-works/pi/pull/5941)** (Closed)  
   Similar to #5962 but with `required` fields.  
   *Note: This PR duplicates #5962 scope; likely superseded.*

6. **[#5955 – fix(coding-agent): add secret-disclosure scope discipline to the default system prompt](https://github.com/earendil-works/pi/pull/5955)** (Closed)  
   Prevents Pi from copying secret files (`.env`, private keys) into world‑readable directories on broad "copy everything" tasks.  
   **Why it matters**: Security fix; addresses real-world incident (#5956) where sensitive data was exposed.

7. **[#5950 – fix: use OpenRouter's actual cost from API response in footer](https://github.com/earendil-works/pi/pull/5950)** (Closed)  
   Replaces static per‑token estimate with `usage.cost` from OpenRouter response.  
   **Why it matters**: Accurate billing display for both built‑in and custom models; reduces surprise charges.

## Feature Request Trends
- **Local & self-hosted LLM support** (#3357, #5871): Dominant demand for dynamic model list fetching and OAuth flexibility.
- **Provider configuration UX** (#5916, #5965): Users want UI/CLI to manage provider aliases, model overrides, and relabel unclear provider names.
- **Session management APIs** (#5804, #5810, #5912, #5932): Extensions and RPC clients need `navigateTree()`, `get_entries`, `get_tree`, and programmatic session switching.
- **Safety & security** (#5939, #5955, #5956): Auto-compaction should be opt-in and safe; secret disclosure prevention in file operations.
- **Performance & reliability** (#5778, #4945, #5571): Streaming resilience, timeout handling, and better error messages for connection drops.

## Developer Pain Points
- **Hangs and freezes**: Non-TTY pipe stalls (#5571), agent loop deadlocks (#5778), openai-codex silent "Working..." (#4945) – all high‑frequency frustrations.
- **Dependency duplication**: Shrinkwrap causing separate module instances breaks extension registries (#5653).
- **Singleton architectural issues**: `globalKeybindings` singleton not shared across extension realms (#4748).
- **Hardcoded provider logic**: OAuth token prefix check (#5871), lack of dynamic model listing (#3357).
- **Missing tooling**: No UI for OpenRouter provider config (#5916), incomplete docs for slash commands (#5959), broken example links (#5957).
- **Accidental secret exposure**: Broad file operations can leak sensitive files (#5955, #5956) – a security concern for automated pipelines.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-06-23

## Today's Highlights
- **Critical release failure** – v0.19.0 and two preview releases failed due to integration tests and publish jobs, prompting a hotfix for settings migration idempotency.
- **Edge‑input validation sweep** – Six PRs landed to enforce integer‑only env values for timeouts, byte limits, and counts, fixing silent rounding bugs in Mermaid renderer, inline media, stop hooks, and compaction.
- **Provider/auth UX improvements** – PRs add fast‑only/voice‑only model flags, restore custom model IDs in the auth wizard, and decouple provider identity from SDK protocol (issue #5090 closed).

---

## Releases
**v0.18.5‑nightly.20260622** – Routine nightly; no user‑facing changes beyond CI adjustments (auto‑publish VSCode companion after stable releases).  
Full changelog: [GitHub Release](https://github.com/QwenLM/qwen-code/releases/tag/v0.18.5-nightly.20260622.6bc3f853e)

---

## Hot Issues (Top 10)

1. **[#5090 – Refactor: Decouple Provider Identity from SDK Protocol](https://github.com/QwenLM/qwen-code/issues/5090)** *(CLOSED)*  
   Proposes making `providerId` a free‑form string and introducing a `Protocol` enum for SDK routing. This unlocks fully custom providers while keeping API calls type‑safe. Community agreed; now merged.

2. **[#5656 – Move tool‑use summaries to loading indicator](https://github.com/QwenLM/qwen-code/issues/5656)** *(OPEN, P3)*  
   Suggests moving short tool‑call summaries (e.g. "Searched in auth/") from conversation history into the loading indicator, reducing visual clutter. Attracted 5 comments debating UX trade‑offs.

3. **[#5634 – Autofix tier‑1 trusts ghost label](https://github.com/QwenLM/qwen-code/issues/5634)** *(OPEN, P2, security)*  
   Highlights a security gap: the autofix pipeline fast‑tracks issues with `status/ready-for-agent` – a label that LLM‑generated issue text can influence without human review. Critical for safe automation.

4. **[#5641 – Repeating shell tool results on npm latest](https://github.com/QwenLM/qwen-code/issues/5641)** *(OPEN, P2, core)*  
   Deterministic providers cause Qwen Code to re‑submit already‑returned shell tool results, leading to loops. A reproducer against the public npm package is included. High confidence bug.

5. **[#5611 – web_fetch can’t fetch JSON APIs (HTTP 415)](https://github.com/QwenLM/qwen-code/issues/5611)** *(OPEN, P2)*  
   The `web_fetch` tool sends only `text/*` Accept headers, so JSON‑only endpoints (e.g. GitHub API) fail. Basic request – a quick fix is in progress (#5660).

6. **[#5672 – Mermaid renderer timeout accepts fractional env values](https://github.com/QwenLM/qwen-code/issues/5672)** *(OPEN, P3)*  
   `QWEN_CODE_MERMAID_RENDER_TIMEOUT_MS=0.4` passes validation then rounds to 0ms, effectively disabling rendering. One example of a wider pattern of weak env parsing.

7. **[#5665 – AI PRs often miss integration‑test updates](https://github.com/QwenLM/qwen-code/issues/5665)** *(OPEN, P2, enhancement)*  
   Pattern observed: AI‑assisted PRs update product code but skip `integration-tests/` changes, surfacing failures only during release. Requests automated diff‑based test reminders.

8. **[#5664 – stopHookBlockingCap accepts fractional values](https://github.com/QwenLM/qwen-code/issues/5664)** *(OPEN, P2, welcome‑pr)*  
   Similar fractional‑parsing bug for `QWEN_CODE_STOP_HOOK_BLOCK_CAP`, causing early hook loop protection to fire unexpectedly. Tagged `welcome-pr`.

9. **[#5675 – Validate QWEN_CODE_IDE_SERVER_PORT before lock file access](https://github.com/QwenLM/qwen-code/issues/5675)** *(OPEN, P2, security)*  
   The port value is used unsanitised to build a lock‑file path, opening a path‑traversal vector. Currently `in-review`.

10. **[#5663 / #5659 / #5653 – Three release failures (v0.19.0, previews)](https://github.com/QwenLM/qwen-code/issues/5663)** *(CLOSED/OPEN)*  
    v0.19.0 failed at `integration_none` and `integration_docker`; two preview releases also failed (publish and Docker). Led to immediate hotfix PR #5676.

---

## Key PR Progress (Top 10)

1. **[#5676 – fix(cli): keep settings v5 migration idempotent](https://github.com/QwenLM/qwen-code/pull/5676)**  
   Directly addresses the v0.19.0 release failure by ensuring the settings migration doesn't break on re‑run. Targeted fix, merged quickly.

2. **[#5666 – docs(tui): design for Ctrl+O transcript view, removing compact mode](https://github.com/QwenLM/qwen-code/pull/5666)**  
   Design‑first PR proposing to replace the global compact/normal toggle with a single baseline view and a separate full‑detail transcript overlay (like Claude Code). Opened for community feedback.

3. **[#5674 – fix(cli): require integer Mermaid render timeout](https://github.com/QwenLM/qwen-code/pull/5674)**  
   Rejects fractional env values for `QWEN_CODE_MERMAID_RENDER_TIMEOUT_MS`. Falls back to default on invalid input. Part of the enviro‑input validation sweep.

4. **[#5632 – feat(core): add fastOnly/voiceOnly flags to hide models from main list](https://github.com/QwenLM/qwen-code/pull/5632)**  
   Allows provider configs to mark models as `fastOnly` or `voiceOnly` so they only appear in the appropriate sub‑selectors (fast model dialog, voice mode). Clean UX improvement.

5. **[#5661 – feat(tui): unify tool output with semantic summaries](https://github.com/QwenLM/qwen-code/pull/5661)**  
   Replaces dual compact/normal tool output modes with a single unified view that shows semantic overviews (e.g. "Read 3 files, edited 2 files"). Removes raw result display clutter.

6. **[#5668 – feat(cli): show model thinking intent in loading indicator](https://github.com/QwenLM/qwen-code/pull/5668)**  
   Replaces random witty loading phrases with the model's real‑time `ThoughtSummary` text. Direct response to issue #5656 (and broader wish for live feedback).

7. **[#5662 – ci(autofix): gate fast path on trusted label](https://github.com/QwenLM/qwen-code/pull/5662)** *(CLOSED)*  
   Mitigates the security issue from #5634 by replacing `status/ready-for-agent` with a dedicated `autofix/ready` label that only CI‑created issues can set. Merged quickly.

8. **[#5660 – fix(core): allow web_fetch JSON fallback](https://github.com/QwenLM/qwen-code/pull/5660)**  
   Adds `*/*;q=0.1` to the Accept header so JSON‑only endpoints respond. Fixes #5611. Simple but high‑impact for API‑driven workflows.

9. **[#5657 – fix(cli): stop repeated duplicate provider responses](https://github.com/QwenLM/qwen-code/pull/5657)**  
   Fixes the tool‑result loop described in #5641 by detecting duplicate provider tool‑call IDs and injecting a synthetic error response to break the cycle.

10. **[#4943 – feat(cli): add --safe-mode flag](https://github.com/QwenLM/qwen-code/pull/4943)**  
    Adds a `--safe-mode` flag (and env var) that disables all customizations (hooks, extensions, MCP, etc.) for troubleshooting. Still in review; has accumulated community interest over two weeks.

---

## Feature Request Trends
- **Custom provider autonomy** – The biggest trend: freeing users from SDK‑enforced provider identities (#5090 closed) and improving the UI for adding arbitrary models to custom providers (#4814). Users want to plug in any OpenAI‑compatible endpoint without workarounds.
- **Leaner conversation UI** – Repeated requests to move tool‑use summaries out of the main history and into loading indicators or status bars (#5656, #5661, #5668). Aligns with Claude Code’s minimal footprint.
- **Daemon/ACP parity** – A tracking issue (#5677) identifies CLI commands (`cd`, `permissions`, `trust`, `lsp`, `setup-github`) that still lack daemon‑side ACP support. Users expect distributed/remote sessions to work fully.
- **Model visibility controls** – `fastOnly` and `voiceOnly` flags (#5632) reflect a desire to reduce model list noise, especially when many aliases or provider‑specific models are registered.

---

## Developer Pain Points
- **Silent env‑value bugs** – Today’s top frustration: fractional or invalid env values (timeouts, byte limits, hook caps, counts) are silently accepted and floored, leading to baffling behaviour. Half a dozen issues (#5672, #5664, #5669, etc.) and matching PRs show this is a systematic validation gap.
- **Release pipeline fragility** – Three separate release failures in one day (#5663, #5659, #5653) due to integration tests and publish steps. Combined with AI‑assisted PRs skipping integration‑test updates (#5665), the community is calling for better CI‑gating and automated test‑path reminders.
- **Provider‑specific tool loops** – Deterministic/re‑entrant providers cause Qwen Code to replay completed tool calls, leading to infinite‑looking loops (#5641). The fix (#5657) is narrow; deeper root cause (history deduplication) may need follow‑up.
- **web_fetch JSON blindness** – Lacking a JSON‑compatible Accept header breaks many common use cases (API fetching). The quick fix (#5660) is welcomed, but users note the tool should ideally inspect Content‑Type and auto‑select parser in future.
- **Autofix security theatre** – The discovery that LLM‑set labels could hijack the automated fix pipeline (#5634) eroded trust in CI automation. The quick label change (#5662) is a band‑aid; proper provenance tracking is now on the roadmap.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest – 2026-06-23

## Today’s Highlights

The CodeWhale v0.8.64 release train is in full swing with a major security hardening push (29-comment tracker #3368), a wave of community-contributed fixes, and several new feature PRs adding dev-server readiness, idle‑timeout suppression, and WeCom integration. Key pain points around agent over‑extension and sandbox worktree support have been addressed through targeted PRs and ongoing discussions.

## Releases

- **v0.8.64** (CodeWhale) – This release marks the official rebrand from `deepseek-tui` to `CodeWhale`. It includes security hardening and code-scanning fixes (see tracker #3368), plus the infrastructure for auto‑review policies and read‑before‑edit guardrails. The legacy npm package `deepseek-tui` is deprecated; users should follow `docs/REBRAND.md` for migration.

## Hot Issues

1. **[#3368 – Security hardening/code-scanning fixes](https://github.com/Hmbown/CodeWhale/issues/3368)**  
   *29 comments* – The central tracker for v0.8.64’s security‑hardening work (CodeQL, advisory reports, local integration commits). Community engagement is high as the team needs a public release gate without disclosing exploit details.

2. **[#3144 – Natural‑language auto‑review policy](https://github.com/Hmbown/CodeWhale/issues/3144)**  
   *12 comments* – Inspired by Cursor’s review workflow, this proposes a pre‑push review gate with configurable policies. Widely discussed as a way to balance manual approval and unchecked autonomy.

3. **[#3275 – Agent over‑extension (self‑questioning)](https://github.com/Hmbown/CodeWhale/issues/3275)**  
   *11 comments* – A regression from #3061; the agent loops by proposing, answering, and executing without waiting for user confirmation. Criticised as a major trust‑breaker for smaller models.

4. **[#3222 – `reasoning_style` override for MiniMax M3 / Qwen / GLM](https://github.com/Hmbown/CodeWhale/issues/3222)**  
   *6 comments* – Broken parsing of reasoning content from non‑OpenAI providers. Users want an inline‑tag based override to keep thinking blocks working.

5. **[#3369 – Restore nightly cross‑target builds and auto‑tag idempotency](https://github.com/Hmbown/CodeWhale/issues/3369)**  
   *2 comments* – CI failures in nightly builds and auto‑tagging are blocking release readiness. The community is watching closely for a fix (PR #3374).

6. **[#3364 – Read‑before‑edit guardrails](https://github.com/Hmbown/CodeWhale/issues/3364)**  
   *1 comment* – Proposed to make edit operations safer by enforcing fresh reads and making edit failures loud. Part of the v0.8.64 reliability sprint.

7. **[#3355 – Sandbox blocks Git write ops on worktree workspaces](https://github.com/Hmbown/CodeWhale/issues/3355)**  
   *3 comments, closed* – The macOS sandbox blocked `git add` when using worktrees. Fix landed in PR #3356; users appreciative of the quick resolution.

8. **[#3189 – Idle‑timeout status line noise](https://github.com/Hmbown/CodeWhale/issues/3189)**  
   *2 comments, closed* – The persistent `waiting for model` countdown was distracting. PR #3375 now thresholds it to show only after 60s or near timeout – a UX win.

9. **[#3360 – Dev‑server readiness primitive](https://github.com/Hmbown/CodeWhale/issues/3360)**  
   *1 comment, closed* – Agents needed a reliable way to check when a local dev server is actually serving. PR #3376 adds a `wait_for_dev_server` tool.

10. **[#3380 – Approval modal key hints more prominent](https://github.com/Hmbown/CodeWhale/issues/3380)**  
    *0 comments* – Labelled “good first issue”. First‑time users miss the low‑contrast `y/a/d` hints. Improving this could lower the learning curve for new contributors.

## Key PR Progress

1. **[#3356 – Allow worktree git metadata writes in sandbox](https://github.com/Hmbown/CodeWhale/pull/3356)**  
   *Closed* – Fixes #3355. Detects linked‑worktree `.git` pointers and allows writes to those paths without requiring trust mode.

2. **[#3374 – Restore nightly cross‑target builds and auto‑tag idempotency](https://github.com/Hmbown/CodeWhale/pull/3374)**  
   *Closed* – Implements artifact‑existence checks and robust idempotency controls for nightly.yml and auto‑tag workflows.

3. **[#3371 – Reduce minimum terminal width for sidebar visibility](https://github.com/Hmbown/CodeWhale/pull/3371)**  
   *Closed* – Lowers the sidebar threshold from 100 to a more typical width (exact value in PR). Fixes #3328.

4. **[#3350 – Add /model pro|flash shortcuts and CLI model set command](https://github.com/Hmbown/CodeWhale/pull/3350)**  
   *Closed* – Introduces `pro` and `flash` aliases for `deepseek-v4-pro`/`flash` and adds `codewhale model set` subcommand. Cleans up model switching.

5. **[#3373 – v0.8.64 security and release integration](https://github.com/Hmbown/CodeWhale/pull/3373)**  
   *Closed* – Draft integration PR carrying local security hardening, auto‑review rails, read‑before‑edit guardrails, CI fixes, and community credit harvests.

6. **[#3345 – Refactor(config): move inline tests to module](https://github.com/Hmbown/CodeWhale/pull/3345)**  
   *Closed* – Moves config tests out of `lib.rs` into `tests.rs`, reducing file size and merge conflict risk. Closes #3307.

7. **[#3377 – Update security contact email to codewhale.com](https://github.com/Hmbown/CodeWhale/pull/3377)**  
   *Closed* – Part of the v0.8.64 security hardening: updates SECURITY.md to reflect the CodeWhale rebrand.

8. **[#3376 – Add dev‑server readiness tool](https://github.com/Hmbown/CodeWhale/pull/3376)**  
   *Open* – Implements `wait_for_dev_server` tool (TCP/HTTP loopback only) wired into TUI tool registry, catelog, approval, and history. Addresses #3360.

9. **[#3375 – Suppress idle timeout countdown in provider‑wait footer](https://github.com/Hmbown/CodeWhale/pull/3375)**  
   *Open* – Changes footer display: hides countdown below 60s, shows `< 75% budget` format only close to limit. Fixes #3189.

10. **[#3370 – Add WeCom (企业微信) intelligent robot bridge](https://github.com/Hmbown/CodeWhale/pull/3370)**  
    *Open* – Adds an integration for WeCom (WeChat Work) robot, enabling CodeWhale to interact within Chinese enterprise messaging ecosystem.

## Feature Request Trends

- **Agent safety & trust** – High demand for auto‑review policies (#3144), read‑before‑edit guardrails (#3364), and approval modal UX improvements (#3380) to prevent unwanted autonomous actions.
- **Context & compaction** – Automatic context compaction with carried‑forward summaries (#3363) is a top‑requested comfort feature for long sessions.
- **Browser automation & testing** – Multiple requests for Playwright‑backed browser automation (#3358), console verification (#3361), and screenshot‑to‑vision feedback (#3362) for iterative UI work.
- **Customisation & modularity** – Users want user‑defined subagent personas via `.codewhale/agents/` (#3367) and ModelProfile descriptors to tailor tool surfaces per model (#3365).
- **Work‑tracking consolidation** – Overlap between plans, todos, tasks, and goals creates confusion; a single canonical work surface (#3366) would simplify model tool choice.
- **Local development workflow** – Dev‑server readiness primitives (#3360), webapp‑testing skills (#3359), and branch hygiene automation (#3214) reflect a push toward seamless local‑app testing.

## Developer Pain Points

- **Agent over‑extension** (#3275) – The agent frequently enters self‑questioning/answering loops, deviating from user intent. This is a regression and a top frustration for those using non‑flagship models.
- **Noisy idle‑timeout** (#3189) – The persistent countdown in the footer was distracting for normal waits. The fix (#3375) is welcome but users still want finer control.
- **Sandbox limitations** (#3355) – macOS seatbelt blocks Git writes on worktree workspaces, forcing users into trust mode. The worktree fix (#3356) should alleviate this.
- **Fragmented work‑tracking** (#3366) – Multiple overlapping work surfaces (plans, todos, tasks) make it hard for models to choose the right tool, especially for smaller models.
- **Lack of browser testing tools** – Absence of a rendered browser automation and console verification (#3358, #3361) forces agents to rely on flaky shell polling for UI tasks.
- **Branch hygiene complexity** (#3214) – Release branches and fork checkouts require manual remote qualification; the community has struggled with post‑merge cleanup.
- **UI discoverability** (#3380) – The approval modal’s key hints are low‑contrast and easy to miss, especially for first‑time users on smaller terminals.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*