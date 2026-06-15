# OpenClaw 生态日报 2026-06-15

> Issues: 260 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-06-15 03:43 UTC

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

好的，作为AI智能体与个人AI助手领域开源项目分析师，根据您提供的OpenClaw项目GitHub数据，我为您生成2026年6月15日的项目动态日报。

---

## OpenClaw 项目动态日报 | 2026-06-15

### 1. 今日速览

今日OpenClaw项目活跃度极高，社区反馈密集且质量较高。过去24小时内，项目处理了260条Issue更新和500条PR更新，显示出巨大的社区参与度和维护工作量。尽管大量PR处于待合并状态，但Bug修复类的PR（尤其是P1/P2级别）提交较多，表明项目正集中精力解决5月中下旬版本引入的稳定性和会话状态问题。**整体项目健康度处于“高活跃、高压力”状态，引擎核心迭代活跃，但稳定性风险需重点关注。**

### 2. 版本发布

- **v2026.6.8-beta.1** 已发布。此版本重点改进了**消息通道的健壮性**，特别是Telegram和WhatsApp：
    - Telegram: 支持结构化富文本（表格、列表、可折叠引文）、更安全的富媒体边界。
    - 改善了CLI后端交付逻辑，并迁移了遗留的原生草稿功能。
    - **未提及明显的破坏性变更或迁移注意事项**，建议用户关注上游更新日志以获取完整细节。

### 3. 项目进展

今日无已合并的重大PR。但从待合并的PR中，可以观察到项目在多个关键领域的积极修复：

- **安全加固**:
    - **PR #93152**: 修复了上下文文件遍历可能越界到系统目录的安全风险，限制了`AGENTS.md`等文件的搜索范围。
    - **PR #93131** (已关闭): 尝试为`web_fetch`工具添加数据出口过滤，防止API Key等敏感信息泄露（虽已关闭，但方向值得关注）。
- **核心稳定性修复**:
    - **PR #93117**: 修复了Anthropic思考模块恢复(`thinking-block recovery`)的重试逻辑，当API错误在控制平面“启动”事件后发生时，系统能正确重试。
    - **PR #93135**: 修复了Cron任务在工具不可用时，向用户发送“自诊”错误消息的问题，改为更明确的失败提示。
    - **PR #93138**: 修复了`sessions.resolve`工具在处理`sessionKey="current"`时，因直接向网关发送字面字符串“current”而导致`INVALID_REQUEST`错误的性能问题。
- **通道功能完善**:
    - **PR #93134**: 修复了飞书(Feishu)通道无法获取完整卡片消息内容的问题。
    - **PR #93148**: 修复了Telegram通道未将发送消息记录到SQLite数据库的遗漏。
    - **PR #93137**: 修复了iMessage通道在禁用回复功能时，仍可能错误携带回复元数据的问题。

**结论**: 项目正在积极修补5.22版本引入的回归问题和长期存在的低级错误。安全、会话状态管理和通道兼容性是当前推进的重点。

### 4. 社区热点

