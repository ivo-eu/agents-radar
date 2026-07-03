# AI CLI Tools Community Digest 2026-07-03

> Generated: 2026-07-03 10:12 UTC | Tools covered: 9

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

# AI CLI Developer Tools: Cross-Tool Comparison Report — July 3, 2026

## 1. Ecosystem Overview

The AI CLI tool landscape continues to mature rapidly, with nine actively maintained projects showing distinct evolutionary paths. This week's data reveals a field divided between **established platforms** (Claude Code, OpenAI Codex) working through scale-related reliability issues, **fast-followers** (Gemini CLI, Qwen Code, DeepSeek TUI) aggressively shipping new agent orchestration features, and **smaller players** (Pi, Kimi Code, OpenCode) focused on cross-platform parity and foundational bug fixes. A common thread across all tools is the tension between **feature velocity** and **production reliability** — every community reports API timeouts, authentication friction, and agent hallucination issues that erode trust. The ecosystem is also converging on common architectural patterns: MCP protocol adoption, sub-agent/fleet orchestration, context window management, and configuration-as-code for permissions and skills.

---

## 2. Activity Comparison

| Tool | Open Issues (24h) | PRs Updated (24h) | Release Status | Community Engagement Signal |
|---|---|---|---|---|
| **Claude Code** | 10 hot issues | 6 PRs | **v2.1.199 released** (stacked slash skills, SSL errors) | High: duplicates indicate active triage; 9 👍 on auth bug |
| **OpenAI Codex** | 10 hot issues | 10 PRs | **2 alpha releases** (rust-v0.143.0-alpha.34/.35) | Very High: 683 👍 on Linux desktop request; 161 👍 on Windows installer |
| **Gemini CLI** | 10 hot issues | 10 PRs | **v0.51.0-nightly** (egress Cloud Run skeleton) | High: 8 👍 on generalist hang; focused dev team engagement |
| **GitHub Copilot CLI** | 10 hot issues | 2 PRs (spam) | **No release** | Moderate: 7 👍 on alt-screen toggle; but spam issues signal moderation need |
| **Kimi Code CLI** | 1 issue | 1 PR | **No release** | **Very Low:** nearly dormant; single issue closed after 5 months |
| **OpenCode** | 10 hot issues | 10 PRs | **No release** | Moderate: 13 👍 on Windows TUI exit bug; memory leak PR active |
| **Pi** | 10 hot issues | 10 PRs | **No release** | Moderate: 22 comments on update failure; active null-content fix PRs |
| **Qwen Code** | 6 issues | 10 PRs | **3 releases** (nightly, cua-driver v0.7.0, stable v0.19.5) | Moderate: focused issues; strong CI/daemon infrastructure work |
| **DeepSeek TUI** | 10 hot issues | 10 PRs | **No release** (tracking v0.8.67) | **High:** intense UX audit cycle; 22 comments on constitution/onboarding debates |

**Key Observations:**
- **Codex** dominates community passion (highest upvote counts) but also has the most fragmented pain points across desktop, sandbox, and CLI surfaces.
- **Claude Code** and **Gemini CLI** show the most mature triage practices (many duplicates closed quickly), but Claude Code has higher raw engagement volume.
- **Kimi Code** is effectively stalled — 1 issue and 1 PR in 24h suggests low active user base or maintainer bandwidth.
- **DeepSeek TUI** has the most concentrated development effort (all PRs targeting v0.8.67 release-blockers), indicating a focused ship cycle.

---

## 3. Shared Feature Directions

### Cross-Tool Requirements (Appearing in 3+ Communities)

