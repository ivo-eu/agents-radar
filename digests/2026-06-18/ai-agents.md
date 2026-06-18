# OpenClaw 生态日报 2026-06-18

> Issues: 252 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-06-18 03:43 UTC

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

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，我已经根据您提供的 OpenClaw GitHub 数据，生成了 2026-06-18 的项目动态日报。

---

### OpenClaw 项目动态日报 (2026-06-18)

**分析师点评：** 今日项目处于“**高活跃度、伴随瓶颈**”的状态。社区提交了大量 Issue 和 PR，但 PR 合并率较低，表明核心维护者或 CI 流程成为当前的主要瓶颈。安全与稳定性是社区讨论的绝对焦点，多个高优问题亟待解决。

---

#### 1. 今日速览
-   **活跃度评估：** 🔥 高度活跃。过去 24 小时内，共有 **252 条 Issue 更新** 和 **500 条 PR 更新**，社区参与度极高。待合并 PR 积压达到 **409 条**，反映出社区贡献意愿强烈但代码审查吞吐量可能已到极限。
-   **核心焦点：** 社区讨论和贡献高度集中在 **安全性（Secrets 管理、权限控制）**、**消息稳定性（消息丢失、路由错误）** 和 **会话状态管理（上下文丢失、内存混乱）** 三大领域。
-   **潜在风险：** 无新版本发布，且关闭/合并的 PR 仅占 18.2%。若合并流程无法提速，可能会影响贡献者积极性。

#### 2. 版本发布
-   **今日无新版本发布。**

#### 3. 项目进展
虽然今日合并率不高，但仍有一些关键性的 PR 被合入，标志着项目在稳定性和架构上的推进：
-   **修复路由与信道问题：**
    - **[已关闭] `#94309`**: 解决了 Telegram Desktop 无法引用/回复 OpenClaw 机器人消息的问题（`Quote & Reply`），改善了用户体验。
    - **[已关闭] `#68936`**: 合并了一个大型 PR，引入了 **基于 AI 的 PR 审查自动修复流水线**，并增加了 Windows 后台守护进程。这是提升开发效率和平台兼容性的重要一步。
-   **关键配置与兼容性修复：**
    - **[`#94020`]**: 修复了浏览器文档中关于 `networkidle` 等待策略的表述矛盾，完善了文档一致性。
    - **[`#93853`]**: 修复了当 embedding 提供商使用自定义基础URL时，内存嵌入路由错误的问题，确保私有部署环境下功能正常。
    - **[`#94054`]**: 修复了 TLS 配置中允许空字符串导致的安全隐患，强化了网关连接的安全性。

#### 4. 社区热点
社区讨论热度高度集中在与**代理安全**和**消息传输可靠性**相关的问题上：

