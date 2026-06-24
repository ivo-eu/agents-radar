# OpenClaw 生态日报 2026-06-24

> Issues: 36 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-06-24 10:35 UTC

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

# OpenClaw 项目日报 — 2026-06-24

## 1. 今日速览

过去 24 小时 OpenClaw 保持高度活跃：共处理 **36 条 Issue**（新开/活跃 22 条，关闭 14 条）和 **500 条 PR**（待合并 454 条，合并/关闭 46 条）。**v2026.6.10 版本** 发布，引入了对话快速模式与更可靠的模型路由，但也伴随多项关键回归（如 Anthropic Vertex 纯文本响应不可见、Lossless-Claw 上下文引擎初始化失败）。社区讨论集中在 MCP 工具注入、内存存储迁移及跨通道消息丢失等问题上。PR 积压量依然庞大（454 条待合并），需维护者加快审阅节奏。

## 2. 版本发布

**v2026.6.10** 主要亮点：
- **对话自动快速模式**：OpenClaw 可在短轮次中启用 fast mode，长轮次自动恢复为正常模式，并带有有界的回退与交付行为（[#85104](https://github.com/openclaw/openclaw/pull/85104)）。
- **更可靠的模型路由**：改进了 Zai 模型的合成与路由选择。

**⚠️ 已知破坏性变更与迁移注意事项**：
- **Anthropic Vertex 路由变更**：`route=native` 导致纯文本响应不可见（[#96337](https://github.com/openclaw/openclaw/issues/96337)），用户需回退到 `2026.6.8` 的 `proxy-like` 配置或等待修复。
- **Lossless-Claw 插件兼容性**：`v2026.6.10` 的只读插件注册机制导致 `lossless-claw` 引擎初始化失败（[#96335](https://github.com/openclaw/openclaw/issues/96335)），建议暂停使用该插件直到上游修复。
- **内存存储位置变更**：从 `2026.6.9` 起，内存向量存储被静默迁移至新路径，无迁移提示，需手动确认路径（[#95495](https://github.com/openclaw/openclaw/issues/95495)）。

## 3. 项目进展

今日合并/关闭的关键 PR：

- **[#96345](https://github.com/openclaw/openclaw/pull/96345) feat(copilot): add BYOK provider parity** — 为 Copilot 添加了 BYOK 提供商对等支持，增强了运行时会话一致性。
- **[#96368](https://github.com/openclaw/openclaw/pull/96368) fix(model): `/model <default>` 写入 override** — 修复当会话使用非默认模型时，`/model <default>` 命令静默失败的问题。
- **[#68936](https://github.com/openclaw/openclaw/pull/68936) Autofix: add PR review autofix pipeline + Windows daemon** — 新增基于 Claude Agent SDK 的 PR 审查自动修复流水线，并附带 Windows 后台守护进程。
- **[#93998](https://github.com/openclaw/openclaw/pull/93998) fix: fail closed when configure is run non-interactively** — `openclaw configure` 在非交互式终端下立即报错退出，防止异常终止。
- **[#96371](https://github.com/openclaw/openclaw/pull/96371) fix(agents): do not suppress tool error warning when it is the only reply** — 解决工具执行失败后模型无文本回复时，用户仅看到 “failed” 草稿而收不到最终回复的问题。

此外，多条针对外部响应读取未绑定的安全性修复 PR（如 [#96043](https://github.com/openclaw/openclaw/pull/96043)、[#96031](https://github.com/openclaw/openclaw/pull/96031)、[#96026](https://github.com/openclaw/openclaw/pull/96026)）处于待审状态，整体项目正向安全性加固与稳定性方向稳步推进。

## 4. 社区热点

过去 24 小时讨论最活跃的 Issue：

- **[#85030](https://github.com/openclaw/openclaw/issues/85030) — MCP tools not injected into subagent sessions**（评论 9，👍 3）  
  用户报告通过 `mcp.servers` 注册的 MCP 工具无法注入到 `sessions_spawn` 子会话中，无论 allowlist 如何配置均被忽略。这是高优先级（P1）行为缺陷，社区期望维护者尽快确认修复方向。

- **[#7722](https://github.com/openclaw/openclaw/issues/7722) — Filesystem Sandboxing Config (tools.fileAccess)**（评论 9，👍 4）  
  用户要求通过配置文件限制文件系统访问路径，此功能请求已开放近 5 个月，社区持续呼吁加入安全边界。

- **[#95495](https://github.com/openclaw/openclaw/issues/95495) — 2026.6.9 silently relocates memory store**（评论 7，👍 1）  
  升级后内存向量存储被静默迁移，导致 1499 个文件重新嵌入，用户抱怨缺乏升级警告，严重性为 P1。

- **[#39847](https://github.com/openclaw/openclaw/issues/39847) — Echo contamination: stripInboundMetadata not called**（评论 6，👍 1）  
  LLM 回复中泄露内部元数据，安全影响高（P1），但已开放三个多月，社区认为修复优先级应提高。

## 5. Bug 与稳定性

以下为今日报告的高严重性 Bug，按优先级排列，并标注是否已有修复 PR：

| Issue | 严重性 | 描述 | 已有 Fix PR |
|-------|--------|------|-------------|
| [#85030](https://github.com/openclaw/openclaw/issues/85030) | P1 | MCP 工具不注入子会话，allowlist 失效 | 否 |
| [#95495](https://github.com/openclaw/openclaw/issues/95495) | P1 | 内存存储静默迁移导致全量重嵌入 | 关闭，但未提供修复说明（可能已回滚） |
| [#39847](https://github.com/openclaw/openclaw/issues/39847) | P1 | 回复中泄露内部元数据（回声污染） | 否，PR [#?](https://github.com/openclaw/openclaw/pulls?q=stripInboundMetadata) 未找到 |
| [#84783](https://github.com/openclaw/openclaw/issues/84783) | P1 | Discord 运行前模型解析延迟 ~30 秒 | 否 |
| [#95658](https://github.com/openclaw/openclaw/issues/95658) | P1 | 升级后 Groq 语音转录彻底中断 | 否 |
| [#96337](https://github.com/openclaw/openclaw/issues/96337) | P1 | Anthropic Vertex native 路由导致纯文本不可见 | 否，需回退配置 |
| [#96335](https://github.com/openclaw/openclaw/issues/96335) | P1 | lossless-claw 引擎因只读插件注册初始化失败 | 否 |
| [#96297](https://github.com/openclaw/openclaw/issues/96297) | P1 | 工具失败时最终回复被抑制 | 是，[#96371](https://github.com/openclaw/openclaw/pull/96371) |
| [#96363](https://github.com/openclaw/openclaw/issues/96363) | P1 | Telegram `/status` 富回复失败 `RICH_MESSAGE_EMAIL_INVALID` | 否 |
| [#96343](https://github.com/openclaw/openclaw/issues/96343) | P1 | 异步命令输出截断丢失登录提示/设备码 | 是，[#96370](https://github.com/openclaw/openclaw/pull/96370) |
| [#96313](https://github.com/openclaw/openclaw/issues/96313) | P2 | Feishu 子账户渲染模式不继承全局设置 | 否 |

此外，多条关于 Telegram 消息重复（[#96242](https://github.com/openclaw/openclaw/issues/96242)）、Web UI 会话重启后上下文丢失（[#96331](https://github.com/openclaw/openclaw/issues/96331)）等 P1 问题也在今日被报告。

## 6. 功能请求与路线图信号

今日新增或活跃的功能请求：

- **[#94285](https://github.com/openclaw/openclaw/issues/94285) — Discord 保留进度草稿并追加最终答案**（P3）  
  用户希望流式模式下不覆盖之前的推理过程，而是追加最终回复，适合偏好看到完整中间步骤的用户。

- **[#94657](https://github.com/openclaw/openclaw/issues/94657) — 将字符制 bootstrap 限制改为基于 token**（P2）  
  社区建议用令牌数代替字符数衡量上下文窗口，该改进已在多个 Thread 中获得支持。

- **[#96350](https://github.com/openclaw/openclaw/issues/96350) — Android 节点系统通知支持 deeplink/contentIntent**（P2）  
  便于点击通知直接跳转至 App 内相关页面。

- **[#20935](https://github.com/openclaw/openclaw/issues/20935) — 代理内存变更审计日志**（P2，已开放 4 个月）  
  社区持续呼吁添加内存修改的只审计日志，用于安全检测与回溯。

结合已有 PR 判断，`#94657`（token 限额）已有相关讨论但无对应 PR；`#96350`（Android 通知 deeplink）可能进入下一版本；`#94285`（Discord 保留草稿）属于体验优化，需维护者评估。

## 7. 用户反馈摘要

从 Issues 评论中提炼的真实用户痛点：

- **升级体验差**：`2026.6.8 → 2026.6.9` 的内存存储静默迁移导致用户丢失嵌入索引，且无任何迁移提示（[#95495](https://github.com/openclaw/openclaw/issues/95495)）。同一版本还破坏了 Groq 语音转录，用户被迫降级（[#95658](https://github.com/openclaw/openclaw/issues/95658)）。
- **MCP 工具配置复杂性高**：即使按照文档设置 `bundle-mcp` 与 allowlist，子代理仍无法使用 MCP 工具，用户认为文档与实际行为不一致（[#85030](https://github.com/openclaw/openclaw/issues/85030)）。
- **Telegram 用户体验不佳**：多步骤任务中，进度编辑会覆盖之前的中间步骤，用户无法回溯（[#96306](https://github.com/openclaw/openclaw/issues/96306)）。`/status` 命令在启用富消息时直接失败（[#96363](https://github.com/openclaw/openclaw/issues/96363)）。
- **Web ChatGPT 集成问题**：登录成功后 OAuth 失败，用户无法正常使用（[#96373](https://github.com/openclaw/openclaw/issues/96373)）。
- **沙盒配置缺失**：文件系统访问限制功能请求（`tools.fileAccess`）已开放近 5 个月，用户在生产环境中担心安全风险（[#7722](https://github.com/openclaw/openclaw/issues/7722)）。

## 8. 待处理积压

以下 Issues/PR 长期未得到充分响应，需提醒维护者关注：

| 编号 | 类型 | 标题 | 创建日期 | 最后更新 | 优先级 |
|------|------|------|----------|----------|--------|
| [#7722](https://github.com/openclaw/openclaw/issues/7722) | Enhancement | Filesystem Sandboxing Config | 2026-02-03 | 2026-06-24 | P2 |
| [#20935](https://github.com/openclaw/openclaw/issues/20935) | Enhancement | Audit log for agent memory changes | 2026-02-19 | 2026-06-24 | P2 |
| [#39847](https://github.com/openclaw/openclaw/issues/39847) | Bug (P1) | Echo contamination (stripInboundMetadata) | 2026-03-08 | 2026-06-24 | P1 |
| [#78493](https://github.com/openclaw/openclaw/issues/78493) | Bug (P1) | `sudo openclaw update` creates mixed ownership | 2026-05-06 | 2026-06-24 | P1 |
| [#69754](https://github.com/openclaw/openclaw/issues/69754) | Bug (P2) | Feishu approval UX unclear fallback | 2026-04-21 | 2026-06-24 | P2 |
| [#65205](https://github.com/openclaw/openclaw/pull/65205) | PR (XL) | feat(discord): canvas-first Discord Activities support | 2026-04-12 | 2026-06-24 | P2 |

这些积压项多涉及安全性与用户核心体验，建议维护团队优先投入资源处理 P1 缺陷（如 `#39847`、`#78493`），并考虑将 `#7722` 文件沙盒纳入近期路线图。

---

## 横向生态对比

# 个人 AI 智能体开源生态横向对比分析报告

**报告时间**：2026-06-24  
**数据来源**：OpenClaw、NanoBot、Hermes Agent、PicoClaw、NanoClaw、NullClaw、IronClaw、LobsterAI、CoPaw、ZeroClaw 等 10 个活跃项目当日动态  
**分析范围**：社区活跃度、技术方向、差异化定位、趋势信号

---

## 1. 生态全景

当前个人 AI 助手与自主智能体开源生态呈现出 **“高活跃、快迭代、强聚焦”** 的态势。核心项目每日处理数十至上百条 Issue/PR，社区贡献意愿强烈，但 **PR 积压普遍** 成为制约合并速度的瓶颈。技术方向上，各项目不约而同地围绕 **模型路由优化、MCP 工具标准、安全权限模型、多平台集成（Slack/Telegram/Feishu）** 展开攻坚；同时，**Token 成本控制** 与 **可观测性建设** 正从个别需求上升为跨项目共识。值得注意的是，多个项目（OpenClaw、ZeroClaw、IronClaw）均报告了 **关键回归或安全漏洞**，表明在快速迭代中稳定性保障仍是最大挑战。总体而言，生态处于 **功能密集拓展与质量基建并行的青春期**，下一阶段有望通过集中发布与安全审计实现成熟度跃升。

---

## 2. 各项目活跃度对比

| 项目 | 当日 Issue 活动 | 当日 PR 活动 | 版本发布 | 健康度评估 |
|------|----------------|-------------|----------|------------|
| **OpenClaw** | 36 条（新开/活跃22，关闭14） | 500 条（待454，合/关46） | ✅ v2026.6.10 | 高度活跃但 PR 积压严重，多回归问题 |
| **NanoBot** | 16 条 | 44 条（合/关21，待23） | 无 | 关键攻坚期，合并率约50%，响应快 |
| **Hermes Agent** | 18 条 | 50 条（待49，关1） | 无 | 社区贡献强烈但零合并，审阅压力大 |
| **PicoClaw** | 2 条（关1，新1） | 17 条（合/关5，待12） | 无 | 健康度良好，待合并队列较长 |
| **NanoClaw** | 1 条（新开） | 12 条（合/关8，待4） | 无 | 贡献活跃，合并效率高 |
| **NullClaw** | 1 条（关闭） | 1 条（更新） | 无 | 低活跃，维护期 |
| **IronClaw** | 14 条 | 48 条（合/关28，待20） | 无 | 高度活跃，但存在基准测试回归（已关闭） |
| **LobsterAI** | 1 条（更新） | 45 条（合/关42，待3） | 无 | 密集收尾，合并率极高（93%），稳定 |
| **CoPaw** | 16 条（新/活11，关5） | 50 条（待40，合/关10） | 无 | 中度活跃，前端严重 Bug 积压 |
| **ZeroClaw** | 12 条 | 50 条（待41，合/关9） | 无 | 非常活跃，安全响应迅速 |
| TinyClaw / Moltis / ZeptoClaw | 0 | 0 | 无 | 无活动，疑似停更或低频维护 |

---

## 3. OpenClaw 在生态中的定位

**优势**：作为生态核心参照项目，OpenClaw 拥有 **最大社区规模**（单日 PR 超 500 条，Issue 超 30 条），功能覆盖面广（多模型路由、MCP 注入、快慢双模式、跨通道支持），版本迭代激进（v2026.6.10 已发布）。其 **模型路由** 与 **快速模式** 设计在同类中领先。

**技术路线差异**：  
- 对比 **Hermes Agent**：Hermes 更侧重 **多代理编排与 ACP 互操作**，而 OpenClaw 以 **个人对话助手** 为核心场景，强调低延迟交互和工具注入。  
- 对比 **ZeroClaw**：ZeroClaw 押注 **WASM 插件安全沙箱** 与 **供应键安全（SLSA）**，技术堆栈更重；OpenClaw 则采用 **插件注册+只读机制**，更易上手但安全边界较弱（如 lossless-claw 初始化失败）。  
- 对比 **NanoBot**：NanoBot 以“超轻量级”为卖点（实际依赖 Python+Node.js 引发争议），目标用户为个人开发者；OpenClaw 依赖程度更高，但提供更完整的平台化能力。

**社区规模**：OpenClaw 的 Issue 和 PR 数量是第二梯队（如 Hermes、ZeroClaw）的 3-5 倍，但其 **PR 积压率高达 90.8%**（454/500），远高于 NanoBot（52%）和 LobsterAI（7%），反映出审阅资源严重不足。OpenClaw 的优势在于品牌认知与用户基数，但 **稳定性口碑因频繁回归受损**（如 Anthropic Vertex 路由、内存存储静默迁移）。

---

## 4. 共同关注的技术方向

以下为过去 24 小时内 **至少两个项目** 集中讨论或修复的热点，按频率排序：

| 技术方向 | 涉及项目 | 具体诉求/表现 |
|----------|----------|--------------|
| **MCP 工具安全与注入** | OpenClaw（#85030）、NanoBot（#4434/#4435）、CoPaw（#5345）、ZeroClaw（#8279） | MCP 工具无法注入子会话、权限配置绕过（enabledTools 失效）、delegate 策略逃逸、自定义提供商 function calling 缺失 |
| **安全权限模型与凭证管理** | OpenClaw（#39847 回声污染）、Hermes（#19566 凭证丢失）、PicoClaw（#3160/#3161 跨站设置拦截）、ZeroClaw（#7623 API Key 泄漏） | 内部元数据泄漏、凭证轮换丢失、请求来源验证、敏感信息替换破坏代码 |
| **Token 成本优化** | Hermes（#6839/#4379）、OpenClaw（上下文压缩）、ZeroClaw（#8065 成本审计） | 73% 固定开销、惰性工具 Schema 加载、压缩引擎钩子、每次调用成本记录 |
| **多平台集成与兼容性** | NanoBot（Telegram 富消息 #4488）、OpenClaw（Telegram/Feishu）、IronClaw（Slack/Telegram）、CoPaw（群聊错发 #5264）、PicoClaw（Windows QQ #3015） | 前端新特性破坏旧客户端、消息格式不一致、隧道端口绑定冲突 |
| **可观测性与调试能力** | IronClaw（#5182）、ZeroClaw（#8065）、OpenClaw（审计日志 #20935） | 托管环境缺少日志、trace_id 关联、内存变更审计 |
| **升级体验与迁移** | OpenClaw（#95495 内存存储迁移）、NanoBot（“轻量级”标签争议 #660） | 静默迁移导致全量重嵌入、文档与实际行为不符 |

---

## 5. 差异化定位分析

| 维度 | OpenClaw | Hermes Agent | ZeroClaw | NanoBot | LobsterAI | CoPaw |
|------|----------|--------------|----------|---------|-----------|-------|
| **核心定位** | 全功能个人 AI 助手 | 多代理编排平台 | 安全优先的智能体框架 | 超轻量级聊天机器人 | 企业级 Agent 网关（网易有道） | 中文社区桌面 Agent |
| **目标用户** | 普通用户 + 开发者 | 开发者 / 运维团队 | 安全敏感型团队 / 企业 | 个人开发者、爱好者 | 企业客户 | 中文用户、桌面端用户 |
| **技术架构** | 模型路由 + 插件系统 | ACP 互操作 + 上下文引擎 | WASM 沙箱 + 供应键安全 | Python+Node.js 混合 | OpenClaw 魔改 + 自有网关 | Tauri 桌面端 + Python后端 |
| **社区规模** | 最大（500 PR/天） | 较大（50 PR/天） | 大（50 PR/天） | 中等（44 PR/天） | 大（45 PR/天） | 中等（50 PR/天） |
| **迭代速度** | 激进（多回归） | 快速但审阅慢 | 快速（安全修复快） | 中等 | 极快（合并率93%） | 中等（重度 Bug 堆积） |
| **差异化亮点** | 对话快速模式、模型路由 | 惰性 schema 加载、ACP 编排 | WASM 插件、SLSA 签名 | PWA 移动端、Kimi Coding | Minimax M3 支持、WeChat 修复 | 桌面端自更新、DataPaw 插件 |
| **主要短板** | PR 积压 90%、升级体验差 | 零合并日、凭证脆弱 | MCP 子进程泄漏（2月未修） | “轻量级”名不副实 | 社区反馈少 | 前端白屏、启动报错 |

---

## 6. 社区热度与成熟度分层

| 阶段 | 项目 | 特征 |
|------|------|------|
| **快速迭代期**（高活跃，功能密集） | OpenClaw、Hermes Agent、ZeroClaw、IronClaw、LobsterAI | 每日 PR 量 40-500，版本发布频繁或即将发布；存在一定回归与 Bug 积压，稳定性波动大 |
| **质量巩固期**（控制合并节奏，修 Bug） | NanoClaw、PicoClaw、CoPaw | 合并率中等，慢性 Bug 较多（如 CoPaw 前端崩溃），但无大规模回归；社区贡献稳定 |
| **维护期**（低活跃，偶有修复） | NullClaw | 单日 1-2 条活动，长期 PR 未合并，核心团队可能转向其他项目 |
| **静默/停滞期** | TinyClaw、Moltis、ZeptoClaw | 24 小时无任何活动，可能已停止维护或处于休眠状态 |

**警示**：CoPaw 虽处于质量巩固期，但其 3 个 P0/P1 前端 Bug 无对应修复 PR，若持续不解决将滑向“低活跃”阶段。

---

## 7. 值得关注的趋势信号

1. **MCP 安全成为首要痛点**：多个项目（ZeroClaw #8279、NanoBot #4434）报告了配置权限绕过或工具注入失败问题。**“安全默认拒绝”将成为各项目下一版本的标配**，预计将出现统一的安全模型 RFC 或最佳实践文档。

2. **Token 成本优化从 feature 走向 infrastructure**：Hermes 的定性分析（73% 固定开销）和 ZeroClaw 的成本审计 PR 表明，**惰性加载、分层注入、上下文压缩** 等机制将内建到框架核心，而非仅作为插件。

3. **可观测性从“可有可无”变为“必须”**：IronClaw #5182（托管日志缺失）和 ZeroClaw #8065（trace_id+成本）表明，用户在生产环境中对 **跨调用追踪、成本核算、故障诊断** 的需求已超过功能本身。

4. **多代理互操作标准（ACP）萌芽**：Hermes #5257 提出通用 ACP 客户端以编排多个 Agent（如 Claude Code、Copilot），这预示着未来 **不同项目间的 Agent 协同** 将成为生态增长点，类似 LLM 的 OpenAI API 标准。

5. **供应链安全开始落地**：ZeroClaw #8177（SLSA 证明）和 #8129（校验和签名）是首个将 **SLSA 级别** 引入个人 AI 智能体开源项目的尝试，可能成为企业选型的门槛。

6. **移动端与桌面端的“跨端一致性”挑战**：NanoBot 的 PWA 支持与 CoPaw 的 Tauri 自更新代表了不同路线，但 **Telegram Web 客户端不兼容**（NanoBot #4488）和 **桌面端内存占用 1.4G**（CoPaw #5441）说明端侧体验仍粗糙。

7. **“轻量级”定调引发认知博弈**：NanoBot #660 的长期争议揭示出 **用户对项目定位的敏感**——社区期望真正的单二进制或低依赖部署，而非“语义上的轻量”。这可能导致更多项目如 PicoClaw 那样采用 Go 编译原生二进制来响应。

**对开发者的建议**：若您是个人开发者，可优先选择 **NanoBot**（快速上手）或 **PicoClaw**（Go 原生）；若需搭建生产级多代理系统，关注 **Hermes** 的 ACP 进展与 **ZeroClaw** 的安全实践；若重视稳定性，**LobsterAI** 的高合并率和低 Bug 率值得观察。无论选择哪个项目，**MCP 安全配置** 与 **Token 成本监控** 应作为初期落地的必选项。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

好的，作为一名 AI 智能体与个人 AI 助手领域开源项目分析师，我将基于您提供的 2026-06-24 的 NanoBot 项目数据，为您生成一份客观、数据驱动的项目日报。

---

## NanoBot 项目日报 (2026-06-24)

### 1. 今日速览

今日 NanoBot 项目进入了一个**关键的技术攻坚与功能收割期**。PR 活动量（44条）远超 Issue 活动量（16条），且 PR 合并/关闭率接近 50%，显示出团队正在加速将新功能与修复代码合并进主线。社区反馈集中在两大痛点：**Telegram 渠道兼容性**（特别是新版富消息与旧版Web客户端的冲突）和 **MCP 安全模型的绕过风险**。WebUI 在 PWA 支持和移动端体验方面取得了显著进展，这将是项目向更广泛用户群体渗透的重要一步。整体来看，项目活跃度较高，正处于快速迭代、解决关键积压问题并拓展生态的阶段。

### 2. 版本发布

无。

### 3. 项目进展

今日共有 **21 条 PR 被合并或关闭**，标志着多项重要功能与修复已进入主线。

-   **WebUI 体验重大升级：** `zpljd258` 贡献的 PWA 支持（#4457, #4480）已合并，将允许用户在移动设备主屏幕上安装应用，并支持离线缓存。同时，配套的侧边栏滑动手势（#4480）也提升了移动端原生操作感。
-   **AI 提供商生态扩展：** 项目新增了三家重要 AI 提供商支持：
    -   **Kimi Coding (kimi_coding)** (#4463, #4464)：为 Kimi 订阅用户提供了专门的编码计划接入。
    -   **OpenCode Zen & Go** (#4475, #4476)：为开发者提供了更灵活、更具性价比的编码模型选择。
-   **渠道集成拓展：** 新增了 Mattermost 频道支持（#4459），进一步丰富了企业级通信场景。
-   **核心稳定性修复：**
    -   **DuckDuckGo 搜索代理问题** (#4484)：修复了在需要代理的网络环境下搜索功能完全不可用的 bug。
    -   **会话归档/心跳问题** (#4468)：修复了心跳检测会包含已归档密钥，以及会话时间戳缺失导致潜在错误的问题。
    -   **配对存储类型转换** (#4433)：修复了发送者 ID 类型不一致导致配对功能失效的 bug。
-   **Dream 功能增强：** `axelray-dev` 修复了 Dream 功能重复创建技能的问题（#4469），使其能正确地将改进合并到现有技能中。

### 4. 社区热点

-   **Telegram 富消息兼容性危机 (Issue #4488 & PR #4489)**
    -   **链接:** [Issue #4488](https://github.com/HKUDS/nanobot/issues/4488), [PR #4489](https://github.com/HKUDS/nanobot/pull/4489)
    -   **分析:** 这是今日最火热的话题。为了提升消息体验，项目实现了 Telegram Bot API 10.1 的富消息功能（#4413）。然而，这导致在 Telegram Web 端（可能因版本较老）显示为“此消息不受支持”，引发用户强烈不满。用户`chengyongru`提交了 bug 报告，开发者`axelray-dev`迅速提出了一个修复 PR，引入开关配置来禁用此功能，以兼容老版本客户端。这场讨论凸显了在快速采纳新特性时，保持向后兼容性的重要性。

-   **“超轻量级”标签的争议 (Issue #660)**
    -   **链接:** [Issue #660](https://github.com/HKUDS/nanobot/issues/660)
    -   **分析:** 尽管此 Issue 已开启 4 个月，但它仍然获得了最多的评论（11条）和点赞（5个）。用户`besoeasy`指出项目自称“超轻量级”，但实际部署却需要同时安装 Python 和 Node.js，这与承诺相悖。此问题直指项目核心品牌的诚信度和入门门槛问题。社区对“轻量级”的定义有不同理解，但此 Issue 的持久热度表明，简化依赖或更清晰地陈述“轻量级”的适用范围是许多潜在用户的深层诉求。

### 5. Bug 与稳定性

按严重程度排列：

-   **严重：MCP 安全策略绕过 (Issues #4434, #4435)**
    -   **链接:** [#4434](https://github.com/HKUDS/nanobot/issues/4434), [#4435](https://github.com/HKUDS/nanobot/issues/4435)
    -   **描述:** 用户`YLChen-007`报告了两个严重的安全问题，指出 `enabledTools` 配置项配置为空数组以拒绝所有工具时，无法生效，导致 MCP 服务器的资源和提示词能力仍可被模型调用。这是一个配置权限绕过漏洞，可能导致未授权的数据访问。**状态：** 已报告，待处理。

-   **较高：Telegram Web 消息不兼容 (Issue #4488)**
    -   **链接:** [Issue #4488](https://github.com/HKUDS/nanobot/issues/4488)
    -   **描述:** 发送的消息在 Telegram Web 客户端上显示“不支持”，但在手机APP上正常。这是一个直接的回归问题。**状态：** 已有修复 PR (#4489)，待审核合并。

-   **中等：WebUI 显示/交互 Bug**
    -   **WebUI 多个文件编辑不保留 (PR #4487):** `chengyongru`提出修复，使一次 `apply_patch` 对多个文件的编辑记录能正确保留，而不是合并。
    -   **`<thinking>` 标签渲染问题 (Issue #4465):** 修复已合并 (PR #4466)，现在会正确地将其渲染为推理过程块，而不是纯文本。

-   **中等：Telegram 消息排版 Bug (Issue #4470)**
    -   **链接:** [#4470](https://github.com/HKUDS/nanobot/issues/4470)
    -   **描述:** 报告了两个问题：换行符被忽略，以及消息在生成过程中被反复编辑导致视觉闪烁。**状态：** 已关闭，说明问题已被修复或已知并接受。

### 6. 功能请求与路线图信号

-   **PWA 与移动端优化 (已实现):** PWA 支持和侧边栏手势（#4479, #4480）几乎是“即提即合”，反映了项目组对改善移动端用户体验的极高优先级。这极有可能在下一版本发布中作为核心特性推出。

-   **MCP Security 审计与强化 (已识别):** 安全绕过报告 (#4434, #4435) 将迫使项目组重新审视 MCP 权限模型。未来可能会看到更严格的“默认拒绝”策略、更清晰的配置文档，甚至引入“允许的资源/提示”白名单配置项，这与 `enabledTools` 形成对称。

-   **只读会话 (PR #4271):** `DreamShepherd2006` 提出的为 pinned 会话（如 FAQ）提供只读模式、避免 LLM 处理的 PR，虽然打开已两周，但潜在价值巨大。若被采纳，将能大幅提升平台构建者在特定场景下的效率和安全性，是向“平台化”发展的重要信号。

-   **OpenAI API 认证 (Issue #4490):** 用户 `xiaweiwei67-stack` 提出的为 OpenAI 兼容 API 增加认证的请求，与 WebSocket 网关已有的策略对齐。这表明社区对将 NanoBot API 暴露在公网的需求正在增加，安全配置的完善是下一阶段的必然要求。

### 7. 用户反馈摘要

-   **正面反馈：** 对于 Kimi Coding 和 OpenCode 等新提供商的加入，社区反应热烈，说明开发者对新模型和新服务有持续的需求。PWA 和侧边栏手势的合并也获得了积极的评价，用户对“原生应用”体验感到满意。

-   **负面反馈/痛点：**
    1.  **“轻量级”认知偏差：** 用户对项目的“超轻量级”定位有质疑，认为 Node.js 依赖违背了初衷（#660）。
    2.  **向前兼容性风险：** 富消息功能的引入破坏了部分用户的 Telegram Web 体验，引发了信任危机。用户更喜欢稳定的“开关”来控制新功能（#4488）。
    3.  **安全担忧：** MCP 权限配置失效的问题（#4434, #4435）引起了使用 MCP 的用户的高度警觉，他们希望项目核心的安全机制是绝对可靠的。
    4.  **小型模型体验不佳：** 用户报告在使用本地模型时频繁陷入无限的工具调用循环（#2298），这指出了项目在处理能力不足的模型时，缺乏足够的“防错”和“补救”逻辑。

### 8. 待处理积压

-   **长期未合并的重要 PR：**
    -   **Zalo 频道重构 (PR #2078):** 自 3 月 16 日开启，已搁置超过 3 个月。作为越南市场的重要渠道，其合并进程应被关注。
    -   **Agent 评估框架 (PR #2283):** 自 3 月 20 日开启，对保障代码质量和 CI 流程至关重要，但迟迟未被合并。
    -   **QQ 频道音频支持 (PR #2316):** 同样是自 3 月开启，这对于希望开通无障碍或语音交互的用户是必备功能。
-   **长期未解决的 Issue：**
    -   **“超轻量级”争议 (Issue #660):** 已存在 4 个月，是高赞问题。虽然讨论活跃，但缺乏项目组的官方回应或处理方案，有冷处理的风险。维护者应考虑给出明确答复，例如调整文档措辞或将 Node.js 依赖设为可选。

建议维护者团队在接下来的开发周期中，优先排序以下事项：1）**合并并发布** MCP 安全绕过问题的修复补丁；2）**审查并合并**如 Agent 评估、Zalo、QQ 等长期搁置的 PR；3）**正式回应** “轻量级” 争议 (Issue #660) 和 “无限循环” 问题 (Issue #2298) 的改进计划。

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent 项目动态日报 | 2026-06-24

---

## 1. 今日速览

过去 24 小时内，Hermes Agent 项目保持 **高度活跃**：社区提交了 **18 个新 Issue** 和 **50 个 Pull Request**，其中 49 个 PR 仍待合并、1 个 PR 被关闭（后被新 PR 取代）。**零合并** 表明维护团队当前审阅压力较大，需加快积压 PR 处理。核心关注点集中在 **Token 开销优化**、**多平台集成**（语音、Zulip）、**稳定性 Bug 修复**（credential 丢失、sandbox 隔离、Telegram 流处理）以及 **安全加固**。项目整体健康度良好，社区贡献意愿强烈，但需平衡维护速度与质量。

---

## 2. 版本发布

**无新版本发布。**

---

## 3. 项目进展

今日 **没有 PR 被合并**。唯一关闭的 PR 是：

- **#47330** `[CLOSED] feat(voice): real-time voice conversation platform (Daily + Deepgram Flux + Cartesia)` [链接](https://github.com/NousResearch/hermes-agent/pull/47330)  
  该 PR 在关闭后，其变化立即被新 PR **#51827** 以更干净的分支替代，确保功能持续。  
  → **项目在语音平台方向实际有推进，但未进入主线。**

其余所有开放 PR 均处于待审状态，涉及：
- 桌面端 slash 命令启用（#51768）
- 凭证轮换优化（#51821）
- 压缩引擎钩子（#51226）
- 安全凭证共享（#51832）
- 多个 Bug 修复（见第 5 节）

**整体评价**：项目今日未向前迈出实质性合并步骤，但社区贡献的广度和深度均显著提升。

---

## 4. 社区热点

以下 Issue/PR 讨论最为活跃、反响最大：

| 编号 | 标题 | 类型 | 评论数 | 👍 数 | 核心诉求 |
|------|------|------|--------|-------|----------|
| [#6839](https://github.com/NousResearch/hermes-agent/issues/6839) | Lazy Tool Schema Loading — Two-Pass Tool Injection | Feature | **26** | **14** | 减少每次 API 调用的 token 浪费（当前 50+ 工具 schema 全量注入） |
| [#5257](https://github.com/NousResearch/hermes-agent/issues/5257) | Generalized ACP client for multi-agent CLI orchestration | Feature | **11** | **16** | 让 Hermes 能编排 Claude Code、Copilot 等 ACP 兼容代理 |
| [#4379](https://github.com/NousResearch/hermes-agent/issues/4379) | Token overhead analysis: 73% of each API call is fixed overhead | Perf | **15** | 0 | 量化分析显示 73% token 为固定开销，要求系统级压缩 |
| [#13834](https://github.com/NousResearch/hermes-agent/issues/13834) | Hermes openai-codex fails on same machine where official Codex CLI works | Bug | **12** | **3** | 用户反馈官方 Codex 正常但 Hermes 集成失败，影响使用信心 |
| [#39691](https://github.com/NousResearch/hermes-agent/issues/39691) | integrate headroom-ai for tool output compression | Feature | 7 | **10** | 引入外部压缩模型优化上下文管理 |

**分析**：社区对 **Token 效率** 和 **多代理互操作** 的呼声最高。两个 Token 相关 Issue（#6839 和 #4379）分别从 schema 加载和总体开销角度提出优化方案，且均获得高赞。ACP 泛化（#5257）则反映了用户希望 Hermes 成为“代理中心”的愿景。

---

## 5. Bug 与稳定性

今日报告 18 个 Issue 中大部分为 Bug，按严重程度排列（附对应修复 PR，如有）：

**P1（高优先级）**
- [#19566](https://github.com/NousResearch/hermes-agent/issues/19566) [Bug] OpenAI-Codex credential pool can drop newly added credential after stale auth.json rewrite  
  → 凭证轮换时新建凭证丢失  
- [#51826](https://github.com/NousResearch/hermes-agent/issues/51826) tirith auto-install retries unboundedly and leaks temp dirs, exhausting disk  
  → 对应 fix PR [#51831](https://github.com/NousResearch/hermes-agent/pull/51831)

**P2（中等）**
- [#33801](https://github.com/NousResearch/hermes-agent/issues/33801) Secret redaction corrupts code syntax in tool output (write_file, execute_code, terminal)  
  → 敏感信息替换破坏 Python/Shell 语法  
- [#48056](https://github.com/NousResearch/hermes-agent/issues/48056) Telegram DM topic cron delivery falls back out of topic  
- [#51691](https://github.com/NousResearch/hermes-agent/issues/51691) skill_view returns UTF-8 decode error on Windows Desktop (cp936)  
  → 对应 fix PR [#51823](https://github.com/NousResearch/hermes-agent/pull/51823)  
- [#51800](https://github.com/NousResearch/hermes-agent/issues/51800) Context compression tail budget ignores reasoning_content, causing near-no-op compaction  
  → 对应 fix PR [#51822](https://github.com/NousResearch/hermes-agent/pull/51822)  
- [#51828](https://github.com/NousResearch/hermes-agent/issues/51828) Telegram streaming truncation triggers full response re-generation with reformulated text  
- [#51820](https://github.com/NousResearch/hermes-agent/issues/51820) execute_code sandbox: generated files invisible to host filesystem  
- [#51825](https://github.com/NousResearch/hermes-agent/issues/51825) All skills run on a single CPU core, cannot utilize multiple cores  
- [#51833](https://github.com/NousResearch/hermes-agent/issues/51833) Model dropdown shows reasoning effort as part of model name  
- [#51829](https://github.com/NousResearch/hermes-agent/issues/51829) /learn slash command shows ack but does not trigger LLM in Desktop GUI  
- [#50663](https://github.com/NousResearch/hermes-agent/issues/50663) z.ai peak hours rate limiting Hermes  
- [#51824](https://github.com/NousResearch/hermes-agent/issues/51824) Path traversal vulnerability in pet thumbnail slug (fix PR #51824)

**P3（低优先级）**
- [#51639](https://github.com/NousResearch/hermes-agent/issues/51639) 等（功能类）

**总结**：今日 Bug 数量多且覆盖面广，从凭证管理到 sandbox 隔离、从压缩逻辑到 GUI 响应，均需维护者及时响应。庆幸的是大部分重要 Bug 已有社区提交的修复 PR。

---

## 6. 功能请求与路线图信号

用户提出的新功能需求及可能纳入下一版本的候选：

| Issue/PR | 描述 | 优先级 | 社区热度 | 路线图信号 |
|----------|------|--------|----------|------------|
| [#6839](https://github.com/NousResearch/hermes-agent/issues/6839) | 惰性工具 Schema 加载 | P3 | ★★★★★ | 可能成为 v0.18 性能优化核心 |
| [#5257](https://github.com/NousResearch/hermes-agent/issues/5257) | 通用 ACP 客户端（编排多代理） | P3 | ★★★★★ | 反映多代理编排趋势 |
| [#39691](https://github.com/NousResearch/hermes-agent/issues/39691) | 集成 headroom-ai 工具输出压缩 | P3 | ★★★★ | 与 token 优化高度协同 |
| [#51816](https://github.com/NousResearch/hermes-agent/issues/51816) | Plugin SDK: 增加 profileScope 表面 | P3 | ★★ | 扩展插件能力 |
| [#51832](https://github.com/NousResearch/hermes-agent/pull/51832) | Valut: 安全凭证共享（`?/.../?` 语法） | P3 | ★★★ | 企业级安全增强 |
| [#51827](https://github.com/NousResearch/hermes-agent/pull/51827) | 实时语音会话平台（Daily+Deepgram+Cartesia） | P3 | ★★★ | 多模态交互 |
| [#51226](https://github.com/NousResearch/hermes-agent/pull/51226) | 上下文引擎钩子 select_context() + on_turn_complete() | P3 | ★★★ | 智能上下文管理 |

**判断**：Token 效率相关功能（#6839、#39691）最可能被优先采纳；多代理编排（#5257）和实时语音（#51827）则体现了 Hermes 向“全平台智能体中心”演进的方向。

---

## 7. 用户反馈摘要

从 Issue 评论中提炼的典型用户痛点和场景：

- **Token 浪费严重**：“每次调用固定消耗约 3500-5000 tokens，即使对话仅需一个简单文件操作”（#6839）。用户 Bichev 通过监控仪表盘发现 73% 的 API 调用为固定开销（#4379），强烈要求引入惰加载或分层注入。
- **凭证管理脆弱**：“新加的 credential 在轮换后神秘消失”（#19566），用户不得不频繁手动检查 `auth.json`，信任度下降。
- **平台兼容性问题**：官方 Codex CLI 在相同机器正常，但 Hermes 集成的 openai-codex 反复失败（#13834），用户对集成质量提出质疑。Windows 用户遭遇编码问题（#51691）和单核性能瓶颈（#51825）。
- **工具输出与安全矛盾**：secret redaction 导致代码语法错误（#33801），用户认为“安全功能反而破坏了可用性”。
- **Telegram 流处理**：长回复被截断后系统重新生成整个响应（#51828），用户吐槽“浪费 tokens 且破坏对话连续性”。
- **社区积极贡献**：多位用户不仅报告问题，还自行开发修复 PR（如 #51822、#51823、#51831），体现了社区的高度参与度。

---

## 8. 待处理积压

以下为长时间未得到官方回应的关键 Issue/PR，建议维护者优先关注：

| 编号 | 类型 | 标题 | 创建时间 | 最后更新 | 重要性 |
|------|------|------|----------|----------|--------|
| [#3335](https://github.com/NousResearch/hermes-agent/pull/3335) | PR | feat(gateway): add Zulip integration | 2026-03-27 | 2026-06-24 | 重要特性，3个月未合并 |
| [#8427](https://github.com/NousResearch/hermes-agent/pull/8427) | PR | feat(vertex): add Vertex AI provider for Gemini | 2026-04-12 | 2026-06-24 | 企业级需求，2个月未审 |
| [#27829](https://github.com/NousResearch/hermes-agent/pull/27829) | PR | fix: use target_model for Bedrock api_mode routing | 2026-05-18 | 2026-06-24 | 功能性修复，1个月未合 |
| [#51226](https://github.com/NousResearch/hermes-agent/pull/51226) | PR | feat(context-engine): add select_context() hooks | 2026-06-23 | 今日 | 新功能，需评审 |
| [#51816](https://github.com/NousResearch/hermes-agent/issues/51816) | Issue | Plugin SDK: additive profileScope surface | 今日 | 今日 | 社区提出的 SDK 改进建议 |

**特别提醒**：Zulip 集成（#3335）与 Vertex AI 支持（#8427）均为跨平台扩展关键，积压过久可能影响社区开发者积极性。建议维护团队在未来一周内至少安排一次集中 PR 审阅活动。

---

*日报生成基于 Hermes Agent GitHub 仓库 2026-06-24 实时数据。数据来源：issues, pulls, releases.*

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

## PicoClaw 项目动态日报 — 2026-06-24

### 今日速览
项目今日保持高活跃度，共处理 **2 条 Issue**（关闭 1 条、新开 1 条）和 **17 条 PR**（合并/关闭 5 条，待合并 12 条）。`openai_compat` 模块是今日修复焦点：两项修复性 PR（#3165、#3166）针对 Doubao Seed 模型的工具调用 XML 泄漏和编译错误，其中昨日提交的应急修复（#3154）已被合并。此外，安全加固（跨站设置请求拦截、执行权限 deny 模式修正）和平台兼容性（Windows 路径处理、Android 崩溃）相关 PR 处于待审状态。整体健康度良好，但待合并队列较长（12 条），需关注合并节奏。

### 项目进展
今日合并/关闭的 **5 条重要 PR** 为项目修复了以下问题：

- **API 兼容性**：[#3154](https://github.com/sipeed/picoclaw/pull/3154) — 修复 OpenAI 兼容接口中 Doubao Seed 模型泄漏的 `<seed:tool_call>` XML，由 `hanZeng-08` 提交并已合并。
- **代码质量**：[#3059](https://github.com/sipeed/picoclaw/pull/3059) — 显式忽略错误路径中的 `Close()` 返回值，消除 linter 警告；[#3054](https://github.com/sipeed/picoclaw/pull/3054) — 为 `LINEChannel.Send` 中 `sync.Map` 类型断言添加 `ok` 检查，防止 panic。
- **前端/会话管理**：[#3047](https://github.com/sipeed/picoclaw/pull/3047) — 恢复会话详情页的完整 JSONL 历史记录，修复归档消息不可见问题。
- **工具配置**：[#2888](https://github.com/sipeed/picoclaw/pull/2888) — 修复工具配置加载图像反应问题（合并自外部贡献者）。

这些合并提升了 OpenAI 兼容性、代码健壮性和会话数据完整性，项目向前迈进了一个小版本迭代。

### 社区热点
今日讨论最活跃的是 **[#3015](https://github.com/sipeed/picoclaw/issues/3015)**（已关闭），关于 Windows 下 QQ 频道连接失败的 Bug 报告，有 **4 条评论**。该 Issue 创建于 6 月 6 日，被标记为 stale 后于今日关闭。用户反馈了 Release 构建中 `picoclaw gateway` 在 Windows 上获取 token 超时的具体场景，Pico 频道不受影响。该问题虽已关闭，但未看到明确的修复 PR 合并，可能需后续验证。

另外，今日新开 **[#3164](https://github.com/sipeed/picoclaw/issues/3164)** 报告了 Android/Termux 上 Process Hooks 导致网关 2 秒内崩溃的问题（v0.2.9，config v3），虽暂无评论但值得关注。

### Bug 与稳定性
按严重程度排列今日报告及修复的 Bug：

| 严重度 | Bug 描述 | 状态 | 是否已有 Fix PR |
|--------|----------|------|----------------|
| **高** | [Android/Termux] Process Hooks（JSON-RPC over stdio）导致网关启动 2 秒内崩溃。即使是 minimal “hello world” hook 也会触发 [#3164](https://github.com/sipeed/picoclaw/issues/3164) | 开放，无评论 | 暂无 |
| **中** | [openai_compat] Doubao Seed 模型的工具调用 XML 泄露到 `message.content`，影响流式输出与最终提取 [#3153](https://github.com/sipeed/picoclaw/issues/3153)（通过 #3154 修复） | 已合并 | ✅ #3154 |
| **低** | [构建] `pkg/providers/openai_compat` 中 `log.Printf` 导致 `undefined: log` 编译失败 [#3166](https://github.com/sipeed/picoclaw/pull/3166) | 待合并 | ✅ #3166（同行 #3165 也涉及） |
| **低** | [Windows] QQ 频道连接失败（token 获取超时） [#3015](https://github.com/sipeed/picoclaw/issues/3015) | 已关闭（stale） | 未确认修复 |

### 功能请求与路线图信号
今日提交的功能性 PR 暗示了下一版本的几个可能方向：

1. **Remote WebSocket 模式**：[#3118](https://github.com/sipeed/picoclaw/pull/3118) — 为 `picoclaw agent` 添加 `--remote` 参数，支持通过 WebSocket 远程控制。该 PR 已存在 12 天，仍在审查。
2. **Bedrock Converse Prompt Caching**：[#3163](https://github.com/sipeed/picoclaw/pull/3163) — 利用 AWS Bedrock 的启动缓存功能，通过显式缓存点降低输入 Token 成本（约 0.1× 输入计费）。
3. **Telegram 群组回复视为提及**：[#2975](https://github.com/sipeed/picoclaw/pull/2975) — 在 `mention_only: true` 配置下，回复机器人消息与 @提及等效，提升 Telegram 群聊交互友好性。
4. **安全加固**：[#3160](https://github.com/sipeed/picoclaw/pull/3160) — 拒绝跨站源设置的 launcher 密码请求，增加 `Sec-Fetch-Site`、`Origin` 检查。 [#3161](https://github.com/sipeed/picoclaw/pull/3161) — 即使在自定义允许规则下也保持 deny 模式活跃，防止 `jq` 等工具读取环境变量。

以上 PR 均已开放待合并，其中 #3118 和 #3163 属于架构级增强，可能纳入下一个中版本；#2975 和 #3160/#3161 属于体验改进与安全补丁，有望快速合入。

### 用户反馈摘要
从 Issue 评论和 PR 描述中萃取的真实用户痛点：

- **Windows 平台体验**：QQ 频道通道在 Windows Release 构建下完全不可用，而 Pico 通道正常。用户 `cuandada` 详细复现了 token 获取超时步骤（#3015）。
- **Android 移动端受阻**：用户 `AMEOBIUS` 在 Android/Termux 上尝试使用 Process Hooks 功能失败，即使是空 hook 也会使网关崩溃，导致无法在移动端部署任何自定义逻辑（#3164）。
- **流式输出数据泄漏**：使用 Doubao Seed 模型的用户可能看到 `<seed:tool_call>` 等 XML 片段裸露在聊天界面，影响观感（#3153 已修复）。
- **会话历史丢失**：用户反馈当消息被 `meta.Skip` 标记后，会话详情页无法看到归档历史（#3047 已修复）。

### 待处理积压
以下开放 Issue/PR 长期未响应或存在合并障碍，提醒维护者关注：

- **[#2975](https://github.com/sipeed/picoclaw/pull/2975)**（feat: telegram → treat reply as mention）—— 提交于 2026-05-30，已开放 25 天，属于功能增强，但无最新讨论。
- **[#3118](https://github.com/sipeed/picoclaw/pull/3118)**（Add remote WebSocket mode）—— 提交于 2026-06-12，开放 12 天，涉及较大改动，可能需要架构评审。
- **[#3104](https://github.com/sipeed/picoclaw/pull/3104) / #3103 / #3100**（Dependabot 依赖更新）—— 三个 PR 均已开放 13 天，涉及前端依赖 `shadcn`、`typescript-eslint`、`@vitejs/plugin-react`，建议定期统一合并以避免安全风险。

另外，已关闭的 [#3015](https://github.com/sipeed/picoclaw/issues/3015) 作为 stale 关闭，但底层 Windows QQ 通道问题未看到明确修复，建议在后续版本中重新评估并纳入测试矩阵。

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

# NanoClaw 项目日报 | 2026‑06‑24

---

## 1. 今日速览

过去24小时内，NanoClaw 合并/关闭了8个 Pull Request，另有4个待合并，整体贡献活跃度较高。核心进展包括：**Chat SDK 全线升级至 4.29.0**（主分支、channels 分支、providers 分支同步推进），**Slack Socket Mode 功能正式落地**，**update-skills 流程得到强化修复**，以及**通用扩展点机制**被引入。社区唯一新开的 Issue 报告了 Slack 安装向导中的端口绑定冲突，值得注意。项目依赖更新和架构改进同天完成，健康度良好。

---

## 2. 版本发布

无新版本发布（最新 Releases 无数据）。

---

## 3. 项目进展

### 已合并/关闭的重要 PR（8 个）

| PR | 标题 | 关键影响 |
|----|------|----------|
| #2834 | `chore(deps): move chat SDK + channel-adapter pins to 4.29.0` | 主分支安装表面全面迁移至 Chat SDK 4.29.0，修复了版本锁定导致的类型不匹配问题 |
| #2835 | `chore(deps): bump @chat-adapter/* + chat to 4.29.0 (channels)` | 配套 channels 分支升级，保持与主分支一致 |
| #2836 | `chore(deps): bump chat SDK to 4.29.0 (providers)` | 同步升级 providers 分支，完成全系统 SDK 版本对齐 |
| #2837 | `feat(slack): Socket Mode — adapter + guided setup (SLACK_APP_TOKEN)` | **新增 Slack Socket Mode**：通过设置 `SLACK_APP_TOKEN` 让 bot 使用 WebSocket 出站连接，无需公开端点，适合本地开发或 NAT 环境 |
| #2839 | `chore(channels): bring Slack Socket Mode (#2837) into channels` | 将 Socket Mode 适配器合并进 channels 分支，确保 channels 用户也能使用该特性 |
| #2826 | `fix(update-skills): nudge into skill updates, rebuild container on re-apply` | 修复 `/update-nanoclaw` 中技能更新提示被误标为可跳过的问题，并增加容器重建逻辑，避免用户遗漏上游重要修复 |
| #2841 | `Generic inert extension-point seams (registerX/applyX) across host + container runtime` | 引入一组通用惰性扩展点（注册/应用对），下游 fork 无需修改上游核心代码即可插入自定义行为，且无注册时行为完全不变 |
| #2833 | `Feat/hook surface guard` | 新增钩子表面守卫，提升扩展点安全性 |

### 待合并的 PR（4 个）

| PR | 标题 | 预期功能 |
|----|------|----------|
| #2843 | `feat: add /learn skill — distill or refine a reusable skill from anything` | 新增 `/learn` 技能，可从目录、URL 等源头蒸馏出可复用的技能 |
| #2842 | `Generic inert extension-point seams + reserve built-in MCP server names` | 与 #2841 类似但增加对内置 MCP 服务器名称的预留保护 |
| #2832 | `feat(approvals): reject with reason` | 为模块审批卡片增加“拒绝并附理由”功能，让请求智能体能够根据理由调整行为 |
| #2838 | `feat(providers): add Manifest model router provider` | 新增 Manifest 模型路由 provider，扩展 provider 生态 |

**整体评估**：项目在依赖版本一致性、Slack 集成实用性、技能更新可靠性以及架构扩展性方面均有显著推进。Chat SDK 的全线升级消除了多个分支间的版本碎片，是当日最核心的工程改进。

---

## 4. 社区热点

### 讨论最活跃的 PR

- **#2832** `feat(approvals): reject with reason`（待合并）  
  该 PR 虽无大量评论，但**功能设计直接回应了用户对审批机制缺乏反馈的长期呼声**。请求智能体目前只能收到“已拒绝”，无法获知拒绝原因，导致无法自适应调整。该特性将显著改善人机协作体验。

### 唯一新 Issue

- **#2840** `Nanoclaw v2 binds port 3000 of external host ip for slack`  
  无评论无点赞，但性质较敏感：Slack 安装指南建议创建隧道到端口 3000，但该端口已被绑定在外部 IP 上，使隧道失去安全意义。可能影响新用户首次配置 Slack 的体验。

---

## 5. Bug 与稳定性

### 新报告 Bug（1 个，中优先级）

| Issue | 描述 | 严重程度 | Fix PR 状态 |
|-------|------|----------|-------------|
| #2840 | NanoClaw v2 在安装 Slack 时占用外部主机 IP 的 3000 端口，导致隧道安全目的失效 | 中 – 影响 Slack 集成初始设置的正确性与安全性 | 暂无 |

### 当日已修复的 Bug / 稳定性改进

- **#2826** 修复了技能更新提示被误判为“可选跳过”的问题，并增加了容器重建逻辑，避免用户遗漏重要上游修复。该问题可能导致已安装技能与新版本不兼容，属于**影响用户实际使用**的中等稳定缺陷。

---

## 6. 功能请求与路线图信号

### 可能纳入下一版本的功能

- **`/learn` 技能（#2843）**：允许从任意来源（目录、URL、历史对话）提取并精炼出可复用的技能。该功能高度符合“个人 AI 助手”场景，预计将进入下一个次要版本。
- **审批拒绝理由（#2832）**：请求智能体在收到拒绝后能读取原因并适应，将显著提升自动化工作流的人机协同能力。
- **通用惰性扩展点（#2841/#2842）**：已部分合并（#2841），#2842 补充了 MCP 名称保留。该架构改动虽不直接对用户可见，但为第三方插件生态奠定了基础，属于**基础设施路线图**信号。
- **Manifest 模型路由 provider（#2838）**：新增 provider 类型，扩展了模型路由选择，暗示项目在支持更多模型后端方面持续投入。

### 用户诉求信号

Issue #2840 虽未直接提出功能请求，但反映了**用户对集成 Wizard 步骤精确性**的期望：向导输出应与实际运行时行为完全一致，否则会造成信任度下降。

---

## 7. 用户反馈摘要

基于当前数据（评论有限），可提炼以下几点：

- **满意度**：无明确负面评论，PR 合并效率较高（8/12 合并），表明协作流程顺畅。
- **痛点**：Issue #2840 显示 Slack 安装向导存在端口冲突，首次使用的用户可能因此遇到排查困难。该问题尚未被回复。
- **使用场景**：PR #2826 修复了 `/update-nanoclaw` 流程，表明用户正在依赖该命令进行版本更新；#2837 的 Slack Socket Mode 为需要本地开发或无法暴露端口的用户提供**关键替代方案**。
- **待验证反馈**：PR #2832 的“拒绝理由”功能若上线，预计会获得正面评价，因为更丰富的审批反馈是许多团队在自动化流程中反复提及的需求。

---

## 8. 待处理积压

### 需维护者关注的重要开放 Issue

| 编号 | 标题 | 创建时间 | 最后更新 | 备注 |
|------|------|----------|----------|------|
| #2840 | Nanoclaw v2 binds port 3000 of external host ip for slack | 2026‑06‑23 | 2026‑06‑23 | 新报问题，无评论，无 PR。建议尽快确认复现并给出解决方案说明 |

### 长期未响应的 PR（当日暂无明显积压）

所有 PR 均于近 3 天内创建或更新，无长时间停滞。其中 #2832（6月22日创建）和 #2842（6月23日创建）待合并，但尚在合理周期内。

---

**日报生成时间**：2026‑06‑24  
**数据来源**：GitHub `nanocoai/nanoclaw`（真实项目名 NanoClaw）  
*本日报由 AI 分析师自动生成，仅供参考。*

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw 项目动态日报（2026-06-24）

---

## 1. 今日速览

过去 24 小时内，NullClaw 项目保持低活跃度：仅关闭 1 个 Bug 类 Issue（#967），并收到 1 个长期存在的 PR（#783）的更新，但该 PR 仍未合并。无新版本发布。整体来看，项目处于维护期，社区贡献节奏放缓，但关键 Bug 得到响应和关闭，稳定性有所提升。

## 2. 版本发布

无新版本发布（上一版本为 v2026.5.29）。

---

## 3. 项目进展

- **Bug 修复关闭**：Issue #967 报告了在 Windows 11 上使用 `nullclaw agent` 命令时出现 `error: NoResponseContent` 的错误，触发频率超过 50%。该 Issue 已于 2026-06-23 被关闭。虽无关联 PR 明确说明修复方式，但关闭表明维护者已定位或处理该问题，用户可关注后续版本确认修复效果。
- **长期 PR 仍在活跃**：PR #783（feat: cron subagent, run history, JSON output, security hardening）已存在两个半月，作者于 2026-06-23 再次更新，显示功能仍在打磨。该 PR 涉及较重大新特性，若能合并将为项目带来定时任务引擎、JSON CLI 输出增强及安全加固，推进项目管理能力。

---

## 4. 社区热点

**Issue #967** 是今日唯一有讨论更新的 Issue（2 条评论），且最终被关闭。用户 `svier0` 详细复现了高频率的 NoResponseContent 错误，并指出相同模型和 API Key 在 picocla**（可能为其他工具）中运行正常，暗示问题集中在 NullClaw 自身。该问题关注度虽不高（0 👍），但反映了 Windows 环境下部分用户的稳定性诉求。

---

## 5. Bug 与稳定性

| 严重程度 | Issue | 描述 | 状态 | Fix PR |
|----------|-------|------|------|--------|
| 🔴 高 | #967 | Windows 11 下 `error: NoResponseContent` 出现概率 >50%，模型为 Agnes-2.0-Flash，导致 agent 对话失败 | 已关闭 | 未明确关联 PR，推测为直接修复或配置调整 |

短期内无新增 Bug 报告，整体稳定性风险可控。

---

## 6. 功能请求与路线图信号

- **PR #783** 包含多项新功能：基于数据库的 Cron 子代理调度器、运行历史记录、JSON 格式 CLI 输出（`cron list --json`）、安全加固等。虽然尚未合并，但持续更新表明作者正积极完成质量保障。若此 PR 被采纳，预计将在下一版本显著增强自动化运维和可扩展性。
- 当前无新增 Feature Request Issue。社区需求暂集中于稳定性修复，功能方向主要依赖维护者的计划。

---

## 7. 用户反馈摘要

来自 Issue #967 的用户反馈：

- **痛点**：在 Windows 11 上调用 `nullclaw agent -m "你好！"` 频繁报错“NoResponseContent”，导致基本对话功能不可用。用户强调“同样的模型同样的 apikey”在其他客户端正常，暗示 NullClaw 在响应处理上存在兼容性或超时问题。
- **使用场景**：即时对话 / 测试交互。
- **满意点**：未直接表达，但能提供详尽的复现步骤（包含版本、模型、频率统计）说明用户对项目有一定信任，且维护者在 3 天内关闭该 Issue，响应速度尚可。

---

## 8. 待处理积压

- **PR #783**（2026-04-07 创建，已持续 78 天）：功能重大但迟迟未合并，涉及定时任务引擎、JSON 输出等核心改动。建议维护者评估是否需拆分为更小粒度以加快合并，或由社区协助 review。  
  [https://github.com/nullclaw/nullclaw/pull/783](https://github.com/nullclaw/nullclaw/pull/783)

- 暂无长时间未响应的 Issue 积压。项目整体积压压力较小。

---

**报告总结**：NullClaw 当前处于平稳维护期，社区贡献集中在已知 Bug 修复与长期功能 PR 的打磨上。建议关注 PR #783 的合并进展，以及确认 #967 的修复是否对 Windows 用户生效。项目健康度良好，但贡献活跃度有待提升。

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是根据 IronClaw (nearai/ironclaw) GitHub 数据生成的 2026-06-24 项目动态日报。

---

## IronClaw 项目日报 | 2026年6月24日

### 1. 今日速览

今日项目活动量极高，社区与核心团队均保持活跃。过去24小时内，共有14条Issue更新和48条PR更新，显示出项目正处于密集的开发与测试周期。核心关注点集中在 **Reborn 回归问题**及**托管服务的可观测性缺失**上。尽管有新功能仍在推进，但一个导致基准测试21/147任务失败的严重回归问题已得到确认和关闭，凸显了持续集成管线的敏感性。总体来看，项目健康度呈现“高度活跃但伴有关键回归”的状态。

### 2. 版本发布

无。

### 3. 项目进展

今日合并/关闭了28条PR，推进了多项关键修复和功能。

- **Slack 集成优化**：合并了一系列针对 Slack 通道的PR，引入了 `IRONCLAW_REBORN_SLACK_ENABLED` 环境变量覆盖开关 ([#5181](https://github.com/nearai/ironclaw/pull/5181), [#5162](https://github.com/nearai/ironclaw/pull/5162))，并清理了遗留配置字段 ([#5161](https://github.com/nearai/ironclaw/pull/5161))，使 Slack 的部署和启用更加灵活。
- **Telegram 集成加固**：合并了 [#5143](https://github.com/nearai/ironclaw/pull/5143)，增加了 Telegram 主机入口验证器的回归测试，并修复了其失败关闭逻辑。
- **NEAR AI MCP 稳定性**：合并了 [#5178](https://github.com/nearai/ironclaw/pull/5178)，修复了在没有持久化认证存储的情况下自动引导 NEAR AI MCP 导致的悬挂状态问题，提升了本地开发体验的健壮性。
- **GitHub API 修复**：合并了 [#5171](https://github.com/nearai/ironclaw/pull/5171)，修正了 Reborn 内置 GitHub WASM 扩展的 API 请求形状，确保了与 GitHub 后端交互的正确性。
- **廉价模型配置固化**：合并了 [#2726](https://github.com/nearai/ironclaw/pull/2726)，完成了 `cheap_model` 和 `smart_routing_cascade` 配置的架构移植和硬化工作，确保其在当前架构下的正确性和持久化。

### 4. 社区热点

今日讨论高度集中于 **Reborn 回归问题** 和 **托管服务的可观测性**。

- **[#5139] reborn regression: web/research tasks hang at init (0 LLM calls) on main HEAD** ([链接](https://github.com/nearai/ironclaw/issues/5139)): 这是今日最受关注的Issue。它报告了主线10次提交后导致的严重回归，导致21/147的基准测试任务因超时而失败，且未发起任何LLM或工具调用。该问题已经关闭，表明团队已快速响应并找到根因，但短期内对社区信心造成了冲击。
- **[#5182] Reborn hosted observability: meaningful logs + failure diagnostics out of the binary** ([链接](https://github.com/nearai/ironclaw/issues/5182)): 社区成员 `serrrfirat` 提出了在托管部署环境下，从二进制文件中提取有意义的诊断信息的强烈需求。当前的操作困难阻碍了故障诊断，社区对增强项目可观测性有很高的呼声。

### 5. Bug 与稳定性

今日报告的Bug数量较多，其中包含一个已修复的关键回归问题。

- **严重 - 回归**:
    - **[#5139] [CLOSED]** `main` HEAD 的10次提交导致 web/research 任务初始化挂起，零 LLM 调用。此问题已被关闭（状态已关闭），推测已通过回滚或热修复解决。 ([链接](https://github.com/nearai/ironclaw/issues/5139))
- **中等 - 功能阻塞**:
    - **[#4986]** Reborn 中创建的自动化任务可能因等待工具批准而永久阻塞。 ([链接](https://github.com/nearai/ironclaw/issues/4986))
    - **[#5169]** 捆绑技能中的API词汇触发了提示安全拒绝列表，导致良性请求被误判为“临时系统问题”而死掉。 ([链接](https://github.com/nearai/ironclaw/issues/5169))
    - **[#5179]** 多租户用户在Web UI中无法访问日志，阻碍了调试。 ([链接](https://github.com/nearai/ironclaw/issues/5179))
- **中等 - 功能异常**:
    - **[#3733]** 无效的 Gmail Token 仍会显示“成功/已激活”的提示，造成用户困惑。 ([链接](https://github.com/nearai/ironclaw/issues/3733))
    - **[#3732]** Gmail 认证在不同对话中显示不一致的UI，有时是OAuth链接，有时是手动Token输入。 ([链接](https://github.com/nearai/ironclaw/issues/3732))

以上多个Bug均有活跃的评论。

### 6. 功能请求与路线图信号

- **可观测性提升**：[#5182](https://github.com/nearai/ironclaw/issues/5182) 明确提出了对托管部署进行更好的日志记录和诊断的需求。这是一个强烈的路线图信号，表明社区（和核心开发者）期望下一版本优先解决可观测性问题。目前尚无直接的PR关联，但优先级可能很高。
- **UI/UX改进**：[#5179](https://github.com/nearai/ironclaw/issues/5179) 要求多租户用户能访问日志，[#5167](https://github.com/nearai/ironclaw/issues/5167) 建议停止在git中跟踪`dist`目录以避免PR中的碎片变更。虽然`dist`的建议是开发流程优化，但用户反馈驱动的UI改进（如日志访问）应被纳入近期路线图。
- **自动化删除支持**：[#5122](https://github.com/nearai/ironclaw/issues/5122) 请求为Reborn自动化添加删除支持。该Issue已被关闭，很可能已通过相关PR实现或被判定为超出当前范围。

### 7. 用户反馈摘要

- **痛点：调试困难**。用户普遍反映在托管环境中调试失败任务非常痛苦 ([#5182](https://github.com/nearai/ironclaw/issues/5182))，尤其是在多租户场景下，无法从UI查看日志 ([#5179](https://github.com/nearai/ironclaw/issues/5179))。
- **误判与误导**：提示安全拒绝列表将正常API词汇误判为高风险内容，导致任务神秘失败 ([#5169](https://github.com/nearai/ironclaw/issues/5169))。同时，无效Token被显示为成功 ([#3733](https://github.com/nearai/ironclaw/issues/3733))，这些误导性信息损害了用户体验。
- **UI不一致**：同一功能的认证流程在不同上下文中显示不同界面，造成用户困惑 ([#3732](https://github.com/nearai/ironclaw/issues/3732))。

### 8. 待处理积压

- **[#4108] Nightly E2E failed:** 自5月底起持续失败，虽然自动化报告，但未看到活跃的讨论或分配。长期存在的CI失败是项目稳定性的一个危险信号，需尽快调查根因。 ([链接](https://github.com/nearai/ironclaw/issues/4108))
- **[#3733] Invalid Gmail token shows success/activated toast:** 存在一个多月，仍为开放状态，未分配处理者。该问题影响用户体验，建议安排在下个迭代中处理。 ([链接](https://github.com/nearai/ironclaw/issues/3733))
- **[#4032] build(deps): bump the wasm group:** 一个来自 `dependabot` 的依赖更新PR，已开放一个月。应检查其合并阻塞原因（如测试失败）并及时处理，以保持依赖项最新。 ([链接](https://github.com/nearai/ironclaw/pull/4032))

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI 项目动态日报 | 2026-06-24

## 1. 今日速览

过去24小时内，LobsterAI 项目保持高活跃度，共处理 **45 条 Pull Request**，其中 **42 条已合并/关闭**，仅剩 **3 条待合并**，表明团队正集中精力推进代码入库和缺陷修复。Issues 侧活跃度较低，仅有 **1 条旧 Issue (#1394)** 被标记为 stale 并更新，暂无新增缺陷报告。整体来看，项目处于 **密集迭代收尾阶段**，大量 PR 聚焦于稳定性修复、网关优化及模型扩展，开发健康度良好。

## 2. 版本发布

**无新版本发布**。上一版本信息请参考项目 Releases 页面。

## 3. 项目进展 — 重要已合并 PR 解读

今日合并的 42 条 PR 覆盖了 OpenClaw 网关、LLM 流输出、工具循环、会话冻结、UI 更新及模型支持等多个领域。以下是其中影响较大的关键 PR：

| PR | 标题 | 主要贡献 |
|-----|------|----------|
| [#2196](https://github.com/netease-youdao/LobsterAI/pull/2196) | fix(openclaw): avoid extra dock apps during shell snapshots | 修复 OpenClaw 网关在 macOS/Linux 下快照时产生多余 dock 应用的问题，统一启动路径 |
| [#2049](https://github.com/netease-youdao/LobsterAI/pull/2049) | fix(openclaw): prevent aborted tool loops from burning tokens | 解决空闲期持续消耗 token 的严重问题，增加中止循环检测机制 |
| [#2048](https://github.com/netease-youdao/LobsterAI/pull/2048) | fix: filter out empty data from LLM streaming output | 过滤 LLM 流式输出中的空数据，提升消息质量 |
| [#2047](https://github.com/netease-youdao/LobsterAI/pull/2047) | fix: solve the problem of session freezing | 解决会话冻结问题，改善用户体验 |
| [#2078](https://github.com/netease-youdao/LobsterAI/pull/2078) | feat(cowork): emit selected-skill routing metadata instead of inlining prompts | 重构 cowork 技能路由，将元数据与提示分离，提升可维护性 |
| [#2089](https://github.com/netease-youdao/LobsterAI/pull/2089) | chore: add minimax m3 and update byok models default context windows | 新增 Minimax M3 模型支持，更新 BYOK 模型的默认上下文窗口 |
| [#2086](https://github.com/netease-youdao/LobsterAI/pull/2086) | fix: fix a WeChat bug during updates/reinstalls | 修复微信集成在更新/重装时的兼容性 Bug |

这些 PR 的批量合并标志着项目在 **稳定性、资源消耗控制、模型生态扩展** 三方面取得了显著进展。特别是工具循环 token 燃烧和会话冻结这类影响核心体验的问题已得到修复。

## 4. 社区热点

### 最活跃 Issue
- **[#1394](https://github.com/netease-youdao/LobsterAI/issues/1394)**：定时任务选择“不重复执行”后，执行一次会被永久删除（用户预期应保留编辑）。该 Issue 创建于 2026-04-03，今日被标记为 stale 但仍有 1 条新评论，说明用户持续关注。此问题涉及任务管理的基础行为，社区诉求明确：**希望保留“仅执行一次”的任务供后续复用编辑**。

### 热点 PR（评论区活跃度较高，虽未显示具体评论数，但标题涉及的修复具有广泛讨论价值）
- [#2049](https://github.com/netease-youdao/LobsterAI/pull/2049) 关于 tool 循环 token 燃烧的修复，直接回应用户报告的“空闲期持续扣费”痛点，是今日社区最关心的 PR 之一。

## 5. Bug 与稳定性

按严重程度排列：

| 严重程度 | 问题描述 | Issue/PR | 状态 | 是否有 Fix PR |
|----------|----------|-----------|------|---------------|
| **高** | 空闲期间工具循环导致无意义 token 消耗 | [#2049](https://github.com/netease-youdao/LobsterAI/pull/2049) | 已合并 | ✅ 已修复 |
| **高** | 会话冻结，用户无法继续交互 | [#2047](https://github.com/netease-youdao/LobsterAI/pull/2047) | 已合并 | ✅ 已修复 |
| **中** | 定时任务执行后自动删除，不符合可编辑预期 | [#1394](https://github.com/netease-youdao/LobsterAI/issues/1394) | OPEN / Stale | ❌ 暂无对应 PR |
| **低** | OpenClaw 网关在 shell 快照时产生多余 dock app | [#2196](https://github.com/netease-youdao/LobsterAI/pull/2196) | 已合并 | ✅ 已修复 |

此外，还包括对 LLM 空数据过滤（#2048）、Gateway 重启问题（#2043）、会话 patch 超时（#2050）等若干稳定性修复，均已合并。

## 6. 功能请求与路线图信号

当前无新的功能请求 Issue 提出，但从已合并 PR 可以看出项目下一阶段的演进方向：

- **模型生态扩展**：新增 Minimax M3 模型、Mimo v2.5 系列、更新默认上下文窗口（#2089, #2102），表明团队正积极适配更多 LLM 提供商。
- **协作能力增强**：cowork 技能路由重构（#2078），为未来更灵活的插件化协同工作流打下基础。
- **平台兼容性**：Windows 下 PowerShell 替换 VBScript（#2057）、统一 OpenClaw 启动路径（#2195），持续打磨多平台体验。

综合来看，下一版本很可能包含 **更多模型支持、自动断点恢复机制、以及更稳定的协作者模块**。

## 7. 用户反馈摘要

从唯一活跃 Issue #1394 的评论内容提炼：

- **痛点场景**：用户创建了一个“不重复”的定时任务，期望在手动触发或自动执行完成后，任务仍保留在列表中供编辑、复制或后续再次使用。当前行为是直接删除，导致用户必须重新创建，操作成本高。
- **使用场景**：适用于一次性提醒、临时脚本执行等，用户希望保留任务配置作为模板。
- **用户预期**：任务状态应标记为“已完成”而非被删除，允许手动重置下次执行时间。

该反馈反映了对任务生命周期管理的精细化需求，属于合理的用户体验优化点。

## 8. 待处理积压

| 类型 | 编号 | 标题 | 创建时间 | 最后更新 | 备注 |
|------|------|------|----------|----------|------|
| Issue | [#1394](https://github.com/netease-youdao/LobsterAI/issues/1394) | 定时任务不重复执行后被自动删除 | 2026-04-03 | 2026-06-24 | 已标记 stale，但今日有评论，需维护者评估保留行为并决定是否修复 |
| PR | 待合并的 3 条 PR | 未在今日展示列表中，请查看项目 PR 页签获取详情 | — | — | 建议优先关注其中的高优先级修复（如涉及安全或核心功能） |

> 提示：项目当前没有长期未响应的重大 Issue，整体积压健康。建议维护者对 #1394 给出回应（确认行为设计或排期修复），避免因 stale 自动关闭后用户不满。

---

**总结**：LobsterAI 在 2026-06-24 展现了强劲的开发执行力，大量关键性修复和功能改进已合并入库。虽然社区反馈较少，但项目团队仍保持了高频迭代节奏。建议在下一次版本发布前，重点确认 #1394 的逻辑调整，并完成剩余 3 条待合并 PR 的审核。

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

# CoPaw 项目动态日报 — 2026-06-24

## 今日速览

过去 24 小时项目保持中度活跃：共产生 16 条 Issue（新开/活跃 11 条，关闭 5 条）和 50 条 PR（待合并 40 条，已合并/关闭 10 条）。未发布新版本。社区讨论集中在**自定义 OpenAI 兼容提供商缺失 function calling**（#5345，8 条评论）与**安装后启动报错**（#5379，5 条评论）两大问题上，前端性能相关 Bug（大会话崩溃、长消息排版错乱）也频繁被报告。合并的 PR 主要修复了 cron 作业兼容性、启用了 cron 编辑/删除能力，并增强了桌面 UI 测试覆盖与移动端响应式布局，整体稳定性和基础设施有所改善。

## 版本发布

- 无新版本发布（当前最新版本为 v1.1.12.post2）。

## 项目进展

今日共合并/关闭 10 条 PR，重要整合如下：

- **cron 功能修复**：  
  - `#5483` [CLOSED] [fix(cron)] 允许在 cron 作业启用时直接编辑和删除，无需先停用后操作。  
  - `#5475` [CLOSED] [fix(cron)] 修复因遗留序列化数据（纯文本 Msg.content、非 JSON 格式 tool_call.input）导致 cron 作业反复失败的问题，增强向下兼容性。  
  - 这两项 PR 显著提升了 cron 功能的健壮性和用户体验，解决了社区长期抱怨的“修改 cron 需删除重建”痛点。

- **桌面 UI 与测试基建**：  
  - `#5428` [CLOSED] [feat(ci)] 新增桌面版本端到端 UI 验证流水线，覆盖 Legacy Win/Mac 和 Tauri Win/Mac 四种平台的安装启动与前端渲染检查，减少发布故障。  
  - `#5368` [CLOSED] [feat(console)] 优化 SkillPool 页面在移动端的响应式布局，解决溢出/错位问题。

- **其他修复**：  
  - `#5487` [OPEN]？实际上是 CLOSED？数据标记为 OPEN，但摘要显示修复 content delta 绕过流路径的问题，暂未合并。  
  - 多个前端单元测试 PR（#5434、#5438）持续增加覆盖率，但尚在开放状态。

总体来看，项目在**核心稳定性和持续集成**上有所推进，但大量功能性 PR（如桌面自更新、Windows GUI 自动化、DataPaw 插件等）仍处于待合并状态，整体前进速度受限于 PR 积压。

## 社区热点

| 标题 | 类型 | 评论数 | 链接 |
|------|------|--------|------|
| [Bug]: Custom OpenAI-compatible providers (e.g. OMLX) don't support function calling | Issue | 8 | [#5345](https://github.com/agentscope-ai/QwenPaw/issues/5345) |
| [Bug]: 通过Python命令安装后启动，直接报错Internal Server Error | Issue | 5 | [#5379](https://github.com/agentscope-ai/QwenPaw/issues/5379) |
| [Question]: 1.1.11及以下版本，windows desktop版本，前端加载不流畅（已关闭） | Issue | 5 | [#5015](https://github.com/agentscope-ai/QwenPaw/issues/5015) |
| [Bug] 群聊消息回复被发送到私聊而非群聊（已关闭） | Issue | 5 | [#5264](https://github.com/agentscope-ai/QwenPaw/issues/5264) |

**热点分析**：  
- **#5345** 是当前社区最关注的功能 Bug——用户期望自定义 OpenAI 兼容提供商（如 OMLX）能正常触发 function calling，但实际只返回纯文本。该问题与模型生态扩展直接相关，尚无对应 PR，开发者已明确期望修复。  
- **#5379** 集中体现了安装体验痛点：最新版本通过 `pip install` 后直接启动即报 `Internal Server Error`，日志指向 `get_remote_addr(transport)`，严重阻碍新用户上手。  
- 已关闭的 #5264（群聊回复错发到私聊）和 #5015（Windows 桌面端卡顿）表明社区对通道一致性和前端性能的长期关注。

## Bug 与稳定性

按严重程度排列（附 fix PR 状态）：

| 严重性 | Bug 描述 | Issue | 是否有修复 PR | 备注 |
|--------|----------|-------|---------------|------|
| **严重** | 单个会话文件 >500KB 时 Web UI 直接白屏崩溃 | [#5479](https://github.com/agentscope-ai/QwenPaw/issues/5479) | ❌ 无 | 影响所有长对话用户，无法浏览历史 |
| **严重** | 前端渲染大工具调用历史时崩溃（`type:"data"` content blocks） | [#5401](https://github.com/agentscope-ai/QwenPaw/issues/5401) | ❌ 无 | Console 全白屏 |
| **严重** | 安装后启动即报 Internal Server Error | [#5379](https://github.com/agentscope-ai/QwenPaw/issues/5379) | ❌ 无 | 阻塞新用户入门 |
| **中等** | 长消息排版错乱（Markdown 格式混在一起），切选项卡后恢复 | [#5480](https://github.com/agentscope-ai/QwenPaw/issues/5480) | ❌ 无 | 视觉影响严重，但可临时规避 |
| **中等** | 通过 OpenCode Go 使用 GLM 5.x 模型时 `json_schema_converter.cc` 报错 | [#5472](https://github.com/agentscope-ai/QwenPaw/issues/5472) | ❌ 无 | 特定模型兼容性问题 |
| **低** | 自定义 OpenAI 提供商不触发 function calling | [#5345](https://github.com/agentscope-ai/QwenPaw/issues/5345) | ❌ 无 | 影响第三方模型接入 |
| 已关闭 | 群聊消息回复错发到私聊 | [#5264](https://github.com/agentscope-ai/QwenPaw/issues/5264) | ✅ 已关闭（推测已修复） | - |
| 已关闭 | 移动端无法切换智能体 | [#5476](https://github.com/agentscope-ai/QwenPaw/issues/5476) | ✅ 已关闭 | - |

**稳定性评估**：项目存在至少 3 个严重级别的前端/启动 Bug，且均无对应修复 PR，需维护者优先关注。#5479 和 #5401 直接导致应用不可用，影响核心用户体验。

## 功能请求与路线图信号

以下来自 Issue 的功能请求反映了用户对**扩展生态与易用性**的明确需求：

| 功能 | Issue | 关联 PR | 预计纳入版本 |
|------|-------|---------|-------------|
| 支持 OpenAI response format（结构化输出） | [#5489](https://github.com/agentscope-ai/QwenPaw/issues/5489) | 无 | 可能为 v2.x 特性 |
| MCP 工具名称在聊天界面显示原始名称（而非清洗后名称） | [#5231](https://github.com/agentscope-ai/QwenPaw/issues/5231) | 无 | 属于界面优化，低风险 |
| 支持通过 pip 从 PyPI 安装插件 | [#5484](https://github.com/agentscope-ai/QwenPaw/issues/5484) | [PR #5484](https://github.com/agentscope-ai/QwenPaw/issues/5484) ？（原始数据将该 PR/Issue 标为 Issue，但标题为 feat） | 社区呼声高，已有提案 |
| Kimi Coding Plan 模型配置支持（Anthropic 兼容端点） | [#5427](https://github.com/agentscope-ai/QwenPaw/issues/5427) | 无 | 特定 provider 适配 |

此外，PR 列表中已出现 **DataPaw 数据分析插件**（#4622）、**Windows 桌面 GUI 自动化**（#5187）、**系统托盘启动**（#4041）等大型功能，均处于待合并状态，说明项目路线图正在向“通用 Agent 平台”演进，但合并速度滞后于开发速度。

## 用户反馈摘要

从 Issue 评论和摘要中提取的真实用户声音：

- **“自定义提供商不能用 function calling，但 Ollama 可以。”**（#5345）—— 用户希望在 QwenPaw 中使用 OMLX 等第三方服务，却因缺乏兼容性而受阻。  
- **“刚 pip 安装完启动就直接报错，看日志也不知道怎么解决。”**（#5379）—— 安装体验成为新用户第一个障碍，缺乏清晰的错误指引。  
- **“大会话文件 500KB 就崩溃，我只能删除会话才能继续用。”**（#5479）—— 高频用户因会话积累无法查看历史，严重影响日常工作流。  
- **“长消息排版完全错乱，切换选项卡又好了，每次都要手动切。”**（#5480）—— 前端 CSS 重计算缺失，用户被迫采取临时措施。  
- **“刚启动内存就 1.4G，能不能优化一下？”**（#5441，已关闭）—— 虽然关闭，但用户不满情绪明显，期望资源占用优化。  
- **“移动端没有切换智能体的按钮，PC 端左上角就有。”**（#5476，已关闭）—— 用户对平台一致性的敏感。

## 待处理积压

以下 Issue/PR 长期未响应或处于阻塞状态，建议维护团队重点关注：

| 项目 | 类型 | 创建时间 | 天数 | 说明 |
|------|------|----------|------|------|
| [#4669](https://github.com/agentscope-ai/QwenPaw/pull/4669) | PR | 2026-05-25 | 30 天 | **Tauri 桌面自更新**，基础功能 PR，长期无 review 或合并，阻碍桌面版本自动更新发布 |
| [#4041](https://github.com/agentscope-ai/QwenPaw/pull/4041) | PR | 2026-05-05 | 50 天 | **系统托盘启动项**（仅 Win32），独立入口，长期未回应 |
| [#4622](https://github.com/agentscope-ai/QwenPaw/pull/4622) | PR | 2026-05-22 | 33 天 | **DataPaw 数据分析插件**，功能完整，但未获 maintainer 评估 |
| [#5210](https://github.com/agentscope-ai/QwenPaw/pull/5210) | PR | 2026-06-15 | 9 天 | **CLI cron update 命令**（实现 `PUT /cron/jobs/{job_id}`），配合已合并的 #5483 可形成完整 cron 管理体验，建议加速 |
| [#5345](https://github.com/agentscope-ai/QwenPaw/issues/5345) | Issue (Bug) | 2026-06-20 | 4 天 | 虽不算长期，但评论数最多且无 assignee，需尽快确认修复计划 |

**项目健康度简评**：Issue 响应及时（多数在 1-2 天内有关闭或回复），但 PR 积压严重，5 月提交的多个功能 PR 至今未合并，可能导致社区贡献者流失。当前 Bug 的修复速度低于用户期望，建议优先解决 #5479、#5401 等前端崩溃问题，并加速合并已有代码的 PR（如 #4669、#4041）以释放贡献者积极性。

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是根据 ZeroClaw 项目数据生成的 2026 年 6 月 24 日项目动态日报。

---

# ZeroClaw 项目日报 - 2026-06-24

## 今日速览

ZeroClaw 项目今日保持极高的开发活跃度。尽管没有新版本发布，但提交的 Issues（12 条）和待合并的 PR（41 条）数量庞大，表明项目正处于密集的功能开发和问题修复冲刺期。社区讨论焦点集中在安全策略、供应键安全（SLSA）和 WASM 插件生命周期三大领域。尤其值得关注的是，多个严重级别（S0/S1）的安全和运行时 Bug 被报告并迅速获得了修复 PR，体现出项目团队对稳定性和安全的快速响应能力。总体来看，项目 **非常活跃**，进展迅速，但版本稳定性可能需要一次集中发布来解决。

## 版本发布

今日无新版本发布。

## 项目进展

过去 24 小时内有 9 个 PR 被合并或关闭，以下为重要进展：

- **安全修复（高优）**：针对 `delegate` 工具绕过父代理权限控制的安全漏洞（#8279），社区提交了至少 **3 个修复 PR**（#8284, #8285, #8286），试图从配置校验和运行时工具筛选两个层面进行防护。这表明问题已得到社区高度重视并迅速产出解决方案。
    - PR [#8286](https://github.com/zeroclaw-labs/zeroclaw/pull/8286): 扩展配置校验，防止子策略通过 `allowed_tools` 和 `excluded_tools` 实现权限升级。
    - PR [#8285](https://github.com/zeroclaw-labs/zeroclaw/pull/8285): 在 `delegate` 边界处对调用者的工具白名单进行强制交集运算。
- **可观测性增强**：PR [#8065](https://github.com/zeroclaw-labs/zeroclaw/pull/8065) 已合并，此重大更新为日志引入了 `trace_id` 进行关联，并记录了每次调用的成本 (`cost_usd`)，显著提升了项目的可调试性和成本监控能力。
- **配置健壮性**：PR [#8098](https://github.com/zeroclaw-labs/zeroclaw/pull/8098) 已合并，禁止创建名为 `default` 的代理，解决了此前因保留名称导致的配置操作不对称（创建/删除/重命名不一致）的问题。

这些推进表明项目在 **安全性** 和 **可观测性** 这两条主线上取得了实质性进展。

## 社区热点

1.  **PR: [RFC]: Supply chain signing** (#8177)
    - **来源**：[Issue #8177](https://github.com/zeroclaw-labs/zeroclaw/issues/8177)
    - **分析**：此 RFC 提出为容器镜像和 Release 二进制文件增加硬件 PGP、SLSA 来源验证等供应键安全措施，是项目安全规范的深化。虽评论数不多，但涉及 `domain:ci` 和 `domain:security` 两个核心领域，并且包含对其的 PR [#8277](https://github.com/zeroclaw-labs/zeroclaw/pull/8277)（为发布流水线添加 SLSA 证明），说明该议题正从讨论走向实施。

2.  **Issue: [Feature]: Prompt-triggered install suggestions** (#6289)
    - **来源**：[Issue #6289](https://github.com/zeroclaw-labs/zeroclaw/issues/6289)
    - **分析**：这是一个长期开放的功能请求，建议在用户发出请求时，如果缺少相应技能或插件，ZeroClaw 应智能提示安装。该议题在 24 小时内有更新，评论 5 条，是社区呼声较高且持续时间较长的需求，涉及用户体验的核心提升。

3.  **Issue: [Bug]: MCP stdio child processes accumulate** (#5903)
    - **来源**：[Issue #5903](https://github.com/zeroclaw-labs/zeroclaw/issues/5903)
    - **分析**：这是一个已持续两个月的严重 Bug（P1），`zeroclaw daemon` 开启心跳后会导致 MCP 子进程泄漏，产生大量孤儿进程。今日的更新表明该问题仍未解决，持续引发社区对运行时稳定性的担忧。

## Bug 与稳定性

- **安全风险 (S0 - 数据丢失/安全风险)**
    - **#8279**: `delegate` 权限绕过。子代理可调用父策略禁止的工具。**已有修复 PR #8284, #8285, #8286**。
    - **#7623**: API Key 泄漏。`delegate` 到需要 OAuth 的子代理时，错误地转发了协调器的 API Key。**状态为 `status:in-progress`**。
- **严重功能异常 (S1 - 工作流阻塞)**
    - **#8151** (已关闭): Matrix 频道中，延迟的图片附件在缓存历史中丢失引用，导致机器人无法再看到该图片。**已关闭，表明已有修复方案或已作为设计问题被接受**。
- **退化/严重行为 (S2)**
    - **#7733**: `mcp_bundles` 配置被解析和显示，但运行时未被强制执行，导致安全隔离的配置成为静默无操作。**有活动评论**。
- **性能/次要问题 (S3)**
    - **#8275**: Scoop 安装包未注册 `zerocode.exe`，导致用户无法通过命令行使用 TUI 客户端。

## 功能请求与路线图信号

- **核心架构演进**：**WASM 插件系统**（#7314, #6140, #7822）仍是 v0.8.2 版本的核心路线。`PluginCapability::Hook` (#7822) 的提出预示着 WASM 插件将具备监听代理生命周期事件的能力，这是一个关键的生态扩展。
- **用户体验提升**：**智能安装建议**（#6289）和 **Web 仪表盘自动升级**（PR #8173）反映出项目在降低用户门槛、改善运维体验上的努力。
- **安全与合规**：**供应键安全**（#8177）的 RFC 及后续 PR（#8129, #8277）是重要的安全基础设施，预计将纳入 v0.8.2 或之后版本。**配置权限模型**的持续加固（#8286 等）也是当前重点。
- **可观测性增强**：PR #8065（已合并）和 #8066（进行中）等表明，项目正在建立从 LLM 调用、日志到成本核算的完整可观测性体系，这对生产环境的运维至关重要。

## 用户反馈摘要

- **正面反馈/期望**：从 #6289 的讨论可以看出，用户对 ZeroClaw 不断增长的技能和插件生态感到兴奋，但**强烈的需求**是平台能主动、智能地引导用户发现和安装这些能力，而不是依赖用户自行搜索。
- **痛点/负面反馈**：
    - **安全问题**：`delegate` 的策略绕过和 API Key 泄漏问题 (#8279, #7623) 暴露了当前复杂代理嵌套场景下的安全风险，是用户的痛点。
    - **配置与行为不一致**：#7733 中 `mcp_bundles` 配置“看起来生效实则无效”的问题，会严重削弱用户对配置系统的信任，并可能引发安全盲区。
    - **部署与使用问题**：Scoop 安装包缺少 `zerocode.exe` (#8275) 虽然是小问题（S3），但影响了用户体验，尤其是在刚接触项目时。
    - **稳定性顾虑**：MCP 子进程泄漏 Bug (#5903) 已存在超过两个月，对运行 `daemon` 模式的用户是持续的性能和稳定性威胁。

## 待处理积压

- **长期未修复的高影响 Bug：**
    - **#5903** - MCP stdio 子进程泄漏。这是 P1 严重性且标记为 `status:accepted` 的 Bug，但自 4 月 19 日报送以来已过去两月，**强烈建议维护者优先处理或给出时间表**。
- **长期开放的重大功能请求：**
    - **#6140** - 混合技能和 WASM 工具插件。作为 WASM 插件路线图的核心部分，虽然标记为 `status:accepted`，但从 4 月 26 日创建以来尚未有具体的工程实施 PR。
    - **#6943** - RFC: 解决插件系统目标冲突的提议。该 RFC 提出了插件架构路线的重大变更（替换 Extism），虽被接受但缺乏后续，需要维护者确认推进计划。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*