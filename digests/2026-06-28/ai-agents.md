# OpenClaw 生态日报 2026-06-28

> Issues: 150 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-06-28 10:09 UTC

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

# OpenClaw 项目动态日报 — 2026-06-28

## 今日速览

过去24小时项目保持高活跃度：共处理 **150 条 Issue**（其中新开/活跃 140 条，关闭 10 条）和 **500 条 PR**（其中 135 条已合并/关闭，待合并 365 条）。虽然无新版本发布，但修复和功能推进密集，尤其是在会话/消息可靠性、多通道支持、安全边界加固等方面。社区讨论热度集中在**重复消息、会话状态丢失**等稳定性问题上，同时多个高优先级 Bug 的修复 PR 已进入评审阶段，项目整体健康度良好，但积压 Issue 数量仍然较大。

## 版本发布

无新版本发布。当前最新 release 仍为 `2026.4.x` 系列。

## 项目进展

今日关闭/合并的重要 PR 包括：

- **#97454** 修复自定义 OpenAI 兼容网关在 `Content-Type: application/json` 下错误包装 SSE 的问题 → [PR #97454](https://github.com/openclaw/openclaw/pull/97454)
- **#97460** 修复 CLI 在 agent reply 时对 `exec` 凭证进行不必要的 `secrets.resolve` 导致错误日志 → [PR #97460](https://github.com/openclaw/openclaw/pull/97460)
- **#97461** 修复 Matrix 投票组件在去重前应用 `max_selections` 导致重复答案占位的问题 → [PR #97461](https://github.com/openclaw/openclaw/pull/97461)
- **#68936** 新增 PR 审查自动修复流水线及 Windows 守护进程（合并） → [PR #68936](https://github.com/openclaw/openclaw/pull/68936)
- **#73311** 为 Control UI 添加会话类型过滤功能 → [PR #73311](https://github.com/openclaw/openclaw/pull/73311)

此外，多个长期 PR 如 **#88968**（防止 memory flush 失败阻断用户回复）、**#74231**（媒体路径拒接时给出可配置根目录提示）仍在积极 review 中，今日有更新。

## 社区热点

最受关注的 Issues 集中在**消息丢失、会话状态损坏**等核心可靠性问题上：

1. **#69208** (评论12, P1) — “跨通道重复转发、回放和上下文组装”的 umbrella issue，涵盖 MSTeams、webchat、Telegram 等多个通道 → [#69208](https://github.com/openclaw/openclaw/issues/69208)
2. **#67777** (评论10, P1) — Subagent 完成投递在直接 announcement 超时或 drain 时丢失 → [#67777](https://github.com/openclaw/openclaw/issues/67777)
3. **#71736** (评论9, P2) — 提议为 Control UI 插件贡献槽标准化，开发者呼吁明确 SDK 接口 → [#71736](https://github.com/openclaw/openclaw/issues/71736)
4. **#76171** (评论7, 3👍, P1) — 高主机负载与慢响应，由陈旧工作进程积累引起，已关闭但被社区多次点赞 → [#76171](https://github.com/openclaw/openclaw/issues/76171)

讨论反映出用户对**稳定运行**（尤其多通道场景）和**插件生态可扩展性**的强烈需求。

## Bug 与稳定性

按严重程度排列今日活跃的重要 Bug（P1/P2，含崩溃、数据丢失、安全）：

| 级别 | Issue | 摘要 | 是否有 fix PR |
|------|-------|------|----------------|
| P1 | #69208 | 跨通道重复转发、回放和上下文组装 | 无明确 PR，umbrella issue |
| P1 | #67777 | Subagent 完成投递丢失 | 无 |
| P1 | #72015 | active-memory 插件导致多 agent 网关回复缓慢甚至崩溃 | 无 |
| P1 | #76171 | 陈旧 worker 进程积累导致高负载（已关闭） | 已合并修复（#？未明确） |
| P1 | #72031 | Bedrock 图片工具因 `requireApiKey` 错误抛异常 | 有 linked PR（#？） |
| P1 | #73182 | Claude 推理默认被静默开启导致费用翻倍、thinking 泄露 | 无 |
| P1 | #69943 | session-memory hook 持久化未清理的模板 token，造成自循环污染 | 无 |
| P1 | #71699 | Windows 下 Mattermost 流式回复导致堆栈缓冲溢出崩溃（STATUS_STACK_BUFFER_OVERRUN） | 无 |
| P2 | #73602 | WhatsApp 反复断连 + Telegram 在 WSL2 下轮询停滞 | 无 |
| P2 | #72704 | Telegram 消息中嵌入过多 JSON 元数据，影响模型理解 | linked PR open |
| P2 | #72418 | loopback 客户端可通过 `shouldSkipBackendSelfPairing` 绕过设备配对（CVSS 8.7/9.3） | linked PR open |
| P2 | #73910 | Codex ACP 使用隔离环境导致认证桥接失败，超时配置不兼容 | 无 |

## 功能请求与路线图信号

用户提出的新功能需求及可能纳入下一版本的信号：

- **#71142** (P2) — 请求为 Control UI 设置可配置的文件上传大小限制（目前硬编码 5MB） → [#71142](https://github.com/openclaw/openclaw/issues/71142)
- **#74077** (P3) — 添加 `/stream` 斜杠命令以在不重启网关下调整预览流式模式 → [#74077](https://github.com/openclaw/openclaw/issues/74077)
- **#73537** (P2) — 请求为发布版本添加“生产就绪”稳定性标签，以便用户选择稳定版 → [#73537](https://github.com/openclaw/openclaw/issues/73537)
- **#71058** (P2) — 支持单个 OpenClaw 网关配置多个 Teams Bot → 已有对应 PR **#97340**（今日新开，涉及多账户支持）→ [PR #97340](https://github.com/openclaw/openclaw/pull/97340)
- **#71195** (P2) — 为 macOS Talk Mode 添加 OpenAI Realtime 语音到语音路径 → [#71195](https://github.com/openclaw/openclaw/issues/71195)
- **#71301** (P3) — 随发布版本捆绑文档，并给予 agent 原生文档检索能力 → 路线图信号

多个 feature request 已有关联 PR 或处于待决策状态，表明社区对**插件 SDK、多通道管理、稳定性标签**等需求迫切。

## 用户反馈摘要

从 Issues 摘要中提炼的真实用户痛点与场景：

- **#76171**（高负载）：用户报告升级到 `2026.4.29` 后响应时间恶化至 2-3 分钟，因陈旧 worker 进程堆积。该问题已关闭，但用户期望更彻底的防复发机制。
- **#73602**（WSL2 通道不可靠）：用户在 Windows WSL2 环境下 WhatsApp 和 Telegram 通道反复失效，DNS 错误频发，甚至假健康状态。
- **#70903**（持久化冷却：用户 Anthropic 账单积分恢复后仍被封锁数小时，因为 `disabledUntil` 参数持续延长。
- **#73910**（Codex ACP 认证失败）：用户能直接通过 ACPX 连接 Codex，但 OpenClaw 管理的 Codex ACP 会话始终失败——问题出在隔离认证桥接。
- **#72176**（重复消息）：用户升级到 `2026.4.24` 后，相同的用户消息偶发交付两次，导致助手回复重复。
- **#73049**（embeddings API key 错误）：用户升级到 `2026.4.25` 后，`memory-core` 插件使用了错误的 API key 进行嵌入，完全打破 agent 功能。

整体用户满意度较高，但**稳定性回归**（尤其是跨版本升级后）是主要不满来源。

## 待处理积压

以下长期未响应或停滞的重要 Issue/PR 需维护者关注：

| 类型 | 编号 | 摘要 | 标记 | 上次更新 |
|------|------|------|------|----------|
| Issue | #71417 | `openclaw agent` 默认使用 `--channel last`，静默恢复最近会话 | stale, P2 | 2026-06-28 |
| Issue | #71712 | 请求 agent 调度 API 并提供不可伪造的 provenance | stale, P2 | 2026-06-28 |
| Issue | #73182 | Claude 推理默认被翻转导致费用翻倍 | stale, P1, needs-security-review | 2026-06-28 |
| Issue | #70903 | 持久化提供商冷却在账单恢复后继续封锁 | stale, P2 | 2026-06-28 |
| Issue | #69572 | Feishu 打字指示器错误使用 Message Reaction API | stale, P2 | 2026-06-28 |
| PR | #64546 | Mattermost 互动令牌 HMAC 硬编码问题（安全风险） | waiting on author, P1 | 2026-06-28 |
| PR | #72713 | 允许信任的 `.openclaw` 符号链接（macOS） | needs-real-behavior-proof | 2026-06-28 |
| PR | #69822 | 会话消息事件 socket.drain 实现 | waiting on author | 2026-06-28 |

这些积压项多涉及安全、会话状态、配置行为，建议维护者优先安排 review 或给出决策意见。

---

