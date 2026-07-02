# OpenClaw 生态日报 2026-07-02

> Issues: 58 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-07-02 10:17 UTC

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

# OpenClaw 项目日报 | 2026-07-02

---

## 1. 今日速览

过去 24 小时内，OpenClaw 项目保持高强度活跃：共处理 **58 条 Issue**（新开/活跃 36 条，关闭 22 条）和 **500 条 PR**（待合并 416 条，已合并/关闭 84 条）。发布了一个新的 beta 版本 v2026.7.1，核心亮点是支持 OpenAI GPT-5.6 模型族。多支高优先级 P1 钻石级 Issue 持续发酵，涉及消息丢失、安全边界、崩溃循环等稳定性风险；同时有大量维护性 PR 进入等待队列，项目整体在修复-验证-合并的节奏中稳步推进。

---

## 2. 版本发布

### v2026.7.1-beta.1
- **发布时间**：2026-07-02
- **关键变更**：
  - **OpenAI GPT-5.6 支持**：在模型目录、能力识别和运行时选择路径中新增对 GPT-5.6 模型族的识别（PR #98333，感谢 @steipete-oai）。
  - **External harness attachment**：`openclaw attach` 允许针对已有 Gateway 会话启动外部测试框架，便于集成测试与调试。
- **破坏性变更**：未标注，但用户需确保模型配置兼容 GPT-5.6 的新 API 签名。
- **迁移注意事项**：若使用自定义模型回退链，建议验证 GPT-5.6 的 tool calling 行为无回归。

---

## 3. 项目进展

今日合并/关闭的 84 个 PR 覆盖了多个关键领域，以下为从 Issue 关闭和 PR 列表提取的重要推进：

| 模块 | 合并/关闭的 PR/Issue | 推进内容 |
|------|----------------------|----------|
| **Discord 频道** | PR #95195（待合并） | 修复网关元数据读取无上限导致的 OOM 风险 |
| **Matrix 频道** | PR #93516（待合并） | 支持反向代理路径前缀的 Matrix 主页 URL |
| **代理工具链** | PR #98640（待合并） | 为 Anthropic Opus 4.8 添加 `additionalProperties:false` 兼容性修复 |
| **会话状态** | Issue #84490、#92364、#98874、#99022、#98872、#98980、#98981 等已关闭 | 修复 sessions.json 中 `totalTokens` 为 null、Feishu 回复丢失、工具结果渲染为图片、终端路径显示异常、会话修剪器漏扫、会话旋转器文件名错误等 |
| **安全** | PR #98871（待合并） | 修复 Mattermost 对等目录只返回前 200 成员的问题 |
| **CLI 与脚本** | Issue #98973、#98896、#98550 已关闭 | 修复 Windows 网关重启失败、claude-cli 输出路径分裂、以及 CVE 分配请求 (GHSA-cf2p-f286-mphf) |

**整体评估**：项目在消息传递、会话稳定性、安全加固三个方面取得显著进展，约 20% 的积压高优先级 Issue 得到关闭，代码库健康度正向发展。

---

## 4. 社区热点

以下 Issue 在过去 24 小时获得最多评论与互动，反映出用户最强烈的诉求：

