# OpenClaw Ecosystem Digest 2026-06-24

> Issues: 36 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-24 10:35 UTC

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

# OpenClaw Project Digest — 2026-06-24

## 1. Today's Overview
Project activity remains extremely high, with **36 issues updated** (22 open, 14 closed) and **500 PRs updated** (454 open, 46 merged/closed) in the last 24 hours. A new release **v2026.6.10** shipped today, introducing automatic fast mode for conversational turns and more reliable model routing. However, the release also introduced several regressions (e.g., Anthropic Vertex routing change, memory store relocation in v2026.6.9, and Groq transcription breakage). The community is actively reporting bugs and contributing fixes—especially around security boundaries, session state, and message delivery.

## 2. Releases
- **v2026.6.10** (openclaw 2026.6.10)
  - **Highlights:**
    - **Automatic fast mode for talks** – OpenClaw now detects short conversational turns and switches to fast mode, falling back to normal mode for longer runs with bounded delivery behavior. ([#85104](https://github.com/openclaw/openclaw/issues/85104)) Thanks @alexph-dev and @vincentkoc.
    - **More reliable model routing** – Zai model synthesis improvements (details not fully shared in highlights).
  - **⚠️ Regressions reported in this release:**
    - `anthropic-vertex` provider `route=native` (changed from `proxy-like` in v2026.6.8) causes non-visible output for pure text responses ([#96337](https://github.com/openclaw/openclaw/issues/96337)).
    - `lossless-claw` context engine fails with *"Engine initialization is disabled during read-only plugin registration"* ([#96335](https://github.com/openclaw/openclaw/issues/96335)).
  - **Migration notes:** Users upgrading from v2026.6.8 or earlier should test `anthropic-vertex` and lossless-claw configurations carefully. No official migration guide yet, but a fix PR is in progress for the lossless-claw issue.

## 3. Project Progress
**46 PRs were merged or closed today.** Notable changes:

- **New features & infrastructure:**
  - `feat(copilot): add BYOK provider parity` ([#96345](https://github.com/openclaw/openclaw/pull/96345)) – Merged; brings Copilot harness up to date with custom provider routing and runtime-session contract.
  - `Autofix: add PR review autofix pipeline + Windows daemon` ([#68936](https://github.com/openclaw/openclaw/pull/68936)) – Merged; automates review comment addressing using Claude Agent SDK.
  - `feat(discord): add canvas-first Discord Activities support` ([#65205](https://github.com/openclaw/openclaw/pull/65205)) – Still open, but significant addition for rich interactive experiences.

- **Bug fixes merged:**
  - `fix(model): /model <default> writes override when session runs non-default model` ([#96368](https://github.com/openclaw/openclaw/pull/96368))
  - `fix: fail closed when configure is run non-interactively` ([#93998](https://github.com/openclaw/openclaw/pull/93998))
  - `fix(tts): defer text settlement for final-mode TTS to eliminate churn` ([#83988](https://github.com/openclaw/openclaw/pull/83988)) – Fixes #83511.
  - `fix(devices): always route stale-approve operator to openclaw devices list` ([#80779](https://github.com/openclaw/openclaw/pull/80779))
  - `fix(config): preserve bootstrapMaxChars/bootstrapTotalMaxChars across…` ([#96294](https://github.com/openclaw/openclaw/pull/96294)) – Fixes #96240 regression.
  - `fix(secrets): inherit full parent env when exec provider passEnv is not explicitly configured` ([#94003](https://github.com/openclaw/openclaw/pull/94003)) – Fixes #93851.
  - `fix(firecrawl): bound successful JSON response reads` ([#96043](https://github.com/openclaw/openclaw/pull/96043))
  - `fix(tavily): bound untrusted JSON response reads` ([#96026](https://github.com/openclaw/openclaw/pull/96026))
  - `fix(nextcloud-talk): bound external send/reaction response reads` ([#96031](https://github.com/openclaw/openclaw/pull/96031))
  - Several fixes for Telegram, QQ bot, and other channels.

## 4. Community Hot Topics
The most active conversations (by comments and reactions) reveal deep user concerns around **security, session state, and data integrity**:

- **[#85030](https://github.com/openclaw/openclaw/issues/85030) – "MCP tools not injected into subagent sessions"**  
  9 comments, 3 👍. Users report that `mcp.servers` tool schemas are completely ignored when agents spawn sub-sessions, even with allowlists configured. This is a **P1 security/behavior bug** with a diamond lobster rating. Community members are actively reproducing and requesting a maintainer decision.

- **[#7722](https://github.com/openclaw/openclaw/issues/7722) – "Filesystem Sandboxing Config (tools.fileAccess)"**  
  9 comments, 4 👍. Long-standing feature request (since February) to restrict filesystem access via configuration. Users demonstrate that the current implementation does not honour `denyPaths` and `allowedPaths` properly. The need for a robust sandbox is critical for multi-tenant or security-conscious deployments.

- **[#95495](https://github.com/openclaw/openclaw/issues/95495) – "Silent memory store relocation with no migration"**  
  7 comments, 1 👍. This closed regression (from v2026.6.9) forced a full re-embed of 1499 files without warning. Users expressed strong dissatisfaction with the lack of upgrade-time warning. The fix has likely been included in v2026.6.10, but it raised trust concerns.

- **[#39847](https://github.com/openclaw/openclaw/issues/39847) – "Echo contamination: stripInboundMetadata not called"**  
  6 comments, 1 👍. Open since March, this P1 security bug causes internal metadata (thread context, sender info) to leak to Discord. Community is frustrated that it remains unaddressed despite being marked `needs-maintainer-review`.

- **[#84084](https://github.com/openclaw/openclaw/issues/84084) – "Codex legacy mirrored-history fallback ignores contextTokenBudget"**  
  5 comments, 1 👍. High-window sessions are capped at ~24k rendered chars due to a hardcoded fallback. Users running long conversations face degraded context quality.

## 5. Bugs & Stability
**Today's bug reports (newly opened or updated) ranked by severity:**

| Severity | Issue | Summary | Fix PR Exists? |
|----------|-------|---------|----------------|
| **P1** (Critical) | [#96337](https://github.com/openclaw/openclaw/issues/96337) | `anthropic-vertex` regression: text responses invisible with `route=native` | No (open) |
| **P1** | [#96335](https://github.com/openclaw/openclaw/issues/96335) | `lossless-claw` engine fails with read-only plugin registration error | No (open) |
| **P1** | [#96331](https://github.com/openclaw/openclaw/issues/96331) | Web UI session context not preserved after restart – model can't see history | No (open) |
| **P1** | [#96242](https://github.com/openclaw/openclaw/issues/96242) | Multiple independent paths cause duplicate Telegram messages | No (open) |
| **P1** | [#96374](https://github.com/openclaw/openclaw/issues/96374) | Stuck "claimed" rows in `channel_ingress_events` block inbound delivery after restart | No (open) |
| **P1** | [#96373](https://github.com/openclaw/openclaw/issues/96373) | ChatGPT OAuth login fails with "unsupported" error | No (open) |
| **P1** | [#96346](https://github.com/openclaw/openclaw/issues/96346) | Async/cron command results lose auth prompt lines due to truncation | Yes [#96370](https://github.com/openclaw/openclaw/pull/96370) |
| **P1** | [#96363](https://github.com/openclaw/openclaw/issues/96363) | Telegram `/status` fails with `RICH_MESSAGE_EMAIL_INVALID` when rich messages enabled | No (open) |
| **P2** | [#96334](https://github.com/openclaw/openclaw/issues/96334) | `lightContext` causes missing tools/rules in isolated sessions | No (open) |
| **P2** | [#96350](https://github.com/openclaw/openclaw/issues/96350) | Deeplink/contentIntent ignored in Android system.notify | No (open) |
| **P2** (Closed) | [#96297](https://github.com/openclaw/openclaw/issues/96297) | Final reply suppressed when tool progress shows 'failed' (exec exit code != 0) | Yes [#96371](https://github.com/openclaw/openclaw/pull/96371) |
| **P2** (Closed) | [#96355](https://github.com/openclaw/openclaw/issues/96355) | `memory_search` fails on subsequent queries after cold start (index metadata missing) | Yes (fix merged) |
| **P2** (Closed) | [#96328](https://github.com/openclaw/openclaw/issues/96328) | OpenAI-compatible embeddings: CLI and runtime compute different providerKey → permanent "index metadata is missing" | Yes (merged) |
| **P2** (Closed) | [#96313](https://github.com/openclaw/openclaw/issues/96313) | Feishu renderMode setting not respected for sub-accounts | Yes (merged) |
| **P2** (Closed) | [#96306](https://github.com/openclaw/openclaw/issues/96306) | Telegram streaming: agent progress edits overwrite previous steps, user loses context | No (closed as duplicate?) |
| **P2** (Closed) | [#96330](https://github.com/openclaw/openclaw/issues/96330) | WebChat doesn't support Skill Workshop approval prompts | No (closed as duplicate) |

Several additional regressions from v2026.6.9 remain: **#95495** (memory store relocation) and **#95658** (Groq transcription broken). Fixes may be coming in the next patch.

**Security bugs** still unpatched: #39847 (echo contamination, P1, open since March), #7722 (filesystem sandboxing, P2, open since Feb), #20935 (audit log missing, P2, enhancement but security impact).

## 6. Feature Requests & Roadmap Signals
The following enhancements were updated today and are strong candidates for the next minor release:

- **Token-based context limits** ([#94657](https://github.com/openclaw/openclaw/issues/94657)) – Replace character-based `bootstrapMaxChars` with token‑based limits. High demand for accurate context window management.
- **Discord progress draft preservation** ([#94285](https://github.com/openclaw/openclaw/issues/94285)) – Option to append final answer instead of replacing the live progress draft. Would improve UX for multi-step tasks.
- **Android deeplink support for `system.notify`** ([#96350](https://github.com/openclaw/openclaw/issues/96350)) – Tapping notifications should open the agent app. Small but impactful for mobile users.
- **Audit log for memory changes** ([#20935](https://github.com/openclaw/openclaw/issues/20935)) – Open since February, this request for append-only memory audit logs could gain traction given recent concerns about silent data changes.

Given the volume of P1 bugs, the next release (v2026.6.11) is likely to be a stability patch. The token-based limits feature (#94657) has been tagged `queueable-fix` and might land sooner.

## 7. User Feedback Summary
Real user voices from today's issues and PRs:

- **"Upgrade broke my setup"** – Multiple users report silent breakage when upgrading from v2026.6.8 or v2026.6.9: memory store relocation without warning ([#95495](https://github.com/openclaw/openclaw/issues/95495)), Groq transcription silently stops working ([#95658](https://github.com/openclaw/openclaw/issues/95658)), and `lossless-claw` engine fails on v2026.6.10 ([#96335](https://github.com/openclaw/openclaw/issues/96335)).
- **"MCP tools are ignored in subagents"** – A recurring pain point for users building multi-agent workflows. The documented allowlist mechanism is completely ineffective ([#85030](https://github.com/openclaw/openclaw/issues/85030)).
- **"Skill Workshop approval times out in code, works in CLI"** – Users trying to automate skill application hit a 90-second timeout in agent-tool mode but not via CLI ([#91266](https://github.com/openclaw/openclaw/issues/91266)). WebChat also doesn't support approval prompts ([#96330](https://github.com/openclaw/openclaw/issues/96330)).
- **"Telegram messages lost or duplicated"** – Users report messages are either not delivered (Telegram `/status` silent, [#95525](https://github.com/openclaw/openclaw/issues/95525)) or delivered twice ([#96242](https://github.com/openclaw/openclaw/issues/96242)). Streaming progress updates overwrite previous steps, making multi-step workflows opaque ([#96306](https://github.com/openclaw/openclaw/issues/96306)).
- **"Filesystem sandboxing doesn't work as documented"** – Users trying to restrict agent file access find that `denyPaths` and `allowedPaths` are not enforced ([#7722](https://github.com/openclaw/openclaw/issues/7722)). This is a security requirement for production deployments.
- **"Memory search fails after cold start"** – A regression in v2026.6.9 caused `memory_search` to fail on subsequent queries; fixed in today's merge ([#96355](https://github.com/openclaw/openclaw/issues/96355)), but incident eroded confidence.

Overall, users are **satisfied with the pace of development** (500 PRs touched today) but **frustrated by regressions in stable releases and the slow resolution of P1 security bugs**.

## 8. Backlog Watch
Issues and PRs that have been open for extended periods without resolution, needing maintainer attention:

| Item | Created | Comments | Status | Reason for concern |
|------|---------|----------|--------|-------------------|
| [#39847](https://github.com/openclaw/openclaw/issues/39847) – Echo contamination (P1 security) | 2026-03-08 | 6 | Open, `needs-maintainer-review` | 3.5 months; security data leak to Discord |
| [#7722](https://github.com/openclaw/openclaw/issues/7722) – Filesystem sandboxing | 2026-02-03 | 9 | Open, `needs-maintainer-review` + `needs-product-decision` | Nearly 5 months; high community demand (4 👍) |
| [#20935](https://github.com/openclaw/openclaw/issues/20935) – Audit log for memory | 2026-02-19 | 4 | Open, `needs-maintainer-review` + `needs-product-decision` | 4 months; security and transparency gap |
| [#78493](https://github.com/openclaw/openclaw/issues/78493) – `sudo openclaw update` causes mixed ownership | 2026-05-06 | 5 | Open, `needs-maintainer-review` | 1.5 months; can break doctor/config |
| [#46794](https://github.com/openclaw/openclaw/pull/46794) – Device pairing bind setup codes to node approvals | 2026-03-15 | 0 | Open, `waiting on author` | 3 months; security improvement blocked |
| [#65205](https://github.com/openclaw/openclaw/pull/65205) – Discord Activities support | 2026-04-12 | 0 | Open, `waiting on author` | 2.5 months; large new feature stalled |
| [#78247](https://github.com/openclaw/openclaw/pull/78247) – Include token usage in model usage summaries | 2026-05-06 | 0 | Open, `waiting on author` | 1.5 months; useful for cost tracking |
| [#84906](https://github.com/openclaw/openclaw/pull/84906) – Gradium realtime transcription provider | 2026-05-21 | 0 | Open, `waiting on author` | 1 month; new provider integration stalled |

**Call to action**: The project would benefit from a **maintainer review blitz** on the oldest P1 issues (#39847, #7722) to restore community trust. Several PRs are waiting on author updates—maintainers may consider reaching out or reassigning.

---

## Cross-Ecosystem Comparison

# AI Agent Open-Source Ecosystem: Cross-Project Comparison Report

**Date:** 2026-06-24  
**Analyst:** Senior Ecosystem Analyst  
**Scope:** 12 active projects in the personal AI assistant / agent open-source landscape

---

## 1. Ecosystem Overview

The personal AI assistant open-source ecosystem is experiencing a maturation phase characterized by rapid iteration on core infrastructure—session state management, model routing, and security boundaries—while simultaneously expanding into new integration surfaces (voice, WebRTC, Slack Socket Mode, Telegram rich messages). The landscape bifurcates between "general-purpose agent frameworks" (OpenClaw, Hermes, NanoClaw, ZeroClaw) and "specialized agents" (CoPaw for Chinese users, NullClaw for Windows users, LobsterAI for enterprise). A clear tension has emerged between feature velocity and stability: several projects shipped regressions in recent releases (OpenClaw v2026.6.10, LobsterAI's token-burn bug) even as they maintain high development throughput. Community feedback converges on three pain points: security isolation for sub-agents, token efficiency for local models, and cross-platform reliability. No single project has achieved market dominance; the ecosystem remains fragmented, with each project differentiating on provider support, UI paradigm, or deployment model.

---

## 2. Activity Comparison

| Project | Issues Updated (Open/Closed) | PRs Updated (Open/Merged) | Release Today? | Health Score |
|---------|------------------------------|---------------------------|----------------|--------------|
| **OpenClaw** | 36 (22 open, 14 closed) | 500 (454 open, 46 merged) | ✅ v2026.6.10 | 8/10 – Very high velocity, but regressions |
| **NanoBot** | 16 (7 open, 9 closed) | 44 (23 open, 21 merged) | ❌ | 7/10 – Responsive, security gaps remain |
| **Hermes Agent** | 18 (18 open, 0 closed) | 50 (49 open, 1 merged) | ❌ | 6/10 – High activity but few closures |
| **PicoClaw** | 2 (1 open, 1 closed) | 17 (12 open, 5 merged) | ❌ | 6/10 – Focused fixes, Android crash risk |
| **NanoClaw** | 1 (1 open, 0 closed) | 12 (4 open, 8 merged) | ❌ | 7/10 – Efficient merging, security issue |
| **NullClaw** | 1 (0 open, 1 closed) | 1 (1 open, 0 merged) | ❌ | 3/10 – Stalled, single bug closed |
| **IronClaw** | 14 (9 open, 5 closed) | 48 (20 open, 28 merged) | ❌ | 8/10 – Strong closure rate, critical fix |
| **LobsterAI** | 1 (1 open, 0 closed) | 45 (3 open, 42 merged) | ❌ | 9/10 – Exceptionally high merge rate |
| **TinyClaw** | 0 | 0 | ❌ | 1/10 – Inactive |
| **Moltis** | 0 | 0 | ❌ | 1/10 – Inactive |
| **CoPaw** | 16 (11 open, 5 closed) | 50 (40 open, 10 merged) | ❌ | 6/10 – Active, frontend stability issues |
| **ZeptoClaw** | 0 | 0 | ❌ | 1/10 – Inactive |
| **ZeroClaw** | 13 (12 open, 1 closed) | 50 (41 open, 9 merged) | ❌ | 7/10 – Security-focused, prepping release |

*Health Score: Composite of issue/PR closure ratio, bug severity, community engagement, and release cadence (1 = stalled, 10 = optimal)*

---

## 3. OpenClaw's Position

**Advantages vs. Peers:**

- **Community scale:** OpenClaw's 500 PRs touched in 24 hours dwarfs the next most active projects (Hermes, CoPaw, ZeroClaw at ~50 each). It has roughly 10× the development throughput of any single peer.
- **Release maturity:** v2026.6.10 demonstrates a structured release process with versioned migration notes, even if imperfect. Most peers (Hermes, NanoBot, CoPaw) lack recent formal releases.
- **Provider ecosystem:** OpenClaw's `anthropic-vertex`, `lossless-claw`, and Groq integrations signal deeper provider support than peers. NanoBot and Hermes focus on OpenAI-compatible APIs with fewer exotic providers.
- **Feature breadth:** Automatic fast mode for conversational turns (#85104) and BYOK provider parity (#96345) show investment in UX and enterprise readiness that peers (NullClaw, PicoClaw) lack.

**Technical Approach Differences:**

- OpenClaw uses a **session-based model** with channel ingress/egress and rich context (bootstrap, memory store). Hermes and ZeroClaw emphasize **sub-agent delegation** and tool orchestration. NanoBot and CoPaw prioritize **WebUI-first experiences** on top of OpenClaw-style backends.
- OpenClaw's architecture is "batteries included" with built-in TTS, device management, and skill workshop. Peers like NullClaw and PicoClaw are more minimal, targeting headless or single-platform deployments.
- OpenClaw's regression-heavy release pattern (3 P1 regressions in v2026.6.10 alone) suggests higher risk tolerance than IronClaw or LobsterAI, which merge more conservatively.

**Community Size Comparison (Inferred):**

- OpenClaw clearly has the largest contributor base (unrelated authors across 500 PRs). IronClaw, ZeroClaw, and Hermes have active but smaller contributor pools (multiple named individuals per digest). NanoBot and CoPaw show strong first-time contributor engagement. NullClaw, TinyClaw, Moltis, and ZeptoClaw appear to have minimal or no community participation.

---

## 4. Shared Technical Focus Areas

*Requirements emerging across multiple projects:*

| Focus Area | Affected Projects | Specific Need |
|------------|------------------|---------------|
| **Security isolation** | **OpenClaw** (#85030, #39847), **NanoBot** (#4434, #4435), **ZeroClaw** (#8279), **PicoClaw** (#3161, #3160) | Sub-agent tool allowlists bypassed; MCP resource exposure; filesystem sandboxing not enforced. This is the #1 cross-project pain point. |
| **Token efficiency** | **OpenClaw** (#84084, #94657), **Hermes** (#6839, #4379), **ZeroClaw** (context budget) | 73% fixed token overhead per call; character-based limits force hard caps; lazy tool schema loading requested. Local model users are most vocal. |
| **Session state reliability** | **OpenClaw** (#96331), **LobsterAI** (#2047, #2050), **CoPaw** (#5401, #5479) | Context not preserved across restarts; session freezing; large sessions crash frontend. |
| **Cross-platform parity** | **OpenClaw** (QQ/Telegram regressions), **PicoClaw** (#3164 Android crash), **NullClaw** (#967 Windows error), **NanoBot** (Telegram Web broken) | Platform-specific bugs erode trust; Windows, Android/Termux, and Telegram Web are common failure points. |
| **Observability & diagnostics** | **IronClaw** (#5182), **ZeroClaw** (#8065), **OpenClaw** (audit log #20935) | Demand for structured logs, cost tracking, and failure diagnostics in hosted deployments. |
| **Plugin/skill discoverability** | **ZeroClaw** (#6289), **OpenClaw** (Skill Workshop), **CoPaw** (#5484 pip plugins) | Users want proactive skill suggestions and easier installation—manual discovery is a friction point. |
| **Cron/scheduled automation** | **ZeroClaw** (PR #783), **CoPaw** (#5475, #5483), **OpenClaw** (cron via configure) | Scheduled task reliability and editability are growing needs as headsless agents become common. |

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | Hermes | NanoBot | ZeroClaw | CoPaw | LobsterAI | PicoClaw |
|-----------|----------|--------|---------|----------|-------|-----------|----------|
| **Primary use case** | General-purpose assistant | Multi-agent orchestration | Lightweight WebUI chat | Tool-delegation agent | Chinese-market agent | Enterprise deployment | Embedded/headless |
| **Target user** | Power users & developers | Agent developers | Casual/non-technical | Security-conscious admins | Chinese-speaking users | Corporate teams | IoT/low-resource |
| **Architecture** | Monolithic + plugins | Service-oriented | WebUI-centric | Plugin + WASM | Platform-integrated | Gateway-focused | Minimal core |
| **Provider support** | 15+ providers | OpenAI-compatible + Codex | OpenAI-compatible + Kimi | OpenAI-compatible + WASM | Chinese models + OMLX | MiniMax, Mimo | OpenAI-compatible + Bedrock |
| **UI paradigm** | CLI + WebUI + Discord/telegram | CLI + Desktop GUI | WebUI (PWA) | CLI + WebUI | WebUI + Feishu/DingTalk | Electron app | CLI-only |
| **Key strength** | Feature breadth & community | Sub-agent tools & ACP | Mobile-first (PWA) | Security model & WASM | Chinese ecosystem | Stability & speed | Minimal footprint |
| **Key weakness** | Regression-prone releases | Low PR merge rate | Security bypasses | Orphan bug #5903 | Frontend crashes | Closed-source? | Android crash (#3164) |

**Notable architectural divergence:** ZeroClaw's WASM plugin model is unique among peers—it runs untrusted code in a sandbox via Extism/wasmtime. Hermes' "Valut" secret reference system (#51832) is novel for credential sharing. OpenClaw's "automatic fast mode" (#85104) uses adaptive delivery detection, a UX pattern no other project has attempted.

---

## 6. Community Momentum & Maturity

**Tier 1: Rapidly iterating (high throughput, active maintainers)**
- **OpenClaw** – 500 PRs/day. Releasing frequently but with regression risk. Community engagement is extremely high but frustrated by P1 security bugs (#39847 open 3.5 months).
- **LobsterAI** – 42 closed PRs today. Highest merge efficiency. Focused on stability and IM integrations. Appears to be a corporate-backed project (Netease Youdao).
- **IronClaw** – 28 merged PRs today. Critical regression #5139 fixed same day. Strong balance of features and bug fixes.
- **ZeroClaw** – 50 PRs touched, prepping v0.8.2. Security-focused development. RFC-driven roadmap (supply-chain signing, WASM plugins).

**Tier 2: Actively developing (steady pace, some bottlenecks)**
- **Hermes** – 50 PRs, but only 1 merged. Many open PRs awaiting review. Token efficiency proposals (#6839) have high community support but no maintainer response.
- **NanoBot** – 44 PRs, 21 merged. Responsive to bugs but security bypasses (#4434/4435) unaddressed. Long-running branding complaint (#660) ignored.
- **CoPaw** – 50 PRs, 10 merged. Frontend stability issues (#5401, #5479) are blocking for heavy users. Active first-time contributor onboarding.
- **PicoClaw** – 17 PRs, 5 merged. Android crash (#3164) is critical for mobile users. Small but responsive team.

**Tier 3: Stabilizing or stalled**
- **NanoClaw** – 12 PRs, 8 merged. Efficient but low volume. Security port binding issue (#2840) needs attention.
- **NullClaw** – Near-zero activity. Single bug closed. Stalled PR #783 (2.5 months). Risk of abandonment.
- **TinyClaw, Moltis, ZeptoClaw** – No activity. Effectively dormant.

---

## 7. Trend Signals

**Five industry trends extracted from community feedback across all projects:**

### 1. Security isolation is the #1 unmet need
Sub-agent tool allowlists are universally bypassable (OpenClaw #85030, NanoBot #4434/4435, ZeroClaw #8279). Filesystem sandboxing is not enforced (OpenClaw #7722, 4.5 months open). Credential leakage between parent and child agents persists (ZeroClaw #7623). **Value for developers:** Projects that solve sub-agent security first will win enterprise trust.

### 2. Token efficiency is becoming a competitive moat
73% token overhead per call (Hermes #4379) is not acceptable for local model users. Lazy tool schema loading (#6839), tool output compression (#39691), and token-based context limits (#94657) are cross-project themes. **Value for developers:** The first project to ship default-on token optimization will attract cost-sensitive and local-model users.

### 3. Multi-agent orchestration is the next UX frontier
Hermes' ACP client (#5257), OpenClaw's subagent refactoring, and ZeroClaw's delegate model all point toward agents managing other agents. Users want to orchestrate Claude Code, Codex CLI, and others from a single control plane. **Value for developers:** Multi-agent protocol standardization (similar to MCP) could emerge from this ecosystem.

### 4. Mobile and cross-platform reliability remains fragile
Android crashes (PicoClaw #3164), Telegram Web incompatibility (NanoBot #4488), Windows-specific errors (NullClaw #967), and missing mobile deeplinks (OpenClaw #96350) show that desktop-first development leaves mobile users underserved. PWA support (NanoBot) and responsive UI (CoPaw) are belated responses. **Value for developers:** Mobile-first design is not optional for consumer adoption.

### 5. Observability and auditability are enterprise gatekeepers
IronClaw users demand better logs (#5182). OpenClaw's missing audit log for memory changes (#20935) has been open 4 months. ZeroClaw is adding trace_id correlation and per-call cost tracking (#8065). **Value for developers:** Structured logging, cost attribution, and change audit trails are prerequisites for production deployment, not afterthoughts.

---

**Summary:** The ecosystem is healthy in velocity but fragmented in approach. The next 6 months will likely see a consolidation around security standards, token optimization, and multi-agent protocols. OpenClaw's community scale gives it the best chance of becoming the "reference implementation," but its regression problem and slow security fixes create an opening for ZeroClaw (security-first) or Hermes (orchestration-first) to capture specific niches.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

## NanoBot Project Digest — 2026-06-24

### 1. Today's Overview
NanoBot continues to show very high development velocity with 44 pull requests updated in the last 24 hours (21 merged/closed) and 16 issues touched (9 closed). The project has no new releases today, but the sheer volume of commits and bug fixes signals a healthy, responsive maintainer community. Security concerns around the MCP `enabledTools` mechanism remain open and are attracting attention, while the team is actively shipping user-facing enhancements like PWA support, new LLM providers, and critical bug fixes.

### 2. Releases
No new releases were published today.

### 3. Project Progress
Today's merged/closed PRs advanced several features and fixed multiple regressions:

- **PWA & Mobile UX** — Two PRs for Progressive Web App support were merged: [#4458](https://github.com/HKUDS/nanobot/pull/4458) and [#4480](https://github.com/HKUDS/nanobot/pull/4480) (both by `zpljd258`). The later PR also adds a mobile swipe gesture for the sidebar.
- **New Providers** — `kimi_coding` (Kimi Coding Plan) was added via [#4464](https://github.com/HKUDS/nanobot/pull/4464), and OpenCode Zen/Go providers via [#4476](https://github.com/HKUDS/nanobot/pull/4476).
- **Bug Fixes:**
  - **[#4466](https://github.com/HKUDS/nanobot/pull/4466)** — Fixed raw `<thinking>` tags leaking into WebUI reasoning output.
  - **[#4484](https://github.com/HKUDS/nanobot/pull/4484)** — DuckDuckGo search now respects proxy configuration.
  - **[#4478](https://github.com/HKUDS/nanobot/pull/4478)** — Dream cron settings are no longer silently removed on config save.
  - **[#4469](https://github.com/HKUDS/nanobot/pull/4469)** — Dream now injects existing workspace skills into its prompt to avoid duplicate skill creation.
  - **[#4468](https://github.com/HKUDS/nanobot/pull/4468)** — Heartbeat excludes archived keys; session timestamps fall back to file modification time.
  - **[#4433](https://github.com/HKUDS/nanobot/pull/4433)** — Pairing store now normalizes sender IDs to strings to prevent silent denial.
- **Telegram** — Fix for duplicated `tool_use` IDs in Anthropic/Kimi Coding streams ([#4473](https://github.com/HKUDS/nanobot/pull/4473) closed), and the Telegram rich messages feature was resolved ([#4413](https://github.com/HKUDS/nanobot/pull/4413) closed).

### 4. Community Hot Topics
- **Issue [#660](https://github.com/HKUDS/nanobot/issues/660)** — *“Project claims to be 'ultra-lightweight' but includes bloated Node.js dependency”* (11 comments, 5 👍)  
  This long-standing criticism about the Docker image requiring both Python and Node.js remains the most active discussion. Users feel the project’s tagline is misleading. No maintainer response visible.

- **Issue [#2298](https://github.com/HKUDS/nanobot/issues/2298)** — *“Breaking endless tool calling loops”* (5 comments)  
  A pain point especially with smaller/local models. Users request loop detection logic. No fix has been merged yet.

- **Security Issues [#4434](https://github.com/HKUDS/nanobot/issues/4434) and [#4435](https://github.com/HKUDS/nanobot/issues/4435)** — MCP `enabledTools` policy bypass (1 comment each)  
  Both were filed by `YLChen-007` on June 21 and remain open. They describe how the allowlist and deny-all configurations can be bypassed, exposing resources and prompts. These are high-severity security findings that should be prioritized.

- **Issue [#4488](https://github.com/HKUDS/nanobot/issues/4488)** — *“Telegram Web: 'This message is not supported on the web version of Telegram'”* (0 comments, very new)  
  A regression from the rich messages feature. A fix PR ([#4489](https://github.com/HKUDS/nanobot/pull/4489)) was submitted the same day.

### 5. Bugs & Stability
| Severity | Issue | Description | Fix PR? |
|----------|-------|-------------|---------|
| **High** | [#4442](https://github.com/HKUDS/nanobot/issues/4442) | Duplicate `tool_use` IDs in streamed responses poison the session – 400 errors forever. | Fixed via related PRs [#4473](https://github.com/HKUDS/nanobot/pull/4473) |
| **High** | [#4434/4435](https://github.com/HKUDS/nanobot/issues/4434) | MCP `enabledTools` security bypass – no fix merged yet. | None |
| **Medium** | [#4488](https://github.com/HKUDS/nanobot/issues/4488) | Rich messages break Telegram Web rendering – regression from new feature. | [#4489](https://github.com/HKUDS/nanobot/pull/4489) open |
| **Medium** | [#4470](https://github.com/HKUDS/nanobot/issues/4470) | Telegram display: lost line breaks and message flickering in v0.2.2. | Closed (likely fixed) |
| **Low** | [#4388](https://github.com/HKUDS/nanobot/issues/4388) | iOS Safari auto-zooms on input click – closed. | Fixed in earlier commit |
| **Low** | [#4465](https://github.com/HKUDS/nanobot/issues/4465) | `<thinking>` tags visible in WebUI – fixed by [#4466](https://github.com/HKUDS/nanobot/pull/4466). | Merged |

### 6. Feature Requests & Roadmap Signals
- **PWA Support** — Merged today; users can now install NanoBot WebUI as a mobile app. Expect this in the next release.
- **New LLM Providers** — Kimi Coding Plan and OpenCode Zen/Go were added. These are live contributions from the community and will be available immediately.
- **Telegram Rich Messages** — The feature (Bot API 10.1) was closed but introduced a regression. Future versions may include a toggle (`rich_messages` config) as proposed in [#4489](https://github.com/HKUDS/nanobot/pull/4489).
- **Mattermost Channel** — PR [#4459](https://github.com/HKUDS/nanobot/pull/4459) adds Mattermost support, still open and likely to be reviewed soon.
- **Read-Only Sessions** — PR [#4271](https://github.com/HKUDS/nanobot/pull/4271) proposes a `read_only` session flag. Still open, but useful for static content pages.
- **Per-Channel System Prompts** — PR [#2866](https://github.com/HKUDS/nanobot/pull/2866) (open since April) would allow channel-specific instructions. Lower urgency.
- **API Authentication** — New issue [#4490](https://github.com/HKUDS/nanobot/issues/4490) asks for auth on the OpenAI-compatible API server. Likely to be addressed given the security trend.

### 7. User Feedback Summary
- **Frustration with “ultra-lightweight” branding** — Issue #660 shows users feel misled by the Node.js dependency in Docker. This has been open for over four months without resolution.
- **Local model loops** — Common complaint: smaller models get stuck repeating tool calls. Users want heuristic or configurable loop breaking.
- **Telegram Web incompatibility** — The rich messages feature, intended as an enhancement, broke Telegram Web for some users. Quick fix PR shows the team is reactive but the feature may need more testing.
- **Positive reception for PWA** — Multiple PRs and issues about mobile UX signal strong satisfaction with the WebUI improvements.
- **Security awareness** — The two MCP bypass issues were filed only three days ago but already have no maintainer response. Users may become concerned if these remain unpatched.

### 8. Backlog Watch
- **Issue [#660](https://github.com/HKUDS/nanobot/issues/660)** — *“Ultra-lightweight claim vs Node.js dependency”* (opened Feb 2026, 5 👍). No maintainer comment. If the dependency is unavoidable, a clear documentation update is needed.
- **Issue [#2298](https://github.com/HKUDS/nanobot/issues/2298)** — *Endless tool loops* (opened March 2026, 5 comments). No resolution or assignment.
- **Security Issues [#4434](https://github.com/HKUDS/nanobot/issues/4434) & [#4435](https://github.com/HKUDS/nanobot/issues/4435)** — MCP bypass (opened June 21, 1 comment each). No fix yet. High priority.
- **PR [#2078](https://github.com/HKUDS/nanobot/pull/2078)** — Zalo integration refactor (opened March 16, still open). Needs review.
- **PR [#2283](https://github.com/HKUDS/nanobot/pull/2283)** — Agent evaluation harness (opened March 20). Valuable for CI, stale.
- **PR [#2316](https://github.com/HKUDS/nanobot/pull/2316)** — QQ voice support (opened March 21). Awaiting attention.
- **PR [#2866](https://github.com/HKUDS/nanobot/pull/2866)** — Per-channel system prompts (opened April 6). Long idle.

*(All issue/PR links are to `github.com/HKUDS/nanobot`)*

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

**Hermes Agent Project Digest — 2026-06-24**

---

## Today's Overview

The Hermes Agent repository saw **18 issues** and **50 pull requests** updated in the past 24 hours. All issues remain open; only **one PR was closed/merged** today. Activity remains high across performance, stability, and feature work — especially around token overhead reduction, credential management, and new gateway integrations. No new releases were cut. The project continues to evolve rapidly with strong community engagement (26+ comment threads, multiple high-reaction feature requests).

---

## Releases

No new releases today.

---

## Project Progress

**Merged/Closed PRs (1 total):**

- **#47330** — Real-time voice conversation platform (Daily + Deepgram Flux + Cartesia) – *superseded by #51827* (same author). The original PR was closed in favor of a clean rebased branch.

**Key open PRs advancing features:**

- **#51768** — Enable `/reasoning`, `/fast`, `/voice`, `/skills` slash commands in Desktop GUI.
- **#51226** — Context engine hooks `select_context()` + `on_turn_complete()` (major architectural RFC implementation).
- **#51827** — Real-time WebRTC voice as an in-process gateway platform plugin.
- **#51832** — "Valut" secret reference system for safe credential sharing in chat.
- **#51795** — Subagent presets skill + web dashboard agents page.
- **#51639** — Terminal billing: `/subscription` view and Remote-Spending gate UX.
- **#48184** — OTLP observability plugin (OpenTelemetry export).
- **#51830** — macOS hotkey help text correction in TUI.
- **#51835** — Auxiliary compression fallback chain when provider auth is unavailable.

---

## Community Hot Topics

The most active discussions and most-upvoted items this cycle:

- **#6839** (🔥 26 comments, 14 👍) — *Feature: Lazy Tool Schema Loading* – Proposes two-pass tool injection to cut per-call token overhead (~3.5–5K tokens saved). High demand from local model users.
- **#5257** (11 comments, 16 👍) — *Generalized ACP client for multi-agent CLI orchestration* – Users want Hermes to orchestrate Claude Code, Codex CLI, etc. through a unified ACP client.
- **#4379** (15 comments) — *Token overhead analysis: 73% fixed overhead (~13.9K tokens)* – Dashboard-based deep-dive into waste; strong resonance with performance-conscious deployers.
- **#13834** (12 comments, 3 👍) — *Hermes openai-codex fails where official Codex CLI works* – Network/environment-specific regression affecting macOS users.
- **#19566** (8 comments, 1 👍) — *OpenAI-Codex credential pool drops newly added credentials after stale auth.json rewrite* – P1 security/stability bug.
- **#33801** (7 comments) — *Secret redaction corrupts code syntax in tool output* – P2 bug impacting tool reliability.
- **#39691** (7 comments, 10 👍) — *Integrate headroom-ai for tool output compression* – Alternative to conversation-level compression, targeting per-tool output.

Underlying needs: **Token cost reduction** (both schema and output), **multi-agent orchestration**, and **credential reliability** dominate community attention.

---

## Bugs & Stability

**New bugs reported today (ranked by severity):**

| Issue | Severity | Description | Fix PR exists? |
|-------|----------|-------------|----------------|
| **#51826** | P1 | `tirith` auto-install retries unboundedly and leaks `/tmp/tirith-install-*` dirs, exhausting disk | ✅ **#51831** |
| **#19566** (updated) | P1 | OpenAI-Codex credential pool drops newly added credential during rotation | PR **#51821** (advisory treatment of `reset_at`) |
| **#51800** | P2 | Context compression tail budget ignores `reasoning_content`, causing near-no-op compaction | ✅ **#51822** |
| **#51691** (updated) | P2 | `skill_view` returns UTF-8 decode error on Windows Desktop (cp936 locale) | ✅ **#51823** |
| **#51833** | P2 | Model dropdown shows reasoning effort as part of model name, duplicate entries | No fix PR yet |
| **#51829** | P2 | `/learn` slash command shows ack but does not trigger LLM in Desktop GUI | No fix PR yet |
| **#51828** | P2 | Telegram streaming truncation triggers full response re-generation (not continuation) | No fix PR yet |
| **#51825** | P2 | All skills run on single CPU core, no multi-core utilization | No fix PR yet |
| **#51820** | P2 | `execute_code` sandbox files invisible to host filesystem (v0.17) | No fix PR yet |
| **#33801** (updated) | P2 | Secret redaction corrupts code syntax in tool output | No fix PR yet |
| **#50663** (updated) | P2 | z.ai rate-limits Hermes during “peak hours” | No fix PR yet |
| **#48056** (updated) | P2 | Telegram DM topic cron delivery falls back out of topic | No fix PR yet |

**Other bug-related PRs:**  
- **#51834** — Fix computer-use wrapper calling nonexistent `screenshot` tool.  
- **#51824** — Harden pet thumbnail_png against path traversal.  
- **#51585** (P0) — Fix content-address prompt cache key for cron jobs hitting provider prefix cache (fixes cache misses).

---

## Feature Requests & Roadmap Signals

Notable proposals that may shape upcoming releases:

- **Lazy Tool Schema Loading (#6839)** – Already high traction; likely candidate for v0.18 or v0.19.
- **Tool Output Compression via headroom-ai (#39691)** – Complementary to context compression; strong community support.
- **Generalized ACP Client (#5257)** – Would make Hermes an orchestration hub; could coexist with existing Copilot client.
- **Plugin SDK additive `profileScope` surface (#51816)** – Enables dashboard plugins to react to profile switches.
- **Real-time Voice (#51827)** – WebRTC voice as gateway platform plugin; likely experimental in next release.
- **Valut Credential Sharing (#51832)** – Client-side secret vault; useful for multi-agent workflows.
- **OTLP Observability Plugin (#48184)** – Enterprise-grade telemetry export.
- **Subagent Presets Skill (#51795)** – Simplifies delegation for non-power-users.
- **Zulip Integration (#3335)** – Long-running PR (since March) adding another messaging platform.

**Token efficiency** and **multi-agent coordination** are the two dominant roadmap themes.

---

## User Feedback Summary

**Pain points voiced today:**

- **Token bloat**: “73% of each API call is fixed overhead” (#4379) — users with limited budgets (local models, capped API keys) feel it most.
- **Credential management fragility**: Dropped credentials (#19566), confusing `reset_at` behavior, and rate limiting (#50663) frustrate reliability.
- **Tool output corruption**: Secret redaction breaking Python/Shell syntax (#33801) and sandbox files invisible (#51820) hurt developer trust.
- **Platform-specific regressions**: Windows UTF-8 decode error (#51691), macOS `codex` CLI incompatibility (#13834), Telegram streaming re-generation (#51828).
- **Missing multi-core support**: All skills pinned to one CPU core (#51825) limits concurrency for heavy workloads.

**Satisfaction signals**: Positive reaction (👍) to token-saving proposals (#6839, #39691) and to multi-agent orchestration (#5257) shows users appreciate proactive improvements. The community is engaged, detailed (dashboard analysis, reproducible bug reports), and willing to contribute code (many fix PRs from external contributors).

---

## Backlog Watch

Older, high-impact items still awaiting maintainer attention:

- **#6839** (Apr 9) — Lazy Tool Schema Loading: 26 comments, 14 👍, no official maintainer response visible.
- **#4379** (Apr 1) — Token overhead analysis: 15 comments, no maintainer reply.
- **#13834** (Apr 22) — openai-codex failure: 12 comments, 3 👍, needs repro guidance.
- **#5257** (Apr 5) — Generalized ACP client: 11 comments, 16 👍, no design sign-off.
- **#19566** (May 4) — Credential pool drop (P1): 8 comments, has a fix PR (#51821) but not yet merged.
- **#33801** (May 28) — Secret redaction corruption: 7 comments, P2, no fix PR yet.
- **#39691** (Jun 5) — headroom-ai compression: 7 comments, 10 👍, still in discussion.
- **PR #3335** (Mar 27) — Zulip integration: 100+ files, heavily reviewed but stalled.
- **PR #8427** (Apr 12) — Vertex AI provider: feature-complete but unmerged.

These items represent **critical long-term value** (token cost, platform parity, credential stability) and would benefit from explicit roadmap prioritization or maintainer feedback.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest – 2026-06-24

## Today’s Overview
PicoClaw shows high activity with 17 pull requests and 2 issues updated in the last 24 hours, signaling an active maintenance and feature-development period. The project closed 5 PRs and 1 issue, while 12 PRs remain open, several addressing critical bugs or introducing substantive features. The development cadence remains healthy, with contributions from multiple authors and a focus on platform compatibility (Windows, Android/Termux), security hardening, and OpenAI-compatible provider integration. No new releases were published today.

## Releases
*None this period.*

## Project Progress (Merged/Closed PRs in Last 24h)
Five pull requests were closed (likely merged) and one issue was resolved:

- **[#3154 (closed)](https://github.com/sipeed/picoclaw/pull/3154)** – `fix(openai_compat): recover Doubao Seed tool calls leaked as <seed:to…` – Fixes a Volcengine Doubao Seed model quirk where tool calls appear as raw XML in message content instead of the standard `tool_calls` field. This improves reliability for users of that model.
- **[#3059 (closed)](https://github.com/sipeed/picoclaw/pull/3059)** – `fix: explicitly ignore Close() errors in error paths and retry loops` – Cleanup of linter warnings by acknowledging ignored errors in resource cleanup.
- **[#3054 (closed)](https://github.com/sipeed/picoclaw/pull/3054)** – `fix(line): add ok checks for sync.Map type assertions in Send` – Prevents potential panics from unchecked type assertions in LINE channel.
- **[#3047 (closed)](https://github.com/sipeed/picoclaw/pull/3047)** – `fix(web): restore full JSONL history for session detail` – Adds a detail-only JSONL reader so archived messages are visible in session details without affecting list pagination.
- **[#2888 (closed)](https://github.com/sipeed/picoclaw/pull/2888)** – `PR: 55N10E/picoclaw-1#1 Fix/tool config load image reaction` – A community-contributed fix for tool config image reactions (details sparse).
- **[#3015 (closed)](https://github.com/sipeed/picoclaw/issues/3015)** – `[BUG] QQ Channel Connection Failed After Windows Release Build` – Resolved as stale; details on Windows QQ channel token retrieval timeout.

## Community Hot Topics
The most discussed item this period is **[Issue #3015](https://github.com/sipeed/picoclaw/issues/3015)** (closed, 4 comments), which describes a Windows-specific QQ channel startup failure due to token retrieval timeout. The issue garnered multiple comments, indicating user frustration with Windows platform support. Although now closed, it highlights ongoing cross-platform gaps.

No other issue or PR received comments today.

## Bugs & Stability
Two bugs were reported or updated in the last 24 hours, with fix PRs already in progress for one:

### High Severity
- **[Issue #3164 – OPEN](https://github.com/sipeed/picoclaw/issues/3164)** – `[BUG] Process hooks crash gateway on Android/Termux (v0.2.9, config v3)` – The gateway crashes within 2 seconds of startup when any process hook (JSON-RPC over stdio) is active. This is a critical regression for Android/Termux users. No fix PR exists yet; the issue is open with zero comments, indicating it needs immediate attention.

### Medium Severity
- **(Closed) [Issue #3015](https://github.com/sipeed/picoclaw/issues/3015)** – QQ channel failure on Windows (token timeout). Closed as stale, but the underlying problem may persist on Windows builds.

### Fix PRs Addressing Other Bugs
- **[PR #3165](https://github.com/sipeed/picoclaw/pull/3165)** – Recovers Doubao Seed XML tool calls (feature fix for OpenAI-compat provider).
- **[PR #3166](https://github.com/sipeed/picoclaw/pull/3166)** – Fixes a build failure (`undefined: log`) in `openai_compat` by switching to structured logger.
- **[PR #3161](https://github.com/sipeed/picoclaw/pull/3161)** – Security fix: keeps deny patterns active even when a command matches custom allow rules.
- **[PR #3160](https://github.com/sipeed/picoclaw/pull/3160)** – Security fix: rejects cross-site launcher setup requests (browser provenance checks).

## Feature Requests & Roadmap Signals
Several open PRs indicate likely directions for the next minor release:

- **Remote Agent Mode** – **[PR #3118](https://github.com/sipeed/picoclaw/pull/3118)** adds a `--remote ws://...` flag to `picoclaw agent`, enabling a remote Pico WebSocket mode. This would allow the agent to connect to a gateway over a WebSocket, useful for distributed deployments.
- **AWS Bedrock Prompt Caching** – **[PR #3163](https://github.com/sipeed/picoclaw/pull/3163)** leverages Bedrock’s Converse API cache points to reduce input costs (~0.1× on reads). This is a clear efficiency gain for Bedrock users.
- **Telegram Reply-as-Mention** – **[PR #2975](https://github.com/sipeed/picoclaw/pull/2975)** (open since May 30) treats replying to a bot in group chats as an @mention. This addresses a common user expectation for Telegram bots.
- **Inline Data URL Fix** – **[PR #3115](https://github.com/sipeed/picoclaw/pull/3115)** prevents session-history corruption when generic tools (read_file, exec) return legitimate data URLs. This is a stability improvement rather than a feature.

Given the volume and maturity of these PRs, the next release will likely include the remote agent mode, Bedrock caching, and the Telegram mention fix.

## User Feedback Summary
User feedback this period is limited but revealing:

- **Windows QQ users** (Issue #3015) experienced a frustrating token retrieval timeout. The issue was closed as stale without a public fix, which may leave some Windows users unsatisfied.
- **Android/Termux users** (Issue #3164) face a complete crash of the gateway with process hooks – a blocker for anyone using custom hooks on mobile. No workaround is documented yet.
- **Positive signals**: Multiple contributors are actively fixing bugs and adding features (e.g., #3154, #3165, #3166), and the project’s test coverage is expanding (PR #3158 adds Windows path handling tests). The ecosystem feels responsive.

## Backlog Watch
Several PRs and issues require maintainer attention:

- **[PR #2975](https://github.com/sipeed/picoclaw/pull/2975)** – `feat(telegram): treat reply to bot message as mention in group chats` – Open since May 30 (stale label), needs review and potential merge.
- **[PR #3118](https://github.com/sipeed/picoclaw/pull/3118)** – Remote Pico WebSocket mode – Open since June 12, active but awaiting final review. Would be a major feature.
- **[PR #3115](https://github.com/sipeed/picoclaw/pull/3115)** – Inline data URL fix – Same author as #3118, likely part of the same effort.
- **Dependabot PRs** – #3104, #3103, #3100 (all open since June 11) bump frontend dependencies. These are low risk but should be merged to stay current.
- **[Issue #3164](https://github.com/sipeed/picoclaw/issues/3164)** – Android/Termux crash – No assignee, no comments, urgent for mobile users.
- **[PR #3163](https://github.com/sipeed/picoclaw/pull/3163)** – Bedrock caching – Very recent (June 23), but an important feature that could use prompt review.

Overall, the project is in a healthy state with strong contributor velocity, though the Android crash and lingering stale PRs warrant priority attention from maintainers.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-24

## 1. Today's Overview

NanoClaw had a highly productive 24 hours with **12 pull requests** updated, of which **8 were merged or closed**. Activity was concentrated on the Chat SDK version bump (4.29.0) across all registry branches (`main`, `channels`, `providers`), a new Slack Socket Mode adapter, and the introduction of generic extension-point seams for downstream forks. The project also received one new open issue reporting a security‑relevant port binding conflict. No new releases were published today, but the incoming features and fixes signal a stable, fast‑moving codebase.

## 2. Releases

*None.* No new version was cut in the last 24 hours.

## 3. Project Progress – Merged & Closed PRs

Eight PRs were merged or closed today, advancing several areas:

| PR | Title | Summary |
|----|-------|---------|
| [#2841](https://github.com/nanocoai/nanoclaw/pull/2841) | Generic inert extension-point seams (registerX/applyX) across host + container runtime | Adds byte‑identical extension points for downstream forks to attach custom behaviour without changing upstream code. |
| [#2839](https://github.com/nanocoai/nanoclaw/pull/2839) | chore(channels): bring Slack Socket Mode into channels | Merges the Slack Socket Mode implementation into the `channels` branch after a branch‑targeting fix. |
| [#2837](https://github.com/nanocoai/nanoclaw/pull/2837) | feat(slack): Socket Mode — adapter + guided setup (SLACK_APP_TOKEN) | Adds an outbound WebSocket adapter for Slack; no public endpoint required. Ideal for local dev or NAT’d hosts. |
| [#2836](https://github.com/nanocoai/nanoclaw/pull/2836) | chore(deps): bump chat SDK to 4.29.0 (providers) | Updates the `providers` registry branch to Chat SDK 4.29.0, keeping version lockstep with `main`. |
| [#2835](https://github.com/nanocoai/nanoclaw/pull/2835) | chore(deps): bump @chat-adapter/* + chat to 4.29.0 (channels) | Same SDK bump for the `channels` branch, updating install pins and setup scripts across 22 files. |
| [#2834](https://github.com/nanocoai/nanoclaw/pull/2834) | chore(deps): move chat SDK + channel-adapter pins to 4.29.0 | Core SDK bump on the `main` branch – foundational for all adapter compatibility. |
| [#2833](https://github.com/nanocoai/nanoclaw/pull/2833) | Feat/hook surface guard | Introduces a guard layer to enforce safe hook surface usage. |
| [#2826](https://github.com/nanocoai/nanoclaw/pull/2826) | fix(update-skills): nudge into skill updates, rebuild container on re‑apply | Fixes a usability issue where skill updates were presented as optional, leading to missed upstream fixes. Now container rebuild is enforced on re‑apply. |

## 4. Community Hot Topics

The only open issue, [#2840](https://github.com/nanocoai/nanoclaw/issues/2840), has sparked the day’s only community discussion (0 comments, but it is the sole active issue). It reports that NanoClaw v2 binds port 3000 on the external host IP, undermining the security purpose of the recommended SSH tunnel for Slack integration. No reactions or additional comments yet, but the bug touches a critical security‑setup flow.

Among open PRs, the most likely to generate discussion are:

- [PR #2843](https://github.com/nanocoai/nanoclaw/pull/2843) – **feat: add /learn skill** – a new skill that distills reusable skills from any source (directory, URL, past history). This adds a powerful meta‑capability.
- [PR #2842](https://github.com/nanocoai/nanoclaw/pull/2842) – **Generic inert extension‑point seams** – complements the already‑merged #2841 with additional seams.
- [PR #2832](https://github.com/nanocoai/nanoclaw/pull/2832) – **feat(approvals): reject with reason** – lets approvers attach a one‑line reason when rejecting a module request, improving agent adaptability.

None of these show comments or reactions yet, but they represent features that users have been requesting (better feedback, extensibility).

## 5. Bugs & Stability

**Moderate severity:** [Issue #2840](https://github.com/nanocoai/nanoclaw/issues/2840) reports that port 3000 (used for Slack tunnel instructions) is already bound on the external IP by NanoClaw itself, making the tunnel pointless. If the user follows the recommended “secure” path, the port is already occupied by the application, exposing the service directly. No fix PR exists yet. This is a design/security regression.

**Low severity (fixed):** [PR #2826](https://github.com/nanocoai/nanoclaw/pull/2826) resolved a stability issue in the `/update-nanoclaw` skill where skill updates were silently skipped, potentially leaving users with outdated channel/provider code.

No crashes or regressions beyond these were reported.

## 6. Feature Requests & Roadmap Signals

Several open PRs point toward likely upcoming features in the next NanoClaw release:

- **`/learn` skill ([#2843](https://github.com/nanocoai/nanoclaw/pull/2843))** – distills new skills from arbitrary sources; a major capability expansion.
- **Reject with reason ([#2832](https://github.com/nanocoai/nanoclaw/pull/2832))** – extends the approval UX to give agents feedback on why a request was declined.
- **Manifest model router provider ([#2838](https://github.com/nanocoai/nanoclaw/pull/2838))** – adds a provider that routes model requests based on a manifest, enabling flexible model selection.
- **Extension‑point seams ([#2842](https://github.com/nanocoai/nanoclaw/pull/2842))** – the open counterpart to the merged #2841, further enabling forks to customise behaviour without modifying upstream code.

The SDK bump to 4.29.0 across all branches suggests the team is aligning the codebase in preparation for a release that depends on the new Chat SDK generation. The next version will likely bundle Slack Socket Mode, the approval UX improvement, and the `/learn` skill.

## 7. User Feedback Summary

- **Pain point – Slack tunnel setup:** The sole issue ([#2840](https://github.com/nanocoai/nanoclaw/issues/2840)) highlights a real user frustration: the recommended secure tunnel is invalidated by the application binding the same port externally. The user expects the tunnel to be genuinely secure; the current behaviour breaks that expectation.
- **Desire for better agent feedback:** The `reject with reason` PR ([#2832](https://github.com/nanocoai/nanoclaw/pull/2832)) stems from users wanting agents to adapt after a declined action, rather than just receiving a “declined” signal.
- **Positive response to Slack Socket Mode:** The merged PRs for Socket Mode ([#2837](https://github.com/nanocoai/nanoclaw/pull/2837), [#2839](https://github.com/nanocoai/nanoclaw/pull/2839)) address a long‑standing request for a setup that does not require a public HTTPS endpoint. This is likely to improve satisfaction for local and behind‑NAT users.
- **Update‑skill usability:** The fix in [#2826](https://github.com/nanocoai/nanoclaw/pull/2826) addresses a situation where users could unknowingly miss critical skill updates – a UX issue that could erode trust in the update mechanism.

## 8. Backlog Watch

No long‑unanswered issues or PRs were identified in today’s data. The only open issue ([#2840](https://github.com/nanocoai/nanoclaw/issues/2840)) is less than 24 hours old and has not yet received a maintainer reply. It should be prioritised due to its security implications. All other open PRs are recent (last 1–2 days) and are under active review.

---

*Generated from GitHub activity data for nanocoai/nanoclaw on 2026-06-24.*

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

Based on the provided GitHub data for the **NullClaw** project (github.com/nullclaw/nullclaw), here is the structured project digest for **2026-06-24**.

---

### Project Digest: NullClaw — 2026-06-24

#### 1. Today's Overview

The NullClaw project shows minimal activity today, with only one issue updated and one open pull request receiving attention. A critical bug affecting Windows 11 users (Issue #967) has been closed after being filed, but no new releases or merged code changes occurred in the last 24 hours. The sole open pull request (#783) remains a significant unmerged feature addition focusing on cron and security, indicating that major development work may be in a "pending review" or "stalled" state. Overall community engagement is very low, with the project appearing stable but not actively advancing.

#### 2. Releases

**No new releases were published** in the last 24 hours. The most recent available version remains v2026.5.29.

#### 3. Project Progress

- **Pull Requests Merged/Closed**: **0** merged or closed PRs today.
- **Active PRs**: **1** pull request (#783) remains open and was updated today.
    - **#783 (Open)**: A large feature branch titled *"feat(cron): cron subagent, run history, JSON output, security hardening"* by **yangf8**. This PR introduces a database-backed cron scheduler, job execution history, JSON CLI output, and security improvements. It has been open since April 7, 2026 (over 2.5 months), suggesting it is a significant, complex feature that is still under review or blocking on requirements.
    - **Link**: [PR #783](https://github.com/nullclaw/nullclaw/pull/783)

#### 4. Community Hot Topics

- **Issue #967 (Closed)**: The only issue updated today generated 2 comments and is the most active item on the board. The user **svier0** reports a high-frequency (>50%) `NoResponseContent` error on Windows 11 with the `Agnes-2.0-Flash` model. The core need appears to be **model provider API stability and compatibility on Windows**, specifically when using specific model endpoints. The fact that the same API key works with “picocla”（likely another client) suggests the bug is specific to NullClaw’s HTTP client or response parsing logic.
    - **Link**: [Issue #967](https://github.com/nullclaw/nullclaw/issues/967)

#### 5. Bugs & Stability

| Severity | Bug ID | Description | Status | Fix PR Exists? |
| :--- | :--- | :--- | :--- | :--- |
| **High** | [#967](https://github.com/nullclaw/nullclaw/issues/967) | `NoResponseContent` error on **Windows 11** with **Agnes-2.0-Flash** model. | Closed | No open fix PR |

**Analysis**: This is the highest severity bug reported recently, as it causes total failure for over half of user interactions on a specific platform (Windows) with a specific model. The issue is now closed, but no associated fix pull request is noted in the data, meaning either the bug was accepted as a known limitation, the user was asked to try a workaround, or a fix was merged without a linked PR in this snapshot.

#### 6. Feature Requests & Roadmap Signals

- **Active Feature Work**: The only clear roadmap signal is the long-running PR #783, which aims to add a **Cron Subagent Engine**, **Job History**, **JSON output** for CLI, and **Security Hardening**. This suggests the maintainers prioritize moving beyond simple chat agents toward **scheduled automation and developer-friendly tooling**.
- **No New Feature Requests**: There are no new feature requests in today’s data. The community focus is exclusively on fixing the stability bug.

**Prediction**: If PR #783 is merged, the next minor version (v2026.6.x) will likely focus on cron capabilities and CLI improvements.

#### 7. User Feedback Summary

- **Pain Point (High)**: **High error rate on Windows 11**. User **svier0** explicitly states the model works elsewhere, but NullClaw fails >50% of the time with a vague internal error. This indicates a stability issue that directly impacts user satisfaction for Windows users.
- **Satisfaction**: No positive feedback or evidence of satisfaction in this 24-hour window. The community seems silent, which could indicate either general satisfaction (no news is good news) or low engagement.
- **Use Case**: The user is clearly trying to run NullClaw as a local AI assistant on a desktop Windows environment using a third-party model provider API.

#### 8. Backlog Watch

| Item | Issue / PR | Reason for Watch |
| :--- | :--- | :--- |
| **Ancient PR (2.5 months)** | [PR #783](https://github.com/nullclaw/nullclaw/pull/783) | This pull request has been **open since April 7, 2026** with no merge. It contains substantial code changes. The prolonged open status may be causing community confusion about the project’s future direction or blocking other work. Maintainer attention is needed to either merge, reject, or update the PR. |
| **No other stale items** | — | No other issues or PRs appear in the data with long wait times. The project backlog appears small and well-handled aside from the large PR. |

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest – 2026-06-24

## 1. Today's Overview

The project saw **high activity** over the past 24 hours: 14 issues updated (9 open, 5 closed) and 48 pull requests updated (28 merged/closed, 20 open). No new releases were published. The team is actively addressing regressions (e.g., the “reborn task hang” bug #5139 was fixed) and advancing major feature work in channel ingress, context management, and WebUI improvements. Several older bugs remain open (e.g., Gmail auth inconsistencies, automation blocker) but have recent comments indicating investigation. Overall project health is solid, with a strong flow of merged PRs and focused issue resolution.

## 2. Releases

**None** – no releases in the last 24 hours.

## 3. Project Progress

The following PRs were **merged or closed** today, advancing key features and fixes:

- **#5181** – Allow Slack enablement env override (size L, low risk) – enables runtime toggling of Slack route without WebUI config.
- **#5143** – Pin Telegram host-ingress verifier fail‑closed (size S) – adds regression test and updates manifest‑driven channel plan.
- **#5178** – Skip NEAR AI MCP without durable auth storage (size M) – prevents dangling MCP activation in local-dev setups.
- **#5162** – [codex] Allow Slack enabled env override (size L) – duplicate/companion PR for Slack env override.
- **#2726** – Harden cheap‑model settings rollout (size L, medium risk) – ports settings persistence for `cheap_model` / `smart_routing_cascade` to current staging.
- **#5161** – Remove legacy Slack fields from hosted config (size XS) – cleans up deprecated env‑key fields.
- **#4492** – Fix configured extension credential staging (size XL, DB migration) – makes local‑dev product‑auth SecretStore‑backed, durable across restarts.
- **#4501** – CI: avoid runtime tests for Dependabot config updates (size M) – skips unnecessary test runs for config‑only bumps.
- **#4550** – Support full SHA GitHub branch creation (size M) – allows 40‑character commit SHAs as direct Git refs.
- **#5171** – Correct Reborn GitHub API requests (size XL) – fixes PAT auth request shapes in first‑party GitHub WASM extension.

Other notable merged PRs (from the full list of 28) include follow‑ups to Slack integration, dependency bumps, and documentation clarifications.

**Overall**: The team merged fixes for credential management, GitHub API, Slack configuration, and cheap‑model settings, while continuing to build toward manifest‑driven channel ingress and improved context handling.

## 4. Community Hot Topics

Despite low reaction counts, the following issues and PRs attracted attention (comments, scope, or urgency):

- **Issue #5182** [OPEN] – *“Reborn hosted observability: meaningful logs + failure diagnostics”* – 1 comment. Community operator wants better diagnostics from the binary, highlighting a gap in hosted deployments.  
  👉 [Issue #5182](https://github.com/nearai/ironclaw/issues/5182)

- **Issue #4986** [OPEN] – *“Recurring automation can become permanently blocked waiting for tool approval”* – 1 comment, 8 days old. A real workflow problem: automations requiring `builtin.http` get stuck in approval loops.  
  👉 [Issue #4986](https://github.com/nearai/ironclaw/issues/4986)

- **Issue #5169** [OPEN] – *“Bundled skills trip the prompt‑safety vocabulary denylist”* – 1 comment. Clean‑setup repro of a benign request failing due to API vocabulary in bundled skill instructions. Highlights a safety‑filter integration issue.  
  👉 [Issue #5169](https://github.com/nearai/ironclaw/issues/5169)

- **PR #5149** [OPEN] – *“Progressive tool disclosure (flag‑gated, default off)”* – size XL, core contributor. This is a major feature aimed at reducing prompt size and model timeout issues. Likely to generate community interest once enabled.  
  👉 [PR #5149](https://github.com/nearai/ironclaw/pull/5149)

- **PR #5107** [OPEN] – *“Manifest‑driven channel ingress contract”* – consolidates four earlier PRs, fundamental for extending channel support.  
  👉 [PR #5107](https://github.com/nearai/ironclaw/pull/5107)

**Underlying needs**: Operators demand better observability (logs, diagnostics); automation users need reliable approval workflows; safety filters must not break legitimate use cases; and the architecture is being refactored for maintainability (manifest‑driven, progressive disclosure).

## 5. Bugs & Stability

Ranked by severity:

| Severity | Issue | Description | Status | Fix PR |
|----------|-------|-------------|--------|--------|
| **Critical** | #5139 | `reborn` web/research tasks hang at init with zero LLM/tool calls, regressed in main (10 commits). Caused 21/147 failures in PinchBench daily. | **CLOSED** | (no explicit PR listed, but fix merged as part of main) |
| **High** | #4986 | Recurring automations blocked forever waiting for tool approval when `builtin.http` is needed. | **OPEN** | None |
| **High** | #5169 | Bundled skills trigger prompt‑safety denylist, causing misleading “temporary system issue” error. | **OPEN** | None |
| **Medium** | #4991 | WASM Google Drive auth failures dead‑end as `operation_failed` without retry or auth gate. | **CLOSED** | (fix likely part of #4492 or similar) |
| **Medium** | #5179 | Web UI logs unavailable for multi‑tenancy users. | **OPEN** | None |
| **Low** | #3733 | Invalid Gmail token shows success/activated toast then immediately re‑asks for OAuth. | **OPEN** | None |
| **Low** | #3732 | Gmail auth gate shows inconsistent UI (OAuth link vs. manual token). | **OPEN** | None |
| **Info** | #5173 | Daily failure taxonomy report (clawbench) – dominated by benchmark defects, not model quality. | **OPEN** | N/A |

Several fix PRs were merged today: #5171 (GitHub API request shapes), #5180 (populate provider on auth‑required gates), and #5178 (skip MCP without durable auth). The critical regression #5139 was closed, but the root cause (10‑commit range `2b2ccc55→704fcd43`) should be audited for long‑term stability.

## 6. Feature Requests & Roadmap Signals

- **#5182** – Reborn hosted observability: request for structured logs and failure diagnostics. Likely to be picked up soon given the operator feedback.
- **#5167** – Stop tracking `dist` in git; build from current code. This is a developer‑experience improvement that will be adopted to reduce PR churn.
- **#5122** – Add Reborn automation delete support (native trigger system). Already closed, meaning implementation is underway.
- **PR #5149** – Progressive tool disclosure (flag‑gated). Expected to be enabled by default in a future release after testing reduces context‑window overhead.
- **PR #5107** – Manifest‑driven channel ingress contract. This is a foundational refactor that will enable easier addition of new channels (e.g., Discord, custom webhooks) in the next minor version.
- **PR #5068** – Tool permissions + global auto‑approve settings in WebUI. This will address automation approval blocking (#4986) and is a strong candidate for the next release.
- **PR #5176** – Subagent thread harness (docs only). Signals a deeper reframing of subagents into resumable threads – a roadmap item that may span several versions.

**Prediction**: The next release will likely include the context management feature (#5149), manifest‑driven ingress (#5107), and the tool permissions UI (#5068), along with fixes for the Gmail auth inconsistencies and automation blocking.

## 7. User Feedback Summary

- **Pain points**:
  - *Automation approval loops* (#4986): “Create an automation that requires `builtin.http`… permanently blocked waiting for tool approval.” – user wants reliable unattended workflows.
  - *Misleading errors from safety denylist* (#5169): “Benign request terminally fails because a bundled skill’s instructions contain ordinary API vocabulary.” – user reports the error message is unhelpful (“temporary system issue”).
  - *Gmail auth inconsistency* (#3732, #3733): User sees different UI paths (OAuth link vs. manual token input) and invalid tokens accepted with false success toasts.
  - *WASM Google Drive auth silent failure* (#4991): expired tokens reported as generic `operation_failed` without refresh retry.

- **Use cases**:
  - Multi‑tenant deployments need WebUI logs (#5179).
  - Operators require better observability of hosted Reborn instances (#5182).
  - Users want finer‑grained tool permissions (auto‑approve per tool) – reflected in PR #5068.

- **Satisfaction**:
  - The rapid closing of critical regression #5139 and many merged PRs today indicates a responsive team.
  - No overt praise found, but the low reaction count on issues may suggest that community involvement is either small or that issues are being addressed before they grow.

## 8. Backlog Watch

The following important issues have remained open for an extended period without a confirmed fix or maintainer response:

| Issue | Created | Days Open | Last Update | Summary |
|-------|---------|-----------|-------------|---------|
| **#3732** | 2026-05-17 | 38 | 2026-06-23 | Gmail auth gate shows inconsistent UI (OAuth link vs. manual token) – 1 comment, no assignee. |
| **#3733** | 2026-05-17 | 38 | 2026-06-23 | Invalid Gmail token shows success toast – same age, same author, no recent progress. |
| **#4986** | 2026-06-16 | 8 | 2026-06-24 | Automation permanently blocked waiting for tool approval – critical, but no fix PR yet. |
| **#4108** | 2026-05-27 | 28 | 2026-06-24 | Nightly E2E failed – auto‑generated, but failure persists; no manual intervention noted. |
| **#5169** | 2026-06-23 | 1 | 2026-06-23 | Bundled skills trigger safety denylist – very new but high impact; no fix PR yet. |

**Maintainer attention needed**: The two‑month‑old Gmail auth issues (#3732, #3733) should be prioritised as they impact user trust. The automation blocker (#4986) and the safety denylist issue (#5169) are high‑severity and lack visible resolution. The nightly E2E failure (#4108) also suggests flaky infrastructure that could erode confidence.

---

*Generated from GitHub data for nearai/ironclaw on 2026-06-24.*

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-06-24

## Today's Overview
The project saw exceptionally high activity, with **45 pull requests updated** in the last 24 hours — an astonishing **42 were merged or closed**, indicating a major consolidation push. In contrast, only **1 issue** received updates, reflecting a focus on shipping fixes rather than surfacing new problems. No new releases were cut today. The bulk of PRs come from core contributors (`fisherdaddy`, `liuzhq1986`) and span stability fixes, cross‑platform gateway improvements, model support updates, and UI polish. The one active issue reports a data‑loss bug in scheduled tasks, which has not yet been resolved.

## Releases
**None** — no new versions were published on this date.

## Project Progress
The 42 merged/closed PRs touch nearly every area of the stack. Key advancements include:

- **OpenClaw Gateway Stability** – PRs #2196, #2195 unify the gateway launch path across macOS, Linux, and Windows using `ELECTRON_RUN_AS_NODE`, preventing spurious dock apps and child‑process misinterpretation.
- **Token Burn & Tool Loop Fixes** – PR #2049 introduces an upstream aborted‑loop breaker that was missing, ending the “infinite token burn” during idle periods reported by users. A follow‑up (#2051) re‑refines the breaker logic.
- **Session & Chat Reliability** – PRs #2050 (gateway sessions timeout), #2047 (session freezing), and #2104 (MCP session timeout during gateway config reload) address concurrency blocks that could freeze the UI or lose messages.
- **IM & Cowork Improvements** – PR #2063 scopes reply assembly to the current turn and strips thinking blocks; PR #2078 emits selected‑skill routing metadata instead of inlining prompts; PR #2058 tightens grace periods when tool results are large.
- **Model & UI Updates** – PR #2089 adds MiniMax M3 and updates BYOK model default context windows; PR #2102 preserves user‑configured context windows and adds Mimo v2.5; PR #2053 fixes model select UI; PR #2088 updates kit UI; PR #2048 filters empty LLM streaming chunks.
- **Platform Fixes** – PR #2057 replaces deprecated VBScript launcher with hidden PowerShell; PR #2086 fixes a WeChat integration bug during reinstalls.

## Community Hot Topics
The only issue updated today is the top active thread:

- **#1394** – **Scheduled tasks are permanently deleted after a single run (when set to “not repeat”)**  
  *Author: zqgittest* | [Issue Link](https://github.com/netease-youdao/LobsterAI/issues/1394)  
  This user reports that creating a non‑repeating scheduled task and triggering it (manually or automatically) causes the task to be auto‑deleted. The expected behavior is that the task should remain editable for future use. The issue has 1 comment and 0 upvotes, but it represents a potential data‑loss concern.

## Bugs & Stability

| Severity | Issue / PR | Description | Status |
|----------|-------------|-------------|--------|
| **Medium** | #1394 | Scheduled tasks vanish after execution (non‑repeating) | **Open** – no fix PR yet |
| **Critical (fixed today)** | #2049 | Infinite token burn from aborted tool loops | **Merged** |
| **High (fixed today)** | #2047 | Session freezing during concurrent operations | **Merged** |
| **High (fixed today)** | #2050 | Gateway `sessions.patch` timeout blocks chat.send | **Merged** |
| **Medium (fixed today)** | #2063 | IM reply assembly not scoped to current turn, leaking thinking blocks | **Merged** |
| **Medium (fixed today)** | #2104 | MCP session timeout during gateway config reload | **Merged** |
| **Low (fixed today)** | #2048 | Empty LLM streaming data in output | **Merged** |

The list shows the team is aggressively closing stability gaps; the remaining open bug (#1394) is the only unresolved high‑visibility item.

## Feature Requests & Roadmap Signals
While no explicit feature requests were filed today, the merged PRs reveal ongoing roadmap investments:

- **Model Support Expansion** – MiniMax M3, Mimo v2.5, and user‑configurable context widths (#2089, #2102) signal a commitment to keeping pace with the fast‑moving LLM landscape.
- **Cowork / Agent Routing** – PR #2078 moves from in‑line prompts to metadata‑based skill routing, laying groundwork for a more modular and composable agent system.
- **UI & Kit Refresh** – PRs #2053, #2088 indicate continued polish of the model selector and kit component UI, likely targeting a smoother user experience in the next release.

## User Feedback Summary
Two concrete pain points emerged from the data:

1. **Scheduled task deletion** – the issue author expects tasks to be preserved after execution for later editing. This contradicts the current “delete after run” behaviour.
2. **Token burn during idle** – multiple users reported continuous token consumption even when no conversation was active (mentioned in PR #2049 description). The merged fix should resolve this.

No explicit satisfaction signals are visible, but the heavy PR volume suggests active testing and iteration based on real usage.

## Backlog Watch
- **#1394** – Scheduled tasks auto‑deletion (open since 2026-04-03, updated 2026-06-24). The issue has only one comment, but the problem affects data persistence. With no assignee or linked PR, it requires maintainer attention.  
  [Issue Link](https://github.com/netease-youdao/LobsterAI/issues/1394)

All other previously open items appear to have been addressed in today’s merges. The 3 open PRs (not listed in the top 20) likely need review or have stalled. Their status should be checked.

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

# CoPaw Project Digest – 2026-06-24

## 1. Today's Overview

CoPaw (QwenPaw) saw elevated activity with **16 issues updated** (11 open, 5 closed) and **50 PRs updated** (40 open, 10 merged/closed) in the last 24 hours. No new releases were published. The project continues to mature rapidly: two significant **frontend stability bugs** were reported (large session crashes, layout rendering failures), while backend fixes for cron jobs and MCP tool handling were merged. Community contributions remain strong, with several first-time contributor PRs under review. Overall project health is active, though memory consumption and custom provider compatibility remain pain points.

## 2. Releases

*No new releases in the last 24 hours.*

## 3. Project Progress

The following PRs were **merged or closed** today, representing concrete progress:

- **fix(cron): restore cron session compatibility and sanitize malformed tool_call history** ([#5475](agentscope-ai/QwenPaw PR #5475)) — Fixes three issues causing cron job failures on legacy or corrupted session state.
- **fix(cron): allow editing and deleting enabled cron jobs** ([#5483](agentscope-ai/QwenPaw PR #5483)) — Enables modification of active cron jobs without requiring delete+recreate.
- **feat(console): improve SkillPool mobile responsive layout** ([#5368](agentscope-ai/QwenPaw PR #5368)) — CSS-only fix for mobile overflow on the settings page.
- **feat(ci): end-to-end UI verification for desktop releases** ([#5428](agentscope-ai/QwenPaw PR #5428)) — Adds automated checks to the desktop release pipeline.

Additionally, 5 issues were closed (including invalid/question items), indicating ongoing triage.

## 4. Community Hot Topics

The most active discussions today involve **custom provider function calling** and **startup crashes**:

- **Issue #5345** ([link](agentscope-ai/QwenPaw Issue #5345)) — *Custom OpenAI-compatible providers (e.g. OMLX) don't support function calling* (8 comments). Users report that OMLX works with other agents but not QwenPaw, highlighting a gap in provider compatibility.
- **Issue #5379** ([link](agentscope-ai/QwenPaw Issue #5379)) — *Internal Server Error after pip install* (5 comments). New users face immediate startup failure, creating a poor first impression.
- **Issue #5015** ([closed](agentscope-ai/QwenPaw Issue #5015)) & **#5264** ([closed](agentscope-ai/QwenPaw Issue #5264)) — Both had 5 comments, discussing frontend lag on Windows and Feishu group-chat misrouting respectively.

**Underlying needs:** Users demand **reliable custom model integration**, **out-of-the-box stability**, and **smooth mobile/responsive UI experiences**.

## 5. Bugs & Stability

New bugs reported today, ranked by severity:

| Severity | Issue | Description | Fix PR Exists? |
|----------|-------|-------------|----------------|
| 🔴 High | [#5401](agentscope-ai/QwenPaw Issue #5401) | Console frontend crashes/whitescreens on sessions with large tool-use history | No |
| 🔴 High | [#5479](agentscope-ai/QwenPaw Issue #5479) | Frontend crash when opening sessions >500KB ("unexpected error") | No |
| 🔴 High | [#5379](agentscope-ai/QwenPaw Issue #5379) | Internal Server Error immediately after pip install | No |
| 🟡 Medium | [#5472](agentscope-ai/QwenPaw Issue #5472) | GLM-5.x model fails with `json_schema_converter.cc` error | No |
| 🟡 Medium | [#5480](agentscope-ai/QwenPaw Issue #5480) | Long messages display with broken layout (fixes on tab switch) | No |
| 🟢 Low | [#5476](agentscope-ai/QwenPaw Issue #5476) | Mobile cannot switch agents (closed, likely duplicate) | N/A |

**Notable fix PRs merged today:** [#5475](agentscope-ai/QwenPaw PR #5475) (cron session compatibility) and [#5483](agentscope-ai/QwenPaw PR #5483) (cron editing). Two PRs opened today address tool input JSON decode ([#5486](agentscope-ai/QwenPaw PR #5486)) and MCP tool name sanitization ([#5485](agentscope-ai/QwenPaw PR #5485)).

## 6. Feature Requests & Roadmap Signals

Several enhancement issues and open PRs point toward the next version's likely focus:

- **Plugin management via pip** ([#5484](agentscope-ai/QwenPaw Issue #5484)) — Allows installing plugins directly from PyPI, reducing dependency on ZIP uploads. Already has a PR.
- **OpenAI response format support** ([#5489](agentscope-ai/QwenPaw Issue #5489)) — Extension for multi-modal message flows.
- **MCP tool name display** ([#5231](agentscope-ai/QwenPaw Issue #5231)) — Show human-readable names in UI while using cleaned names for model calls.
- **Kimi Coding Plan Models** ([#5427](agentscope-ai/QwenPaw Issue #5427)) — Add Anthropic-compatible endpoint for Kimi K2 Code.
- **Per-user timestamp prefix** ([#5455](agentscope-ai/QwenPaw Issue #5455)) — Suggestion to move current time from system context to per-user message prefix.

**Feature PRs in flight:**
- **Desktop GUI automation with UIA** ([#5187](agentscope-ai/QwenPaw PR #5187)) — Windows computer-use tool.
- **Tauri auto-updater** ([#4669](agentscope-ai/QwenPaw PR #4669)) — Desktop update flow.
- **DataPaw plugin** ([#4622](agentscope-ai/QwenPaw PR #4622)) — 12 BI skills for data analysis.
- **User message thumbnail sidebar** ([#5488](agentscope-ai/QwenPaw PR #5488)) — Navigation improvements.
- **Memory search enhancement** ([#5482](agentscope-ai/QwenPaw PR #5482)) — Simplified metadata.

These suggest **v2.0 or imminent release** will include plugin pip support, enhanced desktop features, and frontend UX improvements.

## 7. User Feedback Summary

**Pain points expressed:**
- **Memory consumption** — "刚启动已经1.4g了" (1.4GB at startup, issue [#5441](agentscope-ai/QwenPaw Issue #5441)).
- **Frontend performance** — Multiple users report crashes/lag with large sessions ([#5401](agentscope-ai/QwenPaw Issue #5401), [#5479](agentscope-ai/QwenPaw Issue #5479), [#5015](agentscope-ai/QwenPaw Issue #5015) closed).
- **Custom provider incompatibility** — OMLX/LM Studio users cannot use function calling ([#5345](agentscope-ai/QwenPaw Issue #5345)).
- **Channel misrouting** — Feishu group chat replies sent to private chat ([#5264](agentscope-ai/QwenPaw Issue #5264) closed). DingTalk sessions invisible ([#5177](agentscope-ai/QwenPaw Issue #5177) closed).
- **Startup failure** — pip install leads to immediate 500 error ([#5379](agentscope-ai/QwenPaw Issue #5379)).

**Satisfaction signals:**
- Positive engagement with CLI cron update feature (PR #5210).
- Enthusiasm for desktop GUI automation (PR #5187) and DataPaw plugin (PR #4622).
- Quick closure of invalid/question issues shows responsive maintainers.

## 8. Backlog Watch

The following open PRs and issues have remained unanswered for an extended period and may need maintainer attention:

| Item | Age (Days) | Status |
|------|------------|--------|
| [#4041](agentscope-ai/QwenPaw PR #4041) — System tray startup (win32) | 50 | Under Review (first-time contributor) |
| [#4669](agentscope-ai/QwenPaw PR #4669) — Tauri auto-updater | 30 | Open, needs review |
| [#4622](agentscope-ai/QwenPaw PR #4622) — DataPaw plugin | 33 | Under Review (first-time contributor) |
| [#5210](agentscope-ai/QwenPaw PR #5210) — CLI cron update command | 9 | Under Review (first-time contributor) |
| [#5153](agentscope-ai/QwenPaw PR #5153) — pywebview instant-window startup | 12 | Open |
| [#5213](agentscope-ai/QwenPaw PR #5213) — MCP access policy layout | 8 | Open |

These PRs represent community contributions that risk stagnation; most are feature additions or improvements that would benefit the project. No issues older than 7 days remain open without recent comments, suggesting good triage responsiveness.

---

*Digest generated from CoPaw (github.com/agentscope-ai/CoPaw) activity on 2026-06-24.*

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-06-24

## 1. Today's Overview
ZeroClaw saw heavy activity over the past 24 hours, with **50 pull requests** updated and **13 issues** refreshed (12 still open). The flurry is concentrated on two themes: **hardening the delegate tool security model** — multiple fix PRs target a severe bypass (#8279) — and **preparing the v0.8.2 release**, tracked by two coordination tickets (#7314 for WASM plugins, #8181 for non-plugin work). The supply‑chain signing RFC (#8177) and a large in‑app upgrade feature (#8173) also attracted maintainer attention. No new releases were cut today, but the release pipeline now includes SLSA provenance attestation (PR #8277), signaling that v0.8.2 packaging is being readied.

## 2. Releases
None. The most recent release remains v0.8.1 (no tag shown); v0.8.2 is in active preparation.

## 3. Project Progress
Nine pull requests were merged or closed in the last 24 hours. Key advancements include:

- **Observability** — PR #8065 (`feat(observability): correlate logs by trace_id + record per‑call cost_usd`) was merged, adding trace correlation and cost tracking to logs.
- **Config hardening** — PR #8098 (`fix(config): reject creating the reserved default agent across all operator surfaces`) was merged, preventing creation of the `default` agent through config surfaces, fixing an asymmetry.
- **Bug fix (image attachment)** — Issue #8151 was closed after the underlying fix for deferred image reference loss was applied (likely via a channel‑layer patch).
- **Delegation security fixes** — Three PRs (#8284, #8285, #8286) were opened to address the critical S0 bug #8279 (see §5). These are not yet merged but represent rapid response.
- **Supply‑chain & CI** — PR #8277 (`ci(workflows): add SLSA provenance attestation`) was opened to implement Phase A of RFC #8177; PR #8129 added `cargo-audit` to the PR gate.

Several other bug‑fix PRs (e.g., #8084, #8100, #8163, #8164, #8205, #8278, #8280–#8283) remain open but signal active work on provider vision support, Kimi endpoint, version‑mismatch detection, and voice‑wake channel completeness.

## 4. Community Hot Topics

- **#6289 – Prompt‑triggered install suggestions for missing skills**  
  *zeroclaw-labs/zeroclaw/issues/6289* (5 comments, updated today)  
  Users want ZeroClaw to proactively suggest installing skills or plugins when a user asks for a capability the agent doesn’t have. The discussion reflects a desire for better discoverability of the growing skill ecosystem.

- **#8177 – RFC: Supply chain signing (hardware PGP, hermetic builds, SLSA)**  
  *zeroclaw-labs/zeroclaw/issues/8177* (5 comments, updated today)  
  A detailed proposal for hardware‑backed signing and SLSA provenance, following the “StageX” model. The community and maintainers are engaged on implementation phases, with corresponding CI PRs already landing.

- **#6140 – Hybrid skills + WASM tools**  
  *zeroclaw-labs/zeroclaw/issues/6140* (3 comments, updated today)  
  Discussion of plugins that combine a markdown skill definition with a WASM binary. This is the next step after plain SKILL.md plugins and is a core architectural topic for the plugin system.

- **#6943 – Deconflict plugin system goals (RFC)**  
  *zeroclaw-labs/zeroclaw/issues/6943* (3 comments, updated today)  
  The RFC identifies mutually exclusive commitments in the foundational design document (FND‑001) regarding Extism vs. direct wasmtime integration. Community feedback is shaping Phase 2 of the plugin refactor.

## 5. Bugs & Stability

| Severity | Issue | Summary | Fix Status |
|----------|-------|---------|------------|
| **S0 – data loss/security risk** | [#8279](https://github.com/zeroclaw-labs/zeroclaw/issues/8279) | `delegate` tool bypasses parent’s tool allowlist — sub‑agent can invoke forbidden tools | 3 fix PRs open (#8284, #8285, #8286) |
| **S1 – workflow blocked** | [#8151](https://github.com/zeroclaw-labs/zeroclaw/issues/8151) (closed) | Deferred image attachment loses reference; bot later denies seeing it | Fixed & closed |
| **S2 – degraded behavior** | [#5903](https://github.com/zeroclaw-labs/zeroclaw/issues/5903) | MCP stdio child processes leak ~48 orphans per day with heartbeat | Open, accepted, P1 |
| **S2 – degraded behavior** | [#7733](https://github.com/zeroclaw-labs/zeroclaw/issues/7733) | `mcp_bundles` parsed but never enforced at runtime — silent security no‑op | Open, accepted, P1 |
| **S2 – degraded behavior** | [#7623](https://github.com/zeroclaw-labs/zeroclaw/issues/7623) | Delegate to Codex/OAuth sub‑agent still forwards coordinator’s API key after #7266 | Open, in‑progress, P1 |
| **S3 – minor** | [#8275](https://github.com/zeroclaw-labs/zeroclaw/issues/8275) | Scoop manifest missing `zerocode.exe` from PATH | Open, no assignee |

New bugs today include the critical #8279 (with immediate fix PRs), #8275 (minor packaging gap), and the closed #8151. The long‑standing MCP orphan leak (#5903) and the `mcp_bundles` enforcement gap (#7733) remain high‑priority open issues. Multiple fix PRs for #8279 suggest a fix will land within the next 24–48 hours.

## 6. Feature Requests & Roadmap Signals

The **v0.8.2 release** is the primary focus, with two trackers coordinating work:

- **[#7314] WASM plugin program** – covers plugin architecture, WIT/component‑model, lifecycle hooks, host‑function security.  
- **[#8181] Non‑plugin queue** – includes config fixes, observability, provider updates, CI hardening, and the supply‑chain signing RFC.

Notable features with high community interest and active implementation:

- **In‑app upgrade with auto‑restart** – PR #8173 (large, needs‑maintainer‑review) adds a web dashboard version‑tag upgrade button. Likely for v0.8.2 or v0.9.
- **SLSA provenance & supply‑chain signing** – RFC #8177 and PR #8277 are advancing fast; expect SLSA Level 2 attestation in the next release.
- **WASM plugin lifecycle hooks** – RFC #7822 proposes `PluginCapability::Hook` for sandboxed event listeners. Accepted but no implementation PR yet.
- **Prompt‑triggered skill suggestions** – #6289 is still pending design work; may land in a follow‑up after the plugin system stabilises.

## 7. User Feedback Summary

Real pain points voiced in this batch:

- **Security isolation concerns** — #8279 (tool allowlist bypass) and #7623 (API key bleed) show users are hitting gaps in the delegation model. The rapid response with multiple fix PRs indicates maintainers take these seriously.
- **MCP reliability** — #5903 (orphan processes) and #7733 (unenforced mcp_bundles) frustrate users relying on the MCP integration for long‑running agents.
- **Discovery friction** — #6289 reflects dissatisfaction with having to manually know about and install skills/plugins. Users expect the agent to guide them.
- **Packaging completeness** — #8275 (missing zerocode shim on Windows) is a minor but annoying onboarding gap.
- **Multimodal handling** — #8151 (lost image reference) was a workflow blocker for users sending images in Matrix; the fix is now live.

Overall, users are pushing for **harder security boundaries** and **smoother plugin/workflow discovery**, while the maintainer team is delivering fixes and features at a high velocity.

## 8. Backlog Watch

| Issue | Age | Status | Attention Needed |
|-------|-----|--------|-----------------|
| [#5903](https://github.com/zeroclaw-labs/zeroclaw/issues/5903) MCP orphan leak | Opened Apr 19 (~66 days) | Accepted, P1, no‑stale | Acknowledged but no fix PR; maintainers should prioritise. |
| [#6289](https://github.com/zeroclaw-labs/zeroclaw/issues/6289) Skill suggestions | Opened May 2 (~53 days) | Accepted, no‑stale, P2 | No progress beyond discussion; requires design. |
| [#6140](https://github.com/zeroclaw-labs/zeroclaw/issues/6140) Hybrid skills + WASM | Opened Apr 26 (~59 days) | Accepted, no‑stale, P2 | Blocked until markdown‑only plugins land. |
| [#6943](https://github.com/zeroclaw-labs/zeroclaw/issues/6943) Plugin system deconflict | Opened May 26 (~29 days) | Accepted, P2 | RFC accepted; implementation dependent on Phase 2 decisions. |
| [#7733](https://github.com/zeroclaw-labs/zeroclaw/issues/7733) mcp_bundles not enforced | Opened Jun 15 (~9 days) | Accepted, P1 | Critical security gap; no fix PR yet. |
| [#8177](https://github.com/zeroclaw-labs/zeroclaw/issues/8177) Supply‑chain signing RFC | Opened Jun 22 (~2 days) | Needs‑maintainer‑review | PRs are in progress; maintainer review needed for final approval. |

The oldest high‑severity item is **#5903** (MCP orphan leak), which has been open for over two months and affects all daemon users with MCP stdio servers. Despite the P1 tag and no‑stale status, no corresponding fix PR has been created. This warrants maintainer attention alongside the newer but equally critical #7733.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*