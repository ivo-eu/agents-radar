# OpenClaw 生态日报 2026-06-12

> Issues: 404 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-06-12 06:22 UTC

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

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，现根据 OpenClaw 项目 2026年6月12日的 GitHub 数据，为您呈上今日项目动态日报。

---

# OpenClaw 项目日报 | 2026-06-12

## 1. 今日速览

今日项目活跃度极高，社区讨论与代码提交均达到非常活跃的水平。过去24小时内，共处理了404条Issue和500条PR，显示出强大的社区参与度和维护团队的响应能力。新发布的 **v2026.6.6-beta.2** 版本显著收紧了一系列安全边界和会话状态管理，但同时也引入了关于模型兼容性和 Prompt Cache 失效的回归问题，成为社区讨论的焦点。总体而言，项目在快速迭代中向前迈进，安全性和稳定性是当前的核心主题。

## 2. 版本发布

- **v2026.6.6-beta.2**: 该版本核心聚焦于**安全加固**。
    - **更新内容**：大幅收紧了多个模块的安全边界，包括：
        - 会话记录（transcripts）
        - 沙箱绑定（sandbox binds）
        - 主机环境继承（host environment inheritance）
        - MCP stdio 协议
        - Codex HTTP 访问
        - 原生搜索策略（native search policy）
        - 升级后的发送者检查（elevated sender checks）
        - 已删除代理的 ACP 绕过风险（deleted-agent ACP bypasses）
        - 回环工具（loopback tools）
        - Discord 审核（Discord moderation）
        - Teams 群组操作（Teams group actions）
    - **破坏性变更与迁移注意**：此版本可能对依赖宽松安全策略的配置或自定义插件产生兼容性问题。特别是对 `sandbox`，`MCP stdio` 和 `Codex HTTP` 的访问控制，建议用户仔细检查相关配置，并根据新安全策略进行调整，以避免功能受限或启动失败。

## 3. 项目进展

今日项目在多条关键线上取得进展，多项重要的 Bug 修复和新功能 PR 被合并或推进至待合并状态。

- **核心稳定性修复**：
    - **PR #92331** (merged): 修复了 Webchat 中跨会话消息发送者标签为空的问题。
    - **PR #92336** (merged): 修复了冷启动时 `execSecurity` 配置未正确加载的问题，确保安全策略从网关启动即生效。
    - **PR #86744** (merged): 修复了任务注册表 SQLite 数据库损坏导致网关无法启动的问题，提升了系统健壮性。
    - **PR #92178**: (待合并) 修复了 CLI 命令因设备配对状态异常导致网关 WebSocket 握手崩溃的问题。

- **重要功能推进**：
    - **PR #86655**: 提出为 Claude 模型添加第一方应用服务扩展（Claude-bridge app-server harness），旨在为 Anthropic 模型提供原生工具执行和推理支持，信号强烈。
    - **PR #82434**: 提出为外部插件添加审批验证机制（`externalResolution`），旨在收紧安全模型。
    - **PR #92081**: 提出为 Microsoft Teams 添加语音、聊天和治理集成，显示出项目对协作平台深入集成的决心。

- **生态与开发者体验**：
    - **PR #92347**: 修复了定时任务（cron）无法发出 `model.usage` 诊断事件的问题，增强了可观测性。
    - **PR #92357 & #92341**: 分别修复了混合搜索中关键词结果丢失和 CJK 文本评分过低的问题，改善了记忆搜索质量。

## 4. 社区热点

今日社区讨论焦点主要集中在新版本引发的兼容性问题和尚未修复的长期安全与功能需求上。

- **🔥 Issue #90083**: "[Bug]: 2026.6.1 OpenAI ChatGPT Responses transport fails..." (16条评论, 3个👍)
    - **诉求**：用户升级到 2026.6.1 后，OpenAI 的 GPT-5.4/5.5 模型调用失败。该 Issue 已被关闭，说明修复已合并或提供了解决方案，但短期内仍有大量用户关注此升级带来的兼容性冲击。
- **🔥 Issue #91016**: "⚠️ 升级 2026.6.1 后 DeepSeek Prompt Cache 完全失效..." (5条评论, 4个👍)
    - **诉求**：用户报告升级后 DeepSeek 模型的 Prompt Cache 完全失效，导致成本飙升。虽然评论数较少，但高点赞数表明这是许多用户遇到的普遍痛点，背后是用户对模型推理成本和效率的深切关注。
- **🔥 Issue #32473**: "[Bug]: control ui requires device identity..." (17条评论, 5个👍)
    - **诉求**：用户在 VPS 等非 localhost 环境下无法使用 Control UI，需要 HTTPS 或本地安全上下文。这反映出用户在生产环境部署时的基础设施适配需求。

## 5. Bug 与稳定性

今日报告的 Bug 主要集中在**回归问题**和**配置兼容性**上，严重程度较高。

- **P1 (严重)**
    - **Issue #90083**: `invalid_provider_content_type` for GPT-5.4/5.5. **状态：已关闭**。 *(有修复)*
    - **Issue #91016**: DeepSeek Prompt Cache 失效，成本剧增。 **状态：开放中。** *(未有明确修复 PR)*
    - **Issue #57326**: CLI 辅助路径绕过 CLI 调度。 **状态：开放中。** *(有高风险修复)*
- **P2 (中等)**
    - **Issue #57901**: Safeguard 压缩忽略 `compaction.model` 配置。 **状态：开放中。** *(有修复 PR)*
    - **Issue #52875**: `session_send` 给出无会话的错误。 **状态：已关闭。** *(有修复)*
    - **Issue #63101**: Feishu 频道配置在升级后验证失败。 **状态：已关闭。** *(有修复)*
    - **Issue #63892**: 主动压缩调度器在首次检查点后不再触发。 **状态：已关闭。** *(有修复)*
    - **Issue #63905**: Slack 附件在容器沙箱中失败。 **状态：已关闭。** *(有修复)*
    - **Issue #63612**: 主会话因压缩令牌估算崩溃。 **状态：已关闭。** *(有修复)*

**分析**：今日关闭/修复了大量 P1/P2 级别的 Bug，表明项目维护团队正在积极解决近期的稳定性问题。但 `#91016` 的 Prompt Cache 问题仍未解决，是当前最紧急的风险点。

## 6. 功能请求与路线图信号

社区持续提出高价值的功能请求，反映出项目向更成熟、更安全的企业级平台演进的趋势。

- **高热度功能请求 (高评论数/点赞)**
    - **Issue #9443**: 提供预构建的 Android APK 发布包 (25条评论)。社区对移动端运行 OpenClaw 有强烈需求。
    - **Issue #12602**: 支持 Slack Block Kit (13条评论)。用户渴望更丰富的交互体验，而非简单文本。
    - **Issue #13583**: 预响应执行挂钩（硬性阻拦）(11条评论, 2个👍)。在高风险场景下，需要机械性的强制规则，而非软性提示。
    - **Issue #7707**: 按来源进行内存信任标记 (8条评论)。用户对“记忆投毒”攻击有强烈防范意识，体现了安全需求向 AI 系统内部逻辑渗透。

- **路线图信号**：
    - 目前有多个开放中的 **PR** 正在推进类似功能：**PR #82434** 的插件审批机制与 `#12678`（能力权限）和 `#12219`（技能权限清单）的功能请求高度吻合。**PR #86655** 的 Claude Bridge 反映了社区对多模型深度集成的期待。
    - **Issue #13616** 的备份/恢复工具和 **Issue #13610** 的密钥管理集成虽评论数不多，但它们是生产环境下运维的基石，很可能作为中期路线图中的关键特性被采纳。

## 7. 用户反馈摘要

从今日的 Issue 和 PR 评论中，可以提炼出以下用户声音：

- **升级恐惧与代价**：多位用户报告了从 v2026.5.7 或较旧版本升级到 2026.6.x 后出现的严重问题（`#90083`, `#91016`, `#63101`）。这表明近期的安全加固和依赖升级可能引入了非预期的、影响核心功能的回归。用户付出了时间和金钱成本，对升级的稳定性表达了隐忧。
- **生产环境部署的痛点**：用户不仅关注功能，更关注如何可靠、安全地部署在服务器（VPS）和云环境上。`#32473`（Control UI 需要 HTTPS）和 `#13616`（缺乏备份工具）直接反映了这一需求。
- **对“软规则”的不信任**：在 `#13583` 等讨论中，来自金融、安全行业的用户明确表示，仅依赖 Prompt 指令是不够的，他们需要系统层级的、不可绕过的强制规则。
- **对 AI Agent 能力的更高要求**：用户不再满足于简单的文本回复，`#12602`（Slack Block Kit）、`#33413`（显示工具级进度）和 `#40418`（会话记忆保存）表明，用户期待 Agent 能提供更丰富、更透明、更具有上下文持续性的交互体验。

## 8. 待处理积压

以下为一些长期未响应或存在开发瓶颈的重要 Issue/PR，提醒维护者关注：

- **Issue #7707** "[Feature]: Memory Trust Tagging by Source" (创建于 2026-02-03)
    - **状态**：标记为 `needs-product-decision`。核心安全特性，已搁置超过4个月，社区有明确需求，建议尽快做出产品方向决策。
- **Issue #12219** "[Feature]: Skill Permission Manifest Standard" (创建于 2026-02-09)
    - **状态**：标记为 `needs-maintainer-review` & `needs-security-review`。与近日合并的 `#82434` PR 功能相关，应重新评估优先级，推动其标准化进程。
- **PR #66110** "fix(logging): support root log file overrides" (创建于 2026-04-13)
    - **状态**：`triage: needs-real-behavior-proof`。一个持续近两个月的、关于基本运维信息（日志隔离）的 PR 仍卡在验证环节。对于多实例或多服务的用户来说，这是一个关键的运维痛点，建议尽快完成测试，推动合并。
- **Issue #6615** "[Feature]: Add denylist support for exec-approvals" (创建于 2026-02-01)
    - **状态**：`clawsweeper:needs-product-decision`。该 Issue 获得了社区最多的点赞（7个👍），是用户最强烈的安全功能需求之一，需要产品团队给出明确决策。

---

## 横向生态对比

好的，作为AI智能体与个人AI助手领域开源项目资深技术分析师，现根据您提供的各项目动态日报，为您呈上2026年6月12日的横向对比分析报告。

---

# 个人AI助手开源生态全景分析报告 | 2026-06-12

## 1. 生态全景

当前个人AI助手/自主智能体开源生态呈现**高度活跃、快速迭代、分化与趋同并存**的态势。一方面，以**OpenClaw**和**ZeroClaw**为首的旗舰项目发布了重大版本更新，核心围绕**安全加固**和**多Agent架构**展开，标志着生态从“功能可用”向“生产可用”和“复杂协作”演进。另一方面，大量项目（如**IronClaw、CoPaw、NanoBot**）同步在稳定性、UI体验和开发者基础设施上发力，**模型兼容性**和**跨平台部署**成为普遍痛点。生态整体健康，但**升级引发的回归问题**和**长期积压的功能PR**是各项目共同的挑战。

## 2. 各项目活跃度对比

| 项目名称 | 活跃 Issues (新开/活跃) | 活跃 PRs (提交/待审) | 版本发布 | 健康度评估 |
| :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 404 (高) | 500 (极高) | v2026.6.6-beta.2 | 极高，生态核心，快速迭代，但安全收紧引发兼容性阵痛 |
| **IronClaw** | 16 (高) | 22 (高) | 无 | 高，处于Reborn版本上线前冲刺，Bug修复与功能并重 |
| **CoPaw** | 13 (高) | 21 (高) | v1.1.11.post1/post2 | 高，后端重构与UI修复并行，社区反馈积极 |
| **ZeroClaw** | 新增较多v0.8.0兼容Bug | 中等 | **v0.8.0 (里程碑)** | 高，大版本发布后处于密集修复期，macOS问题突出 |
| **NanoBot** | 3 (中) | 13 (中) | 无 | 良好，功能类PR增多，但MCP网关稳定性是关注点 |
| **PicoClaw** | 低 | 35 (高) | Nightly | 高，代码质量大幅提升，社区贡献踊跃，长期积压问题待解 |
| **Hermes Agent** | 0 (低，大量已关闭) | 50 (极高) | 无 | 优秀，修复效率惊人，专注稳定性，但审核瓶颈明显 |
| **NanoClaw** | 1 (低，但涉及架构重构) | 13 (中) | 无 | 良好，产出质量高，聚焦核心Bug和架构演进 |
| **LobsterAI** | 1 (低) | 6 (低) | 无 | 良好，今日合并大量UX修复，解决历史痛点 |
| **NullClaw** | 1 (低) | 0 | 无 | **停滞**，仅有1个新Bug报告，核心开发活动近乎停止 |
| **TinyClaw** | 0 | 0 | 无 | **停滞** |
| **ZeptoClaw** | 0 | 0 | 无 | **停滞** |
| **Moltis** | 1 (低) | 1 (低) | 无 | **低**，活跃度低，仅有零散讨论 |

