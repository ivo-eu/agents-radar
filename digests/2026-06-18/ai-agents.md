# OpenClaw 生态日报 2026-06-18

> Issues: 200 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-06-18 03:18 UTC

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

好的，作为 AI 智能体与个人 AI 助手开源项目分析师，根据您提供的 OpenClaw 项目 GitHub 数据，我为您生成了 2026-06-18 的项目动态日报。

---

## OpenClaw 项目动态日报 | 2026 年 6 月 18 日

### 1. 今日速览

今日 OpenClaw 项目保持 **极高活跃度**。过去 24 小时内，社区提交和讨论非常密集，共产生 200 条 Issue 更新和 500 条 PR 更新。虽然合并/关闭率（PR 17.2%）相对较低，但大量待合并 PR（414 条）表明项目正处于功能密集开发与集成的阶段。**安全、会话状态和消息丢失**是今日社区讨论的核心关键词，多个高优先级问题围绕这些主题展开。值得注意的是，今日无新版本发布，但大量“回归”问题的报告提示近期版本可能引入了不稳定的变更。

### 2. 版本发布

无，过去 24 小时内无新版本发布。

### 3. 项目进展

尽管无新版本发布，但仍有 86 个 PR 被合并或关闭，体现了持续的内部整合与修复工作。以下是一些已关闭的重要 PR：

