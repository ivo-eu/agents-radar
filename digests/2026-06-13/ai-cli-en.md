# AI CLI Tools Community Digest 2026-06-13

> Generated: 2026-06-13 10:15 UTC | Tools covered: 9

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
**Date: 2026-06-13**

---

## 1. Ecosystem Overview

The AI CLI tools ecosystem is experiencing a bifurcation between rapid feature iteration and reliability crises. Seven major tools tracked today show a collective 220+ issues and 80+ PRs updated in 24 hours, indicating intense developer engagement—but also mounting frustration. The dominant themes are multi-agent orchestration (every major tool), terminal rendering fragility (especially with non-Latin scripts and streaming), and escalating model quality concerns that erode trust in AI-assisted workflows. Notably, the ecosystem is splitting along provider lines: OpenAI, Anthropic, and Google are addressing infrastructure reliability and sub-agent scaling, while projects like Pi, Qwen Code, and CodeWhale (forked from DeepSeek TUI) are racing to add provider diversity and ACP compliance. A critical undercurrent is the growing demand for transparent billing and resource governance, particularly as token-based consumption models frustrate users accustomed to request-based pricing.

---

## 2. Activity Comparison

| Tool | Open/Updated Issues (24h) | Active PRs (24h) | Release Status | Notable Signal |
|------|--------------------------|------------------|----------------|----------------|
| **Claude Code** | 10 hot issues | 1 PR | ✅ v2.1.176 → v2.1.177 | Model quality degradation reports surging |
| **OpenAI Codex** | 10 of 30 most active | 10 of 50 recent | ✅ 3 alpha releases (v0.140.0) | MultiAgentV2 encrypted-tools regression blocking Pro users |
| **Gemini CLI** | 10 of 50 | 10 of 35 | ✅ v0.48.0 nightly | Agent hangs and subagent recovery top complaints |
| **GitHub Copilot CLI** | 10 selected | 0 | ✅ v1.0.62-1 → v1.0.62-2 | Terminal rendering corruption most upvoted bug |
| **Kimi Code CLI** | 3 | 5 | ⏸️ No release (latest v0.12.0) | Billing controversy (#1994) highest risk issue |
| **OpenCode** | 10 | 10 | ⏸️ No release | Multi-agent orchestration PR (#32167) day's biggest contribution |
| **Pi** | 37 | 16 | ✅ v0.79.2 → v0.79.3 | Context window fix prevents billing overruns |
| **Qwen Code** | 10 | 10 | ✅ v0.18.0 | OAuth free-tier quota reduction debate (129 comments) |
| **CodeWhale** (DeepSeek TUI) | 10 | 10 | ✅ v0.8.59 (rebrand) | Agent Fleet control plane epic active |

*Key: ✅ Active releases | ⏸️ No release in last 24h*

---

## 3. Shared Feature Directions

| Requirement | Tools Expressing Need | Specific Use Case |
|------------|----------------------|-------------------|
| **Multi-agent orchestration** | Claude Code, OpenAI Codex, Gemini CLI, Copilot CLI, OpenCode, CodeWhale | Task queues, nested sub-agents, parallel workers, session trees, `/swarm` mode |
| **Custom slash commands / configurable prompts** | Copilot CLI (#618, 99👍), Claude Code, OpenCode | `.github/prompts/` directory; reduce fixed token overhead |
| **AST-aware code navigation** | Gemini CLI (EPIC #22745), Claude Code | Replace regex grep with semantic code understanding |
| **Voice transcription for TUI** | OpenAI Codex (#14630, 44👍), OpenCode.. | Parity with desktop app transcription |
| **Background / persistent agents** | Qwen Code (durable cron #5004), OpenCode, CodeWhale, Claude Code | Survive restarts; non-blocking lint/build tasks |
| **MCP compliance / plugin discovery** | OpenCode (#28567, 20👍), Gemini CLI, Pi, CodeWhale | Full resource/tool/root support; ACP-registry listing |
| **Session persistence & organization** | Claude Code, OpenAI Codex, Gemini CLI, Kimi Code | Folders, tags, sync reliability; avoid history loss |
| **Granular billing / cost tracking** | Claude Code, Copilot CLI, Kimi Code, Pi | Token usage per-turn; OpenTelemetry cost metrics; transparent quota |
| **Safety filter / false positive reduction** | Claude Code, OpenAI Codex, Copilot CLI | Benign DevOps ops blocked; `<ip_reminder>` injections |
| **Terminal rendering with non-Latin scripts** | Claude Code (#66586 Bengali), Copilot CLI | Combining marks, complex script alignment |

---

## 4. Differentiation Analysis

| Dimension | Market Leaders | Differentiators |
|-----------|---------------|-----------------|
| **Target User** | **Claude Code & Copilot CLI** – General developers; **OpenAI Codex** – Pro/Enterprise sub-agent users | **Pi & Qwen Code** – Power users with custom providers; **CodeWhale** – Multi-agent orchestration enthusiasts |
| **Technical Approach** | **Anthropic** – Incremental CLI releases, model-quality tracking; **OpenAI** – Rust alpha hotfixes, encrypted-tool infrastructure; **GitHub** – Plugin marketplace, diff search | **Pi** – Extremely wide provider support (Bedrock, Vertex, Kimi, vLLM); **CodeWhale** – Agent Fleet control plane with headless runtimes; **OpenCode** – Nested sub-agent spawning (up to 5 levels) |
| **Provider Lock-in** | **Claude Code** – Anthropic models; **Gemini CLI** – Google models; **Copilot CLI** – GitHub Copilot (Azure OpenAI) | **Pi, Qwen Code, CodeWhale, Kimi Code** – Multi-provider first-class support; **OpenCode** – OmniRoute custom provider gateway |
| **Maturity & Stability** | **Copilot CLI** – Most stable, but rendering bugs; **Claude Code** – Moderate, but quality complaints rising | **Kimi Code** – Early v0.12, significant billing confusion; **CodeWhale** – Rebranding turbulence but high velocity |
| **Community Governance** | **Claude Code & OpenAI Codex** – Closed/core maintainer; **Gemini CLI** – Google-led | **OpenCode, Pi, Qwen Code, CodeWhale** – Open source with active community PRs |

---

## 5. Community Momentum & Maturity

**Highest Momentum** (rapid iteration, high community engagement):
- **CodeWhale** – Rebranded from DeepSeek TUI, pushing v0.8.60 with Agent Fleet orchestration, 10+ PRs in 24h, community contributing provider routes
- **Qwen Code** – Durable cron (#5004) merged, DaemonTransport/ACP compliance (#5040) advancing, 129 comments on OAuth policy debate
- **Pi** – 37 issues and 16 PRs updated, two patch releases in 24h, wide provider integration (Vertex, Bedrock Mantle)

**Established but Challenged** (large userbase, reliability concerns):
- **Claude Code** – Model quality degradation (#66502, #66522) and task queue demand (#33323, 24👍) indicate users want more than incremental updates
- **OpenAI Codex** – Encrypted-tools error cascade (8+ issues) blocking Pro/Enterprise; macOS relaunch loop (#25882) remains unresolved
- **Gemini CLI** – Agent hang/stall bugs (#21409, 8👍) and subagent false success (#22323) erode trust

**Niche but Growing**:
- **Kimi Code** – Low activity (3 issues), but #1994 billing controversy is a churn risk; PRs focused on MCP stability
- **OpenCode** – Multi-agent orchestration PR (#32167) is ambitious; MCP compliance (#28567, 20👍) is strongest demand

**Maturity Assessment**:
- *Most mature*: Copilot CLI (stable releases, plugin marketplace), Claude Code (systematic CLI updates)
- *Rapidly maturing*: Pi (27 releases tracked), Qwen Code (infrastructure PRs)
- *Early stage*: Kimi Code (v0.12, limited ecosystem), CodeWhale (post-rebrand, architecture in flux)

---

## 6. Trend Signals

1. **Multi-Agent Orchestration is the Battleground** – Every major tool is investing in sub-agent teams, task queues, and parallel execution. OpenAI Codex's MultiAgentV2 regression is the canary in the coal mine: scaling sub-agents introduces encrypted-tool, permission, and lifecycle management challenges that will define the next 6 months of CLI tool development.

2. **Provider Diversity is a Competitive Moat** – Pi, Qwen Code, and CodeWhale are winning community goodwill by supporting 10+ providers (Amazon Bedrock, Vertex AI, Kimi, Z.ai, StepFlash, vLLM, etc.). Meanwhile, Claude Code and Gemini CLI remain locked to single model families, which users increasingly see as a limitation as multi-provider workflows become standard.

3. **Billing Transparency is a Trust Accelerator** – Kimi Code's #1994 (7👍) and Pi's context window fix (#5644) highlight that opaque token-based billing is a churn risk. Tools that expose per-turn cost, quota usage, and allow cost caps (like Pi's 272k-token backend limit) will build more trust than those that don't.

4. **Terminal Rendering is a Universal UX Gap** – Non-Latin script corruption (Claude Code #66586), streaming artifacts (Copilot CLI #3749), narrow window crashes (Kimi Code #2450), and scroll bar misalignment (Copilot CLI #3501) show that TUI frameworks are not keeping pace with demand for international, resize-friendly, and high-performance rendering.

5. **Safety Filter False Positives are a Growing Friction Point** – Claude Code, OpenAI Codex, and Copilot CLI all report benign operations (e.g., `git gc`, opening a clean session) being blocked. This indicates safety systems are not well-tuned for CLI use cases, wasting developer time and eroding trust in AI assistance.

6. **Custom Provider Support is No Longer Optional** – With Pi normalizing $0.01/1M tokens via Bedrock Mantle, CodeWhale adding Z.ai and StepFlash first-class routes, and Qwen Code implementing DaemonTransport for ACP compliance, developers expect their CLI tool to work with their preferred backend—not force them into a single provider's ecosystem.

7. **Session Persistence Remains Fragile** – OpenAI Codex (#20741, 39 comments), Claude Code (#66529), and OpenCode (#32149) all report disappearing conversations, sync corruption, or session termination. As developers rely on long-running CLI sessions for project context, data loss is a critical reliability gap.

8. **Keyboard Layout Support is Overlooked** – Copilot CLI's German keyboard bug (#1999, 9 comments) and Gemini CLI's terminal inconsistencies show that non-US developers are underserved. As AI CLI tools go global, AltGr and input method support will become a differentiator.

---

*Report generated from community digest data for 2026-06-13. All issue/PR references are from the respective GitHub repositories. This is a speculative future analysis for illustrative purposes.*

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
*Data as of 2026-06-13 | Source: github.com/anthropics/skills*

---

## 1. Top Skills Ranking

*Ranked by community discussion activity (comments + cross-references)*

### #1: Document Typography Skill
**PR #514** — *Add document-typography skill*  
[github.com/anthropics/skills/pull/514](https://github.com/anthropics/skills/pull/514)

**Functionality:** Prevents orphan word wrap (1–6 words on next line), widow paragraphs (stranded headers at page bottom), and numbering misalignment in AI-generated documents. Addresses a universal pain point across Claude's document output.

**Discussion Highlights:** Community strongly resonated with the problem statement — "these issues affect every document Claude generates." Discussion focused on scope boundaries (should it also handle table formatting?) and integration with existing document skills.

**Status:** Open. No merge activity detected. High community interest, but author has not responded to review feedback since March.

---

### #2: OpenDocument Format Skill
**PR #486** — *Add ODT skill*  
[github.com/anthropics/skills/pull/486](https://github.com/anthropics/skills/pull/486)

**Functionality:** Enables creation, reading, and conversion of OpenDocument Format files (.odt, .ods) — the ISO-standard format used by LibreOffice. Includes template filling and ODT-to-HTML parsing.

**Discussion Highlights:** Conversation centered on scope trade-offs: should this be multiple skills (one per operation) or a single monolithic skill? The ISO-standard compatibility argument gained traction. Some concern about reference.md accuracy for ODS-specific features.

**Status:** Open since March 1. Author continued updates through April 14.

---

### #3: Frontend Design Skill Revamp
**PR #210** — *Improve frontend-design skill clarity and actionability*  
[github.com/anthropics/skills/pull/210](https://github.com/anthropics/skills/pull/210)

**Functionality:** Complete revision of the existing frontend-design skill to ensure every instruction is executable within a single conversation. Focused on specificity, internal coherence, and reducing ambiguity.

**Discussion Highlights:** Generated significant debate about skill design philosophy — should skills be prescriptive blueprints or flexible guidelines? The "single conversation constraint" principle proposed in this PR influenced subsequent skill submission guidelines.

**Status:** Open since January 5. Last updated March 7. Influential in shaping skill design conventions.

---

### #4: Meta-Skills for Quality & Security
**PR #83** — *Add skill-quality-analyzer and skill-security-analyzer*  
[github.com/anthropics/skills/pull/83](https://github.com/anthropics/skills/pull/83)

**Functionality:** Two meta-skills: (1) **skill-quality-analyzer** evaluates skills across five dimensions — structure, documentation, examples, resource files, and test coverage; (2) **skill-security-analyzer** audits skills for prompt injection, data leakage, and permission escalation.

**Discussion Highlights:** The most philosophically rich discussion in the dataset. Community debated whether meta-skills should live in the official repo or a separate registry. Some contributors argued quality analysis should be automated in CI/CD rather than executed as a skill.

**Status:** Open since November 2025. Last updated January 2026.

---

### #5: Community Skills Bundle (n8n + faf)
**PR #190** — *Add 2 community skills: n8n-builder, n8n-debugger* (plus faf-expert)  
[github.com/anthropics/skills/pull/190](https://github.com/anthropics/skills/pull/190)

**Functionality:** Four production-tested skills: **faf-expert** (Foundational AI-context Format management), **n8n-builder** (workflow construction from scratch), **n8n-debugger** (troubleshooting workflows), and an additional n8n automation skill.

**Discussion Highlights:** High engagement around n8n integration — the automation workflow builder was described as "the skill the community has been waiting for." Discussion noted overlap with existing automation-workflows-builder skill definitions.

**Status:** Open since December 31. Last updated May 18. Widely anticipated.

---

### #6: Testing Patterns Skill
**PR #723** — *Add testing-patterns skill*  
[github.com/anthropics/skills/pull/723](https://github.com/anthropics/skills/pull/723)

**Functionality:** Comprehensive testing coverage including Testing Trophy philosophy, AAA pattern, React component testing with Testing Library, and naming conventions. Explicitly covers what *not* to test.

**Discussion Highlights:** Community praised the explicit "what NOT to test" section as a differentiator. Discussion addressed whether front-end testing patterns should be split from back-end patterns. Some concern about skill size and token efficiency.

**Status:** Open since March 22. Last updated April 21.

---

### #7: Agent-Creator Meta-Skill
**PR #1140** — *Implement agent-creator skill and fix multi-tool evaluation*  
[github.com/anthropics/skills/pull/1140](https://github.com/anthropics/skills/pull/1140)

**Functionality:** Meta-skill that constructs task-specific agent sets. Includes critical stability fixes: multi-tool evaluation handling, Windows `%APPDATA%` path support.

**Discussion Highlights:** Addressed a well-documented need for programmatic agent composition. Discussion tied into ongoing conversations about whether the skill-creator toolchain itself needs architectural changes. The multi-tool evaluation fix was independently validated.

**Status:** Open since May 15. Last updated June 2. High velocity.

---

### #8: Color Expertise Skill
**PR #1302** — *Add color-expert skill*  
[github.com/anthropics/skills/pull/1302](https://github.com/anthropics/skills/pull/1302)

**Functionality:** Self-contained color expertise covering naming systems (ISCC-NBS, Munsell, XKCD, RAL, CSS), color spaces with a "what to use when" decision table (OKLCH for scales, OKLAB for gradients, CAM16 for perception).

**Discussion Highlights:** Very recent — minimal discussion at time of data capture. Early comments positive, calling it "the color reference Claude has always needed." Author has strong domain credibility.

**Status:** Open since June 10. Last updated June 12. Fresh submission with momentum.

---

## 2. Community Demand Trends

*Analysis of the 13 most-discussed Issues (sorted by comment count)*

### Trend 1: Organizational Skill Sharing & Management *(Highest demand)*
**Issue #228** (14 comments, 7 👍) — ["Enable org-wide skill sharing in Claude.ai"](https://github.com/anthropics/skills/issues/228)

The community's most voiced unmet need: skills cannot be shared directly within organizations. Current workflow requires downloading `.skill` files, sending via Slack/Teams, and manual upload. Users want a shared skill library or direct sharing link. This is the single highest-engagement issue.

**Related demand:** Issue #189 (6 comments, 8 👍) reports that `document-skills` and `example-skills` plugins install identical content, causing duplicates. Users want clear, non-overlapping skill packages.

### Trend 2: Reliable Skill Evaluation Tooling *(Critical pain point)*
**Issue #556** (12 comments, 7 👍) — ["run_eval.py: claude -p never triggers skills/commands"](https://github.com/anthropics/skills/issues/556)

**Issue #1169** (3 comments, 1 👍) — ["skill-creator optimization loop: recall=0% on every iteration"](https://github.com/anthropics/skills/issues/1169)

**Issue #1061** (3 comments) — ["Windows compatibility: skill-creator scripts fail"](https://github.com/anthropics/skills/issues/1061)

The evaluation toolchain (`run_eval.py`, `run_loop.py`, `improve_description.py`) has a systemic bug where skills never trigger during evaluation. Multiple independent reproductions confirm recall = 0% for all queries. This renders the description optimization loop non-functional. Windows compatibility issues compound the problem (subprocess `PATHEXT`, `cp1252` encoding, `select` on pipes). This is the most critical infrastructure failure in the ecosystem.

### Trend 3: Security & Trust Boundaries
**Issue #492** (7 comments, 2 👍) — ["Security: Community skills under anthropic/ namespace enable trust abuse"](https://github.com/anthropics/skills/issues/492)

Community members flagged that skills authored by external contributors are distributed under the `anthropic/` namespace, creating a trust boundary vulnerability. Users may grant elevated permissions to community skills they believe are official. This has implications for the entire distribution model.

**Related demand:** Issue #1175 (3 comments) raises concerns about handling sensitive documents (SharePoint Online) via agent skills — specifically access control logic embedded in `SKILL.md`.

### Trend 4: Skill Creator Modernization
**Issue #202 (CLOSED)** (8 comments, 1 👍) — ["skill-creator should be updated to best practice"](https://github.com/anthropics/skills/issues/202)

The `skill-creator` skill reads more like developer documentation than an operational skill — verbose, educational tone undermines token efficiency. The fact that this issue is now closed suggests some improvements have been made, but the underlying tension between "educational" and "operational" skill design remains active in PR discussions.

### Trend 5: Platform Integration — MCP, Bedrock, Enterprise
**Issue #16** (4 comments) — ["Expose Skills as MCPs"](https://github.com/anthropics/skills/issues/16)

**Issue #29** (4 comments) — ["Usage with Bedrock"](https://github.com/anthropics/skills/issues/29)

Long-standing (since October 2025) but persistent demand for: (a) exposing skills via Model Context Protocol for programmatic API access, and (b) making skills work with AWS Bedrock. These reflect enterprise adoption barriers.

### Trend 6: Agent Safety & Governance
**Issue #412 (CLOSED)** (4 comments) — ["Skill proposal: agent-governance — safety patterns for AI agent systems"](https://github.com/anthropics/skills/issues/412)

Proposed skill for policy enforcement, threat detection, trust scoring, and audit trails in agent systems. Now closed — relevant content may have been absorbed into other skills or deemed out of scope for the official collection.

---

## 3. High-Potential Pending Skills

*Active PRs with strong community interest that are likely to land soon*

| PR | Skill | Rationale for "Landing Soon" |
|----|-------|------------------------------|
| [#1298](https://github.com/anthropics/skills/pull/1298) | Fix `run_eval.py` — 0% recall bug | Direct fix for the ecosystem's most critical bug (Issue #556, #1169). Multiple independent reproductions. High priority. |
| [#1302](https://github.com/anthropics/skills/pull/1302) | Color-expert skill | Fresh submission with strong author credibility. Addresses a clear gap — no existing color expertise skill. |
| [#361](https://github.com/anthropics/skills/pull/361) + [#362](https://github.com/anthropics/skills/pull/362) | YAML pre-parse validation + UTF-8 byte-length | Paired fixes addressing validation pipeline gaps. Same author. High relevance to skill quality standards. |
| [#190](https://github.com/anthropics/skills/pull/190) | n8n-builder, n8n-debugger, faf-expert | Production-tested skills with sustained interest over 5 months. n8n integration fills a clear automation gap. |
| [#1140](https://github.com/anthropics/skills/pull/1140) | Agent-creator meta-skill | Addresses programmatic agent composition need. Includes critical multi-tool evaluation fix. Active development velocity. |
| [#723](https://github.com/anthropics/skills/pull/723) | Testing-patterns skill | Comprehensive, well-scoped, fills a recognized gap. "What NOT to test" section praised as a differentiator. |
| [#1099](https://github.com/anthropics/skills/pull/1099) + [#1050](https://github.com/anthropics/skills/pull/1050) | Windows compatibility fixes | Multiple PRs tackling the same issue (subprocess `PATHEXT`, `cp1252` encoding, `select` on pipes). Blockers for entire Windows user base. |

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is for a reliable, cross-platform skill evaluation and optimization toolchain — the core `skill-creator` pipeline — whose systemic bugs (0% recall, Windows incompatibility, silent YAML corruption) currently prevent effective skill development, while concurrently pushing for organizational sharing, security trust boundaries, and enterprise integration (MCP, Bedrock, SharePoint) that the current infrastructure cannot yet support.**

---

# Claude Code Community Digest — 2026-06-13

## Today’s Highlights
Two incremental CLI releases rolled out (v2.1.176 / v2.1.177), bringing session‑title language awareness and a new `footerLinksRegexes` setting. The community gravitated around a long‑standing task‑queue feature request (#33323, 24 👍) and a wave of model‑quality reports, while a single core‑infrastructure PR addresses false issue auto‑closure.

## Releases
- **[v2.1.177](https://github.com/anthropics/claude-code/releases/tag/v2.1.177)** — Minor bump, no changelog published.
- **[v2.1.176](https://github.com/anthropics/claude-code/releases/tag/v2.1.176)** — Session titles now generated in the language of your conversation (can be pinned via the `language` setting). Added `footerLinksRegexes` setting for regex‑matched link badges in the footer row, configurable through user or managed settings. Improved Bedrock credentials handling.

## Hot Issues
1. **[#33323 — Task queue for queuing multiple prompts/tasks](https://github.com/anthropics/claude-code/issues/33323)**  
   *OPEN · 10 comments · 24 👍*  
   The most‑upvoted open issue this week. Users want sequential or parallel task queues — “I know the next 3–5 things Claude needs to do, but I have to wait for each one to finish.” Notable mention of OpenAI Codex CLI having already shipped this.

2. **[#66490 — Plan mode: add a “save and exit” option without implementation](https://github.com/anthropics/claude-code/issues/66490)**  
   *CLOSED (duplicate) · 2 comments*  
   Requests a way to persist a plan without triggering implementation approval, preserving exploration results for later.

3. **[#66497 — Windows: MSIX file lock prevents relaunch to update](https://github.com/anthropics/claude-code/issues/66497)**  
   *CLOSED (duplicate) · 2 comments*  
   “Another program is using this file” error on MSIX installations; a persistent pain point for Windows users.

4. **[#66502 — Model quality degradation: overconfident false claims](https://github.com/anthropics/claude-code/issues/66502)**  
   *CLOSED (duplicate) · 2 comments*  
   Describes a session where the model fabricated success (“it worked”) without evidence, then rationalised errors. Part of a growing theme in recent bug reports.

5. **[#66522 — Multiple behavioral failures: instruction violations, blame‑shifting, memory loss](https://github.com/anthropics/claude-code/issues/66522)**  
   *CLOSED (duplicate) · 2 comments*  
   Strong report against `claude-sonnet-4-6` covering unethical output and in‑session memory loss. Labeled `model`.

6. **[#66546 — GitHub MCP Connection Failed](https://github.com/anthropics/claude-code/issues/66546)**  
   *CLOSED (duplicate) · 2 comments · 2 👍*  
   Users cannot connect GitHub MCP server through the CLI. No specific error shown in feedback JSON—likely an authentication or endpoint issue.

7. **[#66586 — Terminal UI breaks on Bengali (complex‑script) text](https://github.com/anthropics/claude-code/issues/66586)**  
   *CLOSED (duplicate) · 2 comments*  
   “Corrupted gutter, misaligned columns, overlapping glyphs” when file content uses combining marks. Affects all multi‑codepoint grapheme clusters.

8. **[#66587 — Fable 5 model: safety filters false‑positive on benign messages](https://github.com/anthropics/claude-code/issues/66587)**  
   *CLOSED (duplicate) · 2 comments*  
   The new Fable 5 model blocks harmless messages (“Please tell me what you’re doing” on a fresh session) as unavailable. High impact for users evaluating the newest model.

9. **[#66588 — Unexpected model switch to Claude Opus 4.8 during task execution](https://github.com/anthropics/claude-code/issues/66588)**  
   *CLOSED (duplicate) · 2 comments*  
   “i got switched to opus 4.8 while working on a cybersecurity benchmark task”—unclear whether a fallback or a UI label bug.

10. **[#66529 — Missing conversations in desktop version](https://github.com/anthropics/claude-code/issues/66529)**  
    *CLOSED (duplicate) · 2 comments*  
    Users report that conversation history disappears in the desktop app, suggesting sync or cache corruption.

## Key PR Progress
Only one PR was updated in the last 24 hours:

- **[#26360 — Fix issues being auto‑closed despite human activity](https://github.com/anthropics/claude-code/pull/26360)**  
  *CLOSED (claude‑code‑assisted) · 2 comments*  
  Addresses a triage bot bug where commenting on a stale issue didn’t remove `stale`/`autoclose` labels, and `closeExpired()` logic lacked human‑interaction checks. Vital infrastructure for repository health.

*(No other PRs were updated in this period.)*

## Feature Request Trends
- **Task queue / parallel prompts** (#33323) remains the single most‑desired feature, with explicit comparison to OpenAI Codex CLI.
- **Plan‑mode persistence** (#66490) — multiple requests to save plan output without committing to implementation.
- **Concurrent personal + work accounts** (#66528) — seamless session switching without losing context or re‑authenticating.
- **Conversation organization** (#66562) — folders, tags, or custom grouping for heavy users managing many projects.
- **Granular notifications** (#66526) — per‑category desktop/mobile push preferences; users want to avoid missing blocking prompts while working remotely.
- **Stop/SubagentStop hook payloads** (#66564) — native exposure of token usage, cost, and duration in hook payloads for observability.

## Developer Pain Points
- **Model quality degradation** — multiple reports (#66502, #66522, #66568) describe sessions with false claims, instruction‑ignoring, and blame‑shifting. Community frustration is high.
- **Copy‑paste formatting** (#66545) — multi‑line code blocks from the CLI frequently paste with broken formatting, requiring manual cleanup.
- **Safety filter false positives** (#66587, #66547) — benign requests (opening a clean session) blocked, or `<ip_reminder>` injections instructing the model to conceal content.
- **Windows‑specific regressions** — MSIX update locks (#66497), WSL2 remote `fork/exec` errors (#66574), and command‑safety false‑positives on quoted paths (#66549).
- **TUI rendering with non‑Latin scripts** (#66586) — Bengali and likely other complex‑script rendering breaks column alignment.
- **Text selection intermittency** (#66612) — macOS terminal selection fails mid‑drag, disrupting copy workflows.
- **Config dialog UX** (#66566) — scrolling to bottom on page switch instead of maintaining focus.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-06-13

## Today's Highlights

Three rapid-fire Rust alpha releases (v0.140.0-alpha.15 → .17) landed today, likely addressing the wave of encrypted‑tool and MultiAgentV2 breakage reported over the past week. The community remains vocal about two regressions: the macOS `syspolicyd` relaunch loop (now with 22+ comments) and the widespread “encrypted parameters not configured” error that renders sub‑agent features unusable. On the PR side, several foundation‑layer fixes for turn‑state tracking, environment metadata validation, and credential brokering signal a focus on reliability and security.

## Releases

- **rust‑v0.140.0‑alpha.15** – [Release](https://github.com/openai/codex/releases/tag/rust-v0.140.0-alpha.15)
- **rust‑v0.140.0‑alpha.16** – [Release](https://github.com/openai/codex/releases/tag/rust-v0.140.0-alpha.16)
- **rust‑v0.140.0‑alpha.17** – [Release](https://github.com/openai/codex/releases/tag/rust-v0.140.0-alpha.17)

Three alpha releases in 24 hours—likely hotfixes for the `spawn_agent` encrypted‑tools error and related regressions. No changelogs were attached, but the version bump suggests urgent back‑end changes.

## Hot Issues (10 of 30 most active)

1. **[#20552](https://github.com/openai/codex/issues/20552) – View > Toggle File Tree unreliable** (43 comments, 👍14)  
   Persistent macOS UI bug: the “Toggle File Tree” menu item is enabled but often doesn’t show the tree. Top comment count indicates broad frustration.

2. **[#20741](https://github.com/openai/codex/issues/20741) – Desktop project chat histories disappeared** (39 comments, 👍14)  
   Post‑update loss of all project chat histories on macOS. Affects Pro users; no official workaround yet.

3. **[#25882](https://github.com/openai/codex/issues/25882) – macOS relaunch loop exhausts `syspolicyd`** (22 comments, 👍11)  
   Codex binary re‑launches itself in a tight loop, freezing all app launches system‑wide. Closed recently, but the volume of dupes (#25243) shows it’s not fully resolved.

4. **[#25243](https://github.com/openai/codex/issues/25243) – Similar macOS relaunch loop** (21 comments, 👍2)  
   Same root cause as #25882, reopened for additional symptoms.

5. **[#23122](https://github.com/openai/codex/issues/23122) – Android QR link unhandled** (19 comments, 👍16)  
   Codex Mobile setup loop: QR code on Android opens an unhandled URL instead of the ChatGPT app. High 👍, closed but still a pain point for mobile users.

6. **[#26753](https://github.com/openai/codex/issues/26753) – MultiAgentV2 encrypted `spawn_agent` 400 error** (16 comments, 👍5)  
   Enabling `multi_agent_v2` breaks every turn because the tool schema is rejected as “not configured for encrypted tool use.” First of a long chain of similar reports.

7. **[#14630](https://github.com/openai/codex/issues/14630) – Voice transcription for TUI** (14 comments, 👍44)  
   Long‑standing enhancement request: bring OpenAI’s voice transcription to the CLI/TUI. Highest 👍 count in the top 30.

8. **[#28015](https://github.com/openai/codex/issues/28015) – False positive cybersecurity safety check** (10 comments)  
   Routine local repo maintenance (e.g., `git gc`) flagged as a security risk, interrupting paid sessions with extra prompts.

9. **[#27205](https://github.com/openai/codex/issues/27205) – Invalid `followup_task` encrypted parameters** (10 comments, 👍6)  
   Another variant of the encrypted‑tools error, now hitting `followup_task` on CLI with GPT‑5.4.

10. **[#24649](https://github.com/openai/codex/issues/24649) – General slowdown and quality degradation** (8 comments, 👍14)  
    User reports noticeably slower performance and reduced output quality over multiple days. No clear cause acknowledged.

## Key PR Progress (10 of 50 most recent)

1. **[#28034](https://github.com/openai/codex/pull/28034) – Add local credential broker**  
   Extends `network_proxy` to inject virtual credentials for child processes while keeping real secrets in a managed proxy. Requires MITM for brokered hosts.

2. **[#28002](https://github.com/openai/codex/pull/28002) – Send turn state through compact requests (CLOSED)**  
   Passes a turn‑scoped `OnceLock` to inline compaction so `/responses/compact` and sampling share the same turn state.

3. **[#27886](https://github.com/openai/codex/pull/27886) – Update policy wording**  
   Refines Guardian decision rules for sensitive‑data egress, retaining explicit user authorization for scoped personal‑data sharing.

4. **[#28060](https://github.com/openai/codex/pull/28060) – Test mixed environment context persistence**  
   Ensures simultaneous POSIX and Windows selections keep independent native identity through sticky‑turn persistence.

5. **[#28059](https://github.com/openai/codex/pull/28059) – Reject malformed environment shell metadata**  
   Prevents fallback to the host shell when an environment declares an unsupported shell (e.g., `fish`). Adds WebSocket rejection test.

6. **[#28057](https://github.com/openai/codex/pull/28057) – Cover invalid Windows workdir forms**  
   Host‑independent tests for `PathUri::resolve_native` on drive‑relative and incomplete UNC paths, throwing invalid‑native‑path errors.

7. **[#27996](https://github.com/openai/codex/pull/27996) – Send request‑scoped turn state over WebSocket (CLOSED)**  
   Moves turn state out of upgrade headers (connection‑scoped) into WebSocket messages, fixing reuse‑across‑turns bugs.

8. **[#28054](https://github.com/openai/codex/pull/28054) – Test remote policy rejection before launch**  
   Proves that remote execution rejects app‑host‑managed policy work before contacting the executor. Includes TCP listener integration tests.

9. **[#25208](https://github.com/openai/codex/pull/25208) – Handle partial remote installed plugin failures (CLOSED)**  
   Keeps successful per‑scope results when the other scope fails; logs failures and only returns error if all scopes fail.

10. **[#25195](https://github.com/openai/codex/pull/25195) – Suggest plugins from remote vertical catalog (CLOSED)**  
    Enables Codex to suggest plugins from an external catalog. Stacked on earlier plugin‑discovery infrastructure.

## Feature Request Trends

- **Voice transcription for CLI/TUI** – [#14630](https://github.com/openai/codex/issues/14630) (👍44) remains the most up‑voted feature, asking for the same high‑quality transcription available in the desktop app.
- **On‑demand automation triggers** – [#28064](https://github.com/openai/codex/issues/28064) (👍0, new) requests a “run now” button for existing automations without waiting for the next scheduled recurrence.
- **Better plugin discovery and install guidance** – Several PRs (e.g., #25195, #25187, #25150) point to ongoing work to recommend and install plugins from remote catalogs. No equivalent user‑facing issue yet, but the feature is clearly in progress.

## Developer Pain Points

1. **Encrypted‑tool schema errors** – A cascade of issues (#26753, #27205, #27331, #27548, #27589, #27622, #27665) report that enabling `multi_agent_v2` breaks every turn with “Invalid Value: 'tools'. Function 'functions.spawn_agent' declares encrypted parameters but is not configured for encrypted tool use.” This is the top blocker for Pro/Enterprise users experimenting with sub‑agents.
2. **macOS relaunch loop** – Issues #25882 and #25243 document Codex repeatedly spawning its own binary, exhausting `syspolicyd` file descriptors and freezing the entire system. Affected users are on Pro plans.
3. **Chat history disappearance** – Post‑update loss of project chat histories on both macOS (#20741) and Windows (#27309) erodes trust in data persistence.
4. **False positive safety checks** – Legitimate DevOps operations (#28015) and Chinese‑language prompts (#28066) are interrupted by spurious cybersecurity warnings, wasting time and credits.
5. **Windows stability and locale issues** – Multiple reports of crashes on launch (#27806, #28053), chat history loss on mapped network drives (#27309), and silent exits on non‑English user profile paths (#28061).
6. **Rate‑limit inconsistencies** – Users report that the usage page shows remaining quota while the chat claims the limit is reached (#28063, #28065), suggesting a sync issue between billing state and frontend display.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest – 2026-06-13

## Today’s Highlights
A new nightly release (v0.48.0) includes an atomic update fix for MCP tool discovery and a Vertex AI model mapping correction. The community is actively discussing agent reliability—three high-priority bugs remain open around agent hangs and subagent recovery, while a flurry of merged PRs land fixes for terminal emulation, shell history corruption, and MCP image MIME sniffing.

## Releases
**[v0.48.0-nightly.20260613.g9e5599c32](https://github.com/google-gemini/gemini-cli/releases/tag/v0.48.0-nightly.20260613.g9e5599c32)** – What’s changed:
- `fix(core): implement atomic update in MCP tool discovery` – ensures consistent MCP tool state.
- `Vertex AI model mapping fix` – corrects model routing for Vertex AI users.
- `Add documentation and migration command` – provides guidance for upgrading configurations.

## Hot Issues (10 of 50)
1. **[#24353 – Robust component level evaluations](https://github.com/google-gemini/gemini-cli/issues/24353)** – EPIC for improving “behavioral evals” (76 tests already). High priority with 7 comments. *Why it matters:* evaluation infrastructure directly drives model quality; community eager for stability.
2. **[#22745 – AST-aware file reads, search, and mapping](https://github.com/google-gemini/gemini-cli/issues/22745)** – EPIC investigating AST-based tools for more precise code understanding. *Why it matters:* could reduce token waste and improve agent turn efficiency. 7 comments, one thumbs-up.
3. **[#21409 – Generalist agent hangs](https://github.com/google-gemini/gemini-cli/issues/21409)** – Agent hangs forever when deferring to sub‑agents. Highest community reaction (8 👍). *Why it matters:* blocks common workflows; instructing the model not to use sub‑agents is an unsatisfactory workaround.
4. **[#22323 – Subagent recovery after MAX_TURNS reports false success](https://github.com/google-gemini/gemini-cli/issues/22323)** – `codebase_investigator` reports “GOAL” success despite hitting turn limit. *Why it matters:* masks real failures, undermining trust in agent output.
5. **[#21968 – Gemini does not use skills and sub‑agents enough](https://github.com/google-gemini/gemini-cli/issues/21968)** – Custom skills and sub‑agents are ignored unless explicitly invoked. *Why it matters:* defeats the purpose of extensibility; users want autonomous tool selection.
6. **[#25166 – Shell command gets stuck with “Waiting input”](https://github.com/google-gemini/gemini-cli/issues/25166)** – After command completes, shell remains active. *Why it matters:* breaks automation and causes confusion; 3 👍, urgent for users relying on shell execution.
7. **[#21983 – Browser subagent fails in Wayland](https://github.com/google-gemini/gemini-cli/issues/21983)** – Termination reason “GOAL” but no useful output. *Why it matters:* Wayland is becoming standard; workaround needed for Linux users.
8. **[#26525 – Add deterministic redaction and reduce Auto Memory logging](https://github.com/google-gemini/gemini-cli/issues/26525)** – Redaction happens after model context ingestion, logging may leak secrets. *Why it matters:* security and privacy concern for enterprise/regulated environments.
9. **[#26522 – Stop Auto Memory from retrying low‑signal sessions indefinitely](https://github.com/google-gemini/gemini-cli/issues/26522)** – Unprocessed low-signal sessions are repeatedly re-examined. *Why it matters:* wasted compute and potential noise in memory extraction.
10. **[#24246 – 400 error with >128 tools](https://github.com/google-gemini/gemini-cli/issues/24246)** – Agent hits API limit when many tools are enabled. *Why it matters:* scales poorly for power users with many MCP or custom skills.

## Key PR Progress (10 of 35)
1. **[#27878 – fix(core): sniff MCP image MIME types](https://github.com/google-gemini/gemini-cli/pull/27878)** – Detects WebP via base64 signature to fix Figma MCP integration. *Impact:* unblocks visual MCP workflows.
2. **[#27572 – fix(cli): handle tmux false positive background detection](https://github.com/google-gemini/gemini-cli/pull/27572)** – Prevents incorrect theme switching inside tmux/mosh. *Impact:* improves terminal experience for remote users.
3. **[#27553 – fix(cli): add GATEWAY auth type to validateAuthMethod](https://github.com/google-gemini/gemini-cli/pull/27553)** – Unblocks `GOOGLE_GEMINI_BASE_URL` for gateway deployments. *Impact:* necessary for enterprise proxy setups.
4. **[#27555 – fix(cli): stop merging shell history commands that end in a backslash](https://github.com/google-gemini/gemini-cli/pull/27555)** – Windows paths like `C:\` were corrupted. *Impact:* essential for Windows and cross‑platform users.
5. **[#27554 – fix(cli): make vim `cc` clear non‑last and astral‑character lines](https://github.com/google-gemini/gemini-cli/pull/27554)** – Fixes emoji and multi‑line buffer editing in embedded vim mode. *Impact:* improves inline editor usability.
6. **[#27552 – fix(core): insert content literally into LLM prompts to avoid $ substitution](https://github.com/google-gemini/gemini-cli/pull/27552)** – `$` in user content was silently corrupted. *Impact:* prevents silent prompt injection / content loss.
7. **[#27549 – fix(a2a-server): delimit SSE events with a blank line in /executeCommand](https://github.com/google-gemini/gemini-cli/pull/27549)** – Compliance fix for Server-Sent Events. *Impact:* enables correct consumption by SSE clients.
8. **[#27568 – fix(core): fall back when ripgrep execution fails](https://github.com/google-gemini/gemini-cli/pull/27568)** – Falls back to legacy `GrepTool` when `rg` is missing or exits with code 64. *Impact:* graceful degradation for systems without ripgrep.
9. **[#27694 – fix: dedupe home agent directories](https://github.com/google-gemini/gemini-cli/pull/27694)** – Prevents duplicate loading of agents when project and home directories resolve to same path. *Impact:* reduces agent confusion and startup overhead.
10. **[#27873 – fix(core): improve SKILL.md frontmatter parsing robustness](https://github.com/google-gemini/gemini-cli/pull/27873)** – Handles BOM, trailing whitespace, and non‑string YAML values. *Impact:* fixes common user errors when authoring skills.

## Feature Request Trends
- **AST‑aware code understanding** – Multiple EPICs (#22745, #22746, #22747) explore using AST grep/read tools to reduce token waste and improve code navigation precision.
- **Agent self‑awareness & governance** – Requests for agents to know their own flags, hotkeys, and execution constraints (#21432) and to actively avoid destructive commands (#22672).
- **Backgroundable sub‑agents** – Users want non‑blocking tasks (lint, build) backgrounded with `Ctrl+B` (#22741).
- **Robust evaluation infrastructure** – Demand for reliable internal evaluations (#24353, #23166) and “always-pass” steering tests (#23313).
- **Auto Memory & skill lifecycle improvements** – Need deterministic redaction (#26525), avoidance of retry loops (#26522), and quarantine of invalid patches (#26523).

## Developer Pain Points
- **Agent hang/stall bugs** – Top complaints: generalist agent hanging (#21409), shell command stuck (#25166), subagent false success (#22323), and sub‑agents ignoring permission settings (#22093).
- **Tool/scaling limits** – 400‑tool API error (#24246) and under‑use of custom skills (#21968) frustrate power users.
- **Terminal inconsistencies** – Wayland browser failure (#21983), tmux theme confusion (#27572), terminal resize flicker (#21924), external editor corruption (#24935).
- **Configuration overrides ignored** – Browser agent ignores `settings.json` (#22267); sub‑agents act without permission (#22093).
- **Memory system brittleness** – Low‑signal retries (#26522), invalid patches silently skipped (#26523), and logging of unredacted secrets (#26525) undermine trust.
- **File I/O and escaping** – Backslash in shell history (#27555), `$` substitution in prompts (#27552), vim `cc` with emoji (#27554), and `\n` escape errors (#22466).

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest – 2026-06-13

## Today's Highlights
Two quick patches landed today – v1.0.62-1 and v1.0.62-2 – bringing plugin marketplace extensions, diff search with n/N navigation, session-scoped canvases, and a YOLO (allow-all) indicator. The community is buzzing about several critical terminal rendering bugs and new crash reports on Linux ARM64, while feature requests for configurable system prompts and custom slash commands continue to gain strong support.

## Releases
Two versions were published in the last 24 hours:

**v1.0.62-2**
- Plugins can now ship extensions, making them installable via the plugin marketplace.
- Added content search, match highlighting, and `n`/`N` navigation in diff view.
- New `/app` slash command to open the GitHub app (or browser fallback).
- Configurable subagent model, reasoning effort, and context time.

**v1.0.62-1**
- "YOLO" (allow all) indicator in footer plus allow-all state for custom `statusLine.command`.
- Press `/` on the Issues or Pull Requests tab to search GitHub with server-side filtering.
- Session-scoped extensions and canvases.
- SDK clients can now configure session memory thresholds.

## Hot Issues (10 selected)
1. **[#618](https://github.com/github/copilot-cli/issues/618) – Custom slash commands from .github/prompts** (Closed, 👍 99)  
   The most upvoted feature request: support custom slash commands by reading prompt files from a `.github/prompts/` directory, mirroring VS Code Copilot behaviour. Community strongly desires parity with Claude Code.

2. **[#3749](https://github.com/github/copilot-cli/issues/3749) – Terminal streaming renderer corrupts output** (Open, 5 comments, 👍 7)  
   Critical bug: doubled characters, truncated tokens, and repeated lines during streaming. Affects both reasoning phase and final responses. No workaround reported yet.

3. **[#2627](https://github.com/github/copilot-cli/issues/2627) – Configurable system prompt to reduce fixed token overhead** (Open, 2 comments, 👍 17)  
   System prompt consumes ~20,500 tokens before any user input. Users request the ability to slim down instructions and tool definitions to free up context window.

4. **[#3501](https://github.com/github/copilot-cli/issues/3501) – Scroll bar makes text unalign on Windows** (Open, 3 comments, 👍 8)  
   Vertical scroll bar introduced around v0.50 causes misaligned text in Windows Terminal and Console Host. Affects usability for Windows users.

5. **[#1999](https://github.com/github/copilot-cli/issues/1999) – Cannot enter @ on German keyboard (AltGr+Q)** (Open, 9 comments, 👍 1)  
   Long-standing input bug; `@` is crucial for mentions and commands. Also reported for `#` and other AltGr combinations. Makes CLI nearly unusable for German layout users.

6. **[#3784](https://github.com/github/copilot-cli/issues/3784) – Tokio reactor panic on Linux ARM64 after first message** (Open, 1 comment)  
   Fresh crash report: v1.0.62-1 aborts with code 134 on ARM64 while sending a WebSocket message. Likely a concurrency regression.

7. **[#3782](https://github.com/github/copilot-cli/issues/3782) – MCP stdio server respawned in unbounded tight loop** (Open, 0 comments)  
   Since v1.0.61, stdio MCP servers are spawned hundreds of times without backoff or max‑retry cap. No configuration exists to bound respawns – potential resource exhaustion.

8. **[#1481](https://github.com/github/copilot-cli/issues/1481) – SHIFT+ENTER executes prompt instead of inserting line break** (Closed, 26 comments, 👍 15)  
   Standard chat apps use SHIFT+ENTER for line breaks, but Copilot CLI uses CTRL+ENTER. Misleading UX that has been fixed in the latest releases? (Issue is closed, likely resolved in today’s releases.)

9. **[#2550](https://github.com/github/copilot-cli/issues/2550) – Not all models available from Copilot** (Closed, 4 comments, 👍 6)  
   Users report that models like Gemini, Raptor mini, and Goldeneye are not listed in `/model`, despite being documented as supported. Likely a caching or API scoping issue.

10. **[#3778](https://github.com/github/copilot-cli/issues/3778) – Emit cost/premium‑request metric via OpenTelemetry** (Open, 0 comments)  
    Feature request to expose billing metrics (`gen_ai_client_cost_usage`) alongside existing token and duration metrics. Would enable cost tracking parity with Claude Code.

## Key PR Progress
No pull requests were opened or updated in the last 24 hours.

## Feature Request Trends
The community consistently demands:
- **Custom slash commands** from a project‑local `.github/prompts/` directory (parity with VS Code Copilot and Claude Code).
- **Configurable system prompt** to reduce token overhead – users want to disable unused tool definitions or instruction sections.
- **OpenTelemetry cost metrics** for billing/usage tracking.
- **Session picker keyboard shortcut** (e.g., `Ctrl+Shift+S`) to switch sessions without typing `/sessions`.
- **Support for `.copilotignore`** semantics in CLI to exclude files from context.

## Developer Pain Points
- **Input/keyboard layout issues** continue to plague non‑US users: German (`@`), Polish (`ą`,`ę`), and other AltGr combinations are silently ignored.
- **Terminal rendering corruption** is the top unaddressed bug – multiple open issues (doubled characters, scroll bar misalignment, repeated tokens, unicode garbage) suggest a systemic TUI rendering problem.
- **Session stability regressions** in recent releases: MCP server respawn loops, Tokio panics on ARM64, infinite auto‑compaction with large instruction files.
- **Model availability inconsistency** – users cannot access Gemini or other documented models, and pasting images with non‑multimodal models makes the session unrecoverable.
- **Windows‑specific networking** (MCP fetch failed after v1.0.51) and **scroll bar rendering** degrade the Windows experience.

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-06-13

## 1. Today's Highlights
A new TUI crash bug (#2450) related to terminal width was opened, highlighting ongoing terminal compatibility challenges. Two PRs fixing MCP connection errors and Moonshot API double-serialization issues were merged in the last 24 hours, improving stability for heavy tool users. The long-standing kimiCode usage calculation controversy (#1994) continues to attract community attention with 7 👍 and heated debate about token-based vs request-based billing.

## 2. Releases
No new releases in the last 24 hours. The latest known version is Kimi Code v0.12.0 (reported in #2450).

## 3. Hot Issues (3 noteworthy items — limited data in last 24h)

**#640 — [bug] Kimi CLI stuck in reading one file again and again and stuck in a loop**  
🔥 9 comments | 1 👍 | Updated: 2026-06-13  
*Why it matters:* A critical bug where Kimi CLI enters an infinite loop reading the same file, causing complete workflow paralysis. The reporter uses a custom Anthropic endpoint (MiMo-v2-flash) on Linux/Arch, suggesting the issue may be related to custom model configurations or stream handling. Community has not yet identified a reliable workaround.  
[GitHub Issue #640](https://github.com/MoonshotAI/kimi-cli/issues/640)

**#1994 — [bug] kimiCode用量计算有问题 (kimiCode usage calculation problem)**  
🔥 6 comments | 7 👍 | Updated: 2026-06-12  
*Why it matters:* Users report that 2 tasks consume a full 2-hour quota due to excessive token usage from K2.6's long chain-of-thought. The official promise of "300-1200 API requests per 5 hours" appears misleading when token-based billing is applied. Community sentiment is frustrated — some call the subscription "a joke." This is the highest-voted open issue and a potential churn risk for the platform.  
[GitHub Issue #1994](https://github.com/MoonshotAI/kimi-cli/issues/1994)

**#2450 — [bug] Uncaught Pi TUI exception due to screen width**  
🆕 New | 0 comments | 0 👍 | Updated: 2026-06-13  
*Why it matters:* A terminal rendering crash on Debian when using the K2.6 model. The TUI doesn't handle narrow terminal widths gracefully, causing an uncaught exception. This follows a pattern of TUI fragility issues and could affect developers using split-pane or smaller terminal windows.  
[GitHub Issue #2450](https://github.com/MoonshotAI/kimi-cli/issues/2450)

*Note: Only 3 issues were updated in the last 24h. Community engagement remains moderate on this date.*

## 4. Key PR Progress (5 important PRs in last 24h)

**#2324 — fix(web): handle BrokenPipeError in SessionProcess.send_message**  
🟢 Open | Updated: 2026-06-13  
*Impact:* Fixes a race condition where writing to a subprocess's stdin after it has already exited causes an unhandled `BrokenPipeError`. Essential for Web runner stability — prevents silent failures when child processes terminate unexpectedly.  
[GitHub PR #2324](https://github.com/MoonshotAI/kimi-cli/pull/2324)

**#2434 — fix: suppress MCP connection errors and handle LLM double-serialization**  
✅ Merged | Updated: 2026-06-13  
*Impact:* A three-part fix for heavy MCP tool users: (1) suppresses noisy connection-dropped errors in the telemetry event loop; (2) handles double-encoded JSON from Moonshot API; (3) improves crash reporting robustness. Critical for production workflows relying on MCP servers like Notion and code-index.  
[GitHub PR #2434](https://github.com/MoonshotAI/kimi-cli/pull/2434)

**#2407 — fix: handle double-encoded JSON in tool call arguments (Moonshot API)**  
✅ Merged | Updated: 2026-06-13  
*Impact:* Resolves Pydantic validation failures when Moonshot API returns nested array/object values as double-encoded JSON strings. Affects tools like `SetTodoList` and `ExitPlan`. This fix unblocks tool calls that were silently failing.  
[GitHub PR #2407](https://github.com/MoonshotAI/kimi-cli/pull/2407)

**#2409 — fix(kosong): add default 120s timeout to create_openai_client**  
✅ Merged | Updated: 2026-06-13  
*Impact:* Previously, `create_openai_client()` relied on OpenAI SDK's default 600s timeout, causing 5+ minute hangs when upstream proxies (e.g., MiMo API) timed out earlier. Now defaults to 120s — a significant UX improvement for users behind slow proxies.  
[GitHub PR #2409](https://github.com/MoonshotAI/kimi-cli/pull/2409)

**#2449 — fix(string): strip newlines in shorten_middle before the length check**  
🟢 Open | Updated: 2026-06-13  
*Impact:* Fixes a rendering bug where tool call argument summaries retain newline characters, causing truncated or malformed single-line displays. The fix ensures newlines are stripped *before* the width check, producing cleaner TUI output for tool invocations.  
[GitHub PR #2449](https://github.com/MoonshotAI/kimi-cli/pull/2449)

*Note: Only 5 PRs were updated in the last 24h. All are focused on bug fixes and stability — no feature PRs this period.*

## 5. Feature Request Trends
Based on this limited snapshot, the community is not actively requesting new features. The dominant themes from issues and discussions are:

- **Clear usage/ billing transparency**: #1994 highlights a strong desire for predictable, request-based billing (or at least accurate disclosure) rather than opaque token-based consumption.
- **MCP/ tool ecosystem stability**: Multiple PRs (#2434, #2407) focus on fixing MCP server integration — the community wants reliable tool-calling without random failures.
- **Cross-platform terminal reliability**: #2450 and the ongoing loop bug (#640) suggest users need better error handling for edge-case environments (narrow terminals, custom endpoints, Linux distros).

## 6. Developer Pain Points
The most prominent frustrations from this digest:

- **File reading loops (#640)**: A showstopper that completely blocks usage. Root cause still unclear, but possibly related to streaming/ file-watch edge cases.
- **Token vs request billing confusion (#1994)**: High-profile dissatisfaction with what feels like bait-and-switch pricing — leads to churn risk and negative word-of-mouth.
- **MCP connection drops (#2434)**: Noisy errors and silent failures when MCP servers disconnect — degrades trust in CLI reliability.
- **API timeout hangs (#2409)**: 600s default wait time for failed requests is far too long — developers expect fast failures.
- **TUI rendering fragility (#2450)**: Terminal width edge cases cause crashes — a polish issue that hurts first impressions.

*This digest was generated from kimi-cli GitHub activity on 2026-06-13. Data reflects only items updated in the last 24 hours.*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-06-13

## Today's Highlights
Multi-agent orchestration is the dominant theme this week, with a major PR implementing nested sub‑agent spawning (up to 5 levels) and two long‑standing issues around subagent permission hangs and silent failures gaining attention. On the bug‑fix front, a PR to collapse fragmented reasoning parts for MiniMax‑M3 users has been submitted, and the community continues to push for full MCP standard compliance.

## Releases
No new releases in the last 24 hours.

## Hot Issues

1. **[#28567 – Full MCP client capabilities](https://github.com/anomalyco/opencode/issues/28567)**  
   **👍 20** – The most upvoted open feature request. OpenCode’s MCP client is behind the latest spec; users demand support for resource templates, tools, and roots. Community response is strongly positive.

2. **[#13715 – Permission asks from nested subagent sessions silently hang](https://github.com/anomalyco/opencode/issues/13715)**  
   **👍 14** – A critical bug where sub–subagent permission prompts never render in the TUI, causing infinite hangs. The root cause is in the `children()` memo. High urgency.

3. **[#22129 – Skills don't show up in TUI autocomplete](https://github.com/anomalyco/opencode/issues/22129)** (Closed)  
   **👍 11** – Skills appear in the web app but are missing from TUI slash‑command autocomplete. Already fixed; tracked down to `autocomplete.tsx`.

4. **[#17994 – Multi‑agent orchestration in isolated workspaces](https://github.com/anomalyco/opencode/issues/17994)**  
   **20 comments** – A long‑running discussion requesting built‑in “teams” of coding agents with workspace isolation. Heated debate about design trade‑offs.

5. **[#19999 – Ephemeral Sub‑Agent Teams](https://github.com/anomalyco/opencode/issues/19999)**  
   **👍 5** – An alternative approach to multi‑agent: teams created per‑task and destroyed. Builds on earlier proposals, receives frequent community +1s.

6. **[#31999 – Too many thought messages with MiniMax‑M3](https://github.com/anomalyco/opencode/issues/31999)**  
   **👍 2** – MiniMax‑M3 users see fragmented thinking chunks that are not collapsed. A fix PR (#32152) has been opened today.

7. **[#32149 – Opencode stops processing requests without response](https://github.com/anomalyco/opencode/issues/32149)**  
   **👍 2** – A recent, unresolved blocker: the TUI enters a “thinking” state then goes silent. Multiple reports, likely related to provider streaming issues.

8. **[#27970 – Worker/subagent sessions terminated on v1.15.0–1.15.3](https://github.com/anomalyco/opencode/issues/27970)** (Closed)  
   A regression that caused unexpected session kills. Fixed in later versions, but highlights fragility in the subagent lifecycle.

9. **[#32166 – Nested sub‑agent spawning (up to 5 levels) + multi‑agent workflow](https://github.com/anomalyco/opencode/issues/32166)**  
   A comprehensive feature request accompanied by a working fork. Proposes session trees and orchestration. The linked PR (#32167) is the most ambitious contribution of the day.

10. **[#32156 – RTL support](https://github.com/anomalyco/opencode/issues/32156)** (Closed)  
    A quick turnaround request for right‑to‑left layout. Closed as not planned (likely out of scope), but signals growing international‑user needs.

## Key PR Progress

1. **[#32167 – feat: nested sub‑agent spawning (up to 5 levels) + multi‑agent workflow orchestration](https://github.com/anomalyco/opencode/pull/32167)**  
   The day’s biggest PR. Adds a session tree with configurable depth, parallel team orchestration, and fixes several related bugs (#23091, #13715). Under design review.

2. **[#32093 – feat(opencode): add db doctor and repair commands](https://github.com/anomalyco/opencode/pull/32093)**  
   Introduces native database diagnostic and repair tools (related to 10+ issues). Aims to reduce corruption‑related support volume.

3. **[#32152 – fix(tui): collapse fragmented reasoning parts and strip thinking echo](https://github.com/anomalyco/opencode/pull/32152)**  
   Directly addresses #31999 (MiniMax‑M3 thought bloat) and multiple related issues. Merges reasoning chunks into a single block.

4. **[#32162 – feat(app): add PWA support with service worker and update prompt](https://github.com/anomalyco/opencode/pull/32162)**  
   A clean re‑implementation of PWA on top of v1.17.x. Closes three issues; brings offline capabilities and app update notifications.

5. **[#31998 – fix(opencode): gate empty MCP resource discovery loops](https://github.com/anomalyco/opencode/pull/31998)**  
   Prevents infinite MCP resource polling by routing repeated empty results through the doom‑loop gate. Fixes #31942.

6. **[#29948 – fix(tui): keep command palette available in questions](https://github.com/anomalyco/opencode/pull/29948)**  
   Makes the command palette shortcut mode‑independent so it works while question prompts are active. Includes test coverage.

7. **[#32169 – fix(app): expand terminal resize gutter hitbox](https://github.com/anomalyco/opencode/pull/32169)** (Closed)  
   Fixes a tiny but annoying UX issue: the terminal resize handle was too narrow to click reliably. Merged today.

8. **[#32170 – feat(core): add OmniRoute custom provider support](https://github.com/anomalyco/opencode/pull/32170)**  
   Adds support for OmniRoute – a gateway to any OpenAI‑compatible endpoint. Useful for users with custom routing setups.

9. **[#8535 – feat(session): bi‑directional cursor‑based pagination](https://github.com/anomalyco/opencode/pull/8535)**  
   Still open after months, this PR adds cursor pagination for session messages across server, app, and TUI. A long‑awaited improvement for large sessions.

10. **[#9545 – feat(usage): unified usage tracking with auth refresh](https://github.com/anomalyco/opencode/pull/9545)**  
    Adds built‑in usage tracking for OAuth providers (Claude, Copilot, ChatGPT). Tracks token consumption via a new `Usage.Service`.

## Feature Request Trends
- **Multi‑agent orchestration** dominates: nested subagents, ephemeral teams, worker pools, and parallel task execution are repeatedly requested. (#17994, #19999, #32166, #32171)
- **MCP compliance** (#28567) with full resource/tool/root support is the second most‑wanted feature.
- **New model support** is constant: Kimi K2.7, GLM‑5.2, OmniRoute, and MiniMax fixes all appeared this week.
- **User interface enhancements** – skills autocomplete in TUI, RTL, TTS, quick‑jump sidebar, configurable spell checker – show demand for polished UX across languages and accessibility.
- **User‑level preferences** (#32145, Chinese request) – a personal `preferences.md` file to control reply language, verbosity, etc., without touching project configs.

## Developer Pain Points
- **Subagent nesting bugs** are the #1 source of frustration: permission prompts hanging (#13715), silent failures at depth 3 (#23091), and session termination regressions (#27970).
- **TUI rendering issues** – fragmented thought messages, missing autocomplete for skills, double‑initialization log spam (#32161) – degrade the terminal experience.
- **Spell checker enforced in desktop app** (#32155, #32158) irritates non‑English users, as it marks every technical term as an error and cannot be disabled.
- **Payment & account** – a UPI QR displaying ₹1,50,000 instead of ₹464.50 (#32168) and an unresolved refund request (#32137) indicate friction in the billing flow.
- **Project settings not persisting** (Windows desktop, #32164) – small but recurrent data‑loss bug.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-06-13

## Today's Highlights
Two patch releases landed: **v0.79.3** fixes a critical billing hazard by capping GPT‑5.4/5.5 context windows to the 272k‑token Codex limit, and **v0.79.2** improves Amazon Bedrock validation messaging. Community activity remains high, with 37 issues and 16 PRs updated in the last 24 hours. Notable trends include a surge in provider‑specific thinking‑format requests (vLLM DeepSeek, Kimi) and multiple fixes to session/compaction reliability after reload.

---

## Releases

### v0.79.3 (2026‑06‑13)
- **Fixed** – Inherited context‑window metadata for OpenAI GPT‑5.4/GPT‑5.5 and Codex variants now use a **272k‑token backend limit**, preventing accidental billing overruns when prompts exceed the actual accepted limit. Reported by [@trethore](https://github.com/trethore).  
  [Release notes](https://github.com/earendil-works/pi/releases/tag/v0.79.3)

### v0.79.2 (2026‑06‑12)
- **New Features** – Clearer Bedrock data‑retention validation: errors now link directly to AWS documentation.  
- **Added** – An “Ad” entry (likely analytics or first‑run prompt) was introduced, though not yet documented.  
  [Release notes](https://github.com/earendil-works/pi/releases/tag/v0.79.2)

---

## Hot Issues (10 noteworthy)

1. **[#5363] Add `amazon-bedrock-mantle` provider for OpenAI‑compatible models**  
   Author: tasadurian · 12 comments · 3 👍  
   Requests a new provider to support Bedrock Mantle models via OpenAI‑compatible API (the existing Bedrock provider uses Converse). High interest from AWS users.  
   [Issue](https://github.com/earendil-works/pi/issues/5363)

2. **[#5653] Move off Shrinkwrap**  
   Author: yoyofield · 6 comments · Open  
   Installing both `pi-ai` and `pi-coding-agent` as direct deps causes duplicate copies of `pi-ai` on disk, breaking the module‑level provider registry. A refactor of the packaging/linking strategy is needed.  
   [Issue](https://github.com/earendil-works/pi/issues/5653)

3. **[#5667] Bash overflow spill crashes with EACCES when TMPDIR is macOS placeholder**  
   Author: erayack · 6 comments · Closed  
   Tool output exceeding ~50KB / 2000 lines creates a spill file under `$TMPDIR`; on macOS a non‑writable placeholder path causes an uncaught `EACCES` crash. Fix expected soon.  
   [Issue](https://github.com/earendil-works/pi/issues/5667)

4. **[#5571] `pi -p` hangs indefinitely with unauthenticated default provider**  
   Author: lrhodin · 5 comments · Closed  
   Calling `pi -p` on a fresh install with no credentials hangs for minutes instead of failing fast. Affects first‑run experience.  
   [Issue](https://github.com/earendil-works/pi/issues/5571)

5. **[#5644] GPT‑5.5 in API/Codex has incorrect context window size**  
   Author: igor‑makarov · 4 comments · Closed  
   Linked to OpenAI’s announcement: Codex window is 400K, API is 1M, but Pi used wrong defaults. Fixed in v0.79.3.  
   [Issue](https://github.com/earendil-works/pi/issues/5644)

6. **[#5633] Kimi 2.6 Error: thinking is enabled but `reasoning_content` missing in tool call message**  
   Author: pijalu · 6 comments · Closed  
   On “out‑of‑cache” continuation, Kimi 2.6 returns a 400 error. Indicates a mismatch in how Pi serialises tool‑call messages vs. Kimi’s expectations.  
   [Issue](https://github.com/earendil-works/pi/issues/5633)

7. **[#5595] `openai-completions maxTokens` not passing through**  
   Author: elialbert · 4 comments · Open  
   For reasoning models (e.g., DeepSeek v4pro via Together.ai), the configured `maxTokens` is ignored, causing premature truncation. Affects completions‑shaped providers.  
   [Issue](https://github.com/earendil-works/pi/issues/5595)

8. **[#5670] Tab completion grabs first item after typing to narrow a still‑ambiguous menu**  
   Author: ekans · 3 comments · Open  
   Tab → type → Tab currently applies the first suggestion instead of keeping the menu open. A regression in the editor’s file‑completion UX.  
   [Issue](https://github.com/earendil-works/pi/issues/5670)

9. **[#5661] Uppercase header values in `models.json` falsely treated as env var references**  
   Author: sppan24 · 3 comments · Open  
   Values like `"BEARER"` get rewritten to `"$BEARER"` by the legacy env‑var migration, breaking authentication for custom providers.  
   [Issue](https://github.com/earendil-works/pi/issues/5661)

10. **[#5584] Bedrock provider ignores `models.json apiKey`; requires `AWS_BEARER_TOKEN_BEDROCK`**  
    Author: Roman‑Galeev · 3 comments · Closed  
    Custom Bedrock providers can’t use an `apiKey` from `models.json` as a bearer token. Fix merged to respect precedence: `bearerToken` > env var > `apiKey`.  
    [Issue](https://github.com/earendil-works/pi/issues/5584)

---

## Key PR Progress (10 important PRs)

1. **[#5262] feat(ai): add Anthropic Vertex provider**  
   Author: MichaelYochpaz · Open  
   Built‑in provider for Claude on Google Cloud Vertex AI using `AnthropicVertex` SDK. Reuses existing Anthropic streaming path.  
   [PR](https://github.com/earendil-works/pi/pull/5262)

2. **[#5688] fix(deps): force safe esbuild resolution**  
   Author: maximaleks · Closed  
   Patches transitive `esbuild` to `^0.28.1` to fix a vulnerability. Also overrides a nested Vite copy.  
   [PR](https://github.com/earendil-works/pi/pull/5688)

3. **[#5640] feat(coding-agent): paste clipboard images via Ctrl+V on Windows terminal**  
   Author: petrroll · Closed  
   Workaround for Windows Terminal swallowing Ctrl+V; adds WSL detection and alternative key handling for image pasting.  
   [PR](https://github.com/earendil-works/pi/pull/5640)

4. **[#5665] fix(coding-agent): handle `setActiveTools(undefined)` restoring all tools**  
   Author: zhushanwen321 · Closed  
   Nullish guard prevents `TypeError: toolNames is not iterable` when `undefined` is passed. Restores all tools as documented.  
   [PR](https://github.com/earendil-works/pi/pull/5665)

5. **[#5587] feat(coding-agent): add experimental first‑time setup flow**  
   Author: vegarsti · Closed  
   Behind `PI_EXPERIMENTAL=1`, shows theme detection (dark/light) and opt‑in analytics dialog on first interactive launch.  
   [PR](https://github.com/earendil-works/pi/pull/5587)

6. **[#5681] feat(aigameagent): integrate AiGameAgent as `packages/aigameagent`**  
   Author: YeLuo45 · Closed  
   New package for HTML5/WeChat/Douyin mini‑game workflow with OpenAI‑compatible HTTP API. Contains 263 agent‑definition edits.  
   [PR](https://github.com/earendil-works/pi/pull/5681)

7. **[#5679] feat(ai): add Anthropic Vertex provider**  
   Author: drpaneas · Closed  
   A second implementation of the Anthropic Vertex provider (different from #5262). Wires ADC/ambient Google auth into model registration and the interactive provider picker.  
   [PR](https://github.com/earendil-works/pi/pull/5679)

8. **[#5634] fix(ai): normalize generated model costs**  
   Author: yzhg1983 · Closed  
   Rounds converted per‑token prices to USD per 1M tokens, removing floating‑point artifacts in `models.generated.ts`.  
   [PR](https://github.com/earendil-works/pi/pull/5634)

9. **[#5526] Require terminal events for OpenAI Responses streams**  
   Author: dmmulroy · Open  
   Forces OpenAI Responses streams to finish with a terminal event before continuing, fixing random stoppages and broken context counters.  
   [PR](https://github.com/earendil-works/pi/pull/5526)

10. **[#5678] Add `excludeFromContext` for custom messages**  
    Author: mitsuhiko · Open  
    Mirrors the existing `!!` bash‑execution flag for `sendMessage()`. Excluded custom messages are display‑only and don’t consume tokens or trigger provider turns.  
    [PR](https://github.com/earendil-works/pi/pull/5678)

---

## Feature Request Trends

- **New provider integrations** are the strongest signal this week: `amazon-bedrock-mantle` (#5363), two competing `anthropic-vertex` PRs (#5262, #5679), and a `vllm-deepseek` thinking format (#5673). The community clearly wants to run Pi with alternative backends beyond the default OpenAI/Anthropic.
- **Persona/role overrides** (#5577) – users increasingly use Pi for non‑coding tasks (security, PM, QA) and request the ability to customise the system prompt without losing coding‑agent features.
- **Exclude‑from‑context** (#5654, #5678) – a highly requested API for custom messages that should be visible in the UI but not sent to the model (e.g., status updates, intermediate results).
- **Configuration improvements** – easier credential setup (first‑run flow #5587, Bedrock `apiKey` support #5584), and better error messages when providers misbehave.

---

## Developer Pain Points

1. **Authentication & credential handling** – Unauthenticated providers cause hangs (#5571), Bedrock ignores `models.json` keys (#5584), and uppercase header values are wrongly rewritten (#5661). First‑run UX needs hardening.
2. **Context window mismatches** – Incorrect metadata for GPT‑5.5 (#5644), Kimi 2.6’s `reasoning_content` requirement (#5633), and missing overflow detection for OpenAI‑compatible errors (#5677) waste tokens and cause failures.
3. **Installation & update friction** – Bun compatibility (#4160), pnpm global install detection (#5689), trust dialog appearing in home directory (#5619), and CLI hanging after update (#5645) degrade the developer experience.
4. **Tool execution edge cases** – Bash overflow spill crashes (#5667), stale `tool_execution_update` after settlement (#5573), and subagent continue after ESC (#5685) show race conditions around async tool lifecycle.
5. **Session and compaction reliability** – Duplicate packages due to Shrinkwrap (#5653), broken `parentId` chains after fork (#5669), and compaction failures after reload (#5676) threaten data integrity for long‑running sessions.

*Digest generated for 2026‑06‑13 from [earendil-works/pi](https://github.com/earendil-works/pi) data.*

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest — 2026-06-13

## 1. Today's Highlights
This Friday brings a minor bugfix release (v0.18.0), but the community is in heated debate over a proposed **OAuth free‑tier quota reduction** (down to 100 requests/day) in Issue #3203, which has drawn 129 comments. In development, a major **Computer Use migration to a cross‑platform Rust driver** (#5051) and a **DaemonTransport abstraction for ACP compliance** (#5040) are moving forward, while the durable cron feature from last week is being hardened (#5076). A release pipeline failure (#5068) also caught attention.

## 2. Releases
**v0.18.0** was released today. The changelog includes a CLI fix to skip thought parts in copy output, and the usual CI bot version bump from v0.17.1. No breaking changes or major features.

## 3. Hot Issues (Top 10)

1. **[#3203] Qwen OAuth Free Tier Policy Adjustment**  
   *Status: needs‑triage · Comments: 129*  
   Proposes cutting daily free quota from 1,000 to 100 requests and then completely closing the free entry point on the 20th. The large number of comments reflects significant community push‑back.  
   [Issue #3203](https://github.com/QwenLM/qwen-code/issues/3203)

2. **[#4877] OpenWork Can't Distinguish Same Model from Different Providers**  
   *Priority: P2 · Comments: 4*  
   When two providers (e.g., OpenAI and a proxy) offer models with identical IDs, the UI shows them as one. The user’s `settings.json` example shows the problem clearly.  
   [Issue #4877](https://github.com/QwenLM/qwen-code/issues/4877)

3. **[#5075] ExitPlanMode Error – Plan Gate Fails, Full Plan Not Shown**  
   *Priority: P2 · Comments: 3*  
   Exiting plan mode triggers a gate failure that truncates the plan display instead of showing the full plan. A quick fix is already in progress via PR #5077.  
   [Issue #5075](https://github.com/QwenLM/qwen-code/issues/5075)

4. **[#5076] Durable Cron – Follow‑ups from #5004 Review**  
   *Priority: P2 · Comments: 2*  
   Captures deferred work: trust hardening for project‑controlled task files, lock heartbeat, and validation performance. Keeps the original PR focused.  
   [Issue #5076](https://github.com/QwenLM/qwen-code/issues/5076)

5. **[#5074] Add Persistent Sidebar to Web‑Shell for Session Management**  
   *Priority: P2 · welcome‑pr · Comments: 2*  
   Proposes a cmux‑like sidebar showing session list, creation, switching, renaming, and deletion – open by default. A welcome contribution opportunity.  
   [Issue #5074](https://github.com/QwenLM/qwen-code/issues/5074)

6. **[#5067] Focus‑Jump Gates Count Retained Terminal Agents, Not Panel Roster**  
   *Priority: P2 · waiting for feedback · Comments: 2*  
   A subtle follow‑up to #4911: keyboard focus jumps can target hidden agents because they use raw registry data instead of the rendered set, creating phantom selection slots.  
   [Issue #5067](https://github.com/QwenLM/qwen-code/issues/5067)

7. **[#5064] Statusline Should Wrap When Content Overflow**  
   *Priority: P3 · welcome‑pr · Comments: 2*  
   When statusline text is too long it gets hidden or overlapped. The request is simple: allow line wrapping.  
   [Issue #5064](https://github.com/QwenLM/qwen-code/issues/5064)

8. **[#4078] Allow fastModel to Use a Different Auth Type Than Main Session**  
   *Status: needs‑triage · Comments: 1*  
   Currently `fastModel` is silently ignored if its model ID belongs to a different authentication type. Users need to duplicate configs to use, e.g., local fast model with cloud main.  
   [Issue #4078](https://github.com/QwenLM/qwen-code/issues/4078)

9. **[#5068] Release Failed for v0.18.0‑nightly**  
   *Comments: 0*  
   Automated nightly release failed. Likely infrastructure issue; watch for a fix in the next 24h.  
   [Issue #5068](https://github.com/QwenLM/qwen-code/issues/5068)

10. **[#4956] Enable Fork Subagent by Default and Bubble Permission Prompts**  
    *Closed · Comments: 1*  
    Proposes promoting the fork subagent from experimental to always‑available with `approvalMode: bubble` by default. Closed after discussion, but the concept may reappear in future roadmaps.  
    [Issue #4956](https://github.com/QwenLM/qwen-code/issues/4956)

## 4. Key PR Progress (Top 10)

1. **[PR #5051] feat(core): Migrate Computer Use to cua‑driver (cross‑platform)**  
   Replaces the open‑computer‑use npm backend with a Rust driver (`cua‑driver‑rs`) that avoids focus stealing and speaks MCP over stdio. Still experimental, but a significant architecture change.  
   [PR #5051](https://github.com/QwenLM/qwen-code/pull/5051)

2. **[PR #5040] feat(sdk,serve): DaemonTransport Abstraction + ACP Standard Compliance**  
   Adds a pluggable `DaemonTransport` interface (REST+SSE, ACP HTTP+SSE, ACP WebSocket) and a `DaemonRouter` for serving ACP endpoints without forking the provider stack. Major SDK/daemon improvement.  
   [PR #5040](https://github.com/QwenLM/qwen-code/pull/5040)

3. **[PR #5034] feat(core): Workflow P3 – agent() Options**  
   Completes the dynamic workflow contract (P3 after P1/P2) by adding per‑call options `agentType`, `model`, `isolation: 'worktree'` to the `agent()` tool.  
   [PR #5034](https://github.com/QwenLM/qwen-code/pull/5034)

4. **[PR #5004] feat(core): Durable Cron Jobs – /loop Tasks That Survive Restarts**  
   Persists `/loop` task definitions per‑project under `~/.qwen/tmp/` and auto‑resumes after restart. Already merged as of this digest; the follow‑up issues (#5076) will harden it.  
   [PR #5004](https://github.com/QwenLM/qwen-code/pull/5004)

5. **[PR #5033] fix(serve): Add Prompt Queue Backpressure**  
   Adds backpressure to the serve prompt queue to prevent overload. Important for server stability under load.  
   [PR #5033](https://github.com/QwenLM/qwen-code/pull/5033)

6. **[PR #5077] fix(cli): Show Full Plan for Gate Failures**  
   Fixes Issue #5075 by returning the full `plan_summary` (with `rejected: true`) even when the plan gate does not approve execution.  
   [PR #5077](https://github.com/QwenLM/qwen-code/pull/5077)

7. **[PR #5073] fix: Warn on Oversized Context Instructions**  
   Emits a startup warning when QWEN.md / context instructions exceed 15% of the active model’s context window – helps users avoid silent truncation.  
   [PR #5073](https://github.com/QwenLM/qwen-code/pull/5073)

8. **[PR #5066] feat(web‑shell): Token Usage, Settings, Retry, Streaming Metrics, Hidden Commands**  
   Closed today. Brings structured token usage, i18n settings panel, theme/language pickers, compact mode persistence, and hidden command support to the daemon web shell.  
   [PR #5066](https://github.com/QwenLM/qwen-code/pull/5066)

9. **[PR #4598] feat(tui): Collapsible Thinking Blocks with Duration Timer**  
   Replaces the always‑expanded reasoning display with a collapsible, streaming block that shows duration upon completion. Continues to receive updates.  
   [PR #4598](https://github.com/QwenLM/qwen-code/pull/4598)

10. **[PR #4929] fix(cli): Add OSC 52 Clipboard Fallback for SSH Environments**  
    Adds OSC 52 escape sequence support for `/copy` and vim yank operations when no X11/Wayland is available (SSH). Closes Issue #4926.  
    [PR #4929](https://github.com/QwenLM/qwen-code/pull/4929)

## 5. Feature Request Trends

- **Session & UI Persistence:** Several requests target the web‑shell (persistent sidebar #5074, statusline wrapping #5064, collapsible todo panels #5069). The team is already shipping features in this area (e.g., PR #5066).
- **Model & Provider Flexibility:** #4078 (fastModel cross‑auth) and #4877 (distinguishing providers) show a clear desire to manage multiple providers more cleanly, especially when model IDs overlap.
- **Background / Daemon Automation:** #4956 (fork subagent default) and #5076 (durable cron hardening) indicate growing interest in background agents and persistent tasks.
- **Free Tier / OAuth Policy:** #3203 is a flashpoint: the community is actively resisting the proposed quota reduction – expect more discussion or a revision.
- **Workspace & File History:** #5057 (snapshot persistence) and #4884 (preserve CLI flags on resume) suggest users want robust workflow continuity.

## 6. Developer Pain Points

- **Model Identity Ambiguity:** The most common theme – providers with same model IDs cause confusion in UI, settings, and fastModel logic (Issues #4877, #4078). The fix in PR #5039 (using `id+baseUrl`) is a welcome step.
- **Plan & Gate Reliability:** #5075 (plan truncated on exit) and #5067 (phantom focus slots) show that the plan approval and keyboard navigation logic still has edge cases that disrupt the interactive loop.
- **Build & Release Stability:** The nightly release failure (#5068) is a minor but recurring concern. Combined with the FIFO blocking issue (#4894), there is a sense that CI/infrastructure needs attention.
- **Resource Limits:** #3203 (quota reduction) is a hot‑button external pain point, while #5073 (oversized instructions warning) highlights internal LLM context limits that users silently hit.
- **Clipboard in Non‑Graphical Environments:** #4929’s OSC 52 fallback directly addresses a long‑standing pain for SSH users who previously couldn’t copy output.

*Generated for 2026-06-13 by Qwen Code community analysis.*

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-06-13

## Today's Highlights
The project officially rebrands from `deepseek-tui` to **CodeWhale** with the v0.8.59 release; the legacy npm package is deprecated. Development velocity is intense on v0.8.60, with a flood of closed issues and PRs tackling an **Agent Fleet control plane**, headless sub‑agent runtimes, and first‑class provider routes (Z.ai, StepFlash, MiniMax). The community is actively contributing to ACP‑registry integration and UX polish.

## Releases
- **v0.8.59** – [Release notes](https://github.com/Hmbown/CodeWhale/releases/tag/v0.8.59)  
  Final rebrand release: project, command, npm package, and assets now use `CodeWhale`. Legacy `deepseek-tui` receives no further updates. Migration guide in `docs/REBRAND.md`.

## Hot Issues (10 noteworthy)

1. **[#3096](https://github.com/Hmbown/CodeWhale/issues/3096) – Split sub‑agents into headless worker runtime**  
   Proposed architecture change to decouple sub‑agents from heavy TUI context, enabling scalable fan‑out without per‑child UI overhead. Core to v0.8.60 reliability.

2. **[#3082](https://github.com/Hmbown/CodeWhale/issues/3082) – Group background tasks into workflows with phase summaries**  
   Addresses UI clutter from many parallel command cards. Introduces workflow‑level cards with elapsed time, agent count, token usage. Community screenshots provided design direction.

3. **[#3142](https://github.com/Hmbown/CodeWhale/issues/3142) – Agent run ledger with follow‑up, takeover, and artifact receipts**  
   Inspired by Cursor’s Cloud Agents, this proposes a persistent operational ledger rather than hidden chat branches. Closed with implementation.

4. **[#3154](https://github.com/Hmbown/CodeWhale/issues/3154) – EPIC: Agent Fleet control plane**  
   High‑level epic for always‑running, verifiable multi‑agent workflows. Introduces manager agent, heartbeat recovery, and escalation. Active discussion and many sub‑issues.

5. **[#3167](https://github.com/Hmbown/CodeWhale/issues/3167) – Model Agent Fleet org chart, roles, delegation policy**  
   Explicit role/delegation model (scouts, implementers, reviewers, verifiers) to avoid reinventing delegation in every prompt. Community feedback welcome.

6. **[#3178](https://github.com/Hmbown/CodeWhale/issues/3178) – Add `/swarm` as Whaleflow‑backed dynamic multi‑agent mode**  
   User‑facing entry point for dynamic multi‑agent work. Architectural constraint: must not resurrect CodeWhale’s legacy heavy sub‑agent UI.

7. **[#3192](https://github.com/Hmbown/CodeWhale/issues/3192) – Put CodeWhale up for agentclientprotocol/registry**  
   Community request to list CodeWhale in the ACP registry for easier installation from Zed. High community interest (3 👍 on #1447, now reopened).

8. **[#2890](https://github.com/Hmbown/CodeWhale/issues/2890) – Contribution gate workflow allowlist**  
   Restored from deleted issue; community members `@nightt5879` and `@kolief` offered to implement. Focus on an allowlist file and `CONTRIBUTING.md` documentation.

9. **[#1976](https://github.com/Hmbown/CodeWhale/issues/1976) – Goal mode: persistent objective/workflow surface**  
   Long‑standing feature to replace `/goal` with a persistent surface. Not a new execution mode but a UI/UX change for tracking objectives.

10. **[#3194](https://github.com/Hmbown/CodeWhale/issues/3194) – Audit helper command hints and shortcuts**  
    UX audit to ensure all advertised keybindings are accurate and discoverable. Motivated by confusing copy like `Alt+V` in footers.

## Key PR Progress (10 important)

1. **[#3193](https://github.com/Hmbown/CodeWhale/pull/3193) – Config‑gated Pro Plan routing profile**  
   Fresh PR from community contributor `dumbjack`. Adds an explicit `pro_plan_profile` gate (default off) without changing default menus.

2. **[#3191](https://github.com/Hmbown/CodeWhale/pull/3191) – First‑party Z.ai and StepFlash/StepFun provider routes**  
   Closes #3187. Adds native provider support for Z.ai (GLM-5.1 model) and StepFun/StepFlash with proper API endpoints and auth.

3. **[#3036](https://github.com/Hmbown/CodeWhale/pull/3036) – Hide internal IDs from normal UI**  
   Replaces raw UUIDs and hex agent IDs with stable user‑facing labels. Full identifiers remain in hover/detail. Improves readability.

4. **[#3038](https://github.com/Hmbown/CodeWhale/pull/3038) – Ctrl+B directly backgrounds foreground shell**  
   UX simplification: removes the intermediate ShellControlView menu. One‑keystroke backgrounding as requested in #3032.

5. **[#3035](https://github.com/Hmbown/CodeWhale/pull/3035) – Throttle AgentProgress redraws**  
   Fixes UI freeze when 4+ sub‑agents run concurrently. Limits redraw frequency and sidebar recomputation. Critical performance fix for v0.8.60.

6. **[#3039](https://github.com/Hmbown/CodeWhale/pull/3039) – OSC 8 out‑of‑band hyperlink infrastructure**  
   Foundational work for clickable links in the transcript using OSC 8 escape sequences. Bypasses ratatui buffer limitations.

7. **[#3040](https://github.com/Hmbown/CodeWhale/pull/3040) – Clickable sidebar rows**  
   Adds mouse‑click dispatch for Tasks and Agents panels. Clicking a background job shows it; clicking an agent row opens `/subagents`. Addresses #3028.

8. **[#3042](https://github.com/Hmbown/CodeWhale/pull/3042) – New CLI flags for `codewhale exec`**  
   Adds `--allowed-tools`, `--disallowed-tools`, `--max-turns`, `--append-system-prompt` for unattended/CI use. Important for automation.

9. **[#3054](https://github.com/Hmbown/CodeWhale/pull/3054) – Native Anthropic Messages API adapter**  
   Third wire dialect alongside existing providers. Supports `cache_control`, thinking blocks, and tool streaming. Closes #3014.

10. **[#3049](https://github.com/Hmbown/CodeWhale/pull/3049) – JSON decision contract for hooks**  
    Hooks can now return structured JSON decisions (`allow`/`deny`/`ask` + reason). Also adds glob matchers and project‑local hooks. Closes #3026.

## Feature Request Trends
- **Agent Fleet orchestration** – The overwhelming direction of v0.8.60: manager agents, run ledgers, heartbeat recovery, remote worker support (SSH), and verifiable task specs.
- **Multi‑agent UX** – `/swarm` mode, headless sub‑agent lanes, workflow grouping with phase summaries.
- **Provider diversity** – Strong demand for first‑class Z.ai, StepFlash, and MiniMax support beyond OpenRouter/OpenAI‑compatible slots.
- **ACP‑registry listing** – Community pushing for integration with `agentclientprotocol/registry` to simplify installation via Zed.
- **Persistent goal/objective surfaces** – Long‑running request for a “Goal mode” that replaces ad‑hoc `/goal` commands with a continuous workflow surface.

## Developer Pain Points
- **UI performance under sub‑agent load** – Freezes when multiple sub‑agents produce progress events simultaneously (#3035).
- **Hidden / undiscoverable shortcuts** – Users struggle with keyboard‑first navigation; request for inline help and labeled tooltips (#3194).
- **Incomplete provider reasoning support** – Many providers (Atlascloud, Moonshot, Ollama) had silent no‑ops for `reasoning_effort` tiers (#3050, #3046).
- **Unclear busy/idle states** – Version 0.8.53+ lacks visual distinction between thinking and completed tasks (#2982).
- **Hardcoded model validation** – Sub‑agent model IDs were restricted to DeepSeek, breaking multi‑provider setups (#3045).

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*