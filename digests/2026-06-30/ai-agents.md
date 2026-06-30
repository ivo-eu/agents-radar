# OpenClaw 生态日报 2026-06-30

> Issues: 60 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-06-30 10:45 UTC

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

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，我已根据 OpenClaw (github.com/openclaw/openclaw) 的 GitHub 数据，为您生成了 2026年6月30日的项目动态日报。

---

## OpenClaw 项目动态日报 - 2026-06-30

### 1. 今日速览

今日项目社区活跃度极高。**Pull Request 数量激增**（24小时内500条），主要源于维护者 `vincentkoc` 发起的一批针对插件系统稳健性的修复工作，显示出团队正在系统性加固项目基础架构。与此同时，**严重 Bug 报告仍然集中爆发**，尤其是在消息丢失、会话状态异常和 iOS 新版本兼容性方面，P1 (最高优先级) 问题频发，提示稳定性仍是当前最大挑战。尽管没有新版本发布，但大量的修复性 PR 和社区热切的反馈，表明项目正处于一个快速迭代和修复关键问题的关键窗口期。

### 2. 版本发布

*暂无新版本发布。*

### 3. 项目进展

今日共有 **33 个 PR 被合并或关闭**，项目在修复关键 Bug 和加固基础设施方面取得了实质性进展。主要内容包括：

- **关键 Bug 修复**:
    - `#98120`: **已关闭**。修复了在 Windows 上因 `normalizeAccountId` 错误导致的网关 WebSocket 每 6-7 分钟断连一次的回归问题。
    - `#98112`: **已关闭**。修复了全量网关重启时，正在进行的子代理运行因找不到会话条目而永久丢失答案的问题。
    - `#98099`: **已关闭**。修复了 2026.6.9 升级中 SQLite 认证迁移失败，导致配置文件为空的问题。
    - `#98089`: **已关闭**。作为功能请求被合并，为内存压力监控增加了可配置阈值，提高了部署灵活性。
- **基础设施加固**: `vincentkoc` 贡献者提交了约 20 个关于 `fix(plugins)` 的 PR，这些 PR 旨在提高插件系统的鲁棒性，通过在加载时跳过格式错误的配置行或记录诊断信息，而非直接崩溃，提升了系统的稳定性。这些 PR 覆盖了工具注册、元数据解析、提供商发现等多个核心模块。
- **功能推进**:
    - `#97742`: 一个大型 PR，旨在跨提供商保留结构化的工具返回文本，修复了代理回放时可能出现占位符的问题。
    - `#98104`: 一个新的 PR，旨在记录聊天中止时的持久化失败原因，帮助诊断潜在的数据库或文件系统问题。

### 4. 社区热点

今日最热门的讨论集中在几个核心性能和稳定性问题上，用户诉求强烈：

