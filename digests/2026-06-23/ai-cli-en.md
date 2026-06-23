# AI CLI Tools Community Digest 2026-06-23

> Generated: 2026-06-23 10:50 UTC | Tools covered: 9

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

# AI CLI Developer Tools: Cross-Tool Comparison Report
**Date**: 2026-06-23  
**Analyst**: Senior Technical Analyst, AI Developer Tools Ecosystem

---

## 1. Ecosystem Overview

The AI CLI tools landscape on June 23, 2026, reflects a maturing ecosystem where **200+ issues and 170+ pull requests** were filed across seven major tools in a single day. The community is collectively grappling with three universal pressures: **cost unpredictability** (spikes in per-token pricing, unexpected credit consumption), **security hardening** (SSRF vulnerabilities, permission model weaknesses, secret management), and **architectural evolution** (MCP protocol maturity, agent orchestration, multi-provider support). Development velocity is uneven—OpenAI Codex and the newly rebranded CodeWhale (DeepSeek TUI) show the highest PR throughput, while Kimi Code and GitHub Copilot CLI appear more conservative. A clear divide is emerging between **platform-first tools** (Codex, Claude Code) investing in ecosystem lock-in via plugins and MCP governance, and **open-architecture tools** (Pi, Qwen Code, Gemini CLI) prioritizing provider flexibility and local-first workflows. The most striking signal: **permission and authentication systems remain the #1 source of community frustration**, with cross-tool complaints about broken prompts, false approvals, and regressions persisting for 10+ months in some cases.

---

## 2. Activity Comparison (Today)

| Tool | Issues (Today) | PRs (Today) | Release Status | Notable Pattern |
|------|----------------|-------------|----------------|-----------------|
| **Claude Code** | 10 hot issues | 4 open PRs | v2.1.186 (stable) | High duplication (systematic regressions June 19-20) |
| **OpenAI Codex** | 10 hot issues | 10 active PRs | rust-v0.142.0 stable + 6 alphas | Most aggressive release cadence |
| **Gemini CLI** | 10 hot issues | 10 active PRs | No release today | Security-focused (2 SSRF patches merged) |
| **GitHub Copilot CLI** | 10 notable items | 1 PR | v1.0.64-3, v1.0.64-2 (patches) | Low PR throughput, stable patches |
| **Kimi Code CLI** | 3 issues | 2 PRs | v1.48.0 (minor) | Smallest activity footprint |
| **OpenCode** | 10 notable issues | 10 key PRs | No release today | High community engagement (80+ PRs total) |
| **Pi** | 10 hot issues | 10 key PRs | No release today | Broad provider ecosystem expansion |
| **Qwen Code** | 10 hot issues | 10 key PRs | v0.19.0 + v0.19.1 | Architecture-heavy discussions |
| **CodeWhale (ex-DeepSeek TUI)** | 10 notable issues | 10 key PRs (50 total) | v0.8.64 (first rebrand release) | Highest single-day PR count |

**Key observation**: CodeWhale (50 PRs) and OpenCode (80+ PRs total) represent the most active codebases today, though both are smaller-user-base projects. OpenAI Codex's 10 active PRs with 6 alpha releases shows the highest release velocity among major tools.

---

## 3. Shared Feature Directions

Five cross-tool patterns emerged from today's data:

### 3.1 MCP Ecosystem Maturity (6 tools)
- **Claude Code**: New `claude mcp login/logout` commands for CLI-based auth
- **GitHub Copilot CLI**: Policy validation bugs (#2486, #3162), MCP server instructions ignored (#1579)
- **Kimi Code CLI**: MCP server path resolution bug (#2469)
- **OpenCode**: Resource read tools added (PR #33483, closing 2-year-old issues)
- **Qwen Code**: MCP server discovery + resource completion in v0.19.1
- **CodeWhale**: MCP duplicate server processes (#3461)

**Common needs**: CLI-based auth flows, resource read/write parity, proper policy/instructions handling, server lifecycle management, cross-platform path resolution.

### 3.2 Authentication & Credential Management (5 tools)
- **Gemini CLI**: OAuth socket-reuse regression on Node 24 (#28103)
- **GitHub Copilot CLI**: Session resume authentication failures (#3596)
- **Kimi Code CLI**: Approval prompt inconsistency in yolo mode (#2448)
- **Pi**: Hardcoded OAuth token detection (#5871), credential duplication (#5953)
- **CodeWhale**: Provider-scoped API keys from secret managers (#3004)

**Common needs**: Flexible token formats, session persistence, credential deduplication, external secret manager integration.

### 3.3 Permission & Safety Systems (5 tools)
- **Claude Code**: Permissions system broken 10+ months (#30519)
- **OpenCode**: "Allow always" overrides deny rules (#31540), no destructive command guard (#33077)
- **Qwen Code**: Deterministic guards for destructive Git commands (#5749, PR #5754)
- **CodeWhale**: Destructive approval too intrusive (#3466)
- **Kimi Code CLI**: Yolo mode unreliability (#2448)

**Common needs**: Granular deny-first rules, semantic command analysis, configurable approval levels, deterministic kill switches for dangerous operations.

### 3.4 Cost Transparency & Usage Governance (5 tools)
- **Claude Code**: Parabolic usage spike (#69892), billing/VAT errors (#42018)
- **OpenAI Codex**: 10-20x per-token cost jump (#28879), SQLite write amplification cost (#28224)
- **GitHub Copilot CLI**: 174-credit restart consumption (#3886)
- **OpenCode**: Paid balance still hits free limit (#33318)
- **Pi**: Connection reliability stalls (#4945) leading to wasted session costs

**Common needs**: Per-session/step cost breakdowns, rate-limit visibility, predictable pricing tiers, backoff/retry resilience.

### 3.5 Multi-Platform & Cross-Environment Support (5 tools)
- **OpenAI Codex**: Linux desktop app demand (#11023, 628 👍), Windows crashes (#29320)
- **Gemini CLI**: Wayland browser agent failures (#21983)
- **GitHub Copilot CLI**: Windows scroll regression (#1944)
- **OpenCode**: macOS PAC trap crash (#32200), Windows file delete freezes (#33491)
- **Qwen Code**: Chrome extension architecture overhaul (#5626)

**Common needs**: Native Linux support, Wayland compatibility, Windows-specific testing, consistent session state across platforms.

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | GitHub Copilot CLI | Kimi Code CLI | Pi | OpenCode | Qwen Code | CodeWhale |
|-----------|-------------|--------------|------------|--------------------|---------------|-----|-----------|-----------|-----------|
| **Primary Model** | Anthropic Claude | GPT-5.5 / o-series | Gemini 2.5 | GPT-4o / Claude | Moonshot Kimi | Multi-provider | Copilot / Multi | Qwen / Alibaba | DeepSeek / Multi |
| **Dev Philosophy** | Complete platform | Ecosystem lock-in | Architecture-first | Developer simplicity | Minimalist | Provider-agnostic | Open desktop | Local+cloud hybrid | Multi-provider TUI |
| **Strength** | MCP integration, workflows | Plugin ecosystem, cost governance | Security hardening, AST tools | GitHub integration, simplicity | Autonomous mode | Provider flexibility | Community contributions | Architecture innovation | Rebrand agility |
| **Weakness** | Slow bug fixes (10mo+) | Cost unpredictability | Agent reliability | Low innovation velocity | Small community, low activity | Connection reliability | Cross-platform stability | Default token caps | Documentation lag |
| **Target User** | Enterprise teams | Pro developers, Windows | Security-conscious devs | GitHub ecosystem users | Power CLI users | Self-hosters | Desktop-first devs | Asian market, local LLM | Budget-conscious devs |
| **Release Pace** | ~1-2 weeks | Multiple alphas daily | No release today | ~2 patch releases | Single minor release | No release today | No release today | 2 releases today | Single rebrand release |

**Key differentiators**:
- **Claude Code** owns the permission/security conversation but fails to execute on fixes
- **OpenAI Codex** leads in platform features (plugins hub, credit system) but suffers cost trust issues
- **Gemini CLI** most security-conscious (SSRF patches, CI trust remediation)
- **GitHub Copilot CLI** sticks to core simplicity, avoids feature bloat
- **Pi** maximizes provider flexibility (Merge Gateway, Anthropic Vertex incoming)
- **Qwen Code** innovates in architecture (protocol decoupling, visual bridging)
- **CodeWhale** shows highest codebase churn (50 PRs/day) during rebrand transition

---

## 5. Community Momentum & Maturity

### Highly Active & Iterating
- **OpenAI Codex**: Most mature platform, aggressive alpha pipeline, large community (628 👍 for Linux app). Risk: cost erosion of trust.
- **CodeWhale (DeepSeek TUI)**: Highest PR density (50/day), successful rebrand execution. Risk: post-rebrand confusion, documentation lag.
- **OpenCode**: Strongest community contribution pattern (80+ PRs), Turkish localization, mobile improvements. Risk: cross-platform stability concerns.

### Mature & Stable
- **Claude Code**: Largest enterprise user base, high issue visibility, but **dangerously slow response** (30519 unresolved 10 months). Community building workarounds independently—sign of ecosystem resilience but vendor risk.
- **GitHub Copilot CLI**: Lowest innovation velocity but highest reliability (patches only). Risk: falling behind on MCP/feature parity.

### Emerging & Growing
- **Gemini CLI**: Strong architectural foundations (AST tools, SSRF hardening). Currently troubled by agent reliability issues (#21409, #22323) that limit trust.
- **Pi**: Fastest expanding provider ecosystem (4 new providers in PRs). Community engaged (66 comments on connection issue). Risk: connection reliability remains top pain point.

### Quiet & Niche
- **Kimi Code CLI**: Smallest activity footprint (3 issues, 2 PRs). New force-stop feature shows attention to autonomous mode safety, but low engagement suggests limited adoption.
- **Qwen Code**: Strong architecture discussions but small visible community. Asian market focus may limit ecosystem contributions.

### Community Health Signals
- **Most upvoted issue today**: OpenAI Codex Linux app (#11023, 628 👍)
- **Most commented issue today**: Claude Code permissions (#30519, 24 comments, 76 👍)
- **Most active discussion**: Pi OpenAI Codex connection reliability (#4945, 66 comments, 30 👍)
- **Most controversial feature**: CodeWhale destructive approval (#3466, user wants old behavior back)

---

## 6. Trend Signals

### Signal 1: Cost Sensitivity Reaches Breaking Point
The **10-20x per-token cost jump** on OpenAI Codex (#28879), **parabolic usage spikes** on Claude Code (#69892), and **paid balances hitting free limits** on OpenCode (#33318) signal that AI CLI tool pricing is becoming unpredictable and untrustworthy. Developers are actively seeking alternatives: Pi's Merge Gateway extension (40+ models, single API key), Qwen Code's local inference push, and CodeWhale's multi-provider architecture all reflect a hedge against vendor lock-in and cost volatility.

**Actionable insight**: Prioritize cost transparency features (per-step breakdown, rate-limit alerts, budget caps). The "free limit with paid account" bug category is especially damaging to trust.

### Signal 2: Security Hardening Is Now Table Stakes
Gemini CLI merged **two SSRF patches** today (#27744, #27739) addressing DNS hijacking and redirect-based private IP access. OpenCode patched permission bypass (#33077, #31540). Qwen Code added deterministic destructive command guards (#5754). CodeWhale added secret-manager API key integration (#3004). The community is demanding security built into the architecture, not bolted on later.

**Actionable insight**: Every tool should have an auditable permission model, deterministic command guards, and external secret management by default. The age of "allow all" defaults is ending.

### Signal 3: Multi-Provider Architecture Wins
Pi (Merge Gateway, Anthropic Vertex, auto-router), CodeWhale (Xiaomi MiMo, Alibaba Bailian, Baidu Qianfan, OpenRouter), and Qwen Code (visual bridging, protocol decoupling) all demonstrate that **provider flexibility is the defining competitive advantage**. Users want to switch models without switching tools. OpenAI Codex's ecosystem lock-in risks being a liability if cost issues continue.

**Actionable insight**: Invest in provider interface layers, model routing, and cost comparison dashboards. The tool that makes it easiest to swap providers wins the next wave of adoption.

### Signal 4: Agent Reliability Frustration Is Universal
- **Gemini CLI**: Generalist agent hangs (#21409), false "GOAL" success after MAX_TURNS (#22323)
- **Qwen Code**: 8K output cap causes retry loops (#5756)
- **CodeWhale**: Fleet worker freeze regression (#3289)
- **Kimi Code CLI**: Hang on detached child processes (#2468)
- **Pi**: "Working..." freeze on Codex (#4945), Anthropic subscription hangs (#5291)

Every tool with autonomous agent capabilities has a reliability issue. The community wants **deterministic, observable, and interruptible agents**.

**Actionable insight**: Agent observability (subagent trajectory sharing, tool-call transparency, "thinking" visibility) and self-healing (backoff, retry, graceful degradation) are the next battleground features.

### Signal 5: Voice & Multimodal Are Emerging Frontiers
Qwen Code added server-side voice dictation (#5755) and visual bridging for text-only models (#5126). OpenAI Codex has vocal demand for TUI voice transcription return (#16404). These signal that **CLI tools are evolving beyond keyboard-and-text** into multimodal assistants.

**Actionable insight**: Voice input for non-interactive modes, image interpretation for debugging, and screen capture for web automation will become table stakes within 12 months.

### Signal 6: Localization & Internationalization Matter
OpenCode shipped Turkish localization (#33489) and has demand for Vietnamese (#29309). Claude Code has EU B2B VAT errors (#42018). Qwen Code serves the Asian market. CodeWhale added Xiaomi and Alibaba providers. **The AI CLI market is global, and language/currency support is a competitive differentiator.**

**Actionable insight**: International payment handling, locale-aware UX, and localized model providers should be part of roadmap planning, not afterthoughts.

---

## Summary for Technical Decision-Makers

| If you value... | Consider... | Because... |
|-----------------|-------------|------------|
| **Enterprise stability** | Claude Code | Most mature platform, but monitor #30519 closely |
| **Cost control** | CodeWhale / Pi | Multi-provider, auto-routing extensions |
| **Security maturity** | Gemini CLI | SSRF patches, CI trust remediation, AST-aware auditing |
| **GitHub integration** | GitHub Copilot CLI | Simple, reliable, but feature-limited |
| **Desktop experience** | OpenCode | Strongest TUI, active community localization |
| **Asian market / local LLM** | Qwen Code | Alibaba ecosystem, local inference focus |
| **Maximum provider flexibility** | Pi | Merge Gateway + 4+ new providers in PRs |
| **Platform features** | OpenAI Codex | Best plugin UX, but cost trust is eroding |

**Bottom line**: The AI CLI tool landscape is fragmenting along the **platform vs. open architecture** axis. Platform tools (Codex, Claude Code) offer richer integrations but risk vendor lock-in. Open tools (Pi, Qwen Code, CodeWhale) offer flexibility but lag in UX polish. The winners will be those that solve **cost transparency, security hardening, and agent reliability**—the three universal pain points cutting across every community today.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report

*Data snapshot: 2026-06-23 | Source: github.com/anthropics/skills*

---

## 1. Top Skills Ranking

The community's most-discussed contributions cluster around two themes: **critical infrastructure fixes** for the skill-creator toolchain, and **specialized domain skills** for document processing and agent memory. Below are the 8 most-attended Skills by community engagement:

### skill-creator: run_eval.py recall failures (#1298)
**Status:** Open | **Author:** MartinCajiao | **Created:** 2026-06-10  
**Summary:** The single most critical fix in the repository. `run_eval.py`—and by extension `run_loop.py` and `improve_description.py`—reports `recall=0%` for every skill description due to the eval artifact not being installed as a real skill. Multiple independent reproductions (#556, #1169) confirm the optimization loop is optimizing against noise. The PR fixes Windows stream reading, trigger detection, and parallel workers.  
**Discussion highlights:** Directly addresses the most-upvoted open bug (#556, 12 comments, 👍7). Considered the blocker preventing any meaningful description optimization.  
🔗 https://github.com/anthropics/skills/pull/1298

### skill-creator: trigger detection misses real skill name (#1323)
**Status:** Open | **Author:** Polluelo978 | **Created:** 2026-06-16  
**Summary:** Companion fix to #1298. `run_single_query` fails to detect skill triggers because it compares against a hardcoded name rather than the actual installed skill name. Also bails on the first non-Skill tool call, producing `recall=0%` for every should-trigger query.  
**Discussion highlights:** Filed just one week ago but already accumulating engagement. Together with #1298, forms a comprehensive fix for the broken evaluation pipeline.  
🔗 https://github.com/anthropics/skills/pull/1323

### document-typography skill (#514)
**Status:** Open | **Author:** PGTBoos | **Created:** 2026-03-04  
**Summary:** Adds typographic quality control for AI-generated documents. Addresses orphan word wrap (1-6 words on a new line), widow paragraphs (headers stranded at page bottom), and numbering misalignment. The author argues these "affect every document Claude generates" yet users rarely request them explicitly.  
**Discussion highlights:** Well-received as a pragmatic quality-of-life skill. Represents a category of "invisible polish" that Claude would never do unless explicitly instructed.  
🔗 https://github.com/anthropics/skills/pull/514

### ODT skill (OpenDocument text creation) (#486)
**Status:** Open | **Author:** GitHubNewbie0 | **Created:** 2026-03-01  
**Summary:** Comprehensive skill for creating, filling, reading, and converting OpenDocument Format files (.odt, .ods). Handles LibreOffice document generation, template filling, and ODT-to-HTML conversion. Triggers on mentions of "ODT," "ODS," "ODF," "OpenDocument," or requests for open-source standard documents.  
**Discussion highlights:** Fills a clear gap for users working with LibreOffice and ISO-standard formats. Complements the existing PDF and DOCX skills in the document-skills plugin.  
🔗 https://github.com/anthropics/skills/pull/486

### skill-quality-analyzer & skill-security-analyzer (#83)
**Status:** Open | **Author:** eovidiu | **Created:** 2025-11-06  
**Summary:** First major contribution of "meta-skills" that evaluate other skills. The quality analyzer scores across five dimensions (Structure & Documentation 20%, Implementation & Testing 25%, etc.). The security analyzer detects command injection, path traversal, and unsafe file operations.  
**Discussion highlights:** Long-running discussion about meta-skill utility and maintenance burden. Sets precedent for the ecosystem's self-evaluation capabilities.  
🔗 https://github.com/anthropics/skills/pull/83

### AppDeploy skill (#360)
**Status:** Open | **Author:** avimak | **Created:** 2026-02-09  
**Summary:** Enables Claude to deploy and manage full-stack web applications to public URLs using AppDeploy.ai. Covers the full lifecycle: deployment, status checks, versioning, and teardown.  
**Discussion highlights:** Represents a new category of "action skills" that interact with external deployment platforms. Community interest suggests demand for Claude-as-operator capabilities.  
🔗 https://github.com/anthropics/skills/pull/360

### testing-patterns skill (#723)
**Status:** Open | **Author:** 4444J99 | **Created:** 2026-03-22  
**Summary:** Comprehensive testing skill covering the full stack: testing philosophy (Testing Trophy model), unit testing (AAA pattern, pure functions), React component testing (Testing Library), integration, E2E, and mocking strategies.  
**Discussion highlights:** Addresses a clear community need for structured testing guidance. The breadth of coverage (philosophy through implementation) makes it one of the most actionable skills submitted.  
🔗 https://github.com/anthropics/skills/pull/723

### shodh-memory skill (#154)
**Status:** Open | **Author:** varun29ankuS | **Created:** 2025-12-19  
**Summary:** Persistent memory system for AI agents that maintains context across conversations. Teaches Claude when to call `proactive_context`, how to structure rich memories, and how to surface relevant memories per user message.  
**Discussion highlights:** Active discussion about memory persistence patterns and integration with external storage. Competes with the newer compact-memory proposal (#1329).  
🔗 https://github.com/anthropics/skills/pull/154

---

## 2. Community Demand Trends

Analysis of the 14 top-commented issues reveals five concentrated demand vectors:

### 1. Enterprise Sharing & Distribution (#228 — 14 comments, 👍7)
The highest-engagement issue demands organizational skill sharing within Claude.ai. Currently users must download `.skill` files and distribute manually via Slack/Teams. A shared skill library or direct sharing links is the most requested enterprise feature.  
🔗 https://github.com/anthropics/skills/issues/228

### 2. Toolchain Reliability (#556 — 12 comments, 👍7; #1169 — 3 comments; #1061 — 3 comments)
The `run_eval.py` recall=0% bug is the most impactful technical issue. Together with Windows compatibility problems (#1061: PATHEXT, cp1252 encoding, select on pipes), these issues affect every skill developer. The community is investing significant effort to stabilize the skill-creator toolchain before focusing on new skills.  
🔗 https://github.com/anthropics/skills/issues/556  
🔗 https://github.com/anthropics/skills/issues/1061

### 3. Security & Trust Boundaries (#492 — 9 comments, 👍2)
Community-authored skills distributed under the `anthropic/` namespace create a trust boundary vulnerability. Users may grant elevated permissions to skills they believe are official. This is a structural concern for the ecosystem's growth.  
🔗 https://github.com/anthropics/skills/issues/492

### 4. Agent Governance & Safety (#412 — 6 comments)
A proposed skill for agent governance patterns: policy enforcement, threat detection, trust scoring, and audit trails. The skills ecosystem currently covers creative and technical workflows but has no safety governance layer.  
🔗 https://github.com/anthropics/skills/issues/412

### 5. Skills as MCPs (#16 — 4 comments)
A long-standing proposal to expose skills as Model Context Protocol endpoints. This would standardize skill APIs and enable composable AI software—essentially turning each skill into a callable function with typed inputs/outputs.  
🔗 https://github.com/anthropics/skills/issues/16

---

## 3. High-Potential Pending Skills

These PRs have active discussion and are likely to merge in the near term:

| Skill | PR | Why It's Likely to Land | 
|-------|-----|------------------------|
| **skill-creator run_eval fix (recall)** | [#1298](https://github.com/anthropics/skills/pull/1298) | Fixes a blocker bug affecting all skill optimization; multiple reproductions; highest urgency |
| **skill-creator trigger detection fix** | [#1323](https://github.com/anthropics/skills/pull/1323) | Companion to #1298; necessary for correct recall measurement |
| **Windows compatibility (subprocess + encoding)** | [#1050](https://github.com/anthropics/skills/pull/1050) | Two 1-line fixes; low complexity; addresses #1061 |
| **Docx tracked change w:id collision fix** | [#541](https://github.com/anthropics/skills/pull/541) | Fixes document corruption; minimal scope; clear root cause |
| **YAML unquoted description validation** | [#539](https://github.com/anthropics/skills/pull/539) + [#361](https://github.com/anthropics/skills/pull/361) | Prevents silent parsing failures; duplicate efforts suggest consensus |
| **document-typography** | [#514](https://github.com/anthropics/skills/pull/514) | High-value quality-of-life skill; complements existing document skills |
| **testing-patterns** | [#723](https://github.com/anthropics/skills/pull/723) | Fills a clear gap; comprehensive scope; no known blockers |

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is to stabilize the skill-creator toolchain (particularly the evaluation and optimization pipeline) before investing further in new skill content—with enterprise distribution, security, and agent governance emerging as the next frontier for skill architecture.**

---

# Claude Code Community Digest — 2026-06-23

## Today’s Highlights
- **v2.1.186** ships with new `claude mcp login/logout` commands for authenticating MCP servers from the CLI, plus status filtering in the `/workflows` agent.
- A long-standing permissions-matching bug (#30519) continues to draw community frustration (76 👍, 24 comments, no staff engagement since Sep 2025).
- A wave of duplicate issues filed on June 19–20 (rate limiting, clipboard permission, sandbox, plugin SSH) were all closed as duplicates, suggesting systematic regressions hit many users.

## Releases
**v2.1.186** [[release link](https://github.com/anthropics/claude-code/releases/tag/v2.1.186)]
- `claude mcp login <name>` / `claude mcp logout <name>` – authenticate MCP servers without opening the interactive `/mcp` menu. Supports `--no-browser` for stdin-redirect (SSH) flows.
- `/workflows` agent now supports status filtering (press `f`).

## Hot Issues (10 most noteworthy)

1. **[#30519 – Permissions matching fundamentally broken](https://github.com/anthropics/claude-code/issues/30519)**  
   *24 comments, 76 👍* – The core permission matching system has been broken since mid-2025. Over 30 related issues exist; Anthropic left a single workaround suggestion in September 2025 that didn’t resolve it. The community is building workarounds (e.g., custom PreTools). High visibility, unresolved for nearly 10 months.

2. **[#70182 – iOS app silently crashes on Remote Control session tap](https://github.com/anthropics/claude-code/issues/70182)**  
   *14 comments* – Fresh bug (filed today) affecting the Code tab. Crash occurs when tapping a Remote Control session. Duplicated in #70262 (iOS 27 Beta 2).

3. **[#13689 – Improve model’s ability to follow instructions](https://github.com/anthropics/claude-code/issues/13689)**  
   *13 comments* – Feature request asking for better adherence to user instructions in long contexts; a common pain point for complex workflows.

4. **[#42018 – Incorrectly charging VAT to EU B2B customers](https://github.com/anthropics/claude-code/issues/42018)**  
   *7 comments* – Billing bug: EU companies with valid VAT IDs are still being charged VAT. Related docs issue #51310 (missing VAT number entry).

5. **[#26725 – Stale git worktrees never cleaned up](https://github.com/anthropics/claude-code/issues/26725)**  
   *7 comments, 18 👍* – Claude Code creates worktrees for parallel tasks but has no garbage collection. After crashes or interruptions, worktrees accumulate indefinitely.

6. **[#69892 – Parabolic usage spike on June 20](https://github.com/anthropics/claude-code/issues/69892)**  
   *6 comments* – Sudden massive spike in API usage on June 20, 2026 at 7pm PT. Users on Windows/VSCode reported abnormal consumption.

7. **[#65982 – Fact-check gate at response-commit time](https://github.com/anthropics/claude-code/issues/65982)**  
   *6 comments* – Enhancement proposal to add a hook that verifies a verification action ran before the model asserts facts. Addresses hallucination risk in generated output.

8. **[#63839 – Session not found on disk after update to 1.9659.2](https://github.com/anthropics/claude-code/issues/63839)**  
   *4 comments* – Regression on macOS: only the most recent session is accessible; older sessions are “not found on disk”. Affects users who rely on multiple sessions.

9. **[#70279 – Opus model high latency and API errors](https://github.com/anthropics/claude-code/issues/70279)**  
   *2 comments* – Opus model in Claude Code has become “unusable” due to long pickup times, frequent API errors, and classifier failures over the last 2 days.

10. **[#69594 – Server temporarily limiting requests breaks workflows](https://github.com/anthropics/claude-code/issues/69594)**  
    *2 comments* – Rate limiting from the server side is causing long-running workflows to abort. One of many similar reports from June 19 (closed as duplicates, but the issue persists).

## Key PR Progress (4 of 4 open PRs)

- **[#70173 – fix(commit-commands): detect [gone] branches with `git branch -vv`](https://github.com/anthropics/claude-code/pull/70173)**  
  Fixes `/clean_gone` which never deletes any branches because `git branch -v` doesn’t show `[gone]`. Changed to `git branch -vv` and corrected grep pattern. Small but impactful fix for workspace hygiene.

- **[#63686 – Bump stale and autoclose timeouts from 14 to 90 days](https://github.com/anthropics/claude-code/pull/63686)**  
  Proposes extending the issue lifecycle from 14 days to 90 days to reduce premature closure of issues that may still be relevant – particularly important given the long-standing permissions issue (#30519).

- **[#70074 – docs: fix stale marketplace name in plugin-dev README](https://github.com/anthropics/claude-code/pull/70074)**  
  Replaces `claude-code-marketplace` → `claude-code-plugins` in plugin development documentation. References #70064.

- **[#70066 – docs(plugin-dev): update marketplace install docs](https://github.com/anthropics/claude-code/pull/70066)**  
  Updates install instructions to use the official marketplace name, replaces `cc --plugin-dir` examples with `claude --plugin-dir`, and clarifies contribution guidelines.

## Feature Request Trends
- **MCP authentication** – The new `claude mcp login` command addresses a long-standing need; users requested CLI-based auth without `/mcp` menus.
- **Fact verification hooks** – #65982 and similar requests point to a desire for programmable gates that run verification checks before the model commits responses.
- **Better session UX** – Requests for timestamps on messages (#69575), pinned prompt/status bar (#69588), and custom status line whitespace preservation (#69598) indicate UI polish priorities.
- **Plugin ecosystem reliability** – Multiple docs PRs (#70074, #70066) and a bug (#69587) highlight friction around plugin installation: SSH-only cloning, stale marketplace names, and missing HTTPS fallback.
- **Rate limit resilience** – Several feature requests ask for exponential backoff/retry logic (#69599) and better handling of server-side throttling (#69594).
- **Pricing flexibility** – #69601 requests a mid-tier plan (~$40-50/mo) between Pro and Max.

## Developer Pain Points
- **Permissions system broken for 10+ months** (#30519) – The lack of Anthropic engagement and community-built workarounds is the top frustration.
- **Rate limiting and API errors** – A spike in duplicate reports on June 19–20 (server limiting, usage draining 10x faster, Opus latency) indicates a systemic reliability issue.
- **Billing/VAT confusion** – EU B2B VAT errors (#42018) and missing VAT number fields (#51310) are costing users money and causing compliance headaches.
- **Stale worktree accumulation** (#26725) – Without garbage collection, disk usage grows unboundedly; a fix has been requested since February.
- **Sandbox incompatibility** (#69583) – `bwrap` fails on usrmerge Linux systems (WSL2, Ubuntu), making sandboxing effectively broken for many users.
- **Clipboard API denied in desktop preview** (#69576) – Embedded browser denies `clipboard.writeText()`, breaking web app previews.
- **Session persistence regressions** (#63839) – Updates that orphan old sessions erode trust in Claude Code’s reliability.

*Data sourced from github.com/anthropics/claude-code on 2026-06-23.*

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

Here is the OpenAI Codex community digest for June 23, 2026, structured for technical developers.

---

## OpenAI Codex Community Digest — 2026-06-23

### Today's Highlights

The biggest news is the **stable release of `rust-v0.142.0`**, which introduces a highly requested `/usage` credit system and a redesigned `/plugins` hub. This release also arrives alongside a critical fix for the **SQLite feedback log write amplification issue** (#28224), which was consuming up to 640 TB/year on some systems. A flurry of alpha releases for `rust-v0.143.0` is already underway, indicating rapid iteration on the next set of features, particularly around agentic and code-mode cells.

### Releases

A new stable release and six alpha releases landed in the last 24 hours.

- **`rust-v0.142.0` (Stable)**
  - **`/usage`**: Now supports earning, viewing, and redeeming usage-limit reset credits with a full confirmation and retry flow.
  - **`/plugins`**: Remote plugins are now organized into "OpenAI Curated," "Workspace," and "Shared with me" sections, improving discoverability.
- **`rust-v0.143.0-alpha.1` through `rust-v0.143.0-alpha.5`**: A rapid series of pre-release builds, likely containing early-stage work for features like the "continuing code-mode cell actor" and the "code mode host handshake protocol" seen in active PRs.

### Hot Issues

The top 10 most-discussed issues by comment volume highlight a mix of critical bugs and long-standing feature requests.

1.  **#11023 — [enhancement] Codex desktop app for Linux**
    - *Verdict:* Hot. The most-upvoted issue (628 👍). Users are desperate for a native Linux app, citing performance and power issues on macOS. The community is actively coordinating workarounds.
    - [Link](https://github.com/openai/codex/issues/11023)

2.  **#28879 — [bug] Rate-limit cost per token jumped ~10-20x**
    - *Verdict:* Critical. Users on the Plus plan report their budget draining in 2-3 prompts on `gpt-5.5` since June 16. Logs confirm a massive spike in per-token cost. This is a significant cost-of-service regression.
    - [Link](https://github.com/openai/codex/issues/28879)

3.  **#28224 — [bug] SQLite feedback logs write ~640 TB/year** (CLOSED)
    - *Verdict:* Resolved. A major SSD endurance and performance issue was fixed by merging PRs #29432 and #29457, reducing log volume by ~85%. The community is relieved but watching for the fix to land in stable builds. Despite being closed, it remains a top comment thread.
    - [Link](https://github.com/openai/codex/issues/28224)

4.  **#16404 — [enhancement] TUI voice transcription** (CLOSED)
    - *Verdict:* Regretted Removal. The TUI voice feature was removed in v0.118.0, and users are actively pushing for its return or a documented alternative. The desktop app's `Ctrl+M` is not seen as a replacement for a terminal-first workflow.
    - [Link](https://github.com/openai/codex/issues/16404)

5.  **#10665 — [enhancement] Native Azure DevOps Integration**
    - *Verdict:* Persistent Enterprise Request. A long-standing request for parity with the GitHub integration. The 69 👍 indicates strong enterprise demand. No movement yet.
    - [Link](https://github.com/openai/codex/issues/10665)

6.  **#29320 — [bug] Codex app only displays "Something went wrong…"**
    - *Verdict:* New & Blocking. A Windows store app build (26.616.6631.0) renders the app unusable. High-urgency for affected Windows users.
    - [Link](https://github.com/openai/codex/issues/29320)

7.  **#26910 — [bug] GPT 5.5 has a 404**
    - *Verdict:* Resolved Connectivity Issue. A model endpoint bug that caused `404 Not Found` errors when selecting GPT-5.5. Now closed, but it disrupted workflows for over two weeks.
    - [Link](https://github.com/openai/codex/issues/26910)

8.  **#14630 — [enhancement] Voice transcription for TUI**
    - *Verdict:* Redundant Request. A duplicate of #16404, further amplifying the demand for high-quality voice transcription in the CLI.
    - [Link](https://github.com/openai/codex/issues/14630)

9.  **#12773 — [enhancement] Multi-Window Support for macOS**
    - *Verdict:* High Demand for UX Parity. Users with complex projects need to manage multiple independent instances. This is a top UX feature request.
    - [Link](https://github.com/openai/codex/issues/12773)

10. **#18364 — [bug] Mac app hides older local conversations**
    - *Verdict:* Data Loss Concern. Updates are corrupting the local conversation history with bogus `status` sessions. This is a serious usability bug for Pro users with deep histories.
    - [Link](https://github.com/openai/codex/issues/18364)

### Key PR Progress

The most critical PRs in the last 24 hours focus on infrastructure stability, agent isolation, and plugin portability.

1.  **#29608 — Shut down superseded MCP managers on refresh**
    - *Fix:* Prevents zombie MCP processes from accumulating when managers are replaced during refresh. A direct fix for a resource leak.
    - [Link](https://github.com/openai/codex/pull/29608)

2.  **#29591 — feat(app-server): list descendant threads by ancestor**
    - *Feature:* Exposes a new API for clients to retrieve an entire spawned thread subtree in one request, reducing redundant client-side logic.
    - [Link](https://github.com/openai/codex/pull/29591)

3.  **#29566 — [code-reviewed] add continuing code-mode cell actor**
    - *Feature:* A new callback-only cell actor for persistent code execution. This is a foundational piece for more powerful, long-running agentic workflows.
    - [Link](https://github.com/openai/codex/pull/29566)

4.  **#29606 — core: pin yielded code cells to their request runtime**
    - *Fix:* Prevents a code cell that yields from silently switching to a wrong workspace or environment context when execution resumes. Critical for correctness in dynamic environments.
    - [Link](https://github.com/openai/codex/pull/29606)

5.  **#28918 — Make selected plugin roots URI-native**
    - *Feature:* Enforces `file://` URIs for plugin root paths. This is a necessary step for cross-platform plugin support and sandbox compatibility.
    - [Link](https://github.com/openai/codex/pull/28918)

6.  **#29067 — Namespace multi-agent v2 tools under collaboration**
    - *Feature:* Standardizes the tool namespace for multi-agent systems, making them more predictable for models.
    - [Link](https://github.com/openai/codex/pull/29067)

7.  **#29599 — Stop persisting bridged log events (CLOSED)**
    - *Fix:* A follow-up to #28224. The original SQLite fix missed bridged log events, which could still allow high-volume tracing to reach the database. This closes that gap.
    - [Link](https://github.com/openai/codex/pull/29599)

8.  **#29515 — [code-reviewed] define code mode host handshake protocol**
    - *Infrastructure:* Defines the formal protocol for communication between the host and the code execution sandbox. This is essential for security and feature negotiation.
    - [Link](https://github.com/openai/codex/pull/29515)

9.  **#29547 — core: use current step environments for tools**
    - *Fix:* Ensures that tools advertised to the model are always consistent with the environment state at the start of the current request, avoiding stale or phantom tools.
    - [Link](https://github.com/openai/codex/pull/29547)

10. **#28271 — Flatten MCP namespace tools for unsupported providers (CLOSED)**
    - *Fix:* Makes Codex's MCP tools compatible with third-party providers that don't support the proprietary `namespace` tool wrapper. Enables broader MCP interoperability.
    - [Link](https://github.com/openai/codex/pull/28271)

### Feature Request Trends

- **"Desktop Everywhere":** The call for a **native Linux app** (#11023) is the single most demanded feature. This is paired with requests for **multi-window support** (#12773) and **deep-link targeting** (#29605) to manage complex workflows.
- **Terminal-First Voice:** Despite the removal of TUI voice, the community is not backing down. There is a clear desire for **high-quality, server-side voice transcription** within the CLI, not just the desktop app.
- **Enterprise Git Integration:** **Azure DevOps (Azure Repos) integration** (#10665) remains the top enterprise feature request, reflecting Codex's growing use in non-GitHub environments.
- **Smarter Agentic Loops:** PRs like #29566 and #29606 signal a strong push toward **persistent, long-running code agents** that can yield, be pinned to contexts, and survive request boundaries.

### Developer Pain Points

- **"Cost Shock":** The #1 pain point is the **unpredictable rate-limit and cost behavior** (#28879). The 10-20x jump in per-token cost is causing significant frustration and breaking user trust in the pricing model.
- **Connectivity Instability:** A surge of issues related to **persistent 404 errors** (#26910, #28756), **Cloudflare challenges** (#29197), and **token exchange failures** (#26764) suggest an unstable backend API layer.
- **Session & State Management Bugs:** Problems like **conversation history corruption** (#18364), **"Item not found in turn state" errors** (#24263), and **stale app-server backends** (#29138) indicate ongoing issues with local state synchronization and session management.
- **Windows-Specific Fragility:** A disproportionate number of bugs target Windows: the **"Something went wrong" crash** (#29320), **sandbox helper failures** (#27125), **Computer Use plugin failures** (#28713), and **persistent SQLite churn** (#29570). Windows users are having a rougher experience than others.
- **The Slow SQLite Churn:** While issue #28224 is resolved, the community remains hyper-sensitive to **SQLite write amplification**. Any new logging regression is quickly flagged and scrutinized (see #29570). This has become a "will it blend?" test for the engineering team.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-06-23

## Today's Highlights
The team closed two high-severity SSRF vulnerabilities in the `web_fetch` tool via merged PRs, addressing DNS-hostname bypasses and redirect-based private IP access. A critical OAuth token-exchange fix was opened to resolve a Node.js 24.17.0 socket-reuse regression that broke "Sign in with Google" for users on newer runtimes. Meanwhile, the `web_fetch` SSRF patches were backported and finalized, and a workflow-security remediation PR removed excessive trust from public CI pipelines.

## Releases
No new releases were published in the last 24 hours.

## Hot Issues

1. **#21409 — Generalist agent hangs** [P1, 8👍]  
   The most-upvoted open bug: the CLI hangs indefinitely when deferring to the generalist agent for simple tasks like folder creation. Users confirm that instructing the model not to use sub-agents is the only workaround.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/21409)

2. **#22323 — Subagent recovery after MAX_TURNS falsely reports GOAL success** [P1, 2👍]  
   `codebase_investigator` subagent returns `status: "success"` / `Termination Reason: "GOAL"` even when it hit the turn limit before doing any work, masking failures.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/22323)

3. **#24353 — Robust component-level evaluations** [P1, EPIC]  
   Tracks expansion of the behavioral-eval framework from 76 tests across 6 models. The community expects this to become the gate for agent-quality regression detection.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/24353)

4. **#25166 — Shell command gets stuck "Waiting input" after completion** [P1, 3👍]  
   After executing trivial shell commands, the CLI hangs showing "Awaiting user input" even though the command already finished. A common daily frustration.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/25166)

5. **#22745 — Assess AST-aware file reads, search, and mapping** [P2, 1👍]  
   EPIC investigating whether AST-aware tools (tilth, glyph) can reduce turns and token noise by precisely reading method bounds and improving codebase mapping.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/22745)

6. **#26525 — Add deterministic redaction and reduce Auto Memory logging** [P2]  
   Auto Memory sends transcripts to the model before redacting secrets—content is already exposed by the time redaction runs. The issue also flags excessive logging of skill definitions.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/26525)

7. **#21968 — Gemini does not use skills and sub-agents enough** [P2]  
   Anecdotal but consistent: custom skills and sub-agents are ignored unless explicitly instructed, even for closely related tasks. Hinders the value of user-authored agents.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/21968)

8. **#21983 — Browser subagent fails on Wayland** [P1, 1👍]  
   The browser agent terminates with "GOAL" but fails silently under Wayland display servers. Linux users are blocked from web-automation workflows.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/21983)

9. **#22672 — Agent should discourage destructive behavior** [P2, 1👍]  
   Model sometimes uses `git reset --force` or destructive DB commands when safer alternatives exist. Users want situational awareness around irreversible operations.  
   [Issue](https://github.com/google-gemini/gemini-cli/issues/22672)

10. **#22465 — Gemini CLI stuck at interactive prompt creating Vite app** [P2]  
    The agent enters an interactive prompt loop during `npm create vite` and never escapes. A behavioral eval is requested.  
    [Issue](https://github.com/google-gemini/gemini-cli/issues/22465)

## Key PR Progress

1. **#28103 — Fix OAuth keep-alive socket reuse on Node >=24.17.0** [P2]  
   New PR addressing a Node.js regression where pooled sockets cause `ERR_STREAM_PREMATURE_CLOSE` during "Sign in with Google". Disables keep-alive for the token-exchange request.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28103)

2. **#27744 — SSRF: resolve DNS before guard to block hostname-to-private-IP bypass** [CLOSED]  
   Merged. The old `isPrivateIp()` only checked IP-literal hostnames, missing DNS rebinding attacks via services like `127.0.0.1.nip.io`. Now resolves DNS before checking.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/27744)

3. **#27739 — SSRF: prevent private targets via DNS hostnames and redirects** [CLOSED]  
   Merged. Fixes two gaps: hostnames that resolve to private IPs, and redirect chains that redirect to internal addresses.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/27739)

4. **#28098 — Remove excessive CI trust from public workflows** [CLOSED]  
   Remediated potential secret exfiltration by removing `GEMINI_CLI_TRUST_WORKSPACE: 'true'` and restricting tools in workflows processing untrusted issue/PR inputs.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28098)

5. **#28096 — Drop late tool calls after SIGINT cancellation** [P2]  
   Fixes a race where a SIGINT arrives after the stream starts a turn but before tool execution finishes—the tool result still gets submitted. Now discards late calls.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28096)

6. **#27936 — Missing VS Code Disposables due to comma operator** [P2]  
   Two `context.subscriptions.push()` calls wrapped registrations in parentheses, turning them into comma expressions that silently dropped the first Disposable of each pair.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/27936)

7. **#28100 — Same comma-operator leak in VS Code extension** [P2]  
   Parallel fix for the same pattern found in a second location. Both PRs ensure all disposables are registered for clean shutdown.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28100)

8. **#28105 — Fix ellipsis logic in EditTool description** [P1]  
   Corrects a substring-off-by-one bug that could produce incorrect snippet previews in edit-tool descriptions.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28105)

9. **#28000 — Fix Jupyter Notebook and JSON corruption in write_file** [CLOSED]  
   `write_file` silently corrupted `.ipynb` and `.json` files. Now correctly handles non-text binary-safe writes, preventing data loss in Colab/Jupyter environments.  
   [PR](https://github.com/google-gemini/gemini-cli/pull/28000)

10. **#28094 — Deep-merge user and workspace settings in A2A server** [P2]  
    `loadSettings()` used a shallow spread, so nested sections like `tools` or `telemetry` from workspace settings completely replaced user settings. Now deep-merges.  
    [PR](https://github.com/google-gemini/gemini-cli/pull/28094)

## Feature Request Trends

- **AST-aware tooling for codebase understanding**: Multiple issues (#22745, #22746) call for integrating AST-based CLI tools to improve method-bound reading, search precision, and codebase mapping—reducing token waste and turn count. Tied to the `codebase_investigator` subagent.
- **Subagent orchestration transparency**: Feature requests (#22598, #21432) ask for subagent trajectories to be visible via `/chat share` and for the CLI to understand its own mechanics (flags, hotkeys) well enough to act as its own guide.
- **Memory system robustness**: Issues (#26522, #26523, #26516) demand smarter handling of low-signal sessions (stop retrying indefinitely), quarantine of invalid patches, and overall memory-quality improvements.
- **SSRF and security hardening**: The flurry of web-fetch fixes (#27744, #27739) plus requests for deterministic redaction (#26525) and workflow-trust remediation (#28098) show a strong push toward sandboxing and input-validation rigor.
- **Browser agent improvements**: Feature requests (#22232, #22267) ask for automatic session takeover, lock recovery, and respect for `settings.json` overrides like `maxTurns`.

## Developer Pain Points

- **Agent hangs and false success reports**: The generalist agent hangs (#21409), subagents falsely reporting GOAL success after MAX_TURNS (#22323), and infinite thought loops (#16342) erode trust in autonomous modes.
- **Shell command deadlock**: Commands completing but leaving the CLI stuck in "Waiting input" (#25166) is a frequent workflow killer, especially for CI/headless users.
- **Subagent/skill underutilization**: Despite having custom skills and agents configured, the model rarely uses them without explicit instruction (#21968). Combined with subagents executing without permission after updates (#22093), configuration intent is often lost.
- **Terminal and editing glitches**: Corruption after exiting external editors (#24935), flicker on resize (#21924), and broken `\n` escaping (#22466) degrade the interactive experience.
- **Tool-limitation errors**: A 400 error surfaces when more than ~128 tools are loaded (#24246), plus model tendency to scatter temp scripts across the filesystem (#23571) adds cleanup overhead.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-06-23

## Today's Highlights
Two patch releases shipped today, adding user‑configurable HTTP(S) proxies, inline image rendering, and a setting to hide the conversation scrollbar. Meanwhile, the community is calling attention to authentication errors when resuming sessions (#3596, 11 👍) and a long‑standing request for skill subfolder support (#1632, 20 👍). A fresh triage issue reports that OpenAI documentation URLs fail with `WebFetchRedirectError` (#3890).

## Releases
- **[v1.0.64-3](https://github.com/github/copilot-cli/releases/tag/v1.0.64-3)**  
  *Added*: HTTP(S) proxy setting.  
  *Fixed*: Resume sessions by name containing spaces; hide unsupported slash commands in remote‑hosted sessions.

- **[v1.0.64-2](https://github.com/github/copilot-cli/releases/tag/v1.0.64-2)**  
  *Added*: Setting to hide conversation scrollbar; inline image rendering; argument‑hint frontmatter for skills; OpenTelemetry compaction with `gen_ai.conversation.compacted=true`.

## Hot Issues (10 notable items)
1. **[#1944 — Windows mouse scroll regression (CLOSED)](https://github.com/github/copilot-cli/issues/1944)**  
   Mouse wheel scrolling no longer navigates conversation history on Windows; scroll events are captured by the input box.  
   *10 comments, 3 👍* – a noticeable UX regression that has now been resolved.

2. **[#1632 — Subfolders for skills (OPEN)](https://github.com/github/copilot-cli/issues/1632)**  
   Users with 10+ custom skills want subfolder organisation. The CLI currently enforces a flat structure.  
   *8 comments, 20 👍* – the most upvoted open feature request.

3. **[#2486 — MCP server blocked by policy (CLOSED)](https://github.com/github/copilot-cli/issues/2486)**  
   A personal Pro+ user saw MCP servers incorrectly blocked after weeks of normal use. Workaround known but no official reply.  
   *7 comments* – highlights fragility in MCP policy validation.

4. **[#3162 — Registry-listed MCP servers falsely flagged as blocked (CLOSED)](https://github.com/github/copilot-cli/issues/3162)**  
   v1.0.42 reports custom MCP servers from the registry as “blocked by policy”, a false‑negative in registry matching.  
   *7 comments, 1 👍* – second policy‑related bug, now fixed.

5. **[#3596 — “Not authenticated” on session resume (OPEN)](https://github.com/github/copilot-cli/issues/3596)**  
   Resuming a specific session fails with `Error loading model list: Error: Not authenticated`; starting a new session works.  
   *6 comments, 11 👍* – high impact; affects users heavily relying on persistent sessions.

6. **[#3866 — Thinking/reasoning text unreadable on dark backgrounds (OPEN)](https://github.com/github/copilot-cli/issues/3866)**  
   “Thinking…” text uses a hardcoded dark gray foreground that is nearly invisible on dark terminals.  
   *1 comment, 2 👍* – accessibility issue in the latest release.

7. **[#2590 — Plugins from Marketplace not available via ACP (OPEN)](https://github.com/github/copilot-cli/issues/2590)**  
   Plugins visible in CLI are invisible when the same installation is accessed through the Agent Client Protocol (ACP).  
   *1 comment, 3 👍* – important gap for users of both CLI and ACP.

8. **[#1579 — Copilot CLI ignores MCP server “instructions” (OPEN)](https://github.com/github/copilot-cli/issues/1579)**  
   MCP servers return instructions during initialization that should be fed to the LLM, but Copilot CLI discards them.  
   *3 👍* – reduces effectiveness of MCP‑powered workflows.

9. **[#3890 — WebFetchRedirectError for OpenAI docs (OPEN, triage)](https://github.com/github/copilot-cli/issues/3890)**  
   Fetching certain OpenAI documentation URLs fails because 301 redirects are not followed.  
   *New today* – affects users who rely on web fetch for reference material.

10. **[#3888 — Extended thinking as independent control (OPEN, triage)](https://github.com/github/copilot-cli/issues/3888)**  
    For Anthropic models, thinking and reasoning effort are separate API parameters, but the CLI only exposes effort.  
    *New today* – feature request for finer‑grained control.

## Key PR Progress
Only one PR was updated in the last 24 hours:

- **[#3873 — Add initial console log for greeting (OPEN)](https://github.com/github/copilot-cli/pull/3873)**  
  Adds a simple console log when the CLI starts. No comments, no reactions – appears to be a minor cosmetic improvement.  
  PR activity remains low; the main development focus is on issue resolution and patch releases.

## Feature Request Trends
The most‑requested directions emerging from recent issues include:

- **MCP Ecosystem Maturity** – Better policy validation, support for server “instructions”, variable interpolation during `/mcp install`, and stdio transport support in ACP.  
- **Performance Feedback** – Multiple requests (#3278, #3111, #3055) for timer/elapsed‑time indicators during agent thought and shell tool execution.  
- **Skills Organisation** – Subfolder support (#1632) remains the top‑voted feature.  
- **Accessibility & Theming** – Making reasoning text readable on dark backgrounds (#3866).  
- **Extended Model Controls** – Independent toggle for extended thinking on Anthropic models (#3888).

## Developer Pain Points
- **Windows scroll regression** – Recently fixed, but the impact was widespread.  
- **Permission prompts for harmless commands** – Redirects like `2>/dev/null` still trigger unnecessary access dialogs (#2693).  
- **Authentication instability on session resume** – High‑impact authenticated‑state loss (#3596).  
- **AI credit consumption on restart** – Using `/restart` consumes ~174 credits, surprising users (#3886).  
- **MCP server instructions ignored** – Core MCP feature unused, limiting LLM effectiveness (#1579).  
- **Plugin/agent isolation** – Plugins from CLI not available in ACP or VS Code (#2590, #3638).  
- **Hardcoded colour for reasoning text** – Breaks dark‑theme usability (#3866).

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-06-23

## Today’s Highlights

A new minor release (v1.48.0) shipped with a fix for empty reasoning content and a smarter force-stop mechanism for repeated tool-call loops. Meanwhile, three open bugs surfaced: a yolo‑mode approval prompt inconsistency, an MCP server path issue in `kimi web`, and a hang after detached child‑process tool calls. A feature PR proposing a per‑line streaming `Monitor` tool also entered review.

---

## Releases

**v1.48.0** — [Release on GitHub](https://github.com/MoonshotAI/kimi-cli/releases/tag/1.48.0)  
* Fix: Round‑trip empty reasoning content in kosong (#2446)  
* Feature: Escalate repeated‑tool‑call reminders and force‑stop on dead‑end streak (#2466)  
* Internal: Changelog entries omitted for this release  

The force‑stop logic is particularly valuable for long‑running agent loops where the model gets stuck re‑invoking the same tool – now `kimi` will escalate warnings and eventually kill the streak, preventing infinite hangs.

---

## Hot Issues

*Only 3 issues were updated in the last 24h. All are listed below.*

1. **#2448 – [bug] Kimi CLI prompting for approval in yolo mode**  
   [Issue](https://github.com/MoonshotAI/kimi-cli/issues/2448)  
   Opened 2026-06-10, last updated 2026-06-23, 1 comment  
   *Why it matters:* Yolo mode is designed for zero‑approval execution. A user reports continuous approval prompts despite being in yolo mode with an API key. This breaks the core promise of the feature. Community reaction is minimal (no thumbs), but the bug has been open for 13 days, suggesting it may be a low‑priority or tricky regression.

2. **#2469 – [bug] `kimi web` starts MCP servers from the CLI installation directory**  
   [Issue](https://github.com/MoonshotAI/kimi-cli/issues/2469)  
   Opened 2026-06-22, no comments  
   *Why it matters:* MCP (Model Context Protocol) servers configured per workspace should resolve relative paths correctly. Starting them from the global install directory breaks workspace‑relative tool paths, a common source of confusion for teams using project‑specific MCP configurations. No community response yet.

3. **#2468 – [bug] Kimi CLI hangs after detached child‑process tool call**  
   [Issue](https://github.com/MoonshotAI/kimi-cli/issues/2468)  
   Opened 2026-06-22, no comments  
   *Why it matters:* Detached child processes (launched with `&` or `--detach`) should not block the parent CLI. A hang suggests a missing background‑job cleanup or process‑group signal handling issue. The reporter uses a local mock provider, so the problem may be provider‑agnostic. No community engagement yet.

---

## Key PR Progress

*Only 2 PRs were updated in the last 24h. Both are listed below.*

1. **#2471 – [OPEN] feat(tools): add Monitor tool for per‑line stdout streaming**  
   [PR](https://github.com/MoonshotAI/kimi-cli/pull/2471)  
   Author: @Nitjsefnie | Created 2026-06-22  
   *Summary:* Proposes a `Monitor` tool that streams stdout line‑by‑line, acting as a real‑time counterpart to the existing background‑tool. Useful for long‑running processes (e.g., build logs, tail‑f). No maintainer comments yet; community interest unknown.

2. **#2467 – [CLOSED] chore(release): bump kimi-cli to 1.48.0 and kosong to 0.54.0**  
   [PR](https://github.com/MoonshotAI/kimi-cli/pull/2467)  
   Author: @sailist | Merged 2026-06-22  
   *Summary:* Version bump PR that cut the 1.48.0 release. Includes validation scripts for tag consistency. No changelog entries (“internal changes”).

---

## Feature Request Trends

Based on the limited data, two emerging directions can be discerned:

* **Per‑line streaming tools** – PR #2471 proposes a `Monitor` tool for real‑time stdout output. This aligns with a growing need for better observability when the CLI acts as an agent that spawns long‑running processes.
* **Force‑stop / early termination** – The escalation behaviour in v1.48.0 (#2466) addresses a common pain point: agent loops spinning on repeated tool calls. This is a feature that makes autonomous mode safer.

No explicit feature requests appear in the open issues; all three are bug reports.

---

## Developer Pain Points

Three recurring frustrations surface from today’s bugs:

1. **Yolo mode unreliability** – Issue #2448 shows that the “zero‑approval” mode still prompts, undermining trust in autonomous operation.  
2. **Workspace path resolution** – #2469 highlights that `kimi web` fails to respect project‑relative MCP configurations, forcing users to work around global paths.  
3. **Hangs on background processes** – #2468 demonstrates that detached child‑process tool calls can block the whole CLI, a critical issue for users who rely on concurrent tool execution (e.g., running tests while continuing conversation).

These issues reflect a broader need for robust process lifecycle management and consistent configuration scoping. The small number of reports (0–1 comments each) suggests either low user volume or early stages of discovery.

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-23

## Today’s Highlights

No new releases were published in the last 24 hours, but the community saw a flurry of activity across 24 issues and 80+ PRs. Two long‑standing MCP resource‑related issues (#15535, #17543) were finally closed thanks to PR #33483, which adds model‑callable resource read tools. Meanwhile, a critical Copilot API endpoint bug (#31000) and a paid‑balance–still‑hitting‑free‑limit issue (#33318) continue to spark debate and demand maintainer attention.

## Releases

*None.* No new versions were pushed in the last 24 hours.

## Hot Issues (10 notable)

1. **[#31000](https://github.com/anomalyco/opencode/issues/31000) – Copilot provider model‑list fetch fails for github.com users**  
   *Status: Open | 👍 0*  
   The `d7()` function constructs a non‑existent domain for GitHub.com users. Correct endpoint should be `api.githubcopilot.com`. Expected to break model discovery for all free Copilot users.

2. **[#25832](https://github.com/anomalyco/opencode/issues/25832) – opencode cannot read images anymore**  
   *Status: Open | 👍 4*  
   Regression: image input (PNG, JPG) stopped working after April 29. Users report `Bad` error when asking the AI to interpret an image. High impact for visual‑feedback workflows.

3. **[#27755](https://github.com/anomalyco/opencode/issues/27755) – TypeError: Failed to fetch shortly after opening**  
   *Status: Open | 👍 3*  
   App becomes unusable immediately after launch. No prompts can be sent. Multiple reports, likely a runtime initialization race condition.

4. **[#33318](https://github.com/anomalyco/opencode/issues/33318) – [URGENT] Zen paid balance still hits FreeUsageLimitError**  
   *Status: Open | 👍 0*  
   Users with $20+ paid balance still get “Free usage exceeded” after less than 1 hour. Billing bypass logic appears broken–highly disruptive for paying users.

5. **[#33077](https://github.com/anomalyco/opencode/issues/33077) – Bash tool has no protection against destructive commands**  
   *Status: Closed | 👍 0*  
   Permission system only checks the command string as a resource string, not its semantics. `Allow always` can bypass deny rules. Community called for a built‑in safe‑command validator.

6. **[#32200](https://github.com/anomalyco/opencode/issues/32200) – `zsh: trace trap` on launch – macOS 15.3.1 Apple Silicon**  
   *Status: Open | 👍 0*  
   `EXC_BREAKPOINT / SIGTRAP` crash caused by pointer authentication (PAC) trap. Affects Apple Silicon users. Maintainers need to inspect binary signing or compiler flags.

7. **[#22142](https://github.com/anomalyco/opencode/issues/22142) – Repetitive tool‑call loops with `alibaba-coding-plan-cn/qwen3.6-plus`**  
   *Status: Open | 👍 1*  
   Adjacent assistant turns contain identical reasoning, leading to runaway token usage and stalled sessions. Model‑specific pattern likely requires provider‑side tuning.

8. **[#33491](https://github.com/anomalyco/opencode/issues/33491) – Windows: deleted tracked files trigger ENOENT toasts and freeze renderer**  
   *Status: Open | 👍 0*  
   When deleting many tracked files on Windows, repeated errors and renderer freezes. Desktop v1.17.9 regression affecting Git repo housekeeping.

9. **[#32835](https://github.com/anomalyco/opencode/issues/32835) – Review panel: inline comments on diffs broken after v1.17.8**  
   *Status: Closed | 👍 8*  
   “+” button to add inline comments disappeared in Web UI and Desktop. High community engagement (8 👍) and quickly closed, possibly already fixed.

10. **[#31540](https://github.com/anomalyco/opencode/issues/31540) – Permission deny rules overridden by ‘Allow always’ approval**  
    *Status: Open | 👍 0*  
    Clicking “Allow always” stores `{permission: "edit", pattern: "*"}` which overrides all deny rules for the session. Essential for security‑conscious users.

## Key PR Progress (10 important)

1. **[#33483](https://github.com/anomalyco/opencode/pull/33483) – feat(mcp): add resource read tools**  
   *Merged | Fixes #15535, #17543*  
   Adds model‑callable MCP resource list/read tools, treats URIs as opaque identifiers, and limits binary attachments to safe payloads. Long‑requested feature now shipped.

2. **[#33489](https://github.com/anomalyco/opencode/pull/33489) – feat(i18n): complete Turkish localization and sync script**  
   *Merged*  
   Full user‑verified Turkish locale for the desktop app, plus an automated sync script to keep translations up‑to‑date. Good model for future i18n contributions.

3. **[#33482](https://github.com/anomalyco/opencode/pull/33482) – fix(acp): bridge question prompts via extMethod**  
   *Open*  
   Fixes `question` tool hanging forever in ACP mode by forwarding answers via `extMethod`. Closes #17920 and #13752. Important for remote agent workflows.

4. **[#33365](https://github.com/anomalyco/opencode/pull/33365) – feat(tui): add diff viewer keybind**  
   *Open*  
   Adds a keyboard binding to open the diff viewer from the TUI. Improves keyboard‑first workflow.

5. **[#33384](https://github.com/anomalyco/opencode/pull/33384) – [contributor] feat(app): collapsible servers**  
   *Merged*  
   Makes server rows on the homepage collapsible, reducing visual clutter. Community‐contributed UX polish.

6. **[#33267](https://github.com/anomalyco/opencode/pull/33267) – fix(tui): improve worker RPC error handling**  
   *Merged | Closes #33269*  
   Ensures thrown errors in the Bun worker are properly sent back to the caller instead of silently failing.

7. **[#32797](https://github.com/anomalyco/opencode/pull/32797) – feat(app): add mobile bottom navigation**  
   *Merged*  
   Persisted mobile‑only setting to move navigation to the bottom, with safe‑area inset handling. Major improvement for one‑handed use.

8. **[#32798](https://github.com/anomalyco/opencode/pull/32798) – fix(app): improve iOS PWA shell**  
   *Merged*  
   Edge‑to‑edge rendering, updated launch colors, and correct PWA meta tags so internal links stay inside the installed app.

9. **[#32662](https://github.com/anomalyco/opencode/pull/32662) – [contributor] feat(app): new session progress indicator**  
   *Merged*  
   Implements a “dot matrix radar” animation while the agent is running. Community contribution enhancing visual feedback.

10. **[#30942](https://github.com/anomalyco/opencode/pull/30942) – feat(tui): add main branch source to diff mode**  
    *Open*  
    Adds a diff source that shows all changes against the main branch, enabling branch‑wide reviews from the TUI.

## Feature Request Trends

Based on the latest issues, the community is consistently asking for:

- **Better billing/usage visibility** – Users want a Zen dashboard with per‑model breakdowns and invoice download (#13497, 8 👍). The “paid still hits free limit” bug (#33318) amplifies demand.
- **MCP resource integration** – MCP resource read support (#15535) was just delivered, but users still desire tighter integration (e.g., rich @‑mentions for binary resources).
- **Localization** – Vietnamese (#29309) joins Turkish (now landed via PR #33489) as a community‑driven i18n push.
- **UI/UX refinements** – Always‑fixed sidebar (#33487), collapsible servers (#33384), and mobile bottom navigation (#32797) show appetite for a polished desktop/web experience.
- **Permission system hardening** – Issues #31540 and #33077 expose a weak permission model; users want deny‑first rules and command‑level safety analysis.

## Developer Pain Points

Several recurring frustrations cut across multiple issues:

- **Billing/credits confusion** – “Free usage exceeded” even with paid balance (#14273, #33318) and unclear daily limits lead to workflow interruptions.
- **Model‑specific regressions** – GLM‑5.2 rejects `instructions` field (#33490) and returns 400 for content format (#32821). Tool‑call loops with Qwen (#22142) suggest incomplete provider parity.
- **File handling fragility** – Image input stopped working (#25832), Windows delete operations cause freezes (#33491), and the @ file picker ignores `.ignore` negation patterns (#31801).
- **Permission bypass risks** – “Allow always” overrides deny rules (#31540), and Bash destructive commands have no guard (#33077). Security‑conscious users lack granular control.
- **Cross‑platform crashes** – macOS PAC trap (#32200) and TUI “Failed to fetch” (#27755) indicate unstable runtime environments for specific configurations.
- **Regression in key features** – Review panel inline comments broken after v1.17.8 (#32835) and keyboard shortcuts incompatible with Korean IME (#21163) show gaps in regression testing for non‑English locales.

---

*Generated from GitHub data at [anomalyco/opencode](https://github.com/anomalyco/opencode).*

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-06-23

## Today’s Highlights
The community is grappling with **intermittent OpenAI Codex connection stalls** (the top-voted issue this week) and a **long-standing demand for an official local LLM provider** extension. Meanwhile, two new providers (Merge Gateway and Anthropic Vertex) have landed as pull requests, and several UX fixes for the TUI – including URL linkification and dialog flicker – are moving toward merge.

## Releases
**No new versions published in the last 24 hours.** The latest release remains `pi-coding-agent 0.79.1` (referenced in issue #5571).

## Hot Issues (10 most noteworthy)

1. **[#4945 – OpenAI Codex Connection Reliability](https://github.com/earendil-works/pi/issues/4945)**  
   *66 comments, 30 👍*  
   The TUI often freezes on “Working…” with no streamed text, tool call, or error – only pressing Escape recovers. Community reports this happening repeatedly over the last few days, making the Codex provider unreliable for interactive sessions.

2. **[#3357 – Official local LLM provider extension](https://github.com/earendil-works/pi/issues/3357)**  
   *27 comments, 36 👍*  
   Users want Pi to dynamically fetch model lists from `{baseUrl}/models` (e.g. for llama.cpp, Ollama, LM Studio). High demand indicates a strong preference for self‑hosted LLMs, especially for privacy and cost control.

3. **[#5571 – `pi -p` hangs indefinitely on non‑TTY pipe](https://github.com/earendil-works/pi/issues/5571)**  
   *10 comments*  
   When stdin is a pipe that never closes and the default provider lacks credentials, `pi -p` blocks for minutes instead of failing fast. A significant usability issue for CI/CD pipelines.

4. **[#5989 – `pi update` broke extension `pi-lovely-codex`](https://github.com/earendil-works/pi/issues/5989)**  
   *6 comments*  
   A routine update made an existing extension incompatible, causing `Failed to load extension`. Points to the need for better backward‑compatibility guarantees or an extension lockfile.

5. **[#5871 – Anthropic OAuth‑token detection is hardcoded](https://github.com/earendil-works/pi/issues/5871)**  
   *6 comments*  
   The provider only recognises OAuth tokens starting with `sk-ant-oat`. This inflexibility breaks setups with custom or self‑issued tokens, and a configurable `authMode` override is being requested.

6. **[#5291 – Sessions hang on “Working” with Anthropic subscription](https://github.com/earendil-works/pi/issues/5291)**  
   *6 comments, 2 👍*  
   Users with Anthropic Enterprise subscriptions experience intermittent session freezes, all at once. Interrupt/resume works inconsistently – a critical reliability concern for paying customers.

7. **[#5990 – TUI flickers when dialog content exceeds terminal height](https://github.com/earendil-works/pi/issues/5990)**  
   *4 comments*  
   `ctx.ui.confirm()` and similar dialogs cause continuous repaint when the content is taller than the viewport. A UI bug that makes long confirmation prompts unusable.

8. **[#5909 – Coalesce rapid `thinking_level_change` entries to avoid session bloat](https://github.com/earendil-works/pi/issues/5909)**  
   *3 comments*  
   Rapidly cycling the thinking level appends many invisible entries to session JSONL files, bloating the file and slowing down the TUI. A performance optimisation that matters for heavy users.

9. **[#5978 – Plain long URLs lose clickability after TUI wrapping](https://github.com/earendil-works/pi/issues/5978)**  
   *3 comments*  
   Long URLs break across terminal rows; Ctrl‑click only sees the first fragment. This affects OAuth flows from MCP extensions. A fix (linkification via OSC 8) is already in PR #5981.

10. **[#5953 – `models.json` requires `apiKey` even when `auth.json` has credentials](https://github.com/earendil-works/pi/issues/5953)**  
    *3 comments*  
    Custom providers defined in `models.json` force an `apiKey` field, ignoring valid credentials already stored in `auth.json`. Creates unnecessary configuration duplication.

## Key PR Progress (10 important PRs)

1. **[#5994 – fix(ai): route OpenCode Go models through Anthropic](https://github.com/earendil-works/pi/pull/5994)**  
   Redirects OpenCode Go endpoint models (e.g. `minimax-m2.7`) to the Anthropic Messages API instead of OpenAI chat‑completions, fixing compatibility for those models.

2. **[#5526 – Require terminal events for OpenAI Responses streams](https://github.com/earendil-works/pi/pull/5526)**  
   Prevents stalled streams by enforcing that OpenAI Responses must end with a terminal event before completing – solves the frequent “context counter borked” issue and the need to type “continue.”

3. **[#5987 – fix(coding-agent): resolve `--session` by agent name via identity daemon](https://github.com/earendil-works/pi/pull/5987)**  
   Enables the `--session` flag to accept an agent name (e.g. `lucid-gecko-24`) and resolve it to the correct session file path using the identity daemon.

4. **[#5859 – fix(ai): send responses prompts as instructions](https://github.com/earendil-works/pi/pull/5859)**  
   Corrects OpenAI Responses API usage: system prompts go into top‑level `instructions` rather than replayed `input` messages, fixing compatibility with Azure OpenAI and Codex.

5. **[#5985 – feat(ai): add Merge Gateway provider](https://github.com/earendil-works/pi/pull/5985)**  
   Integrates Merge Gateway as a built‑in OpenAI‑compatible provider, giving Pi access to 40+ models from a single API key.

6. **[#5981 – Linkify plain URLs in Text output](https://github.com/earendil-works/pi/pull/5981)**  
   Fixes issue #5978: automatically wraps plain URLs with OSC 8 hyperlinks when the terminal supports it, restoring clickability for OAuth and other long URLs.

7. **[#5979 – Fix `session-id-readonly.test.ts` on clean main](https://github.com/earendil-works/pi/pull/5979)**  
   Mocks the API key pre‑flight check so the test can reach actual session‑ID rejection logic without requiring real credentials – unblocks local development.

8. **[#5977 – feat(ai): allow explicit `authMode` overrides for Anthropic provider](https://github.com/earendil-works/pi/pull/5977)**  
   Introduces a `authMode` compatibility flag to override the hardcoded `sk-ant-oat` check, addressing issue #5871 and supporting custom token formats.

9. **[#5262 – feat(ai): add Anthropic Vertex provider](https://github.com/earendil-works/pi/pull/5262)**  
   Adds a built‑in `anthropic-vertex` provider for Claude on Google Cloud Vertex AI, reusing the existing Anthropic streaming path – requested by enterprise users.

10. **[#5970 – feat: add auto-router extension for DeepSeek V4 cost optimisation](https://github.com/earendil-works/pi/pull/5970)**  
    A new extension that automatically routes simple tasks to DeepSeek V4 Flash and complex ones to V4 Pro, promising 60‑70% API cost savings.

## Feature Request Trends

- **More provider flexibility:** Requests for local LLM (Ollama/LM Studio), Neuralwatt (GLM, Kimi, Qwen), Merge Gateway, and Anthropic Vertex are recurring. The community wants to plug in any OpenAI‑compatible or Anthropic‑compatible endpoint with minimal config.
- **Configuration without redundancy:** Users are pushing to eliminate duplicate credential storage (`models.json` vs. `auth.json`) and to make OAuth detection configurable rather than hardcoded.
- **Better UX for pipes and non‑interactive mode:** The `pi -p` hang, long URL handling, and the need for `--model provider/model` syntax all reflect pain points when Pi is used in scripts or CI.
- **Extension API hardening:** Issues like #5997 (silent failure on Node 24) and #5751 (fire‑and‑forget sendMessage) indicate demand for a more robust, well‑documented extension API with proper async support and error reporting.
- **TUI stability improvements:** Flickering dialogs, session name newlines, and session bloat from rapid thinking changes are small but high‑impact annoyances that the community wants addressed before the next release.

## Developer Pain Points

- **Connection reliability:** OpenAI Codex and Anthropic subscriptions both show intermittent “Working…” freezes with no clear recovery path – top concern.
- **Upgrade regressions:** A simple `pi update` broke a popular extension (#5989), undermining trust in the update process.
- **Hardcoded provider logic:** OAuth token detection and model routing (OpenCode Go) require code changes instead of configuration, frustrating self‑hosters.
- **Test environment friction:** Missing API key mocks cause test suites to fail on clean clones (#5979, #5982), a barrier for new contributors.
- **SDK crashes:** Unhandled `write EPIPE` after heavy tool runs (#5993) and `value.startsWith` crashes on reload (#5992) reveal unhandled edge cases in the SDK and TUI.
- **Package removal not working:** `pi remove` reports success but the extension remains in config (#5966) – a basic CLI command that fails silently.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

好的，作为专注于AI开发者工具的技术分析师，根据提供的GitHub数据，以下是 Qwen Code 社区日报（2026-06-23）。

---

### Qwen Code 社区日报 | 2026年6月23日

#### 1. 今日亮点
今日焦点是架构层面的持续演进和开发者体验优化。随着 `v0.19.1` 的发布，MCP 服务器发现和资源补全功能得到增强。社区对协议/认证解耦的讨论（#5758）尤为活跃，这关乎未来的可扩展性。同时，针对自动模式下的安全性（#5749）和大模型输出限制（#5756）的反馈，显示出社区在追求更高自动化水平的同时，对稳健性和安全性的需求也日益增长。

#### 2. 版本发布
**`v0.19.1`** (最新) 和 **`v0.19.0`** 于今日发布。

*   **v0.19.1** (最新): 主要包含一个功能增强：
    *   **增强的 MCP 支持**: CLI 现在支持通过名称匹配 MCP 资源补全，并能自动发现 MCP 服务器。这提升了与外部工具生态的集成体验。 ([PR #5733](https://github.com/QwenLM/qwen-code/pull/5733))

*   **v0.19.0**: 主要是一个发布流程的里程碑，包含 CI/CD 的改进，例如在稳定版本发布后自动发布 VS Code 扩展。

#### 3. 热点 Issue (10条)
1.  **[#5758] 协议/认证类型解耦: 配置兼容性讨论** (开放性，P2): **关键讨论**。该 Issue 指出了 `providerId` 在 CLI、ACP 和 VSCode 中传递方式的差异，提出了一个最小的代码变更方案——通过协议映射来解耦，以支持更灵活的路由。这是未来支持多种后端架构的核心议题。 ([链接](https://github.com/QwenLM/qwen-code/issues/5758))

2.  **[#5760] 使用 LLaMA.cpp Slot 状态保存/恢复替代基于文本的压缩** (开放性，P2): **性能优化**。提出了一个激进的方案，利用 `llama.cpp` 的 KV-Cache Slice 状态保存/恢复功能来实现上下文压缩，从而避免在触发压缩时的重复预填充（re-prefill），能显著提升长对话的性能。 ([链接](https://github.com/QwenLM/qwen-code/issues/5760))

3.  **[#5749] 为自动模式添加针对破坏性 Git 命令的确定性守卫** (开放性，P2): **安全增强**。社区成员明确要求加入代码级别的安全防护，用以阻止 `git reset --hard` 这类命令在无用户确认的自动模式下执行。 ([链接](https://github.com/QwenLM/qwen-code/issues/5749))

4.  **[#5597] 添加 `/model --vision` 命令以设置备用视觉模型** (开放性，P2): **模型切换**。用户希望能在主模型不支持视觉功能时，通过一个简单命令切换到备用的视觉模型，免去在 UI 中手动切换的麻烦。 ([链接](https://github.com/QwenLM/qwen-code/issues/5597))

5.  **[#5626] 提议：通过守护进程(Daemon) + WebUI 架构复活 Chrome 扩展** (开放性，P2): **架构重构**。一项已被采纳的提案，计划废弃之前基于原生消息的复杂方案，转而利用 `serve` 守护进程作为后端，构建一个更现代、更轻量的 Web UI 扩展。 ([链接](https://github.com/QwenLM/qwen-code/issues/5626))

6.  **[#5756] 默认 8K 输出上限 (CAPPED_DEFAULT_MAX_TOKENS) 反复截断大输出，导致失败重试循环** (开放性，P2): **严重体验问题**。一个关键的代码级限制(`CAPPED_DEFAULT_MAX_TOKENS=8000`)在没有设置 `max_tokens` 时生效，导致在生成大型文件（如Wiki）时被截断，从而触发等待和重试的死循环，严重影响自动化任务的可靠性。 ([链接](https://github.com/QwenLM/qwen-code/issues/5756))

7.  **[#5748] 添加 `/config key=value` 斜杠命令以在提示框中设置任意设置** (开放性，P2): **配置便捷性**。一个受欢迎的功能请求，希望无需打开设置文件或UI，直接在命令行中通过斜杠命令修改任何配置项。 ([链接](https://github.com/QwenLM/qwen-code/issues/5748))

8.  **[#5759] 添加 `ui.history.collapsePreviewCount` 以在恢复折叠会话时显示最后N条消息** (开放性，P2): **UI 优化**。在恢复历史会话时，完全折叠所有消息让用户很难迅速定位之前的进度。此建议旨在提供一个预览功能，显示最后几条消息的上下文。 ([链接](https://github.com/QwenLM/qwen-code/issues/5759))

9.  **[#5766] 性能(CI)：将默认PR检查折叠到单个Ubuntu环境中** (开放性，P2): **开发体验优化**。一个旨在优化 CI/CD 管道的提议，通过将多个检查任务合并到一个Ubuntu环境中运行，以缓解 CI runner 压力并加快 PR 审查流程。 ([链接](https://github.com/QwenLM/qwen-code/issues/5766))

10. **[#5761] UI BUG: 模型选择器显示两项选中，状态栏显示错误的套餐信息** (已关闭，P2): **UI 错误**。新发布的版本中，模型选择下拉菜单出现了 UI 状态不一致问题，同时选中了“Standard”和“Coding Plan”两个同类模型，状态栏显示的信息也不正确。该问题已被快速修复。 ([链接](https://github.com/QwenLM/qwen-code/issues/5761))

#### 4. 关键 PR 进展 (10条)
1.  **[#5030] 功能: 恢复中断的对话轮次而不添加合成的 "continue" 消息** (开放性): **重要架构改进**。通过利用持久化的聊天状态，为 SDK 提供一种一流的恢复方式，能无缝处理中断和恢复，不再需要插入虚假的用户消息，提升对话的连贯性和优雅性。 ([链接](https://github.com/QwenLM/qwen-code/pull/5030))

2.  **[#5754] 功能: 为自动模式中的破坏性命令添加确定性守卫** (开放性): **安全功能**。直接实现了 #5749 中提出的需求，通过一个基于正则表达式的预过滤器，在执行任何 LLM 决策之前，直接阻断破坏性的 Git 和 IaC 命令。 ([链接](https://github.com/QwenLM/qwen-code/pull/5754))

3.  **[#5126] 功能(视觉桥接): 为纯文本模型将图像转录为文本** (开放性): **多模态粘合剂**。这是解决智能体不能“看”的问题的一种创新方案。当用户引用图片时，系统会自动将图片发送给一个支持视觉的“代理”模型进行描述，并将文字描述传递给主模型。 ([链接](https://github.com/QwenLM/qwen-code/pull/5126))

4.  **[#5755] 功能(服务端): 为 Web Shell 添加上下文语音听写功能** (开放性): **新交互模式**。通过在 `qwen serve` 的 Web Shell 中添加语音支持，浏览器捕获麦克风数据流，服务端实时转写，为 Web 端用户带来了便捷的语音输入能力。 ([链接](https://github.com/QwenLM/qwen-code/pull/5755))

5.  **[#5764] 修复(CLI): 从会话级别的聊天对象获取 `/context` 的 Token 总数** (开放性): **修复关键 Bug**。修复了多会话模式下，`/context-usage` 报告的是全局总 token 数面而非某个特定会话的 token 数的问题。这对于守护进程模式下的资源监控至关重要。 ([链接](https://github.com/QwenLM/qwen-code/pull/5764))

6.  **[#5561] 功能(MCP): 在设置变更时实时协调 MCP 服务器** (开放性): **运维便捷性**。实现了 MCP 服务器的热重载，用户修改 `settings.json` 后，不再需要重启整个 CLI 进程即可生效。 ([链接](https://github.com/QwenLM/qwen-code/pull/5561))

7.  **[#5767] CI: 将PR检查合并到单个Ubuntu环境中** (开放性): **效率提升**。实现了 #5766 中的提议，减少CI资源消耗，加速开发流程。 ([链接](https://github.com/QwenLM/qwen-code/pull/5767))

8.  **[#5741] 功能(服务端): 添加远程 LSP 状态路由** (已关闭): **服务化**。为守护进程和 ACP 客户端提供结构化的远程 LSP 状态查询接口，使得在服务端模式下管理语言服务器状态成为可能。 ([链接](https://github.com/QwenLM/qwen-code/pull/5741))

9.  **[#5747] 修复(打包): 为镜像安装捆绑音频捕获模块** (开放性): **打包体验优化**。解决在特定镜像源下依赖缺失的问题，通过将音频捕获的二进制文件预先打包，确保功能完整性。 ([链接](https://github.com/QwenLM/qwen-code/pull/5747))

10. **[#5757] 修复(VSCode): 始终在活动栏侧边栏显示聊天视图** (开放性): **兼容性修复**。修正了差异化代码路径，确保 Qwen Code 聊天视图在不同版本 VS Code 中都能稳定显示在左侧活动栏。 ([链接](https://github.com/QwenLM/qwen-code/pull/5757))

#### 5. 特性请求趋势
从今日的 Issue 和 PR 来看，社区最渴望的几个方向是：
*   **配置与模型管理的便捷性**: 用户希望使用 `/config` 和 `/model --vision` 等斜杠命令实现更快捷的配置和模型切换，减少对 UI 或配置文件的依赖。
*   **协议解耦与架构现代化**: 社区正在积极探索和讨论如何将 Qwen Code 从一个单一 CLI 工具演变为一个更强大、更灵活的平台，典型代表是 `providerId` 的解耦（#5758）和通过守护进程重构 Chrome 扩展（#5626）。
*   **视觉与多模态能力**: 用户对智能体的视觉需求很明确（#5597），社区也提出了创新的“视觉桥接”方案（#5126），优先为纯文本模型赋予基础视觉能力。
*   **开发与运维效率**: 优化 CI/CD 流水线（#5766）和实现 MCP 热加载（#5561）等请求，反映了用户对提升开发者和运维者体验的强烈诉求。

#### 6. 开发者痛点
*   **自动化任务的不可靠性**: 默认的 8K 输出上限 (CAPPED_DEFAULT_MAX_TOKENS) 是最大的痛点，它导致了大型生成任务的失败和重试循环，严重破坏了自动化流 (`#5756`)。这突出了一个事实：AI 开发工具的默认参数必须针对最复杂的使用场景进行优化。
*   **CI/CD 不稳定性**: 多次的“Release Failed”记录（#5725, #5686, #5653）表明，尽管有改进措施，但发布流程的稳定性依然是易碎点，这直接影响了内测用户的体验和迭代速度。
*   **多会话管理混乱**: 在 `serve` 守护进程模式下，会话级别的状态管理（如 token 计数）存在 Bug（#5763），这给需要在服务端运行多个独立会话的开发者带来了困惑和调试困难。

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-06-23

Data source: [github.com/Hmbown/CodeWhale](https://github.com/Hmbown/CodeWhale) (formerly DeepSeek-TUI)

---

## Today's Highlights

The project completed its rebrand to **CodeWhale** and shipped v0.8.64, the first release under the new name. A massive wave of development activity landed today — 50 pull requests — focused on stabilising the Fleet/sub‑agent substrate, provider‑scope routing, and TUI resilience. Several regressions were fixed, including prompt text being misinterpreted as a mode switch and loss of configured instructions after `/model` changes. The community is also buzzing about a new destructive‑approval behaviour that many find overly intrusive.

---

## Releases

- **[v0.8.64 (CodeWhale)](https://github.com/Hmbown/CodeWhale/releases/tag/v0.8.64)** — First release under the `codewhale` branding. The legacy `deepseek-tui` npm package is deprecated; users are expected to migrate following `docs/REBRAND.md`. No other release‑note details provided in the data, but the massive PR volume indicates this release contains many of the fixes and features described below.

---

## Hot Issues

1. **[#1978 (CLOSED) – OpenRouter‑compatible `base_url` fixture](https://github.com/Hmbown/CodeWhale/issues/1978)**  
   Validates custom provider URLs as provider‑scoped route overrides, using ZenMux as the concrete test case. Important for users who route through OpenRouter‑compatible gateways.

2. **[#3289 (CLOSED) – Fleet worker fanout freeze regression](https://github.com/Hmbown/CodeWhale/issues/3289)**  
   Ensures that spawning multiple Fleet/sub‑agent workers cannot freeze TUI input/render/cancel. A critical reliability fix for heavy multi‑agent workloads.

3. **[#3004 (CLOSED) – Provider‑scoped API key from commands/secret managers](https://github.com/Hmbown/CodeWhale/issues/3004)**  
   Allows users to fetch API keys from external commands or secret managers instead of storing them in config files or shell history. A security‑focused enhancement.

4. **[#2900 (CLOSED) – DSML tool‑call streaming regression on SiliconFlow](https://github.com/Hmbown/CodeWhale/issues/2900)**  
   Fixes a bug where tool‑call markup streamed as ordinary text on the SiliconFlow/DeepSeek route. Preserved as a regression fixture for future builds.

5. **[#2621 (CLOSED) – Xiaomi MiMo Token Plan provider integration](https://github.com/Hmbown/CodeWhale/issues/2621)**  
   Adds full provider support for Xiaomi’s MiMo Token Plan, including endpoint/auth, quota display, and rate‑limit handling. Shows the project’s push toward multi‑provider parity.

6. **[#3019 (CLOSED) – Codex/Responses route reliability](https://github.com/Hmbown/CodeWhale/issues/3019)**  
   Brings the OpenAI Codex/Responses route up to chat‑completions reliability: retry/backoff parity, correct tool‑result serialisation, and reasoning‑effort mapping.

7. **[#3079 (CLOSED) – Make `web_search` reliable with SearXNG](https://github.com/Hmbown/CodeWhale/issues/3079)**  
   Overhauls the web search stack to use a SearXNG JSON backend with health checks and visible agent status. Addresses a common complaint that “agent says it’s searching but nothing useful happens”.

8. **[#3320 (CLOSED) – Alibaba Bailian API key routing](https://github.com/Hmbown/CodeWhale/issues/3320)**  
   Adds provider‑scoped auth and route descriptors for the Bailian/Bailian‑compatible API. Another example of the widening provider ecosystem.

9. **[#3466 (OPEN) – How to cancel destructive approval?](https://github.com/Hmbown/CodeWhale/issues/3466)**  
   A user reports that v0.8.64 now requires destructive approval on every operation and wants the old no‑confirmation behaviour back. This is the hottest open issue (2 comments, many 👍 expected). The developer pain is clear.

10. **[#3387 (OPEN) – Prompt text misinterpreted as mode switch](https://github.com/Hmbown/CodeWhale/issues/3387)**  
    A critical UX bug where ordinary user input can be parsed as a mode‑change instruction, silently altering the session. A fix is already underway (PR #3455).

---

## Key PR Progress

1. **[#3455 – Fix prompt text misinterpretation](https://github.com/Hmbown/CodeWhale/pull/3455)**  
   Directly addresses issue #3387. The `parse_mode` catch‑all silently turned unrecognised values into `Agent`, allowing mode leaks. PR tightens the logic to reject invalid mode strings.

2. **[#3469 – Fleet profile types + config plumbing](https://github.com/Hmbown/CodeWhale/pull/3469)**  
   Adds the foundational `FleetProfile`, `FleetRole`, `FleetSlot`, `FleetLoadout` types (additive only). This is the first slice of the Fleet substrate (EPIC #3154).

3. **[#3467 – Per‑automation mode/shell/trust/approval settings](https://github.com/Hmbown/CodeWhale/pull/3467)**  
   Makes scheduled automation runs configurable per task rather than hardcoding `mode=agent / allow_shell=false`. Important for DevOps flexibility.

4. **[#3463 – Shrink eager tool surface](https://github.com/Hmbown/CodeWhale/pull/3463)**  
   Drops `task_shell_start`/`task_shell_wait` from the always‑enabled tool set and folds `tool_search_tool` into the standard search tool. Reduces clutter and improves agent focus.

5. **[#3458 – Route foundation: canonical types + resolver](https://github.com/Hmbown/CodeWhale/pull/3458)**  
   The backbone of EPIC #2608: a new `crates/config/src/route/` module providing `ReadyRouteCandidate` and `RouteResolver` types. Unifies route resolution across the codebase.

6. **[#3454 – Throttle sub‑agent mailbox telemetry](https://github.com/Hmbown/CodeWhale/pull/3454)**  
   Fixes UI lag by throttling per‑agent sub‑agent telemetry events. A chatty child agent can no longer flood the bounded event channel.

7. **[#3465 – Preserve configured instructions after `/model`](https://github.com/Hmbown/CodeWhale/pull/3465)**  
   Fixes issue #3457 (Chinese user report). Now the Alt‑C context inspector retains `Configured instructions` when switching models mid‑session.

8. **[#3446 – Route inline reasoning streams](https://github.com/Hmbown/CodeWhale/pull/3446)**  
   Adds `reasoning_stream_style` config to support `separate_field`, `inline_tags`, and `none`. Streams `<think>...</think>` content into Thinking blocks, even when tags split across SSE chunks.

9. **[#3473 – Fact‑drift CI gate for version, providers, tool inventory](https://github.com/Hmbown/CodeWhale/pull/3473)**  
   Extracts shared derivation logic and adds a CI check that fails when `web/lib/facts.generated.ts` is stale. Prevents documentation drift.

10. **[#3462 – Ablate in‑turn reasoning guardrails + encode dispositions](https://github.com/Hmbown/CodeWhale/pull/3462)**  
    Removes the `loop_guard.rs` (390 lines deleted) and moves reasoning judgement into the constitution as dispositions. A net simplification of the engine core.

---

## Feature Request Trends

- **Provider expansion** – Continuous demand for new model providers: Xiaomi MiMo, Alibaba Bailian, Baidu Qianfan (PR #3425), OpenRouter compatibility, and better Together AI validation ( #3382).
- **Fleet / sub‑agent reliability** – The community strongly desires stable concurrent worker execution without TUI freezes ( #3289) and with clear progress tracking ( #3366).
- **Web search as a first‑class tool** – The `web_search` tool is being rebuilt from the ground up with SearXNG, health checks, and visible agent status ( #3079). Users want search to be transparent and reliable.
- **Secret management** – Many requests for externalising API keys ( #3004) to avoid storing secrets in config files or shell history.
- **Reasoning wire‑protocol mapping** – Users want thinking/reasoning effort to be correctly mapped per provider, not silently ignored ( #3024, #3446).

---

## Developer Pain Points

1. **Intrusive destructive‑approval** – Issue #3466 highlights that v0.8.64’s new approval‑every‑time behaviour is frustrating power users who want a seamless workflow. A documentation PR (#3460) hints at persistence, but the UX remains contentious.
2. **MCP duplicate server processes** – Issue #3461 reports that a single `mcp.json` entry spawns two processes, wasting memory and causing death‑by‑shared‑pipe. This is a reliability concern for tool integrations.
3. **Prompt text misinterpretation** – Issue #3387 shows a serious UX bug where ordinary input can accidentally change the mode. The fix is in PR #3455, but until merged users must be cautious.
4. **Loss of configured instructions on model switch** – Issue #3457 (now closed by PR #3465) shows that mid‑session `/model` commands could drop custom instructions, confusing users.
5. **Stale documentation and tool names** – Multiple PRs (#3472, #3471, #3459) address outdated references to retired tools and the old DeepSeek branding, indicating that documentation lag is a recurring complaint.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*