- **#85888 [OPEN] 定时任务在MiniMax 503错误下持续失败** (12条评论)
    - **链接**: [Issue #85888](https://github.com/openclaw/openclaw/issues/85888)
    - **分析**: 这是今日最受关注的Issue。用户报告OpenClaw在清晨（05:00-07:30 CST）触发MiniMax API定时任务时，连续遭遇503过载错误。问题核心在于，**手动触发同任务却能成功**，暗示问题不在API可用性，而在于OpenClaw自身的调度或认证机制在特定时间段内存在缺陷。该问题影响了大量依赖定时任务的用户，反映了对**任务调度可靠性和API提供者故障隔离**的强烈需求。

- **#86996 [OPEN] Active Memory + Codex 导致严重延迟和启动失败** (9条评论)
    - **链接**: [Issue #86996](https://github.com/openclaw/openclaw/issues/86996)
    - **分析**: 用户反馈在启用`active-memory`和Codex模型时，系统出现长延迟、钩子超时、启动中止等问题。这与#86519和#86508等Issue共同指向了**5.20版本更新后，'active memory'与'Codex app-server'结合使用时，会话管理和事件循环存在严重性能瓶颈**。

- **#86519 [OPEN] Telegram通道代理在5.20更新后重复回复2-10次** (9条评论)
    - **链接**: [Issue #86519](https://github.com/openclaw/openclaw/issues/86519)
    - **分析**: 这是一个典型的回归问题。升级到2026.5.20后，Telegram通道中的AI代理对用户消息重复发送2-10次相同的回复。即使用户升级到5.22版本，问题虽有所缓解（2-3次）但未完全解决。此问题严重影响了用户体验，社区高度关注`session.started`和`session.ended`事件的多次触发，导致消息重复发送。

### 5. Bug 与稳定性

今日报告的P1级Bug众多，主要集中在**会话状态损坏**和**消息丢失/重复**上。

| 严重程度 | Issue/PR 链接 | 摘要 | 状态 |
| :--- | :--- | :--- | :--- |
| **P1** | [Issue #86508](https://github.com/openclaw/openclaw/issues/86508) | Discord运行中出现`EmbeddedAttemptSessionTakeoverError`，会话文件被篡改。 | 无Fix PR |
| **P1** | [Issue #86519](https://github.com/openclaw/openclaw/issues/86519) | **[回归]** Telegram通道代理回复重复2-10次。 | 无Fix PR |
| **P1** | [Issue #86047](https://github.com/openclaw/openclaw/issues/86047) | **[回归]** Nextcloud Talk会话中Codex应用服务器/插件审批停滞导致工具执行超时。 | 无Fix PR |
| **P1** | [Issue #86827](https://github.com/openclaw/openclaw/issues/86827) | 群聊会话失败后进入“failed”状态，静默丢弃所有后续消息。 | 无Fix PR |
| **P1** | [Issue #86214](https://github.com/openclaw/openclaw/issues/86214) | Codex应用服务器在图像/工具请求过程中因`logs_2.sqlite`过大而关闭，导致会话中断。 | 无Fix PR |
| **P1** | [Issue #87109](https://github.com/openclaw/openclaw/issues/87109) | Gateway空闲时内存增长至1073MB+，在内存压力下Cron任务静默失败。 | 无Fix PR |
| **P1** | **PR #93135** | **修复**：Cron任务在工具不可用时失败，避免向用户发送误导性“自诊”消息。 | **有Fix PR (待合并)** |

**总结**: 5.20和5.22版本引入了连锁的P1级回归问题，严重影响了Telegram和Discord等核心通道的稳定性。好消息是，今日已出现针对Cron任务失败问题的修复PR (#93135)，表明团队正在积极应对。

### 6. 功能请求与路线图信号

- **轻量级部署模式**: **Issue #86881** 提出了“Gateway-lite模式”，希望在没有AI引擎的情况下运行OpenClaw作为纯粹的网关、Webhook和调度器。这是一个强烈的信号，表明用户希望**降低部署门槛**并**分离核心AI能力与消息路由**，这可能催生OpenClaw的模块化架构演进。
- **ElevenLabs语音支持**: **Issue #86434** 请求加入ElevenLabs作为“Talk”功能的实时语音提供商。这反映了用户对**高质量、低延迟语音体验**的追求，并希望摆脱对OpenAI和Google的依赖。
- **系统资源利用率优化**: **Issue #86381** 用户抱怨系统资源利用率低并频繁报错。这不仅是Bug，更是对性能优化的呼求，尤其在高并发场景或低配硬件上。
- **增强版Claude-CLI**: **PR #91476** 提议为`claude-cli`后端增加`ultracode`配置标志，以启用更高工作负载和动态工作流编排。这表明OpenClaw正积极与外部高级agent框架(Claude Code、Aider等)集成，扩展其能力边界。

### 7. 用户反馈摘要

- **满意方面**: 用户对v2026.6.8-beta.1在Telegram/WhatsApp通道上的健壮性改进表示肯定（如Issue #86047评论中提到期待修复版本）。
- **痛点与不满**:
    - **版本回滚困难**: **Issue #86616** 用户指出，升级到5.22后无法通过降级来规避新引入的`nativeHook.invoke`错误，因为配置文件中包含了不兼容的forward-incompat配置。**“新版本引入的回归问题导致回滚困难”是当前社区最大的不满之一**。
    - **缓存机制低效**: **Issue #86063** 用户深入分析了Anthropic的缓存机制，指出OpenClaw每次调用都会使对话历史缓存失效，导致不必要的计算和延迟。
    - **新版本忽视工作区文件**: **Issue #85773** 用户报告重新安装新版本后，AI代理完全忽略了工作区文件和技能配置，只进行泛泛回复，导致定制化Agent失效。

### 8. 待处理积压

- **长期未响应的P1 Bug**:
    - **[Issue #85822]**：用户在5月23日报告了Discord会话中48秒的静默延迟问题，至今无修复或明确回应。该问题涉及会话传递核心路径，风险较高。
    - **[Issue #86217]**：关于iOS后台定位权限缺失的安全报告（5月24日提交），至今无维护者回复。
- **长期未合并的P1 Fix PR**:
    - **[PR #87275]**：修复`gateway.ts`中`timeoutMs`为0时错误行为的问题，自5月27日起便处于“需真实行为证明”状态，无维护者推动。
    - **[PR #87330]**：修复`session_yield`可能导致父会话“冻结”的问题，同样自5月27日起处于停滞。

**提醒**: 上述积压的Issue/PR涉及会话状态管理和安全等关键领域，建议维护者优先分配时间进行审查或给出初步回应，以避免引发社区不满并延缓稳定性问题的修复。

---

## 横向生态对比

好的，作为AI智能体与个人AI助手领域开源生态的分析师，根据您提供的2026年6月15日各项目动态数据，我为您梳理并生成以下横向对比分析报告。

---

## 个人AI智能体开源生态横向对比分析报告 (2026-06-15)

### 1. 生态全景

今日的个人AI智能体与AI助手开源生态呈现出 **“核心引擎高速迭代，安全与稳定性矛盾凸显，差异化定位加速形成”** 的态势。一方面，以OpenClaw为代表的头部项目社区活跃度极高，正在快速修复新版本带来的大量回归问题，尤其是会话状态和通道兼容性；另一方面，安全问题成为贯穿全生态的焦点，多个项目（Hermes Agent、NanoClaw、ZeroClaw）均在此次日报周期内收到了高危安全漏洞报告或修复。同时，社区对轻量化部署、多Provider支持、以及将AI能力融入工程研发体系的需求明确增长，标志着该领域正从“原型验证”向“生产级应用”冲刺。整体看，生态正经历从“功能堆叠”向“架构合理化、安全健壮性”转变的关键阵痛期。

### 2. 各项目活跃度对比

| 项目名称 | Issues (新开/更新) | PRs (总/合并) | Release (过去24h) | 项目健康度评估 |
| :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | ~260条更新 | ~500条 / 0 | ✅ v2026.6.8-beta.1 | **高活跃，高压力**；社区贡献巨大，但产品稳定性风险突出，修复和回滚困难是主要痛点。 |
| **NanoBot** | 5条 | 46条 / 27条 | ❌ | **高度活跃且健康**；核心团队与社区协作高效，PR合并率极高，问题响应迅速。 |
| **Hermes Agent** | 6条 | 10条 / 10条 | ❌ | **高参与度，中等健康**；修复与安全增强并进，但待合并PR积压（40条）构成潜在风险。 |
| **PicoClaw** | 5条 | 9条 / 5条 | ✅ v0.2.9-nightly | **稳定，中等健康**；主要修复代码质量小问题和重构，但存在影响核心功能的Bug待解。 |
| **NanoClaw** | 7条 | 10条 / 5条 | ❌ | **极度活跃，需关注安全**；收到高危安全审计报告，核心架构有重大推进，是双刃剑。 |
| **IronClaw** | 22条更新 | 44条 / 15条 | ❌ | **高度活跃，正向进化**；社区与开发者共同贡献大量修复，正探索AI原生开发流程。 |
| **ZeptoClaw** | 0 | 0 | ❌ | **完全静默**。 |
| **NullClaw** | 1条 | 0 | ❌ | **静默潜伏期**；社区关注点单一（云服务认证）。 |
| **LobsterAI** | 2条 | 4条 / 1条 | ❌ | **低活跃，增长停顿**；所有活动和PR均处于“stale”状态，需维护者介入激活。 |
| **TinyClaw** | 0 | 0 | ❌ | **完全静默**。 |
| **Moltis** | 1条 | 0 | ❌ | **静默潜伏期**；社区有增强方向提出，但无代码贡献。 |
| **CoPaw** | 21条 | 16条 / 5条 | ❌ | **高度活跃**；社区反馈和贡献踊跃，生态国际化趋势明显（越南语支持）。 |
| **ZeroClaw** | 1条 (更新) | 50条 / 4条 | ❌ | **高PR量，合并瓶颈**；社区贡献多但合并速度慢，存在高风险生产环境Bug待解决。 |

### 3. OpenClaw 在生态中的定位

- **生态核心参照**：从社区规模看，OpenClaw（~260 Issues, ~500 PRs）的用户和贡献者基数远超其他项目，是当前生态的绝对核心和风向标。其功能和问题（如会话状态损坏、通道兼容性）是其他项目关注和参考的重点。
- **功能全面性与复杂性**：OpenClaw被定位为全功能平台，集成了最多的通道、功能和特性，但这也导致了极高的复杂性。今日日报反映出，其“全”的优势正转化为“稳定性维护”的巨大压力，新版本引入大量连锁回归问题（P1级Bug），版本回滚困难是当前最显著的劣势。
- **对比其他项目**：
    - **vs IronClaw**：IronClaw同样活跃，但其问题更聚焦于特定功能（如认证、审批流）和内在的“AI原生”工程改进，更像一个在特定领域深耕的“高级玩家”。OpenClaw则更像“大而全”的平台，需应对更广泛的使用场景带来的挑战。
    - **vs NanoBot**：NanoBot今日表现出更高的工程效率和稳定性。其高PR合并率和较少的Bug，表明其在控制复杂度方面做得更好，可能在架构清晰度和测试覆盖上更具优势。OpenClaw应在“快”与“稳”之间寻找更好的平衡。

### 4. 共同关注的技术方向

1.  **会话状态持久性与可靠性**：
    - 涉及项目：**OpenClaw、NanoBot、Hermes Agent、LobsterAI、CoPaw**。
    - 具体诉求：用户普遍面临会话数据损坏（被篡改）、任务静默失败、创建的任务重启后不可见等痛点。这表明，当前Agent的会话管理机制在应对并发、重启、异常恢复等场景时存在设计缺陷，亟需更健壮的持久化和状态机模型。

2.  **多Provider接入与厂商中立**：
    - 涉及项目：**OpenClaw、NanoBot、Hermes Agent、NanoClaw、CoPaw**。
    - 具体诉求：用户强烈希望避免被单一模型提供商锁定，要求支持更多API（如Kimi、OmniRoute）或提供身份认证（Azure、Claude订阅用户）。这反映了用户对“成本优化”、“能力互补”和“数据安全”的综合性要求。

3.  **安全与权限管理**：
    - 涉及项目：**OpenClaw、Hermes Agent、NanoClaw、ZeroClaw**。
    - 具体诉求：从文件泄露、凭证跨Profile读取，到气隙（Air-gapped）执行模式，安全问题已从“功能Bug”上升为“架构边界”问题。社区对Agent行为的可审计、可控制、权限最小化提出了极高要求。

4.  **低门槛部署与模块化**：
    - 涉及项目：**OpenClaw、ZeroClaw**。
    - 具体诉求：用户提出“Gateway-lite模式”或“气隙执行模式”，希望分离核心AI引擎与消息路由/调度器，实现更轻量的部署和资源隔离。这表明，生态正在探索更灵活、更适应嵌入式或边缘场景的架构。

### 5. 差异化定位分析

| 项目 | 功能侧重 | 目标用户 | 技术架构关键差异 |
| :--- | :--- | :--- | :--- |
| **OpenClaw** | **全功能堡垒**：集成最多通道、工具和AI引擎，追求极致功能广度。 | 所有人，从个人到企业，追求“一站式”解决方案。 | 高度集成的单体应用，通过插件和配置进行扩展，但架构复杂度高。 |
| **NanoBot** | **轻量级高速引擎**：专注于核心Agent能力和WebUI体验，追求高效率和低资源消耗。 | 对性能敏感的开发者，以及需要快速部署和迭代的团队。 | 架构清晰，PR合并率高，重视代码质量和用户体验，更像一个“精干”的引擎。 |
| **Hermes Agent** | **安全铁桶**：将安全作为第一优先级，大量PR和Issue围绕安全增强。 | 对数据安全有严格要求的企业级用户或安全研究员。 | 严格的权限控制和凭证管理机制，与桌面端深度集成，关注与Claude、Codex的互动。 |
| **NanoClaw** | **核心驱动先锋**：率先在Provider选择和Codex集成上进行架构级扩展，追求前沿能力。 | 愿意尝鲜、具备较强技术能力的创客和早期采用者。 | 强调Provider Seams（接入点）和Codex v2作为一等公民，预示了“插件化、多核心”的架构方向。 |
| **ZeroClaw** | **前沿探索者**：在安全隔离（气隙执行）、Agent自主意识（梦境模式）等前沿方向进行探索。 | 对Agent自主性和安全性有高要求的极客和研究者。 | 强调Unix Socket隔离、WASM插件，是生态中技术前瞻性最强的项目。 |
| **CoPaw** | **国际化与桌面场景**：专注填补海外渠道和国际化的空白，并持续优化桌面端性能。 | 企业微信、钉钉、Zalo等特定区域市场的用户和桌面端重度用户。 | 技术与主流对齐，但在地域化和本土场景适配上有独特优势。 |

### 6. 社区热度与成熟度

- **快速迭代期（高活跃、高变更、高热度）**：**OpenClaw、NanoBot、Hermes Agent、NanoClaw、IronClaw、ZeroClaw**。这些项目社区参与度极高，每天都有大量的Issues和PRs涌入。其中，**NanoBot**表现出更高的成熟度（高PR合并率），而**OpenClaw**和**ZeroClaw**则面临“高活跃”背后的稳定性和管理挑战。
- **质量巩固期（中等活跃、聚焦修复）**：**PicoClaw**。项目在完成初步功能后，正专注于清理技术债务和提升代码质量，社区活动以微小的修复和重构为主。
- **静默潜伏期/边缘化（低活跃、等待激活）**：**NullClaw、Moltis、LobsterAI**。这些项目社区互动稀少，几乎无代码合并，存在被用户遗忘的风险。**LobsterAI**尤为典型，所有PR/Issue都处于“stale”状态，项目如无核心维护者的推动作用，可能陷入停滞。

### 7. 值得关注的趋势信号

1.  **从“功能驱动”到“质量+安全驱动”**：本次日报中，大量核心Bug（会话损坏、任务静默失败）和严重安全漏洞是社区的热点。**这强烈暗示，AI智能体领域已经度过了“跑通功能”的阶段，用户和贡献者的核心诉求正转向“稳定运行”和“数据安全”**。对于开发者而言，这意味着在构建Agent应用时，必须将状态管理和安全架构作为一等公民优先设计。

2.  **“Provider中立”是必由之路，而非可选项**：从OpenClaw、NanoClaw到CoPaw，多Provider接入是跨项目的共性需求。**这表明，用户深度绑定的不再是某个模型，而是Agent平台本身**。Agent平台需要提供无缝的模型切换能力，而开发者则应设计出与模型无关的编排逻辑，以适应快速变化的AI模型生态。

3.  **“模块化”是解决复杂性和安全性的关键架构**：OpenClaw的“Gateway-lite”和ZeroClaw的“气隙执行模式”，共同指向了一个趋势：**将Agent的“思考/推理”核心与“行动/通信”边界进行解耦**。未来，我们可能会看到更多项目采用“独立推理后端 + 轻量网关”的架构，这不仅能降低部署门槛，还能实现更好的安全隔离和资源调度。

4.  **“AI原生工程”从口号走向实践**：IronClaw提出的“构建编码代理云工作流”，并尝试通过AI自动完成PR和代码审查，是本次日报中最具前瞻性的信号之一。**这意味着AI Agent不仅是生产力工具，也开始成为改进软件开发流程本身的工具**。对于团队而言，探索如何让Agent参与CI/CD、代码评审和任务管理，将成为下一个技术热点。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

好的，这是为您生成的 NanoBot 项目动态日报。

---

# NanoBot 项目动态日报 | 2026-06-15

## 1. 今日速览

项目目前处于 **高活跃度** 状态。过去24小时内，开发活动异常密集，共处理了 **46 条 PR**，其中合并或关闭了 **27 条**，显示出核心团队正全力推进多项重构和功能开发。同时，社区也报告了 **5 条 Issue**，其中包含两个较为关键的 Bug：API 返回零 token 用量和图像降级处理泄露文件路径，均已得到快速响应（其中一条已有关联的修复 PR）。项目整体健康度良好，维护者响应速度快，社区贡献活跃。

## 2. 版本发布

无

## 3. 项目进展

今日项目在 **配置架构、WebUI 功能、文档和稳定性** 等多个方面取得了显著进展。核心开发者 `chengyongru` 贡献了大部分合并的 PR，推动了内部架构重构和体验优化。

- **配置与架构重构：** 合并了 **#4314**，通过将共享的 Pydantic 配置基类提取到独立模块 `nanobot.config_base`，成功打破了工具配置模式的循环导入问题，为未来扩展工具模块奠定了更干净的架构基础。这属于一次重要的非功能性改进，提升了代码库的健壮性和可维护性。 [查看PR](https://github.com/HKUDS/nanobot/pull/4314)
- **新功能落地：** 合并了 **#4273**，正式新增了 `tools.exec.pathPrepend` 配置项，允许用户为执行工具指定优先于系统 `PATH` 的查找路径，增强了工具调用的灵活性。 [查看PR](https://github.com/HKUDS/nanobot/pull/4273)
- **用户体验与稳定性：** 
    - 合并了 **#4275**，新增了对配置文件解析失败的快速失败（fail-fast）机制，避免了使用损坏配置启动导致的后续奇怪错误。 [查看PR](https://github.com/HKUDS/nanobot/pull/4275)
    - 合并了 **#4339**，大幅改进了 WebUI 在移动端的响应式表现，优化了导航、间距和内容溢出等问题，提升了移动端用户体验。 [查看PR](https://github.com/HKUDS/nanobot/pull/4339)
    - 合并了 **#4248**，修复了 token 用量热图的时区和标签裁剪问题，并添加了回归测试。 [查看PR](https://github.com/HKUDS/nanobot/pull/4248)
- **文档与入门：** 合并了 **#4177**，重新整理了文档结构，使其对新手更加友好，并添加了多种安装和配置指南。同时合并了 **#4245**，移除了关于“nightly”分支的陈旧指南，使协作文档更清晰。 [查看PR #4177](https://github.com/HKUDS/nanobot/pull/4177) | [查看PR #4245](https://github.com/HKUDS/nanobot/pull/4245)

此外，`discord`、`飞书` 等多个渠道的底层代码也得到优化（如 #4277），提示词历史记录功能 (PR #4274) 和定时任务 (PR #4299) 也得到了会话维度的增强。

## 4. 社区热点

- **#4309 | [BUG] `/v1/chat/completions` API 总是返回零 token 用量**：此问题引发关注，因为它直接影响了与 OpenAI 兼容 API 的核心功能。用户期望在响应中看到准确的 token 消耗，但当前硬编码为 0。这是一个影响标准接口规范性的关键问题，且已存在至少3天，社区期待尽快修复。 [查看Issue](https://github.com/HKUDS/nanobot/issues/4309)
- **#4345 | [BUG] 图像降级回退机制导致模型“幻觉”并泄露文件路径**：此问题因其同时涉及 **隐私泄露** 和 **模型行为异常** 而受到关注。当模型不支持图片输入时，降级机制虽正确重试，但却向模型泄露了本地文件路径并提供误导性文本。此问题已被作者迅速提交PR修复，体现了社区协作的高效性。 [查看Issue](https://github.com/HKUDS/nanobot/issues/4345)
- **#4293 | [PR] 修复子代理结果注入问题**：此PR旨在解决子代理结果无法正确注入到主运行循环的问题，特别影响到定时任务（cron）等直接调用。因其对核心功能的稳定性至关重要，讨论热度较高，目前仍在开放中。 [查看PR](https://github.com/HKUDS/nanobot/pull/4293)

## 5. Bug 与稳定性

| 严重程度 | Issue / PR | 摘要 | 状态 |
| :--- | :--- | :--- | :--- |
| **严重** | [#4309](https://github.com/HKUDS/nanobot/issues/4309) | `/v1/chat/completions` 响应中的 token 用量（`usage`）字段始终返回 0。 | 未修复 |
| **严重** | [#4345](https://github.com/HKUDS/nanobot/issues/4345) | 图片降级处理会向模型泄露文件路径，并导致模型产生“看到图片”的幻觉。 | **已有修复PR** [#4346](https://github.com/HKUDS/nanobot/pull/4346) |
| **中等** | [#4250](https://github.com/HKUDS/nanobot/issues/4250) | Telegram 消息过长被分割时，会破坏代码块（fenced code block），导致渲染错误。 | 已关闭（已修复） |
| **中等** | [#4333](https://github.com/HKUDS/nanobot/issues/4333) | Anthropic 提供者对 `opus-4-8` 模型发送了已弃用的 `temperature` 参数，导致所有请求返回 400 错误。 | 已关闭（已修复） |

## 6. 功能请求与路线图信号

- **WebUI 配置同步（#4313 开放中）**：一个大型 PR，旨在实现 WebUI 设置面板与 `config.json` 的完全功能对等，新增了温度、工具限制、通道等配置的写接口。若合并，将极大降低非 CLI 用户的管理门槛，是向“配置即界面”迈出的重要一步。 [查看PR](https://github.com/HKUDS/nanobot/pull/4313)
- **自动化管理视图（#4330 开放中）**：另一位核心开发者提交了为 WebUI 增加自动化管理界面的 PR，允许用户更直观地管理、暂停、删除自动化任务。这表明项目正在构建更完善的用户自主管理能力。 [查看PR](https://github.com/HKUDS/nanobot/pull/4330)
- **Agent 启动图标（#4262 已关闭）**：用户请求在进入 Agent 模式时，能立即显示用户自定义的 `botIcon`，而不是先显示默认图片。此问题已关闭，表明功能可能已实现或正在实现中。

## 7. 用户反馈摘要

- **痛点：** 从 **#4309** 和 **#4333** 来看，用户对 **API 标准兼容性**（如 Token 用量）和 **提供商模型更新**（Anthropic 新模型参数问题）非常敏感。这些小细节的错配会直接导致用户集成失败或功能感知异常，影响开发体验。
- **使用场景：** **#4345** 的用户“BearMett”在测试中发现了一个非常边界但影响严重的场景：模型仅通过文字描述就“推断”出它“看到了”一张它从未实际接收到的图片，并将本地文件路径暴露在了聊天中。这反应了用户对模型行为透明度和隐私安全的极高要求。
- **满意点：** 多个 Issue 和 PR 的快速关闭（如 #4250, #4262, #4333）显示了项目对社区反馈的积极回应和高效迭代能力，这可能是让贡献者满意的关键因素。**#4273** 的新功能也直接响应了社区对工具配置灵活性的长期需求。

## 8. 待处理积压

- **#4309 零 Token 用量问题**：作为影响核心API的关键Bug，目前已开放3天，虽有关注但尚无公开的修复PR。考虑到 `nanobot serve` 是项目对外暴露的核心服务之一，建议维护者优先投入资源定位并修复此问题，以避免影响更多用户的集成工作。
- **#4293 子代理结果注入 PR**：此PR虽已经过社区讨论，但仍在开放中。它直接关系到定时任务等自动化流程的稳定运行，是一个功能性缺陷的修复，建议尽快完成代码审查和合并，以确保核心功能的健壮性。

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

好的，以下是基于您提供的数据生成的 **Hermes Agent 项目动态日报**，日期为 2026年6月15日。

---

# Hermes Agent 项目日报 | 2026-06-15

## 今日速览

Hermes Agent 项目今日呈现出**极高的社区活跃度**，修复与功能请求双线并进。在**安全加固**方面，社区提交了大量针对凭证存储跨配置文件泄露等关键漏洞的修复 PR，反映出项目在安全性上的持续投入。同时，**功能请求**依然保持热度，尤其是对 Claude 订阅集成的呼声很高。尽管 PR 待合并数量庞大（40条），但维护者响应迅速，今天合并了多个关键 Bug 修复。整体来看，项目社区生态健康，开发节奏紧凑，但需要集中精力处理积压的 PR。

## 版本发布

**无**

## 项目进展

今日合并/关闭了 **10 个 PR**，其中重要的合并包括：

- **`fix(terminal): fall back to .env for env_passthrough in local backend` (#46399)**: 修复了本地后端 `terminal.env_passthrough` 功能无法读取环境变量文件的问题，使本地开发环境与 Docker 后端行为保持一致，提升了开发体验。
- **`fix(ui-tui): stabilize embedded dashboard chat gateway` (#39840)**: 稳定了 TUI 中嵌入式仪表盘的 `/chat` 功能，修复了在特定竞态条件下可能导致的崩溃问题，提升了 TUI 的稳定性。
- **`fix(send_message): support Telegram rich messages` (#46438)**: 为 `send_message` 工具添加了 Telegram 富文本消息的原生支持，使得表格、任务列表等复杂格式能在 Telegram 中正确渲染，提升了 Agent 的交互质量。

这些合并代表了项目在**稳定性修复**、**开发体验优化**和**交互能力增强**上的持续进步。

## 社区热点

今日最受关注的议题当属 **Claude SDK 订阅集成**，这反映了 Claude 用户群体的强烈诉求。

- **[Feature]: Claude Agent SDK model provider with subscription OAuth (Codex-style) (#25267)**：获得了最多的 **20 个 👍** 和 **6 条评论**。用户 `YongboYu` 的核心痛点是，现有的 `anthropic` provider 需要 API Key 并按 token 计费，导致已有 Claude 订阅的用户需要“双重付费”。该请求要求效仿 Codex 的方式，允许用户通过 OAuth 使用其订阅账户。这代表了大量订阅用户对成本优化的迫切需求。
    - 链接: [Issue #25267](https://github.com/NousResearch/hermes-agent/issues/25267)

- 此外，**两项安全跟踪问题**（#46434, #46435）也受到关注，虽然细节未公开，但其 `P1` 的优先级表明安全问题始终是社区焦点。

## Bug 与稳定性

今日报告的严重 Bug 主要集中在以下几个方面：

1.  **关键安全问题 (P1)**: **Cron 任务配置孤儿化**（`#32091`）：在非默认 profile 下创建的 cron 任务会写入错误的配置文件路径，导致 Gateway 无法读取，任务“静默失踪”。这是一个设计缺陷，影响面较大，但目前**暂无公开的 Fix PR**。
    - 链接: [Issue #32091](https://github.com/NousResearch/hermes-agent/issues/32091)

2.  **关键安全问题 (P2)**: 今天报告了三个独立的安全漏洞，且**全部已有对应的 Fix PR**，显示开发者和社区反应迅速。
    - 凭证存储跨 Profile 泄露 (`#46411`)：`read_file` 工具可读取兄弟 profile 的凭证文件。
        - Fix PR: [#46417](https://github.com/NousResearch/hermes-agent/pull/46417)
    - 桌面端文件预览可读取凭证存储 (`#46413`)：桌面应用的 Electron IPC 守卫未拦截 Hermes 自身的凭证文件。
        - Fix PR: [#46430](https://github.com/NousResearch/hermes-agent/pull/46430)
    - Email Gateway 授权绕过 (`#46434`)：追踪中的安全公告，涉及邮件网关适配器的 `From:` 头伪造。
    - Dashboard 插件未授权执行 (`#46435`)：追踪中的安全公告，涉及插件加载器绕过 `plugins.enabled` 门控。

3.  **通用 Bug (P2/P3)**: 包括 `macOS temp/scratch 路径被错误地视为敏感路径`（#46412，已有 PR #46412）、`Hindsight 插件硬编码 UTC 时区`（#46424）等问题，均已有修复措施。

## 功能请求与路线图信号

社区提出的新功能请求表明，用户对 **桌面应用体验** 和 **更灵活的工作流** 寄予厚望。

- **桌面应用增强**:
    - **交付模式 (Deliverable Mode)** (`#46444`): 用户希望桌面应用能像 Telegram 等网关平台一样，自动将 Agent 生成的文件（如 ZIP、PDF）作为附件发送到聊天窗口，而非仅提供文件路径。此功能可显著提升桌面端的使用体验。目前**暂无明确关联 PR**。
    - **x86_64 (Intel) Mac 支持** (`#42199`): 当前桌面版仅提供 ARM64 构建，Intel Mac 用户无法使用。这是一个持续三周的请求，反映了对**硬件兼容性**的切实需求。**暂无回复或 PR**。
- **工作流与集成**:
    - **Kanban 看板的通知继承** (`#46443`): 已有 PR 提出为 Kanban 子任务添加继承根任务通知订阅的功能，这表明社区正在探索更复杂的、有状态的工作流自动化。

## 用户反馈摘要

- **用户痛点**:
    - **成本问题**：`#25267` 的用户明确表达了对“为同一个模型支付两次”的不满，这是当前 `anthropic` provider 为订阅用户带来的核心摩擦。
    - **UI 缺失功能**：`#45865` 和其后续 `#46445` 的用户反映，无法在 TUI 的账户面板中移除已连接的账户，尤其是 Anthropic OAuth 账户，这说明 UI 的完整性和易用性有待加强。
    - **配置孤儿化**：`#32091` 的用户遭遇了 Cron 任务“静默失败”的问题，这严重影响了自动化功能的可靠性和用户对系统的信任。
    - **兼容性限制**：`#42199` 的用户因 Intel Mac 无法运行桌面版而感到困扰。
- **使用场景**:
    - 用户在尝试将 Hermes 集成到 **带订阅的已有服务** 中（如 Claude Pro）。
    - 用户在使用 Hermes 进行 **基于 Profile 的多环境多身份管理**。
    - 用户期望 Hermes Agent 能成为一个**类似交付物的终端**，能直接分享文件。

## 待处理积压

以下为长期未响应或待解决的重要议题，提醒维护者关注：

- **[Issue #32091] Cron jobs created from a profile-scoped agent session go to a jobs.json the gateway never reads (P1)**：此问题已存在 21 天，影响核心自动化功能，但目前**既无社区维护者回复也无相应 PR**，需要优先处理。
    - 链接: [Issue #32091](https://github.com/NousResearch/hermes-agent/issues/32091)
- **[Issue #31584] Treat memory-context as background context, not authoritative user-message content (P2)**：该特性请求涉及 Agent 对记忆上下文的处理方式，关系到 Agent 行为的准确性和安全性，已存在 22 天，需要决策层讨论。
    - 链接: [Issue #31584](https://github.com/NousResearch/hermes-agent/issues/31584)
- **[Issue #42199] Request: x86_64 (Intel) macOS build for Desktop App**：已存在 7 天，是关于桌面版兼容性的基础需求，应评估构建支持的成本和必要性。
    - 链接: [Issue #42199](https://github.com/NousResearch/hermes-agent/issues/42199)

---
**报告结束**

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，根据您提供的 PicoClaw 项目数据，我为您生成了 2026-06-15 的项目动态日报。

---

### PicoClaw 项目动态日报 | 2026年06月15日

#### 1. 今日速览

项目今日活跃度较高，共处理了 5 个 Issues 和 9 个 Pull Requests，并发布了新的 Nightly 构建版本。社区贡献者非常活跃，主要集中在修复小错误（如文件描述符关闭、错误处理）和代码重构（日志系统标准化）上。值得注意的是，**多个关于“AllowFrom”和命令解析的 BUG 报告正处于讨论中，需要维护者重点关注**。整体来看，项目维护节奏稳定，正向提升代码质量和扩展性迈进。

#### 2. 版本发布

- **Nightly Build (v0.2.9-nightly)**
  - **发布概述**: 项目发布了 `v0.2.9-nightly.20260615.13a38bd1` 版本。这是一个自动化构建版本，旨在提供最新的代码集成与测试，项目官方提示其**可能不稳定，使用需谨慎**。
  - **更新内容**: 该版本包含了截至当前所有已合并至 `main` 分支的变更，涵盖了今日合并的多项代码修复和重构工作。
  - **迁移注意事项**: 对于生产环境的用户，建议等待稳定版本发布后再进行升级。若需要使用新功能进行测试，请备份好配置文件及数据。

#### 3. 项目进展

今日项目在代码质量和稳定性方面取得了显著进展，具体表现为：

- **稳定性修复**: 多个 PR 针对代码中潜在的资源泄漏和错误处理不完善进行了修复。
    - **PR #3124**: 修复了 TTS (文本转语音) API 返回错误状态时，`io.ReadAll` 错误被忽略的问题，提升了错误诊断信息的可靠性。 ✅ 已合并
    - **PR #3123**: 修复了文件系统工具中，关闭目录文件描述符时错误被丢弃的问题，使代码意图更明确。 ✅ 已合并
    - **PR #3122**: 修复了 `appendJSONLRecords` 函数中，写入文件后 `Close()` 错误被忽略的问题，防止潜在的磁盘写失败被静默处理。 ✅ 已合并
    - **PR #2904**: 解决了 agent 循环中的 goroutine 泄漏和 panic 问题，通过同步化的 `recover` 流程提升了 agent 的运行稳定性。 ✅ 已关闭
- **代码重构与现代化**: 项目持续进行代码内部清理。
    - **PR #3121**: 将 OpenAI 兼容模块中 `log.Printf` 替换为项目统一的结构化日志记录器，有助于提升日志的可观测性。 ✅ 已合并
- **功能扩展**:
    - **PR #3126**: 改进了 Launcher 的调试诊断，使其可以更清晰地记录白名单绕过行为，提升了安全性审计能力。 🔄 待合并

#### 4. 社区热点

今日社区讨论的热点集中在 **配置解析和功能交互的 BUG** 上，表明用户在使用常规功能时遇到了障碍。

- **Issue #3041**: 关于 `mcp add` 命令无法正确解析全局标志，导致添加 HTTP/SSE 服务失败的问题。该问题被标记为 Bug，且已影响 Linux x86_64 用户，社区用户 `carlosprados` 提供了详细的复现步骤，表明该问题已被清晰定位。👉 [查看详情](https://github.com/sipeed/picoclaw/issues/3041)
- **Issue #3044**: 报告当 Matrix 用户 ID 包含冒号时，`allow_from` 配置完全失效。这直接影响使用标准 Matrix ID 的用户进行频道权限控制。👉 [查看详情](https://github.com/sipeed/picoclaw/issues/3044)

#### 5. Bug 与稳定性

今日上报的 Bug 主要集中在功能交互和平台兼容性上，按严重性排列如下：

- **[严重] Issue #3041**: `mcp add` 命令解析错误，导致无法添加 HTTP/SSE 类型的 MCP 服务，直接破坏了 MCP 功能的使用。
- **[严重] Issue #3044**: Matrix 频道 `allow_from` 功能失效，导致使用标准 Matrix ID 的用户无法被正确识别，破坏了频道权限控制。
- **[中等] Issue #3125**: 在使用 `.security.yml` 配置文件后，`web_search`（Brave API）工具静默失效。虽然没有报错，但功能不可用，会误导用户。
- **[低] Issue #3090**: PicoClaw 面板在 iOS 16.4 以下版本的 Safari 浏览器上无法正常工作，存在兼容性问题。

**修复追踪**：
- **Issue #3041** 和 **#3044** 尚未有关联的修复 PR。
- **Issue #3125** 尚未有关联的修复 PR。

#### 6. 功能请求与路线图信号

今日社区提出的新功能需求和正在开发的 PR 共同指向了项目**扩展性与易用性**的方向：

- **Out-of-tree 频道扩展**: **PR #3120 (待合并)** 尝试添加 `RegisterChannelSettings` 钩子，允许第三方模块无需 fork 主项目即可实现自定义频道。这是一个重要的架构性增强，表明项目正在为模块化生态做准备。
- **远程 Agent 模式**: **PR #3118 (待合并)** 为 `picoclaw agent` 命令增加了通过 WebSocket 连接远程 Pico 实例的功能。这提升了 PicoClaw 作为代理客户端的灵活性。
- **添加新 Provider**: **Issue #2978 (已关闭)** 用户请求添加 `OmniRoute` 作为新的 Provider。虽然当前 Issue 已关闭，但此类请求是项目 Provider 矩阵扩展的持续动力。

#### 7. 用户反馈摘要

从今日的 Issues 评论中可提炼出以下用户痛点：

- **配置复杂性**: 用户 `urtaevS` (**Issue #2978**) 请求添加 Provider，并询问“如何将其添加到配置中”，反映了用户在自定义配置方面的学习成本。
- **功能失效的挫败感**: 用户 `weissfl` (**Issue #3044**) 和 `Giordano10` (**Issue #3125**) 报告的功能在升级或正常使用后突然失效。特别是 `allow_from` 问题的用户，其“Messages are silently rejected”，这种无提示的失败带来了极差的用户体验。
- **命令行工具的易用性**: 用户 `carlosprados` (**Issue #3041**) 在尝试使用 `mcp add` 命令时，遇到了参数解析错误。该用户提供了非常具体的复现步骤，体现了对维护者的尊重和对问题解决的期待。

#### 8. 待处理积压

以下 Issue 或 PR 处于“待处理”状态，可能需要维护者关注：

- **Issue #3041 ( `mcp add` Bug )**: 已存在 8 天且处于活跃状态，是影响核心功能的 Bug，应优先处理。
- **Issue #3044 ( `allow_from` Bug )**: 已存在 8 天，影响 Matrix 用户，涉及到安全与权限核心功能。
- **Issue #3090 ( iOS 兼容性 )**: 已存在 5 天，虽然严重性不高，但可能影响一部分移动端用户的体验。
- **PR #2975 ( Telegram 回复触发 )**: 已存在 16 天，处于停滞状态。这是一个增强 Telegram 频道交互体验的功能性 PR，值得审视。

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，我已根据您提供的 GitHub 数据，为您生成了 Nanoclaw 项目的 2026-06-15 项目动态日报。

---

## NanoClaw 项目动态日报 | 2026-06-15

### 1. 今日速览

今日项目活跃度极高，共收到 7 条 Issue 和 10 个 PR，显示出社区贡献和反馈非常积极。值得关注的是，项目迎来了 **3 个由外部安全研究人员提交的高危安全审计报告** (#2760, #2761, #2762)，涉及文件泄露和权限绕过等严重问题，这可能是近期的核心修复方向。与此同时，**核心架构层有重大推进**，两个关于 Codex 新架构和 Provider 选择机制的 PR (#2756, #2757) 已合并，标志着项目在扩展性与灵活性上迈出了关键一步。基础设施方面，预算耗尽时 LLM 调用静默丢弃的问题终于有了对应的修复 PR (#2759)，有望改善用户体验。**项目健康度评分：7.5/10**（活动性高，但存在高风险漏洞待修复）。

### 2. 版本发布

**无。** 过去 24 小时内未有新版本发布。

### 3. 项目进展

今日合并/关闭了 5 个 PR，构成了项目的重要基础能力推进：

- **核心架构升级**：合并了 `omri-maya` 提交的两个重量级 PR。
    - **PR #2756** (Closed): 引入了操作员驱动的 Provider 选择、切换和记忆迁移功能。这为未来接入更多 AI 服务提供了统一的“接缝（seams）”，奠定了插件化的 Provider 生态基础。
    - **PR #2757** (Closed): 实现了 Codex 的全新 v2 载荷，将其作为一个完整的 Agent Provider，通过 OneCLI 进行纯金库（vault-only）认证。这标志着 Codex 从试验性接入走向正式化和架构一体化。

- **基础设施自动化**：合并了 `gavrielc` 的 **PR #2758** (Closed)，将容器中全局 CLI 工具的安装方式从硬编码 `Dockerfile` 改为数据驱动的 `cli-tools.json` 清单。这极大简化了新 CLI 工具的添加流程，提升了容器化部署的灵活性和可维护性。

- **文档修复**：合并了 `glifocat` 的 **PR #2764** (Closed)，修复了 `CLAUDE.md` 指引中指向已迁移文件的路径错误，修正了开发文档与实际情况的偏差，避免了开发者的困惑。

- **文档修复**：合并了 `Koshkoshinsk` 的 **PR #2769** (Closed)，修复了 `/add-codex` 技能的文档，增加了交互式认证和主机重启步骤的说明，降低了用户上手 Codex 的门槛。

### 4. 社区热点

本日社区讨论的核心集中在 **安全审计报告** 上。由用户 `YLChen-007` 在同一天提交的三个严重安全 Issue (#2760, #2761, #2762) 成为焦点。这三份报告均有详细的分析和利用路径描述，显露出社区安全专家对项目安全性的深度关注。虽然当前评论数为 0，但这类议题通常会迅速引发核心维护者的讨论。这三份报告共同反映了社区对 **Agent 自主行为安全边界** 的核心诉求，即如何防止 Agent 在获得授权后，绕过用户审查，执行恶意操作。

- **相关链接**：
    - #2760 - 文件泄露高危漏洞
    - #2761 - 无认证Webhook权限绕过
    - #2762 - MCP服务器审批隐藏参数风险

### 5. Bug 与稳定性

今日报告的 Bug 呈现“少而精”的特点，但严重性极高。此外，也包含了对性能优化和文档错误的修复。

**严重**

1.  **安全问题 - 任意文件泄露 (Critical)** [#2760](https://nanocoai/nanoclaw/issues/2760): 报告称 `send_file` 工具接受绝对路径，可以将宿主机任意文件复制并发送出去。**暂无关联修复 PR**。
2.  **安全问题 - 本地网关权限绕过 (Critical)** [#2761](https://nanocoai/nanoclaw/issues/2761): 报告称本地环回 Webhook 未认证发送者，导致攻击者可以伪造事件绕过审批流程。**暂无关联修复 PR**。
3.  **安全问题 - MCP服务器审批隐藏参数 (High)** [#2762](https://nanocoai/nanoclaw/issues/2762): 报告称在 `add_mcp_server` 的审批流程中，用户可以看不到实际被提交的 `args` 和 `env` 参数，存在被恶意 Agent 利用的风险。**暂无关联修复 PR**。

**一般**

4.  **预算耗尽静默丢弃 (Medium)** [#2751](https://nanocoai/nanoclaw/issues/2751): 当 LLM 调用因预算耗尽而失败时，用户不会收到任何响应，即 Agent 表现为静默失败。**已有修复 PR #2759** 待合并。
5.  **Telegram 标记兼容性问题 (Low)** [#2767](https://nanocoai/nanoclaw/issues/2767): `channels` 分支上的 `telegram-markdown-sanitize.ts` 用于解决旧版 Telegram 适配器不支持 `MarkdownV2` 的问题，而新版本的适配器已原生支持，此兼容层已过时。
6.  **Claude Provider 性能问题 (Low)** [#2768](https://nanocoai/nanoclaw/issues/2768): 报告指出当前 Claude Provider 未启用 `enablePromptCaching`，导致每次交互都重新发送完整的系统提示，浪费 token 且增加延迟。

### 6. 功能请求与路线图信号

- **核心能力信号**：已合并的 **PR #2756 (Provider Seams)** 和 **PR #2757 (Codex v2)** 强烈暗示了项目路线图的下一个阶段：**构建多 AI Provider 生态**。未来的 Agent 将不再局限于单一模型，用户可以自由切换和升级后端。
- **用户体验增强**：**Issue #2751** 和对应的修复 **PR #2759** 直接回应用户对 Agent **可靠性** 的不满。修复预算耗尽时的静默失败，是提升用户信任度的关键举措。
- **代码质量与文化**：`amit-shafnir` 提交的两个带 `[follows-guidelines]` 标签的 PR (#2765, #2766) 旨在为 Providers 和 Channels 添加 `.format-lint-off` 功能。这可能是社区成员在持续改进代码质量流程，自动处理格式化/ lint 相关的琐碎问题。
- **性能优化**：**Issue #2768** 提出的“开启 Prompt Caching”是一个明确的技术优化请求，若能实现，将对所有使用 Claude Provider 的用户带来成本与速度上的直接好处。

### 7. 用户反馈摘要

从今日的 Issue 中，可以清晰看到几类用户画像和痛点：

- **安全研究人员** (`YLChen-007`): 提交了非常专业、结构化的安全问题，反映出对 AI Agent **安全边界** 的强烈担忧，并希望项目方能堵上所有潜在的后门。他们的诉求是“Agent 能力再强，也不能越界”。
- **开发者/高级用户** (`assapin`, `galmorduku`, `chiptoe-svg`): 反馈集中在技术上难以快速定位或理解的问题。例如预算耗尽静默 (#2751)，他们能通过技术手段找到原因，但对“无反馈”这种 UX 问题感到不满。他们希望系统行为透明、可调试 (`enablePromptCaching`)。
- **文档读者** (`glifocat`): 反映了文档自动化与代码实际状态不同步的痛点。对于依赖 `CLAUDE.md` 引导的开发流程来说，一个错误的路径引用可能导致流程中断，他寻求的是一种“文档即代码”的精确性。

### 8. 待处理积压

以下议题或 PR 未得到及时响应或解决，提醒项目维护者关注：

1.  **安全审计报告积压 (Critical)**：[#2760](https://nanocoai/nanoclaw/issues/2760)、[#2761](https://nanocoai/nanoclaw/issues/2761)、[#2762](https://nanocoai/nanoclaw/issues/2762)。 这三个是今日提交的高危害安全漏洞，**应作为最高优先级处理**，安排安全评审和修复计划。
2.  **主机健康审计修复 PR (Pending)**：[PR #2732](https://nanocoai/nanoclaw/pull/2732)。 `caburi00` 提交的针对安全审计结果的修复 PR，已打开 4 天且仍在等待审查。如果其修复内容与上述新报告的安全问题有重叠，应优先合并。
3.  **预算错误修复 PR (Pending)**：[PR #2759](https://nanocoai/nanoclaw/pull/2759)。 修复“预算耗尽时静默丢弃”问题的 PR，直接关联用户体验。由于该 Issue 于 4 天前提出，PR 也已存在，建议尽快审查合并。
4.  **Codex 文件事件交付修复 PR (Pending)**：[PR #2770](https://nanocoai/nanoclaw/pull/2770)。 `Koshkoshinsk` 提交的修复 Codex 图像生成事件从未到达聊天的 PR。这是 Codex 新架构 (`v2`) 合并后的一个功能完整性修复，应尽快验证并合并。

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw 项目日报 — 2026-06-15

## 1️⃣ 今日速览

- 过去 24 小时内项目仅产生 1 条 Issue（#955，状态 OPEN），无任何 Pull Request 或新版本发布，整体活跃度较低。
- 该 Issue 为增强功能请求，关注 Azure OpenAI LLM Provider 的基于身份认证支持，尚未引起讨论或点赞。
- 无 Bug 报告、无代码合并、无版本迭代，项目处于相对静默期，但社区仍有关注点落在云服务认证集成上。

## 2️⃣ 版本发布

_无新版本发布。_

## 3️⃣ 项目进展

- 今日无 Pull Request 被合并或关闭，无代码提交进入主分支。
- 项目功能推进暂停，整体里程碑与开发进度未发生可量化的前移。

## 4️⃣ 社区热点

- **唯一活动的 Issue**  
  [#955 [enhancement] Identity based authentication support for Azure OpenAI LLM Provider](https://github.com/nullclaw/nullclaw/issues/955)  
  - 作者：kunalk16 | 创建/更新：2026-06-15 | 评论：0 | 👍：0  
  - 诉求：为 Azure OpenAI LLM Provider 添加基于身份认证支持，建议利用 `DefaultTokenCredential` 自动使用 `az CLI` 登录凭据。
  - 背景分析：用户因 Azure 订阅的安全策略无法使用 API Key，需通过托管身份或开发者凭据认证。此需求反映了企业级 Azure 部署中对安全合规认证方式的刚性需求。

## 5️⃣ Bug 与稳定性

- 今日未报告任何 Bug、崩溃或回归问题。
- 项目代码稳定性无明显负面信号。

## 6️⃣ 功能请求与路线图信号

- **唯一功能请求**  
  [#955](https://github.com/nullclaw/nullclaw/issues/955) 明确要求为 Azure OpenAI 提供身份认证支持（非 API Key）。  
  - 技术实现方向：集成 `Azure.Identity` 库与 `DefaultTokenCredential`。  
  - 路线图信号：若被采纳，将增强项目在 Azure 企业环境中的可用性，并可能影响后续多 Provider 统一认证架构的设计。暂无关联 PR，需等待维护团队评估。

## 7️⃣ 用户反馈摘要

- 今日无任何用户评论或讨论（#955 评论数为 0）。  
- 无法从现有数据中提炼真实使用痛点或满意度信息。社区互动稀少，项目需要在后续鼓励用户参与反馈。

## 8️⃣ 待处理积压

- **当前唯一待处理 Issue**  
  [#955 Identity based authentication for Azure OpenAI](https://github.com/nullclaw/nullclaw/issues/955)  
  - 创建于当日，尚无维护者回应。虽然不属于“长期未响应”，但作为唯一活跃问题，建议维护团队尽早回复以明确采纳意向或提供替代方案。

- **长期未响应的历史 Issues/PRs**  
  - 根据提供的数据不足 24 小时，无法识别积压项。项目可定期检查更早的未关闭 Issue，防止被忽略的需求积累。

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw 项目日报 — 2026-06-15

---

## 1. 今日速览

今日 IronClaw 项目活跃度极高。过去 24 小时内，仓库共产生 **22 条 Issue 更新**（新开 17 条，关闭 5 条）和 **44 条 PR 更新**（待合并 29 条，已合并/关闭 15 条），反映出核心团队与社区在功能开发、Bug 修复及工程效率提升方面均投入了密集的精力。值得注意的亮点是，今日社区与核心贡献者共同提交了多项关键修复（如认证恢复、MCP 工具审批、WebSocket 冲突等），同时产出了多个旨在将 IronClaw 本身“AI 原生”化的内部工程改进议题（#4878 及系列子任务），表明项目正在由单纯提供工具向“工具自我进化”阶段演进。


## 2. 版本发布

**无**（过去 24 小时内无官方版本发布）


## 3. 项目进展

以下是今日**已合并/关闭**的关键 PR，覆盖了稳定性修复、测试补充、以及 UX 改进：

- **[PR #4873] (已关闭)** — 测试(slack): 重新安置 approval→auth→final-reply 的 Slack 端到端测试 (修复 #4847)
  - 将此前因引入时即存在问题而被删除的 Slack 审批流程端到端测试重新归位，确保了 Slack 集成场景下的认证与回复链路完整可靠。
  - 链接: https://github.com/nearai/ironclaw/pull/4873

- **[PR #4738] (已关闭)** — 功能(reborn): 附件 Web 上传 UX 在 WebChat v2 SPA 中上线 (#4644 后续)
  - 实现了 WebChat v2 的前端附件上传、暂存、重命名、删除等完整交互，使影像附件可真正发送给视觉模型，而不仅是作为文本引用。
  - 链接: https://github.com/nearai/ironclaw/pull/4738

- **[PR #4848] (已关闭)** — 功能(auth-resume): 改用每次调用唯一标识 (`input_ref`) 匹配待恢复的认证，而非仅匹配 `capability_id`
  - 解决了 `take_if` 闭包在相似能力调用间可能错误匹配 pending auth-resume 的潜在 bug，提升了认证恢复的精确性。
  - 链接: https://github.com/nearai/ironclaw/pull/4848

- **[PR #4851] (已关闭)** — 重构: 将 trusted-trigger 的“来源”从字符串类型提升为正式的类型系统（Option 7）
  - 消除了通过字符串“漂白”触发来源的安全隐患，彻底以类型系统保证调度链中的信任传递。
  - 链接: https://github.com/nearai/ironclaw/pull/4851


## 4. 社区热点

今日讨论最活跃、关注度最高的议题集中在 **认证/审批流程的透明度和准确度** 以及 **安全上下文与用户引导** 方面。

- **[Issue #4887]** [Reborn] 基于 Provider 的 MCP 工具审批恢复可能因过期的功能引用而失败 (作者: krishna-505)
  - **背景**: 用户在 WebUI v2 批准了某项 Provider 背书的 MCP 能力请求后，恢复执行时会因“能力输入引用不在本次循环范围”而直接报错（例如使用 `nearai.web_search` 工具时）。
  - **诉求**: 用户期望审批与执行的衔接应更加可靠，不应在批准后发生不可恢复的运行时错误。
  - 链接: https://github.com/nearai/ironclaw/issues/4887

- **[Issue #4872]** 将运行时上下文中的外部通信标签渲染为转义的非可信数据，而非指令文本 (作者: henrypark133)
  - **背景**: 运行时通信标签（如外部交付目标、通道标签）当前以 `model_safe_label` 方式映射后直接嵌入模型可见的系统提示中，可能被模型误解释为可执行的指令或工具调用。
  - **诉求**: 社区与贡献者普遍认同应将此类标签完全视为非可信数据并严格转义/截断，杜绝任何通过外部标签注入指令的潜在风险。
  - 链接: https://github.com/nearai/ironclaw/issues/4872

- **[Issue #4884]** [Reborn] Google Calendar 认证提示错误：要求提供 Access Token 而非引导用户完成 OAuth 流程 (作者: sunglow666)
  - **背景**: 用户安装并激活 Google Calendar 扩展后，在聊天中要求查看日程，系统直接弹出一个包含混乱语法或不完整权限的 Token 输入框，未引导用户进行标准 OAuth 授权。
  - **诉求**: 用户强烈期望获得流畅的 OAuth 引导体验，而非暴露底层凭证细节。这反映了扩展生态中认证流程的简洁性与用户体验之间存在较大落差。
  - 链接: https://github.com/nearai/ironclaw/issues/4884


## 5. Bug 与稳定性

按严重程度排列如下：

| 严重程度 | Issue/PR | 摘要 | 状态 |
|---|---|---|---|
| **高** | [#4887](https://github.com/nearai/ironclaw/issues/4887) | MCP 工具审批恢复因 `input_ref` 不在当前循环范围而失败 | 今日新开，尚未修复 |
| **高** | [#4874](https://github.com/nearai/ironclaw/issues/4874) | WebChat v2 在非 localhost 的纯 HTTP 访问下发送消息报 `TypeError: Illegal invocation` | 今日新开，尚未修复 |
| **中** | [#4870](https://github.com/nearai/ironclaw/issues/4870) | WebUI WebSocket 助手与 v2 认证合约冲突（用 `?token=` 认证但路由明确拒绝查询参数认证） | 今日新开，尚未修复 |
| **中** | [#4852](https://github.com/nearai/ironclaw/issues/4852) | [Reborn] Shell 命令在审批对话框或活动历史中不可见 | 今日新开，尚未修复 |
| **中** | [#4751](https://github.com/nearai/ironclaw/issues/4751) | [Reborn] 大响应请求因 Provider 工具参数超过 16384 字节而失败 | 已关闭，但尚未见正式修复 PR |
| **低** | [#4886](https://github.com/nearai/ironclaw/issues/4886) | [Reborn] 安装扩展后“AUTH NEEDED”状态提示不清晰，缺乏下一步引导 | 今日新开，尚未修复 |


## 6. 功能请求与路线图信号

今日新提出的功能请求中，最具路线图指示意义的是：

- **[Issue #4882] “构建编码代理云工作流”** — 核心作者 think-in-universe 提出允许通过分配 Issue 或提及 Bot 的方式，让云端编码代理自动完成 PR。这是项目向 AI 原生工程团队演进的关键节点。
  - 链接: https://github.com/nearai/ironclaw/issues/4882

- **[Issue #4880] “自动化代码评审与审查评论解决”** — 提出将 AI 审查自动化到新高度，包括初始 AI Review、处理审查意见、最终准备合并等环节。结合已有的 [PR #3511](https://github.com/nearai/ironclaw/pull/3511)（仍在更新，针对工具 Schema 删除重复查找），可见 AI 内生的工程自动化正在成为长期主线。

- **[Issue #4875] “拆分 runtime_context.rs——将提示渲染器与数据模型+获取合约分离”** — 一项清晰的重构请求，反映了核心代码模块因功能增长（已达 1025 行）导致可维护性下降，是项目架构演进的健康信号。
  - 链接: https://github.com/nearai/ironclaw/issues/4875


## 7. 用户反馈摘要

来自今日 Issue 评论的真实用户痛点及使用场景：

- **“AUTH NEEDED” 状态不明确** (#4886): 用户安装扩展后仅看到“AUTH NEEDED”，但不清楚下一步是“点击配置”、“完成授权”，还是“需要哪些权限”。用户期望获得更明确的分步引导与状态指示。
- **Google Calendar 认证体验脱节** (#4884): 用户期望看到完整的 OAuth 授权界面，而非直接弹出 “请提供 Access Token”。当前实现暴露了底层实现细节，对非技术用户极不友好。
- **Shell 命令审批透明性差** (#4852): 用户发送 shell 命令（如 `pwd`、`ls -la`）后，审批对话框仅显示“Capability: builtin.shell”，未展示具体要执行的命令内容；活动历史中也无记录，导致用户无法确认实际执行了什么操作。
- **非 localhost 的纯 HTTP 访问导致 WebSocket 失效** (#4874): 用户通过内网 IP 或主机名访问 WebUI 时遭遇 `TypeError: Illegal invocation`，被迫只能在开发机上使用，严重限制了对局域网的分享部署场景。


## 8. 待处理积压

以下为长期未合并/响应的关键 PR 与 Issue，建议维护者关注：

| 项目 | 摘要 | 状态 | 提示 |
|---|---|---|---|
| [PR #3708](https://github.com/nearai/ironclaw/pull/3708) | `ironclaw` 版本发布 PR（含多项 API 破坏性变更） | 自 5月16日起待合并，持续更新 | 阻塞其他依赖更新的引入，可能已滞后于当前分支 |
| [PR #4002](https://github.com/nearai/ironclaw/pull/4002) | 批量更新 GitHub Actions 依赖 (16 个更新) | 自 5月24日起待合并 | 可能导致 CI 与当前分支工具链不兼容 |
| [PR #4499](https://github.com/nearai/ironclaw/pull/4499) | 批量更新 tokio 生态依赖（3 个更新） | 自 6月5日起待合并 | 安全补丁与性能改进长期搁置 |
| [Issue #4692](https://github.com/nearai/ironclaw/issues/4692) | “IronClaw Reborn 本地 Dogfooding 发现” 6/08-6/14 周报 | 已于 6/14 更新 | 历史性汇总，反映启动、配置等多方面遗留问题，建议拆分为具体子 Issue |

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI 项目动态日报 | 2026-06-15

---

## 1. 今日速览

过去24小时内，LobsterAI 项目保持中等活跃度：共处理 **2 条 Issue**（均为更新，无新开）、**4 条 PR**（其中 1 条被合并关闭，3 条仍待合并）。项目主要围绕 **定时任务数据残留修复** 以及 **CoWork 会话体验增强** 展开。值得注意的是，当前 Issue 和 PR 均处于“stale”标记状态，多数由同一位贡献者在一个多月前创建，近期才被重新触发更新，社区参与节奏有所放缓。

---

## 2. 版本发布

**无**。过去24小时内未发布新版本。

---

## 3. 项目进展

### 已合并/关闭的 PR

- **[#1465] fix(scheduled-tasks): 已删除的定时任务重启后作为幽灵会话重新出现**  
  作者：linlihua | 创建：2026-04-04 | 关闭：2026-06-14  
  **摘要**：修复了定时任务删除后数据残留的 Bug。根本原因在于删除流程只清除了网关侧（OpenClaw）的任务记录，而未清理本地 SQLite `cowork_sessions` 表中的关联会话记录。重启后这些残留会话会以空内容幽灵形式重新出现。修改后同步清理本地会话数据。  
  **影响**：提升了定时任务功能的可靠性，杜绝反复删除无效会话的用户困扰。  
  链接：https://github.com/netease-youdao/LobsterAI/pull/1465

---

## 4. 社区热点

今日无高热度讨论出现。所有 Issue 和 PR 的评论数均为 0 或 1，无新的 👍 反应。当前最受关注的仍是两条长期存在的 UI 相关问题，反映了 **语言国际化一致性** 和 **表单输入友好性** 两个基础体验点的缺失。

- **#1434**：语言设为中文时，Agent 技能页搜索无数据显示英文提示及按钮（👍 0，评论 1）  
  https://github.com/netease-youdao/LobsterAI/issues/1434

- **#1435**：新建自定义 Agent 名称过长超出弹框，展示不友好（👍 0，评论 1）  
  https://github.com/netease-youdao/LobsterAI/issues/1435

---

## 5. Bug 与稳定性

| 严重程度 | Issue / PR | 描述 | 是否已有修复 PR |
|----------|------------|------|----------------|
| 中等 | #1434 | 语言切换为中文后，搜索无数据时仍展示英文提示和按钮，本地化不完整 | 无 |
| 中等 | #1435 | 自定义 Agent 名称过长时弹框溢出，UI 截断不友好 | 无 |
| 严重（已修复） | #1465 | 定时任务删除后本地会话残留，重启后幽灵会话反复出现 | 已关闭（PR #1465） |

其余 3 条 OPEN PR 均属功能增强，未报告新 Bug。

---

## 6. 功能请求与路线图信号

以下 3 个 OPEN PR 均处于待合并状态，且均为功能增强，可能被纳入下一版本：

- **[#1429] feat(cowork): 添加会话内消息搜索功能**  
  使用 `mark.js` 实现实时高亮、`Cmd/Ctrl+F` 快捷搜索、前后跳转匹配项。  
  路径：https://github.com/netease-youdao/LobsterAI/pull/1429

- **[#1430] feat(cowork): 会话运行期间自动阻止系统休眠**  
  利用 Electron `powerSaveBlocker`，防止长时间任务被系统挂起中断。  
  路径：https://github.com/netease-youdao/LobsterAI/pull/1430

- **[#1431] feat(cowork): StreamingActivityBar 右侧显示会话运行计时器**  
  实时显示已运行时间（如 `已运行 2m 5s`），提升进度感知。  
  路径：https://github.com/netease-youdao/LobsterAI/pull/1431

**路线图信号**：上述 PR 均围绕 CoWork 会话体验优化（搜索、防休眠、计时器），表明项目当前重点向 **会话交互增强** 和 **可靠性提升** 倾斜。

---

## 7. 用户反馈摘要

从 Issue 评论和描述中提炼出的主要用户痛点：

- **本地化不完整**：当系统语言设为中文时，部分 UI 元素（搜索空状态提示、按钮）仍显示英文，破坏一致性预期（#1434）。
- **表单边界处理缺失**：Agent 名称允许输入过长字符，但弹框无溢出限制，导致视觉混乱（#1435）。用户需要更清晰的输入长度提示或截断方案。
- **定时任务删除困惑**：用户反复删除同一任务却重启后重现，体验受挫（#1465 已修复）。

---

## 8. 待处理积压

以下 Issue / PR 已超过两个月未获维护者响应或合并，建议优先审阅：

| 编号 | 标题 | 状态 | 创建时间 | 最后更新 |
|------|------|------|----------|----------|
| #1434 | 中文语言下搜索空状态英文显示 | OPEN | 2026-04-03 | 2026-06-14 |
| #1435 | 新建 Agent 名称过长超出弹框 | OPEN | 2026-04-03 | 2026-06-14 |
| #1429 | CoWork 消息搜索功能 | OPEN | 2026-04-03 | 2026-06-14 |
| #1430 | CoWork 自动防休眠 | OPEN | 2026-04-03 | 2026-06-14 |
| #1431 | CoWork 运行计时器 | OPEN | 2026-04-03 | 2026-06-14 |

上述 Issue 和 PR 均带有 `[stale]` 标签，提示维护者关注。尤其是 **#1434** 和 **#1435** 两个 UI Bug 直接影响新用户首次体验，建议优先分配资源处理。

---

*数据来源：GitHub 仓库 netease-youdao/LobsterAI，统计周期 2026-06-14 00:00 – 2026-06-15 00:00 UTC。*

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis 项目动态日报 (2026-06-15)

## 1. 今日速览
过去24小时项目整体保持稳定，活跃度处于较低水平。社区提交了1个新的增强功能请求（Issue #1123），但未合并任何Pull Request或发布新版本。项目维护者响应及时，但贡献者在本周期内未提交代码变更，表明当前阶段更多聚焦于需求收集与讨论。建议社区关注该新Feature请求的后续讨论进展。

## 2. 版本发布
无

## 3. 项目进展
今日无合并或关闭的Pull Request，项目代码库没有新增代码变更。整体进展缓慢，建议维护团队关注社区提出的新方向，并适时推进PR的创建与合并。

## 4. 社区热点
唯一活跃的Issue为 **#1123**（[链接](https://github.com/moltis-org/moltis/issues/1123)），用户 `joeblew999` 提议为Moltis添加纯Rust实现的 **turbovec** 作为可选的内存后端，主要用于极端边缘压缩场景。该Issue目前0评论、0点赞，但因其独特性（纯Rust、边缘压缩、内存后端）可能吸引后续讨论。用户背后的核心诉求是在资源受限或需要极致数据压缩的环境中，提供不依赖外部C库的轻量级内存管理方案，以提升Moltis在IoT、嵌入式等领域的适用性。

## 5. Bug 与稳定性
今日无新的Bug、崩溃或回归问题报告。项目当前稳定性未见明显风险。

## 6. 功能请求与路线图信号
- **新功能请求**：Issue #1123 提出引入纯Rust的 `turbovec` 内存后端。该需求属于性能与资源优化方向，与现有后端（如标准Vec、外部库）形成互补。鉴于Moltis可能定位为高性能向量存储或张量计算库，此类后端可增强项目在极端条件下的竞争力。目前尚无关联PR，但若社区反响热烈，有望进入下一版本路线图。
- **路线图线索**：无其他明确信号。建议维护团队评估 `turbovec` 的成熟度与集成成本，以决定是否将其纳入短期规划。

## 7. 用户反馈摘要
今日Issue #1123 尚未收到评论，暂无可直接提取的用户痛点或使用场景反馈。但从Feature描述推测，用户可能面临现有内存后端在压缩率或依赖管理上的不足，倾向于纯Rust解决方案以降低交叉编译复杂度。

## 8. 待处理积压
今日无新增长期未响应的Issue或PR。项目当前积压状况良好，维护者对新请求的响应时间（创建于2026-06-14，当日即出现在日报中）表明管理闭环较为及时。建议持续关注#1123的社区参与度，及时回复并引导转化为实际开发任务。

---

*数据来源：GitHub 仓库 moltis-org/moltis，采集时间 2026-06-15 00:00 UTC。*

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

好的，这是为您生成的 CoPaw 项目动态日报。

---

# CoPaw 项目动态日报 | 2026-06-15

## 今日速览

今日项目整体活跃度**极高**，社区反馈和贡献者提交均十分踊跃。过去24小时内，共产生21条Issue更新和16条PR更新，反映出项目正处于快速迭代和社区深度使用阶段。其中，关于桌面端性能、本地模型兼容性以及插件生态的讨论尤为集中。尽管无新版本发布，但多项关键Bug修复和新功能PR正在推进，项目健康度良好，社区参与度持续上升。

## 版本发布

无新版本发布。

## 项目进展

今日共有5个PR被合并/关闭，修复了多个重要问题，并完成了一项关键的功能回滚。

- **关键问题修复：**
    - **PR #5051** (已合并): **修复了Windows桌面版重启后智能体及对话丢失的问题**。此PR解决了长期困扰用户的Issue #4733，通过持久化后端端口，确保桌面应用重启后能正确恢复上次关闭时的状态（智能体及对话session）。
    - **PR #5035** (已合并): **修复了llama.cpp本地模型版本号解析的潜在Bug**。当llama.cpp的构建版本号超过4位数时，原有解析逻辑会出错，该PR采用了更健壮的解析方式。
    - **PR #5038** (已合并): **修复了当消息列表为空时引发的索引错误**。通过增加空列表的守卫判断，增强了`LightContextManager`在极端情况下的稳定性。

- **功能回滚与调整：**
    - **PR #5092** (已合并): 回滚了先前的一个打包修复（#5084），以解决该修复引入的新问题（#5086），体现了项目组对稳定性的快速响应。

- **重要功能推进：**
    - **PR #5189** (待合并): 引入了**前端插件的斜杠命令自动补全功能**。允许插件注册`/command`令牌，用户在聊天输入框输入`/`时即可弹出提示，大幅提升插件的可发现性和交互体验。

## 社区热点

- **Issue #5156: [Feature] 建议支持 kimi-for-coding / 加入 uv 白名单**
  - **热度：** 5条评论 | 社区讨论热烈
  - **诉求分析：** 用户已订阅Kimi的coding套餐，希望能在CoPaw中直接使用，避免重复付费。这反映出用户对**多模型API兼容性**以及**现有付费订阅权益复用**的强烈诉求。许多用户希望一个统一的客户端能接入他们已订阅的各类服务，而非被限制在特定API上。([链接](https://github.com/agentscope-ai/QwenPaw/issues/5156))

- **Issue #5047: [Question] Windows Tauri 桌面端启动特别慢**
  - **热度：** 5条评论 | 持续关注
  - **诉求分析：** 用户反映从Python打包切换到Tauri后，启动时间从一两分钟暴增至十几分钟，严重影响日常使用。这表明**桌面端重写带来的性能回归**是当前用户最核心的痛点之一。社区正在积极讨论并提供诊断信息，期待官方能优先解决此启动性能问题。([链接](https://github.com/agentscope-ai/QwenPaw/issues/5047))

## Bug 与稳定性

今日报告的Bug较多，主要集中在**v1.1.11.post2**版本，显示出新版本引入了一些回归问题。

- **严重程度：高（已确认回归）**
    - **Issue #5163: Gemini 工具调用在 v1.1.11.post2 中失效**。用户在`v1.1.10`版本中可正常使用，升级后出现问题，这是一个明确的**回归Bug**，需优先排查。([链接](https://github.com/agentscope-ai/QwenPaw/issues/5163))
    - **Issue #5184: 本地模型提供商在 v1.1.11.post2 中显示异常**。本地创建的模型提供商无法在设置界面正确显示，影响新版本功能的正常使用。([链接](https://github.com/agentscope-ai/QwenPaw/issues/5184))

- **严重程度：中**
    - **Issue #5190: 企业微信私聊访问控制审批界面不可见**。功能开启后，用户虽然收到“需审批”的提示，但找不到审批入口，导致流程中断。([链接](https://github.com/agentscope-ai/QwenPaw/issues/5190))
    - **Issue #5181: 插件依赖安装失败导致CMD窗口持续弹窗**。当网络环境不佳时，后端的pip安装失败会触发死循环，并不断弹出可见的CMD窗口，严重影响桌面端用户体验。([链接](https://github.com/agentscope-ai/QwenPaw/issues/5181))
    - **Issue #5171: 上下文压缩导致任务中断**。在人设文件token较大时，上下文压缩可能将内容完全清除，导致模型丢失上下文，任务无法继续。([链接](https://github.com/agentscope-ai/QwenPaw/issues/5171))

- **严重程度：低**
    - **Issue #5177: 钉钉Channel的会话在Console前端不可见**。会话虽能正常保存到文件，但在Web控制台的列表中无法显示，存在数据展示的Bug。([链接](https://github.com/agentscope-ai/QwenPaw/issues/5177))
    - **Issue #5161: 长对话后无响应**。长时间对话后，应用可能卡死。([链接](https://github.com/agentscope-ai/QwenPaw/issues/5161))

## 功能请求与路线图信号

今日社区提出的新功能需求丰富，部分已有相关PR跟进，显示出社区对项目未来发展的多样化期望。

- **高优先级信号（已有PR跟进）：**
    - **Issue #5168: 添加Zalo Bot官方渠道支持**。来自越南社区用户的呼声，考虑到Zalo在越南的市场份额，此请求暗示了项目**国际化渠道拓展**的需求。PR #5186（添加越南语界面）也侧面印证了越南社区的高活跃度。([链接](https://github.com/agentscope-ai/QwenPaw/issues/5168))
    - **Issue #5169 (由PR #5186/5175关闭): 添加越南语界面支持**。已由社区贡献者提交了两个不同的PR进行实现，其中PR #5186提供完整翻译，预计将被优先合并。([PR链接](https://github.com/agentscope-ai/QwenPaw/pull/5186))

- **中等优先级信号（社区讨论中）：**
    - **Issue #5156: 支持kimi-for-coding及白名单**。如前所述，反映了用户对于**多供应商API接入**的强烈需求。
    - **Issue #5165: 打包安装后白屏**。社区贡献者指出了`pack-tauri`打包脚本存在模块引用错误，修复此问题对于**降低开发者和用户的使用门槛**至关重要。

- **低优先级信号（未来优化方向）：**
    - **Issue #5185: 为Agent上下文添加实时时间戳**。用户希望Agent能感知到具体的时分秒信息，而不仅仅是日期，以提升 Agent 对时间敏感任务的执行能力。
    - **Issue #5182: 统一模型配置功能**。提出统一向量、文本、音视频等不同模型类型的配置界面，提升易用性。
    - **Issue #5187: Windows桌面GUI自动化**。新增的`computer_use`工具（PR已提交）允许Agent在Windows桌面上执行截图、点击等操作，是一项非常前沿且强大的自动化功能，可能成为未来版本的一大亮点。

## 用户反馈摘要

- **痛点反馈：**
  - **桌面端性能问题**：用户对Windows版Tauri启动慢（#5047）和插件依赖安装弹窗（#5181）表达了强烈不满，称其**严重影响日常使用**。
  - **渠道整合问题**：企业微信（#5190）和钉钉（#5177）的用户都遇到了功能不可用或数据不同步的问题，影响了多平台协作体验。
  - **回归与破坏性变更**：用户对v1.1.11.post2版本出现的Gemini工具调用（#5163）和本地模型（#5184）的回归问题感到困惑和失望。

- **使用场景与建议：**
  - **付费服务复用**：用户明确希望CoPaw能成为一个**“万能客户端”**，将已订阅的Kimi、通义千问等不同服务统一接入使用。
  - **人设与上下文管理**：用户在使用复杂的Agent人设文件时，发现了上下文压缩机制的缺陷（#5171），希望官方能提供更精细的控制，如按条数保留或排除特定文件。
  - **Agent自动化能力**：用户对Cron/心跳Agent的执行能力提出质疑（#5174），认为其无法可靠执行“写文件”等重任务，相关PR #5180 正试图通过增加超时时间和优化上下文提示来解决此问题。

## 待处理积压

- **PR #4622 (DataPaw 数据分析插件)**：该PR自5月22日提出，旨在添加一个包含12项BI技能的数据分析插件。目前仍处于未合并状态，考虑到社区对插件生态的期待，建议维护者评估其成熟度并决定是否纳入主线。([链接](https://github.com/agentscope-ai/QwenPaw/pull/4622))

- **Issue #5009 (可观测性集成)**：用户询问CoPaw是否计划集成Langfuse、OpenTelemetry等LLM可观测性平台。该需求对于企业级用户至关重要，虽已关闭，但建议在官方路线图中给予回应。([链接](https://github.com/agentscope-ai/QwenPaw/issues/5009))

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，请看以下为您生成的 ZeroClaw 项目动态日报。

---

# ZeroClaw 项目动态日报 — 2026-06-15

## 1. 今日速览

ZeroClaw 项目今日保持**高活跃度**，主要体现为 PR 提交量激增（50条）。值得注意的是，**待合并 PR 堆积严重（46条）**，社区贡献者提交了大量修复与功能增强，但合并速度（4条）相对滞后，可能意味着项目维护者正在集中精力进行代码审查或核心重构。Issues 方面，一个关于**封闭执行模式**的高风险 RFC (#6293) 讨论活跃，反映了社区对安全架构的深度关注。整体来看，项目正在进行大量的技术债务清理与功能补全，但合并瓶颈值得关注。

## 2. 版本发布

无新版本发布。

## 3. 项目进展

今日有 4 个 PR 被合并/关闭，同时 2 个高严重性 Issue 得到解决，项目在关键问题修复和核心功能演进上取得进展。

-   **关键 Bug 修复**：
    -   **Windows 编译修复**：`[Bug]: AMQP channel cannot compile on Windows` (#7452) 今日关闭，解决了 AMQP 通道在 Windows 平台因 tokio-reactor-trait 的 Unix-only 实现导致的编译失败问题。
    -   **配置安全修复**：`config: extend #[secret]...` (#6989) 今日关闭，扩展了 `#[secret]` 宏的功能，以支持对 Header 中的 bearer token 进行脱敏处理，增强了配置安全性。
-   **重要功能/修复合并**：
    -   `fix(issue): resolve #7542 [Bug]: ask_user fails instantly...` (#7664) 被合并，修复了 gateway web 界面中 `ask_user` 功能因“Channel closed”错误而立即失败的 Bug。
    -   `feat(cron): add a pause/resume toggle to scheduled tasks` (#7384) 被合并，为定时任务仪表盘增加了暂停/恢复功能，完善了任务管理体验。
-   **活跃开发中的关键演进**：
    -   **核心API重构**: `refactor(api): make before_llm_call take mutable borrows` (#7667) 提交，旨在通过将 `before_llm_call` 钩子函数的参数改为可变引用，避免在每次 LLM 请求时对对话历史进行昂贵的克隆，是提升运行性能的关键一步。
    -   **功能增强**: `feat(gateway/cron): accept enabled on CronPatchBody for pause/resume` (#7666) 提交，通过 HTTP API 实现了定时任务的暂停/恢复，与 #7384 的合并形成完整闭环。

**项目健康度评估**：项目当前正处于积极解决问题的“清理期”和性能优化的“关键期”。社区贡献与核心维护并行，但合并效率亟待提升，以避免贡献者流失和 PR 大量腐烂。

## 4. 社区热点

-   **最热 Issue**：[RFC: Air-gapped execution mode with companion daemon over unix socket (enclave support)](https://github.com/zeroclaw-labs/zeroclaw/issues/6293)
    -   **讨论热度**：5条评论，跨越一个多月持续更新。
    -   **诉求分析**：此 RFC 提出了将 ZeroClaw 拆分为“离线代理容器”和“在线守护进程”的架构设想，通过 Unix socket 连接，以实现**气隙（Air-gapped）执行模式**。这反映了社区中**高级用户**和**企业用户**对在隔离、高安全环境中运行 AI Agent 的强烈需求，尤其是在处理敏感数据或需要遵守严格合规要求时。`domain:security, domain:architecture, status:blocked` 等标签表明这是一个复杂且可能影响深远的讨论，短期内难以定论，但其潜在价值巨大。
-   **高风险待修复 PR**：[fix(providers/runtime): two production blockers — empty tool_choice and orphaned tool_use](https://github.com/zeroclaw-labs/zeroclaw/pull/5892)
    -   **关注焦点**：尽管该 PR 未显示今日评论，但其标题明确指出了两个**生产环境阻断问题**（`tool_choice`为空和`tool_use`孤立），且涉及多个主流 Provider，在社区中具有极高的关注度。其在接近两个月后仍处于 `needs-author-action` 状态，可能反映了修复的复杂性或作者与维护者之间的沟通瓶颈。

## 5. Bug 与稳定性

今日报告了多个 Bug，主要集中在 **跨平台兼容性**、**配置安全**和**运行时稳定性**上。以下按严重程度排列：

-   **高严重性 (High Risk)**：
    -   **[Windows编译失败]**: `AMQP channel cannot compile on Windows` (#7452) **已关闭**。 (修复PR: 已于 #7452 中解决)
    -   **[Web工具安全漏洞]**: `fix(web_fetch): let allowed_private_hosts = ["*"] cover DNS-resolved private hosts` (#7424) **待合并**。修复了 `web_fetch` 工具的通配符 `["*"]` 未能正确覆盖 DNS 解析后的私有IP的问题，是严重的安全漏洞。
    -   **[Agent运行时Bug]**: `fix(runtime): honor profile tool iteration limits` (#7583) **待作者操作**。修复了 cron/CLI 模式下，agent 运行时未能正确遵守配置中 `max_tool_iterations` 限制的回归。
    -   **[运行时/Provider Bug]**: `fix(providers/runtime): two production blockers` (#5892) **需跟进**。如前所述，两个生产环境阻断问题。
-   **中严重性**：
    -   **[插件发现路径不匹配]**: `fix(plugins): align install/discovery paths and add legacy migration` (#7549) **待合并**。修复 CLI 安装的 WASM 插件因路径不统一而无法被运行时发现的核心问题。
    -   **[配置错误]**: `fix(config): warn on extra-nested provider aliases that silently drop fields` (#7617) **待作者操作**。避免因用户错误配置导致 provider 字段被静默丢弃。
-   **低严重性**：
    -   **[安装脚本兼容性]**: `fix(install): detect musl libc for Linux target triple` (#7614)。修复 Linux 安装脚本的 musl 库检测问题。
    -   **[国际化同步]**: `fix(i18n): sync zh-CN Fluent locale files with English source` (#7612)。同步中文翻译文件。

## 6. 功能请求与路线图信号

-   **安全架构革新** (可能纳入远期路线图)：
    -   [RFC: Air-gapped execution mode with companion daemon](https://github.com/zeroclaw-labs/zeroclaw/issues/6293) – 这是一个架构级的功能请求，旨在实现**零信任网络环境**下的 Agent 执行。如果被接受，将成为 ZeroClaw 安全特性的重大里程碑，预计会出现在 v1.0 或更高版本路线图中。
-   **核心能力增强** (预期在下一版本中出现)：
    -   **Matrix频道增强**: `feat(channels): recover Matrix room management` (#7661) – 恢复并增强了Matrix频道的房间管理功能（创建房间、邀请用户），提升了协作场景的通信能力。
    -   **梦境模式 - 记忆整合**: `feat(memory): add dream mode for periodic memory consolidation` (#6693) – 提出的“梦境模式”旨在实现 Agent 的周期性记忆总结与整合，是通往更强大、更连续 Agent 行为的关键功能。
-   **通信延迟优化** (已合并)：
    -   `feat(cron): add a pause/resume toggle` (#7384) 的合并，体现了项目对**任务调度可管理性**的重视。

## 7. 用户反馈摘要

从今日的 Issue 和 PR 讨论中，可以提炼出以下用户声音：

-   **痛点**：
    -   **跨平台体验受阻**: `AMQP channel cannot compile on Windows` (#7452) 的直接抱怨，Windows 用户无法使用核心的 AMQP 功能。
    -   **配置复杂性导致问题**: `extra-nested provider aliases` (#7617) 和 `secret` macro (#6989) 的讨论表明用户在**配置细节上容易出错**，期望有更好的校验和脱敏机制。
    -   **功能缺失**: `ask_user fails instantly` (#7664) 和 `cron pause/resume` (#7384) 都是用户在日常使用中遇到的直接功能缺失。
-   **需求**：
    -   **高级安全需求**: `Air-gapped execution mode` RFC 的提出，代表高级用户或企业用户对 **“数据不离开安全边界”**的强烈诉求，这是一个未被当前版本满足的高层次需求。
    -   **更智能的Agent**: `dream mode` (#6693) 的功能请求表明用户不满足于简单的提示/响应，而是希望 Agent 能够**自主地、周期性地学习和成长**。

## 8. 待处理积压

以下 Issue/PR 长期未得到响应或解决，可能成为维护者需要关注的“烂尾”问题：

-   **[Urgent] 生产环境阻断问题**：[fix(providers/runtime): two production blockers](https://github.com/zeroclaw-labs/zeroclaw/pull/5892) (自 2026-04-19 起，状态: `needs-author-action`) - 此 PR 修复了两个严重 Bug，但长期处于等待作者行动状态。项目维护者应考虑主动联络作者或指派他人接手。
-   **[重要功能] 梦境模式**：[feat(memory): add dream mode for periodic memory consolidation](https://github.com/zeroclaw-labs/zeroclaw/pull/6693) (自 2026-05-16 起，状态: `needs-author-action`) - 这是一个重要的功能提案，但进展缓慢。维护者应考虑是将其纳入核心路线图还是暂时搁置。
-   **[架构级讨论] 气隙执行模式**：[RFC: Air-gapped execution mode...](https://github.com/zeroclaw-labs/zeroclaw/issues/6293) (自 2026-05-03 起，状态: `status:blocked`) - 该 RFC 讨论方向正确，但显然被卡住了。维护者应组织一次线上讨论或给出初步决策，以推动社区共识。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*