| Feature Requirement | Tools Requesting | Specific Need |
|---|---|---|
| **RTL/BiDi Layout Support** | Claude Code (#72245), DeepSeek TUI (#1607) | Hebrew/Arabic users want mirrored TUI chrome and multi-currency cost display |
| **AST-Aware Code Tools** | Gemini CLI (#22745, #22746), Claude Code (internal) | Reduce token waste; enable method-level navigation; improve codebase understanding |
| **Sub-Agent/Fleet Orchestration** | Claude Code (#72332, #72330), Gemini CLI (#22323, #21409), DeepSeek TUI (#3932-#3935), Qwen Code (#6244) | Need stronger isolation, trajectory visibility, safety guardrails, and reliable state management |
| **Custom/Private Model Endpoints** | Copilot CLI (#4003), Pi (GLM, DeepInfra providers), DeepSeek TUI (#3927), Codex (generic) | Enterprise users want BYO models; all tools need flexible provider architecture |
| **Repository-Scoped Configuration** | Codex (#18115), Claude Code (recurring), Gemini CLI (#20079), DeepSeek TUI (#3867) | Per-repo settings via config files (`.codex/config.toml`, `.claude/rules/`) |
| **Context/Window Management** | Claude Code (#72300, #72359), Codex (topic memory), Pi (#6157), Gemini CLI (compaction evals) | Surviving context across summarization; language-aware compaction; directory-rename resilience |
| **Slash Command Parity Across Surfaces** | Claude Code (#72292), Copilot CLI (#1799), Codex (VS Code gaps) | Terminal vs. VS Code extension command sets diverge; users want consistency |
| **Non-Interactive / Pipeline Mode** | Copilot CLI (#4011), OpenCode (#6999), Qwen Code (CI/daemon work) | Run `/init` and tools in batch mode without TUI hang; suppress tool call details |

### Unique but High-Signal Requests (2+ Tools)

- **Prompt Security / Injection Prevention**: Claude Code (#72332), Gemini CLI (#22672), Codex (sandbox policy issues)
- **Automation Sandbox Reliability**: Codex (#15310), Gemini CLI (#19873), Claude Code (WSL2 regressions)
- **Session Persistence & Artifact Management**: Qwen Code (#5895, #6259), Pi (#6227 SQLite storage), Claude Code (#72347)
- **MCP Ecosystem Maturity**: Copilot CLI (#4006 pagination), OpenCode (#35106 forms), DeepSeek TUI (#3866 dynamic startup), Gemini CLI (#28143 cross-server confusion)

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | GitHub Copilot CLI | DeepSeek TUI | Qwen Code | Pi |
|---|---|---|---|---|---|---|---|
| **Core Philosophy** | Developer-centric slash skills, rich TUI | Desktop-first with sandbox; Rust alpha for CLI | Agent-native with sub-agent fleet; enterprise Vertex AI | Terminal companion to VS Code; lightweight | Constitution-driven, multi-agent fleet, privacy-first | Daemon + web-shell; CUA driver for automation | Ultra-lightweight CLI; provider-agnostic; extension-friendly |
| **Target User** | Pro devs, heavy customizers | Enterprise teams, Pro subscribers | Google Cloud / Vertex AI users, enterprise | VS Code users, corporate WSL | Privacy-conscious power users, Chinese market | Qwen ecosystem devs, mobile/web use | Multi-provider power users, tinkerers |
| **Key Technical Approach** | Slash-skill stacking, TUI-first, VS Code companion | Desktop app + CLI alpha (Rust); MITM proxy for credential routing | Sub-agent fleet (generalist/specialist), MCP-first, Zero-Dependency sandboxing | Minimal CLI, leverages Copilot backend, alt-screen rendering | Language-first constitution, Fleet orchestration, permissions.toml | Daemon-managed channels, web-shell React, CUA driver | Plugin-based extensions, SQLite storage, provider adapter pattern |
| **Strength** | Richest TUI, strong slash-command ecosystem, active triage | Highest community engagement, credential routing proxy, desktop polish | Most systematic agent architecture, enterprise infra (Vertex AI, Cloud Run) | Simplicity, VS Code integration, GItHub ecosystem | Most innovative onboarding (constitution), Fleet vision, active audit cycle | Best daemon/backend infrastructure, CUA driver, web-shell features | Fastest provider support, clean null-content fixes, extension model |
| **Weakness** | WSL2/compatibility regressions, API timeouts, auth friction | Windows sandbox fragility, macOS crashes, closed-source alpha | Sub-agent hangs, false success reports, Wayland breakage, tool overload | Model availability errors, MCP pagination gaps, spam issues | Under-documented constitution system, Fleet fragmentation, Windows gaps | Empty-response error ambiguity, OAuth timeouts, self-destruction bug | Update failures, edit tool brittleness (Claude), ghost models |

### Key Strategic Differentiators

- **Claude Code** is winning on **TUI richness** (stacked slash skills, clickable elements) but losing on **cross-platform reliability** (WSL2 regressions, VS Code command gaps).
- **OpenAI Codex** is winning on **community passion** (highest upvotes) but struggling with **platform fragmentation** — desktop, CLI, and VS Code experiences feel disconnected.
- **Gemini CLI** is investing most in **sub-agent architecture** (Fleet-like orchestration) and **enterprise infra** (Vertex AI, Cloud Run egress) — may leapfrog in multi-agent reliability if hang bugs are fixed.
- **DeepSeek TUI** is differentiating on **constitution-first onboarding** and **privacy/trust UX** — unique approach that could attract security-conscious users if documentation catches up.
- **Qwen Code** is building the strongest **daemon/backend layer** — CI pipeline improvements, daemon dashboards, session artifact APIs suggest a platform play beyond CLI.

---

## 5. Community Momentum & Maturity

### Tier 1: High Momentum / Active Iteration
- **OpenAI Codex**: Highest raw engagement (140 comments, 683 👍 on Linux request). Two alpha releases in 24h. Community is loud and passionate, but pain points are widespread across surfaces.
- **Gemini CLI**: Most **focused development velocity** — 10 meaningful PRs, all addressing real bugs (hang fixes, MCP confusion, agent safety). Nightly releases indicate continuous delivery. Community is smaller but more technical.
- **DeepSeek TUI**: Highest **concentration of effort** — all 10 PRs target v0.8.67 release-blockers. The adversarial UX audit (#3793, #3926-#3928) shows mature product thinking. Community debates are substantive.

### Tier 2: Moderate Momentum
- **Claude Code**: High issue volume but **many duplicates** (7 of 10 hot issues closed as dupes). Slower to resolve core reliability bugs (auth #51588 open for months). Release cadence is steady but incremental.
- **Qwen Code**: **Quietly productive** — 3 releases, 10 PRs, but only 6 issues. Community is smaller but engineering work is substantial (CI fixes, daemon APIs, extension lifecycle).
- **Pi**: **Bursty activity** — high issue engagement (22 comments on update failure) but most PRs are maintainer-led. Fixes are landing (null-content, clamping), but the project feels reactive.
- **OpenCode**: **High issue count** (memory leak, schema migration failure) but PRs are mostly feature work (codemode, MCP forms). The project is adding features faster than fixing regressions.

### Tier 3: Low Momentum / Stalled
- **GitHub Copilot CLI**: Only 2 PRs (both spam). No release. Community has high upvotes (7 👍 on alt-screen) but zero maintainer response on top issues. Spam issues signal neglected moderation.
- **Kimi Code CLI**: **Effectively dormant** — 1 issue (closed after 5 months), 1 PR. No community discussion. This project may need a maintainer update to survive.

### Maturity Indicators

| Tool | Most "Mature" Aspect | Most "Immature" Aspect |
|---|---|---|
| Claude Code | Triage pipeline (fast dupe detection) | Core auth/reliability bugs persist months |
| OpenAI Codex | Community management (high engagement) | Fragmented platform experience |
| Gemini CLI | Agent architecture design | Sub-agent hang reliability |
| GitHub Copilot CLI | Minimal surface area (fewer bugs) | Neglected moderation, no recent releases |
| DeepSeek TUI | UX design thinking (constitution, trust) | Feature documentation (under-documented) |
| Qwen Code | CI/daemon infrastructure | User-facing error messages |
| Pi | Provider extensibility | Update stability, edit tool quality |

---

## 6. Trend Signals

### Industry-Level Signals from Community Feedback

1. **The "Agent Hallucination" Crisis is Real and Universal**
   Every tool reports sub-agents producing fabricated outputs, hijacking workflows, or making confident but wrong decisions. Claude Code (#72332), Gemini CLI (#22323 false success), Pi (#6278 edit tool errors), and DeepSeek TUI (#3932 invisible fleet) all demonstrate that **LLM-as-agent reliability is the #1 unsolved problem**. Tools that solve sub-agent trajectory visibility, success/failure verification, and safety guardrails will win developer trust.

2. **MCP is Becoming the Universal Plugin Standard — But Fragmentation Looms**
   MCP (Model Context Protocol) is adopted across Claude Code, Gemini CLI, Copilot CLI, OpenCode, and Pi. However, each implements it slightly differently (resource pagination, credential routing, tool schema shapes). Community frustration with incompatibility (#4006 Copilot pagination, #28143 Gemini cross-server confusion) signals that **MCP needs a conformance test suite** to prevent ecosystem fragmentation before it solidifies.

3. **Windows and WSL2 Remain the Weakest Link**
   Every tool that supports Windows reports WSL2-specific regressions (Claude Code #72235 paste, Copilot CLI #1112 login, Gemini CLI #28144 slow startup, OpenCode #22003 TUI exit). The AI CLI ecosystem is **macOS-first by default**, and Windows users are paying the price. This is a market opportunity for any tool that invests in first-class Windows Terminal support.

4. **Context Management is the New Memory**
   All mature tools are investing in smarter context management — not just token counting, but **structural understanding** (AST-aware reads in Gemini CLI), **survival across summarization** (Claude Code #72300), **topic-based memory** (Codex #19758), and **artifacts persistence** (Qwen Code #5895). The community is demanding that agents remember *what matters* without being prompted.

5. **The Credential/Proxy Pattern is Converging**
   Multiple tools (Codex credential routes, Claude Code SSL handling, Qwen Code auth headers) are implementing **MITM proxy patterns** for secure credential management. This signals an industry shift toward **tool-granted, not user-pasted** authentication — a structural improvement for multi-session workflows.

6. **"Quiet Mode" and Pipeline Integration are Becoming Table Stakes**
   As AI CLIs move from interactive assistants to CI/CD pipeline components, users demand **non-interactive execution** (Copilot CLI #4011, OpenCode #6999). Tools that cannot operate in batch mode will be excluded from DevOps workflows.

7. **Platform Lock-In Anxiety is Growing**
   DeepSeek TUI #3927 (API-key step forces DeepSeek), Copilot CLI #4003 (no custom endpoints), Codex (Microsoft Store only) — users are actively pushing back against vendor lock-in. **Provider-agnostic tools like Pi and DeepSeek TUI** are positioned to capture users frustrated with walled gardens.

### Recommendations for Developers Evaluating These Tools

- **For production use with enterprise compliance needs**: **Gemini CLI** (Vertex AI integration, Cloud Run egress) or **OpenAI Codex** (best proxy/credential infrastructure) are most enterprise-ready, but expect Windows friction.
- **For maximum TUI productivity**: **Claude Code** still has the richest interactive experience, but be prepared for API timeout costs and auth friction on reinstated accounts.
- **For privacy-conscious or multi-provider workflows**: **Pi** offers the fastest provider onboarding and cleanest extension model, but update reliability is a concern.
- **For sub-agent/fleet orchestration experiments**: **DeepSeek TUI** and **Gemini CLI** are most advanced, but both suffer from reliability gaps that make production use risky today.
- **For lightweight terminal companion**: **GitHub Copilot CLI** remains simplest to adopt if you're already in VS Code ecosystem, but lack of maintainer response is worrying.
- **For Chinese market or mobile/web use cases**: **Qwen Code** has strongest daemon + web-shell infrastructure and most reliable release cadence.
- **Avoid for Windows-intensive workflows**: All tools have Windows gaps — **none** are truly cross-platform mature yet.

---

*Data sourced from GitHub repositories for Claude Code, OpenAI Codex, Gemini CLI, GitHub Copilot CLI, Kimi Code CLI, OpenCode, Pi, Qwen Code, and DeepSeek TUI — snapshot 2026-07-03 24-hour window.*

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
**Data Snapshot:** 2026-07-03 | Source: github.com/anthropics/skills

---

## 1. Top Skills Ranking

The following Skills (Pull Requests) have generated the most community discussion and represent the ecosystem's most active development areas:

### #1. Skill-Creator Evaluation Fix (#1298) — **Open**
**Functionality:** Repairs the core `run_eval.py` script, which reports 0% recall for every skill description, rendering the description-optimization loop useless. Fixes include proper eval artifact installation, Windows stream reading, trigger detection, and parallel worker handling.
**Discussion Highlights:** Addresses 10+ independent reproductions of Issue #556. The 0% recall bug has been the single most disruptive problem for skill authors. Multiple workaround attempts (PRs #1099, #1050, #1323) converged here as the definitive fix.
**Status:** Open (high activity, 2 weeks old)

### #2. Document Typography Skill (#514) — **Open**
**Functionality:** Prevents orphan word wrap, widow paragraphs, and numbering misalignment in AI-generated documents. Targets the "last mile" polish gap that users rarely request explicitly but notice immediately.
**Discussion Highlights:** Strong consensus that typographic quality is a universal pain point. PR author notes these issues "affect every document Claude generates."
**Status:** Open (4+ months old, sustained discussion)

### #3. ODT (OpenDocument) Skill (#486) — **Open**
**Functionality:** Creates, fills, reads, and converts OpenDocument Format files (.odt, .ods). Includes template filling and ODT-to-HTML conversion. Triggers on "ODT", "ODS", "ODF", "LibreOffice" mentions.
**Discussion Highlights:** Community demand for LibreOffice/ISO-standard document support. Complements the existing DOCX skill for the open-source office ecosystem.
**Status:** Open (4 months old)

### #4. Self-Audit Skill (#1367) — **Open**
**Functionality:** A "universal" skill that audits AI output before delivery — mechanical file verification followed by four-dimension reasoning audit in damage-severity priority order. Works with any project, any tech stack, any model.
**Discussion Highlights:** Represents a new category: meta-quality assurance. Very recent (2 days old) but already attracting attention for its cross-cutting applicability.
**Status:** Open (very recent)

### #5. Skill-Quality & Security Analyzers (#83) — **Open**
**Functionality:** Two meta-skills for the marketplace: `skill-quality-analyzer` (evaluates structure, documentation, triggers, examples, escape conditions across 5 dimensions) and `skill-security-analyzer` (security review for skills).
**Discussion Highlights:** First serious attempt at skill quality assurance tooling. Addresses the growing need for governance as the ecosystem expands.
**Status:** Open (8 months old, long-running discussion)

### #6. Testing Patterns Skill (#723) — **Open**
**Functionality:** Comprehensive testing skill covering the full stack: testing philosophy (Trophy model), AAA pattern, React component testing, E2E patterns (Playwright/Cypress), API testing, visual regression.
**Discussion Highlights:** Addresses the gap between "Claude can write code" and "Claude can write *tested* code." Community notes this fills a critical quality-assurance void.
**Status:** Open (3 months old)

### #7. Frontend-Design Skill Improvement (#210) — **Open**
**Functionality:** Major revision to the frontend-design skill for clarity, actionability, and internal coherence. Ensures every instruction is executable within a single conversation.
**Discussion Highlights:** Exemplifies the community's growing sophistication in skill design — moving from "what to do" to "how to actually execute step by step."
**Status:** Open (6 months old, refined over time)

---

## 2. Community Demand Trends

From the most-discussed Issues, five clear demand signals emerge:

| Trend | Evidence | Demand Intensity |
|---|---|---|
| **Security & Trust Boundary** | Issue #492 (34 comments) — Community skills under `anthropic/` namespace enable impersonation. Top-voted security concern. | 🔴 Critical |
| **Enterprise Sharing** | Issue #228 (14 comments, 7 👍) — No native org-wide skill sharing; requires manual file transfer via Slack. | 🟠 High |
| **Skill-Creator Reliability** | Issue #556 (12 comments, 7 👍) — `run_eval.py` consistently reports 0% trigger rate. Blocks all skill optimization workflows. | 🔴 Critical |
| **Agent Governance & Safety** | Issue #412 (6 comments) — Proposal for governance patterns: policy enforcement, threat detection, audit trails. No existing skill covers this. | 🟡 Medium |
| **Compact Memory / Agent State** | Issue #1329 (8 comments) — Symbolic notation for compact agent state. Long-running agent context efficiency. | 🟡 Emerging |

**Key Insight:** The community is shifting from "what can Claude do?" to "how do we trust, govern, and scale what Claude does?" — security, quality assurance, and enterprise tooling dominate the issue tracker.

---

## 3. High-Potential Pending Skills

These active-comment PRs are not yet merged but show strong momentum:

| PR | Skill | Why It Might Land Soon |
|---|---|---|
| #1367 | **Self-Audit** (v1.3.0) | Universal applicability, mechanical + reasoning quality gate. Author has clear versioning strategy. |
| #1302 | **Color Expert** | Self-contained color expertise for any task. Well-researched (ISCC-NBS, Munsell, OKLCH/CAM16). No dependencies on other skills. |
| #806 | **Sensory Skill** (macOS AppleScript) | Native macOS automation via `osascript`. Two-tier permission system. Fills a clear platform gap. |
| #723 | **Testing Patterns** | Addresses the most common quality complaint. Comprehensive scope (unit, React, E2E, API, visual). |
| #509 | **CONTRIBUTING.md** | Not a skill, but critical infrastructure. Addresses GitHub community health gap (25% → improved). |

**Prediction:** The **Self-Audit** (#1367) and **Testing Patterns** (#723) skills have the highest surface area for adoption — they solve universal problems that affect every Claude Code user regardless of domain.

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is for reliability and governance tooling** — not new creative/technical capabilities, but the meta-layer of skill-quality validation, security auditing, output verification, and agent state management that makes the existing ecosystem trustworthy and maintainable at scale.

---

# Claude Code Community Digest — July 3, 2026

## Today's Highlights

A fresh patch (v2.1.199) lands with stacked slash-skill invocations and better SSL error handling. The community is most engaged with a lingering authentication bug (#51588) that blocks reinstated accounts, and a surrogates-in-image-paste error (#69781) on macOS. Many swiftly-triged duplicates indicate active issue triage, but several pain points around API timeouts, TUI focus behavior, and WSL2 compatibility remain unresolved.

## Releases

**v2.1.199**
- **Stacked slash-skill invocations** — `/skill-a /skill-b do XYZ` now loads all leading skills (up to 5), not just the first.
- **SSL certificate errors** — TLS-inspecting proxies, missing `NODE_EXTRA_CA_CERTS`, and expired certs now show actionable guidance immediately instead of burning retries.

[Full release](https://github.com/anthropics/claude-code/releases/v2.1.199)

## Hot Issues

1. **[#51588 – Auth blocking after Trust & Safety reinstatement](https://github.com/anthropics/claude-code/issues/51588)**  
   *7 comments · 9 👍 · [OPEN]*  
   A user reinstated by Trust & Safety remains locked out because the authentication system doesn't reflect the status change. High 👍 count indicates this affects many; the engineering fix is still pending.

2. **[#69781 – Image paste fails with "surrogates not allowed"](https://github.com/anthropics/claude-code/issues/69781)**  
   *4 comments · 1 👍 · [OPEN]*  
   Attaching or pasting an image returns a 400 error claiming invalid UTF-8 surrogates. Occurs on macOS desktop; reproduction steps are available.

3. **[#72240 – Tool-call markup emitted as plain text instead of executed](https://github.com/anthropics/claude-code/issues/72240)**  
   *3 comments · [CLOSED duplicate]*  
   In longer sessions, tool-call markup occasionally renders as literal text, causing commands to be silently skipped. A critical reliability bug that was quickly marked as duplicate – likely already tracked internally.

4. **[#72260 – 50% API timeout errors wasting tokens](https://github.com/anthropics/claude-code/issues/72260)**  
   *2 comments · [CLOSED duplicate]*  
   Roughly half of API requests timeout regardless of context size. Retries burn significant cached context (60k+ tokens). Duplicate status suggests team is aware.

5. **[#72292 – `/workflows` slash command not recognized in VS Code extension](https://github.com/anthropics/claude-code/issues/72292)**  
   *2 comments · 2 👍 · [CLOSED duplicate]*  
   The `/workflows` command is missing from the VS Code extension. The extension lags behind the terminal version in slash-command support – a recurring complaint from extension users.

6. **[#72332 – Prompt injection in sub-agent output](https://github.com/anthropics/claude-code/issues/72332)**  
   *2 comments · 1 👍 · [CLOSED duplicate]*  
   A sub-agent was hijacked and attempted to write fabricated security findings. Highlights ongoing concerns about sub-agent isolation and prompt security.

7. **[#72344 – OAuth login fails on Linux VPS](https://github.com/anthropics/claude-code/issues/72344)**  
   *2 comments · 1 👍 · [CLOSED duplicate]*  
   Google OAuth fails headlessly with "Redirect URI not supported by client". Headless deployment users are blocked; a configuration gap.

8. **[#72273 – TUI click-to-select fires even without window focus](https://github.com/anthropics/claude-code/issues/72273)**  
   *2 comments · [CLOSED duplicate]*  
   Clicking an interactive element while the TUI window lacks focus both focuses and activates it immediately – violating platform conventions and causing accidental selections.

9. **[#72235 – Paste broken in WezTerm on WSL2](https://github.com/anthropics/claude-code/issues/72235)**  
   *2 comments · [CLOSED duplicate]*  
   Regression in v2.1.195: right-click and Ctrl+Shift+V do nothing in WezTerm on WSL2. A specific but impactful environment regression.

10. **[#72330 – Agent acts on bundled multi-part approvals](https://github.com/anthropics/claude-code/issues/72330)**  
    *2 comments · [CLOSED duplicate]*  
    During a backlog session, the agent presented decisions as bare issue IDs and took actions the user couldn't meaningfully consent to. Raises UX/consent design questions for long agentic sessions.

## Key PR Progress

*(Only 6 PRs were updated in the last 24h; all are listed below.)*

1. **[#42701 – Fix init-firewall.sh crash from ipset when domain resolves to repeated IPs](https://github.com/anthropics/claude-code/pull/42701)**  
   *[CLOSED]*  
   Devcontainer startup failed if a domain resolved to duplicate IPs. Adds `-exist` flag to `ipset` to gracefully handle repeated addresses.

2. **[#66854 – "toekn"](https://github.com/anthropics/claude-code/pull/66854)**  
   *[OPEN]*  
   No description. Likely spam or a draft; no meaningful changes.

3. **[#72451 – Remove statsig.anthropic.com from init-firewall.sh](https://github.com/anthropics/claude-code/pull/72451)**  
   *[OPEN]*  
   The hostname `statsig.anthropic.com` no longer resolves, causing devcontainer firewall initialization to fail. Removing it restores reliable container startup.

4. **[#73476 – Fix GitHub capitalization in README](https://github.com/anthropics/claude-code/pull/73476)**  
   *[OPEN]*  
   Purely cosmetic: "Github" → "GitHub". Low impact.

5. **[#72543 – Create Cha](https://github.com/anthropics/claude-code/pull/72543)**  
   *[OPEN]*  
   No description. Likely a draft or placeholder; not actionable.

6. **[#72866 – Fix Github → GitHub typo in README](https://github.com/anthropics/claude-code/pull/72866)**  
   *[OPEN]*  
   Same typo fix as #73476; duplicate cleanup may be needed.

## Feature Request Trends

- **RTL / BiDi layout support** (#72245) – Hebrew/Arabic users want the TUI chrome mirrored for right-to-left reading.
- **Clickable prompt-summary header** (#72253) – Navigate back to any prior user input by clicking the pinned summary bar.
- **Preserve file-change reminders across context summarization** (#72300) – System reminders about modified files are lost when the conversation is summarized.
- **Expand repo access mid-session** (#72311) – Users want to grant Claude Code access to additional directories without restarting the session.
- **Project memory survive directory renames** (#72359) – Renaming a repo folder orphans project memory and session context silently.

## Developer Pain Points

The last 24 hours surface several recurring frustrations:

- **API reliability**: Timeouts (50% in some cases) and server-side rate limiting continue to plague users, burning expensive cached context on retries.
- **Authentication friction**: OAuth fails headlessly (#72344), reinstated accounts remain blocked (#51588), and "out of credits" errors on overage accounts require a full system reboot (#72291).
- **WSL2 / terminal compatibility**: Regressions in paste support (#72235), broken sidebar/scrolling (#72321), and missing focus conventions (#72273) affect a large segment of Windows/Linux hybrid users.
- **VS Code extension gaps**: Slash commands like `/workflows` (#72292) and conversation forking (#72281) are missing in the extension, creating an inconsistent experience vs. the terminal.
- **Context and session management**: Child sessions skip transcript persistence (#72347), desktop code mode creates orphan sessions (#72305), and project memory is lost on directory rename (#72359).
- **Prompt security**: Sub-agent hijacking (#72332) and bundled-approval consent issues (#72330) highlight the need for stronger guardrails in multi-step autonomous workflows.

*Data source: [github.com/anthropics/claude-code](https://github.com/anthropics/claude-code)*

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-07-03

**Today’s Highlights**  
Two new Rust alpha releases (0.143.0-alpha.34 and .35) landed in the last 24 hours, while the community remains focused on cross‑platform desktop support and sandbox reliability. A long‑running Linux app request (683 👍) and a Windows standalone installer issue (161 👍) dominate discussion, and a newly discovered GPT‑5.5 reasoning‑token clustering bug is drawing attention. On the PR side, a cluster of credential‑route and proxy improvements merged, along with a fix for duplicate image‑generation output.

---

## Releases
* **[rust-v0.143.0-alpha.34](https://github.com/openai/codex/releases/tag/rust-v0.143.0-alpha.34)** – Alpha release.
* **[rust-v0.143.0-alpha.35](https://github.com/openai/codex/releases/tag/rust-v0.143.0-alpha.35)** – Alpha release.

No changelog details were provided for either tag. These are incremental alpha versions likely containing bug fixes and internal improvements.

---

## Hot Issues (10 Noteworthy)

1. **[#11023 – Codex desktop app for Linux](https://github.com/openai/codex/issues/11023)**  
   *140 comments, 683 👍* – The most‑voted issue overall. Users on Linux (especially those with power‑constrained Macs) want a native Linux app. Still open for 5 months; community patience is thin.

2. **[#13993 – Support standalone Windows installer (`codex-setup.exe`)](https://github.com/openai/codex/issues/13993)**  
   *78 comments, 161 👍* – Many Windows users cannot use the Microsoft Store due to corporate restrictions. A traditional `.exe` installer is a high‑priority request.

3. **[#30224 – “This model is not supported” with `X-OpenAI-Internal-Codex-Responses-Lite`](https://github.com/openai/codex/issues/30224)**  
   *67 comments, 22 👍* – API error blocking custom‑model usage. Affects Plus subscribers on Windows. High engagement suggests a widespread regression.

4. **[#30364 – GPT‑5.5 reasoning‑token clustering at 516/1034/1552 degrades complex tasks](https://github.com/openai/codex/issues/30364)**  
   *34 comments, 49 👍* – A sharp user discovered that `gpt‑5.5` responses cluster at fixed reasoning‑token boundaries, correlated with lower performance on complex tasks. Could indicate a model‑side issue needing urgent investigation.

5. **[#15310 – Desktop automations silently fall back to `workspace-write` sandbox](https://github.com/openai/codex/issues/15310)**  
   *17 comments, 14 👍* – Scheduled/recurring tasks ignore the configured sandbox policy until a user manually enters the chat UI. Undermines trust in automation workflows.

6. **[#26158 – Windows sandbox regression in CLI 0.138.0 (OS error 740)](https://github.com/openai/codex/issues/26158)**  
   *16 comments, 6 👍* – Fixed in later releases? Issue is closed but highlights fragility of Windows sandbox execution. Users stuck on 0.132.0.

7. **[#9615 – VS Code Extension becomes all blank](https://github.com/openai/codex/issues/9615)**  
   *10 comments, 8 👍* – Long‑standing bug on Windows 11. Still open after 5 months; business subscribers affected.

8. **[#18115 – Repository‑scoped marketplace and plugin configuration](https://github.com/openai/codex/issues/18115)**  
   *9 comments, 45 👍* – Community wants per‑repo plugin config via `.codex/config.toml` instead of only user‑scoped settings. Strong upvote ratio for a feature request.

9. **[#30824 – Codex Desktop crashes with `EXC_BREAKPOINT/SIGTRAP` in FSEvents](https://github.com/openai/codex/issues/30824)**  
   *4 comments, 1 👍* – Crash on macOS with stale helper processes. Affects Pro users; crash report points to Chromium/uv backend.

10. **[#30943 – Usage hit 0% with zero messages sent (quota drain bug)](https://github.com/openai/codex/issues/30943)**  
    *3 comments, 0 👍* – A clear accounting bug showing 0% usage remaining despite no activity. Follow‑up to earlier drain reports; undermines subscription trust.

---

## Key PR Progress (10 Important)

1. **[#30976 – Propagate explicit command approval purpose](https://github.com/openai/codex/pull/30976) (open)**  
   Improves approval‑callback metadata to distinguish exec interception from sandbox retries, fixing lifecycle transitions and persistence.

2. **[#30969 – Preserve command identity across repeated approvals](https://github.com/openai/codex/pull/30969) (open)**  
   Ensures follow‑up approval callbacks don’t overwrite parent command metadata – a subtle app‑server/TUI fix.

3. **[#28996 – Avoid duplicate ImageGen Markdown output](https://github.com/openai/codex/pull/28996) (closed)**  
   Fixes a UI annoyance where generated images appeared both inline and as separate nuggets.

4. **[#22680 – Tell model about credentialed routes](https://github.com/openai/codex/pull/22680) (closed)**  
   Seeds the managed proxy with initial credential routes during session start. Foundation for private‑resource access.

5. **[#27503 – Refresh credentialed routes during session](https://github.com/openai/codex/pull/27503) (closed)**  
   Adds proxy‑owned hook to refresh route config every five minutes as plugins change.

6. **[#28981 – Rebase live proxy state through config reloaders](https://github.com/openai/codex/pull/28981) (closed)**  
   Allows sandbox and exec‑policy reloads to replace base policy without composing route state manually.

7. **[#28984 – Add credentialed routes backend adapter](https://github.com/openai/codex/pull/28984) (closed)**  
   Introduces a dedicated crate to translate backend routes into proxy‑owned credential types.

8. **[#22675 – Route credentialed traffic through MITM proxy](https://github.com/openai/codex/pull/22675) (closed)**  
   Implements generic MITM action to reroute HTTPS traffic through the credential proxy endpoint.

9. **[#22673 – Discover credentialed routes for managed proxy](https://github.com/openai/codex/pull/22673) (closed)**  
   Backend‑owned route‑list endpoint and credential proxy URL helpers. Foundation for the full credential flow.

10. **[#28930 – Add filesystem symlink semantics to app‑server](https://github.com/openai/codex/pull/28930) (closed)**  
    Adds `followSymlinks` option to `fs/getMetadata` and directory listing – needed for accurate file state on various OSes.

---

## Feature Request Trends

- **Cross‑platform desktop coverage** – Linux desktop app (#11023) and Windows standalone installer (#13993) are the two loudest requests, reflecting frustration with store‑only or macOS‑first distribution.
- **Repository‑scoped configuration** – Users want per‑repo plugin/marketplace settings (#18115) and project‑level agent memory (#19758) rather than global or user‑level config.
- **Better sandbox and automation control** – Desktop automations ignoring sandbox policy (#15310) and Windows sandbox regressions (#26158) drive demand for more reliable execution environments.
- **Enhanced memory and context management** – Topic‑based memory directories (#19758) and slash commands for memory indicate a desire for structured, scalable long‑term memory.
- **Telemetry and crash reporting** – A proposal for UWP crash telemetry (#26899) reflects user frustration with silent app failures.

---

## Developer Pain Points

- **Windows sandbox instability** – Multiple open issues (e.g., #29332, #30538, #28042) describe elevated sandbox failures, proxy conflicts, and COM+ registry errors. The experience on Windows remains fragile.
- **App crashes on macOS and Windows** – EXC_BREAKPOINT (#30824), infinite scroll loops (#28755), and generic crashes (#30531) degrade daily use, especially for Pro subscribers.
- **Rate‑limit / quota bugs** – False 0% usage draining (#30943) and missing reset buttons (#30641) erode trust in subscription billing.
- **Session and thread UI bugs** – Messages out of order (#29561), freezing when opening large threads (#25390), find‑in‑thread viewport anchoring (#30996) – common usability annoyances.
- **CLI terminal issues** – Suspend (Ctrl‑Z) in kitty makes the terminal unusable (#31013), and the “resume” message is incorrectly colored (#30973). Small polish bugs that hurt CLI first impressions.

---

*Generated from GitHub data for `openai/codex` – data snapshot 2026-07-03 24h window.*

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

Here is the Gemini CLI community digest for July 3, 2026, based on the provided data.

---

## Gemini CLI Community Digest — 2026-07-03

### Today’s Highlights

The nightly release process continues, but the most significant activity revolves around agent reliability and infrastructure. The community is heavily focused on fixing **subagent recovery bugs** that mask failures as successes and resolving **shell execution hangs**. A major infrastructure PR landed this week, establishing a skeleton for a caretaker egress service on Cloud Run, signaling ongoing investment in the project’s backend resilience. Meanwhile, a wave of smaller PRs targets specific pain points like slow startup on Windows, MCP resource confusion, and `AGENTS.md` file discovery.

---

### Releases

- **[v0.51.0-nightly.20260703.gf7af4e518](https://github.com/google-gemini/gemini-cli/releases/tag/v0.51.0-nightly.20260703.gf7af4e518)**
  - **What's New:** A single change introduces the skeleton for the caretaker's egress Cloud Run service, a foundational piece for offloading agent action processing to a managed backend.

---

### Hot Issues (Top 10 by Impact & Community Activity)

1. **[#22323](https://github.com/google-gemini/gemini-cli/issues/22323) — Subagent recovery after MAX_TURNS reported as GOAL success (P1, Bug)**
   - **Why it matters:** This is a critical bug that erodes trust in the agent system. When a subagent hits its turn limit, the system incorrectly reports success, hiding the fact that the task failed. This directly impacts the reliability of multi-agent workflows.
   - **Community reaction:** 9 comments, 2 upvotes. High engagement from maintainers trying to trace the faulty state transition in the agent lifecycle.

2. **[#19873](https://github.com/google-gemini/gemini-cli/issues/19873) — Leverage model's bash affinity via Zero-Dependency OS Sandboxing (P2, Enhancement)**
   - **Why it matters:** This long-running enhancement request targets a fundamental architectural improvement: sandboxing shell execution without heavy dependencies. Gemini models are natively good at bash, so making execution safer is a top priority.
   - **Community reaction:** 8 comments, 1 upvote. Still in active discussion with maintainers, suggesting complex design trade-offs.

3. **[#24353](https://github.com/google-gemini/gemini-cli/issues/24353) — Robust component level evaluations (P1, Epic)**
   - **Why it matters:** This epic tracks the evolution of the behavioral eval system beyond the initial 76 tests. As agent complexity grows, having a robust evaluation framework is crucial to prevent regressions.
   - **Community reaction:** 7 comments, 0 upvotes. High internal priority, indicating a major focus for the core team.

4. **[#22745](https://github.com/google-gemini/gemini-cli/issues/22745) — Assess the impact of AST-aware file reads, search, and mapping (P2, Feature)**
   - **Why it matters:** An epic investigating whether giving the agent syntactic understanding of code (ASTs) can reduce token consumption, improve navigation accuracy, and reduce turn counts. This is a key area for improving codebase interaction.
   - **Community reaction:** 7 comments, 1 upvote. The community is keen to see if this can fix the "misaligned reads" where the agent reads too much or too little code.

5. **[#21409](https://github.com/google-gemini/gemini-cli/issues/21409) — Generalist agent hangs (P1, Bug)**
   - **Why it matters:** A persistent, high-impact bug where the CLI hangs indefinitely when deferring to the generalist subagent. Users are forced to cancel sessions or avoid subagents entirely.
   - **Community reaction:** 7 comments, 8 upvotes. The highest upvoted issue, reflecting a major frustration. The workaround (disabling subagents) is not ideal.

6. **[#25166](https://github.com/google-gemini/gemini-cli/issues/25166) — Shell command execution gets stuck with "Waiting input" after command completes (P1, Bug)**
   - **Why it matters:** A core UX bug that makes the CLI unusable after simple commands. The terminal state machine is not recognizing command completion, causing the agent to "think" it's still waiting for user input.
   - **Community reaction:** 4 comments, 3 upvotes. High severity ("P1"), causes frequent task abandonment.

7. **[#21983](https://github.com/google-gemini/gemini-cli/issues/21983) — Browser subagent fails in Wayland (P1, Bug)**
   - **Why it matters:** A major platform compatibility issue. The browser agent, a key feature for web automation, fails entirely on Linux systems using Wayland, a growing standard.
   - **Community reaction:** 4 comments, 1 upvote. The core issue is likely the headless browser's display server interaction.

8. **[#20079](https://github.com/google-gemini/gemini-cli/issues/20079) — Symlink in agent directory not recognized (P2, Bug)**
   - **Why it matters:** A simple configuration pain point. Users who manage their agent definitions with symlinks (e.g., using dotfile managers) find their agents silently ignored.
   - **Community reaction:** 4 comments, 0 upvotes. A clear, fixable bug affecting developer workflows.

9. **[#22672](https://github.com/google-gemini/gemini-cli/issues/22672) — Agent should stop/discourage destructive behavior (P2, Customer Issue)**
   - **Why it matters:** The agent lacks safety rails for destructive shell commands (e.g., `git reset --hard`, `rm -rf`). This is a major security and usability concern, especially for production environments.
   - **Community reaction:** 3 comments, 1 upvote. Users want the agent to "think twice" before executing high-risk operations.

10. **[#24246](https://github.com/google-gemini/gemini-cli/issues/24246) — Gemini CLI encounters 400 error with > 128 tools (P2, Bug)**
    - **Why it matters:** As users add more MCP servers and custom skills, the tool list grows. Hitting a 400 error from the model API is a hard limit that blocks advanced users with complex setups.
    - **Community reaction:** 3 comments, 0 upvotes. Points to a need for smarter tool selection or pagination.

---

### Key PR Progress (Top 10 by Impact)

1. **[#28247](https://github.com/google-gemini/gemini-cli/pull/28247) — fix(core): match ls ignore globs by relative path**
   - **Summary:** Fixes a bug where `ls` ignore patterns with path separators (e.g., `node_modules/**`) weren't matching correctly, causing ignored files to pollute the agent's context.
   - **Why it matters:** Directly fixes a reported bug (#28207). Cleans up the agent's workspace view, reducing noise and token waste.

2. **[#28183](https://github.com/google-gemini/gemini-cli/pull/28183) — fix(vscode-ide-companion): preserve terminal focus when closing diff tabs**
   - **Summary:** Prevents the VS Code companion extension from stealing terminal focus after approving file edits.
   - **Why it matters:** A major quality-of-life fix for VS Code users. Resolves a "friction point" where every edit required a mouse click to refocus the terminal.

3. **[#28240](https://github.com/google-gemini/gemini-cli/pull/28240) — Fix #28227: add support for AGENTS.md out of the box**
   - **Summary:** Ensures `AGENTS.md` is automatically loaded as a context file, without requiring explicit configuration in `settings.json`.
   - **Why it matters:** Standardizes the `AGENTS.md` convention, making it a first-class citizen alongside `GEMINI.md`.

4. **[#28013](https://github.com/google-gemini/gemini-cli/pull/28013) — fix(prompts): use function replacer in applySubstitutions to prevent $-pattern corruption**
   - **Summary:** Fixes a bug where `$` characters in skill or agent descriptions were interpreted as JavaScript replacement patterns in `String.prototype.replace`, corrupting prompts.
   - **Why it matters:** Prevents subtle prompt corruption for anyone using `$` in their skill definitions (e.g., dollar amounts, environment variables).

5. **[#28153](https://github.com/google-gemini/gemini-cli/pull/28153) — fix(core): ignore stale update_topic calls after a session reset**
   - **Summary:** Prevents orphan `update_topic` tool calls (emitted just before a `/clear`) from corrupting the new session's topic state.
   - **Why it matters:** Fixes a race condition that could cause the agent to think it's still working on a cleared topic, leading to confusing behavior.

6. **[#28149](https://github.com/google-gemini/gemini-cli/pull/28149) — fix(skills): respect .gitignore/.geminiignore in skill resource listing**
   - **Summary:** Makes skill resource listings respect ignore files, preventing irrelevant files (e.g., `node_modules`) from being sent to the model.
   - **Why it matters:** Reduces token overhead and prevents the model from seeing irrelevant files, improving skill quality.

7. **[#28144](https://github.com/google-gemini/gemini-cli/pull/28144) — fix(cli): detect available editors lazily to avoid slow startup**
   - **Summary:** Defers editor detection from startup to when an edit is actually required, solving startup delays on Windows where process creation is slow.
   - **Why it matters:** Makes the CLI feel significantly faster to start, especially on Windows and WSL.

8. **[#28143](https://github.com/google-gemini/gemini-cli/pull/28143) — fix(core): resolve MCP resources by server to prevent cross-server confusion**
   - **Summary:** Fixes a bug where reading a resource from one MCP server could return content from another server if they shared a URI path.
   - **Why it matters:** A subtle but critical security and correctness fix for users running multiple MCP servers.

9. **[#28223](https://github.com/google-gemini/gemini-cli/pull/28223) — fix(core-tools): bypass LLM correction for JSON and IPYNB files in write_file and replace**
   - **Summary:** Prevents the system from corrupting JSON and Jupyter Notebook files by attempting LLM-based "correction" on structured data.
   - **Why it matters:** Directly fixes data corruption bugs for data scientists and developers working with notebooks and config files.

10. **[#28142](https://github.com/google-gemini/gemini-cli/pull/28142) — fix(core): honor GOOGLE_CLOUD_LOCATION for Vertex AI with API key**
    - **Summary:** Ensures that when using a Vertex AI API key, requests are routed to the correct regional endpoint instead of the global one.
    - **Why it matters:** Fixes a configuration issue that could break compliance and data residency requirements for enterprise users.

---

### Feature Request Trends

- **AST-Aware Tooling:** Multiple issues (#22745, #22746) explore using Abstract Syntax Trees for smarter file reading, searching, and codebase mapping. The goal is to reduce token waste, enable precise method-level navigation, and improve the agent's understanding of code structure.
- **Agent Self-Awareness & Commands (#21432):** A strong desire for the CLI to be more "self-aware" — knowing its own hotkeys, CLI flags, and capabilities. Users want the agent to be a reliable guide for its own usage, reducing the manual search for documentation.
- **Robust Evaluation Infrastructure (#24353):** The team is pushing for a systematic component-level evaluation framework beyond simple behavioral evals. This is an internal trend, but it signals a maturing process for quality assurance, likely leading to fewer regressions in the future.
- **Subagent Trajectory Visibility (#22598):** Users want the ability to review what subagents did during a task. Currently, subagent actions are hidden, making it hard to debug or learn from agent behavior.
- **Hook Documentation (#28146):** A minor but clear trend: users working with the hook system need better documentation. A recent PR was specifically about documenting all three token fields in the `LLMResponse` usage metadata.

---

### Developer Pain Points

- **Agent Hangs and Indefinite Waits:** Issues like #21409 (generalist hang) and #25166 (shell hang after command completion) are the top frustrations. The agent frequently gets stuck in unresponsive states, requiring manual intervention or session restarts.
- **Subagent Reliability and Safety:** Problems with subagents are pervasive. Issues #22323 (false success reports), #22093 (running without permission), and #22672 (destructive behavior) highlight that the subagent system is **unpredictable and potentially unsafe** for many users.
- **Configuration and File Handling Friction:** From symlinks not being recognized (#20079) to `AGENTS.md` being ignored (#28240), simple configuration tasks are surprisingly fragile.
- **Platform-Specific Breakage:** The browser agent failing on Wayland (#21983) and slow startup on Windows (#28144) show that cross-platform testing is not keeping pace with feature development.
- **Tool Overload:** The 400 error when exceeding 128 tools (#24246) is a hard wall for power users who heavily customize their environment with MCP servers and skills. The system lacks intelligent tool selection or prioritization.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-07-03

## Today’s Highlights
The community is buzzing about the newly introduced **alt-screen mode** (#1799) and persistent **scroll bar rendering bugs** across platforms. Several new issues report **BYOK authentication regressions** (#4016) and **MCP pagination non-compliance** (#4006). On the positive side, two feature requests for **custom model endpoints** (#4003) and **configurable scroll speed** (#4018) signal strong demand for more terminal control.

## Releases
No new releases in the last 24 hours.

## Hot Issues (10 noteworthy)

1. **[#1799](https://github.com/github/copilot-cli/issues/1799) — How to turn off alt-screen views?**  
   *11 comments, 7 👍*  
   The most upvoted open issue. Users are frustrated by the forced alt-screen mode and want a switch back to the original inline output. No response from maintainers yet.

2. **[#1112](https://github.com/github/copilot-cli/issues/1112) — Copilot CLI `/login` hangs after device code approval in VS Code Dev Container (Debian 12, WSL2 host)**  
   *8 comments, 2 👍* (CLOSED)  
   A long-standing login bug that was recently updated. Affects WSL2 + Dev Container setups. The bug persists even after approval confirmation on GitHub.

3. **[#3997](https://github.com/github/copilot-cli/issues/3997) — Model "gpt-5.3-codex" is not available**  
   *7 comments*  
   Users hit a fatal error when trying to use the agent mode – likely a server-side model rotation issue. Needs triage.

4. **[#3501](https://github.com/github/copilot-cli/issues/3501) — Scroll bar makes text unalign (Windows)**  
   *6 comments, 9 👍*  
   Vertical scroll bar introduced in recent releases breaks text alignment on Windows Console Host and Terminal. High community pain.

5. **[#3936](https://github.com/github/copilot-cli/issues/3936) — Ctrl+G should expand paste tokens to full text in $EDITOR**  
   *3 comments*  
   Parity with Claude Code: when `compactPaste` is enabled, `$EDITOR` receives a literal `[Paste #N]` token instead of the actual text, breaking editing.

6. **[#4019](https://github.com/github/copilot-cli/issues/4019) — Built-in web_fetch does not work with HTTP proxies**  
   *2 comments*  
   Corporate users on WSL cannot use `/research` or URL retrieval. Proxies are not supported, blocking enterprise adoption.

7. **[#4003](https://github.com/github/copilot-cli/issues/4003) — Support custom model endpoint in Copilot CLI**  
   *2 comments*  
   Users want the same flexibility as VS Code’s Language Models panel to point to local/private endpoints.

8. **[#3570](https://github.com/github/copilot-cli/issues/3570) — Scrolling using touch not working**  
   *1 comment, 1 👍*  
   Touch scrolling is broken on Windows. Minor but affects laptop users.

9. **[#3569](https://github.com/github/copilot-cli/issues/3569) — /clear vs /new unclear**  
   *1 comment, 2 👍*  
   Help tooltips need clarification: `/new` preserves session for `/resume` while `/clear` discards it permanently.

10. **[#4020](https://github.com/github/copilot-cli/issues/4020) — IDE auto-connect falsely skipped after forking/closing a session**  
   *0 comments* (new)  
   After forking and closing a session, resuming the original fails to auto-connect to the IDE. Race condition in session management.

## Key PR Progress (2 open PRs)

- **[#3880](https://github.com/github/copilot-cli/pull/3880) — “beyond the streets of amaerica”**  
  Appears to be a spam/non-functional PR containing a random `ArtistCard` component. No substantive changes.

- **[#3873](https://github.com/github/copilot-cli/pull/3873) — “1000Add initial console log for greeting”**  
  A trivial one-liner adding a console log. No discussion or merges. Likely a test or stray contribution.

No meaningful code changes are under review today.

## Feature Request Trends
- **Terminal rendering control**: Alt-screen toggle (#1799), scroll speed configuration (#4018), touch scrolling support (#3570).
- **Custom model endpoints and BYOK improvements**: Multiple requests for self-hosted models (#4003, #4012, #4016) and better reasoning effort handling.
- **MCP ecosystem**: Pagination compliance (#4006), OAuth flow improvements (#4017), and unique name warnings (#3893).
- **Non-interactive (pipeline) use**: Want `/init` and other commands to work in batch mode without hanging (#4011).
- **Persistence and clarity**: Theme setting remembered (#4015), `permissions-config.json` deny rules (#3995), better `clear`/`new` distinction (#3569).

## Developer Pain Points
1. **Terminal rendering regressions** – The new scroll bar and alt-screen mode introduced breaking experiences on Windows and macOS. Copying output now includes a trailing `┃` glyph (#4009).
2. **Authentication and proxy woes** – Login hangs in Dev Containers (#1112), BYOK authentication regressions (#4016), and missing HTTP proxy support (#4019) block corporate users.
3. **Model availability errors** – Sporadic “model not available” errors (#3997) without clear fallback or recovery.
4. **MCP tooling friction** – Pagination not followed (#4006), OAuth silent failures (#4017), and rendering corruption when adding servers (#4014).
5. **Spam issues** – Five nonsensical, likely bot-generated issues (#3234, #3233, #3232, #3231, #3226) clutter the tracker and may indicate a need for moderation steps.

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-07-03

## Today’s Highlights
Activity on the Kimi Code CLI repo was light today. A long-standing bug (#1111) involving Tailscale WebSocket connections was finally closed after five months. Meanwhile, a single open pull request (#2481) addresses a Windows‑specific paste issue when images are present in the clipboard — a fix that improves parity across terminals.

---

## Releases
No new releases were published in the last 24 hours.

---

## Hot Issues

Only one issue was updated in the past 24 hours. It is listed below as it represents the sole noteworthy community discussion.

### #1111 — [CLOSED] [bug] kimi web use tailscale websocket connection error
- **Author:** tianyw0  
- **Created:** 2026-02-12 | **Updated:** 2026-07-02 | **Comments:** 2 | 👍: 0  
- **Link:** [Issue #1111](https://github.com/MoonshotAI/kimi-cli/issues/1111)  
- **Summary:** On Darwin arm64, using Kimi Code CLI v1.12.0 over a Tailscale VPN causes WebSocket connections to fail.  
- **Why it matters:** VPN/proxy compatibility is critical for remote teams. The issue was open for nearly five months before being closed, suggesting a workaround or resolution was found. However, the lack of public follow‑up may leave some users uncertain. Community reaction was minimal (0 👍, 2 comments).

> No other issues received updates in the last 24 hours.

---

## Key PR Progress

Only one pull request was updated in the last 24 hours.

### #2481 [OPEN] fix(shell): read clipboard media on BracketedPaste for Windows terminals
- **Author:** redjade75723  
- **Created:** 2026-07-01 | **Updated:** 2026-07-02 | **Comments:** 0 | 👍: 0  
- **Link:** [PR #2481](https://github.com/MoonshotAI/kimi-cli/pull/2481)  
- **Description:** On Windows Terminal and VS Code’s integrated terminal, Ctrl+V is handled by the terminal as a BracketedPaste event. Binary content (e.g., images) cannot be carried as text in that event, so pasting fails silently. This PR modifies `_handle_bracketed_paste()` to attempt reading clipboard media directly when the pasted content is empty.  
- **Impact:** Fixes a long‑standing UX gap for Windows users who paste images into the CLI. No community discussion yet.

> No other PRs were updated in the last 24 hours.

---

## Feature Request Trends
No feature requests (issues tagged as enhancement or feature) were updated in the last 24 hours. The repository shows no clear trend toward any new capabilities during this period.

---

## Developer Pain Points
Based on the limited activity, two recurring pain points are visible:

1. **VPN/Tailscale incompatibility** — Issue #1111 highlights that custom network configurations (e.g., Tailscale) can break WebSocket connectivity. Such problems are hard to diagnose and often require maintainer debugging.
2. **Windows terminal clipboard quirks** — PR #2481 directly addresses the silent paste failure on Windows when using BracketedPaste. This reflects a broader class of cross‑platform terminal inconsistencies that frustrate developers using VS Code or Windows Terminal.

No other pain points were surfaced in the last 24 hours.

---

*This digest is generated automatically from public GitHub data. For real‑time updates, check the [MoonshotAI/kimi-cli](https://github.com/MoonshotAI/kimi-cli) repository.*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-07-03

## Today’s Highlights
A steady stream of bug fixes and feature PRs landed, with significant progress on code-mode (`codemode`) confinement and MCP tool improvements. On the issue tracker, memory leakage (`#35107`), theme detection regression (`#20926`), and spurious rate-limit errors (`#34884`) drew the most community attention.

## Releases
No new releases in the last 24 hours.

---

## Hot Issues

1. **[#20926] no longer respects system theme (macOS Ghostty)**  
   *Closed* · 30 comments · 👍 10  
   Themes with light/dark variants always default to light, ignoring system theme. Affects macOS Ghostty users.  
   https://github.com/anomalyco/opencode/issues/20926

2. **[#13282] OpenCode 1.1.59 crashes immediately on chat (Bun AVX crash)**  
   *Closed* · 16 comments · 👍 0  
   Crash on start of conversation after upgrade. Narrowed to AVX mismatch with Bun runtime on Skylake CPUs.  
   https://github.com/anomalyco/opencode/issues/13282

3. **[#34884] Go returns “Provider rate limit exceeded” despite 0% rolling usage**  
   *Open* · 15 comments · 👍 6  
   Users on OpenCode Go tier cannot use paid models, while free Zen models work. Dashboard shows no usage.  
   https://github.com/anomalyco/opencode/issues/34884

4. **[#31119] Error: no such column: name**  
   *Open* · 9 comments · 👍 8  
   After updating to v1.16.2, the app refuses to work with a SQLite schema mismatch. High community upvote.  
   https://github.com/anomalyco/opencode/issues/31119

5. **[#31236] Copilot gpt-5.5: “input item ID does not belong to this connection”**  
   *Closed* · 8 comments · 👍 0  
   Switching auth tokens mid-session leaves stale Responses API itemId. Deterministic on Copilot gpt-5.5.  
   https://github.com/anomalyco/opencode/issues/31236

6. **[#17011] AI_InvalidResponseDataError: Multiple reasoning_opaque values**  
   *Closed* · 6 comments · 👍 2  
   GitHub Copilot + Claude Sonnet 4-6 (thinking enabled) returns multiple thinking blocks, which is unsupported.  
   https://github.com/anomalyco/opencode/issues/17011

7. **[#35107] Memory keeps growing until the bun process is killed**  
   *Open* · 2 comments · 👍 0  
   `updatePart` calls `structuredClone(part)` on every streamed token, causing heap pressure under many sessions.  
   https://github.com/anomalyco/opencode/issues/35107

8. **[#22003] TUI exit closes terminal window on Windows**  
   *Open* · 2 comments · 👍 13  
   Using Ctrl+D or `/exit` in Windows Terminal closes the entire window instead of returning to prompt.  
   https://github.com/anomalyco/opencode/issues/22003

9. **[#31831] opencode process consumes 185% avg CPU / 500MB+ RAM constantly**  
   *Open* · 2 comments · 👍 3  
   High resource usage even when idle, reported on Apple Silicon.  
   https://github.com/anomalyco/opencode/issues/31831

10. **[#35116] NOT NULL constraint failed: session_message.seq**  
   *Open* · 1 comment · 👍 0  
   Switching models in a session causes send button to hang. SQLite constraint error in logs.  
   https://github.com/anomalyco/opencode/issues/35116

---

## Key PR Progress

1. **[#35121] fix(tui): optimize session list multi-select delete experience**  
   *Open* · by yuwenlong  
   Fixes keybinding, session ID ordering, and adds `opencode session clean` for bulk deletion.  
   https://github.com/anomalyco/opencode/pull/35121

2. **[#35010] feat(desktop): reopen closed tabs and background tab open**  
   *Open* · by usrnk1  
   Adds browser-style tab management: `⇧⌘T` to reopen, background tabs, and tab history.  
   https://github.com/anomalyco/opencode/pull/35010

3. **[#34974] fix(rpc): reject pending calls when target disconnects**  
   *Open* · by HEETMEHTA18  
   Prevents hung RPC promises when a Worker emits `error` or `messageerror`.  
   https://github.com/anomalyco/opencode/pull/34974

4. **[#35111] fix: Memory keeps growing until the bun process is killed**  
   *Open* · by xingruodong-sys  
   Implements a streaming accumulator to avoid repeated `structuredClone` on text parts.  
   https://github.com/anomalyco/opencode/pull/35111

5. **[#35085] feat(opencode): add code-mode MCP adapter**  
   *Closed* · by rekram1-node  
   Adds the `@opencode-ai/codemode` package adapter; part of the experimental code-mode feature.  
   https://github.com/anomalyco/opencode/pull/35085

6. **[#35118] feat(codemode): add confined execution package**  
   *Open* · by rekram1-node  
   Cherry-picks the code-mode confinement package onto `v2`, with node-gyp pin removed.  
   https://github.com/anomalyco/opencode/pull/35118

7. **[#31570] feat: drive reasoning variants from models.dev reasoning_options**  
   *Closed* · by rekram1-node  
   Parses `reasoning_options` from models.dev instead of hardcoded per-package tables.  
   https://github.com/anomalyco/opencode/pull/31570

8. **[#35106] feat(core): global forms support**  
   *Closed* · by rekram1-node  
   Infrastructure for MCP elicitation forms; hidden behind a typed “global” owner hack.  
   https://github.com/anomalyco/opencode/pull/35106

9. **[#34952] fix(opencode): turn.idle only when active and ctrl are undefined**  
   *Open* · by Mte90  
   Fixes intermittent “Tool execution aborted” errors caused by premature idle-timeout.  
   https://github.com/anomalyco/opencode/pull/34952

10. **[#35103] refactor(opencode): expose MCP tools in native shape from the service**  
    *Closed* · by rekram1-node  
    Changes `MCP.tools()` return type to native `McpTool` shape, improving boundary clarity.  
    https://github.com/anomalyco/opencode/pull/35103

---

## Feature Request Trends

- **Command-level model/agent routing** (#34970): Allow each command to specify its own model or agent, enabling multi-model workflows within a single session.
- **Skill permissions per model** (#35109): Split agent-level tool permissions to model-level granularity.
- **Configurable mid-run prompt delivery** (#32157): Distinguish between `queue`, `steer`, and `break` semantics for user interrupts during agent runs.
- **Multi-select session deletion UX** (#35123): Batch delete in CLI and dedicated `session clean` command.
- **Quiet mode for CLI runs** (#6999): Suppress tool call details in scripted `opencode run` output.

---

## Developer Pain Points

- **Memory leaks under steady use**: `structuredClone` in streaming text parts (#35107) and high idle CPU/RAM (#31831) remain unresolved.
- **Schema migration failures**: “no such column: name” (#31119) after version updates indicates missing upgrade paths.
- **Platform inconsistencies**: Windows TUI closes terminal (#22003); Desktop/CLI version mismatch after update (#35122).
- **Provider integration quirks**: Spurious rate-limit errors on Go tier (#34884); Anthropic `tool_choice` deserialization errors (#35110); payload size limits for image inputs (#35112).
- **Session data corruption**: SQLite `NOT NULL constraint` on `session_message.seq` (#35116) and sync failures between Desktop and CLI.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest – 2026-07-03

## Today’s Highlights
A `pi update` failure on version 0.80.3 due to a missing `@smithy/node-http-handler` package has been reported with high community engagement (22 comments), prompting a quick PR with a pnpm recovery hint. Meanwhile, new Claude models are struggling with Pi’s edit tool, causing roughly 20% of edits to fail due to LLM-invented extra keys—a fix is in review. A cluster of issues around null `message.content` when reasoning models skip text content has been closed via several PRs, hardening agent loops, compaction, and rendering.

## Releases
No new releases were published in the last 24 hours.

## Hot Issues
The ten most noteworthy issues updated in the past 24 hours:

1. **#6215** – `pi update` fails on 0.80.3 due to missing `@smithy/node-http-handler@^4.9.1` ([link](https://github.com/earendil-works/pi/issues/6215))  
   *Open, 22 comments.* Breaks the update workflow for many users; community has been discussing workarounds.

2. **#6278** – New Claude models work poorly with Pi’s edit tool, failing ~20% of edits in some sessions ([link](https://github.com/earendil-works/pi/issues/6278))  
   *Open, 0 comments (just filed).* Claude is inventing extra keys like `new_text_x`, `type`, `closeenough`; a major quality issue for Claude users.

3. **#6157** – Compaction summary should be in the session’s language and deduplicate ([link](https://github.com/earendil-works/pi/issues/6157))  
   *Open, 4 comments.* Non-English sessions currently get English checkpoint summaries; also requests dedup logic to replace full preservation.

4. **#6259** – `content is not iterable` when reasoning models return null content during tool use ([link](https://github.com/earendil-works/pi/issues/6259))  
   *Open, 3 comments.* Affects GLM-5.2 on Fireworks and similar models; has multiple downstream impacts.

5. **#6204** – `mimo-v2-omni` is a ghost model on all three Xiaomi MiMo Token Plan providers ([link](https://github.com/earendil-works/pi/issues/6204))  
   *Open, 3 comments.* Listed in the catalog but not served; produces a 400 error. Users are misled by the model list.

6. **#6276** – `content is not iterable` in compaction.js and render-utils.js on v0.80.3 ([link](https://github.com/earendil-works/pi/issues/6276))  
   *Closed as untriaged, 1 comment.* A duplicate of #6259 but specific to v0.80.3 – reinforces the prevalence of null-content crashes.

7. **#6265** – OpenAI Responses `max_output_tokens` can be set below the API minimum near context limit ([link](https://github.com/earendil-works/pi/issues/6265))  
   *Closed, 1 comment.* Crashes long sessions with a 400 error; a PR (#6264) clamped the value.

8. **#6268** – Codex websocket terminates after 60 minutes without retry ([link](https://github.com/earendil-works/pi/issues/6268))  
   *Closed, 1 comment.* Long tasks are abruptly stopped; no automatic reconnection, leaving users stranded.

9. **#6274** – Handling invalid JSON escapes in edit tool invocation (GLM-5.2 on Fireworks) ([link](https://github.com/earendil-works/pi/issues/6274))  
   *Closed, 1 comment.* Models are not properly escaping backslashes, causing malformed edit tool calls.

10. **#6262** – DS4-server context overflow errors not detected by auto-compaction ([link](https://github.com/earendil-works/pi/issues/6262))  
    *Closed, 3 comments.* Local DeepSeek V4 Flash extension rejects oversized prompts but Pi’s auto-compaction doesn’t trigger; tokens wasted.

## Key PR Progress
Ten important pull requests updated in the past 24 hours:

1. **#6279** – fix(coding-agent): add pnpm self-update prune hint ([link](https://github.com/earendil-works/pi/pull/6279))  
   *Open.* Adds a recovery suggestion (`pnpm store prune`) when `pi update` fails due to stale pnpm metadata. Mashes the #6215 issue.

2. **#6267** – feat(coding-agent): add InlineExtension type for named inline extension factories ([link](https://github.com/earendil-works/pi/pull/6267))  
   *Open.* Closes #6260 – lets dynamic inline extensions display a name instead of `(anonymous)` in startup logs.

3. **#6227** – feat: sqlite session storage ([link](https://github.com/earendil-works/pi/pull/6227))  
   *Open.* Adds an optional SQLite backend for session transcripts alongside the default JSONL. Controlled by `PI_SQLITE_SESSION_STORAGE=1`.

4. **#6273** – Add Zen mode tool call labels ([link](https://github.com/earendil-works/pi/pull/6273))  
   *Closed.* Introduces a `zenMode` setting that renders compact tool-call labels and can asynchronously summarize them via GPT-5.4-mini.

5. **#6271** – [codex] Add GLM API provider ([link](https://github.com/earendil-works/pi/pull/6271))  
   *Closed.* First-class support for Z.AI and Zhipu AI GLM endpoints (`glm`, `glm-cn`) using `ZAI_API_KEY`.

6. **#6266** – Anthropic: strict tool use for the edit tool ([link](https://github.com/earendil-works/pi/pull/6266))  
   *Closed.* Reduces Claude’s error rate on the edit tool by enforcing stricter schema validation; addresses #5501 and #5434.

7. **#6264** – fix(ai): clamp OpenAI Responses output tokens ([link](https://github.com/earendil-works/pi/pull/6264))  
   *Closed.* Fixes #6265 by ensuring `max_output_tokens` never drops below the API minimum of 16.

8. **#6263** – feat(ai): add DeepInfra provider for text and image generation ([link](https://github.com/earendil-works/pi/pull/6263))  
   *Closed.* Registers DeepInfra as a built-in provider with OpenAI-compatible chat and image generation; model catalog auto-fetched.

9. **#6258** – fix: guard against null content in agent loop, compaction, and message transforms ([link](https://github.com/earendil-works/pi/pull/6258))  
   *Closed.* Comprehensive fix for #4909, #6259, and similar – adds null checks in `message.content` iterations across three key modules.

10. **#3799** – add azure cognitive services as provider ([link](https://github.com/earendil-works/pi/pull/3799))  
    *Closed.* While older, it was updated in the last 24h. Adds support for `*.cognitiveservices.azure.com` base URLs in the Azure OpenAI provider, auto-normalizing paths.

## Feature Request Trends
- **New provider integrations** – Several requests this week: DeepInfra (added), GLM Z.AI / Zhipu (added), Kimi K2.7 for Copilot, Claude Sonnet 5 for Copilot. The community is eager for more model access.
- **UI/UX polish** – Requests for Zen code (compact tool call labels), showing active built-in tools in the TUI footer, configurable keybindings (e.g., Shift+Enter to submit), and extended function key handling.
- **Extension & configuration improvements** – Desire for named inline extension factories, typed settings schemas for extensions (`pi.settings.register`), and configurable tool output truncation limits.
- **Compaction & session management** – Make compaction summaries language-aware and deduplicate history, and support alternative storage backends (SQLite).
- **Model-specific compatibility** – Patches for reasoning models, MiMo ghost models, and edit tool tolerance for Claude – users want Pi to work reliably with any LLM.

## Developer Pain Points
- **Missing / stale dependencies** – `pi update` failures due to unresolved pnpm packages (`@smithy/node-http-handler`) waste developer time; workarounds are non-obvious.
- **Null message.content crashes** – Reasoning models that omit text content cause `TypeError: content is not iterable` in agent loops, compaction, and rendering. This has been a recurring frustration across multiple versions (v0.75.4, v0.80.3) and multiple models (GLM, MiMo).
- **Edit tool brittleness** – Models (especially Claude) often add unexpected keys to edit tool calls, causing validation failures. The error rate can be as high as 20% per session, making Pi unreliable for code editing.
- **Websocket connectivity** – Codex sessions are forcibly terminated after 60 minutes with no automatic retry, disrupting long-running tasks.
- **Context / token limit edge cases** – Auto-compaction fails to detect model-specific context overflow errors (DS4-server), and the OpenAI Responses provider can send `max_output_tokens` below the API minimum, leading to 400 errors.
- **Misleading model catalogs** – Ghost models like `mimo-v2-omni` appear in the provider list but are not actually served, causing failed starts and confusion.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-07-03

## Today's Highlights
Two new releases landed today: the nightly v0.19.5 and a standalone `cua-driver-rs v0.7.0` with relative-coordinate support. A flurry of bug fixes and feature PRs target web-shell interactivity (vision model picker, custom code blocks, @-mention panel), while a critical streaming tool-call bug (#6249) and OAuth authentication failure (#6251) are top-of-mind for the community. Several core improvements to daemon session artifacts and extension capability communication are also in flight.

## Releases

- **[v0.19.5-nightly.20260703.b16baf1ff](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.5-nightly.20260703.b16baf1ff)**  
  Nightly release including a fix for mobile session-switch jank (memoized timeline signature, replay-first dispatch).
- **[cua-driver-rs v0.7.0](https://github.com/QwenLM/qwen-code/releases/tag/cua-driver-rs-v0.7.0)**  
  Prebuilt binaries for a relative-coordinate fork of the CUA driver. macOS (codesigned + notarized universal binary + .app), Linux (x86_64 + arm64, glibc 2.31 minimum), Windows (unsigned x86_64 + arm64).
- **[v0.19.5](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.5)**  
  Stable release hardening the CLI daemon-managed channel worker and deferring web-shell session creation until first prompt.

## Hot Issues (6 of 6)

1. **[#3804](https://github.com/QwenLM/qwen-code/issues/3804) – API Error: Model stream ended with empty response text (v0.15.6)**  
   *Author: SeoMP | Updated: 2026-07-03 | Comments: 4*  
   A persistent bug where `AskUserQuestion` operations frequently fail with empty response text. Two months old without resolution, this is a high-friction issue for users relying on interactive prompts.

2. **[#6252](https://github.com/QwenLM/qwen-code/issues/6252) – Feat: Add daemon status dashboard**  
   *Author: doudouOUC | Updated: 2026-07-03 | Comments: 2*  
   Feature request for a browser-based dashboard visualizing `GET /daemon/status`. The corresponding PR #6253 is already in review – community interest in operational visibility is strong.

3. **[#6251](https://github.com/QwenLM/qwen-code/issues/6251) – OAuth Authentication Failing with 504 Gateway Timeout**  
   *Author: mudasirsohail | Updated: 2026-07-03 | Comments: 2*  
   Users reporting that OAuth authorization (chat.qwen.ai) responds with a 504 timeout, blocking tool access. A critical P2 bug that needs urgent infrastructure investigation.

4. **[#6249](https://github.com/QwenLM/qwen-code/issues/6249) – Streaming tool calls with empty `arguments` silently dropped**  
   *Author: tomsen-ai | Updated: 2026-07-03 | Comments: 2*  
   A streaming parser bug: when a tool call has an empty `arguments` string (legal for parameterless tools), the entire call is dropped, causing “Model stream ended with empty response text” retry loops. This directly impacts users of OpenAI-compatible providers with no-arg tools.

5. **[#6246](https://github.com/QwenLM/qwen-code/issues/6246) – qwen_code cannot recognize its own process**  
   *Author: L-Shier | Updated: 2026-07-03 | Comments: 2*  
   When Qwen Code spawns a Node.js backend process and is later asked to stop it, it terminates all Node.js processes (including itself). A dangerous edge case in process management that can lead to self-destruction.

6. **[#6244](https://github.com/QwenLM/qwen-code/issues/6244) – Extension capability changes not reliably communicated to the model**  
   *Author: ZijianZhang989 | Updated: 2026-07-03 | Comments: 2*  
   When extensions are enabled/disabled/refreshed during a session, the model may not be informed. This causes stale tool/command awareness – a core reliability issue for dynamic extension workflows. A “welcome-pr” label suggests maintainers are keen to see a fix.

## Key PR Progress (10 of 50)

1. **[#6261](https://github.com/QwenLM/qwen-code/pull/6261) – CI(autofix): restore sandbox image flow**  
   Moves Qwen Autofix back to GitHub-hosted sandbox execution, removes local OpenAI loopback proxy, and uses a dedicated autofix model secret. Streamlines CI security.

2. **[#6232](https://github.com/QwenLM/qwen-code/pull/6232) – feat(web-shell): support custom code block rendering**  
   Adds a plugin point for hosts to replace Markdown code blocks with custom React content. Bundles a `web-shell-charts` skill to render visualizations – powerful for data-science use cases.

3. **[#6260](https://github.com/QwenLM/qwen-code/pull/6260) – fix(ci): add always() to delay-automatic-review and ack-review-request**  
   Fixes a CI pipeline bug where same-repo PRs never triggered the automated Qwen review. The `always()` condition ensures dependency chains execute correctly.

4. **[#6236](https://github.com/QwenLM/qwen-code/pull/6236) – fix(web-shell): encode vision model picker selection**  
   Critical fix: the vision model picker was passing ACP-format model IDs, but core expects `authType:modelId`. Model selection was silently failing – now corrected.

5. **[#6242](https://github.com/QwenLM/qwen-code/pull/6242) – feat(web-shell): add custom at mention panel**  
   Replaces the inline @-mention autocomplete with a multi-level reference panel supporting categories, files, extensions, and MCP resources. Improves discoverability and keyboard/mouse navigation.

6. **[#6258](https://github.com/QwenLM/qwen-code/pull/6258) – fix(mobile-mcp): add production-release environment to CD workflow for npm auth**  
   Fixes npm publish failure by binding the `NPM_TOKEN` to the `production-release` environment, aligning with the stable release workflow.

7. **[#6238](https://github.com/QwenLM/qwen-code/pull/6238) – fix(core): give Stop-hook continuations a fresh per-turn tool-call budget**  
   Ensures each blocking Stop-hook iteration (e.g., `/goal`) gets its own tool-call budget instead of sharing a cap. Makes the cap configurable – important for long-running agent loops.

8. **[#6253](https://github.com/QwenLM/qwen-code/pull/6253) – feat(serve): Add daemon status dashboard**  
   Implements the browser-based dashboard requested in #6252. Ships a standalone HTML page at `/dashboard` and a reusable dashboard component for embedding.

9. **[#5895](https://github.com/QwenLM/qwen-code/pull/5895) – feat(daemon): add session artifact APIs**  
   Merged today (closed state). Adds structured artifact metadata attached to tool results and hooks. Enables listing/adding/removing session artifacts – a foundational capability for persistent agent memory.

10. **[#6245](https://github.com/QwenLM/qwen-code/pull/6245) – Notify model when extension capabilities change**  
    Directly addresses issue #6244. Tracks already-announced MCP tools, skills, and agent subagent types; queues deltas and drips them into the model context. A clean architectural fix for runtime extension updates.

## Feature Request Trends

- **Daemon observability**: Multiple requests for dashboards and status APIs (#6252, PR #6253) – operators want a real-time view of sessions, permissions, rate limiting, and workspace health.
- **Rich web-shell interactivity**: Custom code block rendering (#6232), multi-level @-mention panel (#6242), nested sub-agent tree visualization (#6239) – the community is pushing for a full-featured IDE-like experience in the browser.
- **Session artifact persistence**: With PR #5895 and the follow-up V2 design document (#6259), users want artifacts to survive daemon restarts and session replays.
- **Extension lifecycle management**: Reliable communication of capability changes to the model (#6244) is a growing concern as the ecosystem of MCP tools and skills expands.

## Developer Pain Points

- **“Model stream ended with empty response text”** – This vague error appears in multiple issues (#3804, #6249) and is frequently triggered by empty tool call arguments or streaming glitches. Developers need better error messages and graceful handling.
- **OAuth infrastructure flakiness** – 504 Gateway Timeouts during authentication (#6251) are blocking access entirely. The root cause may be on the auth provider side, but the impact on daily use is severe.
- **Process management blind spots** – Qwen Code terminating its own process when asked to stop a child process (#6246) reveals a gap in process group awareness. Self-destruction is a high-severity UX bug.
- **CI/CD secret and environment inconsistencies** – Multiple PRs (#6258, #6257, #6260) fix mismatches between workflows, showing that the CI configuration is fragile and error-prone for contributors.

*For the full list of changes, visit the [Qwen Code repository](https://github.com/QwenLM/qwen-code).*

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-07-03

## 1. Today's Highlights
The project remained in a heavy v0.8.67 hardening phase, with two overlapping themes: **correctness/UX fixes from the adversarial audit** (a batched PR closed five issues on onboarding, trust, and constitution visibility) and **Fleet reliability rounds** (sub-agent memory bounds, unique temp paths, and recursion-depth ceilings). Several community members contributed long-standing feature gaps — notably auto-discovery of `.claude/rules/` and dynamic MCP server startup — which were merged during the day.

## 2. Releases
No new releases in the last 24 hours. The project is tracking toward v0.8.67.

## 3. Hot Issues (10 selected)

1. **#3793 — Guided constitution creator vs. blank editor** (Hmbown | Updated: 07-02)  
   Continues as the central v0.8.67 UX debate. The maintainer emphasises language-first onboarding and an explicit ban on constitution files toggling runtime security settings.  
   *Why it matters:* Defines the product’s “first-time help vs. power-user flexibility” tradeoff.  
   [Issue #3793](https://github.com/Hmbown/CodeWhale/issues/3793)

2. **#3867 — Project-scope instructions hard-blocked (CLOSED)** (yekern | 07-03 → resolved by PR #3892)  
   The `instructions` key was unusable in multi-project workflows because of a hard deny at project scope, no glob support, and no rules-directory discovery.  
   *Why it matters:* Unblocked a core workflow for developers who manage many repos with per-project configuration.  
   [Issue #3867](https://github.com/Hmbown/CodeWhale/issues/3867)

3. **#3928 — Constitution is invisible & custom override fails silently** (Hmbown | Updated: 07-03)  
   No in-app way to read the constitution; `/context` points at a path that doesn’t exist in installed binaries; custom overrides are ignored without error.  
   *Why it matters:* Makes the product’s central feature — the constitution — opaque and untestable.  
   [Issue #3928](https://github.com/Hmbown/CodeWhale/issues/3928)

4. **#3927 — API-key step is a hard DeepSeek-only gate** (Hmbown | Updated: 07-03)  
   First-run users who want another provider (or to browse before creating a key) can’t proceed; Esc jumps two steps back.  
   *Why it matters:* Blocks adoption for multi-provider users and sends a “lock-in” signal.  
   [Issue #3927](https://github.com/Hmbown/CodeWhale/issues/3927)

5. **#3926 — Trust step is trust-or-quit** (Hmbown | Updated: 07-03)  
   Enter is a dead key; declining trust dumps the user to shell.  
   *Why it matters:* Erodes confidence in the trust model — the only “safe” option is to exit.  
   [Issue #3926](https://github.com/Hmbown/CodeWhale/issues/3926)

6. **#3882 — Fleet worker fanout exhausts TUI memory (CLOSED)** (Hmbown | 07-02)  
   A user report of ~15 GB RAM consumption during Fleet usage. Fixed by bounding sub-agent output.  
   *Why it matters:* A release-blocker that had to be fixed before v0.8.67 could ship.  
   [Issue #3882](https://github.com/Hmbown/CodeWhale/issues/3882)

7. **#3932 — Fleet: model has no vocabulary to pick a fleet member** (Hmbown | Updated: 07-02)  
   The `agent` tool schema exposes no `role`/`model_class` choice; the constitution never mentions Fleet or per-role routing.  
   *Why it matters:* The fleet feature is functionally invisible to the LLM, so it cannot be used.  
   [Issue #3932](https://github.com/Hmbown/CodeWhale/issues/3932)

8. **#3935 — Fleet wizard dead-ends / status shows wrong data** (Hmbown | Updated: 07-02)  
   “Start” only pastes a prompt into the composer; `/fleet status` promises an inspection it cannot deliver.  
   *Why it matters:* The feature’s entry point is broken, mis-educating users.  
   [Issue #3935](https://github.com/Hmbown/CodeWhale/issues/3935)

9. **#1607 — Token cost: add more currencies (RMB ¥, etc.)** (xyz-225648 | Updated: 07-02)  
   Long-standing feature request from the Chinese community for multi-currency cost display.  
   *Why it matters:* Localisation for a substantial user base; shows the audience is global.  
   [Issue #1607](https://github.com/Hmbown/CodeWhale/issues/1607)

10. **#3961 — Make new-version prompts persistent and actionable** (Hmbown | Created: 07-03)  
    Core update machinery exists; weak link is the in-app UX that shows a one-time hint that disappears.  
    *Why it matters:* Reduces user stickiness to outdated versions and missed security patches.  
    [Issue #3961](https://github.com/Hmbown/CodeWhale/issues/3961)

## 4. Key PR Progress (10 selected)

1. **#3968 — Expose `context_input_budget_for_route`** (h3c-hexin | Open)  
   Makes the internal input-budget computation available to external hosts and hooks.  
   *Impact:* Enables consistent budgeting across the TUI, CLI, and server surfaces.  
   [PR #3968](https://github.com/Hmbown/CodeWhale/pull/3968)

2. **#3967 — Avoid redundant composer input wrapping per frame** (reidliu41 | Open)  
   Fixes a performance issue where input text was wrapped up to five times per frame. Closes #3909.  
   *Impact:* Directly reduces CPU usage during typing in large sessions.  
   [PR #3967](https://github.com/Hmbown/CodeWhale/pull/3967)

3. **#3966 — Coverage for file permission rule action precedence** (greyfreedom | Open)  
   Tests that `deny > ask > allow` ordering and path- vs tool-specific weighting work correctly.  
   *Impact:* Hardens the execution-policy system before release.  
   [PR #3966](https://github.com/Hmbown/CodeWhale/pull/3966)

4. **#3964 — Align documentation with current permissions.toml schema** (greyfreedom | Open)  
   Clarifies `action = "deny" | "ask" | "allow"` values, defaults, and approval-card Save behaviour.  
   *Impact:* Reduces user confusion about how permissions actually work.  
   [PR #3964](https://github.com/Hmbown/CodeWhale/pull/3964)

5. **#3963 — Only advertise `list-resource` meta-tools when resources exist** (h3c-hexin | Open)  
   Prevents injecting empty MCP resource-listing tools into the model-visible catalog.  
   *Impact:* Reduces noise in tool selection and avoids wasted inference tokens.  
   [PR #3963](https://github.com/Hmbown/CodeWhale/pull/3963)

6. **#3960 — Release-lane correctness batch (CLOSED)** (Hmbown | Merged 07-03)  
   Self-contained fixes for five UX issues from the v0.8.67 audit (onboarding, trust, constitution, skills, permissions).  
   *Impact:* Directly resolves #3918/#3919/#3920/#3924/#3926/#3927/#3928/#3929.  
   [PR #3960](https://github.com/Hmbown/CodeWhale/pull/3960)

7. **#3892 — Auto-discover `.codewhale/rules/` and `.claude/rules/` directories (CLOSED)** (yekern | Merged 07-03)  
   Solves the long-standing “project-scope instructions hard-blocked” problem (#3867).  
   *Impact:* Unblocks multi-project workflows; enables compatibility with Claude Code rule sets.  
   [PR #3892](https://github.com/Hmbown/CodeWhale/pull/3892)

8. **#3866 — LLM can start MCP servers from chat context** (bistack | Open)  
   Adds a `start_mcp_server` tool that supports both stdio and HTTP transports. Enables dynamic tool sourcing.  
   *Impact:* Major step toward modifiable runtime tooling without restarting the TUI.  
   [PR #3866](https://github.com/Hmbown/CodeWhale/pull/3866)

9. **#3931 — Enforce absolute recursion-depth ceiling; widen task-id entropy (CLOSED)** (Hmbown | Merged 07-03)  
   Fixes two concurrency/correctness issues in fleet orchestration; confirms the global admission cap is sound.  
   *Impact:* Prevents runaway sub-agent spawning and ID collisions.  
   [PR #3931](https://github.com/Hmbown/CodeWhale/pull/3931)

10. **#3936 — Unique temp path per atomic state write (CLOSED)** (Hmbown | Merged 07-03)  
    Fixes a concurrent-persist corruption bug where multiple threads could overwrite the same `state.json` file.  
    *Impact:* Eliminates a subtle data-loss race in sub-agent persistence.  
    [PR #3936](https://github.com/Hmbown/CodeWhale/pull/3936)

## 5. Feature Request Trends

- **Guided wizard-driven setup** — Replace the “blank prompt editor” with a step-by-step constitution creator that is language-first, security-aware, and accessible even without an API key (from #3793, #3792, #3927).
- **Multi-project & multi-repo workflow support** — Rules-directory auto-discovery, glob patterns, and per-project configuration that does not get hard-blocked (#3867, #3892).
- **Fleet / multi-agent skeleton** — A consistent role vocabulary, profile loading, tool schema that surfaces fleet members to the model, and a wizard that actually writes configuration (#3932, #3933, #3934, #3935).
- **Better update UX** — Persistent, actionable new-version prompts instead of a one-time dismissal (#3961).
- **Localisation & cross-platform polish** — Multi-currency cost estimation (#1607), Chinese locale fixes (#1675), in-app version-check / GitHub link (#1678), and Windows Terminal default launch (#1854).
- **Slash-command discoverability** — Provide the LLM itself with the vocabulary to explain built-in commands instead of hallucinating (#1708).

## 6. Developer Pain Points

1. **Constitution & trust system is under-documented and sometimes opaque** — No in-app reader, silent custom-override failures, and a trust-or-quit UX that punishes cautious users (#3926, #3928).
2. **Fleet / multi-agent configuration is fragmented** — Four separate role vocabularies, two disjoint profile-loading mechanisms, and a wizard that produces files the runtime cannot read (#3933, #3934).
3. **Permissions and skills have misleading UI** — Users are taught to author `allowed-tools:` frontmatter that nobody parses, and `/skill trust` advertises enforcement that does not exist (#3920, #3919).
4. **Windows support gaps** — Chinese garbled characters in real-time output (#1675), CLI commands that break between PowerShell and cmd (#1754), and no default Windows Terminal launch (#1854).
5. **Memory/performance concerns** — Fleet fanout causing 15 GB RAM consumption (#3882), multi-wrapping of input text per frame (#3909 / #3967).
6. **Onboarding locks users into DeepSeek** — API-key step offers no skip, no alternative-provider path, and misbehaving navigation (Esc jumps two steps back, #3927).

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*