# AI CLI Tools Community Digest 2026-06-15

> Generated: 2026-06-15 03:43 UTC | Tools covered: 9

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

# AI CLI Developer Tools — Cross-Tool Comparison Report
**Date:** 2026-06-15 | **Audience:** Technical decision-makers and developers

---

## 1. Ecosystem Overview

The AI CLI tools landscape is nearing a critical inflection point. While Claude Code and OpenAI Codex remain the most mature and feature-rich platforms, both are hobbled by billing transparency issues, cross-platform regression bugs, and growing community frustration over unaddressed feature requests. Gemini CLI is investing heavily in evaluation infrastructure and AST-aware tooling, signaling a shift from raw agent capability toward *measured reliability*. The mid-tier tools — OpenCode, Pi, and the rebranded CodeWhale (formerly DeepSeek TUI) — are converging on three shared priorities: extension APIs, multi-agent orchestration, and better memory/context management. Meanwhile, GitHub Copilot CLI and Kimi Code occupy quiet niches (enterprise GitHub integration and Chinese-market pricing, respectively), generating lower community velocity. The overall signal: **users are demanding predictable, auditable, and cross-platform tools** — not just smarter agents.

---

## 2. Activity Comparison

| Tool | Hot Issues (notable) | PRs updated (24h) | Release (24h) | Community Signal |
|------|----------------------|-------------------|---------------|------------------|
| **Claude Code** | 10 | 5 (2 open, 3 closed) | None | Highest engagement (#17432: 442👍), but critical bugs unaddressed |
| **OpenAI Codex** | 10 | 10 (all merged or in-progress) | None | Strong engineering velocity; Windows regressions piling up |
| **Gemini CLI** | 10 | 10 (heavy dependency updates) | None | High activity but many P1 bugs unresolved; 43 PRs updated |
| **GitHub Copilot CLI** | 6 | 0 | None | Lowest activity; only 6 actionable issues |
| **Kimi Code CLI** | 3 | 4 (2 merged) | None | Small volume but high-impact rate-limiting complaints |
| **OpenCode** | 10 | 10 | **v1.17.7** | Healthy release cadence; security and UX regressions |
| **Pi** | 10 | 10 | None | Very high activity (40 issues, 13 PRs updated); extension API focus |
| **Qwen Code** | 10 | 10 | None | Active triage; VSIX false-positive blocking adoption |
| **DeepSeek/CodeWhale** | 10 | 10 | **v0.8.60** | Rebrand causing migration breakage; high PR throughput |

**Key observation:** Only OpenCode and CodeWhale shipped releases in the last 24 hours. All others are in a maintenance/regression-fix phase.

---

## 3. Shared Feature Directions

### 3.1 Regional Pricing & Payment Flexibility
- **Claude Code** (#17432, 442👍): India-specific INR pricing, 5+ months without response
- **Kimi Code** (#2123): Rate limiting gap between advertised and actual capacity for Chinese subscribers
- **CodeWhale** (#2629): Authentication failures with SiliconFlow and Tencent Cloud TokenHub block Chinese users

**Trend:** Anthropic's silence on regional pricing is a reputational risk; competitors who offer local currency support gain adoption in price-sensitive markets.

### 3.2 Subagent/Spawn Control & Resource Governance
- **Claude Code** (#68430): 50+ level recursive subagent spawn, exponential token burn
- **Claude Code** (#68165, #68411): Per-message model selection, model/effort inheritance
- **Gemini CLI** (#21409, #22323): Agent hangs, false success on turn limits
- **OpenCode** (#31072, #31487): Race conditions in subagent message events, duplicate streams
- **CodeWhale** (#2652): Clipped subagent output mistaken as complete
- **CodeWhale** (#3226, #3229, #3230): Whaleflow swarm orchestration (parent-visible workers, fleet ledger)

**Common theme:** Users want **configurable depth limits, per-spawn model selection, and reliable turn accounting** for subagents. This is the #1 technical gap across all tools.

### 3.3 Session & History Management
- **Claude Code** (#41458): `cleanupPeriodDays` ignored, sessions silently deleted
- **OpenAI Codex** (#12564, 111👍): Allow renaming task/thread titles (CLOSED)
- **OpenAI Codex** (#25500, #27353): Chat history disappears after updates (cross-platform)
- **OpenCode** (#32368): Revertible compaction
- **Pi** (#5654): `excludeFromContext` for custom messages
- **Qwen Code** (#5120): Skip auto-title when history has no user message

**Takeaway:** Session persistence and organization remain fragile. Users want control over what persists, what's sent to the model, and how history is navigated.

### 3.4 Rate-Limit & Billing Transparency
- **Claude Code** (#32544): Charged despite available capacity; false rate-limit errors
- **OpenAI Codex** (#15281, 15👍): Expose full usage/limits in CLI `/status`
- **OpenAI Codex** (#28154, #28249): Rate-limit reset redemption in TUI
- **Kimi Code** (#2123): Severe rate limiting misalignment
- **Pi** (#5738): Anthropic 1-hour cache write pricing fix

**Observation:** Every tool with paid tiers faces billing complaints. The gap between advertised limits and actual behavior erodes trust. OpenAI Codex is leading by actively exposing reset credits.

### 3.5 Cross-Platform Parity
- **Claude Code** (#51143): Windows blank screen (2 months open)
- **Claude Code** (#66020): macOS kernel memory leak
- **OpenAI Codex** (#27979, #27367): Windows app crashes after update
- **Gemini CLI** (#21983): Browser agent fails on Wayland
- **OpenCode** (#13984, 48 comments): Copy-paste broken in CLI
- **OpenCode** (#32370): Linux clipboard selection (PR just merged)
- **Pi** (#5103): Windows bash detector fails on non-default drive
- **Qwen Code** (#3979): Ghostty terminal flicker on macOS

**Hard truth:** macOS is the primary target for all tools. Windows and Linux users experience second-class support with longer bug lifetimes.

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | Copilot CLI | Kimi Code | OpenCode | Pi | Qwen Code | CodeWhale |
|-----------|-------------|--------------|------------|-------------|-----------|----------|-----|-----------|-----------|
| **Primary user** | Power developers | Full-stack devs | Research/ML | Enterprise orgs | Chinese market | Plugin devs | Extension builders | OSS contributors | Multi-provider orchestration |
| **Core strength** | Agent recursion depth | Rate-limit transparency | AST-aware tools | GitHub integration | Affordability | MCP sandboxing | Extension API | Safe-mode debugging | Swarm orchestration |
| **Weakness** | Billing opacity | Windows stability | Agent hangs | Low community | Rate limits | TUI regressions | Provider compatibility | Trojan false-positive | Rebrand migration |
| **Release cadence** | Slow | Steady | Very high (43 PRs/24h) | Minimal | Moderate | Healthy | Very high | Active | High |
| **Community engagement** | Very high (442👍 issues) | High (111👍 issues) | Moderate | Low | Low | Moderate | High | Moderate | Moderate |

**Key differentiators:**
- **Gemini CLI** stands out for its investment in evaluation infrastructure (#24353, 76 behavioral tests) — a level of quality assurance no other tool matches.
- **Pi** is the most forward-looking on extension APIs, with `excludeFromContext`, `setPromptGuidelines()`, and safe reloads — positioning itself as the Emacs of AI CLI tools.
- **OpenCode** has the best release cadence and MCP security posture, but TUI regressions remain a weak point.
- **CodeWhale** is the only tool explicitly designed for multi-provider swarm orchestration — risky but differentiated.

---

## 5. Community Momentum & Maturity

### Tier 1: Mature, high-engagement
- **Claude Code** — Most upvoted issues, sustained discussion, but Anthropic's slow response risks erosion
- **OpenAI Codex** — Strong engineering response on rate-limit features; Windows stability is a growing liability

### Tier 2: High iteration velocity
- **Gemini CLI** — 43 PRs updated in 24h, heavy dependency upgrades, but P1 bugs linger
- **Pi** — 40 issues + 13 PRs; extension API advances are market-leading
- **CodeWhale** — Rapid rebrand + swarm features; migration pain is temporary but acute

### Tier 3: Stable niche
- **OpenCode** — Healthy releases, but community size is moderate; security-focused
- **Qwen Code** — Active but impacted by false-positive CVE; OSS-centric

### Tier 4: Low engagement
- **GitHub Copilot CLI** — Only 6 actionable issues; 0 PRs; appears minimally maintained
- **Kimi Code** — Small community; rate-limiting complaints dominate; limited feature velocity

---

## 6. Trend Signals — Implications for Developers

### Signal 1: Trust is the new competitive moat
Billing opacity (#32544, #15281), silent data loss (#41458), and misattributed errors (#68502) are the top community pain points. **Tools that invest in transparent usage dashboards, audit logs, and predictable behavior will win adoption.** OpenAI Codex's rate-limit reset credits are a template others should follow.

### Signal 2: Subagent governance is table-stakes
The recursive-spawn bugs at Claude Code (#68430) and Gemini CLI (#21409) are not isolated — they reflect a systemic lack of guardrails. **Developers need configurable depth limits, per-spawn model controls, and reliable turn accounting.** Tools that ship subagent governance frameworks first gain a structural advantage.

### Signal 3: Cross-platform support is a hidden barrier
Windows blank screens (Claude Code #51143, months open) and Wayland failures (Gemini CLI #21983) exclude large developer populations. **Any tool targeting enterprise adoption must treat Windows and Linux as first-class platforms** — not afterthoughts. OpenCode's Linux clipboard fix (#32370, just merged) is a model for incremental improvement.

### Signal 4: Extension APIs are becoming the differentiator
Pi's extension API expansion (prompt guidelines, excludeFromContext, safe reloads) and OpenCode's plugin ecosystem show that **the future belongs to extensible tools**. Claude Code and OpenAI Codex lack robust extension models — a gap that mid-tier tools are exploiting.

### Signal 5: Multi-agent orchestration is early but inevitable
CodeWhale's WhaleFlow swarm (synthesis/reduce, fleet ledger, parent-visible workers) represents the most ambitious multi-agent vision in the ecosystem. **The industry is moving from single-agent assistants toward agent swarms organized in DAGs.** Developers should track CodeWhale's progress as a leading indicator, even if they don't use the tool.

### Signal 6: Chinese-market tools are bridging
Kimi Code and CodeWhale are building affordances for Chinese AI providers (SiliconFlow, TokenHub, GLM-5.2). **Global tool developers ignore regional AI ecosystems at their peril** — the next DeepSeek moment will come from a provider without Western pricing.

---

**Bottom line for decision-makers:** If you need reliability today, **OpenCode** (best release discipline) or **OpenAI Codex** (best billing transparency) are safe bets. If you need extensibility, **Pi** is the most forward-looking. If you're building multi-agent systems, watch **CodeWhale** closely but expect migration friction. Avoid **Claude Code** until Anthropic addresses the recursive-spawn and billing trust crises. For enterprise GitHub shops, **Copilot CLI** is adequate but shows minimal innovation velocity.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
**Data Snapshot**: 2026-06-15 | Source: github.com/anthropics/skills

---

## 1. Top Skills Ranking

The most-discussed Skills (by PR activity/comments) reflect a community focused on **document quality**, **meta-tooling**, and **cross-platform reliability**.

### #1: Document Typography Skill (#514)
- **Functionality**: Prevents orphan word wrap, widow paragraphs, and numbering misalignment in AI-generated documents.
- **Discussion Highlights**: Identifies a universal pain point—Claude-generated documents consistently exhibit these typographic issues. The PR summary notes these affect "every document Claude generates," indicating broad user frustration.
- **Status**: OPEN — active discussion, not yet merged.
- **Link**: [PR #514](https://github.com/anthropics/skills/pull/514)

### #2: ODT Skill (#486)
- **Functionality**: OpenDocument text creation, template filling, and ODT-to-HTML conversion for LibreOffice/ISO standard formats.
- **Discussion Highlights**: Addresses demand for open-source document format support beyond DOCX/PDF. Long-running discussion (March–April 2026).
- **Status**: OPEN.
- **Link**: [PR #486](https://github.com/anthropics/skills/pull/486)

### #3: Skill Quality & Security Analyzers (#83)
- **Functionality**: Two meta-skills: *skill-quality-analyzer* evaluates Skills across five dimensions (structure, documentation, examples, etc.); *skill-security-analyzer* audits for security issues.
- **Discussion Highlights**: Represents the community's self-improvement impulse—building tools to evaluate other Skills. Five-dimensional scoring system (20% per dimension) generated substantial debate.
- **Status**: OPEN.
- **Link**: [PR #83](https://github.com/anthropics/skills/pull/83)

### #4: Agent-Creator Skill (#1140)
- **Functionality**: Meta-skill for generating task-specific agent sets. Also fixes multi-tool evaluation and adds Windows support for the recalc script.
- **Discussion Highlights**: Addresses Issue #1120 directly. Combines new feature delivery with critical stability fixes, suggesting high priority for both maintainers and users.
- **Status**: OPEN — recently updated (2026-06-02), likely close to merge.
- **Link**: [PR #1140](https://github.com/anthropics/skills/pull/1140)

### #5: Testing Patterns Skill (#723)
- **Functionality**: Comprehensive testing skill covering the Testing Trophy model, unit testing (AAA pattern), React component testing (Testing Library), and test coverage strategy.
- **Discussion Highlights**: Addresses a clear gap—there was no dedicated testing skill in the collection. The Testing Trophy vs. Testing Pyramid framing drew attention.
- **Status**: OPEN.
- **Link**: [PR #723](https://github.com/anthropics/skills/pull/723)

### #6: run_eval.py Fix (0% Recall) — Multiple PRs (#1298, #1099, #1050)
- **Functionality**: A cluster of PRs fixing the same root cause: `run_eval.py` reports `recall=0%` on every skill description, making the entire optimization loop meaningless.
- **Discussion Highlights**: Issue #556 (12 comments, 7👍) documents 10+ independent reproductions. PR #1298 (MartinCajiao) provides the most comprehensive fix—installing eval artifacts as real skills, fixing Windows stream reading, trigger detection, and parallel workers.
- **Status**: OPEN — this is the ecosystem's most critical bug.
- **Links**: [PR #1298](https://github.com/anthropics/skills/pull/1298), [PR #1099](https://github.com/anthropics/skills/pull/1099), [PR #1050](https://github.com/anthropics/skills/pull/1050)

### #7: Frontend Design Skill Improvement (#210)
- **Functionality**: Revises the existing frontend-design skill for clarity, actionability, and single-conversation usability.
- **Discussion Highlights**: Focused on making instructions "actually followable" by Claude within one conversation—a quality standard that could set a precedent for Skill design.
- **Status**: OPEN.
- **Link**: [PR #210](https://github.com/anthropics/skills/pull/210)

---

## 2. Community Demand Trends

From Issues (sorted by highest comment count), four dominant demand themes emerge:

### 🔴 Critical: Skill Testing & Validation Infrastructure (Issue #556, #1169)
**14 total comments across two issues, 8👍**
The `run_eval.py` 0% recall bug (#556) is the community's most vocal pain point. Every skill creator using the optimize loop gets misleading metrics. Issue #1169 confirms the same behavior on literal slash-command queries. This blocks the entire skill-creation pipeline.

### 🟠 High: Organizational Skill Sharing (Issue #228)
**14 comments, 7👍**
Users want to share skills within organizations without manual `.skill` file downloads and Slack forwarding. The request for a "shared skill library or direct sharing link" reflects enterprise adoption growing beyond individual users.

### 🟠 High: Security & Trust Boundaries (Issue #492)
**7 comments, 2👍**
Community skills distributed under the `anthropic/` namespace impersonate official skills. This trust-boundary vulnerability could lead to privilege escalation. Users want namespace verification or sandboxing.

### 🟡 Medium: Windows Compatibility (Issue #1061)
**3 comments**
Three distinct blockers (`PATHEXT` handling, `cp1252` encoding, `select()` on pipes) prevent Windows users from running skill-creator scripts. Multiple PRs (#1099, #1050, #1298) attempt fixes, indicating a fragmented solution landscape.

### 🟡 Medium: Document Format Expansion (Issue #189)
**6 comments, 8👍**
Duplicate content between `document-skills` and `example-skills` plugins points to demand for clearer skill organization and deduplicated document-format support (ODT, advanced PDF, etc.).

---

## 3. High-Potential Pending Skills

These PRs have active discussions and may merge soon:

| Skill | PR | Last Update | Why It's Close |
|-------|-----|-------------|----------------|
| **Agent-Creator** | [#1140](https://github.com/anthropics/skills/pull/1140) | 2026-06-02 | Fixes Issue #1120, includes Windows fix, multi-tool evaluation fix. High-value meta-skill. |
| **run_eval.py Comprehensive Fix** | [#1298](https://github.com/anthropics/skills/pull/1298) | 2026-06-11 | Addresses the #1 community bug (#556). Author MartinCajiao provides a holistic fix. |
| **Windows subprocess + encoding fixes** | [#1050](https://github.com/anthropics/skills/pull/1050) | 2026-05-24 | Minimal 1-line changes targeting Issue #1061. Low risk, high impact. |
| **YAML description validation** | [#539](https://github.com/anthropics/skills/pull/539) | 2026-04-16 | Prevents silent YAML parsing failures from unquoted `:` characters. Essential for skill reliability. |
| **Testing Patterns Skill** | [#723](https://github.com/anthropics/skills/pull/723) | 2026-04-21 | Fills a clear gap in the Skills collection with comprehensive coverage. |

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is for *reliable meta-tooling*—fixing the skill-creation and evaluation pipeline (run_eval.py, Windows compatibility, YAML validation) so that authors can trust their Skill optimization loops, rather than for any single new Skill category.**

The volume of discussion around `run_eval.py` 0% recall (three PRs, two issues, 10+ reproductions) dwarfs all feature requests. Until the evaluation infrastructure is stable, the entire ecosystem of community-contributed Skills operates without reliable quality feedback.

---

# Claude Code Community Digest — 2026-06-15

---

## Today's Highlights

The community is rallying around a long-standing **India-specific pricing request** (#17432) that has amassed 442 upvotes and 194 comments, now spanning over five months without official response. Meanwhile, a **critical subagent recursion bug** (#68430) reported yesterday describes exponential token burn with agents spawning 50+ levels deep, and a **kernel-level memory leak on macOS 26.5.1** (#66020) is causing `claude.exe` to panic at ~20GB RAM usage — both demanding urgent attention from the Anthropic team.

---

## Releases

No new releases in the last 24 hours.

---

## Hot Issues

### 1. India-Specific Pricing Plans — #17432
**Link:** https://github.com/anthropics/claude-code/issues/17432  
**Signal:** 442 👍 | 194 comments | Open since Jan 2026  
The most-voted issue on the repository. Users request INR-denominated pricing for Claude Pro and Claude Code, citing competitors (OpenAI, Google) already offering regional pricing. The high comment count reflects sustained community frustration over currency conversion costs and lack of local payment methods. No official response from Anthropic in over five months — a growing reputational risk.

### 2. Cowork Edit/Write Tools Silently Truncate Files — #53940
**Link:** https://github.com/anthropics/claude-code/issues/53940  
**Signal:** 12 👍 | 31 comments  
A deterministic bug where tools truncate file content due to a byte-conservation buffer cap. This affects all file sizes and can cause silent data loss during edit operations. The detailed reproduction steps have made this a community-pinned debugging reference.

### 3. Session Cleanup Ignoring `cleanupPeriodDays: 99999` — #41458
**Link:** https://github.com/anthropics/claude-code/issues/41458  
**Signal:** 1 👍 | 16 comments  
Despite an explicit configuration to never auto-clean, 490 sessions were silently deleted. Labelled as a regression with data-loss severity. Users report this as a trust-breaking bug — you cannot rely on session persistence even with explicit settings.

### 4. Extra Usage Charged Despite Available Plan Capacity — #32544
**Link:** https://github.com/anthropics/claude-code/issues/32544  
**Signal:** 14 👍 | 15 comments  
Linux users report being billed for usage beyond plan limits even when unused capacity remains, combined with false rate-limit errors. This creates a double pain point: overcharging and blocked productivity. The cross-platform nature (tagged Linux, but affects cost/auth) suggests a systemic billing issue.

### 5. Windows Desktop Persistent Blank/White Screen — #51143
**Link:** https://github.com/anthropics/claude-code/issues/51143  
**Signal:** 12 👍 | 13 comments  
Cowork mode is entirely unusable on Windows due to a persistent blank/white screen. Multiple reinstalls have no effect. This has been open for nearly two months, blocking an entire platform.

### 6. Bash Tool Calls Emitted as Raw `<invoke>` Text — #63870
**Link:** https://github.com/anthropics/claude-code/issues/63870  
**Signal:** 13 👍 | 11 comments  
The model emits raw XML invocation tags instead of executing them, causing silent failures. The reporter provides JSONL evidence from 23 malformed calls. Multiple duplicates suggest this is a systemic model-side regression, not an isolated incident.

### 7. Copy-Paste Broken on macOS — #66192
**Link:** https://github.com/anthropics/claude-code/issues/66192  
**Signal:** 10 👍 | 11 comments  
A basic UX regression — copy-paste does not work in the TUI on macOS. This is a fundamental workflow blocker for any developer relying on text manipulation within the tool.

### 8. CRITICAL: Subagent Infinite Recursion — #68430
**Link:** https://github.com/anthropics/claude-code/issues/68430  
**Signal:** 0 👍 | 7 comments (filed yesterday)  
Multiple compound regressions: subagents spawn recursively 50+ levels deep, ignore `CLAUDE_CODE_FORK_SUBAGENT=0`, and fetch individual files via HTTP instead of using the workspace. The author describes this as a "catastrophic token burn scenario." Highest urgency among all open issues.

### 9. macOS Kernel Zone Memory Leak — #66020
**Link:** https://github.com/anthropics/claude-code/issues/66020  
**Signal:** 0 👍 | 7 comments  
A `data.kalloc.1024` zone leak in the macOS kernel attributed to Claude Code CLI. Leak rate scales from 21 to 1027/sec under agent load, causing kernel panics at ~20GB. This is a platform-level stability issue that could crash the entire system.

### 10. Pty File Descriptor Leak — #65995 / #66434  
**Links:** https://github.com/anthropics/claude-code/issues/65995 | https://github.com/anthropics/claude-code/issues/66434  
**Signal:** 2-3 comments each  
Multiple reports confirm that Claude Desktop leaks ptmx file descriptors on macOS, eventually exhausting the system pty pool and killing all terminals with `forkpty: Device not configured`. A system-wide DoS caused by a client application.

---

## Key PR Progress

### 1. #43598 — Add upstream issue sync workflow (CLOSED)
**Link:** https://github.com/anthropics/claude-code/pull/43598  
A script and documentation for fetching and normalizing upstream issues from `anthropics/claude-code`, with robust GitHub CLI pagination handling. Merged — improves issue triage infrastructure.

### 2. #68423 — Don't auto-close assigned issues in sweep (OPEN)
**Link:** https://github.com/anthropics/claude-code/pull/68423  
Fixes `scripts/sweep.ts` where `markStale` correctly skips assigned issues but `closeExpired` does not — leads to auto-closing issues actively owned by someone. A small but critical workflow fix for maintainers.

### 3. #67699 — Fix for autonomous background script execution (OPEN)
**Link:** https://github.com/anthropics/claude-code/pull/67699  
Bounty: $29. Addresses a security/privacy concern where Claude autonomously runs background scripts that call paid external APIs. Implemented via NVIDIA AI automation. Community-funded fix.

### 4. #67409 — Fix for account downgrade due to billing error (OPEN)
**Link:** https://github.com/anthropics/claude-code/pull/67409  
Bounty: $200. The highest-bountied open PR. Addresses incorrect account downgrades caused by billing system errors. Also NVIDIA AI-driven implementation. Significant community investment in billing reliability.

### 5. #67722 — Claude deduplication workflow (CLOSED)
**Link:** https://github.com/anthropics/claude-code/pull/67722  
A workflow for generating and deduplicating issues automatically. Approved and merged. Aims to reduce duplicate bug reports, which are currently flooding the issue tracker.

---

## Feature Request Trends

### 1. Regional Pricing & Localization
The overwhelming theme. #17432 (India INR pricing) dominates, but there are also requests for localized timezone handling (#64988) and project-scoped views (#68495). The community sees regional pricing as table-stakes for global adoption and is vocal about Anthropic lagging behind OpenAI and Google.

### 2. Agent & Subagent Configuration Control
Multiple requests ask for **per-message model selection** (#68165), **cwd parameter for Task tool** (#12748), and **model/effort inheritance in agent teams** (#68411). Users want fine-grained control over how subagents are spawned, what models they use, and what resources they can access. The infinite recursion bugs (#68430, #68110) amplify this demand — the system needs guardrails.

### 3. UI/UX Improvements
- **Appshots-style window capture** (#68498) — inspired by OpenAI Codex, this would capture full window text including scrolled content via macOS accessibility APIs.
- **Project-scoped conversation view** (#68495) — a regression complaint that the new home screen shows conversations from all projects, leaking context.
- **Per-message model switching** (#68165) — users want to use cheap models for simple lookups and expensive models for complex reasoning within the same session.

### 4. Status & Billing Transparency
Requests for richer statusline hooks (#62082), clearer billing for `remote-control` mode (#59823), and granular rate-limit visibility. Users want to understand what they're being charged for and when they're approaching limits before hitting errors.

---

## Developer Pain Points

### 1. Billing & Usage Ambiguity
This is the #1 systemic frustration. Users report being charged for usage when they have available capacity (#32544), unclear billing for `remote-control` mode (#59823), and missing rate-limit data in status hooks (#62082). The India pricing issue (#17432) amplifies this with the added pain of USD-only conversion costs. **The billing system lacks transparency and users feel nickel-and-dimed.**

### 2. Silent Data Loss & Configuration Non-Compliance
Multiple bugs where the tool ignores explicit user configuration:
- `cleanupPeriodDays: 99999` ignored — sessions deleted (#41458)
- Files silently truncated by Cowork tools (#53940)
- Auto-compact stopped for third-party providers (#65585)
- Model/effort settings ignored by subagents (#68411)

**Developers cannot trust that their configuration will be respected**, which undermines the entire value proposition of a deterministic dev tool.

### 3. Unbounded Resource Consumption
The subagent recursion bugs (#68430, #68110) describe exponential token burn with no depth limits. Combined with the macOS kernel memory leak (#66020) and pty exhaustion (#65995), there is a pattern of **insufficient resource governance**. The tool can crash not just itself but the entire operating system.

### 4. Windows & Linux Second-Class Support
- Windows: blank screen (#51143), file lock errors (#51847), MCP integration issues (#68462)
- Linux: extra billing (#32544), false rate limits (#68502), rootfs checksum errors (#68514)

**macOS receives the most attention, while developers on other platforms experience persistent, unremediated bugs.** The Windows blank screen issue has been open for nearly two months with no fix.

### 5. Rate Limiting That Doesn't Distinguish
Issue #68502 highlights a critical UX failure: HTTP 529 overload errors are rendered as "Rate limited" even though the problem is server-side capacity, not user quota. This leads developers to believe they're being throttled when in fact Anthropic's infrastructure is struggling. **Misleading error messages waste developer time and erode trust.**

---

*Data sourced from github.com/anthropics/claude-code issues and PRs updated as of 2026-06-15 23:59 UTC.*

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest – 2026-06-15

## Today’s Highlights

The Codex community continues to experience a mix of regression bugs and feature demands. Most notably, a long‑standing enhancement request to **rename task/thread titles** (Issue #12564) was closed after 80 comments and 111 upvotes, signaling strong user desire for better session organisation. Meanwhile, the **Windows desktop app** is the source of multiple crash reports following the June 12 update, affecting both WSL mode and standard launches, with several issues accumulating over 20 comments. On the development side, OpenAI engineers are shipping improvements to **rate‑limit credit management**, **async hook execution**, and **MITM CA handling** for child sandboxes.

## Releases

No new versions were published in the last 24 hours.

## Hot Issues

*(10 noteworthy items, ordered by community engagement)*

| Issue | Summary | Comments | 👍 | Why it matters |
|-------|---------|----------|---|----------------|
| [#12564](https://github.com/openai/codex/issues/12564) (CLOSED) | **Allow renaming task/thread titles** to improve history navigation | 80 | 111 | Most upvoted issue ever. Users want to organise long sessions – closed as enhancement. |
| [#27979](https://github.com/openai/codex/issues/27979) | **Windows Codex App 26.609.4994.0 no longer opens** after update | 21 | 6 | Critical regression affecting Windows Pro users; app crashes on launch. |
| [#25500](https://github.com/openai/codex/issues/25500) | **Project sidebar shows “No chats”** for older conversations after update | 18 | 2 | Chat history disappearance is a top frustration; affects productivity. |
| [#23840](https://github.com/openai/codex/issues/23840) | **Computer Use MCP initialize times out** from desktop but works from terminal | 9 | 0 | Highlights inconsistency between desktop and CLI MCP handshakes. |
| [#27367](https://github.com/openai/codex/issues/27367) | **Windows 10 app immediately exits** after update; CLI works fine | 9 | 0 | Another Windows crash, specifically on 22H2. |
| [#27353](https://github.com/openai/codex/issues/27353) | **Project chat history disappeared** after Codex app update (macOS) | 7 | 3 | Cross‑platform issue with session persistence. |
| [#15281](https://github.com/openai/codex/issues/15281) | **Expose full usage/limits data** in CLI `/status` command | 6 | 15 | Users want real‑time rate‑limit visibility; high demand for CLI transparency. |
| [#28180](https://github.com/openai/codex/issues/28180) (CLOSED) | **Remotion causes 100% CPU** from syspolicyd/trustd on macOS | 5 | 0 | Performance bug tied to Chrome rendering – closed quickly by team. |
| [#28103](https://github.com/openai/codex/issues/28103) | **MSIX build missing Linux `codex` binary** for “Run agent in WSL” | 5 | 9 | Blocks WSL workflow for Windows users; 9 upvotes indicate frustration. |
| [#25431](https://github.com/openai/codex/issues/25431) | **Expose spellcheck toggle** in Windows Desktop settings | 5 | 14 | Simple UX request, high upvote count; lacking user control. |

## Key PR Progress

*(10 important pull requests updated in the last 24h)*

| PR | What it does | Impact |
|----|-------------|--------|
| [#25888](https://github.com/openai/codex/pull/25888) | Prepare managed child MITM CA environment | Parent PR in a stack to securely proxy HTTPS for sandboxed agents. |
| [#28008](https://github.com/openai/codex/pull/28008) | Add external agent import result accounting | Gives clients a stable import ID for tracking async plugin/session imports. |
| [#26315](https://github.com/openai/codex/pull/26315) | Materialise child MITM CA bundles | Makes child‑selected CA material readable and immutable – critical for managed sandboxes. |
| [#27963](https://github.com/openai/codex/pull/27963) | Reference writable roots from environment context | Deduplicates path info in permissions messages; cleaner UX for developer permissions. |
| [#28143](https://github.com/openai/codex/pull/28143) | Expose rate‑limit reset credits in app‑server API | Backend foundation for redeeming personal reset credits – paired with TUI PR #28154. |
| [#27640](https://github.com/openai/codex/pull/27640) | Support multi‑tool install requests | Lets the model request multiple plugin installs in one operation; improves efficiency. |
| [#28235](https://github.com/openai/codex/pull/28235) | Add `request_user_input` auto‑resolution timer | TUI auto‑answers after 2 minutes of inactivity – reduces stuck sessions. |
| [#28154](https://github.com/openai/codex/pull/28154) | Add rate‑limit reset redemption to `/usage` | CLI users can now view and redeem personal reset credits directly from the TUI. |
| [#28232](https://github.com/openai/codex/pull/28232) | Add workspace headline statusline item | Enterprise‑gated feature to show workspace messages in the TUI status bar. |
| [#27794](https://github.com/openai/codex/pull/27794) | Remove terminal resize reflow flag gates | Stabilises terminal reflow; old config no longer needed. |

## Feature Request Trends

Several high‑demand features appear repeatedly across issues:

- **Session/thread organisation** – Ability to rename tasks, group conversations, and recover lost history (Issues #12564, #25500, #27353).
- **Rate‑limit transparency** – Expose detailed usage, quotas, and reset credits in both CLI and desktop (Issues #15281, #28246, #28249).
- **Better MCP tool call feedback** – Users want visible progress indicators, cancellation signals, and timeout handling (Issues #28003, #26956, #23840).
- **Customisation of UI/UX** – Spellcheck toggle, longer response scrolling, configurable status bar (Issues #25431, #23280, #28232).
- **Cross‑platform parity** – Windows users consistently ask for the same level of support as macOS (e.g., WSL binary, spellcheck, sandbox stability).

## Developer Pain Points

Recurring frustrations and high‑frequency bug reports:

- **Windows app instability** – Multiple crash reports after the June 12 update (#27979, #27367, #28245), plus missing Linux binary for WSL (#28103) and overlapping UI elements (#28243).
- **Session and project history loss** – Updates or power outages cause chats to “disappear” or show incorrect sandbox permissions (#25500, #27353, #25590, #28248).
- **MCP reliability** – Handshake timeouts, missing stop signals, and ignored OAuth keyring persistence (#23840, #26956, #28201).
- **Performance issues** – Zombie process leaks on macOS (#28244), 100% CPU from system daemons (#28180), and quota window anchoring bugs that lose subscription time (#28246).
- **Plugin/sandbox lifecycle** – Stale subagents that cannot be closed (#25179), plugin reinstall loop on Windows (#28247), and missing metric sanitisation for skills with colons (#27659).

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — June 15, 2026

## Today’s Highlights
No new releases landed in the past 24 hours, but the repository saw heavy activity with 50 issues and 43 PRs updated. A major dependency batch update (53 packages) was merged, and several high-priority bugs around agent hangs, subagent recovery, and telemetry errors remain open. The community continues to push for tighter agent safety, better code awareness via AST tools, and more reliable auto-memory handling.

## Releases
No new versions were published in the last 24 hours.

---

## Hot Issues
*(Top 10 by comment count and community engagement)*

1. **#21409 – Generalist agent hangs**  
   [Link](https://github.com/google-gemini/gemini-cli/issues/21409)  
   P1 bug with 8 👍. Simple tasks like folder creation cause indefinite hangs; users must explicitly disable sub‑agent delegation to work around it. High frustration.

2. **#24353 – Robust component‑level evaluations**  
   [Link](https://github.com/google-gemini/gemini-cli/issues/24353)  
   Epic tracking 76 behavioral eval tests across 6 Gemini models. Aims to formalize evaluation infrastructure for agent quality.

3. **#22745 – Assess impact of AST‑aware file reads, search, and mapping**  
   [Link](https://github.com/google-gemini/gemini-cli/issues/22745)  
   P2 investigation into whether AST‑aware tools can reduce token waste and improve agent accuracy. Linked to #22746 and #22747.

4. **#22323 – Subagent recovery after MAX_TURNS reports false success**  
   [Link](https://github.com/google-gemini/gemini-cli/issues/22323)  
   P1 bug: `codebase_investigator` subagent says “GOAL” even after hitting turn limits, hiding real failures. 2 👍 from affected users.

5. **#21968 – Gemini does not use custom skills and sub‑agents enough**  
   [Link](https://github.com/google-gemini/gemini-cli/issues/21968)  
   P2 anecdotal report that user‑defined skills (e.g., “gradle”, “git”) are almost never invoked unless explicitly forced.

6. **#26525 – Add deterministic redaction and reduce Auto Memory logging**  
   [Link](https://github.com/google-gemini/gemini-cli/issues/26525)  
   Security concern: transcript content is sent to models before redaction; logging may expose secrets. P2 – part of memory quality push.

7. **#26522 – Stop Auto Memory from retrying low‑signal sessions indefinitely**  
   [Link](https://github.com/google-gemini/gemini-cli/issues/26522)  
   Sessions deemed low‑signal remain unprocessed and keep being surfaced, causing infinite retries. P2 bug.

8. **#25166 – Shell command execution gets stuck with “Waiting input” after completion**  
   [Link](https://github.com/google-gemini/gemini-cli/issues/25166)  
   P1 with 3 👍. After simple commands, the CLI hangs showing “Awaiting user input” even though the command finished.

9. **#21983 – Browser subagent fails on Wayland**  
   [Link](https://github.com/google-gemini/gemini-cli/issues/21983)  
   P1 bug: browser agent terminates with “GOAL” but never actually renders on Wayland‑based systems.

10. **#20079 – Symlinked agent files not recognized**  
    [Link](https://github.com/google-gemini/gemini-cli/issues/20079)  
    P2: `~/.gemini/agents/filename.md` symlinks are ignored; users cannot organise agents with links.

---

## Key PR Progress
*(Highlighting fixes, features, and major dependency upgrades)*

1. **#27729 – Fix telemetry metric attribute truncation**  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27729)  
   P2. Truncates attributes to 1024 chars to prevent GCP export errors that flooded terminal with stack traces. (Open)

2. **#27730 – Keep array tool results out of `structuredContent`**  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27730)  
   P1 fix for MCP compliance: JSON arrays from calendar‑style payloads no longer break tool output. (Open)

3. **#27718 – Keep `auto` model alias visible without preview access**  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27718)  
   P2. Restores the top‑level `auto` alias in `/model` when dynamic model config is enabled, while filtering preview‑only aliases. (Open)

4. **#23030 – Non‑invasive UX Journey testing framework**  
   [Link](https://github.com/google-gemini/gemini-cli/pull/23030)  
   Closed. Introduces “White Box” terminal UI testing for React components – enables automated verification without manual instrumentation.

5. **#22456 – New interactive policies dialog**  
   [Link](https://github.com/google-gemini/gemini-cli/pull/22456)  
   Closed. Replaces text‑based `/policies` output with a searchable, tabbed dialog categorising Allow/Ask/Deny rules.

6. **#27925 – Bulk npm dependencies update (53 packages)**  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27925)  
   Closed. Includes major bumps for `@google/genai`, `undici`, `puppeteer-core`, `google-auth-library`, and more. Keeps the CLI on modern runtimes.

7. **#27929 – Bump `@google/genai` from 1.30.0 to 2.8.0**  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27929)  
   Major version step that likely brings new GenAI SDK capabilities.

8. **#27931 – Bump `puppeteer-core` from 24.39.0 to 25.1.0**  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27931)  
   Critical for the browser agent – may fix some Wayland‑related issues.

9. **#27928 – Bump `undici` from 7.24.5 to 8.4.0**  
   [Link](https://github.com/google-gemini/gemini-cli/pull/27928)  
   HTTP client upgrade – likely includes connection handling improvements.

10. **#27926 – Bump `google-auth-library` from 9.15.1 to 10.7.0**  
    [Link](https://github.com/google-gemini/gemini-cli/pull/27926)  
    Major update for authentication backend, needed for enterprise/remote agent features.

---

## Feature Request Trends
Several high‑level themes emerge from recent issues:

- **AST‑aware code understanding** — Multiple issues (#22745, #22746, #22747) propose using AST tools for file reads, search, and codebase mapping to reduce token waste and improve agent precision.
- **Evaluation infrastructure** — Strong push for robust, stable component‑level evals (#24353, #23166) to track quality regressions and ensure test reliability.
- **Agent self‑awareness & tool usage** — Users want the CLI to automatically invoke custom skills and sub‑agents (#21968, #21432) and to expose accurate metadata about its own capabilities.
- **Auto‑memory & session management** — Improvements to memory redaction (#26525), avoidance of infinite retry loops (#26522), and better patch validation (#26523) are all in flight.
- **Safety & non‑destructive behavior** — Requests for the agent to avoid `git reset --force` and destructive DB operations (#22672) highlight a need for better guardrails.

---

## Developer Pain Points
Recurring frustrations reported by the community:

- **Agent hangs** – The generalist agent hangs on simple tasks (#21409), shell commands stall after completion (#25166), and the “get‑shit‑done” output hook crashes (#22186).
- **Sub‑agent unreliability** – False success reports after hitting turn limits (#22323), browser agent failures on Wayland (#21983), and ignored `settings.json` overrides (#22267).
- **Permission & detection issues** – Sub‑agents run without consent after updates (#22093), symlinked agent files are ignored (#20079), and the CLI fails with 400 errors when >128 tools are available (#24246).
- **Terminal & UI glitches** – Terminal corruption after exiting external editors (#24935), poor resize performance (#21924), and stray temp scripts cluttering workspaces (#23571).
- **Memory system overhead** – Auto‑memory retries low‑signal sessions forever (#26522), and secrets may be exposed before redaction (#26525).

These pain points are driving ongoing work in agent stability, resource management, and user‑control over tool execution.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-06-15

## Today’s Highlights
No new releases or pull requests landed in the last 24 hours. The community surfaced a new UI inconsistency bug (#3797) and raised two feature requests—model discovery for BYOK customers (#3795) and Azure DevOps work item integration (#3794). Two long-running issues, duplicate item errors (#3558) and agent skills folder execution (#956), continue to attract attention and comments.

## Releases
No new versions were published in the past 24 hours.

## Hot Issues
Only six actionable issues were updated in the last 24 hours; the most noteworthy are highlighted below.

**#956 – [area:agents] Agent skills scripts executed in wrong folder**  
Author: msundman78 · Updated: Jun 14 · Comments: 6 · 👍: 2  
A skill defined with a script reference (e.g. `scripts/myscript.sh`) is executed from the wrong working directory, breaking the [Agent Skills spec](https://agentskills.io/specification#file-references). Community discussion continues about proper path resolution.  
[Issue #956](https://github.com/github/copilot-cli/issues/956)

**#3558 – [area:context-memory, area:models] Duplicate Item Errors**  
Author: psulightning · Updated: Jun 14 · Comments: 4 · 👍: 7  
After the initial prompt, users receive a `CAPIError: 400` with `"Duplicate item found with id fc_call_..."`. The error persists across sessions and appears to be a server-side deduplication issue. High reaction count indicates broad impact.  
[Issue #3558](https://github.com/github/copilot-cli/issues/3558)

**#3797 – [triage] Different prompt input box layout in two cmd tabs in the same window**  
Author: kunalk16 · Updated: Jun 15 · Comments: 1 · 👍: 0  
A UI rendering bug where two command tabs within the same window display different prompt input box layouts, causing confusion for multi-tab users.  
[Issue #3797](https://github.com/github/copilot-cli/issues/3797)

**#3795 – [triage] Feature request: opt-in model discovery for BYOK / custom providers**  
Author: aosama · Updated: Jun 14 · Comments: 0 · 👍: 0  
When using a custom provider (BYOK mode), users must manually set `COPILOT_MODEL` or pass `--model`. The CLI does not query the provider for available models, making configuration error-prone.  
[Issue #3795](https://github.com/github/copilot-cli/issues/3795)

**#3794 – [triage] Add Azure DevOps work items to Up next**  
Author: OmerMicro · Updated: Jun 14 · Comments: 0 · 👍: 0  
The “Up next” panel currently only surfaces GitHub issues and PRs. For projects backed by Azure DevOps repos, the panel is empty. Request to also surface assigned ADO work items.  
[Issue #3794](https://github.com/github/copilot-cli/issues/3794)

**#3791 – [triage] Malformed attachment poisons session; all subsequent turns fail with 400**  
Author: jay‑tau · Updated: Jun 14 · Comments: 0 · 👍: 0  
A password-protected `.xlsx` file (or any unsupported attachment) triggers a CAPI 400 error that persists for **all** future turns in the same session, even after removing the attachment. This is a session‑state bug that forces a full restart.  
[Issue #3791](https://github.com/github/copilot-cli/issues/3791)

*Two other issues (#3796 closed as invalid, #3793 with no description) were not actionable and are excluded.*

## Key PR Progress
No pull requests were updated in the last 24 hours.

## Feature Request Trends
The most-requested feature directions emerging from recent issues are:

- **Model discovery for BYOK / custom providers** — Users want automatic querying of model identifiers instead of manual `COPILOT_MODEL` configuration. (#3795)
- **Azure DevOps integration in “Up next”** — Extending the cross-session inbox to include Azure DevOps work items for teams using ADO as their project backend. (#3794)
- **Agent skills execution context** — Clearer guarantees about the working directory when skills run scripts, to match the official Agent Skills specification. (#956)

## Developer Pain Points
Two high‑friction patterns are evident:

1. **Session‑poisoning bugs** — Once a malformed attachment (e.g., encrypted `.xlsx`) causes a 400 error, the CLI session becomes irrecoverable; all subsequent turns fail with the same error. (#3791)
2. **Duplicate item errors** — A server-side duplicate ID check (`fc_call_*`) fires even when no obvious duplication exists, forcing users to restart sessions repeatedly. (#3558) The high reaction count (👍7) suggests this is a frequent blocking issue.

---

*Digest generated from [github.com/github/copilot-cli](https://github.com/github/copilot-cli) activity on 2026-06-15.*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-06-15

## Today’s Highlights
This week’s digest is light on new releases but heavy on community feedback. Three issues and four pull requests were updated in the last 24 hours, highlighting growing developer concerns around rate limiting and system prompt inflexibility. Meanwhile, contributions from the community continue to improve Windows support, with two closed PRs for terminal paste and log file handling, and a new open fix for multi-edit file replacement.

## Releases
No new releases in the last 24 hours.

## Hot Issues (3 items)

1. **[#850 – Auto-load project context/rules](https://github.com/MoonshotAI/kimi-cli/issues/850)** (CLOSED)  
   *Author: Al4ric*  
   A feature request to replicate Claude Code’s automatic loading of `CLAUDE.md`-style project rules. The user switched from Claude Code and expects similar seamless context injection. The issue was closed, possibly implemented or deferred. Community reaction was muted (1 👍, 3 comments).

2. **[#2123 – 限速，限额严重 (Severe rate limiting)](https://github.com/MoonshotAI/kimi-cli/issues/2123)** (OPEN)  
   *Author: littlePoBoy*  
   A critical complaint about aggressive rate limiting on the “Code Plan” subscription, which promises 300–1200 requests per 5 hours but delivers only ~60+ in practice. The user calls out unclear disclosure of quotas, and mentions being denied a refund. This is a high-impact pain point for paying users and has attracted 2 comments. The issue remains open, suggesting the team is still evaluating or addressing it.

3. **[#2451 – System prompt conflicting with desired workflow](https://github.com/MoonshotAI/kimi-cli/issues/2451)** (OPEN)  
   *Author: iaindooley*  
   A bug report about the built-in system prompt overriding user-defined workflow guidelines. The user uses an API key with `k2.7-coding` model and has strict rules in a separate file, but Kimi Code appears to inject its own instructions. Zero comments yet, but it signals a need for more flexible prompt configuration.

## Key PR Progress (4 items)

1. **[#2452 – fix(tools): fail StrReplaceFile when a multi-edit hunk is unmatched](https://github.com/MoonshotAI/kimi-cli/pull/2452)** (OPEN)  
   *Author: Osamaali313*  
   Fixes a bug where `StrReplaceFile` only errors if the entire result is unchanged, instead of reporting partial match failures. This improves reliability of multi-edit operations. No comments yet; likely still under review.

2. **[#2018 – feat: add Alt+V paste support for Windows Terminal](https://github.com/MoonshotAI/kimi-cli/pull/2018)** (CLOSED, merged)  
   *Author: LittleDrinks*  
   Adds Alt+V as an alternative paste keybinding because Windows Terminal intercepts Ctrl+V. A practical Windows UX improvement that has been merged.

3. **[#2020 – fix: use per-process log filenames to prevent rotation lock on Windows](https://github.com/MoonshotAI/kimi-cli/pull/2020)** (CLOSED, merged)  
   *Author: LittleDrinks*  
   Resolves `PermissionError` when multiple Kimi processes run concurrently by naming log files `kimi.{pid}.log`. Also merged.

4. **[#839 – feat(shell): add configurable shell support for Windows](https://github.com/MoonshotAI/kimi-cli/pull/839)** (CLOSED)  
   *Author: HamzaETTH*  
   Adds the ability to configure the shell used by Kimi Code on Windows. Merged, expanding platform flexibility.

## Feature Request Trends
The most requested feature direction from recent issues is **automatic project context loading** (e.g., from `AGENTS.md` or `.cursorrules`), as seen in #850. Users migrating from competing tools expect zero-config onboarding. A secondary trend is demand for **more transparent and generous rate limits** (#2123), particularly for paying subscribers. The system prompt control issue (#2451) also hints at a desire for user-defined instructions to take precedence.

## Developer Pain Points
The dominant pain point is **rate limiting** (#2123) — the gap between advertised capacity and real-world usage is causing frustration, refund disputes, and potential churn. A secondary but recurring frustration is **lack of control over the system prompt** (#2451), which breaks established workflows when users have strict guidelines. Both issues remain open, so the community is watching for official responses and fixes.

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest – 2026-06-15

## Today’s Highlights
The team shipped **v1.17.7** with several core bug fixes (reusing active server for plugin clients, fixing ACP shell tool output, and better PTY environment support). Meanwhile, the community flagged two critical regressions: the TUI hangs on a black screen since v1.17.0 (issue #32361) and the desktop UI gets stuck indefinitely on “thinking” after a stream error (#32366). A security vulnerability was also opened disclosing that MCP server subprocesses receive the full `process.env` (#31778).

## Releases
**v1.17.7** (latest)  
- *Bug fixes:* Plugin client requests now reuse the active server instead of assuming the default local port. ACP shell tool calls display the command and working directory from the start. Plugin-provided shell environment variables now apply to PTY sessions.  
- *Improvements:* MCP-related improvements rolled in (details in changelog).  
[View release](https://github.com/anomalyco/opencode/releases/tag/v1.17.7)

## Hot Issues (10 notable)
1. **[#13984 – Can not copy and paste in OpenCode CLI](https://github.com/anomalyco/opencode/issues/13984)**  
   Long-standing usability blocker (48 comments, 20 👍). Copy appears to succeed but paste yields nothing. Still open, likely requires terminal/clipboard integration overhaul.

2. **[#32366 – UI stuck on ‘thinking’ indefinitely after stream error](https://github.com/anomalyco/opencode/issues/32366)**  
   Critical usability bug: desktop UI freezes with no error display after an `AI_APICallError` or socket close. No state recovery; requires app restart. High impact.

3. **[#32361 – TUI hangs on black screen at boot since v1.17.0](https://github.com/anomalyco/opencode/issues/32361)**  
   Regression on Ubuntu/Wayland – terminal freezes silently. v1.16.2 last working version. 1 👍, one reporter but likely affects many.

4. **[#31778 – MCP server subprocess receives full process.env (API keys leaked)](https://github.com/anomalyco/opencode/issues/31778)**  
   Security issue: `process.env` is passed wholesale to every local MCP server. Exposes API keys, tokens, etc. Open and flagged as high priority.

5. **[#31072 – Subagent sessions fail to project first message due to race in commitSyncEvent](https://github.com/anomalyco/opencode/issues/31072)**  
   Subtle race condition: first message event for a subagent session is lost, leading to an empty session row. Affects v1.16.2 and earlier.

6. **[#31487 – Duplicate overlapping LLM streams in one session after background task injection](https://github.com/anomalyco/opencode/issues/31487)**  
   Concurrency bug: distinct assistant messages/parts appear for the same parent user message, persisted and not UI-only. 1 👍.

7. **[#32172 – [FEATURE] Add GLM-5.2 model support for Z.AI provider](https://github.com/anomalyco/opencode/issues/32172)**  
   Request to integrate the newly released GLM-5.2 reasoning model. 7 comments, open for two days.

8. **[#32369 – [FEATURE] Allow project-level tool description augmentation for MCP/built-in tools](https://github.com/anomalyco/opencode/issues/32369)**  
   Users want to modify tool descriptions per project to influence model selection, e.g., preferring CodeGraph over `grep`.

9. **[#32368 – [FEATURE] Make compaction revertible (undo / explicit dialog)](https://github.com/anomalyco/opencode/issues/32368)**  
   Closed as implemented? Request for compaction to be undoable, either via dialog or manual revert. Community response positive.

10. **[#28202 – Plugin async prompts can overlap with Web prompt_async and create duplicate siblings](https://github.com/anomalyco/opencode/issues/28202)**  
    Closed! Bug fixed where overlapping async prompts produced duplicate assistant messages. 4 👍, important for plugin stability.

## Key PR Progress (10 important)
1. **[#30977 – feat(tui): attach to configured server by default](https://github.com/anomalyco/opencode/pull/30977)**  
   New TUI config `server.attach` allows the TUI to connect to a pre-configured server automatically. Closes #17322.

2. **[#32364 – fix: reset terminal modes on TUI shutdown](https://github.com/anomalyco/opencode/pull/32364)**  
   Ensures terminal state is properly restored on exit, fixing issues with broken subsequent terminal sessions.

3. **[#32373 – feat(opencode): support models.dev reasoning options](https://github.com/anomalyco/opencode/pull/32373)**  
   Adds `reasoning_options` to the models.dev provider schema, generating proper model variants for reasoning effort levels.

4. **[#32241 – fix(tui): render move errors inline](https://github.com/anomalyco/opencode/pull/32241)**  
   Improves error handling for file move operations by showing an inline error row with recovery guidance instead of a broken dialog.

5. **[#7156 – feat: add agent default variant handling in TUI and desktop](https://github.com/anomalyco/opencode/pull/7156)**  
   Long-running PR (since January) that respects an agent’s configured model variant. Important for multi-model workflows.

6. **[#9545 – feat(usage): unified usage tracking with auth refresh](https://github.com/anomalyco/opencode/pull/9545)**  
   Large feature adding `Usage.Service` for OAuth providers (Anthropic, GitHub Copilot, OpenAI). Exposes `/usage` API endpoint.

7. **[#8535 – feat(session): bi-directional cursor-based pagination](https://github.com/anomalyco/opencode/pull/8535)**  
   Enables efficient forward/backward pagination for session messages across server, app, and TUI.

8. **[#31132 – fix(tui): load root sessions safely in dialogs](https://github.com/anomalyco/opencode/pull/31132)**  
   Fixes multiple session dialog crashes by safely loading root sessions. Closes #16270, #31125, and part of #13877.

9. **[#32370 – Linux clipboard selection](https://github.com/anomalyco/opencode/pull/32370)**  
   Implements PRIMARY buffer support for Linux terminal selection, fixing #29963. New feature.

10. **[#32245 – fix(mcp): stop idle OAuth callback server](https://github.com/anomalyco/opencode/pull/32245)**  
    Ensures the MCP OAuth callback listener shuts down properly after authentication, releasing the port. Closes related issues.

## Feature Request Trends
- **Session lifecycle management** – rename sessions (#32375), compaction revertibility (#32368), session title in footer (#32372).
- **Model/provider flexibility** – adding new models (GLM-5.2), reasoning options, and per-project tool description augmentation to influence model selection (#32172, #32373, #32369).
- **MCP/ACP cleanup** – ensuring temporary resources (OAuth servers, MCP registrations) are cleaned up after sessions close (#32371, #32245).
- **Linux-first features** – clipboard selection (#32370) and terminal mode safety (#32364) reflect growing Linux user base.

## Developer Pain Points
- **Terminal/CLI regressions** – Copy-paste not working (#13984, 48 comments), TUI black screen hang (#32361), and terminal mode not resetting on shutdown have been persistent frustrations.
- **State & concurrency bugs** – Duplicate LLM streams (#31487), subagent first message loss (#31072), and plugin permission replies silently dropped (#28037) indicate deep synchronization issues.
- **Security leakage** – MCP subprocess env exposure (#31778) and OAuth callback server not stopping (#23563, #32245) raise security concerns.
- **UI freezing after errors** – “thinking” hang (#32366) and tool call blocking (#32363) show poor error recovery paths that force app restarts.
- **Inconsistent tool call display** – MCP `isError=true` shown as completed (#16969) and missing error details for move operations (#32241) erode developer trust.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-06-15

## Today’s Highlights
Activity remained high with 40 issues and 13 pull requests updated in the last 24h. The community converged around two major themes: improving the extension developer experience (excludeFromContext, prompt guidelines, safe reloads) and cleaning up the model-registry build system (decomposing the monolithic `generate-models.ts`). Several TUI regressions — including a broken Escape interrupt and WezTerm image rendering — also drew attention and have been addressed.

## Releases
No new releases in the last 24 hours.

## Hot Issues
1. **#5103 – Windows bash detector fails with Git Bash on non‑default drive**  
   *CXwudi* — The built-in bash detection only looks under `C:\Program Files`, missing installations on other drives. 18 comments; no 👍 yet.  
   [GitHub](https://github.com/earendil-works/pi/issues/5103)

2. **#5702 – `prompt_cache_retention` sent to providers that reject it + maintainability concern**  
   *devasur* — A 400 error when using OpenAI‑compatible proxies that don’t support `cache_control`. The issue also flags the growing complexity of `generate-models.ts` (30‑branch cascade). 6 comments.  
   [GitHub](https://github.com/earendil-works/pi/issues/5702)

3. **#5736 – Escape no longer interrupts active interactive task**  
   *charles-cooper* — Pressing Escape during an agent run doesn’t reliably cancel the task. UI still advertises it as the interrupt key. 6 comments, marked `inprogress`.  
   [GitHub](https://github.com/earendil-works/pi/issues/5736)

4. **#5654 – Add `excludeFromContext` to custom messages sent via `sendMessage()`**  
   *zachmeador* — Mirror the flag already used for bash execution messages so extensions can inject messages that are persisted but not sent to the model. 6 comments, 1 👍.  
   [GitHub](https://github.com/earendil-works/pi/issues/5654)

5. **#5671 – `~/.pi` and `cwd/.pi` overlap**  
   *mitsuhiko* — Global settings stored in `~/.pi` conflict with project‑local `.pi` directories when the project is `$HOME`. 5 comments, 3 👍.  
   [GitHub](https://github.com/earendil-works/pi/issues/5671)

6. **#5710 – Extension‑level prompt guidelines**  
   *xl0* — Proposes a `pi.setPromptGuidelines()` API so extensions can influence model behavior without modifying system prompts. 4 comments.  
   [GitHub](https://github.com/earendil-works/pi/issues/5710)

7. **#5208 – `pi` crashes when background process exits late**  
   *kolt-mcb* – `ProcessRegistry` calls `output.finish()` on `exit`, but stdout/stderr can still emit after that, causing `uncaughtException`. 4 comments, `inprogress`.  
   [GitHub](https://github.com/earendil-works/pi/issues/5208)

8. **#5618 – WezTerm fails rendering images**  
   *vekexasia* — Since PR #4461, only a small chunk of header images is shown in WezTerm. 4 comments, closed as `inprogress`.  
   [GitHub](https://github.com/earendil-works/pi/issues/5618)

9. **#5575 – `kimi-k2.6` via OpenCode Go fails with JSON Schema conflict when tools enabled**  
   *josegil1909* — 400 Bad Request due to incompatible tool definitions. Highlights a broader provider compatibility problem. 4 comments.  
   [GitHub](https://github.com/earendil-works/pi/issues/5575)

10. **#5700 – Multiple live agent sessions with TUI switching**  
    *shmuelamit* — Wants concurrent agents without tearing down the current session. 4 comments.  
    [GitHub](https://github.com/earendil-works/pi/issues/5700)

## Key PR Progress
1. **#5743 – Decompose `generate-models.ts` into a data‑driven generator**  
   *devasur* — Behaviour‑preserving refactor to replace the ~30‑branch cascade with declarative descriptors. Draft, tied to #5702.  
   [GitHub](https://github.com/earendil-works/pi/pull/5743)

2. **#5738 – Fix pricing for Anthropic 1‑hour cache writes**  
   *theBucky* — Corrects under‑charging of 1‑hour cache writes by reading `ephemeral_1h_input_tokens`.  
   [GitHub](https://github.com/earendil-works/pi/pull/5738)

3. **#5678 – Add `excludeFromContext` for custom messages**  
   *mitsuhiko* — Extends the flag across agent harness, extension APIs, compaction, and branch summarization.  
   [GitHub](https://github.com/earendil-works/pi/pull/5678)

4. **#5735 – Defer extension reload requests safely**  
   *mitsuhiko* — Makes `ctx.reload()` available on base `ExtensionContext` with a deferral mechanism to avoid unsafe mid‑operation reloads.  
   [GitHub](https://github.com/earendil-works/pi/pull/5735)

5. **#5732 – Support `allowCommands` option in `sendUserMessage`**  
   *max-elia* — Enables extension‑injected messages to trigger slash commands (e.g., session resets) by enabling prompt template expansion.  
   [GitHub](https://github.com/earendil-works/pi/pull/5732)

6. **#5731 – Tool instrumentation for execution profiling**  
   *RHarshith* — Adds profiling hooks for tool execution, helping developers understand latency and resource usage.  
   [GitHub](https://github.com/earendil-works/pi/pull/5731)

7. **#5711 – Add extension prompt guideline API**  
   *xl0* — Implements the `pi.setPromptGuidelines()` proposal from #5710. Verified working.  
   [GitHub](https://github.com/earendil-works/pi/pull/5711)

8. **#5714 – Add xAI Grok account OAuth login**  
   *hyiiiii* — Built‑in OAuth provider for xAI Grok using device‑code flow; adds Grok subscription models to the provider list.  
   [GitHub](https://github.com/earendil-works/pi/pull/5714)

9. **#5385 – Detect first‑run terminal theme**  
   *vegarsti* – Queries terminal via OSC to auto‑set light/dark theme on first launch.  
   [GitHub](https://github.com/earendil-works/pi/pull/5385)

10. **#2331 – Vim‑like modal editor extension**  
    *Nokodoko* — Adds Normal, Insert, Visual, Command‑line modes with motions, operators, yank/paste, and ex commands. Still open after several months.  
    [GitHub](https://github.com/earendil-works/pi/pull/2331)

## Feature Request Trends
- **Extension API expansion** – The most active area: `excludeFromContext`, custom prompt guidelines, `allowCommands` for injected messages, safe extension reloads, and exposing raw provider responses in hooks.
- **Model/provider flexibility** – Requests for model‑specific compaction settings, adding missing models (Gemini 3.5 Flash, GLM‑5.2), provider‑specific config in `auth.json` (e.g., `accountId`), and OAuth support for xAI Grok.
- **TUI improvements** – Multiple‑session switching, rich slash‑command autocomplete popups, focus control when custom overlays are open, and scroll‑to‑bottom fixes during streaming.
- **Security & integrity** – Demand for SHA256SUMS and provenance attestations for binary releases.
- **First‑run experience** – Auto‑detecting terminal theme (light/dark) to set the default Pi theme.

## Developer Pain Points
- **Windows path handling** – The bash detector fails on non‑default drives; several issues cite incomplete platform support.
- **Caching/TTL quirks** – `cache_control` blocks with TTLs are sent to providers that reject them, and the Anthropic cache write pricing mismatch for 1‑hour vs 5‑minute TTLs causes cost errors.
- **TUI rendering regressions** – CJK wide characters split vertical borders, WezTerm images are truncated, chat view jumps to top during streaming, and overlays block transcript scrolling.
- **Process output truncation** – Bash tool loses output when child processes hold stdout past exit; `git commit` with pre‑commit hooks is a common trigger.
- **Terminal state corruption** – Killing Pi with SIGTERM leaves raw mode and Kitty keyboard protocol enabled, breaking the parent shell.
- **Model schema conflicts** – Custom providers (e.g., kimi‑k2.6 via OpenCode Go) fail with 400 due to tool definition incompatibility.
- **Extension reliability** – Extension reload requests are unsafe from non‑slash‑command contexts; custom overlays cannot yield focus to the transcript.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-06-15

**Today’s Highlights:**  
A significant **Trojan false-positive** (#5055) affecting the VSIX packaging of v0.18.0 remains open and under triage, while the team closed a critical CI “green-success” fake-out bug (#5052) and a multi-agent `/review` breakage (#5100). On the feature front, PRs continue to land for **interactive extension management** (#4850), **per-task resource telemetry** (#5118), and a **safe-mode troubleshooting flag** (#4943).

## Releases  
*None in the last 24 hours.*

## Hot Issues *(10 selected)*

1. **[OPEN] #5055 — Trojan:JS/ShaiWorm.DBA!MTB**  
   A user reports that the VSIX for v0.18.0 triggers Windows Defender. No model source has been tampered, but the false-positive blocks enterprise deployment. *5 comments, no resolution yet.*  
   → [Issue #5055](https://github.com/QwenLM/qwen-code/issues/5055)

2. **[OPEN] #5119 — Sudo commands fail ungracefully**  
   Agents cannot execute sudo even after user confirmation; they must manually provide the command. Community wants a permission dialog that detects and elevates privilege. *1 comment, priority P2.*  
   → [Issue #5119](https://github.com/QwenLM/qwen-code/issues/5119)

3. **[CLOSED] #5052 — CI reports green success when API error kills mid-review**  
   A job exits 0 despite a mid-stream API disconnect, posting zero review comments. Fixed by adding exit-code & comment-count assertions. *2 comments.*  
   → [Issue #5052](https://github.com/QwenLM/qwen-code/issues/5052)

4. **[CLOSED] #5100 — Agent team `name` breaks `/review` skill**  
   Setting a `name` on an Agent Team causes the bundled review skill to fail with “no active team” and then loop into provider aborts. Closed with a fix. *2 comments.*  
   → [Issue #5100](https://github.com/QwenLM/qwen-code/issues/5100)

5. **[OPEN] #3979 — Ghostty terminal flicker after plan-mode completion**  
   On macOS with Ghostty, the terminal flickers continuously after Qwen Code finishes a plan reply. *2 comments, autofix in-progress.*  
   → [Issue #3979](https://github.com/QwenLM/qwen-code/issues/3979)

6. **[CLOSED] #4727 — Dual Output mode TUI unresponsive**  
   Using `--json-file` + `--input-file` with a named pipe leaves the TUI dead. Closed after 5 comments.  
   → [Issue #4727](https://github.com/QwenLM/qwen-code/issues/4727)

7. **[OPEN] #4349 — estimatePromptTokens misses mid-session token count**  
   The steady-state branch omits the model’s previous response tokens, leading to under-estimation during long sessions. *1 comment, 1 like.*  
   → [Issue #4349](https://github.com/QwenLM/qwen-code/issues/4349)

8. **[OPEN] #3184 — Infinite loop when edit fails**  
   When Qwen Code cannot find a string to edit, it loops indefinitely rather than asking the user. *1 like, 0 comments.*  
   → [Issue #3184](https://github.com/QwenLM/qwen-code/issues/3184)

9. **[OPEN] #3424 — Generic 403 “Access to model denied”**  
   No helpful error context; user just sees “403”. *0 comments.*  
   → [Issue #3424](https://github.com/QwenLM/qwen-code/issues/3424)

10. **[OPEN] #3884 — Model reads /home/userfile outside project**  
    The agent suddenly tries to open an absolute path outside the permitted workspace. *0 comments.*  
    → [Issue #3884](https://github.com/QwenLM/qwen-code/issues/3884)

## Key PR Progress *(10 selected)*

1. **[OPEN] #5073 — Warn on oversized context instructions**  
   Startup warning when `QWEN.md` / instructions exceed 15% of model context. Small UX win for long-context setups.  
   → [PR #5073](https://github.com/QwenLM/qwen-code/pull/5073)

2. **[OPEN] #4850 — Interactive multi-tab /extensions manager**  
   Turns the flat `/extensions` list into a three-tab manager (Installed / Discover / Sources) covering the full lifecycle.  
   → [PR #4850](https://github.com/QwenLM/qwen-code/pull/4850)

3. **[OPEN] #5094 — Workflow meta-extraction (P4a)**  
   First half of phase P4 of the Dynamic Workflows port: extracts and strips meta from run outcomes. Builds on P1–P3.  
   → [PR #5094](https://github.com/QwenLM/qwen-code/pull/5094)

4. **[OPEN] #5030 — Resume interrupted turn without synthetic “continue”**  
   First-class continuation of unfinished assistant turns after crash or resume. Closes the gap in #4679.  
   → [PR #5030](https://github.com/QwenLM/qwen-code/pull/5030)

5. **[OPEN] #5118 — Per-task token & time detail in web-shell todos**  
   Clicking a completed todo now shows start/end time, elapsed duration, token counts, and API time.  
   → [PR #5118](https://github.com/QwenLM/qwen-code/pull/5118)

6. **[OPEN] #5120 — Skip auto-title when history has no user message**  
   Guards `tryGenerateSessionTitle` to return `empty_history` until a real user message exists. Prevents bogus titles on daemon sessions.  
   → [PR #5120](https://github.com/QwenLM/qwen-code/pull/5120)

7. **[OPEN] #5122 — Configurable screenshot max dimension (CUA)**  
   Adds a user knob + env var for the screenshot longest-edge cap in the computer-use driver.  
   → [PR #5122](https://github.com/QwenLM/qwen-code/pull/5122)

8. **[OPEN] #4866 — Split PR triage into 4-job pipeline**  
   Replaces monolithic `/triage` with a staged `qwen-pr-triage.yml` pipeline (resolve → product-decision → findings → doc).  
   → [PR #4866](https://github.com/QwenLM/qwen-code/pull/4866)

9. **[CLOSED] #4520 — Truncate model-facing tool output**  
   Moves truncation from shell tool into `CoreToolScheduler` so any tool returning string output is bounded before entering history.  
   → [PR #4520](https://github.com/QwenLM/qwen-code/pull/4520)

10. **[OPEN] #4943 — `--safe-mode` CLI flag**  
    Disables all customizations (context files, hooks, extensions, MCP servers) for a clean baseline.  
    → [PR #4943](https://github.com/QwenLM/qwen-code/pull/4943)

## Feature Request Trends

- **Security & permissions** – multiple requests for sudo elevation flow (#5119), safe-mode troubleshooting flag (#4943), and quarantine of virus false-positives (#5055).  
- **Extension & agent UX** – rich extension management (#4850), agent team configuration (#5100), and per-task resource telemetry (#5118).  
- **Context & workspace control** – warnings for oversized instructions (#5073), prevention of rogue path access (#3884), and custom ignore files (#4653).  
- **Session & history improvements** – persistent history collapse (#4085), resume without synthetic messages (#5030), and timestamps in CLI (#5001).  

## Developer Pain Points

1. **False-positive CVEs** – The VirusTotal hit on v0.18.0 VSIX (#5055) blocks adoption on locked-down Windows workstations.  
2. **Silent failures** – CI greencheck on API errors (#5052) and looping on edit failures (#3184) erode trust in automation.  
3. **Invasive prompt scope** – Models reading files outside the project directory (#3884) or looping indefinitely (#3184) show insufficient boundary enforcement.  
4. **Auth & error opacity** – Generic 403 errors (#3424) and token-estimation drift (#4349) make debugging expensive.  
5. **Terminal compatibility** – Ghostty flicker (#3979) and TUI lock-up in Dual Output mode (#4727) signal ongoing terminal-rendering fragility.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest – 2026-06-15

*Project has been rebranded to **CodeWhale**. Legacy names `deepseek-tui` / `deepseek` are deprecated.*

---

## Today's Highlights

The v0.8.60 release makes the **CodeWhale** rebrand official, but the community is feeling the migration pains: npm/cargo update paths are broken for many users, and authentication fails with popular Chinese AI providers like SiliconFlow and Tencent Cloud TokenHub. On the bright side, the maintainer has merged a large batch of community harvest PRs (`#3197`, `#3051`, `#2811` etc.) and opened a draft `v0.8.61` with WhaleFlow swarm foundations and a freeze fix. The ecosystem is clearly pivoting toward multi-provider orchestration, but legacy compatibility and documentation gaps remain the top blockers.

---

## Releases

- **[v0.8.60](https://github.com/Hmbown/CodeWhale/releases/tag/v0.8.60)** – Rebrand release: project, command, npm package, and release assets now use the `codewhale` name. The legacy `deepseek-tui` npm package is deprecated and receives no further releases. Users on `v0.8.x` legacy names must follow `docs/REBRAND.md` to migrate.

---

## Hot Issues (10 of 21)

1. **[#2629 – SiliconFlow & Tencent Cloud TokenHub auth error (401)](https://github.com/Hmbown/CodeWhale/issues/2629)**  
   • *Created 2026-06-03, updated 2026-06-14, 3 comments*  
   • Configuring the OpenAI-compatible endpoint for these providers always returns `401 invalid api key`. Blocks a large segment of Chinese users. Community is waiting for a fix.

2. **[#1067 – glibc 2.38/2.39 requirement too high for many servers](https://github.com/Hmbown/CodeWhale/issues/1067)**  
   • *Created 2026-05-07, updated 2026-06-14, 3 comments*  
   • Ubuntu servers with glibc 2.35 cannot run the prebuilt binary. No workaround yet; users request a statically linked build or lower requirement.

3. **[#3102 – First-class clarifying question requests for agents](https://github.com/Hmbown/CodeWhale/issues/3102)**  
   • *Created 2026-06-11, updated 2026-06-14, 3 comments*  
   • Feature request to give agents a structured way to ask the user for clarification, rather than hiding in normal chat messages. High priority for agent UX.

4. **[#3231 – Support DeepInfra as a provider](https://github.com/Hmbown/CodeWhale/issues/3231)**  
   • *Created 2026-06-15, 1 comment*  
   • Fresh request; official provider list does not include DeepInfra. Quick community ask.

5. **[#2924 – Can't update to new version using npm](https://github.com/Hmbown/CodeWhale/issues/2924)**  
   • *Created 2026-06-09, updated 2026-06-14, 1 comment, 1 👍*  
   • `npm update deepseek-tui` (or `codewhale`) fails silently or with confusing errors. A symptom of the rebrand migration.

6. **[#2917 – Cargo distribution: `codewhale` not found after rename](https://github.com/Hmbown/CodeWhale/issues/2917)**  
   • *Created 2026-06-08, updated 2026-06-14, 1 comment*  
   • Users who installed via `cargo install deepseek-tui` now get `error: failed to spawn 'codewhale'` when running the old command.

7. **[#2652 – Subagent clipped output mistaken as complete evidence](https://github.com/Hmbown/CodeWhale/issues/2652)**  
   • *Created 2026-06-03, updated 2026-06-14, 1 comment*  
   • Sub-agent tool outputs are truncated with "Alt+V for details" but the model treats the truncated view as full context. Reliability bug.

8. **[#3222 – Add `reasoning_style` override for inline-tag thinking blocks](https://github.com/Hmbown/CodeWhale/issues/3222)**  
   • *Created 2026-06-14, 2 comments*  
   • Parsing of reasoning content from MiniMax M3, Qwen, GLM is broken. Requests an override mechanism for OpenAI chat-completions style.

9. **[#3224 – Keyboard shortcut renovation for composer and queued steering](https://github.com/Hmbown/CodeWhale/issues/3224)**  
   • *Created 2026-06-14, 0 comments (picked for importance)*  
   • New `Ctrl+S` send behavior conflicts with other shortcuts; the keyboard scheme is hard to discover. Design feedback requested.

10. **[#3230 – WhaleFlow synthesis/reduce pass for many workers](https://github.com/Hmbown/CodeWhale/issues/3230)**  
    • *Created 2026-06-14, 1 comment*  
    • WhaleFlow swarm lacks a reduce stage to merge multiple worker outputs into one coherent result. Core architectural gap for ultracode-style orchestration.

---

## Key PR Progress (10 of 44)

1. **[#3197 – Rename DeepSeek blue consumers to whale accent](https://github.com/Hmbown/CodeWhale/pull/3197)** *(CLOSED)*  
   • Replaces 40+ `DEEPSEEK_BLUE` references with `WHALE_ACCENT_PRIMARY`. Deprecated aliases kept for migration. Important cosmetic step in rebrand.

2. **[#3051 – Voice input via `/voice` slash command](https://github.com/Hmbown/CodeWhale/pull/3051)** *(CLOSED)*  
   • Adds three new slash commands for speech recording, AI transcription, and composer insertion. Inspired by MiMo Code's UX.

3. **[#3225 – v0.8.61 community harvest + freeze fix + WhaleFlow foundation (draft)](https://github.com/Hmbown/CodeWhale/pull/3225)** *(CLOSED, draft)*  
   • 28 commits on `codex/v0.8.61` including many community PRs, a launch-blocker fix, and the base for WhaleFlow orchestration.

4. **[#2811 – VS Code extension scaffold](https://github.com/Hmbown/CodeWhale/pull/2811)** *(CLOSED)*  
   • Adds Phase 0 VS Code extension with commands to open CodeWhale, start HTTP server, check local runtime. Important for IDE integration.

5. **[#2102 – Defer low-value native tools by default](https://github.com/Hmbown/CodeWhale/pull/2102)** *(CLOSED)*  
   • Tools beyond core are lazy-loaded; adds `[tools] always_load` config. Reduces startup cost and memory.

6. **[#2646 – Harden Cargo publish ordering and crates.io checks](https://github.com/Hmbown/CodeWhale/pull/2646)** *(CLOSED)*  
   • Fixes release tooling: correct workspace dependency order, crates.io TTL check. Prevents failed `cargo publish`.

7. **[#2771 – Harvest LLM-guided AGENTS.md init](https://github.com/Hmbown/CodeWhale/pull/2771)** *(CLOSED)*  
   • The `/init` command now delegates AGENTS.md generation to the agent, with credential-safe HTTP(S) support and workspace context detection.

8. **[#2802 – Hugging Face MCP helpers](https://github.com/Hmbown/CodeWhale/pull/2802)** *(CLOSED)*  
   • Adds `/hf mcp status`, `/hf mcp setup`, `/hf concepts` slash commands. Bridges Hugging Face settings with CodeWhale MCP.

9. **[#2795 – Enrich auth errors with request context](https://github.com/Hmbown/CodeWhale/pull/2795)** *(CLOSED)*  
   • Authentication failures now show provider, base URL, model, key source/fingerprint. Drastically improves debugging of 401 errors like #2629.

10. **[#2103 – Fix mouse capture on Windows keeping history arrows](https://github.com/Hmbown/CodeWhale/pull/2103)** *(CLOSED)*  
    • Removes blanket Windows override for empty-composer arrow keys; preserves scroll fallback. Fixes #1720.

---

## Feature Request Trends

- **Multi-provider ecosystem expansion** – Strong demand for first-class support beyond DeepSeek: DeepInfra (#3231), SiliconFlow/TokenHub (#2629), MiniMax M3, Qwen, GLM thinking blocks (#3222). The pricing table is also dead for all non-DeepSeek models (#3066).
- **Agent orchestration & swarm** – WhaleFlow synthesis/reduce (#3230), fleet ledger coordination (#3229), parent-visible worker interaction (#3226). The community wants scalable multi-agent workflows.
- **UX & keyboard renovation** – Shortcut conflicts (#3224), clarifying questions popup (#3102), sidebar management, voice input (#3051), and better queued steering.
- **Global context & instructions** – Auto-load global `~/.codewhale/instructions.md` (#3012) alongside project-level files.
- **Provider/model isolation per session** – Multiple TUI sessions should not share provider/model state (#3227).

---

## Developer Pain Points

1. **Rebrand migration breakage** – `npm update`, `cargo install`, and `deepseek update` all fail or produce cryptic errors (#2924, #2917, #2960). Users on legacy installs are stuck.
2. **glibc compatibility** – Prebuilt binaries require glibc ≥2.38/2.39, excluding many Linux servers (#1067).
3. **Authentication errors with non-DeepSeek providers** – 401 errors with Chinese providers (#2629) are hard to debug even with the new enriched error messages (#2795 now merged).
4. **Confusing release assets** – Both `codewhale-linux-x64` (binary) and `.tar.gz` archives are published without clear documentation (#3208).
5. **Cost tracking dead for non-DeepSeek** – Pricing table returns `None` for most models, so session cost and cache savings readouts are broken (#3066).
6. **Sub-agent output clipping** – Truncated tool output misleads the model, causing unreliable agent behavior (#2652).
7. **Legacy `.deepseek/` paths** – Multiple code paths still reference the old directory; migration status is undocumented (#3068).

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*