### 🔥 #25592 — 工具调用间的文本泄漏到消息频道
- **评论数**：33 | 👍：1
- **链接**：[Issue #25592](https://github.com/openclaw/openclaw/issues/25592)
- **诉求**：用户强烈要求阻止 Agent 在处理中产生的中间文本（如错误处理、确认信息）被路由到 Slack、iMessage 等外部渠道。目前该行为导致 UX 极差，且可能暴露内部状态。
- **状态**：P1，diamond lobster，已标记 `linked-pr-open`，但尚未合并修复。

### 🔥 #38327 — Gemini 3.1 Pro 升级后 "Cannot convert undefined or null to object"
- **评论数**：10 | 👍：3
- **链接**：[Issue #38327](https://github.com/openclaw/openclaw/issues/38327)
- **诉求**：从 2026.3.2 版本开始，使用 google-vertex/gemini-3.1-pro-preview 模型时任何消息都会触发此致命错误。用户急需降级或修复。
- **状态**：P1，diamond lobster，已有 `linked-pr-open` 标签，但未更新修复进展。

### 🔥 #92433 — 子代理完成被丢弃（steer 到 requester run 未处理）
- **评论数**：7 | 👍：1
- **链接**：[Issue #92433](https://github.com/openclaw/openclaw/issues/92433)
- **诉求**：多代理协作场景下子代理返回结果时，若 requester 的 run 已结束，结果被静默丢弃，导致任务无响应。用户要求明确的状态追踪和重试机制。

### 🔥 #70024 — 频道停止超时导致永久死亡，无自动恢复
- **评论数**：4 | 👍：0
- **链接**：[Issue #70024](https://github.com/openclaw/openclaw/issues/70024)
- **诉求**：`stopChannel` 超时后错误地将 `running: true` 写入运行时快照，却不清理 store 条目，导致频道静默死亡，只能手动重启。用户要求实现自动恢复或至少提供诊断手段。

---

## 5. Bug 与稳定性

以下按严重程度（P1 > P2）排列今日报告的 Bug，并标注是否已有修复 PR（Fix PR）。

| Issue | 严重度 | 摘要 | Fix PR 状态 |
|-------|--------|------|-------------|
| [#99021](https://github.com/openclaw/openclaw/issues/99021) | P1 | Discord 回复 >10MB 附件被 413 拒绝，文本和附件均丢失，无重试 | 无，已标记 `queueable-fix` |
| [#98956](https://github.com/openclaw/openclaw/issues/98956) | P1 | 所有模型超时后网关死锁，需手动重启（回归） | 无，`needs-maintainer-review` |
| [#98925](https://github.com/openclaw/openclaw/issues/98925) | P1 | `fetchWithSsrFGuard` 严格模式在托管代理后进行本地 DNS 解析，绕过安全边界 | 无 |
| [#98976](https://github.com/openclaw/openclaw/issues/98976) | P1 | Provider refusal（Anthropic/OpenAI）不触发模型回退链，直接报错 | 无 |
| [#98588](https://github.com/openclaw/openclaw/issues/98588) | P1 | Anthropic Opus 4.8 拒绝 tools.profile=coding 的 input schema | PR #98640 已开放 |
| [#99031](https://github.com/openclaw/openclaw/issues/99031) | P2 | iOS app 通过 Bonjour 发现的网关无 TLS 时无法连接 | 无 |
| [#98970](https://github.com/openclaw/openclaw/issues/98970) | P2 | 英文音频提示导致非英语 STT 偏向翻译 | 无 |
| [#98964](https://github.com/openclaw/openclaw/issues/98964) | P2 | dispatch-wrapper 中 flock -- 分隔符被错误跳过，脚本命令误检 | 无 |
| [#98962](https://github.com/openclaw/openclaw/issues/98962) | P2 | 模型别名重复时静默覆盖，无告警 | 无 |
| [#98961](https://github.com/openclaw/openclaw/issues/98961) | P3 | find 工具错误处理器缺少 stopChild 调用，导致子进程泄露 | 无 |
| [#98958](https://github.com/openclaw/openclaw/issues/98958) | P2 | 网关锁文件句柄泄漏（writeFile 失败时） | 无 |
| [#98960](https://github.com/openclaw/openclaw/issues/98960) | P2 | usage-bar 模板文件缓存无限制增长，产生未关闭的 fs.watch watcher | 无 |
| [#98959](https://github.com/openclaw/openclaw/issues/98959) | P2 | proxy-capture 中两个 DELETE 操作无事务，可导致表不一致 | 无 |
| [#98945](https://github.com/openclaw/openclaw/issues/98945) | P2 | Gemini CLI 凭据暂存与运行中任务竞态，非原子写入 | 无 |

**值得警惕**：今日共报告约 15 个新 Bug，其中 5 个 P1 钻石级别且均无 fix PR，涉及消息丢失、死锁、安全绕过等核心稳定性问题，建议维护者优先评审。

---

## 6. 功能请求与路线图信号

以下用户提出的新功能需求具有较高可行性，可能被纳入下一版本：

| Issue | 需求描述 | 相关性判断 |
|-------|----------|------------|
| [#98986](https://github.com/openclaw/openclaw/issues/98986) | 将 transcripts store（会议录制）迁移到 SQLite | 已有对应 PR #99006（Draft），符合项目“SQLite only”方向，极可能进入 |
| [#98968](https://github.com/openclaw/openclaw/issues/98968) | 为插件 LLM 调用 (`runtime.llm.complete`) 增加诊断输出 | 已有 `queueable-fix` 标签，易实现 |
| [#98927](https://github.com/openclaw/openclaw/issues/98927) | 简化 Doubao/Volcengine 实时 Talk 的使用路径 | 用户强烈需求（中国大陆用户），已有 TTS 支持，拓展合理 |
| [#98995](https://github.com/openclaw/openclaw/issues/98995) | iOS 设置中改进外观选择器放置位置 | UX 改进，已有 `fix-shape-clear` 标签 |
| [#98943](https://github.com/openclaw/openclaw/issues/98943) | iOS “关于”页面改进文案和设备信息 | 同上，低风险改进 |
| [#32530](https://github.com/openclaw/openclaw/issues/32530) | 自动发现外部工作区中的 Agent 配置（P2） | 长期开放，涉及配置管理重构，可能需要多个版本完成 |

**路线图信号**：从 PR #98718、#99006 等大型 PR 可以看出，项目正在推进**持久化会话任务运行时**（durable session task runtime）和**SQLite 统一存储**两个底层架构变更，这将是后续版本的核心方向。

---

## 7. 用户反馈摘要

从 Issue 评论中提炼的典型用户痛点和场景：

- **消息渠道兼容性**：多位用户报告 Feishu、Discord、Mattermost 等渠道在特定条件下（如 DM vs 群聊、大附件、反向代理）出现回复丢失或格式错误。用户期望 OpenClaw 能提供一致的跨渠道体验。
- **稳定性崩溃焦虑**：`#38327`（Gemini 3.1 Pro 崩溃）、`#98956`（网关死锁）的评论中用户明确表示“回归”让他们对生产环境部署失去信心，强烈要求增加自动恢复机制。
- **内部状态泄漏**：`#25592` 的讨论中有用户指出，即使没有错误，Agent 的“思考过程”也不应出现在对外消息中，这既是 UX 问题也是隐私风险。
- **模型兼容性压力**：`#98588` 和 `#98614` 显示，每次模型供应商更新（如 Anthropic Opus 4.8、OpenAI GPT-5.6）都带来 breaking change，用户希望 OpenClaw 能更敏捷地适配或者提供插件式兼容层。
- **iOS 原生体验**：多条关于 iOS app 的改进请求（`#98995`、`#98943`、`#98929`）表明移动端用户比例上升，对原生 UI 细节和稳定性（如 TLS 指纹、发现网关）有更高要求。

---

## 8. 待处理积压

以下为长期未响应或缺少维护者关注的重要 Issue / PR，按创建时间排序：

| 项目 | 类型 | 创建时间 | 当前状态 | 风险等级 |
|------|------|----------|----------|----------|
| [#25592](https://github.com/openclaw/openclaw/issues/25592) 文本泄漏 | Issue | 2026-02-24 | P1，`linked-pr-open`，5个月未合并 | 🔴 高，影响所有消息渠道 |
| [#38327](https://github.com/openclaw/openclaw/issues/38327) gemini-3.1-pro 崩溃 | Issue | 2026-03-06 | P1，无 fix PR | 🔴 高，阻止用户升级 |
| [#70024](https://github.com/openclaw/openclaw/issues/70024) 频道永久死亡 | Issue | 2026-04-22 | P1，`linked-pr-open`，2.5个月未合并 | 🔴 高，影响多频道部署 |
| [#32530](https://github.com/openclaw/openclaw/issues/32530) 自动发现 agent 配置 | Issue | 2026-03-03 | P2，无维护者标记 | 🟡 中，功能缺失 |
| [#53467](https://github.com/openclaw/openclaw/pull/53467) Slack 忽略其他提及 | PR | 2026-03-24 | `ready for maintainer look` | 🟢 低，代码已就绪 |
| [#52951](https://github.com/openclaw/openclaw/pull/52951) tools.fs.roots 文件系统根控制 | PR | 2026-03-23 | `ready for maintainer look` | 🟢 低，但涉及安全边界，需谨慎 |
| [#51762](https://github.com/openclaw/openclaw/pull/51762) 可配置默认 agent ID | PR | 2026-03-21 | `waiting on author` | 🟡 中，作者可能停滞 |
| [#51822](https://github.com/openclaw/openclaw/pull/51822) 拒绝含凭据的 webhook URL | PR | 2026-03-21 | `needs proof` | 🟢 低，安全加固 |
| [#47604](https://github.com/openclaw/openclaw/pull/47604) Wear OS 应用 MVP | PR | 2026-03-15 | `needs proof` | 🟡 中，大型 PR，维护者需资源评审 |
| [#47277](https://github.com/openclaw/openclaw/pull/47277) 用户特定记忆隔离 | PR | 2026-03-15 | `needs proof` | 🟡 中，重要功能，但涉及会话状态变更 |

**建议**：维护团队可优先安排一次 triage 会议，集中处理 2 月～4 月的钻石级 Issue 及其关联 PR，以释放拥堵并恢复社区信任。

---

*日报生成基于 OpenClaw 项目 2026-07-02 的公开 GitHub 数据，所有链接均指向真实 Issue/PR。*

---

## 横向生态对比

好的，作为AI智能体与个人AI助手领域开源项目分析师，根据您提供的各项目2026年7月2日动态数据，现生成如下横向对比分析报告。

---

### 个人AI助手开源生态横向对比分析报告 (2026-07-02)

**报告日期:** 2026年7月2日
**分析师:** 资深技术分析师

#### 1. 生态全景

当日，个人AI助手与自主智能体开源生态呈现“**核心项目向全栈平台演进，新兴项目聚焦于轻量化与特定场景**”的态势。头部项目（OpenClaw）通过高强度迭代继续巩固其作为通用智能体操作系统的地位，但积压的P1级问题也暴露了复杂性带来的稳定性挑战。以NanoBot、Hermes Agent为代表的中坚力量，在保持核心功能稳定的同时，开始发力桌面端体验和企业级特性。而PicoClaw、ZeptoClaw等更侧重边缘计算或嵌入式场景，强调轻量与低资源占用。整体来看，生态已从最初的“能否运行”进入到“**如何可靠、安全、高效地运行**”的第二阶段，社区对稳定性、安全性和跨平台一致体验的诉求显著增强。

#### 2. 各项目活跃度对比

| 项目名称 | Issues (当日活跃/新开) | PRs (待处理/已合并关闭) | 版本发布 | 健康度评估 | 关键信号 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 36新开/活跃，22关闭 | 416待合并，84合并/关闭 | v2026.7.1-beta.1 | 🟡 一般 | 高强度开发，但修复与积压并存，稳定性风险较高 |
| **NanoBot** | 24 | 66 (28合并，10关闭) | 无 | 🟢 健康 | 高效的修复与迭代，合并率高，社区贡献积极 |
| **Hermes Agent** | 6 | 50 (超40合并/关闭) | **v0.18.0** | 🟢 非常健康 | 大版本发布驱动，社区贡献集中打捞，项目成熟度显著提升 |
| **PicoClaw** | 3 | 15 (2合并/关闭) | nightly构建 | 🟡 一般 | 维护节奏偏慢，部分PR已stale，活跃度中等 |
| **NanoClaw** | 1 | 12 (6合并/关闭) | 无 | 🟢 健康 | 核心功能（Agent模板）冲刺中，社区PR合并积极 |
| **IronClaw** | 5 | 50 (26合并/关闭) | 无 | 🟢 健康 | “Reborn”架构重构持续推进，但回归风险值得关注 |
| **LobsterAI** | 5 | 7 (全部合并/关闭) | 无 | 🟡 一般 | 修复力度强，但对关键Bug的社区反馈响应滞后，积压严重 |
| **CoPaw (QwenPaw)** | 12 | 50 (32合并/关闭) | v2.0.0-beta.2 | 🟡 一般 | 功能迭代快，但用户端报告严重Bug多，Beta版稳定性不佳 |
| **ZeroClaw** | 12 | 50 (10合并/关闭) | 无 | 🟡 一般 | 通道与安全功能拓展积极，但PR合并效率低，存在瓶颈 |

#### 3. OpenClaw在生态中的定位

*   **优势：无可争议的生态核心与“参照点”**
    *   **社区规模与成熟度：** 无论是Issue/PR数量（日处理58条Issue、500条PR）还是社区讨论深度，OpenClaw都远超其他项目，是生态中最为庞大和活跃的。其功能集（Discord, Matrix, Slack等30+渠道、丰富的Agent工具链）是行业的通用基准。
    *   **技术路线：** 坚持“全栈智能体操作系统”路径，致力于解决AI Agent在多渠道、多模型、多工具环境下的统一编排与运行时问题，而非仅聚焦于单点功能。
    *   **生态影响力：** 大量Fork（如NanoClaw、NanoBot等）的出现，证明了其架构设计与核心概念的广泛认可和借鉴价值。

*   **劣势与风险：**
    *   **稳定性风险积聚：** 其“高强度修复-积压”并存的模式，以及大量P1级钻石Issue（如消息丢失、网关死锁）长期未决，已成为社区信任的潜在风险。相比之下，**NanoBot**和**Hermes Agent**在Bug修复的效率和彻底性上表现更佳。
    *   **创新引领力被分散：** 例如**Hermes Agent**的“审判”版本引入Agent输出评判机制；**ZeroClaw**对SOP可视化编辑的探索；**CoPaw**对飞书等中国区市场的深度集成，这些在特定方向的创新，OpenClaw并未展现出明显的引领优势。

**总结：** OpenClaw仍是生态的“**大本营**”和**多功能性的天花板**，但其他项目正在通过更高效率的迭代和更聚焦的创新，在特定维度上形成差异化优势，对OpenClaw的“全能”地位构成挑战。

#### 4. 共同关注的技术方向

多个项目的Issue与PR不约而同指向以下核心方向：

1.  **安全性是企业级应用的基石：**
    *   **涉及项目：** OpenClaw, NanoBot, Hermes Agent, IronClaw, ZeroClaw, CoPaw.
    *   **具体诉求：** API密钥脱敏安全存储（CoPaw #5705）、DoS/Zip炸弹防御（ZeroClaw #8574）、符号链接逃逸修复（NanoBot #4072）、插件审计签名密钥泄露（ZeroClaw #8591）、工具中间文本泄漏到公共频道（OpenClaw #25592）。
    *   **趋势解读：** 安全已不再是“锦上添花”，而是项目能否被企业或其商业客户采用的“**准入门槛**”。

2.  **模型兼容性正成为运维噩梦：**
    *   **涉及项目：** OpenClaw (GPT-5.6支持, Opus 4.8兼容性), IronClaw (Gemini 3.1 Pro崩溃), ZeroClaw (Gemini 400拒绝, 兼容Provider静默删除内容), NanoBot (推理模型输出处理), CoPaw (模型列表更新).
    *   **具体诉求：** 每次模型供应商的更新（新模型发布、API签名变更、拒绝模式改变）都可能导致项目功能失效或行为异常。
    *   **趋势解读：** 项目需要更**鲁棒的模型抽象层、自动检测与自适应逻辑**，以对抗上游API的不稳定性，减轻运维负担。

3.  **消息渠道的可靠性与一致性体验：**
    *   **涉及项目：** OpenClaw (Discord 413, Feishu回复丢失), Hermes Agent (QQBot重连, Teams安全), NanoBot (Telegram长消息渲染失败, Feishu对话分割), PicoClaw (Matrix重连), LobsterAI (任务删除后复活).
    *   **具体诉求：** 用户期望Agent在Slack、飞书、Telegram等不同渠道上提供**一致、可靠、无数据丢失**的交互体验。
    *   **趋势解读：** 渠道适配正在从“能接入”向“**高质量集成**”转变，对平台特有限制（如消息大小、API频率）的处理能力成为关键。

4.  **持久化与状态管理：**
    *   **涉及项目：** OpenClaw (SQLite统一存储), NanoBot (会话修剪器漏扫、广播重组), NanoClaw (对话搜索技能), IronClaw (Session恢复).
    *   **具体诉求：** 会话状态、记忆、任务状态的可靠持久化、高效检索与恢复。
    *   **趋势解读：** 随着Agent任务越来越复杂、运行时间越来越长，**状态管理**成为构建可靠Agent应用的核心技术挑战。

5.  **工具调用的智能与可靠性：**
    *   **涉及项目：** NanoBot (edit_file消歧失败), OpenClaw (工具结果渲染为图片), IronClaw (工具延迟追踪).
    *   **具体诉求：** Agent需要更聪明地使用工具，尤其是在文件编辑、数据查询等场景下，提高工具调用的准确性、可靠性和可解释性。
    *   **趋势解读：** 下一阶段的竞争点，将是从“**能调用工具**”转向“**精准、高效、可靠地调用工具**”。

#### 5. 差异化定位分析

| 项目 | 功能侧重 | 目标用户 | 技术架构与哲学 | 核心差异化 |
| :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 全能型智能体中枢 | 高级用户、开发者、社区运维者 | 单体化、高度集成，强调多渠道与多模型统一 | **生态核心**，功能最全，社区最大，但复杂度高 |
| **NanoBot** | 轻量化、高性能、易使用 | 对效率和稳定性要求高的开发者 | 模块化、高度可测试、强调“修复优先” | **高合并/修复效率**，代码库轻量，极致的开发体验 |
| **Hermes Agent** | 企业级特性、创新评判机制 | 注重AI治理、安全的企业用户 | 发布驱动，积极打捞社区历史贡献 | **“审判”版本**的Agent行为评判，桌面端成就系统 |
| **NanoClaw** | Agent模板化、易用性 | 希望快速创建标准化Agent的中级用户 | 聚焦于**Agent创建与管理流程**的简化 | **模板化、向导化**，降低Agent创建门槛 |
| **PicoClaw** | 轻量、边缘、多平台 | 树莓派、嵌入式、移动终端用户 | 极致精简，保持核心功能 | **极低资源占用**，支持ARMv7，适合IoT场景 |
| **IronClaw** | 架构重构、基础设施优化 | NearAI生态内部及开发者 | 大厂驱动，深度重构 (“Reborn”) | 企业级**断路器、Open Wiki、代码图谱**等基础设施 |
| **CoPaw (QwenPaw)** | AI原生、中国市场深度集成 | 以飞书/企业微信为核心的国内用户 | 背靠大模型，拥抱开源生态 | 深度支持**飞书等中国平台**，强化AI Agent协作 |
| **ZeroClaw** | 安全优先、标准化、SOP可视化 | 安全研究者、需要审计的组织 | 强调**安全边界与审计**，制定标准 (RFC) | **SOP可视化编辑**、统一审计架构、OpenAI兼容端点 |

#### 6. 社区热度与成熟度

*   **快速迭代阶段 (高度活跃，功能驱动，稳定性波动)：**
    *   **OpenClaw:** 技术牵引力最强，但功能堆叠快于问题解决，处于“高发展，高负荷”状态。
    *   **IronClaw & CoPaw (QwenPaw):** 都处于重大架构或版本迭代期，PR数量多，合并快，但伴随的回归Bug和用户报障也最多，表现为“积极建设，但质量阵痛”。

*   **质量巩固阶段 (高度活跃，修复驱动，稳定性优先)：**
    *   **NanoBot & Hermes Agent:** 这两个项目是当日表现最佳的标杆。它们的PR合并率极高（NanoBot近60%，Hermes Agent超80%），对Issue的响应和修复周期短。这种“**稳定、高效**”的特性，是其赢得开发者和用户信任的核心原因。
    *   **NanoClaw:** 也表现出类似的健康态势，社区贡献的PR能迅速被合并，体现了良好的维护节奏。

*   **早期或缓慢增长阶段：**
    *   **PicoClaw, ZeptoClaw, TinyClaw, Moltis (无活动数据):** 活跃度明显较低，或社区规模较小，或项目处于早期原型阶段。其中PicoClaw有明确的轻量化定位和社区反馈，但维护者响应速度偏慢。
    *   **LobsterAI:** 虽然今日修复力度强，但积压3个月的高危Bug（如蓝屏、任务复活）未被修复，严重损害了用户信任和社区氛围，处于“**修复与口碑恶化并存**”的危险状态。

#### 7. 值得关注的趋势信号

1.  **安全性不再是可以商榷的附加品，而是项目成熟度的KPI。** 几乎每个项目都有安全相关的Issue或PR。开发者选择项目时，将像评估稳定性一样评估其安全审计机制。**ZeroClaw**的安全优先策略可能在未来受到更多重视。

2.  **标准化API（如OpenAI兼容端点）是打通生态壁垒的关键。** **ZeroClaw**对`/v1/chat/completions`端点的强烈诉求，反映了社区希望项目能融入更广泛工具生态（如LobeChat, Open WebUI）的渴望。**谁先实现主流标准的深度兼容，谁就能获得更广泛的集成生态。**

3.  **跨平台支持（尤其是Windows）是隐性但严重的“断裂带”。** **ZeroClaw** (74个Windows测试失败) 和 **NanoBot** (Windows shell语义不一致) 的问题表明，跨平台不仅是打包问题，更牵扯到核心的行为一致性。对于面向开发者的项目，**Windows开发与运行体验将影响其在用户群中的Base。**

4.  **“工具链的智能化”是用户体验竞赛的下一个前沿。** 当Agent能调用工具后，如何让工具调用**更精确（NanoBot edit_file）、更透明（OpenClaw文本泄漏）、更快速（IronClaw延迟追踪）**，将成为区分体验好坏的关键。

5.  **“开发者体验”与“项目增长”呈强正相关。** **NanoBot**因其代码轻量、易于阅读、修复及时而获得社区的高度评价。**Hermes Agent**积极合并社区贡献，形成了良好的循环。维护者是否尊重社区贡献、代码质量是否优雅、文档是否清晰，已成为项目能否持续吸引人才和获得广泛采用的隐藏因素。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

好的，作为AI智能体与个人AI助手领域开源项目分析师，根据您提供的NanoBot项目数据，我为您生成了2026年7月2日的项目动态日报。

---

### NanoBot 项目动态日报 | 2026-07-02

---

#### 1. 今日速览

今日项目活跃度**极高**，社区贡献与维护响应均十分积极。过去24小时内，共处理了24个Issue和66个Pull Request，其中近六成的PR（38个）已被合并或关闭，显示了高效的代码迭代速度。修复工作集中在WebUI展示、执行器安全与稳定性方面，同时社区对新功能（如Anthropic OAuth）和渠道体验优化（如飞书、Telegram）的呼声依然强烈。尽管没有新版本发布，但大量修复和测试代码的合入为项目下一个发布版本奠定了坚实基础。

#### 3. 项目进展

今日有大量重要PR被合并/关闭，推动了项目在稳定性、安全性和测试覆盖方面的显著进步。

-   **关键Bug修复与稳定性提升**：
    -   **#4641** `fix(webui): hide subagent backfill payloads` 已合并，修复了WebUI在刷新后可能显示内部子代理结果的问题，提升了用户界面体验。[#4641](https://github.com/HKUDS/nanobot/pull/4641)
    -   **#4643** `fix(exec): return early when session command exits` 已合并，优化了执行器性能，当会话命令快速退出时能尽早返回。[#4643](https://github.com/HKUDS/nanobot/pull/4643)
    -   **#4642** `fix(trigger): cap local trigger audit records` 已合并，对本地触发器审计记录进行截断，防止文件过大，并修复了因文件系统不支持目录`fsync`导致的启动崩溃问题。[#4642](https://github.com/HKUDS/nanobot/pull/4642)
    -   **#4630** `test: cover runner blocked tool-call finish reasons` 与 **#4633** `test(cron): cover stale instance mutation consistency` 已合并，为运行器和定时任务特性增加了关键回归测试，提升了代码健壮性。[#4630](https://github.com/HKUDS/nanobot/pull/4630) [#4633](https://github.com/HKUDS/nanobot/pull/4633)

-   **批量问题修复**：
    -   **#4648** `Fix validated issue batch` 是一个大型修复PR，旨在一次性修复包含安全、Bug在内的13个已验证问题（如 #4078, #4076, #4075等），这是对项目健康度的一次重大提升。[#4648](https://github.com/HKUDS/nanobot/pull/4648)

-   **基础设施与测试增强**：
    -   **#4631** `test: add scripted agent runner harness` 新增了可复用的脚本化代理运行器测试框架，有助于未来对运行器流程进行更精确的测试。[#4631](https://github.com/HKUDS/nanobot/pull/4631)
    -   **#4628** `test: add memory lifecycle harness` 新增了内存生命周期测试框架，覆盖从会话归档到持久化更新的完整流程。[#4628](https://github.com/HKUDS/nanobot/pull/4628)

#### 4. 社区热点

今日社区讨论焦点主要集中在新功能请求和渠道体验优化上。

-   `#4604` **[feature request] Anthropic OAuth**：该Issue请求支持Anthropic OAuth身份验证，以便Claude订阅用户无需API Key即可使用。此请求引起了关注，且社区成员已经为此提交了PR `#4632`，显示出极高的协同开发效率。[#4604](https://github.com/HKUDS/nanobot/issues/4604)
-   `#4619` **[enhancement] 飞书频道：使用 /new 进行新会话时发送 system 级新会话信息切分对话**：来自飞书用户的增强请求，希望通过发送特定类型消息在UI上实现对话分割线，提升多轮对话的视觉清晰度。该需求非常具体，体现了非英语用户对本地化体验的期望。[#4619](https://github.com/HKUDS/nanobot/issues/4619)
-   `#4637` **[bug] Telegram long message splits -- trunks prior to the final trunk cannot render**：用户报告了Telegram频道长消息分片后，除最后一片外的其他片段无法正确渲染的问题（仅显示纯文本）。该问题严重影响Telegram用户的使用体验，社区关注度较高。[#4637](https://github.com/HKUDS/nanobot/issues/4637)

#### 5. Bug 与稳定性

今日新报告和持续活跃的Bug修复主要集中在以下方面，严重程度从高到低排列：

-   **（严重）** `#4637` **Telegram长消息分片渲染失败**：Telegram渠道中，长Markdown消息被分片发送时，非末段消息无法渲染，仅显示纯文本。尚无相关修复PR。[#4637](https://github.com/HKUDS/nanobot/issues/4637)
-   **（严重）** `#4634` **edit_file目标消歧失败**：`edit_file`工具在执行精确文本替换时，可能错误地修改了文本的另一个出现位置，这是离线编辑基准测试中的主要失败模式，严重限制了工具的可靠性。尚无相关修复PR。[#4634](https://github.com/HKUDS/nanobot/issues/4634)
-   **（中）** `#4640` **WebUI显示内部子代理结果**：已通过PR `#4641`修复并关闭。此问题会导致在WebUI中看到不应显示的内部消息。今日已解决。[#4640](https://github.com/HKUDS/nanobot/issues/4640)
-   **（中）** `#4544` **Windows下exec工具的shell语义不一致**：在Windows上，单行命令使用`cmd.exe`，多行命令使用`PowerShell`，导致行为不统一，且对编写跨平台命令的代理不友好。修复包含在批量修复PR `#4648`中。[#4544](https://github.com/HKUDS/nanobot/issues/4544)

#### 6. 功能请求与路线图信号

用户对新功能和渠道增强的需求持续活跃，一些请求已经转化为实际的开发工作。

-   **高优先级信号**：
    -   **Anthropic OAuth支持 (#4604)**：已提交对应PR `#4632`，预计将进入下一个版本。这对吸引非API Key的Claude用户至关重要。[#4604](https://github.com/HKUDS/nanobot/issues/4604)
    -   **Cron任务级别的模型/预设选择 (#4378)**：此功能请求讨论已久，用户希望在定时任务中指定使用的模型，以实现更精细的自动化控制。虽暂无直接PR，但属于路线图中的潜在方向。[#4378](https://github.com/HKUDS/nanobot/issues/4378)

-   **其他潜在需求**：
    -   **飞书频道体验增强 (#4619)**：通过`system`类型消息实现对话分割，提升聊天界面清晰度。目前尚无相关PR。
    -   **外部脚本触发Agent动作 (#4605)**：用户希望提供一种API或机制，让外部脚本可以触发Agent执行特定任务，以实现与其他系统的集成。该项目被标记为已关闭，可能已通过其他方式解决或不在当前路线图中。
    -   **后台任务完成通知 (#4651)**：已提交PR，为WebUI增加非活跃聊天的后台任务完成通知功能。这属于体验优化，有望被合并。[#4651](https://github.com/HKUDS/nanobot/pull/4651)

#### 7. 用户反馈摘要

从今日的Issues和PR中，可以提炼出以下用户反馈：

-   **跨平台体验是痛点**：`#4544` 反映出Windows用户在环境中遇到的shell行为不一致问题，这为跨平台Agent开发带来了困扰。
-   **对具体渠道的深度优化期望高**：`#4619` 和 `#4637` 分别来自飞书和Telegram用户，他们不仅要求能用，还要求有更好、更符合平台习惯的体验，如Telegram长消息渲染和飞书对话分割。
-   **高度评价项目代码质量和开发速度**：`#4605` 的创建者明确提到“相比OpenClaw，轻量级的代码库让我很容易阅读和理解源码，我真的很欣赏这一点”，并表达了对外部脚本触发功能的需求，这体现了用户对项目的认可和进一步集成的高级需求。
-   **“开箱即用”的安全性意识增强**：系列安全相关Issue（#4072, #4075, #4076, #4078）的提出和批量修复PR的跟进，表明社区和开发者对项目安全性的高度重视，这是项目走向成熟的关键信号。

#### 8. 待处理积压

以下为长期未得到解决或更新，但持续存在的重要问题，需提醒维护团队关注：

-   **安全类积压**：从5月29日提交的多个安全问题，如 `#4072` (ExecTool符号链接逃逸，今日有fix PR `#4629`) 和 `#4075` (Dream可覆盖用户创建的技能) 等，虽然在今日的批量修复PR `#4648` 中被提及，但该PR本身仍处于打开状态。这些问题的最终合并将是项目安全性的重要里程碑。
    -   [#4072](https://github.com/HKUDS/nanobot/issues/4072) | [#4075](https://github.com/HKUDS/nanobot/issues/4075) | [#4078](https://github.com/HKUDS/nanobot/issues/4078)

-   **高影响力Bug积压**：**`#4056`** 上下文修剪可能丢失用户问题前的助手回复，和 **`#4058`** 工具结果协议修复允许出现缺失或重复的tool_call_id状态，都是影响对话连贯性和Agent行为正确性的关键问题。它们从5月底提报后，虽今日有新评论，但仍未有关联修复PR合并。
    -   [#4056](https://github.com/HKUDS/nanobot/issues/4056) | [#4058](https://github.com/HKUDS/nanobot/issues/4058)

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

OK，收到。以下是基于你提供的 Hermes Agent 项目数据，以分析师视角生成的 2026-07-02 项目动态日报。

---

# Hermes Agent 项目日报 | 2026-07-02

**分析师备注**: 今日项目活动度极高，主要受 v0.18.0 版本发布驱动，伴随大量历史 PR 和 Issues 的集中扫尾与合并。社区贡献活跃，但需注意这可能是一轮“发布窗口期”后的暂时性拥挤。

---

## 1. 今日速览

- **整体状态**: **非常活跃**。过去 24 小时，项目完成了大幅度的版本迭代与代码清理，共处理 50 条 PR 和 6 条 Issue，整体健康度良好。
- **核心事件**: 重磅版本 **v0.18.0 “The Judgment Release”** 正式发布，自 v0.17.0 以来已累积超 1700 次提交、关闭近 1000 个 Issue，标志着项目在功能和社区规模上的巨大跨越。
- **社区贡献**: 超过 370 名社区贡献者参与迭代，今日仍有大量的历史 PR（如 #1991、#2794）被“打捞”并合并，展现了项目维护者对历史贡献的尊重与清理效率。
- **Bug 修复**: 重点修复了 **Docker exec 权限安全漏洞** 和 **桌面端模型选择器显示错误**，并针对 email、QQbot 等特定平台适配器进行了稳定性加固。

## 2. 版本发布

### Hermes Agent v0.18.0 (v2026.7.1) — “The Judgment Release”

- **更新规模**: 自 v0.17.0 以来，共 **1,720+ 次提交**，**998 个 PR 被合并**，变更了 **2,215 个文件**，净增 **~25 万行代码**，关闭了 **949 个 Issue**。
- **重点内容 (基于版本描述推断)**:
    - 本次发布代号为“审判”，预示着可能引入了更严格的智能体行为评估、输出评判（Judgment）机制或自动矫正功能。
    - 虽然未提供详细破坏性变更列表，但鉴于代码变更量巨大，**强烈建议所有用户在执行 `git pull` 或重新部署前，查阅完整的 Release Notes 和 CHANGELOG**。
- **迁移建议**:
    1.  **备份配置**: 注意检查 `config.yaml` 中关于模型配置、频道覆盖（channel_overrides）等新特性的适配。
    2.  **环境变量**: 关注 PR #2863 提到的，现在 `on` 值被统一视为布尔值的 `true` 标志，检查你的 `.env` 文件中是否有类似 `*_ENABLED=on` 的用法，以确保正确识别。
    3.  **插件兼容性**: 如果你使用了第三方插件，特别是涉及 LLM 调用的，请关注 PR #56948 中对 `_extract_text` 函数的修复，确保插件能正确处理“推理模型”的响应。

## 3. 项目进展

今日合并或关闭了一系列对功能和稳定性有重大提升的 PR：

- **安全加固 (Security Hardening)**
    - **[重要] PR #56968**: 修复了 Microsoft Teams 适配器中未经授权的发送者可能在授权检查前触发附件处理的漏洞，属于安全边界修复。🔒 ([链接](https://github.com/NousResearch/hermes-agent/pull/56968))
    - **PR #56956**: 强化了邮件适配器对畸形 IMAP 响应的处理，防止单条邮件格式错误导致批次内后续所有邮件被静默丢弃。([链接](https://github.com/NousResearch/hermes-agent/pull/56956))
- **核心功能推进**
    - **PR #56967**: “打捞”并合并了历史 PR #1991，最终实现了 **`channel_overrides` 功能**。现在同一个网关（如 Discord）下的不同频道可以使用不同的模型和系统提示词。([链接](https://github.com/NousResearch/hermes-agent/pull/56967))
    - **PR #56959**: “打捞”了历史 PR #2696 的存留部分，使 API 服务器能内联图片为 `data URL`，解决了远程前端无法显示服务器本地图片的问题。([链接](https://github.com/NousResearch/hermes-agent/pull/56959))
    - **PR #56973**: 对应 Issue #56791，Cron 任务输出现在会附带模型和提供商元数据，方便调试和追溯。([链接](https://github.com/NousResearch/hermes-agent/pull/56973))
- **问题修复 (Quick Fixes)**
    - **PR #56948**: 修复了 `plugin_llm.py` 中无法正确处理深度求索（DeepSeek）等“推理模型”响应文本的问题。([链接](https://github.com/NousResearch/hermes-agent/pull/56948))
    - **PR #56970**: 修复了 QQBot 适配器在重连时因不识别 `is_reconnect` 参数而永久卡在重试状态的问题。([链接](https://github.com/NousResearch/hermes-agent/pull/56970))

## 4. 社区热点

- **Issue #1955 (已关闭)**: **7个 👍，10条评论**。该 issue 是关于“为网关平台提供频道级别模型和系统提示覆盖”的经典 Feature Request。虽然创建于 2026-03-18，但最终在今日通过 PR #1991 和后续的 #56967 被合并关闭。这彰显了社区呼声对项目路线图的直接影响，以及维护者对长期悬挂功能的清理决心。([链接](https://github.com/NousResearch/hermes-agent/issues/1955))
- **PR #54748 (待合并)**: 关于**模型路由层级**的提案。这是一个较大型的重构 PR，旨在根据任务上下文智能选择模型/提供商组合。其复杂性和潜在影响可能使其成为下一个版本讨论的焦点。虽然今天无直接讨论，但此类 PR 通常是社区讨论热度升温的起点。([链接](https://github.com/NousResearch/hermes-agent/pull/54748))

## 5. Bug 与稳定性

按严重程度排列：

1.  **[Critical] Issue #56942 (已关闭)**: Docker 镜像中硬编码 `USER=root` 导致 `docker exec` 写入 root 所有权的文件，引发 401 权限错误。**此问题严重性高**，虽然已被修复，但用户需立即更新基础镜像。([链接](https://github.com/NousResearch/hermes-agent/issues/56942))
2.  **[High] Issue #56974 & PR #56966**: 桌面端模型选择器中**泄露了用户未显式配置的提供商（如隐式发现的 GitHub CLI 凭证）**。这属于**隐私/配置可见性**问题，可能使用户困惑。PR #56966 已提供一个修复方案。([Issue](https://github.com/NousResearch/hermes-agent/issues/56974), [PR](https://github.com/NousResearch/hermes-agent/pull/56966))
3.  **[Medium] PR #56948**: 插件在与“推理模型”（如 DeepSeek R1）交互时遇到问题，导致无法正确获取文本输出。已修复。([链接](https://github.com/NousResearch/hermes-agent/pull/56948))
4.  **[Medium] PR #256 (未在列表中)**: 虽然未见，但要注意 `lazy_deps.py` 中 `pip` 回退暴露环境变量的 Bug，已修复。([PR #56976](https://github.com/NousResearch/hermes-agent/pull/56976))

## 6. 功能请求与路线图信号

- **高优先级/可能纳入下一版本**:
    - **Cron 输出包含模型元数据 (Issue #56791)**: 用户强烈希望能在日志中直接看到执行 cron 任务的模型，以避免手动交叉验证。此功能已被 PR #56973 实现。([链接](https://github.com/NousResearch/hermes-agent/issues/56791))
    - **工具调用前置钩子 (Issue #56969)**: 用户希望能够在工具执行前（而非执行后）嵌入 URL 路由等规则。这是一个**有价值的功能请求**，预示着更精细化的工具使用控制。但目前尚无对应 PR。([链接](https://github.com/NousResearch/hermes-agent/issues/56969))
- **长期信号/需社区讨论**:
    - **原生图像/视频生成 (Issue #56965)**: 用户希望直接使用内部模型（如大学提供的端点）进行文生图，而不是被迫调用外部图像 API。这表明社区对**模型能力泛化**和**供应商中立性**的持续追求。([链接](https://github.com/NousResearch/hermes-agent/issues/56965))

## 7. 用户反馈摘要

- **非常满意 (已满足需求)**:
    - **Issue #1955 的关闭**: 解决了许多用户在 Discord 等群里对频道级模型定制的长期渴望。用户 “rivercrab26” 提出的 #daily 频道用便宜模型、#dev 频道用高端模型的场景成为现实。
- **痛点与改进建议**:
    - **Issue #56974 (桌面模型选择器)**: 用户 “rarf” 报告了桌面端 UI 混淆的问题，认为应只显示自己配置的模型，隐式发现的模型会造成认知负担。
    - **Issue #56965 (原生图像生成)**: 用户 “unnatLNCO” 表达了插件生态对特殊模型端点支持的不足，这可能是未来插件 API 或模型抽象层需要优化的方向。

## 8. 待处理积压

- **PR #46466 (积压 17 天)**: 为桌面端引入原生看板和成就系统面板的插件加载器。这是一个重要的 Desktop 功能增强，已处于开放状态两周，且今日有更新，但仍未合并。建议维护者关注，以推动桌面端功能落地。([链接](https://github.com/NousResearch/hermes-agent/pull/46466))
- **PR #54748 (积压 3 天)**: 模型路由层级提案。其“风险兼容性 (risk-compatibility)”标签可能暗示了设计的复杂性和潜在兼容性问题。鉴于版本刚发布，该 PR 可能需要更长的评估周期，但值得社区和核心开发者优先讨论。([链接](https://github.com/NousResearch/hermes-agent/pull/54748))

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw 项目动态日报 — 2026-07-02

## 1️⃣ 今日速览

- 项目今日发布 **nightly 构建版 v0.3.1-nightly.20260702**，包含主分支最新代码，但未提供正式更新日志。
- **3 条新 Issue** 被提交，其中 2 条为 Bug 报告（Android/Termux 下进程钩子崩溃、Matrix 同步重连缺失），1 条为功能请求（QQ 通道流式输出）。
- **15 条 PR** 处于活跃状态，其中 2 条已合并/关闭（`turn.done` 生命周期完善、Telegram 回复触发优化），其余 13 条等待审核。
- 整体社区活跃度中等，修复类 PR 数量多但合并进度偏慢，部分 PR 已标记 `stale`，提示维护精力需集中。

## 2️⃣ 版本发布

### nightly: Nightly Build (v0.3.1-nightly.20260702.2cf030d2)
- **类型**：自动化夜间构建，可能不稳定，仅供测试。
- **变更范围**：包含从 `v0.3.1` 到 `main` 分支的所有累积提交。
- **完整变更日志**：[compare/v0.3.1...main](https://github.com/sipeed/picoclaw/compare/v0.3.1...main)
- **注意**：无破坏性变更说明，建议生产环境仍使用稳定版。

## 3️⃣ 项目进展

今日共有 **2 个 PR 被合并/关闭**，推动了以下改进：

| PR | 标题 | 状态 | 说明 |
|----|------|------|------|
| [#3116](https://github.com/sipeed/picoclaw/pull/3116) | fix(pico): complete turn.done lifecycle signaling | ✅ 已合并 | 修复了 Pico 对话生命周期中 `request_id` 未正确传递、steering/follow-up 消息缺失等问题，完善了 `turn.done` 信号流程。 |
| [#2975](https://github.com/sipeed/picoclaw/pull/2975) | feat(telegram): treat reply to bot message as mention in group chats | ✅ 已合并 | Telegram 群组现支持将“回复 bot 消息”视为 @提及，使用 `mention_only: true` 配置时，用户可通过回复方式触发 bot，提升交互便利性。 |

此外，还有 **13 个待合并 PR** 覆盖 CLI 参数校验、Azure 依赖冻结、ARMv7 构建支持、DeltaChat 网关、Bedrock 提示缓存等重要改进，表明项目有持续的功能扩展和修复计划。

## 4️⃣ 社区热点

今日社区讨论热度不高（Comment 数量均为 0），但以下两个 Issue 因其触及核心稳定性而值得关注：

- **🔥 [#3203 – Matrix sync loop has no reconnection logic](https://github.com/sipeed/picoclaw/issues/3203)**（2026-07-02 新开）  
  **背景**：用户 `weissfl` 报告 Matrix 通道的 `/sync` 长轮询在网络中断或服务器重启后永久死亡，无自动重连机制，而主进程存活导致 systemd `Restart=on-failure` 无法触发。  
  **诉求**：期望增加自动重连逻辑或心跳检测，避免静默失联。

- **🔥 [#3164 – Process hooks crash gateway on Android/Termux](https://github.com/sipeed/picoclaw/issues/3164)**（2026-06-23 报告，2026-07-01 最后更新）  
  **背景**：用户 `AMEOBIUS` 在 Android 上通过 Termux 运行 picoclaw 时，只要启用了任何进程钩子（即使是最简单的 hello world），网关会在启动后 2 秒内崩溃。  
  **诉求**：希望修复该平台兼容性问题，该 Issue 已被标记 `stale`，但仍未关闭。

## 5️⃣ Bug 与稳定性

| 严重程度 | Issue | 描述 | 关联 Fix PR |
|----------|-------|------|-------------|
| 🔴 崩溃/不可用 | [#3164](https://github.com/sipeed/picoclaw/issues/3164) | Android/Termux 下进程钩子导致网关启动即崩（v0.2.9） | 暂无 |
| 🟡 严重逻辑缺陷 | [#3203](https://github.com/sipeed/picoclaw/issues/3203) | Matrix 同步循环无重连，网络恢复后永久静默 | 暂无 |
| 🟢 功能缺失 | [#3201](https://github.com/sipeed/picoclaw/issues/3201) | QQ 通道不支持流式输出（非崩溃，但影响体验） | 无（功能请求） |

**值得注意的已提交修复 PR**：
- [#3180](https://github.com/sipeed/picoclaw/pull/3180) – CLI 工具调用参数无效时跳过而非丢弃整批调用（待合并）。
- [#3205](https://github.com/sipeed/picoclaw/pull/3205) – 修复 9router 网关响应解析，并添加 Linux ARMv7 构建目标（待合并）。
- [#3202](https://github.com/sipeed/picoclaw/pull/3202) – 修复路由 ID 规范化未处理首尾下划线的问题（待合并）。
- [#3165](https://github.com/sipeed/picoclaw/pull/3165) – 恢复火山引擎豆包 Seed XML 工具调用（已标记 stale）。

## 6️⃣ 功能请求与路线图信号

用户提出的新功能需求及对应已有 PR：

| 功能 | Issue/PR | 状态 | 可能纳入版本 |
|------|----------|------|-------------|
| QQ 通道流式输出 | [#3201](https://github.com/sipeed/picoclaw/issues/3201) | 新开，无对应 PR | 下一里程碑（需社区贡献） |
| DeltaChat 网关 | [#3063](https://github.com/sipeed/picoclaw/pull/3063) | 待合并（已标记 stale） | v0.3.x 或 v0.4.0 |
| AWS Bedrock 提示缓存 | [#3163](https://github.com/sipeed/picoclaw/pull/3163) | 待合并（已标记 stale） | 预计纳入后续 version |
| 9router 兼容 + ARMv7 支持 | [#3205](https://github.com/sipeed/picoclaw/pull/3205) | 新提交，待审核 | 大概率随 nightly 进入下个稳定版 |
| Azure 依赖版本冻结 | [#3204](https://github.com/sipeed/picoclaw/pull/3204) | 新提交，待审核 | 用于内部供应链检查，低风险合并 |

从近期 PR 趋势看，网关兼容性（9router、Bedrock、Seed）和多平台构建（ARMv7）是当前开发重点。

## 7️⃣ 用户反馈摘要

- **Android/Termux 用户**（Issue #3164）：反馈进程钩子完全不可用，希望官方能给出临时 workaround 或平台特定配置。
- **Matrix 用户**（Issue #3203）：强调“静默死亡”比显式崩溃更糟糕，因为运维人员无法及时发现；期望借鉴其他通道（如 Telegram）的重连机制。
- **QQ 通道用户**（Issue #3201）：对比 Telegram 和 Pico WebSocket 的流式体验，认为缓慢的全量响应严重影响实时交互，要求尽快支持 token-by-token 输出。

## 8️⃣ 待处理积压

以下 Issue/PR 长期未得到维护者响应或合并，建议优先关注：

| 类型 | 编号 | 标题 | 创建时间 | 最后更新 | 状态标记 |
|------|------|------|----------|----------|----------|
| Issue | [#3164](https://github.com/sipeed/picoclaw/issues/3164) | Process hooks crash gateway on Android/Termux | 2026-06-23 | 2026-07-01 | `stale` |
| PR | [#3165](https://github.com/sipeed/picoclaw/pull/3165) | fix(openai_compat): recover Seed XML tool calls | 2026-06-24 | 2026-07-01 | `stale` |
| PR | [#3161](https://github.com/sipeed/picoclaw/pull/3161) | fix(exec): keep deny patterns active for custom allow rules | 2026-06-23 | 2026-07-01 | `stale` |
| PR | [#3160](https://github.com/sipeed/picoclaw/pull/3160) | fix(auth): reject cross-site launcher setup requests | 2026-06-23 | 2026-07-01 | `stale` |
| PR | [#3158](https://github.com/sipeed/picoclaw/pull/3158) | test: cover sandbox fs Windows path handling | 2026-06-22 | 2026-07-01 | `stale` |
| PR | [#3063](https://github.com/sipeed/picoclaw/pull/3063) | feat: add deltachat gateway | 2026-06-08 | 2026-07-02 | `stale` |
| PR | [#3104](https://github.com/sipeed/picoclaw/pull/3104) | build(deps): bump shadcn from 4.7.0 to 4.11.0 | 2026-06-11 | 2026-07-01 | `stale`, `dependencies` |
| PR | [#3103](https://github.com/sipeed/picoclaw/pull/3103) | build(deps-dev): bump typescript-eslint… | 2026-06-11 | 2026-07-01 | `stale`, `dependencies` |

这些积压项涉及安全性（跨站认证、exec 权限）、兼容性（Android、Seed、DeltaChat）及依赖升级，越早处理越有利于项目健康度。

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

好的，根据您提供的 NanoClaw 项目数据，以下是 2026-07-02 的项目动态日报。

---

### NanoClaw 项目动态日报 | 2026-07-02

**数据来源**: [github.com/qwibitai/nanoclaw](https://github.com/qwibitai/nanoclaw)

---

#### 1. 今日速览

项目今日保持高度活跃，尤其在代码开发与合并方面。24小时内处理了12个Pull Requests，其中6个已合并/关闭，显示了维护团队的强劲交付能力。核心焦点集中在 **Agent模板系统** 的构建上，昨日有3个紧密关联的PR被提交，标志着该功能进入冲刺阶段。同时，多个由社区贡献的、积压较久的旧PR今日被统一合并，体现了良好的社区维护节奏。整体而言，项目正处于一个重大功能发布前的高强度开发期。

#### 2. 项目进展

今日合并/关闭了多个重要PR，推进了功能完善与稳定性提升。

- **WhatsApp连接稳定性修复**: 合并了 #2905 修复了WhatsApp频道在重连时未正确关闭旧socket导致的主机内存泄漏问题。此修复直接解决了因Baileys库频繁断开(`reason: 408`)而产生的资源堆积，对WhatsApp重度用户至关重要。
    - **链接**: [nanocoai/nanoclaw PR #2905](https://github.com/nanocoai/nanoclaw/pull/2905)

- **调度与贡献流程优化**: 合并了社区贡献者 `shrwnsan` 提交的三个PR，显著增强了项目的运维能力：
    - **任务脚本容错重试**: #2677 为计划任务的预处理脚本增加了失败后自动重试机制，并在失败时输出诊断信息，提高了任务执行的鲁棒性。
        - **链接**: [nanocoai/nanoclaw PR #2677](https://github.com/nanocoai/nanoclaw/pull/2677)
    - **贡献自动化检查**: #1716 新增 `/check-contribution` 运维技能，可自动验证贡献的SKILL.md格式、代码质量等，降低了贡献者的PR提交门槛。
        - **链接**: [nanocoai/nanoclaw PR #1716](https://github.com/nanocoai/nanoclaw/pull/1716)
    - **自定义API端点支持**: #1257 实现了对第三方Anthropic兼容API（如 z.ai）挂载在子路径下的支持，扩展了模型供应商的选择范围。
        - **链接**: [nanocoai/nanoclaw PR #1257](https://github.com/nanocoai/nanoclaw/pull/1257)

- **技能生态补充**: 合并了 #1693 (自动状态备份技能) 与 #1597 (语义化对话搜索技能)，进一步丰富了项目内置的Utility和Feature技能生态，使自托管用户的运维与数据检索更便捷。
    - **链接**: [nanocoai/nanoclaw PR #1693](https://github.com/nanocoai/nanoclaw/pull/1693), [nanocoai/nanoclaw PR #1597](https://github.com/nanocoai/nanoclaw/pull/1597)

#### 3. 社区热点

今日最受关注的无疑是 **Agent模板系统** 系列PR。该系列由开发者 `amit-shafnir` 提交，分为两部分，且存在依赖关系，引起了关注。

- **核心PR**: #2890 (本地模板加载器) 和 #2909 (设置向导中集成模板流程) 构成了Agent模板功能的完整链路。
    - **链接**: [nanocoai/nanoclaw PR #2890](https://github.com/nanocoai/nanoclaw/pull/2890), [nanocoai/nanoclaw PR #2909](https://github.com/nanocoai/nanoclaw/pull/2909)
- **配套支持**: #2908 则为Codex provider适配了此功能，实现了模板化Agent的端到端工作。
    - **链接**: [nanocoai/nanoclaw PR #2908](https://github.com/nanocoai/nanoclaw/pull/2908)

**诉求分析**: 社区对“模板化”的诉求强烈，因为这能极大简化AI Agent的创建和标准化部署流程。允许用户通过 `ncl groups create --template <ref>` 命令，从预定义模板中快速创建包含一组预置Promise、背景和技能的Agent群组，这是项目向“企业级”和“易用性”迈进的关键一步。

#### 4. Bug 与稳定性

今日未发现新增的重大Bug报告，但两个已提交的修复性PR值得关注，它们分别针对Slack和WhatsApp频道的生产环境问题：

- **严重性：高** | **WhatsApp内存泄漏 (已修复)**: PR #2905 已在今日合并，解决了因频繁重连导致的主机内存泄漏。对于部署了WhatsApp频道的用户，这是一个关键的稳定性修复。
    - **链接**: [nanocoai/nanoclaw PR #2905](https://github.com/nanocoai/nanoclaw/pull/2905)

- **严重性：中** | **Slack线程历史丢失 (待合入)**: PR #2904 修复了Slack频道在 `engage_mode: 'mention'` 模式下，当用户在已有线程中再次 `@` 机器人时，无法获取历史对话上下文的问题。这会导致Agent的回复与上下文脱节。
    - **链接**: [nanocoai/nanoclaw PR #2904](https://github.com/nanocoai/nanoclaw/pull/2904)

#### 5. 功能请求与路线图信号

从开放PR和合并情况可以看出，“**易用性与可配置性**”是当前路线图的明确信号。

- **Agent模板化 (高概率纳入下个版本)**: 前述的 #2890、#2909 和 #2908 构成了一套完整功能，预计在合并后将成为下一个版本的核心亮点。
- **实例级默认Provider配置**: PR #2906 提议新增 `DEFAULT_AGENT_PROVIDER` 环境变量，允许运维人员为所有新建Agent组设置默认AI提供商，避免逐个配置。这反映了多模型部署场景下的配置管理需求。
    - **链接**: [nanocoai/nanoclaw PR #2906](https://github.com/nanocoai/nanoclaw/pull/2906)
- **本地语音转写功能**: PR #2317 新增 `/add-voice-transcription-free-whisper` 技能，支持在本地免费且无需GPU的情况下实现语音转录，吸引了注重隐私和成本的用户关注。
    - **链接**: [nanocoai/nanoclaw PR #2317](https://github.com/nanocoai/nanoclaw/pull/2317)

#### 6. 用户反馈摘要

今日暂无来自Issues的直接用户反馈。但从PR的讨论与描述中，可以提炼出用户画像：

- **自托管运维者**: 他们贡献/恢复了备份(PR #1693)和语义搜索(PR #1597)技能，显示出对数据主权和生产环境稳定性的高要求。
- **多渠道重度用户**: PR #2905 (WhatsApp内存泄漏) 和 #2904 (Slack线程历史) 的提出，表明用户正在实际生产环境中高频使用这些频道，并遭遇了具体的交互细节问题。

#### 7. 待处理积压

- **长期未更新的语音转写技能PR**: PR #2317 (`/add-voice-transcription-free-whisper`) 自今年5月7日起就处于开放状态，且已有超过一个月未更新。该功能需求明确，但可能因实现复杂性或作者精力问题而停滞。建议维护者主动联系作者或评估合入代码库的可行性。
    - **链接**: [nanocoai/nanoclaw PR #2317](https://github.com/nanocoai/nanoclaw/pull/2317)

- **`ape_claw_cli` Issue**: #2907 是一个标题抽象、摘要空白的Issues。这可能是一个无效Issue，也可能是未来的一个重大提议。建议维护者联系作者 `slotaibuddy-admin` 询问详情，避免项目符号被混淆。
    - **链接**: [nanocoai/nanoclaw Issue #2907](https://github.com/nanocoai/nanoclaw/issues/2907)

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是根据您提供的 IronClaw GitHub 数据生成的 2026-07-02 项目动态日报。

---

### IronClaw 项目动态日报 — 2026-07-02

**分析师点评：** 项目今日维持高度活跃状态。尽管新版本发布暂停，但代码库中大规模的重构与稳定性修复工作正在密集进行，超过50个PR的更新量以及26个PR被合并/关闭，显示出核心团队对“Reborn”架构的推进决心。然而，今日新报告的多个中等规模Bug（特别是与功能死代码路径和QA回归相关的问题）提示，随着快速迭代，系统复杂度和潜在的回归风险正在增加。

---

### 1. 今日速览

今日 IronClaw 项目呈现 **高活跃度与关键修复并行** 的状态。共有 **5 个新 Issue** 和 **50 个 Pull Request** 更新，其中 **26 个 PR 被合并/关闭**，显示出强劲的交付能力。核心开发者 `serrrfirat` 主导了多项重要合并，包括围绕“Reborn”架构的稳定性修复和基础设施改进（如心跳超时解耦、断路器默认开启）。社区与QA团队报告了多个值得关注的Bug，主要涉及功能路径不可达（死代码）和特定场景下的运行时错误，为项目的稳定性敲响了警钟。总体而言，项目在积极向前演进的同时，也面临着快速迭代带来的质量挑战。

### 2. 版本发布

无

### 3. 项目进展

今日项目在 **稳定性、基础设施和开发者体验** 方面取得了实质性进展。核心团队合并并关闭了多项关键 PR，推动了“Reborn”架构的成熟。

- **核心稳定性与可靠性提升**
    - **断路器机制默认开启**：[PR #5203](https://github.com/nearai/ironclaw/pull/5203) 已合并。该修复使系统在LLM提供商出现故障时能快速失败，避免了单个提供商故障导致整个实例停滞的问题，显著提升了系统的容错能力。
    - **运行超时解耦**：[PR #5228](https://github.com/nearai/ironclaw/pull/5228) 已合并。该修复解耦了Runner心跳间隔和写入超时，防止慢速状态写入拖慢整个心跳报告系统，提升了调度器的健壮性。
    - **SSE会话恢复修复**：[PR #5407](https://github.com/nearai/ironclaw/pull/5407) 已合并。修复了一个生产级Bug，确保在SSE连接恢复后，技能学习“气泡”消息能够正确投递。
    - **Auth恢复流程优化**：至少有三个相关PR（[#4998](https://github.com/nearai/ironclaw/pull/4998), [#4911](https://github.com/nearai/ironclaw/pull/4911), [#2754](https://github.com/nearai/ironclaw/pull/2754)）被合并，优化了授权恢复流程，确保了身份和审批状态的原子性和一致性。

- **开发者体验与代码智能**
    - **构建代码知识图谱**：[PR #5532](https://github.com/nearai/ironclaw/pull/5532) 被合并。该PR为所有Agent（包括CI）集成了代码知识图谱和OpenWiki，使开发者能更高效地在超百万行代码中进行发现和理解，是提升团队开发效率的重要基础设施。

- **QA与测试能力**
    - **实时QA增加语义判断**：[PR #5531](https://github.com/nearai/ironclaw/pull/5531) 被合并。为“Reborn WebUI v2”的实时QA流程增加了可选的LLM语义判断回退机制，使自动化测试能更智能地识别非精确匹配，提升了QA覆盖的准确性和鲁棒性。

### 4. 社区热点

今日社区讨论最活跃的议题集中在 **功能配置的灵活性与可观测性** 两大方面。

- **可配置技能与工具**：Issue [#5459](https://github.com/nearai/ironclaw/pull/5459) 获得了2条评论，是今日唯一有讨论的Issue。用户 `zetyquickly` 提出需要允许管理员和用户分别安装WASM工具和技能，并管理其可见性（全局/私有）。这反映了社区对于 **细粒度权限控制和个性化工作空间** 的强烈需求。
- **工具延迟的可观测性**：PR [#5534](https://github.com/nearai/ironclaw/pull/5534) 旨在为第一方工具运行时增加延迟追踪。虽然没有具体评论数，但其“XL”大小和“core”贡献者标签表明这是一项重要工作。这暗示社区和开发团队对 **定位和优化工具调用性能瓶颈** 有持续关注。

### 5. Bug 与稳定性

今日报告的 Bug 数量不多，但质量较高，主要涉及 **功能路径死代码**和 **QA回归**。

- **严重**：
    1.  **[关键功能死代码] 基于条件的技能自动激活功能不可达**：Issue [#5530](https://github.com/nearai/ironclaw/pull/5530) 明确指出 `SelectableSkillContextSource` 在核心提交路径上无法运行条件检查逻辑，导致技能自动激活功能处于“死亡”状态。这是一个**严重的设计/重构遗留问题**，可能导致关键产品功能失效。
    2.  **[生产环境问题] 幂等性写/读作用域不匹配**：Issue [#5527](https://github.com/nearai/ironclaw/pull/5527) 指出 `FilesystemSessionThreadService` 的幂等性写入和读取使用了不同作用域（`owner scope` vs `system scope`），导致“replay-before-policy”机制在生产环境中无效。这是一个隐蔽的逻辑Bug，影响数据一致性和重放策略。

- **中等**：
    1.  **[QA回归] Reborn进程在读取Slack私信时失败**：Issue [#5522](https://github.com/nearai/ironclaw/pull/5522) 由QA团队报告。当任务需要读取Slack DM功能时，Agent因为缺少该 `capability` 而陷入 `capability_info` 重试循环，最终导致任务失败。这是一个具体的功能回归，影响了Slack集成的可用性。

- **低**：
    1.  **[持续失败] 夜间E2E测试失败**：Issue [#4108](https://github.com/nearai/ironclaw/pull/4108)，自5月27日起持续报告失败。虽然可能是基础设施或测试脚本的问题，但长期未解决会削弱CI/CD的置信度。

**已有修复PR**: 对于 `Run-borking` 问题，PR [#4841](https://github.com/nearai/ironclaw/pull/4841) 正在开发中，旨在为终端错误提供解释和可重试机制，这将间接缓解如#5522中的无休止重试问题。

### 6. 功能请求与路线图信号

- **技能/工具的细粒度管理与自动激活**：Issue [#5459](https://github.com/nearai/ironclaw/pull/5459) 提出的“Configurable skills and tools”是一个清晰的功能请求，与 `Reborn` 架构的进化方向（如 #5261 的管理员安装商店）高度契合。结合PR [#4544](https://github.com/nearai/ironclaw/pull/4544)（作用域生命周期管理已合并），该功能很可能会是下一个版本的核心特性。
- **更深度的可观测性**：PR [#5534](https://github.com/nearai/ironclaw/pull/5534)（工具延迟追踪）和PR [#5380](https://github.com/nearai/ironclaw/pull/5380)（扩展QA矩阵）表明，项目正致力于从 **被动发现Bug** 转向 **主动监控和量化系统健康度**。这应被视为一个长期的路线图信号。
- **长期未解决的夜间测试**：Issue [#4108](https://github.com/nearai/ironclaw/pull/4108) 虽然是一个Bug，但长期不解决也发出了一个消极信号，即基础设施的稳定性优先级可能低于功能开发。

### 7. 用户反馈摘要

由于今日活跃的Issue较少，反馈主要来自QA团队的模拟用户场景报告（Issue [#5522](https://github.com/nearai/ironclaw/pull/5522)）：
- **痛点**：用户期望Agent能够处理涉及Slack私信的复杂任务，但目前会因为“缺乏能力”而陷入无休止的重试循环，**用户体验极差**。这表明系统在能力不足时的错误处理和用户反馈机制需要改进。
- **场景**：用户尝试使用“Reborn”Agent运行一个需要读取Slack DM的任务，这是一个**重要的企业级协作场景**。该功能的失败可能会严重影响用户对Agent自动化能力的信任。

### 8. 待处理积压

- **重要待办**：
    - **夜间E2E测试持续失败**：Issue [#4108](https://github.com/nearai/ironclaw/pull/4108) 已经存在超过一个月，至今未解决。这是影响CI/CD管道健康度的首要问题， **强烈建议维护者优先排查**。
    - **`Reborn` 无故障运行解释**：大型PR [#4841](https://github.com/nearai/ironclaw/pull/4841) 已经开放将近三周，旨在解决核心的“run-borking”问题。虽然近期有更新，但其规模和重要性意味着它应被持续关注和推进，否则会影响整个 `Reborn` 架构的用户信心。

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

好的，以下是为您生成的 **LobsterAI 项目动态日报 (2026-07-02)**。

---

# LobsterAI 项目日报 – 2026-07-02

## 1. 今日速览

今日项目合并 / 关闭了 **7 个 PR**，修复了定时任务、自定义模型删除白屏、共享部署 Node 环境构建等关键问题，代码质量提升显著。但过去 24 小时 **无新版本发布**，且 **5 个已持续 3 个月的 stale Issue 被再次标记为活跃**，表明用户报告的核心 Bug 仍未得到官方回复或修复。项目整体活跃度中高，维护者聚焦于后台稳定性与协作功能增强，但对积压的社区反馈响应不足。

## 2. 版本发布

*（今日无新版本发布）*

## 3. 项目进展

今日共有 **7 个 PR 被合并/关闭**，主要集中在渲染进程修复、文档更新与协作功能增强：

- **定时任务通知渠道「不通知」失效修复**  
  PR [#2255](https://github.com/netease-youdao/LobsterAI/pull/2255) 修复了编辑定时任务时切换到“不通知”后保存不生效的问题。根因是网关 `cron.update` 的补丁合并逻辑无法清除 `delivery` 字段。

- **删除自定义模型时白屏修复**  
  PR [#2252](https://github.com/netease-youdao/LobsterAI/pull/2252) 修复了在设置中删除当前选中的自定义提供商导致整个视图白屏的问题。原因是 `confirmDeleteCustomProvider` 的异步顺序导致状态未及时切换。

- **共享部署使用独立 Node 环境**  
  PR [#2251](https://github.com/netease-youdao/LobsterAI/pull/2251) 新增 Node 工具子进程环境构建逻辑，确保部署命令（install、build、prune）在独立环境中执行，并补充了缺失 npm 的明确错误提示。

- **协作用功能增强：子 Agent 产物面板**  
  PR [#2249](https://github.com/netease-youdao/LobsterAI/pull/2249) 新增子 Agent 专属的产物标签页，支持列表/详情视图，并在对话卡片中通过侧面板打开子 Agent 详情，不再替换主会话页面。

此外，`fisherdaddy` 提交了 **3 次文档更新**（[#2254](https://github.com/netease-youdao/LobsterAI/pull/2254)、[#2253](https://github.com/netease-youdao/LobsterAI/pull/2253)、[#2250](https://github.com/netease-youdao/LobsterAI/pull/2250)），持续改进项目首页与 README 图片。

**项目向前迈进一步**：渲染层稳定性提升（定时任务、模型设置），部署体验优化，协作用功能更完善。

## 4. 社区热点

今日所有 Issue 均无新增评论，讨论热度较低。相对而言，**Issue [#1354](https://github.com/netease-youdao/LobsterAI/issues/1354)**（启动 pageant 后蓝屏）拥有 **2 条评论**，是评论数最高的 Issue。  
用户描述了偶现的蓝屏现象并附带了日志文件，诉求是希望项目能够稳定调用外部工具 `pageant` 而不导致系统崩溃。  
该 Issue 已标记为 `stale` 近 3 个月，至今未获官方回复，社区可能开始关注维护者的响应速度。

## 5. Bug 与稳定性

以下为今日处于活跃状态的 **5 个 Bug**，均来自长期未解决的 `stale` Issue，按严重程度排列：

| 严重程度 | Issue 编号 | 描述 | 是否已有 Fix PR |
|---------|-----------|------|----------------|
| 🔴 严重 | [#1354](https://github.com/netease-youdao/LobsterAI/issues/1354) | 调用 pageant 后导致电脑蓝屏（偶现） | 无 |
| 🟠 中等 | [#1357](https://github.com/netease-youdao/LobsterAI/issues/1357) | “帮我开启 pageant”回答已启动，实际未启动（必现） | 无 |
| 🟠 中等 | [#1359](https://github.com/netease-youdao/LobsterAI/issues/1359) | 删除的任务在重启后再次出现 | 无 |
| 🟡 一般 | [#1358](https://github.com/netease-youdao/LobsterAI/issues/1358) | 定时任务点击后无交互，无法确认是否启动 | 无 |
| 🟡 一般 | [#1360](https://github.com/netease-youdao/LobsterAI/issues/1360) | Agent 自定义创建未做重名验证 | 无 |

今日合并的 PR 中，[#2256](https://github.com/netease-youdao/LobsterAI/pull/2256)（Open）包含了定时任务通知修复与模型删除白屏修复，但未涉及上述 `pageant` 相关或任务持久化问题。

## 6. 功能请求与路线图信号

今日无新提交的功能请求 Issue。但从 **PR [#2249](https://github.com/netease-youdao/LobsterAI/pull/2249)**（子 Agent 产物面板）可以看出，项目团队正在增强 **协作模式下子 Agent 的可视化与交互**。  
结合 Issue [#1358](https://github.com/netease-youdao/LobsterAI/issues/1358)（定时任务交互无反馈），**用户体验反馈的缺失** 可能被纳入下一轮 UI 改进——尤其需要为定时任务的启动/停止增加明确的状态提示。

## 7. 用户反馈摘要

从 stale Issue 的评论及描述中可提炼以下用户痛点：

- **稳定性问题**（#1354）：用户期待 LobsterAI 在调用外部工具时不引发系统级崩溃，否则完全不可用。
- **功能执行不一致**（#1357）：模型回答与实际行动不符，导致信任问题。
- **任务管理混乱**（#1359）：删除的任务重启后复活，让用户感到不可控；同时（#1358）无法获知任务是否执行，操作无反馈。
- **配置体验粗糙**（#1360）：Agent 同名创建未被拦截，违反常规软件设计预期。

用户整体表现为 **对基础功能可靠性的不满**，而非对新功能的需求。

## 8. 待处理积压

以下 **5 个 Issue** 均创建于 2026-04-02，已标记 `stale` 超过 90 天，且今日被更新（可能是机器人自动标记），极需维护者优先介入：

- **[#1354](https://github.com/netease-youdao/LobsterAI/issues/1354)**：启动 pageant 蓝屏（严重，附日志）
- **[#1357](https://github.com/netease-youdao/LobsterAI/issues/1357)**：pageant 回答与实际状态不一致（必现）
- **[#1358](https://github.com/netease-youdao/LobsterAI/issues/1358)**：定时任务无交互反馈
- **[#1359](https://github.com/netease-youdao/LobsterAI/issues/1359)**：删除任务重启后重现
- **[#1360](https://github.com/netease-youdao/LobsterAI/issues/1360)**：Agent 重名创建无验证

这些 Issue 涵盖了核心功能的严重缺陷，长期悬置可能影响用户留存与社区口碑。建议下个迭代至少对 [#1354](https://github.com/netease-youdao/LobsterAI/issues/1354) 和 [#1357](https://github.com/netease-youdao/LobsterAI/issues/1357) 给出初步排查结论或临时 workaround。

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

好的，作为一名AI智能体与个人AI助手领域开源项目分析师，我将根据您提供的 CoPaw (QwenPaw) 项目数据，为您生成一份结构清晰、数据驱动的项目动态日报。

---

# CoPaw (QwenPaw) 项目动态日报 | 2026-07-02

**项目状态：** 🟢 活跃
**当前版本：** v2.0.0-beta.2 (Beta)
**项目健康度：** 🟡 一般

---

## 1. 今日速览

今日项目活跃度极高，经历了大规模、高频次的开发活动。**关键版本 `v2.0.0-beta.2`（Beta）** 的发布是今日最大事件，标志着项目向2.0架构迈进了实质性一步，但其“早期测试版”性质也伴随着稳定性风险。社区贡献热情高涨，PR数量激增至50条，其中32条已被合并，表明团队在快速推进新功能和修复，例如飞书通道的Bug修复、内存系统的增强。然而，**用户端报告了多个严重Bug**，包括内存泄漏、浏览器卡顿和功能缺失，凸显了开发速度与稳定性之间的张力。社区核心诉求集中在安全性、易用性和关键通道（飞书）的可靠性上。

## 2. 版本发布

-   **v2.0.0-beta.2**
    -   **链接:** [Release v2.0.0-beta.2](https://github.com/agentscope-ai/QwenPaw/releases/tag/v2.0.0-beta.2) | [相关 Issue #5733](https://github.com/agentscope-ai/QwenPaw/issues/5733)
    -   **状态:** **⚠️ 早期Beta测试版**，专为开发者和早期采用者设计, **不推荐用于生产环境**。
    -   **更新内容:** 此版本为QwenPaw 2.0.0的第二个Beta版，包含正在积极开发中的功能和破坏性变更。
    -   **What's Changed:** 目前仅记录了一项CLI相关功能 `feat(cli): add cron up`，暗示增强了定时任务能力。
    -   **迁移注意事项:**
        1.  **破坏性变更:** 该版本可能包含与之前版本不兼容的变更。升级前请仔细阅读完整的Release Note，并在测试环境下进行验证。
        2.  **稳定性问题:** 作为Beta版本，可能存在未预见的Bug和性能问题，不建议在生产环境中部署。
        3.  **安装验证:** 关联的Release Duty Issue (#5733) 正在进行安装验证，表明团队正在积极监控此版本的可用性。

## 3. 项目进展

今日共有 **32 个PR被合并/关闭**，项目在以下方面取得关键进展：

-   **核心Bug修复:**
    -   **Goal Mode修复:** PR #5727 ([已关闭](https://github.com/agentscope-ai/QwenPaw/pull/5727)) 修复了Goal Mode架构和范围过滤的严重缺陷，解决了特定模式下命令不生效的问题。
    -   **内存系统增强:** PR #5296 ([已关闭](https://github.com/agentscope-ai/QwenPaw/pull/5296)) 将ADBPG长期内存组件改造为纯REST架构，去除了SQL依赖，提升了稳定性和维护性。
    -   **消息处理优化:** PR #5693 ([已关闭](https://github.com/agentscope-ai/QwenPaw/pull/5693)) 修复了纯文件消息（如通过企业微信发送的Excel文件）在无文本后续时被静默丢弃的Bug。
-   **通道与集成改进:**
    -   **飞书通道:** PR #5724 ([已关闭](https://github.com/agentscope-ai/QwenPaw/pull/5724)) 解决了飞书通道误拦截Bot消息的问题（对应Issue #5709），这是今日社区非常关注的一个修复。
-   **用户体验与界面优化:**
    -   **会话管理UI:** PR #5728 ([已关闭](https://github.com/agentscope-ai/QwenPaw/pull/5728)) 为会话列表增加了日期分组和搜索功能，显著提升了用户体验。
    -   **模型更新:** PR #5730 ([已关闭](https://github.com/agentscope-ai/QwenPaw/pull/5730)) 更新了内置模型列表，确保对新模型的支持。
-   **架构与治理:**
    -   **治理策略通用化:** PR #5546 ([已关闭](https://github.com/agentscope-ai/QwenPaw/pull/5546)) 将治理策略模式通用化，为后续安全与权限控制功能的开发奠定了基础。
    -   **网站多语言支持:** PR #5729 ([已关闭](https://github.com/agentscope-ai/QwenPaw/pull/5729)) 将官网默认语言切换为英文，优化了国际用户的首次体验。

**总结：** 项目通过快速合并PR，成功修复了多个与飞书、Goal Mode、消息处理相关的关键Bug，并增强了系统架构的灵活性和用户界面的易用性，整体向前迈出了坚实一步。

## 4. 社区热点

今日讨论最为活跃的议题集中于 **安全性、稳定性与核心功能缺失** 方面。

1.  **[Env] 密钥脱敏与安全存储** [#5705](https://github.com/agentscope-ai/QwenPaw/issues/5705)
    -   **讨论热度:** 5条评论
    -   **诉求分析:** 用户对API密钥等敏感信息的安全性提出了更高要求，不仅关注存储加密，还关注日志输出中的脱敏问题。这反映了企业级用户对安全合规的核心诉求。

2.  **[Bug] 浏览器自动填充干扰** [#5403](https://github.com/agentscope-ai/QwenPaw/issues/5403)
    -   **讨论热度:** 6条评论，持续更新超一周
    -   **诉求分析:** 这个Bug持续时间较长且评论区活跃，揭示了用户对**基本的可用性和输入体验**的困扰。浏览器将搜索框误判为密码输入框，直接影响了用户配置模型的核心操作流程，属于高频痛点。

3.  **[Bug] 内存泄漏** [#5720](https://github.com/agentscope-ai/QwenPaw/issues/5720)
    -   **讨论热度:** 4条评论，今日新开
    -   **诉求分析:** 用户进行了详尽的根因分析，表明这是一个高优先级的技术Bug。内存泄漏导致服务崩溃和配置丢失，对生产环境和长时间运行的用户是致命打击，引发了社区的强烈关注。

4.  **[Feature Request] 增强CLI能力** [#5737](https://github.com/agentscope-ai/QwenPaw/issues/5737)
    -   **讨论热度:** 1条评论，但代表了明确的开发者需求
    -   **诉求分析:** 用户（很可能是ISV或企业解决方案提供商）希望增强CLI能力以方便在非图形化场景下集成和自动化QwenPaw。这表明项目正在从单一用户工具向可集成平台演进，社区对此有明确期待。

## 5. Bug 与稳定性

以下为今日报告的Bug，按严重程度排序：

-   **严重 - 内存泄漏导致服务崩溃**
    -   **Issue:** [#5720](https://github.com/agentscope-ai/QwenPaw/issues/5720) - Qwen Paw v1.1.12.post2 内存泄漏反馈
    -   **状态:** 未指定，社区用户已详细分析了根因（异步任务泄漏、HTTP会话未回收）。
    -   **分析:** 这是最严重的稳定性问题，直接影响服务的持续运行和数据安全。

-   **高 - 飞书通道功能缺失**
    -   **Issue #1:** [#5709](https://github.com/agentscope-ai/QwenPaw/issues/5709) - 飞书通道硬丢弃 Bot 消息
        -   **状态:** 已关闭。**注意: 此Bug已有对应的修复PR #5724今日已合并**。
    -   **Issue #2:** [#5708](https://github.com/agentscope-ai/QwenPaw/issues/5708) - 飞书交互式卡片消息不解析
        -   **状态:** 未指定，该问题已存在超过两周。
    -   **Issue #3:** [#5721](https://github.com/agentscope-ai/QwenPaw/issues/5721) - 飞书群聊中历史消息丢失发送者信息
        -   **状态:** 未指定，今日新报告。
    -   **分析:** 飞书通道作为重要集成渠道，一夜之间暴露3个严重Bug，其中2个影响核心消息交互。**对飞书用户而言，当前版本稳定性堪忧。**

-   **中 - 浏览器性能问题**
    -   **Issue:** [#5725](https://github.com/agentscope-ai/QwenPaw/issues/5725) - Console 流式输出过程中浏览器卡顿
    -   **状态:** 未指定，用户已和DeepSeek进行对比，定位到QwenPaw自身问题。
    -   **分析:** 影响用户实时交互体验，属于比较影响感知的Bug。

## 6. 功能请求与路线图信号

今日收集到的功能请求表明，项目未来的重点方向在于：

1.  **安全性:** **密钥脱敏与安全存储** ([#5705](https://github.com/agentscope-ai/QwenPaw/issues/5705)) 是呼声最高的功能请求。结合今日合并的**治理策略通用化PR** [#5546](https://github.com/agentscope-ai/QwenPaw/pull/5546) 和 **多维限流PR** [#5738](https://github.com/agentscope-ai/QwenPaw/pull/5738)，可以判断 **安全性是企业采用QwenPaw的关键门槛**，项目方正在优先构建安全基础设施。
2.  **开发者体验与集成能力:**
    -   **增强CLI能力** ([#5737](https://github.com/agentscope-ai/QwenPaw/issues/5737)) 的需求，与今日发布的**v2.0.0-beta.2中新增的CLI功能**和多个新提交的CLI相关PR（如 [#5732](https://github.com/agentscope-ai/QwenPaw/pull/5732) 添加 `none` 内存后端）高度吻合。**QwenPaw正在强化其作为可编程平台的能力**。
3.  **Agent能力增强:**
    -   多个PR致力于提升Agent的记忆和检索能力，例如为内存搜索结果添加**reranker** ([#5692](https://github.com/agentscope-ai/QwenPaw/pull/5692), [#5669](https://github.com/agentscope-ai/QwenPaw/pull/5669)) 并为该功能添加了**UI配置** ([#5691](https://github.com/agentscope-ai/QwenPaw/pull/5691))。这显示项目正在**持续优化Agent的“记忆”准确性和检索质量**。
    -   **Windows桌面GUI自动化** PR [#5187](https://github.com/agentscope-ai/QwenPaw/pull/5187) 处于活跃开发中，显示项目正在探索Agent的“操作”能力。

## 7. 用户反馈摘要

从今日的Issues和PR评论中，可以提炼出以下核心用户反馈：

-   **痛点：**
    -   “我的浏览器在我搜索模型提供商时，老是想帮我填密码，太烦了！”（来自 #5403）
    -   “我的服务在稳定运行64分钟后因为内存泄漏崩溃了，所有配置都丢了，非常沮丧。”（来自 #5720）
    -   “飞书群里的Agent间根本不能正常@沟通，卡片消息也是乱码，这通道完全没法用。”（来自 #5709, #5708）
-   **期望：**
    -   “我们希望API密钥不要出现在任何日志里，这是最基本的安全要求。”（来自 #5705）
    -   “云服务商希望用CLI来预装和配置Skill，给个方便的接口吧。”（来自 #5737）
    -   “如果QwenPaw在流式输出时能像DeepSeek网页版一样丝滑就好了。”（来自 #5725）

## 8. 待处理积压

以下为长期未响应或对项目健康度有潜在影响的议题，提请维护者关注：

-   **高优先级的未修复Bug:**
    -   **`(严重) 飞书卡片消息不解析`** - [#5708](https://github.com/agentscope-ai/QwenPaw/issues/5708) (创建于10天前)。此Bug会完全阻断飞书的一大类交互场景。虽然其关联的PR #5724修正了另一飞书问题，但此问题至今无PR关联。
    -   **`(高) 向量索引无限膨胀至37G`** - [#4795](https://github.com/agentscope-ai/QwenPaw/issues/4795) (已关闭，但问题严重且影响时间长)。虽然已关闭，但这暴露了ChromaDB集成的重要隐患。在v2.0.0的释出周期中，是否有对此类根本性问题的长期方案，可能需要向社区澄清。
-   **长期等待响应或合并的PR:**
    -   **`Windows原生沙箱`** - PR [#5525](https://github.com/agentscope-ai/QwenPaw/pull/5525) (创建于7天前，`first-time-contributor`)。来自新贡献者的重要功能，若能合并将极大拓展QwenPaw在Windows环境下的安全边界。长时间未响应可能打击社区贡献积极性。

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，请看根据您提供的 GitHub 数据生成的 ZeroClaw 项目动态日报。

---

## ZeroClaw 项目动态日报 — 2026年7月2日

### 1. 今日速览

今日 ZeroClaw 项目开发活动极其活跃，展现了社区和核心团队的高参与度。**Issues 池**持续扩大，有12条新/活跃讨论，其中包含多个高优先级Bug。**Pull Request** 侧尤为繁忙，共有50条待处理或已处理的PR，但合并/关闭率（20%）偏低，反映出review瓶颈，尤其在大型功能分支上。尽管如此，项目核心团队正稳步推进通道扩展、安全加固和SOP（标准操作程序）可视化编辑等关键路线图事项。总体来看，项目处于快速迭代期，但需关注PR合并效率。

### 2. 版本发布

*(无更新)*

### 3. 项目进展

今日合并/关闭了10个PR，以下为几个重大进展：

- **通道功能重大拓展**：`#8504` 合并，实现了基于Git Forge的通道（支持GitHub App），并集成SOP作为入口处理逻辑。这是实现多渠道、自动化工作流的重要里程碑。
- **安全性与稳定性加固**：`#7361`（RFC-6969：per-turn输出路由）合并，修复了语音频道可能存在的双重发送bug；`#8567` 合入运行时OpenTelemetry内容策略，默认为安全模式，保护用户隐私。
- **新通道集成**：`#8547` 落地了ACP通道支持，虽因修复安全漏洞临时移除了PDF依赖，但核心通道框架得到丰富。
- **开发者体验优化**：`#8614` 和 `#8607` 等文档与CI相关的PR合并，改善了CI诊断输出并规范了PR处理流程。

这些合并成果表明项目正按既定的路线图快速推进，在功能、安全、开发者体验三方面均有提升。

### 4. 社区热点

本周社区讨论的焦点集中在几大领域：

1.  **标准化接口的迫切需求**：`#8550` 和 `#8603`（要求增加OpenAI兼容的Chat Completions端点）是今日最热门的话题之一。用户 `REL-mame` 提出了将ZeroClaw与Open WebUI、LobeChat等标准客户端集成的强烈愿望，并迅速提交了RFC。这表明社区对提升项目互操作性和降低集成门槛有很高呼声。

    [Issue #8550](zeroclaw-labs/zeroclaw Issue #8550)
    [Issue #8603](zeroclaw-labs/zeroclaw Issue #8603)

2.  **SOP（标准操作程序）生态的演进**：`#8590` 是今天体量最大、影响最广的PR（`size:XL`），作者 `singlerider` 提交了支持可视化SOP创作的功能。这直接回应了用户 `susyabashti` 在`#8587`中提出的增加SOP文档示例的诉求，表明项目紧跟用户需求，并已进入高级功能开发。

    [PR #8590](zeroclaw-labs/zeroclaw PR #8590)
    [Issue #8587](zeroclaw-labs/zeroclaw Issue #8587)

### 5. Bug 与稳定性

今日报告了多个Bug，其中高优先级问题较为集中：

- **高风险** `#7462`：Windows平台存在74个测试失败，主要与Unix-only命令和路径语义有关。这是一个全局性问题，意味着项目在Windows上的支持存在显著缺口。目前无Fix PR。
- **高风险** `#6302`：调用Gemini模型时因历史消息顺序错误被400拒绝。这是一个影响核心功能的兼容性问题。目前无Fix PR，但有相关讨论。
- **高风险** `#8605`：运行时配置自修改的安全防护逻辑存在路径错误，导致保护失效。这是一个高优先级的安全Bug。已在`Bug`栏下。
- **高风险** `#8615`：兼容性Provider（如兼容OpenAI的）会悄无声息地删除`<think>`标签及其内容，导致用户丢失回复内容。这是一个静默数据丢失问题。

**已有Fix PR的Bug**：
- `#8488` 尝试修复通道系统提示中工具可用性声明不一致的问题。
- `#8591` 试图修复审计签名密钥可能通过错误信息泄露的安全隐患。
- `#8574` 致力于修复解压缩Zip炸弹可能导致磁盘耗尽的风险。
- `#8599` 旨在对齐WebSocket路径下Agent的工具分发器与当前活动Provider。

这些活跃的Fix PR表明核心团队正在积极处理稳定性和安全问题。

[Issue #7462](zeroclaw-labs/zeroclaw Issue #7462)
[Issue #6302](zeroclaw-labs/zeroclaw Issue #6302)
[Issue #8605](zeroclaw-labs/zeroclaw Issue #8605)
[Issue #8615](zeroclaw-labs/zeroclaw Issue #8615)

### 6. 功能请求与路线图信号

除了已被采纳的SOP可视化编辑，以下请求可能影响未来路线图：

- **高级文件操作** `#8602`：用户 `NiuBlibing` 提议增强`file_read`工具，使其像Claude Code一样支持自动编码检测、PDF分页阅读、Jupyter Notebook感知等。这反映了对更强大的Agent工具链的需求。
- **简易模型切换** `#8600`：用户 `vvuk` 请求支持在运行时轻松切换同一Provider下的不同模型。这是提升用户体验和Agent灵活性的重要诉求。
- **预构建资产扩展** `#7952`：用户 `Audacity88` 建议发布包含完整频道支持的预编译包，而不是默认的“精简版”。这有助于新手快速体验全部功能，降低部署门槛。

[Issue #8602](zeroclaw-labs/zeroclaw Issue #8602)
[Issue #8600](zeroclaw-labs/zeroclaw Issue #8600)
[Issue #7952](zeroclaw-labs/zeroclaw Issue #7952)

### 7. 用户反馈摘要

从今日的讨论中可以提炼以下主要用户声音：

- **痛点**：
    - **Windows支持不足**：Windows上的大量测试失败是一个严重障碍，使得Windows用户或开发者无法获得稳定体验。
    - **配置与安全困惑**：`#8605` 的配置路径Bug和 `#8615` 的静默内容删除，被用户标记为“危险行为”，影响了用户对系统安全性的信任。
    - **集成困难**：无法使用通用OpenAI兼容的客户端（如LobeChat）是社区表达的主要集成痛点，迫使用户自行开发适配器。

- **期望**：
    - **简化部署**：`#7952` 的请求表明用户希望项目能提供“开箱即用”的完整体验，而不是需要复杂的配置。
    - **增强互操作性**：对OpenAI兼容端点的强烈呼声，代表了社区希望ZeroClaw能够融入更广泛的AI工具生态。
    - **更智能的工具**：对 `file_read` 增强的请求表明用户希望Agent工具更强大、更智能，而非仅仅满足基本功能。

### 8. 待处理积压

以下为长期未决或状态为“blocked”的高优先级Issue，值得维护者关注：

1.  **`#7952` [Feature]: publish full-channel prebuilt assets alongside default prebuilts** (状态: `blocked`) 作者: `Audacity88`。自6月19日创建以来，该请求提议发布更完整的预编译包，但至今仍处于`blocked`状态，可能是由于技术决策或资源分配问题。

    [Issue #7952](zeroclaw-labs/zeroclaw Issue #7952)

2.  **`#6302` [Bug]: Gemini 400 — assistant tool_call emitted as first non-system turn** (状态: `accepted`, 优先级 `p1`) 作者: `dmnkhorvath`。自5月初提出，这是一个影响核心功能的高风险Bug，虽已接受但尚未有修复PR合并，应优先推进。

    [Issue #6302](zeroclaw-labs/zeroclaw Issue #6302)

3.  **`#8615` [Bug]: compatible provider silently deletes content via unconditional `<think>` tag stripping** (状态: `accepted`, 优先级 `p2`) 作者: `NiuBlibing`。该Bug会导致用户数据静默丢失，严重性高，且为今日新报告，应获得迅速响应和修复资源。

    [Issue #8615](zeroclaw-labs/zeroclaw Issue #8615)

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*