1.  **[#25592 - 工具调用间的文本泄露到消息频道]**: 获得 **32条评论**，是今日最热的 Issue。该问题描述了代理在处理工具调用时产生的内部文本（如错误处理、处理确认）意外泄露到用户可见的聊天频道中。这是一个严重的 **UX 和安全隐患**，社区有大量关于如何隔离内部输出与用户输出的讨论。
2.  **[#9443 - 请求预构建的 Android APK 发布]**: 获得 **25条评论**。时间跨度长、关注度高，表明社区对移动端（Android）体验有强烈需求，而当前编译流程对普通用户有门槛。
3.  **[#10659 - 请求“遮盖式”密钥系统]**: 获得 **13条评论**，并收获 **4个👍**。社区主流观点认为，代理不应能直接看到 API 密钥的明文，以防止泄漏和被注入攻击利用。这个需求与 [#25592] 共同构成了社区对“防泄漏”机制的核心诉求。

#### 5. Bug 与稳定性
今日报告了多个关键 Bug，部分已有修复 PR 链接，反映出社区正在积极自救：

-   **致命级 (Crash / Wedge)：**
    -   `#45224` (P1, Platinum Hermit): **Playwright 断言错误导致 Gateway 崩溃**。一个未被捕获的 Playwright 错误直接导致整个 Gateway 进程退出，需要重启。严重性高，暂无 fix PR。
    -   `#86215` (P1, Platinum Hermit): **Codex OAuth 刷新失败导致代理“卡死”**。认证问题无法被有效恢复，导致代理长时间不可用。暂无 fix PR。
-   **高严重级 (Message Loss / Session State / Security)：**
    -   `#41744` (P1, Diamond Lobster): **飞书 (Feishu) 图片消息丢失**。代理读取并回复图片后，图片在发送过程中丢失。已有 **linked PR**，表明修复在望。
    -   `#41165` (P1, Diamond Lobster): **Telegram 私信路由到主会话**。修复了上一个 Bug 后又出现回归，私信仍然会污染主工作流会话，非常影响多频道并发的稳定性。已有 **linked PR**。
    -   `#40255` (P1, Diamond Lobster): **用户自定义的心跳提示词被覆盖**。一个 Commit 引入了额外的文本，导致用户配置的提示词无效。存在回归问题。已有 **linked PR**。
    -   `#43747` (P2, Diamond Lobster): **内存管理“一团糟”**。不同用户的内存管理行为不一致（有的存储到 SQLite，有的到向量数据库等），引发用户困惑和功能失效。暂无 fix PR。
    -   `#41201` (P2, Diamond Lobster): **控制 UI 头像无法显示**。该回归问题导致用户无法通过 UI 识别和区分不同代理。暂无 fix PR。

#### 6. 功能请求与路线图信号
用户在新功能上表现出对**企业级管控、安全治理和可观测性**的强烈期待：

-   **企业级预算与成本管控：**
    -   **[#42475] (P2): 请求基于代理的成本预算强制功能** - 在网关层面实现每日/每月开销上限，防止预算超支。
    -   **[#13219] (P2): 请求按模型进行使用量日志记录** - 方便管理员进行成本追踪和模型混合优化。
-   **深度安全与策略引擎：**
    -   **[#13583] (P2): 请求“硬门闩”（Pre-response enforcement hooks）** - 要求代理在执行特定操作（如查询数据库）前必须调用预定义的工具，而不能仅依赖软提示词。这是一个向“策略即代码”演进的信号。
    -   **[#7707] (P2): 请求按来源进行内存信任标记** - 防止来自网页、第三方技能等不可信来源的“记忆投毒”攻击。
-   **架构演进与可观测性：**
    -   **[#42026] (P2): 提议分布式代理运行时** - 分离控制平面和代理计算平面。这是一个重要的架构演进提案，旨在提高系统的可扩展性和隔离性。
    -   **[#28300] (P2): 请求主题定制系统** - 控制 UI 的预设主题 + 自定义主题工作室，虽非核心功能但收到 **5个👍**，反映了社区对优化用户体验细节的渴望。

#### 7. 用户反馈摘要
-   **用户痛点（高频场景）：**
    -   **多代理/多信道混乱：** 用户在真实多代理群聊和跨信道（TUI、iMessage、Telegram）使用中，普遍遇到了路由不一致、消息可见性丢失、回复错位等问题。`#43747` (`AM-young-fun` 反馈团队3人使用体验不同)、`#40678` 和 `#41165` 都反映了这一点。
    -   **“软规则”不可靠：** 在金融、安全等高风险场景中，用户发现仅通过提示词来指导代理行为的“软规则”不可接受。`#13583` 的作者直言需要“机械性地阻止”未满足条件的回复。
    -   **文档与实践脱节：** `#41372` 的作者发长文列出了从自托管生产中收集的 25 项发现，涵盖了配置崩溃、CLI 文档与行为不符、Cron 作业等多个方面，表示文档更新未能跟上快速迭代的步伐。
-   **用户期望：**
    -   **更强的透明度和可控性：** 用户希望看到推理过程 (`#42276`)、能控制消息的发送方式（是否作为文档而非图片发送 `#41965`），并希望有全局“干跑”模式 (`#41418`) 来预览工具执行结果。
    -   **开箱即用体验：** `#9443` 请求预编译的 APK、`#13597` 请求 AWS 部署指南，都表明用户希望降低项目的使用和部署门槛。

#### 8. 待处理积压
-   **`#41372` (P2, Platinum Hermit):** **25项自托管生产使用发现**。这份“场外报告”包含了大量未解决的配置崩溃、CLI 文档误导、Discord 和 Cron 行为异常问题。虽然作者提供了工作流建议，但其中许多 Bug 和文档问题仍需维护者系统性地跟进和修复。**建议优先处理。**
-   **`#42273` (P2, Silver Shellfish):** **备份命令在大目录下卡死**。`openclaw backup create` 在数据量达到 4GB+ 时会导致静默失败。这是一个严重的运维问题，但至今未得到维护者确认或分配。
-   **`#45224` (P1, Platinum Hermit):** **Playwright 断言错误导致整个 Gateway 崩溃**。该 Bug 会导致服务完全中断，影响所有依赖 Playwright（浏览器自动化）功能的用户。当前无任何 assignee 或 linked PR，需尽快处理。

---

## 横向生态对比

好的，作为资深技术分析师，我已根据您提供的各项目2026-06-18动态日报，为您整合生成一份横向对比分析报告。

---

### AI智能体与个人AI助手开源生态横向分析报告 (2026-06-18)

#### 1. 生态全景

当前，AI智能体与个人AI助手开源生态正处于**高强度的“功能爆发”与“稳定性阵痛”并存**的快速演进期。一方面，以`computer use`、`多租户架构`、`跨Agent协作`为代表的前沿功能正被快速集成；另一方面，社区反馈高度集中在**消息路由混乱、安全漏洞频发、跨平台兼容性差**等核心稳定性问题上。`OpenClaw`作为生态核心，其PR合并瓶颈已构成整个生态的“限速器”，迫使许多下游项目（如`CoPaw`）启动自有的2.0架构升级。同时，**安全治理**（SSRF、文件泄露、OAuth劫持）正从最佳实践转变为急迫的补丁需求，成为衡量项目成熟度的新标尺。

#### 2. 各项目活跃度对比

| 项目 | 今日新Issues | 今日新PRs | 版本发布 | 健康度评估 | 分析师点评 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 252 (更新量) | 500 (更新量) | 无 | ⚠️ 高活跃，但合并效率低 | 社区贡献旺盛，核心维护成为瓶颈，安全与稳定性是绝对焦点。 |
| **NanoBot** | 1 | 多 (17 PR合并/关闭) | 无 | ✅ 健康，修复效率高 | 项目维护响应积极，Bug修复与完善功能并重，社区活跃度良好。 |
| **Hermes Agent** | 10 | 50 | 无 | ⚠️ 活跃，但核心修复缓慢 | 功能请求多，但核心团队合并动作少。OAuth安全与桌面端体验是热点。 |
| **PicoClaw** | 2 | 6 (合并/关闭) | 无 | ✅ 健康，响应迅速 | 项目维护者响应积极，快速修复了SSRF和高模型兼容性等关键问题。 |
| **NanoClaw** | - (多个文档相关) | 20 (含3合并) | **v2.1.17** | ✅ 健康，迭代迅速 | 版本发布后快速跟进修复，同时暴露了严重的路径遍历与文件读取漏洞。 |
| **NullClaw** | 0 | 2 (待合并) | 无 | ⚠️ 停滞，积压待解决 | 无Bug修复，核心调度器Bug与文档易用性问题长期积压，社区缺乏回应。 |
| **IronClaw** | 11 | 50 | 无 | ✅ 健康，开发密集 | 核心团队与社区共同推进，安全与无进展检测等功能是近期重点。 |
| **LobsterAI** | 1 (安全) | 10 (合并) | **2026.6.15** | ✅ 健康，技术债务清理 | 发布新版本并引入重要功能，同时聚焦对话体验和资源管理优化。 |
| **TinyClaw** | 3 (安全) | 0 | 无 | ❌ 危急，安全风险高 | 一天内被曝出3个严重未认证攻击漏洞，无任何修复响应，风险极高。 |
| **Moltis** | 2 | 1 (待合并) | 无 | ⚠️ 低活跃，稳步前进 | 活跃度较低，但功能请求清晰，集中在输出格式灵活性与数据导出。 |
| **CoPaw** | 45 (更新量) | 50 (更新量) | **v1.1.12** | ✅ 健康，生态迁移意图明显 | 发布正式版，启动2.0规划。跨Agent死循环和审批流问题是核心痛点。 |
| **ZeptoClaw** | 0 | 0 | 无 | ⚠️ 沉默 | 24小时内无任何活动，项目可能处于暂停或维护周期。 |
| **ZeroClaw** | 17 | 50 | 无 | ✅ 健康，平台化演进 | 核心功能（配置级联、WASM插件）推进顺利，Windows兼容性为关键短板。 |

*(注：由于数据源格式差异，OpenClaw和CoPaw的Issues/PRs为更新总量，其他项目为新开/创建量)*

#### 3. OpenClaw在生态中的定位

- **核心参照与生态基石**: `OpenClaw` 凭借其庞大的社区规模（单日252条Issue，500条PR）和详尽的功能集，仍然是该生态中无可争议的**参照性项目和功能先行者**。社区提出的许多需求，如成本预算管控、策略引擎等，都代表了整个行业对Agent平台化的深度思考。
- **优势**: 社区活跃度“断层领先”，功能覆盖最全面，是企业级和个人高级用户的首选实验场。几乎所有下游项目都在不同程度上参考或兼容其设计。
- **核心瓶颈与差异**:
    - **瓶颈**: 其**PR合并率极低（18.2%）**，核心维护或CI流程成为当前限制生态整体创新速度的最大障碍。大量社区贡献被积压，可能挫伤贡献者积极性。
    - **技术路线差异**: 相较于`CoPaw`、`IronClaw`等积极启动**2.0/架构升级**的项目，`OpenClaw`当前更侧重于修补既有问题（安全、消息丢失），而非宣布重大架构重构。这反映了不同项目在“稳定存量”与“定义增量”上的策略分化。

#### 4. 共同关注的技术方向

1.  **开放架构与插件生态系统**: 多个项目（`ZeroClaw`、`NanoClaw`、`PicoClaw`）都在推动**技能/插件/WASM**平台建设，旨在让Agent能力通过第三方或社区贡献进行热插拔和自我进化。
2.  **安全治理与合规**: 这是今日最普遍的主题：
    - **SSRF与文件读取**: `PicoClaw` (修复)、`NanoClaw` (新发现严重漏洞)、`LobsterAI` (新发现严重漏洞)、`TinyClaw` (新发现多个严重漏洞)。
    - **权限与策略**: `OpenClaw` (策略引擎需求)、`IronClaw` (修复Slack OAuth安全)、`Hermes Agent` (记忆工具绕过漏洞)。
    - **密钥与计费**: `OpenClaw` (遮盖式密钥)、`Hermes Agent` (OAuth计费错误)。
    - **结论**: 安全已从“加分项”变为“必考题”，未通过安全审查的项目可能迅速流失用户信任。
3.  **成本优化与可观测性**: `OpenClaw` (预算管控、按模型日志)、`NanoBot` (按模型配置上下文)、`ZeroClaw` (原生上下文压缩RFC) 和 `IronClaw` (文件系统浏览器) 均反映了用户对**透明、可控成本**和**内部状态可视化**的强烈诉求。
4.  **多平台与多通道兼容**:
    - **桌面端**: `Hermes Agent` (纯客户端部署)、`ZeroClaw` (Windows更新修复)、`CoPaw` (WebUI性能反馈)。
    - **移动端**: `OpenClaw` (Android APK请求)、`ZeroClaw` (Android Termux反馈)。
    - **消息渠道**: `OpenClaw`、`PicoClaw`、`CoPaw` 都在处理特定渠道（Telegram, Discord, Matrix, QQ）的路由、审批和消息一致性问题。

#### 5. 差异化定位分析

| 维度 | OpenClaw | NanoBot | Hermes Agent | PicoClaw | CoPaw | ZeroClaw |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **功能侧重** | 全能型、企业级平台 | 易用性、快速修复 | 桌面体验、创意工作流 | 轻量、嵌入式、移动优先 | 多Agent协作、A2A协议 | WASM插件、平台安全、云原生 |
| **目标用户** | 高级开发者、企业架构师 | 普通开发者、个人用户 | 个人高级用户、创意工作者 | 嵌入式和IoT开发者 | 企业团队、协作场景 | 平台开发者、安全敏感用户 |
| **技术架构关键差异** | 单实例、插件驱动 | 网关-代理模式、MCP支持 | 桌面客户端捆绑运行时 | OneBot协议、资源占用低 | 多Agent、A2A、2.0架构规划 | WASM沙箱、配置级联、结构化安全策略 |

#### 6. 社区热度与成熟度

- **第一梯队（高度活跃，快速迭代）**: **OpenClaw** (生态中心), **NanoClaw** (版本发布密集), **IronClaw** (开发速度快), **ZeroClaw** (核心功能推进快)。这些项目社区贡献度高，团队响应及时，是生态创新的主要引擎。
- **第二梯队（稳定活跃，质量巩固）**: **NanoBot**, **PicoClaw**, **CoPaw**, **LobsterAI**。这些项目注重Bug修复和体验优化，项目健康度良好，是生态的中坚力量。
- **第三梯队（低活跃，风险累积）**: **Hermes Agent**, **NullClaw**, **Moltis**。这些项目或有核心Bug未解，或社区互动不足，活跃度相对偏低，可能进入“跟随创新”阶段。
- **高危项目**: **TinyClaw**。一天内发现多个严重漏洞但无响应，项目可能存在严重的“维护者缺失”风险，用户应立即暂停使用并关注后续。

#### 7. 值得关注的趋势信号

1.  **安全治理成为准入标准**: `TinyClaw`的事件是一个强烈警告。对于个人AI助手这类可访问API密钥和本地文件的软件，**基础的安全审计（如输入校验、认证、权限隔离）** 已成为用户选择项目的首要考量。开发者应将安全设计作为架构的默认考虑，而非事后补丁。
2.  **从“软规则”到“策略即代码”**: `OpenClaw`社区对“硬门闩”的呼唤，以及`ZeroClaw`对结构化安全策略的探索，标志着行业正从依赖LLM提示词（软规则）来约束行为，转向**代码级的、不可绕过的策略引擎**。这将是企业级Agent落地的关键一步。
3.  **“瘦客户端”与“分布式Agent”架构崛起**: `Hermes Agent`的纯客户端部署需求、`OpenClaw`的分布式运行时提案以及`CoPaw`的跨Agent协作讨论，共同指向一个趋势：**Agent的计算与控制平面正在分离**。未来，个人拥有轻量级交互终端，而计算和智能解算在云端或边缘自治集群中完成。
4.  **Agent“记忆”的信任与治理**: 记忆投毒、上下文泄露（如`OpenClaw`的工具调用文本泄露）等问题正在催生新的需求：**来源可追溯”和“能力可隔离”的记忆系统**。开发者需要思考如何让Agent在利用历史信息时，能区分可信与不可信的来源。

**总结建议**：对于技术决策者，在选择平台时，应优先评估项目的**安全治理能力**和**核心维护团队的响应速度**。对于开发者，**跨Agent安全协作**、**策略即代码** 以及 **可观测性** 是未来最值得投入的技术方向。`OpenClaw` 生态依然是首选，但跟踪其合并瓶颈的解决方案，以及关注`CoPaw`、`ZeroClaw`等项目的架构演进，将是获取先发优势的关键。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，我已根据您提供的 NanoBot 项目 GitHub 数据，生成以下 2026-06-18 的项目动态日报。

---

# NanoBot 项目动态日报 | 2026-06-18

## 今日速览

今日 NanoBot 项目社区活动保持高度活跃。Issues 和 PRs 的更新数量均处于高位，反映出用户持续的反馈热情和开发团队的积极响应。本日有多达 17 个 PR 被合并或关闭，显示了项目在 Bug 修复、功能增强和稳定性改进方面的强劲推进力。无新版本发布，但多个关键性修复已进入待合并状态，预示着下一个修补版本可能即将到来。

## 项目进展

今日项目有 17 个 PR 被合并或关闭，主要集中在 **Bug 修复** 和 **功能完善** 两个方向，体现了项目维护者对稳定性和用户体验的高度重视。

- **核心稳定性提升**：
  - `fix(mcp): close tracked generators in _close_server to prevent GC crash (#4302)` [PR #4303](https://github.com/HKUDS/nanobot/pull/4303) 已合并，修复了 MCP 服务器重连时因生成器未正确关闭导致的运行时崩溃问题。
  - `fix(anthropic): sanitize tool_use/tool_result IDs to API pattern` [PR #4356](https://github.com/HKUDS/nanobot/pull/4356) 已修复，解决了从其他模型或会话恢复时，工具 ID 包含非法字符导致 Anthropic API 请求失败的问题。
  - `fix(providers): disable proxy for local endpoints, respect env proxy for cloud` [PR #4367](https://github.com/HKUDS/nanobot/pull/4367) 已合并，修复了本地模型服务（如 Ollama）因全局代理设置而连接失败的问题，提升了开发环境的兼容性。

- **功能与体验优化**：
  - `fix: allow git commands in workspace subdirectories` [PR #4380](https://github.com/HKUDS/nanobot/pull/4380) 已合并，修复了工作区安全策略限制导致 Git 命令无法在子目录执行的 Bug。
  - `fix: recover failed Feishu streaming updates` [PR #4381](https://github.com/HKUDS/nanobot/pull/4381) 已合并，通过重试机制增强了飞书流式更新功能的健壮性。
  - `feat(web): add Keenable search provider` [PR #4350](https://github.com/HKUDS/nanobot/pull/4350) 已合并，新增 Keenable 作为内置网络搜索提供商，丰富了工具生态。

- **架构与代码质量**：
  - `feat(providers): better Mistral support` [PR #4351](https://github.com/HKUDS/nanobot/pull/4351) 已合并，针对 Mistral API 的特定限制进行了适配，改进了兼容性。
  - `feat(bridge): send read receipts (blue ticks) for incoming WhatsApp messages` [PR #4354](https://github.com/HKUDS/nanobot/pull/4354) 已合并，为 WhatsApp 桥接增加了“已读回执”功能。

## 社区热点

- **工作区文件读写不对称问题** `#4374`：该 Issue 指出 WebUI 项目工作区从项目根目录读取 `SOUL.md`，但在写入时却写到了默认工作区，存在“读写不对称”问题。这触及了工作区隔离和用户自定义 Agent 行为的核心逻辑，引发了社区开发者对“什么是项目根目录”的讨论。
  - Issue: [HKUDS/nanobot Issue #4374](https://github.com/HKUDS/nanobot/issues/4374)

- **新用户向导增强需求** `#4376`：用户提议改进 `nanobot onboard --wizard` 的易用性，以降低非技术用户的上手门槛。此建议获得了社区的正向反馈，反映了项目从小众开发者向更广泛用户群拓展的诉求。
  - Issue: [HKUDS/nanobot Issue #4376](https://github.com/HKUDS/nanobot/issues/4376)

## Bug 与稳定性

- **[严重] Git 命令被工作区安全策略误拦截** `#4375`：在安全策略允许的工作区子目录内执行 Git 操作失败。该问题已由 PR `#4380` 和 PR `#4393` 修复和测试覆盖，风险已解除。
  - Issue: [HKUDS/nanobot Issue #4375](https://github.com/HKUDS/nanobot/issues/4375)

- **[中] iOS Safari 输入框触发页面放大** `#4388`：在 iOS Safari 浏览器上点击 WebUI 输入框时页面自动放大，影响移动端使用体验。目前尚未有对应修复 PR，是移动端体验的一个已知痛点。
  - Issue: [HKUDS/nanobot Issue #4388](https://github.com/HKUDS/nanobot/issues/4388)

- **[低] `session_key` 未定义导致启动崩溃** `#4322`：此问题由合并操作引起，已被标记为已关闭，表明修复已合并或已确认非普遍性问题。
  - Issue: [HKUDS/nanobot Issue #4322](https://github.com/HKUDS/nanobot/issues/4322)

## 功能请求与路线图信号

- **多租户网关** `#936`：一个持续数月的长期功能请求，要求支持单一网关实例管理多个 Agent，以降低资源消耗。虽然 PR 尚未合并，但该功能属于架构级增强，契合企业级部署场景，是未来路线图的重要候选。
  - Issue: [HKUDS/nanobot Issue #936](https://github.com/HKUDS/nanobot/issues/936)
- **按模型配置上下文窗口** `#4389`：用户要求在全局层面外，允许为 fallback 模型单独配置 `contextWindowTokens`，以避免因 fallback 模型上下文窗口过小导致内容截断错误。这是一个精细化的配置增强，有望提升模型切换的稳定性。
  - Issue: [HKUDS/nanobot Issue #4389](https://github.com/HKUDS/nanobot/issues/4389)
- **简化多实例管理** `#4390`：用户希望能在单个机器上通过文件夹组织多个 Agent 实例，并能通过 UI 进行管理。这与 `#936` 的目标一致，都是为了让实例管理更简单。
  - Issue: [HKUDS/nanobot Issue #4390](https://github.com/HKUDS/nanobot/issues/4390)

## 用户反馈摘要

- **痛点**：用户在 `#4374` 的评论中反映，项目工作区的文件读写行为与预期不一致，`SOUL.md` 等指导性文件“从哪读就应该写回哪去”，而不是默认写到一个全局位置，这导致了数据和配置的混乱。
- **痛点**：在 `#4376` 的讨论中，用户明确指出当前的初始化向导对新用户极其不友好，技术门槛高，期望能有一个更简单、更“傻瓜式”的配置流程。

## 待处理积压

- **`#936` 多租户网关需求**：自 2026年2月提出至今超过4个月，目前仍为开放状态，无关联 PR。作为一项能够显著降低部署成本和复杂度的架构升级，建议维护者基于社区反馈评估其优先级，并考虑纳入版本规划。
  - Issue: [HKUDS/nanobot Issue #936](https://github.com/HKUDS/nanobot/issues/936)
- **`#3437` 按需心跳触发调试功能**：自2026年4月底提出，允许开发者在不消耗令牌的情况下，手动触发调试心跳。目前仍为 RFC 状态，但已有实现细节讨论，建议明确是否进入开发管线。
  - Issue: [HKUDS/nanobot Issue #3437](https://github.com/HKUDS/nanobot/issues/3437)

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

好的，这是为您生成的 Hermes Agent 项目动态日报。

---

# Hermes Agent 项目日报 | 2026-06-18

## 今日速览

今日项目社区活跃度极高，共产生 10 个新 Issue 和 50 个 PR。然而，核心开发团队的合并/关闭动作较少，仅处理了 2 个 PR，大部分工作仍处于社区贡献和问题发现阶段。Issue 方面，功能请求和安全/计费相关的 Bug 讨论最为热烈，显示出用户对桌面体验优化、OAuth 计费兼容性以及安全性有迫切需求。项目整体处于功能密集提交但合并速度暂缓的活跃状态。

## 项目进展

今日 PR 合并/关闭数量仅为 2 个，进度相对放缓。主要推进的贡献如下：

- **技能与工作流修复 (PR #48145)**：修复了捆绑的 ComfyUI 工作流无法提交的问题，原因是工作流 JSON 文件中的 `_comment` 键未被清理。此 PR 已关闭，直接提升了相关创意技能的开箱即用性。
  - [NousResearch/hermes-agent PR #48145](https://github.com/NousResearch/hermes-agent/pull/48145)

## 社区热点

- **桌面客户端纯客户端安装 (Issue #38602, 6 评论, 18 👍)**：这是社区讨论最热门的 Issue，用户强烈希望 Hermes Desktop 能够作为一个纯客户端连接到远程服务器，而不是强制在本地启动完整的 Agent 运行时。该诉求获得 18 个点赞，表明许多用户倾向于更轻量、分布式的部署方式。
  - [NousResearch/hermes-agent Issue #38602](https://github.com/NousResearch/hermes-agent/issues/38602)

- **OAuth 计费问题 (Issue #48176, 1 评论)**：该 Issue 报告了使用 Claude Pro/Max/Team OAuth 凭证时，请求被返回 HTTP 400 错误，提示“第三方应用使用了额外用量而非计划内额度”。此问题直接影响了所有使用 Anthropic 平台的付费用户，是当日最紧迫的 Bug 之一，并且已有对应的 fix PR (#48177)。社区对此高度关注，核心问题在于缺少必要的计费归属头部信息。
  - [NousResearch/hermes-agent Issue #48176](https://github.com/NousResearch/hermes-agent/issues/48176)

- **Discord 多频道提示词配置 (Issue #48175, 2 评论)**：用户寻求如何为多个 Discord 频道配置完全不同的 AI 人设和系统提示词。这是一个典型的高级用户配置需求，反映出 Hermes 在多平台、多场景下的配置灵活性仍有提升空间。
  - [NousResearch/hermes-agent Issue #48175](https://github.com/NousResearch/hermes-agent/issues/48175)

## Bug 与稳定性

- **[严重] OAuth 计费错误 (Issue #48176, P1)**：Anthropic OAuth 请求因缺少 `x-anthropic-billing-header` 被拒绝，导致 Pro/Max/Team 用户无法使用。**已有对应 fix PR (#48177)**。
  - [NousResearch/hermes-agent Issue #48176](https://github.com/NousResearch/hermes-agent/issues/48176)

- **[严重] 记忆工具绕过安全漏洞 (Issue #48181, P2, 重复)**：报告指出，即使配置了 `disabled` 基于 `memory` 的 toolset，攻击者仍可通过后期注入其他工具的 memory 提供者来绕过此限制。这是一个重要的安全漏洞，可能破坏用户配置的隔离策略。被标注为重复，可能已有内部或社区的跟踪。
  - [NousResearch/hermes-agent Issue #48181](https://github.com/NousResearch/hermes-agent/issues/48181)

- **[中等] 桌面版 Profile 切换问题 (Issue #48183, P2)**：在桌面应用中切换 Profile 后，会话列表显示异常，会话数据访问错乱。这直接破坏了多 Profile 场景下的核心用户体验。
  - [NousResearch/hermes-agent Issue #48183](https://github.com/NousResearch/hermes-agent/issues/48183)

- **[中等] Docker 环境更新问题 (PR #48188, P2)**：修复了桌面应用和 Docker 容器共享 `~/.hermes` 时，`hermes update` 命令无法正确应用更新的问题。该 PR 已提交，解决了混合部署场景下的核心痛点。
  - [NousResearch/hermes-agent PR #48188](https://github.com/NousResearch/hermes-agent/pull/48188)

- **[轻微] React 错误 Dashboard 聊天页 (Issue #41808, P3)**：“Maximum update depth exceeded” 渲染错误发生在通过 Tailscale IP 等外部连接访问 Dashboard 时，但本地访问正常。
  - [NousResearch/hermes-agent Issue #41808](https://github.com/NousResearch/hermes-agent/issues/41808)

- **[轻微] 剪贴板粘贴失败 - WSL (PR #48186, P3)**：在 WSL 环境下，Hermes Desktop 无法粘贴图片，原因是 Electron 和 WSLg 的剪贴板路径冲突。已有修复 PR。
  - [NousResearch/hermes-agent PR #48186](https://github.com/NousResearch/hermes-agent/pull/48186)

- **[轻微] macOS Chrome 检测失败 (PR #48185, P2)**：运行时浏览器可用性检查无法识别 macOS 应用包形式安装的 Chrome，导致相关浏览器工具被隐藏。已有修复 PR。
  - [NousResearch/hermes-agent PR #48185](https://github.com/NousResearch/hermes-agent/pull/48185)

## 功能请求与路线图信号

- **增强桌面客户端与 CLI 体验**：多个 Issue 指向改进桌面和命令行交互体验。例如，为桌面端添加命令中心（#48189）启动/停止消息网关的功能，以及为 CLI 引入会话与工作区绑定的概念（#48190），允许记录并恢复 `cwd` 和 `git repo`。这暗示着 Hermes 正朝着更强大的 “AI 工作环境管理” 方向演进。

- **系统托盘与更好的交互**：社区贡献了系统托盘支持 (PR #48163)，并提出了包括技能高亮、更频繁的联系人在列表展示等一系列交互改进建议 (Issue #48182)。这些信号表明用户对优雅的、低侵入性的后台运行和流畅的日常使用体验有较高期望。

- **扩展 “托管系统” 概念**：Issue #48179 提出将 “托管系统” 功能（目前支持 NixOS 和 Homebrew）扩展到 Fedora 等其他 Linux 发行版。这表明项目在寻求更广泛的平台兼容性，尤其是系统包管理器集成方面。

## 用户反馈摘要

- **痛点：桌面端安装流程不灵活**：Issue #38602 揭示了用户对“强制捆绑启动 Agent 运行时”的不满，部分用户希望将桌面端作为纯粹的瘦客户端，这与当前的架构设计存在矛盾。
- **痛点：OAuth 计费问题导致服务中断**：Issue #48176 引发了使用 Anthropic 付费服务的用户焦虑，无法使用核心功能是用户最直接的挫折感来源。社区强烈期望快速修复。
- **诉求：更精细的配置能力**：Issue #48175 的讨论显示了高级用户希望完全掌控多频道的 AI 人设和提示词，当前配置系统对于这种场景的简洁支持不足。
- **安全担忧**：Issue #48181 揭露的记忆工具绕过漏洞引发了社区的潜在安全担忧，用户期望更严格的沙箱机制。

## 待处理积压

以下是一些创建较久或涉及核心功能，但仍未合并或得到充分响应的 PR/Issue，值得项目维护者关注：

- **`cognee_query` 工具 (PR #19331)**：创建于 2026-05-03，旨在为基于 `cognee` 的知识库添加只读查询工具。长期未合并可能因设计决策或测试问题而延迟。
  - [NousResearch/hermes-agent PR #19331](https://github.com/NousResearch/hermes-agent/pull/19331)

- **`agent_loop_stopped` 插件钩子 (PR #27208)**：创建于 2026-05-17，为用户中断 Agent 时提供了关键的资源释放机会。此 PR 对于构建健壮、资源友好的插件生态至关重要，但仍在开放中。
  - [NousResearch/hermes-agent PR #27208](https://github.com/NousResearch/hermes-agent/pull/27208)

- **Dashboard 会话持久化 (PR #41768)**：创建于 2026-06-08，旨在修复侧边栏 Chat 链接无法保持恢复的会话 ID 问题。该 Bug 影响用户日常导航体验，需尽快解决。
  - [NousResearch/hermes-agent PR #41768](https://github.com/NousResearch/hermes-agent/pull/41768)

- **安装器 Node.js 检测问题 (PR #42698)**：创建于 2026-06-09，修复安装程序在 PATH 未加载时无法识别 NVM 等环境管理工具安装的 Node.js。对于新用户安装体验至关重要。
  - [NousResearch/hermes-agent PR #42698](https://github.com/NousResearch/hermes-agent/pull/42698)

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，根据 PicoClaw 项目 2026-06-18 的 GitHub 数据，现为您呈上项目动态日报。

---

### PicoClaw 项目动态日报 | 2026-06-18

---

#### 1. 今日速览

今日项目活跃度**较高**。虽然无新版本发布，但社区在问题解决和功能修复上表现积极。过去24小时内，共有6个 Pull Request (PR) 被合并或关闭，其中包含对**关键安全漏洞 (SSRF)** 的修复，以及对 **Gemini 3.5 Flash 模型兼容性** 问题的补丁。项目目前有4个待合并PR和2个新开的活跃Issue，显示社区贡献持续且维护者响应迅速。

#### 2. 版本发布

无

#### 3. 项目进展

今日项目在**安全加固、模型兼容性和用户体验**方面取得了显著进展。以下为合并/关闭的重要PR：

- **修复关键安全漏洞 (SSRF):** `#3140 [CLOSED]` 修复了OneBot频道在处理外部媒体URL时，可能被攻击者利用来请求内网地址的问题，增强了主机侧的安全性。
- **修复Gemini 3.5 Flash兼容性:** `#3136 [CLOSED]` 针对新模型要求，在API请求中补充了蛇形命名法 (`snake_case`) 的`thought_signature`字段，解决了工具调用失败 (`400 Bad Request`) 的问题。
- **修复Web搜索工具:** `#3139 [CLOSED]` 修复了因搜狗搜索页面HTML结构调整导致的解析失败问题，保证了Web搜索工具的可用性。
- **优化历史会话显示:** `#2990 [CLOSED]` 修复了Web UI中会话历史只能显示最后一条用户消息的问题，改善了用户浏览对话记录时的体验。

#### 4. 社区热点

今日社区讨论的焦点集中在**安全与性能**上，两个高优先级的Issues引发了关注：

- **[Feature] 使用 vodozemac 替换 libolm:** `#3088` [OPEN] - 该Issue获得2个👍，提议将项目中不再维护且存在安全隐患的 `libolm` 库替换为官方推荐的 `vodozemac`。这不仅是技术升级，也反映了社区对项目长期安全性和可维护性的高度关注。
- **[Security] OneBot SSRF 漏洞:** `#3070` [CLOSED] - 虽然该安全漏洞已在 `#3140` 中被修复，但在修复前获得了社区成员的积极讨论和报告，体现了社区在安全审计上的主动性和责任感。

#### 5. Bug 与稳定性

今日无新的严重级Bug报告，但有几个重要的修复值得关注：

- **[严重-已修复] OneBot SSRF 漏洞:** `#3070` 报告的允许攻击者诱导主机构造内部请求的漏洞，已在 `#3140` PR中紧急修复。
- **[高-已修复] Gemini 3.5 Flash 模型工具调用失败:** `#3111` 报告的使用新模型时API返回400错误的问题，已在 `#3136` PR中修复。
- **[中-待审查] 子代理消息重复:** `#3142` [OPEN] 报告了当子代理完成任务时，会导致重复消息投递的问题。目前已有修复PR等待合并。
- **[低-待审查] Brave搜索静默失败:** `#3141` [OPEN] 指出了Brave搜索API在返回空结果时缺乏诊断日志，导致问题难以定位。提交的PR旨在为该场景增加日志，属于可观测性改进。

#### 6. 功能请求与路线图信号

- **核心依赖升级（高优先级信号）:** `#3088` [OPEN] 提出的替换 `libolm` 为 `vodozemac` 的请求，因其高安全性和维护性影响，极有可能被列入下一版本的开发路线图。该Issue已被标记为 `priority: high`。
- **新网关支持（中优先级信号）:** `#3093` [OPEN] 用户请求增加SimpleX或Tox网关支持，拓宽了PicoClaw的通信渠道。与此同时，已有 `#3063` [OPEN] PR在实现DeltaChat网关，表明项目正在积极扩展网关生态。这可能成为下一个版本的功能亮点。
- **新增云提供商（已实现）:** 通过 `#2917` [CLOSED] 的合并，项目已成功集成NEAR AI Cloud作为新的LLM提供商，这反映了项目向去中心化AI基础设施扩展的一步。

#### 7. 用户反馈摘要

- **用户对废弃依赖感到担忧:** 在 `#3088` 中，用户明确指出 `libolm` 已不再维护且不安全，建议立即替换。这反映了用户对项目技术债务和安全风险的清晰认知与关切。
- **对特定网关的强烈需求:** `#3093` 和 `#3063` 中的用户表达了超越现有渠道的通信需求，期望能集成更多去中心化或隐私保护的网关，如SimpleX和DeltaChat。这表明用户群体不仅限于标准的聊天平台。
- **对模型兼容性敏感:** `#3111` 中用户积极测试新模型 `gemini-3.5-flash` 并快速报告了兼容性问题，显示了用户作为早期采纳者的热情以及项目对新模型支持的重要性。

#### 8. 待处理积压

- **`#3092 [stale, OPEN]`**: 修复技能安装时类型断言检查的PR已打开一周，目前状态为 `stale`。该问题会导致非预期的静默失败，影响用户体验，建议维护者优先处理。
    - 链接: `sipeed/picoclaw PR #3092`
- **`#3063 [OPEN]`**: 增加DeltaChat网关功能的PR已打开10天，但未获得合并或维护者反馈。该功能与社区反馈的热点方向一致，建议给予明确反馈或推进合并，以避免社区贡献者流失。
    - 链接: `sipeed/picoclaw PR #3063`

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，根据您提供的 NanoClaw (github.com/qwibitai/nanoclaw) 数据，我为您生成了 2026-06-18 的项目动态日报。

---

## NanoClaw 项目动态日报 | 2026-06-18

### 1. 今日速览

过去 24 小时，NanoClaw 项目展现出极高的活跃度。PR 数量激增（20 条），其中包含多个紧急 Bug 修复和安全性加固，反映了项目在 v2.1.17 发布后的快速响应和迭代能力。社区贡献者积极参与，不仅修复了关键性的“单会话故障导致全局瘫痪”问题，还提交了包括韩语 README、新的 CLI 仪表盘技能在内的增强功能。项目整体处于高强度的功能迭代与稳定维护并行阶段。

### 2. 版本发布

- **v2.1.17**
  - **更新内容**：这是一个累积性版本，包含了从 v2.1.1 到 v2.1.17 的所有 `package.json` 更新，是 v2.1.0 发布以来功能与修复的集合。
  - **破坏性变更**：
    - **`@onecli-sh/sdk` 升级**：SDK 版本从 0.5.0 直接跳升至 2.2.1。此项变更要求 OneCLI 服务器必须提供 `/v1` API，旧版服务器将无法处理 SDK 发出的任何请求（返回 404 错误）。项目已锁定官方网关和 CLI 版本。
    - **启动标记要求**：系统启动时增加了升级标记检查，如果 `data/upgrade-state.json` 文件未记录当前版本是通过正式渠道升级而来，主机将拒绝启动。
  - **迁移注意事项**：
    1.  **强制升级 OneCLI 服务器**：所有运行 NanoClaw 的实例必须确保其 OneCLI 服务器版本支持 `/v1` API。否则，安装 v2.1.17 后所有 SDK 功能将立刻失效。
    2.  **启动前检查**：请检查 `data/upgrade-state.json` 文件是否存在且状态正确。对于使用不可变镜像部署的托管集群，需设置环境变量 `NANOCLAW_DISABLE_UPGRADE_TRIPWIRE=1` 来绕过此检查。
    - 链接：[v2.1.17 Release](https://github.com/qwibitai/nanoclaw/releases/tag/v2.1.17)

### 3. 项目进展

今日虽有多达 17 个待合并的 PR，但 3 个已合并/关闭的 PR 均对项目稳定性有重大推进：

- **[已合并] 修复单会话故障导致全局投递瘫痪**：PR #2797 成功合入，解决了长期存在的 Issue #2796。该修复将每个会话的错误处理隔离开来，防止一个会话 `outbound.db` 的读取失败影响所有其他代理的消息投递，极大地增强了系统健壮性。
    - 链接：[PR #2797](https://github.com/qwibitai/nanoclaw/pull/2797)
- **[已合并] 修复托管集群 LLM 认证失效**：PR #2794 修复了 v2.1.17 引入的一个回归问题。该问题导致基于不可变镜像的托管集群在启动后无法完成 LLM API 的认证，所有智能体对话均因“401 无凭证”错误而失败。
    - 链接：[PR #2794](https://github.com/qwibitai/nanoclaw/pull/2794)
- **[已合并] 新增启动标志的环境变量 Opt-out**：PR #2780 为 v2.1.17 新增的启动检查提供了环境变量 `NANOCLAW_DISABLE_UPGRADE_TRIPWIRE` 的退出机制，使得托管集群等特殊部署场景可以平滑升级，体现了对用户反馈的快速响应。
    - 链接：[PR #2780](https://github.com/qwibitai/nanoclaw/pull/2780)

### 4. 社区热点

今日社区最为活跃的讨论集中在**文档完善**与**新手引导**上。贡献者 `specterslient95-lgtm` 提交了一系列 Issues 和对应的修复 PR，直指项目文档中的细节问题，引发了核心维护者的快速响应。

- **核心诉求**：降低新用户的上手门槛，完善技能（Skill）文档的清晰度和完整性。
- **典型示例**：
    - Issue #2787 指出“初始化 OneCLI”技能的文档中，OneCLI 的服务端口（10254）仅在最后的故障排除部分被提及，而在前文从未声明，导致用户困惑。对应的 PR #2788 已提交修复。
        - 链接：[Issue #2787](https://github.com/qwibitai/nanoclaw/issues/2787), [PR #2788](https://github.com/qwibitai/nanoclaw/pull/2788)
    - Issue #2791 指出“添加 iMessage 通道”技能的步骤 2 中，`git show` 重定向到文件的操作会因为 `src/channels/` 目录不存在而失败。对应的 PR #2792 已提交修复。
        - 链接：[Issue #2791](https://github.com/qwibitai/nanoclaw/issues/2791), [PR #2792](https://github.com/qwibitai/nanoclaw/pull/2792)

### 5. Bug 与稳定性

今日报告的 Bug 覆盖了从文档细节到严重安全漏洞的多个层面。

- **严重 Bug**
    - **CLI 命令“创建消息组”完全失效** (PR #2804)：`ncl messaging-groups create` 命令因数据库字段约束问题持续抛出 `NOT NULL constraint failed` 错误。已有 fix PR。
        - 链接：[PR #2804](https://github.com/qwibitai/nanoclaw/pull/2804)
- **安全漏洞**
    - **CLI 分组创建存在路径遍历漏洞** (PR #2800)：`ncl groups create` 命令未对输入的 `--folder` 参数进行有效校验，允许 `../../etc` 这样的路径，可能造成配置文件覆盖。已有 fix PR。
        - 链接：[PR #2800](https://github.com/qwibitai/nanoclaw/pull/2800)
    - **`send_file` 功能存在任意文件读取漏洞** (PR #2799)：攻击者可通过此功能读取容器内任意文件，包括凭证和机密数据。已有 fix PR。
        - 链接：[PR #2799](https://github.com/qwibitai/nanoclaw/pull/2799)
- **文档/体验 Bug**
    - **`setup` 技能文档过于简略** (Issue #2789, PR #2790)：10 行的文档仅有进度条文案，无具体恢复步骤。
    - **`migrate-nanoclaw` 技能文档标题不当** (Issue #2785, PR #2786)：一级标题为泛泛的“Context”，无描述性。
    - **`safeParseContent` 函数解析非对象 JSON 失败** (PR #2801)：处理像 `"5"` 或 `"true"` 这样的原始值时，导致调用者获取不到预期的 `.text` 属性。已有 fix PR。
        - 链接：[PR #2801](https://github.com/qwibitai/nanoclaw/pull/2801)

### 6. 功能请求与路线图信号

今日虽然没有直接的“功能请求” Issue，但多个开放的 PR 清晰地展示了社区和项目未来可能的发展方向：

- **代理间消息治理 (PR #2793)**：`moshe-nanoco` 提交的 PR 引入了在已连接的代理之间，按消息级别添加“需要批准”的策略。这标志着项目开始向更细粒度的多代理协作与治理迈进，可能成为 v2.2 版本的核心功能之一。
    - 链接：[PR #2793](https://github.com/qwibitai/nanoclaw/pull/2793)
- **国际化与社区建设 (PR #2806)**：`arkjun` 提交的韩语 README 翻译，表明社区正在自发性地进行国际化推广，有助于项目吸引更广泛的开发者群体。
    - 链接：[PR #2806](https://github.com/qwibitai/nanoclaw/pull/2806)
- **提升用户可观察性 (PR #2795)**：`leetwito` 提交的 “/add-clidash” 技能，旨在提供一个只读的 CLI 仪表盘，帮助用户更好地理解和管理其智能体实例的运行状态。
    - 链接：[PR #2795](https://github.com/qwibitai/nanoclaw/pull/2795)

### 7. 用户反馈摘要

从今日的 Issues 和 PR 讨论中，可以提炼出以下用户痛点：

- **文档门控 (Documentation as Gatekeeper)**：多数新用户问题并非代码 Bug，而是文档中的不一致或缺失。例如“端口 10254 未提前声明”、“技能目录需要手动创建”等，这些问题直接影响了上手体验和信心。
- **“黑盒”故障体验**：当 `setup` 或 `migrate` 这样的关键流程失败时，用户希望得到具体、可执行的恢复步骤，而不是单纯的非正常结束或指向一个笼统的脚本。
- **功能可发现性**：部分功能或参数（如 `--folder` 校验）仅在出错或故障排除时被提及，缺乏前置引导，使用户在出错后才了解其存在与规则。

### 8. 待处理积压

以下 PR 已开放较长时间，解决的关键问题可能影响项目稳定性或生态建设，建议维护者予以关注：

- **PR #2750：修复容器异常死亡后 `outbound.db` 的日志恢复问题**，开放时间已超过 6 天。该 PR 解决了主机端只读数据库句柄的两个已知失效模式（issue #2516, #2640），对系统在容器故障后的自我修复能力至关重要。
    - 链接：[PR #2750](https://github.com/qwibitai/nanoclaw/pull/2750)
- **PR #2717：为 README 添加 Atlas Cloud 作为 OpenAI 兼容的后端选项**，开放时间已超过 9 天。该 PR 是一个纯文档贡献，旨在增加项目生态系统的选择，降低用户接入不同 LLM 提供商的成本。
    - 链接：[PR #2717](https://github.com/qwibitai/nanoclaw/pull/2717)
- **PR #2802：为 `ncl` socket 客户端添加请求超时和响应大小限制**，这是防止因异常主机导致客户端无限等待或内存溢出的关键稳定性更新。
    - 链接：[PR #2802](https://github.com/qwibitai/nanoclaw/pull/2802)

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

好的，我是 NullClaw 项目的 AI 智能体与个人 AI 助手领域开源项目分析师。以下是为您生成的 2026年6月18日 项目动态日报。

---

## NullClaw 项目动态日报 (2026-06-18)

### 1. 今日速览

今日项目整体活跃度 **中等**，处于“开发进行中”状态。虽然过去24小时内没有新的版本发布，但有两个关键性的 Pull Request (PR) 正在等待合并，分别针对长期存在的 CLI 控制字符问题和内存系统可配置性。Issues 数量保持不变（3个），且均为历史遗留问题，无新增 Bug 报告，表明项目近期未引入明显的回归问题。项目维护者在推动 Bug 修复和功能增强方面显示出积极迹象，但社区对于部分功能的易用性和稳定性仍有期待。

### 2. 版本发布

无

### 3. 项目进展

今日无已合并或关闭的 PR。但以下两个待合并的 PR 是项目推进的关键信号：

- **修复 CLI 交互体验 (PR #960)**：此 PR 解决了长期困扰用户的 CLI 控制字符问题。它引入了一个轻量级、零分配的行编辑器，支持在终端环境下的上下键历史导航、光标左右移动等标准操作。这直接提升了命令行用户的核心交互体验，是影响广泛的用户体验优化。
    - **链接**: [PR #960 - fix(cli): handle arrow keys in agent REPL](https://github.com/nullclaw/nullclaw/pull/960)

- **增强记忆系统可配置性 (PR #961)**：这是一个重要的功能增强。它为 `memory` 配置块增加了三个新选项：`auto_recall`（可选是否跳过记忆检索）、`recall_limit`（控制每次注入的记忆条目上限，默认5条）和 `max_context_bytes`（控制记忆内容的最大上下文字节数）。这为用户提供了更精细的控制权，尤其是在上下文窗口有限或希望优化性能的场景下尤为重要。
    - **链接**: [PR #961 - feat(memory): add configurable auto-recall, recall_limit, max_context_bytes](https://github.com/nullclaw/nullclaw/pull/961)

这两个 PR 的合并将标志着项目在 **“流程优化”** 和 **“系统可配置性”** 两个维度上的重要进步。

### 4. 社区热点

今日社区讨论热度不高，但以下两个议题值得关注：

- **调度器认证问题 (Issue #915)**：该 Issue 讨论了在使用外部 Ollama 主机时，调度器功能（Scheduler）无法正常工作的问题。虽然评论数不多，但该问题已存在一个月，且至今未有修复 PR，已成为一个影响特定用户群体（特别是依赖外部 LLM 服务的用户）的“顽疾”。
    - **链接**: [Issue #915 - [bug] Problem with scheduler unauthorized](https://github.com/nullclaw/nullclaw/issue/915)

- **Web UI 配置困惑 (Issue #861)**：用户对如何在没有图形界面的 VPS 上启用 Web UI 感到困惑。该问题的标签是“How to enable...”，这表明问题不在于 Bug，而在于文档的易用性和清晰度。用户表达了希望获得“非专业术语”的、针对“无头服务器”的具体指导。这反映了项目文档在指导用户进行一些高级配置（如 Web UI 的隧道浏览器）时存在短板。
    - **链接**: [Issue #861 - How to enable the Web UI on headless VPS server?](https://github.com/nullclaw/nullclaw/issue/861)

**分析**：社区热点集中在 **“功能性障碍”**（调度器）和 **“入门门槛”**（Web UI 配置）两个方面。前者是稳定性的痛点，后者是用户体验的瓶颈。

### 5. Bug 与稳定性

今日无新增 Bug 报告。目前在追踪的 3 个开放 Issue 均属于历史遗留问题，且均为 Bug 或使用问题，按严重程度排列如下：

1.  **严重 Bug: 调度器未授权 (Issue #915)**
    - **状态**: 开放，无关联修复 PR。
    - **影响**: 核心功能“调度器”在特定配置（外部 Ollama）下完全失效，严重影响用户对该功能的信任和使用。
    - **链接**: [Issue #915](https://github.com/nullclaw/nullclaw/issue/915)

2.  **可用性 Bug: CLI 控制字符 (Issue #865)**
    - **状态**: 开放，已有修复 PR #960 待合并。
    - **影响**: CLI 交互体验严重受损，不符合行业标准。PR #960 已直接针对此问题给出了解决方案。
    - **链接**: [Issue #865](https://github.com/nullclaw/nullclaw/issue/865)

3.  **文档/配置问题: Web UI 配置困难 (Issue #861)**
    - **状态**: 开放，无关联修复 PR。
    - **影响**: 阻碍了非图形界面环境下的用户部署 Web UI，降低了产品的可及性。
    - **链接**: [Issue #861](https://github.com/nullclaw/nullclaw/issue/861)

**稳定性信号**：项目近期未引入新的严重错误，稳定性尚可。但核心的调度器 Bug 若长期得不到解决，将对项目口碑产生负面影响。

### 6. 功能请求与路线图信号

今日无明确的新功能请求 Issue。但 **PR #961** 是一个非常重要的信号，它直接响应了用户对于内存系统进行更细致控制的需求。这个 PR 允许用户：

- **关闭自动回忆**（`auto_recall: false`）: 适用于隐私敏感或不需要长期记忆的场景。
- **调整回忆条目数量**（`recall_limit`）: 平衡对话质量与上下文窗口消耗。
- **控制回忆内容大小**（`max_context_bytes`）: 精确控制 token 使用量。

这表明开发者正积极吸纳和实现 **“可配置性”** 和 **“性能优化”** 方向的社区需求。此功能极有可能被包含在下一个正式版本中。

### 7. 用户反馈摘要

从今日活跃的 Issues 评论中提炼的用户反馈：

- **痛点 1：调度器失效**。用户 `scabros` 在 Issue #915 中报告，其工作流程（LLM 正常、工具调用正常）因调度器故障而中断，这是一个“局部可用”的典型负面体验。
- **痛点 2：CLI 体验差**。 用户 `eabase` 在 Issue #865 中报告控制字符问题，直接影响了其使用内建命令行工具的日常操作效率。
- **痛点 3：文档可读性不足**。 用户 `eabase` 在 Issue #861 中明确表达了对 README 文档“70%看不懂”的反馈，并请求“非专业术语”的解释，这反映了专业用户（能部署 VPS）在面对不友好的文档时的挫败感。

### 8. 待处理积压

以下 Issue 长期未获得有效响应或推进，提醒维护者关注：

- **Issue #915 - [bug] Problem with scheduler unauthorized** (创建于 **2026-05-15**)
    - **摘要**：已在多个段落中提及，这是一个严重且持续了34天的 Bug。用户已提供了详细的配置信息（Ubuntu, Ollama, Qwen3.6），但没有得到开发者的明确回复或指派修复。
    - **链接**: [Issue #915](https://github.com/nullclaw/nullclaw/issue/915)

- **Issue #861 - How to enable the Web UI on headless VPS server?** (创建于 **2026-04-22**)
    - **摘要**：这是一个典型的文档改进和用户帮助类 Issue。用户表达了清晰的痛点和期望（“human terms”），但自创建近两个月后，依然只有1条评论，没有得到指引或文档更新。
    - **链接**: [Issue #861](https://github.com/nullclaw/nullclaw/issue/861)

---

**总结**：NullClaw 项目今日在代码层面的进展积极（两个高质量 PR），但社区沟通和用户支持方面存在滞后。解决积压的核心 Bug 和改善文档易用性，将是提升项目健康度的两个关键突破口。

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

好的，这是根据您提供的 IronClaw 项目 GitHub 数据生成的 2026-06-18 项目动态日报。

---

## 🗞️ IronClaw 项目动态日报 | 2026年6月18日

### 1. 今日速览

今日项目活跃度**非常高**。社区贡献活跃，共产生 **50 条 PR** 和 **11 条 Issues**。核心开发团队正密集推进“Reborn Projects”这一大型功能（5/5的PR栈皆已提交），同时社区贡献者也积极修复“无进展检测”、“自动模型重试”等关键稳定性问题。值得注意的是，有 **33 条 PR** 处于待合并状态，可能预示着项目正等待一次集中评审和合并窗口。整体来看，项目正处于功能密集开发与稳定性修复并重的阶段，健康状况良好。

### 3. 项目进展

今日合并/关闭了 **17 条 PR**，其中一些关键进展如下：

- **核心功能修复：** 合并了 **PR #5052**，修复了 Slack OAuth 路径中的一个关键安全漏洞（对应 Issue #5009），确保了实时（非触发）Slack OAuth 流程与触发式流程具有同等的结构性DM安全校验。
- **代理循环增强：** 合并了 **PR #5022** 和 **PR #5000**，这两者是“无进展检测”重设计的第2和第3部分。`PR #5000` 添加了记录每一步能力的输出摘要（`ContentDigest`）的底层框架，为后续更智能的检测奠定了基础。
- **自动化发布流程：** **PR #3708**（chore: release）已合并，为项目准备了一次新版本发布，包含 `ironclaw_common` 和 `ironclaw_skills` 的破坏性API变更。

### 4. 社区热点

今日讨论最活跃的内容主要集中在 **Slack OAuth 安全性和第一次使用体验** 上。

- **`#5009` [Security] Slack OAuth路径安全加固**：该Issue详细描述了如何确保非触发的Slack OAuth流程也具备与触发式流程同等级别的DM安全校验。开发者和安全审查员对此进行了深入讨论，显示出社区对集成安全性的高度重视。已通过 PR #5052 解决。
- **`#4793` [Question] 首次运行引导是否应屏蔽“扩展”和“自动化”？**：该 Issue 探讨了用户首次安装后的引导流程，发现未配置提供者时，导航到“扩展”或“自动化”会重定向回欢迎页。社区讨论的核心在于如何平衡用户体验与安全引导。

### 5. Bug 与稳定性

今日报告了多项 Bug，涉及功能异常、UI 缺陷和外部服务集成问题。按严重程度排列如下：

- **严重 - 功能阻断：**
    - **`#5060` [Bug] GitHub 分析工作流陷入重复审批循环**：导致工作流永远无法产生结果。当前无修复PR。
    - **`#5058` [Bug] AWS Bedrock 在独立二进制文件中不可用**：由于编译特性（feature）和工具 schema 问题，导致用户无法使用 Bedrock。**已有修复PR #5059**。
- **中等 - 功能异常：**
    - **`#4761` [Bug] 代理在工具反复失败后停止运行，无法恢复**：影响代理的鲁棒性和任务完成率。当前无修复PR。
    - **`#5007` [Bug] 技能验证错误在填写必填字段后不会清除**：影响用户创建技能时的交互体验。当前无修复PR。
- **低 - UI/交互问题：**
    - **`#4823` [Bug] 删除运行中的会话无UI反馈**（已关闭）。
    - **`#4974` [Bug] 活动工具行显示重复的“...”操作按钮**（已关闭）。

### 6. 功能请求与路线图信号

今日提交的 PR 和 Issues 明确指示了项目的几个短期路线图方向：

- **Agent 自我进化能力：** **PR #5061** 提出了一个大型新功能——“技能提取与自我进化”。该 PR 旨在让代理在成功完成任务后，自动将过程提炼为可复用的技能（SKILL.md）并安装，使代理能够“边用边学”。这是一个重要的路线图信号，表明项目正朝着更智能、更自主的 AI Agent 方向发展。
- **“项目”管理功能（Reborn Projects）：** `ilblackdragon` 提交了一个由 **5 个 PR 组成的栈**（`#5015` 至 `#5019`），系统地实现了在 Reborn 版本中的“项目”实体、服务端口、HTTP 接口和前端页面。这表明项目管理工作流将成为 Reborn 平台的核心能力。
- **内部文件系统可视化：** **PR #5057** 增加了代理内部文件系统的只读浏览器，方便用户查看代理的内存存储和项目目录，提升了可观测性。

### 7. 用户反馈摘要

- **关于代理恢复能力：** 用户 `sunglow666` 在 Issue `#4761` 中描述了一个典型场景：当代理连续多次工具调用失败时，会直接停止运行而不会尝试恢复。这反映出在复杂的自动化任务中，用户对 Agent 的容错性和自愈能力有更高期待。
- **关于UI/UX困惑：** 用户 `sunglow666` 在 Issue `#5007` 中抱怨技能创建表单的验证错误提示不会即时消失，需要提交后才能清除。这种“滞后”的验证反馈造成了困惑，用户希望UI能提供更即时的反馈。
- **关于首次使用体验：** 用户 `sunglow666` 在 Issue `#4793` 中提出了一个关键问题：是应该让用户自由探索新功能（如扩展、自动化），还是应该在配置好基础模型之前对其进行屏蔽？这反映了新用户引导与功能可发现性之间的矛盾。

### 8. 待处理积压

- **`#4879` [Open] 本地“Dogfooding”问题跟踪**：这是一个持续跟踪周（6/15-6/21）的 Issue，记录了使用本地 Reborn WebUI 时发现的启动、配置等问题。虽然活跃，但作为一个为期一周的跟踪项，需要确保所有发现的问题都能被分解为独立的、可追踪的 Issue，避免跟踪项本身被长期搁置。
- **`#4761` [Open] 代理反复工具失败后不恢复**：该 Issue 创建于6月11日，虽然核心团队可能已知晓，但至今没有关联的修复 PR 或明确的解决方案，且有3条评论，说明这是一个社区关心的痛点，应引起维护者的关注。
- **`#4793` [Open] 首次运行引导问题**：该 Issue 涉及新用户引导，如无明确结论或设计决策，可能会影响新用户的留存率，建议维护者尽快就此进行讨论并给出方向性指引。

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI 项目动态日报 — 2026-06-18

---

## 1. 今日速览

- 项目在过去24小时内高度活跃：共关闭10个Pull Request，发布1个新版本（2026.6.15），新开1个安全相关Issue。
- 版本更新带来了**computer use**、**实时ASR语音输入**及**上下文压缩连续性改进**三个重要能力，标志着LobsterAI向多模态Agent方向迈出关键一步。
- 新开的Issue #2176披露了一个**严重任意文件读取漏洞**，涉及消息自动加载机制，已引起社区关注。
- 开发团队持续清理技术债务，多个跨模块fix（模型选择、停止流、OOM崩溃等）密集合并，项目健康度良好。

---

## 2. 版本发布

**LobsterAI 2026.6.15**（发布于2026-06-15）

### 新增功能
- **Computer Use（计算机操作）**：Agent可直接控制桌面或应用，扩展自主任务执行能力。
- **实时ASR语音输入**：Cowork模块新增低延迟语音识别，支持会话中实时转写。
- **上下文压缩连续性改进**：针对OpenClaw历史压缩后的上下文丢失问题，增加了LobsterAI自有的连续性保障层，确保长任务可可靠恢复。

### 破坏性变更
- 无明确破坏性变更报告。但若使用了自定义`MEDIA:`文件引用，请注意新版本的安全更新（见下文Issue #2176），可能需调整配置或升级权限控制。

### 迁移注意事项
- 如果之前启用了自动工件加载，建议立即审查`MEDIA:`解析逻辑，或暂时关闭该功能直到官方修复。
- 语音输入依赖ASR服务，确保服务端已部署实时ASR模型。

---

## 3. 项目进展

今日合并/关闭的PR共10条，主要覆盖以下方向：

| 方向 | PR | 关键变化 |
|------|----|---------|
| **稳定性/性能** | #2149 | 提升OpenClaw网关进程V8堆内存上限，减少长时多通道OOM崩溃。 |
| | #2147 | 修复用户停止操作与启动轮次之间的竞态条件，避免误发聊天消息。 |
| | #2174 | 修正自动滚动到底部的定位偏移，优化会话切换时的清理逻辑。 |
| **功能完善** | #2145 | 改进Cowork上下文压缩后任务连续性，增加诊断、任务状态、轻量工作空间等。 |
| | #2153 | 修复同名称包/自定义模型选择冲突，保持显式的`lobsterai-server/...`引用。 |
| | #2154 | 停止流后保留模型元数据，手动停止的回复不再丢失模型信息。 |
| | #2162 | 合并后保留语音输入取消保护与诊断日志。 |
| | #2144 | 更新门户fallback URL，切换生产/测试域名。 |
| **文档与清理** | #2175 | 优化README文档。 |
| | #1463 | 修复长模态标题截断问题（旧PR，今闭）。 |

整体来看，项目在**对话体验精细化**（滚动、模型选择、语音输入）、**资源管理**（堆内存、停止竞争）和**核心Agent能力**（computer use、ASR）上均有实质性推进，代码库质量持续提升。

---

## 4. 社区热点

**#2176 [Security] LobsterAI automatic artifact loading allows message-derived arbitrary local file reads**  
- 作者: YLChen-007  
- 创建: 2026-06-18 | 评论: 1 | 👍: 0  
- 链接: [Issue #2176](https://github.com/netease-youdao/LobsterAI/issues/2176)

该Issue报告了一个**严重安全漏洞**：LobsterAI自动解析助手或工具输出中的`MEDIA:`文件引用，并将生成的路径传递给Electron的渲染进程，导致攻击者可利用消息内容读取本地任意文件。虽然评论数不多，但这是今日唯一新开Issue，且涉及安全红线，预计很快会引发维护者和用户的高度关注。社区诉求是尽快获得补丁修复和临时缓解措施。

---

## 5. Bug 与稳定性

| 严重程度 | Issue/PR | 描述 | 状态 |
|----------|----------|------|------|
| **严重** | #2176 | 自动工件加载导致任意本地文件读取（Electron权限绕过） | 新开，无修复PR |
| 高 | #2149 (PR) | OpenClaw网关进程在长时多通道场景下OOM | ✅ 已修复（提升堆限制） |
| 中 | #2147 (PR) | 停止操作与启动轮次竞态导致误发消息 | ✅ 已修复（取消启动并发送空闲状态） |
| 中 | #2154 (PR) | 停止流后模型元数据丢失 | ✅ 已修复（在清理前完成流信息） |
| 低 | #2174 (PR) | 自动滚动到底部位置偏差、会话切换未清理定时器 | ✅ 已修复 |

今日未发现回归类Bug。所有已修复问题均通过明确的PR解决，稳定性风险可控。

---

## 6. 功能请求与路线图信号

从已合并的PR和版本更新中，可看出以下功能正在被积极开发并很可能纳入下一版本：

- **Computer Use（计算机操作）**：在2026.6.15版本中首次引入，路线图优先级较高。
- **实时ASR语音输入**：Cowork模块已支持，未来可能扩展多语言和降噪。
- **上下文压缩连续性增强**：针对长对话场景，持续优化Agent记忆能力。
- **门户域名迁移**：PR #2144已完成URL切换，表明团队正在更新基础设施。
- **模型选择精细化**：PR #2153修复了同名包模型冲突，暗示用户对多模型管理有需求。

社区当前未在Issue中提出新的功能请求（除安全修复）。如果后续有用户建议，可能会集中在对Computer Use的更细粒度控制、语音输入的离线支持等方面。

---

## 7. 用户反馈摘要

- **安全反馈（Issue #2176）**：用户YLChen-007以安全研究者身份报告了一个Electron端任意文件读取漏洞，描述清晰，包含攻击向量（通过`MEDIA:`引用）。该反馈属于负面但专业，暴露出自动工件加载机制的权限分级不足。
- 其余PR和Issue无用户评论或讨论。鉴于项目开源且由网易有道主导，用户反馈主要通过Issue渠道表达；今日整体反馈量低，但安全反馈质量高。

---

## 8. 待处理积压

- **#2176 [Security] 任意文件读取漏洞**：当前无任何PR或Assignment。这是**最高优先级**待处理项，建议维护者在24小时内确认漏洞、发布安全公告（临时禁用`MEDIA:`或添加路径白名单），并关联修复PR。
- **#1463（已关闭）** 属于老旧积压，已于今日关闭，无需进一步关注。
- 无其他长期未响应的公开Issue或PR。

> 注：数据基于2026-06-18 00:00 UTC至2026-06-18 24:00 UTC的GitHub活动。

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

# TinyClaw 项目日报 — 2026-06-18

## 今日速览

过去24小时内，TinyClaw项目**新增3个安全相关Issue**（全部由同一研究员提交），均为未修复的严重漏洞；**无新Pull Request、无版本发布**。项目整体活跃度集中于**安全审计反馈**，但缺乏对应修复进展，社区响应有待提升。三个漏洞均涉及未认证接口或不可信输入导致的数据泄露问题，需维护团队优先评估并加速修复。

---

## 版本发布

无新版本发布。

---

## 项目进展

今日无任何Pull Request被合并或关闭，项目在功能推进、Bug修复及代码合并方面**无可见进展**。

---

## 社区热点

今日唯一的活动集中于3个安全Issue，全部为同一作者（@YLChen-007）提交，虽暂无评论或点赞，但其内容揭示了三个**高危乃至严重级别的未认证攻击向量**：

- **#284** [Security] TinyAGI allows unauthenticated API messages to invoke Claude with provider permission checks disabled by default  
  - 详述：未认证的 `POST /api/message` 接口可直接将攻击者控制的消息转发至 Claude 提供者，且默认关闭权限检查。  
  - 链接：https://github.com/TinyAGI/tinyagi/issues/284

- **#283** [Security] Unauthenticated `prompt_file` agent configuration allows arbitrary local file disclosure to the model provider  
  - 详述：未认证的 agent 配置接口接受攻击者控制的 `prompt_file` 参数，可导致任意本地文件泄露至模型提供商。  
  - 链接：https://github.com/TinyAGI/tinyagi/issues/283

- **#282** [Security] Untrusted `[send_file: ...]` response tags allow arbitrary host file attachment delivery  
  - 详述：不安全的响应附件的处理逻辑使攻击者能通过控制模型输出中的 `[send_file:...]` 标签，强制将任意主机文件作为附件发送。  
  - 链接：https://github.com/TinyAGI/tinyagi/issues/282

**分析**：这三个Issue构成了一个**典型的未认证攻击链**，从消息注入、配置篡改到响应劫持，覆盖了TinyClaw核心的API、配置和输出处理模块。社区潜在诉求是**快速确认漏洞影响范围并制定修复计划**。建议维护者立即启动安全应急流程，考虑暂时关闭受影响接口或增加认证与输入校验。

---

## Bug 与稳定性

| 严重程度 | Issue # | 标题摘要 | 状态 |
|----------|---------|----------|------|
| **严重（Severe）** | #284 | 未认证API消息可调用Claude（默认绕过权限检查） | Open，无Fix PR |
| **严重（Severe）** | #283 | 未认证配置接口允许任意本地文件泄露 | Open，无Fix PR |
| **严重（Severe）** | #282 | 不可信 `[send_file:]` 标签可强制发送任意主机文件 | Open，无Fix PR |

所有三个漏洞均为**远程未认证攻击**，可在无任何权限下导致数据泄露或代理滥用，需作为最高优先级处理。

---

## 功能请求与路线图信号

今日无任何功能请求或路线图相关Issue/PR。项目当前聚焦安全事件，暂无新功能规划可见。

---

## 用户反馈摘要

由于三个Issue均无评论，尚无用户反馈可供提炼。但提交者 `@YLChen-007` 的漏洞描述详细且具有实际攻击场景，可能暗示真实环境已存在风险。建议维护者主动在下游社区（如 Discord、邮件列表）收集用户是否遭遇类似异常行为。

---

## 待处理积压

今日新增的三个安全Issue均为**当天提交、当天未响应**，属于紧急积压。建议维护者在24小时内至少对每个Issue做出**确认回复**（如标注受影响版本、评估严重性、给出预计修复时间线或临时缓解措施）。此外，项目此前是否存在尚未响应的长期Issue/PR？当前数据未提供，请维护者自行检查。

---

**项目健康度提示**：连续无代码合并、无版本发布，叠加新增三个严重安全漏洞，社区可能对项目安全治理能力产生疑虑。建议尽快发布安全修复补丁（哪怕仅针对关键路径进行hotfix），并公开安全响应流程以重建信任。

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis 项目动态日报 — 2026-06-18

## 今日速览
过去 24 小时内，Moltis 项目共产生 **2 个新 Issue** 和 **1 个待合并 PR**，无版本发布。社区活跃度中等偏低，主要聚焦于用户界面与输出格式的灵活性增强。两个新 Issue 均为功能请求（enhancement），尚无 Bug 或回归报告。一个围绕 WebUI RPC 超时配置的 PR 已提交并等待审查，表明项目在基础设施可配置性方面持续推进。整体项目健康度良好，但合并节奏较缓。

---

## 版本发布
今日无新版本发布。

---

## 项目进展
今日 **无 PR 被合并或关闭**，仅有 1 个开放的 PR：

- **[#1130] feat: make webui rpc timeout configurable**（作者：khimaros）  
  该 PR 旨在将 WebUI 的 RPC 超时时间从硬编码改为可配置项，并关联修复了 Issue #1127（该 Issue 未出现在今日数据中，推测为已关闭的 Bug）。  
  链接：https://github.com/moltis-org/moltis/pull/1130

该 PR 若被合并，将提升 WebUI 在高延迟场景下的稳定性，可视为对基础设施层的细小改进。

---

## 社区热点
今日讨论最活跃的 Issue 为 **#1126**，共 3 条评论，作者为 khimaros。

- **[#1126] [Feature]: allow to configure the format of tts output**  
  用户提出希望允许配置文本转语音（TTS）的输出格式（如文件类型、采样率等）。评论区暂未披露具体细节，但从标题和仓库定位（AI 智能体与个人助手）可推测，该需求来源于用户对多模态输出控制的普遍期待。  
  链接：https://github.com/moltis-org/moltis/issues/1126

另一新 Issue #1131（Markdown 复制/导出功能）暂无评论，但缺乏初始讨论热度。

---

## Bug 与稳定性
今日 **未报告任何 Bug、崩溃或回归问题**。项目稳定性表现良好，无紧急修复需求。

*注意：PR #1130 关联的 Issue #1127 可能涉及 WebUI 的稳定性问题，但该 Issue 未被收录在今日数据中，推测已在更早时间解决。*

---

## 功能请求与路线图信号
今日两个新 Feature Request 均来自社区用户，反映了对**输出端灵活性和内容导出能力**的诉求：

1. **#1126** – 允许配置 TTS 输出格式  
   - 涉及音频后处理与输出多样性，可能需扩展平台层配置。  
2. **#1131** – 添加复制 + 导出为 Markdown  
   - 类似功能常见于聊天/助手类工具，便于用户记录对话。

结合今日 PR #1130（WebUI RPC 超时可配置），可见社区当前关注点在于**使项目更易集成和个性化**。这些功能请求有较大概率被纳入下一迭代（v0.x 或 v1.0 后续版本），但需维护团队评估实现成本。

---

## 用户反馈摘要
由于今日仅有 2 个新 Issue，且评论内容未完整展示，从简短摘要可提取以下用户痛点：

- **TTS 输出控制缺失**：Issue #1126 的提交者明确表示当前缺少对 TTS 输出格式的配置能力，可能影响在多场景（如流式播放、文件保存）下的兼容性。  
- **数据便携性需求**：Issue #1131 提出复制和导出为 Markdown，暗示用户希望将对话内容带出系统，用于二次编辑或存档，这是个人助手类工具常见的实用场景。

暂无负面评价或严重不满意反馈。

---

## 待处理积压
今日数据中无长期未响应的关键 Issue 或 PR。以下为值得关注的**待办项**：

- **PR #1130 待审查** – 该 PR 来自活跃作者 khimaros，且关联已修复的 #1127，建议维护者尽快 Review 并合并，以释放可配置超时功能。  
- **Issue #1126** 虽已有 3 条评论，但尚未标记为“计划中”或分配标签；若该需求获得社区 +1，建议纳入路线图投票。

整体来看，Moltis 项目活跃度维持稳定，社区需求清晰，维护者响应速度尚可。建议未来增加对功能请求的优先级标注，以帮助贡献者聚焦。

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw 项目动态日报 (2026-06-18)

## 1. 今日速览
项目在 **2026-06-18** 继续保持高活跃度，24 小时内共处理 **45 条 Issue**（新开/活跃 11，关闭 34）和 **50 条 PR**（待合并 17，合并/关闭 33）。**正式发布 v1.1.12**，带来控制台模型页面重构、简单模式等多项新特性，同时启动 **v2.0.0-alpha 里程碑**（PR #5281）。社区反馈集中于 WebUI 性能（#4108）、升级后技能重置（#5262）以及跨 Agent 死循环（#5204）等稳定性问题。项目整体健康状况良好，维护响应积极。

## 2. 版本发布

### v1.1.12 (正式版)
- **主要新增：**
  - **模型页面重构**：采用统一的 Provider 聚合卡片 UI 和布局重设计（#5203）
  - **简单模式**：扁平导航、按更新时间排序的会话列表（#5222）
  - **会话过滤**：支持按标题搜索会话（#5178，合入此版本）
- **性能优化**：移除 agent 配置中不必要的深层拷贝（#5240）
- **发布号**：从 `v1.1.12-beta.2` 正式升为 `v1.1.12`
- **兼容性**：无破坏性变更，可通过 `pip install --upgrade qwenpaw` 或直接替换 Docker 镜像升级。若使用自定义部署，需注意 `uv tool install` 环境下钉钉渠道的 SSL 证书问题（已提供修复 PR #5291）
- **迁移注意**：升级后内置技能可能恢复默认启用状态，需手动重新禁用（已知问题 #5262，尚未完全修复）

### v1.1.12-beta.2 (预发布，已被 v1.1.12 取代)
- 包含上述部分功能的前置提交，无需独立部署。

## 3. 项目进展

### 已合并/关闭的重要 PR

| PR | 说明 | 状态 |
|----|------|------|
| [#5280](https://github.com/agentscope-ai/QwenPaw/pull/5280) | 发布 v1.1.12 版本 | 已合并 |
| [#5281](https://github.com/agentscope-ai/QwenPaw/pull/5281) | 版本号升至 `2.0.0a1`，开启 AgentScope 2.0 alpha 里程碑 | 已合并 |
| [#5288](https://github.com/agentscope-ai/QwenPaw/pull/5288) | 修复预发布脚本的参数展开问题，并更新版本号 | 已合并 |
| [#5289](https://github.com/agentscope-ai/QwenPaw/pull/5289) | 修复 ChromaDB 探针集合名称为 `probe-test`，避免混淆 | 已合并 |
| [#5275](https://github.com/agentscope-ai/QwenPaw/pull/5275) | 修复主动响应模块缓存污染，防止 `load_agent_config()` 被意外修改 | 已合并 |
| [#5274](https://github.com/agentscope-ai/QwenPaw/pull/5274) | 重构 XiaoYi 渠道为双 WebSocket 连接，修复通道不可用问题 | 已合并 |
| [#5026](https://github.com/agentscope-ai/QwenPaw/pull/5026) | 修复当 user_id 与 session_id 相同时文件名重复的问题 | 已合并 |
| [#5041](https://github.com/agentscope-ai/QwenPaw/pull/5041) | 备份时跳过不可读文件，避免整个备份失败 (Windows 权限问题) | 已合并 |
| [#3839](https://github.com/agentscope-ai/QwenPaw/pull/3839) | 修复 XiaoYi A2A 协议实现与测试用例 | 已合并 |
| [#5277](https://github.com/agentscope-ai/QwenPaw/pull/5277) | 更新路线图文档 | 已合并 |

**项目进展总结**：项目在 24 小时内完成了 v1.1.12 正式版本的发布，同时宣布启动 2.0 大版本开发计划。超过 10 个重要修复合入主干，涵盖连接稳定性、备份可靠性、内存管理、缓存一致性等核心领域。XiaoYi 渠道获得重大重构，提升了 A2A 协议的兼容性。

## 4. 社区热点

| 链接 | 标题 | 评论数 | 关注点 |
|------|------|--------|--------|
| [#280](https://github.com/agentscope-ai/QwenPaw/issues/280) | Discussion: Which Skills and MCPs Can Be Built-in? | 27 | 讨论内置技能与 MCP 的预装策略，虽是 3 月发起但今日仍有更新，社区对“开箱即用”需求强烈 |
| [#4108](https://github.com/agentscope-ai/QwenPaw/issues/4108) | [Bug] 新版本 WebUI 非常卡顿 | 8 | 用户反馈 v1.5.1.post2 在生成回复时导致整个系统卡顿，Win11+32GB 内存仍有问题，引发社区共鸣 |
| [#2229](https://github.com/agentscope-ai/QwenPaw/issues/2229) | [Bug] failed to forward request（Docker 升级后启动失败） | 7 | NAS 容器化部署用户升级 0.2 后前端报错，影响范围较大 |

**分析**：#280 虽已关闭，但 27 条评论表明用户对预置技能生态非常关注，建议项目团队参考此讨论结果优化默认技能集。#4108 暴露了新版 WebUI 在高负载下的性能瓶颈，可能是 React 渲染或消息流处理问题，需要紧急排查优化。

## 5. Bug 与稳定性

### 严重级

| Issue | 描述 | 严重性 | 修复进展 |
|-------|------|--------|----------|
| [#5262](https://github.com/agentscope-ai/QwenPaw/issues/5262) | 每次升级后被禁用的内置技能自动重新启用 | ⚠ 中等 | 社区抱怨第二次出现，暂无关联 PR |
| [#5204](https://github.com/agentscope-ai/QwenPaw/issues/5204) | 两个 QwenPaw Agent 通过 Matrix 互聊陷入无限循环 | 🔴 高 | OPEN，无明确修复方案，属于架构级问题 |
| [#5295](https://github.com/agentscope-ai/QwenPaw/issues/5295) | Subagent 的审批请求未推送到外部渠道（QQ 等） | 🔴 高 | OPEN，影响安全审批流程 |
| [#5292](https://github.com/agentscope-ai/QwenPaw/issues/5292) | 手动添加的模型提供方不出现在对话页面模型选择中 | 🟡 低 | 已关闭（可能为重复或已修复） |
| [#5290](https://github.com/agentscope-ai/QwenPaw/issues/5290) | 安装 discord.py 2.7.1 后 huggingface_hub import 失败 | 🟡 低 | OPEN，依赖版本冲突 |

### 当日修复的稳定性相关 PR
- [#5275](https://github.com/agentscope-ai/QwenPaw/pull/5275)：修复 `load_agent_config()` 缓存污染
- [#5289](https://github.com/agentscope-ai/QwenPaw/pull/5289)：修复 ChromaDB 探针集合名错误
- [#5041](https://github.com/agentscope-ai/QwenPaw/pull/5041)：跳过不可读文件避免备份失败
- [#5274](https://github.com/agentscope-ai/QwenPaw/pull/5274)：修复 XiaoYi 渠道不可用

## 6. 功能请求与路线图信号

| Issue/PR | 功能描述 | 状态 | 技术判断 |
|----------|----------|------|----------|
| [#4965](https://github.com/agentscope-ai/QwenPaw/issues/4965) | 合并同一品牌（如智谱）的多个 Provider 卡片为单个下拉选择 | ✅ 已关闭，很可能已在 v1.1.12 中实现 | 与 #5203 Models Page Overhaul 相关 |
| [#2202](https://github.com/agentscope-ai/QwenPaw/issues/2202) | 公开 HTTP API，允许直接调用而不依赖 WebUI | ❌ 已关闭但未实现，仅记录 | 社区呼声较高，可考虑在 2.0 路线图中加入 |
| [#2235](https://github.com/agentscope-ai/QwenPaw/issues/2235) | 通过 Web 控制台远程升级 CoPaw | ❌ 已关闭 | 需评估安全性，可能纳入 2.0 计划 |
| [#1366](https://github.com/agentscope-ai/QwenPaw/issues/1366) | 增加定时任务的历史执行记录 | ❌ 已关闭，用户有本地实现但未合并 | 建议合入主干，提升可观测性 |
| [#5293](https://github.com/agentscope-ai/QwenPaw/pull/5293) | 对话框历史改为右侧内嵌面板（永久侧边栏） | 🆕 OPEN | 改善多会话切换体验，预计合入下个小版本 |
| [#5276](https://github.com/agentscope-ai/QwenPaw/pull/5276) | 新增 `qwenpaw migrate openclaw` CLI 工具，从 OpenClaw 生态导入配置 | 🆕 OPEN | 吸引其他平台用户迁移，有战略价值 |
| [#5210](https://github.com/agentscope-ai/QwenPaw/pull/5210) | 新增 `cron update` 命令编辑已有定时任务 | 🆕 OPEN | 后端已支持 PUT，社区期待该 CLI |

**路线图信号**：v2.0.0a1 的启动表明项目将开始重构或引入重大新功能。建议优先处理跨 Agent 死循环（#5204）和审批推送（#5295）等架构问题，同时收集用户对 HTTP API、远程升级等高票需求的反馈。

## 7. 用户反馈摘要

- **性能不满**：用户 `pengliang159` 在 #4108 中描述“新版本 WebUI 卡到鼠标掉帧”，恢复到旧版才正常，说明优化工作仍需加强。
- **升级体验差**：用户 `daigoopautoy` 在 #5262 中抱怨每次升级都需要手动禁用不需要的内置技能，“体验越来越差”，建议维护用户偏好持久化。
- **安装依赖问题**：用户 `wangqiangqiangcx` 在 #5290 中报告 `discord.py` 与 `huggingface_hub` 的 import 冲突，在 Windows 打包环境下尤其突出，需排查依赖锁定策略。
- **跨 Agent 协作痛点**：用户 `laeni` 在 #5204 中详细描述了通过 Matrix 互聊导致的无限循环，并区分了与单 Agent 死循环的不同，说明高级用户对 Agent 间通信有明确需求。
- **积极贡献**：多位首次贡献者提交了高质量 PR（如 #5241 定时任务宽容时间、#5287 摘要截断保护），社区参与度健康。

## 8. 待处理积压

| 链接 | 类型 | 创建时间 | 最后更新 | 问题描述 |
|------|------|----------|----------|----------|
| [#5204](https://github.com/agentscope-ai/QwenPaw/issues/5204) | 🐛 Bug | 2026-06-15 | 2026-06-18 | 跨 Agent 通过 Matrix 互聊无限循环，无明确修复方向 |
| [#5295](https://github.com/agentscope-ai/QwenPaw/issues/5295) | 🐛 Bug | 2026-06-18 | 2026-06-18 | Subagent 审批请求未推送至外部渠道，影响安全管控 |
| [#5262](https://github.com/agentscope-ai/QwenPaw/issues/5262) | 🐛 Bug | 2026-06-17 | 2026-06-18 | 升级后内置技能自动启用（第二次反馈） |
| [#4975](https://github.com/agentscope-ai/QwenPaw/pull/4975) | 🚀 PR | 2026-06-05 | 2026-06-18 | 可定制会话页面列顺序，已标记 Under Review 但长期未合入 |
| [#5241](https://github.com/agentscope-ai/QwenPaw/pull/5241) | 🚀 PR | 2026-06-16 | 2026-06-17 | 增加定时任务 misfire_grace_seconds 默认值，首次贡献者 PR 待审核 |

**建议**：项目维护者应优先关注 #5204 和 #5295 这两个影响 Agent 协作与安全的严重 bug；#5262 用户已第二次反馈，建议在 v1.1.13 加入配置持久化逻辑。同时，对标记“Under Review”且提交时间超过一周的 PR（如 #4975）尽快完成 review 并决定合入或关闭。

---

*报告生成时间：2026-06-18 18:00 UTC*  
*数据来源：[CoPaw GitHub 仓库](https://github.com/agentscope-ai/QwenPaw)*

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，作为 ZeroClaw 项目的 AI 分析师，我已根据您提供的 GitHub 数据，为您整理并生成 2026-06-18 的项目动态日报。

---

# ZeroClaw 项目日报 | 2026-06-18

## 1. 今日速览

项目今日呈现**高活跃度**态势。社区在代码提交（50条PRs）和问题反馈（17条Issues）上均有显著贡献，表明开发与社区双轮驱动模式运转良好。核心开发团队正密集推进 **v0.8.2 (Skills平台, WASM插件)** 和 **v0.9.0 (安全与认证)** 里程碑。今日重点在于**Windows平台兼容性修复**、**更新管道加固**以及**一系列严重级Bug的专项修复**，项目健康度稳步提升。尽管有1个严重问题被修复关闭，但仍有数个S1级工作流阻塞Bug待解，维护团队反应迅速。

## 2. ~~版本发布~~

无

## 3. 项目进展

今日有 **11个 PR 被合并或关闭**，其中包括一个重要的8连PR栈的一部分，标志着核心配置管理、网关API和技能平台功能取得重大进展。

- **配置管理与重命名级联功能落地**：
    - **合并** `#7842` (feat(cli): agents/providers/channels CRUD + skill-bundle cascade) 和 `#7841` (feat(gateway): agent owned-state rename cascade + rename wiring)。这两个PR是核心系列的第7和第8部分，为 `providers`、`agents` 和 `channels` 的删除/重命名提供了完整的端到端级联处理，显著提升了配置的健壮性与一致性。
    - **合并** `#7840` (feat(config): rename_with_cascade for aliased entries)。作为上述系列的核心组件，该PR实现了配置层级的重命名逻辑。
- **用户界面与体验改进**：
    - **合并** `#7684` (fix(acp): surface history-pruner and turn-cancel as visible events)。改进ACP协议的事件展示，将历史修剪和回合取消操作以可见事件形式呈现给用户，增强了UI的透明度与可理解性。

- **持续进行的修复工作（PR #7853）**：一个关于“修复Windows自我更新并加固更新管道”的PR正在活跃审查中。它解决了Windows上二进制文件替换的核心缺陷（由于文件锁定），并设计了更健壮的旁路启动方案。这是向多平台稳定发布迈出的关键一步。

总体来看，项目在**配置管理的深度**、**Windows平台的兼容性**和**用户操作反馈**上迈出了坚实的一步。

## 4. 社区热点

- **RFC: 原生上下文压缩**
    - **Issue**：`#7673` [RFC: Native context compression as a provider pipeline decorator](zeroclaw-labs/zeroclaw Issue #7673)
    - **热度**：评论数3，与v0.8.1 跟踪器并列今日最高。
    - **分析**：这是社区对LLM调用成本与响应延迟的主动优化提案。作者提出通过一个“压缩装饰器”在请求发送给大模型之前压缩上下文，旨在降低token消耗。该讨论反映了用户对运行成本**高度敏感**，并希望通过框架级别的标准化手段来解决，而非依赖各Provider的私有特性。这是提升框架经济性的重要信号。

## 5. Bug 与稳定性

今日报告了多个高影响度Bug，其中已有部分得到快速修复。

- **S1 - 工作流阻塞**：
    - **`#7907`** [Bug]: agent rename can move owned state before config persistence
        - **分析**：一个深刻的状态管理Bug。代理重命名操作在配置持久化成功前就移动了外部状态，若重命名中途失败，可能导致状态损坏和配置不一致。无fix PR关联，需高度关注。
        - **链接**: [zeroclaw-labs/zeroclaw Issue #7907](zeroclaw-labs/zeroclaw Issue #7907)
    - **`#7563`** **已关闭** [Bug]: canvas-store regression in WS chat/ACP sessions breaks /canvas after #6986
        - **分析**：这是一个严重的回归问题，导致基于WebSocket的AI对话界面出现空白。该问题已在今日被修复并关闭，说明团队对S1级Bug的响应和修复速度非常快。
        - **链接**: [zeroclaw-labs/zeroclaw Issue #7563](zeroclaw-labs/zeroclaw Issue #7563)

- **S2 - 功能降级**：
    - **`#7462`** [Bug]: 74 test failures on Windows
        - **分析**：这是一个持续存在的严重跨平台问题。CI缺失Windows环境，导致大量Unix依赖的测试命令和路径语义问题未被发现。已经有相关的修复工作（如PR #7853）在推进，但根治仍需CI基础设施的改进。
        - **链接**: [zeroclaw-labs/zeroclaw Issue #7462](zeroclaw-labs/zeroclaw Issue #7462)

- **主要修复性PR**：
    - **`#7909`** (fix(providers): include tool name in native tool-result messages)：修复了Groq原生工具调用因缺少 `name` 字段而失败的Bug。
    - **`#7908`** (fix(browser): repair WebDriver snapshot returns and CSS selector escaping)：修复了浏览器工具在WebDriver模式下快照返回空值和CSS选择器转义问题。
    - **`#7901`** (fix(runtime): bound repeated shell approval loops)：增加了对重复shell命令审批请求的本地限制，防止在用户确认前出现循环请求。
    - **`#7902`** (fix(tools): pin http_request requests to vetted DNS addresses)：为 `http_request` 工具增加了SSRF防护，通过DNS校验来对抗DNS重绑定攻击，显著提升了网络工具的安全性。

## 6. 功能请求与路线图信号

- **平台级功能成为主流**：
    - **`#7822`** (WASM plugin lifecycle hook subscriptions)：提议让WASM插件像内置Hook一样订阅Agent生命周期事件。这与v0.8.2的WASM插件计划高度契合，极有可能被采纳。
    - **`#7673`** (Native context compression)：正如社区热点所提，对标准化的、跨提供者的上下文压缩需求强烈，是平台工程化的一个标志性功能。
- **运维与安全能力增强**：
    - **`#7897`** (Apply security policy and channel config updates without full daemon reload)：体现了社区对生产环境零停机变更的迫切需求，是平台成熟度的关键指标。
    - **`#7883`** (Expose intra-family provider fallback notices)：提供了更精细的模型回退透明度，帮助用户控制成本和了解性能表现。
- **路线图协同**：
    - 这些新功能请求与当前**v0.8.2 (Skills/WASM)** 和 **v0.9.0 (安全/认证)** 的里程碑高度吻合，表明社区需求与团队规划保持一致。

## 7. 用户反馈摘要

- **痛点**：
    - **Windows兼容性**：用户 `state-space-swarm` 报告在Android Termux安装失败（`#7911`），用户 `NiuBlibing` 指出Windows有74个测试失败（`#7462`），修复PR `#7853` 的讨论也详细揭示了Windows更新功能的脆弱性都是用户面临的主要障碍。
    - **功能可用性**：用户 `tidux` 的PR `#7903` 指出 `session/messages` 端点无法正确返回ACP协议会话的历史记录，影响调试与审计。
- **使用场景**：
    - **Web UI体验**：用户 `vrurg` 报告的 Canvas 回归问题（`#7563`），表明浏览器工具（如 canvas）是用户高频使用的核心工作流之一。
    - **移动端部署**：对Android Termux的支持需求反映了用户希望在更多样化的硬件上运行ZeroClaw的探索趋势。

## 8. 待处理积压

- **等待作者行动的重要 PR**：
    - **`#7821`** (feat(config): add schema struct & risk field)：为安全策略配置引入基础架构，是后续安全功能的关键依赖。但该PR标记为 `needs-author-action`，可能因等待作者补充测试或说明而导致阻塞。
        - **链接**: [zeroclaw-labs/zeroclaw PR #7821](zeroclaw-labs/zeroclaw PR #7821)
    - **`#7098`** (feat(channel/mattermost): add optional WebSocket listener mode)：一个长期挂起的增强请求，为Mattermost频道增加WebSocket监听以减少延迟。从6月2日至今无作者响应（`needs-author-action`），可能面临被关闭的风险。
        - **链接**: [zeroclaw-labs/zeroclaw PR #7098](zeroclaw-labs/zeroclaw PR #7098)
    - **`#7835`** (fix(tools/git_operations): add recovery hint and path context to non-repository error)：一个小而提升体验的修复，同样因 `needs-author-action` 停滞。
        - **链接**: [zeroclaw-labs/zeroclaw PR #7835](zeroclaw-labs/zeroclaw PR #7835)

**维护者提醒**：以上3个PR已等待作者反馈较长时间。若无响应，项目将视其为弃坑并重新评估。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*