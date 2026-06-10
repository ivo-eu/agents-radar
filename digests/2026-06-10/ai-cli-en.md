# AI CLI Tools Community Digest 2026-06-10

> Generated: 2026-06-10 05:08 UTC | Tools covered: 9

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
**Date:** 2026-06-10 | **Prepared by:** Senior Technical Analyst

---

## 1. Ecosystem Overview

The AI developer CLI tools landscape is experiencing a period of intense, parallel innovation across seven major projects, each iterating at different velocities and with distinct philosophical approaches. The dominant signal this week is the Claude Fable 5 model launch—triggering both capability excitement and significant cost/blocking issues that ripple across the ecosystem. A shared tension is emerging between rapid feature delivery (multi-agent orchestration, MCP integration, provider plurality) and the resulting quality regressions, particularly around session management, terminal rendering, and cross-platform reliability. The ecosystem is fragmenting along provider lines (Anthropic-centric vs. multi-provider) while converging on shared pain points around model costs, subagent reliability, and hook/plugin extensibility.

---

## 2. Activity Comparison

| Tool | Issues Updated (24h) | PRs Active (24h) | Release Status | Community Engagement |
|------|---------------------|------------------|----------------|---------------------|
| **Claude Code** | 10 hot issues | 9 notable PRs | **v2.1.170** (major: Fable 5) | High: 36-comment cost squeeze (#55053) |
| **OpenAI Codex** | 15 updated | 10 notable PRs | **3 releases** (rust-v0.139.0, 2 alphas) | Moderate: regional 404s, Desktop regressions |
| **Gemini CLI** | 10 of 44 tracked | 10 of 41 PRs | **v0.47.0-preview.0** + v0.46.0 stable | Moderate: PTY fix, security gating |
| **GitHub Copilot CLI** | 10 hot issues | 1 PR (spam) | **v1.0.61** patch | Moderate: 75👍 demand for classic commands |
| **Kimi Code** | 2 issues | 1 PR | No release today | **Low**: low-activity day |
| **OpenCode** | 10 picks | 10 picks | **v1.17.0** (file search, Cohere) | High: MCP timeout bugs, TUI crash |
| **Pi** | 10 hot issues | 10 PRs | **v0.79.1** (Fable 5 + trust system) | **Very High**: #5514 backlash (24 comments) |
| **CodeWhale** | 10 hot issues | 10 PRs | **v0.8.56** "Community Harvest" | High: 14 merged PRs, 9 contributors |
| **Qwen Code** | 8 tracked | 8 notable PRs | **v0.18.0-preview.1** (failed preview.2) | Moderate: OOM hardening, CI failures |

**Key Takeaway:** Claude Code and Pi lead in feature velocity with new model support; CodeWhale leads in community contribution volume; Kimi Code is the outlier with minimal activity.

---

## 3. Shared Feature Directions

Across 5+ tools, the following requirements appear with notable convergence:

| Requirement | Tools Expressing Need | Specifics |
|------------|----------------------|-----------|
| **AST-aware code navigation** | Claude Code, Gemini CLI, OpenCode | Replace naive grep/file-read with semantic code understanding, method boundary detection |
| **Multi-agent orchestration** | Claude Code (workflows), Gemini CLI (subagents), CodeWhale (#2007), Qwen Code (Agent Team) | Parallel sub-agents, role assignment, consensus/reconciliation |
| **Session/resume reliability** | **All tools** | Blank-screen-on-resume (Copilot), session window cost squeeze (Claude Code), model revert after compact (Claude Code), session file corruption (Gemini, OpenCode) |
| **MCP integration** | Gemini CLI, Copilot CLI, OpenCode, Pi, CodeWhale | Timeout handling (OpenCode), OAuth 2.1 support (CodeWhale #1409), custom registry URLs (Copilot) |
| **Terminal rendering fixes** | **All tools** | Resize handling (Qwen Code, Gemini CLI), mouse reporting (Claude Code, CodeWhale), scrollback corruption, viewport bugs |
| **Cost/usage transparency** | Claude Code (#55053, #66807), CodeWhale (#1607), Pi (context monitoring) | Real-time token budget, compression ratio, usage drain warnings |
| **Plugin/hook extensibility** | Copilot CLI, OpenCode, Claude Code, CodeWhale | Pre/post tool hooks, permission hooks, terminalSequence hook parity |
| **Globalization/i18n** | CodeWhale (7 locales), OpenCode (zh-CN request), Pi (CNY support) | Full localization of dialogs, sandbox elevation, token cost in local currencies |
| **Configurable trust/approval** | Pi (Project Trust #5514), Qwen Code (Plan Approval Gate #4853) | Per-folder/per-team trust gating with parent inheritance |
| **Provider plurality** | Gemini CLI, Qwen Code, CodeWhale, OpenCode | Identical model names across providers causing selection confusion (#4877 Qwen) |

---

## 4. Differentiation Analysis

| Dimension | Claude Code | OpenAI Codex | Gemini CLI | Copilot CLI | OpenCode | Pi | CodeWhale | Qwen Code |
|-----------|------------|--------------|------------|-------------|----------|-----|-----------|-----------|
| **Primary Model** | Fable 5 (Anthropic) | GPT-5.x (OpenAI) | Gemini 3.x (Google) | Multi-model (Copilot infra) | Multi-provider | Multi-provider | Multi-provider | Qwen/OpenAI-compat |
| **Philosophy** | Deep integration, opinionated | Refactoring for perf | Security-conscious, methodical | Enterprise-gated ecosystem | Plugins-first, extensible | Experimental, rapid iteration | Community-driven, localization-heavy | Stability & production readiness |
| **Target User** | Power users, early adopters | Cross-platform teams | Google Cloud ecosystem | GitHub/Enterprise orgs | Plugin developers | Tinkerers, local LLM users | Chinese/global bilingual | Qwen ecosystem, researchers |
| **Release Cadence** | Major + patch weekly | 3 releases/day | Preview + stable | Patch monthly | Weekly stable | Near-daily | Weekly community | Preview weekly |
| **Community Model** | Issues-driven, less PR engagement | Closed core, open plugin APIs | Open source, automated bots | Low OSS contribution | Open source, active | Open source, responsive | **Highest OSS contribution** | Open source, process-focused |
| **Biggest Risk** | Cost squeeze (#55053) | Platform regressions (Win/Mac) | Documentation drift | Stagnation (1 PR) | MCP timeout bugs | Trust/approval backlash | Windows compatibility (#765) | CI/CD fragility (#4864) |
| **Unique Strength** | Agent workflow reliability (when working) | Debug build perf (-78% async_trait PR) | Vertex AI integration, eval tooling | GitHub ecosystem lock-in | Plugin architecture | Model flexibility (local + cloud) | i18n, HarmanyOS port | OOM prevention, Agent Team |

---

## 5. Community Momentum & Maturity

### High Momentum / Rapid Iteration
- **Pi** — Most active single-day feature velocity (Fable 5 + Bedrock Mantle provider + prompt template defaults), but facing backlash from poorly-tested Project Trust feature. Community is vocal and maintainers respond within hours.
- **CodeWhale** — Strongest OSS contributor community (14 PRs from 9 contributors in one release). Deep global i18n investment and multi-agent orchestration vision. Windows blocker (#765) remains critical.
- **OpenCode** — Steady weekly releases with significant plugin architecture investment. MCP timeout bugs dominate issue tracker but PR pipeline is active.

### Established / Mature
- **Claude Code** — Highest engagement on individual issues (36 comments on cost squeeze). Model leadership (Fable 5) drives both excitement and pain. Slower PR acceptance rate; fewer external contributions.
- **Gemini CLI** — Most disciplined development process (stable + preview branches, systematic model promotion). PR activity is internally-driven (bots + Google engineers). Lower community participation volume.

### Plateau / Maintenance Phase
- **GitHub Copilot CLI** — Only 1 PR in 24h (spam). Community energy has shifted to forks (e.g., `shell-ai`, Issue #53). Enterprise-gated model access continues to frustrate users.
- **Kimi Code** — Lowest activity. Only 2 issues and 1 PR. File-reading loop bug (#640) has been open for 6 months.

### Mixed Signals
- **Qwen Code** — Strong architectural PRs (Agent Team, OOM hardening, prompt cache) but blocked by CI/CD process failures (#4864) and failed release pipeline (#4920). Technical quality is high; process maturity is catching up.

---

## 6. Trend Signals for Developers

### 1. **Cost-Aware Agent Design Is Becoming Non-Negotiable**
The #1 pain point across the ecosystem is cost unpredictability. Claude Code's 5-hour window squeeze (#55053) and OpenCode's image-processing usage drain (#66815) signal that developers need: per-turn token budgets, real-time context monitoring, and kill-switch mechanisms. Expect tools to add cost caps, compression ratio dashboards, and automatic model downgrades on cost thresholds.

### 2. **Multi-Provider Strategy Is Fragmenting the UX**
The proliferation of providers (Anthropic, OpenAI, Gemini, Bedrock, Together AI, local LLMs, MCP) is creating a new class of bugs: identical model names from different providers not being distinguishable (Qwen #4877), model switching mid-session breaking thinking (MiniMax M3 on Pi), and session state contamination across provider switches. Developers should demand provider-agnostic session management and model-version checkpointing.

### 3. **Sub-Agent Reliability Is the Next Frontier**
Every major tool is investing in multi-agent orchestration (Claude Code workflows, CodeWhale #2007, Qwen Agent Team, Gemini subagent eval). But reliability is consistently poor: false success after MAX_TURNS (Gemini #22323), background agents competing for resources (Pi #5035), and agent header omissions (Claude Code #66761). The maturity model is shifting from "does it work?" to "can I trust it not to silently fail?"

### 4. **Windows and Linux Are Becoming Second-Class Citizens**
The asymmetry is stark: macOS receives first-class treatment while Windows faces mapped drive history loss (Codex #27309), PTY file descriptor leaks (Claude Code), TUI crashes (OpenCode #31607), and clipboard/font/keyboard incompatibilities (Copilot CLI). For any team with heterogeneous environments, this is a risk factor in tool selection.

### 5. **Security and Trust Gating Is Tearing Communities Apart**
Pi's Project Trust backlash (#5514) and Qwen's Plan Approval Gate (#4853) reveal a fundamental tension: users want safety for destructive actions (git resets, database writes, file deletions) but reject friction that slows their flow. The winning approach will be configurable per-user/per-team trust policies with clear opt-out paths, not blanket gating.

### 6. **The "Claude Code Compatibility" Standard Is Emerging**
Multiple tools (CodeWhale #2865, OpenCode hooks, Pi prompts) are explicitly targeting parity with Claude Code's lifecycle model—hooks, skills, agents, prompt templates. This suggests that Claude Code's architecture is becoming the de facto reference design for AI CLI tools, similar to how VS Code's extension model became the standard for editor extensibility.

### **Recommendation for Developers:**
- **For stability:** Gemini CLI (disciplined releases) or OpenCode (active fix pipeline)
- **For bleeding edge:** Pi or Claude Code (Fable 5 access, but budget for costs)
- **For cross-platform teams:** CodeWhale (best i18n) or OpenCode (most platforms covered)
- **For enterprise:** Copilot CLI (GitHub integration) or Gemini CLI (GCP ecosystem)
- **To contribute to:** CodeWhale (most open to OSS contributors) or OpenCode (plugin architecture)

---

## Per-Tool Reports

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills Highlights

> Source: [anthropics/skills](https://github.com/anthropics/skills)

# Claude Code Skills Community Highlights Report  
*Analysis of `github.com/anthropics/skills` as of 2026-06-10*

## 1. Top Skills Ranking

The following pull requests attracted the most discussion and represent the community’s highest-priority skill contributions. All are currently **open**.

| Rank | PR | Skill | Functionality | Discussion Highlights | Status |
|------|----|-------|---------------|-----------------------|--------|
| 1 | [#514](https://github.com/anthropics/skills/pull/514) | **document-typography** | Prevents orphan words, widowed headers, and numbering misalignment in AI-generated documents. | Widely recognized pain point in Claude’s document output; strong support for a focused typography fix. Early adopters testing on long-form reports. | Open |
| 2 | [#486](https://github.com/anthropics/skills/pull/486) | **odt** | Full OpenDocument (.odt, .ods) creation, template filling, and conversion to HTML. | First native ODF support for LibreOffice users; community requested ISO-standard format handling. Conversations around edge cases (embedded images, styles). | Open |
| 3 | [#210](https://github.com/anthropics/skills/pull/210) | **frontend-design** (improvement) | Revises existing skill for clarity and actionability – each instruction must be executable in a single conversation. | Long-running discussion about balancing conciseness with guidance; multiple reviewers suggesting refinements to user‑facing examples. | Open |
| 4 | [#83](https://github.com/anthropics/skills/pull/83) | **skill-quality-analyzer** & **skill-security-analyzer** | Meta‑skills that evaluate other skills across structure, documentation, security, and output consistency. | First dedicated quality assurance tooling for the skills ecosystem; debate on scoring thresholds and false positives in security checks. | Open |
| 5 | [#181](https://github.com/anthropics/skills/pull/181) | **SAP-RPT-1-OSS predictor** | Integrates SAP’s open‑source tabular foundation model for predictive analytics on SAP business data. | Enterprise interest is high; discussion focuses on model loading requirements and data privacy when processing sensitive SAP records. | Open |
| 6 | [#1140](https://github.com/anthropics/skills/pull/1140) | **agent-creator** (meta‑skill) | Generates task‑specific agent sets from a single description. Includes fixes for multi‑tool evaluation and Windows path support. | Addresses long‑standing feature request (#1120); commenters testing meta‑skill composition with existing document and analysis skills. | Open |
| 7 | [#568](https://github.com/anthropics/skills/pull/568) | **servicenow** | Broad ServiceNow platform assistant covering ITSM, ITOM, SecOps, ITAM/SAM, FSM, SPM, CSDM, and IntegrationHub. | One of the largest skill scopes submitted; community caution about maintainability and token budget for such a wide skill. | Open |

*Other notable entries:* [#723 testing-patterns](https://github.com/anthropics/skills/pull/723) (testing trophy model, React, unit tests), [#154 shodh-memory](https://github.com/anthropics/skills/pull/154) (persistent context across conversations), [#335 masonry-generate-image-and-video](https://github.com/anthropics/skills/pull/335) (Imagen 3.0 integration).

---

## 2. Community Demand Trends (from Issues)

| Issue | Topic | Key Requests |
|-------|-------|--------------|
| [#228](https://github.com/anthropics/skills/issues/228) | Org‑wide skill sharing | Direct sharing links and shared skill libraries for teams; currently requires manual file transfer. |
| [#556](https://github.com/anthropics/skills/issues/556) | `run_eval.py` unreliability | Evaluation tool consistently reports 0% trigger rates; need robust testing infrastructure. |
| [#202](https://github.com/anthropics/skills/issues/202) | Skill‑creator skill overhaul | Request to rewrite as actionable instructions for Claude, not human documentation. |
| [#492](https://github.com/anthropics/skills/issues/492) | Namespace security | Community skills under `anthropic/` namespace risk trust‑boundary abuse; demand for clear provenance labeling. |
| [#189](https://github.com/anthropics/skills/issues/189) | Duplicate skills from plugins | `document-skills` and `example-skills` install identical content; need deduplication mechanism. |
| [#412](https://github.com/anthropics/skills/issues/412) | Agent governance skill | Proposal for audit trails, policy enforcement, and threat detection in multi‑agent setups. |
| [#16](https://github.com/anthropics/skills/issues/16) | MCP exposure | Demand to expose skill functionality as reusable MCP tools for broader programmatic use. |

**Key themes:**  
- **Enterprise tooling** (ODT, SAP, ServiceNow, SharePoint) and **meta‑skills** (quality, security, governance) dominate.  
- **Cross‑platform stability** (Windows subprocess bugs, path encoding) is a recurring pain point.  
- **Evaluation and testing** infrastructure needs urgent attention ([#556](https://github.com/anthropics/skills/issues/556), [#202](https://github.com/anthropics/skills/issues/202)).  

---

## 3. High‑Potential Pending Skills

These open PRs have active discussion and are structurally complete; they are likely to be merged soon or serve as templates for future work.

| PR | Skill | Reason for Near‑Term Impact |
|----|-------|------------------------------|
| [#514](https://github.com/anthropics/skills/pull/514) | **document-typography** | Simple, non‑invasive skill addressing a universal frustration; minimal scope reduces merge friction. |
| [#486](https://github.com/anthropics/skills/pull/486) | **odt** | First ISO‑standard document format support; strong community demand from LibreOffice/enterprise users. |
| [#1099](https://github.com/anthropics/skills/pull/1099) | **skill-creator: Windows fix** | Unblocks Windows users from using the evaluation loop – a critical infrastructure fix. |
| [#723](https://github.com/anthropics/skills/pull/723) | **testing-patterns** | Comprehensive testing skill covering modern practices (AAA, React Testing Library, mocking); likely to be merged as a reference skill for future PRs. |

---

## 4. Skills Ecosystem Insight

**One‑sentence summary:** The community’s most concentrated demand is for **enterprise‑grade document handling** (ODT, typography, PDF fixes, SharePoint) paired with **meta‑skills that assure quality, security, and governance** – revealing a shift from single‑task novelty to production‑ready reliability.

---

# Claude Code Community Digest — 2026-06-10

## Today's Highlights
Anthropic released **Claude Code v2.1.170**, introducing **Claude Fable 5** — a Mythos-class model now available generally. The community is buzzing about a sharp session-window cost squeeze (issue #55053) and a PTY file descriptor leak on macOS (#57580). Several high-severity bugs around tool-call text leakage and agent header omissions are also drawing attention as users begin stress-testing Fable 5.

## Releases
- **v2.1.170** ([release](https://github.com/anthropics/claude-code/releases/tag/v2.1.170))  
  - Introduces **Claude Fable 5**, a Mythos-class model with capabilities exceeding any previously released model.  
  - Fixed a session-related issue.  
  - Announcement: [Anthropic blog](https://www.anthropic.com/news/claude-fable-5-mythos-5)

## Hot Issues (10 noteworthy)
1. **#55053 – Sudden 5-hour session window squeeze** ([link](https://github.com/anthropics/claude-code/issues/55053))  
   *36 comments, 12 👍*  
   Users report the 5-hour session window depleting 5–10× faster since April 29, even for light editing. Sonnet sub-agents burn ~20–25% in under an hour. This is the top-cost complaint.

2. **#57580 – macOS PTY file descriptor leak** ([link](https://github.com/anthropics/claude-code/issues/57580))  
   *13 comments, 20 👍*  
   Long Bash-heavy sessions leak `/dev/ptmx` FDs, eventually exhausting `kern.tty.ptmx_max` and causing system-wide `ENXIO` errors. No official fix yet.

3. **#63870 – Bash tool calls emitted as raw `<invoke>` text** ([link](https://github.com/anthropics/claude-code/issues/63870))  
   *8 comments, 12 👍*  
   Detailed JSONL evidence of 23 malformed Bash calls in one session. Similar reports (#61122 etc.) suggest this is a recurring regression.

4. **#66761 – Workflow-tool agent() subagents omit x-claude-code-agent-id** ([link](https://github.com/anthropics/claude-code/issues/66761))  
   *6 comments, 6 👍*  
   Task/Agent-tool subagents get correct headers, but Workflow-tool subagents do not. Breaks agent tracking for complex workflows.

5. **#61889 – CVP user blocked on benign queries** ([link](https://github.com/anthropics/claude-code/issues/61889))  
   *30 comments, 3 👍*  
   Approved users on claude.ai (not Claude Code) being blocked for non-security queries. High frustration but may be a separate product issue.

6. **#62706 – Mouse reporting breaks copy/paste in SSH terminals** ([link](https://github.com/anthropics/claude-code/issues/62706))  
   *5 comments, 0 👍*  
   `tengu_pewter_brook` flag enables `ESC[?1000h`, making text selection impossible in MobaXterm. Workaround: `CLAUDE_CODE_DISABLE_MOUSE=1`.

7. **#66671 – Fable 5 safety measures blocking normal conversation** ([link](https://github.com/anthropics/claude-code/issues/66671))  
   *4 comments*  
   User simply says “Hi” and gets blocked. Indicates overly aggressive safety classifier for the new model.

8. **#65731 – /deep-research workflow hits rate limit errors** ([link](https://github.com/anthropics/claude-code/issues/65731))  
   *3 comments, 1 👍*  
   Default `/deep-research` workflow frequently fails with “Server is temporarily limiting requests.” Blocks research-heavy use.

9. **#66801 – Tool dispatch stall regression still in 2.1.168** ([link](https://github.com/anthropics/claude-code/issues/66801))  
   *2 comments*  
   Original #54847 auto-closed as stale; bug persists. Three independent reporters with cross-platform forensic data. No Anthropic engineer engagement.

10. **#66815 – Repeated image processing errors consume ~60% usage limit** ([link](https://github.com/anthropics/claude-code/issues/66815))  
    *2 comments*  
    19 consecutive “image could not be processed” failures, each re-sending full 34MB context. Cost impact is severe.

## Key PR Progress (9 notable)
1. **#66813 – Automated fix for bug bounty reporting issue** ([link](https://github.com/anthropics/claude-code/pull/66813))  
   REAPR bot attempts to fix #66806 by instructing users to report security issues via Anthropic’s responsible disclosure channel.

2. **#66608 – Fix false positive usage policy block on lattice gauge theory** ([link](https://github.com/anthropics/claude-code/pull/66608))  
   Automated fix for #66592 — Fable 5 incorrectly flagged academic physics content.

3. **#66607 – Fix Fable 5 safety classifier auto-switching to Opus** ([link](https://github.com/anthropics/claude-code/pull/66607))  
   Addresses #66595 where safety classifier downgrades model mid-session during authorized security testing.

4. **#66577 – Sync security-guidance plugin version and description** ([link](https://github.com/anthropics/claude-code/pull/66577))  
   Corrects `marketplace.json` to match plugin’s own `plugin.json` (version 2.0.0, updated description).

5. **#66575 – Use full author name in pr-review-toolkit plugin** ([link](https://github.com/anthropics/claude-code/pull/66575))  
   Changes “Daisy” to “Daisy Hollman” for consistency across repository plugins.

6. **#66573 – Restore dead error handlers broken by `set -euo pipefail`** ([link](https://github.com/anthropics/claude-code/pull/66573))  
   Fixes `ralph-wiggum` shell script where `set -e` caused early exits before error handling code could run.

7. **#66650 – Use full author name in pr-review-toolkit manifest** ([link](https://github.com/anthropics/claude-code/pull/66650))  
   Complementary to #66575, fixes the plugin manifest.

8. **#66572 – WIP: Fix repeated image processing errors consuming usage limit** ([link](https://github.com/anthropics/claude-code/pull/66572))  
   First attempt to address image processing failures causing usage drain. Work-in-progress.

9. **#65723 – Claude/subscription debate chat (unrelated)** ([link](https://github.com/anthropics/claude-code/pull/65723))  
   Ambiguous PR; likely not a viable contribution.

## Feature Request Trends
- **Configurable Spinner** (#62392) — users want to customize or disable the animated spinner and its verbs.
- **Context-Health Monitoring** (#66807) — expose real-time context usage, compression ratio, and token budget in the TUI.
- **Install Location Options** (#66794) — allow specifying a custom path during `claude install` (reopened from #21019).
- **Subtractive Exclude-List for Spinner Verbs** (#66814) — enable users to remove individual verbs from the spinner rotation without rewriting the entire list.

## Developer Pain Points
1. **Session cost squeeze** (#55053) — sudden acceleration of session window depletion is the #1 complaint.
2. **PTY file descriptor leak** (#57580) — macOS users risk system-wide terminal failure after long sessions.
3. **Tool call text leakage** (#63870, #66809) — Bash tool calls rendered as raw `<invoke>` text, breaking automation.
4. **Rate limits on /deep-research** (#65731) — blocks multi-agent research workflows.
5. **Safety classifier false positives** (#66671, #66592, #66607) — Fable 5 blocks benign content and model-switches mid-session.
6. **Mouse reporting breaking copy/paste** (#62706) — impacts SSH/WSL users.
7. **Image processing errors draining usage** (#66815) — repeated retries waste up to 60% of limit.
8. **Agent header omissions** (#66761) — Workflow subagents lack tracking IDs.
9. **Tool dispatch stall regression** (#66801) — unresolved for months, no Anthropic engagement.
10. **Model revert after compact** (#47325) — when multiple sessions use different models, compacting one session reverts to another’s model.

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

## OpenAI Codex Community Digest – 2026-06-10

A busy day on the Codex repo with three releases, a flurry of Windows and macOS regression reports, and an intensive push to remove `async_trait` from the core executor. The infrastructure team landed a series of refactors that improve debug build times by ~78%, while the community flagged several post-update regressions in the Desktop app.

---

### Releases

**rust-v0.139.0**  
- Code mode can now call standalone web search directly, including from nested JavaScript tool calls, and receive plaintext search results ([#26719](https://github.com/openai/codex/issues/26719)).  
- Tool and connector input schemas now preserve `oneOf` and `allOf`; large schemas keep more shallow structure when compacted.

**rust-v0.140.0-alpha.2** and **rust-v0.139.0-alpha.3**  
- Alpha releases, no detailed changelog provided.

[All releases](https://github.com/openai/codex/releases)

---

### Hot Issues (10 of 15 updated in last 24h)

| # | Issue | Why it matters | Community reaction |
|---|-------|----------------|--------------------|
| [#26916](https://github.com/openai/codex/issues/26916) | `gpt-5.5` returns 404 on first message (Brazil region) | Blocks CLI usage in Brazil; potentially a regional routing bug | 6 comments, no reactions yet |
| [#18115](https://github.com/openai/codex/issues/18115) | Repository-scoped marketplace & plugin config | Top feature request (38 👍); would let teams share plugins per repo | 5 comments, high demand |
| [#27205](https://github.com/openai/codex/issues/27205) | Encrypted parameters tool call error with `gpt-5.4` | Prevents use of encrypted tools in CLI; affects Pro users | 4 comments, 4 👍 |
| [#27297](https://github.com/openai/codex/issues/27297) | macOS Desktop update regression: `service_tier` default rejected & model selector missing | Breaks model selection after latest app update | 2 comments |
| [#27310](https://github.com/openai/codex/issues/27310) | Appshot fails with `completed_without_screenshot` on macOS | Screenshot capture pipeline fails silently | 1 comment |
| [#27309](https://github.com/openai/codex/issues/27309) | Windows Desktop: chat history on mapped network drives disappears after update | Data loss for network drive users | 1 comment |
| [#27308](https://github.com/openai/codex/issues/27308) | In-app browser using MSAL popup crashes Codex app | Blocks OAuth flows for apps using MSAL | 1 comment |
| [#27307](https://github.com/openai/codex/issues/27307) | Windows Desktop: mapped SMB/NAS project history lost after `\\?\UNC` canonicalization | Pervasive issue for NAS users | 1 comment |
| [#27296](https://github.com/openai/codex/issues/27296) | Fn global dictation hotkey stops working across apps after update | macOS system integration regression | 1 comment |
| [#27305](https://github.com/openai/codex/issues/27305) | `codex-cli 0.138.0` crashes with `zsh: trace trap` during code review | Hard crash on macOS, likely memory corruption | 0 comments, but critical |

---

### Key PR Progress (10 of 50 updated in last 24h)

| PR | What it does | Why it matters |
|----|--------------|----------------|
| [#27062](https://github.com/openai/codex/pull/27062) | Retry transient `Guardian` review failures | Improves reliability of auto-review (permission checks) |
| [#27304](https://github.com/openai/codex/pull/27304) | Remove `async_trait` from `ToolExecutor` — 78% debug build speedup | Major build performance improvement for core tests |
| [#27299](https://github.com/openai/codex/pull/27299) | Outline `ToolExecutor` handler bodies (prefactor) | Makes the `async_trait` removal reviewable step by step |
| [#27281](https://github.com/openai/codex/pull/27281) | Remove `EnvironmentPathRef` before filesystem migration | Simplifies path handling ahead of `PathUri` migration |
| [#27282](https://github.com/openai/codex/pull/27282) | Migrate `ExecutorFileSystem` to `PathUri` | Foundation for cross-environment file access |
| [#27312](https://github.com/openai/codex/pull/27312) | Reuse release artifacts for npm staging | Optimizes CI artifact caching, reducing release time |
| [#27311](https://github.com/openai/codex/pull/27311) | Skip local curated discovery for remote plugins | Fixes loading order; remote plugins now avoid unnecessary local lookup |
| [#26754](https://github.com/openai/codex/pull/26754) | Prepare side threads off TUI event loop to prevent deadlock | Fixes a deadlock when `/side` is used with slow forks |
| [#27069](https://github.com/openai/codex/pull/27069) | `fix(doctor)`: recognize legacy rollout files | Prevents `codex doctor` from marking old rollout files as corrupt |
| [#26259](https://github.com/openai/codex/pull/26259) | Add advisory `Interrupt` hooks for interrupted turns | New hook pattern for gather-and-report on user interruptions |

---

### Feature Request Trends

- **Repository-scoped plugin configuration** — Issue [#18115](https://github.com/openai/codex/issues/18115) (38 👍) leads demand; teams want to ship `.codex/config.toml` with plugin/marketplace settings that apply to all contributors.  
- **Better multi-platform reliability** — No single feature request, but the volume of Windows and macOS regression bugs suggests a strong desire for more rigorous update testing.

---

### Developer Pain Points

- **Windows Desktop regressions** — Four issues filed today alone: mapped drive history loss, MSAL popup crash, SMB canonicalization breakage, and sidebar disappearance. The platform feels underserved.  
- **macOS post-update breakage** — Hotkey, appshot, and model selector regressions from update `26.608.12217` suggest insufficient integration testing.  
- **CLI hard crashes** — `zsh: trace trap` ([#27305](https://github.com/openai/codex/issues/27305)) and the Brazil 404 ([#26916](https://github.com/openai/codex/issues/26916)) indicate fragile network handling.  
- **Encrypted tool configuration friction** — Model-gated validation errors ([#27205](https://github.com/openai/codex/issues/27205)) frustrate Pro users trying to use encryption features.

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

# Gemini CLI Community Digest — 2026-06-10

## 1. Today's Highlights

A busy release day brings **v0.47.0-preview.0** (the latest preview) and stable **v0.46.0**, which hardens PTY resize handling against native crashes. Several patch releases (v0.45.3, v0.46.0-preview.3) also land a critical Vertex AI model mapping fix. On the security front, a new PR gates chained E2E workflows to prevent fork-based API key exfiltration, reflecting growing awareness of supply-chain risks in CI.

---

## 2. Releases

| Version | Type | Key Changes |
|---------|------|-------------|
| **v0.47.0-preview.0** | Preview | Bumps to nightly base; adds "Respect backend def" config fix. [Full changelog](https://github.com/google-gemini/gemini-cli/pull/27644) |
| **v0.46.0** | Stable | Fixes PTY resize crashes via `harden PTY resize against native crashes` (PR #27496). Includes changelogs from v0.44.0 onward. [Full changelog](https://github.com/google-gemini/gemini-cli/pull/27495) |
| **v0.46.0-preview.3** | Preview Patch | Cherry-picks Vertex AI model mapping fix (f08b4af) into preview branch. [PR #27768](https://github.com/google-gemini/gemini-cli/pull/27768) |
| **v0.45.3** | Stable Patch | Same Vertex AI mapping fix backported to v0.45.2. [PR #27769](https://github.com/google-gemini/gemini-cli/pull/27769) |

*Takeaway:* The dual-patch strategy (preview + stable) for the Vertex AI model fix signals urgency—users on non-API-key auth flows were likely hitting routing failures with gemini-3.5-flash.

---

## 3. Hot Issues (10 of 44)

1. **[#24353 — Robust component-level evaluations](https://github.com/google-gemini/gemini-cli/issues/24353)** (P1, agent, 7 comments)  
   *Why it matters:* The epic tracks scaling behavioral evals from 76 to a sustainable baseline. With `aiq/eval_infra` label, this is foundational for quality gating all agent components.

2. **[#22745 — AST-aware file reads, search, codebase mapping](https://github.com/google-gemini/gemini-cli/issues/22745)** (P2, agent, 7 comments, 1 👍)  
   *Why it matters:* Proposes replacing naive grep/file-read with AST-aware tools to reduce token waste and misaligned reads. Community interest is high—this could meaningfully improve agent efficiency.

3. **[#22323 — Subagent false success after MAX_TURNS](https://github.com/google-gemini/gemini-cli/issues/22323)** (P1, bug, 6 comments, 2 👍)  
   *Why it matters:* `codebase_investigator` reports `status: "success"` and `Termination Reason: "GOAL"` even when it hit max turns before any analysis. This hides real failures from users. High upvote count shows frustration.

4. **[#21968 — Gemini does not use skills/sub-agents enough](https://github.com/google-gemini/gemini-cli/issues/21968)** (P2, bug, 6 comments)  
   *Why it matters:* Custom skills like "gradle" or "git" exist but Gemini rarely invokes them automatically. Undermines the entire extensibility value prop.

5. **[#26525 — Add deterministic redaction and reduce Auto Memory logging](https://github.com/google-gemini/gemini-cli/issues/26525)** (P2, security, 5 comments)  
   *Why it matters:* Auto Memory sends transcript content to the model *before* redaction prompts; secrets could leak. Logging also may expose skill data. Security-conscious users are watching.

6. **[#26522 — Stop Auto Memory from retrying low-signal sessions indefinitely](https://github.com/google-gemini/gemini-cli/issues/26522)** (P2, bug, 5 comments)  
   *Why it matters:* Sessions that look low-signal never get marked "processed," so extractor retries them forever. Wastes API quota and compute.

7. **[#25166 — Shell command hangs with "Waiting input" after completion](https://github.com/google-gemini/gemini-cli/issues/25166)** (P1, core, 4 comments, 3 👍)  
   *Why it matters:* Causes blocking UX for simple commands. Highest 👍 count in this batch—clearly a common pain point.

8. **[#21983 — Browser subagent fails on Wayland](https://github.com/google-gemini/gemini-cli/issues/21983)** (P1, bug, 4 comments, 1 👍)  
   *Why it matters:* Linux/Wayland users cannot use browser agent at all. P1 severity indicates this is a known blocker.

9. **[#27761 — Cloud Code Assist API 429 quota error for paid users](https://github.com/google-gemini/gemini-cli/issues/27761)** (P2, platform, 3 comments)  
   *Why it matters:* A user with Gemini Pro plan sees "Individual quota reached" despite 0% usage. Indicates a misconfiguration or backend bug immediately after plan purchase.

10. **[#24246 — 400 error with >128 tools](https://github.com/google-gemini/gemini-cli/issues/24246)** (P2, agent, 3 comments)  
    *Why it matters:* As users enable more extensions/MCP tools, the agent hits client-side tool limits. Needs dynamic tool pruning.

---

## 4. Key PR Progress (10 of 41)

1. **[#27780 — Security: gate chained E2E on same-repository checkout](https://github.com/google-gemini/gemini-cli/pull/27780)** (OPEN, size/xs)  
   *Why it matters:* Prevents fork PRs from exfiltrating `GEMINI_API_KEY` via `workflow_run` artifact injection. Proactive supply-chain hardening.

2. **[#27705 — Promote Gemini 3.1 Flash Lite to GA + support Gemini 3.5 Flash](https://github.com/google-gemini/gemini-cli/pull/27705)** (OPEN, size/xl)  
   *Why it matters:* Unifies three changes: retiring preview models, GA promotion, and 3.5 Flash support. A major model plumbing PR for the release branch.

3. **[#27465 — Surface extension disable/enable feedback to terminal](https://github.com/google-gemini/gemini-cli/pull/27465)** (CLOSED, P2, documentation/CLI)  
   *Why it matters:* `gemini extensions disable` previously produced zero output. Community frustration with silent commands is now fixed.

4. **[#27453 — Re-seed metadata when chat session file is recreated mid-session](https://github.com/google-gemini/gemini-cli/pull/27453)** (CLOSED, P2, agent)  
   *Why it matters:* External cleanup or tool removal could delete the session file mid-session, causing load failures. Now recovers gracefully.

5. **[#27455 — Amazon URL parsing and product metadata extraction](https://github.com/google-gemini/gemini-cli/pull/27455)** (CLOSED, P3, agent)  
   *Why it matters:* Adds `web-fetch` support for `amzn.in`/`amzn.to` short URLs and product data extraction. New e-commerce workflow capability.

6. **[#27643 — Fix parallel workspace compilation race condition](https://github.com/google-gemini/gemini-cli/pull/27643)** (OPEN, build)  
   *Why it matters:* Splits build into topological stages (core → libraries → apps) to prevent race conditions in monorepo builds. Improves CI reliability.

7. **[#27631 — Add static eval source analyzer](https://github.com/google-gemini/gemini-cli/pull/27631)** (OPEN, eval)  
   *Why it matters:* First piece of eval dev tooling—parses TypeScript AST to extract helper names, base hashes, and metadata. Will help diagnose failed evals faster.

8. **[#23948 — Fix infinite re-render loop in useFlickerDetector/useSessionResume](https://github.com/google-gemini/gemini-cli/pull/23948)** (CLOSED, P0, core)  
   *Why it matters:* Critical UI lockup fix where `useEffect` hooks missing dependency arrays caused freezing. Now closed and merged.

9. **[#27772 — Standardize tool output formatting](https://github.com/google-gemini/gemini-cli/pull/27772)** (OPEN, core)  
   *Why it matters:* Introduces `wrapUntrusted` helper to unify MCP, shell, and web-fetch output formats. Reduces parsing bugs downstream.

10. **[#27770 — Avoid persisting empty resume sessions](https://github.com/google-gemini/gemini-cli/pull/27770)** (CLOSED, agent)  
    *Why it matters:* Prevents empty or command-only sessions from appearing in `/resume` or `--list-sessions`. Clean UX improvement.

---

## 5. Feature Request Trends

1. **AST-aware tooling for agent efficiency** (Issues #22745, #22746, #22747)  
   A strong push to replace naive grep/file-read with AST-based code understanding—method boundary detection, semantic search, and smarter codebase mapping. The goal is fewer turns and less token waste.

2. **Subagent orchestration improvements** (Issues #22601, #22600, #22741)  
   Requests for: (a) dedicated subagent eval strategy, (b) dev+review loop subagents for code quality, and (c) ability to background subagents with Ctrl+B. The community wants subagents to be first-class parallel workers, not blocking callers.

3. **Agent self-awareness and self-documentation** (Issue #21432)  
   Users want Gemini CLI to know its own hotkeys, CLI flags, and settings—so it can answer "How do I run myself for X?" without external docs.

4. **Destructive action guardrails** (Issue #22672)  
   Users want the agent to flag or block unsafe git operations (`git reset --hard`, `--force` pushes) and database modifications unless explicitly confirmed.

5. **Memory system reliability** (Issues #26516, #26522, #26523, #26525)  
   A cluster of Auto Memory bugs: indefinite retries, silent patch invalidation, and pre-redaction secret exposure. The community wants deterministic extraction, quarantine for invalid patches, and reduced logging.

---

## 6. Developer Pain Points

- **Subagent false positives and ignore of settings** — Issues #22323 (false success after MAX_TURNS), #22093 (subagents running despite being disabled), and #22267 (browser agent ignores `settings.json` overrides). Confidence in agent behavior is low.
- **Shell execution hangs** — Issue #25166: simple commands hang with "Waiting input" after completion. High 👍 count suggests this is a daily frustration.
- **Terminal/UI flicker and corruption** — Issues #21924 (high performance on resize), #23948 (infinite re-render lockup), #24935 (corruption after exiting external editors). Terminal UX fragility is a recurring theme.
- **Quota and billing confusion** — Issue #27761: paid Pro users hitting 429 quota errors with 0% usage. Suggests backend provisioning or model routing issues immediately after plan purchase.
- **Documentation drift** — Issue #27449 (policy engine docs showing wrong tier numbers) and the general need for changelog automation (multiple PRs from `gemini-cli-robot`) indicate documentation is hard to keep current.
- **Inconsistent tool limits** — Issue #24246: >128 tools causes 400 errors. With more MCP extensions being added, this limit will bite more users unless dynamic pruning is implemented.

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI Community Digest – 2026-06-10

## Today’s Highlights
A new patch release (v1.0.61) landed with a polished agents picker, a dedicated `/settings` dialog, and a fix for blank-screen-on-resume. Community attention remains focused on the long-running **Issue #53** (75 👍) – the request to restore classic Copilot CLI commands – which has driven users to create their own forks. Several regressions in MCP and plugin hooks are also drawing scrutiny, alongside continued model availability gaps for enterprise users.

## Releases
**v1.0.61** – 2026-06-09
- Polish `/agents` picker and “Create New Agent” wizard with consistent borders, headers, and styled inputs.
- Fixed a bug where resuming a session could leave the screen blank.
- Added `/settings` interactive dialog to browse and edit all user settings in one place.
- Improved local session resuming with better state handling.

[View release](https://github.com/github/copilot-cli/releases/tag/v1.0.61)

## Hot Issues (Top 10 by community activity & impact)

1. **[Issue #53 – Bring back GitHub Copilot in the CLI commands](https://github.com/github/copilot-cli/issues/53)**  
   After six months without official word, the community is building alternatives (e.g., `shell-ai`). 75 👍, 31 comments. A critical workflow-breaking regression for many power users.

2. **[Issue #1703 – Copilot CLI does not list all org-enabled models (e.g. Gemini 3.1 Pro)](https://github.com/github/copilot-cli/issues/1703)**  
   While VS Code Copilot sees the full model list, CLI is stuck with a subset, even when models are enabled in org settings. 54 👍, 29 comments. Blocks enterprise adoption.

3. **[Issue #2082 – ctrl+shift+c no longer copies to clipboard on Linux](https://github.com/github/copilot-cli/issues/2082)**  
   A basic terminal shortcut broken since v1.0.4. 8 👍, 20 comments. Affects daily usability on Ubuntu.

4. **[Issue #135 – Light theme doesn’t work](https://github.com/github/copilot-cli/issues/135)**  
   Long-standing cosmetic issue with no fix yet. 11 👍, 11 comments. Minor but irritating for light-terminal users.

5. **[Issue #2050 – Claude Sonnet 4.6 – Execution failed: Error 503 / connection errors](https://github.com/github/copilot-cli/issues/2050)**  
   Repeated failures with `claude-sonnet-4.6` (medium) while Gemini works fine. 4 👍, 8 comments. Affects heavy spec generation tasks.

6. **[Issue #3436 – /mcp search constructs wrong URL for custom MCP registries (CLOSED)](https://github.com/github/copilot-cli/issues/3436)**  
   Missing `/v0.1/` segment caused 404’s against self-hosted registries. Fixed in a recent release. 1 👍, 7 comments. Important for enterprise MCP setups.

7. **[Issue #2540 – Plugin-defined preToolUse hooks do not fire in main session or subagents](https://github.com/github/copilot-cli/issues/2540)**  
   `hooks.json` preToolUse hooks are silently ignored. 3 👍, 7 comments. Blocks plugin author workflows.

8. **[Issue #3596 – Error loading model list: Not authenticated when resuming a session](https://github.com/github/copilot-cli/issues/3596)**  
   Resuming a specific session yields authentication errors, while a fresh session works. 10 👍, 3 comments. Disrupts session continuity.

9. **[Issue #2243 – Worktrees are a nightmare, should be disabled by default](https://github.com/github/copilot-cli/issues/2243)**  
   Auto-created worktrees cause thousands of unmergeable lines. 8 👍, 2 comments. User frustration with uncontrolled side effects.

10. **[Issue #1613 – Feature request: Built-in git worktree lifecycle management](https://github.com/github/copilot-cli/issues/1613)**  
    The community wants Copilot CLI to create/destroy worktrees safely as part of its workflow. 31 👍, 2 comments. A complementary ask to #2243.

## Key PR Progress
Only one pull request was opened in the last 24 hours:

- **[PR #3737 – Jigg empire ai](https://github.com/github/copilot-cli/pull/3737)**  
  A low-quality, likely spam PR with no meaningful description. Not merged. No other significant PR activity – the repository appears to be in a maintenance phase with few open source contributions landing this cycle.

## Feature Request Trends
- **Git worktree lifecycle management** (#1613, 31 👍) – users want safe, automated worktree creation/cleanup.
- **Local session sharing across machines** (#3729) – ability to continue a session on another machine.
- **Enterprise custom model support** (#3730) – bring admin-configured models into CLI.
- **Restore `web_fetch` access to private networks** (#3731) – security restrictions in v1.0.60 broke internal corporate agentic files.
- **Classic CLI command restoration** (#53) – still the top community ask; forks are emerging.

## Developer Pain Points
- **Model list inconsistency** – CLI doesn’t show all models available in VS Code, blocking org users.
- **Clipboard & input handling** – `ctrl+shift+c` broken on Linux, Chinese character double-encoding, Ctrl+wheel zoom blocked on Windows.
- **Session bugs** – blank screen on resume, authentication failures when resuming sessions, session store corruption (`cwd`/`branch` null since v1.0.13).
- **Plugin & MCP regressions** – hooks not firing, MCP server spawning loops, missing URL segments in custom registries, OAuth fan-out causing rate limits.
- **Non-ASCII & encoding issues** – Bash tool strips non-ASCII characters, `edit` tool corrupts non-UTF-8 bytes.
- **Worktree chaos** – auto-created worktrees produce unmergeable code, lack of lifecycle management.
- **Windows installation & uninstallation** – cannot uninstall via Control Panel, Ctrl+G fails with `code-insiders --wait`.

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI Community Digest — 2026-06-10

## Today's Highlights
Only one pull request and two issues were updated in the last 24 hours, indicating a low-activity day for the Kimi CLI project. The community’s attention is focused on a longstanding file-reading loop bug (Issue #640) and a recent edit-tool failure in Kimi Code (Issue #2443). The sole PR (#2445) improves LLM context by surfacing stderr from the `PostToolUse` hook, which could help debug tool interactions.

## Releases
No new releases in the last 24 hours.

## Hot Issues
*(Only 2 issues were updated in the last 24 hours; both are listed below)*

1. **#640 – [bug] Kimi CLI stuck in reading one file again and again and stuck in a loop**  
   *Author: isbafatima90-arch* | Created: 2026-01-19 | Updated: 2026-06-10  
   *Comments: 7* | 👍: 1  
   **Summary:** User on Linux reports that Kimi CLI v0.76 enters an infinite loop repeatedly reading the same file. The issue is using a custom Anthropic endpoint (`mimo-v2-flash`) via `config.toml`.  
   **Significance:** This is a critical, long-standing bug (6+ months old) that affects users relying on custom endpoints. The recent update suggests maintainers may be looking into it.  
   [GitHub](https://github.com/MoonshotAI/kimi-cli/issues/640)

2. **#2443 – [bug] Edit tool keeps failing in new kimi-code**  
   *Author: iaindooley* | Created: 2026-06-09 | Updated: 2026-06-09  
   *Comments: 0* | 👍: 0  
   **Summary:** User on Debian reports that the edit tool frequently fails in Kimi Code v0.12.0, running with the `k2.6` model after login.  
   **Significance:** Zero comments suggests the issue is new and may still be unconfirmed. However, “fails frequently” implies a regression in the edit pipeline, which is core to the CLI’s value proposition.  
   [GitHub](https://github.com/MoonshotAI/kimi-cli/issues/2443)

## Key PR Progress
*(Only 1 PR was updated in the last 24 hours)*

1. **#2445 – feat(hooks): surface PostToolUse hook stderr to LLM context**  
   *Author: zwpdbh* | Created: 2026-06-10 | Updated: 2026-06-10  
   *Comments: 0* | 👍: 0  
   **Summary:** Converts the `PostToolUse` hook from fire-and-forget to awaited, so that stderr output from the hook is captured and appended to the tool result message. This gives the LLM visibility into hook-side errors or logs.  
   **Significance:** A small but valuable improvement for debugging custom hooks. It enables the model to see errors that would otherwise be silently dropped, which is especially useful for users integrating custom toolchains.  
   [GitHub](https://github.com/MoonshotAI/kimi-cli/pull/2445)

## Feature Request Trends
Given the limited data (only 2 issues and 1 PR), explicit feature requests are sparse. However, the following directions can be inferred:
- **Better error recovery for file reading**: Issue #640 implies a need for loop detection or timeouts when reading files, especially with non-standard models/endpoints.
- **Robustness of edit tool**: Issue #2443 suggests users want the edit tool to be more reliable, possibly with clearer error messages or fallback mechanisms.
- **Hook improvements**: PR #2445 indicates interest in tighter integration between custom hooks and the LLM context, hinting at a broader desire for extensibility.

## Developer Pain Points
- **File-reading infinite loops**: The oldest unresolved bug (#640) highlights that custom endpoint configurations can lead to unrecoverable loops, causing frustration and lost work.
- **Edit tool failures in new releases**: The recent v0.12.0 appears to have introduced regressions in the edit functionality, a core workflow for code generation.
- **Lack of community engagement on new bugs**: Issue #2443 has zero comments despite being open for a day, suggesting either low maintainer bandwidth or difficulty reproducing the issue.
- **Dependency on platform-specific models**: Both issues mention custom endpoints or specific models, indicating that users are experimenting beyond the default setup and encountering inconsistencies.

*Generated from GitHub data for MoonshotAI/kimi-cli on 2026-06-10.*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode Community Digest – 2026-06-10

## Today's Highlights
Version **1.17.0** shipped with a faster `fff`-backed file search and Cohere North support, while the community tackled a flurry of regressions—session title auto-generation broke for opencode‑provider models, and a new crash hit TUI startup on Windows. MCP timeout bugs continue to dominate the issue tracker, but two key PRs landed fixes for both Anthropic fallback responses and MCP client failure safety.

## Releases
**v1.17.0**  
[View release](https://github.com/anomalyco/opencode/releases/tag/v1.17.0)  
- **Faster file search** across large projects using the `fff`-backed tools (@dmtrKovalenko)  
- **`X-Session-Id` headers** for proxy setups that need sticky routing (@songchaow)  
- **Cohere North** model support  
- **`reasoning` as an interleaved field** option for provider messages

---

## Hot Issues (10 picks)

1. **[#30662 – Auto session title generation fails for opencode provider models](https://github.com/anomalyco/opencode/issues/30662)**  
   *7 comments* – Titles stay as “New session…” when using opencode‑provider models (e.g., `big-pickle`). Root cause: missing provider config in the title agent’s `smallOptions`. Community suspects a config merge bug.  

2. **[#31337 – The status endpoint is unusable in server mode](https://github.com/anomalyco/opencode/issues/31337)**  
   *5 comments* – `GET /session/status` returns 404 on v1.15.12 (macOS). Blocks SDK clients that rely on it for session monitoring.

3. **[#31437 – [FEATURE] add `/bottom` in CLI to scroll terminal to bottom](https://github.com/anomalyco/opencode/issues/31437)**  
   *👍2* – Already partially implemented? Users want an explicit command rather than manual scrolling. Quick community engagement.

4. **[#26412 – Custom OpenAI-compatible provider: “Expected 'function.name' to be a string” on streaming tool calls](https://github.com/anomalyco/opencode/issues/26412)**  
   *4 comments* – vLLM backend users hit validation errors on streaming tool‑call chunks. Impacts self‑hosted deployments.

5. **[#31579 – `@ai-sdk/anthropic` 3.0.71 rejects `usage.iterations[].type: "fallback_message"`](https://github.com/anomalyco/opencode/issues/31579)**  
   *2 comments, closed* – Anthropic’s new fallback API emits a `fallback_message` type that the bundled SDK disallows, breaking streams. Fixed in PR #31611.

6. **[#31588 – Tool stderr leaks into message input field on bash timeout](https://github.com/anomalyco/opencode/issues/31588)**  
   *2 comments, closed* – Multi‑line stderr after a timed‑out `sudo` command pollutes the user’s input buffer. Workaround: spam spacebar.

7. **[#31592 – Session title no longer auto updates](https://github.com/anomalyco/opencode/issues/31592)**  
   *2 comments* – Regression after v1.16.2; titles stay static. Possibly related to #30662.

8. **[#31607 – Launch opencode 1.17.00 TUI may cause crash](https://github.com/anomalyco/opencode/issues/31607)**  
   *1 comment* – Segmentation fault on Windows (Bun standalone). Critical blocker for Windows users.

9. **[#31606 – Switching model mid-session causes SQLiteError: NOT NULL constraint failed: session_message.seq](https://github.com/anomalyco/opencode/issues/31606)**  
   *1 comment* – All subsequent messages fail after a model switch. Makes sessions unusable until restart.

10. **[#31604 – Read tool displays garbled text for GBK encoded files](https://github.com/anomalyco/opencode/issues/31604)**  
    *1 comment* – Chinese Windows users see mojibake. Common in .NET projects.

---

## Key PR Progress (10 picks)

1. **[#31618 – fix(mcp): apply timeouts to catalog requests](https://github.com/anomalyco/opencode/pull/31618)**  
   Centralizes timeout precedence for MCP `prompts/list` and `resources/list` pagination. Preserves runtime override + experimental fallback.

2. **[#31611 – fix(opencode): support Anthropic fallback responses](https://github.com/anomalyco/opencode/pull/31611)**  
   Updates `@ai-sdk/anthropic` to 3.0.82 and accepts `fallback_message` iteration types. Closes #31579.

3. **[#31613 – fix(session): preserve model message cache snapshots](https://github.com/anomalyco/opencode/pull/31613)**  
   Prevents mutable DB message objects from corrupting cache. Closes #25366 and #24841.

4. **[#30509 – feat(permission): wire permission.ask plugin hook](https://github.com/anomalyco/opencode/pull/30509)**  
   Adds `Plugin.trigger` before permission prompts, enabling plugins to intercept and approve/deny requests.

5. **[#30508 – fix(permission): prevent doom_loop infinite popups](https://github.com/anomalyco/opencode/pull/30508)**  
   Catches `PermissionV1.RejectedError` to stop the loop when a permission is denied.

6. **[#31602 – fix(tui): clear leader sequence when typing in prompt input](https://github.com/anomalyco/opencode/pull/31602)**  
   Prevents leader key shortcuts (e.g., `l` for /login) from being triggered while typing. Fixes #31568.

7. **[#31595 – fix(mcp): make client creation failure-safe](https://github.com/anomalyco/opencode/pull/31595)**  
   Wraps MCP client initialization in a total boundary; returns explicit failure statuses instead of crashing.

8. **[#31609 – fix(opencode): start git HEAD watcher at instance bootstrap](https://github.com/anomalyco/opencode/pull/31609)**  
   Restores live branch name updates in the TUI footer after `git checkout` outside OpenCode. Closes #30956.

9. **[#31601 – feat(core): add project reference guidance](https://github.com/anomalyco/opencode/pull/31601)**  
   Describes local/Git project references, advertises them to agents, and auto‑allows external directories.

10. **[#31578 – fix(cli): stream run output, add empty-text warning, flush race-late parts](https://github.com/anomalyco/opencode/pull/31578)**  
    Fixes three silent‑exit bugs in `opencode run`: streaming text, empty warnings, and out‑of‑order parts.

---

## Feature Request Trends
- **New provider integrations** – oMLX (Apple MLX‑native inference) requests appeared twice ([#31587](https://github.com/anomalyco/opencode/issues/31587), [#31603](https://github.com/anomalyco/opencode/issues/31603)), plus Cohere North support just landed in v1.17.0.
- **CLI and TUI enhancements** – `/bottom` scroll command ([#31437](https://github.com/anomalyco/opencode/issues/31437)), Chinese (zh‑CN) TUI localization ([#31585](https://github.com/anomalyco/opencode/issues/31585)), and dynamic Tmux window names ([#24822](https://github.com/anomalyco/opencode/issues/24822)).
- **Platform‑specific support** – x86_64‑v1 CPU compatibility ([#31605](https://github.com/anomalyco/opencode/issues/31605)) and improved Windows experience (session export [#19513](https://github.com/anomalyco/opencode/issues/19513)).
- **Skills and metadata** – Users request better documentation and actual usage of skill metadata ([#31616](https://github.com/anomalyco/opencode/issues/31616)).

---

## Developer Pain Points
- **MCP timeout / progress notification failures** – Multiple reports (#28186, #14499, #24965) where long‑running tool calls time out because `onprogress` is not wired or progress notifications are ignored by the MCP client.
- **Session title regressions** – Both opencode‑provider models (#30662) and the latest update (#31592) broke auto‑title generation, frustrating users who rely on context‑aware naming.
- **Windows‑specific bugs** – Inconsistent path separators in SQLite break session listing (#31597); TUI crash on v1.17.0 (#31607); session export still missing (#19513).
- **Provider integration friction** – Custom OpenAI‑compatible providers fail on streaming tool calls (#26412); Anthropic’s new fallback API broke stream validation (#31579).
- **Data integrity issues** – Model switching corrupts `session_message.seq` (#31606); GBK encoding not handled (#31604); CLI input gets polluted by stderr (#31588).

---

*Data sourced from [anomalyco/opencode](https://github.com/anomalyco/opencode) – generated for 2026-06-10.*

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# Pi Community Digest — 2026-06-10

## Today's Highlights

Pi v0.79.1 shipped today with **Claude Fable 5** support across Anthropic and Amazon Bedrock providers, alongside a new prompt template argument defaults system. The **Project Trust** feature landed to mixed reactions—Issue #5514 attracted 24 comments and 12 upvotes, prompting an immediate PR (#5549) to improve the UX. Additionally, the ecosystem saw the first PR for an **Amazon Bedrock Mantle** provider, adding GPT-5.5/5.4 support via OpenAI-compatible endpoints.

---

## Releases

**[v0.79.1](https://github.com/badlogic/pi-mono/releases/tag/v0.79.1)** — Released today

- **Claude Fable 5** — Now available on Anthropic and Amazon Bedrock providers, with adaptive thinking and `xhigh` effort support.
- **Prompt template defaults** — Templates can use default positional arguments such as `${1:-7}` for optional values.

---

## Hot Issues

1. **[#5514 – Project Trust Feature Feedback](https://github.com/badlogic/pi-mono/issues/5514)** *(CLOSED, 24 comments, 12 👍)*  
   The new trust-gating system landed to immediate backlash. Users find the per-folder trust prompts annoying across multiple machines, especially when they already know which folders they open. The high engagement triggered PR #5549 within hours.

2. **[#4180 – Links not clickable anymore](https://github.com/badlogic/pi-mono/issues/4180)** *(CLOSED, 13 comments)*  
   After a recent update switching to alternate terminal mode, all hyperlinks in agent output stopped working. This is a significant UX regression for users who rely on clickable references.

3. **[#4984 – Interactive mode crash on transient terminal EPIPE](https://github.com/badlogic/pi-mono/issues/4984)** *(CLOSED, 13 comments)*  
   Users experience crashes during `edit` tool calls due to unhandled `write EPIPE` errors. The crash surfaces as an uncaught exception, forcing session restarts. Community members report this has become more frequent in the last few days.

4. **[#4877 – Session folder collision](https://github.com/badlogic/pi-mono/issues/4877)** *(OPEN, 11 comments)*  
   Sessions at distinct paths like `/a/b/c/d` and `/a-b/c-d` can collide to the same session folder due to naive path-to-folder mapping. Not a data-loss risk, but a latent surprise for users with diverse project structures.

5. **[#3715 – `local-llm` streams terminate at 5 minutes](https://github.com/badlogic/pi-mono/issues/3715)** *(CLOSED, 10 comments, 4 👍)*  
   Long `Write` tool calls against local vLLM backends die after exactly 5 minutes from undici's default `bodyTimeout`. The `retry.provider.timeoutMs` setting doesn't override it. This affects users running local Qwen3 with thinking models.

6. **[#5363 – Add amazon-bedrock-mantle provider](https://github.com/badlogic/pi-mono/issues/5363)** *(OPEN, 7 comments, 3 👍)*  
   A feature request for a new Bedrock provider using the OpenAI-compatible Mantle API, which Converse doesn't support. This is the community's gateway to GPT-5.5/5.4 on AWS.

7. **[#5350 – Custom tool operations receive host-OS-resolved paths](https://github.com/badlogic/pi-mono/issues/5350)** *(OPEN, 6 comments)*  
   Windows users injecting custom SSH-based file operations find that paths are resolved to host-OS format (e.g., `C:\...`) instead of the remote Linux path. Breaks cross-platform tooling.

8. **[#5511 – Error: context shift is disabled](https://github.com/badlogic/pi-mono/issues/5511)** *(CLOSED, 4 comments)*  
   Users hitting 51.1% context usage see "Thinking..." stall and `/compact` fails with a 502 from the summarization endpoint. Likely a backend availability issue, but blocking for deep-context sessions.

9. **[#5035 – Background subagents compete Telegram getUpdates](https://github.com/badlogic/pi-mono/issues/5035)** *(CLOSED, 4 comments)*  
   Spawned background agents inherit Telegram bot credentials and start competing long-polling sessions, causing HTTP 409 conflicts. Integration bug for Telegram-connected workflows.

10. **[#5541 – MiniMax M3 model switching causes no thinking](https://github.com/badlogic/pi-mono/issues/5541)** *(CLOSED, 3 comments)*  
    Switching to minimax-m3 mid-session from Claude causes thinking to never engage. Fresh context works fine. Suggests session state contamination across model switches.

---

## Key PR Progress

1. **[#5563/#5564 – Claude Fable 5 and Mythos 5 model metadata](https://github.com/badlogic/pi-mono/pull/5563)** *(CLOSED)*  
   Dual PRs adding model metadata for Claude Fable 5 and Mythos 5 to the Anthropic provider, marking them as always-adaptive thinking. Omitted unsupported `thinking.type: "disabled"` and temperature payloads. Core infrastructure for the v0.79.1 feature.

2. **[#5561 – Claude Fable 5 on Amazon Bedrock](https://github.com/badlogic/pi-mono/pull/5561)** *(OPEN)*  
   Adapts Bedrock's reasoning requests to use `thinking.type=adaptive` with `output_config.effort` instead of `budget_tokens`. Unlocks `xhigh` effort support natively.

3. **[#5509 – Amazon Bedrock Mantle OpenAI Responses provider](https://github.com/badlogic/pi-mono/pull/5509)** *(OPEN)*  
   New provider for Bedrock Mantle's OpenAI-compatible API, adding GPT-5.5 and GPT-5.4 support. Modeled after Azure's OpenAI provider. A significant expansion of cloud model options.

4. **[#5549 – Improved project approval settings](https://github.com/badlogic/pi-mono/pull/5549)** *(CLOSED)*  
   Direct response to the #5514 backlash: adds global enable/disable, parent-folder inheritance, and a "trust parent" button in the approval dialog (VSCode-style). Aligns config and list commands with the new trust behavior.

5. **[#5547 – Experimental feature guard](https://github.com/badlogic/pi-mono/pull/5547)** *(CLOSED)*  
   Implements RFC 0043: running Pi with `PI_EXPERIMENTAL=1` gates experimental features. Allows the team to ship riskier changes behind a flag.

6. **[#5555 – Attach reasoning_details streamed before tool_calls](https://github.com/badlogic/pi-mono/pull/5555)** *(CLOSED)*  
   Fixes a silent data-loss bug where `reasoning_details` signatures arriving before `tool_calls` chunks (e.g., OpenRouter + Gemini) were dropped. Critical for encrypted reasoning integrity.

7. **[#5554 – Add opus-4-8 to supportsAdaptiveThinking](https://github.com/badlogic/pi-mono/pull/5554)** *(CLOSED)*  
   Quick fix: Claude Opus 4.8 wasn't in the adaptive thinking list, causing 400 errors from the Anthropic API. One-line change with large impact.

8. **[#5553 – Prompt template argument defaults](https://github.com/badlogic/pi-mono/pull/5553)** *(CLOSED)*  
   Implements `${N:-default}` syntax for prompt templates, enabling optional positional arguments. Includes documentation and regression tests.

9. **[#5270 – Ephemeral session model and thinking level selection](https://github.com/badlogic/pi-mono/pull/5270)** *(CLOSED)*  
   Makes in-session configuration changes (Ctrl+P, Ctrl+T, `/model`) session-only by default, requiring explicit `{ persist: true }` to overwrite global defaults. Prevents accidental global overrides.

10. **[#5385 – Detect first-run terminal theme](https://github.com/badlogic/pi-mono/pull/5385)** *(OPEN)*  
    Queries the terminal via OSC to auto-detect light/dark theme on first run and persist it to settings. Quality-of-life improvement for new users.

---

## Feature Request Trends

- **Trust/Approval Workflow** – The #5514 backlash has created immediate demand for parent-folder inheritance, global disable, and VSCode-style "always trust parent" dialogs. PR #5549 addresses this, but users want even finer control (per-path/network-based trust).
- **Remote/WSL Path Support** – Issues like #5350 highlight a growing need for first-class Windows-to-Linux path translation in custom tool operations. Cross-platform developers are a vocal minority.
- **Plugin & Shell Integration** – Requests for `/about` command (#5548), shell plugin support (#5578), and persona overrides (#5577) point to users treating Pi as a general agentic harness, not just a coding tool.
- **API and Provider Expansion** – Bedrock Mantle (#5363, #5509) and OpenRouter custom model support (#5544) show demand for more cloud/third-party provider options, especially for GPT-5.x models.
- **Session & Configuration Management** – Ephemeral model changes (#5270) and session-only thinking levels reflect a desire for per-session flexibility without clobbering global config.

---

## Developer Pain Points

- **Streaming/Terminal Issues** – Multiple reports of viewport locking during rendering (#5192), chat view jumping during streaming (#5576), and widget flickering (#5579). TTY rendering remains a persistent pain point, especially on Windows.
- **Model-Specific Compatibility** – Switching models mid-session breaks thinking for MiniMax M3 (#5541). Thinking disabled/enabled toggle causes 400 errors for adaptive models (#5569). Each model has slightly different API semantics, and keeping Pi compatible is a delicate dance.
- **Session State Confusion** – Session folder collisions (#4877), stale tool execution callbacks (#5573), and context shift failures (#5511) all point to state management that doesn't yet handle all edge cases gracefully.
- **Tool Execution Reliability** – EPIPE crashes (#4984), 5-minute timeout ceilings (#3715), and agent output swallowing (#5580) frustrate users who depend on long-running tool calls. The `retry.timeoutMs` setting being ineffective compounds the issue.
- **Local Model Latency** – The 3-5 minute "Working" status delay on basic messages (#5464) makes local model usage nearly unusable for interactive workflows. Users want faster feedback loops.

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

Here is the Qwen Code community digest for 2026-06-10, based on the provided GitHub data.

---

### Qwen Code Community Digest – 2026-06-10

#### 1. Today's Highlights

Today's activity is dominated by a significant push toward stability and production readiness. Key releases and PRs focus on hardening the core: preventing out-of-memory (OOM) errors, stabilizing the prompt cache against dynamic MCP/skill changes, and fixing a blocking startup issue with dual-output FIFOs. Alongside this, a major experimental feature for parallel sub-agent coordination (Agent Team) has been proposed, and the community continues to advocate for better multi-provider model management and enhanced terminal rendering.

#### 2. Releases

- **[v0.18.0-preview.1](https://github.com/QwenLM/qwen-code/releases/tag/v0.18.0-preview.1)**: A new preview release was cut, but the initial release workflow for `v0.18.0-preview.2` ([Issue #4920](https://github.com/QwenLM/qwen-code/issues/4920)) failed this morning. The `v0.18.0-preview.1` changelog is minimal, including a chore for the previous release and a CLI fix to skip "thought" parts during copy output.

#### 3. Hot Issues

- **[#4877: OpenWork can't distinguish same model from different providers](https://github.com/QwenLM/qwen-code/issues/4877)** – [P2, Bug]  
  Users with multiple OpenAI-compatible providers (e.g., `glm-5` and `qwen-turbo`) find that OpenWork treats identically-named models as one, preventing proper provider selection. This is a core UX issue for users managing heterogeneous backends.

- **[#4891: Terminal resize during streaming leaves fragmented content](https://github.com/QwenLM/qwen-code/issues/4891)** – [P2, Bug]  
  A persistent rendering bug where resizing the terminal during generation creates visually fragmented scrollback. The `welcome-pr` tag suggests the team is open to contributions for this fix.

- **[#4882: [FEATURE] terminalSequence field on hook](https://github.com/QwenLM/qwen-code/issues/4882)** – [P3, Feature Request]  
  The community is requesting parity with Claude Code's hook system, specifically the ability for hooks to emit terminal-side effects (notifications, title updates) without needing a controlling terminal.

- **[#4921: "Virtualized History" causes viewport height bug](https://github.com/QwenLM/qwen-code/issues/4921)** – [P3, Bug]  
  A new bug report from today shows that enabling the "Virtualized History" setting breaks the terminal viewport, causing an unexpected scrollbar and cursor offset.

- **[#4910: Support installing extensions from archive files and URLs](https://github.com/QwenLM/qwen-code/issues/4910)** – [P2, Feature Request]  
  Following the addition of multiple extension sources, the community now wants the ability to install directly from `.zip`/`.tar.gz` archives and arbitrary download URLs.

- **[#4864: CI: enable required status checks on main branch protection](https://github.com/QwenLM/qwen-code/issues/4864)** – [P2, Enhancement]  
  A serious process issue: a PR was merged with CI failures, breaking the `main` branch. The community is pushing for mandatory status checks to prevent merge train derailments.

- **[#4423: MaxListenersExceededWarning (AbortSignal leak)](https://github.com/QwenLM/qwen-code/issues/4423)** – [Bug]  
  A long-standing memory leak issue on macOS (iTerm2) has been closed after weeks. The high number of abort listeners (1596) suggests a deep lifecycle management problem that may now be resolved.

- **[#4920: Release Failed for v0.18.0-preview.2](https://github.com/QwenLM/qwen-code/issues/4920)** – [CI Failure]  
  An automated release pipeline failure was logged this morning, indicating potential issues with the build or publishing infrastructure.

- **[#4864 (continued): CI enforcement](https://github.com/QwenLM/qwen-code/issues/4864)** – The community's push for CI discipline is a recurring theme, highlighting a developer pain point regarding code quality gates.

- **[#4896: Stabilize prompt-cache prefix against MCP/skills churn](https://github.com/QwenLM/qwen-code/pull/4896)** – *Linked PR; see PR section* – This issue is tied to a key performance fix, making it a top discussion point.

#### 4. Key PR Progress

- **[#4850: Interactive multi-tab /extensions manager](https://github.com/QwenLM/qwen-code/pull/4850)**  
  A comprehensive rework of the `/extensions` command into a full manager UI with Installed, Discover, and Sources tabs, covering the entire extension lifecycle.

- **[#4844: Experimental Agent Team feature](https://github.com/QwenLM/qwen-code/pull/4844)**  
  A major new feature enabling parallel sub-agents (teammates) that can communicate, share task lists, and consolidate results under a leader agent.

- **[#4914: Harden OOM prevention – idempotent compaction tests, explicit GC](https://github.com/QwenLM/qwen-code/pull/4914)**  
  Critical stability work. Fixes a counting bug in tool group compaction and adds regression tests for OOM prevention, including explicit garbage collection.

- **[#4922: Support image upload and echo in WebShell](https://github.com/QwenLM/qwen-code/pull/4922)**  
  Extends the WebShell UI to support multimodal image uploads, allowing users to paste or upload images that are echoed back as thumbnails.

- **[#4853: Add enter_plan_mode tool and Plan Approval Gate](https://github.com/QwenLM/qwen-code/pull/4853)**  
  Adds a proactive planning tool where the model can self-lower into plan mode for complex tasks, with an approval gate for manual oversight.

- **[#4896: Stabilize prompt-cache prefix against MCP/skills churn](https://github.com/QwenLM/qwen-code/pull/4896)**  
  A performance optimization that decouples skill visibility from validation, ensuring mid-session MCP changes don't invalidate the entire prompt cache.

- **[#4880: Layered tool-output truncation](https://github.com/QwenLM/qwen-code/pull/4880)**  
  Implements a three-layer truncation model for tool output (single-result, per-message, per-tool) to prevent context window overflow from oversized tool returns.

- **[#4924: POST /workspace/reload-env for hot-reloading env vars](https://github.com/QwenLM/qwen-code/pull/4924)**  
  Adds a daemon endpoint to reload environment variables and API keys without restarting the service, refreshing auth on all active sessions.

- **[#4894: Prevent FIFO blocking on startup](https://github.com/QwenLM/qwen-code/pull/4894)**  
  Fixes a startup hang when `--json-file` points to a named pipe (FIFO) that has no reader attached, by switching to non-blocking open mode.

- **[#4919: Debounce resize repaint and clear stale scrollback](https://github.com/QwenLM/qwen-code/pull/4919)**  
  A direct fix to Issue #4891, this PR debounces terminal resize events and clears fragmented scrollback after a settle timeout.

- **[#4903: Auto-detect SYSTEM account and default PATH scope to machine](https://github.com/QwenLM/qwen-code/pull/4903)** (Runner up)  
  Improves the Windows installer to correctly handle SYSTEM account installations (e.g., via CI), ensuring `qwen` is added to the machine-wide PATH.

#### 5. Feature Request Trends

The community's feature requests are converging on a few key themes:

- **Advanced Extension Installation**: Moving beyond Git/marketplace sources to support direct archive files (`.zip`, `.tar.gz`) and arbitrary URLs ([#4910](https://github.com/QwenLM/qwen-code/issues/4910)).
- **Rich Hook & Event System**: Seeking parity with Claude Code's hook protocol, specifically the `terminalSequence` field for terminal-side effects ([#4882](https://github.com/QwenLM/qwen-code/issues/4882)).
- **Better Terminal & Rendering UX**: Persistent requests for improvements to terminal resize handling, virtualized history, and general scrollback fidelity ([#4891](https://github.com/QwenLM/qwen-code/issues/4891), [#4921](https://github.com/QwenLM/qwen-code/issues/4921)).
- **OOM Prevention & Memory Management**: The community is actively engaged in hardening the application against memory leaks and out-of-memory crashes ([#4914](https://github.com/QwenLM/qwen-code/pull/4914), [#4423](https://github.com/QwenLM/qwen-code/issues/4423)).

#### 6. Developer Pain Points

- **Multi-Provider Model Management**: The inability to distinguish identical model names from different providers (e.g., `glm-5` from two OpenAI-compatible endpoints) is a persistent UX pain point ([#4877](https://github.com/QwenLM/qwen-code/issues/4877)).
- **CI/CD Fragility**: The accidental merge of a failing PR into `main` highlights a critical process failure. Developers are frustrated by the lack of mandatory status checks on branch protection ([#4864](https://github.com/QwenLM/qwen-code/issues/4864)).
- **Terminal Rendering Edge Cases**: Resizing the terminal during streaming and the new "Virtualized History" viewport bug are causing significant visual clutter and user confusion ([#4891](https://github.com/QwenLM/qwen-code/issues/4891), [#4921](https://github.com/QwenLM/qwen-code/issues/4921)).
- **Dual-Output / FIFO Blocking**: Users relying on programmatic output (e.g., daemon mode) face startup stalls when the output pipe isn't immediately read ([#4894](https://github.com/QwenLM/qwen-code/pull/4894)).
- **Frequent UX/Config Adjustments**: The community shows a pattern of needing to frequently adjust their configuration (`settings.json`) or UI settings (e.g., Virtualized History) to address regressions or new bugs.

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

# DeepSeek TUI (CodeWhale) Community Digest — 2026-06-10

**Note:** The project has fully rebranded to **CodeWhale**. The legacy npm package `deepseek-tui` is deprecated. All new releases are published under the `codewhale` name.

---

## Today's Highlights

The community shipped **v0.8.56 "Community Harvest"** with 14 merged PRs from 9 contributors, delivering major strides in localization, provider routing, prefix-cache stability, and shell safety. Alongside the release, a critical Windows npm install bug (#765) remains the most active open issue, while the team is pushing forward with the **Hotbar** (MMO-style quick action bar) feature stack and a multi-agent orchestration surface. An important process fix (#2985) was raised to ensure release commits land on `main` so `Closes #` lines actually resolve.

---

## Releases

### v0.8.56 — Community Harvest
- **14 merged PRs** from 9 contributors
- **Added:** Status picker localization across all 7 locales via 19 new `MessageId` variants (#2896)
- **Fixes:** Provider routing for `SiliconflowCN` config field; oversized paste files now write to `.codewhale/pastes/` (legacy `.deepseek/pastes/` removed); shell tools blocking diagnostic when `allow_shell` is disabled; mouse scrolling no longer clears Composer input
- **Refactors:** `allow_shell` moved from static system prompt prefix to per-turn `<runtime_prompt>` tag for better prefix-cache hit rates
- → [Release details](https://github.com/Hmbown/CodeWhale/releases/tag/v0.8.56)

### v0.8.55 — Together AI, OpenAI Codex, Model Catalog
- Added **Together AI** and **OpenAI Codex** as provider endpoints
- Introduced a structured **Model Catalog**
- Finalized the rebrand from `deepseek-tui` → `codewhale` (npm package `deepseek-tui` is deprecated; see `docs/REBRAND.md`)
- → [Release details](https://github.com/Hmbown/CodeWhale/releases/tag/v0.8.55)

---

## Hot Issues (Top 10)

1. **#765 — [BUG] Conversation cannot be triggered on Windows via npm global install (Stuck at "Working")**
   - **Votes:** 0 | **Comments:** 8
   - **Why it matters:** The most-commented open issue. Windows + WSL2 users cannot send chat messages — the UI hangs indefinitely on "Working". Multiple users report the same behavior after fresh installs.
   - → [Issue #765](https://github.com/Hmbown/CodeWhale/issues/765)

2. **#2007 — [ENHANCEMENT] Migration runs for coordinated multi-agent work**
   - **Votes:** 0 | **Comments:** 6
   - **Why it matters:** Proposes replacing the School-mode exploration with a visible, coordinated multi-agent orchestration surface — bounded parallel workers, role assignment, disagreement reconciliation. This is a v0.8.x major feature direction.
   - → [Issue #2007](https://github.com/Hmbown/CodeWhale/issues/2007)

3. **#1607 — [ENHANCEMENT] Token cost estimation: add more currency units (CNY, etc.)**
   - **Votes:** 0 | **Comments:** 5
   - **Why it matters:** Chinese-speaking users want RMB (CNY) and other currencies in the token cost display. A clear, actionable localization/financial UX request.
   - → [Issue #1607](https://github.com/Hmbown/CodeWhale/issues/1607)

4. **#1871 — [ENHANCEMENT] QoL: taskbar progress, animated title spinner, and configurable completion sound**
   - **Votes:** 1 | **Comments:** 4
   - **Why it matters:** Users want visual+audible feedback when alt-tabbed away — taskbar progress, terminal title animation, and a completion sound. High-value ergonomics for power users.
   - → [Issue #1871](https://github.com/Hmbown/CodeWhale/issues/1871)

5. **#2985 — [BUG] Release flow: commits should land on `main` so `Closes #` lines actually process**
   - **Votes:** 0 | **Comments:** 0 (new)
   - **Why it matters:** The v0.8.55 release branch was tagged but never merged back to `main`, breaking issue auto-closure. This is a critical process-fix for maintainers.
   - → [Issue #2985](https://github.com/Hmbown/CodeWhale/issues/2985)

6. **#1172 — [ENHANCEMENT] Support plugin/agent workflows from Cursor, Claude Code, Codex**
   - **Votes:** 0 | **Comments:** 2
   - **Why it matters:** Users want a portable agent workflow system (plugin market + hooks + skills + agents) that works across multiple coding agent harnesses. High demand for cross-platform portability.
   - → [Issue #1172](https://github.com/Hmbown/CodeWhale/issues/1172)

7. **#1778 — [BUG] Mouse scrolling clears user input in Composer**
   - **Votes:** 0 | **Comments:** 1
   - **Why it matters:** A frustrating UX bug — accidentally scrolling the mouse wheel while typing wipes the entire Composer text. Reproducible across sessions.
   - → [Issue #1778](https://github.com/Hmbown/CodeWhale/issues/1778)

8. **#1425 — [BUG] Long text processing session hangs (agent_wait timeout)**
   - **Votes:** 0 | **Comments:** 1
   - **Why it matters:** Analyzing large documents (3M+ character novels) spawns subagents that time out, causing session hangs. Users need better subagent lifecycle management for heavy workloads.
   - → [Issue #1425](https://github.com/Hmbown/CodeWhale/issues/1425)

9. **#1888 — [ENHANCEMENT] Slash commands: control-plane semantics for agents, jobs, hooks, recovery**
   - **Votes:** 0 | **Comments:** 1
   - **Why it matters:** Proposes a unified control-plane contract for long-running slash commands — inspect, pause/redirect/cancel, resume/recover. This is foundational for reliable agent workflows.
   - → [Issue #1888](https://github.com/Hmbown/CodeWhale/issues/1888)

10. **#1409 — [ENHANCEMENT] Support OAuth 2.1 for MCP authentication**
    - **Votes:** 1 | **Comments:** 1
    - **Why it matters:** MCP servers increasingly require OAuth 2.1 (not just API keys). Users cannot connect to `tinyfish` and other services. A high-impact integration gap.
    - → [Issue #1409](https://github.com/Hmbown/CodeWhale/issues/1409)

---

## Key PR Progress (Top 10)

1. **#2892 — feat(i18n): localize sandbox elevation dialog across 7 locales**
   - Migrates the sandbox elevation dialog from hardcoded English to `MessageId`-based translations (En, Ja, ZhHans, ZhHant, PtBr, Es419, Vi). 18 new `MessageId` variants.
   - → [PR #2892](https://github.com/Hmbown/CodeWhale/pull/2892)

2. **#2945 — feat(tui): render hotbar in sidebar**
   - First visible UI for the Hotbar feature — renders configurable slots at the bottom of the right sidebar with compact two-row layout and highlight.
   - → [PR #2945](https://github.com/Hmbown/CodeWhale/pull/2945)

3. **#2773 — feat(provider): complete provider fallback chain (#2574)**
   - Adds automatic fallback to the next configured provider on retryable errors (429, 5xx, timeout, network). Configurable via `fallback_providers` in TOML.
   - → [PR #2773](https://github.com/Hmbown/CodeWhale/pull/2773)

4. **#2865 — Modernize toward latest Claude Code (prompts, hooks, skills, agents, UI)**
   - Large gap-analysis PR that aligns CodeWhale with latest Claude Code behavior across lifecycle, skills/agents, and UI. Closes the long tail of feature parity.
   - → [PR #2865](https://github.com/Hmbown/CodeWhale/pull/2865)

5. **#2920 — fix(tui): write oversized paste files to .codewhale/pastes/**
   - Migrates paste storage from legacy `.deepseek/pastes/` to `.codewhale/pastes/`, completing the rebranding for file-system paths.
   - → [PR #2920](https://github.com/Hmbown/CodeWhale/pull/2920)

6. **#2579 — refs(#2264): Phase 4 — replace Session.messages with AppendLog**
   - Replaces `Vec<Message>` with an append-only `AppendLog` wrapper, improving session data model safety and enabling future streaming/checkpoint features.
   - → [PR #2579](https://github.com/Hmbown/CodeWhale/pull/2579)

7. **#2634 — feat: porting to HarmonyOS**
   - Makes CodeWhale compile on HarmonyOS/OpenHarmony by conditionally excluding Linux-only dependencies. A significant cross-platform port.
   - → [PR #2634](https://github.com/Hmbown/CodeWhale/pull/2634)

8. **#2971 — feat(execpolicy): expose matched approval rule metadata**
   - Adds explainability to execution-policy approvals — surfaces which `permissions.toml` rule matched on each approval request event.
   - → [PR #2971](https://github.com/Hmbown/CodeWhale/pull/2971)

9. **#2947 — fix(tui): guide long shell work to background**
   - Adds >5-second threshold guidance to the model prompt for shell work, aligning execution schema with background task expectations.
   - → [PR #2947](https://github.com/Hmbown/CodeWhale/pull/2947)

10. **#2905 — fix(tui): name `allow_shell` blocker for shell tools**
    - Improves the diagnostic message when `allow_shell` is disabled — now explicitly names the blocker and provides a fix command.
    - → [PR #2905](https://github.com/Hmbown/CodeWhale/pull/2905)

---

## Feature Request Trends

| Trend | Description | Key Issues |
|-------|-------------|------------|
| **Multi-agent orchestration** | Replace single-agent with coordinated multi-agent runs — scouts, subagents, reconciliation, durable work surfaces | #2007, #2024, #1976 |
| **Hotbar / quick-action bar** | MMO-style single-keypress slot bar for frequent toggles (voice, mode, palette) | #2061 (umbrella), #2063, #2064, #2065 |
| **Globalization & localization** | More currencies (CNY, JPY, EUR) for token costs, full localization for sandbox/settings dialogs | #1607, #2892 |
| **Plugin / agent workflow portability** | Cross-platform portable agent workflows that work across Cursor, Claude Code, Codex, and CodeWhale | #1172 |
| **Control-plane for commands** | Inspect, pause/resume, cancel, recover for long-running slash commands | #1888, #1890 |
| **MCP authentication** | Support OAuth 2.1 in addition to API key auth for MCP tools | #1409 |
| **Session continuity** | Non-interactive mode with `--resume` / `--session-id` for multi-turn workflows | #1530 |
| **QoL feedback** | Taskbar progress, terminal title spinner, configurable sound on completion | #1871 |

---

## Developer Pain Points

1. **Windows / WSL2 compatibility:** The #1 blocker — npm global install on Windows gets stuck at "Working" with no response. Multiple users report the same issue across different WSL2 builds.
2. **Session hangs on large workloads:** Processing large documents (3M+ characters) causes subagent timeouts and session hangs. Users need better timeout handling and session recovery.
3. **Mouse interaction bugs:** Scrolling in the Composer clears text — a persistent UX regression that breaks input on desktop.
4. **Prefix-cache instability:** Static system prompt prefix prevents effective caching. The `allow_shell` move is a first step, but more work needed.
5. **Cross-platform gaps:** `loongarch64` support blocked by `portable-pty` 0.8 dependency on `nix` 0.25.x. HarmonyOS port (#2634) shows growing demand for non-x86 platforms.
6. **Process hygiene:** Release branches that never merge to `main` break issue auto-closure and contributor attribution (#2985). The team is addressing this with a harvest-credit template and stewardship docs.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*