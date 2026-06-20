# OpenClaw 生态日报 2026-06-20

> Issues: 116 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-06-20 10:17 UTC

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

好的，作为一名 AI 智能体与个人 AI 助手领域开源项目分析师，我将根据您提供的 OpenClaw (github.com/openclaw/openclaw) GitHub 数据，为您呈现 2026 年 6 月 20 日的项目动态日报。

---

# OpenClaw 项目日报 | 2026-06-20

## 1. 今日速览

今日项目活跃度极高，共产生 **116 条 Issue 更新**和 **500 条 PR 更新**，显示出社区极高的参与度和强烈的反馈意愿。然而，**Issue 关闭率极低**（仅 3 条），**PR 合并率也较低**（仅 19 条），表明项目维护容量已成为**当前最显著的瓶颈**。大量高优先级 Bug 和待审核的 PR 积压严重，尤其是 `v2026.5.28` 版本引入的多个回归问题仍在等待修复，项目健康度受此影响较大。虽然社区讨论热烈，但**项目的前进速度明显落后于问题发现的速度**。

## 2. 版本发布

**无新版本发布。**

## 3. 项目进展

今日合并/关闭的重大事件相对较少，但少数关键 PR 的推进值得关注：

- **功能增强：**
    - **微信（WeChat）支持**：PR [#85866](https://github.com/openclaw/openclaw/pull/85866) 为 WhatsApp 添加了电话号码登录功能，展示了跨渠道认证能力的通用化。
    - **模型生态扩展**：PR [#92217](https://github.com/openclaw/openclaw/pull/92217) 为 Fireworks 服务目录新增了 DeepSeek V4 Pro 等多个推理模型。
- **Bug 修复尝试：**
    - **关键消息丢失问题**：PR [#89039](https://github.com/openclaw/openclaw/pull/89039) 修复了一个因 `EmbeddedAttemptSessionTakeoverError` 导致静默消息丢失的严重问题，目前处于重新审查阶段。
    - **TTS 问题**：PR [#83988](https://github.com/openclaw/openclaw/pull/83988) 解决了 Telegram 中 TTS “最终模式”下的文本闪烁问题，已经准备好等待维护者审查。
- **维护与文档：**
    - **文档完善**：PR [#89142](https://github.com/openclaw/openclaw/pull/89142) 试图通过生成 SecretRef 的参考文档，提升配置安全性。
    - **问题诊断**：PR [#89226](https://github.com/openclaw/openclaw/pull/89226) 修复了 `openclaw doctor` 命令在特定环境下崩溃的问题。

**总体来看，项目在处理小规模修复和模型扩展上有所进展，但核心稳定性问题和大量待合并 PR 的积压是当前最主要的风险。**

## 4. 社区热点

今日讨论最热门的议题集中在 **`v2026.5.28` 版本的回归问题**和**长期存在的架构性缺陷**上。

1.  **GHCR 镜像配置问题 ([#88788](https://github.com/openclaw/openclaw/issues/88788))**：关于 `2026.5.28` 版本 Docker 镜像的配置架构与运行时文件不一致的问题，获得了 12 条评论。这反映了**版本发布质量与部署一致性**的社区担忧。
2.  **OAuth 刷新失败导致Agent僵死 ([#86215](https://github.com/openclaw/openclaw/issues/86215))**：Codex OAuth 刷新失败导致 Agent 数小时无响应的问题持续受到关注（9 条评论）。用户要求更激进的配置轮换，其背后是对**Agent高可用性**的迫切需求。
3.  **Gateway内存泄漏 ([#89315](https://github.com/openclaw/openclaw/issues/89315))**：Gateway 堆内存无限制增长导致被 OOM 杀死的问题，获得 6 条评论和 3 个 👍。这被标记为 P1 且评级为“白金隐士”，表明**长期运行部署的稳定性**是社区升级的关键障碍。
4.  **子Agent超时无通知 ([#88856](https://github.com/openclaw/openclaw/issues/88856))**：`sessions_spawn` 工具调用后，子 Agent 静默丢失的问题回溯到约 3.8% 的历史发生率。这揭示了**多Agent编排中的可靠性和可观测性**存在严重缺陷。
5.  **Bedrock回归 ([#88707](https://github.com/openclaw/openclaw/issues/88707))**：`2026.5.27` 到 `2026.5.28` 的升级导致 Bedrock 用户完全无法使用，直接阻碍了特定用户群体的升级路径，社区反响强烈。

## 5. Bug 与稳定性

今日报告的 Bug 数量多、影响面广，按严重程度排列如下：

- **严重回归/崩溃 (P1 / 评级🐚 白金隐士)：**
    - **Gateway 内存泄漏**：`[Bug]: gateway heap grows unbounded over time, gets killed by cgroup OOM...` ([#89315](https://github.com/openclaw/openclaw/issues/89315))。无 fix PR。
    - **备份功能损坏**：`openclaw backup create --verify exits 13, leaves corrupt .tmp archive...` ([#89257](https://github.com/openclaw/openclaw/issues/89257))。无 fix PR。
    - **Session 状态导致 TUI 崩溃**：`Session in status:failed remains bound ... crashing next TUI launch` ([#89167](https://github.com/openclaw/openclaw/issues/89167))。无 fix PR。
    - **Agent 长期僵死**：`Chronic agent failures on Telegram — LLM timeouts...` ([#88907](https://github.com/openclaw/openclaw/issues/88907))。无 fix PR。

- **高优先级 Bug (P1)：**
    - **Bedrock 升级回归**：`[Bug] Regression 2026.5.27→2026.5.28: "No API provider registered for api: bedrock-converse-stream"...` ([#88707](https://github.com/openclaw/openclaw/issues/88707))。无 fix PR。
    - **安全：API 密钥明文写入**：`[Bug]: models.json generator writes apiKey as plain string instead of secret-ref object` ([#88562](https://github.com/openclaw/openclaw/issues/88562))。无 fix PR。
    - **安全问题：推理令牌泄露**：`[Bug]: Reasoning tokens leak to chat channels...` ([#89473](https://github.com/openclaw/openclaw/issues/89473))。无 fix PR。
    - **Codex OAuth 僵死**：`Codex OAuth refresh failures can wedge an agent for hours...` ([#86215](https://github.com/openclaw/openclaw/issues/86215))。有 linked PR (#86213)，但未指定。
    - **子Agent超时无通知**：`Sub-agent run timeout does not notify parent session...` ([#89095](https://github.com/openclaw/openclaw/issues/89095))。无 fix PR。
    - **`NODE_USE_SYSTEM_CA=1` 回归**：`[Bug]: NODE_USE_SYSTEM_CA=1 breaks openai-codex auth...` ([#89655](https://github.com/openclaw/openclaw/issues/89655))。无 fix PR。

- **其他显著 Bug：**
    - **P2**: Feishu 会话 `/reset` 后无法使用 ([#89245](https://github.com/openclaw/openclaw/issues/89245))，QQBot WebSocket 重连后出站配置丢失 ([#88955](https://github.com/openclaw/openclaw/issues/88955))。
    - **P2**: 状态目录权限被更新检查重置 ([#89589](https://github.com/openclaw/openclaw/issues/89589))，Telegram 论坛主题回复丢失 ([#89352](https://github.com/openclaw/openclaw/issues/89352))。

**总结：** 今日没有高优先级 Bug 的 fix PR 被合并，这意味着所有关键问题仍处于待解决状态，稳定性风险较高。

## 6. 功能请求与路线图信号

- **模型与提供商支持**：
    - **阿里云模型工作室**：[#94418](https://github.com/openclaw/openclaw/issues/94418) 请求为阿里云的 Token Plan（团队版）提供原生支持。
    - **xAI 新模型**：PR [#89190](https://github.com/openclaw/openclaw/pull/89190) 提议添加 `grok-composer-2.5-fast` 模型，反映出快速跟进最新前沿模型的社区需求。

- **用户界面与体验**：
    - **子Agent会话可见性**：[#95295](https://github.com/openclaw/openclaw/issues/95295) 指出 Web UI 侧边栏隐藏了通过 `sessions_spawn` 创建的子 Agent 会话，这是对多Agent管理可视化的明确需求。
    - **初始化反馈**：[#89274](https://github.com/openclaw/openclaw/issues/89274) 建议在 `/new` 或 `/reset` 后显示已解析的模型和思考级别，提升用户对Agent状态的感知。

- **性能与可观测性**：
    - **OpenClaw 自身延迟**：[#88812](https://github.com/openclaw/openclaw/issues/88812) 作为 maintainer 创建的问题，请求测量和减少 OpenClaw 自身在频道回复中的前置模型延迟，这表明了官方对优化整体系统性能的意向。
    - **入站消息装饰**：[#95279](https://github.com/openclaw/openclaw/issues/95279) 提出为入站消息装饰提供一个可信的合约，以便消费者可以可靠地进行剥离/去重，反映了对更清晰、更健壮的数据管道的需求。

这些功能请求，尤其是 UI 可视化和性能优化方面的需求，结合已有的 PR，很可能成为下一版本关注的重点。

## 7. 用户反馈摘要

- **痛点：升级风险高**。大量用户报告从 `v2026.5.27` 到 `v2026.5.28` 的升级导致了关键功能（如 Bedrock）的回归，使得用户对升级保持警惕。`#88707` 提到这是一个“硬性的升级阻塞器”。
- **痛点：稳定性和可靠性不足**。用户反馈长期运行的服务出现 Gateway OOM 被杀（`#89315`）、Agent 因 OAuth 问题而僵死（`#86215`）、子 Agent 静默丢失（`#88856`）等问题，严重影响了生产环境的信心。
- **满意度：功能持续完善**。社区对添加新模型（如 xAI 模型 PR `#89190`）和扩展第三方服务（如 Ace Data Cloud 提供商 PR `#88680`）的热情很高，表明对 OpenClaw 作为聚合平台的扩展性感到满意。
- **痛点：配置和操作复杂**。用户对 Nix 模式下的写保护块（`#95135`）、模型提供商 ID 重命名后 OAuth 配置孤立（`#95136`）等问题感到困惑，表明管理配置的复杂性和潜在破坏性是用户的持续痛点。

## 8. 待处理积压

以下是一些需要维护者特别关注的、长期未解决或高优先级的 Issue/PR：

1.  **长期搁置的核心 Bug**：
    - Issue `#86215`：**Codex OAuth 刷新僵死问题**（P1，自 2026-05-24 起）。这是影响 Agent 可用性的核心问题，虽然活跃但无明确合并进度的 fix PR。
    - Issue `#87857`：**Agent 跳过 AGENTS.md 启动序列**（P2，自 2026-05-29 起）。关于 Agent 行为一致性的根本性问题，已标记为 `stale`，需要决策。
    - Issue `#87689`：**QMD 迁移期间禁止会话转录**（P2，自 2026-05-28 起）。关乎系统核心功能“Dreaming”的稳定性，需要一个支持的守护机制。

2.  **高风险/高价值的待合并 PR**：
    - PR `#89039`：**修复静默消息丢失**（P1）。这是对 `EmbeddedAttemptSessionTakeoverError` 的修复，对所有使用 Agent 内嵌功能的用户至关重要，目前处于重新审查循环中。
    - PR `#83988`：**修复 TTS 文本闪烁**（P2）。针对 Telegram 用户的长期痛点，已完全准备好等待审查，但已搁置一个多月。

3.  **安全审查积压**：
    - 多个涉及安全的 Issue（如 `#88562`、`#89147`、`#89228`、`#89473`、`#89589`、`#89607`、`#89254`）都在 `needs-security-review` 或 `needs-product-decision` 的队列中，表明安全审查是当前维护流程中的一个显著瓶颈。

**呼吁维护者优先关注这些积压事项，特别是修复关键退化和合并核心功能的 PR，以确保项目的健康发展。**

---

## 横向生态对比

好的，作为一名专注于 AI 智能体与个人 AI 助手开源生态的资深技术分析师，我将基于您提供的各项目动态数据，为您呈现一份深度横向对比分析报告。

---

## AI 智能体与个人助手开源生态横向对比分析报告 (2026-06-20)

### 1. 生态全景

当前，AI智能体与个人助手开源生态呈现出 **“头部项目巩固基础设施，尾部项目分化求存”** 的两极分化态势。以 **OpenClaw** 为代表的成熟项目，虽凭借庞大的社区积累了显著影响力，但正深陷于版本迭代引发的稳定性危机与维护容量瓶颈中，社区呼声与项目实际吞吐能力存在巨大鸿沟。与此同时，**IronClaw**、**ZeroClaw** 和 **CoPaw** 等第二梯队项目，则通过聚焦企业级特性（如多租户、安全加固）或特定场景优化（如第三方模型兼容性），展现出强劲的迭代活力。一个值得关注的趋势是，**安全问题**（如任意文件读取、API密钥泄露）正从单个项目的偶发事件演变为整个生态的普遍性风险，多个项目在同一天内收到了严重安全漏洞报告，这标志着行业正从功能驱动转向“安全与稳定性优先”的阶段。

### 2. 各项目活跃度对比

| 项目名称 | 今日 Issues 数 | 今日 PR 数 | Release 情况 | 健康度评估 |
| :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 116 更新 | 500 更新 | 无 | **待改进** - 极高的社区噪音，但维护响应严重滞后，Bug 关闭率极低。 |
| **NanoBot** | 7 (3 关闭) | 20 (3 合并) | 无 | **良好** - 社区反馈与修复节奏较为匹配，性能优化是重点。 |
| **Hermes Agent** | *高* | 50+ | **v0.17.0 发布** | **关注** - 重磅Release后出现显著的平台回归问题，修复及时但暴露了测试覆盖不足。 |
| **PicoClaw** | 4 更新 | 7 更新 | 1 Nightly | **稳定** - 开发节奏稳定，但部分重要 PR 处于停滞状态。 |
| **NanoClaw** | 0 | 6 (0 合并) | 无 | **健康** - 无新反馈，团队集中精力修复内部积压的安全和功能性 Bug。 |
| **NullClaw** | 1 (1 关闭) | 0 | 无 | **低活跃** - 仅有 Bug 报告，无代码贡献，项目依赖零星维护。 |
| **IronClaw** | *中* | 32 (4 合并) | 无 | **积极** - 密集的架构重构期，测试基础设施先行，为后续合并铺路。 |
| **LobsterAI** | 0 (5 关闭) | 0 | 无 | **维护模式** - 仅清理旧 Issue，无新功能或修复引入。 |
| **TinyClaw** | 1 (严重) | 0 | 无 | **高风险** - 报告了严重安全漏洞 (CVE级)，亟需维护者回应。 |
| **Moltis** | 0 | 1 (Dependabot) | 无 | **停滞** - 社区互动几乎为零，项目生命力堪忧。 |
| **CoPaw** | 8 | 15 (6 合并) | 无 | **健康** - 高吞吐量，社区反馈与开发响应形成正向循环。 |
| **ZeptoClaw** | 0 | 0 | 无 | **极低活跃** - 过去24小时无任何活动。 |
| **ZeroClaw** | 11 | 50 (少量合并) | 无 | **高速爬坡** - 密集开发期，大量高风险、高价值 PR 正在处理中。 |

### 3. OpenClaw 在生态中的定位

- **优势：生态规模与影响力**。OpenClaw 依然是生态中的“巨无霸”，其 Issue 和 PR 的日更新量远超其他项目，反映了其最庞大的用户基础和社区规模。它是生态内事实上的“核心参照”项目。
- **劣势：维护效率与版本质量**。与 **NanoBot** 和 **CoPaw** 等小团队但高机动性的项目相比，OpenClaw 的维护效率成为其最大短板，出现了“版本越新，问题越多”的负面体验。其修复速度远跟不上问题爆发速度，Bug 修复 PR 的合并周期很长。
- **技术路径差异**：与专注于 **架构重构**（IronClaw、ZeroClaw）或 **特定渠道深度优化**（CoPaw）的项目不同，OpenClaw 更像一个“全能选手”，试图覆盖所有功能和平台，这导致了其代码复杂度极高，也是维护困难的根源。
- **社区规模**：从数据量级来看，OpenClaw 的社区是当前生态中当之无愧的 No.1，但其引发的“噪音”和用户不满也最多。相比之下，**Hermes Agent** 和 **ZeroClaw** 的社区虽然规模较小，但显得更专业、更聚焦于技术讨论和高质量贡献。

### 4. 共同关注的技术方向

多项目同日涌现的需求，揭示了当前生态的技术热点：

1.  **安全加固成为头等大事 (涉及: OpenClaw, TinyClaw, NanoClaw, ZeroClaw)**：API密钥明文写入、任意文件读取、SSRF绕过、OAuth僵死等安全问题在同一天内被多个项目报告。这已不再是单个项目的偶发事件，而是普遍性的行业痛点。用户对**数据安全**和**系统隔离性**（Docker隔离、权限控制）的要求正被提升到前所未有的高度。
2.  **多通道体验的“最后一公里”优化 (涉及: Hermes Agent, CoPaw, ZeroClaw)**：项目普遍从“支持渠道”转向“优化渠道体验”。如 Hermes Agent 修复 Telegram 文本重叠，ZeroClaw 完善 Discord 斜杠命令，CoPaw 解决移动端侧边栏问题。这表明**提升已有渠道的用户界面一致性和交互流畅度**是当前各项目的重点。
3.  **Agent 协作与任务可靠性 (涉及: OpenClaw, NanoBot, PicoClaw)**：多个项目都在讨论“子Agent超时无通知”、“静默消息丢失”、“cron任务结果路由错误”等问题。这表明随着 Agent 能力增强，**多Agent编排的可靠性和可观测性**成为了共同的挑战。
4.  **模型兼容性与性能 (涉及: OpenClaw, NullClaw, CoPaw)**：从“支持更多模型”转向“让模型跑得更稳、更快”。如 CoPaw 修复第三方模型 API 兼容性问题，NanoBot 优化 Token 估算性能。这表明用户不再满足于“能用”，而是追求**极致的兼容性和响应速度**。

### 5. 差异化定位分析

- **OpenClaw**: **全栈多平台领导者**。目标用户是追求最大兼容性和社区支持的个人开发者和玩家。技术架构庞大，强调功能广度，但代价是复杂度和维护压力。
- **NanoBot**: **轻量级性能与子代理专家**。更侧重工具定义、性能缓存、子代理协作等“黑科技”，目标用户是追求效率和深度定制的开发者。
- **Hermes Agent**: **社交平台统一体验追求者**。强力推进 Telegram/WhatsApp/Discord 等核心社交渠道的体验一致性，v0.17.0 是重要里程碑。
- **IronClaw**: **企业级AI平台奠基者**。正从个人工具向企业平台转型，强调多租户、特性标志、托管基础设施，技术架构最前沿。
- **ZeroClaw**: **架构现代化与安全革新者**。在代码层面进行大规模重构（多数据库后端、权限强化），意图解决 Rust 生态下的性能和安全问题，是技术层面的“深度玩家”。
- **CoPaw**: **用户体验与生态扩展深耕者**。快速响应用户反馈，对第三方模型（如智谱）和第三方可观测性（如Langfuse）的兼容性做得较好，社区互动质量高。
- **PicoClaw / NanoClaw / NullClaw**: **特定场景的轻量级分支或学生项目**。通常规模较小，开发节奏不稳定，但可能在特定平台（如Windows、移动端）或特定用户群体（学生开发者）中拥有忠实用户。

### 6. 社区热度与成熟度

- **第一梯队 (快速迭代，高活跃度)**：**ZeroClaw**, **IronClaw**, **CoPaw**。这些项目社区讨论专业、开发者响应迅速，PR 吞吐量大，体现了极高的迭代效率和执行力。
- **第二梯队 (高影响力，但显疲态)**：**OpenClaw**。虽然数据量庞大，但呈现出“有量无质”的局面，大量 Issue 和 PR 无法被有效处理，可能导致核心贡献者流失。
- **第三梯队 (稳定发展，质量巩固)**：**NanoBot**, **Hermes Agent**, **PicoClaw**。这些项目保持了稳定的开发节奏，版本发布有计划性，但社区规模相对较小，处于从“做出来”到“做更好”的过渡期。
- **第四梯队 (低活跃或停滞)**：**NullClaw**, **LobsterAI**, **TinyClaw**, **Moltis**, **ZeptoClaw**。这些项目普遍缺乏维护者互动，社区活跃度极低，除非有重大版本发布或关键修复，否则将逐渐淡出视野。

### 7. 值得关注的趋势信号

- **信号一：安全不再是“可选”，而是“准入”条件**。TinyClaw 的 CVE 级别漏洞与多个项目的安全 Issue 提醒开发者，**在功能开发的前期就应嵌入安全设计**（Security by Design），否则会在生态竞争中迅速失去信任。
- **信号二：多通道竞争进入“平台之战”，而非“管道之战”**。项目间的主要差异不再是“你支持多少种渠道”，而是**你在每个渠道上提供的体验有多一致、多流畅**。Hermes Agent 的更新和趋势分析完美印证了这一点。
- **信号三：“Agent 协作”正从“技术 demo”走向“生产力瓶颈”**。当单一 Agent 能力趋于饱和，如何编排多个 Agent 协同工作、如何处理任务超时和失败、如何确保可观测性，将成为下一代智能体平台的核心竞争力。
- **信号四：性能优化的“微观战场”已经打响**。从 NanoBot 的 “token 估算缓存” 到 ZeroClaw 的 “流式叙述重复修复”，这些看似细微的优化，实则在用户体验上能带来质的飞跃。**对于AI助手而言，“T+1秒的延迟”带来的体验差距远比想象中巨大**。

**对于开发者的建议**：如果您追求最广泛的社区支持和功能覆盖，OpenClaw 依然是最佳起点，但需对其版本稳定性风险有所准备；如果您是注重开发效率和代码质量的中坚力量，**CoPaw** 和 **NanoBot** 可能更符合您的品味；如果您在探索企业级部署或架构前沿，**IronClaw** 和 **ZeroClaw** 则代表了未来的方向。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot 项目动态日报 | 2026-06-20

---

## 1️⃣ 今日速览

过去24小时，NanoBot 项目保持活跃：共处理 7 条 Issue，其中 3 条已关闭；PR 侧有 20 条更新，包含 3 条合并/关闭。社区聚焦于**性能优化**（工具定义 token 编码缓存）、**Telegram 富消息支持**以及**内存/子代理功能增强**。值得注意的趋势是：多位贡献者针对“推理努力度自动调节”“心跳任务输出路由”等用户体验问题提交了 feature request，项目整体健康度良好。

---

## 3️⃣ 项目进展

今日合并/关闭的重要 PR：

- **#4138** `[enhancement, valid] Add tools.file.enable to toggle built-in filesystem tools`  
  由 @niradler 提交，新增 `tools.file.enable` 开关，允许部署环境完全禁用内置文件系统工具，仅使用 MCP 服务器。该 PR 提供了与 `tools.exec.enable`、`tools.web.enable` 一致的配置能力，增强了安全控制。

- **#4230** `fix: set httpx timeout for streamableHttp transport`  
  修复了 `streamableHttp` MCP 连接因 `timeout=None` 导致无限等待的问题，现在具有默认超时保护。

- **#4246** `fix(session): delete_session also removes legacy path files`  
  解决了删除会话时遗留全局目录文件未被清理的 bug，防止旧数据被意外恢复。

- **#4296** `feat(sdk): expand Python SDK runtime controls`（仍处于 open 状态，今日有更新）  
  该 PR 扩展了 Python SDK，从简单的 `bot.run()` 变为更完整的开发者 API，新增 `RunResult` 元数据、稳定的 session/memory 客户端以及 Cursor/OpenAI 风格的回调钩子，有助于生态集成。

- **#4416** `[codex] feat(cron): support job model presets`（open，今日有更新）  
  为定时任务新增可选的 `model_preset` 字段，允许每个 cron job 使用不同的模型/provider/上下文窗口覆盖，无需修改 live agent 配置。

- **#4421** `perf(utils): cache tool-definition JSON in estimate_prompt_tokens`（今日新开）  
  直接对应 Issue #4420，将工具定义的 `json.dumps` + `tiktoken` 编码结果缓存起来，避免每轮 agent turn 重复计算，预计可显著降低响应延迟。

项目整体在**安全性**（文件工具开关）、**稳定性**（超时、会话清理）、**性能**（token 估算缓存）以及**功能**（cron 模型预设、SDK 扩展）方面均有进展。

---

## 4️⃣ 社区热点

**最活跃 Issue：**  
- **#4013** `Error calling LLM: stream stalled for more than 90 seconds`（已关闭，5 条评论）  
  用户反馈从 0.1.5post2 升级到 0.2.0 后出现流超时错误，严重影响使用。社区参与度高，最终关闭，但未明确 fix PR 指向，建议关注是否有隐藏回归。

**最热门 PR：**  
- **#4296** `feat(sdk): expand Python SDK runtime controls`（今日有更新，持续讨论）  
  该 PR 因涉及 SDK 架构变更，社区关注度较高，可能影响下游集成方式。

**分析：**  
用户对“模型流中断”“token 估算慢”等与响应速度相关的问题非常敏感。同时“心跳任务结果路由”和“Telegram 富消息”等新需求反映了社区对多渠道体验和自动化可靠性的期待。

---

## 5️⃣ Bug 与稳定性

| Issue / PR | 严重程度 | 描述 | 状态 | Fix PR |
|-----------|--------|------|------|--------|
| **#4013** `stream stalled >90s` | 严重 | 升级后 LLM 流超时，导致无法正常使用 | 已关闭 | 未标注 |
| **#4374** `SOUL.md/USER.md 读写不对称` | 中等 | WebUI 项目工作区的 bootstrap 文件读取路径和写入路径不一致 | 已关闭 | 未知（可能随 #4007 修复） |
| **#4389** `fallback 模型上下文窗口未适配` | 中等 | 当 fallback 模型窗口小于主模型时，prompt 不自动裁剪 | 已关闭 | 未知 |
| **#4417** `[CI/CD] test(mcp): use resolvable timeout URL` | 低 | 测试中使用了不可解析的 hostname 导致 CI 不稳定 | Open | #4417 本身即为修复 |

**关键风险：**  
#4013 虽然已关闭，但未公开明确修复方法，用户若仍遇到类似问题需关注底层流式超时逻辑是否彻底解决。

---

## 6️⃣ 功能请求与路线图信号

- **#4419** `Automatic reasoning effort escalation`  
  用户请求为推理模型添加“默认”和“升级”两级推理努力度，当任务复杂时自动切换。该功能与现存 `reasoningEffort` 配置字段兼容，实现成本较低，可能很快被纳入。

- **#4422** `Telegram sendRichMessage support`  
  用户 @zpljd258 同时提交了 Issue 和 PR (#4423)，增加 Telegram Bot API 10.1 的富消息支持（表格、任务列表、折叠块、数学公式）。由于 PR 已存在，极大概率合入下一个版本。

- **#4418** `Heartbeat tasks deliver results to the channel where added`  
  用户指出心跳任务的结果未投递到任务原始渠道，而是投递到最近活跃频道。该问题影响多渠道用户，已有对应 PR #3590（手动触发心跳）部分覆盖，但路由逻辑仍在讨论。

- **#4420** `性能优化：estimate_prompt_tokens 缓存`  
  已由 PR #4421 快速响应，预计会很快合并。

**路线图信号：**  
- **子代理** 相关 PR 密集（#4414 聚合结果模式、#4415 spawn 模型覆盖），表明团队正在强化子代理能力。
- **内存/归档** 方面 #4424 引入 provenance 上下文，提升事实归档准确性。
- **WebUI** 侧 #4329 新增内联终端 TUI，#4412 抑制常规 cron 通知，用户界面体验持续优化。

---

## 7️⃣ 用户反馈摘要

- **正面：**  
  - #4013 作者在关闭前表示“it's been very good (way to say ty)”，对 0.1.5post2 版本满意度高。
  - #4420 的提交者 @codeLong1024 在描述中感谢项目，并基于自身项目 `nanobee` 发现性能瓶颈，说明用户愿意深入代码自调试并回馈社区。

- **负面/痛点：**  
  - 升级到 0.2.0 后流超时（#4013）导致用户无法正常工作，体验断崖式下降。
  - 项目工作区的 SOUL.md 写入路径不对称（#4374）导致用户精心维护的文件被写到错误位置。
  - Heartbeat 任务结果路由错误（#4418）让多频道的用户感到困惑。

这些反馈提示项目需更注重**升级兼容性测试**和**多工作区场景验证**。

---

## 8️⃣ 待处理积压

- **#1945** `XMPP channel`（2026-03-12 创建，至今未合并）  
  XMPP 支持 PR，作者声称已在实际环境中运行，但长期未收到核心团队的 review。建议维护者评估并合并或给出明确延期说明。

- **#3590** `feat(heartbeat): add manual trigger command`（2026-05-02）  
  虽被提及，但一直没有合并。结合 #4418 的需求，手动触发心跳可部分解决结果路由问题，值得推进。

- **#3662** `fix(tokens): avoid network loads during estimation`（2026-05-06）  
  与今日合并的 #4421 目标类似但更早提交，然而未被采用。需要确认 #3662 的解决方案是否被 #4421 替代，或者需要进一步讨论。

- **#3591** `feat(dream): add update scope controls`（2026-05-02）  
  允许用户限制 Dream 功能的影响范围（memory/context only），防止 skill drift。长期未合并，建议推进。

**提醒维护者：** 上述积压 PR 均来自活跃贡献者，长时间搁置可能挫伤社区参与热情。建议安排定期 Triage 会议，对老旧 PR 给出决策。

---

**生成时间：** 2026-06-20  
**数据来源：** GitHub - HKUDS/nanobot（过去24小时活动）

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是根据提供的 Hermes Agent 项目数据生成的 2026-06-20 项目动态日报。

---

# Hermes Agent 项目动态日报 | 2026年6月20日

**分析师:** AI 开源项目分析师
**日期:** 2026-06-20

## 1. 今日速览

今日 Hermes Agent 项目活跃度极高，社区贡献密集。昨日发布了里程碑式的 `v0.17.0`，标志着项目进入全新阶段。社区反馈主要聚焦于新版本的稳定性问题，尤其是 Telegram 和 WhatsApp 平台适配器出现了回归性 Bug。项目组反应迅速，已针对多个关键 Bug 提交了修复 PR，同时也在积极回应社区关于功能配置和用户体验的新诉求。今日 PR 更新达到 50 条，显示出强大的社区推动力。

## 2. 版本发布

- **Hermes Agent v0.17.0 (v2026.6.19)**
    - **重磅发布**：这是“Reach Release”，标志着 Hermes Agent 在上一版本实现桌面端部署后，进一步在可访问性和平台覆盖范围上取得重大进展。
    - **规模惊人**：自 v0.16.0 以来，收到了来自 245 位社区贡献者的超过 1,475 次提交，合并了近 800 个 PR，关闭了 300 多个 Issue，版本迭代力度前所未有。
    - **破坏性变更与迁移**：
        - 由于本次发布规模巨大，涉及约 1,693 个文件变更，预计存在显著的配置和行为变更。
        - **强烈建议所有用户在升级前** 仔细阅读 [v0.17.0 完整发布说明](https://github.com/nousresearch/hermes-agent/releases/tag/v2026.6.19)。特别是 Docker 用户，需注意新版本中 WhatsApp 和 Dashboard 的已知问题（见下文）。
        - 已知的迁移问题包括：Telegram 消息渲染逻辑变更、WhatsApp 桥接依赖安装流程调整、以及可能存在的配置文件结构更新。

## 3. 项目进展

今日项目向前迈出了坚实的一步，多个长期存在的问题得到解决，关键功能得到增强：

- **安全性强化**：PR [#7073](https://github.com/nousresearch/hermes-agent/pull/7073) 修复了 WhatsApp 桥接 `bridge.js` 中的高危安全漏洞（V-004），该漏洞可能暴露关键信息，项目组已标记为 **P1 (最高优先级)**。
- **跨平台适配完善**：
    - PR [#46388](https://github.com/nousresearch/hermes-agent/pull/46388) 和 [#49563](https://github.com/nousresearch/hermes-agent/pull/49563) 完成了对 Signal 平台引用回复上下文的保存功能，增强了对会话上下文的处理。
    - Telegram 平台的 Bug 修复和功能增强是今日重点，多个相关 PR 正在处理中。
- **核心运行时稳定性**：已关闭的 PR [#30828](https://github.com/nousresearch/hermes-agent/pull/30828) 和 [#32544](https://github.com/nousresearch/hermes-agent/pull/32544) 分别强化了网关预算和上下文溢出保护，并修复了网关在每次轮次配置重载后缓存 `max_iterations` 过时的稳定性问题。
- **CLI 工具改进**：PR [#49564](https://github.com/nousresearch/hermes-agent/pull/49564) 修复了 `hermes doctor` 命令在非 UTF-8 系统语言环境下的崩溃问题，提升了工具兼容性。PR [#49566](https://github.com/nousresearch/hermes-agent/pull/49566) 修复了 `doctor` 命令中关于虚拟环境入口点的误报警告。

## 4. 社区热点

今日社区讨论热度最高、最为关注的几个议题集中在 WhatsApp 和 Telegram 平台适配器上。

1.  **WhatsApp Docker 桥接完全无法使用**
    - **链接**: Issue [#49569](https://github.com/nousresearch/hermes-agent/issues/49569)
    - **背后诉求**: 用户在升级到 `v2026.6.19` 后尝试在 Docker 环境中使用 WhatsApp，但遭遇了两个阻塞性 Bug，导致 WhatsApp 功能在 Docker 中完全不可用。这反映了社区希望关键平台适配器能开箱即用的强烈需求，任何配置障碍都会极大地影响用户体验和对新版本的接纳度。

2.  **Telegram 消息显示文本重叠**
    - **链接**: Issue [#49536](https://github.com/nousresearch/hermes-agent/issues/49536)
    - **背后诉求**: 这是一个影响视觉体验的 Bug，导致用户在 Telegram 上与 Agent 交互时，最终回复文本与流式预览文本重叠，需要手动刷新才能清除。用户期望无缝的流式交互体验，该 Bug 破坏了消息的完整性和可读性，影响了核心的用户体验。

3.  **社区老牌 PR 被重拾**
    - **链接**: PR [#12987](https://github.com/nousresearch/hermes-agent/pull/12987)， PR [#15872](https://github.com/nousresearch/hermes-agent/pull/15872)
    - **背后诉求**: 这两个由 Community Contributor `quinnmacro` 发起的关于 Mem0 记忆插件和系统时间戳的 PR 在今日被重新激活并更新。这显示了社区对更智能的上下文管理和时间感知能力的长期兴趣。`quinnmacro` 的持续投入也体现了核心贡献者对项目 AI 能力深度改进的承诺。

## 5. Bug 与稳定性

今日报告的 Bug 主要集中在版本 `v2026.6.19` 的回归问题上，以下是按严重程度排列的列表。

- **[Critical]** **WhatsApp Docker 桥接 (P2)**: Issue [#49569](https://github.com/nousresearch/hermes-agent/issues/49569) 报告两个阻塞性 Bug，导致 WhatsApp 在 Docker 中无法使用。
    - **已有 Fix PR**: 否（但有同类型的 Issue [#49561](https://github.com/nousresearch/hermes-agent/issues/49561) 描述了类似问题，同样无 fix PR）。
- **[High]** **Telegram 消息文本重叠 (P2)**: Issue [#49536](https://github.com/nousresearch/hermes-agent/issues/49536) 报告流式消息最终化时的显示错误。
    - **已有 Fix PR**: **是**，PR [#49537](https://github.com/nousresearch/hermes-agent/pull/49537) 已提交，通过强制所有最终化过程走富文本编辑路由来修复此问题。
- **[High]** **WhatsApp 桥接依赖安装失败 (P2)**: Issue [#49561](https://github.com/nousresearch/hermes-agent/issues/49561) 报告在升级到 v2026.6.19 后，Docker 中无法安装 WhatsApp 桥接依赖。
    - **已有 Fix PR**: 否。
- **[Medium]** **Dashboard 启动失败 (P2)**: Issue [#49567](https://github.com/nousresearch/hermes-agent/issues/49567) 报告按官方文档设置 Dashboard，启动后因绑定了 `0.0.0.0` 而被认证网关拒绝。
    - **已有 Fix PR**: 否。
- **[Medium]** **CLI `doctor` 命令误报 (P2)**: Issue/PR [#49566](https://github.com/nousresearch/hermes-agent/pull/49566) 修复了 `hermes doctor` 在通过 pip wheel 安装时产生的误报警告。
- **[Medium]** **MiniMax 视觉路由 (P3)**: Issue/PR [#49568](https://github.com/nousresearch/hermes-agent/pull/49568) 修复了 MiniMax-CN 提供商的视觉路由和 API base_url 问题。
- **[Low]** **插件安装器不兼容 (P2)**: PR [#49562](https://github.com/nousresearch/hermes-agent/pull/49562) 修复了某些插件安装器在没有 `pip` 的 Hermes 管理环境中失败的问题。

## 6. 功能请求与路线图信号

今日新提出的功能请求显示用户对平台配置粒度和多模态能力有显著需求。

- **可配置的交互相应表情**：
    - **链接**: Issue [#49570](https://github.com/nousresearch/hermes-agent/issues/49570)
    - **描述**: 用户请求让 Telegram 平台的“正在输入”、“成功”、“失败”等反应透出的表情符号变得可配置。
    - **路线图信号**: 该 Issue 立即得到了开发团队响应，**已有对应 PR [#49571](https://github.com/nousresearch/hermes-agent/pull/49571) 提交**。这表明项目对于提升用户自定义配置权限持开放态度，此类“小而美”的配置项很可能快速进入下一个版本。
- **视频多模态输入支持**：
    - **链接**: Issue [#49565](https://github.com/nousresearch/hermes-agent/issues/49565)
    - **描述**: 用户希望 Agent 能将视频（如 `.mp4`）作为原生多模态内容直接传递给支持视频的模型，而非仅提供文件路径。
    - **路线图信号**: 此功能涉及到核心 Agent 与模型的交互逻辑，技术挑战较大。虽然暂无关联 PR，但它指出了 Agent 能力演进的重要方向——拥抱多模态。参考已有 PR [#19331](https://github.com/nousresearch/hermes-agent/pull/19331)（添加cognee查询工具），项目正积极拓展数据处理和检索能力，未来支持原生视频输入并非不可能。
- **隐藏的配置项**：
    - PR [#39853](https://github.com/nousresearch/hermes-agent/pull/39853) 修复了 `notification_sources` 配置项从未被实际读入的 Bug。这揭示了项目文档中存在与实际代码逻辑不符的“幽灵配置”，提醒项目组需要加强文档与实际功能的同步审查。

## 7. 用户反馈摘要

从今日的 Issues 和 PR 中，可以清晰听到以下用户声音：

- **痛点: Docker 环境是个大问题**：在 `v0.17.0` 发布后，大量 Docker 用户反馈 WhatsApp 和 Dashboard 功能阻塞或无法使用（Issue #49561, #49567, #49569）。这“开箱即用”的体验不佳是新版本发布后最大的摩擦点。用户期待 Docker 镜像发布前能有更充分的集成测试。
- **体验诉求：精细控制与更好交互**：在 Telegram 平台上，用户不仅希望修复“文本重叠”这个 Bug（Issue #49536），还希望自定义机器人发出的表情（Issue #49570），这说明用户对交互体验的细节有很高要求，并有强烈的个性化配置意愿。
- **对 Agent 能力的更高期待**：用户不满足于 Agent 仅能处理文本和图片，视频文件的传输方式（Issue #49565）表明社区希望 Agent 能直接理解视频内容，对 Agent 的多模态处理能力提出了更高要求。

## 8. 待处理积压

以下是一些长期未响应但重要性高的 PR，可能影响到项目的安全性与AI能力深度，建议项目维护者们关注。

1.  **高危安全漏洞修复 (WhatsApp)**
    - **链接**: PR [#7073](https://github.com/nousresearch/hermes-agent/pull/7073)
    - **创建时间**: 2026-04-10
    - **摘要**: 修复 `bridge.js` 中的高严重性安全漏洞。目前已接近 **70天** 未合并，尽管今日有更新，但作为 **P1** 安全修复，其积压状态仍然值得重点关注。
2.  **Mem0 记忆插件 Ebbinghaus 衰减**
    - **链接**: PR [#12987](https://github.com/nousresearch/hermes-agent/pull/12987)
    - **创建时间**: 2026-04-20
    - **摘要**: 为 Mem0 插件引入基于艾宾浩斯遗忘曲线的时间序列记忆感知，是提升 Agent 长期记忆和上下文相关性的关键技术。两个月的积压可能延迟了 Agent 在复杂会话场景下的性能提升。
3.  **多项目管理器问题**
    - **链接**: PR [#15872](https://github.com/nousresearch/hermes-agent/pull/15872)
    - **创建时间**: 2026-04-26
    - **摘要**: 修复 Agent 感知到过时时间戳的问题，防止因时间感知错误导致的逻辑混乱。近两个月的积压，可能会持续导致用户遇到 Agent 时间判断不准确的困扰。
4.  **Cognee 查询工具**
    - **链接**: PR [#19331](https://github.com/nousresearch/hermes-agent/pull/19331)
    - **创建时间**: 2026-05-03
    - **摘要**: 添加只读的 `cognee_query` 工具，扩展 Agent 的知识检索能力。积压近 50 天，可能阻碍 Agent 在处理特定领域知识时的表现。

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw 项目日报 — 2026-06-20

## 1. 今日速览

过去24小时内，PicoClaw收到4条Issue更新、7条PR更新，并发布了1个nightly构建版本（v0.3.0-nightly）。项目社区讨论活跃，但多数PR和部分Issue标有`stale`标签，反映了一定的积压。一位用户报告了中文字幕“自己给自己整失忆”的对话历史丢失类Bug，值得关注。总体而言，项目保持稳定开发节奏，但关键Bug修复和大型功能PR等待合并时间较长，团队需注意积压风险。

## 2. 版本发布

- **nightly**: `v0.3.0-nightly.20260620.287853ab`  
  [GitHub Release](https://github.com/sipeed/picoclaw/releases/tag/v0.3.0-nightly.20260620.287853ab)  
  **说明**：自动化每日构建，基于`main`分支最新commit（`287853ab`），可能不稳定。无破坏性变更或迁移指南。建议用户谨慎用于生产环境。完整更新日志参见 [v0.3.0…main](https://github.com/sipeed/picoclaw/compare/v0.3.0...main)。

## 3. 项目进展

今日仅合并/关闭了1个PR：

- **#2956** [已关闭] `fix: preserve channel enabled state when merging security.yml`  
  - 作者：yuxuan-7814  
  - 描述：修复当用户在`security.yml`中添加凭证（如`telegram.token`）而未显式设置`enabled: true`时，`config.json`中已启用的通道被意外禁用的问题。  
  - 链接：https://github.com/sipeed/picoclaw/pull/2956  
  - **影响**：解决了配置合并逻辑中的回归，对多通道配置用户尤为重要。

其他6个PR仍保持开放状态，项目整体向前推进了一小步，但多个已标记`stale`的功能性PR（如Agent协作、SSRF修复）尚未合并。

## 4. 社区热点

今日最活跃的议题：

- **#2472** `[BUG] list_dir returns "invalid argument" on Windows...`  
  - 评论数：6，👍：1  
  - 作者：ut2or1 | 创建：2026-04-10 | 最后更新：2026-06-19  
  - 摘要：`list_dir`在Windows下因路径分隔符不符导致报错。用户要求修复平台兼容性。  
  - 链接：https://github.com/sipeed/picoclaw/issues/2472

- **#3150** `[BUG]它给自己整失忆了`  
  - 评论数：2，👍：0  
  - 作者：svier0 | 创建：2026-06-19 | 最后更新：2026-06-19  
  - 摘要：用户报告PicoClaw对话历史丢失（“失忆”），描述不完整，需进一步确认。中文字幕表明中国大陆用户活跃。  
  - 链接：https://github.com/sipeed/picoclaw/issues/3150

**分析**：Windows兼容性问题是长期痛点，而“失忆”Bug可能涉及会话持久化或上下文窗口管理，社区对这两类问题反馈强烈，维护者应优先回应。

## 5. Bug 与稳定性

| 严重程度 | Issue/PR | 描述 | 关联修复PR |
|----------|----------|------|------------|
| **高** | #3150 | 对话历史丢失，具体原因未知 | 无 |
| **中** | #2472 | Windows下`list_dir`因路径分隔符失败 | 无 | 
| **中** | #3143 (PR) | `web_fetch` SSRF防护绕过（ISATAP字面量嵌入私有IPv4） | 已有修复PR #3143，等待合并 |
| **低** | #3091 (PR) | `native_search`类型断言缺少`ok`检查，可能导致静默关闭搜索 | 已有修复PR #3091 |

此外，#2472 已存在2个月，虽有多条评论但未分配修复PR，建议维护者评估优先级。

## 6. 功能请求与路线图信号

以下功能请求/PR具有较高讨论热度或与路线图相关：

- **#348** `[Feature] General Attachment Support`  
  作者：Zepan | 创建：2026-02-17 | 更新：2026-06-19  
  目标：支持Telegram、Discord等渠道的附件（文本、多媒体）处理。该Issue优先级高，且为路线图项。  
  链接：https://github.com/sipeed/picoclaw/issues/348

- **#3114** `[Future Request] Telegram 渠道按对话类型（私聊/群组/频道）的权限分级控制`  
  作者：v2up-32mb | 创建：2026-06-12 | 更新：2026-06-19  
  需求：希望区分私聊/群组/频道权限，防止群内随意执行危险操作（如shell命令）。已获得1个👍。  
  链接：https://github.com/sipeed/picoclaw/issues/3114

- **#2937** `[PR] Feat/agent collaboration`  
  作者：afjcjsbx | 创建：2026-05-24 | 更新：2026-06-19  
  描述：首次引入Agent内部协作总线，包含邮件系统、协作线程、权限感知等。该PR较大，可能影响核心架构，但仍标记`stale`，需及时审阅。  
  链接：https://github.com/sipeed/picoclaw/pull/2937

**判断**：附件支持和Agent协作可能纳入v0.4.0；权限分级可能作为安全增强在后续版本中优先考虑。

## 7. 用户反馈摘要

- **Windows用户痛点**（#2472）：  
  `list_dir`在Windows上失败，用户需手动转义路径分隔符。用户“ut2or1”指出“platform-specific backslashes are passed directly to Go's fs.FS/os.Root”，期望跨平台一致性。  
  > *“这个Bug导致我在Windows上无法使用文件浏览功能，每次都要手动替换反斜杠……”*（评论摘录）

- **“失忆”问题**（#3150）：  
  用户“svier0”报告对话历史莫名丢失，虽未提供详细复现步骤，但情绪急切（中文标题）。类似问题曾在其他AI工具中出现，可能与上下文窗口清理或存储损坏有关。

- **安全边界需求**（#3114）：  
  用户“v2up-32mb”强调将PicoClaw加入群组后，如果没有按对话类型区分权限，任何人都可以要求机器人执行危险命令。  
  > *“目前只能控制谁可以用机器人，但不能控制在哪里可以用，这是很大的安全隐患。”*

## 8. 待处理积压

| 编号 | 类型 | 创建时间 | 最后更新 | 状态 | 备注 |
|------|------|----------|----------|------|------|
| **#2472** | Issue | 2026-04-10 | 2026-06-19 | `stale` | Windows路径Bug，2个月未分配修复 |
| **#2937** | PR | 2026-05-24 | 2026-06-19 | `stale` | Agent协作大型功能，等待review |
| **#3045** | PR | 2026-06-07 | 2026-06-19 | `stale` | Matrix用户ID解析修复，需合并 |
| **#3048** | PR | 2026-06-07 | 2026-06-19 | `stale` | MCP `add` 参数解析修复 |
| **#3053** | PR | 2026-06-08 | 2026-06-19 | `stale` | `lockStoreFile`类型断言panic修复 |
| **#3091** | PR | 2026-06-10 | 2026-06-19 | `stale` | `native_search`类型断言修复 |
| **#3143** | PR | 2026-06-18 | 2026-06-19 | OPEN | SSRF防护增强，最新活动 |

**提醒**：以上积压项均为社区贡献的修复或功能，长期未合并可能打击贡献者积极性。建议维护者安排周审阅会议，优先处理已标注`stale`且有测试覆盖的PR（如#3143、#3045）。

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

好的，以下是根据您提供的 NanoClaw 项目数据生成的 2026-06-20 项目动态日报。

---

# NanoClaw 项目动态日报 (2026-06-20)

**日报编写时间:** 2026-06-20
**数据来源:** github.com/qwibitai/nanoclaw

## 1. 今日速览

过去24小时，NanoClaw 项目未收到新的 Issue 报告，社区反馈趋于平静。然而，项目开发侧保持了高强度的修复与开发节奏，共有6个 Pull Request (PR) 处于待合并状态，远超日常水平。这表明团队正集中精力解决遗留的安全漏洞、功能缺陷，并引入新的容器运行时支持。项目整体处于“开发活跃、社区反馈稳定”的健康状态。

## 2. 版本发布

无。

## 3. 项目进展

过去24小时内没有 PR 被合并或关闭，所有活跃的 PR 均处于等待审查或测试的阶段。但从这些 PR 的内容看，项目正在多个关键方向取得实质性进展：
- **安全性强化：** PR #2799 针对因路径限制缺失导致的任意文件读取漏洞 (CVE-2026-29611) 提供了修复方案。
- **数据完整性：** PR #2820 修复了审批流程中交付目标信息缺失的问题，确保了审批记录的数据完整性。
- **平台兼容性：** PR #2809 引入了 Apple Container 运行时支持和远程 OneCLI 网关，朝向多平台部署迈出重要一步。
- **用户体验改善：** PR #2812 通过将长回复分块发送，替代原有的截断处理，提升了 Discord 平台上的用户体验。

## 4. 社区热点

由于缺乏评论数据，本次无法基于讨论热度进行量化分析。但从 PR 的内容来看，以下两个请求可能成为社区关注的焦点：

- **安全 CVE 修复 (PR #2799)**：
  链接: nanocoai/nanoclaw PR #2799
  分析：此 PR 直接关联一个 CVE 编号 (CVE-2026-29611)，预示着这是一个已公开的安全漏洞。对任何生产环境的用户而言，此类 PR 都至关重要，背后是用户对数据安全和系统健壮性的核心诉求。

- **Apple Container 运行时支持 (PR #2809)**：
  链接: nanocoai/nanoclaw PR #2809
  分析：支持在 macOS 上作为 Apple Container 运行是一项重大功能扩展。这显示出项目正在积极拥抱非 Linux 生态，特别是满足 Mac 开发者和希望在个人设备上部署 AI Agent 用户的需求。

## 5. Bug 与稳定性

过去24小时内未报告新 Bug。但当前待合并的 PR 中包含了多个关键修复，按严重程度排列如下：

1.  **极度严重 - 安全漏洞 (已修复，待合并)**
    - **文件读取漏洞 (CVE-2026-29611)**：`send_file` 函数缺少路径限制，恶意 Agent 可读取容器内的任意文件。修复方案在 PR #2799 中。
    - 链接: nanocoai/nanoclaw PR #2799

2.  **中等 - 数据完整性问题 (已修复，待合并)**
    - **审批记录目标缺失**：`pending_approvals` 表未保留审批消息的投递目标信息，导致审批列表功能异常。修复方案在 PR #2820 中。
    - 链接: nanocoai/nanoclaw PR #2820

3.  **低 - 功能缺陷 (已修复，待合并)**
    - **JSON 解析错误**：`safeParseContent` 函数未正确处理非对象类型的 JSON 原始数据（如字符串），导致部分场景下的数据读取失败。修复方案在 PR #2801 中。
    - 链接: nanocoai/nanoclaw PR #2801

4.  **易用性问题 (已修复，待合并)**
    - **Discord 消息截断**：Discord 回应超过2000字符时直接截断，导致信息丢失。修复方案在 PR #2812 中，通过分块发送解决。
    - 链接: nanocoai/nanoclaw PR #2812

## 6. 功能请求与路线图信号

虽然没有新的 Issue 提出需求，但本次 PR 列表强烈暗示了以下路线图动向：

- **安全架构重构**：PR #2799 和 PR #2801 的提交者均为 `sturdy4days`，表明安全性是当前的优先开发项。对文件路径和输入解析的强化，可能是更大范围安全审计的组成部分。
- **全平台部署策略**：PR #2809 的重点 (`Apple Container`) 表明项目正计划支持 macOS 作为一等公民，这与“个人 AI 助手”的定位高度吻合。
- **审批流程的稳定与可追溯性**：PR #2820 对审批记录空字段的修复，表明项目正在完善企业级工作流功能，提升系统的可审计性。
- **遗留功能的处理**：PR #2605 关注的子代理权限继承问题，已经存在近一个月，其被更新可能意味着团队正在重新评估其优先级。

## 7. 用户反馈摘要

尽管没有直接的用户评论数据，但从 PR 的描述中可以提炼出用户侧潜在的真实痛点：

- **安全担忧**：文件读取漏洞 (CVE-2026-29611) 的存在，意味着任何使用联网或第三方插件的用户都面临凭证泄露的风险。用户对此类安全事件的容忍度极低。
- **信息丢失**：Discord 用户在使用长回复时遭遇信息截断，严重影响对话的完整性和可用性。这说明在跨平台消息处理上，对不同平台限制的适配仍显不足。
- **管理混乱**：审批功能失效（无法查看审批历史）会严重阻碍用户对 Agent 行为的管控，尤其是在多 Agent 或多人协作场景下，会导致信任危机和管理混乱。

## 8. 待处理积压

本次数据中，以下 PR 需要维护者特别关注：

- **PR #2605:** [PR: Fix, follows-guidelines] feat: inherit parent agent permissions via OneCLI
    **状态:** 创建于 2026-05-24，至今未合并，但最近一次更新在 2026-06-19。
    **链接:** nanocoai/nanoclaw PR #2605
    **分析与提醒：** 这是一个功能性的 PR，旨在解决子 Agent 继承父 Agent 权限的问题。该 PR 已经存在近一个月，尽管近期有更新，但仍未合并。如果这是一个重要的功能需求，建议维护者尽快安排审查，避免功能积压，同时也能给社区一个明确的信号。

---
**总结：** 今日 NanoClaw 项目无新 Issue，但开发者集中精力处理了多项 Bug 和安全修复，整体开发势头强劲。社区用户主要受安全问题困扰，对平台适配和功能完整性有较高期待。建议核心团队优先合并并发布 CVE 安全修复 PR (#2799)。

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，我将根据您提供的 NullClaw 项目数据，为您生成一份结构清晰、数据驱动的项目动态日报。

---

# NullClaw 项目动态日报 | 2026-06-20

## 1. 今日速览

今日 NullClaw 项目活跃度中等偏低，社区焦点集中在 **本地模型集成** 和 **API响应稳定性** 两个核心痛点上。过去24小时内，项目没有新的代码贡献（Pull Request）或版本发布，但成功关闭了一个关于“本地Ollama模型回答不完整”的 Bug (#952)，表明团队仍在持续进行稳定性修复。同时，一个关于“NoResponseContent”错误的新 Bug (#967) 被报告，出现频率较高，可能成为下一个急需解决的稳定性问题。

## 2. 版本发布

**无**

过去24小时内未有新版本发布。

## 3. 项目进展

今日无新合并或关闭的 Pull Request，项目核心代码开发工作暂无明显推进。值得注意的是，**Bug #952** 在今日被标记为已关闭，这是一个积极的信号，表明开发者修复了本地模型集成中的一个关键问题。

-   **[Bug修复] 修复本地模型回答不完整的问题**：议题 #952 报告了使用 Ollama 运行的 Gemma 模型无法返回完整句子的故障。该问题已确认修复并关闭，提升了本地部署场景下的用户体验。([链接](https://github.com/nullclaw/nullclaw/issues/952))

## 4. 社区热点

今日社区讨论主要集中在已关闭的 **Bug #952** 上，该项目共获得 3 条评论，是当前最活跃的议题。

-   **热点议题：Local model using ollama returns incomplete answers (#952)**
    -   **诉求分析**：用户 `bloodgroup-cplusplus` 在使用 Ollama 运行 Gemma 模型时发现 Agent 回答总是“话只说一半”，并通过截图详细展示了问题。该议题获得了积极的关注和修复，反映了社区对于 **本地模型（Ollama、Gemma）集成质量和输出完整性** 有很高的要求。用户希望 Agent 能够稳定地生成完整、可用的回答，而不是截断的内容。([链接](https://github.com/nullclaw/nullclaw/issues/952))

## 5. Bug 与稳定性

今日共报告 1 个新 Bug，另有关闭 1 个旧 Bug。按严重程度排列如下：

-   **严重：error: NoResponseContent (#967)** (NEW)
    -   **描述**：用户 `svier0` 在 Win11系统上，使用 v2026.5.29 版本和 `Agnes-2.0-Flash` 模型时，报告出现 `error: NoResponseContent` 错误。该错误的出现频率非常高（>50%），严重影响了基本对话功能。
    -   **影响分析**：该问题影响 Windows 用户，且与特定模型 (`Agnes-2.0-Flash`) 相关联，可能导致使用了该模型的用户频繁遇到服务不可用的情况。
    -   **状态**：**未修复，暂无对应 PR**。([链接](https://github.com/nullclaw/nullclaw/issues/967))

-   **中等：Local model using ollama returns incomplete answers (#952)** (已关闭)
    -   **描述**：使用 Ollama 时，Gemma 模型返回的回答不完整。
    -   **状态**：**已关闭**。([链接](https://github.com/nullclaw/nullclaw/issues/952))

## 6. 功能请求与路线图信号

今日未有新的功能请求提交。所有新议题均为 Bug 报告，表明当前社区用户更关注于 **现有功能的稳定性和可靠性**。从报告的 Bug 来看（涉及 Ollama、特定模型、高频错误），项目下一版本的优化重点很可能指向 **API返回内容的健壮性** 和 **多种模型兼容性的深度测试**。

## 7. 用户反馈摘要

从今日的议题评论中，可以提炼以下用户反馈：

-   **痛点一：本地模型集成不稳定**
    -   用户 `bloodgroup-cplusplus` 指出了使用 Ollama + Gemma 模型时输出不完整的问题，表明即使用户成功启动了本地模型，其对话质量也可能不达预期，反馈过程需要“看图说话”，说明用户花了较多精力来调试和报告此问题。

-   **痛点二：高频的“静默”错误**
    -   用户 `svier0` 在 Issue #967 中详细描述了 `NoResponseContent` 错误的触发环境（Windows、特定模型、长响应时间），并明确指出此错误出现频率超过 50%。这表明该问题对用户的工作流程造成了严重干扰，且错误信息“NoResponseContent”本身对用户而言不够直观，需要开发者进一步调查根因并提供更好的错误提示。

-   **积极信号：用户愿意提供详细Debug信息**
    -   无论是 #952 附带的截图，还是 #967 中包含的具体版本、模型名、响应速度和复现概率，都反映出用户愿意投入精力协助项目改进。这是开源社区健康运转的良好迹象。

## 8. 待处理积压

当前最需要关注的待处理事项是昨日报告的紧急 Bug：

-   **Issue #967: [bug] error: NoResponseContent** - 该议题今日刚由用户报告，尚未获得任何维护者回应或 Labels。由于该问题严重性高（>50% 出现频率）且影响核心对话功能，建议维护者尽快介入，进行原因排查并给出初步反馈。([链接](https://github.com/nullclaw/nullclaw/issues/967))

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

好的，作为AI智能体与个人AI助手领域开源项目分析师，根据您提供的IronClaw项目GitHub数据，我已为您生成了2026年6月20日的项目动态日报。

---

### **IronClaw 项目动态日报 | 2026-06-20**

---

#### **1. 今日速览**

今日IronClaw项目开发活动极其活跃，PR提交量达到32条，其中21条处于待合并状态，显示出团队正在密集推进多项新功能。开发重点集中在**Reborn阶段**的架构重构和功能完善，特别是围绕扩展系统、入站流量管理（Slack/Telegram）和外部工具集成。同时，一条**Nightly E2E测试持续失败**的Issue表明基础稳定性存在待解决的问题，这是需要优先关注的健康度信号。整体而言，项目处于功能高速迭代期，但持续集成管道稳定性和工作负载管理是潜在风险点。

#### **2. 版本发布**

无新版本发布。

#### **3. 项目进展**

今日主要进展体现在**测试基础设施加固**和**文档规范**的完善上，为后续高风险的合并提供了质量保障。以下是今日合并的重要PR：

- **[CLOSED] (#5095) test(reborn-qa): add recorded fixtures**: 为Reborn QA系统添加了LLM追踪的录制和回放测试夹具，覆盖了连接、例程和网络抓取等关键场景，显著提升了测试的可重复性和覆盖率。
- **[CLOSED] (#5096) test(reborn-qa): port project-setup automation-workflow benchmarks to QA record/replay**: 将7个项目的设置基准测试迁移至录制/回放框架，使这些测试更可靠且更易于调试。
- **[CLOSED] (#5097) docs: add Reborn QA guidance to agent rules**: 新增了针对跨层和用户可见行为的测试指导，明确了录制QA与集成测试的使用场景，提升了团队协作规范。
- **[CLOSED] (#5092) ci(spike): A/B sccache (GHA) vs rust-cache on a heavy Reborn build**: 完成了Sccache与Rust-cache的A/B测试实验工作流，旨在优化大型Reborn构建的编译缓存效率。

**里程碑评估**：这些PR的合并巩固了项目的测试和CI基础设施，为后续待合并的21个功能特性PR（如Telegram/Slack Ingress、外部工具响应等）的安全发布铺平了道路。

#### **4. 社区热点**

今日在PR数量和复杂度方面，社区（主要是核心团队）的注意力和工作重点集中在以下几个大型功能模块的持续栈式开发上。虽然评论数未显示，但可以通过PR的“size: XL”标签和主题判断其关注度：

1.  **[#5081 - Add hosted single-tenant Postgres profile](https://github.com/nearai/ironclaw/pull/5081)** (serrrfirat): 这是为Reborn添加单租户Postgres配置文件的基础性工作，标志着Reborn向“托管预览”环境的演进。该PR涉及数据库迁移和多个模块，是基础设施层的热点。
2.  **[#5103 - feat(host-ingress): project IngressPolicy from the manifest](https://github.com/nearai/ironclaw/pull/5103)** (serrrfirat): 这是“清单驱动通道”工作的基石，通过将Ingress策略从硬编码枚举改为内联声明，大幅提升了灵活性和可扩展性。
3.  **[#5100 / #5093 - feat(reborn): project Telegram/Slack ingress from extension state](https://github.com/nearai/ironclaw/pull/5100)** (abbyshekit / serrrfirat): 这两个是相互关联的PR，旨在将Slack和Telegram的入站流量管理从硬编码逻辑转变为基于扩展状态的声明式配置，代表了Reborn在开放集成生态上的关键进展。

**分析**：社区热点主要集中在**Reborn架构的深度重构**，核心开发者正在将系统从静态、硬编码的配置向动态、模块化和声明式的方向演进。这反映了项目在追求更灵活、更易于扩展的企业级部署方案。

#### **5. Bug 与稳定性**

- **[高] (#4108) Nightly E2E failed**: 该Issue报告了**Nightly E2E测试持续失败**，失败工作流为“Full E2E / E2E (features)”。这是一个由自动化工具报告的稳定性回归问题。截至今日，此问题仍然**开放**且**无评论**，**尚无明确关联的修复PR**。
    - 链接: [Issue #4108](https://github.com/nearai/ironclaw/issues/4108)

**健康度信号**：E2E测试持续失败是严重的健康度红旗，它可能隐藏了近期代码合并导致的回归。维护团队应优先排查此问题，以防止问题蔓延到生产环境。

#### **6. 功能请求与路线图信号**

- **[#5091 - [enhancement] Unified feature-flag system for Reborn](https://github.com/nearai/ironclaw/issues/5091)** : 这是一项重要的功能请求，提出为Reborn创建一个统一的特性标志系统，以替代当前零散的`env`变量检查，支持动态切换、用户定位和灰度发布。该Issue被标记了多个模块标签（如 `module:M1-webui-product`），表明这是一个跨模块的、影响深远的架构改进。虽然尚无直接的PR，但它与当前PR中的许多动态配置工作高度相关，很可能被纳入未来的路线图。

**趋势判断**：项目明显正向**企业级特性**演进，包括灰度发布（Feature Flag）、多租户隔离（Postgres Profile）、可扩展的集成生态（Slack/Telegram Ingress）和更复杂的权限模型（#5062 工具权限覆盖）。这些信号表明IronClaw正在从个人开发者工具向企业级AI平台转型。

#### **7. 用户反馈摘要**

今日数据中，与用户直接相关的互动较少。主要反馈来源是**自动化工具**（如E2E测试失败报告）。社区的讨论主要集中在**开发者对架构的深度讨论**，而非用户端问题。这表明项目目前可能处于内部重写或重构期，用户层面的反馈较少。

#### **8. 待处理积压**

- **[#4002 - build(deps): bump the actions group](https://github.com/nearai/ironclaw/pull/4002)** (dependabot[bot]): 这是一个由`dependabot`自动创建的依赖更新PR，涉及16个Actions包的升级，已开放近一个月（自2026-05-24起）。长期未合并的依赖更新可能导致安全隐患和CI环境的版本落后。建议维护者尽快处理。
- **[#4108 - Nightly E2E failed](https://github.com/nearai/ironclaw/issues/4108)**: 如第5节所述，这是长期存在的E2E稳定性问题，且无任何响应或进展，亟需关注。

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

好的，这是为您生成的 LobsterAI 项目动态日报。

---

# LobsterAI 项目动态日报 (2026-06-20)

## 1. 今日速览

今日项目总体状态**平稳，活跃度较低**。过去24小时内，项目没有新的Issue或PR开启，也没有新版本发布，表明开发活动暂时进入间歇期。项目维护团队的主要动作是清除了5个标记为“stale”（过时）的旧Issue，这些Issue主要涉及Agent创建、配置修改时的数据丢失问题以及任务执行异常。批量关闭旧Issue有助于清理积压，保持仓库整洁，但并未引入新功能或修复。综合来看，项目目前处于一个静态维护阶段。

## 2. 版本发布

无新版本发布。

## 3. 项目进展

今日无新的Pull Request被提交或合并，项目核心代码和功能没有向前推进。今日的主要进展体现在Issue管理上：维护团队关闭了5个长期未更新的旧Issue，执行了仓库的日常清理工作。

## 4. 社区热点

今日社区讨论热度较低，所有关闭的Issue评论数均不多（2-3条）。从历史数据看，以下两个方向曾获得较多关注：

-   **[#1495] 无缘无故中断进程**：该Issue获得1个👍，反映了用户对任务执行稳定性的核心关切。用户报告在执行任务时进程意外中断，并质疑是客户端问题还是大模型问题，这指向了项目运行稳定性的一个潜在痛点。
    - 链接: [netease-youdao/LobsterAI Issue #1495](https://github.com/netease-youdao/LobsterAI/issues/1495)
-   **[#1468] ~ [#1470] 系列“未保存确认”Bug**：这一组3个Issue均由同一用户提出，集中反映了用户对**数据安全性**和**用户体验**的强烈诉求。用户普遍对在填写大量配置信息（如系统提示词、环境变量）时因误操作导致数据丢失感到不满，期望增加防误触机制。
    - 链接系列: [#1468](https://github.com/netease-youdao/LobsterAI/issues/1468), [#1469](https://github.com/netease-youdao/LobsterAI/issues/1469), [#1470](https://github.com/netease-youdao/LobsterAI/issues/1470)

## 5. Bug 与稳定性

今日无新报告的Bug。被关闭的5个Issue均为历史Bug，按严重程度排列如下：

-   **高（功能性中断）**：
    -   **[#1495] 无缘无故中断进程**：进程意外中断，严重影响任务执行的连续性和用户体验。
        - 链接: [netease-youdao/LobsterAI Issue #1495](https://github.com/netease-youdao/LobsterAI/issues/1495)
-   **中（数据丢失/体验受损）**：
    -   **[#1468] 创建Agent弹窗关闭时无未保存确认**：静默丢失用户已填写的Agent配置。
        - 链接: [netease-youdao/LobsterAI Issue #1468](https://github.com/netease-youdao/LobsterAI/issues/1468)
    -   **[#1469] Agent设置面板关闭时无未保存确认**：静默丢失对已有Agent的修改。
        - 链接: [netease-youdao/LobsterAI Issue #1469](https://github.com/netease-youdao/LobsterAI/issues/1469)
    -   **[#1470] MCP服务器配置弹窗关闭或按Escape时无未保存确认**：静默丢失重要的MCP服务器及环境变量配置。
        - 链接: [netease-youdao/LobsterAI Issue #1470](https://github.com/netease-youdao/LobsterAI/issues/1470)
-   **低（行为异常/误导）**：
    -   **[#1496] 任务显示完成，但是没有返回**：任务状态显示与结果返回不匹配，属于逻辑或前端展示问题。
        - 链接: [netease-youdao/LobsterAI Issue #1496](https://github.com/netease-youdao/LobsterAI/issues/1496)

**注意**：以上Bug已被关闭，但关闭原因（如“已修复”、“无法复现”或“暂时搁置”）未在数据中提供，因此无法判断问题是否已解决。

## 6. 功能请求与路线图信号

今日无新的功能请求提出。但从今日关闭的Issue（[#1468], [#1469], [#1470]）中可以提取出明确的**增强诉求**：
-   **增强用户输入保护**：用户在提交或离开包含未保存编辑的弹窗/面板时，系统应给出“未保存更改”的确认提示。这是一个典型的用户体验改进，预计未来版本可能会将此纳入，以提升产品的可靠性和用户满意度。

## 7. 用户反馈摘要

从今日关闭的Issue评论中，可以提炼出以下用户痛点：
-   **数据安全焦虑**：用户对在Agent创建、MCP服务器配置等高频操作场景下，因错误点击而导致“静默丢失”长时间的工作成果感到非常不满。这直接影响了用户对工具的信任感。
-   **运行稳定性诉求**：用户核心诉求是任务能够可靠地运行。进程“无缘无故”中断令用户感到困惑和挫败，并且难以定位是客户端、网络还是大模型的问题，增加了使用成本。
-   **状态一致性要求**：任务显示“已完成”但结果为空，这种状态不一致会严重干扰用户的工作流，使其无法判断任务实际是否成功。

## 8. 待处理积压

今日维护团队清理了部分积压，但并未引入新的待处理项。由于所有被关闭的Issue均为“stale”状态，且关闭原因未知，建议项目维护者对以下长期受关注的问题进行明确回应，无论最终决定是修复、推迟还是无法复现，都应在Issue中给出说明，以维护社区透明度：
-   核心稳定性问题：**[#1495] 无缘无故中断进程** – 获得用户认可，直接关联核心体验。
-   用户体验问题：**[#1468]**, **[#1469]**, **[#1470]** 系列“未保存确认”需求 – 反映了用户对高品质交互的期望，相关反馈较多，建议评估是否纳入后续迭代计划。

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是基于所提供数据的 TinyClaw 项目动态日报。

---

### TinyClaw 项目动态日报 | 报告周期：2026-06-19 至 2026-06-20

**核心分析**：今日项目活跃度较低，仅有一条新创建的 Issues，但内容涉及严重安全漏洞，需要项目维护者高度关注。社区其他方面（PR、版本发布）无新进展。

---

#### 1. 今日速览

今日 TinyClaw 项目活跃度较低，主要活动集中在一项新提交的安全问题报告上。项目未产生新的代码合并（PR），也无版本发布。当前项目的核心关注点应转向一个由外部安全研究人员报告的未认证接口漏洞，该漏洞被评为严重级别，可能导致本地文件泄露。整体来看，项目处于维护模式，但需优先处理这一突发安全风险。

#### 2. 版本发布

无新版本发布。

#### 3. 项目进展

今日无任何 Pull Requests 被合并、关闭或更新。项目在功能开发和 Bug 修复方面无实质性推进。

#### 4. 社区热点

- **[Security]** `prompt_file` 更新接口存在权限缺失漏洞
    - **Issue #285**：由社区成员 `YLChen-007` 于今日提交。该报告指出 TinyAGI `<= 0.0.20` 版本中，HTTP 管理 API 存在未授权访问风险，任何能够访问该 API 的客户端均可设置代理的 `prompt_file` 为任意可读的本地路径。
    - **分析**：该 Issue 是今日唯一且最重要的讨论热点。虽然目前尚无评论，但其所涉及的安全漏洞（任意文件读取）严重性极高，直接威胁用户数据和系统安全。社区的核心诉求是希望项目维护者能**迅速确认此漏洞，并发布修复版本或提供临时缓解措施**。
    - **链接**：[Issue #285](https://github.com/TinyAGI/tinyagi/issues/285)

#### 5. Bug 与稳定性

- **严重漏洞（未修复）**：**任意文件读取**
    - **Issue #285**：报告了一个严重安全漏洞。攻击者可通过未认证的 HTTP 管理 API，将 `prompt_file` 指向系统敏感文件（如 `/etc/passwd` 或配置文件），从而在后续 prompt 交互中泄露这些文件内容。
    - **影响范围**：所有运行 `<= 0.0.20` 版本且暴露了 HTTP 管理接口的 TinyAGI 实例。
    - **修复状态**：尚无关联的修复 PR。

#### 6. 功能请求与路线图信号

今日无任何新功能请求提交。项目路线图信号主要来自上述安全 Issue，暗示项目在 API 鉴权和输入验证方面存在不足，这可能会成为下一版本（`0.0.21`）发布的首要目标。

#### 7. 用户反馈摘要

今日无来自用户评论的有效反馈。唯一的 Issue 是由安全研究人员提交的漏洞报告，而非普通用户的日常使用反馈。

#### 8. 待处理积压

- **[严重安全漏洞] Issue #285**：此 Issue 虽为新提交，但因问题严重性高，应被视为最高优先级的待处理项。维护者需尽快响应，确认问题、提供临时解决步骤或发布紧急修复版本。目前该 Issue 尚无任何来自项目维护者或社区的回复。
    - **链接**：[Issue #285](https://github.com/TinyAGI/tinyagi/issues/285)

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis 项目日报 — 2026-06-20

## 1. 今日速览
- 项目过去24小时非常安静，未产生任何新Issue或版本发布，社区互动几乎为零。
- 唯一活跃的贡献来自自动化依赖更新机器人（Dependabot），提交了1个关于文档依赖升级的PR（#1133），目前处于等待审查状态。
- 整体活跃度评估：**极低**，项目进入典型的静默维护期，无功能开发或问题修复的动态。

## 2. 版本发布
*无新版本发布。*

## 3. 项目进展
今日无任何PR被合并或关闭。唯一活跃的PR为：
- **#1133 (OPEN)** — `chore(deps): bump astro from 6.3.3 to 6.4.8 in /docs`  
  由 Dependabot 自动发起，将文档目录下的 Astro 框架从 6.3.3 升级至 6.4.8。该版本包含多项上游修复与特性改进，但项目本身核心代码未受影响。  
  [查看PR](https://github.com/moltis-org/moltis/pull/1133)

> 说明：项目今日无实质性功能推进或Bug修复合并，核心代码状态未发生变化。

## 4. 社区热点
- **今日无高活跃度讨论**。PR #1133 未产生任何评论或反应（👍: 0），无社区成员参与互动。  
- 项目整体讨论热度处于低谷，可能因为近期无重大更新或用户关注点转向其他方向。

## 5. Bug 与稳定性
- 今日未报告任何新的Bug、崩溃或回归问题。  
- 无已存在的紧急Bug需要跟踪。

## 6. 功能请求与路线图信号
- 今日未收到任何新功能请求或Feature Request Issue。  
- 现有PR #1133 仅涉及依赖升级，不反映路线图变化。  
- 建议维护者关注用户长期未满足的需求（若有），但今日数据中无相关信号。

## 7. 用户反馈摘要
- 无有效用户反馈（今日无新Issue评论或讨论）。  
- 项目处于低交互状态，无法从数据中提取用户痛点或使用场景。

## 8. 待处理积压
- **无长期未响应的Issue或PR**。当前唯一待处理项为PR #1133，创建于今日，尚未获得维护者审查。  
- 建议维护者及时审查该依赖升级PR，避免因版本滞后引入安全风险或兼容性问题。

---

**项目健康度评估**：今日数据表明Moltis处于低活跃维护模式。虽然无负面问题，但长期缺乏社区互动和功能迭代可能影响项目生命力。建议团队考虑发布新版本或推动关键PR合并以提振活跃度。

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是基于 CoPaw (github.com/agentscope-ai/CoPaw) 提供的 GitHub 数据生成的 2026-06-20 项目动态日报。

---

## CoPaw 项目动态日报 | 2026-06-20

### 1. 今日速览

今日项目活跃度较高。过去24小时内，处理了8个Issues和15个Pull Requests，其中有6个PR已成功合入。尽管没有新版本发布，但社区贡献与维护响应均十分积极。**项目在稳定性（超时保护、cron任务优化）和用户体验（侧边栏交互、模型排序）方面取得了显著进展。** 社区讨论聚焦于第三方模型兼容性（如智谱AI、OMLX）与消息队列的健壮性，反映了项目生态扩展期的关键诉求。

### 2. 版本发布
**无**

### 3. 项目进展

今日合并/关闭的重要PR主要集中在以下方面，推动了项目稳定性和功能的实质性提升：

- **稳定性增强**：PR #5242 为上下文压缩过程增加了超时保护，防止因模型 API 挂起导致整个进程冻结；PR #5241 将定时任务的默认宽限时间延长，避免了长时间任务执行时错过预定任务的情况。
- **可观测性改进**：PR #5128 合并，成功将 Langfuse 可观测性数据按 Agent 循环进行分组，使调试和性能分析更加清晰。
- **第三方兼容性修复**：PR #5337 / #5338 修复了智谱AI (Zhipu AI) 模型连接测试失败的问题，确认是消息格式差异导致。
- **用户体验优化**：PR #5179 通过扩展关键词，提升了多智能体协作模式的触发成功率。

### 4. 社区热点

今日讨论最活跃的议题反映了用户对**第三方模型集成**和**核心交互稳定性**的集中关注：

- **模型提供商兼容性** (`#5208`, 6条评论)：用户反馈使用 `LongCat-2.0-Preview` 模型时，因 `reasoning` 块类型不兼容导致消息计数告警。这揭示了项目在适配非标准 OpenAI API 实现时仍存在隐性问题。
    - [Issue #5208](https://github.com/agentscope-ai/QwenPaw/issues/5208)

- **移动端/侧边栏交互** (`#5329`, 3条评论)：用户提出在手机浏览器上因侧边栏缩起而无法切换 Agent。该问题紧随其后的 PR #5334 提供了点击图标弹出菜单的解决方案，显示社区反馈与开发响应非常迅速。
    - [Issue #5329](https://github.com/agentscope-ai/QwenPaw/issues/5329)
    - [PR #5334](https://github.com/agentscope-ai/QwenPaw/pull/5334)

- **定时任务与主对话冲突** (`#5250`, 2条评论)：用户报告 cron 任务消息被错误地注入到用户主对话流中，导致 Agent 中断当前工作。这暴露了任务上下文隔离的机制问题。
    - [Issue #5250](https://github.com/agentscope-ai/QwenPaw/issues/5250)

### 5. Bug 与稳定性

今日报告的 Bug 主要集中在 **API 行为不一致** 和 **数据安全** 方面，按严重程度排列如下：

| 严重程度 | Issue ID | 描述 | 状态 / 修复PR |
| :--- | :--- | :--- | :--- |
| **严重** | [#5344](https://github.com/agentscope-ai/QwenPaw/issues/5344) | 当 Agent 忙碌时，通过 API 发送的消息返回 200 但被静默丢弃，损害系统可靠性。 | Open |
| **高** | [#5208](https://github.com/agentscope-ai/QwenPaw/issues/5208) | 特定模型返回 `reasoning` 类型块时导致功能退化（消息计数错误/注入失败）。 | Closed |
| **高** | [#5345](https://github.com/agentscope-ai/QwenPaw/issues/5345) | 自定义 OpenAI 兼容提供商（如 OMLX）不支持 Function Calling，限制了Agent能力。 | Open |
| **中** | [#5342](https://github.com/agentscope-ai/QwenPaw/issues/5342) | 当 LLM 调用失败时，现有工具结果修剪机制被跳过，可能导致上下文爆炸和级联故障。 | Open |

**总结**：项目当前面临的主要稳定性风险来自对非标准 API 的错误处理（`#5208`）和在异常情况下的消息队列健壮性（`#5344`）。`#5342` 提议的硬限制是重要的防御性措施。

### 6. 功能请求与路线图信号

今日的功能请求清晰地指向了**更好的用户体验**和**更强的边界控制**：

- **UI/UX 改进**:
    - `#5329` (Agent切换按钮)：已在 PR `#5334` 中得到解决，很可能包含在下一版本中。
    - `#5267` (模型列表自定义排序)：PR `#5336` 已提交，通过引入 `sort_order` 字段支持了自定义排序。该项功能有很高的落地概率。
        - [PR #5336](https://github.com/agentscope-ai/QwenPaw/pull/5336)

- **安全与隔离性**:
    - `#5346` (工具运行在 Docker 中)：这是一个来自首次贡献者的提案，旨在将工具执行与主环境隔离，显著提升安全性。
        - [PR #5346](https://github.com/agentscope-ai/QwenPaw/pull/5346)
    - `#5342` (工具结果硬限制)：作为防止上下文爆炸的防御性设计，若被采纳将显著提升系统健壮性。

### 7. 用户反馈摘要

从今日的 Issues 和 PR 评论中，我们可以洞察到以下用户诉求与痛点：

- **痛点：第三方模型集成兼容性不足**。用户在使用 `LongCat`、`OMLX` 等非标准 API 时遇到了多种问题（消息格式、工具调用），表明项目在兼容性测试方面需要更多努力。用户期望的不仅是“能用”，还要“无缝”。
- **痛点：移动端/窄屏场景体验受限**。用户明确提出了在手机浏览器上无法切换 Agent 的问题，凸显了移动端体验对项目普及的重要性。
- **肯定：安全边界受到开发者和用户关注**。PR `#5341` 提议将文件工具约束到工作区，这获得了社区正面响应，体现了开发者对安全性的重视。
    - [PR #5341](https://github.com/agentscope-ai/QwenPaw/pull/5341)
- **期待：更稳定的后台任务管理**。用户对 cron 任务与主对话的交互冲突感到困扰，希望有更清晰的任务隔离机制（如独立的上下文空间或明确的优先级规则）。

### 8. 待处理积压

今日数据中未发现明显被忽视的“积压” Issue。但以下两个首次贡献者的 PR 已打开但进入审核流程，需维护者持续关注以鼓励社区参与：

- **PR #5321 (scroll context strategy)**：一个大型特性 PR，引入了新的上下文管理策略，截止至报告日已打开超过24小时，需尽快安排Review，以决定是否合入及后续方向。
    - [PR #5321](https://github.com/agentscope-ai/QwenPaw/pull/5321)
- **PR #5346 (tool run in docker)**：作为`first-time-contributor`的 PR，需要提供更多指导或明确项目是否接受此类架构变更，避免长时间搁置打击贡献者积极性。
    - [PR #5346](https://github.com/agentscope-ai/QwenPaw/pull/5346)

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，作为AI智能体与个人AI助手领域开源项目分析师，我将根据您提供的ZeroClaw项目数据，为您生成2026年6月20日的项目动态日报。

---

# ZeroClaw 项目动态日报 | 2026年6月20日

**项目名称:** ZeroClaw (AI 智能体 & 个人 AI 助手)
**仓库地址:** [github.com/zeroclaw-labs/zeroclaw](https://github.com/zeroclaw-labs/zeroclaw)
**数据统计周期:** 2026年6月19日 00:00 UTC - 2026年6月20日 00:00 UTC

## 1. 今日速览

今日项目活跃度极高，社区与开发团队运作高效。共收到**11条** Issue 更新和**50条** Pull Request (PR) 更新，表明项目处于密集开发与迭代阶段。尽管没有新版本发布，但大量涉及高风险、高优先级特性的 PR 与 Issue 正在处理中，显示了项目在核心功能（如网关、运行时、安全）和用户体验（如 ZeroCode 界面）上的快速推进。社区互动热烈，开发者积极回应并提交修复方案，项目整体健康状况良好，处于快速爬坡期。

## 2. 版本发布

**今日无新版本发布。**

## 3. 项目进展

今日虽然合并的 PR 不多，但大量待合并的 PR 揭示了项目在多个关键领域取得的实质性进展。

- **核心基础设施增强：**
    - **多数据库会话后端 (#6893, perlowja)**: 一个大规模 PR，为多代理集群引入可选、功能门控的会话持久化后端（PostgreSQL, Oracle, MySQL, Db2），是提升项目企业级能力的重要一步。
    - **xAI OAuth 登录支持 (#7945, legokichi)**: 增加对 xAI/Grok 的 OAuth 认证支持，简化了用户通过 xAI 进行身份验证的流程，不再需要手动复制 API 密钥。
    - **自动清理临时文件 (#7923, hfut-GYH)**: 新增 `[files_cleanup]` 配置区块和规则，允许运维人员自动化清理临时文件，提升系统健壮性。

- **通信渠道与用户界面优化：**
    - **Discord 频道功能完善 (#7922, Nillth)**: 完成了 Discord 斜杠命令的所有功能，包括支持多语言命令描述和公会（Guild）范围注册，使 Discord 集成达到巅峰水平。
    - **WhatsApp Web 允许列表 (#7720, perlowja)**: 为 WhatsApp 通道增加了基于群组 JID 的允许列表功能，解决运维人员无法将机器人限定在特定群组的问题。
    - **ZeroCode 用户界面改进 (#8000, ConYel)**: 添加了“浏览模式”状态徽章、从文件树中自动退出浏览模式、制表符补全对齐等多项细节优化，显著提升了 TUI 用户体验。

- **稳定性与代码质量：**
    - **修复流式叙述重复问题 (#8014, singlerider)**: 修复了在原生工具调用前，流式叙述文本被重复发送的 Bug，优化了用户体验。
    - **裁剪函数 UTF-8 安全防护 (#7962, perlowja)**: 修复了 `truncate_output` 函数在 UTF-8 字符边界处切片可能导致崩溃的安全隐患。

**项目整体向前迈进了一大步**，尤其是在向多平台、企业级会话管理、渠道完善和代码健壮性方面进行了深度拓展。

## 4. 社区热点

今日最受关注的议题围绕 **网关安全** 和 **开发者体验/合规性**。

- **热点 Issue: #6127 "gateway: silent-fallback hardening"** (评论数: 7)
    - **链接**: [zeroclaw-labs/zeroclaw Issue #6127](https://github.com/zeroclaw-labs/zeroclaw/issues/6127)
    - **诉求分析**: 这是一个关于**网关静默降级（silent-fallback）加固**的追踪问题。社区开发者 `@WareWolf-MoonWall` 的合并条件触发了对 #6099 的后续讨论。核心诉求是进一步强化网关凭证解析的逻辑，确保在降级发生时不再是静默的，而是有明确的信号和隔离，反映了社区对系统安全性和可观测性的高度重视。

- **热点 PR: #8008 "feat(auth): add email-login subcommand"** (评论数: 11)
    - **链接**: [zeroclaw-labs/zeroclaw PR #8008](https://github.com/zeroclaw-labs/zeroclaw/pull/8008)
    - **诉求分析**: 此 PR 提议为电子邮件频道增加基于 OAuth2 设备授权码流程的登录子命令。虽然评论数未明确给出，但从 `mov-xound-glitch` 作者的身份来看，以及 `email-login` 功能的实用性，很可能引起了关于如何平衡命令行安全性与用户体验的讨论。这反映了社区对**安全且在无浏览器环境下进行身份认证**的需求。

## 5. Bug 与稳定性

今日报告了多个 Bug，其中一些已有修复 PR 跟进，显示团队响应迅速。

**按严重程度排列：**

- **S2 - 退化行为 (Degraded Behavior):**
    - **Issue #6651 [已关闭]**: [Matrix 频道重载导致内存泄漏](https://github.com/zeroclaw-labs/zeroclaw/issues/6651)。每次 `/admin/reload` 操作会泄漏约 1MB 内存，由上游依赖 `matrix-sdk` 的 `Arc` 循环引用导致。该问题已被关闭，表明可能已被修复或找到规避方案。
    - **Issue #8047 [开放]**: [`ReadSkillTool` 查找路径错误](https://github.com/zeroclaw-labs/zeroclaw/issues/8047)。`read_skill` 工具在技能目录查找技能文件时定位错误，导致智能体无法读取和使用技能。**目前尚无关联的修复 PR。**
    - **Issue #8039 [开放]**: [`fill-translations` 工具导致数据丢失](https://github.com/zeroclaw-labs/zeroclaw/issues/8039)。工具在修复“提示泄露”的翻译条目时，未能移除残留的多行续接符，导致生成的 `.po` 文件数据丢失。**目前尚无关联的修复 PR。**

- **运行时错误 / 逻辑错误:**
    - **PR #8023 [开放]**: [MCP Stdio 子进程泄露修复](https://github.com/zeroclaw-labs/zeroclaw/pull/8023)。合并分支的 PR，旨在修复 MCP 工具在每个心跳周期泄露 stdio 子进程的问题，这是一个明确的稳定性修复。
    - **PR #8014 [开放]**: [流式叙述重复修复](https://github.com/zeroclaw-labs/zeroclaw/pull/8014)。如前所述，修复了一段文本被重复叙述的Bug。
    - **PR #7345 [开放]**: [路径列举工具结果被错误路由到视觉模型](https://github.com/zeroclaw-labs/zeroclaw/pull/7345)。一个高风险的修复，修正了智能体循环中错误地将文件搜索/列表结果视作图像标记的问题，可能导致不必要的 API 成本。

## 6. 功能请求与路线图信号

今日涌现了多个新功能请求，其中一些与已有开发轨迹高度吻合，很可能被纳入后续版本。

- **潜在纳入 v0.8.2/v0.8.3 版本的功能：**
    - **Issue #8046 [开放]**: [可选的 Telegram Webhook 模式](https://github.com/zeroclaw-labs/zeroclaw/issues/8046)。用户 `dakaii` 提出支持 Webhook 作为 `getUpdates` 长轮询的替代方案。这与项目在 `v0.8.2` 和 `v0.8.3` 版本中对**多通道**及**网络/插件管理**的路线图信号相符，可能是一个受欢迎的补充。
    - **Issue #8044 [开放]**: [`/model --agent` 命令的权限强化](https://github.com/zeroclaw-labs/zeroclaw/issues/8044)。用户 `Nillth` 指出为所有用户更改模型的命令缺乏权限校验，这是一个重要的**安全与管理功能**请求，与项目对 **Gateway** 和 **Security** 的持续关注一致。

- **其他值得关注的功能请求：**
    - **Issue #8043 [开放]**: [RFC: 废弃独立 `aardvark-sys` crate](https://github.com/zeroclaw-labs/zeroclaw/issues/8043)。这是一个架构层面的提议，旨在将独立的 `aardvark-sys` crate 合并到 `zeroclaw-hardware` 中，属于**代码库清理与重构**，对长期维护有利。

## 7. 用户反馈摘要

从今日的 Issues 和 PRs 评论中，可以听到社区用户和开发者的真实声音：

- **点出权限与安全痛点**: 用户 `Nillth` 在 Issue #8044 中敏锐地指出了 `/model --agent` 命令存在的权限问题，强调“**任何能向智能体发送消息的人都可以更改模型**”，这反映了用户对**多用户环境下的安全治理**有着深切的关切。
- **表达对操作体验的改进诉求**: 用户 `ConYel` 在 Issue #8047 中描述了智能体“**无法找到技能**”的困惑，并提供了清晰的复现步骤，这体现了对 **AI 工具调用可靠性**和**反馈准确性**的高要求。
- **积极贡献解决方案**: 开发者 `perlowja` 和 `singlerider` 等持续活跃在修复一线，如 PR #8023 和 #8014，展现了社区**快速响应和主动解决问题的文化**，用户（开发者用户）对项目的参与度和归属感很强。

## 8. 待处理积压

以下为长期未响应或处于关键阻塞状态的重要 Issue/PR，需要维护团队重点关注。

- **Issue #6651 (已关闭 - 内存泄漏)**: [Matrix 频道内存泄漏](https://github.com/zeroclaw-labs/zeroclaw/issues/6651)。虽然已关闭，但作为影响稳定性的 S2 级内存泄漏，建议确保上游 `matrix-sdk` 的 fix 已应用且验证无误。
- **Issue #6127 (已关闭 - 追踪问题)**: [网关静默降级加固](https://github.com/zeroclaw-labs/zeroclaw/issues/6127)。此追踪问题关闭了，但其关联的 `#6099` 的后续统一合并条件仍需确认，确保所有相关 PR 的加固逻辑是完整和一致的。
- **PR #6893 (+30天，高风险，开放中)**: [多数据库会话后端](https://github.com/zeroclaw-labs/zeroclaw/pull/6893)。这是一个已经开放近一个月的超大规模 PR（size: XL，risk: high）。涉及核心基础设施变更，贡献者 `perlowja` 一直在更新，但未被合并。需要核心维护者优先安排审查，以避免与其他正在进行的重大重构（如网关安全）产生冲突。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*