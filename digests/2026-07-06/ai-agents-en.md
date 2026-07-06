# OpenClaw Ecosystem Digest 2026-07-06

> Issues: 149 | PRs: 500 | Projects covered: 13 | Generated: 2026-07-06 13:05 UTC

- [OpenClaw](https://github.com/openclaw/openclaw)
- [NanoBot](https://github.com/HKUDS/nanobot)
- [Hermes Agent](https://github.com/nousresearch/hermes-agent)
- [PicoClaw](https://github.com/sipeed/picoclaw)
- [NanoClaw](https://github.com/qwibitai/nanoclaw)
- [NullClaw](https://github.com/nullclaw/nullclaw)
- [IronClaw](https://github.com/nearai/ironclaw)
- [LobsterAI](https://github.com/netease-youdao/LobsterAI)
- [TinyClaw](https://github.com/TinyAGI/tinyagi)
- [Moltis](https://github.com/moltis-org/moltis)
- [CoPaw](https://github.com/agentscope-ai/CoPaw)
- [ZeptoClaw](https://github.com/qhkm/zeptoclaw)
- [ZeroClaw](https://github.com/zeroclaw-labs/zeroclaw)

---

## OpenClaw Deep Dive

# OpenClaw Project Digest — 2026-07-06

## 1. Today's Overview

OpenClaw is experiencing extremely high activity: **149 issues** and **500 PRs** updated in the last 24 hours, with **237 PRs merged or closed** — a clear sign of rapid, sustained development. Despite no new releases today, the volume of fixes (especially crash, regression, and security-related) indicates the project is stabilising and hardening for scale. Community engagement remains strong, with the top issue (Linux/Windows desktop apps) gathering 110 comments and 81 reactions. The maintainer team is actively reviewing and merging contributions, though a long tail of feature requests and open PRs reveals areas needing product decisions.

## 2. Releases

**None.** No new releases were recorded for 2026-07-06. The project appears focused on merging bug fixes and features into `main` without creating a formal release today.

## 3. Project Progress

In the last 24 hours, **237 PRs were merged or closed**. Notable merged PRs include:

- **[#100893]** `fix(auto-reply): treat transient preflight compaction failures as skip reasons` — resolves a P1 bug that permanently locked the Composer.
- **[#100882]** `refactor: remove unused internal exports` — improves code quality.
- **[#100711]** `fix(acp): close SQLite event ledger on shutdown to prevent file lock leaks` — addresses file descriptor leaks.
- **[#98857]** `fix(codex-supervisor): guard loadCodexSupervisorEndpoints JSON.parse against malformed env input` — hardens startup error handling.
- **[#100890]** `fix(diagnostics): reconcile queueDepth to zero on session idle transition` — fixes false queued signal in diagnostics.
- **[#85616]** `fix(auto-reply): fast path text control commands` — improves Feishu/Lark command handling.
- **[#68936]** `Autofix: add PR review autofix pipeline + Windows daemon` — introduces automated review response and Windows gateway supervision.
- **[#96525]** (issue closed) — cron model-fallback chain bug fixed.
- **[#100600]** (issue closed) — Azure AI Foundry misclassification fixed.

33 issues were closed, reflecting progress on many reported bugs. Feature work also continues: e.g., [#100789](feat: show background tasks in Control UI) and [#100814](feat: session grouping and unread state across clients) are open but advancing.

## 4. Community Hot Topics

The most active discussions and reactions in the last 24 hours (all issues remain open unless noted):

- **[#75]** `Linux/Windows Clawdbot Apps` — **110 comments, 81 👍**. The overwhelming community priority. Users demand desktop parity with macOS/iOS/Android, signalling a need for broader platform support.
- **[#7707]** `Feature Request: Memory Trust Tagging by Source` — **16 comments**. A security-conscious request to prevent memory poisoning by tagging entries by origin (user input, web scrape, third-party skill). Reflects growing enterprise use cases.
- **[#99241]** `Tool outputs sometimes render as image attachments and become unreadable to the agent` — **12 comments, P1**. A critical usability bug that breaks long-running tool workflows. Community frustration is high because the agent loses ability to read evidence.
- **[#13583]** `Pre-response enforcement hooks (hard gates)` — **12 comments**. Demand for mechanical enforcement of mandatory tool calls (quant/finance, security). Reveals need for policy-as-code in high-stakes environments.
- **[#63829]** `Per-agent memory-wiki vault configuration` — **11 comments, 9 👍**. Multi-agent isolation is a top ask. Users want each agent to maintain its own knowledge vault instead of a global shared one.
- **[#14785]** `Reduce tool schema token overhead (~3,500 tok/session)` — **8 comments**. Efficiency concern: every session pays a fixed token tax for all tool schemas, regardless of usage. Users want optimizations for cost and latency.
- **[#16670]** `Onboarding Wizard should include Memory/Embedding setup as a mandatory step` — **8 comments**. Critical UX gap: memory (a core feature) is not configured during setup, leading to confusion later.

These threads collectively show a community that is advanced, security-aware, and demanding of both new features and polish.

## 5. Bugs & Stability

High-severity bugs reported today, ranked by priority and whether fix PRs exist:

| Issue | P | Summary | Fix PR | Status |
|-------|---|---------|--------|--------|
| #99241 | P1 | Tool outputs collapse to image attachment, agent cannot read text | — | Open |
| #100778 | P1 | Preflight compaction failure permanently locks Composer into "terminated" state | #100893 | **Merged** (fixed) |
| #100600 | P1 | Azure Foundry response.failed misclassified as auth failure (cooldown cascades) | — | **Closed** (fix applied) |
| #96525 | P1 | Isolated cron model-fallback chain never engages on result-level failures | — | **Closed** (fix applied) |
| #87058 | P1 | Android node connects but advertises zero commands (node-connect-reconcile bug) | — | Open |
| #53486 | P2 | Feishu card JSON rendered as plain text (regression) | #100883 | Open (fix submitted, waiting review) |
| #49205 | P2 | Control UI messages reach shared context but not visible in Open WebUI chat | — | Open |
| #92478 | P2 | Native Talk button triggers OpenAI Realtime WebRTC call and fails with 500 when not configured | — | Open |
| #87980 | P2 | exec tool silently corrupts 2>&1 / 2>/dev/null shell redirect arguments | — | Open |
| #100883 | P2 | Feishu card JSON message params as cards (fix PR) | — | Open (ready for maintainer) |

Several other P2 bugs (e.g., [#84623] non-streaming duplicate replies, [#97877] provider-error retry gating) have fix PRs open: [#100828] and [#100669] respectively.

**Critical regression risk:** The Feishu card regression (#53486) indicates a testing gap in the messaging pipeline, and the Talk button crash (#92478) points to incomplete feature gating.

## 6. Feature Requests & Roadmap Signals

The following high-value feature requests are likely candidates for upcoming releases based on community momentum and maintainer engagement:

- **Cross-platform desktop apps** (#75) – Longest-running, highest-reaction request. Predict inclusion in a future stable release.
- **Memory Trust Tagging** (#7707) – Directly related to security incidents. P2 but with multiple maintainer labels.
- **Pre-response enforcement hooks** (#13583) – P2 with security review needed. For high-stakes users.
- **Per-agent memory-wiki vault** (#63829) – P1 enhancement. Strong demand from multi-agent users.
- **Capability-based permissions for skills/tools** (#12678, #12219) – Core security infrastructure. Likely planned for next major.
- **Tool schema token overhead reduction** (#14785) – Efficiency win. Could be implemented as incremental optimization.
- **Onboarding wizard including memory setup** (#16670) – Easy UX win. Low-hanging fruit for next patch.
- **Context file elevation (SOUL.md)** (#17810) – Creative feature for persona-heavy users. Less common but highly appreciated.
- **Mid-stream message injection (soft steer)** (#10960) – Real-time interaction improvement. Could be designed for chat platforms.
- **Multi-role support in Chat UI** (#22774) – Needed for multi-agent workflows. Roadmap signal.

These features point toward consolidation: improved security model, better multi-agent and memory isolation, and reduction of token waste.

## 7. User Feedback Summary

**Satisfactions:**
- The project is very active; many issues resolved quickly (237 PRs merged in 24h).
- Security-conscious features (memory tagging, enforcement hooks) are being actively discussed and prioritised.
- Cross-platform ecosystem is expanding (note Android node issues show Android is in use).

**Dissatisfactions / Pain Points:**
- **Platform gaps:** Linux/Windows desktop apps missing (#75). Android node unreliable (#87058).
- **Onboarding friction:** Memory and embedding setup not guided, leading to broken experiences (#16670).
- **Tool output reliability:** Long-running or ANSI-heavy workflows break, agents become blind (#99241).
- **Token cost:** Fixed schema tax ~3,500 tokens per session is wasteful (#14785).
- **Channel regressions:** Feishu card handling broke (#53486), Telegram `notify=false` leaked (#80569).
- **Security model:** No enforceable permission system for skills; memory poisoning risk (#7707, #12678).
- **Confusing error states:** Preflight compaction lock (#100778), cron fallback silent drops (#96525).

These pain points are well-documented and many have active fix PRs, indicating responsive development.

## 8. Backlog Watch

The following issues and PRs have been open for extended periods and need maintainer attention:

| Item | Type | Created | Priority | Last Activity | Notes |
|------|------|---------|----------|---------------|-------|
| [#75] Linux/Windows Clawdbot Apps | Issue | 2026-01-01 | P2 | Updated today | 110 comments, 81 👍, still open. No assigned owner or milestone. |
| [#14785] Reduce tool schema token overhead | Issue | 2026-02-12 | P2 | Updated today | 8 comments, no PR yet. Could be a quick win. |
| [#16670] Onboarding Wizard memory setup step | Issue | 2026-02-15 | P2 | Updated today | 8 comments. UX gap with low implementation effort. |
| [#13364] Expose before/after_tool_call in hooks system | Issue | 2026-02-10 | P2 | Updated today | 7 comments, security review pending. |
| [#13597] AWS deployment guide | Issue | 2026-02-10 | P2 | Updated today | 6 comments, 3 👍. Documentation needed for cloud adoption. |
| [#13751] Feishu plugin: remove contact permission for sender name | Issue | 2026-02-11 | P1 | Updated today | Security request, P1 but still open. |
| [#10118] TUI Shift+Enter for newline | Issue | 2026-02-06 | P3 | Updated today | Small UX improvement, 5 comments, 4 👍. |
| [#100840] fix(matrix): fallback to copy+unlink on EXDEV | PR | 2026-07-06 | P2 | Open | Status: 📣 needs proof. Low merge risk. |
| [#100846] fix(cron): do not set delivery mode to announce | PR | 2026-07-06 | P2 | Open | Status: 📣 needs proof. Requires maintainer review. |
| [#100814] feat(sessions): grouping, unread state, full controls | PR | 2026-07-06 | P2 | Open | Status: ⏳ waiting on author. Large change (XL), merges expected after revision. |

**Attention needed:** The top-voted issue (#75) remains without a maintainer decision after 6 months. Many P2 enhancement issues have been pending review since February. The Feishu card fix PR (#100883) and the Android notification fix (#100888) are ready for maintainer review.

## Links

All issues and PRs can be found at `openclaw/openclaw` on GitHub. Specific URLs are listed inline above (e.g., `openclaw/openclaw Issue #75`).

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report — AI Agent Open-Source Ecosystem
**2026-07-06**

## 1. Ecosystem Overview

The personal AI assistant and agent open-source landscape is experiencing a **security-driven maturation phase**, with multiple projects undergoing comprehensive audits and hardening sprints. Activity is concentrated around three axes: **Anthropic/Claude provider optimization**, **cross-platform desktop parity**, and **multi-agent memory isolation**. The ecosystem is bifurcated between high-velocity, well-funded reference implementations (OpenClaw, IronClaw) and smaller specialized forks (PicoClaw, NanoClaw) that prioritize specific provider or security niches. Several projects show signs of **community bottlenecking**, where issues exceed maintainer review capacity despite high contributor engagement. The overall trend is toward production-grade reliability with enterprise security requirements becoming baseline expectations rather than differentiators.

## 2. Activity Comparison

| Project | Issues (24h) | PRs (24h) | PRs Merged/Closed | Release Today | Health Score |
|---------|-------------|-----------|-------------------|---------------|--------------|
| **OpenClaw** | 149 | 500 | 237 | None | ★★★★★ |
| **NanoBot** | 40 | 500 | 4 | None | ★★★☆☆ |
| **Hermes Agent** | 7 | 50 | 16 | None | ★★★★☆ |
| **PicoClaw** | 3 | 6 | 1 | None | ★★★★☆ |
| **NanoClaw** | 2 | 7 | 0 | None | ★★☆☆☆ |
| **NullClaw** | 0 | 1 | 0 | None | ★☆☆☆☆ |
| **IronClaw** | 18 | 38 | 11 | None | ★★★★☆ |
| **LobsterAI** | 0 | 12 | 11 | None | ★★★★☆ |
| **Moltis** | 0 | 5 | 3 | None | ★★★☆☆ |
| **CoPaw** | 16 | 45 | 24 | **v1.1.12.post3** | ★★★★★ |
| **ZeroClaw** | 3 | 50 | 7 | None | ★★★★☆ |

*Health Score reflects merge velocity, community engagement, release cadence, and stability. Projects with critical security audits in progress (NanoBot) score lower due to open risk, not low activity.*

## 3. OpenClaw's Position

OpenClaw maintains a **clear leadership position** as the ecosystem's core reference implementation, outpacing all peers in raw activity (500 PRs/day, 149 issues updated). Key advantages:

- **Scale**: 81 community reactions on top-voted feature request (#75, desktop apps) dwarfs any single issue in other projects. 237 PRs merged in 24 hours exceeds the weekly output of most competitors.
- **Technical approach**: Modular micro-service architecture with ACP protocol and composable supervisor patterns (codex-supervisor, preflight compaction) contrasts with NanoBot's monolithic gateway or IronClaw's monorepo structure. This design enables the rapid fix cycle observed.
- **Community density**: The 110-comment thread on Linux/Windows desktop parity signals a broader, more demanding user base than any single competitor. However, the **6-month unresolved status of this top issue** indicates maintainer bandwidth constraints despite high merge velocity.
- **Gap**: Compared to IronClaw's deliberate Slack OAuth and HST Postgres stacks, OpenClaw's PRs are more tactical (bug fixes) than strategic. The project lacks visible long-term infrastructure investments visible in IronClaw's 4-PR Postgres latency stack.

## 4. Shared Technical Focus Areas

| Focus Area | Affected Projects | Specific Requirements |
|------------|------------------|----------------------|
| **Security hardening** | NanoBot, ZeroClaw, PicoClaw, OpenClaw | Plaintext credential storage fix (#4803 NanoBot), SSRF protection (#8713 ZeroClaw), libolm→vodozemac migration (#3088 PicoClaw), memory trust tagging (#7707 OpenClaw) |
| **Cross-platform desktop apps** | OpenClaw (#75), Hermes Agent (#56835) | Linux/Windows app gap; crash on sleep/resume in remote setups |
| **Tool output reliability** | OpenClaw (#99241), CoPaw (#5717) | Image-attachment rendering breaks agent reading; truncated tool_call.input causes infinite loops |
| **Multi-agent memory isolation** | OpenClaw (#63829), Hermes Agent (implicit in #56004) | Per-agent vault configuration; thinking context loss on replay |
| **Provider OAuth integration** | IronClaw (#5669), LobsterAI (#2276), Hermes (#34393) | Slack least-privilege OAuth, xAI/Grok login, Proton Pass secret management |
| **Anthropic provider maturity** | PicoClaw (#3228, #3229), OpenClaw (caching fixes) | Per-block cache_control, rolling conversation breakpoints, prompt caching support |
| **CI reliability** | IronClaw (#5684, #5687), ZeroClaw (#8753), CoPaw (#5736) | Canary report automation, workspace test gap, AI review bot integration |
| **Onboarding friction** | OpenClaw (#16670), ZeroClaw (#8718) | Memory setup not guided; config template silently breaks transcription |

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | IronClaw | CoPaw | LobsterAI | ZeroClaw |
|-----------|----------|---------|----------|-------|-----------|----------|
| **Primary focus** | Core platform stability | Security audit & remediation | Enterprise scaling | IM channel reliability | Consumer UX polish | SOP authoring & auth |
| **Target user** | Developers, power users | Security-conscious operators | Enterprise teams | Chinese IM users | General consumers | Workflow engineers |
| **Architecture** | Modular micro-services | Monolithic gateway | Monorepo with feature stacks | Bifurcated 1.x/2.x | Full-stack (renderer + agent) | Monorepo with schema evolution |
| **Release cadence** | Continuous, no formal release this week | Heavy triage, few merges | Stack-based, deliberate | Patch release today | Rapid merge, likely release soon | Blocked by large features |
| **Community language** | English-dominant | English-dominant | English | Chinese-dominant | Chinese-dominant | English |
| **Security posture** | Reactive (fixes after reports) | Proactive (audit with 35 findings) | Mitigated (OAuth migration) | Moderate | Moderate | Improving (env var blocking) |
| **Feature maturity** | Mature core, long feature backlog | Pre-audit, now stabilizing | Maturing via feature stacks | Stable 1.x, risky 2.0 beta | Feature-rich, low bugs | Schema evolution in progress |

**Emerging pattern**: Chinese-ecosystem projects (CoPaw, LobsterAI) focus on IM channel reliability (Feishu, DingTalk, WeChat) and consumer-grade UI, while English-ecosystem projects prioritize security, observability, and enterprise integration.

## 6. Community Momentum & Maturity

### Tier 1: Very High Velocity (Rapid Iteration)
- **OpenClaw**: 237 PRs merged/day, 149 issues updated. Stabilizing after period of rapid feature growth. **Risk**: maintainer attention is spread thin; top issue (#75) unresolved for 6 months.
- **NanoBot**: 500 PRs updated, but only 4 merged. Intense triage after audit. **Risk**: merge bottleneck could demotivate contributors.
- **IronClaw**: 38 PRs updated, 11 merged. Strategic feature stacks (Slack OAuth, HST Postgres) with clear ownership. **Healthiest** cadence among enterprise-focused projects.
- **CoPaw**: 24 merged/closed, patch release. Balanced 1.x maintenance with v2.0 pre-release work. **Highest stability** among active projects.

### Tier 2: High Velocity (Active Development)
- **Hermes Agent**: 16 merged/closed, focused bug fixes. Good responsiveness but lacks strategic roadmap visibility.
- **LobsterAI**: 11 merged/closed, single sprint-dominated day. **Risk**: low community engagement (0 issues with comments). Internal team-driven rather than community-driven.
- **ZeroClaw**: 7 merged/closed, but 43 open PRs. Large feature stacks (SOP authoring, multi-user auth) indicate **pre-release blockage** rather than maintenance churn.

### Tier 3: Moderate Steady State
- **PicoClaw**: Low volume but high quality (targeted Anthropic fixes). **Healthiest niche project** with responsive maintainers and low friction.
- **Moltis**: Low volume, steady maintenance. Acceptable for a mature, low-complexity project.

### Tier 4: Low Activity / Stalled
- **NanoClaw**: Documentation-only updates. No merged code. **Risk**: may be project-maintained but not actively developed.
- **NullClaw**: Single Dependabot PR. **Effectively dormant**.
- **TinyClaw, ZeptoClaw**: Zero activity. **Suspected unmaintained**.

## 7. Trend Signals

### 1. Security is Becoming Table Stakes, Not Differentiation
The NanoBot audit (35 findings including command injection, plaintext credentials) and ZeroClaw's SSRF vulnerability reflects a **rapidly maturing security baseline**. Projects without published security audits (Moltis, LobsterAI) face growing trust gaps as the ecosystem professionalizes. The pattern of "security first, features second" in NanoBot's 500-PR audit response will likely become standard practice.

### 2. Anthropic/Claude Provider Optimization is the Dominant Engineering Effort
Three projects (PicoClaw, OpenClaw, CoPaw) have active PRs or issues addressing prompt caching, tool-call reliability, or system message optimization for Claude. This reflects the **de facto standard status** of Anthropic's API for agentic workloads. Developers are optimizing for Claude's tool-use execution model rather than generic LLM abstractions.

### 3. OAuth is Replacing API Keys for Enterprise Integration
IronClaw's Slack OAuth migration, LobsterAI's xAI/Grok OAuth login, and Hermes Agent's Proton Pass integration signal a **systematic shift away from static credentials**. The ecosystem is converging on OAuth/PKCE patterns for both provider authentication and channel authorization (Slack, Telegram). Projects without OAuth support (NanoClaw, NullClaw) risk irrelevance for team deployments.

### 4. Multi-Agent and Memory Isolation Requirements are Converging
OpenClaw (#63829, per-agent memory vaults), Hermes Agent (#56004, thinking context loss), and PicoClaw (#3229, conversation cache breakpoints) all address variants of the same problem: **agents need independent memory spaces that persist correctly across sessions**. This is a prerequisite for production multi-agent deployments, which remain unsupported in most projects.

### 5. The TypeScript/Rust Frontier is Quietly Shifting
While Python remains the dominant language (OpenClaw, NanoBot, LobsterAI), Rust infrastructure is appearing in critical paths: PicoClaw's vodozemac migration, Moltis's WhatsApp library, and ZeroClaw's CI tooling. IronClaw's WASM linker optimization (#5677) hints at browser-based agent execution. **Observation**: no project has yet committed to a full Rust rewrite, but all are selectively adopting Rust for performance-critical components.

### 6. Feedback Loop Quality is Emerging as a Competitive Advantage
CoPaw's Feishu issue (#5757) with 11 comments and OpenClaw's tool-output bug (#99241) with 12 comments demonstrate that **projects with clear, actionable bug reports** (including reproduction steps, expected vs actual behavior) get faster fixes. Projects with low-quality issues (NanoClaw's "help with logo design") struggle to convert feedback into value. **Implication**: maintainers should invest in issue templates and reproduction guides to improve triage efficiency.

### 7. Chinese Ecosystem is an Underserved but Growing Segment
CoPaw, LobsterAI, and the NullClaw ecosystem serve Chinese users with specialized IM support (Feishu, DingTalk, WeChat) that no Western project matches. ZeroClaw's explicit addition of Bocha AI search (PR #8737) for China-based deployments signals recognition of this gap. **Opportunity**: Western projects that add WeChat or Feishu support could capture a significant underserved user base.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

Based on the provided GitHub data for NanoBot (github.com/HKUDS/nanobot) updated on **2026-07-06**, here is the project digest.

## Today's Overview

NanoBot saw an extraordinary surge in activity over the past 24 hours, with **500 pull requests updated** (496 open) and **40 issues updated** (38 open). This spike is largely driven by a comprehensive **security and correctness audit** (issue #4815) that delivered **35 findings**, each filed as a separate bug or refactor ticket. Only **4 PRs were merged/closed** and **2 issues closed**, indicating the project is currently in a heavy triage and fix‑creation phase rather than a release cycle. No new releases were published. The community response is strong but the maintainers face a growing backlog of critical security and stability concerns.

## Releases

**None** – no new versions were published on this date.

## Project Progress

Four pull requests were merged or closed today:

- [PR #4770](https://github.com/HKUDS/nanobot/pull/4770) – **fix(gateway): resolve config path for state refresh** – repairs a CI regression in gateway state self‑healing.
- [PR #4017](https://github.com/HKUDS/nanobot/pull/4017) – **fix(providers): parse text-format tool_calls in openai-compat responses** – addresses a provider compatibility issue (Xiaomi MiMo).
- [PR #4547](https://github.com/HKUDS/nanobot/pull/4547) – **fix(gateway): self-heal state file PID on server startup** – resolves the Windows `/restart` PID mismatch (related to issue #4511).
- [PR #4654](https://github.com/HKUDS/nanobot/pull/4654) – **fix(cli): print response text when streaming fails in interactive mode** – ensures complete answers are not lost when streaming fails silently.

Two issues were closed:

- [#4511](https://github.com/HKUDS/nanobot/issues/4511) – Windows `--background` gateway state file inconsistency (fixed by #4547).
- [#4765](https://github.com/HKUDS/nanobot/issues/4765) – Python SDK `async with` context manager protocol error (likely a documentation fix).

These moves show the team is actively resolving long‑standing bugs and improving stability, but the large number of newly opened issues suggests the pace of merging will need to accelerate.

## Community Hot Topics

The most active discussions by comment count:

1. **[#4511 (CLOSED) – Windows gateway `--background` bug](https://github.com/HKUDS/nanobot/issues/4511)** – 4 comments. Underlying need: reliable background process management across OS platforms (now resolved).
2. **[#4637 (OPEN) – Telegram long message truncation / rendering](https://github.com/HKUDS/nanobot/issues/4637)** – 2 comments. Underlying need: proper splitting and rendering of long Markdown messages on Telegram.
3. **[#4765 (CLOSED) – Python SDK async context manager error](https://github.com/HKUDS/nanobot/issues/4765)** – 2 comments. Underlying need: correct, working Python SDK examples in documentation.

Many newly filed issues (e.g., #4815, the audit summary) have zero comments but are highly significant due to their breadth. The PRs with most activity are the **fix‑series** authored by `axelray-dev` (e.g., #4816, #4814, #4813, #4812, #4811) and the **WebUI improvements** by `chengyongru` (#4766, #4771, #4769). These indicate the community is actively responding to the audit findings.

## Bugs & Stability

Today’s filings reveal **critical security and stability vulnerabilities**, ranked by severity:

| Severity | Issue | Summary |
|----------|-------|---------|
| **P0 / Critical** | [#4815](https://github.com/HKUDS/nanobot/issues/4815) | Audit summary – 35 findings including command injection, path escape, auth bypass, secrets in plaintext. |
| **P0 / Critical** | [#4803](https://github.com/HKUDS/nanobot/issues/4803) | API keys stored in plaintext in `~/.nanobot/config.json`. |
| **P1 / High** | [#4796](https://github.com/HKUDS/nanobot/issues/4796) | `restrict_to_workspace` defaults to `False` – full filesystem accessible to agent tools. |
| **P1 / High** | [#4797](https://github.com/HKUDS/nanobot/issues/4797) | No resource limits on shell subprocesses (CPU, memory, ulimit). |
| **P1 / High** | [#4790](https://github.com/HKUDS/nanobot/issues/4790) | Symlink TOCTOU in filesystem tools – escape possible during race window. |
| **P1 / High** | [#4785](https://github.com/HKUDS/nanobot/issues/4785) | `read_file` loads entire file before truncation – OOM risk on large files. |
| **P2 / Medium** | [#4802](https://github.com/HKUDS/nanobot/issues/4802) | Token budget returns 128 tokens when context window is disabled (`0`). |
| **P2 / Medium** | [#4804](https://github.com/HKUDS/nanobot/issues/4804) | Leaked `CancelledError` silently swallowed in main agent loop (MCP interaction). |
| **P2 / Medium** | [#4798](https://github.com/HKUDS/nanobot/issues/4798) | Concurrent file writes from different sessions not serialized – workspace file corruption. |
| **P2 / Medium** | [#4791](https://github.com/HKUDS/nanobot/issues/4791) | No channel‑level rate limiting – potential DoS on LLM tokens and session memory. |
| **P3 / Low** | [#4805](https://github.com/HKUDS/nanobot/issues/4805) | `suppress(Exception)` in tool runner swallows validation errors. |
| **P3 / Low** | [#4789](https://github.com/HKUDS/nanobot/issues/4789) | `WeakValueDictionary` for locks – GC can break mutual exclusion. |
| **P3 / Low** | [#4787](https://github.com/HKUDS/nanobot/issues/4787), [#4786](https://github.com/HKUDS/nanobot/issues/4786) | Unbounded `Session.messages` and `SessionManager._cache` – resource leaks. |

**Fix PRs exist** for several of these:

- [#4816](https://github.com/HKUDS/nanobot/pull/4816) – narrows `BaseException` catch to `Exception` in `_run_tool()`.
- [#4814](https://github.com/HKUDS/nanobot/pull/4814) – propagates `CancelledError` instead of swallowing it.
- [#4813](https://github.com/HKUDS/nanobot/pull/4813) – guards `.strip()` on multimodal content.
- [#4812](https://github.com/HKUDS/nanobot/pull/4812) – uses `.get()` for `role` key to prevent `KeyError`.
- [#4811](https://github.com/HKUDS/nanobot/pull/4811) – logs suppressed exceptions instead of silencing them.
- [#4701](https://github.com/HKUDS/nanobot/pull/4701) – catches `BaseException` in MCP wrappers.
- [#4671](https://github.com/HKUDS/nanobot/pull/4671) – pins validated DNS for SSRF checks.
- [#4768](https://github.com/HKUDS/nanobot/pull/4768) – adds exponential backoff to QQ WebSocket reconnection.
- [#4664](https://github.com/HKUDS/nanobot/pull/4664) – protects dream history during compaction.

## Feature Requests & Roadmap Signals

Several feature PRs remain open and may be candidates for the next release:

- [PR #4145](https://github.com/HKUDS/nanobot/pull/4145) – **Weather Skill** (new skill example). Open since June 1, awaiting review.
- [PR #4771](https://github.com/HKUDS/nanobot/pull/4771) – **Document attachments in WebUI** – support PDF/doc upload and validation. This PR has backend and frontend tests, suggesting near-term inclusion.
- [PR #4689](https://github.com/HKUDS/nanobot/pull/4689) – **OAuth status and expiry warnings** – adds CLI and WebUI visibility into OAuth tokens.
- [PR #4769](https://github.com/HKUDS/nanobot/pull/4769) – **Centralize native runtime access in WebUI** – refactoring to support better cross‑platform UI behavior.
- [PR #364](https://github.com/HKUDS/nanobot/pull/364) – **Comprehensive Cron Service Upgrade** (hot reload, heartbeat, echo mode). A long‑standing PR (since Feb 6) that would be a significant quality‑of‑life improvement.
- [PR #216](https://github.com/HKUDS/nanobot/pull/216) – **A2A Protocol support** for agent‑to‑agent communication. Another old PR (since Feb 6) that would enable inter‑agent collaboration.

Given the current focus on fixing critical bugs, the next version is likely to prioritise **security hardening** (plaintext credentials, filesystem restrictions, rate limiting) and **stability fixes** (timeouts, resource leaks, token budget). Feature inclusion of document uploads and OAuth status is plausible, while larger efforts like A2A or cron hot‑reload may be deferred.

## User Feedback Summary

Real user pain points reflected in recent issues:

- **Windows compatibility**: The `/restart` command on Windows left stale PID files (now fixed). Windows users continue to need first‑class support.
- **Telegram integration**: Long Markdown messages are split but only the final trunk renders correctly ([#4637](https://github.com/HKUDS/nanobot/issues/4637)) – a regression that impacts a widely used channel.
- **Python SDK documentation**: The example code for the async context manager was broken ([#4765](https://github.com/HKUDS/nanobot/issues/4765)), causing immediate errors for developers trying to embed NanoBot programmatically.
- **Security concerns**: The audit findings, especially plaintext API keys and unrestricted filesystem access, are likely to alarm users and operators. The community response (dozens of fix PRs in a single day) signals dissatisfaction with the previous security posture.

Overall satisfaction appears mixed: the platform is actively used and developers are contributing fixes, but the volume of high‑severity bugs indicates that recent rapid growth may have outpaced code quality controls.

## Backlog Watch

Several important items have not received maintainer attention for extended periods:

| Item | Age | Notes |
|------|-----|-------|
| [PR #364](https://github.com/HKUDS/nanobot/pull/364) – Cron Hot Reload | 5 months | Labeled `[conflict]` – unresolved merge conflicts block progress. |
| [PR #216](https://github.com/HKUDS/nanobot/pull/216) – A2A Protocol | 5 months | Also `[conflict]` – large feature but no recent updates. |
| [PR #4145](https://github.com/HKUDS/nanobot/pull/4145) – Weather Skill | 5 weeks | No maintainer review since June 1. |
| [Issue #4637](https://github.com/HKUDS/nanobot/issues/4637) – Telegram truncation bug | 5 days | Opened July 1, still unaddressed even though fix PRs exist for other bugs. |
| [Issue #4511](https://github.com/HKUDS/nanobot/issues/4511) – Windows background (now closed) | 11 days | Only resolved today after extended discussion. |

The maintainers should prioritise clearing the conflict‑labeled PRs (#364, #216) to avoid bitrot, and review the long‑standing Weather Skill PR to encourage new skill contributions. The Telegram rendering bug (#4637) also deserves a fix before it affects more users.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-07-06

## 1. Today's Overview

Hermes Agent maintains high development velocity: 50 pull requests were updated in the last 24 hours (34 open, 16 merged/closed), alongside 7 issues (5 open, 2 closed). No new releases were published. The activity is dominated by bug fixes and infrastructure improvements, with a notable cluster of work around gateway routing, Telegram reliability, and threat‑scanner user‑visible warnings. Community engagement remains healthy, with several multi‑comment discussions on memory context preservation and network resilience.

## 2. Releases

No new releases today.

## 3. Project Progress

**16 PRs were merged or closed today.** Highlights of changes that advanced the codebase:

- **Telegram Business Mode (observe‑with‑approval)** — PR [#30055](https://github.com/NousResearch/hermes-agent/pull/30055) (still open) adds Secretary‑Bot style draft‑and‑approve replies for Telegram Business accounts. This is a long‑standing feature that continues to receive updates.

- **Tool‑call execution from model `content`** — PR [#35129](https://github.com/NousResearch/hermes-agent/pull/35129) merged, enabling Ollama/Kimi/MiniMax/Gemma models to emit tool calls in the response `content` field and have them executed instead of leaked as chat text.

- **Image task deadline env var** — PR [#59629](https://github.com/NousResearch/hermes-agent/pull/59629) merged, introducing a 600‑second default deadline for image generation tasks and fixing a misnamed timeout variable.

- **Provider `default_headers` alias** — PR [#57086](https://github.com/NousResearch/hermes-agent/pull/57086) merged, accepting `default_headers` as a compatibility alias for `extra_headers` on providers (e.g., W&B/CoreWeave).

- **Memory provider pre‑compression insights** — PR [#43567](https://github.com/NousResearch/hermes-agent/pull/43567) merged, fixing the silent discard of `on_pre_compress()` return values so compression summaries now include provider insights.

- **Surface blocked context‑file warnings** — Two PRs ([#59625](https://github.com/NousResearch/hermes-agent/pull/59625) and [#59622](https://github.com/NousResearch/hermes-agent/pull/59622)) are open to address Issue [#59612](https://github.com/NousResearch/hermes-agent/issues/59612), making threat‑scanner blocks visible to the user.

## 4. Community Hot Topics

| Item | Type | Comments | Reactions | Summary |
|------|------|----------|-----------|---------|
| [#56004](https://github.com/NousResearch/hermes-agent/issues/56004) | Issue (open, P2) | 4 | 👍 3 | Qwen3.6 / vLLM multi‑turn thinking context lost on replay — prior‑turn reasoning stripped. User expects maintainers to implement fix via agentic tools. |
| [#56835](https://github.com/NousResearch/hermes-agent/issues/56835) | Issue (open, P3) | 4 | 👍 0 | Desktop client crashes with `ERR_NETWORK_IO_SUSPENDED` after sleep/resume when connected to remote gateway. Affects Debian server + Windows/macOS clients. |
| [#5000](https://github.com/NousResearch/hermes-agent/issues/5000) | Issue (closed) | 4 | 👍 0 | Qodo AntiSlop scan found 27 issues across 10 recent PRs. Closed as `cannot‑reproduce`, but underlying code‑quality concerns remain. |

**Analysis:** The most active discussion centers on [#56004](https://github.com/NousResearch/hermes-agent/issues/56004) (thinking context loss), which received three upvotes and detailed technical context from the reporter. This indicates a significant need for reliable multi‑turn reasoning preservation, especially as models like Qwen3.6 see wider adoption. The network‑resume crash ([#56835](https://github.com/NousResearch/hermes-agent/issues/56835)) also draws engagement, highlighting the importance of session stability across desktop sleep cycles.

## 5. Bugs & Stability

Bugs reported today, ranked by severity (P1 highest):

- **P1 — Telegram polling reconnect deadlock** ([#59614](https://github.com/NousResearch/hermes-agent/issues/59614)): When both primary and fallback Telegram servers become unreachable, `start_polling()` hangs indefinitely because the reconnect ladder lacks a timeout. **Fix in progress:** PR [#59618](https://github.com/NousResearch/hermes-agent/pull/59618) adds a timeout wrapper.

- **P2 — Multi‑turn thinking context lost** ([#56004](https://github.com/NousResearch/hermes-agent/issues/56004)): Qwen3.6 via vLLM strips prior‑turn reasoning on replay. No fix PR yet, but reporter offered to let agentic tools implement the solution.

- **P2 — Silent threat‑scanner block** ([#59612](https://github.com/NousResearch/hermes-agent/issues/59612)): Project context files (AGENTS.md, etc.) are silently replaced when blocked by the threat scanner. **Fix in progress:** PRs [#59625](https://github.com/NousResearch/hermes-agent/pull/59625) and [#59622](https://github.com/NousResearch/hermes-agent/pull/59622) add user‑visible warnings.

- **P2 — Session routing corruption with `tui_shutdown`** ([#59627](https://github.com/NousResearch/hermes-agent/pull/59627) — PR fixing a bug): TUI shutdown end_reason prevents correct session recovery. PR open to treat `tui_shutdown` as recoverable.

- **P3 — Desktop crash after network resume** ([#56835](https://github.com/NousResearch/hermes-agent/issues/56835)): `ERR_NETWORK_IO_SUSPENDED` on WebSocket reconnect. No fix PR yet.

- **P3 — Dashboard false positive** ([#59626](https://github.com/NousResearch/hermes-agent/issues/59626)): `dashboard --status` reports a running server when only a wrapper/sandbox command matches the string.

- **P3 — Cron tick() None lock_fd crash** ([#59623](https://github.com/NousResearch/hermes-agent/pull/59623) — PR fixing a bug): Guard against crash when `open()` fails in finally block.

Additional fixes in open PRs today address Slack `invalid_blocks` ([#59621](https://github.com/NousResearch/hermes-agent/pull/59621)), Daytona env passthrough ([#59620](https://github.com/NousResearch/hermes-agent/pull/59620)), URL credential stripping for fork detection ([#59619](https://github.com/NousResearch/hermes-agent/pull/59619)), and gateway restart‑notification markers ([#59079](https://github.com/NousResearch/hermes-agent/pull/59079)).

## 6. Feature Requests & Roadmap Signals

Notable feature‑type pull requests open or recently merged:

| PR | Feature | Status | Likely next release? |
|----|---------|--------|----------------------|
| [#30055](https://github.com/NousResearch/hermes-agent/pull/30055) | Telegram Business Mode (observe‑with‑approval) | Open, updated today | Possibly not yet – still active development |
| [#36145](https://github.com/NousResearch/hermes-agent/pull/36145) | `/upload-trace` and `hermes trace upload` to HF trace viewer | Open, updated today | High probability – builds on existing Hugging Face integration |
| [#34393](https://github.com/NousResearch/hermes-agent/pull/34393) | Proton Pass as external secret source | Open, updated today | Medium – complements Bitwarden integration |
| [#35129](https://github.com/NousResearch/hermes-agent/pull/35129) | Execute tool calls from model `content` (Ollama/Kimi/MiniMax/Gemma) | **Merged** | Already in main |
| [#59629](https://github.com/NousResearch/hermes-agent/pull/59629) | Image task deadline env var (600s default) | **Merged** | Already in main |
| [#40624](https://github.com/NousResearch/hermes-agent/pull/40624) | Startup banner terminal‑width‑aware skills display | Open, updated today | Likely low priority cosmetic improvement |

The sustained activity on [#30055](https://github.com/NousResearch/hermes-agent/pull/30055) and [#36145](https://github.com/NousResearch/hermes-agent/pull/36145) suggests these features are nearing completion and could land in the next minor release. The addition of Proton Pass secrets ([#34393](https://github.com/NousResearch/hermes-agent/pull/34393)) signals a roadmap focus on flexible secret management.

## 7. User Feedback Summary

Real pain points expressed in today’s issues:

- **Session reliability after network disruption** – Users on remote desktop setups (Debian server + Windows/macOS client) experience crashes on sleep/wake ([#56835](https://github.com/NousResearch/hermes-agent/issues/56835)).
- **Invisible context filtering** – The threat scanner’s silent block of project instructions (AGENTS.md) frustrates users who expect clear feedback when their configuration is overridden ([#59612](https://github.com/NousResearch/hermes-agent/issues/59612)).
- **Multi‑turn thinking loss** – Users of Qwen3.6 models lose historical reasoning when the conversation replays, which undermines multi‑step agent tasks ([#56004](https://github.com/NousResearch/hermes-agent/issues/56004)).
- **Telegram bot hang on network failure** – When Telegram endpoints go down, the bot hangs forever, stopping all message processing ([#59614](https://github.com/NousResearch/hermes-agent/issues/59614)).
- **Fake process detection** – Wrapper commands containing `hermes dashboard --status` cause false positives in process listing ([#59626](https://github.com/NousResearch/hermes-agent/issues/59626)).

Overall, users are actively testing edge cases (network sleep, server unreachability, proxy configurations) and reporting them in detail. The project’s high PR activity suggests maintainers are responsive, with fix PRs already open for the most critical P1‑P2 bugs.

## 8. Backlog Watch

Older, still‑open issues/PRs that may need maintainer attention:

- **Issue [#5000](https://github.com/NousResearch/hermes-agent/issues/5000) (closed Apr 4, reopened? no, closed)** – Qodo AntiSlop scan flagged 27 issues across 10 PRs. While marked `cannot‑reproduce`, the underlying code quality concerns remain unaddressed. Worth a follow‑up audit.

- **PR [#30055](https://github.com/NousResearch/hermes-agent/pull/30055)** – Telegram Business Mode (May 21) – still open with no maintainer comments. Large feature with significant UI implications. If not actively pursued, should be updated with status.

- **PR [#36145](https://github.com/NousResearch/hermes-agent/pull/36145)** – Trace upload (June 1) – still open, updated today. No maintainer review yet. Could be merged soon given the recent update.

- **PR [#34393](https://github.com/NousResearch/hermes-agent/pull/34393)** – Proton Pass secrets (May 29) – open, no maintainer feedback. Adding new secret sources increases the maintenance surface; should be reviewed for security.

- **PR [#40624](https://github.com/NousResearch/hermes-agent/pull/40624)** – Terminal‑width banner (June 6) – open with 4 passing tests but unreviewed. Minor cosmetic change unlikely to cause regressions; easy merge candidate.

- **Issue [#56835](https://github.com/NousResearch/hermes-agent/issues/56835)** (July 2) – Desktop crash on network resume. No fix PR exists. Could benefit from a maintainer investigating the WebSocket reconnect logic.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest — 2026-07-06

## Today's Overview
The project shows steady, collaborative activity: 3 issues and 6 pull requests were updated in the last 24 hours, including a critical bug fix merged for Anthropic prompt caching and a fix for session‑history corruption moving through review. No new release was published today, but the development velocity indicates active maintenance. The community is contributing high‑quality patches, particularly around Anthropic‑provider reliability and tool‑use behavior.

## Releases
*No new releases today.*

## Project Progress
One pull request was merged/closed today:

- **[PR #3227] fix(providers): resolve tool_use name/args from Function on reloaded history** (`AayushGupta16`) — closed 2026-07-06.  
  Fixes a round‑trip bug where `ToolCall.Name` and `ToolCall.Arguments` (marked `json:"-"`) were lost after chat history was reloaded, causing deserialization errors on both Anthropic providers. This fix ensures tool‑use blocks survive session restores.

Additionally, **Issue #2191** (anthropic_messages provider ignores SystemParts, breaking prompt caching) was closed today – resolved by the incoming PR #3228 and related work.

## Community Hot Topics

### Most active issue
- **[Issue #3088] [Feature] Use vodozemac instead of libolm**  
  `@pbsds` | 6 comments | 2 👍 | **Priority: high** | Opened 2026-06-09  
  Discussion centres on replacing the unmaintained `libolm` with the official Rust‑based successor `vodozemac`. Community agrees on the security urgency; no PR yet.  
  [Link](sipeed/picoclaw Issue #3088)

### Most active pull requests
- **[PR #3115] Fix inline data URL media extraction for generic tool output**  
  `@jp39` | updated today | 6 comments? (not shown) | Still open. Addresses a session‑history corruption bug where `data:` URIs in text tool outputs were misidentified as media attachments.  
  [Link](sipeed/picoclaw PR #3115)

- **[PR #3228] fix(anthropic‑messages): send SystemParts as system blocks with cache_control**  
  `@AayushGupta16` | created today | immediate attention. Directly fixes #2191 by allowing per‑block `cache_control` markers on the `anthropic_messages` provider, enabling Anthropic prompt caching for system messages.  
  [Link](sipeed/picoclaw PR #3228)

- **[PR #3226] fix(tools): stop write_file from coaching destructive overwrite**  
  `@ACMYuechen` | updated today | Fixes a behavioral issue where the tool’s documentation pushed the agent toward destructive overwrites instead of offering safe alternatives.  
  [Link](sipeed/picoclaw PR #3226)

### New proposal
- **[Issue #3229] Proposal: rolling conversation cache breakpoints for anthropic‑messages**  
  `@AayushGupta16` | created today | No comments yet, but builds on #3228 to extend caching beyond system prompts to long conversation histories – a key cost‑saver for agentic workloads.  
  [Link](sipeed/picoclaw Issue #3229)

---

## Bugs & Stability

| Severity | Bug | Status | Fix PR |
|----------|-----|--------|--------|
| **Critical** | `tool_use` name/args lost on history reload (both Anthropic providers) | **Fixed** (PR #3227 merged) | #3227 |
| **High** | Data URI strings in tool output cause session‑history corruption (#3115) | **Open** (PR #3115 awaiting merge) | #3115 |
| **High** | `write_file` coaches agent toward destructive overwrite (#3226) | **Open** (PR #3226 in review) | #3226 |
| **High** | `anthropic_messages` ignores SystemParts, kills prompt caching (#2191) | **Closed** – fix in #3228 | #3228 (open) |
| **Low** | Duplicate `build/` entry in `.gitignore` (#3191) | Open (chore PR) | #3191 |

No new crashes or regression reports appeared today. The remaining open bugs have active, targeted PRs.

## Feature Requests & Roadmap Signals

The community’s strongest signals point toward **Anthropic‑provider maturity and security**:

1. **Replace libolm with vodozemac** (#3088, high priority) – the most‑wanted feature, likely to land in the next minor release if maintainers pick up the implementation. No PR exists yet.
2. **Per‑block `cache_control` for `anthropic_messages`** (#3228) – already in PR, expected in next release.
3. **Rolling conversation cache breakpoints** (#3229) – a next‑step proposal after #3228, targeting cost reduction for long, multi‑turn agentic sessions.

These features signal that the user base is shifting toward heavy, production‑grade Anthropic usage where both security (libolm → vodozemac) and cost (caching) are critical.

## User Feedback Summary

- **Pain points:** Users express frustration with `libolm` being unmaintained (security risk) and with the lack of Anthropic prompt caching on the `anthropic_messages` provider, which inflates API costs. The `write_file` tool’s destructive‑overwrite coaching also frustrates those using memory files.
- **Use cases dominating discussion:** Agentic tool‑calling workflows where conversation history dominates input tokens; secure chat applications that depend on Olm/Megolm encryption.
- **Satisfaction indicators:** Quick turnaround on bug fixes (e.g., #2191 → PR #3228 within hours, #3227 merged same day) indicates a responsive maintainer team. Community contributions are high‑quality and well‑targeted.
- **Dissatisfaction:** No clear complaints other than the above feature gaps, which are being actively addressed.

## Backlog Watch

The following items have been open for extended periods and need maintainer attention:

1. **Issue #3088 (Feature: vodozemac) – Open since 2026-06-09**  
   High priority, help‑wanted label, no assigned owner. Implementation requires significant dependency changes; maintainers should either assign or provide guidance.  
   [Link](sipeed/picoclaw Issue #3088)

2. **PR #3115 (data URI fix) – Open since 2026-06-12**  
   A well‑described fix for a session‑history corruption bug. Updated today but not merged; unclear if it needs re‑review or additional tests.  
   [Link](sipeed/picoclaw PR #3115)

3. **PR #3192 (Alpine base image bump) – Open since 2026-06-27**  
   Low risk, trivial update. Should be merged promptly to keep Docker builds consistent.  
   [Link](sipeed/picoclaw PR #3192)

4. **PR #3191 (duplicate `.gitignore` entry) – Open since 2026-06-27**  
   Similarly low‑risk cleanup.  
   [Link](sipeed/picoclaw PR #3191)

The project overall is healthy and well‑maintained; these items are minor hoppers relative to the breakneck pace of fixes and features this week.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-07-06

## 1. Today’s Overview
The project saw light activity over the past 24 hours. Two issues were updated (one closed, one open) and seven pull requests were updated, all remaining open. No new releases were published. The main event was a design proposal for a Zoom voice agent that was reviewed and closed, while a documentation sweep by contributor **glifocat** accounts for the majority of PR updates, aiming to bring in-repo docs back in line with current code. Overall, the project appears stable with a focus on documentation correctness and incremental skill additions.

## 2. Releases
None in the last 24 hours.

## 3. Project Progress
No PRs were merged or closed today. Seven open PRs were updated, indicating ongoing work:

- **#2910** – *docs: update SDK deep-dive from 0.2.x to 0.3.197* (glifocat) — re-verifies documentation against the actual SDK package.  
  [nanocoai/nanoclaw#2964](https://github.com/nanocoai/nanoclaw/pull/2964)

- **#2963** – *docs: rewrite architecture.md and agent-runner-details.md* (glifocat) — rewrites the most stale sections of these documents.  
  [nanocoai/nanoclaw#2963](https://github.com/nanocoai/nanoclaw/pull/2963)

- **#2962** – *docs: sync DB schema and entity docs with migrations 010-018* (glifocat).  
  [nanocoai/nanoclaw#2962](https://github.com/nanocoai/nanoclaw/pull/2962)

- **#2961** – *docs: fix stale claims across README, CONTRIBUTING, CLAUDE.md and operational docs* (glifocat) — a batch of small factual fixes.  
  [nanocoai/nanoclaw#2961](https://github.com/nanocoai/nanoclaw/pull/2961)

- **#2958** – *add-teams: Teams-CLI-first credentials flow in SSF directive grammar* (Koshkoshinsk) — rebuilds the add-teams skill on the structured-skill-format.  
  [nanocoai/nanoclaw#2958](https://github.com/nanocoai/nanoclaw/pull/2958)

- **#2949** – *feat(skill): /add-litellm — minimal model router* (javexed) — adds a utility skill for local model routing.  
  [nanocoai/nanoclaw#2949](https://github.com/nanocoai/nanoclaw/pull/2949)

- **#2909** – *feat(setup): template setup flow in the wizard and first-agent stamping* (amit-shafnir) — second part of agent template support; remains open since July 2.  
  [nanocoai/nanoclaw#2909](https://github.com/nanocoai/nanoclaw/pull/2909)

## 4. Community Hot Topics
Activity levels are low. The most commented item today was:

- **#2960** [CLOSED] – *Proposal: Live Zoom voice agent + K-ai KB integration — review for Kumuda* (vishalsachdev)  
  This design proposal received one comment and was closed on the same day, suggesting it was reviewed and likely accepted or merged for Kumuda. The underlying need is for real-time voice interaction with knowledge-base retrieval in meeting contexts.  
  [nanocoai/nanoclaw#2960](https://github.com/nanocoai/nanoclaw/issues/2960)

No other issues or PRs drew comments or reactions.

## 5. Bugs & Stability
No bugs, crashes, or regressions were reported today. The project remains stable with no stability-related issues under discussion.

## 6. Feature Requests & Roadmap Signals
- **#2959** [OPEN] – *Image generation* (rajpoot713) requests help generating a shop logo (“Dream design make a ascetic logo”). This is likely a user misunderstanding of NanoClaw’s capabilities rather than a genuine feature request, but it may signal a desire for image-generation skills within the platform.  
  [nanocoai/nanoclaw#2959](https://github.com/nanocoai/nanoclaw/issues/2959)

- **#2960** (closed) introduces a complex feature: **Zoom voice agent + KB integration** using Azure OpenAI Realtime API. If adopted (the “review for Kumuda” suggests it was), this could be a roadmap item for upcoming releases, enabling voice-first agent interaction in meetings.

- The ongoing PR **#2909** (template setup flow) is a strong roadmap signal — it completes the second part of agent templates and improves first-run onboarding, likely targeting the next minor release.

## 7. User Feedback Summary
Explicit user feedback is scarce. The only user-initiated issue (#2959) indicates a possible expectation that NanoClaw can generate images, or at least help with graphic design — a gap that may need clearer documentation of what skills are available. The Zoom voice agent proposal (#2960) comes from a contributor (vishalsachdev) and reflects a sophisticated use case, suggesting the community is exploring advanced integrations. No dissatisfaction was expressed.

## 8. Backlog Watch
All open issues and PRs were created or updated within the last five days. There are no long-unanswered or abandoned items that require maintainer attention at this time. The oldest open PR **#2909** (July 2) has been updated recently and is under active discussion. No stale issues were identified.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-07-06

## Today’s Overview
The project experienced minimal activity over the past 24 hours. No new issues were opened or updated, and no releases were published. The only activity is a single open pull request (#956) that bumps the base Alpine image from 3.23 to 3.24, which was last updated today by Dependabot. No code changes or feature work were merged. The project appears to be in a low-activity phase, with maintenance automation (dependency updates) being the sole driver of visible changes.

## Releases
*(No new releases today.)*

## Project Progress
- **Merged/Closed PRs:** 0  
- **Notable advances:** None  

The lone open PR (#956) remains unmerged and does not represent completed work.

## Community Hot Topics
**Only active PR:**  
- [#956 – `[dependencies, docker] ci(deps): bump alpine from 3.23 to 3.24`](https://github.com/nullclaw/nullclaw/pull/956)  
  *Author:* dependabot[bot] | *Updated:* 2026-07-06 | *Comments:* 0 | *Reactions:* 0  

This PR has no comments or reactions; it reflects routine dependency maintenance. There is no community discussion or debate to highlight.

## Bugs & Stability
No bugs, crashes, or regressions were reported in the last 24 hours. No fixes are in flight.

## Feature Requests & Roadmap Signals
No new feature requests were submitted or commented on today. Given the lack of user proposals and the maintenance-only activity, no clear signals indicate where the next version may focus. Continued dependency updates (as seen in #956) represent the only observable roadmap direction.

## User Feedback Summary
No user feedback (comments, reactions, or issue discussions) was recorded in the last 24 hours. The project’s silence does not allow assessment of satisfaction or pain points.

## Backlog Watch
- **Open PR #956** (Alpine bump) has remained open since June 15, 2026 (21 days). Though a low-risk dependency update, its prolonged unmerged status may indicate maintainer bandwidth constraints or a deliberate holding pattern.  
- No other long-unanswered issues or PRs exist in the current dataset.

---

**Overall Project Health:** Low activity; no known regressions or community friction. Dependency maintenance is the only visible signal. Maintainers may benefit from reviewing the stale Alpine bump PR to keep the build environment current.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest – 2026-07-06

## Today's Overview

The IronClaw project remains highly active with 18 issues and 38 pull requests updated in the last 24 hours. Development is concentrated on two major feature stacks: Slack personal OAuth migration (4–7 PRs across the integration) and hosted-single-tenant (HST) Postgres latency improvements (4 stacked PRs). A burst of 8 new performance issues was filed today by a core contributor, targeting inefficient cloning, serialization, and polling in the agent loop, event store, and web UI layers. No new releases were cut; the pending release PR (#5598) remains open. CI reliability received two targeted improvements (canary reporting). Overall, the project is in a high-velocity development phase with significant infrastructure and performance work ongoing.

## Releases

None (no new releases today).

## Project Progress

**Merged/Closed PRs (11 total, 2 detailed):**
- [#5687](https://github.com/nearai/ironclaw/pull/5687) – **CI: comment canary results on triggering PR** (closed). Finalizes live canary report by posting results as GitHub Markdown on the originating PR, keeping the comment path independent from Slack.
- [#5684](https://github.com/nearai/ironclaw/pull/5684) – **CI: include PR context in canary Slack reports** (closed). Adds PR link, branch name, and target SHA to Slack canary reports when triggered manually.

**Closed Issues (4):**
- [#5507](https://github.com/nearai/ironclaw/issues/5507) – [bug_bash_P2] Failed routine run shows "No thread attached" (QA bug – fix applied).
- [#5676](https://github.com/nearai/ironclaw/issues/5676) – [performance] N+1 record fetches in `records_for_scope` (performance defect – fix merged).
- [#5556](https://github.com/nearai/ironclaw/issues/5556) – [bug_bash_P3] Active chat highlight remains in sidebar (UI bug – fix merged).
- [#5555](https://github.com/nearai/ironclaw/issues/5555) – [bug_bash_P2] Terminal floating button overlaps chat composer (UI bug – fix merged).

## Community Hot Topics

**Highest activity by comment count:**
- [#5553](https://github.com/nearai/ironclaw/issues/5553) – Approval notifications disappear (2 comments, P2 bug). Users report that approval prompts for network capabilities vanish from the notification panel, causing lost workflow steps.

**Active feature stacks drawing attention:**
- **Slack OAuth remodel** – A series of 7 stacked PRs (#5643–#5646, #5668, #5670, and the issue #5669) is reshaping Slack integration from pairing codes to personal OAuth, with least-privilege tool scopes. The PRs involve 121+ file changes and are driven by experienced contributors.
- **HST Postgres latency** – A 4-PR stack (#5688–#5691) introducing RootFilesystem-backed row stores to match Postgres latency for single-tenant hosted deployments. Core contributor serrrfirat authored extensive documentation and a standalone latency harness.

**Underlying needs:** The Slack work addresses security (least-privilege OAuth) and user experience (simpler, more granular Slack permissions). The Postgres latency stack targets operational parity for hosted tenants, reducing tail latency from event replay. The cluster of 8 new performance issues (e.g., #5674, #5675, #5677) signals a focused effort to eliminate unnecessary cloning, serialization, and polling across the entire runtime stack.

## Bugs & Stability

**New open bugs (severity: P2/P3):**
- [#5553](https://github.com/nearai/ironclaw/issues/5553) (P2) – Approval notifications disappear instead of remaining in notification history. No fix PR yet. **High impact** for users relying on approval flows.
- [#5554](https://github.com/nearai/ironclaw/issues/5554) (P2) – Mobile chat layout breaks with horizontal overflow. No fix PR yet. **Medium impact** for mobile users.
- [#5557](https://github.com/nearai/ironclaw/issues/5557) (P3) – Logs deep link requires two clicks to open selected conversation. Minor UX friction.
- [#4108](https://github.com/nearai/ironclaw/issues/4108) – Nightly E2E test failure (open since May 27, no comments). **Stability regression** still unresolved.

**Performance bugs (new today, 8 issues):**
Seriously filed by serrrfirat: #5671 (`LeakDetector` rebuilt per string), #5672 (SSE polling inefficient), #5673 (event stream replay on every call), #5674 (tool schema deep-cloned per LLM request), #5675 (conversation state cloned per message), #5677 (WASM linker rebuilt per instantiation), #5678 (agent loop clones per turn), #5679 (events fully deserialized before filter), #5680 (live-progress items cloned per delta). These are **critical for performance** but not user-facing bugs. No fix PRs yet, but the pattern suggests a planned perf optimization sprint.

**Existing bug fixes (closed today):**
- [#5507](https://github.com/nearai/ironclaw/issues/5507) (QA P2) – "No thread attached" on failed runs: no longer blocks debugging.
- [#5555](https://github.com/nearai/ironclaw/issues/5555) (P2) – Terminal button overlap: fixed.
- [#5556](https://github.com/nearai/ironclaw/issues/5556) (P3) – Sidebar highlight: fixed.

## Feature Requests & Roadmap Signals

**Explicit feature requests (from issues):**
- [#5669](https://github.com/nearai/ironclaw/issues/5669) – **Slack least-privilege OAuth**: reduce write scopes for read-only users. Already tracked in the Slack stack (PR #5670).

**Features in active development (PRs):**
- **Slack personal OAuth** (#5644, #5645, #5646) – Transition from pairing codes to OAuth. Expected in next release (stack 4/4 ready).
- **Slack bot channel entrypoint + installable tools** (#5668) – Model-B Slack remodel (stack 5/7). Likely to follow OAuth release.
- **HST Postgres latency parity** (#5688–#5691) – Row-store implementation for turn state and runtime stores. Targeted at hosted tenants.
- **i18n localization** (#5685) – Add translations for Reborn shell, chat, and extensions. Nearing completion.
- **No run-borking failures** (#5692, #4841) – Collapsed recoverability stack to eliminate opaque terminal errors and provide retry/fault explanation. Large, core contributor effort.

**Predictions for next version:** The Slack OAuth + least-privilege scopes are the most mature (stack 4/7 already closing). The HST Postgres work is foundational but likely experimental. The localization PR and mobile overflow fix (PR #5682) are small and could ship quickly.

## User Feedback Summary

**Pain points surfaced in the data:**
- **Approval notification reliability** (#5553): Users cannot consistently approve automation actions, leading to blocked workflows.
- **Mobile experience** (#5554): Chat is unusable on mobile due to horizontal overflow.
- **Debugging blocked runs** (#5507, closed): Previously users could not inspect failed run threads – now fixed.
- **Logs deep link** (#5557): Navigating to a specific conversation from a run requires two clicks, frustrating power users.
- **Chat highlight confusion** (#5556, fixed): Users were misled about active context after navigating away.
- **Terminal button overlap** (#5555, fixed): Obstructed chat input area on desktop.

**Satisfaction indicators:**
- Quick closure of three UI bugs (P2/P3) and one QA bug within 4 days suggests responsive maintenance.
- No negative sentiment evident in issue comments (all brief).

## Backlog Watch

**Long-unanswered important issues:**
- [#4108](https://github.com/nearai/ironclaw/issues/4108) – **Nightly E2E failed** (open since May 27, 0 comments). This is a recurring automation failure with no maintainer engagement. At 40+ days, it represents a **stability risk** and may indicate gaps in CI monitoring or test flakiness. Needs triage.
- [#4841](https://github.com/nearai/ironclaw/issues/4841 as PR) – **No run-borking failures** (PR open since June 13). While the PR is active (updated today), its size (XL) and risk (low) suggest it may take additional cycles to merge. Stale risk if reviewer bandwidth is limited.

**Other old open issues:**
- [#5557](https://github.com/nearai/ironclaw/issues/5557) (P3, Jul 2) – Logs deep link: no activity from maintainers.
- [#5554](https://github.com/nearai/ironclaw/issues/5554) (P2, Jul 2) – Mobile chat overflow: no fix PR yet, but a PR (#5682) addressing it was opened today by italic-jinxin. This should be resolved soon.

**Recommendation:** The team should prioritize triage on #4108 and ensure the canary CI improvements (#5687, #5684) help catch regressions before they become nightly failures.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest – 2026-07-06

## 1. Today's Overview
Today marks one of the most active development days for LobsterAI in recent weeks, with **11 pull requests merged/closed** and only **1 open PR** (a long-standing dependency update). The team focused heavily on the **OpenClaw** agent engine, adding heartbeat cost controls, a heartbeat toggle setting, and OAuth login support for **xAI (Grok)** providers. Several cross-cutting fixes landed as well, including a cowork stale‑sync guard, MCP transport config cleanup, and a redesigned scheduled‑task list. No new releases were published, but the rapid merge cadence suggests a release is imminent.

---

## 2. Releases
**No new releases today.** The last tag remains the previous version. Given the volume of merged feature work (especially around OpenClaw, Cowork, and email), a release may follow within the week.

---

## 3. Project Progress
All merged PRs advanced core areas of the project:

| PR | Area | Description |
|----|------|-------------|
| [#2256](https://github.com/netease-youdao/LobsterAI/pull/2256) | renderer | Fix scheduled-task notification channel “不通知” not saving; fix white screen when deleting active model in settings |
| [#2281](https://github.com/netease-youdao/LobsterAI/pull/2281) | main | Prevent stale final sync from restarting context maintenance after cowork chat errors |
| [#2280](https://github.com/netease-youdao/LobsterAI/pull/2280) | openclaw, docs | Add heartbeat cost-control policy, repair legacy `HEARTBEAT.md` files, strip proactive-heartbeat guidance from bundled templates |
| [#2279](https://github.com/netease-youdao/LobsterAI/pull/2279) | main | Hide bundled xAI plugin from user sync to avoid accidental exposure |
| [#2278](https://github.com/netease-youdao/LobsterAI/pull/2278) | openclaw, cowork, renderer | Add a Settings toggle for OpenClaw heartbeat (enabled by default), persist it in Cowork config |
| [#2277](https://github.com/netease-youdao/LobsterAI/pull/2277) | openclaw, main, renderer | Clear stale MCP server transport config when switching types; fix header/env/arg persistence |
| [#2276](https://github.com/netease-youdao/LobsterAI/pull/2276) | openclaw, main, renderer | Add xAI (Grok) OAuth login support via PKCE with device‑code fallback |
| [#2275](https://github.com/netease-youdao/LobsterAI/pull/2275) | skills, docs, main, renderer | Add multi‑account support for the built-in `imap-smtp-email` skill, including account management UI |
| [#2274](https://github.com/netease-youdao/LobsterAI/pull/2274) | cowork, renderer | Add time‑aware greeting and recent tasks to the Cowork home screen |
| [#2273](https://github.com/netease-youdao/LobsterAI/pull/2273) | scheduledTask, renderer | Redesign task list cards with status chip, toggle, search, and optimistic UI feedback |

Additionally, **PR #2282** (a copy of #2256) was closed without merge, likely as a duplicate.

---

## 4. Community Hot Topics
- **#1277** ([dependabot PR](https://github.com/netease-youdao/LobsterAI/pull/1277)) – an open Electron dependency bump from v40.2.1 to v43.0.0, created April 2. It has not been merged for **over 3 months**, which may raise concerns about version lag or breaking changes. No comments or reactions are recorded.
- No other PRs or issues received comments or reactions today. The high volume of internal team contributions suggests a focused development sprint rather than broad community engagement.

*Underlying need*: The unmerged dependency upgrade could indicate unresolved compatibility issues or a deliberate hold. The community might benefit from a maintainer update on the status.

---

## 5. Bugs & Stability
Three bug fixes landed today, ranked by potential user impact:

| Severity | Bug | Fix PR |
|----------|-----|--------|
| **High** | Scheduled-task notification channel “不通知” not saving, causing tasks to always notify | [#2256](https://github.com/netease-youdao/LobsterAI/pull/2256) |
| **Medium** | Stale cowork chat history sync could erroneously restart context maintenance after errors | [#2281](https://github.com/netease-youdao/LobsterAI/pull/2281) |
| **Medium** | MCP transport config retained stale headers/env/args when switching transport type | [#2277](https://github.com/netease-youdao/LobsterAI/pull/2277) |
| **Low** | White screen when deleting an active model in settings | Covered by [#2256](https://github.com/netease-youdao/LobsterAI/pull/2256) |

No new regressions or crashes were reported. The fixes address persistent UI and state‑management issues.

---

## 6. Feature Requests & Roadmap Signals
The following merged features likely respond to user requests or internal roadmap goals:

| Feature | Likely Motivation |
|---------|-------------------|
| **xAI (Grok) OAuth login** | Users wanted a native Grok provider; adds OAuth flow with fallback. |
| **Heartbeat cost‑control policy & toggle** | Users complained about unnecessary model calls from empty heartbeat files. |
| **Email multi‑account support** | A common request for managing multiple email accounts via the built-in skill. |
| **Cowork time‑aware greeting + recent tasks** | Improves daily workflow visibility; likely influenced by usability studies. |
| **Scheduled‑task list redesign** | Better task management UX (status chips, toggle, search, feedback). |

*Prediction*: These features are likely candidates for the next minor release, possibly alongside the Electron upgrade (once #1277 is resolved).

---

## 7. User Feedback Summary
No explicit user feedback (comments, issues, reactions) was captured today. However, the changes themselves reveal pain points:

- **Heartbeat cost confusion** – the addition of a policy and toggle suggests users were surprised by hidden model invocations.
- **Email configuration friction** – multi‑account support and account management UI address the need to manage multiple email identities without `.env` hacks.
- **Task notification reliability** – the fix for “不通知” indicates a frustrating workflow where disabling notifications didn’t stick.
- **Provider diversity** – xAI/Grok integration broadens model choice, a recurring request in many AI tool communities.

Overall, user satisfaction should improve with today’s batch of fixes and features.

---

## 8. Backlog Watch
- **#1277** – [Open Dependabot PR](https://github.com/netease-youdao/LobsterAI/pull/1277) (Electron `40.2.1` → `43.0.0`). Open since **April 2, 2026** (over 3 months). The update is major (three semver bumps) and may involve breaking Electron API changes. The lack of action is a potential technical debt risk. A maintainer response or merge would be beneficial.

No other long‑standing issues or PRs appear in today’s data.

---

*Data source: GitHub – netease-youdao/LobsterAI. Activity as of 2026-07-06 23:59 UTC.*

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest – 2026-07-06

## 1. Today's Overview
Project activity today was moderate with five pull requests updated, three of which were merged or closed. No new issues were reported, and no new releases were published. The focus was on bug fixes and dependency updates, particularly for MCP OAuth handling, Telegram streaming, and Docker deployment practices. A notable feature bump for WhatsApp integration was also completed. Overall, the project remains stable with steady maintenance efforts.

## 2. Releases
No new versions were released in the last 24 hours. The latest published release remains unchanged.

## 3. Project Progress
Three pull requests were merged or closed today, advancing both fixes and features:

- **#1122 [CLOSED] – fix: drop VOLUME declarations that shadow the home bind mount**  
  Removed Dockerfile `VOLUME` declarations that conflicted with bind-mounting the entire home directory. This resolves a pathological deployment issue where volumes shadowed persistent data.  
  *Author: sayotte | [GitHub](https://github.com/moltis-org/moltis/pull/1122)*

- **#1113 [CLOSED] – hotfix(telegram): stream final replies without completion notify**  
  Restores expected Telegram streaming behavior when completion notifications are disabled. The final answer is now correctly treated as a streamed final reply.  
  *Author: s-salamatov | [GitHub](https://github.com/moltis-org/moltis/pull/1113)*

- **#1144 [CLOSED] – feat(whatsapp): bump whatsapp-rust 0.5 -> 0.6 with LID-native addressing**  
  Upgraded the WhatsApp integration library to use LID-native addressing, fixing inbound message failures after WhatsApp’s device registration migration.  
  *Author: juanlotito | [GitHub](https://github.com/moltis-org/moltis/pull/1144)*

## 4. Community Hot Topics
No issues were updated today, and the open PRs have no comments or reactions. The most notable open discussion revolves around:

- **#1120 [OPEN] – fix(mcp): use direct fetch for resource_metadata URL from WWW-Authenticate**  
  This PR addresses a real-world failure where MCP OAuth authentication breaks with servers like Notion and Linear that include a `resource_metadata` URL in the `WWW-Authenticate` header. Although no comments are recorded, it is a critical interoperability fix that may generate community interest once reviewed.  
  *Author: xzavrel | [GitHub](https://github.com/moltis-org/moltis/pull/1120)*

Additionally, a routine dependency bump PR (#1087) remains open but is inert.

## 5. Bugs & Stability
The following bugs were fixed today, ranked by severity:

- **Medium – MCP OAuth failure with `invalid_target` (PR #1120, still open)**  
  Servers (Notion, Linear) that return `resource_metadata` in the `WWW-Authenticate` header cause `discover_and_register()` to pass an incorrect URL, leading to an OAuth error. A fix is proposed but not yet merged.  

- **Low – Telegram streaming incomplete without completion notify (PR #1113, closed)**  
  When completion notifications are disabled, the final answer was not streamed properly. The hotfix restores correct behavior.  

- **Low – Docker volume shadowing home bind mount (PR #1122, closed)**  
  Explicit `VOLUME` declarations in the Dockerfile prevented bind-mounted home directories from being persisted, causing data loss on container restart. The fix removes those declarations.  

No new crashes, regressions, or security issues were reported today.

## 6. Feature Requests & Roadmap Signals
The only feature-level change today is the WhatsApp protocol upgrade (#1144), which is a necessary adaptation to upstream changes rather than a user-requested feature. No new feature requests appeared in issues or PRs. Based on recent activity, the next minor release is likely to include:

- WhatsApp LID-native addressing (merged)
- Telegram streaming fix (merged)
- Docker volume fix (merged)
- MCP OAuth fix (pending merge of #1120)

No major roadmap shifts are indicated.

## 7. User Feedback Summary
No direct user feedback was captured in issues today. However, the fixes merged indicate real pain points:

- **Telegram users** experienced broken streaming when disabling completion notifications – a configuration some users may prefer for brevity.
- **Docker users** deploying with bind-mounted home directories faced data persistence inconsistencies.
- **WhatsApp users** saw inbound messages fail after device migration – a critical issue for multi-device workflows.
- **MCP users** integrating with Notion or Linear encountered OAuth authentication failures that required manual workarounds.

Overall satisfaction is implied by the rapid closure of several bugs (three fixes merged today).

## 8. Backlog Watch
- **PR #1120 [OPEN] – MCP OAuth fix**  
  Created 2026-06-13, last updated today. No comments or reviews. This is an important bug affecting popular MCP servers and needs maintainer attention to be reviewed and merged.  
  [GitHub](https://github.com/moltis-org/moltis/pull/1120)

- **PR #1087 [OPEN] – Dependency bump `tar` 0.4.45→0.4.46**  
  Opened 2026-05-29, still open. Automated dependabot PRs usually require manual merge; if left unaddressed, it may block other dependency updates.  
  [GitHub](https://github.com/moltis-org/moltis/pull/1087)

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest – 2026-07-06

## 1. Today's Overview

The CoPaw (QwenPaw) project showed very high development activity on July 6, 2026, with **16 issues updated** (14 open, 2 closed) and **45 pull requests updated** (21 open, 24 merged/closed). A **patch release v1.1.12.post3** was published to fix a critical compatibility regression with the ACP protocol. The volume of PRs is notably high, largely driven by a coordinated push of regression tests (12 test-focused PRs from one contributor) and several bug fixes across the backend and frontend. The project continues to balance maintenance on the stable 1.x line while progressing toward the **v2.0.0 beta** (tracked in issue #5273). Overall, the project appears healthy with active community engagement and frequent releases.

## 2. Releases

**v1.1.12.post3** – patch release on the 1.1.12 line.  

**Changelog (from PR #5818):**  
- Pinned `agent-client-protocol` (ACP) to `>=0.9.0,<0.11.0` to resolve a breaking change in ACP that caused import errors (`ImportError: cannot import name 'SetSessionModelResponse' from 'acp'`) for 1.x users.  
- Synced PyInstaller build scripts with the same ACP constraint.  

**Compatibility note for 1.x users:** Users of v1.1.12.post2 and earlier should upgrade to post3 to avoid ACP version conflicts. No breaking changes introduced.

Release page: [https://github.com/agentscope-ai/QwenPaw/releases/tag/v1.1.12.post3](https://github.com/agentscope-ai/QwenPaw/releases/tag/v1.1.12.post3)

## 3. Project Progress

**24 PRs were merged or closed today.** Key advancements:

- **Release and compatibility:** PR #5818 (merged) – prepared v1.1.12.post3 and pinned ACP, fixing a critical import error.
- **Cron CLI enhancement:** PR #5210 (merged) – added `qwenpaw cron update <job_id>` command to modify existing cron jobs via PUT API, eliminating the need for delete+recreate.
- **DingTalk channel stability:** PR #5654 (merged) – surfaces delivery failures for DingTalk proactive/cron sends, skips empty text, raises `ChannelError` appropriately.
- **AI review bot:** PR #5736 (merged) – added a GitHub Actions workflow for automated code review using QwenPaw itself.
- **Timezone bugfix:** PR #5768 (merged) – fixed naive datetime formatting in `AgentMdManager` by adding `timezone.utc` to `datetime.fromtimestamp()` calls.
- **Plugin market link safety:** PR #5750 (merged) – routes plugin market details through the console's external-link guard instead of raw `window.open()`.
- **Runtime tool registration:** PR #5524 (merged) – registers `spawn_subagent` in Runtime 2.0 tool discovery, adds end-to-end test.
- **Streaming API error retry:** PR #5799 (open, but noted) – fixes retry logic for OpenAI `APIError` with body status codes (5xx from custom gateways). Still open but likely to merge soon.
- **Comprehensive regression tests (PRs #5807–#5813):** Six PRs from `hanson-hex` adding unit and contract-guard tests for inbox, channels, approvals, console hooks/stores, API modules, runtime security, and large session rendering. These will improve coverage and prevent regressions.

## 4. Community Hot Topics

| Issue | Title | Comments | 🔗 |
|-------|-------|----------|-----|
| #5757 | [Bug] 飞书信息不回复情况 (Feishu not replying) | 11 | [Issue](https://github.com/agentscope-ai/QwenPaw/issues/5757) |
| #5401 | Console: session with large tool-use history fails to render | 8 | [Issue](https://github.com/agentscope-ai/QwenPaw/issues/5401) |
| #5273 | QwenPaw v2.0.0 Pre-release Bug & Issue Tracker | 5 | [Issue](https://github.com/agentscope-ai/QwenPaw/issues/5273) |
| #5725 | Console 流式输出过程中浏览器卡顿 (browser lag during streaming) | 4 | [Issue](https://github.com/agentscope-ai/QwenPaw/issues/5725) |

**Analysis:**  
- #5757 reports Feishu (Lark) channel not responding after the first message, affecting users of both Docker and AgentScope Platform. This is a reliability issue on a major IM channel; community is actively troubleshooting but no fix PR exists yet.  
- #5401 is a console frontend crash when loading sessions with large tool-use histories. Backend sends `type: "data"` content blocks that the frontend cannot render. No fix PR yet; likely a v2.0 priority.  
- #5273 is the central tracker for v2.0.0 pre-release bugs, currently with 5 comments and 1 reaction. It aggregates issues for the beta.  
- #5725 describes browser lag during streaming output, contrasting with DeepSeek's smooth streaming. Likely a frontend rendering or state management issue – no fix PR yet.

## 5. Bugs & Stability

**Critical severity:**
- **#5816** (CLOSED with release) – `ImportError: cannot import name 'SetSessionModelResponse' from 'acp'`. Immediate break on startup. Fixed by patching ACP pin in v1.1.12.post3.
- **#5789** (OPEN) – Context compression crashes when model output exceeds JSON Schema `maxLength`. The `jsonschema.validate()` fails and halts the pipeline. No fix PR yet.

**High severity:**
- **#5717** (OPEN) – Runtime 2.0: truncated `tool_call.input` causes infinite repeated tool execution. Reproduced with large `write_file` calls. No fix PR yet.
- **#5775** (OPEN) – Auto-memory interval never triggers because `MemoryMiddleware` state is lost across per-request agent rebuilds. Affects v2.0.0b3. No fix PR yet.
- **#5776** (OPEN) – Stale pinned first user message treated as active task in long-lived IM sessions (v2.0.0b3). Context pollution leads to wrong task execution.

**Medium severity:**
- **#5401** – Console frontend crash on large tool-use history (see Community Hot Topics). No fix PR.
- **#5725** – Browser lag during streaming (see Community Hot Topics). No fix PR.
- **#5779** (CLOSED) – Cron state API returns UTC time instead of job's configured timezone. Fixed previously by PR #5768? Actually PR #5768 fixed `AgentMdManager` timezone, but the cron state API (#5779) may still be open? The issue is CLOSED, presumably a fix was merged, though not explicitly listed in today's PRs.

**Lower severity:**
- **#5781** (OPEN) – Offline code mode cannot preview files because resources require online download. Workaround needed.
- **#5795** (OPEN) – Web console chat page does not auto-refresh when new messages arrive via WeChat channel. User requests real-time push.

## 6. Feature Requests & Roadmap Signals

| Issue | Title | 💬 | 🔗 |
|-------|-------|----|-----|
| #5785 | Coding mode cannot select hidden folders (dot-prefixed) | 3 | [Issue](https://github.com/agentscope-ai/QwenPaw/issues/5785) |
| #5780 | Multi-user account management for team use | 2 | [Issue](https://github.com/agentscope-ai/QwenPaw/issues/5780) |
| #5795 | Web console auto-refresh on WeChat new messages | 1 | [Issue](https://github.com/agentscope-ai/QwenPaw/issues/5795) |
| #5567 | QwenPaw GitHub Issue helper Skill (modelscope) | 2 | [Issue](https://github.com/agentscope-ai/QwenPaw/issues/5567) |

**Predictions for next version:**
- **Multi-user account management (#5780)** – A highly requested feature for team deployments. Given the project's focus on IM channels and v2.0 roadmap, this is likely to appear in v2.0 or a subsequent release.
- **Hidden folder selection in coding mode (#5785)** – A small UX improvement that could land in a 1.1.x patch.
- **Web console auto-refresh (#5795)** – Possibly bundled with frontend rework for v2.0 if not addressed sooner.
- **GitHub Issue helper (#5567)** – An external skill rather than core feature, but shows community innovation.

## 7. User Feedback Summary

**Pain points expressed clearly:**
- **Feishu unreliability** (Chinese user, #5757): "First message replies, then subsequent messages are received by the bot but no reply." This is a critical channel reliability issue.
- **Browser lag during streaming** (#5725): User compares to DeepSeek's smooth streaming – "the lag becomes more obvious the longer the answer." Suggests frontend performance regression or missing incremental rendering optimization.
- **Context compression crash** (#5789): User reports structured output validation error causing complete pipeline failure. Affects accuracy of long conversations.
- **Offline code mode** (#5781): User frustrated that file preview requires internet – "offline" mode is not truly offline.
- **Cron timezone confusion** (#5779, now closed): User discovered API returning UTC despite job timezone setting. Now fixed.

**Positive signals:**
- Community members are contributing meaningfully: `manjieqi` added cron update CLI, `hanson-hex` contributed a large test suite, `hehuang139` fixed streaming retry logic, `jinglinpeng` added Node runtime bundling for ACP desktop.
- Users like `tecgic` are creating community skills (#5567) to improve issue reporting quality.
- The release of v1.1.12.post3 was published and accepted quickly, indicating responsive maintainers.

## 8. Backlog Watch

**Issues needing maintainer attention (unresolved for >1 week, no fix PR):**

| Issue | Title | Opened | Last Updated | 🔗 |
|-------|-------|--------|--------------|-----|
| #5401 | Console crash on large tool-use history | 2026-06-23 | 2026-07-06 | [Issue](https://github.com/agentscope-ai/QwenPaw/issues/5401) |
| #5725 | Browser lag during streaming | 2026-07-02 | 2026-07-06 | [Issue](https://github.com/agentscope-ai/QwenPaw/issues/5725) |
| #5789 | Context compression crash on JSON Schema maxLength | 2026-07-05 | 2026-07-06 | [Issue](https://github.com/agentscope-ai/QwenPaw/issues/5789) |
| #5717 | Tool-call truncation causes endless execution (Runtime 2.0) | 2026-07-02 | 2026-07-06 | [Issue](https://github.com/agentscope-ai/QwenPaw/issues/5717) |
| #5775 | Auto-memory never triggers (v2.0.0b3) | 2026-07-04 | 2026-07-06 | [Issue](https://github.com/agentscope-ai/QwenPaw/issues/5775) |
| #5776 | Stale pinned message in IM sessions (v2.0.0b3) | 2026-07-04 | 2026-07-06 | [Issue](https://github.com/agentscope-ai/QwenPaw/issues/5776) |

**PRs awaiting review (open with no recent comment):**
- #5654 (DingTalk delivery fix) – Under Review since June 30, still open.
- #5799 (streaming retry fix) – open since July 6, likely needs final approval.
- #5805 (Tauri DevTools) – new PR, needs review.
- #5814 (Bundle Node runtime for ACP desktop) – new PR, needs review.

**Long-open tracker:**
- #5273 (v2.0.0 pre-release tracker) – opened June 17, still active with 5 comments. While not an individual bug, it represents the overall v2 stability effort. Maintainers should triage remaining child issues there.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-07-06

## 1. Today’s Overview
The ZeroClaw project saw **high activity** with **50 pull requests updated** in the last 24 hours (43 open, 7 merged/closed) and **3 new issues** filed (all open). No new releases were published. The development pace remains strong, with large cross-cutting features (SOP authoring, multi-user auth, schema V4) and security-focused fixes advancing through review. The bug tracker shows two high-severity issues — a broken onboarding config template and a CI gap that allows broken test code to land on `master` — both actively being addressed. Community engagement is evident in the volume of PRs and the respectful developer–user interactions on the open issues.

## 2. Releases
**None.**  
No new versions were tagged or published in the last 24 hours.

## 3. Project Progress
Seven pull requests were merged or closed today. Notable among them:

- **#8726** ([fix(runtime/shell): block dangerous env vars from TUI overlay](https://github.com/zeroclaw-labs/zeroclaw/pull/8726)) – Merged. Adds a static deny-list for environment variables that are never forwarded from TUI client environments, closing a defense-in-depth gap.
- **#8728** ([docs(agents): sync model_provider trait rename in extension points](https://github.com/zeroclaw-labs/zeroclaw/pull/8728)) – Merged. Documentation update to reflect a recent trait rename.
- **#8032** ([fix(web): gate MCP server command/url required-ness on transport](https://github.com/zeroclaw-labs/zeroclaw/pull/8032)) – Closed (merged). Fixes a UI validation bug in the Operator Console config editor for MCP servers.

The remaining merged/closed PRs are not individually listed in the top 20 but include smaller CI and documentation changes.

## 4. Community Hot Topics

- **#6641** – [Feature: Turn-level OTel trace correlation](https://github.com/zeroclaw-labs/zeroclaw/issues/6641) (4 comments)  
  The most-discussed open issue. Author **JordanTheJet** praises maintainer responsiveness and frames this as a natural follow-up to #6009 and #6190. The proposal to nest `llm.call` / `tool.call` / `memory.*` spans under a single turn trace has been accepted, signalling strong community interest in observability improvements.

- **#8718** – [Bug: `zeroclaw config init` ships a config template its own daemon rejects, silently disabling transcription](https://github.com/zeroclaw-labs/zeroclaw/issues/8718) (1 comment)  
  Filed by **lynnkeele**, this is a high-severity onboarding bug that has garnered immediate attention (label `priority:p1`). The community pain point is clear: first-time users get a broken experience.

- **#8746** – [fix(goal): stop active goal self-resume loops](https://github.com/zeroclaw-labs/zeroclaw/pull/8746)  
  A large, cross-cutting fix (tags span 30+ components) addressing a behavioral loop in the goal system. The PR description notes it depends on another fork-based branch, reflecting a complex dependency chain.

- **#8590** – [feat(sop): visual authoring with channel fan-in, selectable executing agent, and tests](https://github.com/zeroclaw-labs/zeroclaw/pull/8590)  
  A massive PR (size XL, risk high) that introduces deterministic SOP workflows. The author explicitly calls for beta testers, indicating the feature is nearing production readiness and the community is engaged.

## 5. Bugs & Stability

Ranked by severity (P1 highest):

| Issue/PR | Description | Status | Risk |
|----------|-------------|--------|------|
| **#8718** | `zeroclaw config init` writes a template with incompatible `max_audio_bytes` and `record_timeout` values, causing `local_whisper` transcription to silently fail on fresh installs. | **Open** (no fix PR yet) | high |
| **#8753** | `rust_quality_gate.sh` omits `--workspace`, so member-crate test targets can fail to compile while the gate passes, allowing broken code to land on `master`. | **Open** (new, no fix PR yet) | high |
| **#8713** (PR) | Fix for `file_download` tool: adds `allowed_private_hosts` opt-in to close the remaining SSRF surface from the July 2026 security audit. | **Open** (PR #8713) | high |
| **#8726** (PR) | Merged fix for TUI environment variable blocking (defense-in-depth). | **Merged** | low-medium |

Two of the three open issues today are bug reports; both are accepted with `priority:p1` or `risk:high`. The CI gap (#8753) is particularly concerning for long-term code quality.

## 6. Feature Requests & Roadmap Signals

The following features appear to be actively under development and likely candidates for the next release:

- **Observability** – Issue #6641 (turn-level OTel trace correlation) is accepted and has a clear implementation path. Expect a PR soon.
- **Multi-User Auth** – PR #8672 ([feat(security): multi-user auth providers, permission profiles, and principal isolation](https://github.com/zeroclaw-labs/zeroclaw/pull/8672)) is a size XL, high-risk enhancement implementing RFC #7141. This is a foundational security feature.
- **SOP Visual Authoring** – PR #8590 is near completion and calling for beta testers. Likely to land in the next minor release.
- **Git Forge Channel** – PR #8609 ([feat(channels): add Git forge channel (GitHub provider) with SOP ingress](https://github.com/zeroclaw-labs/zeroclaw/pull/8609)) adds a new channel integration.
- **Config Schema V4** – PR #8754 ([feat(config)!: schema V4 cut of skills, inert tunable, and summary_model cruft](https://github.com/zeroclaw-labs/zeroclaw/pull/8754)) is a breaking change that removes deprecated config surface. The `!` in the title indicates a breaking semver change.
- **Bocha AI Search** – PR #8737 ([feat(tools): add Bocha AI web search provider](https://github.com/zeroclaw-labs/zeroclaw/pull/8737)) addresses a real user need: none of the existing search providers work from mainland China.

## 7. User Feedback Summary

- **Positive** – In #6641, the author explicitly thanks maintainer `@alexandme` for “excellent responsiveness.” This reflects well on the project’s maintainer engagement.
- **Pain point (Onboarding)** – #8718 reports a direct-to-broken experience for new users running `zeroclaw config init`. The issue is labelled `quickstart` and `severity S2`, highlighting that the first-user journey is currently degraded.
- **Pain point (CI Reliability)** – #8753 (filed by maintainer `alexandme`) shows an internal tooling gap that could silently ship broken tests. While not a user-facing issue, it affects the project’s ability to maintain quality.
- **Use case (China-based deployments)** – PR #8737 explicitly states that all existing web search providers are network-blocked from mainland China. The addition of Bocha AI directly addresses this deployment constraint.

## 8. Backlog Watch

- **Issue #6641** (filed 2026-05-13) – [Feature: Turn-level OTel trace correlation](https://github.com/zeroclaw-labs/zeroclaw/issues/6641)  
  Open for nearly two months with `status:accepted` and a clear implementation plan. While the issue remains active with updates, it has not yet seen a corresponding PR. Given the positive community sentiment and maintainer interest, this should be prioritised.

- **Issue #8032** – This was a long-standing bug (filed 2026-06-19) that was finally closed today. Its long wait before merge may indicate that the `needs-author-action` label can delay fixes. Maintainers should keep an eye on other issues with that label.

No other issues in the backlog appear to have been waiting for extended periods without activity.

---

*Generated from GitHub data: 3 issues, 50 PRs updated in the last 24 hours. No new releases.*

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*