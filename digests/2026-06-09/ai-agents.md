# OpenClaw 生态日报 2026-06-09

> Issues: 147 | PRs: 475 | 覆盖项目: 13 个 | 生成时间: 2026-06-09 04:18 UTC

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

# OpenClaw 项目动态日报 | 2026-06-09

---

## 1. 今日速览

过去24小时项目保持高度活跃：共处理 **147 条 Issue**（新开/活跃 83 条，关闭 64 条）及 **475 条 PR**（待合并 314 条，已合并/关闭 161 条），并发布 **2 个 beta 版本**。社区反馈集中在 **升级后兼容性问题**（OpenAI/ChatGPT、DeepSeek、Telegram 回归）与 **会话状态异常**（重复回复、内存泄露、SQLite 向量索引）上，维护者已快速响应多个 P1 级 Bug。同时，**Windows 原生测试覆盖**与 **ACP 错误类型保留**等基础设施改进持续推进，项目整体健康度中等偏上，但 Beta 版本稳定性仍有待观察。

---

## 2. 版本发布

### v2026.6.5-beta.5 / v2026.6.5-beta.3

两个 Beta 版本于今日发布，核心变更一致：

- **QQBot 消息过滤增强**：自动剥离模型推理/思考框架标记（`<thinking>`）以防原始内容泄漏到频道回复中（参考 #89913, #90132，感谢 @openperf）。
- **MCP 工具结果规范化**：强制将 `resource_link`、`resource`、`audio`、格式异常图片等类型转换为标准响应格式，并对未来新增类型提供兜底处理。

> ⚠️ **迁移注意**：Beta 版本中包含 **API 行为变更**——MCP 工具结果的结构化字段可能重新排序，依赖原始 JSON 结构的外部插件需要适配。建议升级前备份工作空间，并等待稳定版（预计 2026.6.5 正式版）。

---

## 3. 项目进展

今日共 **合并/关闭 161 个 PR**，重点修复如下：

