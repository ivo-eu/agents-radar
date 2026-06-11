# OpenClaw Ecosystem Digest 2026-06-11

> Issues: 320 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-11 03:32 UTC

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

# OpenClaw Project Digest — 2026-06-11

## 1. Today's Overview

Project activity remains very high, with 320 issues and 500 pull requests updated in the last 24 hours. A new **v2026.6.6-beta.1** release tightens security boundaries across transcripts, sandbox binds, MCP stdio, and more. Of the updated PRs, 104 were merged or closed, indicating steady forward momentum, but the open issue count (302) and unresolved high-severity bugs (e.g., silent subagent failures, Discord internal leaks) show that the community is facing meaningful stability challenges. The maintainer team has labelled many issues as `needs-product-decision` or `needs-security-review`, suggesting active triage but also some decision bottlenecks.

## 2. Releases

**v2026.6.6-beta.1** – OpenClaw 2026.6.6-beta.1  
[Release link](https://github.com/openclaw/openclaw/releases/tag/v2026.6.6-beta.1)

- **Highlights**: Security boundaries are substantially tighter across transcripts, sandbox binds, host environment inheritance, MCP stdio, Codex HTTP access, native search policy, elevated sender checks, deleted-agent ACP bypasses, loopback tools, Discord moderation, and Teams group access.
- **Breaking changes**: Not explicitly documented; beta tag suggests a hardening release with no major API shifts.
- **Migration notes**: Users should review their `tools.elevated`, `sandbox`, and MCP configurations for potential stricter defaults.

## 3. Project Progress

In the past 24 hours, 104 PRs were merged or closed. Notable featured PRs that advanced the codebase:

- **#90128** [CLOSED] – `fix(sessions): preserve user /model override across daily/idle session rollover` – Fixes a bug where a user’s `/model` override was silently dropped on session rollover.  
- **#84938** [CLOSED] – `fix: forward reasoning_content from OpenAI-compatible providers` – Adds support for MiMo v2.5 and others that return `reasoning_content`.  
- **#77359** [CLOSED] – `[Bug]: Slash commands not registered for non-default Discord accounts in multi-bot setup` – Resolved a missing registration problem in multi-account Discord configurations.  
- **#78381** [OPEN] – `feat(embedded-runner): expose prep stage timings` – Adds prep stage snapshots to `EmbeddedRunAttemptResult` and the `agent_end` hook, improving plugin observability.  
- **#92079** [OPEN] – `fix(memory): auto-fix providerKey mismatch from CLI index --force` – Addresses memory index corruption when CLI and runtime have different provider cache keys.

## 4. Community Hot Topics

The most active issues (by comment count and reactions) reveal deep concerns about reliability and transparency:

- **Issue #44925** – [Bug]: Subagent completion silently lost — no retry, no notification, no auto-restart on timeout (19 comments, 👍 1)  
  [openclaw/openclaw Issue #44925](https://github.com/openclaw/openclaw/issues/44925)  
  *Underlying need*: Users require guaranteed task execution with explicit failure signals; silent losses break trust in subagent orchestration.

- **Issue #45740** – gh-issues skill: untrusted issue body injected directly into sub-agent prompt (13 comments, 👍 1)  
  [openclaw/openclaw Issue #45740](https://github.com/openclaw/openclaw/issues/45740)  
  *Underlying need*: Stronger input sanitisation for skills that process external content – critical for security-minded deployments.

- **Issue #39604** – Feature: Add `tools.web.fetch.allowPrivateNetwork` to allow private network access (13 comments, 👍 9)  
  [openclaw/openclaw Issue #39604](https://github.com/openclaw/openclaw/issues/39604)  
  *Underlying need*: Users want an opt-in for internal network access, balancing security with automation needs.

- **Issue #85888** – Cron jobs consistently fail with MiniMax 503 overload during early morning (11 comments, 👍 1)  
  [openclaw/openclaw Issue #85888](https://github.com/openclaw/openclaw/issues/85888)  
  *Underlying need*: Scheduled reliability; cron failures that don’t occur on manual triggers point to scheduling/retry logic issues.

- **Issue #40001** – Write tool lacks append mode — isolated cron sessions destroy shared files (11 comments, 👍 1)  
  [openclaw/openclaw Issue #40001](https://github.com/openclaw/openclaw/issues/40001)  
  *Underlying need*: Basic file operation for stateful cron tasks – a clear missing feature causing data loss.

- **Issue #44905** – Discord leaks internal tool-call traces (NO_REPLY, commentary, to=functions) to channel (10 comments, 👍 1)  
  [openclaw/openclaw Issue #44905](https://github.com/openclaw/openclaw/issues/44905)  
  *Underlying need*: Privacy and trust – internal debug output must never reach end-users.

## 5. Bugs & Stability

Several high-severity bugs remain unresolved, many with linked PRs still open:

**P1 / Diamond Lobster (highest impact):**

| Issue | Summary | Fix PR? |
|-------|---------|---------|
| [#44925](https://github.com/openclaw/openclaw/issues/44925) | Subagent completion silently lost – no retry, notification, or restart | Linked PR open |
| [#45740](https://github.com/openclaw/openclaw/issues/45740) | gh-issues skill untrusted body injection into sub-agent prompts | Linked PR open |
| [#41744](https://github.com/openclaw/openclaw/issues/41744) | Feishu read image tool loses media before outbound payload | Linked PR open |
| [#85888](https://github.com/openclaw/openclaw/issues/85888) | Cron jobs fail with MiniMax 503 during early morning (manual triggers succeed) | Linked PR open |
| [#44905](https://github.com/openclaw/openclaw/issues/44905) | Discord leaks internal tool-call traces to channel | No linked PR |
| [#74586](https://github.com/openclaw/openclaw/issues/74586) | AM embedded run aborts `memory_search` tool calls; classifies as timeout despite completion | No linked PR |
| [#43367](https://github.com/openclaw/openclaw/issues/43367) | Multi-agent orchestration unstable: concurrent add/config overwrites, session-lock failures | Linked PR open |
| [#40540](https://github.com/openclaw/openclaw/issues/40540) | `openclaw update` fails with EBUSY on Windows | No linked PR |
| [#43661](https://github.com/openclaw/openclaw/issues/43661) | Session hangs indefinitely on compaction timeout, causing duplicate message sends | No linked PR |
| [#38327](https://github.com/openclaw/openclaw/issues/38327) | "Cannot convert undefined or null to object" with google-vertex/gemini-3.1-pro-preview (regression) | No linked PR |

**P2 / Diamond Lobster (notable regressions):**

- [#43747](https://github.com/openclaw/openclaw/issues/43747) – Memory management in chaos (different users see wildly different behavior) – no fix PR.
- [#46786](https://github.com/openclaw/openclaw/issues/46786) – `tools.elevated.enabled: true` breaks exec routing logic – all calls go to host instead of sandbox (regression) – no fix PR.
- [#43996](https://github.com/openclaw/openclaw/issues/43996) – Sandbox container exits immediately when `no-new-privileges` applied – no fix PR.

Many of these issues have been open since March 2026 and carry the `clawsweeper:no-new-fix-pr` label, indicating no new fix PR has been submitted. This signals a gap between reported problems and available solutions.

## 6. Feature Requests & Roadmap Signals

The community is requesting capabilities that would make OpenClaw more flexible and production-ready:

- **Per-agent cost budget enforcement** ([#42475](https://github.com/openclaw/openclaw/issues/42475)) – gateway-level daily/monthly caps to prevent runaway spend. High interest (👍 1).
- **Per-skill model routing** ([#43260](https://github.com/openclaw/openclaw/issues/43260)) – `model` field in SKILL.md frontmatter for task-specific models. 👍 0 but practical.
- **Telegram reaction triggers** ([#47677](https://github.com/openclaw/openclaw/issues/47677)) – first-class reaction-driven agent wake-up. 👍 2.
- **MathJax/LaTeX support in Control UI** ([#42840](https://github.com/openclaw/openclaw/issues/42840)) – for math formulas. 👍 6.
- **Configurable mediaLocalRoots for image tool** ([#47856](https://github.com/openclaw/openclaw/issues/47856)) – needed by iMessage users. 👍 1.
- **Distributed Agent Runtime** ([#42026](https://github.com/openclaw/openclaw/issues/42026)) – separate control plane from agent compute (RFC). 👍 3.
- **Session resetPrompt** ([#45501](https://github.com/openclaw/openclaw/issues/45501)) – configurable startup message after `/new`.
- **Gateway lifecycle hooks** ([#43454](https://github.com/openclaw/openclaw/issues/43454)) – `onSubagentComplete`, `onToolCallThreshold`, etc. 👍 1.
- **Backup exclude patterns** ([#40786](https://github.com/openclaw/openclaw/issues/40786)) – .gitignore-like patterns for `openclaw backup create`. 👍 0.
- **Provider fallback by failure class** ([#47910](https://github.com/openclaw/openclaw/issues/47910)) – quarantine auth-broken providers. 👍 0.

**Likely for next version**: The beta release focused on security. Next stable (v2026.6.x) may incorporate **per-agent cost budgets** (#42475) and **per-skill model routing** (#43260) as they have clear implementations and community need. The **Telegram reaction triggers** (#47677) and **MathJax support** (#42840) are popular but may be pushed to later milestones.

## 7. User Feedback Summary

**Pain points (dissatisfaction):**

- **Silent failures**: Multiple reports of subagent completions lost, cron jobs failing without error propagation, and session hangs with no recovery. Users feel they cannot trust automation.
- **Session state chaos**: Memory inconsistency between users, `/model` overrides dropped on rollover, and session locks failing under concurrency.
- **Security surprises**: Internal tool traces leaking on Discord, untrusted data injected into prompts, sandbox non-compliance with security policies.
- **Missing basic file operations**: No append mode for `write` tool – a commonly needed feature for stateful tasks like daily notes.
- **Platform-specific regressions**: Windows update (EBUSY) and arm64/memory issues hamper diverse deployments.

**Use cases expressed:**

- Multi-user team environments (memory isolation, cost budgets).
- Bot platforms (Telegram, Discord, Feishu, QQ) with multi-account setups.
- Scheduled cron jobs for memory maintenance and data syncing.
- Mathematical and scientific communication (LaTeX rendering).
- Long-running conversations with frequent session resets.

**Satisfaction**: The high number of contributors (500 PRs updated in 24h) and community engagement (many issue discussions) signal strong interest, but the backlog of unresolved P1 bugs suggests user patience may be wearing thin.

## 8. Backlog Watch

Issues and PRs that have been waiting for maintainer attention for an extended time, often with `clawsweeper:needs-maintainer-review` or `clawsweeper:needs-product-decision`:

- **[#44925](https://github.com/openclaw/openclaw/issues/44925)** – Subagent completion silently lost (P1, diamond lobster). Created March 13, last updated June 11. Needs maintainer review and product decision. No fix PR yet.
- **[#45740](https://github.com/openclaw/openclaw/issues/45740)** – gh-issues skill untrusted injection (P2, diamond lobster). Created March 14. Needs security review and product decision.
- **[#39604](https://github.com/openclaw/openclaw/issues/39604)** – Private network fetch opt-in (P2, diamond lobster). Created March 8. Needs security review and product decision. Very popular (👍 9).
- **[#42475](https://github.com/openclaw/openclaw/issues/42475)** – Per-agent cost budget (P2, off-meta tidepool). Created March 10. Needs product decision.
- **[#40001](https://github.com/openclaw/openclaw/issues/40001)** – Write tool lacks append mode (P1, diamond lobster). Created March 8. Needs product decision.
- **[#43747](https://github.com/openclaw/openclaw/issues/43747)** – Memory management regression (P2, diamond lobster). Created March 12. Needs maintainer review and product decision.
- **[#42026](https://github.com/openclaw/openclaw/issues/42026)** – Distributed Agent Runtime RFC (P2, off-meta tidepool). Created March 10. No maintainer response.

**PRs waiting for maintainer look:**

- **[#47523](https://github.com/openclaw/openclaw/pull/47523)** – “Agents: tighten tool name trust and preflight tool collisions” (size L, P1, ready for maintainer look). Open since March 15. A security-critical clean-up.
- **[#78441](https://github.com/openclaw/openclaw/pull/78441)** – “feat(subagents): forward toolsAllow from sessions_spawn” (size L, ready for maintainer look). Open since May 6. Would give subagent permission control.
- **[#90110](https://github.com/openclaw/openclaw/pull/90110)** – “fix(anthropic): resolve Claude Haiku 4.5 static catalog refs” (size S, ready for maintainer look). Open since June 4. Small but needed fix.

These items represent critical decisions and fixes that have been pending for weeks. The project’s velocity on new features is high, but clearing the backlog of well-characterised P1 bugs and RFCs will be essential to maintain community trust.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: Personal AI Assistant Open-Source Ecosystem
**Date: 2026-06-11**

---

## 1. Ecosystem Overview

The personal AI assistant open-source ecosystem is experiencing an unprecedented surge in development velocity, with seven projects showing high activity simultaneously. The landscape is consolidating around a shared architecture—agent core, sandboxed execution, tool-use framework, and multi-channel gateway—but diverging sharply in target user profiles (enterprise vs. hobbyist vs. embedded) and stability maturity. The critical tension across all projects is between rapid feature expansion and unresolved reliability debt, particularly around scheduled tasks, multi-agent orchestration, and cross-platform compatibility. The ecosystem is maturing rapidly but still lacks a clear category winner; each project occupies a defensible niche with distinct trade-offs between feature richness, security posture, and developer experience.

---

## 2. Activity Comparison

| Project | Issues Updated (24h) | PRs Updated (24h) | Merged/Closed PRs (24h) | Release Today? | Health Signal |
|---------|---------------------|-------------------|------------------------|----------------|---------------|
| **OpenClaw** | 320 | 500 | 104 | ✅ v2026.6.6-beta.1 | Growing pains – high volume but P1 backlog |
| **NanoBot** | 11 | 33 | 19 | ❌ | Healthy, responsive maintainers |
| **Hermes Agent** | 6 | 50 | 3 | ❌ | Healthy but slow merge velocity |
| **PicoClaw** | 5 | 15 | 6 | ✅ nightly | Stable, security-focused |
| **NanoClaw** | 2 | 12 | 4 | ❌ | Moderate, growing skill ecosystem |
| **NullClaw** | 0 | 4 | 0 | ❌ | Low activity, pre-release |
| **IronClaw** | 18 | 50 | 23 | ❌ | Very high activity, reborn v2 push |
| **LobsterAI** | 0 | 25 | 24 | ✅ v2026.6.10 | Strong, cleanup mode post-release |
| **ZeroClaw** | 7 | 50 | 11 | ❌ | High activity, CI/UX focus |
| **TinyClaw** | 0 | 0 | 0 | ❌ | Dormant |
| **Moltis** | 0 | 0 | 0 | ❌ | Dormant |
| **CoPaw** | 16 | 49 | 30 | ✅ v1.1.11 | Very active, Windows issues prominent |

**Key insight:** OpenClaw maintains the largest absolute activity but also the highest unresolved bug count (302 open issues, 10+ P1 bugs). IronClaw and CoPaw are the most responsive to critical bugs, often merging fixes within hours of reports.

---

## 3. OpenClaw's Position

**Advantages vs. peers:**
- **Massive contributor base** – 500 PRs/day, 104 merges/day, indicating broad community investment
- **Most comprehensive feature surface** – Discord, Feishu, Teams, MCP, Codex, ACP, subagent orchestration
- **Strong security hardening release (v2026.6.6-beta.1)** – tightened sandbox, MCP, elevated sender checks
- **Detailed backlog triage** – explicit `needs-product-decision`, `needs-security-review` labels

**Technical approach differences:**
- Uses dual-plugin architecture (skills + MCP) more extensively than NanoBot or PicoClaw
- Subagent spawning model is more granular than Hermes Agent's "Action Runtime" abstraction
- Security model is more restrictive than ZeroClaw (which recently blocked `allowed_tools: []` for delegates)

**Community size (inferred from activity):**
- OpenClaw dwarfs all peers in raw issue/PR volume (320/500 vs. next highest ~50/50)
- But ratio of open bugs to merges (302/104) is worse than NanoBot (4/19), PicoClaw (4/6), or LobsterAI (0/24)
- Maintainer response latency appears higher than IronClaw or CoPaw for critical bugs

**Verdict:** OpenClaw is the "Windows of AI agents" – broadest ecosystem, heaviest maintenance burden, and most vocal user frustration. It leads in capability but trails in stability and decision velocity.

---

## 4. Shared Technical Focus Areas

The following requirements are emerging independently across **multiple projects**, indicating ecosystem-wide gaps:

| Focus Area | Projects Affected | Specific Needs |
|------------|------------------|----------------|
| **Reliable scheduled execution** | OpenClaw, NanoBot, Hermes Agent, CoPaw | Cron failures on timeout (OpenClaw #85888), early termination on subagent spawn (NanoBot #4290), model parameter required errors (Hermes #43899), agent-created tasks don't trigger (CoPaw #5064) |
| **Subagent/multi-agent orchestration reliability** | OpenClaw, NanoBot, Hermes Agent, PicoClaw | Silent completion loss (OpenClaw #44925), duplicate message pushes (PicoClaw #3094), result injection breaks cron (NanoBot #4293), concurrency session-lock failures (OpenClaw #43367) |
| **Security hardening parity** | OpenClaw, Hermes Agent, PicoClaw, NanoClaw, CoPaw | SSRF bypass prevention, tar extraction escapes, Discord trace leaks, untrusted prompt injection, `rm` bypass via Python (CoPaw #5090) |
| **Cross-platform compatibility** | PicoClaw, CoPaw, ZeroClaw, Hermes Agent | Windows path separator bugs (PicoClaw #2472), OpenSSL regression on Windows (CoPaw #5086), 74 Windows test failures (ZeroClaw #7462), `vi` not present in containers (ZeroClaw #7469) |
| **Tool file operations – append mode** | OpenClaw, NanoBot | Missing `write` tool append (OpenClaw #40001); NanoBot's `exec` tool can install pip packages but no native append |
| **Per-agent or per-skill cost controls** | OpenClaw, Hermes Agent | Budget enforcement (OpenClaw #42475), subagent model presets (NanoBot #4291), per-skill model routing (OpenClaw #43260) |
| **Observability / diagnostics** | OpenClaw, PicoClaw, Hermes Agent, IronClaw, CoPaw | Debug trace viewers (PicoClaw #2945), token usage display (CoPaw #4433), `turn_id` for OTel correlation (ZeroClaw #7385), trajectory observers (IronClaw #4588) |

**Implication:** Any project that solves scheduled task reliability, sandbox-level security, or cross-platform CI will have immediate cross-project impact. These are ecosystem bottlenecks, not single-project concerns.

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | Hermes Agent | PicoClaw | IronClaw | CoPaw |
|-----------|----------|---------|--------------|----------|----------|-------|
| **Target user** | Power users, enterprise ops | Hobbyists, Telegram-first | Researchers, Bedrock users | Embedded/edge, Windows devs | Enterprise, NEAR AI ecosystem | Chinese market, consumer desktop |
| **Primary language** | TypeScript/Node.js | Python | Rust | Go | Rust | Python |
| **Gateway focus** | Discord, Feishu, Teams, generic | Telegram, Feishu | WeCom, Matrix, generic | Minimal (CLI + SDK) | Slack, WebUI | WeChat, DingTalk, desktop Tauri |
| **Agent orchestration** | Complex subagent spawning, ACP | Simple subagent spawn | Action Runtime refactor (in progress) | Agent Collaboration Bus (stale PR) | Triggered automations (reborn) | Cowork sessions, scheduled tasks |
| **Security posture** | High (recent hardening release) | Moderate (sandbox basic) | High (constant-time comparisons, SSRF) | High (quick SSRF patch) | Moderate (approval gates, egress audit) | Low (rm bypass unaddressed) |
| **Desktop app** | None (CLI + WebUI) | None (CLI) | None (CLI + TUI) | None (CLI) | WebUI v2 (reborn) | Windows Tauri client |
| **Latest release maturity** | Beta – hardening | Pre-0.3.0 rapid iteration | No recent release | Nightly (unstable) | No release (reborn v2) | Stable v1.1.11 |
| **Community language** | English global | English, Chinese | English | Chinese (sipeed) | English | Chinese (primarily) |

**Key differentiators:**
- **IronClaw** uniquely targets the NEAR AI ecosystem with OAuth flows and private.near.ai integration – a specific platform lock-in that others avoid.
- **CoPaw** is the only project with a polished Windows desktop Tauri client and aggressive Chinese market features (WeChat, DingTalk, Xiaomi MiMo).
- **PicoClaw** occupies the minimal/low-power niche, with a Go codebase optimized for small devices.
- **Hermes Agent** is the only project with serious Rust investment, focusing on sub-millisecond latency for WeCom enterprise users.
- **NanoBot** has the best signal-to-noise ratio: small team, fast merges, low bug count relative to features added.

---

## 6. Community Momentum & Maturity

**Tier 1: Hyperactive (risk of burnout / quality dilution)**
- **OpenClaw** – 500 PRs/day, 302 open issues, 10+ unresolved P1 bugs. High contributor engagement but maintainers are decision-bottlenecked. Risk: community frustration tipping point.
- **Hermes Agent** – 50 PRs/day, 6 issues, but 3 PRs merged – low merge rate relative to update volume suggests review capacity is strained.
- **IronClaw** – 50 PRs/day, 23 merged. Strong velocity but WebUI v2 has 13 new issues filed today alone, many UX/usability.

**Tier 2: Steady iteration (healthy cadence)**
- **NanoBot** – 33 PRs/day, 19 merged, 4 open issues. Best ratio of fixes to open bugs. Pre-0.3.0, feature-complete but not yet stable release.
- **PicoClaw** – 15 PRs/day, 6 merged, 5 issues. Small team, consistent output, quick security patches.
- **CoPaw** – 49 PRs/day, 30 merged, 16 issues. Very high output but Windows desktop bugs are a persistent drag on release quality.
- **LobsterAI** – 25 PRs/day, 24 merged, 0 new issues. Post-release cleanup mode – healthiest backlog of any active project.

**Tier 3: Early/low activity**
- **NanoClaw** – 12 PRs/day, 4 merged. Skills ecosystem growing but maintainer attention is inconsistent (PR #2937 stale for 18 days).
- **NullClaw** – 4 PRs open, 0 merged. Pre-release, no community engagement yet.
- **ZeroClaw** – 50 PRs/day but documentation-heavy; CI and UX fixes dominate. Still maturing core.

**Tier 4: Dormant**
- TinyClaw, Moltis, ZeptoClaw – No activity in 24h. Likely inactive or in winter.

---

## 7. Trend Signals

Extracted from cross-project community feedback and issue patterns:

1. **"Silent failure is the #1 user trust killer"** – Every high-activity project has at least one P1 bug where failures are invisible: subagents lost (OpenClaw), cron jobs silent (NanoBot, CoPaw), model parameter errors (Hermes). Users are demanding explicit failure propagation, not just retry logic.

2. **"Configuration is still too hard"** – Proxy/env loading issues (NanoClaw #2730), editor fallback failures (ZeroClaw #7469), permission scope confusion (OpenClaw #39604, #46786). The next breakthrough will be "clone and run" – zero-config OAuth (CoPaw v1.1.11) is the first mov in this direction.

3. **"Multi-agent is the new default, not a feature"** – Spawn/subagent, delegate modes, trigger automations, Action Runtime – every project is investing in orchestration. But reliability lags: subagent results corrupting cron (NanoBot), session state corruption (OpenClaw), duplicate message delivery (PicoClaw). The ecosystem needs a canonical pattern for safe multi-agent execution.

4. **"Security is becoming a primary UX feature, not afterthought"** – OpenClaw's hardening release, Hermes' constant-time comparisons and SSRF fixes, PicoClaw's quick SSRF patch, ZeroClaw's delegate tool restrictions. Users increasingly evaluate projects on "can I safely connect this to my internal tools?" rather than feature count.

5. **"Windows is the neglected but growing platform"** – PicoClaw's 2-month-old path separator bug, CoPaw's OpenSSL regression, ZeroClaw's 74 Windows test failures. The desktop Windows user base is growing (CoPaw Tauri client, LobsterAI NSIS installer), but CI and testing remain Linux-first.

6. **"Community wants observability as a built-in, not bolt-on"** – Debug trace viewers (PicoClaw, Hermes), token usage badges (CoPaw), OTel span correlation (ZeroClaw), trajectory observers (IronClaw). Users no longer accept black-box agents; they want to inspect every decision.

7. **"Local/free models are the wedge for adoption"** – CoPaw's free model OAuth and Xiaomi MiMo support, Hermes' Gemma 4 reasoning, NanoBot's StepFun ASR. The value proposition is shifting from "best model" to "most flexible model orchestration" with fallback chains.

**Actionable guidance for AI agent developers:**
- Prioritize **explicit failure reporting** over retry loops – users told the truth they can act on.
- Invest in **zero-config onboarding** (OAuth + auto-detection) – this separates leading projects from the pack.
- Build **cross-platform CI** before releasing Windows features – the cost of retrofitting is visible in every project's backlog.
- Treat **security as a UX feature** – sandbox violations and trace leaks erode trust faster than missing features.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest – 2026-06-11

## Today's Overview
NanoBot saw very high development activity over the past 24 hours, with 33 pull requests updated (19 merged/closed) and 11 issues updated (7 closed). No new release was published, but the velocity suggests a near-term release candidate may be forming. The community continues to contribute bug fixes and enhancements across the core agent, sandbox, WebUI, and channel integrations. Four open issues remain active, spanning empty model responses, cron job interruptions, missing context, and a feature request for aggregated subagent notifications. Overall project health is strong, with maintainers actively reviewing and merging contributions.

## Releases
None. No new releases were tagged on this date.

## Project Progress
Nineteen PRs were merged or closed today, advancing multiple areas:

- **Agent & Execution**:  
  - `#4273` – Added `tools.exec.pathPrepend` config for PATH lookup precedence, complementing the existing `pathAppend`.  
  - `#4275` – Fail fast on invalid config files, improving error handling.  
  - `#4272` – Providers now retry and fall back on stream stalled timeouts (closed issue `#4013`).  
  - `#4274` – Scoped `# Recent History` prompt injection by session, fixing cross-session context contamination (issue `#4259`).  

- **WebUI & Storage**:  
  - `#4247` – Auto-compacts transcripts when file exceeds 8 MB limit, preventing full history loss.  
  - `#4278` – Segmented transcript storage for large chats; older turns rotate into immutable segment files, speeding up session loads.  
  - `#4255` – On-demand version check in Settings > About, replacing background polling.  

- **Channels & Providers**:  
  - `#4281` – Added SiliconFlow as a transcription provider (reuses OpenAI‑compatible adapter).  
  - `#4277` – Feishu channel now lazy-loads the `lark_oapi` SDK during gateway startup.  
  - `#4261` – Fixed `max_tokens` vs `max_completion_tokens` for GPT‑5.x models in `OpenAICompatProvider`.  
  - `#4237` – Fixed Bubblewrap sandbox not resetting `$HOME`, breaking tool writes.  

- **Other Merged**:  
  - `#4000` – Added StepFun native ASR provider for transcription.  
  - `#4233` – Version now displayed in WebUI (user‑facing).  
  - `#3934` – `exec` tool can now install pip packages via `pathPrepend` (see `#4273`).  

## Community Hot Topics
The most active discussions revolve around subagent behaviour and model fallback reliability:

- **Issue `#4290` (open) – Cron job ends early when subagent is spawned**  
  A user reports that the main agent stops listening after a subagent finishes, breaking subsequent workflow steps. This directly impacts users relying on cron‑driven automation. No comment yet, but the problem description is clear.  
  *Link: [HKUDS/nanobot Issue #4290](https://github.com/HKUDS/nanobot/issues/4290)*

- **Issue `#4287` (open) – Empty model responses not triggering fallback**  
  During peak hours, DeepSeek returns empty choices; the system classifies this as non‑fallbackable, causing silent failures. A fix PR `#4288` is open immediately addressing this.  
  *Link: [HKUDS/nanobot Issue #4287](https://github.com/HKUDS/nanobot/issues/4287)*

- **Issue `#4279` (open) – Support aggregated notifications for subagents**  
  Real‑time delivery of subagent results can cause LLM hallucinations when many subagents return simultaneously. A user requests batching/aggregation to reduce noise. This is a thoughtful quality‑of‑life request that could influence roadmap.  
  *Link: [HKUDS/nanobot Issue #4279](https://github.com/HKUDS/nanobot/issues/4279)*

- **PR `#4293` (open) – Fix subagent result injection in cron jobs**  
  A direct fix to issue `#4290`, adding `pending_queue` to `process_direct()` so the runner waits for subagent results mid‑turn. This PR is likely to be merged soon given the severity.  
  *Link: [HKUDS/nanobot PR #4293](https://github.com/HKUDS/nanobot/pull/4293)*

## Bugs & Stability
Several bugs were reported or fixed today, ranked by severity:

- **High: Cron job early termination with subagents** (`#4290`, open) – Breaks scheduled automation. Fix in progress via `#4293`.  
- **High: Empty model responses not triggering fallback** (`#4287`, open) – Causes silent failures during peak hours. Fix PR `#4288` open.  
- **High: Cross‑session context contamination** (`#4259`, closed) – Fixed by `#4274`.  
- **Medium: Bubblewrap sandbox doesn't reset `$HOME`** (`#4237`, closed) – Fixed; tool writes would fail inside sandbox.  
- **Medium: Stream stalled for 90+ seconds not recoverable** (`#4013`, closed) – Fixed by retry/fallback logic in `#4272`.  
- **Low: `max_tokens` vs `max_completion_tokens` for GPT‑5.x** (`#4261`, closed) – Fixed model‑specific parameter.  
- **Low: Missing "sustained goal" context** (`#4286`, open) – User reports repeated unexpected missing context after assigning article creation. No fix PR yet, but seems user‑specific configuration issue.

## Feature Requests & Roadmap Signals
User‑requested features from today’s activity:

- **Aggregated subagent notifications** (`#4279`) – Likely to be considered for next minor release to improve LLM reliability when many subagents run in parallel.  
- **Configurable model presets for subagents** (PR `#4291` open) – Allows subagents to use a different model from the parent, specified in a `spawnPresets` list. This is a substantial new capability that may land soon if reviewers approve.  
- **Slack `groupRequireMention`** (PR `#4289` open) – Scopes allowlist channels to @‑mentions only, giving operators finer control.  
- **Skill activation from WebUI slash palette** (PR `#4284` open) – Enhances discoverability of skills, a common UX request.  
- **File management in WebUI settings** (PR `#4282` open) – Lets users browse and modify agent/SOUL config files from the UI, reducing host‑side operations.  

These features indicate the project is maturing its UI/UX while also tackling advanced agent orchestration.

## User Feedback Summary
Based on issue descriptions and PR discussions, user sentiment is generally positive but with clear pain points:

- **Positive**: “0.1.5post2 was very good” (issue `#4013`). Users appreciate the feature set and are actively upgrading.  
- **Dissatisfaction**: The 0.2.0 upgrade caused a regression (stream stalled) that made real work “useless” until retry logic was added (`#4013`). Context contamination (`#4259`) also frustrated users who saw their sessions polluted by unrelated history.  
- **Use cases**: Telegram bot runtime (`#4287`), cron‑driven automation (`#4290`), article creation (`#4286`), and Python script execution via `exec` tool (`#3934`).  
- **Feature requests show demand for better subagent control**, model fallback, and config‑file management from the UI.

Overall, the community is engaged and contributing fixes quickly. The maintainers are responsive, merging multiple PRs per day.

## Backlog Watch
No long‑unanswered important issues were identified in this digest. The oldest open issues among today’s active set date back to 2026-05-20 (`#3934`, closed). The project seems to have a healthy turnaround. However, maintainers should keep an eye on:

- **Issue `#4286` (missing sustained goal context)** – No fix proposed yet; user may need guidance or a deeper investigation.  
- **Issue `#4279` (aggregated subagent notifications)** – While not critical, it could improve multi‑agent reliability and prevent LLM hallucinations; consider scheduling for next milestone.  

No PRs are stalled or awaiting response from maintainers beyond normal review cycles.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-06-11

## 1. Today’s Overview
Project activity remains high, with **50 pull requests updated** and **6 issues reported** in the last 24 hours. No new releases were cut today. The community is heavily focused on **security hardening** (constant‑time signature comparison, SSRF prevention, tar‑extraction escape fixes) and **reliability improvements** (memory‑provider contract fixes, persistent‑state redaction gaps). Several critical bugs surfaced for cron jobs, non‑Claude Bedrock models, and TTS voice playback, but no fix PRs have yet been attached to those reports. Overall, the project is healthy and responsive, with maintainers and contributors actively merging small fixes while larger refactors (e.g., Action Runtime, Matrix gateway parity) continue to evolve.

## 2. Releases
*No new releases were published on 2026‑06‑11.*

## 3. Project Progress
Three pull requests were **merged or closed** today:

- **#43602** [docs] – Clarified that Baoyu image‑generation skills use Hermes’ `image_generate` and prohibited SVG/HTML/manual‑rendering alternatives.  
  🔗 [NousResearch/hermes-agent #43602](https://github.com/NousResearch/hermes-agent/pull/43602)

- **#43948** [fix] – Fixed `MemoryProvider.on_turn_start` so that the live model name is passed as a kwarg, aligning the call site with the documented contract.  
  🔗 [NousResearch/hermes-agent #43948](https://github.com/NousResearch/hermes-agent/pull/43948)

- **#21105** (issue) – The renamed root‑profile alias bug was closed after a fix was merged earlier (the associated PR is not shown in the top‑20 list).  
  🔗 [NousResearch/hermes-agent #21105](https://github.com/NousResearch/hermes-agent/issues/21105)

Additionally, issue **#43601** (Baoyu skills guidance) was closed as resolved by the docs PR.

## 4. Community Hot Topics
The most active discussions and work‑in‑progress revolve around:

- **Action Runtime refactor (#43039)** – A large PR unifying exec handlers behind a single schema, making gateway exec “honest” (failed actions are always detectable) and landing a typed `SessionState`. It has been updated today and is a significant architectural change.  
  🔗 [NousResearch/hermes-agent #43039](https://github.com/NousResearch/hermes-agent/pull/43039)

- **Matrix gateway parity (#18505, #18506)** – Two stacked PRs (1/3 and 2/3) that add room isolation, native tools, and reaction controls for the Matrix platform. They have been active since May 1 and continue to receive updates, indicating strong community interest in Matrix support.  
  🔗 [NousResearch/hermes-agent #18505](https://github.com/NousResearch/hermes-agent/pull/18505)  
  🔗 [NousResearch/hermes-agent #18506](https://github.com/NousResearch/hermes-agent/pull/18506)

- **Security fixes** – Multiple PRs from contributor `zapabob` today address WeCom callback signature comparison (#43937), SSRF in yuanbao download URL (#43938), and curator rollback tar extraction harden (#43942). These reflect a proactive security posture.  
  🔗 [NousResearch/hermes-agent #43937](https://github.com/NousResearch/hermes-agent/pull/43937)  
  🔗 [NousResearch/hermes-agent #43938](https://github.com/NousResearch/hermes-agent/pull/43938)  
  🔗 [NousResearch/hermes-agent #43942](https://github.com/NousResearch/hermes-agent/pull/43942)

- **Persistence redaction gaps (#43940)** – A fix closes three boundary‑redaction holes where credentials could leak into `state.db`.  
  🔗 [NousResearch/hermes-agent #43940](https://github.com/NousResearch/hermes-agent/pull/43940)

Underlying needs: the community is demanding **robust security**, **cross‑platform parity** (especially Matrix and Telegram), and **reliable execution** (Action Runtime, memory, context preservation).

## 5. Bugs & Stability
Several bugs were reported today, ranked by estimated severity:

| Issue | Severity | Description | Fix PR Exists? |
|-------|----------|-------------|----------------|
| **#43899** | **P1** – Cron jobs fail with `Model parameter is required` even when `config.yaml` has a default model. Blocks all scheduled automation. | No |
| **#43946** | **P1** – Non‑Claude Bedrock models (Kimi, DeepSeek, Amazon Nova) fail with malformed tool/message payloads via the Bedrock provider. Breaks all Bedrock traffic for non‑Claude models. | No |
| **#43944** | **P1** – Voice conversation TTS cuts off mid‑message; speech buffer resets lose remaining text. Affects voice‑mode UX. | No |
| **#43951** | **P3** – `vision_analyze` tool is exposed to models that lack vision capability, causing unnecessary fallback to text mode. | No |
| **#21105** | – (Closed) Renamed root profile alias was not honored; now fixed. | ✅ Closed |

All three P1 bugs lack an associated fix PR as of today. The Bedrock and cron issues appear related to model‑picker logic and provider‑specific payload construction.

## 6. Feature Requests & Roadmap Signals
Notable feature requests and signals from today’s data:

- **Auto‑disable vision tools for non‑vision models (#43951)** – A clean UX improvement that could land quickly. Likely to be included in the next patch release.
- **Telegram config‑driven callback‑to‑text routes (#43949)** – Allows inline keyboard `callback_data` to be injected as user text. A concrete enhancement for Telegram gateway users.
- **Gemma 4 reasoning token normalization (#43950)** – Normalises non‑standard control tokens so they render correctly in the UI. Important for local provider compatibility.
- **Opt‑in per‑user memory scope via `X-Hermes-User-Id` header (#34609)** – An architectural feature for multi‑tenant deployments. Still open since May 29.
- **Egress security audit / doctor (#35149)** – Follow‑up to a security review, adding audit and Anthropic‑native egress improvements. Signals a push for production‑grade network security.

Predictions for next version: Bedrock compatibility fixes, vision tool auto‑disable, Gemma 4 token support, and the Action Runtime refactor (if it matures) could all appear in the next minor release.

## 7. User Feedback Summary
Real user pain points surfaced through issue reports:

- **Cron automation broken (#43899)** – Users expect the global config default to be respected when no explicit model is set.
- **Non‑Claude Bedrock models unusable (#43946)** – Users who want to use cost‑effective or region‑specific models on Bedrock face 400 errors.
- **TTS truncation in voice mode (#43944)** – Content after code blocks or unpunctuated sentences is lost, breaking read‑aloud functionality.
- **Profile rename ignored (#21105, now fixed)** – Users customizing profiles encountered unexpected alias behaviour; now resolved.

On the positive side, the community is actively contributing fixes—especially around security (zapabob’s three PRs) and memory reliability (diskerror’s #43948). This demonstrates high satisfaction with the project’s development velocity and openness.

## 8. Backlog Watch
Issues and PRs that have remained open for an extended period and may require maintainer attention:

- **#18505 / #18506 (Matrix parity, since May 1)** – Stacked PRs with many updates but no merge yet. Dependencies between them may need coordination.
- **#26051 (context preservation on compression failures, since May 15)** – Open for nearly a month; addresses a subtle bug in conversation history handling.
- **#31477 (agent recovery‑path blank stream, since May 24)** – Fixes a critical blank‑stream gap in the conversation loop. No merge yet.
- **#34460 (per‑thread SessionDB connections, since May 29)** – Could significantly improve concurrency performance, but still awaiting review.
- **#35149 (egress audit, since May 30)** – A security‑focused PR from an invited contributor; may need maintainer bandwidth to review.
- **#43039 (Action Runtime refactor, since June 9)** – A large, cross‑cutting change that touches CLI, gateway, and TUI. Likely requires careful review and testing.

These languishing items risk accumulating technical debt if not addressed soon, particularly the concurrency and context‑preservation fixes.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest — 2026-06-11

## 1. Today's Overview

Project activity remains high, with **5 issues updated** (4 open, 1 closed) and **15 pull requests** (9 open, 6 merged/closed) in the last 24 hours. A new **nightly build (v0.2.9-nightly)** was published, incorporating several bug fixes and a feature addition. The merged PRs focus on Windows compatibility, SSRF protection hardening, error handling, and OpenAI API compatibility. A security vulnerability (SSRF bypass via `198.18.0.0/15`) was closed with a fix merged. Overall project health is stable, with active maintainer engagement and a steady flow of community contributions.

## 2. Releases

**nightly** — automated build for `v0.2.9-nightly.20260611.d955d5bb`
- This is an **automated, potentially unstable** build. No guaranteed breaking changes are documented, but users should treat it as a work-in-progress.
- Full changelog: [compare/v0.2.9...main](https://github.com/sipeed/picoclaw/compare/v0.2.9...main)
- **Note:** No versioning or migration notes are provided for nightly releases.

## 3. Project Progress

**6 pull requests were merged/closed today** (all merged, none rejected):

- **[#3089](https://github.com/sipeed/picoclaw/pull/3089)** — `fix os.Root api on windows issue`  
  Resolves `list_dir` failure on Windows caused by path separator mismatch (`#2472`). Merged.

- **[#3085](https://github.com/sipeed/picoclaw/pull/3085)** — `fix(tools): block 198.18.0.0/15 in SSRF guard`  
  Addresses security advisory `#3077`. Blocks RFC 2544 benchmark address range in `web_fetch` SSRF protection. Merged.

- **[#3043](https://github.com/sipeed/picoclaw/pull/3043)** — `fix: check strconv.Atoi and json.Unmarshal errors`  
  Silently discarded errors in `seahorse` and `seahorse/short_retrieval` now properly handled. Merged.

- **[#2951](https://github.com/sipeed/picoclaw/pull/2951)** — `fix: use function-type web_search for better API compatibility`  
  Fixes HTTP 400 errors when using native web search on OpenAI endpoints that reject `web_search_preview`. Merged.

- **[#2948](https://github.com/sipeed/picoclaw/pull/2948)** — `fix: skip temperature parameter for claude-opus-4-7 models`  
  Prevents HTTP 400 errors for Anthropic models that deprecate temperature. Merged.

- **[#2945](https://github.com/sipeed/picoclaw/pull/2945)** — `feat: add debug trace viewer (picoclaw-tracer)`  
  New standalone web UI for real-time LLM trace inspection. Merged.

These changes improve **cross-platform stability, security, error resilience, model compatibility, and debugging capability**.

## 4. Community Hot Topics

- **Issue [#2472](https://github.com/sipeed/picoclaw/issues/2472)** — `[BUG] list_dir returns "invalid argument" on Windows due to path separator mismatch with os.Root`  
  **5 comments, 1 👍**. A long-standing Windows-specific bug (created April 2026) that prevented file system tools from working on Windows. The fix has now been merged in PR #3089.

- **Issue [#3094](https://github.com/sipeed/picoclaw/issues/3094)** — `[Bug] 异步子代理(spawn)任务完成时，ForUser字段被同时用于直接推送和主代理汇总，导致重复消息`  
  New today, no comments yet, but addresses a **duplicate message** issue in asynchronous subagent (spawn) workflows. This is likely to attract attention from users relying on multi-agent patterns.

- **Issue [#3077](https://github.com/sipeed/picoclaw/issues/3077)** — `[Security] PicoClaw web_fetch SSRF restriction bypass via special-use IPv4 literal 198.18.0.0/15`  
  **Closed and fixed** (PR #3085). Quick turnaround from report to patch.

- **PR [#2937](https://github.com/sipeed/picoclaw/pull/2937)** — `Feat/agent collaboration` (stale, open since May 24)  
  A large feature adding a “first-class internal Agent Collaboration Bus” with per-agent mailboxes and collaboration threads. No new activity in the last 24h, but it remains the most significant open feature PR.

## 5. Bugs & Stability

| Severity | Issue | Description | Fix |
|----------|-------|-------------|-----|
| **High** | [#3094](https://github.com/sipeed/picoclaw/issues/3094) | Duplicate messages when asynchronous subagent completes (ForUser field double-pushed) | No fix PR yet |
| **Medium** | [#3090](https://github.com/sipeed/picoclaw/issues/3090) | Panel does not work on Safari iOS <16.4 | No fix PR yet |
| **Low** | [#2472](https://github.com/sipeed/picoclaw/issues/2472) | `list_dir` fails on Windows (now fixed) | PR #3089 merged |
| **Critical** | [#3077](https://github.com/sipeed/picoclaw/issues/3077) | SSRF bypass via `198.18.0.0/15` (closed) | PR #3085 merged |

The most impactful active bug is **#3094** (duplicate messages), which can confuse users of the spawn tool. No immediate fix is available.

## 6. Feature Requests & Roadmap Signals

- **Issue [#3093](https://github.com/sipeed/picoclaw/issues/3093)** — `[Feature] I need SimpleX or tox`  
  A user request for support of two additional messaging protocols (SimpleX and Tox) as integration gateways. This signals demand for **decentralized/private communication channels**. Unlikely to appear in the next nightly, but could influence future release planning.

- **PR [#2945](https://github.com/sipeed/picoclaw/pull/2945)** — `feat: add debug trace viewer (picoclaw‑tracer)` has been merged, indicating **developer tooling improvements** are a priority.

- **PR [#2937](https://github.com/sipeed/picoclaw/pull/2937)** — **Agent collaboration** remains open and stale. If reviewed and merged in the coming weeks, it would be a major v0.3.0 feature.

- Several smaller feature PRs (e.g., `feat(web): harden launcher access control` [#3083](https://github.com/sipeed/picoclaw/pull/3083)) are still open, pointing to ongoing **security hardening** efforts.

## 7. User Feedback Summary

- **Pain points (Windows users):** The long-standing path separator bug (now fixed) caused frustration for Windows users of file system tools. The fix is a positive signal.
- **Async subagent users:** Duplicate messages ([#3094](https://github.com/sipeed/picoclaw/issues/3094)) degrade UX for multi-step workflows. This is a newly reported issue that requires prompt attention.
- **Safari/iOS users:** Inability to use the Panel on older iOS versions ([#3090](https://github.com/sipeed/picoclaw/issues/3090)) may limit mobile access for some users.
- **Security-conscious users:** Quick response to SSRF vulnerability ([#3077](https://github.com/sipeed/picoclaw/issues/3077)) demonstrates responsible maintenance.
- **General satisfaction:** No complaints about feature completeness or performance beyond the reported bugs.

## 8. Backlog Watch

| Item | Type | Age / Status | Reason for Concern |
|------|------|--------------|--------------------|
| [#2937](https://github.com/sipeed/picoclaw/pull/2937) `Feat/agent collaboration` | PR | Open since May 24, stale | Major feature awaiting review; no activity for 18 days. Risk of bitrot or merge conflicts. |
| [#3067](https://github.com/sipeed/picoclaw/pull/3067) `fix: add DmScope field to SessionConfig` | PR | Open since June 9 | Fixes a UI config save regression; no objections but not yet merged. |
| [#3045](https://github.com/sipeed/picoclaw/pull/3045) `fix(identity): allow_from fallthrough for Matrix user IDs with colon` | PR | Open since June 7 | Affects Matrix users; no maintainer comment in 4 days. |
| [#3087](https://github.com/sipeed/picoclaw/pull/3087) `fix(tools): allow workspace relative exec paths` | PR | Open since June 9 | Fixes false positive in exec safety guard; no maintainer feedback yet. |
| [#2472](https://github.com/sipeed/picoclaw/issues/2472) `list_dir Windows bug` (now fixed) | Issue | Closed today after 2 months | Satisfactory resolution, but illustrates slow turnaround for platform-specific issues. |

*All links are to `https://github.com/sipeed/picoclaw` paths as shown.*

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-11

## 1. Today’s Overview

NanoClaw saw moderate activity over the past 24 hours, with **2 open issues**, **12 pull requests** updated (4 merged/closed), and **no new releases**. The community is actively contributing both bug fixes and new skill modules; the most significant technical conversation revolves around a multi-runtime agent SDK abstraction (#1690). Several PRs address concrete operational bugs, including a high-severity issue where egress lockdown blocks host-local services (#2731) and a missing `.env` loading problem (#2730). The project’s skill ecosystem continues to expand with proposals for guardrails, web-search, and log persistence.

---

## 2. Releases

No new releases were published today.

---

## 3. Project Progress (Merged / Closed PRs)

Four PRs were merged or closed in the last 24 hours:

- **[#2719] feat: add uninstall.sh — per-copy uninstaller**  
  Merged. Adds a standalone uninstall script with confirmation, dry-run, and OneCLI agent cleanup.  
  *Author: amit-shafnir* | [GitHub](https://github.com/nanocoai/nanoclaw/pull/2719)

- **[#2721] docs: customizing intro, skills model, and skill guidelines**  
  Merged. Provides three new public docs that formalise the skills-based customisation contract.  
  *Author: gavrielc* | [GitHub](https://github.com/nanocoai/nanoclaw/pull/2721)

- **[#2724] Opened against wrong repo — disregard**  
  Closed immediately as a mistake. No code changes.  
  *Author: Pineacles* | [GitHub](https://github.com/nanocoai/nanoclaw/pull/2724)

- **[#3] Secure IPC with per-group namespaces**  
  Merged (older PR, updated today). Enforces isolation by giving each container its own IPC directory, preventing privilege escalation.  
  *Author: gavrielc* | [GitHub](https://github.com/nanocoai/nanoclaw/pull/3)

---

## 4. Community Hot Topics

The most active issue and PRs in terms of discussion and reactions:

### Issue #1690 — Multi-runtime agent SDK abstraction
- **Comments: 6 | 👍: 3**  
- **Summary:** Proposes an `AgentRuntime` interface that allows different agent SDKs (Claude, Codex, local models) to be installed as modular skills, mirroring the channel pattern (`/add-telegram`).  
- **Underlying need:** Users want to swap or combine agent backends without fork-and-fight, and the existing skill model is seen as a natural extension.  
- **Link:** [Issue #1690](https://github.com/nanocoai/nanoclaw/issues/1690)

### Issue #2731 — Egress lockdown breaks host.docker.internal
- **New today, 0 comments** but high-severity signal. It reveals a critical regression where `NANOCLAW_EGRESS_LOCKDOWN=true` attaches the gateway container to a network that disconnects agents from host-local services (Ollama, proxies).  
- **Link:** [Issue #2731](https://github.com/nanocoai/nanoclaw/issues/2731)

### PR #2211 — Tool-visibility skill (live tool-call previews)
- **Open since May 3, updated today.** Implements hooks (`PreToolUse`/`PostToolUse`) as a self-contained skill. The long open duration suggests maintainers are weighing its integration.  
- **Link:** [PR #2211](https://github.com/nanocoai/nanoclaw/pull/2211)

---

## 5. Bugs & Stability

Three bugs were reported or fixed today, ranked by severity:

| Severity | Issue/PR | Description | Fix Status |
|----------|----------|-------------|------------|
| **High** | Issue #2731 | `NANOCLAW_EGRESS_LOCKDOWN` hijacks `host.docker.internal`, making agents unreachable inside Docker. | No fix PR yet |
| **Medium** | PR #2730 | `.env` file variables (e.g., `NANOCLAW_EGRESS_LOCKDOWN`) are never loaded under `launchd/systemd`. A PR fixes the boot sequence. | **Open** |
| **Medium** | PR #2728 | Telegram pairing with `--intent wire-to:<folder>` does not create the expected database row (`messaging_group_agents`). | **Open** |
| **Low** | PR #2729 | Documentation mismatch in the `/add-telegram` skill: status block names and adapter pin are wrong. | **Open** |

Additionally, **PR #2611** (security fix) preserves the original caller context after approval-gated commands are replayed — this addresses a potential privilege issue when using `ncl` approval flow. That PR remains open.

---

## 6. Feature Requests & Roadmap Signals

Several feature PRs and issues point to likely future additions:

- **Multi-runtime abstraction** (#1690) — If merged, this would be a major architectural change, allowing users to plug in different agent providers. High community interest.
- **Guardrails skill** (PR #2726) — Per-agent-group input/output guardrails (regex/keyphrase blocking, quarantine audit trail). Replies to demand for safety in multi-user deployments.
- **Container log persistence** (PR #2727) — Stores `stdout`/`stderr` from agent containers to disk. Important for debugging and auditing.
- **Web-search-plus skill** (PR #2725) — Multi-provider web search + URL extraction, no MCP. Expands agent capabilities.
- **Tool-visibility skill** (PR #2211) — Live tool-call previews in chat. Improves transparency.

**Prediction for next release (v0.x):** The guardrails skill and container log persistence have the highest momentum (both authored by active contributors and aligned with security/observability needs). The multi-runtime abstraction may be deferred to a larger 1.0 milestone.

---

## 7. User Feedback Summary

From today’s issues and PRs, the following real pain points emerged:

- **“Egress lockdown breaks my host services”** — Users relying on Ollama or local proxies behind `host.docker.internal` are completely blocked when lockdown is enabled. The workaround is to disable lockdown, which defeats its purpose.
- **“.env variables don’t work under systemd”** — Users following the documented setup (`NANOCLAW_EGRESS_LOCKDOWN=true`) find the flag is never read, causing silent misconfiguration.
- **“Telegram pairing seems to succeed but doesn’t wire agents”** — The `wire-to` intent logs a success but no actual agent-group connection is created, leaving users confused.
- **“No easy way to uninstall”** — Addressed by the newly merged `uninstall.sh` skill (PR #2719).

Overall sentiment is positive toward the skill ecosystem, but the volume of open bugs indicates the project is still maturing its edge cases.

---

## 8. Backlog Watch

These important items have been open for more than a week with no resolution or recent maintainer comment:

| Item | Age | Description | Maintainer Action Needed |
|------|-----|-------------|--------------------------|
| #1690 | 2 months | Multi-runtime SDK abstraction – active discussion but no maintainer response since creation. | Acknowledge and indicate roadmap fit. |
| #2211 | 5 weeks | Tool-visibility skill – updated today but no maintainer review. | Review and merge or provide feedback. |
| #2611 | 2 weeks | Security fix for caller context preservation – critical for approval workflows. | Request missing information or merge. |
| #2727 | 1 day | Container log persistence – new, but author is from Microsoft (sibling PR in amplifier-app). | Likely needs cross-repo alignment; should be triaged. |

No issues older than 30 days with zero maintainer attention were noted.

---

*Digest generated from GitHub data on 2026-06-11. All links refer to the [nanocoai/nanoclaw](https://github.com/nanocoai/nanoclaw) repository.*

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-06-11

## 1. Today's Overview
The project saw no new issues or releases in the last 24 hours, and all four recently updated pull requests remain open with no merges or closures. This low activity level suggests a temporary lull in development or maintainer focus. The open PRs target bug fixes in agent output handling, queue configuration, cron delivery attribution, and gateway testing, indicating ongoing maintenance efforts. No community discussion or user feedback has been recorded for any of these changes.

## 2. Releases
No new releases were published today.

## 3. Project Progress
No pull requests were merged or closed in the last 24 hours. The four open PRs (all created on 2026-06-10) are:
- **#951** – Fix to suppress stderr initialization logs on agent failure (author: vernonstinebaker)
- **#949** – Adds `agent.default_queue_mode` config field and moves QueueMode enum to shared config (author: vernonstinebaker)
- **#948** – Fixes cron agent delivery attribution (author: DonPrus)
- **#950** – Fixes gateway port probe ordering to prevent resource leaks (author: addadi)

None of these have been reviewed, commented on, or reacted to.

## 4. Community Hot Topics
No issues or pull requests have received any comments or reactions in the last 24 hours. All activity is developer-driven with no observable community engagement at this time.

## 5. Bugs & Stability
Four bug‑fix PRs are currently open, none yet merged. Ranked by potential impact:
- **High:** [#951](https://github.com/nullclaw/nullclaw/pull/951) – Agent failure could post initialization logs (memory plan, MCP server registration) to channels as false responses. Fix pending.
- **Medium:** [#950](https://github.com/nullclaw/nullclaw/pull/950) – Gateway test runs failing with `AddressInUse` could leak memory if port probe happens after allocations. Fix reorders the probe.
- **Low:** [#948](https://github.com/nullclaw/nullclaw/pull/948) – Cron‑delivered agent responses may lack correct channel/account attribution; patch ensures metadata flows into spawned subprocesses.
- **Enhancement:** [#949](https://github.com/nullclaw/nullclaw/pull/949) – Makes queue mode configurable; fixes a potential gap where default mode was hardcoded.

No crashes or regression reports have been filed outside these open PRs.

## 6. Feature Requests & Roadmap Signals
The only feature‑adjacent change is **PR #949**, which introduces a config‑file option for `agent.default_queue_mode` and refactors the `QueueMode` enum to a shared location. This suggests the team is moving toward more user‑configurable agent behavior. If merged, this could pave the way for per‑session queue mode control in the next release. No other feature requests were recorded.

## 7. User Feedback Summary
No user issues, comments, or reactions were posted in the last 24 hours. The project’s current interaction base appears to be limited to its core contributors. No pain points or use cases are observable from today’s data.

## 8. Backlog Watch
No issues exist in the repository (total: 0 items). All open PRs are brand new (created yesterday) and have not yet received maintainer attention. No long‑unanswered items require escalation at this time.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-06-11

## 1. Today’s Overview

IronClaw saw extremely high activity over the past 24 hours, with **50 PRs** and **18 issues** updated. The reborn (v2) stack dominated the day’s work: 23 PRs were merged or closed, fixing provider configuration persistence, auth-gate resume for OAuth flows, Slack DM delivery, and attachment upload UX. However, a flood of new issues (13 opened today) surfaced usability gaps in the WebUI – from missing syntax highlighting and small font sizes to opaque error messages and broken NEAR AI login flows for local builds. The overall health is strong: the core team is responding quickly with fix PRs, but first-run and diagnostic quality still need polish before the reborn release feels production-ready.

## 2. Releases

No new releases were published today.

## 3. Project Progress (Merged/Closed PRs Today)

The following non‑trivial PRs were merged or closed, advancing key reborn features and stability:

- **[#4746 – Auth-gate resume: re-dispatch original capability call after OAuth completes](https://github.com/nearai/ironclaw/pull/4746)** (XL, core)  
  Fixes “next meeting” returning stale memory by persisting `PendingAuthResume` records, mirroring the approval-gate resume pattern.

- **[#4745 – Refactor automations panel facade with TriggerRepository](https://github.com/nearai/ironclaw/pull/4745)** (XL, core)  
  Replaces capability dispatch read path with a direct repository query, simplifying the panel and removing a synthetic trust decision.

- **[#4743 – Fix NEAR context overflow classification](https://github.com/nearai/ironclaw/pull/4743)** (M, core)  
  Classifies NEAR’s `prompt is too long` 400 as `ContextLengthExceeded` with token counts, preventing silent failures.

- **[#4742 – Fix manual token runtime credential selection](https://github.com/nearai/ironclaw/pull/4742)** (L, core)  
  Ensures manual‑token credentials can satisfy runtime requests without stored provider‑scope metadata.

- **[#4652 – Document Reborn serve/WebUI testing flow + launcher script](https://github.com/nearai/ironclaw/pull/4652)** (M, regular)  
  Adds a full walk‑through and one‑command `run-reborn-webui.sh` launcher.

- **[#4730 – Personal triggered‑event delivery: Slack DM end to end](https://github.com/nearai/ironclaw/pull/4730)** (XL, core)  
  Completes the personal scope for triggered automations – Slack pairing, DM provisioning, delivery of final replies/approvals/auth notices.

- **[#4739 – Enable Slack for QA Railway Reborn](https://github.com/nearai/ironclaw/pull/4739)** (S, core)  
  Activates Slack integration in Docker configs used by Railway QA deployments.

- **[#4717 – Restore WebUI v2 always‑approval affordance](https://github.com/nearai/ironclaw/pull/4717)** (XL, experienced)  
  Adds `allow_always` for typed approval gates, preserving persistent approval policies for known‑safe tool calls.

## 4. Community Hot Topics

Most commented / reacted issues in the last 24 hours:

- **[#3036 – [EPIC] Configuration‑as‑Code for IronClaw Reborn](https://github.com/nearai/ironclaw/issues/3036)** (6 comments, 👍1)  
  This long‑running epic describes the need for declarative tenant blueprints and harnesses. The underlying ask: operators want a single schema to replace the current mix of `.env`, workspace docs, settings JSON, and runtime flags. Still awaiting implementation signal.

- **[#4703 – Conversation cannot use NEAR AI provider after successful setup](https://github.com/nearai/ironclaw/issues/4703)** (2 comments)  
  A user reports that clicking “Save” after testing the NEAR AI provider does not persist configuration. This is a high‑friction onboarding blocker – fixed by PR **[#4731](https://github.com/nearai/ironclaw/pull/4731)** (still open but addresses the root cause).

- **[#4632 – [Epic] Build out Reborn WebUI v2 end‑to‑end smoke coverage](https://github.com/nearai/ironclaw/issues/4632)** (0 comments, but many child issues)  
  A roadmap signal: the project aims to add browser‑driven full‑stack E2E tests, which would catch the kind of UX regressions reported today.

## 5. Bugs & Stability

Thirteen new issues filed today. Severity ranking:

| Severity | Issue | Summary | Fix in progress? |
|----------|-------|---------|------------------|
| 🔴 Critical | [#4741](https://github.com/nearai/ironclaw/issues/4741) | Opaque “Invalid master key” on corrupt secret‑store key file in local dev | No |
| 🔴 Critical | [#4729](https://github.com/nearai/ironclaw/issues/4729) | NEAR AI login broken for local/desktop builds – `private.near.ai` rejects non‑`private.near.ai` callback origins | No |
| 🟠 High | [#4703](https://github.com/nearai/ironclaw/issues/4703) | NEAR AI provider not persisted after “Save” | PR [#4731](https://github.com/nearai/ironclaw/pull/4731) |
| 🟠 High | [#4704](https://github.com/nearai/ironclaw/issues/4704) | `builtin.http` approval loop after `invalid_input` failure without actionable details | No |
| 🟠 High | [#4683](https://github.com/nearai/ironclaw/issues/4683) | Generic “driver unavailable” error for invalid model config | No |
| 🟡 Medium | [#4701](https://github.com/nearai/ironclaw/issues/4701) | Approval modal lacks context for `builtin.http` tool requests | No |
| 🟡 Medium | [#4708](https://github.com/nearai/ironclaw/issues/4708) | Code blocks have no syntax highlighting | No |
| 🟡 Medium | [#4707](https://github.com/nearai/ironclaw/issues/4707) | Font size too small in WebUI conversation page | No |
| 🟡 Medium | [#4733](https://github.com/nearai/ironclaw/issues/4733) | Clicking response links navigates away from conversation | No |
| 🟡 Medium | [#4740](https://github.com/nearai/ironclaw/issues/4740) | Slack tool `parameters_schema` missing types for `channel`, `limit`, etc. | No |
| 🔵 Low | [#4700](https://github.com/nearai/ironclaw/issues/4700) | NEAR AI MCP should auto‑enable when credentials are configured | No |
| 🔵 Low | [#4734](https://github.com/nearai/ironclaw/issues/4734) | Agent avatar shows “IC” instead of IronClaw icon | Closed (patched) |
| 🔵 Low | [#4594](https://github.com/nearai/ironclaw/issues/4594) | Unsupported config diagnostics for not‑yet‑wired settings | Closed |

**Notable:** Three closed issues show the team is responsive, and the fix PRs for #4703, #4746 (auth‑gate), #4717 (always‑approval) address some of the highest‑impact blocker reports.

## 6. Feature Requests & Roadmap Signals

- **[#3036 – Configuration‑as‑Code epic](https://github.com/nearai/ironclaw/issues/3036)** remains open with no implementation yet, but it is labelled “reborn” and “suggested_P2”. Likely targeting a post‑Beta release.

- **[#4632 – E2E smoke coverage epic](https://github.com/nearai/ironclaw/issues/4632)** signals the team plans browser‑driven tests after the current manual UX churn stabilises.

- **[#4700 – Auto‑enable NEAR AI MCP](https://github.com/nearai/ironclaw/issues/4700)** when credentials are present – a small UX win, likely to be picked up soon.

- **[#4735 – Programmatic MCP server config + PATCH update](https://github.com/nearai/ironclaw/pull/4735)** (open PR) extends the Extensions API to allow one‑shot server setup and incremental updates without teardown. This is a contributor PR and could land in the next minor release.

- **[#4588 – Observability seams (trajectory observer + LLM provider injection)](https://github.com/nearai/ironclaw/pull/4588)** (open PR) adds hooks for external hosts like `nearai-bench`. Suggests growing interest in benchmarking and observability.

## 7. User Feedback Summary

Real users surfaced several pain points during local testing (`think-in-universe`, `sunglow666`, `abbyshekit`, `pranavraja99`):

- **Onboarding friction:** NEAR AI provider setup fails silently (Save not working, login origin blocked). Users have to re‑enter credentials repeatedly.
- **Poor diagnostics:** “driver unavailable” and “invalid_input” errors give no actionable details; the secret‑store “Invalid master key” message lacks a fix hint.
- **UX regressions:** No syntax highlighting, very small font, approval modals showing only generic buttons, and external links navigating away from the chat.
- **Tool configuration:** Slack tool’s incomplete parameter schema forces models to guess types, causing incorrect tool calls.
- **Positive signal:** Users are actively testing and reporting, and several fixes merged the same day (e.g., #4734 icon, #4673 provider save via #4731) show the team values this feedback.

## 8. Backlog Watch

- **[#3036 – Configuration‑as‑Code epic](https://github.com/nearai/ironclaw/issues/3036)** opened 2026‑04‑28, last updated 2026‑06‑10. Despite being an epic with 6 comments and a thumbs‑up, no PRs are linked. It needs maintainer triage to decide milestone or design phase.

- **[#3708 – chore: release](https://github.com/nearai/ironclaw/pull/3708)** opened 2026‑05‑16, still open. This is the release PR that includes API‑breaking changes in `ironclaw_common` and `ironclaw_skills`. The project has not cut a release in over three weeks, which may delay downstream consumers.

- **[#4559 – Agent‑driven Trace Commons onboarding](https://github.com/nearai/ironclaw/pull/4559)** (XL, opened 2026‑06‑08) is a complex PR with no recent activity on the issue side, but the PR is still open. It could benefit from reviewer bandwidth.

- **[#4499 – Dependabot tokio‑ecosystem bump](https://github.com/nearai/ironclaw/pull/4499)** (opened 2026‑06‑05) is low‑risk but waiting for merge – typical backlog for dependency updates.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

## LobsterAI Project Digest – 2026-06-11

### 1. Today's Overview
The project saw high activity with **25 pull requests updated in the last 24 hours**, of which **24 were merged or closed** and **1 remains open**. A new patch release (`2026.6.10`) was published yesterday, bundling major features including user data backup/migration, local callback login, and task completion notifications. No new issues were created or updated, indicating a focus on stabilizing the codebase after the recent release. The overall project health is strong, with continuous feature development and bugfixes.

### 2. Releases
**New version:** [LobsterAI 2026.6.10](https://github.com/netease-youdao/LobsterAI/releases/tag/2026.6.10) (released 2026-06-10)

**What's changed (from release notes):**
- **Data migration:** Added user data backup and restore functionality.
- **Auth:** Introduced local callback login flow for improved authentication handling.
- **Settings:** Surfaces for OpenCla (context compaction settings) added.
- Bundled in the same release PR ([#2140](https://github.com/netease-youdao/LobsterAI/pull/2140)): task completion notifications, Windows update fixes, and UI refinements.

**Breaking changes / migration notes:** None documented.

### 3. Project Progress
Today’s merged/closed PRs show steady progress across multiple areas:

- **Computer Use MVP** ([#2143](https://github.com/netease-youdao/LobsterAI/pull/2143)) – Adds a Windows x64 built-in Computer Use kit with skill bundle integrity, install/uninstall handling, and an MCP server bridge for app/window management.
- **Context continuity** ([#2145](https://github.com/netease-youdao/LobsterAI/pull/2145)) – Improves Cowork session context compaction reliability after OpenClaw compression, adding session‑scoped task state and lightweight workspace snapshots.
- **NSIS installer fix** ([#2142](https://github.com/netease-youdao/LobsterAI/pull/2142)) – Fixes destructive initialization and redesigns the engine loading page on Windows.
- **Portal fallback URLs** ([#2144](https://github.com/netease-youdao/LobsterAI/pull/2144)) – Updates authentication fallback and upgrade links to new LobsterAI portal domains.
- **Windows in‑app updates** ([#2141](https://github.com/netease-youdao/LobsterAI/pull/2141)) – Resolves update workflow issues on Windows.
- **UI polish** ([#2139](https://github.com/netease-youdao/LobsterAI/pull/2139)) – Refines markdown rendering, code block syntax highlighting, and model selector styling.
- **Task completion notifications** ([#2134](https://github.com/netease-youdao/LobsterAI/pull/2134)) – Restores LobsterAI from notifications when the main window is closed/destroyed.

Additionally, several older PRs (from April) were closed today as part of cleanup (e.g., dependency bumps, stale bugfixes for disabled skills, scheduled task improvements).

### 4. Community Hot Topics
No issues were reported or updated in the last 24 hours. Among pull requests, comment counts are not provided (likely zero), so no single topic generated active discussion. The most significant updates that may attract attention are the **Computer Use MVP** and **context continuity** improvements, both of which signal major upcoming capabilities.

### 5. Bugs & Stability
Several bugs were fixed today, ranked by severity:

| Severity | Bug | Fix PR |
|----------|-----|--------|
| **High** | NSIS installer destructive initialization on Windows (could break existing installs). | [#2142](https://github.com/netease-youdao/LobsterAI/pull/2142) |
| **High** | Windows in‑app update process broken. | [#2141](https://github.com/netease-youdao/LobsterAI/pull/2141) |
| **Medium** | Portal fallback URLs pointing to old domains causing auth failures. | [#2144](https://github.com/netease-youdao/LobsterAI/pull/2144) |
| **Medium** | Task completion notifications not restoring the main window on macOS. | [#2134](https://github.com/netease-youdao/LobsterAI/pull/2134) |
| **Low** | Disabled skills still being injected into system prompts (fix from April, closed today) | [#1485](https://github.com/netease-youdao/LobsterAI/pull/1485) |

All reported bugs appear to have fix PRs merged today or earlier.

### 6. Feature Requests & Roadmap Signals
New features landed today point to near‑term roadmap priorities:

- **Computer Use** – the MVP adds a foundation for Windows app control via natural language. Expect further refinements and cross‑platform support in upcoming releases.
- **Context management** – the continuity layer (`#2145`) and session pruning (earlier PR [#1499](https://github.com/netease-youdao/LobsterAI/pull/1499)) indicate a focus on longer, more reliable agent sessions.
- **Task notifications** – full support for restoring app state from notifications, plus scheduled task enhancements (macOS notifications, test task button) suggest ongoing investment in autonomous agents.
- **Settings UX** – rich markdown editors for agent identity files (PR [#1503](https://github.com/netease-youdao/LobsterAI/pull/1503)) and UI polish (`#2139`) show attention to user configuration experience.

The next minor release will likely include Computer Use stability updates and deeper context compaction tuning.

### 7. User Feedback Summary
No direct user complaints were filed in the last 24 hours, but the fixes merged today reveal recurring pain points:

- **Windows installer reliability** – multiple users encountered NSIS destructive init failures.
- **In‑app updates** – Windows users could not update via the app UI.
- **Task notification behavior** – on macOS, clicking notifications did not bring the app to focus; also scheduled tasks had inconsistent channel delivery.
- **Skill management** – disabled skills remained active, confusing users.
- **Context limits** – long‑running Cowork sessions hit input‑too‑long errors until session pruning was introduced (PR [#1499](https://github.com/netease-youdao/LobsterAI/pull/1499)).

These fixes indicate that the team is responsive to real user issues.

### 8. Backlog Watch
Only **one pull request remains open**:

- [#1277](https://github.com/netease-youdao/LobsterAI/pull/1277) – **chore(deps-dev): bump electron group** (dependabot). Open since April 2, 2026, last updated June 10. This dependency update bumps Electron from 40.2.1 to 42.3.3. It has no comments and no maintainer response. While not urgent, it should be reviewed to avoid accumulating version debt.

No open issues exist in the tracker. The backlog is otherwise clean, with many stale PRs being closed today. The project appears well‑managed.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest — 2026-06-11

## 1. Today's Overview
CoPaw is experiencing high development activity, with 16 issues and 49 pull requests updated in the last 24 hours. Two new releases were published (v1.1.11 stable and v1.1.11-beta.3), introducing zero-config OAuth authentication, a Xiaomi MiMo provider, and self-evolving skill creation. The project is actively addressing regressions, with several critical desktop startup bugs reported and fixed within the same day. Overall, the community remains engaged, though a handful of recurring stability issues (scheduled task failures, tool call errors) signal areas needing deeper attention.

## 2. Releases

### v1.1.11 (stable)
- **Free Model OAuth**: Zero-config free models with one-click OAuth authentication (#5049)
- **Xiaomi MiMo Provider**: Built-in support for Xiaomi MiMo Token Plan (#4722)
- No breaking changes reported.
- *Migration note*: Standard upgrade via `pip install qwenpaw==1.1.11`; no config changes required.

### v1.1.11-beta.3
- Self-evolving skill creation flow (#4857)
- Skill improvements and CI pipeline cleanup (#5056)
- Targeted at developers experimenting with skill-based agent behavior.

## 3. Project Progress
**30 PRs merged/closed today**, including:
- **Bug fixes**:
  - Shield icon vertical centering in Security settings (#5094, closed)
  - Windows OpenSSL regression: `openssl=3.5.7` pinned to `3.5.6` for desktop builds (#5096, open, fix submitted)
  - Environment variables page scrollbar flickering removed (#4766, closed)
  - `aiohttp` pinned to ≤3.14.0 to avoid SSL error on Windows builds (#5082, closed)
  - Use `certifi` CA bundle for Windows build verification (#5083, closed)
  - Compile-check discord after conda-unpack (#5084, closed)
- **Infrastructure**:
  - Version bumped to 1.1.11.post1 (#5093, closed)
  - Revert of a previous pack fix that broke desktop (#5092, closed)
- **Open but advancing PRs**:
  - **Runtime 2.0** – modular architecture with enhanced tool-call coordination (#5078)
  - **Agent OS Driver** – unified abstraction for MCP/A2A/ACP (#5067)
  - **DataPaw plugin** – data-analysis plugin with 12 BI skills (#4622)
  - **PRD CRUD tool** – built-in with frontend renderer (#4902)
  - **Token usage display** – per-chat token/context usage badge (#4433)

## 4. Community Hot Topics

| Issue/PR | Comments | Topic |
|----------|----------|-------|
| [#4878](https://github.com/agentscope-ai/QwenPaw/issues/4878) *closed* | 7 | WeChat delivery failure for scheduled tasks – root-caused to `to_user_id` parameter mismatch in channel logic |
| [#4989](https://github.com/agentscope-ai/QwenPaw/issues/4989) *closed* | 5 | Local Qwen 3.6-27B model unresponsive in v1.1.9/1.1.10 (works in 1.1.5) |
| [#4992](https://github.com/agentscope-ai/QwenPaw/issues/4992) *open* | 4 | Request for independent visual model fallback when main model lacks vision |
| [#5064](https://github.com/agentscope-ai/QwenPaw/issues/5064) *open* | 5 | Agent-created scheduled tasks fail to trigger and cannot be edited |
| [#5095](https://github.com/agentscope-ai/QwenPaw/issues/5095) *closed* | 3 | Windows desktop client v1.1.11 cannot start |
| [#5086](https://github.com/agentscope-ai/QwenPaw/issues/5086) *open* | 3 | OpenSSL 3.5 regression causing desktop startup failure (fix in progress) |

**Key insight**: A recurring theme is **delivery reliability** – scheduled tasks (WeChat, agent-created) fail silently or cannot be triggered. The community wants better error visibility and fallback strategies.

## 5. Bugs & Stability

### Critical (active fix)
- **Windows Desktop cannot start** (v1.1.11) – two related reports:
  - [#5095](https://github.com/agentscope-ai/QwenPaw/issues/5095): Installer creates non-launchable app.
  - [#5086](https://github.com/agentscope-ai/QwenPaw/issues/5086): Root cause – OpenSSL 3.5.7 regression on DER certificate loading.
  - *Fix PRs*: [#5096](https://github.com/agentscope-ai/QwenPaw/pull/5096) (pin OpenSSL to 3.5.6); [#5093](https://github.com/agentscope-ai/QwenPaw/pull/5093) (post-release patch).
- **Security bypass** ([#5090](https://github.com/agentscope-ai/QwenPaw/issues/5090)): `rm` command intercepted, but agent used Python to delete files. No fix PR yet.

### High
- **Agent-created scheduled tasks fail** ([#5064](https://github.com/agentscope-ai/QwenPaw/issues/5064)) – tasks created via agent never trigger and cannot be edited manually.
- **Local model no response** ([#4989](https://github.com/agentscope-ai/QwenPaw/issues/4989)) – regression between 1.1.5 and 1.1.9/1.1.10 for self-hosted Qwen models.
- **Tool call 'unexpected keyword argument'** ([#5052](https://github.com/agentscope-ai/QwenPaw/issues/5052)) – after several successful calls, all tools fail with argument error.
- **Session return failure** ([#5089](https://github.com/agentscope-ai/QwenPaw/issues/5089)) – after running a new session, previous session cannot be resumed.

### Medium
- WeChat text delivery rejected (`ret=-3`) – closed but root cause identified as channel handler mismatch ([#4878](https://github.com/agentscope-ai/QwenPaw/issues/4878)).
- Tool call content not streamed in UI during file generation ([#4865](https://github.com/agentscope-ai/QwenPaw/issues/4865)).
- UI lag with 4+ tabs on Windows Tauri client ([#5053](https://github.com/agentscope-ai/QwenPaw/issues/5053)).

## 6. Feature Requests & Roadmap Signals

### Likely for next release
- **Visual model fallback** ([#4992](https://github.com/agentscope-ai/QwenPaw/issues/4992)) – high demand, aligns with multi-model strategy.
- **Agent avatars** ([#4974](https://github.com/agentscope-ai/QwenPaw/issues/4974)) – UI enhancement, straightforward to implement.
- **DingTalk private deployment custom endpoint** ([#4887](https://github.com/agentscope-ai/QwenPaw/issues/4887)) – enterprise request, low complexity.
- **Token usage display** (PR [#4433](https://github.com/agentscope-ai/QwenPaw/pull/4433)) – under review for a while, may merge soon.

### Strategic signals
- **Runtime 2.0** (PR [#5078](https://github.com/agentscope-ai/QwenPaw/pull/5078)) – monolithic runner replaced with modular coordinator; major refactor.
- **Agent OS Driver** (PR [#5067](https://github.com/agentscope-ai/QwenPaw/pull/5067)) – unified interface for MCP/A2A/ACP – indicates expansion to multi-agent ecosystems.
- **Governance & sandbox** (PR [#5088](https://github.com/agentscope-ai/QwenPaw/pull/5088)) – early discussion on security isolation.

## 7. User Feedback Summary

- **Satisfaction**: Positive reaction to v1.1.11 new providers (OAuth, Xiaomi) and self-evolving skills. Community appreciates rapid bug-fix turnaround (e.g., OpenSSL regression fixed within hours).
- **Pain points**:
  - Desktop startup failures on Windows (multiple reports).
  - Scheduled tasks (agent-created or via `home_agent`) are unreliable – tasks disappear, don't trigger, or fail to deliver to WeChat.
  - Local model compatibility is fragile across versions – downgrading to 1.1.5 restores functionality.
  - Long-running file generation freezes the UI because tool arguments are not streamed.
  - Security controls are easily bypassed (rm → Python workaround).
- **Use cases**: Users heavily rely on scheduled tasks for monitoring/notification, and local model deployment for privacy. The inability to edit agent-created tasks is a major workflow blocker.

## 8. Backlog Watch

| Item | Age | Status | Why Important |
|------|-----|--------|---------------|
| [#4433](https://github.com/agentscope-ai/QwenPaw/pull/4433) – Token usage display | Since May 15 | Under review | High user desire, already has implementation |
| [#4622](https://github.com/agentscope-ai/QwenPaw/pull/4622) – DataPaw plugin | Since May 22 | Open, first-time contributor | 12 BI skills; needs maintainer review |
| [#4902](https://github.com/agentscope-ai/QwenPaw/pull/4902) – PRD CRUD built-in tool | Since Jun 2 | Open | Documentation/management use case |
| [#5051](https://github.com/agentscope-ai/QwenPaw/pull/5051) – Persist backend port across restarts | Since Jun 9 | Open | Fixes agent selection reset on desktop restart (#4733) |

**Maintainer attention needed**: The DataPaw plugin and token usage PRs have been open for 3+ weeks without maintainer comments. The Runtime 2.0 PR (#5078) and Agent OS Driver (#5067) are large changes that require thorough review. The security bypass issue (#5090) has no fix PR yet and may need architectural remediation (e.g., sandboxing).

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-06-11

## 1. Today's Overview
The ZeroClaw project remains **highly active** with 50 pull requests updated in the last 24 hours (39 open, 11 merged/closed) and 7 issues touched (6 still open). Activity is split between **bug squashing** (three P1/P2 bugs opened today), **feature work** (alias renaming, editor improvements, cross-platform CI), and **documentation/operational fixes** (Discord invite, rootless Docker, MCP clarity). No new releases were published. The project health is good but some areas — particularly Windows test coverage, delegate agentic mode, and missing tools in containers — are attracting high-priority attention.

## 2. Releases
**None** – no tags or releases were cut in the reporting period.

## 3. Project Progress
The following **11 PRs were merged or closed** today (note: the data only shows a subset; here are those explicitly listed as closed):

- **#7096** – `docs(readme): use stable Discord vanity invite` – Fixes the invalid Discord link in the README (replacing expired invite with `https://discord.gg/zeroclaw`).
- **#7458** – `revert(ci): remove cross-platform clippy lint gate` – Rolls back a previous PR that added required macOS/Windows Clippy checks; returns CI to Linux-only for the required lint job to avoid blocking routine PRs.

Additionally, many **open PRs** are progressing toward merge (e.g., editor fallback from `vi` to `nano`, Cmd-C keymap fix, theme markdown body, locale‑aware string routing). The **high volume of documentation PRs** (#7473, #7474, #7475, #7479) suggests a push to improve user onboarding and API clarity.

## 4. Community Hot Topics
Only two issues received comments (1 each); all shown PRs have no comment count listed. Nonetheless, the following generated the most recent discussion:

- **#7468** [OPEN] `[Feature]: Allow aliases to be renamed` – 1 comment, 0 reactions. User `damajor` requests the ability to rename aliases for model providers and agents in the TUI. The TUI currently requires re‑creation if a typo is made. Underlying need: **improving configuration ergonomics**.

- **#7467** [OPEN] `[Feature]: More flexibility in edit strings` – 1 comment. Also by `damajor`, asks for arrow‑key navigation when editing strings (e.g., to fix typos without retyping everything). Together with #7468 this signals a user desire for **better interactive configuration editing**.

- **#7037** [CLOSED] `[Bug]: Discord invite link in README is invalid` – 1 comment, now fixed via PR #7096. Community trust issue resolved.

- **#7470** [OPEN] `[Bug]: delegate agentic mode rejects empty risk_profile.allowed_tools` – 0 comments but high‑priority (P1) and authored by `vrurg`. The silent failure blocks practical multi‑agent reviewer/re‑search setups. Expect discussion soon.

*No PRs had comment counts in the provided data.*

## 5. Bugs & Stability
Three new bugs reported today, ranked by severity:

| Issue | Severity | Component | Summary | Fix PR Exists? |
|-------|----------|-----------|---------|----------------|
| **#7470** | **S1 – Workflow blocked** | runtime/daemon | Delegate agentic mode rejects empty `allowed_tools`; same‑profile gating blocks stricter targets. | None yet |
| **#7462** | **S2 – Degraded** | tooling/ci | 74 test failures on Windows (Unix‑only commands, path encoding). CI does not catch because only Linux tested. | Enhancement #7461 proposes adding Windows/macOS to CI. |
| **#7469** | **S3 – Minor** | config / container | The TUI defaults to `vi` as editor, but container images lack `vi` (have `nano`). | Yes – #7483 (change fallback to `nano`) and #7476 (environment variable fallback). |

Additionally, **#7037** (invalid Discord invite, S3) was fixed and closed.

## 6. Feature Requests & Roadmap Signals
Today’s feature requests cluster around three themes:

- **TUI configuration workflow** – #7468 (rename aliases) and #7467 (arrow‑key navigation for string editing). These are likely to land in the next minor release as they improve the `zerocode` CLI experience with low risk.
- **Cross‑platform testing** – #7461 (run tests on Windows/macOS in CI). Given the 74 Windows failures, this is a strong candidate for a near‑term CI overhaul.
- **Observability & turn metadata** – PR #7385 (add turn_id to observer events) is a large (size:XL) enhancement that extends OTel span correlation. This may be aimed at a v0.6+ release.

Other in‑flight mentions: `cron pause/resume` (#7384), model list improvements in `doctor` (#7450), and tool localization (#7480). The project appears to be gradually polishing the user‑facing CLI and gateway experience.

## 7. User Feedback Summary
Real user pain points surfaced in the last 24 hours:

- **Editor fallback** – Container users (`damajor`) found that the text editor launch fails with “command not found” because `vi` is not packaged in Debian/Alpine images. The expected default would be `nano`.
- **MacOS key mapping** – User issue #7378 (linked from #7484, #7477) reports that `Cmd‑C` triggers the Quit chord instead of copying text. A fix is in PR review.
- **Configuration rigidity** – `damajor`’s two feature requests (alias rename, string editing) indicate that setting up model providers and agent names is currently cumbersome and error‑prone.
- **Delegate agent limits** – `vrurg` reports a workflow‑blocking bug when trying to use delegate agents with restricted tool sets; the security gating is too strict.

Overall, satisfaction seems high for core agent capabilities, but **configuration and cross‑platform support** are areas of dissatisfaction.

## 8. Backlog Watch
No issues or PRs were flagged as long‑unanswered in the provided data. However, the following items may warrant maintainer attention:

- **#7469** (vi fallback) – Two overlapping fix PRs (#7483 and #7476) exist; maintainers should pick one or merge them to avoid duplication.
- **#7462** (Windows test failures) – No fix PR yet, but its companion enhancement #7461 is open. Without a PR that actually fixes the 74 failures, the project remains blind to Windows regressions.
- **PR #7215** (webhook port field in quickstart) – Open since June 4, needs‑author‑action; if the author is unresponsive, maintainers may need to take over or close.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*