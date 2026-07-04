# OpenClaw Ecosystem Digest 2026-07-04

> Issues: 117 | PRs: 500 | Projects covered: 13 | Generated: 2026-07-04 09:06 UTC

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

# OpenClaw Project Digest — 2026-07-04

## Today’s Overview

Activity on the OpenClaw repository remains very high, with **117 issues** and **500 pull requests** updated in the past 24 hours. Of those, **17 issues were closed** and **88 PRs were merged or closed**, indicating steady progress. No new releases have been published; the project is in a dense development cycle. The maintenance team is actively processing contributions, though a significant number of open items (100 open issues, 412 open PRs) reflect the project's scale and community momentum. Priority levels (P1/P2) dominate the most active discussions, pointing to both critical bug fixes and high-impact feature requests.

## Releases

None in the past 24 hours.

## Project Progress

In the last day, **88 pull requests were merged or closed**, spanning agents, channels (Discord, Feishu, LINE, WhatsApp), UI, security, and plugin consolidation. Notable merged/closed PRs from the top comment list:

- **#97790** (CLOSED) – Estimates DeepSeek spend when API reports zero cost, fixing a zero-spend display issue.
- **#99850** (CLOSED) – Major consolidation of duplicated plugin state and doctor migration plumbing onto SDK seams (closes #99841).
- **#99893** (CLOSED) – Hides internal Code Mode wait progress in Discord channel.
- **#99821** (CLOSED) – Shares native threads across Codex clients, enabling thread visibility between OpenClaw and native Codex apps (closes #99781).
- **#94012** (CLOSED) – Routes canonical provider models through ClawRouter, adding generic plugin hooks for credential-scoped model transport.

These merges significantly reduce code duplication, improve plugin architecture, and extend interop with Codex and credential-aware routing.

## Community Hot Topics

**Most active issues** (by comment count):

- **#22438** (17 comments) – *Tiered bootstrap file loading* – Users want to reduce LLM token waste by loading bootstrap files only as needed.
- **#22676** (17 comments) – *Signal daemon race condition on SIGUSR1 restart* – Produces orphaned processes and send failures; a high-severity P1 bug.
- **#48003** (15 comments, 3 👍) – *Steer mode does not inject messages mid-turn* – A session-state bug that breaks real-time steering.
- **#29387** (14 comments, 5 👍) – *Bootstrap files silently ignored in agentDir* – Long-standing configuration inconsistency between workspace and per-agent directories.
- **#73148** (14 comments, 3 👍) – *Image tool opaque failure when sharp is missing* – Users frustrated by missing dependency error messages.
- **#22358** (12 comments, 1 👍) – *Post-subagent completion extension hook* – Demand for workflow orchestration and trajectory logging.
- **#86215** (10 comments, 1 👍) – *Codex OAuth refresh failure wedges agent* – P1 impacting availability and alerting.

**Underlying needs**: Community discussions reveal a strong desire for better context management (tiered loading, bootstrap isolation), reliable messaging (steer mode, signal race), improved error messaging (image tool), and configurable extensions (hooks, OAuth resiliency).

## Bugs & Stability

The following severe bugs were active in the last 24 hours, all rated P1 by the project:

- **#22676** (P1) – Signal daemon race condition during SIGUSR1 restarts causes orphaned processes and send failures. Linked PR open.
- **#48003** (P1) – Steer mode fails to inject messages mid-turn for main sessions. Identified root cause in `KeyedAsyncQueue`. Linked PR open.
- **#29387** (P1) – Bootstrap files in per-agent directories are silently ignored. Multiple security and session-state impacts. Linked PR open.
- **#86215** (P1) – Codex OAuth refresh failures can wedge an agent for hours. Needs clearer alerting and profile rotation. Linked PR open.
- **#31331** (P1) – Docker + Sandbox cannot access workspace when using Docker-outside-of-Docker. Breaks fundamental container usage.
- **#58514** (P1) – Google Chat space/group messages silently ignored. DMs work, but group messages are lost.
- **#25574** (P1) – Config warnings logged repeatedly on every reload, spamming error logs with thousands of duplicates.
- **#94846** (P2) – Cron isolated agentTurn skips delivery after a recovered early tool error is misclassified as fatal.

Each of these bugs has a linked PR or is awaiting maintainer review. The density of P1 issues indicates the project is actively squashing regressions while adding new features.

## Feature Requests & Roadmap Signals

New feature requests with significant community support (by reactions or cross-linking):

- **#20786** (6 👍) – Telegram Business Bot support – high demand for business messaging integration.
- **#29387** (5 👍) – AgentDir bootstrap file loading – a configuration fix that acts like a feature request.
- **#7722** (4 👍) – Filesystem sandboxing config – user safety and security requirement.
- **#31331** (4 👍) – Docker+Sandbox workspace access – blocks container deployment.
- **#10118** (4 👍) – TUI support for Shift+Enter for newlines – small UX improvement with high approval.
- **#14438** (4 👍) – Plugin hot-reload without container restart – developer experience essential for plugin authors.
- **#9986** (0 👍 but high discussion) – Trigger model fallback on context length exceeded – missing feature that would reduce LLM errors.
- **#13615** (2 👍) – Rate limiting and throttling for external API calls – cost control and reliability.
- **#17810** (Request) – Configurable context file elevation for persona files (e.g., SOUL.md) – community sentiment: agents need identity priority.