## 3. OpenClaw在生态中的定位

OpenClaw凭借其**空前活跃的社区规模**（24小时内处理404条Issue和500条PR）和**系统化的安全哲学**，毫无疑问是当前生态的**领航者和参照标准**。

- **核心优势**：其 **“安全先行”** 的迭代策略（如v2026.6.6-beta.2的全面安全收紧）领先于几乎所有同类项目，彰显了对企业级场景的野心。在**多平台集成**（Discord、Teams）和**协议支持**（MCP、Codex）的广度上也处于领先地位。
- **技术路线差异**：与追求轻量化和速度的**PicoClaw**、**NanoClaw**不同，OpenClaw更强调**复杂治理模型**（如审批、审计、权限隔离）。与正进行大规模后端重构的**CoPaw**和**IronClaw**相比，OpenClaw更侧重于**跨平台网关的安全性**和**生态兼容性**。
- **社区规模**：OpenClaw的Issues和PR活跃度是第二名**IronClaw**的数十倍，远超其他项目。这表明其拥有最庞大的用户基础和开发者生态，但也意味着其面临的噪音和社区管理压力最大。

## 4. 共同关注的技术方向

| 技术方向 | 核心诉求 | 涉及项目 |
| :--- | :--- | :--- |
| **模型兼容性与成本** | 1. 新模型（Gemini 3.5、Kimi K2、GPT-5.x）的快速适配问题频发。 <br> 2. Prompt Cache失效导致成本飙升（OpenClaw）。 <br> 3. Token消耗透明化及审计（CoPaw、LobsterAI）。 | OpenClaw, PicoClaw, Hermes Agent, ZeroClaw, CoPaw |
| **多Agent协作与通信** | 1. 子Agent/多Agent架构落地（ZeroClaw v0.8.0）。 <br> 2. Agent间状态同步、会话继承、工作流执行（NanoClaw, ZeroClaw）。 | OpenClaw, ZeroClaw, PicoClaw, NanoClaw, CoPaw |
| **安全与治理增强** | 1. 从软提示转向系统级强制规则，如执行前挂钩、插件审批。 <br> 2. 多租户/频道级安全配置（OpenClaw, PicoClaw）。 | OpenClaw, NanoBot, Hermes Agent, PicoClaw, ZeroClaw |
| **平台集成深入** | 1. 渠道功能对等性（如Slack Block Kit、Signal Reactions）。 <br> 2. 消息投递可靠性（WhatsApp、Mattermost）。 | OpenClaw, NanoBot, PicoClaw, NanoClaw, Moltis |
| **部署与运维痛** | 1. 跨平台兼容性，特别是Windows/macOS和反向代理环境。 <br> 2. 内存泄漏、OOM、网关崩溃等稳定性问题。 | OpenClaw, ZeroClaw, NanoBot, Hermes Agent, CoPaw, IronClaw |

## 5. 差异化定位分析

| 项目 | 功能侧重 | 目标用户 | 技术架构关键差异 |
| :--- | :--- | :--- | :--- |
| **OpenClaw** | 全能型企业级平台，安全与多平台集成 | 大型团队、企业、对安全合规要求高的组织 | 模块化网关 + 策略引擎，强调安全隔离与审计 |
| **NanoBot** | 灵活开发框架，MCP网关与SDK扩展 | 开发者、需要深度定制AI助手的创客 | 轻量级核心 + 丰富的SDK和Provider抽象层 |
| **PicoClaw** | 轻量级、嵌入式友好，注重跨平台 | 个人开发者、边缘设备、Windows/Mac用户 | 代码精简，依赖较少，社区贡献门槛低 |
| **ZeroClaw** | **多Agent架构**的先行者，聚焦工作流自动化 | 需要复杂任务编排、DevOps集成的组织 | 守护进程管理多具名Agent，原生支持ACP协议 |
| **IronClaw** | **Reborn**版本的质量巩固，UI与可观测性 | 追求稳定、一流用户体验的终端用户 | 大型前端（WebUI）与后端（Reborn引擎）分离架构 |
| **CoPaw** | **智能化重写**，聚焦记忆系统与Plugin生态 | 追求AI原生体验的知识工作者 | 后端向AgentScope 2.0迁移，强调Plugin与知识管理 |
| **Hermes Agent** | **最快的Bug修复响应**，稳定性优先 | 注重服务连续性和稳定性的生产环境用户 | 高度自动化的CI/CD和Bug分类流程 |
| **NanoClaw** | **架构演进**，探索内存系统与多实例 | 对底层技术架构和扩展性敏感的高级用户 | 强调核心架构的重新设计，解决可扩展性瓶颈 |
| **LobsterAI** | **C端体验**优化，实时协作与多媒体输入 | 普通互联网用户、追求流畅交互的群体 | 集成网易云服务，重前后端一体化体验 |
| **NullClaw** | **停滞** | - | - |

## 6. 社区热度与成熟度

- **快速迭代、高度活跃**：**OpenClaw、IronClaw、ZeroClaw** 处于功能发布或重大重构的密集期，社区反馈和代码提交均异常活跃，是最有看点的发展梯队。
- **质量巩固、稳定输出**：**Hermes Agent、NanoBot、PicoClaw、NanoClaw、LobsterAI** 在体验和稳定性上持续打磨，修复和PR合并率高，是可靠的生态组成部分。
- **转型调整、健康度高**：**CoPaw** 虽然活跃，但后端重构导致部分Bug堆积，处于“建设性阵痛”期。**Moltis** 活跃度偏低，但尚未出现负面信号。
- **生态边缘、关注度低**：**NullClaw、TinyClaw、ZeptoClaw** 今日近乎无活动，核心开发动力不足，除非有重大发布，否则正逐步退出竞争核心区。

## 7. 值得关注的趋势信号

1. **“安全”从软性建议走向硬性规则**：社区（尤其是金融等行业用户）对**系统级强制规则（硬性阻拦、执行前钩子）** 的需求明确且强烈（OpenClaw #13583, NanoBot #82434）。这预示着下一阶段智能体安全将从“提示对抗”转向“架构隔离与控制”。
2. **模型兼容性成最大“成本陷阱”**：新模型（Gemini 3.5, DeepSeek）的快速迭代频繁破坏Agent框架的兼容性（PicoClaw #3111, OpenClaw #91016），而**Prompt Cache失效**等问题直接导致用户成本不可控。这意味着**模型网关的兼容适配速度和成本审计能力**将成为Agent框架的关键竞争力。
3. **多Agent协作从“概念验证”走向“生产部署”**：ZeroClaw v0.8.0的发布，以及多个项目（NanoClaw, PicoClaw）对协作总线的探索，表明**多Agent编排、会话传递、工作流调度**不再是愿景，而是正在落地的现实。专注于**子Agent生命周期管理和跨Agent通信协议（ACP/A2A）** 的项目将占据先机。
4. **跨平台“最后一公里”仍是普遍短板**：macOS启动失败（ZeroClaw）、Windows路径分隔符（PicoClaw）、反向代理下WebUI崩溃（Hermes Agent）等问题反复出现。这表明，**将Agent从开发环境平滑部署到桌面和服务器的完整流程**，仍是所有项目尚未完美解决的共性问题，也是差异化竞争的机会点。

**对AI智能体开发者的参考价值**: 选择框架时，应将**其安全架构的成熟度**（是否能满足你的合规要求）、**模型适配的灵活性和透明度**（是否可能导致意外成本）以及**跨平台/生产部署的成熟度**，置于是否满足功能需求之上。同时，**多Agent协作**已在路上，提前评估框架的扩展性至关重要。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot 项目动态日报 (2026-06-12)

---

## 1. 今日速览

过去24小时内，NanoBot 项目保持高活跃度：累计处理 **4 个 Issue**（新开/活跃 3 个，关闭 1 个），**17 个 PR**（其中 13 个待合并，4 个已合并/关闭）。社区聚焦于 **多自定义提供商支持**、**MCP 网关稳定性** 以及 **后轮 consolidation 导致的上下文丢失** 等核心问题。功能类 PR 数量显著增加，表明项目在持续拓展 SDK、调度、技能缓存等基础设施能力，同时 Bug 修复也同步跟进，整体健康度良好。

---

## 2. 版本发布

无新版本发布。

---

## 3. 项目进展

本日共有 **4 个 PR 被合并或关闭**，主要推进方向如下：

