# OpenClaw Ecosystem Digest 2026-06-14

> Issues: 230 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-14 11:10 UTC

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

# OpenClaw Project Digest — 2026-06-14

A hyperactive day across the repository: 230 issues updated, 500 pull requests touched, and 134 PRs closed or merged in the last 24 hours. A new beta release (v2026.6.8) landed with richer Telegram/WhatsApp delivery and several stability fixes. However, the community continues to wrestle with gateway memory leaks (P0), session compaction bugs, and channel-specific delivery gaps. Maintainers remain the bottleneck on many high‑impact items.

---

## Releases

### v2026.6.8‑beta.1 (2026.6.8)

**Highlights** – Telegram and WhatsApp channel delivery are richer and less brittle:

- Telegram now supports structured rich text (tables, lists, expandable blockquotes), prompt‑preserving CLI backend delivery, retired native draft migration, and safer rich‑media boundaries.
- WhatsApp delivery hardening (further details not excerpted but part of the release).

📦 [Release page](https://github.com/openclaw/openclaw/releases/tag/v2026.6.8-beta.1)

---

## Project Progress (Merged/Closed PRs Today)

134 PRs were closed or merged in the past 24 hours. Among the most notable from the top issues/PRs:

- **#53399** (Browser control server hang) – closed; fix likely merged.  
  [Issue](https://github.com/openclaw/openclaw/issues/53399)

- **#91552** (Codex app‑server binding sidecar stale model) – closed; fix merged.  
  [Issue](https://github.com/openclaw/openclaw/issues/91552)

- **#89996** (GPT‑5.5 Codex WebChat timeout after completed tool calls) – closed.  
  [Issue](https://github.com/openclaw/openclaw/issues/89996)

- **#78897** (OpenAI Responses provider `store=true` fallback) – closed.  
  [Issue](https://github.com/openclaw/openclaw/issues/78897)

The **open PR pipeline** shows substantial forward motion: many PRs tagged `size:XL` or `proof:supplied` are awaiting maintainer review, including a GTK‑native Linux companion app (#59859) and a lane‑oriented channel interface (#59986).

---

## Community Hot Topics

The most engaged discussions (by comment count and reactions):

### 🔥 #58450 – Agent can promise a later follow‑up without starting any follow‑up action
- **15 comments, 3 👍**  
  Users report that the agent says “I’ll check project memory and come back” but never actually fires a sub‑agent or cron job. This leads to dropped user expectations.  
  [Issue](https://github.com/openclaw/openclaw/issues/58450)

### 🔥 #57901 – Safeguard compaction ignores `compaction.model` config
- **14 comments, 1 👍**  
  The safeguard extension incorrectly uses the session’s main model instead of the explicitly configured compaction model.  
  [Issue](https://github.com/openclaw/openclaw/issues/57901)

### 🔥 #57326 – CLI‑backed helper paths still bypass CLI dispatch
- **13 comments, 1 👍**  
  Despite an earlier fix, some helper paths continue to bypass `runCliAgent()` for CLI‑backed models.  
  [Issue](https://github.com/openclaw/openclaw/issues/57326)

### 🔥 #51429 – Hardcoded working path (`/Users/wangtao`) merged into published code
- **12 comments**  
  A Chinese user discovered a developer’s hard‑coded workspace path that creates unwanted directories.  
  [Issue](https://github.com/openclaw/openclaw/issues/51429)

### 🔥 #27445 – Feature: `announceTarget` for sub‑agent completion routing (11 comments, 5 👍)
- Long‑standing request to route sub‑agent completion announcements as user messages (back to parent session) for multi‑step orchestration.  
  [Issue](https://github.com/openclaw/openclaw/issues/27445)

### 🔥 #91588 – Gateway Memory Leak – RSS grows from 350 MB to 15.5 GB over days (P0, 11 comments)
- Critical memory leak causing repeated OOM kills. Affects all deployments.  
  [Issue](https://github.com/openclaw/openclaw/issues/91588)

### 🔥 #59330 – Control UI Raw mode permanently disabled (8 comments, 14 👍)
- A regression in 2026.3.31 that breaks config editing for power users.  
  [Issue](https://github.com/openclaw/openclaw/issues/59330)

**Underlying needs**: Users want reliable background execution (follow‑ups), predictable memory usage (leak fix), config control (Raw mode), and channel fidelity (CLI dispatch, bad compaction).

---

## Bugs & Stability

### 🔴 Critical (P0)
- **#91588** – Gateway memory leak: `sessions.json` unbounded growth from duplicated `skillsSnapshot`, no pruning of ephemeral sessions.  
  [Issue](https://github.com/openclaw/openclaw/issues/91588)  
  *No merged fix PR yet; PR #91712 (prune diagnostic sessions) is a partial solution.*

### 🟠 High (P1)
- **#55334** – `sessions.json` duplication leads to gateway OOM (linked to #91588).  
  [Issue](https://github.com/openclaw/openclaw/issues/55334)
- **#57326** – CLI‑backed helpers still bypass CLI dispatch.  
  [Issue](https://github.com/openclaw/openclaw/issues/57326)
- **#56733** – Gateway event loop frozen; all HTTP requests timeout.  
  [Issue](https://github.com/openclaw/openclaw/issues/56733)
- **#57567** – Configuration migration failure during upgrade (v3.24 → v3.28).  
  [Issue](https://github.com/openclaw/openclaw/issues/57567)
- **#54634** – Update silently drops config when `HOME` changes.  
  [Issue](https://github.com/openclaw/openclaw/issues/54634)
- **#55694** – Agent tool‑call dead‑loop causes repeated message spam (Chinese user report).  
  [Issue](https://github.com/openclaw/openclaw/issues/55694)
- **#56096** – Telegram `sendChatAction` infinite retry loop with no backoff.  
  [Issue](https://github.com/openclaw/openclaw/issues/56096)
- **#56488** – Cron `deleteAfterRun` re‑triggers instead of deleting.  
  [Issue](https://github.com/openclaw/openclaw/issues/56488)
- **#54488** – Session lane starvation (follow‑up drain monopolizes lane).  
  [Issue](https://github.com/openclaw/openclaw/issues/54488)

### 🟡 Medium (P2)
- **#51429** – Hardcoded path in published code.
- **#58514** – Google Chat group messages silently ignored (DMs work).
- **#58737** – Slack display name reverts to bot default on edited messages.
- **#58702** – WebChat text covered by action icons after upgrade.
- **#56692** – Group chat context blurring (Telegram with multiple agents).
- **#59662** – Anthropic usage alert blocks delivered as assistant messages.
- **#58289** – Non‑main agent may miss `models.json`, fail with “Unknown model”.
- **#57256** – `openclaw status` falsely reports mem0 as unavailable.

**Fix PRs in flight**:  
- PR #92914 – Clamp unsupported thinking for subagent spawns (fixes half‑failure).  
- PR #91712 – Prune non‑idle stale diagnostic session entries (partial memory fix).  
- PR #91448 – Validate Telegram cron delivery chat ID.  
- PR #89081 – Skip unreadable tool names in construction plan.  
- PR #89092 – Fix LaunchAgent plist path for external `HOME`.  
- PR #89074 – Let gateway queue active‑run sends in TUI.

---

## Feature Requests & Roadmap Signals

The community is pushing for:

| Feature | Issue / PR | Likely Next Version? |
|---------|------------|----------------------|
| `announceTarget` for sub‑agent completion routing | [#27445](https://github.com/openclaw/openclaw/issues/27445) | ✅ – 11 comments, 5 👍, often requested |
| Guarantee last N raw messages survive compaction/reset | [#58818](https://github.com/openclaw/openclaw/issues/58818) | ⏳ – P2, community strong interest |
| Force reply to originating channel (Telegram/Discord/WhatsApp) | [#54531](https://github.com/openclaw/openclaw/issues/54531) | ⏳ – P1, linked PR open |
| Telegram Inline Query support (`@botname query`) | [#54794](https://github.com/openclaw/openclaw/issues/54794) | 🟡 – P2, good candidate |
| Outbound task calls (phone calls on user’s behalf) | [#59245](https://github.com/openclaw/openclaw/issues/59245) | 🟡 – P2, novel use case |
| Session labels / nicknames | [#55249](https://github.com/openclaw/openclaw/issues/55249) | 🟡 – P2, quality‑of‑life |
| Unbypassable outbound policy enforcement | [#56349](https://github.com/openclaw/openclaw/issues/56349) | 🟡 – P2, security‑focused |
| Adopt Claude Code’s multi‑layer compaction architecture | [#58398](https://github.com/openclaw/openclaw/issues/58398) | ⏳ – P2, architectural effort |
| `exec()` sandbox isolation (from Claude Code leak analysis) | [#58730](https://github.com/openclaw/openclaw/issues/58730) | ⏳ – P1, security review needed |
| GTK‑native Linux companion app | PR [#59859](https://github.com/openclaw/openclaw/pull/59859) | ✅ – size XL, proof supplied |
| Channel echo / session pinning | PR [#88815](https://github.com/openclaw/openclaw/pull/88815) | ✅ – size XL, waiting on author |

High‑certainty near‑term additions: `announceTarget`, GTK app, session pinning, and raw mode fix (the latter already has a high‑upvoted bug report #59330).

---

## User Feedback Summary

**Satisfaction drivers**:  
- Telegram/WhatsApp delivery improvements in v2026.6.8 are well received.  
- The community appreciates the ongoing work on CLI dispatch and compaction (though still buggy).  
- The GTK Linux app PR (#59859) signals commitment to desktop Linux.

**Dissatisfaction / pain points**:  
- **Memory leaks** (#91588, #55334) cause repeated crashes; users on macOS, WSL2, and Ubuntu all affected.  
- **Context loss** after compaction/reset (#58818) forces agents to “forget” ongoing tasks.  
- **Channel delivery gaps**: Google Chat group messages ignored (#58514), Slack avatar reverts (#58737), Telegram infinite retry (#56096).  
- **Upgrade breakage**: Config migration fails (#57567), config silently dropped (#54634), Control UI Raw mode disabled (#59330).  
- **Agent reliability**: Follow‑up promises never kept (#58450), tool‑call infinite loops (#55694), session lane starvation (#54488).  
- **Multilingual friction**: Chinese users report hardcoded paths (#51429) and Feishu streaming issues (PR #59771).  

Overall sentiment: “Exciting progress, but stability regressions are hurting day‑to‑day use.”

---

## Backlog Watch

Issues and PRs that have been open for weeks (or months) and need maintainer attention:

| Item | Age | Status Signal |
|------|-----|---------------|
| [#27445 – `announceTarget` feature](https://github.com/openclaw/openclaw/issues/27445) | ~3.5 months (2026‑02‑26) | 5 👍, 11 comments, `needs‑product‑decision`, `needs‑maintainer‑review` |
| [#51429 – Hardcoded path](https://github.com/openclaw/openclaw/issues/51429) | ~2.5 months (2026‑03‑21) | 12 comments, `linked‑pr‑open` but no movement; high embarrassment factor |
| [#57901 – Safeguard compaction model](https://github.com/openclaw/openclaw/issues/57901) | ~2.5 months (2026‑03‑30) | 14 comments, `needs‑maintainer‑review`, `needs‑product‑decision` |
| [#57326 – CLI bypass](https://github.com/openclaw/openclaw/issues/57326) | ~2.5 months (2026‑03‑29) | 13 comments, `needs‑maintainer‑review` |
| [#54531 – Force reply to originating channel](https://github.com/openclaw/openclaw/issues/54531) | ~2.5 months (2026‑03‑25) | 10 comments, 1 👍, `needs‑security‑review` |
| [#58398 – Multi‑layer compaction](https://github.com/openclaw/openclaw/issues/58398) | ~2.5 months (2026‑03‑31) | 5 comments, architectural proposal awaiting decision |
| PR [#59986 – Lane‑oriented channel interface](https://github.com/openclaw/openclaw/pull/59986) | ~2.5 months (2026‑04‑03) | size XL, `waiting on author` – may be superseded? |
| PR [#59859 – GTK Linux app](https://github.com/openclaw/openclaw/pull/59859) | ~2.5 months (2026‑04‑02) | size XL, `waiting on author` – needs final polish |

Most of these are tagged `needs‑maintainer‑review` or `needs‑product‑decision`. Tackling the P0 memory leak (#91588) is the single highest impact item.

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report — 2026-06-14

## 1. Ecosystem Overview

The open-source personal AI agent landscape is undergoing rapid maturation, characterized by a shift from experimental prototyping to production-grade reliability. Across the ecosystem, projects are converging on common challenges: memory management (leaks, compaction, context persistence), channel delivery fidelity (Telegram, Slack, WhatsApp, DingTalk), and security hardening against injection and data-loss vectors. A clear bifurcation is emerging between full-stack reference implementations (OpenClaw, Hermes Agent) and specialized, lighter-weight alternatives (NanoClaw, PicoClaw, TinyClaw). The ecosystem's health is mixed—while some projects are stabilizing after aggressive feature pushes, others are actively wrestling with critical P0 bugs that erode user trust. Notably, Chinese-language user communities (CoPaw, PicoClaw) are growing rapidly and surfacing distinct localization needs. Overall, the ecosystem is *developmentally active but stability-constrained*, with maintainer bandwidth remaining the primary bottleneck across most projects.

---

## 2. Activity Comparison

| Project | Issues Updated (24h) | PRs Updated (24h) | PRs Merged/Closed (24h) | Release Status | Health Score |
|---|---|---|---|---|---|
| **OpenClaw** | 230 | 500+ | 134 | v2026.6.8-beta.1 (today) | ⚠️ Moderate (P0 memory leak, critical bug) |
| **Hermes Agent** | 13 | 50 | 11 | None | ✅ Good (high velocity, responsive maintainers) |
| **ZeroClaw** | 4 | 50 | 11 | None | ✅ Good (active CI, architecturally advancing) |
| **IronClaw** | 11 | 21 | 10 | None | ✅ Good (focused epic delivery) |
| **NanoBot** | 6 | 31 | 11 | None | ✅ Good (security hardening in progress) |
| **CoPaw** | 9 | 12 | 1 | None | ⚠️ Fair (stale PRs piling up) |
| **NanoClaw** | 2 | 10 | 5 | None | ✅ Good (responsive maintainers) |
| **PicoClaw** | 2 | 7 | 5 | Nightly unstable | ⚠️ Fair (token drain bug unresolved) |
| **Moltis** | 2 | 3 | 0 | None | ✅ Good (focused on OAuth fix) |
| **NullClaw** | 1 | 1 | 0 | None | ⚠️ Stale (critical delivery bug open 14 days) |
| **LobsterAI** | 2 | 4 | 1 | None | ⚠️ Stale (stale-labeled items only) |
| **TinyClaw** | 0 | 0 | 0 | None | 🔴 Inactive |
| **ZeptoClaw** | 0 | 0 | 0 | None | 🔴 Inactive |

**Note:** Health Score is qualitative based on PR velocity, bug severity, maintainer responsiveness, and backlog age.

---

## 3. OpenClaw's Position

**Advantages over peers:**
- **Scale and community mass:** 500+ PRs and 230 issues in 24 hours dwarfs every other project—OpenClaw is the undisputed reference implementation. No other project comes within an order of magnitude.
- **Channel richness:** Telegram/WhatsApp structured rich text, prompt-preserving CLI backend delivery, and multi-channel support (Telegram, Discord, WhatsApp, Google Chat, Slack) far exceeds any peer. Moltis and IronClaw are closest but serve narrower use cases.
- **Feature breadth:** The roadmap includes `announceTarget`, GTK Linux app, session pinning, raw mode, and multi-layer compaction—a feature set that covers power users, desktop users, and enterprise orchestration.

**Technical approach differences:**
- OpenClaw uses a monolithic gateway architecture with `sessions.json` as the core persistence mechanism, which is now causing the P0 memory leak (#91588). In contrast, NanoClaw uses containerized runners with declarative manifests, and ZeroClaw employs a daemon+gateway split with RPC-based ACP bridge.
- OpenClaw's configuration system relies on a unified JSON schema with TUI/CLI control; IronClaw uses Rust-native config with typed structs. This makes OpenClaw more accessible but also more fragile (witness the Raw mode regression #59330 and hardcoded path #51429).

**Community size comparison:**
- OpenClaw's 230 issues and 134 merged PRs daily versus 13 issues and 11 merged for Hermes Agent, or 2 issues and 5 merged for PicoClaw, indicates a community perhaps 10-20x larger. However, *density of critical bugs is also proportionally higher*—the P0 memory leak affects all deployments, whereas NanoBot's critical bug (Anthropic temperature) was isolated to a specific model version.

**Vulnerability: maintainer bottleneck.** With 134 PRs merged but hundreds still awaiting review (many `size:XL` with proof supplied), OpenClaw's velocity is gated by core team capacity. Hermes Agent, by contrast, merges 11 PRs from 50 updated, suggesting a more efficient review pipeline relative to scale.

---

## 4. Shared Technical Focus Areas

The following requirements appear across multiple projects, indicating ecosystem-level pain points:

| Requirement | Projects Affected | Specific Needs |
|---|---|---|
| **Memory/context leak prevention** | **OpenClaw** (#91588, 15.5GB RSS), **CoPaw** (#5171, context reduced to zero), **Hermes Agent** (#46090, growing context slowdown) | Proper session pruning, compaction model correctness, bounded memory growth |
| **Channel delivery reliability** | **OpenClaw** (Telegram infinite retry #56096, Google Chat ignored #58514), **Hermes Agent** (Telegram Web unsupported #46098, WeChat mislabeling #46086), **CoPaw** (DingTalk chats not registered #5177), **NullClaw** (cron delivery fails #941) | Idempotent delivery, channel-specific error handling, chat persistence |
| **Configuration/upgrade stability** | **OpenClaw** (migration failure #57567, config silently dropped #54634, Raw mode disabled #59330), **NanoBot** (#4322, `session_key` crash after merge), **ZeroClaw** (capitalized alias invalidates quickstart #7591) | Backward-compatible migration, graceful fallback, clear error messages |
| **Security hardening** | **NanoBot** (13 open security PRs: symlink escapes, malformed payload rejection, cursor monotonicity), **Hermes Agent** (#44073, gateway fail-closed defaults; #10857, shell injection), **OpenClaw** (#56349, unbypassable policy enforcement) | Input validation, sandboxing, credential management, CVE patching |
| **Multi-provider support** | **NanoBot** (Ollama API #193), **Hermes Agent** (GLM-5.2, OpenRouter Fusion), **NanoClaw** (operator-driven provider switching #2756), **PicoClaw** (OpenRouter TTS), **OpenClaw** (OpenAI Responses fallback) | Provider-agnostic API shimming, runtime switching, credential injection |
| **Background/cron execution** | **CoPaw** (#5174, cron can't write files or spawn subagents), **OpenClaw** (#58450, follow-up promises never kept), **NullClaw** (#941, cron subprocess never starts) | Autonomous action capability, timeout management, delivery guarantees |
| **Localization/i18n** | **PicoClaw** (Traditional Chinese locale merged), **CoPaw** (Vietnamese #5169, Chinese chat hangs #5172), **LobsterAI** (#1434, Chinese empty-state in English) | Consistent i18n across UI, documentation, and error messages |

**Common architectural pattern:** *session/blob persistence* is a top-3 pain point in 5 of 7 active projects. The ecosystem has not yet converged on a durable, leak-free memory model.

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | Hermes Agent | IronClaw | NanoClaw | CoPaw | PicoClaw |
|---|---|---|---|---|---|---|
| **Primary user** | Power users, multi-channel deployers | Developers, security-conscious teams | Rust-copilot integrators | Provider-switching users | Chinese-language users | Lightweight multimodal |
| **Architecture** | Monolithic gateway + JSON persistence | Gateway + profile runtime | Rust crate ecosystem, Reborn pipeline | Containerized runners + declarative manifests | Electron desktop + cron agents | Nightly builds, low-ceremony |
| **Channel maturity** | ★★★★★ (Telegram, WhatsApp, Discord, Google Chat, Slack) | ★★★★☆ (Telegram, Slack, WeChat, WhatsApp voice) | ★★★☆☆ (Slack, WebChat) | ★★☆☆☆ (limited, via OneCLI) | ★★★☆☆ (DingTalk, QQ, WeChat) | ★★☆☆☆ (OpenRouter, basic) |
| **Security posture** | Reactive (bugs being found) | ★★★★★ (systematic audit underway) | ★★★★☆ (gate re-approval fix) | ★★★☆☆ (credential-stub mounts) | ★★☆☆☆ (plugin popups) | ★★★☆☆ (minimal attack surface) |
| **Innovation** | Multi-layer compaction, announceTarget | Telegram guest bots, WhatsApp voice sidecar | Universal attachments, runtime context | Operator-driven provider selection | Kimi-for-coding integration | Out-of-tree channel plugins |
| **Stability** | ⚠️ P0 leak in production | ⚠️ Data loss bugs resolved quickly | ✅ Nightly E2E red but feature areas solid | ✅ Silent budget drops fixed within hours | ⚠️ Context compression destroys tasks | ⚠️ Token drain bug open 9 days |
| **Community origin** | Global (US/EU heavy) | Global (US heavy, enterprise) | Rust ecosystem | English/tech-focused | China-heavy | OpenRouter/vision community |

**Key insight:** Hermes Agent is the closest direct competitor to OpenClaw in terms of channel breadth and developer focus, but leans heavily on security as a differentiator. IronClaw is the most architecturally robust (Rust-native, crates-based) but serves a narrower integrator audience. CoPaw has the strongest Chinese-language community but the most stability issues. PicoClaw is the most lightweight and experimental, targeting multimodal via OpenRouter.

---

## 6. Community Momentum & Maturity

### Tier 1: Rapidly Iterating (High Velocity, Active Maintainers)
- **OpenClaw** — Despite the P0 bug, 134 PRs merged/day signals immense development capacity. The ecosystem is *too big to fail* but must solve the memory leak to retain trust.
- **Hermes Agent** — 50 PRs/day with rapid bug closure (Anthropic bug fixed in <24h). Security audit underway suggests maturity push.
- **ZeroClaw** — 50 PRs/day with architectural consolidation (RFC #7415 closed, ACP bridge active). Strong engineering discipline.
- **NanoBot** — 31 PRs/day with systematic hardening. The `yu-xin-c` PR cluster signals a deliberate security review.
- **IronClaw** — 21 PRs/day focused on a single epic (attachments). Efficient, directed iteration.

### Tier 2: Steady State (Moderate Velocity, Stable but Needs Attention)
- **CoPaw** — 12 PRs/day but maintainer bottleneck on 5 critical PRs from `ly-wang19` that have been unreviewed for 5 days. Risk of contributor disillusionment.
- **PicoClaw** — 7 PRs/day, responsive to image hallucination (#3108 fixed quickly), but token drain bug (#3012) lingering.
- **NanoClaw** — 10 PRs/day, excellent responsiveness (budget drop fixed <24h), but low community size.
- **Moltis** — 3 PRs/day, healthy for its niche (MCP OAuth), but OAuth fix (#1120) must merge to maintain user trust.

### Tier 3: Stabilizing / Low Activity
- **NullClaw** — 1 PR/day, critical cron delivery bug open 14 days. Risk of abandonment if not addressed soon.
- **LobsterAI** — 4 stale items updated (likely stale-bot touch). No new work. Maintenance-only phase.
- **TinyClaw** — No activity. Effectively dormant.
- **ZeptoClaw** — No activity. Dormant.

**Overall ecosystem trajectory:** The top 6 projects are healthier than they were 3 months ago (more contributors, more structured roadmaps), but the median project is struggling to keep up with bug reports. The ecosystem is consolidating around a few strong implementations while lesser-maintained projects risk falling into disuse.

---

## 7. Trend Signals

### Emerging Industry Trends from Community Feedback

1. **Reliability over features.** Users across OpenClaw (#91588, memory leak), CoPaw (#5171, context loss), and NullClaw (#941, cron failure) are explicitly prioritizing predictable, leak-free operation over new capabilities. The "excitement" around Telegram/WhatsApp features in OpenClaw is tempered by "stability regressions are hurting day-to-day use."

2. **Guest/zero-registration agent access.** The Telegram Guest Bots feature (#21587 in Hermes Agent, 11 comments) and the push for `announceTarget` (#27445 in OpenClaw) signal demand for agents that can participate in group conversations without requiring bot admin rights or full channel registration—a multi-agent, organic interaction model.

3. **Multi-provider runtime switching.** NanoClaw's operator-driven provider selection (#2756), Hermes Agent's multiple model additions (GLM-5.2, OpenRouter Fusion), and CoPaw's Kimi-for-coding request (#5156) all point to users wanting to *dynamically swap backends* based on cost, capability, or latency—not just configure once at install time.

4. **Security as a differentiator.** Hermes Agent's fail-closed gateway PR (#44073) and NanoBot's 13-PR security audit cluster (symlink escapes, malformed payload rejection) are among the most significant engineering efforts in the ecosystem. This follows recent Claude Code leak analysis (#58730 in OpenClaw). Expect security to become a headline feature for agent frameworks in the next 6 months.

5. **Chinese-language ecosystem growth.** CoPaw (DingTalk, QQ, WeChat), PicoClaw (Traditional Chinese locale), and Hermes Agent (WeChat/Weixin fixes) all show active Chinese-language user bases. The hardcoded `/Users/wangtao` path in OpenClaw (#51429) is a symptom of this demographic shift—developers from Asian markets are increasingly contributing but need i18n-aware code practices.

6. **Desktop and multimodal convergence.** OpenClaw's GTK Linux app PR (#59859), IronClaw's attachment pipeline (#4644 epic), and PicoClaw's image compression PR (#2964) all point toward rich desktop and vision-capable agents. The mobile-only era is giving way to desktop-first agents with camera/image input.

### Value for AI Agent Developers

- **Invest in memory management.** Leaks and context corruption are the #1 source of user dissatisfaction across the entire ecosystem. Any new agent framework should prioritize a `context_blob` + `session_pruning` architecture from day one.
- **Provide zero-config but allow deep customization.** The push for raw mode (#59330 in OpenClaw) and the backlash against silent defaults (#2751 in NanoClaw) show that users want both turnkey operation *and* escape hatches for power users.
- **Design for multi-tenant environments.** IronClaw's Railway QA issues (tool activity disappears, #4853/#4852) and ZeroClaw's quickstart fragility (#7591) highlight that agent frameworks need robust isolation and reproducible setup flows.
- **Build for maintenance efficiency.** The maintainer bottleneck is the ecosystem's single greatest risk. Projects like NanoBot (community-driven security PRs) and CoPaw (first-time contributor localization) show that concrete, well-documented contribution pathways reduce this bottleneck. OpenClaw's 134 merged PRs/day is impressive but unsustainable if 200+ remain in review.

**Bottom line for decision-makers:** If you need the broadest channel support and community mass, choose **OpenClaw** but budget for memory overhead. If security and developer UX matter most, consider **Hermes Agent** or **IronClaw**. For lightweight, Chinese-market deployments, **CoPaw** is the strong candidate. For hyper-efficient containerized deployments, **NanoClaw** is promising but early. Avoid **NullClaw** until its cron delivery bug is resolved, and skip dormant projects (**TinyClaw**, **ZeptoClaw**) entirely.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-06-14

## 1. Today's Overview

NanoBot saw steady activity over the past 24 hours, with **6 issues updated** (5 closed, 1 open) and **31 pull requests touched** (11 merged/closed, 20 still open). No new releases were tagged. The project continues to focus heavily on **robustness and security hardening** — a cluster of defensive fixes from contributor `yu-xin-c` (media validation, monotonic cursors, symlink escapes, malformed payload rejection) now spans over a dozen open PRs. The WebUI team contributed a new automation management view and mobile responsiveness improvements, while documentation updates added affiliate links for Kimi and MiniMax. Overall health is good, though several long-running test and fix PRs remain open, indicating ongoing investment in quality assurance and edge-case handling.

---

## 2. Releases

**None** — no new versions were published in the last 24 hours.

---

## 3. Project Progress

**Merged/closed PRs today (11 total; key items from top-20 list):**

| PR | Description | Category |
|---|---|---|
| [#4270](https://github.com/HKUDS/nanobot/pull/4270) | Fix `compact_idle_session` to archive full history, preserving corrections in the last messages | Bug fix |
| [#4295](https://github.com/HKUDS/nanobot/pull/4295) | Add Kimi and MiniMax partner links in docs and README | Docs |
| [#4331](https://github.com/HKUDS/nanobot/pull/4331) | Localize update check copy in WebUI settings | WebUI improvement |
| [#4338](https://github.com/HKUDS/nanobot/pull/4338) | Make Kimi partner banner clickable in README | Docs |

Additional closed PRs (not shown in the top 20) contributed to the 11 count, including smaller doc fixes and configuration tweaks. The most significant merged fix is **#4270**, which resolves the “idleCompact” bug that could leave incorrect summaries in `history.jsonl` (see also Issue #4264).

---

## 4. Community Hot Topics

### Most commented issue
- **[#193 – Ollama API support?](https://github.com/HKUDS/nanobot/issues/193)** — 15 comments, closed in February but updated again yesterday. The user asks for Ollama API integration alongside existing vLLM support. Although closed without implementation, the continued commentary suggests lingering community demand.

### Most active open issues
- **[#4322 – NameError: 'session_key' is not defined](https://github.com/HKUDS/nanobot/issues/4322)** — Opened yesterday, 1 comment. Crashes agent startup after a merge. Root cause identified in `_build_memory_context`. **This is currently the only open bug issue** and is actively being discussed.

### Underlying community needs
- The Ollama request (#193) points to a desire for **broader local model provider support** beyond vLLM.
- The `session_key` crash (#4322) reveals pain around **branch merging and method extraction** – users who merge from `main` into custom branches risk breaking their agent startup.
- The closed Anthropic `temperature` bug (#4333, see below) highlights **sensitivity to API deprecations** – the community needs faster provider parameter shimming.

---

## 5. Bugs & Stability

Bugs reported today, ranked by severity:

| Severity | Issue | Description | Fix status |
|---|---|---|---|
| **Critical** | [#4333](https://github.com/HKUDS/nanobot/issues/4333) | Anthropic provider sends deprecated `temperature` to `opus-4-8`/Fable → 400 on every request. Only `opus-4-7` was exempted. | **Closed** (fixed; likely via a separate PR) |
| **High** | [#4322](https://github.com/HKUDS/nanobot/issues/4322) | `NameError: session_key` crashes agent at startup after a merge. Affects any branch merging `origin/main`. | **Open** – root cause identified; no fix PR yet. |
| **Medium** | [#4264](https://github.com/HKUDS/nanobot/issues/4264) | `idleCompact` uses truncated history (removes last 8 messages), leaving wrong summaries. User corrections lost. | **Closed** – fixed by merge of [#4270](https://github.com/HKUDS/nanobot/pull/4270). |
| **Medium** | [#4083](https://github.com/HKUDS/nanobot/issues/4083) | `pathAppend` appends after existing `PATH`, so configured tool executables don’t take precedence. | **Closed** – fix merged previously. |

### Additional fix PRs in progress (open)
Multiple defensive patches from `yu-xin-c` are still open and target stability:
- [#4340](https://github.com/HKUDS/nanobot/pull/4340) – preserve fenced code blocks when splitting messages
- [#4337](https://github.com/HKUDS/nanobot/pull/4337) – ignore empty injected payloads
- [#4336](https://github.com/HKUDS/nanobot/pull/4336) – reject malformed CLI `argv` payloads
- [#4315](https://github.com/HKUDS/nanobot/pull/4315) – ignore malformed history entries
- [#4312](https://github.com/HKUDS/nanobot/pull/4312) – reject malformed media attachments
- [#4311](https://github.com/HKUDS/nanobot/pull/4311) – reject non-positive file pagination limits
- [#4256](https://github.com/HKUDS/nanobot/pull/4256) – keep history cursor monotonic
- [#4119](https://github.com/HKUDS/nanobot/pull/4119) – block relative symlink workspace escapes
- [#4053](https://github.com/HKUDS/nanobot/pull/4053) – keep read-only roots out of write paths
- [#4323](https://github.com/HKUDS/nanobot/pull/4323) – resolve env vars before transcription config lookup
- [#4325](https://github.com/HKUDS/nanobot/pull/4325) – resolve env-var templates in WebUI settings update paths

These collectively strengthen input validation, memory integrity, and security boundaries.

---

## 6. Feature Requests & Roadmap Signals

### Recently requested features
- **[#4262](https://github.com/HKUDS/nanobot/issues/4262) – Use `botIcon` at agent startup** (closed). The user wants the configured `botIcon` to appear immediately instead of the default “puppy”. A straightforward UX polish. **Likely candidate for next minor release** (already closed, possible quick PR).
- **[#4330](https://github.com/HKUDS/nanobot/pull/4330) – WebUI Automation Management View** (open). A new surface for listing, filtering, running, pausing/resuming, and deleting automations. This is a **major feature** being actively developed by the core WebUI team — strong signal for the next version.
- **[#4339](https://github.com/HKUDS/nanobot/pull/4339) – Improve WebUI mobile responsiveness** (open). Tightens spacing, wraps overflow, compacts heatmaps. **Targeted for next UI release** to support mobile users.
- **Ollama API support** (#193, closed) – though closed, the 15 comments indicate persistent interest. Could be reconsidered if community demand grows further.

### Predicted roadmap items
- **Security hardening** – the steady stream of `yu-xin-c` PRs (input validation, symlink protection, cursor monotonicity) suggests a **systematic security audit** is underway. These will likely be merged in a batch soon.
- **Localisation expansion** – PR #4331 and #4338 show continued investment in i18n and partner integrations.
- **Memory & history reliability** – the `idleCompact` fix (#4270) and memory lifecycle harness (#4193) indicate the team is solidifying the memory subsystem for production use.

---

## 7. User Feedback Summary

### Pain points expressed
- **Anthropic model breakage** (Issue #4333): User `Ulef1005` reports complete request failure when using `opus-4-8` or Fable. This likely caused frustration until a fix was applied – now closed.
- **History summary corruption** (Issue #4264): User `imkuang` explains a real conversation scenario where the last corrective messages are lost in `idleCompact`, leaving erroneous records. The fix (#4270) directly addresses this.
- **Startup crash** (Issue #4322): User `professionelle-hypnose` encountered a crash after merging `main`. The dependency on `_build_memory_context` extraction broke their custom branch. They need better backward compatibility or clearer merge instructions.
- **Tool precedence misconfiguration** (Issue #4083): User `hamb1y` found that `pathAppend` doesn’t let configured tools override system executables – a subtle but important usability issue for tool developers.

### Use cases
- **Conversation correction flow** (from #4264): A user repeatedly corrects the model, and the correct final answer must be archived. This is a core use case for any assistant that learns from corrections.
- **Ollama local inference** (from #193): Users want to run NanoBot completely offline with Ollama’s API, not just vLLM.

### Satisfaction signals
- The **quick closure of #4333** (Anthropic temperature bug) within a day suggests responsive maintainers.
- The **rapid merge of #4270** (idleCompact fix) shows that community-reported regressions are taken seriously.

---

## 8. Backlog Watch

Items that may need maintainer attention:

| Item | Age | Significance |
|---|---|---|
| [#193](https://github.com/HKUDS/nanobot/issues/193) – Ollama API support | 4+ months (Feb 2026), closed | 15 comments, persistent community demand. If not planned, a clear “wontfix” communication would help manage expectations. |
| [#3982](https://github.com/HKUDS/nanobot/pull/3982) – Scripted agent runner harness | Open since May 24 (3 weeks) | Foundational test infrastructure, untouched. Could block other runner changes. |
| [#3983](https://github.com/HKUDS/nanobot/pull/3983) – Test blocked tool-call finish reasons | Open since May 24 | Companion to #3982, also untouched. |
| [#4053](https://github.com/HKUDS/nanobot/pull/4053) – Keep read-only roots out of write paths | Open since May 29 (2 weeks) | Security fix, no recent activity. |
| [#4193](https://github.com/HKUDS/nanobot/pull/4193) – Memory lifecycle harness | Open since June 4 (10 days) | Large test PR; could benefit from review. |
| [#4256](https://github.com/HKUDS/nanobot/pull/4256) – Keep history cursor monotonic | Open since June 8 (6 days) | Memory stability fix. |
| [#4322](https://github.com/HKUDS/nanobot/issues/4322) – `session_key` NameError | Open since yesterday | **Critical**, needs a fix PR soon. |

The cluster of open, unreviewed PRs from `yu-xin-c` (13+ items) suggests the maintainer team may be resource-constrained or waiting for a consolidated review pass. Designating a “hardening sprint” could help clear these.

---

*Digest generated from GitHub data for 2026-06-14.*

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest – 2026-06-14

## 1. Today’s Overview

Hermes Agent saw extremely high activity today: 13 issues updated (9 open, 4 closed) and 50 pull requests updated (39 open, 11 merged/closed). No new releases were published. The project is in a major development sprint with focus on security hardening, platform expansions (Telegram guest bots, WhatsApp voice, WeChat/Weixin correction), MCP reliability, and fixing several high-severity data-loss and race-condition bugs. The burst of 24 new issues (many opened today) indicates both heavy community testing and rapid iteration on stability.

## 2. Releases

*None.* No new versions were published today.

## 3. Project Progress

**Merged / closed PRs today (11 total):**  
Notable closures from the top-contributed PRs include:

- [#45793](https://github.com/NousResearch/hermes-agent/pull/45793) (P1, closed) – Fixed a critical Windows data-loss bug where desktop app restart/crash reset profile configs and deleted `.env` files.

**Closed issues today (4 total):**

- [#13248](https://github.com/NousResearch/hermes-agent/issues/13248) (P1) – Fixed empty-response retry loop on `claude-opus-4-7` in Slack group threads. The model would emit no text, triggering infinite retries.
- [#25114](https://github.com/NousResearch/hermes-agent/issues/25114) (P2) – Resolved `$HOME` environment variable corruption inside Hermes profile runtime contexts.
- [#45682](https://github.com/NousResearch/hermes-agent/issues/45682) (P1) – Fixed a structural race condition during gateway restart where `auto-resume` competed with inbound message processing.
- [#46053](https://github.com/NousResearch/hermes-agent/issues/46053) (P1, duplicate) – Fixed assistant responses silently dropped after `repair_message_sequence` due to conversation history mismatch.

These fixes address serious stability and data-integrity issues.

## 4. Community Hot Topics

**Most active issue:**

- [Issue #21587](https://github.com/NousResearch/hermes-agent/issues/21587) – *“Feature: Telegram Guest Bots, Bot-to-Bot, Stickers and Chat Automation”* (11 comments, 1 👍, open since May 8). This feature request draws on Telegram’s May 2026 Bot API update and has generated sustained discussion around multi-agent collaboration, guest-mode @mentions, and automation workflows.

**Most active PRs (by engagement, despite missing comment counts):**

- [#44073](https://github.com/NousResearch/hermes-agent/pull/44073) (P1, security) – Fail-closed gateway defaults and external-surface hardening. Represents a systemic security improvement.
- [#43049](https://github.com/NousResearch/hermes-agent/pull/43049) (P3) – Implements Telegram guest mode (`@mention` from non-member chats) directly linked to issue #21587.
- [#42334](https://github.com/NousResearch/hermes-agent/pull/42334) (P1, security/deps) – Bumps `aiohttp`, `anthropic`, and `cryptography` for CVE patches; highlights inconsistency between install paths.

The underlying community need is clear: users are eager for deeper Telegram integration (guest bots, voice, automation) and for a hardened, predictable security posture.

## 5. Bugs & Stability

**Bugs reported today (ranked by severity):**

| Severity | Issue | Description | Fix PR exists? |
|----------|-------|-------------|----------------|
| **P1** | [#46088](https://github.com/NousResearch/hermes-agent/issues/46088) | Gateway restart causes assistant messages to be lost from session transcript – permanent data loss. | No open fix yet. |
| **P1** | [#46090](https://github.com/NousResearch/hermes-agent/issues/46090) | Agent becomes extremely slow for basic tasks, even after restart – suspected growing context / compression issue. | No. |
| **P1** | [#46053](https://github.com/NousResearch/hermes-agent/issues/46053) | Assistant responses silently dropped (already closed with fix). | Yes (closed). |
| **P2** | [#46099](https://github.com/NousResearch/hermes-agent/issues/46099) | MCP client `TextReceiveStream` uses `errors='strict'` – fails on Windows with UnicodeDecodeError. | No. |
| **P2** | [#46100](https://github.com/NousResearch/hermes-agent/issues/46100) | Telegram batch-forward attachments can miss the initial text turn. | No. |
| **P3** | [#46098](https://github.com/NousResearch/hermes-agent/issues/46098) | Telegram Web shows “This message is not supported on the web version of Telegram”. | No. |
| **P3** | [#46086](https://github.com/NousResearch/hermes-agent/issues/46086) | Desktop UI mislabels Weixin (personal WeChat) as “WeChat (Official Account)”. | Yes: [#46096](https://github.com/NousResearch/hermes-agent/pull/46096) and [#46093](https://github.com/NousResearch/hermes-agent/pull/46093) (both open). |
| **P2** | [#46046](https://github.com/NousResearch/hermes-agent/issues/46046) | State.db accumulates too many sessions – no janitor/cleanup tool. Dashboard creates empty sessions. | No. |

The assistant message loss bug (both #46088 and the already-fixed #46053) is the most critical open stability concern. The performance regression (#46090) is alarming but lacks reproduction details.

## 6. Feature Requests & Roadmap Signals

**User-requested features in today’s issues:**

- **Telegram Guest Bots** ([#21587](https://github.com/NousResearch/hermes-agent/issues/21587), open, high engagement) – Bot-to-bot, stickers, chat automation. A PR [#43049](https://github.com/NousResearch/hermes-agent/pull/43049) already exists.
- **Session janitor / cleanup** ([#46046](https://github.com/NousResearch/hermes-agent/issues/46046)) – User requests a proper way to compact/clean session state.
- **Desktop font size setting** ([#46097](https://github.com/NousResearch/hermes-agent/issues/46097)) – Small text on high-DPI displays.
- **WeChat/Weixin correct labeling** ([#46086](https://github.com/NousResearch/hermes-agent/issues/46086)) – Already has two fix PRs.

**Signpost PRs indicating roadmap direction:**

- [PR #46028](https://github.com/NousResearch/hermes-agent/pull/46028) – WhatsApp voice-native audio and calling sidecar (merges two existing branches).
- [PR #46092](https://github.com/NousResearch/hermes-agent/pull/46092) – Adds GLM-5.2 to Z.AI provider.
- [PR #46094](https://github.com/NousResearch/hermes-agent/pull/46094) – Exposes OpenRouter Fusion as a model slug.
- [PR #45225](https://github.com/NousResearch/hermes-agent/pull/45225) – Perplexity web search backend plugin.
- [PR #41326](https://github.com/NousResearch/hermes-agent/pull/41326) – Configurable terminal shell (not just bash).
- [PR #44073](https://github.com/NousResearch/hermes-agent/pull/44073) – Security hardening of gateway defaults.

**Predictions for next minor release:**  
Likely includes Telegram guest bots (PR #43049), WeChat labeling fix (PR #46096/#46093), security patches (PR #44073, #42334), and at least one new model provider (GLM-5.2 or OpenRouter Fusion).

## 7. User Feedback Summary

**Pain points expressed today:**

- **Data loss & corruption:** Several users report assistant messages disappearing or session transcripts broken after gateway restarts or long sessions (#46053, #46088). One user describes “literally hours for nothing” due to slowdown (#46090).
- **Windows compatibility:** Configuration reset on crash (#45793, now fixed) and MCP Unicode errors (#46099) frustrate Windows desktop users.
- **Session bloat:** The dashboard auto-creates empty sessions, and there is no way to clean the state database (#46046).
- **Telegram Web issues:** A user cannot see messages on Telegram Web after the latest update (#46098).
- **Mislabeled platform:** Personal WeChat users are confused by the UI claiming “Official Account” (#46086).

**Satisfaction signals:**  
High community contribution rate (many PRs from diverse contributors) and detailed bug reports indicate an invested, technically sophisticated user base that values the project’s extensibility. The rapid closure of several P1 bugs (e.g., #45682, #46053) suggests responsive maintainers.

## 8. Backlog Watch

**Long-unanswered issues demanding attention:**

- [#21587](https://github.com/NousResearch/hermes-agent/issues/21587) (open since May 8, 11 comments) – Feature request for Telegram guest bots. While a PR exists (#43049), the issue remains open and discussion continues. High community interest.
- [#10857](https://github.com/NousResearch/hermes-agent/pull/10857) (open since April 16, P1 security) – Shell injection and SSRF via URL scheme validation in `hermes_cli/`. This is the oldest open PR in the top-20 list and addresses a HIGH severity static analysis finding. Needs maintainer review.
- [#36286](https://github.com/NousResearch/hermes-agent/pull/36286) (open since June 1, P3) – Adds China-region OAuth provider (minimax-cn-oauth). No comments from maintainers yet.
- [#42334](https://github.com/NousResearch/hermes-agent/pull/42334) (open since June 8, P1 security/deps) – CVE bumps with install-path inconsistency. Security-critical and still unmerged.

These items represent potential blockers or unmet needs that could erode user trust if left unresolved.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# 🐙 PicoClaw Project Digest — 2026-06-14

## 1. Today's Overview

The PicoClaw project continues with steady activity: 2 issues were updated in the last 24 hours (1 open, 1 closed), and 7 pull requests received updates (5 merged/closed, 2 open). A new nightly build was published. The community’s focus remains on stability and extensibility — key fixes landed for image‑description hallucinations and media routing, while a long‑standing token consumption bug is still under investigation. The nightly release (v0.2.9‑nightly) is marked unstable, so production users should stick with the stable v0.2.9 tag. Overall project health looks good, with maintainers actively merging contributions and closing bugs.

---

## 2. Releases

- **nightly**: Nightly Build  
  Tag: `v0.2.9-nightly.20260614.cf67dd38`  
  Status: Automated build, may be unstable.  
  [Full Changelog](https://github.com/sipeed/picoclaw/compare/v0.2.9...main)  
  No additional changelog, breaking changes, or migration notes were provided for this nightly.

---

## 3. Project Progress

Five pull requests were merged/closed today, advancing several areas:

- **[#2935] docs(i18n): add Traditional Chinese (zh‑TW) locale and READMEs**  
  *Merged* — Full Traditional Chinese documentation translations (README, CONTRIBUTING, frontend i18n).  
  URL: https://github.com/sipeed/picoclaw/pull/2935

- **[#3065] fix(seahorse): explicitly ignore Close() errors on PRAGMA/migration failure paths**  
  *Merged* — Silences linter warnings by explicitly discarding `db.Close()` errors.  
  URL: https://github.com/sipeed/picoclaw/pull/3065

- **[#3066] fix: explicitly ignore Close() errors on temp file write/sync failure paths**  
  *Merged* — Same pattern applied to three file‑handling paths.  
  URL: https://github.com/sipeed/picoclaw/pull/3066

- **[#3119] fix(tts): support OpenRouter voice overrides and fallback**  
  *Merged* — Adds per‑model TTS voice/response‑format overrides via `extra_body` and a single‑retry fallback mechanism.  
  URL: https://github.com/sipeed/picoclaw/pull/3119

- **[#3117] fix(agent): route media turns to image models**  
  *Merged* — Fixes hallucination issue #3108 by routing image‑related turns to the configured image model instead of falling back to a text‑only model.  
  URL: https://github.com/sipeed/picoclaw/pull/3117

Two pull requests remain open:

- **[#3120] feat(config): add RegisterChannelSettings hook for out‑of‑tree channels**  
  *Open* — Enables third‑party channel plugins without forking PicoClaw.  
  URL: https://github.com/sipeed/picoclaw/pull/3120

- **[#2964] Feat/image input compression**  
  *Open* — Adds configurable inbound image compression to the vision pipeline.  
  URL: https://github.com/sipeed/picoclaw/pull/2964

---

## 4. Community Hot Topics

The most active issue today is:

**[#3012] [BUG] Continuous consumption of tokens every minute when evolution is enabled**  
  - Author: xpader  
  - Comments: 3 (most commented item)  
  - Status: Open, last updated 2026‑06‑13  
  - URL: https://github.com/sipeed/picoclaw/issues/3012  

**Analysis:** This bug reports a potential cost‑intensive issue where tokens are consumed every minute when the evolution feature is active. The user provided detailed reproduction steps (FreeBSD, MiniMax model, Evolution Mode set to Draft, Code Path Trigger). With only three comments and no assignee, this is a high‑priority concern that likely will draw maintainer attention soon. Community members are probably waiting for troubleshooting guidance or a fix.

---

## 5. Bugs & Stability

| Severity | Issue / PR | Description | Status |
|----------|------------|-------------|--------|
| **High** | [#3012](https://github.com/sipeed/picoclaw/issues/3012) | Continuous token consumption when evolution is enabled – potential runaway costs. | Open, no fix yet. |
| **Medium** | [#3108](https://github.com/sipeed/picoclaw/issues/3108) (closed) | Image description requests hallucinate when active model lacks vision support. | Fixed by PR [#3117](https://github.com/sipeed/picoclaw/pull/3117) (merged today). |
| **Low** | Linter‑warning fixes ([#3065](https://github.com/sipeed/picoclaw/pull/3065), [#3066](https://github.com/sipeed/picoclaw/pull/3066)) | Explicitly ignoring `Close()` errors to suppress static analysis warnings. | Both merged. |

The nightly build is explicitly marked as unstable; users relying on `main` should expect potential regressions.

---

## 6. Feature Requests & Roadmap Signals

- **Out‑of‑tree channels** — PR [#3120](https://github.com/sipeed/picoclaw/pull/3120) adds a configuration hook for third‑party channels, signaling a move toward a plugin ecosystem. Likely to land in the next stable release (v0.3.0 or later).

- **Inbound image compression** — PR [#2964](https://github.com/sipeed/picoclaw/pull/2964) from May 28 is still open. It adds configurable compression before model payload building, addressing previous `max_media_size` limitations. Expect this to be reviewed and merged in the coming weeks.

- **TTS voice overrides** — Merged PR [#3119](https://github.com/sipeed/picoclaw/pull/3119) introduces per‑model voice/format overrides for OpenRouter TTS, a user‑requested flexibility improvement.

- **i18n expansion** — The merged Traditional Chinese locale (PR [#2935](https://github.com/sipeed/picoclaw/pull/2935)) shows ongoing commitment to internationalization.

**Prediction for next version:** Image compression and the channel plugin hook are the most likely additions. The token consumption bug (issue #3012) will likely be fixed before a stable release.

---

## 7. User Feedback Summary

- **Pain Points:**
  - *Token drain* (issue #3012): A user on FreeBSD with MiniMax reports excessive token usage when evolution is on — a high‑cost risk for API‑based deployments.
  - *Image hallucination* (issue #3108): When using `deepseek/deepseek-v4-flash` via OpenRouter, image descriptions were hallucinated. This is now fixed by PR #3117.
- **Use Cases:**
  - The merged TTS fix (#3119) supports per‑model voice overrides, indicating demand for fine‑grained control of multimodal outputs.
  - The Traditional Chinese locale (PR #2935) suggests a growing non‑English user base, especially among Mandarin‑speaking communities.
- **Satisfaction:** Quick turnaround on bug #3108 (reported June 11, fixed by June 13) demonstrates responsive maintainers. However, the token consumption bug (#3012) has been open since June 5 with no public fix yet, which may cause frustration for affected users.

---

## 8. Backlog Watch

Two items require maintainer attention:

- **[#3012] Token consumption bug** — Open since June 5, with 3 comments but no assignee or milestone. The issue includes detailed reproduction steps. Given the high severity, a maintainer response or a fix PR is overdue.

- **[#2964] Image compression feature** — Open since May 28, untouched for 17 days. This PR has no comments from maintainers. While not urgent, awaiting review could delay a valuable improvement for vision workloads.

No other long‑unanswered issues or PRs were detected in the 24‑hour window.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-14

## 1. Today’s Overview
The project saw **moderate activity** over the last 24 hours: 2 issues were touched (1 open bug, 1 self-deleted duplicate) and 10 pull requests were updated, evenly split between **5 merged/closed** and **5 still open**. No new releases were published. The work is heavily focused on **provider extensibility** (a new capability seam, operator-driven provider selection, a Codex v2 payload, memory scaffold) and **stability hardening** (recovered from a health audit, fixes for silent budget drops and stale database journals). The rapid follow-up on issue #2751 with a dedicated fix PR (within 24 hours) signals strong maintainer responsiveness.

## 2. Releases
**None** – No new versions were cut today.

## 3. Project Progress
Five pull requests were merged or closed today, reflecting convergent work on core infrastructure and provider scaffolding:

- **#2758** [CLOSED] – `feat(container): data-drive global CLI installs from cli-tools.json`  
  *Moves Dockerfile hard‑coded `ARG`+`RUN` blocks to a declarative manifest; skills add tools via JSON merge.*  
  [PR #2758](https://github.com/nanocoai/nanoclaw/pull/2758)

- **#2754** [CLOSED] – `feat(runner): onExchangeComplete provider hook + slash-command interruption`  
  *Adds an optional hook at the end of an LLM exchange and enables interruption via slash commands.*  
  [PR #2754](https://github.com/nanocoai/nanoclaw/pull/2754)

- **#2747** [CLOSED] – `feat(onecli): SDK 2.2.1 — credential-stub mounts + machine-checkable pins`  
  *Bumps the OneCLI SDK to 2.2.1; introduces credential‑stub mounts and pin‑based machine verification.*  
  [PR #2747](https://github.com/nanocoai/nanoclaw/pull/2747)

- **#2746** [CLOSED] – `feat(providers): agent-surfaces capability seam`  
  *Creates a host‑side registry where providers declare capabilities (e.g., supports memory scaffold, persistent state).*  
  [PR #2746](https://github.com/nanocoai/nanoclaw/pull/2746)

- **#2745** [CLOSED] – `feat(memory): opt-in persistent memory scaffold for providers`  
  *Adds a `usesMemoryScaffold` provider capability and a container‑based memory scaffold for providers that opt in.*  
  [PR #2745](https://github.com/nanocoai/nanoclaw/pull/2745)

These merged PRs collectively advance **provider consistency**, **memory persistence**, and **operational reliability** (Dockerfile simplification, SDK bump).

## 4. Community Hot Topics
No issues or PRs attracted comments or reactions today; community discussion was minimal. However, the **open bug #2751** — “Budget‑exhausted LLM turns are silently dropped” — is the most significant community‑facing issue, as it directly affects user experience (no feedback when spend or token budgets are exhausted). A fix PR (#2759) was opened within hours of the bug being filed.

- **#2751** [OPEN] – *User reports silent drops on budget‑limited turns.*  
  [Issue #2751](https://github.com/nanocoai/nanoclaw/issues/2751)

The cluster of **open PRs from contributor `omri‑maya`** (#2756, #2757) on provider selection and Codex v2 signals a major feature push that is likely to attract attention once merged.

## 5. Bugs & Stability
**High‑severity** (fix PR exists):  
- **#2751** – Budget/token‑exhausted LLM turns are dropped with no feedback to the user. *Fix PR #2759 is open.*  
  [Issue #2751](https://github.com/nanocoai/nanoclaw/issues/2751) | [Fix PR #2759](https://github.com/nanocoai/nanoclaw/pull/2759)

**Medium‑severity** (fix PRs open, addressing two related failure modes):  
- **#2750** – Stale `outbound.db` journals after container kills; hot‑journal poll races. *PR #2750 itself is the fix.*  
  [PR #2750](https://github.com/nanocoai/nanoclaw/pull/2750) (addressing issues #2516, #2640)

**Hardening (low‑severity in practice but proactive):**  
- **#2732** – Multi‑agent health audit findings: bind‑mount realpath resolution, spawn circuit‑breaker, concurrency limits, Docker fallback kill. *PR #2732 is the hardening patch.*  
  [PR #2732](https://github.com/nanocoai/nanoclaw/pull/2732)

No crashes or regressions were reported specifically today; the above bugs are the only stability topics.

## 6. Feature Requests & Roadmap Signals
User‑observed gaps and maintainer‑driven features suggest the following are likely to land in the next release:

- **Budget/billing feedback** – Issue #2751 strongly implies users want explicit error messages rather than silent failure. Already worked on in PR #2759.
- **Operator‑driven provider switching** – PR #2756 introduces a provider registry, setup picker, installer, and memory migration skill. This will allow users to switch between backend model providers at runtime.
- **Codex as a full provider** – PR #2757 delivers a v2 Codex payload that runs on the host’s capability seams with vault‑only authentication.
- **Persistent memory** – Merged PR #2745 provides an opt‑in scaffold; future releases will likely make it the default and tie it to provider selection.

The **closed seam‑based PRs (#2746, #2745)** were essential groundwork; the **open PRs (#2756, #2757)** are the payload. This points to a **provider‑switching UI/UX** being the headline feature of the next minor or major release.

## 7. User Feedback Summary
The only explicit user report today is **frustration with silent drops** (issue #2751). The user expected to be notified when a turn was cut off due to budget limits. This is a clear pain point for anyone running budget‑constrained agents.

Additional implicit feedback can be inferred from the **volume of provider‑related work**: users and maintainers are actively pushing toward a multi‑provider ecosystem (OpenCode, Codex, Anthropic direct, OneCLI gateway). The merged seam work suggests satisfaction with the modular direction, but the delay in delivering actual provider payloads (still open) may cause some impatience.

No praise or negative sentiment was recorded in the 24‑hour window beyond the bug report.

## 8. Backlog Watch
No issues or PRs are dangerously stale; the oldest open item updated in the last 24h is **PR #2732** (health audit, opened June 11 – 3 days ago). It still requires review/merge. **PR #2750** (outbound.db fix, opened June 12) is also awaiting attention. Neither is critically overdue, but maintainers should aim to merge these before the week’s end to avoid accumulating open PRs.

- **PR #2732** – Harden host + agent‑runner from health audit findings (3 days open).  
  [PR #2732](https://github.com/nanocoai/nanoclaw/pull/2732)

- **PR #2750** – Recover stale outbound.db journals; classify poll races (2 days open).  
  [PR #2750](https://github.com/nanocoai/nanoclaw/pull/2750)

All other open items are from June 14 and are being actively worked (e.g., #2759, #2756, #2757). No long‑unanswered issues were identified.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

## NullClaw Project Digest — 2026-06-14

### 1. Today's Overview
Project activity remains low today, with only a single open issue and one open pull request updated in the last 24 hours. No new releases were published. The spotlight is on a critical bug where agent-type scheduled jobs silently fail to deliver messages—the root cause has been identified in a pending PR, but the fix has not yet been merged. Overall, the project appears stable but is currently stalled on a high-impact defect that affects core delivery functionality.

### 2. Releases
No releases were made today. The latest release remains unknown (none listed).

### 3. Project Progress
- **No PRs were merged or closed today.**  
- An open PR ([#954](https://github.com/nullclaw/nullclaw/pull/954)) by vernonstinebaker proposes a fix for the agent-cron bug, but it has not yet been reviewed or merged.

### 4. Community Hot Topics
The most active discussion this period centers on a single open issue:

- **[#941 – Agent-type cron jobs don't spawn a subprocess — Telegram delivery never happens](https://github.com/nullclaw/nullclaw/issues/941)**  
  *Author: weissfl | Created: 2026-05-31 | Updated: 2026-06-13 | Comments: 7 | 👍: 0*  
  This issue has garnered the most comments (7) and describes a complete failure of scheduled agent jobs to execute and deliver messages via Telegram (and presumably other channels). The underlying need is for reliable, unattended task execution—a core feature for users relying on the `schedule` system. The community is clearly following this because it breaks a fundamental use case (e.g., daily reports, alerts). PR #954 directly addresses this problem.

### 5. Bugs & Stability
| Severity | Bug / Issue | Status |
|----------|-------------|--------|
| **High** | [#941 – Agent-type cron jobs don't spawn a subprocess](https://github.com/nullclaw/nullclaw/issues/941) – scheduled jobs marked complete but no delivery occurs. | Open; fix proposed in PR #954 (use-after-free in `OutboundMessage.channel`). |
| Medium | None reported today. | – |
| Low | None reported today. | – |

The sole bug is severe: it causes silent data loss for any user who relies on one-shot or recurring agent tasks. The proposed fix (PR #954) is promising but awaits maintainer review.

### 6. Feature Requests & Roadmap Signals
No new feature requests or roadmap signals were observed today. The only activity is a bug fix, which is likely to be included in the next patch release if merged. No speculative features can be predicted for the next version based on today’s data.

### 7. User Feedback Summary
- **Pain point:** A user (weissfl) reported that scheduled agent tasks silently fail to deliver messages, even when configured with correct delivery parameters (`delivery_mode: "always"`, `delivery_channel: "telegram"`). The job is recorded as completed, but the agent subprocess never starts.  
- **Use case affected:** Unattended cron-like execution of agent tasks, which is a common pattern for automation and alerts.  
- **Satisfaction/Dissatisfaction:** The community has not expressed overt frustration (no reactions on the issue), but the 7 comments suggest active troubleshooting. The existence of a PR indicates that a contributor found the bug reproducible and important enough to fix.

### 8. Backlog Watch
- **[Issue #941 – Agent-type cron jobs don't spawn a subprocess](https://github.com/nullclaw/nullclaw/issues/941)** has been open since 2026-05-31 (14 days) and remains unmerged. Although a fix PR [#954](https://github.com/nullclaw/nullclaw/pull/954) now exists, **maintainer attention is required to review, test, and merge it.** No other long-unanswered issues were identified in today’s data.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-06-14

## 1. Today's Overview
The project saw high activity with **11 open issues** and **21 PRs** updated in the last 24 hours, including **10 merged/closed PRs**. Development momentum is concentrated on two fronts: the **Reborn attachment pipeline** (issue #4644 and its series of stacked PRs) and **Slack delivery reliability** (auth-resume identity, gate delivery, single-flight fixes). Bug reports from the Railway QA environment (tool activity disappearance, invisible shell commands, excessive approval prompts) indicate real-user friction in multi-tenant setups. No new releases were cut today.

## 2. Releases
*No new releases today.*

## 3. Project Progress (Merged/Closed PRs Today)

Ten PRs were merged or closed, spanning attachments, Slack fixes, and housekeeping:

- **[#4839 – fix: preserve invocation identity across auth-gate re-dispatch (Slack re-approval loop)](https://github.com/nearai/ironclaw/pull/4839)** — Fixes the core Slack re-approval loop bug where a capability requiring both approval and credential triggered multiple human approvals. Introduces `pending_auth_resume` identity matching and lease management for `Dispatching` state.
- **[#4680 – fix(openai-compat): stop emitting `[non_text_content]` for non-text parts](https://github.com/nearai/ironclaw/pull/4680)** — Kills the opaque `[non_text_content]` canary that collapsed image/audio/file parts to text, part of attachment pipeline cleanup.
- **[#4677 – feat(threads): fold attachment text into model-visible context](https://github.com/nearai/ironclaw/pull/4677)** — Attachments now reach the model via `ContextMessage` rendered with `<attachments>` blocks.
- **[#4676 – feat(attachments): extract document text on the inbound landing path](https://github.com/nearai/ironclaw/pull/4676)** — `land_inbound_attachments` now fills `AttachmentRef.extracted_text` using the new extractors crate.
- **[#4675 – refactor(extractors): extract file text-extraction into ironclaw_extractors crate](https://github.com/nearai/ironclaw/pull/4675)** — Moves byte→text extraction to a standalone leaf crate, enabling reuse.
- **[#4672 – feat(reborn): accept inline attachment uploads on the WebChat v2 send path](https://github.com/nearai/ironclaw/pull/4672)** — End-to-end ingress wiring for browser file attachments (bytes → storage → transcript refs).
- **[#4670 – feat(attachments): bridge inbound bytes into transcript AttachmentRefs](https://github.com/nearai/ironclaw/pull/4670)** — Connects byte landing to durable transcript references.
- **[#4843 – fix(slack): single-flight gate delivery per run_id (resolution-ack fanout)](https://github.com/nearai/ironclaw/pull/4843)** — Fixes Bug 3 of Slack re-approval-loop triage: gate-resolution ack now respects `submitted_run_id` to avoid duplicate deliveries.
- **[#4242 – chore(deps): bump tar from 0.4.45 to 0.4.46](https://github.com/nearai/ironclaw/pull/4242)** — Security patch for PAX header handling.
- **[#4677](https://github.com/nearai/ironclaw/pull/4677) and [#4676](https://github.com/nearai/ironclaw/pull/4676) also listed above (cross-reference).**

## 4. Community Hot Topics

- **Issue #4644 – Universal attachments across all channels**  
  [nearai/ironclaw#4644](https://github.com/nearai/ironclaw/issues/4644) – The longest-lived epic (created June 9, 1 comment) driving the entire attachment pipeline. It defines three gaps: Reborn’s text-only contract, duplicated format logic across call sites, and missing extensible format registry + polished web UX. The active PR stack (5 PRs merged today) shows strong maintainer investment.

- **Issue #4854 – Excessive approval prompts for simple GitHub extension requests**  
  [nearai/ironclaw#4854](https://github.com/nearai/ironclaw/issues/4854) – A user reports that a read-only GitHub query required multiple approval dialogs, indicating an approval-identity or gate-filtering regression.

- **Issue #4853 / #4852 – Tool activity disappears after completion; shell command not visible**  
  [#4853](https://github.com/nearai/ironclaw/issues/4853) and [#4852](https://github.com/nearai/ironclaw/issues/4852) – Both from the same user (sunglow666) on Railway QA and local main. Shell activity not shown in approval dialog or activity history, and tool activity count disappears after completion. Likely a frontend state-management issue in Reborn’s multi-tenant render path.

- **PR #4841 – reborn: no run-borking failures**  
  [nearai/ironclaw#4841](https://github.com/nearai/ironclaw/pull/4841) – XL-sized PR aiming to explain or recover from terminal errors (`HostUnavailable`, model failures). High community interest as it affects run reliability.

- **PR #4842 – test(reborn): record QA-trace phrases to terminal-or-gate**  
  [nearai/ironclaw#4842](https://github.com/nearai/ironclaw/pull/4842) – Fixes QA-trace recorder hanging on auth gates; directly tied to the user-reported issues.

Underlying need: Users want **reliable, visible, and minimally intrusive** tool execution in Reborn, especially in multi-tenant environments. The Slack re-approval loop and the QA-trace hang are symptoms of a broader capability-resumption design that is being refactored.

## 5. Bugs & Stability

Ranked by severity:

| Severity | Issue | Description | Fix PR(s) |
|----------|-------|-------------|-----------|
| **High** | [#4854](https://github.com/nearai/ironclaw/issues/4854) | Excessive approval prompts for simple read-only GitHub requests — security UX regression. | Open |
| **High** | [#4108](https://github.com/nearai/ironclaw/issues/4108) | Nightly E2E scheduled run failed (full E2E + extensions). Persistent since May 27. | Not yet |
| **Medium** | [#4853](https://github.com/nearai/ironclaw/issues/4853) | Tool activity disappears after completion on Railway (multi-tenant). | Open |
| **Medium** | [#4852](https://github.com/nearai/ironclaw/issues/4852) | Shell command not visible in approval dialog or Activity history. | Open |
| **Low** | [#4851](https://github.com/nearai/ironclaw/issues/4851) | Trusted-trigger origin laundered through `adapter_kind` string – type safety gap. | Open |
| **Low** | [#4849](https://github.com/nearai/ironclaw/issues/4849) | Auth-resume leases orphaned in `Dispatching` by mid-dispatch crash (TTL/sweep needed). | Open |
| **Low** | [#4848](https://github.com/nearai/ironclaw/issues/4848) | `pending_auth_resume` matched only by `capability_id`, not per-invocation identity (`input_ref`). | Open |

Fix PRs in progress: [#4844](https://github.com/nearai/ironclaw/pull/4844) (gate filter fix), [#4843](https://github.com/nearai/ironclaw/pull/4843) (single-flight gate delivery), [#4838](https://github.com/nearai/ironclaw/pull/4838) (explicit gate feedback), [#4841](https://github.com/nearai/ironclaw/pull/4841) (run-borking recovery), [#4842](https://github.com/nearai/ironclaw/pull/4842) (QA-trace hang).

## 6. Feature Requests & Roadmap Signals

- **Universal attachments (epic #4644)** – The most visible roadmap item. With five PRs merged today (pipeline crates, extraction, web upload), the v1 bridge is nearly complete. The remaining pieces: wire into Reborn’s `MessageContent`, extensible format registry, and polished WebChat UX (PR #4738 open). Likely to be in the next minor release.
- **Runtime context awareness** – PR [#4836](https://github.com/nearai/ironclaw/pull/4836) (open) surfaces connected channels, delivery state, and run origin to the model. This is foundational for context-aware tool selection.
- **Gated final-answer nudge** – PR [#4837](https://github.com/nearai/ironclaw/pull/4837) (open) adds a tool-free model call when the agent ends a turn empty or with `NoProgressDetected`, improving conversational quality.
- **Slack connected state persistence** – PR [#4777](https://github.com/nearai/ironclaw/pull/4777) (open) fixes the Slack reconnect loop in WebUI by persisting delivery connection state.
- **Steer routine delivery through outbound targets** – PR [#4780](https://github.com/nearai/ironclaw/pull/4780) (open) adds model guidance to select delivery targets before creating routines/triggers.

Predictions: The next version (v0.29.x?) will include the full attachment pipeline (backend + basic WebChat UX), the Slack re-approval loop fix, and possibly the runtime-context surface. The “no run-borking” (#4841) and QA-trace (#4842) improvements may also land.

## 7. User Feedback Summary

- **Pain point: excessive approval prompts** – User *sunglow666* reports simple read-only GitHub requests generating multiple approval dialogs (issue [#4854](https://github.com/nearai/ironclaw/issues/4854)). This degrades trust and usability.
- **Pain point: missing tool visibility** – Same user reports shell commands are not shown in approval dialogs or activity history (issues [#4852](https://github.com/nearai/ironclaw/issues/4852), [#4853](https://github.com/nearai/ironclaw/issues/4853)). This undermines the transparency expected from an agent framework.
- **Multi-tenant instability** – The `Railway QA` environment is suffering from tool activity disappearing after completion and missing tool command previews, suggesting race conditions or state reset issues in multi-tenant deployments.
- **Satisfaction signals** – No explicit positive feedback in today’s data, but the high volume of merged attachment-related PRs indicates rapid progress on a long-requested feature.

## 8. Backlog Watch

| Item | Type | Age | Status | Notes |
|------|------|-----|--------|-------|
| [#4108 – Nightly E2E failed](https://github.com/nearai/ironclaw/issues/4108) | Bug | Since May 27 (18 days) | Open, no comments | Automated bot report with no human follow-up. E2E CI is consistently red; may block releases. |
| [#3708 – chore: release](https://github.com/nearai/ironclaw/pull/3708) | PR | Since May 16 (29 days) | Open | Release PR with API breaking changes (`ironclaw_common` 0.5.0, `ironclaw_skills` 0.4.0). Stale; merging would cut a new version incorporating many fixes. |
| [#4644 – Universal attachments](https://github.com/nearai/ironclaw/issues/4644) | Epic | Since June 9 (5 days) | Active | Well-tracked with stacked PRs; not yet complete but actively worked. |

The nightly E2E failure (#4108) is the most concerning backlog item—it has been red for 18 days with no response from maintainers, which risks regressions going undetected. The release PR (#3708) is also stale despite containing breaking changes that likely encompass recent work.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-06-14

## Today's Overview

The project shows a quiet day with no new releases and only two issues and four pull requests updated in the last 24 hours. All of these items were originally created in early April 2026 and carry the `[stale]` label, indicating they have been lingering without recent maintainer activity until today. One PR (#1465) was finally closed/merged, suggesting some progress on a recurring bug. Overall, the repository is in a maintenance phase with low visible activity, but the fact that stale items are being touched could signal a cleanup push.

## Releases

No new releases were published today. The latest official release remains unknown from this data snapshot.

## Project Progress

One pull request was merged/closed today:

- **#1465 (merged/closed)** — `fix(scheduled-tasks): 已删除的定时任务重启后作为幽灵会话重新出现`  
  This fix addresses a bug where deleted scheduled tasks reappeared as ghost sessions after a restart. The root cause was a missing cleanup of the local SQLite `cowork_sessions` record when a cron task was removed. The PR also links to issue #1359.  
  [GitHub](https://github.com/netease-youdao/LobsterAI/pull/1465)

No other PRs were merged; three remain open.

## Community Hot Topics

The most active items today are the two open issues, each with one comment (likely from the maintainer or author). They have no reactions, but their UI/UX nature makes them user-facing pain points:

- **#1434** — *[stale] Language setting to Chinese shows English empty-state text and buttons in agent skill tab*  
  Users expect full i18n consistency. The screenshot shows English fallback when search yields no results.  
  [GitHub](https://github.com/netease-youdao/LobsterAI/issues/1434)

- **#1435** — *[stale] Custom agent name too long overflows the modal UI*  
  A straightforward UI overflow issue that degrades the creation experience.  
  [GitHub](https://github.com/netease-youdao/LobsterAI/issues/1435)

Both issues were created over two months ago and have received no visible maintainer response until today's update (likely just a stale-bot touch). The underlying need is clear: polish of high‑touch UI elements for international users and better input validation.

## Bugs & Stability

- **Severity: Medium** — #1434 (i18n gap for Chinese empty-state) – No fix PR associated.
- **Severity: Low** — #1435 (UI overflow for long names) – No fix PR associated.
- **Severity: High (resolved)** — #1465 (ghost sessions after deleting scheduled tasks) – Fixed and merged today. This was a data‑consistency bug potentially causing user confusion.

No new bugs or crashes were reported today.

## Feature Requests & Roadmap Signals

Three open PRs, all stale from April, represent features that may land in the next release:

- **#1429** — In‑session message search with `mark.js` highlighting (Ctrl+F) for Cowork view.  
- **#1430** — Automatic prevention of system sleep during active sessions using Electron `powerSaveBlocker`.  
- **#1431** — Real‑time elapsed timer displayed in the StreamingActivityBar.

These features improve the Cowork experience for long‑running agent tasks. The absence of recent activity suggests they may be blocked or waiting for review. If merged, they would be strong additions to the next minor release.

## User Feedback Summary

The only voiced user pain points today come from the two open issues:

- **i18n inconsistency**: Users setting LobsterAI to Chinese encounter untranslated UI elements (English empty‑state text and buttons). This points to a broader need for thorough localization testing.
- **UI overflow**: Long agent names break out of modal dialogs, indicating missing truncation or responsive design.

No satisfaction feedback or positive use‑case mentions were observed in this snapshot.

## Backlog Watch

All four open items (two issues, two PRs) are marked `[stale]` and were created in early April 2026. They have not been actively worked on for over two months. Notable:

- **#1429** (message search) and **#1430** (sleep prevention) are feature PRs that would benefit from maintainer review. Their stagnation may be a bottleneck for the Cowork feature set.
- **#1434** and **#1435** are small, well‑defined UI bugs that could be good candidates for first‑time contributors or quick fixes.

The project maintainers should evaluate whether these items can be closed or advanced, as a large stale backlog may erode contributor confidence.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest — 2026-06-14

## 1. Today's Overview

Activity on the Moltis repository is moderate, with two open issues and three open pull requests updated in the last 24 hours. No new releases were published. The project is actively addressing a critical OAuth interoperability bug affecting remote MCP servers (Notion, Linear) — a fix is already in review. Additionally, a user-contributed Dockerfile improvement and a dependency update are pending, alongside a new feature request for an alternative memory backend. Overall, the project shows healthy development momentum with a clear focus on edge‑case fixes and infrastructure hardening.

## 2. Releases

*No new releases were published on 2026-06-14.*

## 3. Project Progress

No pull requests were merged or closed in the last 24 hours. Three PRs remain open:

- **[#1122](https://github.com/moltis-org/moltis/pull/1122)** — **fix: drop VOLUME declarations that shadow the home bind mount**  
  Addresses a Dockerfile issue where declared `VOLUME` paths prevented bind‑mount overrides for the entire home directory. Submitted by contributor `sayotte`.

- **[#1121](https://github.com/moltis-org/moltis/pull/1121)** — **chore(deps-dev): bump esbuild from 0.25.12 to 0.28.1**  
  Automated dependency update by Dependabot for the web UI crate.

- **[#1120](https://github.com/moltis-org/moltis/pull/1120)** — **fix(mcp): use direct fetch for resource_metadata URL from WWW-Authenticate**  
  Proposed fix for bug #1119 (see *Bugs & Stability*). Implements a direct fetch instead of passing the URL through an intermediate method that caused `invalid_target` errors.

## 4. Community Hot Topics

The most actively discussed item is **Issue #1119** ([Bug]: MCP OAuth fails with `invalid_target` for servers using `resource_metadata` in WWW-Authenticate](https://github.com/moltis-org/moltis/issues/1119)), which received one comment and has an associated open fix PR (#1120). This bug directly impacts users of popular third‑party MCP services like Notion and Linear, making it a high‑visibility issue. The underlying need is for robust OAuth flow handling when servers embed non‑standard parameters in authentication challenge headers.

No other issues or PRs currently have comments or reactions, indicating that the community’s attention is concentrated on this bug and its resolution.

## 5. Bugs & Stability

One bug was reported in the last 24 hours, ranked **critical** due to its impact on core functionality:

- **Issue #1119** — [MCP OAuth fails with `invalid_target` for servers using `resource_metadata` in WWW-Authenticate (Notion, Linear)](https://github.com/moltis-org/moltis/issues/1119)  
  **Severity:** High – prevents adding remote MCP servers that require OAuth with a `resource_metadata` parameter (e.g., Notion, Linear).  
  **Status:** A fix PR (#1120) is already open and awaiting review. The error manifests in the browser during authorization, returning a JSON error. The proposed solution replaces an indirect URL pass through `fetch_resource_metadata()` with a direct fetch.

No crashes, regressions, or other stability issues were reported.

## 6. Feature Requests & Roadmap Signals

One new feature request was opened:

- **Issue #1123** — [[Feature]: Add pure-Rust turbovec as an alternative memory backend for extreme edge compression](https://github.com/moltis-org/moltis/issues/1123)  
  Proposed by `joeblew999`, this request asks for a drop‑in memory backend using the `turbovec` Rust crate, targeting extreme compression for edge deployments. The request includes a preflight checklist and implies a need for lower storage overhead or faster serialization for constrained environments.  
  **Prediction:** If this aligns with the project’s roadmap for lightweight deployments, it could be considered for a future minor release. No maintainer response yet.

No other feature requests were updated.

## 7. User Feedback Summary

User feedback this period is centred around two technical pain points:

- **OAuth interoperability (Issue #1119):** Users attempting to connect to Notion and Linear MCP servers encounter a broken OAuth flow. The error `invalid_target` suggests a parsing mismatch in how the `resource_metadata` URL from `WWW-Authenticate` headers is handled. This indicates that the current implementation assumes a simpler OAuth negotiation format.

- **Docker deployment (PR #1122):** Contributor `sayotte` identified that declaring `VOLUME` instructions for subdirectories (`/home/moltis/.config`, `.moltis`, etc.) prevents users from bind‑mounting the entire home directory — a common deployment pattern. The fix removes those explicit `VOLUME` declarations, reflecting a desire for more flexible container configurations.

Overall, the community appears technically engaged, with contributions addressing real‑world use‑case friction.

## 8. Backlog Watch

No long‑unanswered issues or PRs are apparent in the 24‑hour window. All open items are recent (created or updated within the last 2 days). No items currently require escalated maintainer attention beyond the normal review cycle.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest — 2026-06-14

## 1. Today's Overview

CoPaw is experiencing a high level of community activity, with **9 issues updated in the last 24 hours** (8 open, 1 closed) and **12 pull requests updated** (11 open, 1 merged/closed). The project remains in a rapid development phase, with several first-time contributors submitting fixes and features. A major focus this week is on improving stability and reliability across core components (cron/heartbeat agents, context management, backup, and channel integration). User feedback highlights persistent pain points around session persistence, context compression, and dependency installation, but the community is actively engaging through detailed bug reports and enhancement suggestions. No new releases were published today.

---

## 2. Releases

**No new releases** in the last 24 hours. The project remains at the previously reported versions.

---

## 3. Project Progress

Only one PR was merged/closed today:

- **[PR #2498](https://github.com/agentscope-ai/QwenPaw/pull/2498)** *(merged/closed)* — `fix(agents): use console language when creating agent and fallback unsupported langs`  
  Ensures that newly created agents respect the user’s UI language and fall back gracefully if a requested language is unsupported. This addresses a long-standing inconsistency between console language and agent configuration.

In addition, one issue was closed:

- **[Issue #5172](https://github.com/agentscope-ai/QwenPaw/issues/5172)** *(closed)* — *“聊天总出现问完问题没反应一直等待”* (Chat hangs indefinitely after a period of inactivity). The issue was reproduced and appears to have been resolved, though the specific fix is not yet merged as a PR (the reporter acknowledged a workaround).

Other PRs remain open (see *Backlog Watch*).

---

## 4. Community Hot Topics

| Item | Type | Comments | Summary |
|------|------|----------|---------|
| [#5156](https://github.com/agentscope-ai/QwenPaw/issues/5156) | Issue | 5 | Request to add `kimi-for-coding` to the `uv` (accelerated Python resolver) whitelist so that users with a Kimi coding subscription can use their existing paid capacity directly within CoPaw. |
| [#5173](https://github.com/agentscope-ai/QwenPaw/issues/5173) | Issue | 2 | Feature request (description incomplete, but tagged as enhancement). |
| [#5169](https://github.com/agentscope-ai/QwenPaw/issues/5169) | Issue | 2 | Request for Vietnamese UI language support. |
| [#5182](https://github.com/agentscope-ai/QwenPaw/issues/5182) | Issue | 1 | Proposal to unify model configuration (text, vector, audio/video) under a single schema. |
| [#5181](https://github.com/agentscope-ai/QwenPaw/issues/5181) | Bug | 1 | Plugin dependency installation causes persistent cmd window popups. |

**Analysis**: The most active thread (**#5156**) reveals an unmet integration need — users who have paid for `kimi-for-coding` cannot use their subscription inside CoPaw because the tool only supports the standard Kimi API. This reflects a broader demand for compatibility with third-party coding models and accelerated package managers. The Vietnamese language request (**#5169**) already has a corresponding PR (**#5175**) from a first-time contributor, demonstrating healthy community-driven localization efforts.

---

## 5. Bugs & Stability

### Critical / High Severity
- **[Issue #5172](https://github.com/agentscope-ai/QwenPaw/issues/5172)** — *(Closed)* Chat session hangs indefinitely after a period of inactivity. User must hit “Stop” to recover. The issue was marked closed, but no fix PR is linked; it may have been a configuration or environment-specific problem.
- **[Issue #5181](https://github.com/agentscope-ai/QwenPaw/issues/5181)** — Plugin installation via `pip` spawns visible `cmd.exe` windows and retries infinitely if PyPI is unavailable. **No fix PR yet.** Severity: high for users with unreliable network.
- **[Issue #5177](https://github.com/agentscope-ai/QwenPaw/issues/5177)** — DingTalk channel messages are not registered in `chats.json`, making them invisible in the console frontend. A detailed bug report with logs was provided. **No fix PR yet.** Severity: medium-high for DingTalk users.
- **[Issue #5171](https://github.com/agentscope-ai/QwenPaw/issues/5171)** — Context compression can reduce context to zero when the persona file token count exceeds the retain threshold, causing task failure. **No fix PR yet.** Severity: high for agents with large persona files.
- **[Issue #5174](https://github.com/agentscope-ai/QwenPaw/issues/5174)** — Cron/heartbeat agents cannot perform `write_file` or `spawn_subagent` due to timeout or autonomy limitations. **No fix PR yet**, but a proposed fix exists in PR #5180 (see below).

### Fix PRs in Review
Several PRs from first-time contributor `ly-wang19` target stability improvements:

| PR | Issue(s) Addressed | Severity |
|----|-------------------|----------|
| [#5035](https://github.com/agentscope-ai/QwenPaw/pull/5035) – llama.cpp build number parsing (fixed-width slice) | Minor | Low (failure after build number >9999) |
| [#5037](https://github.com/agentscope-ai/QwenPaw/pull/5037) – Empty `Exec=` line in Linux browser detection | Medium (UI crash) | Medium |
| [#5038](https://github.com/agentscope-ai/QwenPaw/pull/5038) – Empty message list in `LightContextManager` | Medium (crash) | Medium |
| [#5040](https://github.com/agentscope-ai/QwenPaw/pull/5040) – Invalid job in `jobs.json` causes total failure | High (loss of all cron jobs) | High |
| [#5041](https://github.com/agentscope-ai/QwenPaw/pull/5041) – Unreadable files abort entire backup | High (data loss risk) | High |

All five PRs have been “Under Review” since June 9 and have not received maintainer feedback for at least 1 day.

Additionally, three PRs from `nguyenthanhthe` opened today address related issues:
- **[#5180](https://github.com/agentscope-ai/QwenPaw/pull/5180)** – Increase cron/heartbeat timeout and add autonomous context prompt (fixes the core of #5174).
- **[#5176](https://github.com/agentscope-ai/QwenPaw/pull/5176)** – Fix horizontal overflow of approval command text (UX bug, low severity).
- **[#5178](https://github.com/agentscope-ai/QwenPaw/pull/5178)** – Add session filter by title (feature, not a bug fix).

---

## 6. Feature Requests & Roadmap Signals

### High-Interest Requests
1. **Kimi-for-Coding integration** ([#5156](https://github.com/agentscope-ai/QwenPaw/issues/5156)) – Adding `kimi-for-coding` to the `uv` whitelist. This would unlock paid Kimi coding subscriptions for CoPaw users. A natural fit for the next minor release, likely requiring a quick configuration addition.
2. **Vietnamese UI language** ([#5169](https://github.com/agentscope-ai/QwenPaw/issues/5169)) – Already implemented as PR #5175 (by the same requestor). Likely to be merged soon given the clear implementation and low risk.
3. **Unified model configuration** ([#5182](https://github.com/agentscope-ai/QwenPaw/issues/5182)) – Proposes a single schema for text, vector, and audio/video models. This is a significant architectural change; would be a major version feature if accepted.

### Smaller Enhancements
- Session filter by title ([#5178](https://github.com/agentscope-ai/QwenPaw/pull/5178)) – already a PR; useful for power users.
- Multi-agent collaboration skill trigger keywords expansion ([#5179](https://github.com/agentscope-ai/QwenPaw/pull/5179)) – improves reliability of “team collaboration” mode.

**Prediction for next release**: Vietnamese language support, session title filter, and the cron/heartbeat timeout fix are the most likely candidates given they have clean PRs with no breaking changes.

---

## 7. User Feedback Summary

### Pain Points (real user quotes paraphrased)
- **Chat hangs after idle** (#5172): *“After chatting, wait a bit, then ask a question → infinite wait. Must press Stop to recover. For QQ/WeChat channels there’s no Stop button, so the agent becomes dead.”*
- **Context loss** (#5171): *“When persona file token exceeds the retain threshold, compression reduces context to 0 → task is completely interrupted.”*
- **Plugin installation spamming** (#5181): *“v1.1.11.post2: plugin dependencies cause endless popups; can’t close them without killing the process.”*
- **DingTalk sessions invisible** (#5177): *“Agent replies, state saves, but `chats.json` has no record → no conversation in console frontend.”*
- **Cron/heartbeat not doing real work** (#5174): *“Cron agent runs Python script but can’t write files or spawn subagents. Heartbeat only follows a checklist.”*

### Positive Signals
- Users are actively contributing: three first-time PRs today.
- The Vietnamese language request was quickly implemented by a community member.
- Detailed bug reports with logs, screenshots, and reproduction steps indicate a technically engaged user base.

### Overall Sentiment
Frustration about persistent stability issues (hangs, context loss, channel incompatibility) is balanced by appreciation for the project’s rapid iteration and openness to community contributions.

---

## 8. Backlog Watch

### Long-Unanswered PRs (Need Maintainer Attention)
| PR | Author | Created | Last Updated | Status | Risk |
|----|--------|---------|--------------|--------|------|
| [#5035](https://github.com/agentscope-ai/QwenPaw/pull/5035) | ly-wang19 | Jun 9 | Jun 13 | Under Review | Low (fix llama.cpp parsing) |
| [#5037](https://github.com/agentscope-ai/QwenPaw/pull/5037) | ly-wang19 | Jun 9 | Jun 13 | Under Review | Medium (browser detection crash) |
| [#5038](https://github.com/agentscope-ai/QwenPaw/pull/5038) | ly-wang19 | Jun 9 | Jun 13 | Under Review | Medium (empty msg crash) |
| [#5040](https://github.com/agentscope-ai/QwenPaw/pull/5040) | ly-wang19 | Jun 9 | Jun 13 | Under Review | High (cron job corruption) |
| [#5041](https://github.com/agentscope-ai/QwenPaw/pull/5041) | ly-wang19 | Jun 9 | Jun 13 | Under Review | High (backup failure) |
| [#5170](https://github.com/agentscope-ai/QwenPaw/pull/5170) | ly-wang19 | Jun 13 | Jun 13 | Open | Performance (agent list caching) |

All five PRs from `ly-wang19` have been open for ~5 days without maintainer review. Given that they address critical stability issues (backup corruption, cron job failure, crashes), they should be prioritized. The performance improvement PR (#5170) is less urgent.

### Open Issues with No Activity
- **[#5156](https://github.com/agentscope-ai/QwenPaw/issues/5156)** – 2 days old, 5 comments, no maintainer response. A simple configuration addition could resolve this; a comment from the team would be reassuring.
- **[#5171](https://github.com/agentscope-ai/QwenPaw/issues/5171)** – 1 day old, 1 comment. Context compression bug is severe but has not yet received any maintainer reply.
- **[#5174](https://github.com/agentscope-ai/QwenPaw/issues/5174)** – 1 day old, 1 comment. The proposed fix (#5180) was opened today but has not been acknowledged.

**Action recommended**: Review and merge/comment on the five `ly-wang19` PRs as soon as possible to avoid accumulation of risk.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-06-14

## 1. Today's Overview

The project saw **extremely high activity** with **50 pull requests updated** in the last 24 hours (39 open, 11 merged/closed) and **4 issues updated** (3 open, 1 closed). No new releases were published. Development velocity is concentrated on bug fixes, stability improvements, and incremental feature work, with multiple contributors driving concurrent patches. The merged/closed PRs focus on code quality (replacing `unwrap()` with `expect()`, removing `println!` in daemon paths) and targeted bug fixes in the gateway and cron subsystems. The community is actively discussing the completed unification of the three agent turn engines (RFC #7415) and the ongoing ACP Bridge tracker (#6823).

## 2. Releases

None.

## 3. Project Progress – Merged/Closed PRs Today (11 total)

The following 6 closed PRs appear in the top-20 list; the remaining 5 are from the other 30 PRs not shown:

- **#7607** – `fix(gateway): replace HeaderValue unwrap() with .expect() in api_config`  
  *(gateway)* – Replaced bare `unwrap()` calls with documented `.expect()` for etag handling.  
  [zeroclaw-labs/zeroclaw PR #7607](https://github.com/zeroclaw-labs/zeroclaw/pull/7607)

- **#7606** – `fix(gateway): replace Mutex lock unwrap() with .expect()`  
  *(agent, gateway, runtime)* – Improved lock failure documentation in `EventBuffer` and tool execution.  
  [zeroclaw-labs/zeroclaw PR #7606](https://github.com/zeroclaw-labs/zeroclaw/pull/7606)

- **#7605** – `fix(runtime): replace println! with zeroclaw_log::record! in cron store`  
  *(cron, runtime)* – Removed banned `println!` in daemon cron management path.  
  [zeroclaw-labs/zeroclaw PR #7605](https://github.com/zeroclaw-labs/zeroclaw/pull/7605)

- **#7604** – `fix(tools): replace partial_cmp().unwrap() with total_cmp() in calculator`  
  *(tool)* – Fixed potential NaN panic in `calc_median`/`calc_percentile` by using `total_cmp`.  
  [zeroclaw-labs/zeroclaw PR #7604](https://github.com/zeroclaw-labs/zeroclaw/pull/7604)

- **#7599** – `fix(runtime): replace println! with zeroclaw_log::record! in daemon startup`  
  *(daemon, runtime)* – Removed 6 `println!` calls from daemon startup banner.  
  [zeroclaw-labs/zeroclaw PR #7599](https://github.com/zeroclaw-labs/zeroclaw/pull/7599)

- **#7417** – `fix(cron): expose all fields in cron job edit modal to match add form (#6891)`  
  *(bug, size:M, risk:medium)* – Backend and frontend updates to make cron editing fully feature-complete.  
  [zeroclaw-labs/zeroclaw PR #7417](https://github.com/zeroclaw-labs/zeroclaw/pull/7417)

Additionally, the **RFC #7415** (Unify the three agent turn engines) was **closed** today, with maintainer direction to execute as a single consolidation PR (see implementation update comment). This represents a major architecture completion.

## 4. Community Hot Topics

- **#7415 (CLOSED RFC)** – *Unify the three agent turn engines*  
  Comments: 5 | Author: Nillth | Updated: 2026-06-14  
  This RFC sparked significant discussion. The maintainer decided to implement it as a single consolidation PR (#7540) rather than a phased migration, signaling a push for rapid architectural simplification.  
  [zeroclaw-labs/zeroclaw Issue #7415](https://github.com/zeroclaw-labs/zeroclaw/issues/7415)

- **#6823 (OPEN)** – *[Tracker]: Zerocode ACP Bridge*  
  Comments: 2 | Author: singlerider | Updated: 2026-06-14  
  This long-running tracker for the client-side connection layer (TUI ↔ daemon over RPC) remains active. The 2 comments likely reflect implementation progress updates.  
  [zeroclaw-labs/zeroclaw Issue #6823](https://github.com/zeroclaw-labs/zeroclaw/issues/6823)

No other issues or PRs in the top-20 list have visible comment counts (data shows `Comments: undefined`), indicating that most activity is in code contributions rather than discussion.

## 5. Bugs & Stability

Reported bugs (by severity):

- **S1 – Workflow blocked:**
  - `#7542` (referenced in PRs #7618, #7613) – `ask_user` fails instantly with “Channel closed before” from a WebSocket dashboard session. Two fix PRs submitted today by xuwei-xy, both under review.  
  - `#7551` (referenced in PR #7615) – Free-form `ask_user` not supported by WebSocket approval channel, misleading error. Fix proposes `supports_free_form_ask() = false`.

- **S2 – Degraded behavior:**
  - **#7591** – `zeroclaw quickstart` does not invalidate capitalized agent aliases early; after failure, all previous input is discarded, requiring restart from scratch.  
    Author: ax-fe | Created/Updated: 2026-06-14 | Fix PR #7609 (open) adds validation at input time.  
    [zeroclaw-labs/zeroclaw Issue #7591](https://github.com/zeroclaw-labs/zeroclaw/issues/7591)

- **Other fixes in progress:**
  - #7546 (open, risk:high) – Unify `SopEngine` construction to single instance per daemon to fix duplicate state.  
  - #7547 (open, risk:high) – Auto-include discovered MCP tools in `risk_profile allowed_tools` after default `mcp.enabled` flip.  
  - #7617 (open) – Warn on extra-nested provider aliases that silently drop fields.  
  - #7616 (open) – Strip assistant reasoning on outbound replay for Groq provider.  
  - #7614 (open) – Detect musl libc for correct Linux target triple in install script.

## 6. Feature Requests & Roadmap Signals

- **#7611 (OPEN)** – *Support markdown in telegram*  
  Author: kshcherban | Created: 2026-06-14  
  Request to enable Telegram’s new bot markdown formatting as default. Given the project’s active channel support and the PR track record of similar enhancements, this is likely to be implemented in the next minor release.  
  [zeroclaw-labs/zeroclaw Issue #7611](https://github.com/zeroclaw-labs/zeroclaw/issues/7611)

- **#7162 (OPEN enhancement)** – *feat(channels): notify before context compression*  
  Author: sbenedetto | Created: 2026-06-03 | Updated: 2026-06-14  
  Adds user-visible notice before proactive context compression. Open for 11 days, no maintainer comments yet.  
  [zeroclaw-labs/zeroclaw PR #7162](https://github.com/zeroclaw-labs/zeroclaw/pull/7162)

- **#7170 (OPEN enhancement)** – *feat(slack): upload outbound attachments*  
  Author: sbenedetto | Created: 2026-06-03 | Updated: 2026-06-14  
  Adds support for `[IMAGE:]` and `[FILE:]` markers in Slack messages. Also open without maintainer feedback.

- **#6823 (ACP Bridge tracker)** – Remains a roadmap item, expected to be a major integration for TUI/remote connectivity.

## 7. User Feedback Summary

Real user pain points observed:

- **Quickstart workflow fragility** (Issue #7591): A simple typo in alias capitalization forces the user to restart the entire `zeroclaw quickstart` process and lose all prior input. This causes significant dissatisfaction for new users setting up from scratch. A fix PR (#7609) is open.

- **`ask_user` tool fails on WebSocket dashboard** (Issue #7542): Users relying on gateway web dashboards find their agent workflows blocked when the tool “Channel closed before” error appears. Two fix PRs are under review.

- **Missing Telegram markdown** (Issue #7611): User request to leverage Telegram’s new formatting options; expected to be well-received.

- **Provider configuration pitfalls** (Issues #7616, #7617): Groq provider compatibility issues and silent field dropping with nested TOML tables indicate users may experience unexpected provider behavior. Fixes are in review.

Overall satisfaction seems moderate – the community is actively contributing fixes, but the number of S1/S2 bugs suggests stability is a current focus area.

## 8. Backlog Watch

- **#6823 – Zerocode ACP Bridge tracker**  
  Open since 2026-05-21 (24 days), last updated 2026-06-14. Labeled `status:accepted` and `status:no-stale`, but no visible code PRs linked. Maintainers may need to assign implementation ownership.  
  [zeroclaw-labs/zeroclaw Issue #6823](https://github.com/zeroclaw-labs/zeroclaw/issues/6823)

- **#7162 & #7170 – Channel enhancements (context compression notice, Slack attachments)**  
  Open since 2026-06-03 (11 days). No maintainer comments or reviews. These are low-risk, well-defined features that could be merged with quick maintainer attention.

- **#7547 – Auto-include MCP tools in risk_profile allowed_tools**  
  Open since 2026-06-12 (2 days), but labeled `risk:high`. Needs review as it affects all agents using risk profiles.  
  [zeroclaw-labs/zeroclaw PR #7547](https://github.com/zeroclaw-labs/zeroclaw/pull/7547)

- **#7546 – Unify SopEngine construction**  
  Also `risk:high` and open since 2026-06-12. Addresses a duplicate instance bug that could lead to inconsistent tool state.  
  [zeroclaw-labs/zeroclaw PR #7546](https://github.com/zeroclaw-labs/zeroclaw/pull/7546)

No issues or PRs older than 30 days appear to be completely abandoned.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*