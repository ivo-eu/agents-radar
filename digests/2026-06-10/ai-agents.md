# OpenClaw 生态日报 2026-06-10

> Issues: 258 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-06-10 05:08 UTC

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

好的，遵照您的指示，以下是为 **OpenClaw** 项目生成的 **2026-06-10** 项目动态日报。

---

# OpenClaw 项目动态日报 | 2026-06-10

## 1. 今日速览

今日项目活跃度**极高**，社区反馈和开发迭代均保持密集节奏。过去 24 小时内，Issues 更新达 258 条，其中新开和活跃 181 条，显示出用户反馈和问题发现的迅猛势头；Pull Requests 更新量更是高达 500 条，表明团队也在积极跟进解决。项目发布了新的小版本 `v2026.6.5` 和 `v2026.6.5-beta.6`，核心修复内容包括 QQBot 对模型推理框架的剥离以及 MCP 工具结果的类型强制转换。同时，多个高优先级（P1）的 Bug 仍在持续讨论中，社区对安全、消息丢失和会话状态一致性等稳定性问题表现出较高关注。

## 2. 版本发布

今日发布了 **2 个**新版本，内容如下：

- **`v2026.6.5` (正式版)** & **`v2026.6.5-beta.6` (测试版)**
    - **主要变更 (Highlights):**
        - **QQBot 体验优化**: 现在 QQ 机器人会在本地交付模型回复前，剥离掉模型的推（思考）过程框架 (`<thinking>...`)，防止原始的思维链内容泄漏到聊天频道中。相关 Issue: [#89913](https://github.com/openclaw/openclaw/issues/89913), [#90132](https://github.com/openclaw/openclaw/issues/90132)。
        - **MCP 工具结果处理增强**: MCP（Model Context Protocol）工具的结果现在会对 `resource_link`, `resource`, `audio`, 格式错误的 `image` 等类型进行强制类型转换，以增强兼容性和稳定性。
    - **破坏性变更 (Breaking Changes)**: 根据版本推测，本次更新涉及 MCP 工具结果的处理逻辑，如果其他第三方工具或插件依赖原始格式的 MCP 结果，可能需要进行适配。
    - **迁移注意事项**: 建议所有用户尤其是在使用 QQBot 和 MCP 工具的用户进行升级。测试版 (`-beta.6`) 可能包含进一步的实验性调整，建议生产环境谨慎使用。

## 3. 项目进展

今日项目整体处于**高速解决与争议并存**的状态。虽然有大量 PR 被提出，但关闭/合并的 PR 数量（131）远少于待合并的（369），表明很多修复方案仍在讨论和审查中。重要进展包括：

- **核心 Bug 修复推进**:
    - **会话状态持久化**: PR [#57825](https://github.com/openclaw/openclaw/pull/57825) 修复了一个关键问题，即通用的会话持久化过程会错误地清除之前设置的 `/model` 模型覆写。
    - **Cron 任务行为修复**: PR [#57831](https://github.com/openclaw/openclaw/pull/57831) 修复了 Cron 任务可能错误地写入用户实时 DM 会话的问题，避免了会话污染。
- **功能开发**: PR [#56023](https://github.com/openclaw/openclaw/pull/56023) 为 `memory-core` 扩展引入了来自 SkillFoundry 的质量门控和加权召回功能，增强了代理（Agent）在处理复杂任务时的代码质量。
- **安全加固**: PR [#85196](https://github.com/openclaw/openclaw/pull/85196) 开篇已有一个雄心勃勃的提案，旨在扩展工具输出的秘密/密钥红action（Redact）覆盖范围，这是一个重要的安全改进，目前仍在审查中。

## 4. 社区热点

今日最受关注的议题集中在 **会话上下文错乱** 和 **子任务无声丢失** 这两个核心痛点上。

1.  **会话上下文错乱 (Session Context Confusion)**
    -   **Issue**: [#32296](https://github.com/openclaw/openclaw/issues/32296) `[Bug]: Agent replies to previous message instead of current message (session context confusion)`
    -   **热度**: 15 条评论，P1 优先级。这是一个高优 Bug，表现为代理经常回复用户的上一条消息，导致对话“牛头不对马嘴”。
    -   **诉求**: 用户报告了该问题的症状，但根本原因尚未明确定论，社区在积极讨论可能的复现方式和对 `session` 处理逻辑的质疑。

2.  **子任务完成无声丢失 (Subagent Completion Silently Lost)**
    -   **Issue**: [#44925](https://github.com/openclaw/openclaw/issues/44925) `[Bug]: Subagent completion silently lost — no retry, no notification, no auto-restart on timeout`
    -   **热度**: 19 条评论，P1 优先级。用户详细描述了子任务编排的多种失败模式，导致任务结果在无人知晓的情况下被丢弃，严重影响工作流的可靠性。
    -   **诉求**: 社区呼吁增加子任务的**重试机制**、**通知机制**和**超时自动重启**的能力。这反映了用户对代理系统可靠性要求的提升，而不仅仅是功能的有无。

## 5. Bug 与稳定性

今日报告了大量 Bug，以下按严重程度排列，并标注是否有修复 PR。

- **【P1-高严重性】**
    - **会话上下文错乱** (`#32296`): 如前所述，核心交互问题。**尚无明确修复 PR。**
    - **子任务完成无声丢失** (`#44925`): 严重的工作流可靠性问题。**尚无明确修复 PR。**
    - **OpenAI ChatGPT Responses 传输失败** (`#90083`，**已关闭**): 针对 `gpt-5.4/5.5` 模型因 `invalid_provider_content_type` 错误而无法使用的问题，**已修复并关闭**。
    - **网关事件循环冻结** (`#56733`): 进程存活但无响应，所有 HTTP 请求超时，是一个严重的“心脏骤停”问题。**尚无明确修复 PR。**

- **【P2-中等严重性】**
    - **文本在工具调用间泄漏** (`#25592`): 代理内部的中间文本被意外发送到消息频道，造成隐私和UX问题。**有相关 PR `#57483` 试图从权限控制角度解决 `session_send` 的问题，但未直接解决此Bug。**
    - **Safeguard Compaction 忽略配置** (`#57901`): `compaction.model` 配置被忽视，无法选择指定的模型进行压缩。**有相关 PR `#58455` 试图改进压缩质量。**
    - **Cron 任务 `deleteAfterRun` 故障** (`#56488`): 一次性任务在执行后未被删除，而是反复触发。**尚无明确修复 PR。**

## 6. 功能请求与路线图信号

社区提出的功能需求反映了用户对更精细控制和更好用户体验的追求。

- **从“有”到“优”的精细化控制**: 多个 Feature Request 集中在**可配置的抑制（Suppress）** 上，例如：
    - `#39406`: 要求配置选项来抑制瞬时工具错误的警告信息。
    - `#8299`: 要求配置选项来抑制子任务完成后的自动“宣布（Announce）”步骤。
- **安全与权限深化**:
    - `#10659`: 引入“Masked Secrets（遮蔽密钥）”系统，让代理**能用**但不能**看**密钥，防御提示注入导致的关键信息泄漏。此需求有较高的关注度 (👍: 4)。
    - `#6615`: 在 `exec-approvals` 中加入黑名单支持，实现“除X外全部放行”的安全策略。该需求获得了**7个赞**，呼声很高。
    - `#56263`: 允许配置文件权限 (`chmod 0o640/0o750`) 以支持多用户部署。
- **未来版本的潜在方向**: **Slack Modal 支持 (`#88154`)** 和 **WebChat 界面优化（如 `#84216` 的最近会话下拉菜单）** 等需求，配合已有的 UI 优化 PR （如 `#90057` 的 Workboard 改进），表明下一个小版本可能会聚焦于提升**协作和交互体验**。

## 7. 用户反馈摘要

从今日的 Issue 讨论中，可以提炼出以下用户情绪和痛点：

- **严重不满、期待变革**: 在 `#32296`（会话错乱）和`#44925`（子任务丢失）的讨论中，用户语气趋于焦急和不满。他们期望的是一个“可靠的工具”，而不仅仅是“能用的功能”。这些 Bug 对日常使用和工作流造成了实质性障碍。
- **“承诺但不执行”的空头支票**: Issue `#58450` 中，用户（@al-osokin）生动地描述了代理的“画饼”行为：“我会查一下，然后...”，但实际上却没有执行任何后续动作。这是一种非常具体的、破坏信任的UX缺陷。
- **新版本伴随新问题**: 多个 Issue 提到升级到 `2026.6.1` 后出现了新的回归问题，如 OpenAI 传输失败 (`#90083`)、`/compact` 命令回复被静默丢弃 (`#90185`)。这表明虽然迭代快，但稳定性保障仍有提升空间，用户的升级意愿可能因此降低。
- **对黑屏/静默失败零容忍**: 用户对消息无声丢失（Google Chat `#58514`、WhatsApp `#50093`、WebSocket断连 `#45952`）表现出极低的容忍度。大量 Issue 集中于“silently ignored”、“silently lost”、“replies=0”，显示社区将代理通信的**透明度和可靠性**视为刚需。

## 8. 待处理积压

以下为长期未获回应或进展缓慢的高优先级（P1）Issue，需提醒维护者关注。

1.  **Issue:** `#75` `[Feature] Linux/Windows Clawdbot Apps`
    -   **创建时间**: 2026-01-01 (距今已**6个月**)
    -   **热度**: 109 条评论，79 个 👍
    -   **状态**: 标签众多，但仍为 `OPEN`。这是社区呼声最高的跨平台需求，但进展未明。
    -   [链接](https://github.com/openclaw/openclaw/issues/75)

2.  **Issue:** `#10659` `[Feature] Masked Secrets - Prevent Agent from Accessing Raw API Keys`
    -   **创建时间**: 2026-02-06 (距今约**4个月**)
    -   **热度**: 13 条评论，4 个 👍
    -   **状态**: P1 优先级，但仍有多个 `clawsweeper:needs-*` 标签，无明确的 PR 指向此 Issue。
    -   [链接](https://github.com/openclaw/openclaw/issues/10659)

3.  **Issue:** `#6731` `[Feature]: safe/unsafe ClawdBot`
    -   **创建时间**: 2026-02-02 (距今约**4个月**)
    -   **热度**: 12 条评论
    -   **状态**: 这是一个纯功能增强需求，但涉及 Rust 重写这一宏大设想，可能超出了当前路线图的范畴，导致讨论陷入停滞。
    -   [链接](https://github.com/openclaw/openclaw/issues/6731)

---

## 横向生态对比

好的，作为一名专注于 AI 智能体与个人 AI 助手开源生态的资深技术分析师，我已仔细阅读并分析了您提供的各项目动态日报。以下是根据这些数据生成的横向对比分析报告，旨在为技术决策者和开发者提供一份清晰的生态全景图。

---

### AI智能体与个人AI助手开源生态横向对比分析报告 (2026-06-10)

#### 1. 生态全景

当前，个人 AI 助手与自主智能体开源生态正处于一个 **“功能狂飙”与“可靠性阵痛”并存的阶段**。一方面，各项目核心功能日趋完善，多代理协作、技能市场、多平台集成等高级特性快速涌现，社区贡献异常活跃。另一方面，随着功能复杂度的提升，**会话上下文错乱、子任务无声失败、消息静默丢失**等稳定性问题成为社区普遍痛点，用户对可靠性的呼声已超越对“新奇功能”的追求。与此同时，**安全性**正从一个可选项演变为项目的核心生命线，多项目不约而同地开始深度加固安全边界。整体生态呈现出高创新、高迭代、但尚未形成绝对稳定范式的繁荣初期特征。

#### 2. 各项目活跃度对比 (过去24小时)

| 项目名称 | Issues (新开/活跃) | PRs (总数/合并关闭) | 新版本发布 | 社区健康度评估 | 核心特征 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 181 | 500 / 131 | ✅ v2026.6.5 | 非常高 | 社区规模巨大，问题发现全面，迭代速度极快，但稳定性挑战突出 |
| **Hermes Agent** | 15 | 50 / 18 | ❌ | 高 | 修复近期回归Bug为主，桌面端与Docker集成成重点，功能稳步推进 |
| **PicoClaw** | 21 | 16 / 5 | ✅ Nightly | 高 (安全事件驱动) | 安全事件成为焦点，暴露多模块漏洞，团队应急响应迅速 |
| **NanoBot** | 7 | 27 / 12 | ❌ | 中等 | 核心稳定性和基础设施得到增强（如HookCenter），趋向成熟稳健 |
| **CoPaw** | 13 | 35 / 16 | ✅ Beta | 高 | 安全加固、插件生态扩展，关注前沿特性（如学习循环） |
| **NanoClaw** | 1 | 43 / 39 | ❌ | 高 | 大量积压PR被清理，模块化、插件化、技能市场体系趋于成型 |
| **ZeroClaw** | 0 | 47 / 极少数 | ❌ | 中等 | 提交大量新功能PR，但合并率低，处于功能“堆砌”待整合阶段 |
| **NullClaw** | 1 | 7 / 7 | ❌ | 中等 | 快速修补社区报告的小问题，但一个影响Telegram定时任务的严重Bug待修复 |
| **IronClaw** | 7 | 50 / 4 | ❌ | 高 (开发密集) | 大型功能PR集中推进（附件、项目模型、审批策略），架构升级期 |
| **LobsterAI** | 1 | 4 / 4 | ❌ | 中等 | 专注数据备份、窗口恢复等桌面端体验优化，节奏稳健 |
| **Moltis** | 1 | 0 / 0 | ❌ | 低 | 24小时内无开发活动，只有一条新Issue，项目处于静默期 |
| **TinyClaw** | - | - | - | 沉默 | 24小时内无任何活动 |
| **ZeptoClaw** | - | - | - | 沉默 | 24小时内无任何活动 |

#### 3. OpenClaw 在生态中的定位

- **核心参照与“大象”**：从其“今日速览”即可看出，OpenClaw 的社区规模和迭代强度（188条活跃Issue, 500条PR/日）在所有项目中一骑绝尘，这确立了其作为生态中“核心参照”和“问题发现机”的地位。大多数其他项目遇到的问题，很可能已经在OpenClaw的Issue区被讨论过。

- **优势与差异**：
    - **协议领先 (MCP)**：OpenClaw 是MCP（Model Context Protocol）最强的支持者和推动者，其对MCP工具结果的类型强制转换等修复，体现了对“智能体即工具调用者”这一核心范式的深刻理解。
    - **“大而全”的综合能力**：不同于其他项目的某一侧重点，OpenClaw 试图覆盖从个人助手到多代理协作的方方面面，其问题覆盖的广度（会话、子任务、安全、网关、记忆）也最全面。

- **定位与短板**：OpenClaw 更像是一个“全能型”操作系统，其强大的功能以极高的复杂性和“可靠性阵痛”为代价。相比之下，其他项目致力于在特定场景下提供更“开箱即用”或更“极致可靠”的体验。

#### 4. 共同关注的技术方向 (多项目涌现)

社区的反馈揭示了多个跨项目的共识性需求，这些是当前AI智能体发展的核心瓶颈和未来方向：

1.  **会话管理与可靠性**：
    - **涉及项目**: **OpenClaw** (#32296 会话错乱)、**NanoBot** (#4259 跨会话注入)、**Hermes Agent** (#40416 压缩导致消息丢失)、**NullClaw** (#941 Agent定时任务不触发)。
    - **核心诉求**: 确保对话上下文的正确性和隔离性；分布式或复杂工作流下的任务状态一致性；消息传递的可靠性和透明性。

2.  **安全与权限控制**：
    - **涉及项目**: **PicoClaw** (#3068-3082 批量SSRF/CSRF漏洞)、**OpenClaw** (#10659 Masked Secrets)、**ZeroClaw** (#5982 多租户RBAC, #5775 技能级权限)、**NanoClaw** (#2722 配对码可预测漏洞)。
    - **核心诉求**: 从“功能实现”转向“安全防护”。防御提示注入、权限提升、敏感信息泄露，并提供细粒度的、可管理的权限控制（从用户级到技能级）。

3.  **多模型/多后端支持**：
    - **涉及项目**: **OpenClaw** (多种模型适配)、**NanoBot** (#4253 会话级模型切换)、**NanoClaw** (#1690 多运行时抽象)、**Hermes Agent** (#43245 模型不理解sudo)。
    - **核心诉求**: 抽象底层LLM提供商差异，实现模型热切换，支持本地/云端混用。核心矛盾在于模型对特定指令的理解和执行能力存在偏差。

4.  **成本与效率优化**：
    - **涉及项目**: **Hermes Agent** (#6839 懒加载工具提示)、**OpenClaw** (大量关于Token消耗的讨论)、**CoPaw** (#5063 Headroom集成)。
    - **核心诉求**: 减少Token浪费和API调用成本。通过`按需加载`、`智能压缩`、`中断长链思考`等机制，提升经济性和响应速度。

#### 5. 差异化定位分析 (关键差异)

| 项目 | 功能侧重 | 目标用户 | 技术架构 | 社区文化 |
| :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 全功能个人/企业AI中心，MCP核心，多Agent协作 | 技术极客、开发者、企业团队 | 后端进程，多网关，插件化 | 大而全，问题驱动，迭代快 |
| **NanoBot** | 轻量级AI助手，核心LLM调用与上下文管理者 | 个人开发者，追求简洁与可控 | 模块化，Plugin/Hook架构 | 稳健务实，重度依赖核心开发者 |
| **Hermes Agent** | 桌面端优先的AI Agent，注重UI/UX | 桌面硬件用户，追求沉浸式体验 | 桌面应用，对Docker有优化 | 活跃，对成本敏感，注重软件工程 |
| **PicoClaw** | 高度关注安全的个人助手 | 安全研究员，隐私敏感用户 | 前端(Launcher)+后端，多通道 | **安全事件驱动**，应急响应快 |
| **NanoClaw** | 可插拔、模块化的多技能平台 | 生态构建者，集成开发者 | 技能市场、多运行时、插件系统 | 社区贡献者多，贡献质量高 |
| **ZeroClaw** | 标准化流程(SOP)、多租户企业级部署 | 企业IT部门、运维人员 | AMQP/MQTT驱动，强调流程自动化 | 功能导向，合并流程偏慢 |
| **NullClaw** | 快速修复Bug，精简稳定的个人助手 | 普通用户，追求“稳定即正义” | 对单一问题快速定位修复 | 小团队维护，问题响应快 |
| **IronClaw** | 大规模功能迭代（附件、多租户、可观测性） | 依赖Near生态的开发者 | Rust底层，复杂功能模块化 | 开发密集，架构升级期 |

#### 6. 社区热度与成熟度

- **快速迭代与问题高发期**：
    - **OpenClaw, PicoClaw, ZeroClaw, IronClaw**: 这些项目处于最前沿，社区最活跃，但同时也是Bug和问题发现最多的地方。它们为整个生态贡献了大量的试错成本和新方向探索。
    - **特征**: Issues和PR数量极高，稳定性问题频发，但功能迭代也最快。

- **质量巩固与成熟稳健期**：
    - **NanoBot, NullClaw, LobsterAI**: 这些项目主要活动是修复Bug、提升文档、打磨现有功能。其核心架构相对稳定，社区贡献更偏向于“完善”而非“突破”。
    - **特征**: 新Feature Request较少，主要集中在对现有系统的小修小补和体验优化上。

- **特色专注期**：
    - **CoPaw, NanoClaw, Hermes Agent**: 它们在其特定的细分领域（如本地化体验、插件化、桌面交互）中拥有极高的社区忠诚度和高质量的贡献。
    - **特征**: 在特定维度上的迭代强度很高，体现出差异化优势。

#### 7. 值得关注的趋势信号 (对AI智能体开发者的参考)

1.  **“可靠性”是通向大众市场的门票**：大量用户反馈揭示，**会话混乱、任务无声丢失、消息静默失败**是目前智能体最需要解决的“最后一公里”问题。开发者应优先考虑**可观测性、状态一致性、失败重试与通知机制**，否则功能再强也无法建立用户信任。
2.  **安全不再是“锦上添花”**：PicoClaw 的批量安全报告是一个强烈信号。**提示注入防御、密钥管理、权限隔离**必须成为整个开发流程的一部分，而非事后修补。任何对API Key的暴露都是不可接受的。
3.  **“成本效率”正在工具化**：从手动寻找到**工具化**。如 Hermes Agent 的“懒加载工具”和 CoPaw 的“集成Headroom”，表明社区正在将Token优化开发成通用的、可插拔的组件，这将是未来评估智能体框架的重要指标。
4.  **从“单体Agent”到“智能体生态”**：NanoClaw 的技能市场和多运行时抽象，NanoBot 的HookCenter，OpenClaw 的MCP，都指向一个共同的未来：**构建一个开放、可组合、跨平台的智能体生态系统**，而不仅仅是开发一个单一的Agent应用。开发者应采纳**插件化、协议化和微内核**的设计思想。
5.  **“配置即代码”与“可调试性”**：多个项目的Issue表明，复杂的配置往往成为用户踩坑的重灾区。简化并可视化配置（如ZeroClaw的配置陷阱），同时提供强大的调试工具（如OpenClaw的`/compact`命令引发的困惑），是提升用户体验（DX）的关键。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot 项目动态日报 — 2026-06-10

---

## 1. 今日速览

- **社区与开发活跃度较高**：过去24小时新开/活跃Issue 7条（均未关闭），PR总数27条（其中12条已合并/关闭），无新版本发布。  
- **核心稳定性修复集中**：多个与模型调用、上下文管理、工具执行相关的Bug被提交，同时有对应修复PR快速响应（如#4263、#4270）。  
- **基础设施与用户体验并行推进**：HookCenter事件钩子系统（#3564）已合并，WebUI新增“Fork from here”功能（#4208），文档重写（#4177）也在今天关闭。  
- **测试与安全性持续投入**：多名贡献者（如yu-xin-c）提交了内存生命周期测试、runner工具覆盖、symlink逃逸阻断等测试/安全PR，表明项目对质量与安全的重视。  

---

## 2. 版本发布

**无**（过去24小时无新Release）。

---

## 3. 项目进展 — 重要合并/关闭 PR

以下PR在今天被合并或关闭，代表了项目在功能、修复、基础设施上的实质性推进：

| PR | 标签 | 要点 | 链接 |
|----|------|------|------|
| #3564 | feat, hooks | **HookCenter 类型化事件钩子系统** — 引入插件化hook分发机制（observe/transform/guard），替代旧AgentHook方法重写，外部开发者可通过`nanobot.hooks` entry points分发插件，旧系统完全向后兼容。这是架构升级的关键步骤。 | [PR #3564](https://github.com/HKUDS/nanobot/pull/3564) |
| #4263 | fix, provider | **修复GPT-5.x / o系列模型拒绝`max_tokens`问题** — 对未声明`supports_max_completion_tokens`的provider，自动根据模型名回退使用`max_completion_tokens`，使Azure等自定义端点能正常调用推理模型。 | [PR #4263](https://github.com/HKUDS/nanobot/pull/4263) |
| #4208 | feat, webui | **WebUI 新增“Fork from here”** — 允许用户从任意助手回复处派生新会话，保留前缀，原会话不变。提升对话管理灵活性。 | [PR #4208](https://github.com/HKUDS/nanobot/pull/4208) |
| #4177 | docs | **入门文档重写** — 重新组织文档入口，区分零基础用户、CLI快速上手、WebUI、部署等路径，降低新手门槛。 | [PR #4177](https://github.com/HKUDS/nanobot/pull/4177) |
| #4265 | feat, skill | **英语阅读技能调度改为每2天一次** — 调整cron表达式从每日变为每2天。 | [PR #4265](https://github.com/HKUDS/nanobot/pull/4265) |

> **评估**：当日合并的PR覆盖了架构改进、关键兼容性修复、WebUI易用性增强以及文档优化，项目整体健康度良好。

---

## 4. 社区热点 — 最活跃议题

| 议题 | 类型 | 评论数 | 热度原因 | 链接 |
|------|------|--------|----------|------|
| #4013 — [bug, question] Error calling LLM: stream stalled for more than 90 seconds | Issue | 3 | 用户升级到0.2.0后高频出现流中断，影响正常使用，社区关注度较高 | [Issue #4013](https://github.com/HKUDS/nanobot/issues/4013) |
| #4253 — [enhancement] support overriding model per conversation | Issue | 3 | 用户希望根据隐私/时效性需求在不同会话间切换模型，使用场景具体，讨论活跃 | [Issue #4253](https://github.com/HKUDS/nanobot/issues/4253) |
| #4259 — [enhancement, refactor] `history.jsonl`跨会话注入导致上下文污染 | Issue | 2 | 深入分析了`Consolidator`与`ContextBuilder`之间的会话隔离缺失问题，技术细节丰富，引发架构讨论 | [Issue #4259](https://github.com/HKUDS/nanobot/issues/4259) |

**热点诉求分析**：  
- 用户对**模型调用稳定性**（#4013）和**会话隔离**（#4259）的抱怨集中，前者影响工作效率，后者可能导致长期上下文混乱。  
- **会话级模型切换**（#4253）是高级用户的核心需求，展示了项目在灵活配置上的缺口。  

---

## 5. Bug 与稳定性

| 严重度 | Issue | 描述 | 修复状态 |
|--------|-------|------|----------|
| **高** | #4013 | LLM流中断超过90秒，升级0.2.0后复现，疑似客户端/服务端超时硬编码 | 未关联PR，待调查 |
| **高** | #4259 | `history.jsonl`跨会话注入导致的上下文污染，可能输出错误信息 | 未关联PR，但有详细分析 |
| **高** | #4061 | OpenAI兼容的文本格式工具调用未被解析为结构化调用，导致无法触发工具 | 未关联PR |
| **中** | #4264 | `idleCompact`仅使用除最后8条外的历史进行总结，导致修正信息丢失 | **已有修复PR #4270** |
| **中** | #4261 | `max_tokens`/`max_completion_tokens`歧义导致GPT-5.x调用失败 | **已于#4263合并修复** |
| **低** | #4262 | Agent模式启动时botIcon未优先使用 | 未关联PR，属UI细节 |

**关键关注**：  
- #4264的修复PR #4270（axelray-dev）今日已提交，将idleCompact改为使用完整会话历史进行归档，预计很快合并。  
- #4261已由PR #4263修复并合并，解决了Azure自定义provider调用推理模型的阻塞问题。  
- 高严重度的#4013、#4259、#4061尚无修复PR，社区期待官方跟进。

---

## 6. 功能请求与路线图信号

| 需求 | 链接 | 相关PR/状态 | 纳入下版本可能性 |
|------|------|-------------|------------------|
| **每个会话覆盖模型** (#4253) | [Issue](https://github.com/HKUDS/nanobot/issues/4253) | 无关联PR，社区讨论中 | 高 — 用户刚需，且已有全局配置基础 |
| **使用botIcon作为默认头像** (#4262) | [Issue](https://github.com/HKUDS/nanobot/issues/4262) | 无关联PR | 低 — 简单增强，可能很快有PR |
| **WebUI版本按需检查** (#4255) | [PR #4255](https://github.com/HKUDS/nanobot/pull/4255) | 开放PR，移除后台轮询 | 高 — 已实现，等合并 |
| **StepFun ASR转录提供商** (#4260) | [PR #4260](https://github.com/HKUDS/nanobot/pull/4260) | 开放PR，新增SSE流式转录 | 中 — 功能扩展，由外部贡献者提交 |
| **fenced-code-block感知的split_message** (#4257) | [PR #4257](https://github.com/HKUDS/nanobot/pull/4257) | 开放PR，修复长消息拆分破坏代码块渲染 | 高 — 修复bug，提升消息展示质量 |
| **max-iteration最终回复改进** (#4269) | [PR #4269](https://github.com/HKUDS/nanobot/pull/4269) | 开放PR，增加无工具最终化pass | 中 — 提升Agent用户体验 |

**路线图信号**：  
- 多个开放PR集中于**消息渲染改进**（#4257、#4269）、**测试与安全**（#3982、#3983、#4053、#4119、#4193、#4256）以及**提供者兼容性**（#3869、#4268）。这些提示项目在稳定性和扩展性上持续投入。  
- #4255（按需版本检查）若合并将为WebUI瘦身并减少网络请求，符合轻量设计理念。

---

## 7. 用户反馈摘要

从Issue评论和PR讨论中提炼真实用户声音：

- **“升级0.2.0后无法正常工作”** （#4013）: 用户从0.1.5post2升级后遭遇流中断，需要反复“keep going”，认为“硬编码的/goal”导致无法完成实际工作。表达对之前版本的满意（“very good”），对升级的不满。  
- **“希望按隐私/时效性切换模型”** （#4253）: 用户使用OpenRouter（快速）和本地llamacpp（私有）两种预设，期望在对话中灵活切换，认为“全局设置不够”。  
- **“idleCompact丢失修正信息”** （#4264）: 用户描述典型场景——模型犯错、用户纠正、得到正确结果，但idleCompact仅保留最后8条外的历史，导致纠正行为被排除，历史存档留下错误结论。表达了对总结机制准确性的关切。  
- **“GPT-5.4需要max_completion_tokens”** （#4261）: 用户在使用Azure部署GPT-5.4时遇到参数不匹配错误，指出需要区分API参数，贡献了修复PR #4263。  
- **“WebUI会话内容被静默丢弃”** （#4267）: 用户报告间歇性助手回复在渲染中消失但已保存，认为与token生成速度相关。目前有修复PR #4267。

**整体感受**：社区活跃且技术细节丰富，用户多为有经验的开发者，对模型行为、会话一致性要求高；对升级带来的回归较敏感，期望官方快速响应。

---

## 8. 待处理积压

以下重要Issue/PR长期未响应或进展缓慢，提醒维护者关注：

| 类型 | 编号 | 标题 | 最后更新 | 延迟时间 | 建议行动 |
|------|------|------|----------|----------|----------|
| Issue | #4013 | Error calling LLM: stream stalled for more than 90 seconds | 2026-06-10（作者更新） | 创建于2026-05-26，持续收到评论 | 需确认是否与provider超时设置有关，建议指派负责人 |
| Issue | #4061 | OpenAI-compatible text-format tool calls not parsed | 2026-06-09 | 创建于2026-05-29 | 影响多个provider，建议制定解析方案 |
| Issue | #4259 | `history.jsonl`跨会话注入上下文污染 | 2026-06-09 | 创建于2026-06-09（新汇报） | 已有详细分析，建议尽快验证并修复，防止数据损坏 |
| PR | #3869 | fix(providers): DeepSeek message hardening | 2026-06-10（最后提交） | 创建于2026-05-16，长时间开放 | 需要review或merge决策，DeepSeek兼容性对部分用户重要 |
| PR | #3983 | test: cover runner blocked tool-call finish reasons | 2026-06-09 | 创建于2026-05-24，多次更新 | 测试PR较安全，建议优先合并以提升覆盖率 |
| PR | #3982 | test: add scripted agent runner harness | 2026-06-09 | 创建于2026-05-24 | 同上，作为runner测试基础设施，建议合并 |

> **注意**：以上积压项截至2026-06-10尚未获得官方回应或合并，部分PR已由同一贡献者持续更新（如yu-xin-c），建议维护者安排review轮次。

---

*生成时间：2026-06-10 23:59 UTC | 数据来源：GitHub HKUDS/nanobot*

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，我已根据 Hermes Agent 项目的 GitHub 数据生成以下项目动态日报。

---

# Hermes Agent 项目动态日报 | 2026年6月10日

## 1. 今日速览

今日 Hermes Agent 项目活跃度极高，24 小时内产生了 65 次更新事件（15 条 Issues 和 50 条 PRs）。社区贡献主要集中在**修复近期重构引入的 UI/UX 回归问题**、**完善 Docker 与系统服务集成**以及**推进一套复杂的新认知技能模块**。尽管今日无新版本发布，但多个关于核心功能的 PR 被合并（如代理网关持久化、内存写入控制），标志着项目稳定性与可定制性向前迈进了一步。值得关注的是，一个关于**懒加载工具模式节省 Token** 的高赞讨论（#6839）持续发酵，反映了社区对成本优化的强烈诉求。

## 2. 版本发布

**无新版本发布。**

## 3. 项目进展

今日合并/关闭的 18 个 PR 中，以下推进了项目在关键领域的进展：

- **代理网关稳定性与 Docker 集成**：
    - **PR #33619**: 修复了容器重启时网关未能自动恢复`running`状态的问题（[链接](https://github.com/NousResearch/hermes-agent/pull/33619)）。
    - **PRs #28888, #28891, #30887, #39683**: 一系列针对 `terminal.docker_extra_args` 配置项的修复被相继合并，彻解决了该配置在 CLI、网关和 Docker 环境下的“静默失效”Bug。这展现了社区对**Docker 安全加固（如 `--read-only`）**需求的落实。
- **核心功能增强**：
    - **PR #38199**: `feat(memory,skills): approve/deny gate for memory + skill writes` ( [链接](https://github.com/NousResearch/hermes-agent/pull/38199) ) 被合并。此功能为用户提供了对代理内存和技能写入的细粒度控制，直接回应了关于“代理自己学会错误知识”的担忧，是提升代理长期行为可控性的重要一步。
    - **PR #43036**: `fix(cli): run one-shot query cleanup before lease release` ( [链接](https://github.com/NousResearch/hermes-agent/pull/43036) ) 被合并。此修复确保了一次性查询命令能正确释放资源，避免了潜在的会话冲突。
- **桌面应用体验提升**：
    - **PR #43292**: `feat(desktop): install any VS Code theme from the Marketplace` ( [链接](https://github.com/NousResearch/hermes-agent/pull/43292) ) 被合并。此功能允许用户直接从 VS Code 市场安装主题，极大地扩展了桌面端的个性化能力，并改善了终端着色器、快捷键 HUD 的视觉风格。
    - **PR #39335**: `feat(desktop): full Russian localization (i18n)` ( [链接](https://github.com/NousResearch/hermes-agent/pull/39335) ) 被合并。项目国际化取得进展，增加了俄语支持。

**总结**: 今日合并的 PR 共同提升了核心组件的稳定性（网关、Docker）、可扩展性（技能门控）和终端用户满意度（主题、国际化）。

## 4. 社区热点

今日讨论最热烈的议题反映了社区的两大核心关注点：**优化成本**与**改善用户体验**。

1.  **Issue #6839: “[Feature] Lazy Tool Schema Loading — Two-Pass Tool Injection to Reduce Token Overhead”** ( [链接](https://github.com/NousResearch/hermes-agent/issues/6839) )
    - **热度**: 23条评论，13个👍
    - **诉求**: 用户核心痛点在于当前每次API调用都会注入全部工具的Schema，导致严重的Token浪费（估算3500-5000 tokens/次），尤其是在本地模型上，成本高昂。提议采用“两遍式工具注入”机制，按需加载工具Schema。
    - **分析**: 这是目前社区最关注的性能优化方向，直接关系到用户的使用成本，特别是对于使用付费API或在消费级硬件上运行本地模型的用户。该功能若被采纳，将是项目在效率优化上的一个里程碑。

2.  **Issue #40416: “[Bug] Context compaction visually deletes messages from user's chat — terrible UX”** ( [链接](https://github.com/NousResearch/hermes-agent/issues/40416) )
    - **热度**: 3条评论，1个👍
    - **诉求**: 用户在Telegram上进行对话时，上下文压缩功能会将历史消息从界面上“删除”，给用户造成消息被误删的极差体验。
    - **分析**: 这是一个典型的“设计意图”与“用户感知”不匹配的问题。技术上的“压缩”在用户看来就是“消失”，严重破坏了用户对代理的信任感。此问题急需UI/UX层面的改进，而非单纯的技术修复。

3.  **Issue #43245: “[Feature] most models ignore the use of the PTY terminal tool when needing sudo”** ( [链接](https://github.com/NousResearch/hermes-agent/issues/43245) )
    - **热度**: 2条评论
    - **诉求**: 用户反映，在需要sudo权限的场景下，LLM总是尝试各种“创造性”但不符合项目设计的方案来获取权限，而不是使用项目推荐的PTY终端工具。
    - **分析**: 这揭示了LLM的安全约束与既定指令执行之间的深层矛盾。问题的核心可能在于LLM本身对安全策略的“理解”不足，而非单纯的技能实现问题，这需要项目组在底层指令工程或模型选择上寻找更根本的解决方案。

## 5. Bug 与稳定性

今日报告的 Bug 集中在**桌面端UI回归**和**系统集成兼容性**两类。

|严重程度| Issue 链接 | 摘要 | 状态分析 |
| :--- | :--- | :--- | :--- |
| **P1** | [#42675 (CLOSED)](https://github.com/NousResearch/hermes-agent/issues/42675) | Docker容器重启后，网关无法自启动，导致消息渠道（微信/Telegram等）静默失效。 | **已修复**: 该问题已被PR #33619解决。 |
| **P2** | [#43232 (OPEN)](https://github.com/NousResearch/hermes-agent/issues/43232) | 当 `HERMES_HOME` 位于NAS网络文件系统(NFS/SMB)上时，`hermes dashboard` 进程启动时进入D状态（不可中断睡眠），无法响应。 | **待处理**: 严重影响使用NAS作为存储的高端用户。 |
| **P2** | [#43228 (OPEN)](https://github.com/NousResearch/hermes-agent/issues/43228) | 飞书/Lark消息卡片Footer显示的模型名称不正确，被后台辅助任务（如视觉理解）使用的模型名覆盖。 | **待处理**: 影响飞书用户的日常体验，错误引导用户认知。 |
| **P2** | [#40416 (OPEN)](https://github.com/NousResearch/hermes-agent/issues/40416) | 上下文压缩在Telegram上视觉上删除了用户消息。 | **待处理**: 严重影响用户体验，被标记为“糟糕的UX”。 |
| **P2** | [#43295 (OPEN)](https://github.com/NousResearch/hermes-agent/issues/43295) | Docker Compose文件中指定特定配置文件的`command`参数被入口点脚本忽略。 | **已有修复PR**: PR #43314 已提交修复。 |
| **P3** | [#43309 (OPEN)](https://github.com/NousResearch/hermes-agent/issues/43309) | 桌面端侧边栏的Cron任务区域在缩放时会跳动。 | **已有修复PR**: PR #43310 已提交修复。 |

## 6. 功能请求与路线图信号

今日提出的新功能请求，结合已有 PR 趋势，揭示了项目未来的可能发展方向：

- **核心效率优化**: Issue #6839 关于**Lazy Tool Schema Loading**（懒加载工具模式）的提议，极有可能成为下一个版本的开发重点之一，因为它触及了几乎所有用户的使用成本痛点。
- **智能体行为可塑性**: PR #38199 的合并，以及新Issue #43245 对 `sudo` 权限控制不力的探讨，都指向了同一个方向：**让代理更“听话”**。给代理引入一个“控制器”层（`write_mode`门控、`sudo`技能约束），是项目正在加强的方向。
- **桌面端成为标配**: 新功能PR #43292 (VS Code主题) 和被快速响应的视觉回归Issue #43309，表明**Hermes Desktop**正在成为社区的主力使用平台，其UI/UX的打磨和功能增强将加速。
- **深度用户个性化**: Issue #40399（请求桌面端主题和字体定制）与合并的PR #43292相呼应，说明个性化定制需求强烈。新出现的`qca-cycle`认知技能PR #43306，虽然目前看较为试验性，但代表了社区对“**赋予代理独特人格**”这一方向的探索。
- **企业级/场景化适配**: Issue #43294 (自动信任企业CA证书) 和 #43307 (请求CLI界面支持中文) 反映了不同场景下用户的特定需求，这表明项目正在从通用工具向解决更多细分领域问题的方向发展。

## 7. 用户反馈摘要

从今日的 Issues 和 PR 评论中，可以提炼出以下真实的用户声音：

- **痛并热爱着**: 用户一方面对项目的高度可定制性和强大的能力表示满意（例如在 #40399 中提到“渴望一个作为日常驱动器的桌面应用的精致感”），但另一方面也对一些基础体验问题感到困扰，如上下文压缩导致的消息丢失（#40416）。
- **对“成本”的敏感**: 社区对 Token 消耗异常敏感。Issue #6839 的提出者并非抱怨，而是以一种非常技术性的方式提出了解决方案，这表明用户愿意为优化成本做出贡献。
- **代码质量被“看见”**: 针对`docker_extra_args`的多个冗余PR虽被合并，但也引出了关于代码库健康度的问题。有评论（尽管未在数据中直接展现）暗示项目可能存在“重复造轮子”或代码审查流程的冗余，这可能是未来需要关注的内部运维问题。
- **关于“门槛”的讨论**: Issue #43245 (sudo权限) 的讨论暗示，即使是最先进的LLM，在遵循复杂、非通用的安全策略时也表现不佳。用户期望的是“模型能够理解并遵循我们设定的规则”，而非强行绕过。

## 8. 待处理积压

今日数据中未发现长期（超过1月）未响应的重要Issue或PR。多数活跃议题（如 #6839）都在持续被讨论和更新。建议社区维护者关注以下较新但影响面广的议题：

- **P2 级别，已有修复PR的Bug**: `#43295 (Docker entrypoint ignores compose command arguments)` ( [链接](https://github.com/NousResearch/hermes-agent/issues/43295) ) 的修复PR `#43314`应尽快审查合并。
- **P2 级别，影响终端用户反馈的Bug**: `#43228 (Feishu footer model name polluted)` ( [链接](https://github.com/NousResearch/hermes-agent/issues/43228) ) 和 `#40416 (Context compaction deletes messages)` ( [链接](https://github.com/NousResearch/hermes-agent/issues/40416) ) 虽无阻塞性，但对特定平台用户影响大，建议分配资源。
- **反映开发者需求的Issue**: `#43036` 的提交者 `mnajafian-nv` 是项目活跃贡献者，其提交的关于 `one-shot query` 的修复已合并，体现了良好协作。

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

好的，这是为您生成的 PicoClaw 项目动态日报。

---

# PicoClaw 项目动态日报 | 2026-06-10

## 1. 今日速览

今日项目整体活跃度极高，**安全事件成为绝对焦点**。过去24小时内，社区共提交并处理了21条Issue和16条PR，并发布了新的nightly版本。值得注意的是，社区成员 **@YLChen-007** 一次性提交了12个安全相关的Issues（#3068至#3082），揭示了包括SSRF绕过、CSRF、授权绕过在内的多个关键漏洞。项目维护者响应迅速，已提交针对其中部分问题的修复PR（如 #3083、#3085），显示出强大的安全应急能力。总体来说，项目正处于快速迭代与安全加固并行的阶段。

## 2. 版本发布

- **nightly (v0.2.9-nightly.20260610.b9a8fad6)**
  - **类型**: 自动构建的Nightly版本。
  - **说明**: 此版本为自动构建，可能包含不稳定特性。建议谨慎使用。
  - **更新内容**: 
    > **Full Changelog**: https://github.com/sipeed/picoclaw/compare/v0.2.9...main
  - **迁移注意事项**: Nightly版本主要用于测试新功能，不建议在生产环境中直接使用。

## 3. 项目进展

今日共关闭/合并了5个PR，主要涉及配置迁移稳定性、模型兼容性修复和文档更新。具体进展如下：

- **修复默认Claude模型ID错误** (#2942): 修复了因默认配置中使用了带点的模型ID（`claude-sonnet-4.6`），导致新安装的Launcher首次对话即失败的问题。现已修正为连字符格式（`claude-sonnet-4-6`）。
- **解决Claude Opus 4-7的兼容性问题** (#2940): 针对`claude-opus-4-7`模型已弃用`temperature`参数导致请求失败（HTTP 400）的问题，合入修复，当调用该模型时，会在请求体中省略`temperature`字段。
- **强化配置迁移类型安全** (#3064): 修复了配置迁移(`pkg/config/migration.go`)中一条未检查类型断言的代码，该问题可能导致在特定情况下发生运行时panic。增加安全检查后，会在遇到错误配置时返回错误信息而非崩溃。
- **修复受限工作区中相对路径的误判** (#3087): 修复了当启用`restrict_to_workspace`时，对于类似 `skills/calendar-query/scripts/query_calendar.py` 的相对路径，执行安全检查会错误地阻止其执行的问题。
- **更新微信二维码** (#3086): 更新了文档中的微信联系方式二维码。
- **基础设施更新** (#3084): 标准化了`.gitignore`文件的文本编码，修复了因格式错误导致Git将其视为二进制文件的潜在问题。

## 4. 社区热点

今日社区讨论最热烈、关注度最高的事件集中在 **安全漏洞批量报告** 和 **功能请求的长期讨论** 上。

- **【安全焦点】批量SSRF及授权绕过报告**
  - **链接**: [#3077](https://github.com/sipeed/picoclaw/issues/3077), [#3078](https://github.com/sipeed/picoclaw/issues/3078), [#3080](https://github.com/sipeed/picoclaw/issues/3080), [#3082](https://github.com/sipeed/picoclaw/issues/3082) 等
  - **分析**: 安全研究员 **@YLChen-007** 提交的一批高质量安全报告成为今日社区讨论的绝对核心。这些报告指出`web_fetch`工具、Launcher权限控制、Feishu通道等多个模块存在多种绕过方式（如利用特殊IPv6地址、环境代理、ISATAP隧道等）。这反映了项目在安全边界防御上仍有提升空间，也极大地激发了社区对项目安全性的关注。

- **【长期讨论】支持流式HTTP请求**
  - **链接**: [#2404](https://github.com/sipeed/picoclaw/issues/2404)
  - **分析**: 该Issue自4月提出以来，已获得11条评论和1个赞，是今日讨论热度最高的非安全类Issue。用户 **@OuSatoru** 希望能在配置文件中添加 `"streaming": true` 选项，以支持像OpenAI Python客户端那样的流式LLM请求。这反映了用户对**低延迟、实时交互体验**的强烈需求。尽管该功能目前尚未有关联PR，但代表了社区一个重要的功能诉求。

## 5. Bug 与稳定性

今日报告的Bug主要集中在安全漏洞和模型兼容性问题，部分已有修复PR在跟进。

- **【严重/高】多项安全漏洞（SSRF、CSRF、授权绕过等）**
  - **描述**: 如前文所述，`@YLChen-007` 提交的12个Security Issues覆盖了SSRF、CSRF、命令执行白名单绕过、配置重载权限滥用等多个高危领域。这些漏洞可能被利用进行内网探测、数据泄露甚至控制权接管。
  - **状态**: 部分漏洞已有快速响应性修复PR。例如：
    - **#3085** (`fix(tools): block 198.18.0.0/15 in SSRF guard`) 修复了`#3077`。
    - **#3083** (`feat(web): harden launcher access control`) 修复了`#3069`和`#3080`。

- **【中】Claude Opus 4-7模型调用失败** (#2939，已关闭)
  - **描述**: 由于`temperature`参数被弃用，调用该模型时返回HTTP 400错误。
  - **状态**: 已通过PR #2940修复并关闭。

- **【中】对话记录显示不完整** (#2796，已关闭)
  - **描述**: 用户反映在历史对话中，如果同一轮对话有多条用户消息，只能看到最后一条。
  - **状态**: 已有关联的修复PR #2990处于待合并状态，等待进一步审查。

- **【低】Windows平台下`list_dir`工具报错** (#2472，已关闭)
  - **描述**: 由于路径分隔符不匹配，`list_dir`工具在Windows系统上返回"invalid argument"错误。
  - **状态**: 已关闭，推测已修复。

## 6. 功能请求与路线图信号

今日涌现的功能请求与项目现有PR显示出几个可能的发展方向：

- **协议与扩展性增强**
  - **流式响应**: Issue [#2404](https://github.com/sipeed/picoclaw/issues/2404) 强烈要求增加对LLM流式响应的支持。
  - **WebSocket协议完善**: Issue [#2984](https://github.com/sipeed/picoclaw/issues/2984) 请求增加明确的“回合结束”信号，以便外部客户端能准确判断Agent处理状态。
  - **Agent协作**: PR [#2937](https://github.com/sipeed/picoclaw/pull/2937)（搁置中）引入了一个全新的“Agent协作总线”，表明项目正在探索多Agent通信与协作这一高级功能。

- **安全性基础设施升级**
  - **加密库替换**: Issue [#3088](https://github.com/sipeed/picoclaw/issues/3088) 建议用更安全、且维护中的 `vodozemac` 替换已不安全的 `libolm` 加密库。这是一个重要的安全基础设施改进信号。

- **新平台/提供商支持**
  - **新网关**: PR [#3063](https://github.com/sipeed/picoclaw/pull/3063) 正在为项目添加DeltaChat网关支持。
  - **新LLM提供商**: PR [#2917](https://github.com/sipeed/picoclaw/pull/2917) 正在集成NEAR AI Cloud作为新的AI提供商。

## 7. 用户反馈摘要

从Issues的评论中，我们可以提炼出以下用户反馈：

- **痛点与不满**:
  - **模型兼容性**: 用户 `@LegendAlessandro-Liguori` 在使用 `claude-opus-4-7` 时遇到的直接失败问题，虽然已修复，但反映出项目对新模型更新的兼容性测试需要加强。
  - **对话体验**: 用户 `@EverestSnow` 反馈的对话历史记录显示不完整问题，影响了回顾和编辑长对话的体验。
  - **配置困惑**: 用户 `@yuxuan-7814` 在PR #2988中修正了`/context`命令显示内容与配置不符的bug，说明配置的生效逻辑可能不够直观。

- **使用场景与期望**:
  - **开发工具集成**: 用户希望能够通过流式HTTP请求与LLM后端交互，类似OpenAI Client库的用法，这表明开发者倾向于将PicoClaw作为后端组件集成到自有应用中。
  - **安全部署**: 大量的安全报告来自同一位研究员，表明有用户或团队正在对PicoClaw进行深度的安全审计，这对于项目在敏感或企业级环境中的长期部署是积极的信号。

## 8. 待处理积压

以下为一些长期未更新或未合并的重要Issue/PR，建议维护者关注：

- **重要PR搁置中**:
  - **[#2937] Feat/agent collaboration** (https://github.com/sipeed/picoclaw/pull/2937): 这是一个重要的新功能PR，引入了多Agent协作机制。自5月底以来已被标注为`[stale]`，可能需要维护者评估其合并或调整方向。
  - **[#2917] feat(provider): add NEAR AI Cloud provider** (https://github.com/sipeed/picoclaw/pull/2917): 自5月21日以来未有更新。
  - **[#2983] fix(agent): retry empty llm response** (https://github.com/sipeed/picoclaw/pull/2983): 修复一个可能导致请求失败的边缘情况，已搁置近10天。

- **长期未响应的重要Issue**:
  - **[#2404] [Feature] Add in config to send streaming HTTP request** (https://github.com/sipeed/picoclaw/issues/2404): 这是一个呼声很高的功能需求，评论持续到6月9日，但尚未有官方维护者认可或指派的迹象。

---

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw 项目日报（2026-06-10）

## 今日速览
过去 24 小时 NanoClaw 项目保持高活跃度：共处理 43 条 Pull Request，其中 39 条被合并或关闭，仅剩 4 条待合并；新开 Issue 1 条（#1690），聚焦多运行时抽象，社区讨论活跃。没有新版本发布，但大量积累数月的 PR 得到清理，代码库在功能、安全、文档三方面均有实质推进。整体项目健康度良好，社区贡献持续涌入。

## 版本发布
无新版本发布。

## 项目进展
今日合并/关闭的重要 PR 涵盖功能增强、安全修复、平台适配和文档完善，标志着项目向模块化、可观测性和易用性迈出重要一步：

- **安全加固**：`#2722` 将 Telegram 配对码生成从 `Math.random` 切换为 `crypto.randomInt`，防止可预测攻击，已提交待合并。`#2718` 修复飞书通道在生产环境中因 agent-runner 异常退出导致交互卡片卡在“运行中”状态的僵尸卡片问题，已合并。
- **模块化与生态**：`#1309` 实现技能市场/注册表系统（CLI 发现、安装、管理 GitHub 托管技能），`#1387` 新增类似 channel 的插件系统，`#1245` 引入 `/approve` 和 `/reject` 审批技能，`#1285` 增加直接运行模式（无需 Docker 容器，降低资源开销）。这些功能均已合并，极大提升了 NanoClaw 的可扩展性和部署灵活性。
- **可观测性**：`#1202` 新增 agent 调用 trace 记录及 Web UI（端口 3001），`#337` 增加 prompt trace 日志（JSONL 格式，支持脱敏和截断），`#1333` 在日志中嵌入构建时元数据（git commit、构建时间）。
- **文档与入门**：`#2721`（待合并）新增三份自定义文档（customizing.md、skills model、skill guidelines）；`#1161` 合并 `/setup-dev` skill 简化本地开发环境搭建；`#1084` 合并容器沙箱系统设计文档；`#214` 合并安全审计报告。
- **积压清理**：`#212`（WebUI）、`#357`（Markdown 种子文件）、`#379`/`#380`（JSDoc）、`#481`（CLAUDE.md 示例）等十余条自二月起因阻塞或待审核的 PR 在本日被统一关闭或合并，表明维护者正在系统性处理积压。

## 社区热点
今日唯一的新 Issue `#1690` 是社区热点（5 条评论，3 个 👍）。该 Issue 由用户 `chiptoe-svg` 提交，提出一种多运行时抽象层，允许将不同 Agent SDK（Claude + Codex + 本地模型）作为模块化 skill 安装，沿用了现有的 `/add-telegram`、`/add-slack` 模式。其核心诉求是**让 NanoClaw 成为多后端 Agent 路由平台**，通过统一接口 `runtime.run()` 屏蔽底层差异。这一方向与项目近期的插件化、模块化趋势高度吻合，反映了社区对多模型兼容的强烈期待。

👉 [Issue #1690：Multi-runtime agent SDK abstraction](https://github.com/nanocoai/nanoclaw/issues/1690)

## Bug 与稳定性
| 严重程度 | Bug 描述 | 状态 | PR |
|---------|---------|------|----|
| **高** | 飞书通道交互卡片在 agent-runner 因 `PROCESS_TIMEOUT` 被杀死后，仍显示“运行中”超过 50 分钟，造成用户困惑和混淆。根因：`deleteActiveCard(jid)` 仅在 SDK final 事件中触发，Kill 时未清理。 | 已合并 | [#2718](https://github.com/nanocoai/nanoclaw/pull/2718) |
| **中** | Telegram 配对码使用 `Math.random`，由于输出可预测，攻击者可能观察多次配对码后计算出后续码值，进而绕过注册验证成为管理员。 | 待合并（PR #2722 已提交） | [#2722](https://github.com/nanocoai/nanoclaw/pull/2722) |

此外，`#2723`（`Finance dd agent`）虽已关闭，但类型为 Feature skill，无直接关联稳定性。

## 功能请求与路线图信号
- **多运行时支持**：Issue #1690 明确要求将不同 Agent SDK 抽象为可插拔 skill，该方向与近期已合并的 `#1285`（直接运行模式）和 `#1309`（技能市场）构成组合拳，有望成为下一版本核心能力。
- **审批工作流**：`#1245`（已合并）的 `/approve`/`/reject` 技能实现了条件授权，类似功能在其他开源 Agent 平台（如 AutoGPT、CrewAI）中正成为标配，NanoClaw 在此方向已先行一步。
- **直接运行模式**：`#1285` 使 NanoClaw 可脱离 Docker 运行，大幅降低个人开发者和小团队的使用门槛，该 PR 得到 0 评论但仍被合入，说明维护者认可其必要性。
- **文档体系完善**：多份文档 PR（#2721、#1084、#214）表明项目正系统化建设开发者入⻔和安全参考，后续版本可能推出正式文档站点。

## 用户反馈摘要
- **正面贡献**：用户 `chiptoe-svg` 在 #1690 中不仅提出需求，还提供了完整的多运行时抽象实现，并已在其 fork 中成功运行。这种“先做后提”的贡献模式体现了社区对 NanoClaw 扩展能力的高度认可。
- **使用场景痛点**：#2718 的修复源于生产环境实际遇到飞书卡片无法自动清除的问题，报告者 `brookgao` 给出了 50 分钟观察数据，反映了真实部署下 agent 进程管理的不完善，修复后直接提升用户体验。
- **易用性反馈**：多个文档 PR（#2721、#1161）来自贡献者而非核心维护者，说明社区在自行编写文档以弥补入⻔资料不足，侧面暗示项目前期文档覆盖有待加强。

## 待处理积压
- **待合并 PR**：目前有 4 条 PR 处于 OPEN 状态，其中 `#2722`（Telegram 安全修复）和 `#2721`（文档）值得优先审核合并，前者涉及安全性，后者影响新用户入⻔。另有两条未在展示列表的 PR 需维护者关注。
- **长期未响应的 Issue**：今日无新增长期未响应 Issue，但 #1690 虽然创建于 4 月 7 日，2 个多月后才于 6 月 10 日被更新，说明社区曾一度缺乏回应。建议维护者尽快就多运行时抽象方案发表官方意见（接受、拒绝或修改方向），以避免社区重复劳动。
- **已关闭但未实际解决的 PR**：`#2723`（Finance dd agent）被直接关闭，标签仅为 `follows-guidelines`，未说明拒绝理由或指引作者，可能影响贡献者积极性。建议对同类 PR 附加评论，告知原因或引导至正确提交路径。

👉 [待合并 PR #2722](https://github.com/nanocoai/nanoclaw/pull/2722)  
👉 [待合并 PR #2721](https://github.com/nanocoai/nanoclaw/pull/2721)  
👉 [待回复 Issue #1690](https://github.com/nanocoai/nanoclaw/issues/1690)

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw 项目日报 — 2026-06-10

---

## 1. 今日速览

过去24小时内项目保持着较高的活跃度：关闭了4个Issue和7个PR，其中绝大部分是Bug修复类合并，有效解决了PII误报、Telegram交互体验、自定义Provider兼容性以及配置死标志等已知问题。同时，一个新的严重Bug（#941）被报告并遗留未修，社区及时提交了对应的修复PR（#948）但尚未合并。项目整体呈“快速补丁+功能演进”态势，健康度良好，但依赖Telegram交付的Agent定时任务仍存在阻断性隐患。

---

## 2. 版本发布

（无新版本发布，本节略）

---

## 3. 项目进展

今日合并/关闭了7个PR，主要推进方向如下：

| PR | 标题 | 类型 | 对应Issue | 影响评估 |
|----|------|------|-----------|----------|
| #945 | fix(redaction): reject ISO date/time patterns as false-positive phone matches | Bug修复 | #944 | 消除PII过滤对系统日期输出的误报，提升Agent输出的可读性 |
| #946 | fix(agent): filter tools in system prompt text by tool_filter_groups | 功能优化 | – | 减少无用工具描述进入文本提示，降低token消耗与上下文污染 |
| #947 | feat(providers): add Evolink as an OpenAI-compatible provider | 新功能 | – | 扩展多模型接入能力，新增Evolink网关（支持GPT-5、Gemini等） |
| #943 | fix(telegram): show typing indicator during callback-query processing | Bug修复 | #942 | 补齐Telegram内联按钮操作的反馈体验 |
| #940 | fix(models): query base_url for custom OpenAI-compatible providers | Bug修复 | #936 | 解决自定义Provider无法获取真实模型列表的回归 |
| #939 | fix(agent): honor compact_context flag instead of always compacting | Bug修复 | #937 | 使`compact_context`配置标志真正生效，避免不必要的历史压缩 |
| #711 | Feat/cross memory | 功能合并 | – | 新增跨Agent确定性内存事件流，为多实例同步提供底层原语 |

项目在**PII处理、Telegram交互、Provider兼容性、Agent配置有效性**四个维度均有明显改进，另有一项长期开发的跨内存功能正式落地，表明核心架构仍在稳步强化。

---

## 4. 社区热点

**最活跃 Issue：** [#941 Agent-type cron jobs don't spawn a subprocess — Telegram delivery never happens](https://github.com/nullclaw/nullclaw/issues/941)  
- 作者 `weissfl` 报告了关键问题：配置为`job_type: "agent"`的定时任务被标记为完成但子进程从未启动，导致Telegram消息永远不发送。  
- 该Issue自5月31日创建后至6月9日收到1条评论，社区关注度较高（因其影响交付可靠性）。  
- 同日合并的PR #948（由 `DonPrus` 提交）直接尝试从根源修复该问题（传递cron交付元数据），但尚未合并，说明团队正在积极处理。

**另外两个前期热点议题已于今日关闭：**  
- #936（自定义Provider回退）和#942（Telegram内联按钮无打字指示）在关闭前均有社区成员参与讨论，反映用户对**Telegram集成体验**与**Provider自定义灵活性**的强烈诉求。

---

## 5. Bug 与稳定性

按严重程度排列：

| 严重程度 | Issue | 描述 | 状态 | 修复PR |
|---------|-------|------|------|--------|
| **严重（阻断）** | [#941](https://github.com/nullclaw/nullclaw/issues/941) | 配置`agent`类型定时任务时子进程未启动，Telegram交付永久缺失 | Open | 待合并：PR #948 |
| 中等 | [#944](https://github.com/nullclaw/nullclaw/issues/944) | PII过滤器将系统日期输出误判为电话号码，输出被替换为`[PHONE_X]` | Closed | 已合并：PR #945 |
| 中等 | [#936](https://github.com/nullclaw/nullclaw/issues/936) | 自定义OpenAI兼容Provider无法查询自身模型列表，回退至硬编码Claude模型 | Closed | 已合并：PR #940 |
| 低 | [#942](https://github.com/nullclaw/nullclaw/issues/942) | Telegram内联按钮缺少“正在输入…”指示器，用户无反馈等待 | Closed | 已合并：PR #943 |
| 低 | [#937](https://github.com/nullclaw/nullclaw/issues/937) | `compact_context`配置标志被解析但运行时忽略，导致始终压缩 | Closed | 已合并：PR #939 |

**总结：** 大部分已报告的Bug均于当日完成修复并合并。唯一未修复的#941是影响Telegram定时交付的阻断级问题，不过已有对应修复PR（#948），预计很快能进入合并流程。

---

## 6. 功能请求与路线图信号

- **Provider生态扩展**：PR #947 将Evolink作为一等OpenAI兼容Provider加入，表明项目有意继续扩展多模型网关支持。用户若反馈缺少其他主流提供商（如Groq、Together等），可能被纳入后续版本。  
- **MCP工具过滤**：PR #946 新增按`tool_filter_groups`过滤系统提示中的工具列表，这为动态分组MCP工具提供了更精细的控制，暗示未来可能支持更灵活的权限或场景切换。  
- **跨Agent内存同步**：PR #711 合并的“跨内存”功能提供了确定性事件流，路线图上可能导向**多Agent协作**或**状态持久化**的高级特性。  
- **定时任务可靠性**：#941 暴露了Agent定时任务交付的脆弱性，社区提交的修复PR #948 虽然针对cron交付元数据，但说明维护团队正将**cron子系统稳定性**作为短期重点。

---

## 7. 用户反馈摘要

从近期Issue讨论中提炼的真实使用痛点：

- **“定时任务无声失败”**（#941）：`weissfl` 描述了完整的配置过程——设置`schedule`、`job_type: "agent"`、`delivery_mode: "always"`、`delivery_channel: "telegram"`，结果任务标记为完成但消息从未到达。用户预期“子进程启动并发送消息”，实际是“静默空转”。  
- **“时间输出变电话号码”**（#944）：`vernonstinebaker` 反馈当Agent执行`date`命令时，输出中的时间格式被PII红色遮蔽成`[PHONE_X]`，导致系统无法向用户提供准确时间信息。用户认为该功能“过于激进”。  
- **“自定义Provider无用”**（#936）：`weissfl` 配置了任意名称的OpenAI兼容Provider并指定`base_url`，但在Telegram`/models`菜单中看到的仍是硬编码的Claude模型名。用户需要手动绕过或通过其他方式查询真实模型。  
- **“内联按钮无反馈”**（#942）：`weissfl` 对比了普通文本消息（有打字指示）和内联按钮（无指示），认为后者让用户感觉“卡死”或“未响应”。

所有反馈均已被对应的修复PR解决或正在解决，社区参与质量较高。

---

## 8. 待处理积压

| 项目 | 创建时间 | 最后更新 | 说明 | 链接 |
|------|----------|----------|------|------|
| **#941 (Issue)** | 2026-05-31 | 2026-06-09 | Agent定时任务子进程不启动，影响Telegram交付，严重度阻断 | [Issue #941](https://github.com/nullclaw/nullclaw/issues/941) |
| **#948 (PR)** | 2026-06-10 | 2026-06-10 | 由`DonPrus`提交的修复PR，传递cron交付元数据，意图修复#941 | [PR #948](https://github.com/nullclaw/nullclaw/pull/948) |

**建议：** 优先审查并合并PR #948，以关闭#941这一影响Telegram定时交付的关键Bug。此外，#711（跨内存）虽已关闭，但涉及核心架构变更，建议关注上线后是否引入回归。目前无其他长期未响应的重要Issue或PR。

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

# IronClaw 项目日报 - 2026-06-10

## 1. 今日速览

过去24小时项目维持**极高活跃度**：共处理7个Issue和50个Pull Request，其中46个PR仍处于待合并状态，4个已合并/关闭。没有新版本发布，但大量功能性的巨型PR（XL/L）集中推进，覆盖附件处理、项目所有权模型、审批策略持久化、可观测性等关键领域。同时，社区报告的Gmail工具缺失归档操作、NEAR AI提供者配置保存失败等Bug已获新Issue记录，部分已有修复PR在审。整体来看，项目正处于**大规模功能迭代替换旧架构**的阶段，代码库健康度一般（存在长期E2E失败与文件大小超线问题），但团队响应迅速。

## 2. 版本发布

无

## 3. 项目进展

今日合并/关闭了2个重要PR和1个关联Issue，标志着多个子功能走向稳定：

- **PR #4600** [CLOSED] – **Slack个人DM外发目标实现**（Phase 4 C2）。在现有共享频道目标库基础上增加以用户为粒度的个人DM目标，后端合约保持频道中立，为后续自动化触达提供基础。  
  [nearai/ironclaw#4600](https://github.com/nearai/ironclaw/pull/4600)

- **PR #4553** [CLOSED] – **WebUI v2 国际化（i18n）补全**。移除扩展、任务、常规、侧边栏中残留的硬编码英文字符串，补全缺失翻译键，修复部分区域运行时崩溃。  
  [nearai/ironclaw#4553](https://github.com/nearai/ironclaw/pull/4553)

- **Issue #4554** [CLOSED] – 关联前述PR，i18n覆盖不完整及翻译运行时崩溃问题已解决。  
  [nearai/ironclaw#4554](https://github.com/nearai/ironclaw/issues/4554)

此外，**附件支持**的多条追踪PR（#4644）今日密集入栈：  
- `ironclaw_extractors` 独立crate（#4675）、文档文本提取（#4676）、附件模型可见上下文注入（#4677）、附件存储落地层（#4668）均在审查中，表明附件全链路即将完成。

## 4. 社区热点

过去24小时单个Issue/PR的评论数均较低，但以下条目因其**功能性缺口或用户触发频率**成为社区关注焦点：

- **Issue #4674** – Gmail工具缺少 `archive` 和 `label-modify` 操作（只能trash），用户通过WASM枚举验证了缺失。这是对核心工具功能的直接诉求，可能影响重度邮件用户。  
  [nearai/ironclaw#4674](https://github.com/nearai/ironclaw/issues/4674)

- **Issue #4673** – NEAR AI提供者配置在成功“测试连接”后无法保存，导致用户返回欢迎页仍显示未配置。该问题阻止新用户完成onboarding，优先级较高。  
  [nearai/ironclaw#4673](https://github.com/nearai/ironclaw/issues/4673)

- **PR #3708** – 持续开放超过三周的版本发布PR，涉及 `ironclaw_common` 和 `ironclaw_skills` 的API破坏性变更，尽管本周继续更新，但尚未合并，阻塞下游依赖方升级。  
  [nearai/ironclaw#3708](https://github.com/nearai/ironclaw/pull/3708)

## 5. Bug 与稳定性

| 严重度 | Issue / PR | 描述 | 状态 | 关联修复 |
|--------|-------------|------|------|----------|
| 高 | #4108 | Nightly E2E调度运行持续失败（自5月27日起），至今已14天未修复，影响回归信心。 | OPEN | 无 |
| 高 | #4673 | NEAR AI提供者配置保存失败，用户无法完成onboarding。 | OPEN | 无（今日新建） |
| 中 | #4680 (PR) | OpenAI兼容API中非文本内容被压缩为 `[non_text_content]` 占位符，导致模型接收错误文本。 | OPEN（待合并） | #4680 修复 |
| 低 | #4554 | i18n覆盖不完整及运行时崩溃，已通过#4553合并修复。 | CLOSED | ✅ 已修复 |

长期稳定性隐患：`slack_host_state.rs`（2,823行）接近文件大小上限，`slack_host_beta.rs`（3,359行）已超标，团队已建立追踪Issue（#4666、#4665）并要求后续PR不得增加行数。

## 6. 功能请求与路线图信号

今日新增的功能请求/设计信号：

- **Gmail archive/label-modify**（#4674）—— 单工具扩展需求，可能被纳入下一次工具更新。  
- **Ask-gated能力审批**（#4667）—— Reborn REPL无法处理需要人工授权的capability，解决后才可支持高级安全场景。  
- **项目范围所有权模型**（#4662、#4663、#4664）—— 今日提交了系列PR，从设计文档、核心模型到产品表面重命名，标志着多租户/项目级自动化控制即将落地，是下一版本的核心路线图项。  
- **持久化审批策略**（#4613）—— 支持 `AlwaysAllow` 的scope式存储，为未来产品化控制面板铺路。  
- **Trace Commons agent驱动加入**（#4559）—— 通过邀请链接简化注册流程，提升可观测性平台的易用性。  
- **附件全链路**（#4644系列）—— 从存储、提取到注入模型上下文的完整方案，预计将大幅改善文档/图片对话体验。

以上功能多数已有对应PR在审查中，有望在下一版本（0.30.0）中批量合入。

## 7. 用户反馈摘要

基于过去24小时Issue评论（仅#4674有1条评论，内容未提供）及Issue描述，可提炼以下用户痛点：

- **Gmail工具能力不足**：用户明确表示只能 `trash` 而无法 `archive` 或管理标签，与实际工作流脱节，期望增加与 `trash` 并列的 `archive` 和 `label-modify` 操作。  
- **NEAR AI配置体验断裂**：用户通过“测试连接”成功验证，但“保存”按钮静默失败，返回后配置未持久化，易造成困惑和信任损耗。  
- **Reborn REPL错误反馈模糊**：PR #4679 的动机即用户看到 `(no assistant text; ...)` 而非具体失败原因，说明当前诊断信息不足，影响调试效率。

## 8. 待处理积压

- **Issue #4108** – Nightly E2E 持续失败（自2026-05-27），至今无任何进展或公关，严重威胁持续集成稳定性。  
  [nearai/ironclaw#4108](https://github.com/nearai/ironclaw/issues/4108)

- **PR #3708** – 版本发布PR（`chore: release`）已开放25天，包含 `ironclaw_common` 和 `ironclaw_skills` 的破坏性API变更，建议维护者评估风险后尽快合并或关闭，避免阻塞社区升级。  
  [nearai/ironclaw#3708](https://github.com/nearai/ironclaw/pull/3708)

- **PR #4501** – Dependabot配置更新优化（避免因 `unknown_error` 失败），自2026-06-05起无新活动，可安排合并。  
  [nearai/ironclaw#4501](https://github.com/nearai/ironclaw/pull/4501)

---

*生成时间：2026-06-10 基于 GitHub 公开数据自动分析*

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

**LobsterAI 项目动态日报 (2026-06-10)**  

---

## 1. 今日速览  

过去 24 小时内，LobsterAI 项目保持中等活跃度：共处理 4 个 Pull Request（全部合并/关闭），新增 1 个 Issue。无新版本发布。核心进展集中在渲染器导出修复、数据备份功能实现及主窗口恢复机制。社区唯一新增 Issue 关注跨模型子任务协作机制，提出明确的优化方向。整体来看，项目在稳定修复与功能增强上均有推进，社区反馈聚焦于多智能体协同场景。  

---

## 2. 版本发布  

无新版本发布。  

---

## 3. 项目进展  

今日合并/关闭的 4 个 PR 中，涉及渲染器、协同工作（cowork）、文档及主模块多个领域：  

- **`#2133`** [CLOSED] [area: renderer, area: cowork] fix: fix export and code copy bugs  
  修复了导出和代码复制相关的 Bug，提升了渲染器与协同工作区的可用性。  
  → [netease-youdao/LobsterAI PR #2133](https://github.com/netease-youdao/LobsterAI/pull/2133)  

- **`#2136`** [CLOSED] [area: renderer, area: docs, area: main] feature: data backup and migration  
  新增**数据备份与迁移**功能，覆盖渲染器、文档及主模块，为用户提供数据安全与跨环境移动能力。  
  → [netease-youdao/LobsterAI PR #2136](https://github.com/netease-youdao/LobsterAI/pull/2136)  

- **`#2135`** [CLOSED] [area: renderer] chore: temporary close databackup  
  临时关闭数据备份功能（可能为后续重构或测试做准备），属于工程清理动作。  
  → [netease-youdao/LobsterAI PR #2135](https://github.com/netease-youdao/LobsterAI/pull/2135)  

- **`#2134`** [CLOSED] [area: renderer, area: docs, area: main] Liuzhq/task complete notice  
  实现**任务完成通知优化**：主窗口关闭或销毁后，任务完成时可恢复 LobsterAI 窗口；等待渲染器通知处理器就绪后再打开目标协同会话；保持 macOS 通知中心点击动作的有效性。显著提升了协同场景下的用户恢复体验。  
  → [netease-youdao/LobsterAI PR #2134](https://github.com/netease-youdao/LobsterAI/pull/2134)  

**小结**：今日 PR 密集合并，除了一项临时关闭操作外，均属功能性增强或关键 Bug 修复，项目在数据处理、窗口协调和通知稳定性方面迈出扎实一步。  

---

## 4. 社区热点  

今日唯一活跃的 Issue 为 **`#2132`**，讨论跨模型子任务调用问题：  

- **作者**: woxinsj  
- **状态**: OPEN（无评论、无点赞）  
- **核心诉求**：  
  - 主任务（如 M3 模型）擅长规划与验收监督，子任务（如 deepseek 模型）擅长快速执行，但当前跨模型调用机制存在缺陷。  
  - 提出优化方案：借鉴同模型子任务的通知机制（子任务完成后主任务第一时间知晓）；子任务完成或卡点时主动对接主任务，并通过明确的跨模型调用要求落地。  
  - 摘要中展示了诊断过程：发现函数调用 `call_function_gblu0nmqpcej_1` 不在 sessions_list 和 subagents 中，属于网关级函数调用，而非 `sessions_spawn` 创建的子任务。  

**分析**：该 Issue 反映了用户对**多模型协作**的强烈需求，且已给出具体技术分析和改进建议。虽然目前无回复，但内容质量高，值得项目组优先关注。  
→ [netease-youdao/LobsterAI Issue #2132](https://github.com/netease-youdao/LobsterAI/issues/2132)  

---

## 5. Bug 与稳定性  

- **已修复 Bug**：`#2133` 修复了导出和代码复制 Bug，直接提升用户体验。  
- **临时关闭**：`#2135` 临时关闭数据备份功能，可能暗示该功能在某场景下存在稳定性风险（如崩溃、资源泄漏等），需关注后续重构版本。  
- **无新报告的严重 Bug**：今日未出现崩溃、回归或数据丢失类报告。  

总体稳定性良好，但数据备份模块的临时关闭需要项目组在后续 PR 中说明具体原因。  

---

## 6. 功能请求与路线图信号  

- **跨模型子任务协作**（`#2132`）：用户明确提出了同模型通知机制借鉴、主动对接等设计思路，属于多智能体协作的核心能力。结合已有 PR `#2134`（任务完成通知）的优化，项目可能在下一版本中增强跨模型间的消息传递。  
- **数据备份与迁移**（`#2136` 已合并）：已落地为正式功能，预计将在下个 Release 中公开。  
- **任务完成通知恢复窗口**（`#2134` 已合并）：针对 macOS 用户优化，属于桌面端体验的重要补全。  

这些特性暗示项目正向**多智能体协作 + 可靠持久化**方向演进。  

---

## 7. 用户反馈摘要  

基于 `#2132` 的 Issue 内容，提炼核心用户痛点：  

- **使用场景**：用户期望用一个“M3 模型”作为主任务负责规划与验收，同时用“deepseek 模型”作为子任务快速执行具体操作。  
- **痛点**：当前跨模型调用时，子任务的完成状态无法像同模型那样自动回传至主任务，需要人工介入或导致协作卡顿。  
- **满意点**：用户认可项目现有的同模型子任务通知机制（“这个机制值得借鉴”），表明基础架构设计合理。  

无其他负面反馈记录。  

---

## 8. 待处理积压  

当前无长期未响应的重要 Issue 或 PR。仅 `#2132` 处于 OPEN 状态，创建于 2026-06-09，距今 1 天，尚未有项目组回复。建议维护者尽快参与讨论，尤其是用户已提供根因分析及修复方案，可加速问题解决。  

---

*报告生成时间：2026-06-10 23:59 UTC*  
*数据来源：netease-youdao/LobsterAI GitHub repository*

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis 项目动态日报 – 2026-06-10

## 1. 今日速览
- 项目过去24小时整体活跃度极低：仅产生1条新 Issue，无 Pull Request 提交或合并，无新版本发布。
- 唯一的新 Issue 为 `[Bug]: provider 'coqui' not configured`，严重程度标记为 minor，尚未获得任何评论或点赞。
- 从数据看，社区贡献与维护者响应均处于静默状态，项目当前无明显推进或修复活动。
- 待办积压无新增，但长期健康度需关注 Issue 反馈的响应效率。

---

## 2. 版本发布
无

---

## 3. 项目进展
- 今日无 Pull Request 被合并或关闭，项目代码库无新增提交。功能开发与 Bug 修复均无可见进展。

---

## 4. 社区热点
- 今日无活跃讨论或高互动 Issue/PR。唯一 Issue [#1114](https://github.com/moltis-org/moltis/issues/1114) 为零评论、零点赞，尚未引发社区关注。潜在诉求：用户在使用 Coqui 语音提供商时遭遇配置缺失，可能期望项目提供更清晰的文档或自动检测机制。

---

## 5. Bug 与稳定性
| Issue | 严重程度 | 描述 | 是否有修复 PR |
|-------|----------|------|---------------|
| [#1114](https://github.com/moltis-org/moltis/issues/1114) | minor | `provider 'coqui' not configured` – 用户无法使用 Coqui TTS 提供商，提示未配置。 | 无 |

**分析**：该 Bug 可能源于配置校验不完善或环境变量缺失。轻微影响使用特定 TTS 后端的用户，但无崩溃或数据丢失风险。

---

## 6. 功能请求与路线图信号
- 今日无新增功能请求 Issue。但结合现有 Bug 反馈，Coqui 提供商的配置问题可能暗示用户对多 TTS 后端支持有隐性需求，未来路线图或需强化 provider 配置的容错与默认行为。

---

## 7. 用户反馈摘要
- 从 [#1114](https://github.com/moltis-org/moltis/issues/1114) 的描述可看出用户已完成基本排查（搜索已有 Issue、使用最新版本），但未在聊天会话中保留上下文。这表明用户对配置流程的直觉性期望与实际实现存在差距，建议维护者在文档或错误提示中明确列出可用的 provider 及配置方法。

---

## 8. 待处理积压
- 今日无长期未响应的 Issue 或 PR。唯一 Issue #1114 为当日新开，尚未超过合理响应窗口。仍需关注其后续的维护者回应与解决方案。

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw 项目日报 (2026-06-10)

## 今日速览

项目今日保持高活跃度：过去24小时共处理 **25 条 Issue**（新开/活跃13，关闭12）与 **35 条 PR**（待合并19，已合并/关闭16），并发布了 **v1.1.11-beta.2** 版本。主要亮点包括：**AgentScope 2.0 迁移讨论**持续深入（#4727），社区对 **Hermes Agent 的“学习循环”特性**高度关注（#5017），同时多项关键 Bug 如 MCP 子进程累积、微信推送失败已完成修复。整体项目健康度良好，迭代节奏稳定。

---

## 版本发布

### v1.1.11-beta.2
- **发布时间**：2026-06-10  
- **主要更新**：
  - **浏览器控制增强**：新增页面坐标点击能力（`browser_control`）  
  - **跨浏览器切换优化**：添加 CDP 超时参数，实现浏览器配置文件隔离  
  - 详见 [Release 页面](https://github.com/agentscope-ai/QwenPaw/releases/tag/v1.1.11-beta.2)
- **破坏性变更**：无明确说明，属于功能补丁
- **迁移注意**：由于是 beta 版本，建议生产环境用户谨慎升级，测试确认兼容性后再部署。

---

## 项目进展

今日共合并/关闭 **16 个 PR**，以下为关键推进：

| PR | 标题 | 状态 | 说明 |
|----|------|------|------|
| [#4981](https://github.com/agentscope-ai/QwenPaw/pull/4981) | fix(security): restrict file preview to WORKING_DIR and block sensitive paths | 已合并 | 修复文件预览路径穿越漏洞，增强安全边界 |
| [#5033](https://github.com/agentscope-ai/QwenPaw/pull/5033) | feat(plugin/cloudpaw): support importing agents from AgentHub and enhance A2A capabilities | 已合并 | 首次贡献者 PR，扩展插件生态，支持从 AgentHub 导入代理并增强 Agent-to-Agent 能力 |
| [#4969](https://github.com/agentscope-ai/QwenPaw/pull/4969) | feat(skill): Add skill tag batch download | 已合并 | 实现技能按标签批量下载，对应 Issue #2961（技能分类功能） |
| [#5014](https://github.com/agentscope-ai/QwenPaw/pull/5014) | fix(mcp): prevent subprocess accumulation across restarts (#4834) | 已合并 | 彻底修复 MCP 子进程在 Docker 重启后残留的问题 |
| [#5062](https://github.com/agentscope-ai/QwenPaw/pull/5062) | fix(e2e): handle text-based empty state in token usage test | 已合并 | 修复 E2E 测试中空状态判定的兼容性问题 |

项目整体在 **安全加固、插件生态、技能管理、稳定性** 四个维度均有实质性进展。

---

## 社区热点

### 1. 🔥 建议关注 Hermes Agent 发展 (#5017，评论10，👍3)
**链接**：[#5017](https://github.com/agentscope-ai/QwenPaw/issues/5017)  
**核心诉求**：用户强烈建议项目关注 **Hermes Agent** 的“学习循环（Learning Loop）”——Agent 能自动从自身行为中创建并迭代技能。该功能若引入将显著提升 CoPaw 的自主进化能力。  
**社区反应**：评论区10条讨论，用户对 CoPaw 本地化体验高度肯定，同时表达了“博采众长”的开放期待。

### 2. 🚀 AgentScope 2.0 迁移讨论 (#4727，评论8，👍2)
**链接**：[#4727](https://github.com/agentscope-ai/QwenPaw/issues/4727)  
**核心诉求**：后端从 AgentScope 1.x 迁移至 2.0 的 breaking change 提案，涉及架构、API、运行时模型全面升级。社区关注迁移后的兼容性与预期收益（性能、可扩展性）。  
**当前状态**：已进入方案细化阶段，预计将影响后续多个版本。

### 3. 🐛 会话模型配置丢失 Bug (#4666，评论7)
**链接**：[#4666](https://github.com/agentscope-ai/QwenPaw/issues/4666)  
**背景**：新建会话后 Models 配置页面丢失且无法加载，只能重启解决。用户投诉影响日常使用，该 Issue 今日已关闭（疑似修复或定位）。

---

## Bug 与稳定性

按严重程度排列，标注修复进展：

| 严重程度 | Issue | 描述 | 状态 |
|----------|-------|------|------|
| 🔴 **阻塞** | [#5065](https://github.com/agentscope-ai/QwenPaw/issues/5065) | `MODEL_EXECUTION_FAILED` 频繁报错，导致任务完全无法执行 | OPEN，无 PR |
| 🔴 **阻塞** | [#5064](https://github.com/agentscope-ai/QwenPaw/issues/5064) | Agent 创建的定时任务无法自动触发且无法手动编辑 | OPEN，无 PR |
| 🟠 **功能性** | [#4878](https://github.com/agentscope-ai/QwenPaw/issues/4878) | 微信定时任务推送失败（`send_text rejected`） | 已关闭，关联 PR [#5060](https://github.com/agentscope-ai/QwenPaw/pull/5060) 合并修复 |
| 🟠 **功能性** | [#4834](https://github.com/agentscope-ai/QwenPaw/issues/4834) | MCP 子进程跨重启累积导致加载缓慢 | 已关闭，PR [#5014](https://github.com/agentscope-ai/QwenPaw/pull/5014) 合并修复 |
| 🟡 **体验** | [#4993](https://github.com/agentscope-ai/QwenPaw/issues/4993) | 图片预览放大后拖动异常抖动 | 已关闭 |
| 🟡 **体验** | [#5031](https://github.com/agentscope-ai/QwenPaw/issues/5031) | 技能斜杠调用在 Console 中显示为完整 SKILL.md 内容 | OPEN，无 PR |
| 🟡 **体验** | [#4170](https://github.com/agentscope-ai/QwenPaw/issues/4170) | Agent 动作信息仅在全部完成后才展示，导致无法及时干预 | 已关闭 |
| ⚪ **稳定性** | [#4006](https://github.com/agentscope-ai/QwenPaw/issues/4006) | OpenAI 兼容 provider 未过滤推理内容 | 已关闭 |
| ⚪ **兼容性** | [#4989](https://github.com/agentscope-ai/QwenPaw/issues/4989) | 使用本地千问3.6-27B 模型时对话无响应（1.1.9/1.1.10） | OPEN，无 PR |

**关键发现**：今日新增 2 个阻塞级 Bug（#5065、#5064），均与模型执行和定时任务相关，需优先响应。

---

## 功能请求与路线图信号

### 高潜力 Feature（可能被纳入下一版本）

| Issue/PR | 描述 | 信号强度 |
|----------|------|----------|
| [#5017](https://github.com/agentscope-ai/QwenPaw/issues/5017) | 借鉴 Hermes Agent 的学习循环功能 | ⭐⭐⭐ 社区呼声极高 |
| [#5063](https://github.com/agentscope-ai/QwenPaw/issues/5063) | 集成 Headroom 实现上下文压缩，降低 60-95% Token 消耗 | ⭐⭐⭐ 刚提出，但价值明确 |
| [#3751](https://github.com/agentscope-ai/QwenPaw/issues/3751) | Windows 系统托盘支持 | ⭐⭐ 长期未解决，但 PR [#4669](https://github.com/agentscope-ai/QwenPaw/pull/4669) 已进入 Tauri 自动更新方向 |
| [#4057](https://github.com/agentscope-ai/QwenPaw/issues/4057) | 统一 AgentScope tracing 初始化入口 | ⭐⭐ 为集成监控平台铺路 |
| [#4975](https://github.com/agentscope-ai/QwenPaw/pull/4975) | 会话页面可自定义列顺序 | ⭐⭐ 已提交 PR，有合并潜力 |
| [#2961](https://github.com/agentscope-ai/QwenPaw/issues/2961) | 技能分类（文件夹管理） | ✅ 已通过 PR [#4969](https://github.com/agentscope-ai/QwenPaw/pull/4969) 实现标签批量下载，基础分类功能已完成 |

### 路线图信号
- **AgentScope 2.0 迁移**（#4727）将是未来 1-2 个月最重大的架构升级，涉及多个模块重构。
- **插件生态扩展**：PR [#5033](https://github.com/agentscope-ai/QwenPaw/pull/5033) 和 [#5066](https://github.com/agentscope-ai/QwenPaw/pull/5066) 表明团队正在推进前端插件文档化和 A2A 能力。

---

## 用户反馈摘要

### 正面评价
- **本地化体验优秀**（#5017）：“国内用起来特别舒服——本地化做得很到位，设置清晰无门槛，开箱即用。”
- **技能管理进步**：用户在 #2961 中提出技能分类需求，PR #4969 实现后社区反响积极。

### 痛点与不满
1. **桌面端启动极慢**（#5047）：切换到 Tauri 后启动时间从 1-2 分钟增至十几分钟，且易进入无响应状态。影响 Windows 用户基础体验。
2. **大对话卡死**（#4213、#4917）：当历史消息过多（百万级 Token）时，网页重新加载或切换页面导致严重卡顿，用户期望分片传输或分页。
3. **工具调用多次后失效**（#5052）：前几次工具调用成功，若干轮后全部报 `got an unexpected keyword argument 'arguments'`，疑似状态污染或 SDK 版本冲突。
4. **定时任务不可编辑**（#5064）：Agent 生成的定时任务无法手动修改时间或内容，灵活性不足。
5. **模型配置丢失**（#4666）：新建会话导致配置页面崩溃，需重启解决，严重影响连续工作流。

### 建议期望
- 用户频繁提及 **Hermes Agent** 的自动技能学习能力，希望 CoPaw 能借鉴。
- 对 **AgentScope 2.0 迁移**持谨慎乐观态度，但要求保证向后兼容。
- 多个用户请求 **上下文压缩/Token 节省** 功能（#5063），表明大模型成本仍是痛点。

---

## 待处理积压

以下 Issue/PR 长期未得到维护者响应或停滞，建议关注：

| 类别 | 编号 | 标题 | 创建时间 | 最后更新 | 说明 |
|------|------|------|----------|----------|------|
| Issue | [#3751](https://github.com/agentscope-ai/QwenPaw/issues/3751) | Windows 系统托盘功能 | 2026-04-23 | 2026-06-10 | 已超过1.5个月无实质性进展，虽有 PR [#4669](https://github.com/agentscope-ai/QwenPaw/pull/4669) 但未合并 |
| PR | [#4669](https://github.com/agentscope-ai/QwenPaw/pull/4669) | feat(desktop): add tauri auto updater | 2026-05-25 | 2026-06-10 | 长期 OPEN，缺少 review，影响桌面端自动更新能力 |
| Issue | [#4057](https://github.com/agentscope-ai/QwenPaw/issues/4057) | 支持 AgentScope tracing 初始化 | 2026-05-06 | 2026-06-10 | 评论数4，无 assignee，技术细节清晰但未进入开发 |
| Issue | [#4989](https://github.com/agentscope-ai/QwenPaw/issues/4989) | 本地千问3.6-27B 对话无响应 | 2026-06-06 | 2026-06-09 | 影响使用本地模型的用户，尚无社区或官方回复 |

---

**总结**：CoPaw 项目今日在 Bug 修复和功能完善上取得扎实进展，社区活跃度处于高位。需要警惕两个新阻塞 Bug 以及桌面端性能问题，同时尽快响应学习循环、上下文压缩等高价值功能请求，以巩固用户满意度。

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，作为 ZeroClaw 项目的 AI 智能体与个人 AI 助手领域开源项目分析师，根据您提供的 2026-06-10 数据，我已为您生成了以下项目动态日报。

---

# ZeroClaw 项目动态日报 | 2026-06-10

## 1. 今日速览

ZeroClaw 项目今日保持高活跃度，社区贡献者提交了大量的 Pull Request，项目在稳定性修复和功能增强上齐头并进。核心关注点为：内存管理、安全性（多租户 RBAC、MCP 工具权限）、以及平台兼容性（Discord、Telegram、macOS TUI）。虽然今日无新版本发布，但社区提交的 47 个待合并 PR 表明项目正处于密集的功能迭代期，项目健康度良好，合并率（MRM）有待提升。

## 2. 版本发布

**无**

## 3. 项目进展

过去24小时内，项目完成了少数但关键的 PR 合并与关闭，显示出核心维护者在处理社区贡献上的谨慎态度。

- **[已关闭] PR #7362 feat(zerocode): hold back context-window usage bar until it bakes**
    - **内容**: 暂时隐藏了 `zerocode` TUI 中的上下文窗口使用量进度条，因其会计功能尚不成熟，避免导致用户困惑。
    - **意义**: 体现了对用户体验的重视，避免了过早暴露不准确数据可能带来的误导。
- **[已关闭] PR #7369 feat(channels/sop): AMQP inbound channel with mutual TLS...**
    - **内容**: 引入了一个通过 AMQP（高级消息队列协议）与双向 TLS（传输层安全性协议）的完整、可部署的 SOP（标准操作程序）用例场景。
    - **意义**: 为 ZeroClaw 在自动化运维和企业级集成场景打下了坚实基础，扩展了项目的应用边界。
- **[已合并/关闭] - 另有 1 条 PR 在统计中被合并/关闭，具体信息未在提供的数据中详细列出。**

**项目向前迈进**：尽管合并数量不多，但 `AMQP` 通道的引入和 `SOP` 流程的完善是项目架构上的一个里程碑。同时，大量悬而未决的 PR 也为项目的下一个大版本积累了丰富的特性。

## 4. 社区热点

今日社区讨论的热点主要集中在几个长期演进的功能和安全议题上。

- **Multi-tenant RBAC（多租户角色权限控制） - Issue #5982** ([链接](https://github.com/zeroclaw-labs/zeroclaw/issues/5982))
    - **热度**: 9条评论，被标记为 `status:blocked` 且 `needs-author-action`。
    - **核心诉求**: 用户 `metalmon` 提出的“为单个ZeroClaw实例添加可选的、基于发送者的角色权限控制”功能，引发了广泛讨论。该功能旨在支持不同用户类（客户、操作员、开发者）拥有隔离的工作空间、工具集、速率限制和系统提示。这反映出社区对ZeroClaw在企业级多租户部署场景下的安全与隔离需求日益增长。

- **Per-skill Security Permissions（技能级安全权限） - Issue #5775** ([链接](https://github.com/zeroclaw-labs/zeroclaw/issues/5775))
    - **热度**: 3条评论。
    - **核心诉求**: 详细讨论了一个安全问题：`allow_scripts` 和 `allowed_commands` 是全局配置，启用后对所有已安装的技能开放。社区希望引入“per-skill”的细粒度安全权限，以确保安装一个需要 `python3` 的技能不会使系统暴露在巨大的安全风险下。这与上面 #5982 的讨论方向一致，都指向了权限的细分和隔离。

- **Image Info & Vision Models（图像信息与视觉模型） - PR #7446** ([链接](https://github.com/zeroclaw-labs/zeroclaw/pull/7446))
    - **热度**: 最新提交的 PR，快速吸引了注意。
    - **核心诉求**: 修复了 `image_info` 工具在交付图像给视觉模型时，仅支持绝对POSIX路径的问题。这表明社区在积极打磨多模态交互能力，尤其是与视觉模型的结合，这是一个非常受欢迎的功能领域。

## 5. Bug 与稳定性

今日报告了多个严重程度较高的 Bug，尤其集中在内存管理、上下文预算、工具权限和特定平台兼容性上。部分已有对应的修复 PR。

- **严重等级 S1 (工作流阻塞)**
    - **[Bug] #5844: Too much emphasis on memory** ([链接](https://github.com/zeroclaw-labs/zeroclaw/issues/5844))
        - **问题**: 系统提示对记忆（memory）的权重过高，导致在处理当前 prompt 时（尤其是在 cron 任务中）不准确。
        - **现状**: 被标记为 `accepted`，等待修复。
    - **[Bug] #6002: Not clearly addressed to the assistant (Telegram)** ([链接](https://github.com/zeroclaw-labs/zeroclaw/issues/6002))
        - **问题**: 在 Telegram 频道中，ZeroClaw 有时无法识别消息是明确发给它的，导致不响应。
        - **现状**: 被标记为 `accepted`，等待修复。
    - **[Bug] #5808: Default 32k context budget exceeded on iteration 1** ([链接](https://github.com/zeroclaw-labs/zeroclaw/issues/5808))
        - **问题**: 默认的 32k 上下文预算在第一次对话迭代中就被系统提示和工具定义超出约3.3倍，导致模型不断被迫裁剪上下文（Preemptive Trim），严重影响对话质量。
        - **现状**: 被标记为 `accepted`，等待修复。
    - **[Bug] #6646: web_search_tool and web_fetch not firing via Telegram** ([链接](https://github.com/zeroclaw-labs/zeroclaw/issues/6646))
        - **问题**: 在 v0.7.5 版本中，通过 Telegram 频道使用 `web_search_tool` 和 `web_fetch` 工具失效。
        - **现状**: 被标记为 `accepted`，等待修复。
    - **[Bug] #6687: Two independent SopEngine instances** ([链接](https://github.com/zeroclaw-labs/zeroclaw/issues/6687))
        - **问题**: 单个 ZeroClaw 守护进程创建了两个独立的 SOP 引擎实例（一个给 agent，一个给 MQTT），导致通过 MQTT 启动的 SOP 运行状态对 agent 不可见。
        - **现状**: 被标记为 `accepted`，等待修复。

- **严重等级 S2 (行为降级)**
    - **[Bug] #7376: zerocode Dashboard hides unavailable/error states** ([链接](https://github.com/zeroclaw-labs/zeroclaw/issues/7376))
        - **问题**: zerocode 仪表盘隐藏了不可用/错误状态，并将历史会话标签为活跃会话，存在误导性。
        - **现状**: **已有修复 PR #7444**，状态为 OPEN。
    - **[Bug] #7377: zerocode dark themes inherit unreadable foreground text** ([链接](https://github.com/zeroclaw-labs/zeroclaw/issues/7377))
        - **问题**: 在浅色终端配置文件下使用深色主题时，文本不可读。
        - **现状**: 被标记为 `accepted`，等待修复。

## 6. 功能请求与路线图信号

今日社区提出的新功能需求反映了对安全性、可观测性和集成度的更高要求。

- **Webhook Signing Secrets 动态读取 - Issue #7410** ([链接](https://github.com/zeroclaw-labs/zeroclaw/issues/7410))
    - **诉求**: 建议网关的 Webhook 签名密钥不要在启动时缓存，而应在处理时从 `AppState.config` 动态读取，以支持配置的热更新，避免因配置变更而重启服务。
    - **路线图信号**: 这指向了项目的运维友好性，未来可能会加强对配置文件动态加载的支持。其关联的 PR #7367（通道别名路由）也将进一步完善 Webhook 功能。

- **Persist Cached Input Tokens for Cost Accounting - Issue #7248** ([链接](https://github.com/zeroclaw-labs/zeroclaw/issues/7248))
    - **诉求**: 请求在运行时路径中持久化缓存输入令牌（Cached Input Tokens）的数据，并将其纳入成本核算。
    - **路线图信号**: 这表明社区对精确的成本控制和审计需求正在上升，未来可能在记账、计费模块上有所增强。

## 7. 用户反馈摘要

从 Issues 和 PR 的评论中，可以提炼出以下真实用户痛点：

- **“我的上下文预算总是不够用”**: 用户 `JordanTheJet` 在 Issue #5808 中详细描述了默认32k上下文预算在第一次对话迭代中就爆炸的问题，这极大地阻碍了复杂任务（如大型文档处理、多步骤工具调用）的执行。
- **“配置陷阱太多”**: 用户 `nick-pape` 在 Issue #6721 中指出 `tool_search` 不在 `default_auto_approve()` 列表中，导致在非交互式 Webhook 模式下，MCP 工具的加载会静默挂起2分钟然后自动拒绝，这是一个深度的配置陷阱，严重影响自动化体验。
- **“多平台支持仍需打磨”**: 用户在 Issue #7376 至 #7400 中反馈了 `zerocode` TUI 在 macOS 上的多个问题，如 `Cmd-C` 被误认为是退出快捷键、深色主题下的文本可读性问题等，说明跨平台兼容性仍是社区关注的重点。
- **“自定义Provider的集成体验需要优化”**: 用户 `databillm` 在 Issue #7439 和 PR #7441 中提到，自定义 Provider 在实际工作中运行正常，但 ZeroClaw 的 Doctor 工具会错误地报告其为“无效”，这给用户带来了困扰，表明项目的诊断工具与配置系统之间可能存在脱节。

## 8. 待处理积压

以下 Issue 和 PR 长时间未得到关闭或明确响应，建议项目维护者关注：

- **Issue #4853: Installing skills from `.well-known` URI** ([链接](https://github.com/zeroclaw-labs/zeroclaw/issues/4853))
    - **创建时间**: 2026-03-27
    - **标签**: `enhancement`, `status:accepted`
    - **状态**: 社区期待已久的功能，关于如何标准化安装技能，但至今未有实质性进展。虽然已被接受，但长期未转交开发。

- **Issue #5982: Per-sender RBAC for multi-tenant** ([链接](https://github.com/zeroclaw-labs/zeroclaw/issues/5982))
    - **创建时间**: 2026-04-22
    - **标签**: `status:blocked`, `needs-author-action`
    - **状态**: 这个重要功能请求处于“阻塞”状态，需要发起人 `metalmon` 提供更多信息或更新。这是一个高风险的特性，影响到项目在复杂部署场景下的核心能力，建议维护者主动联系该用户推动进展。

- **Issue #5775: per-skill security permissions** ([链接](https://github.com/zeroclaw-labs/zeroclaw/issues/5775))
    - **创建时间**: 2026-04-15
    - **标签**: `status:blocked`
    - **状态**: 同样处于阻塞状态。该问题与 #5982 共同构成了安全权限体系的两大支柱，这些功能的缺失正在成为影响社区采用的关键瓶颈。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*