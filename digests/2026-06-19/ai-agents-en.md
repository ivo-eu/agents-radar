# OpenClaw Ecosystem Digest 2026-06-19

> Issues: 104 | PRs: 500 | Projects covered: 13 | Generated: 2026-06-19 12:58 UTC

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

# OpenClaw Project Digest — 2026-06-19

## Today's Overview
Project activity remains very high with 500 pull requests and 104 issues updated in the last 24 hours, and 32 PRs merged or closed. One new beta release (v2026.6.9-beta.1) shipped, focusing on richer Telegram delivery performance. The community is actively reporting regressions and stability issues from the v2026.6.1 stable release, with several P0/P1 bugs still awaiting fixes. Overall project health is strong but under strain from migration-related defects and connector-specific deadlocks.

## Releases
**v2026.6.9-beta.1** ([release notes](https://github.com/openclaw/openclaw/releases/tag/v2026.6.9-beta.1))
- Highlights: Telegram delivery now sends rich HTML, preserves rich markdown and sticker paths, renders progress drafts and command output faithfully, and keeps mentions and spooled handlers on the correct delivery path.
- This is a beta release; no breaking changes or migration notes were published. Use with caution in production deployments.

## Project Progress
While 32 PRs were merged or closed today, the top 30 open PRs by comment count give a sense of active work:
- **[#94881](https://github.com/openclaw/openclaw/pull/94881)** (XL) – Slack global and message shortcuts handled; pending proof.
- **[#78536](https://github.com/openclaw/openclaw/pull/78536)** (XL) – Slack thread reply buffering fix; ready for maintainer review.
- **[#90500](https://github.com/openclaw/openclaw/pull/90500)** (L) – Stale session routes for removed providers; ready for review.
- **[#94828](https://github.com/openclaw/openclaw/pull/94828)** (M) – New `openclaw skills uninstall` CLI command; waiting on author.
- **[#90855](https://github.com/openclaw/openclaw/pull/90855)** (L) – Recover stale pending final deliveries; waiting for proof.

The merge pace suggests sustained progress on core connectivity, CLI, and Slack/Matrix integrations.

## Community Hot Topics
Most active issues by comment count (6 comments each):
- **[#90595 – Cron "failed" notifications fire during hot reload and retries](https://github.com/openclaw/openclaw/issues/90595)** (P2, diamond lobster) – Alert fatigue from noisy cron failure reporting. Users want better deduplication.
- **[#90378 – Cron store silently migrated to SQLite, new jobs default to announce mode causing errors](https://github.com/openclaw/openclaw/issues/90378)** (P0, diamond lobster) – Silent data migration from JSON to SQLite during upgrade breaks cron workflows.
- **[#90213 – Legacy state migration warnings persist even after `doctor --fix`](https://github.com/openclaw/openclaw/issues/90213)** (P2, diamond lobster) – Upgrade process not idempotent; users frustrated by repeated warnings.

Other high-reaction issues:
- **[#90361](https://github.com/openclaw/openclaw/issues/90361)** (3 👍) – Intermittent `memory_search "index metadata is missing"` despite valid index; likely search/reindex race.
- **[#90082](https://github.com/openclaw/openclaw/issues/90082)** (2 👍) – Active-memory circuit breaker too aggressive, polluting main session with fallback prompt.

Underlying need: Users are encountering systemic issues with the new memory indexing and migration logic introduced in v2026.6.1. The community desires more robust migration testing and clearer error messaging.

## Bugs & Stability
**P0 / Critical:**
- **[#90559](https://github.com/openclaw/openclaw/issues/90559)** – Feishu connector: sequential queue deadlock + self-talk feedback loop under slow LLM. No fix PR yet. Causes complete session loss.
- **[#90378](https://github.com/openclaw/openclaw/issues/90378)** – Silent cron store migration leads to channel errors (linked PR open).
- **[#89994](https://github.com/openclaw/openclaw/issues/89994)** – Fuzzy edit silently normalizes and rewrites unrelated lines across whole file (linked PR open).

**P1 / High Severity:**
- **[#90840](https://github.com/openclaw/openclaw/issues/90840)** – Subagent run completion delivered as raw worker output instead of parent summary (linked PR open).
- **[#90588](https://github.com/openclaw/openclaw/issues/90588)** – All QQ Bot agents unresponsive: `Cannot read properties of undefined (reading 'run')` (no fix PR).
- **[#89374](https://github.com/openclaw/openclaw/issues/89374)** – Timeout compaction leaves Codex channel session unrecoverable (no fix PR).
- **[#90082](https://github.com/openclaw/openclaw/issues/90082)** – Active-memory circuit breaker too aggressive (no fix PR).
- **[#90536](https://github.com/openclaw/openclaw/issues/90536)** – OpenAI OAuth missing `model.request` scope – GPT-5.5 falls back silently (no fix PR).

Many of these regressions stem from the v2026.6.1 migration. Several have linked PRs, but over half lack a fix in progress, indicating maintainer bottleneck.

## Feature Requests & Roadmap Signals
- **[#89871](https://github.com/openclaw/openclaw/issues/89871)** – `replyInThread` channel config for Slack & Discord (P3). Likely to be implemented given strong community need for thread management.
- **[#90442](https://github.com/openclaw/openclaw/issues/90442)** – Support Slack-native Block Kit for structured cron reports (P3). Could be fast-tracked for enterprise users.
- **[#90818](https://github.com/openclaw/openclaw/issues/90818)** – Expose `agentId/agentName` in plugin SDK `register()` (P2). Small change, high value for multi-agent setups; likely next minor release.
- **[#90732](https://github.com/openclaw/openclaw/issues/90732)** – Rename "Steer now" button to non-technical language (P3). Simple UX fix that could come in next patch.
- **[#90608](https://github.com/openclaw/openclaw/issues/90608)** – Configurable default locale for interface (P3). Low effort, high satisfaction for international users.

**Prediction for next release (v2026.6.9 stable):** The beta release focuses on Telegram delivery; the stable will likely include fixes for the most critical regressions reported last week, plus the `skills uninstall` command and Slack shortcut support.

## User Feedback Summary
- **Pain points:** Silent data migration (JSON→SQLite), persistent "index metadata is missing" errors, connector deadlocks (Feishu, QQ Bot), confusing "Steer now" terminology, and non-idempotent upgrade commands.
- **Use cases:** Multi-agent teams, cron-based reporting, rich media delivery (images in Telegrams, Slack threads), and CPU-only VPS deployments with QMD backend.
- **Satisfaction:** Users appreciate the new Telegram delivery improvements in the beta, but overall satisfaction is dented by the number of regressions introduced in v2026.6.1. The “diamond lobster” issue rating on many reports indicates high community engagement and frustration.
- **Quotes:** "I hate [the German interface]" (#90608) and "not enough info" on token cost regression (#90170) show a mix of emotional and data-driven feedback.

## Backlog Watch
Several high-severity issues lack maintainer activity or linked fix PRs:
- **[#66360](https://github.com/openclaw/openclaw/issues/66360)** (P1, diamond lobster, opened 2026-04-14) – Session maintenance has no size cap for transcript `.jsonl` files, causing CPU 100% crashes. Needs product decision.
- **[#90588](https://github.com/openclaw/openclaw/issues/90588)** (P1, platinum hermit, no fix PR) – QQ Bot agents unresponsive after v2026.5.28→v2026.6.1 upgrade. Critical for Chinese-market users.
- **[#90595](https://github.com/openclaw/openclaw/issues/90595)** (P2, diamond lobster, no new fix PR) – Cron alert fatigue. Maintainer review needed to decide deduplication approach.
- **[#89926](https://github.com/openclaw/openclaw/issues/89926)** (P2, platinum hermit, no fix PR) – OpenCode ACP runtime spawn fails 40% when wrapped by npx. Live repro needed.
- **[#90254](https://github.com/openclaw/openclaw/issues/90254)** (P2, platinum hermit, linked PR open) – Telegram webhook continuously wiped by long-polling transport. Needs fix merge.

These issues should be prioritized to improve stability and user trust.

---

## Cross-Ecosystem Comparison

好的，作为一名资深分析师，以下是根据您提供的2026年6月19日社区摘要生成的跨项目比较报告。

***

### 跨项目对比报告：个人AI智能体/助手开源生态系统 (2026-06-19)

本报告对核心参考实现 **OpenClaw** 及其同生态相关项目（NanoBot， Hermes Agent， PicoClaw， NanoClaw, NullClaw, IronClaw， LobsterAI, Moltis, CoPaw, ZeroClaw）在2026年6月19日的活动状态、社区焦点和技术趋势进行了横向比较分析，旨在为技术决策者和开发者提供数据驱动的洞察。

#### 1. 生态系统概述

当前，个人AI智能体/助手开源生态系统正处于一个从“能力堆叠”向“稳定可靠”过渡的关键时期。以 `*Claw` 家族为代表的多个项目围绕“连接器-代理-记忆-工具”的核心架构进行竞争与分化。尽管社区贡献活跃，但多个项目都因近期的重大版本更新（尤其是v2026.6.1系列）而面临迁移稳定性、连接器死锁和回归问题的挑战。用户对无缝升级、鲁棒错误处理和云端-边缘混合部署（如移动端、嵌入式设备）的需求尤为突出，驱动着项目从功能创新转向稳定性与兼容性优化。

#### 2. 活动对比

下表总结了各项目在2026年6月19日过去24小时内的核心活动指标。

| 项目 | 问题更新数 (近24h) | PR更新数 (近24h) | 发布状态 (近24h) | 健康度评级 |
| :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 104 | 500 | 1个Beta (v2026.6.9-beta.1) | **Strong (高活跃，但迁移压力大)** |
| **ZeroClaw** | 8 | 50 | 1个Patch (v0.8.1) | **Very Strong (快速增长，大量特性开发)** |
| **Hermes Agent** | 8 | 50 | 无 | **Strong (高活跃，桌面UI和子代理是焦点)** |
| **NanoClaw** | 4 | 19 | 无 | **Strong (社区贡献活跃，安全修复集中)** |
| **NanoBot** | 6 | 38 | 无 | **Strong (问题修复迅速，社区健康)** |
| **CoPaw** | 10 | 11 | 无 | **Moderate (记忆管理和移动端是热点)** |
| **IronClaw** | 7 | 21 | 无 | **Good (响应迅速，OAuth和审批流程改进)** |
| **PicoClaw** | 1 | 12 | 1个Nightly (v0.3.0-nightly) | **Moderate (依赖维护为主，核心Bug待解)** |
| **NullClaw** | 4 | 4 | 无 | **Low (平台兼容性问题拖累进展)** |
| **LobsterAI** | 4 | 0 | 1个Patch (2026.6.18) | **Low (贡献动量低，PR活动为零)** |
| **Moltis** | 1 | 0 | 无 | **Inactive (单日无代码合并，开发疑似暂停)** |
| **TinyClaw/ZeptoClaw** | 0 | 0 | 无 | **Dormant (24小时内无任何活动)** |

**结论**: **OpenClaw** 作为核心参考实现，拥有压倒性的社区规模和活动量，但也暴露出相应的维护瓶颈。**ZeroClaw** 和 **Hermes Agent** 呈现出非常强劲的增长和创新势头。**NanoClaw** 和 **NanoBot** 则体现了稳定的社区贡献和完善的问题修复机制。

#### 3. OpenClaw 的定位

- **优势与社区规模**: OpenClaw 是整个“Claw生态”的绝对核心。其500个PR和104个问题更新的单日数据，是其他任何项目的5-10倍，这反映了其作为“事实标准”的巨大社区基数和开发者注意力。其发布版本（v2026.6.9）是所有项目中最频繁和详细的，具备最全面的连接器支持（如Feishu， QQ Bot）。
- **技术差异**: 相比更轻量的 **NanoBot** 和 **NanoClaw**，OpenClaw 的代码库体量更大，引入了“技能(Skills)”、“内存索引(Memory Indexing)”等更复杂的抽象层。这也是其在v2026.6.1版本迁移中遭遇多项P0/P1回归问题的根本原因。
- **社区对比**: OpenClaw 面临的是“幸福的烦恼”：社区规模过大导致维护者成为瓶颈。许多高优先级Bug（如 #90588, #90595）缺乏明确的修复PR，而 **IronClaw** 和 **NanoBot** 等规模较小的项目则能以天为单位快速响应用户问题。这暗示OpenClaw可能需要在流程（如从社区招募核心维护者）或架构（如模块化）上进行优化。

#### 4. 共享技术焦点领域

多个项目不约而同地将开发重点集中在以下几个领域，显示了行业共识：

- **稳定迁移与状态管理**:
    - **项目**: **OpenClaw， NanoClaw， CoPaw**
    - **需求**: 用户对非侵入式、无故障的版本升级有强烈需求。具体表现为：**OpenClaw** 的无声数据迁移（JSON→SQLite）和非幂等修复命令、**NanoClaw** 的用户对v1→v2迁移中Telegram群组身份状态的不确定性、**CoPaw** 的v1.1.12版本导致的图片显示回归。
- **丰富消息通道与原生体验**:
    - **项目**: **OpenClaw， Hermes Agent， NanoClaw, ZeroClaw, CoPaw**
    - **需求**: 不仅是要“能发消息”，更是要利用各平台原生UI能力。例如：**OpenClaw** 的Telegram富文本/贴纸交付、**Hermes Agent** 的Feishu/Lark富内容可见性、**NanoClaw** 和 **ZeroClaw** 的Discord消息分块和嵌入，以及 **CoPaw** 的Webui实时推送。
- **子代理（Sub-agent）通信与消息去重**:
    - **项目**: **OpenClaw， PicoClaw， CoPaw**
    - **需求**: 在异步子代理任务完成后，如何优雅地提供结果。**PicoClaw** 遇到了明确的“重复消息”Bug（#3094），即子代理的原始输出和父代理的摘要会同时推送给用户。**OpenClaw** 也有子代理原始输出而非摘要的Bug。这表明子代理结果的生命周期管理是一个普遍性痛点。
- **安全性与沙箱**:
    - **项目**: **NanoClaw, NullClaw， CoPaw**
    - **需求**: 保护主机环境免受恶意或错误代理行为的侵害。**NanoClaw** 重点修复了文件读取未限制工作区和路径遍历的问题，**CoPav** 新增了 `bubblewrap` Linux沙箱。**NullClaw** 则聚焦于Android平台下的权限和安全问题。

#### 5. 差异化分析

| 维度 | OpenClaw | NanoClaw | Hermes Agent | ZeroClaw | CoPaw |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **目标用户** | 技术发烧友（任何开发者） | 系统管理员/DevOps | 前端/全栈开发者 | 性能优先的高级用户 | 注重记忆和长期任务的团队 |
| **核心特色** | 「参考实现」，生态最全，社区最大 | 安全优先，CLI和CI/CD流程强大 | 桌面UI体验，子代理和Copilot深度整合 | 性能极致，架构革新（WASM, 路由） | 强大的记忆与上下文管理，移动端适配 |
| **技术架构** | 模块化，社区驱动 | 可插拔，注重环境变量 | Rust主架构，Electron桌面壳 | Rust主架构，试验性WASM插件 | Python，Tauri桌面壳 |
| **沟通焦点** | 稳定性与迁移痛点 | 安全攻击面和CLI缺陷 | 桌面UI和Copilot体验 | 会话持久化和WASM插件 | 记忆崩溃和DeepSeek兼容 |

**结论**: 与OpenClaw相比，**NanoClaw** 更像一个“安全加固版本”，**Hermes Agent** 提供了更现代的桌面和Copilot体验，而 **ZeroClaw** 则在敏捷性和技术前瞻性上领先。

#### 6. 社区势头与成熟度

- **第一梯队: 高速迭代 & 高活跃**
    - **OpenClaw， ZeroClaw, Hermes Agent**: 这些项目处于活动量的顶峰，社区贡献源源不断。**OpenClaw** 是“巨无霸”，但正经历成长的阵痛。**ZeroClaw** 和 **Hermes Agent** 则表现出更健康的快速增长，修复和功能开发并行不悖，且修复周期更短。
- **第二梯队: 稳定递增 & 功能完善**
    - **NanoClaw, NanoBot, IronClaw**: 这些项目活动量虽不如第一梯队，但项目成熟度较高。**IronClaw** 在OAuth和审批流程上响应迅速；**NanoClaw** 在安全修复上极其高效；**NanoBot** 展现出快速修复社区报告的Claw问题的能力。
- **第三梯队: 低活跃或停滞**
    - **PicoClaw， NullClaw， LobsterAI, Moltis**: 这些项目或因核心Bug（如PicoClaw的子代理消息重复、NullClaw的Android构建失败）而停滞，或因贡献动量不足（LobsterAI零PR）而呈现“假死”状态。**Moltis** 的单日无活动尤其值得注意，可能表明项目已暂停开发或重构中。

#### 7. 趋势信号

从社区反馈和开发重点中，可以提取出以下关键趋势：

1.  **“无摩擦”升级是基本盘**: 用户不再容忍破坏性迁移（如OpenClaw的无声数据库迁移）。任何需要手动修复脚本或丢失状态的版本都是不可接受的。自动化、幂等的迁移是使项目进入生产就绪状态的门槛。
2.  **从“可用”到“好用”的连接器**: 开发者对连接器（Telegram， Slack, Discord， Feishu）的要求不仅是基础收发，而是要求原生排版（富文本， 卡片， 嵌入）、原生命令（Slack斜杠命令、Telegram菜单）和无损媒体支持。
3.  **从“单兵作战”到“多代理协作”: ** 子代理（Sub-agent/Spaw）的普及带来了新的复杂性：消息管理（去重/汇总）、状态管理（会话持久化）和权限管理（代理间审批）。这是未来AI Agent系统必须解决的核心中间件问题。
4.  **安全不再是附加功能，而是架构的一部分**: `NanoClaw` 和 `CoPaw` 的实践表明，安全（文件系统隔离、沙箱、输入验证）正在被嵌入到CLI工具、SDK甚至核心运行时中。
5.  **对“模型中立性”提出了更高要求**: 随着DeepSeek， Zhipu等模型在亚洲市场的流行，**CoPaw** 和 **IronClaw** 凸显了兼容性Bug。一个成功的项目必须主动测试和适应非OpenAI的API（如思考标记， 工具调用格式）。

---

## Peer Project Reports

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

## NanoBot Project Digest — 2026-06-19

### 1. Today’s Overview
The project saw **high activity** with 6 issues updated (3 closed, 3 open) and **38 pull requests updated** (20 merged/closed, 18 open). No new releases were published, but the maintainer merged a batch of older PRs (many from March–April) while also addressing recent bugs and enhancements. Community engagement remains strong, particularly around fallback reliability, concurrency safety, and channel-specific rendering.

---

### 2. Releases
*No new releases today.*

---

### 3. Project Progress — Merged/Closed PRs (20 total)
Notable merged/closed PRs that advanced features or fixed issues:

- **[fix(providers): use non-descriptive placeholder when stripping images](https://github.com/HKUDS/nanobot/pull/4401)** — Prevents LLMs from hallucinating about stripped image content by replacing the placeholder with a generic `[image omitted]` (closes issue #4345).
- **[fix(fallback): treat empty API choices as fallbackable error](https://github.com/HKUDS/nanobot/pull/4288)** — Ensures empty responses from providers like DeepSeek trigger fallback models (closes #4287).
- **[fix: support OpenAI image reference edits](https://github.com/HKUDS/nanobot/pull/4394)** — Routes image editing requests to `/images/edits` with proper multipart handling.
- **[fix(feishu): support reading WebSocket rendered card content](https://github.com/HKUDS/nanobot/pull/4342)** — Fixes structural mismatches in Feishu card parsing.
- **[fix MCP malformed progress notifications](https://github.com/HKUDS/nanobot/pull/4372)** — Filters invalid `notifications/progress` messages causing Pydantic errors (closes #4052).

Also merged: multiple older PRs for Discord (#2655), history consolidation (#2692), workspace restrictions (#2725), SQLite dream backend (#3015), `/clear` command (#3467), token optimization (#3497), Matrix list rendering (#2872), and more.

---

### 4. Community Hot Topics

- **[#4287 — Empty model responses not triggering fallback (CLOSED)](https://github.com/HKUDS/nanobot/issues/4287)** — 2 comments. User reported DeepSeek returning empty `choices` during peak hours. The fix PR (#4288) was merged today, addressing a critical reliability gap.
- **[#4408 — Nanobot.run() per-run hooks concurrency bug (OPEN)](https://github.com/HKUDS/nanobot/issues/4408)** — 1 comment. Reports that `_extra_hooks` is mutated in place under concurrent runs, causing hook clobbering. No fix PR yet; high severity.
- **[#4410 — Cron job unwanted message after upgrade (OPEN)](https://github.com/HKUDS/nanobot/issues/4410)** — User describes that bound cron jobs now always deliver LLM responses, even for “nothing to report” cases. A fix PR [#4412](https://github.com/HKUDS/nanobot/pull/4412) was opened today.

**Underlying needs:** Users demand robust fallback chains, thread-safe hook management, and granular control over cron-job output suppression.

---

### 5. Bugs & Stability

| Issue | Severity | Summary | Fix PR |
|-------|----------|---------|--------|
| [#4408 — Hook concurrency clobbering](https://github.com/HKUDS/nanobot/issues/4408) | **High** – shared mutable state can corrupt per-run hooks | Not yet |
| [#4410 — Unwanted cron messages](https://github.com/HKUDS/nanobot/issues/4410) | **Medium** – spurious notifications after upgrade | [#4412](https://github.com/HKUDS/nanobot/pull/4412) (open) |
| [#4345 — Image-strip fallback hallucination](https://github.com/HKUDS/nanobot/issues/4345) | **Medium** – LLM acts as if it saw an image | [#4401](https://github.com/HKUDS/nanobot/pull/4401) (merged today) |
| [#4287 — Empty choices not falling back](https://github.com/HKUDS/nanobot/issues/4287) | **Critical** – main model failure blocks conversation | [#4288](https://github.com/HKUDS/nanobot/pull/4288) (merged today) |
| [#4052 — MCP progress notifications rejected](https://github.com/HKUDS/nanobot/issues/4052) | **Medium** – Pydantic validation error | [#4372](https://github.com/HKUDS/nanobot/pull/4372) (merged today) |

All three closed bugs now have merged fix PRs; the two open bugs need maintainer attention (especially #4408).

---

### 6. Feature Requests & Roadmap Signals

- **[#4413 — Telegram Bot API 10.1 rich messages](https://github.com/HKUDS/nanobot/issues/4413)** – User requests conversion of standard Markdown to Telegram’s rich message format. Likely to be targeted for next minor release.
- **[PR #4411 — SuspendTurn for async/human-in-the-loop](https://github.com/HKUDS/nanobot/pull/4411)** – Adds a sentinel tool to pause a turn cleanly without calling the model; enables async continuations.
- **[PR #4395 — Improve onboard wizard setup flow](https://github.com/HKUDS/nanobot/pull/4395)** – JetBrains-inspired terminal palette and fallback defaults for non-TTY runs.

**Prediction:** Next minor version (v0.2.2 or v0.3.0) will likely include rich message support, SuspendTurn, and the improved onboarding wizard, building on today’s stability fixes.

---

### 7. User Feedback Summary

- **Pain points**: Frequent fallback failures on empty API responses (now fixed), concurrency issues in multi-tenancy scenarios (#4408), and unexpected cron job output after schema changes (#4410).
- **Use cases**: Telegram bot runtime, cron-driven heartbeat tasks, MCP server integration, Discord channel management.
- **Satisfaction**: Users appreciate rapid bug fixes (e.g., #4287, #4345, #4052 all fixed within days). The community is actively contributing tests and enhancements (#4393, #4412).

---

### 8. Backlog Watch

- **[#4408 — Concurrency-safe hooks](https://github.com/HKUDS/nanobot/issues/4408)** – Opened 2026-06-18, 0 comments. No maintainer response yet. Affects any deployment running concurrent `Nanobot.run()` calls (e.g., multiple webhook triggers). High priority.
- **[#4410 — Cron unwanted messages](https://github.com/HKUDS/nanobot/issues/4410)** – Opened today, fix PR #4412 exists but not yet merged. Monitor for maintainer review.
- **Old PRs closed today** – Many stale PRs (#2655, #2692, #2725, etc.) were closed, indicating maintainer cleanup. No ancient unaddressed issues remain.

Overall, project health is **strong** with active bug fixing and steady feature development.

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent Project Digest — 2026-06-19

## Today's Overview
The project saw high activity with **8 issues** (7 open, 1 closed) and **50 pull requests** (42 open, 8 merged/closed) updated in the last 24 hours. No new releases were published. The community continues to contribute heavily on both bug fixes and new features, with a strong emphasis on Desktop UI polish, provider/copilot integration, and subagent lifecycle management. The volume of open PRs suggests healthy development momentum, though several important Windows-specific bugs and configuration issues remain unresolved.

## Releases
*No new releases published today.*

## Project Progress
**8 PRs were merged or closed** in the last 24 hours. Notable among them:
- **#37482** – *feat(stt): add Deepgram transcription provider*  
  (Closed, [GitHub](https://github.com/nousresearch/hermes-agent/pull/37482))  
  Adds Deepgram as a selectable speech-to-text provider, wired through the transcription registry, CLI/config, and web server. This feature was in review since June 2 and finally merged.

Additional closed PRs are not individually listed in the top 20 by comments, but the overall closed count (8) indicates active project maintenance and integration of community contributions.

## Community Hot Topics
- **Issue #42042** – *Desktop UI: add API key field for Local / custom endpoint*  
  ([GitHub](https://github.com/nousresearch/hermes-agent/issues/42042), 2 comments)  
  The most-discussed issue today. Users want to configure custom OpenAI-compatible endpoints entirely from the UI. The first-run flow already supports auto-discovery for no-key endpoints, but adding a dedicated API key field is the remaining gap. Strong demand from self-hosters and local server users.

- **PR #37482** (closed) also generated conversation during its review phase, though comment counts are not available.

- **Issue #49007** – *Add reorder/move support to fallback provider interactive manager*  
  ([GitHub](https://github.com/nousresearch/hermes-agent/issues/49007))  
  A new feature request filed today with a clear use case (fallback priority). The community describes exactly how the config works and why reordering is needed.

- **Issue #49005** – *Feature request: configurable pre-session tool auto-invocation*  
  ([GitHub](https://github.com/nousresearch/hermes-agent/issues/49005))  
  Another feature request with a well-articulated scenario (MCP servers like MemPalace). Users want automatic memory tool invocations at the start of every session.

## Bugs & Stability
Several bugs were reported or actively fixed today, ranked by severity (P2 = medium, P3 = low/moderate):

| Severity | Issue | Description | Fix PR? |
|----------|-------|-------------|---------|
| **P2** | [#48808 (CLOSED)](https://github.com/nousresearch/hermes-agent/issues/48808) | Windows Desktop: clicking notification shows "session not found" – likely a navigation/state bug in the Electron shell. | Closed (fix presumably merged) |
| **P2** | [#46759 (OPEN)](https://github.com/nousresearch/hermes-agent/issues/46759) | Desktop model switch leaves stale `base_url` in config.yaml, breaking subsequent API calls when switching providers. | No PR linked yet |
| **P2** | [#49014 (OPEN)](https://github.com/nousresearch/hermes-agent/issues/49014) | `ssl_guard` throws `NotImplementedError` on Windows due to `truststore` overriding `ssl.SSLContext.get_ca_certs()`. Affects all Windows installs. | No PR linked |
| **P2** | [#49008 (OPEN)](https://github.com/nousresearch/hermes-agent/issues/49008) | `openai-codex` image gen plugin fails with HTTP 400 because the Codex backend rejects `tool_choice` for `image_generation`. | No PR linked |
| **P2** | [#49002 (OPEN)](https://github.com/nousresearch/hermes-agent/issues/49002) | Fresh install auto-detects `gh auth` token and populates Copilot models even without proper scope – pollutes the model picker with dead entries. | No PR linked |
| **P2** | [#49019 (OPEN PR)](https://github.com/nousresearch/hermes-agent/pull/49019) | *fix(terminal): stop flagging benign nonzero exit codes as failures* – fixes classifiers that treat `grep` (exit 1) or user-interrupt (130) as failures. | PR actively being reviewed |
| **P2** | [#49013 (OPEN PR)](https://github.com/nousresearch/hermes-agent/pull/49013) | *fix(compression): prevent duplicate side-effect actions after compaction* – context compaction can cause re-execution of completed actions (sent messages, created PRs). | PR active |
| **P2** | [#49015 (OPEN PR)](https://github.com/nousresearch/hermes-agent/pull/49015) | *fix(auth): preserve docs URL during OAuth prompt sanitization* – greedy name replacement clobbered the hermes-agent docs URL. | PR active |
| **P2** | [#49012 (OPEN PR)](https://github.com/nousresearch/hermes-agent/pull/49012) | *fix(agent): inherit main runtime connection fields for auto vision* – vision auto path was not inheriting custom `base_url`/`api_key`. | PR active |
| **P2** | [#49001 (OPEN PR)](https://github.com/nousresearch/hermes-agent/pull/49001) | *fix(feishu): improve reply-to message content visibility* – quoted images and rich content were invisible in Feishu/Lark. | PR active |
| **P3** | [#49018 (OPEN PR)](https://github.com/nousresearch/hermes-agent/pull/49018) | *fix(desktop): close idle runtime sessions at chat boundaries* – Desktop could leave orphaned backend sessions. | PR active |
| **P3** | [#49020 (OPEN PR)](https://github.com/nousresearch/hermes-agent/pull/49020) | *fix(desktop): clear subagent and todo state when a session is deleted* – per-session stores not pruned on deletion. | PR active |
| **P3** | [#49016 (OPEN PR)](https://github.com/nousresearch/hermes-agent/pull/49016) | *fix(gateway): prevent clipping of messages containing '<'* – think-tag filter held back single `<` characters. | PR active |
| **P3** | [#49017 (OPEN PR)](https://github.com/nousresearch/hermes-agent/pull/49017) | *fix(telegram): bump MAX_COMMANDS_PER_SCOPE from 30 to 75* – skill-registered commands were being dropped from the BotCommand menu. | PR active |

## Feature Requests & Roadmap Signals
New or updated feature requests today point to several likely directions for the next release:

- **Desktop API key field for custom endpoints** (#42042) – The most requested UI improvement. Likely to ship soon given the existing groundwork.
- **Reorder fallback providers interactively** (#49007) – A small but high-usability enhancement for CLI power users.
- **Configurable pre-session tool auto-invocation** (#49005) – Enables persistent memory tools (e.g., MemPalace) to be run automatically every session. Could land as a configuration option in `agent.pre_session_tools`.
- **Unify provider list across Desktop and CLI** (PR #48800) – A significant refactoring to consolidate three hand-maintained provider lists into a single source.
- **LSP idle timeout** (PR #47332) – Automatically shut down inactive language server processes to save memory on small VPS instances (default 600s).
- **French locale for Desktop** (PR #48070) – Complete translation, ~2100 strings, indicates growing international community.
- **SQLite persistence for shadow clone tasks** (PR #48907) – Gateway crash recovery for background delegation tasks.
- **Honcho filterable memory toolset** (PR #49011) – Exposes a named toolset for restricted sessions (cron, subagents).
- **.sql document and MEDIA delivery support** (PR #49021) – Support SQL files in gateway, a small but practical addition for agent workflows involving databases.

These features, especially #42042, #49007, and #49005, are strong candidates for inclusion in the next minor release.

## User Feedback Summary
**Pain points** reported today include:
- Windows users are hit hardest: the notification “session not found” bug (#48808), `ssl_guard` crash (#49014), and stale `base_url` on model switch (#46759) directly affect reliability.
- Copilot users on fresh installs experience a confusing model picker populated with invalid entries (#49002).
- Codex plugin users cannot use image generation due to `tool_choice` rejection (#49008).
- Terminal power users are frustrated by spurious failure messages from commands like `grep` (#49019).
- Feishu (Lark) users cannot see quoted images or rich content in replies (#49001).

**Satisfaction indicators**: Community is actively submitting PRs and well-scoped issues, suggesting the project is welcoming and responsive. The detailed issue descriptions (e.g., #49005, #49007) show users understand the architecture and provide clear use cases.

## Backlog Watch
- **PR #33032** – *fix(acp): recover silently lost tool completion events with retry mechanism*  
  ([GitHub](https://github.com/nousresearch/hermes-agent/pull/33032), open since May 27, P2)  
  This PR addresses a fire-and-forget pattern that silently drops tool completion events when connected to ACP clients (AionUI, VS Code, etc.). It has been open for over three weeks with no recent review activity. Maintainer attention is needed to prevent potential tool UI inconsistency.

- **Issue #46759** – *Desktop model switch leaves stale base_url* (P2, open since June 15)  
  No linked fix PR. This bug affects all users who switch providers via the Desktop app and is a candidate for immediate triage.

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

Based on the provided GitHub data for PicoClaw (github.com/sipeed/picoclaw) from 2026-06-19, here is a structured project digest.

**Note on Data Context:** All data points reflect activity within the last 24 hours. The project shows a high degree of automated maintenance activity (Dependabot PRs and a nightly release), but low direct human contribution or community discussion volume in this specific window.

---

### 1. Today's Overview

PicoClaw is in a phase of automated maintenance and incremental dependency updates, with a nightly build (v0.3.0-nightly) released today. The issue tracker shows a single, important open bug concerning duplicate message delivery for sub-agents. Pull request (PR) activity is high (12 items), but dominated by Dependabot's automated dependency bumps, with half merged quickly and the other half still open. Overall, the project appears stable but is dealing with a notable user-facing concurrency bug (Issue #3094) that has not yet had a related fix PR surfaced in this window.

### 2. Releases

- **New Release:** `nightly` (v0.3.0-nightly.20260619.287853ab)
  - **Description:** Automated nightly build that may be unstable.
  - **Changelog:** [Full Changelog](https://github.com/sipeed/picoclaw/compare/v0.3.0...main) (compares to the latest stable tag)
  - **Analysis:** This is a standard automated build from the `main` branch. There are no documented breaking changes or specific migration notes for this release. It serves as a preview of the next formal release (v0.3.0).

### 3. Project Progress

This section focuses on PRs that were merged/closed today, as these represent completed work that advances the codebase.

- **Dependency Updates (Merged):** Six Dependabot PRs were closed/merged. These indicate ongoing maintenance and compatibility updates.
    - `actions/checkout` from v6 to v7 ([PR #3144](https://github.com/sipeed/picoclaw/pull/3144))
    - `golang.org/x/term` from v0.43.0 to v0.44.0 ([PR #3146](https://github.com/sipeed/picoclaw/pull/3146))
    - `github.com/Azure/azure-sdk-for-go/sdk/azidentity` from v1.13.1 to v1.14.0 ([PR #3147](https://github.com/sipeed/picoclaw/pull/3147))
    - `github.com/anthropics/anthropic-sdk-go` from v1.46.0 to v1.50.2 ([PR #3149](https://github.com/sipeed/picoclaw/pull/3149))
    - `golang.org/x/sys` from v0.45.0 to v0.46.0 ([PR #3148](https://github.com/sipeed/picoclaw/pull/3148))
    - `github.com/github/copilot-sdk/go` from v0.2.0 to v1.0.1 ([PR #3107](https://github.com/sipeed/picoclaw/pull/3107))

- **Note:** No bug-fix PRs or feature PRs from human contributors were merged in this window.

### 4. Community Hot Topics

The GitHub activity is very low in terms of community discussion. The only active issue dominates the conversation.

- **Most Active Issue (1 item, sole discussion):**
    - `#3094 [Bug] 异步子代理(spawn)任务完成时，ForUser字段被同时用于直接推送和主代理汇总，导致重复消息` ([Issue #3094](https://github.com/sipeed/picoclaw/issues/3094))
    - **Analysis:** This is the only issue updated in the last 24 hours, indicating it is the primary topic of concern for users currently. The underlying need is for cleaner, non-redundant messaging when using sub-agents. Users expect a single, well-formatted, and summarized output, not a raw one followed by a processed one.

### 5. Bugs & Stability

- **Active Bug (High Priority):**
    - **Title:** `#3094 [Bug] 异步子代理(spawn)任务完成时，ForUser字段被同时用于直接推送和主代理汇总，导致重复消息` ([Issue #3094](https://github.com/sipeed/picoclaw/issues/3094))
    - **Severity:** High. The bug causes direct user-facing message duplication across channels (Feishu/Telegram), degrading user experience and creating confusion.
    - **Status:** Open, tagged with `bug`, `subagent`, `duplicate-messages`. It has been open for 9 days with no linked fix PR visible in this data.

- **Regression or Fix PRs:** None were identified in the 24-hour window for this bug.

### 6. Feature Requests & Roadmap Signals

No explicit feature requests were created as issues or PRs in this 24-hour window. However, a clear roadmap signal can be inferred from the dependency updates:

- **GitHub Copilot Integration Focus:** Two separate PRs involved updating the `github.com/github/copilot-sdk/go` dependency (one merged to v1.0.1, one open to v1.0.2). The major version jump (from v0.2.0 to v1.0.x) signals that the Copilot SDK has stabilized and PicoClaw is actively working on integrating or improving its Copilot capabilities. This is likely a core feature for the upcoming v0.3.0 stable release.

- **Prediction for Next Version (v0.3.0):** It will likely include a stable integration with the GitHub Copilot SDK (v1.0+). The nightly release suggests the team is iterating toward this.

### 7. User Feedback Summary

The primary user pain point evident in the data is the **messaging behavior of the async sub-agent (`spawn`) feature**.

- **Pain Point:** Users are experiencing "double messages" – a raw, unformatted message followed by the agent's final summary. This is confusing and noisy.
- **Use Case Affected:** Users who rely on `spawn` for parallel or background tasks (e.g., research, data processing) and expect a clean final report.
- **Satisfaction:** Not directly measurable, but the filing of a well-detailed bug report (Issue #3094) suggests dissatisfaction with the current behavior and a desire for a more polished user experience.

### 8. Backlog Watch

No issues are explicitly marked as `stale` from a lack of activity within this very narrow 24-hour view, but the following are notable for attention:

- **Key Unresolved Bug:** `#3094 [Bug] 异步子代理(spawn)任务完成时...` ([Issue #3094](https://github.com/sipeed/picoclaw/issues/3094)). This is the most critical open item and, having been open for 9 days without a fix PR, requires maintainer prioritization and a response.

- **Open Dependabot PRs (backlog):** Six Dependabot PRs remain open, waiting for maintainer review and merge. These are low risk but could become problematic if left open too long.
    - `eslint` v10.4.1 ([PR #3105](https://github.com/sipeed/picoclaw/pull/3105))
    - `shadcn` v4.11.0 ([PR #3104](https://github.com/sipeed/picoclaw/pull/3104))
    - `typescript-eslint` v8.61.0 ([PR #3103](https://github.com/sipeed/picoclaw/pull/3103))
    - `vite` v8.0.16 ([PR #3101](https://github.com/sipeed/picoclaw/pull/3101))
    - `@vitejs/plugin-react` v6.0.2 ([PR #3100](https://github.com/sipeed/picoclaw/pull/3100))
    - `github.com/github/copilot-sdk/go` v1.0.2 ([PR #3145](https://github.com/sipeed/picoclaw/pull/3145))

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw Project Digest — 2026-06-19

## 1. Today's Overview
The project continues to see **high activity**: 4 issues were updated in the last 24 hours (2 closed, 2 remaining open), and **19 pull requests** were updated (16 still open, 3 merged/closed). No new releases were published today. The community is contributing actively, especially around security hardening, CLI reliability, and communication channel improvements. Several critical bugs (CLI `create` commands completely broken, socket timeouts, session source staleness checks) have dedicated fix PRs under review, indicating strong momentum toward a stable v2.2 release.

## 2. Releases
**None** — no new releases were published as of 2026-06-19.

## 3. Project Progress (Merged/Closed PRs Today)
Three pull requests were merged or closed in the last 24 hours:

- **#2793** (closed) — *feat(agent-to-agent): per-message approval policies on connected agents*  
  Adds an optional, directed approval gate for agent-to-agent messages. When a policy exists for A→B, each message is held until B's human approves. Fully backward-compatible (no policy = free flow). This is a significant feature for controlled inter-agent communication.

- **#2811** (closed) — *fix(setup): allow env-selected agent provider*  
  Enables setting the agent provider via environment variable during setup, improving deployment flexibility.

- **#2810** (closed) — *refactor: mirror .claude skills + CLAUDE.md into .agents via symlinks*  
  Makes `.claude` the single source of truth by creating symlinks from `.agents/skills` → `.claude/skills` and `AGENTS.md` → `CLAUDE.md`. This improves compatibility with agent-convention-aware tools (e.g., Codex) without duplication.

## 4. Community Hot Topics

### Most Active Issues
- **#957** [CLOSED] *Suggest supporting Podman as an alternative to Docker*  
  **10 comments, 7 👍**  
  User request to document Podman as a Docker alternative on macOS/Linux. Community engagement is high, reflecting strong interest in container runtime flexibility.  
  [Issue #957](https://github.com/nanocoai/nanoclaw/issues/957)

- **#29** [CLOSED] *feat: Add Signal as messaging channel*  
  **7 comments, 4 👍**  
  Feature request to add Signal communication integration, following patterns of existing `/add-telegram`, `/add-slack`, etc. Indicates demand for encrypted messaging support.  
  [Issue #29](https://github.com/nanocoai/nanoclaw/issues/29)

- **#2632** [OPEN] *Clarify status of Telegram agent-swarm / multi-bot identity in v2*  
  **2 comments**  
  A user planning a v1→v2 migration is confused about the `add-telegram-swarm` feature’s status. The issue remains open without maintainer reply, causing uncertainty.  
  [Issue #2632](https://github.com/nanocoai/nanoclaw/issues/2632)

### Most Active Pull Requests (still open)
- **#2812** — *fix(discord): chunk replies over 2000 chars instead of truncating*  
  Addresses a long-standing pain point where long Discord replies were silently cut off. The PR introduces proper chunking via the Chat SDK bridge.  
  [PR #2812](https://github.com/nanocoai/nanoclaw/pull/2812)

- **#2809** — *feat(apple-container): Apple Container runtime + remote OneCLI gateway*  
  Adds experimental Apple Container support and remote OneCLI gateway, gated by `CONTAINER_RUNTIME` env var. Highlights community interest in macOS-native containerisation.  
  [PR #2809](https://github.com/nanocoai/nanoclaw/pull/2809)

**Analysis of underlying needs:** The community is pushing for **container runtime diversity** (Podman, Apple Container), **additional messaging channels** (Signal, improved Discord chunking), and **clearer upgrade guidance** (Telegram swarm v2 status). There is also a strong emphasis on **security hardening** (multiple PRs for confined file reads, input validation).

## 5. Bugs & Stability

Several bugs were reported today, with fix PRs already in review. Severity ranking:

| Severity | Issue/PR | Description | Status |
|----------|----------|-------------|--------|
| **Critical** | #2804 (fix open) | `ncl messaging-groups create` always throws `NOT NULL constraint failed` — CLI create path completely dead. | Fix PR #2804 exists |
| **Critical** | #2802 / #2813 (fixes open) | `SocketTransport.sendFrame` has no request timeout and no response buffer bound; unresponsive hosts hang forever. | Fix PRs #2802 and #2813 exist |
| **High** | #2784 (issue open, no fix PR) | `container-runner.ts` staleness check only watches `index.ts`, missing changes in `ipc-mcp-stdio.ts` — may cause stale code sync in multi-file container agent runners. | No fix PR yet |
| **High** | #2818 / #2817 (fixes open) | `send_file` reads not confined to agent workspace — allows reading sibling mounts like `/workspace/extra`. | Fix PRs #2817 and #2818 exist |
| **Medium** | #2814 (fix open) | CLI `groups create` lacks validation of group folder path; traversal attacks possible. | Fix PR #2814 exists |
| **Medium** | #2801 / #2815 (fixes open) | `safeParseContent` returns `undefined` for primitive JSON (e.g., `"5"`, `true`) because callers expect `.text` / `.sender` on non-objects. | Fix PRs #2801 and #2815 exist |
| **Low** | #2792 (fix open) | `add-imessage` skill fails on fresh checkouts because `src/channels/` directory doesn't exist. | Fix PR #2792 exists |

**Other open bug-fix PRs:** #2808 (idempotent `insertMessage` + missing `agent_group_id` on approval rows), #2816 (Discord chunking regression test).

## 6. Feature Requests & Roadmap Signals

### User-Requested Features (active/open)
- **Podman as alternative to Docker** (#957) — closed, likely merged into docs or config.
- **Signal messaging channel** (#29) — closed, feature implemented.
- **Apple Container runtime** (#2809) — new experimental feature, may be folded into next release.
- **Read-only CLI-derived dashboard skill** (#2795) — utility skill proposed.

### Predicted Next Version Highlights
Based on merged PRs and open activity, v2.2 (or v2.1.19) is likely to include:
1. **Agent-to-agent approval policies** (already merged in #2793).
2. **Apple Container / remote OneCLI gateway** (#2809) if reviewed quickly.
3. **Security fixes** for file read confinement, CLI validation, and socket handling.
4. **Discord chunking** fix (#2812 / #2816).
5. **Telegram swarm status clarification** — maintainers may need to address #2632 before migration guides are updated.

## 7. User Feedback Summary

### Pain Points
- **Docker dependency** (Issue #957): Users want Podman as a lighter alternative, especially on macOS/Linux.
- **CLI completely broken for group creation** (Issue related to #2804): Blocks users from setting up messaging groups.
- **Telegram swarm v2 status unclear** (Issue #2632): User planning migration is stuck.
- **Discord replies silently truncated** (Issue behind #2812): Long messages lost without warning.
- **Session source staleness check misses files** (Issue #2784): Potential silent data loss in container-runners.

### Satisfaction / Use Cases
- Positive remarks about project being “very useful and well designed” (#957).
- Active contribution of fix PRs indicates a healthy community willing to invest in stability.
- Multiple security-focused PRs suggest users trust NanoClaw for production but demand robust boundaries.

## 8. Backlog Watch

Issues and PRs that appear to lack maintainer attention or remain unanswered for extended periods:

- **Issue #2632** (created 2026-05-28, last updated 2026-06-18, 2 comments) — *Clarify Telegram swarm status*  
  No maintainer reply. User explicitly asks for guidance before committing to a migration path. Needs prompt response.

- **Issue #2784** (created 2026-06-16, 1 comment) — *Session source staleness check misses ipc-mcp-stdio.ts*  
  Bug with clear reproduction. No fix PR yet; maintainers should acknowledge and assign.

- **PR #2792** (created 2026-06-17, still open) — *fix(add-imessage): mkdir src/channels before writing*  
  Simple one-line fix for a regression. Waiting for merge.

- **PR #2795** (created 2026-06-17) — *feat: add /add-clidash — read-only CLI-derived dashboard skill*  
  Skill addition following contribution guidelines. No reviewer comments yet.

- **PR #2809** (created 2026-06-18) — *Apple Container runtime*  
  Significant new feature; needs architectural review and testing plan.

**Recommendation:** Maintainers should prioritize #2632 (user guidance) and #2784 (bug fix) to avoid community frustration. The critical CLI bug fix (#2804) should be fast-tracked to restore functionality.

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw Project Digest – 2026-06-19

**Date:** 2026-06-19  
**Prepared by:** AI Agent Analysis (data from github.com/nullclaw/nullclaw)

---

## 1. Today's Overview

NullClaw saw moderate activity over the past 24 hours, with **4 open issues** and **4 open pull requests** updated — none of which were closed or merged. No new releases were published. The project remains in an active development phase, with significant contributions focused on **Android/Termux (aarch64) compatibility**, **streaming tool-call improvements**, and **documentation for WeChat QR code login**. Community engagement is steady, though several issues have been open for months without resolution. The maintainer response rate appears to be the primary bottleneck.

---

## 2. Releases

**None.** No new releases were published today.

---

## 3. Project Progress

**No PRs were merged or closed today.** However, four open PRs demonstrate ongoing development:

- [#966 – fix(http): route stdlib HTTP through curl on aarch64-linux-android](https://github.com/nullclaw/nullclaw/pull/966)  
  Addresses a runtime HTTP failure (`error.NameServerFailure`) on Termux by using an alternative transport.

- [#965 – Proposal: structured streaming tool-call support for SSE parser](https://github.com/nullclaw/nullclaw/pull/965)  
  Companion to the root fix for streaming tool calls; improves the SSE parser to handle model‑emitted XML tool calls.

- [#964 – Enable native API-level tool calls during streaming](https://github.com/nullclaw/nullclaw/pull/964)  
  Fixes a bug where native tools were not sent during streaming requests, preventing tool-use in real-time conversations.

- [#963 – docs(channels): document weixin personal WeChat QR code login channel](https://github.com/nullclaw/nullclaw/pull/963)  
  Adds English and Chinese documentation for the recently‑added WeChat personal account QR login feature (closes #817).

These PRs indicate that the project is actively improving **cross‑platform reliability** (Android), **real‑time tool calling**, and **documentation**.

---

## 4. Community Hot Topics

The most active discussions revolve around **platform compatibility** and **Chinese‑market integrations**:

- **[#50 – Can this run on an Esp32?](https://github.com/nullclaw/nullclaw/issues/50) (4 comments)**  
  User requests ESP32 support. The issue has been open since February 2026, indicating sustained interest in embedded/IoT deployments.

- **[#484 – Feishu (飞书) cannot connect to internet](https://github.com/nullclaw/nullclaw/issues/484) (3 comments)**  
  A bug report with a screenshot showing a network failure. Likely an authentication or proxy issue specific to the Feishu channel. No fix PR yet.

- **[#817 – Does nullclaw support WeChat QR code login?](https://github.com/nullclaw/nullclaw/issues/817) (2 comments)**  
  Feature request now addressed by PR #963. Community eagerly awaiting documentation to enable the feature.

- **[#868 – zig build fails on Android/Termux (aarch64) with AccessDenied on options.zig linkat](https://github.com/nullclaw/nullclaw/issues/868) (2 comments)**  
  Blocks Android builds completely. The issue has been open since late April, though PR #966 addresses a related runtime issue.

**Underlying needs:** Users are looking for broader platform support (ESP32, Android), seamless integration with Chinese messaging apps (WeChat, Feishu), and reliable builds on non‑standard environments.

---

## 5. Bugs & Stability

| Issue | Severity | Description | Fix PR? |
|-------|----------|-------------|---------|
| [#868](https://github.com/nullclaw/nullclaw/issues/868) | **High** | `zig build` fails on Android/Termux (aarch64) with `AccessDenied` on `linkat`. Prevents compilation. | No direct PR, but [#966](https://github.com/nullclaw/nullclaw/pull/966) fixes related HTTP runtime issue on same platform. |
| [#484](https://github.com/nullclaw/nullclaw/issues/484) | **Medium** | Feishu (飞书) channel cannot connect to internet. Blocks functionality for Chinese users. | No. |
| [#966](https://github.com/nullclaw/nullclaw/pull/966) | **Medium** (runtime) | `std.http.Client` fails with DNS errors on Android. PR is open but not yet merged. | Open (this PR is the fix itself). |

**Summary:** The most severe bug is the Android build failure (#868), which has been open for nearly two months. The runtime HTTP issue on the same platform is being addressed but not yet merged. No regressions were reported in the last 24 hours.

---

## 6. Feature Requests & Roadmap Signals

- **ESP32 support** ([#50](https://github.com/nullclaw/nullclaw/issues/50)) – Unlikely to be in next release given the embedded constraints, but interest persists.
- **WeChat QR code login** ([#817](https://github.com/nullclaw/nullclaw/issues/817)) – Already implemented per PR #963; documentation only is pending merge. Likely in next release.
- **Streaming tool‑call support** ([#965](https://github.com/nullclaw/nullclaw/pull/965), [#964](https://github.com/nullclaw/nullclaw/pull/964)) – Two PRs target enabling tool use during streaming. If merged, next release will include this feature.
- **Android HTTP stability** ([#966](https://github.com/nullclaw/nullclaw/pull/966)) – Essential for Android users; high‑priority fix.

**Prediction:** The next release (assuming one is near) will likely contain WeChat QR login documentation, the HTTP fix for Android, and streaming tool‑call improvements.

---

## 7. User Feedback Summary

**Real pain points:**
- Android/Termux users cannot build or run the project reliably (issues #868, #966).
- Chinese users need Feishu connectivity fixes and WeChat login instructions.
- Users exploring IoT/embedded (ESP32) face unclear compatibility status.

**Use cases expressed:**
- “I would love to get your thoughts on this” (ESP32) – enthusiasm for expansion.
- “Does nullclaw support WeChat QR code login?” – practical need for personal account access.
- “Zig build fails” – directly blocks development on mobile devices.

**Satisfaction/dissatisfaction:** No explicit complaints, but the lack of merged fixes for months‑old issues (e.g., #868) may frustrate Android contributors. The new PRs and documentation suggest the maintainers are actively responding.

---

## 8. Backlog Watch

The following open items have not received maintainer attention recently and may block community contributions:

| Issue/PR | Date Created | Last Updated | Needs Maintainer Action |
|----------|--------------|--------------|-------------------------|
| [#50 – ESP32 support](https://github.com/nullclaw/nullclaw/issues/50) | 2026-02-21 | 2026-06-18 | Clarify feasibility or close as “won’t implement”. |
| [#484 – Feishu not connecting](https://github.com/nullclaw/nullclaw/issues/484) | 2026-03-13 | 2026-06-19 | Investigate network failure; reproduce or request logs. |
| [#868 – Android build failure](https://github.com/nullclaw/nullclaw/issues/868) | 2026-04-23 | 2026-06-19 | High‑priority; consider merging related PR #966 first, then address build issue. |
| [#817 – WeChat QR login](https://github.com/nullclaw/nullclaw/issues/817) | 2026-04-14 | 2026-06-18 | Nearly resolved by PR #963 – merge and close. |

**Recommendation:** The maintainers should prioritize reviewing and merging PRs #963, #964, #965, and #966 to unblock users and reduce open issue count. A comment on #50 and #484 would also improve community trust.

---

*End of digest for 2026-06-19.*

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw Project Digest — 2026-06-19

## Today's Overview
IronClaw shows elevated activity with **7 issues updated** (3 closed, 4 open) and **21 pull requests** touched in the last 24 hours. The team continues to focus on stabilising the Reborn runtime, especially around OAuth flows, approval UX, and automation workflows. Two PRs were closed/merged today, and several high-severity bugs were addressed promptly. The project appears healthy and responsive, though a few longstanding infrastructure issues remain open.

## Releases
No new releases in the last 24 hours.

## Project Progress
**Closed/merged PRs today (2):**
- **#4968** – Fix GSuite refresh auth classification: credentials rejected after refresh are now correctly flagged as `auth-required` instead of `backend auth`.
- **#5082** – Bound approval command previews: truncates long shell commands in approval modals with a toggle (fixes #5078).

**Other notable merged/closed PRs from the last 48 hours:**
- **#5065** – Fire-once (one-shot) scheduled triggers (merged 2026-06-18).
- **#5055** – Soften automation run errors in WebUI (merged 2026-06-18).
- **#5018** – Projects WebChat v2 HTTP endpoints (merged 2026-06-18).

**Open PRs advancing key features:**
- #5081 – Hosted single-tenant Postgres profile (DB migration, low risk).
- #4989 – Persist Engine V2 LLM usage for admin billing.
- #5085 – Concurrent turn execution via TurnRunScheduler.
- #5019 – Light up the Projects page (frontend, stack 5/5).
- #5087 – Proactively refresh Google OAuth tokens before expiry.
- #5084 – Redesign automations page (frontend).
- #5072 – Generic host-ingress for Slack.
- #5063 – Per-turn auto-approve resolution with hard floor.

## Community Hot Topics
**Most commented issue:**
- [**#4907**](nearai/ironclaw Issue #4907) – (CLOSED) Run fails after successful Google OAuth instead of resuming execution. 3 comments. User `sunglow666` identified a critical UX break after authentication. The issue is now closed, indicating a fix was deployed.

**Most active issues (by resonance):**
- [**#1012**](nearai/ironclaw Issue #1012) – (OPEN) Alibaba Coding Plan not usable in openai_compatible mode. 1 comment, 1 👍. Seeking parity with other claw frameworks.
- [**#5088**](nearai/ironclaw Issue #5088) – (OPEN) Shell approval prompt sometimes asks to approve `reads` as a command. Newly filed, reflecting confusion in approval UX.

**Underlying needs:** Users are demanding smoother authentication hand-off (OAuth), clearer approval dialogs, and better compatibility with third-party LLM providers. The rapid closure of #4907 and #5078 shows the team is responsive to these pain points.

## Bugs & Stability
| Severity | Issue | Status | Description | Fix PR |
|----------|-------|--------|-------------|--------|
| **High** | [#4907](nearai/ironclaw Issue #4907) | CLOSED | Google OAuth completes but original run fails | Fixed (no explicit PR link) |
| **High** | [#5060](nearai/ironclaw Issue #5060) | CLOSED | Repeated approval loops on GitHub analysis workflows | Fixed (likely #5055 + #5082) |
| **High** | [#4108](nearai/ironclaw Issue #4108) | OPEN | Nightly E2E scheduled run failing (since May 27) | No fix PR yet |
| **Medium** | [#5078](nearai/ironclaw Issue #5078) | CLOSED | Approval modal hard to review with large commands | Fixed in #5082 |
| **Medium** | [#5083](nearai/ironclaw Issue #5083) | OPEN | Scoped index for automations list state-unaware | No fix PR |
| **Low** | [#5088](nearai/ironclaw Issue #5088) | OPEN | Shell approval prompt shows misleading `reads` | No fix PR |

**Other bug-fix PRs active:**
- #5045, #5043 – Fail fast on HTTP 400 invalid model (stop retry hang).
- #4990 – Fix NEAR AI MCP ready state projection.
- #5087 – Proactive Google OAuth token refresh (closes #5071).

## Feature Requests & Roadmap Signals
- **#1012** – Alibaba Coding Plan support in openai_compatible mode. Low complexity, likely to be addressed in a future release.
- **#5088** – Clarify shell approval language. Suggests a UX polish sprint.
- **#5084** – Redesign automations page (PR open). Indicates a UX refresh is incoming.
- **#5085** – Concurrent turn execution. Performance improvement likely targeted for next minor version.
- **#5081** – Hosted single-tenant Postgres profile. Signals multi-tenant infrastructure prep.
- **#5072** – Generic Slack ingress. Shows platform expansion beyond chat.

## User Feedback Summary
Real user pain points surfaced this week:
- **sunglow666** (test engineer) reported three distinct bugs (OAuth resume, approval modal overflow, approval loops), all of which were closed quickly. This reflects high satisfaction with the team’s responsiveness.
- **think-in-universe** filed #5088 about confusing `reads` approval prompt – indicates that even experienced users find the approval language ambiguous.
- **wznmickey** (#1012) wants broader LLM provider compatibility – dissatisfaction with locked-in OpenAI mode.
- **abbyshekit** submitted two small PRs (#5045, #5043) to fix the `auto` model retry hang – a real frustration for desktop app users.

Overall sentiment: The project is actively improving the Reborn UX, but users expect smoother authentication and more transparent approval controls.

## Backlog Watch
- [**#4108**](nearai/ironclaw Issue #4108) – **Nightly E2E failing** (opened 2026-05-27, no comments). Critical infrastructure issue; if the nightly pipeline is broken, regressions may go undetected. Needs maintainer attention.
- [**#1012**](nearai/ironclaw Issue #1012) – **Alibaba Coding Plan** (opened 2026-03-12, 1 comment). Low engagement but a valid feature request. Possibly deprioritised due to scope.
- [**#5083**](nearai/ironclaw Issue #5083) – **Triggers scoped index** (opened 2026-06-18, 0 comments). Unbounded prefix scan; performance issue that could bite at scale. No PR yet.

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI Project Digest — 2026-06-19

## 1. Today's Overview
Project activity remains moderate today, with **4 issues updated** in the last 24 hours (1 newly opened, 3 closed) and **zero PR activity**. A new patch release **2026.6.18** shipped yesterday, expanding artifact sharing support to more file types (Word, PPT, Excel, PDF, Markdown, Mermaid) and cleaning up the voice input subsystem. The three closed issues were all stale bugs from early April — none linked to today's PRs — indicating a possible maintenance sweep. One high-impact feature request (#2180) opened today, proposing a major platform upgrade. Overall, the project appears stable but with low short-term contribution momentum (zero PRs).

## 2. Releases
**LobsterAI 2026.6.18** (released 2026-06-18)
- **New features:** Artifact sharing capabilities upgraded to support Word, PPT, Excel, PDF, Markdown, and Mermaid file types.
- **Bug fixes / improvements:** `keep only realtime asr` — voice input subsystem simplified by removing non-realtime ASR modes.
- **Breaking changes:** None reported.
- **Migration notes:** Users relying on non-realtime ASR for voice input should update their workflows to realtime-only paths.

👉 [Release link](https://github.com/netease-youdao/LobsterAI/releases/tag/2026.6.18)

## 3. Project Progress
**Today's PRs merged/closed:** 0
No pull requests were merged or closed today. The three closed issues (#1487, #1471, #1472) were automatically marked stale and then closed — none have attached fix PRs. No feature or bugfix advances visible from PR activity.

## 4. Community Hot Topics
Most active issues today (by comments / reactions):
- **#2180 [OPEN]** — *Build "AI Collaborator" Form: Introduce Natural Language Command Bar and Task Dispatch Console for Cross-Model Orchestration and Project-Level Memory*  
  Author: woxinsj | Comments: 0 | 👍: 0  
  Summary: Proposes upgrading OpenClaw from a low-level toolset to an "AI Collaborator" platform for "tech-savvy non-elite programmers," enabling continuous cross-model orchestration and persistent project-level memory.  
  [Issue link](https://github.com/netease-youdao/LobsterAI/issues/2180)

No other issues or PRs attracted notable discussion today. The stale issues from April had 2–3 comments each but are now closed.

**Underlying need:** Users are increasingly looking for higher-level abstractions beyond tool orchestration — a unified command interface that coordinates multiple AI models and retains long-term project context.

## 5. Bugs & Stability
No new bugs were reported today. The three closed stale issues (all from April) documented UX regressions:
- **#1471 (Closed)** — Draft content loss when switching sessions/views due to debounce not flushing before unmount (moderate severity, UX impact).
- **#1472 (Closed)** — "Re-edit" button overwrites unsent input without confirmation (moderate severity, data loss risk).
- **#1487 (Closed)** — Python script execution fails locally with 30B model (medium severity, model-dependent).

None of these have associated fix PRs. The voice-input fix in the latest release (`keep only realtime asr`) resolved a previously unreported ASR subsystem issue.

**Overall stability:** No critical regressions or crashes reported today. The project is in a stable state with low bug pressure.

## 6. Feature Requests & Roadmap Signals
The only feature request opened today is **#2180** — a comprehensive proposal for an "AI Collaborator" layer. It signals a demand for:
- Natural language command interfacing
- Cross-model orchestration
- Persistent project-level memory

This aligns with an emerging theme: developers want the agent to act as a **collaborative partner**, not just a tool. If the maintainers respond positively, this could be the basis for a major v2026.7 or later release.

No other feature requests were created today. The artifact sharing enhancement in the latest release was a user-facing improvement that likely addresses prior requests.

## 7. User Feedback Summary
Based on recent closed issues and one new feature request:
- **Pain points:**
  - Draft content loss when navigating between views (#1471)
  - Unintended overwrite of unsent input by re-edit (#1472)
  - Local model-dependent Python script failures (#1487)
- **Use cases:**
  - Artifact sharing across multiple file types (Word, PPT, Excel, PDF, etc.) — received positively in the latest release
  - Voice input for faster interaction — simplified to realtime-only, which may frustrate users who relied on offline ASR
- **Satisfaction / dissatisfaction:**
  - No direct positive/negative sentiment expressed in issues. The closed stale bugs suggest some users experienced friction, but no outcry.
  - The #2180 proposal implies a desire for more integrated, higher-level agent capabilities.

## 8. Backlog Watch
No long-unanswered issues or PRs requiring maintainer attention today. The three stale issues were all closed — no open backlog items predating this week. The newly opened #2180 has zero comments and may become a candidate for maintainer response in the coming days.

**Potential risk:** Zero PR activity over the last 24 hours, combined with only one new issue, may indicate low contributor engagement. Maintainers should monitor for contributor burnout or declining community momentum.

---

*Generated from GitHub data on 2026-06-19. All links refer to [github.com/netease-youdao/LobsterAI](https://github.com/netease-youdao/LobsterAI).*

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis Project Digest – 2026-06-19

## 1. Today's Overview
Project activity remains very low, with only one issue updated in the last 24 hours and no pull requests or releases. The single open bug report suggests users are encountering a functional limitation regarding session management. No new code merges, feature work, or maintainer responses have occurred, indicating a quiet period for the project. The lack of any merged or closed PRs suggests development may be paused or focused on longer-term tasks.

## 2. Releases
*No new releases were published today.*  
*(No release data to report.)*

## 3. Project Progress
- **Merged/Closed PRs:** 0  
- **Features advanced or fixed:** None. No pull requests were updated, merged, or closed in the last 24 hours.

## 4. Community Hot Topics
- **#1132** (Open, Bug) – ["main" session can't be deleted/archived](https://github.com/moltis-org/moltis/issues/1132)  
  *Author: vvuk | Created: 2026-06-18 | Comments: 0 | 👍: 0*  
  This is the only active discussion. Although it has zero comments or reactions, the issue highlights a potential usability gap: users cannot delete or archive the default “main” session. Underlying need: users want full control over session lifecycle, including the ability to remove default sessions to avoid clutter or to reset state. The lack of engagement may indicate either a recent report or low community awareness.

## 5. Bugs & Stability
| Severity | Issue | Description | Fix PR Exists? |
|----------|-------|-------------|----------------|
| Medium   | [#1132](https://github.com/moltis-org/moltis/issues/1132) | "main" session cannot be deleted or archived. User reports having searched existing issues and using latest version. | No |

**Severity rationale:** The bug blocks a core workflow (session management) but does not cause crashes or data loss. Without a fix PR, the issue remains unaddressed.

## 6. Feature Requests & Roadmap Signals
- No explicit feature requests were filed today.  
- The bug in #1132 implicitly suggests a **session management enhancement**: allowing users to delete or archive all sessions, including the default “main” one. This could be a candidate for the next minor release if maintainers prioritize it.

## 7. User Feedback Summary
- **Pain point:** A user (vvuk) is blocked from cleaning up or resetting the default session, indicating dissatisfaction with session management flexibility.  
- **Use case:** Users likely want to start fresh or reorganize their chat history, but the restriction prevents that.  
- **Satisfaction:** Low perceived ease-of-use for this specific workflow. No other feedback signals available.

## 8. Backlog Watch
- No long-unanswered issues or PRs were identified in the provided data. All tracked items are recent and open.  
- Maintainer attention may be needed on issue #1132, as it directly affects user workflow and has no response after one day.

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw Project Digest – 2026-06-19

## Today's Overview

The project saw moderate activity with 10 issues and 11 PRs updated in the last 24 hours. The community remains highly engaged, reporting several regressions from the recent v1.1.12 release and requesting improvements to mobile and multi-agent workflows. A key stability fix for ChromaDB memory bloat (Issue #4795) was merged, addressing a critical crash. New contributors are active, submitting fixes for context compaction, sandboxing, and memory ranking. No new releases were published today.

## Releases

No new releases today. The latest available version remains **v1.1.12.post1** (referenced in several bug reports). Users upgrading from v1.1.11 to v1.1.12 have noted a regression in image display via `send_file_to_user` (Issue #5320).

## Project Progress

**Merged/Closed PR today:**
- **[PR #5332](https://github.com/agentscope-ai/QwenPaw/pull/5332)** – fix(memory): added index maintenance and timeout protection for ChromaDB (closed). This directly addresses the vector index runaway bug (#4795). It introduces `compact_index()`, `purge_index()`, auto-compaction at configurable thresholds, and a 5‑second query timeout.
- **Issue #5319 (closed)** – Console channel "Answers have stopped" resolved after reinstallation.

**Other notable open PRs (not merged yet):**
- [#5334](https://github.com/agentscope-ai/QwenPaw/pull/5334) – feat(ui): allow switching agent in collapsed sidebar mode (addresses #5329).
- [#5331](https://github.com/agentscope-ai/QwenPaw/pull/5331) – feat(console): real-time SSE push notifications with voice beep (addresses #5322).
- [#5326](https://github.com/agentscope-ai/QwenPaw/pull/5326) – feat: minimize window to system tray on close.
- [#5325](https://github.com/agentscope-ai/QwenPaw/pull/5325) – feat(memory): add optional recency-aware ranking for daily notes.
- [#5323](https://github.com/agentscope-ai/QwenPaw/pull/5323) – feat(plan): add todo_write progress panel for multi-step agent tasks.
- [#5324](https://github.com/agentscope-ai/QwenPaw/pull/5324) – fix(files): use inline content-disposition for file preview (addresses #5320).
- [#5321](https://github.com/agentscope-ai/QwenPaw/pull/5321) – feat(context): scroll context manager – durable history + recall REPL.
- [#5310](https://github.com/agentscope-ai/QwenPaw/pull/5310) – feat(sandbox): add bubblewrap Linux sandbox with mount namespace isolation.
- [#5241](https://github.com/agentscope-ai/QwenPaw/pull/5241) – fix(cron): increase default misfire_grace_seconds from 60 to 3600.
- [#5287](https://github.com/agentscope-ai/QwenPaw/pull/5287) – fix(context): don't crash compaction when summary exceeds schema maxLength.

## Community Hot Topics

The most active discussions centre on usability and model compatibility:

1. **[Issue #5329](https://github.com/agentscope-ai/QwenPaw/issues/5329) (3 comments, open)** – User requests a button to switch agents when the sidebar is in collapsed/compact mode (mobile use case). A fix PR (#5334) was opened the same day.

2. **[Issue #4795](https://github.com/agentscope-ai/QwenPaw/issues/4795) (3 comments, open)** – ChromaDB index ballooning to 37 GB, causing repeated `memory_search` crashes. The fix PR (#5332) was merged today, providing immediate relief.

3. **[Issue #5328](https://github.com/agentscope-ai/QwenPaw/issues/5328) (2 comments, open)** – DeepSeek agents getting stuck during “thinking” state across all UIs. Workaround: manual stop then continue. Indicates a broader compatibility issue with DeepSeek streaming or thinking tokens.

4. **[Issue #5322](https://github.com/agentscope-ai/QwenPaw/issues/5322) (1 comment, open)** – Request for real-time UI updates and voice notifications when API messages arrive. PR #5331 addresses this with SSE push.

The underlying needs reflect a desire for better mobile support, stable memory management, and seamless interaction with third‑party model providers.

## Bugs & Stability

Bugs reported today are **moderate to critical**. The merged fix for #4795 addresses the most severe (memory crash). Remaining issues:

| Severity | Issue | Description | Fix exist? |
|----------|-------|-------------|------------|
| **High** | [#5330](https://github.com/agentscope-ai/QwenPaw/issues/5330) | Zhipu API: supplier test passes, but all model connections fail. Likely model routing/name resolution bug. | No |
| **High** | [#5333](https://github.com/agentscope-ai/QwenPaw/issues/5333) | After submitting instruction, agent freezes; text box remains editable (no stop button). Possibly DeepSeek-related. | No |
| **High** | [#5328](https://github.com/agentscope-ai/QwenPaw/issues/5328) | DeepSeek thinking stuck across all UIs. | No |
| **Medium** | [#5320](https://github.com/agentscope-ai/QwenPaw/issues/5320) | v1.1.12 regression: `send_file_to_user` no longer displays images inline (download only). PR #5324 provides fix. | Yes (#5324) |
| **Medium** | [#5317](https://github.com/agentscope-ai/QwenPaw/issues/5317) | Tauri on Windows cannot find Python after previous conda‑bundled version. Skills cannot run Python scripts. | No |
| **Low** | [#5319](https://github.com/agentscope-ai/QwenPaw/issues/5319) | Console channel shows "Answers have stopped" – resolved by reinstall (closed). | N/A |

## Feature Requests & Roadmap Signals

- **Mobile & Sidebar** – [#5329](https://github.com/agentscope-ai/QwenPaw/issues/5329) (agent switch in collapsed sidebar) and related PR #5334. Likely to land in next patch.
- **Agent Office Interactions** – [#5327](https://github.com/agentscope-ai/QwenPaw/issues/5327) requests chat buttons on agent cards, plus session switching. Already has a PR (#5327 itself? no, but community interest is high).
- **Real‑time Notifications** – [#5322](https://github.com/agentscope-ai/QwenPaw/issues/5322) with voice beep is being addressed by PR #5331.
- **Sandbox** – [#5310](https://github.com/agentscope-ai/QwenPaw/pull/5310) adds `bubblewrap` isolation for Linux. Indicates growing concern for security.
- **Context & Memory** – Multiple PRs enhance context strategies (scroll, compaction) and memory ranking (recency). These reflect advanced user needs for long‑running agents.
- **Cron Reliability** – [#5241](https://github.com/agentscope-ai/QwenPaw/pull/5241) increases misfire grace period, a practical fix for busy agents.

The next release (likely v1.1.13) will probably include the collapsed sidebar fix, image display regression fix, real‑time push, and the ChromaDB maintenance tools.

## User Feedback Summary

Users are **satisfied** with the responsiveness of the maintainers (fixes often appear the same day) but **frustrated** by regressions in v1.1.12 (image display, API model failures). Mobile web usage is on the rise, with clear desire for a touch‑friendly UI. DeepSeek users are experiencing intermittent freezes, which may need coordination with upstream model changes. The Zhipu API inconsistency (supplier pass vs model fail) suggests an internal routing bug. Overall, the project is seen as powerful but sometimes brittle under heavy memory or specific model configurations.

## Backlog Watch

- **[Issue #4795](https://github.com/agentscope-ai/QwenPaw/issues/4795)** – ChromaDB memory bloat (open since May 29). Fix merged today (#5332); users should test the update.
- **[PR #5241](https://github.com/agentscope-ai/QwenPaw/pull/5241)** – cron misfire grace period (first-time contributor, open since June 16). No comments from maintainers; may need review.
- **[#5317](https://github.com/agentscope-ai/QwenPaw/issues/5317)** – Tauri Python path issue (open June 18, 2 comments). No assigned dev or PR yet.
- **[#5330](https://github.com/agentscope-ai/QwenPaw/issues/5330)** – Zhipu model connectivity (reported today, no fix). Requires investigation into model name resolution logic.

These items should be prioritised by maintainers to prevent user churn and ensure platform stability.

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

No activity in the last 24 hours.

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw Project Digest – 2026-06-19

## Today's Overview

ZeroClaw saw extremely high activity on 2026-06-19, with 8 open issues and 50 pull requests updated in the last 24 hours (49 open, 1 merged/closed). A new patch release **v0.8.1** landed, packing 207 commits from 45 contributors and emphasizing bug fixes (123) with substantial new features (46). The project continues to stabilize its multi-agent runtime, channels, and provider stack while simultaneously pushing forward several large‑scope features (Discord parity, WASM plugin host, per‑turn output routing, authentication). The one merged PR today—a refactor bundling `run_tool_call_loop` arguments into a `ToolLoop` struct—signals ongoing architectural cleanup as v0.8.x matures toward v0.9.0.

## Releases

**ZeroClaw v0.8.1** – First patch on the v0.8.x line, focused on stabilizing the multi‑agent runtime, channels, and provider stack introduced in v0.8.0. Highlights include 123 bug fixes, 46 new features, and contributions from 45 developers. No breaking changes or migration notes are documented in the release body; it is expected to be a drop‑in update.  
🔗 [Release v0.8.1](https://github.com/zeroclaw-labs/zeroclaw/releases/tag/v0.8.1)

## Project Progress

The only PR merged/closed today was **#7969**, a refactor that bundles the 34 positional arguments of `run_tool_call_loop` into a `ToolLoop` struct, reducing code complexity and removing a `#[allow(clippy::too_many_arguments)]` suppression. This is a foundational change that simplifies the call sites for the upcoming universal ingress policy layer (#7997, stacked on top).  
🔗 [PR #7969](https://github.com/zeroclaw-labs/zeroclaw/pull/7969)

Beyond the closed PR, 49 other PRs remain open and active, many approaching merge readiness (see feature highlights below). The broad range of updates indicates that the project is simultaneously landing fixes and building toward v0.8.2/v0.8.3 and v0.9.0 milestones.

## Community Hot Topics

The most engaged issues and PRs (by comment count) reflect user demand for session persistence, model switching flexibility, and Discord feature completeness.

- **#7759** (3 comments) – Gateway WebSocket lifetime decoupling: treat the WebSocket as a transport channel, not as the owner of the agent turn, so client disconnects don’t cancel in‑flight turns. This is a high‑risk, priority‑p1 feature in progress.  
  🔗 [Issue #7759](https://github.com/zeroclaw-labs/zeroclaw/issues/7759)

- **#6557** (2 comments) – Reconciliation of runtime model switching with the provider structure for v0.8.0. A priority‑p2 feature that touches channel slash commands and agent config.  
  🔗 [Issue #6557](https://github.com/zeroclaw-labs/zeroclaw/issues/6557)

- **#7432** (1 comment) – v0.9.0 tracker for auth, security, gateway, and breaking‑change backlog. Signals the community’s desire for hardened security boundaries.  
  🔗 [Issue #7432](https://github.com/zeroclaw-labs/zeroclaw/issues/7432)

- **#7831** (no comments yet but widely referenced) – Discord channel interaction‑surface parity tracker covering embeds, typed slash options, components, and voice. This is a coordination hub for multiple PRs; the related PR #7833 (Discord embeds) is large (XL) and currently waiting on author action.  
  🔗 [Issue #7831](https://github.com/zeroclaw-labs/zeroclaw/issues/7831)

The underlying needs are clear: users want **reliable long‑running sessions** (WebSocket decoupling), **flexible model selection across channels**, and a **feature‑complete Discord integration** comparable to other mature agent bots.

## Bugs & Stability

Several bugs were fixed or reported today, with most having corresponding fix PRs. Severity is generally medium, with no critical regressions observed.

| Bug / Issue | Severity | Description | Fix PR (if exists) |
|------------|----------|-------------|--------------------|
| HMAC tool receipts not wired through all agent turn paths (ACP, gateway WS, CLI) | Medium | Receipt generation was scoped only to channel orchestrator, missing other paths | [#8009](https://github.com/zeroclaw-labs/zeroclaw/pull/8009) |
| `cron_update` and `cron_remove` only accepted UUIDs, not job names | Medium | Forced extra round‑trip call for human‑readable names | [#8007](https://github.com/zeroclaw-labs/zeroclaw/pull/8007) |
| Cost budget config frozen at boot (not reloadable) | Medium | `CostConfig` held in `OnceLock` prevented runtime updates | [#8004](https://github.com/zeroclaw-labs/zeroclaw/pull/8004) |
| Session termination hook never fired (`session_end` was dead code) | Medium | Existing `HookRunner::fire_session_end()` never called | [#8003](https://github.com/zeroclaw-labs/zeroclaw/pull/8003) |
| Codex JWT account extraction missed `chatgpt_account_id` claim | Low | Account ID preference order updated | [#8002](https://github.com/zeroclaw-labs/zeroclaw/pull/8002) |
| Memory embedder had no dedicated API key (borrowed from first provider) | Medium | Fixed by decoupling embedding key from chat provider | [#7942](https://github.com/zeroclaw-labs/zeroclaw/pull/7942) |
| Non‑opus TTS not delivered correctly to Telegram voice peers | Medium | Fixed `sendAudio` wiring and agent transcription provider | [#7019](https://github.com/zeroclaw-labs/zeroclaw/pull/7019) |
| Custom model providers not validated correctly by `zeroclaw-doctor` | Low | Made helper public and updated error display | [#8010](https://github.com/zeroclaw-labs/zeroclaw/pull/8010) |

Additionally, a new issue **#8012** tracks removal of the legacy log cursor now that byte‑offset pagination is primary (cleanup, low severity).

## Feature Requests & Roadmap Signals

Today’s issue and PR activity reveals several features likely to land in the next minor releases (v0.8.2 / v0.8.3 / v0.9.0):

- **Scope‑selectable `/model` overrides** (#7998) – Users can target a model change at the chat, user, or agent level. Highly requested for managing multiple chat contexts under one bot.
- **Restored per‑sender `/thinking` overrides** (#8011) – Brings back `/thinking` commands for configuring reasoning level per conversation.
- **Email‑login auth via OAuth2** (#8008) – Adds device authorization flow for email channels, a stepping stone to broader auth (tracked in #7432).
- **WASM component‑model plugin host** (#7928) – Initial code for plugin isolation using WASM, an important architectural addition for extensibility.
- **Per‑turn output routing (RFC‑6969)** (#7361) – Large feature for sending agent responses to different destinations (e.g., voice‑only peers), fixing double‑send bugs and adding voice delivery.
- **Discord rich embeds** (#7833) – Renders `[EMBED:{…}]` markers as native Discord embeds, part of the parity tracker #7831.
- **ZeroCode UI improvements** (#8000, #8006) – Browse mode badge, Aliases/Costs tabs in provider alias list, and other quality‑of‑life enhancements for the TUI.
- **Configurable temp‑file cleanup** (#7996) – Addresses storage‑constrained deployments (low‑end devices, embedded systems).

Given the project’s milestone trackers, **v0.8.3** (MCP dashboard, web/plugin surfaces) and **v0.9.0** (auth, security, breaking changes) are the next major targets.

## User Feedback Summary

User feedback, extracted from issue descriptions and PR motivations, highlights the following pain points and use cases:

1. **Session persistence** – #7759: Users expect that temporarily losing a WebSocket connection should not abort an ongoing agent turn. Real‑time chat applications and long‑running tool calls are the driving use case.
2. **Model flexibility** – #6557: Channel slash commands (`/model`, `/models`, `/config`) are confusing when multiple providers are configured. Users want predictable, granular control without breaking provider semantics.
3. **Discord feature gap** – #7831: Discord users cannot send rich embeds, use typed slash options, or handle voice interactions to the same degree as other channel integrations. The tracker is the most active community coordination point.
4. **Prebuilt bundle completeness** – #7952: Users who configure non‑default channels (e.g., Telegram, Discord, Slack) find that the default prebuilt binary does not include those channel runtimes, causing confusion and extra build steps.
5. **Resource management** – #7996: Deployments on Raspberry‑Pi‑like devices accumulate temporary files (attachments, snapshots, downloads) that eventually fill storage. Configurable cleanup is requested.

Overall satisfaction appears high given the volume of contributions, but the above issues indicate a desire for **production‑ready session handling**, **full channel support without compilation**, and **mature Discord integration**.

## Backlog Watch

Several items require maintainer or author attention to unblock progress:

- **#7952** – Publish full‑channel prebuilt assets. Labeled `needs-maintainer-review`; this enhancement with high risk has no corresponding PR yet.  
  🔗 [Issue #7952](https://github.com/zeroclaw-labs/zeroclaw/issues/7952)

- **#7833** (PR) – Discord embeds (XL, risk medium). Labeled `needs-author-action`; the author must address review feedback before merging.  
  🔗 [PR #7833](https://github.com/zeroclaw-labs/zeroclaw/pull/7833)

- **#7019** (PR) – Telegram voice fixes (L, risk medium). Also labeled `needs-author-action`; waiting on author response.  
  🔗 [PR #7019](https://github.com/zeroclaw-labs/zeroclaw/pull/7019)

- **#6557** – Model switching reconciliation (p2, in‑progress since May 10). Two comments but no recent updates; may be stalled despite acceptance.  
  🔗 [Issue #6557](https://github.com/zeroclaw-labs/zeroclaw/issues/6557)

These items, if left unattended, could delay delivery of v0.8.2 or v0.8.3. The community would benefit from timely maintainer responses and author revisions.

</details>

---
*This digest is auto-generated by [agents-radar](https://github.com/ivo-eu/agents-radar).*