1.  **Codex App-Server 回复截断问题 (`#84516` - 11条评论, 2个👍)**: 这是今日讨论最激烈的问题。用户报告在无头模式下，配置了 `gpt-5.5` 的回复在约 1000-1100 字符处被**静默截断**。由于模型并未报告错误（`stop=null`），这使得问题极难排查。该问题引发了开发者对底层数据流和处理逻辑的深入讨论。 [链接](https://github.com/openclaw/openclaw/issues/84516)

2.  **子代理完成消息丢失 (`#92433` - 6条评论, 1个👍)**: 此问题同样关乎用户体验。当子代理在请求者运行时完成工作，其完成消息未能正确传递给用户，导致消息丢失。这暴露了 `maybeSteerSubagentAnnounce` 函数在复杂并发场景下的逻辑缺陷。 [链接](https://github.com/openclaw/openclaw/issues/92433)

3.  **WhatsApp 会话在长时间模型调用后卡死 (`#84569` - 6条评论, 3个👍)**: 用户在 WhatsApp 上发送长耗时任务时，模型调用超过 120-240 秒后，会话状态变为 `stalled_agent_run` 并最终终止，导致回复无法送达。这是渠道集成稳定性的一大痛点。 [链接](https://github.com/openclaw/openclaw/issues/84569)

### 5. Bug 与稳定性

今日报告的 Bug 主要集中在**消息/会话状态丢失**和**新版本兼容性问题**，严重影响用户体验。

**严重 (P1):**
- **会话/消息丢失**: `#84516` (Codex 回复截断), `#92433` (子代理完成丢失), `#84569` (WhatsApp 会话卡死), `#84536` (上下文溢出静默杀死子会话), `#91456` (Telegram 车道守卫), `#98081` (Discord 重启后丢失), `#98076` (Telegram 重启时文件丢失), `#98112` (已关闭, 子代理重启丢失), `#98121` (Cron 任务锁死 Telegram 会话) 等。这些是当前稳定性危机的核心，修复 PR 大部分仍在开放或等待审核中。
- **iOS 兼容性与连接问题**: `#98062` (无法通过 Tailscale CGNAT 连接), `#98064` (局域网配对令牌丢失), `#98033` (已关闭, 二维码配对后令牌丢失), `#98116` (会话消息重复)。iOS 新版本发布带来了大量兼容性问题。
- **安全/数据风险**: `#98107` (网关重启重置凭证), `#97911` (`tools.deny` 无法隐藏技能工坊), `#97457` (网页聊天中止导致上下文丢失)。

**中等级别 (P2/P3):**
- **功能异常**: `#97826` (视频上传宽高比错误), `#98044` (安卓摄像头权限未提示), `#98101` (HTTP 429 错误误报为速率限制) 等。这些 Bug 影响了特定功能的正确性。

### 6. 功能请求与路线图信号

- **高优先级功能需求**: `#84527` (10👍) 要求在 Google Gemini CLI 即将退役前，支持其继任者 **Antigravity CLI (agy)** 作为新的 CLI 后端。这得到了社区的强力支持，且已有相关 PR (`#90101`) 在推进。
- **订阅/认证革新**: `#83954` 和 `#84528` 均呼吁支持使用用户现有 AI 模型订阅（如 ChatGPT Pro, Gemini Pro）直接进行认证，以替代按量付费的 API 模式。这表明社区对成本优化和简化配置有强烈需求。
- **核心功能增强**: `#27482` (直接视频上传到LLM), `#98092` (智能蒸馏内存), `#98069` (技能工作坊的运行时测试) 等，体现了社区对提升模型交互能力和开发体验的期待。

### 7. 用户反馈摘要

- **核心痛点**: 用户对“静默”失效模式感到非常沮丧。`#84516` 的提交者明确指出“模型没有中止，却截断回复”，这种无反馈的错误极难排查。`#84092` 的用户反映升级后长回复被“静默”丢弃，迫使他们回滚版本。这表明透明和非静默的错误处理是用户最迫切的需求。
- **配置复杂性**: 多个 Issue (如 `#83954`, `#84212`) 反映了配置过程不清晰或易出错的问题，特别是指定模型、路由和认证信息时。用户希望有更清晰的向导和文档。
- **特殊功能期望**: 除了主流渠道，用户对 Signal (`#84120`), Feishu (`#84486`) 等渠道的功能稳定性有很高期望，尤其是打字指示器这类看似微小但能极大改善体验的功能。
- **情感反馈**: Issue `#84500` 的标题“能不能精简一下代码”（意为“能不能精简一下代码”）直接表达了用户对项目代码臃肿、运作低效的不满，反映了对性能和资源消耗的普遍关注。

### 8. 待处理积压

以下长期未解决的重要 Issue 和 PR 需要维护者重点关注：

- **`#81099` (P1, 5月12日)**: **Claude CLI 后端的 `AskUserQuestion` 工具始终无响应**。这是一个影响核心交互功能的严重Bug，但自创建以来几乎没有进展。 [链接](https://github.com/openclaw/openclaw/issues/81099)
- **`#27482` (P2, 2月26日)**: **支持直接上传视频到 LLM**。这是一个具有普遍需求的增强功能，已开放超过4个月，进展缓慢。 [链接](https://github.com/openclaw/openclaw/issues/27482)
- **`vincentkoc` 的系列修复 PR (如 #90099, #90073 等, 6月4日)**: 虽然已被标记为“ready for maintainer look”，但**超过20个旨在提升系统稳健性的 PR 在过去26天里一直处于待合并状态**。这些 PR 对于防止因插件配置问题导致的崩溃至关重要，应优先处理。

---

## 横向生态对比

好的，作为资深技术分析师，我已仔细审阅了您提供的所有项目动态日报。基于这些数据，以下是为您生成的横向对比分析报告。

---

### AI 智能体与个人 AI 助手开源生态横向分析报告 (2026-06-30)

#### 1. 生态全景

当前，AI 智能体与个人 AI 助手开源生态正处于 **高速迭代与深度分化** 的关键阶段。一方面，核心项目（如 OpenClaw、NanoBot）正在围绕“稳定性”和“可扩展性”进行激烈的军备竞赛，社区 Bug 报告和修复 PR 数量激增，显示出用户对 **开箱即用的可靠体验** 有极高期待。另一方面，**多元化集成** 成为共识，多个项目（PicoClaw、CoPaw、ZeroClaw）争相集成 Telegram、WhatsApp、钉钉、Matrix 甚至去中心化协议，力求成为个人数字生活的“统一入口”。值得注意的是，对 **模型效率、Token 成本** 的质疑（如 LobsterAI）和 **零配置/简化配置** 的呼声（如 IronClaw）日益强烈，标志着市场正从“能用”向“好用、省钱”过渡。

#### 2. 各项目活跃度对比

| 项目名称 | 活跃度评级 | 新 Issues | 新/合并 PRs | 版本发布 | 健康度评估 |
| :--- | :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 极高 | 多 (密集) | 33 合并 | 无 | 积极迭代，但稳定性是核心挑战 |
| **NanoBot** | 高 | 5 / 9 关闭 | 52 提交 / 31 合并 | 无 | 功能开发与自动化推进顺利，健康度良好 |
| **Hermes Agent** | 极高 | 7 / 4 关闭 | 50 提交 / 32 合并 | 无 | 安全与稳定性修复是主线，响应迅速 |
| **PicoClaw** | 中高 | 5 | 7 提交 / 2 合并 | Nightly 版本 | 功能开发活跃，但合并效率需提升 |
| **NanoClaw** | 中 | 1 | 6 提交 / 0 合并 | 无 | 功能开发密集期，但合并流程存在瓶颈 |
| **NullClaw** | 中 | 1 | 9 提交 / 1 合并 | 无 | 核心修复与功能开发并行，进展平稳 |
| **IronClaw** | 极高 | 9 | 50 提交 / 多重合并 | 无 | 开发与测试并行，社区反馈集中 |
| **LobsterAI** | 极高 | 4 / 1 关闭 | 13 合并 | 有 (v2026.6.29) | 交付效率高，但面临性能信任危机 |
| **TinyClaw** | 低 | 0 | 0 | 无 | 无活动 |
| **Moltis** | 低 | 0 | 0 | 无 | 无活动 |
| **CoPaw** | 极高 | 3 | 50 提交 / 23 合并 | 无 | 开发节奏迅猛，多渠道集成亮点多 |
| **ZeptoClaw** | 低 | 0 | 0 | 无 | 无活动 |
| **ZeroClaw** | 极高 | 1 | 50 提交 / 9 合并 | 无 | v0.8.3 冲刺阶段，开发动力强劲，审查是瓶颈 |

#### 3. OpenClaw 在生态中的定位

- **优势与生态中心地位**：作为“核心参照”项目，OpenClaw 社区规模最大，问题反馈最密集。它是生态的 **“压力测试场”**，率先暴露了插件系统健壮性、会话状态管理、消息丢失等普适性难题。其解决方案（如 `vincentkoc` 的系列修复 PR）对全行业有指导意义。
- **技术路线差异**：与强调 **“超轻量”** 的 NanoBot 或 **“安全优先”** 的 Hermes Agent 相比，OpenClaw 更倾向于 **功能全面与平台灵活性**。它的挑战也源于此：过多的功能和配置选项导致了更高的复杂度和稳定性风险。相比之下，NanoBot 的自动化重构 (`#4382`) 和 Hermes 的安全修复 (`#55576`) 执行得更快、更聚焦。
- **社区规模与健康度**：OpenClaw 的 Issue 和 PR 绝对数量远超其他项目，但 **“积压”问题**（如 `#81099` 长期无进展）也最突出。这反映出社区虽大，但核心维护者的带宽可能成为瓶颈。而 Hermes Agent 和 CoPaw 则显示出对 Bug 更快的响应和合并速度，社区效率更高。

#### 4. 共同关注的技术方向

- **消息与会话状态一致性（全线爆发）**：**OpenClaw** (`#84516` 回复截断，`#92433` 消息丢失)、**NanoBot** (`#4595` 工具调用 ID 变异)、**Hermes Agent** (`#15093` 陈旧消息泄漏)、**PicoClaw** (聊天历史错乱) 均报告了此类问题。这是当前个人 AI 助手 **最大的稳定性痛点**，直接影响用户对“AI 是否可靠”的信任。
- **渠道集成（特别是 WhatsApp、Telegram、钉钉）**：**Hermes** (`#55584` WhatsApp 权限)、**CoPaw** (`#5590` 钉钉@提及，`#5654` 钉钉消息可靠性)、**PicoClaw** (`#3093` 去中心化协议请求)、**IronClaw** (`#55589` WhatsApp cron 任务丢失 [从数据看，应属于 Hermes 的 Issue，此处为 IronClaw 的数据，显示跨项目关注]) 都在积极完善渠道适配。**渠道稳定性和体验一致性** 是所有多平台 Agent 的必修课。
- **配置简化与零配置体验**：**IronClaw** (`#5421` Web Search 需二次认证)、**NanoBot** (`#660` 超轻量宣称与 Node.js 依赖矛盾)、**PicoClaw** (`#3195` GPT 默认配置不工作) 的反馈显示，**降低上手门槛** 是拉拢用户的关键。
- **Agent 记忆与工具调用可靠性**：**NanoBot** (`#4595` 工具调用 ID 被覆写)、**CoPaw** (`#5588` 记忆搜索需 Reranker)、**OpenClaw** (`#98112` 子代理重启丢失) 共同指向：**Agent 的长期记忆和工具调用逻辑** 需要更健壮、更智能的设计。

#### 5. 差异化定位分析

| 特性 | OpenClaw | NanoBot | Hermes Agent | PicoClaw | IronClaw | CoPaw |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| **核心定位** | 通用型 AI 智能体框架 | 超轻量、自动化优先的个人助手 | 安全优先、MoA (混合代理) 先驱 | 边缘计算、硬件集成 (NanoKVM) | 企业级、平台化 AI 助手 | 多渠道协作与知识管理型 Agent |
| **目标用户** | 高级开发者、系统集成商 | 注重效率、轻量部署的个人用户 | 安全敏感、追求模型组合效果的用户 | 嵌入式开发者、硬件爱好者 | 组织、团队、自动化运营者 | 需要强协作与记忆搜索的团队 |
| **技术架构差异** | 高度模块化、插件系统复杂 | 自动化层 (cron/脚本) 自成体系 | 强调安全审计、MoA 预设管理 | 微控制器 / 边缘设备优先 | Reborn WebUI、E2E 测试覆盖强 | 多渠道统一、记忆系统 RAG 化 |
| **社区生态** | 活跃度最高，问题最多元 | 开发节奏快，自动化方向明确 | 修复效率高，安全社区认可度高 | 新生力量，专注于特定硬件场景 | 企业文化重测试，QA 流程清晰 | 贡献者活跃，渠道集成贡献多 |

#### 6. 社区热度与成熟度

- **高活跃度与快速迭代阶段**：**OpenClaw、NanoBot、Hermes Agent、IronClaw、CoPaw、ZeroClaw**。这些项目社区用户活跃，开发者提交的 PR 数量大。其中 **Hermes Agent** 和 **CoPaw** 显示出最快的问题响应和修复能力，项目成熟度增长迅速。
- **质量巩固与功能优化阶段**：**PicoClaw** 和 **NanoClaw**、**LobsterAI**。它们的功能开发活跃，但面临 PR 合并瓶颈（NanoClaw/PicoClaw）或核心性能信任危机（LobsterAI）。LobsterAI 已发布正式版本，表明在版本化方面走在前列，但需要解决效率问题。
- **长尾/停滞阶段**：**TinyClaw、Moltis、ZeptoClaw**。过去24小时无活动，暗示这些项目可能进入维护期或社区关注度较低。

#### 7. 值得关注的趋势信号

1.  **“模型寄生”关系紧张**：用户对 AI 助手的 **Token 消耗和性能** 越来越敏感（LobsterAI `#2230`，NanoBot `#660`）。未来，AI 开发者需要将 **“模型经济学”** 纳入架构设计，如引入智能缓存、蒸馏记忆（OpenClaw `#98092`）、以及更透明的 Token 开销计量功能，否则将面临用户流失。
2.  **从“聊天机器人”向“数字员工”进化**：高频出现的 **Cron 任务**（NanoBot `#4382`，NullClaw `#783`，IronClaw `#5420`）和 **多渠道集成** 表明，用户不再满足于被动问答，而是希望 Agent 能充当 **主动执行、跨平台协调的工作流引擎**。这要求项目提供更强大的后台任务调度和状态审计能力。
3.  **安全与隐私成为硬需求**：Hermes Agent 的 ZIP 炸弹修复、API 密钥脱敏，以及 OpenClaw 的会话丢失问题，共同指向：**数据安全与可靠性是个人 AI 助手走向大规模应用的门槛**。提供沙箱执行、精细化的权限控制（如 ZeroClaw `#8496`）和透明的错误报告机制，将成为下一代 Agent 平台的标准配置。
4.  **去中心化与数据主权**：PicoClaw 社区对 **SimpleX、Tox** 等去中心化通信协议的呼声，与 DeltaChat 网关的 PR，共同昭示了 **部分用户对数据主权和隐私自建** 的强烈追求。这为专注于隐私保护的 Agent 项目提供了差异化发展的土壤。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

# NanoBot 项目每日动态日报 — 2026-06-30

## 1. 今日速览

过去 24 小时项目保持高活跃度：共处理 14 条 Issue（新开 5 条，关闭 9 条），提交 52 条 PR（合并/关闭 31 条，待合并 21 条）。无新版本发布。社区讨论集中在自动化功能完善、工具调用健壮性以及跨平台适配上。项目团队在自动化、WebUI 和 AI Agent 核心逻辑上持续推进，整体健康度良好。

## 2. 版本发布

无。

## 3. 项目进展

今日合并/关闭的重要 PR 展现了项目在多个维度的实质性推进：

- **自动化层重构**：`#4382` 将定时任务（cron）和脚本触发机制统一纳入 `automations` 层级，为后续扩展自动化类型奠定基础。同时 `#4330` 在 WebUI 中增加了自动化管理视图，支持用户创建、编辑、暂停自动化任务。
- **Agent 核心逻辑优化**：`#4359` 实现了长时间任务中目标延续上下文的延迟加载，确保后续提示能包含最新创建的目标。`#4349` 修复了回放窗口历史裁剪时可能丢弃用户长回合的问题。
- **Dream 记忆系统增强**：`#4369` 替换了空白的 `/dream` 响应为可解释提示，帮助用户理解为何无历史可学习；`#4370` 将空闲自动压缩的默认时间从 0 分钟改为 15 分钟，提升默认体验。
- **平台与通道适配**：`#4381` 修复了飞书（Feishu）流式更新失败时的重试机制；`#4545` 将 Windows 默认命令 Shell 切换为 PowerShell 并支持 `shell` 参数，解决跨盘符 `cd` 失败等问题。
- **安全与稳定性**：`#4202` 明确了文件系统写入策略，将 `extra_allowed_dirs` 设为只读并引入读写分离配置；`#4385` 在主模型失败日志中记录具体错误信息便于调试；`#4386` 抑制了 CLI 中不可路由的进度噪音。
- **持续集成**：`#4400` 跳过仅修改文档的 CI 运行，节约资源。

今日共合并/关闭 31 个 PR，项目在自动化能力、Agent 可靠性、跨平台兼容性上迈出了坚实一步。

## 4. 社区热点

- **`#660`** — **“超轻量”宣称与 Node.js 依赖的矛盾**（15 条评论，5 👍）。用户指出项目 Dockerfile 同时依赖 Python 和 Node.js，与文档中“ultra-lightweight”不符。虽然 Issue 已关闭，但引发了广泛讨论，反映出用户对项目资源占用敏感，呼吁纯 Python 化或提供可选 Node.js 方案。  
  [Issue #660](https://github.com/HKUDS/nanobot/issues/660)

- **`#4418`** — **Heartbeat 任务结果应发送到添加任务的频道**（4 条评论）。用户抱怨当前 heartbeat 将结果发送到“最近活跃的聊天频道”，而非原始添加频道，导致跨频道协调混乱。社区支持将此视为合理需求。  
  [Issue #4418](https://github.com/HKUDS/nanobot/issues/4418)

- **`#4419`** — **自动推理努力级别（reasoning effort）升降级**（4 条评论）。用户希望在 agent 配置中实现默认/升级两级努力控制，让模型在复杂任务时自动加深思考。该功能已有初步配置支持，但缺少自动化逻辑。  
  [Issue #4419](https://github.com/HKUDS/nanobot/issues/4419)

## 5. Bug 与稳定性

| 严重程度 | Bug 简述 | 状态 | 链接 |
|----------|----------|------|------|
| **严重** | `StreamingFileEditTracker.apply_final_call_ids()` 会覆盖非文件编辑工具的 `tool_call.id`，导致会话永久中毒。影响所有工具类型（如 `read_file`），被用户标记为“Session Poisoning”。 | 已报告，暂无直接 fix PR；关联 `#4603`（refactor 提议） | [Issue #4595](https://github.com/HKUDS/nanobot/issues/4595) |
| **高** | 使用 `nssm` 将 NanoBot 设为 Windows 系统服务后，`/restart` 命令导致端口占用或无限重启。Windows 用户服务管理受阻。 | 已报告，无固定 PR | [Issue #4513](https://github.com/HKUDS/nanobot/issues/4513) |
| **中** | Linux 安装脚本在 TUI 界面出现未交互即崩溃。影响新用户首次体验。 | 已关闭（可能是误报告或已修复） | [Issue #4599](https://github.com/HKUDS/nanobot/issues/4599) |
| **中** | ExecTool 路径提取正则无法识别等号后的绝对路径（如 `--output=/tmp/file`），导致 `restrictToWorkspace` 误判。 | 新开待处理 | [Issue #4592](https://github.com/HKUDS/nanobot/issues/4592) |
| **低** | 旧 Issue `#1023`：`provider login` 的 OAuth Token 未持久化，且配置刷新会丢弃未知 provider。 | 已关闭（可能未完全修复） | [Issue #1023](https://github.com/HKUDS/nanobot/issues/1023) |

## 6. 功能请求与路线图信号

- **推理努力自动升降级 (`#4419`)**：呼声较高，且配置字段 `reasoningEffort` 已存在，只需实现自动化切换逻辑。大概率进入下一版本。
- **Anthropic OAuth 支持 (`#4604`)**：用户请求从外部讨论迁移而来，涉及企业级认证，若采纳将扩展提供商生态。
- **外部脚本触发 agent 动作 (`#4605`)**：用户期望通过 HTTP 或 CLI 方式调用 Agent 执行特定任务，这符合“个人 AI 助手”的集成需求，与自动化方向一致。
- **`tool_call.id` 不再被 WebUI 进度跟踪变异 (`#4603`)**：针对 `#4595` 的根源修复提案，属于必要重构，预计将被采纳。
- **MCP 凭证脱敏 (`#4584`)**：安全改进，防日志泄露 token，开放 PR 待合并。
- **WeChat 通道流式支持 (`#4567`)**：解决 Anthropic 兼容代理在非流式下丢弃 `tool_use` 字段的问题，推进中。
- **Dream 提示覆盖 (`#4491`)**：允许 workspace 级别的 Dream 指令，开放 PR 待合并，提升记忆定制能力。

## 7. 用户反馈摘要

- **轻量级质疑**：`#660` 中用户指出“ultra-lightweight”名不副实，Docker 镜像需同时安装 Python 和 Node.js。项目团队虽已关闭 Issue，但未给出明确解释或改进计划，建议在文档中澄清或提供最小化构建。
- **Windows 服务部署痛点**：`#4513` 中文用户详细描述了 `nssm` 服务化后的重启异常，称“程序不断尝试重启”或“服务停止但实际运行”。Windows 用户群体对系统服务稳定性较差。
- **Heartbeat 工作流混乱**：`#4418` 用户指出 heartbeat 将结果定向到最近频道而非原始频道，导致“无法追踪任务来源”。该功能在多人通道场景下问题突出。
- **工具调用 ID 变异导致不可预知行为**：`#4595` 用户详细描述了 `apply_final_call_ids` 如何错误覆盖所有工具的调用 ID，称其为“Session Poisoning”，呼吁立即修复。
- **安装脚本友好性**：`#4599` 用户通过 bot 报告 Linux 安装脚本崩溃，虽可能是测试反馈，但说明安装流程仍存在未捕获的异常。

## 8. 待处理积压

- **`#4419` — 推理努力自动升降级**（2026-06-20 创建，4 评论，0 assignee）：功能请求已获社区初步认可，但未分配到任何 milestone 或特定维护者。建议在下一版本规划中优先评估。
- **`#4513` — Windows nssm 服务重启异常**（2026-06-25 创建，2 评论）：属于平台关键 Bug，但至今无 PR 响应。Windows 用户数量虽少于 Linux，但此问题影响服务化部署体验，应标记为 high priority。
- **`#1023` — Provider Token 持久化与配置刷新问题**（2026-02-22 创建，已关闭但讨论未结）：用户指出 OAuth token 不持久化且刷新会丢失部分 provider 配置。虽已关闭，但未看到具体修复 PR，建议重新验证是否彻底解决。

> 以上为 2026-06-30 动态日报。所有数据来源于 NanoBot GitHub 仓库（[hkuDS/nanobot](https://github.com/HKUDS/nanobot)），统计周期为过去 24 小时。

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

# Hermes Agent 项目动态日报

**日期：2026-06-30**  
**分析师：AI 开源项目分析师**  
**数据来源：GitHub (NousResearch/hermes-agent)**

---

## 今日速览

- 过去24小时内项目活跃度极高：**50条PR更新**（其中32条已合并/关闭，18条待合并），**11条Issue更新**（7条新开/活跃，4条已关闭），无新版本发布。
- 安全与稳定性修复是今日主线：至少5条安全相关PR被合并（API密钥日志脱敏、WhatsApp转发权限闸门、ZIP炸弹防护等），多项网关和代理核心bug得到根治。
- 社区反馈集中在**MoA（混合代理）预设行为异常**（3个相关bug已全部关闭）、**WhatsApp集成缺失**以及**会话状态丢失**等影响日常使用的痛点。
- 维护者反应迅速：多个当天提交的严重bug（如#55572、#55589）已获得fix PR或处于待合并状态，项目健康度处于良好水平。

---

## 版本发布

无新版本发布。

---

## 项目进展

今日共合并/关闭 **32条PR**，以下为关键推进项：

### 🛡️ 安全加固（5条合并）
- **#55576** — [`fix: redact status --all API keys, harden compression message handling`](https://github.com/NousResearch/hermes-agent/pull/55576)  
  `hermes status --all` 现在会遮掩API密钥（符合帮助文本“redacted for sharing”的承诺），且上下文压缩处理更健壮。
- **#55583** — [`fix(security): redact Fireworks AI API keys in logs`](https://github.com/NousResearch/hermes-agent/pull/55583)  
  新增对Fireworks AI密钥格式的日志脱敏，补全了此前遗漏的注册提供商。
- **#55584** — [`fix(whatsapp): gate opt-in owner-typed forwards on customer chatId allowlist`](https://github.com/NousResearch/hermes-agent/pull/55584)  
  限制WhatsApp Bot模式下“转发所有者消息”功能仅对白名单用户生效，防止意外泄露。
- **#29450** — [`fix(skills-hub): stream ZIP extraction to prevent zip bomb via manipulated metadata`](https://github.com/NousResearch/hermes-agent/pull/29450)  
  修复ZIP炸弹漏洞：不再依赖ZIP中央目录元数据（可被篡改），采用流式解压并监控实际数据量。
- **#27500** — 与#55583等效的旧PR最终关闭（被新PR取代），确认安全性修复已到位。

### 🐛 关键Bug修复
- **#55582 / #27426** — [`fix(tools): defense-in-depth guard against hallucinated acp_command crashing the gateway`](https://github.com/NousResearch/hermes-agent/pull/55582)  
  防止模型在`delegate_task`中虚构`acp_command`（如`"copilot"`）导致缺少对应CLI的网关崩溃。新增两层防御：schema校验+运行时降级。
- **#55559 / #15093** — [`fix(gateway): add session staleness guard to stream consumer`](https://github.com/NousResearch/hermes-agent/pull/55559)  
  `GatewayStreamConsumer`新增会话陈旧性检查，用户执行`/new`或`/stop`后立即丢弃残余delta，避免陈旧片段与“会话重置”确认消息一起发送。
- **#55585** — [`fix(kanban): unknown skill warns instead of crashing the worker`](https://github.com/NousResearch/hermes-agent/pull/55585)  
  Kanban工作进程遇到不存在的技能名不再崩溃，而是跳过并记录警告。对应关闭Issue #27136。
- **#54021** — [`fix(agent): resolve MoA preset to its aggregator for auxiliary tasks`](https://github.com/NousResearch/hermes-agent/pull/54021)  
  修复当主模型为MoA预设时，辅助任务（标题生成、上下文压缩、视觉）因模型ID无效而失败的bug。
- **#54384** — [`fix(moa): propagate api_mode from slot runtime to call_llm`](https://github.com/NousResearch/hermes-agent/pull/54384)  
  MoA参考槽位的`api_mode`传递丢失导致Copilot GPT-5.x模型被错误路由到`/chat/completions`引发400错误，现已修复。
- **#27123** — [`fix(gemini): add missing role field to systemInstruction for Gemma compatibility`](https://github.com/NousResearch/hermes-agent/pull/27123)  
  解决Gemini原生适配器缺少`"role": "system"`字段导致Gemma模型返回500错误的问题。

### ✨ 新功能/增强
- **#54349** — [`feat(cron): audit log for missed jobs to surface scheduler gaps`](https://github.com/NousResearch/hermes-agent/pull/54349)（待合并）  
  新增cron作业调度审计日志，记录丢失的调度时刻，方便运维排查。
- **#55581** — [`fix(desktop): guard voice loop during playback`](https://github.com/NousResearch/hermes-agent/pull/55581)（待合并）  
  防止桌面端TTS播放过程中误触发新麦序；改进语音循环错误恢复。

### 📌 项目整体进展评估
- 今日合并的修复覆盖了**6个严重程度P1/P2的bug**，且均为用户反馈的真实崩溃或数据丢失问题。
- 安全性方面，**3个不同攻击面**（API密钥泄露、ZIP炸弹、消息转发越权）得到修补，防御纵深加强。
- 待合并的18个PR中包含多个功能请求（如cron审计日志、多模型配置支持），预示着下一版本的功能丰富度。

---

## 社区热点

### 🔥 最高讨论量 Issue
单条Issue评论数虽然不高（最多3条），但以下两项引发的关联PR数量最多，反映社区关切：

1. **[#54379] MoA参考槽位api_mode丢失** — 1条评论，但触发PR #54384立即修复并关闭。  
   [Issue链接](https://github.com/NousResearch/hermes-agent/issues/54379) | [PR链接](https://github.com/NousResearch/hermes-agent/pull/54384)  
   **诉求**: 使用Copilot GPT-5.x模型的MoA预设用户遭遇400错误，影响生产使用。

2. **[#52445] GUI cron表单缺少WhatsApp选项** — 3条评论（最高）。  
   [Issue链接](https://github.com/NousResearch/hermes-agent/issues/52445)  
   **诉求**: 尽管已配置WhatsApp且网关运行，GUI新建cron任务的下拉菜单仍硬编码缺少WhatsApp。用户期望统一的配置体验。该项目标注为`duplicate`，但尚未关闭。

### 🔥 关注度最高的PR
- **[#55576] `status --all` API密钥脱敏** — 虽无评论，但涉及安全+日常运维，影响力大。  
  链接：[PR #55576](https://github.com/NousResearch/hermes-agent/pull/55576)

---

## Bug 与稳定性（按严重程度排列）

| 严重程度 | Issue # | 描述 | 状态 | 是否有 fix PR |
|---------|---------|------|------|---------------|
| **P1** | #27426 / #55582 | 模型虚构 `acp_command` 导致网关崩溃 | 已合并 | ✅ #55582 |
| **P1** | #15093 / #55559 | 网关流消费者未检查会话陈旧性，导致陈旧消息泄漏 | 已合并 | ✅ #55559 |
| **P1** | #29450 | ZIP炸弹漏洞（攻击者可构造恶意ZIP） | 已合并 | ✅ #29450 |
| **P2** | #55572 | 尾部保护token估计未计入 `codex_reasoning_items`，导致压缩时机延迟 | 新开（当日） | ❌ 无PR |
| **P2** | #55589 | WhatsApp回复cron投递丢失任务上下文 | 新开（当日） | ❌ 无PR |
| **P2** | #55588 | 压缩链会话在根会话归档后从侧栏消失 | 新开（当日） | ❌ 无PR |
| **P2** | #54669 | `moa.enabled: true` 静默污染 `model.base_url` 为 `moa://local` | 已关闭 | ✅ #54021 等配套修复 |
| **P2** | #54259 | MoA网关主/行动调用解析为 `moa://local` -> 404 -> fallback | 已关闭 | ✅ #54021 |
| **P3** | #55594 | TUI长响应滚动出界后不可达 | 新开（当日） | ❌ 无PR |
| **P3** | #55578 | 桌面异步委托完成可复活旧会话，后续提示却创建新会话 | 新开（当日） | ❌ 无PR |
| **P3** | #52445 | GUI cron表单缺少WhatsApp选项（duplicate） | 开放 | ❌ 无PR |

**小结**：7个新bug中4个为P2等级，其中#55572和#55589直接影响会话压缩和WhatsApp任务连续性，需优先处理。已有P1/P2 bug当日全部通过PR解决，维护响应效率高。

---

## 功能请求与路线图信号

| Issue/PR # | 功能描述 | 状态 | 可能的纳入版本 |
|-----------|---------|------|----------------|
| [#55573](https://github.com/NousResearch/hermes-agent/issues/55573) | 支持为自定义OpenAI兼容提供商配置多个模型 | 新开 | 下一版本（0.18.0）或后续 |
| [#54349](https://github.com/NousResearch/hermes-agent/pull/54349) | Cron作业审计日志（调度缺口感知） | 待合并 | 0.18.0（若获批） |
| [#55581](https://github.com/NousResearch/hermes-agent/pull/55581) | 桌面语音循环保护 | 待合并 | 0.18.0 |
| [#55586](https://github.com/NousResearch/hermes-agent/pull/55586) | Anthropic凭证池同步access_token变化 | 开放 | 0.18.0 |
| [#55587](https://github.com/NousResearch/hermes-agent/pull/55587) | API-key提供商增加config.yaml中的`model.api_key`备用读取 | 开放 | 0.18.0 |

**趋势判断**：
- 用户对**多模型支持**的呼声持续，尤其本地/自定义提供商场景。
- **运维可见性**（cron审计日志）和**会话状态一致性**（语音、压缩链）是社区明确期望增强的方向。

---

## 用户反馈摘要

### 典型痛点（来自Issue评论）
1. **WHATSAPP集成体验割裂**  
   “我已正确配置WHATSAPP_ENABLED=true，但GUI新建cron任务下拉菜单仍然没有WhatsApp选项。”  
   → 反映UI未动态读取运行时配置，硬编码导致配置与实际能力不一致。

2. **MoA预设破坏辅助功能**  
   “使用MoA预设后，标题生成永远失败，显示`⚠ Auxiliary title generation failed`。”  
   → 对日常多轮对话影响大，用户不得不回退到单模型。

3. **桌面端长回答“消失”**  
   “这是日常使用中排名第一的UX挫折。长回答滚出可视区域后无法恢复。”  
   → 虚拟滚动与粘性定位交互的bug，严重影响阅读体验。

### 满意场景
- 安全社区对API密钥脱敏和ZIP炸弹修复表示认可（PR #55576、#29450快速合并）。
- 多名用户在Issue评论区感谢维护者快速响应MoA相关问题（#54021、#54384）。

---

## 待处理积压

以下为长期未响应或阻塞时间较长的重要事项，需维护者关注：

| 类型 | 编号 | 描述 | 创建时间 | 最后更新 | 建议动作 |
|------|------|------|---------|---------|---------|
| Issue | [#52445](https://github.com/NousResearch/hermes-agent/issues/52445) | GUI cron表单缺失WhatsApp选项（标记为duplicate但未关闭） | 2026-06-25 | 2026-06-30 | 确认是否被其他PR修复，或与#55589合并处理 |
| PR | [#54349](https://github.com/NousResearch/hermes-agent/pull/54349) | Cron审计日志功能（已2天无review） | 2026-06-28 | 2026-06-30 | 需reviewer决策是否合并 |
| PR | [#55586](https://github.com/NousResearch/hermes-agent/pull/55586) | Anthropic凭证池同步access_token变更 | 2026-06-30 | 2026-06-30 | 新提交，需评审 |
| PR | [#55587](https://github.com/NousResearch/hermes-agent/pull/55587) | API-key提供商增加config.yaml fallback | 2026-06-30 | 2026-06-30 | 新提交，需评审 |

此外，**18条待合并PR**中可能包含其他未被高亮的功能修复，建议维护者对队列进行优先级排序，优先合入P1安全修复及代码冻结窗口前的功能。

---

**备注**：本日报基于2026-06-29至2026-06-30的GitHub活动数据生成，部分结论可能随后续PR合并或新Issue开设而变化。

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

好的，以下是为您生成的 PicoClaw 项目动态日报。

---

# PicoClaw 项目动态日报 | 2026-06-30

## 1. 今日速览

PicoClaw 项目今日保持高活跃度。社区在 24 小时内提交了 5 个新 Issue 和 7 个 PR，显示用户积极参与问题报告与功能建议。一个针对 SSRF 漏洞的修复和一项提升开发者体验的认证错误处理 PR 已被合并，对项目稳定性和可用性有直接贡献。同时，新发布的 Nightly 版本包含了多项进入 `main` 分支的改进，项目整体处于快速迭代期。

## 2. 版本发布

*   **Nightly Build (v0.3.1-nightly)**
    *   **链接**: [v0.3.1-nightly.20260630.52320f48](https://github.com/sipeed/picoclaw/releases/tag/v0.3.1-nightly.20260630.52320f48)
    *   **说明**: 这是一个自动化构建的每日构建版本，可能包含不稳定因素。
    *   **主要变更**: 该版本包含了合并到 `main` 分支的最新修改，具体变更可参见 [完整更新日志](https://github.com/sipeed/picoclaw/compare/v0.3.1...main)。

## 3. 项目进展

今日有两个重要的 Pull Request 被合并关闭，标志着项目在安全性和易用性上迈出一步：

*   **修复 SSRF 绕过漏洞**: PR [#3143](https://github.com/sipeed/picoclaw/pull/3143) 修复了 `web_fetch` 功能中的一个 SSRF 防护绕过问题，该项漏洞由 Issue [#3074](https://github.com/sipeed/picoclaw/issues/3074) 报告。该项目通过增强 IP 分类器，使其能够识别内嵌私有 IPv4 地址的 ISATAP IPv6 字面量，从而防止攻击者利用此特性绕过安全限制。
*   **优化认证错误提示**: PR [#3198](https://github.com/sipeed/picoclaw/pull/3198) 被合并，它改进了提供商的认证错误处理逻辑。当用户配置的 API 密钥或令牌失效时，系统现在会返回更清晰、更友好的错误提示，帮助用户快速定位和解决问题，显著提升了配置环节的用户体验。

## 4. 社区热点

今日社区讨论热度最高的议题集中在 **新网关功能请求** 和 **对特定平台兼容性的反馈** 上。

*   **最活跃议题**: [Feature: I need SimpleX or tox (#3093)](https://github.com/sipeed/picoclaw/issues/3093)
    *   **详情**: 该议题共收到 4 条评论和 1 个赞。用户 `Damian-o2` 明确提出希望 PicoClaw 能支持 `SimpleX` 和 `Tox` 这两种去中心化通信协议作为消息网关。这表明社区中部分用户对去中心化、隐私保护的通信方式有强烈需求，这类需求在过去几周也伴随相关 PR（如 [#3063 DeltaChat 网关）持续出现。
*   **值得关注的合并**: PR [#3063](https://github.com/sipeed/picoclaw/pull/3063)（feat: add deltachat gateway） 仍处于开放状态，自 6 月 8 日提出以来持续获得更新，反映了社区对引入新消息网关的持续关注和努力。

## 5. Bug 与稳定性

今日报告了多个 Bug，严重程度不一。值得注意的有：

*   **功能失效 (高)**:
    *   [#3195](https://github.com/sipeed/picoclaw/issues/3195): OpenAI GPT 在 NanoKVM 上默认配置无法工作。用户 `rtadams89` 报告在 NanoKVM 2.4.0 环境下配置 GPT-5.4 后，所有交互均无响应。这可能是 NanoKVM 集成或特定模型兼容性问题。
*   **登录问题 (高)**:
    *   [#3196](https://github.com/sipeed/picoclaw/issues/3196) 与 [#3197](https://github.com/sipeed/picoclaw/issues/3197): 用户 `nyawitniorang` 在同一时段提交了两个几乎相同的 Issue，报告 `Codex` 和 `Antygravity` 的 OAuth 登录在 v0.2.9 上失效。这暗示可能存在与 OAuth 流程相关的回归问题或配置变更。
*   **数据泄露 (中)**:
    *   [#3153](https://github.com/sipeed/picoclaw/issues/3153): 当使用 `Volcengine` 的 `doubao-seed-2.0-pro` 模型时，工具调用偶尔会以原始文本 `<seed:tool_call>` 返回，而非执行。此 Bug 会直接导致 AI 功能异常，影响用户体验，目前仍在开放中。

## 6. 功能请求与路线图信号

用户需求呈现多元化趋势，尤其在消息网关和模型集成方面。

*   **消息网关**: 除了 #3093 中提到的 `SimpleX` 和 `Tox`，PR [#3063](https://github.com/sipeed/picoclaw/pull/3063) 正在尝试增加 `DeltaChat` 作为新网关。这表明 PicoClaw 正在向支持更广泛、更注重隐私的通信协议方向演进。
*   **模型增强**: PR [#3163](https://github.com/sipeed/picoclaw/pull/3163)（AWS Bedrock 提示缓存） 和 PR [#3156](https://github.com/sipeed/picoclaw/pull/3156)（逐轮 Token 用量统计） 都提出了非常具体且实用的模型层优化。前者能降低 AWS Bedrock 的推理成本，后者能为开发者提供更精细的消费监控，这两项功能很可能会被纳入下一个稳定版本。
*   **远程 Agent 模式**: PR [#3118](https://github.com/sipeed/picoclaw/pull/3118) 提出的“远程 Pico WebSocket 模式”正在开发中，它为 `picoclaw agent` 命令增加了远程连接功能，这将极大地扩展 PicoClaw Agent 的使用场景，例如在边缘设备上部署 agent。

## 7. 用户反馈摘要

从今日的 Issue 评论中，我们可以提炼出以下用户痛点：

*   **配置与兼容性痛点**: 用户 `rtadams89` 和 `nyawitniorang` 遇到的配置问题（OpenAI 无法工作、OAuth 登录失败）反映了新版本（v0.2.9）或新环境（NanoKVM）中存在未完全解决的兼容性问题，导致用户上手和使用受阻。
*   **特定场景下的功能缺陷**: 用户 `ms8great` 报告的工具调用文本泄露问题 (`#3153`)，直接影响了 AI 功能的可靠性和准确性，属于影响核心体验的严重问题。
*   **用户界面兼容性问题**: 已关闭的 Issue [#3090](https://github.com/sipeed/picoclaw/issues/3090) 说明，部分用户仍在用较老版本的 iOS Safari 浏览器，这提醒项目组在 UI 更新时需注意对旧版浏览器的兼容性支持。

## 8. 待处理积压

*   **重要的未解决问题**: Issue [#3153](https://github.com/sipeed/picoclaw/issues/3153)（Volcengine 工具调用泄漏）已经报告一周多，仍然为开放状态，且影响核心功能，建议优先处理。
*   **长期开放的 PR**: PR [#3063](https://github.com/sipeed/picoclaw/pull/3063)（DeltaChat 网关）和 PR [#3118](https://github.com/sipeed/picoclaw/pull/3118)（远程 Agent 模式）均是引入重要新功能的 PR，但它们均已在开放状态超过一周。考虑到社区对这些功能有较高呼声，建议维护者评估进度，给出明确的合并或修改方向，以保持社区参与度。

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

好的，遵照您的指示，以下是基于所提供数据生成的 NanoClaw 项目动态日报。

---

### NanoClaw 项目动态日报 | 2026-06-30

---

#### 1. 今日速览

今日项目活跃度中等偏上，**未发布新版本**，但 Pull Request (PR) 提交活跃，共有 **6** 个新 PR，表明开发团队正在积极推进多项新功能与修复。与此相对，过去24小时内**没有PR被合并**，合并流程可能存在瓶颈。Issues 方面，社区报告了1个关于 Discord 频道附件处理的功能性 Bug，已收到开发者初步分析与回应。整体来看，项目处于**功能密集开发与稳定并重**的阶段。

#### 2. 版本发布

*无*

---

#### 3. 项目进展

过去24小时内**无 PR 被合并或关闭**，但新提交的 PR 展现了项目在以下几个方向的重要突破：

- **渠道适配器扩展**：PR #2884 ([链接](nanocoai/nanoclaw PR #2884)) 正式添加了 **Discord 渠道适配器**，并通过 Chat SDK 桥接实现网关模式连接，填补了 Discord 原生集成的空白。修复了 Discord DM 审批按钮路由问题，也是该 PR 的重要贡献。
- **模板系统初步实现**：PR #2890 ([链接](nanocoai/nanoclaw PR #2890)) 引入了 **Agent 模板系统**，允许用户从可复用的指令、MCP 工具和服务包中快速创建智能体组，极大提升了上手效率和配置复用性。
- **功能型 Agent 开发**：PR #2889 ([链接](nanocoai/nanoclaw PR #2889)) 贡献了一个 **Daily-News-Agent 示例**，整合了 HN、RSS 抓取、LLM 摘要与定时任务，并新增了**微信渠道适配器**，展示了项目在具体场景下的应用能力。
- **安全与稳定性修复**：PR #2880 ([链接](nanocoai/nanoclaw PR #2880)) 正在进行一项重要的安全修复，旨在解决因容器内符号链接逃逸导致的任意文件写入漏洞（CWE-59）。此修复覆盖了所有文件写入路径，对安全性有重大意义。

---

#### 4. 社区热点

- **热度最高的 Issue: #2888 - Discord 附件丢失问题** ([链接](nanocoai/nanoclaw Issue #2888))
  - **诉求分析**：用户直观地报告了 Discord 频道（以及其他基于 URL 的聊天 SDK适配器）在接收图片、文件等附件时，Agent 仅能看到文件名和元数据，而非实际内容。这与 Telegram 的完美体验形成对比，直接影响了产品的核心可用性。用户社区对此类功能缺陷高度敏感，因为它破坏了用户“发送即分享”的预期。该 Issue 已有开发者在摘要中定位到问题根因，引发了对渠道适配器一致性的讨论。

---

#### 5. Bug 与稳定性

按严重程度排列：

1.  **严重 - 安全漏洞 (CWE-59)**：PR #2880 ([链接](nanocoai/nanoclaw PR #2880)) 正修复一个**符号链接逃逸漏洞**。攻击者在容器内的会话目录中预置符号链接，可导致宿主机器文件被任意写入。此问题影响面广，涉及所有文件写入路径，已有修复方案待合并。
2.  **高 - 功能缺失/Discord 附件**：Issue #2888 ([链接](nanocoai/nanoclaw Issue #2888)) 报告了**Discord 频道Agent无法解析图片/文件附件内容**，仅提取元数据。这严重影响用户通过 Discord 使用 NanoClaw 的体验，已被定位为代码层问题 (`messageToInbound` 方法)。
3.  **中 - 配置继承 Bug**：PR #2886 ([链接](nanocoai/nanoclaw PR #2886)) 修复了**新建 Agent 不继承安装时的默认 LLM 提供者**的问题。在单提供者环境下，这会导致 Agent 因使用错误的默认值而出现 401 认证失败，阻碍用户正常使用。
4.  **低 - 设置流程不一致**：PR #2885 ([链接](nanocoai/nanoclaw PR #2885)) 指出**Slack 的 Socket Mode 集成功能**已开发完成，但因其合并到非主线分支，导致主分支上的引导设置流程仍只支持 webhook 模式，功能上线不完整。

---

#### 6. 功能请求与路线图信号

- **多渠道原生集成**：PR #2884 (Discord) 和 #2889 (WeChat) 的提交，以及 Issue #2888 对 Discord 附件问题的反馈，强烈表明 **完善和统一多通道适配器体验** 是当前社区的迫切需求和项目的优先发展方向。
- **模板化与快速部署**：PR #2890 引入的 Agent 模板系统暗示项目正在向 **“平台化”** 演进，允许用户和社区分享、复用最佳实践，这将是提升开发者生态活力的重要组件。
- **场景化 Agent 示例**：PR #2889 提交的 Daily-News-Agent 是一个完整的场景化应用。这表明用户社区和贡献者希望看到更具体的 **“开箱即用”** 示例，以降低学习成本和推广项目价值。

---

#### 7. 用户反馈摘要

从 Issue #2888 的评论中可提炼出真实用户痛点：
- **核心体验不一致**：用户能发送文件到 Telegram 并得到 Agent 的有效处理，但在 Discord 中同样的操作失败了。这种跨平台的体验不一致是用户不满的主要来源，表明适配器的完整性和质量需要统一标准。
- **对透明度和诊断的期待**：用户报告问题时，开发者能快速定位到代码层面 (`messageToInbound` 方法) 的原因，这虽然展示了开发团队的响应速度，但也隐含了用户对当前系统**缺乏清晰的错误信息或调试日志**的不满，迫使他们依赖经验或底层代码分析。

---

#### 8. 待处理积压

- **待合并的关键修复**：
  - **PR #2880** ([链接](nanocoai/nanoclaw PR #2880))：安全漏洞修复。此PR已打开超过48小时，考虑到其涉及“严重”级别的安全漏洞，建议维护者尽快跟进审查与合并，以避免潜在风险。
  - **PR #2886** ([链接](nanocoai/nanoclaw PR #2886))：配置继承修复。此修复影响单提供者环境下的新用户开箱体验，建议在下次发布前完成合并。
- **待解决的问题**：
  - **Issue #2888** ([链接](nanocoai/nanoclaw Issue #2888))：关于Discord附件处理的严重 Bug。虽然已有初步分析，但需要尽快分配开发资源进行修复，并同步到所有依赖 Chat-SDK 的适配器，防止跨平台问题扩散。

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

好的，遵照您的指示。以下是根据 NullClaw 项目数据生成的2026年6月30日项目动态日报。

---

# NullClaw 项目动态日报 | 2026-06-30

## 1. 今日速览

今日项目整体活跃度处于**中等偏上**水平。主要动态集中在Pull Request（PR）的推进与修复上，其中一项重要的CLI交互修复被合并，同时两项新的功能特性（流式工具调用和Cron子代理）仍在积极迭代中。Bug方面，有用户报告了一个关于Telegram频道空闲后停止响应的关键问题，但目前尚未有对应的修复PR。项目整体处于 **“功能开发与关键问题修复并行”** 的状态，长期未合并的PR（如#783）持续更新，表明核心开发工作仍在稳步进行。

## 2. 版本发布

无新版本发布。

## 3. 项目进展

今日项目向前迈进了重要一步，主要是对一个优先级较高的交互体验问题进行了修复。

- **【已合并】修复：CLI Agent REPL中的键盘方向键支持 (#960)**
  - **作者**: vernonstinebaker
  - **状态**: 已合并/关闭
  - **详情**: 此PR为交互式 `nullclaw agent` REPL添加了一个轻量级的行编辑器。它通过启用POSIX原始模式，使得用户按键（如方向键、Home/End、删除键等）能正常工作，而不是被打印为乱码。这显著改善了用户在终端中进行交互式对话的体验。
  - **链接**: [nullclaw/nullclaw PR #960](https://github.com/nullclaw/nullclaw/pull/960)

此外，今日有多项开放中的PR仍在活跃推进，预示着未来的功能更新：

- **【新增功能】支持SSE流式传输中的原生工具调用 (#971)**: 此PR解耦了原生工具调用与流式传输路径，允许在流式响应期间实现原生工具调用，避免了之前必须将工具强制注入prompt的问题，有望显著提升部分模型提供商的流式响应质量。
- **【新增功能】Cron子代理引擎（长期开发中）(#783)**: 此PR自4月份起开发，今日仍有更新。它致力于打造一个完整的后台任务调度系统，包括数据库支持、多种任务类型、JSON输出以及安全增强。这是项目向平台化演进的重要模块。
- **【修复】CLI Agent REPL中的键盘方向键支持 (#970)**: 这是对#960合并后同名但未合并PR的补充，可能包含了额外的优化或调整。

## 4. 社区热点

今日社区讨论热点集中在用户反馈的一个稳定性问题上。

- **Issue #972：Telegram频道在空闲一段时间后停止响应（1条讨论）**
  - **状态**: 开启
  - **分析**: 这是今天唯一的新增Issue，且用户反馈了较严重的问题。用户描述Telegram频道在正常工作一天后，第二天早上便“宕掉”，而后台的NullClaw进程看似仍在运行（能正常响应 `ping`）。社区的核心诉求是**排查后台进程与Telegram适配器之间的连接保活或资源泄漏问题**。虽然目前评论数为0，但这类偶发性、影响用户实际使用的问题往往能引发广泛关注。
  - **链接**: [nullclaw/nullclaw Issue #972](https://github.com/nullclaw/nullclaw/issues/972)

## 5. Bug 与稳定性

今日仅报告了一个严重级别的稳定性问题。

- **严重：Telegram频道空闲后停止响应 (Issue #972)**
  - **描述**: 用户报告Telegram机器人信道在闲置一夜后停止响应，但后端进程本身运行正常。这暗示问题可能出在与Telegram API的WebSocket或长轮询连接上，或存在某种后台资源未正确释放。目前**无关联的Fix PR**。
  - **链接**: [nullclaw/nullclaw Issue #972](https://github.com/nullclaw/nullclaw/issues/972)

## 6. 功能请求与路线图信号

根据目前开放的PR，以下功能被明确提上日程，可能进入下一版本：

1.  **Cron子代理引擎 (#783)**: 强大的后台任务调度系统。
2.  **SSE流式传输中的原生工具调用 (#971)**: 提升流式API的灵活性和效率。
3.  **CLI Agent REPL交互增强 (#970)**: 对刚刚合并的#960进行补充或优化。

此外，Issue #972报告的“长时间空闲断开”问题，可能暗示社区对 **连接保活与稳定性** 有更强的潜在需求。

## 7. 用户反馈摘要

从今日有限的反馈中，可以提炼出以下痛点：

- **稳定性痛点**: 用户（i11010520）反馈了Telegram频道在长时间空闲后自动“宕掉”的问题，这是一个非常影响实际部署体验的稳定性问题。用户花费了时间部署，却发现系统无法长期稳定运行。后台“看似正常”的现象增加了排查难度，让用户感到困惑。

## 8. 待处理积压

今日没有长期无人响应的新Issue或PR。但可以注意到，**PR #783 (Cron子代理)** 从4月7日创建至今已近3个月，尽管仍在更新，但始终未合并。如此大规模的变更如果持续未合并，可能面临与主干代码冲突的风险。维护团队可考虑评估其成熟度，决定是否加快审查或拆分合并。

- **链接**: [nullclaw/nullclaw PR #783](https://github.com/nullclaw/nullclaw/pull/783)

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，根据您提供的 IronClaw 项目 GitHub 数据，现为您呈上 **2026-06-30** 的项目动态日报。

---

### IronClaw 项目动态日报 — 2026-06-30

#### 1. 今日速览

今日 IronClaw 项目状态**非常活跃**，开发与测试工作并行推进。过去 24 小时内，项目共产生了 9 条新 Issues 和高达 50 条 PR，显示出极高的社区和开发团队参与度。尽管没有新版本发布，但多个涉及核心功能修复与 QA 改进的 PR 已被合并，表明项目正在积极解决稳定性问题并提升测试覆盖率。值得注意的是，Issues 中集中反馈了 `Routine` (自动化) 功能在交付渠道配置上的全局默认行为问题，以及 `Web Search` 功能的零配置化问题，这是当前社区关注的热点。整体而言，项目处于快速迭代和问题修复并重的阶段。

#### 2. 版本发布

*(今日无新版本发布)*

#### 3. 项目进展

今日项目有多项实质性进展，主要体现在 PR 的合并与关闭上，推进了以下几个方向：
- **WebUI 功能完整性与测试覆盖**：核心贡献者 `ilblackdragon` 系列 [codex] PR（#5373、#5374、#5375、#5376）被合并或保持开放，将频道配对、扩展管理、项目设置等遗留浏览器覆盖测试迁移至 Reborn WebUI，增强了新 UI 的稳定性和功能完备性。
- **QA 与测试基础设施优化**：贡献者 `serrrfirat` 的 PR #5436 被合并，增强了对 Live QA 回复成功变体的容错性。PR #5424 也已被合并，将失败的 QA 用例链接到具体的构建工件 (Artifacts)，极大提升了问题排查效率。
- **核心功能修复**：PR #5395 被合并，修复了 Web Access 功能中 `Exa` 内容获取的缓存与直接查询问题。PR #5438 保持开放，计划修复 GitHub 搜索验证恢复和输入时区偏移等问题。
- **通知系统初步引入**：贡献者 `italic-jinxin` 提交了 PR #5441，为 WebUI 增加了自动化任务运行时的**头部通知 (Header Notifications)** 功能，旨在解决用户难以察觉自动化结果的问题。

#### 4. 社区热点

本周最受关注的议题主要围绕**配置体验与用户可见性**。

1.  **日常投递的全局默认路由 Bug (#5420)**：
    - **链接**: [Issue #5420](https://github.com/nearai/ironclaw/issues/5420)
    - **分析**: 该 Issue 由 `thisisjoshford` 提出，批评了 `Routine` (自动化) 的交付目标被设计成一个**全局用户默认值**，而非按 `Routine` 独立设置。这导致用户为某个 Routine 设置 Slack 通道后，所有其他 Routine 的交付目标也会被篡改。这触犯了用户对“配置隔离”的基本预期，对运营自动化工作流的用户来说是一个**严重的可用性倒退**，得到了社区的强烈共鸣（尽管 👍 数为 0，但话题重要性极高）。

2.  **自动化任务结果的可发现性 (#5443 / #5441)**：
    - **链接**: [Issue #5443](https://github.com/nearai/ironclaw/issues/5443) & [PR #5441](https://github.com/nearai/ironclaw/pull/5441)
    - **分析**: 问题提出者 `italic-jinxin` 指出，自动化任务的结果很容易被忽略，因为用户必须手动打开“Automations”页面才能看到。针对此问题，**贡献者同步提交了 PR #5441**，实现了**头部通知中心**功能。这表明社区和核心团队已就此痛点达成共识，并迅速进入开发阶段，响应非常及时。

3.  **项目内部工程质量的自我监督 (#5437)**：
    - **链接**: [Issue #5437](https://github.com/nearai/ironclaw/issues/5437)
    - **分析**: `pranavraja99` 提交了名为“每日失败分类”的 Issue，披露了 `pinchbench` 测试套件中 146 个失败案例，且错误原因完全一致（`HTTP 400`）。这种主动、公开的失败报告机制，是项目追求高质量和透明度的信号，虽然不直接面向用户，但对开发者社区建立信心至关重要。

#### 5. Bug 与稳定性

今日报告的 Bug 主要集中于以下方面：

- **严重 (Critical)**:
    - **#5420: 自动化交付目标全局默认问题**: 严重的设计缺陷，**尚未有关联的 Fix PR**。
        - [Issue #5420](https://github.com/nearai/ironclaw/issues/5420)
    - **#5426: 创建 Routine 时系统驱动器不可用**: 严重阻碍用户创建自动化流程。
        - [Issue #5426](https://github.com/nearai/ironclaw/issues/5426) (状态: 开放)

- **中等 (Medium)**:
    - **#5429: Web 搜索需要额外 Token**: 即使聊天功能正常，启用 Web 搜索仍需再次认证 NEAR AI API。
        - [Issue #5429](https://github.com/nearai/ironclaw/issues/5429) (状态: 开放)
    - **#5421: Web 搜索非零配置，且重复提示认证**: 进一步确认了 Web 搜索功能在零配置和认证流程上的断裂。
        - [Issue #5421](https://github.com/nearai/ironclaw/issues/5421) (状态: 开放)
    - **#5412: WebUI v2 日志文本不可选中/复制**: 影响开发者调试体验的严重问题。*(已关闭)*
        - [Issue #5412](https://github.com/nearai/ironclaw/issues/5412)

- **测试与工程 (Test/Engineering)**:
    - **#5428: Mock-MCP 测试层缺陷**: 跟踪三个未解决但被归类为测试代码的缺陷。
        - [Issue #5428](https://github.com/nearai/ironclaw/issues/5428) (状态: 开放)
    - **#4108: Nightly E2E 持续失败**: 持续超过一个月未解决的 CI 失败，表明 E2E 测试稳定性是当前短板。
        - [Issue #4108](https://github.com/nearai/ironclaw/issues/4108) (状态: 开放，已超一个月)

#### 6. 功能请求与路线图信号

社区和 QA 团队提出了几个明确的功能需求，其中一些已有相应的 PR 开发中：

- **自动化执行通知** (`#5443/#5441`): **呼声最强**。用户强烈希望在页面任何位置都能感知到自动化任务的完成。关联的 `feat(webui): add header notifications` PR (#5441) 正在开放中，已匹配此需求。
- **零配置 Web Search** (`#5421`): 用户期望 `ironclaw-reborn serve` 能像聊天一样，开箱即用地支持 Web 搜索，无需额外配置 API 密钥。目前仍在 Issue 讨论阶段。
- **日志文本可选性** (`#5412`): 这是一个提升开发者体验的小而重要的功能，已经被修复并关闭。
- **可恢复的失败运行** (`#4841`): 核心 PR #4841 旨在消除“运行出错 (run-borking)”的最终错误，使每次失败都被“要么恢复，要么解释清楚，要么可重试”。此 PR 已保持开放 17 天，**可能是下一个里程碑版本的核心特性**。
- **WebUI 工具权限设置导航** (`#5247`): 一个 UX 改进，旨在让用户更容易找到全局自动批准的设置，体现了对细节体验的打磨。

#### 7. 用户反馈摘要

从 Issues 评论和描述中，可以提炼出以下真实用户痛点：

1.  **配置体验断裂** (#5420, #5421): 用户对部分功能（如 Routine 交付路由、Web Search 认证）的配置必须依赖**全局状态**感到困扰。他们期望更细粒度的、可独立管理的配置，这是对平台成熟度的基本要求。
2.  **测试基础设施不稳定** (#5437, #4108): `Daily ironclaw failure taxonomy` 这样的 Issue 表明内部 QA 工程师也在为此头疼。大量测试因基础设施问题（如模型服务返回 `HTTP 400`）而失败，而非代码逻辑错误。这会侵蚀开发者对 CI 结果的信心。
3.  **前端可用性细节** (#5412): 日志文本无法复制，虽然是小问题，但在调试时**严重打断工作流**。用户对此有明确的不满。
4.  **行为确定性** (#5404): 从 PR #5404 的“聊天输入框发送后清空”功能中，可以看到用户对**提交即确认**这个行为的依赖性。PR 增加了“发送失败后恢复草稿”的逻辑，这表明用户在之前可能因为发送失败而丢失了长篇输入内容。

#### 8. 待处理积压

以下 Issue 或 PR 长期未获更新或关闭，可能成为项目的隐患，需提醒维护者关注：

- **`#4108` Nightly E2E 失败**: 自 **2026-05-27** 以来已报告超过一个月，目前状态仍为“开放”。持续性的 CI 失败会严重拖累开发效率。
    - [Issue #4108](https://github.com/nearai/ironclaw/issues/4108)
- **`#4841` “无运行出错”核心修复**: 该 PR 自 **2026-06-13** 开放，已存在 17 天，属于复杂度高 (XL)、影响面广的核心架构改进。若无明确进展，会影响项目整体的健壮性声誉。
    - [PR #4841](https://github.com/nearai/ironclaw/pull/4841)

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，根据您提供的 LobsterAI GitHub 数据，我为您生成了以下项目动态日报。

---

### LobsterAI 项目动态日报 (2026-06-30)

---

### 1. 今日速览

今日项目活跃度极高，24小时内合并/关闭了13个Pull Request并发布了新版本，表明开发节奏紧凑且交付效率高。社区反馈活跃，主要集中在**性能问题**（如模型响应速度慢、窗口假死）和**功能体验优化**（如定时任务、技能管理）上。新版本 `2026.6.29` 主要聚焦于 OpenClaw 稳定性修复和 Cowork 功能优化，是一条健康的发布线。整体来看，项目处于快速迭代期，社区参与度与代码贡献度均表现出色。

### 2. 版本发布

**新版本: `LobsterAI 2026.6.29`**
- **发布日期:** 2026-06-29
- **主要内容:**
    - **OpenClaw 稳定性:** 修复了 agent 引导工作区与任务工作目录分离的问题，确保 agent 身份/记忆文件不被项目目录覆盖。
    - **Cowork 协作功能增强:** 清理了导航轨的预览界面，优化了会话切换体验。同时，修复了 `feature` 分支合并 `release` 分支时的意外回滚。
    - **权限与插件:** 改进了 OpenClaw 插件的审批路由。
- **破坏性变更:** 无明确提及。
- **迁移注意事项:** 建议所有用户升级至此版本以获得稳定性改进。对于使用 OpenClaw 并有自定义 agent 配置的用户尤其重要。

### 3. 项目进展

今日团队通过高效的代码合并，推动项目在稳定性和用户体验上迈进了关键一步。

- **OpenClaw 任务执行优化 (#2234, #2232):** 修复了 cron 任务中，子 agent 完成后无法驱动父 agent 继续执行的问题，并提供了内置的令牌限制 (`maxTokens`) 后备方案，防止因配置文件缺失导致模型调用失败。
- **Cowork 会话界面修复 (#2226, #2223):** 重构了会话轨（Conversation Rail）的功能，修复了工具提示显示干扰、导航懒加载异常等问题，提升了用户浏览长会话时的流畅度和信息清晰度。
- **诊断与可观测性增强 (#2229):** 合并了针对 Cowork 会话加载和 OpenClaw 运行时错误的诊断日志，这将极大提升未来线上问题的排查效率。
- **版本合入 (#2228):** 将 `release/2026.6.29` 分支成功合入 `main` 主分支，标志着该版本的开发周期正式结束，所有修复已进入主代码库。

### 4. 社区热点

今日社区讨论热度最高的议题集中在**性能体验**与**功能合理性**上。

- **性能对比与 Token 消耗 (#2230):** 用户 `woxinsj` 报告了严重的性能不一致问题。在相同模型和提示词下，LobsterAI 耗时 25 分钟并消耗 60M Token，而竞品 CodeBuddy 仅耗时 2 分 24 秒且消耗 67K Token。此问题引起了高度关注，因为它触及 AI 助手核心的效率和成本痛点。
    - 链接: https://github.com/netease-youdao/LobsterAI/issues/2230
- **Task 预输入与运行时长 (#2120):** 用户 `nbjoe` 提出的关于任务预输入（参照 WorkBuddy）和延长单次任务运行时间的建议，获得了 2 条讨论。这反映了高级用户期望在自动化工作流中获得更连续、更可靠的执行体验。
    - 链接: https://github.com/netease-youdao/LobsterAI/issues/2120

### 5. Bug 与稳定性

今日报告的 Bug 主要集中在以下几个方面：

- **[严重] 极端性能问题 (#2230):** 同一模型在 LobsterAI 上的运行耗时和 Token 消耗是竞品的数十倍甚至上百倍。此问题可能指向框架下层（如模型调用、上下文管理或流式处理）存在重大问题。**尚无对应 fix PR。**
- **[高] 执行结果窗口顶端假死 (#2079):** 在 `2026.5.27` 版本中，滚动到窗口顶部会导致应用假死。此问题已持续一个月，可能是一个较为顽固的 UI 渲染或事件处理 Bug。**尚无对应 fix PR。**
- **[中] 重复输出消耗 Token (#2121):** 用户 `nbjoe` 怀疑系统存在重复输出文本（疑似 Claude 的“思维链”或中间产物被显示），导致不必要的 Token 消耗。这可能是 UI 显示层或流式处理的 Bug。**尚无对应 fix PR。**
- **[低] 文件选择与任务管理遗留问题 (#1426, #1427):** 虽然这两个 Issue 今日被关闭，但其中提到的“无上传成功提示”、“技能重复添加”等问题，在之前的版本中影响了用户添加技能的基本体验。它们已被标记为 `stale`，意味着长期未解决，今日关闭可能是由于版本迭代或不再重现。

### 6. 功能请求与路线图信号

用户提出的功能请求预示着项目未来的发展方向：

- **任务连续性增强:** 用户 `nbjoe` 明确提出了“任务预输入”功能（类似 WorkBuddy），这能极大提升自动化工作流的编排体验。结合今日已合入的 PR (#2234) 中对 cron 任务父子 agent 逻辑的修复，研发团队很可能正在加强任务调度这一块。
- **UI/UX 优化建议 (#2120):** 关于技能展示界面从双列改为三列的建议，以及会话分组管理的需求，表明用户对更高信息密度和更优的信息架构有强烈需求。
- **特殊 Agent 支持 (#2131):** 用户询问对 `hermes agent` 的支持计划，表明社区对集成更多类型或特定功能的 Agent 有期待。虽然目前无直接回复，但这是一个值得关注的路线图信号。

### 7. 用户反馈摘要

从今日的 Issues 和 PR 评论中，可以提炼出以下真实用户反馈：

- **核心痛点: Token 成本担忧:** 用户 `nbjoe` (#2121) 和 `woxinsj` (#2230) 都明确表达了因系统行为导致 Token 消耗过大带来的经济和使用效率上的困扰。这是一个需要优先解决的信任问题。
- **任务管理体验有待提升:** 用户 `nbjoe` (#2120) 期望能预输入任务，并在任务中断后能有更长或更智能的运行控制，说明当前的任务编排和执行模型相对“原子化”，缺乏流程管理的灵活性和鲁棒性。
- **对功能期望的落差:** 用户 `isaiah5818` (#1381) 期望 cron 任务的输出能聚合在一个会话中而非每次都新建，这表明对“定时任务”的认知更接近于一个长期运行的“自动化流程”，而非独立的“查询请求”。

### 8. 待处理积压

以下为长期未获回应的旧 Issue，可能影响用户满意度，建议维护团队予以关注：

- **[长期积压] 定时任务会话聚合 (#1381):** 创建于 2026-04-03，至今未关闭。用户的核心诉求（cron 任务结果聚合到一个会话）是一个合理的功能改进点。
    - 链接: https://github.com/netease-youdao/LobsterAI/issues/1381
- **[长期积压] 微信机器人消息同步问题 (#1383):** 创建于 2026-04-03，影响与机器人用户的交互基础功能。
    - 链接: https://github.com/netease-youdao/LobsterAI/issues/1383
- **[长期积压] 多文件上传问题 (#1384):** 即使有对应的 PR 修复，但 Issue 本身已开放超过两个月，应确认是否已修复并关闭它。
    - 链接: https://github.com/netease-youdao/LobsterAI/issues/1384
- **[长期未合并 PR] 修复多文件选择问题 (#1372):** 该 PR 已于 4 月 2 日创建以修复上述 Issue，至今未合并也未关闭，存在资源浪费和社区贡献者的挫败感风险。
    - 链接: https://github.com/netease-youdao/LobsterAI/pull/1372
- **[依赖更新 PR] (#1277):** 一个来自 `dependabot` 的 Electron 依赖更新 PR，自 4 月 2 日创建以来一直处于打开状态。长期搁置依赖可能导致安全风险和兼容性问题。
    - 链接: https://github.com/netease-youdao/LobsterAI/pull/1277

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

好的，作为AI智能体与个人AI助手领域的开源项目分析师，我将根据您提供的CoPaw项目数据，为您生成一份结构清晰、数据驱动的项目动态日报。

---

### CoPaw 项目动态日报 | 2026-06-30

**项目健康度评估：** ⭐⭐⭐⭐⭐ (非常活跃)

---

#### 1. 今日速览

今日项目保持极高的活跃度，开发节奏迅猛。尽管没有新版本发布，但提交了数量惊人的代码贡献（50条PR），其中近半数已被合并。社区反馈同样踊跃，既有对关键Bug的彻底修复（涉及前端崩溃和模型兼容性），也涌现了大量关于模型配置、记忆搜索和多渠道集成的新功能请求。项目整体正处于深度的功能开发与优化阶段，社区参与度与开发响应速度均处于峰值。

#### 2. 版本发布

*无新版本发布*

#### 3. 项目进展

今日项目在功能增强和Bug修复上取得了显著进展，多个重要Pull Request被合并。

- **主要功能推进：**
    - **[PR #5590]** - 在`/messages/send` API和`channels send` CLI中正式支持**钉钉@提及**功能。这解决了多Agent协作场景下的关键需求，允许Agent在群聊中明确@另一个Agent，提升协作可见性。（[链接](https://github.com/agentscope-ai/QwenPaw/pull/5590)）
    - **[PR #5639]** - 新增技能自动同步功能，允许将技能池中的技能自动同步到选定的Agent，提升了多Agent场景下技能管理的效率。（[链接](https://github.com/agentscope-ai/QwenPaw/pull/5639)）
    - **[PR #5664]** - 实现了多标签页模式下非所有者标签页的提示信息，明确了“仅入队”的职责，优化了前端会话管理的用户体验。（[链接](https://github.com/agentscope-ai/QwenPaw/pull/5664)）

- **关键修复与优化：**
    - **[PR #5654]** - 修复了钉钉渠道在cron任务或静默执行时因发送空消息导致无效通知的问题，提升了渠道消息发送的可靠性。（[链接](https://github.com/agentscope-ai/QwenPaw/pull/5654)）
    - **[PR #5660]** - 修复了Runtime 2.0迁移后`spawn_subagent`功能失效的回归问题，确保了子Agent派生子系统的稳定性。（[链接](https://github.com/agentscope-ai/QwenPaw/pull/5660)）
    - **[PR #5656]** - 修复了侧边栏会话列表滚动在非简单模式下错位的问题。（[链接](https://github.com/agentscope-ai/QwenPaw/pull/5656)）

**项目进展总结：** 项目正快速迭代，社区贡献的修复和功能得到了积极响应，尤其是钉钉集成的完善和Agent核心功能的回归修复，显著提升了项目的实用性和稳定性。

#### 4. 社区热点

- **Issue #5401 (CLOSED)** - **高风险Bug引发广泛关注**：关于“Console前端因大量工具调用历史导致崩溃”的问题，获得了6条评论。该问题被快速定位并关闭，社区成员讨论了前端`DataContent`渲染与后端API响应之间的类型不匹配问题。（[链接](https://github.com/agentscope-ai/QwenPaw/issues/5401)）
- **Issue #5588 (OPEN)** - **核心功能增强呼声最高**：用户提出的“记忆搜索支持专用Reranker模型”的建议，获得了4条评论。这反映了社区对检索增强生成（RAG）质量提升的强烈需求，希望从单阶段检索升级为“粗筛+精排”的两阶段模式。（[链接](https://github.com/agentscope-ai/QwenPaw/issues/5588)）
- **Issue #5403 (OPEN)** - **UI交互痛点的典型代表**：关于“浏览器自动填充功能干扰模型配置页搜索输入框”的问题，引发了5条评论。用户抱怨浏览器会错误地将搜索框识别为密码输入框，这是一个常见且影响使用体验的前端Bug。（[链接](https://github.com/agentscope-ai/QwenPaw/issues/5403)）

#### 5. Bug 与稳定性

今日报告的Bug中，部分问题已得到修复或已有PR修复方案。

- **严重级别 Bug：**
    1.  **Console前端崩溃（已解决）**：[Issue #5401](https://github.com/agentscope-ai/QwenPaw/issues/5401) 描述了因渲染大型工具调用历史导致的前端崩溃和白屏问题。该问题已被关闭，但未提及具体的修复PR，需确认是已合并还是暂时关闭。
    2.  **DeepSeek V4模型兼容性问题（已解决）**：[Issue #5573](https://github.com/agentscope-ai/QwenPaw/issues/5573) 报告了与DeepSeek V4模型在非官方端点上的两类400错误（`reasoning_content`缺失和Schema null类型）。该问题已被关闭，社区用户提出了修复方案并成功验证。

- **中高级别 Bug：**
    3.  **Qwen-Image工具安装失败**：[Issue #5587](https://github.com/agentscope-ai/QwenPaw/issues/5587) 报告了`Qwen-Image`工具安装错误，影响多模态功能。
    4.  **自动化任务无故终止**：[Issue #5616](https://github.com/agentscope-ai/QwenPaw/issues/5616) 报告了自动化任务在无人工干预的情况下莫名终止，这是影响生产服务稳定性的潜在风险。

- **修复中的Bug：**
    5.  **钉钉消息发送失效**：[Issue #5566](https://github.com/agentscope-ai/QwenPaw/issues/5566) 描述的cron任务静默执行和通知不可达问题，已有对应的 **[PR #5654](https://github.com/agentscope-ai/QwenPaw/pull/5654)** 提出修复方案，目前正在审核中。
    6.  **子Agent功能回归**：[Issue #5523](https://github.com/agentscope-ai/QwenPaw/issues/5523) 报告的`spawn_subagent`功能失效，已有对应的 **[PR #5660](https://github.com/agentscope-ai/QwenPaw/pull/5660)** 提出完整修复方案，目前已合并。

#### 6. 功能请求与路线图信号

用户今日提出的新功能请求主要集中在 **Agent记忆能力增强** 和 **多渠道集成精细化** 两方面，这些方向很可能被纳入下一版本的规划。

- **Agent核心能力增强：**
    - **两阶段记忆检索**：[Issue #5588](https://github.com/agentscope-ai/QwenPaw/issues/5588) 建议引入专用的Reranker模型对记忆搜索结果进行二次排序。这符合当前AI Agent提升长期记忆精度的主流趋势。
    - **循环检测机制**：[Issue #5657](https://github.com/agentscope-ai/QwenPaw/issues/5657) 建议增加循环检测机制，防止Agent在工作流中陷入无限循环，这对于提升复杂任务的鲁棒性至关重要。

- **集成与配置优化：**
    - **Telegram自定义BaseURL**：[Issue #5630](https://github.com/agentscope-ai/QwenPaw/issues/5630) 和 **[PR #5651](https://github.com/agentscope-ai/QwenPaw/pull/5651)** 请求支持Telegram频道的自定义API域名，这主要服务于需要代理或有特殊网络环境的用户。相关PR已提交并待合并。
    - **前端配置“绕过Debounce”开关**：[Issue #5663](https://github.com/agentscope-ai/QwenPaw/issues/5663) 建议新增一个开关，让用户决定发送图片/文件后是否绕过输入Debounce，以手动触发Agent响应。这反映了用户对交互控制权的更高需求。

#### 7. 用户反馈摘要

从今日的Issues评论中，可以提炼出用户的核心痛点与使用场景：

- **痛点：** 
    - **性能瓶颈**：用户反馈“钉钉通道卡片流传输输出速度过慢”（[Issue #5603](https://github.com/agentscope-ai/QwenPaw/issues/5603)），对比QwenPaw控制台页面，钉钉端的逐字输出严重影响长内容的查看效率。
    - **消息不可达**：用户反馈“长信息在飞书无法收到，只能通过文件发送”（[Issue #5561](https://github.com/agentscope-ai/QwenPaw/issues/5561)），以及“cron任务静默执行和`channels send`命令在后台不可达”（[Issue #5566](https://github.com/agentscope-ai/QwenPaw/issues/5566)），表明消息分发系统的稳定性是当前用户关注的重中之重。
    - **模型兼容性问题**：用户在连接非官方模型端点时遇到困难，例如无法连接9router转发的模型（[Issue #5658](https://github.com/agentscope-ai/QwenPaw/issues/5658)），以及DeepSeek V4的400错误（[Issue #5573](https://github.com/agentscope-ai/QwenPaw/issues/5573)）。

- **满意点：** 
    - 社区对于 **DeepSeek V4模型问题** 和 **钉钉@提及功能** 的快速响应与解决表示肯定，显示了项目团队和社区对关键问题的解决能力。

#### 8. 待处理积压

以下是一些可能已存在一段时间、尚未被解决或确认的重要Issue和PR，建议维护者关注。

- **长期开放的Bug跟踪：**
    - **[Issue #5273](https://github.com/agentscope-ai/QwenPaw/issues/5273)**：QwenPaw v2.0.0预发布版本问题集中跟踪。作为版本大更新的核心跟踪单，持续更新但长期开放，其内部依赖的子问题解决进度值得关注。
    - **[Issue #5151 (PR)](https://github.com/agentscope-ai/QwenPaw/pull/5151)**：修复GitPanel中由于CSS类名前缀错误导致样式无法生效的问题。该PR已开放超过两周，是影响用户体验的前端显示问题。

- **长期等待合入的功能PR：**
    - **[PR #5187](https://github.com/agentscope-ai/QwenPaw/pull/5187)**：Windows桌面GUI自动化功能。这是一个重量级的新特性，已开放多日，代码量较大，可能正在等待更详细的Review或测试。
    - **[PR #5525](https://github.com/agentscope-ai/QwenPaw/pull/5525)**：实现Windows原生沙箱。同样是一个涉及平台底层的复杂功能，开发时间较长。

**分析师注：** 项目今日的PR处理效率极高，但仍有大量待合并的PR（26条）。考虑到这些积压PR中包含多个关键功能，建议维护团队加快审核速度，以保持社区的贡献热情。同时，钉钉、飞书等渠道的性能和稳定性问题，以及模型兼容性问题，是提升用户体验的核心瓶颈，应立即投入资源解决。

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是基于您提供的 ZeroClaw (github.com/zeroclaw-labs/zeroclaw) GitHub 数据生成的 2026-06-30 项目动态日报。

---

## ZeroClaw 项目日报 — 2026-06-30

### 1. 今日速览

项目今日处于 **v0.8.3 版本的密集冲刺阶段**，开发活动极为活跃。昨日共提交了 50 个 PR，但合并效率略有下降（仅合并完成 9 个），导致待合并积压升至 41 个。核心维护者 `Audacity88` 今日新开和更新了多个 v0.8.3 的子追踪 Issue，明确将庞大的版本工作分解为 Provider 序列化、Channel 适配器、运行时策略等多个并行执行流。社区贡献持续涌入，尤其以新的渠道集成（Inkbox、Matrix流式消息、GitForge）最为瞩目。总体来看，项目健康度 **优秀**，开发动力强劲，但需关注 PR 审查及合并效率。

### 2. 版本发布

*(无新版本发布)*

### 3. 项目进展

今日合并/关闭了 9 个 PR，但关键成果集中在修复长期存在的 Bug 和提升代码质量上，为 v0.8.3 的稳定发布奠定基础。

- **核心运行时 Bug 修复**：PR [#8003](https://github.com/zeroclaw-labs/zeroclaw/pull/8003) 已合并，解决了 [#7889](https://github.com/zeroclaw-labs/zeroclaw/issues/7889)`session_end` 钩子从未被触发的问题，这是一个长期存在的“死代码”Bug，其修复提升了会话生命周期的完整性。
- **Provider 稳定性提升**：PR [#8148](https://github.com/zeroclaw-labs/zeroclaw/pull/8148) 已合并，修复了 Anthropic 提供者在流式请求构建器中因 `serde_json` 序列化失败导致 `panic` 的潜在问题，将其转为优雅的错误传播。
- **测试覆盖增强**：贡献者 `Alix-007` 持续发力，其多个仅含测试的 PR（[#8458](https://github.com/zeroclaw-labs/zeroclaw/pull/8458)、[#8267](https://github.com/zeroclaw-labs/zeroclaw/pull/8267)、[#8269](https://github.com/zeroclaw-labs/zeroclaw/pull/8269) 等）已被合并，为日志模块、工具 I/O 捕获和配置文件解析等关键组件增加了单元测试，项目整体代码质量和回归防护能力得到稳步提升。

### 4. 社区热点

尽管 PR 评论数据未提供，但从 PR 的规模和类型可以判断，以下议题是今日社区关注的焦点：

- **新渠道集成成为热点**：贡献来自 `dimavrem22` 的 [PR #8384](https://github.com/zeroclaw-labs/zeroclaw/pull/8384)（新增 Inkbox 渠道）和 `Nillth` 的 [PR #8504](https://github.com/zeroclaw-labs/zeroclaw/pull/8504)（新增 Git Forge 渠道）均为 **XL 规模**的巨型 PR。这反映出社区对扩展 ZeroClaw Agent 触达能力（尤其是 Email、SMS 及开发者工作流集成）的强烈需求。用户不仅希望 Agent 在聊天应用中存在，更希望它能主动参与和响应更多的数字交互场景。
- **Matrix 渠道体验优化**：`vrurg` 提交的 [PR #8443](https://github.com/zeroclaw-labs/zeroclaw/pull/8443) 为 Matrix 渠道引入了“单消息流式草稿”模式，这直接回应了用户对更流畅、更自然的聊天体验的期待，将 Agent 内部思考过程与最终回复在视觉上优雅区分，是提升产品可用性的重要一步。

### 5. Bug 与稳定性

- **中风险**：修复了 `session_end` 钩子长期未触发的 Bug ([`#7889`](https://github.com/zeroclaw-labs/zeroclaw/issues/7889))。该问题导致会话结束后的清理或通知逻辑无法执行。**已有 Fix PR [#8003](https://github.com/zeroclaw-labs/zeroclaw/pull/8003) 已合并。**
- **中风险**：新增 [PR #8465](https://github.com/zeroclaw-labs/zeroclaw/pull/8465) 修复了 `cron` 定时任务的优雅关闭问题，通过引入 `CancellationToken` 确保调度器能够随守护进程一同安全停止，防止资源泄漏。
- **低风险**：PR [#8535](https://github.com/zeroclaw-labs/zeroclaw/pull/8535) 修复了在 Windows 平台下编译时，因 Shell 测试函数未正确使用 `#[cfg(unix)]` 导致编译警告的问题。
- **高风险**：高风险的 [PR #8496](https://github.com/zeroclaw-labs/zeroclaw/pull/8496) 旨在集中化 MCP 工具访问策略，解决一个重要的安全/权限一致性问题，目前仍处于开放状态，需要维护者重点审查。

### 6. 功能请求与路线图信号

- **v0.8.3 发布路线图明确化**：核心开发者 `Audacity88` 今日维护了多个 v0.8.3 的跟踪 Issue，包括 [`#8360`](https://github.com/zeroclaw-labs/zeroclaw/issues/8360) (Provider 序列化)、[`#8362`](https://github.com/zeroclaw-labs/zeroclaw/issues/8362) (Channel 适配器)、[`#8363`](https://github.com/zeroclaw-labs/zeroclaw/issues/8363) (运行时策略) 等。这表明下一版本的开发重点在于 **底层消息序列化、渠道行为标准化、以及配置驱动的灵活运行时策略**。
- **新的渠道功能**：多个 Open PR 预示着新功能即将到来：
    - **Git Forge 集成**：`Nillth` 的 [PR #8504](https://github.com/zeroclaw-labs/zeroclaw/pull/8504) 将为 Agent 增加与 GitHub 等代码托管平台交互的能力，这是一个重要的开发者工具功能。
    - **地理位置支持**：`Leuca` 的 [PR #8427](https://github.com/zeroclaw-labs/zeroclaw/pull/8427) 为 WhatsApp 渠道添加了发送和接收地理位置信息的能力，拓展了 Agent 在移动场景下的应用。
- **用户界面改进**：`LiLan0125` 的 [PR #8477](https://github.com/zeroclaw-labs/zeroclaw/pull/8477) 让 ZeroCode 用户在活动会话中可以切换 Agent，这直接改善了用户交互流畅度，很可能被纳入后续版本。

### 7. 用户反馈摘要

由于今日数据中 Issues 评论较少，我们从活跃的 PR 中提炼出以下用户诉求：

- **稳定性和可靠性是首要关注**：社区贡献者在解决 `session_end` 钩子、串行化 Panic、Shell 编译警告等问题上投入了大量精力，这表明用户在实际运行中遇到了因这些细节导致的不稳定或兼容性问题，对系统的稳健运行有较高期待。
- **对扩展 Agent 交互能力有强烈需求**：Inkbox、Git Forge 等大 PR 的出现，表明用户社区不仅满足于聊天机器人，更希望 ZeroClaw Agent 能成为真正参与多平台、多形态任务的“数字员工”，尤其是深入到 Email、短信、以及开发者惯用的 Git 工作流中。
- **对可用性细节的关注**：Matrix 的流式消息草稿、ZeroCode 内的 Agent 切换、WhatsApp 的 emoji 回复和定位支持，这些都反映了用户对“体验细节”的追求，期望 ZeroClaw 在各个渠道上都能提供与原生应用媲美的交互体验。

### 8. 待处理积压

以下 PR/Issue 因长期未获得响应或处于“待作者操作”状态，值得维护者关注：

- **高风险、需审查**：[PR #8496](https://github.com/zeroclaw-labs/zeroclaw/pull/8496) (`fix(tools/mcp): centralize deferred-MCP access policy...`) 状态为 `needs-author-action`，但其解决的核心安全问题影响较大，维护者可能需要主动介入审查和沟通。
- **长期未合入的渠道功能**：[PR #7535](https://github.com/zeroclaw-labs/zeroclaw/pull/7535) (`feat(whatsapp-web): implement add_reaction and remove_reaction`) 和 [PR #7637](https://github.com/zeroclaw-labs/zeroclaw/pull/7637) (`fix(zerocode): auto-normalise agent alias input...`) 均已被标记为 `stale-candidate`。这两个 PR 分别涉及功能增强和用户体验修复，虽非核心变更，但长期搁置可能导致代码冲突和社区贡献者积极性受挫，建议维护者尽快决定是否纳入合并。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*