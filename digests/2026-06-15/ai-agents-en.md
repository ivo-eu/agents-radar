# OpenClaw Ecosystem Digest 2026-06-15

> Issues: 260 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-15 03:43 UTC

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

# OpenClaw Project Digest — 2026-06-15

## 1. Today's Overview
Activity remains extremely high with **260 issues** and **500 pull requests** updated in the last 24 hours. The vast majority (232 issues, 425 PRs) are still open, signaling both an active contributor base and a growing maintenance backlog. One new release, **v2026.6.8-beta.1**, landed today, bringing richer Telegram and WhatsApp channel delivery. However, the project continues to wrestle with a cluster of P1 regressions concentrated around session state, message duplication, and memory leaks — many of which are still awaiting maintainer review or product decisions.

## 2. Releases
### [v2026.6.8-beta.1](https://github.com/openclaw/openclaw/releases/tag/v2026.6.8-beta.1)
- **Highlights:** Telegram structured rich text (tables, lists, expandable blockquotes) and safer rich-media boundaries; WhatsApp delivery improved with prompt-preserving CLI backend migration and retired native draft handling. The original changelog entry truncates mid-sentence (ending with “WhatsAp”), suggesting the full release notes may be incomplete.
- **Breaking changes / migration notes:** None documented. Operators upgrading from v2026.5.x should review Telegram and WhatsApp plugin configs for any new options.

## 3. Project Progress
**75 PRs were merged or closed** in the last 24 hours. Notable among the top-30-by-comment list:

| PR | Description | Status |
|----|-------------|--------|
| [#93131](https://github.com/openclaw/openclaw/pull/93131) | Data Egress Filter for `web_fetch` fallback – prevents credential leakage | **Closed** |
| [#93133](https://github.com/openclaw/openclaw/pull/93133) | Adds `zod-compiler` for 2–75× faster Zod schema validation | **Closed** |
| [#93117](https://github.com/openclaw/openclaw/pull/93117) | Fix thinking-block recovery retry after control-plane start event | Open, proof supplied |
| [#93148](https://github.com/openclaw/openclaw/pull/93148) | Record every sent message in Telegram SQLite database | Open |
| [#93137](https://github.com/openclaw/openclaw/pull/93137) | Fix iMessage sender to honor disabled reply actions | Open, ready for maintainer look |

**Closed issues** (28 total) include high-priority items such as [#86184](https://github.com/openclaw/openclaw/issues/86184) (Telegram direct get generic `/new` fallback), [#86231](https://github.com/openclaw/openclaw/issues/86231) (Codex harness prompt latency tracking), and [#85692](https://github.com/openclaw/openclaw/issues/85692) (Feishu agent silent failure). These closures indicate active bug-fixing effort, though the sheer volume of open items remains concerning.

## 4. Community Hot Topics

### Most Commented Issues
- **[#85888](https://github.com/openclaw/openclaw/issues/85888)** (12 comments) – Cron jobs consistently fail with MiniMax 503 overload during early morning CST; manual triggers succeed. The community suspects a scheduling/retry logic issue rather than API availability.  
- **[#86996](https://github.com/openclaw/openclaw/issues/86996)** (9 comments) – Active Memory + Codex app-server path causes long latency, hook timeouts, and gateway event-loop stalls. Requires product decision and live repro.  
- **[#86519](https://github.com/openclaw/openclaw/issues/86519)** (9 comments) – Agent repeats identical replies 2–10× on Telegram after 5.20 update. Regression not fully resolved in 5.22.  
- **[#86508](https://github.com/openclaw/openclaw/issues/86508)** (9 comments, 4 👍) – `EmbeddedAttemptSessionTakeoverError` during Discord runs: session file changed while lock released. High community pain point.

### Most Reacted Issues
- [#86508](https://github.com/openclaw/openclaw/issues/86508) – 4 👍 (session takeover)  
- [#86047](https://github.com/openclaw/openclaw/issues/86047) – 3 👍 (Codex app-server stalls)  
- [#86845](https://github.com/openclaw/openclaw/issues/86845) – 2 👍 (session takeover cluster)  
- [#86616](https://github.com/openclaw/openclaw/issues/86616) – 2 👍 (voice flow regression + downgrade blocked)  

**Underlying needs:** The community is frustrated by repeated session-state corruption, message duplication, and silent cron failures — all of which erode trust in the gateway’s reliability. Many of these are regressions from the 5.20/5.22 releases, and users are eager for root-cause fixes rather than workarounds.

## 5. Bugs & Stability

### P1 Regressions (Severity: Critical)
| Issue | Description | Fix PR Exists? |
|-------|-------------|----------------|
| [#86996](https://github.com/openclaw/openclaw/issues/86996) | Active Memory + Codex: long latency, hook timeouts, event-loop stalls | No |
| [#86519](https://github.com/openclaw/openclaw/issues/86519) | Duplicate replies on Telegram (2–10×) after 5.20 | No |
| [#86508](https://github.com/openclaw/openclaw/issues/86508) | `EmbeddedAttemptSessionTakeoverError` in Discord runs | No |
| [#86047](https://github.com/openclaw/openclaw/issues/86047) | Codex app-server / plugin approval stalls in Nextcloud Talk | No |
| [#85822](https://github.com/openclaw/openclaw/issues/85822) | Silent ~48s lane retention after `embedded run done` | [Yes](https://github.com/openclaw/openclaw/pull/93110) (PR #93110) |
| [#87109](https://github.com/openclaw/openclaw/issues/87109) | Gateway heap grows to 1073MB+ at idle, cron fails silently | No |
| [#87068](https://github.com/openclaw/openclaw/issues/87068) | Telegram duplicate outbound sends (2–4× per turn) | No |
| [#86616](https://github.com/openclaw/openclaw/issues/86616) | `nativeHook.invoke` regression breaks Telegram voice + blocks downgrade | No |
| [#86592](https://github.com/openclaw/openclaw/issues/86592) | Inbound user messages not persisted to session JSONL on agent throw | [Yes](https://github.com/openclaw/openclaw/pull/93138) (PR #93138) |
| [#86845](https://github.com/openclaw/openclaw/issues/86845) | Cluster of 13 `EmbeddedAttemptSessionTakeoverError` events in 42h | No |

**Additional P2 issues** (e.g., [#86827](https://github.com/openclaw/openclaw/issues/86827) – group chat stuck in failed state, [#86215](https://github.com/openclaw/openclaw/issues/86215) – Codex OAuth refresh wedging agents) are also open with no fix PRs yet.

**Stability assessment:** The project is facing a **regression crisis** in its 5.20–5.22 lineage, with at least a dozen high-severity bugs that directly impact end users. While some fixes are in-flight (e.g., PR #93110 for cron delivery leases, PR #93138 for session persistence), the majority remain in triage limbo. The release of v2026.6.8-beta.1 may introduce new features but does not address these core stability issues.

## 6. Feature Requests & Roadmap Signals

### Top User-Requested Features
- **[#86881](https://github.com/openclaw/openclaw/issues/86881)** – Gateway-lite mode without AI harness for deterministic deployments (P2, 7 comments). This would allow lightweight gateway use for webhooks, cron, and plugins.
- **[#86434](https://github.com/openclaw/openclaw/issues/86434)** – ElevenLabs Realtime Voice Provider for Talk (P3, 3 comments).
- **[#86425](https://github.com/openclaw/openclaw/issues/86425)** – `describe_view` camera frame support for OpenAI Realtime Talk (P2, 2 👍).
- **[#86534](https://github.com/openclaw/openclaw/issues/86534)** – TUI competitive analysis: 15 gaps vs Claude Code, Aider, etc. (P3, 3 comments).
- **[#87002](https://github.com/openclaw/openclaw/issues/87002)** – Telegram DM debounce / message aggregation window (P3, 3 comments).
- **[#86237](https://github.com/openclaw/openclaw/issues/86237)** – Rename internal `cron` subsystem to avoid collision with system cron (P3, 4 comments).

### In-Flight Features (PRs)
- **[#89569](https://github.com/openclaw/openclaw/pull/89569)** – Pre-auth access requests and grouped DM allowlists for Telegram/WhatsApp (size XL, waiting on author).
- **[#91476](https://github.com/openclaw/openclaw/pull/91476)** – Ultracode backend flag for Claude CLI (ready for maintainer look).

**Prediction for next minor release:** The gateway-lite mode (#86881) and ElevenLabs voice (#86434) have strong community support but are P2/P3. The **cron rename** (#86237) is a simpler change and may land soon. Given the regression load, the project may prioritize bug fixes over new features in the immediate next patch.

## 7. User Feedback Summary

**Pain points (most frequently voiced):**
- **Message duplication:** Telegram users see identical replies 2–10× after 5.20 ([#86519](https://github.com/openclaw/openclaw/issues/86519), [#87068](https://github.com/openclaw/openclaw/issues/87068)).
- **Session takeover errors:** Cron jobs and Discord runs fail with `EmbeddedAttemptSessionTakeoverError` ([#86508](https://github.com/openclaw/openclaw/issues/86508), [#86845](https://github.com/openclaw/openclaw/issues/86845)).
- **Silent failures:** Cron jobs that produce no output or push notification ([#87109](https://github.com/openclaw/openclaw/issues/87109), [#86827](https://github.com/openclaw/openclaw/issues/86827)).
- **Memory/resource leaks:** Gateway heap grows to over 1GB at idle on macOS/Docker ([#87109](https://github.com/openclaw/openclaw/issues/87109)).
- **Regression risk:** Multiple users report downgrading to 5.12 as the only stable version ([#86047](https://github.com/openclaw/openclaw/issues/86047), [#86616](https://github.com/openclaw/openclaw/issues/86616)).

**Positive signals:** The new release’s improvements to Telegram and WhatsApp formatting were well received by early testers. The community also appreciates rapid issue filing and the wealth of debugging information provided in reports.

**Satisfaction:** Mixed. While the project’s feature set is praised (“key differentiator is its multi-agent orchestration”), reliability concerns are eroding trust. Users are watching for a stable release that resolves the regression cluster.

## 8. Backlog Watch

### Critical Issues Needing Maintainer Attention
| Issue | Age (days) | Status | Reason |
|-------|------------|--------|--------|
| [#86996](https://github.com/openclaw/openclaw/issues/86996) | 20 | Open, P1 | Needs maintainer review & product decision; no fix PR |
| [#86519](https://github.com/openclaw/openclaw/issues/86519) | 21 | Open, P1 | Needs maintainer review; no fix PR |
| [#86047](https://github.com/openclaw/openclaw/issues/86047) | 22 | Open, P1 | Needs maintainer review & product decision; no fix PR |
| [#86215](https://github.com/openclaw/openclaw/issues/86215) | 22 | Open, P1 | Needs product decision & live repro; no fix PR |
| [#86214](https://github.com/openclaw/openclaw/issues/86214) | 22 | Open, P1 | Needs product decision & live repro; no fix PR |
| [#87109](https://github.com/openclaw/openclaw/issues/87109) | 19 | Open, P1 | Needs maintainer review & product decision & live repro |

### PRs Waiting on Author
- [#89569](https://github.com/openclaw/openclaw/pull/89569) – Pre-auth access requests (waiting on author since June 2).
- [#93141](https://github.com/openclaw/openclaw/pull/93141) – Feishu channel crash fix (waiting on author).
- [#87260](https://github.com/openclaw/openclaw/pull/87260) – Secure multi-agent routing docs (waiting on author since May 27).

### Long-Unanswered Issues
- [#85461](https://github.com/openclaw/openclaw/issues/85461) (P2, created May 22) – Capture image-generation provider usage metadata. No maintainer response in 24 days.

These items represent a growing backlog that, if unaddressed, may further strain community confidence and increase the number of stale issues.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report — AI Agent Open-Source Ecosystem

## 1. Ecosystem Overview

The personal AI assistant and agent open-source landscape in June 2026 is characterized by explosive development activity across a dozen projects, but with a widening gap between feature velocity and production reliability. Most projects are in a rapid iteration phase, shipping new capabilities (multi-agent orchestration, provider switching, WebUI parity) while simultaneously accumulating regressions—particularly around session state, message duplication, and credential security. The two largest communities (OpenClaw and ZeroClaw) are dealing with critical stability backlogs, while smaller projects like NanoBot and IronClaw maintain a healthier fix-to-feature ratio. A clear architectural split is emerging between "gateway-first" projects (OpenClaw, Hermes, NanoClaw) that treat the AI agent as a service endpoint and "workbench-first" projects (IronClaw, CoPaw) that emphasize a rich desktop/web UI for interactive debugging.

## 2. Activity Comparison

| Project | Issues Updated (24h) | PRs Updated (24h) | New Release? | Overall Health | Notes |
|--------|----------------------|-------------------|--------------|----------------|-------|
| OpenClaw | 260 (232 open) | 500 (425 open) | Yes (beta) | **C** – Regression crisis; 10+ P1 bugs unaddressed |
| NanoBot | 5 (2 open) | 46 (19 open) | No | **A** – Healthy throughput, fast fix turnaround |
| Hermes Agent | 14 (12 open) | 50 (34 open) | No | **B** – Security advisories under triage; active fixes |
| PicoClaw | 5 (3 open) | 9 (4 open) | Nightly | **B** – Steady, but silent tool failures |
| NanoClaw | 7 (6 open) | 10 (5 open) | No | **B** – Security issues filed, budget-bug fix in review |
| NullClaw | 1 (1 open) | 0 | No | **D** – Minimal activity, single feature request |
| IronClaw | 22 (17 open) | 44 (29 open) | No | **B+** – High feature velocity, but release blocked |
| LobsterAI | 0 new issues | 1 merged PR | No | **C+** – Low churn, stale PRs since April |
| CoPaw | 21 (17 open) | 16 (11 open) | No | **C** – Two critical regressions in v1.1.11.post2 |
| ZeroClaw | 3 (1 open) | 50 (46 open) | No | **B–** – Intense but high-risk open PRs, air-gapped RFC stalled |
| Moltis | 1 (0 open) | 0 | No | **D** – Effectively dormant; single enhancement |
| TinyClaw | 0 | 0 | No | **D** – No activity |
| ZeptoClaw | 0 | 0 | No | **D** – No activity |

## 3. OpenClaw's Position

OpenClaw remains the **largest and most feature-rich** project in the ecosystem, with 260 issues and 500 PRs in a single day—far exceeding any peer. It is the de facto reference implementation for gateway-based agent orchestration, supporting the widest range of messaging channels (Telegram, WhatsApp, Discord, Feishu, etc.). Its technical advantage lies in the **unified session and cron engine**, which enables cross-channel continuity. However, this complexity is also its Achilles’ heel: the 5.20–5.22 lineage has introduced a **cluster of P1 regressions** (session takeover errors, duplicate messages, memory leaks to 1GB+) that are eroding user trust. Competitors like **NanoBot** and **IronClaw** are closing feature gaps with more stable release cycles. OpenClaw’s community engagement is unmatched (75 PRs merged in 24h), but its **maintainer review bottleneck** (6 P1 issues >20 days old without a fix PR) is becoming a structural risk. For adopters prioritizing reliability over bleeding-edge features, OpenClaw currently lags behind NanoBot and Hermes Agent.

## 4. Shared Technical Focus Areas

The following cross-project patterns indicate where the ecosystem as a whole is investing:

| Focus Area | Affected Projects | Specific Gaps |
|------------|------------------|---------------|
| **Session state & duplication** | OpenClaw, NanoClaw, CoPaw, ZeroClaw | Repeated reply delivery, session takeover errors, state file corruption |
| **Credential & secret management** | Hermes Agent, NanoClaw, ZeroClaw | File-path leaks in image fallback, `#[secret]` not redacting bearer tokens, auth JSON exfiltration |
| **Cron & scheduling reliability** | OpenClaw, ZeroClaw, LobsterAI | Silent failures after restart, pause/resume missing, MiniMax 503 not retried |
| **Provider integration friction** | Hermes Agent, CoPaw, NullClaw, NanoBot | OAuth vs API-key double billing, temperature deprecation, Claude SDK subscription, Azure identity-based auth |
| **WebUI–config parity** | NanoBot, IronClaw, ZeroClaw, CoPaw | Admin controls missing in UI, settings not reflected in config.json |
| **Cross-platform support** | Hermes Agent, CoPaw, ZeroClaw, PicoClaw | Intel macOS builds, Windows AMQP compilation, iOS Safari <16.4, Wayland broken |
| **Long-running session robustness** | OpenClaw, IronClaw, CoPaw | Active Memory + Codex latency, timeout loops, context compression discarding all state |

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | Hermes Agent | IronClaw | CoPaw | ZeroClaw |
|-----------|----------|---------|--------------|----------|-------|----------|
| **Primary paradigm** | Gateway-first | Gateway + WebUI | Gateway + TUI | Agent workbench + WebUI | Workbench + plugins | Rust core + config-driven |
| **Target user** | Power users, multi-channel | Developers, self-hosters | Security-conscious, terminal users | Enterprise, observability | Desktop users, Chinese market | Low-resource, air-gapped |
| **Architecture language** | TypeScript/Node | Python | TypeScript/Rust | Rust + Tauri | Python | Rust |
| **Unique strength** | Channel breadth (10+ platforms) | Fast schema validation (zod-compiler) | Desktop app + security audit | Observability seams, dogfooding | Plugin system, Vietnamese community | Type-driven config, cron pause/resume |
| **Current weakness** | Regression density | Zero usage-token bug | Intel macOS missing; memory-context semantics | Release blocked 30+ days | Two critical regressions in v1.1.11.post2 | High-risk open PRs, straw man RFCs |
| **Release stability** | Beta regressions | Stable | Stable with P1 advisories | Feature-rich but stuck | Regression-prone | Mixed |

## 6. Community Momentum & Maturity

**Tier 1 – Rapidly iterating (high throughput, high risk):**  
- **OpenClaw** – 500+ PRs/day, but drowning in regressions. Community trust is at risk if the regression crisis is not resolved within 2–3 weeks.  
- **ZeroClaw** – 50 PRs/day, config framework overhaul merged, but 6 out of top-10 open PRs need author action. High potential if maintainers clear the bottleneck.  
- **IronClaw** – 44 PRs/day, observability and extensibility focused. Release PR blocked for 30 days is the main drag.  

**Tier 2 – Healthy growth (moderate throughput, stable):**  
- **NanoBot** – Low bug count, rapid fix turnaround, WebUI–config parity almost complete. Best overall health score.  
- **Hermes Agent** – Security hardening phase; 2 P1 advisories under triage but fix PRs already exist for credential exfiltration. Steady.  
- **NanoClaw** – Feature-advancement mode; multi-provider infrastructure merged. Security issues need maintainer response soon.  

**Tier 3 – Stabilizing or stalled:**  
- **PicoClaw** – Nightly builds, 5 merged PRs today (error handling). Silent tool failures need attention but overall steady.  
- **LobsterAI** – Low churn; three feature PRs languishing since April. Needs maintainer push to merge or close.  
- **CoPaw** – High community activity but two critical regressions in the latest release. Must ship v1.1.12 soon to avoid user exodus.  
- **NullClaw, Moltis, TinyClaw, ZeptoClaw** – Minimal or no activity. Effectively stalled.

## 7. Trend Signals

1. **Security is the new differentiator.** Three projects had security advisories filed in the last 24h (Hermes, NanoClaw, ZeroClaw) – all involving credential exfiltration or authorization bypass. Users are demanding audit trails, secret redaction, and air-gapped execution (ZeroClaw RFC #6293). The ecosystem is moving from "it works" to "it's secure."

2. **Provider billing friction drives feature requests.** Across Hermes (Claude SDK OAuth), CoPaw (Kimi subscription), and NullClaw (Azure identity-based auth), users are pushing for **bring-your-own-subscription** models that avoid double-billing API keys. This will likely become a baseline requirement for any agent framework targeting paid AI services.

3. **Long-running session resilience is a universal pain point.** Every major project reported issues with session state corruption, memory leaks, context loss, or timeout loops. As agents graduate from chat to background workflows (cron, knowledge generation), the community needs **transactional session management** with rollback and checkpointing.

4. **WebUI parity is table stakes.** NanoBot, IronClaw, and ZeroClaw are all adding admin controls (pause/resume cron, temperature sliders, tool toggles) to their dashboards. The CLI-only era is ending; users expect full graphical management of agent configurations.

5. **Cross-platform expansion is accelerating.** Hermes Agent’s Intel macOS gap, CoPaw’s Windows Tauri slow start, and ZeroClaw’s AMQP compile issue on Windows show that Windows and Linux desktop support are now critical for adoption beyond developer laptops. Vietnamese community growth (CoPaw) signals increasing internationalization demand.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-06-15

## 1. Today’s Overview
NanoBot saw high activity over the past 24 hours with **5 issues updated** (2 open, 3 closed) and **46 pull requests updated** (19 open, 27 merged/closed). No new releases were cut. The project continues to focus on stability fixes, WebUI–config parity, and refactoring the agent loop boundaries. Significant merged PRs addressed token usage tracking, session scoping, config validation, and mobile responsiveness, while two new open bugs (zero token usage in the `/v1/chat/completions` endpoint and a file-path leak in the image-strip fallback) surfaced from the community. The overall pace indicates healthy maintenance and active feature development.

## 2. Releases
*No new releases.* No version bumps, changelogs, or migration notes are available.

## 3. Project Progress
The following notable PRs were merged or closed (from the top 20 listed, 17 of 20 are closed):
- **Token usage heatmap & timezone fix** ([PR #4248](https://github.com/HKUDS/nanobot/pull/4248)) — aligns date window with agent timezone, prevents label clipping.
- **Beginner-friendly documentation** ([PR #4177](https://github.com/HKUDS/nanobot/pull/4177)) — restructured entry points, added CLI chooser, config task map.
- **No-tools finalization when max iterations reached** ([PR #4269](https://github.com/HKUDS/nanobot/pull/4269)) — provides concise status instead of generic budget message.
- **`pathPrepend` for exec tool** ([PR #4273](https://github.com/HKUDS/nanobot/pull/4273)) — new config field to give tool directories PATH lookup precedence.
- **Fail fast on invalid config files** ([PR #4275](https://github.com/HKUDS/nanobot/pull/4275)) — improved error handling for malformed configs.
- **Session-scoped recent history** ([PR #4274](https://github.com/HKUDS/nanobot/pull/4274)) — filters `# Recent History` prompt by session in non-unified mode, excludes cron/dream/heartbeat in unified mode.
- **Feishu channel lazy-load fix** ([PR #4277](https://github.com/HKUDS/nanobot/pull/4277)) — defer `lark_oapi` import to avoid WebSocket loop conflicts.
- **Cron session binding** ([PR #4299](https://github.com/HKUDS/nanobot/pull/4299)) — cron jobs now own the creating session instead of defaulting to `unified:default`.
- **Config schema import cycle resolution** ([PR #4314](https://github.com/HKUDS/nanobot/pull/4314)) — moved shared Pydantic base into `nanobot.config_base`, breaking circular imports.
- **Mobile WebUI responsiveness** ([PR #4339](https://github.com/HKUDS/nanobot/pull/4339)) — tightened spacing, prevented overflow, improved heatmap on small screens.
- **Built-in filesystem tool toggle** ([PR #4138](https://github.com/HKUDS/nanobot/pull/4138)) — added `tools.file.enable` flag, aligning with exec/web groups.
- **Desktop restart token & replay gaps** ([PR #4210](https://github.com/HKUDS/nanobot/pull/4210)) — refreshed WebUI tokens after engine restart, persisted WebSocket events for replay.
- **Telegram code-block split fix** ([#4250](https://github.com/HKUDS/nanobot/pull/4250), closed) — prevents `split_message` from breaking fenced code blocks across chunks.
- **Anthropic `temperature` deprecation** ([#4333](https://github.com/HKUDS/nanobot/pull/4333), closed) — suppressed `temperature` for `opus-4-8`/Fable models.

## 4. Community Hot Topics
- **Zero usage tokens in `/v1/chat/completions`** ([#4309](https://github.com/HKUDS/nanobot/issues/4309)) — the open issue reports that every response returns hardcoded `"usage": {"prompt_tokens": 0, ...}`. The reporter notes the agent loop already tracks real usage. This is a high-visibility regression for any tool consuming the OpenAI-compatible API, including billing dashboards. No fix PR exists yet, but the issue has one comment (likely a maintainer acknowledgment).
- **Image-strip fallback leaks file path** ([#4345](https://github.com/HKUDS/nanobot/issues/4345)) — filed today, describes that when image stripping occurs, the injected text contains the original file path, creating a privacy leak and causing the model to hallucinate “seeing” an image. A companion fix PR ([#4346](https://github.com/HKUDS/nanobot/pull/4346)) was opened simultaneously, suggesting strong community/maintainer alignment.

*Other open PRs* with potential for discussion:
- **WebUI/config.json parity** ([#4313](https://github.com/HKUDS/nanobot/pull/4313)) — large feature PR adding write endpoints for temperature, tool limits, dream, channels, memory, and new UI controls. Still open, likely under review.
- **Automation management view** ([#4330](https://github.com/HKUDS/nanobot/pull/4330)) — adds a WebUI surface for listing/filtering/running/pausing automations. Open, awaiting merge.

## 5. Bugs & Stability
| Severity | Bug | Status | Fix PR |
|----------|-----|--------|--------|
| **High** | `/v1/chat/completions` always returns zero usage tokens ([#4309](https://github.com/HKUDS/nanobot/issues/4309)) | Open, no PR | — |
| **High** | Image-strip fallback leaks file path and misleads model ([#4345](https://github.com/HKUDS/nanobot/issues/4345)) | Open, fix PR exists | [#4346](https://github.com/HKUDS/nanobot/pull/4346) |
| **Medium** | Anthropic provider sends deprecated `temperature` to opus-4-8 / Fable (400 error) | Closed by [#4333](https://github.com/HKUDS/nanobot/pull/4333) | Merged |
| **Medium** | Telegram `split_message` breaks fenced code blocks | Closed by [#4250](https://github.com/HKUDS/nanobot/pull/4250) | Merged |
| **Low** | `tools.exec.pathPrepend` missing (now addressed) | Closed by [#4273](https://github.com/HKUDS/nanobot/pull/4273) | Merged |
| **Low** | Config file validation not failing fast on malformed YAML | Closed by [#4275](https://github.com/HKUDS/nanobot/pull/4275) | Merged |

**Critical open bugs:** #4309 (usage stats) and #4345 (file-path leak). The latter already has a proposed fix; the former still needs attention from the maintainers.

## 6. Feature Requests & Roadmap Signals
- **BotIcon on agent start** ([#4262](https://github.com/HKUDS/nanobot/issues/4262), closed) — user requested that `botIcon` from `config.json` be displayed immediately instead of a default puppy; merged.
- **Filesystem tool toggle** ([PR #4138](https://github.com/HKUDS/nanobot/pull/4138)) — now available, matching exec/web enable patterns.
- **WebUI automations manager** ([PR #4330](https://github.com/HKUDS/nanobot/pull/4330)) — in review, likely for next minor release.
- **Config/WebUI settings parity** ([PR #4313](https://github.com/HKUDS/nanobot/pull/4313)) — large surface area, may land in multiple patches.
- **PathPrepend for exec tool** ([PR #4273](https://github.com/HKUDS/nanobot/pull/4273)) — merged, ready for users needing custom PATH overrides.
- **Strict tool parameter validation** ([#4343](https://github.com/HKUDS/nanobot/pull/4343), open) — would reject unknown parameters before execution, improving debugging.

**Prediction:** The next release will likely include the automation management view, WebUI settings parity for temperature/tool limits, and the image-strip fallback fix. The zero-usage-token bug is a strong candidate for a hotfix.

## 7. User Feedback Summary
- **Pain points:**
  - *Zero usage tokens* — breaks compatibility with tools that rely on token counts (e.g., cost monitoring).
  - *Image fallback leaking file paths* — privacy concern and model hallucination.
  - *Anthropic temperature deprecation* (now fixed) — users of newer Claude models couldn’t use NanoBot.
  - *Telegram code-block rendering* (now fixed) — broke multi-message responses.
  - *Mobile WebUI overflow* (addressed in PR #4339) — improved usability on small screens.
- **Satisfaction signals:** Users contributed enhancements (botIcon startup, pathPrepend), and the community is actively filing detailed bug reports with reproduction steps and even providing fix PRs (e.g., #4346). This indicates a healthy, engaged user base.

## 8. Backlog Watch
No issues or PRs older than a week remain unanswered. The following open items might benefit from maintainer attention:
- **PR #4293** (open, subagent result injection) — still unmerged; could affect cron + subagent workflows.
- **PR #4343** (open, reject unknown builtin parameters) — small but important for tool safety.
- **Issue #4309** (zero usage tokens) — no fix PR yet, needs maintainer investigation and assignment.

Overall project health is strong: rapid PR throughput, responsive issue triage (3 closed issues in 24h), and a clear direction toward WebUI–config parity and agent loop cleanup.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-06-15

## 1. Today's Overview
Project activity is **high** today with 50 pull requests updated and 14 issues refreshed in the last 24 hours. Of the 10 merged or closed PRs, several deliver critical fixes for credential‑store exfiltration, Copilot leakage, and Telegram rich message support. Two new **P1 security advisories** (email gateway spoofing and unreviewed plugin execution) have been opened and are under triage. The community is particularly engaged around a Claude SDK subscription‑based provider and a request to remove provider accounts from the TUI. While no new release was cut, the volume of security‑related fixes and feature PRs suggests a stability and hardening phase is underway.

## 2. Releases
No new releases were published in the last 24 hours.

## 3. Project Progress — Merged / Closed PRs
Ten PRs were merged or closed today. Notable among them:

- **[#39840] fix(ui‑tui): stabilize embedded dashboard chat gateway** (closed) – Fixed a crash when the dashboard embedded the TUI gateway while the React hook tree was still rendering.
- **[#46399] fix(terminal): fall back to .env for env_passthrough in local backend** (closed) – Mirrors the Docker backend’s behaviour: `terminal.env_passthrough` now reads variables from `~/.hermes/.env` even when not exported to `os.environ`.
- **[#46438] fix(send_message): support Telegram rich messages** (closed) – Adds an opt‑in fast path to send raw Markdown for tables, task lists, etc. via Telegram’s Bot API 10.1.

Other closed items include duplicate/bug‑fix PRs for file permissions, TTS on Feishu, and test issues.

## 4. Community Hot Topics
The most active issues and PRs, ranked by comments and reactions:

- **[#25267] [Feature]: Claude Agent SDK model provider with subscription OAuth**  
  **20 👍, 6 comments** – Strong demand for a `provider/anthropic` that uses a Claude subscription OAuth (like Codex) instead of requiring a separate API key. Users want to avoid double billing.  
  📎 [NousResearch/hermes-agent Issue #25267](https://github.com/NousResearch/hermes-agent/issues/25267)

- **[#31584] [Feature]: Treat memory‑context as background context, not authoritative user‑message content**  
  **0 👍, 6 comments** – A nuanced request to stop agents from treating recalled memories as hard user instructions, both to reduce confusion and to close a potential attack surface.  
  📎 [NousResearch/hermes-agent Issue #31584](https://github.com/NousResearch/hermes-agent/issues/31584)

- **[#45865 / #46445] [Feature]: Add ability to remove provider accounts**  
  **0 👍, 4 comments (closed), new issue resurfaced** – The original feature was closed but the user reports that the trash icon now appears only for Nous Portal accounts, not for Anthropic OAuth. The follow‑up (#46445) asks for consistent UI.  
  📎 [NousResearch/hermes-agent Issue #45865](https://github.com/NousResearch/hermes-agent/issues/45865)  
  📎 [NousResearch/hermes-agent Issue #46445](https://github.com/NousResearch/hermes-agent/issues/46445)

- **[#42199] Request: x86_64 (Intel) macOS build for Desktop App**  
  **0 👍, 4 comments** – Intel Mac users are blocked because the DMG ships ARM‑only. Rosetta 2 cannot run Intel code.  
  📎 [NousResearch/hermes-agent Issue #42199](https://github.com/NousResearch/hermes-agent/issues/42199)

## 5. Bugs & Stability
Several high‑severity bugs were reported or addressed today:

| Severity | Issue | Description | Fix PR? |
|----------|-------|-------------|---------|
| **P1** | [#46434] | **Authorization Bypass via Spoofed `From:` Header in Email Gateway** (private advisory GHSA‑rxqh‑5572‑8m77) – under triage. | Not yet public |
| **P1** | [#46435] | **Unreviewed Plugin Code Execution via Dashboard Loader** (GHSA‑mcfc‑hp25‑cjv7) – `plugins.enabled` gate missing. | Not yet public |
| **P1** | [#32091] | **Cron jobs created in profile‑scoped sessions become silent orphans** – gateway never reads profile‑local `cron/jobs.json`. No fix PR yet. | — |
| **P2** | [#46411] | **read_file can exfiltrate credential stores from sibling profiles** – `get_read_block_error()` does not block `<root>/profiles/<other>/auth.json`. | [#46417] (open) |
| **P2** | [#46413] | **Desktop file preview can read Hermes credential stores** – `sensitiveFileBlockReason()` returns `null` for Hermes token stores. | [#46430] (open) |
| **P2** | [#46421] | **Corrupt auth.json can be silently overwritten** – fix preserves the corrupt file and refuses to overwrite. | [#46421] (open – fix itself) |
| **P3** | [#46424] | **Hindsight plugin hardcodes UTC timestamps** – memory extraction loses local timezone. | Not yet |

The project is responding quickly: PRs #46417 and #46430 directly address the credential exfiltration bugs.

## 6. Feature Requests & Roadmap Signals
High‑impact feature requests that may influence the next release:

- **Claude subscription OAuth provider** (#25267) – Would enable users to bring their own Claude subscription without incurring additional API costs. This is the most upvoted issue today and could become a high‑priority roadmap item.
- **Memory‑context as background, not authoritative** (#31584) – Addresses a systemic agent behaviour problem that affects both accuracy and security. Likely to be considered for a near‑future release.
- **Consistent “remove provider account” UI** (#45865 / #46445) – The partial implementation frustrates users. A full solution may land in the next TUI update.
- **Intel macOS build** (#42199) – Blocks a significant user segment; infrastructure work may be planned.
- **Desktop “deliverable mode” with auto‑attachment** (#46444) – Users want generated files (ZIP, PDF) to be automatically attached in the chat, mirroring gateway platforms.
- **Kanban epoch callback for spiral workflows** (PR #46360) – A feature PR adding orchestrator epoch callbacks for autonomous measure‑dispatch‑wait cycles.

## 7. User Feedback Summary
Real pain points and satisfaction signals from today’s activity:

- **Billing friction** – Several users express frustration about double‑paying (subscription + API key) when using Claude with Hermes. The request for OAuth‑based provider is a direct response.
- **Memory confusion** – A user (who admits the details exceed their comprehension) reports that agents treat recalled memory as authoritative user input, causing erratic behaviour and potential misuse.
- **UI incompleteness** – The ability to add but not remove provider accounts is a clear UX gap, now reported twice.
- **macOS exclusion** – Intel Mac users feel left out when only ARM64 builds are provided.
- **Triage decomposition** – A feature PR (#46443) to inherit notification subscriptions to child tasks addresses a workflow need for Kanban users.
- **Security awareness** – The two new P1 advisories indicate that the community and maintainers are actively scanning for vulnerabilities; users reporting these issues are contributing to project health.

## 8. Backlog Watch
Issues and PRs that have been open for extended periods without maintainer resolution:

- **[#25267] Claude SDK OAuth provider** – Open since **2026-05-13** (33 days). Although highly upvoted, no maintainer response or PR is visible. The community is waiting for a signal.
- **[#31584] Memory‑context semantics** – Open since **2026-05-24** (22 days). Requires a deeper architectural discussion; no maintainer reply yet.
- **[#32091] Cron orphan in profile‑scoped sessions** – Open since **2026-05-25** (21 days). This is a **P1** bug with a clear reproducer and no associated fix PR. It silently loses user jobs.
- **[#42199] Intel macOS build** – Open since **2026-06-08** (7 days). While newer, it has no maintainer acknowledgment or workaround note.

These items would benefit from a maintainer triage comment to set expectations or provide a timeline.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest – 2026-06-15

## 1. Today's Overview

Project activity is moderate with a nightly build cutting and five pull requests merged. Five issues and nine PRs were touched in the last 24 hours. The nightly release (`v0.2.9-nightly.20260615`) brings no documented changes beyond the automated build, which is flagged as potentially unstable. Several bug reports and feature-oriented pull requests indicate ongoing community interest in extensibility and cross-platform compatibility. Overall project health appears steady, though a handful of unresolved bugs and long‑standing PRs require maintainer attention.

## 2. Releases

**One new release:** [`nightly`](https://github.com/sipeed/picoclaw/releases/tag/main) (nightly build for `v0.2.9-nightly.20260615.13a38bd1`).  
This is an automated build and may be unstable.  
**Full changelog:** [compare/v0.2.9...main](https://github.com/sipeed/picoclaw/compare/v0.2.9...main)  
No migration notes or breaking changes are mentioned.

## 3. Project Progress

Five pull requests were merged/closed in the last 24 hours:

- [#2904](https://github.com/sipeed/picoclaw/pull/2904) **Fix agent loop reload and panic cleanup stability** — eliminates a detached goroutine that could cause blocked goroutines and improves cleanup reliability.
- [#3124](https://github.com/sipeed/picoclaw/pull/3124) **fix(tts): handle io.ReadAll error in error response path** — prevents silent degradation of diagnostic messages when the TTS API returns a non-200 status.
- [#3123](https://github.com/sipeed/picoclaw/pull/3123) **fix(filesystem): explicitly ignore Close() error on directory file descriptor** — makes the intentional error discard explicit, matching project style.
- [#3122](https://github.com/sipeed/picoclaw/pull/3122) **fix(evolution): capture Close() error on write file in appendJSONLRecords** — surfaces delayed write failures (e.g., disk full) that were previously ignored.
- [#3121](https://github.com/sipeed/picoclaw/pull/3121) **refactor(openai_compat): replace log.Printf with structured logger** — completes the migration to structured logging in the OpenAI compatibility module.

These closures improve stability, error reporting, and code consistency across the agent loop, TTS, filesystem, evolution, and provider layers.

## 4. Community Hot Topics

No single issue or PR received more than two comments. The most discussed items are:

- [#2978](https://github.com/sipeed/picoclaw/issues/2978) (closed, 2 comments) — **Request to add OmniRoute as a provider**. The user asked for integration or guidance on adding a combo to the configuration. This was closed as stale, but the underlying need for easier external provider integration remains.
- [#3044](https://github.com/sipeed/picoclaw/issues/3044) (1 comment) — **allow_from fails for Matrix user IDs containing a colon**. The user reports silent rejection of messages despite matching configuration. This points to a parsing or validation regression.
- [#3041](https://github.com/sipeed/picoclaw/issues/3041) (1 comment) — **`mcp add` mis-parses global flags into positionals**. The CLI command breaks when using `-t http` or `-f` flags, causing silent misnaming or failure.

The common thread is a desire for robust configuration and CLI tooling, with users encountering unexpected silent failures.

## 5. Bugs & Stability

Five bugs were updated in the last 24 hours, ranked by severity:

- **High: [#3125](https://github.com/sipeed/picoclaw/issues/3125) — `web_search` tool fails silently with Brave API key from `.security.yml`**.  
  After the API key migration to `.security.yml`, the tool returns `"No results for: [query]"` without any error. This breaks a core tool silently. No fix PR exists yet.

- **High: [#3044](https://github.com/sipeed/picoclaw/issues/3044) — `allow_from` fails for Matrix user IDs containing a colon**.  
  Standard Matrix IDs (`@localpart:domain`) are silently rejected. This blocks legitimate use of Matrix channels with access control. No fix PR is linked.

- **Medium: [#3041](https://github.com/sipeed/picoclaw/issues/3041) — `mcp add` mis-parses global flags into positionals**.  
  The command ignores flags like `--no-color` when followed by `mcp add -t http ...`. This affects all users of the MCP subcommand. No fix PR is linked.

- **Medium: [#3090](https://github.com/sipeed/picoclaw/issues/3090) — Panel does not work on Safari on iOS below 16.4**.  
  Users on older iOS devices cannot access the web panel after login. The issue affects a growing segment of mobile users.

- **Low: [#3123/#3122/#3124](https://github.com/sipeed/picoclaw/pull/3123) (already fixed)** — The error‑handling gaps in filesystem, evolution, and TTS modules have been closed in today’s merged PRs.

## 6. Feature Requests & Roadmap Signals

- **Out‑of‑tree channels** — [#3120](https://github.com/sipeed/picoclaw/pull/3120) adds a `RegisterChannelSettings` hook so third‑party channel implementations can integrate without forking PicoClaw. This is the top candidate for the next minor release.
- **Remote Pico WebSocket mode for the CLI agent** — [#3118](https://github.com/sipeed/picoclaw/pull/3118) enables `picoclaw agent --remote` to connect to a remote WebSocket endpoint. Likely to be included after review.
- **Telegram: treat reply to bot as mention** — [#2975](https://github.com/sipeed/picoclaw/pull/2975) (open since May 30) allows group chat reply triggers without explicit @mention. Awaiting merge.
- **OmniRoute provider** — [#2978](https://github.com/sipeed/picoclaw/issues/2978) was closed stale, but the request signals user demand for more provider options.
- **Web search durability** — [#3125](https://github.com/sipeed/picoclaw/issues/3125) suggests that the `.security.yml` migration broke tool functionality; a future release should address this regression.

## 7. User Feedback Summary

Real pain points expressed in today’s issues:

- **Silent tool failure**: The `web_search` tool returns no results without logging an error, making debugging impossible ([#3125](https://github.com/sipeed/picoclaw/issues/3125)).
- **Configuration frustration**: Using standard Matrix user IDs with `allow_from` fails silently ([#3044](https://github.com/sipeed/picoclaw/issues/3044)); CLI flags in `mcp add` are mis‑parsed ([#3041](https://github.com/sipeed/picoclaw/issues/3041)).
- **Mobile/web panel accessibility**: iOS users below 16.4 cannot use the panel at all ([#3090](https://github.com/sipeed/picoclaw/issues/3090)).
- **Positive signals**: The five merged PRs today were all contributed by external developers, indicating a healthy community of contributors who care about code quality.

## 8. Backlog Watch

Items that have been open for more than a week without maintainer response or fix:

- **[#2975](https://github.com/sipeed/picoclaw/pull/2975)** — PR: Telegram reply as mention. Opened May 30, no maintainer comments.
- **[#3044](https://github.com/sipeed/picoclaw/issues/3044)** — Bug: `allow_from` fails for Matrix colon IDs. Opened June 7, one user comment.
- **[#3041](https://github.com/sipeed/picoclaw/issues/3041)** — Bug: `mcp add` flag parsing. Opened June 7, one user comment.
- **[#3090](https://github.com/sipeed/picoclaw/issues/3090)** — Bug: Panel on Safari iOS <16.4. Opened June 10, one user comment.

These issues and PRs affect core functionality (Matrix channels, CLI, UI) and would benefit from maintainer triage or prioritization.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-15

## 1. Today's Overview

NanoClaw saw **high activity** over the past 24 hours, with 7 issues updated (6 open, 1 closed) and 10 pull requests updated (5 merged/closed, 5 still open). The project is clearly in a **feature-advancement and security hardening phase**: three coordinated security advisories were filed (issues #2760, #2761, #2762), a critical budget-silent-drop bug has a fix PR in review (#2759), and major feature branches for **Codex agent-provider v2** and **operator-driven provider selection** were merged. No new releases were cut, but the merged PRs point toward an imminent release containing multi-provider support and data-driven container builds.

## 2. Releases

**None.** No new releases were published in the last 24 hours.

## 3. Project Progress

**5 pull requests were merged or closed today:**

- **PR #2769** *(merged)* — Documentation fix for `/add-codex` skill: flags interactive auth step and adds host-restart step.  
  *Link: [PR #2769](nanocoai/nanoclaw PR #2769)*

- **PR #2757** *(merged)* — **Codex agent-provider payload v2**: Codex now operates on host capability seams with vault-only authentication through OneCLI.  
  *Link: [PR #2757](nanocoai/nanoclaw PR #2757)*

- **PR #2756** *(merged)* — **Operator-driven provider selection, switching, and memory migration**. Introduces provider registry, setup picker, installer, vault auth walkthrough, and memory-migration skill. This is the infrastructure for non-default providers.  
  *Link: [PR #2756](nanocoai/nanoclaw PR #2756)*

- **PR #2764** *(merged)* — Fixes two relocated file paths in `CLAUDE.md` Key Files table. Closes issue #2763.  
  *Link: [PR #2764](nanocoai/nanoclaw PR #2764)*

- **PR #2758** *(merged)* — Data-driven global CLI installs from `cli-tools.json` manifest, replacing hardcoded Dockerfile `ARG`+`RUN` blocks.  
  *Link: [PR #2758](nanocoai/nanoclaw PR #2758)*

**Additional progress:** Issue #2763 (docs path errors) was closed after its fix PR merged.

## 4. Community Hot Topics

While **no issues or PRs have more than 0 comments or reactions**, the following items draw attention by their nature:

- **Security issues #2760, #2761, #2762** (all filed by YLChen-007) describe vulnerabilities in file exfiltration, approval bypass via unauthenticated loopback webhook, and hidden MCP server arguments. These are high-severity reports that demand immediate maintainer response.  
  *Issues: [#2760](nanocoai/nanoclaw Issue #2760), [#2761](nanocoai/nanoclaw Issue #2761), [#2762](nanocoai/nanoclaw Issue #2762)*

- **Budget-exhausted turns dropped silently (#2751)** by assapin generated the fix PR #2759. This is a user-facing reliability issue that likely affects many deployments.  
  *Issue: [#2751](nanocoai/nanoclaw Issue #2751)*

- **Prompt caching default (#2768)** from galmorduku points to a performance optimization gap in the Claude provider. No discussion yet, but impactful for token cost.  
  *Issue: [#2768](nanocoai/nanoclaw Issue #2768)*

## 5. Bugs & Stability

**Bugs reported today (ranked by severity):**

| Severity | Issue | Description | Fix available? |
|----------|-------|-------------|----------------|
| **Critical** | [#2762](nanocoai/nanoclaw Issue #2762) | `add_mcp_server` approval flow hides `args` and `env` from approver – attacker can inject arbitrary arguments/environment | No PR yet |
| **Critical** | [#2761](nanocoai/nanoclaw Issue #2761) | Local gateway approval bypass via unauthenticated loopback webhook – unauthenticated requests can trigger privileged actions | No PR yet |
| **High** | [#2760](nanocoai/nanoclaw Issue #2760) | `send_file` tool allows arbitrary file exfiltration by accepting absolute paths without read restriction | No PR yet |
| **High** | [#2751](nanocoai/nanoclaw Issue #2751) | Budget/token-exhausted LLM turns silently dropped – user gets no reply | PR #2759 open (fixes) |
| **Medium** | [#2768](nanocoai/nanoclaw Issue #2768) | Prompt caching not enabled by default in Claude provider – every turn re-sends full system prompt, wasting tokens/cost | No PR yet |
| **Low** | [#2767](nanocoai/nanoclaw Issue #2767) | Telegram adapter on `channels` branch still uses legacy Markdown sanitizer; upstream fixed in v4.30.0 – migrating to native MarkdownV2 | Upstream fix available, no internal PR |

**Stability note:** PR #2732 (still open) aims to harden host + agent-runner from health audit findings, but is not yet merged.

## 6. Feature Requests & Roadmap Signals

**Merged features (likely in next version):**
- **Operator-driven provider selection** (PR #2756) – users can switch between AI providers and migrate memory.
- **Codex agent provider v2** (PR #2757) – Codex integrated as a full provider on host capability seams.
- **Data-driven CLI tooling** (PR #2758) – simplifies adding new CLI dependencies.

**Open feature PRs under review:**
- **PR #2770** – Fixes Codex file event delivery (image generation output was dropped) and adds `file` to `ProviderEvent` type.
- **PR #2766 & #2765** – Add `.format-lint-off` skills for `channels` and `providers` respectively (likely formatting/linting tooling for contributors).

**Roadmap prediction:** The next NanoClaw release will ship multi-provider infrastructure, Codex support, and the budget-error delivery fix (#2759). Security patches for the three advisories are urgent and may be fast-tracked.

## 7. User Feedback Summary

- **Silent budget drops (#2751):** User assapin reports that when an LLM turn exhausts token/spend budgets, the agent-runner drops the error and users receive no reply – a significant UX pain point. The PR #2759 delivers the error to the user.
- **Security concerns (#2760-2762):** YLChen-007 (likely a security researcher) disclosed three vulnerabilities with clear exploit paths. No maintainer response yet, but these issues are critical for trust.
- **Performance optimization request (#2768):** galmorduku asks for prompt caching to be enabled by default to reduce cost and latency – a quality-of-life improvement for Claude users.
- **Telegram Markdown incompatibility (#2767):** chiptoe-svg reports that the `channels` branch still uses a workaround for an upstream bug that is now fixed – suggests dropping the workaround.

Overall, users are actively contributing fixes and reporting real-world pain points, indicating healthy community engagement.

## 8. Backlog Watch

- **PR #2732 (Harden host + agent-runner from health audit findings)** – Open since June 11, last updated June 14, awaiting review/merge. This is a broad security and stability improvement (19 files) that could address some of the same concerns raised in the new security issues.  
  *Link: [PR #2732](nanocoai/nanoclaw PR #2732)*

- **Issue #2751 (Budget-exhausted turns dropped)** – Created June 12, still has 0 comments from maintainers despite an open fix PR. The maintainers should acknowledge and shepherd the fix PR #2759 to merge quickly.

- **No response yet from maintainers on the three security issues (#2760, #2761, #2762)** filed on June 14. This is a backlog priority; coordination is recommended to avoid disclosure without patch.

*End of digest for 2026-06-15.*

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-06-15

## 1. Today’s Overview

Project activity over the past 24 hours was minimal, with exactly one new issue opened and no pull requests or releases. The open issue count remains low (1), and no existing work items were closed or merged. While the repository appears to have a healthy historical trajectory, today’s data indicates a quiet development cadence with no signs of regressions or urgent maintenance demands.

## 2. Releases

No new releases were published today. The latest release in the repository remains unchanged; users are advised to continue using the current version.

## 3. Project Progress

No pull requests were updated, merged, or closed in the last 24 hours. No features, bug fixes, or other code changes advanced today.

## 4. Community Hot Topics

The only item with any activity is:

- **[Issue #955 – [enhancement] Identity based authentication support for Azure OpenAI LLM Provider](https://github.com/nullclaw/nullclaw/issues/955)** (opened today by **kunalk16**, 0 comments, 0 reactions)  
  *Underlying need:* The user requests support for `DefaultTokenCredential` (from Azure.Identity) to allow identity‑based authentication in the Azure OpenAI LLM Provider. This is driven by enterprise security policies that restrict the use of API keys or connection strings. Although the issue has no comments or upvotes yet, it highlights a growing enterprise requirement for zero‑trust, credential‑free authentication via Azure CLI login.

## 5. Bugs & Stability

No bugs, crashes, or regressions were reported in the last 24 hours. The project’s stability profile remains unchanged.

## 6. Feature Requests & Roadmap Signals

The sole feature request today is **Issue #955** (identity‑based authentication for Azure OpenAI). Given the increasing adoption of Azure OpenAI in regulated environments, this request is likely to receive attention soon. A plausible next version may introduce an optional authentication mode that uses `DefaultAzureCredential` when no explicit key is provided. Users should watch for PRs related to `providers/azure_openai`.

## 7. User Feedback Summary

The only direct user signal comes from the issue author, who describes a concrete pain point: **security policies in Azure subscriptions that prohibit API‑key‑based access** to Azure OpenAI services. The request implies a desire for seamless integration with existing Azure infrastructure (e.g., managed identities, Azure CLI credentials). While no other feedback was recorded today, this single request represents a real enterprise‑grade requirement that could affect adoption.

## 8. Backlog Watch

No long‑unanswered issues or stale pull requests were detected. The backlog is effectively empty, which may indicate either a well‑maintained project or a period of low feature‑request volume. No intervention from maintainers is required at this time.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-06-15

## 1. Today's Overview

The project remains highly active, with **22 issues updated** (17 open, 5 closed) and **44 pull requests updated** (29 open, 15 merged/closed) in the last 24 hours. No new releases were published, but the pipeline is busy: the Reborn (v2) WebUI and agent loop are seeing steady bug fixes and feature landings, particularly around capability auth, attachment support, and runtime‑context transparency. A new high‑level initiative (#4878) aims to dogfood IronClaw for its own engineering productivity, spawning sub‑issues for preview deployments, automated code review, test coverage tracking, and a cloud coding agent workflow. The project is clearly in an intensive development phase with both user‑facing and infrastructure improvements in flight.

## 2. Releases

No new releases in this period.

## 3. Project Progress (Merged/Closed PRs)

**15 pull requests were merged or closed** in the last 24 hours. Notable closures include:

*   **Image attachment support (WebChat v2)** – PR [#4738](https://github.com/nearai/ironclaw/pull/4738) (merged) wires the frontend upload UX for the Reborn SPA, completing the user‑visible side of the attachment feature (#4644). The backend stack had already landed.
*   **Runtime‑context slice** – PR [#4836](https://github.com/nearai/ironclaw/pull/4836) (merged) implements a new model‑visible system prompt section that tells the agent which channels are connected, where outbound delivery points, and how the run originated. This directly improves model situational awareness.
*   **Slack approval→auth→final-reply e2e re‑homed** – PR [#4873](https://github.com/nearai/ironclaw/pull/4873) (merged) re‑introduces a critical Slack e2e test that was previously removed as “born broken”. It now passes and improves CI reliability.

Other closed PRs likely include dependency bumps and smaller fixes among the 15 merged/closed items.

## 4. Community Hot Topics

Activity is concentrated on several sizable PRs that are still open and under review:

*   **[#4871](https://github.com/nearai/ironclaw/pull/4871) – Image attachment support for vision-capable models** (size L, low risk, core contributor). Follows the backend attachment work by actually sending image content as multimodal parts to vision models. This closes a gap documented in #4644.
*   **[#4588](https://github.com/nearai/ironclaw/pull/4588) – Observability seams for Reborn** (size XL, experienced contributor). Adds a trajectory observer hook and LLM provider injection so external benchmarking tools (nearai‑bench) can drive and observe a Reborn agent run with parity to the legacy path.
*   **[#4778](https://github.com/nearai/ironclaw/pull/4778) – Slack as a product‑adapter extension** (size XL, core). Moves Slack integration from a hardcoded built‑in channel to a Reborn extension, making the channel list itself extensible via the hosting application.
*   **[#4840](https://github.com/nearai/ironclaw/pull/4840) – Missing‑credential auth gate before approval** (size L, core). Reorders the gate pipeline so that a missing credential is surfaced *before* the human approval prompt, preventing “approve then fail” UX.
*   **[#4838](https://github.com/nearai/ironclaw/pull/4838) – Explicit gate‑open feedback for busy threads** (size XL, core). Replaces a previous park‑and‑resubmit approach with a clear rejection message when a thread is busy, making the user the retry actor.

Several issues with zero comments reflect active early‑stage tracking or simple bug reports, but the above PRs drive the most discussion and implementation risk.

## 5. Bugs & Stability

**High severity:**

*   **Provider‑backed MCP tool approval resume failure** – Issue [#4887](https://github.com/nearai/ironclaw/issues/4887) (open today): after approval in WebUI v2, the resumed invocation can fail with `capability input ref is not scoped to this loop run`. This blocks usage of tools like `nearai.web_search`. A fix is likely needed in the executor’s resume matching logic.
*   **“Illegal invocation” on non‑localhost HTTP** – Issue [#4874](https://github.com/nearai/ironclaw/issues/4874): chat send fails with `TypeError: Illegal invocation` when the WebChat v2 SPA is accessed over plain HTTP from a network hostname/IP. This is a browser‑side restriction that must be addressed for LAN deployments.

**Medium severity:**

*   **WebSocket helper conflicts with v2 auth contract** – Issue [#4870](https://github.com/nearai/ironclaw/issues/4870): the browser‑side `openEventSocket()` authenticates with a query token, while the v2 auth contract explicitly rejects that. The issue notes it may not currently hit the production path, but it’s a latent mis‑alignment.
*   **Shell command not visible in approval dialog** – Issue [#4852](https://github.com/nearai/ironclaw/issues/4852): the approval dialog only shows `Capability: builtin.shell` without the actual command, and the Activity history also omits it. Users cannot audit what they are approving.

**Low severity / UX:**

*   **Google Calendar auth flow** – Issue [#4884](https://github.com/nearai/ironclaw/issues/4884): the extension requests an access token from the user rather than guiding through OAuth.
*   **Extension post‑install guidance** – Issue [#4886](https://github.com/nearai/ironclaw/issues/4886): the “AUTH NEEDED” status does not explain what step is next (Configure, authorize, etc.).

No crash or regression appears to be unaddressed; several of these bugs already have open PRs or are being actively diagnosed.

## 6. Feature Requests & Roadmap Signals

The clearest roadmap signal is the **“Improve IronClaw Engineering Productivity”** initiative (#4878), which has already spawned concrete sub‑issues:

*   **Preview Deployments** (#4881) – Vercel‑like preview links for IronClaw PRs.
*   **Automated Code Review** (#4880) – AI review and resolution of comments.
*   **Coding Agent Cloud Workflow** (#4882) – assign issues to a cloud agent that produces PRs.
*   **Test Coverage Tracking** (#4883) – enforce coverage thresholds per PR.

These signal that the team intends to dogfood its own AI agent technology for development itself, a strong indicator of maturity and ambition.

Other recurring themes:

*   **Runtime context transparency** – The merged #4836 and the related cost/reward planning are making the agent loop more predictable for end users.
*   **Slack channel as extension** (#4778) – Extensibility of channel integrations is a clear architectural goal.
*   **Persistent “always allow” across threads** (#4835) – User‐facing feature to avoid repeated approvals per thread.
*   **Gated final‑answer nudge** (#4837) – When an agent ends an empty turn, the loop issues one extra model call to try to produce a real answer.

Prediction: Most of the open “size L/XL” PRs above (#4840, #4838, #4835, #4837, #4871) are likely to be merged in the next release (v0.30?). The productivity sub‑issues (#4881‑4883) are longer‑term but signal the direction for v0.31+.

## 7. User Feedback Summary

Real user pain points emerge from the issue tracker:

*   **Approval UX confusion** – Users do not see the actual shell command when approving `builtin.shell` (#4852). They also see “AUTH NEEDED” without knowing what to do (#4886). Both degrade trust and usability.
*   **Auth flow friction** – Google Calendar users are asked for an access token instead of being guided through OAuth (#4884). This is a poor developer experience.
*   **Network deployment restrictions** – The `Illegal invocation` error (#4874) prevents using the SPA from a LAN address, limiting local testing and team usage.
*   **MCP tool failures** – The approval resume bug (#4887) breaks a core use case (web search) after approval, frustrating users who already granted permission.

Satisfaction signals are less visible but the rapid closing of bugs (e.g., #4751 large response error closed on June 14) suggests the team is responsive. The dogfooding findings issues (#4879, #4692) aggregate usability problems and will likely lead to further UX improvements.

## 8. Backlog Watch

| Item | Age | Priority |
|------|-----|----------|
| [#3511](https://github.com/nearai/ironclaw/pull/3511) – Fix repeated compact tool schema lookups (PR, S size) | Opened 2026‑05‑12 (34 days) | Low risk, should be simple to merge. |
| [#3708](https://github.com/nearai/ironclaw/pull/3708) – Release PR (M size) | Opened 2026‑05‑16 (30 days) | **High priority** – this is the blocked release that should deliver many new features. Depends on CI and approval. |
| [#4002](https://github.com/nearai/ironclaw/pull/4002) – Bump actions group (dependabot, L size) | Opened 2026‑05‑24 (22 days) | Medium – dependency hygiene, but blocking CI if not updated. |
| [#4032](https://github.com/nearai/ironclaw/pull/4032) – Bump wasm group (dependabot, M size) | Opened 2026‑05‑25 (21 days) | Medium – similar to above. |
| [#4499](https://github.com/nearai/ironclaw/pull/4499) – Bump tokio ecosystem (dependabot, M size) | Opened 2026‑06‑05 (10 days) | Lower urgency but still accumulating. |

The release PR (#3708) is the most important backlog item – it has been open for over a month and its closure would unblock the delivery of all merged features to users. The dependabot PRs could be reviewed and merged in batches to reduce noise.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest – 2026-06-15

## 1. Today's Overview
The project saw no new issues or pull requests created in the last 24 hours, but six existing items received updates, indicating ongoing work on long-standing tasks. Activity was concentrated on three unmerged feature PRs for the Cowork session experience (in-session search, system sleep prevention, and a session timer) and one bug-fix PR that was closed (merged). Two UI-related issues remain open and stale since early April. No new releases were published. Overall, the pace is steady but driven by refinement of existing functionality rather than new feature churn.

## 2. Releases
*No new releases.* (The latest available release data is empty.)

## 3. Project Progress
Only one pull request was closed/merged in the last 24 hours:

- **PR #1465 (merged)** – `fix(scheduled-tasks): deleted scheduled tasks reappear as ghost sessions after restart` by @linlihua.  
  This fix addresses a data‑cleaning oversight: when a scheduled task is deleted, its associated local SQLite session (`cowork_sessions`) was not removed, causing ghost sessions to reappear after an application restart. The fix ensures both the gateway‑side (`cron.remove`) and the local session are cleaned.  
  [Link to PR](https://github.com/netease-youdao/LobsterAI/pull/1465)

No other PRs were closed or merged today.

## 4. Community Hot Topics
No single issue or PR attracted unusual discussion or reactions (all have 0 👍 and ≤1 comment). The most actively referenced items are the three still‑open feature PRs that the team continues to update:

- **#1429** – *feat(cowork): add in-session message search with mark.js highlighting*  
  [Link to PR](https://github.com/netease-youdao/LobsterAI/pull/1429)
- **#1430** – *feat(cowork): prevent system sleep during session runtime*  
  [Link to PR](https://github.com/netease-youdao/LobsterAI/pull/1430)
- **#1431** – *feat(cowork): streaming activity bar shows session‑elapsed timer*  
  [Link to PR](https://github.com/netease-youdao/LobsterAI/pull/1431)

These PRs address core usability gaps (search within a running conversation, preventing interruptions from system sleep, and giving users a visible progress indicator). They were created in early April and last updated on June 14, suggesting the team is actively reviewing or polishing them.

## 5. Bugs & Stability
One stability bug was fixed today (see §3). Two open UI bugs remain:

- **#1434** – *[OPEN]* Language inconsistency in search results when LobsterAI UI is set to Chinese (English text and button labels shown).  
  **Severity:** Low (localisation glitch). No fix PR yet.  
  [Link to Issue](https://github.com/netease-youdao/LobsterAI/issues/1434)

- **#1435** – *[OPEN]* Agent name overflow in custom agent creation modal (name extends beyond dialog boundary).  
  **Severity:** Low (visual polish). No fix PR yet.  
  [Link to Issue](https://github.com/netease-youdao/LobsterAI/issues/1435)

No new crashes, regressions, or critical stability issues were reported today.

## 6. Feature Requests & Roadmap Signals
The three open PRs (#1429, #1430, #1431) represent features that are likely to land in the next release:

- **In‑session message search** – users have requested the ability to find specific messages in long Cowork sessions.
- **System sleep prevention** – a reliability improvement for long‑running Agent tasks; users reported sessions being interrupted by system sleep.
- **Session timer** – provides real‑time elapsed time in the streaming status bar, analogous to tools like Claude Code.

No new feature requests were filed today. These PRs align with observed user needs for better real‑time feedback and session robustness.

## 7. User Feedback Summary
User pain points captured in the two open issues:

- **Language switching confusion** (Issue #1434): When the interface language is set to Chinese, search‑result empty‑state messages and button labels remain in English, breaking the expected immersion.
- **Poor agent naming UX** (Issue #1435): Long agent names cause UI overflow in the creation dialog, making the dialog unusable without truncation or scrolling.

Positive signals can be inferred from the PR descriptions: users are valuing session continuity (sleep prevention) and progress visibility (timer, search). The merged fix for ghost sessions addresses a subtle but frustrating data‑persistence bug that likely caused repeated user complaints.

## 8. Backlog Watch
Several items have been open since early April 2026 and were only updated notionally (stale labels) without closure:

| Item | Created | Last Updated | Status | Notes |
|------|---------|--------------|--------|-------|
| Issue #1434 – Chinese language inconsistency | 2026-04-03 | 2026-06-14 | Open | No assignee, no PR attached. Needs maintainer triage. |
| Issue #1435 – Agent name overflow | 2026-04-03 | 2026-06-14 | Open | Same as above. Low priority but should be fixed before next release. |
| PR #1429 – In-session search | 2026-04-03 | 2026-06-14 | Open, stale | No comments from maintainers; awaiting review merge. |
| PR #1430 – Sleep prevention | 2026-04-03 | 2026-06-14 | Open, stale | Same – appears ready but unmerged. |
| PR #1431 – Timer | 2026-04-03 | 2026-06-14 | Open, stale | Same. |

All three PRs and two issues have been untouched for over two months. The project maintainers should prioritise reviewing and merging these PRs, as they address clear user needs and have been in the pipeline for a long time.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest – 2026-06-15

*Generated from GitHub activity (moltis-org/moltis)*

---

## 1. Today's Overview
Project activity was very low over the last 24 hours. No pull requests were opened, updated, or closed, and no new releases were published. A single open enhancement issue (#1123) was updated, suggesting ongoing interest in expanding memory backend options. Overall, the project appears to be in a **quiet maintenance phase**, with no signs of urgent bug fixes or regression reports.

---

## 2. Releases
**None** – No new versions were tagged or published in the last 24 hours. The latest release remains unchanged (not specified in this snapshot). No migration notes or breaking changes to report.

---

## 3. Project Progress
**Merged/closed PRs today:** None.  
No feature branches or bug-fix pull requests were advanced. The repository shows zero PR activity for the reporting period.

---

## 4. Community Hot Topics
**Most active (and only) issue:**  

- **[#1123] [Feature]: Add pure-Rust turbovec as an alternative memory backend for extreme edge compression**  
  Author: `joeblew999` | Created: 2026-06-14 | Updated: 2026-06-14 | Comments: 0 | 👍: 0  
  [GitHub](https://github.com/moltis-org/moltis/issues/1123)

This is a standalone feature request with no discussion yet. The proposal highlights a desire for a **Rust-native vector compression library** that could replace or complement existing backends in memory-constrained edge environments. Although lacking community engagement so far, the request signals interest in **extreme compression** and **pure-Rust dependencies**—a common theme in embedded/edge AI.

---

## 5. Bugs & Stability
**No bugs, crashes, or regressions** were reported in the last 24 hours. The only issue is an enhancement, not a defect. Project stability appears stable at this snapshot.

---

## 6. Feature Requests & Roadmap Signals
**New feature request:**  
- `#1123` proposes integrating [turbovec](https://crates.io/crates/turbovec) as a memory backend for Moltis, targeting scenarios where **storage space is extremely limited**.  

**Prediction for next version:**  
Given the low activity and absence of other recent features, this request is unlikely to be implemented in the immediate next release unless it gains maintainer traction or community support. However, if the Moltis team prioritizes **edge deployment optimization**, turbovec integration could appear in a future minor release.

---

## 7. User Feedback Summary
**Available feedback is limited** to the single issue description. The author (`joeblew999`) expresses a need for **compact, efficient memory backends** that do not rely on C bindings (pure-Rust). The underlying pain point is likely **binary size or memory overhead** when deploying Moltis on low-resource devices. No other user satisfaction/dissatisfaction signals are present in this snapshot.

---

## 8. Backlog Watch
**No long-unanswered issues or PRs** were flagged. The repository has no items that require maintainer unblocking. The single open issue is less than 24 hours old and has not yet received a response.

---

*End of digest. Data refreshed as of 2026-06-15.*

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw (QwenPaw) Project Digest — 2026-06-15

*Generated from GitHub data for [agentscope-ai/QwenPaw](https://github.com/agentscope-ai/CoPaw)*

---

## 1. Today's Overview

The QwenPaw/CoPaw project shows **high community activity** with 21 Issues and 16 PRs updated in the last 24 hours. 4 Issues were closed and 5 PRs were merged/closed, indicating steady progress on bug fixes and feature work. However, the open bug count (17 issues) and several regressions in the latest `v1.1.11.post2` release suggest **stability challenges** that require maintainer attention. The community remains engaged, contributing both bug reports and pull requests (five first-time contributors today). No new releases were published.

---

## 2. Releases

*None today.*

---

## 3. Project Progress

**PRs merged/closed today** (5 items):

- [#5092](https://github.com/agentscope-ai/QwenPaw/pull/5092) — **Revert** `fix(pack): compile-check discord after conda-unpack`. Related to build issue #5086.
- [#5188](https://github.com/agentscope-ai/QwenPaw/pull/5188) — **Feature**: Added `Request payload transforms` — a chat registry list slot allowing plugins to intercept and modify outgoing API requests. Exposed via `window.QwenPaw.chat.requestPayload.add(...)`.
- [#5051](https://github.com/agentscope-ai/QwenPaw/pull/5051) — **Fix (desktop)**: Persist backend port across restarts to preserve localStorage, resolving [#4733](https://github.com/agentscope-ai/QwenPaw/issues/4733) (Windows desktop losing last-used agent on restart).
- [#5035](https://github.com/agentscope-ai/QwenPaw/pull/5035) — **Fix (local_models)**: Parse llama.cpp server version without fixed-width slice, preventing breakage when build numbers exceed 4 digits.
- [#5038](https://github.com/agentscope-ai/QwenPaw/pull/5038) — **Fix (context)**: Guard empty message list in `LightContextManager.pre_reply()` to prevent `IndexError`.

**Notable open PRs under review**:

- [#5189](https://github.com/agentscope-ai/QwenPaw/pull/5189) — Plugin command suggestions with cross-tab language sync (feature).
- [#5187](https://github.com/agentscope-ai/QwenPaw/pull/5187) — Windows desktop GUI automation via UIA + Tauri control mode (feature).
- [#5179](https://github.com/agentscope-ai/QwenPaw/pull/5179) — Expand multi-agent collaboration trigger keywords (fix).
- [#5186](https://github.com/agentscope-ai/QwenPaw/pull/5186) — Complete Vietnamese locale support (feature).

---

## 4. Community Hot Topics

**Most discussed Issues** (by comment count):

| Issue | Comments | Summary |
|-------|----------|---------|
| [#5047](https://github.com/agentscope-ai/QwenPaw/issues/5047) (CLOSED) | 5 | Windows Tauri desktop startup extremely slow (10+ minutes) — closed as known issue. |
| [#5156](https://github.com/agentscope-ai/QwenPaw/issues/5156) (OPEN) | 5 | Request to allow `kimi-for-coding` via `uv` whitelist — users with paid Kimi coding subscription cannot use it in QwenPaw. |
| [#5009](https://github.com/agentscope-ai/QwenPaw/issues/5009) (CLOSED) | 3 | Question about observability/tracing integration (Langfuse, OpenTelemetry). Closed, but interest remains. |
| [#5190](https://github.com/agentscope-ai/QwenPaw/issues/5190) (OPEN) | 2 | WeChat Work private chat approval UI missing — user can trigger permission approval but cannot see/review it. |
| [#5184](https://github.com/agentscope-ai/QwenPaw/issues/5184) (OPEN) | 2 | Local model providers not showing in v1.1.11.post2 — regression. |

**Underlying needs**: Users want better **integration with third-party services** (Kimi, Zalo, observability tools), **reliable desktop experience** (slow startup, white screen), and **transparent permission workflows** (WeChat approval UI). The high comment count on slow startup (#5047) indicates a **critical pain point** for Windows users.

---

## 5. Bugs & Stability

**Reported today (June 14–15)**, sorted by estimated severity:

| Severity | Issue | Description | Fix PR exists? |
|----------|-------|-------------|----------------|
| **Critical** | [#5184](https://github.com/agentscope-ai/QwenPaw/issues/5184) | Local model providers not displaying in v1.1.11.post2 — blocks local model usage. | No |
| **Critical** | [#5163](https://github.com/agentscope-ai/QwenPaw/issues/5163) | Gemini tool calling regression between v1.1.10 and v1.1.11.post2 — no known fix. | No |
| **High** | [#5161](https://github.com/agentscope-ai/QwenPaw/issues/5161) | Long conversations cause QwenPaw to stop responding entirely. | No |
| **High** | [#5171](https://github.com/agentscope-ai/QwenPaw/issues/5171) | Context compression can discard all context when persona file tokens exceed threshold, breaking task continuity. | No |
| **High** | [#5166](https://github.com/agentscope-ai/QwenPaw/issues/5166) | Plugin installation fails on Python 3.13 due to deprecated `imghdr` module. | No |
| **Medium** | [#5183](https://github.com/agentscope-ai/QwenPaw/issues/5183) | Pet feature broken on Wayland (Niri window manager). | No |
| **Medium** | [#5177](https://github.com/agentscope-ai/QwenPaw/issues/5177) | DingTalk channel messages not saved to `chats.json` — sessions invisible in console frontend. | No |
| **Medium** | [#5181](https://github.com/agentscope-ai/QwenPaw/issues/5181) | Plugin dependency pip install causes persistent cmd popup loops when PyPI is unreachable (v1.1.11.post2). | No |
| **Low** | [#5174](https://github.com/agentscope-ai/QwenPaw/issues/5174) | Cron/heartbeat agents unable to execute heavy tasks (write_file, spawn_subagent) due to timeout/design limits. | [#5180](https://github.com/agentscope-ai/QwenPaw/pull/5180) (open) |
| **Low** | [#5162](https://github.com/agentscope-ai/QwenPaw/issues/5162) | Thinking logic enters infinite loop. | No |
| **Low** | [#5165](https://github.com/agentscope-ai/QwenPaw/issues/5165) | Packaged installer results in white screen due to missing modules in spec file. | No |

**Regression alert**: Two critical regressions in `v1.1.11.post2` (local providers, Gemini tools) were confirmed working in `v1.1.10`. The community is advised to **avoid upgrading** until these are resolved.

---

## 6. Feature Requests & Roadmap Signals

**Top requested features today**:

| Issue | Feature | Likely priority for next release? |
|-------|---------|-----------------------------------|
| [#5156](https://github.com/agentscope-ai/QwenPaw/issues/5156) | Support `kimi-for-coding` via `uv` whitelist | High — directly affects paid subscribers. |
| [#5185](https://github.com/agentscope-ai/QwenPaw/issues/5185) | Add real-time timestamp (HH:MM:SS) to agent context | Medium — reduces redundant tool calls. |
| [#5168](https://github.com/agentscope-ai/QwenPaw/issues/5168) | Official Zalo bot channel | Medium — growing Vietnamese user base. |
| [#5167](https://github.com/agentscope-ai/QwenPaw/issues/5167) | Optimize Feishu CardKit streaming card for long replies | Medium — usability improvement. |
| [#5182](https://github.com/agentscope-ai/QwenPaw/issues/5182) | Unify model configuration for text/vector/audio/video | Low — likely longer-term redesign. |
| [#5178](https://github.com/agentscope-ai/QwenPaw/pull/5178) | Add session filter by title (PR open) | High — already implemented in PR; likely merged soon. |
| [#5189](https://github.com/agentscope-ai/QwenPaw/pull/5189) | Plugin command suggestions with autocomplete (PR open) | High — frontend improvement under review. |

**Seasonal trend**: Multiple requests from **Vietnamese community** (Zalo channel, Vietnamese locale) and a **new contributor spike** (6 PRs from first-time contributors, 3 from Vietnamese names) suggests growing Southeast Asian adoption.

---

## 7. User Feedback Summary

**Positive signals**:
- Community contributors are actively submitting PRs (5 first-time contributors today) — project health is good.
- Request for Vietnamese locale (#5175/#5186) was quickly addressed by community members, indicating responsive collaboration.

**Pain points & dissatisfaction**:

| Theme | Evidence | Sentiment |
|-------|----------|-----------|
| **Desktop performance** | #5047: Tauri startup takes 10+ minutes on Windows 11. #5165: White screen after packaging. | High frustration — users report "no response" and reinstalling doesn't help. |
| **Regression anxiety** | #5163, #5184: Features working in v1.1.10 broken in v1.1.11.post2. | Trust erosion — users are hesitant to upgrade. |
| **Context loss** | #5171: Task interruption due to compression discarding everything. | Critical for production use — agents become useless. |
| **Plugin UX** | #5181: cmd window popup storm on dependency install. | Annoying for desktop users, interferes with workflow. |
| **Permission workflow** | #5190: WeChat approval UI invisible — cannot approve pending requests. | Blocks enterprise adoption. |

**Real use cases mentioned**:
- Data analysis with `datapaw` plugin (PR #4622)
- Cron-based knowledge generation (#5174)
- Feishu (Lark) streaming card for long responses (#5167)
- Multi-agent team collaboration (#5179)

---

## 8. Backlog Watch

Issues and PRs that remain **open for a significant time** or lack maintainer response:

| Item | Age | Issue | Maintainer action needed |
|------|-----|-------|---------------------------|
| [#4622](https://github.com/agentscope-ai/QwenPaw/pull/4622) | 24 days | DataPaw plugin PR (May 22) — large feature, under review but no merge. | Review / request changes or merge. |
| [#4902](https://github.com/agentscope-ai/QwenPaw/pull/4902) | 13 days | Built-in PRD CRUD tool with frontend renderer (June 2). | Review / merge after conflict resolution. |
| [#5161](https://github.com/agentscope-ai/QwenPaw/issues/5161) | 3 days | Long conversation non-response — reported June 12, no official workaround. | Acknowledge and provide temporary fix (e.g., context limit). |
| [#5163](https://github.com/agentscope-ai/QwenPaw/issues/5163) | 3 days | Gemini regression — confirmed between v1.1.10 and v1.1.11.post2. | Priority regression; needs hotfix. |
| [#5184](https://github.com/agentscope-ai/QwenPaw/issues/5184) | 1 day | Local model providers missing in v1.1.11.post2 — same regression window. | Critical hotfix required. |

**Call to action**: The two critical regressions in `v1.1.11.post2` and the long-unreviewed DataPaw plugin are the highest-priority items. A patch release (`v1.1.12`) is strongly recommended to restore stability.

---

*Digest prepared for 2026-06-15. Data source: [agentscope-ai/QwenPaw](https://github.com/agentscope-ai/CoPaw).*

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest – 2026-06-15

**Project:** ZeroClaw (zeroclaw-labs/zeroclaw)  
**Data snapshot:** 2026-06-15 00:00 UTC – 24h window

---

## 1. Today’s Overview

ZeroClaw saw another day of intense development activity with **50 pull requests updated** (46 open, 4 merged/closed) and **3 issues updated** (1 open, 2 closed). The project remains release‑stable (no new releases today) while addressing a wide range of defects and enhancements. Team velocity is high, but a substantial backlog of high‑risk, open PRs (many requiring author action) signals that maintainer bandwidth may be stretched. The most noteworthy landing today is a **config‑framework overhaul** (PR #7594) that eliminates hard‑coded path special‑casing, and a **Windows‑compilation fix** for the AMQP channel (Issue #7452). A long‑standing RFC for air‑gapped execution (Issue #6293) remains blocked, awaiting maintainer review.

---

## 2. Releases

**No new releases** were published in the last 24 hours. The latest available release remains unchanged.

---

## 3. Project Progress (Merged/Closed PRs Today)

| PR | Title | Type | Summary |
|----|-------|------|---------|
| [#7664](https://github.com/zeroclaw-labs/zeroclaw/pull/7664) (CLOSED) | fix(issue): resolve #7542 `ask_user` failure | Bug fix | Fixes a regression where `ask_user` instantly failed with “Channel closed before receiving a response” in the gateway web dashboard. |
| [#7594](https://github.com/zeroclaw-labs/zeroclaw/pull/7594) (CLOSED) | feat(config): type‑driven alias‑ref pickers and self‑declaring config enums | Enhancement | Removes hard‑coded path special‑casing; config field UX now derived from Rust types. Internal representation only – no breaking TOML changes. |
| [#7384](https://github.com/zeroclaw-labs/zeroclaw/pull/7384) (CLOSED) | feat(cron): add pause/resume toggle to scheduled tasks | Enhancement | Exposes the already‑existing `enabled` field via the gateway PATCH endpoint, allowing dashboard users to pause jobs without deleting them. |
| Not shown in top 20 | (one additional merged PR) | – | – |

Key advances: the config refactor (#7594) reduces maintenance burden, and the cron pause/resume (#7384) fills a long‑standing UX gap. The `ask_user` fix (#7664) restores a critical interaction path.

---

## 4. Community Hot Topics

**Most active issue**  
- [#6293 – RFC: Air‑gapped execution mode with companion daemon over unix socket (enclave support)](https://github.com/zeroclaw-labs/zeroclaw/issues/6293) – **5 comments**; opened 2026-05-03, updated today. Labeled `blocked` and `needs-maintainer-review`. This RFC proposes splitting ZeroClaw into offline and online processes, enabling secure air‑gapped operation. High community interest, but no maintainer response visible.

**Most active open PRs (by comment count – all shown as undefined, but by volume of discussion inferred)**  
- [#5892 – fix(providers/runtime): two production blockers](https://github.com/zeroclaw-labs/zeroclaw/pull/5892) – stale‑candidate, high risk. Fixes empty `tool_choice` and orphaned `tool_use` – affects OpenAI, Bedrock, OpenRouter, Azure, Copilot providers. Rebased onto recent engine consolidation.  
- [#6693 – feat(memory): add dream mode for periodic memory consolidation](https://github.com/zeroclaw-labs/zeroclaw/pull/6693) – XL enhancement, needs author action. Adds five‑phase memory consolidation with optional LLM reflection.  

These PRs reflect deep user interest in **production reliability** (tool handling) and **long‑term memory** (dream mode).

---

## 5. Bugs & Stability

**High‑severity bugs reported/active today:**

| Issue | Title | Status | Notes |
|-------|-------|--------|-------|
| [#7542](https://github.com/zeroclaw-labs/zeroclaw/issues/7542) (underlying) | `ask_user` fails instantly in gateway web | **Fixed** via PR #7664 | Was a blocking user interaction bug. |
| [#7452](https://github.com/zeroclaw-labs/zeroclaw/issues/7452) | AMQP channel cannot compile on Windows – Unix‑only tokio reactor | **Closed** | S2 severity; fix is Windows‑specific. |
| [#6989](https://github.com/zeroclaw-labs/zeroclaw/issues/6989) | `#[secret]` does not redact bearer tokens in header maps | **Closed** | P1 bug; extend secret derive to `HashMap`. |
| [#7424](https://github.com/zeroclaw-labs/zeroclaw/pull/7424) (open) | `allowed_private_hosts = ["*"]` doesn’t cover DNS‑resolved private hosts | Open, high risk | Fix required for correct wildcard behavior. |
| [#7616](https://github.com/zeroclaw-labs/zeroclaw/pull/7616) (open) | Groq rejects `reasoning_content` on outbound replay | Open, medium risk | Workaround needed for Groq compatibility. |
| [#7617](https://github.com/zeroclaw-labs/zeroclaw/pull/7617) (open) | Extra‑nested provider aliases silently drop fields | Open, medium risk | Adds a warning detector. |
| [#7610](https://github.com/zeroclaw-labs/zeroclaw/pull/7610) (open) | Quickstart does not prompt for webhook port | Open, high risk | Breaks webhook channel setup for newcomers. |
| [#7608](https://github.com/zeroclaw-labs/zeroclaw/pull/7608) (open) | Deferred MCP tools not exposed to delegates | Open, high risk | Fixes #6136; delegate tool visibility broken. |
| [#7583](https://github.com/zeroclaw-labs/zeroclaw/pull/7583) (open) | Runtime does not honor profile tool iteration limits | Open, high risk | Affects cron/CLI agent runs; bypasses limits. |

**Stability note:** 8 open PRs with `risk: high` were updated today, many with `needs-author-action`. While the project is fixing bugs quickly, multiple production‑blocking issues remain unresolved.

---

## 6. Feature Requests & Roadmap Signals

**Notable open RFCs and enhancements:**

- **Air‑gapped execution** ([#6293](https://github.com/zeroclaw-labs/zeroclaw/issues/6293)) – Designs a secure split‑architecture. Could be a major next version feature if approved.
- **Dream mode memory consolidation** ([#6693](https://github.com/zeroclaw-labs/zeroclaw/pull/6693)) – XL PR implementing periodic memory pruning and reflection. Local‑only by default.
- **Matrix room management** ([#7661](https://github.com/zeroclaw-labs/zeroclaw/pull/7661), open) – Adds `create_room` and `invite_user` APIs for the Matrix channel. Enhances collaboration.
- **Cron pause/resume** – Already merged (#7384); now live.

**Prediction for next minor version:** The air‑gapped RFC (#6293) and dream mode (#6693) are both large, but the config type‑driven picker (#7594) just landed and will likely be part of the next release. The Matrix room management PR (#7661) and deferred MCP visibility fix (#7608) have high user demand and may be merged soon.

---

## 7. User Feedback Summary

- **Pain point – Windows compatibility:** Issue #7452 (AMQP compile failure) was closed today, but the root cause (Unix‑only reactor) hints at ongoing gaps in Windows support.  
- **Pain point – Configuration complexity:** Multiple issues/PRs this week (#6989, #7617, #7610) reveal that configuration options (secret redaction, provider aliases, quickstart port) are error‑prone. Users are hitting silent failures and need better defaults/validation.  
- **Satisfaction – Cron pause/resume:** Users requested the ability to pause scheduled tasks via dashboard; PR #7384 landed today, likely addressing a common workflow need.  
- **Frustration – Production reliability:** The stale but still open #5892 (two production blockers in provider/tool handling) has been rebased but remains unmerged, indicating users running in production may still hit `tool_choice`/`tool_use` failures.

---

## 8. Backlog Watch

**Issues and PRs needing maintainer attention (unanswered > 7 days, blocked, or stale):**

| Item | Days since last maintainer action | Status | Reason for concern |
|------|-----------------------------------|--------|-------------------|
| [#6293 – Air‑gapped RFC](https://github.com/zeroclaw-labs/zeroclaw/issues/6293) | 43 days (created May 3) | Open, blocked | No maintainer review; 5 comments; high community interest. |
| [#5892 – Production blocker PR](https://github.com/zeroclaw-labs/zeroclaw/pull/5892) | 57 days (created Apr 19) | Open, stale‑candidate | Last update today (rebase) but still not merged; impacts many providers. |
| [#6693 – Dream mode PR](https://github.com/zeroclaw-labs/zeroclaw/pull/6693) | 30 days (created May 16) | Open, needs‑author‑action | Author has not responded to review; no maintainer forcing hand. |
| [#7616 – Groq reasoning fix](https://github.com/zeroclaw-labs/zeroclaw/pull/7616) | 1 day (updated today) | Open, needs‑author‑action | Fresh PR, but label indicates author must address feedback. |
| [#7617 – Extra‑nested alias warning](https://github.com/zeroclaw-labs/zeroclaw/pull/7617) | 1 day | Open, needs‑author‑action | Same pattern. |

**Notable:** 6 out of the top 10 open PRs carry the `needs-author-action` label. Without author engagement, these fixes risk stalling. The `blocked` label on #6293 and `stale-candidate` on #5892 suggest these may require maintainer triage to unblock.

---

**Risk Summary**  
Project health is **active but fragile**: high throughput of bug fixes is counterbalanced by a heavy tail of high‑risk open items. The absence of a maintainer response on the air‑gapped RFC (#6293) is the most visible gap. The next 48‑72 hours of maintainer reviews on `needs-author-action` PRs will determine whether the backlog shrinks or grows.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*