**Prediction for next version**: The number of P1 bugs with linked PRs suggests a stabilization release. Given community demand, **tiered bootstrap loading** (#22438), **Steer mode fix** (#48003), and **OAuth resilience** (#86215) are likely candidates. Also, **multi-file Discord attachments** (#99515) and **pre-auth access requests** (#89569) are nearing maintainer approval and could be merged soon.

## User Feedback Summary

Real pain points from today’s data:

- **Frustration with silent failures**: Users report messages being ignored (Google Chat), steer mode not injecting, and bootstrap files having no effect without warning (#58514, #48003, #29387).  
- **Context window waste**: Users want to control what gets loaded into every session to save tokens (#22438, #17810).  
- **Deployment friction**: Docker+Sandbox workspace access issues (#31331) and cross-device move errors on plugin install (#99896) hinder container adoption.  
- **Debugging difficulty**: Opaque errors like “Failed to optimize image” (#73148) and missing error handling in JSON parsing (#99894) make troubleshooting hard.  
- **Missing integrations**: Business bots (Telegram), custom emoji support, and full channel parity (Mattermost, LINE) are repeatedly requested.  

**Satisfaction signals**: High PR throughput and maintainer responsiveness (many PRs moved to “ready for maintainer look”) indicate an active and engaged team. The plugin duplication consolidation (#99850, #99901) shows care for code quality.

## Backlog Watch

Issues and PRs that have been open for a long time or need urgent maintainer attention:

- **#7722** (Feb 3, 9 comments, 4 👍) – Filesystem sandboxing config. Open since early February, labeled `needs-maintainer-review` and `needs-product-decision`.  
- **#7639** (not in top 50 but referenced) – Telegram default allowlist mode. Open since Feb 3.  
- **#10687** (Feb 6, 9 comments) – Dynamic model discovery for OpenRouter. Stalled on product decision.  
- **#12855** (Feb 9, 7 comments) – Built-in auto-update workflow. Needs product decision.  
- **#7722** and **#9986** (Feb 5) – Model fallback on context exceeded. 5 comments, needs product decision.  
- **#22358** (Feb 21, 12 comments) – Post-subagent completion hook. Has linked PR but still needs maintainer review.  
- **#26370** (Feb 25, 5 comments) – Per-agent cron isolation in multi-user deployments. Multi-label high-priority but no movement.  

Many of these issues are tagged `clawsweeper:needs-product-decision` – they require architectural guidance from the core team. The community is actively discussing solutions, but decisions appear to be a bottleneck.

---

*Data compiled from GitHub activity for `openclaw/openclaw` on 2026-07-04. All issue/PR links: https://github.com/openclaw/openclaw/issues and https://github.com/openclaw/openclaw/pulls.*

---

## Cross-Ecosystem Comparison

# Cross-Project Comparison Report: Personal AI Agent Open-Source Ecosystem

**Date:** 2026-07-04  
**Scope:** 12 projects across the AI agent and personal assistant landscape  
**Audience:** Technical decision-makers, developers, and ecosystem analysts

---

## 1. Ecosystem Overview

The personal AI agent open-source ecosystem is undergoing a **rapid consolidation phase**, where foundational plumbing (credential routing, context management, plugin systems) is being standardized while differentiation shifts toward channel integrations, deployment models, and reasoning capabilities. Seven of twelve projects show high or moderate activity, with **OpenClaw, Hermes Agent, IronClaw, and ZeroClaw** leading in commit velocity and community engagement. A clear divide is emerging between **general-purpose agent frameworks** (OpenClaw, Hermes, IronClaw) and **specialized assistants** (CoPaw, LobsterAI, NanoClaw), with the latter focusing on polished desktop/WebUI experiences and goal-oriented workflows. Across all active projects, **context management, multi-agent orchestration, and channel reliability** emerge as the most urgent shared challenges, while Mobile/Windows support and structured reasoning (goal modes, SOPs) represent the next competitive frontier.

---

## 2. Activity Comparison

| Project | Issues Updated (24h) | PRs Updated (24h) | Release This Week | Health Score |
|---|---|---|---|---|
| **OpenClaw** | 117 | 500 | No | 🔴 **Intense** – High throughput but 412 open PRs indicate bottleneck |
| **Hermes Agent** | 12 | 50 | No | 🟡 **High** – Critical bug (#58168) threatens session stability |
| **IronClaw** | 8 | 50 | No | 🟡 **High** – CI failures and E2E flakiness need attention |
| **NanoClaw** | 0 | 20 | No | 🟢 **Moderate** – Steady fixes, security hardening |
| **ZeroClaw** | 13 | 50 | No | 🟢 **High** – Fixed 3 S1 bugs, large feature branches active |
| **CoPaw** | 7 | 17 | Yes (2026.7.3) | 🟢 **High** – Healthy mix of features and community contributions |
| **LobsterAI** | 1 | 5 | Yes (2026.7.3) | 🟢 **Moderate** – Shipping cowork features, stale backlog concerns |
| **NanoBot** | 3 | 17 | No | 🟢 **Moderate** – Active fixes, MCP reliability gaps |
| **PicoClaw** | 2 | 12 | No | 🟡 **Moderate** – Stale bugs on Android and WhatsApp |
| **NullClaw** | 1 | 0 | No | 🔴 **Low** – Single open issue, zero PR activity |
| **TinyClaw** | 0 | 0 | No | ⚪ **Inactive** |
| **Moltis** | 0 | 0 | No | ⚪ **Inactive** |
| **ZeptoClaw** | 0 | 0 | No | ⚪ **Inactive** |

**Health Score Key:** 🔴 Critical issues or stalled / 🟡 Active with notable risks / 🟢 Healthy / ⚪ No activity

---

## 3. OpenClaw's Position

### Advantages vs. Peers
- **Community size:** With 117 issues and 500 PRs updated in 24 hours, OpenClaw commands the **largest contributor base** in the ecosystem by a factor of ~10x over the next most active projects. This creates unmatched plugin availability and channel support (Discord, Feishu, LINE, WhatsApp, Telegram, Google Chat).
- **Architectural maturity:** The `ClawRouter` for credential-scoped model routing (PR #94012) and plugin state consolidation onto SDK seams (PR #99850) demonstrate a **systematic approach to technical debt** that smaller projects lack.
- **Channel breadth:** Only OpenClaw and ZeroClaw support the full range of consumer messaging platforms; others focus on 2-3 primary channels.

### Technical Approach Differences
- **Plugin system:** OpenClaw’s approach to plugin state migration and doctor plumbing (PR #99850) is more rigorous than NanoClaw’s optional plugin system (#4396) or Hermes’s ad-hoc extensions.
- **Context management:** The tiered bootstrap file loading proposal (#22438) signals an architectural investment in LLM token optimization that peer projects have not yet formalized.
- **Codex interop:** OpenClaw is unique in providing native thread sharing with Codex clients (#99821), creating a **cross-platform continuity** that Hermes and IronClaw do not offer.

### Community Size Comparison
| Metric | OpenClaw | Next Largest (Hermes / IronClaw) | Ratio |
|---|---|---|---|
| Daily PR updates | 500 | 50 | 10:1 |
| Daily issue updates | 117 | 12 | 10:1 |
| Open PR backlog | 412 | ~15-30 | 14-27:1 |
| Distinct active PR authors (est.) | 30+ | 5-10 | 3-5:1 |

### Risk
The **412 open PRs** and 100 open issues represent a review bottleneck that could slow response time for critical fixes. Hermes’s ability to merge 18 PRs/day with a smaller team suggests OpenClaw may benefit from automated triage or additional maintainers.

---

## 4. Shared Technical Focus Areas

The following requirements appear across **3+ projects**, indicating ecosystem-wide pain points:

| Focus Area | Affected Projects | Specific Need |
|---|---|---|
| **Context Management** | OpenClaw, Hermes, IronClaw, NanoClaw, CoPaw, ZeroClaw | Tiered bootstrap loading (#22438), protected compression anchors (#5710), ticket-based context (#78155), cron memory suppression (#8695) |
| **Multi-Agent Orchestration** | OpenClaw, Hermes, PicoClaw, CoPaw, ZeroClaw | Sub-agent completion hooks (#22358), goal-mode cowork (LobsterAI), agent collaboration bus (#2937), per-agent overrides (#3225) |
| **MCP Reliability** | NanoBot, PicoClaw, ZeroClaw, Hermes | Crash on tool call errors (#4652), reconnect failures (#4302), missing tools in TUI (#8193) |
| **Channel Parity** | OpenClaw, Hermes, NanoClaw, PicoClaw, ZeroClaw | WhatsApp routing bugs (#37906), Google Chat silent drops (#58514), Signal DM issues (#2694), Telegram idle timeout (#972) |
| **Credential & OAuth** | Hermes, IronClaw, NanoClaw, OpenClaw | Auth store encoding on Windows (#58158), OAuth refresh wedging (#86215), credential obligation alignment (#5512), caller context preservation (#2611) |
| **Mobile & Windows Support** | NanoBot, Hermes, PicoClaw, ZeroClaw | Mobile WebUI clipping (#4693), Windows sandbox (#5525), Android service crash (#3182), MSYS path normalization (#57696) |

### Pattern Recognition
**Context management** is the #1 shared pain point across the ecosystem. Seven of twelve active projects have open issues or PRs addressing how context is loaded, compressed, or preserved. This reflects a maturing understanding that LLM token budgets require **explicit developer control** rather than relying on generic conversation history.

---

## 5. Differentiation Analysis

| Project | Primary Differentiator | Target User | Architecture |
|---|---|---|---|
| **OpenClaw** | **All-purpose agent OS** – broadest channel coverage, largest plugin market, Codex interop | Power users, multi-platform teams | Monorepo, Python/TypeScript, middleware-based routing |
| **Hermes Agent** | **Enterprise reasoning** – MiMo/multi-modal support, profiled workers, SSH integration | Enterprise teams, research labs | Modular, credential-isolated workers |
| **IronClaw** | **Rust-native agent framework** – WASM credential system, Reborn architecture, Slack ingress | Developers building custom agents | Rust, WASM-powered, manifest-driven routing |
| **NanoClaw** | **Safety-first assistant** – session approval guards, path traversal prevention, container security | Security-conscious users, compliance teams | Go-based, sandbox-by-default |
| **ZeroClaw** | **Low-code agent platform** – visual SOP authoring, Inkbox channel, OpenAI-compat gateway | Non-developers, citizen developers | Rust, visual tools, .ignore mechanisms |
| **CoPaw (QwenPaw)** | **Desktop-first UX** – Tauri desktop, memory reranker, LLM fallback, Azure Bot channel | Desktop users, Chinese market | TypeScript/React, Qwen-integrated |
| **LobsterAI** | **Goal-oriented Cowork** – sub-agent artifacts, goal mode, service deployment | Collaboration-focused teams | TypeScript, OpenClaw RPC integration |
| **NanoBot** | **Lightweight embeddable agent** – CLI first, plugin controls, heartbeats | Developers embedding assistants | Python, simple gateway model |
| **PicoClaw** | **Edge agent** – DeltaChat, Simplex, WhatsApp focus | Privacy-focused users, IoT | Go, minimal dependencies |

### Key Strategic Takeaways
- **OpenClaw** competes on breadth and ecosystem lock-in; its challenge is maintaining review velocity.
- **IronClaw** and **ZeroClaw** are architecturally the most forward-looking (Rust, WASM, manifest-driven), suggesting they may set future standards.
- **CoPaw** and **LobsterAI** prioritize user experience over plumbing depth, targeting the "AI assistant" rather than "agent framework" market.
- **NullClaw, TinyClaw, Moltis, ZeptoClaw** are effectively dormant; users should not rely on them for active development.

---

## 6. Community Momentum & Maturity

### Tier 1: High Velocity, Active Daily (Intense Development)
| Project | Velocity Signal | Maturity Risk |
|---|---|---|
| **OpenClaw** | 500 PR updates/day, 88 merges | Review bottleneck; 412 open PRs |
| **Hermes Agent** | 50 PR updates/day, 18 merges | Critical session-breaking bug (#58168) |
| **IronClaw** | 50 PR updates/day, 23 merges | CI flakiness, nightly E2E failures |
| **ZeroClaw** | 50 PR updates/day, 6 merges | Large feature branches may slow releases |

### Tier 2: Steady Growth, Weekly Releases
| Project | Velocity Signal | Maturity Risk |
|---|---|---|
| **NanoClaw** | 20 PR updates/day, security hardening | 17 open PRs, some since May |
| **CoPaw** | 17 PR updates/day, healthy community | SDK architecture blocking multi-agent (#5767) |

### Tier 3: Focused Iteration
| Project | Velocity Signal | Maturity Risk |
|---|---|---|
| **LobsterAI** | 5 PR updates/day, shipping features | Two stale items from April |
| **PicoClaw** | 12 PR updates/day, channel expansion | Android bug unaddressed |
| **NanoBot** | 17 PR updates/day, MCP fixes | Gateway crash (#4302) unresolved |

### Tier 4: Low / No Activity
| Project | Status | Recommendation |
|---|---|---|
| **NullClaw** | 1 open issue, 0 PRs | Consider archived status |
| **TinyClaw** | No activity | Likely abandoned |
| **Moltis** | No activity | Likely abandoned |
| **ZeptoClaw** | No activity | Likely abandoned |

---

## 7. Trend Signals

### Observed Industry Trends (From Community Feedback)

1. **Context is the new memory** – Across all active projects, developers are realizing that generic conversation history management is insufficient. The shift toward **tiered bootstrap loading** (OpenClaw #22438), **protected compression anchors** (CoPaw #5710), and **per-agent context overrides** (PicoClaw #3225) signals that context management is becoming a **first-class architectural concern**, not a UX polish.

2. **Multi-agent orchestration is the next frontier** – Five projects (OpenClaw, Hermes, PicoClaw, CoPaw, ZeroClaw) have active work on sub-agent completion hooks, goal modes, agent collaboration buses, or per-agent model overrides. The ecosystem is moving beyond "one agent fits all" toward **composable agent teams** with specialized roles and routing.

3. **MCP reliability is table stakes** – The fact that three projects (NanoBot, PicoClaw, ZeroClaw) shipped fixes for MCP tool crash bugs in the same 24-hour window indicates that the **Model Context Protocol is entering production reality**, and crashes are no longer acceptable. Expect MCP compliance to become a selection criterion.

4. **Channel parity is a competitive differentiator** – WhatsApp, Signal, Google Chat, and Telegram bugs persist across multiple projects. Users expect **full channel parity** before considering agent deployment. The pattern of "DMs work but group messages fail" (OpenClaw #58514, Hermes #37906) is a recurring complaint.

5. **Credential security is getting serious** – Hermes’s auth store encoding fix (#58158), IronClaw’s WASM credential alignment (#5512), and NanoClaw’s caller context preservation (#2611) all address **credential exposure risks**. This suggests production deployments are demanding better isolation and auditability.

6. **Mobile and desktop UX is expanding** – CoPaw’s Tauri desktop pipeline (#5734), NanoBot’s mobile WebUI fix (#4694), and PicoClaw’s Android bug (#3182) show that the ecosystem is **closing the gap** from CLI-first to consumer-friendly interfaces.

7. **Structured reasoning is emerging** – Goal modes (LobsterAI), SOP authoring (ZeroClaw), and hierarchical completion hooks (OpenClaw) indicate a shift from free-form conversation toward **structured, goal-directed agent behaviors** – a pattern that echoes the transition from chatbots to AI agents.

### Value for AI Agent Developers

- **If you need maximum channel coverage and plugin ecosystem** → **OpenClaw** remains the default choice, but prepare for PR backlog friction.
- **If you value architecture quality and future-proofing** → **IronClaw** (Rust/WASM) or **ZeroClaw** (manifest-driven) offer cleaner foundations.
- **If you target desktop users or Chinese market** → **CoPaw** delivers the best UX, especially with Qwen integration.
- **If security and credential isolation are paramount** → **NanoClaw** and **IronClaw** lead in this dimension.
- **Avoid NullClaw, TinyClaw, Moltis, and ZeptoClaw** – their inactivity makes them unsuitable for new projects.

**Recommendation:** The ecosystem is healthy but fragmented. For long-term investment, bet on Rust-based architectures (IronClaw, ZeroClaw) or the community-momentum leader (OpenClaw). For immediate deployment with good UX, CoPaw and LobsterAI are the most polished options.

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot Project Digest — 2026-07-04

## 1. Today’s Overview
NanoBot shows moderate development activity with **17 pull requests** updated in the last 24 hours (5 merged/closed) and **3 open issues** receiving attention. The project is actively addressing stability concerns around MCP tool exception handling and gateway reconnection, while also polishing the WebUI for mobile users and expanding channel integrations (Mattermost, OpenCode, Anthropic OAuth). No new releases were published, but several fixes and enhancements are nearing completion. Overall project health is good, with a healthy mix of bug fixes, feature additions, and community-driven improvements.

## 2. Releases
None. No new versions were published in the last 24 hours.

## 3. Project Progress (Merged/Closed PRs Today)
Five PRs were merged or closed, advancing both stability and feature work:

- **[#4692] fix(config): serialize model presets as camelCase**  
  Closed. Harmonizes the `model_presets` field with documentation’s `modelPresets` while accepting both forms.  
  → [PR #4692](https://github.com/HKUDS/nanobot/pull/4692)

- **[#4632] feat(providers): add Anthropic OAuth**  
  Closed. Adds support for Claude Code tokens via `claude setup-token`, enabling OAuth-based provider usage without an API key.  
  → [PR #4632](https://github.com/HKUDS/nanobot/pull/4632)

- **[#4688] feat(cli): add safe WebUI first-run launcher**  
  Closed. Introduces `nanobot webui` command that verifies config/provider before launching, with Quick Start fallback.  
  → [PR #4688](https://github.com/HKUDS/nanobot/pull/4688)

- **[#4396] Add optional Nanobot plugin controls**  
  Closed. Implements `nanobot plugins list/enable/disable` commands and moves heavier capabilities behind explicit enablement.  
  → [PR #4396](https://github.com/HKUDS/nanobot/pull/4396)

- **[#4691] fix(plugins): polish optional feature controls**  
  Closed. Follow-up polish for the plugin system: warns on missing deps, keeps disabled plugins simply inactive, and adds recovery instructions.  
  → [PR #4691](https://github.com/HKUDS/nanobot/pull/4691)

## 4. Community Hot Topics
The most active issues and PRs (by comments/reactions) indicate two pressing pain points:

- **[Issue #4652] [bug] Nanobot process crashes directly when MCP tool call exception**  
  *3 comments, created Jul 2* – Reports that MCP tool call errors or empty data crash the process. A fix PR #4666 is already open.  
  → [Issue #4652](https://github.com/HKUDS/nanobot/issues/4652)

- **[Issue #4302] [bug] nanobot gateway crashes after MCP reconnect**  
  *2 comments, last updated today* – Similar crash pattern at the gateway level after session termination. Comments link to #4211.  
  → [Issue #4302](https://github.com/HKUDS/nanobot/issues/4302)

- **[Issue #4693] WebUI: improve mobile responsive layout for chat viewport and composer**  
  *0 comments but followed by PR #4694* – Users report clipped UI on mobile browsers. PR #4694 (fix) was opened today.  
  → [Issue #4693](https://github.com/HKUDS/nanobot/issues/4693)

**Underlying need**: Users depend on MCP tools for production use and encounter crashes under error conditions. The mobile experience gap suggests growing usage from mobile devices.

## 5. Bugs & Stability
Three bugs were reported or updated today, ranked by severity:

| Severity | Issue/PR | Description | Fix status |
|----------|----------|-------------|------------|
| **Critical** | [#4652](https://github.com/HKUDS/nanobot/issues/4652) | Process crashes on MCP tool call exception | PR [#4666](https://github.com/HKUDS/nanobot/pull/4666) open (contains malformed tool results) |
| **High** | [#4302](https://github.com/HKUDS/nanobot/issues/4302) | Gateway crashes after MCP reconnect (since Jun 11) | No fix PR yet; maintainers aware |
| **Medium** | [#4690](https://github.com/HKUDS/nanobot/pull/4690) | `nanobot gateway stop` crashes on Windows with `OSError: [WinError 87]` | PR open (fix: handle Windows stop fallback) |

Additionally, PR [#4666](https://github.com/HKUDS/nanobot/pull/4666) addresses #4652 by wrapping MCP result rendering, marking timeouts and failures as structured errors.

## 6. Feature Requests & Roadmap Signals
Several open PRs indicate upcoming capabilities likely to land in the next release:

- **Mattermost channel support** – [#4459](https://github.com/HKUDS/nanobot/pull/4459) (open, updated today) – Real-time messaging with streaming responses.
- **Canonical OpenCode provider** – [#4686](https://github.com/HKUDS/nanobot/pull/4686) (open) – Adds `opencode` alongside existing `opencode_zen`, keeping compatibility.
- **Cron job model presets** – [#4622](https://github.com/HKUDS/nanobot/pull/4622) (open) – Allows per-run provider/model overrides for scheduled jobs.
- **Memory: archive facts with provenance** – [#4621](https://github.com/HKUDS/nanobot/pull/4621) (open) – Includes MEMORY.md excerpt in Consolidator prompts to reduce duplicates.
- **Heartbeat trigger command** – [#4620](https://github.com/HKUDS/nanobot/pull/4620) (open) – Adds `nanobot heartbeat trigger` for CLI and gateway timer paths.
- **Dream duplicate skill guard** – [#4554](https://github.com/HKUDS/nanobot/pull/4554) (open) – Prevents Dream from overwriting existing skill directories.
- **Search_history tool** – [#4439](https://github.com/HKUDS/nanobot/pull/4439) (open) – Read-only memory recall tool.

**Prediction**: The next minor release will likely include the MCP crash fix (#4666), mobile layout fix (#4694), and the merged plugin controls (#4396). The Anthropic OAuth provider (#4632) and config camelCase fix (#4692) landed today and may be released soon.

## 7. User Feedback Summary
Pain points surfaced in the last 24 hours:

- **MCP reliability**: Users report crashes when tools return errors or after reconnection – this is the top stability issue.
- **Mobile usability**: “WebUI is hard to use from a mobile browser” – clipping and horizontal scrolling frustrate mobile users.
- **Windows compatibility**: Gateway stop command crashes on Windows, indicating less-tested platform edge case.
- **Dream inconsistency**: Users (via PR #4673) report that Dream’s commit logs don’t match actual file diffs – the team is working on grounding memory audits in real diff.
- **Plugin workflow**: No explicit complaints, but the new plugin system (#4396) has been polished based on early user feedback (#4691).

Positive signals: The community is actively contributing (many PRs from external authors), and the team quickly reacts with fixes (e.g., #4694 opened same day as issue #4693).

## 8. Backlog Watch
- **[Issue #4302] Gateway crash after MCP reconnect** (created Jun 11) – No fix PR yet; the issue’s age (24 days) and severity (crash) call for maintainer attention. The similar #4652 has a fix, but #4302 may need a different approach.
- **[Issue #4440] (referenced by PR #4439)** – Not directly in the 24h dataset, but the PR for search_history tool (open since Jun 21) references this issue – it may be a long-standing feature request.
- **[Issue #4211] (mentioned in #4302)** – Not listed in the recent data but is the root cause reference; likely still lingering.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-07-04

## Today’s Overview
Hermes Agent shows **high-activity** today with 50 PRs updated in the last 24 hours (18 merged/closed) and 12 issues updated (11 still open). No new releases were published. The project is processing a **significant backlog of bug fixes and feature work**, particularly around credential isolation, session state integrity, and multi-model support. However, a critical new bug—context compaction producing invalid message sequences (#58168)—represents a potential **session-breaking regression** that will require immediate attention. The community is actively contributing patches, especially for Telegram routing, SSH terminal behavior, and authentication improvements.

---

## Releases
**None.** No new releases were published today.

---

## Project Progress
**18 pull requests were merged or closed today.** Key areas that advanced:

- **Agent runtime improvements:**
  - [PR #58155](https://github.com/NousResearch/hermes-agent/pull/58155) — Codex `max_output_tokens` recovery: turns that hit `status=incomplete / reason=max_output_tokens` now resume instead of dead-ending; mid-turn compaction added to prevent context exhaustion.
  - [PR #58159](https://github.com/NousResearch/hermes-agent/pull/58159) — SSH `~`/`~/…` remote `cwd` preserved instead of expanded against host HOME.
  - [PR #58156](https://github.com/NousResearch/hermes-agent/pull/58156) — `reasoning_effort` now forwarded to custom OpenAI-compatible providers (GLM-5.2 on ARK, vLLM, Ollama, etc.).

- **Cron and Kanban fixes:**
  - [PR #52079](https://github.com/NousResearch/hermes-agent/pull/52079) — Cron deliveries to Telegram forum topics now route correctly via `message_thread_id`.
  - [PR #57713](https://github.com/NousResearch/hermes-agent/pull/57713) — Kanban workers terminated when tasks are blocked or archived.
  - [PR #57536](https://github.com/NousResearch/hermes-agent/pull/57536) — Fix double-encoding of reasoning columns on session fork round-trip.

- **Authentication and state fixes:**
  - [PR #57548](https://github.com/NousResearch/hermes-agent/pull/57548) — Custom provider `max_tokens` now honored at startup.
  - [PR #57696](https://github.com/NousResearch/hermes-agent/pull/57696) — Windows MSYS path normalization in gateway media delivery.
  - [PR #57601](https://github.com/NousResearch/hermes-agent/pull/57601) — `reasoning_effort` forward for custom providers (GLM-5.2 on ARK).

---

## Community Hot Topics

### Most Active Issues (by comments/reactions)

1. **[#24443](https://github.com/NousResearch/hermes-agent/issues/24443) — MiMo reasoning models fail because `reasoning_content` not preserved in chat history**
   - *5 comments, 1 👍, open since 2026-05-12*
   - Underlying need: MiMo’s OpenAI-compatible API requires clients to echo back `reasoning_content` from prior assistant messages. Hermes strips it, breaking multi-turn reasoning. **High community interest** for multi-modal reasoning support.

2. **[#38129](https://github.com/NousResearch/hermes-agent/issues/38129) — Cron sessions expose the memory tool but memory is unavailable at runtime**
   - *3 comments, 1 👍, open since 2026-06-03*
   - Pain point: Inconsistent UX—tools visible but non-functional in cron jobs. Confuses users and wastes agent turns.

3. **[#29530](https://github.com/NousResearch/hermes-agent/issues/29530) — Profiled workers need shared auth home separate from isolated `HERMES_HOME`**
   - *3 comments, open since 2026-05-20*
   - Risk: Split-brain OAuth state when multiple profiled workers share refresh tokens. Potential credential corruption.

4. **[#37906](https://github.com/NousResearch/hermes-agent/issues/37906) — WhatsApp `send_message` fails: `@lid` JIDs not recognized, phone numbers cause `jidDecode` error**
   - *2 comments, open since 2026-06-03*
   - Three related bugs in WhatsApp message routing. Affects users relying on WhatsApp as a primary channel.

5. **[#57836](https://github.com/NousResearch/hermes-agent/issues/57836) — Headless MCP OAuth blocks gateway startup with stale cached tokens** *(CLOSED)*
   - *2 comments, closed 2026-07-04*
   - Gateway startup blocked by MCP OAuth discovery timeout. Fixed via [PR #57563](https://github.com/NousResearch/hermes-agent/pull/57563).

---

## Bugs & Stability

### Critical (session-breaking)

- **[#58168](https://github.com/NousResearch/hermes-agent/issues/58168) — Context compaction produces invalid message sequences (orphaned tool messages)** *(NEW, 0 comments)*
  - **Severity: Critical.** The trajectory compressor (`trajectory_compressor.py`) generates sequences where `role: tool` messages lack a preceding assistant call, causing the DeepSeek API to reject requests with HTTP 400. Sessions are **permanently broken** after compaction.
  - **No fix PR yet.** An immediate stopgap or fix is needed.

### High (platform stability, data integrity)

- **[#58150](https://github.com/NousResearch/hermes-agent/issues/58150) — Hermes uses Windows system Node.js instead of its built-in node** *(NEW, 1 comment)*
  - **Severity: High (Windows).** Causes node terminal flashing with version mismatch.
  - **No fix PR yet.**

- **[#58167](https://github.com/NousResearch/hermes-agent/issues/58167) — SendGrid prefix pattern masks only key-id, secret stays in cleartext** *(NEW, 0 comments)*
  - **Severity: High (security).** `_PREFIX_PATTERNS` regex matches only two of three SendGrid key segments, leaving the key-secret segment unmasked.
  - **No fix PR yet.**

- **[#58166](https://github.com/NousResearch/hermes-agent/issues/58166) — `GET /auth/login?provider=basic` returns HTTP 500** *(NEW, 0 comments)*
  - **Severity: High.** Dashboard login with the built-in `basic` auth provider returns a 500 error instead of the password form.
  - **No fix PR yet.**

- **[#58158](https://github.com/NousResearch/hermes-agent/pull/58158) — Auth store reads as locale-default encoding (cp1252 on Windows), wiping credentials** *(PR OPEN)*
  - **Severity: High (Windows).** Non-ASCII bytes cause `UnicodeDecodeError` and credential loss. Fix PR is open.
  - **Fix PR exists:** [PR #58158](https://github.com/NousResearch/hermes-agent/pull/58158)

### Medium

- **[#58160](https://github.com/NousResearch/hermes-agent/pull/58160) — Cron jobs whose final response admits unfinished guardrail work** *(PR OPEN)*
  - Cron marks `ok` for jobs that state pending validation/cleanup. Fix PR is open.

---

## Feature Requests & Roadmap Signals

- **[#58161](https://github.com/NousResearch/hermes-agent/issues/58161) — Desktop app: adjustable interface font size** *(NEW, duplicate)*
  - User with poor eyesight requests font size customization. Likely to be addressed in a future desktop release if the project maintains accessibility focus.

- **[#58151](https://github.com/NousResearch/hermes-agent/issues/58151) — CLI commands and agent-assisted session archiving/organization** *(NEW)*
  - Real-world session management pain point: no way to archive, tag, or clean up old sessions. Could drive a **session management feature** in a minor release—especially if the agent itself can assist with organization.

**Prediction:** Desktop accessibility (#58161) may land in a patch release if the project has a desktop team. Session archiving (#58151) is a stronger roadmap signal—introducing tags or folder organization aligns with growing enterprise adoption.

---

## User Feedback Summary

- **Pain point: Multi-turn reasoning loss.** Multiple users report MiMo and reasoning model failures when `reasoning_content` is not preserved across turns (#24443). This is the **most-voted issue** in recent data.
- **Pain point: Cron tool inconsistency.** Users expect cron-available tools to work; the memory tool failing at runtime (#38129) erodes trust in cron jobs.
- **Pain point: WhatsApp reliability.** Three distinct routing bugs (#37906) frustrate WhatsApp-first users.
- **Pain point: Configuration confusion.** Provider names silently replaced with `"custom"` (#26879), reasoning_effort dropped (#57601), and BWS secret injection failures (#58163) indicate **config churn** that needs UX polish.
- **Positive signal:** Community contributions are strong—multiple PRs from new contributors (kshitijk4poor, nankingjing, AlexFucuson9) fixing long-standing issues.

---

## Backlog Watch

These issues are **open for >30 days** without a fix PR, and may need maintainer attention:

1. **[#24443](https://github.com/NousResearch/hermes-agent/issues/24443) — MiMo reasoning models: `reasoning_content` not preserved** (P2, open since 2026-05-12, 5 comments)
   - High community interest (1 👍). Requires API contract alignment with MiMo’s OpenAI-compatible endpoint.

2. **[#26879](https://github.com/NousResearch/hermes-agent/issues/26879) — Auxiliary task provider identity lost when `base_url` + `api_key` set** (P3, open since 2026-05-16, 2 comments)
   - Silent replacement of provider name with `"custom"` breaks provider-specific code paths for ZAI, Anthropic, etc.

3. **[#29530](https://github.com/NousResearch/hermes-agent/issues/29530) — Profiled workers need shared auth home** (P2, open since 2026-05-20, 3 comments)
   - Credential isolation is a **security risk**. One partial fix is in [PR #57563](https://github.com/NousResearch/hermes-agent/pull/57563) (still open).

4. **[#38129](https://github.com/NousResearch/hermes-agent/issues/38129) — Cron memory tool unavailable at runtime** (P3, open since 2026-06-03, 3 comments)
   - No PR yet. A relatively contained fix: either hide the tool or make memory available during cron execution.

5. **[#37906](https://github.com/NousResearch/hermes-agent/issues/37906) — WhatsApp `send_message` three-bug cluster** (P2, open since 2026-06-03, 2 comments)
   - No PR yet. The `@lid` JID parsing issue alone blocks a core use case for WhatsApp users.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw Project Digest — 2026-07-04

## 1. Today’s Overview
Activity remains moderate with **2 open issues** (both stale) and **12 PRs updated** in the last 24 hours. **4 PRs were merged or closed** today, mainly stability fixes and code cleanup. No new releases were published. The project continues to see steady contributor work on agent collaboration, WhatsApp resilience, and new channel support, but several older bug reports are lingering without clear resolution.

## 2. Releases
**None.** No new versions were published.

## 3. Project Progress (Merged/Closed PRs Today)
- **#3223** – `fix(agent): clear routed agent session` (closed in favor of a corrected PR #3224)  
  URL: https://github.com/sipeed/picoclaw/pull/3223
- **#3128** – `fix(web): explicitly ignore resp.Body.Close() errors after io.ReadAll` (merged)  
  URL: https://github.com/sipeed/picoclaw/pull/3128  
  *Stability fix for search providers (Bing, Tavily, Sogou, Perplexity).*
- **#3142** – `fix(spawn): clear ForUser in sub-turn ToolResult to prevent duplicate messages` (merged)  
  URL: https://github.com/sipeed/picoclaw/pull/3142  
  *Fixes a duplicate push issue when async sub-agents complete.*
- **#3063** – `feat: add deltachat gateway` (merged)  
  URL: https://github.com/sipeed/picoclaw/pull/3063  
  *DeltaChat integration landed and is now being refined (see open PR #3222).*

These merges improve agent routing, web search robustness, sub-agent messaging, and expand channel support.

## 4. Community Hot Topics
- **#2937 – Feat/agent collaboration** (open since May 24)  
  URL: https://github.com/sipeed/picoclaw/pull/2937  
  *A large feature adding a first-class inter‑agent collaboration bus. No recent comments but still actively updated – likely a major roadmap item.*
- **#3225 – Support agent-specific runtime overrides** (opened today)  
  URL: https://github.com/sipeed/picoclaw/pull/3225  
  *Quickly gaining attention; allows per-agent `max_tokens`, summarization thresholds, and split_on_marker. Signals strong user demand for flexible agent configuration.*
- **#3193 – Added simplex channel type** (open since June 27)  
  URL: https://github.com/sipeed/picoclaw/pull/3193  
  *New channel integration; community interest from users wanting secure/p2p messaging.*

Underlying need: Users want more control over agent behavior and broader channel support, as reflected in both collaboration features and configuration overrides.

## 5. Bugs & Stability
**Moderate severity, both stale:**

- **#3182 – [BUG] Android version – cannot launch service**  
  URL: https://github.com/sipeed/picoclaw/issues/3182  
  *User reports inability to launch the service despite full permissions. No fix PR identified; remains unaddressed.*
- **#3178 – [BUG] WhatsApp Websocket Timeout**  
  URL: https://github.com/sipeed/picoclaw/issues/3178  
  *WebSocket connection drops after adding a schedule. PR #3179 (open) attempts to fix reconnection logic – not yet merged.*  
  URL: https://github.com/sipeed/picoclaw/pull/3179

**Severity ranking:** Android launch failure is critical for mobile users; WhatsApp timeout is high (core connectivity). Both need maintainer attention.

## 6. Feature Requests & Roadmap Signals
- **Agent-specific runtime overrides** (PR #3225) – likely to land in next minor release.
- **Agent collaboration bus** (PR #2937) – ambitious feature; could become a v0.3.0 highlight.
- **DeltaChat gateway refactoring** (PR #3222) – cleanup and documentation after initial merge; signals ongoing polish.
- **Configurable model fallback chain** (PR #3200) – adds default fallback UI and persistence, a frequent user request.
- **Simplex channel** (PR #3193) – new channel type.

**Predictions:** Next release (v0.3.0 or v0.2.10) will likely include per‑agent overrides, model fallback, and possibly the agent collaboration bus as beta.

## 7. User Feedback Summary
- **Pain point:** Android support is broken for at least one user; no workaround provided.
- **Pain point:** WhatsApp WebSocket timeout disrupts scheduled tasks; PR #3179 exists but not merged.
- **Use case:** Multiple agents with different models/configs (PR #3225, #3200) – users want granular control.
- **Satisfaction:** DeltaChat integration and collaboration bus are well‑received; clean‑up PR #3222 shows active engagement from contributor `trufae`.

Overall sentiment seems constructive, but the two stale bugs may indicate maintainer bandwidth constraints.

## 8. Backlog Watch
- **#2937 – Feat/agent collaboration** (open 42 days) – large, complex, may need maintainer review to merge.
- **#3165 – fix(openai_compat): recover Seed XML tool calls** (open 10 days) – awaits review; could unblock Volcengine users.
- **#3178 / #3182** – Both stale bugs with no maintainer response; risk of user frustration.
- **#3128** – Already merged, but its old creation date (June 15) shows how long code review can take.

**Call to action:** Maintainers should prioritise the two stale bugs and the open #3165 fix to keep community trust.

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest – 2026-07-04

## Today’s Overview

NanoClaw saw no new issues or releases in the past 24 hours, but significant PR activity continued with 20 pull requests updated, of which 3 were merged or closed (PRs #2611, #2630, #2765). The project remains in a steady maintenance and feature‑integration phase, with a concentration of security‑hardening and bug‑fix PRs moving toward completion. Community contributions remain strong, with 17 open PRs actively receiving updates from both maintainers and external authors. The overall project health is positive, though the backlog of long‑standing open PRs (some dating to early May) indicates that review throughput could benefit from additional focus.

## Releases

No new releases were published today. The latest stable version remains unchanged; users should refer to the project’s main branch for recent fixes.

## Project Progress

Three pull requests were closed/merged today:

- **[PR #2611 – fix(cli): preserve caller context after approval](https://github.com/nanocoai/nanoclaw/pull/2611)** – Merged. This security fix ensures that when an admin approves a CLI command, the original caller context is preserved during replay, preventing privilege escalation or context mix‑ups.
- **[PR #2630 – fix(session-manager): confine target inbox roots](https://github.com/nanocoai/nanoclaw/pull/2630)** – Merged. Hardens inbound attachment handling by refusing to write through symlinked session inbox roots, closing a potential path‑traversal vector.
- **[PR #2765 – feat(providers): add .format-lint-off](https://github.com/nanocoai/nanoclaw/pull/2765)** – Closed (merged). Adds a configuration directive to suppress formatting/linting on specific sections, a minor quality‑of‑life improvement for provider configurations.

Additionally, a new bug‑fix PR was opened today:

- **[PR #2922 – fix(discord): unwrap forwarded-message snapshots](https://github.com/nanocoai/nanoclaw/pull/2922)** – Addresses an issue where Discord forwarded messages were presented as opaque snapshots, making the original content invisible to agents. The fix extracts the embedded message so agents can process it normally.

## Community Hot Topics

No issues were reported or commented on in the last 24 hours, and none of the updated PRs received comments or reactions (all showing 0). The most active discussions are occurring implicitly through ongoing PR updates. Notable long‑running PRs that continue to receive attention from multiple contributors include:

- **[PR #2208](https://github.com/nanocoai/nanoclaw/pull/2208) – feat(mcp): support http and sse MCP server transports** (updated Jul 3) – A feature addition that would unlock MCP servers using HTTP/SSE, currently in review.
- **[PR #2863](https://github.com/nanocoai/nanoclaw/pull/2863) – feat(skills): add /setup-system-digest and /system-digest** (updated Jul 3) – A utility skill for automated system health digests, contributed by grantland.

The lack of explicit comments or reactions suggests that most review and discussion is happening offline or in dedicated channels; however, contributors are actively pushing updates to their PRs.

## Bugs & Stability

No new bugs were *reported* today (zero issues opened), but several open bug‑fix PRs are in the pipeline, many of which address high‑severity issues:

| Priority | Issue / PR | Description | Status |
|----------|------------|-------------|--------|
| **High** | **PR #2920** (open) – [DB connection leak in container restart](https://github.com/nanocoai/nanoclaw/pull/2920) | `openInboundDb()` was never closed after checking pending messages, leaking file descriptors on each container restart. Also fixes stale documentation and a duplicate script. | Open, created today |
| **High** | **PR #2694** (open) – [Signal DM silently dropped](https://github.com/nanocoai/nanoclaw/pull/2694) | Signal adapter does not set `isMention`/`isGroup` on inbound DMs, causing the router to drop them entirely. | Open, updated Jul 3 |
| **High** | **PR #2695** (open) – [Signal image attachments unreadable in container](https://github.com/nanocoai/nanoclaw/pull/2695) | Host‑side attachment paths are emitted but not mounted in the agent container; fix stages images as base64. | Open, updated Jul 3 |
| **Medium** | **PR #2531** (open) – [Duplicate text when send_message fires mid‑turn](https://github.com/nanocoai/nanoclaw/pull/2531) | Poll loop duplicates output text if the agent calls `send_message` while still generating a response. | Open, updated Jul 3 |
| **Medium** | **PR #2184** (open) – [Stale session error delivered to user instead of retry](https://github.com/nanocoai/nanoclaw/pull/2184) | After session expiry, an error message was shown before the next poll could restart fresh; fix retries immediately. | Open, updated Jul 3 |
| **Low** | **PR #2348** (open) – [WhatsApp single-timer reconnect + clean teardown](https://github.com/nanocoai/nanoclaw/pull/2348) | Ensures safe reconnection and teardown for the WhatsApp channel. | Open, updated Jul 3 |

No regressions from recent merges were reported. The closed security PRs (#2611, #2630) directly improve stability.

## Feature Requests & Roadmap Signals

Several feature additions are on the verge of landing or under active development:

- **MCP transport expansion** – [#2208](https://github.com/nanocoai/nanoclaw/pull/2208) would allow NanoClaw tools to connect to MCP servers over HTTP and SSE, a frequently requested capability.
- **Utility skills** – [#2863](https://github.com/nanocoai/nanoclaw/pull/2863) (system digest) and [#2530](https://github.com/nanocoai/nanoclaw/pull/2530) (CalDAV tool), along with [#2693](https://github.com/nanocoai/nanoclaw/pull/2693) (Google Contacts tool), indicate growing investment in the skill ecosystem.
- **Container runner improvements** – [#2230](https://github.com/nanocoai/nanoclaw/pull/2230) adds rootless Podman support via `keep-id`, enabling secure container execution on systems without root privilege.
- **Compose fragment gating** – [#2921](https://github.com/nanocoai/nanoclaw/pull/2921) fixes an important compose bug that inlined all skill instructions into every group’s CLAUDE.md, regardless of per‑group skill selection. Once merged, this will improve customization for multi‑group setups.

No explicit user feature requests were filed today, but the volume of utility‑skill PRs suggests a community appetite for plug‑and‑play tools.

## User Feedback Summary

User pain points visible through PR descriptions:

- **Discord forwarded messages invisible to agents** (PR #2922) – Users sending forwarded content could not have agents react to it.
- **Signal integration broken for DMs and images** (PRs #2694, #2695) – Two separate issues made Signal essentially unusable in some scenarios: DMs were dropped, and image attachments could not be read by the containerised agent.
- **Duplicate text output during agent turns** (PR #2531) – Mid‑turn tool calls caused the same text to appear twice, frustrating interaction flow.
- **Configuration confusion** (PR #2920) – Stale references to `NANOCLAW_ADMIN_USER_IDS` in documentation caused user confusion about v2 migration.

Overall sentiment appears satisfied with the pace of bug fixes, though the Signal issues have likely been a source of frustration for users relying on that channel.

## Backlog Watch

Several important PRs have been open for more than a month and need maintainer attention or final review:

- **[PR #2184](https://github.com/nanocoai/nanoclaw/pull/2184) – fix(poll-loop): retry immediately on stale session** (opened May 2) – Addresses a user‑visible error scenario. Updated Jul 3 with fixes; ready for merge.
- **[PR #2208](https://github.com/nanocoai/nanoclaw/pull/2208) – feat(mcp): support http and sse** (opened May 3) – A highly requested feature; has been in review for two months.
- **[PR #2230](https://github.com/nanocoai/nanoclaw/pull/2230) – fix(container-runner): map host user via keep-id on rootless podman** (opened May 3) – Important for container security.
- **[PR #2416](https://github.com/nanocoai/nanoclaw/pull/2416) – fix(cli): provision companion rows on `ncl groups create`** (opened May 11) – A CLI correctness fix that may cause DB inconsistencies if not merged.
- **[PR #2348](https://github.com/nanocoai/nanoclaw/pull/2348) – fix(channels/whatsapp): single-timer reconnect** (opened May 8) – Stability fix for WhatsApp channel.

None of these have received negative feedback, but the age suggests they may be blocked by review capacity. Maintainers should prioritise #2184 and #2694/2695, as they directly affect core user experience (session errors and Signal connectivity).

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest — 2026-07-04

## 1. Today's Overview
Project activity is minimal over the last 24 hours. No new releases or pull requests were created or updated. The only movement is a single open bug report (Issue #972) that received its last comment on July 3 and remains unaddressed by maintainers. With zero merged PRs and no new contributions, the project appears to be in a low‑activity phase, with no visible progress on features or fixes.

## 2. Releases
- **No new releases** in the past 24 hours. The most recent release information is unavailable.

## 3. Project Progress
- **No pull requests** were updated (open, merged, or closed) in the last 24 hours. No features, fixes, or improvements have been merged.

## 4. Community Hot Topics
- **[Issue #972 – [bug] telegram channel stop respond after some idle time](https://github.com/nullclaw/nullclaw/issues/972)**  
  *Author: i11010520 | Created: 2026‑06‑30 | Updated: 2026‑07‑03 | 1 comment*  
  This is the only active discussion. The reporter describes that the Telegram channel stops responding after a period of inactivity (e.g. overnight), while the NullClaw backend continues to operate normally (ping works). The underlying need appears to be a **persistent connection or reconnection mechanism** for Telegram adapters that can handle idle timeouts without manual intervention. The issue has no negative reactions, but also no maintainer response yet.

## 5. Bugs & Stability
- **Severity: Medium** — The only reported bug is Issue #972 (Telegram channel silent after idle). It disrupts the user experience for Telegram‑based interactions but does not crash the backend. No fix PR exists. The root cause is likely a connection timeout or session expiry in the Telegram API or the channel adapter. Until addressed, users relying on Telegram may need to restart the channel manually after idle periods.

## 6. Feature Requests & Roadmap Signals
- **No explicit feature requests** were filed in the last 24 hours. The single bug report does not suggest a new feature, but it highlights a reliability gap that could be addressed by adding an automatic reconnection or heartbeat mechanism in a future release.

## 7. User Feedback Summary
- **Pain point:** The main (and only) user feedback is the Telegram channel becoming unresponsive after idle time. The reporter indicates that the core agent works fine, pointing to a channel‑specific connectivity issue. User satisfaction is likely reduced for those using Telegram, though no expressions of dissatisfaction beyond the bug report are present. No praise or other use cases were shared.

## 8. Backlog Watch
- **[Issue #972 – Telegram channel idle timeout](https://github.com/nullclaw/nullclaw/issues/972)**  
  Opened 4 days ago, last updated yesterday. While not extremely old, it is the **only open issue** and has received no maintainer acknowledgement or assignment. With zero PRs in flight, this issue may become a higher priority signal if more users report the same problem. Maintainers should triage and either confirm the bug or request additional information to reproduce it.

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

## IronClaw Project Digest — 2026-07-04

### 1. Today’s Overview

IronClaw maintained a high level of activity over the past 24 hours, with **50 pull requests** updated (23 merged/closed) and **8 issues** updated (4 closed). The project remains heavily focused on the **Reborn architecture**, particularly credential-injection semantics, ingress route management, and integration test coverage. CI stability is receiving urgent attention—several recent workflow failures on `main` are being addressed through dedicated fixes and build-time optimizations. A new contributor wave is also visible, with multiple open PRs fixing desktop-app model handling and trigger persistence.

### 2. Releases

**No new releases** in the last 24 hours.

### 3. Project Progress — Merged/Closed PRs

The following pull requests were merged or closed today, representing concrete progress:

- **CI & Build**  
  - [#5591](https://github.com/nearai/ironclaw/pull/5591) — Stabilized `main`-equivalent clippy and coverage checks (fixing WASIp1 compile path, coverage failures around Reborn live streams).  
  - [#5628](https://github.com/nearai/ironclaw/pull/5628) — Documented build-time optimization targets and added baseline performance scripts for CI.  
  - [#5630](https://github.com/nearai/ironclaw/pull/5630) — Fixed Docker build by copying the `prompts/` directory into the legacy builder stage (open but likely to merge shortly).

- **Testing & Test Infrastructure**  
  - [#5610](https://github.com/nearai/ironclaw/pull/5610) — Wave‑4 integration coverage for Reborn: auth-gate wire regression, triggered auth delivery, attachments, and golden/synthetic expansions.  
  - [#5609](https://github.com/nearai/ironclaw/pull/5609) — Extracted trigger-prompt materializer into shared test support, reducing duplication in integration harnesses.

- **Documentation**  
  - [#5628](https://github.com/nearai/ironclaw/pull/5628) (also listed above) provided formal CI build-time loss targets and anti-cheat constraints.

Additionally, **4 issues were closed**: [#3087](https://github.com/nearai/ironclaw/issues/3087) (Reborn host runtime services composition), [#3231](https://github.com/nearai/ironclaw/issues/3231) (architecture deepening follow‑up), [#3068](https://github.com/nearai/ironclaw/issues/3068) (credential injection preservation), and [#5590](https://github.com/nearai/ironclaw/issues/5590) (main CI checks green). These closures mark progress on Reborn cutover blockers and overall stability.

### 4. Community Hot Topics

The most active discussions occurred on older architecture issues, with top comment counts:

- **[#3087](https://github.com/nearai/ironclaw/issues/3087)** (7 comments) — Closed today; tracked the composition of `ironclaw_host_runtime` services.  
- **[#3278](https://github.com/nearai/ironclaw/issues/3278)** (3 comments) — Still open; defines the `MissionService` integration with `TurnCoordinator` and touches multiple related contracts.  
- **[#3231](https://github.com/nearai/ironclaw/issues/3231)** (3 comments) — Closed; tracked follow‑up architecture deepening after the Reborn substrate landed.  
- **[#3238](https://github.com/nearai/ironclaw/issues/3238)** (2 comments) — Open; seeks comprehensive cancellation semantics for Reborn turn runs.

Among PRs, the most commented (though `Comments: undefined` in the data) are likely those with multiple reviewers or complex changes:

- [#5623](https://github.com/nearai/ironclaw/pull/5623) — Honor staged credential obligations for WASM egress (size L).  
- [#5598](https://github.com/nearai/ironclaw/pull/5598) — Release automation bumping several crate versions.  
- [#5550](https://github.com/nearai/ironclaw/pull/5550) — Dependency update affecting 13 packages (including agent-client-protocol 0.10.4 → 1.0.1).

The underlying need across these discussions is the community’s desire for **clear, well-documented interfaces** (ingress routes, credential injection, cancellation) before the Reborn architecture becomes the default.

### 5. Bugs & Stability

Several stability issues were highlighted in the last 24 hours, ranked by severity:

- **Critical — CI failure on `main`**  
  [#5590](https://github.com/nearai/ironclaw/issues/5590) reported that recent `main` branch checks are not green across code style, deterministic tests, and live/browser QA workflows. **Fix PRs** [#5591](https://github.com/nearai/ironclaw/pull/5591) and [#5630](https://github.com/nearai/ironclaw/pull/5630) address the root causes.

- **High — Nightly E2E failure**  
  [#4108](https://github.com/nearai/ironclaw/issues/4108) (open since May 27) continues to fail. The latest run (2026-07-04) shows a full E2E job failure. No dedicated fix PR is yet linked.

- **Medium — WASM credential provider bypasses authorizer**  
  [#5512](https://github.com/nearai/ironclaw/issues/5512) (open) — The WASM credential provider re-derives injection eligibility from the manifest instead of consulting the authorizer’s `Decision.obligations`. This can cause incorrect credential injection. **Fix PR** [#5623](https://github.com/nearai/ironclaw/pull/5623) is open and targets exactly this issue.

- **Low — Retry hang on invalid model**  
  PRs [#5043](https://github.com/nearai/ironclaw/pull/5043) and [#5045](https://github.com/nearai/ironclaw/pull/5045) (both open) fix silent hangs when `NEARAI_MODEL=auto` is sent to a cloud API that does not recognize the “auto” alias.

### 6. Feature Requests & Roadmap Signals

The following feature-related issues and PRs point toward upcoming work:

- **Reborn Cancellation Semantics** ([#3238](https://github.com/nearai/ironclaw/issues/3238)) — A comprehensive spec for end‑to‑end turn cancellation, including who may cancel, state transitions, and recovery. Likely to be implemented after the current credential‑injection work lands.

- **MissionService & TurnCoordinator Integration** ([#3278](https://github.com/nearai/ironclaw/issues/3278)) — Ties together product-surface migration, conversation binding, and session thread contracts. A major next phase for Reborn.

- **Slack Ingress from Manifest** ([#5626](https://github.com/nearai/ironclaw/pull/5626)) — Moves Slack’s two inbound routes from Rust policy literals to manifest‑driven data. This reduces code duplication and improves extensibility.

- **Skill‑Learning Enhancements** ([#5156](https://github.com/nearai/ironclaw/pull/5156), open) — Lands an approval gate, learned‑only scoping, and persisted switches. Previously deferred as residual risk, this PR makes freshly‑learned skills inactive until human approval.

- **WebUI v2 Improvements** — Several open PRs from new contributors fix SSE tool‑activity streaming ([#5160](https://github.com/nearai/ironclaw/pull/5160)), invalid chat routes ([#5132](https://github.com/nearai/ironclaw/pull/5132)), and headless routine persistence ([#5041](https://github.com/nearai/ironclaw/pull/5041)).

Predicting next‑version inclusion: The credential‑obligation fix [#5623](https://github.com/nearai/ironclaw/pull/5623) and the CI stabilisation PRs are likely to merge first. The Slack ingress manifest PR [#5626](https://github.com/nearai/ironclaw/pull/5626) and the skill‑learning approval gate [#5156](https://github.com/nearai/ironclaw/pull/5156) appear mature enough for the upcoming release.

### 7. User Feedback Summary

Direct user feedback is not abundant in this dataset, but the following pain points emerge from bug reports and PR descriptions:

- **Desktop‑app model configuration** — Users setting `NEARAI_MODEL=auto` experience silent hangs because the cloud API rejects “auto”. PRs [#5043](https://github.com/nearai/ironclaw/pull/5043) and [#5045](https://github.com/nearai/ironclaw/pull/5045) address this by failing fast and resolving the alias.

- **Headless trigger persistence** — Routines created with “every N minutes” in WebChat v2 do not persist due to policy‑gate denial and missing cron‑seconds. [#5041](https://github.com/nearai/ironclaw/pull/5041) fixes three root causes.

- **One‑line answers containing tool names** — A genuine final answer like “Tool result from web__fetch: …” was misclassified as a transcript artifact and lost. [#5042](https://github.com/nearai/ironclaw/pull/5042) provides a weak‑model heuristic.

- **Credential injection inconsistency** — The WASM credential provider ignoring the authorizer’s decision (issue [#5512](https://github.com/nearai/ironclaw/issues/5512)) could cause incorrect credential exposure, a safety concern for extension developers.

Overall satisfaction appears moderate: the project is making solid architectural progress, but recurring CI failures and credential‑related regressions create friction for contributors and early adopters.

### 8. Backlog Watch

The following issues and PRs have been open for an extended period without recent maintainer action:

- **[#3278](https://github.com/nearai/ironclaw/issues/3278)** (created May 6, updated Jul 3) — MissionService integration with TurnCoordinator. No assignee or recent comments from maintainers.

- **[#3238](https://github.com/nearai/ironclaw/issues/3238)** (created May 3, updated Jul 3) — Cancellation semantics. This spec is foundational for Reborn but appears to be waiting for earlier work (credential injection) to land.

- **[#4108](https://github.com/nearai/ironclaw/issues/4108)** (created May 27, updated Jul 4) — Nightly E2E failure. Still open with no linked fix PR; the failure is recurring.

- Pull requests from new contributors that have been open for 2–3 weeks without merge or review:
  - [#5041](https://github.com/nearai/ironclaw/pull/5041) (Jun 17) — Trigger persistence fix.
  - [#5042](https://github.com/nearai/ironclaw/pull/5042) (Jun 17) — One‑line answer misclassification.
  - [#5043](https://github.com/nearai/ironclaw/pull/5043) (Jun 17) — Fail‑fast on 400 invalid model.
  - [#5045](https://github.com/nearai/ironclaw/pull/5045) (Jun 17) — `auto` model resolution.
  - [#5132](https://github.com/nearai/ironclaw/pull/5132) (Jun 22) — Invalid chat route redirect.
  - [#5156](https://github.com/nearai/ironclaw/pull/5156) (Jun 23) — Skill‑learning approval gate.
  - [#5160](https://github.com/nearai/ironclaw/pull/5160) (Jun 23) — WebChat v2 SSE tool activity.
  - [#5170](https://github.com/nearai/ironclaw/pull/5170) (Jun 23) — Subagent spawn run failure.

These items may benefit from maintainer triage to avoid contributor stagnation.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-07-04

## 1. Today's Overview

The project showed moderate activity in the last 24 hours, with one new release (2026.7.3) and five pull requests merged, indicating a continued focus on cowork features and stability fixes. One open issue and one open PR remain stale from April, suggesting some long-standing community concerns that have not yet been addressed. Overall, the project is actively shipping improvements, particularly in the areas of service deployment, goal-mode cowork, and context handling.

## 2. Releases

**New version: [LobsterAI 2026.7.3](https://github.com/netease-youdao/LobsterAI/releases/tag/2026.7.3) (2026-07-03)**

Highlights from the release:
- **feat: service deployment** — enables deployment of services directly from the application.
- **feat(cowork): add goal mode** — introduces a goal-oriented mode for collaborative sessions.
- **feat(cowork): add subagent artifact panel** — new UI panel for managing sub-agent artifacts within cowork sessions.

No breaking changes or migration notes were mentioned in the changelog. Users are expected to upgrade normally.

## 3. Project Progress

Five pull requests were merged/closed in the last 24 hours, advancing both features and fixes:

| PR | Description | Status |
|----|-------------|--------|
| [#2270](https://github.com/netease-youdao/LobsterAI/pull/2270) | `chore(release): merge release/2026.7.1 into main` — brings in cowork goal mode, OpenClaw RPC, prompt tool entry points, and recovery improvements | Merged |
| [#2269](https://github.com/netease-youdao/LobsterAI/pull/2269) | `chore: add tooltip to create agent button & guide users to auth requirement when toggling disabled provider` — improves UI usability | Merged |
| [#2268](https://github.com/netease-youdao/LobsterAI/pull/2268) | `fix(cowork): restore compact add menu width` — reverts an unintended width change after removing goal helper text | Merged |
| [#2267](https://github.com/netease-youdao/LobsterAI/pull/2267) | `fix(cowork): sync channel session model override from OpenClaw gateway` — ensures model switches made outside the app are reflected in IM/channel sessions | Merged |
| [#2266](https://github.com/netease-youdao/LobsterAI/pull/2266) | `fix(cowork): clear context maintenance on chat errors` — prevents UI from getting stuck in a “context整理/压缩” state after LLM timeouts | Merged |

These merges indicate that the team is actively polishing the cowork experience, fixing synchronization bugs, and addressing UI regressions.

## 4. Community Hot Topics

Only one open issue and one open PR were updated in the last 24 hours, both of which are stale (created on April 2, 2026):

- **[Issue #1352](https://github.com/netease-youdao/LobsterAI/issues/1352)** — *“任务对话框，任务运行中，附件无法上传（点击上传附件无反应）”*  
  A user reports that file attachment upload fails during an active task (clicking upload has no effect). Has 1 comment, no reactions. The lack of recent responses suggests this might be a low-priority or unconfirmed bug.

- **[PR #1350](https://github.com/netease-youdao/LobsterAI/pull/1350)** — *“skills文件长时间生成阻塞无法感知，中间态过程无展示，用户无法进行下一步…”*  
  This open PR describes multiple user experience pain points: skills generation blocks with no feedback, no intermediate state display, and model understanding issues compared to OpenClaw. It has 0 comments from maintainers.

While neither item shows high engagement, they highlight two persistent community concerns: **file upload reliability** and **lack of progress visibility during long-running operations**.

## 5. Bugs & Stability

The following bugs were addressed or reported today, ranked by severity:

- **High severity – Fix merged**: [PR #2266](https://github.com/netease-youdao/LobsterAI/pull/2266) fixes a cowork bug where chat errors could leave the UI stuck in a “context maintenance” state, preventing further interaction. This is a notable stability improvement.
- **Medium severity – Fix merged**: [PR #2267](https://github.com/netease-youdao/LobsterAI/pull/2267) resolves a model override synchronization issue between OpenClaw gateway and IM/channel sessions, preventing divergence between local and server-side model selections.
- **Low severity – Fix merged**: [PR #2268](https://github.com/netease-youdao/LobsterAI/pull/2268) restores the compact width of the prompt add menu, fixing a minor UI regression.

**Unresolved bug** – [Issue #1352](https://github.com/netease-youdao/LobsterAI/issues/1352) (attachment upload not working during task execution) remains open without a fix PR. Its severity is unclear (no reproduction steps beyond screenshot), but it directly affects task workflow.

## 6. Feature Requests & Roadmap Signals

The most significant feature signal from today’s release is the addition of **cowork goal mode** and **subagent artifact panel**. These suggest the team is investing in structured, goal-driven collaboration workflows.

The open [PR #1350](https://github.com/netease-youdao/LobsterAI/pull/1350) implicitly requests:
- **Progress indicators / intermediate state display** for long-running skills generation.
- **Better model understanding consistency** across different contexts (user notes that the same prompt works better in OpenClaw).

Given the team’s focus on cowork improvements, a progress-indication feature (e.g., streaming status or task breakdown) could appear in the next few releases.

## 7. User Feedback Summary

User feedback from the past 24 hours is limited to the two stale items, but they reveal real pain points:

- **Frustration with opaque long-running operations** – In PR #1350, a user explicitly states *“无中间思考过程态，无法知道龙虾到底是否在操作，用户感知不到下一步的动作和问题”* (no intermediate thought process, can’t tell if the system is working, no visibility into next steps). This indicates dissatisfaction with the current feedback model.
- **Attachment upload failure** – Issue #1352 shows that task-mode upload is broken, which blocks a core use case (e.g., sending files during an active task).
- **Comparisons to alternative tools** – The same user notes that OpenClaw with the same model handles skills generation correctly, implying users are evaluating LobsterAI against other platforms.

Overall satisfaction is mixed: the project is shipping valuable cowork features, but fundamental UX issues like progress visibility and basic file upload reliability remain unaddressed.

## 8. Backlog Watch

Two items have been open since **April 2, 2026** (over 3 months) and have received no maintainer comments or updates:

1. **[Issue #1352](https://github.com/netease-youdao/LobsterAI/issues/1352)** — *Attachment upload unresponsive during task execution.*  
   Low engagement but potentially blocking for users who rely on file sharing in tasks.

2. **[PR #1350](https://github.com/netease-youdao/LobsterAI/pull/1350)** — *Skills generation blocks with no progress feedback; model understanding inconsistent.*  
   This is a user-submitted PR that combines a bug report and feature request. Its lack of response suggests the maintainers have not yet triaged it or it may be considered low priority.

Both items would benefit from a maintainer acknowledgement or a status update to manage community expectations.

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

# CoPaw Project Digest – 2026-07-04

*Based on activity from `agentscope-ai/QwenPaw` (CoPaw) – a personal AI agent and assistant platform.*

---

## 1. Today's Overview

The project saw sustained high activity with **17 PRs updated** (11 open, 6 merged/closed) and **7 issues updated** (6 open, 1 closed). No new releases were cut today, but the development branch is moving rapidly toward the anticipated **QwenPaw 2.0**. The merged PRs bring concrete improvements in memory search (configurable reranker), Windows native sandbox support, GitHub Models endpoint migration, and request timeout/retry infrastructure. Meanwhile, the open PR queue includes several first-time-contributor fixes, a new Azure Bot channel, and a Tauri-based desktop release pipeline. The issue tracker shows a healthy mix of user-reported bugs, architectural discussions, and community encouragement for the upcoming v2.0.

---

## 2. Releases

*No releases were published today.*

---

## 3. Project Progress – Merged / Closed PRs Today

Six pull requests were merged or closed in the last 24 hours:

| PR | Title | Contributor | Significance |
|---|---|---|---|
| [#5648](https://github.com/agentscope-ai/QwenPaw/pull/5648) | feat(memory): add configurable reranker for memory search | `lecheng2018` | Adds reranking (e.g. SiliconFlow) to hybrid vector+BM25 search, improving memory retrieval quality. |
| [#5647](https://github.com/agentscope-ai/QwenPaw/pull/5647) | feat(memory): add reranker config panel to memory settings | `lecheng2018` | Frontend counterpart – collapsible UI for enabling/disabling reranker and configuring API credentials. |
| [#5525](https://github.com/agentscope-ai/QwenPaw/pull/5525) | feat(sandbox): implement windows native sandbox | `ustc-mkh` *(first-time)* | Community contribution bringing native Windows sandbox isolation for agent code execution. |
| [#5735](https://github.com/agentscope-ai/QwenPaw/pull/5735) | fix(providers): update GitHub Models to new endpoint and support fine-grained PAT | `wangfei010313` | Migrates from deprecated Azure-hosted endpoint to `models.github.ai/inference` and adds fine-grained token support. |
| [#5764](https://github.com/agentscope-ai/QwenPaw/pull/5764) | feat: add request timeout, retry and AbortSignal support | `zhaozhuang521` | Improves robustness of outbound HTTP calls with configurable timeout (default 30s), retries, and proper abort handling. |
| [#5506](https://github.com/agentscope-ai/QwenPaw/pull/5506) | fix: sync execution_level to policy.yaml on frontend policy update | `chenzhengcai` *(first-time)* | Fixes tool execution policy not persisting to `policy.yaml` in v2.0, and corrects `off` behavior to truly skip approval. |

**Notable:** Three first-time contributors merged today – a strong sign of growing community engagement.

---

## 4. Community Hot Topics

The most active discussions (by comments, reactions, or underlying architectural weight):

- **[Issue #5767](https://github.com/agentscope-ai/QwenPaw/issues/5767)** – *Console session/message layer limited by single-session pull model* (2 comments, 1 👍)  
  A deep architectural issue: the underlying `@agentscope-ai/chat` SDK only supports pulling one session at a time, blocking multi-agent and multi-workspace evolution. User `Zedthm` describes concrete blocking scenarios. This is a **key bottleneck** for v2.0’s scalability goals.

- **[Issue #5710](https://github.com/agentscope-ai/QwenPaw/issues/5710)** – *Context compression has no protected anchors* (2 comments)  
  Submitted by `ZhaoX666` with detailed impact analysis: critical messages (channel identity, task instructions, board content) are evicted during compression, causing agents to lose situational awareness. Linked to real-world failures in Feishu group chats.

- **[Issue #5770](https://github.com/agentscope-ai/QwenPaw/issues/5770)** – *Hope v2.0 will amaze everyone!* (2 comments)  
  While not a technical bug, this issue by `vipcys001-bot` reflects **strong community anticipation** for the next major release. The two responses show that users are actively waiting.

- **[Issue #5769](https://github.com/agentscope-ai/QwenPaw/issues/5769)** – *Double /api prefix causes 404 for workspace commands* (1 comment)  
  A regression in the 2.0.0b2 console build that breaks workspace command endpoint resolution.

- **[PR #5598 / #5597](https://github.com/agentscope-ai/QwenPaw/pull/5598)** – *LLM fallback configuration UI & backend* (updated today, 0 comments)  
  These paired PRs introduce per-agent and global LLM model fallback with safe retry boundaries – a feature that addresses a common user pain point (model failures). No public discussion yet, but the implementation is substantial.

---

## 5. Bugs & Stability

**Critical severity:**
- **[Issue #5769](https://github.com/agentscope-ai/QwenPaw/issues/5769)** – Double `/api` prefix in v2.0.0b2 console assets causes 404 on workspace command API. Blocks frontend functionality on the beta. No fix PR open yet.

**High severity:**
- **[Issue #5772](https://github.com/agentscope-ai/QwenPaw/issues/5772)** – `_is_bad_request_or_media_error()` treats all HTTP 400 as media rejection, poisoning capability cache when switching LM Studio models. After a model switch, all subsequent image messages get silently stripped. User `quanrennsxsb` reports an easy reproducer with LM Studio. No existing fix PR.
- **[Issue #5710](https://github.com/agentscope-ai/QwenPaw/issues/5710)** – Context compression lacks anchor protection – critical messages get evicted, causing Agent to lose channel awareness or task instructions. A proposed fix is in **[PR #5765](https://github.com/agentscope-ai/QwenPaw/pull/5765)** (fix scroll strategy to protect active turn), which supersedes an earlier patch. This PR also addresses a related scroll strategy eviction bug.

**Low severity:**
- **[Issue #5771](https://github.com/agentscope-ai/QwenPaw/issues/5771)** – Debug logs in `model_factory.py` incorrectly use WARNING level, causing log spam. Trivial to fix.
- **[Issue #5456](https://github.com/agentscope-ai/QwenPaw/issues/5456)** – *(Now closed)* Wrong agent identity for channel-built requests. Likely fixed by an earlier PR.

**Regression note:** The console 404 (Issue #5769) is a clear blocker for anyone testing the 2.0.0b2 release.

---

## 6. Feature Requests & Roadmap Signals

The following features, visible in open PRs and discussed in issues, point to the likely direction for QwenPaw 2.0 and beyond:

| Feature | PR / Issue | Expected Impact |
|---|---|---|
| **LLM model fallback** (UI + backend) | [#5598](https://github.com/agentscope-ai/QwenPaw/pull/5598), [#5597](https://github.com/agentscope-ai/QwenPaw/pull/5597) | Resilience when the primary model fails – a top request for production use. |
| **Memory reranker** (configurable, SiliconFlow) | [#5648](https://github.com/agentscope-ai/QwenPaw/pull/5648), [#5647](https://github.com/agentscope-ai/QwenPaw/pull/5647) | Better relevance ranking in memory search. |
| **Azure Bot channel** | [#5762](https://github.com/agentscope-ai/QwenPaw/pull/5762) | Connects QwenPaw to Teams, Slack, Web Chat, Telegram, etc. via a single channel. |
| **Tauri desktop release** | [#5734](https://github.com/agentscope-ai/QwenPaw/pull/5734) | New cross-platform desktop build pipeline, replacing legacy conda-pack. |
| **AI code review bot** | [#5736](https://github.com/agentscope-ai/QwenPaw/pull/5736) | Automated PR review using QwenPaw itself – dogfooding and developer tooling. |
| **Per-request model override** | [#5731](https://github.com/agentscope-ai/QwenPaw/pull/5731) | Allows cron jobs to specify a different model without changing the agent's persisted model. |
| **Windows native sandbox** | [#5525](https://github.com/agentscope-ai/QwenPaw/pull/5525) (merged) | Full Windows sandbox support for agent execution. |

**Prediction for next release (likely v2.0):** LLM fallback, memory reranker, Tauri desktop, and the Azure Bot channel appear close to merging. The context compression fix (PR #5765) will likely land soon to resolve a high-severity bug.

---

## 7. User Feedback Summary

**Pain points expressed in the last 24 hours:**
- Context compression destroying agent awareness (#5710) – “Agent forgets which channel it is in” – especially painful for multi-channel deployments.
- Console SDK architecture blocking multi-agent workflows (#5767) – user `Zedthm` explicitly says it’s blocking evolution.
- Model switching with LM Studio causes image input to silently fail (#5772) – confusing UX.
- Debug log spam (#5771) – noisy, not dangerous.
- 404 on beta console (#5769) – immediate usability regression for beta testers.

**Satisfaction signals:**
- Issue #5770 explicitly cheers for v2.0: “希望V2.0的正式版推出之后，能够惊艳所有人！” (“Hope v2.0 will amaze everyone!”) – shows strong positive anticipation.
- Several PRs from first-time contributors indicate the project is welcoming and the docs/channels are effective.

**Underlying need:** Users want a more robust multi-context, multi-channel, multi-model system. The recurring theme is reliability under real-world conditions (fallback, compression anchors, channel identity).

---

## 8. Backlog Watch

No issues or PRs stand out as long-unanswered with stale maintainer attention. The most active items are recent. However, one item worth monitoring:

- **[PR #5514](https://github.com/agentscope-ai/QwenPaw/pull/5514)** – *fix chat input queue session id migration* – Opened June 25, last updated July 3. This PR touches the SDK integration layer and is blocked on a dependency upgrade (`@agentscope-ai/chat` to `1.1.71-beta`). It’s been open over a week; maintainers may need to prioritize reviewing the SDK upgrade path to unblock it.

All other open items have received recent updates or are actively being discussed.

---

*Generated from `agentscope-ai/QwenPaw` GitHub data, snapshot 2026-07-04.*

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest — 2026-07-04

## Today's Overview

Activity remains high, with **13 issues** and **50 pull requests** updated in the past 24 hours. The project closed **3 issues** and merged/closed **6 PRs**, while 10 issues and 44 PRs remain active. No new releases were published today. The team focused on fixing critical bugs—especially around MCP tools not appearing in TUI sessions and OpenAI‑compatible provider serialization—while advancing large feature branches for SOP authoring, gateway OpenAI‑compatible endpoints, and runtime OTel content policies. Several RFCs and trackers signal active architectural planning for a `.ignore` file mechanism, goal‑mode implementation split, and ADR restoration.

## Releases

No new releases were published today. The last release remains unknown from the provided data.

## Project Progress

**Closed issues** (bugs fixed):
- [#8193](https://github.com/zeroclaw-labs/zeroclaw/issues/8193) – MCP tools missing from TUI sessions (high severity, fixed via PR #8634).
- [#7862](https://github.com/zeroclaw-labs/zeroclaw/issues/7862) – OpenAI‑compat providers sending `tool_choice` with empty tools list → vLLM 400 error (medium severity, fixed via PR #8667).
- [#8632](https://github.com/zeroclaw-labs/zeroclaw/issues/8632) – Source install with `embedded-web` failing before generated API client exists (blocked workflow, fixed via PR #8643).

**Merged/closed PRs today** (6 total):
- [#8596](https://github.com/zeroclaw-labs/zeroclaw/pull/8596) – `fix(channels)`: carry WeCom reply scope as structured metadata (L).
- [#8667](https://github.com/zeroclaw-labs/zeroclaw/pull/8667) – `fix(providers)`: omit `tool_choice` for empty tools on Responses API path (XS).
- [#8665](https://github.com/zeroclaw-labs/zeroclaw/pull/8665) – `fix(windows)`: silence config shell clippy warnings (XS).
- [#8643](https://github.com/zeroclaw-labs/zeroclaw/pull/8643) – `fix(install)`: prebuild dashboard for embedded web (XS).
- [#8634](https://github.com/zeroclaw-labs/zeroclaw/pull/8634) – `fix(zerocode)`: advertise deferred MCP tools in Chat TUI system prompt (XS).

These fixes resolved the two S1 (workflow‑blocked) bugs reported in recent days and cleaned up Windows platform support.

## Community Hot Topics

**Most active issues** (by comment count):
- [#8193](https://github.com/zeroclaw-labs/zeroclaw/issues/8193) – **MCP tools missing from TUI** (15 comments) – now closed, but generated high engagement as users reported workflow blockage. The discussion identified root cause as the `Agent::from_config_with_tui_env` path not building the `tool_search` tool.
- [#8424](https://github.com/zeroclaw-labs/zeroclaw/issues/8424) – **RFC: .ignore File Mechanism** (7 comments) – users request a way to protect workspace‑internal sensitive files from AI agent access. Currently blocked awaiting author action.
- [#8681](https://github.com/zeroclaw-labs/zeroclaw/issues/8681) – **Tracker: Goal mode implementation split stack** (7 comments) – coordinates splitting an already‑implemented goal‑mode feature into reviewable PRs. Signals a large incoming feature.

**Most active PRs** (comment counts not provided, but size and label indicate heavy review):
- [#8590](https://github.com/zeroclaw-labs/zeroclaw/pull/8590) – **feat(sop): visual SOP authoring surfaces** (XL, 20+ labels) – calls for beta testers; likely a community‑facing milestone.
- [#8384](https://github.com/zeroclaw-labs/zeroclaw/pull/8384) – **feat(inkbox): native Inkbox channel** (XL) – adds email, SMS, voice, iMessage channel with Quickstart onboarding.
- [#8486](https://github.com/zeroclaw-labs/zeroclaw/pull/8486) – **feat(gateway): OpenAI chat completions endpoint** (XL) – seeks to enable compatibility with LangChain, Continue.dev, etc. Currently blocked on author action.

## Bugs & Stability

**New bugs opened today** (4 issues):
- [#8695](https://github.com/zeroclaw-labs/zeroclaw/issues/8695) – **Cron jobs still recall memory when `uses_memory = false`** (S2 – degraded behavior). Memory context is not fully suppressed. No fix PR yet.
- [#8693](https://github.com/zeroclaw-labs/zeroclaw/issues/8693) – **ZeroCode /model picker shows switched model but provider still uses pinned configured model** (S2). Inconsistency between UI and actual dispatch. No fix PR yet.
- [#8646](https://github.com/zeroclaw-labs/zeroclaw/issues/8646) – **ZeroCode Logs detail hiding event attributes behind preview-only rows** (S2). Log inspection degraded.
- [#8664](https://github.com/zeroclaw-labs/zeroclaw/issues/8664) – **ZeroCode code-block Copy includes Markdown fences** (S2). Usability issue.

**Previously reported high‑priority bugs still open**:
- [#5869](https://github.com/zeroclaw-labs/zeroclaw/issues/5869) – **Security: rumqttc v0.25.1 pins vulnerable transitive dependencies** (P1, blocked). Four RUSTSEC advisories. No progress since April.

**Stability note**: The three bugs closed today (#8193, #7862, #8632) were all S1 (workflow‑blocked). Their resolution significantly improves the developer and user experience for MCP tool use and OpenAI‑compatible provider integrations.

## Feature Requests & Roadmap Signals

**Most notable RFCs and trackers**:
- [#8424](https://github.com/zeroclaw-labs/zeroclaw/issues/8424) – **.ignore File Mechanism** – high demand for workspace file protection. Likely to be accepted once author responds.
- [#8681](https://github.com/zeroclaw-labs/zeroclaw/issues/8681) – **Goal mode implementation split stack** – indicates a new agent behavior mode is nearing review.
- [#8070](https://github.com/zeroclaw-labs/zeroclaw/issues/8070) – **Tracker: v0.8.3 gateway, web, ZeroCode, and onboarding** – umbrella for the next minor release. Includes web dashboard, desktop/Tauri work.
- [#8692](https://github.com/zeroclaw-labs/zeroclaw/issues/8692) – **Tracker: Active RFC review queue** – makes RFC review state visible; suggests maintainers are systematically processing proposals.
- [#8691](https://github.com/zeroclaw-labs/zeroclaw/issues/8691) – **Tracker: Restore ADR baseline** – aims to audit and restore architecture decision records. PR #8694 already submitted for ADR-001 and ADR-002.

**Prediction**: The v0.8.3 release (tracked in #8070) will likely include the OpenAI chat completions endpoint (PR #8486), visual SOP authoring (PR #8590), and the Inkbox channel (PR #8384), alongside the hotfixes closed today. Goal mode may land shortly after.

## User Feedback Summary

**Common pain points expressed**:
- **MCP tool visibility** (#8193): users reported that MCP servers connect but tools never appear in TUI sessions. Fixed today.
- **Log inspection** (#8646): selection of events only shows preview fields, hiding full payloads.
- **Code block copy** (#8664): copying Markdown fences and highlighted messages is confusing.
- **Model picker inconsistency** (#8693): UI shows model change but actual dispatch uses old pinned model.
- **Cron memory leak** (#8695): scheduled tasks ignore the `uses_memory=false` flag.

**Satisfaction signals**:
- The quick closure of S1 bugs (MCP, vLLM 400, install failure) demonstrates responsive maintainer engagement.
- Large feature PRs like SOP authoring (#8590) are explicitly calling for beta testers, indicating community involvement is welcomed.

## Backlog Watch

**Long‑standing open issues needing maintainer attention**:
- [#5869](https://github.com/zeroclaw-labs/zeroclaw/issues/5869) – **Security: rumqttc dependency** – open since April, blocked on `rumqttc` upstream. No update in months. Risk high.
- [#8424](https://github.com/zeroclaw-labs/zeroclaw/issues/8424) – **RFC: .ignore File Mechanism** – blocked on author action; maintainer review pending.
- [#8070](https://github.com/zeroclaw-labs/zeroclaw/issues/8070) – **v0.8.3 tracker** – no activity since creation, though sub‑PRs are moving.

**PRs needing author action** (blocked):
- [#8486](https://github.com/zeroclaw-labs/zeroclaw/pull/8486) – OpenAI chat completions endpoint; labelled `needs-author-action`.
- [#8384](https://github.com/zeroclaw-labs/zeroclaw/pull/8384) – Inkbox channel; labelled `needs-author-action`.

**Tracker with low activity**:
- [#8692](https://github.com/zeroclaw-labs/zeroclaw/issues/8692) – RFC review queue – created today; will need regular updates to serve its purpose.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*