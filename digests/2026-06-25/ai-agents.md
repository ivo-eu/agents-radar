# OpenClaw 生态日报 2026-06-25

> Issues: 305 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-06-25 10:25 UTC

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

好的，以下是为您生成的 **OpenClaw 项目动态日报 (2026-06-25)**。

---

## OpenClaw 项目动态日报 | 2026-06-25

### 1. 今日速览

今日项目社区高度活跃，Issue 与 PR 数量巨大，表明项目处于快速迭代与大规模问题反馈期。在过去24小时内，共有 **305条 Issues** 更新，其中 **245条** 处于活跃讨论状态；同时有 **500条 PR** 被提交或更新，其中 **402条** 待合并，但仅有 **98条** 被合并或关闭，**合并效率有待提升**。项目发布了一个新版本 `v2026.6.11-beta.1`，带来了强化的频道控制能力。从社区反馈来看，**消息泄漏 (Message Leak)** 与 **会话稳定性 (Session State)** 是当前最为突出的问题，大量高优先级（P1）的 Bug 和功能请求围绕这两个主题展开。

### 2. 版本发布

**版本：`v2026.6.11-beta.1` (2026-06-11)**
- **链接**: [点此查看](https://github.com/openclaw/openclaw/releases/tag/v2026.6.11-beta.1)
- **亮点**:
    - **更强的频道控制能力**: 新增 Slack 中继模式、Mattermost 原生 `/oc_queue` 支持以及每个私信 (DM) 的独立模型覆盖配置。这使得频道操作的自动化和微调更加灵活。
    - **致谢**: 感谢 `@sjf-oa`, `@amknight`, `@xydigit-zt`, `@thomaszta`, 和 `@gandalf-at-lerian` 等贡献者。
- **迁移注意事项**: 本次为 Beta 版本，建议用户关注新引入的频道配置项，特别是私信模型覆盖与 Slack 中继配置，确保与现有工作流兼容。

### 3. 项目进展

虽然合并/关闭的 PR 数量相对较少（98条），但其中不乏关键修复。在过去24小时内，项目在以下方面取得进展：
- **代理 (Agents) 恢复机制改进**: PR [#95899](https://github.com/openclaw/openclaw/pull/95899) 提出了一种恢复机制，当网关重启后无法恢复代理的“助手回复(Tail)”时，将重新展示用户的最后一条消息，以维持会话连续性。
- **Cron 任务行为修正**: PR [#83933](https://github.com/openclaw/openclaw/pull/83933) 修复了手动执行 `openclaw cron run` 时，意外删除 `deleteAfterRun` 任务并干扰计数器状态的问题。
- **内存 (Memory) 管理优化**: PR [#96465](https://github.com/openclaw/openclaw/pull/96465) 将活跃内存的资源释放与内存管理器关闭操作分离，提升了资源管理的细粒度。
- **代码库健康度**: PR [#96736](https://github.com/openclaw/openclaw/pull/96736) 和 PR [#96711](https://github.com/openclaw/openclaw/pull/96711) 分别增加了终端运行时和临时目录的单元测试，持续提升代码质量。

### 4. 社区热点

过去24小时，社区讨论集中于几个核心痛点，评论热度极高：

- **消息泄漏问题 (Message Leak)**: 这是社区目前最关注的问题，多个高评论数 Issue 与此相关。
    - **Issue #25592 (32条评论)**: 代理在工具调用之间产生的文本（如错误处理、处理确认）被泄漏到 Slack、iMessage 等消息频道中，造成严重的 UX 问题。([链接](https://github.com/openclaw/openclaw/issues/25592))
    - **Issue #44905 (10条评论)**: Discord 频道泄漏了内部 LLM 工具调用痕迹，如 `NO_REPLY`, `commentary` 以及原始工具调用 JSON，安全风险高。([链接](https://github.com/openclaw/openclaw/issues/44905))
    - **Issue #44910 (6条评论)**: OpenAI Codex 相关的错误信息直接泄漏到用户的 Telegram 聊天中，带来了糟糕的使用体验。([链接](https://github.com/openclaw/openclaw/issues/44910))

- **子代理 (Subagent) 任务状态丢失**:
    - **Issue #44925 (20条评论)**: 子代理的完成状态在超时等情况下会被“静默丢失”，既无重试、无通知，也不会自动重启，导致任务不可靠。([链接](https://github.com/openclaw/openclaw/issues/44925))

这些热点反映了用户对**稳定、安全、无泄漏的通信体验**的迫切需求，是目前影响项目口碑的关键障碍。

### 5. Bug 与稳定性

以下根据严重程度列出今日活跃或新报告的 Bug 与回归问题：

| 严重程度 | 问题描述 | 链接 | 状态 |
| :--- | :--- | :--- | :--- |
| **P1** | 代理产生的内部文本泄漏到消息频道 (Slack, iMessage等) | [Link](https://github.com/openclaw/openclaw/issues/25592) | 打开，无Fix PR |
| **P1** | 子代理完成状态在超时后静默丢失 | [Link](https://github.com/openclaw/openclaw/issues/44925) | 打开，有关联PR未合并 |
| **P1** | Discord 频道泄漏内部工具调用痕迹 | [Link](https://github.com/openclaw/openclaw/issues/44905) | 打开，无Fix PR |
| **P1** | 多代理并发时，`agents add` 配置会被覆盖，会话锁失败 | [Link](https://github.com/openclaw/openclaw/issues/43367) | 打开，无Fix PR |
| **P1** (Regression) | 用户自定义的心跳 (Heartbeat) 提示词在更新后不再生效 | [Link](https://github.com/openclaw/openclaw/issues/40255) | 打开，有关联PR未合并 |
| **P2** (Regression) | 内存管理混乱，不同用户/实例的行为表现不一致 (Chunking, 存储) | [Link](https://github.com/openclaw/openclaw/issues/43747) | 打开，无Fix PR |
| **P2** | Telegram 群聊中的图片消息未被正确解析和存储（仅保留占位符） | [Link](https://github.com/openclaw/openclaw/issues/40440) | 打开，无Fix PR |
| **P1** | 与Google Antigravity模型频繁交互导致账号被误判为滥用并封禁 | [Link](https://github.com/openclaw/openclaw/issues/44134) | 打开，无Fix PR |
| **P2** | Docker中Browser路径因Playwright版本更新而改变，缺乏软链接支持 | [Link](https://github.com/openclaw/openclaw/issues/42819) | 打开，有关联PR未合并 |

**小结**: 当前Bug核心集中在 **通信隐私/安全 (Leak)**, **任务可靠性 (Subagent/Timeout)**, 和 **多代理并发稳定性** 三大领域，均为影响用户体验的严重问题。

### 6. 功能请求与路线图信号

以下功能请求反映了社区的显性或潜在需求，结合已有PR，部分功能已进入开发阶段：

- **安全性增强**:
    - **遮罩密钥 (Masked Secrets)**: Issue [#10659](https://github.com/openclaw/openclaw/issues/10659) 要求系统能让代理使用API Key但无法读取原始值，防护注入攻击。**呼声很高** (13条评论, 4个👍)。
    - **配置加密**: Issue [#43794](https://github.com/openclaw/openclaw/issues/43794) 要求对凭证文件进行加密，以应对文件泄露风险。
- **管理性提升**:
    - **代理级成本预算**: Issue [#42475](https://github.com/openclaw/openclaw/issues/42475) 要求在网关上添加每日/每月成本上限，防止财务失控。
    - **预编译Android APK**: Issue [#9443](https://github.com/openclaw/openclaw/issues/9443) 请求提供开箱即用的Android客户端APK，降低使用门槛。**呼声很高** (25条评论, 2个👍)。
- **功能优化**:
    - **会话快照 (Session Snapshot)**: Issue [#13700](https://github.com/openclaw/openclaw/issues/13700) 请求类似游戏存档能力的“保存/加载会话”功能，便于分支测试和回滚。
    - **A2A单向派遣模式**: Issue [#44309](https://github.com/openclaw/openclaw/issues/44309) 建议为智能体间的消息传递增加“只发送、不等待回复”的模式，简化某些自动化场景。

**路线图信号**: 从PR [#88968](https://github.com/openclaw/openclaw/pull/88968) (防止内存刷新失败导致用户回复中断) 和 [#96465](https://github.com/openclaw/openclaw/pull/96465) (内存资源管理分离) 来看，项目核心团队正在积极解决 **会话稳定性** 和 **内存管理** 这两个被社区广泛诟病的问题。

### 7. 用户反馈摘要

从今日的 Issues 评论中，我们可以提炼出以下真实用户反馈：

- **核心痛点**: “Text between tool calls leaks to messaging channels” ([#25592](https://github.com/openclaw/openclaw/issues/25592)). 内部处理逻辑暴露给最终用户是一个**主流且破坏性极强**的UX问题。
- **配置与管理复杂**: “Me and my colleagues (3 people) are using openclaw. I never see any of our memory is managed in same way.” ([#43747](https://github.com/openclaw/openclaw/issues/43747)). 用户反映内存管理行为不统一，造成困惑。
- **稳定性待提升**: “Subagent task orchestration has multiple failure modes where results are silently lost” ([#44925](https://github.com/openclaw/openclaw/issues/44925)). 自动化任务（如子代理）的可靠性是企业级应用的基石，当前静默丢失问题是**严重的信任危机**。
- **部署门槛**: “As a user who wants to run the OpenClaw Android companion app, I would like to have prebuilt APK downloads” ([#9443](https://github.com/openclaw/openclaw/issues/9443)). 用户希望有更便捷的安装方式，而非从源码编译。
- **功能期待**: “Is it possible to have openclaw write and overwrite lines like openai and grok do to indicate the thinking process?” ([#42276](https://github.com/openclaw/openclaw/issues/42276)). 用户羡慕竞品（如OpenAI、Grok）的流式思考过程展示，希望OpenClaw也能支持。

### 8. 待处理积压

以下为长期未更新或未分配，但影响重大的 Issue/PR，提醒维护者关注：

- **高优先级Bug，长期未决**:
    - **Issue #10659** ([Masked Secrets](https://github.com/openclaw/openclaw/issues/10659)): 自2026-02-06起开放，对安全影响巨大（P1），但至今仍`needs-product-decision`。
    - **Issue #9443** ([Prebuilt Android APK](https://github.com/openclaw/openclaw/issues/9443)): 自2026-02-05起开放，用户需求强烈，但同样被`needs-product-decision`卡住。
- **重要PR等待审核**:
    - **PR #88968** ([Prevent memory flush failure from aborting user reply](https://github.com/openclaw/openclaw/pull/88968)): 该PR直接解决会话被中断的关键Bug，审查充分但已等待近一个月，存在5个 `merge-risk` 标签，风险较高，需要谨慎评估。
    - **PR #83933** ([fix(cron): skip deleteAfterRun and preserve counters for manual runs](https://github.com/openclaw/openclaw/pull/83933)): 也是解决核心Cron逻辑错误的PR，处于`needs proof`状态，等待验证。
- **低活跃度高价值PR**:
    - **PR #61519** ([CI: report circular dependencies in PRs](https://github.com/openclaw/openclaw/pull/61519)): 提升代码质量的CI工具，已等待作者响应超过2个月。

---

## 横向生态对比

好的，作为您的资深技术分析师，我已审阅您提供的所有项目动态日报。基于这些详实的数据，以下是为您生成的横向对比分析报告。

---

### **AI 智能体与个人 AI 助手开源生态横向对比分析报告 (2026-06-25)**

**分析师：** 资深技术分析师
**报告日期：** 2026-06-25

---

#### **1. 生态全景**

当前，个人 AI 助手/自主智能体（Agent）开源生态正处于**爆发式增长与快速迭代的“战国时代”**。一方面，项目通过强化频道控制、多代理协作（如A2A）、工具调用（MCP）等核心能力，从“单机玩具”向“生产力基础设施”演进；另一方面，**安全风险**（消息泄漏、权限绕过）和**稳定性问题**（进程泄漏、会话状态丢失）已成为所有主流项目面临的共性挑战。社区对**低延迟、高稳定、强安全**的诉求空前强烈，这标志着生态正从“功能竞赛”转向“质量与信任的攻坚期”。

---

#### **2. 各项目活跃度对比**

| 项目 | 24h活跃 Issues | 24h活跃 PRs | 版本发布 | 核心动态 | 健康度评估 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 245 活跃 (305更新) | 402 待合并 (500更新) | `v2026.6.11-beta.1` | **消息泄漏**问题成为社区第一大痛点，合并效率低。 | **高风险**：社区极度活跃，但核心Bug影响面大，PR积压严重。 |
| **NanoBot** | 20 | 36 | 无 | **安全审计**中心，集中修复MCP和`exec`工具的权限绕过漏洞。 | **中等风险**：安全问题集中爆发，但修复响应迅速有力。 |
| **Hermes Agent** | 18 | 50 (8合并) | 无 | **性能优化**（Token开销）和**互联互通**（A2A协议）是社区热点。 | **健康**：功能迭代与稳定性修复并行，社区讨论深度高。 |
| **PicoClaw** | 少量 | 4 | 无 | **批量关闭安全漏洞**，并修复进化模式Token消耗问题。 | **稳健**：侧重于安全加固和代码质量打磨，生态平稳。 |
| **NanoClaw** | 1 | 19 (4合并) | 无 | 贡献者踊跃，但**审核效率是瓶颈**；安全与多实例支持是亮点。 | **活力-瓶颈**：社区贡献热情高，但维护团队需提速。 |
| **IronClaw** | 22 | 39 | 无 | **高强度迭代**，紧急修复WASM执行导致的**服务冻结**事故。 | **高风险**：生产级事故频发，凸显前沿技术的高风险性。 |
| **LobsterAI** | 少量 (1更新) | 10合并 | 无 | 高效修复**CoWork模式**消息重复和子代理稳定性问题。 | **稳健**：开发节奏高效，持续打磨核心协作功能。 |
| **CoPaw** | 10 | 50 (28合并) | 无 | 修复浏览器资源泄漏，增强自定义提供商支持，合并效率高。 | **健康**：功能与稳定并进，社区与核心团队协作良好。 |
| **ZeroClaw** | 少量 | 50 (12合并) | 无 | 修复**委派绕过白名单**严重Bug，**进程泄漏**问题积压。 | **活跃-隐忧**：安全响应快，但长期Bug积压影响声誉。 |
| **TinyClaw** | 0 | 1 (已关闭) | 无 | 修复Windows原生环境兼容性。 | **冷清**：社区参与度低，但项目健康度良好，无积压。 |
| **NullClaw, Moltis, ZeptoClaw** | 无 | 无 | 无 | 无活动。 | **静默**：可能处于搁置或重构阶段，需关注。 |

---

#### **3. OpenClaw 在生态中的定位**

作为生态的“核心参照”项目，OpenClaw 的核心优势在于**功能全面性**和**社区规模**（24小时内305条Issue、500条PR是其他项目数倍），但这也带来了**巨大的质量维护压力**。

- **与同类相比的优势**：
    - **功能广度**：频道控制、Cron任务、内存管理等生态中最全面的功能集。
    - **社区影响力**：作为参照系，其问题讨论和功能设计往往引领其他项目的开发方向。
- **技术路线差异**：
    - 对比 **NanoBot**（强安全审计）和 **PicoClaw**（嵌入式硬件、轻量），OpenClaw 走的是“**通用平台**”路线，试图覆盖尽可能多的场景。
    - 对比 **Hermes Agent**（侧重性能优化、协议标准化）和 **IronClaw**（前沿技术探索），OpenClaw 更**注重“基础”**，其当前最大问题正是基础能力（消息路由、会话稳定性）的可靠性不足。
- **社区规模对比**：
    - OpenClaw 的Issue/PR数量远超其他项目，但其**PR合并率极低**（98/500 ≈ 20%），而 CoPaw（28/50 ≈ 56%）和 LobsterAI（10/10 ≈ 100%）的合并效率远高于它。这表明OpenClaw可能面临**治理危机**：社区贡献大量涌入，但核心团队的审核能力和方向决策跟不上，导致大量有价值的修复和新功能被“卡住”。

**结论**：OpenClaw 在生态中扮演着“**基础设施平台**”的角色，拥有最庞大的用户和贡献者基础。然而，其当前最大的风险是 **“规模不经济”** ，若不能解决核心稳定性和PR审核效率，其领导地位可能被更敏捷的项目（如NanoBot、Hermes Agent）所侵蚀。

---

#### **4. 共同关注的技术方向**

生态中多个项目同时涌现出以下共性需求，代表着当前行业的技术焦点：

1.  **消息泄漏与隐私安全**：(OpenClaw, NanoBot, PicoClaw)
    - **具体诉求**：防止Agent内部逻辑、工具调用痕迹、错误信息泄漏到用户聊天界面。
    - **项目数量**：3+。这是影响用户体验和信任的**头号杀手**。

2.  **会话稳定性与状态持久化**：(OpenClaw, NanoClaw, IronClaw, ZeroClaw)
    - **具体诉求**：修复子代理状态丢失、会话意外中断、进程泄漏、取消响应不可靠等问题。
    - **项目数量**：4+。Agent任务的**可靠性**是企业级应用的基石。

3.  **多代理协作与互联互通（A2A/ACP）**：(Hermes Agent, ZeroClaw, LobsterAI)
    - **具体诉求**：支持Agent间独立对话、委派、发现与互操作，超越简单的工具调用。
    - **项目数量**：3+。生态从“单Agent”向“多Agent协作系统”演进的明确信号。

4.  **性能优化与Token控制**：(Hermes Agent, PicoClaw)
    - **具体诉求**：减少Token浪费（如Lazy Tool Schema）、优化API调用成本。
    - **项目数量**：2+。尤其在本地模型和付费API场景下，这是**刚性需求**。

5.  **跨平台体验**：(TinyClaw, NanoClaw, CoPaw, Hermes Agent)
    - **具体诉求**：Windows原生支持、移动端（Android APK、PWA）、桌面端UI（系统托盘、时间戳）等。
    - **项目数量**：4+。工具从“CLI神器的”走向“大众消费品”的必经之路。

---

#### **5. 差异化定位分析**

| 项目 | 核心定位 | 目标用户 & 场景 | 技术侧重点 |
| :--- | :--- | :--- | :--- |
| **OpenClaw** | **全能型平台** | 技术发烧友、团队协作、多平台管理 | 功能广度、频道集成、配置灵活性 |
| **NanoBot** | **安全堡垒型** | 安全敏感用户、企业部署、自动化工作流 | 安全审计、权限控制、漏洞响应速度 |
| **Hermes Agent** | **性能先锋型** | 性能优化党、高级开发者、本地模型爱好者 | Token优化、协议标准化（A2A）、数据驱动决策 |
| **PicoClaw** | **嵌入式轻量型** | 嵌入式设备、边缘计算、资源受限场景 | 轻量化、低功耗、快速启动、安全加固 |
| **NanoClaw** | **协作互通型** | 团队协作、多实例管理、端到端加密通信 | Telegram多Bot、Matrix E2EE、远程MCP |
| **IronClaw** | **前沿探索型** | 追求最新技术的开发者、研究机构 | WASM执行、代码生成、系统集成、高性能并发 |
| **LobsterAI** | **协作实用型** | 团队日常工作流、IM平台深度集成 | CoWork模式、OpenClaw兼容、即时消息优化 |
| **CoPaw** | **企业生态型** | 企业用户、复杂内部系统（如小米IoT） | 自定义提供商、浏览器自动化、Tauri桌面端、IoT集成 |
| **ZeroClaw** | **流程自动型** | 任务自动化、Cron调度、多代理委派 | 流程编排、Cron工具、代理委派、进程管理 |
| **TinyClaw** | **极简插件型** | 个人开发者、快速原型、跨平台CLI用户 | 极简设计、跨平台兼容、易上手 |

**关键差异**：
- **安全优先 vs 功能优先**：NanoBot 和 PicoClaw 将安全作为最高优先级，而 OpenClaw 功能虽多但安全漏洞是其最致命短板。
- **通用平台 vs 特定场景**：OpenClaw, Hermes 是通用平台；PicoClaw, ZeroClaw, LobsterAI 则在特定场景（嵌入式、自动化、协作）上更深入。
- **前沿探索 vs 稳定可靠**：IronClaw 在探索WASM等高风险新技术；而 TinyClaw, NanoBot 则更注重基础的稳定可靠。

---

#### **6. 社区热度与成熟度**

- **第一梯队-快速迭代期（极高热度，问题多，风险高）**：**OpenClaw, NanoBot, IronClaw, ZeroClaw**
    - 特点：日活Issue/PR超30，社区贡献活跃，但同时伴随大量安全和稳定性问题。项目处于“野蛮生长”后的“补坑”阶段。开发者应**谨慎选择**，适合愿意承担风险并积极参与贡献的用户。
- **第二梯队-质量巩固期（热度高，迭代稳，健康度好）**：**Hermes Agent, CoPaw, LobsterAI**
    - 特点：日活Issue/PR在10-50之间，有明确的迭代方向（性能、生态、协作），Bug修复效率高，社区反馈良好。适合**寻求稳定生产级工具**的团队和个人。
- **第三梯队-冷启动/低活跃期（热度低，关注少）**：**PicoClaw, NanoClaw, TinyClaw, NullClaw**
    - 特点：社区活跃度低，贡献者少，项目可能处于早期或维护停滞状态。适合**特定场景**的深度用户或想要低风险观察趋势的开发者。

---

#### **7. 值得关注的趋势信号**

1.  **“安全墙”成为生态准入门槛**：NanoBot和PicoClaw的实践表明，未来各项目的核心竞争力将从“能做什么”转变为“如何保证安全和隐私”。**对AI智能体开发者而言**，在构建Agent时，**默认安全（Secure by Default）** 架构（如最小权限原则、沙箱隔离）将成为必须，而非可选项。

2.  **从“工具调用”到“任务编排”**：Hermes Agent的A2A协议、LobsterAI的CoWork模式、ZeroClaw的代理委派，共同指向一个趋势：Agent的价值不再仅仅是调用几个API，而是**成为一个可编程的、协作的、可靠的任务执行网格**。**开发者应关注**事件驱动架构、状态机、以及Agent间通信协议的设计。

3.  **“Token意识”觉醒**：Hermes Agent社区对73%固定开销的量化分析是一个强烈信号。**这意味着AI应用的成本和性能优化将成为独立的技术领域**。开发者需要学会度量、分析和优化Token消耗，例如通过按需加载工具Schema、缓存历史记录等方式。

4.  **跨平台边界消融**：TinyClaw、CoPaw对Windows和移动端的支持，预示着AI助手正从“服务器端守护程序”向“用户个人设备上的原生应用”演进。**未来**，本地优先（Local-first）、边缘智能、PWA等将是重要方向。

5.  **供应链安全初现端倪**：ZeroClaw的SBOM、依赖安全审查提案，预示着一个更成熟的开源生态将关注**软件供应链的透明度和安全性**。对于依赖多个开源项目的开发者而言，这是一项重要的前瞻性考量。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，根据您提供的 NanoBot GitHub 数据，我已生成 2026 年 6 月 25 日的项目动态日报。

---

### NanoBot 项目动态日报 (2026-06-25)

#### 1. 今日速览

今日 NanoBot 项目处于**极高活跃度**状态，核心焦点集中在**安全审计与漏洞修复**上。过去 24 小时内，有 20 个 Issues 和 36 个 PRs 被更新，其中安全相关议题占据主导地位，尤其是关于 MCP 和 `exec` 工具的权限绕过高危漏洞。项目维护者响应迅速，已合并关键的安全修复 PR (#4452, #4436) 以解决部分已报告的严重问题，展现出强健的安全应急响应能力。同时，社区也提交了多项功能增强和稳定性修复，表明项目在快速迭代新功能的同时，正积极地加固安全防线。

#### 2. 版本发布

(无新版本发布)

#### 3. 项目进展

今日项目在**安全加固**和**稳定性提升**方面取得了关键进展。以下为重点合并/关闭的 PR：

- **[安全] MCP 权限控制系统加固 (#4452, #4436, #4434, #4435)**: 针对此前爆出的 `enabledTools` 配置未能限制 MCP 资源和提示词暴露的严重问题（#4434, #4435），项目团队合并了两个修复 PR：
    - PR #4452: 统一将 `enabledTools` 的限制应用至所有 MCP 模型可调用的能力（工具、资源和提示词）。
    - PR #4436: 为 MCP 资源和提示词的注册添加了 `enabledTools` 的检查关卡。
    - **意义**: 这两项修复堵住了严重的安全漏洞，确保了“拒绝一切”策略（`enabledTools: []`）的有效性，是今日项目最重要的安全里程碑。

- **[修复] Telegram Web 兼容性修复 (#4489, #4488)**: 针对 Telegram Web 端因富文本消息渲染问题导致的空白消息 Bug (#4488)，PR #4489 被合并，通过引入 `rich_messages` 配置项，允许用户禁用富文本消息以兼容 Telegram Web。

- **[修复] CLI 向导 `onboard --wizard` 问题修复 (#4438)**: 修复了配置向导中搜索引擎（如 Keenable）不显示及配置字段错误的问题，提升了新用户的上手体验。

- **[CI/CD] 测试套件性能优化 (#4507)**: 合并了提速测试套件的 PR (#4507)，减少了不必要的时间等待，提升了开发及 CI/CD 流程的效率。

- **[WebUI] 代码块复制功能修复 (#4509)**: 修复了 WebUI 中代码块复制在部分环境下失效的问题，恢复了一项关键的用户体验功能。

这些进展表明，项目在快速响应社区反馈的同时，正在系统性地提升产品的安全水平和稳定性。

#### 4. 社区热点

今日社区讨论的核心是 **安全漏洞**，尤其是由安全研究员 `YLChen-007` 报告的一系列安全问题。虽然后续单个 Issue/PR 的评论数不多，但议题数量庞大且影响严重，已成为社区的绝对焦点。

- **[Security] MCP `enabled_tools` 配置白名单绕过 (#4517)**: 这是对之前问题 (#4434) 的进一步深挖，指出即便修复后，MCP 能力注册环节仍可能存在绕过风险。链接: [Issue #4517](https://github.com/HKUDS/nanobot/issues/4517)
- **[Security] `exec` 工具命令执行绕过 (#4521, #4520, #4519, #4518, #4516, #4515, #4514)**: `YLChen-007` 今天连续提交了 7 个关于 `exec.allowPatterns` 配置的绕过报告，涵盖了 Shell 链式执行、注释尾部绕过、包装前缀绕过等多种攻击路径。这反映出 `exec` 工具的权限控制实现存在系统性的缺陷，是当前最受关注和亟待解决的安全挑战。链接: [Issue #4521](https://github.com/HKUDS/nanobot/issues/4521)
- **[Bug] Windows 系统服务重启问题 (#4513)**: 用户 `Quincy-Zh` 报告了在 Windows 下使用 `nssm` 将 NanoBot 设为服务后，调用 `/restart` 指令会导致端口占用或服务状态不一致的问题。这吸引了少量社区讨论。链接: [Issue #4513](https://github.com/HKUDS/nanobot/issues/4513)

社区讨论背后的主要诉求是 **追求更安全的默认配置和更稳健的权限控制体系**，尤其是在高风险的操作（如执行 Shell 命令）上。

#### 5. Bug 与稳定性

今日报告的 Bug 主要集中在**安全缺陷**和**特定环境（Windows、Telegram Web）下的兼容性问题**。

- **严重**
    - **[Security] `exec` 工具多个白名单绕过漏洞**: 包括 Shell 链式执行 (#4514, #4520)、注释尾部绕过 (#4515)、包装前缀绕过 (#4516)、默认登录 Shell 泄露秘钥 (#4518) 等。这些漏洞允许攻击者在受限配置下执行未授权的 Shell 命令，风险极高。**目前无 Fix PR。**
    - **[Security] MCP `enabledTools` 范围绕过 (#4519)**: 资源与提示词包装器仍可能被暴露，风险高。**目前无 Fix PR。**
    - **[Bug] Windows 系统服务 `/restart` 异常 (#4513)**: 在 NSSM 服务环境下重启程序导致端口占用或服务状态错误。**目前无 Fix PR。**

- **中等**
    - **[Bug] 钉钉（DingTalk）通道问题 (#4497)**: HTTP 超时和不支持富文本格式导致消息丢失。已有 Fix PR #4501。
    - **[Bug] Telegram Web 空白消息回复 (#4499, #4488)**: 富文本消息在 Telegram Web 端渲染失败。已合并 Fix PR #4489。
    - **[Bug] Windows `--background` 选项状态不一致 (#4511)**: 程序重启后，记录的进程信息与实际运行状态不符。**目前无 Fix PR。**

- **轻微**
    - **[Bug] WebUI 首页消息发送不跳转 / 自重启后停止按钮失效 (#4500)**: WebUI 的一些交互和状态管理问题。**目前无 Fix PR。**

#### 6. 功能请求与路线图信号

今日收到多个有潜力的功能请求和增强 PR，可能纳入后续版本开发计划：

- **新增 Webhook 触发器 (#4502)**: PR #4502 提议为 Gateway 添加通用 Webhook 触发能力，增强与外部系统的集成，这在自动化工作流中非常有价值。
- **子代理聚合结果模式 (#4414)**: PR #4414 为子代理添加 `aggregated` 结果模式，可缓冲所有子代理结果后再统一推送，更适合需要汇总分析的场景。
- **暂停回合（SuspendTurn）机制 (#4411)**: PR #4411 添加 `SuspendTurn` 能力，允许工具在调用后暂停当前回合，等待异步回调或人工介入，这对于构建“人在回路”（Human-in-the-Loop）应用至关重要。
- **新增 Serper.dev 搜索提供商 (#4406)**: PR #4406 为 `WebSearchTool` 添加了 Serper.dev 支持，丰富了用户的搜索引擎选择。
- **PWA 支持与移动端侧边栏手势 (#4479)**: Issue #4479 提出了为 WebUI 添加 PWA 支持和移动端滑动操作，将显著提升移动端用户体验。
- **`ask_clarification` 工具 (#4508)**: 用户提议新增一个工具，用于在需求不明确或操作风险较高时，向用户提问以澄清意图，这能有效减少 AI 的“幻觉”和误解。

这些功能请求反映了社区对项目**可扩展性、鲁棒性、单/多人协作模式**以及**移动端体验**的更高期待。

#### 7. 用户反馈摘要

- **用户痛点**：
    - **安全配置失效**：尽管有详细文档，但用户发现核心安全配置（如 MCP 和 `exec` 的白名单）存在绕过风险，这严重打击了用户对项目安全性的信心。
    - **平台特性兼容性**：在 Windows 系统服务管理和特定聊天平台（如 Telegram Web、钉钉）的兼容性问题，给用户的实际部署和日常使用带来了困扰。
    - **日志与状态一致性**：例如 `--background` 模式和 `/restart` 后状态不一致，导致运维人员难以判断程序真实运行情况，增加了维护成本。

- **用户诉求**：
    - **更高安全标准**：用户希望安全功能不仅是“白名单”形式，更需要从架构上防止绕过，提供默认安全的配置。
    - **更佳的平台支持**：期望项目对各种部署环境和客户端平台提供更好的测试和适配，减少“水土不服”的情况。

#### 8. 待处理积压

以下为社区关注度较高或已存在一段时间，但仍待维护者响应的关键议题，长期搁置可能影响项目声誉和发展：

- **待处理的关键安全漏洞**: 今日 `YLChen-007` 报告的至少 7 个 `exec` 相关的绕过漏洞（#4514 至 #4521）仍无 Fix PR。这些议题的解决方案对于平台安全至关重要，需优先处理。
- **长期未合并的重要增强 PR**:
    - **Serper.dev 搜索提供商 (#4406)**: 创建于 6 月 18 日，未合并。这个 PR 具有明确的实际应用价值。
    - **子代理聚合结果模式 (#4414)**: 创建于 6 月 19 日，未合并。这是一个有潜力的核心功能增强。
    - **暂停回合（SuspendTurn）机制 (#4411)**: 创建于 6 月 19 日，未合并。对于推进复杂的交互模式至关重要。
- **WebUI 相关 Bug 修复**: Bug #4500 和 PR #4510 若无冲突，建议尽快评估和合并，以避免在主要安全修复之外积累过多的 user-facing 问题。

**总结**：NanoBot 项目今日展现了强大的社区活跃度和快速迭代能力，但安全问题已成为当前最重要的矛盾点。项目组今日已快速修复了 MCP 的关键漏洞，但从如此多的新安全报告来看，对 `exec` 工具和其他模块的安全性需要进行一次系统性的审查和重构。处理好这些安全问题，将是项目迈向更高成熟度的关键一步。

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent 项目动态日报 (2026-06-25)

---

## 1. 今日速览

过去 24 小时，Hermes Agent 项目保持高度活跃：**18 条新/活跃 Issue** 和 **50 条 PR 更新**，其中 8 个 PR 已完成合并/关闭。无新版本发布。社区讨论集中在**性能优化（Token 开销）、跨 Agent 协议（A2A）、桌面端体验**及**各提供商兼容性**上。项目整体健康度良好，但需关注多个 P2 级 Bug 的修复进度。

---

## 2. 版本发布

**无**（过去 24 小时无新版 Release）。

---

## 3. 项目进展

今日合并/关闭了 **8 个 PR**，代表性进展如下：

| PR | 类型 | 摘要 |
|----|------|------|
| [#43637](https://github.com/NousResearch/hermes-agent/pull/43637) (已关闭) | feat(gateway) | 为消息网关（Telegram、Discord 等）添加 `/sessions` 命令处理，此前仅 TUI/桌面端可用。 |
| [#49348](https://github.com/NousResearch/hermes-agent/pull/49348) (已关闭) | feat(tool/skills) | 新增官方可选技能 `hermes-best-practices`，包含 6 大类操作习惯检查清单。 |

其他 6 个已关闭 PR 涉及文档改进、小范围 Bug 修复和配置优化。项目整体在**网关功能完整性**和**用户引导**方面向前迈进一步。

---

## 4. 社区热点

以下 Issues 在过去 24 小时获得最多讨论（按评论数排序），反映社区核心诉求：

| Issue | 标题 | 评论 | 👍 | 核心诉求 |
|-------|------|------|----|----------|
| [#6839](https://github.com/NousResearch/hermes-agent/issues/6839) | Feature: Lazy Tool Schema Loading — Two-Pass Tool Injection to Reduce Token Overhead | 28 | 15 | 减少每次 API 调用中不必要的工具 schema 开销（约 3500–5000 tokens），提升本地模型效率。 |
| [#514](https://github.com/NousResearch/hermes-agent/issues/514) | Feature: A2A (Agent-to-Agent) Protocol Support | 22 | 18 | 支持 Google A2A 协议，实现不同框架 Agent 的发现、通信与互操作，补充 MCP 的“工具级”能力。 |
| [#4379](https://github.com/NousResearch/hermes-agent/issues/4379) | Token overhead analysis: 73% of each API call is fixed overhead (~13.9K tokens) | 16 | 0 | 通过监控面板定量分析 Token 浪费，请求优化基础架构。 |
| [#10567](https://github.com/NousResearch/hermes-agent/issues/10567) | Add --host and CORS config for dashboard to enable Tailscale/VPN access | 12 | 13 | 远程访问 Dashbord 的需求强烈，当前硬编码 127.0.0.1 及 CORS 规则限制使用场景。 |
| [#13834](https://github.com/NousResearch/hermes-agent/issues/13834) | Hermes openai-codex fails on same machine where official Codex CLI works | 12 | 3 | 与 OpenAI Codex CLI 的兼容性问题，影响部分 macOS 用户。 |

**分析**：性能（Token 节省）和互操作性（A2A）是当前社区最关注的两大方向。远程桌面访问与 Dashbord 配置也是高频诉求。

---

## 5. Bug 与稳定性

今日报告的 Bug 及已有 Fix PR 汇总（按严重程度排列）：

| 严重程度 | Issue/PR | 摘要 | 是否有 Fix PR |
|----------|----------|------|---------------|
| **P2** | [#13834](https://github.com/NousResearch/hermes-agent/issues/13834) | `openai-codex` 在相同环境下失败，官方 CLI 正常。 | 无（需进一步复现） |
| **P2** | [#50663](https://github.com/NousResearch/hermes-agent/issues/50663) | z.ai 在“高峰时段”对 Hermes Agent 进行限流，而其他客户端正常。 | 无（外部限制） |
| **P2** | [#52444](https://github.com/NousResearch/hermes-agent/pull/52444) | 图片负载返回 413 时导致会话“崩溃”，无法恢复。 | **已有 Fix PR #52444** |
| **P2** | [#52441](https://github.com/NousResearch/hermes-agent/pull/52441) | CLI TUI 中交互 stderr 未重定向，导致警告信息污染界面。 | **已有 Fix PR #52441** |
| **P2** | [#52432](https://github.com/NousResearch/hermes-agent/pull/52432) | Gateway 运行时 `agent.request_overrides` 被无条件覆盖，丢失自定义 provider 配置。 | **已有 Fix PR #52432** |
| **P2** | [#52431](https://github.com/NousResearch/hermes-agent/pull/52431) | `_summarize_tool_result` 因非字符串参数导致 `TypeError`，引发 TUI 无限崩溃。 | **已有 Fix PR #52431** |
| **P2** | [#51741](https://github.com/NousResearch/hermes-agent/pull/51741) | `/compress` 命令在 TUI 中因未绑定变量而失效（LM Studio 场景）。 | **已有 Fix PR #51741** |
| **P3** | [#52428](https://github.com/NousResearch/hermes-agent/issues/52428) | Windows 下 `browser_navigate` 调用报 `name '_hermes_read_browser_output' is not defined`。 | 无（需复现） |

**注意**：今日许多 P2 级 Bug 已有对应 PR 提交，说明社区响应迅速，但需维护者及时 review 合并。

---

## 6. 功能请求与路线图信号

以下新提出的功能请求结合已有 PR，可推测下一版本可能纳入的方向：

| 功能请求 | 相关 PR/证据 | 说明 |
|----------|--------------|------|
| [Lazy Tool Schema Loading (#6839)](https://github.com/NousResearch/hermes-agent/issues/6839) | 无直接 PR，但 #4379 提供数据支撑 | 高 token 节省潜力，可能进入性能优化主线。 |
| [A2A 协议支持 (#514)](https://github.com/NousResearch/hermes-agent/issues/514) | [#5257](https://github.com/NousResearch/hermes-agent/issues/5257)（通用 ACP 客户端） | 跨 Agent 协作生态，可能与本 PR 联动。 |
| [桌面端系统托盘 (#52434)](https://github.com/NousResearch/hermes-agent/issues/52434) | 无 PR | 常见桌面基础功能，需求明确。 |
| [对话可读性改进 (#52426)](https://github.com/NousResearch/hermes-agent/issues/52426) | 无 PR | UI 易用性改进，低风险。 |
| [Dashbord 远程访问 (#10567)](https://github.com/NousResearch/hermes-agent/issues/10567) | 无 PR | 企业/自托管用户强需求。 |
| [Rocket Chat 网关 (#3725)](https://github.com/NousResearch/hermes-agent/issues/3725) | 无 PR | 扩展消息通道。 |
| [Vertex AI 提供商 (#8427)](https://github.com/NousResearch/hermes-agent/pull/8427) | **已有 PR #8427 (OPEN)** | 企业级 Google Gemini 支持，需合并。 |
| [Ollama Cloud 插件 (#22648)](https://github.com/NousResearch/hermes-agent/pull/22648) | **已有 PR #22648 (OPEN)** | 云端搜索/视觉降级，丰富插件生态。 |
| [Shadow Clone 异步委托+SQLite 持久化 (#48907)](https://github.com/NousResearch/hermes-agent/pull/48907) | **已有 PR #48907 (OPEN)** | 提高后台任务可靠性，应对网关重启。 |
| [Mem0 自定义 API URL (#51654)](https://github.com/NousResearch/hermes-agent/pull/51654) | **已有 PR #51654 (OPEN)** | 自托管记忆层灵活性。 |

**路线图信号**：性能（Token 优化）、企业级提供商（Vertex AI、Ollama Cloud）、任务持久化、桌面体验改进是当前最活跃的开发方向。

---

## 7. 用户反馈摘要

从今日 Issues 评论中提炼真实用户声音：

1. **Token 浪费是最大痛点**（#6839、#4379）：用户构建监控后指出 73% 的 API 调用为固定开销，尤其在本地模型上影响显著。希望引入“懒加载”或“按需注入”工具 schema。

2. **远程桌面访问受阻**（#10567）：多位用户反馈希望在家用 Tailscale 或 VPN 下安全访问 Dashboard，当前绑定 127.0.0.1 和硬编码 CORS 是阻碍。

3. **z.ai 限流体验差**（#50663）：用户使用 Max 套餐仍被限流，竞争对手（opencode、claude）无此问题，怀疑是 Hermes 请求模式触发限制。

4. **桌面 UI 细节**：用户期望区分用户/助手消息（#52426）、显示原始模型 ID（#52442）、系统托盘支持（#52434）、文件拖拽跳过 staging（#52427）。反映桌面端进入精细打磨阶段。

5. **网关命令缺失**：虽然今日合并了 `/sessions` 命令，但仍有用户抱怨多个网关未实现 TUI/桌面已有的 slash 命令。

---

## 8. 待处理积压

以下长期开放的重要 Issue 或 PR 尚未得到维护者足够关注，建议优先 review：

| 项目 | 创建时间 | 最后更新 | 摘要 | 建议动作 |
|------|----------|----------|------|----------|
| [#514](https://github.com/NousResearch/hermes-agent/issues/514) (A2A) | 2026-03-06 | 2026-06-25 | 已获 18 👍，22 评论，但仍无 PR。 | 标记为路线图里程碑并分配责任人。 |
| [#8427](https://github.com/NousResearch/hermes-agent/pull/8427) (Vertex AI) | 2026-04-12 | 2026-06-25 | 功能完整，但长期处于 OPEN 状态。 | 安排 review 并尝试合并。 |
| [#22648](https://github.com/NousResearch/hermes-agent/pull/22648) (Ollama Cloud) | 2026-05-09 | 2026-06-25 | 重写后解决冲突，但未合并。 | 检查 CI 状态，加速合并。 |
| [#5257](https://github.com/NousResearch/hermes-agent/issues/5257) (泛化 ACP 客户端) | 2026-04-05 | 2026-06-25 | 与 #514 互补，16 👍。 | 与 #514 合并考虑。 |
| [#3725](https://github.com/NousResearch/hermes-agent/issues/3725) (Rocket Chat) | 2026-03-29 | 2026-06-25 | 11 👍，无 PR。 | 征集社区贡献或标记为 good first issue。 |
| [#21172](https://github.com/NousResearch/hermes-agent/issues/21172) (Loop Contract) | 2026-05-07 | 2026-06-25 | 较新但设计讨论活跃，建议跟踪。 | 等待更详细设计提案。 |

---

*本日报数据来源：NousResearch/hermes-agent GitHub 仓库，统计截止 2026-06-25 UTC。*

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，根据您提供的 PicoClaw 项目数据，我为您生成了 2026-06-25 的项目动态日报。

---

### PicoClaw 项目动态日报

**日期：** 2026-06-25
**分析师：** AI 智能体与个人 AI 助手领域开源项目分析师

---

#### 1. 今日速览

- **活跃度评估：** **极高**。今日项目在安全修复和代码质量提升方面表现出非常高的活跃度。
- **核心动态：** 批量关闭大量安全漏洞报告，显示维护者对安全的重视程度极高。同时，社区贡献者提交了一系列代码优化和改进性 PR，体现项目生态的健康发展。
- **总体状态：** 项目当前处于一个 **“安全加固 + 稳定演进”** 的阶段。大量已报告的安全问题得到闭环，内部逻辑和兼容性得到优化，为下一阶段的功能开发奠定了基础。

---

#### 2. 版本发布

无新版本发布。

---

#### 3. 项目进展

今日共有 **4 个 PR** 被合并或关闭，主要集中在代码修复和兼容性优化，有效提升了项目的稳定性和安全性。

- **[PR #3169] - fix(evolution): 跳过心跳轮的冷路径调度 (已合并)**
    - **贡献者:** Alix-007
    - **内容:** 修复了 evolution（进化）模块在心跳检查时仍触发冷启动路径的问题，避免了在进化草稿模式下，周期性心跳检查消耗不必要的 tokens。
    - **链接：** [查看 PR](https://github.com/sipeed/picoclaw/pull/3169)

- **[PR #3168] - fix(model): 处理错误响应读取失败 (已合并)**
    - **贡献者:** Alix-007
    - **内容:** 修复了当获取 OpenAI 兼容模型列表时，即使收到非 200 状态码但无法读取错误体时，会返回空或误导性错误信息的问题。现在会正确返回底层读取错误。
    - **链接：** [查看 PR](https://github.com/sipeed/picoclaw/pull/3168)

- **[PR #3166] - fix(openai_compat): 使用结构化日志记录 native_search 警告 (已合并)**
    - **贡献者:** Alix-007
    - **内容:** 修复了 `openai_compat` 模块中一处 `log.Printf` 调用导致构建失败的问题，将其替换为包内统一的结构化日志记录器。
    - **链接：** [查看 PR](https://github.com/sipeed/picoclaw/pull/3166)

- **[PR #3045] - fix(identity): `allow_from` 对带冒号的 Matrix 用户 ID 的降级处理 (已关闭)**
    - **贡献者:** chengzhichao-xydt
    - **内容:** 修复了 `allow_from` 配置无法正确处理 `@alice:example.com` 格式的 Matrix 用户 ID 的问题。原因在于解析逻辑错误地将冒号前的部分作为平台名，导致认证失败。
    - **链接：** [查看 PR](https://github.com/sipeed/picoclaw/pull/3045)

---

#### 4. 社区热点

今日社区讨论热度最高的两个话题反映了用户对 **核心体验** 和 **未来方向** 的关注。

1.  **[Issue #2404] - [需求] 支持在配置中添加流式 HTTP 请求 (已关闭)**
    - **热度：** 13 条评论，1 个 👍
    - **内容：** 用户 `OuSatoru` 强烈建议在配置文件中增加 `streaming` 选项，使 PicoClaw 能像 OpenAI Python 客户端一样向 LLM 后端发送流式 HTTP 请求。这个需求非常明确，旨在减少用户感知的延迟，提升交互实时性。该 issue 虽已关闭，但其讨论热度凸显了社区对 **流式交互** 功能的迫切期待。
    - **链接：** [查看 Issue](https://github.com/sipeed/picoclaw/issues/2404)

2.  **[Issue #3167] - 咨询：PageAgent 对 Vue 等 MVVM 架构的适配方案或规划**
    - **热度：** 用户 `Wavekip` 发起
    - **内容：** 用户在使用 Vue 2 和 Element UI 构建的企业后台系统时，测试了 PageAgent，并提出了一个非常有价值的咨询。他们希望了解 PageAgent 在操作高度动态、依赖 v-model 和组件内部状态的 MVVM 框架（如 Vue）页面时，是否有成熟的适配方案或规划。这反映了项目向 **复杂、动态网页自动化** 领域拓展的重要潜在需求。
    - **链接：** [查看 Issue](https://github.com/sipeed/picoclaw/issues/3167)

---

#### 5. Bug 与稳定性

今日是安全漏洞修复的集中日。共有 **10 个** 安全相关 Issues 在同一天（2026-06-09）被报告，并于今日完成关闭。这些漏洞覆盖面广，严重程度高，但均已得到响应和关闭，表明项目团队有专门的漏洞处理流程。

- **【严重】** **[Issue #3079] - `exec` 命令白名单允许 `jq` 环境信息泄露**
    - **描述：** 通过特定命令可绕过 deny-pattern，利用 `jq` 命令泄露环境变量。
    - **状态：** 已关闭。
    - **链接：** [查看 Issue](https://github.com/sipeed/picoclaw/issues/3079)

- **【严重】** **[Issue #3078] - `web_fetch` SSRF 防护可被环境配置的 HTTP 代理绕过**
    - **描述：** 通过配置系统环境变量中的 HTTP 代理，可以绕过 `web_fetch` 工具对 SSRF 攻击的防护。
    - **状态：** 已关闭。
    - **链接：** [查看 Issue](https://github.com/sipeed/picoclaw/issues/3078)

- **【严重】** **[Issue #3075] - 不受信任的本地仓库 `skills/` 元数据被自动加载到系统提示词中**
    - **描述：** 当从不可信的仓库运行 PicoClaw 时，`skills/` 目录下的内容会被自动加载，存在提示词注入风险。
    - **状态：** 已关闭。
    - **链接：** [查看 Issue](https://github.com/sipeed/picoclaw/issues/3075)

- **【高】** **[Issue #3074] - `web_fetch` SSRF 防护可通过 ISATAP IPv6 地址绕过**
    - **描述：** 利用嵌入回环/私有 IPv4 的 ISATAP IPv6 地址形式，可以绕过对本地/内网 IP 的访问限制。
    - **状态：** 已关闭。
    - **链接：** [查看 Issue](https://github.com/sipeed/picoclaw/issues/3074)

- **【高】** **[Issue #3072] - 启动器首次密码设置存在 CSRF 漏洞，可导致本地控制面接管**
    - **描述：** 在首次设置密码时存在跨站请求伪造漏洞，攻击者可利用本地已认证用户设置密码，从而控制 PicoClaw Launcher。
    - **状态：** 已关闭。
    - **链接：** [查看 Issue](https://github.com/sipeed/picoclaw/issues/3072)

- **【其他已修复的安全问题】** 此外，还包括 Issue #3082 (飞书上下文扩展绕过 `allow_from`)、#3081 (`exec` 工作目录符号链接竞态)、#3076 (企业微信群触发策略绕过)、#3073 (LINE webhook 重放攻击)、#3071 (已认证 WebSocket 客户端可触发 `/reload`) 和 #3068 (MQTT `allow_from` 绕过) 等，均已在今日被标记为已关闭。

- **【一般 Bug】** **[Issue #3012] - [BUG] 启用 evolution 后，每分钟持续消耗 tokens (已关闭，有 PR)**
    - **链接：** [查看 Issue](https://github.com/sipeed/picoclaw/issues/3012)
    - **描述：** 用户反馈当 evolution 功能启用时，即使在非活跃状态，系统每分钟都会消耗 tokens。此问题已在今日的 **[PR #3169]** (修复心跳消耗) 中得到解决。
    - **当前状态：** 已通过合并 PR 修复。

---

#### 6. 功能请求与路线图信号

- **流式传输（Streaming）**：用户 [Issue #2404] 对添加流式 HTTP 请求功能的呼声很高。这是一个核心用户体验改进，可能被纳入 **下一个版本的候选功能**。
- **MVP/MVVM 框架适配**：用户 [Issue #3167] 对 PageAgent 适配 Vue 等现代框架的关注，预示着未来项目功能可能会向更复杂的 **“智能体操作系统”或“自动化浏览器代理”** 方向演进。
- **DeltaChat 网关**：待合并的 **[PR #3063]** 尝试增加 DeltaChat 网关，这表明社区希望扩展 PicoClaw 的通讯渠道，向更私密、去中心化的聊天网络迈进。
- **远程 Pico WebSocket 模式**：待合并的 **[PR #3118]** 允许多个机器/用户连接到同一个 PicoClaw 实例，这是一个关键的协作和扩展功能，很可能成为 **下一个里程碑版本的核心特性**。

---

#### 7. 用户反馈摘要

- **核心痛点：Tokens 消耗敏感**：用户 `xpader` 报告进化模式下的 token 浪费问题（Issue #3012），揭示出用户对调用成本有很高的敏感度。维护者已通过 PR #3169 迅速修复，展示了良好的响应力。
- **核心期望：低延迟交互**：用户 `OuSatoru` 对“流式传输”功能的强烈需求（Issue #2404），体现了用户渴望更流畅、实时的人机交互体验。
- **前沿需求：动态页面支持**：用户 `Wavekip` 关于 PageAgent 对 Vue 等 MVVM 框架支持的询问（Issue #3167），虽然只是一个咨询，但预示着企业级用户开始尝试将 PicoClaw 用于真正的生产环境，他们对工具在新一代前端架构下的兼容性充满期待，但目前看来尚未得到明确答复。

---

#### 8. 待处理积压

尽管今日清理了大量 Issues，但仍有部分重要的 PR 处于“待合并”状态，需要维护者审阅和推进，以避免社区贡献者的积极性受挫。

- **[PR #3118] - 增加远程 Pico WebSocket 模式 (已标记为 stale)**
    - **重要性：** **高**。这是一个被社区期待的功能，是项目向分布式协作演进的关键一步。
    - **链接：** [查看 PR](https://github.com/sipeed/picoclaw/pull/3118)

- **[PR #3115] - 修复通用工具输出的内联数据 URL 媒体提取 (stale)**
    - **重要性：** **高**。该 PR 修复了导致会话历史错乱的Bug，直接影响数据完整性。
    - **链接：** [查看 PR](https://github.com/sipeed/picoclaw/pull/3115)

- **[PR #3063] - 新增 DeltaChat 网关 (stale)**
    - **重要性：** **中**。这是一项新功能，扩展了项目生态，但可能不是当前最优先的事项。
    - **链接：** [查看 PR](https://github.com/sipeed/picoclaw/pull/3063)

- **[PR #3116] - 完成 turn.done 生命周期信号 (stale)**
    - **重要性：** **中**。完善了 Pico 协议的生命周期管理，对于正确性和健壮性很重要。
    - **链接：** [查看 PR](https://github.com/sipeed/picoclaw/pull/3116)

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

好的，作为AI智能体与个人AI助手领域开源项目分析师，以下是基于您提供的数据生成的NanoClaw项目动态日报。

---

### NanoClaw 项目日报 — 2026年6月25日

#### 1. 今日速览

今日 NanoClaw 项目活跃度极高，尤其在 Pull Request (PR) 方面，24小时内收到了19条PR，显示出社区贡献者极大的积极性。然而，合并/关闭的仅有4条，大量PR（15条）处于待合并状态，表明项目维护团队的审核和合并速度可能成为当前瓶颈。Issue 方面仅产生1条新反馈，社区讨论热度相对较低。整体来看，项目处于**快速迭代**阶段，社区贡献踊跃，但维护效率有待提高。

#### 2. 版本发布

无

#### 3. 项目进展

今日合并/关闭的4个重要PR，标志着项目在**健壮性、安全性和多实例支持**方面迈出了关键步伐：

- **核心健壮性与清理机制提升**：[#2830](https://github.com/nanocoai/nanoclaw/pull/2830) (已关闭) 修复了因直接删除NanoClaw目录而未运行卸载程序导致的“僵尸”服务注册问题。该修复提升了项目`setup`的健壮性，确保服务注册能够随着二进制文件的消失而被自动清理。

- **认证架构演进**：[#2855](https://github.com/nanocoai/nanoclaw/pull/2855) (已关闭) 实现了一项重大特性：将Claude订阅(OAuth)设为主要认证方式，同时在订阅失败时自动回退到`ANTHROPIC_API_KEY`作为备用。这为用户提供了更无缝和可靠的API访问体验，是基础设施的重要升级。

- **平台功能完善**：[#2849](https://github.com/nanocoai/nanoclaw/pull/2849) (已关闭) 实现了对单实例支持多Telegram Bot的需求。通过`TELEGRAM_BOT_TOKEN_<SUFFIX>`环境变量，用户可在同一个NanoClaw实例下运行多个具有不同功能的机器人。

- **安全加固**：[#2799](https://github.com/nanocoai/nanoclaw/pull/2799) (已关闭) 修复了`send_file`功能中的一个安全漏洞（CVE-2026-29611），该漏洞可能允许被攻击的Agent读取容器内的任意文件。此修复将文件读取路径限制在`/workspace`内，显著提升了沙箱安全性。

这些合并操作表明，项目正在积极回应社区贡献，并持续优化核心体验与安全基线。

#### 4. 社区热点

- **热点 Issue：Telegram 多 Bot 功能回归**：[#2852](https://github.com/nanocoai/nanoclaw/issues/2852) 是今日唯一新开的Issue，且是**社区热点**。作者对“多Bot实例”功能被移除表示困惑，并指出Claude无法让已有的“instance”支持正常工作。此诉求恰好与今天合并的PR [#2849](https://github.com/nanocoai/nanoclaw/pull/2849) 完美呼应，反映出社区对该功能有强烈且一致的需求。该Issue的关闭可能预示着功能回归，值得作者关注。

- **热点 PR：Matrix 原生 E2EE 适配器**：[#2844](https://github.com/nanocoai/nanoclaw/pull/2844) 提出的是一项重大基础设施重构，用基于`matrix-bot-sdk`的原生适配器替换了`@beeper/chat-adapter-matrix`桥接方案，以支持持久化的端到端加密 (E2EE)。尽管评论较少，但此举涉及对核心通信模块的重写，关系到安全性、性能和协议栈的长期维护，是技术决策上的关键讨论点。

#### 5. Bug 与稳定性

今日报告的Bug和稳定性问题主要围绕**安全性**和**关键路径可靠性**：

- **严重 (安全相关)**:
  - **send_file 路径遍历漏洞**：PR [#2799](https://github.com/nanocoai/nanoclaw/pull/2799) (已于今日合并) 修复了此漏洞，是最高优的安全修复。
  - **ncl Socket 与文件路径控制**：有3个关联的待处理PR，涉及`ncl` socket的客户端/服务器超时与容量限制 ([#2802](https://github.com/nanocoai/nanoclaw/pull/2802))、`folder`参数路径遍历 ([#2800](https://github.com/nanocoai/nanoclaw/pull/2800)) 和`safeParseContent`输入校验 ([#2801](https://github.com/nanocoai/nanoclaw/pull/2815))。**关键**：这些是重要的安全防线，需尽快合并。

- **中等 (关键路径 Bug)**:
  - **macOS CA 证书挂载问题**：PR [#2854](https://github.com/nanocoai/nanoclaw/pull/2854) (待合并) 修复了在macOS Rancher Desktop上，因`TMPDIR`未正确设置导致网关CA证书无法挂载到容器内，引发`Self-signed certificate` API调用错误。这是影响特定平台用户核心功能的问题。
  - **测试框架稳定性**：PR [#2851](https://github.com/nanocoai/nanoclaw/pull/2851) (待合并) 修复了由于轮询循环未正确停止而导致窃取后续测试消息的Bug，直接影响CI/CD的可靠性。
  - **outbound.db 日志恢复**：PR [#2750](https://github.com/nanocoai/nanoclaw/pull/2750) (待合并，已存在较久) 修复了容器被Kill后，`outbound.db`日志文件可能损坏或状态不一致的问题，影响数据持久性和恢复能力。

#### 6. 功能请求与路线图信号

今日社区（通过PR）提出了多个明确的功能扩展请求，其中一些很可能被纳入下一版本：

- **Telegram 多 Bot 实例**：Issue [#2852](https://github.com/nanocoai/nanoclaw/issues/2852) 的强烈诉求，已有修复性PR [#2853](https://github.com/nanocoai/nanoclaw/pull/2853) 提交（注意：此PR与已合并的#2849内容相似，可能为相同功能的不同实现或补丁）。**信号强烈**：此功能已被定为高优。
- **原生 Matrix E2EE 支持**：PR [#2844](https://github.com/nanocoai/nanoclaw/pull/2844) 是一个标志性的基础设施变更，虽然不是直接新功能，但它是支持未来安全可靠Matrix通信的基石。**意义重大**：表明项目在向更强大、更安全的消息基础设施迁移。
- **远程 MCP 服务器支持**：PR [#2847](https://github.com/nanocoai/nanoclaw/pull/2847) (待合并) 请求支持连接到基于HTTP/SSE的远程MCP服务器，将Agent的能力扩展到本地进程之外。**趋势信号**：反映了Agent生态的“去中心化”和“云化”趋势，是项目扩展性的关键一步。
- **Docker-in-Docker 支持**：PR [#2846](https://github.com/nanocoai/nanoclaw/pull/2846) (待合并) 旨在解决在容器内运行Docker命令的问题。**信号**：此需求来自高级用户和CI/CD场景，表明项目正被用于更复杂的部署中。

#### 7. 用户反馈摘要

今日唯一的用户反馈来自Issue [#2852](https://github.com/nanocoai/nanoclaw/issues/2852)：
- **用户痛点**：**功能回溯与文档未更新**。用户Kwisss表达了对“telegram multi-bot”功能被移除的沮丧，并指出文档宣称支持“instance”，但实际上（至少通过Claude）无法正常工作。用户明确询问此功能是否会回归，否则将考虑其他方案。这表明功能的一致性和文档的准确性对用户信心至关重要。

#### 8. 待处理积压

以下 Issue/PR 存在时间较长且重要，提醒维护者重点关注：

- **PR #2750 (`outbound.db` 日志恢复)**：自2026-06-12起待合并。这是一个影响数据可靠性的关键Bug修复，积压过久有损系统健壮性。
- **PR #2800, #2801, #2802 (安全修复系列)**：三个紧密相关的安全加固PR，自2026-06-17起待合并。它们是防御路径遍历、socket滥用和输入校验的防线，优先级应很高。
- **PR #2815 (`safeParseContent` 守卫)**：自2026-06-18起待合并，是上述#2801的替代方案，同样提升了输入安全性。

**总结**：项目社区动力强劲，贡献井喷，但**审核合并效率是当前最大挑战**。请维护团队优先处理待合并的安全及关键稳定性PR，并以清晰的沟通回应社区在Issue [#2852](https://github.com/nanocoai/nanoclaw/issues/2852) 中的关切，以维持社区信任。

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域的开源项目分析师，这是根据您提供的 IronClaw 项目数据生成的 2026-06-25 项目动态日报。

---

### **IronClaw 项目动态日报 — 2026-06-25**

#### **1. 今日速览**

过去 24 小时，IronClaw 项目展现出极高的开发活跃度。社区和核心团队围绕 **“Reborn”** 重构版本，在稳定性、并发控制和用户体验方面进行了大量修复和优化。今日共有 22 个 Issue 更新和 39 个 PR 更新，其中包含多项针对生产环境问题的紧急修复（如 WASM 执行导致的 Tokio 工作池饥饿、Slack Webhook 死锁等）。此外，一系列关于 `activity identity` 和 `gate lifecycle` 的架构重构 PR 被合并，标志着项目在核心数据模型一致性上迈出了重要一步。当前项目整体状态为 **“高强度迭代，聚焦稳定性与架构加固”**。

---

#### **3. 项目进展**

今日合并/关闭了多项关键 PR，显著提升了项目的稳定性和架构健壮性：

- **核心架构重构里程碑**：
    - **[PR #5145] [CLOSED]** `refactor(reborn): clean up capability activity lifecycle`：这是一个 XL 规模的重构，旨在统一 Capability Activity 的生命周期管理，确保 activity 行、gate 提示、终端状态和 WebUI 渲染都使用同一个稳定的身份标识（`CapabilityActivityId`），而非从 UI 顺序或运行状态推断。这解决了深层次的数据一致性问题，为防止活动身份分裂或丢失奠定了基础。

- **关键 Bug 修复**：
    - **[PR #5225] [CLOSED]** `fix(reborn): release Slack admission permit once inbound is durably accepted`：修复了 Slack Webhook 路径中的一个严重死锁问题。该 PR 调整了并发控制准入许可（admission permit）的释放时机，确保一旦入站消息被持久化接受，许可即被释放，避免了在处理缓慢时阻塞后续入站请求。
    - **[PR #5202] [CLOSED]** `[codex] fix recurring trigger poller hang`：修复了重复触发轮询器（trigger poller）的挂起问题，确保来自 Slack 等渠道的触发事件在被接受后能立即持久化，不会因后续处理慢而阻塞轮询本身。

- **持续集成与工具链**：
    - **[PR #5206] [OPEN]** `fix(reborn): stop WASM execution from starving the tokio worker pool`：这是一个针对 **2026-06-24 生产事故**的紧急修复 PR，该事故导致服务**完全冻结约 4 分钟**。原因是 WASM 工具执行占用了 Tokio 异步工作池，导致其他所有任务（心跳、日志等）无法运行。此 PR 将 WASM 执行卸载到阻塞线程池，是项目稳定性修复的**重中之重**。
    - **[PR #5224] [OPEN]** `fix(reborn): reconcile scheduler concurrency default and guard worker_count=1`：经深入调查，排除了对调度器并发问题的误判，最终确认问题源于配置引导中的一个竞争条件（race condition），并进行了修复。

**结论**：项目在前沿架构（代码生成、WASM 执行）的生产化过程中，正在经历关键的稳定性提升阶段。今日的进展表明团队正全力解决由高并发和高负载暴露出的深层系统性问题。

---

#### **4. 社区热点**

今日讨论最活跃的 Issue 与 PR 主要集中在新功能的用户体验和架构一致性上。

- **Issue [#5196] “Ask each time” tool permission failure**: [🔗](https://github.com/nearai/ironclaw/issues/5196)
    - **热度**: 1 条评论
    - **诉求分析**：用户报告了一个影响核心工作流的 Bug：配置为“每次询问”的工具在用户批准后，仍会返回授权错误，并触发重复的批准流程。这表明批准状态与执行授权之间存在未同步的状态或竞态条件。这直接影响了用户对代理工具的信任和操作流畅性，诉求是**确保“批准”操作具有确定性和一次性**。

- **PR [#5145] Activity Lifecycle 重构**： [🔗](https://github.com/nearai/ironclaw/pull/5145)
    - **热点分析**: 虽然已于昨日合并，但基于此产生了 **5 个** follow-up Issues (#5200, #5216, #5219, #5218, #5215)。这表明该 PR 是当前开发路线的核心，社区和核心贡献者都在密切关注其引入的新概念（如 `CapabilityActivityId`）能否彻底解决过去活动中普遍存在的身份模糊问题。后续的 Issues 详细探讨了如何强化其不变量、处理过期 Gate 状态以及优化 WASM 资源计费等，显示了社区对架构严谨性的高要求。

---

#### **5. Bug 与稳定性**

今日报告的 Bug 数量较多，按严重程度排列如下：

| 严重程度 | Issue/PR 链接 | 描述 | 修复状态 |
| :--- | :--- | :--- | :--- |
| **P0 (崩溃/服务不可用)** | **[PR #5206]** | **WASM 执行导致 Tokio 工作池饥饿**，使服务完全挂起数分钟。 | 已有 **[Fix PR #5206]** |
| **P1 (核心功能异常)** | **[Issue #5231]** | **自动化设置失败**，显示 `HostUnavailable`。 | 待修复 |
| **P1 (核心功能异常)** | **[Issue #5196]** | “每次询问”工具权限批准后，仍报授权错误并触发重复流程。 | 待修复 |
| **P1 (核心功能异常)** | **[Issue #5227]** | **运行失败消息可能附加到错误的对话轮次**，造成对话时间线混乱。 | 待修复 |
| **P1 (核心功能异常)** | **[Issue #5209]** | **取消响应不可靠**，取消后发送新消息，代理仍在处理旧消息。 | 待修复 |
| **P2 (严重缺陷)** | **[Issue #5229]** | 多租户环境下，Capability 预览使用错误的作用域，导致日志中 `unknown thread` 错误。 | 已有 **[Fix PR #5230]** |
| **P2 (严重缺陷)** | **[Issue #5189]** | 成功的工具调用在运行时**不显示活动详情面板**，而失败的调用可以。 | 待修复 |

---

#### **6. 功能请求与路线图信号**

今日的讨论和 PR 揭示了项目未来的发展方向：

- **稳定性与基础设施强化**：多个新开的 PR（如 [#5232] 持久化 Runner 租约、[#5233] 持久化触发器访问、[#5228] 解耦心跳超时）均指向一个核心目标：**为 Reborn 版本构建更健壮、可恢复的运行时基础设施**。这包括将状态持久化与执行分离、改进文件系统抽象等，是应对生产环境的必要举措。
- **LLM 提供商与 WebUI 体验优化**：
    - **[PR #5214]** 提议强制要求所有 LLM 提供商使用**加强的 HTTP 客户端**（连接超时、keepalive 等），避免因个别提供商问题拖垮整个系统。
    - **[PR #5213]** 指出当前模型调用超时栈无法服务于大规模上下文（1M token）或大规模输出，暗示了处理更大规模任务的路线图需求。
    - **[PR #5212]** 和 **[PR #5226]** 请求并实现在对话中**一致显示消息时间戳**，这是提升 WebUI 用户体验的细节改进。

---

#### **7. 用户反馈摘要**

从今日的 Issues 中可以提取到一些核心用户痛点：

- **自动化体验不佳**：`Issue #5231` 的报告中，用户尝试设置自动化监控 GitHub 仓库，但遭遇了 `HostUnavailable` 错误，表明自动化功能在新环境中不够稳定。
- **“禁止”操作语义不统一**：`Issue #5120` (已关闭) 和 `Issue #5196` 指出，在授权、批准和活动投影等场景中，“拒绝”/“取消”等概念的术语和语义不一致，导致用户和 UI 层理解混乱。这表明用户对操作结果的**确定性和可理解性**有较高要求。
- **时间戳显示不一致**：`Issue #5212` 用户反映消息时间戳“在生成时短暂出现，完成后消失”，这是一个典型的 UI/UX 体验不稳定的例子。
- **对架构改进的积极反馈**：从多个 follow-up Issue (如 #5200, #5216) 可以看出，社区贡献者针对 PR #5145 的 `activity identity` 重构提出了非常具体和深入的优化建议，显示了对项目核心数据模型一致性的**高度认可和积极参与**。

---

#### **8. 待处理积压**

- **[Issue #4108] Nightly E2E failed**: [🔗](https://github.com/nearai/ironclaw/issues/4108)
    - **状态**: 自 2026-05-27 起持续 Open。
    - **分析**：这是一个持续了近一个月的 Nightly E2E 测试失败报告，由 GitHub Actions 机器人自动创建。虽然更新日期显示在 2026-06-25 有更新，但该 Issue 已存在近一个月，且长期没有关闭。这可能表明存在一个间歇性或难以复现的端到端测试问题，需要维护者关注并评估其对发布流程的影响。

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

好的，这是为您生成的 LobsterAI 项目动态日报。

---

# LobsterAI 项目动态日报 — 2026-06-25

## 1. 今日速览

今日项目活跃度非常高，核心团队在一天内合并了 10 个 Pull Request，展现出了高效的开发节奏。修复重点集中在 **CoWork 协作模式**的稳定性与 **OpenClaw 扩展框架**的兼容性上，同时解决了多个导致消息重复、子代理轮询中断和插件加载失败的 Bug。社区反馈方面，有一项关于“定时任务开关失效”的旧 Issue 被再次更新，表明该问题仍对用户造成困扰。整体来看，项目正稳步修复遗留问题并巩固新特性，代码库健康状况良好。

## 2. 版本发布

今日无新版本发布。

## 3. 项目进展

今日共合并/关闭了 10 个 PR，主要推进了以下工作：

-   **OpenClaw 扩展与健全性**
    -   **`#2203`** 修复了本地扩展入口文件的预编译问题，确保打包后的扩展能被正确加载。
    -   **`#2202`** 修复了在严格的插件白名单模式下，浏览器控制插件被错误禁用的问题。
    -   **`#2201`** 修复了子会话（subagent）产出后，最终同步时导致助手回复重复的问题。
    -   **`#2198`** 实现了 QQ 和 Discord 频道插件的预安装，简化了 IM 平台的集成流程。

-   **CoWork 协作模式与设置**
    -   **`#2206`** 修复了“开机自启”设置项状态与操作系统实际状态不同步的问题，并增加了诊断日志和错误提示。
    -   **`#2205`** 更新了计划模式图标，使其能更好地适配浅色/深色主题。
    -   **`#2204`** 修复了 CoWork 计划模式中，未能正确解析块级标签导致消息泄漏的问题。
    -   **`#2200`** 解决了因流式响应抖动导致计划模式消息重复的问题。
    -   **`#2199`** 修复了父会话完成后，子代理仍在运行时轮询过早停止的问题，并设置了 5 分钟的超时保护。
    -   **`#2197`** 修复了在 OpenClaw 回退到 `chat.history` 获取最终回答时，导致摘要重复的问题。

**小结**：项目在修复 CoWork 模式的消息一致性和子任务管理问题上取得了显著进展，同时增强了 OpenClaw 框架的健壮性和插件管理能力。

## 4. 社区热点

今日没有新开的热点讨论。不过，编号为 **`#1392`** 的 Issue 获得了新的更新，值得关注。

-   **`#1392` [OPEN] 定时任务开关点击无反应** - [链接](netease-youdao/LobsterAI Issue #1392)
    -   **状态**：该 Issue 于 2026-04-03 创建，今日（2026-06-25）被更新（内容为图片显示，可能为复现截图）。评论数量为 1。
    -   **分析**：尽管该问题已有近三个月历史，但今日的更新表明用户仍受此困扰。问题描述为部分任务的开关无法通过点击关闭，这直接影响了核心的自动化功能。鉴于 PR 列表中未见直接相关的修复，该项目可能是一个较难复现的偶发 Bug，或者是特定用户环境的兼容性问题。社区的沉默也暗示了问题复现门槛较高。

## 5. Bug 与稳定性

今日未报告新的致命性崩溃或回归问题。修复工作主要围绕以下已知或新发现的稳定性问题：

-   **严重程度：中**
    -   **`#1392` [OPEN] 定时任务开关点击无反应**：这是一个影响用户对自动化任务控制的 Bug。虽然今日无直接修复 PR，但其被再次更新，表明问题依然存在。严重程度中等，因为它影响功能但未导致崩溃。
-   **严重程度：低 (今日已修复)**
    -   **消息重复/泄漏**：一系列 PR (`#2201`, `#2200`, `#2204`, `#2197`) 集中修复了 CoWork 模式下因各种边界情况（子会话同步、流抖动、标签解析、历史回退）导致的助手消息重复或泄漏问题。
    -   **启动项同步**：`#2206` 修复了“开机自启”设置状态不同步的问题。
    -   **插件/扩展加载**：`#2202` 和 `#2203` 修复了插件因白名单或打包问题无法加载的 Bug。

## 6. 功能请求与路线图信号

今日暂无新的功能请求 Issue。从今日合并的 PR 中，我们可以解读出项目的发展方向：

-   **优化 UI 体验**：`#2205` 将计划模式图标替换为支持主题的 SVG 组件，暗示了下一波 UI/UX 打磨工作的开始。
-   **扩展生态完善**：`#2203` 和 `#2198` 均在改进 OpenClaw 扩展系统的易用性和兼容性（预装插件、本地扩展加载），这预示着项目正在积极构建和推广其第三方扩展生态。
-   **强化流式响应稳定性**：`#2200` 和 `#2197` 针对流式传输中的抖动和回退机制进行了精细化处理，表明项目团队正在将流式体验的“终极稳定性”视为下一个重要里程碑。

## 7. 用户反馈摘要

-   **痛点**：
    -   **定时任务控制问题**：Issue `#1392` 的用户尝试关闭定时任务开关，但操作无响应。这可能导致用户无法正常停掉不想继续的任务，是一个明确且令人沮丧的交互障碍。
-   **近期改进**：
    -   虽然用户未直接评论，但 `#2206` 修复的“开机自启”同步问题是一个常见痛点。Windows 用户可能曾遇到过“设置里关了自启，但下次开机软件还是自己开了”的困惑，此修复将直接解决此类体验问题。

## 8. 待处理积压

-   **`#1392` [OPEN] 定时任务开关点击无反应** - [链接](netease-youdao/LobsterAI Issue #1392)
    -   **创建日期**：2026-04-03
    -   **状态**：已标记为 `[stale]`，但今日有更新。
    -   **分析**：该 Issue 已存在近三个月，虽然标记为陈旧，但用户的再次更新使其回到活跃状态。这是一个值得维护团队重视的信号。建议维护者重新评估此问题，尝试与用户沟通以获取更多复现信息，或者进行内部复现测试。如果确实难以复现，也应在 Issue 中同步排查进展，以免让用户感到其反馈被忽视。

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

# TinyClaw 项目日报 | 2026-06-25

---

## 1. 今日速览

- 项目在过去24小时内无新增 Issue，关闭1个PR（#281），无新版本发布，整体活跃度较低。
- 该PR专注于修复Windows原生环境（非WSL）下的CLI兼容性问题，表明项目正在稳步提升跨平台用户体验。
- 无待合并的PR，积压问题清零，项目健康度良好，但社区讨论明显偏冷，需警惕长期低活跃度。

---

## 2. 版本发布

**无新版本发布。**

---

## 3. 项目进展

- **PR #281（已合并/关闭）** — 修复Windows跨平台支持  
  - **作者**：mperkins0155  
  - **时间线**：2026-06-16 创建 → 2026-06-24 更新并关闭  
  - **内容**：修复了三个阻止`tinyagi` CLI在原生Windows上运行的Bug：  
    1. `new URL('.', import.meta.url).pathname` 在Windows上返回 `/C:/Users/...`，导致`path.resolve`将路径解析为 `C:...`（双驱动器号），引发 `MODULE_NOT_FOUND`。  
    2. 其他两个Bug未在摘要中详细说明，但均属于Windows特有路径处理与进程行为差异。  
  - **影响**：此PR移除了WSL依赖，使得Windows用户可以无需额外配置直接运行CLI，显著扩大潜在用户群。

**小结**：项目在跨平台兼容性上迈出实质性一步，向“零配置即用”目标靠近。

---

## 4. 社区热点

今日无活跃讨论或高互动Issue/PR。唯一被关闭的PR #281 未产生评论或点赞（评论: undefined, 👍: 0），社区参与度极低。这表明：
- 项目目前仍处于早期或小众阶段，外部贡献者及用户反馈匮乏。
- 需考虑通过论坛、Discord或文档改进引导用户反馈；该PR的修复本身可能已经解决了大量用户的隐性问题，但缺乏公开讨论记录。

---

## 5. Bug 与稳定性

今日**无新增Bug报告**。但PR #281 所修复的三个Windows Bug可视为近期影响稳定性的重要问题：

| 严重程度 | Bug描述 | 状态 |
|----------|---------|------|
| **高** | 原生Windows下`pathname`返回带斜杠的驱动器号，导致模块加载失败（`MODULE_NOT_FOUND`） | 已修复（PR #281） |
| **高** | 其他Windows特有的CLI崩溃行为（未在摘要展开） | 已修复 |

**结论**：当前无已知公开Bug积压，稳定性较好。

---

## 6. 功能请求与路线图信号

今日无新功能请求（Issue 0条）。但PR #281 的修复内容揭示了用户（尤其是Windows用户）的隐性需求：**原生Windows环境的一流支持**。  
结合项目名称“TinyClaw”及AI Agent定位，可以推测下一版本可能的改进方向：
- 进一步完善Windows下的路径规范化、子进程符号链接处理。
- 增加对PowerShell/cmd的彩蛋或默认配置适配。
- 考虑支持跨平台自动检测（如通过`os.platform()`动态选择路径处理逻辑）。

目前暂无明确路线图信号，建议维护者将Windows支持纳入下一个小版本（如v0.2.x）的更新目标。

---

## 7. 用户反馈摘要

今日无用户反馈（Issues评论为空）。但从PR #281 的摘要可以间接获得用户痛点：
> “`tinyagi` CLI无法在原生Windows（非WSL）上运行”

这意味着：
- **使用场景**：Windows开发者/用户期望直接使用CLI而无需安装WSL依赖。
- **痛点**：路径处理Bug导致工具完全不可用，造成配置门槛过高。
- **满意点**：该PR已关闭，说明问题已得到解决；但缺乏后续用户确认（如测试报告）。

建议维护者发布一个小版本或测试标签，并主动在相关论坛（如r/programming、Reddit、Hacker News）征集Windows用户试用反馈。

---

## 8. 待处理积压

**无**。所有打开的PR和Issues已清零（Issues 0，PR 0）。  
项目当前处于“干净”状态，但长期无积压也可能意味着用户参与度不足。建议维护者：
- 主动提出`help wanted`标签的Issue（如“增加单元测试”、“文档翻译”）。
- 检查是否有已经关闭但未完成回归测试的功能。
- 考虑发布一个包含Windows修复的正式版本（如v0.1.1），并撰写更新日志以吸引用户更新。

---

**报告生成时间**：2026-06-25 12:00 UTC  
**数据来源**：[TinyAGI/tinyagi](https://github.com/TinyAGI/tinyagi)

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

## CoPaw 项目动态日报 — 2026-06-25

### 1. 今日速览

项目保持高活跃度：24 小时内共处理 **14 条 Issue**（新开/活跃 10 条，关闭 4 条）和 **50 条 PR**（新增 22 条待合并，28 条已合并/关闭）。社区提交集中在 **浏览器自动化资源泄漏修复**、**自定义提供商 Function Calling 支持**、**Runtime 2.0 工具注册** 以及 **前端性能优化** 等方向。无新版本发布，但多项重要修复已合入主干，项目稳定性持续改善。

---

### 3. 项目进展

今日合并/关闭了 **28 条 PR**，以下为关键推进项：

| PR | 标题 | 类型 | 影响 |
|----|------|------|------|
| [#5532](https://github.com/agentscope-ai/QwenPaw/pull/5532) | fix(skillpool): restore SkillPool styles | Bugfix | 修复 PR #5521 引起的技能池样式丢失 |
| [#5522](https://github.com/agentscope-ai/QwenPaw/pull/5522) | feat(inbox): add source type filter | Feature | 收件箱新增 Cron/Heartbeat/Memory 来源筛选器，提升消息管理体验 |
| [#5059](https://github.com/agentscope-ai/QwenPaw/pull/5059) | fix(matrix): use nio client.download for encrypted media | Bugfix | 修复 Matrix 端到端加密房间中媒体下载失败问题 |
| [#5485](https://github.com/agentscope-ai/QwenPaw/pull/5485) | fix: refine mcp tool name to parse openai api | Bugfix | 修复 MCP 工具名解析错误 |
| [#5486](https://github.com/agentscope-ai/QwenPaw/pull/5486) | fix: fix tool input json decode | Bugfix | 修复工具输入 JSON 解码异常 |

此外，社区贡献者首次提交的 PR 如 [#5533](https://github.com/agentscope-ai/QwenPaw/pull/5533)（避免内容安全错误误判为媒体能力）和 [#5524](https://github.com/agentscope-ai/QwenPaw/pull/5524)（在 Runtime 2.0 注册 spawn_subagent）均已完成或正处于积极审核中，体现了社区对关键 bug 的快速响应。

---

### 4. 社区热点

当日讨论最活跃的 Issues 及背后诉求：

- **[#5345](https://github.com/agentscope-ai/QwenPaw/issues/5345) [CLOSED]** — 自定义 OpenAI 兼容提供商（如 OMLX）不支持 Function Calling，用户期望与 Ollama 一致。虽已关闭，但 8 条评论反映该功能对使用非标准提供商用户的刚需。
- **[#5505](https://github.com/agentscope-ai/QwenPaw/issues/5505) [OPEN]** — MiniMax-M3 图片安全审核错误被缓存为 `rejects_media=True`，导致后续视觉请求被剥离。3 条评论，已有对应 PR [#5533](https://github.com/agentscope-ai/QwenPaw/pull/5533) 进行修复，社区对模型能力误判机制高度关注。
- **[#5480](https://github.com/agentscope-ai/QwenPaw/issues/5480) [OPEN]** — Console 长消息排版错乱，需切换选项卡恢复。4 条评论暴露了前端 CSS 布局重计算缺失问题，对日常使用体验影响较大。

---

### 5. Bug 与稳定性

当日报告的 Bug 按严重程度排列如下：

| 严重度 | Issue | 摘要 | 当前状态 | 修复 PR |
|--------|-------|------|----------|---------|
| 🔴 严重 | [#5162](https://github.com/agentscope-ai/QwenPaw/issues/5162) | 对话思考逻辑陷入死循环 | OPEN，无关联 PR | — |
| 🔴 严重 | [#5505](https://github.com/agentscope-ai/QwenPaw/issues/5505) | MiniMax-M3 安全审核错误导致的视觉功能永久失效 | OPEN | [#5533](https://github.com/agentscope-ai/QwenPaw/pull/5533) |
| 🟠 高 | [#5520](https://github.com/agentscope-ai/QwenPaw/issues/5520) | `browser_use stop()` 残留 Chrome 渲染进程，导致内存泄漏（#2733 回退） | OPEN | 暂无 |
| 🟠 高 | [#5479](https://github.com/agentscope-ai/QwenPaw/issues/5479) | 大会话文件（>500KB）打开时前端崩溃 | OPEN | 暂无 |
| 🟠 高 | [#5480](https://github.com/agentscope-ai/QwenPaw/issues/5480) | Console 长消息排版错乱 | OPEN | 暂无 |
| 🟡 中 | [#5528](https://github.com/agentscope-ai/QwenPaw/issues/5528) | Linux 桌面 IME 包装的默认浏览器导致 browser tool 启动超时 | OPEN | [#5526](https://github.com/agentscope-ai/QwenPaw/pull/5526) |
| 🟡 中 | [#5523](https://github.com/agentscope-ai/QwenPaw/issues/5523) | `spawn_subagent` 在 Runtime 2.0 中未注册 | OPEN | [#5524](https://github.com/agentscope-ai/QwenPaw/pull/5524) |
| 🟢 低 | [#5508](https://github.com/agentscope-ai/QwenPaw/issues/5508) | Windows 本地版文件预览链接返回 404 | CLOSED | 已修复（未提供 PR 号） |

---

### 6. 功能请求与路线图信号

- **动态模型切换** — [#5527](https://github.com/agentscope-ai/QwenPaw/issues/5527) 提出当模型限流或不可用时自动切换到备用模型，防止任务中断。社区呼声较高，目前无对应 PR，但可能与 [#5399](https://github.com/agentscope-ai/QwenPaw/pull/5399)（自定义模型排序）形成组合能力。
- **Slack 通道集成** — [#5193](https://github.com/agentscope-ai/QwenPaw/pull/5193) 已提交完整的多模态 Slack Channel（Socket Mode），支持流式输出，是渠道拓展的重要信号。
- **小米智能家居集成** — [#5375](https://github.com/agentscope-ai/QwenPaw/pull/5375) 增加了 Miloco（小米智能家居）集成指南和工作区模板，朝向 IoT 方向延伸。
- **Tauri 自动更新** — [#4669](https://github.com/agentscope-ai/QwenPaw/pull/4669) 和 [#4041](https://github.com/agentscope-ai/QwenPaw/pull/4041) 持续推进桌面版体验，包括系统托盘行为和自动更新，表明桌面端正在走向成熟。

---

### 7. 用户反馈摘要

从当日 Issues 评论中提炼的典型用户声音：

- **“OMLX 在 Reasonix 上 Agent 能力正常，但在 QwenPaw 中手动添加后模型只返回文本，不会调用工具。”**（[#5345](https://github.com/agentscope-ai/QwenPaw/issues/5345)）—— 自定义提供商 Function Calling 支持缺失让用户被迫切换工具。
- **“每次调用 `send_file_to_user` 后收到的文件链接不可用，无法预览或下载”**（[#5508](https://github.com/agentscope-ai/QwenPaw/issues/5508)）—— Windows 用户基础功能受阻，虽已关闭但影响面大。
- **“即使只有一次图片请求被拒绝，后续所有图片请求都被剥离，模型实际没看到图片”**（[#5505](https://github.com/agentscope-ai/QwenPaw/issues/5505)）—— 模型能力缓存逻辑过于激进，导致用户误解模型无法处理图片。
- **“死循环问题让我不得不重启整个进程，希望有进度终止机制”**（[#5162](https://github.com/agentscope-ai/QwenPaw/issues/5162)）—— 对话逻辑缺乏边界条件，用户期待预设最大循环次数。
- **“大会话文件打开直接崩溃，我只能删除会话继续使用”**（[#5479](https://github.com/agentscope-ai/QwenPaw/issues/5479)）—— 前端渲染性能瓶颈影响重度用户。

---

### 8. 待处理积压

以下为长期开放且未获得充分关注的 Issue/PR，建议维护者优先评估：

| 类型 | 编号 | 标题 | 创建时间 | 最后更新 | 风险 |
|------|------|------|----------|----------|------|
| Issue | [#5162](https://github.com/agentscope-ai/QwenPaw/issues/5162) | 对话思考逻辑进入死循环 | 2026-06-12 | 2026-06-25 | 严重影响用户体验，无对应修复 PR |
| Issue | [#5479](https://github.com/agentscope-ai/QwenPaw/issues/5479) | 大会话文件（>500KB）打开报错 | 2026-06-24 | 2026-06-25 | 前端性能瓶颈，影响大会话用户 |
| Issue | [#5480](https://github.com/agentscope-ai/QwenPaw/issues/5480) | Console 长消息排版错乱 | 2026-06-24 | 2026-06-25 | 排版缺陷降低专业感 |
| PR | [#4041](https://github.com/agentscope-ai/QwenPaw/pull/4041) | feat(desktop): add Tauri tray behavior | 2026-05-05 | 2026-06-25 | 长时间搁置，与 #4669 可能冲突 |
| PR | [#4669](https://github.com/agentscope-ai/QwenPaw/pull/4669) | feat(desktop): add tauri auto updater | 2026-05-25 | 2026-06-25 | Tauri 桌面更新能力，需合并决策 |

---

**备注**：所有链接基于 `agentscope-ai/QwenPaw` 仓库，项目名称在文档中统一为 **CoPaw**。

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，根据 ZeroClaw 项目在 2026-06-25 的 GitHub 活动数据，为您呈现以下项目动态日报。

---

# ZeroClaw 项目动态日报 | 2026-06-25

## 1. 今日速览

项目今日保持高活跃度，尤其在开发与测试环节。**Pull Request 数量激增**，24 小时内更新达 50 条，显示团队正在密集推进功能开发和缺陷修复，为即将到来的 v0.8.3 版本做冲刺准备。**安全问题成为焦点**，一个关于“委派（delegate）机制绕过父进程工具白名单”的严重 Bug 被报告并迅速修复。此外，社区围绕“独立委派模式”和“进程泄漏”等关键问题展开了深入讨论，项目健康度良好，但稳定性与安全性的打磨是当前重点。

## 2. 版本发布

**无新版本发布。** 项目当前处于 v0.8.2 向 v0.8.3 过渡的关键开发期，多个跟踪器（Tracker）Issues 正在协调相关工作。

## 3. 项目进展

今日合并/关闭了 12 个 PR，解决了多个关键问题，项目在稳定性和测试覆盖方面迈出了坚实的一步。

- **修复高风险安全问题**：PR #8190 合并，修复了技能审核（skill audit）中对远程 Markdown 文档链接的误报问题，对应关闭了 `Bug` #6714。
- **修复严重安全漏洞**：Issue #8279 报告了委派（delegate）机制可绕过父进程工具白名单的严重 Bug（S0 级），已在 24 小时内被标记为 `CLOSED`，体现出团队对安全问题的快速响应。
- **增强测试稳健性**：多个测试相关的 PR 被合并，包括：
    - PR #8224：为渠道（channels）添加回归测试，确保 JSON 工具信封在历史修剪时能保留关键字段。
    - PR #8201：将文档链接检查脚本集成到 CI 流程中，防止 PR 引入无效文档链接。
    - PR #8084：修复了诊断模块中自定义提供商验证的问题。
- **CI/CD 与基础设施**：PR #8158 和 #8188 虽然未合并，但分别提出了在 CI 中生成 SBOM（软件物料清单）和加强依赖安全性审查的方案，预示着项目在供应链安全方面的投入。

## 4. 社区热点

今日最受关注的议题集中在**性能稳定性**和**架构演进**上。

- **进程泄漏严重 Bug**：**Issue #5903** `[Bug]: MCP stdio child processes accumulate on daemon with heartbeat.enabled=true (one orphan per tick)` 获得了最多的评论（4 条）。该问题描述了启用心跳检测后，MCP 进程会持续泄漏的严重问题，直接关系到守护进程的稳定性和服务器资源。用户 `rordd` 精确描述了问题复现路径和影响，引发了团队的高度关注。

    [查看 Issue](zeroclaw-labs/zeroclaw Issue #5903)

- **委派机制的功能讨论**：**Issue #8238** `[Feature]: Add independent delegate mode for specialist handoffs` 和 **Issue #7743** `[Feature]: support explicit target-profile authority for delegate handoffs` 均围绕代理委派（Agent Delegation）机制展开。社区希望在默认的“受限委派”基础上，增加一种“独立委派”模式，允许被委派的专家代理使用自己的策略和工具集。这反映了用户对构建更复杂、更专业的多代理协作系统的迫切需求。

    [查看 Issue #8238](zeroclaw-labs/zeroclaw Issue #8238)
    [查看 Issue #7743](zeroclaw-labs/zeroclaw Issue #7743)

## 5. Bug 与稳定性

今日报告了多个高风险 Bug，安全问题尤为突出。

- **S0 - 数据泄露/安全风险**
    - **Issue #8279** `delegate bypasses parent's tool allowlist — sub-agent can invoke tools the parent policy excludes`：**已关闭**。这是最严重的 Bug，子代理通过 `delegate` 工具可以调用父进程白名单之外的工具。庆幸的是，该问题已被快速修复。
    - **Issue #8219** `gpt-oss-120b on Groq fails native multi-turn tool loops`：**已关闭**。特定模型在多轮工具调用循环中因序列化问题而失败。

- **高风险**
    - **Issue #5903** `MCP stdio child processes accumulate on daemon`：**开放中**。这是影响稳定性的关键 Bug，进程泄漏累积会耗尽系统资源，当前尚无关联的 Fix PR，需要重点关注。
    - **PR #8146** `fix(observability): CLI one-shot loses telemetry and token totals on exit`：**开放中**。CLI 单次运行场景下遥测数据丢失，影响使用体验和成本追踪。
    - **PR #8007** `fix(tools/cron): accept job name in cron_update and cron_remove`：**开放中**。Cron 工具只能接受 UUID 作为操作对象，增加用户使用难度。

## 6. 功能请求与路线图信号

多个功能请求表现活跃，并已有对应的实现 PR，很可能被纳入下一版本。

- **独立委派模式 (Independent Delegate Mode)**：Issue #8238 和 #7743 提出的增强功能，旨在提供更灵活的代理协作方式。虽然尚无直接关联的 PR，但这一直是 v0.8.3 路线图上的重要议题，预计后续会有更多进展。
- **LAN 对等发现**：PR #8325 `feat(gateway): add LAN peer discovery hints` 提出了通过 mDNS 进行局域网内服务发现的方案，这对于家庭或小型团队搭建多节点系统非常实用。
- **CI 与安全强化**：PR #8158 (SBOM 生成) 和 PR #8188 (依赖审查强化) 虽然尚未合并，但它们直接回应了社区对供应链透明度和安全性的关切，是项目长期健康发展的积极信号。

## 7. 用户反馈摘要

从今日的 Issue 评论中，可以提炼出用户的几个核心痛点：

- **对稳定性的不满**：用户 `rordd` 在 Issue #5903 中详细描述了 MCP 进程泄漏的现象，并使用“**accumulate（累积）**”和“**orphan（孤儿进程）**”等词，表明该问题对长时间运行的守护进程带来了严重困扰。
- **对安全性的担忧**：用户 `wangmiao0668000666` 报告了 Issue #8279 中委派绕过白名单的严重缺陷，并明确指出其严重性为“**S0 - data loss / security risk**”，反映了专业用户对代理系统安全边界的高度敏感性。
- **对灵活性的需求**：用户 `vrurg` 在 Issue #8238 中对新功能“独立委派模式”的诉求，以及用户 `Audacity88` 在 #7743 中提出的“目标配置文件授权”，都体现出高级用户不满足于默认行为，希望拥有更精细、更自主的控制权，以构建定制化的工作流。

## 8. 待处理积压

以下为需要维护者重点关注、时间跨度较长或风险等级高的开放问题：

- **高风险进程泄漏**：**Issue #5903**（创建于 2026-04-19）是关于 MCP 进程泄漏的核心 Bug，开放已超过两个月且没有关联的修复 PR，可能存在实现难度。这个问题是守护进程稳定性的关键阻碍，应优先投入资源。

    [查看 Issue](zeroclaw-labs/zeroclaw Issue #5903)

- **高风险 Cron 工具可用性**：**PR #8007**（创建于 2026-06-19）旨在解决 Cron 工具的用户体验问题。虽然修复方法看起来直接，但至今仍处于开放待合并状态，可能因为涉及与现有 API 的兼容性讨论。

    [查看 PR](zeroclaw-labs/zeroclaw PR #8007)

- **高风险遥测数据丢失**：**PR #8146**（创建于 2026-06-22）影响了 CLI 的调试和监控能力，作为基础体验问题，其修复应被优先考虑。

    [查看 PR](zeroclaw-labs/zeroclaw PR #8146)

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*