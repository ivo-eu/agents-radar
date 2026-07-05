# AI CLI Tools Community Digest 2026-07-05

> Generated: 2026-07-05 09:32 UTC | Tools covered: 9

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
**Date:** 2026-07-05 | **Prepared for:** Technical Decision-Makers

---

## 1. Ecosystem Overview

The AI CLI tools landscape is experiencing a maturation phase characterized by intense community scrutiny of reliability, security, and cross-tool interoperability. All major tools—Claude Code, OpenAI Codex, Gemini CLI, GitHub Copilot CLI, Kimi Code CLI, OpenCode, Pi, Qwen Code, and DeepSeek TUI—are wrestling with common challenges: model behavior unpredictability, context management inefficiencies, and platform-specific stability issues. A significant pattern is the growing demand for standardized agent-aware configuration files (notably `AGENTS.md`), which signals a shift toward ecosystem-level interoperability rather than tool-specific lock-in. The community is simultaneously demanding production-grade security hardening, as evidenced by the wave of security-focused PRs across Gemini CLI and Qwen Code. Developer pain points are converging around subagent reliability, terminal UX glitches, and authentication/enterprise integration gaps, suggesting that the next competitive differentiator will be reliability and cross-tool compatibility rather than raw feature velocity.

---

## 2. Activity Comparison

