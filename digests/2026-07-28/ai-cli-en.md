# AI CLI Tools Community Digest 2026-07-28

> Generated: 2026-07-28 00:11 UTC | Tools covered: 9

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

# AI CLI Tools Cross-Tool Comparison Report — 2026-07-28

## 1. Ecosystem Overview

The AI CLI tools landscape remains intensely competitive, with nine major tools all targeting developer workflows ranging from autonomous code generation to multi-agent orchestration. The digest summaries reveal a maturing but still brittle ecosystem: billing trust, Windows GPU stability, model-access confusion, and subagent reliability are cross-cutting pain points. Release cadence varies from multiple daily alpha builds (OpenAI Codex, OpenCode) to quiet days with only PR activity (Kimi Code, Pi). Community engagement is highest for Claude Code and Codex, where single issues can attract 60+ upvotes and 50+ comments, while smaller tools like DeepSeek TUI and Qwen Code show focused, lower-volume participation. The overall trajectory points toward deeper MCP integration, better multi-agent orchestration, and renewed emphasis on terminal UX and cross-platform stability.

## 2. Activity Comparison

| Tool | Noteworthy Issues (today) | Key PRs (today) | Release Today? |
|------|---------------------------|----------------|----------------|
| Claude Code | 10 | 6 | No |
| OpenAI Codex | 10 | 10 | Yes (2 alpha builds) |
| Gemini CLI | 10 | 10 | Yes (nightly) |
| GitHub Copilot CLI | 10 | 8 | Yes (v1.0.76-0) |
| Kimi Code CLI | 4 | 4 | No |
| OpenCode | 10 | 10 | Yes (v1.18.6, v1.18.7) |
| Pi | 10 | 10 | No |
| Qwen Code | 4 | 10 | Yes (nightly + prerelease) |
| DeepSeek TUI | 5 | 10 | No |

**Notes:**  
- Issues count = hot issues highlighted in each digest (proxy for community pain points surfaced today).  
- PRs count = key PRs highlighted (proxy for maintainer/contributor output).  
- Release = any version bump in the last 24 hours (including alpha/nightly/patches).

## 3. Shared Feature Directions

Several cross-tool patterns emerge from community feedback and merged PRs:

### 3.1 Model Access & Metadata Transparency
- **Claude Code** – Fable 5 walled behind inconsistent plan/max-UI (#79337, #79597).  
- **OpenAI Codex** – Sol alias missing (#34027), Multi-Agent V1/V2 mismatch (#35097).  
- **Gemini CLI** – `gemini-3.5-flash` not selectable, fixed in #28485.  
- **GitHub Copilot CLI** – 0x‑multiplier models misrepresented in docs (#1116).  
- **Pi** – Z.AI `max_tokens` vs `max_completion_tokens` confusion (#7174).  
- **Need**: A standard, tier‑aware model registry that dynamically resolves aliases and capabilities.

### 3.2 Windows GPU/Rendering Stability
- **OpenAI Codex** – 5+ open issues (e.g., #32683, #34133, #35352) tied to SwiftShader/Code Integrity.  
- **Claude Code** – Desktop Browser pane crashes with GPU exit (#81275).  
- **Kimi Code CLI** – Two Unicode‑encoding crashes on non‑UTF‑8 consoles (#2561, #2560).  
- **Need**: Bundled software renderers that pass Windows Code Integrity; robust fallback paths.

### 3.3 Subagent/Agents Reliability & Transparency
- **Gemini CLI** – Subagent false‑positive "GOAL" success (#22323), unexpected invocation (#22093).  
- **OpenCode** – Missing `task_id` on subagent failure (#39196), infinite tool loops (#28596).  
- **OpenAI Codex** – Subagent disk bloat (#34061), quota drain overnight (#35463).  
- **GitHub Copilot CLI** – `task_complete` tool disappears in autopilot (#4161), Zombie processes (#4163).  
- **Need**: Clear success/failure feedback, resource cleanup, and usage accounting.

### 3.4 Authentication & Billing Trust
- **Claude Code** – July 17 billing incident with $700+ disputed charge (#81703).  
- **OpenAI Codex** – Reset counter fails to reset (#31606), Pro 20x users hit cap overnight (#35463).  
- **GitHub Copilot CLI** – macOS keychain prompts on every launch (#4273).  
- **Gemini CLI** – MCP OAuth token silently deleted (#28481 fix).  
- **Need**: Real‑time, auditable usage dashboards; automatic reconciliation; stable credential storage.

### 3.5 Terminal/CLI UX Refinements
- **DeepSeek TUI** – Foreground shell blocks user input (#4930), SSH/tmux key capture (#4925).  
- **Pi** – Mouse support in prompt (#7185), full re‑render every 1s (#7194).  
- **OpenCode** – Desktop theme only changeable once (#39205).  
- **GitHub Copilot CLI** – Blank screens on Windows Terminal (#4263).  
- **Need**: Deterministic input routing, configurable key bindings, performance profiling for large sessions.

## 4. Differentiation Analysis

| Tool | Core Focus | Target User | Technical Approach | Key Differentiators |
|------|------------|-------------|-------------------|---------------------|
| **Claude Code** | Full‑stack agent with collaboration (Cowork) | Enterprise teams | Anthropic‑optimized hooks & plugins; plan‑mode workflow | GitHub connector, Fable 5 model, billing complexity |
| **OpenAI Codex** | Multi‑agent orchestration + desktop app | Power users / Pro subscribers | Rust‑based CLI + Electron desktop; sandboxed subagents | Two alpha builds/day; strong model portfolio; Windows GPU pain |
| **Gemini CLI** | MCP‑first agent with sandboxing | Google ecosystem users | Node.js; deep MCP integration; seatbelt sandbox | Generalist agent delegation; OAuth token management |
| **GitHub Copilot CLI** | GitHub‑native autopilot + ACP protocol | GitHub users / CI/CD | TypeScript; ACP server; plans as first‑class objects | Tight GitHub integration; autopilot persistence; plan‑mode regression |
| **Kimi Code CLI** | VSCode extension + CLI | Moonshot AI users | Python; hooks system; focus on Chinese locale | Small but targeted; Windows encoding fixes; GC hook issue |
| **OpenCode** | Multi‑platform desktop + TUI | Open‑source community | Electron + Python; task‑based subagent model | Two dash releases; strong subagent orchestration; model‑agnostic |
| **Pi** | Extensible TUI platform | Developer power users | Node.js/Go; rich extension API; SQLite search | Extension‑first; provider interoperability (Bedrock, Z.AI); cache optimization |
| **Qwen Code** | CI/CD integration + Web Shell | Alibaba cloud / enterprise | TypeScript; Web Shell + Git workflow | Benchmark focus (SWE‑bench); CI automation; voice hold mode |
| **DeepSeek TUI** | Rust TUI for multi‑agent collaboration | Terminal‑first teams | Rust; Fleet concept; SSH/tmux friendly | Highly customizable TUI; v0.9.2 RC; avante.nvim compatibility fix |

## 5. Community Momentum & Maturity

**High Momentum / Mature Communities** (rapid iteration, high engagement, multiple daily updates):  
- **OpenAI Codex** – 10 key PRs + 2 alpha releases today; issues with 61 upvotes (#31606) show a vocal, demanding user base.  
- **OpenCode** – 10 key PRs + 2 patch releases; active subagent/task discussions.  
- **Claude Code** – 10 hot issues with 47 comments on top issue; billing and model access dominate, indicating a large paying user base.  

**Moderate Momentum** (steady PR flow, active but fewer high‑impact issues):  
- **Gemini CLI** – 10 PRs, nightly release, but community upvotes lower (max 8).  
- **GitHub Copilot CLI** – 8 PRs, patch release, but many issues are low‑engagement (3–6 comments).  
- **Pi** – 10 PRs, no release, but extension API work signals growing plugin ecosystem.  
- **Qwen Code** – 10 PRs, nightly release, but only 4 hot issues; community appears smaller but focused on CI and benchmarks.  

**Lower Momentum / Emerging** (fewer hot issues, slower iteration):  
- **Kimi Code CLI** – 4 PRs, no release; VSCode extension bugs unresolved.  
- **DeepSeek TUI** – 10 PRs but no release; RC integration still in progress; community comments low.  

**Overall**: The ecosystem is bifurcated. A few large vendors (Anthropic, OpenAI, Microsoft) dominate mindshare and pain‑point volume, while independent projects (Pi, DeepSeek, OpenCode) compensate with architectural innovation and community responsiveness.

## 6. Trend Signals

From the aggregated community feedback, several industry‑level trends are visible:

- **GPU & Rendering Stability Is the #1 Windows Blockade.** At least 6 distinct issues across Codex and Claude Code involve GPU process crashes, SwiftShader rejection, or code‑integrity failures. Any developer targeting Windows must prioritize Electron/Chromium fallback rendering and secure‑boot compliance.  
- **“Model Zoo” Fragmentation Is Accelerating.** Multiple tools report users unable to select newly released models (Gemini 3.5‑flash, Fable 5, Sol Luna) due to stale hardcoded lists. Tools need a dynamic model registry that queries provider APIs—or they risk becoming gatekeepers to the latest capabilities.  
- **Usage Billing Is a Trust‑Critical Feature.** Claude Code’s $700+ disputed charge and Codex’s reset‑counter bug show that when billing fails, user trust erodes overnight. Real‑time usage dashboards, proactive notifications, and automatic refund policies are table stakes.  
- **Subagent Orchestration Remains Immature.** “False success” (Gemini), “missing task IDs” (OpenCode), “zombie processes” (Copilot), and “quota drain” (Codex) all point to the same root cause: agent workflows lack reliable state tracking and resource isolation. This is the primary barrier to trust in autonomous mode.  
- **Terminal UX Is Being Rediscovered.** As tools add mouse support, configurable key bindings, and SSH‑friendly designs (DeepSeek, Pi, Copilot), it signals that the terminal is not dying—it’s becoming a richer interaction surface.  
- **MCP Is Standardizing but Still Fragile.** Every tool with MCP integration (Gemini, Copilot, Kimi, OpenCode, Pi) has fixed or is fixing OAuth token management, tool name normalization, or caching. The protocol is converging, but implementation quality varies.  
- **Extension Ecosystems Are Growing.** Pi’s extension API (scopedModels, message hooks, terminal colour schemes) and Qwen’s external context profiles indicate that tool vendors are betting on plugins to differentiate without bloating core functionality.  

**For developers evaluating AI CLI tools:**  
- Prioritize tools with transparent billing and active model registry updates.  
- If you use Windows, test GPU path thoroughly—or plan to use headless/remote sandbox.  
- For multi‑agent workflows, avoid tools with known false‑success bugs; prefer those with clear subagent status reporting.  
- Terminal purists will appreciate Pi and DeepSeek TUI for their extensibility and keyboard‑friendly defaults.  
- Enterprise teams may lean toward Claude Code (GitHub connector, plan‑mode) or Codex (sandboxed subagents, heavy automation), but must budget for ongoing stability issues.

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report
**Data snapshot:** 2026-07-28 | Source: github.com/anthropics/skills

---

## 1. Top Skills Ranking

Most-discussed Pull Requests by community engagement, ordered by comment volume:

### #1 — skill-creator: `run_eval.py` always reports 0% recall
**PR:** [github.com/anthropics/skills/pull/1298](https://github.com/anthropics/skills/pull/1298) | **Status:** Open (since 2026-06-10)
**Functionality:** Fixes the core evaluation script in the skill-creator toolchain — the description-optimization loop (used by `run_loop.py` and `improve_description.py`) has been optimizing against noise because `run_eval.py` reports `recall=0%` for every skill description. This is the culmination of a long-standing bug chain (see Issues #556, #1169).
**Discussion highlights:** Community members report 10+ independent reproductions. The fix requires installing the eval artifact as a real skill, correcting Windows stream reading, trigger detection logic, and parallel worker handling. This is the most-commented PR in the repository and represents the community's primary pain point.

### #2 — document-typography skill
**PR:** [github.com/anthropics/skills/pull/514](https://github.com/anthropics/skills/pull/514) | **Status:** Open (since 2026-03-04)
**Functionality:** Typographic quality control for AI-generated documents — catches orphan word wrap, widow paragraphs, and numbering misalignment. Addresses a universal quality issue across all Claude-generated documents.
**Discussion highlights:** Users note these are "problems every document Claude generates" faces, rarely caught by users who don't know to ask for typographic fixes.

### #3 — fix(pdf): case-sensitive file references in SKILL.md
**PR:** [github.com/anthropics/skills/pull/538](https://github.com/anthropics/skills/pull/538) | **Status:** Open (since 2026-03-06)
**Functionality:** Corrects 8 case-sensitivity mismatches in the PDF skill's SKILL.md where `REFERENCE.md` and `FORMS.md` were referenced in uppercase while actual files are lowercase — breaking on case-sensitive filesystems.
**Discussion highlights:** Part of a broader quality assurance wave; the community is increasingly catching cross-platform portability bugs in existing skills.

### #4 — ODT skill (OpenDocument)
**PR:** [github.com/anthropics/skills/pull/486](https://github.com/anthropics/skills/pull/486) | **Status:** Open (since 2026-03-01)
**Functionality:** Creates, fills, reads, and converts OpenDocument Format files (.odt, .ods). Covers LibreOffice document creation, template filling, and ODT-to-HTML conversion.
**Discussion highlights:** Strong demand for ISO standard / open-source document formats beyond the existing DOCX skill. The skill covers both reading and writing ODF files.

### #5 — frontend-design skill clarity improvement
**PR:** [github.com/anthropics/skills/pull/210](https://github.com/anthropics/skills/pull/210) | **Status:** Open (since 2026-01-05)
**Functionality:** Revises the existing frontend-design skill to improve actionability — ensuring every instruction is followable within a single conversation and guidance is specific enough to steer behavior without being overly prescriptive.
**Discussion highlights:** Exemplifies the community's growing concern with skill quality: making instructions Claude-executable rather than human-readable documentation.

### #6 — skill-quality-analyzer and skill-security-analyzer
**PR:** [github.com/anthropics/skills/pull/83](https://github.com/anthropics/skills/pull/83) | **Status:** Open (since 2025-11-06)
**Functionality:** Two meta-skills: one evaluates skills across five quality dimensions (structure, documentation, examples, resource usage, compliance); the other scans for security concerns in skill code.
**Discussion highlights:** Longest-running open PR with sustained interest. The community recognizes the need for self-evaluation tools as the skills ecosystem grows.

### #7 — testing-patterns skill
**PR:** [github.com/anthropics/skills/pull/723](https://github.com/anthropics/skills/pull/723) | **Status:** Open (since 2026-03-22)
**Functionality:** Comprehensive testing skill covering the testing trophy model, unit testing (AAA pattern), React component testing, integration testing, end-to-end testing, and mocking strategies.
**Discussion highlights:** Broad community enthusiasm for a "one skill to rule all testing" — consolidates fragmented testing guidance into a single, opinionated skill.

### #8 — Pyxel retro game development skill
**PR:** [github.com/anthropics/skills/pull/525](https://github.com/anthropics/skills/pull/525) | **Status:** Open (since 2026-03-05)
**Functionality:** Integrates with the Pyxel retro game engine via MCP, enabling Claude to create pixel-art/8-bit games with an iterative write → run → inspect → refine workflow.
**Discussion highlights:** Uniquely creative skill targeting game development enthusiasts. Demonstrates the community's appetite for domain-specific skills beyond enterprise use cases.

---

## 2. Community Demand Trends

From Issues sorted by comment volume, the community's most-anticipated skill directions are:

**1. Security and Trust Boundaries** — Issue [#492](https://github.com/anthropics/skills/issues/492) (43 comments) is the single most-commented issue in the repository. The community is deeply concerned about community skills distributed under the `anthropic/` namespace being perceived as official. This has sparked a wider conversation about trust boundary enforcement, skill provenance, and permission management. **Related:** Issue [#1175](https://github.com/anthropics/skills/issues/1175) on security concerns with SharePoint document handling via agent skills, and the [skill-security-analyzer PR #83](https://github.com/anthropics/skills/pull/83).

**2. Organizational Skill Sharing** — Issue [#228](https://github.com/anthropics/skills/issues/228) (16 comments, 8 👍) asks for org-wide skill libraries and direct sharing links rather than manual file downloads and uploads. This is the highest-voted issue and reflects enterprise adoption friction.

**3. Skill-Creator Toolchain Reliability** — Issues [#556](https://github.com/anthropics/skills/issues/556) (12 comments), [#1169](https://github.com/anthropics/skills/issues/1169) (3 comments), [#1061](https://github.com/anthropics/skills/issues/1061) (3 comments), and [#202](https://github.com/anthropics/skills/issues/202) (8 comments) collectively document the broken `run_eval.py` / `run_loop.py` pipeline, Windows compatibility failures, and the need to rewrite `skill-creator` as an operational skill rather than developer documentation. This is the infrastructure bottleneck preventing effective skill creation by the community.

**4. Agent Memory & Reasoning Quality** — Issue [#1329](https://github.com/anthropics/skills/issues/1329) (9 comments) proposes `compact-memory` — a symbolic notation skill for managing long-running agent state efficiently. Issue [#1385](https://github.com/anthropics/skills/issues/1385) proposes a three-gate reasoning quality pipeline. These signal growing sophistication: the community wants skills that manage *Claude's own cognitive processes*, not just output formatting.

**5. Document Format Expansion** — Issues and PRs around ODT ([#486](https://github.com/anthropics/skills/pull/486)), typography ([#514](https://github.com/anthropics/skills/pull/514)), and context window exhaustion from `claude-api` ([#1487](https://github.com/anthropics/skills/issues/1487)) show sustained demand for document-centric skills across formats and quality dimensions.

**6. Duplicate Skill Management** — Issue [#189](https://github.com/anthropics/skills/issues/189) (6 comments, 9 👍) reports that `document-skills` and `example-skills` plugins install identical content, wasting context window. Community wants deduplication and clear ownership boundaries between skill packages.

---

## 3. High-Potential Pending Skills

These active PRs have strong community discussion and may land soon:

**self-audit — Mechanical Verification + Reasoning Quality Gate** ([#1367](https://github.com/anthropics/skills/pull/1367))
*Author: YuhaoLin2005 | Updated: 2026-07-02*
A universal skill that audits AI output before delivery: mechanical file verification followed by four-dimension reasoning audit (in damage-severity priority order). Works across any project, tech stack, or model. Directly addresses the reasoning quality concerns raised in Issue [#1385](https://github.com/anthropics/skills/issues/1385).

**plan-file-hygiene** ([#1479](https://github.com/anthropics/skills/pull/1479))
*Author: Palo-Alto-AI-Research-Lab | Updated: 2026-07-27*
Addresses the problem of planning artifacts accumulating in project directories with no lifecycle management. Built on community framing from Issue [#1417](https://github.com/anthropics/skills/issues/1417). Very recent — could merge quickly.

**testing-patterns** ([#723](https://github.com/anthropics/skills/pull/723))
*Author: 4444J99 | Updated: 2026-04-21*
Comprehensive testing skill covering the full testing stack. Broad community interest, sustained discussion over three months.

**pyxel retro game development** ([#525](https://github.com/anthropics/skills/pull/525))
*Author: kitao | Updated: 2026-07-15*
Mature PR (4+ months) with steady updates. Integrates with an existing MCP server. Unusual for being a creative/entertainment skill in an otherwise enterprise-oriented collection.

**document-typography** ([#514](https://github.com/anthropics/skills/pull/514))
*Author: PGTBoos | Updated: 2026-03-13*
Second most-commented PR. Addresses a universal document quality issue. Has been open since March — could benefit from a maintainer to push it across the finish line.

---

## 4. Skills Ecosystem Insight

**The community's most concentrated demand is for foundational toolchain reliability (fixing the skill-creator evaluation pipeline, especially on Windows) combined with a converging push toward new skill categories in document formatting/typography, testing patterns, security governance, and agent memory management — reflecting a community that has shifted from "can we make skills?" to "how do we make good skills, share them safely, and manage their lifecycle?"**

---

# Claude Code Community Digest — 2026-07-28

## Today’s Highlights
Fable 5 continues to dominate the issue tracker: multiple Max-plan users report being falsely blocked by a “usage credits required” wall, with the model picker and usage dashboard giving contradictory information. Meanwhile, a long‑running GitHub connector regression (#71542) remains unresolved, leaving many unable to access any repository content. A July 17 billing incident (#81703) has also drawn fresh attention, with one user disputing $704.71 in erroneous usage‑credit charges.

## Releases
No new releases in the last 24 hours.

## Hot Issues
1. **[#79337 – Fable 5 prompts ‘usage credits required’ on Max plan](https://github.com/anthropics/claude-code/issues/79337)** — Opened Jul 20, 47 comments, 16 👍. The first day Fable 5 became part of Max plans, sessions were silently downgraded to Opus 4.8. Community response is intense: many Max subscribers feel misled.

2. **[#71542 – GitHub connector cannot access any repository content](https://github.com/anthropics/claude-code/issues/71542)** — 43 comments, 37 👍. An account‑wide regression: connections succeed but content retrieval fails for both public and private repos. High severity for CI/CD workflows.

3. **[#57371 – Disable Windows Cowork background service](https://github.com/anthropics/claude-code/issues/57371)** — 15 comments, 39 👍. Users who don’t use Cowork want to stop the always‑running CoworkVMService. Strong demand for opt‑out control.

4. **[#79597 – Fable 5 falsely walled in interactive picker for headless auth](https://github.com/anthropics/claude-code/issues/79597)** — 8 comments, 9 👍. Similar to #79337 but specific to `setup-token` authentication; `-p` headless mode works, interactive picker does not.

5. **[#81703 – July 17 mass billing incident: usage credits charged despite plan](https://github.com/anthropics/claude-code/issues/81703)** — 5 comments. Anthropic acknowledged the incident but charges remain unreconciled. One user reports $704.71 in disputed charges.

6. **[#81275 – Desktop Browser pane crashes on Windows (GPU process exit)](https://github.com/anthropics/claude-code/issues/81275)** — 5 comments. Opening the Cowork browser preview crashes Claude Desktop with a consistent GPU exit code. Affects Intel, NVIDIA, and software rendering.

7. **[#72455 – Fullscreen renderer corrupts system‑wide macOS clipboard](https://github.com/anthropics/claude-code/issues/72455)** — 5 comments, 5 👍. Extreme severity: fullscreen Claude Code breaks copy/paste in all applications.

8. **[#78315 – Browser tool read actions ignore allowed sites list](https://github.com/anthropics/claude-code/issues/78315)** — 6 comments. Navigation to allowed domains works without prompting, but read/interact actions still require per‑action approval. Closed as invalid? Actually status is CLOSED invalid – but worth noting as it highlights confusion.

9. **[#57882 – Android image attachments not transmitted in shared sessions](https://github.com/anthropics/claude-code/issues/57882)** — 5 comments, 5 👍. Camera/gallery images fail to sync in Cowork shared sessions; text syncs fine.

10. **[#70115 – Existing Max subscriber locked out across all surfaces](https://github.com/anthropics/claude-code/issues/70115)** — 2 comments. Recurring auth‑routing failure: magic‑link/OAuth routes to “create account”. Cross‑referenced with five similar past issues → indicates systemic backend problem.

## Key PR Progress
*(Only 6 PRs were updated in the last 24 hours; all are listed.)*

1. **[#81673 – fix(devcontainer): don’t abort firewall setup on DNS failure](https://github.com/anthropics/claude-code/pull/81673)** — Fixes #55623. `init-firewall.sh` exits early when `statsig.anthropic.com` fails to resolve, leaving a half‑populated ipset. PR makes domain resolution non‑fatal.

2. **[#81672 – fix(hookify): make import independent of install directory name](https://github.com/anthropics/claude-code/pull/81672)** — Fixes #69665 and #81448. Hook entry points assumed the plugin directory is always named `hookify`; marketplace installs break. PR uses a more robust path resolution.

3. **[#81670 – fix(plugins): quote ${CLAUDE_PLUGIN_ROOT} in hook commands](https://github.com/anthropics/claude-code/pull/81670)** — Fixes #78490 and #79143. Unquoted variable expansions break hooks when paths contain spaces. Also prefixes hookify examples for clarity.

4. **[#20448 – Add web4‑governance plugin](https://github.com/anthropics/claude-code/pull/20448)** — A long‑standing PR (since January) adding a trust‑governance plugin with T3 trust tensors and R6 audit trails. Still open with no recent activity – may be stale.

5. **[#81576 – docs: fix security‑guidance plugin entry in README](https://github.com/anthropics/claude-code/pull/81576)** — Corrects inaccurate description: plugin has no `PreToolUse` hook, has 25 patterns (not 9), and three distinct triggers.

6. **[#81540 – Fix #80705: usage leak bug (Atlas 2 automated)](https://github.com/anthropics/claude-code/pull/81540)** — Automated contribution from the “Atlas 2” bot, claiming to fix a usage leak. Stated reward $200. Repository validation was run.

## Feature Request Trends
Several recurring themes emerge from recent issues:

- **Plan approval UX** – Multiple users request “accept, clear context and auto mode” parity on remote‑control surfaces (e.g., headless) (#81393).
- **Config vs. state separation** – Strong desire to split `~/.claude` into portable settings (`settings.json`, `rules/`, `CLAUDE.md`) and machine‑local ephemera (`cache/`, `sessions/`) (#81392). A recommended `.gitignore` would help.
- **Stable project identity for sync** – Auto‑memory is keyed by absolute path, preventing cross‑machine synchronization (#81391). Users want a first‑class project ID.
- **Workflow tool token waste** – The `Workflow` tool eagerly loads ~4k tokens even when opt‑in only; request to defer loading or add a toggle (#79504).
- **Cowork / Browser tool permissions** – Users want consistent behaviour: “Allowed sites” should cover all actions, and the Windows Cowork service should be disableable (#57371, #78315).

## Developer Pain Points
The most impactful recurring frustrations this week:

- **Fable 5 access confusion** – Multiple issues (#79337, #79597, #79412, #81350) show a systematic bug: the model picker and `/usage` dashboard disagree on whether Fable 5 is covered by Max plans. Headless workflows break or behave differently.
- **GitHub connector regression** – #71542 has now been open for over a month with 37 upvotes but no fix; developers relying on repository integration are blocked.
- **Auth‑routing failures** – #70115 and five linked past issues show recurring problems with OAuth/magic‑link flows routing to account creation instead of login.
- **Desktop stability** – The Browser pane crash on Windows (#81275) and clipboard corruption on macOS (#72455) are severe, system‑affecting bugs.
- **Billing trust eroded** – The July 17 incident (#81703) saw usage‑credit charges despite plan allowances, with one user disputing over $700. Lack of automatic reconciliation adds to frustration.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex Community Digest — 2026-07-28

## Today’s Highlights

Two new Rust alpha builds (0.146.0-alpha.12 and .13) landed overnight, but the community’s attention remains fixed on a wave of Windows‑specific crashes—especially GPU process failures triggered by Code Integrity blocking bundled SwiftShader DLLs. Meanwhile, a long‑standing reset‑counter bug (#31606) and a new “Codex Diff” crash on macOS (#35058) are drawing heavy engagement, with 61 and 48 upvotes respectively.

## Releases

- [rust-v0.146.0-alpha.12](https://github.com/openai/codex/releases/tag/rust-v0.146.0-alpha.12) — Alpha release with no further details.
- [rust-v0.146.0-alpha.13](https://github.com/openai/codex/releases/tag/rust-v0.146.0-alpha.13) — Alpha release on top of .12; likely contains the crossterm fork patch and the Windows exec yield floor changes merged in today’s PRs.

## Hot Issues (10 Noteworthy)

1. **[#31606 – Reset failed, did not apply and 1 reset is wasted](https://github.com/openai/codex/issues/31606)**  
   *52 comments, 61 👍*  
   Pro users on Windows report that using a reset token decrements the counter without actually resetting anything. The high vote count suggests this is a widespread cap‑exhaustion pain point.

2. **[#32683 – Codex App crashes in CrBrowserMain when Browser Use opens a page](https://github.com/openai/codex/issues/32683)**  
   *27 comments, 8 👍*  
   A `0xC0000005` access violation in `chrome.dll` on Windows. The in‑app browser, used by both Computer Use and web automation, is a frequent crash vector.

3. **[#34133 – Page.captureScreenshot crashes GPU process after Code Integrity rejects vk_swiftshader.dll](https://github.com/openai/codex/issues/34133)**  
   *24 comments*  
   Windows 10 users see full freezes after an agent takes a screenshot. The GPU process crash is tied to Code Integrity blocking the bundled Vulkan software renderer—a known tension between security policies and Electron’s GPU stack.

4. **[#35058 – Codex Diff crashes with “Oops, an error has occurred” in VS Code on macOS](https://github.com/openai/codex/issues/35058)**  
   *20 comments, 48 👍*  
   Apple Silicon users cannot review Codex changes because the diff view is broken in every workspace. A top upvote count signals high demand for a fix.

5. **[#34061 – Insane Codex Disk Usage from Subagents](https://github.com/openai/codex/issues/34061)**  
   *14 comments*  
   CLI users on macOS report that subagents amass gigabytes of disk space without cleanup. This is a quality‑of‑life issue for anyone running multi‑agent workflows locally.

6. **[#35352 – Codex Desktop exits when embedded browser GPU process crashes (unsigned SwiftShader blocked)](https://github.com/openai/codex/issues/35352)**  
   *12 comments*  
   Another Windows GPU crash variant—the desktop app exits entirely when the in‑app browser’s GPU process dies and the fallback SwiftShader is rejected.

7. **[#24268 – Codex Desktop Windows+WSL resolves bundled plugin cache as invalid C:\mnt\c path](https://github.com/openai/codex/issues/24268)**  
   *10 comments, 3 👍*  
   A long‑standing bug (since May) where the Desktop plugin cache misinterprets WSL paths. Still open, indicating a tricky cross‑environment fix.

8. **[#34027 – The 'gpt-5.6-sol' model is not supported when using Codex with a ChatGPT account](https://github.com/openai/codex/issues/34027)**  
   *5 comments, 5 👍 (Closed)*  
   The “Sol” model alias disappeared after an update, breaking CLI users. Closed quickly, but the model‑compatibility gap between ChatGPT accounts and Codex remains a recurring theme.

9. **[#35097 – gpt-5.6-luna is marked as MultiAgent V1, so V2 spawn_agent rejects it](https://github.com/openai/codex/issues/35097)**  
   *3 comments, 5 👍*  
   Multi‑Agent version mismatch—a configuration error that prevents subagent spawning for certain models. Points to a broader need for better model metadata management.

10. **[#35463 – Codex subagents drain full week quota overnight – usage counting broken](https://github.com/openai/codex/issues/35463)**  
    *3 comments*  
    Pro 20x users hit their weekly cap in a single overnight run because subagent usage is incorrectly counted. This is a critical billing/availability bug for power users.

## Key PR Progress (10 Important Merges)

1. **[#35688 – Point crossterm patch to the OpenAI OSS fork](https://github.com/openai/codex/pull/35688)**  
   Updates the terminal library dependency to an OpenAI‑maintained fork, likely to work around upstream bugs affecting Windows or macOS rendering.

2. **[#35685 – Load cloud‑managed profiles for `codex sandbox`](https://github.com/openai/codex/pull/35685)**  
   Enables sandboxed containers to accept permission profiles from cloud configuration, improving multi‑tenant isolation and policy enforcement.

3. **[#35678 – Preserve paginated thread metadata across resumes](https://github.com/openai/codex/pull/35678)**  
   Fixes a loss of thread title/preview when paginated history is reloaded, which has been confusing users resuming long sessions.

4. **[#35675 – Prepare MCP and plugin recommendations concurrently](https://github.com/openai/codex/pull/35675)**  
   Reduces startup latency by parallelising MCP runtime preparation and plugin discovery—a welcome optimisation for users with many connected tools.

5. **[#35671 – Route curated plugins by authentication mode](https://github.com/openai/codex/pull/35671)**  
   Ensures the correct plugin set is shown after an account switch or when using API key authentication, closing a source of confusion in the merged ChatGPT/Codex desktop.

6. **[#35670 – Raise the Windows exec yield floor to 10 seconds](https://github.com/openai/codex/pull/35670)**  
   A direct response to Windows crashes: clamping the minimum `exec_command` yield time prevents premature timeouts that could race with GPU/process startup.

7. **[#35655 – Terminate Windows non‑TTY processes on interrupt](https://github.com/openai/codex/pull/35655)**  
   Fixes a long‑standing gap where Ctrl‑C did not work for non‑interactive exec sessions on Windows, an important fix for automation scripts.

8. **[#35652 – Enable network policy callbacks for remote exec](https://github.com/openai/codex/pull/35652)**  
   Allows remote execution to participate in network‑policy decisions, aligning remote and local sandbox behaviour—critical for enterprise security.

9. **[#35649 – Preserve TUI input when terminal focus returns](https://github.com/openai/codex/pull/35649)**  
   Stops keystrokes from being lost when the terminal regains focus (e.g., after a palette refresh), a subtle but annoying UX bug in CLI mode.

10. **[#35642 – Make OpenTelemetry provider shutdown idempotent](https://github.com/openai/codex/pull/35642)**  
    Prevents double‑flush or panic during graceful shutdown, improving telemetry reliability for developers monitoring their Codex agents.

## Feature Request Trends

- **State & session fidelity** – Several issues ask for better “residual” tracking (#35528): when tool output is truncated or context compacted, the system should remember what was omitted and whether it can be recovered. This is driven by multi‑step agent workflows where lost context leads to repeated work.
- **OAuth/MCP lifecycle** – Issue #35006 is an umbrella for enterprise SSO login reliability, including reauthentication flows and credential‑store locking. Enterprise users need robust OAuth token management for custom MCP servers.
- **Voice integration for mobile remote** – #35687 requests real‑time speech conversations on Codex Mobile Remote, mirroring the desktop voice mode. This reflects growing demand for “code with me” voice interactions across devices.
- **Better model alias handling** – Issues like #34027 and #35097 highlight confusion around model naming (Sol, Luna, Terra) and compatibility between ChatGPT accounts and Codex. Users want a clear, consistent model picker that respects account tier.

## Developer Pain Points

1. **Windows GPU/rendering crashes** – At least five open issues (#32683, #34133, #35352, #34450, #35637) involve crashes, freezes, or input lag on Windows, often linked to SwiftShader, GPU process crashes, or Code Integrity policies. This is the single biggest stability problem for Windows users.
2. **Rate‑limit & usage miscounting** – Issues #31606 (reset not applying), #35463 (quota drain overnight), and #32613 (cache invalidation spikes) show that the usage accounting system is unreliable, causing real monetary or availability loss for Pro subscribers.
3. **Subagent resource leaks** – Disk bloat from subagents (#34061) and orphaned node_repl workers (#35582) are common complaints among power users running multi‑agent workflows on macOS and Linux.
4. **Model configuration confusion** – Duplicate model entries (#35493), missing Sol alias (#34027), and Multi‑Agent V1/V2 mismatches (#35097) indicate that the model registry and the client‑side model picker are not well synchronised. Developers frequently have to guess which model is actually available.
5. **Cross‑platform path handling** – The Windows+WSL plugin cache bug (#24268) and the GitLab webview crash (#35637) suggest that Codex still struggles with mixed filesystem environments, especially on Windows with WSL or long paths.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-07-28

## Today's Highlights
The v0.54.0 nightly was published, bundling a critical fix for macOS sandbox crashes (seatbelt profiles) and a fix for MCP OAuth token refresh failures that were silently deleting user credentials. On the issue tracker, long-running bugs around agent hangs and subagent false success reporting continue to draw community attention, while the "generalist agent hang" (#21409) remains one of the most upvoted unresolved issues.

## Releases

**v0.54.0-nightly.20260727.g3818efbbf**  
[Full Changelog](https://github.com/google-gemini/gemini-cli/compare/v0.54.0-nightly.20260726.g3818efbbf...v0.54.0-nightly.20260727.g3818efbbf)

This nightly picks up several fixes merged earlier this week, including:
- Fallback to embedded macOS seatbelt profiles if system files are missing (PR #28551)
- MCP OAuth token refresh using the stored client ID (PR #28481)
- Model selector support for `gemini-3.5-flash` (PR #28485)

No stable release today.

---

## Hot Issues (10 notable)

**[#22323] Subagent recovery after MAX_TURNS reported as GOAL success**  
*P1 bug | 12 comments | 2 👍*  
[Issue Link](https://github.com/google-gemini/gemini-cli/issues/22323)  

A `codebase_investigator` subagent hits max turns but returns `status: "success"` with `Termination Reason: "GOAL"` — hiding the interruption from users. This is a dangerous false positive that masks incomplete work. Community: high engagement, marked `need-retesting`.

**[#21409] Generalist agent hangs forever**  
*P1 bug | 8 comments | 8 👍*  
[Issue Link](https://github.com/google-gemini/gemini-cli/issues/21409)  

When Gemini CLI defers to the generalist agent for simple tasks (e.g., folder creation), it hangs indefinitely. The workaround (instructing the model not to use subagents) is impractical. The highest 👍 count among open bugs.

**[#25166] Shell command execution stuck with "Waiting input"**  
*P1 bug | 4 comments | 3 👍*  
[Issue Link](https://github.com/google-gemini/gemini-cli/issues/25166)  

After completing a simple CLI command, the agent hangs with "Awaiting user input" even for commands that cannot prompt. Affects basic shell execution reliability.

**[#21983] Browser subagent fails on Wayland**  
*P1 bug | 4 comments | 1 👍*  
[Issue Link](https://github.com/google-gemini/gemini-cli/issues/21983)  

Browser agent errors out on Wayland display servers, reporting `Termination Reason: GOAL` on failure. A platform-specific blocker affecting Linux users.

**[#22093] Subagents running without permission since v0.33.0**  
*P2 bug | 3 comments | 0 👍*  
[Issue Link](https://github.com/google-gemini/gemini-cli/issues/22093)  

Subagents (e.g., generalist) are invoked even when disabled in all configurations. User expected only MCP functionality. Suggests a permissions regression.

**[#11799] Gemini CLI ignores `GEMINI.md` context file**  
*P1 bug (closed) | 5 comments | 4 👍*  
[Issue Link](https://github.com/google-gemini/gemini-cli/issues/11799)  

While closed, this remains highly relevant: `GEMINI.md` shows in `/memory show` but the model ignores it unless explicitly mentioned with `@GEMINI.md`. High community reaction suggests this "closed but not fixed" frustration.

**[#24246] 400 error with > 128 tools**  
*P2 bug | 3 comments | 0 👍*  
[Issue Link](https://github.com/google-gemini/gemini-cli/issues/24246)  

Enabling many tools causes a 400 error. The agent lacks tool-scoping awareness. A scalability limitation for power users.

**[#22672] Agent should stop/discourage destructive behavior**  
*P2 feature | 3 comments | 1 👍*  
[Issue Link](https://github.com/google-gemini/gemini-cli/issues/22672)  

Model uses `git reset`, `--force`, or dangerous DB commands when safer alternatives exist. Community: wants safeguards and prompt-level guardrails.

**[#26516] Memory system bugs and quality improvements**  
*P2 bug (meta) | 2 comments | 0 👍*  
[Issue Link](https://github.com/google-gemini/gemini-cli/issues/26516)  

Tracker for Auto Memory issues: indefinite retrying of low-signal sessions (#26522), invalid inbox patches (#26523), missing deterministic redaction (#26525). A systemic quality concern.

**[#21000] Experiment with native file tools for task tracker**  
*P3 bug | 4 comments | 0 👍*  
[Issue Link](https://github.com/google-gemini/gemini-cli/issues/21000)  

Proposes using native file tools for creating and maintaining the task tracker, potentially improving reliability over the current agent-driven approach.

---

## Key PR Progress (10 notable)

**[#28551] fix(cli): fall back to embedded macOS seatbelt profiles if missing**  
[PR Link](https://github.com/google-gemini/gemini-cli/pull/28551)  

Critical fix for sandbox mode (-s) crashing on macOS/gMac when `.sb` profiles are absent from runfiles or bundle folder. Uses embedded copies as fallback. Merged into today's nightly.

**[#28481] fix(core): refresh MCP OAuth tokens with the stored client ID**  
[PR Link](https://github.com/google-gemini/gemini-cli/pull/28481)  

Fixes OAuth token refresh failure that was silently deleting stored credentials, forcing re-authentication. Affects MCP servers configured via OAuth discovery.

**[#28485] fix(cli): add gemini-3.5-flash to model selector for all users**  
[PR Link](https://github.com/google-gemini/gemini-cli/pull/28485)  

Resolves #28483 — users on v0.51.0+ could not select `gemini-3.5-flash` or `gemini-3.6-flash` due to a stale default model constant.

**[#28549] fix(mcp): disclose that Plan Mode read-only status is a server claim**  
[PR Link](https://github.com/google-gemini/gemini-cli/pull/28549)  

Plan Mode's read-only status relies on an unverified `readOnlyHint` from the MCP server. This PR adds transparency by marking tools as "server claims read-only." Addresses a trust boundary issue.

**[#28546] fix(core): strip Authorization header when using GEMINI_API_KEY auth**  
[PR Link](https://github.com/google-gemini/gemini-cli/pull/28546)  

Fixes #28538 — leftover `Authorization` headers from custom config caused Google API endpoints to reject requests when authenticating via `x-goog-api-key`.

**[#28364] fix(core): deep-merge user model config over defaults**  
[PR Link](https://github.com/google-gemini/gemini-cli/pull/28364)  

Shallow merges were overwriting deeply nested default configuration (aliases, overrides). This PR implements proper deep merging, fixing config inheritance issues.

**[#28369] feat(evals): add local report command and developer documentation**  
[PR Link](https://github.com/google-gemini/gemini-cli/pull/28369)  

Adds `npm run eval:report` to aggregate pass rates from Vitest `report.json`, mapping results back to inventory policies with duplicate test support. Improves developer tooling for behavioral evaluations.

**[#28363] fix(core): prevent AbortSignal listener leak in ShellExecutionService**  
[PR Link](https://github.com/google-gemini/gemini-cli/pull/28363)  

Fixes #28280 — AbortSignal listeners were not removed when a process finished naturally, causing memory leaks in long-lived sessions.

**[#28442] Main (large PR — details unclear)**  
[PR Link](https://github.com/google-gemini/gemini-cli/pull/28442)  

A large PR (size/xl) with limited description. Marked P1 and `status/pr-nudge-sent`. Requires review for scope and impact.

**[#28539] chore(deps): bump the npm-dependencies group with 75 updates**  
[PR Link](https://github.com/google-gemini/gemini-cli/pull/28539)  

Massive dependency update covering `@modelcontextprotocol/sdk`, `simple-git`, and 73 others. Merged today — potential for regressions.

---

## Feature Request Trends

1. **Agent self-awareness and tool utilization** — Multiple requests (e.g., #21432, #21968) ask the agent to better understand its own flags, hotkeys, and when to delegate to sub-agents or use custom skills. Users want the model to act as its own expert guide.

2. **AST-aware code operations** — Issues #22745 and #22746 explore using AST-aware file reads, searches, and codebase mapping to reduce turn count, improve accuracy, and reduce token noise. Active investigation with tools like `tilth`/`glyph`.

3. **Subagent trajectory visibility** — #22598 and #21763 request that subagent sessions be included in `bug` reports and `chat share`, enabling easier debugging and evaluation of sub-agent behavior.

4. **Evaluation and robustness infrastructure** — #24353 (76 behavioral evals across 6 models) tracks building a structured evaluation pipeline. Community interest in reproducible quality metrics.

5. **Destructive action safeguards** — #22672 and #23571 ask for prompt-level guardrails against dangerous git/DB operations and randomized tmp script creation. A safety and workspace hygiene concern.

---

## Developer Pain Points

- **Agent hangs and false success reporting** — Bugs #21409 (generalist hang), #22323 (max turns reported as success), and #25166 (shell wait-input) represent a cluster of reliability issues that erode trust in autonomous mode.

- **Subagent invocation problems** — #22093 (unexpected invocation), #22323 (false positives), and #21924 (flicker) show ongoing challenges with the agent delegation architecture.

- **Configuration ignored or overridden** — #11799 (`GEMINI.md` ignored), #22267 (browser agent ignores `settings.json`), and #24246 (tool limits) frustrate users who expect their configuration to be respected.

- **Terminal and shell integration issues** — #21924 (flicker on resize), #12045 (PowerShell crashes), #22465 (stuck at interactive prompts like Vite init) impair day-to-day usage across platforms.

- **Memory system instability** — #26516 tracker highlights indefinite retries, invalid patches, missing redaction, and silent failures. The Auto Memory feature is perceived as unreliable, especially for security-sensitive workflows.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest — 2026-07-28

## Today’s Highlights
A new patch release (v1.0.76-0) brings faster MCP tool loading and defaults autopilot to stay active after task completion. The community flagged a serious regression in plan‑mode that blocks shell commands, and a critical memory‑limit bug remains unfixed. Several new issues highlight gaps in ACP protocol support and macOS credential handling.

## Releases
**[v1.0.76-0](https://github.com/github/copilot-cli/releases/tag/v1.0.76-0)**
- **Improved:** MCP tools load faster from definition-scoped snapshots, with process‑wide and per‑server cache opt‑outs; autopilot stays selected after `task_complete` by default (set `stayInAutopilot: false` to return to interactive mode).
- **Fixed:** Restored early warning when un… (text truncated in source).

## Hot Issues
(Picked 10 by community impact, comment count, and severity. All links open in new tab.)

1. **[#4188 – [area:permissions, tools] Regression on plan‑mode](https://github.com/github/copilot-cli/issues/4188)**  
   Plan‑mode now blocks shell commands that were previously allowed (e.g., `gh` CLI). Users call this a regression that breaks workflow enrichment. (6 comments, 3 👍)

2. **[#2792 – [closed] Automatic switching between model for planning and execution](https://github.com/github/copilot-cli/issues/2792)**  
   Feature request to let users configure different models for planning vs. execution. Very popular (16 👍) – though closed, it signals strong demand.

3. **[#4163 – Zombie processes accumulate on Linux](https://github.com/github/copilot-cli/issues/4163)**  
   Finished subprocesses are not reaped, causing zombies (~2/min). Critical for long‑running sessions. (5 comments, 3 👍)

4. **[#4183 – Auto‑compaction does not prevent CAPI 5 MB failure](https://github.com/github/copilot-cli/issues/4183)**  
   Even with compaction, long tool‑heavy sessions hit the 5 MB serialization limit and stop responding. High‑severity. (4 comments, 10 👍)

5. **[#1381 – “Rewind is not available because you’re not in a git repository.”](https://github.com/github/copilot-cli/issues/1381)**  
   Users of non‑Git VCS (e.g., jj) cannot use Rewind. VS Code Copilot works without Git – CLI should follow suit. (3 comments, 9 👍)

6. **[#4233 – ACP mode lacks `usage_update` emission](https://github.com/github/copilot-cli/issues/4233)**  
   Clients like Zed cannot show context‑window or AI‑credit usage because `copilot --acp` never sends `usage_update`. (2 comments, 2 👍)

7. **[#4161 – `task_complete` tool unavailable after switching back to autopilot](https://github.com/github/copilot-cli/issues/4161)**  
   Regression of a previously fixed bug – the `task_complete` tool is again filtered out in some autopilot mode transitions. (2 comments, 3 👍)

8. **[#4118 – `/app` command does not select current working directory](https://github.com/github/copilot-cli/issues/4118)**  
   High demand (35 👍) – users must manually pick the directory every time they open the Copilot app via `/app`.

9. **[#4273 – macOS keychain prompts on every launch (XARA partition mismatch)](https://github.com/github/copilot-cli/issues/4273)**  
   GitHub‑signed and Microsoft‑signed Copilot binaries share login‑keychain items, causing repeated keychain prompts. (0 comments but new and impactful.)

10. **[#4271 – `glob` tool false‑negatives on multi‑segment patterns](https://github.com/github/copilot-cli/issues/4271)**  
    Any pattern with a path separator (e.g., `2026/07/*.md`) fails unless prefixed with `**/`. Breaks common file‑matching workflows.

## Key PR Progress
(Most open PRs are stale or spam; the following represent meaningful contributions.)

1. **[#1598 – fix: add trap to clean up temp directory on unexpected exit](https://github.com/github/copilot-cli/pull/1598)**  
   Prevents `/tmp` leaks when `install.sh` fails mid‑download.

2. **[#1116 – Fix misleading doc – 0x models don’t reduce quota](https://github.com/github/copilot-cli/pull/1116)**  
   Clarifies the README: 0x‑multiplier models do not consume quota per use.

3. **[#1609 – Update instructions for adding permissions in PAT](https://github.com/github/copilot-cli/pull/1609)**  
   Better navigation path for the “Copilot Requests” PAT permission, often missed.

4. **[#988 – chore(docs): add missing prefix to brew command](https://github.com/github/copilot-cli/pull/988)**  
   Fixes `brew install copilot-cli` → correct formula `github/copilot-cli/copilot-cli`.

5. **[#1333 – Fix minor grammar and Markdown formatting](https://github.com/github/copilot-cli/pull/1333)**  
   Adds missing article “an” and removes extra blank line.

6. **[#4030 – Add GitHub Actions workflow for Jekyll deployment](https://github.com/github/copilot-cli/pull/4030)**  
   Automates building and deploying a Jekyll site to GitHub Pages.

7. **[#3928 – Add .gitignore and settings configuration](https://github.com/github/copilot-cli/pull/3928)**  
   Adds project‑level `.gitignore` and editor settings.

8. **[#2800 – Add initial devcontainer configuration](https://github.com/github/copilot-cli/pull/2800)**  
   Provides a devcontainer for contributors.

*(Remaining PRs are either spam or very trivial – no other significant merging or review activity in the last 24h.)*

## Feature Request Trends
- **ACP protocol parity:** Multiple requests ([#4233](https://github.com/github/copilot-cli/issues/4233), [#4174](https://github.com/github/copilot-cli/issues/4174), [#4275](https://github.com/github/copilot-cli/issues/4275)) ask for exposing context usage, AI credits, and `contextTier` configuration in the ACP server.
- **Model flexibility:** Persistent demand for automatic model switching between planning/execution ([#2792](https://github.com/github/copilot-cli/issues/2792)) and better BYOK provider handling ([#4258](https://github.com/github/copilot-cli/issues/4258)).
- **Autopilot persistence:** Users want a launch flag or setting to keep autopilot mode across tasks ([#3977](https://github.com/github/copilot-cli/issues/3977), partially addressed in v1.0.76-0 by `stayInAutopilot`).
- **Non‑Git VCS support:** Rewind and other features that assume Git should work with alternative version control systems ([#1381](https://github.com/github/copilot-cli/issues/1381)).
- **Configuration enhancements:** Symlink support for `~/.copilot/` files ([#3264](https://github.com/github/copilot-cli/issues/3264)), and the ability to set context tier in ACP sessions.

## Developer Pain Points
- **Plan‑mode regression:** Blocking `gh` and other shell commands breaks existing workflows ([#4188](https://github.com/github/copilot-cli/issues/4188)).
- **Process management:** Zombie accumulation on Linux ([#4163](https://github.com/github/copilot-cli/issues/4163)) and missing cleanup on script failure.
- **Memory/context limits:** The 5 MB CAPI serialization cap stops sessions even when token context is healthy ([#4183](https://github.com/github/copilot-cli/issues/4183)).
- **Terminal rendering:** Blank screens on Windows Terminal ([#4263](https://github.com/github/copilot-cli/issues/4263), [#4159](https://github.com/github/copilot-cli/issues/4159)), “pending” message not cleared ([#4281](https://github.com/github/copilot-cli/issues/4281)), and clipboard failures in tmux/WSL ([#4191](https://github.com/github/copilot-cli/issues/4191)).
- **Tool misbehavior:** `glob` false negatives ([#4271](https://github.com/github/copilot-cli/issues/4271)), `task_complete` disappearing in autopilot ([#4161](https://github.com/github/copilot-cli/issues/4161)), and `subagent` OTEL spans missing billing attributes ([#4224](https://github.com/github/copilot-cli/issues/4224)).
- **Platform friction:** macOS keychain prompts due to signing mismatches ([#4273](https://github.com/github/copilot-cli/issues/4273)), and BYOK provider not accepting startup prompts in TTY mode ([#4258](https://github.com/github/copilot-cli/issues/4258)).

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest – 2026-07-28

## 📌 Today's Highlights

Community activity remains focused on fixing Windows encoding crashes and VSCode extension reliability. Two Unicode-related PRs target startup and web‑banner failures on non‑UTF‑8 consoles. Two VSCode extension bugs – unclickable file paths in Plan mode and stalled approval prompts – are generating user frustration. Meanwhile, a critical bug where `PostToolUse` hooks are silently garbage‑collected has been root‑caused, and a fix for MCP tool name normalization is open.

---

## 🚀 Releases

No new releases in the last 24 hours.

---

## 🔥 Hot Issues

[#1070 – Login failed: Network is unreachable](https://github.com/MoonshotAI/kimi-cli/issues/1070)  
*Closed, 8 comments, 👍 0*  
**Why it matters:** While closed, the issue saw recent activity, suggesting users still encounter sporadic connectivity to `auth.kimi.com:443`. The root cause remains environment‑specific (proxy, firewall, or DNS). Community has not yet confirmed a permanent workaround.

[#2317 – [VSCode Extension] Plan mode file path not clickable in chat webview](https://github.com/MoonshotAI/kimi-cli/issues/2317)  
*Open, 3 comments, 👍 0*  
**Why it matters:** A core workflow blocker – developers using Plan mode cannot navigate to referenced files via click. The extension version `0.5.10` is affected. No maintainer response yet.

[#2564 – fix(hooks): PostToolUse / PostToolUseFailure tasks collected by GC before completion](https://github.com/MoonshotAI/kimi-cli/issues/2564)  
*Open, 0 comments, 👍 0*  
**Why it matters:** A silent, non‑deterministic drop of hook subprocesses. Root cause identified as garbage collection of the `tokio::task::JoinHandle` in `kimi_cli/soul/toolse...`. This undermines reliability for users relying on `config.toml` hooks. No PR linked yet.

[#2563 – [Bug] VS Code extension: approval prompts (ExitPlanMode / tool permissions) intermittently never render, causing indefinite stalls or 600s timeout](https://github.com/MoonshotAI/kimi-cli/issues/2563)  
*Open, 0 comments, 👍 0*  
**Why it matters:** Another VSCode extension instability – approval prompts fail to appear, locking the developer for up to 10 minutes. Extension version `0.6.4` on macOS. No workaround documented.

---

## 🛠️ Key PR Progress

[#2539 – fix(mcp): normalize tools for Moonshot API](https://github.com/MoonshotAI/kimi-cli/pull/2539)  
*Open, created 2026-07-23, updated 2026-07-27*  
**What it does:** Generates stable Moonshot‑compatible aliases for MCP tool names while keeping original names for upstream routing. Fixes missing `object` type and corrects `anyOf`/`required` schema shapes. Critical for MCP tool interoperability.

[#2562 – fix(llm): allow disabling prompt cache key](https://github.com/MoonshotAI/kimi-cli/pull/2562)  
*Open, created/updated 2026-07-27*  
**What it does:** Adds a `prompt_cache_key` boolean to Kimi provider configuration. When `false`, the session‑derived `prompt_cache_key` field is omitted from requests. Maintains default behavior. Useful for users who want to bypass cache for debugging or consistency.

[#2561 – Fix UnicodeEncodeError on startup when stdio uses a non‑UTF-8 encoding](https://github.com/MoonshotAI/kimi-cli/pull/2561)  
*Open, created/updated 2026-07-27*  
**What it does:** Fixes `gbk` codec failure on the welcome banner character `▐` (from `ui/shell/__init__.py`). Resolves crash when launching `kimi` from Git Bash on Windows. Addresses issue #1436.

[#2560 – Fix UnicodeEncodeError in web banner when stdout is non‑UTF‑8 (Windows)](https://github.com/MoonshotAI/kimi-cli/pull/2560)  
*Open, created/updated 2026-07-27*  
**What it does:** Fixes `gbk` codec failure on the arrow character `➜` in the web server banner. Previously blocked `kimi web` from binding to a port on Chinese‑locale Windows. Addresses issue #2532.

---

## 💡 Feature Request Trends

No new feature requests were filed in the last 24 hours. The community’s attention is dominated by bug fixes rather than new capabilities. The PRs indicate a broader push toward:

- **Windows compatibility** – two Unicode‑related fixes.
- **MCP tooling stability** – normalization and schema corrections.
- **Control over caching** – the ability to disable prompt cache keys.

---

## ⚠️ Developer Pain Points

- **VSCode extension reliability** – two critical UX bugs (unclickable file paths, unresponsive approval prompts) remain open with no maintainer comment, causing workflow interruptions.
- **Hook execution failures** – `PostToolUse` hooks are silently dropped under GC, making automation unreliable.
- **Windows encoding crashes** – non‑UTF‑8 consoles (Git Bash, Command Prompt) trigger fatal errors on startup and web mode; the two pending PRs directly address this.
- **Connectivity issues** – the lingering #1070 (closed) suggests network configuration problems still plague a subset of users.

---

*Generated from GitHub data on 2026-07-28. All links open directly to the respective issues/PRs.*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest — 2026-07-28

## Today’s Highlights
Two patch releases (v1.18.7, v1.18.6) landed today, fixing macOS fullscreen insets, stale command palette entries, and branch‑specific cache corruption. The community flagged several regressions: models entering infinite tool‑call loops, a missing `task_id` on subagent failure, and unbounded `glob`/`grep` timeouts. A new PR to improve edit tool output and a docs PR clarifying subagent resume guidance both moved quickly to review.

## Releases

**[v1.18.7](https://github.com/anomalyco/opencode/releases/tag/v1.18.7)**  
- **Desktop** – Removed extra titlebar inset in fullscreen on macOS.  
- Fixed command palette entries reappearing after shadowed commands were removed.  
- Added scrolling to the project selector dropdown for long lists (thanks @david1gp).

**[v1.18.6](https://github.com/anomalyco/opencode/releases/tag/v1.18.6)**  
- **Core** – Fixed branch‑specific repository caches so refreshing one reference no longer moves another branch checkout.  
- **Desktop** – Improved compatibility with newer client API across directory, project, session, and terminal flows.  
- Fixed legacy MCP compatibility.

---

## Hot Issues (10 of 12 updated)

1. **[#28596 – Bug: repeated tool calls](https://github.com/anomalyco/opencode/issues/28596)**  
   *Model sometimes loops tool/exec calls indefinitely. Community has reproduced with various models; no fix yet. 5 comments, 0 👍.*

2. **[#38384 – Missing required parameter: 'input[8].arguments' at startup](https://github.com/anomalyco/opencode/issues/38384)**  
   *Persistent startup warning displayed in TUI – seems harmless but user reports no clue to root cause. 2 comments.*

3. **[#39196 – Foreground subagent failure returns no task_id](https://github.com/anomalyco/opencode/issues/39196)**  
   *When a foreground subagent fails, the parent gets a bare error string, leaving partial work stranded. Critical for agent orchestration. 2 comments.*

4. **[#39204 – deepseek-v4-flash-free stops agent loop after nearly every tool call](https://github.com/anomalyco/opencode/issues/39204)**  
   *Build agent halts after single tool call (read, grep, glob, todowrite) with free-tier DeepSeek model. Workaround: type “continue”. 1 comment.*

5. **[#35863 – Context window hardcoded to 200k for many models](https://github.com/anomalyco/opencode/issues/35863)**  
   *Static values cause premature auto‑compaction and overflow checks. Wanted: dynamic resolution from provider metadata. 1 comment, 1 👍.*

6. **[#39208 – v2 glob/grep run unbounded – no default timeout](https://github.com/anomalyco/opencode/issues/39208)**  
   *A single glob call hung for 21+ minutes. V1 had the same missing timeout. Marked [bug, 2.0]. 0 comments.*

7. **[#39207 – GitHub OAuth login fails with SQL error](https://github.com/anomalyco/opencode/issues/39207)**  
   *`email` param comes back empty from GitHub OAuth, breaking `update user` query. 0 comments.*

8. **[#39205 – Desktop: can only change theme once per settings page](https://github.com/anomalyco/opencode/issues/39205)**  
   *Theme dropdown works only on first open; subsequent changes require closing and reopening settings. 1 comment.*

9. **[#39210 – [needs:compliance] No response after prompt](https://github.com/anomalyco/opencode/issues/39210)**  
   *User reports the app accepts prompts but returns nothing – even after switching models. Language: Spanish. Likely compliance‑bot. 1 comment.*

10. **[#39212 – [FEATURE]: Clarify task_id source and resume guidance](https://github.com/anomalyco/opencode/issues/39212)**  
    *Task tool description says output includes a task_id, but actual output never provides it. Quickly followed by a docs PR. 0 comments.*

---

## Key PR Progress (10 of 50)

1. **[#39211 – feat(core): improve edit tool output](https://github.com/anomalyco/opencode/pull/39211)**  
   *Replaces synthetic diff preview with concise replacement‑count output, reports actual match count on ambiguous edits, adds target path on failure. Merged as [CLOSED].*

2. **[#39213 – docs(opencode): clarify task_id source and when to resume a subagent](https://github.com/anomalyco/opencode/pull/39213)**  
   *Addresses #39212 – updates task.txt prompt to note where task_id comes from and how to resume. [OPEN].*

3. **[#33453 – fix(provider): default custom models to image input](https://github.com/anomalyco/opencode/pull/33453)**  
   *Defaults new custom models to text+image input, preserves legacy attachment support. Merged.*

4. **[#39203 – refactor(core): manage watcher lifecycle with RcMap](https://github.com/anomalyco/opencode/pull/39203)**  
   *Makes watcher acquisition interrupt‑safe; pending Parcel subscription now respects 10‑second timeout. Merged.*

5. **[#39209 – fix(desktop): use channel database in local runs](https://github.com/anomalyco/opencode/pull/39209)**  
   *Stops disabling channel database for unpackaged desktop runs – improves local development consistency. [OPEN].*

6. **[#39206 – fix(desktop): make file:// chat links clickable](https://github.com/anomalyco/opencode/pull/39206)**  
   *Fixes DOMPurify stripping file:// links and missing event handler. Closes #37891. [OPEN].*

7. **[#29831 – fix(core): resolve spawn completion on exit, not only close](https://github.com/anomalyco/opencode/pull/29831)**  
   *Fixes Windows detached‑child hang – commands with background processes now complete properly. [OPEN].*

8. **[#38534 – feat(tui): emit toast mount event](https://github.com/anomalyco/opencode/pull/38534)**  
   *Adds `tui.toast.mount` lifecycle event for server plugins via POST. Closes #38527. [OPEN].*

9. **[#37625 – fix(provider): normalize kimi tool schemas for mfjs](https://github.com/anomalyco/opencode/pull/37625)**  
   *Projects tool schemas through compatibility layer to prevent incompatible custom/MCP tools from rejecting entire prompt. [OPEN].*

10. **[#38060 – fix(opencode): exclude denied MCP tools from provider requests](https://github.com/anomalyco/opencode/pull/38060)**  
    *Ensures global `tools` deny rules (e.g. `{"mymcp_*": false}`) actually exclude tools from API calls. Closes #37675. [OPEN].*

---

## Feature Request Trends

- **Task/Subagent clarity** – Multiple users (and the self‑correcting PR #39213) want clearer documentation about `task_id` availability and how to resume failed subagents.  
- **Localisation** – Issue #39202 requests translating “agent” to “智能体” (intelligent agent) instead of “代理” in Chinese docs.  
- **Dynamic context window** – Long‑standing #35863 asks for model‑specific context windows resolved from provider metadata rather than hardcoded 200k.  
- **Tool output transparency** – PR #39211 and related issues show demand for more actionable edit/schema failure messages.

---

## Developer Pain Points

1. **Infinite tool call loops** (#28596) – Models can enter a never‑ending loop of tool/exec calls, requiring manual kill.  
2. **Missing task_id on subagent failure** (#39196) – Parent agents cannot recover partial work from failed subagents, breaking multi‑step orchestration.  
3. **Unbounded glob/grep timeouts** (#39208) – V2 tools have no wall‑clock timeout; users waste minutes waiting for hung calls.  
4. **Hardcoded context window** (#35863) – Premature overflow and compaction for models with high actual context limits.  
5. **Startup parameter warnings** (#38384) – `Missing required parameter` message raises alarm but is harmless – lack of root cause insight.  
6. **Desktop theming limitation** (#39205) – Settings UX allows only one theme change per page open.  
7. **OAuth login failure** (#39207) – GitHub sign‑in broken by empty email field; requires database‑side fix.  
8. **Model‑specific stalls** (#39204) – DeepSeek free‑tier model stops after every tool call, forcing manual “continue” prompts.

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest – 2026-07-28

## Today's Highlights

A burst of 28 issues and 26 pull requests landed in the last 24 hours, with strong community engagement around extension API surfaces, provider compatibility, and caching performance. Notable is the rapid closure of several bugs—many filed and fixed within the same day—alongside continued work on search infrastructure (SQLite FTS5) and deeper model-scope control for extensions. The `anthropic-messages` path finally gets `x-client-request-id` session affinity, and a persistent crash in package schema validation has been addressed.

## Releases

No new releases in the last 24 hours.

## Hot Issues (10 noteworthy)

1. **[#5263 – Make in-session model and thinking-level changes ephemeral by default](https://github.com/earendil-works/pi/issues/5263)**  
   Proposes that runtime model/thinking changes affect only the active session, with a new "Default model" entry in `/settings`. 10 👍 and 10 comments show strong support; this would greatly simplify session management.

2. **[#6747 – API for enhancing agent message markdown](https://github.com/earendil-works/pi/issues/6747)**  
   Wants an extension hook to mutate agent message rendering (e.g., formula renderers) without changing the LLM content. 8 comments, 2 👍 – a clean separation-of-concerns feature.

3. **[#7157 – OpenCode Go provider displays as "OpenCode Zen Go"](https://github.com/earendil-works/pi/issues/7157)**  
   Simple display‑name bug in `pi --list-models`. Already fixed in PR #7173.

4. **[#7161 – anthropic-messages never sends x-client-request-id](https://github.com/earendil-works/pi/issues/7161)**  
   Missing header prevents gateways from grouping Anthropic conversations into sessions. 4 comments, immediate PR #7172.

5. **[#7190 – setCustomEditorComponent copies stale borderColor from defaultEditor](https://github.com/earendil-works/pi/issues/7190)**  
   A subtle bug where active editor’s border colour is ignored in favour of the default editor’s. Quick community reproduction.

6. **[#7196 – visibleWidth cache for large buffers: use LRU instead of FIFO](https://github.com/earendil-works/pi/issues/7196)**  
   FIFO eviction thrashes the cache, burning 15% CPU on large transcripts. LRU proposal with a PR linked.

7. **[#7194 – Pi does full re-render every 1s when tool card scrolls outside viewport](https://github.com/earendil-works/pi/issues/7194)**  
   Performance bug affecting remote sandbox users – wasteful redraws even when nothing changed.

8. **[#7187 – Silent crash caused by inconsistent error handling and schema validation](https://github.com/earendil-works/pi/issues/7187)**  
   A third‑party package manifest typo could permanently crash all sessions. Core package resolution crash before extensions run. Critical bug filed and closed within hours.

9. **[#7193 – Extension event-bus listeners survive session reloads and disposal](https://github.com/earendil-works/pi/issues/7193)**  
   Lifecycle leak in embedded Pi agent: listeners survive extension reload. Zero comments yet, but a serious isolation concern.

10. **[#7185 – Basic mouse support in the prompt input](https://github.com/earendil-works/pi/issues/7185)**  
    Simple UX request to click in the prompt to move cursor instead of arrow‑key scrolling. No comments, but mirrors popular demand.

## Key PR Progress (10 important)

1. **[#7022 – WIP: guard tree navigation during responses](https://github.com/earendil-works/pi/pull/7022)**  
   Blocks `/tree` while agent is streaming to prevent broken state. Work‑in‑progress but essential for interactive reliability.

2. **[#7163 – feat: search index sqlite](https://github.com/earendil-works/pi/pull/7163)**  
   Adds `SessionRepo.search()` with a contentless FTS5 virtual‑table migration for SQLite. Lay the groundwork for full‑text session search.

3. **[#7191 – feat(extensions): expose ctx.scopedModels to extensions](https://github.com/earendil-works/pi/pull/7191)**  
   Matches issue #7192; exposes the session’s resolved model set to extension contexts, enabling model‑picker companions.

4. **[#7081 – feat(ai): support Claude Opus 5 on Bedrock](https://github.com/earendil-works/pi/pull/7081)**  
   Configures adaptive thinking for Claude Opus 5 (required by the model) and improves error message sanitisation.

5. **[#7184 – fix(ai): strip multimodal media markers from tool results to prevent tokenizer crashes](https://github.com/earendil-works/pi/pull/7184)**  
   Fixes tokenizer crash when tool results contain orphaned `|image|` markers without actual bitmap data. Immediate crash fix.

6. **[#7176 – fix(ai): prefer configured Bedrock profile over ambient AWS keys](https://github.com/earendil-works/pi/pull/7176)**  
   Respects the Bedrock profile set via Pi’s auth flow even when `AWS_ACCESS_KEY_ID` environment variables are present.

7. **[#7174 – fix(ai): send max_tokens for Z.AI providers](https://github.com/earendil-works/pi/pull/7174)**  
   Z.AI ignores `max_completion_tokens` and only honours `max_tokens`, causing truncated turns. Adds `isZAI` flag to fix output cap.

8. **[#7172 – fix(ai): send x-client-request-id on anthropic-messages](https://github.com/earendil-works/pi/pull/7172)**  
   Adds session‑affinity header to Anthropic path, mirroring OpenAI paths. Fixes #7161.

9. **[#7169 – fix(coding-agent): dedupe byte-identical context files](https://github.com/earendil-works/pi/pull/7169)**  
   Dedupes `AGENTS.md`/`CLAUDE.md` by content hash instead of path, preventing duplicate context when worktrees mirror repo root.

10. **[#7103 – fix(coding-agent): support concurrent user bash cancellation](https://github.com/earendil-works/pi/pull/7103)**  
    Allows users to cancel a running bash command even when multiple bash tools are active. Approved and merged.

## Feature Request Trends

- **Extension API expansion**: Multiple requests for new context surfaces (`ctx.scopedModels`, message markdown hooks, terminal colour‑scheme API) – the community wants Pi to be a composable platform.
- **Session customisation and persistence**: Ephemeral model changes (#5263), better autocomplete settings persistence (#7179), and exposing the Responses API `phase` for richer UX (#7142) show a desire for fine‑grained session control.
- **Terminal UX improvements**: Mouse support in the prompt (#7185), deterministic input readiness after negotiation (#7177), and status lines for toggles (#7180) indicate that the TUI experience is a top priority.
- **Provider interoperability**: Requests for merging gateways (#5986), header alignment (`x-client-request-id`), and provider‑specific token parameters (Z.AI `max_tokens`) point to friction in multi‑provider setups.

## Developer Pain Points

- **Provider inconsistencies**: Headers (`x-client-request-id` missing on Anthropic, `max_tokens` vs `max_completion_tokens` on Z.AI) and display‑name mismatches (#7157) are recurring headaches.
- **Cache and performance regressions**: FIFO cache thrashing (#7196) and full‑transcript re‑renders (#7194) degrade the TUI experience, especially on large sessions.
- **Extension lifecycle bugs**: Listeners surviving session reloads (#7193), package‑install poisoning (#7189), and symlink‑resolving failures (#7195) make extension development fragile.
- **Windows compatibility**: Shift+Enter submission vs newline (#7175) remains a point of friction for Windows Terminal users.
- **Silent crashes from schema validation**: A single manifest typo taking down the entire agent (#7187) highlights the need for robust error boundaries in package loading.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code Community Digest – 2026-07-28

## Today’s Highlights
A new nightly release (v0.21.0‑nightly) shipped with a timezone fix for CLI insight measurements and a refactor of autofix extensions. Meanwhile, a non‑production benchmark prerelease (`dsw‑manual‑poc`) shows **376/500 SWE-bench Verified** tasks resolved, although the dataset is quarantined. On the development front, several Web Shell PRs landed—voice hold mode, a native folder picker, and a Git branch picker with commit/PR creation—while a main‑branch CI failure triggered an auto‑filed bug report.

## Releases
- **v0.21.0-nightly.20260727.c003e1718** – [Release](https://github.com/QwenLM/qwen-code/releases/tag/v0.21.0-nightly.20260727.c003e1718)  
  ✅ `fix(cli): measure insight days and hours in local time everywhere`  
  ✅ `refactor(autofix): ext`  
- **dsw-manual-poc-20260727-2** – Non‑production benchmark prerelease (based on `v0.20.0-nightly.20260722`). SWE‑bench Verified: 376 resolved, 116 unresolved, 1 execution error (dataset quarantined).

## Hot Issues (4 items updated in last 24h)
1. **[#7585] – Proposal: Add a direct external context provider profile** (P3, feature‑request)  
   *Author: @doudouOUC* – Requests an extension that allows an interactive CLI process to retrieve repository‑shared context from an external memory/knowledge service without modifying Qwen Core. Community discussion ongoing (9 comments).  
   → [Issue #7585](https://github.com/QwenLM/qwen-code/issues/7585)

2. **[#7449] – Proposal: Define an enterprise external‑memory integration profile** (P3, feature‑request)  
   *Author: @doudouOUC* – Proposes a provider‑neutral, documentation‑first profile for enterprise external memory integration. Builds on earlier triage feedback.  
   → [Issue #7449](https://github.com/QwenLM/qwen-code/issues/7449)

3. **[#7167] – Fleet Shepherd Dashboard** (need‑information, CI/CD)  
   *Author: @qwen-code-dev-bot* – Auto‑maintained dashboard tracking PR scan signals, syncs, and dispatches. Useful for monitoring CI health. Last tick: 2026-07-27.  
   → [Issue #7167](https://github.com/QwenLM/qwen-code/issues/7167)

4. **[#7861] – Main CI failed: E2E Tests** (bug, ready‑for‑agent)  
   *Author: @qwen-code-dev-bot* – A main‑branch E2E test workflow failed on commit `3209b89`. Auto‑labelled for agent remediation. Early discussion (2 comments).  
   → [Issue #7861](https://github.com/QwenLM/qwen-code/issues/7861)

## Key PR Progress (10 important PRs)
1. **[#7839] – feat(web-shell): honor voice hold mode**  
   Adds tap‑vs‑hold microphone interaction for Web Shell voice commands.  
   → [PR #7839](https://github.com/QwenLM/qwen-code/pull/7839)

2. **[#7414] – feat(triage): add revert‑pattern high‑risk path detection**  
   Replaces a low‑hit‑rate PR filter with a data‑backed triage gate based on revert history analysis (111 revert commits).  
   → [PR #7414](https://github.com/QwenLM/qwen-code/pull/7414)

3. **[#7856] – feat(web-shell): add composer footer renderer**  
   Provides an optional hook for hosts to render contextual content after the composer in chat/ split‑view.  
   → [PR #7856](https://github.com/QwenLM/qwen-code/pull/7856)

4. **[#7731] – feat(web-shell): add git branch picker, commit dialog, and create PR flow**  
   IntelliJ‑style branch picker with search, checkout, new branch, and commit/PR creation from the web shell.  
   → [PR #7731](https://github.com/QwenLM/qwen-code/pull/7731)

5. **[#7877] – feat(external-context): Add submitted‑prompt auto recall**  
   Opt‑in deterministic Auto Recall profile for the Direct External Context integration, running as a `UserPromptSubmit` hook.  
   → [PR #7877](https://github.com/QwenLM/qwen-code/pull/7877)

6. **[#7863] – fix(core): pass Grep pattern behind `-e` so a leading dash is not an option**  
   Prevents misinterpretation of patterns beginning with a dash as CLI flags.  
   → [PR #7863](https://github.com/QwenLM/qwen-code/pull/7863)

7. **[#7827] – fix(safe-mode): preserve caller‑supplied top‑tier MCP servers**  
   Ensures `--safe-mode` only drops ambient MCP servers, not those explicitly passed via `--mcp-config` or session API.  
   → [PR #7827](https://github.com/QwenLM/qwen-code/pull/7827)

8. **[#7859] – feat(web-shell): add native Live Voice**  
   macOS‑only opt‑in live voice experience: double‑tap Command to start conversation from any app.  
   → [PR #7859](https://github.com/QwenLM/qwen-code/pull/7859)

9. **[#7842] – fix(core): fast‑fail permanent quota‑exhaustion 429s**  
   Recognises 429 responses with reset‑time headers and fails immediately instead of silent retry.  
   → [PR #7842](https://github.com/QwenLM/qwen-code/pull/7842)

10. **[#7826] – feat(channels): dispatch GitHub notifications by reason**  
    Routes notifications (mentions, review requests, assignments) to appropriate agent inputs, improving automation accuracy.  
    → [PR #7826](https://github.com/QwenLM/qwen-code/pull/7826)

## Feature Request Trends
The two most active feature‑request issues this week both target **external context and memory integration**:
- A **direct external context provider profile** (#7585) to let CLI processes pull shared repository context from external services without core changes.
- An **enterprise external‑memory integration profile** (#7449) for provider‑neutral, documentation‑first support.  
PR #7877 already implements the first concrete “auto recall” profile, showing rapid community-driven implementation.

## Developer Pain Points
- **CI reliability**: The auto‑filed failure #7861 highlights ongoing E2E test instability on `main`. The Fleet Shepherd dashboard (#7167) provides visibility but does not prevent failures.
- **CLI edge cases**: The fix for Grep patterns with leading dashes (#7863) and the quota‑exhaustion 429 fast‑fail (#7842) address subtle but frustrating CLI/user‑experience bugs.
- **Safe mode complexity**: Users reported that `--safe-mode` inadvertently removed explicitly provided MCP servers; PR #7827 corrects this behaviour.
- **Benchmark data quality**: The SWE‑bench Verified dataset is quarantined, indicating unresolved issues with the benchmark infrastructure or result validation.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI Community Digest — 2026-07-28

## Today's Highlights

The v0.9.2 release candidate has entered active integration, with 82 commits batched via umbrella PR #4911. Community contributions continue to refine the TUI: a new `think_default_expanded` setting (#4928) addresses SSH/tmux usability, and a critical fix for avante.nvim compatibility lands (#4929). Meanwhile, issue #4930 highlights a blocker for shell-interactive workflows, and a large-scale zh-Hans translation update (#4908) improves localization quality.

## Releases

No new releases in the last 24 hours.

## Hot Issues

*Pick 5 noteworthy issues (all that were updated in the last 24h).*

1. **#4930 [OPEN] Enter during foreground shell should detach it before steering**  
   *Author: M-Maciej*  
   When the agent is blocked by a foreground Bash command, typing and hitting Enter fails silently. Users expect the shell to detach so their message can steer the agent. Practical UX sore point.  
   [GitHub](https://github.com/Hmbown/CodeWhale/issues/4930) | 👍 0 | Comments: 1

2. **#4925 [CLOSED] Add thinking_default_expanded setting**  
   *Author: M-Maciej*  
   Thinking blocks collapse by default, but the Space key to expand them is often intercepted over SSH/tmux. The accepted fix introduces a setting to always show reasoning expanded. Merged as PR #4928.  
   [GitHub](https://github.com/Hmbown/CodeWhale/issues/4925) | 👍 0 | Comments: 1

3. **#4907 [CLOSED] Web CI deploy trigger contradicts manual-only preflight**  
   *Author: Hmbown*  
   The Web Frontend workflow deterministically fails on main because the deploy step runs on push but the preflight is manual-only. Closed by trimming the trigger. Important for release pipeline reliability.  
   [GitHub](https://github.com/Hmbown/CodeWhale/issues/4907) | 👍 0 | Comments: 1

4. **#4751 [CLOSED] Settings IA rework: Fleet/Models boundaries**  
   *Author: Hmbown*  
   User screenshots showed misplaced toggles in the Settings UI. The Fleet section contained unrelated Goal-command and Workflow toggles, and a legacy row needed removal. Closed with a UI restructure.  
   [GitHub](https://github.com/Hmbown/CodeWhale/issues/4751) | 👍 0 | Comments: 1

5. **#4042 [CLOSED] Environment-level tool sandboxing for sub-agents**  
   *Author: JayBeest*  
   Long-running feature to enforce tool restrictions across sessions, sub-agents, Fleet workers, and MCP servers. Confirmed `--disallowed-tools` works. Community had 20 comments — high engagement.  
   [GitHub](https://github.com/Hmbown/CodeWhale/issues/4042) | 👍 0 | Comments: 20

## Key PR Progress

*Pick 10 important PRs from the 24 updated in the last 24h.*

1. **#4911 [CLOSED] v0.9.2 release candidate integration (umbrella)**  
   *Author: Hmbown*  
   Umbrella draft unifying 82 commits for the v0.9.2 RC. Opened to run CI and gather review on the combined candidate. All subsequent harvest PRs target this branch.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/4911)

2. **#4928 [CLOSED] feat(tui): add thinking_default_expanded setting**  
   *Author: M-Maciej*  
   Implements issue #4925. When enabled, thinking blocks start expanded. Space still toggles. Useful for SSH/tmux users.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/4928)

3. **#4929 [OPEN] fix(acp): preserve numeric JSON-RPC IDs for avante.nvim**  
   *Author: atmosuwiryo*  
   The ACP response helper coerced numeric IDs to strings (for Zed), which broke avante.nvim because Lua table keys distinguish `1` from `"1"`. Fix preserves the type by default with an opt-in string coercion.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/4929)

4. **#4932 [CLOSED] test(cli): satisfy strict all-target clippy**  
   *Author: Hmbown*  
   Post-v0.9.2 maintenance: replaces a `vec!` with a fixed-size array to pass `clippy::useless_vec` in a test. Ensures release gate passes.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/4932)

5. **#4931 [OPEN] Migrate QA PTY test harness from vt100 to rio-vt**  
   *Author: raphamorim*  
   Swaps the terminal emulator used in TUI test assertions from `vt100` to `rio-vt`. Promises better VT parsing; closes a long-standing maintenance gap.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/4931)

6. **#4913 [OPEN] test(preview): provider-free manifest×wire matrix**  
   *Author: Hmbown*  
   Adds a wiremock-based test matrix for four v0.9.2 benchmark routes (e.g., GLM-5.2, kimi-k3). No live calls — deterministic verification of request manifest against captured wire body.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/4913)

7. **#4912 [OPEN] feat(web): v0.9.2 docs guide/vocabulary, getting-started path**  
   *Author: Hmbown*  
   Harvests web docs onto the release candidate: new `/docs/guide` and `/docs/vocabulary` routes, homepage getting-started, a11y landmarks, and real-session media manifest.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/4912)

8. **#4908 [CLOSED] I18n(zh-Hans): update simplified-Chinese translations**  
   *Author: SparkofSpike*  
   Second round of zh-Hans translation improvements: adversarial review of all 1134 keys against `en.json`. A reviewer sub-agent independently verified the changes.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/4908)

9. **#4927 [CLOSED] fix(billing): dispatch-receipt classification, Moonshot/MiniMax product truth**  
   *Author: Hmbown*  
   Comprehensive billing fix cluster: a turn is billed from its dispatch receipt (not live config), Moonshot splits platform vs direct pricing, and route-scoped env URLs are enforced.  
   [GitHub](https://github.com/Hmbown/CodeWhale/pull/4927)

10. **#4923 [CLOSED] feat(tui): visual program slices — luminance audit, selection vocabulary, focus texture**  
    *Author: Hmbown*  
    Five visual-supervision slices: theme contrast audit (3:1 minimum for secondary chrome), menu selection vocabulary, focus texture, opt-in sound events, and a jellyfish animation fix.  
    [GitHub](https://github.com/Hmbown/CodeWhale/pull/4923)

## Feature Request Trends

- **Tool sandboxing across execution contexts**: Issue #4042 (now closed) reflects strong demand for granular tool access control per session, sub-agent, Fleet worker, and MCP server. The `--disallowed-tools` flag is confirmed, but environment-level enforcement remains a popular direction.
- **Always-expanded thinking blocks**: Issue #4925 / PR #4928 — users want reasoning blocks visible by default, especially when terminal capture interferes with toggling. This setting is now merged.
- **Remote/mobile/chat-bridge modes**: PR #4926 harvests an onboarding lane that includes a remote mode matrix, offline explore exit, and contributor skill onboarding. Suggests interest in multi-platform usage.
- **Persistent sessions and auto-resume**: PR #4922 implements persistent session rails, archived flags, and opt-in auto-resume. Users want to pick up where they left off without losing context.
- **Fleet configuration with exact routing**: PR #4924 introduces saved exact Fleets with provider/model pinning, permission ceilings, and a reasoning router. Shows desire for repeatable, safe multi-agent setups.

## Developer Pain Points

- **Foreground shell blocking**: Issue #4930 — when a Bash command blocks the turn, user input is ignored. No workaround for “steering” a blocked agent. High frustration in interactive usage.
- **SSH/tmux key capture**: Related to #4925 — Space key for expanding thinking blocks is intercepted by the terminal layer. The fix (new setting) is a workaround, but the underlying terminal compatibility issue remains.
- **CI pipeline fragility**: Issue #4907 — Web Frontend CI failed deterministically due to conflicting deploy triggers. Such reliability bugs waste developer time and reduce trust in the release process.
- **Settings information architecture**: Issue #4751 — users reported misplaced controls in the Settings UI. While fixed, this indicates that the settings layout is still maturing and may need further refinement as new features are added.
- **Lua/JSON-RPC type coercion**: PR #4929 — the ACP’s string coercion of numeric IDs broke avante.nvim. A reminder that cross-ecosystem compatibility (especially with Lua-based editors) requires careful type preservation.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*