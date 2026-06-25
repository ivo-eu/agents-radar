# OpenClaw Ecosystem Digest 2026-06-25

> Issues: 305 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-25 10:25 UTC

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

# OpenClaw Project Digest — 2026-06-25

## Today's Overview
Project activity remains very high: **305 issues were updated in the last 24 hours** (245 open/active, 60 closed), and **500 pull requests were updated** (402 open, 98 merged/closed). One new release candidate (`v2026.6.11-beta.1`) is now available. The community is deeply engaged with critical security and reliability bugs—especially around message leaks and silent agent failures. Development velocity is strong, with nearly one hundred PRs merged today, though many high-severity issues remain open and require maintainer attention.

---

## Releases
**New release: `v2026.6.11-beta.1`** ([GitHub Releases](https://github.com/openclaw/openclaw/releases/tag/v2026.6.11-beta.1))  
Key highlights (excerpt from release notes):

- **More capable channel control**: Slack relay mode, native Mattermost `/oc_queue` command, and per-DM model overrides.  
  (Thanks @sjf-oa, @amknight, @xydigit-zt, @thomaszta, @gandalf-at-lerian)

No breaking changes or migration notes were provided in the available snippet; the full changelog should be consulted for any additional updates or deprecations.

---

## Project Progress
Today **98 pull requests were merged or closed**. Notable merges and closures from the top 30 tracked PRs include:

- **#96737** (`fix(llm): trust control escapes in JSON strings`) – closed as “too-many-prs”.  
- **#96731** (`fix(cron): wire result classifier into isolated cron model fallback chain`) – merged, addresses silent drop of results on terminal failures.  
- **#96733** (`fix(telegram): label forum topic sessions`) – merged, improves session identification for Telegram forum topics.  
- **#68936** (`Autofix: add PR review autofix pipeline + Windows daemon`) – closed; adds automation for PR reviews.

Other important *open* PRs that advanced today (still under review) include:

- **#96522** – `fix(cron): persist session-id targets in destination session` – addresses cron response routing.  
- **#96719** – `fix(discord): preserve tool updates when progress preview is empty`.  
- **#96730** – `fix(codex): prefer desktop app-server for Computer Use on macOS`.  
- **#88968** – `fix: prevent memory flush failure from aborting user reply` – ready for maintainer review.  
- **#95899** – `fix(agents): recover from non-resumable assistant tail by re-presenting last user message`.

These fixes indicate focused effort on stability, session recovery, and channel-specific behavior.

---

## Community Hot Topics
The most active discussions (by comment count and reactions) highlight key community concerns:

| Issue | Title | Comments | 👍 | Priority & Impact |
|-------|-------|----------|----|-------------------|
| [#25592](https://github.com/openclaw/openclaw/issues/25592) | Text between tool calls leaks to messaging channels | **32** | 1 | P1, 🦞 diamond lobster, impact:security, message-loss |
| [#9443](https://github.com/openclaw/openclaw/issues/9443) | Request: Prebuilt Android APK releases | **25** | 2 | P2, enhancement |
| [#44925](https://github.com/openclaw/openclaw/issues/44925) | Subagent completion silently lost – no retry, no notification | **20** | 1 | P1, 🦞 diamond lobster, impact:session-state, message-loss |
| [#10659](https://github.com/openclaw/openclaw/issues/10659) | Masked Secrets – prevent agent from accessing raw API keys | **13** | **4** | P1, 🦞 diamond lobster, impact:security |
| [#42475](https://github.com/openclaw/openclaw/issues/42475) | Per-agent cost budget enforcement at gateway level | **12** | 1 | P2, 🦞 diamond lobster, impact:auth-provider |
| [#40215](https://github.com/openclaw/openclaw/issues/40215) | Show cumulative context usage in /usage footer | **5** | **3** | P2, closed but high interest |

**Underlying needs:**  
- **Security and privacy**: Users are demanding that internal processing (tool call text, credentials) never leak to public channels. Masked secrets and hardened routing are top requests.  
- **Reliability**: Silent failures in sub-agent orchestration and session state corruption are causing significant real-world pain.  
- **Platform reach**: Android APK and better cloud deployment docs (e.g., AWS) remain high-demand items.

---

## Bugs & Stability
Several high-severity bugs were reported or updated today. Ranked by severity (P1 = critical):

| Issue | Title | Severity / Impact | Fix PR exists? |
|-------|-------|-------------------|----------------|
| [#25592](https://github.com/openclaw/openclaw/issues/25592) | Text between tool calls leaks to messaging channels | **P1, security, message-loss** | No fix PR (linked PR open) |
| [#44925](https://github.com/openclaw/openclaw/issues/44925) | Subagent completion silently lost | **P1, session-state, message-loss** | No fix PR |
| [#43367](https://github.com/openclaw/openclaw/issues/43367) | Multi-agent orchestration unstable (config overwrites, session-lock failures) | **P1, session-state, message-loss** | No fix PR (linked PR open) |
| [#43661](https://github.com/openclaw/openclaw/issues/43661) | Session hangs on compaction timeout, causing duplicate sends | **P1, crash-loop, message-loss** | No fix PR |
| [#45049](https://github.com/openclaw/openclaw/issues/45049) | Agent loop allows simulated tool calls instead of real invocation | **P1, security, message-loss** | No fix PR |
| [#44905](https://github.com/openclaw/openclaw/issues/44905) | Discord leaks internal tool-call traces (NO_REPLY, JSON args) | **P1, security, message-loss** | No fix PR |
| [#41165](https://github.com/openclaw/openclaw/issues/41165) | Telegram DMs still land in agent:main:main despite #40519 | **P1, session-state, message-loss** | No fix PR (linked PR open) |
| [#40255](https://github.com/openclaw/openclaw/issues/40255) | Regression: User-configured heartbeat prompt no longer respected | **P1, session-state, message-loss** | No fix PR (linked PR open) |

**Key observations:**  
- Many P1 issues involve **message loss or security leaks**, and nearly all lack a submitted fix PR despite being open for weeks.  
- **Regressions** are a recurring pattern (e.g., heartbeat prompt, Telegram routing).  
- No new crashes were reported today, but session hangs and silent failures continue to dominate.

---

## Feature Requests & Roadmap Signals
The following user-requested features received recent updates and are likely candidates for upcoming releases:

| Issue | Title | Likely to ship next? |
|-------|-------|----------------------|
| [#10659](https://github.com/openclaw/openclaw/issues/10659) | Masked Secrets – prevent agent from reading raw API keys | **High** – strong community interest (4 👍) + security governor priority |
| [#42475](https://github.com/openclaw/openclaw/issues/42475) | Per-agent cost budget enforcement at gateway level | **Medium** – increasingly important for production deployments |
| [#42026](https://github.com/openclaw/openclaw/issues/42026) | Distributed Agent Runtime (control plane / compute split) | **Low** – major architectural change, only 3 👍 |
| [#42276](https://github.com/openclaw/openclaw/issues/42276) | Reasoning stream (thinking indicator) | **Medium** – common UX request |
| [#44395](https://github.com/openclaw/openclaw/issues/44395) | Heading-aware chunking + entity extraction for memory | **Medium** – improves memory quality |
| [#43117](https://github.com/openclaw/openclaw/issues/43117) | Multimodal embedding with gemini-embedding-2-preview | **Low-Medium** – niche but innovative |
| [#42877](https://github.com/openclaw/openclaw/issues/42877) | Bounded memory tool with hard character limits | **High** – direct answer to memory bloat |
| [#43794](https://github.com/openclaw/openclaw/issues/43794) | Config encryption for credentials at rest | **High** – security-critical, part of threat model |
| [#40786](https://github.com/openclaw/openclaw/issues/40786) | Add .gitignore-like exclude patterns to backup CLI | **Medium** – practical improvement |
| [#13700](https://github.com/openclaw/openclaw/issues/13700) | Session snapshots (save/load context checkpoints) | **Low-Medium** – power-user feature |

Several of these align with the project’s threat model and user pain points (config encryption, masked secrets) and may be prioritized in the next beta.

---

## User Feedback Summary
Real-world reports from self-hosted users reveal both satisfaction and frustration:

- **Pain points**:  
  - *Message leaks* to Discord/Telegram channels (internal tool calls, credentials) are the most vocal issue (#25592, #44905).  
  - *Silent failures* in sub-agent orchestration (#44925) and session compaction (#43661) erode trust.  
  - *Configuration regressions* (heartbeat prompt overrides, Telegram routing) undermine upgrades.  
  - *Large deployment stalls* – backup commands fail on 4GB+ directories (#42273).  

- **Positive signals**:  
  - Users actively request *advanced features* (AW5 deployment guides, per-agent budgets, multimodal memory), indicating ongoing adoption and scaling.  
  - One user shared a detailed *field report* ([#41372](https://github.com/openclaw/openclaw/issues/41372)) with 25 findings after 4 weeks of production use, demonstrating deep engagement and willingness to contribute upstream.  
  - The community is responsive: issues quickly accumulate comments and reactions.

- **Overall sentiment**: Mixed – high enthusiasm for the project’s capabilities but growing impatience with stability and security issues that affect daily use.

---

## Backlog Watch
The following important issues have been open for an extended period without a fix PR or maintainer decision:

| Issue | Created | Updated | Comments | Status | Reason for concern |
|-------|---------|---------|----------|--------|--------------------|
| [#25592](https://github.com/openclaw/openclaw/issues/25592) | 2026-02-24 | 2026-06-25 | 32 | Open, P1, diamond lobster | No fix PR; text leak is a critical security bug |
| [#9443](https://github.com/openclaw/openclaw/issues/9443) | 2026-02-05 | 2026-06-25 | 25 | Open, P2, enhancement | High community interest, no maintainer decision |
| [#10659](https://github.com/openclaw/openclaw/issues/10659) | 2026-02-06 | 2026-06-25 | 13 | Open, P1, diamond lobster | Feature request with 4 👍; no fix PR |
| [#13751](https://github.com/openclaw/openclaw/issues/13751) | 2026-02-11 | 2026-06-25 | 6 | Open, P1, diamond lobster | Feishu plugin overly broad permissions |
| [#40255](https://github.com/openclaw/openclaw/issues/40255) | 2026-03-08 | 2026-06-25 | 5 | Open, P1, regression | Heartbeat prompt regression; linked PR open but not merged |
| [#40611](https://github.com/openclaw/openclaw/issues/40611) | 2026-03-09 | 2026-06-25 | 6 | Open, P1, diamond lobster | Heartbeat drift causes Telegram blocking; linked PR open but no resolution |

**Maintainer attention needed**: Many P1 issues have “clawsweeper:needs-maintainer-review” and “clawsweeper:needs-product-decision” labels, indicating they are stalled awaiting product direction. The backlog of security-critical bugs without fix PRs is concerning for a project with this level of community activity.

---

*Digest generated from GitHub data for openclaw/openclaw as of 2026-06-25. Data includes 305 issues and 500 PRs updated in the last 24 hours, plus the top 50 issues and top 30 PRs by comment count.*

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report — AI Agent / Personal Assistant Ecosystem

## 1. Ecosystem Overview

The personal AI assistant open-source ecosystem is experiencing an intense period of maturation, characterized by simultaneous pushes for security hardening, multi-platform compatibility, and production-grade reliability. Across eleven tracked projects, the dominant themes are **message leak prevention**, **tool-call execution safety**, and **inter-agent delegation governance**, reflecting a shift from experimental prototypes to real-world deployment concerns. Development velocity remains high—several projects process 30–50 pull requests daily—but a growing backlog of high-severity bugs without fix PRs signals a tension between feature velocity and stability. The ecosystem is consolidating around three architectural models: monorepo reference implementations (OpenClaw), modular plugin frameworks (PicoClaw, NanoClaw), and vertically integrated agent platforms (Hermes Agent, ZeroClaw).

---

## 2. Activity Comparison

| Project | Open Issues (updated last 24h) | Open PRs (updated last 24h) | PRs Merged/Closed (last 24h) | New Release | Health Assessment |
|---------|-------------------------------|-----------------------------|------------------------------|-------------|-------------------|
| OpenClaw | 245 | 402 | 98 | `v2026.6.11-beta.1` | High velocity, critical P1 bugs without fix PRs |
| NanoBot | 14 (20 total updated) | 26 (36 total updated) | 10 | None | Reactive, strong security response |
| Hermes Agent | 18 | 50 | 8 | None | Steady, token optimization focus |
| PicoClaw | 0 issues open (14 closed) | 9 (13 updated) | 4 | None | Stable, security hardening wave |
| NanoClaw | 1 | 15 (19 updated) | 4 | None | Active, security/feature mix |
| NullClaw | 0 | 0 | 0 | N/A | Inactive |
| IronClaw | 18 (22 total updated) | 26 (39 total updated) | 13 | None | Strong bug fix velocity |
| LobsterAI | 1 | 0 (10 merged) | 10 | None | Polish phase, clean backlog |
| TinyClaw | 0 | 0 | 1 | None | Maintenance mode |
| Moltis | 0 | 0 | 0 | N/A | Inactive |
| CoPaw | 10 (14 total updated) | 22 (50 total updated) | 28 | None | Very high activity, regression risk |
| ZeptoClaw | 0 | 0 | 0 | N/A | Inactive |
| ZeroClaw | 7 (10 total updated) | 38 (50 total updated) | 12 | None | Intense pre-release, S0 bug fixed same day |

---

## 3. OpenClaw's Position

**Advantages over peers:**
- **Largest community and velocity**: 500 PRs and 305 issues updated daily, far exceeding the next most active project (CoPaw at 50 PRs). This creates faster bug discovery but also higher noise.
- **Broadest channel coverage**: Slack relay mode, native Mattermost commands, Telegram forum topics, Discord integration—channel breadth exceeds NanoBot and PicoClaw.
- **Release maturity**: `v2026.6.11-beta.1` indicates structured release cadence; most peers (Hermes, ZeroClaw, CoPaw) lack recent releases.

**Technical approach differences:**
- OpenClaw uses a **monorepo reference architecture** where core capabilities are tightly coupled; NanoBot and PicoClaw use more modular plugin systems.
- OpenClaw's diamond-lobster priority tagging system is unique; peers use standard labels or severity-only systems.
- OpenClaw has explicit "clawsweeper" maintainer-review and product-decision labels, indicating formalized governance lacking in peer projects.

**Community size comparison:**
- OpenClaw: 500 PRs/day, 305 issues/day → likely 500+ contributors
- CoPaw: 50 PRs/day, 14 issues/day → ~100–200 contributors
- ZeroClaw: 50 PRs/day, 10 issues/day → ~100 contributors
- Hermes Agent: 50 PRs/day, 18 issues/day → ~80 contributors
- Remaining projects: 10–20 PRs/day → <50 contributors each

**Risk**: OpenClaw's high volume of open P1 bugs without fix PRs (7 critical, each open for months) threatens its leadership position. Peers like ZeroClaw demonstrate faster turnaround on security bugs (S0 fixed same day).

---

## 4. Shared Technical Focus Areas

Five requirements emerge consistently across projects, indicating ecosystem-wide priorities:

| Requirement | Projects Affected | Specific Pain Points |
|-------------|-------------------|----------------------|
| **Message leak prevention** | OpenClaw (#25592, #44905), NanoBot, PicoClaw (security batch) | Tool-call text leaking to public channels; internal processing visible to end users |
| **Silent failure recovery** | OpenClaw (#44925, #43661), Hermes Agent (#52431), ZeroClaw (#5903) | Subagent completions lost; session hangs with no notification; orphan processes accumulating |
| **Multi-platform channel parity** | OpenClaw (Telegram routing regression), NanoBot (Telegram Web issues), PicoClaw (Delta Chat, Matrix), ZeroClaw (Groq compatibility) | Features work on one platform but break on others; inconsistent behavior degrades user trust |
| **Token/cost optimization** | Hermes Agent (73% overhead, lazy schema loading), OpenClaw (#42475 – per-agent budgets), PicoClaw (streaming config) | Fixed tool-schema overhead 3.5–5k tokens per call; high API costs for production deployments |
| **Security hardening for credential management** | OpenClaw (#10659 – masked secrets), NanoClaw (#2799 – CVE-2026-29611), PicoClaw (SSRF, symlink race), ZeroClaw (#8279 – delegate bypass) | Prompt injection can leak API keys; path traversal; delegation allowlist bypass |

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | Hermes Agent | ZeroClaw | CoPaw | PicoClaw / NanoClaw |
|-----------|----------|---------|--------------|----------|-------|---------------------|
| **Target user** | Self-hosted power users, production deployments | Discord/Telegram chatbot hosts | Developer agent orchestrator | Enterprise inter‑agent workflows | Desktop + browser automation | Edge/low‑resource deployments |
| **Core strength** | Channel breadth, community scale | Security responsiveness | Token efficiency, agent-to-agent protocols | Delegation governance, SBOM transparency | Browser automation, Runtime 2.0 | Lightweight, embedded‑friendly |
| **Architecture** | Monorepo reference | Modular plugin | Provider-agnostic orchestration | Policy-enforced delegation | Plugin + Chrome integration | Minimal core + community plugins |
| **Key differentiator** | Highest issue/PR volume, formal priority system | Coordinated security disclosure handling | Lazy tool loading, A2A protocol ambition | Independent delegate mode, in-app upgrade | Browser resource cleanup, Slack channel | Multi-bot support, native E2EE |
| **Risk profile** | Backlog of critical unfixed bugs | Outstanding exec allowlist bypasses | Token overhead not yet addressed | Unresolved MCP process leak (2+ months) | Browser memory leak regression | Few contributors, may lack review bandwidth |

---

## 6. Community Momentum & Maturity

**Tier 1 — High velocity, rapid iteration:**
- **OpenClaw**: Despite bug backlog, community engagement is unmatched. Mature governance process (maintainer-review labels). Most likely to drive ecosystem standards.
- **ZeroClaw**: Security‑focused velocity. S0 bugs fixed same day. Tracked release milestones (v0.8.3). Strong candidate for enterprise adoption.
- **CoPaw**: Second‑highest PR volume. Fast feedback loops (bug → PR same day). However, regression risk is elevated (Chrome leak returned after earlier fix).

**Tier 2 — Steady development, stabilizing:**
- **Hermes Agent**: Strong token optimization research. A2A protocol ambition is forward‑looking. Maintainer‑acknowledged issues but older PRs need review.
- **NanoBot**: Excellent security disclosure practice. Functional improvements (PWA, DingTalk) alongside security work. Good balance.
- **IronClaw**: Rapid bug fix merging (heartbeat deadlock fixed same day). Reborn architecture progressing. Nightly E2E failures are a risk.

**Tier 3 — Polish/maintenance phase:**
- **LobsterAI**: Clean backlog, 10 PRs merged in one day. Focused on cowork plan mode stability rather than new features.
- **PicoClaw**: Security advisory wave resolved. Streaming feature closed. Appears to be hardening before next release.
- **NanoClaw**: Active but low volume. Telegram multi-bot and Matrix E2EE are meaningful, but reviewer bandwidth is thin.

**Tier 4 — Inactive or stalled:**
- **NullClaw**, **Moltis**, **ZeptoClaw**: No activity in 24 hours. Likely abandoned or in extended hibernation. Community should evaluate for code/security decay.
- **TinyClaw**: Single Windows fix merged. Clean backlog but no roadmap signal. May be maintained in name only.

---

## 7. Trend Signals

Five industry trends extracted from community feedback across projects:

**1. Security is the primary adoption blocker**
Users across OpenClaw, NanoBot, PicoClaw, and ZeroClaw are demanding masked secrets, path traversal prevention, and delegation allowlist enforcement. The volume of coordinated security disclosures (eight in one day for NanoBot; ten in one batch for PicoClaw) indicates vulnerability research is accelerating. **For developers**: expect security requirements to dominate feature prioritization in Q3 2026.

**2. Token waste is a cost crisis**
Hermes Agent's finding that 73% of API call tokens are fixed overhead, combined with OpenClaw's per-agent budget requests and PicoClaw's streaming config demand, signals that **unsustainable API costs** will drive architectural changes. Projects that implement lazy/late-binding tool schema injection will have a competitive advantage.

**3. Multi-agent orchestration is the next frontier**
ZeroClaw's independent delegate mode (#8238) and target-profile authority (#7743), combined with Hermes Agent's A2A protocol ambition (#514), and CoPaw's spawn_subagent (#5523), point to a clear shift from single-agent chatbots to **agent networks**. The winning architectures will need robust delegation governance, session recovery, and consistent error propagation across hops.

**4. Platform-specific polish is non-negotiable**
NanoBot's Telegram Web workaround (#4489), CoPaw's Linux browser detection fix (#5528), ZeroClaw's Windows UTF-8 BOM parse fix (#8326) all stem from the same root cause: **cross-platform parity gaps degrade user trust**. Developers building for this ecosystem must prioritize testing on Windows, macOS, Linux, and at least two messaging platforms (Telegram + one other) before ship.

**5. Production readiness requires operational tools**
OpenClaw's distributed agent runtime (#42026), ZeroClaw's in-app upgrade (#8173) and LAN peer discovery (#8325), and IronClaw's local service lifecycle management (#4598) all point to **operational maturity** as a differentiator. The market is moving beyond "it works on my machine" to "it survives a reboot and auto-updates." CI/CD reliability (e.g., IronClaw's 30-day failing nightly E2E) is emerging as a key trust signal.

**Value for AI agent developers**: Invest in security-first design patterns (least privilege for tool calls, credential isolation), implement lazy schema loading regardless of current token budget, and target multi-agent delegation with explicit policy governance early—these will be baseline requirements within 6–12 months.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-06-25

## 1. Today’s Overview
Today saw heavy activity driven by a coordinated disclosure of eight high-priority security vulnerabilities (MCP allowlist bypasses and `exec` shell‑chain bypasses). The team responded quickly, with **10 PRs merged/closed** and **6 issues resolved** in the last 24 hours. Community engagement is high, evidenced by **36 PRs updated** and **20 issues updated**. Despite the security focus, several functional improvements (PWA, DingTalk fixes, Telegram compatibility) also progressed. Overall, the project is in a reactive state but demonstrating strong maintenance velocity.

## 2. Releases
No new releases today. The latest published release remains **v0.2.x** (based on issue context). Security fixes merged today are expected to land in a point release soon.

## 3. Project Progress
**10 PRs merged or closed today** (out of 36 updated). Key merged/closed items:

- **MCP `enabledTools` enforcement** — PR [#4452](https://github.com/HKUDS/nanobot/pull/4452) and [#4436](https://github.com/HKUDS/nanobot/pull/4436) now correctly gate resources and prompts behind the allowlist, addressing security issues [#4434](https://github.com/HKUDS/nanobot/issues/4434) and [#4435](https://github.com/HKUDS/nanobot/issues/4435).
- **Telegram rich‑message config** — PR [#4489](https://github.com/HKUDS/nanobot/pull/4489) (and duplicate [#4505](https://github.com/HKUDS/nanobot/pull/4505)) added a `rich_messages` toggle to fix unsupported‑message errors on Telegram Web.
- **WebUI clipboard fallback** — PR [#4509](https://github.com/HKUDS/nanobot/pull/4509) fixed code‑block copy when the Clipboard API is unavailable.
- **Test suite speedup** — PR [#4507](https://github.com/HKUDS/nanobot/pull/4507) removed avoidable waits, improving CI efficiency.
- **Other** — PR [#4512](https://github.com/HKUDS/nanobot/pull/4512) (test/CLAUDE.md updates), issue [#4503](https://github.com/HKUDS/nanobot/issues/4503) (HVTracker badge, closed), issue [#4499](https://github.com/HKUDS/nanobot/issues/4499) (Telegram empty messages, closed).

## 4. Community Hot Topics
The most active discussions center on **security disclosures** and **Telegram compatibility**:

- **Security disclosures (8 issues)** — All filed by YLChen-007 and spanning MCP allowlist bypass ([#4517](https://github.com/HKUDS/nanobot/issues/4517), [#4519](https://github.com/HKUDS/nanobot/issues/4519) etc.) and `exec` allowlist bypass ([#4521](https://github.com/HKUDS/nanobot/issues/4521), [#4520](https://github.com/HKUDS/nanobot/issues/4520), [#4518](https://github.com/HKUDS/nanobot/issues/4518), [#4516](https://github.com/HKUDS/nanobot/issues/4516), [#4515](https://github.com/HKUDS/nanobot/issues/4515), [#4514](https://github.com/HKUDS/nanobot/issues/4514)). These have **zero to one comments** but high urgency; the maintainers have already merged fixes for the MCP branch.
- **Telegram Web “not supported”** — Issue [#4488](https://github.com/HKUDS/nanobot/issues/4488) (closed) and the rich‑message workaround received quick PRs ([#4489](https://github.com/HKUDS/nanobot/pull/4489), [#4505](https://github.com/HKUDS/nanobot/pull/4505)).
- **DingTalk formatting/timeout** — Issue [#4497](https://github.com/HKUDS/nanobot/issues/4497) with logs showing `ConnectTimeout` and dropped `richText` messages, prompting fix PR [#4501](https://github.com/HKUDS/nanobot/pull/4501).
- **WebUI navigation / streaming** — Issue [#4500](https://github.com/HKUDS/nanobot/issues/4500) reports three distinct WebUI bugs (welcome‑screen send, restart stalling, “No active task to stop”). No fix PR yet.

The underlying need is clear: users demand **both security hardening** and **platform‑specific polish** (Telegram Web, DingTalk rich text, Windows service mode).

## 5. Bugs & Stability
Reported issues ranked by severity:

| Severity | Issue | Description | Fix PR exists? |
|----------|-------|-------------|----------------|
| **Critical** | [#4517](https://github.com/HKUDS/nanobot/issues/4517) + related | Eight security bypasses in MCP allowlist and `exec` allowlist | Yes ([#4452](https://github.com/HKUDS/nanobot/pull/4452), [#4436](https://github.com/HKUDS/nanobot/pull/4436)); `exec` bypasses still open as of today |
| **High** | [#4499](https://github.com/HKUDS/nanobot/issues/4499) | Telegram messages sent as empty (Android app works) | Closed (no explicit fix PR, but likely resolved by rich‑message config) |
| **High** | [#4488](https://github.com/HKUDS/nanobot/issues/4488) | Telegram Web unsupported‑message error | Yes ([#4489](https://github.com/HKUDS/nanobot/pull/4489)) |
| **Medium** | [#4513](https://github.com/HKUDS/nanobot/issues/4513) | Windows `nssm` service restart causes port‑conflict loop or service‑running confusion | No |
| **Medium** | [#4497](https://github.com/HKUDS/nanobot/issues/4497) | DingTalk rich text dropped & HTTP timeout | Yes ([#4501](https://github.com/HKUDS/nanobot/pull/4501)) |
| **Medium** | [#4511](https://github.com/HKUDS/nanobot/issues/4511) | Windows `--background` restart mismatches PID/log files | No |
| **Low** | [#4500](https://github.com/HKUDS/nanobot/issues/4500) | WebUI navigation, streaming stuck, stop‑button failure | No |
| **Low** | [#4492](https://github.com/HKUDS/nanobot/issues/4492) | Xiaomi MiMo ASR transcription fails (WebM→WAV needed) | Yes ([#4493](https://github.com/HKUDS/nanobot/pull/4493)) |

## 6. Feature Requests & Roadmap Signals
Several enhancement PRs and issues indicate likely upcoming features:

| Request | Issue/PR | Likelihood for next release |
|---------|----------|-----------------------------|
| PWA support & mobile sidebar swipe | [#4479](https://github.com/HKUDS/nanobot/issues/4479) | High – PR already open with complete implementation |
| Gateway webhook triggers | [#4502](https://github.com/HKUDS/nanobot/pull/4502) | High – structured PR with config example |
| `ask_clarification` tool for ambiguous requests | [#4508](https://github.com/HKUDS/nanobot/issues/4508) | Medium – detailed proposal, no PR yet |
| Subagent aggregated result mode | [#4414](https://github.com/HKUDS/nanobot/pull/4414) | Medium – open since June 19, may need review |
| `SuspendTurn` for async/human-in-the-loop | [#4411](https://github.com/HKUDS/nanobot/pull/4411) | Medium – new design, needs maintainer feedback |
| Serper.dev web search provider | [#4406](https://github.com/HKUDS/nanobot/pull/4406) | High – straightforward provider addition |
| Custom provider thinking style config | [#4482](https://github.com/HKUDS/nanobot/pull/4482) | High – fixes reported issue [#4429](https://github.com/HKUDS/nanobot/issues/4429) |
| Subagent `fail_on_tool_error` config | [#4485](https://github.com/HKUDS/nanobot/pull/4485) | Medium – PR, addresses [#4198](https://github.com/HKUDS/nanobot/issues/4198) |

The security fixes will almost certainly be backported into a v0.2.x patch release. PWA and DingTalk fixes are also strong candidates for the next release.

## 7. User Feedback Summary
Real pain points and use cases surfaced today:

- **Security is top‑of‑mind** — Multiple users (YLChen-007) responsibly disclosed serious bypasses. The community expects a prompt patch release.
- **Telegram Web users are frustrated** — Rich messages break the Web interface entirely; the team responded with a config toggle.
- **Windows administrators need reliability** — Reports around `nssm` service management and the `--background` flag highlight gaps in Windows tooling.
- **WebUI UX issues** — Navigation, streaming state, and stop‑button reliability degrade the user experience. One reporter (@zpljd258) provided detailed reproduction steps.
- **DingTalk channel needs robustness** — HTTP timeouts and unsupported message types cause silent failures.
- **Positive signal**: The Xiaomi MiMo ASR fix was quickly accepted, showing responsiveness to hardware‑specific needs.

## 8. Backlog Watch
No clearly long‑standing unanswered issues are present (most issues are from the last 48 hours). However, several open PRs may need maintainer attention:

- **PR [#4414](https://github.com/HKUDS/nanobot/pull/4414)** (subagent aggregated result) – open since June 19 with no maintainer comment; still updated today.
- **PR [#4411](https://github.com/HKUDS/nanobot/pull/4411)** (`SuspendTurn`) – open since June 19; no maintainer review.
- **PR [#4482](https://github.com/HKUDS/nanobot/pull/4482)** (thinking style config) – opened June 23, needs acceptance.
- **Security `exec` bypasses** – Issues [#4521](https://github.com/HKUDS/nanobot/issues/4521) through [#4514](https://github.com/HKUDS/nanobot/issues/4514) remain open; no fix PR yet. Urgent maintainer attention required given the exploit potential.

The project maintainers should prioritize the outstanding `exec` allowlist bypasses and review the older enhancement PRs to keep velocity high.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-06-25

## 1. Today's Overview

The Hermes Agent project shows **high activity** today with 50 pull requests updated and 18 issues active. No new releases were cut. The community is heavily engaged in two major directions: **token optimization** (reducing per-call overhead) and **inter‑agent interoperability** (A2A/ACP protocols). A significant number of open PRs address **desktop and gateway UX improvements**, while several bug fixes target session recovery and credential handling. The project maintainers are actively merging contributions (8 PRs closed/merged today), though many long‑standing feature proposals remain open and await decision.

## 2. Releases

No new releases were published on 2026-06-25.

## 3. Project Progress

**8 pull requests** were closed or merged today, signaling steady iterative development. Notable among the top‑20 most‑commented PRs:

- [#43637](https://github.com/NousResearch/hermes-agent/pull/43637) (closed) – Adds a `/sessions` slash‑command handler for messaging gateways (Telegram, Discord, etc.), resolving a gap where the command was registered only for TUI/desktop.
- [#49348](https://github.com/NousResearch/hermes-agent/pull/49348) (closed) – Introduces an optional `hermes-best-practices` skill – an operational checklist covering configuration, delegation habits, memory management, etc.

Additional merged PRs (not individually listed in the top 20 digest) likely include smaller fixes and doc improvements. The open PR queue (42 items) reveals strong ongoing work on **provider integrations** (Vertex AI, Ollama Cloud), **delegation persistence** (SQLite in‑flight tracking), and **gateway robustness**.

## 4. Community Hot Topics

The most active discussions and highest‑voted items reveal clear community priorities:

- **Token overhead reduction** – [#6839 “Lazy Tool Schema Loading”](https://github.com/NousResearch/hermes-agent/issues/6839) (28 comments, 15 👍) and [#4379 “73% of each API call is fixed overhead”](https://github.com/NousResearch/hermes-agent/issues/4379) (16 comments) reflect a broad desire to cut the ~3.5k–5k tokens wasted on tool schemas per call. The community suggests a two‑pass injection strategy and has even built a monitoring dashboard to quantify the issue.

- **Agent‑to‑Agent communication** – [#514 “A2A Protocol Support”](https://github.com/NousResearch/hermes-agent/issues/514) (22 comments, 18 👍) and [#5257 “Generalized ACP client”](https://github.com/NousResearch/hermes-agent/issues/5257) (11 comments, 16 👍) signal strong interest in making Hermes a hub for orchestrating other coding agents (Claude Code, Copilot, etc.) using Google’s A2A standard and the existing ACP protocol.

- **Remote access** – [#10567 “Add --host and CORS config for dashboard”](https://github.com/NousResearch/hermes-agent/issues/10567) (12 comments, 13 👍) addresses the hurdle of accessing the web dashboard over Tailscale/VPN when it currently binds only to localhost.

- **Gateway expansion** – [#3725 “Rocket Chat support”](https://github.com/NousResearch/hermes-agent/issues/3725) (11 comments, 11 👍) is a relatively small request that would widen Hermes’ reach into enterprise chat platforms.

*Underlying needs:* The community is calling for **efficiency‑first tooling** (fewer tokens wasted), **interoperability** (orchestrating other agents), and **better remote‑first design** (dashboard access, thin desktop client).

## 5. Bugs & Stability

Several bugs were reported or fixed today, ranked by severity:

- **P2 – Session‑bricking / data loss**  
  - [#52444](https://github.com/NousResearch/hermes-agent/pull/52444) (fix PR) – Recovers from image‑payload 413 errors by shrinking embedded images, preventing the session from getting stuck.  
  - [#52431](https://github.com/NousResearch/hermes-agent/pull/52431) (fix PR) – Prevents `TypeError` in `_summarize_tool_result` when tool arguments contain non‑string values (booleans, etc.). Without this fix, the context compressor crashes and causes an infinite crash loop in the TUI.  
  - [#52432](https://github.com/NousResearch/hermes-agent/pull/52432) (fix PR) – Gateway runner now merges (instead of overwriting) `agent.request_overrides`, fixing a regression where custom provider extra body settings were lost per turn.  
  - [#52439](https://github.com/NousResearch/hermes-agent/pull/52439) (fix PR) – Resolves env‑var resolution in custom provider credentials and unblocks `delegate_task` to a different provider than the parent.

- **P2 – External provider issues**  
  - [#13834](https://github.com/NousResearch/hermes-agent/issues/13834) (open) – `openai-codex` provider fails on the same machine where official Codex CLI works. Needs reproduction help.  
  - [#50663](https://github.com/NousResearch/hermes-agent/issues/50663) (open) – z.ai rate‑limits Hermes during “peak hours” while other clients work fine.

- **P3 – Browser & CLI bugs**  
  - [#52428](https://github.com/NousResearch/hermes-agent/issues/52428) (open) – `browser_navigate` tool crashes on Windows 10: `name '_hermes_read_browser_output' is not defined`.  
  - [#51741](https://github.com/NousResearch/hermes-agent/pull/51741) (fix PR) – Fixes two bugs in the `/compress` command that prevented it from working in the TUI (UnboundLocalError and race condition).  
  - [#52441](https://github.com/NousResearch/hermes-agent/pull/52441) (fix PR) – Redirects interactive stderr away from the TUI to a log file, preventing UI corruption from raw stderr writes.

**Overall stability:** The project is actively addressing session‑breaking bugs; the rapid submission of fix PRs (especially for 413 and `TypeError`) shows strong maintainer responsiveness.

## 6. Feature Requests & Roadmap Signals

Today’s issue tracker reveals a rich set of user‑requested features. The following are most likely to land in the next minor release:

- **Lazy Tool Schema Loading** (#6839) – High demand, multiple acknowledgements from maintainers, and a clear technical path (two‑pass injection).  
- **Desktop thin‑client installation** (#38602) – 26 👍 and a well‑defined scope; the current desktop app always bootstraps a local runtime, but users want a “connect to remote” mode.  
- **System tray support** (#52434) – Duplicate but with 0 comments; trivial to implement and highly requested in desktop apps.  
- **Conversation readability improvements** (#52426) – Visual distinction between user/assistant messages; low effort, high UX impact.  
- **Session content search** (#52429) – Extend search to include message bodies, not just titles.  
- **Skip desktop‑attachments staging for local gateways** (#52427) – Avoid unnecessary uploads when the gateway shares the filesystem.

*Predictions:* Next minor version (v0.6.1 or v0.7.0) is likely to include lazy tool loading, a desktop tray option, and the `/sessions` gateway command (already merged). The A2A protocol support (#514) may be deferred to a subsequent major release due to its architectural scope.

## 7. User Feedback Summary

Real pain points expressed by users today:

- **Token waste** – “73% of every API call is fixed overhead” (#4379). Users are building custom dashboards to profile consumption and demand schema‑on‑demand.
- **Remote access friction** – The dashboard binds to `127.0.0.1` with hardcoded CORS, forcing users either to work locally or patch the source (#10567).  
- **Desktop UX gaps** – Slash commands like `/model`, `/fast`, `/voice` are blocked or require mouse interaction (#51754). Closing the app exits completely; no system tray (#52434). User/assistant messages look identical (#52426).  
- **Cross‑platform inconsistencies** – The browser tool breaks on Windows (#52428); z.ai rate‑limits only Hermes but not official clients (#50663).  
- **Memory plugin fragility** – The Honcho memory integration had a silent context‑loss bug where `honcho_search` never searched messages – fixed today in PR #52395.

Satisfaction is evident in the high interaction counts: users are investing significant effort to propose, test, and profile features, indicating a dedicated community that expects Hermes to become a production‑ready personal AI assistant.

## 8. Backlog Watch

Several important issues and PRs have been open for weeks without a maintainer decision or reproducer:

- **#6839** – “Lazy Tool Schema Loading” (open since Apr 9, 28 comments, labelled `needs-decision`). No maintainer response on implementation direction.
- **#514** – “A2A Protocol Support” (open since Mar 6, 22 comments, 18 👍). No assignment or milestone.
- **#4379** – “Token overhead analysis: 73%” (open since Apr 1, 16 comments, 0 👍 but rich data) – maintainers have not acknowledged the dashboard or the data.
- **#13834** – “Hermes openai-codex fails” (open since Apr 22, 12 comments, labelled `needs-repro`) – still awaiting a minimal reproduction from the reporter.
- **#21172** – “First‑class Loop Contract” (open since May 7, 4 comments, labelled P3) – a thoughtful design proposal for cron‑backed persistent agent loops; no maintainer reply.
- **#8427** – PR “Vertex AI provider” (open since Apr 12, no maintainer reviews). This is a well‑scoped contribution that would expand Hermes’ enterprise reach.

**Recommendation:** Prioritize a decision on token optimization (#6839) and A2A (#514), as these dominate community attention. Also, provide repro guidance for #13834 to unblock a likely provider bug.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest — 2026-06-25

## 1. Today's Overview

The project saw elevated activity with **14 issues closed** (all stale/security advisories from a single batch) and **13 pull requests updated**, 4 of which were merged or closed. No new releases were published. The maintainers appear to be focusing on **security hardening and code quality**: a wave of vulnerability reports from the same security researcher (YLChen-007) were triaged and closed, while several PRs addressing error handling, type safety, and concurrency were merged. At the same time, **9 PRs remain open**, including a new remote WebSocket mode, a Delta Chat gateway, and several bug fixes that may need further review. Overall project health is stable with a slightly over-loaded open-PR backlog.

## 2. Releases

No new releases were published in the last 24 hours. The latest tag remains **v0.2.9** (as referenced in issue #3012). This section is omitted.

## 3. Project Progress

**Merged/Closed PRs today (2026-06-25):**

| PR # | Title | Status | Area |
|------|-------|--------|------|
| [#3169](https://github.com/sipeed/picoclaw/pull/3169) | fix(evolution): skip cold path for heartbeat turns | **Closed** (merged) | Evolution, token saving |
| [#3166](https://github.com/sipeed/picoclaw/pull/3166) | fix(openai_compat): use structured logger for native_search warning | **Closed** (merged) | Logging, OpenAI-compat |
| [#3168](https://github.com/sipeed/picoclaw/pull/3168) | fix(model): handle error response read failures | **Closed** (merged) | Model list fetch, error handling |
| [#3045](https://github.com/sipeed/picoclaw/pull/3045) | fix(identity): allow_from fallthrough for Matrix user IDs with colon | **Closed** (merged) | Identity parsing, Matrix channel |

**Key fixes advanced:**
- **Evolution module** now skips cold-path scheduling for heartbeat turns, preventing unnecessary token consumption in draft mode.
- **OpenAI-compat providers** fixed a build failure caused by a stray `log.Printf` and improved error reporting when model list responses fail.
- **Matrix identity** bug resolved: users with standard `@user:server` IDs were incorrectly rejected by `allow_from` (fixes [#3044](https://github.com/sipeed/picoclaw/issues/3044)).

**Open PRs that are actively developing new features:**
- [#3118](https://github.com/sipeed/picoclaw/pull/3118) – Remote Pico WebSocket mode for agent
- [#3063](https://github.com/sipeed/picoclaw/pull/3063) – Delta Chat gateway (new channel)
- [#3165](https://github.com/sipeed/picoclaw/pull/3165) – Recover Seed XML tool calls from Volcengine Doubao

## 4. Community Hot Topics

**Most active Issues by comments/reactions:**

1. **[#2404](https://github.com/sipeed/picoclaw/issues/2404) – [Feature] Add streaming HTTP request config**  
   *13 comments, 1 👍* – Originally opened 2026-04-07, closed today. User requested a `"streaming": true` config option for LLM backends (like OpenAI client). The high comment count indicates strong community interest in streaming responses. This feature was likely implemented or resolved in a recent commit.

2. **[#3012](https://github.com/sipeed/picoclaw/issues/3012) – Continuous token consumption with evolution enabled**  
   *5 comments* – Reported token waste every minute when evolution is active. This bug is now closed, and the fix in PR [#3169](https://github.com/sipeed/picoclaw/pull/3169) addresses the root cause (heartbeat turns triggering cold path).

3. **Batch security issues (10 issues)** – All filed by **YLChen-007** on 2026-06-09 and closed today. Topics include:
   - SSRF bypass via ISATAP IPv6 and environment proxies
   - LINE webhook replay attack
   - Feishu reply-context bypass of `allow_from`
   - Cross-Site Request Forgery on launcher first-run setup
   - MQTT `allow_from` spoofing via client_id
   - WeCom group trigger policy bypass
   - Symlink race in approval hooks
   - Command whitelist bypass using `jq`
   Each received 2 comments, indicating the team responded and closed them (likely patched or documented mitigations).

**Underlying needs:**  
The surge in security issues shows that PicoClaw’s multi-channel (Feishu, WeCom, LINE, MQTT, Matrix) architecture is attracting scrutiny. The community wants robust controls against SSRF, message spoofing, and policy bypass. The streaming feature request (#2404) reflects a desire for real-time interaction with LLMs, which is critical for chat-oriented agents.

## 5. Bugs & Stability

**Severity: High**

- **[#3012](https://github.com/sipeed/picoclaw/issues/3012) – Continuous token consumption** (closed)  
  *Impact:* Token waste every minute when evolution draft mode enabled. Fixed by [#3169](https://github.com/sipeed/picoclaw/pull/3169).

**Severity: Medium**

- **[#3167](https://github.com/sipeed/picoclaw/issues/3167) – PageAgent compatibility with Vue/Element UI**  
  *Impact:* Users testing PageAgent on Vue 2 apps report that `v-model` and component state are not easily manipulated. No fix yet; primarily a feature request/discussion.

**New fix PRs submitted today for latent bugs:**
- [#3172](https://github.com/sipeed/picoclaw/pull/3172) – Explicitly ignore `Close()` errors in error paths/retry loops (8 call sites)
- [#3171](https://github.com/sipeed/picoclaw/pull/3171) – Add `ok` checks for `sync.Map` type assertions in LINE Send to prevent panics
- [#3170](https://github.com/sipeed/picoclaw/pull/3170) – Close base64 encoder on `io.Copy` error path to prevent resource leaks

These are all defensive fixes, not triggered by user reports, indicating proactive code hygiene.

**Stale open bugs (no maintainer response yet):**  
- [#3079](https://github.com/sipeed/picoclaw/issues/3079) (jq environment disclosure) – closed, but likely still awaiting a code patch? The issue is closed but no linked PR found.  
- [#3078](https://github.com/sipeed/picoclaw/issues/3078) (SSRF via proxy) – similar pattern.

## 6. Feature Requests & Roadmap Signals

**Notable requests with clear demand:**
- **[#2404](https://github.com/sipeed/picoclaw/issues/2404) – Streaming HTTP request config** (closed) – Likely already implemented or planned for next release.
- **[#3118](https://github.com/sipeed/picoclaw/pull/3118) – Remote Pico WebSocket mode** – Open PR from @jp39, adds `--remote` flag to agent. This could land in v0.3.0.
- **[#3063](https://github.com/sipeed/picoclaw/pull/3063) – Delta Chat gateway** – Open PR from @trufae, new channel connector. Still open since 2026-06-08, may need review.
- **[#3165](https://github.com/sipeed/picoclaw/pull/3165) – Seed XML tool calls** – Adds support for Volcengine Doubao’s `<seed:tool_call>` blocks. Important for Chinese LLM ecosystem.

**Roadmap signals:**  
The closure of 10 security advisories suggests the team is prioritizing security hardening for the next release. The streaming feature (#2404) and evolution token fix (#3012) are low-hanging fruit for a minor release. The **PageAgent Vue question** (#3167) hints that users want better framework-agnostic DOM traversal — this could become a future enhancement.

## 7. User Feedback Summary

**Pain points:**
- **Token waste with evolution** (#3012) – user @xpader reported continuous every-minute consumption. The fix was merged same day, indicating responsive support.
- **Vue/Element UI compatibility** (#3167) – @Wavekip notes that PageAgent struggles with reactive frameworks; suggests a need for `v-model` detection or component snapshot APIs. This is a usability gap for enterprise users.
- **Matrix user identity breakage** (#3044) – users could not use standard Matrix IDs in `allow_from`. Fixed in PR #3045.

**Satisfaction (implied):**
- The **security batch** was closed without further escalation – likely users are satisfied with the response.
- The **streaming request** (#2404) received positive reaction and was closed, suggesting the feature was implemented or deemed unnecessary.

**Use cases highlighted:**
- Real-time LLM interaction (streaming)
- Multi-channel enterprise deployment (Feishu, WeCom, LINE, MQTT, Matrix, Delta Chat)
- Automated code review (evolution draft mode)
- Testing of internal business systems (PageAgent on Vue)

## 8. Backlog Watch

**Long-unanswered open PRs needing maintainer attention:**

| PR # | Title | Created | Last Updated | Stale? | Notes |
|------|-------|---------|--------------|--------|-------|
| [#3063](https://github.com/sipeed/picoclaw/pull/3063) | feat: add deltachat gateway | 2026-06-08 | 2026-06-25 | **Stale** (17 days) | Major new feature, no reviewer comments. Risk of merge conflict. |
| [#3116](https://github.com/sipeed/picoclaw/pull/3116) | fix(pico): complete turn.done lifecycle signaling | 2026-06-12 | 2026-06-24 | **Stale** (13 days) | Addresses issue #2984 – preserving request_id for queued messages. |
| [#3128](https://github.com/sipeed/picoclaw/pull/3128) | fix(web): ignore resp.Body.Close() errors after io.ReadAll | 2026-06-15 | 2026-06-25 | **Stale** (10 days) | Low-risk, trivial fix for code quality. |
| [#3115](https://github.com/sipeed/picoclaw/pull/3115) | Fix inline data URL media extraction | 2026-06-12 | 2026-06-25 | **Stale** (13 days) | Important bug fix for session history corruption. |

**Open issues without maintainer response (none – all issues are closed).**  
However, the **PageAgent question** (#3167) is a support request with 0 comments from maintainers – while not a bug, it remains unanswered.

**Recommendation:**  
The maintainers should prioritize reviewing the Delta Chat gateway and the `turn.done` lifecycle fix, as they represent significant new capabilities and bug fixes that have been waiting for over two weeks. The inline data URL fix (#3115) also directly impacts user experience.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-25

## 1. Today's Overview
NanoClaw showed high activity in the last 24 hours, with 19 pull requests updated and one open issue. Four PRs were merged or closed, including two security fixes and a new authentication posture. The project remains in active development, with contributions across multiple areas: Telegram multi-bot support, Matrix E2EE, remote MCP servers, security hardening, and container/infrastructure fixes. No new releases were published.

## 2. Releases
No new releases in the last 24 hours. The latest release remains unchanged.

## 3. Project Progress — Merged/Closed PRs Today
Four pull requests were merged or closed within the last 24 hours:

- **#2830 [closed]** `fix(setup): reap dead peer service registrations whose binary is gone` — Addresses launchd/systemd ghost entries left after removing a NanoClaw checkout. (by amit-shafnir) [PR #2830](https://github.com/nanocoai/nanoclaw/pull/2830)
- **#2855 [closed]** `feat(auth): subscription-primary credential posture with API-key failover` — Makes OAuth primary and falls back to `ANTHROPIC_API_KEY` on subscription failure, with operator alerts. (by bogdano2) [PR #2855](https://github.com/nanocoai/nanoclaw/pull/2855)
- **#2849 [closed]** `feat(telegram): support multiple bot instances via TELEGRAM_BOT_TOKEN_<SUFFIX>` — Enables running multiple Telegram bots from a single instance. (by grantland) [PR #2849](https://github.com/nanocoai/nanoclaw/pull/2849)
- **#2799 [closed]** `fix(security): confine send_file reads to /workspace (CVE-2026-29611)` — Restricts file reads to `/workspace` to prevent path traversal and credential exfiltration. (by sturdy4days) [PR #2799](https://github.com/nanocoai/nanoclaw/pull/2799)

These merges indicate progress in credential management, Telegram multi-bot, and security hardening.

## 4. Community Hot Topics
The only open issue updated in the last 24 hours is **#2852 [OPEN]** about Telegram multi-bot support:  
> "we had it, and then it got removed... Is it ever going to be implemented?"  
This issue has no comments or reactions yet, but it reflects a recurring user demand. Notably, PR #2849 (closed) implemented the feature, and a follow-up PR **#2853 [OPEN]** (`feat(telegram): support multiple bot instances`) was also opened, suggesting active work to resolve the concern.  
[Issue #2852](https://github.com/nanocoai/nanoclaw/issues/2852) | [PR #2853](https://github.com/nanocoai/nanoclaw/pull/2853)

Another notable PR is **#2844 [OPEN]** (`feat(matrix): native persistent E2EE adapter via matrix-bot-sdk`), which replaces the Beeper Chat SDK bridge with native Matrix crypto using Rust bindings. This is a significant architectural change for Matrix integration.  
[PR #2844](https://github.com/nanocoai/nanoclaw/pull/2844)

## 5. Bugs & Stability
Multiple bug fixes and security improvements are in the pipeline. Ranked by severity:

**High severity (security):**
- **CVE-2026-29611** — `send_file` path traversal (fixed in #2799, already closed). Allows prompt-injected agents to read any container file.  
- **#2800 [OPEN]** — `folder` validation bypass in `ncl groups create/update` (CWE-22) and missing image tag pinning.  
- **#2802 [OPEN]** — Socket transport has no timeout and no buffer limit, enabling indefinite hangs and memory growth.  
- **#2801 [OPEN]** — `safeParseContent` returns non-object JSON primitives, causing undefined fields and potential routing misbehavior.

**Medium severity:**
- **#2854 [OPEN]** — `fix(onecli): redirect TMPDIR so gateway CA mounts into containers on macOS` — Rancher Desktop users face self-signed certificate errors.  
- **#2850 [OPEN]** — Signal group messages lack `isMention` and `isGroup` fields, preventing proper routing.  
- **#2851 [OPEN]** — Abandoned poll loops in test helpers steal messages from later tests.  
- **#2848 [OPEN]** — OpenCode provider fails to find `.env` and provider/model env vars without explicit cwd.

**Lower severity / infrastructure:**
- **#2750 [OPEN]** — Stale `outbound.db` journals after container kills; hot-journal poll races.  
- **#2830** (closed today) — Dead peer service registrations accumulating on macOS/Linux.  
- **#2845 [OPEN]** — `q.ts` not forwarding positional parameters for SQL queries.  
- **#2815 [OPEN]** — Router `safeParseContent` fails on primitive JSON (regression test added).

## 6. Feature Requests & Roadmap Signals
The open PRs and issue reveal several directions for upcoming releases:

- **Telegram multi-bot** — Requested in #2852, implemented in #2849 (closed), but #2853 (open) may refine or extend it. Likely to land in next release.
- **Native Matrix E2EE** (#2844) — A major uplift for Matrix users, replacing the WASM-based bridge. Could be included after review.
- **Remote MCP servers** (#2847) — Support for HTTP/SSE-based MCP servers in addition to local stdio, expanding agent tool connectivity.
- **Extension-point seams** (#2842) — Generic `registerX/applyX` hooks for host and container runtime, enabling third-party extensions without forking.
- **Docker-in-Docker agent groups** (#2846) — Mounting Docker socket with correct permissions, enabling agents to spawn Docker containers.
- **Subscription-primary auth** (#2855, closed) — Already merged; signals ongoing investment in robust authentication fallback.

Predictions: The next version will likely include Telegram multi-bot, remote MCP servers, and several security fixes (especially #2800, #2802, #2801). The Matrix E2EE rewrite may take longer due to its scope.

## 7. User Feedback Summary
The only user-reported pain point in the last 24 hours is **Issue #2852**: a user who previously had Telegram multi-bot functionality is frustrated it was removed, and reports that "Claude cannot get it to work" with the promised "instance" support. This indicates a gap between documentation and actual usability, and a desire for reliable multi-bot support. No other direct user feedback (comments, reactions) is present in the data. The high volume of security fixes suggests users or testers have encountered reliability and safety issues in practice.

## 8. Backlog Watch
Several important PRs remain open and may require maintainer attention:

- **#2750 [OPEN]** (created June 12) — Stale `outbound.db` journals after container kills. Involves complex race conditions and two linked issues (#2516, #2640). No comments from maintainers.  
- **#2800, #2801, #2802 [OPEN]** (created June 17) — Three security fixes by sturdy4days. All have no merges or reviews yet. These address CWE-22, unsafe JSON parsing, and socket vulnerabilities — should be prioritized.  
- **#2815 [OPEN]** (created June 18) — Router regression test for primitive JSON; may be waiting for #2801 resolution.  
- **#2842 [OPEN]** (created June 23) — Extension-point seams; no comments.  
- **#2844 [OPEN]** (created June 24) — Matrix E2EE rewrite; no comments yet but significant code change.

None of these have been explicitly acknowledged by maintainers in the data. Given the security implications, #2800–#2802 and #2750 should be reviewed as soon as possible.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-06-25

## 1. Today's Overview
The project saw elevated activity with 22 issues and 39 pull requests updated in the last 24 hours, indicating a sustained push to address Reborn stability and feature completeness. No new releases were published. The team closed 13 PRs (merged or closed) and resolved 4 issues, while 18 issues and 26 PRs remain open. The overall health is positive, with the core contributor pipeline delivering several critical fixes and architectural improvements, though a handful of high-impact bugs (duplicate approval flows, cancellation reliability, and timeout edge cases) remain under active investigation.

## 2. Releases
**No new releases.** The latest release remains the previous tag; no version bumps or changelogs appeared today.

## 3. Project Progress — Merged/Closed PRs Today
The following pull requests were merged or closed in the last 24 hours, representing concrete progress:

- **#5145 (closed) — refactor(reborn): clean up capability activity lifecycle**  
  Core refactor that stabilises activity identity (`CapabilityActivityId`) across the gate lifecycle, fixing three related issues where activity rows, gate prompts, and WebUI rendering used inconsistent identity. The closed status indicates the fix has been integrated.  
  [nearai/ironclaw PR #5145](https://github.com/nearai/ironclaw/pull/5145)

- **#5202 (closed) — fix recurring trigger poller hang**  
  Makes trigger post-submit delivery hooks asynchronous to prevent the poller from blocking, and adds regressions for slow Slack delivery. **Fixes #5148** (turn scheduler heartbeat self-deadlock).  
  [nearai/ironclaw PR #5202](https://github.com/nearai/ironclaw/pull/5202)

- **#5225 (closed) — fix(reborn): release Slack admission permit once inbound is durably accepted**  
  Resolves a semaphore leak that could cause Slack webhook intake to stall. Early ACK path now releases the admission permit after durable acceptance, preventing concurrency exhaustion under load.  
  [nearai/ironclaw PR #5225](https://github.com/nearai/ironclaw/pull/5225)

- **#5028 (closed) — Follow up: make denied activity ids explicit and stable**  
  Follow-up from #4978; stabilises activity identity for denied tool calls (merged earlier).  
  [nearai/ironclaw Issue #5028](https://github.com/nearai/ironclaw/issue/5028)

- **#4598 (closed) — Local service lifecycle install, start, stop, and status support**  
  Delivers Reborn local service lifecycle management for supported OS targets (parent ticket #4533).  
  [nearai/ironclaw Issue #4598](https://github.com/nearai/ironclaw/issue/4598)

Additionally, several other closed PRs (e.g., dependabot updates, documentation changes) were merged.

## 4. Community Hot Topics
The following issues and pull requests generated the most discussion or are widely referenced:

- **#5196 — [Reborn] "Ask each time" tool permission may fail with authorization error and trigger duplicate approval flow**  
  A user-facing bug that creates a poor experience: after approving a tool, it fails with an `authorization` error and the assistant re-prompts. One comment so far, but the impact is high as it breaks the core approval trust model.  
  [nearai/ironclaw Issue #5196](https://github.com/nearai/ironclaw/issue/5196)

- **#5148 — Turn scheduler heartbeat can self-deadlock while a run holds transition state**  
  A severe deadlock scenario that can freeze a running turn indefinitely. Already fixed by PR #5202 (closed today), but generated significant internal investigation due to production impact.  
  [nearai/ironclaw Issue #5148](https://github.com/nearai/ironclaw/issue/5148)

- **#5232 — [codex] durable runner lease sidecar**  
  A large (XL) PR that moves heartbeat persistence into per-run durable lease sidecars to reduce filesystem contention. Despite being a core-internal change, its scope suggests it is a response to production incidents.  
  [nearai/ironclaw PR #5232](https://github.com/nearai/ironclaw/pull/5232)

- **#5156 — feat(skill-learning): any-backend distillation, approval gate, learned-only scoping, persisted switches**  
  A major skill-learning feature that introduces an approval gate for freshly learned skills and persists user preferences. This PR has been open for 2 days and touches on user-controlled AI learning, which is a high-interest area.  
  [nearai/ironclaw PR #5156](https://github.com/nearai/ironclaw/pull/5156)

## 5. Bugs & Stability
Several bugs were reported or tracked today. Ranked by severity:

| Severity | Issue | Description | Fix PR exists? |
|----------|-------|-------------|----------------|
| **Critical** | #5196 | “Ask each time” approval triggers duplicate auth flow; tool fails with authorization error | No (open) |
| **Critical** | #5209 | Canceling a response does not reliably stop execution before sending a new message | No (open) |
| **High** | #5231 | Automation setup fails with `HostUnavailable` at planned driver prompt stage | No (open) |
| **High** | #5227 | Run failure messages attach to wrong conversation turn | No (open) |
| **High** | #5189 | Successful tool runs do not show activity details while running | No (open) |
| **High** | #5148 | Scheduler heartbeat self-deadlock | Yes – #5202 (closed) |
| **High** | #5213 | Turn-model timeout stack can't serve large-context (1M) / large-output generations | No (open) |
| **Medium** | #5229 | Durable capability display previews use runtime owner scope in WebUI runs | Yes – #5230 (open) |
| **Medium** | #5120 | Unify gate declined semantics across auth, approval, and activity projection | Yes – #5145 (closed) |
| **Low** | #5144 | Show NEAR AI default base URL in provider card | Yes – #5217 (open) |
| **Ongoing** | #4108 | Nightly E2E failed (recurring since 2026-05-27) | No (open) |

**Notable fixes merged today:** #5202 (heartbeat deadlock), #5225 (Slack admission permit leak), #5145 (activity lifecycle).

## 6. Feature Requests & Roadmap Signals
The following issues and PRs indicate likely upcoming features or user-demanded improvements:

- **#5212 — Show message timestamps consistently in the conversation**  
  This enhancement request to display timestamps for both user and assistant messages is addressed by PR #5226 (open). Expect it in the next WebUI update.  
  [nearai/ironclaw Issue #5212](https://github.com/nearai/ironclaw/issue/5212)

- **#5149 — Context management: progressive tool disclosure (flag-gated, default off)**  
  A large PR that reduces per-call prompt size by disclosing only relevant tool schemas. This could significantly lower latency and timeout rates. Currently gated behind a feature flag.  
  [nearai/ironclaw PR #5149](https://github.com/nearai/ironclaw/pull/5149)

- **#5068 — Tool permissions + global auto-approve settings surface**  
  End-to-end wiring for approval settings in the Reborn WebUI. This PR has been open since 2026-06-18 and is a key part of the permission system.  
  [nearai/ironclaw PR #5068](https://github.com/nearai/ironclaw/pull/5068)

- **#5156 — Skill-learning with approval gate**  
  Freshly learned skills will be saved as inactive and require human approval before activation. This is a privacy/control feature that may land in the next minor release.  
  [nearai/ironclaw PR #5156](https://github.com/nearai/ironclaw/pull/5156)

## 7. User Feedback Summary
User-reported pain points this week centre on **approval/reliability** and **UI consistency**:

- **Approval/persistence:** Users are frustrated when “Ask each time” settings fail (issue #5196) and when tool permissions are not consistently reflected (PR #5195 addresses this).
- **Conversation timeline:** Failures attaching to the wrong turn ( #5227 ) and missing timestamps ( #5212 ) degrade trust in the chat UI.
- **Cancellation:** The inability to reliably cancel and start a new message ( #5209 ) is disruptive in fast-paced workflows.
- **Activity visibility:** Tool runs that succeed show no details during execution ( #5189 ), while failures do – causing confusion.
- **Automation reliability:** Automated workflows failing due to host unavailability ( #5231 ) or recurring nightly test failures ( #4108 ) undermine developer confidence.

Positive signals: the team’s rapid response to production deadlocks (heartbeat and Slack admission) and the transparent backlog updates ( #5220, #5221 ) are appreciated by community observers.

## 8. Backlog Watch
The following open issues and PRs have been unanswered or idle for an extended period, requiring maintainer attention:

- **#4108 — Nightly E2E failed**  
  First reported 2026-05-27, updated today but still open. The nightly E2E suite has been failing for nearly a month with no resolution or root-cause comment. Urgent attention needed.  
  [nearai/ironclaw Issue #4108](https://github.com/nearai/ironclaw/issue/4108)

- **#4598 — Local service lifecycle install, start, stop, and status support**  
  Despite being closed today, this issue sat open for 16 days (since 2026-06-09). Its resolution is good news, but similar lifecycle issues (e.g., OS-specific unsupported targets) may linger.

- **#5144 — Show NEAR AI default base URL in provider card**  
  Opened 2026-06-23, a simple UI fix that has a corresponding PR #5217 but remains open after 2 days. Low risk but would improve user clarity.

- **#5120 — Unify gate declined semantics across auth, approval, and activity projection**  
  Open since 2026-06-22, closed today via PR #5145. This was actively worked and resolved, but the persistence of terminology issues in earlier releases suggests a need for forward-looking documentation.

*No other issues remain unanswered beyond 5 days without updates, indicating healthy maintainer responsiveness.*

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-06-25

## 1. Today's Overview
The project saw high development activity with **10 pull requests merged** in the last 24 hours, all closed without open PRs. A single open issue (#1392) remains unresolved, flagged as stale but still updated today. The merged PRs focus heavily on **stability enhancements** for the cowork (collaborative) mode, OpenClaw extension framework, and settings synchronization. No new releases were cut, but the rapid pace of fixes suggests a strong push toward a more robust release candidate. Overall project health appears solid, with active maintainer engagement and a clear emphasis on bug fixing and polish.

## 2. Releases
*No new releases were published today.*

## 3. Project Progress — Merged PRs Today
Ten PRs were merged, advancing several areas:

- **Settings & Auto‑Launch** (#2206): Fixed state sync for launch‑at‑login, including Windows registry cleanup for historical argument variants.
- **Cowork Plan Mode** (#2205, #2204, #2200, #2199, #2197): Improved plan icon styling, fixed block‑level plan tag parsing, prevented duplicate plan messages caused by stream jitter, ensured subagent polling continues after parent completion (with a 5‑minute timeout), and deduplicated final assistant prefixes after history fallback.
- **OpenClaw Framework** (#2203, #2202, #2201, #2198): Precompiled local extension entries, kept browser plugin allowlisted, fixed deduplication of assistant final sync after `sessions_yield`, and preinstalled QQ and Discord channel plugins with correct account indexing.
- **Infrastructure** (#2203, #2202, #2198): Tightened packaging checks, added runtime sync assertions for plugin configuration, and fixed NIM account/environment‑variable indexing.

All PRs are linked below:
- [#2206](https://github.com/netease-youdao/LobsterAI/pull/2206)
- [#2205](https://github.com/netease-youdao/LobsterAI/pull/2205)
- [#2204](https://github.com/netease-youdao/LobsterAI/pull/2204)
- [#2203](https://github.com/netease-youdao/LobsterAI/pull/2203)
- [#2202](https://github.com/netease-youdao/LobsterAI/pull/2202)
- [#2201](https://github.com/netease-youdao/LobsterAI/pull/2201)
- [#2200](https://github.com/netease-youdao/LobsterAI/pull/2200)
- [#2199](https://github.com/netease-youdao/LobsterAI/pull/2199)
- [#2198](https://github.com/netease-youdao/LobsterAI/pull/2198)
- [#2197](https://github.com/netease-youdao/LobsterAI/pull/2197)

## 4. Community Hot Topics
The only open issue (#1392) is the most commented and reacted item. It reports a **scheduled task toggle** that fails to respond for some tasks. Despite being created in April, it was updated today, possibly due to recent investigation or user follow‑up. The issue has 1 comment and 0 👍, but its persistence and staleness signal a potential pain point. No PR directly addresses this issue yet.

- [#1392](https://github.com/netease-youdao/LobsterAI/issues/1392) – *[stale]* Timer switch unresponsive for some scheduled tasks.

All merged PRs had zero comments, indicating a streamlined review process with no extended community discussion.

## 5. Bugs & Stability
**High‑severity bug (open):**  
- **#1392** – Scheduled task toggle unresponsive for a subset of tasks. No fix PR exists; bug is 2.5 months old and marked stale.

**Stability fixes merged today (indirectly addressing potential crashes/data issues):**  
- **#2206** – Settings sync not reflecting OS state (e.g., auto‑launch changes lost on Windows).
- **#2204** – Plan mode leaking tags into chat messages instead of rendering plan cards.
- **#2201** – Duplicate assistant replies after subagent yield.
- **#2200** – Duplicate plan messages due to stream snapshot jitter.
- **#2199** – Polling stopped prematurely for subagent runs, potentially missing terminal updates.
- **#2197** – Redundant assistant prefix after history fallback in OpenClaw runs.
- **#2203/#2202** – Packaging/build issues causing extensions or browser plugin to be missing in precompiled builds.
- **#2198** – NIM account/env‑var indexing failures prevented QQ/Discord plugin initialization.

These fixes improve data consistency, stream reliability, and extension loading. The overall stability posture has improved substantially today.

## 6. Feature Requests & Roadmap Signals
No new feature requests were filed today. However, the merged PRs indicate active roadmap focus on:

- **Improved collaborative AI (cowork) interactions**: Real plan rendering, subagent lifecycle, and deduplication point to a more robust multi‑agent experience.
- **OpenClaw plugin ecosystem**: Pre‑installing QQ and Discord plugins, browser plugin allowlisting, and local extension compilation suggest the team is preparing for broader plugin availability and possibly a public marketplace.
- **Settings reliability**: The auto‑launch fix (#2206) hints at better cross‑platform settings synchronization, likely a precursor to more advanced system integration.

Predicting next‑version priorities: stabilization of the cowork plan mode, expansion of OpenClaw plugin support (more channels), and resolution of the long‑standing timer toggle bug.

## 7. User Feedback Summary
The only direct user feedback today comes from issue #1392: a user unable to disable some scheduled tasks via the toggle switch. The lack of maintainer response for three months may indicate either low priority or complexity. The user provided a screenshot, but no additional context was given. This single data point suggests dissatisfaction with the reliability of the task scheduler UI. On the positive side, the merged fixes today address several user‑facing problems introduced in previous versions (e.g., duplicate messages, missing plan cards, broken extensions), which should improve overall satisfaction.

## 8. Backlog Watch
**Critical item needing maintainer attention:**  
- **#1392** (open since 2026‑04‑03, updated 2026‑06‑25) – Timer switch unresponsive. Despite being marked stale, it was updated today and remains unresolved. No assignee or milestone. This bug directly impacts a core feature (scheduled tasks) and should be prioritized.

No other open issues or long‑standing PRs were reported. The backlog appears clean apart from this single issue.

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

# TinyClaw Project Digest – 2026-06-25

*Generated from GitHub data for [TinyAGI/tinyagi](https://github.com/TinyAGI/tinyagi)*

---

## 1. Today's Overview

The TinyClaw project saw minimal activity in the past 24 hours. No new issues were created or updated, and no releases were published. A single pull request (PR #281) was closed after merging, addressing Windows-specific CLI bugs that had prevented native (non-WSL) usage. The project remains stable, with no open issues or pending PRs to track. The quiet period suggests the team may be consolidating the recent Windows fix or preparing for the next development cycle.

---

## 2. Releases

*None.* No new versions were published in the last 24 hours.

---

## 3. Project Progress

- **PR #281 (merged/closed):** *fix: Windows cross-platform support in CLI*  
  Author: `mperkins0155` | [View PR](https://github.com/TinyAGI/tinyagi/pull/281)  
  **What was fixed:** Three Windows-only bugs that prevented the `tinyagi` CLI from running natively (without WSL):
  1. **Doubled drive letter bug** – `new URL('.', import.meta.url).pathname` returned `/C:/Users/...` on Windows, causing `path.resolve` to treat the leading `/` and colon incorrectly, leading to `MODULE_NOT_FOUND`.
  2. (Two additional bugs, presumably related to path handling and process spawning on Windows, as described in the PR summary.)
  
  This fix improves cross-platform compatibility, making TinyClaw accessible to Windows users without requiring WSL.

---

## 4. Community Hot Topics

No issues or PRs received comments or reactions in the last 24 hours. The only merged PR (#281) had no discussion activity; the fix appears straightforward and uncontroversial.

---

## 5. Bugs & Stability

**No new bugs reported today.** The recently merged PR #281 addressed three existing Windows-only bugs that affected CLI startup. These were blocking issues for native Windows users, now resolved. No regressions or stability concerns have emerged from this fix.

---

## 6. Feature Requests & Roadmap Signals

No feature requests were submitted or discussed in the last 24 hours. The project appears to be in a maintenance phase, with the most recent activity focused on platform compatibility rather than new functionality.

---

## 7. User Feedback Summary

No user feedback or pain points were expressed in issues or PR comments today. The Windows fix suggests that native Windows support was a known, unaddressed need—now met. Without additional input, the project’s health remains neutral; users appear satisfied or quiet.

---

## 8. Backlog Watch

**No items require maintainer attention.** All issues are resolved, all PRs are merged, and there are no unanswered or stale items in the tracker. The project backlog is currently clean.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest – 2026-06-25

## 1. Today's Overview

CoPaw experienced a very high activity day, with 50 pull requests updated (28 merged/closed, 22 open) and 14 issues updated (10 open, 4 closed). No new releases were published. The community contributed several first-time patches addressing critical bugs, especially around browser automation resource leaks, media capability detection, and Linux desktop compatibility. The project is in a rapid development phase, with significant progress on the Runtime 2.0 migration, Slack channel integration, and desktop auto‑update features.

## 2. Releases

None. No new versions were released today. The latest available version remains v1.1.12.post2 (as referenced in recent bug reports).

## 3. Project Progress

**Merged/Closed PRs Today (28 total)**

Key merged changes:

- **#5532** – Restored SkillPool styles broken by PR #5521 (CSS grid layout, independent styling).
- **#5522** – Added source-type filter (Cron, Heartbeat, Memory) to the Inbox push messages tab, improving message triage.
- **#5059** – Fixed Matrix encrypted media download by using `nio.client.download` instead of raw HTTP, resolving “Encrypted media unavailable” errors.
- **#5485** – Refined MCP tool name parsing for OpenAI API compatibility.
- **#5486** – Fixed JSON decoding of tool inputs.

These merges indicate continued focus on UI polish, communication channel reliability, and tool integration robustness. The closed issues #5345 (custom OpenAI function calling) and #2733 (Chrome process cleanup) were also resolved – the latter via the earlier PR #2843, but a regression (#5520) has since been reported.

## 4. Community Hot Topics

Most commented issues today:

| Issue | Comments | Status | Topic |
|-------|----------|--------|-------|
| [#5345](https://github.com/agentscope-ai/QwenPaw/issues/5345) | 8 | Closed | Custom OpenAI providers (OMLX) not supporting function calling |
| [#2733](https://github.com/agentscope-ai/QwenPaw/issues/2733) | 6 | Closed | Chrome processes not closing after browser automation (resource exhaustion) |
| [#5162](https://github.com/agentscope-ai/QwenPaw/issues/5162) | 5 | Open | Conversation thinking logic enters infinite loop |
| [#5480](https://github.com/agentscope-ai/QwenPaw/issues/5480) | 4 | Open | Console long-message layout corruption (CSS recalculation missing) |
| [#5505](https://github.com/agentscope-ai/QwenPaw/issues/5505) | 3 | Open | MiniMax-M3 image safety errors cached as `rejects_media=True` |

The underlying needs are clear: users demand reliable model compatibility (custom providers and media handling), stable resource management (Chrome processes), and a polished web UI (layout rendering, large session support). The high activity on the MiniMax caching bug (#5505) and the immediate PR (#5533) from the same reporter shows strong community engagement.

## 5. Bugs & Stability

**High Severity** (potential crashes, data loss, or memory leaks):

- **[#5520](https://github.com/agentscope-ai/QwenPaw/issues/5520) – Browser tool `stop()` leaves renderer processes (memory leak, regression)**  
  Reported today: `stop()` does not fully clean Chrome renderer subprocesses (~150–210MB each), accumulating over repeated start/stop cycles. This is a regression from the fix in #2733 / PR #2843. No fix PR yet.

- **[#5162](https://github.com/agentscope-ai/QwenPaw/issues/5162) – Conversation thinking loop**  
  Open for 13 days, 5 comments. The agent enters an infinite thinking loop, preventing normal conversation. Root cause not yet identified.

- **[#5505](https://github.com/agentscope-ai/QwenPaw/issues/5505) – MiniMax-M3 image safety errors cached as media rejection**  
  A false positive `rejects_media=True` is persisted, stripping images from later requests. A fix PR [#5533](https://github.com/agentscope-ai/QwenPaw/pull/5533) was opened by the same reporter.

**Medium Severity** (partial functionality loss or regression):

- **[#5528](https://github.com/agentscope-ai/QwenPaw/issues/5528) – Browser tool fails on Linux when default browser `Exec` line wraps `env`**  
  The default browser detection picks only the first token, causing `env` to be used as the binary. Fix PR [#5526](https://github.com/agentscope-ai/QwenPaw/pull/5526) submitted.

- **[#5480](https://github.com/agentscope-ai/QwenPaw/issues/5480) – Console long messages layout corrupt, fixed by tab switching**  
  CSS layout recalculation is missing. No fix PR yet.

- **[#5479](https://github.com/agentscope-ai/QwenPaw/issues/5479) – Web UI crashes when opening conversations >500KB**  
  Frontend renders an error page; the session must be deleted. No fix PR.

**Low Severity / Environment-Specific**:

- Panel #5508 (file preview 404 on Windows app) – closed as a question with workaround provided.
- Panel #5512 (provider count mismatch on stats panel) – open, low impact.

## 6. Feature Requests & Roadmap Signals

Several community requests point toward upcoming releases:

- **Dynamic Model Switching** ([#5527](https://github.com/agentscope-ai/QwenPaw/issues/5527)) – User asks for automatic fallback when a model is rate-limited or unavailable, avoiding task interruption. This aligns with the ongoing Runtime 2.0 migration and the custom model ordering PR [#5399](https://github.com/agentscope-ai/QwenPaw/pull/5399).

- **Command Auto‑complete Conflict** ([#5529](https://github.com/agentscope-ai/QwenPaw/issues/5529)) – `/new` command conflicts with skills starting with `/news`. A UI disambiguation improvement is needed.

- **Spawn Subagent in Runtime 2.0** ([#5523](https://github.com/agentscope-ai/QwenPaw/issues/5523)) – The `spawn_subagent` tool is missing from the new runtime registry, blocking migration. PR [#5524](https://github.com/agentscope-ai/QwenPaw/pull/5524) opened same day to register it.

- **New Channel & Desktop Features** – PR [#5193](https://github.com/agentscope-ai/QwenPaw/pull/5193) (Slack channel with multimodal/streaming), [#4669](https://github.com/agentscope-ai/QwenPaw/pull/4669) (Tauri auto‑updater), and [#4041](https://github.com/agentscope-ai/QwenPaw/pull/4041) (Tauri tray) are long‑standing feature PRs awaiting merging. They signal strong demand for enterprise‑grade messaging and desktop UX.

## 7. User Feedback Summary

Real user pain points expressed today:

- **“I can’t use my preferred model provider”** – Issue #5345 (OMLX) was closed but highlights a broader wish for custom‑provider parity with Ollama/OpenAI.
- **“My browser automation drains RAM”** – Issue #5520 shows users rely heavily on `browser_use` and are annoyed by incomplete resource cleanup.
- **“The web UI breaks on long sessions”** – Multiple reports (#5480, #5479) about layout corruption and crashes, suggesting the frontend is a key friction point.
- **“Windows file preview doesn’t work”** – Issue #5508 was closed as a question, but the underlying path handling may still need attention.
- **“Model switches should be seamless”** – Issue #5527 reflects a production‑grade expectation for model fallback.

Overall, users are actively testing edge cases and contributing fixes, indicating a strong, technically engaged community. Satisfaction is mixed: the software is powerful and rapidly improving, but regressions (like the Chrome leak) and UI instabilities are causing frustration.

## 8. Backlog Watch

Issues and PRs that have been open for an extended time without maintainer response or action:

| Item | Days Open | Type | Notes |
|------|-----------|------|-------|
| [#5162](https://github.com/agentscope-ai/QwenPaw/issues/5162) | 13 | Bug | Infinite loop in conversation thinking – no fix PR yet, high impact. |
| [#5479](https://github.com/agentscope-ai/QwenPaw/issues/5479) | 1 | Bug | Large session crash – no PR, but highly disruptive to users. |
| [#5480](https://github.com/agentscope-ai/QwenPaw/issues/5480) | 1 | Bug | Console layout corruption – no PR. |
| [#5193](https://github.com/agentscope-ai/QwenPaw/pull/5193) | 10 | Feature | Slack channel integration – no recent updates; may need review. |
| [#4669](https://github.com/agentscope-ai/QwenPaw/pull/4669) | 31 | Feature | Tauri auto‑updater – long open, likely awaiting final QA. |
| [#4041](https://github.com/agentscope-ai/QwenPaw/pull/4041) | 51 | Feature | Tauri tray behavior – very old, may be superseded or need rebase. |

Maintainer attention is especially needed on the memory‑leak regression (#5520) and the infinite‑loop bug (#5162), both affecting core usability. The feature PRs for Slack and desktop updates would significantly expand the project’s ecosystem.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-06-25

## Today’s Overview

ZeroClaw saw intense development activity today, with **50 pull requests updated** (12 merged/closed) and **10 issues updated** (3 closed). The project remains in a high-velocity pre-release phase, with three v0.8.3 tracker issues (#8071, #8070, #8072) and one v0.8.2 release-support tracker (#8181) actively coordinated. No new release was published. The highlight of the day is the closure of a **S0‑severity security bug** (#8279) where the delegate tool could bypass the parent’s tool allowlist – fixed and merged. Concurrently, a critical MCP stdio process leak (#5903) remains open and is the most discussed open issue. Overall project health is strong, with rapid turnaround on high-risk bugs but a worrying accumulation of orphan processes in daemon mode.

## Releases

No new releases today. The latest public release remains v0.8.1 (or earlier); no changelog or migration notes to report.

## Project Progress

**Merged/closed pull requests in the last 24 hours (4 of 12 visible in sample):**

- **#8224** (closed) — `test(channels): add regression for JSON tool envelope preservation during proactive trim` – hardens channel orchestrator against tool_call_id loss.  
  [PR #8224](zeroclaw-labs/zeroclaw PR #8224)
- **#8084** (closed) — `fix(doctor): pass Config to provider_validation_error for custom providers` – enables custom providers in the `zeroclaw doctor` diagnostic.  
  [PR #8084](zeroclaw-labs/zeroclaw PR #8084)
- **#8190** (closed) — `fix(skills): stop flagging remote markdown documentation links in skill audit` – resolves a high false‑positive rate that blocked real skills.  
  [PR #8190](zeroclaw-labs/zeroclaw PR #8190)
- **#8201** (closed) — `ci: run docs_links_gate.sh in PR CI for docs changes` – automates link validation for documentation PRs.  
  [PR #8201](zeroclaw-labs/zeroclaw PR #8201)

**Closed issues (3):**

- **#8279** – [Bug] delegate bypasses parent's tool allowlist (S0 severity). Fixed and closed.  
  [Issue #8279](zeroclaw-labs/zeroclaw Issue #8279)
- **#6714** – [Feature] Remove remote-markdown-link block from skill audit (implemented via #8190).  
  [Issue #6714](zeroclaw-labs/zeroclaw Issue #6714)
- **#8219** – [Bug] gpt-oss-120b on Groq fails native multi-turn tool loops (closed without linked PR; fix assumed merged).  
  [Issue #8219](zeroclaw-labs/zeroclaw Issue #8219)

**Open PRs worth noting (advanced work in progress):**

- **#8158** – CycloneDX SBOM generation for Rust and npm (supply‑chain transparency).  
  [PR #8158](zeroclaw-labs/zeroclaw PR #8158)
- **#8173** – In-app upgrade with auto-restart from the web dashboard (gateway side).  
  [PR #8173](zeroclaw-labs/zeroclaw PR #8173)
- **#8325** – LAN peer discovery hints via mDNS (default‑off).  
  [PR #8325](zeroclaw-labs/zeroclaw PR #8325)

## Community Hot Topics

Issues with the most discussion activity (comments) today:

| Issue | Comments | Summary | Link |
|-------|----------|---------|------|
| **#5903** | 4 | MCP stdio child processes leak one orphan per heartbeat tick – highest severity open bug. | [Issue #5903](zeroclaw-labs/zeroclaw Issue #5903) |
| **#8238** | 3 | Feature request for an independent delegate mode so specialist agents run with their own policy/tools. | [Issue #8238](zeroclaw-labs/zeroclaw Issue #8238) |
| **#8279** | 2 | (closed) S0 security bypass – delegate can invoke tools the parent policy excludes. | [Issue #8279](zeroclaw-labs/zeroclaw Issue #8279) |
| **#7743** | 2 | Feature request for explicit target-profile authority in delegate handoffs (complement to #8238). | [Issue #7743](zeroclaw-labs/zeroclaw Issue #7743) |
| **#6714** | 2 | (closed) Skill audit false positives from `.md` documentation links. | [Issue #6714](zeroclaw-labs/zeroclaw Issue #6714) |

**Analysis**: The community is heavily focused on **delegation security and flexibility** (#8279, #8238, #7743). The MCP process leak (#5903) is a long‑standing runtime stability concern that draws consistent attention. No PR on this issue has been submitted yet, indicating a possible need for architectural discussion.

## Bugs & Stability

**Critical (S0) – Security:**
- **#8279** – Delegate tool allowlist bypass. **Closed** (fixed). Severity S0 – data loss / security risk.  
  [Issue #8279](zeroclaw-labs/zeroclaw Issue #8279)

**High – Runtime stability & data loss:**
- **#5903** – MCP stdio child processes leak on daemon with heartbeat. Still **open**. Default interval (30 min) yields ~48 orphans per day. No fix PR exists.  
  [Issue #5903](zeroclaw-labs/zeroclaw Issue #5903)
- **#8219** – Groq multi-turn tool loop failures (tool_call_id serialized as null, reasoning_content rejected). **Closed** (assumed fixed, no linked PR visible).  
  [Issue #8219](zeroclaw-labs/zeroclaw Issue #8219)

**Medium – Configuration/parsing:**
- **#8326** (PR, open) – UTF‑8 BOM in config.toml causes parse failure on Windows (ACP bridge).  
  [PR #8326](zeroclaw-labs/zeroclaw PR #8326)

**Low – Observability / test issues:**
- **#8146** (PR, open) – CLI one-shot loses telemetry and token totals on exit.  
  [PR #8146](zeroclaw-labs/zeroclaw PR #8146)

All high‑severity bugs found a fix or a PR within 24–48 hours except the MCP process leak, which remains unaddressed for over two months.

## Feature Requests & Roadmap Signals

**Actively discussed requests:**
- **Independent delegation mode** (#8238) – lets specialist agents run under their own tool/approval policy.  
- **Target‑profile authority for delegation** (#7743) – adds explicit allow/deny per handoff target.  
  Both are likely candidates for the **v0.8.3** milestone, as they align with the delegation work tracked in #8071.

**Other roadmap signals from open PRs:**
- **In-app upgrade from dashboard** (#8173) – a high‑visibility UX improvement for end users.
- **LAN peer discovery** (#8325) – networking enhancement for local multi‑node setups.
- **Allow HTTP URLs in browser tool** (#8136) – requested by users who run local HTTP services.
- **SBOM generation** (#8158) – supply‑chain security, likely a CI/audit requirement for a future release.

**Prediction**: The delegation authority features (#8238, #7743) and LAN discovery (#8325) are strong candidates for v0.8.3. The in‑app upgrade (#8173) may land sooner as a quality‑of‑life improvement.

## User Feedback Summary

**Real pain points reported:**
- **Resource exhaustion** – MCP daemon spawns one orphan process per heartbeat (rordd, #5903). Mitigation requires manual cleanup or process restarts.
- **Security bypass** – Delegate tool allowed sub‑agents to invoke prohibited tools (wangmiao0668000666, #8279). Rapidly fixed, restoring trust in delegation sandboxing.
- **Model compatibility** – Groq strictness broke multi‑turn tool loops (perlowja, #8219). Closed, but users may still face edge cases with other OpenAI‑compatible providers.
- **Skill audit friction** – False positives on `.md` documentation links blocked marketplace skill installations (xiongzubiao, #6714). Now resolved.

**User satisfaction signals:**
- Fast turnaround on high‑severity issues (hours to days).
- Transparent tracking via version‑specific trackers (#8070, #8071, #8072) suggests the team is responsive to user‑reported regressions.

**Dissatisfaction:**
- The open MCP leak (#5903) has been reported since April 19 and remains unaddressed – a growing source of frustration for daemon operators.

## Backlog Watch

| Item | Created | Status | Notes |
|------|---------|--------|-------|
| **#5903** – MCP stdio child process leak | 2026-04-19 | Open (accepted, no‑stale) | Oldest open issue with no fix PR. Priority P1, risk high. Lack of traction is concerning. |
| **#7743** – Target‑profile authority for delegation | 2026-06-15 | Open (accepted) | Related to #8238 which is in‑progress; likely to move forward soon. |
| **#8071, #8070, #8072** – v0.8.3 trackers | 2026-06-20 | Open (accepted) | No recent maintainer comments; coordinated via PRs. Healthy activity. |
| **#8181** – v0.8.2 release-support tracker | 2026-06-22 | Open (accepted) | 29 non‑plugin items listed – release appears imminent but no official date. |

**No PRs appear stuck without maintainer review.** All open PRs in the top‑20 list have at least one recent update. The main risk is the unresolved MCP leak, which if left to v0.8.3 could degrade daemon reliability for production users.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*