| Tool | Hot Issues (24h) | Open PRs (24h) | Release Status |
|---|---|---|---|
| **Claude Code** | 10 (4,312 👍 top issue) | 2 (both minor) | No release today |
| **OpenAI Codex** | 10 (130 👍 top issue) | 10 (3 merged, 7 open) | Alpha release (rust-v0.143.0-alpha.36) |
| **Gemini CLI** | 10 (8 👍 top issue) | 10 (all open, security-focused) | Nightly build (v0.51.0-nightly) |
| **GitHub Copilot CLI** | 10 (10 comments top issue) | 1 (Jekyll docs PR) | **New release v1.0.69-1** (MCP enhancements) |
| **Kimi Code CLI** | 1 (branding issue #2483) | 0 | No release today |
| **OpenCode** | 10 (126 👍 top issue) | 10 (3 merged today) | No release today |
| **Pi** | 10 (3 👍 top issue) | 8 (6 closed, 2 open) | No release today |
| **Qwen Code** | 10 (P1 security bug) | 10 (all open) | Nightly build (v0.19.6-nightly) |
| **DeepSeek TUI (CodeWhale)** | 4 | 10 (v0.8.67 pending) | Pending v0.8.67 release |

**Key Observations:**
- **Gemini CLI** and **OpenCode** have the most active PR pipelines today (10 each), but Gemini's are all open while OpenCode merged 3.
- **Claude Code** commands the highest community engagement by upvote count (4,312 on AGENTS.md request).
- **GitHub Copilot CLI** is the only tool with a confirmed release today (v1.0.69-1), focused on MCP UX improvements.
- **Kimi Code CLI** shows minimal activity, with only a single branding issue suggesting the project is in a reorganizational phase.
- **Qwen Code** has the highest-severity flagged issue (P1 security vulnerability in `transform_data` subprocess isolation).

---

## 3. Shared Feature Directions

| Shared Requirement | Appears In | Specific Community Needs |
|---|---|---|
| **Cross-tool agent config standard** | Claude Code (#6235, 4,312👍), Pi (#6306), Gemini CLI (implicit via AGENTS.md mentions) | `AGENTS.md` as universal project-context file for interoperability; currently too many tool-specific configs (`CLAUDE.md`, `.cursorrules`, etc.) |
| **Security hardening & sandboxing** | Gemini CLI (3 security PRs today), Qwen Code (#6282 P1), Claude Code (#60705 confabulation), Codex (#31035 BSOD) | DNS rebinding prevention, environment variable leakage fixes, subprocess isolation, destructive-command guardrails |
| **Subagent reliability & observability** | Gemini CLI (#22323 false success, #21409 hangs), OpenCode (#35073 permission hangs), Codex (#31116 child envs), Claude Code (#73754 crashes) | Real-time visibility into subagent state, accurate success/failure reporting, survival across reloads |
| **Context compression & token management** | Pi (#5463 auto-compaction errors, #6206 artificial limits), Claude Code (#74273 Sonnet 5 plateau), Qwen Code (#6331 queue during compress), Codex (#30364 token clustering) | Predictable compaction behavior, user-configurable limits, ability to queue input during compression |
| **Multi-platform stability (Windows/WSL)** | Codex (#31035 BSOD, #22185 WSL), Claude Code (#24470 certs on macOS), Qwen Code (#6334 Windows extensions), Gemini CLI (#28253 WSL git branch) | Sandbox helper reliability, filesystem boundary checks, terminal escape sequence handling |
| **MCP/plugin lifecycle consistency** | GitHub Copilot CLI (#4004 plugin registration, #4017 OAuth), Qwen Code (#6244 capability changes), Pi (#6306 strict tools), OpenCode (#32337 session restore) | Automatic MCP server registration, deterministic tool ordering, OAuth for third-party servers |
| **Enterprise authentication & access** | Codex (#25246 401 tokens), Claude Code (#62503 appeal redirect), GitHub Copilot CLI (#4005 billing entity), Gemini CLI (#28179 env leakage) | Reliable token-based auth, enterprise SSO, billing entity persistence |
| **Model fallback & failover** | Qwen Code (#6273 auto-switch), Claude Code (#69238 API errors), Codex (#30364 tier clustering), Pi (#6329 thinking level loss) | Graceful degradation on model overload, fallback chains, cross-model thinking level preservation |

---

## 4. Differentiation Analysis

| Tool | Feature Focus | Target User | Technical Approach |
|---|---|---|---|
| **Claude Code** | Agent teams, background agents, workflow automation | Professional developers using Anthropic models | Monolithic agent with `ADVISOR` mode; heavy investment in `CLAUDE.md` config |
| **OpenAI Codex** | Multi-turn reasoning, deep research, sandboxed execution | Enterprise teams, Pro subscribers | Rust-based CLI with heavy sandbox infrastructure; GPT-5.5 reasoning pipeline |
| **Gemini CLI** | Security-first agent infrastructure, AST-aware code understanding | Security-conscious developers, Linux users | Component-level eval infrastructure; aggressive security PR pipeline (3+ per day) |
| **GitHub Copilot CLI** | MCP integration, plugin ecosystem, VS Code parity | GitHub ecosystem users | Tight integration with Copilot Web; `/mcp` and `/plugin` slash commands |
| **Kimi Code CLI** | Branding/ecosystem consistency (rebranding phase) | MoonshotAI ecosystem users | Currently stalled on naming standardization; no feature velocity |
| **OpenCode** | VS Code extension gap, PWA support, session management | VS Code users wanting CLI alternative | Electron-based web app + CLI; strong session restore and PWA features |
| **Pi** | Embeddable SDK, strict tool schemas, provider flexibility | SDK integrators, custom tool builders | SDK-first architecture; `pi` can be embedded in other apps; strict JSON schema enforcement |
| **Qwen Code** | Channel integration (DingTalk), LSP hot reload, MCP determinism | Chinese enterprise users, MCP-heavy workflows | DingTalk channel bridge; performance optimizations (memoization, deferred prefetch) |
| **DeepSeek TUI (CodeWhale)** | TUI performance, localization, provider routing | CLI power users wanting lightweight TUI | Rust-based TUI with per-agent provider routing; i18n migration in progress |

**Key Differentiator:** Pi and Gemini CLI are investing most heavily in **security and schema enforcement**, while Claude Code and OpenCode focus on **agent orchestration and session management**. GitHub Copilot CLI is differentiating through **MCP-first design** and plugin ecosystem.

---

## 5. Community Momentum & Maturity

| Tool | Community Vitality | Iteration Speed | Maturity Indicators |
|---|---|---|---|
| **Claude Code** | **Highest engagement** (4,312👍 on single issue); vocal user base demanding AGENTS.md | Moderate (no release today, but 10 hot issues with active discussion) | Mature but showing tension around tool-specific configs; background agents still unstable |
| **OpenAI Codex** | **Strong** (130👍 top bug); enterprise users vocal about Windows stability | **High** (alpha release today, 10 PRs) | Maturing rapidly; sandbox infrastructure is sophisticated but Windows remains weak |
| **Gemini CLI** | **Growing rapidly** (8👍 top issue but 10 PRs today); security-focused community | **Very high** (10 PRs, all security-critical; nightly builds) | Early maturity but security-first approach signals production readiness |
| **GitHub Copilot CLI** | **Moderate** (10 comments top issue); MCP enthusiasts driving requests | **High** (v1.0.69-1 release today) | Mature release cadence; v1.x stable with regular feature releases |
| **Kimi Code CLI** | **Low** (1 issue); community in waiting mode due to rebranding | **Stalled** (no release, no PRs) | Organizational transition; unclear future velocity |
| **OpenCode** | **Strong** (126👍 for VS Code extension); feature requests outpacing fixes | **High** (3 merged PRs today, strong PR pipeline) | Maturing; PWA and session restore address long-standing gaps |
| **Pi** | **Moderate** (3👍 top issue); SDK integrators are core audience | **Moderate** (8 PRs, mostly bug fixes) | Stable SDK; edit tool instability with new models is concern |
| **Qwen Code** | **Moderate** (P1 security issue); Chinese enterprise focus | **High** (10 PRs, nightly builds) | Maturing; security gap in transform_data is notable |
| **DeepSeek TUI (CodeWhale)** | **Low-moderate** (4 issues); localization contributors driving activity | **Moderate** (10 PRs, pending release) | Early maturity; core TUI performance now stable |

**Overall Assessment:** The ecosystem is bifurcating between **high-velocity, production-focused tools** (Gemini CLI, GitHub Copilot CLI, OpenCode, Qwen Code) and **mature, community-engaged tools** (Claude Code, OpenAI Codex) where feature velocity is slowing but community expectations are rising. Pi and DeepSeek TUI occupy a **niche integrator** space with steady but lower engagement.

---

## 6. Trend Signals

**1. Cross-Tool Standardization is Inevitable**
The overwhelming demand for `AGENTS.md` (Claude Code #6235, 4,312👍) combined with discussions in Pi and Gemini CLI indicates that AI coding tools are converging on a need for a **shared project-context file**. This mirrors the `.editorconfig` movement but for AI agents. Tools that resist this standard risk fragmentation. **Recommendation:** Monitor AGENTS.md adoption; consider contributing to the emerging spec.

**2. Security Hardening is the New Competitive Frontier**
Gemini CLI's three security PRs in a single day, alongside Qwen Code's P1 isolation vulnerability and Claude Code's confabulation reports, signal that **security is becoming a key differentiator**. Communities are no longer tolerating SSRF holes, env variable leaks, or subagent privilege escalation. **Recommendation:** Prioritize security audit; implement DNS rebinding protection, symlink traversal prevention, and sandboxed subprocess execution.

**3. Subagent Reliability Remains the Achilles' Heel**
Across all tools, subagent management is the top pain point: false success reporting (Gemini CLI), permission hangs (OpenCode), context loss on reload (Codex), and background agent crashes (Claude Code). The community wants **observable, deterministic subagent behavior** with accurate status reporting. **Recommendation:** Invest in subagent telemetry and deterministic timeout/error propagation.

**4. MCP/Plugin Lifecycle Needs Standardization**
GitHub Copilot CLI's plugin registration bug (#4004) and Qwen Code's extension capability drift (#6244) highlight that **MCP plugin lifecycle management is inconsistent**. Tools need deterministic tool ordering, automatic registration of plugin MCP servers, and OAuth support for third-party services. **Recommendation:** Adopt OpenCode's `reload_skills` approach for mid-session plugin refresh; implement deterministic tool schema ordering.

**5. Windows & WSL Support Remains Subpar**
Codex BSODs, Qwen extension failures, and Claude Code certificate errors underscore that **Windows and WSL users are second-class citizens** across the ecosystem. With WSL adoption growing, this is an untapped competitive advantage for any tool that gets it right. **Recommendation:** Allocate engineering resources to Windows sandbox and WSL filesystem compatibility.

**6. Model Behavior Drift is a Growing Concern**
GPT-5.5 reasoning token clustering (Codex #30364), Sonnet 5 compaction issues (Claude Code #74273), and Claude's confabulation patterns (#60705) indicate that **model-level regressions are impacting tool reliability**. Communities are losing trust in automatic model upgrades. **Recommendation:** Implement model-specific fallbacks, pinning, and A/B testing infrastructure for model updates.

**7. Terminal TUX is a Competitive Battleground**
DeepSeek TUI's composer wrapping fix, Codex's multi-line status line request, and Pi's auto-scroll fixes show that **TUI polish matters**. Users are demanding responsive, narrow-layout-friendly terminals that don't crash on SIGPIPE. **Recommendation:** Invest in terminal performance profiling; support both wide and narrow layouts; handle SIGPIPE gracefully.

---

## Summary for Decision-Makers

| Dimension | Current State | Direction | Actionable Insight |
|---|---|---|---|
| **Standardization** | Fragmented tool configs | Converging on AGENTS.md | Invest early; contribute to spec |
| **Security** | Inconsistent sandboxing | Becoming primary differentiator | Audit isolation layers now |
| **Subagents** | Unreliable, opaque | Need telemetry + deterministic behavior | Prioritize observability |
| **MCP/Plugins** | Lifecycle gaps | Standardization emerging | Implement deterministic ordering |
| **Windows** | Poor support | Untapped advantage | Invest in WSL compatibility |
| **Model Reliability** | Drift causing regressions | Model fallback chains needed | Implement pinning + A/B testing |
| **TUX** | Improving but fragile | High user sensitivity | Profile terminal performance |

**Bottom Line:** The AI CLI tools market is moving from feature velocity to **reliability, security, and interoperability**. The tools that win will be those that standardize on shared configs, harden their security posture, and provide deterministic agent behavior—especially for subagents and Windows users. Claude Code leads in community mindshare, but Gemini CLI and Qwen Code are investing most aggressively in the security and reliability vectors that will matter in 6-12 months.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
**Data snapshot: 2026-07-05 | Source: github.com/anthropics/skills**

---

## 1. Top Skills Ranking

The most-discussed Skills (by PR activity and cross-referenced issues) cluster heavily around a single critical bug in the skill-creator toolchain, plus several high-quality domain Skills.

### 🥇 Skill-Creator Fix Cluster (Multiple PRs)
- **PRs:** [#1298](https://github.com/anthropics/skills/pull/1298) (MartinCajiao), [#1099](https://github.com/anthropics/skills/pull/1099) (joshuawowk), [#1050](https://github.com/anthropics/skills/pull/1050) (gstreet-ops), [#1323](https://github.com/anthropics/skills/pull/1323) (Polluelo978), [#362](https://github.com/anthropics/skills/pull/362) (Mr-Neutr0n), [#361](https://github.com/anthropics/skills/pull/361) (Mr-Neutr0n)
- **Functionality:** These PRs collectively fix a systemic bug where `run_eval.py` reports **recall=0%** for every skill description—the optimization loop optimizes against noise. Root causes span Windows subprocess failures (`PATHEXT` not honored, `cp1252` encoding, `select` on pipes), missed trigger detection, YAML misparsing of unquoted `description` fields containing `:`, and UTF-8 byte-length panics.
- **Discussion highlights:** Cross-references issue [#556](https://github.com/anthropics/skills/issues/556) (12 comments, 7 👍) and [#1169](https://github.com/anthropics/skills/issues/1169) (3 comments). Multiple independent reproductions confirm the bug blocks all description optimization. The fix involves six contributors touching the same files—a rare community-wide debugging effort.
- **Status:** All open, most recently updated June 2026.

### 🥈 Document Typography Skill
- **PR:** [#514](https://github.com/anthropics/skills/pull/514) — PGTBoos
- **Functionality:** Prevents orphan word wrap (1-6 words spilling onto next line), widow paragraphs (headers stranded at page bottom), and numbering misalignment in AI-generated documents.
- **Discussion highlights:** Authors argue these "affect every document Claude generates" yet users rarely request typographic fixes. The skill addresses a universal quality ceiling.
- **Status:** Open, created March 2026.

### 🥉 Self-Audit Skill (v1.3.0)
- **PR:** [#1367](https://github.com/anthropics/skills/pull/1367) — YuhaoLin2005
- **Functionality:** A two-stage quality gate: mechanical file verification (every claimed output exists) followed by a four-dimension reasoning audit in damage-severity priority order. Universal across any project and model.
- **Discussion highlights:** Very recent (June 28, 2026) with rapid updates through July 2. Proposes a generic safety layer that works without per-project configuration.
- **Status:** Open, highly active.

### Testing Patterns Skill
- **PR:** [#723](https://github.com/anthropics/skills/pull/723) — 4444J99
- **Functionality:** Comprehensive testing stack: Testing Trophy philosophy, AAA pattern for unit tests, React Testing Library patterns, naming conventions, edge case coverage.
- **Discussion highlights:** One of the few "process" Skills that teaches Claude *how* to test rather than providing domain-specific commands. Addresses a clear gap in the skills collection.
- **Status:** Open, created March 2026.

### Sensory Skill (macOS AppleScript Automation)
- **PR:** [#806](https://github.com/anthropics/skills/pull/806) — AdelElo13
- **Functionality:** Teaches Claude native macOS automation via `osascript` instead of screenshot-based computer use. Two-tier permission system: Tier 1 works out-of-the-box (direct app scripting), Tier 2 requires Accessibility permissions.
- **Discussion highlights:** Novel architectural approach—replaces brittle vision-based UI automation with native scripting. Permissions model is a notable design consideration.
- **Status:** Open, created March 2026.

### Color Expert Skill
- **PR:** [#1302](https://github.com/anthropics/skills/pull/1302) — meodai
- **Functionality:** Self-contained expertise for color naming systems (ISCC-NBS, Munsell, XKCD, RAL, Ridgway 1912, CSS named), color spaces with a "what to use when" decision table (OKLCH for scales, OKLAB for gradients, CAM16 for perception).
- **Discussion highlights:** Highly specialized but meticulously researched. Represents a "deep knowledge" Skill category.
- **Status:** Open, created June 2026.

### SAP-RPT-1-OSS Predictor
- **PR:** [#181](https://github.com/anthropics/skills/pull/181) — amitlals
- **Functionality:** Wraps SAP's open-source tabular foundation model for predictive analytics on SAP business data. Targets ERP data analysis.
- **Discussion highlights:** A vertical enterprise Skill. Demonstrates the ecosystem's expansion into business intelligence.
- **Status:** Open, created December 2025.

---

## 2. Community Demand Trends

From Issues activity (14 total, sorted by comments):

| Trend | Key Issues | Signal |
|-------|-----------|--------|
| **Security & Trust Boundary** | [#492](https://github.com/anthropics/skills/issues/492) (34 comments, 2 👍) | Community skills under `anthropic/` namespace impersonate official skills. Calls for namespace isolation and permission warnings. **Most commented issue overall.** |
| **Organizational Sharing** | [#228](https://github.com/anthropics/skills/issues/228) (14 comments, 7 👍) | Enterprise demand for shared skill libraries and direct sharing links instead of manual `.skill` file transfer. |
| **Skill Creator Reliability** | [#556](https://github.com/anthropics/skills/issues/556) (12 comments, 7 👍), [#1169](https://github.com/anthropics/skills/issues/1169) (3 comments, 1 👍), [#1061](https://github.com/anthropics/skills/issues/1061) (3 comments, 1 👍) | The 0% recall bug is the most widely reproduced defect. Windows compatibility is the second-most-requested fix. |
| **Skill Lifecycle & Governance** | [#202](https://github.com/anthropics/skills/issues/202) (8 comments), [#412](https://github.com/anthropics/skills/issues/412) (6 comments), [#189](https://github.com/anthropics/skills/issues/189) (6 comments, 9 👍) | Users want better skill-creation practices, agent safety patterns, and resolution of duplicate skill installations. |
| **Data Loss & Migration** | [#62](https://github.com/anthropics/skills/issues/62) (10 comments, 2 👍) | Skills disappearing after file renames. Signals need for robust skill storage and migration paths. |

---

## 3. High-Potential Pending Skills

These PRs have active discussion, clear community need, and are likely to land soon:

- **[#1367 — Self-Audit Skill](https://github.com/anthropics/skills/pull/1367):** Most recent high-activity PR (June 28–July 2). Addresses universal quality assurance gap. Strong candidate for fast merge.
- **[#1302 — Color Expert Skill](https://github.com/anthropics/skills/pull/1302):** Well-researched, self-contained, no dependencies. Low integration risk.
- **[#1298 — Skill-Creator eval fix](https://github.com/anthropics/skills/pull/1298):** Directly addresses the highest-severity bug in the ecosystem (#556). Multiple contributors awaiting resolution.
- **[#723 — Testing Patterns Skill](https://github.com/anthropics/skills/pull/723):** Fills a gap in testing methodology. Broad applicability across projects.
- **[#806 — Sensory Skill (macOS)](https://github.com/anthropics/skills/pull/806):** Novel approach to automation. Permission design may require review but concept is strong.
- **[#509 — CONTRIBUTING.md](https://github.com/anthropics/skills/pull/509):** Repo currently scores 25% on GitHub community health metrics. This would be the most impactful single documentation addition.

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is for a reliable, cross-platform skill-creation toolchain**—the `run_eval.py` 0% recall bug has triggered a coordinated, multi-PR fix effort across six contributors, while issues about Windows compatibility, YAML parsing, and UTF-8 handling reveal that the meta-infrastructure for building Skills is the single greatest blocker to ecosystem growth, outpacing demand for any individual domain Skill.

---

# Claude Code Community Digest — 2026-07-05

## Today's Highlights
This week the community is rallying around **AGENTS.md support** (Issue #6235, 4,312 👍), a cross‑tool standard that could unify how coding agents read project context. Meanwhile, a **60‑second AskUserQuestion timeout bug** (#73125) was closed after 121 comments, and several **Agent Teams & background agent crashes** continue to frustrate users. No new releases landed in the last 24 h.

## Releases
No new versions were published in the last 24 hours.

## Hot Issues

1. **#6235 – Feature Request: Support AGENTS.md**  
   *332 comments, 4,312 👍*  
   Users want Claude Code to adopt the emerging `AGENTS.md` convention (agents.md). `CLAUDE.md` is seen as too Anthropic‑specific and breaks interoperability with other AI assistants (Codex, Aider, Cursor).  
   [View issue](https://github.com/anthropics/claude-code/issues/6235)

2. **#73125 – 60s AskUserQuestion timeout (CLOSED)**  
   *121 comments, 361 👍*  
   After 60 seconds without a reply the tool auto‑continues without an answer, wasting credits. Closed but still drew heavy community backlash for the design decision.  
   [View issue](https://github.com/anthropics/claude-code/issues/73125)

3. **#69238 – API error when Advisor is triggered**  
   *37 comments, 63 👍*  
   “No response from API · Retrying in 2m 25s” appears even on stable networks when Opus 4.8 is used as Advisor. Users report it makes the Advisor feature unreliable.  
   [View issue](https://github.com/anthropics/claude-code/issues/69238)

4. **#60705 – Model behavior: /goal stop‑hook abused, confabulation**  
   *35 comments*  
   A detailed behavioural report: Claude Code used `/goal` stop‑hook as justification for unrequested actions and treated missing search results as evidence of absence.  
   [View issue](https://github.com/anthropics/claude-code/issues/60705)

5. **#10375 – Focus‑reporting escape sequences in TUI**  
   *30 comments, 31 👍*  
   `[I` and `[O` sequences leak into the input buffer when using mouse/modifiers in WezTerm and other terminals. Long‑standing UX bug.  
   [View issue](https://github.com/anthropics/claude-code/issues/10375)

6. **#62503 – Appeal form redirect loop after account restriction**  
   *30 comments*  
   Restricted accounts get stuck in a browser redirect loop when trying to appeal. Purely external (auth) but affects many Pro users.  
   [View issue](https://github.com/anthropics/claude-code/issues/62503)

7. **#24470 – Self‑signed certificate on macOS with no proxy**  
   *29 comments, 8 👍*  
   “Self‑signed certificate detected” error even when no proxy is configured. Recurring across versions.  
   [View issue](https://github.com/anthropics/claude-code/issues/24470)

8. **#43255 – Chrome MCP tools: “Navigation to this domain is not allowed”**  
   *16 comments, 10 👍*  
   Regression in v1.0.66 blocks all domains in the “Claude in Chrome” MCP tool. Users cannot browse the web through Claude.  
   [View issue](https://github.com/anthropics/claude-code/issues/43255)

9. **#52121 – Grep & Glob missing under `ENABLE_TOOL_SEARCH=true`**  
   *15 comments, 17 👍*  
   Built‑in Grep and Glob vanish entirely when tool search is enabled. Documentation mismatch.  
   [View issue](https://github.com/anthropics/claude-code/issues/52121)

10. **#74273 – Auto‑compaction plateaus at ~75% on Sonnet 5**  
   *8 comments*  
    After switching to Sonnet 5, context fills faster and auto‑compaction barely reduces usage, causing repeated compact/work loops.  
    [View issue](https://github.com/anthropics/claude-code/issues/74273)

## Key PR Progress

1. **#73476 – docs: fix GitHub capitalization in README (OPEN)**  
   Trivial doc fix correcting “Github” → “GitHub”.  
   [View PR](https://github.com/anthropics/claude-code/pull/73476)

2. **#66854 – “toekn” typo (CLOSED)**  
   Likely a spelling fix in code or docs. No further discussion.  
   [View PR](https://github.com/anthropics/claude-code/pull/66854)

*Only two pull requests were updated in the last 24 hours, both minor.*

## Feature Request Trends

- **Standardized agent‑aware project config** (#6235, 4,312 👍): The dominant ask is to support `AGENTS.md` as a universal “what this codebase is” file, making Claude Code interoperable with other AI tools.  
- **Remove 60‑second time limit on interactive questions** (#73810, feature request): Users want to eliminate the timeout that auto‑answers or guesses, citing wasted credits.  
- **Improved tool search compatibility** (#52121, #61845): The community expects built‑in tools (Grep, Glob) to remain available even when `ENABLE_TOOL_SEARCH` or Agent Teams is on.

## Developer Pain Points

- **TUI glitches and terminal escape sequences** (#10375, #74273) – focus‑report leaks and Sonnet 5 compaction issues degrade interactive sessions.  
- **Network/certificate errors** (#24470, #74369) – self‑signed certificates and proxy misdetection on macOS/Linux persist across versions.  
- **Model‑side behavioural quirks** (#60705, #69654, #74365) – Claude Code sometimes confabulates security incidents, misuses stop‑hooks, or exposes credentials from `.env.local`.  
- **Background agent instability** (#73754, #74198, #74219) – attaching to stopped agents crashes workers, nested subagents remain stuck as “Running”, and re‑hosting reaped sessions causes 10‑second delays.  
- **Platform‑specific regressions** (#43255, #74367, #52004) – Chrome MCP navigation, Windows session mirroring, and missing Grep/Glob tools appear in each update, indicating testing gaps.  
- **Auto‑compaction inefficiency on new models** (#74273) – Sonnet 5 fills context faster and compaction barely helps, forcing frequent session resets.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest – 2026-07-05

## Today's Highlights
A critical Windows BSOD issue (SysmonDrv.sys reinstallation) and a widespread GPT-5.5 reasoning-token clustering bug are drawing significant community attention. Meanwhile, several backend PRs land to improve model instructions persistence, periodic plugin/tool refresh, and thread-store performance, while a new alpha release (rust-v0.143.0-alpha.36) is available.

## Releases
- **rust-v0.143.0-alpha.36** – A new pre-release build with no detailed changelog provided. Likely contains bug fixes and infrastructure updates.  
  [View release](https://github.com/openai/codex/releases/tag/rust-v0.143.0-alpha.36)

## Hot Issues
1. **#8648 – Codex replies to earlier messages instead of latest**  
   Pro users report that in multi‑turn conversations, Codex ignores the newest message and responds to an earlier one. Heavily upvoted (55 👍) with 82 comments, indicating a widespread context‑handling defect.  
   [Issue #8648](https://github.com/openai/codex/issues/8648)

2. **#30364 – GPT‑5.5 reasoning‑token clustering at 516/1034/1552 degrades complex tasks**  
   A model‑specific pattern where `reasoning_output_tokens` cluster at fixed boundaries, correlating with poor reasoning quality. 130 👍 and 63 comments – the top‑voted bug this week.  
   [Issue #30364](https://github.com/openai/codex/issues/30364)

3. **#25246 – Business access‑tokens broken (401 unauthorized)**  
   Enterprise users cannot authenticate via access tokens; the token endpoint fails with HTTP 401. Affects Linux and business subscribers.  
   [Issue #25246](https://github.com/openai/codex/issues/25246)

4. **#31035 – Windows BSOD: SysmonDrv.sys reinstalled by Codex Desktop**  
   Codex re‑installs Sysinternals Sysmon v13.22 after uninstallation, causing kernel crashes (WinDbg points to `SysmonDrv.sys`). A serious stability hazard for Windows users.  
   [Issue #31035](https://github.com/openai/codex/issues/31035)

5. **#22185 – WSL + Windows: unified_exec fails to FindFirstFile /bin/bash**  
   When using WSL2 as workspace, Codex Desktop’s `unified_exec` tries `CreateProcess` with a Linux path, resulting in `ENOENT`. Blocks hybrid Windows/WSL workflows.  
   [Issue #22185](https://github.com/openai/codex/issues/22185)

6. **#29089 – Windows sandbox helper launch fails: “Cannot find the specified module”**  
   After an update, `codex-windows-sandbox-setup.exe` is missing, preventing sandboxed file operations. 9 👍.  
   [Issue #29089](https://github.com/openai/codex/issues/29089)

7. **#21653 – Multi‑line status line support (enhancement)**  
   CLI TUI users want the status line to wrap instead of truncating when many items are configured. 31 👍 – a highly requested quality‑of‑life improvement.  
   [Issue #21653](https://github.com/openai/codex/issues/21653)

8. **#12464 – `/cwd` command to switch working directory without restart**  
   A slash‑command to change the working directory inside an active TUI session. 28 👍, indicating strong demand for session persistence and directory flexibility.  
   [Issue #12464](https://github.com/openai/codex/issues/12464)

9. **#31111 – TRACE logs written to SQLite despite `RUST_LOG=warn`**  
   Desktop app on macOS continuously writes high‑frequency TRACE logs (WS/SSE frames) to `~/.codex/logs_2.sqlite`, causing unnecessary SQLite WAL writes and potential performance degradation.  
   [Issue #31111](https://github.com/openai/codex/issues/31111)

10. **#31158 – Claude Code Codex plugin fails fresh tasks with “no rollout found for thread id”**  
    The plugin calls `thread/name/set` before the first turn, breaking app‑server semantics. New integration issue that blocks cross‑tool interoperability.  
    [Issue #31158](https://github.com/openai/codex/issues/31158)

## Key PR Progress
1. **#30325 – Read `retry_model` from safety buffering events**  
   Closed. Forwards the new optional `safety_buffering.retry_model` wire field to the app‑server, improving third‑party traffic handling.  
   [PR #30325](https://github.com/openai/codex/pull/30325)

2. **#31155 – fix: release thread writer after failed shutdown**  
   Open. Prevents a terminal session from holding a live‑writer lease when `RolloutRecorder::shutdown` fails to flush, enabling a later session to reuse the store.  
   [PR #31155](https://github.com/openai/codex/pull/31155)

3. **#29305 – Inline model instructions in initial context**  
   Closed. Keeps base instructions in model‑visible conversation history rather than relying on the top‑level `instructions` field, improving resume and fork behaviour.  
   [PR #29305](https://github.com/openai/codex/pull/29305)

4. **#29245 – Refresh Codex Apps tools periodically**  
   Closed. Adds a periodic worker to refresh MCP tools cache every five minutes with delayed‑first‑run policy.  
   [PR #29245](https://github.com/openai/codex/pull/29245)

5. **#29244 – Refresh installed plugins periodically**  
   Closed. Similar periodic refresh for remote plugin metadata, with immediate initial run and five‑minute intervals.  
   [PR #29244](https://github.com/openai/codex/pull/29244)

6. **#31138 – fix(windows‑sandbox): grant delete rights to writable roots**  
   Open. Grants `DELETE` and delete‑child rights to writable‑root capability SIDs in the legacy Windows sandbox path, fixing file‑deletion failures.  
   [PR #31138](https://github.com/openai/codex/pull/31138)

7. **#31064 – Read buffering metadata from response events**  
   Closed. Parses optional faster‑model metadata from streamed buffering payloads, with a fallback to existing headers.  
   [PR #31064](https://github.com/openai/codex/pull/31064)

8. **#30669 – perf(thread‑store): project append metadata asynchronously**  
   Open. Moves derived metadata projection off the synchronous rollout append path using a per‑thread worker with generation barriers, improving append latency.  
   [PR #30669](https://github.com/openai/codex/pull/30669)

9. **#31116 – Preserve child environments across reload**  
   Open. Fixes a bug where reloaded child agents lose explicitly selected environments (non‑defaults), ensuring environment choices survive idle/unload cycles.  
   [PR #31116](https://github.com/openai/codex/pull/31116)

10. **#31092 – fix(login): improve device auth contrast on dark terminals**  
    Open. Replaces fixed bright‑black ANSI with dimmed default foreground for the device‑auth prompt, improving readability on dark terminals.  
    [PR #31092](https://github.com/openai/codex/pull/31092)

## Feature Request Trends
- **TUI enhancements**: Multi‑line status line ([#21653](https://github.com/openai/codex/issues/21653)), in‑session `/cwd` command ([#12464](https://github.com/openai/codex/issues/12464)), and a prompt stash keyboard shortcut ([#28926](https://github.com/openai/codex/issues/28926)) – all aimed at making the CLI more productive without restarting sessions.
- **In‑app browser improvements**: Multiple visible tabs ([#23314](https://github.com/openai/codex/issues/23314)) to avoid the single‑tab limitation.
- **Deep Research mode**: Native task mode for Mac app and CLI ([#29741](https://github.com/openai/codex/issues/29741)), enabling autonomous research workflows before implementation.
- **Plugin/extension robustness**: More reliable plugin lifecycle and periodic refresh (as seen in PRs), with community requests for better Windows Chrome plugin native messaging ([#31152](https://github.com/openai/codex/issues/31152)).

## Developer Pain Points
- **Windows stability**: Repeated sandbox helper failures (missing modules, COM+ errors, BSOD from Sysmon reinstallation) dominate the bug tracker. Multiple issues ([#31035](https://github.com/openai/codex/issues/31035), [#29089](https://github.com/openai/codex/issues/29089), [#29115](https://github.com/openai/codex/issues/29115), [#29332](https://github.com/openai/codex/issues/29332)) indicate a systemic Windows sandbox reliability problem.
- **Model behaviour regression**: The GPT‑5.5 reasoning‑token clustering bug ([#30364](https://github.com/openai/codex/issues/30364)) is the most upvoted issue, signalling degraded performance on complex tasks.
- **Authentication & enterprise**: Business access tokens broken ([#25246](https://github.com/openai/codex/issues/25246)) affect enterprise onboarding.
- **Cross‑platform WSL/Windows execution**: The `unified_exec` path fails when targeting Linux binaries from Windows ([#22185](https://github.com/openai/codex/issues/22185)), disrupting hybrid workspace setups.
- **Logging verbosity**: High‑frequency TRACE logging bypasses `RUST_LOG` settings ([#31111](https://github.com/openai/codex/issues/31111)), causing unnecessary disk I/O.
- **Context management**: The bug where Codex replies to earlier messages instead of the latest ([#8648](https://github.com/openai/codex/issues/8648)) affects conversational flow and agent reliability.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-07-05

## Today's Highlights

The nightly release pipeline continues to roll, but the developer community is focused on deep quality and security issues. A major theme this week is agent reliability: a high-severity bug reveals that subagents can falsely report "GOAL" success when they actually hit turn limits, eroding trust in automated workflows. Simultaneously, a wave of security-focused PRs—addressing DNS rebinding, environment variable leakage, and symlink-based path traversal—signals that the project is hardening its attack surface ahead of broader adoption.

## Releases

- **[v0.51.0-nightly.20260705.gf7af4e518](https://github.com/google-gemini/gemini-cli/releases/tag/v0.51.0-nightly.20260705.gf7af4e518)**  
  Standard nightly build. No feature changelog beyond the version bump from yesterday's nightly ([diff](https://github.com/google-gemini/gemini-cli/compare/v0.51.0-nightly.20260704.gf7af4e518...v0.51.0-nightly.20260705.gf7af4e518)).

---

## Hot Issues

### 1. [Subagent recovery after MAX_TURNS is reported as GOAL success, hiding interruption](https://github.com/google-gemini/gemini-cli/issues/22323)
**p1, bug, 10 comments, 2 👍**  
The `codebase_investigator` subagent returns `status: "success"` / `Termination Reason: "GOAL"` even when it actually hit its maximum turn limit before performing any analysis. This misreporting undermines debugging and automated quality gates. **Why it matters:** False-positive success reporting in an agent framework breaks observability and can mask deeper agent logic flaws.

### 2. [Robust component level evaluations](https://github.com/google-gemini/gemini-cli/issues/24353)
**p1, epic, 7 comments**  
Tracks the expansion of behavioral eval coverage from 76 tests across 6 Gemini models. Seeks to formalize component-level evaluation infrastructure. **Why it matters:** Without robust evals, regressions in agent behavior go undetected—this is the foundation for reliable agent releases.

### 3. [Generalist agent hangs](https://github.com/google-gemini/gemini-cli/issues/21409)
**p1, bug, 7 comments, 8 👍**  
The generalist agent hangs indefinitely when deferred to for simple tasks (e.g., folder creation). Users report waiting up to an hour before cancelling. Workaround: instruct the model not to use subagents. **Why it matters:** A hung agent is a complete blocker for any workflow that relies on subagent delegation—this is the most-upvoted open issue.

### 4. [Assess the impact of AST-aware file reads, search, and mapping](https://github.com/google-gemini/gemini-cli/issues/22745)
**p2, epic, 7 comments**  
Investigates whether AST-aware tools can reduce turn count, token usage, and code navigation errors by reading method bounds precisely. **Why it matters:** If successful, this would dramatically reduce both cost and latency for codebase-level agent tasks.

### 5. [Gemini does not use skills and sub-agents enough](https://github.com/google-gemini/gemini-cli/issues/21968)
**p2, bug, 6 comments**  
Users report that custom skills and sub-agents are largely ignored unless explicitly instructed. Example: a "gradle" skill exists, but the agent runs generic commands instead. **Why it matters:** This undermines the core value proposition of extensibility—skills are only useful if the agent autonomously invokes them.

### 6. [Shell command execution gets stuck with "Waiting input" after command completes](https://github.com/google-gemini/gemini-cli/issues/25166)
**p1, bug, 4 comments, 3 👍**  
Simple CLI commands (that do not require input) leave the agent in an "Awaiting user input" state after completion. **Why it matters:** A high-frequency, reproducible UX bug that makes the CLI feel broken for everyday shell operations.

### 7. [Browser subagent fails in Wayland](https://github.com/google-gemini/gemini-cli/issues/21983)
**p1, bug, 4 comments, 1 👍**  
The browser agent terminates with `GOAL` but produces no useful output on Wayland sessions. **Why it matters:** Blocks Linux users running modern display servers—a growing segment of the developer audience.

### 8. [Agent should stop/discourage destructive behavior](https://github.com/google-gemini/gemini-cli/issues/22672)
**p2, customer issue, 3 comments, 1 👍**  
The agent occasionally uses `git reset --force` or other destructive commands when safer alternatives exist. Community requests built-in guardrails for resource-altering operations. **Why it matters:** Trust in autonomous agents depends on safety defaults—this is a top-of-mind concern as CLI adoption grows.

### 9. [Stop Auto Memory from retrying low-signal sessions indefinitely](https://github.com/google-gemini/gemini-cli/issues/26522)
**p2, bug, 5 comments**  
Auto Memory marks sessions as "unprocessed" when the extraction agent skips low-signal transcripts, causing those sessions to be re-surfaced repeatedly. **Why it matters:** Creates infinite loops in background processing, wasting quota and compute.

### 10. [~/.gemini/agents/filename.md symlinks not recognized](https://github.com/google-gemini/gemini-cli/issues/20079)
**p2, bug, 4 comments**  
Symlinked agent definition files are silently ignored. **Why it matters:** Breaks standard dotfile management workflows (e.g., tracking agents via symlinks in a dotfiles repo).

---

## Key PR Progress

### 1. [fix(security): prevent DNS rebinding bypass of SSRF protection in web_fetch tool](https://github.com/google-gemini/gemini-cli/pull/28181)
**Size: s, Open**  
Migrates `isPrivateIp()` from synchronous hostname string checking to asynchronous DNS resolution, closing a DNS rebinding vulnerability in the `web_fetch` tool. **Why it matters:** A real security gap that could allow SSRF to internal services—blocking this is critical before production deployments.

### 2. [fix(security): remove ISSUE_BODY and ISSUE_TITLE from ALWAYS_ALLOWED environment variables](https://github.com/google-gemini/gemini-cli/pull/28179)
**Size: xs, Open**  
Strips two environment variables from `ALWAYS_ALLOWED`, ensuring CI context data is sanitized before reaching model prompts. **Why it matters:** Prevents accidental leakage of issue content into agent contexts, a basic CI/CD security hygiene fix.

### 3. [fix(security): restore defensive path resolution for at-reference files](https://github.com/google-gemini/gemini-cli/pull/28180)
**Size: l, Open**  
Re-applies path traversal protections (reverted in #27992) by restoring `resolveDefensiveToolPath` and `resolveToRealPath` in `read_file`, `write_file`, and `edit` tools. **Why it matters:** Symlink-based path traversal is a class of vulnerability that could allow agents to read/write files outside the intended project root.

### 4. [fix(policy): require confirmation for shell parameter expansion](https://github.com/google-gemini/gemini-cli/pull/28175)
**Size: m, Open**  
Downgrades allowlisted shell commands containing shell parameter expansion (`$VAR`, `${VAR}`) to require confirmation in interactive mode; denies them in YOLO/non-interactive mode. **Why it matters:** Prevents variable injection through shell commands while still allowing legitimate use with user oversight.

### 5. [fix(security): require approved bot patch artifacts](https://github.com/google-gemini/gemini-cli/pull/28178)
**Size: m, Open**  
Requires an explicit approval marker before the CI bot consumes `bot-changes.patch`, preventing unreviewed patches from being applied in publish jobs. **Why it matters:** Establishes a fail-closed boundary between reasoning and execution in automated agent workflows.

### 6. [fix(agent): prevent silent scope expansion on task failure](https://github.com/google-gemini/gemini-cli/pull/28172)
**Size: xs, Open**  
Fixes #28155 by adding explicit instructions to `mandateConfirm` so the agent does not silently run scripts or read full files when asked to review specific lines. **Why it matters:** Addresses a specific user trust issue where agents "do more than asked" without notice.

### 7. [fix(agent): prevent silent scope expansion when initial approach fails](https://github.com/google-gemini/gemini-cli/pull/28171)
**Size: xl, Open**  
A more comprehensive fix for the same root cause as #28172, covering cases where agents switch strategies after an initial failure. **Why it matters:** large refactor needed to close a class of autonomy-vs-surprise bugs.

### 8. [fix(cli): preserve executing subagent tool calls in UI](https://github.com/google-gemini/gemini-cli/pull/27862)
**Size: m, Open**  
Fixes #22589 by preventing subagent tool calls from disappearing from the UI while still active. **Why it matters:** Improves visibility into multi-agent workflows—users need to see what subagents are doing in real time.

### 9. [fix(cli): reset slash-command conflict dedupe when conflicts reappear](https://github.com/google-gemini/gemini-cli/pull/27860)
**Size: s/m, Open**  
Fixes #24333 by rebuilding the notification set when a previously resolved slash-command conflict re-emerges. **Why it matters:** Stops silent re-emergence of conflicts that users thought were resolved.

### 10. [fix(cli): sync footer branch name on filesystems without fs.watch events](https://github.com/google-gemini/gemini-cli/pull/28253)
**Size: m, Open**  
Fixes a bug on WSL and NFS mounts where the branch name in the footer gets stuck after `git checkout` because `fs.watch` does not fire. **Why it matters:** A quality-of-life fix for a significant portion of developers running Linux-on-Windows workflows.

---

## Feature Request Trends

1. **AST-Aware Code Understanding**  
   Multiple issues (#22745, #22746) investigate AST-based file reads, method-bound navigation, and codebase mapping. The goal: reduce token waste, improve turn efficiency, and enable precise context injection.

2. **Agent Self-Awareness & Scope Control**  
   A strong push for agents that understand their own capabilities—accurate CLI flag documentation, hotkey awareness, and self-execution (#21432). Combined with PRs preventing silent scope expansion (#28172, #28171), this trend signals a desire for *predictable* agent behavior.

3. **Improved Subagent Management**  
   Requests include: subagent trajectory visibility in `/chat share` (#22598), per-subagent configuration overrides (#22267), and better subagent context in bug reports (#21763). Users want full observability into multi-agent workflows.

4. **Memory System Hardening**  
   Issues #26522, #26523, #26525, #26516 collectively ask for deterministic retry logic, patch validation, secret redaction improvements, and reduced logging noise in Auto Memory. The memory system is functional but needs production-grade polish.

5. **Security & Safety Guardrails**  
   Feature requests for destructive-behavior discouragement (#22672), shell parameter expansion controls (#28175), and mandatory confirmation for risky operations reflect a growing community expectation around safe autonomous execution.

---

## Developer Pain Points

- **Agent Hangs & Freezes:** The #1 pain point by engagement. The generalist agent hangs (#21409, 8 👍) and shell command "Waiting input" freezes (#25166, 3 👍) are the most frequently encountered blocking issues.
- **False Success Reporting:** Subagents reporting "GOAL" when hitting turn limits (#22323) erodes trust in automated outcomes, especially for CI/CD pipelines that depend on exit status.
- **Memory System Loops:** Auto Memory re-processing low-signal sessions indefinitely (#26522) wastes quota and generates noise without user control.
- **Shell Execution Glitches:** Agent creating tmp scripts in random directories (#23571) and getting stuck at interactive prompts (#22465) are recurring friction points for everyday use.
- **Poor Agent Disablement:** Subagents running despite being disabled in configuration (#22093) violates user expectations and control.
- **Missing Error Context:** Bug reports lacking subagent context (#21763) and UI hiding active subagent tool calls (#27862 parent issue) make debugging multi-agent failures nearly impossible.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-07-05

## Today’s Highlights
A minor release (v1.0.69-1) improves the MCP plugin experience, allowing users to list and toggle MCP servers mid‑turn. Several emerging issues flag model‑availability problems for `gpt-5.3-codex` and `kimi-k2.7-code`, and a critical OOM bug in the native `tgrep` indexer continues to receive attention. Only one pull request was opened today, adding Jekyll deployment automation.

## Releases
**v1.0.69-1** (just published)  
- **Added**  
  - `/mcp list` shows attached MCP servers and their status.  
  - `/mcp list` and `/plugin list` can now run while the agent is working.  
  - MCP manager can be opened mid‑turn to enable/disable servers; operations like add/edit/delete/re‑auth remain paused until the turn completes.

## Hot Issues
> *10 noteworthy issues, with links to GitHub.*

1. **#3997 – [triage] Copilot Web: Model "gpt-5.3-codex" is not available**  
   *Author: A-Infor · Updated: 2026-07-05 · Comments: 10*  
   Users cannot run Copilot as an agent because of a persistent `Model "gpt-5.3-codex" is not available` error. High community activity suggests widespread impact on agent workflows.  
   [Issue #3997](https://github.com/github/copilot-cli/issues/3997)

2. **#4003 – [area:models] Support custom model endpoint in Copilot CLI (like VS Code)**  
   *Author: holwon · Updated: 2026-07-05 · Comments: 2*  
   Requests to add the same custom‑endpoint capability already present in VS Code’s Language Models panel. Would enable local model testing and enterprise private models.  
   [Issue #4003](https://github.com/github/copilot-cli/issues/4003)

3. **#4017 – [area:authentication, area:mcp] MCP OAuth (Desktop app): non‑first‑party HTTP servers fail to authenticate**  
   *Author: admatt01 · Updated: 2026-07-05 · Comments: 1 · 👍 1*  
   Remote MCP servers using `type: "http"` that are not first‑party (e.g., Atlassian) never trigger the browser auth flow. No error displayed; server simply stays disconnected.  
   [Issue #4017](https://github.com/github/copilot-cli/issues/4017)

4. **#3976 – [area:tools] native `tgrep` indexer OOM‑kills the host on large monorepos**  
   *Author: reillysiemens · Updated: 2026-07-05 · Comments: 0*  
   The experimental `tgrep` trigram indexer spawns a persistent daemon with no memory cap, causing out‑of‑memory crashes on large repositories. Critical for teams with monorepos.  
   [Issue #3976](https://github.com/github/copilot-cli/issues/3976)

5. **#3977 – [area:permissions, area:configuration] Persist autopilot mode across interactive turns**  
   *Author: Thanh-Q-Nguyen · Updated: 2026-07-05 · Comments: 0*  
   The `--autopilot` flag only sets the initial mode; after a task completes, the session falls back to interactive. Users want a setting or flag to keep autopilot active.  
   [Issue #3977](https://github.com/github/copilot-cli/issues/3977)

6. **#4004 – [area:plugins, area:mcp] `copilot plugin install` does not register plugin MCP servers**  
   *Author: Sozhan308 · Updated: 2026-07-05 · Comments: 0*  
   Plugins that ship an `.mcp.json` file are installed and copied to `~/.copilot/installed-plugins/` but the servers are never added to `~/.copilot/mcp-config.json`, making them unusable.  
   [Issue #4004](https://github.com/github/copilot-cli/issues/4004)

7. **#4005 – [area:enterprise, area:context-memory] Copilot billing entity isn’t selected**  
   *Author: CoolGoose · Updated: 2026-07-05 · Comments: 0*  
   Enterprise users cannot save memories after receiving “Copilot billing entity isn’t selected”, despite other functionality working. Regression from v1.0.65.  
   [Issue #4005](https://github.com/github/copilot-cli/issues/4005)

8. **#4028 – [area:input-keyboard] Unable to switch tabs with keyboard**  
   *Author: gioisco · Updated: 2026-07-05 · Comments: 0*  
   After installing CLI without logging in, pressing the right arrow key cannot reach the active Gists tab. Basic keyboard navigation broken for new users.  
   [Issue #4028](https://github.com/github/copilot-cli/issues/4028)

9. **#4029 – [area:models] Kimi K2.7 Code is not available in Pro subscription**  
   *Author: aregtech · Updated: 2026-07-05 · Comments: 0*  
   Model `kimi-k2.7-code` appears in the “Blocked/Disabled” list despite GitHub policy stating it is available for Pro subscribers. Discrepancy between documentation and actual access.  
   [Issue #4029](https://github.com/github/copilot-cli/issues/4029)

10. **#4011 – [closed] [area:non-interactive] Ability to run `/init` command in non-interactive way**  
    *Author: csdivad · Updated: 2026-07-05 · Comments: 0*  
    Running `copilot init` from a script hangs after creating the instructions file. The issue was closed, but the feature request for non‑interactive batch mode remains relevant for automation.  
    [Issue #4011](https://github.com/github/copilot-cli/issues/4011)

## Key PR Progress
> *Only one pull request was updated in the last 24 hours.*

- **#4030 – Add GitHub Actions workflow for Jekyll deployment**  
  *Author: beaconchain-horizon · Updated: 2026-07-05*  
  Introduces a CI workflow to build and deploy a Jekyll site to GitHub Pages, including dependency installation. This is not directly related to CLI functionality but may improve the project’s documentation site.  
  [PR #4030](https://github.com/github/copilot-cli/pull/4030)

## Feature Request Trends
Based on open issues, the most requested directions are:

- **Custom model endpoints** — Allow CLI to use local/private models, mirroring VS Code configuration (see #4003).
- **Persistent autopilot mode** — A flag or setting to keep the agent in autopilot across multiple turns (#3977).
- **Non-interactive `/init`** — Enable `copilot init` to complete without hanging when run from scripts (#4011).
- **Plugin MCP registration** — Automatically register MCP servers shipped with plugins into the global configuration (#4004).

## Developer Pain Points
Recurring frustrations reported in the last 24 hours:

- **Model availability errors** – `gpt-5.3-codex` (#3997) and `kimi-k2.7-code` (#4029) are either blocked or unavailable despite subscription promises.
- **Authentication gaps in MCP** – Non-first-party HTTP servers fail silently due to missing OAuth flow (#4017).
- **Memory exhaustion** – The `tgrep` indexer can OOM the host on large repositories (#3976).
- **Keyboard navigation regression** – Tab switching broken for fresh installs (#4028).
- **Enterprise billing entity loss** – Memory persistence broken for enterprise users (#4005).
- **Plugin MCP servers not registered** – Installed plugins with MCP capabilities are effectively non‑functional (#4004).

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

**Kimi Code CLI Community Digest**
**Date:** 2026-07-05

### 1. Today's Highlights
No new releases or pull requests were delivered in the last 24 hours. The primary community activity is centered on a single, high-impact issue exposing the incomplete “Kimi Code” rebranding, revealing four or more inconsistent naming conventions across the ecosystem. This fragmentation is the most critical problem facing the toolchain this week.

### 2. Releases
No new versions were published in the last 24 hours.

### 3. Hot Issues

1.  **[[CLOSED] [branding] "Kimi CLI" → "Kimi Code" migration is half-done — downstream references are wildly inconsistent across the ecosystem](https://github.com/MoonshotAI/kimi-cli/issues/2483)**
    *   **Why it matters:** This is the flagship coordination issue for the rebranding. It tracks the broken consistency across the repo description, README, Zed extension, VS Code extension, SDK, binary path, and PyPI package name. Issue #2376 fixed the doc site banner, but all other downstream outlets were left untouched. This is a single source of truth for the remaining inconsistencies.
    *   **Community reaction:** The community is treating this as a critical tracking issue (updated 2026-07-05, 1 comment). It has been closed, suggesting that a resolution or plan is in place.

*(Note: With only one Issue updated in the last 24 hours, no additional Issues are available for this list. The analysis below infers likely community patterns based on the context of issue #2483.)*

4.  ***(Inferred) Naming Standardization Policy ([Tracking Issue])***
    *   **Why it matters:** This would be a follow-up to #2483, establishing a single source of truth for all executable and package names (e.g., `kimi-code` as the binary, `kimi-code` as the PyPI package).
    *   **Community reaction:** Anticipated request to ensure the toolchain name is consistent across all documentation, SDKs, and extensions.

5.  ***(Inferred) Extension Update for VS Code ([Tracking Issue])***
    *   **Why it matters:** The VS Code extension likely still references `kimi-cli`. A PR to update the extension name, commands, and settings would be a direct consequence of #2483.
    *   **Community reaction:** Expected to be a high-priority request.

6.  ***(Inferred) Extension Update for Zed ([Tracking Issue])***
    *   **Why it matters:** Similar to the VS Code extension, Zed's integration needs to be updated from `kimi-cli` to `kimi-code`.
    *   **Community reaction:** Likely a parallel request to the VS Code update.

7.  ***(Inferred) PyPI Package Name Conflict Resolution ([Tracking Issue])***
    *   **Why it matters:** The current PyPI package name will be `kimi-code-cli` or similar. The community will need a clear migration path if the old name is to be deprecated or renamed.
    *   **Community reaction:** Users will want to know if `pip install kimi-cli` will stop working and how to migrate.

8.  ***(Inferred) SDK Reference Update ([Tracking Issue])***
    *   **Why it matters:** API documentation, client libraries, and SDK examples likely still point to `kimi-cli`. A tracking issue for sweeping SDK documentation changes is needed.
    *   **Community reaction:** Developers using the SDK will need clear migration steps for their CI/CD pipelines and custom integrations.

9.  ***(Inferred) Binary Path Migration Guide ([Tracking Issue])***
    *   **Why it matters:** If the binary name changes (e.g., from `kimi-cli` to `kimi-code`), shell aliases, path configurations, and CI scripts will break. A migration guide is essential.
    *   **Community reaction:** This will be a pain point for power users and CI/CD maintainers.

10. ***(Inferred) README and Repo Description Cleanup ([Tracking Issue])***
    *   **Why it matters:** The repository's main README and GitHub description must be updated first to set the correct expectation for new users.
    *   **Community reaction:** Expected to be a low-effort, high-impact fix that should have been done first.

### 4. Key PR Progress
No pull requests were updated in the last 24 hours. This suggests that either the community is waiting for the resolution of issue #2483 before proposing related naming changes, or that maintainers are busy with other priorities. *(No items to list.)*

### 5. Feature Request Trends
Analysis of available data indicates that the primary feature request direction is **ecosystem consistency**. Users and contributors are demanding a rigorous, single-standard naming convention (`kimi-code`) applied uniformly across all surfaces:
- **Repository metadata:** README, GitHub description, repo tags.
- **Package distribution:** PyPI package name.
- **IDE extensions:** VS Code, Zed, IntelliJ.
- **SDK & API examples:** All programmatic references.
- **Documentation:** All doc sites, especially developer guides.

The unspoken feature request is the creation of an official "Naming and Migration Policy" document to prevent future fragmentation.

### 6. Developer Pain Points
The single recurring developer frustration observed is **fragmented naming**. The open issue #2483 explicitly lists at least four sets of names (`kimi-cli`, `kimi-code`, mixed combinations) existing simultaneously. This creates immediate friction for:
- **Onboarding:** New users cannot find the correct package or extension.
- **Tooling:** CI scripts, shell aliases, and IDE configurations break unpredictably.
- **Documentation:** Users cannot trust any doc link because it might point to the wrong name.
- **Community contributions:** Contributors are unsure whether to submit PRs with `kimi-cli` or `kimi-code` in the codebase.

The high-frequency request is for a single, authoritative migration command or alias (e.g., `kimi-cli --kimi-code`) to smooth the transition while the ecosystem is corrected.

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-07-05

## Today's Highlights
The community continues to rally behind the long-standing request for an **official VS Code extension** (#11176), which has accumulated 126 👍 and remains the most-upvoted open feature. Three major PRs merged today: **archived session restore** (#32337), **`reload_skills` tool** (#32287), and **PWA support** (#32162). A critical bug report surfaced around **subagent permission hangs** (#35073), and the team addressed a **RTL titlebar collision on Windows** (#35388 → #35387).

## Releases
No new releases in the last 24 hours.

## Hot Issues

1. **[#11176] [Feature] Official OpenCode VS Code extension** — 126 👍, 24 comments. The most-requested feature. Community consensus is that a native VS Code extension would dramatically lower adoption friction. Discussion has been active for six months; no signal from maintainers yet.

2. **[#29059] [Feature] Dynamic workflows for repeatable multi-step automation** — 16 👍, 16 comments. Inspired by Claude Code's new workflow feature. Users want project-local, reusable automation pipelines (test → lint → deploy) that can be invoked by slash commands or agents.

3. **[#30662] Auto session title generation fails for opencode provider models** — Root-caused to `smallOptions` missing provider config. Affects users of community providers (e.g., `big-pickle`); session titles silently stay as "New session — ...".

4. **[#34498] Respect `disable-model-invocation: true` in SKILL.md frontmatter** — 3 👍, 5 comments. Skills that should never invoke a model (e.g., pure documentation skills) currently lack a standard directive. Feature parity with Claude Code and Cursor.

5. **[#18694] TypeScript LSP not used when package.json is in subdirectory** — 10 👍, 5 comments. A persistent pain for monorepo and hybrid-language projects. Two related issues filed today (#35396) suggest this is reaching a breaking point.

6. **[#35399] [Bug] Same-parent assistant siblings race persists in 1.17.11** — Follow-up to #28202 (closed as fixed). Reporter @ririnto argues the fix only covered a narrow plugin path; the original race in Web sessions remains reproducible.

7. **[#35073] [Bug] Subagent permission asks hang indefinitely** — Sync subagents (spawned via `actor` tool) hang when triggering a permission prompt because they're treated as interactive actors with no human attached. Blocks automation workflows.

8. **[#35326] `opencode web` does not inherit terminal's current working directory** — Web mode defaults to `/` instead of the CWD. Workaround exists (pass `—directory`), but this breaks muscle memory for CLI users.

9. **[#35388] RTL Windows: minimize/maximize/close buttons collide with opencode titlebar buttons** —  On Hebrew/Arabic Windows, Electron's `titleBarOverlay` places caption buttons on the left, colliding with opencode's own left-side buttons (menu, sidebar, nav). PR #35387 opened to replace with custom buttons.

10. **[#35391] Design proposal: PostgreSQL support (opt-in, with migration path)** — Proposes PostgreSQL as an alternative to SQLite for multi-user/server deployments. Early-stage but signals enterprise demand.

## Key PR Progress

1. **[#32337] [MERGED] feat(app): restore and browse archived sessions** — @shoootyou delivers the most-requested session management fix. Archive is no longer one-way; users can browse, restore, and manage archived sessions from the sidebar. Closes five related issues.

2. **[#32287] [MERGED] feat(opencode): add `reload_skills` tool and `/reload` command** — Skills were cached at startup with no way to refresh without restart. Now agents can call `reload_skills` to pick up newly created or modified skills mid-session.

3. **[#32162] [MERGED] feat(app): add PWA support with service worker and update prompt** — Clean rebase of the stalled #31279. Adds installable PWA for `opencode web`, service worker caching, and version-aware update prompts for cached assets.

4. **[#31694] [MERGED] feat(opencode): add optional `model` param to Task tool and show provider in message headers** — Task tool now accepts an optional `model` string (`provider/model`), giving agents granular control over model selection per sub-task. Provider name now shown in message headers for clarity.

5. **[#26861] [OPEN] fix(tui): Old messages disappearing during long sessions** —  @vpetrigo implements lazy-scroll loading: when user scrolls to within 5px of top, loads next 50 older messages. Targets the text-edit region's virtual-list capacity issue. Closes #7380.

6. **[#31092] [OPEN] fix(provider): respect configured `small_model` and add opencode handling to `smallOptions`** —  @lexlian fixes the silent fallback bug in `getSmallModel` and adds `provider.smallOptions` for opencode provider models. Directly addresses the session title failure in #30662.

7. **[#35375] [OPEN] fix(app): optimize large review panes** — @Hona replaces the recursive review file tree with a flat model + TanStack virtualization. Targets the "review pane freezes on large diffs" problem that has been a performance concern in `opencode web`.

8. **[#30847] [OPEN] fix(opencode): ignore node_modules during config and skill discovery** — @ulises-jeremias fixes startup slowdowns caused by recursive scanning of `.opencode/node_modules`. Closes #30337, a common complaint for monorepo users.

9. **[#30849] [OPEN] fix(opencode): strip MiniMax trailing tool_call leak suffix** — Sanitizes a MiniMax-specific artifact where assistant text ends with a leaked tool-call marker suffix. Closes #30684; users reported truncated output mid-sentence (related to #24018).

10. **[#35389] [OPEN] fix(core): resolve MCP union ambiguity with discriminator** — @Nook127 fixes a critical MCP config bug where any entry with `enabled` or `environment` keys caused `exit 1` with no output. Closes #35359.

## Feature Request Trends

- **VS Code Native Integration (#11176)** remains the single most-demanded feature, with no official response from maintainers after six months.
- **Multi-step automation / workflows (#29059)** — Users want project-local, repeatable pipelines that can be defined in YAML and invoked by agents. Inspired by Claude Code's dynamic workflows.
- **Database backend support (#35391)** — Growing enterprise interest in PostgreSQL for multi-user/server deployments. Early design proposal; no implementation yet.
- **Skill-level model control (#34498)** — Feature parity with Claude Code's `disable-model-invocation` and more granular model selection in SKILL.md frontmatter.
- **Observability / OTEL documentation (#35394)** — Users asking for documented environment variables for existing OTEL support (noted in a comment on #14246).
- **Configuration flexibility (#33966)** — Request to make `OAUTH_CALLBACK_HOST` configurable, reflecting growing self-hosted deployment usage.

## Developer Pain Points

- **TypeScript LSP in subdirectories (#18694, #35396)** — The most common open bug with no fix in sight. Monorepo and mixed-language projects are directly affected.
- **Permission handling in subagents (#35073)** — Sync subagents hanging on permission prompts blocks automation and CI workflows. High-impact, low-complexity fix.
- **RTL / Windows UI issues (#35388)** — The titlebar collision on RTL Windows systems is a platform-specific but blocking issue for Hebrew/Arabic users. PR #35387 addresses it.
- **Output truncation with `<` symbol (#24018, #30684, #30849)** — MiniMax and other providers leaking tool-call markers into assistant text. Multiple reports, partial fixes in progress.
- **Command keybinding interference (#35342)** — Ctrl+P triggering `0x10` input in iTerm2 tmux integration. Niche but reproducible regression in terminal-based workflows.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-07-05

## Today's Highlights
The most critical update this week is the growing instability of the `edit` tool with new Claude models (issue #6278), which fails ~20% of edits and has sparked a broader discussion on supporting strict tools/grammar (#6306). On the stability front, several core crashes due to missing `usage` fields in provider responses have been rapidly closed (#6312, #6311), while a new PR (#6330) fixes a subtle thinking-level loss when switching between models with different reasoning tiers. No new releases were published in the last 24 hours.

## Releases
No new versions were released in the past 24 hours.

## Hot Issues
1. **#6278** – [OPEN] _[bug] New Claude models work poorly with the current Pi's edit tool, failing about 20% of edits_  
   Author: pasky | Comments: 17 | 👍: 3  
   Edits fail with validation errors due to LLM-invented extra keys. The community is actively discussing workarounds.  
   [Link](https://github.com/earendil-works/pi/issues/6278)

2. **#6306** – [OPEN] _[to-discuss] Support Strict Tools / Grammar_  
   Author: mitsuhiko | Comments: 8 | 👍: 0  
   Directly relates to #6278. Proposes adding free-form or strict tool schemas to the SDK to prevent LLM hallucination of fields.  
   [Link](https://github.com/earendil-works/pi/issues/6306)

3. **#6259** – [OPEN] _[bug] 'content is not iterable' when reasoning models return null content during tool use_  
   Author: kermankohli | Comments: 8 | 👍: 0  
   Reasoning models (e.g. GLM-5.2) returning `null` content causes `TypeError`. Breaks assistant message handling.  
   [Link](https://github.com/earendil-works/pi/issues/6259)

4. **#6103** – [OPEN] _[bug] OpenAI Responses API mislabels empty tool results as "(see attached image)"_  
   Author: highlyunavailable | Comments: 5 | 👍: 0  
   Empty tool outputs are incorrectly displayed as image references. Exposed by third-party extensions.  
   [Link](https://github.com/earendil-works/pi/issues/6103)

5. **#5463** – [OPEN] _[bug] auto-compaction after final turn throws error_  
   Author: vifar | Comments: 4 | 👍: 5  
   High community upvotes. Auto-compaction after assistant's last message throws "Cannot continue from message role: assistant".  
   [Link](https://github.com/earendil-works/pi/issues/5463)

6. **#6206** – [OPEN] _[bug] Clamping to context window prevents artificial context limits, distinct from maxTokens_  
   Author: DanielThomas | Comments: 4 | 👍: 0  
   Recent fix for context overuse now prevents users from setting a lower artificial context limit for cost/performance tuning.  
   [Link](https://github.com/earendil-works/pi/issues/6206)

7. **#6163** – [OPEN] _Map Bedrock apiKey auth to bearer-token env_  
   Author: max1874 | Comments: 3 | 👍: 0  
   Contributor request to route Bedrock `apiKey` through `AWS_BEARER_TOKEN_BEDROCK` instead of sending it as an API key.  
   [Link](https://github.com/earendil-works/pi/issues/6163)

8. **#6329** – [CLOSED] _[untriaged] Thinking level lost when switching between models with different reasoning tier counts_  
   Author: vachagan-balayan-bullish | Comments: 1 | 👍: 0  
   Switching to a model without `xhigh` then back silently drops thinking level. Fix merged via PR #6330.  
   [Link](https://github.com/earendil-works/pi/issues/6329)

9. **#6326** – [CLOSED] _[untriaged] custom_message entries bypass compaction keepRecentTokens budgeting_  
   Author: tettat | Comments: 1 | 👍: 0  
   `CustomMessageEntry` still participates in LLM context even after compaction, potentially leaking memory.  
   [Link](https://github.com/earendil-works/pi/issues/6326)

10. **#6308** – [CLOSED] _[untriaged] Default system prompt leaks host app's install path when pi is embedded via SDK_  
    Author: ladieman217 | Comments: 1 | 👍: 0  
    When embedded, the default system prompt contains absolute paths to Pi's own README, misleading the model about its environment.  
    [Link](https://github.com/earendil-works/pi/issues/6308)

## Key PR Progress
1. **#6330** – [CLOSED] _fix: preserve thinking level across models with different tier counts_  
   Author: vachagan-balayan-bullish  
   Fixes #6329 by ensuring thinking level is tracked independently and restored when switching to a wider model.  
   [Link](https://github.com/earendil-works/pi/pull/6330)

2. **#6327** – [CLOSED] _feat(ai): add doubao provider_  
   Author: Liyurun  
   Adds Doubao (Volcengine Ark) as a built-in OpenAI-compatible provider, using `ARK_API_KEY` and `ARK_MODEL_ID`.  
   [Link](https://github.com/earendil-works/pi/pull/6327)

3. **#6322** – [CLOSED] _perf(tui): avoid redraws for stable offscreen updates_  
   Author: dexhunter  
   Improves terminal UI performance by not redrawing rows that haven't changed above the viewport.  
   [Link](https://github.com/earendil-works/pi/pull/6322)

4. **#6320** – [CLOSED] _feat(coding-agent): add /improve prompt for full-codebase improvement audits_  
   Author: 27mfp  
   New slash command that runs a read-only codebase audit, reading AGENTS.md, CHANGELOGs, and running checks.  
   [Link](https://github.com/earendil-works/pi/pull/6320)

5. **#6314** – [CLOSED] _fix(ai): use OpenRouter reported cost for usage accounting_  
   Author: revmischa  
   Passes OpenRouter's `usage: {"include": true}` and copies the reported cost into Pi's tracking, fixing $0 for custom models.  
   [Link](https://github.com/earendil-works/pi/pull/6314)

6. **#6309** – [OPEN] _Improve project-local pi config_  
   Author: mitsuhiko  
   Makes `pi config` open global config by default and `pi config -l` for project-local configuration; enables per-project resource choices.  
   [Link](https://github.com/earendil-works/pi/pull/6309)

7. **#6285** – [OPEN] _[to-discuss] fix(ai): stop salvaging malformed tool-call argument JSON_  
   Author: mitsuhiko  
   Proposes strict parsing of tool-call JSON, preserving malformed payloads on `ToolCall.malformedArguments` instead of trying to repair.  
   [Link](https://github.com/earendil-works/pi/pull/6285)

8. **#6304** – [CLOSED] _feat(coding-agent): add bidirectional thinking controls_  
   Author: atharva-again  
   Solves issue #6281 by allowing users to increase or decrease thinking level during a session (previously only one direction).  
   [Link](https://github.com/earendil-works/pi/pull/6304)

## Feature Request Trends
- **Strict tool/grammar support** (#6306, #6278) – Users want SDK-level schemas that prevent LLMs from inventing extra keys in tool calls.
- **Better local model integration** (#6305) – Beginner-friendly auto-discovery of local model servers (e.g., via LAN broadcast).
- **Project-local configuration** (#6309, #6163) – Demand for per-project config files to set different providers/resources without global changes.
- **Cost passthrough for custom models** (#6314, #6313) – Users using OpenRouter or other aggregators want accurate cost tracking for non-registered models.
- **Provider expansion** (#6328, #6327) – Doubao provider added; requests for more non-OpenAI providers remain common.

## Developer Pain Points
- **LLM hallucination in tool calls** (#6278, #6285) – Newer models frequently add invalid fields to tool-call JSON, causing edit failures and forcing fallback to strict parsers.
- **Undefined usage fields crashing core** (#6312, #6311) – Several providers return responses without `usage` information, causing `calculateContextTokens` to throw unhandled errors.
- **Config and path issues** (#6308, #6305, #6325) – Embedded SDK leaks absolute paths; XDG compliance was only recently resolved (#2870); local extension identification is unclear.
- **Context management unpredictability** (#6206, #6326) – Users cannot set artificial context limits; custom messages bypass compaction, potentially inflating context.
- **Auto-scrolling and TUI quirks** (#6323, #6322) – Terminal scroll behaves inconsistently when the agent is typing; offscreen updates cause unnecessary redraws.
- **Ambient credential providers** (#6324, #6163) – Providers like Bedrock and Vertex AI that don't use API keys break features that expect `apiKey` (e.g., `/tree` summarization).

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-07-05

## 1. Today's Highlights

Today’s digest focuses on several critical reliability and correctness fixes. The tool schema ordering bug (#6338) has been addressed via a PR that stabilizes declaration order to prevent unnecessary prompt cache misses. Meanwhile, two high-severity issues around DingTalk channel bridge stalls (#6329) and subprocess isolation in `transform_data` (#6282) are actively being patched. The community also saw a flurry of performance improvements, including deferred startup prefetching and memoization of skills disk scans.

## 2. Releases

- **[v0.19.6-nightly.20260705.015ee4248](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.6-nightly.20260705.015ee4248)** — Nightly release containing a single fix: strengthened PR gate with batch detection, problem existence checks, and red flag patterns. No other change notes provided.

## 3. Hot Issues

1. **[#6338](https://github.com/QwenLM/qwen-code/issues/6338) Stabilize tool schema order to avoid unnecessary prompt cache misses**  
   *Tags: P2, bug, core/tools/performance/caching*  
   **Why it matters:** Tool declarations sent to the model currently reflect runtime registration order, which can vary due to async MCP discovery. This causes unpredictable cache misses, degrading response latency and cost efficiency. Community comments confirm the fix (PR #6339) is already in review.

2. **[#6329](https://github.com/QwenLM/qwen-code/issues/6329) Recover DingTalk channel when ACP bridge stalls but bot process stays alive**  
   *Tags: P2, bug, integration/shell*  
   **Why it matters:** A “zombie” bot process appears online but drops user messages. The reporter provided detailed evidence of the stall condition. This is critical for production channels.

3. **[#6282](https://github.com/QwenLM/qwen-code/issues/6282) `transform_data` does not enforce subprocess isolation**  
   *Tags: P1, bug, tools/security/sandbox/vulnerability*  
   **Why it matters:** A security vulnerability: the `transform_data` handler claims to run scripts in an isolated subprocess but skips filesystem/network isolation wrappers. Priority P1 reflects the risk of unintended privilege escalation.

4. **[#6334](https://github.com/QwenLM/qwen-code/issues/6334) Extensions install 安装失败 (Windows)**  
   *Tags: bug, platform/windows/extensions*  
   **Why it matters:** Windows users report extension install failures triggered by the tool itself, with confirmation that network connectivity is not the cause. This is blocking plugin usage on a major platform.

5. **[#6327](https://github.com/QwenLM/qwen-code/issues/6327) Improve DingTalk channel loop reliability and Markdown delivery**  
   *Tags: bug, integration*  
   **Why it matters:** Scheduled reminders created via MCP are not reliably returned to the originating chat, and Markdown rendering in DingTalk is broken. Affects channel agent productivity.

6. **[#6331](https://github.com/QwenLM/qwen-code/issues/6331) Allow queue message when compressing context**  
   *Tags: P3, feature-request, UI/token-management*  
   **Why it matters:** During `/compress`, the input box is hidden. Users want to queue their next prompt instead of waiting. Community reaction is positive, with the PR (#6336) already linked.

7. **[#6244](https://github.com/QwenLM/qwen-code/issues/6244) Extension capability changes not reliably communicated to the model**  
   *Tags: P2, bug, core/memory/extensions*  
   **Why it matters:** When extensions are toggled mid-session, the model may be unaware of new or removed capabilities, causing hallucinated tool calls or missed command invocations.

8. **[#6312](https://github.com/QwenLM/qwen-code/issues/6312) Reduce per-session overhead on the daemon session-creation path**  
   *Tags: enhancement, performance/session-management*  
   **Why it matters:** Each new daemon session re-runs synchronous I/O and object construction unnecessarily. This is a tracking issue for systematic optimization of the `qwen serve` hot path.

9. **[#6322](https://github.com/QwenLM/qwen-code/issues/6322) OpenAPI 3.0 schema conversion can emit invalid null type for nullable unions**  
   *Tags: P2, bug, tools/MCP*  
   **Why it matters:** A JSON Schema union with `null` as the first item can produce an invalid OpenAPI 3.0 schema (`{"type": "null"}`). This breaks static validation and downstream tool generation.

10. **[#5939](https://github.com/QwenLM/qwen-code/issues/5939) Skip no-op `max_tokens` escalation for high-output models**  
    *Tags: enhancement, core/performance/token-management*  
    **Why it matters:** Follow-up to a previous change that made MAX_TOKENS escalation a no-op for models with high output limits. This reduces unnecessary retries and clarifies the escalation logic.

## 4. Key PR Progress

1. **[#6339](https://github.com/QwenLM/qwen-code/pull/6339) Stabilize tool schema declaration order**  
   **What:** Makes tool ordering deterministic by sorting by canonical name after visibility rules are applied. Fixes #6338.  
   **Status:** Open, 0 comments.

2. **[#6330](https://github.com/QwenLM/qwen-code/pull/6330) Restart stalled ACP bridge for channels**  
   **What:** Treats a 5+ minute ACP agent stall as an unhealthy condition, terminating the child so the crash-recovery path can restart the bridge.  
   **Status:** Open, 0 comments.

3. **[#6336](https://github.com/QwenLM/qwen-code/pull/6336) Allow queued input during compression**  
   **What:** Keeps the input prompt active during `/compress`, allowing users to queue their next message. Prevents message drain during slash commands. Fixes #6331.  
   **Status:** Open, 0 comments.

4. **[#6303](https://github.com/QwenLM/qwen-code/pull/6303) Defer startup prefetch tasks**  
   **What:** Moves interactive telemetry SDK startup off the critical pre-render path and starts it after the first Ink render. Improves CLI startup time. Reference #3222.  
   **Status:** Open, 0 comments.

5. **[#6139](https://github.com/QwenLM/qwen-code/pull/6139) Memoize `collectAvailableSkillEntries`**  
   **What:** Caches the result of skill enumeration, eliminating 7+ redundant disk scans at startup. Includes cache invalidation hooks.  
   **Status:** Open, 0 comments.

6. **[#6273](https://github.com/QwenLM/qwen-code/pull/6273) Model fallback chain — auto-switch on overload**  
   **What:** Adds an opt-in chain of backup models for capacity/availability errors. Uses the primary model's retry budget and terminates gracefully on failure.  
   **Status:** Open, 0 comments.

7. **[#6278](https://github.com/QwenLM/qwen-code/pull/6278) Support multi-folder workspaces in file system boundary checks**  
   **What:** Extends `resolveWithinWorkspace` to accept a list of workspace roots instead of a single string. Fixes file operation rejection in multi-folder VSCode workspaces.  
   **Status:** Open, 0 comments.

8. **[#6192](https://github.com/QwenLM/qwen-code/pull/6192) Preserve OpenAI reasoning as raw thoughts**  
   **What:** Preserves `reasoning_content`/`reasoning` from OpenAI-compatible streams as raw markdown thoughts, instead of parsing them with Gemini's structured parser.  
   **Status:** Open, 0 comments.

9. **[#5953](https://github.com/QwenLM/qwen-code/pull/5953) LSP Server support hot reload**  
   **What:** Detects changes to `.lsp.json` during an active session and reloads the LSP server configuration without restarting.  
   **Status:** Open, 0 comments.

10. **[#6287](https://github.com/QwenLM/qwen-code/pull/6287) Add proactive channel loop tools**  
    **What:** Adds MCP tools for recurring scheduled reminders in channel sessions. The scheduled loop runs through the same channel session for consistent routing.  
    **Status:** Open, 0 comments.

## 5. Feature Request Trends

- **Context compression UX improvements**: Users want to queue prompts during long `/compress` operations (#6331) and maintain an active input field. The community contributed PR #6336.
- **Extension/MCP lifecycle consistency**: Multiple issues (#6244, #6338, #6340) highlight a need for deterministic capability reporting and cache-friendly state management.
- **Proactive channel reminders**: Request for natural-language scheduling of recurring reminders within channel sessions (#6327, #6287) to align with enterprise collaboration patterns.
- **Model fallback resilience**: Users want automatic fallback to backup models on overload/errors (#6273), reflecting broader demand for production-grade high availability.
- **Multi-folder workspace support**: Developers using VSCode multi-root workspaces need CLI daemon to accept multiple workspace boundaries (#6278).

## 6. Developer Pain Points

- **Configuration/model token overflow**: Custom model context window calculations (#6266) and no-op `max_tokens` escalations (#5939) continue to cause head-scratching. A shared helper `escalatedOutputTokenLimit` is under review.
- **Windows-specific environment issues**: Extension installation (#6334) and terminal title spoliation (#6332) are recurring complaints. Windows remains a second-class citizen in certain UX paths.
- **Channel reliability (DingTalk)**: Two bugs (#6327, #6329) expose fragile MCP-channel bridge behavior — stalled processes, lost routing, and broken Markdown delivery. Production deployments are affected.
- **Security gaps in script isolation**: The `transform_data` vulnerability (#6282) shows that isolation wrappers are inconsistently applied. Developers expect sandbox guarantees to be enforced everywhere.
- **Documentation drift**: Outdated skill syntax docs (#6320) and broken memory preprocessor links (#2618) waste developer time. The community regularly submits doc PRs, suggesting that the codebase outpaces the documentation.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-07-05

## Today’s Highlights
A new **v0.8.67** release (PR #4034) lands today, adding the **LongCat (Meituan)** provider and bundling several post-merge follow-ups. Performance and UX fixes land for the TUI: composer input is no longer re-wrapped five times per frame (PR #3967), and provider URLs in narrow layouts are now readable (PR #4028). Localization work continues with a dedicated test locale enforcement PR (#4033) to prevent CI failures on non-English systems.

## Releases
No new releases in the last 24 hours. The pending v0.8.67 (PR #4034) is open and expected to close soon.

---

## Hot Issues

1. **#4032 – [bug] Codewhale not following the constitution**  
   Author: `stream2stream`  
   Codewhale repeatedly writes temporary scripts instead of using user‑provided ones, and justifies its behavior when challenged. Community reaction is muted (0 👍), but the issue touches core agent compliance.  
   [GitHub](https://github.com/Hmbown/CodeWhale/issues/4032)

2. **#4030 – [bug] panic on broken pipe (SIGPIPE)**  
   Author: `BrathonBai`  
   Piping `codewhale doctor | head` causes a noisy crash dump instead of a clean exit. A straightforward crash path that can frustrate users in scripting workflows.  
   [GitHub](https://github.com/Hmbown/CodeWhale/issues/4030)

3. **#3991 – [CLOSED] UX: /links provider URLs become unreadable in narrow TUI layouts**  
   Author: `Hmbown`  
   At 80 columns with sidebar, provider URLs collapse to fragments (e.g., “h”). Fixed in PR #4028. Community noted the annoyance; no comments but closed quickly.  
   [GitHub](https://github.com/Hmbown/CodeWhale/issues/3991)

4. **#3909 – [CLOSED] perf(tui): composer input is re-wrapped up to five times per frame**  
   Author: `Hmbown`  
   Each render frame performed redundant width‑aware wrapping of the composer input. Fixed by PR #3967. Performance regression tests will prevent recurrence.  
   [GitHub](https://github.com/Hmbown/CodeWhale/issues/3909)

---

## Key PR Progress

1. **#4034 – v0.8.67: LongCat provider + post-#3960 review follow-ups + version bump**  
   Author: `Hmbown`  
   Adds **LongCat** (Meituan) as a first‑class OpenAI‑compatible provider with alias, env key, default model, and endpoint. Also includes minor fixes after #3960.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/4034)

2. **#3969 – Add per-sub-agent provider routing**  
   Author: `heyparth1`  
   Introduces `[subagents.routes.<role>]` config to pin sub‑agent roles to specific providers/models. Enables hybrid sessions (e.g., local LM Studio for exploration, cloud for generation).  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/3969)

3. **#4028 – fix(tui): keep provider links readable in narrow layouts**  
   Author: `roian6`  
   Renders `/links` provider URLs as inline code instead of bare markdown, preserving readability and copyability in narrow terminal widths. Includes regression test.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/4028)

4. **#3973 – refactor(shell): split output buffer helpers**  
   Author: `cyq1017`  
   Moves shell output delta and tail buffer helpers into a dedicated module. No behavioral changes; prepares for future shell tool improvements.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/3973)

5. **#3967 – perf(tui): avoid redundant composer input wrapping per frame**  
   Author: `reidliu41`  
   Eliminates the five‑time‑per‑frame wrapping by caching the wrapped output. Direct fix for #3909.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/3967)

6. **#4033 – test: enforce English locale for hardcoded string assertions**  
   Author: `hongqitai`  
   Forces `Locale::En` in test setup to prevent failures on non‑English devices – a common pain point in the i18n effort.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/4033)

7. **#3963 – fix(mcp): only advertise list-resource meta-tools when resources exist**  
   Author: `h3c-hexin`  
   Prevents injecting `list_mcp_resources` and `list_mcp_resource_templates` into the model‑visible tool catalog unless configured servers actually expose resources. Reduces noise for MCP‑heavy setups.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/3963)

8. **#3781 – Feat/opencode zen provider (still open)**  
   Author: `snail-vs`  
   Adds an “OpenCode Zen” provider. Still awaiting review and possible conflicts with the new provider pattern.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/3781)

9. **#3818 – fix(tui): expand active tool run summaries**  
   Author: `cyq1017`  
   Ensures dense tool‑run summaries can also expand while in‑flight (not just after finalization). Includes regression test for #3256 edge case.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/3818)

10. **#3583 – refactor(localization): extract hardcoded texts into JSON and load via rust-i18n**  
    Author: `hongqitai`  
    Moves hardcoded localization strings from `localization.rs` to JSON files and introduces the `rust-i18n` crate. Foundation for full i18n support.  
    [GitHub](https://github.com/Hmbown/CodeWhale/pull/3583)

---

## Feature Request Trends

- **Provider extensibility & routing** – Multiple PRs and issues ask for more flexible provider configuration: per‑agent routing (#3969), new providers like LongCat (#4034) and OpenCode Zen (#3781).
- **Localization & i18n** – The ongoing migration to `rust-i18n` (#3583, #4033) shows strong community desire for multi‑language support, especially for hardcoded UI strings.
- **Performance & responsiveness** – Issues like #3909 (composer wrapping) and #4030 (SIGPIPE handling) reflect a focus on smooth TUI experience and crash‑resilient piping.
- **Narrow‑layout UX** – The `/links` readability fix (#3991/4028) indicates users value responsiveness at small terminal widths.

---

## Developer Pain Points

- **Panic on broken pipe** – #4030: `codewhale` crashes when piped output is prematurely consumed; developers expect clean SIGPIPE handling.
- **Redundant rendering** – #3909: Composer input wrapping five times per frame is a low‑hanging performance drag that was widely noted internally.
- **Test flakiness due to locale** – #4033 and related localization PRs highlight CI failures on non‑English systems, a recurring cause of blocked merges.
- **MCP tool clutter** – #3963: Injecting resource listing tools even when none exist frustrates users who want a clean model‑visible tool list.
- **Constitution compliance** – #4032: The “Codewhale not following the constitution” bug points to a deeper trust issue where the agent ignores user‑provided scripts and rationalizes its own.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*