### 稳定性与兼容性
| PR | 标题 | 状态 | 关键影响 |
|---|---|---|---|
| [#89034](https://github.com/openclaw/openclaw/pull/89034) | fix(usage): return all agent sessions when no agentId specified | **已合并** | 恢复 v2026.5.17 前多 Agent 配置下 Web UI 会话过滤的正确行为 |
| [#90369](https://github.com/openclaw/openclaw/pull/90369) | test(browser): replace broad win32 skips with dynamic capability checks | **已合并** | 提升浏览器路径测试在 Windows 上的覆盖率 |
| [#90374](https://github.com/openclaw/openclaw/pull/90374) | 类似替换 win32 硬编码跳过 | **已合并** | Telegram、IRC、LINE、WhatsApp 等多通道 Windows 兼容性提升 |
| [#90387](https://github.com/openclaw/openclaw/pull/90387) | 同上 Telegram token 测试 | **已合并** | |
| [#90393](https://github.com/openclaw/openclaw/pull/90393) | 同上 WhatsApp auth 测试 | **已合并** | |

### 性能与延迟
| PR | 标题 | 状态 | 关键影响 |
|---|---|---|---|
| [#86981](https://github.com/openclaw/openclaw/pull/86981) | fix(tui): prewarm agent runtime before first send | **已合并** | 首条消息发送前预热 Agent 运行时，减少首次响应延迟 |
| [#91598](https://github.com/openclaw/openclaw/pull/91598) | perf(control-ui): lazy load slash commands | **已合并** | 优化控制面板启动速度，延迟加载斜杠命令直到实际使用 |

### 诊断与日志
- [#91600](https://github.com/openclaw/openclaw/pull/91600) 已打开（待审核）：修复 App-Agent/OAuth 会话的诊断日志上下文问题。

> 项目整体向前迈进了 **约 50 个功能点/修复**（按合并 PR 估算），尤其 Windows 平台测试基础设施大幅增强，为后续跨平台稳定性打下基础。

---

## 4. 社区热点

### 🔥 最活跃 Issue
| Issue | 标题 | 评论数 | 点赞 | 类别 |
|---|---|---|---|---|
| [#90083](https://github.com/openclaw/openclaw/issues/90083) | OpenAI ChatGPT Responses transport fails with `invalid_provider_content_type` for gpt-5.4/5.5 | **15** | 3 | **升级兼容性** |
| [#32296](https://github.com/openclaw/openclaw/issues/32296) | Agent replies to previous message instead of current (session context confusion) | **14** | 1 | **会话状态** |
| [#85888](https://github.com/openclaw/openclaw/issues/85888) | Cron jobs consistently fail with MiniMax 503 overload during early morning | **10** | 1 | **定时任务可靠性** |

**分析**：
- #90083 表明自 v2026.6.1 升级后，OpenAI 最新模型（gpt-5.4/5.5）在 OpenClaw 中完全不可用，用户强烈要求紧急修复（P1 级，且标记为 `platinum hermit`）。目前尚无关联 PR。
- #32296 持续对会话上下文混淆表示不满，该问题自 3 月报告至今已超 3 个月仍未关闭，用户对修复速度有期待。
- #85888 暴露了 MiniMax API 在特定时段（05:00-07:30 CST）的过载问题与 OpenClaw 调度策略的交互缺陷——手动触发可成功，但定时任务失败，说明调度器重试/退避逻辑存在缺陷。

### 🔥 最受关注 PR
| PR | 标题 | 状态 | 点赞 |
|---|---|---|---|
| [#91302](https://github.com/openclaw/openclaw/issues/91302) | (已关闭) claude-cli tool-call 后空文本导致重复回复 | 已修复 | 2 |
| [#91150](https://github.com/openclaw/openclaw/issues/91150) | /status 误导 Context 计算 | 已修复 | 2 |

这些关注点反映了用户对 **消息不透传**、**状态显示不准确** 等体验类问题的敏感度。

---

## 5. Bug 与稳定性

### 🔴 P1 级 Bug（当日活跃）

| ID | 标题 | 是否有 Fix PR | 影响域 |
|---|---|---|---|
| #90083 | OpenAI ChatGPT Responses 传输失败 (gpt-5.4/5.5) | ❌ 无 | 核心推理 |
| #91302 | claude-cli 空文本导致重复回复 (Telegram) | ✅ 已修复 | 消息传递 |
| #90456 | OpenAI OAuth Realtime Talk 返回 HTTP 500 | ❌ 无 | 语音功能 |
| #91191 | WhatsApp 监控静默丢弃某个号码的消息 | ❌ 无 | 通道接收 |
| #87109 | macOS 上 Gateway 堆内存增长至 1GB+ 导致定时任务静默失败 | ❌ 无 | 稳定性 |
| #88657 | DeepSeek V4 Flash 不完整轮次 (payloads=0) | ❌ 无 | 推理 |
| #91434 | LM Studio 会话启动时工具集为空 | ❌ 无 | 推理 |
| #91433 | MCP 远程 OAuth 连接失败 (两个独立 Bug) | ❌ 无 | 外部集成 |

### 🟡 P2 级回归/行为 Bug（当日修复）

- **#89525** Telegram 压缩命令 `/compact` 失效（已关闭，修复待合并）
- **#88933** 压缩操作在异步回调中 `notifyUser` 被阻塞（已关闭，修复待合并）
- **#88128** Telegram 内部消息泄漏到用户聊天（已关闭，修复待合并）
- **#90891** Doctor 未检测官方插件版本漂移（已关闭，修复待合并）
- **#90007** 升级后插件索引迁移失败 `TypeError: statement.columns is not a function`（已关闭，修复待合并）

### 趋势总结
今日 **回归类 Bug 数量下降**（昨日 10+ 个 → 今日 5 个），表明 Beta 版本的热修补正在生效。但 **P1 级新 Bug**（#91434、#91433）同时暴露了 LM Studio 与 MCP OAuth 的深度集成问题，需警惕 Beta.5 后的集成兼容性。

---

## 6. 功能请求与路线图信号

### 可能进入下一版本的请求

| Issue/PR | 标题 | 状态 | 信号强度 |
|---|---|---|---|
| [#44294](https://github.com/openclaw/openclaw/issues/44294) | 保留结构化 ACP 后端错误类型（不映射到 `end_turn`） | 开放，已有 PR #91603 | **高**（PR 已就绪） |
| [#44297](https://github.com/openclaw/openclaw/issues/44297) | 将 Slack 外部 arg-menu 回退暴露为健康信号 | 开放 | 中（已有清晰设计） |
| [#79223](https://github.com/openclaw/openclaw/issues/79223) | 可配置 Dream Diary 语言（非英语用户） | 开放（`stale` 标签） | **低**（长期未响应） |
| [#66705](https://github.com/openclaw/openclaw/issues/66705) | 在 exec 子进程中暴露 `OPENCLAW_SESSION_KEY` 环境变量 | 开放（有 Open PR） | **高**（安全审查中） |
| [#78441](https://github.com/openclaw/openclaw/pull/78441) | 子 Agent 转发 `toolsAllow` 策略 | 开放（待审核） | **高**（功能验证已通过） |

> 特别关注 **#91603 PR**（ACP 错误类型保留），一旦合并将显著改善开发者对 LLM 错误原因的定位能力。**#66705** 涉及安全边界，需优先完成安全审查。

---

## 7. 用户反馈摘要

### 正向反馈
- **@losts1** 在 #89434 中感谢 WebRTC Talk 修复“peer is null”的迅速响应。
- **@wilfried-codex** 在 #91482 中赞扬 `dir_list` 暴露问题已被快速识别。

### 痛点与不满
- **升级体验**：多个用户（@TreyLawrence, @ramitrkar-hash, @laurenceputra）抱怨从 v2026.5.x 升级到 v2026.6.x 后出现 `statement.columns is not a function`、SQLite 插件冲突、构建失败等问题，提示升级路径测试仍不充分。
- **会话状态混乱**：@survivor998（#32296）和 @mikefaierberg-byte（#88657）重复报告“Agent 回复上一条消息”和“不完整轮次”，严重影响日常使用。
- **定时任务**：@xxtyyq（#85888）指出 MiniMax 过载问题的手动/定时行为不一致，判断为调度器缺陷，而非单纯 API 问题。
- **Windows 支持**：@jackmtl71（#90157）反映“Open config”按钮在 PowerShell 中抛出异常，这一体验问题在 Beta 版本中仍未解决。

### 使用场景亮点
- **Dreaming 功能**：@stevenjj33（#91051）报告 Deep Sleep 阶段未写入 DREAMS.md，表明有用户正在深入使用记忆自动摘要功能。
- **微信集成**：@zaakwang-lam（#78754）仍在依赖 `openclaw-weixin` 通道，并遇到 cron 任务兼容问题，说明该通道仍有实际用户。

---

## 8. 待处理积压

### ⏳ 长期未响应的 Issue（stale 标签，>30 天无活动）

| Issue | 标题 | 创建时间 | 标签 | 优先级 |
|---|---|---|---|---|
| [#44289](https://github.com/openclaw/openclaw/issues/44289) | Generate secretref reference docs from secret target registry metadata | 2026-03-12 | stale, P2 | **中**（已有 PR #89142 待审核） |
| [#44291](https://github.com/openclaw/openclaw/issues/44291) | Add native PowerShell smoke coverage for contributor commands | 2026-03-12 | stale, P2 | **中**（Windows 生态缺失） |
| [#44292](https://github.com/openclaw/openclaw/issues/44292) | Add autofix for missing config field labels | 2026-03-12 | stale, P3 | **低**（已有 CI 检测兜底） |
| [#66705](https://github.com/openclaw/openclaw/issues/66705) | Expose session key env vars in exec child processes | 2026-04-14 | stale, P2, 安全 | **高**（有 Open PR，安全审查停滞） |
| [#79223](https://github.com/openclaw/openclaw/issues/79223) | Configurable Dream Diary language / prompt | 2026-05-08 | stale | **低**（需求明确但无 PR） |

### ⏳ 长期未合并的 PR

| PR | 标题 | 创建时间 | 状态 | 阻塞原因 |
|---|---|---|---|---|
| [#41299](https://github.com/openclaw/openclaw/pull/41299) | fix: add clear separator between metadata and user message | 2026-03-09 | 开放，stale | 需要真实行为验证 |
| [#78441](https://github.com/openclaw/openclaw/pull/78441) | feat(subagents): forward toolsAllow | 2026-05-06 | 开放，等待审核 | 安全边界审查 |
| [#83169](https://github.com/openclaw/openclaw/pull/83169) | Discord reaction notification wake policy | 2026-05-17 | 开放 | 需要产品决策 |

> **维护者提醒**：积压中最值得优先处理的是 **#66705**（exec 环境变量）和 **#41299**（消息分隔符），两者均有社区 PR 且分别涉及安全和体验，长期搁置可能影响贡献者积极性。此外，**#44289** 的 PR #89142 已准备就绪，仅需最终审核。

---

## 总结

OpenClaw 在 2026-06-09 展现出高强度的社区纠错能力，尤其是对 **回归问题** 的快速响应（如 Telegram / DeepSeek 修复）。但 **P1 级新 Bug 不减反增**（#91434、#91433），提示 Beta.5 版本可能存在未发现的集成裂缝。推荐用户耐心等待稳定版 v2026.6.5 正式发布后再升级，生产环境建议锁定在 v2026.5.x 分支。

*报告生成时间：2026-06-09 UTC，数据来源：GitHub API（openclaw/openclaw）。*

---

## 横向生态对比

好的，作为一名专注于AI智能体与个人AI助手开源生态的资深技术分析师，以下是根据您提供的2026-06-09各项目日报生成的横向对比分析报告。

---

### AI智能体/个人AI助手开源生态横向对比分析报告 (2026-06-09)

**报告时间：** 2026-06-09
**数据来源：** 各项目GitHub仓库当日动态

#### 1. 生态全景

截至2026年6月9日，个人AI助手/自主智能体开源生态呈现出**高度活跃、快速分化、安全与稳定成为普遍痛点**的态势。头部项目如OpenClaw、IronClaw、Hermes Agent等正围绕“稳定性修复”和“下一代架构（Reborn、ZeroClaw）”进行密集迭代，社区贡献量巨大。同时，一个明确的趋势是，各项目在追求功能完整性的同时，开始将焦点转向**生产级特性**，如跨Agent协作、审批工作流、出口封锁、安全审计等，标志着该生态正从“技术演示”阶段向“可用、可信”的实用阶段迈进。然而，广泛的升级兼容性问题（OpenAI、DeepSeek）和会话状态混乱表明，快速迭代也带来了稳定性挑战。

#### 2. 各项目活跃度对比

| 项目 | 活跃 Issues | 活跃 PRs | 版本发布 | 健康度评估 | 核心方向 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 83 | 475 | 2 Beta | 中等偏上 | **稳定性修复/回归测试** |
| **NanoBot** | 3 | 34 | 无 | 良好 | **多提供者/安全性/消息处理** |
| **Hermes Agent** | 4 | 50 | 无 | 高 | **桌面端/安全/AI交互精细化** |
| **PicoClaw** | 2 | 20 | 1 Nightly | 良好 | **代码质量/跨平台(Telegram/DeltaChat)** |
| **NanoClaw** | 1 | 4 | 无 | 良好 | **安全加固(出口封锁)** |
| **NullClaw** | - | - | - | 停滞 | - |
| **IronClaw** | 11 | 50 | 无 | **极高** | **Reborn架构/工作流集成/审批门控** |
| **LobsterAI** | 1 | 17 | 无 | **高** | **功能交付/协作/数据迁移** |
| **TinyClaw** | 0 | 1 | 无 | 低 | **安装体验优化** |
| **Moltis** | - | - | - | 停滞 | - |
| **CoPaw** | 8 | 47 | 无 | 高 | **稳定性修复(OneBot)/插件生态/TUI** |
| **ZeptoClaw** | - | - | - | 停滞 | - |
| **ZeroClaw** | 17 | 50 | 无 | **极高** | **v0.8冲刺/WASM/安全/配置精炼** |

#### 3. OpenClaw 在生态中的定位

- **优势与定位**：OpenClaw是目前生态中**用户基础最大、社区讨论最活跃、Issue/PR量级最高**的“核心参照”项目。其定位是功能全面的“个人AI助手操作系统”，整合了最广泛的LLM提供商、通信渠道和工具链。
- **技术路线差异**：相比Hermes Agent的“智能体大脑”特性（如学习循环）和IronClaw的“企业级工作流”架构，OpenClaw更侧重于**“桥梁”角色**，连接前端（多种聊天/UI通道）与后端（LLM、MCP、工具）。其优势在于生态的广度（Telegram, QQ, WhatsApp, LINE等全渠道覆盖），但这也带来了更复杂的兼容性维护挑战。
- **社区规模与挑战**：尽管活跃度压倒性领先，但其“中等偏上”的健康度评分表明，规模也为OpenClaw带来了更严重的**“稳定性债”**。大量P1级Bug（如OpenAI gpt-5.4兼容性）和回归问题凸显了其快速迭代与用户基数庞大之间的矛盾。相比之下，IronClaw和ZeroClaw虽然活跃度稍低，但其“极高”的健康度说明其迭代更聚焦、更有控制。

#### 4. 共同关注的技术方向

- **平台兼容性与稳定性（所有活跃项目）**：这是一个普遍痛点。
    - **OpenAI/DeepSeek兼容性**：OpenClaw (gpt-5.4/5.5)、PicoClaw (RISC-V OpenAI) 均报告了与最新LLM提供商的兼容问题。IronClaw 也有 Minimax 配置问题。
    - **Windows原生支持**：OpenClaw、PicoClaw (QQ频道)、CoPaw (Shell闪窗)、ZeroClaw (跨平台代码质量) 均有涉及，表明跨平台稳定性是行业级挑战。
- **跨Agent协作与工作流**：这是从“单兵作战”到“团队协同”的关键。
    - **NanoBot** (#3992): 实现跨Agent消息总线和实例协作。
    - **Hermes Agent** (PR #42551): 核心工具渐进式披露，优化多工具场景。
    - **IronClaw** (Reborn架构): 审批门控、子代理(Subagent)规划，强调流程控制。
- **配置精细化与安全控制**：用户对“有效控制”的需求前所未有地强烈。
    - **OpenClaw** (工具过滤/ACL)、**Hermes Agent** (Provider选择器)、**NanoClaw** (出口封锁)`、**ZeroClaw** (MCP工具过滤器/私有域名过滤) 均致力于让用户精确配置Agent的行为边界和安全策略。
- **多模态输入与处理**：从文本向富媒体交互进化。
    - **NanoBot** (#4217, #4224, #4232): 支持语音转写(AssemblyAI, 小米)、共享语音输入、文件上传解析。
    - **Hermes Agent** (PR #42507): iMessage 全双向媒体交互。

#### 5. 差异化定位分析

| 项目 | 功能侧重 | 目标用户 | 技术架构/语言 | 核心差异点 |
| :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 全能型渠道管理 | 开发者/高级用户 (核心参照) | 模块化、TypeScript | **渠道最广**，问题最多，社区最成熟，是了解生态的入口。 |
| **NanoBot** | 技术极简、安全 | 重视隐私和简单部署的用户 | 轻量级、Go语言 | **代码轻量、部署简单、安全响应快**，是“即插即用”的代表。 |
| **Hermes Agent** | 桌面端深度体验 | 桌面重度用户、追求AI交互质量 | Electron/React等 | **桌面端功能最丰富**（桌面Onboarding、气泡布局），AI交互控制精细（角色扮演语音、学习循环概念）。 |
| **PicoClaw** | 跨平台嵌入式 | 嵌入式/IoT开发者 | Go语言、跨平台 | **轻量级“下载即用”**，注重RISC-V/Windows嵌入式场景，与OpenClaw互补。 |
| **IronClaw** | 企业级工作流 | 企业/专业用户、需要审批流程 | 强调架构 (Reborn) | **架构最大胆、最现代**，聚焦审批、工作流、生产级可用性，是“企业版”的领跑者。 |
| **LobsterAI** | 协作与数据管理 | 团队用户、注重数据备份 | 桌面客户端 | **功能交付效率高**，专注数据迁移、协作安全、定时任务等，是“协作助手”的典范。 |
| **ZeroClaw** | 深度开发者定制 | 追求极致灵活性和升级体验的用户 | 高度模块化、Rust/Go等 | **版本升级体验和插件生态 (WASM) 是其核心亮点**，代码质量和配置体验要求高。 |
| **CoPaw** | 国内生态整合 | 主要面向中国用户 | - | **深度整合钉钉、元宝等国产渠道**，并积极构建TUI和插件市场，是“中国本土化”的代表。 |

#### 6. 社区热度与成熟度

- **第一梯队：快速迭代 / 功能扩展期** - **Hermes Agent、IronClaw、ZeroClaw、LobsterAI、CoPaw**。这些项目合并/关闭PR的效率高，或者正围绕核心架构（Reborn、ZeroClaw v0.8）进行大规模开发，社区讨论活跃，开发者贡献热情高。它们代表了生态的未来方向，但稳定性可能处于新的“阵痛期”。
- **第二梯队：质量巩固 / 稳定优化期** - **OpenClaw、NanoBot、PicoClaw、NanoClaw**。这些项目具备坚实的功能基础，当前重心正转向修复高优先级Bug、提升安全性（出口封锁）、优化跨平台支持和代码质量。它们是生态的“压舱石”，稳定性更高但功能创新节奏稍缓。
- **第三梯队：沉寂 / 低活跃度** - **NullClaw、Moltis、ZeptoClaw、TinyClaw**。这些项目在过去24小时内无实质性活动或其进展极其微小，无法反映其真实健康度。TinyClaw仅有一个修复PR，处于“被关注”但“低活跃”状态。

#### 7. 值得关注的趋势信号

1.  **Agent自主进化能力成为新焦点**：
    - **案例**：CoPaw (Issue #5017) 明确提出借鉴Hermes Agent的“学习循环”（Learning Loop），让Agent能自动从行为中创建和迭代技能。
    - **价值**：这标志着社区对AI Agent的期待已超越“执行指令”的工具层面，转向“自我改进”的智能体层面。对于开发者而言，未来的Agent框架可能不再是静态配置，而是具备元学习能力的动态系统。

2.  **跨Agent通信与协作从概念走向原型**：
    - **案例**：NanoBot (#3992) 实现了跨Agent消息总线和协作功能。
    - **价值**：这打破了单一Agent的能力边界，预示着未来的Agent应用将像微服务一样，由多个专门化的Agent（搜索Agent、规划Agent、执行Agent）协同工作。这为构建复杂、模块化的AI系统提供了新的思路。

3.  **安全与权限控制从“可有可无”变为“核心刚需”**：
    - **案例**：NanoClaw (#2713, 出口封锁)、Hermes Agent (#40663, Shell注入检测)、IronClaw (#3957, 第三方钩子安全审计)、ZeroClaw (#7412, 域名过滤Bug)。
    - **价值**：随着Agent获得读取文件、执行Shell命令、调用API等更强大的能力，安全地从外部世界隔离Agent已成为所有项目的共识。对于企业级应用来说，细粒度的安全策略（如出口封锁、审批门控）是部署前的必要条件。

4.  **从“聊天工具”演进为“生产力平台”**：
    - **案例**：IronClaw (审批工作流)、LobsterAI (协作任务通知)、Hermes Agent (Discord会议模式)。
    - **价值**：AI Agent正在从简单的对话机器人，演变为集成任务审批、团队协作通知、无干扰会议的“数字化工人”。这表明其应用场景正扩展到更复杂的商业流程，开发者在设计时应将Agent视为工作流中的一环，而不仅仅是一个接口。

5.  **社区治理的“规模病”显现**：
    - **案例**：OpenClaw庞大的Issue/PR积压和回归Bug频发。
    - **价值**：当一个开源项目达到OpenClaw的体量后，其发展节奏将从“野蛮生长”转向“精耕细作”。快速合并PR带来活跃的同时，也带来了测试不充分和稳定性风险。这提醒其他项目，在扩张期需要投入更多资源在自动化测试和回归保护上，以避免“大而不稳”。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot 项目动态日报 | 2026-06-09

---

## 1. 今日速览

NanoBot 项目在 2026-06-09 延续高活跃度：过去 24 小时共处理 34 个 Pull Request（其中 15 个已合并/关闭），以及 7 个 Issue（关闭 4 个，新开/活跃 3 个）。核心贡献集中在**多提供者支持**（Azure 兼容额外查询参数）、**安全性修复**（SSRF 验证、符号链接逃逸阻断）、**消息分割优化**（代码块识别）以及**跨 Agent 协作**等方向。社区提交的多个功能请求（如每对话模型切换、WebUI 版本显示）正在被积极讨论，项目整体健康度良好。

---

## 2. 版本发布

过去 24 小时内无新版本发布。

---

## 3. 项目进展

今日合并/关闭了多项关键 PR，推动了以下主要功能与修复：

| PR # | 标题 | 关键进展 |
|------|------|----------|
| #4217 | feat(providers): add `extra_query` config for OpenAI-compatible providers | 支持 Azure 等网关所需的 `?api-version=` 查询参数，解决了部分端点 404 的问题 |
| #4224 | feat(transcription): add AssemblyAI as transcription provider | 增加 AssemblyAI 语音转写提供商，丰富 STT 选择 |
| #4175 | feat(transcription): add Xiaomi MiMo ASR provider | 集成小米 MiMo ASR 模型，增强中文语音识别 |
| #4113 | feat(transcription): configurable STT model + OpenRouter provider | 使 OpenRouter 可同时用于文本与语音转写 |
| #4232 | feat(transcription): add shared voice input support | 将语音输入提升为全局能力，支持 WebUI 和桌面端 |
| #4219 | fix(session): drop orphan tool results before trimming history | 修复会话历史裁剪时残留孤立工具结果的问题 |
| #4221 | fix(exec): block relative symlink workspace escapes | 阻断通过相对路径符号链接逃逸工作目录的安全漏洞 |
| #4232 | fix(weixin): reload session state after pause expiry | 修复微信频道 Token 过期后永久静默的死循环 |
| #4257 | fix(utils): make split_message fenced-code-block-aware | 修复长消息切分时破坏代码块语法的问题 |
| #3992 | feat(agent-collab): enable cross agent messaging | 实现多 Agent 实例间的消息总线，支持跨实例协作 |

此外，多个测试基础设施 PR（#3982、#3983、#4193）以及内存游标单调性修复（#4256）也处于开放并持续改进状态，项目在**功能扩展、稳定性、安全性**三个维度同步推进。

---

## 4. 社区热点

以下 Issue/PR 在报告期内获得较多讨论或社区关注：

- **[#4253] support overriding model per conversation**  
  用户希望能在不同对话中使用不同的模型（如公开的 OpenRouter 与私有的本地 LlamaCpp），以平衡速度与隐私需求。该需求属于**用户场景驱动的功能请求**，已有 1 条评论，暂无对应 PR。

- **[#4233] Show the nanobot version in the webui somewhere**  
  用户希望 WebUI 显式显示版本号，并最好能提示是否有更新。这是一个**易用性改进**，已标记为 `good first issue`，社区参与门槛低。

- **[#4251] 支持在输入框上传文件或图片，根据输入处理答案**  
  中文用户请求直接上传图片/PDF 并由 AI 解析。该需求属于**多模态输入支持**，与现有 transcription 能力互补，未来可能纳入路线图。

---

## 5. Bug 与稳定性

按严重程度排列：

| 严重程度 | Issue/PR | 描述 | 状态 |
|----------|----------|------|------|
| **高** | #4074 | MCP HTTP/SSE 配置在拒绝前尝试回环连接，存在 SSRF 风险 | **已关闭**（涉及 MCP 连接逻辑修复） |
| **高** | #4221 | 符号链接相对路径可逃逸工作目录，导致任意文件读取 | **已关闭**（PR #4221 已合并） |
| **中** | #4250 / #4257 | `split_message` 在代码块中间切分导致语法破坏 | **#4250 开放**，**PR #4257 已提交**修复 |
| **中** | #4223 | 微信频道 Token 过期后因未重载状态而永久静默 | **PR #4223 已提交**（未合并） |
| **中** | #4219 | 会话历史裁剪时孤立工具结果丢失 | **已关闭**（PR #4219 已合并） |
| **低** | #4256 | MemoryStore 游标在压缩后可能非单调，影响写入 | **PR #4256 已提交**（开放中） |

项目安全响应及时，#4074 和 #4221 等关键漏洞均已在修复后关闭。

---

## 6. 功能请求与路线图信号

结合新增 Issue 与已合并 PR，以下功能请求可能进入下一版本：

- **每对话模型覆盖（#4253）**：用户需求强烈，且已有全局配置基础，扩展难度适中。
- **文件/图片上传与解析（#4251）**：与刚完成的共享语音输入（#4232）逻辑类似，可复用多模态管道。
- **WebUI 版本显示（#4233）**：已标记 `good first issue`，预计短期内会有社区贡献。
- **Azure 兼容额外查询参数（#4204 → #4217）**：**已合并**，将随下一发布版本提供。
- **AssemblyAI / 小米 ASR / OpenRouter 转录（多个 PR）**：转录提供商生态快速扩充，路线图明确。

此外，跨 Agent 协作（#3992）作为大型特性已合入，说明项目正从单实例向多智能体协作演进。

---

## 7. 用户反馈摘要

从 Issue 评论中提炼出真实用户痛点与场景：

- **模型切换需求**（#4253）：用户日常使用 OpenRouter（快速）和本地 LlamaCpp（隐私/慢）两种预设，希望按任务自由切换，而不是全局固定。
- **版本信息可见性**（#4233）：用户知道 `/status` 端点可查版本，但希望 WebUI 界面直接显示，并提示更新以降低维护成本。
- **多模态上传**（#4251）：用户希望上传图片让 AI 分析含义，或上传 PDF 让 AI 总结，此需求在中文用户群中尤为突出。
- **Azure 网关的兼容性**（#4204）：用户部署在 Azure 后遭遇 404，由于缺少 `api-version` 参数，该问题直接阻碍使用，感谢社区快速修复。
- **微信频道稳定性**（#4223）：用户反映 Token 过期后频道永久静默，需要手动重新扫码，影响无人值守场景。

整体来看，用户对项目扩展性（更多提供商、多模态）和终端易用性（版本显示、会话级配置）有较高期待。

---

## 8. 待处理积压

以下为长期未合并或未回复的重要 PR/Issue，建议维护者关注：

| 项目 | 标题 | 创建时间 | 说明 |
|------|------|----------|------|
| PR #3982 | test: add scripted agent runner harness | 2026-05-24 | 核心测试基础设施，覆盖工具循环场景，持续未合并 |
| PR #3983 | test: cover runner blocked tool-call finish reasons | 2026-05-24 | 对应同上，增加拒绝、内容过滤等 finish reason 覆盖 |
| PR #4053 | fix(tools): keep read-only roots out of write paths | 2026-05-29 | 安全加固：防止只读路径被意外写入 |
| PR #4119 | fix(exec): block relative symlink workspace escapes | 2026-05-31 | 与 #4221 类似，但基于不同实现路径，未关闭（可能重复） |
| PR #4177 | docs: make onboarding friendlier for beginners | 2026-06-03 | 文档改进，涉及 README、快速入门、配置示例等 |
| PR #4170 | feat(email): add configurable IMAP post-actions | 2026-06-03 | 邮件通道增强，允许处理后的消息自动归档/删除 |
| Issue #4233 | Show the nanobot version in the webui | 2026-06-07 | 易用性改进，无相关 PR，适合社区贡献 |

以上项目多数停留时间较长（超过 7 天），建议安排 reviewer 进行最终审查或合并，以避免社区贡献热情消退。

---

*日报生成时间：2026-06-09 10:30 UTC*  
*数据来源：[HKUDS/nanobot GitHub](https://github.com/HKUDS/nanobot)*

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是根据你提供的 Hermes Agent 项目数据生成的 2026-06-09 项目动态日报。

---

# Hermes Agent 项目动态日报 | 2026-06-09

## 1. 今日速览

Hermes Agent 项目今日维持超高活跃度，开发与社区讨论均十分密集。昨日新增了 50 条 Pull Request，虽然其中 45 条仍处于待合并状态，但这反映了社区贡献与团队开发工作的井喷式增长。同时，Issue 更新数量也达到 4 条，其中 3 条为新提交或活跃讨论。项目整体处于快速迭代期，尤其在安全修复、桌面端功能增强、工具链优化和 AI 行为精细化控制方面有大量工作推进。尽管合并速度（5条）稍显滞后于提交速度，但项目健康度依然很高，代码库正在经历显著的扩张与优化。

## 2. 版本发布

无

## 3. 项目进展

昨日共合并/关闭了 5 条 PR 和 1 个 Issue，项目在这些领域取得了明确进展：

- **Slack 通知样式优化**：Issue #42601 已被关闭，其对应的 PR 实现了使用 Slack Block Kit 发送结构化通知，取代了传统的纯文本回复，将普通对话和重要通知区分开，提升了 Slack 集成体验。
- **大量质量与安全修复**：多个 PR 正在处理关键 Bug 和安全问题，如 #40663 检测 shell 命令混淆攻击，和 #42547 防止 ZAI 提供商钱包资金静默流失，这些皆属于高风险、高价值的修复，一旦合并将大幅提升项目健壮性。
- **桌面端功能丰富化**：桌面端新增了远程 Onboarding 路径 (#42282)、气泡聊天布局 (#41402) 和可调整大小的主题终端面板 (#42521)，表明项目在优化独立桌面应用的用户体验和功能完整性方面投入了显著精力。
- **AI 交互模式优化**：PR #42596 解决了 tool call 后模型仅回复“进度消息”而未实际处理结果的常见问题，通过引入进度占位检测和重新提示机制，有望显著提升对话智能体的可靠性。

## 4. 社区热点

尽管 PR 数量庞大，但大部分公开讨论集中在较老的 Issue 上。今日最能体现社区诉求的热点如下：

- **热门 Issue：Provider 选择器配置** (Issue #12655)
  - **链接**: [NousResearch/hermes-agent Issue #12655](https://github.com/NousResearch/hermes-agent/issues/12655)
  - **热度**: 6条评论，虽创建较早却持续活跃，说明该需求存在广泛的长期用户群体。
  - **诉求分析**: 用户希望能在 `/model` 命令的显示结果中，通过 `picker_providers` 配置项隐藏不需要的提供商（如 Anthropic, OpenRouter）。核心诉求是**用户自主权和体验精细化**，特别是针对那些仅使用私有/自定义端点的用户，他们希望界面更清爽，避免无关选项干扰。

## 5. Bug 与稳定性

昨日社区报告了 3 个新的 Bug，按严重程度排列如下，其中部分已有对应的修复 PR：

- **严重: ZAI 提供商钱包资金静默流失 (Issue #42536 → PR #42547)**
  - **描述**: 若用户使用 Coding Plan 密钥，请求可能被错误路由到按量计费端点，导致意外扣费。这是一个**资金安全**问题。
  - **状态**: 已有修复 PR #42547，通过调整端点探测顺序来规避。

- **中: Web 搜索工具在禁用后仍可通过 Browser 工具暴露 (Issue #42585)**
  - **描述**: `web_search` 同时属于 `web` 和 `browser` 两个工具集。当用户禁用 `web` 工具集时，该功能仍可通过 `browser` 工具集使用，导致 `hermes tools disable web` 命令具有误导性，带来安全策略失效风险。
  - **状态**: 新提交的 Bug，暂无对应的修复 PR。

- **低/UI: 文件浏览器悬停展开功能遮挡滚动条 (Issue #42593)**
  - **描述**: 当文件浏览器折叠时，鼠标移到屏幕右侧边缘意图操作聊天区滚动条，会误触文件浏览器的悬停展开区域，导致滚动条被遮盖。
  - **状态**: 新提交的用户体验 Bug，暂无修复 PR。

## 6. 功能请求与路线图信号

社区和开发者在积极推动项目向更强大、更灵活的方向发展，以下功能请求具有很强的信号意义，且部分已有对应的 PR 实现：

- **AI 交互精细化控制**：
    - **核心工具渐进式披露 (PR #42551)**：允许用户配置哪些核心工具在搜索时优先显示或延迟披露。这适合大型工具集场景，有助于提升 LLM 的工具选择效率和准确性，很可能被纳入下一版本。
    - **Gemini TTS 角色扮演提示 (PR #42595)**：允许为 Gemini 语音合成添加独立的人物角色提示文件，实现更具表现力的语音输出。这直接回应了社区对于 AI 个性和交互深度的需求。

- **桌面端功能扩展**：
    - **本地知识库集成 (PR #42598)**：新增 “Graphify“ 本地笔记面板，支持读写本地 Graphify 图数据库。这表明项目正在从单纯的对话助手向集成本地知识管理的全能型个人助手演进。
    - **仪表盘本地回环认证 (PR #42546)**：简化桌面客户端与本地仪表盘的连接配置，提升易用性。

- **平台集成增强**：
    - **Discord 语音会议模式 (PR #42597)**：优化 Discord 集成，新增仅转录不触发电竞语音回复的“会议模式”，这是向企业级/实用性办公场景的重要一步。
    - **iMessage 双向媒体交互 (PR #42507)**：完成 iMessage 插件的全双向媒体支持（图片、语音、视频、Tapbacks 等），补齐了苹果生态信息集成的重要功能。

## 7. 用户反馈摘要

从今日的 Issue 和 PR 评论中，可以提炼出以下用户反馈：

- **对自主权和控制力的渴望 (Issue #12655)**: 用户明确表示需要更精细的提供商可见性控制，以便管理混乱的界面，反映了用户对工具和界面高度定制化的期望。
- **对安全性的敏锐感知 (Issue #42585)**: 用户对“禁用”功能未能彻底生效表示担忧，这直接关系到系统的安全策略和用户信任。用户不是仅仅报告一个 Bug，而是点明了该 Bug 的**误导性**和潜在的**安全风险**。
- **对实用场景的优化要求 (PR #42596)**: 用户观察并报告了 AI 模型在工具调用后的行为缺陷，即模型回答“好的，我看看”等无实质内容的进度消息。这显示了用户对 AI 交互质量的敏锐观察和对一致、实质性回应的期待。

## 8. 待处理积压

当前核心积压在于 **PR 队列的合并效率**。尽管项目活跃度高是好事，但 45 条待合并的 PR 形成了一定积压。社区贡献者和团队开发者提交了大量高质量的代码，若等待时间过长，可能影响贡献者积极性并拖后项目迭代速度。

- **关键需关注的 PR**: #40663 (Shell命令注入检测) 和 #42547 (ZAI钱包保护) 均为安全修复，应被优先审查和合并。
- **大型功能 PR**: #41871 (Gateway Gateway 适配器) 和 #42282 (桌面端 Onboarding) 属于架构级或核心体验的大型功能，需要充分的代码审查和测试，建议项目维护者投入精力优先推进。

**提醒**: 关注 Issue #42585 和 #42593，它们分别涉及安全策略漏洞和用户体验问题，建议在后续的迭代中尽快给出方案或进行响应。

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw 项目动态日报 | 2026-06-09

## 今日速览

过去 24 小时内，PicoClaw 项目保持了较为活跃的开发节奏。共处理 3 个 Issue（新开 2 个，关闭 1 个），并提交了 20 个 Pull Request（合并/关闭 9 个，待合并 11 个）。团队集中修复了多个长期存在的代码质量隐患（如未检查类型断言、未处理错误返回值），并合并了 Telegram 位置消息忽略问题的修复。同时，一个针对 DeltaChat 网关的新特性 PR 已提交，展示了项目在跨平台拓展上的持续投入。整体项目健康度良好，社区响应及时，但个别遗留 Bug（如 RISC-V 平台兼容性）仍需关注。

## 版本发布

**nightly build：`v0.2.9-nightly.20260609.46b29a0a`**  
- 自动构建版本，基于 `main` 分支最新提交，可能包含不稳定变更。  
- 完整变更日志：[Compare v0.2.9...main](https://github.com/sipeed/picoclaw/compare/v0.2.9...main)  
- 无破坏性变更或迁移说明。

## 项目进展

过去 24 小时合并/关闭的 9 个 PR 涵盖了以下关键改进：

- **Bug 修复**  
  - [#3052](https://github.com/sipeed/picoclaw/pull/3052) 解决了 Telegram 通道忽略位置消息（`message.location`）的问题，将其转换为文本并推送至代理管线。  
  - [#3062](https://github.com/sipeed/picoclaw/pull/3062) 修复了健康检查始终返回“not ready”的 bug。  
  - [#3058](https://github.com/sipeed/picoclaw/pull/3058) 为 `isAllowedFirstHopHost` 添加类型断言检查，防止潜在的空字符串默认行为。  
  - [#3057](https://github.com/sipeed/picoclaw/pull/3057)、[#3056](https://github.com/sipeed/picoclaw/pull/3056) 在 `subagent`、`spawn`、`base` 等多个工具函数中补充了未检查的类型断言。  
  - [#3055](https://github.com/sipeed/picoclaw/pull/3055) 处理 `NewContextBuilder` 中 `os.Getwd` 错误，使用空字符串作为回退。  
  - [#3018](https://github.com/sipeed/picoclaw/pull/3018) 修复 LINE 通道和 Evolution store 中的类型断言与错误处理。

- **代码质量与重构**  
  - [#3051](https://github.com/sipeed/picoclaw/pull/3051) 将 `channels` 和 `mcp` 中的错误包装改用 `%w`，确保 `errors.Is/As` 链正常工作。  
  - [#3050](https://github.com/sipeed/picoclaw/pull/3050) 将多处生产代码中的 `log.Printf`/`fmt.Printf` 替换为结构化日志基础设施，提升可观测性。

以上合并表明项目正在系统性地提升代码健壮性与可维护性，尤其针对错误处理的标准化改进值得关注。

## 社区热点

- **[#2887 – .deb version on RISC-V is not functional with OpenAI model](https://github.com/sipeed/picoclaw/issues/2887)**  
  本日评论数最多（9 条），虽是旧 Issue（5月17日创建），但在过去 24 小时内有更新。用户报告 `0.2.8` 版本的 `.deb` 包在 RISC-V 架构上无法正常使用 OpenAI 模型（gpt-5.4-2026-03-05），且 Go 版本显示为 `1.25.9`。社区讨论集中在对 RISC-V 构建依赖及证书/API 通道的排查上。该问题已标记为 `stale`，但近期活跃度上升，值得维护者优先响应。

- **[#3015 – QQ channel connection failure on Windows](https://github.com/sipeed/picoclaw/issues/3015)**  
  新开仅 2 天，已有 2 条评论。用户反馈 Windows 版 `picoclaw gateway` 运行时无法获取 QQ 频道的 App Access Token，超时失败。其他通道（如 Pico）工作正常。该问题尚未有修复 PR，可能涉及 Windows 网络栈或证书问题，需要进一步追踪。

## Bug 与稳定性

| 严重程度 | Issue / PR | 状态 | 说明 |
|----------|------------|------|------|
| **高** | [#2887](https://github.com/sipeed/picoclaw/issues/2887) RISC-V .deb 无法使用 OpenAI | OPEN / 无 fix PR | 平台兼容性缺陷，影响 RISC-V 用户 |
| **中** | [#3015](https://github.com/sipeed/picoclaw/issues/3015) QQ 频道 Windows 连接失败 | OPEN / 无 fix PR | 首次上报，需环境复现 |
| **低** | [#3049](https://github.com/sipeed/picoclaw/issues/3049) Telegram 忽略位置消息 | **已关闭**，PR [#3052](https://github.com/sipeed/picoclaw/pull/3052) 已合并 | 24 小时内快速修复，响应及时 |

此外，今日合并的多个 PR（如 #3055、#3056、#3057 等）均针对运行时的潜在 panic（未检查类型断言、未处理错误），这些隐患虽未产生用户报告，但修复后能够提升整体稳定性。

## 功能请求与路线图信号

- **DeltaChat 网关** – [#3063](https://github.com/sipeed/picoclaw/pull/3063)（OPEN）  
  该 PR 新增了 DeltaChat 通道支持，属于新特性扩展。目前尚无后续版本规划，但若顺利合并，将为用户提供基于电子邮件协议的聊天集成选项，可能成为下一版本亮点。

- **UI 配置持久化** – [#3067](https://github.com/sipeed/picoclaw/pull/3067)（OPEN）  
  修复运行时会话隔离范围（`DmScope`）设置无法保存的问题。该功能属于配置管理优化，被多次用户提及，预计会跟随下一次正式版发布。

- **Windows 子进程控制台闪烁消除** – [#3061](https://github.com/sipeed/picoclaw/pull/3061)（OPEN）  
  对 `gateway.go` 以外的多个 `exec.Command` 调用添加 `HideWindow`，提升 Windows 用户体验，属于持续的跨平台打磨。

## 用户反馈摘要

- **RISC-V 用户**（[#2887](https://github.com/sipeed/picoclaw/issues/2887) 评论）表示“.deb 包安装后无法连接 OpenAI，希望提供预编译二进制或修复依赖”。社区中已有用户尝试通过源码编译绕过，但过程较为繁琐。
- **Windows QQ 频道用户**（[#3015](https://github.com/sipeed/picoclaw/issues/3015) 评论）：“Pico 通道正常，但 QQ 一直卡在 token 获取；Win 防火墙已关闭，怀疑是证书或 DNS 问题。”该问题暂无官方回应，用户希望尽快给出临时解决方案。
- **Telegram 用户**（[#3049](https://github.com/sipeed/picoclaw/issues/3049) ）：“发送位置后完全无响应，日志无输出。期望支持位置消息解析。”该反馈已在 PR #3052 中实现，将转换为 `[User location: lat=..., lng=...]` 文本。

## 待处理积压

- **[#2887 – RISC-V 兼容性](https://github.com/sipeed/picoclaw/issues/2887)**  
  已存在近 1 个月（标注 `stale`），但近日重新活跃。缺乏明确的修复方案或维护者回应。建议优先分配资源进行平台构建测试，或提供临时 workaround。

- **[#2904 – Agent 循环重载与 panic 清理稳定性](https://github.com/sipeed/picoclaw/pull/2904)**  
  该 PR 创建于 5 月 20 日，至今未合并或关闭。涉及到 `agent` 包的 `ReloadProviderAndConfig` 并发安全与资源泄漏问题，属于核心稳定性提升。维护者应重新评估优先级，避免长期积压导致分支冲突。

---

**报告生成时间**：2026-06-09 23:59 UTC

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw 项目动态日报 | 2026-06-09

---

## 1. 今日速览

过去24小时项目保持中等活跃度：1个新Issue上报了影响实际使用的WhatsApp媒体附件不可达问题，同时有4个Pull Request被提出，其中3个已合并/关闭。安全加固是今日主旋律——“出口封锁”特性（#2713）已合并，四项认证/安全修复（#2714）处于待合并状态，表明项目团队正积极提升容器隔离与访问控制能力。整体项目健康度良好，但新报的Bug需要尽快定位修复。

---

## 2. 版本发布

今日无新版本发布。

---

## 3. 项目进展

### 已合并/关闭的PR（重要）

| PR | 标题 | 说明 |
|----|------|------|
| [#2713](https://github.com/nanocoai/nanoclaw/pull/2713) | feat(security): egress lockdown (opt-in, off by default) | **核心安全更新**：通过将每个Agent容器置于`--internal`网络（无互联网路由）并将OneCLI网关作为`host.docker.internal`别名暴露，实现可选的出口流量封锁。Agent只能通过经过代理的网关访问外部，大幅降低了数据泄露风险。该特性默认关闭，用户可通过配置启用。 |
| [#2716](https://github.com/nanocoai/nanoclaw/pull/2716) | [follows-guidelines] resolve check test 0609 once | 测试/CI相关合并，旨在解决一次性检查测试问题。 |
| [#2712](https://github.com/nanocoai/nanoclaw/pull/2712) | [follows-guidelines] pull request | 流程性或测试性提交，已关闭。 |

### 待合并的PR（关键）

| PR | 标题 | 说明 |
|----|------|------|
| [#2714](https://github.com/nanocoai/nanoclaw/pull/2714) | security: fix four auth/security issues | 修复四项安全/认证问题：① webhook-server默认绑定至`127.0.0.1`而非`0.0.0.0`，并提供`WEBHOOK_BIND_HOST`环境变量；② 用`crypto.randomUUID()`替代`Math.random()`生成审批ID，防止时序预测攻击；③ 其他两项审计修复。该PR合并后将进一步加固项目安全性。 |

**项目整体进展**：今日完成了出口封锁特性（#2713）的合并，标志着安全隔离能力迈出重要一步。安全修复PR（#2714）正在等待最终审查，预计未来24小时内可合并。

---

## 4. 社区热点

**#2715** — [Inbound WhatsApp media (images/docs/audio) is unreachable by the agent](https://github.com/nanocoai/nanoclaw/issues/2715)  
- 作者：jon-ruth | 创建：2026-06-08 | 评论：0 | 👍：0  

该Issue虽暂时无评论和点赞，但内容直接指向v2版本中WhatsApp媒体附件功能的不可用问题。**核心诉求**：当用户发送图片、文档、音频时，文件被下载到宿主机的`DATA_DIR/attachments`目录（该目录未挂载到Agent容器中），而Agent却被传入一个`/workspace/attachments/...`的路径，导致Agent无法打开任何媒体文件。这是一个影响核心功能使用的阻塞性问题，需要社区和维护者优先关注。

---

## 5. Bug 与稳定性

| 严重程度 | Issue | 描述 | 修复状态 |
|----------|-------|------|----------|
| **严重** | [#2715](https://github.com/nanocoai/nanoclaw/issues/2715) | WhatsApp媒体文件保存至未挂载的主机目录，Agent无法访问，导致图片/文档/音频处理功能完全失效。 | 暂无修复PR，需确认是否为配置问题或代码缺陷。 |
| **中等** | [#2714](https://github.com/nanocoai/nanoclaw/pull/2714)（PR） | 包含四项安全修复，其中webhook-server默认绑定`0.0.0.0`可能暴露服务；`Math.random()`生成审批ID存在预测攻击风险。 | PR已提交待合并，修复方案清晰。 |

此外，出口封锁特性（#2713）本身不修复Bug，但通过网络隔离减少了潜在攻击面。

---

## 6. 功能请求与路线图信号

- **出口封锁**（#2713，已合并）：用户可选的出口流量控制，属于安全增强型功能。虽然默认关闭，但为有高安全需求的部署提供了强力工具，未来很可能成为推荐配置。
- **安全修复**（#2714，待合并）：四项认证/安全修复均为2026年AI Agent安全基线要求（如禁用全局绑定、强随机数生成），极大概率纳入下一版本。
- **WhatsApp附件挂载**（#2715）：当前为Bug报告，但若用户启用WhatsApp通道，该功能是必备能力。从Issue描述看，可能涉及文件路径映射逻辑变更，属于急需修复的缺陷而非新功能请求。

---

## 7. 用户反馈摘要

从Issue #2715的描述中提炼真实用户痛点：

> *“在v2版本中，WhatsApp输入的媒体文件（图片、文档、音频）被下载到一个宿主机的目录，该目录**未被挂载**到Agent容器内。Agent获得的是一个不存在的`/workspace/attachments/...`路径，因此无法打开任何用户发送的媒体。”*

- **使用场景**：用户通过WhatsApp与NanoClaw Agent交互，发送图片/文档/音频以触发Agent处理（如OCR、文件分析等）。
- **不满意之处**：核心功能在v2上完全不可用，影响实际业务流程。用户预期文件应能被Agent直接访问。
- **隐含建议**：需要修正文件存储路径，或者将宿主机的`DATA_DIR/attachments`挂载到Agent容器的相同路径下，并确保Agent收到的路径与容器内实际路径一致。

目前该Issue尚无维护者回应，建议尽快确认并给出解决方案或临时Workaround。

---

## 8. 待处理积压

| 条目 | 状态 | 创建时间 | 备注 |
|------|------|----------|------|
| [#2715](https://github.com/nanocoai/nanoclaw/issues/2715) | OPEN，无回复 | 2026-06-08 | **高优先级**：WhatsApp媒体附件阻塞性问题，虽新但影响核心功能，应尽快指派。 |
| [#2714](https://github.com/nanocoai/nanoclaw/pull/2714) | OPEN，待合并 | 2026-06-08 | 安全修复PR，已获得`security`标签，需维护者Review并合并。 |
| 无长期未响应的老旧Issue（>7天） | — | — | 项目积压管理良好，当前无其他显著积压。 |

---

**总结**：今日项目在安全方向上取得实质进展（出口封锁合并、安全修复待合并），但一个功能级Bug（WhatsApp附件不可达）需要紧急关注。建议维护团队在下一次迭代中优先处理#2715，并尽快合并#2714以巩固安全基线。

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，这是根据您提供的 IronClaw 项目数据生成的 2026-06-09 项目动态日报。

---

# IronClaw 项目动态日报 | 2026-06-09

## 今日速览

过去24小时内，IronClaw 项目活跃度**极高**，共有 50 条 PR 和 11 条 Issue 更新，表明团队正在密集推进核心功能的迭代。PR 合并/关闭率（23/50）接近 50%，显示出高效的协作与审查流程。核心开发方向集中在 **“Reborn”架构的完善**，特别是在审批机制、工作流集成（ProductWorkflow）以及钩子系统（Hooks）的安全审计方面。社区反馈以配置问题和功能请求为主，总体趋势向好，但需关注部分因版本僵化导致的兼容性问题。

## 版本发布

**无**。项目未发布新版本。需要注意的是，一个长期搁置的自动化发布 PR（[#3708](https://github.com/ironclaw-ai/ironclaw/pull/3708)）仍在持续更新中，其包含了对 `ironclaw_common` 和 `ironclaw_skills` 等核心库的重大 API 破坏性变更，这可能意味着下一版本（预计为 0.30.0 或更高）将是一次大规模的架构升级。

## 项目进展

今日合并了 23 个 PR，主要推动了以下关键功能与修复：

1.  **Reborn 工作流集成深化**：**核心里程碑**。多个大型 PR 正在将 Reborn 架构与 OpenAI 兼容的生产工作流（ProductWorkflow）融合。
    -   [#4580](https://github.com/ironclaw-ai/ironclaw/pull/4580) (已合并)：为自动化触发器添加了运行历史 UI，增强了可观测性。
    -   [#4581](https://github.com/ironclaw-ai/ironclaw/pull/4581) (已合并)：实现了分层级的触发投递默认项，为复杂的个性化场景提供了基础设施。
    -   [#4476](https://github.com/ironclaw-ai/ironclaw/pull/4476) (已合并)：通过 `NormalizingProvider` 装饰器标准化了 LLM Provider 的行为，并强制修复了部分提供商的 `finish_reason` 问题。
    -   [#4472](https://github.com/ironclaw-ai/ironclaw/pull/4472) (已合并)：重构了子代理（Subagent）的创建逻辑，并引入了新的“规划者（Planner）”类型。

2.  **工具层稳定性修复**：
    -   [#4578](https://github.com/ironclaw-ai/ironclaw/pull/4578) (已合并)：修复了谷歌日历工具 `list_events` 默认返回最早事件而非未来事件的逻辑错误。
    -   [#4523](https://github.com/ironclaw-ai/ironclaw/pull/4523) 和 [#4575](https://github.com/ironclaw-ai/ironclaw/pull/4575) (已合并)：修复了系统相关的 `ResourceScope` 序列化/反序列化问题，解决了 LLM 设置报错的问题。

**项目向前迈进了一大步**。“Reborn”架构已从理论设计进入密集的代码落地阶段，特别是与 OpenAI 兼容 API 的深度集成，标志着项目正向一个功能更强大、架构更现代的新版本过渡。

## 社区热点

今日社区讨论的焦点集中在一系列**由核心成员发起的大型架构性 PR** 上。虽然没有单一 Issue 拥有特别高的评论数，但多个 XL 规模的 PR 共同构成了今日最活跃的讨论区域。

1.  **[PR #4590](https://github.com/ironclaw-ai/ironclaw/pull/4590): Add Slack shared outbound targets** - 该 PR 引入了 Slack 频道作为外发目标，表明项目正在扩展智能体触发后通知的交付渠道，社区对此功能兴趣浓厚，希望看到更多第三方集成的支持。
2.  **[PR #4559](https://github.com/ironclaw-ai/ironclaw/pull/4559):  agent-driven Trace Commons onboarding via invite link** - 此 PR 实现了用户通过链接让智能体自行完成 Trace Commons 注册的流程，大幅简化了操作，反响热烈。其背后反映了用户对“零配置”或“低配置”使用体验的强烈诉求。
3.  **[PR #4495](https://github.com/ironclaw-ai/ironclaw/pull/4495) 与 [PR #4546](https://github.com/ironclaw-ai/ironclaw/pull/4546): 路由Chat/Responses通过 ProductWorkflow** - 这两个 PR 是今日最核心的架构变更。开发者和社区正在围绕其实现细节、与现有系统的兼容性以及性能影响进行深入探讨，体现了社区对项目核心架构发展的极高关注度。
4.  **[PR #4186](https://github.com/ironclaw-ai/ironclaw/pull/4186): [codex] Wire local-dev approval gates** - 这个 PR 虽然创建较早（5月28日），但仍在更新中，它连接了本地开发的审批门控逻辑，是响应社区“审批功能缺失”诉求的重要一步。

**分析**：社区的主要诉求是**功能的完整性和易用性**。用户不仅希望核心功能（如聊天补全）通过新架构（Reborn）跑通，更希望获得像 Slack 集成、自动化投递、简化注册等提升生产力和体验的配套设施。

## Bug 与稳定性

今日报告的 Bug 属于中等严重程度，主要集中在功能行为异常和配置兼容性上。

-   **SP1 (严重)**:
    -   **[#4587](https://github.com/ironclaw-ai/ironclaw/issues/4587): 无法配置 Minimax Provider**。用户在配置 Minimax 并使用时遇到秘钥凭据读取失败。这直接导致该 LLM Provider 不可用。目前有 PR正在讨论中，但尚未有针对性的修复 PR 被合并。
-   **SP2 (行为异常)**:
    -   **[#4577](https://github.com/ironclaw-ai/ironclaw/issues/4577): Google Calendar 返回最旧事件**。该 Bug 会导致“查看我最近的会议”这类常用功能返回错误数据，严重影响用户体验。**已通过 PR [#4578](https://github.com/ironclaw-ai/ironclaw/pull/4578) 修复**。
-   **SP3 (边缘情况)**:
    -   **[#4560](https://github.com/ironclaw-ai/ironclaw/issues/4560): Trace Commons 上线流程绕过了网络出口策略**。这在多租户生产环境中可能成为安全或策略执行的漏洞。
-   **SP4 (拦截器)**:
    -   **[#4585](https://github.com/ironclaw-ai/ironclaw/issues/4585): Reborn 认证证据应携带租户身份**。这是一个架构层面未覆盖到的场景，可能导致跨租户的认证和授权验证不完整，被标记为由之前 PR 审查中发现的跟进项。

## 功能请求与路线图信号

-   **核心信号：审批功能对等性**。Issue **[#4539](https://github.com/ironclaw-ai/ironclaw/issues/4539)** 明确提出了 Reborn 架构需要达到 V1 版本的审批功能对等性。这包括“批准一次”、“拒绝”和“始终允许”等操作。结合相关的核心开发 PR（如[#4186](https://github.com/ironclaw-ai/ironclaw/pull/4186)），**可以判断审批功能的补全是下一版本发版前的核心要务**。
-   **新需求：超时与配额检查**。Issue **[#4586](https://github.com/ironclaw-ai/ironclaw/issues/4586)** 提出了验证 OpenAI 兼容接口中的超时行为是否受配额限制。虽然创建者本人随后澄清该问题并非要新增一个配额系统，但揭示了对运行时边界控制和资源管理的关注，这将是 Reborn 架构成熟化的必经之路。
-   **路线图印证：Reborn 操作能力**。Issue **[#4533](https://github.com/ironclaw-ai/ironclaw/issues/4533)** 作为 Epic，总结了 Reborn 上线前必须解决的安装、配置、诊断和生命周期管理等问题。这直接印证了路线图中从“功能开发”转向“运营就绪”的阶段。

## 用户反馈摘要

-   **可靠性质疑**：来自 Issue [#4587](https://github.com/ironclaw-ai/ironclaw/issues/4587) 的用户 `matiasbenary`。用户按照文档配置了 Minimax Provider 后，发现服务无法使用，错误日志暴露出“could not read stored key metadata”。这是一个明确的痛点：**配置体验不佳且缺乏清晰的错误指引**。用户可能会因此认为该 Provider 支持不成熟。
-   **心智负担**：同一个用户在同一 Issue 的评论中可能透露出对复杂配置的挫败感。他们期望能像使用 OpenAI 那样简单的 API Key 配置方式，而不希望暴露复杂的内部 Hash 结构。
-   **功能缺失的挫败感**：针对审批功能（[#4539](https://github.com/ironclaw-ai/ironclaw/issues/4539)），虽然社区没有留下长篇大论，但创建 Epic 本身就是一个强烈的信号，表明“不能批准”或“只能全部放行”的工作流对于非个人开发者用户来说是**不可接受的**。这直接影响项目能否从个人爱好工具升级为团队协作产品。

## 待处理积压

-   **[#3957](https://github.com/ironclaw-ai/ironclaw/issues/3957) (开放): 第三方钩子激活强化跟进** - 创建于 2026-05-23。这是一个安全相关的关键 Issue，涉及 `HOOKS_THIRD_PARTY_ENABLED` 特性在生产环境开启前的最终安全检查。已有 2 条评论但无核心维护者的近期公开回应，作为安全里程碑，需提醒关注以免阻塞发布。
-   **[#3959](https://github.com/ironclaw-ai/ironclaw/issues/3959) (开放): 在剩余边界调用点采用 SecurityAuditSink** - 创建于 2026-05-23。同样是安全审计基础设施的全面铺开工作，与 #3957 相同，是提升系统安全可观测性的关键未完成任务。
-   **[#4186](https://github.com/ironclaw-ai/ironclaw/pull/4186) (开放): 连接本地开发的审批门控** - 这个 PR 从 5 月 28 日开放至今，作为连接 #4539 Epic 的关键组件，其长时间未合并可能会阻塞后续相关功能的开发。维护者应优先审查此 PR。

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

好的，作为 LobsterAI 开源项目的分析师，以下是基于 2026-06-09 GitHub 数据生成的项目动态日报。

***

# LobsterAI 项目动态日报 | 2026-06-09

## 1. 今日速览

今日项目活跃度**高**，主要集中在 **PR 的快速合并与清理**上。过去 24 小时内，共有 **17 条 PR 被处理**，其中 16 条已合并/关闭，显示出团队高效的协作与交付能力。合并内容覆盖了数据迁移、协作功能增强、认证流程优化以及多项 Bug 修复，项目代码库向前迈进了坚实的一步。社区讨论方面相对平静，仅有一条新 Issues 提出关于支持 Hermes Agent 的功能请求。

## 2. 版本发布

无新版本发布。

## 3. 项目进展

今日有 **16 条 PR 被合并或关闭**，项目在以下方面取得了显著进展：

- **数据迁移与备份（Data Migration）**：合并了 `#2125`, `#2126`, `#2128` 三个 PR，引入了完整的用户数据备份与恢复功能，并修复了在恢复过程中排除网络目录、保留运行时锁文件等关键问题。这标志着 LobsterAI 的数据安全性和可移植性得到了大幅提升。
- **协作功能增强（Cowork）**：合并了 `#2130` 功能，为协作会话添加了隐私安全的任务完成通知，即使应用在后台也能通过系统通知、macOS Dock 角标和 Windows 任务栏提醒用户。
- **认证流程优化（Auth）**：合并了 `#2127`, `#2129` 两个 PR，重点修复了 Windows 平台下回调登录后窗口无法自动获得焦点的问题，并增加了登录回调的诊断日志，方便排查问题。
- **Bug 修复与功能打磨**：合入了多个标记为 `[stale]` 的 PR，这通常意味着较长时间未处理的修复或功能。包括：
    - `#1510`：修复了定时任务在选择IM通知频道但未选会话时，任务静默失败的问题。
    - `#1514`：修复了 QQ Bot 群组白名单设置缺少输入框，无法通过 UI 配置的问题。
    - `#1515`：修复了导出日志时因压缩耗时过长而超时的问题。
    - `#1517`：修复了关闭设置面板时，GitHub Copilot OAuth 轮询未被取消导致 Token 丢失的问题。
    - `#1521`：修复了 OpenClaw 网关因 `skills-changed` 消息导致意外重启的问题。
    - `#1522`：新增了“动态获取模型列表”功能，允许用户从模型提供商API自动拉取最新模型列表。
    - `#1524`：优化了“测试连接”失败时的错误信息，提供更具体的网络、鉴权等上下文。
    - `#1526`：为协作会话列表添加了颜色标注功能，支持 7 种颜色，方便用户视觉区分不同会话。
- **OpenClaw 网关可观测性**：合并了 `#2123`，在设置中公开了 OpenClaw 网关的 URL 和运行状态，增强了用户对运行时信息的掌控和排障能力。

## 4. 社区热点

今日社区讨论热度一般，但有以下值得关注的点：

- **Issue #2131：支持 Hermes Agent 的计划**
    **热度**：新开 Issues，1条评论。
    **诉求**：用户询问 LobsterAI 未来是否有计划支持 Hermes Agent。这反映了社区对连接和使用多元化 AI Agent 生态的兴趣，希望 LobsterAI 成为更开放的 Agent 管理平台。
    **链接**：[#2131](https://github.com/netease-youdao/LobsterAI/issues/2131)

- **PR #2130：协作任务完成通知**
    **热度**：作为今天合并的核心功能之一，受关注度较高。
    **诉求**：该 PR 的诉求非常明确——提升活跃用户的协作体验。当 LobsterAI 不在前台时，用户仍能即时获知任务完成状态，无需反复切换窗口检查，体现了对“非侵入式”交互体验的重视。
    **链接**：[#2130](https://github.com/netease-youdao/LobsterAI/pull/2130)

## 5. Bug 与稳定性

今日未报告新的严重 Bug。但通过合入大量 `[stale]` 的 PR，项目集中解决了一批长期存在的 Bug：

- **`[高]` 定时任务 IM 通知静默失败 (PR #1510)**：当用户未选择会话时，任务创建时无提示，运行时通知静默失败。**已修复**。
- **`[中]` QQ Bot 白名单 UI 缺失 (PR #1514)**：切换至白名单模式后，无法通过 UI 添加群组 ID，功能完全不可用。**已修复**。
- **`[中]` 日志导出超时 (PR #1515)**：在配置较低的机器上，大量日志压缩过程超过30秒超时限制，导致导出失败。**已修复**。
- **`[中]` Copilot OAuth Token 静默丢失 (PR #1517)**：在轮询认证期间关闭设置面板，会导致 Token 丢失，用户需要重新认证。**已修复**。
- **`[低]` OpenClaw 网关意外重启 (PR #1521)**：`skills-changed` 消息触发了不必要的网关重启。**已修复**。

这些修复体现了项目对稳定性和用户体验的持续打磨，消除了多个潜在的“隐形”问题。

## 6. 功能请求与路线图信号

- **`[新功能请求]` 支持 Hermes Agent (Issue #2131)**：这是一个明确的路线图信号，表明一些用户希望 LobsterAI 不仅仅局限于自身或特定的 Agent，而是能成为一个支持多种 Agent 类型的统一平台。此请求的优先级取决于团队的战略方向。

- **`[已实现]` 动态获取模型列表 (PR #1522)**：用户之前只能在设置中手动添加新发布的模型。现在 PR #1522 已合并，支持从 Provider API 动态拉取模型列表，极大提升了用户体验，该功能很可能会被包含在下一个版本中。

- **`[已实现]` 会话颜色标注 (PR #1526)**：用户此前缺乏可视化的方式来区分不同类型的会话。此 PR 实现了 7 种颜色标注功能，满足了用户进行个性化组织和管理会话的需求。

- **`[技术债务]` 依赖更新 (PR #1277)**：一个已开放两个多月的 PR #1277，旨在将 `electron` 从 40.2.1 更新至 42.3.3，并更新 `electron-builder`。尽管未合并，但依赖更新对安全性和性能至关重要，通常会在版本发布前被纳入。

## 7. 用户反馈摘要

今日 Issues/PRs 评论中公开的用户反馈较少，但从修复的 Bug 中可以反推痛点：

- **用户对交互反馈的透明度和及时性有较高要求**。如“日志导出超时”和“定时任务 IM 通知静默失败”修复表明，用户对长时间无反馈的操作（如导出、后台任务）持负面看法，期望系统能提供清晰的进度或失败提示。
- **配置的易用性直接影响用户满意度**。“QQ Bot 白名单 UI 缺失”和“Copilot 认证 Token 丢失”都源于用户界面或流程上的设计缺陷，导致功能无法使用或体验断层。此次集中修复将显著提升相关用户的满意度。
- **用户希望终端应用更“智能”和“开放”**。“动态获取模型列表”和“支持 Hermes Agent”的诉求，都指向用户希望 LobsterAI 能自动感知并接入外部的、不断发展的 AI 能力，而不是一个静态的封闭系统。

## 8. 待处理积压

- **PR #1277：依赖更新 (electron 40 -> 42)**：此 PR 由 `dependabot[bot]` 创建于 2026-04-02，至今已超过两个月未获合并或实质性回复。从 Electron 40 到 42 的跨版本更新通常涉及重大的安全修复和性能改进，以及对新 API 的支持。长期搁置可能带来技术债务和潜在的安全风险。
    **链接**：[#1277](https://github.com/netease-youdao/LobsterAI/pull/1277)

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

# TinyClaw 项目日报 (2026-06-09)

## 1. 今日速览

过去24小时内项目整体活跃度处于低位：Issues 无新增或关闭，无新版本发布。唯一值得关注的是 PR #280（待合并）被提出，旨在通过添加 postinstall 脚本自动重建 `better-sqlite3` 原生模块，消除用户手动 `npm rebuild` 的步骤。该 PR 触及了项目安装流程的痛点，对提升首次使用体验有积极意义。整体来看，项目当前处于功能优化与稳定性打磨阶段，开发者互动频率较低。

## 2. 版本发布

无新版本发布。

## 3. 项目进展

今日无任何 PR 被合并或关闭，项目核心代码库无新功能落地。唯一的开放 PR #280 尚待评审，其若被合并将解决因 `better-sqlite3` 原生模块未能自动编译导致的安装错误，属于安装流程的自动化改进。建议维护者优先评审该 PR 以避免用户重复遇到相同问题。

## 4. 社区热点

- **PR #280** – [[OPEN] fix(install): add postinstall script to auto-rebuild better-sqlite3](https://github.com/TinyAGI/tinyagi/pull/280)  
  作者：dsy122 | 创建：2026-06-08  
  该 PR 虽无评论互动，但其摘要直指用户频繁报告的安装失败场景。因为 `better-sqlite3` 是 C++ 原生模块，不同 Node.js 运行时版本需要重新编译，而当前 `npm install` 后无自动重建步骤，导致新用户/新环境首次运行极易报错。该 PR 的提出反映了社区对安装体验自动化的迫切需求，是项目优化运维体验的重要信号。

## 5. Bug 与稳定性

今日未报告新的 Bug 或崩溃问题。但 PR #280 间接指出了现有安装流程中的稳定性缺陷：`better-sqlite3` 未在 `postinstall` 阶段自动重建，使得依赖该模块的数据库功能在初次部署时不可用。此问题虽无 issue 追踪，但若该 PR 合并，将被视为一次稳定性修复。

## 6. 功能请求与路线图信号

无新的功能请求出现。PR #280 的性质属于安装流程优化，并非新功能，但反映了项目对“零配置”方向的重视。如果维护者接受此模式，未来可能进一步推广到其他原生依赖模块（如 `node-canvas` 等）的自动重建，降低用户上手门槛。

## 7. 用户反馈摘要

今日无用户反馈记录。但 PR #280 的摘要可视为间接用户反馈：“用户会在全新安装时遇到错误，因为预编译的 `better-sqlite3` 与当前 Node.js 运行时不匹配”。这提示项目需要更健壮的安装脚本或文档指导。

## 8. 待处理积压

- **PR #280** – 自提交后已过去1天仍未收到任何评审意见或评论。该 PR 改动较小（仅新增 `postinstall` 脚本），且能显著改善安装体验，建议维护者尽快分配 reviewer 并推动合并，避免同类投诉反复出现。  
  链接：https://github.com/TinyAGI/tinyagi/pull/280

---

*数据来源：TinyClaw GitHub 仓库 (github.com/TinyAGI/tinyagi) *

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

好的，这是根据您提供的 CoPaw 项目 GitHub 数据生成的 2026-06-09 项目动态日报。

---

## CoPaw 项目动态日报 — 2026-06-09

### 1. 今日速览

过去 24 小时，CoPaw 项目保持高度活跃态势，尤其是拉取请求（PR）数量激增至 47 条，表明核心开发工作正在快速推进。合并/关闭的 PR 数量（24）与新开 PR 数量（23）基本持平，显示出良好的代码审查与合并循环。Issues 方面，虽然新开/活跃数量（8）不多，但关闭数量（13）更高，体现了项目在回馈社区反馈方面的高效率。整体来看，项目在修复稳定性问题（如 OneBot 频道、MCP 工具兼容性）和推进新功能（如 TUI、Plugin Market）方面均取得显著进展。

### 2. 版本发布

无新版本发布。

### 3. 项目进展

今日有多个关键 PR 被合并，推进了项目的核心功能和稳定性：

- **OneBot 频道稳定性修复**：PR #5010 被合并，解决了 OneBot 频道在零停机重载时因端口冲突导致的永久性故障，改为此问题的优雅降级处理。这直接回应了 Issue #4926 和社区反馈。
- **钉钉频道消息合并 Bug 修复**：PR #4932 被合并，修复了因 `conversation_id` 后缀冲突导致不同用户消息被错误合并的 Bug，提升了多用户场景下的消息隔离性。
- **新功能：元宝（Yuanbao）频道“正在输入”提示**：PR #5020 被合并，为 Yuanbao 频道增加了通过心跳 API 实现的“正在输入”状态指示器，提升了用户交互体验。
- **新功能：控制台插件扩展基础设施**：PR #4997 被合并，为前端插件提供了统一的菜单、路由和 UI 插槽注册机制，是构建丰富插件生态的关键基础设施。
- **TUI 终端聊天界面第一阶段**：PR #5032 处于开放状态，旨在将终端 TUI 集成到主 `qwenpaw` 命令中，目标是提供无需图形界面的轻量级交互方式。

这些进展表明，项目正同时着力于修复高优先级 Bug 和建设面向未来的功能架构。

### 4. 社区热点

今日社区讨论热度最高的议题集中在以下方面：

- **Agent 学习与进化能力**：Issue #5017 获得了 7 条评论和 2 个 👍，社区成员建议 CoPaw 关注并借鉴 Hermes Agent 的“学习循环”（Learning Loop）特性，让 Agent 能自动从自身行为中创建和迭代技能。这反映出用户对 Agent 自主进化能力的强烈期望。
- **稳定性和体验问题**：Issue #4477（15 条评论）关于微信定时任务推送失败，Issue #4123（9 条评论）关于 Windows 下 Shell 命令执行闪窗，以及 Issue #5003（8 条评论）关于阿里 Coding Plan 模型卡死，是今日讨论最多的 Bug 报告，直接影响了用户的核心使用体验。

这些热点清晰指向了用户的两大核心诉求：一是希望 Agent 更“智能”，能自主学习和进化；二是希望现有功能的稳定性和体验能进一步提升。

### 5. Bug 与稳定性

今日报告的 Bug 主要包括：

| 严重程度 | 问题描述 | Issue 链接 | 状态 | 修复 PR |
| :--- | :--- | :--- | :--- | :--- |
| **高** | **MCP 工具名包含“.”导致 OpenAI API 报错 400**，模型无法调用工具。 | [#5034](https://github.com/agentscope-ai/QwenPaw/issues/5034) | 已关闭 | 已声称在后续版本修复，参考 `stateful_client.py`。 |
| **高** | **内存压缩期间 `AttributeError`**，导致 Agent 崩溃。 | [#5019](https://github.com/agentscope-ai/QwenPaw/issues/5019) | 已关闭 | 未提及具体 PR，但 Issue 标记为已关闭，疑似已找到原因或修复。 |
| 中 | **`submit_to_agent` 会话文件路径 Bug**，导致后台任务执行失败。 | [#5025](https://github.com/agentscope-ai/QwenPaw/issues/5025) | 开放 | [#5036](https://github.com/agentscope-ai/QwenPaw/pull/5036) 已提交修复。 |
| 中 | **Windows Desktop 前端加载卡顿**，任务执行时 CPU 激增。 | [#5015](https://github.com/agentscope-ai/QwenPaw/issues/5015) | 开放 | 暂无 |
| 低 | **Skill 斜杠命令在 Console 中显示为完整 Markdown**，而非简洁形式。 | [#5031](https://github.com/agentscope-ai/QwenPaw/issues/5031) | 开放 | 暂无 |

此外，`ly-wang19` 作为首次贡献者提交了 3 个针对边界情况的 Bug 修复 PR，覆盖了空消息列表、空 `Exec=` 配置和 llama.cpp 版本号解析，表明社区正在积极修补代码健壮性。

### 6. 功能请求与路线图信号

- **Agent 学习与进化**：Issue #5017 建议借鉴 Hermes Agent 的“学习循环”功能。目前尚未有对应的 PR 或路线图计划，但社区关注度高，可能成为未来重要发展方向。
- **可观测性集成**：Issue #5009 询问了集成 Langfuse、OpenTelemetry 等可观测性平台的计划。这表明在社区和企业用户中，对生产级监控、链路追踪和成本分析有明确需求。
- **Plugin Market（插件市场）**：PR #5023 正在开发中，旨在社区内提供一个可浏览、搜索和安装插件的市场。这与 PR #4997（插件基础设施）相辅相成，表明一个开放的插件生态系统正在构建中，很可能在下一版本中上线。
- **轻量级目标模式**：PR #4443 提出添加 `/goal` 模式，允许用户为会话设定一个长期目标并持续跟踪。该 PR 仍在开放状态，但若被采纳，将助力 Agent 实现更复杂的任务。

### 7. 用户反馈摘要

- **积极反馈**：用户 `tecgic` 在 Issue #5017 中称赞 CoPaw“本地化做得很到位，设置清晰无门槛，开箱即用”，这是对项目国内体验的极高评价。
- **稳定性痛点**：
  - 用户 `QinQiang123` 反馈，开启主动模式后微信频道出现一个问题两次回复的异常，表明该模式下消息去重机制存在缺陷。
  - 用户 `frederickwang-blip` 则对 Pet 功能的稳定性和性能表示不满，称其“闪退、卡顿严重，体验极差”，并建议标记为实验性功能。
  - 用户 `zhanggegehao` 详细描述了微信定时任务推送失败的问题，核心在于 `context_token` 过期后缺乏重试机制，以及文件发送失败时无日志记录。
- **功能期望**：用户 `flyrae` 提出的可观测性需求，以及用户 `tecgic` 提出的学习循环建议，都代表了对 Agent 框架深度能力和生产化应用的前瞻性思考。

### 8. 待处理积压

- **长期未决 Issue**：
  - **Issue #2777**：关于 GPT-5.x 模型 `max_tokens` 参数错误的问题，自 4月1日创建以来，虽持续更新但尚未关闭。这是一个影响主流模型使用的 Bug，应优先处理。

- **长期未结 PR**：
  - **PR #4669**：桌面端 Tauri 自动更新功能，自 5月25日提交，处于开放状态。对于桌面应用的体验至关重要，建议维护者加快审查与合并进程。
  - **PR #4443**：轻量级目标模式，虽不是高优先级 Bug，但其概念新颖且社区有潜在需求，建议在下一次核心功能选型时讨论。

---
*数据来源：CoPaw GitHub 仓库 (github.com/agentscope-ai/CoPaw)，数据采集于 2026-06-09。*

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域的开源项目分析师，以下是根据您提供的 ZeroClaw 项目数据生成的 2026 年 6 月 9 日项目动态日报。

---

### ZeroClaw 项目动态日报 | 2026-06-09

**分析师点评：** 项目今日维持高活跃度，Issues 与 PR 数量均处于高位。核心团队正围绕 `v0.8.0` 系列版本进行最后的冲刺，大量修复和功能 PR 处于待合并状态，表明发布节奏在加快。重要 Bug (如 WASM 插件路径、Matrix 会话隔离) 的修复方案已提出，社区提出的新问题也得到了快速响应。

---

### 1. 今日速览

-   **高活跃度，发布冲刺期**：过去 24 小时内有 17 条 Issue 更新和 50 条 PR 更新，整体活跃度极高。这表明项目正处于密集的开发迭代和问题修复阶段。
-   **修复与功能并重**：大量 PR 集中在修复核心 Bug（如插件路径、Cron 调度、Matrix 会话）和推进新功能（如观测性、Webhook 路由、Slash 命令）。
-   **`v0.8.0` 系列版本接近但尚未发布**：虽然多个 `v0.8.x` 的跟踪 Issue（`#7112`, `#6970`, `#7314`）正在积极更新，但今日无新版本发布，说明团队仍在谨慎地合并关键修复。
-   **社区参与度高**：多个由社区成员提交的 PR (如 `#7410`, `#7412`, `#7409`) 和 Bug 报告得到了及时响应，项目生态活跃。

### 2. 版本发布

-   无。

### 3. 项目进展

今日有两项重要的 PR 被关闭，代表着关键问题的解决：

-   **Matrix 会话隔离修复**：PR [#7388](https://github.com/zeroclaw-labs/zeroclaw/pull/7388) (`singlerider`) 已合并/关闭。该 PR 直接解决了长期存在的 `[Bug]: Multi-alias Matrix instances share one session store and clobber each other` (`#6487`)，通过为每个 Matrix 别名创建独立的会话存储路径，彻底修复了多实例相互覆盖的关键问题。
-   **ESP32 智能家居 Demo 完成**：PR [#6148](https://github.com/zeroclaw-labs/zeroclaw/pull/6148) (`Rhoahndur`) 已合并/关闭。这标志着一段为期较长的 Hackathon 项目最终落地，展示了 ZeroClaw 与 Telegram 及硬件 (ESP32) 集成的端到端能力。

这些进展表明项目在稳定性（Matrix 修复）和功能性（硬件Demo落地）上均迈出了坚实的一步。

### 4. 社区热点

-   **MCP 工具过滤器失效 (`#6699`)**：该 Issue 获得了最多的评论 (7 条)，讨论热度最高。用户 `nick-pape` 报告了 `tool_filter_groups` 配置对真实的 MCP 工具完全无效，且与延迟加载机制存在冲突。这是影响 `agent` 和 `tool` 配置的核心 Bug，社区对此高度关注。 [👉 参与讨论](https://github.com/zeroclaw-labs/zeroclaw/issues/6699)
-   **核心团队在 PR 评论区密集互动**：今日无单一 PR 获得极高评论数，但核心成员 `singlerider`, `Audacity88` 等在多条 PR 中活跃，并对新报告的 Issue (`#7412`) 快速响应。这表明核心团队正在积极吸收社区反馈。

**分析**：社区热点集中在配置生效和核心功能的可靠性上。关于 MCP、插件路径、安全过滤等问题的讨论，反映了用户希望 ZeroClaw 作为 Agent 框架，其配置系统需要高度的确定性和可预测性。

### 5. Bug 与稳定性

-   **严重 Bug (P0/P1) - 已修复**：
    -   **[P0 Blocker] Matrix 多实例会话冲突** (`#6487`)：已通过 PR `#7388` 修复。
    -   **[P1] 工具过滤器对 MCP 工具无效** (`#6699`)：暂无已关联的修复 PR，但已被核心团队接受 (`status:accepted`)，预计将是 `v0.8.0` 的修复重点。
    -   **[P1] 系统过度强调记忆** (`#5844`)：暂无修复 PR，该问题影响用户体验，尤其是在定时任务场景下。

-   **严重 Bug (P1/P2) - 已有修复 PR**：
    -   **[P1] WASM 插件安装与扫描路径不一致** (`#6254`)：PR `#7413` (修复: 插件路径) 直接指向此 Issue。该 PR 正在开放中，等待合并。
    -   **[P2] Cron 跳过逾期任务功能失效** (`#7250`)：PR `#7348` 提出了修复方案，目前在待合并状态。

-   **新报告的 Bug**：
    -   **[P0] Webhook 签名密钥缓存问题** (`#7410`)：`IftekharUddin` 报告，建议改为运行时从配置读取，而非启动时缓存。暂无修复 PR，但已提交。
    -   **[P2] `allowed_private_hosts` 配置对域名解析的私有 IP 不生效** (`#7412`)：`NiuBlibing` 报告了一个安全相关的重要 Bug。
    -   **[P2] Clippy 检查仅限 Linux，导致跨平台代码质量下降** (`#7409`)：`NiuBlibing` 指出了 CI 流程中的一个潜在缝隙。

### 6. 功能请求与路线图信号

-   **Web 与体验优化**：PR `#7223` (Web Slash 命令) 和 `#7367` (Webhook 路由) 正在推进中，表明下一版本将显著增强 Web 交互和渠道集成体验。Issue `#4832` (禁用高熵令牌脱敏) 虽然目标是 `v0.8.0`，但尚未有关联 PR。
-   **安全架构升级**：Issue `#7142` (可插拔安全提供者) 是一个大型的 RFC，目标是 `v0.9.0`，表明项目正在规划长远的安全架构演进。
-   **WASM 与插件生态**：`v0.8.2` 和 `v0.8.3` 的跟踪 Issue (`#7314`, `#7320`) 聚焦于 WASM 插件和 MCP 仪表盘，这是构建强大开发者生态的关键基础设施。
-   **配置编辑体验**：PR `#7267` 提出了对 `[[mcp.servers]]` 进行字段级编辑的增强，这表明项目正在持续优化用户配置体验。

### 7. 用户反馈摘要

-   **痛点**：
    -   **配置复杂性**：多位用户反馈配置不生效或行为与预期不符 (如 `#6699` 工具过滤器、`#6254` 插件路径、`#7412` 私有域名过滤)。
    -   **行为不一致**：用户 `databillm` 在 `#5844` 中抱怨系统“过度强调记忆”而忽略当前 Prompt，影响了 Agent 对任务的专注度。
    -   **跨平台支持**：用户 `NiuBlibing` 在 `#7409` 中指出 Windows/macOS 平台的代码质量由于 CI 限制无法得到保证。
-   **满意之处**：
    -   社区对 Bug 的响应速度感到满意，如 `#7412` 报告后很快得到维护者关注。
    -   功能请求如 Telegram 智能截断 `#6225` 和 ESP32 Demo `#6148` 得到认可并最终实现。

### 8. 待处理积压

-   **关键 Issue 待关注**：
    -   `#4832` [P2, 高影响]: **禁用高熵令牌脱敏功能** (创建于 2026-03-27)。该 issue 已经开放近 2.5 个月，虽然优先级较低，但其背后涉及误报问题，对部分用户有实际影响。
    -   `#6645` [P2, 高风险]: **SkillImprover 无法识别 `manifest.toml`** (创建于 2026-05-14)。该问题会导致内置 Skill 无法被改进和管理，影响核心功能，建议尽快确定修复方案。

-   **重要 PR 等待回复/合并**：
    -   `#6973` [P1]: **WhatsApp 通道 LID JID 修复**。该 PR 已有 `needs-author-action` 标签，需要作者更新。
    -   `#6970` [P2]: **`v0.8.1` 集成/渠道/提供者 PR 队列**。这是一个重要的跟踪 Issue，但目前无任何评论，可能需要定期审视其中的 PR 状态。

---
**报告结束。**

</details>