*报告生成日期：2026-06-28，数据来源：[OpenClaw GitHub 仓库](https://github.com/openclaw/openclaw)*

---

## 横向生态对比

好的，作为资深技术分析师，以下是根据您提供的各项目动态日报生成的横向对比分析报告。

---

## 个人 AI 智能体开源生态横向对比分析报告 (2026-06-28)

### 1. 生态全景

当前个人 AI 智能体（Personal AI Agent）与自主智能体（Autonomous Agent）开源生态正处于 **功能密集迭代与稳定性阵痛并存** 的快速发展期。一方面，以 OpenClaw 为代表的核心参照项目在会话可靠性、多通道支持、安全边界加固等领域持续深挖，而 NanoBot、IronClaw 等新锐项目则在多智能体协作（A2A）、能力策略（Capability Policy）等前瞻性方向上快速布局。另一方面，社区反馈普遍集中于 **稳定性回归**（跨版本升级后 Bug 频发）、**平台兼容性**（Windows/Android/MacOS 适配不足）以及 **插件/SDK 标准化** 等实际问题。整体而言，生态正从“能用”向“好用、安全、可扩展”过渡，开发者需在跟进新功能与规避稳定性风险之间取得平衡。

### 2. 各项目活跃度对比 (2026-06-28)

| 项目 | Issues (新开/活跃) | PR (新) | PR 合并/关闭 | 版本发布 | 项目健康度 |
|------|------------------|---------|--------------|---------|-----------|
| **OpenClaw** | 140 | 500 | 135 | 无 | **优秀**（高活跃，但积压多） |
| **NanoBot** | 4 (新) | 24 | 5 | 无 | **优秀**（高质量 PR 密集） |
| **Hermes Agent** | 10 | 50 | 15+ | 无 | **良好**（安全加固与稳定性修复并行） |
| **PicoClaw** | 3 | 3 | 2 | 无 | **中等**（关键 Bug 待修，合并效率低） |
| **NanoClaw** | 1 (新) | 7 | 0 | 无 | **中等**（合并瓶颈严重，核心功能有 Bug） |
| **NullClaw** | 0 (新) | 1 | 0 | 无 | **低**（低频，Android 构建 Bug 长期未解） |
| **IronClaw** | 10+ | 50 | 24 | 无 | **极高**（Capability Policy 里程碑合并） |
| **LobsterAI** | 2 (新开, 关闭) | 4 (关闭) | 4 | 无 | **低**（主要靠自动清理陈旧项，缺乏新进展） |
| **Moltis** | 1 (新) | 4 | 0 | 无 | **中等**（修复方向明确，但长期 PR 搁置） |
| **CoPaw** | 2 (新) | 6 | 0 | 无 | **中等**（测试基础设施推进，关键回归 Bug 待修） |
| **ZeptoClaw** | - | - | - | 无 | **无活动** |
| **TinyClaw** | - | - | - | 无 | **无活动** |
| **ZeroClaw** | 10+ | 50 | 10 | 无 | **极高**（功能需求与 RFC 活跃，合并待提升） |

**分析**：
- **IronClaw** 和 **ZeroClaw** 在功能推进上最为激进，PR 提交和合并/关闭量居前。
- **OpenClaw** 虽活跃度最高，但待合并 PR 积压 365 条，显示维护瓶颈。
- **NullClaw、ZeptoClaw、TinyClaw** 处于停滞或极低频状态。

### 3. OpenClaw 在生态中的定位

- **优势**：
  - **规模最大**：Issues/PR 数量远超其他项目，社区反馈最为全面，是事实上的“生态参照系”。
  - **深度覆盖多通道**：跨通道重复转发、Matrix/Telegram/WhatsApp 等通道的修复持续进行，通道兼容性是 OpenClaw 的核心壁垒。
  - **安全与稳定性投入**：对内存溢出（STATUS_STACK_BUFFER_OVERRUN）、凭证泄露、会话状态丢失等 P1 级 Bug 有系统性跟踪。
- **技术路线差异**：
  - **OpenClaw** 采用 **微内核 + 插件** 架构，强调可插拔性与多模型/多通道适配，近期重点在会话可靠性（如 active-memory 插件优化）。
  - **IronClaw** 倾向于 **高度集成的能力策略模型**，通过四维权限模型控制 agent 行为，更聚焦企业级权限管理。
  - **NanoBot** 则强调 **超轻量部署** 与 **Agent-to-Agent 协作**，技术路线更为激进（如原生 A2A、子代理模型覆盖）。
- **社区规模对比**：
  - OpenClaw 的 Issue 评论量（如 #69208 达 12 条）和 PR 数量（单日 500 条）远超其他项目，社区贡献者最活跃。
  - NanoBot、Hermes Agent 虽数据量较小，但贡献者专注度高，高质量修复占比大。

### 4. 共同关注的技术方向

| 技术方向 | 涉及项目 | 具体诉求 |
|----------|----------|----------|
| **会话/消息可靠性** | OpenClaw, NanoBot, PicoClaw, NanoClaw, Moltis | 重复消息、上下文丢失、流式中断、SSE 封装错误等 |
| **多通道稳定与扩展** | OpenClaw, NanoBot, PicoClaw, NanoClaw, ZeroClaw, CoPaw | Telegram/WhatsApp/Matrix/Signal 断连、流式支持、富文本渲染 |
| **小/本地模型兼容性** | Moltis, ZeroClaw, NanoBot | 小模型输出格式不标准（字符串参数、null 处理）、上下文预算溢出 |
| **安全性加固** | OpenClaw, Hermes Agent, ZeroClaw, IronClaw | API 密钥脱敏、权限提前检查、OAuth 令牌刷新、MCP 源锁定、loopback 绕过 |
| **插件/技能生态标准化** | OpenClaw, LobsterAI, NanoClaw, ZeroClaw | 插件 SDK 接口可预期、技能热更新、重复校验、CDN 安装修复 |
| **Agent 协作与任务编排** | NanoBot (A2A), PicoClaw (协作总线), ZeroClaw (SOP), IronClaw (Capability Policy) | 子代理模型覆盖、逻辑邮箱、Kanban 调度、定时任务 SOP |

**分析**：**会话可靠性**与**多通道稳定**是跨项目的共同难题，反映了个人 AI 智能体在生产环境中面临的首要挑战。**小模型兼容**和**安全加固**则体现了生态对低成本与合规性的双重追求。

### 5. 差异化定位分析

| 项目 | 功能侧重 | 目标用户 | 技术架构 | 关键差异化 |
|------|----------|----------|----------|-----------|
| **OpenClaw** | 多通道/多模型通用平台 | 个人开发者、中小团队 | 微内核 + 插件 | 通道兼容性最强，社区最庞大 |
| **NanoBot** | 轻量级个人助手 | 技术爱好者、独立开发者 | 超轻量 Python/Node | 强调低资源消耗，早期集成 A2A |
| **Hermes Agent** | 企业级安全与自动化 | 企业运维/安全团队 | Gateway + 调度器 | 安全加固（API 密钥脱敏、权限检查）与 Kanban 调度 |
| **PicoClaw** | 嵌入式/低功耗设备 | IoT 开发者、边缘计算用户 | Go/CLI | 极简部署，代理协作总线创新 |
| **NanoClaw** | 容器化多提供商支持 | DevOps、云原生用户 | Docker/K8s | 围绕 OpenAI/Anthropic 等提供商容器化，注重可部署性 |
| **NullClaw** | 实验性极简项目 | 高级开发者/研究者 | Zig | 极度轻量，构建门槛高（Android 兼容性差） |
| **IronClaw** | 企业级权限与控制 | 企业客户、平台开发者 | Rust 后端 + WebUI | Capability Policy（四维权限模型）、集成测试框架 |
| **LobsterAI** | 技能管理与定时任务 | 中文用户、教学场景 | 前端重 | 技能导入校验、定时任务 UI 升级，中文社区 |
| **Moltis** | 本地模型优化 | 本地模型爱好者、隐私敏感用户 | Rust | 小模型兼容（预处理、类型强制转换）、图像降采样 |
| **CoPaw** | Chinese LLM 集成 | 国内开发者、华为昇腾用户 | Python | 依赖 AgentScope，测试覆盖推进，适配 Ascend-vLLM |
| **ZeroClaw** | 高扩展性多通道平台 | 重度多通道用户、自动化场景 | 多通道/协议 | WhatsApp 被动上下文、Wire-Protocol 重构 RFC、SOP 文件系统触发 |

### 6. 社区热度与成熟度

- **快速迭代阶段**（高活跃，功能与 Bug 并行）：
  - **IronClaw** (🔥 核心功能冲刺，测试框架完善)
  - **ZeroClaw** (🔥 大量新功能 PR 与 RFC，社区讨论活跃)
  - **OpenClaw** (🔥 规模最大，但维护压力大，积压多)

- **质量巩固阶段**（聚焦修复与稳定，新功能较少）：
  - **NanoBot** ✅ 高质量修复 PR 密集，但项目本身体量较小
  - **Hermes Agent** ✅ 安全与稳定性投入显著，Kanban 调度优化
  - **PicoClaw** ✅ 关键协作总线落地，但 Matrix 加密 Bug 待解决
  - **Moltis** ✅ 针对小模型兼容做专项修复，但长期 PR 搁置

- **低频/停滞阶段**（活动稀少）：
  - **NullClaw** 🔴 仅 1 个 PR，Android 构建 Bug 2 个月未响应
  - **LobsterAI** 🔴 主要靠自动清理，新进展极少
  - **ZeptoClaw / TinyClaw** 🔴 无活动

### 7. 值得关注的趋势信号

1. **隐私与本地化模型需求显著增长**：Moltis、ZeroClaw 的多项 PR 和 Issue 直接回应小模型兼容、本地部署优化；NanoBot 社区质疑“超轻量”声明时也隐含对轻量化的较高标准。**对开发者启示**：应优先支持开源/本地模型，提供预处理降级策略，减少对云端大模型的硬依赖。

2. **Agent 协作（A2A）从概念走向实现**：NanoBot 的 A2A PR、PicoClaw 的代理协作总线、ZeroClaw 的 SOP 事件源，都表明多智能体间的结构化通信（邮箱、权限、协议）正成为标配。**对开发者启示**：尽快设计统一的 Agent 间消息协议，避免各自为政。

3. **安全合规成为必备而非附加**：Hermes Agent 的 API 密钥脱敏、IronClaw 的四维权限模型、ZeroClaw 的 Wire-Protocol 重构，反映出用户对个人智能体应用场景（如自动化、远程控制）的信任门槛在提高。**对开发者启示**：从架构初期引入细粒度权限模型和密钥脱敏机制。

4. **平台兼容性仍是最大短板**：Windows（反斜杠路径问题）、WSL2（DNS 错误）、Android（NullClaw 构建失败）、macOS（Talk Mode 路径缺失）等平台问题在各项目中反复出现，严重阻碍了用户群体的扩大。**对开发者启示**：建立跨平台 CI 矩阵，至少覆盖 Linux/macOS/Windows 三大平台，并做好容器化适配。

5. **社区对“稳定性回归”的容忍度下降**：OpenClaw 高赞 Issue #76171（worker 进程积累导致高负载）、NanoClaw 频繁报告跨版本升级后崩溃，表明用户期望项目在发布新功能的同时保持向后兼容。**对开发者启示**：建立完整的回归测试套件和兼容性测试，避免在快速迭代中忽视基础稳定性。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，根据您提供的 NanoBot 项目数据，我为您生成了 2026-06-28 的项目动态日报。

---

## NanoBot 项目动态日报 | 2026-06-28

### 1. 今日速览

NanoBot 项目今日呈现 **极高活跃度**，主要由密集的 Pull Request 提交驱动。过去 24 小时内共有 24 条 PR 更新，其中 19 条处于待合并状态，表明项目团队正在集中精力进行 Bug 修复和功能开发。社区互动集中在几个关键问题上，特别是关于 WebUI 稳定性、子代理模型自定义和上下文缓存优化。尽管没有新版本发布，但大量高质量 PR 的涌现预示着项目即将迎来一次重要的功能迭代和稳定性提升。项目整体健康度 **优秀**。

### 2. 版本发布

今日无新版本发布。

### 3. 项目进展

今日有 5 条 PR 被合并或关闭，推动了以下关键进展：

- **Agent 鲁棒性提升 (PR #4510, CLOSED)**：该 PR 修复了大型语言模型返回格式错误的工具调用（如名称缺失或为空）时导致 Agent 崩溃的问题。同时，它还修复了历史记录中可能存在的损坏数据，提升了 Agent 在处理异常响应时的自我修复能力。([链接](https://github.com/HKUDS/nanobot/pull/4510))

- **定时任务功能增强 (PR #4225 & #4357, CLOSED)**：这两条 PR 为定时任务引入了 `silent`（静默）模式和 `lock_recipient`（锁定接收者）功能。静默模式允许后台监控任务仅在需要报告时才输出结果，避免了不必要的通知干扰。这体现了项目在功能精细化上的持续投入。([链接1](https://github.com/HKUDS/nanobot/pull/4225), [链接2](https://github.com/HKUDS/nanobot/pull/4357))

- **仓库指南与协作规范 (PR #4575, CLOSED)**：该 PR 为仓库添加了统一的贡献指南，有助于降低新贡献者参与的门槛，优化社区协作流程。([链接](https://github.com/HKUDS/nanobot/pull/4575))

**整体来看**，项目在 Agent 稳定性、定时任务功能和社区协作规范方面都取得了实质性进展。

### 4. 社区热点

- **#660 “超轻量”声明引发的争议 (CLOSED)**： 尽管已关闭，但此 Issue 获得了 14 条评论和 5 个👍，是讨论最激烈的议题。用户 `besoeasy` 质疑项目自称“超轻量”（ultra-lightweight）但 Dockerfile 中同时依赖 Python 和 Node.js 的矛盾。这反映了社区对项目定位和宣传准确性的高度关注。([链接](https://github.com/HKUDS/nanobot/issues/660))

- **#4500 WebUI 自重启导致界面卡死 (OPEN, 高热度)**： 用户 `zpljd258` 报告了一个关键 Bug，即 WebUI 在 Nanobot 自重启后，界面会陷入持续的“处理中”状态且无法停止，严重影响使用体验。此问题已得到开发者关注，并有对应的修复 PR (#4565)。([链接](https://github.com/HKUDS/nanobot/issues/4500))

**分析**：社区热点主要集中在 **项目真实性与用户体验** 两方面。对宣传声明的质疑显示用户对项目有较高期待和审视标准；而 WebUI 的稳定性问题则是直接影响用户粘性的核心痛点。

### 5. Bug 与稳定性

今日共有 4 条新开的 Issues，均涉及 Bug 或稳定性问题，严重程度按顺序排列如下：

- **严重**：
    - **#4500 WebUI 自重启后界面卡死**： Bug 复现路径清晰，影响用户的核心交互。**已有对应修复 PR #4565**。([链接](https://github.com/HKUDS/nanobot/issues/4500))
    - **#4222 上下文管理导致提示缓存失效**： 该 Bug 指出 `max_messages` 截断和 `microcompact` 机制持续导致 LLM 的提示/前缀缓存失效，降低了响应速度并增加了推理成本。这是一个较为底层的性能问题。**已有对应修复 PR #4568**。([链接](https://github.com/HKUDS/nanobot/issues/4222))

- **中等**：
    - **#4569 工具调用路径对格式错乱的中继响应缺乏防护**： 当上游兼容 Anthropic 的中继服务返回格式错误的 `tool_use` 数据时，Agent 可能会崩溃。已提出 **PR #4569** 进行修复。([链接](https://github.com/HKUDS/nanobot/pull/4569))
    - **#4567 微信通道不支持流式传输**： 当前微信通道仅支持非流式 API，导致在特定中继环境下工具调用信息丢失。已提出 **PR #4567** 进行修复。([链接](https://github.com/HKUDS/nanobot/pull/4567))

### 6. 功能请求与路线图信号

- **#4231 子代理模型覆盖 (OPEN)**： 用户 `jsapede` 请求为 `spawn`（生成子代理）工具增加 `model` 参数，允许为不同子任务指定不同的模型。这是一个热门需求，**对应的 PR #4570 今日已经提交**，这表明该功能极有可能出现在下一个版本中。([链接](https://github.com/HKUDS/nanobot/issues/4231))

- **#4179 原生 A2A (Agent-to-Agent) 协作 (间接信号)**： 虽然此 Issue 不在今日数据内，但 **PR #4571** 正是为了解决其核心诉求：允许多个 Agent（如 Supervisor, Researcher）以原生方式进行协作，而非仅通过匿名的 `spawn` 创建子代理。这标志着项目正从单一的“主-从”模式向更灵活的“平级协作”架构演进。([链接](https://github.com/HKUDS/nanobot/pull/4571))

- **其他增强请求**：
    - **#4504** 允许将技能文件放置在子目录中以便于组织。([链接](https://github.com/HKUDS/nanobot/pull/4504))
    - **#4459** 增加对 Mattermost 消息通道的支持。([链接](https://github.com/HKUDS/nanobot/pull/4459))
    - **#4542** 改进 MCP 工具对图片内容的处理，将其作为附件 (artifact) 传递。([链接](https://github.com/HKUDS/nanobot/pull/4542))

**信号分析**：项目路线图正朝着 **更精细化的 Agent 控制**（模型选择、A2A协作）和 **更广泛的渠道集成**（Mattermost）方向迈进。

### 7. 用户反馈摘要

- **用户痛点**：
    - **体验不一致**：用户 `zpljd258` 指出 WebUI 自重启后界面卡死，以及停止按钮失效，这暴露了 WebSocket 连接的健壮性问题。([链接](https://github.com/HKUDS/nanobot/issues/4500))
    - **宣传与实际不符**：用户 `besoeasy` 对“超轻量”的声明提出质疑，反映出用户在选择项目时对技术栈精简度的期待。([链接](https://github.com/HKUDS/nanobot/issues/660))

- **使用场景与需求**：
    - **成本与性能优化**：用户 `imkuang` 报告的缓存失效问题（#4222），表明有用户在大规模或高频对话场景下使用 NanoBot，他们对推理延迟和成本非常敏感。([链接](https://github.com/HKUDS/nanobot/issues/4222))
    - **灵活性与可定制性**：用户 `jsapede` 要求子代理模型可配置（#4231），说明开发者在实际搭建复杂 Agent 工作流时，希望为不同任务分配最合适的模型，以实现性能与成本的平衡。([链接](https://github.com/HKUDS/nanobot/issues/4231))

### 8. 待处理积压

今日数据中未发现长期未响应的重要 Issue 或 PR。不过，以下两个 Issue 虽然已有对应的修复 PR，但其本身作为优先级较高的 Bug，建议维护者在合并 PR 后及时关闭，并督促测试以验证修复效果：

- **#4500** (WebUI 卡死) - 已有关联修复 PR #4565
- **#4222** (缓存失效) - 已有关联修复 PR #4568

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

好的，作为 Hermes Agent 项目的 AI 分析师，根据您提供的 GitHub 数据，我为您生成一份 2026 年 6 月 28 日的项目动态日报。

---

## Hermes Agent 项目日报 | 2026-06-28

### 1. 今日速览

今日项目活跃度极高。**过去24小时内，共有10个Issue和50个PR被更新**。项目团队在修补关键Bug（特别是关于环境安装和平台锁死问题）的同时，社区也贡献了大量功能请求和安全加固PR，显示出健康且充满活力的社区生态。尽管没有新版本发布，但合并了15个以上PR，整体项目正在快速迭代和稳定化。

### 2. 版本发布

今日无新版本发布。

### 3. 项目进展

今日项目进展主要集中在 **Bug 修复、安全加固和核心功能增强** 上。多个高优先级（P1/P2）的 PR 被提出或正在进行审查，旨在解决关键的稳定性和安全性问题。

*   **稳定性修复**
    *   **[PR #44772]** (修复中) **修复了 `hermes update` 删除根目录 Node 依赖的问题**，通过使用 `--workspaces=false` 并最后安装根目录依赖，确保了诸如 `agent-browser` 等核心依赖不受影响。
    *   **[PR #54162]** (待合并) **修复了 `--replace` 参数导致 systemd 闪退的问题**。通过正确处理僵尸进程（zombie PID），防止了 `--replace` 操作因等待死进程而失败并导致 systemd 无限重启。
    *   **[PR #54159]** (待合并) **修复了 macOS 上 `hermes update` 后网关脱离 launchd 监管** 的问题，确保更新后网关仍能保持开机自启和自动重启能力。
*   **安全加固**
    *   **[PR #54134]** (待合并) **安全加固**：建议将 `n8n MCP` 包锁定到不可变commit，并阻止目录清单（catalog manifests）中的浮动git引用。
    *   **[PR #54166]** (待合并) **安全修复**：防止 API 密钥在代理读取配置/数据文件时暴露或损坏。文件内容中的密钥将被替换为**非可重用标记**（如 `«redacted:ghp_…»`），而不是容易泄露的截断密钥。
    *   **[PR #54164]** (待合并) **安全修复**：在 Telegram 网关中，对已移除/未授权的用户的权限检查提前到事件构建之前，防止这类用户注入 Prompt 内容。
*   **核心功能增强**
    *   **[PR #54139]** (待合并) **修复了 MCP 重连机制**，确保断开后持续尝试重连并恢复注册的工具。
    *   **[PR #54145]** (待合并) **引入内存感知调度**，并将 `_recent_worker_exits` 持久化到磁盘，提升 Kanban 调度的可靠性和网关重启后的稳定性。
    *   **[PR #54146]** (待合并) 新增 `appnz` **LLM 网关提供商支持**，丰富了模型提供商生态。
    *   **[PR #54151]** (待合并) 减少对预期可选工具不可用时的警告日志，降低日志噪音。

### 4. 社区热点

今日社区讨论热度最高的议题集中在**平台兼容性与原生客户端需求**上。

1.  **[Issue #35966]**: **原生桌面/移动客户端应用 (Feature Request)**
    *   该功能请求要求开发一个不依赖第三方平台的本地客户端，直接与 Hermes 的 Gateway API 交互。评论数 (2) 和点赞数 (2) 都很高，说明社区对摆脱第三方平台（如 Telegram、Discord）的依赖有强烈诉求。这是一个长期需求，今日被再次提及并更新。

2.  **[Issue #50745]**: **移动端 App (Duplicate of #35966)**
    *   该问题同样提出了开发一款 iOS/Android App 的需求，以实现远程连接回桌面端会话。虽然被标记为 `duplicate`，但它的存在本身进一步印证了社区对 #35966 中所述的原生客户端的强烈呼声。

3.  **[Issue #39714]**: **Bug: `hermes update` 使用了错误的venv目录 (已关闭)**
    *   尽管已经关闭，但作为今日唯一关闭的高优Issue，其讨论热度（3条评论）表明该Bug对基于 `uv` 安装的用户造成了困扰。社区贡献者 `LineckerN` 报告了因硬编码 `venv/` 和实际使用的 `.venv/` 目录不一致导致的依赖安装失败问题，该BUG的快速关闭显示了项目对用户痛点的响应速度。

### 5. Bug 与稳定性

今日报告的 Bug 主要集中在对核心组件稳定性的威胁，且大多数已有修复 PR。

*   **[已修复] 严重 Issue**
    *   **[Issue #39714]** **[P1]** **修复 `hermes update` 使用错误的 venv 目录**。该问题影响所有 `uv` 安装的用户，会导致更新后核心包安装失败。**已通过相关 PR 合并关闭。**

*   **[待修复] 严重 Issue**
    *   **[Issue #54167]** **[P1]** **Telegram 机器人令牌锁文件永久死锁**。当 `hermes gateway` 在初始化 Telegram 时被异常杀死，锁文件残留会导致下次启动时整个 Telegram 平台功能瘫痪。**已有修复 PR #54176** 将该锁定失败标记为可重试，而非永久性致命错误。
    *   **[Issue #54174]** **[Bug]** **Docker 容器中 `hermes profile install` 后网关启动失败**。安装成功但未注册 s6 服务配置文件，导致 `hermes gateway start` 报 “no such gateway” 错误。这直接影响了 Docker 用户的正常使用。

*   **[中等] 稳定性问题**
    *   **[PR #54165]** **[P2]** **QQ 适配器重连循环**。`QQAdapter.connect()` 缺少 `is_reconnect` 参数，会导致无限重试，最终淹没日志。
    *   **[PR #54157]** **[P2]** **防止在 Gateway 会话中再次启动 Gateway**。修复因递归调用可能导致的状态混乱。
    *   **[PR #54139]** **[P2]** **MCP 重连可能进入无限等待**。确保重连尝试不会无限期暂停。

### 6. 功能请求与路线图信号

今日出现多个来自社区（特别是用户 `c03rad0r`）的新功能请求，这些建议具体且实用，很可能被纳入下一版本规划。

*   **高价值与具体化信号**
    *   **[Issue #54152]** **`delegate_task` 的子超时按需覆盖**：为任务委派功能提供细粒度控制，满足如固件编译等耗时任务的需求。
    *   **[Issue #54153]** **软警告机制**：在工具调用预算使用率达到80%时提示agent保存状态，避免因达到上限而丢失中间结果。这直接关系到任务成功率。
    *   **[Issue #54155]** **`command_allowlist` 路径模式匹配**：实现基于路径前缀/Glob 模式的命令自动批准，在安全性和易用性之间取得更好平衡。
    *   **[Issue #54154] & [Issue #54156]** **Kanban调度器增强**：包括捕获工作进程崩溃的诊断信息、在任务失败后提供结构化总结和分解建议。这体现了运维侧对自动化任务管理深度的追求。
*   **路线图指向**
    *   **[PR #54145]** 已经实现了 Kanban 调度器的内存感知和持久化，与上述 Issue #54154 和 #54156 共同指向 **Kanban 调度系统的智能化与健壮性提升**，是当前开发的重点方向之一。
    *   **[Issue #35966]** 对原生客户端的呼声持续高涨，但尚无相关 PR，可能仍在规划或调研阶段，是中期路线上潜在的重量级特性。

### 7. 用户反馈摘要

*   **痛点表达**
    *   *开发环境和安装问题*：`uv` 安装用户反馈 `hermes update` 导致 “orphan virtualenv”，暴露出不同安装路径下的环境隔离问题。Docker 用户遭遇 `profile install` 后无法启动网关的障碍。
    *   *平台锁与死锁*：`jerryfang527` 报告了 Telegram 网关因锁文件残留而完全无法启动的问题，这对于依赖 Telegram 进行交互的用户是灾难性体验。
    *   *工具使用限制*：`c03rad0r` 接连提出多项改进建议，其深层诉求是目前 `agent.max_turns` 和任务委派缺乏透明度和预警，导致长耗时任务容易失败，体现了专业用户对工具可控性的高要求。
*   **满意/潜在满意点**
    *   *社区反馈闭环*：用户 `LineckerN` 报告的 venv 问题被快速关闭，体现了项目组对高影响 Bug 的响应速度，这会是提升用户满意度的关键。
    *   *自动化与安全增强*：`chadsm-sys` 和 `teknium1` 提交的多项安全修复（MCP 源锁定、密钥脱敏、Telegram 权限提前检查）受到了社区的关注，这些措施增强了用户对项目和代码安全性的信心。

### 8. 待处理积压

*   **重点功能请求 (待长期关注)**
    *   **[Issue #35966]** **原生桌面/移动客户端应用**。创建于 2026-05-31，至今超过一个月，虽持续有社区反馈和更新，但状态为 `OPEN` 且无关联 PR。这是社区呼声最高的功能之一，建议维护者给予官方回应或纳入路线图，以回应社区期待。
    *   **[Issue #50745]** **移动端 App**。作为 #35966 的重复问题，同样无进展。多个用户提出相同需求，不能忽视。

*   **待审阅的 PR**
    *   **[PR #44772]** **修复 `hermes update` 删除根目录 Node 依赖**。创建于2026-06-12，至今已近半个月仍未合并。此问题与已关闭的 #39714 性质相似，理论上应优先处理以避免更多用户反馈，建议维护者尽快审阅。

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw 项目动态日报 (2026-06-28)

---

## 1. 今日速览

过去 24 小时，PicoClaw 社区保持中等活跃度：共处理 3 个 Issue（1 个新开、2 个关闭）和 3 个 PR（1 个开放待合并、2 个合并/关闭）。无新版本发布。**亮点**在于两项重要 PR 已被合并——修复了 `mcp add` 命令的参数解析 bug，并引入了内置的代理协作总线，增强了多智能体通信能力。同时，用户报告了一个与 Matrix 加密相关的严重 Bug，尚未有修复 PR。

---

## 2. 版本发布

无新版本发布。

---

## 3. 项目进展

### ✅ 已合并/关闭的重要 PR

- **[#3048] fix(mcp): reject unknown pre-positional flags in add**（已关闭）  
  修复了 `mcp add` 命令中根级持久化标志（如 `--no-color`）被误解析为位置参数的问题。  
  链接：[sipeed/picoclaw PR #3048](https://github.com/sipeed/picoclaw/pull/3048)

- **[#2937] Feat/agent collaboration**（已关闭）  
  引入一等公民的**内部代理协作总线**，支持：  
  - 每个 agent 拥有独立的 mailbox  
  - 协作线程与会话隔离  
  - 结构化消息信封与投递状态追踪  
  - 权限感知的通信机制  
  此功能为多 agent 协同工作奠定了基础设施。  
  链接：[sipeed/picoclaw PR #2937](https://github.com/sipeed/picoclaw/pull/2937)

### 🔄 待合并 PR

- **[#3193] Added simplex channel type**（开放状态）  
  新增 `simplex` 通道类型，扩展消息渠道支持。  
  链接：[sipeed/picoclaw PR #3193](https://github.com/sipeed/picoclaw/pull/3193)

项目整体在**稳定性修复**与**核心能力扩展**两方面均取得进展，尤其是代理协作总线的落地，为未来复杂的 multi-agent 场景扫清关键障碍。

---

## 4. 社区热点

- **#2472** [BUG] `list_dir` 在 Windows 上因路径分隔符不兼容报错（**已关闭**，评论 7 条，👍 1）  
  这是近段时间讨论最集中的 Issue，用户详细描述了 Windows 环境下 `fs.FS` / `os.Root` 强制要求正斜杠而传入反斜杠导致失败。该问题已在 0.2.6 版本中修复，但社区仍关注跨平台兼容性。  
  链接：[sipeed/picoclaw Issue #2472](https://github.com/sipeed/picoclaw/issues/2472)

- **#3114** [Future Request] Telegram 渠道按对话类型的权限分级控制（**已关闭**，评论 2 条，👍 1）  
  用户提出在 Telegram 渠道中区分私聊、群组、频道的操作权限，防止危险命令泄露。虽然该请求被标记为 `stale` 并关闭，但反映出用户对安全边界的迫切需求。  
  链接：[sipeed/picoclaw Issue #3114](https://github.com/sipeed/picoclaw/issues/3114)

---

## 5. Bug 与稳定性

| 严重程度 | Issue | 描述 | 状况 |
|----------|-------|------|------|
| **高** | [#3194](https://github.com/sipeed/picoclaw/issues/3194) | **Received encrypted message but crypto is not enabled** – 用户在使用 Matrix 通道时收到加密消息，但 PicoClaw 未启用加密支持，导致无法正常处理。 | 新开，无评论，无关联修复 PR，需紧急评估。 |
| 中 | [#2472](https://github.com/sipeed/picoclaw/issues/2472) | `list_dir` 在 Windows 因路径分隔符报错 | 已关闭（0.2.6 已修复），无遗留风险。 |

**注意**：#3194 涉及加密信道的兼容性问题，若 Matrix 是项目重点渠道，建议尽快安排修复。

---

## 6. 功能请求与路线图信号

- **Telegram 权限分级控制**（#3114）虽已关闭，但社区呼声较高。结合刚合并的代理协作总线，未来或许可通过 agent 级别的权限模型间接实现。
- **新增 Simplex 通道类型**（PR #3193）正在等待合并，将扩展 PicoClaw 的消息入口，迎合去中心化通信趋势。
- **代理协作总线**（#2937）已合并，明确指向 PicoClaw 的路线图：**从单 agent 向多 agent 协作演进**，预计后续会围绕 mailbox、会话隔离、权限检查等发布更完善的 API 文档。

---

## 7. 用户反馈摘要

- **Windows 兼容性痛点**（#2472）：用户 `ut2or1` 表示，`list_dir` 在 Windows 上因反斜杠问题无法使用，且调试过程发现 Go 标准库 `os.DirFS` 对路径格式敏感。此问题已修复，但反映出当前项目对 Windows 的测试覆盖不足。
- **安全边界需求**（#3114）：用户 `v2up-32mb` 指出，将 PicoClaw 加入 Telegram 群组后，任何成员（若未严格配置白名单）均可执行 `exec`、`write_file` 等危险操作，强烈建议按对话类型（私聊/群组/频道）限制权限。
- **加密消息处理**（#3194）：用户 `Damian-o2` 刚报告了 Matrix 加密消息的异常，尚未提供更多环境细节，但该问题可能影响所有使用加密 Matrix 房间的用户。

---

## 8. 待处理积压

- **#3194** – Received encrypted message but crypto is not enabled  
  状态：OPEN，创建于 2026-06-27，无任何响应。虽有 0 评论，但属于功能性 Bug，影响核心 Matrix 通道的正常使用。建议维护者在下一轮 sprinit 中优先复现并修复。  
  链接：[sipeed/picoclaw Issue #3194](https://github.com/sipeed/picoclaw/issues/3194)

- **#3114** – Telegram 权限分级控制（已关闭但未实际解决）  
  请求虽标记为 `stale` 并关闭，但社区共识仍在，若未来考虑重新开放或作为议题讨论会更合理。目前暂无替代方案，属于**潜在路线图需求**。  
  链接：[sipeed/picoclaw Issue #3114](https://github.com/sipeed/picoclaw/issues/3114)

---

*数据截止时间：2026-06-28 06:00 UTC*  
*项目地址：https://github.com/sipeed/picoclaw*

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是为 NanoClaw 项目生成的 2026-06-28 项目动态日报。

---

# NanoClaw 开源项目动态日报 | 2026-06-28

---

## 1. 今日速览

今日项目整体活跃度较高，但以“提出待办”为主，而非“完成落地”。过去 24 小时内，有 7 个 Pull Request (PR) 处于“待合并”状态（其中 3 个为开放超过一周的老 PR），但无一合并或关闭，显示出维护者在代码审核与合并环节可能存在瓶颈。同时，社区报告了一个关键的 **OpenAI 提供商容器运行时崩溃问题 (Issue #2876)**，涉及核心的 Agent Runner 功能，需要优先关注。此外，一个全新的 **Telegram 富文本渲染特性 PR** 已提交，展示了社区对新平台能力的强烈需求。

**项目健康度评估：** **中高活跃度，但合并效率显著低于提交效率，存在潜在风险。**

---

## 2. 版本发布
今日无新版本发布。

---

## 3. 项目进展
**核心指标：** 过去 24 小时内，**0 个 PR 被合并或关闭**，项目代码进展停滞。

尽管无代码落地，今日堆积的 7 个待合并 PR 内容值得关注，它们代表了项目未来的重要方向：
- **#2877 (feat):** Telegram 原生富文本渲染，将极大提升在该渠道上的消息体验。
- **#2875 (feat):** Coolify 部署方案，这意味着将增加一个新的、更易用的第三方平台部署选项。
- **#2874 (fix):** Signal 集成稳定性修复，防止因 Signal CLI 启动波动导致的容器崩溃循环。
- **#2873 (fix):** 技能代码热更新修复，此 PR 试图解决 `/update-skills` 命令无法正确刷新代码的问题，对开发者体验至关重要。

**总结：** 项目的主要进展体现在“功能提案”和“问题修复方案”的提出上，但代码合并的瓶颈正在形成。

---

## 4. 社区热点
**讨论焦点：** 社区当前讨论的热点主要集中在两个方向：**多提供商支持**和**新渠道功能拓展**。

1.  **[OPEN] Issue #2876: OpenAI 提供商容器崩溃问题**
    - **链接：** `nanocoai/nanoclaw Issue #2876`
    - **热度分析：** 作为今日唯一活跃的 Issue，虽然暂无评论，但其内容直接指向核心功能的严重故障（容器崩溃），预计将很快成为社区关注的焦点。用户尝试将 AI 提供商从默认切换至 OpenAI 时遭遇失败，这暴露了项目在插件化提供商体系入口的健壮性不足。**背后的核心诉求是：用户渴望使用 OpenAI 等主流模型，但系统对配置变更的兼容性测试不足。**

2.  **[OPEN] PR #2877: Telegram 原生富文本渲染**
    - **链接：** `nanocoai/nanoclaw PR #2877`
    - **热度分析：** 该 PR 虽为新提交，但“富文本渲染”是 IM 平台几乎必不可少的功能。这说明有特定用户/开发者对 NanoClaw 在 Telegram 上的消息呈现质量不满意，并愿意贡献代码来解决。**背后诉求是：提升 Telegram 用户体验，使其消息展示能与原生 Bot 应用程序看齐。**

---

## 5. Bug 与稳定性
**严重程度: 高**
-   **[OPEN] Issue #2876: OpenAI 提供商容器崩溃**
    - **描述：** 用户在配置 agent group 的 provider 为 `openai` 后，配置成功保存，但 Agent 在接收到消息并尝试生成新容器时直接崩溃。
    - **链接：** `nanocoai/nanoclaw Issue #2876`
    - **状态：** **无关联修复 PR**。这目前是一个孤立 Bug，可能涉及容器初始化逻辑中对非默认 provider 的环境变量、依赖或挂载点处理不当。此问题应被标记为 **Blocking**，因为它阻止了用户使用 OpenAI 这一最主流的模型。

**严重程度: 中**
-   **[OPEN] PR #2874: Signal 集成启动时崩溃循环**
    - **描述：** PR 描述了 Signal CLI 在启动时的“boot flaps”（启动波动）会导致 NanoClaw 的 Signal channel 不断崩溃重启。
    - **链接：** `nanocoai/nanoclaw PR #2874`
    - **状态：** **已有修复 PR 待合并**。此修复增加了容错和重试机制，一旦合并，将显著提升 Signal 集成的稳定性。

---

## 6. 功能请求与路线图信号
根据今日的 PR 和 Issue，以下功能请求很可能被纳入下一版本的考虑范围：

-   **多 AI 提供商支持:** Issue #2876 虽是 Bug，但背后是用户对“更换 AI 模型提供商”的强烈需求。项目可能需要正视这一点，完善 provider 的容器化支持和测试，确保 `OpenAI`, `Anthropic`, `Google` 等主流提供商能即插即用。
-   **平台/渠道能力增强:**
    - **Telegram 富文本消息:** PR #2877 直接指向下一版本发布清单中的一个重要 Feature。
    - **Signal 稳定性:** PR #2874 是提升渠道稳定性的关键修复，几乎肯定会进入下一个补丁版本。
-   **部署便利性:** PR #2875 (Coolify 部署) 表明社区正在积极为项目寻找更友好的自托管方案，这是项目生态繁荣的信号。

---

## 7. 用户反馈摘要
-   **痛点与场景：**
    - **多模型难以切换：** 用户尝试使用 OpenAI (`gpt-4o`) 作为底层模型，但在切换配置后遭遇了严重的运行时崩溃 (`crash on spawn`)，导致完全无法使用。这表明当前的 provider 架构在运行时的隔离性方面存在问题，**用户期望“配置即生效”的无缝切换体验**。
    - **容器环境健壮性不足:** 用户反馈在 `agent-runner` 中的容器在启动新实例时崩溃，暗示容器创建、网络、或环境变量传递的流程有缺陷。

---

## 8. 待处理积压
**警报：** 以下 PR 长期处于开放且无被合并迹象，可能因维护者无暇处理或存在争议，建议维护团队尽快评估。

1.  **[OPEN] PR #2822: refactor(container-runner): drop dead /workspace/global mount**
    - **创建时间：** 2026-06-20 (已开放 8 天)
    - **更新于：** 2026-06-27
    - **链接：** `nanocoai/nanoclaw PR #2822`
    - **重要性：** 此重构涉及到移除一个“死”的全局挂载点，有助于清理代码和潜在的容器问题，属于技术债务清理。

2.  **[OPEN] PR #2823: fix: remove groups/global/CLAUDE.md**
    - **创建时间：** 2026-06-20 (已开放 8 天)
    - **更新于：** 2026-06-27
    - **链接：** `nanocoai/nanoclaw PR #2823`
    - **重要性：** 修复了宿主机在每次启动时都删除该文件的问题。这是一个稳定性和流程上的 Bug 修复。

3.  **[OPEN] PR #2824: fix: drop stale "Global Memory" instruction**
    - **创建时间：** 2026-06-20 (已开放 8 天)
    - **更新于：** 2026-06-27
    - **链接：** `nanocoai/nanoclaw PR #2824`
    - **重要性：** 从 Prompt 中移除过时的全局记忆指令，这对保持 Agent 行为的准确性和一致性至关重要。

**建议：** 这三条 PR 均由同一贡献者 (`CutSnake01`) 提交，都是低风险、高价值的修复。长期不合并不仅会打击贡献者的积极性，也可能导致项目代码落后于社区的最佳实践。

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

好的，这是根据您提供的 GitHub 数据生成的 NullClaw 项目动态日报。

---

# NullClaw 项目动态日报 | 2026-06-28

## 今日速览

项目昨日呈现低频但高效的活跃状态。核心亮点是社区贡献者提交了一个关于 **Agent 工具审批流程** 的重要功能 PR，填补了此前版本中该机制的缺失，对提升 AI Agent 的安全交互性有显著价值。另一方面，一个在 Android/Termux 环境下的 **构建 Bug** 依然处于未解决状态，且长期缺乏维护者关注，是当前项目稳定性的一个隐忧。总体来说，项目在功能演进上迈出了坚实一步，但遗留的跨平台兼容性问题需要优先处理。

## 版本发布
无

## 项目进展

今日无合并或关闭的 Pull Request，但有一项重要的新功能 PR 提交。

-   **PR #969 (待合并):** [feat(agent): structured approval_request / approval_response flow](https://github.com/nullclaw/nullclaw/pull/969)
    -   **核心贡献:** 由社区成员 `valonmulolli` 提交，实现了针对 `shell tool`（及任何需要用户批准才能执行的工具）的 **结构化审批流程**。
    -   **工作原理:** 当工具需要用户批准时，Agent 会捕获该请求，通过 Server-Sent Events (SSE) 推送给客户端，由客户端渲染审批 UI。用户确认后，结果再通过 `approval_response` 返回给 Agent 继续执行。
    -   **项目意义:** 这项功能是构建安全、可控的 AI Agent 的关键一环，解决了此前用户无法对高危操作（如执行 Shell 命令）进行中断或确认的痛点，显著提升了项目的可用性和安全性。

## 社区热点

-   **#868 [Bug] zig build fails on Android/Termux**: [Issue 链接](https://github.com/nullclaw/nullclaw/issues/868)
    -   **热度分析:** 该 Issue 已存在超过两个月，共有 4 条评论。虽然评论数不多，但问题本身影响特定用户群体（Android 开发者和 Termux 用户），其长期未解决的状态已成为社区关注点。
    -   **背后诉求:** 用户期望 NullClaw 能在移动端（Android）的终端环境（Termux）中顺畅运行，这反映了项目向更广泛开发者场景拓展的需求，目前该能力受到阻碍。

## Bug 与稳定性

-   **严重 Bug (长期未解决):**
    -   **#868:** [zig build fails on Android/Termux (aarch64) with AccessDenied on options.zig linkat](https://github.com/nullclaw/nullclaw/issues/868)
    -   **严重程度:** 高
    -   **问题描述:** 在 aarch64 架构的 Android 设备上，通过 Termux 环境使用 `zig build` 构建项目时，因 `options.zig linkat` 操作出现 `AccessDenied` 错误而失败。该问题阻碍了用户在移动端进行本地构建和开发。
    -   **状态:** 无关联的修复 PR，维护者未响应。

## 功能请求与路线图信号

-   **确定性候选功能：**
    -   **PR #969 (结构化审批流):** 已被提交为 PR，说明该功能已从需求阶段进入实现阶段。这很可能成为下一版本 `v2026.x` 的重要组成部分，标志着 Agent 交互模式的成熟化。
-   **潜在需求：**
    -   **Issue #868 (Android/Termux 构建支持):** 虽然是一个 Bug，但其本质是用户对 **跨平台构建支持** 的功能请求。修复此问题将直接扩展项目的支持平台列表。

## 用户反馈摘要

-   **真实用户痛点 (来自 #868):**
    -   **环境:** 用户明确描述了其设备（Xiaomi Redmi Note 9）、操作系统（LineageOS 22.2，基于 Android）、终端（Termux）和 Zig 版本（0.16.0）。
    -   **问题复现:** 用户成功通过 `pkg install zig` 安装了 Zig，但在执行构建时失败，错误信息指向文件系统的权限问题（`AccessDenied` 和 `linkat`）。
    -   **不满意点:** 构建流程在标准 Linux/桌面环境下可能正常，但在 Android 这一特定但重要的平台上存在障碍，且长期未获响应。

## 待处理积压

-   **关键积压 Issue:**
    -   **#868:** [Android/Termux 环境构建失败](https://github.com/nullclaw/nullclaw/issues/868) (创建于2026-04-23，最后更新于2026-06-27)
    -   **建议:** 该 Issue 已存在 2 个月，且项目维护者至今未予反馈。强烈建议维护者或具有相关经验的社区成员介入排查，这既是对现有用户负责，也能吸引更多移动端开发者参与项目。修复方向可能涉及调整 `linkat` 系统调用的使用方式或依赖的 Zig 标准库版本。

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，根据您提供的 IronClaw GitHub 数据，我为您生成了 2026 年 6 月 28 日的项目动态日报。

---

## IronClaw 开源项目动态日报 (2026-06-28)

**分析师点评：** 项目今日活跃度极高，状态为 **高度活跃**。核心开发者 (zetyquickly) 主导的 **Capability Policy (能力策略)** 功能模块（Epic #5261）已发生关键进展，系列相关的 7 个 Issue 在今日关闭，标志着核心功能的框架代码已合并。同时，社区贡献持续涌入，涵盖 WebUI 修复、集成测试框架以及 CI 测试矩阵扩展。新出现的 Google OAuth 令牌刷新 BUG 需维护者重点关注。

### 1. 今日速览

-   **核心功能冲刺：** 备受期待的 **Capability Policy (能力策略)** 功能取得里程碑式进展，其包含的 `REST-created local users`、`Admin REST surface`、`Four-dimension policy` 等 7 个关联 Issue 于今日全部关闭，代码已成功合入主分支。
-   **社区贡献活跃：** 项目收到 50 条 PR 更新，其中 24 条已合并/关闭。集成测试框架 (Reborn Integration-Test Framework) 及 WebUI v2 改进（如 Retry 按钮、活动卡片显示）等方面的贡献尤为突出。
-   **自动化测试警报：** 持续了一个月的 Nightly E2E 测试 (Issue #4108) 仍在失败，这已成为一个长期存在的稳定性隐患。
-   **新出现 BUG：** 报告了一个关于 Google OAuth 令牌约一小时强制过期重认证的 BUG (Issue #5378)，影响体验。
-   **社区热点：** 围绕 `Reborn` 栈的稳定性改进和开发者体验是该社区当前关注的焦点，多篇提交旨在解决工具调用错误详情展示、流式响应等问题。

### 3. 版本发布

**无新版本发布**

---

### 4. 项目进展

今日项目迈出了关键一步，即 **Capability Policy** 功能的代码基础已准备完毕。以下是今日完成的重要 PR：

-   **`[CLOSED]` [Capability-Policy] 核心模型入库**
    -   **PR:** [#5262 feat(capability-policy): policy model — ironclaw_capability_policy crate (epic #5261)](https://github.com/nearai/ironclaw/issues/5262)
    -   **概要：** 作为 Epic #5261 的第一部分，合并了 `ironclaw_capability_policy` crate，定义了四维权限模型、优先级规则和内存存储解析器。这是实施后续所有权限控制逻辑的基石。

-   **`[CLOSED]` [Reborn] 修复非 OAuth 触发运行记录**
    -   **PR:** [#4970 fix(reborn): record Delivered (not Skipped) for non-OAuth auth-denied triggered runs](https://github.com/nearai/ironclaw/issues/4970)
    -   **概要：** 修复了自动化运行在遇到非 OAuth (如 API Key) 认证失败时，状态被错误记录为 "Skipped" 的问题，修正为 "Delivered"，保证了投递通知到用户的准确性。

-   **`[CLOSED]` [Reborn] 集成测试框架基础建立**
    -   **PR:** [#5381 Reborn integration-test framework (slices 1–2)](https://github.com/nearai/ironclaw/issues/5381)
    -   **概要：** 建立了全新的 Reborn 集成测试框架，允许在进程中模拟真实堆栈（LLM Provider、文件系统等）运行完整的 Reborn 对话流程，仅对模型返回 API 进行桩打，极大提升了后端逻辑的测试效率和覆盖度。

-   **`[CLOSED]` [WebUI v2] 前端工具链锁定**
    -   **PR:** [#5384 [codex] Pin WebUI v2 frontend Node tooling](https://github.com/nearai/ironclaw/issues/5384)
    -   **概要：** 锁定了 WebUI v2 前端的 Node.js 版本 (Node 22)，并添加了构建脚本重跑提示，解决了因工具链不一致导致的开发环境问题。

### 5. 社区热点

-   **Issue #5378: Google OAuth 令牌刷新失败**
    -   **链接：** [Google OAuth token refresh fails with BackendUnavailable](https://github.com/nearai/ironclaw/issues/5378)
    -   **诉求：** 用户报告在使用 Google 驱动功能（如 Gmail 发送）时，OAuth 令牌约每小时强制过期，导致需要频繁重认证。
    -   **分析:** 这直接影响了依赖 Google 服务的自动化任务和集成体验，是影响用户粘性的关键痛点。社区关注度极高，目前已有 fix PR [#5388](https://github.com/nearai/ironclaw/issues/5388) 正在处理中，预计很快会得到解决。
-   **PR #5306: `ask-each-time` 批准恢复循环修复**
    -   **链接：** [fix(reborn): ask-each-time approval resume loop](https://github.com/nearai/ironclaw/issues/5306)
    -   **分析：** 标题明确指向一个用户交互流程上的 BUG，即用户批准一次 “每次询问” 的工具后，下次调用时仍会再次要求批准。这个话题触及了核心的用户体验，讨论热度很高。

### 6. Bug 与稳定性

| 严重程度 | 问题描述 | 链接 | 关联 PR | 状态 |
| :--- | :--- | :--- | :--- | :--- |
| **严重** | Nightly E2E 自动化测试持续失败已超过一个月，反映测试基础设施或代码存在不稳定。 | [Issue #4108](https://github.com/nearai/ironclaw/issues/4108) | N/A | **长期开放** |
| **高** | Google OAuth 令牌因 `BackendUnavailable` 错误频繁过期，严重阻塞 Google 功能。 | [Issue #5378](https://github.com/nearai/ironclaw/issues/5378) | [#5388](https://github.com/nearai/ironclaw/issues/5388) | **有修复 PR** |
| **中** | WebUI v2 聊天中的 “重试（Retry）” 按钮失效，点击后并未重新发送消息。 | [PR #5365](https://github.com/nearai/ironclaw/issues/5365) | #5365 (本身即为修复 PR) | **有修复 PR** |
| **中** | Reborn 中工具调用失败后返回错误的 `invalid_input` 通用错误信息，不便于排查。 | [PR #5338](https://github.com/nearai/ironclaw/issues/5338) | #5338 (本身即为修复 PR) | **有修复 PR** |

### 7. 功能请求与路线图信号

-   **`[OPEN]` Capability Policy 细节配置 (Issue #5385)**
    -   **链接：** [Add Capability Policy](https://github.com/nearai/ironclaw/issues/5385)
    -   **信号：** 尽管核心模型已合并，但该 Issue 提出了更细粒度的用户（Owner/Admin/Member）配置需求，包括更具体的逻辑（如：用户能否创建Agent、访问API等）。这表明用户期待的 **Capability Policy** 远不止于单纯的权限定义，而是更丰富的角色管理。这很可能是下一阶段的开发重点。

-   **WebUI v2 布局/功能改进**
    -   **信号：**
        -   PR [#5084](https://github.com/nearai/ironclaw/issues/5084) (Redesign the automations page): 社区成员正在对自动化页面进行UI/UX重设计，表明前端用户体验是社区关注的重点。
        -   PR [#5365](https://github.com/nearai/ironclaw/issues/5365) (fix Retry button) 和 [#5297](https://github.com/nearai/ironclaw/issues/5297) (fix stale gate projection): 显示出开发团队正在系统性地排查并修复 WebUI v2 里的交互问题，提升产品成熟度。

-   **测试框架扩展**
    -   **信号：**
        -   PR [#5354](https://github.com/nearai/ironclaw/issues/5354) (Add Reborn WebUI v2 live QA canary) & [#5380](https://github.com/nearai/ironclaw/issues/5380) (expand Reborn WebUIv2 QA matrix coverage): 显示了项目正在大力加强质量保障体系，引入端到端和金丝雀测试。
        -   PR [#5387](https://github.com/nearai/ironclaw/issues/5387) (URL-keyed HTTP matcher) & [#5386](https://github.com/nearai/ironclaw/issues/5386) (embeddings fake): 是昨天 #5381 建立的集成测试框架的有力补充，正在快速完善。

### 8. 用户反馈摘要

-   **OAuth 问题**：用户（`thisisjoshford`）描述了使用 Google OAuth 时，每小时重新认证的困扰。这一体验非常糟糕，会打断工作流和自动化任务。
-   **体验修复**：`italic-jinxin` 在 PR #5306 中修复的批准恢复循环直接反映了用户在需要审批的工具调用时的困惑，用户希望“一次性批准”能够生效。而在 PR #5338 的修复中，用户得到的是模糊的 "driver protocol error" 错误，这不利于其理解问题根源。

### 9. 待处理积压

-   **长期失效 Nightly E2E 测试**
    -   **链接：** [Issue #4108 Nightly E2E failed](https://github.com/nearai/ironclaw/issues/4108)
    -   **警告：** 此 Issue 自 2026-05-27 起因失败被开启，距今已超过一个月，且至今无任何评论或进展。这是项目稳定性的一个红灯信号，提醒维护者需要关注 CI 基础设施的健康度。

-   **长期未合并的 Dependencies Bump PR**
    -   **链接：** [PR #3706 chore(deps): bump postcss, @remotion/cli...](https://github.com/nearai/ironclaw/issues/3706)
    -   **警告：** 此 PR 由 `dependabot` 于 2026-05-16 创建，旨在更新文档视频目录的依赖，但已长时间未合并，可能导致文档构建环境漏洞累积。建议维护者检查并处理。

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

## LobsterAI 项目动态日报 — 2026-06-28

### 1. 今日速览
过去24小时内，项目无新版本发布，活跃度维持在较低水平。主要动作为**自动清理陈旧的 Issues 和 PR**：4 条长期未跟进的 Issues 被关闭（均为 stale 标记），4 条 PR 被合并/关闭（其中 3 条为此前遗留的 PR，1 条为较旧的 hotfix）。目前仍有 2 条开放 PR 处于 stale 状态，等待维护者审核。整体来看，项目近期开发节奏放缓，社区讨论热度不高，但核心功能模块（技能、定时任务、会话管理）仍在持续优化中。

### 2. 版本发布
无新版本发布。

### 3. 项目进展
今日合并/关闭的 PR 集中在**技能管理与制品预览**方向的遗留改进：

- **PR #1441** — `feat(artifacts): add extensible preview pipeline for HTML, React and Mermaid`  
  该 PR 是一个冲突解决与 bug 修复版本，源自 #1011，为 Cowork 会话增加了 HTML、React 和 Mermaid 的可扩展预览流水线。目前已合并，标志着「制品预览」能力正式落地。  
  [查看 PR](https://github.com/netease-youdao/LobsterAI/pull/1441)

- **PR #1440** — `feat(cowork): 将已选技能标签移至输入框内顶部展示`  
  优化了技能标签在输入框内的布局，将 ActiveSkillBadge 从底部工具栏移至 textarea 上方，并增加 className prop 支持。合并后改善了多技能选择时的界面拥挤问题。  
  [查看 PR](https://github.com/netease-youdao/LobsterAI/pull/1440)

- **PR #1445** — `fix(skills): 修复技能重复导入无校验及 zip 导入目录名异常的问题`  
  修复了三个导入渠道（zip、文件夹、GitHub）中技能重复无校验、目录名随机化的 Bug，增加了从 SKILL.md 读取技能名的辅助函数，并引入重复检测逻辑。  
  [查看 PR](https://github.com/netease-youdao/LobsterAI/pull/1445)

- **PR #1001** — `hotfix：增加对 sse 和 流式http 的mcp支持`（较旧 PR，今日关闭）  
  解决了 MCP 配置中仅支持 stdio 传输类型的问题，使 SSE 和流式 HTTP 配置可正常生效。  
  [查看 PR](https://github.com/netease-youdao/LobsterAI/pull/1001)

> 以上 4 条 PR 均已被合并/关闭，项目在**技能体验、制品预览、MCP 兼容性**方面向前迈进了关键步子。

### 4. 社区热点
今日无高热度讨论。评论最多的 Issue 为 #1443（3 条评论），主题为**openclaw 新版本适配问询**，用户报告升级至 v2026.3.24 后服务启动失败，询问团队是否计划支持。该 Issue 已被标记为 stale 并关闭，但未获得官方明确答复。  
[查看 Issue #1443](https://github.com/netease-youdao/LobsterAI/issues/1443)

### 5. Bug 与稳定性
今日被关闭的 4 条 Issues 均为陈旧 Bug，严重程度中等，且均已有对应修复或关闭理由：

| Issue | 标题 | 严重程度 | 修复状态 |
|-------|------|---------|---------|
| #1443 | 有计划支持新版本的openclaw吗？ | 中（兼容性问题） | 已关闭（stale），无明确修复 |
| #1437 | 创建定时任务时，计划选择不重复，清空日历，点击【创建任务】按钮没反应 | 中（功能阻塞） | 已关闭（stale） |
| #1439 | 上传技能已停用，对话中仍然可以调用 | 高（逻辑错误） | 已关闭（stale），但对应修复 PR #1445 已合并 |
| #1442 | Agent添加技能，对话后引用的技能不展示 | 中（体验问题） | 已关闭（stale），对应修复 PR #1494（待合并） |

其中 #1439 的问题（技能停用后仍被调用）已被 #1445 修复，而 #1442 的技能展示丢失问题，开放 PR #1494 已提出按会话独立管理技能选择状态的方案，但尚未合并。

### 6. 功能请求与路线图信号
今日无新功能请求。但从开放 PR 可看出，项目正在推进以下方向：

- **定时任务模块 UI 全面升级**（PR #1488）：卡片网格视图、搜索筛选、历史任务按日期分组展示。该 PR 仍开放且处于 stale，若获得合并，将显著提升定时任务管理体验。  
  [查看 PR #1488](https://github.com/netease-youdao/LobsterAI/pull/1488)

- **技能选择状态按会话独立管理**（PR #1494）：将全局存储改为按会话存储，避免技能选择在切换会话时串扰。此改动与 #1442 的用户反馈直接对应，预计将被纳入下一版本。  
  [查看 PR #1494](https://github.com/netease-youdao/LobsterAI/pull/1494)

### 7. 用户反馈摘要
从近期 Issues 评论中提炼的典型用户痛点：

- **技能停用失效**：用户 `devilszy` 反馈关闭技能后对话中仍能触发调用，质疑技能关闭与路由之间的联动逻辑。（#1439）
- **技能选择状态混乱**：同一用户还指出添加技能后在 Agent 对话中技能不展示，切换会话后又复现，对其作用机制表示困惑（#1442）
- **定时任务创建无响应**：用户 `xuzx-code` 在创建一次性定时任务时，点击创建按钮无反应且无错误提示，体验差（#1437）
- **openclaw 兼容性**：用户 `Juzisuan965` 升级 openclaw 后无法启动，期望官方尽快适配（#1443）

### 8. 待处理积压
以下长期未响应的 PR 和 Issue 值得维护者优先关注：

- **PR #1488**（`feat(scheduledTask): 定时任务模块 UI 全面升级`）  
  创建于 2026-04-05，更新于 2026-06-28，仍处于 open+stale 状态。该 PR 改动较多，涉及核心功能重构，建议尽快安排 Code Review。  
  [查看 PR #1488](https://github.com/netease-youdao/LobsterAI/pull/1488)

- **PR #1494**（`fix(cowork): 技能选择状态改为按会话独立管理`）  
  创建于 2026-04-06，更新于 2026-06-28，同样处于 stale。该 PR 直接关联 #1442 的用户反馈，合并后可解决一个长期体验痛点。  
  [查看 PR #1494](https://github.com/netease-youdao/LobsterAI/pull/1494)

- **Issue #1443**（openclaw 新版本适配）  
  虽已关闭，但作为外部依赖兼容性问题，若用户社区仍有需求，建议在路线图中评估适配优先级。  
  [查看 Issue #1443](https://github.com/netease-youdao/LobsterAI/issues/1443)

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis 项目日报 | 2026-06-28

---

## 1. 今日速览

- 过去24小时项目活跃度中等偏低：仅新增1个Bug Issue，无Issue关闭；4个Pull Request全部处于待合并状态，无已合并或关闭的PR。
- 贡献者`resumeparseeval`提交了3个修复性PR，聚焦于代理（agent）与网关（gateway）模块的可靠性与兼容性，表明团队正在持续解决社区反馈的痛点。
- 无新版本发布，项目当前处于迭代修复阶段，整体健康度良好，但缺乏关键性进展合并动作，长期未合并的PR（#1098）仍悬而未决。
- 社区热点集中在**模型兼容性**与**资源管理**上：小型本地模型（如Gemma 4、oMLX）的参数序列化问题持续成为修复主线。

---

## 2. 版本发布

无新版本发布。

---

## 3. 项目进展

今日无任何PR被合并或关闭，仅新增4个待合并PR。这些PR若合并后将解决以下核心问题：

- **#1139 [fix(gateway)]**：修复`metrics`特性强制启用`moltis-matrix`依赖的问题，避免非Matrix用户引入不必要的`matrix-sdk`编译开销。属于构建优化与依赖隔离。
- **#1138 [fix(agents)]**：解决全分辨率图片（如4032×3024）以Base64内联嵌入时，单张图片消耗约350K tokens导致上下文溢出的问题。提出在进入模型上下文前先降采样图像，这是对多模态场景下资源预算的关键改进。
- **#1136 [fix(agents)]**：针对小模型（Gemma 4、oMLX）将标量工具参数以JSON字符串形式输出（如`"true"`、`"5000"`）而非原始值的情况，增加预分派前的类型强制转换（coerce），提升工具调用成功率。
- **#1098 [fix(browser)]**：已有25天未合并，解决浏览器工具调用中可选参数显式传入`null`时serde反序列化失败的问题，同样针对小模型的行为不一致。

**项目整体迈进的判断**：修复方向明确——强化对小模型/本地模型的兼容性、优化资源使用（图像上下文体积、依赖链）、提升构建灵活性。但**合并速度偏慢**，特别是#1098已长时间等待。

---

## 4. 社区热点

由于今日仅有1条新Issue且无评论，PR也均无评论数据，暂无高互动话题。但从PR摘要及Issue内容可提炼潜在热点：

- **#1137 [Bug] Apple Container ID 超过名称长度限制**：作者`holgzn`报告在Apple平台（可能指macOS/iOS容器）上Container ID超出名称长度限制。这是与平台集成相关的边界错误，虽无评论但属于平台适配问题，可能影响macOS用户。
- **#1136、#1138、#1098** 共同指向**小模型工具调用兼容性**，这是社区用户（尤其是本地模型使用者）最关心的痛点。PR作者`resumeparseeval`持续贡献，说明该方向需求迫切。

---

## 5. Bug 与稳定性

| 严重程度 | Issue/PR | 描述 | 是否有修复PR |
|----------|----------|------|--------------|
| 高 | #1137 | Apple Container ID超长导致创建容器失败，阻塞macOS用户 | 无，待调查 |
| 高 | #1138 | 全分辨率图片嵌入导致上下文预算溢出，每次对话均被拒绝 | 已有PR #1138 待合并 |
| 中 | #1136 | 小模型标量参数以字符串形式传递导致验证失败 | 已有PR #1136 待合并 |
| 中 | #1098 | 可选参数显式`null`导致反序列化错误 | 已有PR #1098 待合并（已25天） |

**长期未合并风险**：#1098 状态为Open超过25天，可能因多轮评审或兼容性测试滞后，建议维护者优先合并或给出明确反馈。

---

## 6. 功能请求与路线图信号

今日无明确的新功能请求Issue。但从PR摘要可判断以下方向很可能被纳入下一版本：

- **图像预处理管线**：#1138 提出的“进入模型上下文前降采样超大图片”是一项功能增强，将改变多模态代理的行为逻辑，属于路线图上的重要优化。
- **小模型适配层**：#1136 与 #1098 分别针对字符串参数和null参数的处理，暗示团队计划建立一套对“非严格JSON输出”模型的兼容机制，可能作为agent模块的通用预处理层。

**潜在路线图信号**：对**本地模型生态**（如Gemma 4、oMLX）的支持优先级正在提升，与当前开源社区的趋势一致。

---

## 7. 用户反馈摘要

基于今日唯一Issue #1137 的初始描述，作者报告了Apple平台容器ID名称超限的bug。虽无评论，但可以推测作者：

- 在使用Apple设备（可能是macOS）运行Moltis容器化部署时遇到。
- 已严格按照bug报告模板检查了现有issue并更新了最新版本，说明用户有一定技术素养且积极反馈。

其他PR暂无用户评论数据，但从PR解决的具体问题（图片过大、模型输出格式异常）可推断真实用户场景：

- 用户使用手机拍摄的照片直接通过Agent上传对话，导致上下文爆炸。
- 用户使用本地小模型（如Gemma 4）调用浏览器工具时频繁遇到参数错误。

**满意度推断**：社区对性能（图片处理）和稳定性（工具调用）有更高期待；当前修复若合并将显著提升体验。

---

## 8. 待处理积压

| 类型 | 编号 | 标题 | 创建时间 | 当前状态 | 建议 |
|------|------|------|----------|----------|------|
| PR | #1098 | fix(browser): tolerate null optional params... | 2026-06-03 (25天) | Open，待合并 | 优先级应设为高；长期未合并可能让提PR者失去动力，建议维护者本周内评审或提出修改意见 |
| Issue | #1137 | Apple Container ID exceeds name limit | 2026-06-27 (1天) | Open，无评论 | 新报告，需尽快确认复现并分配处理人 |

**其他注意**：今日4个PR均为同一作者`resumeparseeval`，该贡献者近期活跃度极高，其贡献应得到及时响应以维持贡献热情。

---

*数据来源：Moltis GitHub 仓库 (github.com/moltis-org/moltis)，截止 2026-06-28 23:59 UTC。*

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw 项目日报 – 2026-06-28

## 1. 今日速览

过去 24 小时项目保持了中等活跃度：2 个新 Issue 开启，6 个 Pull Request 处于开放待审状态。测试覆盖工作持续推进：三个包含 120 个单元测试的 PR（#5581、#5422、#5423）已更新至 **AgentScope 2.0** 主线，标志着后端基础设施测试已接近完成。社区问题方面，一个涉及 Ascend-vLLM 连接回归的严重 Issue 被提出，同时有两个功能型 PR（Matrix 流式输出、上下文压缩阈值修复）和插件安装修复 PR 仍在等待合入。无新版本发布。

## 2. 版本发布

无

## 3. 项目进展

今日没有 PR 被合并或关闭，但以下开放 PR 标志着关键进展：

- **测试基础设施（W1–W3 Sprint）**  
  #5423（crons 模块，51 用例）、#5422（chats 模块，38 用例）、#5581（app-infra 层，31 用例）均已完成对 **AgentScope 2.0**（`agentscope==2.0.2`）的适配，并更新至最新上游。这三组 PR 合计 **120 个单元测试**，覆盖了 `agent_context`、`console_push_store`、`workspace_migration`、`inbox_trace_store`、`JsonChatRepository`、`CronManager`、`JsonJobRepository` 等核心模块。一旦合入，将大幅提升后端回归检测能力。

- **关键修复**  
  - #5586（首次贡献者，`fix(context): prioritize runtime model over static config`）修复了用户通过 UI 切换会话内模型后，`light_context_config` 的压缩阈值仍读取静态配置的 bug。该问题直接影响长对话的上下文管理精度。  
  - #5568 解决了因 AgentScope 1.x → 2.0 迁移导致 5 个官方插件从 CDN 目录安装失败的问题，待合入后可恢复插件市场可用性。

- **功能开发**  
  - #5585 为 Matrix 频道添加了类似 Discord 的流式输出模式，减少用户等待首 Token 时间（TTFT）。

## 4. 社区热点

- **Issues**  
  - **#5584 [Question]: 无法连接自定义的 ascend-vllm 模型**  
    （https://agentscope-ai/QwenPaw/issues/5584）  
    用户报告从 1.1.7 版本升级后，QwenPaw 对话框始终提示连接错误，但 vLLM 后端正常，且其他客户端可正常交互。错误为 `openai.APIConnectionError`。该问题已收获 1 条评论，被标注为 `question`。背后反映的是模型配置向后兼容性问题，可能涉及 AgentScope 2.0 中 API 客户端初始化逻辑的变更。目前缺乏项目维护者的明确回复，需要优先调查。

- **PRs**  
  - #5586（首次贡献者）因修复了会话模型切换这一用户高频操作中的逻辑缺陷，虽无大量评论，但具有较高实践价值。

## 5. Bug 与稳定性

| 严重程度 | Issue/PR | 描述 | 修复状态 |
|----------|----------|------|----------|
| **高** | #5584 | 自定义 Ascend-vLLM 模型连接失败（`APIConnectionError`），1.1.7 正常，之后版本全量回归。影响已配置自定义模型的所有用户。 | 无关联修复 PR |
| **中** | #5586 | 对话内切换模型后上下文压缩阈值错误，导致可能触发不必要的截断或超出长度限制。 | 已有修复 PR #5586（待合入） |
| **低** | #5583 | 聊天界面右侧弹出层默认选中背景不明显，属于 UI 易用性瑕疵。 | 无关联修复 |
| **中** | #5568 | 5 个官方插件在 QwenPaw 2.0 上从 CDN 安装失败，阻断插件市场核心功能。 | 已有修复 PR #5568（待合入） |

## 6. 功能请求与路线图信号

- **Matrix 流式输出（#5585）**  
  用户 `Morxi` 提交了为 Matrix 频道增加流式回复模式的功能 PR，使 Matrix 端行为与 Discord 端一致。该功能可提升 Matrix 用户的交互体验，如果社区反馈积极，很可能被纳入下一修补版本。

- **无新功能请求 Issue**  
  今日开启的两个 Issue 均为 `question` 类型（#5584 连接问题、#5583 UI 问题），未出现正式的功能请求（feature request）。路线图信号主要来自 PR 侧。

## 7. 用户反馈摘要

- **Ascend-vLLM 连接回归**（#5584）：用户明确表示“1.1.7 版本可连，后续版本均无法连接”，且“vLLM 后端显示一切正常，其他软件均可正常对话”。这指向 AgentScope 2.0 在 API 客户端错误处理或超时设置上的退化，用户痛点清晰且验证条件简单，建议维护者优先复现。
- **UI 细节**（#5583）：用户反映“聊天界面右侧对话弹出层默认选中背景不明显”，虽为小问题，但侧面说明社区对 UI 细节有较高敏感度，项目可考虑在下一版调整 CSS 对比度。

## 8. 待处理积压

- **长期未合并的测试 PR**  
  #5422（6 月 23 日开启）、#5423（6 月 23 日开启）虽已更新至最新 AgentScope 2.0，但已待审超过 5 天，可能阻塞后续基础设施变更。建议维护者尽快 review 并合入。
- **插件安装修复 PR**（#5568，6 月 26 日开启）同样等待合并，若继续积压会导致插件市场长期不可用，影响用户体验。
- **无其他长期未响应的重要 Issue**（#5584 为 1 天前新开，尚未到“积压”程度）。

---

*报表生成时间：2026-06-28 23:59 UTC | 数据来源：GitHub repository agentscope-ai/CoPaw（仓库名为 QwenPaw）*

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是基于您提供的 GitHub 数据生成的 ZeroClaw 项目动态日报。

---

## ZeroClaw 项目动态日报 | 2026-06-28

### 今日速览

ZeroClaw 项目今日社区活跃度极高，但合并/关闭率较低。过去24小时内产生了 **50 条 Pull Request**，但其中 **40 条仍处于待合并状态**，仅有 10 条被处理，这暗示了维护团队在审查速度上可能面临瓶颈。Issue 方面，社区讨论热烈，除常规 Bug 修复外，**一份关于重构底层通信协议（Wire Protocol）的 RFC（#8396）** 以及 **多项涉及 Telegram、WhatsApp、文件系统的新功能请求** 成为今日焦点。总体来看，项目处于功能密集开发期，但合并效率有待提升。

### 项目进展

过去24小时内，项目的主要进展体现在对已合并/关闭 PR 的持续跟进和 Bug 修复上。

- **配置与体验优化**：`[CLOSED] #7808` - **Config** 的 Bug 被成功解决。该问题修复了用户在设置加密密钥时，CLI 提示完全隐藏输入而无任何反馈（如字符数变化）的问题，提升了配置界面的用户体验。
- **核心功能修复**：多个旨在修复核心功能的 PR 正在积极推进更新中，表明项目在稳定现有功能上投入了较多精力。
    - `#8414` - **WhatsApp Web Channel**：正在修复一条**回归问题**，即允许 WhatsApp Web 使用 `/models`、`/model` 等运行时命令。在 v0.7.5 补丁中丢失的功能正在恢复。
    - `#7858` - **Channel 本地化**：正在为频道核心运行时命令（如 `/new`、`/stop`）添加多语言支持，不再仅限硬编码英文。
    - `#8343` - **CI 发布流程**：正在改进发布流程，使其从统一的特性注册表中构建工件，以确保所有分发渠道（如安装脚本、容器）的一致性。

### 社区热点

今日社区讨论和关注的热点主要集中在技术架构演进和特定平台的功能增强上。

1.  **RFC: Wire-Protocol-First 提供者模型 (#8396)**：该 Issue 提出对 ZeroClaw 的提供者（Provider）模型进行根本性重构，意图将 `wire_api` 作为核心组织轴。这是一项**高风险**但可能带来巨大架构优势的提案，引发了社区的初步关注和讨论（1条评论）。这反映了社区对项目长期技术债务和扩展性的思考。
    - 链接：https://github.com/zeroclaw-labs/zeroclaw/issues/8396

2.  **Telegram Bot API 10.1 富文本消息支持 (#8415)**：用户 `klonuo` 提出了在 Telegram Bot 中实现富文本消息的需求，以提供更好的用户体验。这源自之前关于 Telegram 表格渲染问题的讨论 (#8390)。该请求强调了对特定平台能力的适配和补齐。
    - 链接：https://github.com/zeroclaw-labs/zeroclaw/issues/8415

3.  **WhatsApp Group 被动上下文 (#8389)**：PR `#8389` 是一个大规模、高风险的功能增强，旨在为 WhatsApp 群组添加“被动通道上下文”（passive channel context）支持。该特性允许群组将某些消息标记为“仅历史记录”，从而优化上下文管理。此 PR 涉及超过20个模块，显示了它的影响范围之广和重要性之高。
    - 链接：https://github.com/zeroclaw-labs/zeroclaw/pull/8389

### Bug 与稳定性

今日没有发现严重级别为 `S0` 或 `S1` 的崩溃类 Bug。已报告的 Bug 主要涉及功能退化与体验问题。

- **`[Bug] CLI secret prompts give no feedback after paste (#7808)` - [S2 功能退化] - 已关闭**：该问题已通过维护者的工作得到修复。
    - 链接：https://github.com/zeroclaw-labs/zeroclaw/issues/7808
- **`[Fix] 恢复 WhatsApp Web 模型命令 (#8414)` - [回归问题]**：这是一个已确认的回归问题，对应的修复 PR 正在审查中。
    - 链接：https://github.com/zeroclaw-labs/zeroclaw/pull/8414
- **`[Fix] 网关端口占用无提示 (#8115)` - [中等风险]**：该 PR 修复了当网关端口已被占用时，守护进程静默降级的问题，现在会直接报错退出，提升了启动时的错误报告能力。
    - 链接：https://github.com/zeroclaw-labs/zeroclaw/pull/8115

### 功能请求与路线图信号

今日的功能请求显示出社区对 **新集成**、**协议演进** 和 **任务编排扩展** 的强烈兴趣。

- **协议层重构（路线图信号）**：RFC `#8396` 提出的“Wire-Protocol-First”提供者模型，若被接受，将是 ZeroClaw 架构的一次重大升级，可能影响未来的所有集成方式。
- **新平台/协议集成**：
    - **`[Feature] 实现 Telegram Bot API 10.1 富文本消息 (#8415)`**：提升 Telegram 平台上的用户体验。
    - **`[Feature] 添加 channel-filesystem SOP 事件源 (#8413)`**：允许 SOP（Standard Operating Procedure）工作流由文件系统事件（如文件创建、修改）触发，极大拓展了其自动化场景。
        - 链接：https://github.com/zeroclaw-labs/zeroclaw/issues/8413
- **核心能力增强**：`PR #8400` 成功将 Cron 触发器集成到 SOP 维护循环中，使定时任务 SOP 具备了生产环境可用的调用者，这是对 SOP 编排能力的关键补齐。
    - 链接：https://github.com/zeroclaw-labs/zeroclaw/pull/8400

### 用户反馈摘要

- **用户痛点**：用户 `Audacity88` 反馈的 CLI 秘钥输入无反馈问题 (#7808) 暴露了配置流程中的可用性盲区，已得到修复。
- **使用场景**：用户 `klonuo` 的请求 (#8415) 和 用户 `singlerider` 的请求 (#8413) 反映了社区在使用 ZeroClaw 时，正将其应用于更复杂的自动化和交互场景（如机器人富文本聊天和文件驱动的自动化工作流）。用户 `singlerider` 对将 SOP 集成到文件系统事件中的需求，显示其已达到一定的使用深度。

### 待处理积压

以下是一些重要的待处理工作项，它们年龄较长或涉及高风险变更，需要维护者关注。

1.  **高风险、长期待审查的 PR**：
    - `#7162` - **通知压缩功能**：该 PR 创建于6月3日，旨在在压缩上下文前通知用户，属于用户体验优化。至今已25天未合并。
    - `#7529` - **网关 URL 打印 Bug**：创建于6月12日，修复了一个误导性打印问题。至今已16天。
    - `#7923` & `#8368` - **大规模新功能**：这两项分别是“自动清理临时文件”和“替换 Wasm 运行时”，均为高风险（`risk:high`）、大工作量（`size:XL`）的 PR，虽在积极更新，但合并周期可能较长，需持续关注。

2.  **需要作者操作的 PR**：
    - `#8350` - **Web 搜索性能优化 PR**：此 PR 被标记为 `needs-author-action`，可能需要作者回应 review 意见或更新代码。
        - 链接：https://github.com/zeroclaw-labs/zeroclaw/pull/8350

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*