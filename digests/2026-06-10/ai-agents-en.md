# OpenClaw Ecosystem Digest 2026-06-10

> Issues: 258 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-10 05:08 UTC

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

# OpenClaw Project Digest — 2026-06-10

## Today’s Overview
OpenClaw shows exceptionally high activity with **258 issues** and **500 pull requests** updated in the last 24 hours. Of those, 77 issues were closed and 131 PRs were merged or closed, indicating a strong cadence of development and bug fixing. Two new releases were published today (v2026.6.5 and v2026.6.5-beta.6), both carrying critical fixes for QQBot reasoning leak and MCP tool result coercion. However, a backlog of high-severity bugs (especially around message loss, session state corruption, and cross-agent communication) remains under active triage, with many open issues awaiting product decisions or security review.

## Releases
Two versions shipped today:
- **v2026.6.5** (stable)  
- **v2026.6.5-beta.6**  

Both include the same highlights:
- **QQBot**: The bot now strips model reasoning/thinking scaffolding before native delivery, preventing raw `<thinking>` content from leaking into channel replies. (PRs #89913, #90132 — thanks @openperf)  
- **MCP tool results**: Now coerce `resource_link`, `resource`, `audio`, malformed image, and future unknown types to safe fallback formats.  

No known breaking changes or migration notes were included in the release highlights. Users on the beta channel should upgrade to beta.6 to get the QQBot fix.

## Project Progress
Of the 131 merged/closed PRs today, several notable ones appear in the top-30 list:
- **#57831** (closed) – `fix(cron): keep current/session-bound jobs out of direct channel sessions` – Prevents cron jobs from writing back into live direct-message sessions (Feishu/Lark), reducing silent message pollution.  
- **#56023** (closed) – `feat: SkillFoundry quality gates + weighted recall for memory-core` – Adds agent skills (`$anvil`, `$gate-keeper`, `$layer-check`) and ports SkillFoundry capabilities into memory-core for better retrieval quality.  
- **#55757** (closed) – `cli-runner: disable no-output watchdog when stdin is closed` – Fixes a false positive watchdog trigger when providing input to CLI backends.  

Other merged PRs likely include fixes for the `openai/gpt-5.x` transport (#90083) and various WebChat session rotation issues (#70330, #45952).

## Community Hot Topics
| Issue / PR | Comments | 👍 | Summary |
|------------|----------|----|---------|
| [#75 – Linux/Windows Clawdbot Apps](https://github.com/openclaw/openclaw/issues/75) | 109 | 79 | Long-standing feature request for desktop apps on Linux/Windows; created in Jan 2026, still open with high demand. |
| [#25592 – Text between tool calls leaks](https://github.com/openclaw/openclaw/issues/25592) | 30 | 1 | Agent internal processing output (error handling, narration) is sent to messaging channels as visible messages — a UX and security risk. Open, P1, diamond lobster rating. |
| [#52875 – Session_send "no session found"](https://github.com/openclaw/openclaw/issues/52875) | 22 | 1 | Regression after upgrade where agents cannot contact each other; closed today. |
| [#44925 – Subagent completion silently lost](https://github.com/openclaw/openclaw/issues/44925) | 19 | 1 | Multiple failure modes cause subagent orchestration results to be lost without notification. P1, open since March. |
| [#90083 – ChatGPT Responses transport fails with invalid_provider_content_type](https://github.com/openclaw/openclaw/issues/90083) | 16 | 3 | After upgrading to 2026.6.1, `gpt-5.4`/`gpt-5.5` inference fails — closed today, likely fixed in the new release. |
| [#32296 – Agent replies to previous message](https://github.com/openclaw/openclaw/issues/32296) | 15 | 1 | Session context confusion where agent responds to wrong user message. P1, platinum hermit rating. |
| [#58450 – Agent promises follow-up without starting any](https://github.com/openclaw/openclaw/issues/58450) | 14 | 2 | Agent says "I'll check and follow up" but doesn't actually start any action — user trust issue. |

The community is most vocal about missing cross-platform support (#75) and systemic message loss/leak patterns (#25592, #44925, #32296). Security concerns (leaked reasoning, lost subagent work) dominate the most-commented issues.

## Bugs & Stability
### Critical / High Severity (P1)
- **#25592** – Text between tool calls leaks to channels. Still open; no fix PR linked.  
- **#44925** – Subagent completion silently lost (no retry, no notification). Open since March with no fix.  
- **#32296** – Agent responds to previous message (session context confusion). Open, P1, needs live repro.  
- **#90083** – ChatGPT `gpt-5.x` transport broken. **Closed today** — likely fixed in v2026.6.5.  
- **#56733** – Gateway process alive but event loop frozen (all HTTP timeouts). Open, P1, WSL2-specific.  
- **#56488** – Cron `deleteAfterRun` causes repeated re-triggering instead of deletion. Open, P1.  

### Medium / Moderate (P2)
- **#52875** – Regression: `session_send` fails after upgrade – **closed today**.  
- **#57901** – Safeguard compaction ignores `compaction.model` config. Open with linked PR (#? not shown).  
- **#57447** – `sessions_send` blocked by visibility guard even with A2A policy. Open.  
- **#58514** – Google Chat space/group messages silently ignored (DMs work). Open since March.  
- **#56498** – Unable to connect WhatsApp and Telegram accounts. Open, lacks maintainer info.  
- **#58737** – Slack agent display name/avatar reverts on edited messages. Open, regression.  

### Fixes in Progress
- [#85196 – Redact tool output secrets](https://github.com/openclaw/openclaw/pull/85196) – Large PR (XL) aiming to expand secret redaction coverage. Open since May 22, ready for maintainer look.  
- [#57483 – fix(sessions): let a2a policy gate cross-agent sends](https://github.com/openclaw/openclaw/pull/57483) – Directly addresses #57447. Open, P2.  
- [#57831 – fix cron session pollution](https://github.com/openclaw/openclaw/pull/57831) – Merged today.  

Overall, message loss and session state corruption remain the most common stability themes. Several older regressions (March-April) are still open.

## Feature Requests & Roadmap Signals
| Issue | Summary | Likelihood for Next Version |
|-------|---------|----------------------------|
| [#75](https://github.com/openclaw/openclaw/issues/75) | Linux/Windows Clawdbot Apps (109 👍) | High demand but likely longer-term; not yet prioritized. |
| [#10659](https://github.com/openclaw/openclaw/issues/10659) | Masked Secrets – prevent agent from seeing raw API keys (4 👍) | P1, diamond lobster rating; security-critical, may appear soon. |
| [#6615](https://github.com/openclaw/openclaw/issues/6615) | Denylist support for exec-approvals (7 👍) | Well-liked; complements existing allowlist, could ship in 2026.7.x. |
| [#88154](https://github.com/openclaw/openclaw/issues/88154) | Slack Modal Support for interactive workflows (1 👍) | Newer request (May 29). P2, but explicit design would improve UX. |
| [#38626](https://github.com/openclaw/openclaw/issues/38626) | Subagent lifecycle observability + async supervision | P2, off-meta tidepool; comprehensive, but no clear timeline. |
| [#55914](https://github.com/openclaw/openclaw/issues/55914) | `openclaw invite` – shareable invite codes for mobile pairing (1 👍) | Interesting quality-of-life addition; P2. |

We predict **Masked Secrets** (#10659) and **exec-approvals denylist** (#6615) could land in the next minor release given their security impact and community support.

## User Feedback Summary
- **Pain Points**:  
  - **Message loss** is the #1 complaint: subagent results, cron outputs, Google Chat group messages, WhatsApp mid-reconnect messages all vanish silently.  
  - **Session confusion**: agents replying to wrong threads or previous messages, especially in group chats with multiple agents.  
  - **Security**: reasoning scaffolding leaking to channels (#25592), API keys fully accessible to agents (#10659), tool output secrets not redacted.  
  - **Cross-platform**: lack of Linux/Windows desktop apps forces users to rely on terminal/TUI.  
  - **Multi-user setups**: hardcoded `chmod 0o600/0o700` breaks shared containers (#56263).  
- **Satisfaction Indicators**:  
  - Two releases today show responsive maintainers.  
  - Many issues closed promptly (e.g., #90083, #52875, #70164).  
  - PRs like #85196 (secret redaction) demonstrate active security hardening.  
- **Use Cases**:  
  - Heavy use in Telegram group chats with multiple agents.  
  - WhatsApp sticker support requested (#7476).  
  - Docker/CI automated deployment blocked by TTY-dependent auth (#56650).  

## Backlog Watch
The following issues and PRs have been open for extended periods (≥2 months) with high priority/impact and no merged fix:
- **#75** (Jan 1) – Linux/Windows apps – 109 comments, 79 👍, no fix PR. Needs product decision.  
- **#10659** (Feb 2) – Masked Secrets – P1, diamond lobster, no fix PR.  
- **#6731** (Feb 2) – Safe/unsafe ClawdBot – No fix PR, needs security review.  
- **#10687** (Feb 6) – Dynamic model discovery (OpenRouter) – P2, no fix PR.  
- **#6615** (Feb 1) – Exec-approvals denylist – P2, needs product decision.  
- **#25592** (Feb 24) – Text between tool calls leaks – P1, no fix PR, needs product decision and security review.  
- **#57901** (Mar 30) – Safeguard compaction ignores config – P2, linked PR open but not merged.  
- **#85196** (May 22) – Redact tool output secrets – Large PR awaiting maintainer review.  

These items represent the largest gap between community demand and delivery. Maintainers should prioritize the P1/P2 issues with security and message-loss impact for the next sprint.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report — 2026-06-10

## 1. Ecosystem Overview

The personal AI agent open-source landscape is experiencing a **rapid maturation phase** characterized by intense bug-fixing, security hardening, and platform expansion. Major projects are shipping daily releases or significant feature branches, while community feedback is converging around a shared set of pain points: message loss, session context corruption, token inefficiency, and lack of cross-platform desktop support. The ecosystem is bifurcating between **platform-level orchestrators** (OpenClaw, ZeroClaw, IronClaw) and **lightweight, embeddable agent frameworks** (NanoBot, NullClaw, PicoClaw), with several projects actively adopting **plugin/skill marketplaces** and **multi-runtime abstractions** to avoid vendor lock-in. Overall, the space is healthy, with maintainers responding quickly to regressions and the community contributing substantial security audits and feature requests.

---

## 2. Activity Comparison

| Project | Issues Updated (last 24h) | PRs Updated (last 24h) | New Releases | Health Score | Notes |
|---------|---------------------------|-------------------------|--------------|--------------|-------|
| **OpenClaw** | 258 (77 closed) | 500 (131 merged) | **2** (v2026.6.5, beta.6) | **High** | Extreme throughput; backlog of P1 bugs |
| **ZeroClaw** | 24 (1 closed) | 50 (3 merged) | 0 | **High** | Heavy feature work; MCP and provider refactoring |
| **CoPaw** | 25 (12 closed) | 35 (16 merged) | **1** (v1.1.11-beta.2) | **High** | Rapid bug fixing; first-time contributor friendly |
| **IronClaw** | 7 (1 closed) | 50 (4 merged) | 0 | **High** | Feature-heavy (Reborn v2 attachments, i18n) |
| **Hermes Agent** | 15 (0 closed) | 50 (18 merged) | 0 | **Moderate** | Very active PRs; some open bugs lack maintainer response |
| **NanoClaw** | 1 (0 closed) | 43 (39 merged) | 0 | **High** | Massive backlog clearing; major features merged |
| **PicoClaw** | 21 (0 closed) | 16 (6 merged) | **1** (nightly) | **Moderate** | Security audit dominated; fixes in flight |
| **NanoBot** | 7 (0 closed) | 27 (5 merged) | 0 | **Moderate** | Regression from 0.2.0; stream timeout critical |
| **NullClaw** | 5 (4 closed) | 8 (7 merged) | 0 | **High** | All bugs except one fixed within 24h |
| **LobsterAI** | 1 (0 closed) | 4 (4 merged) | 0 | **Moderate** | Consolidation; cross-model orchestration gap |
| **Moltis** | 1 (0 closed) | 0 | 0 | **Low** | Minimal activity; one open config bug |
| **TinyClaw** | 0 | 0 | 0 | **Low** | No activity |
| **ZeptoClaw** | 0 | 0 | 0 | **Low** | No activity |

*Health Score: Based on issue closure rate, PR merge velocity, responsiveness to critical bugs, and release cadence.*

---

## 3. OpenClaw's Position

**Advantages over peers:**

- **Largest community and activity**: 258 issues and 500 PRs in 24 hours dwarfs all others; 2 releases in one day demonstrates unparalleled maintainer capacity.
- **Core reference implementation**: As the foundational project, it sets standards for agent communication, session management, and channel integration. Other projects (e.g., NullClaw, PicoClaw) explicitly borrow patterns from OpenClaw.
- **Broadest channel support**: QQBot, Feishu/Lark, Slack, Telegram, Google Chat, WhatsApp – coverage exceeds most peers.
- **Active security hardening**: Secret redaction PR (#85196) and reasoning leak fixes (#89913) show proactive security stance.

**Technical approach differences:**

- **Session-first architecture**: OpenClaw treats sessions as first-class entities with explicit lifecycle management, unlike NanoBot’s simpler `history.jsonl` approach. This enables rich cross-agent communication but also introduces complexity (session corruption bugs #32296, #44925).
- **Monolithic vs. modular**: OpenClaw is a monolithic core, whereas NanoClaw and NullClaw adopt plugin/skill marketplaces. OpenClaw’s `SkillFoundry` (#56023) is recent, indicating a pivot toward modularity.
- **Community size**: OpenClaw’s issue/PR volume is ~5–10× that of Hermes or ZeroClaw, reflecting the largest user and contributor base.

**Weaknesses:**

- **Backlog of P1 bugs**: Message loss (#25592, #44925) and session confusion (#32296) are open for weeks with no fix PRs. User trust is eroding.
- **Desktop support gap**: No native Linux/Windows apps (#75), while CoPaw has Tauri desktop and Hermes has desktop app with VS Code theme integration.
- **Token inefficiency**: No lazy tool loading or context compression (unlike Hermes #6839 and CoPaw #5063).

---

## 4. Shared Technical Focus Areas

| Focus Area | Projects Involved | Specific Needs |
|------------|-------------------|----------------|
| **Session context corruption / message loss** | OpenClaw, NanoBot, Hermes, ZeroClaw | Subagent results lost, cron outputs vanish, agent replies to wrong message, cross-session pollution |
| **Tool call reliability & token efficiency** | Hermes, PicoClaw, CoPaw, OpenClaw | Lazy tool schema loading, streaming config, empty response retries, tool call format parsing |
| **Security hardening** | PicoClaw, OpenClaw, NanoClaw, ZeroClaw | SSRF bypass, secret redaction, reasoning leak, approval gating, predictable pairing codes |
| **Provider compatibility & multi-runtime** | NanoBot, NullClaw, ZeroClaw, IronClaw, LobsterAI | OpenAI-compatible custom providers, per-conversation model override, cross-model orchestration |
| **Multi-platform desktop & mobile** | OpenClaw (#75), CoPaw (Tauri), Hermes (desktop app), PicoClaw | Native apps, system tray, startup performance, platform-specific auth (macOS Feishu, Windows CA) |
| **Scheduled tasks / cron reliability** | OpenClaw, NullClaw, CoPaw, ZeroClaw | Agent-type cron jobs not spawning, cron session pollution, delete-after-run loops |
| **Memory & context weighting** | ZeroClaw, NullClaw, Hermes | Memory overemphasis in system prompt, cross-agent memory sync, unwanted auto-saves |

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | ZeroClaw | CoPaw | Hermes | NanoBot | NullClaw | PicoClaw | IronClaw |
|-----------|----------|----------|-------|--------|---------|----------|----------|----------|
| **Primary use case** | General-purpose agent orchestration | OS-level agent integration (SOP, AMQP) | Chinese desktop agent (Tauri) | Developer agent with skill learning | Lightweight headless assistant | Minimalist agent with reliable cron | IoT/edge agent with security focus | NEAR AI cloud agent (TEE, Reborn v2) |
| **Target user** | DevOps, advanced users | Enterprise, multi-tenant | Chinese-speaking desktop users | Developers, local model enthusiasts | Hobbyists, CLI-first | Cron-task heavy users | Security-conscious, embedded | Cloud/NEAR ecosystem |
| **Architecture** | Monolithic core with channels | Plugin-based (skills, SOPs) | Python backend + Tauri frontend | Agent skills + gateway | Hook-based event system | Minimal dependencies | Provider-agnostic, nightly builds | Rust-based Reborn architecture |
| **Desktop support** | None (only TUI) | None (focus on server) | **Yes** (Windows Tauri) | **Yes** (Desktop app) | None | None | None | None (web UI only) |
| **Community engagement** | Very high (500 PRs/day) | High (50 PRs/day) | High (35 PRs/day) | Moderate (50 PRs/day) | Low (27 PRs/day) | Low (8 PRs/day) | Moderate (16 PRs/day) | High (50 PRs/day) |
| **Key strength** | Breadth of integrations, rapid releases | Multi-tenant RBAC, SOP automation | Out-of-box Chinese UX, fast bug fixes | Token efficiency demand, learning loops | Simplicity, small footprint | Cross-memory, reliable cron | Proactive security audit | Attachment pipeline, i18n completeness |
| **Key weakness** | Backlog of P1 bugs | Unaddressed provider issues | Windows startup regression | Lazy tool loading unimplemented | Stream timeout regression | Cron job bug unfixed | Security audit volume | NEAR provider config save broken |

---

## 6. Community Momentum & Maturity

**Tier 1 – Rapidly iterating (daily releases, high velocity):**
- **OpenClaw**: Exceptional release frequency (2 per day), but strain shows in backlog of critical bugs. Mature project with largest ecosystem.
- **CoPaw**: Fast bug-fix turnaround, welcoming first-time contributors. Beta release today; expects stable v1.2 soon.
- **ZeroClaw**: Heavy feature PRs (provider refactor, RBAC, MCP). No release today but clear roadmap.
- **IronClaw**: 50 PRs in 24h, focused on major Reborn v2 features. Release blocked by PR #3708.
- **NanoClaw**: 39 PRs merged in one day – likely preparing for a major release (skill marketplace, direct runner mode).

**Tier 2 – Actively developing (multiple PRs per day, occasional releases):**
- **NanoBot**: Moderate velocity, but 0.2.0 regression undermines user trust. Stream timeout (#4013) is blocker.
- **Hermes Agent**: High PR volume (50) but many open bugs lack maintainer response. Token efficiency feature (#6839) is two months old.
- **PicoClaw**: Security audit creates urgency; fixes are in flight. Nightly releases keep community engaged.

**Tier 3 – Stabilizing or quiet:**
- **NullClaw**: All but one bug fixed within 24h – maturing codebase with low churn.
- **LobsterAI**: Consolidation phase; cross-model orchestration gap may drive next feature cycle.
- **Moltis, TinyClaw, ZeptoClaw**: Minimal to no activity; may be dormant or in pre-release planning.

---

## 7. Trend Signals

1. **Security as a first-class concern**: The PicoClaw security audit (14 issues in one day) and OpenClaw’s reasoning leak fix (#89913) signal a shift from "move fast" to "ship safe." Expect more SSRF, CSRF, and secret redaction work across all projects.

2. **Token efficiency is the new performance battleground**: Users demand lazy tool loading (Hermes #6839), context compression (CoPaw #5063), and streaming config (PicoClaw #2404). Projects that deliver 50–95% token savings will win power users.

3. **Multi-runtime / provider abstraction**: No single LLM backend dominates. Projects are adding pluggable runtimes (NanoClaw #1690, NullClaw #947, ZeroClaw #5937) to avoid vendor lock-in and support local/edge models.

4. **Desktop and mobile are the next frontier**: CoPaw’s Tauri migration, Hermes’ desktop app, and OpenClaw’s long-standing #75 request show that CLI/TUI is insufficient for mainstream adoption. Expect more cross-platform desktop releases by year-end.

5. **Agent self-improvement loops**: CoPaw community explicitly calls out Hermes’ "Learning Loop" (#5017) as a must-have. Projects that combine memory, skill generation, and retry logic will define the next generation of "autonomous" agents.

6. **Enterprise and multi-tenant needs growing**: ZeroClaw’s RBAC (#5982), per-skill permissions (#5775), and AMQP SOP channel (#7369) indicate that agents are moving beyond single-user hobbyist use into production environments.

7. **Stability regression sensitivity**: Users upgrade eagerly but penalize regressions harshly (NanoBot 0.2.0, CoPaw Tauri startup). Projects must invest in regression testing and rollback mechanisms to maintain trust.

**For AI agent developers**: The ecosystem is converging on modularity (skills, plugins, runtimes), token optimization, and security-first design. The biggest immediate value opportunity lies in **session context isolation** and **efficient tool injection** – both are universally requested and under-delivered across top projects.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest – 2026-06-10

## 1. Today's Overview

The project shows high activity: 7 open issues and 27 PRs updated in the last 24 hours, with 12 PRs merged or closed and 15 still open. No new releases were published. The community is actively reporting regressions from the recent 0.2.0 release, particularly around stream timeouts and context pollution, while maintainers are responding with targeted fixes. Several infrastructure improvements (HookCenter, memory lifecycle testing, documentation) have been landed, and a steady stream of provider-specific bug fixes (GPT-5.x, DeepSeek, OpenAI-compatible tool calls) indicates strong cross-provider compatibility attention.

## 2. Releases

**None.** No new releases were published in the last 24 hours. The latest known release remains 0.2.0, which is the version implicated in several recent bug reports.

## 3. Project Progress

Five PRs were merged or closed in the last day:

- **[#3564 – feat(hooks): HookCenter typed-event hook system with plugin support](https://github.com/HKUDS/nanobot/pull/3564)** [CLOSED]  
  Landed a major refactor of the hook system, replacing `AgentHook` method-override patterns with a typed-event architecture. External developers can now distribute hook plugins via `entry_points(group="nanobot.hooks")` with observe/transform/guard modes. Backward compatibility is maintained via an adapter.

- **[#4263 – fix(providers): use max_completion_tokens for GPT-5.x and reasoning models](https://github.com/HKUDS/nanobot/pull/4263)** [CLOSED]  
  Fixes bug #4261 by adding model-name-based fallback so that GPT-5.x, o1, o3, o4 series automatically use `max_completion_tokens` even when the provider spec does not declare `supports_max_completion_tokens`.

- **[#4208 – feat(webui): add assistant reply fork-from-here](https://github.com/HKUDS/nanobot/pull/4208)** [CLOSED]  
  Introduces a simple “Fork from here” button for completed assistant replies in the WebUI, creating a new chat containing the conversation prefix through that reply.

- **[#4177 – docs: make onboarding friendlier for beginners](https://github.com/HKUDS/nanobot/pull/4177)** [CLOSED]  
  Reworked documentation entry points to help users choose the right path (no-background setup, CLI, WebUI, provider recipes, deployment, contributor docs). Added a CLI command chooser, configuration task map, and deployment readiness guides.

- **[#4265 – feat(english-read): change cron schedule from daily to every 2 days](https://github.com/HKUDS/nanobot/pull/4265)** [CLOSED]  
  Adjusted a skill’s cron schedule from daily to every two days.

## 4. Community Hot Topics

The two most commented issues (3 comments each) reflect distinct user pain points:

- **[#4013 – [bug, question] Error calling LLM: stream stalled for more than 90 seconds](https://github.com/HKUDS/nanobot/issues/4013)**  
  Author reports that after upgrading from 0.1.5post2 to 0.2.0, long streams stall after 90 seconds, rendering real work impossible. Comments suggest the timeout is hardcoded and related to `/goal`. This is the highest-severity user-facing regression.

- **[#4253 – [enhancement] support overriding model per conversation](https://github.com/HKUDS/nanobot/issues/4253)**  
  User works with two model presets (fast cloud vs private local) and wants to switch per conversation based on privacy/speed needs. Currently only a global setting exists. This feature request has been upvoted and discussed as a practical workflow improvement.

Other notable discussions:  
- [#4259 – `history.jsonl` cross-session context pollution](https://github.com/HKUDS/nanobot/issues/4259) (2 comments) – a detailed Chinese-language bug report with code-level analysis showing that session summaries leak between conversations.  
- [#4264 – idleCompact should use complete session history](https://github.com/HKUDS/nanobot/issues/4264) (0 comments but filed with clear repro) – describes how the current idle-compaction excludes the last 8 messages, causing incorrect summarization.

## 5. Bugs & Stability

**Critical**  
- **[#4013 – Stream stalled for >90 seconds](https://github.com/HKUDS/nanobot/issues/4013)** – Hardcoded timeout after upgrading to 0.2.0. Blocks all long-running LLM interactions. No fix PR yet; community is waiting for maintainer response.  
- **[#4259 – Cross-session context pollution in `history.jsonl`](https://github.com/HKUDS/nanobot/issues/4259)** – The `ContextBuilder.build_system_prompt()` injects summaries from all sessions, not just the current one, causing context contamination. Detailed root cause analysis provided. No fix merged yet.

**High**  
- **[#4264 – idleCompact removes last 8 messages before summarization](https://github.com/HKUDS/nanobot/issues/4264)** – User corrections in the final messages are excluded from the summary, leaving false conclusions. A fix PR is open: [#4270 – fix: archive full session history in idle compact](https://github.com/HKUDS/nanobot/pull/4270).  
- **[#4061 – OpenAI-compatible text-format tool calls not parsed](https://github.com/HKUDS/nanobot/issues/4061)** – Some providers emit tool calls as plain text instead of structured `tool_calls`, causing raw markup display and tool dispatch failure. Root cause identified but no fix PR yet.

**Medium**  
- **[#4261 – GPT-5.x expects `max_completion_tokens` not `max_tokens`](https://github.com/HKUDS/nanobot/issues/4261)** – Already fixed by merged PR #4263. Also a second open PR [#4268](https://github.com/HKUDS/nanobot/pull/4268) provides an alternative approach.

**Low**  
- [#4267 – WebUI silently drops assistant replies](https://github.com/HKUDS/nanobot/pull/4267) – Intermittent rendering loss, fix PR open.  
- Several open PRs addressing security (symlink escape [#4119](https://github.com/HKUDS/nanobot/pull/4119), read-only root isolation [#4053](https://github.com/HKUDS/nanobot/pull/4053)) are in review.

## 6. Feature Requests & Roadmap Signals

- **Per-conversation model override** [#4253](https://github.com/HKUDS/nanobot/issues/4253) – High demand; likely to be implemented in an upcoming minor release given its simplicity and utility.  
- **botIcon support at agent mode startup** [#4262](https://github.com/HKUDS/nanobot/issues/4262) – Small UX improvement, could ship in next patch.  
- **On-demand version check (no background polling)** – PR [#4255](https://github.com/HKUDS/nanobot/pull/4255) refactors the Settings > About page; design aligns with user privacy preferences.  
- **StepFun ASR transcription provider** – PR [#4260](https://github.com/HKUDS/nanobot/pull/4260) adds a new SSE-based speech recognition provider, indicating interest in voice input support.  
- **Better final responses after max tool iterations** – PR [#4269](https://github.com/HKUDS/nanobot/pull/4269) improves agent output when tool budget is exhausted, a quality-of-life fix.

These signals point to a roadmap focused on provider flexibility (model switching, new ASR), user experience polish (forking, version check, botIcon), and stability (timeout, compaction, context isolation).

## 7. User Feedback Summary

**Positive**  
- User mxnbf (#4013) explicitly praised version 0.1.5post2: “it’s been very good (way to say ty).”

**Pain Points**  
- The same user expressed strong frustration with the 0.2.0 regression: “this renders any real work useless” – the hardware timeout makes the tool unusable for long tasks.  
- Multiple users (chxuan #4259, imkuang #4264, hamb1y #4061) reported subtle but critical bugs that lead to incorrect outputs (wrong summaries, tool call failures, context pollution).  
- User rombert (#4253) described a clear two-model workflow that is currently unsupported, indicating a gap between advanced user needs and single-model design.

**Overall sentiment**: The 0.2.0 release brought new features (HookCenter, forking, documentation) but introduced regressions that impact daily productivity. Users value the project’s rapid development but need more regression testing before major releases.

## 8. Backlog Watch

- **[#4013 – Stream stalled error](https://github.com/HKUDS/nanobot/issues/4013)** – Opened 2026-05-26, 15 days ago, no official fix or workaround posted by maintainers. Highest priority for project health.  
- **[#4061 – Unparsed text tool calls](https://github.com/HKUDS/nanobot/issues/4061)** – Opened 2026-05-29, 12 days ago. Root cause identified but no PR yet.  
- **[#3869 – DeepSeek message hardening PR](https://github.com/HKUDS/nanobot/pull/3869)** – Opened 2026-05-16, 25 days old, still open with no merge. This PR fixes null content errors, empty placeholder leakage, and assistant text loss for DeepSeek providers.  
- **[#3983 – Runner blocked tool-call finish reasons test](https://github.com/HKUDS/nanobot/pull/3983)** and **[#3982 – Scripted agent runner harness](https://github.com/HKUDS/nanobot/pull/3982)** – Both opened 2026-05-24, still open. These are test infrastructure PRs that would improve code quality but need review.

These items represent gaps in maintainer attention that could lead to continued user frustration or missed quality improvements.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-06-10

## 1. Today’s Overview

Hermes Agent saw a **very high activity day** with 50 pull requests and 15 issues updated in the last 24 hours. The team merged or closed 18 PRs, addressing long-standing configuration gaps (`docker_extra_args` bridging), gateway lifecycle stability under Docker restarts, and user-facing features such as VS Code theme installation and Russian localization. Open issue counts remain high (11 active bugs/features), but the volume of incoming PRs indicates active development. The most discussed topic remains the **lazy tool schema loading feature** (#6839), reflecting ongoing user concern over token efficiency on local models.

## 2. Releases

No new releases were published today.

## 3. Project Progress — Merged/Closed PRs Today

18 pull requests were closed or merged. The most significant are:

- **Docker `docker_extra_args` bridging cluster**: Three PRs (#28888, #28891, #30887, #39683) finally closed — all fixing the same bug where `terminal.docker_extra_args` from `config.yaml` was silently dropped in CLI and gateway startup. This resolves a reported security hardening gap (e.g., `--read-only` and `--security-opt` flags).  
  (PR #28888: [link](https://github.com/NousResearch/hermes-agent/pull/28888))

- **Gateway Docker restart persistence** (#33619 closed): Preserves `running` state in `gateway_state.json` across container restarts, preventing silent channel outage after `docker compose up -d --force-recreate`.  
  ([PR #33619](https://github.com/NousResearch/hermes-agent/pull/33619))

- **Kanban archived board guard** (#43286 closed): Refuses to create database stubs for archived boards, preventing resurrection in the Web UI.  
  ([PR #43286](https://github.com/NousResearch/hermes-agent/pull/43286))

- **Desktop VS Code theme installation** (#43292 merged): Users can now install any theme from the VS Code Marketplace directly from the desktop app.  
  ([PR #43292](https://github.com/NousResearch/hermes-agent/pull/43292))

- **Russian localization** (#39335 merged): Full i18n for Hermes Desktop settings, sidebar, statusbar, and gateway panel.  
  ([PR #39335](https://github.com/NousResearch/hermes-agent/pull/39335))

- **Memory/Skills write approval gate** (#38199 merged): Introduces `write_mode` tri-state gate (`on`, `approve`, `deny`) for memory and skill writes, addressing unprompted “wrong assumption” saves by background self-improvement forks.  
  ([PR #38199](https://github.com/NousResearch/hermes-agent/pull/38199))

- **One-shot CLI query cleanup** (#43036 closed): Plugin session finalization now runs before lease release, preventing stale locks.  
  ([PR #43036](https://github.com/NousResearch/hermes-agent/pull/43036))

- **Desktop styling and theming** (#40399 closed): Feature request resolved (likely part of #43292 or earlier work).  
  ([Issue #40399](https://github.com/NousResearch/hermes-agent/issues/40399))

- **Gateway auto-start after container restart** (#42675 closed): Root cause fixed — all shutdown paths now write `planned-stop` markers so that unexpected signals (container restart) leave the gateway ready to auto-start.  
  ([Issue #42675](https://github.com/NousResearch/hermes-agent/issues/42675))

- **Kanban access control** (#43312 closed): Feature request for user-level access control in Web UI; closed, likely merged or deferred.  
  ([Issue #43312](https://github.com/NousResearch/hermes-agent/issues/43312))

## 4. Community Hot Topics

**Most Active Issue:**  
**[#6839 – Feature: Lazy Tool Schema Loading — Two-Pass Tool Injection to Reduce Token Overhead](https://github.com/NousResearch/hermes-agent/issues/6839)**  
- 23 comments, 13 👍  
- Created April 9, updated today. The community strongly desires a smarter tool injection mechanism: instead of dumping all 50+ tool schemas (consuming ~3,500–5,000 tokens/call), the agent should inject only schemas relevant to the current conversation via a two-pass approach. This is especially critical for local model users who pay token costs and for latency-sensitive workflows.

**Other notable discussions:**  
- **[#40416 – Context compaction visually deletes messages on Telegram](https://github.com/NousResearch/hermes-agent/issues/40416)** – 3 comments, 1 👍. Users report a jarring UX where compacted messages vanish from the chat view.  
- **[#43245 – Models ignore PTY terminal tool when needing sudo](https://github.com/NousResearch/hermes-agent/issues/43245)** – 2 comments. Despite explicit `sudo` skills, models try creative alternative paths, bypassing Hermes’ intended sudo mechanism.

**Hottest Open PR:**  
**[#43306 – New skill: qca-cycle — persistent identity & cognitive memory](https://github.com/NousResearch/hermes-agent/pull/43306)** – A new `cognitive` skill category offering a signed personality kernel, graph memory, and novelty gate. This is a large, novel contribution with potential for community debate.

## 5. Bugs & Stability

**Bugs reported today (open):**

| Severity | Issue | Description | Fix PR exists? |
|----------|-------|-------------|----------------|
| **P2** | [#40416](https://github.com/NousResearch/hermes-agent/issues/40416) | Context compaction visually deletes messages on Telegram | No |
| **P2** | [#43232](https://github.com/NousResearch/hermes-agent/issues/43232) | Dashboard enters D (uninterruptible sleep) on startup with NAS-mounted `HERMES_HOME` (NFS/SMB) | No |
| **P2** | [#43228](https://github.com/NousResearch/hermes-agent/issues/43228) | Feishu message footer model name polluted by auxiliary channel calls (vision/compression) | No |
| **P2** | [#43295](https://github.com/NousResearch/hermes-agent/issues/43295) | Docker entrypoint ignores compose command arguments for profile-specific gateway startup | **Yes** – [#43314](https://github.com/NousResearch/hermes-agent/pull/43314) (open) |
| **P2** | [#43296](https://github.com/NousResearch/hermes-agent/issues/43296) | `gateway run --replace` causes Feishu WebSocket auth conflict on macOS | No |
| **P2** | [#42517](https://github.com/NousResearch/hermes-agent/issues/42517) | Gateway `ExecStop` should write planned-stop marker instead of inferring signal source | No |
| **P3** | [#43309](https://github.com/NousResearch/hermes-agent/issues/43309) | Cron jobs section jumps on zoom – sidebar rework regression | **Yes** – [#43310](https://github.com/NousResearch/hermes-agent/pull/43310) (open) |
| **P3** | [#43245](https://github.com/NousResearch/hermes-agent/issues/43245) | Models ignore PTY terminal tool for sudo | No |
| **P3** | [#43294](https://github.com/NousResearch/hermes-agent/issues/43294) | Corporate CA trust not auto-activated on Windows; manual wiring required | No |
| **P3** | [#43248](https://github.com/NousResearch/hermes-agent/issues/43248) *(closed)* | Content zoom reduces font rendering clarity (bilinear downscaling) | Closed (likely fixed by other means) |

**Analysis:** The highest-impact bugs are the **Telegram compaction UX** (#40416) and the **Dashboard D state under NAS** (#43232), both P2 and both lacking fix PRs. The Docker entrypoint bug (#43295) already has an open fix. The Feishu auth conflict (#43296) on macOS needs attention as it affects a growing user base. The cron zoom regression (#43309) has a candidate fix (#43310).

## 6. Feature Requests & Roadmap Signals

**High-interest user-requested features:**

- **Lazy Tool Schema Loading** (#6839 – 13 👍) – Top feature request; would drastically reduce token consumption for local models. Likely on the roadmap, but no PR yet.  
- **Chinese CLI language support** (#43307) – User asks for a command to switch CLI display language to Chinese. Low scope (<300 lines), may be implemented by the community.  
- **Corporate CA trust auto-activation on Windows** (#43294) – Enterprise users need automatic trust for Python-side CA certificates; currently requires manual wiring.  
- **Per-task model/provider routing for delegate subagents** (PR #36790, still open) – Enables different models within the same batch, resolving multiple linked issues (#6306, #10995, #5012). Likely to land in next minor release.  
- **qca-cycle cognitive skill** (PR #43306) – If merged, this would add a permanent persona to every Hermes installation; may face scrutiny over design decisions.

**Predictions for next release (v0.9.x or v1.x):**  
- Lazy tool loading is high priority but seems complex – maybe not in the very next release.  
- Delegate per-task routing (#36790) is close to merge (open since June 1).  
- Docker entrypoint fix (#43314) and desktop sidebar regression (#43310) are likely to be merged quickly.  
- The `write_mode` gate (#38199) is already merged; backport to stable may appear.  
- Chinese CLI support (#43307) may be deferred or contributed externally.

## 7. User Feedback Summary

**Pain points:**
- **Token overhead**: The forced injection of all 50+ tool schemas is the most echoed concern, wasting 3,500–5,000 tokens per call regardless of need.  
- **Telegram UX**: Context compaction making messages disappear feels “terrible” and “extremely disruptive”. Users expect compacted messages to be hidden, not deleted.  
- **Sudo handling**: Users who write explicit `sudo` skills find models ignoring them and trying unconventional paths.  
- **Docker configuration**: Multiple reports of `docker_extra_args` not being honored (now fixed with today’s merged trio).  
- **Cross-platform quirks**: macOS Feishu auth conflict, Windows CA trust gaps, Termux doctor false positives indicate friction for non-Linux users.  
- **Visual regressions**: The sidebar zoom issue (#43309) and font clarity on high-DPI (#43248) degrade the desktop experience.

**Satisfaction signals:**
- Fast turnaround on `docker_extra_args` bugs (multiple PRs today) shows maintainers respond quickly to regression clusters.  
- The merge of VS Code theme installation and Russian localization demonstrates attentiveness to customisation and global audience needs.  
- The `write_mode` approval gate for memory/skills directly addresses user feedback about unwanted auto-saves.  
- The Kanban archived board guard (#43286) prevents a confusing UI resurrection.

**Overall**: The community is vocal about token efficiency and UX polish, but appreciates rapid bug squashing and feature delivery.

## 8. Backlog Watch

Issues and PRs needing maintainer attention (based on age, discussion volume, or lack of response):

| Item | Age | Issue | Notes |
|------|-----|-------|-------|
| **#6839** – Lazy tool schema loading | Created Apr 9 (2 months) | [Issue link](https://github.com/NousResearch/hermes-agent/issues/6839) | High community interest (23 comments, 13 👍). No maintainer comment since creation. Needs a decision: pursue as feature, or close as wontfix? |
| **#36790** – Per-task model routing for delegate | Created Jun 1 (9 days) | [PR link](https://github.com/NousResearch/hermes-agent/pull/36790) | Open, no merge status; touches multiple linked issues. May need rebase or reviewer bandwidth. |
| **#40416** – Telegram compaction UX | Created Jun 6 (4 days) | [Issue link](https://github.com/NousResearch/hermes-agent/issues/40416) | P2 bug, no fix PR yet. Should be triaged for priority. |
| **#43232** – Dashboard D state with NAS | Created today | [Issue link](https://github.com/NousResearch/hermes-agent/issues/43232) | P2, potentially critical for users with network-mounted home directories. No response yet. |
| **#43296** – Feishu WebSocket auth conflict macOS | Created today | [Issue link](https://github.com/NousResearch/hermes-agent/issues/43296) | P2, macOS-specific, no fix yet. |
| **#43278** – Firecrawl scrape timeout (PR) | Created today | [PR link](https://github.com/NousResearch/hermes-agent/pull/43278) | Tagged as duplicate – needs maintainer triage. |
| **#43286** – Kanban archived board guard | Already closed today | [PR link](https://github.com/NousResearch/hermes-agent/pull/43286) | Closed, but may benefit from backport to stable branch. |
| **#43306** – qca-cycle cognitive skill (PR) | Created today | [PR link](https://github.com/NousResearch/hermes-agent/pull/43306) | Large novel feature; requires scrutiny for design, security, and stability. |

**General observation**: Several P2 bugs reported today lack maintainer responses, suggesting the team may be overwhelmed by the high volume of incoming issues. The backlog on #6839 is the most visible gap, as it has been open for two months with no maintainer engagement despite 23 comments and 13 thumbs-up.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest — 2026-06-10

## 1. Today's Overview
PicoClaw sees an exceptionally active 24 hours with **21 issues updated** and **16 PRs updated**, signalling intense development and community engagement. A major coordinated security audit (14 issues filed by YLChen-007) dominates the open issue tracker, covering SSRF bypasses, authorization holes, and control‑plane vulnerabilities. Several fixes are already in flight via PRs like #3083 and #3085. One **nightly release** (v0.2.9-nightly) was cut automatically, and six PRs were merged/closed, including a large agent‑collaboration feature. The project remains in a healthy, fast-paced cycle.

## 2. Releases
- **Nightly build** `v0.2.9-nightly.20260610.b9a8fad6`  
  Automated, may be unstable. No breaking changes or migration notes provided.  
  Full changelog: [v0.2.9...main](https://github.com/sipeed/picoclaw/compare/v0.2.9...main)

## 3. Project Progress (Merged/Closed PRs Today)
- **#3064** (`fix(config): add ok check for type assertion in migration model name indexing`) — prevents a panic from malformed config entries. [PR link](https://github.com/sipeed/picoclaw/pull/3064)
- **#3086** (`docs: update wechat qrcode`) — documentation update. [PR link](https://github.com/sipeed/picoclaw/pull/3086)
- **#2942** (`fix(config): use canonical hyphenated model ID for default claude-sonnet entry`) — fixes out‑of‑box failure on Anthropic. [PR link](https://github.com/sipeed/picoclaw/pull/2942)
- **#2940** (`fix(providers): omit temperature for claude-opus-4-7`) — resolves HTTP 400 error for Opus 4-7. [PR link](https://github.com/sipeed/picoclaw/pull/2940)
- **#2937** (`Feat/agent collaboration`) — introduces a durable inter‑agent communication bus with mailboxes, collaboration threads, and permission‑aware delivery. [PR link](https://github.com/sipeed/picoclaw/pull/2937)

## 4. Community Hot Topics
- **[Issue #2404](https://github.com/sipeed/picoclaw/issues/2404) — Streaming HTTP request config** (11 comments, 1 👍)  
  Long‑running feature request: add `"streaming": true` to config for OpenAI‑compatible backends. Community discussion shows strong desire for real‑time output without polling.  
- **[Issue #2984](https://github.com/sipeed/picoclaw/issues/2984) — Explicit turn completion signal for WebSocket** (1 comment, 1 👍)  
  External clients need a deterministic "done" event. Drives toward a richer Pico protocol.  
- **[Issue #3088](https://github.com/sipeed/picoclaw/issues/3088) — Replace libolm with vodozemac** (new, 0 comments)  
  Security‑driven dependency swap; signals community prioritisation of modern crypto libraries.

**Underlying needs**: Real‑time streaming, standardised protocol events, and hardened dependency supply chain.

## 5. Bugs & Stability
A bulk of **14 security issues** were filed today, all by YLChen-007. Severity is critical in most cases:

| Issue | Severity | Summary | Fix in flight? |
|-------|----------|---------|----------------|
| [#3072](https://github.com/sipeed/picoclaw/issues/3072) | **Critical** | CSRF on first‑run password setup → full control‑plane takeover | No dedicated fix PR yet |
| [#3071](https://github.com/sipeed/picoclaw/issues/3071) | **Critical** | Authenticated WebSocket clients can trigger `/reload` without proper auth | Addressed by PR #3083 (in review) |
| [#3068](https://github.com/sipeed/picoclaw/issues/3068) | **Critical** | MQTT `allow_from` bypass via spoofed `client_id` | No dedicated fix yet |
| [#3078](https://github.com/sipeed/picoclaw/issues/3078) | **High** | `web_fetch` SSRF via environment proxy | No dedicated fix yet |
| [#3077](https://github.com/sipeed/picoclaw/issues/3077) | **High** | SSRF bypass via `198.18.0.0/15` | Fix PR #3085 (open, blocks the range) |
| [#3075](https://github.com/sipeed/picoclaw/issues/3075) | **High** | Untrusted `skills/` auto‑loaded into system prompt | No dedicated fix yet |
| [#3074](https://github.com/sipeed/picoclaw/issues/3074) | **High** | ISATAP IPv6 SSRF bypass | No dedicated fix yet |
| [#3073](https://github.com/sipeed/picoclaw/issues/3073) | **Medium** | Signed LINE webhook replay | No dedicated fix yet |
| [#3070](https://github.com/sipeed/picoclaw/issues/3070) | **High** | OneBot media URL arbitrary fetch | No dedicated fix yet |
| [#3069](https://github.com/sipeed/picoclaw/issues/3069) | **High** | `allowed_cidrs` bypass via reverse proxy | Addressed by PR #3083 |
| [#3076](https://github.com/sipeed/picoclaw/issues/3076) | **Medium** | WeCom group trigger policy bypass | No dedicated fix yet |
| [#3079](https://github.com/sipeed/picoclaw/issues/3079) | **High** | `exec` whitelist allows `jq` environment disclosure | No dedicated fix yet |
| [#3080](https://github.com/sipeed/picoclaw/issues/3080) | **High** | `allowed_cidrs` bypass via loopback proxying during setup | Addressed by PR #3083 |
| [#3081](https://github.com/sipeed/picoclaw/issues/3081) | **High** | Approval hook `cwd` symlink race for `exec` | No dedicated fix yet |

Additionally, a long‑standing Windows bug ([#2472](https://github.com/sipeed/picoclaw/issues/2472), `list_dir` path separator mismatch) was closed today, indicating the fix has been shipped.

## 6. Feature Requests & Roadmap Signals
- **[#2404](https://github.com/sipeed/picoclaw/issues/2404) — Streaming config** — already has discussion and a clear implementation path. Likely to land in next stable release.
- **[#2984](https://github.com/sipeed/picoclaw/issues/2984) — WebSocket turn completion signal** — needed for external client integration; roadmap signal for protocol maturation.
- **[#3088](https://github.com/sipeed/picoclaw/issues/3088) — Replace libolm with vodozemac** — security‑critical dependency upgrade; may be fast‑tracked.
- **[PR #3063](https://github.com/sipeed/picoclaw/pull/3063) — Delta Chat gateway** — new channel integration, under review.
- **[PR #2917](https://github.com/sipeed/picoclaw/pull/2917) — NEAR AI Cloud provider** — adds TEE‑capable LLM support; open since May 21, may need maintainer attention.

**Prediction for next version (v0.3.0?)**: streaming config, WebSocket protocol enhancements, removal of libolm, and at least one new provider (NEAR AI).

## 7. User Feedback Summary
- **Pain points**:  
  - Multi‑message history not showing all user messages (bug #2796, now fixed via PR #2990).  
  - Windows `list_dir` failure (#2472, now fixed).  
  - Anthropic `claude-opus-4-7` temperature deprecation (#2939, fixed by PR #2940).  
  - Empty LLM response not retried (PR #2983 addresses this).  
- **Use cases driving feature requests**: real‑time streaming, inter‑agent collaboration, multi‑protocol gateways (Delta Chat).  
- **Satisfaction mirrors responsiveness** – most reported bugs are fixed or have open fix PRs within days. The security audit being addressed by immediate PRs indicates strong maintainer attention.

## 8. Backlog Watch
- **[Issue #2404](https://github.com/sipeed/picoclaw/issues/2404) — Streaming config** — open since April 7, 2 months without merge. Discussion is mature; maintainer decision needed.
- **[PR #2917](https://github.com/sipeed/picoclaw/pull/2917) — NEAR AI Cloud provider** — stale since May 21, no updates from maintainers. Should be reviewed or merged.
- **[PR #2990](https://github.com/sipeed/picoclaw/pull/2990) — Fix session history display** — stale since June 2, even though the related bug #2796 was closed. Needs final review.
- **[PR #2988](https://github.com/sipeed/picoclaw/pull/2988) — Use `summarize_token_percent` config** — stale, addresses #2968. Should be merged to unblock user configuration.
- **[PR #2983](https://github.com/sipeed/picoclaw/pull/2983) — Retry empty LLM response** — stale, fixes an important silent failure. Maintainer attention needed.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest – 2026-06-10

## 1. Today’s Overview  
The project saw an extraordinary spike in activity with **43 pull requests updated** in the last 24 hours, **39 of which were merged or closed** — an unusually high throughput that suggests a major housekeeping or release preparation effort. Only **one issue** (#1690) was updated, remaining open with growing community interest. No new releases were published today, but the volume of merged PRs points to a large batch of feature work and documentation nearing completion. Overall, project health appears very strong, with maintainers actively processing contributions and resolving long-standing PRs.

## 2. Releases  
No new releases were published today. The last release (if any) is not listed, so no migration notes or breaking changes to report.

## 3. Project Progress – Merged/Closed PRs  
The following significant features and fixes were merged or closed today (list reflects top-20 by comment activity; 39 total closed):

- **Agent Trace Observability** – [#1202](https://github.com/nanocoai/nanoclaw/pull/1202) adds per-invocation traces and a web UI on port 3001 for debugging.  
- **Direct Runner Mode** – [#1285](https://github.com/nanocoai/nanoclaw/pull/1285) introduces `NANOCLAW_DIRECT_RUNNER=1` to run agents in-process, bypassing Docker containers.  
- **Skill Marketplace / Registry** – [#1309](https://github.com/nanocoai/nanoclaw/pull/1309) implements CLI commands for discovering and installing skills from GitHub-hosted repositories.  
- **Approval-Gated Capabilities** – [#1245](https://github.com/nanocoai/nanoclaw/pull/1245) adds `/approve` and `/reject` skills for mutation workflows.  
- **Plugin System** – [#1387](https://github.com/nanocoai/nanoclaw/pull/1387) introduces a plugin architecture analogous to the channel pattern.  
- **Multi-runtime Agent SDK Abstraction** – [#1690](https://github.com/nanocoai/nanoclaw/issues/1690) (issue only, not a PR) remains open; see Community Hot Topics.  
- **Prompt Trace Logging** – [#337](https://github.com/nanocoai/nanoclaw/pull/337) logs prompt/response flows to JSONL with redaction support.  
- **External Markdown Seed Files** – [#357](https://github.com/nanocoai/nanoclaw/pull/357) allows injecting `.md` files into persistent context.  
- **Setup-Dev Skill** – [#1161](https://github.com/nanocoai/nanoclaw/pull/1161) simplifies local development setup.  
- **Explicit Claude Model** – [#1192](https://github.com/nanocoai/nanoclaw/pull/1192) makes the model choice explicit in code.  
- **WebUI Control Panel** – [#212](https://github.com/nanocoai/nanoclaw/pull/212) adds a full web control panel (closed after long blocking).  
- **Container Sandbox Design Doc** – [#1084](https://github.com/nanocoai/nanoclaw/pull/1084) provides a comprehensive design reference.  
- **Feishu Zombie Cards Fix** – [#2718](https://github.com/nanocoai/nanoclaw/pull/2718) resolves a production bug with stuck “running” cards (see Bugs).  
- **Finance DD Agent Skill** – [#2723](https://github.com/nanocoai/nanoclaw/pull/2723) adds a new skill (closed).  
- Multiple documentation PRs improved JSDoc, security audit docs, and group-level CLAUDE.md examples.

## 4. Community Hot Topics  
The single active issue, **[#1690 – Multi-runtime agent SDK abstraction](https://github.com/nanocoai/nanoclaw/issues/1690)**, generated the most discussion (5 comments, 3 👍). The author proposes a `AgentRuntime` interface that allows plugging different SDKs (Claude, Codex, local models) as modular skills, mirroring the `/add-telegram` channel pattern. This underscores a growing desire among users for flexibility beyond the default Claude Agent SDK, especially for cost control and offline use. The high reaction count and recent update suggest this feature could become a roadmap priority.

All 43 PRs updated today had undefined comment counts, so no other threads drove active debate.

## 5. Bugs & Stability  
Two bug-fix PRs were active today, both of high severity:

- **[#2718 – fix(feishu): cleanup zombie active_cards](https://github.com/nanocoai/nanoclaw/pull/2718)** (merged/closed) – Resolves a production bug where interactive cards remained stuck in a “running” state for >50 minutes after the agent runner timed out. Root cause: the cleanup handler only fired on SDK’s `final` event, which never executed for killed processes. Fix adds a fallback cleanup. **Severity: High** – affected real users.

- **[#2722 – fix(telegram): use CSPRNG for pairing codes](https://github.com/nanocoai/nanoclaw/pull/2722)** (open) – Switches code generation from `Math.random` to `crypto.randomInt` to prevent predictable pairing codes, which could allow unauthorized chat registration and owner escalation. **Severity: High (security)** – fix pending review.

Additionally, **[#1333](https://github.com/nanocoai/nanoclaw/pull/1333)** added build-time version metadata to logs, improving debuggability (low-severity improvement, not a bug fix).

No crashes or regressions were reported today.

## 6. Feature Requests & Roadmap Signals  
The most prominent feature signal is **Issue #1690** (multi-runtime abstraction), which could unify support for Codex, local models, and future SDKs. Its modular “skill” approach is already partly realized by the merged **[#1387 Plugin System](https://github.com/nanocoai/nanoclaw/pull/1387)** and **[#1309 Skill Marketplace](https://github.com/nanocoai/nanoclaw/pull/1309)**.

Other patterns pointing to future releases:
- **Direct runner mode** ([#1285](https://github.com/nanocoai/nanoclaw/pull/1285)) – likely to be promoted to default or more documented.
- **Approval workflows** ([#1245](https://github.com/nanocoai/nanoclaw/pull/1245)) and **trace observability** ([#1202](https://github.com/nanocoai/nanoclaw/pull/1202)) – both nearly ready for inclusion in the next stable release.
- Several PRs (e.g., [#212 WebUI](https://github.com/nanocoai/nanoclaw/pull/212), [#337 prompt logging](https://github.com/nanocoai/nanoclaw/pull/337)) were closed after long blocking, possibly to consolidate before a major version bump.

Prediction: The next release (v1.x or v2.0) will likely include the skill marketplace, direct runner mode, approval-gated skills, and the observability dashboard.

## 7. User Feedback Summary  
- **Pain point: SDK lock-in** – Issue #1690 explicitly requests multi-runtime support, indicating dissatisfaction with being tied to a single agent SDK.  
- **Pain point: Stuck interactive cards** – PR #2718 was motivated by a real production outage; user reported cards hung for 50+ minutes.  
- **Pain point: Security** – PR #2722 was opened by a user concerned about predictable pairing codes, reflecting a need for stronger defaults.  
- **Satisfaction: Self-service skills** – The marketplace ([#1309](https://github.com/nanocoai/nanoclaw/pull/1309)) and approval skills ([#1245](https://github.com/nanocoai/nanoclaw/pull/1245)) indicate eager adoption of modular, community-driven extensibility.  
- **Positive contributor activity** – 39 merged/closed PRs in one day suggests a healthy, engaged contributor base with responsive maintainers.

## 8. Backlog Watch  
No critical issues or PRs are languishing without attention. The only open issue ([#1690](https://github.com/nanocoai/nanoclaw/issues/1690)) was updated today, so it is active. The two open PRs from today ([#2722](https://github.com/nanocoai/nanoclaw/pull/2722) and [#2721](https://github.com/nanocoai/nanoclaw/pull/2721)) are fresh and awaiting review. Older PRs (#212, #214, #337, etc.) that had been blocked for months were all closed today, indicating the maintainers are clearing backlog. No items currently appear at risk of neglect.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-06-10

## 1. Today’s Overview

The NullClaw project saw a healthy burst of activity over the past 24 hours, with **5 issues updated** (1 open, 4 closed) and **8 pull requests updated** (7 merged/closed, 1 open). Bug-fix velocity remains high: three critical issues—PII false positives, missing typing indicators, and the dormant `compact_context` flag—were resolved and merged. One significant new feature, an **Evolink provider integration**, landed, and the long-running **cross-memory feature** (#711) was finally closed. The sole open issue (agent-type cron jobs failing to spawn) is likely being addressed by an open PR (#948). Overall, the project demonstrates strong maintainer engagement and a steady cadence of incremental improvements.

## 2. Releases

**None.** No new versions were published in the last 24 hours.

## 3. Project Progress — Merged/Closed PRs (7 items)

All seven closed PRs from the past 24 hours represent meaningful fixes or additions:

- **[#945 — fix(redaction): reject ISO date/time patterns as false-positive phone matches](https://github.com/nullclaw/nullclaw/pull/945)**  
  Resolves #944. Adds a guard to the phone-number matcher to exclude date/time patterns (e.g., `2026-06-02 20:17`), preventing the PII redactor from mangling agent timestamp output.

- **[#946 — fix(agent): filter tools in system prompt text by tool_filter_groups](https://github.com/nullclaw/nullclaw/pull/946)**  
  Ensures that only built-in tools and MCP tools from `always` filter groups appear in the text-based system prompt. Dynamic-group MCP tool schemas are still sent via native API tool-calling, improving prompt clarity and reducing token waste.

- **[#947 — feat(providers): add Evolink as an OpenAI-compatible provider](https://github.com/nullclaw/nullclaw/pull/947)**  
  Adds Evolink (a multi-model gateway) as a first-class provider, slotting into the existing OpenAI-compatible handler. This expands the list of supported backends beyond hardcoded Claude models.

- **[#943 — fix(telegram): show typing indicator during callback-query processing](https://github.com/nullclaw/nullclaw/pull/943)**  
  Closes #942. Now presses of inline buttons (e.g., `nc_choices`) trigger a “typing…” indicator, eliminating silent pauses during model inference.

- **[#940 — fix(models): query base_url for custom OpenAI-compatible providers](https://github.com/nullclaw/nullclaw/pull/940)**  
  Closes #936. Instead of falling back to a hardcoded Claude model list, the `/models` menu now queries the provider’s actual `/v1/models` endpoint when a custom `base_url` is configured.

- **[#939 — fix(agent): honor compact_context flag instead of always compacting](https://github.com/nullclaw/nullclaw/pull/939)**  
  Closes #937. The `compact_context` field in `AgentConfig` is now read at runtime; previously the code was always compacting history regardless of the flag’s value.

- **[#711 — Feat/cross memory (closed after long review)](https://github.com/nullclaw/nullclaw/pull/711)**  
  Merges a deterministic memory event stream enabling memory synchronisation across agent instances. This provides primitives for cross-agent knowledge sharing (e.g., Agent B knowing preferences established in Agent A).

## 4. Community Hot Topics

Most issues and PRs generated **zero reactions** and only **1 comment each**, suggesting focused, low-noise collaboration. The following items stand out by nature of being open or involving multiple participants:

- **[Issue #941 — Agent-type cron jobs don't spawn a subprocess](https://github.com/nullclaw/nullclaw/issues/941)**  
  The **only open issue** updated in the last 24h. Reported by weissfl, it describes scheduled tasks with `job_type: "agent"` completing without ever starting the agent subprocess, so Telegram delivery never happens. This blocks a core use case (scheduled agent runs). One comment exists; no fix merged yet, but PR #948 may address it.

- **[PR #948 — fix cron agent delivery attribution (OPEN)](https://github.com/nullclaw/nullclaw/pull/948)**  
  The **only open PR**. Author DonPrus aims to pass cron delivery metadata into spawned `nullclaw agent` subprocesses and update the `once-agent` handler. This likely targets Issue #941, though not explicitly linked. It was created on the digest date and may merge soon.

- **[PR #711 — Feat/cross memory](https://github.com/nullclaw/nullclaw/pull/711)**  
  Though closed, this feature was heavily anticipated (created March 2026, merged now). It introduces cross-agent memory, a frequently requested capability for multi-agent workflows.

**Underlying needs:** Users want reliable scheduled agent execution, transparent cross-agent memory, and accurate provider model detection. The community is clearly invested in making NullClaw production-ready for autonomous, scheduled tasks.

## 5. Bugs & Stability

| Bug | Severity | Status | Fix PR |
|-----|----------|--------|--------|
| **Agent cron jobs don’t spawn subprocess** (#941) | **High** – scheduled tasks fail silently, core functionality broken | **Open** | PR #948 (open, may fix) |
| PII redactor false positives on date/time (#944) | Medium – agent command output corrupted | **Closed** | #945 (merged) |
| Missing typing indicator on inline buttons (#942) | Low – UX annoyance only | **Closed** | #943 (merged) |
| Custom OpenAI provider falls back to Claude (#936) | Medium – incorrect model listing | **Closed** | #940 (merged) |
| `compact_context` flag dead (#937) | Low – flag ignored, always compacting | **Closed** | #939 (merged) |

The most pressing stability risk is **#941**; no merged fix exists yet, though #948 is promising. All other bugs were resolved within 24 hours of their reports, indicating a responsive maintenance pipeline.

## 6. Feature Requests & Roadmap Signals

- **Provider expansion** (PR #947): Adding Evolink suggests ongoing work to support more OpenAI-compatible backends. Expect more gateway providers (e.g., OpenRouter, Together) in future releases.
- **Cross-memory synchronization** (PR #711, now merged): This long-requested feature paves the way for multi-agent workflows, shared context, and persistent memory across sessions.
- **Tool filtering in prompts** (PR #946): A refinement to keep system prompts lean; signals future work on dynamic tool selection and context optimization.
- **Cron job reliability** (Issue #941, PR #948): The community is pushing for robust scheduled agent execution. Likely the next hotfix release will address this.

Predictions for next version: a minor release (v0.x) containing the cron fix, the Evolink provider, and the cross-memory feature, plus the accumulated bug fixes.

## 7. User Feedback Summary

**Reporters:** weissfl (4 issues), vernonstinebaker (1 issue, 2 PRs), DonPrus (2 PRs), raskevichai (3 PRs), EvoLinkAI (1 PR).

**Pain points expressed:**
- Scheduled agent tasks are unreliable (“job marked as completed but agent subprocess never started” – weissfl).
- PII redactor overzealous on timestamps, breaking agent outputs (vernonstinebaker).
- Telegram UX gaps: no typing indicator on inline buttons (weissfl).
- Misleading model selection: custom providers display wrong models (weissfl).
- Configuration flags that have no effect (weissfl).

**Satisfaction indicators:** All reported bugs (except #941) were closed within 1–14 days with targeted fixes. The project’s rapid response time suggests high maintainer attentiveness and a quality-focused culture.

## 8. Backlog Watch

- **[Issue #941 — Agent-type cron jobs don't spawn a subprocess](https://github.com/nullclaw/nullclaw/issues/941)**  
  Created 2026-05-31, still open. No maintainer comment yet; the only activity is a linked PR (#948) opened today. If #948 merges, this issue may soon close. If not, it risks becoming a persistent blocker for users relying on scheduled agents.

- **No other long-unanswered issues** of note. The oldest issue updated in the last 24h was #711 (PR, now closed). The project backlog appears well-managed with no stale critical items.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

## IronClaw Project Digest – 2026-06-10

### 1. Today's Overview
IronClaw saw extremely high activity, with **50 pull requests** updated in the last 24 hours (46 open, 4 closed/merged) and 7 issues updated (6 open, 1 closed). Development is concentrated on the Reborn v2 architecture: attachment handling, project-scoped ownership, Slack personal DM targets, persistent approval policies, and i18n completion. No new releases were published today. The project remains in a feature-heavy phase with steady CI and stability work, though a nightly E2E failure (open since May 27) and a reported provider-configuration save bug are notable stability concerns.

### 2. Releases
_No new releases were created in the last 24 hours._ (Previous release work is tracked in PR #3708, still open.)

### 3. Project Progress
Four pull requests were merged or closed today:

- **[#4600] – Add Slack personal DM outbound targets** (core, XL)  
  Merged. Implements Phase 4 C2 of the trigger delivery default outbound plan by adding provider-side Slack personal DM target authority alongside existing shared Slack channel targets. Keeps the backend contract channel-neutral.  
  [GitHub](https://github.com/nearai/ironclaw/pull/4600)

- **[#4553] – Complete webui v2 i18n tokens** (core, XL)  
  Merged. Removes remaining hardcoded English strings from extensions, jobs, routines, and sidebar components; bridges translation gaps across locales.  
  [GitHub](https://github.com/nearai/ironclaw/pull/4553)

- **[#4554] – Incomplete i18n coverage and translator runtime crashes** (issue)  
  Closed alongside the PR above.  
  [GitHub](https://github.com/nearai/ironclaw/issues/4554)

In addition, the following PRs advanced key features (still open as of digest time):

- Attachment pipeline Tracks 2, 6, and extraction crate (#4655, #4668, #4675, #4676, #4677) – building byte-storage and model-visible context for attachments.
- Project-scoped ownership model and vocabulary (#4662, #4663, #4664) – design docs and core implementation.
- Persistent approval policies (#4613) – stores and wires `AlwaysAllow` through Reborn.
- Trajectory observer + LLM injection (#4588) – observability seams for external hosts.

### 4. Community Hot Topics
While reaction counts are low across all items today, the most discussed pieces (by size, scope, or author attention) are:

- **#4674 – Gmail tool: no archive / label-modify action** (issue, 1 comment)  
  A clear user-requested extension to the Gmail tool – only trash is available; archive and label modification are missing.  
  [GitHub](https://github.com/nearai/ironclaw/issues/4674)

- **#4667 – Support Ask-gated capability approvals in Reborn REPL** (issue)  
  Raises the need to surface host approval requests for `PermissionMode::Ask` capabilities – a critical UX gap for the Reborn CLI.  
  [GitHub](https://github.com/nearai/ironclaw/issues/4667)

- **#4680 – Fix OpenAI-compat non-text content emission** (open PR)  
  Addresses the `[non_text_content]` canary that opaque-ed non-text parts to the model – a correctness issue affecting multimodal workflows.  
  [GitHub](https://github.com/nearai/ironclaw/pull/4680)

The underlying need across these is **multimodal and multi-tool completeness** – users want to handle images, audio, and Gmail labels, and the Reborn CLI must properly gate and render capability requests.

### 5. Bugs & Stability

| Severity | Issue / PR | Description | Status |
|----------|-------------|-------------|--------|
| **High** | **#4673** – NEAR AI provider config cannot be saved after successful test connection | Silent failure on Save; provider appears unconfigured. Blocks onboarding for NEAR AI users on Reborn. No fix PR yet. | Open |
| **Medium** | **#4108** – Nightly E2E failed | CI scheduled run failed. Open since May 27. No comments or linked fix. | Open |
| **Medium** | **#4678** – Scrub provider reasoning metadata (PR) | Addresses leak of secret-like tokens in provider reasoning strings. Prevents sensitive data from reaching model or logs. | Open, risk low |
| **Low** | **#4680** – Stop emitting `[non_text_content]` for non-text parts (PR) | Fixes a canary that turned images/audio into opaque text. Non-critical but important for multimodal correctness. | Open, risk low |
| **Info** | **#4666 / #4665** – File size cap tracking for `slack_host_state.rs` (2,823 lines) and `slack_host_beta.rs` (3,359 lines) | Architecture rule violations; no regression but technical debt. PRs must leave these files shorter. | Open, tracking |

No crash or regression with a direct fix PR exists for the highest-severity bug (#4673) yet.

### 6. Feature Requests & Roadmap Signals
Strong signals from issues and PRs indicate the upcoming roadmap focuses on:

- **Attachment pipeline** (Tracks 1-6) – extractors, storage, model-visible context, landing crate. Likely part of next minor release (#4644 epic).
- **Project-scoped ownership & automation** – design and core model nearing completion; surface-level API renaming in #4664. Foundation for multi-tenant or project-based workflows.
- **Persistent approval policies** – `AlwaysAllow` and revocable scoped storage (#4613). Expected to ship alongside Reborn CLI improvements (#4667).
- **Ask-gated capability approvals in REPL** – directly requested by #4667. Likely to follow the persistent approval policy work.
- **Observability seams** – trajectory observer and LLM provider injection (#4588) point toward external benchmarking (nearai-bench) and monitoring.
- **Gmail archive/label-modify** (#4674) – low implementation complexity; could appear in a subsequent release.

### 7. User Feedback Summary
Real user pain points surfaced today:

- **Onboarding friction**: NEAR AI provider configuration cannot be saved (#4673) – blocks first-time setup, likely causing frustration.
- **Missing Gmail functionality**: Only trash, no archive or label management (#4674) – users managing email expect parity with standard Gmail tools.
- **CLI transparency**: When a model calls an Ask-gated capability, the REPL doesn’t show the approval request (#4667) – leads to confusing silent failures.
- **Stability concerns**: Nightly E2E failures (#4108) and the ongoing release PR (#3708) suggest CI/release pipeline needs attention.

Overall satisfaction is moderate: core feature development (attachments, project ownership) is moving fast, but missing day-to-day tooling (Gmail, provider save) and CI instability dampen the experience.

### 8. Backlog Watch
The following important items have gone unanswered for notable periods:

- **[#4108] – Nightly E2E failed**  
  Opened May 27, last updated June 10 (auto-update from CI). No maintainer comment or fix PR. A recurring failure that may indicate a deeper stability issue.  
  [GitHub](https://github.com/nearai/ironclaw/issues/4108)

- **[#3708] – chore: release**  
  PR opened May 16, still open. A release automation PR that bumps crates (`ironclaw_common` 0.4.2→0.5.0, breaking). Remains unmerged, blocking new version publication.  
  [GitHub](https://github.com/nearai/ironclaw/pull/3708)

- **[#4501] – CI: avoid runtime tests for dependabot config updates**  
  Opened June 5, no merge or further discussion. Small CI hygiene fix.  
  [GitHub](https://github.com/nearai/ironclaw/pull/4501)

These items should be reviewed by maintainers to prevent accumulation of technical debt and release blockers.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

Here is the LobsterAI project digest for 2026-06-10, generated from the provided GitHub data.

---

### LobsterAI Project Digest — 2026-06-10

#### 1. Today's Overview
The project exhibited moderate activity over the last 24 hours, focused on consolidation and bug fixing rather than new feature launches. Four pull requests (PRs) were merged, addressing critical bug fixes in the renderer, a new data backup and migration feature, and improvements to system notifications. The only new issue (#2132) is a high-value user report concerning cross-model sub-task orchestration, which remains open for maintainer discussion. No new releases were published, indicating the team is likely stabilizing code for a future version.

#### 2. Releases
**None.** No new official releases were published on this date.

#### 3. Project Progress
Four PRs were merged/closed today, advancing the codebase in three key areas:

- **Bug Fixes & UX:** [#2133](https://github.com/netease-youdao/LobsterAI/pull/2133) (by fisherdaddy) fixed export and code copy bugs in the renderer and cowork areas.
- **Data Management:** [#2136](https://github.com/netease-youdao/LobsterAI/pull/2136) (by fisherdaddy) added a full data backup and migration feature spanning the renderer, docs, and main areas. Note: A complementary chore PR ([#2135](https://github.com/netease-youdao/LobsterAI/pull/2135)) temporarily disabled this backup feature shortly after its introduction, likely for testing or hotfix purposes.
- **System Integration:** [#2134](https://github.com/netease-youdao/LobsterAI/pull/2134) (by liuzhq1986) improved the application lifecycle by allowing it to restore from a task completion notification even when the main window was closed. It also ensured the notification handler is ready before opening a target session and kept active references for macOS Notification Center click actions.

#### 4. Community Hot Topics
The sole active issue represents the most significant community discussion point:

- **[#2132: 跨模型子任务调用的问题](https://github.com/netease-youdao/LobsterAI/issues/2132)** (by woxinsj)
  - *Status:* Open, 0 comments.
  - *Analysis:* This detailed report highlights a core user pain point: the lack of robust **cross-model sub-task orchestration**. The author describes a scenario where a planning-oriented main model (M3) delegates to an execution-oriented model (deepseek). The bug report's technical analysis reveals a `call_function` that is a gateway-level function call, not a proper `sessions_spawn` sub-task, leading to state tracking failures. The user’s underlying need is for a transparent, unified task monitoring system, regardless of which backend model is handling the sub-task.

#### 5. Bugs & Stability
- **Severity: Medium** — **Renderer Bug (Fixed):** PR [#2133](https://github.com/netease-youdao/LobsterAI/pull/2133) fixed export and code copy bugs in the renderer area. No open issues suggest residual problems.
- **Severity: Low** — **Task State Tracking Bug (Open):** Issue [#2132](https://github.com/netease-youdao/LobsterAI/issues/2132) reports a bug where cross-model sub-tasks fail to register correctly in the system’s session/sub-agent lists. While not a crash, this is a functional bug that breaks task monitoring for complex, multi-model workflows. The report includes the root cause but lacks a fix PR.

#### 6. Feature Requests & Roadmap Signals
- **Cross-Model Task Orchestration (High Probability):** The detailed request and technical diagnosis in issue [#2132](https://github.com/netease-youdao/LobsterAI/issues/2132) strongly signals a need for a formal cross-model sub-task protocol. Given the author’s specific improvement suggestions (e.g., using notification patterns for cross-model tasks, similar to existing same-model callbacks), this feature is likely to be prioritized for the next development cycle.
- **Data Backup & Migration (Integrated):** The feature requested by the community for data safety was already implemented in PR [#2136](https://github.com/netease-youdao/LobsterAI/pull/2136), though it is temporarily disabled. This indicates the team is validating it for a stable rollout in the next release.

#### 7. User Feedback Summary
- **Pain Point:** Users are dissatisfied with the current black-box nature of task execution when using different AI models for sub-tasks. There is a clear desire for visibility and control, where a controlling "main task" can track execution status and receive notifications from any "sub-task," regardless of the underlying AI model.
- **Use Case:** A power user is attempting to create a pipeline with a strong planner (M3) and a fast executor (deepseek). The current architecture fails to manage this hybrid workflow natively.
- **Satisfaction:** Low for this specific orchestration scenario. Other users likely benefit from the merged bug fixes and notification improvements.

#### 8. Backlog Watch
No long-standing or high-value unanswered issues or PRs are present in this data snapshot. The project maintainers (fisherdaddy, liuzhq1986) appear responsive, as demonstrated by the four rapidly merged PRs today. The only item to watch is **[Issue #2132](https://github.com/netease-youdao/LobsterAI/issues/2132)**, which, if not addressed soon by a developer, could become a blocker for users attempting advanced multi-agent workflows.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest – 2026-06-10

## Today's Overview
The Moltis project saw minimal activity over the last 24 hours. Only one issue was updated, and no pull requests or releases were published. The single open issue is a bug report regarding a missing Coqui provider configuration. With zero merged PRs and no new releases, the project appears to be in a quiet phase, possibly awaiting maintainer intervention or community contributions. Overall project health is stable but shows low development momentum on this particular day.

## Releases
*No new releases were published today.*

## Project Progress
- **Merged/Closed PRs:** 0
- **Features advanced or fixed:** None.

## Community Hot Topics
- **Issue #1114 – [Bug]: provider 'coqui' not configured** [🔗](https://github.com/moltis-org/moltis/issues/1114)  
  *Author: vvuk* | *Comments: 0* | *Reactions: 0*  
  This is the only active discussion today. While it has no comments or reactions yet, the bug pertains to a missing provider configuration—a common integration pain point. The issue’s presence signals that users are relying on the Coqui text-to-speech provider and encountering friction when setting it up. Maintainers may need to clarify configuration requirements or add documentation.

## Bugs & Stability
- **#1114 – provider 'coqui' not configured** *(minor)*  
  The only bug reported today is a configuration error for the Coqui provider. Severity is minor; it does not crash the application but blocks usage of that specific provider. No associated fix PR exists yet. Users can work around the issue by manually ensuring the provider is set up according to Moltis documentation.

## Feature Requests & Roadmap Signals
No feature requests were submitted today. However, the Coqui provider configuration issue (#1114) could hint at a desire for automatic provider detection or clearer setup guidance. Future releases might include improved environment variable checks or a configuration wizard to avoid similar user errors.

## User Feedback Summary
- **Pain Points:** One user (vvuk) encountered a configuration hurdle when trying to use the Coqui provider. The lack of comments suggests either the issue is straightforward or users have not yet engaged with it.
- **Satisfaction / Dissatisfaction:** With no closed issues or PRs today, it’s difficult to gauge overall user sentiment. The single open bug indicates a minor dissatisfaction with provider setup, but no widespread complaints.

## Backlog Watch
- **Issue #1114** (open since 2026-06-10) – While very recent, this bug may require maintainer confirmation or a quick documentation update. If left unaddressed, it could accumulate confusion for other Coqui users.  
- No other long-standing issues or stale PRs were flagged in today’s data. Maintainers should monitor #1114 for follow-up comments.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest — 2026-06-10

## Today's Overview

CoPaw (QwenPaw) is in a period of **high activity and rapid stabilization**, with 25 issues and 35 PRs updated in the last 24 hours and a new beta release published. The closure rate is strong: 12 of 25 issues and 16 of 35 PRs were closed or merged, indicating an active, responsive maintainer team. Engagement leans heavily toward **bug fixing and quality-of-life improvements**, with several first-time contributors landing fixes. The community is engaged on feature discussions (especially around agent learning loops and context compression), while backend stability and desktop UX (Tauri startup, localStorage persistence) remain focus areas. Overall project health appears **positive and developer-community friendly** with an accelerating pace of small, targeted fixes.

## Releases

- **v1.1.11-beta.2** was released today. According to changelog entries (partially visible):  
  - **feat(browser):** Added page-coordinate click support to `browser_control`.  
  - **fix(browser):** Added CDP timeout parameter and browser profile isolation for cross-browser switching.  

  No breaking changes or migration notes were indicated. The release appears to be a minor feature/bugfix iteration focused on browser automation reliability.  
  **Link:** [1.1.11-beta.2](https://github.com/agentscope-ai/QwenPaw/releases/tag/1.1.11-beta.2) (URL inferred from pattern; not fully visible in data)

## Project Progress

**Merged/Closed PRs today (16):**

- **MCP subprocess accumulation fix** — PR #5014 (rayrayraykk): Prevents MCP server processes from accumulating across restarts (linked to Issue #4834). Now merged.
- **Security: file preview path restriction** — PR #4981 (zhijianma): Adds `WORKING_DIR` boundary check and sensitive-path deny list to the `/files/preview` endpoint. Merged.
- **DingTalk AI Card empty output fix** — PR #5061 (celestialhorse51D): Defers card creation to prevent sending "Processing..." cards when agent output is empty. Merged.
- **Plugin infrastructure: CloudPaw / A2A** — PR #5033 (Saint-Yin): Refactors A2A router with rename/import endpoints, alias derivation from Agent Card, and plugin router registration. Merged.
- **Skill tag batch download** — PR #4969 (Leirunlin): Adds tag filtering to skill batch download (fixes Feature Request #2961). Merged.
- **E2E test fix for token usage empty state** — PR #5062 (yutai78786): Handles text-based empty state in token usage test. Merged.
- Various **first-time contributor fixes** for edge cases in cron jobs, context manager, Linux browser detection, backup, and llama.cpp version parsing (PRs #5040, #5038, #5037, #5041, #5035) — all merged today.

**Non-merged but Under Review/Active PRs (19 open):**
- #4973 (EaveLuo): Large unit test addition (129 test cases, Phase 5 modules).
- #4669 (jinglinpeng): Tauri auto-updater — slow-moving but high-impact desktop feature.
- #4891 (Leirunlin): Multi-path support for skill pool — enhancement still in review.
- #4975 (saltapp): Customizable column order in sessions page.
- #5051 (wangfei010313): Persist Tauri backend port across restarts to prevent localStorage loss.

## Community Hot Topics

1. **[#5017 — Enhancement: 建议关注 Hermes Agent 的"学习循环"特性]** (CLOSED, 10 comments, 3 👍)  
   - **Context:** A feature suggestion urging the project to study Hermes Agent's "Learning Loop" capability, where agents automatically create and iterate skills from their own behavior.  
   - **Analysis:** This is a strategic discussion about the project's feature roadmap. The community is signaling that **self-improving agent behavior** is the next frontier and expects QwenPaw to compete with emerging disruptive projects (Hermes 46k+ stars).  
   - **Link:** [Issue #5017](https://github.com/agentscope-ai/QwenPaw/issues/5017)

2. **[#4727 — Breaking Change: Migrate backend from AgentScope 1.x to 2.0]** (OPEN, 8 comments, 2 👍)  
   - **Context:** Plan to upgrade backend dependency from AgentScope 1.0.20 to AgentScope 2.0, adopting new architecture, APIs, and runtime model.  
   - **Analysis:** This is a major architectural change that affects every underlying component. Community is watching closely — migration is likely to ship in next 1-2 versions.  
   - **Link:** [Issue #4727](https://github.com/agentscope-ai/QwenPaw/issues/4727)

3. **[#4878 — WeChat delivery failure in cron jobs]** (CLOSED, 6 comments, 0 👍)  
   - **Context:** A user reported that `home_agent` cron jobs configured with WeChat delivery fail silently due to routing logic returning `session_id` instead of `user_id` in `to_handle_from_target`.  
   - **Analysis:** This was a moderately high-impact bug that prevented an entire channel from functioning with scheduled tasks. It has been closed (and fix PR #5060 merged).  
   - **Link:** [Issue #4878](https://github.com/agentscope-ai/QwenPaw/issues/4878)

## Bugs & Stability

**High Severity:**
- **[#5064 — Agent-generated cron tasks fail to trigger and cannot be edited]** (OPEN)  
  A user reports that scheduled tasks created by the agent appear in the UI but never trigger, and cannot be manually edited. This is a **critical usability bug** in the cron automation feature, potentially blocking users from relying on agent-initiated scheduling. No fix PR yet.  
  **Link:** [Issue #5064](https://github.com/agentscope-ai/QwenPaw/issues/5064)

- **[#5052 — Tool calls fail after several rounds with "got an unexpected keyword argument 'arguments'"]]** (OPEN)  
  All tool calls fail after a few rounds of a session, even though the OpenAI-compatible model returns standard tool-call format. This suggests a **model/server-side deserialization bug** in the runtime. No fix PR linked.  
  **Link:** [Issue #5052](https://github.com/agentscope-ai/QwenPaw/issues/5052)

**Medium Severity:**
- **[#5047 — Windows Tauri desktop app starts extremely slowly (10+ minutes)]** (OPEN)  
  Since packaging switched from Python to Tauri, startup time increased from 1-2 minutes to 10+ minutes, often with unresponsiveness. This is a **major regression** for desktop users. No fix PR yet.  
  **Link:** [Issue #5047](https://github.com/agentscope-ai/QwenPaw/issues/5047)

- **[#5031 — Skill slash invocation displays expanded SKILL.md content in Console]** (OPEN)  
  When using `/skill` commands, the raw `SKILL.md` file contents appear in the UI instead of being rendered properly. A console frontend rendering bug.  
  **Link:** [Issue #5031](https://github.com/agentscope-ai/QwenPaw/issues/5031)

**Low/Cosmetic:**
- [#5001 — 9router model connection failure] (OPEN) — Unsupported provider, likely configuration issue.
- [#4991 — Unanswered question about UI layout] (OPEN) — Minimal information.

**Bugs with Fix PRs Already Submitted:**
- [#5060 → PR #5061 (merged) — WeChat delivery session_id/user_id routing fix]
- [#4834 → PR #5014 (merged) — MCP process accumulation]
- [#4917 (slow chat switching) → no fix yet; #5051 (port persistence) → PR #5051 open]

## Feature Requests & Roadmap Signals

1. **[#5063 — Integrate Headroom context compression layer]** (OPEN, today)  
   A feature request to integrate Headroom as an optional proxy that compresses tool outputs, conversation history, and RAG chunks by 60-95% before sending to the LLM. This aligns with **user demand for token efficiency without quality loss**. Likelihood: **Medium-High** — the project has shown openness to context optimization features.  
   **Link:** [Issue #5063](https://github.com/agentscope-ai/QwenPaw/issues/5063)

2. **[#3751 — Windows system tray icon]** (OPEN, 3 comments)  
   First raised April 23, still open. Users want minimize-to-tray and quick-access menus. Given the Tauri desktop migration, this is a natural roadmap item for the desktop team.  
   **Link:** [Issue #3751](https://github.com/agentscope-ai/QwenPaw/issues/3751)

3. **[#4057 — AgentScope tracing initialization support]** (OPEN)  
   User request to support `agentscope.init(tracing_url=...)` within QwenPaw startup for integrating with Arize-Phoenix monitoring. This is blocked by the AgentScope 2.0 migration (#4727). Likely to be resolved when that migration ships.  
   **Link:** [Issue #4057](https://github.com/agentscope-ai/QwenPaw/issues/4057)

4. **[#2961 — Skill classification/folders]** (already partially addressed via PR #4969 — batch download with tags). The full folder-based skill pool organization is not yet implemented but is in design via PR #4891 (multi-path support, still open).  
   **Link:** [Issue #2961](https://github.com/agentscope-ai/QwenPaw/issues/2961)

## User Feedback Summary

**Pain Points (Real Reported):**
- **Desktop startup performance regression** — Tauri migration drastically slowed Windows startup (10+ min) — users are dissatisfied and consider this a blocker (#5047).
- **Cron job dysfunction** — agent-created tasks are invisible to the scheduling system — this undermines one of QwenPaw's core automation use cases (#5064).
- **Tool call decay over long sessions** — after a few rounds, tool invocation breaks entirely with a parser error — highly disruptive for power users (#5052).
- **Large chat history forces re-render every time** — users switching between tabs experience "spinning wheel" and long freezes (#4917 still open).
- **Skill slash command displays raw Markdown** — UI polish issue for skill users (#5031).

**Satisfaction Signals:**
- Users praise the **localized / out-of-box Chinese experience** ("开箱即用") and call the project "非常棒" (great) in multiple issues.
- The community is **actively contributing**: 5+ first-time contributors landed PRs today.
- The **fast closure rate** (12 issues, 16 PRs closed in 24h) shows maintainer responsiveness is appreciated.

## Backlog Watch

**Issues Needing Maintainer Attention:**

1. **[#3751 — Windows system tray icon]** — Open since April 23, 2.5+ months with no maintainer response. Low technical complexity; would significantly improve desktop UX at low cost.  
   **Link:** [Issue #3751](https://github.com/agentscope-ai/QwenPaw/issues/3751)

2. **[#4057 — AgentScope tracing init]** — Open since May 6 with 4 comments but no maintainer update. A straightforward configuration feature that could attract observability-minded users.  
   **Link:** [Issue #4057](https://github.com/agentscope-ai/QwenPaw/issues/4057)

3. **[#4669 — Tauri auto-updater]** — PR open since May 25; no recent activity. This is a critical desktop feature for sustainability. If merged, it would eliminate the need for manual downloads.  
   **Link:** [PR #4669](https://github.com/agentscope-ai/QwenPaw/pull/4669)

4. **[#4917 — Chat tab switching performance]** — Open since June 2 with 3 comments. Relates to core UX on large conversation histories. No fix PR or maintainer response seen.  
   **Link:** [Issue #4917](https://github.com/agentscope-ai/QwenPaw/issues/4917)

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-06-10

## 1. Today's Overview
The ZeroClaw project remains in a high-velocity development cycle with **24 issues** and **50 pull requests** updated in the last 24 hours. Activity is balanced across bug fixes, feature PRs, and documentation changes. Only **1 issue** was closed (#4710 – logo finalized) and **3 PRs** were merged/closed (including the AMQP SOP channel and zerocode UI polish). No new releases were cut. The project continues to address critical architectural debt (provider unification, security hardening, multi-tenant RBAC) while triaging real-world deployment bugs from the Telegram, MCP, and custom-provider front.

---

## 2. Releases
No new releases in this period.

---

## 3. Project Progress (Merged/Closed Items)
Three PRs and one issue were closed today:

- **Issue #4710** – [Feature]: A better LOGO of ZeroClaw (closed, 20 comments) — the community design initiative completed.
- **PR #7362** – `feat(zerocode): hold back context-window usage bar until it bakes` — merged. Gates a partially implemented UX feature behind a compile-time flag.
- **PR #7369** – `feat(channels/sop): AMQP inbound channel with mutual TLS, and drive deterministic SOP runs to completion` — merged. Enables a complete end-to-end SOP pipeline via AMQP, including deterministic run-to-completion semantics.

These represent progress on OS-level integrations and TUI polish, respectively.

---

## 4. Community Hot Topics
The most active discussions (by comment count and reactions) highlight core architecture and security concerns:

- **#4710** (CLOSED, 20 comments, 2 👍) – Logo design: community engagement on branding, now resolved.
- **#5937** (OPEN, 10 comments) – *“refactor: Unify providers architecture and reqwest client management”* – High-risk, accepted feature affecting every model provider. The community is pushing for consistent parameter handling and reduced code duplication.
- **#5982** (OPEN, 9 comments) – *“Per-sender RBAC for multi-tenant agent deployments”* – A widely requested feature for isolating workspaces, tools, and system prompts per user class.
- **#6378** (OPEN, 7 comments) – *“Discord Bot respond only in specific Discord channels”* – Request for `allowed_channels` config matching Matrix/Nextcloud Talk patterns.
- **#5844** (OPEN, 6 comments) – *“Too much emphasis on memory”* – Real operational pain where system prompt weighting of memory over current context degrades cron-job behavior.
- **#6721** (OPEN, 4 comments) – *“tool_search not in default_auto_approve → deferred_loading+webhook silently hangs 120s then auto-denies”* – Critical MCP workflow blocker.

[Issue #4710](https://zeroclaw-labs/zeroclaw/issues/4710) | [#5937](https://zeroclaw-labs/zeroclaw/issues/5937) | [#5982](https://zeroclaw-labs/zeroclaw/issues/5982) | [#6378](https://zeroclaw-labs/zeroclaw/issues/6378) | [#5844](https://zeroclaw-labs/zeroclaw/issues/5844) | [#6721](https://zeroclaw-labs/zeroclaw/issues/6721)

---

## 5. Bugs & Stability
Several new and prior bugs were updated today, many with associated fix PRs. Severity ranking follows:

### Critical / S1 (workflow blocked)
| Issue | Summary | Fix PR? |
|-------|---------|---------|
| [#6721](https://zeroclaw-labs/zeroclaw/issues/6721) | `tool_search` not auto-approved → MCP tool loading hangs 120s then denies (webhook mode) | No PR yet |
| [#5808](https://zeroclaw-labs/zeroclaw/issues/5808) | Default 32k context budget exceeded 3.3× on iteration 1 → perpetual trim | None |
| [#6646](https://zeroclaw-labs/zeroclaw/issues/6646) | `web_search_tool` / `web_fetch` not firing via Telegram in v0.7.5 | None |
| [#7439](https://zeroclaw-labs/zeroclaw/issues/7439) | Custom provider causes Doctor error `model_provider "custom.<alias>" is invalid` | **PR #7441** (open) |
| [#6687](https://zeroclaw-labs/zeroclaw/issues/6687) | Two independent SopEngine instances per daemon – MQTT runs invisible to `sop_status` | None |

### High / S2 (degraded behavior)
| Issue | Summary | Fix PR? |
|-------|---------|---------|
| [#5844](https://zeroclaw-labs/zeroclaw/issues/5844) | Memory overemphasis in system prompt (especially cron) | None |
| [#6876](https://zeroclaw-labs/zeroclaw/issues/6876) | `risk_profile.allowed_tools` does not restrict MCP tools | None |
| [#7376](https://zeroclaw-labs/zeroclaw/issues/7376) | zerocode Dashboard hides unavailable/error states, labels history as active | **PR #7444** (open) |
| [#7377](https://zeroclaw-labs/zeroclaw/issues/7377) | Dark themes inherit unreadable terminal foreground text | None |
| [#7378](https://zeroclaw-labs/zeroclaw/issues/7378) | macOS Cmd-C copy triggers quit chord | None |

### Medium / S3 (minor)
| Issue | Summary |
|-------|---------|
| [#7400](https://zeroclaw-labs/zeroclaw/issues/7400) | Locale selection appears to do nothing until restart and download |
| [#6002](https://zeroclaw-labs/zeroclaw/issues/6002) | Not clearly addressed to the assistant (Telegram/container scenario) |

Other open bug PRs targeting related issues: **#7446** (fix `image_info` for vision models), **#7348** (cron skip overdue jobs), **#7426** (warn on unimplemented `rerank_enabled`), **#7427** (warn on unimplemented `context_aware_tools`).

---

## 6. Feature Requests & Roadmap Signals
The following accepted/enhancement features with high risk are strong candidates for the next release:

- **#5937** – Unify providers architecture (reqwest client management, consistent params)
- **#5982** – Per-sender RBAC (multi-tenant isolation)
- **#6378** – Discord `allowed_channels` configuration
- **#5775** – Per-skill security permissions (scoped `allow_scripts` / `allowed_commands`)
- **#6916** – Process-memory limits on shell/skill_tool subprocess execution
- **#6917** – Honor action-scope filter in Composio tool dispatch
- **#7248** – Persist cached input tokens and include in cost accounting
- **#7410** – Read webhook signing secrets at handler time instead of caching at startup

PR activity suggests the **provider refactor (#5937)** and **process-memory limits (#6916)** are near completion. The **RBAC (#5982)** and **Discord channel filtering (#6378)** are still in early discussion but have strong community backing.

---

## 7. User Feedback Summary
Real-world user reports cluster around three pain points:

1. **Model provider integration** – Custom providers (LM Studio, OpenAI-compatible) face validation errors (`#7439`) and missing features (`#6646` Telegram tools not firing, `#6002` unclear addressing).
2. **Memory vs. context weighting** – In cron and long-running scenarios, the agent over-prioritizes old memories over the immediate prompt (`#5844`). Suggests tuning or configurable weighting.
3. **MCP tool approval flow** – Non-interactive webhook deployments hit silent 120s timeouts because `tool_search` is not in the auto-approve list (`#6721`). This is a design gap for headless operation.

Positive signals: Users are actively testing custom providers (vLLM, LM Studio) and contributing detailed bug reports with reproduction steps. The zerocode TUI received multiple UX defect reports from macOS users (`#7376–#7400`), indicating growing adoption outside Linux.

---

## 8. Backlog Watch
Several important issues remain open for 1+ months without a fix PR or maintainer response:

| Issue | Age (created) | Last Update | Summary |
|-------|---------------|-------------|---------|
| [#4853](https://zeroclaw-labs/zeroclaw/issues/4853) | 2026-03-27 | 2026-06-09 | Installing skills from `.well-known` URI – accepted, no visible PR activity |
| [#5775](https://zeroclaw-labs/zeroclaw/issues/5775) | 2026-04-15 | 2026-06-09 | Per-skill security permissions – accepted but blocked (needs author action) |
| [#5842](https://zeroclaw-labs/zeroclaw/issues/5842) | 2026-04-17 | 2026-06-09 | Evaluate `extra_args` validation/allowlist for Codex CLI – accepted, no PR |
| [#5982](https://zeroclaw-labs/zeroclaw/issues/5982) | 2026-04-22 | 2026-06-09 | RBAC multi-tenant – accepted but still in discussion phase |
| [#6250](https://zeroclaw-labs/zeroclaw/issues/6250) | 2026-05-01 | 2026-06-09 | Extract `require_auth` to route-layer middleware – accepted, no-stale |

These items represent unaddressed architectural improvements that could cause friction as adoption scales. Maintainers should prioritize triage, especially for the `.well-known` skills standard and per-skill security, which align with external industry initiatives (Agent Skills group, Cloudflare integration).

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*