# OpenClaw 生态日报 2026-06-14

> Issues: 230 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-06-14 11:10 UTC

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

## OpenClaw 项目深度报告

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，现根据 OpenClaw 项目在 2026 年 6 月 14 日的 GitHub 数据，为您呈上项目动态日报。

---

### **OpenClaw 项目日报 | 2026-06-14**

#### **1. 今日速览**

OpenClaw 项目今日处于**极高活跃度**状态。社区贡献踊跃，过去 24 小时内产生了 230 条 Issue 更新和 500 条 PR 更新，刷新了近期活跃度记录。项目发布了 `v2026.6.8-beta.1` 版本，重点增强了 Telegram 和 WhatsApp 等消息渠道的健壮性与功能。然而，社区反馈主要集中于**稳定性**和**会话状态可靠性**，特别是多个活跃 Issue 指向了严重的内存泄漏和 Agent 逻辑缺陷。项目在修复关键 Bug（如子 Agent 的 `thinking` 参数问题）上有所进展，但仍有大量高优先级问题等待解决。

---

#### **2. 版本发布**

**版本: `v2026.6.8-beta.1`**
- **链接:** [openclaw/openclaw Releases v2026.6.8-beta.1](https://github.com/openclaw/openclaw/releases/tag/v2026.6.8-beta.1)
- **主要内容:**
    - **Telegram 频道增强:** 支持发送结构化富文本（表格、列表）和可折叠引用。改进了 CLI 后端交付和原生草稿迁移的健壮性。
    - **WhatsApp 频道增强:** 提升了消息发送的健壮性。
    - **安全性/边界处理:** 优化了富媒体边界处理，防止解析错误。
- **破坏性变更/迁移注意事项:** 本次发布为 Beta 版本，未明确指出破坏性变更，但建议用户关注 Telegram 富文本解析规则的变化，确保自定义格式能正确渲染。

---

#### **3. 项目进展**

今日，项目在修复核心 Agent 行为和系统稳定性方面取得了关键进展，合并/关闭了 134 项 PR。

- **`sessions.json` 无限制增长问题:** 针对 Issue #55334（会话文件导致 OOM），PR #91712 已被合并，通过清理非空闲的旧诊断条目来缓解日志膨胀。这是一个对长期运行服务至关重要的修复。
- **GPT-5.5 Codex 超时问题:** Issue #89996 被关闭，该问题涉及特定模型在完成工具调用后超时，相关修复已合入。
- **子 Agent `thinking` 参数失败问题:** PR #92914 已打开并标记为“ready for maintainer look”，旨在修复子 Agent 在请求不支持的 `thinking` 级别时静默失败的问题。这直接回应了社区在 #92412 中的反馈。
- **记忆索引状态同步:** PR #92943 已打开，修复了外部重索引后内存网关可能使用旧句柄的问题，提升了记忆子系统的数据一致性。
- **CLI 多通道交付:** PR #89074 被提出，旨在修复 TUI 在活跃运行时发送被阻塞的问题，确保后端策略能正常处理 `/queue` 命令。

**项目整体向前迈进:** 核心稳定性（内存泄漏、日志管理）和 Agent 行为（子 Agent 逻辑、模型切换）是今日的修复重点。

---

#### **4. 社区热点**

今日社区讨论热度最高的议题集中在**AI Agent 的可靠性与透明度**以及**系统稳定性**。

1.  **Agent “虚假承诺”问题 (Issue #58450):**
    - **链接:** [Issue #58450](https://github.com/openclaw/openclaw/issues/58450)
    - **热度:** 15 条评论, 3 个 👍。
    - **分析:** 用户反馈 Agent 会以“我会检查项目记忆并稍后回复”结束对话，但实际上并未启动任何后台任务。这严重影响了用户对 Agent 完成度的信任，揭示了 Agent 在任务调度和状态机管理上的深层缺陷。该问题被标记为“铂金海螺”级，表明其复杂性和高影响。

2.  **网关进程严重内存泄漏 (Issue #91588):**
    - **链接:** [Issue #91588](https://github.com/openclaw/openclaw/issues/91588)
    - **热度:** 11 条评论, P0 优先级。
    - **分析:** 用户报告网关在 2-3 天内 RSS 从 350MB 增长到 15.5GB，最终导致 OOM 崩溃。这是影响所有长期运行实例的“杀手级” Bug。用户对其稳定性的担忧非常强烈。

3.  **Control UI Raw 模式被禁用 (Issue #59330):**
    - **链接:** [Issue #59330](https://github.com/openclaw/openclaw/issues/59330)
    - **热度:** 8 条评论, 14 个 👍。
    - **分析:** 这个回归 Bug 获得了今日最高的点赞数。用户对 UI 功能的下降非常敏感，`Raw` 模式是高级用户进行配置调试的必备功能，其被强制禁用严重影响了他们的工作流。

---

#### **5. Bug 与稳定性**

今日报告的 Bug 中，稳定性和会话状态问题是重灾区。

**严重 Bug (P0/P1):**
- **`[P0]` 网关进程内存泄漏:**
    - **链接:** [Issue #91588](https://github.com/openclaw/openclaw/issues/91588)
    - **描述:** RSS 内存从 350MB 激增至 15.5GB 并导致 OOM 崩溃。
    - **状态:** 无 Fix PR。

- **`[P1]` `sessions.json` 无限制增长导致 OOM:**
    - **链接:** [Issue #55334](https://github.com/openclaw/openclaw/issues/55334)
    - **描述:** 会话文件条目从不修剪，导致内存消耗飙升。
    - **状态:** 有相关的 Fix PR（#91712）已在日志层面缓解问题，但根因（会话条目不修剪）可能仍需处理。

- **`[P1]` 子 Agent `thinking` 参数静默失败 (Issue #92412子问题):**
    - **链接:** [PR #92914](https://github.com/openclaw/openclaw/pull/92914) (解决方案)
    - **描述:** `sessions_spawn` 在请求不支持的 `thinking` 级别时，错误地返回成功状态，但实际未生效。
    - **状态:** 已提出修复，等待审核。

**中等严重 Bug (P2):**
- **Safeguard Compaction 模型配置被忽略:**
    - **链接:** [Issue #57901](https://github.com/openclaw/openclaw/issues/57901)
    - **描述:** `compaction.model` 配置被忽略，转而使用会话主模型。
- **Agent 换行后无后续动作:**
    - **链接:** [Issue #58450](https://github.com/openclaw/openclaw/issues/58450)
    - **描述:** Agent 承诺跟进，但实际上没有启动任何操作。
- **Google Chat 组消息被静默忽略:**
    - **链接:** [Issue #58514](https://github.com/openclaw/openclaw/issues/58514)
    - **描述:** 仅在组/空间消息中失效，DMs 正常。
- **更新后 UI 问题:**
    - **链接:** [Issue #58702](https://github.com/openclaw/openclaw/issues/58702)
    - **描述:** WebChat 消息文本被操作图标遮挡。

---

#### **6. 功能请求与路线图信号**

今日用户提出的功能请求指向了更成熟的 AI Agent 能力和更细粒度的控制。

- **增强 Agent 会话管理:**
    - **请求:** 支持会话标签/昵称 (Issue #55249) 和状态文件 (`STATE.md`) 以在压缩后保留上下文 (Issue #56110)，表明用户希望更直观地管理多个会话并提升 Agent 的长期记忆能力。
- **严格的出站策略:**
    - **请求:** “Unbypassable outbound policy enforcement” (Issue #56349) 被提出，要求所有出站消息必须经过一个统一的、不可绕过的策略边界，这反映了企业级或高安全需求场景下的核心诉求。
- **新渠道/功能探索:**
    - **请求:** 支持 Telegram Inline Query (Issue #54794) 和 Agent 主动拨打外呼电话 (Issue #59245)。`feat(typing)` (PR #58887) 要求在语音消息接收后立即显示打字指示器，体现了对用户体验细节的精益求精。
- **架构借鉴 Claude Code:**
    - **请求:** 社区积极分析并建议借鉴 Claude Code 的开源代码，提出引入多层压缩架构 (Issue #58398) 和 `exec()` 沙箱隔离 (Issue #58730)。这表明社区希望 OpenClaw 在核心安全和上下文处理能力上向业界最佳实践看齐。

**可能纳入下一版本的功能:** 基于现有 PR 和 Issue 热度，`feat(typing)` (PR #58887)、会话标签 (Issue #55249) 以及参考 Claude Code 的隔离特性 (Issue #58730) 是很有可能被优先引入的。

---

#### **7. 用户反馈摘要**

- **痛点聚焦:**
    - **信任与一致性:** 用户对 Agent 的“虚假承诺”（Issue #58450）感到沮丧，认为这破坏了 agent 的可信度。
    - **稳定性焦虑:** 内存泄漏 (Issue #91588, #55334) 导致服务频繁崩溃，用户担忧长期运行的可靠性。
    - **安全和隐私:** 代码中包含硬编码路径 (Issue #51429) 引发了用户对项目治理和代码审核流程的质疑。
    - **功能回退:** Control UI Raw 模式被禁用 (Issue #59330) 这类回归问题，直接导致了用户对版本升级持谨慎态度。
- **使用场景:**
    - **企业级使用:** 多工作区 Slack (Issue #58523)、严格的安全策略 (Issue #56349) 等需求，表明项目正被用于生产环境。
    - **开发者驱动:** 对 Claude Code 源代码的分析和模仿 (Issue #58398)，显示出用户具备较高的技术水平和改进意愿。

---

#### **8. 待处理积压**

以下 Issue/PR 长期未被响应或解决，可能正在阻碍项目健康发展。

- **`[P2]` Safeguard Compaction模型配置被忽略 (Issue #57901):**
    - **链接:** [Issue #57901](https://github.com/openclaw/openclaw/issues/57901)
    - **问题:** 该问题已存在近 3 个月，影响使用了自定义 `compaction.model` 的用户，可能导致数据丢失或不期望的模型调用。
- **`[P2]` 高赞 Issue (14 👍) - Control UI Raw 模式被禁用 (Issue #59330):**
    - **链接:** [Issue #59330](https://github.com/openclaw/openclaw/issues/59330)
    - **问题:** 该 Issue 获得大量社区共鸣，但状态仍为 `[OPEN]`，虽然已被标记挂上 Fix PR，但进展不明。
- **`[P1]` 清理不稳定消息渠道问题 (如 #58523, #57447):**
    - **链接:** [Issue #58523](https://github.com/openclaw/openclaw/issues/58523), [Issue #57447](https://github.com/openclaw/openclaw/issues/57447)
    - **问题:** 多个渠道（Slack多工作区、跨Agent消息）的入站消息处理存在问题，影响了核心的消息路由与通信功能。
- **长期未合并的 PR:**
    - **`feat(slack): add ignoreOtherMentions` (PR #53467):** 一个已准备好被维护者查看的 Slack 功能，自 3 月开放至今，可能影响工作组内其他成员对 Slack bot 的使用体验。
    - **`feat: include provider/model/... context in error messages` (PR #55851):** 一个旨在提升诊断效率的 PR，已准备好近 3 个月，其延迟合并可能会拖慢运营排查问题的速度。

---
**分析师总结:** OpenClaw 项目正处在一个快速迭代但稳定性承压的关键时期。社区贡献热情高涨，但同时也暴露出核心系统在资源管理（内存泄漏）和 Agent 逻辑（任务兑现）上的短板。维护者需优先处理 P0/P1 级别的稳定性问题，并加快对高赞社区问题（如 Raw 模式回归）的响应速度，以巩固用户信任。

---

## 横向生态对比

好的，作为AI智能体与个人AI助手开源生态的资深技术分析师，基于您提供的2026-06-14各项目动态摘要，我为您呈现以下横向对比分析报告。

---

### **2026年6月14日：个人AI助手开源生态横向对比分析报告**

#### **1. 生态全景**

当前个人AI助手/自主智能体开源生态正处于 **“能力爆发与稳定性阵痛并存的青春期”**。一方面，社区创新活力空前，多模型支持、平台集成（Telegram、Slack、微信等）、多Agent协作与记忆持久化成为普遍追求；另一方面，核心系统的资源管理（内存泄漏）、运行可靠性（任务静默失败、上下文丢失）和交互透明度（“虚假承诺”、审批冗余）成为阻碍用户体验进一步跃升的关键瓶颈。生态内项目在技术路线上呈现分化：以OpenClaw为代表的“整合派”追求全功能、高集成的单体平台，而以NanoBot为代表的“轻量派”则强调模块化与易用性，社区正围绕“全能”与“专注”两种哲学进行探索。

#### **2. 各项目活跃度对比**

| 项目名称 | 今日Issues(新开/更新) | 今日PRs(新/合并) | 版本发布 | 健康度评估 |
| :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 230 (极高) | 500 (极高) | Beta 版本 | **迭代狂热，但稳定性承压**。社区贡献爆炸，但P0/P1级内存泄漏与Agent逻辑缺陷威胁生产部署。 |
| **NanoBot** | 6 (中) | 20+ (高) | 无 | **高强度质量加固**。重点修复数据校验、路径安全等边界问题，处于密集迭代期。 |
| **Hermes Agent** | 11 (中高) | 50 (高) | 无 | **跨平台兼容性攻坚**。集中修复Windows数据丢失、Gateway重启一致性等平台级Bug，进展显著。 |
| **PicoClaw** | 2 (低) | 7 (中) | Nightly (1) | **小步快跑，隐患犹存**。修复了图像幻觉Bug，但Token持续消耗问题待解决。 |
| **NanoClaw** | 1 (低) | 10 (中高) | 无 | **基础设施奠基**。重点构建Provider能力框架与持久化内存，为后续功能扩展打基础。 |
| **NullClaw** | 1 (低) | 1 (低) | 无 | **问题僵持，信心危机**。一个严重Bug（定时任务静默失败）持续13天无官方回应，社区信任度面临挑战。 |
| **IronClaw** | 11 (中) | 21 (高) | 无 | **“Slack+审批”专项修复冲刺**。核心开发团队针对Reborn平台进行集中修复，但E2E测试持续失败。 |
| **LobsterAI** | 0 | 0 (1个历史PR合并) | 无 | **社区“静默”期**。无新动态，但积压了3个与核心体验（Cowork模块增强）相关的PR，贡献者热情可能消退。 |
| **CoPaw** | 9 (中) | 12 (中高) | 无 | **活跃的社区共建**。大量来自首次贡献者的PR，专注于UI多语言、细节修复，生态健康度高。 |
| **Moltis** | 1 (低) | 3 (低) | 无 | **快速响应典型**。昨日报的严重Bug（MCP OAuth认证失败），今天就有修复PR，效率值得肯定。 |
| **TinyClaw** | 0 | 0 | 无 | **沉寂** |
| **ZeptoClaw** | 0 | 0 | 无 | **沉寂** |
| **ZeroClaw** | 4 (中) | 50 (极高) | 无 | **代码规范与架构梳理**。大量PR集中在消除`unwrap`、统一日志，并完成了一项重要的内部引擎统一工作。 |

#### **3. OpenClaw在生态中的定位**

*   **优势与定位**：作为核心参考项目，OpenClaw是当前生态中 **功能最全面、社区最活跃** 的“航空母舰”。其社区规模（230 Issue/500 PR）远超其他项目，是生态创新的主要策源地。技术路线偏向 **全栈整合**，统一管理会话、记忆、Agent、渠道，追求开箱即用的完整体验。
*   **技术路线差异**：与 **NanoBot** 的模块化、轻量化（侧重数据校验与稳健性）和 **IronClaw** 的企业级审批流（侧重安全与合规）相比，OpenClaw更像一个“万能框架”，牺牲了一定深度以换取广度。
*   **核心隐忧**：社区繁荣的背后，是P0级内存泄漏（#91588）和Agent“虚假承诺”（#58450）等动摇根基的稳定性问题。这使其在追求极致稳定性的企业级用户眼中，信任度可能不如专注于特定场景的 **NanoBot** 或 **IronClaw**。

#### **4. 共同关注的技术方向**

| 技术方向 | 涉及项目 | 具体诉求与表现 |
| :--- | :--- | :--- |
| **稳定性与资源管理** | OpenClaw, NanoBot, PicoClaw, NullClaw, CoPaw | **内存泄漏** (OpenClaw/15.5GB增长, NanoBot/sessions.json膨胀)；**数据丢失/静默失败** (OpenClaw/虚假承诺, NullClaw/定时任务, CoPaw/上下文丢失)；**资源滥用** (PicoClaw/Token持续消耗)。 |
| **平台扩展与多形态交互** | OpenClaw, Hermes Agent, CoPaw, Moltis, ZeroClaw | **深度集成** (Hermes/Telegram AI Bot, CoPaw/钉钉, Moltis/MCP OAuth)；**非文本交互** (Hermes/WhatsApp语音, PicoClaw/TTS增强)；**跨渠道一致** (Hermes/微信个人号适配, OpenClaw/Google Chat修复)。 |
| **安全、隐私与合规** | OpenClaw, Hermes Agent, IronClaw, ZeroClaw | **出站策略** (OpenClaw/“Unbypassable outbound policy enforcement”, IronClaw/审批流程)；**配置安全** (Hermes/Gateway外部安全默认配置)；**代码质量** (ZeroClaw/消除`unwrap`防止panic)。 |
| **用户体验与交互透明** | OpenClaw, NanoBot, CoPaw, ZeroClaw | **Agent可靠** (OpenClaw/“虚假承诺”，影响信任)；**通知与反馈** (ZeroClaw/通道压缩前通知, CoPaw/会话内搜索，IronClaw/明确拒绝而非后台排队)；**国际化** (CoPaw/越南语支持, PicoClaw/繁中支持)。 |
| **多模型与记忆管理** | OpenClaw, NanoClaw, IronClaw, Moltis | **Provider能力框架** (NanoClaw/`capability seam`)；**持久化记忆** (NanoClaw/`usesMemoryScaffold`)；**向量库优化** (Moltis/纯Rust `turbovec` 后端)。 |

#### **5. 差异化定位分析**

| 对比维度 | OpenClaw (全能旗舰) | NanoBot (稳健基石) | IronClaw (企业安全) | NullClaw (专注场景) | CoPaw (社区体验) |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **功能侧重** | 全功能整合 (Agent, 渠道, 记忆) | 核心健壮性 (数据校验, 边界处理) | 审批流与安全 (Reborn平台) | 定时任务与Agent自动化 | UI/UX打磨与国际化 |
| **目标用户** | AI爱好者/开发者/深度使用者 | 对稳定性和资源有高要求的开发者 | 企业级用户、合规需求强的团队 | 需要定时触发Agent场景的用户 | 多语言用户、注重交互体验的用户 |
| **技术架构** | 单体应用，全栈整合 | 模块化，侧重底层稳健 | 微服务/平台化 (Reborn, Gateway) | 简洁，功能聚焦 | 轻量级，快速迭代 |
| **当前阶段** | 快速迭代，稳定性承压 | 质量加固，稳步前进 | 专项修复，架构优化 | Bug修复，社区停滞 | 社区共建，活力充沛 |

#### **6. 社区热度与成熟度**

*   **第一梯队 (高速迭代/“爆炸”阶段)**：**OpenClaw**。社区贡献量惊人，但问题数量也多，处于“野蛮生长”期。**ZeroClaw** 也属此类，但更偏向内部代码重构。
*   **第二梯队 (活跃迭代/质量巩固阶段)**：**NanoBot**、**Hermes Agent**、**IronClaw**、**CoPaw**、**NanoClaw**。这些项目有清晰的迭代方向，社区活跃且聚焦，Bug修复迅速，是生态中最稳定、产出最扎实的部分。
*   **第三梯队 (温和发展/静默阶段)**：**PicoClaw**、**Moltis**。有一定活跃度但幅度不大，主要在解决特定问题或进行小步改进。
*   **第四梯队 (停滞/危险阶段)**：**NullClaw**、**LobsterAI**、**TinyClaw**、**ZeptoClaw**。社区活动极低，关键Bug长期未解决，有被边缘化的风险。尤其是 **NullClaw**，一个严重Bug13天无官方回应，社区信心正在流失。

#### **7. 值得关注的趋势信号**

1.  **从“能用”到“可靠”的范式转变**：用户已不再满足于AI助手能回答问题，而是追求 **“任务兑现率”** 和 **“确定性”**。OpenClaw的“虚假承诺” (Issue #58450) 和NullClaw的“静默失败” (Issue #941) 被视为信任危机，这预示着下一个竞争焦点将是Agent工作流的**可审计性、可预测性和自愈能力**。

2.  **内存泄漏成为“杀手级”共性BUG**：OpenClaw (15.5GB) 和NanoBot (sessions.json) 的内存泄漏问题表明，对于需要长期运行的个人AI助手，资源管理是比模型能力更严峻的技术挑战。未来将催生更高效的**上下文压缩、增量存储和内存健康监控**方案。

3.  **“热猫腻”模型支持成为基本盘，而非差异化优势**：几乎所有活跃项目都在积极适配新模型（GLM-5.2, Claude Opus等）。当模型接入变得像插件一样简单，项目的核心竞争力将转向**独特的数据管理、安全的工具调用框架和流畅的用户交互设计**。

4.  **社区开始向DeepCode/Claude Code等前沿架构“取经”**：OpenClaw社区明确提出借鉴Claude Code的**多层压缩架构**和`exec()`沙箱隔离。这预示着个人AI助手社区不再闭门造车，而是主动吸收业界最佳实践，**安全、高效的代码执行环境**将成为下一个技术高点。

**对AI智能体开发者的参考价值**：在构建您的AI智能体时，请优先关注**系统稳定性**（尤其是内存和任务调度的可靠性）和**交互透明度**（让用户明确知道Agent的状态和决策依据）。其次，应关注**平台扩展的深度**（如Telegram新特性）而非广度。最后，**向Claude Code学习**其安全、高效的设计哲学，将是提升您产品竞争力的关键。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

好的，这是为您生成的 NanoBot 项目动态日报（2026-06-14）。

---

# NanoBot 项目动态日报 — 2026-06-14

## 1. 今日速览

今日项目活跃度极高，社区贡献与核心维护并行推进。过去24小时内，共处理了31个Pull Request，其中11个已合并/关闭，20个处于待审核状态，显示出密集的代码迭代和功能交付节奏。Issues方面，有6个进行了更新，5个已关闭，问题响应迅速。值得注意的是，**稳定性与安全性修复**成为今日的核心主题，多个PR聚焦于数据校验、路径安全及会话管理的边界情况处理。此外，WebUI方面也迎来了多项功能性改进，包括自动化管理视图和移动端响应式适配。整体而言，项目处于一个**高强度的迭代与质量加固周期**。

## 2. 版本发布

无新版本发布。

## 3. 项目进展

今日合并/关闭了多个关键PR，显著推动了项目的稳定性和功能性。

- **核心稳定性提升**：合并了修复 `idleCompact` 机制的 `#4270`。该修复解决了之前仅对部分对话历史进行总结的问题，现在可以**对完整的会话历史进行存档**，确保了对话存档的准确性和可靠性。
  - [PR #4270 fix: archive full session history in idle compact](https://github.com/HKUDS/nanobot/pull/4270)

- **WebUI 功能增强**：
    - 合并了提供跨平台本地化支持的 `#4331`，为“设置”中的“关于/更新检查”页面添加了多语言文案，提升了国际化体验。
      - [PR #4331 [valid, webui] fix(webui): localize update check copy](https://github.com/HKUDS/nanobot/pull/4331)
    - 合并了`#4338`，为README中的Kimi合作伙伴Banner添加了可直接点击的链接，改善了外部合作伙伴的导流体验。
      - [PR #4338 docs(readme): link Kimi partner banner](https://github.com/HKUDS/nanobot/pull/4338)
    - 合并了合作伙伴文档更新 `#4295`，在配置文档中增加了Kimi的附链接，并在README中新增了“开源合作伙伴”板块。
      - [PR #4295 docs: add Kimi and MiniMax partner links](https://github.com/HKUDS/nanobot/pull/4295)

- **数据健壮性修复**：合并了多个由 `yu-xin-c` 提交的修复，针对恶意或格式错误的输入数据进行校验，如忽略空的注入内容、拒绝不合法的路径链接和分页参数等，从系统底层增强了项目的稳健性。
  - [PR #4337 fix(runner): ignore empty injected payloads](https://github.com/HKUDS/nanobot/pull/4337)
  - [PR #4336 fix(cli-apps): reject malformed argv payloads](https://github.com/HKUDS/nanobot/pull/4336)
  - [PR #4315 fix(memory): ignore malformed history entries](https://github.com/HKUDS/nanobot/pull/4315)
  - [PR #4312 fix(message): reject malformed media attachments](https://github.com/HKUDS/nanobot/pull/4312)
  - [PR #4311 fix(tools): reject non-positive file pagination limits](https://github.com/HKUDS/nanobot/pull/4311)

## 4. 社区热点

今日社区讨论最活跃的议题仍然是**对更多模型提供商的支持**。

- **[Issue #193 Ollama api support?](https://github.com/HKUDS/nanobot/issues/193)**：该问题自2026年2月提出，共有15条评论，并在今日再次更新。用户明确提出了对Ollama API的支持需求，表示当前只看到vLLM的支持。这反映了社区用户对于**拥抱本地轻量级推理方案（如Ollama）的强烈诉求**，旨在降低使用门槛和成本。

- **[PR #4330 feat(webui): add automation management view](https://github.com/HKUDS/nanobot/pull/4330)**：由核心贡献者 `chengyongru` 提交的新PR，为WebUI增加了自动化管理视图。尽管评论数未明确，但从其“添加列表、过滤、运行、暂停/恢复、删除自动化”的摘要描述来看，这显著提升了用户对自动化流程的可视化控制能力，是社区期待的高级管理功能。

## 5. Bug 与稳定性

今日报告的严重BUG较少，大部分已通过PR快速解决。

- **[Issue #4333 Anthropic provider sends deprecated `temperature` to opus-4-8 / Fable](https://github.com/HKUDS/nanobot/issues/4333)** （严重，已关闭）：
  - **问题**：Anthropic 提供商仅在 `opus-4-7` 模型中抑制了已废弃的`temperature`参数，导致新模型 `opus-4-8` 和 `Fable` 的 API 请求被拒绝（返回400）。
  - **影响**：影响所有使用新版 Claude Opus 模型的用户，导致API请求全面失败。
  - **状态**：已于今日关闭，表明已修复。

- **[Issue #4322 [question] NameError: 'session_key' is not defined](https://github.com/HKUDS/nanobot/issues/4322)**（严重，待解决）：
  - **问题**：在合并 `fix/prompt-caching` 分支后，用户报告了一个 `NameError`，导致Agent在启动时崩溃。根因指向了一次合并操作中提取新方法时的代码错误。
  - **影响**：严重，影响开发者使用该特定分支。目前为开放状态，尚无明确的修复PR与之关联。

- **[Issue #4264 [bug] idleCompact should use the complete session history](https://github.com/HKUDS/nanobot/issues/4264)** （中等，已关闭）：
  - **问题**：`idleCompact` 机制仅对会话中的最后8条消息进行存档，忽略了对用户纠正行为的总结，可能导致`history.jsonl`中保留错误的记录。
  - **影响**：影响对话存档的准确性。
  - **状态**：已由 [PR #4270](https://github.com/HKUDS/nanobot/pull/4270) 修复并关闭。

## 6. 功能请求与路线图信号

- **Ollama API 支持**：Issue #193 是一个长期未解决且社区呼声很高的功能请求。考虑到NanoBot对本地部署和轻量化推理的潜在定位，将此功能纳入下一版本的可能性较高。
- **WebUI 深度升级**：今日的PR清晰地展示了WebUI的发展方向：
  - **自动化管理界面** (PR #4330) 和 **移动端响应式适配** (PR #4339) 的提出，表明项目正致力于打造一个**完整、现代化且全平台兼容的用户界面**。这些功能很可能被优先合并入下一个版本。
  - **本地化完善**：PR #4331 和 #4338 的工作表明项目对国际化体验的持续投入，这通常是成熟项目的重要标志。

## 7. 用户反馈摘要

- **痛点**：
  - 模型支持不足：用户 (neil0306) 明确表示需要Ollama支持，因为vLLM并非唯一选择，这限制了本地使用场景。
  - 会话存档不准确：用户 (imkuang) 在 #4264 中详细描述了 `idleCompact` 机制的缺陷，指出当前的“尾巴切除”策略会**丢失用户最后的关键纠正信息**，导致错误的存档结论。
  - 协作开发风险：用户 (professionelle-hypnose) 在 #4322 中遇到了分支合并导致的应用崩溃问题，这反映了在代码快速迭代过程中，分支合并可能引入的稳定性风险。
- **使用场景**：
  - 多数用户（如 #193, #4264）处于“对话纠错”和“结果验证”的典型AI辅助工作流中，对对话历史的完整性和准确性有较高要求。

## 8. 待处理积压

- **[Issue #193 Ollama api support?](https://github.com/HKUDS/nanobot/issues/193)**：自2026年2月提出至今已4个多月，仍然是社区共鸣最强的功能需求。建议项目维护者评估其优先级并给出明确的路线图回应。
- **[PR #4053 fix(tools): keep read-only roots out of write paths](https://github.com/HKUDS/nanobot/pull/4053)**：由 `yu-xin-c` 提交，旨在防止安全绕过的重要修复。自5月29日提交，已开放超过两周，值得关注。
- **[PR #4119 fix(exec): block relative symlink workspace escapes](https://github.com/HKUDS/nanobot/pull/4119)**：另一个重要的安全检查，防止通过符号链接逃逸工作区。同样来自 `yu-xin-c`，开放两周有余，建议优先审阅。
- **[Issue #4322 NameError: 'session_key' is not defined](https://github.com/HKUDS/nanobot/issues/4322)**：虽然被称为“question”，但实际上是一个会导致应用崩溃的严重BUG。目前仅有问题报告，无修复PR，需维护者立即介入。

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

好的，作为AI智能体与个人AI助手领域开源项目分析师，这是根据您提供的Hermes Agent项目数据生成的2026年6月14日项目动态日报。

---

## Hermes Agent 项目动态日报 | 2026年06月14日

### 1. 今日速览

今日项目活跃度极高，共处理了13个Issue和50个Pull Request，社区讨论和技术开发均非常活跃。项目维护者表现出强大的代码合并与问题解决能力，今日成功合并/关闭了11个PR和4个Issue，其中包含多个P1级别的关键Bug修复（如Gateway重启状态恢复、消息序列化问题）。社区功能请求呈现出鲜明的“平台扩展”与“多形态交互”趋势，特别是对Telegram和微信（WeChat）生态的深度集成需求显著。总体来看，项目处于一个高速迭代和持续优化的健康状态。

### 3. 项目进展

今日项目在稳定性和平台兼容性方面取得了重要进展，多项关键Bug修复被合并。

- **修复Windows平台数据丢失Bug**: PR [#45793](https://github.com/NousResearch/hermes-agent/pull/45793) 已被合并。该PR解决了Windows系统上桌面应用重启或崩溃时，用户自定义配置文件 (`config.yaml`) 和 `.env` 文件被重置或删除的关键数据丢失问题。此修复极大地提升了Windows用户的使用安全性。
- **修复Gateway重启消息丢失Bug**: 多个关联的Bug修复被合并或关闭。
    - Issue [#46053](https://github.com/NousResearch/hermes-agent/issues/46053) 及其修复PR（推断）已关闭。该问题涉及长对话中，AI助手的回复在 `repair_message_sequence` 后沉默丢失，根本原因是 `conversation_history` 与 `messages` 不匹配。相关修复确保了对话上下文的一致性。
    - Issue [#45682](https://github.com/NousResearch/hermes-agent/issues/45682) 已关闭。修复了Gateway重启时，`auto-resume` 并发处理与入站消息处理间的结构性问题（竞争条件），防止了相同会话的AI实例被重复创建。
- **增强Gateway状态查询**: PR [#46091](https://github.com/NousResearch/hermes-agent/pull/46091) 被提出。该PR修复了 `/model` 命令在模型因限速（429错误）触发后备模型时，仍显示已失效的配置模型的问题。现在该命令能正确显示当前实际在线的模型，提高了运维透明度。
- **跨进程会话缓存一致性**: PR [#46095](https://github.com/NousResearch/hermes-agent/pull/46095) 被提出。此修复解决了当Dashboard等外部进程通过状态数据库修改会话时，导致AI Agent的内存副本（Transcript）变为脏数据的问题，确保了不同进程间状态的一致性。

### 4. 社区热点

今日社区讨论焦点集中在**平台集成与多模态交互**上。

- **Telegram AI Bot新特性集成 (Issue #21587 - 11条评论)**: 该Issue是今日讨论次数最多的议题。用户 [`Editorenbici`](https://github.com/Editorenbici) 详细阐述了Telegram近期发布的AI Bot大更新（11项新特性），并提出了将“Guest AI Bots”、“Bot-to-Bot交互”、“贴纸与聊天自动化”等四项变革性功能集成到Hermes Agent的详细提案。这反映了社区对将Hermes Agent深度融入主流即时通讯生态，实现更复杂、更自然的多人/多Agent协作场景的强烈渴望。相关的实现PR [#43049](https://github.com/NousResearch/hermes-agent/pull/43049) 也已提出来实现“Guest Bots”功能。
- **安全与平台合规 (PR #44073 & #46028)**: 两个高评论量的开放PR也反映了社区的关注点。
    - PR [#44073](https://github.com/NousResearch/hermes-agent/pull/44073) 聚焦于Gateway的外部安全默认配置，旨在将“默认拒绝”的安全策略落实到代码层面，防止因配置疏忽导致外部未授权访问。这体现了项目对安全基线的重视。
    - PR [#46028](https://github.com/NousResearch/hermes-agent/pull/46028) 则是一个重量级的WhatsApp语音原生集成，包含语音笔记直传和Cloud Calling功能，代表了社区对“语音交互”这一非文本交互形态的探索。

### 5. Bug 与稳定性

今日报告的Bug主要集中在连接器稳定性、跨平台兼容性和核心数据持久化方面。

- **P1 严重级别**:
    - **Gateway重启消息丢失 (Issue #46088)**: 报告指出，在Gateway进程重启后，短暂的窗口内AI的响应消息会从`state.db`的会话记录中**永久丢失**。这是一个严重的会话一致性问题，后续用户消息将基于不完整的上下文生成回复。**状态**: 开放，暂无直接关联的Fix PR，但社区已高度关注。
    - **会话上下文损坏 (Issue #46090)**: 用户报告Agent执行基础任务时速度变得**极其缓慢**，无论重启还是新建会话都无法恢复。用户怀疑是上下文膨胀或持久化数据损坏，设置了压缩率为0.5也未解决。**状态**: 开放，需维护者深入调查。
    - **Assistant回复静默丢弃 (Issue #46053)**: 该问题已修复并关闭。根本原因是 `repair_message_sequence` 机制导致的 `conversation_history` 与 `messages` 不匹配，使模型无法看到自己之前的回复。
    - **非@提及消息触发空响应重试循环 (Issue #13248)**: 该Claude Opus-4模型在Slack群组中的已知Bug已被修复并关闭。问题在于模型对非@提及的讨论消息生成空回复，被Gateway解读为失败而触发无限重试。
- **P2 严重级别**:
    - **HOME环境变量污染 (Issue #25114)**: 该Bug已被修复并关闭。修复了在Hermes Profile运行时上下文中，`$HOME` 环境变量被错误地重定向到Profile目录的问题，避免了依赖`Path.home()`的Python脚本等子进程行为异常。
- **其他严重Bug (未区分P级别)**:
    - **Windows平台Unicode解码错误 (Issue #46099)**: 当Windows系统MCP子进程的输出流包含非UTF-8字节时，`TextReceiveStream` 会抛出 `UnicodeDecodeError`。**状态**: 开放，影响Windows用户使用MCP工具。
    - **Telegram Web消息显示异常 (Issue #46098)**: 更新后，通过Telegram Web访问的Agent会显示“此消息在Web版本中不受支持”，但Telegram App端正常。**状态**: 开放。
    - **Telegram批量转发消息处理时序问题 (Issue #46100)**: 当用户在Telegram中先发送简短评论，紧接着批量转发多条消息/附件时，Agent可能会在附件下载完成前就开始处理之前的评论，导致回复内容不完整。**状态**: 开放。

### 6. 功能请求与路线图信号

今日的功能请求呈现“平台扩展”和“个性化体验”两大方向。

- **潜在早期路线图信号**:
    - **Telegram深度集成 (Issue #21587)**: 社区提出的Guest Bot、Bot-to-Bot、贴纸和聊天自动化等11项新特性是明确且强烈的需求信号。结合已有PR [#43049](https://github.com/NousResearch/hermes-agent/pull/43049) 对该功能的实现，可以合理推测，Telegram Bot API 10.0的全面拥抱将成为**下一个版本的核心特性**。
    - **微信个人号适配器 (Issue #46086, PR #46096, #46093)**: 社区发现桌面UI将微信（Weixin）个人号适配器错误地标记为“微信公众号”。这一发现以及随之而来的多个修正PR，表明社区不仅在使用，而且对平台适配器的**正确性和完整性**有很高要求。对微信个人号的支持是项目中一个重要的、正在被社区积极关注的生态特性。
- **即将到来的功能**:
    - **桌面UI可选字体大小 (Issue #46097)**: 用户提出为桌面应用添加字号设置，以在高分辨率屏幕或长时间使用时改善阅读体验。这是一个提升用户体验的高价值小特性，实现门槛低，极有可能被纳入小版本。
    - **会话清理/回收机制 (Issue #46046)**: 用户抱怨Dashboard会在每次访问时创建空会话，导致会话列表膨胀。用户期望一个“Janitor/清洁工”机制来批量清理无用会话。这可能是对现有会话管理功能的一个补充。
    - **新模型支持**: PR [#46092](https://github.com/NousResearch/hermes-agent/pull/46092) 为Z.AI (智谱) 提供商添加了GLM-5.2模型支持；PR [#46094](https://github.com/NousResearch/hermes-agent/pull/46094) 为OpenRouter添加了Fusion模型支持。这表明项目正在积极跟进最新最强的LLM模型，保持对关键提供商的支持。

### 7. 用户反馈摘要

从今日的Issue评论中，可以观察到以下用户声音：

- **对数据持久化的担忧**: 用户 `kaluluosi` 在Issue [#46088](https://github.com/NousResearch/hermes-agent/issues/46088) 和 [#45682](https://github.com/NousResearch/hermes-agent/issues/45682) 中反复报告Gateway重启导致消息丢失的问题，反映出用户对“对话记录永不丢失”的**核心期望**。任何程度的会话丢失都会严重打击用户信任。
- **对性能退化感到沮丧**: 用户 `hhajj` 在Issue [#46090](https://github.com/NousResearch/hermes-agent/issues/46090) 中描述了Agent“变得极其缓慢”的体验，即使重启也无济于事。这表明用户已将Hermes Agent视为日常工具，对其性能的可靠性有很高要求。这种“无故变慢”的问题如果无法解决，会导致用户流失。
- **平台适配器正确性的重要性**: Issue [#46086](https://github.com/NousResearch/hermes-agent/issues/46086) 的作者 `52yibaobao` 精准地指出了UI标签和实际功能的错配，这不仅仅是文字错误，更可能误导用户对产品能力的认知。这反映出社区中有相当一部分**技术能力强的用户**，他们严格审查产品细节，并乐于提供精确反馈。

### 8. 待处理积压

以下是一些值得维护者关注的长时间未响应或未解决的开放式议题和PR。

- **安全PR积压严重**: 部分重要的安全修复PR已开放超过两个月，急需审查。
    - **PR [#10857](https://github.com/NousResearch/hermes-agent/pull/10857)**: 一项修复Shell注入和SSRF漏洞的P1级安全PR，已开放59天，最后一次更新是今天（2026-06-14）。考虑到其严重性，建议优先处理。
    - **PR [#21363](https://github.com/NousResearch/hermes-agent/pull/21363)**: 同样是P1级安全PR，旨在通过逐步废弃和隐藏 `v1/completions` 端点来加固系统。已开放，但需关注。
- **平台集成PR等待合并**:
    - **PR [#36286](https://github.com/NousResearch/hermes-agent/pull/36286)**: 为MiniMax中国区OAuth添加支持。对于中国区域用户至关重要，已开放13天。
- **无明确Fix PR的高严重性Bug**:
    - **Issue [#46088](https://github.com/NousResearch/hermes-agent/issues/46088) (Gateway重启消息丢失)**: 此P1级Bug目前没有明确的修复PR，其影响范围（可能导致会话永久性损坏）使其成为当前项目最需要关注的稳定性问题之一。
    - **Issue [#46090](https://github.com/NousResearch/hermes-agent/issues/46090) (Agent速度极慢)**: 此Bug影响了核心使用体验，但目前信息不足，需要维护者介入复现和定位根因。

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，根据您提供的 PicoClaw 项目数据，我为您生成了 2026-06-14 的项目动态日报。

---

### PicoClaw 项目动态日报 | 2026-06-14

---

#### 1. 今日速览

项目今日活跃度极高，共处理了 7 个 PR 和 2 个 Issue。核心进展集中在 **AI 模型兼容性修复**（TTS 与视觉模型）与 **项目可扩展性建设**（通道配置钩子）。一个关键的 Bug 修复（Issue #3108 图像幻觉）已被合并，而另一个关于 Token 持续消耗的 Bug 仍在讨论中。总体来看，项目处于快速迭代与修复的积极状态。

#### 2. 版本发布

发布了 **Nightly Build** 版本 `v0.2.9-nightly.20260614.cf67dd38`。这是一个自动构建的预发布版本，旨在让社区提前体验 `main` 分支的最新更改，但**可能不稳定**，建议谨慎使用。此次构建的变更日志可以在此查看：
- [Full Changelog: v0.2.9...main](https://github.com/sipeed/picoclaw/compare/v0.2.9...main)

#### 3. 项目进展

今日项目在多个方面取得了实质性推进，共有 5 个 PR 被合并或关闭：

- **修复视觉模型支持**：PR [#3117](https://github.com/sipeed/picoclaw/pull/3117) 已合并，解决了当活跃模型不支持视觉时，图像描述请求产生“幻觉”的问题。现在系统会自动将媒体请求路由到配置的图像模型。
- **增强 TTS 能力**：PR [#3119](https://github.com/sipeed/picoclaw/pull/3119) 已合并，为OpenRouter的TTS功能增加了 `voice` 和 `response_format` 的覆盖设置，并实现了自动重试降级机制。
- **提升代码健壮性**：两个来自 `chengzhichao-xydt` 的 PR（[#3065](https://github.com/sipeed/picoclaw/pull/3065), [#3066](https://github.com/sipeed/picoclaw/pull/3066)）被合并，通过在错误路径上显式忽略 `Close()` 错误，清理了 linter 警告，提升了代码质量。
- **推进国际化**：合并了添加繁体中文 (zh-TW) 支持的多项文档和前端配置的PR [#2935](https://github.com/sipeed/picoclaw/pull/2935)，有助于扩大项目在中文社区的覆盖面。

#### 4. 社区热点

今日最受关注的 Issue 是 **[#3012: Continuous consumption of tokens every minutes when evolution is enabled](https://github.com/sipeed/picoclaw/issues/3012)**，获得了 3 条评论。该问题讨论了在启用进化（Evolution）功能后，Token 持续消耗的问题。用户 `xpader` 提供了详细的复现步骤和环境信息，引发了社区对该功能资源消耗的关注。这反映了用户对AI助手 **成本控制** 和 **资源效率** 的敏感需求。

#### 5. Bug 与稳定性

- **[严重] Token 持续消耗** (Issue [#3012](https://github.com/sipeed/picoclaw/issues/3012)): 用户报告在开启“进化”功能后，每分钟都会持续消耗 Token。该 Bug 目前处于开启状态，尚未有相关的修复 PR，需要项目组重点关注。
- **[已修复] 图像描述幻觉** (Issue [#3108](https://github.com/sipeed/picoclaw/issues/3108)): 当模型不支持视觉时，图像描述请求会产生不相关的结果。此问题已被 PR [#3117](https://github.com/sipeed/picoclaw/pull/3117) 修复，并已关闭，状态良好。

#### 6. 功能请求与路线图信号

- **可扩展性增强**：PR [#3120](https://github.com/sipeed/picoclaw/pull/3120) 提出为“带外通道”（out-of-tree channels）增加配置注册钩子，这将成为 PicoClaw 插件化生态的重要基石，很可能会被纳入下一个正式版本。
- **图像压缩管线**：PR [#2964](https://github.com/sipeed/picoclaw/pull/2964) 提议为视觉管线增加可配置的入站图像压缩功能，以解决因图像过大导致的资源浪费问题。该 PR 仍处于开放状态，但其价值与 Issue [#3012] 关注的资源成本问题高度契合，预计未来会得到更多关注和推动。
- **TTS 灵活性与稳定性**：已合并的 PR [#3119](https://github.com/sipeed/picoclaw/pull/3119) 展示了社区对 TTS 功能个性化控制和稳定性的需求，通过 `extra_body` 提供覆盖设置并实现自动降级，这是一个非常实用的功能点。

#### 7. 用户反馈摘要

- **核心痛点**：从 Issue [#3012](https://github.com/sipeed/picoclaw/issues/3012) 的讨论中可以提炼出，用户对后台行为（如Evolution）导致的 **无感知资源消耗** 非常敏感。用户期望能够精确控制和预见 AI 服务的成本。
- **使用场景**：Issue [#3108](https://github.com/sipeed/picoclaw/issues/3108) 反映了用户在多模型切换（如使用文本模型处理图像需求）的复杂场景下，期望系统能有更智能的路由和错误处理逻辑，而不是产生迷惑性的错误输出。

#### 8. 待处理积压

- **功能请求**:**PR [#2964](https://github.com/sipeed/picoclaw/pull/2964) (Feat/image input compression)** 自 2026-05-28 提出后已开放超过两周，尚未收到项目维护者的审核或合并。鉴于其与用户成本控制的核心诉求相关，建议项目组安排评估。
- **技术优化**:**PR [#3120](https://github.com/sipeed/picoclaw/pull/3120) (feat(config): add RegisterChannelSettings hook)** 是提升项目架构扩展性的重要特性，于今日提出，应尽快进入审查流程。
- **关键 Bug**:**Issue [#3012](https://github.com/sipeed/picoclaw/issues/3012) (Token 持续消耗)** 目前无任何进展或指派，是当前最需要维护者介入的问题。

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw 项目动态日报 — 2026-06-14

---

## 今日速览

过去 24 小时项目活跃度较高：新增 1 个 Issue（#2751 关于 LLM 预算耗尽后响应静默丢弃的严重 bug），另有 1 个误报 Issue 已关闭。Pull Request 方面活跃度显著，共 10 条，其中 5 条已合并/关闭，5 条待审查。无新版本发布。合并的 PR 集中在 provider 能力框架、持久化内存支架、OneCLI SDK 升级及容器部署数据驱动化等底层基础设施改进，同时社区最关注的预算耗尽问题已有修复 PR（#2759）待合并。项目整体处于功能拓展与稳定性加固并行推进的状态。

---

## 版本发布

无新版本发布。

---

## 项目进展

今日合并/关闭的 PR 共 5 条，标志着多个关键基础设施模块的落地：

- **#2758** [CLOSED] — `feat(container): data-drive global CLI installs from cli-tools.json`  
  将容器内全局 Node CLI（如 `claude-code`、`agent-browser`、`vercel`）的安装方式从硬编码 Dockerfile 改为基于数据清单 `container/cli-tools.json`，方便技能模块以 JSON merge 方式声明依赖。这简化了容器构建流程，提升了可维护性。  
  [PR #2758](https://github.com/nanocoai/nanoclaw/pull/2758)

- **#2754** [CLOSED] — `feat(runner): onExchangeComplete provider hook + slash-command interruption`  
  为 agent-runner 引入了可选的 `onExchangeComplete` provider 钩子以及斜杠命令中断机制，使 provider 能在每次 LLM 回合完成后执行自定义逻辑，并为用户提供更灵活的交互中断方式。  
  [PR #2754](https://github.com/nanocoai/nanoclaw/pull/2754)

- **#2747** [CLOSED] — `feat(onecli): SDK 2.2.1 — credential-stub mounts + machine-checkable pins`  
  将 `@onecli-sh/sdk` 从 0.5.0 升级至 2.2.1，引入凭证存根挂载和可机器验证的 pin 机制，强化了 OneCLI 网关的安全性与兼容性。  
  [PR #2747](https://github.com/nanocoai/nanoclaw/pull/2747)

- **#2746** [CLOSED] — `feat(providers): agent-surfaces capability seam`  
  在 host 侧新增 provider 注册表，允许 provider 声明自身的能力范围（capability seam），为后续多 provider 动态选择奠定基础。  
  [PR #2746](https://github.com/nanocoai/nanoclaw/pull/2746)

- **#2745** [CLOSED] — `feat(memory): opt-in persistent memory scaffold for providers`  
  为 provider 添加可选的持久化内存支架（`usesMemoryScaffold` 能力标志及容器级支持），使 provider 能够跨会话保留记忆，属于长期记忆功能的基础组件。  
  [PR #2745](https://github.com/nanocoai/nanoclaw/pull/2745)

这些合并推进了项目在多 provider 支持、容器部署标准化、用户交互钩子、安全凭证管理以及会话记忆持久化等方向的重要进展。

---

## 社区热点

今日仅有 2 条 Issue 更新，其中 **#2751** 是最受关注的实体问题，尽管评论数为 0，但其描述的**预算/Token 耗尽后 LLM 响应被静默丢弃**的问题直接触及用户核心体验。该 Issue 由 `assapin` 提交，并已附带修复 PR #2759，成为社区当前最高优先级的讨论热点。  
[Issue #2751](https://github.com/nanocoai/nanoclaw/issues/2751)

另一条待合并的 PR **#2756** （`operator-driven provider selection`）与 **#2757**（`Codex provider payload v2`）则代表了社区对多 provider 切换与扩展的长期需求，虽无大量评论，但体现了项目路线图中的关键方向。

---

## Bug 与稳定性

| 严重程度 | Issue/PR | 描述 | 状态 |
|----------|----------|------|------|
| **严重** | [#2751](https://github.com/nanocoai/nanoclaw/issues/2751) | LLM 预算/Token 耗尽后 agent-runner 静默丢弃响应，用户得不到任何反馈。 | 已关联修复 PR [#2759](https://github.com/nanocoai/nanoclaw/pull/2759)（待合并） |
| **高** | [#2750](https://github.com/nanocoai/nanoclaw/pull/2750) | 容器被 `SIGKILL` 后 `outbound.db` 的 WAL 日志残留导致数据不一致；同时修复了热轮询竞争条件（关联 #2516、#2640）。 | PR 开放中，已获得正确诊断与修复方案 |
| **中** | [#2732](https://github.com/nanocoai/nanoclaw/pull/2732) | 多智能体健康审计发现的系列问题：bind-mount 真实路径解析（解决 Docker Desktop drvfs 导致 exit 127 崩溃）、容器启动断路器、最大并发数限制、daemon 级 `docker kill` 回退、agent-runner 超时与重试。 | PR 开放中，已获多轮 adversarial 验证 |

其中 **#2751** 是用户直接报告的线上故障，应优先合并对应修复。

---

## 功能请求与路线图信号

- **多 Provider 动态选择与切换**：PR #2756 将 provider 选择权交予 operator，并配套 provider 注册表、安装向导、vault 认证及内存迁移技能，是项目走向多模型、多平台运营的核心功能。结合已合并的 #2746（capability seam），该方向正逐步成型。  
  [PR #2756](https://github.com/nanocoai/nanoclaw/pull/2756)

- **Codex 作为完整 provider 的 v2 版本**：PR #2757 将 Codex 改造为基于 host 能力 seam 的完整 provider，仅通过 vault 认证，替代旧版 payload。这暗示项目可能重点支持 Codex 作为除默认 provider 之外的首选替代方案。  
  [PR #2757](https://github.com/nanocoai/nanoclaw/pull/2757)

- **容器部署数据驱动化**：PR #2758 的合并表明项目正引入声明式工具依赖管理，未来技能模块可通过简单的 JSON 条目声明所需 CLI 及版本，无需修改 Dockerfile，降低了社区贡献容器技能的门槛。

---

## 用户反馈摘要

- **#2751** 用户 `assapin` 报告了核心痛点：当 LLM 调用因 Token/预算限制而失败时，agent-runner 完全静默处理，用户既看不到错误消息也得不到替代响应。此问题直接影响可观测性与用户体验，且用户已自行提供修复方案（#2759），表明社区用户有深度参与调试和贡献的意愿。

其他 Issues/PR 暂无直接的用户评论反馈。

---

## 待处理积压

以下为当前需要维护者重点关注的开放 PR 和 Issue：

| 编号 | 类型 | 标题 | 创建时间 | 持续周期 | 备注 |
|------|------|------|----------|----------|------|
| [#2751](https://github.com/nanocoai/nanoclaw/issues/2751) | Issue | Budget-exhausted LLM turns are silently dropped | 2026-06-12 | 2 天 | 严重 bug，已有修复 PR，建议尽快合并 |
| [#2759](https://github.com/nanocoai/nanoclaw/pull/2759) | PR | fix: deliver budget/billing error turns instead of dropping them | 2026-06-14 | 0 天 | 对应 #2751 的修复，需 review 与测试 |
| [#2750](https://github.com/nanocoai/nanoclaw/pull/2750) | PR | recover stale outbound.db journals after container kills | 2026-06-12 | 2 天 | 涉及持久化数据完整性，建议验证后合并 |
| [#2732](https://github.com/nanocoai/nanoclaw/pull/2732) | PR | Harden host + agent-runner from health audit findings | 2026-06-11 | 3 天 | 批量修复多项稳定性问题，需关注是否覆盖全面 |
| [#2756](https://github.com/nanocoai/nanoclaw/pull/2756) | PR | operator-driven provider selection | 2026-06-14 | 0 天 | 核心特性 PR，但配套的 Codex payload (#2757) 也可能需要同步评审 |
| [#2757](https://github.com/nanocoai/nanoclaw/pull/2757) | PR | Codex agent-provider payload v2 | 2026-06-14 | 0 天 | 依赖 #2756 的能力 seam，建议协同处理 |

> 提醒：以上 PR 均无 review comments 或 approval 标记，建议维护者分配 reviewer，优先处理 #2759（修复严重 bug）与 #2750（数据完整性修复），其次推进 #2756/#2757 特性分支的架构讨论。

---

*数据来源：GitHub [nanocoai/nanoclaw](https://github.com/nanocoai/nanoclaw) 2026-06-14 快照*

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw 项目动态日报 — 2026-06-14

## 1. 今日速览
- 过去24小时内项目活跃度中等，共产生1条新Issue和1条新PR，无新版本发布。
- 核心问题#941（Agent类型定时任务无法触发子进程）持续引起社区关注，已获7条评论，但尚未有官方修复PR。
- 新提交的PR #954针对#941提出了根本原因分析（use-after-free）并给出了修复方案，目前处于待合并状态。
- 整体来看，项目在消息投递稳定性方面存在一例已知严重缺陷，社区正在积极提交修复，但维护者尚未回应。

## 2. 版本发布
*（无新版本发布，本节省略）*

## 3. 项目进展
今日无任何PR被合并或关闭，项目主要进展体现在一条新提交的修复PR：
- **PR #954**（由 vernonstinebaker 提交）：修复一次性Agent类型Cron任务静默投递失败的问题。该PR指出根因是`OutboundMessage.channel`存在use-after-free，执行后消息对象被释放，导致无法通过Telegram、Mattermost等渠道发送。修复后应能使`schedule`中`job_type: "agent"`的定时任务正常投递消息。  
  🔗 [PR #954](https://github.com/nullclaw/nullclaw/pull/954)

目前项目尚未推进任何功能开发，处于bug修复集中期。

## 4. 社区热点
当前社区讨论最集中的是**Issue #941**，共7条评论，由用户 weissfl 在2026-05-31提出，更新于2026-06-13。  
- **核心诉求**：配置了`job_type: "agent"`、`delivery_mode: "always"`、`delivery_channel: "telegram"`的定时任务，状态显示已完成，但Agent子进程从未启动，Telegram也收不到消息。用户期望能正常触发Agent并投递。  
- **社区反响**：评论中多名用户确认复现，并指出该问题影响所有“一次性”Agent定时任务（包括Mattermost等渠道）。  
- **关联PR**：PR #954 已被社区用户标记为#941的解决方案，但维护者未回复。  
🔗 [Issue #941](https://github.com/nullclaw/nullclaw/issues/941)

## 5. Bug 与稳定性
今日报告的Bug主要围绕消息投递失败，按严重程度排列如下：

| 严重程度 | 问题描述 | 对应Issue/PR | 是否有修复PR |
|----------|----------|--------------|--------------|
| **严重** | Agent类型一次性Cron任务静默失败，子进程不启动，消息不投递 | #941 | ✅ 有（#954，待合并） |
| **严重** | 根因：`OutboundMessage.channel`在消息发送后被释放，导致use-after-free | PR #954 （分析部分） | ✅ 同一PR已包含修复 |

暂无其他崩溃、回归或安全问题报告。

## 6. 功能请求与路线图信号
今日未收到新功能请求。当前社区焦点完全集中在修复上述Bug上，未出现新功能提案。若PR #954被合并，可间接恢复“Agent定时任务”功能的可用性，这可能是下一版本（若有）的必备修复。

## 7. 用户反馈摘要
从Issue #941的评论及描述中可提炼以下用户痛点：
- **核心痛点**：配置了完整参数（`delivery_mode: "always"`、`delivery_channel: "telegram"`）的定时任务“看起来”执行了（标记完成），但实际Agent未启动，用户无法通过任何日志或界面获知失败。
- **使用场景**：用户希望利用`schedule`与Agent实现自动化消息推送（如定时报告、告警通知），但该Bug导致此场景完全不可用。
- **满意度**：用户对“静默失败”的体验表示失望，认为这类关键链路错误应至少记录错误日志或弹出警告，而非默默标记完成。

## 8. 待处理积压
- **Issue #941**（创建于2026-05-31，13天未获维护者回应）：严重Bug，已有社区提供的修复PR，但维护者尚未review或merge。建议维护者尽快确认修复方案并合并，避免社区信任度下降。  
  🔗 [Issue #941](https://github.com/nullclaw/nullclaw/issues/941)
- **PR #954**（创建于2026-06-13，今日提交）：修复代码，目前无冲突，分支名称明确针对#941。若合并可立即解决该积压问题。  
  🔗 [PR #954](https://github.com/nullclaw/nullclaw/pull/954)

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

好的，各位关注 IronClaw 项目的同仁，以下是 2026-06-14 的项目动态日报。

---

## IronClaw 项目日报 — 2026-06-14

### 1. 今日速览

今日项目活跃度极高，核心开发者在 **Reborn 平台稳定性** 和 **Slack 交付** 两个关键领域进行了一场“集中修复冲刺”。共有 11 个 Issues 被开启，21 个 PR 被更新，其中 10 个 PR 已成功合并。特别值得注意的是，一个针对 Slack 审批循环的复杂修复 PR (#4839) 已合并，尽管留下了一些需要后续解决的跟踪问题。社区用户提交的 Issues 反映了对审批流程透明度和平台稳定性的更高期望。项目整体处于向 Reborn 平台过渡的关键时期。

### 2. 项目进展

今日合并/关闭了多个高价值 PR，项目在 **Reborn 平台稳定性**、**Slack 集成修复** 和 **附件功能集成** 三个核心线上取得了实质性进展。

-   **Reborn 稳定性与 Slack 交付：**
    -   **`#4839` [已合并]**：修复了 Reborn 中 Slack 渠道的重新审批循环问题。该问题导致一个需要同时进行审批和身份验证的控制调用，每次恢复时都会要求新的审批。修复通过在重新调度时保留调用身份来解决。 **重要提示**：此 PR 的合并也产生了多个用于追踪后续分解工作的 Issues（如 `#4849`, `#4850`）。
        [链接](https://github.com/nearai/ironclaw/pull/4839)
    -   **`#4843` [已合并]**：对 Slack 交付层进行“单路飞行”优化，确保同一个 `run_id` 的门控解决通知不会被多次提交，解决了另一个审批重复的 bug。
        [链接](https://github.com/nearai/ironclaw/pull/4843)
    -   **`#4844` [已合并]**：修复了 Slack 消息路由中，由于 `GateRef` 过滤函数使用不当，导致审批和认证消息可能发送到错误路由通道的问题。
        [链接](https://github.com/nearai/ironclaw/pull/4844)
    -   **`#4838` [已合并]**：**关键设计决策**。该 PR 用“显式拒绝繁忙线程”的模式替代了之前的“后台排队再重试”模式。当一条消息到达时，如果线程正被另一个运行中的任务占用，系统会直接向用户回复明确的通知，而不是偷偷排队，让用户自己决定何时重试。
        [链接](https://github.com/nearai/ironclaw/pull/4838)
    -   **`#4841` [已合并]**：旨在消除 Reborn 中导致运行彻底失败的错误。现在，诸如模型不可用等错误会向用户解释原因，并允许用户重试，而不是以不透明的代码终止运行。
        [链接](https://github.com/nearai/ironclaw/pull/4841)

-   **附件功能集成：**
    -   **`#4680`, `#4677`, `#4676`, `#4675`, `#4672`, `#4670` [均已合并]**：由核心贡献者 `ilblackdragon` 推动，这些 PR 共同构建并串联了 Reborn 平台的附件处理管道。从提取文档内容的独立 crate (`#4675`)，到将附件文本注入模型上下文的逻辑 (`#4677`)，再到 WebChat v2 的前端与后端上传集成 (`#4672`)，最后是 web 用户界面的体验打磨 (`#4738`)。这是一个功能完善的重要里程碑。
        [链接](https://github.com/nearai/ironclaw/pull/4680) | [链接](https://github.com/nearai/ironclaw/pull/4677) | [链接：](https://github.com/nearai/ironclaw/pull/4676) | ...（更多链接可在原始数据中找到）

### 3. 社区热点

今日没有出现评论数特别多的“爆款”议题，但 `#4850`, `#4849`, `#4848` 等一组由核心开发者 `henrypark133` 提交的 Issues 集中爆发，反映了 **PR #4839 合并后，社区和核心团队对代码质量和架构分解的高度关注**。这些 Issues 本身就是对 PR #4839 的 review 总结，意味着项目在快速迭代的同时，也在严格执行自身的架构规范（如文件大小限制、模块化）。这体现了项目追求长期健康度的工程文化。

### 4. Bug 与稳定性

今日报告了数个 Bug，主要集中在 Reborn 平台的用户体验和稳定性上。

-   **严重**:
    -   **`#4854` [审批提示过多]**: 用户报告执行一个简单的只读 GitHub 扩展请求时，遇到了多次不必要的审批提示。这反映了审批流程的粒度或逻辑存在问题，可能与 PR #4839 和 #4838 涉及的门控逻辑有关。
        [链接](https://github.com/nearai/ironclaw/issue/4854)
    -   **`#4853` [工具活动消失]**: 在 Railway 多租户环境中，工具执行时活动计数显示正常，但执行完成后活动记录会**消失**。这是一个严重的状态显示问题，可能导致用户困惑。
        [链接](https://github.com/nearai/ironclaw/issue/4853)

-   **中等**:
    -   **`#4852` [Shell命令不可见]**: 用户在审批弹窗和活动历史中只能看到 `Capability: builtin.shell`，但**看不到具体的 Shell 命令**。这严重影响了用户对工具调用的审查能力，是潜在的安全和信任隐患。
        [链接](https://github.com/nearai/ironclaw/issue/4852)

-   **长期存在的回归**:
    -   **`#4108` [夜间E2E测试失败]**: 自5月27日起的夜间端到端测试持续失败，涉及 `extensions` 等关键模块。这是一个需要优先解决的持续集成问题，影响新代码的质量保障。
        [链接](https://github.com/nearai/ironclaw/issue/4108)

-   **已有修复PR的Bug**:
    -   上述 `#4854`, `#4853` 等 Issue 目前尚未关联 fix PR。但 PR `#4841`（失败解释与重试）、`#4838`（拒绝而非排队）的合并，从架构上解决了导致 Bug 的一类根因。

### 5. 功能请求与路线图信号

今日的 Issues 和 PRs 清晰地指向了项目的两个核心演进方向：

-   **用户控制与透明性**: Issues `#4854` (审批过多)和 `#4852` (命令不可见) 表明用户对 Reborn 平台的交互控制权和决策透明度有更高要求。PR `#4838`（拒绝繁忙线程）和 `#4836`（向模型暴露运行时上下文）正是为了满足此类需求而设计的功能。

-   **多通道附件统一**: Issue `#4644` 是横跨多个 PR 和多次发布的史诗级任务。随着今日多批附件相关 PR 的合并，此功能已经接近尾声。WebChat v2 的增强和附件格式注册表（PR #4738）是下一步的重点。这很可能是一个即将发布的重大版本的核心功能。

### 6. 用户反馈摘要

-   **对审批透明度的强烈不满**：用户 `sunglow666` 在 Issue `#4854` 中描述其“简单只读请求”遭受了“多次不必要的审批提示”，并要求其审批“预先运行的工具调用”。这反映了用户对当前 Reborn 审批流程的感知是**冗余、低效且不信任**的。
-   **对运行状态显示的困惑**：用户 `sunglow666` 在 Issue `#4853` 中抱怨“工具活动执行后消失”，在 Issue `#4852` 中抱怨“看不到具体的命令”。这表明用户界面在提供**确定性反馈**方面存在明显不足，影响了用户对 Agent 行为的监控和信心。
-   **对构建/部署体验的隐忧**：问题 `#4108` 持续失败，虽然没有直接用户评论，但它象征着项目基础设施的健康度问题，最终会影响开发者和高级用户的自建和部署体验。

### 7. 待处理积压

-   **`#3708` [发版 PR]**: 这个发版 PR 已从 5 月 16 日开放至今，长达近一个月，且包含 `ironclaw_common` 和 `ironclaw_skills` 等核心 crate 的破坏性变更。长时间的阻塞可能意味着内部在协调版本发布的策略，或存在一些尚未解决的依赖冲突。
    [链接](https://github.com/nearai/ironclaw/pull/3708)
-   **`#4108` [E2E 测试失败]**: 这个持续 18 天的夜间 E2E 测试失败问题，尽管关闭了 0 条评论，但其重要性不言而喻。它会影响新代码合并的信心，建议维护组优先查明 root cause 并修复。
    [链接](https://github.com/nearai/ironclaw/issue/4108)

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是基于您提供的 GitHub 数据生成的 **LobsterAI 项目动态日报**。

---

# LobsterAI 项目动态日报 | 2026-06-14

## 1. 今日速览

过去24小时，LobsterAI 项目动态整体呈现 **低活跃度** 状态。核心表现为：没有新版本发布，且所有更新的 Issues 和 PR 均为创建于数月前的“陈年积压”，今日并无任何真正意义上的新提交或新讨论启动。尽管有 1 个重要的 Bug 修复 PR（#1465）被合并关闭，但项目活跃度评估为 **“静默”**，团队注意力可能已转向内部开发或其他模块。社区反馈集中在遗留的 UI 本地化与交互体验问题上，暂无新的激烈讨论。

## 2. 版本发布

无。

## 3. 项目进展

今日唯一重要的正向进展是 **1 个 Bug 修复 PR 已被合并关闭**，标志着项目在修复数据持久化可靠性上迈出了一步。

- **PR #1465 [已关闭]**: **修复定时任务的“幽灵会话”Bug**。
    - **作者**: linlihua
    - **概要**: 解决了用户删除定时任务后，重启应用该任务仍以空内容幽灵会话形式出现的顽固问题。根本原因是此前删除逻辑只清理了网关端（OpenClaw）的任务记录，未清理本地 SQLite 数据库中对应的会话记录。本次修复通过补充对 `cowork_sessions` 表中本地会话记录的清理逻辑，彻底解决了该数据不一致问题。
    - **链接**: [netease-youdao/LobsterAI PR #1465](https://github.com/netease-youdao/LobsterAI/pull/1465)

## 4. 社区热点

今日社区讨论活跃度极低，所有项目均只有 1 条评论且无反应（👍: 0）。但从 Issue 内容和 PR 标题来看，以下事项代表了用户长期关注的焦点：

- **热点 #1**: **UI 本地化不彻底** (Issue #1434)
    - **痛点**: 用户明确指出，当系统语言设为中文时，搜索无数据状态的提示文字和按钮仍显示为英文，影响产品专业度与用户体验。
    - **链接**: [netease-youdao/LobsterAI Issue #1434](https://github.com/netease-youdao/LobsterAI/issues/1434)

- **热点 #2**: **Agent 创建界面不友好** (Issue #1435)
    - **痛点**: 新建自定义 Agent 时，名称过长会直接超出弹框边界，属于明显的 UI 溢出 bug，反馈了细节打磨不足。
    - **链接**: [netease-youdao/LobsterAI Issue #1435](https://github.com/netease-youdao/LobsterAI/issues/1435)

- **热点 #3**: **Cowork 模块交互增强** (PR #1429, #1430, #1431)
    - **诉求**: 三个开放的 PR 都致力于增强 Cowork（协作）会话体验，包括会话内消息搜索、防止系统休眠、显示运行计时器。这些功能反映了用户对长时间运行任务时的进度感知和系统稳定性的迫切需求。
    - **链接**: [PR #1429](https://github.com/netease-youdao/LobsterAI/pull/1429), [PR #1430](https://github.com/netease-youdao/LobsterAI/pull/1430), [PR #1431](https://github.com/netease-youdao/LobsterAI/pull/1431)

## 5. Bug 与稳定性

今日无新报告的 Bug。以下是唯一被修复的、影响较大的历史遗留 Bug：

- **[严重]: 定时任务“幽灵会话”Bug (PR #1465) (已修复)**
    - **描述**: 用户删除定时任务后，重启应用该任务会以空内容的“幽灵”会话形式重新出现，反复删除无效。该问题严重影响定时任务功能的可用性与用户信任度。
    - **状态**: **已于今日合并关闭**，修复逻辑为补充本地数据库清理。
    - **关联 Issue**: #1359

- **【低】UI 本地化问题 (Issue #1434)**
    - 属于外观/本地化 Bug，非功能性问题。目前处于开放状态且标记为 `stale`（陈旧的），未有对应的修复 PR。

- **【低】UI 溢出问题 (Issue #1435)**
    - 属于前端展示 Bug，用户交互体验不佳。同样处于开放 `stale` 状态。

## 6. 功能请求与路线图信号

今日无新功能请求。但 **3 个长期开放且标记为 `stale` 的 PR (1429, 1430, 1431)** 集中指向了对 **“Cowork 会话增强”** 的强烈信号，这很可能是下一阶段版本的重要规划方向：

1.  **会话内消息实时搜索与高亮**：提升用户在长对话中检索信息的效率（PR #1429）。
2.  **防止系统休眠**：解决 Agent 执行耗时任务时因系统休眠导致中断的可靠性问题（PR #1430）。
3.  **会话运行计时器**：增强用户在等待任务完成时的进度感知与透明度（PR #1431）。

这些 PR 若被纳入后续版本，将显著提升 Cowork 模块的健壮性和用户体验。

## 7. 用户反馈摘要

从仅有的 Issues 评论中，可以提炼出用户对 LobsterAI 的具体意见：

- **用户在 Issue #1434 中反馈**：**对语言本地化的完成度不满意**。产品在多语言支持上存在明显遗漏（中文环境下仍有英文残留），表明项目在国际化和细节打磨上缺少最后一步的校验。
- **用户在 Issue #1435 中反馈**：**对新建 Agent 的交互表示困惑**。一个基础的输入框长度未做限制，产品设计对极端输入场景的容错性不足。

## 8. 待处理积压

以下项目创建于 4 月初，已超过 2 个月未得到实质性推进，均被标记为 `stale`，亟需维护者关注并决策：

1.  **[PR #1429] feat(cowork): 会话内消息搜索功能**
    - **创建**: 2026-04-03 | **状态**: OPEN & Stale
    - **问题**: 这是一个已具备完整实现代码的特性增强（带 `mark.js` 高亮），长期未被审查或合并，可能导致社区贡献者挫败感。
    - **链接**: [netease-youdao/LobsterAI PR #1429](https://github.com/netease-youdao/LobsterAI/pull/1429)

2.  **[PR #1430] feat(cowork): 阻止系统休眠功能**
    - **创建**: 2026-04-03 | **状态**: OPEN & Stale
    - **问题**: 这是针对长时间任务的可靠性修复，对于重度用户至关重要，积压时间过长。
    - **链接**: [netease-youdao/LobsterAI PR #1430](https://github.com/netease-youdao/LobsterAI/pull/1430)

3.  **[PR #1431] feat(cowork): 会话运行计时器**
    - **创建**: 2026-04-03 | **状态**: OPEN & Stale
    - **问题**: 同属 Cowork 模块直接提升用户感知的小而美的功能。
    - **链接**: [netease-youdao/LobsterAI PR #1431](https://github.com/netease-youdao/LobsterAI/pull/1431)

4.  **[Issue #1434] & [Issue #1435]**
    - **创建**: 2026-04-03 | **状态**: OPEN & Stale
    - **问题**: 两个 UI 层面的 bug，虽然优先级不高，但对用户体验有直接影响，应尽快确认并分配修复资源。
    - **链接**: [Issue #1434](https://github.com/netease-youdao/LobsterAI/issues/1434), [Issue #1435](https://github.com/netease-youdao/LobsterAI/issues/1435)

---
**总结**：LobsterAI 项目在过去24小时进入静默期。社区声量微弱，但积压的 PR 和 Issue 清晰地指出了软件在 **本地化细节**、**Agent 创建交互** 和 **Cowork 模块健壮性** 上的改进空间。维护团队应尽快审视这些积压项，以保持社区贡献的积极性并提升产品成熟度。

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域的开源项目分析师，我将根据您提供的 Moltis GitHub 数据，生成一份结构清晰、数据驱动的项目动态日报。

---

### Moltis 项目日报 | 2026年6月14日

---

### 1. 今日速览

今日项目活跃度 **中等偏高**。没有新版本发布，但社区的贡献和反馈非常积极。一个涉及 **MCP OAuth 认证失败** 的严重 Bug 在昨天被报告后，今天立即得到了修复 PR，体现了项目的快速响应能力。同时，基础设施（Docker 镜像）、核心依赖（esbuild）以及功能扩展（新内存后端）均有更新或提议，显示出项目在社区维护和性能优化方面的持续投入。

### 3. 项目进展

今日未有 PR 被合并或关闭，但**三个关键的待合并 PR** 反映了项目正在推进的重要方向：

- **核心功能修复：MCP OAuth 认证**（[PR #1120](https://github.com/moltis-org/moltis/pull/1120)）：作者 `xzavrel` 提交了针对 Issue #1119 的修复方案。该 PR 瞄准了因 `WWW-Authenticate` 头中包含 `resource_metadata` 参数而导致 OAuth 流程中断的严重问题，直接影响到 Notion、Linear 等主流 MCP 服务器的集成。此 PR 是项目修复核心集成能力的关键一步。
- **基础设施改进：Docker 构建优化**（[PR #1122](https://github.com/moltis-org/moltis/pull/1122)）：作者 `sayotte` 修复了一个影响部署稳定性的 Dockerfile 问题。通过移除与动态挂载用户目录冲突的 `VOLUME` 指令，解决了特定部署场景下的挂载覆盖问题，提升了容器化部署的健壮性和可预测性。
- **依赖更新：esbuild 升级**（[PR #1121](https://github.com/moltis-org/moltis/pull/1121)）：由 Dependabot 自动发起，将前端构建工具 esbuild 从 0.25.12 升级至 0.28.1。这通常包含性能改进、Bug 修复和安全更新，保持项目基础库的现代化。

### 4. 社区热点

今日讨论的焦点无疑是 **MCP OAuth 认证失败** 的问题。

- **[Issue #1119 - MCP OAuth fails with `invalid_target`](https://github.com/moltis-org/moltis/issues/1119)**：该 Issue 由用户 `xzavrel` 报告，详细描述了当 MCP 服务器 (如 Notion, Linear) 在 `WWW-Authenticate` 头中包含 `resource_metadata` 参数时，OAuth 授权流程会失败并返回 `invalid_target` 错误。此问题直击项目作为 “MCP 客户端” 的核心功能，阻碍了用户与这些流行服务的深度集成。
- **背后的诉求**：用户的核心需求是 **无缝、可靠地连接主流第三方 MCP 服务器**。问题暴露了 Moltis 在处理特定协议细节（`WWW-Authenticate` 扩展参数）上的兼容性不足。有趣的是，报告者与修复 PR 的作者为同一人，这是典型的“社区驱动开发”，用户不仅发现问题，还贡献了解决方案，大大加快了问题的解决速度。

### 5. Bug 与稳定性

今日报告的 Bug 数量少但**影响面广**。

| 严重程度 | Issue | 描述 | 是否已有修复 PR |
| :--- | :--- | :--- | :--- |
| **严重** | [#1119](https://github.com/moltis-org/moltis/issues/1119) | **MCP OAuth 集成完全失效**：当连接 Notion、Linear 等使用带 `resource_metadata` 的 OAuth 的服务器时，授权流程失败。 | **是** (PR #1120) |
| **中等** | [PR #1122](https://github.com/moltis-org/moltis/pull/1122) 关联 | **Docker 部署路径问题**：`VOLUME` 指令与用户主目录的 bind mount 冲突，可能导致数据或配置被意外覆盖。 | **是** (PR #1122 即为修复) |

### 6. 功能请求与路线图信号

今日出现了一个值得关注的功能请求，可能预示着项目未来的性能优化方向。

- **[Issue #1123 - Feature: 添加纯 Rust turbovec 内存后端](https://github.com/moltis-org/moltis/issues/1123)**：用户 `joeblew999` 提出为 Moltis 添加一个基于 `turbovec` 的纯 Rust 内存后端，以实现“极端边缘压缩”。结合当前 AI 领域对本地向量数据库性能的追求，此请求具有很强的合理性。

**路线图信号分析**：
- **性能优化是持续方向**：该请求指向了为 Moltis 的“记忆”或“上下文”系统寻找更高效、更紧凑的存储方案。这表明社区期望项目在资源受限的设备上也能有出色表现。
- **与现有 PR 的关联判断**：目前没有直接关联的 PR。但考虑到 `turbovec` 是专门为边缘和嵌入式场景优化的向量库，如果项目计划在 AI Agent 中采用更多本地化、低延迟的向量搜索，此功能有很高概率被纳入未来版本。

### 7. 用户反馈摘要

今日用户反馈主要集中在一个具体的技术痛点上。

- **用户痛点**：来自 Issue #1119 的贡献者 `xzavrel` 的反馈非常典型：“使用 Notion、Linear 的 MCP 服务器时，OAuth 流程总是失败”。这直接反映了用户对于 **“即插即用”的第三方服务连接体验** 的高期待，以及标准实现（OAuth）与服务器特有扩展（`resource_metadata`）之间兼容性冲突带来的困扰。
- **用户积极贡献**：值得注意的是，该用户不仅仅提交了 Bug 报告，还快速提交了修复 PR（#1120）。这表明 Moltis 的社区具有较高的技术能力和参与意愿，是项目健康发展的积极信号。

### 8. 待处理积压

目前没有明显的长期未响应 Issue 或 PR。今日的动态表明项目维护者反应迅速，对于严重 Bug 和基础设施改进都有及时跟进。所有待处理的 PR (#1120, #1121, #1122) 和 Issues (#1119, #1123) 都是近日产生，属于正常的工作积压。

**建议维护者优先处理**：
- **[PR #1120](https://github.com/moltis-org/moltis/pull/1120)**：作为严重 Bug 的修复，建议尽快 Code Review 并合并，恢复主流 MCP 服务器的连接功能。
- **[Issue #1123](https://github.com/moltis-org/moltis/issues/1123)**：建议维护者对该性能优化提议给予明确回应。即使是短暂的对齐需求、技术难度或路线图匹配度的讨论，也能有效激励社区贡献者。

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw 项目动态日报 — 2026-06-14

## 今日速览

过去 24 小时内，CoPaw 项目保持了较高的社区活跃度，共产生 9 条 Issues 和 12 条 Pull Requests。其中 1 条 Issue 被关闭，1 条 PR 被合并，其余均在讨论或审查中。多份新提交的 PR 来自首次贡献者（*first-time-contributor*），覆盖钉钉通道、UI 多语言、定时任务超时修复等多个方向，反映出社区参与度正在上升。整体来看，项目处于 **高活跃、小步快跑** 的健康状态，核心团队对用户反馈响应及时，但部分长期待审的 PR 仍需推动合并。

## 项目进展

### 合并/关闭的 PR

- **[#2498] fix(agents): use console language when creating agent and fallback unsupported langs**  
  作者: Alneys | 合并于 2026-06-13  
  该 PR 解决了新建 Agent 时始终默认使用英语（en）及中文人设文件的问题，改为读取用户 UI 语言设置，并添加了服务端语言回退逻辑。提升多语言用户的配置体验，对国际化部署有积极影响。

### 关键进展说明

- 社区贡献者 **nguyenthanhthe** 连开 4 个新 PR（#5175、#5176、#5178、#5179、#5180），集中解决 Console 前端 UI 显示、会话过滤、多 Agent 协作触发词等细节问题，表明团队正对前端体验进行系统性打磨。
- **ly-wang19** 提交的 5 个质量修复 PR（#5035、#5037、#5038、#5040、#5041）仍在 Under Review 状态，内容涉及 llama.cpp 版本解析、Linux 浏览器检测、空消息列表保护、备份容错等底层稳定性提升，若合并将显著降低边缘崩溃率。

## 社区热点

### 最活跃 Issue

- **[#5156] [Feature]: 建议支持 kimi-for-coding / 加入 uv 白名单**  
  作者: wjt0321 | 评论: 5 | [链接](https://github.com/agentscope-ai/QwenPaw/issues/5156)  
  该 Issue 持续获得讨论，用户明确表达了已订阅 Kimi coding 套餐却无法在 CoPaw 中接入的痛点。核心诉求是支持 `kimi-for-coding` 端点并加入 `uv` 白名单，以便利用已有的付费服务。该需求若被采纳，将直接拓宽模型接入生态，值得团队评估优先级。

- **[#5172] [Bug]: 聊天总出现问完问题没反应一直等待**（已关闭）  
  作者: kfrtiamo | 评论: 1 | [链接](https://github.com/agentscope-ai/QwenPaw/issues/5172)  
  虽已关闭，但该问题反映了用户在高频使用时遇到的严重卡顿现象（Task has been cancelled!），并明确指出在 QQ/微信等渠道中无法手动停止的危害。该 Bug 虽已标记为关闭，但根本原因是否已根除需关注。

## Bug 与稳定性

按严重程度排列如下：

1. **[#5181] [Bug]: 插件依赖安装导致 cmd 窗口持续弹窗** (OPEN)  
   - 严重度：高 (影响用户体验，可能泄露隐私)  
   - 描述：升级至 v1.1.11.post2 后，pip 安装插件依赖时未隐藏 cmd 窗口，且网络不稳定时反复重试，桌面频繁闪现黑窗。  
   - Fix PR: 暂无  
   - [链接](https://github.com/agentscope-ai/QwenPaw/issues/5181)

2. **[#5171] [Bug]: 上下文压缩保留缺少按条数保留，导致信息完全丢失** (OPEN)  
   - 严重度：高 (导致 Agent 上下文丢失、任务中断)  
   - 描述：当人设文件 token 大于保留阈值时，压缩后上下文变为 0，模型无法继续任务。  
   - Fix PR: 暂无  
   - [链接](https://github.com/agentscope-ai/QwenPaw/issues/5171)

3. **[#5177] [Bug]: DingTalk channel 消息未注册到 chats.json** (OPEN)  
   - 严重度：中 (前端 Console 无法显示钉钉会话)  
   - 描述：钉钉消息可正常回复，但未写入 `chats.json`，导致前端看不见会话列表。  
   - 关联 PR: 暂无专门修复，但可能是通道注册逻辑问题。  
   - [链接](https://github.com/agentscope-ai/QwenPaw/issues/5177)

4. **[#5174] [Bug]: 定时任务和心跳机制缺陷** (OPEN)  
   - 严重度：中 (功能部分失效)  
   - 描述：Cron agent 无法产出知识文件（不会 write_file / spawn_subagent），心跳 agent 不按 HEARTBEAT.md 执行重任务。  
   - 关联 PR: [#5180](https://github.com/agentscope-ai/QwenPaw/pull/5180) 尝试增加超时时间和上下文提示，但尚未合并。  
   - [链接](https://github.com/agentscope-ai/QwenPaw/issues/5174)

5. **[#5172] [Bug]: 聊天等待问题** (已关闭)  
   - 虽关闭，潜在根本原因（如 session 超时或消息队列阻塞）仍可能影响其他用户，建议关注复现环境。

## 功能请求与路线图信号

| 功能需求                                                  | 来源 Issue | 对应 PR 或路线图判断                                                                 |
| --------------------------------------------------------- | ---------- | ------------------------------------------------------------------------------------ |
| 支持 `kimi-for-coding` / 加入 `uv` 白名单                 | #5156      | 无 PR，但讨论热度高，可能在下个版本纳入模型接入层扩展。                              |
| 添加越南语 (vi) 界面语言                                  | #5169      | 同日已有对应 PR [#5175](https://github.com/agentscope-ai/QwenPaw/pull/5175)，几乎确定合并。 |
| 优化模型配置功能，统一向量/文本/音视频模型配置           | #5182      | 无 PR，属于架构改进，但作者 hongweifei 为社区核心贡献者，可能进入 Sprint 规划。       |
| Console 会话列表增加标题过滤                              | 隐含需求   | PR [#5178](https://github.com/agentscope-ai/QwenPaw/pull/5178) 已提交，方向明确。     |
| 增强多 Agent 协作技能触发词（如“团队协作”模式跳转）       | 隐含需求   | PR [#5179](https://github.com/agentscope-ai/QwenPaw/pull/5179) 已提交，增加关键词匹配。 |

## 用户反馈摘要

- **付费用户痛点**：多位用户（#5156、#5172）提及已订阅第三方服务（Kimi coding）或使用 QQ/微信接入，但在当前版本中遇到功能缺失或卡死问题，反馈情绪急切。表明 CoPaw 的通道和模型集成层需要更灵活的支持，以降低用户迁移成本。
- **稳定性抱怨**：Issue #5171 的用户直言“上下文完全丢失导致任务中断”，且强调该问题在 v1.1.11 版本中存在，说明长期运行的 Agent 需要更鲁棒的压缩策略。
- **国际化诉求**：越南语用户（#5169）感谢团队对印尼语/巴西葡语的支持，并主动提供翻译，体现了社区对多语言功能的正向反馈。
- **贡献者态度**：首次贡献者 nguyenthanhthe 和 ly-wang19 集中提交多个 PR，说明项目文档和入门流程对新手友好，但部分 PR 等待评审超过 5 天（#5035 #5037 等），需注意避免挫伤贡献热情。

## 待处理积压

### 长期未响应的重要 Issue

| Issue                                                       | 创建时间   | 状态     | 备注                                                              |
| ----------------------------------------------------------- | ---------- | -------- | ----------------------------------------------------------------- |
| [#5035] fix(local_models): parse llama.cpp server version   | 2026-06-09 | Under Review | 已等待 5 天，涉及构建号位数增长导致的解析崩溃，属稳定性隐患。     |
| [#5040] fix(crons): tolerate invalid jobs in jobs.json      | 2026-06-09 | Under Review | 关键容错修复，单个任务损坏即可导致全部 cron 功能瘫痪。            |
| [#5041] fix(backup): skip unreadable files                  | 2026-06-09 | Under Review | Windows 环境下权限文件导致整个备份失败，影响数据安全。            |
| #5156 支持 kimi-for-coding                                  | 2026-06-12 | OPEN      | 用户呼声高，但尚未有维护者响应或指派，建议由产品经理评估优先级。  |

### 建议行动

1. 优先审查 ly-wang19 的 5 个 PR，尤其是 #5035、#5040、#5041 三个累计待审已超一周的修复。
2. 对 #5181（cmd 弹窗）给出热修复或临时工作区建议，避免高频重复弹窗影响用户信任。
3. 在项目 Roadmap 中将“多模型端点（kimi-for-coding）”标记为待调研，回应社区诉求。

---

**报告生成时间**：2026-06-14 23:59 UTC  
**数据来源**：CoPaw GitHub 仓库（agentscope-ai/QwenPaw）

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

# ZeroClaw 项目动态日报 — 2026-06-14

## 今日速览

- 项目今日保持 **高度活跃**，共处理 50 个 PR（其中 11 个已合并/关闭，39 个待合并）和 4 个 Issue（3 个新开，1 个关闭）。  
- 社区焦点集中在 **quickstart 流程改进**（别名校验、Webhook 端口）、**gateway 稳定性**（`ask_user` 工具在 WebSocket 下的阻塞故障）以及 **代码质量提升**（消除裸 `unwrap`、禁用 `println!`）。  
- 一项重要的 **RFC 合并**（#7415）将三个 agent 运行循环统一为一个 consolidation PR，标志着架构级优化进入落地阶段。  
- 无新版本发布，但多项修复和增强已合入主干，预计近期有版本计划。

---

## 版本发布

无新版本发布。

---

## 项目进展 — 今日合并/关闭的重要 PR

以下 PR 已合入 `master`，显著提升了代码健壮性、用户体验和可维护性：

| PR | 描述 | 影响 |
|----|------|------|
| [#7417](https://github.com/zeroclaw-labs/zeroclaw/pull/7417) | 修复 cron 编辑模态框无法匹配新增表单所有字段的问题，后端扩展 `CronPatchBody` 字段，前端补全 UI | 消除 cron 任务编辑时的字段遗漏，提升管理体验 |
| [#7607](https://github.com/zeroclaw-labs/zeroclaw/pull/7607) | 将 gateway 中 `HeaderValue` 的裸 `unwrap()` 替换为带说明的 `expect()` | 提高 panic 时的诊断信息，降低隐含风险 |
| [#7606](https://github.com/zeroclaw-labs/zeroclaw/pull/7606) | 将 `EventBuffer` 及工具查找中的 `Mutex::lock().unwrap()` 替换为 `.expect("...not poisoned")` | 明确 panic 原因，防止静默误用 |
| [#7605](https://github.com/zeroclaw-labs/zeroclaw/pull/7605) | 将 cron store 中的 `println!` 替换为 `zeroclaw_log::record!` | 符合项目 clippy 规范，统一日志输出 |
| [#7604](https://github.com/zeroclaw-labs/zeroclaw/pull/7604) | 修复计算器工具中 `partial_cmp` 在 NaN 时引发 panic 的问题，改用 `total_cmp` | 提升工具数学计算的健壮性 |
| [#7599](https://github.com/zeroclaw-labs/zeroclaw/pull/7599) | 将 daemon 启动 banner 中的 6 个 `println!` 替换为结构化日志 | 标准化守护进程日志，便于监控 |

此外，**[#7415 Issue](https://github.com/zeroclaw-labs/zeroclaw/issues/7415)（RFC: Unify the three agent turn engines）已于今日关闭**，其对应的 consolidation PR（#7540）已执行，将 `run_tool_call_loop`、`turn_streamed` 和 `Agent::turn` 统一为一个引擎。这标志着 agent 执行层架构的重大演进。

---

## 社区热点

### 1. 统一 turn 引擎 RFC（#7415）—— 架构级讨论
- **链接**：[Issue #7415](https://github.com/zeroclaw-labs/zeroclaw/issues/7415)
- **背景**：该 RFC 最初提出分阶段迁移，但在 maintainer 指导下直接以单个 consolidation PR 执行。5 条评论集中在实现细节和 test coverage 上。
- **诉求**：社区希望减少 agent 执行路径的重复逻辑，降低维护成本，并确保向后兼容。

### 2. Zerocode ACP Bridge（#6823）—— 客户端连接层
- **链接**：[Issue #6823](https://github.com/zeroclaw-labs/zeroclaw/issues/6823)
- **背景**：自 5 月 21 日创建以来持续活跃，今日仍有人关注（评论 2 条）。该 tracker 负责 TUI 到 daemon 的 JSON-RPC 连接层，是远程管理的核心组件。
- **诉求**：社区期待该功能尽快稳定，以便在 TUI 端实现完整的 daemon 控制。

### 3. 今日大量新开 PR 集中在 `ask_user` 修复
- **PR #7618**、**#7615**、**#7613** 均由 `xuwei-xy` 提交，针对 issue #7542 和 #7551，修复 gateway WebSocket 下 `ask_user` 工具“通道关闭”的无提示失败。这些 PR 获得快速关注，因为它们影响 S1 级工作流阻塞。

---

## Bug 与稳定性

| 严重度 | 问题 | 状态 | 修复 PR |
|--------|------|------|---------|
| **S1 - 工作流阻塞** | `ask_user` 在 gateway WebSocket 会话中立即失败，返回“Channel closed before…”误导性消息（#7542） | 已修复 | [#7618](https://github.com/zeroclaw-labs/zeroclaw/pull/7618) / [#7613](https://github.com/zeroclaw-labs/zeroclaw/pull/7613) |
| **S1 - 工作流阻塞** | WebSocket 审批通道 `supports_free_form_ask()` 未覆盖，导致 `ask_user` 返回错误而非清晰提示（#7551） | 已修复 | [#7615](https://github.com/zeroclaw-labs/zeroclaw/pull/7615) |
| **S2 - 行为退化** | `zeroclaw quickstart` 中若输入大写字母的 agent alias，不会在输入时校验，导致后续创建失败且丢失所有已填信息（#7591） | 已提交修复 | [#7609](https://github.com/zeroclaw-labs/zeroclaw/pull/7609) |
| **S2 - 行为退化** | MCP 工具启用后，若 agent `risk_profile` 设置了显式 `allowed_tools`，枚举到的工具不会自动加入（#7547衍生） | 待合入 | [#7547](https://github.com/zeroclaw-labs/zeroclaw/pull/7547) |
| **S3 - 代码质量** | 多个仓库文件存在裸 `unwrap()`（gateway、cron、calculator 等） | 已修复 | 多条 PR 合并 |

---

## 功能请求与路线图信号

### 可能纳入下一版本的新功能
1. **Telegram Markdown 支持**（[#7611](https://github.com/zeroclaw-labs/zeroclaw/issues/7611)）  
   - 用户请求在 Telegram 频道中使用富文本格式化，Telegram Bot API 已支持。该请求无 PR 但门槛低，大概率被快速实现。

2. **quickstart 中增加 Webhook 端口提示**（[#7610](https://github.com/zeroclaw-labs/zeroclaw/pull/7610)）  
   - 已提交 PR，修复了用户选择 webhook 渠道后未询问端口的问题，有望随下个小版本发布。

3. **通道压缩前通知**（[#7162](https://github.com/zeroclaw-labs/zeroclaw/pull/7162)）  
   - 6 月 3 日提交的增强，让用户在上下文压缩前收到提示，已获得 maintainer 初步认可，仍待 review。

4. **Slack 附件上传**（[#7170](https://github.com/zeroclaw-labs/zeroclaw/pull/7170)）  
   - 同为 6 月 3 日提交，支持向 Slack 上传图片/文件附件，对于文档生成类 agent 非常实用。

### 长期路线图信号
- **Zerocode ACP Bridge**（[#6823](https://github.com/zeroclaw-labs/zeroclaw/issues/6823)）作为 tracker，意味着项目正在构建独立的 TUI 连接层，最终可能成为跨机器管理 daemon 的标准接口。

---

## 用户反馈摘要

从今日活跃的 Issues 评论和 PR 讨论中提炼：

1. **“别名校验缺失”严重影响首次使用体验**  
   - 用户 `ax-fe` 在 [#7591](https://github.com/zeroclaw-labs/zeroclaw/issues/7591) 中抱怨：`zeroclaw quickstart` 中输入大写 alias 时无反馈，直到最后 “apply” 阶段才报错，且所有已填信息丢失，必须从头开始。这种 “功亏一篑” 的流程设计对新手非常不友好。

2. **“ask_user 在 Web 端工作流中彻底阻断”**  
   - 多个用户（尤其是通过 gateway 使用 WebSocket 的团队）报告 `ask_user` 工具无法使用，返回的错误信息无助于诊断。修复 PR 获得迅速响应，说明该问题影响面广。

3. **“希望 Telegram 能跟上官方 Bot API 的最新特性”**  
   - 用户 `kshcherban` 提出 Telegram 已支持 Markdown 格式，ZeroClaw 应默认启用，以提升消息可读性。

4. **“配置文件嵌套太容易写错”**  
   - PR [#7617](https://github.com/zeroclaw-labs/zeroclaw/pull/7617) 针对额外嵌套的 provider alias（如多写一层 `default`）增加了警告，侧面反映用户常因 TOML 层级错误导致配置静默失效。

---

## 待处理积压

以下为长期开放但近期有活动或较高的重要 Issue/PR，提醒维护者关注：

| 编号 | 类型 | 描述 | 创建时间 | 最后更新 | 备注 |
|------|------|------|----------|----------|------|
| [#6823](https://github.com/zeroclaw-labs/zeroclaw/issues/6823) | Issue | Zerocode ACP Bridge tracker | 2026-05-21 | 2026-06-14 | 核心功能 tracker，需定期同步进展 |
| [#7162](https://github.com/zeroclaw-labs/zeroclaw/pull/7162) | PR | 通道压缩前通知 | 2026-06-03 | 2026-06-14 | 已超 11 天未 review |
| [#7170](https://github.com/zeroclaw-labs/zeroclaw/pull/7170) | PR | Slack 附件上传 | 2026-06-03 | 2026-06-14 | 同上，需 maintainer 评估 |
| [#7450+](https://github.com/zeroclaw-labs/zeroclaw/issues?q=is%3Aopen+sort%3Aupdated-asc) | - | 其他早期未标记的开放问题 | - | - | 建议使用 [no-stale](https://github.com/zeroclaw-labs/zeroclaw/issues?q=is%3Aopen+label%3Ano-stale) 标签筛选长期关注事项 |

---

**总结**：ZeroClaw 今日在 **用户体验修复**（quickstart 校验、WebSocket 工具阻塞）和 **代码规范统一**（去除裸 unwrap、日志标准化）方面取得实质性进展。RFC 统一 turn 引擎的落地标志着核心架构收敛，多项长期增强 PR 仍待 review，希望 maintainer 能尽快纳入下一版本。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*