# OpenClaw 生态日报 2026-06-23

> Issues: 39 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-06-23 10:50 UTC

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

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，我将根据您提供的 OpenClaw 项目 GitHub 数据，为您生成 2026-06-23 的项目动态日报。

---

### OpenClaw 项目日报 (2026-06-23)

**数据快照时间范围：2026-06-22 至 2026-06-23 (UTC)**

---

#### 1. 今日速览

今日 OpenClaw 项目处于**高活跃度**状态。社区提交了创纪录的 500 条 PR，但其中绝大多数（416 条）处于待合并状态，显示出巨大的潜在贡献量。核心开发与维护活动集中在**会话状态恢复、ACP 运行时兼容性**以及**跨平台稳定性**等关键领域。虽然过去 24 小时内无新版本发布，但多个高优先级（P1）Bug 获得修复，项目稳定性有望提升。大量长时间未合并的 PR 积压是项目当前需要关注的风险点。

#### 2. 版本发布

无。

#### 3. 项目进展

尽管今日无 PR 合并，但活跃的讨论和 Bug 修复关闭标志着项目在以下方面取得进展：

- **核心稳定性提升**：多个影响会话状态恢复的 Bug 被关闭，例如 #89868 中修复了 `Stop-button` 与 `compaction` 完成时的竞态条件。这直接提升了生产环境应用的可靠性。
- **跨平台兼容性**：针对 Windows 和 WSL 环境的 Bug 得到确认和修复，如 #89231 和 #96023，体现了项目对非 Linux 平台用户的支持正在加强。
- **开发者工具链优化**：PR #78664 (Tool schema normalization) 等长期存在的性能优化 PR 状态更新为 `ready for maintainer look`，表明核心开发者正在审视这类提升性能的贡献，这对降低运行成本至关重要。
- **插件生态修复**：针对 Telegram、Discord 等渠道的特定 Bug (如 #95878, #96007) 有对应的修复 PR，社区的自愈能力较强。

#### 4. 社区热点

今日社区讨论最活跃的议题主要围绕**会话状态持久化**与**ACP (Agent Communication Protocol)** 运行时的稳定性。

1.  **`#92201` - Anthropic 思考签名验证失败 (Diamond Lobster Priority)**
    - *链接*: [Issue #92201](https://github.com/openclaw/openclaw/issues/92201)
    - *分析*：社区讨论最多的议题。该问题直指核心 LLM 交互流程，在重放流式思考块时签名验证不通过，且错误处理机制失效。高星级（Diamond Lobster）表明这是一个棘手且影响巨大的核心Bug，开发者可能正在紧急讨论最佳修复方案。

2.  **`#92043` - 180秒压缩超时问题 (Diamond Lobster Priority)**
    - *链接*: [Issue #92043](https://github.com/openclaw/openclaw/issues/92043)
    - *分析*：此议题涉及会话历史压缩（Compaction）的核心机制。社区讨论集中在一个设计缺陷：超时计时器是全局的一次性计时器，而不是按阶段分配的。这导致大量用户（尤其是使用慢速或本地模型）的会话在压缩时反复失败，无法从长对话中恢复。

**核心诉求**：社区渴望项目能在核心数据流（签名、超时、恢复）上提供更高的**可靠性**和**容错性**。当前的实现方式在边缘场景下表现脆弱，影响用户体验。

#### 5. Bug 与稳定性

过去 24 小时报告的 Bug 数量较多（27个新开），严重程度普遍较高。按优先级排列如下：

- **严重 (P1 / Diamond Lobster) 风险高**:
    - **`#95750` - 主会话重启恢复无重试预算** - 损坏的会话导致网关跨重启死循环。**已有 PR**。
    - **`#96023` (已关闭)** - WSL 环境下 `interop=false` 导致设置失败。**已修复**。
    - **`#93465` - Windows 嵌入式 ACPX 运行时不可用** (`spawn EINVAL`) - 影响所有 Windows 用户使用 Claude Code 等 ACP 适配器。**已有 PR**。
    - **`#95141` - 配置变更重启导致会话丢失** - 用户的重要对话数据在重启过程中不可见地丢失。
    - **`#95998` - 代理设置破坏腾讯云 COS 上传** - 特定环境下的回归Bug，影响中国大陆用户使用相关服务。

- **中等 (P2 / Platinum Hermit)**:
    - **`#95701` - 视频生成任务死锁** - 背景任务完成后状态未及时更新，阻碍后续任务。**已有 PR**。
    - **`#96009` - 其他** - 工具执行错误信息未正确隐藏。
    - **`#96020` - 会话消息投递状态错误** - 系统误以为消息未送达。

**稳定性总结**：**ACP 运行时**和**会话状态管理**是当前稳定性的最大短板。多个 P1 级别 Bug 显示，在不同环境下（Windows、特定模型、网络条件），系统健壮性不足。好消息是，大量此类 Bug 已有对应的修复 PR 在审查中。

#### 6. 功能请求与路线图信号

- **`#88032` - Telegram 引用/回复作为一等耐久组件** - 用户强烈希望 Telegram 的回复功能能成为一个核心、稳定、有文档支持的合约，而非需要频繁打补丁的表面功能。这预示着项目可能在未来强化对 Telegram 平台的高级特性支持。
- **`#96056` - OpenShell 沙箱支持非秘钥环境变量** - 社区要求为 OpenShell 沙箱添加传递普通环境变量的能力，以使沙箱配置更灵活，避免将非敏感配置（如 API 基础 URL）误当作秘钥处理。
- **`#95752` - Control UI 折叠工具调用** - 用户希望 Dashboard UI 默认折叠冗长的工具执行过程，以减少视觉干扰，只展示最终结果。这是一个提升用户体验的常见请求。

**路线图信号**：社区需求正在从“能否运行”转向“运行体验是否优雅”。请求集中在**UI/UX 优化**和**配置灵活性**上。如果项目要发布大版本，`telegram quote/reply` 的稳定化和 `OpenShell` 配置增强很可能会被纳入范围。

#### 7. 用户反馈摘要

- **正面反馈（隐含）**: Issue #96023 (WSL 无法配对) 在短时间内被提出并关闭，说明**配置流程文档**和**WSL 社区**的支持是有效的。
- **痛点**:
    - **Windows 支持不足**：`#93465` 明确指出了 Windows 用户无法使用 ACPX 运行时，这是对 Windows 用户群体的重大阻碍。
    - **数据丢失焦虑**：`#95141` (配置变更导致会话丢失) 和 `#95750` (主会话死循环) 的讨论反映了用户对数据持久化和恢复机制的担忧。用户希望即使出现问题，对话也能被安全地保存和恢复。
    - **“静默”失败**：`#95759` (ACP 会话输出为 0 字节，无错误) 和 `#92043` (压缩反复失败) 让用户在发现问题前浪费了大量时间。用户明确希望系统在非正常状态下能给出**明确的错误提示**。

#### 8. 待处理积压

以下为长期未关闭或状态停滞的重要议题，需要维护者关注：

- **`#71712` - Agent 调度 API (已逾2月，P2)**
    - *链接*: [Issue #71712](https://github.com/openclaw/openclaw/issues/71712)
    - *风险*：这是一个关于**Agent 自主性**的重要 RFC。长期无人回应可能导致设计思想过时，或社区贡献者转向其他项目。

- **`#79200` - 添加 `--message-file` CLI 标志 (已逾45天，P3)**
    - *链接*: [PR #79200](https://github.com/openclaw/openclaw/pull/79200)
    - *风险*：该 PR 解决了一个常见的 CLI 使用痛点（Shell 转义），但状态长期为 `waiting on author`。如果作者失活，这个有价值的贡献将永远无法落地。

- **大量“等待证明/等待作者”的 PR**: 如 `#79342`, `#78742`, `#78958` 等，状态为 `needs proof` 或 `waiting on author`。这表明项目对 PR 的审查标准较严格，但也可能导致许多高质量贡献被长期搁置，挫伤贡献者积极性。建议项目组考虑引入社区 reviewer 机制。

---

## 横向生态对比

好的，作为AI智能体与个人AI助手领域的资深技术分析师，我已审阅您提供的2026年6月23日各开源项目动态日报。现基于这些一手数据，为您呈现一份横向对比分析报告。

---

### **个人AI助手与自主智能体开源生态全景洞察（2026-06-23）**

#### **1. 生态全景**

今日，个人AI助手与自主智能体开源生态呈现出 **“核心分化，边缘激进”** 的态势。以OpenClaw为核心的项目家族（如PicoClaw、NanoClaw）正面临大规模PR积压与核心会话稳定性的双重挑战，这表明生态在急速扩张后进入了关键的 **质量巩固期**。与此同时，NanoBot、CoPaw、ZeroClaw等外围项目则通过高频发版和快速修复Bug来争夺特定场景的用户，尤其是在**定时任务可靠性**和**多模型/多平台兼容性**上展开激烈竞争。整体来看，社区集体关注的焦点已从“能否实现功能”转向“功能是否稳定”，会话持久化、跨平台兼容性（特别是对国产模型和Windows的支持）以及开发者体验（工具链、配置灵活性）成为决定下一阶段项目竞争力的核心要素。

#### **2. 各项目活跃度对比**

| 项目 | Issues (更新/新开) | PRs (更新/新开) | 版本发布 | 健康度评估 |
| :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 27+ 新开 | 500+ (多数待合并) | 无 | **风险累积** (PR积压严重，核心Bug频发) |
| **NanoBot** | 5+ 新开/更新 | 33+ (9个合并) | **v0.2.2** | **快速迭代** (发布后紧密修复回归) |
| **Hermes Agent**| 21 更新 | 50 更新 | 无 | **健康高产** (社区贡献活跃，多发P1 Bug有PR) |
| **PicoClaw** | 2 新开 | 18 (6个合并) | 无 | **稳定修复** (侧重底层健壮性与安全性) |
| **NanoClaw** | 无 | 10 (1个合并) | 无 | **集成构建** (大规模版本对齐与功能开发) |
| **NullClaw** | 无 | 1 (待合并) | 无 | **低活跃维护** (单一Bug修复等待合并) |
| **IronClaw** | 10 更新 | 33 更新 (12个合并) | 无 | **功能跃进** (合并大量自动化与权限功能) |
| **LobsterAI** | 1 更新 | 3 合并 | 无 | **社区停滞** (严重Bug 80天未获官方回应) |
| **Moltis** | 无 | 1 (合并) | 无 | **温和整合** (长期PR合并，基础功能补全) |
| **TinyClaw** | - | - | - | **静止** (无活动) |
| **ZeptoClaw** | - | - | - | **静止** (无活动) |
| **ZeroClaw** | 1+ 新开 | 50+ (7个合并) | 无 | **高度活跃** (密集合并，响应迅速) |
| **CoPaw** | 14 更新 | 50 更新 | **v1.1.12.post2** | **积极迭代** (大量合并，修补定时任务等核心功能) |

#### **3. OpenClaw 在生态中的定位**

*   **核心参照与问题发现者**：OpenClaw是生态中当之无愧的核心参照项目和“金丝雀”。它承受着最复杂的场景（如会话压缩、ACP运行时），因此**最先暴露**出共性痛点，例如会话状态恢复、跨平台兼容性（WSL/Windows）、工具执行反馈等问题。这些问题是整个生态都需要面对和解决的。
*   **独特优势**：其社区提出的Bug报告（如 `Diamond Lobster` 优先级）和议题讨论深度最高，是理解行业痛点的**最佳情报源**。其 `OpenShell` 沙箱、`Agent 调度 API` 等设计也是其他项目（如PicoClaw）功能演进的参考。
*   **技术路线差异**：相比NanoBot、CoPaw等每月发版的项目，OpenClaw的**发布节奏更保守**，更注重架构的健壮性。这导致了巨大的PR积压（500条），使其在**反应速度**上落后于外围项目。当前其路线图信号（如UI优化、Telegram高级特性）显示其正试图从底层架构向用户体验延伸。
*   **社区规模**：从每日PR/Iusse数量看，OpenClaw拥有**最大的贡献者社区**。然而，社区规模大也带来了维护挑战，大量“需要等待作者/需要证明”状态的PR反映了代码审查流程可能成为其发展的瓶颈。

#### **4. 共同关注的技术方向**

1.  **会话状态持久化与恢复**：几乎成为所有活跃项目的核心痛点。OpenClaw的关键Bug (#89868, #92043, #95141) 和Hermes Agent的Bug (#27038)、NullClaw的修复PR (#968) 都指向了这一问题。社区对 **数据不丢失、服务重启后无缝恢复** 的需求极其强烈。
2.  **定时任务（Cron/Trigger）的可靠性**：CoPaw、LobsterAI、NanoBot 均报告了定时任务不执行、状态错乱、心跳消息错误发送等问题。这表明定时自动化是用户高频使用的核心场景，但其实现远未成熟，错误处理（如 `misfire_grace_time`）和状态管理是薄弱环节。
3.  **跨平台与多模型兼容性**：OpenClaw、NanoBot、Hermes Agent、CoPaw 均报告了Windows、WSL及特定LLM Provider（如DeepSeek、Z.AI、Groq）的兼容性Bug。**支持更广泛的模型和操作系统**是项目扩大用户基础的关键。
4.  **Gateway/频道精细化控制**：Hermes Agent 社区的 **per-channel 配置 (Issue #1955)** 和 **多Telegram Bot管理 (Issue #10452)** 请求，以及NanoClaw的Slack Socket Mode PR，共同指向了用户对**灵活、隔离、可编程的通信渠道管理**的强烈需求。
5.  **开发者体验与可观测性**：ZeroClaw 的trace_id和请求负载捕获讨论、OpenClaw的 `--message-file` CLI标志、NullClaw 的测试环境隔离，都表明**高级用户和贡献者**对更好的调试、测试和运维工具有普遍诉求。

#### **5. 差异化定位分析**

| 项目 | 功能侧重 | 目标用户 | 技术架构关键差异 |
| :--- | :--- | :--- | :--- |
| **OpenClaw** | **通用型自主Agent**、高度可定制化 | 核心开发者、高级用户、深度定制者 | 自有一套完整的 `OpenShell` 沙箱和 `Agent Session Runtime`，架构复杂但高度灵活。 |
| **NanoBot** | **快速启动**、`Dream` (自进化技能)、Telegram深度集成 | 追求即时反馈、喜欢“魔法”功能的个人用户 | 迭代速度最快，“Dream”功能是其杀手锏，但新特性（如Telegram 10.1）引入的回归较多。 |
| **CoPaw** | **定时任务与Dream任务**、桌面端体验 | 需要自动化工作流的C端用户 | 与OpenClaw体系关联但更聚焦于C端体验（如实时SSE通知、多步骤进度面板），依赖Tauri客户端。 |
| **Hermes Agent** | **Gateway平台 (Discord/Telegram) 管理**、可扩展性 | 需要管理多频道/多Bot的运维者 | **Gateway层是核心优势**，社区对per-channel配置和多Bot支持的需求最强。 |
| **PicoClaw** | **代码安全性与健壮性**、跨平台工具 | 关注代码质量、安全可靠运行的开发者 | 代码审查严格，侧重修复底层类型安全、连接稳定性、安全漏洞，是“稳定派”代表。 |
| **ZeroClaw** | **可观测性、SOP (标准操作程序)**、多Provider | 运维/DevOps背景的开发者、企业级场景 | 近期亮点在可观测性（trace_id）和SOP持久化，显示出对**审计和运维**的重视。 |
| **LobsterAI** | OpenClaw的**中国本土化**和UI集成 | 中文用户群体 | 中文环境下与OpenClaw联动，但更新缓慢，积压了大量功能和安全修复。 |

#### **6. 社区热度与成熟度**

*   **快速发展与迭代期 (高热度 + 快速发版)**：**NanoBot, CoPaw, ZeroClaw, IronClaw**。这些项目在24小时内均有大量PR合并或小版本发布，快速响应社区反馈，但Bug修复与功能开发并行，稳定性波动较大。
*   **质量巩固与功能完善期 (高讨论度 + 审查严格)**：**OpenClaw, Hermes Agent, PicoClaw**。这些项目社区活跃，贡献量大，但合并标准高，PR积压严重。它们正经历从功能堆叠到质量巩固的痛苦转型期。**PicoClaw** 在这一阶段表现出色，在稳定性和安全性上口碑较好。
*   **稳定维护或低活跃期**：**NullClaw, Moltis, TinyClaw, ZeptoClaw, LobsterAI**。这些项目更新频率低，多为 Bug 修复或长期 PR 的合并。其中 **LobsterAI** 因严重 Bug 长期未获回应，社区信任度正在流失。

#### **7. 值得关注的趋势信号**

1.  **从“能用”到“好用”的体验转向**：社区不再满足于Agent能执行任务，而是要求**任务执行过程的透明度**（CoPaw的进度面板、OpenClaw的折叠工具调用）、**交互的流畅性**（NanoBot的UI闪烁修复、CoPaw的实时通知）和**反馈的完整性**（NanoClaw的审批拒绝理由）。
2.  **从单一集成到生态兼容**：支持多模型（如Z.AI、Groq、DeepSeek等国产/新兴模型）和去中心化协议（如PicoClaw的SimpleX请求）成为共同趋势。**对特定服务的绑定对项目发展是重大风险。**
3.  **从功能开发到治理挑战**：OpenClaw 的500条PR积压、LobsterAI 80天未响应的严重Bug，都揭示了**社区治理和代码审查流程**成为项目能否健康持续发展的关键瓶颈。未来，引入社区维护者角色、自动化审查工具（如Dependabot，但需要更智能）可能是破解之道。
4.  **从通用Agent到垂直场景渗透**：除了通用对话，NanoBot的“Dream”自进化技能、ZeroClaw的SOP持久化、CoPaw的定时任务，均显示出Agent正在向 **“代码维护”、“流程自动化”、“定时执行”** 等明确场景深入，从“聊天机器人”向“数字员工”的角色转变。

**结论与建议**：
对于技术决策者，**选择项目时需要权衡“功能丰富度”与“生态稳定性”**。近期关注 **PicoClaw** 和 **ZeroClaw** 在稳定性与可观测性上的进展。对于开发者和贡献者，当前进入 **CoPaw、NanoBot、IronClaw** 等快速迭代项目能获得更快的反馈，但代码生命周期可能更短。而深入 **OpenClaw** 或其家族项目（如PicoClaw），则需要做好应对PR审查周期长、积压问题的准备，但其掌握的底层技术更具深度。整体上，**会话状态管理、跨平台兼容性、Gateway灵活性**是未来半年内最值得关注的投资方向。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

好的，作为AI智能体与个人AI助手领域开源项目分析师，我根据NanoBot项目在2026年6月23日的GitHub数据，为您生成了以下项目动态日报。

---

# NanoBot 项目动态日报 | 2026-06-23

项目状态：**高度活跃** | 健康状况：**良好** (修复与功能开发并行)

### 1. 今日速览

今日项目迎来 **v0.2.2** 正式发布，这是自v0.15以来的一次重大升级，累计合并了140个PR，吸引了21位新贡献者。项目在“坚固性”方面有显著提升，但也引入了部分回归性问题，主要集中在新集成的 **Telegram Bot API 10.1** 富消息功能上。社区响应迅速，今日**33个PR**中包含了多个针对该回归问题的修复提案，展现了项目快速迭代的风范。同时，老代码中的一些长期BUG（如心跳机制）也在此次升级后被暴露和修复。整体来看，项目正处于向新架构和大版本跃进的关键磨合期，社区活跃度极高。

### 2. 版本发布

- **版本**: v0.2.2
- **发布日期**: 2026-06-23 (基于PR #4461)
- **发布说明**:
  - 本次发布是v0.15后的一个重大里程碑，标题强调“坚固性 (durability)”的提升。
  - WebUI 对话体验得到强化：对话记录被分段存储，不再依赖单一的脆弱文件；分叉对话更可靠地保留回复。
  - 此版本是140个PR和21位新贡献者辛勤工作的结晶。
- **破坏性变更 / 迁移注意事项**:
  - **Telegram Bot API 10.1 集成**: 引入了 `sendRichMessage`，目的是支持更丰富的消息格式，但导致了两个关键回归性问题：
    1. **换行符丢失**：消息内容被合并为单个文本块。
    2. **消息闪烁**：网关在生成过程中反复编辑消息，导致UI闪烁。
  - **心跳机制行为变更**: 此前绑定会话的定时任务（Cron）会无条件发送LLM返回的信息。现在即使任务“无事可报”，也可能发送消息。 (Issue #4410)
  - **建议**: 如果您是Telegram重度用户或依赖心跳功能，需关注针对上述问题的修复PR (`#4472`, `#4412`)，建议等待下一个补丁版本。

### 3. 项目进展

今日项目向前迈进了关键一步，已合并/关闭了9个PR，主要集中在以下几个方向：

- **新版本发布与文档**: PR #4461 正式将 `v0.2.2` 写入文档，将新闻列表中的旧条目归档。
- **稳定性修复**:
  - **Cron通知优化**: PR #4412 合并，修复了`v0.2.2`中定时任务“无事也汇报”的问题，增加了抑制常规性回答的机制。
- **功能与架构演进**:
  - **PWA支持**: PR #4458 (WebUI的PWA支持) 被创建并标记为开放中，显示了项目向移动端原生体验演进的方向。
  - **Mattermost集成**: PR #4459 提出了全新的Mattermost频道支持，正在等待合并，这将进一步扩大NanoBot的平台覆盖范围。
  - **Dream功能优化**: PR #4469 提出在Dream运行时注入现有工作区技能，以避免重复创建，该PR与 `#4467` 功能请求强相关，表明社区需求已获得核心开发者响应。

### 4. 社区热点

今日社区讨论最热烈的议题集中在 `v0.2.2` 带来的界面与交互体验问题。

1. **Telegram 显示回归 (Issue #4470 & PR #4472)**
   - **热度**: 该issue在提出后很快便有开发者跟进并提交了修复PR (`#4472`)，交互非常迅速。
   - **诉求**: 用户 `chengyongru` 明确指出了`v0.2.2`后Telegram上的两个具体BUG：换行失效和消息闪烁。这反映了用户对即时通讯工具中基础排版功能和交互流畅性的高要求。
   - **链接**: [Issue #4470](https://github.com/HKUDS/nanobot/issues/4470), [PR #4472](https://github.com/HKUDS/nanobot/pull/4472)

2. **WebUI 思考标签渲染问题 (Issue #4465 & PR #4466)**
   - **热度**: 该问题被提出后，作者`ZhouJ-sh`立即提交了修复PR `#4466`，显示出社区对UI细节的重视和开发者积极响应的态度。
   - **诉求**: 用户希望`<thinking>`标签能被正确处理为推理内容块，而非直接显示为可见文本。这涉及到模型输出在UI层的语义化渲染，是提升用户体验的重要细节。
   - **链接**: [Issue #4465](https://github.com/HKUDS/nanobot/issues/4465), [PR #4466](https://github.com/HKUDS/nanobot/pull/4466)

3. **Dream功能重复创建技能 (Issue #4467 & PR #4469)**
   - **热度**: 虽然Issue没有直接引发大量评论，但它迅速催生了修复PR (`#4469`)，表明这是一个影响用户工作效率的核心痛点，开发者也认同其重要性。
   - **诉求**: 用户 `songsong-hui` 提出了一个典型的“自动化期望”与“现实行为”的冲突：用户期望Dream能**更新**现有技能，而不是每次都**新建**一个副本，导致工作区杂乱且手动合并工作量大。
   - **链接**: [Issue #4467](https://github.com/HKUDS/nanobot/issues/4467), [PR #4469](https://github.com/HKUDS/nanobot/pull/4469)

### 5. Bug 与稳定性

今日报告的Bug集中爆发于 `v0.2.2` 的新特性，按严重程度排列如下：

- **严重: Telegram 消息显示异常 (Issue #4470)**
  - **描述**: `v0.2.2` 后，Telegram消息换行符被忽略，且生成过程中消息频繁闪烁。这严重影响了用户体验。
  - **状态**: 已有修复PR `#4472` 提交。
  - **链接**: [Issue #4470](https://github.com/HKUDS/nanobot/issues/4470)

- **高: WebUI 渲染`<thinking>`标签为可见文本 (Issue #4465)**
  - **描述**: 模型输出的`<thinking>`标签未被成功解析为推理块，而是作为纯文本暴露给用户。
  - **状态**: 已有修复PR `#4466` 提交。
  - **链接**: [Issue #4465](https://github.com/HKUDS/nanobot/issues/4465)

- **中: 心跳任务错误发送消息 (Issue #4410)**
  - **描述**: 升级后，即使心跳任务的LLM回答为“无需操作”，消息也会被发送到对话中。
  - **状态**: 已修复，PR `#4412` 已合并。
  - **链接**: [Issue #4410](https://github.com/HKUDS/nanobot/issues/4410)

- **中: 心跳任务应排除归档会话 (PR #4468)**
  - **描述**: 心跳任务仍会检查归档的session，可能导致不必要的行为。
  - **状态**: 修复PR `#4468` 已提交。
  - **链接**: [PR #4468](https://github.com/HKUDS/nanobot/pull/4468)

### 6. 功能请求与路线图信号

- **高可能纳入下版本: Dream技能更新机制 (Issue #4467)**
  - 与已提交的PR `#4469` 对应，表明这是当前开发重点之一。此功能将大幅提升用户利用`Dream`维护代码的体验。
  - **链接**: [Issue #4467](https://github.com/HKUDS/nanobot/issues/4467)

- **中可能纳入下版本: Kimi Coding Plan 支持 (Issue #4463)**
  - 用户`zpljd258`同时提出了Issue和PR `#4464`，显示出清晰的规划和开发意愿。如果Kimi是合作伙伴，这一集成将增强NanoBot在编码领域的竞争力。
  - **链接**: [Issue #4463](https://github.com/HKUDS/nanobot/issues/4463), [PR #4464](https://github.com/HKUDS/nanobot/pull/4464)

- **低/观望: PWA 支持 (Issue #4457)**
  - 已提交相关PR `#4458`，但被标记为`invalid`，可能当前合并优先级不高。但PWA是提升用户体验的重要方向。
  - **链接**: [Issue #4457](https://github.com/HKUDS/nanobot/issues/4457)

### 7. 用户反馈摘要

- **对v0.2.2新特性的直接反馈**:
  - 用户`chengyongru`在反馈Telegram问题时，详细描述了“2个Bug”，并给出了明确的现象（换行丢失、闪烁），这种高质量的反馈对开发者定位问题非常有帮助。
- **对长期功能点的期望**:
  - 用户`madIlama`和 `songsong-hui` 提出的功能请求都伴随着具体的痛点和使用场景。`madIlama`关注于消息格式本身的先进性，而`songsong-hui`则在意工作流的效率与整洁性。这些反馈揭示了用户从“能用”到“好用”、“智能”的进阶需求。
- **用户对“自然语言睡眠搜索”的喜爱**:
  - 通过观察v0.2.2的release notes，虽然未在今日数据中详细展开，但可以推断用户社区对“自然语言睡眠搜索”这一创新功能反响热烈，它是本次版本的一大亮点。

### 8. 待处理积压

- **长期未解决的功能请求 (Issue #2305)**
  - **内容**: 支持隐藏推理步骤显示。该Issue从3月就已开启，虽然社区用户`EvanNotFound`提了早期版本，但直到今日`#4465`提出类似问题并出现`#2305`的链接时，它才被重新激活。
  - **建议**: 鉴于WebUI渲染`<thinking>`标签已成为一个热点问题，维护者可以在修复PR `#4466`的基础上，考虑同时实现`#2305`中提出的“隐藏推理步骤”功能，作为完整的UI改进方案。
  - **链接**: [Issue #2305](https://github.com/HKUDS/nanobot/issues/2305)

- **待合并的`fix(providers): DeepSeek...` PR (PR #3869)**
  - **内容**: 从5月就提出的针对DeepSeek provider的数据清理和格式修复PR。
  - **状态**: 仍标记为`[question]`，等待审核或更多讨论。虽然功能正常，但该PR能防止潜在的API错误，应尽快处理。
  - **链接**: [PR #3869](https://github.com/HKUDS/nanobot/pull/3869)

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

好的，遵照您的指示，以下是为期 2026-06-23 的 Hermes Agent 项目动态日报。

---

# Hermes Agent 项目日报 | 2026-06-23

## 今日速览

今日项目活跃度非常高，共产生 21 条 Issue 更新和 50 条 PR 更新。社区提交的 Bug 修复和功能请求数量显著，涉及桌面端、Dashboard、CLI、Gateway 及多个平台适配器。尽管没有新版本发布，但社区贡献者积极提交 PR 解决高优先级 Bug，项目整体处于健康且高产的迭代状态。**值得关注的是，今日出现了多个由社区贡献的针对关键 Bug 的修复 PR，显示出社区维护力量的快速增长。**

## 版本发布

今日无新版本发布。

## 项目进展

尽管没有正式版本发布，但今日社区贡献者在修复 Bug 和推进功能方面取得了实质性进展。以下为今日最关键的合并/关闭的 PR：

1.  **[PR #51311] fix(gateway): runner takes over TTS when adapter won't auto-TTS for Discord DM voice_only** - 已关闭 (合并)
    -   **摘要**：修复了 Discord 平台在私信（DM）场景下，语音输入无法自动触发 TTS（文本转语音）回复的问题。现在 Gateway Runner 会在适配器无法处理时接管 TTS 任务，确保了语音交互的完整性。
    -   **影响**：提升了 Discord 用户的语音交互体验，修复了一个重要的功能缺失。
    -   **链接**: NousResearch/hermes-agent PR #51311

## 社区热点

今日讨论最热烈的议题集中在**多平台适配**和**核心用户体验**方面。

1.  **Issue #1955 - feat: per-channel model and system prompt overrides for gateway platforms**
    -   **热度**：9 条评论，5 个 👍
    -   **分析**：这是社区呼声很高的功能请求，希望能够在 Discord/Telegram 等网关平台的**不同频道**内，为 Agent 配置不同的模型和系统提示词。例如，一个“日常”频道使用低成本模型，“开发”频道使用高性能模型。这表明用户对精细化的网关管理和成本控制有强烈需求。
    -   **链接**: NousResearch/hermes-agent Issue #1955

2.  **Issue #10452 - Support multi Telegram bots for gateway routing and send_message**
    -   **热度**：7 条评论，4 个 👍
    -   **分析**：同样是一个关于 Gateway 灵活性的需求。用户希望单个 Hermes 实例能同时管理多个 Telegram 机器人。这通常用于生产环境，例如将主要助手机器人和内部开发机器人分离。该议题与 #1955 共同指向了 Gateway 平台多实例、精细化管理的核心诉求。
    -   **链接**: NousResearch/hermes-agent Issue #10452

## Bug 与稳定性

今日报告的 Bug 数量较多，其中不乏影响核心体验的严重问题，但社区已迅速提供了部分修复 PR。

**高优先级 (P1-P2)**

1.  **Issue #4707 (P1): cron under profile-scoped launchd gateway falls back to default `~/.hermes`**
    -   **问题**：macOS 上使用 `launchd` 运行多配置文件网关时，定时任务（cron）会错误地回退到默认配置目录，导致配置丢失。
    -   **严重程度**：高。影响使用 macOS 并依赖多配置文件的专业用户。
    -   **链接**: NousResearch/hermes-agent Issue #4707

2.  **Issue #27038 (P2): Codex Responses API rejects replayed assistant message items with long id fields**
    -   **问题**：在恢复或继续会话时，由于 `id` 字段过长，被 Codex Responses API 拒绝，导致会话无法继续。
    -   **严重程度**：高。直接导致“继续会话”功能失败，严重影响使用体验。
    -   **已有修复 PR**：注意到 **PR #25268** 和 **PR #51312** 等修复了相关的会话状态问题，但此特定问题尚未有直接关联的修复 PR。
    -   **链接**: NousResearch/hermes-agent Issue #27038

3.  **Issue #49200 (P2): Configured memory provider fails silently**
    -   **问题**：当外部存储提供者（如 Mnemosyne）加载失败时，Hermes 会静默回退到内置存储，不向用户发出任何警告。这可能导致用户误以为自己的记忆配置生效，而实际上并未使用。
    -   **严重程度**：高。属于“静默失败”问题，对用户造成误导。
    -   **链接**: NousResearch/hermes-agent Issue #49200

4.  **Issue #47685 (P2): Z.AI GLM-5.2 Coding Plan returns 429 when system prompt contains "Hermes Agent"**
    -   **问题**：一个非常有趣的 Bug。当系统提示词中包含“Hermes Agent”字样时，Z.AI 的 GLM-5.2 模型会返回 429 限流错误。这看起来像是 Z.AI 服务端的一个内容过滤或匹配逻辑问题。
    -   **严重程度**：中高。虽非崩溃性 Bug，但行为诡异，难以排查。
    -   **链接**: NousResearch/hermes-agent Issue #47685

**中低优先级 (P3)**

1.  **Issue #51321 (P3): 桌面端 "Artifacts" 视图显示文件时间错误**
    -   **问题**：Windows 桌面端文件时间显示错误。
    -   **链接**: NousResearch/hermes-agent Issue #51321
2.  **Issue #51320 (P3): 桌面端聊天输入框文字发送后消失**
    -   **问题**：macOS 桌面端输入文字发送后，文字在输入框和聊天记录中短暂消失，降低用户信心。
    -   **链接**: NousResearch/hermes-agent Issue #51320

## 功能请求与路线图信号

1.  **Gateway 精细化管理**：**Issue #1955 (per-channel 配置)** 和 **Issue #10452 (多 Telegram Bot)** 是今日最突出的功能请求，预计会是未来 Gateway 模块升级的重要方向。
2.  **桌面端增强**：**Issue #41222 (集成看板到桌面应用)** 和 **Issue #34390 (Dashboard 添加 `--allowed-hosts` 标志)** 显示用户希望桌面端功能更强大、更易于集成到现有工作流中。
3.  **新 Provider 支持**：**Issue #51319 (新增豆包/火山引擎)** 是一个由中国用户发起的需求，要求将字节跳动的模型作为原生 provider 支持。这与支持 DeepSeek、Kimi 等趋势一致，表明中国市场对多样化 provider 的需求。
4.  **TUI 可配置性**：**Issue #51288 (添加 `HERMES_TUI_WS_WRITE_TIMEOUT_S` 环境变量)** 表明高级用户希望有更多底层参数控制权，以应对特定的网络环境。
5.  **RTL 文本支持优化**：**Issue #51318 (RTL/Bidi 文本方向检测问题)** 提出了一个重要的国际化问题，建议改进对阿拉伯语等 RTL 语言的文本方向确定算法。

## 用户反馈摘要

-   **对 Gateway 灵活性的强烈渴求**：从 #1955 和 #10452 的讨论中可以看出，用户不满足于“一个机器人对应一个模型”的模式，期待“一个网关，多个频道/机器人，独立配置模型、系统提示和路由”的架构。
-   **对特定 Provider 的痛点**：**Issue #50663** 的用户反馈 z.ai 在高峰时段对 Hermes 的限流非常严格，影响了 Coding Plan 的使用体验。
-   **对桌面端稳定性的关注**：**Issue #51320** 的用户虽然表达了对项目的感谢，但“发送文字消失”的 Bug 明显降低了其使用信心。
-   **开发者体验的抱怨**：**Issue #41499** 指出，Mac 开发者在构建桌面版时会因为代码签名证书问题而失败，这阻碍了潜在的贡献者参与开发。

## 待处理积压

以下是一些长期存在或对项目健康度至关重要，但目前未得到足够响应的 Issue/PR：

1.  **Issue #1955 (per-channel config)**：此 Feature Request 创建于 2026-03-18，讨论热度高且需求明确，建议项目维护者考虑将其纳入短期路线图，并给出官方反馈。
    -   **链接**: NousResearch/hermes-agent Issue #1955
2.  **Issue #10452 (Multi Telegram Bot)**：与 #1955 类似的 Gateway 增强需求，同样值得维护者关注并回应。
    -   **链接**: NousResearch/hermes-agent Issue #10452
3.  **Issue #27038 (Codex id field rejection)**：严重影响核心会话功能的 P2 Bug，目前没有看到直接的 fix PR 或官方人员的回复，需要优先关注。
    -   **链接**: NousResearch/hermes-agent Issue #27038
4.  **Issue #4707 (macOS cron fallback)**：一个 P1 级别的 Bug，影响特定用户群体的核心功能，亟待被处理或分配。
    -   **链接**: NousResearch/hermes-agent Issue #4707
5.  **Issue #47685 (Z.AI 429 on "Hermes Agent")**：一个令人困惑的 Bug，可能指向上游 Z.AI 平台的兼容性问题。建议优先排查是 Hermes 自身问题还是需要与 Z.AI 沟通。
    -   **链接**: NousResearch/hermes-agent Issue #47685

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

好的，以下是根据您提供的 GitHub 数据生成的 **PicoClaw 项目动态日报**。

---

# PicoClaw 项目动态日报 | 2026-06-23

## 1. 今日速览

过去24小时内，项目保持高度活跃，共处理 **18 条 Pull Requests**，其中 **6 条成功合并或关闭**，工作量显著。Issues 方面有 **2 条新提交**，其中一条为功能请求，另一条为 Bug 报告。未发布新版本。整体来看，社区提交代码的热情高涨，尤其在修复与稳定性方面取得了实质性进展，但关闭的 PR 多为功能修复，待合并的 PR 数量（12 条）仍较多，说明代码审核流程可能成为下一个瓶颈。

## 2. 版本发布

无新版本发布。

## 3. 项目进展

今日有 **6 条 PR 被合并或关闭**，主要集中在代码健壮性提升与基础设施优化方面：

- **[#3162] [已合并] fix(whatsapp): add reconnection and async message processing**  
  修复了 WhatsApp 集成中的 WebSocket 自动断开问题，通过添加 pong 处理、读取超时与指数退避的重连机制，显著增强了消息通道的稳定性。  
  链接: [PR #3162](https://github.com/sipeed/picoclaw/pull/3162)

- **[#3152] [已合并] add installation instructions to picoclaw skills search**  
  增强了 `picoclaw skills search` 命令的输出，在搜索结果中直接显示如何安装该技能，提升了用户体验。  
  链接: [PR #3152](https://github.com/sipeed/picoclaw/pull/3152)

- **[#3053] [已关闭] fix(evolution): add ok check for LoadOrStore type assertion**  
  修复了 `lockStoreFile` 函数中缺少类型断言检查导致的潜在 panic 问题，提高了并发场景下的稳定性。  
  链接: [PR #3053](https://github.com/sipeed/picoclaw/pull/3053)

- **[#3091] [已关闭] fix(openai_compat): add ok check for native_search type assertion**  
  修复了 OpenAI 兼容提供者中，`native_search` 类型断言失败时会静默禁用搜索功能的问题，避免了无声的故障模式。  
  链接: [PR #3091](https://github.com/sipeed/picoclaw/pull/3091)

- **[#3101] [已关闭] build(deps-dev): bump vite**  
  **[#3105] [已关闭] build(deps-dev): bump eslint**  
  两项依赖升级（Vite 8.0.13 -> 8.0.16，ESLint 10.2.1 -> 10.4.1）由 Dependabot 自动完成，保持了工具链的更新。

**小结**：项目今日在 **WhatsApp 连接可靠性** 与 **代码类型安全** 方面取得了关键突破，同时补齐了技能搜索中缺失的指导信息。这些改动共同降低了底层故障率，并改善了最终用户的操作便利性。

## 4. 社区热点

- **[Issue #3159] 经常重复任务**  
  此 Bug 报告获得了较多关注。用户描述了一个清晰且严重的缺陷：当连续询问不同国家新闻时，AI 会在回答第二个问题时重复执行第一个任务。该问题直接影响了用户对 AI 任务调度逻辑的信任，预计将很快收到维护团队的响应。  
  链接: [Issue #3159](https://github.com/sipeed/picoclaw/issues/3159)

- **[PR #3161] fix(exec): keep deny patterns active for custom allow rules**  
  这是一项重要的安全修复，发现自定义允许规则会完全绕过拒绝规则。用户担心这会引发严重的安全风险。开发者积极响应并提交了补丁，当前该 PR（#3161）已开启并等待审核。  
  链接: [PR #3161](https://github.com/sipeed/picoclaw/pull/3161)

- **[Issue #3093] [Feature] I need SimpleX or tox**  
  该功能请求持续获得社区支持（👍: 1），用户明确表示需要支持 SimpleX、Wire 或 Tox 协议作为新的通信网关，反映出社区对去中心化、隐私优先通信通道有持续且明确的需求。  
  链接: [Issue #3093](https://github.com/sipeed/picoclaw/issues/3093)

## 5. Bug 与稳定性

| 严重程度 | Issue / PR | 摘要 | 状态 | 描述 |
| :--- | :--- | :--- | :--- | :--- |
| **严重** | **[#3159](https://github.com/sipeed/picoclaw/issues/3159)** | 任务重复执行（美国新闻+法国新闻案例） | 无修复 PR | AI 在处理连续、不同的请求时会错误地重新执行前一个任务的逻辑，导致计算资源浪费和输出结果错误。 |
| **高** | **[#3161](https://github.com/sipeed/picoclaw/pull/3161)** | 自定义允许规则绕过 `exec` 拒绝模式 | 🛜 **Fix PR 已提交** | 这是一个安全漏洞，攻击者可能利用放行的命令绕过拒绝列表。已提交修复方案。 |
| **中** | **[#3115](https://github.com/sipeed/picoclaw/pull/3115)** | 内联 data URL 导致的会话历史损坏 | 🛜 **Fix PR 待合并** | 通用工具输出中内嵌的 `data:image/...` 字符串被误当作媒体附件，导致历史记录损坏。修复方案已存在。 |
| **低** | **[#3128](https://github.com/sipeed/picoclaw/pull/3128)**, **[#3131](https://github.com/sipeed/picoclaw/pull/3131)** | 多处类型安全检查缺失 | 🛜 **Fix PR 待合并** | 涉及 HTTP 响应体关闭错误忽略及工具 Schema 类型断言，修复可减少静默错误和潜在 panic。 |

**分析**：项目稳定性正经历考验。#3159 是一个典型的“任务状态” Bug，尚未有修复 PR，建议维护者优先处理。#3161 的安全修复提交后，需尽快合并以避免潜在风险。

## 6. 功能请求与路线图信号

- **通信网关扩展** ([Issue #3093](https://github.com/sipeed/picoclaw/issues/3093))：用户需求强烈，要求支持 SimpleX 等去中心化协议。目前项目主要依赖 WebSocket 等协议，此需求若被采纳将拓展项目的连接生态。
- **Android ADB 远程工具** ([PR #3157](https://github.com/sipeed/picoclaw/pull/3157))：实验性功能，允许通过 ADB 操控 Android 设备。这显示出 PicoClaw 正从 AI 对话代理向**物理世界控制中枢**演变，未来可能进入路线图。
- **LLM 令牌用量追踪** ([PR #3156](https://github.com/sipeed/picoclaw/pull/3156))：新增了在消息中携带 Token 消耗的功能，这对于有成本敏感诉求的企业用户至关重要，很可能被合并到正式版本。
- **远程遥控模式** ([PR #3118](https://github.com/sipeed/picoclaw/pull/3118))：增强了 agent 命令的远程连接能力，支持 `--remote` 参数，使其能作为客户端连接到远程 Pico 实例。

**路线图信号**：以上功能请求和 PR 表明，项目正朝着**跨平台操控工具**与**企业级成本管理**两个方向演进，同时保持了对去中心化趋势的敏感。

## 7. 用户反馈摘要

- **痛点 — 任务执行异常** (Issue #3159)：一位使用 `deepseek-v4-flash-free` 模型和 Web UI 的用户反馈，AI 在处理多轮问答时存在严重的**任务重复执行**问题。用户期望的是一一对应的问答，但系统却出现了记忆混乱和行为异常。这反映出当前的任务调度器在处理“链式”请求时，状态追踪逻辑存在缺陷。
- **功能诉求 — 更广泛的连接协议** (Issue #3093)：用户 `Damian-o2` 明确表达了“我需要 SimpleX 或 Tox”，表明用户群体中有一部分是高度注重隐私和去中心化的开发者，他们希望 PicoClaw 能够作为这类自定义协议的局域网/广域网网关。

## 8. 待处理积压

以下为长期未响应或处于停滞状态的重要 Issue 和 PR，建议维护者重点关注：

1. **[PR #3118] [Stale] Add remote Pico WebSocket mode**  
   已停滞 11 天，涉及远程 agent 模式，扩展性强，但仍需确认设计与现有系统架构是否一致。  
   链接: [PR #3118](https://github.com/sipeed/picoclaw/pull/3118)

2. **[PR #3131] [Stale] fix(registry): add ok checks for tool schema type assertions**  
   已停滞 8 天，修复工具注册中的类型安全检查，是底层稳定性改进，应尽快审核合并。  
   链接: [PR #3131](https://github.com/sipeed/picoclaw/pull/3131)

3. **[PR #3104] [Stale] build(deps): bump shadcn**  
   依赖升级 PR，虽不紧急，但长期停滞可能导致安全漏洞，建议尽快合并。  
   链接: [PR #3104](https://github.com/sipeed/picoclaw/pull/3104)

---

**日报总结**：项目开发进展迅速，稳定性修复与功能扩展并重。Bug #3159 是当前最需要关注的用户体验问题。本周建议维护者优先合并安全补丁 (#3161) 及快速修复 #3159，并评审 @danmobot 提交的 ADB 及配置文件相关的 PR 集合。

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

好的，作为AI智能体与个人AI助手领域开源项目的分析师，以下是根据您提供的NanoClaw项目数据生成的2026年6月23日项目动态日报。

---

## NanoClaw 项目动态日报 (2026-06-23)

### 1. 今日速览

今日项目活跃度**高**，主要体现在密集的依赖升级与功能开发PR上。团队正在进行一次大规模的版本对齐工作，将核心Chat SDK及所有通道适配器统一升级至4.29.0版本（PR #2834, #2835, #2836）。同时，两项重要功能开发取得进展：Slack Socket Mode支持（PR #2837）与审批拒绝理由功能（PR #2832）。尽管今日无新Issue或版本发布，但10个PR的更新表明项目处于快速迭代期。

### 2. 版本发布

**无新版本发布。**

### 3. 项目进展

今日项目向前迈进的关键一步是完成了一次**核心依赖的版本对齐**。虽然尚未合并，但团队已提交三个关联PR，意图将 `main` 分支、`providers` 分支、`channels` 分支的 **Chat SDK 版本统一锁定至 `4.29.0`**。这解决了因版本不一致导致 `createChatSdkBridge(...)` 类型检查失败的问题，是提升代码健壮性的重要举措。

另一项已合并的PR (#2833) 是对贡献者指南和PR模板的规范化践行，这有助于提升社区贡献的质量和一致性。

### 4. 社区热点

今日讨论最活跃的PR主要围绕新功能开发：

1.  **feat(slack): Socket Mode — adapter + guided setup** ([PR #2837](https://github.com/nanocoai/nanoclaw/pull/2837)): 该PR为Slack集成增加了Socket Mode支持，允许Bot通过出站WebSocket连接，无需公网端点即可运行。这直接回应了**本地开发者和内网环境用户对简化Slack Bot部署的迫切需求**，是当前社区关注的热点。

2.  **feat(approvals): reject with reason** ([PR #2832](https://github.com/nanocoai/nanoclaw/pull/2832)): 该PR允许审批者在拒绝时附带理由，并将此信息反馈给请求Agent。这一改进显著提升了人机协作的体验，解决了“仅收到拒绝但不知原因”的痛点，预计将受到需要精细审批流程的用户广泛欢迎。

### 5. Bug 与稳定性

今日未报告新的严重Bug。但在稳定性方面有一项重要的修复PR待处理：

- **[严重] fix(setup): reap dead peer service registrations whose binary is gone** ([PR #2830](https://github.com/nanocoai/nanoclaw/pull/2830)): 该PR修复了一个系统层面的潜在问题。当删除NanoClaw安装目录而未运行卸载程序时，系统服务（launchd/systemd）会残留失效注册，持续尝试启动不存在的二进制文件。此修复将自动清除这些“僵尸”注册，**对系统长期稳定性和资源消耗至关重要**。

### 6. 功能请求与路线图信号

从今天的PR中可以清晰看到未来版本的几个方向：

- **增强连通性**：`feat(slack): Socket Mode` (PR #2837) 表明项目正在致力于降低集成门槛，特别是针对开发环境和受限网络。
- **提升协作体验**：`feat(approvals): reject with reason` (PR #2832) 指向了对Agent与人类审批者交互流程的精细化打磨，是提升Agent实用性的关键路径。
- **扩展生态**：`feat: add IMAP/SMTP email integration` ([PR #1235](https://github.com/nanocoai/nanoclaw/pull/1235)) 虽然创建于3月，但今日未关闭且有更新，表明邮件通道作为一项重要的“遗留系统”集成，仍在路线图规划中，可能成为下一版本的重点功能之一。

### 7. 用户反馈摘要

今日无直接的用户Issue讨论，但从相关PR的摘要中可以提炼出用户的潜在痛点：

- **“部署Slack Bot太麻烦”**：PR #2837 的诞生直接反映了部分用户希望摆脱公网端点限制的需求。在没有Socket Mode前，在本机或内网服务器开发测试Slack Bot存在障碍。
- **“Agent拒绝我的请求却不说原因”**：PR #2832 的动机正是为了解决用户在与Agent协作时反馈链路不完整的问题。用户期望获得更智能、有解释性的交互，而不是简单的“是/否”二元反馈。

### 8. 待处理积压

**功能请求积压**

- **IMAP/SMTP Email Integration** ([PR #1235](https://github.com/nanocoai/nanoclaw/pull/1235)): 该PR已开放超过3个月，但仍收到更新。这是一个重要的传统渠道集成，需求明确，建议维护团队评估其技术挑战及与现有架构的兼容性，以决定是否纳入近期里程碑。

**稳定性修复积压**

- **Agent Container `--shm-size` + `--init`** ([PR #2771](https://github.com/nanocoai/nanoclaw/pull/2771)): 该PR旨在解决Agent容器中Headless Chrome因Docker默认共享内存（/dev/shm）过小（64MB）而导致的崩溃问题。该问题对依赖浏览器自动化的Agent稳定性影响较大，建议合并优先级提升。

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw 项目动态日报 — 2026-06-23

---

## 1. 今日速览

- 项目过去 24 小时活跃度较低，无新 Issue 产生，唯一动态是 PR #968 仍处于待合并状态。
- 该 PR 修复了 Matrix 频道模块中 `next_batch` 光标仅存于内存，导致服务重启后触发初始同步的稳定性问题。
- 目前代码库在 `matrix` 相关逻辑上存在明确可复现的 Bug，但修复方案已经提交，整体项目处于“等待合并 → 测试验证”的稳步推进阶段。
- 无新版本发布，也无关闭的 Issue/PR，维护节奏偏向保守修复而非功能迭代。

---

## 2. 版本发布

**无**（过去 24 小时内无新 Release）。

---

## 3. 项目进展

### 待合并 PR（唯一重要变更）

| PR # | 标题 | 作者 | 状态 | 创建时间 |
|------|------|------|------|----------|
| [#968](https://github.com/nullclaw/nullclaw/pull/968) | fix(matrix): persist next_batch across restart + test env isolation | addadi | **待合并** | 2026-06-22 |

- **功能/修复**：将 Matrix 频道当前的 `/sync` 游标（`next_batch`）写入持久化存储，使 `buildSyncUrl` 在服务重启后能正确带上 `&since=` 参数，避免每次重启都执行完整初始同步。
- **附带改进**：PR 同时为相关单元测试添加了环境隔离，减少测试之间的状态污染。
- **项目进展评估**：该 PR 修复了一个中等严重度的数据一致性问题（重启后可能重复拉取历史消息或状态同步延迟），合并后将显著提升 Matrix 频道的健壮性。整体项目在“稳定性修复”维度向前迈出一步。

---

## 4. 社区热点

- 过去 24 小时内，唯一被更新的 PR [#968](https://github.com/nullclaw/nullclaw/pull/968) 未产生任何评论或反应（👍:0），也未进入讨论。  
- 目前项目社区活跃度较低，无高热度 Issue/PR 出现。维护者应留意该 PR 是否因缺乏 reviewer 或测试通过条件而停滞。

---

## 5. Bug 与稳定性

### 已知 Bug（严重程度：中等）

- **问题描述**：Matrix channel 的 `next_batch` 光标仅保存在内存结构体中，每当服务进程重启，`buildSyncUrl` 无法获取 `since` 参数，导致 homeserver 返回完整初始同步。`pollMessages` 虽然能检测到初始同步并捕获新 token，但此行为会引发：
  - 不必要的带宽和计算消耗（重复同步大量历史消息）
  - 可能触发重复通知或状态错乱（取决于下游处理逻辑）
- **修复 PR**：✅ 已提交 [#968](https://github.com/nullclaw/nullclaw/pull/968)（待合并）

未报告其他崩溃、回归或安全问题。

---

## 6. 功能请求与路线图信号

- 今日无新增功能请求 Issue。
- 从 PR #968 的“test env isolation”部分看，开发者在测试基础设施上投入了精力，这可能暗示团队正在推进 CI/测试稳定性，为后续功能迭代铺路。但该信号较弱，未直接关联新功能路线图。

---

## 7. 用户反馈摘要

- 过去 24 小时无用户评论或反馈。
- 根据 PR #968 的问题描述（Problem），可推断真实用户场景：  
  **痛点**：Matrix 频道在部署或扩缩容重启后，会反复执行初始同步，导致消息重复、同步延迟或资源浪费。用户可能观察到消息条数异常或通知滞后。  
  **预期**：重启时应能从持久化状态恢复同步断点，实现无缝恢复。

---

## 8. 待处理积压

- 当前无长期未响应的重要 Issue 或 PR。唯一待合并的 PR [#968](https://github.com/nullclaw/nullclaw/pull/968) 距离创建已过 24 小时，尚未有 reviewer 标记或合并动作，建议维护者尽快安排 code review，避免修复陷入积压。

---

**总结**：项目今日处于低活跃修复窗口，核心进展是 Matrix 稳定性 Bug 的修复等待合并。整体健康度中等，社区参与度低，建议维护者主动推动 PR 合并并关注后续测试反馈。

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

好的，以下是基于您提供的 IronClaw (github.com/nearai/ironclaw) GitHub 数据生成的 2026-06-23 项目动态日报。

---

# IronClaw 项目日报 2026-06-23

## 1. 今日速览

今日项目活动度极高，共有 **10 个 Issues 更新**和 **33 个 PR 更新**，表明团队正在快速推进功能开发和问题修复。核心开发工作围绕 **Reborn** 子系统的自动化管理（暂停/恢复/删除）、工具权限系统与用户界面集成展开。然而，一个严重的**回归问题**（#5139）导致大量 PinchBench 任务初始化失败，凸显了近期合并变更的稳定性风险。同时，一个用以解决关键性能瓶颈的 **PR #5149**（上下文管理与渐进式工具披露）被提出，有望显著降低 Token 消耗和请求超时。

## 2. 版本发布

今日无新版本发布。

## 3. 项目进展

今日合并/关闭了 **12 个 PR**，项目在以下方面取得了显著进展：

- **核心稳定性与性能修复**:
    - **[CLOSED] #5142: fix(turns): prevent turn-state write convoy**: 移除了 per-user 的写入门控，转而采用带版本控制的 CAS 操作与退避重试，旨在解决高并发下的写锁竞争问题。
    - **[CLOSED] #5141: fix(triggers): complete once permanent failures**: 修复了 `Once` 类型触发器在永久性失败后无限重试的问题，确保此类任务失败后能正确终止。
- **自动化与权限系统增强**:
    - **[CLOSED] #5061: feat(reborn): skill extraction & self-evolution with activation controls**: 合并了 Hermes 风格的技能提取与自我进化功能，允许 Reborn 在成功完成任务后自动从对话记录中提炼技能。
    - **[CLOSED] #4959 & #4958**: 这两个关于 **全局自动批准设置** 和 **per-tool 权限模型** 的问题被关闭，对应的 UI 实现 PR #5068 正在开发中，表明其底层设计已达成共识并完成。
- **基础设施与配置**:
    - **[CLOSED] #5150: restore GSuite duplicate account fallback**: 修复了 GSuite 账户解析器的一个回归问题，恢复了重复账户的回退选择逻辑。

整体而言，项目正向**更稳定、更自动化、用户可控性更强**的方向迈进。重要基础功能（如技能进化、Triggers 状态机修复）已落地，为后续迭代打下基础。

## 4. 社区热点

今日社区讨论焦点集中在两个问题上：

1.  **回归问题与性能**:
    - **[ISSUE #5139] [OPEN] reborn regression: web/research tasks hang at init (0 LLM calls)**: 这是目前最严重的问题，直指近期 10 个提交（`704fcd43`）引发了回归，导致 Web/研究任务在初始化阶段即挂起，影响面很大（波及 21/147 的 PinchBench 任务）。此问题立即引发了开发者关注，与 PR #5149 提出的“上下文过大导致超时”的分析方向一致。
        - 链接: <https://github.com/nearai/ironclaw/issues/5139>

2.  **模型兼容性与自动化**:
    - **[ISSUE #5151] [OPEN] Claude fails to create Reborn automation after trigger pause/resume tools are exposed**: 用户 `italic-jinxin` 报告了在与 `anthropic/claude-sonnet-4-5` 模型交互时，Reborn 无法按指令创建自动化任务的问题。模型错误地调用了 `capability_info` 等探索性工具而非 `trigger_create`。这反映了不同模型对新功能的适配差异，是 AI Agent 应用中常见的痛点。
        - 链接: <https://github.com/nearai/ironclaw/issues/5151>

## 5. Bug 与稳定性

今日报告的 Bug 按严重程度排列如下：

- **严重 (P0)**:
    - **[(ISSUE #5139) Reborn 回归问题：Web/研究任务在初始化时挂起](https://github.com/nearai/ironclaw/issues/5139)**: 影响大量任务执行，无任何 LLM 调用，是当前最高优先级问题。**尚无 Fix PR**，但 PR #5149 可能是其解决方案的前奏。
- **高 (P1)**:
    - **[(ISSUE #5148) Turn调度器心跳与状态更新锁导致死锁](https://github.com/nearai/ironclaw/issues/5148)**: 在特定流程（如 GitHub 扩展安装）中，调度器心跳触发时若持有状态锁，会导致任务永久挂起。**尚无 Fix PR**。
    - **[(ISSUE #5151) Claude 模型无法创建自动化任务](https://github.com/nearai/ironclaw/issues/5151)**: 特定模型与新功能的兼容性问题，影响用户体验。**尚无 Fix PR**。
- **中 (P2)**:
    - **[(ISSUE #5147) 波动测试: trigger_poller 测试间歇性失败](https://github.com/nearai/ironclaw/issues/5147)**: 阻碍了 PR 合并队列，约 1/3 的失败率。**尚无 Fix PR**。
    - **[(ISSUE #5146) 扩展页面缺少“停用”按钮](https://github.com/nearai/ironclaw/issues/5146)**: WebUI 功能缺失。**尚无 Fix PR**。
- **低/持续 (P3)**:
    - **[(ISSUE #4108) Nightly E2E 持续失败](https://github.com/nearai/ironclaw/issues/4108)**: 长期存在的 E2E 测试问题，今日再次报告失败。

## 6. 功能请求与路线图信号

今日收到的功能请求和路线图信号包括：

- **明确纳入路线图**:
    - **工具渐进式披露 (PR #5149)**: 该 PR 直接针对性能瓶颈，通过只发送当前步骤相关工具来大幅减少 Token 消耗。这是一个**高优先级**的性能优化，很可能被合并到下个版本。
    - **WebUI 集成自动化管理 (PR #5131, #5133)**: 自动化的暂停/恢复/删除功能已有成熟的 PR 在审查中，是下一版本路线图中的核心功能。
    - **用户权限管理 UI (PR #5068)**: 与已关闭的 #4958、#4959 配套的 WebUI，是用户非常期待的功能，预计会很快被合并。
- **潜在新功能**:
    - **[(ISSUE #5144) 在 Provider 卡片中显示 NEAR AI 默认 Base URL](https://github.com/nearai/ironclaw/issues/5144)**: UI 可用性改进建议，关联到更大的 #5119 父 Issue。
    - **[(ISSUE #5146) 添加停用扩展的按钮](https://github.com/nearai/ironclaw/issues/5146)**: 基础的扩展管理功能请求，用户期望在界面上直接操作。

## 7. 用户反馈摘要

从今日的 Issues 中，可以提炼出以下用户痛点和使用场景：

- **稳定性是核心痛点**: #5139 和 #5148 直接反映了用户（或内部QA）在核心工作流中遇到的卡死和超时问题，严重影响了任务执行效率和可靠性。用户 `pranavraja99` 通过 PinchBench 测试量化了回归的影响范围。
- **模型行为非预期导致信任危机**: #5151 展示了当 AI Agent 无法正确理解用户意图并执行操作时，会严重损害用户的使用信心。用户期望模型能准确调用暴露出的工具，而不是进行“多此一举”的探索。
- **配置与管理的易用性需求**: #5144 中“None”字段带来的困惑和 #5146 中缺少停用按钮，都表明用户希望 WebUI 能提供更直观、完整的信息和操作入口，减少命令行或文档查阅的依赖。
- **自动化期望**: #5151 的用户明确期望创建一个**周期性自动化任务**，这表明 Reborn 的自动化（Triggers）功能是用户的核心使用场景之一，其可靠性至关重要。

## 8. 待处理积压

以下 Issue/PR 长期存在，可能需要维护者关注：

- **[(ISSUE #4108) Nightly E2E 持续失败](https://github.com/nearai/ironclaw/issues/4108)**: 自 2026-05-27 起已存在近一个月，一直由 `github-actions[bot]` 自动报告失败。虽然未分配优先处理，但长期失败的 E2E 测试会降低 CI/CD 的可靠性，并可能掩盖其他回归问题。**建议评估其根因并决定是否需要重置基线或修复。**
- **[(PR #2863) feat(llm): add Manifest as built-in LLM provider](https://github.com/nearai/ironclaw/pull/2863)**: 该 PR 由新贡献者 `SebConejo` 于两个月前提交，目前仍为开放状态。此功能可扩展 Reborn 的 LLM 提供商支持范围，增加平台的多样性和厂商中立性。**建议维护者评估其技术可行性并给予反馈，以避免贡献者流失。**

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

好的，以下是根据你提供的 LobsterAI GitHub 数据生成的 **2026-06-23 项目动态日报**。

---

# LobsterAI 项目日报 | 2026-06-23

## 1. 今日速览

- 项目今日有 **3 个 PR 被合并/关闭**（#2190、#2189、#2188），主要涉及 OpenClaw 的 Cron 存储迁移和会话同步，维护工作持续推进。
- 遗留的 **6 个开放 PR 仍处于 stale 状态**（超过 80 天无更新），其中包含安全修复、UI 优化等重要内容，合并进度缓慢。
- **唯一活跃的 Issue #1400 报告了严重的 4.1 版本升级故障**（网关反复重启），社区用户已留联系方式求助，但尚未出现 fix PR 或官方回应。
- 整体活跃度偏低，核心维护者仅在 OpenClaw 领域有少量活动，积压的 bug 和功能改进尚未得到有效处理。

## 2. 版本发布

今日无新版本发布。

## 3. 项目进展

今日合并/关闭的 PR 均聚焦于 **OpenClaw 的定时任务（Cron）系统**：

- **PR #2190** `fix(openclaw): sync cron run sessions`  
  修复 OpenClaw 定时任务的运行会话同步问题，确保重复运行使用本地 Cowork 会话，并标准化缓存键。  
  [链接](https://github.com/netease-youdao/LobsterAI/pull/2190)

- **PR #2189** `fix(openclaw): migrate legacy cron storage on startup`  
  在网关启动时自动检测并迁移旧版 OpenClaw Cron JSON 存储，执行官方 doctor 迁移，并记录文档。  
  [链接](https://github.com/netease-youdao/LobsterAI/pull/2189)

- **PR #2188** `Liuzhq/rlog`  
  关联多个区域（renderer, docs, main, openclaw, cowork），具体变更未详细描述，但已合并。  
  [链接](https://github.com/netease-youdao/LobsterAI/pull/2188)

以上三个 PR 共同推进了 OpenClaw 定时任务模块的稳定性与兼容性，为后续版本升级做了铺垫。

## 4. 社区热点

今日最受关注的 Issue 是：

- **#1400** `4.1版本严重bug，网关反复启动失败，反复重启，无限循环！`  
  创建于 2026-04-03，今日有更新。作者来自社区，详细描述了从 3.30 升级到 4.1 后的完全瘫痪问题，并主动提供了联系方式。评论数达 6 条，是当前社区反应最集中的问题。  
  [链接](https://github.com/netease-youdao/LobsterAI/issues/1400)

该 Issue 暴露出版本升级可能存在的关键兼容性缺陷，用户期望团队尽快介入排查。截至目前尚无官方回复或关联 PR。

## 5. Bug 与稳定性

按严重程度排列：

| 严重程度 | 编号 | 标题 | 状态 | 是否有 fix PR |
|----------|------|------|------|---------------|
| **P0（严重）** | #1400 | 4.1版本严重bug，网关反复重启 | OPEN（stale） | 无 |
| P2（功能性） | #1402 | 多选附件只保留最后一个文件 | OPEN（stale） | ✅ #1402 已提交，但未合并 |
| P3（安全性） | #1401 | SSE 请求 ID 使用伪随机数，可被预测 | OPEN（stale） | ✅ #1401 已提交，但未合并 |
| P3（本地化） | #1403 | 中文界面下“delete”未翻译 | OPEN（stale） | ✅ #1403 已提交，但未合并 |
| P3（UI） | #1406 | 通知渠道列表在 IM 未开启时为空 | OPEN（stale） | ✅ #1406 已提交，但未合并 |

**关键风险**：#1400 高严重性问题尚无任何解决迹象，可能导致用户流失。安全 PR #1401 积压 80 余天未合并，构成潜在攻击面。

## 6. 功能请求与路线图信号

- **PR #1404** `feat(scheduledTasks): 定时任务创建界面时间控件优化`  
  提出使用自定义组件替代原生控件，统一 Electron 环境下的交互和主题样式。该 PR 已 stale 但表明社区对定时任务 UI 体验有改进需求。  
  [链接](https://github.com/netease-youdao/LobsterAI/pull/1404)

- **PR #1277** `chore(deps-dev): bump the electron group`  
  Dependabot 自动发起的 Electron 版本升级（40.2.1 → 42.4.0），涉及大版本跳跃，若合并将提升安全性与性能。该 PR 已 stale，反映项目对依赖更新的响应滞后。  
  [链接](https://github.com/netease-youdao/LobsterAI/pull/1277)

从当前动态看，OpenClaw 定时任务和 Electron 升级可能是近期路线图的一部分，但具体时间未知。

## 7. 用户反馈摘要

来自 Issue #1400：

> “从3.30版本升级到4.1版本后反复重启，始终无法正常启动！……另外升级前还有一个bug：自定义配置的LLM - qwen3.5-plus 无法调用……反正现在是彻底瘫痪了！请团队协助解决，我的联系方式……”

**用户痛点**：
1. **升级可靠性**：3.30 → 4.1 升级路径存在严重阻塞，无平滑迁移方案。
2. **自定义 LLM 配置冲突**：登录后自动配置可能与手动设置冲突，导致 `web-extractor` 无法启动。
3. **紧急响应缺失**：用户主动留联系方式，但过去 80 天未获官方回复，社区信任度受损。

## 8. 待处理积压

以下 Issue/PR 已超 **80 天**无有效响应或合并，建议维护者优先关注：

| 编号 | 类型 | 标题 | 最后更新 | 备注 |
|------|------|------|----------|------|
| #1400 | Issue | 4.1版本严重bug，网关反复重启 | 2026-06-23 | 高严重性，用户急需支持 |
| #1401 | PR | fix: 修复请求安全性问题 | 2026-06-23 | 安全修复，不可忽视 |
| #1402 | PR | fix(cowork): 多选附件只保留最后一个 | 2026-06-23 | 功能回归修复 |
| #1403 | PR | fix(i18n): 添加delete翻译键 | 2026-06-23 | 本地化质量 |
| #1404 | PR | feat(scheduledTasks): 定时任务界面优化 | 2026-06-23 | 用户体验改进 |
| #1406 | PR | fix(scheduled-task): 通知渠道列表回退 | 2026-06-23 | 功能性修复 |
| #1277 | PR | chore: 升级Electron依赖 | 2026-06-22 | 依赖安全与兼容性 |

**建议**：在下一个版本迭代前，至少合并 #1401（安全）、#1402（功能修复）和 #1277（依赖升级），并针对 #1400 发布热修复。

---

*报告生成时间：2026-06-23 23:59 UTC*  
*数据来源：[netease-youdao/LobsterAI](https://github.com/netease-youdao/LobsterAI)*

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

好的，没问题。作为AI智能体与个人AI助手领域的开源项目分析师，以下是为您生成的Moltis项目2026年6月23日动态日报。

---

### Moltis 项目日常动态日报 | 2026-06-23

#### 1. 今日速览

今日项目整体活跃度较低，过去24小时内无新增Issue或新版本发布。项目主要动态为一项始于今年2月的功能PR (#215) 最终被合并关闭，为频道图像传递能力带来了关键性支持。这表明项目正在稳步推进基础功能完善，但社区公共讨论热度较低，目前处于温和的整合期。

#### 2. 版本发布

N/A

#### 3. 项目进展

今日项目最重要的进展是合并了 **PR #215**，这是一个里程碑式的功能补全。

-   **#215 [已合并] feat(tools): add send_image tool for channel image delivery**
    - **状态**：已合并 (由最大化者合并)
    - **概要**：此PR引入了一个名为 `send_image` 的新工具。该工具允许AI技能向Telegram等聊天频道发送本地图像文件（支持PNG, JPEG, GIF, WebP格式）。其设计精妙之处在于复用了项目现有的截图管线，无需重复造轮子。它会在返回的 `screenshot` 键中包含一个 `data:` URI，聊天运行器会自动拾取并传输。同时，该工具还支持可选的 `caption` 参数，允许为图像添加描述文字。
    - **项目意义**：这是对Moltis跨渠道通信能力的重大增强。之前可能仅限于文本或通过复杂代码实现的图像发送，现在通过一个标准化的 `send_image` 工具即可实现，降低了开发者和用户构建多媒体交互体验的门槛。项目在“工具-渠道”连接能力上向前迈出了一大步。

#### 4. 社区热点

当前无特别活跃的Discussion或Issue。今日唯一合并的PR #215的评论数为undefined（可能为0），表明该PR的讨论和合并在内部完成，并未引发大规模社区公开讨论。社区整体处于静默状态。

#### 5. Bug 与稳定性

过去24小时内无新提交的Bug报告。

#### 6. 功能请求与路线图信号

今日无新的功能请求提交。但 **PR #215** (send_image工具) 的成功合并，强烈暗示了项目路线图的下一个方向：

-   **核心体验优化**：提升AI与用户交互的丰富度，从纯文本向“图文混合”甚至“多媒体”交互演进。
-   **标准化工具库**：`send_image` 是 `screenshot` 管线的复用，表明项目正致力于构建一个模块化、可复用的工具标准库。未来类似功能的开发（如 `send_audio`、`send_file`）可能会遵循相同的技术路径。

因此，可以推测在未来一到两个版本中，图像相关的场景（如AI生成图像后直接发送、AI解读并返回图表等）将成为核心用例。

#### 7. 用户反馈摘要

今日无有效的用户反馈提交。但从 **PR #215** 的功能描述可以推断出用户场景中的潜在痛点：

-   **场景**：用户希望AI助手能直接分享一张本地的分析图表、截图或记忆中的照片给Telegram群组。
-   **痛点**：之前要实现这个功能，用户可能需要手动上传或依赖外部服务，流程割裂。新增的 `send_image` 工具将这一系列操作封装成一个简单的工具调用，直接提升了用户体验的连贯性。

#### 8. 待处理积压

根据现有数据，当前无 “待处理积压” 的Issues或PRs。PR #215 已成功合并且没有新的积压问题。需维持对 **PR #180**（或其他未关闭的早期PR）的持续关注，以确保项目进展的连续性。

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，我已根据您提供的 CoPaw（github.com/agentscope-ai/CoPaw）GitHub 数据，生成以下项目动态日报。

---

# CoPaw 项目动态日报 | 2026-06-23

---

## 1. 今日速览

项目今日迎来 **v1.1.12.post2 补丁版本**，展现出极高的开发活跃度，24小时内共有50条PR和14条Issues更新。开发团队修复了多个与**定时任务挂起**、**Dream任务执行失败**以及**第三方模型兼容性**相关的核心Bug，显著提升了系统的稳定性。社区对**移动端支持**和**Agent思考过程卡死**问题的讨论热度很高，这反映了用户对更流畅、更可靠的使用体验的追求。整体来看，项目正处于快速迭代、积极解决用户痛点的阶段。

## 2. 版本发布

### v1.1.12.post2 (Post 版本)
- **发布时间**: 2026-06-23
- **主要更新内容**:
    - **Bug 修复**: 修复了删除当前会话后无法导航到新聊天的问题。
    - **功能增强**: 增强了文件预览功能，支持显示相对路径。
    - **稳定性**: 修复了 `dream` 任务因触发时间错失而报错的问题，并将默认 `misfire_grace_time` 设置为 600 秒。
    - **破坏性变更**: 无。
    - **迁移注意**: 建议所有 v1.1.12 系列的用户升级到此版本，以解决部分定时任务和 Dream 任务执行失败的问题。升级方式为 `pip install --upgrade qwenpaw==1.1.12.post2`。

## 3. 项目进展

过去24小时内，项目取得了重要进展，大量核心功能和稳定性修复被合并到主分支：

- **即时通知与UI流畅度**：
    - [#5331](https://github.com/agentscope-ai/QwenPaw/pull/5331) **实时SSE推送通知**：将控制台通知（如cron、API）从轮询（1-3秒延迟）改为SSE推送（<50毫秒），并支持可选的音频提示音，极大提升了通知的实时性。
    - [#5326](https://github.com/agentscope-ai/QwenPaw/pull/5326) **最小化到系统托盘**：Tauri客户端现在关闭窗口时默认隐藏到系统托盘，而非退出应用，改善了常驻用户的使用体验。

- **核心功能增强**：
    - [#5325](https://github.com/agentscope-ai/QwenPaw/pull/5325) **记忆搜索新近度排序**：为 `memory_search` 功能增加了可选的“新近度感知”排序，使最近的记忆在搜索结果中排名更高，增强了核心记忆功能的智能性。
    - [#5304](https://github.com/agentscope-ai/QwenPaw/pull/5304) **终端编码模式**：引入 `qwenpaw terminal` CLI命令，提供了一个连接到后台守护进程的交互式编码终端，方便开发者在终端环境中直接与Agent交互。
    - [#5323](https://github.com/agentscope-ai/QwenPaw/pull/5323) **多步骤任务进度面板**：为 Agent 的多步骤任务（Plan）添加了原生 `todo_write` 工具和前端进度面板，用户可在控制台实时查看任务执行进度。

- **修复回归与兼容性问题**：
    - [#5335](https://github.com/agentscope-ai/QwenPaw/pull/5335) **修复UI卡死**：修复了后端模型执行异常时，SSE连接关闭但未返回错误事件，导致前端UI卡死的问题。
    - [#5339](https://github.com/agentscope-ai/QwenPaw/pull/5339) **修复智普AI连接测试**：修复了智谱AI provider的模型连接测试失败问题，原因是消息格式与OpenAI API不兼容。
    - [#5324](https://github.com/agentscope-ai/QwenPaw/pull/5324) **修复文件预览**：修复了 `send_file_to_user` 发送图片时被强制下载而非预览的回归问题。

## 4. 社区热点

- **Issue #5064 - 定时任务无法触发**: 虽然被标记为 `[invalid]` 并已关闭，但收获了12条评论，是今日讨论最活跃的议题。用户反映Agent创建的定时任务无法在设定时间自动执行，且无法手动编辑，严重影响了核心功能的可用性。此问题与 #5235、#5398 高度相关，共同揭示了定时任务模块的不稳定性，是社区最关注的痛点之一。
- **Issue #5345 - 自定义OpenAI兼容提供商不支持Function Calling**: 该问题收到了6条评论。社区用户尝试集成OMLX等第三方服务时发现，尽管这些服务提供了OpenAI兼容API，但Agent无法进行工具调用。这说明用户对模型和服务的选择范围有很高要求，而当前的兼容层未能完美支持所有服务商。
- **PR #5394 - 移动端卡片布局**: 这是一个关于移动端UI响应式设计的PR，引发了广泛关注。用户普遍反映在手机上访问WebUI体验不佳，社区对于“为Agent提供移动端访问入口”的呼声很高，该PR直接回应了这一需求。

## 5. Bug 与稳定性

| 严重程度 | Issue ID | 问题描述 | 状态 | 关联修复 PR |
| :--- | :--- | :--- | :--- | :--- |
| **严重** | [#5210](https://github.com/agentscope-ai/QwenPaw/issues/5210) | Cron定时任务停止调度，应用存活但任务不执行。 | **已关闭** | 已有PR [#5347](https://github.com/agentscope-ai/QwenPaw/pull/5347) 进行根因修复。 |
| **严重** | [#5402](https://github.com/agentscope-ai/QwenPaw/issues/5402) | Dream任务执行失败，多个Agent同时报错。 | **已关闭** | [#5426](https://github.com/agentscope-ai/QwenPaw/pull/5426) |
| **中等** | [#5328](https://github.com/agentscope-ai/QwenPaw/issues/5328) | 使用DeepSeek时，Agent在思考过程经常卡死。 | **开启** | 暂无 |
| **中等** | [#5421](https://github.com/agentscope-ai/QwenPaw/issues/5421) | 在Agent和聊天窗口之间切换时出现严重卡顿。 | **开启** | 暂无 |
| **低** | [#5379](https://github.com/agentscope-ai/QwenPaw/issues/5379) | Python命令安装后启动报内部服务器错误。 | **开启** | 暂无 |

## 6. 功能请求与路线图信号

- **第三方模型集成增强**：Issues #5345 和 #5427 表明社区对更广泛的模型兼容性需求强烈。用户不仅希望支持标准的 OpenAI API格式，还希望支持 Anthropic 格式等，这暗示项目可能需要一个更灵活的 Provider 抽象层。
- **移动端体验优先**：Issue #4635 虽然已关闭，但与之相关的 PR #5394（移动端卡片布局）和 PR #5367/#5431（安全页面自适应）正在积极开发。这表明项目路线图可能已将“提升移动端WebUI体验”作为下一阶段重点目标。
- **CLI 增强**：已合并的 PR #5304（终端编码模式）和正在开发的 PR #5210（`cron update`命令）表明，项目正在持续强化命令行工具，使其成为一个更强大的管理终端。
- **桌面端自动化**：新开的 PR #5428 为桌面版本增加了端到端UI验证CI，这预示着项目可能计划在下一版本（2.0）中发布更稳定的桌面客户端。

## 7. 用户反馈摘要

- **正面反馈（隐含）**：定时任务功能的“自动生成并展示”体验良好（#5064），表明基础交互逻辑是清晰的。
- **核心痛点**：
    - **定时任务不稳定**: 用户反复报告Agent创建的定时任务“无法正常触发”、“停止调度”、“执行失败”，这是当前最令人沮丧的体验。社区对此功能的可靠性期待很高，因为它关系到“Agent自主执行”的承诺。
    - **思考与输出不一致**: 多位用户反映，DeepSeek等模型的思考内容（`thinking`）与最终回复（`content`）处理不当，导致用户看不到回复或界面卡死（#5328, #5416）。这说明在处理非标准API响应格式时存在缺陷。
    - **部署与集成困难**: 用户报告了自定义Provider集成失败（#5345, #5427）和安装后“服务器内部错误”（#5379）等问题，入门门槛仍然较高，影响了潜在用户的转化。

## 8. 待处理积压

- **[PR #5153](https://github.com/agentscope-ai/QwenPaw/pull/5153) (开启, 11天) - 启动优化同步**: 将Tauri客户端的“瞬时启动”优化移植到pywebview客户端。该PR已打开11天，未获得合并或明确的下个版本计划。鉴于桌面体验是项目的关键入口之一，建议维护者评估并推进此PR。
- **[PR #4669](https://github.com/agentscope-ai/QwenPaw/pull/4669) (开启, 29天) - Tauri自动更新**: 添加桌面客户端自动更新功能的PR已近一个月未合并。自动更新是桌面应用维护的基本功能，它的长期积压可能使用户无法及时获得最新的bug修复和安全更新，建议优先处理。
- **近期Bug的高密度复现**: Issues #5235 (cron不执行)、#5064 (cron无法触发) 和 #5398 (cron停止调度) 三者在短时间内被不同用户以不同场景报告，虽然部分已标为 `[invalid]` 或已关闭，但问题根因可能并未完全解决。PR #5347 正在尝试做根本性的修复，建议维护者密切关注其合并状态和后续测试结果，防止问题再次复发。

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，作为AI智能体与个人AI助手领域开源项目分析师，根据您提供的ZeroClaw项目数据，以下是2026年6月23日的项目动态日报。

---

## ZeroClaw 项目日报 - 2026年6月23日

### 1. 今日速览

ZeroClaw 项目今日活跃度 **极高**。单日 **50 条 PR** 更新（其中43条待合并）是目前最显著的信号，表明代码库正经历一轮密集的功能合入与修复冲刺。社区贡献者非常活跃，主要集中在**可观测性**、**运行时稳定性**和**提供者兼容性**三大领域。一个涉及Groq提供商的严重Bug（#8219）已于当日定位并修复（PR #8227），体现了项目的快速响应能力。

### 3. 项目进展

过去24小时，共有 **7 条PR** 被合并或关闭，项目在多个关键领域取得了实质性进展：

- **修复 Groq 提供商兼容性 (Bug 修复):** [PR #8227](https://github.com/zeroclaw-labs/zeroclaw/pull/8227) 被合并，针对性地修复了 `gpt-oss-120b` 模型在 Groq 上执行多轮工具调用时因序列化问题导致的崩溃。此举迅速响应了当日报告的严重Bug（#8219），保障了核心工作流的稳定性。
- **增强测试与稳定性:** [PR #8220](https://github.com/zeroclaw-labs/zeroclaw/pull/8220) 被合并，新增了针对通道层（channels）历史修剪（proactive trim）时JSON工具调用ID（`tool_call_id`）保留的回归测试，强化了边缘场景下的数据一致性。
- **基础设施与CI优化 (关闭):** [Issue #8143](https://github.com/zeroclaw-labs/zeroclaw/issues/8143) 已被关闭，标志着 `.po` 翻译文件迁移至独立Git子模块的任务完成，此举将有效精简主仓库大小并提升CI管理效率。

### 4. 社区热点

- **Groq 提供者兼容性问题 ([Issue #8219](https://github.com/zeroclaw-labs/zeroclaw/issues/8219)):** 此Bug报告是本日的焦点。它直接影响了使用Groq平台运行特定模型的用户，导致多轮对话中断。问题暴露了两个独立但耦合的缺陷（`tool_call_id`序列化错误和`reasoning_content`被拒绝）。该问题的快速定位和修复（PR #8227）获得了社区积极关注，展示了项目维护的敏捷性。
- **增强可观测性讨论 ([PR #8065](https://github.com/zeroclaw-labs/zeroclaw/pull/8065), [#8066](https://github.com/zeroclaw-labs/zeroclaw/pull/8066)):** 这两条关于日志关联（trace_id）和LLM请求负载捕获（opt-in）的 PR 虽然创建于几日前，但在24小时内仍持续有更新，表明社区对“可观测性”的浓厚兴趣。贡献者与维护者正围绕如何在保证隐私（默认关闭）与增强调试能力之间取得平衡进行深入讨论。

### 5. Bug 与稳定性

今日新报告 **1 个** Bug，并有 **1 个** 长期运行的Bug待办事项被关注。

- **[严重] Groq 多轮工具调用失败 (`tool_call_id` / `reasoning_content`):** [Issue #8219](https://github.com/zeroclaw-labs/zeroclaw/issues/8219)
    - **描述:** 使用 `gpt-oss-120b` 模型在Groq上运行时，第二次请求会因 `tool_call_id` 为 null 和 `reasoning_content` 被拒绝而失败。这是一个会中断核心智能体循环的严重问题。
    - **状态:** 🛠️ **已修复！** 相应的修复PR [[#8227]](https://github.com/zeroclaw-labs/zeroclaw/pull/8227) 已于当日被合并至主分支，解决了此问题。
- **[高] 路径列举工具输出误触视觉路由 ([PR #7345](https://github.com/zeroclaw-labs/zeroclaw/pull/7345)):**
    - **描述:** 一个存在了16天的Bug PR。`content_search` 等路径列举工具在结果中输出本地图片路径，导致智能体循环错误地将其识别为需要视觉模型处理，从而触发不必要的视觉路由。
    - **状态:** 🔄 **待合并**。该PR仍在开放状态，是影响运行时准确性的一个潜在问题。

### 6. 功能请求与路线图信号

- **[高优先级] 支持为每个 Agent 配置自定义环境变量 ([Issue #8226](https://github.com/zeroclaw-labs/zeroclaw/issues/8226)):**
    - **描述:** 用户 `susyabashti` 提出了一个清晰的功能需求：为 `AliasedAgentConfig` 增加声明式 `env` 配置块，允许开发者为不同Agent安全地注入独立的环境变量。
    - **信号分析:** 这是一个合理的、面向开发者体验和安全性的增强需求。结合近期活跃的PR（如[#8066](https://github.com/zeroclaw-labs/zeroclaw/pull/8066)对配置项的增强），该功能有很大概率纳入下一版本的路线图。
- **应用内升级功能 ([PR #8173](https://github.com/zeroclaw-labs/zeroclaw/pull/8173)):**
    - **描述:** 实现了RFC #8170中描述的仪表板升级功能，允许用户通过点击侧边栏版本标签进行一键升级并重启。
    - **信号分析:** 这是一个以用户为中心的重大功能（`size: L`），表明项目正从底层稳定性向用户界面和易用性进发，提升端到端用户体验。
- **SOP 运行状态持久化存储 ([PR #8206](https://github.com/zeroclaw-labs/zeroclaw/pull/8206)):**
    - **描述:** 增加了支持WAL模式的SQLite存储后端，用于持久化SOP运行状态，并配套了实时运行指标。
    - **信号分析:** 这是对SOP（标准操作程序）功能的重大增强，使其更具实用性。释放信号表明ZeroClaw正在构建更健壮、可审计的工作流执行能力。

### 7. 用户反馈摘要

从今日的Issue和PR中，可以提炼出以下用户痛点与诉求：

- **痛点：** 用户 `perlowja` 在使用 Groq 的 `gpt-oss-120b` 模型时遭遇了**核心工作流中断**，暴露出项目在支持非标准AI提供者（特别是遵循OpenAI兼容协议但具有特殊约束的平台）时可能存在的兼容性问题。用户期望的是“开箱即用”的体验。
- **诉求：** 用户 `susyabashti` 提出了一个清晰的 **“为开发者赋能”** 的诉求。他们不希望为了给不同Agent设置环境变量而去修改全局配置文件或依赖复杂的Inject机制，而是期望一个直接在Agent配置中声明式定义的方式，体现了对**模块化和隔离性**的高要求。
- **持续关注：** 从对可观测性PR（#8065, #8066）的持续更新可以看出，社区工程师普遍有**运维和调试的深层需求**，他们不仅关心“结果正确”，更关心“为何如此”以及“成本如何”。

### 8. 待处理积压

以下为长期未处理或需要维护者关注的重要项目：

- **[关键] NVIDIA NIM 提供者视觉支持 ([PR #8100](https://github.com/zeroclaw-labs/zeroclaw/pull/8100)):**
    - **状态:** 开放，已持续 **2** 天。
    - **影响:** 修复了NVIDIA NIM提供者硬编码禁用视觉支持的问题，是支持特定GPU推理平台的关键。
    - **建议:** 优先合并，以确保对更多硬件/平台的支持。
- **[重要] 智能体生命周期观察者事件传播不完整 ([PR #7771](https://github.com/zeroclaw-labs/zeroclaw/pull/7771)):**
    - **状态:** 开放，已持续 **7** 天。
    - **影响:** 涉及可观测性基座，`channel`, `agent_alias`等关键元数据在大部分路径上未被正确传播，导致追踪数据不完整。
    - **建议:** 作为可观测性增强的前置条件，需要尽快推进审查与合并。
- **[长期] WebSocket通道的活跃连接管理 (假设性积压):** 虽然今日数据未提及，但从架构上看，由于项目涉及长连接（如网关、通道），应持续关注并积累相关议题和PR，进行主动管理。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*