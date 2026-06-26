# OpenClaw Ecosystem Digest 2026-06-26

> Issues: 154 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-26 10:38 UTC

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

# OpenClaw Project Digest — 2026-06-26

---

## 1. Today’s Overview

OpenClaw saw **high activity** over the last 24 hours: **154 issues updated** (142 still open, 12 closed) and **500 pull requests updated** (438 open, 62 merged or closed). No new releases were published. The most active discussions centre on session write-lock timeouts, multi-agent orchestration instability, and security‑related regressions. The project remains under active maintenance with a steady stream of contributions, but several high‑impact bugs continue to block user workflows, especially around concurrency and gateway reliability.

---

## 2. Releases

**None.** No new versions or release notes were created in the reported period.

---

## 3. Project Progress

Despite the absence of a release, **62 pull requests were merged or closed** today. Notable examples from the top‑30 by comment count include:

- **[#96978]** – `fix(together): bound video generation JSON response reads`. Closed (merged) – prevents runaway memory usage from misbehaving Together AI endpoints.
- **[#89038]** – `fix: skip setup-only channel plugins in outbound resolution and drain pending deliveries on qqbot reconnect`. Still open, but demonstrates active work on channel reliability.
- **[#79753]** – `[Bug]: Cron announce fallback delivery reports success but message never arrives (WeChat + Feishu)`. Closed (likely resolved after investigation).

Other merged/closed PRs (not individually listed in the top 30) likely included smaller bug fixes, dependency updates, and documentation improvements. The high PR count suggests a large contributor pipeline, but many open PRs are awaiting author or maintainer review.

---

## 4. Community Hot Topics

The following issues generated the most discussion (top 5 by comment count):

| Issue | Title (abbreviated) | Comments | 💬 | Link |
|-------|---------------------|----------|----|------|
| #86538 | Session write-lock timeouts block subagent delivery lanes | 15 | P1, Diamond Lobster | [Issue #86538](https://github.com/openclaw/openclaw/issues/86538) |
| #43367 | Multi-agent orchestration unstable: config overwrites, session-lock failures, detached child work | 13 | P1, Diamond Lobster | [Issue #43367](https://github.com/openclaw/openclaw/issues/43367) |
| #71736 | [RFC] Control UI plugin contribution slots | 9 | P2, Off‑meta Tidepool | [Issue #71736](https://github.com/openclaw/openclaw/issues/71736) |
| #72015 | Active-memory blocks replies; QMD boot can overload multi-agent gateways | 8 | P1, Diamond Lobster | [Issue #72015](https://github.com/openclaw/openclaw/issues/72015) |
| #43996 | Sandbox container exits immediately with `no-new-privileges` | 7 | P1, Diamond Lobster | [Issue #43996](https://github.com/openclaw/openclaw/issues/43996) |

**Underlying needs:** The community is experiencing **reliability and concurrency pain** – session‑lock contention, configuration overwrites under parallel agent launches, and plugin‑induced gateway overload. These issues affect power users running multi‑agent setups (parallel batch coding, cron‑nested subagents, gateway proxy deployments). The RFC on Control UI plugin slots (#71736) signals a desire for a more modular, extensible chat surface that doesn’t require hacking core components.

---

## 5. Bugs & Stability

Critical and high‑severity bugs reported or updated in the last 24 hours:

| Issue | Severity | Summary | Fix PR exists? |
|-------|----------|---------|----------------|
| #86538 | 🔴 P1 | Session JSONL write‑lock timeouts block main, cron‑nested, and subagent lanes | Not yet linked |
| #43367 | 🔴 P1 | Multi‑agent orchestration: concurrent `agents add` config overwrites, session‑lock failures | Multiple linked PRs open (e.g., #73704) |
| #72015 | 🔴 P1 | Active‑memory plugin blocks replies, QMD boot overloads gateway | Not yet linked |
| #43996 | 🔴 P1 | Sandbox exits with `operation not permitted` when `no-new-privileges` enabled | Not yet linked |
| #73182 | 🔴 P1 | Reasoning default silently flipped to “on” for Claude – cost doubling | Not yet linked |
| #63664 | 🔴 P1 | Session flush blocks write tool completely during compaction | Not yet linked |
| #72418 | 🔴 P1 | `shouldSkipBackendSelfPairing` allows loopback clients to bypass device pairing (CVSS 9.3) | Not yet linked |
| #71699 | 🔴 P1 | Gateway crashes on Windows with `STATUS_STACK_BUFFER_OVERRUN` | Not yet linked |
| #72176 | 🟡 P2 | Intermittent duplicate message delivery across all channels after 2026.4.24 | Not yet linked |
| #71326 | 🟡 P1 | Cross‑exec stale file reads (vnode/dentry cache race) after 2026.4.20 | Not yet linked |

**Analysis:** The regression list is long and touches core subsystems (session persistence, multi‑agent orchestration, sandbox security, auth). Several issues have been open for weeks or months with no linked fix PR – most still carry the `needs-maintainer-review` or `needs-product-decision` label. The “diamond lobster” severity rating (the highest) on many indicates they are considered blockers by both the community and maintainers.

---

## 6. Feature Requests & Roadmap Signals

Top‑requested features from active issues:

| Issue | Feature | Priority | Likely in next version? |
|-------|---------|----------|--------------------------|
| #71142 | Configurable upload size limit for Control UI | P2 | Possible – low‑risk config change |
| #71736 | Control UI plugin contribution slots | P2 | Unlikely – still RFC stage, no implementation |
| #74077 | Slash command to set preview streaming mode per session | P3 | Low probability – niche use case |
| #50798 | Visible agent‑to‑agent messaging for ACP thread‑bound sessions | P2 | Possible if multi‑agent reliability improves |
| #70266 | Use assistant avatar in macOS Talk Mode overlay | P3 | Low – macOS‑only cosmetic |
| #71195 | OpenAI Realtime (speech‑to‑speech) path for Mac Talk Mode | P2 | Could be fast‑tracked – many upvotes, PR open (#73606) |
| #73537 | Production‑readiness stability labels on releases | P2 | Likely – community demand for clearer release quality |
| #71301 | Version‑matched bundled docs + native docs retrieval | P3 | Low – large scope |
| #50809 | Add SMS as a messaging channel and `sms.read` command | P2 | Possible – new channel plugin |
| #26037 | Ali Bailian coding plan support with reasoning | P2 | Stalled since Feb – unclear prioritisation |

**Signal:** The most voted feature (#26037, 4 👍) has been open since February without maintainer action. The spike in multi‑agent reliability issues may push the team to prioritise session‑level improvements before new channels.

---

## 7. User Feedback Summary

Real user experiences captured in recent issues:

- **Session lock hang:** “Session JSONL write‑lock timeouts block main, cron‑nested, and subagent lanes” – operator running parallel coding tasks on CLI ([#86538]).
- **Configuration overwrites:** “`openclaw agents add` appears unsafe when invoked concurrently: config gets overwritten repeatedly” – batch orchestration user ([#43367]).
- **Silent cost increase:** “Reasoning default silently flipped to on for Claude models … doubles Anthropic spend” – production gateway operator ([#73182]).
- **Platform regressions:** “Gateway crashes on Windows with `0xC0000409` … auto‑respawn frequently wedges” – Mattermost user on Windows ([#71699]).
- **Channel reliability:** “WhatsApp flaps and Telegram polling stalls on WSL2” – home automation user ([#73602]).
- **Cron delivery failures:** “Cron announce fallback reports success but message never arrives” – cron user on WeChat + Feishu ([#79753], now closed).

Overall sentiment is a mix of **enthusiasm** (many users rely on OpenClaw daily) and **frustration** with regressions that break existing workflows after minor version upgrades. The community provides detailed reproduction steps and even PRs, but several high‑severity issues remain unaddressed for weeks.

---

## 8. Backlog Watch

Issues and PRs that have been open for weeks or months and need maintainer attention:

| ID | Title | Open Since | Priority | Status |
|----|-------|------------|----------|--------|
| #26037 | Ali Bailian coding plan support with thinking | Feb 25, 2026 | P2 | No fix PR; no maintainer response in months |
| #42648 | Memory MVP: write pipeline with classification, dedupe, merge, conflict handling | Mar 11, 2026 | P2 | No fix PR; `needs-product-decision` |
| #48641 | Discord DMs: inbound messages silently dropped | Mar 17, 2026 | P2 | No fix PR; `needs-live-repro` |
| #43996 | Sandbox exits immediately with `no-new-privileges` | Mar 12, 2026 | P1 | No linked fix PR; `needs-maintainer-review` |
| #73182 | Reasoning default silently flipped for Claude | Apr 28, 2026 | P1 | `stale`; no fix PR |
| #72418 | Loopback auth bypass (CVSS 9.3) | Apr 26, 2026 | P1 | No fix PR; `needs-product-decision` |
| #89038 | qqbot reconnect: skip setup‑only channel plugins | Jun 1, 2026 | P1 | PR open, `waiting on author` |
| #94537 | Harden memory‑core dreaming daily‑file writes | Jun 18, 2026 | P0 | PR open, `waiting on author` |

**Watch list:** The security‑critical issue #72418 (auth bypass) has been without a fix PR for over two months despite a CVSS 9.3 rating. The P0 memory‑core fix #94537 is awaiting author updates. Many P1 bugs lack even a linked PR, suggesting the maintainer team may be stretched or prioritising different areas.

---

*Data as of 2026-06-26, based on GitHub issues and PRs updated in the last 24 hours from the `openclaw/openclaw` repository.*

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: Personal AI Assistant Open-Source Ecosystem

## 1. Ecosystem Overview

The personal AI assistant open-source landscape in mid-2026 is characterized by **rapid feature expansion** colliding with **foundational reliability challenges**. Projects are racing to deliver multi-agent orchestration, richer desktop experiences, and platform interoperability—while simultaneously wrestling with session concurrency bugs, security vulnerabilities, and integration regressions that erode user trust. A clear bifurcation is emerging: "heavyweight" frameworks (OpenClaw, IronClaw, ZeroClaw) target power users and production deployments with complex architectures, while more focused projects (NanoBot, PicoClaw) prioritize developer ergonomics and rapid iteration. The ecosystem remains fragmented, with no single project achieving dominant mindshare, though OpenClaw's reference implementation status continues to influence downstream forks and derivatives.

---

## 2. Activity Comparison (Last 24 Hours)

| Project | Issues Updated | PRs Updated | Releases Today | Health Score |
|---------|---------------|-------------|----------------|--------------|
| **OpenClaw** | 154 | 500 | None | 🟡 **Moderate** – High volume but many P1 bugs unaddressed, backlog growing |
| **NanoBot** | 18 | 37 | None | 🟢 **Strong** – Rapid security response, feature PRs advancing |
| **Hermes Agent** | 13 | 50 | None | 🟡 **Moderate** – High PR volume from one contributor, platform stability issues |
| **PicoClaw** | 3 | 16 | None | 🟢 **Strong** – Healthy merge rate, targeted bug fixes, dependency upkeep |
| **NanoClaw** | 2 | 22 | None | 🟢 **Strong** – Very high merge velocity (11/22 PRs), clear progress |
| **IronClaw** | 7 | 50 | None | 🟢 **Strong** – Reborn stack advancing, 15 PRs merged, focused bug fixes |
| **LobsterAI** | 1 | 6 | None | 🟢 **Strong** – All PRs merged, Cowork stability improvements shipped |
| **CoPaw (QwenPaw)** | 18 | 50 | **1 (v2.0.0-beta.1)** | 🟡 **Moderate** – Beta release addresses refactor; several critical bugs open |
| **ZeroClaw** | 5 | 50 | **1 (v0.8.2)** | 🟡 **Moderate** – New release with A2A/skills; low merge throughput relative to PR backlog |
| **NullClaw / TinyClaw / Moltis / ZeptoClaw** | 0 | 0 | None | ⚪ **Inactive** – No activity in 24 hours |

*Health Score reflects merge efficiency, security response, bug fix velocity, and backlog management.*

---

## 3. OpenClaw's Position

**Advantages vs. Peers:**
- **Reference implementation status**: As the core project from which many others fork or derive, OpenClaw sets architectural patterns for session management, multi-agent orchestration, and gateway design.
- **Scale of community**: 500 PRs and 154 issues in 24 hours dwarfs all peers; contributor pipeline is largest in ecosystem.
- **Feature breadth**: The most comprehensive set of capabilities—Control UI plugins, sandbox execution, multi-channel gateways, cron scheduling, and active-memory systems.

**Technical Approach Differences:**
- OpenClaw's **heavyweight architecture** (session write-locks, JSONL persistence, complex multi-agent orchestration) provides flexibility but introduces concurrency bugs that simpler projects (NanoBot, PicoClaw) avoid by design.
- Relies on **Go runtime** with extensive dependency chain, contrasting with NanoBot's Python + Node.js hybrid and IronClaw's Rust-based Reborn rewrite.

**Community Size Comparison:**
- OpenClaw: ~500 daily PRs, ~150 daily issues → **Largest contributor base**
- ZeroClaw/IronClaw: ~50 daily PRs → **Active mid-tier**
- NanoClaw/PicoClaw: ~15-22 daily PRs → **Healthy small teams**
- NanoBot/CoPaw: ~30-50 daily PRs → **Growing fast**
- Hermes Agent: 50 PRs driven largely by one contributor → **Dependency risk**

**Weakness**: High-severity bugs (session locks, auth bypass CVSS 9.3) remain unpatched for weeks, eroding the "reference quality" reputation. Meanwhile, NanoBot demonstrates 24-hour fix turnaround for critical security issues.

---

## 4. Shared Technical Focus Areas

Recurring themes across **multiple projects**:

| Technical Need | Affected Projects | Specific Requirements |
|---------------|-------------------|----------------------|
| **Multi-agent reliability** | OpenClaw, NanoClaw, LobsterAI, ZeroClaw | Session lock contention, config overwrites under parallel launches, subagent completion events |
| **Security hardening** | OpenClaw, NanoBot, NanoClaw, PicoClaw, ZeroClaw | Shell allowlist bypasses, sandbox escapes, auth bypass, encryption library upgrades (libolm→vodozemac) |
| **Approval flow maturity** | IronClaw, NanoClaw, ZeroClaw | Persistent "always allow" settings, multi-admin fallback, approval routing to distinct channels |
| **Desktop platform stability** | OpenClaw, Hermes Agent, NanoBot, CoPaw | Windows build failures, macOS certificate issues, update process hangs, font/zoom persistence |
| **Third-party model/provider compatibility** | OpenClaw, Hermes Agent, CoPaw, ZeroClaw | OpenAI-compatible endpoint issues, DeepSeek reasoning stability, custom provider URL handling |
| **Skill/plugin ecosystem** | OpenClaw, ZeroClaw, CoPaw, NanoBot | CRUD hooks, skill persistence across upgrades, typed slash commands, dependency conflict resolution |
| **Message/channel reliability** | OpenClaw, PicoClaw, CoPaw, Hermes Agent | WhatsApp timeout/reconnect, WeCom file delivery, Feishu long text failures, duplicate delivery |
| **Cost transparency** | OpenClaw, ZeroClaw | Silently flipped reasoning defaults causing cost doubling, live pricing from gateway |

**Observation**: No single project has solved multi-agent reliability—all implementations exhibit session-lock or completion-event issues. This represents the ecosystem's most critical open problem.

---

## 5. Differentiation Analysis

| Dimension | OpenClaw | NanoBot | IronClaw | ZeroClaw | CoPaw | NanoClaw | PicoClaw |
|-----------|----------|---------|----------|----------|-------|----------|----------|
| **Target User** | Power users, production operators | Developers, security-conscious teams | Enterprise teams | Plugin developers, composability-focused | Chinese ecosystem, WeChat/Feishu | Small teams, rapid deployment | Embedded/low-resource environments |
| **Architecture** | Monolithic Go + heavy deps | Lightweight Python + Node.js hybrid | Rust Reborn rewrite | Modular Rust + WASM plugins | Python with China-first channels | Python, CLI-centric | Minimal Go, ARM-friendly |
| **Key Strength** | Reference breadth, largest community | Security response speed (24h) | Architectural modernization | Plugin extensibility | China integration depth | Merge velocity, rapid fixes | Stability focus, low overhead |
| **Key Weakness** | Bug backlog, slow P1 fixes | Node.js dependency contention | Community still small | Review bottleneck (44 open PRs) | Beta instability, channel gaps | Small community size | Limited feature scope |
| **Release Maturity** | Stable but regression-prone | Pre-1.0 (v0.2.x) | Pre-1.0 (Reborn) | v0.8.x | v2.0.0-beta | ~v2.x | v0.2.x |
| **Primary Language** | Go | Python + Node | Rust | Rust | Python | Python | Go |

**Strategic Differentiation:**
- **IronClaw** is the only project betting entirely on Rust, positioning for performance and memory safety. Its Reborn rewrite, if completed, could leapfrog OpenClaw on stability.
- **ZeroClaw** uniquely invests in WASM component-model plugins, enabling language-agnostic extension—a potential long-term differentiator.
- **CoPaw** maintains exclusive depth in Chinese enterprise channels (WeCom, Feishu, DingTalk), carving a geographic niche.
- **NanoBot** differentiates on security responsiveness, critical for trust-sensitive deployments.
- **PicoClaw** targets resource-constrained environments (Raspberry Pi Zero) where OpenClaw cannot run.

---

## 6. Community Momentum & Maturity

**Tier 1: Rapidly Iterating (High Merge Velocity + Active Feature Development)**
- **NanoClaw** – 50% PR merge rate, new skills, Slack threading, migration fixes
- **LobsterAI** – 100% PR merge rate, Cowork stabilization
- **IronClaw** – 30% PR merge rate, Reborn stack advancing across multiple fronts

**Tier 2: High Volume, Moderate Stability**
- **OpenClaw** – Massive scale but merge throughput (~12%) doesn't keep pace with issue creation
- **ZeroClaw** – New release with A2A/skills but 44 open PRs suggest review bottleneck
- **CoPaw** – Beta release signals major refactor; bugs accumulating faster than fixes

**Tier 3: Stabilizing / Consolidating**
- **NanoBot** – High merge rate for security fixes, transition from bug-fixing to feature rollout
- **PicoClaw** – Sustained dependency upkeep and targeted fixes; growth trajectory is steady rather than explosive

**Tier 4: Inactive**
- NullClaw, TinyClaw, Moltis, ZeptoClaw – No activity indicates stalled or abandoned projects

**Sentiment Analysis:**
- **Enthusiasm**: Users across projects actively contribute repro steps, test fixes, and request features. Contribution culture is healthy.
- **Frustration**: Recurring themes of regressions after upgrades (OpenClaw, CoPaw), slow response to P1 bugs (OpenClaw), and stale feature requests (Hermes Agent #13181, OpenClaw #26037). The OpenClaw auth bypass (CVSS 9.3) sitting unpatched for 2 months is a significant trust concern.

---

## 7. Trend Signals

**For AI Agent Developers and Decision-Makers:**

1. **Security is now table stakes, not differentiator.** NanoBot's 24-hour fix turnaround sets an expectation that other projects will need to match. The libolm→vodozemac migration across PicoClaw and ZeroClaw signals industry movement away from unmaintained cryptographic dependencies.

2. **Multi-agent orchestration is the next frontier—and no one has solved it.** Session locking, completion event loss, and configuration corruption are the top blockers across OpenClaw, NanoClaw, LobsterAI, and ZeroClaw. Projects that deliver reliable subagent routing and durable approval flows will capture the production deployment market.

3. **Desktop platform stability is a make-or-break for mass adoption.** Three separate projects (OpenClaw, Hermes Agent, CoPaw) report Windows build/update issues. macOS certificate problems affect two projects. Enterprise desktop adoption requires platform-native behavior as a baseline, not a nice-to-have.

4. **"Ultra-lightweight" claims are under scrutiny.** NanoBot's removal of the "ultra-lightweight" label after community backlash reflects growing sophistication among users who audit dependency trees. Honest sizing and clear architectural trade-offs are valued over marketing language.

5. **Regional integration depth matters.** CoPaw's exclusive support for WeCom, Feishu, and DingTalk illustrates that global projects (OpenClaw) are losing ground to regionally-specialized forks in specific markets. Decision-makers should evaluate channel coverage relative to their deployment geography.

6. **The Rust migration trend is real.** IronClaw's Reborn and ZeroClaw's WASM host signal that projects investing early in memory safety and performance optimization (at the cost of ecosystem maturity today) may emerge as long-term winners as Go/Python projects accumulate technical debt.

7. **Cost visibility is becoming a user expectation.** OpenClaw's reasoning cost-doubling bug (#73182) and ZeroClaw's live pricing PR (#8233) indicate that users are demanding transparent billing. Future projects that provide built-in cost tracking and quota management will have an adoption advantage.

**Bottom Line**: The ecosystem is healthy but fragmented. No single project offers production-grade multi-agent reliability today. The next 3-6 months will likely see consolidation as users gravitate toward projects that solve the session concurrency problem while maintaining security response parity with NanoBot's demonstrated cadence.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-06-26

## 1. Today’s Overview

The project is experiencing a **security-focused cleanup wave**, with 14 issues closed and 11 PRs merged/closed in the past 24 hours. The majority of closed items involve critical shell command allowlist bypasses, MCP scope bypasses, and filesystem permission flaws — all addressed within the same day of disclosure. Meanwhile, 26 open PRs and 4 open issues indicate strong forward momentum on features (ask_clarification, PWA support, MCP idle timeout, webhook triggers). The **ultra-lightweight branding controversy** (Issue #660) remains the most commented-open item, reflecting ongoing user scrutiny of project claims. Overall, the project demonstrates rapid incident response and active feature development, though dependencies on Node.js continue to raise consistency questions.

## 2. Releases

No new releases were published today (2026-06-26). The latest stable version remains at a tag prior to today’s fixes. Given the volume of security patches merged, a release (likely v0.2.2 or v0.3.0) is expected shortly.

## 3. Project Progress

Three pull requests were merged or closed today, all addressing security or documentation:

- **PR #4099** (merged) — *Keep filesystem extra roots read-only*. Fixes #4073, ensuring `extra_allowed_dirs` are not treated as writable for `write_file`/`edit_file`.  
  **Link**: https://github.com/HKUDS/nanobot/pull/4099

- **PR #4524** (merged) — *Fix enabledTools filtering for MCP resources and prompts*. Addresses the bypass reported in Issue #4519.  
  **Link**: https://github.com/HKUDS/nanobot/pull/4524

- **PR #4536** (merged) — *Remove misleading “ultra-lightweight” claim from README*.  
  **Link**: https://github.com/HKUDS/nanobot/pull/4536

Additionally, several critical open PRs (e.g., #4525, #4526, #4522, #4506) are nearing readiness and already carry “valid” tags.

## 4. Community Hot Topics

- **Issue #660** — *Project claims to be “ultra-lightweight” but includes bloated Node.js dependency*  
  - 👍 5, 💬 12 comments  
  - Created Feb 2026, still open. The community is debating whether the Dockerfile’s dual‑runtime requirement (Python + Node.js) contradicts the “ultra-lightweight” tagline. The linked PR #4536 (merged today) removes the claim, signalling maintainer agreement.  
  **Link**: https://github.com/HKUDS/nanobot/issues/660

- **Issue #2439** — *Malicious data‑exfiltration code in litellm_init.pth*  
  - 👍 4, 💬 6 comments  
  - A critical security incident that was closed after investigation and mitigation. Despite the closure, user trust concerns linger.  
  **Link**: https://github.com/HKUDS/nanobot/issues/2439

- **Security disclosure cluster** (Issues #4514–#4520) — Multiple allowlist bypass techniques for the `exec` tool and MCP tool scoping, reported by **YLChen-007**. All were closed within 24 hours with matching fix PRs. This coordinated responsible disclosure is a positive sign for project transparency.  
  **Links**: #4514, #4515, #4516, #4519, #4520

- **Issue #4508** — *Feature request: add ask_clarification tool*  
  - 💬 1 comment, 👍 0  
  - A feature that addresses real user pain points when instructions are ambiguous. A PR (#4527) is already open.  
  **Link**: https://github.com/HKUDS/nanobot/issues/4508

## 5. Bugs & Stability

| Severity | Bug | Status | Fix PR(s) |
|----------|-----|--------|-----------|
| **Critical** | MCP `enabledTools` scope bypass exposes resource/prompt wrappers (#4519) | Closed | #4524 (merged) |
| **Critical** | `exec.allowPatterns` whitelist bypass via chained commands (#4514, #4515, #4516, #4520) | Closed | #4526 (open) |
| **Critical** | Filesystem `extra_allowed_dirs` writable (#4073) | Closed | #4099 (merged) |
| **High** | `exec` login-shell default leaks secrets from `.bash_profile` (ref #4518) | Open (PR #4525) | #4525 (open) |
| **Medium** | Windows `--background` restarts inconsistent with JSON file (#4511) | Open | None yet |
| **Low** | DingTalk timeout / unsupported rich‑text (#4497) | Closed | #4497 (PR?) |
| **Low** | WebUI voice transcription fails with Xiaomi MiMo provider (#4492) | Closed | #4492 (PR) |

Most critical bugs were reported and resolved the same day, demonstrating a strong security response cadence. The Windows gateway restart bug (#4511) is the only open issue without an immediate fix PR.

## 6. Feature Requests & Roadmap Signals

The following feature requests have corresponding open PRs and are strong candidates for the next release:

- **ask_clarification tool** (#4508, PR #4527) — Short‑circuits agent turns when user input is ambiguous.
- **PWA support + mobile sidebar swipes** (#4479, PR #4494) — Enhances WebUI on mobile devices.
- **read‑only search_history tool** (#4439, PR) — Enables agent to recall past conversations.
- **MCP server idle timeout auto‑kill** (PR #4506) — Prevents zombie processes and memory leaks.
- **Gateway webhook triggers** (PR #4502) — Extends gateway to handle inbound webhooks from external services.
- **Skills in subdirectories** (PR #4504) — Organizational improvement for skill management.
- **Custom provider thinking style configuration** (Issue #4429, closed) — Already merged, allows non‑OpenAI reasoning parameters.

The **ask_clarification tool** and **MCP idle timeout** appear highest priority given user demand and resource leak concerns. Expect both in the next release.

## 7. User Feedback Summary

- **“Ultra‑lightweight” claim backlash**: Issue #660 shows that users are paying close attention to project claims; the PR to remove the phrase (#4536) was merged today, indicating maintainers listen to feedback.
- **Security trust erosion**: Issue #2439 (malicious code bundle) generated strong negative reactions. While the package was cleaned, users may remain wary of supply‑chain integrity.
- **Reliability concerns**: Issue #1710 (frequent “no response to give” with Qwen 3.5) and #4242 (dream function misbehaviour) indicate occasional UI/agent loop frustrations. Both closed, but recurring complaints could resurface.
- **Integration pain points**: DingTalk users reported timeouts (#4497) and Xiaomi MiMo voice issues (#4492) – both resolved quickly, showing attentiveness to third‑party channel stability.
- **Positive sentiment**: Rapid closure of security issues and merging of community PRs (e.g., #4099, #4536) suggests a responsive and welcoming contribution environment.

## 8. Backlog Watch

- **Issue #660** (created 2026‑02‑14) — Though now partially addressed by README removal, the underlying runtime‑bloat concern remains. No formal discussion on whether Node.js can be reduced to an optional dependency. Maintainers should consider a dedicated issue to track lightweight image efforts.
- **No other long‑unanswered issues**: All other open items are from late June 2026. The maintainers appear to be keeping the backlog well below the 1‑month threshold. No stale PRs requiring attention were identified.

**Overall Assessment**: Project is in a healthy state with a high pace of vulnerability fixes and feature development. The security incidents were handled transparently and quickly. The removal of the “ultra‑lightweight” claim shows responsiveness to community criticism. The coming release should solidify security improvements and introduce several user‑requested features.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-06-26

## 1. Today's Overview
The project shows **high activity** with **13 issues** (12 open, 1 closed) and **50 pull requests** (49 open, 1 merged/closed) updated in the last 24 hours. A burst of desktop feature contributions from community member **professorpalmer** drives the PR volume, covering source control, plan mode, edit review, and in-app browser panels. Bug reports focus on platform-specific build failures, missing keyboard shortcuts, and async delegation reliability. No new releases were published today.

---

## 2. Releases
No new releases were created on 2026-06-26.

---

## 3. Project Progress
One pull request was merged or closed today. The specific PR is not listed in the provided top-20 data, but the project's merge pipeline remains active. No new features or fixes were explicitly recorded as merged; most contributed PRs remain open for review.

---

## 4. Community Hot Topics
The most active discussions and highly-upvoted items reveal user interest in:

- **OpenCode Go model integration** – [Issue #13181](https://github.com/NousResearch/hermes-agent/issues/13181) (👍12, 4 comments) asks for a documented way to add lightweight, self-hosted model backends. The request has been open since April 20 and remains unanswered by maintainers.
- **Desktop rebuild failure on Windows** – [Issue #46677](https://github.com/NousResearch/hermes-agent/issues/46677) (👍6, 4 comments) describes a broken `@assistant-ui/tap` export that blocks `vite build`. Duplicate of #46692, indicating a widespread regression.
- **Auxiliary title generation 404 on custom Anthropic-mode providers** – [Issue #19753](https://github.com/NousResearch/hermes-agent/issues/19753) (👍2, 6 comments) reports a URL double-slash bug affecting Kimi and MiniMax providers using `custom:` provider definitions.
- **Desktop build fails on Windows (package-lock + junction path)** – [Issue #44902](https://github.com/NousResearch/hermes-agent/issues/44902) (👍1, 5 comments) documents a bootstrap installer failure on Windows 11.

The **professorpalmer** PR series (e.g., #48777, #48781, #48813, #48818, #48821, #48825, #48842, #48849) has garnered attention for adding VS Code/Cursor-style features like Source Control, Plan mode, and edit review, though none have official maintainer review comments yet.

---

## 5. Bugs & Stability
New bugs reported today, ranked by potential impact:

| Issue | Component | Severity | Description | Fix PR Exists? |
|-------|-----------|----------|-------------|----------------|
| [#53027](https://github.com/NousResearch/hermes-agent/issues/53027) | cron/agents | **High** | Async delegation completion events lost in cron jobs – subagent results never reach the parent agent. | None yet |
| [#53026](https://github.com/NousResearch/hermes-agent/issues/53026) | Matrix adapter | **Medium** | Hardcoded `MAX_MESSAGE_LENGTH=4000` breaks Markdown tables; Matrix spec allows 65KB. | None yet |
| [#53021](https://github.com/NousResearch/hermes-agent/issues/53021) | terminal/auth | **Medium** | Feature request (deny-by-default terminal allowlist) but also highlights security gap for unattended agents. | PR [#53021](https://github.com/NousResearch/hermes-agent/pull/53021) (open) |

**Previously reported bugs still active:**

- **Desktop zoom not working on Windows** ([#37917](https://github.com/NousResearch/hermes-agent/issues/37917), P3) – no fix yet.
- **Desktop "Update" button hangs on Windows** ([#48854](https://github.com/NousResearch/hermes-agent/issues/48854)) – closed today, fix implied but not confirmed in data.
- **Early-sent queued message disappears from transcript** ([#48825](https://github.com/NousResearch/hermes-agent/pull/48825)) – a fix PR is open from professorpalmer.
- **Desktop build fails on macOS** ([#46692](https://github.com/NousResearch/hermes-agent/issues/46692), duplicate of #46677) – missing `./react-shim` export from `@assistant-ui/tap`.

---

## 6. Feature Requests & Roadmap Signals
New feature requests filed today indicate growing demand for:

- **Session-scoped terminal allowlist mode** ([#53021](https://github.com/NousResearch/hermes-agent/issues/53021)) – users want to run untrusted agents as monetized services (e.g., x402 micropayments) with deny-by-default terminal access. This could become a high-priority security feature.
- **Ctrl+PageUp/PageDown for session switching** ([#53017](https://github.com/NousResearch/hermes-agent/issues/53017)) – familiar tab-navigation pattern for desktop users.
- **10 new Discord channel/role management actions** ([#53019](https://github.com/NousResearch/hermes-agent/issues/53019)) – request for extended admin capabilities in the Discord gateway.

Existing active requests that may land in the next desktop release (based on PR activity):

- Persistent font size / zoom setting ([#51918](https://github.com/NousResearch/hermes-agent/issues/51918))
- OpenCode Go model integration ([#13181](https://github.com/NousResearch/hermes-agent/issues/13181))
- Native local OSS web stack (SearXNG + Crawl4AI) – PR [#6325](https://github.com/NousResearch/hermes-agent/pull/6325) is open.

The professorpalmer PRs (source control, plan mode, edit review) strongly suggest the next release will ship an **IDE-like desktop experience** if those PRs are merged.

---

## 7. User Feedback Summary
**Pain points voiced by the community:**

- **Desktop font size is too small** – multiple users (e.g., [#51918](https://github.com/NousResearch/hermes-agent/issues/51918)) report that the default 13px is uncomfortable and zoom resets on restart.
- **Windows build and update are fragile** – issues [#44902](https://github.com/NousResearch/hermes-agent/issues/44902) and [#48854](https://github.com/NousResearch/hermes-agent/issues/48854) (now closed) highlight deadlocks from file locks and junction path problems.
- **Matrix messages break tables** – [#53026](https://github.com/NousResearch/hermes-agent/issues/53026) shows that the conservative 4KB limit damages formatting unnecessarily.
- **Cron job delegation is unreliable** – [#53027](https://github.com/NousResearch/hermes-agent/issues/53027) reports lost completion events, which could be a critical issue for production automation.

**Use cases driving new features:**
- Unattended/monetized agents (x402 pay-per-turn) – [#53021](https://github.com/NousResearch/hermes-agent/issues/53021)
- Self-hosted/lightweight model backends – [#13181](https://github.com/NousResearch/hermes-agent/issues/13181)
- Desktop as a code editor substitute – see professorpalmer PRs

Overall sentiment is **active and constructive** but with frustration around platform stability (Windows/macOS builds) and missing quality-of-life settings.

---

## 8. Backlog Watch
Issues that have been open for an extended period without maintainer resolution:

| Issue | Date Created | Last Update | Comments | Summary |
|-------|--------------|-------------|----------|---------|
| [#13181](https://github.com/NousResearch/hermes-agent/issues/13181) | 2026-04-20 | 2026-06-26 | 4 | OpenCode Go model integration (👍12) – no maintainer response. |
| [#19753](https://github.com/NousResearch/hermes-agent/issues/19753) | 2026-05-04 | 2026-06-26 | 6 | Anthropic-mode provider double-/v1 bug – no fix merged. |
| [#44902](https://github.com/NousResearch/hermes-agent/issues/44902) | 2026-06-12 | 2026-06-26 | 5 | Windows build fails – stale package-lock + junction paths. |

These issues represent long-standing technical debt and community needs that have not yet been addressed by maintainers, risking user frustration with the desktop platform and model extensibility.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest – 2026-06-26

## 1. Today's Overview

Project activity remains high, with **3 issues** and **16 pull requests** updated in the last 24 hours. Nine PRs were merged or closed, indicating a focused effort on bug fixes, resource‑leak corrections, and routine dependency upgrades. No new releases were published today. The community is actively reporting real‑world problems (especially around channel scheduling and WhatsApp connectivity), while maintainers are responding with targeted fix PRs. Overall, the project shows healthy maintenance velocity and steady progress toward stability.

## 2. Releases

No new releases were created in the last 24 hours.

## 3. Project Progress (Merged/Closed PRs Today)

Nine PRs were merged or closed today, spanning three main categories:

**Bug Fixes & Code Quality**  
- [#3172 – fix: explicitly ignore Close() errors in error paths and retry loops](https://github.com/sipeed/picoclaw/pull/3172)  
  Ignores secondary `Close()` errors in 8 call sites across 4 files where the primary error already provides sufficient context.  
- [#3170 – fix(agent): close base64 encoder on io.Copy error path](https://github.com/sipeed/picoclaw/pull/3170)  
  Ensures the base64 encoder’s internal buffer is always flushed/closed after `io.Copy`, preventing resource leaks.  
- [#3092 – fix(skills_install): add ok checks for version and force type assertions](https://github.com/sipeed/picoclaw/pull/3092)  
  Adds missing `ok` checks in type assertions for `version` and `force` arguments, avoiding silent zero‑value fallbacks.

**Dependency Updates (Dependabot)**  
- [#3176 – build(deps): bump github.com/mymmrac/telego from 1.9.0 to 1.10.0](https://github.com/sipeed/picoclaw/pull/3176)  
- [#3175 – build(deps): bump fyne.io/systray from 1.12.1 to 1.12.2](https://github.com/sipeed/picoclaw/pull/3175)  
- [#3174 – build(deps): bump github.com/line/line-bot-sdk-go/v8 from 8.20.0 to 8.20.1](https://github.com/sipeed/picoclaw/pull/3174)  
- [#3173 – build(deps): bump modernc.org/sqlite from 1.51.0 to 1.53.0](https://github.com/sipeed/picoclaw/pull/3173)  
- [#3145 – build(deps): bump github.com/github/copilot-sdk/go from 0.2.0 to 1.0.2](https://github.com/sipeed/picoclaw/pull/3145)  

**Documentation**  
- [#1892 – docs: add gitcgr code graph badge](https://github.com/sipeed/picoclaw/pull/1892)  
  Adds a visual code‑graph badge to the README, improving repository discoverability.

## 4. Community Hot Topics

The most active discussions involve two issues:

- **#1757 (CLOSED) – [BUG] when I ask my agent to do a task every hour of the day I now get channel error**  
  *Author: dhensen* | [Issue Link](https://github.com/sipeed/picoclaw/issues/1757)  
  *Comments: 10* | *Reactions: 0*  
  This long‑running issue (created March 18, closed today) detailed a channel‑specific error when using scheduled hourly tasks via Telegram. The cause appears resolved; the high comment count reflects sustained community interest in cron‑based agent scheduling.

- **#3088 (OPEN) – [Feature] use vodozemac instead of libolm**  
  *Author: pbsds* | [Issue Link](https://github.com/sipeed/picoclaw/issues/3088)  
  *Comments: 3* | *Reactions: 👍 2*  
  A security‑critical feature request to replace the unmaintained `libolm` with the official `vodozemac` library. The community is actively supporting this change, and it has been labelled `help wanted` and `priority: high`.

## 5. Bugs & Stability

**New Bug Report (High Severity)**  
- **#3178 (OPEN) – [BUG] WhatsApp Websocket Timeout**  
  *Author: Jh123x* | [Issue Link](https://github.com/sipeed/picoclaw/issues/3178)  
  *Environment:* PicoClaw v0.2.9, Go 1.25.11, deepseek‑v4‑pro, Docker/launchpad, WhatsApp channel.  
  A user reports that after connecting to WhatsApp via WebSocket and adding a scheduled task, the connection times out. No comments or fix yet.

**Direct Fix PR for WhatsApp Connectivity**  
- **#3179 (OPEN) – fix(whatsapp): reconnect after websocket drops**  
  *Author: Alix-007* | [PR Link](https://github.com/sipeed/picoclaw/pull/3179)  
  Implements automatic reconnection, read deadlines, and asynchronous message dispatching for the WhatsApp bridge. This PR directly addresses the root cause of #3178.

**Additional Stability Fixes (merged today)**  
- #3172 (ignoring spurious Close errors) and #3170 (base64 encoder leak) both reduce resource‑leak risks under error conditions.

## 6. Feature Requests & Roadmap Signals

- **#3088 – Replace libolm with vodozemac** (see above) – Security modernization likely to be included in the next minor release given its `priority: high` label.  
- **#3063 (OPEN PR) – feat: add deltachat gateway**  
  *Author: trufae* | [PR Link](https://github.com/sipeed/picoclaw/pull/3063)  
  A new gateway for Delta Chat, still open and updated today. Extends PicoClaw’s multi‑channel support. If merged, it will debut in an upcoming release.  
- **#3177 (OPEN, Dependabot) – build(deps): bump github.com/github/copilot-sdk/go from 0.2.0 to 1.0.4**  
  [PR Link](https://github.com/sipeed/picoclaw/pull/3177)  
  This major version jump (0.2.0 → 1.0.4) suggests significant Copilot SDK changes. Once merged, it could enable new Copilot integration features but may also require code adaptations.

## 7. User Feedback Summary

- **Scheduled Task Channel Errors** (#1757) – A user running PicoClaw on a Raspberry Pi Zero W with Telegram and Ollama encountered persistent channel errors when asking agents to perform hourly tasks. The issue is now closed, implying a fix has been deployed.  
- **WhatsApp Timeout** (#3178) – A Docker user with deepseek‑v4‑pro reports timeouts immediately after adding a schedule to WhatsApp. This indicates that the WhatsApp WebSocket connection is fragile when combined with cron jobs.  
- **Security Concern** (#3088) – A contributor explicitly calls `libolm` “unmaintained and insecure,” reflecting a community desire for proactive security hardening. No user satisfaction comments were recorded today.

## 8. Backlog Watch

Several important items lack recent maintainer action:

- **#3088 – Feature: use vodozemac** – Open since June 9, with 2 upvotes but no assignee. Labelled `help wanted` and `priority: high`.  
- **#3063 – PR: add deltachat gateway** – Open since June 8, no recent review comments despite today’s update.  
- **#3142 – fix(spawn): clear ForUser in sub-turn ToolResult** – Open since June 17, marked `stale`. No maintainer feedback.  
- **#3128 – fix(web): explicitly ignore resp.Body.Close() errors** – Open since June 15, no engagement.  
- **#3177 – Dependabot copilot-sdk bump** – A major dependency update that requires careful review to avoid breaking changes.

These items, especially the security‑critical #3088, would benefit from prompt maintainer triage.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-26

## 1. Today's Overview
NanoClaw shows **very high development velocity** today: 22 pull requests were updated, with 11 merged or closed and 11 still open. The activity is concentrated on **stability fixes** (migration crashes, security confinement, macOS compatibility) and **feature expansion** (new operational skills, approval enhancements, Slack threading). Only two issues were touched, one closed and one new feature request, suggesting the team is focused on addressing community-reported bugs and prepping the next release. Overall project health is strong, with rapid review cycles and a steady stream of contributions from multiple developers.

## 2. Releases
No new releases were published today.

## 3. Project Progress
**11 PRs were merged or closed today**, advancing the following areas:

- **Data Migration Fix**: [#2859](nanocoai/nanoclaw PR #2859) — Fixes a crash in the v1→v2 DB migration when the `is_main` column is missing, preventing cascading failures in session and task tables. **Critical for users upgrading from early v1 versions.**
- **Slack Threading**: [#2472](nanocoai/nanoclaw PR #2472) and [#2471](nanocoai/nanoclaw PR #2471) — Two related PRs now give Slack DMs proper per-thread session isolation and in-thread bot replies. This resolves a long-standing collapse issue where all top-level messages shared one session.
- **Approval Flow**: [#2832](nanocoai/nanoclaw PR #2832) — Adds an optional "Reject with reason" path to modular approval cards, allowing approvers to provide feedback that the requesting agent can interpret and adapt to.
- **Security Fix**: [#2817](nanocoai/nanoclaw PR #2817) — Confines `send_file` reads to the workspace using strict canonical-path validation and blocks symlink escapes. **Important security hardening.**
- **Router & CLI Stability**: [#2815](nanocoai/nanoclaw PR #2815) prevents `safeParseContent` from treating JSON primitives as objects; [#2813](nanocoai/nanoclaw PR #2813) fixes the CLI socket response cap to count bytes instead of characters, avoiding truncation on multi-byte UTF-8 payloads.
- **macOS OneCLI Fix**: [#2854](nanocoai/nanoclaw PR #2854) — Redirects `TMPDIR` so gateway CA certificates mount correctly into containers on macOS (Rancher Desktop / Apple container). Resolves self-signed certificate errors.
- **New Skill**: [#2843](nanocoai/nanoclaw PR #2843) — Adds `/learn`, a skill that distills a reusable skill from any directory, URL, or pasted content.
- **Test/Misc**: [#2867](nanocoai/nanoclaw PR #2867) — A test finding PR (closed, minor).

## 4. Community Hot Topics
No issues or PRs attracted more than 0 comments today, but open PRs indicate strong community interest in:

- **Operational Skills**: [#2863](nanocoai/nanoclaw PR #2863) (`/setup-system-digest` and `/system-digest`) and [#2862](nanocoai/nanoclaw PR #2862) (`/manage-agents` and `/manage-schedules`) — both by **grantland**, adding utility and operational tools that reduce the need for raw CLI access.
- **Environment Variable Expansion in MCP**: [#2861](nanocoai/nanoclaw PR #2861) — Expands `${VAR_NAME}` references in MCP server env at spawn time, making configuration more dynamic.
- **Session Rotation**: Two PRs ([#2865](nanocoai/nanoclaw PR #2865) and [#2864](nanocoai/nanoclaw PR #2864)) address stale session handling in OpenCode and provider connections, rotating sessions on ceiling-kill signals or age thresholds.

The underlying theme is **operational maturity**: users want more self-service controls and fewer manual restarts.

## 5. Bugs & Stability
No new bug reports were filed as issues today, but several severity-level bugs were fixed via PRs:

| Severity | Bug / Fix | PR / Issue |
|----------|-----------|------------|
| **High** | **Migration crash** on older v1 installs (no `is_main` column) prevents v2 DB creation, cascading to session/task steps. Fixed by [#2859](nanocoai/nanoclaw PR #2859). | Merged |
| **High** | **Security: `send_file` workspace escape** via symlinks. Fixed by [#2817](nanocoai/nanoclaw PR #2817) with realpath validation. | Merged |
| **Medium** | **macOS CA certificate failure** in containers – all agent API calls fail with self-signed cert error. Fixed by [#2854](nanocoai/nanoclaw PR #2854). | Merged |
| **Medium** | **Discord attachments not reaching agents** – images and text files appear as bare filenames. Fix open in [#2752](nanocoai/nanoclaw PR #2752) (open, needs review). | Open |
| **Medium** | **Legacy Telegram markdown sanitizer** conflicts with MarkdownV2 adapter. Fix by [#2866](nanocoai/nanoclaw PR #2866) (open). | Open |
| **Low** | **libsignal debug spam** leaking key material in logs. Fix by [#2860](nanocoai/nanoclaw PR #2860) (open). | Open |
| **Low** | **Stale "Global Memory" instruction** in seed prompt. Fix by [#2824](nanocoai/nanoclaw PR #2824) (open). | Open |

No regressions were reported.

## 6. Feature Requests & Roadmap Signals
The **most significant user-requested feature** is **multi-admin approval** ([#2857](nanocoai/nanoclaw Issue #2857) – open, created yesterday). Currently approval requests target a single admin; if unavailable, others cannot approve. The request includes re-asking for approval from the next admin and CLI-based approval for machine owners. This is likely to land in the next minor release (v2.x), given the related approval-with-reason feature was just merged.

Other roadmap signals from open PRs:
- **New Skills**: `/setup-system-digest`, `/system-digest`, `/manage-agents`, `/manage-schedules` (all by grantland) point to a **dashboard-lite** approach via chat commands, reducing reliance on external admin UIs.
- **Dynamic MCP Configuration** ([#2861](nanocoai/nanoclaw PR #2861)) enables environment variable interpolation, increasing deployment flexibility.
- **Session Rotation Automation** ([#2864](nanocoai/nanoclaw PR #2864), [#2865](nanocoai/nanoclaw PR #2865)) suggests the team is tackling long-running session staleness proactively.

The closed issue [#1275](nanocoai/nanoclaw Issue #1275) (auto-prompt registration when added to new Telegram group) was resolved today — likely merged or rejected, indicating responsiveness to community onboarding friction.

## 7. User Feedback Summary
Real user pain points surfaced in recent issues and PRs:

- **Onboarding UX**: In [#1275](nanocoai/nanoclaw Issue #1275), a user reported that when NanoClaw is added to a new Telegram group, it silently ignores messages. The request is for automatic detection and a clear prompt to register from the main chat. This was addressed (closed), improving first-time user experience.
- **Approval Bottleneck**: [#2857](nanocoai/nanoclaw Issue #2857) highlights that single-admin approvals block workflows when that admin is unavailable. The community wants fallback to other admins and/or CLI approval.
- **macOS Users**: Multiple users on Rancher Desktop faced API connection failures due to certificate path issues ([#2854](nanocoai/nanoclaw PR #2854)). The fix was merged quickly, showing responsiveness to platform-specific pain.
- **Discord Integration**: PR [#2752](nanocoai/nanoclaw PR #2752) addresses missing attachment content for Discord users — images and text appear as useless placeholders. This remains open and is likely a top annoyance for Discord channel operators.

Overall sentiment appears **positive**: contributors are actively submitting fixes and maintainers are merging them rapidly. The approval feature expansion and new skills indicate growing adoption among teams managing multiple agents.

## 8. Backlog Watch
No critical long-unanswered items were found in today’s data. However, two open PRs merit attention:

- **[#2752](nanocoai/nanoclaw PR #2752)** — Discord inbound attachments fix (opened June 12, still open). This affects a core integration and has no recent activity from reviewers. If not merged soon, users relying on Discord channels may continue to see degraded functionality.
- **[#2824](nanocoai/nanoclaw PR #2824)** — Drop stale "Global Memory" instruction (opened June 20). This is a low-severity cleanup but has been open for 6 days without merge. The instruction may confuse new agents, so a timely review is advisable.

The issue [#1275](nanocoai/nanoclaw Issue #1275) was closed today after being open since March 19, illustrating that the maintainers do eventually address older items, but the response time (3+ months) suggests the project could benefit from a more systematic backlog triage process.

**No issues or PRs remain unanswered from before June 2026.**

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-06-26

## 1. Today’s Overview
Activity remains very high, with 50 PRs touched in the last 24 hours and 7 issues updated, though no new releases were cut. The project is firmly focused on the **Reborn stack** — a major architectural rewrite — with ongoing PRs addressing capability policy, persistent approvals, Slack host bindings, storage stress testing, and build/dependency housekeeping. The community submitted several usability bugs around tool approval flows, while core contributors continued chipping away at the large “capability-policy” epic and test infrastructure porting. Overall project health appears solid, with a balanced mix of feature work, bug fixes, and CI improvements.

## 2. Releases
**None.** No new releases were published today.

## 3. Project Progress
Fifteen PRs were merged or closed today, including:

- **CI/branch-protection fixes**  
  - [#5317](https://github.com/nearai/ironclaw/pull/5317) — Restored the `Run Tests` required check marker for branch protection.  
  - [#5308](https://github.com/nearai/ironclaw/pull/5308) — Moved legacy test suite triggers to nightly CI, keeping PR CI lean.  
  - [#5281](https://github.com/nearai/ironclaw/pull/5281) — Unblocked `main` by fixing libsql feature, apt retry, fail-fast, and `.codegraph` issues.

- **Bug fixes targeting Reborn WebUI v2**  
  - [#5299](https://github.com/nearai/ironclaw/pull/5299) — Anchored run failure messages to their correct positions on timeline refreshes.  
  - [#5309](https://github.com/nearai/ironclaw/pull/5309) — Fixed “Approve & always allow” persistence for shared registry tools like `nearai.web_search`.  
  - [#5306](https://github.com/nearai/ironclaw/pull/5306) — Broke the approval resume loop when a one-shot capability lease satisfies the ask-each-time gate.  
  - [#5310](https://github.com/nearai/ironclaw/pull/5310) — Surfaced schema-validation failures as `model_error` instead of generic `driver_protocol_violation`.  
  - [#5297](https://github.com/nearai/ironclaw/pull/5297) — Suppressed stale blocked-gate rows in WebUI v2 stream projections.  
  - [#5252](https://github.com/nearai/ironclaw/pull/5252) — Fixed Slack host conversation bindings to persist through durable service restarts.

- **Feature/architecture PRs that remain open (notable advances)**  
  - [#5313](https://github.com/nearai/ironclaw/pull/5313) — Added a storage stress harness for filesystem and SQL backends.  
  - [#5305](https://github.com/nearai/ironclaw/pull/5305) — Ported legacy e2e/Playwright tests to the Reborn harness.  
  - [#5244](https://github.com/nearai/ironclaw/pull/5244) — Removed generated `dist/` artifacts from source control, building them into Cargo `OUT_DIR`.  
  - [#5312](https://github.com/nearai/ironclaw/pull/5312) — Build Reborn CLI as a single default binary bundling WebUI v2, Slack, OpenAI-compatible routes, and PostgreSQL.

## 4. Community Hot Topics
Only a handful of issues attracted comments today, indicating that the community is still relatively small or that active discussion happens mostly on PRs. The most commented items were:

- **[#5283](https://github.com/nearai/ironclaw/issues/5283) — “Approve & always allow” not persisted for `nearai.web_search`**  
  2 comments. User sunglow666 reported that after checking the “always allow” checkbox, the setting is lost on next tool invocation. The underlying need is for a reliable per-tool persistent approval, which the team addressed in the open PR [#5309](https://github.com/nearai/ironclaw/pull/5309).

- **[#5272](https://github.com/nearai/ironclaw/issues/5272) — REST-created local users for capability policy testing**  
  1 comment. This is a prerequisite for the capability-policy epic; the ask is to support multiple user roles on localhost so that per-user auth flows can be tested manually.

- **[#5315](https://github.com/nearai/ironclaw/issues/5315) — Daily failure taxonomy**  
  0 comments but notable as it aggregates 8 non-passing ClawBench tasks, all failing inside the benchmark adapter. This is an internal tracker rather than a community discussion, but it signals a recurring infrastructure issue.

PRs with high visibility (based on size/risk labels) include [#5313](https://github.com/nearai/ironclaw/pull/5313) (XL, storage stress harness) and [#5305](https://github.com/nearai/ironclaw/pull/5305) (XL, porting legacy tests). Both are core-contributor driven and have not yet received external feedback.

## 5. Bugs & Stability
Several bugs were reported today, ranked by severity:

- **High severity** — [#5283](https://github.com/nearai/ironclaw/issues/5283): “Approve & always allow” not persisted for `nearai.web_search`. A fix is proposed in [#5309](https://github.com/nearai/ironclaw/pull/5309) (still open). This directly impacts user trust in tool permissions.

- **Medium severity** — [#5302](https://github.com/nearai/ironclaw/issues/5302): Pending approval in one conversation blocks sending messages in other conversations until manual refresh. No dedicated fix PR yet, but related to the approval-state management work underway (e.g., [#5306](https://github.com/nearai/ironclaw/pull/5306)).

- **Medium severity** — [#5316](https://github.com/nearai/ironclaw/issues/5316): Gmail extension discovery/install is inconsistent — sometimes the agent reports no email extension, sometimes it works. Likely a race condition in extension registry lookup.

- **Infrastructure** — [#5315](https://github.com/nearai/ironclaw/issues/5315) profiles 8 failing ClawBench tasks, all due to the benchmark adapter itself (not the agent), which is a CI/stability concern rather than a product bug.

- **Low severity** — Several PRs merge fixes for UI glitches: run failure drifting ([#5299](https://github.com/nearai/ironclaw/pull/5299)), stale gate rows ([#5297](https://github.com/nearai/ironclaw/pull/5297)), and tool input validation errors ([#5310](https://github.com/nearai/ironclaw/pull/5310)).

All high- and medium-severity bugs have open fix PRs or are directly addressed by ongoing work in the capability-policy and approval-resume areas.

## 6. Feature Requests & Roadmap Signals
No explicit user-requested features were filed today, but several ongoing epics and PRs signal near-term roadmap direction:

- **Capability policy with per-user auth** ([#5261](https://github.com/nearai/ironclaw/issues/5261)) — The single epic for shipping admin-shared tools and skills with per-user auth on the Reborn stack. Sub-issue [#5272](https://github.com/nearai/ironclaw/issues/5272) (REST-created local users) is a gating item. Likely to land in the next minor release once testing infrastructure is ready.

- **Storage stress and multi-backend support** — PR [#5313](https://github.com/nearai/ironclaw/pull/5313) adds a stress harness for filesystem, libSQL, and Postgres backends. This suggests upcoming work on resource governor tuning and possibly new storage profiles (e.g., [#5259](https://github.com/nearai/ironclaw/pull/5259) – hosted single-tenant volume profile).

- **Default Reborn CLI** ([#5312](https://github.com/nearai/ironclaw/pull/5312)) — Bundling all major features into a single binary signals that Reborn is approaching production readiness, with the legacy engine expected to be phased out.

- **Porting legacy tests** ([#5305](https://github.com/nearai/ironclaw/pull/5305)) indicates the team is serious about achieving full parity before cutting over to Reborn.

Predictions for the next version (0.30.x or 1.0 preview): capability policy with per-user auth, one-binary Reborn CLI, and resolution of the WebUI v2 approval-persistence bugs.

## 7. User Feedback Summary
The most vocal user today is **sunglow666**, who filed three bugs:

1. **Frustration with broken “always allow” checkbox** ([#5283](https://github.com/nearai/ironclaw/issues/5283)) — “I checked ‘Approve & always allow’ but it still asks every time.” This undermines confidence in the permission system.

2. **Inconsistent Gmail extension** ([#5316](https://github.com/nearai/ironclaw/issues/5316)) — “Same prompt sometimes says no email extension, sometimes works.” Suggests interleaving of extension discovery with conversation state.

3. **Message sending blocked by unresolved approval in another conversation** ([#5302](https://github.com/nearai/ironclaw/issues/5302)) — “I switched to another conversation, typed a message, pressed Enter, nothing happens until I refresh.” This is a significant workflow blocker for multi-tab users.

Overall user sentiment appears cautiously negative on the Reborn WebUI v2’s reliability regarding tool permissions and multi-conversation management. However, the rapid response from core contributors (multiple fix PRs already open) suggests the team is responsive to these pain points.

## 8. Backlog Watch
No issues or PRs are languishing for an extended period; the oldest items in today’s update are from 2026-06-25 (yesterday). However, attention should be paid to:

- **[#5221](https://github.com/nearai/ironclaw/issues/5221) — Ironclaw harness backlog for deepseek-v4-flash**  
  Opened 2026-06-25, 0 comments. This is an internal tracking issue for hillclimb candidates. It has not been updated with any PR yet, so it may stall without owner follow-up.

- **[#5261](https://github.com/nearai/ironclaw/issues/5261) — Capability policy epic**  
  Although very recent, it has no assignee and no linked PRs beyond sub-issues. The epic is critical for the Reborn roadmap; maintainers should ensure progress is not blocked.

- **[#5272](https://github.com/nearai/ironclaw/issues/5272) — REST-created local users**  
  This sub-issue has a PR (likely to be created soon) but currently has no assignee. It is a prerequisite for manual testing of the capability policy, so any delay here blocks the entire epic.

All other open items from today are actively being worked on or have fix PRs attached. No P0 issues are abandoned.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-06-26

## Today's Overview
The project saw a burst of activity with **6 pull requests merged** and **1 issue closed** in the last 24 hours, though no new releases were published. The majority of changes focused on the **Cowork multi-agent collaboration feature** and a **major runtime upgrade** (OpenClaw v2026.6.1). The lone issue—a feature request for per-agent model binding and agent groups—was closed as stale, but its comments indicate strong user interest. Overall, the project appears healthy, with sustained engineering effort on both stability improvements and new capabilities.

---

## Releases
No new releases were tagged today. The runtime upgrade merged via PR #2209 will likely form the basis of a future release.

---

## Project Progress — Merged Pull Requests (6 items)

### Runtime & Build
- **[PR #2209](netease-youdao/LobsterAI PR #2209)** *(area: multiple)* — **Major runtime upgrade** from OpenClaw v2026.4.14 to v2026.6.1. Includes required patches, plugin upgrades, build script updates, Cowork integration fixes, and app version bump to 2026.6.26.
- **[PR #2211](netease-youdao/LobsterAI PR #2211)** *(area: main)* — Sorted test imports in `finalUpgradePatchDecisions.test.ts` to comply with linting rules.

### Cowork Multi-Agent Collaboration
- **[PR #2207](netease-youdao/LobsterAI PR #2207)** *(area: renderer, main)* — **Stabilized subagent progress tracking**. Now derives progress from local `subagent_runs` table instead of model-authored announce text. Fixes stale progress display (e.g., showing 3/5 when actual state is 5/5). Preserves failed spawn rows while appending retry replacements.
- **[PR #2208](netease-youdao/LobsterAI PR #2208)** *(area: renderer, main)* — **Freezed terminal subagent duration** in sidebar. Persists `endedAt` for finished subagents and uses `endedAt - createdAt` to avoid drifting durations. Running subagents continue to update against current time.

### Artifacts & UI
- **[PR #2210](netease-youdao/LobsterAI PR #2210)** *(area: renderer, artifacts)* — **Prevents Mermaid error SVG leaking**. Validates diagram content with `mermaid.parse()` before rendering; syntax errors now show the controlled artifact error UI instead of raw error SVGs. Also cleans up Mermaid-generated DOM nodes after render/error.
- **[PR #1459](netease-youdao/LobsterAI PR #1459)** *(area: skills)* — **Skill tooltip on hover**. New `SkillTooltip.tsx` component displays full name, official identifier, and complete description in a rich tooltip. Smart four-direction positioning with 300ms delay and two-stage height measurement. Resolves long-standing complaint about truncated skill descriptions.

---

## Community Hot Topics
The only issue updated today is the most active:
- **[Issue #1462](netease-youdao/LobsterAI Issue #1462)** (CLOSED, 3 comments) — User requested **per-agent model binding** and a **multi-agent group mode** with a manager agent that can dynamically dispatch sub-agents. The user acknowledged that the current multi-instance per-channel feature (v4.3) is useful, but expressed a desire for finer-grained control. They also compared LobsterAI favorably to Alibaba’s Hiclaw, noting better interaction design. The issue was closed by the stale bot, which may frustrate users who still want this functionality.

---

## Bugs & Stability
All bugs addressed today were fixed via merged PRs:

| Bug | Severity | Fixed By |
|-----|----------|----------|
| Mermaid syntax errors producing raw, uncontrolled SVG error artifacts | Medium (UI/UX) | [#2210](netease-youdao/LobsterAI PR #2210) |
| Cowork subagent progress showing stale/incorrect values (e.g., 3/5 instead of 5/5) | High (functional) | [#2207](netease-youdao/LobsterAI PR #2207) |
| Terminal subagent durations in sidebar continuing to tick after completion | Medium (UI) | [#2208](netease-youdao/LobsterAI PR #2208) |

No new critical crashes or regressions were reported. The Cowork progress fix (#2207) addresses a specific accuracy issue that could affect user trust in multi-agent workflows.

---

## Feature Requests & Roadmap Signals
The strongest signal comes from **Issue #1462**:
- **Per-agent model binding** — users want each agent to use a separate LLM model.
- **Agent group/room mode** — a manager agent that can orchestrate sub-agents on demand.

These requests align with the ongoing Cowork development seen in today’s PRs (#2207, #2208), suggesting the team is actively improving the multi-agent foundation. It is plausible that the **next minor release** (e.g., v4.4 or v2026.7) will introduce a higher-level orchestration layer, building on the subagent progress and duration fixes.

Additionally, the **skill tooltip feature** (#1459) addresses a long-standing UX pain point and indicates attention to polish.

---

## User Feedback Summary
- **Positive**: User in #1462 praised the v4.3 multi-instance per-channel feature as “very useful” and stated that LobsterAI’s interaction design is superior to Alibaba’s Hiclaw.
- **Gap**: The same user expressed a clear need for **per-agent model configuration** and **dynamic multi-agent coordination**. The issue was closed without resolution, which may lead to dissatisfaction if not addressed in a future release.
- **General**: The volume of Cowork-related PRs suggests that multi-agent collaboration is the area users interact with most and where improvements are most demanded.

---

## Backlog Watch
- **[Issue #1462](netease-youdao/LobsterAI Issue #1462)** — Closed as stale today despite having 3 comments and representing a well-articulated feature request. The maintainers may want to **reopen or track** the underlying ideas (per-agent model binding, agent groups) on a roadmap or an enhancement epic. The user’s comparison with Hiclaw adds competitive context. No other long-unanswered issues were identified in today’s data.

*Note: The stale bot closure may have been automatic based on inactivity since April 4; manual review could prevent valuable feedback from being lost.*

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

# CoPaw Project Digest — 2026-06-26

## 1. Today's Overview

Project activity remains very high, with **18 issues** and **50 pull requests** updated in the last 24 hours. A total of **24 PRs were merged or closed**, alongside **1 new release** — the early beta of QwenPaw v2.0.0-beta.1. Community engagement is strong, with users reporting bugs across desktop, channel integrations (WeCom, Feishu, DingTalk), and agent runtime (DeepSeek thinking, silent cron). The maintainers are actively triaging and pushing fixes, especially around the v2.0 migration and plugin ecosystem, though several open high-severity bugs still lack resolved PRs. Overall, the project is in a rapid development and stabilization phase, with many contributions from first-time contributors.

## 2. Releases

**v2.0.0-beta.1** — [Release Page](https://github.com/agentscope-ai/QwenPaw/releases/tag/v2.0.0-beta.1)  
⚠️ **Early beta** for QwenPaw 2.0.0. Contains breaking changes and instability; intended for developers and early adopters only, **not recommended for production**.

- **Changes:** `refactor: migrate agent` — foundational refactor of the agent subsystem for the 2.0 architecture.
- **Migration notes:** Users upgrading from 1.x should expect incompatibilities, especially with plugins and agent configurations. The release includes a pre-release checklist issue (#5571) for installation verification across platforms.

## 3. Project Progress

Today **24 pull requests were merged or closed**. Highlights:

- **#5340** (merged) — `fix(model_factory): switch formatter drop detection from blacklist to whitelist`  
  Fixes a regression where interrupted agent generations lost messages on next turn.
- **#5559** (merged) — `Feat/session switch performance`  
  Improves session switching latency, addressing performance regressions reported in recent builds.
- **#5560** (merged) — `fix(channels): process WeCom media-only messages`  
  Fixes a bug where file/image messages sent via WeCom were buffered and never delivered to the agent (fixes #5554).
- **#5536** (under review) — `fix: kill orphaned Chrome renderer processes on browser stop`  
  Addresses memory leaks from Playwright subprocesses accumulating over repeated start/stop cycles.
- **#5549** (open) — `fix(providers): sanitize nullable tool schemas`  
  Fixes 400 errors on OpenAI-compatible endpoints that reject JSON Schema `anyOf` with `null` types.
- **#5570** (open) — `fix(desktop): stop plugin dep install storm and orphaned backends`  
  Direct fix for #5550, adding locks to prevent fork-bomb-like `pip install` loops.

Other merged PRs include plugin fixes for QwenPaw 2.0 (#5568), pet plugin fixes (#5565), and heartbeat timeout configurability (#5557).

## 4. Community Hot Topics

The most active discussions today center around critical stability issues and feature requests:

- **[#5550] Remote SSH plugin dependency install loop + orphaned backends** (4 comments)  
  A severe bug affecting macOS Desktop users: `pip install` storms and accumulating backend processes can exhaust memory. Fix PR #5570 is open.  
  [Link](https://github.com/agentscope-ai/QwenPaw/issues/5550)

- **[#5262] Disabled built-in skills re-enable after upgrade** (12 comments, closed)  
  A long-standing user frustration: skills like `docx`, `xlsx` reset to enabled on each update. The issue was closed, but no explicit fix PR is linked — likely resolved in v2.0 migration.  
  [Link](https://github.com/agentscope-ai/QwenPaw/issues/5262)

- **[#5379] Internal Server Error on fresh Python install** (7 comments)  
  Users unable to start QwenPaw after `pip install` due to `get_remote_addr(transport)` failure. No fix PR yet.  
  [Link](https://github.com/agentscope-ai/QwenPaw/issues/5379)

- **[#5563] Feature request: aggregate multi-step agent responses** (3 comments)  
  Users want message bundling instead of receiving 10+ cards per task, to reduce chat spam.  
  [Link](https://github.com/agentscope-ai/QwenPaw/issues/5563)

- **[#5328] Agent freezes during DeepSeek thinking** (3 comments)  
  Reproducible across Web, Tauri, and Console — agent gets stuck in `thinking` and requires manual intervention.  
  [Link](https://github.com/agentscope-ai/QwenPaw/issues/5328)

Underlying needs: **installation reliability**, **channel maturity** (especially WeCom, Feishu, DingTalk), and **agent execution robustness** (timeouts, message aggregation).

## 5. Bugs & Stability

Reported bugs today ranked by severity:

| Severity | Issue | Summary | Fix PR exists? |
|----------|-------|---------|----------------|
| **Critical** | #5550 | Remote SSH plugin install loop + orphan backends (memory exhaustion) | ✅ #5570 |
| **Critical** | #5379 | Internal Server Error on `pip install` – cannot start | ❌ |
| **High** | #5328 | DeepSeek thinking freeze (multi-platform) | ❌ |
| **High** | #5573 | DeepSeek V4 400 errors: missing `reasoning_content` + nullable tool schemas | ✅ #5549 |
| **High** | #5554 | WeCom file messages not reaching agent | ✅ #5560 |
| **Medium** | #5566 | Cron task cannot be silent + `channels send` unreachable from background scripts | ❌ |
| **Medium** | #5561 | Feishu bot fails on long text replies (sends as file instead) | ❌ |
| **Low** | #5556 | `reme-ai` dependency not on PyPI (source install) | Closed, works around |
| **Low** | #5555 | General performance degradation ("getting stuck") – vague | ❌ |

Additionally, the v2.0.0-beta.1 release is flagged as unstable by design. Users are advised to stay on v1.1.x for production use.

## 6. Feature Requests & Roadmap Signals

Several feature requests with clear user demand were filed today:

- **Message aggregation** (#5563) — users want multi-step agent responses batched into a single message card. High impact on UX.
- **DingTalk @mention support** (#5564) — ability to `@` agents via CLI/API for multi-agent workflows.
- **Model auto-fallback** (#5572) — automatic switch to backup model on quota exhaustion or timeout. Common enterprise ask.
- **WeCom attachment-only send** (#5558) — users want to send files without needing accompanying text (e.g., CAD/screenshot analysis).
- **Computer use support** (#5551) — question about future plans for GUI automation. Not yet confirmed.

Ongoing requests: **Slack channel integration** (#5152, closed as a duplicate or superseded) — likely under consideration. The **GitHub Issue helper skill** (#5567) demonstrates community desire for better feedback tooling.

**Predictions for next release (v2.0.0 stable?):**  
The team appears focused on stabilizing the 2.0 refactor, fixing plugin compatibility (#5568), and addressing the DeepSeek provider issues (#5549). Likely to also include the heartbeat configurability (#5557) and session switch performance improvement (#5559). Message aggregation and model fallback may appear in a minor release after beta.

## 7. User Feedback Summary

Real user pain points from today’s issues:

- **"I have to manually disable skills every update"** — #5262 highlights configuration persistence failures.
- **"Agent freezes on thinking and I have to stop and resume"** — #5328, a cross-platform annoyance.
- **"WeCom can't handle file-only messages"** — #5554, #5558 show incomplete channel UX.
- **"Feishu bot fails on slightly long replies"** — #5561, users forced to send files instead.
- **"Cron tasks spam my DingTalk even when there's nothing to report"** — #5566, no silent execution mode.
- **"SSH plugin causes a fork bomb on macOS"** — #5550, memory exhaustion under desktop app.
- **"Installation gives Internal Server Error immediately"** — #5379, a blocker for new users.

Overall sentiment: users are engaged and willing to report bugs in detail, but there is clear dissatisfaction with regressions in recent versions and missing polish in channel integrations. On the positive side, many issues receive fast responses and are paired with fix PRs within hours.

## 8. Backlog Watch

Several important issues and PRs appear to be long-unanswered or stuck under review:

| Item | Author | Created | Status | Concern |
|------|--------|---------|--------|---------|
| **PR #4622** — DataPaw data-analysis plugin | EliasMei | 2026-05-22 | Under Review | No updates in over a month; plugin adds 12 BI skills. |
| **PR #4041** — Tauri tray behavior | wfeng007 | 2026-05-05 | Under Review | Stalled for nearly 2 months; foundational desktop UX. |
| **PR #5525** — Windows native sandbox | ustc-mkh | 2026-06-25 | Open | Awaiting review; important for sandboxed execution on Windows. |
| **PR #5515** — Enable latest chat beta UI | sanfran1068 | 2026-06-25 | Under Review | Blocked on dependency update; critical for frontend improvements. |
| **Issue #5379** — Internal Server Error on install | luo201227 | 2026-06-22 | Open | No maintainer response or fix PR after 4 days. |

These items represent a growing maintenance debt that could hinder contributor trust. The community would benefit from a triage update on these older contributions.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest – 2026-06-26

## 1. Today's Overview

ZeroClaw shipped **v0.8.2** with A2A agent discovery and enriched skills configuration, while the development pipeline remains highly active with **50 PRs updated** in the last 24 hours (44 open, 6 merged/closed). The project continues to focus on interoperability, security hardening, and operator experience. Five issues were updated (4 open, 1 closed), with two newly filed bugs—one blocking the multi‑agent skill‑install flow (S2 severity) and one affecting native tool‑calling with image markers. The release brings a net positive feature advance, but the high open‑PR count (44) and low merge throughput (6 today) suggest a growing review backlog that may slow feature velocity.

## 2. Releases

**v0.8.2** – published today (2026-06-26)  
[GitHub Release](https://github.com/zeroclaw-labs/zeroclaw/releases/tag/v0.8.2)

Key changes:
- **A2A agent discovery** – foundation for agent‑to‑agent interoperability.
- **Richer skills story** – user‑configured extra registries and typed slash‑command options.
- **Security posture** – sharper security controls across plugins, channels, and other components.

**Breaking changes:** None noted in the release notes.  
**Migration notes:** Users upgrading from v0.8.1 should verify their skills registries and slash‑command definitions, as the new typed‑options format may require minor config updates. No other migration steps are documented.

## 3. Project Progress

**Merged/closed PRs today (6 total):**  
Only the following merged/closed PR appears in the top‑20 list by comment count:

- `#8022 [CLOSED]` – [feat(web): themed, click‑to‑open config pickers + dashboard health fix‑modal + storage nav fix](https://github.com/zeroclaw-labs/zeroclaw/pull/8022)  
  *Enhancement. Improves Operator Console UX with dark‑theme‑friendly pickers and fixes storage navigation.*  
  (Expected to be part of the v0.8.2 release or a follow‑up patch.)

The remaining 5 merged/closed PRs (not shown in the top‑20 list) are unconfirmed but likely represent minor fixes or documentation updates.

**Features advanced (not yet merged):**
- **WASM component‑model plugin host** (#7928) – large PR, still open.
- **Two‑path onboard tree** (#8033) – LLM + deterministic flow over RPC/CLI.
- **In‑app upgrade from dashboard** (#8173) – detect, show notes, apply, restart.
- **Rotating log persistence** (#8307) – size/date/retention control.
- **Cost tracking from live gateway pricing** (#8233) – fill unpriced models dynamically.
- **Tool approvals routed to distinct approver channel** (#8231).
- **LAN peer discovery hints** (#8325) – MDNS multicast announce/discover.

## 4. Community Hot Topics

Most issues and PRs see modest discussion (1–2 comments). The most active items are:

- **Issue #8327** – [Native tool calling: [IMAGE:data:...] markers sent as plain text](https://github.com/zeroclaw-labs/zeroclaw/issues/8327)  
  *2 comments.* User `optman` reports that image‑data markers are treated as text, inflating token count with llama.cpp. The underlying need is proper `image_url`‑part serialisation for OpenAI‑compatible providers. This is a high‑impact bug for users relying on vision tools.

- **Issue #8334** – [[Bug]: `skills install`/`list`/`remove` target `data_dir`, which no multi‑agent runtime loads](https://github.com/zeroclaw-labs/zeroclaw/issues/8334)  
  *1 comment.* Labeled S2 severity. The “pull a skill and use it” flow is broken on multi‑agent setups. This directly affects usability of the new skills platform.

- **PR #8329** – [fix(runtime): forward narration emitted after a native tool call](https://github.com/zeroclaw-labs/zeroclaw/pull/8329)  
  *Undefined comment count but appears in top‑20.* Fixes streaming consumer latching `suppress_forwarding` prematurely. Essential for correct tool‑call response handling.

**Underlying needs:** Users are demanding reliable multi‑agent composition, proper handling of rich media (images), and seamless tool interaction. The v0.8.2 focus on A2A and skills resonates, but bugs like #8334 and #8327 erode trust in those features immediately after release.

## 5. Bugs & Stability

| Issue | Severity | Description | Fix PR exists? |
|-------|----------|-------------|----------------|
| [#8334](https://github.com/zeroclaw-labs/zeroclaw/issues/8334) | **S2 – degraded** | `skills install/list/remove` broken on multi‑agent runtimes (uses wrong `data_dir`). | Not yet |
| [#8327](https://github.com/zeroclaw-labs/zeroclaw/issues/8327) | Not specified | Native tool‑calling sends `[IMAGE:data:...]` as plain text instead of `image_url` parts, inflating token count. | No direct fix PR, but [#8329](https://github.com/zeroclaw-labs/zeroclaw/pull/8329) addresses related narration‑forwarding bug |
| [#8146](https://github.com/zeroclaw-labs/zeroclaw/pull/8146) (PR, open) | High (risk) | CLI one‑shot runs lose telemetry and token totals on exit. | Yes, PR open (not yet merged) |
| [#8330](https://github.com/zeroclaw-labs/zeroclaw/pull/8330) (PR, open) | Medium (risk) | `render_conversation` clones entire cache per frame; fixed by viewport‑only rendering. | Yes, PR open |
| [#8350](https://github.com/zeroclaw-labs/zeroclaw/pull/8350) (PR, open) | Low (perf, panic risk) | `strip_tags` regex recompiled every call, panics on bad pattern. | Yes, PR open |

**Observation:** Three bug‑fix PRs are open but none have been merged today. The S2 severity bug (#8334) currently has no fix in flight, which is concerning given it blocks a headline feature.

## 6. Feature Requests & Roadmap Signals

**New feature requests (today):**

- **Issue #8348** – [Skill CRUD hook/event for observing skill changes](https://github.com/zeroclaw-labs/zeroclaw/issues/8348)  
  User `chdeepexi` wants official hooks to notify external systems when skills are created, installed, updated, deleted, archived, or bundled. This is a natural extension of the skills platform and aligns with the project’s growing emphasis on agent‑to‑agent and external integration.

**Ongoing feature PRs (likely candidates for next minor release):**
- Skill‑dashboard write‑guard and shadowed‑by audit (#8082) – follow‑up to v0.8.2 skills UI.
- WASM plugin host (#7928) – major platform expansion, likely v0.9.0.
- In‑app upgrade (#8173) – user‑facing quality‑of‑life feature.
- Cost tracking from live pricing (#8233) – addresses silent zero‑cost recording.
- Tool approval routing (#8231) – closes a cross‑channel HITL gap.

**Prediction for v0.8.3 / v0.9.0:** The next release will likely include the skill CRUD hook, the rotating log mode, and one or two of the larger features (WASM host or onboard tree). The A2A and skills foundations are now released, so follow‑up bugs and polish will dominate the near term.

## 7. User Feedback Summary

**Pain points (real user reports):**

- “Image markers are sent as plain text, inflating token count” (#8327) – affects users of llama.cpp and vision tools; demands structured image‑url parts.
- “Skills install is broken in multi‑agent runtime” (#8334) – a user attempting the “pull a skill and use it” demo flow hit a dead end; the new skills feature is effectively broken for multi‑agent setups.
- “No hook to observe skill CRUD changes” (#8348) – external system integrators need event notifications.

**Satisfaction signals:**

- The v0.8.2 release note emphasises A2A and skills enrichment, which directly addresses two long‑requested areas. The community likely welcomes the direction, even if initial bugs hamper immediate use.
- Multiple PRs from different contributors (JordanTheJet, Nillth, Audacity88) indicate healthy external contribution.

**Overall tone:** Enthusiastic about new features but frustration with release‑day regressions. Users expect smoother quality after a major version bump.

## 8. Backlog Watch

**Long‑unanswered Issues/PRs needing maintainer attention:**

- **PR #7928** – [feat(wasi): initial WASM component‑model plugin host code](https://github.com/zeroclaw-labs/zeroclaw/pull/7928)  
  *Opened 2026-06-18, size XL, risk high.* This is a foundational piece for plugin extensibility but has not received maintainer review. It may block downstream work.

- **PR #8033** – [feat(onboard): two‑path onboard tree wired end‑to‑end](https://github.com/zeroclaw-labs/zeroclaw/pull/8033)  
  *Opened 2026-06-20, size XL, risk high.* Complex PR touching multiple crates; needs thorough review to avoid merge conflicts and regressions.

- **PR #8082** – [feat(skills): dashboard write‑guard, skipped‑audit, and shadowed_by](https://github.com/zeroclaw-labs/zeroclaw/pull/8082)  
  *Opened 2026-06-21.* Deferred follow‑ups from the skills dashboard. With v0.8.2 released, this PR should be prioritised to stabilise the skills UI.

- **Issue #8181** – [Tracker: v0.8.2 release‑support and non‑plugin queue](https://github.com/zeroclaw-labs/zeroclaw/issues/8181)  
  *Updated today.* With v0.8.2 now out, this tracker should be closed or migrated to a v0.8.3 tracker. 37 items remain open with no clear path.

**Recommendation:** The maintainers should triage the 44 open PRs, merge the simplest bug‑fix PRs first (e.g., #8330, #8350), and schedule a review session for the three large PRs (#7928, #8033, #8082) to unblock progress.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*