- **修复与稳定性提升：**
    - **[PR #94366] fix(cron): allow empty assistant replies in cron lane**：修复了定时任务（cron）在产生空回复时触发失败的问题，提升了自动化任务的鲁棒性。
    - **[PR #94397] fix(agents): render identity name in runtime prompt**：将 Agent 的配置名称渲染到运行时提示词中，改善了模型对自身身份的认知。
    - **[PR #94398] fix(gateway): expose idempotencyKey in chat history metadata**：在聊天历史中暴露幂等键，增强了消息去重和事务可靠性。
    - **[PR #92682] fix(read): use system encoding fallback for non-UTF-8 text files on Windows**：修复了 Windows 系统上读取非 UTF-8 编码（如中文 GBK）文件时出现乱码的问题。
    - **[PR #94209] fix(model): cap contextWindow at native runtime catalog limit**：当用户配置的上下文窗口超过模型实际限制时，自动降低至实际大小，避免了配置错误导致的问题。
    - **[PR #94234] fix(anthropic): allow failover for thinking signature replay errors**：允许在重放思考签名失败时进行故障转移，提升了长对话的稳定性。
    - **[PR #94277] fix(exec): add advisory file lock to prevent concurrent approval writes**：修复了高并发情况下 exec 授权文件写入丢失的问题，提升了安全性。
- **基础设施与维护：**
    - **[PR #68936] Autofix: add PR review autofix pipeline + Windows daemon**：一个大型自动化修复流水线和 Windows 后台守护进程被合并（此前已打开一段时间），表明项目在自动化基础设施上的持续投入。

**小结**：项目在解决兼容性问题、提升系统鲁棒性和自动化运维能力方面取得了稳步进展。大量针对近期回归问题的修复被合并，显示出维护者正在积极应对稳定性的挑战。

### 4. 社区热点

今日讨论最热烈、关注度最高的议题集中在以下几个领域：

- **AI 行为与 UX 问题**：
    - **[Issue #25592] Text between tool calls leaks to messaging channels**：这是目前最热门的问题，拥有 **32 条评论**。用户报告 AI 在执行工具过程中的中间思考文本（如错误处理、处理确认）被错误地发送到聊天渠道，造成严重的 UX 问题。该问题标有 `impact: message-loss` 和 `impact: security` 标签，反映出社区对 AI 行为可控性的高度关注。[链接](openclaw/openclaw Issue #25592)
    - **[Issue #9443] Request: Prebuilt Android APK releases**：拥有 **25 条评论**，社区用户强烈要求提供预构建的 Android APK，以便更便捷地使用 OpenClaw 的安卓伴侣应用。这反映了用户对移动端体验的迫切需求。[链接](openclaw/openclaw Issue #9443)

- **配置灵活性与模型兼容性**：
    - **[Issue #68596] Feature Request: Configurable streaming watchdog timeout threshold**：拥有 **15 条评论**和 **8 个赞**。用户反映在使用长推理模型时，现有的流式看门狗经常超时，强烈要求支持可配置的超时阈值。[链接](openclaw/openclaw Issue #68596)
    - **[Issue #93794] Messages on v2026.6.8 no longer supported on telegram web**：拥有 **8 个赞**，一个 P1 级别的回归问题，严重影响了 Telegram Web 用户的使用，引发了社区的强烈反应。[链接](openclaw/openclaw Issue #93794)

**分析**：社区讨论的焦点正在从单纯的“功能有无”转向“功能的智能性和稳定性”。用户对 AI 行为的预期更高，要求其不要展示内部“思考”过程，同时也对配置灵活性提出了更精细的需求。移动化和跨平台兼容性问题仍然是用户的痛点。

### 5. Bug 与稳定性

今日报告的 Bug 和回归问题较多，按严重程度排列如下：

- **严重 (P1, Critical)**：
    - **[#93794] Messages on v2026.6.8 no longer supported on telegram web**：P1 回归，严重阻断 Telegram Web 用户使用。无关联修复 PR。[链接](openclaw/openclaw Issue #93794)
    - **[#86215] Codex OAuth refresh failures can wedge an agent for hours**：P1，OAuth 刷新失败可能导致 Agent 长时间无响应，影响核心功能。无关联修复 PR。[链接](openclaw/openclaw Issue #86215)
    - **[#45224] Unhandled Playwright assertion error in CRSession._onMessage crashes Gateway**：P1，未捕获的 Playwright 断言错误会导致整个网关进程崩溃，影响服务稳定性。无关联修复 PR。[链接](openclaw/openclaw Issue #45224)
    - **[#31583] `exec` tool does not inherit `skills.entries.*.env` environment variables**：P1 回归，导致技能无法正确获取环境变量（如密钥），影响安全性和功能。有关联 PR ([#?])。[链接](openclaw/openclaw Issue #31583)
    - **[#43374] All LLM API calls time out simultaneously**：P1，多 Agent 并发时导致所有 LLM API 调用同时超时，严重影响多任务效率。无关联修复 PR。[链接](openclaw/openclaw Issue #43374)
    - **[#40255] Regression: User-configured heartbeat prompt no longer respected**：P1 回归，用户自定义的心跳提示词被忽略，影响 Agent 行为可控性。有关联 PR。[链接](openclaw/openclaw Issue #40255)

- **中等 (P2, High/Medium)**：
    - **[#43747] Memory management is in chaos**：P2 回归，用户报告内存管理行为不一致，且出现预期外的行为。无关联修复 PR。[链接](openclaw/openclaw Issue #43747)
    - **[#92451] v2026.6.x system prompt bloat causes instruction following degradation**：P2 回归，新版本系统提示词膨胀导致小模型指令遵循能力下降。该问题反映了版本迭代中可能存在的设计权衡问题。无关联修复 PR。[链接](openclaw/openclaw Issue #92451)
    - **[#44502] Discord routing / mention-gating issue**：P1/P2 回归，Discord 消息路由和提及门控逻辑出错。无关联修复 PR。[链接](openclaw/openclaw Issue #44502)
    - **[#38091] OpenClaw UI WebSocket reconnect 导致会话 terminated**：P1，WebSocket 重连导致会话异常终止，影响 Web 端用户体验。无关联修复 PR。[链接](openclaw/openclaw Issue #38091)

**小结**：Bug 报告密集，且多为 “regression”（回归），提示近期版本变更可能引入了风险。尽管有部分修复 PR 已提交，但仍有多个高优先级问题等待解决方案，项目的稳定性面临挑战。

### 6. 功能请求与路线图信号

今日用户提交的功能请求反映了以下路线图信号：

- **长期积压功能**：
    - **[#9443] Prebuilt Android APK releases**：呼声很高，应是下一阶段移动端策略的重要选项。
    - **[#7707] Memory Trust Tagging by Source**：社区对安全性的关注度很高，按源标记内存信任级别是防范“记忆投毒”的关键功能。[链接](openclaw/openclaw Issue #7707)
    - **[#41366] Durable natural-language rule learning**：多 Agent 群聊中的自然语言规则学习能力，指向更复杂的团队协作场景。[链接](openclaw/openclaw Issue #41366)

- **即将可能纳入的功能**：
    - **[#68596] Configurable streaming watchdog timeout**：开发者社区对该 PR 的讨论和点赞数很高，且逻辑相对独立，很可能被快速实现。
    - **[#14785] Reduce tool schema token overhead**：减少 Token 开销是提升性能和降低成本的关键优化点，且已有 Token 开销分析，实现路径清晰。[链接](openclaw/openclaw Issue #14785)
    - **[#92451] System prompt bloat**：该问题的出现本身就是信号，项目可能需要引入动态或按需加载系统提示词的机制。

### 7. 用户反馈摘要

- **痛点与不满**：
    - **内存管理混乱**：用户在 `[#43747]` 中抱怨不同的机器上内存管理行为不一致，有的进行分块嵌入，有的则不然，感觉缺乏统一标准。
    - **系统提示词膨胀**：用户在 `[#92451]` 反馈新版本在初始上下文中加入了约 20 多个新指令，严重稀释了小模型的注意力，导致指令遵循能力下降。
    - **多 Agent 并发问题**：用户在 `[#43374]` 描述了多 Agent 运行时所有 API 调用同时超时的现象，这与底层并发控制逻辑有关，严重影响效率。
    - **回退问题**：多个“回归”报告的提交者都表达了不满，特别是 `[#93794]` Telegram Web 消息不支持问题和 `[#40255]` 心跳提示词被忽略的问题，都是“以前能用，现在坏了”的典型情况。

- **使用场景与诉求**：
    - **复杂开发场景**：用户 `[#13700]` 请求会话快照功能，希望通过 `/session save|load` 在长开发会话中进行 A/B 测试或回滚，这指向了高级用户和开发者的使用场景。[链接](openclaw/openclaw Issue #13700)
    - **企业级部署**：用户 `[#43673]` 提出了组织/团队部署的需求，包括工作区脚手架、RBAC 和部署清单，标志着 OpenClaw 正在从小型个人项目向企业级应用演进。[链接](openclaw/openclaw Issue #43673)
    - **成本追踪**：用户 `[#13219]` 和 `[#9016]` 都提到了成本追踪，表明用户不仅关注功能，也开始关注运行成本和模型使用的优化。[链接](openclaw/openclaw Issue #13219)

### 8. 待处理积压

以下是一些长期未响应或已标记为“stale”的重要 Issue，建议维护者关注：

- **[#87857] Agent skips AGENTS.md startup sequence**：创建于 5 月 29 日，已标记为“stale”。该问题指出 Agent 会跳过 `AGENTS.md` 定义的强制性启动序列，这是一个核心行为问题，长时间未解决可能导致用户对 Agent 行为可控性失去信心。[链接](openclaw/openclaw Issue #87857)
- **[#7707] Feature Request: Memory Trust Tagging by Source**：创建于 2 月 3 日，是一个重要的安全特性请求。随着 LLM agents 和外部工具的深入集成，此类安全机制的重要性日益凸显，但其优先级似乎不高。
- **[#9443] Request: Prebuilt Android APK releases**：创建于 2 月 5 日，评论数高达 25 条，但至今仍未明确回应或排期，这可能会影响用户对项目交付能力的看法。
- **大量“待合并”PR**：目前有 **414 条 PR 处于待合并状态**，这个数字非常庞大。虽然部分可能因为需要更多审查或测试而延迟，但这通常也暗示了维护者审阅能力的瓶颈。建议对 PR 进行优先级排序，并建立僵尸 PR 清理机制。

---

## 横向生态对比

好的，作为一名资深技术分析师，我仔细审阅了过去24小时内上述13个开源项目的动态日报。尽管其中两个项目（TinyClaw, ZeptoClaw）在过去24小时无活动，其余项目均呈现出高度的活跃性与多样性。以下是根据您的要求生成的横向对比分析报告。

---

## AI 智能体与个人 AI 助手开源生态横向对比报告 (2026-06-18)

### 1. 生态全景

2026年6月18日，个人 AI 智能体开源生态呈现 **“功能密集迭代与稳定性阵痛并存”** 的态势。一方面，项目普遍向 **多智能体协作、安全加固、企业级部署、跨平台兼容** 等方向快速演进；另一方面，大量 **回归（regression）Bug**、**安全漏洞** 与 **资源管理问题** 频繁出现，暴露出在快速引入新功能时，对核心稳定性的测试覆盖不足。社区需求正从“能否工作”转向“是否可靠可控”，对 **配置灵活性、性能优化 (Token/内存)、以及可观测性** 的呼声显著增强。生态整体处于由社区驱动向工程化成熟度迈进的关键转型期。

### 2. 各项目活跃度对比 (2026-06-18)

| 项目 | Issues 更新 | PR 更新 (总/合并) | 版本发布 | 核心活跃度评估 | 健康度评估 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 200 | 500 / 86 | 无 | 🔥极高 | ⚠️中等（大量回归Bug报告） |
| **NanoBot** | 7 | 32 / 18 | 无 | 🔥极高 | ✅良好（合并效率高） |
| **Hermes Agent** | 10 | 50 / 4 | 无 | 🔥极高 | ⚠️中等（PR积压严重，高优先级Bug多） |
| **PicoClaw** | 4 | 10 / 6 | 无 | 🔥高 | ✅良好（安全与兼容修复快） |
| **NanoClaw** | 5 | 20 / 4 (2版本含破坏性变更) | **v2.1.17, v2.1.0** | 🔥极高 | ⚠️中等（文档与容器稳定性问题） |
| **NullClaw** | 3 (更新) | 2 / 0 | 无 | 🔸温和 | ✅良好（无严重Bug，社区响应及时） |
| **IronClaw** | 11 | 50 / 17 | 无 | 🔥极高 | ✅良好（Projects特性栈推动路线图） |
| **LobsterAI** | 1 | 11 / 11 | **v2026.6.15** | 🔥极高 | ⚠️中等（安全漏洞待处理） |
| **TinyClaw** | 0 | 0 / 0 | 无 | 💤无活动 | - |
| **Moltis** | 2 | 1 / 0 | 无 | 🔸温和 | ✅良好（功能积累期，无严重Bug） |
| **CoPaw** | 5+ | 8+ / 6+ | **v1.1.12, v1.1.12-beta.2** | 🔥高 | ⚠️中等（进程冻结等重度Bug） |
| **ZeptoClaw** | 0 | 0 / 0 | 无 | 💤无活动 | - |
| **ZeroClaw** | 17 | 50 / 11 | 无 | 🔥极高 | ⚠️中等（Windows/Termux兼容问题，S1阻塞） |

**注**：活跃度依据更新总量与核心变更密度；健康度基于Bug严重度、修复响应速度及社区反馈质量。

### 3. OpenClaw 在生态中的定位

- **优势**：作为生态核心参照项目，OpenClaw拥有 **最大的社区规模（200 Issues / 500 PR 日更新）**、最庞大的待处理PR池（414条），表明其功能开发投入最大。其在 **修复非UTF-8编码、幂等键、负载均衡** 等工程细节上的投入，体现了其作为“工业级” AI Agent 的标准地位。
- **技术路线差异**：与 **NanoBot** 的“轻量快速合并”和 **Hermes Agent** 的“深度安全与可观测性”不同，OpenClaw更倾向于通过 **广泛的社区贡献（大量PR涌入）** 来构建一个 **庞大而复杂的超级Agent**。这导致其稳定性控制面临比其他项目更严峻的挑战。
- **社区规模对比**：OpenClaw 的日更新量是 **NanoBot（7/32）**、**Hermes Agent（10/50）** 或 **ZeroClaw（17/50）** 的数倍至十倍，社区体量无疑最大。但 **NanoBot 与 IronClaw** 在合并效率上更优，生态密度更高。

### 4. 共同关注的技术方向

多个项目不约而同地涌现出以下需求，反映行业共性痛点：

| 技术方向 | 涉及项目 | 具体诉求 |
| :--- | :--- | :--- |
| **行为安全与可控性** | OpenClaw, Hermes Agent, PicoClaw, LobsterAI, ZeroClaw | 1) 工具执行中间思考不泄露至通信渠道；2) 内存/隐私工具可被禁用且不可绕过；3) 文件读取/路径遍历漏洞修复；4) 配置回退（fallback）与审批日志。 |
| **配置灵活性与可定制化** | OpenClaw, NanoBot, NullClaw, Moltis | 1) 按模型配置上下文窗口限制；2) 流式看门狗超时可配置；3) TTS输出格式可配置；4) 全局与局部环境变量继承。 |
| **多智能体/多租户管理** | NanoBot, PicoClaw, ZeroClaw, CoPaw | 1) 多Agent统一网关；2) 跨Agent消息循环防范；3) 子进程隔离与并发控制；4) 群聊/团队协作的规则学习。 |
| **跨平台兼容性** | OpenClaw, NullClaw, ZeroClaw | 1) Windows 非UTF-8编码；2) iOS Safari WebUI适配；3) Android Termux部署；4) 桌面版Profile切换问题。 |
| **性能与Token优化** | OpenClaw, ZeroClaw | 1) 系统提示词膨胀导致注意力下降；2) 原生上下文压缩（Provider管道装饰器）；3) 工具schema Token开销减少。 |

### 5. 差异化定位分析

- **功能侧重**：
  - **OpenClaw**：**全能型/工业级**。覆盖聊天、任务、自动化、网关、多媒体，最大社区，但稳定性挑战大。
  - **NanoBot**：**轻量型/快速迭代**。重点在飞书、文件系统安全、模型适配，PR合并高效，适合快速部署。
  - **Hermes Agent**：**服务型/可观测性**。引入OTLP、MCP二进制协议、安全审计，强调运维与安全性。
  - **PicoClaw**：**安全型/协议多样**。加密库替换、SSRF修复、Delta Chat/SimpleX网关，专注底层安全与去中心化。
  - **NanoClaw**：**实用型/敏捷发布**。专注会话隔离、容器稳定性、路径遍历修复，发布频繁但含破坏性变更。
  - **NullClaw**：**社区导向型/轻量功能**。CLI体验、内存可配置，更新温和但社区参与积极。
  - **IronClaw**：**架构与协作型**。Projects特性、自我进化、输出感知，侧重项目管理和Agent成长。
  - **LobsterAI**：**交互型/多模态**。Computer Use、实时ASR语音输入，强调人机交互体验。
  - **CoPaw**：**中文生态/重用户体验**。Console重构、简单模式、AgentScope 2.0架构升级，面向中文开发者。
  - **ZeroClaw**：**开源标准型/架构探索**。别名级联、ACP协议、WASM插件、零宕机热重载，技术探索激进。
- **目标用户**：
  - **OpenClaw, ZeroClaw**：高级开发者、企业用户。
  - **NanoBot, IronClaw**：开发者、DevOps、团队协作。
  - **Hermes Agent**：运维、安全专家、SRE。
  - **PicoClaw**：隐私敏感用户、去中心化爱好者。
  - **LobsterAI, CoPaw**：普通用户、中文社区。
- **技术架构差异**：OpenClaw（多语言/插件/网关生态）、NanoBot（Python/MCP/轻量）、Hermes Agent（Rust/安全/可观测）、ZeroClaw（Rust/WASM/声明式）。

### 6. 社区热度与成熟度分层

- **🔥快速迭代阶段（高活跃，功能优先）**：
  - **OpenClaw, NanoBot, Hermes Agent, IronClaw, ZeroClaw**：日更新量大，PR/Issue密集，正在积极拥抱新功能，但稳定性控制面临考验。
- **🔸质量巩固阶段（中活跃，修复与增强并行）**：
  - **PicoClaw, NullClaw, Moltis, CoPaw**：更新量适中，合并效率高，Bug报告较少或快速修复，项目相对健康稳定。
- **💤低活跃/无活动**：
  - **TinyClaw, ZeptoClaw**：过去24小无任何动态，可能处于开发间歇或维护停滞状态，需关注是否“弃坑”风险。

### 7. 值得关注的趋势信号

1.  **“内省”与“自我进化”成为新能力模型**：
    - **IronClaw PR #5061 (技能提取)** 和 **ZeroClaw RFC #7673 (上下文压缩)** 表明，Agent 正从“工具使用者”向“自动化自我优化者”演进。开发者应关注如何让 Agent 从对话历史或运行时数据中学习，形成可复用的“技能”或“策略”。

2.  **可观测性与审计成为企业级部署刚需**：
    - **Hermes Agent 的 OTLP 集成**、**ZeroClaw 的 RISC-V/SBOM 安全 CI**、**OpenClaw 的幂等键与日志暴露**，共同指向 AI Agent 在生产环境中需要与现有运维体系无缝对接。开发者应尽早集成 OpenTelemetry、安全扫描与审计日志。

3.  **多智能体交互的稳定性是下一个关键挑战**：
    - **CoPaw Issue #5204 (无限循环)**、**ZeroClaw 的多 Agent 并发问题** 以及 **NanoBot 的多 Agent 网关需求** 表明，当 Agent 之间能够相互通信时，如何防止“死循环”或“资源耗尽”成为了核心架构问题。开发者需在交互协议设计（如有限次转发、环检测）上提前布局。

4.  **固件级安全漏洞 (SSRF/路径遍历) 进入高发期**：
    - **PicoClaw 的 OneBot SSRF**、**NanoClaw 的 CWE-22 路径遍历**、**LobsterAI 的任意文件读取** 揭示了一个严峻趋势：随着 Agent 获取调用外部工具（文件、网络、POSIX命令）的能力，其本身成为了新的攻击面。**运行时沙箱与最小权限原则** 已从“功能增强”变为“生存必备”。

5.  **移动与边缘部署场景开始萌芽**：
    - **NullClaw 的手机端友好向导**、**ZeroClaw 的 Android Termux 请求** 以及 **CoPaw 的简单模式**，表明用户不仅希望 Agent 运行在云端或桌面，还期待在手机、树莓派等边缘设备上直接运行，并要求开箱即用的体验。低功耗、小内存、跨编译支持将成为未来项目竞争的新维度。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，我根据您提供的 NanoBot 项目 GitHub 数据，为您整理出 2026-06-18 的项目动态日报。

---

### NanoBot 开源项目日报 (2026-06-18)

**项目名称：** NanoBot
**数据统计周期：** 2026-06-17 至 2026-06-18
**分析师：** AI 智能体分析师

---

#### 1. 今日速览

今日 NanoBot 项目社区活跃度极高。**Pull Request (PR) 处理量达到了 32 条，其中 18 条已被合并或关闭，体现了维护团队高效的代码集成与问题解决能力。** 同时，有 7 条 Issues 更新，其中包含 6 条新开或活跃的讨论，显示社区反馈和需求持续涌入。虽然未有新版本发布，但大量已合并的 PR 表明项目处于快速迭代和功能完善阶段，项目整体健康度良好，向前迈进了坚实的一步。核心开发工作集中在提升飞书（Feishu）通道的稳定性与功能、加固文件系统安全策略，以及优化会话管理和模型适配方面。

#### 2. 版本发布

本次日报统计周期内，**无新版本发布**。

#### 3. 项目进展

今日项目取得了显著进展，共有 **18 个 PR 被合并或关闭**，主要集中在以下几个方面：

- **飞书（Feishu）通道增强与修复：**
    - **流更新恢复**：合并了 [#4381](https://github.com/HKUDS/nanobot/pull/4381) 以修复飞书流式更新失败后无法恢复的问题，通过重试前重新打开流模式来提升推送可靠性。
    - **WebSocket 卡片渲染**：[#4342](https://github.com/HKUDS/nanobot/pull/4342) 解决了飞书通过 WebSocket 接收到的卡片因结构不匹配导致无法正确渲染的问题。

- **安全性与配置修复：**
    - **工作区 Git 命令**：通过合并 [#4380](https://github.com/HKUDS/nanobot/pull/4380) 和 [#4393](https://github.com/HKUDS/nanobot/pull/4393)，修复了在 `restrictToWorkspace` 模式下，无法在工作区子目录内执行 git 命令的回归问题，并新增了测试覆盖。
    - **文件系统写入策略**：[#4053](https://github.com/HKUDS/nanobot/pull/4053) 和 [#4202](https://github.com/HKUDS/nanobot/pull/4202) 的合并，显著加固了文件系统的安全模型。明确将 `extra_allowed_dirs` 设为只读，并引入了读写分离的目录配置，防止工具意外修改重要文件。
    - **本地代理问题**：[#4367](https://github.com/HKUDS/nanobot/pull/4367) 修复了因环境变量设置了全局代理，导致无法连接本地模型服务（如 Ollama）的问题。

- **模型与工具改进：**
    - **Mistral 模型适配**：通过合并 [#4351](https://github.com/HKUDS/nanobot/pull/4351)，更好地支持了 Mistral API，解决了 `reasoning_effort` 参数、`tool_use` 内容类型等特定限制问题。
    - **智能体模型预设**：[#4347](https://github.com/HKUDS/nanobot/pull/4347) 修复了 “My Tool” 在切换模型预设时的行为，提供了更清晰的输出和错误提示。
    - **新搜索提供商**：[#4350](https://github.com/HKUDS/nanobot/pull/4350) 集成了 **Keenable** 作为新的内置网络搜索提供商。

- **其他改进：**
    - **WhatsApp 已读回执**：[#4354](https://github.com/HKUDS/nanobot/pull/4354) 为 WhatsApp 通道添加了消息已读状态标记功能。
    - **回退日志**：[#4385](https://github.com/HKUDS/nanobot/pull/4385) 优化了回退机制日志，在尝试回退模型前，会先打印主模型的错误信息，便于排查问题。
    - **WebUI 活动时长**：[#4283](https://github.com/HKUDS/nanobot/pull/4283) 修复了 WebUI 中活动持续时间的显示问题，使其更准确地反映任务的完成时长。

#### 4. 社区热点

今日讨论热度最高的议题反映了用户对**多智能体管理**、**新手易用性**和**高级功能**的强烈需求。

- **需求一：简化多智能体部署（多租户网关）**
    - **Issue**: [#936 - Feature Request: Add Multi-Tenant Gateway for Multiple Agents](https://github.com/HKUDS/nanobot/issues/936)
    - **热度**: 创建于 2 月，昨日再次被更新，引发讨论。这显示了用户对运行多个智能体时资源消耗和管理复杂度的痛点，期望通过一个统一的网关实例来管理所有智能体，是社区对**企业级、轻量化部署**的重要呼声。

- **需求二：提升新手用户体验（用户友好向导）**
    - **Issue**: [#4376 - enhancement: user friendly wizard](https://github.com/HKUDS/nanobot/issues/4376)
    - **热度**: 有 1 个 👍 支持，且评论活跃。用户 `chengyongru` 直接指出了当前 `nanobot onboard --wizard` 对非技术用户不友好，要求提供更便捷的无痛配置体验。这表明项目在向**更广泛的用户群体渗透**，易用性成为当前的关键改进点。

- **需求三：精细化的模型回退配置**
    - **Issue**: [#4389 - Feature Request: Per-model contextWindowTokens for fallback models](https://github.com/HKUDS/nanobot/issues/4389)
    - **热度**: 昨日新开即获得评论。该需求讨论了当回退模型上下文窗口小于主模型时，因缺乏按模型配置 `contextWindowTokens` 而导致的提示截断问题。这反映了用户在使用模型回退机制时，对**精确控制**和**模型特性适配**的深层次需求。

#### 5. Bug 与稳定性

今日报告的 Bug 主要集中在 WebUI 跨端兼容性和代码合并引入的问题，其中大部分已有修复。

| 严重程度 | Issue/PR | 描述 | 状态/进展 |
| :--- | :--- | :--- | :--- |
| 中 | [#4388 - [WebUI] iOS Safari 点击输入框触发页面放大](https://github.com/HKUDS/nanobot/issues/4388) | 在 iOS 26.5 的 Safari 上，点击 WebUI 输入框会导致页面自动放大和 UI 错位。即使在已包含移动端 UI 修复的版本中仍存在。 | **待修复**，评论区已提出该问题，需要前端开发者介入解决移动端适配。 |
| 高 | [#4322 - NameError: 'session_key' is not defined in context.py after merge](https://github.com/HKUDS/nanobot/issues/4322) | 合并代码后，智能体启动时崩溃。根因是提取 `_build_memory_context` 方法时引入的问题。 | **已修复**，该 Issue 已在昨日的更新中被关闭（Closed），问题已解决。 |
| 低 | [#4392 - fix(agent): make tool microcompaction configurable](https://github.com/HKUDS/nanobot/pull/4392) | 此 PR 反映工具结果的微压缩（microcompaction）行为不可配置，对缓存敏感的部署不利。 | **有 fix PR**，PR [#4392](https://github.com/HKUDS/nanobot/pull/4392) 提议添加 `microcompactToolResults` 配置项，使之可动态启用或关闭。 |

#### 6. 功能请求与路线图信号

除了上述社区热点中的功能请求，还有几个值得我们关注的信号，它们可能预示着项目下一版本的演进方向。

- **多实例 (Multi-instances)**：Issue [#4390](https://github.com/HKUDS/nanobot/issues/4390) 请求为普通用户简化多实例管理，希望免去手动配置 UI 和选项的麻烦，直接通过文件夹和 `config.json` 来管理。这与 [#936](https://github.com/HKUDS/nanobot/issues/936) 的多租户诉求相辅相成，都指向了**更高级的部署和管理能力**。

- **按需心跳调试**：早在4月提出的 Issue [#3437](https://github.com/HKUDS/nanobot/issues/3437) 昨日再次被更新。用户请求实现一个按需触发心跳的机制，用于调试 `HEARTBEAT.md`，并希望在跳过执行阶段（Phase 1）时无需支付执行成本（Phase 2）。这体现了高级用户对**内部机制的可观测性和可控性**的追求。

- **飞书 CLI 登录**：PR [#4391](https://github.com/HKUDS/nanobot/pull/4391) 提出了通过命令行扫码创建飞书机器人应用，简化了飞书通道的集成流程，具有很高的实用价值，很可能被合并。

#### 7. 用户反馈摘要

从近期 Issues 和 PR 的评论中，可以提炼出几类真实用户的反馈：

- **痛点反馈**：iOS Safari 用户在 Issue [#4388](https://github.com/HKUDS/nanobot/issues/4388) 中表达了对于页面放大和 UI 错位的强烈不满，并特别注明“即使在声称已修复移动端 UI 的版本中仍然存在”，这表明用户对项目解决该问题的预期很高，但实际情况与之存在差距。

- **使用场景反馈**：Issue [#4390](https://github.com/HKUDS/nanobot/issues/4390) 的作者描述了其 “一台机器，多实例管理” 的个人用例，并结合官方文档明确指出了当前配置方式的繁琐。这体现了**文档与实际用户体验之间的 Gap** 是用户痛点的重要来源。

- **用户体验反馈**：Issue [#4376](https://github.com/HKUDS/nanobot/issues/4376) 的作者 `chengyongru` 直言不讳地批评了当前向导“对新手或非技术人员不友好”，这是一个典型的对**入门体验不佳**的反馈，暗示着项目需要降低使用门槛。

#### 8. 待处理积压

以下是一些长期未得到响应或解决的重要 Issue，提请维护者关注：

- **功能请求：多租户网关** ([#936](https://github.com/HKUDS/nanobot/issues/936))
    - **提出时间**：2026-02-21 (距今约 4 个月)
    - **现状**：这是一个与项目规模化部署息息相关的核心功能请求，虽然昨日有更新，但长期处于开放状态，缺乏官方的路线图回应。社区对此功能的需求声音日益增加，建议项目核心团队评估并给出回应。

- **增强请求：按需心跳触发器** ([#3437](https://github.com/HKUDS/nanobot/issues/3437))
    - **提出时间**：2026-04-25 (距今近 2 个月)
    - **现状**：此 Issue 是对项目内部心跳机制的深度优化请求，虽昨日有更新，但尚未有明确进展。对于使用 `HEARTBEAT.md` 进行调试的开发者来说，这是一个有价值的改进。

- **待合并 PR：飞书扫码登录** ([#4391](https://github.com/HKUDS/nanobot/pull/4391))
    - **提出时间**：今日 (刚刚提出)
    - **现状**：该 PR 功能完整且有价值（简化飞书集成），是昨日刚提出的新 PR，应被优先评估和合并，以避免 PR 积压。

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是根据您提供的 Hermes Agent 项目 GitHub 数据生成的 2026 年 6 月 18 日项目动态日报。

---

# Hermes Agent 项目日报 - 2026-06-18

## 1. 今日速览

今日 Hermes Agent 项目整体活跃度**极高**，社区贡献和问题反馈均十分踊跃。过去24小时内共产生50条 PR 和10条 Issue，显示出项目正处于快速迭代期。然而，PR 积压情况严重（待合并46条），而同时有2个中等优先级的 Bug 已存在相应修复 PR，且有多个高优先级（P1）的 Bug 被报告，主要集中在 **OAuth 认证**和**安全绕过**等关键领域，需维护团队重点关注合并效率和紧急安全修复。项目版本层面暂无新发布，但功能与修复推进势头强劲。

## 2. 版本发布

- **无新版本发布**。

## 3. 项目进展

今日合并/关闭了4个 PR，主要成果如下：

- **安全计算改进**： #43051 已关闭，该 PR 修复了全局命令白名单（`command_allowlist`）中对 shell 样式 glob 模式（如 `podman *`）匹配失效的问题。这使得安全策略更为精确和灵活，是一项重要的代码执行安全加固。
- **系统服务修复**： #48178 解决了在 macOS 上通过 `launchd` 以 plist 格式配置网关时，XML 特殊字符（如 `&`, `<`, `>`）未转义导致服务启动失败的问题。此修复使用标准序列化工具替换了拼接模板，提高了系统集成的健壮性。
- **技能系统修复**： 修复了 ComfyUI 工作流 JSON 文件中包含 `_comment` 元数据键，导致向 ComfyUI API 提交请求时失败的问题。虽然出现了重复 PR（#48143， #48144， #48145），但问题本身已得到解决，非功能性障碍已清除。

**项目向前迈进的标志**：修复了安全策略执行和系统集成启动的核心阻塞问题，同时解决了 ComfyUI 插件的开箱即用体验，项目在稳定性和可用性上迈出了坚实一步。

## 4. 社区热点

今日最受关注的 Issue 是 **#38602**，该议题是关于“桌面客户端纯净安装”的功能需求，共获得 **18 个 👍** 和 **6 条评论**，充分反映了社区对于更轻量级、可远程连接桌面客户端的强烈诉求。用户描述的场景是希望安装 Hermes 桌面端作为“瘦客户端”，连接到远端已运行的 Hermes Agent，而非每次启动都要本地自举新实例。这表明社区用户对灵活部署架构（如客户端-服务器分离）有明确需求。

此外，**#48175** 关于“为多个 Discord 频道配置不同系统提示词”的提问也形成了讨论，用户希望在实践中能更优雅地管理 AI 人格复杂配置，体现了社区对精细化配置管理的好奇与探索。

## 5. Bug 与稳定性

今日报告的 Bug 分布在认证、安全、桌面端、工具兼容性等多个方面，其中部分风险较高。按严重程度排列如下：

- **P1 (Highest)**
  - **#48176** `[Bug]: OAuth Pro/Max/Team requests rejected...` (**已有 Fix PR #48177**)：报告了使用 Claude Pro/Max/Team 的 OAuth 凭证时，请求被返回 HTTP 400 错误，提示缺少第三方计费归属头部。这是一个严重的认证阻塞问题。幸运的是，PR #48177 已提交，向 OAuth 请求中添加了计费归属系统块。
- **P2 (High)**
  - **#48181** `[Security] Disabled memory toolsets can be bypassed...`：报告了即使配置“禁用” `memory` 工具集，晚期的内存提供者工具注入仍可绕过该限制。这是一个安全漏洞，可能导致用户隐私信息泄露。目前**暂无**对应的修复 PR 被标记。
  - **#48183** `[Bug]: 桌面版 Profile 切换后会话列表不显示...`：报告了桌面版在多 Profile 切换时出现重大问题，会话数据无法正确显示且可能分散。这直接影响到桌面产品的核心用户数据管理体验。目前**暂无**修复 PR。
  - **#48173** `Mid-session model switch leaves stale Model/Provider...`：报告了在会话中切换模型后，UI 及元数据中显示的模型和提供商信息仍然过时。这是一个典型的状态一致性问题，影响用户体验与调试。**暂无**修复 PR。

- **P3 (Medium)**
  - **#48172** `macOS app-bundle Chrome is ignored...` (**已有 Fix PR #48185**)：macOS 上无法自动检测到通过 App Bundle 安装的 Chrome 浏览器。PR #48185 已提交，将修复这一浏览器工具可用性检查的疏漏。
  - **#41808** `Dashboard Chat tab: React error #301`：报告了仪表盘 Web 聊天标签页在使用外部网络（如 Tailscale IP）连接时，会出现 React “最大更新深度” 错误。这是前端的一个非功能性 Bug，影响远程用户的使用。**暂无**修复 PR。

## 6. 功能请求与路线图信号

今日用户提交的功能请求不仅限于“锦上添花”，部分具有重要的架构和生态意义：

- **有望进入下一版本**：
  - **#48184** `feat(agent): add OTLP observability plugin` (**PR 已提交**)：这是一个非常有价值的信号，意味着社区正在积极为 Hermes Agent 引入标准的 **OpenTelemetry (OTLP)** 可观测性能力。这将对运维和生产化部署产生积极影响，很可能被核心团队采纳。
  - **#47740** `feat(mcp): add LUMEN binary protocol transport support` (**PR 已提交**)：提供了一个可选的二进制协议传输方案，号称有32-80%的带宽压缩和多代理共享能力。这符合 Agent 间高效通信的趋势，表明社区正探索更先进的 MCP 扩展。
  - **#38602** `[Feature]: Desktop Client-Only Installation`：此需求呼声极高，对应社区对更灵活部署模式的渴望。PR #48163 (系统托盘支持) 可以看作是向此方向迈进的一小步，但核心的“纯净客户端”功能很可能成为后续版本的重点考量。

- **信号较弱但仍值得关注**：
  - **#48179** `Extend managed system concept to be general for systems besides NixOS and Homebrew`：用户希望扩展“托管系统”概念以支持 Fedora (dnf)。这表明社区用户尝试将 Hermes 打包到更多 Linux 发行版中，对扩大生态覆盖面有参考价值。
  - **#48182** `Better interactive experience`：用户提出希望进行技能高亮显示和 `/` 斜杠命令补全等类似 Codex/Claude 的风格改进。这体现了社区对交互体验的精细化要求。

## 7. 用户反馈摘要

从今日的 Issues 和讨论中可以提炼出以下用户痛点和诉求：

- **核心痛点：安全与认证**
  - **OAuth 计费问题 (P1)**: 使用付费 Claude 账户的用户无法通过 OAuth 正常使用，被三方应用计费策略拦截，这可能是当前最影响付费用户使用的严重问题。
  - **安全配置可绕过 (P2)**: 用户对“禁用”内存工具的安全性表示担忧，担心隐私数据因框架的设计疏漏而泄露。

- **桌面端体验不佳**
  - **多 Profile 管理混乱**: 桌面版用户对于切换 Profile 后会话数据错乱的问题感到困扰，这破坏了核心的用户数据组织和访问逻辑。
  - **浏览器检测不智能**: macOS 用户抱怨首次启动时，系统不会自动识别本地安装的 Chrome 浏览器，需要手动配置环境变量，体验不顺畅。

- **配置管理的呼声**
  - **复杂的 Discord 人格配置**: 高级用户希望有一种更清晰、外部化的方式来管理多个 Discord 频道的不同角色和系统提示词，而非依赖复杂的单一配置文件。

## 8. 待处理积压

以下为长期开放、未获得及时响应的重要 Issue 或 PR，提醒维护者关注：

- **PR #12794** `feat(delegate_task): per-subagent model/provider overrides...` (创建于2026-04-20，距今约2个月)
  - **重要性**: **高**。它提出了为子代理单独指定模型/提供商的特性，是增强 Hermes Agent 中复杂任务委托与编排能力的关键功能。
  - **状态**: 超过60天无核心成员评论或进展标记，有被淹没的风险。

- **PR #19331** `feat: add source-bound cognee query tool` (创建于2026-05-03，距今约1.5个月)
  - **重要性**: **中**。整合外部知识库工具，增强 Agent 对特定领域知识的处理能力。
  - **状态**: 长期 Open，无互动。

- **PR #24923** `fix(clarify): treat timeout as refusal instead of implicit consent` (创建于2026-05-13，距今1个月)
  - **重要性**: **高**。这是一个关键的**安全**倾向修复。它将“超时”从“默认同意”改为“默认拒绝”，对于安全敏感的 Agent 操作至关重要。尽管已标记 P1，但长时间未被合并可能会带来潜在的安全风险。

**总结**：项目今日活跃度极高，同时伴随着高优先级的 Bug 报告和大量待处理的 PR。核心团队当前可能需要优先解决 OAuth 和安全绕过这两个 P1/P2 级别的障碍，并加速跟进如 delegate_task 扩展和超时安全修复等积压已久但价值重大的 PR。

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

## PicoClaw 项目日报 — 2026-06-18

### 1. 今日速览
过去24小时内，项目共处理4条Issue和10条Pull Request，整体活跃度较高。其中2条新Issue（#3088、#3093）均为功能请求，2条Bug/安全相关Issue已关闭。PR方面，6条已合并/关闭（含安全修复、Gemini API兼容性修复、搜索解析修复等），4条待合并（含子代理重复消息修复、日志增强、Delta Chat网关等功能）。无新版本发布。项目在安全性、API兼容性和社区功能扩展上均有实质推进。

---

### 2. 版本发布
无。

---

### 3. 项目进展
今日合并/关闭的重要PR：

- **安全修复：OneBot 入站媒体获取限制** — PR #3140 修复了攻击者可通过控制媒体URL使PicoClaw主机发起任意网络请求的漏洞（对应Issue #3070）。该PR复用了已有的HTTP守卫逻辑，阻断私有网络和元数据地址的请求。
- **Gemini 3.5 Flash 兼容性修复** — PR #3136 在工具调用请求体中同时设置 camelCase 和 snake_case 的 `thought_signature` 字段，解决Gemini 3.5 Flash Agentic reasoning所需的格式问题（对应Issue #3111）。
- **搜狗搜索解析修复** — PR #3139 修复因搜狗WAP搜索页HTML结构变化导致的解析失败。
- **Web UI会话历史显示修复** — PR #2990 修复用户只能看到最后一条用户消息的bug，完整读取会话历史。
- **NEAR AI Cloud 提供商支持** — PR #2917 增加对NEAR AI Cloud的OpenAI兼容LLM提供商支持，含模型列表获取和配置向导。

这些修复与功能新增使PicoClaw在安全性、多模型兼容性和用户体验上均向前迈进一步。

---

### 4. 社区热点
- **Issue #3088** — [使用vodozemac替换libolm](https://github.com/sipeed/picoclaw/issues/3088)：获得2个👍，社区关注度高。libolm已无人维护且存在安全问题，替换为官方推荐库vodozemac是提升加密可靠性的关键诉求。项目已标记为`help wanted`和`priority: high`，但尚未有对应PR。
- **Issue #3093** — [需求SimpleX或Tox网关](https://github.com/sipeed/picoclaw/issues/3093)：用户明确需求集成SimpleX、Wire或Tox协议，背后是对去中心化通信和隐私保护的需求。当前PR #3063（Delta Chat网关）仍在开放中，可能与这类需求形成互补。

---

### 5. Bug 与稳定性
今日报告的Bug均已关闭或已有修复PR：

| 严重程度 | Issue/PR | 描述 | 状态 |
|----------|----------|------|------|
| **高危（安全）** | #3070 / PR #3140 | OneBot入站媒体URL允许主机端任意请求，可SSRF攻击 | 已关闭，已合并修复 |
| **中危（API兼容）** | #3111 / PR #3136 | Gemini 3.5 Flash 因缺少`thought_signature`字段返回400错误 | 已关闭，已合并修复 |
| **低危（搜索解析）** | PR #3139 | 搜狗搜索解析因HTML结构变化失败 | 已合并修复 |

其他稳定性改进：PR #3142（子代理重复消息推送）、PR #3141（Brave搜索空结果诊断日志）等仍在开放中，将在后续版本中提升可靠性。

---

### 6. 功能请求与路线图信号
- **加密库替换**（#3088，高优先级）：替换libolm为vodozemac，已获社区认可但尚未开发。该变更影响所有使用端到端加密的功能模块。
- **新通信协议网关**（#3093）：用户希望集成SimpleX、Tox等隐私优先协议；已有PR #3063（Delta Chat网关）开放中，表明项目有意扩展网关生态，可能纳入下一版本。
- **NEAR AI Cloud提供商**（PR #2917已合并）：标志项目正式支持TEE（可信执行环境）能力的AI云，可切入企业级隐私计算场景。

---

### 7. 用户反馈摘要
- **#3088** 用户`pbsds`明确指出现用libolm不安全且无人维护，建议编译时可选，降低迁移风险。
- **#3093** 用户`Damian-o2`直接请求“I need SimpleX or Tox”，反映出对去中心化通信的强烈需求，但未提供具体使用场景。
- **#3070** 安全研究人员`YLChen-007`提交了详细的安全洞见报告（SSRF攻击向量），促使项目快速响应并合并修复。

总体用户反馈集中在安全增强和协议扩展方面，对项目响应速度（24小时内关闭安全漏洞）表示认可。

---

### 8. 待处理积压
以下Issue/PR长期未响应或接近停滞，建议维护者关注：

| ID | 标题 | 状态 | 最后更新 | 停滞原因 |
|----|------|------|----------|----------|
| #3092 | fix(skills_install): add ok checks for version and force type assertions | 开放 | 2026-06-10 | 已标记`stale`，未获维护者review |
| #3063 | feat: add deltachat gateway | 开放 | 2026-06-08 | 功能分支长期未合并，存在冲突风险 |
| #3093 | [Feature] I need SimpleX or tox | 开放 | 2026-06-10 | 仅1条评论，缺乏实现细节和团队反馈 |

此外，建议高优先级Issue #3088（vodozemac替换）在路线图中明确时间窗口，避免libolm安全问题长期暴露。

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

好的，这是为您生成的 NanoClaw 项目动态日报。

---

## NanoClaw 项目动态日报 | 2026-06-18

### 1. 今日速览

今日 NanoClaw 项目活跃度极高，共计处理了 5 个 Issues 和 20 个 Pull Requests，并发布了两个包含破坏性变更的新版本。社区贡献热情高涨，主要集中在 **解决由单一节点故障引发的系统级崩溃**、**修复安全漏洞(CWE-22路径遍历)** 以及 **提升容器环境与文档的用户体验**。尤其是在 **会话故障隔离 (PR #2797)** 和 **安全限制 (PR #2799, #2800)** 方面的修复，表明项目正重点加固稳定性和安全性，整体向生产级可用迈出坚实一步。

### 2. 版本发布

今日发布了 **v2.1.17** 和 **v2.1.0** 两个版本，均包含破坏性变更，用户需特别注意。

- **v2.1.17**: [查看发布说明](https://github.com/nanocoai/nanoclaw/releases/tag/v2.1.17)
    - **核心更新**：**`@onecli-sh/sdk` 从 0.5.0 升级到 2.2.1**。
    - **破坏性变更 (BREAKING)**：此 SDK 版本仅兼容 `/v1` API 的 OneCLI 服务器。旧版 OneCLI 服务器（仅支持旧 API）将无法响应任何 SDK 调用。**官方（sanctioned）的 OneCLI 网关和 CLI 版本已被固定**，以强制要求管理员升级服务端。用户需要确保部署的 OneCLI 版本支持新 API，否则所有通过 SDK 发起的交互都将失败。

- **v2.1.0**: [查看发布说明](https://github.com/nanocoai/nanoclaw/releases/tag/v2.1.0)
    - **核心更新**：引入了 **启动升级标记 (Startup upgrade marker)** 机制。
    - **破坏性变更 (BREAKING)**：现在，服务主机在启动时会检查 `data/upgrade-state.json` 文件。如果该文件记录的信息表明安装状态未达到当前版本，主机将**拒绝启动 (refuses to boot)**。这是为了确保系统状态一致性，防止因升级中断或状态不匹配导致的未定义行为。
    - **迁移注意事项**：用户只需正常通过新版本的安装/升级流程（如运行 `bash nanoclaw.sh`）即可自动创建或更新此标记文件。

### 3. 项目进展

今日项目核心进展体现在成功合并的 PR 上，解决了两个关键的系统性问题：

- **修复会话级故障的“雪崩”效应**: [#2797 [CLOSED] fix(delivery): isolate per-session failures](https://github.com/nanocoai/nanoclaw/pull/2797)
    - **问题**：当某个 agent 的 `outbound.db` 出现瞬时读取故障时，会导致**整个消息投递任务 (delivery tick) 完全停止**，影响所有其他 agent 的消息收发，直到服务重启。这是 Issue [#2796](https://github.com/nanocoai/nanoclaw/issues/2796) 报告的重大稳定性问题。
    - **解决**：通过捕获单个会话处理中的异常，将故障隔离在特定会话内，确保一个不健康的会话不会影响全局消息投递。此 PR 的合入是项目走向高可用架构的关键一步。

- **恢复托管集群的LLM认证**: [#2794 [CLOSED] fix(providers): restore env-var gateway auth for managed fleets](https://github.com/nanocoai/nanoclaw/pull/2794)
    - **背景**：在 v2.1.17 版本后，所有运行在不可变 VM 镜像（managed-fleet）中的 agent 无法进行 LLM（大语言模型）认证，报错 `401 No credentials config`。此 PR 修复了这一回归，保证托管集群用户能正常使用。

### 4. 社区热点

今日社区讨论主要围绕**文档改进**和**安全问题**，体现了社区用户对项目易用性和安全性的高度关注。

- **最活跃的议题组**: 贡献者 `specterslient95-lgtm` 今日集中提交了 4 个 Issues（#2791, #2789, #2787, #2785）和 4 个对应的修复 PR（#2792, #2790, #2788, #2786），全部聚焦于 `.claude/skills` 目录下的文档缺陷。
    - **[#2791](https://github.com/nanocoai/nanoclaw/issues/2791)**: `add-imessage` 技能因缺少 `mkdir` 命令导致文件重定向失败。
    - **[#2789](https://github.com/nanocoai/nanoclaw/issues/2789)**: `setup` 技能文档过于简单，缺乏具体的问题解决步骤。
    - **[#2787](https://github.com/nanocoai/nanoclaw/issues/2787)**: `init-onecli` 文档的端口（10254）声明缺失。
    - **[#2785](https://github.com/nanocoai/nanoclaw/issues/2785)**: `migrate-nanoclaw` 文档标题不具描述性。
    - **分析**：这批议题反映了**项目文档在新用户 onboarding 和操作指引上的薄弱环节**。尽管社区贡献者迅速提供了 PR，但这也暴露了项目文档质量参差不齐，缺乏统一规范的问题。

- **安全相关 PR 受关注**: 来自 `sturdy4days` 的一系列安全修复 PR (#2799, #2800) 虽然没有产生大量评论，但其修复的 **CWE-22 路径遍历**和**任意文件读取**风险，是项目安全领域的重大更新，是今日最具技术深度和影响力的热点之一。

### 5. Bug 与稳定性

今日报告的 Bug 主要集中在以下几个方面，按严重程度排列：

- **严重 (Critical) - 已有 Fix PR**:
    - **[#2796](https://github.com/nanocoai/nanoclaw/issues/2796)**: **会话故障全局阻塞**。如前所述，单一节点故障导致全局消息投递服务宕机。**现已通过 PR #2797 修复**，状态已关闭。
- **高 (High) - 已有 Fix PR**:
    - **[#2804](https://github.com/nanocoai/nanoclaw/pull/2804) (PR)**: `ncl messaging-groups create` 命令因数据库 `NOT NULL` 约束而**完全不可用**。
    - **[#2799](https://github.com/nanocoai/nanoclaw/pull/2799) (PR)**: `send_file` 功能允许**任意文件读取**，存在数据泄露风险。
    - **[#2800](https://github.com/nanocoai/nanoclaw/pull/2800) (PR)**: `ncl groups create` 命令存在**路径遍历漏洞**（CWE-22），允许创建文件夹逃逸到系统目录。
- **中 (Medium) - 已有 Fix PR**:
    - **[#2802](https://github.com/nanocoai/nanoclaw/pull/2802) (PR)**: Socket 客户端存在**无限期等待**和**无上限响应缓冲**问题，可能导致资源耗尽。
    - **[#2801](https://github.com/nanocoai/nanoclaw/pull/2801) (PR)**: JSON 解析器在处理非对象原始类型（如 `"5"`, `true`）时行为异常，可能导致后续逻辑错误。
- **低 (Low) - 已有 Fix PR**:
    - **[#2791](https://github.com/nanocoai/nanoclaw/issues/2791)**: `add-imessage` 技能因目录不存在而失败。

### 6. 功能请求与路线图信号

- **Agent 间消息审批策略**: **[PR #2793](https://github.com/nanocoai/nanoclaw/pull/2793) (Open)** 提出了一个极具潜力的新特性：在 agent 间通信链路上引入**可选的、按消息级别的审批策略**。这为构建需要人工介入或安全审计的复杂 AI 工作流提供了基础，信号明确，**极有可能被纳入下一个次要版本**。

- **CLI 仪表盘技能**: **[PR #2795](https://github.com/nanocoai/nanoclaw/pull/2795) (Open)** 新增了一个名为 `/add-clidash` 的只读仪表盘技能，允许用户通过 CLI 查看代理状态、任务进度等，这能显著改善运维体验。

- **新增 AI 后端**: **[PR #2717](https://github.com/nanocoai/nanoclaw/pull/2717) (Open)** 提议将 Atlas Cloud 作为新的 OpenAI 兼容 LLM 后端。这表明社区正在积极拓展项目可用的 AI 模型生态，增加用户选择。

### 7. 用户反馈摘要

从今日的 Issues 和 PR 评论中，可以提炼出以下用户痛点：

- **良好的社区协作**：`specterslient95-lgtm` 用户发现的文档问题后，项目维护者（如 `sturdy4days`）和其他参与者迅速响应，确认问题并讨论解决方案，展现了良好的社区治理氛围。
- **文档不完善是核心痛点**：多个报告指向了 `SKILL.md` 文件质量不佳，如内容过于简略、关键步骤缺失、表述不清晰等问题。这直接影响了新用户的上手效率和成功率，是当前项目的关键短板。
- **对新手的“零配置”幻灭**：用户对 `setup` 技能的期望是能获得一步到位的引导和常见问题排查路径，而非仅指向一个脚本。这表明用户渴望更完善的 Onboarding 体验。

### 8. 待处理积压

- **长期未解决的修复工作**: **[PR #2750](https://github.com/nanocoai/nanoclaw/pull/2750) (Open)**，旨在修复容器被强制杀死后 `outbound.db` 日志文件损坏的问题。此 PR 从 6月12日开启，至今已有 6 天，且关联了两个重要的 Issue（#2516, #2640）。虽然今天有更新，但仍未合并。此问题影响容器环境的稳定性，建议维护者优先关注和评审。

- **长期待处理的文档更新**: **[PR #2717](https://github.com/nanocoai/nanoclaw/pull/2717) (Open)**，自 6月9日开启已有一周多，内容为新增 Atlas Cloud 作为 LLM 后端的文档。此 PR 不涉及代码变更，评审周期较长，可能反映了项目对新集成方案的审核较为谨慎，或维护者人手不足。

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

# NullClaw 项目动态日报 — 2026-06-18

## 1. 今日速览
过去24小时，NullClaw社区保持活跃，有3个Issue获得更新（均为早期问题，无新关闭），2个新提交的PR待审查，其中包含一项重要的内存召回功能扩展和一项CLI交互修复。项目暂无新版本发布，整体处于功能迭代与问题修复并行阶段。社区对CLI体验和Web UI易用性的反馈持续，贡献者积极提出解决方案。

## 2. 版本发布
（无）

## 3. 项目进展
今日未合并或关闭任何PR，但有两项重要PR新提交，标志着项目在以下方面向前推进：
- **CLI交互修复**：PR #960 解决了长期存在的方向键控制字符显示问题，通过引入轻量级行编辑器替代标准输入处理，预期将提升终端用户使用体验。
- **内存系统可配置化**：PR #961 为memory模块新增三个JSON配置项（auto_recall、recall_limit、max_context_bytes），允许用户精细控制记忆召回行为，是响应社区对内存管理灵活性需求的直接成果。

以上PR均待维护者审查合并，项目整体健康度良好，贡献者响应及时。

## 4. 社区热点
- **Issue #915**（评论2，👍0）"Problem with scheduler unauthorized"：用户scabros报告在Ubuntu上使用外部Ollama Qwen3.6:27b时调度器无法工作（Telegram与CLI均受影响）。该问题自5月15日创建，6月17日有更新，但尚未有官方回复或修复PR，属于用户痛点较高的阻塞性Bug。
- **Issue #865**（评论2，👍0）"CLI shows ctrl characters for up/down/left/right keys"：该问题已存在近两个月，今日PR #960 正是针对此问题的修复尝试，社区关注度较高，反映了对CLI基础交互的强烈需求。
- **Issue #861**（评论1，👍0）"How to enable the Web UI on headless VPS server?"：用户eabase明确请求更易懂的Web UI部署文档，暴露出README中关于Web UI的说明对非专业用户不够友好。

## 5. Bug与稳定性
以下为今日活跃的Bug报告，按严重程度排列：

| 严重程度 | Issue | 摘要 | 是否有Fix PR |
|---------|-------|------|-------------|
| **高** | #915 | 调度器因未授权无法工作，影响自动化任务与Telegram集成 | 无 |
| **中** | #865 | CLI方向键显示乱码，影响基础命令行操作 | 是（PR #960） |

此外，Issue #861 虽非直接Bug，但造成用户部署障碍，可归类为文档缺陷。

## 6. 功能请求与路线图信号
- **PR #961** (feat(memory): auto_recall, recall_limit, max_context_bytes)：用户valonmulolli提交了内存召回机制的可配置化方案，允许关闭自动召回、限制召回条目数和上下文token数。该功能若合并，将直接回应用户对内存消耗控制的需求，很可能被纳入下一版本。
- 社区未在今日提出全新功能请求，但Issue #861 隐含对Web UI部署简化的需求，建议维护者考虑在文档或配置向导中提供更清晰的指引。

## 7. 用户反馈摘要
- **痛点**：
  - 调度器在Ollama+lora环境中无法工作，用户scabros在#915中详细描述了配置与日志，期望获得根本性修复。
  - CLI终端基础键绑定（方向键、退格、Home/End）异常，用户eabase在#865中强调该问题影响日常使用，期望行为符合终端标准。
  - Web UI部署流程晦涩，用户eabase在#861中明确表示“70%的README看不懂”，希望用“人类语言”描述如何通过隧道访问浏览器界面。
- **满意点**：暂无正面反馈，但两例用户均提供了清晰的复现步骤和耐心等待，表明社区对新功能（如调度器、Web UI）有较高期待。

## 8. 待处理积压
以下为长时间未获得官方响应或修复的重要Issue，建议维护者优先关注：

| Issue | 描述 | 创建时间 | 最后更新 | 建议行动 |
|-------|------|---------|---------|---------|
| #915 | 调度器未授权错误，影响核心自动化能力 | 2026-05-15 | 2026-06-17 | 确认是否与Ollama认证机制兼容，考虑添加调试日志或提供临时解决方案 |
| #865 | CLI方向键乱码，已有PR #960 但尚未合并 | 2026-04-23 | 2026-06-17 | 尽快审查并合并PR #960，或给出替代修复方案 |
| #861 | Web UI部署文档过于技术化 | 2026-04-22 | 2026-06-17 | 在README中添加简化版向导或提供示例配置脚本，降低门槛 |

---

*数据来源：GitHub NullClaw/NullClaw 仓库，截至2026-06-18 23:59 UTC。*

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

好的，这是为您生成的 IronClaw 项目动态日报。

---

# IronClaw 项目动态日报 — 2026-06-18

## 1. 今日速览

今日项目活跃度**极高**。过去 24 小时内产生了 50 条 PR 和 11 条 Issue，表明核心团队和社区贡献者都在密集推进开发。**“Projects”特性**作为一项重量级更新，其五个分阶段的 PR 栈已全部提交，标志着项目在数据模型和协作功能上迈出重要一步。社区反馈集中在 WebUI 体验和模型兼容性上，多个关于 UI 反馈、错误处理和 Slack OAuth 安全性的问题已被关闭或正在修复。项目整体处于新功能密集开发与稳定性并行修复的健康状态。

## 2. 版本发布

本日无新版本发布。

## 3. 项目进展

本日合并/关闭了 17 条 PR，主要集中在功能完善与 Bug 修复。最关键的进展是 **“Projects” 特性** 的完整 PR 栈（#5015-#5019）被提交，标志着 IronClaw Reborn 引入了独立于旧引擎的**项目**实体和完整的 CRUD + 成员管理功能，为团队协作奠定了基础。

其他重要进展包括：
- **输出感知的“无进展”检测**：PR #5022 合并，引入了 `ContentDigest` 以识别对话是否陷入循环，提升了 Agent 的智能决策能力。
- **Slack OAuth 安全加固**：PR #5052 关闭 Issue #5009，为**非触发的**（live）Slack OAuth 路径也增加了结构性的 DM 限制，实现了与触发式路径的安全对等。
- **自动化错误软化**：PR #5055 (待合并) 提出将自动化运行失败的状态从红色“错误”改为黄色“需注意”，优化了前端用户体验。
- **Bedrock 模型支持恢复**：PR #5059 针对 Issue #5058 修复了 AWS Bedrock 在独立二进制文件中不可用以及 Converse 工具架构的问题。

## 4. 社区热点

本日最受关注的讨论聚焦于新功能的实现与安全优化。

- **`Sprint: IronClaw Reborn Local Dogfooding`** (Issue #4879): 作为一周记录，持续跟踪本地自用中发现的问题，是社区质量反馈的重要来源。
- **`[security] Bring live Slack OAuth path to structural DM-parity`** (Issue #5009): 对 Slack OAuth 安全性的深度讨论，引发了 PR #5052 的快速跟进修复，体现了项目对安全的重视。
- **`[Reborn] GitHub analysis workflows may enter repeated approval loops`** (Issue #5060): 报告了一个可能导致工作流卡死的问题，需要项目组关注。
- **`feat(reborn): skill extraction & self-evolution with activation controls`** (PR #5061): 一个大型 XL 规模的 PR，引入了技能提取和自我进化能力，引起了广泛关注，尽管是新提交，但代表了项目在 Agent 智能化方向上的探索。

## 5. Bug 与稳定性

今日报告的 Bug 主要集中在 WebUI 体验和模型兼容性方面，严重程度中等。

- **[Bug] GitHub 分析工作流卡死** (Issue #5060): 报告 CI 工作流可能陷入重复审批循环，最终不产生任何结果。这是一个影响 CI 效率的问题。
- **[Bug] Reborn 独立二进制无法使用 AWS Bedrock** (Issue #5058): 功能性 Bug，阻止用户使用 Bedrock 作为模型后端。**已有修复 PR #5059**。
- **[Bug] “auto” 模型ID被 cloud-api 拒绝** (Issue #5044): 配置 `NEARAI_MODEL=auto` 会导致 HTTP 400 错误，影响桌面端用户。**已有修复 PR #5045**。
- **[Bug] UI 反馈缺失与显示错误** (Issues #4823, #5007, #4974): 多个 UI 相关问题，包括删除对话无反馈、验证错误不消失、操作按钮重复等，影响用户体验。

## 6. 功能请求与路线图信号

本日没有直接的用户功能请求。但通过活跃的 PR 可以判断出项目的研发方向。

- **“Projects” 特性** (PR #5015-#5019): 这是一个明确的路线图信号，表明 IronClaw 正在构建自己的项目管理能力。
- **Agent 自我进化** (PR #5061): “技能提取”功能表明项目正在探索让 Agent 从对话历史中学习并形成可复用技能。
- **输出感知的进度检测** (PR #5022): 提升 Agent 的智能逻辑，避免陷入无效循环，这是 Agent 核心智能性的关键优化。

这些 PR 大部分是针对内部路线图的实现，而非社区直接反馈，但它们将是未来版本的核心特性。

## 7. 用户反馈摘要

从今日的 Issue 和 PR 评论中，可以提炼以下用户反馈：

- **常见痛点**：
  - **删除运行中的对话无反馈** (Issue #4823): 用户期望每次操作都有明确的视觉反馈。
  - **表单验证错误不消失** (Issue #5007): 用户期望更智能的即时校验反馈，而不是提交后才清除错误标记。
  - **Activity 界面按钮冗余** (Issue #4974): UI 界面混乱，操作按钮重复，影响使用。
  - **模型配置问题** (Issue #5044): 用户通过桌面端配置后，因模型名“auto”无效导致服务不可用，体验较差。

- **正面反馈**：
  - 社区对 OAuth 安全性的关注和快速修复得到了积极反响 (通过 #5009 的关闭)。

## 8. 待处理积压

以下 Issue/PR 更新较少，可能被忽略或需要社区维护者关注：

- **`[Reborn] 首次引导应阻止访问扩展和自动化`** (Issue #4793): 此问题虽已关闭，但讨论了首次引导的体验设计。相关讨论和决策可能需要跟进文档化或实现。
- **大量依赖更新 PR** (Issue #4876): 这是一个大型依赖更新 PR，涉及 43 个 Rust crate 更新。此类 PR 虽然重要，但通常需要较长审查时间以避免兼容性问题。维护者需关注其审查状态，避免积压过久导致与 `main` 分支冲突加剧。

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

好的，作为一名 AI 智能体与个人 AI 助手领域开源项目分析师，我将根据您提供的数据，为您生成 LobsterAI 项目的 2026-06-18 项目动态日报。

---

# LobsterAI 项目动态日报 (2026-06-18)

**分析日期:** 2026年6月18日
**数据来源:** github.com/netease-youdao/LobsterAI

### 1. 今日速览

项目今日活跃度极高。过去24小时内，开发团队合并了11个 Pull Request，并发布了重要新版本，展现出强大的迭代能力。核心开发团队（特别是 `liuzhq1986`）正聚焦于 Cowork 模式的稳定性、上下文连续性及细节优化，同时推进了 Computer Use 和实时 ASR 语音输入等重磅功能。**项目健康度非常好，处于快速演进阶段。** 但与此同时，社区反馈了一个关于自动工件加载的高危安全漏洞，需要团队优先响应。

### 2. 版本发布

**新版本：LobsterAI 2026.6.15**

-   **发布时间:** 2026年6月15日
-   **更新亮点:**
    -   **新功能：Computer Use（计算机使用）** ([#2143](https://github.com/netease-youdao/LobsterAI/pull/2143)): 新增了让AI直接操控计算机的能力，这可能是向自主Agent迈出的关键一步。
    -   **新功能：Cowork模式实时ASR语音输入** ([#2148](https://github.com/netease-youdao/LobsterAI/pull/2148)): 在协作模式下引入了实时的语音识别功能，显著提升了交互体验。
    -   **优化：Cowork上下文压缩后连续性增强** ([#2145](https://github.com/netease-youdao/LobsterAI/pull/2145)): 改进了对话历史压缩的质量，使Agent在长对话中能更可靠地维持任务上下文。
-   **破坏性变更：** 未明确提及。
-   **迁移注意事项：** 建议所有用户升级到此版本，以体验新功能和关键的稳定性改进。

### 3. 项目进展

今日合并/关闭了11个PR，其中大部分由核心开发者 `liuzhq1986` 完成，主要聚焦于Cowork模式的稳定性和细节打磨。

-   **核心功能修复与增强：**
    -   **上下文连续性（Context Continuity）:** `#2145` 显著提升了Cowork模式下上下文压缩后的任务连续性，这对长对话和复杂任务至关重要。
    -   **语音输入（Voice Input）:** `#2162` 修复了语音输入取消守卫在合并后失效的问题，`#2148` 新增了实时ASR功能，显示团队正大力投入语音交互。
    -   **滚动与渲染（Scrolling & Rendering）:** `#2174` 修复了自动滚到底部的定位问题，`#2173` 修复了用户消息渲染，现在能正确保留用户输入的换行符。
-   **稳定性与性能：**
    -   **Gateway内存限制（OOM Fix）:** `#2149` 为OpenClaw网关设置了V8旧空间内存限制，以应对长时间、多通道工作负载下的OOM崩溃，这是一个重要的稳定性提升。
    -   **竞态条件修复（Race Condition Fix）:** `#2147` 修复了用户停止操作与启动流程之间的竞态条件，防止在启动阶段被中断时发送错误数据。
-   **其他优化：**
    -   **模型选择（Model Selection）:** `#2153` 修复了同名的Package模型和自定义模型的选择问题。
    -   **文档与信息（Docs & Info）:** `#2175` 优化了README，`#2154` 确保在手动停止流式输出后，模型元数据能正确显示。

**总结：** 项目在修复了多个关键内部状态逻辑问题的同时，通过引入新功能和修复细节，大幅提升了Cowork模式的成熟度和用户体验。

### 4. 社区热点

**最受关注 Issue:** **[Security] LobsterAI automatic artifact loading allows message-derived arbitrary local file reads** (#2176)

-   **链接:** [Issue #2176](https://github.com/netease-youdao/LobsterAI/issues/2176)
-   **热度分析:** 这是今日**唯一**的新Issue，且被标记为 `Security`。报告者指出，LobsterAI 自动解析来自助手或工具的 `MEDIA:` 文件引用，并将结果路径传递给一个特权的Electron进程，从而导致可被利用的任意本地文件读取漏洞。
-   **背后诉求:** 这是一个严重的安全问题。社区希望项目维护者能**立即**关注并修复该漏洞，可能需要修改自动工件加载的逻辑或加入严格的路径校验。当前有一条评论，可能涉及与社区或作者的沟通。

### 5. Bug 与稳定性

| 严重程度 | 描述 | 状态/修复进展 | 链接 |
| :--- | :--- | :--- | :--- |
| **严重 (Security)** | 自动工件加载漏洞，允许通过消息触发任意本地文件读取。 | **待处理**（今日刚报告）。 | [Issue #2176](https://github.com/netease-youdao/LobsterAI/issues/2176) |
| **高** | OpenClaw Gateway进程的OOM崩溃。 | **已修复**（已在`v2026.6.15`中被合入）。 | [PR #2149](https://github.com/netease-youdao/LobsterAI/pull/2149) |
| **中** | 语音输入取消功能在合并后失效；启动时停止操作导致竞态问题。 | **已修复**（`#2162`, `#2147`）。 | [PR #2162](https://github.com/netease-youdao/LobsterAI/pull/2162), [PR #2147](https://github.com/netease-youdao/LobsterAI/pull/2147) |
| **低** | 滚动到底部位置不准确；用户消息渲染丢失换行。 | **已修复**（`#2174`, `#2173`）。 | [PR #2174](https://github.com/netease-youdao/LobsterAI/pull/2174), [PR #2173](https://github.com/netease-youdao/LobsterAI/pull/2173) |

### 6. 功能请求与路线图信号

-   **潜伏信号：Computer Use**
    -   最新版本中已发布的 `Computer Use` 功能（`#2143`）表明，LobsterAI的路线图正向“AI自主操作电脑”的方向迈进。这很可能成为下一个版本的宣传重点，并孕育更多相关的功能和插件。
-   **潜在需求：实时语音交互的深化**
    -   实时ASR语音输入（`#2148`）的加入，可能意味着社区将提出更多关于语音助手场景的请求，如语音唤醒、多语言支持、声纹识别等。

### 7. 用户反馈摘要

基于 Issue #2176 的初始报告和现有数据，用户反馈集中于对**安全性和可靠性的关切**。

-   **痛点:** 报告者明确指出自动加载机制存在安全漏洞，这可能导致严重的隐私和数据泄露问题。这是用户使用此类AI助手工具时最根本的担忧。
-   **使用场景:** 报告者应该是在测试或使用LobsterAI的过程中，尝试让AI工具解析或读取本地文件时发现了问题。
-   **满意/不满意:** 社区对当前漏洞的发现表示担忧（不满意的地方），但其发现本身对项目价值的提升是有帮助的。对于已经合并的诸如OOM修复等稳定性改进，社区应该是满意的，但缺乏具体的评论数据。

### 8. 待处理积压

-   **长期未响应的 PR [#1463](https://github.com/netease-youdao/LobsterAI/pull/1463)**
    -   **状态:** 已标记为`stale`并已关闭。
    -   **摘要:** 修复了模态框标题过长的问题。
    -   **分析:** 该PR由非核心开发者创建，于4月4日提交，在经历了长达2个多月的沉寂后被关闭。这反映出项目团队可能对非核心贡献者或较小的UI修复PR的响应和处理优先级较低。**建议团队成员审查此类PR，如果仍有价值，应尽快合并或给出明确的拒绝理由，以维护贡献者的积极性。**

---

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis 项目日报 — 2026-06-18

## 1. 今日速览

- 过去24小时内，项目保持温和活跃，共新增2个Issue和1个PR，无版本发布，无已合并的PR。  
- 两个新Issue均为增强型功能请求，分别涉及TTS输出格式配置和Markdown复制/导出，社区需求集中在**输出可定制性**方向。  
- 一个待合并PR（#1130）试图使WebUI的RPC超时可配置，直接关联未关闭的Issue #1127，显示维护者对功能扩展的响应较快。  
- 整体来看，Moltis处于**功能积累期**，社区参与度适中，未出现严重Bug或回归报告，项目健康度良好。

> 链接：https://github.com/moltis-org/moltis

---

## 2. 版本发布

**无**（过去24小时无新版本发布）

---

## 3. 项目进展

### 今日无合并/关闭的PR

- 当前仅有的一个开放PR **#1130**（`feat: make webui rpc timeout configurable`）仍处于待合并状态，未完成集成。  
- 该PR由贡献者 khimaros 提交，直接关联 Issue #1127，旨在允许用户自定义WebUI与后端RPC通信的超时时间，对高延迟或复杂查询场景具有重要意义。  
- 项目整体进展体现在**功能扩展准备阶段**：多个增强型Issue和PR被提出，但尚未进入主干，预期后续版本将整合这些能力。

> 链接：https://github.com/moltis-org/moltis/pull/1130

---

## 4. 社区热点

- **Issue #1126** – `[Feature]: allow to configure the format of tts output`  
  - 作者：khimaros | 评论：3条 | 创建后持续活跃  
  - 讨论核心：用户希望TTS（文本转语音）输出格式（如WAV、MP3、OPUS等）可由用户配置。当前硬编码格式限制了某些场景（如语音助手需要低延迟流式输出）。  
  - 分析：这是AI助手常见的痛点，尤其是需要与其他系统集成时。评论中已有用户提出“streaming chunk format”的需求，显示该功能可能带动后续PR。

- **Issue #1131** – `[Feature]: Add copy + export as Markdown`  
  - 作者：vvuk | 评论：0条（新开）  
  - 背景：用户需要将聊天记录或模型输出一键复制/导出为Markdown格式，方便文档整理。  
  - 分析：该需求反映了用户对**可移植性**的重视，与TTS输出格式请求形成互补——都在强调“数据应适配用户工具链”。

> 链接：Issue #1126: https://github.com/moltis-org/moltis/issues/1126  
> 链接：Issue #1131: https://github.com/moltis-org/moltis/issues/1131

---

## 5. Bug 与稳定性

**无新增Bug、崩溃或回归报告**（过去24小时内所有Issue均为增强请求，无错误标签）。

> 说明：当前项目稳定性良好，未发现严重问题。需持续关注后续用户反馈。

---

## 6. 功能请求与路线图信号

| 功能请求 | 关联链接 | 是否有可能纳入下一版本 |
|----------|----------|------------------------|
| 允许配置TTS输出格式 | Issue #1126 | 高可能性。同类需求已有PR #1130（配置化思路），且该Issue有3条评论表示支持。 |
| 添加Markdown复制/导出 | Issue #1131 | 中高可能性。基础UI操作，实现成本低，符合“用户工具链友好”方向。 |
| WebUI RPC超时可配置 | PR #1130 | 极高可能性。PR已写就，仅待合并，预计下一补丁版本即包含。 |

- 信号分析：三大需求均指向**增强配置能力**和**输出适配**，暗示项目下一版本会专注提升用户对输出格式和交互行为的控制力。

---

## 7. 用户反馈摘要

基于 Issue #1126 的3条评论，提炼真实用户声音：

- **痛点**：现有TTS输出格式不可调，导致用户无法直接用于某些应用（如Telegram语音、流式浏览器播放）。用户写道：“I need to pipe the output to a Discord bot, but it only accepts specific audio formats.”
- **使用场景**：多平台集成、语音助手前端、文档标注系统。
- **期望**：支持类似 `--output-format wav` 的命令行参数或WebUI下拉选择，并希望保留原始原始音频的PCM流式输出。

> Issue #1131 暂无评论，但标题已明确表达对“Markdown导出”的渴望，推测用户多为内容创作者或开发者，需将对话记录导入笔记软件。

---

## 8. 待处理积压

**无显著积压**。当前所有Issue和PR均在一周内创建/更新：

- Issue #1126 最后更新于2026-06-17，距今天仅1天，无需提醒。
- Issue #1131 创建于当天（2026-06-17），属于新需求。
- PR #1130 创建于当天，待审。

> 说明：项目维护者响应良好，无长期未响应的关键项。建议保持节奏，优先合并PR #1130以解决RPC超时问题。

---

**总结**：Moltis在过去24小时内展现了健康的功能扩展节奏，社区对输出可配置性的需求集中且清晰。建议尽快合并PR #1130，并考虑将 #1126 和 #1131 标记为 `help wanted` 以吸引更多贡献。

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，我已基于您提供的 CoPaw (github.com/agentscope-ai/CoPaw) GitHub 数据，生成以下项目动态日报。

---

# CoPaw 项目动态日报 | 2026-06-18

## 1. 今日速览

项目今日保持高活跃度，主要围绕 **v1.1.12正式版发布** 以及 **v2.0.0架构升级** 的前期准备工作展开。核心动态包括：社区在经历大规模Issue清理后，焦点重新汇聚至几个关键稳定性 Bug（如进程冻结、死循环）和新功能讨论上。项目发布了两个新版本，并开启了通往2.0时代的里程碑。整体而言，项目健康度良好，维护者与社区的互动高效。

## 2. 版本发布

今日发布了 **2个** 新版本，显示项目迭代速度加快。

- **`v1.1.12` (正式版)**:
  - **新增**: 全面重做了 Console 端的 **模型提供商页面**，支持提供商聚合、统一卡片 UI 和布局重新设计 ([#5203](https://github.com/agentscope-ai/QwenPaw/pull/5203))。新增 **“简单模式”** ，提供扁平导航和按更新时间排序的会话列表，降低了新用户的上手门槛 ([#5222](http://...))。
  - **迁移/破坏性变更**: 本次更新主要为UI增强，暂无破坏性变更报告。

- **`v1.1.12-beta.2`**:
  - **性能优化**: 移除了 Agent 配置中的非必要深拷贝操作，提升了配置加载性能 ([#5240](https://github.com/agentscope-ai/QwenPaw/pull/5240))。
  - **新功能**: 在管理后台 (Console) 增加了按标题过滤会话的功能，方便用户管理大量对话 ([#5178](https://github.com/agentscope-ai/QwenPaw/pull/5178))。

## 3. 项目进展

今日合并/关闭的 PR 修复了多个重要问题，并启动了架构升级的进程。

- **🔥 启动架构升级**: **`#5281`** [merged](https://github.com/agentscope-ai/QwenPaw/pull/5281) 将版本号推至 `2.0.0a1`，标志着 CoPaw 向 AgentScope 2.0 架构的 Alpha 阶段迈出关键一步。
- **ChromaDB 兼容性修复**: **`#5289`** [merged](https://github.com/agentscope-ai/QwenPaw/pull/5289) 重命名了 ChromaDB 运行探针的集合名称，修复了因命名规则不符导致启动时 `InvalidArgumentError` 的问题 (Issues #5284)。
- **发布流程修复**: **`#5288`** [merged](https://github.com/agentscope-ai/QwenPaw/pull/5288) 修复了版本发布脚本中的参数扩展问题，确保了未来版本发布的顺畅。
- **小易 (XiaoYi) 频道修复**: **`#5274`** [merged](https://github.com/agentscope-ai/QwenPaw/pull/5274) 对小易渠道进行了重大重构，将其连接机制改为双 WebSocket 架构，解决了该渠道不可用的问题，并适配了官方协议。
- **备份稳定性提升**: **`#5041`** [merged](https://github.com/agentscope-ai/QwenPaw/pull/5041) 修复了备份功能中，当遇到无法读取的文件时，整个备份过程会直接失败的Bug，改为跳过不可读文件，提升了备份的鲁棒性。

## 4. 社区热点

今日讨论最活跃的议题揭示了社区在 **Agent自主性与可靠性** 方面的核心关注点。

- **子Agent上下文压缩导致进程冻结 (#5218)**: [Open](https://github.com/agentscope-ai/QwenPaw/issues/5218)
    - **分析**: 该Bug报告了子Agent在执行上下文压缩时，会导致整个QwenPaw进程完全冻结，只能强制重启。这是**严重稳定性问题**。评论区有16条讨论，社区正在积极分析根因，并已有PR尝试修复。这暴露了当前上下文管理机制在高负载或特定场景下的脆弱性。

- **跨Agent通信无限循环 (#5204)**: [Open](https://github.com/agentscope-ai/QwenPaw/issues/5204)
    - **分析**: 报告了两个通过Matrix频道互聊的Agent会陷入无限循环的问题。这是**Agent交互模式下的一个核心缺陷**，揭示了运行时层面缺乏防止此类“双向唤醒链”的机制。用户（laeni）清晰地区分了此场景与单Agent内部的ReAct死循环，指出了问题的特殊性，对项目路线图有重要参考价值。

## 5. Bug 与稳定性

今日新增/活跃的Bug Issue主要集中在运行时稳定性和UI/UX缺陷上。

| 严重程度 | Issue | 描述 | 状态 |
| :--- | :--- | :--- | :--- |
| **严重** | [#5218](https://github.com/agentscope-ai/QwenPaw/issues/5218) | 子Agent上下文压缩导致进程冻结 | 活跃讨论，已有相关PR |
| **严重** | [#5204](https://github.com/agentscope-ai/QwenPaw/issues/5204) | 两个Agent交互时陷入无限循环 | Open |
| **中** | [#4967](https://github.com/agentscope-ai/QwenPaw/issues/4967) | 执行过程陷入死循环无法退出 | Open |
| **中** | [#5264](https://github.com/agentscope-ai/QwenPaw/issues/5264) | 群聊回复错误地发到私聊 (飞书) | Open |
| **低** | [#5262](https://github.com/agentscope-ai/QwenPaw/issues/5262) | 升级后，被禁用的内置技能会重新启用 | Open |
| **低** | [#5284](https://github.com/agentscope-ai/QwenPaw/issues/5284) | ChromaDB探针因命名规则失败 | **已修复** (PR #5289已合并) |
| **低** | [#5292](https://github.com/agentscope-ai/QwenPaw/issues/5292) | 手动添加的模型不显示在对话选择中 | Closed |

## 6. 功能请求与路线图信号

从今日的Issues和PR中，可以捕捉到几个明确的功能需求信号，部分已有实现或正在开发中。

- **内置技能/MCPs预装**: Issue **#280** (讨论内建技能) 重新活跃，社区对提升“开箱即用”体验有强烈需求。
- **配置迁移工具**: PR **#5276**  (Open) 提出了 `qwenpaw migrate openclaw` 的命令行工具，用于从OpenClaw/Hermes Agent生态系统迁移配置。这表明项目正在积极吸引其他平台的用户。
- **定时任务管理增强**: PR **#5210** (Open) 添加了 `qwenpaw cron update` 命令，允许用户直接修改现有定时任务，替代“删除-重建”的繁琐流程。同时，PR **#5241** (Open) 将定时任务的 `misfire_grace_seconds` 从60秒延长至3600秒，解决因任务繁忙导致的任务被跳过问题。这些信号显示定时任务功能正在快速完善。
- **HTTP API公开**: Issue **#2202** 再次被提及，用户 (MagcirHu) 希望公开HTTP API以摆脱Web UI的限制，实现自动化多Agent对话。这是一个高级用户需求，可能影响未来的集成方式。

## 7. 用户反馈摘要

- **进程冻结 (严重痛)**: 用户 `malongan` 报告子Agent上下文压缩导致“完全无响应”，需要手动重启，严重影响了使用体验。这可能是当前版本最严重的稳定性问题。
- **升级困扰**: 用户 `daigoopautoy` 和 `GroAries` 多次抱怨升级或下载新技能后，原有的禁用状态会被重置，导致每次都需手动调整，感到非常“繁琐”。这是一个影响长期体验的UX问题。
- **对AI写代码的接纳**: 用户 `mofeiss` 在Issue **#2677** 中提出一个尖锐问题，即社区是否欢迎使用 AI Agent (如Codex/Claude Code) 编写的 PR。该用户表达了因“用AI写PR”而被其他项目维护者“羞辱”的经历，显示出在开源社区中，对于AI辅助编程的接受度和文化正在形成。

## 8. 待处理积压

- **跨Agent无限循环 (#5204)**: 这是一个关乎 **Agent间交互范式** 的根本性问题。在目前的框架下，如何防止Agent间的消息产生“递归”或“反馈回路”尚待解决。此问题对于构建复杂的、多Agent协作系统至关重要，建议维护者优先评估并将其纳入路线图讨论。
- **升级后技能重置问题 (#5262, #3090)**: 这是一个已报告多次的“甜蜜点”Bug，虽然严重度不高，但持续影响大量用户的实际使用体验。建议在下个小版本中优先修复。

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，我将根据您提供的 ZeroClaw 项目数据，生成一份结构清晰、数据驱动的项目动态日报。

---

## ZeroClaw 项目动态日报 | 2026-06-18

### 1. 今日速览

ZeroClaw 项目今日处于**高度活跃的开发期**，社区提交与合并活动频繁。过去 24 小时共产生 50 条 Pull Request，其中 11 条已合入主干，显示出核心维护团队正积极推动多个功能模块的进展。同时，17 条 Issue 的更新表明社区反馈活跃，但风险等级为 **high** 的议题占比偏高（新开的 16 条 Issue 中，过半被标记为 high risk），项目在快速迭代中正面临**稳定性与跨平台兼容性**方面的集中挑战。总体来看，项目正处于一个功能密集推进与 Bug 修复并行的关键阶段。

### 2. 版本发布

无新版本发布。

### 3. 项目进展

今日合并或关闭了 11 个 Pull Request，主要推进了核心功能的几个重要方面：

-   **核心配置与状态管理**：完成了关于**别名条目级联删除与重命名**的系列 PR（#[7841] 与 #[7840]）。这是对一个长期存在的核心功能（Issue #7175 与 #7468）的实现，使得修改 Agent、Provider 等别名时，其关联的配置与状态能同步更新，为后续 CRUD 功能奠定了基础。
-   **ACP 协议与用户体验**：合并了 PR #[7684]，该修复将 ACP 会话中的**历史修剪器**和**对话取消**操作作为可见事件呈现给用户，不再悄然静默处理，显著提升了用户在 ACP 会话中的感知和控制体验。
-   **CLI 工具链**：合并了 PR #[7842]，为 CLI 添加了对 Agents、Providers、Channels 的 CRUD 操作以及技能包的级联管理功能，让用户在命令行下管理配置更加得心应手。

这些合并推进了 **v0.8.2/v0.8.3** 里程碑的进展，特别是关于 “Skills Platform” 和 “MCP Dashboard” 的基础能力。

### 4. 社区热点

今日讨论最活跃的议题主要集中在项目架构方向的重大规划上：

-   **[RFC] Native context compression as a provider pipeline decorator (#7673)**：虽然评论数只有 3 条，但这是一个关于**架构级别优化**的 RFC。讨论围绕在 Provider 层引入压缩装饰器以减少 token 消耗和延迟的提议，触及到核心性能问题，是社区高级用户和核心开发者关注的焦点。
-   **[Tracker]: v0.8.1 integration/channel/provider/tool queue and history (#6970)**：作为 **v0.8.1** 版本的进展追踪，持续获得关注。它汇总了多个 Channel、Provider 和 Tool 的集成工作，是社区了解项目短期迭代方向的核心窗口。
-   **[Support]: Android Termux Setup (#7911)**：这是一个新提出的支持请求，首次出现了**在 Android Termux 上运行 ZeroClaw** 的需求。虽然目前无人评论，但这代表了从移动端服务器运行 AI Agent 的新场景，可能会吸引一部分对私有部署和边缘计算感兴趣的开发者。

### 5. Bug 与稳定性

今日报告的 Bug 中，有两个 `S1 - workflow blocked`（严重级别）的问题，值得高度关注：

-   **【严重】Agent 重命名导致状态写入时序问题 (#7907，S1)**：`rename_agent_cascade()` 函数在执行重命名时，会先修改外部状态，然后才持久化配置。如果中间步骤失败，会导致状态与配置不一致，阻塞后续工作流。**目前暂无直接的 fix PR**，但应结合刚合并的 PR #7841 进行审查。
-   **【严重】Canvas 回归问题 (#7563，S1)**：该问题在 #6986 之后出现，导致 WebSocket 会话中使用 Canvas 工具时，Web 界面 `/canvas` 显示空白。该 Issue 已于今日**关闭**，表明修复已合并，是一个值得庆祝的稳定性回正。
-   **【高】Windows 平台测试失败 (#7462，S2)**：在 Windows 11 上，有 **74 个测试用例**失败，根源在于 Unix-only 的命令、路径语义和终端编码问题。CI 未覆盖此场景，是目前平台兼容性的重大缺口。关联的修复 PR #[7853]（Windows 自更新修复）仍在开放中，仅部分解决了问题。
-   **【高】Canvas-store 回归 (#7563)** **【已修复】**。

### 6. 功能请求与路线图信号

从今日的 Issue 和 PR 中，可以识别出几个明确的路线图信号：

-   **WASM 插件生命周期管理 (#7822)**：提出的 `PluginCapability::Hook` 特性，让 WASM 插件可以订阅 Agent 生命周期事件。这被视为 **v0.8.2 WASM 插件项目**的一部分，很可能被纳入该里程碑。
-   **增强的安全与供应链 CI (#7675)**：提出的 RFC 希望加固 CI 流水线，加入供应链扫描、软件来源验证和 SBOM 生成。这与 **v0.9.0 安全加固**的目标高度一致。
-   **零宕机配置热重载 (#7897)**：希望在不重启守护进程的情况下，应用安全策略和频道配置的更新。这属于 **v0.9.0** 里程碑中的一项重要工程挑战。
-   **原生上下文压缩 (#7673)**：作为 Provider 管道装饰器的 RFC，旨在降低 token 消耗。虽然可能是一个重大变更，但对降低用户成本有明显吸引力，有可能在未来版本中成为核心特性。

### 7. 用户反馈摘要

从今日的 Issue 中，可以提炼出以下用户痛点和使用场景：

-   **跨平台痛点明显**：用户尝试在 **Android Termux** 和 **Windows** 上运行和更新 ZeroClaw 都遇到了障碍。Windows 上存在 74 个测试失败和自更新机制损坏的问题，而 Android 上则是二进制安装和编译均失败。这强烈表明项目当前**对主流 PC（Linux/Mac）之外平台的支持并非首要优化目标**。
-   **功能体验细节**：
    -   **可视化反馈缺失**：用户在输入密码等敏感信息时，由于 CLI 完全隐藏输入，无法确认输入是否被系统正确接收，这是一个小而关键的体验问题。PR #7856 正在尝试解决。
    -   **回退通知不明**：用户希望当模型 Provider 内部发生智能降级（fallback）时，能收到明确的提示信息，而不仅仅是无感切换（Issue #7883）。
    -   **安装场景拓展**：#7911 提出了在 Android 上运行 ZeroClaw 的需求，这可能用于移动办公或家庭自动化场景，显示出社区对轻量级、可随处部署的 Agent 的兴趣。

### 8. 待处理积压

以下是一些值得关注的长期未响应或停滞的议题与 PR：

-   **`[needs-author-action]` 的 Issue 与 PR**：
    -   **Issue #7673**：`RFC: Native context compression as a provider pipeline decorator`，标记为 `needs-author-action`，需要作者根据反馈更新提案。
    -   **PR #7821**：`feat(config): add schema struct & risk field`，标记为 `needs-author-action`，需要作者处理后续的代码修改。
    -   **PR #7098**：`feat(channel/mattermost): add optional WebSocket listener mode`，创建于 6 月 2 日，状态为 `needs-author-action`，是社区期待的一个功能（减少延迟），已停滞超过两周，建议维护者关注或协助推动。
-   **长期跟踪的里程碑 Issue**：
    -   `v0.8.1` 追踪 (#6970)，`v0.8.2` 追踪 (#7852, #7314)，`v0.8.3` 追踪 (#7320)，以及 `v0.9.0` 追踪 (#7432) 这些大型跟踪 Issue，虽然持续有更新，但其中包含大量子任务和链接，需要维护者定期评估每个子任务的健康度，避免有重要依赖项被遗忘。

[#7841]: zeroclaw-labs/zeroclaw PR #7841
[#7840]: zeroclaw-labs/zeroclaw PR #7840
[#7684]: zeroclaw-labs/zeroclaw PR #7684
[#7842]: zeroclaw-labs/zeroclaw PR #7842
[#7673]: zeroclaw-labs/zeroclaw Issue #7673
[#6970]: zeroclaw-labs/zeroclaw Issue #6970
[#7911]: zeroclaw-labs/zeroclaw Issue #7911
[#7907]: zeroclaw-labs/zeroclaw Issue #7907
[#7563]: zeroclaw-labs/zeroclaw Issue #7563
[#7462]: zeroclaw-labs/zeroclaw Issue #7462
[#7853]: zeroclaw-labs/zeroclaw PR #7853
[#7822]: zeroclaw-labs/zeroclaw Issue #7822
[#7675]: zeroclaw-labs/zeroclaw Issue #7675
[#7897]: zeroclaw-labs/zeroclaw Issue #7897
[#7883]: zeroclaw-labs/zeroclaw Issue #7883
[#7856]: zeroclaw-labs/zeroclaw PR #7856
[#7821]: zeroclaw-labs/zeroclaw PR #7821
[#7098]: zeroclaw-labs/zeroclaw PR #7098
[#7852]: zeroclaw-labs/zeroclaw Issue #7852
[#7314]: zeroclaw-labs/zeroclaw Issue #7314
[#7320]: zeroclaw-labs/zeroclaw Issue #7320
[#7432]: zeroclaw-labs/zeroclaw Issue #7432

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*