| PR 编号 | 类型 | 标题 | 说明 |
|---------|------|------|------|
| [#4289](https://github.com/HKUDS/nanobot/pull/4289) | ✅ 合并 | feat(slack): add groupRequireMention to scope allowlist channels | 为 Slack 频道增加 `groupRequireMention` 选项，可在白名单模式下仅响应 @提及，提升频道管理灵活性。 |
| [#4020](https://github.com/HKUDS/nanobot/pull/4020) | ✅ 合并 | feat(providers): make stream-idle timeout configurable per-provider | 将流式空闲超时 (原硬编码 90s) 改为按 provider 可配置，解决本地模型在大提示下超时问题。 |
| [#4298](https://github.com/HKUDS/nanobot/pull/4298) / [#4297](https://github.com/HKUDS/nanobot/pull/4297) | ✅ 关闭 | Worktree feature+hermes research doc | 两份研究文档 PR，虽已关闭但可能为后续特性提供设计参考。 |

**项目整体向前迈进**：Slack 集成更精细、Provider 配置更灵活、同时通过多个开放 PR（如 cron 绑定会话、子代理模型预设、Python SDK 扩展等）持续增强核心能力。

---

## 4. 社区热点

### 最活跃讨论：多自定义提供商支持

- **Issue [#4305](https://github.com/HKUDS/nanobot/issues/4305)** （enhancement，24h 内 1 条评论）  
  用户 `smurfix` 提出需要多个 `custom` 和 `openai` 类型的 provider，建议在 Provider 节增加 `template` 参数以选择内置 provider。该诉求与长期开放的 **PR [#3239](https://github.com/HKUDS/nanobot/pull/3239)**（支持多自定义 OpenAI 兼容提供商）直接对应，PR 自 4 月 17 日开启至今未合并，评论数为 0，但 Issue 的热度可能推动其重新评估。

- **PR [#3239](https://github.com/HKUDS/nanobot/pull/3239)**（4月17日开启，至今 60+ 天）  
  该 PR 是本日 Issue #4305 的代码实现，两者形成强烈呼应。社区对该功能的实际需求已从“单例”升级为“多实例”，预计会在后续版本中被优先考虑。

### 其他高关注 Issue

- **Issue [#4307](https://github.com/HKUDS/nanobot/issues/4307)**（bug，24h 内新开，0 条评论）  
  用户 `MARJORIESHA-pBAD` 报告了后轮 consolidation 导致 agent 自身回复被归档、用户后续追问丢失的问题，涉及上下文窗口管理的关键逻辑。

---

## 5. Bug 与稳定性

### 严重程度：高

| 编号 | 标题 | 状态 | 是否已有 Fix PR |
|------|------|------|----------------|
| [#4302](https://github.com/HKUDS/nanobot/issues/4302) | nanobot gateway crashes after MCP reconnect | 🟡 开放 | ✅ **PR [#4303](https://github.com/HKUDS/nanobot/pull/4303)**（已提交，待合并） |
| [#4307](https://github.com/HKUDS/nanobot/issues/4307) | Post-turn consolidation wipes agent’s own delivery message | 🟡 开放 | ❌ 尚无 PR |
| [#4236](https://github.com/HKUDS/nanobot/issues/4236) | bwrap sandbox fails on Ubuntu 24.04 (restricted user namespaces) | ✅ 已关闭（已处理？） | 用户 `primit1v0` 报告，但 Issue 已关闭，未说明解决方案，需确认是否已在某 PR 中修复。 |

### 严重程度：中

- **PR [#4306](https://github.com/HKUDS/nanobot/pull/4306)**（bug fix）  
  修复了会话历史中出现孤立 tool 结果的问题（#4006），避免被严格 API 拒绝或渲染异常。该问题难以复现，但维护者已给出修复。

- **PR [#4304](https://github.com/HKUDS/nanobot/pull/4304)**（bug fix）  
  修复 Cron 任务在子代理尚未完成时即标记为完成的问题（关联 #4290），确保子代理任务彻底结束。

---

## 6. 功能请求与路线图信号

### 高优先级信号

1. **多自定义 Provider（#4305 + PR #3239）**  
   需求明确且已有代码实现，预计将在下一轮合入。若未及时合并，可能因 conflict 导致 review 成本增加。

2. **子代理模型预设（PR #4291）**  
   允许 spawn 子代理时使用不同的模型 preset，并通过 `agents.defaults.spawnPresets` 控制。该功能对多模型 pipeline 场景至关重要，已处于开放待审状态。

3. **Cron 绑定到会话（PR #4299）**  
   将计划任务绑定到指定会话，实现更精细的自动化调度。此特性与 spawn 结合将显著增强 agent 工作流能力。

### 中等信号

- **Python SDK 扩展（PR #4296）**：提升开发者体验，增加 RunResult、会话持久化等能力。
- **技能缓存（PR #4301）**：缓存 skill 元数据，减少重复扫描，提升 agent 启动性能。
- **技能类型要求检查（PR #4300）**：为技能增加类型依赖管理，便于组合复杂技能。

### 长期积压

- **PR #3538**：添加 gateway start/stop/restart 命令，已开放 44 天，尚未合并。该 PR 补充了 CLI 运维能力，可能因 gateway 重构而被搁置。

---

## 7. 用户反馈摘要

从本日活跃 Issues 的评论中可提炼以下真实用户痛点：

- **Provider 灵活性不足**：“I need more than one 'custom' (and 'openai') provider.” (#4305) —— 用户期望一个配置模板机制，以便对接多个内部 API 或云厂商。
- **上下文管理敏感**：当 `context_window_tokens` 设为 40k 时，长时间多轮对话积累超过 100k token，consolidation 仅在回合结束后触发，导致 agent 自身回复被归档，用户后续引用丢失 (#4307)。该问题直接影响多轮交互的连续性。
- **沙箱兼容性问题**：Ubuntu 24.04 默认限制用户命名空间，导致 bwrap 沙箱失效 (#4236)。用户期望更清晰的错误提示或自动降级方案。
- **MCP 重连稳定性**：streamableHttp MCP 服务器 session 终止后，gateway 在重连时崩溃，日志显示 `RuntimeError: Attempted to exit cancel scope in a different task`。用户希望崩溃变重连 (#4302)。

---

## 8. 待处理积压

以下为开放时间较长、尚未获得维护者响应的关键 PR/Issue，建议关注：

| 类型 | 编号 | 标题 | 开放天数 | 状态 |
|------|------|------|---------|------|
| PR | [#3239](https://github.com/HKUDS/nanobot/pull/3239) | feat: support multiple custom OpenAI-compatible providers | 56 天 | 开放，无维护者评论 |
| PR | [#3538](https://github.com/HKUDS/nanobot/pull/3538) | feat: add gateway start/stop/restart commands | 44 天 | 开放，无维护者评论 |
| PR | [#4021](https://github.com/HKUDS/nanobot/pull/4021) | fix(codex): dedup reasoning items before send | 16 天 | 开放，AI-assisted，待合并 |
| Issue | [#4307](https://github.com/HKUDS/nanobot/issues/4307) | Bug: Post-turn consolidation wipes agent's own delivery message | 0 天（新开） | 尚未有 PR 匹配，需迅速定位 |
| Issue | [#4302](https://github.com/HKUDS/nanobot/issues/4302) | nanobot gateway crashes after MCP reconnect | 1 天 | 已有 PR #4303 等待合并 |

**特别提醒**：PR #3239 与今日新增的 Issue #4305 直接相关，若长期搁置可能影响社区对功能开发效率的信心。建议维护者评估其合理性并推动合并或给出明确拒绝理由。

---

**数据来源**：GitHub 仓库 [HKUDS/nanobot](https://github.com/HKUDS/nanobot)  
**生成时间**：2026-06-12  
**分析师**：AI 智能体与个人 AI 助手领域开源项目分析师

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，根据您提供的 Hermes Agent 项目数据，我为您生成了 **2026年6月12日** 的项目动态日报。

---

### Hermes Agent 项目日报 | 2026年6月12日

#### 1. 今日速览

本项目今日活跃度极高。虽然无新版本发布，但社区贡献和问题解决效率惊人：**24小时内关闭了14个Issue**，且所有关闭的Issue均由自动化工具或快速修复完成（`sweeper:implemented-on-main`）。**50条Pull Request的更新量**表明项目正处于密集的开发与集成阶段。总体来看，项目健康度优秀，修复速度快，但大量PR待合并暗示了核心维护者的审核压力。

- **项目状态**: 🔴 **高活跃度** | **稳定性修复为主**
- **核心指标**: `Issues 关闭数` > `新开数` (14:0) | `PR 处理量` 50条

#### 2. 版本发布

**无新版本发布。**

#### 3. 项目进展

今日项目在稳定性和兼容性修复上取得了显著进展，27个PR被合并或关闭，主要工作集中在修复各类生产环境中的关键问题，提升了网关、CLI和核心Agent的健壮性。

- **✅ 网关稳定性增强**：
    - **修复内存泄漏 (PR #25315)**: 针对网关进程在长期运行后内存暴涨至20-37GB导致OOM的问题，从根源上修复了`_evict_cached_agent`、`_session_messages`等多个缓存和会话生命周期管理漏洞。
    - **修复重启机制 (PR #25298, #25217)**: 解决了Docker容器内`/restart`命令导致永久关闭而非重启的Bug，并修复了系统服务用户与执行用户不同时，`hermes update`命令路径配置错误的问题。
    - **平台消息传递修复 (PR #25265, #25258, #25181)**: 修复了Slack平台在流式传输失败时产生重复回复的问题，并解决了Mattermost网关在首次回复后无法正确追踪线程的Bug。
- **✅ Agent 核心机制修复**:
    - **修复自动故障切换 (PR #25261)**: 解决了当主要API提供商返回429（限流）或账单错误时，Agent不会自动切换到备用模型的问题，确保服务连续性。
    - **修复上下文压缩错误 (PR #25265)**: 修复了会话摘要会保留失败的shell命令，导致后续Agent上下文被误导的问题。
- **✅ 兼容性与跨平台修复**:
    - **AI提供商兼容性 (PR #44676, #25301)**: 修复了在Fireworks上托管Kimi K2模型时，因JSON Schema关键字`$ref`与`default`冲突导致HTTP 400错误的问题。同时修复了DeepSeek插件未能正确传递`reasoning_effort`参数的问题。
    - **Windows平台支持 (PR #25340)**: 修复了在Windows无头服务（如计划任务）中运行时，因缺少Win32控制台缓冲区导致程序直接崩溃的问题。

#### 4. 社区热点

今日社区讨论的热点集中在**对稳定性问题的快速反馈与修复**上。虽然所有活跃讨论的Issues都已被关闭，但其背后的用户诉求非常明确。

1.  **Issue #25378 - Azure OpenAI 配置问题**
    - **链接**: [NousResearch/hermes-agent Issue #25378](https://github.com/NousResearch/hermes-agent/issues/25378)
    - **热度**: 4条评论
    - **分析**: 用户报告在配置Azure OpenAI环境时遇到`NotFoundError [HTTP 404]`。这反映了企业用户对官方支持的云服务提供商有明确的集成需求。虽然Issue已关闭，但背后可能意味着Azure代理的配置文档或自动检测逻辑需要进一步优化。

2.  **Issue #25281 - 更新Dashboard删除所有定时任务**
    - **链接**: [NousResearch/hermes-agent Issue #25281](https://github.com/NousResearch/hermes-agent/issues/25281)
    - **热度**: 4条评论
    - **分析**: 用户报告一个严重P1 Bug：点击Dashboard的“更新”按钮会清空所有定时Cron任务。这暴露了运维上的重大痛点，用户强烈期望有更健壮的持久化存储方式（如数据库或配置文件替代内存），而非在更新时被覆盖。

3.  **Issue #25351 - Dashboard聊天在反向代理后静默崩溃**
    - **链接**: [NousResearch/hermes-agent Issue #25351](https://github.com/NousResearch/hermes-agent/issues/25351)
    - **热度**: 3条评论
    - **分析**: 技术用户报告在反向代理环境下，每次打开`/chat`标签页都会触发一次耗时15秒的`npm run build`，阻塞整个事件循环，导致WebSocket连接超时。该问题揭示了部署场景下的架构缺陷，尤其是在容器化或代理环境下的启动流程优化需求。

#### 5. Bug 与稳定性

今日报告的14个Bug全部为已有复现/修复的状态。按严重程度排列：

- **P1 (严重)**:
    - **内存泄漏导致OOM (Issue #25315)**: 网关进程在20-35小时内内存增长到20-37GB。✅ **已有对应的修复PR被合并。**
    - **更新操作删除定时任务 (Issue #25281)**: `Update`按钮导致Cron任务丢失。✅ **标记为已由sweeper修复。**
    - **上下文压缩错误 (Issue #9631)**: 迭代上下文压缩导致旧话题抢占当前活跃话题。✅ **已关闭，可能是旧Bug被彻底修复。**
- **P2 (主要)**:
    - **反向代理下聊天崩溃 (Issue #25351)**: 由于`npm run build`阻塞导致。✅ **标记为已由sweeper修复。**
    - **Telegram网关“假死” (Issue #25221)**: 轮询冲突恢复失败后，网关进程存活但无法处理消息。✅ **标记为已由sweeper修复。**
    - **Mattermost线程失效 (Issue #25181)**: 首次回复后无法继续线程对话。✅ **标记为已由sweeper修复。**
    - **Shell钩子不触发 (Issue #25204)**: `pre_tool_call`和`on_session_finalize`钩子在特定上下文中失效。✅ **标记为已由sweeper修复。**
    - **Dashboard指令显示截断 (Issue #18376)**: 安全审批弹窗命令不换行，回归问题。✅ **已关闭。**
- **P3 (较低)**:
    - **Kimi编码设置冲突 (Issue #25104)**: Anthropic模式路径生成错误。✅ **标记为已由sweeper修复。**

#### 6. 功能请求与路线图信号

- **高频需求：插件系统扩展 (Issue #25262, PR #44703)**: 用户提出为插件上下文API添加`pre/post tool-call`钩子，以便于实现安全审计、干运行或自定义逻辑。该需求与社区热门的`everything-claude-code`项目模式一致，且已有相关的PR (`#44703`) 提出了具体实现方案。这将是**下个版本最可能被纳入的功能之一**，也是项目向更开放、可扩展Agent框架演进的重要信号。
- **平台及执行环境扩展 (PR #44158)**: 有PR提议添加WSL2执行后端，让Windows用户能原生体验Linux环境，这反映了**提升跨平台开发体验**的强烈社区意愿。

#### 7. 用户反馈摘要

- **痛点**:
    - **配置门槛高**: 用户反映`Azure OpenAI`配置困难，`hermes gateway install`命令因代码内嵌emoji导致执行失败 (`Issue #25191`)。
    - **运维风险大**: “更新按钮会删除定时任务”这类P1 Bug让用户对Dashboard的操作缺乏信任感，期望更健壮的持久化方案 (“Cron jobs should be stored in a database”)。
    - **部署陷阱多**: 在Docker和反向代理后的部署都需要特定的workaround (“/restart”命令不工作，“chat”页面卡死)，表明在非标准部署场景下的测试和文档仍有欠缺。
    - **功能期待**: “希望能在Windows上无Docker运行Linux环境” (WSL2 backend PR)。“希望插件系统能拦截和分析工具调用” (pre/post hooks)。

#### 8. 待处理积压

以下PR处于**开放状态**已超过一个月，可能因为需要更多讨论、测试或维护者时间，值得关注：

- **PR #12525 - Home Assistant服务调用状态验证**: `2026-04-19` 创建，已开放超50天。该功能旨在解决HA API虚假成功响应的问题，对智能家居场景至关重要，属于稳定性增强。
- **PR #11457 - TTS OGG转换平台限定**: `2026-04-17` 创建，已开放超50天。该功能让TTS音频格式更合理，避免非必要转换。虽为P2，但卡顿已久，可能需维护者介入决策。
- **PR #11455 - Mistral API兼容性补丁**: `2026-04-17` 创建，已开放超50天。该PR对Mistral用户至关重要。如果Mistral用户增长，此PR的合并优先级应提高。

---
**分析总结**: 今日日报显示项目正处于一个**以稳定性巩固为主导的密集维护期**。问题修复速度极快（尤其依赖自动化工具），社区活跃度很高。但积压的长期PR表明，在核心功能开发路径上存在一些“瓶颈”，需要核心团队对这些长期开放的功能性PR进行优先级排序和决策信号释放，以保持社区的长期贡献热情。

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

好的，作为 PicoClaw 开源项目的 AI 分析师，根据您提供的 2026-06-12 GitHub 数据，我为您生成了以下项目动态日报。

---

# PicoClaw 项目动态日报 | 2026年6月12日

## 1. 今日速览

今日 PicoClaw 项目继续保持高歌猛进的活跃态势。社区贡献者不仅提交了 35 条 Pull Request，其中修复了大量关键代码质量和稳定性问题，还关闭了 4 个历史 Issue，显示出项目在 bug 修复和代码健壮性上的投入。尽管新发布的 Nightly 版本提示可能存在不稳定性，但社区围绕 Windows 兼容性、新模型 Gemini 3.5 适配以及 WebSocket 协议完善等核心议题的讨论持续升温，推动项目向更好用、更通用的方向演进。项目整体健康度良好，社区贡献者与维护者间的协作高效。

## 2. 版本发布

- **Nightly 构建发布**: `v0.2.9-nightly.20260612.413d3749`
  - **概述**: 此版本为自动化构建的 Nightly 版本，主要包含最新的主分支代码，旨在让用户和开发者提前体验最新功能。
  - **更新内容**: 完整变更日志见 [v0.2.9...main](https://github.com/sipeed/picoclaw/compare/v0.2.9...main)。根据今日的 PR 活动，此版本应包含了多项错误处理改进、依赖更新以及一些社区贡献的 Bug 修复。
  - **破坏性变更**: 该版本为 Nightly 构建，未明确标注破坏性变更。
  - **迁移注意事项**: 由于是自动化构建，可能存在不稳定性。建议用于测试和开发环境，生产环境用户请谨慎升级。在升级前，建议关注近期关闭的 Issue 和合并的 PR，特别是涉及权限、核心库（如 AWS SDK、MCP SDK）的更新，确保与新版本的兼容性。

## 3. 项目进展

今日项目在代码质量和依赖管理方面取得了显著进展。

- **依赖大规模升级**: 多个自动依赖更新 PR 已合并，包括 AWS SDK 相关组件 (`#3106`, `#3102`) 和 MCP SDK (`#3098`)，确保了项目与外部生态的同步。
- **关键功能修复与合并**:
  - **Copilot SDK 升级**: `#3107` 将 `copilot-sdk/go` 从 `0.2.0` 重磅升级至 `1.0.1`，这可能引入了不兼容的变更，标志着 PicoClaw 对 GitHub Copilot 集成的正式化支持进入稳定阶段。
  - **WhatsApp 原生模式修复**: PR `#2934` 已合并，修复了 WhatsApp 通道在启用 `use_native` 标志后无法正常工作的问题，为使用 whatsmeow 的用户扫清了障碍。
  - **MCP 动态消息头支持**: PR `#2696` 的合并是一项重要的功能增强，允许渠道根据上下文动态地向 MCP 服务器传递自定义 HTTP 头部（例如从用户上下文传递 `Authorization`），极大地增强了 MCP 集成的灵活性和安全性。
  - **通道配置保留修复**: PR `#2956` 修复了加载安全配置文件时可能错误覆盖通道启用状态的问题，确保了配置的一致性。

## 4. 社区热点

今日社区讨论热度集中在以下几个议题：

- **WebSocket 协议完善 (热度最高)**: Issue `#2984` “为 WebSocket 客户端添加显式完成信号” 获得了 2 个 👍 和持续的讨论。这表明社区开发者正尝试基于 PicoClaw 构建更复杂的应用（如聊天界面），他们需要一个清晰、确定的 API 信号来判断 Agent 是否已完成处理。这不仅是功能请求，更是对 PicoClaw 作为平台型基础设施的期望。
  - [Issue #2984](https://github.com/sipeed/picoclaw/issues/2984)

- **Windows 平台兼容性 (持续关注)**: Issue `#2472` 关于 `list_dir` 在 Windows 上报错的问题，虽非今日新开，但仍有 5 条评论。这反映了跨平台兼容性，尤其是处理路径分隔符等细节问题，是用户迁移和使用的关键痛点。
  - [Issue #2472](https://github.com/sipeed/picoclaw/issues/2472)

- **新模型支持兼容困境**: Issue `#3111` 指出新兴模型 `gemini-3.5-flash` 的工具调用因 `thought_signature` 参数问题而失败。这显示快速迭代的 LLM 厂商不断推出新特性，给下游的 Agent 框架带来了持续适配的压力。该问题暂无评论和修复 PR，需要维护者关注。
  - [Issue #3111](https://github.com/sipeed/picoclaw/issues/3111)

## 5. Bug 与稳定性

今日报告了多个 Bug，按严重程度排列如下：

- **高**:
  - **`gemini-3.5-flash` 工具执行失败** (`#3111`): 新模型与现有 schema 不兼容，直接导致工具调用功能完全不可用。**无对应 Fix PR**。
- **中高**:
  - **子代理任务完成导致消息重复** (`#3094`): 核心协作机制出现缺陷，会导致用户收到重复消息，严重影响体验。**无对应 Fix PR**。
  - **Windows 路径分隔符错误** (`#2472`): 长期存在的平台兼容性问题，影响核心文件操作功能。
- **中**:
  - **Telegram Forum 话题回复错乱** (`#3110`): 在 Telegram 群组话题模式下，机器人回复会被发送到错误的默认话题，破坏了群组秩序。**无对应 Fix PR**。
  - **无视觉能力的模型图像描述产生幻觉** (`#3108`): 当模型不支持视觉时，其对图像的描述会是虚假的。这属于模型选型不当导致的非预期行为，但框架应提供更好的兜底或提示机制。
- **低**:
  - **JSON 序列化错误静默忽略** ( `#3113`, `#3112`): 这两条已提交的 PR 共同修复了在工具调用和配置管理中因丢失 JSON 错误检查而导致的数据丢失风险。虽然目前可能未触发严重问题，但修复后显著提升了代码健壮性，防患于未然。

## 6. 功能请求与路线图信号

- **核心协议增强**:
  - **Agent 完成信号**: Issue `#2984` 提出的为 WebSocket 客户端添加明确的 `<turn_end>` 信号，是构建无缝交互体验的基础。预期此特性将很快被纳入开发路线图。
  - **频道级权限控制**: Issue `#3109` 已关闭，但其提出的“在频道内区分安全/不安全操作”的想法，是应对多租户、群组协作场景的关键。虽然被关闭，但该话题可能在未来以其他形式（如更细粒度的角色权限管理）被重新提出。
- **潜在重大升级**:
  - **Agent 协作总线**: PR `#2937` “引入内部 Agent 协作总线” 虽然当前状态为 OPEN，但其描述的功能（持久化邮箱、协作线程等）代表了 PicoClaw 架构演进的重要方向。如果此 PR 成功合并，将标志着 PicoClaw 从单一 Agent 向多 Agent 协作系统迈出一大步。

## 7. 用户反馈摘要

- **痛点与不满**:
  - **Windows 用户**: 核心文件操作 `list_dir` 在 Windows 上失败，严重影响了 Windows 作为开发或服务平台的可用性。
  - **高级用户**: 在为 PicoClaw 构建自定义 UI 时，认为控制流信号（如 Agent 是否完成思考）不够透明和确定，增加了开发复杂性。
  - **功能用户**: 在升级或尝试新模型（如 Gemini 3.5 Flash）时遇到兼容性问题，对框架的模型兼容性表示担忧。
- **正向反馈**:
  - 社区对于代码质量的关注度很高，多位贡献者提交了修复静默错误处理的 PR，这从侧面反映了代码库的健康度。
  - 自动依赖更新 PR 的迅速合并，让开发者对项目能及时同步外部安全更新和功能保持信心。

## 8. 待处理积压

以下 Issue/PR 已长时间未得到有效回应或解决，可能成为项目健康度的潜在风险，建议维护者给予关注：

- **Issue #2374**: [BUG] `.gitignore` 文件内容被覆盖 (`tools/` 目录) - 这是一个可能影响开发工作流的低级错误，自2025年10月后已无活动。
- **Issue #1135**: [Feature] 支持 RSS 订阅作为输入通道 - 这是一个被高度期望的功能请求，但由于长期未更新，社区可能已失去讨论热度。
- **PR #2950**: [Stale] fix: 飞书 (Feishu) 通道消息线程ID处理 - 该 PR 旨在修复一个具体的逻辑错误，但停留在一个月前的讨论状态，需要维护者介入审核或推动。

---

**总结**：PicoClaw 项目正处于活跃的迭代期，社区力量充沛，但同时也面临着跨平台兼容性、新模型适配等快速成长中的挑战。建议项目维护者在追求新功能的同时，优先解决报告中标记为“高”和“中高”的 Bug，特别是与核心功能（如Windows兼容性）和模型兼容性相关的问题。同时，对于社区呼声高的协议级增强（如完成信号），应尽快给出明确的规划或响应，以维持社区贡献热情。

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，我已根据您提供的 GitHub 数据，为您生成了 NanoClaw 项目 2026 年 6 月 12 日的项目动态日报。

---

### NanoClaw 项目动态日报 | 2026-06-12

---

#### 1. 今日速览

今日 NanoClaw 项目活跃度极高，呈现出明显的“高产”状态，主要贡献者集中在核心团队。过去24小时内，共有 13 个 PR 被更新，其中 9 个已合并或关闭，合并/关闭率高达 69%，这通常反映了高效的开发流程和明确的版本迭代目标。尽管仅关闭了 1 个 Issue，但新开的 Issue 聚焦于核心架构的“内存系统重新设计”，表明项目正从功能修补向深层次架构重构迈进。综合来看，项目目前处于 **高活跃度、高质量产出** 的健康开发周期。

---

#### 2. 版本发布

无新版本发布。

---

#### 3. 项目进展

今日项目完成了 9 个 PR 的合并与关闭，主要集中在 Bug 修复和核心功能增强上。贡献者 `gavrielc` 表现突出，提交并合并了多个关键修复与功能。具体进展如下：

- **核心 Bug 修复：**
    - **修复 Session Manager 关键数据丢失问题**: **PR #2738** 修复了 `writeOutboundDirect` 函数以只读模式打开数据库，导致命令门控（command-gate）的拒绝响应被静默丢弃的严重 Bug。此修复是保证通信可靠性的关键一步。
    - **修复容器调度稳定性**: **PR #2736** 为刚唤醒的容器（可能含有陈旧的 `processing claims`）引入了宽限期，防止 Host Sweep 机制误判并回收这些容器，提升了系统的整体健壮性。
    - **优化用户交互体验**: **PR #2735** 修复了 Chat-SDK 桥接中，已解决的审批卡片（Approval Card）无法记录执行用户的 Bug，解决了审计和责任追溯的痛点。
    - **修复 CLI 功能旁路**: **PR #2743** (待合并) 指出了 `ncl wirings create` 命令会绕过必要的 `agent_destinations` 副作用，导致新创建的聊天中，Agent 发送的消息被丢弃的问题。

- **新功能与架构增强：**
    - **多机器人基础架构**: **PR #2733** 引入了原生通道实例维度（Channel-instance dimension），为多机器人的部署（Multi-bot substrate）提供了基础支持，是平台扩展性的重要一步。
    - **Webhook 扩展能力**: **PR #2739** 添加了“原始路由注册表”，允许非 Chat-SDK 的 Webhook 作为附加组件被集成，降低了第三方系统集成的门槛。
    - **审批流程现察性**: **PR #2737** 实现了审批解决后的回调注册机制，允许不同模块以“加法”方式监听和响应审批结果，增强了系统的模块化和扩展性。
    - **消息投递增强**: **PR #2734** 为 Action 注册表添加了 `getDeliveryAction` 读取接口，完善了消息投递的查询能力。
    - **容器生命周期管理**: **PR #2740** 为容器添加了“按组空闲超时”功能，支持临时会话（Ephemeral Session）的自动清理退出。
    - **交互流程优化**: **PR #2741** 修复了交互式设置流程中，向 Claude 传递上下文时因缺少用户消息而无法自动处理的问题，提升了一体化体验。

**总结:** 项目今日在 **稳定性、扩展性、用户交互** 三个方面均取得了实质性的进展，尤其是对多个潜在数据丢失和流程阻断 Bug 的快速修复，展现了团队对产品质量的严格要求。

---

#### 4. 社区热点

在今日活跃的议题中，讨论焦点主要集中在核心架构和体验问题上。

- **最受关注 Issue**: **#1356 [OPEN] Agent memory system redesign**
  - **链接**: [nanocoai/nanoclaw Issue #1356](https://github.com/qwibitai/nanoclaw/issues/1356)
  - **分析**: 该 Issue 获得了 6 个👍，是目前社区最关注的核心顶层设计议题。虽然创建于3月，但昨日仍在更新，显示出其重要性。社区对当前基于文件的内存系统在扩展性上的瓶颈表示担忧，并积极寻求一个更全面、可扩展的解决方案。这代表了随着项目用户量和使用深度增加，对 **“记忆与持久化”** 这一核心技术能力的需求正在急剧上升。

- **高价值 PR**: **PR #2744 [OPEN] fix(signal): deliver agent reactions and forward inbound reactions**
  - **链接**: [nanocoai/nanoclaw PR #2744](https://github.com/qwibitai/nanoclaw/pull/2744)
  - **分析**: 尽管没有直接评论，但该 PR 直接关联到Signal适配器的功能完整性。它指出Agent的 `add_reaction` 工具输出被静默丢弃，且入站反应（reaction）被完全忽略。这对于依赖 Signal 进行生产通信的用户来说是一个重大体验缺失，表明 **“渠道适配器功能对等性”** 是用户的重要诉求。

---

#### 5. Bug 与稳定性

今日报告的 Bug 和修复主要集中在功能性逻辑错误和数据一致性问题，未发现致命的崩溃或安全漏洞。

- **严重 - Bug 已修复**: **Session Manager 只读数据库写入失败 (Issue #2495)**
  - **状态**: 已关闭
  - **修复 PR**: **PR #2738** (已合并)
  - **描述**: `writeOutboundDirect()` 函数在写入时错误地以只读模式打开了 `outbound.db`，导致涉及命令门控的拒绝响应等关键数据写入失败，从而被静默丢弃。这是一个可能导致 **关键通信丢失** 的严重逻辑错误。

- **严重 - Bug 已修复**: **Host Sweep 对刚唤醒容器的误判**
  - **状态**: 已关闭
  - **修复 PR**: **PR #2736** (已合并)
  - **描述**: 容器从休眠中唤醒后，Sweep 机制可能会立即将其标记为“处理陈旧 claim”并回收，导致容器刚启动就立刻被杀死，影响服务稳定性。

- **中等 - Bug 已修复**: **审批卡片无法记录操作人 (Issue #2495 相关)**
  - **状态**: 已关闭
  - **修复 PR**: **PR #2735** (已合并)
  - **描述**: 审批流程中，谁批准或拒绝了请求无法被记录，导致 **追溯和审计困难**。

- **中等 - Bug 已被定位，有修复 PR**: **CLI `wirings create` 命令绕过副作用 (PR #2743)**
  - **状态**: 待合并
  - **影响**: 使用 CLI 创建线路（wiring）时，Agent 可能无法正常向新创建的聊天发送消息，是一个功能性缺陷。

- **中等 - Bug 已被定位，有修复 PR**: **Signal 适配器反应功能缺失 (PR #2744)**
  - **状态**: 待合并
  - **影响**: Signal 渠道上的 Agent 和用户均无法使用消息反应功能。

---

#### 6. 功能请求与路线图信号

从今日的数据中，可以提炼出以下对未来路线图的指引信号：

- **顶层架构重构信号**: **Issue #1356 (内存系统重新设计)** 是最明确的路线图信号。这表明，随着项目规模的扩大，当前单一的内存架构正在成为瓶颈。后续版本极有可能围绕 **更可扩展、更强大的记忆系统** 进行重大投入。
- **多实例与集群化**: **PR #2733 (原生通道实例维度 - 多机器人基板)** 的合并是重大架构信号。这表明项目正在为 **支持单一 Agent 的多个实例** 或 **多 Agent 协同工作** 的集群化部署铺平道路，不再局限于单机单 Agent 模式。
- **自动化与 DevOps 集成**: **PR #2742 (PR Factory - 一个关于 PR 审查、分类和测试的发布食谱)** 是一个颇具特色的功能。它将项目自身的开发流程自动化，利用 Agent 来辅助代码审查、测试和发布管理，这展示了项目在 **“用 AI 管理 AI 项目”** 这一前沿实践上的探索。
- **开发体验 (DX) 提升**: **PR #2741 (自动提交上下文给 Claude)** 和 **PR #2739 (原始路由注册表)** 的修复和功能，表明项目持续关注并优化着开发者/运维人员的上手和使用体验，降低集成和故障排查的摩擦。

---

#### 7. 用户反馈摘要

- **对于内存系统的反馈** (来自 Issue #1356): 用户 `Ordinath` 通过 Issue 指出，当前内存系统在小规模运行时表现尚可，但随着 Agent 文件数量和大小增长（~54个文件，~83KB），其扩展性限制已经显现。这反映了 **社区用户面临真实可扩展性瓶颈**，他们对一个“更全面、可扩展”的方案有强烈的迫切需求。
- **关于数据静默丢失的痛点** (来自 Issue #2495): 用户 `SebTardif` 发现并报告了 `writeOutboundDirect` 的 Bug。这个 Bug 的可怕之处在于它是**静默失败**的——写入失败的异常被静默捕获，导致用户误以为指令已发送（命令门控的“拒绝”响应），但实际上并未送达。这暴露了 **在生产环境中，命令的无提示丢失** 是对用户信任的巨大打击。

---

#### 8. 待处理积压

- **Issue #1356: Agent memory system redesign** (创建于 2026-03-23, 更新于 2026-06-11):
  - **链接**: [nanocoai/nanoclaw Issue #1356](https://github.com/qwibitai/nanoclaw/issues/1356)
  - **原因**: 该议题已存在近3个月，但每次更新都表明社区和核心开发者仍在持续关注和思考。这既是一个 **需要长期投入的技术债务**，也是 **路线图规划的关键信号**。维护者可以考虑发布一份初步的设计文档或 RFC，与社区进行更深入的探讨，而不是停留在 Issue 层面。
- **PR #2685: docs(signal): group typing, outbound reactions, quote-reply fix** (创建于 2026-06-04, 更新于 2026-06-11):
  - **链接**: [nanocoai/nanoclaw PR #2685](https://github.com/qwibitai/nanoclaw/pull/2685)
  - **原因**: 此 PR 是关于 Signal 渠道新功能（群组打字、出站反应、引用回复）的文档更新，但已存在超过一周未被合并。鉴于当前已有多个 Signal 相关的功能 PR（如 #2744），此文档 PR 的长期未合并可能会造成代码与文档的脱节，建议维护者优先合并或关闭。

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw 项目动态日报 — 2026-06-12

## 1. 今日速览
- 过去 24 小时内项目活跃度较低，仅收到 **1 条新 Issue**，无新 Pull Request 或版本发布。
- 该 Issue 报告了使用 Ollama 加载本地模型时，Agent 回答不完整的问题，目前尚无社区讨论或修复进展。
- 整体来看，项目处于低速运维状态，核心开发活动接近停滞，社区对模型兼容性的关注度有所上升。

## 2. 版本发布
今日无新版本发布。

## 3. 项目进展
今日无 PR 被合并或关闭，项目主干无实质推进。

## 4. 社区热点
- **Issue #952**（[链接](https://github.com/nullclaw/nullclaw/issues/952)）：用户 `bloodgroup-cplusplus` 报告了使用 `ollama` 拉取 `gemma` 模型后，Agent 的回答总是截断、不完整的现象。该 Issue 为今日唯一活跃话题，虽无人评论，但暴露了本地模型集成的一个关键交互缺陷，可能影响所有通过 Ollama 使用非流式或短响应模型的用户。

## 5. Bug 与稳定性
| 严重程度 | Issue | 状态 | 是否已有 fix PR |
|---------|-------|------|----------------|
| **高** | [#952] 本地模型回答不完整 | OPEN | 无 |

- 该 Bug 导致 Agent 无法正常输出完整句子，严重影响使用体验。用户提供了截图（未在文本中复现），表明现象可复现。目前无关联 PR，需维护者定位是否为 Ollama 响应截断或 Agent 输出处理逻辑问题。

## 6. 功能请求与路线图信号
今日无新增功能请求。从 Issue #952 看，用户隐含需求是：
- 确保 Agent 对 Ollama 返回的流式/非流式输出能正确拼接和截断；
- 提升与主流本地模型（如 Gemma）的兼容性。

此类问题若持续涌现，可能推动项目增加 `response_completion_check` 机制或配置项，允许用户设定最大等待/重试次数。

## 7. 用户反馈摘要
- **用户痛点**：`bloodgroup-cplusplus` 描述了完整操作步骤（pull → start → 回答不全），表明该流程不具备预期可用性。用户未表达愤怒，但上传截图说明其试图复现并寻求帮助。
- **使用场景**：本地部署 AI Agent，依赖 Ollama 作为模型后端。
- **不满意点**：Agent 输出不完整，无法直接用于生产或测试。暂无正面反馈。

## 8. 待处理积压
目前无长期未响应的重要 Issue。建议维护者尽快跟进 #952，因为该 Bug 涉及本地模型基本可用性，且用户已给出明确复现信息。如能快速修复或提供临时工作区（如设置 `max_tokens` 或启用流式模式），将有助于提升社区信心。

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，我将根据您提供的 IronClaw 项目数据，为您生成 2026 年 6 月 12 日的项目动态日报。

---

### **IronClaw 项目动态日报 | 2026-06-12**

**数据快照：** 过去24小时，项目共处理 19 个 Issues 和 40 个 PR，其中新开/活跃 Issue 16 个，待合并 PR 22 个。

---

#### **1. 今日速览**

今日 IronClaw 项目状态呈现 **高活跃度**，开发节奏紧凑。尽管没有新版本发布，但项目核心开发团队与社区成员针对 **Reborn** 版本的稳定性和体验优化展开了高强度工作。一方面，大量由社区成员（尤其是一位 ID 为 `sunglow666` 的用户）提交的高质量 Bug 报告被快速定位，并已有多个高优先级的修复 PR 提交；另一方面，项目也在积极推进**配置即代码**、**自动化 QA**、**开发者可观测性**等基础设施建设工作。整体来看，项目正处在 **Reborn 全面上线前关键的加固与冲刺阶段**。

---

#### **2. 版本发布**

过去24小时无新版本发布。

---

#### **3. 项目进展**

今日项目推进迅速，多个关键性 PR 被合并或关闭，标志着以下方向取得重要进展：

- **Reborn 生产就绪度提升：** PR [#4763](https://github.com/nearai/ironclaw/pull/4763) 的合并，关闭了关于生产环境切换就绪度的 Issue [#4621](https://github.com/nearai/ironclaw/issues/4621)。该 PR 补充了最后一份切换就绪的地图，并明确了向 Reborn 运行时迁移的默认关闭和回滚策略，是 Reborn 走向生产环境的里程碑式进展。
- **WebUI 核心功能与稳定性修复：**
    - **日志系统：** PR [#4760](https://github.com/nearai/ironclaw/pull/4760) 已合并，解决了 Issue [#4758](https://github.com/nearai/ironclaw/issues/4758) 中 WebUI 日志页面空白的问题，成功将真实的运行日志接入前端，并带有缓存、过滤、密钥脱敏等功能。
    - **UI Bug 修复：** PR [#4772](https://github.com/nearai/ironclaw/pull/4772) 合并，批量修复了 WebChat v2 的一系列前端问题，包括模型选择器 Bug、代码块换行、自动滚动等。
    - **自动化流程打通：** PR [#4757](https://github.com/nearai/ironclaw/pull/4757) 合并，修复了自动化触发任务点开后页面空白的问题，成功打通了“自动化 -> 查看/审批执行”的端到端流程。
- **开发者体验与生态：**
    - **能力可观测性：** 核心 PR [#4588](https://github.com/nearai/ironclaw/pull/4588) 依然开放，但已获得持续关注。该 PR 引入了轨迹观察者和 LLM 提供者注入两个关键钩子，为外部程序（如基准测试工具）驱动和观察 Reborn Agent 提供了标准接口，是提升框架可扩展性的重要一步。
    - **CI 流程优化：** PR [#4774](https://github.com/nearai/ironclaw/pull/4774) 被提交，提议引入 CodeRabbit 作为新的 AI 代码审查工具，与现有审查者进行对比评估，以提高代码审查效率。

---

#### **4. 社区热点**

今日社区活动主要由一个活跃的测试者 `sunglow666` 驱动，其发现的多个稳定性和可用性问题引发了核心团队的高度关注和快速响应，形成了 **“发现-分析-修复”** 的快速循环。最受关注的热点包括：

- **多工具调用时 Agent 循环崩溃：** Issue [#4789](https://github.com/nearai/ironclaw/issues/4789) 报告了一个严重问题：当模型一次返回多个工具调用时，Agent 会因 `driver_unavailable` 错误彻底崩溃。此 Bug 阻塞了 Agent 执行复杂的多步骤任务。
    - **关联修复：** 核心成员 `henrypark133` 在一个小时内就提交了修复 PR [#4790](https://github.com/nearai/ironclaw/pull/4790)，分析并解决了底层连接池的竞争条件问题，响应极为迅速。
- **Agents 交互死锁与不一致：** 用户测试了 Agent 在“工具失败后”的交互行为，提出了两个深度相关的问题：
    - Issue [#4761](https://github.com/nearai/ironclaw/issues/4761) 指出，在多次工具调用失败后，Agent 会陷入无法恢复的“死锁”状态。
    - Issue [#4762](https://github.com/nearai/ironclaw/issues/4762) 则发现，失败的工具体验会导致后续对话消息顺序错乱。
- **日常操作的可用性反馈：** `sunglow666` 还报告了一系列影响基本操作的 Bug，如工作区路径重复（[#4759](https://github.com/nearai/ironclaw/issues/4759)）、生成文件在 WebUI 不可见（[#4750](https://github.com/nearai/ironclaw/issues/4750)）、拒绝 Shell 审批无反馈（[#4764](https://github.com/nearai/ironclaw/issues/4764)）等，这些反馈直接关联到 Reborn 版本提供给最终用户的“第一印象”。

---

#### **5. Bug 与稳定性**

今日 Bug 报告数量较多，按严重程度排列如下：

| 严重程度 | Issue | 描述 | 状态 |
| :--- | :--- | :--- | :--- |
| **严重** | [#4789](https://github.com/nearai/ironclaw/issues/4789) | **Agent 在批量工具调用时崩溃** (`driver_unavailable`)，导致 Agent 循环完全中断。 | 已有修复 PR [#4790](https://github.com/nearai/ironclaw/pull/4790) |
| **严重** | [#4761](https://github.com/nearai/ironclaw/issues/4761) | **Agent 在重复工具失败后停止恢复**，陷入死锁，无法完成后续正常任务。 | 待修复 |
| **严重** | [#4762](https://github.com/nearai/ironclaw/issues/4762) | **失败的工具操作导致后续对话消息与活动顺序不一致**，破坏对话体验。 | 待修复 |
| **高** | [#4759](https://github.com/nearai/ironclaw/issues/4759) | **工作区路径重复**：使用工作区相对路径创建文件时，路径被错误拼接。 | 待修复 |
| **中** | [#4750](https://github.com/nearai/ironclaw/issues/4750) | **WebUI 中无法发现或预览工作区文件**，用户不知道文件是否成功创建。 | 待修复 |
| **中** | [#4751](https://github.com/nearai/ironclaw/issues/4751) | **大型响应导致 Provider 工具参数超限** (16384 bytes)，限制 Agent 处理长内容。 | 待修复 |
| **低** | [#4764](https://github.com/nearai/ironclaw/issues/4764) | **拒绝 Shell 执行后无反馈**，用户界面上操作被挂起，没有成功或失败提示。 | 待修复 |

此外，**持续 E2E 测试失败** Issue [#4108](https://github.com/nearai/ironclaw/issues/4108) (已存在2周以上) 仍为 Open 状态，表明存在需要关注的长期稳定性问题。

---

#### **6. 功能请求与路线图信号**

今日的新功能请求显示了社区对未来功能的期待，并结合已有 PR 分析其潜在走向：

- **核心：配置即代码 (Configuration-as-Code)：** 史诗级 Issue [#3036](https://github.com/nearai/ironclaw/issues/3036) 今日获得更新。该提议希望引入一套基于 Schema 的声明式配置系统来管理 “Blueprints” 和 “Harnesses”，取代目前零散的 `.env` 和配置文件。这符合大型项目和云平台对配置可管理性、版本化和审计的需求，一旦实现，将极大提升 IronClaw 的企业级部署能力。
- **可用性提升：全局“始终允许”设置：** Issue [#4776](https://github.com/nearai/ironclaw/issues/4776) 提出增加一个全局开关，允许用户不必每次都为安全工具授权。这是从社区测试反馈中提炼出的直接需求，表明用户在执行简单任务时对繁琐的审批流程感到困扰。
- **质量保障：自动化 QA 大坝：** Issue [#4775](https://github.com/nearai/ironclaw/issues/4775) 开启了一个新的史诗任务，目标是为 Reborn 二进制文件建立从“无外部依赖的封闭测试”到“模拟 E2E 测试”再到“真实集成测试”的完整自动化 QA 体系。这标志着项目正在系统性地构建一个质量护城河，为后续正式版本的稳定性提供保障。已有相关 PR [#4769](https://github.com/nearai/ironclaw/pull/4769) 部分实现了此目标。
- **开发者可观测性：** PR [#4588](https://github.com/nearai/ironclaw/pull/4588) 的长期存在和不断更新，证明社区和开发者对能够“白盒”观察和驱动 Agent 运行的需求强烈，这将是 Reborn 生态发展的关键基础设施。

---

#### **7. 用户反馈摘要**

从今日的 Issue 评论中，可以提炼出以下用户痛点与使用场景：

- **核心痛点：Reborn 版本的稳定性。** 用户 `sunglow666` 提交的大量 Bug 暴露了 Reborn 在任务执行、UI 交互、工具状态管理等多个方面存在稳定性缺口。尤其是在处理“工具执行失败”这一边缘场景时，用户体验下降非常明显。这表明 Reborn 在鲁棒性方面还有很大的提升空间。
- **对“所见即所得”的期望。** 从 [#4750](https://github.com/nearai/ironclaw/issues/4750) 和 [#4764](https://github.com/nearai/ironclaw/issues/4764) 中可以看出，用户对系统状态有极强的可视性需求。当 Agent 声称“文件已创建”但 WebUI 看不到，或者拒绝一个操作后界面无反应时，用户会感到困惑和不安。
- **配置的复杂性是掣肘。** [#4711](https://github.com/nearai/ironclaw/issues/4711)(关于ChatGPT订阅流程困惑) 和 [#4776](https://github.com/nearai/ironclaw/issues/4776) (关于全局允许) 都指向一个共同问题：**当前配置和授权流程对用户不够友好**。用户希望在导入复杂的外部服务（如 ChatGPT、Slack）时能有更清晰、更一体化的引导，同时在日常使用中能跳过重复的确认步骤。

---

#### **8. 待处理积压**

以下为需要项目维护者关注的重要待办事项：

- **高优先级：E2E 测试持续失败。** Issue [#4108](https://github.com/nearai/ironclaw/issues/4108) 自 5月27日起持续报告 nightly E2E 测试失败，涉及 `v2-engine` 等核心组件。此问题已开放超过2周，严重影响项目稳定性信心，应作为最高优先级排查。
- **长期议题：配置即代码。** Issue [#3036](https://github.com/nearai/ironclaw/issues/3036) 已开放超过1个月，作为一个里程碑级的功能请求，其设计方案的制定和原型开发进度可能影响了项目的整体规划，建议维护者能给出一个初步的时间表或设计决策。
- **长期搁置 PR：版本发布。** PR [#3708](https://github.com/nearai/ironclaw/pull/3708) 是一个自动化的版本发布 PR，距今已有近一个月，包含了多项 API breaking changes。其长期未合并意味着新功能无法通过版本升级交付给用户，建议团队评估其阻塞点并尽快处理。
- **新星 PR 观察：** PR [#4588](https://github.com/nearai/ironclaw/pull/4588) 作为提升可观测性的核心 PR，虽然每日有更新，但进展缓慢，建议维护者给予更多关注，推动其落地。

---

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

```markdown
# LobsterAI 项目动态日报 (2026-06-12)

## 1. 今日速览
过去 24 小时项目活跃度较高。共处理 2 条 Issue（1 条新开/活跃、1 条已关闭），19 条 Pull Request（13 条已合并/关闭、6 条待合并）。合并的 PR 集中在 Cowork 模块的模型选择、流停止元数据、实时 ASR 语音输入、共享文件等功能的修复与增强，同时对 OpenClaw 网关、技能注入、快捷键冲突等长期存在的 Bug 进行了修复。没有新版本发布，但大量修复 PR 的合入表明项目正在快速推进稳定性与功能完善。

## 2. 版本发布
无。

## 3. 项目进展
今日合并/关闭了 **13 条 PR**，其中以下变更对项目功能与稳定性有重要推进：

| PR # | 功能/修复 | 状态 |
|------|----------|------|
| [#2154](https://github.com/netease-youdao/LobsterAI/pull/2154) | 修复流停止后模型元数据不显示的问题 | 已合并 |
| [#2153](https://github.com/netease-youdao/LobsterAI/pull/2153) | 保留同名称包模型选择，区分包模型与自定义模型 | 已合并 |
| [#2152](https://github.com/netease-youdao/LobsterAI/pull/2152) | 延长发送前模型同步超时时间至 90s，避免冷启动丢消息 | 已合并 |
| [#2151](https://github.com/netease-youdao/LobsterAI/pull/2151) | 新增共享文件功能 | 已合并 |
| [#2150](https://github.com/netease-youdao/LobsterAI/pull/2150) | 修复 Kits 专家套件顶部栏固定问题 | 已合并 |
| [#2149](https://github.com/netease-youdao/LobsterAI/pull/2149) | 提升 OpenClaw 网关进程 V8 堆内存限制，减少 OOM 崩溃 | 已合并 |
| [#2148](https://github.com/netease-youdao/LobsterAI/pull/2148) | 新增实时 ASR 语音输入功能（WebSocket 流式识别） | 已合并 |
| [#2147](https://github.com/netease-youdao/LobsterAI/pull/2147) | 修复启动阶段用户主动停止时仍发送 Chat 的竞态问题 | 已合并 |
| [#1473](https://github.com/netease-youdao/LobsterAI/pull/1473) 等 5 条 | 为 Agent 创建、设置、MCP 配置等弹窗添加未保存确认，防止内容静默丢失 | 已合并 |

这些修复覆盖了用户近期反馈的多个痛点，如输入草稿丢失、快捷键冲突、定时任务创建无响应等。项目在 **稳定性和用户体验** 方面向前迈进了重要一步。

## 4. 社区热点
- **[Issue #2121 – 重复输出文字浪费 token 的疑问](https://github.com/netease-youdao/LobsterAI/issues/2121)**  
  用户提交截图，怀疑系统在重复输出内容导致大量 token 浪费，并询问是否是 claw 的问题。该 Issue 目前有 1 条评论，讨论未定论，但反映了用户对 token 消耗的敏感度。

- **[PR #1446 – 修复网关无限重启循环](https://github.com/netease-youdao/LobsterAI/pull/1446)**  
  尽管已停留多日，今天有更新（6 月 12 日），可能是维护者重新评估。该 PR 关联 Issue #1400，涉及 OpenClaw 网关在 4.1 版本后的严重无限重启问题，是社区关注的稳定性重点。

## 5. Bug 与稳定性
| 严重等级 | Bug 描述 | 状态 |
|----------|----------|------|
| 🔴 严重 | OpenClaw 网关进程因竞态条件陷入无限重启循环，导致整个应用瘫痪 (PR [#1446](https://github.com/netease-youdao/LobsterAI/pull/1446)) | 有 fix PR，待合并 |
| 🟠 较高 | 已停用的技能仍被注入对话提示词，导致用户无法真正关闭技能 (PR [#1453](https://github.com/netease-youdao/LobsterAI/pull/1453)) | 有 fix PR，待合并 |
| 🟠 较高 | 定时任务选择「不重复」模式、清空日期后点击创建无反应，无任何错误提示 (PR [#1454](https://github.com/netease-youdao/LobsterAI/pull/1454)) | 有 fix PR，待合并 |
| 🟡 中等 | 快捷键设置缺少重复检测，多个功能可绑定同一组合导致部分快捷键静默失效 (PR [#1456](https://github.com/netease-youdao/LobsterAI/pull/1456)) | 有 fix PR，待合并 |
| 🟡 中等 | 创建 Agent、设置面板、MCP 配置等弹窗关闭时未提示未保存更改，导致内容静默丢失 (PR [#1473-#1477](https://github.com/netease-youdao/LobsterAI/pulls?q=is%3Apr+label%3Afix+created%3A%3E%3D2026-04-04)) | 已合并修复 |
| 🟢 低 | 流停止后模型元数据不显示、启动阶段停止仍发 Chat 等 (PR #2154, #2147) | 已合并修复 |

今日新增 Bug 报告：**[Issue #2121](https://github.com/netease-youdao/LobsterAI/issues/2121)** 疑似重复输出文本浪费 token，尚未有 fix PR。

## 6. 功能请求与路线图信号
- **实时 ASR 语音输入**（[#2148](https://github.com/netease-youdao/LobsterAI/pull/2148)）今日已合并，预计将在下一版本中提供流式语音输入选项，用户可在设置中切换实时 / 一次性录入模式。
- **文件共享功能**（[#2151](https://github.com/netease-youdao/LobsterAI/pull/2151)）已合并，为 Cowork 模块增加文件分享能力。
- **定时任务执行记录折叠分组**（[#1449](https://github.com/netease-youdao/LobsterAI/pull/1449)）虽仍待合并，但作为 UX 优化持续被关注，有望在后续迭代中减少侧栏历史记录堆积。

用户反馈中（Issue #2121）对 token 浪费的担忧可能暗示社区对**成本透明化**的需求，建议路线图中考虑加入提示词/输出长度审计功能。

## 7. 用户反馈摘要
- **simson2010** 在 [Issue #1](https://github.com/netease-youdao/LobsterAI/issues/1) 中报告：在 Mac OS 13.7.8 上配置 MiniMaxi API 测试通过，但切换到 OpenAI 消息类型时遇到 “invalid params” 错误，该问题已于今日已关闭（推测已在早期版本修复或用户自行解决）。
- **nbjoe** 在 [Issue #2121](https://github.com/netease-youdao/LobsterAI/issues/2121) 中提出：`claw` 在对话中重复输出文字，怀疑大量消耗 token，影响使用体验。目前仅有 1 条评论，未得到官方回应。

用户对 **输入内容安全** 有较高期待：多个 PR 针对未保存确认的修复（#1473-#1477）直接回应了用户“不小心关闭导致内容丢失”的痛点。

## 8. 待处理积压
以下 **6 条 PR** 创建于 4 月 3 日，标记为 `stale`，但今日（6 月 12 日）有更新活动，可能正被重新审查或等待最终合并。建议维护者优先关注：

| PR # | 描述 | 链接 |
|------|------|------|
| #1446 | 修复 OpenClaw 网关无限重启循环 | [链接](https://github.com/netease-youdao/LobsterAI/pull/1446) |
| #1448 | Agent 设置页删除按钮及技能选择器未国际化 | [链接](https://github.com/netease-youdao/LobsterAI/pull/1448) |
| #1449 | 定时任务多次执行记录折叠分组展示 | [链接](https://github.com/netease-youdao/LobsterAI/pull/1449) |
| #1453 | 已停用技能仍被注入对话提示词 | [链接](https://github.com/netease-youdao/LobsterAI/pull/1453) |
| #1454 | 不重复模式清空日期后创建任务无响应 | [链接](https://github.com/netease-youdao/LobsterAI/pull/1454) |
| #1456 | 快捷键设置缺少重复检测 | [链接](https://github.com/netease-youdao/LobsterAI/pull/1456) |

另外，**[Issue #2121](https://github.com/netease-youdao/LobsterAI/issues/2121)** 尚无官方回复，如确认为 Bug 建议尽快分配。
```

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis 项目动态日报 (2026-06-12)

---

## 1. 今日速览

过去24小时项目整体活跃度偏低，仅有1个新Issue和1个待合并PR产生，无新版本发布，也无任何已关闭或合并的事件。社区焦点集中在两个独立的修复与反馈上：一个涉及Fastmail MCP授权问题的Bug报告，另一个是针对WhatsApp“@lid”聊天回复投递失败的修复PR。项目当前处于小幅迭代状态，未见大规模功能推进或紧急回归问题。

---

## 2. 版本发布

无新版本发布。

---

## 3. 项目进展

**今日无任何PR被合并或关闭**，项目停滞在“待合并”队列等待代码审查。唯一活跃的PR #1116 由 juanlotito 提交，旨在修复 WhatsApp 渠道中通过 PN JID 重写机制来投递对“@lid”聊天的回复。该问题表现为：网关正确生成了回复（Web UI可看到），但通过 outbound 发送后消息从未到达用户，也未收到送达回执。修复后有望改善隐私模式下的消息交付可靠性。该PR处于开放状态，尚未被合并。

> 🔗 相关链接：[PR #1116](https://github.com/moltis-org/moltis/pull/1116)

---

## 4. 社区热点

今日最受关注的议题是 **Issue #1115** —— 用户 **kmath313** 报告在使用 Fastmail MCP（Model Context Protocol）时遇到授权问题。该Issue创建于6月11日，已有1条评论，目前仍保持开放状态，且未获得任何 👍 标记。虽然讨论热度不高，但这是过去24小时内唯一的新Issue，且涉及外部邮件服务集成，可能对依赖 Fastmail 的用户构成使用障碍。

> 🔗 相关链接：[Issue #1115](https://github.com/moltis-org/moltis/issues/1115)

---

## 5. Bug 与稳定性

唯一新报告的Bug（严重等级：中等）：

- **#1115 [Bug] Fastmail MCP Authorisation**  
  用户已按模板提交预检清单（确认已搜索过现有Issue、使用最新版本），但未提供完整会话上下文。问题归因于Fastmail MCP授权流程出错。目前无关联的Fix PR，也未见临时解决方案。

项目未出现崩溃、回归或严重数据丢失类Bug。

---

## 6. 功能请求与路线图信号

今日无新功能请求提出。PR #1116 属于Bug修复而非新功能，但改善WhatsApp消息投递的可靠性可能对后续隐私增强路线有参考价值。暂无迹象表明该修复会被纳入下一版本。

---

## 7. 用户反馈摘要

从唯一Issue #1115的摘要中可提取出用户痛点：  
- 用户 **kmath313** 在配置Fastmail MCP授权时受阻，尽管已确认使用最新版本并搜索过重复问题。  
- 用户遵循了模板预检流程，但问题依然存在，说明当前文档或错误处理可能不够清晰。

由于无更多评论，无法进一步了解用户满意度或具体使用场景。

---

## 8. 待处理积压

截至今日，项目中暂无超过数日未响应的长期积压Issue或PR。唯一活跃的Issue #1115 和 PR #1116 均为昨日或今日创建，正处于正常响应周期内，无需特别提醒维护者。

| 类型 | 编号 | 标题 | 创建日 | 状态 | 链接 |
|------|------|------|--------|------|------|
| Issue | #1115 | Fastmail MCP Authorisation | 2026-06-11 | 开放 | [链接](https://github.com/moltis-org/moltis/issues/1115) |
| PR | #1116 | fix(whatsapp): deliver replies to @lid chats via PN JID rewrite | 2026-06-12 | 开放 (待合并) | [链接](https://github.com/moltis-org/moltis/pull/1116) |

---

**日报生成分析结论**：项目今日整体节奏平稳，无重大进展或风险。建议维护者优先审查PR #1116 的代码质量，并尽快回应Issue #1115 以降低用户等待时间。若Fastmail MCP授权问题影响面扩大，应评估是否需要立即发布包含临时修复的补丁版本。

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw 项目动态日报 | 2026-06-12

---

## 1. 今日速览

CoPaw（QwenPaw）过去24小时保持了高产且高质量的社区活跃度：共处理 **23 条 Issue**（13 条新开/活跃，10 条关闭）和 **37 条 Pull Request**（21 条待合并，16 条已合并/关闭），并发布了 2 个补丁版本。项目正处在**后端重大重构（AgentScope 1.x → 2.0 迁移）与前端稳定性修复并重**的阶段，社区反馈集中于**桌面端启动失败、进程内存泄漏、记忆配置丢失**等可用性问题，同时也有多项功能增强 PR 进入审查。整体项目健康度良好，维护者响应及时，但需关注累积的积压问题。

---

## 2. 版本发布

过去 24 小时内发布了两个补丁版本，均为 v1.1.11 系列的后缀修正：

- **v1.1.11.post2**  
  - 变更：修整工具卡片标题样式（单行省略）、版本号提升。  
  - 破坏性变更：无。  
  - 迁移注意事项：建议所有桌面端用户升级以修复已知异常。

- **v1.1.11.post1**  
  - 变更：回退了一个可能导致打包后 Discord 检查失败的修复，补充发布职责检查清单。  
  - 破坏性变更：无。

📎 [v1.1.11.post2 Release](https://github.com/agentscope-ai/QwenPaw/releases/tag/v1.1.11.post2)  
📎 [v1.1.11.post1 Release](https://github.com/agentscope-ai/QwenPaw/releases/tag/v1.1.11.post1)

---

## 3. 项目进展

今日合并或关闭了多项重要 PR，标志着关键修复与功能推进：

- **安全加固**：`fix(security): isolate keychain master key per install`（PR #5028）已合并，解决了多安装实例共享同一 keychain 项的安全隐患。  
- **CI 稳定性**：`fix(desktop): harden Tauri Windows CI against crates.io fetch failures`（PR #5125）已合并，加入重试与离线构建，减少 CI 不稳定。  
- **UI 增强**：`feat(ui): apply AionUi design language to Console layout`（PR #5133）和 `feat(i18n): pt-BR translation`（PR #5136）今日由首次贡献者提交并合并，丰富了用户界面与国际化支持。  
- **自动化流程**：`feat(.claude): qwenpaw-changelog historian agent`（PR #5134）合并，新增自动记录变更日志的 agent，提升开发团队效率。  
- **插件系统**：`Decouple plugin loader initialization from agent startup`（PR #4900）仍在审查中，但已包含对 frozen 环境（PyInstaller/Tauri）的关键修复。

总体上，项目在**安全性、国际化、可观测性**方面有明显推进。

---

## 4. 社区热点

| Issue/PR | 类型 | 评论数 | 核心诉求 |
|----------|------|--------|----------|
| [#4727](https://github.com/agentscope-ai/QwenPaw/issues/4727) [OPEN] | Issue | 9 | **后端迁移至 AgentScope 2.0**：社区高度关注这一重大重构，讨论了迁移时间线与兼容性影响。 |
| [#5064](https://github.com/agentscope-ai/QwenPaw/issues/5064) [OPEN] | Issue | 8 | **Agent 创建的定时任务无法触发**：用户反馈任务创建后无报错但到点不执行，且无法手动编辑，涉及核心功能缺陷。 |
| [#5106](https://github.com/agentscope-ai/QwenPaw/issues/5106) [CLOSED] | Issue | 7 | **桌面端 SSL 证书错误 + 无限进程占满内存**：Windows 用户遭遇严重可用性问题，已在 v1.1.11.post2 中部分解决但仍需关注。 |
| [#5137](https://github.com/agentscope-ai/QwenPaw/issues/5137) [OPEN] | Issue | 4 | **记忆搜索配置保存丢失**：UI 交互细节导致用户配置无法持久化，修复 PR #5144 已提交。 |
| [#5139](https://github.com/agentscope-ai/QwenPaw/issues/5139) [OPEN] | Feature | 2 | **Agent 团队/群组协作能力**：用户希望引入类似 WorkBuddy 的多 agent 协作模式，标志着项目功能需求方向。 |

**分析**：社区对 **稳定性与可用性** 的诉求最为强烈（SSL、进程泄漏、定时任务、配置保存），同时用户也对 **新功能（Agent 团队协作、token 统计）** 抱有期待，说明项目正处于“稳定现有功能”与“扩展新能力”的平衡期。

---

## 5. Bug 与稳定性

下列 Bug 按严重程度排列，标注是否已有修复 PR：

1. **🔴 致命**：`[Bug]: Windows客户端进程持续增加`（#5138）——打开客户端后进程无限增长，内存占用 >90%。**修复状态**：暂无关联 PR，严重性最高。  
2. **🔴 致命**：`[Bug]: OpenSSL 3.5 回归 bug 导致 Desktop 无法启动`（#5086 已关闭）——该 Bug 已在 v1.1.11.post1/post2 中修复，但需持续监控是否复发。  
3. **🔴 严重**：`[Bug]: 新版Tauri端SSL证书错误+无限进程占满内存致黑屏`（#5106 已关闭）——关联 PR 已合并，但仍需跟进用户实际升级后的反馈。  
4. **🟠 中等**：`[Bug]: 向量模型自动记忆搜索配置丢失`（#5137）——保存前需展开折叠面板。**修复 PR**：[#5144](https://github.com/agentscope-ai/QwenPaw/pull/5144) 已提交（`forceRender: true`）。  
5. **🟠 中等**：`[Bug]: 附件下载报错404（docx/pdf等）`（#5140）——纯文本可下载，docx/pdf 返回 404。**修复状态**：暂无关联 PR。  
6. **🟡 一般**：`[Bug]: 编码模式刷新后 Session 丢失`（#5142）——修复 PR：[#5147](https://github.com/agentscope-ai/QwenPaw/pull/5147) 已提交。  
7. **🟡 一般**：`[Bug]: 上下文压缩统计值与实际 API 输入体量不符`（#5122）——怀疑技能/MCP 导致膨胀。**修复状态**：暂无 PR，需进一步排查。

另外，`[Bug]: 记忆搜索工具无结果显示 bug`（#5098）和 `[Bug]: 网页UI渲染数学公式根号显示问题`（#5143）均为 UI 渲染问题，已有部分修复。

---

## 6. 功能请求与路线图信号

| Issue | 描述 | 可能纳入下一版本的信号 |
|-------|------|-----------------------|
| [#5139](https://github.com/agentscope-ai/QwenPaw/issues/5139) | 原生 Agent 团队协作能力，类似 WorkBuddy Expert Team | 已获 2 个评论，且项目正在推进 [#5067](https://github.com/agentscope-ai/QwenPaw/pull/5067)（Agent OS Driver 统一抽象层），为多 Agent 协作打下基础。 |
| [#5103](https://github.com/agentscope-ai/QwenPaw/issues/5103) | 对话队列、token 统计、准确时间戳 | 已有对应 PR [#5130](https://github.com/agentscope-ai/QwenPaw/pull/5130)（per-turn token usage popover）进入审查，预计在下一版本中部分实现。 |
| [#5131](https://github.com/agentscope-ai/QwenPaw/issues/5131) | Coding 模式支持代码补全 | 开发者未回复，但项目已在改进 Coding 模式路由（PR [#5147](https://github.com/agentscope-ai/QwenPaw/pull/5147)），后续可能延伸。 |

**路线图信号**：项目正通过 `Agent OS Driver`（PR [#5067](https://github.com/agentscope-ai/QwenPaw/pull/5067)）统一 MCP/A2A/ACP 等外部协议，为多 Agent 协作等高级功能奠定架构基础；同时 token 统计功能（PR [#5130](https://github.com/agentscope-ai/QwenPaw/pull/5130)）直接呼应社区高频诉求。

---

## 7. 用户反馈摘要

从今日活跃的 Issues 评论中提炼以下真实用户痛点：

- **“新版 UI 更简洁，但附件无法下载，回退到 1.1.10 就正常。”**（#5102）  
- **“每次容器重启后向量模型配置被重置为空串，WebUI 保存不生效。”**（#3817）  
- **“新旧两个桌面版均无法使用，新版 SSL 错误加内存占满死机，旧版也打不开了。”**（#5106）  
- **“auto_memory_search 功能本身能用，但 UI 表格显示异常，file 列显示 unknown。”**（#5098）  
- **“配置了 `enable_thinking: false` 但对话依然显示 Thinking，是配置不对吗？”**（#5132）  
- **“希望像 openclaw 那样不等回复结束就能输入下一请求，形成队列；还要有 token 统计和准确时间。”**（#5103）  
- **“Coding 模式刷新后 session 丢失，普通 chat 正常。”**（#5142）  
- **“执行 compact 压缩后显示占用 0.9%，但实际上发给 API 的输入凭空多几十 KB，怀疑技能/MCP 导致膨胀。”**（#5122）

**满意点**：用户对新版 UI 的简洁性表示认可（#5102），但功能退化和 Bug 抵消了满意度。社区对维护者的快速响应（如 v1.1.11.post2 修复 SSL 问题）持正面态度。

---

## 8. 待处理积压

以下为长期未响应或解决方案进展缓慢的重要 Issue / PR，提醒维护者关注：

1. **#4727 [OPEN] 后端迁移至 AgentScope 2.0**  
   创建于 5 月 27 日，累计 9 条评论，涉及重大变更。迁移计划直接影响所有插件和功能开发，建议尽快发布 RFC 或里程碑计划。

2. **#4900 [OPEN] 解耦插件加载器初始化**（PR）  
   创建于 6 月 2 日，至今仍在审查。该 PR 修复了 frozen 环境中插件系统静默超时的问题，对桌面端用户至关重要，建议优先合并。

3. **#4975 [OPEN] 会话页面列顺序自定义**（PR）  
   创建于 6 月 5 日，无标签更新，无合并。功能较为积极但缺少维护者反馈，建议明确是否需要此项增强。

4. **#4622 [OPEN] DataPaw 数据分析插件**（PR）  
   包含 12 个 BI 技能的插件，创建于 5 月 22 日，代码量较大，但长期未获得审查。社区贡献者的首次贡献可能因等待过久而丧失积极性。

5. **#3817 [CLOSED] 容器重启后向量模型配置丢失**  
   Issue 已于 6 月 11 日关闭，但评论中仍有用户表示问题未完全解决（配置重置问题偶发）。建议维护者重新开启并跟踪根本原因——初始化覆盖逻辑。

---

*本日报由 AI 自动生成，基于 2026-06-12 的 GitHub 数据。所有链接指向 agentscope-ai/QwenPaw 仓库。*

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，这是根据您提供的 ZeroClaw 项目数据生成的 2026-06-12 项目动态日报。

---

# ZeroClaw 项目日报 — 2026-06-12

## 1. 今日速览

今日项目活跃度极高。v0.8.0 大版本已于近日发布，带来了多代理架构、安全策略隔离等重大更新，社区反馈热烈，但同时也导致新版本兼容性 Bug 报告集中涌现，尤其是 macOS 平台的问题。开发团队在积极处理新版本 Bug 的同时，也合并了此前积累的多个高优先级修复 PR，项目整体处于“发布后密集迭代与修复”阶段。主要动态包括：修复了子代理 CWD 继承、配置持久化等关键问题；社区用户围绕新特性和新 Bug 展开了积极讨论。

- **活跃度评估**: 🔥🔥🔥🔥🔥 (极高)

## 2. 版本发布

### [v0.8.0](https://github.com/zeroclaw-labs/zeroclaw/releases/tag/v0.8.0)

**这是 ZeroClaw 的一个里程碑式版本**，核心变更如下：

- **核心架构变革**: 引入多代理（Multi-Agent）架构。现在，单个守护进程（Daemon）可以运行多个具名代理（Agent），每个代理拥有独立的工作空间、内存、模型供应商配置、安全策略、通信渠道和个性。这为构建复杂的多智能体系统奠定了基础。
- **配置文件重写**: 采用全新的配置模式，以支持多代理架构。官方声称该版本支持从旧版配置文件自动迁移。
- **破坏性变更与迁移注意事项**:
    - 配置文件结构已完全改变。虽然官方声称支持自动迁移，但用户（特别是拥有复杂自定义配置的用户）手动检查并更新 `~/.zeroclaw/config.yaml` 是必要的。
    - 任何依赖旧版单实例 API 或目录结构的第三方工具或脚本都可能会失效。
    - 建议用户在升级前备份旧版配置文件。

## 3. 项目进展

今日共合并/关闭了 3 个 PR，标志着项目在稳定性和兼容性方面取得了重要进展：

- **[PR #7520] fix(ci): install cross g++ for ARM glibc release builds** (已关闭): 修复了 v0.8.0 发布流水线中 ARM Linux 构建失败的问题，确保了该版本的跨平台兼容性。
- **[PR #7517] fix(runtime/subagent): inherit ACP session cwd into spawn_subagent and delegate** (已关闭): 修复了一个严重的 Bug，即子代理在 ACP（Agent Communication Protocol）会话中无法继承父会话的当前工作目录（CWD），导致工作流被阻塞。
- **[PR #7519] fix(config): persist [[mcp.servers]] per-field edits via natural-key dirty-path walker** (已关闭): 修复了配置编辑后 MCP 服务器配置无法正确持久化到磁盘的问题，这是一个导致用户配置丢失的“灵异事件”。

这些修复直接对应了社区中长时间讨论的几个高优先级问题，有力地提升了 v0.8.0 版本的稳定度。

## 4. 社区热点

以下 Issues 和 PRs 在今日讨论最为活跃（评论数均为 4 条），反映了社区的关注焦点：

1.  **[Issue #5542] [Bug]: consecutive OOM in wsl2** (4 条评论): 一个长期存在的 WSL2 环境下的 OOM（内存溢出）Bug。用户“Themoonshinesontheriver”正在积极与维护者沟通，要求提供复现步骤（needs-repro）。这表明 WSL2 用户群体仍在期待该严重性能问题得到解决。[链接](https://github.com/zeroclaw-labs/zeroclaw/issues/5542)

2.  **[Issue #6302] [Bug]: Gemini 400 — assistant tool_call emitted as first non-system turn** (4 条评论): 一个与 Gemini 模型 API 兼容性的关键 Bug。用户“dmnkhorvath”指出 ZeroClaw 构造的对话历史违反了 Gemini 的 API 规范，导致工具调用失败。这反映出社区对多模型支持（特别是 Gemini）的稳定性诉求强烈。[链接](https://github.com/zeroclaw-labs/zeroclaw/issues/6302)

3.  **[Issue #6312] feat(gateway): per-alias webhook path routing for multi-instance channels** (4 条评论): 该功能请求讨论热度极高，用户希望实现更精细的 Webhook 路由，以支持同一网关下多个 Bot 实例的独立运行。虽然已有 `/webhook?agent=` 的解决方案，但社区仍倾向于路径路由方案。[链接](https://github.com/zeroclaw-labs/zeroclaw/issues/6312)

4.  **[Issue #6391] [Feature]: real heartbeat tracking for daemon nodes** (4 条评论): 社区对多节点管理功能表现出浓厚兴趣。用户“theonlyhennygod”提出的实时心跳检测功能，是实现可靠仪表盘监控和集群管理的关键一步，相关讨论非常深入。[链接](https://github.com/zeroclaw-labs/zeroclaw/issues/6391)

## 5. Bug 与稳定性

今日新报告的 Bug 集中在 v0.8.0 新版本的兼容性问题上，严重程度高：

- **S1 - 工作流阻塞**:
    - **`macos app not work`** ([#7527]): 用户报告 macOS 15.7.7 系统上 ZeroClaw 应用无法检测权限、显示空白页甚至窗口消失。 **（暂无对应修复 PR）**
    - **`dashboard not available`** ([#7523]): macOS Brew 安装的用户无法访问 Web 仪表盘，需要手动运行 `cargo web build` 才能工作。 **（暂无对应修复 PR）**

- **S0 - 数据丢失/安全风险**:
    - **`consecutive OOM in wsl2`** ([#5542]): 长期未决，表现为 WSL2 环境下进程被内核 OOM Killer 杀死。 **（暂无对应修复 PR，等待复现）**

- **S1 - 工作流阻塞**:
    - **`WhatsApp Web — allowed-numbers bypassed`** ([#6350]): 配置中的允许号码列表被绕过，导致消息被静默丢弃。 **（暂无对应修复 PR）**

- **S2 - 功能降级**:
    - **`model_switch tool does not persist`** ([#6173]): 模型切换工具不生效，切换不跨对话轮次持久化。 **（暂无对应修复 PR）**
    - **`file_read — decode non-UTF-8 text`** ([#7521]): `file_read`工具对非UTF-8文件内容处理不当，输出乱码。幸运的是，已有 PR [#7522](https://github.com/zeroclaw-labs/zeroclaw/pull/7522) 尝试修复此问题。

## 6. 功能请求与路线图信号

- **高优先级**:
    - **[Issue #7521] file_read — decode non-UTF-8 text**：社区用户强烈需要 `file_read` 工具支持检测和正确处理非 UTF-8 编码的文本文件（如中文的 GBK 编码）。**已有 PR #7522 尝试解决**，表明维护者已采纳此建议。
    - **[Issue #6365] Dashboard "Update ZeroClaw" button**：用户希望通过 Web 仪表盘直接更新 ZeroClaw 至最新版，无需使用命令行。这是一个提升用户体验的关键功能。

- **中优先级**:
    - **[Issue #6346] zeroclaw node CLI + dashboard health & management**：与 #6391 相关，社区对完整的节点管理和健康监控功能需求明确，这将是实现“舰队管理”的重要一环。
    - **[Issue #6443] Add Twitch chat channel**：一个低成本的功能扩展请求，利用现有 IRC 模块适配 Twitch 平台，可能吸引更多直播领域的用户。**该 Issue 已于今日关闭**，但具体实现尚未合并。

这些功能请求与 v0.8.0 引入的多代理、多节点架构方向一致，预计将在后续的 v0.9.x 版本中逐步实现。

## 7. 用户反馈摘要

- **积极反馈**:
    - v0.8.0 版本发布，用户对多代理架构和自动迁移功能充满期待。
    - 社区对于子代理 CWD 继承 Bug（#7263）的修复（#7517）表示欢迎，这解除了许多开发者使用 ZeroClaw 作为开发工具时的障碍。

- **痛点与抱怨**:
    - **macOS 体验糟糕**: 多个 macOS 用户报告应用无法正常工作（#7527），且仪表盘需要手动构建（#7523），表明 v0.8.0 在 macOS 平台的发布准备不足，用户体验较差。
    - **配置持久化问题**: Issue #7519 修复的配置无法保存问题，此前已导致多名用户在修改 MCP 服务器配置时感到困惑和不满。
    - **部署稳定性**: 用户报告在新版中仍遇到老问题，如 OOM（#5542）和 Cron 任务重复执行（#6037），反映出部分系统稳定性 Bug 积压未解。

## 8. 待处理积压

- **多条“等待作者处理”的 PR**: 包括修复 Windows Claude Code 的 [#7214](https://github.com/zeroclaw-labs/zeroclaw/pull/7214)、MCP 自动重连的 [#7351](https://github.com/zeroclaw-labs/zeroclaw/pull/7351) 以及长期存在的 fuzz 测试 [#5516](https://github.com/zeroclaw-labs/zeroclaw/pull/5516) 等。这些 PR 作者在收到维护者反馈后未能及时更新，若不尽快回应有被标记为“陈腐”的风险。

- **高风险的“陈腐候选”PR**:
    - **[PR #5892] fix(providers,runtime)**: 该 PR 解决了 vLLM、Mistral 和 OpenRouter 的 3 个生产级 Bug，但已超过 50 天无活动，被标记为 `stale-candidate`。鉴于其重要性（“生产级阻塞器”），维护者应优先处理。
    - **[PR #6143] feat(skills)**: 这是一个大规模的技能注册表功能，但同样处于停滞状态。长期搁置可能导致大量代码冲突。

- **未分配的高风险 Issues**:
    - **[Issue #5903] MCP子进程泄漏**: 一个严重的内存/进程泄漏问题，已被标记为 `priority:p1` 但至今无对应 PR。该问题在默认开启心跳机制时就会触发，影响面极广。
    - **[Issue #6173] model_switch 不生效**: 该 Bug 影响 Agent 工具链的可靠性，但修复优先级似乎不高，长期处于 `in-progress` 状态。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*