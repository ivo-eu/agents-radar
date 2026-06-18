# OpenClaw 生态日报 2026-06-18

> Issues: 289 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-06-18 12:31 UTC

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

# OpenClaw 项目日报 — 2026-06-18

## 今日速览

过去 24 小时项目保持极高活跃度：共产生 **289 条 Issue 更新**（其中新开/活跃 277 条，关闭 12 条）和 **500 条 PR 更新**（其中合并/关闭 43 条，待合并 457 条）。社区贡献热情高涨，但 PR 合并率仅约 8.6% (43/500)，说明审查资源仍面临较大压力。**无新版本发布**，多数修复和功能请求仍处于讨论或等待维护者决策阶段。安全、会话状态持久性和消息丢失类问题成为今日讨论焦点。

## 版本发布

（无新版本发布）

## 项目进展

今日共有 **43 个 PR 被合并/关闭**，其中重点推进包括：

- **Autofix 流水线与 Windows 守护进程**（[PR #68936](https://github.com/openclaw/openclaw/pull/68936)）—— 已合并，增加了基于 Claude Agent SDK 的自动审查修复管道和 Windows 后台守护进程，提升工程效率。
- **CLI 状态/探测路径回归修复**（[PR #75223](https://github.com/openclaw/openclaw/pull/75223)）—— ClawSweeper 机器人提交的回归修复，已进入维护者就绪状态。
- **打字的启动变为即发即忘**（[PR #75403](https://github.com/openclaw/openclaw/pull/75403)）—— 修复了打字指示器在清理时可能残留的问题。
- **XMemo 云端内存插件**（[PR #94435](https://github.com/openclaw/openclaw/pull/94435)）—— 虽已关闭（可能因重复或外部插件候选），但展示了社区对第三方记忆后端的兴趣。
- 此外，**43 个关闭的 PR** 中包括多个小尺寸清理、文档修正和测试增强，整体上项目在 **编排稳定性、Telegram 通道行为、内存索引一致性**等方面持续微调。

## 社区热点

| Issue | 标题 | 评论数 | 👍 数 | 核心诉求 |
|-------|------|--------|-------|---------|
| [#50090](https://github.com/openclaw/openclaw/issues/50090) | Community Skill Development & ClawHub | 15 | 2 | 推动技能生态从“承诺”走向“实践”，降低社区贡献门槛 |
| [#58450](https://github.com/openclaw/openclaw/issues/58450) | Agent can promise a later follow-up without starting any actual follow-up action | 15 | 3 | Agent 虚假承诺行为导致用户信任受损，要求明确后续动作执行 |
| [#65161](https://github.com/openclaw/openclaw/issues/65161) | Heartbeat isolated mode: cadence stalls, … | 14 | 1 | 心跳隔离模式多项回归，影响系统可用性 |
| [#57326](https://github.com/openclaw/openclaw/issues/57326) | CLI-backed helper paths still bypass CLI dispatch | 13 | 1 | 安全/路由绕过问题，仍有残存表面未修复 |
| [#91016](https://github.com/openclaw/openclaw/issues/91016) | ⚠️ 升级后 DeepSeek Prompt Cache 完全失效，一小时烧掉 ~$6 | 12 | 6 | **严重成本损失**，用户升级路径后缓存失效 |

**分析**：今日最热 Issue 集中在 **Agent 行为可靠性**和**成本控制**上。#58450 获得的 👍 较多，表明用户对 Agent “说谎” 的容忍度极低。#91016 虽然已关闭（可能是已识别或临时缓解），但 6 个 👍 和 12 条评论说明该问题在用户群体中引发广泛关注。

## Bug 与稳定性

按严重程度排列，重点 Bug 如下：

| ID | 标题 | 等级 | 影响 | 是否有 Fix PR |
|----|------|------|------|--------------|
| [#91016](https://github.com/openclaw/openclaw/issues/91016) | DeepSeek Prompt Cache 完全失效烧钱 | **高（成本）** | 用户直接经济损失 | 已关闭，可能有热修复 |
| [#63216](https://github.com/openclaw/openclaw/issues/63216) | Repeated hard resets on same session key despite high reserveTokensFloor | P1 | 会话状态丢失、循环重启 | 无 |
| [#67777](https://github.com/openclaw/openclaw/issues/67777) | Subagent completion delivery lost on timeout/drain/orphan prune | P1 | 消息丢失 | 无 |
| [#57326](https://github.com/openclaw/openclaw/issues/57326) | CLI-backed helper paths still bypass CLI dispatch | P1 | 安全旁路 | 无 |
| [#65624](https://github.com/openclaw/openclaw/issues/65624) | Mattermost slash commands expose reusable tokens via cleartext URLs | P1 | 安全（CVSS 7.6~8.6） | [PR 关联](https://github.com/openclaw/openclaw/issues/65624) |
| [#64267](https://github.com/openclaw/openclaw/issues/64267) | Agent internal thinking (English) exposed to user | P1 | 信息泄露 | 无 |
| [#65374](https://github.com/openclaw/openclaw/issues/65374) | Dreaming system contaminates agent identity in multi-agent setups | P1 | 数据污染 | 无 |
| [#64810](https://github.com/openclaw/openclaw/issues/64810) | Heartbeat/async events swallow in-progress replies in Telegram | P1 | 消息丢失 | 无 |
| [#66443](https://github.com/openclaw/openclaw/issues/66443) | Overflow recovery duplicates user messages | P1 | 会话膨胀 | [PR 关联](https://github.com/openclaw/openclaw/issues/66443) |
| [#67419](https://github.com/openclaw/openclaw/issues/67419) | Bootstrap files re-injected every turn, wasting 20-30% tokens | P2 | 性能/成本 | 无 |

另外有多项 **P2 级别**的稳定性 Bug 如 [#67288](https://github.com/openclaw/openclaw/issues/67288)（Bedrock 插件无谓的 IAM 发现）、[#67417](https://github.com/openclaw/openclaw/issues/67417)（备份因会话文件清理而失败）等，显示项目在边缘路径处理上仍有较大改进空间。

> **注**：上述标有“PR 关联”的 Issue 已有对应的开放 PR，但尚未合并。

## 功能请求与路线图信号

今日收集到的功能请求中，以下 4 项关注度最高且与当前 PR 方向存在关联：

1. **Per-agent memory-wiki vault configuration**（[#63829](https://github.com/openclaw/openclaw/issues/63829)）—— 9👍，要求多 agent 独立知识库，与社区对隔离性的普遍需求一致。同类诉求还有 **Per-agent dreaming configuration**（[#67413](https://github.com/openclaw/openclaw/issues/67413)，5👍）。
2. **敏感数据脱敏**（[#64046](https://github.com/openclaw/openclaw/issues/64046)）—— 配置文件、日志、UI 中 API Key 等敏感信息明文存储，用户呼吁统一脱敏方案。
3. **Per-agent TTS/STT configuration overrides**（[#66252](https://github.com/openclaw/openclaw/issues/66252)）—— 支持多语言/多语音代理，符合国际化部署趋势。
4. **Plugin UI Extension System**（[#66944](https://github.com/openclaw/openclaw/issues/66944)）—— 允许插件向 Control UI 贡献原生页面，4👍，提升可扩展性。

**路线图信号**：上述功能请求均无对应的已合并 PR，但部分已有 `linked-pr-open` 标记（如 #66944 关联 PR 未展示），可能已在开发中。此外，针对内存索引、技能生态（#50090）、多索引嵌入（#63990）等基础设施级请求长期未取得实质进展，建议维护团队优先投入资源。

## 用户反馈摘要

- **成本敏感**：用户 `RavenSS213` 在 [#91016](https://github.com/openclaw/openclaw/issues/91016) 中详细记录了升级后 Prompt Cache 失效导致每小时多花 ~$6，要求紧急修复并增加缓存状态监控。
- **信任危机**：用户 `al-osokin` 报告 [#58450](https://github.com/openclaw/openclaw/issues/58450) 中 Agent 承诺“稍后跟进”但实际无任何后台动作，认为这严重损害用户体验，建议在对话末尾显示实际启动的任务列表。
- **多 agent 混淆**：用户 `SweetSophia` 在 [#65374](https://github.com/openclaw/openclaw/issues/65374) 中指出内置 dreaming 系统将不同 agent 的会话混合，导致身份污染，期望增加 agent 边界。
- **集成兼容性**：用户 `RavenSS213` 之外，多位用户反映 Telegram、Slack、QQ 等通道存在消息丢失或状态不一致问题，社区对多通道的稳定性测试覆盖率表示疑虑。
- **正面反馈**：以 `liuhao1024` 为代表的贡献者持续提交小尺寸修复（如 Telegram thread label、macOS 权限缓存），社区对这类细节改进给予肯定。

## 待处理积压

以下为长期未响应或状态卡顿但影响重大的 Issue/PR，提醒维护者关注：

| 条目 | 创建日期 | 最后更新 | 当前状态 | 为什么需要关注 |
|------|----------|----------|----------|--------------|
| [#50090](https://github.com/openclaw/openclaw/issues/50090) Community Skill Development & ClawHub | 2026-03-19 | 2026-06-18 | OPEN, 15评论 | 社区技能生态是项目增长关键，已停滞 3 个月 |
| [#58957](https://github.com/openclaw/openclaw/issues/58957) Model switch fails silently with large context | 2026-04-01 | 2026-06-18 | stale, P1, 7评论 | 影响模型切换体验，无任何 Fix PR |
| [#59662](https://github.com/openclaw/openclaw/issues/59662) Anthropic Max usage alert text delivered as assistant messages | 2026-04-02 | 2026-06-18 | stale, P2, 5评论 | 频道中暴露无关内容，影响专业场景 |
| [#67820](https://github.com/openclaw/openclaw/pull/67820) fix(whatsapp): reuse active QR and preserve runtime warnings | 2026-04-16 | 2026-06-18 | OPEN, waiting on author | WhatsApp 通道关键修复，但作者未响应维护者请求 |
| [#75403](https://github.com/openclaw/openclaw/pull/75403) fix: typing start fire-and-forget | 2026-05-01 | 2026-06-18 | OPEN, actively grinding | ClawSweeper 机器人提交，但维护者尚未最终审核 |

此外，大量 **P1 级别**的 Bug（如 #63216、#67777、#64267）至今无对应 Fix PR，已积压超过 2 个月，建议项目核心团队在下个迭代窗口优先解决这些问题以挽回用户信任。

---

**总结**：OpenClaw 社区活跃度虽高，但合并瓶颈和长期积压问题日益显著。安全、稳定性、成本控制是当前用户最关心的三大痛点。若能加快关键 Bug 的修复节奏并推进技能生态、多 agent 隔离等重点功能，项目有望进入更健康的增长周期。

---

## 横向生态对比

好的，作为专注于AI智能体与个人AI助手开源生态的资深技术分析师，以下是我基于您提供的2026-06-18各项目动态，生成的横向对比分析报告。

---

### AI智能体与个人AI助手开源生态横向分析报告 (2026-06-18)

**报告摘要：** 2026年6月18日，AI智能体开源生态呈现出“高活跃、强分化、深忧虑”的特点。一方面，社区贡献热情空前高涨，核心项目的功能迭代和版本发布密集；另一方面，安全漏洞、会话稳定性、成本失控等问题成为跨项目的共同痛症。生态正处于从“功能证明”到“生产可用”的残酷转型期，谁能率先解决信任与成本问题，谁就将赢得下一阶段。

#### 1. 生态全景

当前，个人AI助手开源生态整体上已从早期的“模型对话壳”阶段，全面进入“自主智能体”体系构建期。各项目疯狂“上分”，争夺开发者和高级用户的青睐。然而，野蛮生长的代价是系统性问题的集中爆发：安全漏洞频现（SSRF、权限绕过、本地文件泄露）、核心功能可靠性堪忧（上下文丢失、会话冻结、Agent“说谎”）、以及成本控制意识的全面觉醒。生态正处于一个**痛苦的“功能补全”向“质量巩固”的过渡期**，社区的耐心正在被测试，对稳定性和安全性的诉求已压倒对功能数量的追求。

#### 2. 各项目活跃度对比

| 项目名称 | Issues (新/活跃) | PRs (新/活跃/合并) | Release (今日) | 健康度评估 |
| :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 289 (277/12) | 500 (457/43) | 无 | **中高** (社区极度活跃，但合并瓶颈严重，积压问题多) |
| **NanoBot** | ~7 (新) | 31 (16/15) | 无 | **高** (协作高效，PR合并率高，迭代节奏好) |
| **Hermes Agent** | 6 (新) | 50 (47/3) | 无 | **中高** (PR积压严重，但核心问题修复及时，方向明确) |
| **PicoClaw** | 4 (4/2) | 8 (4/4) | 无 | **中等** (修复与安全为今日重点，功能请求停滞) |
| **NanoClaw** | 6 (新) | 15 (11/4) | 2个 (v2.1.0, v2.1.17) | **高** (高效发布，同时推进修复与新特性，破坏性变更需关注) |
| **NullClaw** | 3 (活跃) | 3 (待合并) | 无 | **中等** (项目活跃度不高，但三个PR精准回应了社区长期痛点) |
| **IronClaw** | ~7 (新) | 高 | 无 | **高** (开发与修复闭环高效，社区关注点从“能用”向“好用”升级) |
| **LobsterAI** | 2 (新) | ~10 (活跃) | 1个 (2026.6.15) | **中高** (新功能“Computer Use”发布是关键动作，但伴随安全漏洞) |
| **TinyClaw** | 3 (新) | 0 | 无 | **低** (三个严重安全漏洞，项目无任何响应，处于“安全危机”状态) |
| **CoPaw** | 33 (活跃) | 31 (活跃) | 1个 (v1.1.12.post1) | **高** (修复密度高，社区对核心问题反馈激烈，迭代动力足) |
| **Moltis** | 2 (新) | 1 (待合并) | 无 | **中等** (活跃度平稳，关注点集中在配置化和导出功能) |
| **ZeptoClaw** | N/A | N/A | N/A | **休眠** (过去24小时无活动) |
| **ZeroClaw** | 12 (11/1) | 50 (41/9) | 无 | **高** (PR量巨大，修复与功能并行，但合并率偏低，处于密集开发期) |

#### 3. OpenClaw 在生态中的定位

- **优势地位：** OpenClaw以“核心参照”项目自居，其社区规模（单日289条Issue更新、500条PR更新）和影响力无疑是生态中的“巨无霸”。它是社区创新的主要策源地，大量功能请求（如Per-agent memory、敏感数据脱敏）从这里发源并辐射到其他项目。
- **技术路线差异：** 相比于追求“精简实用”的NanoBot或“企业级稳定”的IronClaw，OpenClaw更倾向于一种“包罗万象、社区驱动”的宏大架构。这导致其功能极其丰富，但也带来了**合并瓶颈**（PR合并率仅8.6%）和**长期积压**（关键Bug无人认领）的严重问题。
- **社区规模：** OpenClaw的社区参与度在生态中居首。然而，高参与度并未完全转化为高质量产出（合并率低），说明其社区贡献存在“堰塞湖”效应。**NanoBot**和**IronClaw**则展现出更高效的社区协作模式（合并率更高、版本发布更规律）。

#### 4. 共同关注的技术方向

多个项目不约而同地聚焦于以下技术痛点：

- **安全与权限管理：** （涉及项目：**OpenClaw, NanoBot, PicoClaw, IronClaw, TinyClaw, LobsterAI**）核心诉求是杜绝SSRF（PicoClaw #3140），强化API鉴权（TinyClaw #284），实现细粒度的工具权限和文件访问控制（IronClaw #4960系列），并防止敏感信息泄露（OpenClaw #64046）。
- **会话状态与记忆管理：** （涉及项目：**OpenClaw, ZeroClaw, CoPaw, Hermes Agent**）核心诉求是解决会话丢失（ZeroClaw #7799）、上下文污染（CoPaw #65374）、以及记忆压缩导致的“失忆”问题（CoPaw #5171, OpenClaw #58450）。这与Agent的信任度直接相关。
- **Agent行为可靠性：** （涉及项目：**OpenClaw, ZeroClaw, CoPaw**）社区对Agent“说谎”（虚假承诺）、挂起（进程冻结）、以及忽略安全限制等行为零容忍。这推动了确保工具调用结果、设计可预期的自动化流程等诉求。
- **成本控制：** （涉及项目：**OpenClaw, ZeroClaw, CoPaw**）DeepSeek Prompt Cache失效（OpenClaw #91016）、Bootstrap文件浪费Token（OpenClaw #67419）、以及对子进程OOM的担忧（ZeroClaw #6916），都指向了成本问题。用户从“追求功能”转向“精打细算”。
- **多Agent隔离与独立配置：** （涉及项目：**OpenClaw, NanoBot, IronClaw**）社区对多Agent的想象已经从“聊天”升级到“工作”。他们要求每个Agent拥有独立的记忆、模型、工具、知识库，甚至是认证流程，以实现真正意义上的多任务并行与隔离。

#### 5. 差异化定位分析

- **OpenClaw**: **全能型社区平台**。功能最丰富，社区最大，但资源错配严重。适合希望探索所有前沿功能、愿意忍耐不稳定并参与深度贡献的“硬核”玩家。
- **NanoBot**: **易用型个人助手**。定位清晰，注重开箱即用和配置简化（`--wizard` PR）。合并率高，迭代快，面向“怕麻烦”的个人开发者和小团队。
- **Hermes Agent**: **多通道、重体验的中枢**。强调TUI/GUI和Telegram/Discord等通道的体验一致性。其核心挑战是确保不同交互界面下功能完备且稳定。
- **PicoClaw**: **安全与兼容性的平衡者**。今日修复以安全（SSRF）和LLM兼容性（Gemini）为主，像是一位“稳健的守成者”，适合看重安全与数据隐私的小型托管商或隐私爱好者。
- **ZeroClaw**: **企业级自动化先驱**。代码量巨大，关注点在生产环境下的稳定性和资源管理（OOM控制、定时任务、会话恢复），是企业级用户和运维人员的选择。
- **IronClaw**: **企业级权限与流程管控者**。核心词是“审批”和“权限”。用户关注点从“AI能不能干”转向了“AI怎么被管控”，特别适合需要在安全合规前提下使用AI的公司。
- **CoPaw**: **协作与插件生态探索者**。强调Agent间协作（Cowork）和强大的插件系统（datapaw），旨在构建一个能完成复杂分析和任务的AI Agent网络。
- **LobsterAI**: **桌面端原生Agent试点**。发布“Computer Use”MVP，剑指Windows桌面自动化，开辟了与其他基于云/API的项目截然不同的赛道。

#### 6. 社区热度与成熟度

- **第一梯队 (快速迭代/功能扩展期)： OpenClaw, NanoBot, Hermes Agent, ZeroClaw, CoPaw。** 这些项目日活跃度极高，代码提交频繁，是生态创新的主要引擎。但普遍面临“壮骨（稳定性）”的挑战。
- **第二梯队 (质量巩固/细节优化期)： IronClaw, PicoClaw, LobsterAI。** 项目已具备核心功能，当前阶段侧重于修复关键Bug、提升性能和用户体验，并为下一阶段的重磅功能（如认证系统、桌面Agent）做准备。
- **第三梯队 (特定领域/问题修复期)： NullClaw, Moltis。** 活跃度相对平缓，主要聚焦于解决特定领域（如Moltis的TTS配置）或历史遗留问题（如NullClaw的CLI控制字符）。
- **第四梯队 (安全危机/低活跃期)： TinyClaw, ZeptoClaw。** 前者因安全漏洞而陷入被动，后者则完全进入休眠状态。代表生态中的“失联者”。

#### 7. 值得关注的趋势信号

- **“安全”成为新晋的“生死线”**: TinyClaw和PicoClaw今日的安全报告并非个例。当AI Agent拥有操作文件和网络的能力后，安全已从“加分项”变为“准入门槛”。能够率先提供开箱即用的安全框架（如细粒度权限、沙盒执行、输入输出清洗）的项目，将获得巨大优势。
- **“多Agent”不是选择题，是必答题**: 从OpenClaw的“Per-agent”配置，到NanoBot的“多实例”讨论，再到IronClaw的“工具权限”体系，都指向一个问题：用户不只想要一个Agent，而是希望拥有一个**Agent团队**。如何管理这个团队（隔离、路由、通信、授权），是未来12-18个月的核心战场。
- **“成本”意识从个人蔓延至全行业**: 用户不再是“我觉得好用就行”，而是开始计算Token、缓存、API调用带来的实际开销（OpenClaw #91016, ZeroClaw #7492）。这将催生“成本控制即feature”的潮流，如Prompt Cache管理、多模型路由、以及更智能的上下文压缩。
- **从“听命令”到“负责任”**: Agent“说谎”不再被容忍（OpenClaw #58450）。行业趋势要求Agent从“确定性响应”迈向“可解释的行动”。用户期望看到Agent做出决策的完整证据链和可追溯的执行日志。
- **统一体验是通往“大众化”的必经之路**: Hermes Agent的TUI功能审计和NullClaw的CLI控制字符修复，都指向用户对**一致、流畅**的交互体验的渴望。无论最终形态是WebUI、TUI还是聊天通道，一个统一且鲁棒的交互框架将是成为“大众应用”的基石。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，根据您提供的 NanoBot 项目 2026-06-18 的 GitHub 数据，为您呈上今日项目动态日报。

---

### **NanoBot 项目动态日报 | 2026-06-18**

**分析师点评：** 项目今日保持高度活跃状态，尤其在 PR 层面呈现出密集的开发与合并节奏。社区贡献者积极性高，不仅有大量新功能提交，也迅速跟进修复了多个关键 Bug。项目整体健康度良好，正向“更易用、更稳定、更专业”的方向稳步演进。

---

#### **1. 今日速览**

今日项目活跃度极高。**PR 更新量大（31条），其中接近半数（15条）已合并/关闭**，显示出维护者与社区的高效协作。**Issues 方面，以 Bug 和功能请求为主（7条开放）**，问题类型指向明确。特别值得注意的是，**关于多实例、多租户的讨论热度持续**，反映了小型团队和高级用户对部署灵活性的真实需求。项目正处于功能扩展与稳定性加固并重的阶段。

#### **2. 版本发布**

今日无新版本发布。

#### **3. 项目进展**

今日合并/关闭了多个关键 PR，显著推进了项目在多方面能力：

- **新服务商与平台支持（扩展生态）：** **七牛云（Qiniu）AI**、**Mistral AI**、**Keenable 搜索** 已被正式支持为内置 Provider 和服务商，为开发者提供了更多选择。 [PR #3643](https://github.com/HKUDS/nanobot/pull/3643), [PR #4351](https://github.com/HKUDS/nanobot/pull/4351), [PR #4350](https://github.com/HKUDS/nanobot/pull/4350)
- **核心通道与用户体验改进：** **WhatsApp 已读回执（蓝勾）** 功能已实现，提升了消息体验的完整性。 [PR #4354](https://github.com/HKUDS/nanobot/pull/4354)
- **安全与权限策略落地：** 合并了关于 **Filesystem 工作区写入策略的澄清与执行** 的 PR，解决了读写路径混乱的潜在安全风险，并明确了 `extra_allowed_dirs` 为只读的行为。[PR #4202](https://github.com/HKUDS/nanobot/pull/4202), [PR #4053](https://github.com/HKUDS/nanobot/pull/4053)

#### **4. 社区热点**

- **#936 [Feature Request: Add Multi-Tenant Gateway for Multiple Agents]**
  [Issue](https://github.com/HKUDS/nanobot/issues/936)
  - **分析：** 此 Issue 虽然创建于 2 月，但在今日依旧被更新和关注。其背后反映了**中小规模开发者运行多个 Agent 时存在的资源浪费和管理复杂度问题。** 社区期望能通过一个统一 Gateway 管理多个 Agent，这与项目近期在向导（Wizard）和多实例（Multi-instances）方面的改进方向高度一致，是社区的核心诉求之一。
- **#4374 [bug: SOUL.md/USER.md read/write asymmetry]**
  [Issue](https://github.com/HKUDS/nanobot/issues/4374)
  - **分析：** 该 Issue 精准指出了 WebUI 项目工作区中配置文件读写路径不一致的 Bug，获得了社区成员 `maximilize` 的深入分析与报告。这表明用户已经开始利用项目的高级功能（如项目工作区），并对其行为逻辑提出了更高要求。

#### **5. Bug 与稳定性**

今日报告的 Bug 涵盖了从核心逻辑到前端 UI 的多个层面，按严重程度排列如下：

- **严重（可能导致数据丢失或逻辑错误）：**
    - **#4307 [Bug: Post-turn consolidation wipes the agent's own delivery message]**
      [Issue](https://github.com/HKUDS/nanobot/issues/4307)
      - **状态：** 已有关联修复 PR [#4373](https://github.com/HKUDS/nanobot/pull/4373)（`fix(memory): preserve delivery context during consolidation`）。
      - **分析：** 这是一个关于 Agent 记忆/上下文管理的高危 Bug。当 token 限制触发摘要/归档时，属于 Agent 自身的“交付消息”被错误清除，导致后续用户引用丢失。**此问题已得到及时响应，PR #4373 正在审查中，是今日最高优先级的修复项。**
- **中等（影响特定功能使用）：**
    - **#4375 [bug: Git Command Execution Blocked by Workspace Security Policy]**
      [Issue](https://github.com/HKUDS/nanobot/issues/4375)
      - **状态：** 已有测试性 PR [#4393](https://github.com/HKUDS/nanobot/pull/4393)（`test(exec): cover git commands in workspace subdirectories`）。修复代码已通过 PR #4380 合并，本 PR 仅增加回归测试。
      - **分析：** Git 命令在子目录中被安全策略错误阻止，影响开发者工作流。该问题已被定位并修复，现正在完善测试，稳定性得到增强。
- **较低（影响特定平台或配置）：**
    - **#4388 [bug: [WebUI] iOS Safari 点击输入框触发页面放大]**
      [Issue](https://github.com/HKUDS/nanobot/issues/4388)
      - **分析：** 移动端 UI 体验问题，影响 iOS 用户的使用。

#### **6. 功能请求与路线图信号**

- **易用性与向导优化（高优先级）：**
    - `#4376` 要求改善 `nanobot onboard --wizard` 的体验，使其对非技术用户更友好。
    - `#4390` 提出为“普通人”提供多实例管理的简化方案，隐藏复杂设置。
    - **关联 PR：** [#4395](https://github.com/HKUDS/nanobot/pull/4395) `Improve onboard wizard setup flow` 和 [#4399](https://github.com/HKUDS/nanobot/pull/4399) `add configurable hidden_settings_sections` 正是针对这些诉求的代码提交，**极有可能被纳入下一版本的核心改进中。**
- **多租户与多实例（长期需求）：**
    - `#936` 要求 Multi-Tenant Gateway。
    - `#4390` 要求 Multi-instances 简化。
    - **分析：** 这两条 Issue 揭示了社区从“能用”到“用好”的进阶需求。`Multi-instances` 问题已有部分进展（隐藏设置），但 `Multi-Tenant Gateway` 仍在前序讨论阶段，是未来版本的重要方向。
- **模型配置精细化：**
    - `#4389` 要求为 fallback 模型独立配置 `contextWindowTokens`。
    - **分析：** 这是一个精细化资源管理需求，确保当主模型失败回滚到小窗口模型时，不会被截断。这表明用户正在生产环境中使用多模型策略，对健壮性有较高要求。

#### **7. 用户反馈摘要**

- **痛点确认：** 用户 `maximilize` 在 #4374 中的报告，清晰地展示了具备技术背景的用户如何通过配置项目工作区来组织 Agent。他发现并明确指出了**读写路径不对称**这一细节 Bug，反馈质量非常高。
- **使用场景：** 用户 `bukit-kronik` 在 #4390 中描述了“单机多实例”的具体用例，即通过不同文件夹管理不同配置的 Agent。他希望能**隐藏UI中“普通人”不关心的复杂设置**，这反映了从个人开发者向团队/产品化方向使用的过渡。
- **表达认可：** Issue #4376 获得了 1 个 👍，说明有不少用户认可“改进向导”这一需求，认为当前初始配置对新人不够友好，是普遍存在的痛点。

#### **8. 待处理积压**

- **#936 [Feature Request: Add Multi-Tenant Gateway for Multiple Agents]**
  [Issue](https://github.com/HKUDS/nanobot/issues/936)
  - **状态：** 自 2026年2月21日创建，至今已近 4 个月，更新停留在 6月17日。
  - **建议：** 此请求是社区对高级部署的核心诉求，已有多条后续 Issue 讨论类似话题（如 #4390）。虽然短期内可能不会实现，但**建议项目维护者对此 Issue 打上 `needs-design` 或 `future-roadmap` 标签，并给予简短回应**，告知社区已注意到该需求并在考虑中，以避免长期未响应导致用户期待落空。

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

好的，作为一名专注于 AI 智能体与个人 AI 助手领域的开源项目分析师，我将根据您提供的 Hermes Agent 项目数据，为您生成了2026年6月18日的项目动态日报。

---

### Hermes Agent 项目动态日报 | 2026年6月18日

---

#### 1. 今日速览

今日 **Hermes Agent** 项目保持了极高的活跃度，尤其在开发者侧 (PR) 非常繁忙。虽然无新版本发布，但社区提交了 50 条 Pull Request，其中 47 条待合并，显示出核心功能与 Bug 修复的工作量巨大。Issue 方面，除了常规 Bug 报告，涌现出多个关于 **TUI v2 功能完整性** 和 **跨平台/协议行为一致性** 的讨论，表明项目在从 CLI 向 GUI/TUI 转型及扩展平台支持（Telegram, WhatsApp, Discord 等）的过程中，正在经历精细化打磨阶段，社区对细节体验的要求很高。项目整体健康度良好，但存在大量 PR 积压待处理。

#### 2. 版本发布

- **无新版本发布。**

#### 3. 项目进展

今日合并/关闭的 PR 数量虽少，但意义重大，表明项目在核心功能和关键 Bug 修复上取得了进展。

- **新增特性 (已合并):**
    - **[PR #45987] feat(mem0): add app_id support for project-level memory isolation**：合并了为 Mem0 内存插件添加 `app_id` 支持的功能，实现了项目级别的内存隔离。这是对开发者和大型项目用户非常有用的后端增强。
    - **[PR #48396] fix(spotify): handle invalid_grant refresh token expiry**：合并了针对 Spotify 令牌过期 (`invalid_grant`) 的修复，解决了由于外部服务策略变更（2026年7月20日起刷新令牌每6个月过期）导致的认证问题，体现了项目对第三方服务兼容性的快速响应。

- **安全问题修复 (已合并):**
    - **[Issue #7651] Telegram: Missing user-level access control**：该高优先级安全 Issue 已在今日被关闭，虽然没有明确指明对应 PR，但通常意味着修复已被合并。这表明团队已将 Telegram 集成中的用户级访问控制漏洞视为高优先级的解决目标。
    - **[Issue #48394] WhatsApp: missing per-sender user-ID gating**：作为 Telegram 安全修复的延续，此 Issue 被标记为重复并关闭，暗示团队可能会采用相似策略来统一解决各平台权限问题。

- **项目整体推进**：今日对 **存储层 (Memory)** 和 **外部服务集成 (Spotify)** 的改进，以及对 **安全权限控制 (Telegram, WhatsApp)** 的重视，标志着 Hermes Agent 正从“功能可用”向“安全可靠、企业级可用”迈进。

#### 4. 社区热点

今日社区讨论热度集中在以下几个痛点：

- **[Issue #47917] Desktop build fails after update - electronDist does not exist** (9 条评论, 👍 1)
    - **分析**：这是今日讨论最激烈的 Issue。用户报告了桌面版构建失败的回归问题，即使在应用了前一个 PR 的修复后仍然失败。核心矛盾在于 **Electron 二进制缓存** 在项目更新时被清除，导致重复报错。这暴露了开发环境的构建流程不够健壮，严重影响了贡献者和测试者的开发体验。社区诉求是 **要求一个稳定、可复现的本地构建环境**。

- **[Issue #19753] Auxiliary title generation 404 on custom Anthropic-mode providers** (4 条评论, 👍 2)
    - **分析**：这是一个持续了一个多月的经典Bug。用户在使用自定义 Anthropic 模式提供商（如 Kimi Coding Plan）时，辅助任务（标题生成）因 URL 拼接问题（`/v1` 重复）导致 404。这反映了项目在处理 **非标准或自定义 API Provider** 时的兼容性问题，是阻碍用户使用更灵活、低成本模型的重要障碍。

- **[Issue #12130] TUI v2 feature-parity audit vs v1 CLI** (4 条评论)
    - **分析**：这是一份详细的“功能差异清单”，列举了 TUI v2 相比 CLI v1 缺失的约 23 个命令、`@` 引用、界面元素等。这表明 **TUI v2 的用户界面用户体验远未达到 CLI v1 的完备性**。社区的诉求是强烈希望 TUI 能成为 CLI 的“替代品”而非“降级品”。

#### 5. Bug 与稳定性

今日报告的 Bug 主要集中在 **GUI/TUI**、**多平台兼容性** 和 **后端稳定性** 上。

- **P1 (严重):**
    - **[Issue #48176] OAuth Pro/Max/Team requests rejected with HTTP 400**：OAuth 认证失败。用户使用高级 Claude 账号（Pro/Max/Team）时，请求被拒绝。这直接影响了付费用户的正常使用，是当前的最高优先级 Bug。尚无明确修复 PR。
    - **[PR #47219] fix(approval): close config/.env write-gate bypass**：这是一个待合并的 P1 安全修复 PR，旨在堵住配置文件和 `.env` 文件的写入权限绕过漏洞，维护者需优先关注。

- **P2 (中):**
    - **[Issue #47917] Desktop build fails after update**：如前所述，桌面环境构建回归问题，严重影响开发者。
    - **[Issue #48388] Desktop GUI `_save_cfg()` sorts config.yaml keys**：这是一个新增但影响用户体验的 Bug，桌面 GUI 在保存配置时会打乱用户手动排序的 `config.yaml`。虽不影响核心功能，但降低了高级用户的满意度。**已有对应的修复 PR #48399**。
    - **[Issue #48338] `_append_model_switch_marker` injects “system” role mid-conversation**：在会话中切换模型时，代码会插入一个 `system` 角色消息，这违反了 VLLM/Qwen 等严格 API 的规范，导致 HTTP 400 错误。
    - **[Issue #48386] Gateway sends timestamp metadata → HTTP 422 on strict endpoints**：网关向严格 API（如 Mistral）传递了不兼容的 `timestamp` 字段，导致 422 错误。

- **P3 (低):**
    - **[Issue #20866] 400 format_error on Qwen3.6-27B**：辅助任务中系统消息位置错误，可能与多轮对话处理有关。
    - **[Issue #48406] Desktop inline edit loses unsaved text**：桌面端内联编辑器的焦点丢失导致文本丢失，影响编辑体验。

#### 6. 功能请求与路线图信号

今日用户提出的功能请求显示出强烈的 **平台集成深化** 与 **用户体验优化** 需求。

- **平台特定功能 (高频需求):**
    - **[Issue #48413] Feature: persist /model selection per Discord channel**：用户希望在 Discord 频道级别持久化模型选择，而非仅限全局或单次会话。这表明多频道、多模型场景是真实且普遍的需求。**可能被纳入下一版本**。
    - **[Issue #48400] Feature: Delete visible Telegram messages when using /undo**：用户要求 `/undo` 命令不仅能回滚会话状态，还能同步删除 Telegram 聊天界面中的消息。这是对 **多端状态一致性** 的明确需求。**已有对应的修复 PR #48401**。

- **TUI 与用户体验优化:**
    - **[Issue #48404] Feature: In-Session Raw/Compressed View Toggle**：用户希望在会话中能直观地看到上下文压缩前后的状态，并手动切换。这反映了对 Agent “黑盒”行为的透明性要求。**可能被纳入中长期路线图**。
    - **[Issue #12130] TUI v2 feature-parity audit**：来自社区的详细审查为 TUI v2 的优化提供了清晰的“路线图”。项目团队很可能根据此清单制定 TUI v2 的迭代计划。

#### 7. 用户反馈摘要

- **痛点:** 桌面端构建失败 (`#47917`) 是最具破坏性的痛点，直接阻碍开发者贡献。
- **等待:** 对自定义 Provider (`#19753`) 的兼容性问题已存在一个多月，用户长时间无法使用自己想要的模型，感到沮丧。
- **安全担忧:** 社区对 Telegram 集成中的安全漏洞 (`#7651`) 和 OAuth 认证失败 (`#48176`) 非常敏感，这表明用户对 AI Agent 的安全性和数据隐私有很高期望。
- **误报通知:** 一位用户 (`#48411`) 提醒社区 `uv.exe` 被杀毒软件误报为木马，这虽然是假阳性，但反映了打包和分发过程中潜在的干扰因素。

#### 8. 待处理积压

以下为长期未结但重要的 Issue，需要维护者予以关注：

- **[Issue #19753] Auxiliary title generation 404 on custom Anthropic-mode providers** (创建于 2026-05-04，P3)
    - **状态:** 长期未解决。这是一个影响自定义 API 提供商生态的关键问题，不利于吸引使用第三方 API 的用户。
    - **信号:** 该 Issue 积压时间超过一个月，且有 2 个 👍，表明需求存在但优先级较低。鉴于 P3 优先级，建议在下一次功能排期中考虑。

- **[Issue #12130] TUI v2 feature-parity audit vs v1 CLI** (创建于 2026-04-18，P2)
    - **状态:** 长尾 Issue。TUI 作为项目的未来方向，其功能完备性至关重要。虽然已经获得社区的大量反馈，但 2 个月过去，仍处于开放状态。
    - **信号:** 这是一个极好的社区协作指南，维护者应基于此清单在项目看板中创建任务，并向社区同步进度。

- **[PR #43091] fix(cron): avoid consuming skipped manual triggers** (创建于 2026-06-09，P2)
    - **状态:** 待合并 PR。该 PR 修复了 Cron 调度器中的一个边缘情况，可能导致手动触发的任务被跳过。这对依赖定时任务自动化工作流的用户很重要。

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw 项目日报 2026-06-18

## 1. 今日速览

过去24小时内项目保持中等活跃度：共处理4条Issue（新开2条、关闭2条）和8条PR（待合并4条、已合并/关闭4条）。安全修复与LLM兼容性改进是今日主要方向，重点合入了Gemini 3.5 Flash工具调用修复（#3136）、OneBot内网SSRF防护（#3140）以及Brave搜索静默失败的诊断日志增强（#3141）。同时，社区对去中心化通信协议（SimpleX/Tox、DeltaChat）和新认证库（vodozemac）的需求持续发酵，但部分长期PR仍需维护者关注。

## 2. 版本发布

今日无新版本发布。

## 3. 项目进展

今日合并/关闭了4个重要PR，项目在安全性、LLM兼容性和诊断能力上取得实质进展：

| PR | 标题 | 状态 | 核心价值 |
|----|------|------|----------|
| [#3141](https://github.com/sipeed/picoclaw/pull/3141) | fix(web_search): add diagnostic logging for Brave empty results | 已合并 | 帮助定位Brave API静默返回空结果的问题，为#3125提供诊断手段 |
| [#3136](https://github.com/sipeed/picoclaw/pull/3136) | fix(gemini): set both camelCase and snake_case thought_signature | 已合并 | 修复Gemini 3.5 Flash Agentic推理时工具调用400错误，确保与Google API兼容 |
| [#3140](https://github.com/sipeed/picoclaw/pull/3140) | fix(onebot): block private inbound media fetches | 已合并 | 修复OneBot多媒体URL解析中的SSRF漏洞，禁止拉取内网/本地地址 |
| [#2917](https://github.com/sipeed/picoclaw/pull/2917) | feat(provider): add NEAR AI Cloud provider | 已合并 | 新增OpenAI兼容的NEAR AI云提供商，扩展TEE能力模型选项 |

此外，安全相关PR [#3143](https://github.com/sipeed/picoclaw/pull/3143)（修复ISATAP IPv6下的SSRF绕过）仍处于开放状态，预计近期可合并。

## 4. 社区热点

- **[#3088 [Feature] 使用vodozemac替代libolm](https://github.com/sipeed/picoclaw/issues/3088)**  
  获得2个👍，标记为 `priority: high`。用户强调libolm已停止维护且存在安全风险，vodozemac是官方替代。该议题直接关乎项目核心加密库的升级，社区对此关注度较高，但尚未有对应PR。

- **[#3093 [Feature] 需要SimpleX或Tox网关](https://github.com/sipeed/picoclaw/issues/3093)**  
  用户明确要求增加去中心化通信协议支持（SimpleX, Wire, Tox）。结合已有的PR [#3063](https://github.com/sipeed/picoclaw/pull/3063)（添加DeltaChat网关），反映出社区对隐私保护、联邦通信的强烈偏好，但这些PR均处于待合并/停滞状态，需维护者评估优先级。

## 5. Bug 与稳定性

今日无新报告的崩溃或回归问题，但以下两个已关闭的Issue值得关注：

| Issue | 标题 | 严重程度 | 修复状态 |
|-------|------|----------|----------|
| [#3125](https://github.com/sipeed/picoclaw/issues/3125) | web_search工具在Brave API key迁移至.security.yml后静默失败 | **中** – 功能不可用但无错误提示 | 已关闭；通过PR [#3141](https://github.com/sipeed/picoclaw/pull/3141) 添加诊断日志，但根本原因（是否为Brave响应格式变化）仍需进一步排查 |
| [#3111](https://github.com/sipeed/picoclaw/issues/3111) | Gemini 3.5 Flash工具执行返回400 Bad Request | **高** – 阻断使用 | 已关闭；PR [#3136](https://github.com/sipeed/picoclaw/pull/3136) 修复了thought_signature字段命名问题 |

**值得注意**：PR [#3142](https://github.com/sipeed/picoclaw/pull/3142)（清除子回合ToolResult中的ForUser字段）正在解决异步子Agent完成时可能出现的重复消息推送问题，该问题影响多人对话体验。

## 6. 功能请求与路线图信号

- **加密库升级**： [#3088](https://github.com/sipeed/picoclaw/issues/3088) 要求默认使用vodozemac，并支持编译时选择libolm。该项目与PicoClaw的安全基线高度相关，可能被纳入下一小版本。
- **去中心化通信协议**： DeltaChat网关PR [#3063](https://github.com/sipeed/picoclaw/pull/3063) 和SimpleX/Tox请求 [#3093](https://github.com/sipeed/picoclaw/issues/3093) 显示用户对非中心化IM后端的兴趣。考虑到NEAR AI Cloud提供商已合并（[#2917](https://github.com/sipeed/picoclaw/pull/2917)），项目可能在扩展“自有网关”生态。
- **技能安装类型断言修复**： [#3092](https://github.com/sipeed/picoclaw/pull/3092) 修复skills_install中version/force类型断言忽略ok检查的问题，虽小但影响安装稳定性，已停滞两周。

## 7. 用户反馈摘要

从近期Issue评论中提炼的典型痛点：

- **API配置迁移副作用**：用户`Giordano10`反映将API Key迁移至`.security.yml`后，Brave搜索工具未给出任何错误提示，仅返回空结果，导致调试困难（[#3125](https://github.com/sipeed/picoclaw/issues/3125)）。该问题虽已通过增加日志缓解，但用户希望更好的自动检测或降级提示。
- **Gemini模型兼容性**：同一用户指出Gemini 3.5 Flash要求`thought_signature`使用snake_case字段，而PicoClaw只发送camelCase，导致请求被拒（[#3111](https://github.com/sipeed/picoclaw/issues/3111)）。修复后用户无后续反馈，预计问题已解决。
- **SSRF防护仍有盲区**：PR [#3143](https://github.com/sipeed/picoclaw/pull/3143) 的提交者发现ISATAP IPv6地址可嵌入私有IPv4绕过`web_fetch`的SSRF守卫，说明安全测试需要持续覆盖IPv6转换机制。

## 8. 待处理积压

以下议题/PR长期未获响应，建议维护者优先关注：

| 序号 | 链接 | 类型 | 创建时间 | 状态 | 提醒原因 |
|------|------|------|----------|------|----------|
| 1 | [#3088](https://github.com/sipeed/picoclaw/issues/3088) | Issue (Feature) | 2026-06-09 | 开放，priority: high | 核心加密库安全升级，9天无实质性进展 |
| 2 | [#3093](https://github.com/sipeed/picoclaw/issues/3093) | Issue (Feature) | 2026-06-10 | 开放，已被标记stale | 用户多次提及SimpleX/Tox，但无维护者回应 |
| 3 | [#3092](https://github.com/sipeed/picoclaw/pull/3092) | PR (fix) | 2026-06-10 | 开放，stale | 修复类型断言bug，8天无review/merge |
| 4 | [#3063](https://github.com/sipeed/picoclaw/pull/3063) | PR (feat) | 2026-06-08 | 开放 | DeltaChat网关新功能，10天未合并，存在冲突风险 |
| 5 | [#3143](https://github.com/sipeed/picoclaw/pull/3143) | PR (fix) | 今日新开 | 开放 | 修复SSRF绕过，属于安全补丁，建议快速review |

---

**总结**：PicoClaw今日在兼容性、安全性和诊断能力上均有正向推进，但去中心化通信和加密库升级等“硬骨头”需求仍悬而未决。建议维护者在下一个版本发布前，优先合并安全修复PR [#3143](https://github.com/sipeed/picoclaw/pull/3143) 和类型断言修复 [#3092](https://github.com/sipeed/picoclaw/pull/3092)，并对 [#3088](https://github.com/sipeed/picoclaw/issues/3088) 给出明确的路线图回应。

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

好的，这是为您生成的 NanoClaw 项目动态日报。

---

# NanoClaw 项目日报 | 2026-06-18

## 1. 今日速览

过去 24 小时，NanoClaw 项目呈现高度活跃状态，社区贡献和核心团队响应均非常积极。共处理了 15 个 Pull Request，其中包含多个重要的安全修复和功能增强。值得注意的是，项目同时发布了两个新版本（v2.1.0 和 v2.1.17），均包含破坏性变更，需要用户注意。此外，一个关于会话故障导致全局消息投递停滞的严重问题（#2796）已得到修复，但一个关于群组成员权限管理的安全漏洞（#2807）被新报告出来，正在等待处理。

## 2. 版本发布

项目在过去 24 小时内发布了两个新版本，均包含破坏性变更，升级时需特别注意。

-   **v2.1.17**: [发布链接](https://github.com/nanocoai/nanoclaw/releases/tag/v2.1.17)
    -   **更新内容**：这是一个从 v2.1.1 到 v2.1.17 的滚装发布，包含了此间所有对 `package.json` 的更新。
    -   **破坏性变更**：
        -   **`@onecli-sh/sdk` 版本要求**：SDK 版本从 0.5.0 跳跃至 2.2.1，现在强制要求 OneCLI 服务器支持 `/v1` API。旧版服务器的 SDK 调用将全部返回 404 错误。
        -   **迁移注意事项**：所有使用 OneCLI 的用户必须在升级前确保其所连接的 OneCLI 网关（gateway）和 CLI 版本已更新至与 v2.1.17 兼容的版本。

-   **v2.1.0**: [发布链接](https://github.com/nanocoai/nanoclaw/releases/tag/v2.1.0)
    -   **更新内容**：这是从 v2.0.65 到 v2.1.0 的滚装发布。
    -   **破坏性变更**：
        -   **启动前提条件变更**：现在主机（host）在启动时必须检测到 `data/upgrade-state.json` 文件，该文件需记录该次安装已通过合规渠道升级至当前版本，否则主机将拒绝启动。
        -   **迁移注意事项**：所有部署在升级到 v2.1.0 时，除了代码更新，还需要手动创建或更新 `data/upgrade-state.json` 文件。具体的文件格式和部署脚本请参考官方升级指南。

## 3. 项目进展

今日有 4 个 PR 被合并或关闭，项目在稳定性和国际化方面取得进展。

-   **清理废弃代码**：`#2803` [已合并] 移除了在 v2 架构中已无生产调用者的 `resolveGroupIpcPath` 函数，清理了代码库。
-   **增强社区贡献**：`#2806` [已合并] 添加了韩文版 README，并更新了语言切换器，降低了韩国开发者的参与门槛。
-   **修复关键稳定性问题**：`#2797` [已合并] 修复了 `#2796` 报告的问题，实现了单个会话失败不影响其他会话的消息投递，显著提升了系统整体的健壮性。
-   **修复安装流程**：`#2805` [已合并] 修复了 `claude setup-token` 命令在被 PTY 包装时，输出内容可能被截断导致 token 解析失败的问题，完善了初始化流程。

## 4. 社区热点

-   **话题热度：会话隔离与安全**：被关闭的 Issue `#2796` (一个异常会话导致所有消息投递卡死) 和其修复 PR `#2797` 引发了关于系统容错能力的讨论。
-   **功能提议：Apple Container 原生支持**：PR `#2809` 添加了 macOS 的 Apple Container 运行时支持，并支持指向远程 OneCLI 网关。这表明社区对非 Docker 的容器化运行环境有明确需求。
-   **安全漏洞报告**：新开放的 Issue `#2807` 报告了一个严重的安全漏洞：在所有者初始化的群组中，非所有者成员可以未经审批创建持久子代理。该议题虽无评论，但其“安全”标签使其成为社区关注的焦点。

## 5. Bug 与稳定性

今日报告的 Bug 集中在代码健壮性和安全防护方面，大部分已有对应的修复 PR。

-   **严重 - 权限管理漏洞**：`#2807` [开放] 报告了群组中非所有者成员可绕过审批创建子代理的安全漏洞。**目前无对应的 fix PR**。
-   **严重 - 路径遍历漏洞**：`#2799` [开放](#) 和 `#2800` [开放](#) 分别修复了 `send_file` 功能和 `ncl groups create` 命令中的路径遍历漏洞，防止恶意操作读取或写入容器内任意文件。
-   **高 - 功能阻塞**：`#2804` [开放] 修复了 `ncl messaging-groups create` CLI 命令因数据库约束问题始终报错，完全无法使用的问题。
-   **高 - 请求超时**：`#2802` [开放] 修复了 `ncl` 客户端在 socket 通信中无请求超时和响应大小限制的缺陷，防止因服务端异常导致客户端永久挂起。
-   **中 - 数据类型错误**：`#2801` [开放] 修复了 `safeParseContent` 函数在处理非对象 JSON（如数字、字符串）时返回类型错误，导致后续逻辑无法正确处理的 bug。
-   **中 - 数据一致性问题**：`#2808` [开放] 修复了 `insertMessage` 操作在重复 ID 下不幂等，以及审批行缺少 `agent_group_id` 的两个 bug。

## 6. 功能请求与路线图信号

-   **细粒度权限控制**：PR `#2793` 为“代理间”消息引入了可选的“每消息审批策略”。这回应了企业级用户对 Agent 间自动化协作进行管控的需求，很可能被纳入下一个次要版本。
-   **非 Docker 运行时支持**：PR `#2809` 对 Apple Container 的支持，表明路线图可能正在考虑拥抱更多样的容器化方案，以覆盖 macOS 用户群体。
-   **可观测性工具**：PR `#2795` 添加了一个名为 `/add-clidash` 的只读仪表盘技能，这是一个基于 CLI 派生的实用工具，体现了社区对便捷运维工具的需求。

## 7. 用户反馈摘要

-   **负面反馈：单体故障问题**：`#2796` 中的用户 `mashkovtsevlx` 提供了一个典型的负面场景——一个会话的健康问题（无法读取 `outbound.db`）就能导致整个实例的消息投递服务瘫痪，直到重启守护进程。这暴露了系统在应对局部故障时缺乏隔离机制。该问题已被修复。
-   **安全问题担忧**：虽然 `#2807` 报告者 `YLChen-007` 未发表评论，但报告的内容本身反映了用户对群组内成员权限清晰划分和管理的担忧，尤其是在多租户或协作场景下。

## 8. 待处理积压

-   **PR #2750：容器异常退出导致 `outbound.db` 日志资源锁定** [链接](https://github.com/nanocoai/nanoclaw/pull/2750)
    -   **状态**：由 `sturdy4days` 发起，已开放 6 天，无新评论。
    -   **详情**：该 PR 旨在修复两个紧密相关的数据库日志恢复问题。这是影响生产环境稳定性的重要问题，建议维护者尽快审核并推进合并。

-   **PR #2804：CLI 命令创建群组功能完全不可用** [链接](https://github.com/nanocoai/nanoclaw/pull/2804)
    -   **状态**：由 `sturdy4days` 发起，已开放 1 天。
    -   **详情**：这是一个破坏性 Bug，导致 `ncl messaging-groups create` 命令无效。对于依赖 CLI 进行管理的用户影响很大，建议优先处理。

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

好的，这是为您生成的 NullClaw 项目动态日报。

---

# NullClaw 项目动态日报 (2026-06-18)

## 今日速览

今日项目活跃度中等，主要以修复和文档更新为主。过去24小时内没有新版本发布，但有3个活跃的Issue和3个待合并的Pull Request，表明社区和开发者仍在积极互动。**核心动态集中在两个方向**：一是对CLI交互体验的修复（`#960`），旨在解决长期困扰用户的终端控制字符显示问题；二是对核心功能（调度器Scheduler）和配置文档（Anthropic Provider、Memory系统）的改进与补充。整体来看，项目正处于解决历史遗留问题并提升可配置性的阶段。

## 项目进展

由于今日没有PR被合并或关闭，项目整体状态没有向前推进。然而，`#960`、`#961`和`#962`三个PR均已提交，预示着以下改进即将落地：
- **CLI交互优化**：`#960` 提供了针对`nullclaw agent` REPL的箭头键支持，解决了`#865`报告中提到的核心痛点。
- **记忆系统增强**：`#961` 引入了三个新的记忆配置项，允许用户更精细地控制记忆召回行为，提升了系统的灵活性和可定制性。
- **文档完善**：`#962` 补充了原生Anthropic提供商的详细配置文档，有助于扩大用户群体并降低使用门槛。

## 社区热点

今日最受关注的讨论集中在两个长期未解决的Bug上：
- **`#915` [调度器授权问题]**：此Issue已存在一个多月，今日仍有新评论。用户报告在Telegram和CLI中都无法使用调度器，仅得到“unauthorized”响应。尽管模型和工具调用正常，但调度器作为核心自动化功能受阻，影响了用户的关键使用场景。 [查看详情](https://github.com/nullclaw/nullclaw/issues/915)
- **`#865` [CLI控制字符问题]**：此Bug对日常使用体验影响极大。用户反馈CLI无法正确处理方向键，导致命令行编辑效率低下。今日提交的`#960` PR直接针对此问题提供了修复方案，因此该Issue的关注度也随之上升。 [查看详情](https://github.com/nullclaw/nullclaw/issues/865)

## Bug 与稳定性

今日没有新的Bug报告，但两个待解决的长期Bug是稳定性关注的重点：
1.  **[中度/功能Bug] 调度器返回“unauthorized” (`#915`)**：严重影响Scheduler功能的使用。此问题已存在超过一个月，目前尚无明确的修复PR，可能需要维护者优先排查权限验证或插件兼容性问题。 [查看详情](https://github.com/nullclaw/nullclaw/issues/915)
2.  **[低度/体验Bug] CLI显示控制字符 (`#865`)**：虽然不导致程序崩溃，但严重影响了命令行交互体验。好消息是，PR `#960` 已经提交了修复方案，预计将在后续合并中解决此问题。 [查看详情](https://github.com/nullclaw/nullclaw/issues/865)

## 功能请求与路线图信号

今日无直接的功能请求Issue，但从新提交的PR中可以观察到社区的演进方向：
- **可配置的记忆系统**：PR `#961` 增加了`auto_recall`、`recall_limit`和`max_context_bytes`配置项，这表明社区和贡献者正在关注如何让AI Agent的记忆行为更可控、更高效，这很可能成为下一版本的重要特性。
- **原生提供商支持**：PR `#962` 完善了对Anthropic原生API和OAuth的文档支持，表明项目正在积极拓展对主流模型提供商的原生集成，减少对第三方代理的依赖。

## 用户反馈摘要

从今日活跃的Issue评论中，可以提炼出以下用户反馈：
- **痛点**：用户对**基础功能故障的容忍度较低**。例如`#915`中的调度器问题，用户已搭建了完整的硬件和软件环境，但核心功能无法使用，这直接影响了其部署和信任。这类问题需要快速响应。
- **使用场景**：用户案例`#915`展示了一个典型的**私有化部署**场景：在Ubuntu上使用外部Ollama托管模型和RTX 3090显卡。这表明NullClaw的核心用户群体是具有一定技术能力、追求本地化运行和API灵活性的开发者和爱好者。
- **不满**：用户`#861`的诉求（如何配置Web UI）反映出**现有文档对新手不够友好**。用户明确要求“非术语的人话”来解释如何部署，说明文档的易读性和步骤清晰度还有提升空间。

## 待处理积压

- **`#861` [Web UI配置求助]**：此Issue已存在近两个月，是关于如何配置Web UI的求助帖，目前只有1条评论。作为社区用户普遍关心的功能，该问题长期无人响应，可能会对新用户的留存率产生负面影响。建议维护者在该Issue下提供更清晰的配置指南，或考虑优化README中的相关章节。 [查看详情](https://github.com/nullclaw/nullclaw/issues/861)
- **`#865` & `#915`**：这两个Bug分别影响CLI和Scheduler，属于核心功能。其中`#865`已有修复PR，但`#915`仍待跟进，建议将其标记为高优先级。 [查看详情](https://github.com/nullclaw/nullclaw/issues/865) | [查看详情](https://github.com/nullclaw/nullclaw/issues/915)

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是根据您提供的 IronClaw 项目 GitHub 数据生成的 2026-06-18 项目动态日报。

---

# IronClaw 项目动态日报 | 2026 年 06 月 18 日

## 1. 今日速览

今日项目整体活跃度 **高**，社区与核心开发团队均保持强劲的贡献势头。PR 与 Issue 的合入与关闭数量显著，展现了高效的开发闭环。核心开发聚焦于 **“Reborn”版本** 的稳定性、WebUI 的 **用户体验一致性**、**OAuth 认证流程** 的健壮性，以及 **CI/CD 基础设施** 的现代化改造。此外，自动化的 **工具权限管理** 和 **审批流程** 正成为新的开发热点。尽管发布节奏放缓（今日无新版本），但项目内部的迭代速度和质量控制投入正在加强。

## 2. 版本发布
*（今日无新版本发布）*

## 3. 项目进展

今日项目在多个关键领域取得了实质性进展，尤其是 **“Reborn”版本的核心修复** 和 **工程效率提升**。

- **WebUI 稳定性与 Bug 修复**：
    - **修复了工具调用失败在 WebUI 中无法即时显示的问题 (#4984)**：确保 `builtin.http` 等工具失败时，用户能立即在界面上看到失败的 Tool Card，而无需手动刷新页面。这对于用户反馈和诊断工具错误至关重要。
    - **修复了技能 (Skills) 验证错误不清除的问题 (#5047)**：用户在添加技能时，填写完必填字段后，之前的验证错误消息会立即消失，提升了表单交互的流畅性。
    - **改进日志与文档图标 (#4923)**：将导航栏中容易误解的图标替换为清晰易懂的文字标签，降低了新用户的学习成本。

- **工程效率与 CI/CD 改进**：
    - **引入统一的 CI 判定检查 (#5075)**：创建了 `ci-verdict` 工作流，作为合入前的唯一绿灯信号，简化了审查者和合并队列的判断逻辑。
    - **修复了夜间全量 E2E 测试失败 (#5073)**：针对已删除的场景文件导致 `v2-engine` 矩阵测试失败的问题进行了修复，并增加了护栏，防止类似问题再次发生。
    - **优化了 Rust 缓存策略 (#5074)**：确保合并队列中的代码能够正确地被缓存，从而显著加快后续 PR 和合并操作的编译速度。

- **OAuth 与认证流程**：
    - **修复了 OAuth 认证回调时的权限刷新问题 (#5053)**：移除了 OAuth 凭据的进程级一次性缓存，确保 Google 等令牌在过期后能正确刷新，而不是被无限复用。

**总结**：今日项目在 **“修复”** 和 **“基建”** 两条线上齐头并进。三笔关键的 Bug 修复直接提升了 Reborn 版本的用户体验，而 CI/CD 的重构则将为整个团队未来的开发效率带来长期红利。

## 4. 社区热点

今日社区讨论热度分散，但反映出对 **“Reborn”版本深度使用** 和 **用户体验细节** 的高度关注。

- **Local Dogfooding 发现 (#4879)**：[链接](https://github.com/nearai/ironclaw/issues/4879)
    - **诉求**：社区成员（think-in-universe）正在进行深度的本地“吃狗粮”测试，集中反馈 Reborn 启动、配置、模型提供商设置等 “First-Run” 体验问题。这暗示了团队需要一个更平滑、更友好的开箱即用体验。
- **Google OAuth 令牌主动刷新 (#5071)**：[链接](https://github.com/nearai/ironclaw/issues/5071)
    - **诉求**：这是一个高风险的 Bug。用户明确指出，Google OAuth 的短生命周期（~1小时）要求用户频繁重新认证，严重破坏了使用 GSuite 工具的用户体验。用户期望的是后台自动刷新，而非手动干预，这直接关系到 Reborn 在长周期自动化任务中的可用性。
- **认证流程重放问题 (#5070)**：[链接](https://github.com/nearai/ironclaw/issues/5070)
    - **诉求**：该 Issue 描述了在审批流后取消 OAuth 认证，会导致工具活动“幽灵”运行或重复弹出认证提示。这反映了认证流程在状态管理上的缺陷，可能对用户造成困惑和操作死锁。

**分析**：社区热点已经从“能否用”（功能性）转向了“用得好”（流程与易用性）。OAuth 流程的处理成为核心痛点，表明 Reborn 在集成第三方服务时的认证持久性和流程完整性有待加强。

## 5. Bug 与稳定性

今日报告的 Bug 主要集中在 **“Reborn”** 分支的高风险区域。

| 严重程度 | Bug 描述 | Issue 链接 | 是否有 Fix PR |
| :--- | :--- | :--- | :--- |
| **高风险** | **Google OAuth 令牌过期后需重复认证**，影响 GSuite 工具使用。 | [#5071](https://github.com/nearai/ironclaw/issues/5071) | 有，[#5054](https://github.com/nearai/ironclaw/pull/5054) 与 [#5053](https://github.com/nearai/ironclaw/pull/5053) 已合并/提出解决策略。 |
| **中风险** | **无效的聊天 URL 会导致错误页面而非重定向**，用户体验差。 | [#5077](https://github.com/nearai/ironclaw/issues/5077) | 无 |
| **中风险** | **侧边栏在非聊天页面仍高亮显示聊天线程**，导航逻辑混乱。 | [#5076](https://github.com/nearai/ironclaw/issues/5076) | 无 |
| **中风险** | **认证流程取消后，工具活动状态未正确清理**，可能导致界面假死。 | [#5070](https://github.com/nearai/ironclaw/issues/5070) | 无 |
| **低风险** | **OAuth 认证卡片在授权 URL 不可用时未正确显示**，信息不透明。 | [#5066](https://github.com/nearai/ironclaw/issues/5066) | 有，[#5067](https://github.com/nearai/ironclaw/pull/5067) 已提出。 |

此外，还存在一个**长期未解决的高危信号**：**夜间 E2E 测试持续失败**（[#4108](https://github.com/nearai/ironclaw/issues/4108)），虽然已有修复 PR（[#5073](https://github.com/nearai/ironclaw/pull/5073)），但其长期存在表明测试基础设施需要更强的鲁棒性监督。

## 6. 功能请求与路线图信号

今日的功能请求显示出用户和开发者正积极为 Reborn 版本的 **自动化能力** 和 **精细化管理** 铺路。

- **自动化 UX 重新设计 (#5069)**：[链接](https://github.com/nearai/ironclaw/issues/5069)
    - **信号**：纯粹的功能性 Issue，标题直接指向“重新设计”。这表明当前 Reborn 的自动化（Scheduled Tasks）页面 UX 可能不足以支撑其复杂功能，预示着团队计划对其进行一次大的 UI 改造。
- **工具权限与全局自动审批设置 (#4960 系列 PR)**：[#5068](https://github.com/nearai/ironclaw/pull/5068)，[#5063](https://github.com/nearai/ironclaw/pull/5063)，[#5062](https://github.com/nearai/ironclaw/pull/5062)
    - **信号**：这是一个非常强烈的路线图信号。今日有三笔“XL”尺寸的 PR 集中合入，旨在建立一个端到端的 **工具权限管理系统**。包括“单个工具权限覆盖”（允许/拒绝/每次询问）和“全局自动审批开启”设置。这标志着 IronClaw 正在从“所有工具调用都需要交互确认”的模式，迈向 **“可配置、可信任、可自动化”** 的智能体权限管理新模式，是向更高阶自动化演进的关键一步。

## 7. 用户反馈摘要

从今日的 Issues 中，可以提炼出几类核心用户痛点：

- **挫折感：工具失败的反馈延迟与信息不透明**
    - 用户抱怨 `builtin.http` 工具失败后，错误信息仅为 `invalid_input`（[#4704](https://github.com/nearai/ironclaw/issues/4704)），没有任何具体细节，且会陷入循环。此外，失败的 Tool Card 需要刷新页面才能看到（[#4942](https://github.com/nearai/ironclaw/issues/4942)）也增加了用户的沮丧感。**用户渴望即时、清晰、可操作的错误信息。**
- **中断感：频繁的认证打断工作流**
    - 对于 Google Calendar、Drive 等 GSuite 工具，Google 令牌的短有效期导致用户需要反复进行 OAuth 确认（[#5071](https://github.com/nearai/ironclaw/issues/5071)）。这表明在长时间任务中，频繁的身份请求会严重割裂用户的使用体验。**用户期望“一次认证，持续可用”。**
- **困惑感：WebUI 交互的语义不一致**
    - 侧边栏高亮的误导（[#5076](https://github.com/nearai/ironclaw/issues/5076)）、无效 URL 的错误页面（[#5077](https://github.com/nearai/ironclaw/issues/5077)）以及图标含义不清（[#4923](https://github.com/nearai/ironclaw/issues/4923)），都让新用户感到困惑。**用户对 UI 的“心理模型”要求一致且直观。**

## 8. 待处理积压

以下列出了一些值得维护者关注的长期未解决或重要性较高的议题：

- **夜间 E2E 测试持续失败** ([#4108](https://github.com/nearai/ironclaw/issues/4108))：该 Issue 已存在超过三周，虽今日有修复 PR（[#5073](https://github.com/nearai/ironclaw/pull/5073)），但根本原因及其对主分支健康度的威胁需要持续关注。这是一个项目稳定性的**红旗信号**。
- **企微群组对话标题过于模糊** ([#4505](https://github.com/nearai/ironclaw/issues/4505))：创建于 6 月 5 日，至今无更新。对于使用 WeCom 的用户来说，群聊无法区分是个实际痛点，可能影响其 Reborn 版本的采用。
- **Qwen 模型 Provider 错误** ([#1520](https://github.com/nearai/ironclaw/issues/1520))：创建于 3 月，长期未分配。虽然可能是特定 Provider 的兼容性问题，但对于依赖 Qwen 的用户来说是一个阻塞性 Bug，建议根据当前优先级评估是否处理。
- **依赖更新 PR 等待合并** ([#4002](https://github.com/nearai/ironclaw/pull/4002))：一次涉及 16 个 GitHub Actions 依赖的批量更新 PR，因风险较高（`risk: medium`）已开放近一个月。虽然安全性和稳定性风险可控，但建议尽快安排审查和合并，以避免因版本落后导致的潜在兼容性问题。

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

好的，作为 LobsterAI 开源项目分析师，我已根据您提供的 GitHub 数据，为您生成 2026-06-18 的项目动态日报。

---

### LobsterAI 项目动态日报 | 2026年6月18日

---

### 1. 今日速览

今日项目活跃度**极高**。核心团队正在进行一轮重要的版本迭代收尾工作，24小时内合并了14个Pull Request，并发布了新版本。社区方面，有一个关于“Artifact自动加载导致本地文件读取”的安全漏洞报告 (#2176) 成为今日焦点，需要维护者高度关注。整体来看，项目开发节奏紧凑，功能迭代与安全维护并行推进。

### 2. 版本发布

项目于昨日（2026-06-15）发布了 **LobsterAI 2026.6.15** 版本。

- **主要更新内容**:
    - **新增 “Computer Use” MVP 功能**: 该功能允许AI模型通过内置的“技能包”直接操控Windows桌面应用程序，例如列出窗口、启动应用、模拟鼠标键盘操作等。这是一个重大的能力拓展，标志着LobsterAI向Agent形态迈进的重要一步。
    - **Cowork 协作模式增强**: 新增了**实时ASR（自动语音识别）输入**功能，用户现在可以在协作编辑时通过语音输入，并由AI实时转写为文本，极大提升了输入效率。
    - **上下文理解优化**: 改进了协作编辑中的上下文压缩（post-compaction）逻辑，使模型能更准确地理解长篇对话的要点。
- **破坏性变更**: 本次发布中，语音输入模块进行了彻底的**重构**。原有的短语音上传识别（`asr:recognize` IPC接口）被**完全移除**，并统一为“实时ASR”模式。相关的设置项 (`voiceInput.recognitionMode`) 也被删除。
- **迁移注意事项**: 
    - 使用自定义脚本或外部工具调用旧版ASR接口的开发者需要特别注意，需迁移到新的基于WebSocket的实时ASR流程。
    - 用户界面上，原“听写”相关的文字描述已全部更新为“语音输入”。

### 3. 项目进展

今日项目取得了显著进展，多项核心功能得到推进和修复。

- **Computer Use 功能完善**: PR [#2156](https://github.com/netease-youdao/LobsterAI/pull/2156) 将 Computer Use 运行时 (runtime) 升级至 1.0.7 版本，以解决稳定性问题。
- **语音输入功能终版定型**: 通过一系列PR（[#2160](https://github.com/netease-youdao/LobsterAI/pull/2160) 和 [#2163](https://github.com/netease-youdao/LobsterAI/pull/2163)），团队彻底移除了旧版ASR流程，并优化了实时语音输入的录音界面和配额处理逻辑，标志着该功能已趋于成熟稳定。
- **Artifact 分享能力扩展**: PR [#2178](https://github.com/netease-youdao/LobsterAI/pull/2178) 为 Artifact 面板新增了对 **Markdown (.md)** 和 **Mermaid** 文件的支持，并可进行打包分享，进一步提升了功能的通用性。
- **代码合并与发布准备**: PR [#2179](https://github.com/netease-youdao/LobsterAI/pull/2179) 将早前发布的 2026.6.11 版本分支合并回主分支，为未来版本整合了完整的DOCX、PPTX、XLSX、PDF等文档分享预览能力。

### 4. 社区热点

今日社区讨论最集中的是 **Issue #2176: [Security] LobsterAI automatic artifact loading allows message-derived arbitrary local file reads**。
- **链接**: [Issue #2176](https://github.com/netease-youdao/LobsterAI/issues/2176)
- **分析**: 这是一个被标记为`[Security]`的高风险漏洞。报告者发现，LobsterAI在自动解析“助手”或“工具”输出中的媒体文件引用 (`MEDIA:`）时，存在缺陷，可能允许攻击者通过精心构造的对话消息，诱导AI读取用户本地任意文件。该问题在创建后迅速获得了一个评论，社区关注度极高，凸显了用户对AI Agent安全性的担忧。
- **诉求**: 社区期待维护团队能迅速确认漏洞，给出修复计划时间表，并在修复前提供临时缓解措施。

### 5. Bug 与稳定性

- **【严重】安全问题: 本地文件读取漏洞**
    - **报告**: [Issue #2176](https://github.com/netease-youdao/LobsterAI/issues/2176)
    - **状态**: 已报告，尚未有关联的修复PR。此问题为高优先级，需要立即评估和修复。

- **【低】UI显示问题: MCP服务名称过长**
    - **报告**: [Issue #1422](https://github.com/netease-youdao/LobsterAI/issues/1422)
    - **状态**: 旧Issue，于今日被重新激活（更新）。问题描述为MCP自定义页面中，当服务名称较长时，删除弹窗的显示效果不佳。此问题对功能无影响，主要影响用户体验。

### 6. 功能请求与路线图信号

- **“Computer Use” 功能**: 作为MVP刚刚发布，社区必然会围绕其能力边界、安全性、支持的操作系统（目前仅为Windows x64）和可用性提出大量反馈和新需求。未来版本很可能聚焦于扩展这一能力，例如支持macOS、增加更多内置操作等。
- **语音输入体验优化**: 随着实时ASR的正式落地，可以预见用户会提出关于多语言支持、自定义唤醒词、语速适配、离线模式等更精细化的需求。
- **文档分享能力增强**: PR [#2178](https://github.com/netease-youdao/LobsterAI/pull/2178) 的合并表明团队正在积极扩展Artifact的分享能力。未来很可能支持更多文件格式，或提供在线预览/编辑功能。

### 7. 用户反馈摘要

- **（安全隐患）** Issue [#2176](https://github.com/netease-youdao/LobsterAI/issues/2176) 的评论中，用户表达了对“AI助手可能被用于恶意读取文件”的担忧，这反映了高级用户对模型输出内容安全性的高度警惕。
- **（体验建议）** Issue [#1422](https://github.com/netease-youdao/LobsterAI/issues/1422) 虽然评论不多，但其作为“stale”标签的Issue被重新激活，说明用户界面细节的友好性仍是社区长期关注的痛点。

### 8. 待处理积压

- **PR #1277: Electron 依赖升级**
    - **链接**: [PR #1277](https://github.com/netease-youdao/LobsterAI/pull/1277)
    - **状态**: 由Dependabot自动创建的依赖包升级PR，提议将Electron从40.2.1升级到42.4.0。
    - **建议**: 此PR已存在超过2个月且未合并。大型框架升级通常涉及重大变更和回归测试，但长期滞后可能导致安全风险累积。建议维护者评估并将其纳入下一阶段的维护计划。

- **Issue #1422: MCP自定义页面显示问题**
    - **链接**: [Issue #1422](https://github.com/netease-youdao/LobsterAI/issues/1422)
    - **状态**: 这是一个长期悬而未决的UI小问题。虽然不紧急，但作为体验类反馈，长期不解决会造成社区对维护者响应速度的负面印象。建议在下一个UI迭代周期内予以修复。

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

# TinyClaw 项目动态日报｜2026-06-18

## 今日速览
- **活跃度评估**：低（仅 3 条新 Issue，无 PR 活动，无版本发布）。项目今日未见代码提交或合并活动，但集中暴露了 **3 个严重安全漏洞**，涉及未认证 API 入口和任意文件泄露，项目健康度面临重大威胁。
- 所有新 Issue 均由同一安全研究员提交，均为 **Security Advisory** 级别，且截至日报生成时尚未得到维护者响应。
- 社区讨论和贡献度几乎为零（无评论、无点赞），但漏洞的严重性应引起核心团队高度重视。

## 版本发布
无（过去 24 小时无新版本发布）。

## 项目进展
今日无合并或关闭的 PR，项目在功能演进上无明显推进。需优先处理安全漏洞。

## 社区热点
今日唯一活跃的议题是 **3 个安全相关 Issue**，均无评论但内容极为敏感，以下是按严重程度排列：

- **#284 – [Security] TinyAGI allows unauthenticated API messages to invoke Claude with provider permission checks disabled by default**  
  链接：[https://github.com/TinyAGI/tinyagi/issues/284](https://github.com/TinyAGI/tinyagi/issues/284)  
  核心问题：`POST /api/message` 入口无身份验证，且默认未启用 provider 权限检查，攻击者可任意调用 Claude 模型，导致资源滥用和数据泄露。

- **#283 – [Security] Unauthenticated `prompt_file` agent configuration allows arbitrary local file disclosure to the model provider**  
  链接：[https://github.com/TinyAGI/tinyagi/issues/283](https://github.com/TinyAGI/tinyagi/issues/283)  
  核心问题：代理配置 API 未认证，攻击者可指定 `prompt_file` 路径，服务端读取本地任意文件并发送至模型提供商，造成敏感信息外泄。

- **#282 – [Security] Untrusted `[send_file: ...]` response tags allow arbitrary host file attachment delivery in TinyAGI**  
  链接：[https://github.com/TinyAGI/tinyagi/issues/282](https://github.com/TinyAGI/tinyagi/issues/282)  
  核心问题：`[send_file: ...]` 响应标签缺乏安全校验，攻击者可诱导服务端输出恶意标签，导致任意主机文件被附送到其控制的端点。

这些 Issue 背后的诉求非常明确：**紧急修复安全边界，阻止未授权访问和文件泄露**。社区期待维护者立即回应并给出修复时间表。

## Bug 与稳定性
今日报告了 **3 个安全漏洞**，均属于 **严重/高危** 级别，尚未有对应的修复 PR：

| Issue | 严重程度 | 描述 | 状态 |
|-------|----------|------|------|
| #284 | 严重 | 未认证 API 可调用 Claude，绕过了 provider 权限检查 | 待修复 |
| #283 | 严重 | 未认证配置接口可导致本地文件泄露到模型提供商 | 待修复 |
| #282 | 严重 | send_file 标签允许任意主机文件外发 | 待修复 |

暂无其他稳定性回归或崩溃报告。

## 功能请求与路线图信号
今日无用户提出的新功能需求。项目应优先将 **安全加固** 作为下一个版本的核心方向，尤其建议：
- 为所有 API 端点增加身份认证机制
- 对 `prompt_file` 和 `send_file` 等文件操作路径进行白名单限制
- 在默认配置中启用 provider 权限检查

## 用户反馈摘要
由于三个 Issue 均无评论和点赞，无法提炼真实用户痛点。但漏洞报告本身隐含了用户（安全研究员）的反馈：**项目现有的安全设计存在严重缺失，缺乏身份认证和输入校验**，这可能影响实际部署用户的信任度。

## 待处理积压
目前所有 3 个 Issue 均为 **当天创建、待响应** 状态，无长期未处理的积压。但维护者应 **尽快回复** 并提供预计修复时间，避免漏洞信息扩散后造成负面影响。

---
**总结**：项目今日处于 **安全危机** 状态，核心功能开发停滞。建议立即组织安全审计，安排至少一个补丁版本（如 v0.x.x-security）修复上述三个漏洞，并在发布后更新文档。

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis 项目日报 | 2026-06-18

---

## 1. 今日速览  
过去24小时内，Moltis 项目保持平稳的社区活跃度：收到 **2 个新功能请求（Issue）** 和 **1 个待合并的 Pull Request**，无新版本发布，无重大 Bug 报告。社区关注点集中在 **TTS 输出格式可配置性** 与 **Markdown 复制/导出功能** 上，而 PR #1130 针对 WebUI RPC 超时时间可配置的改进已提交，整体项目处于持续打磨用户体验的阶段。

---

## 2. 版本发布  
无新版本发布。

---

## 3. 项目进展  
- **PR #1130** [OPEN] `feat: make webui rpc timeout configurable`  
  - 作者：khimaros  
  - 该 PR 允许用户配置 WebUI 与后端之间的 RPC 超时时间，解决了此前超时硬编码导致的灵活性不足问题。  
  - 关联 Issue：#1127（未出现在今日活跃列表中，但说明该改进对应了具体的用户反馈）。  
  - 当前状态：待合并，尚未收到审阅评论。  

**项目整体推进**：虽然暂无已合并的 PR，但此改进代表项目在**配置化与可定制性**方向上迈出了明确一步，符合社区近期提出的多个“让行为可配置”的需求趋势。

---

## 4. 社区热点  
- **Issue #1126** [OPEN] `[Feature]: allow to configure the format of tts output`  
  - 👤 作者：khimaros | 💬 评论数：3 | 👍 0  
  - 链接：[#1126](https://github.com/moltis-org/moltis/issues/1126)  
  - **分析**：这是今日最活跃的 Issue（3条评论），用户希望 TTS（文本转语音）输出格式能够被自定义（例如选择 wav、mp3、opus 等编码或采样率）。背后的诉求是**增强输出兼容性**，使 Moltis 能更灵活地集成到不同的下游应用或设备中。评论详情未提供，但从标签来看，该功能已被社区标记为 enhancement，后续可能会与 PR #1130 的配置化思路合并考虑。

- **Issue #1131** [OPEN] `[Feature]: Add copy + export as Markdown`  
  - 👤 作者：vvuk | 💬 评论数：0 | 👍 0  
  - 链接：[#1131](https://github.com/moltis-org/moltis/issues/1131)  
  - **分析**：用户希望为对话内容增加“复制为 Markdown”和“导出为 Markdown”功能。这符合知识工作者及开发者使用场景——将 AI 对话结果直接粘贴到笔记或文档中。目前无评论，属于相对明确的产品需求。

---

## 5. Bug 与稳定性  
今日无新提交的 Bug、崩溃或回归问题。项目稳定性表现良好，未见严重缺陷报告。

---

## 6. 功能请求与路线图信号  
| Issue / PR | 功能 | 是否可能纳入下一版本 |
|------------|------|---------------------|
| #1126 | TTS 输出格式可配置 | ⭐ 高可能性。与 PR #1130 的“配置化”主题一致，且作者相同，可能作为同一批改进推出。 |
| #1131 | Markdown 复制/导出 | ⭐ 中等偏高，属于常见的 UI 交互增强，实现成本较低（基于现有渲染机制）。 |

此外，PR #1130 自身也是功能改进（RPC 超时可配置），对应 Issue #1127，预计合并后将为后续类似配置项（如 #1126）提供模式参考。

---

## 7. 用户反馈摘要  
- **TTS 格式需求**（#1126 评论）：用户希望能在不同输出格式间切换，推测当前输出格式固定（如仅 wav）导致在某些场景下文件过大或不兼容。  
- **Markdown 导出**（#1131）：用户期望便捷地将 AI 回复以结构化文本形式保存，常见于撰写报告、整理知识库、或与协作文档工具配合使用。  
- **RPC 超时配置**（#1130 的 Issue #1127）：有用户报告 WebUI 在长时间等待后端处理时发生超时断开，配置化可让用户根据网络状况或任务复杂度调整等待阈值。

整体来看，用户对 Moltis 的诉求集中在**可配置性、输出兼容性和实用导出功能**上，项目维护者对这些反馈响应较为积极（已提交 PR #1130 对应 #1127）。

---

## 8. 待处理积压  
今日无长期未响应的重要 Issue 或 PR。  
- 此前已打开的 #1126 和 #1131 均为两天内创建，尚属新鲜状态，无需特别提醒。  
- PR #1130 已提交但尚未获得任何代码审查，建议维护者尽快分配 Reviewer，以推进功能合并。

---

**数据来源**：Moltis GitHub 仓库，更新至 2026-06-18 UTC。  
**分析师备注**：项目健康度良好，社区参与稳步上升，建议优先合并 PR #1130 并跟进 #1126 的讨论，以巩固“可配置性”这一关键产品优势。

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

好的，作为CoPaw开源项目的AI分析师，基于您提供的2026年6月18日的数据，我为您生成了以下项目动态日报。

---

# CoPaw 项目动态日报 | 2026-06-18

## 1. 今日速览

今日项目整体**高度活跃**，共产生33条Issue和31条PR，反映出社区极高的参与度和维护团队的快速响应能力。核心进展集中在**上下文压缩机制的稳定性与功能性修复**（多个相关Bug与PR）、**渠道兼容性（钉钉/Windows SSL）的修复**，以及**插件系统的能力扩展（卸载钩子、数据分析插件）**。尽管修复密度很高，但社区反馈中关于“上下文压缩导致进程冻结”的严重Bug（Issue #5218）仍处于开放状态，成为当前项目稳定性的突出风险点。

## 2. 版本发布

**发布版本**: `v1.1.12.post1`
- **更新内容**: 本次为补丁版本更新，主要包含两项修复：
    1.  **修复发布脚本**: 修正了发布脚本中的参数展开问题 `(fix(scripts): correct prerelease arguments expansion)`。
    2.  **修复内存模块**: 将ChromaDB的探针集合重命名为'probe-test'，以避免可能的命名冲突或混淆 `(fix(memory): rename ChromaDB probe collection to 'probe-test')`。
- **破坏性变更**: 无。
- **迁移注意事项**: 对于普通用户平滑升级。该版本主要面向开发者修复了CI/CD流程，并对内部存储逻辑进行了清理，无用户侧配置或数据变更。

## 3. 项目进展

今日项目在多条线路上取得了实质性进展，正稳步迈向更稳定、功能更丰富的方向。

- **核心上下文管理重构**: **PR #5309** (由 @qbc2016 提交并已合并) 是一个重要里程碑，它将自定义的 `LightContextManager` 替换为 **AgentScope 2.0 的原生上下文压缩机制**，此举将提升上下文处理的核心性能和可维护性。
- **渠道稳定性增强**: **PR #5291** (由 @wangfei010313 提交并已合并) 修复了通过 `uv` 安装时，钉钉渠道因SSL证书配置缺失导致的通信失败问题。**PR #5298** (已合并) 一并处理了Windows构建验证中的SSL错误。
- **用户体验改进**: **PR #5303** 和 **PR #5306** (均已合并) 修复了Web端对话窗口上下文占用比例显示不准确的问题，现能正确匹配当前模型的 `max_input_length`。**PR #5293** (由 @zhaozhuang521 提交并已合并) 将历史会话列表从弹出式抽屉改为**右侧常驻面板**，提升了对话切换的流畅性。
- **插件系统预演**: **PR #4900** (由 @wangfei010313 提交) 提出了将插件加载器与Agent启动解耦的改进，以解决在冻结环境（如PyInstaller/Tauri）下插件系统超时无法启动的问题。此PR仍在开放中，是未来插件生态发展的重要基础。
- **本地模型就绪性检查**: **PR #5305** (由 @zhijianma 提交并已合并) 增强了对本地模型提供者的就绪性检查，将有助于减少因模型未准备就绪而导致的运行时错误。

## 4. 社区热点

今日社区讨论的热点高度集中在**上下文压缩机制的可靠性与智能性**上，反映了用户对核心功能的深度依赖及对现有问题的担忧。

- **[Bug] #5218 - 子Agent触发上下文压缩时QwenPaw进程冻结无响应 (16条评论)**
    - **链接**: [agentscope-ai/QwenPaw Issue #5218](https://github.com/agentscope-ai/QwenPaw/issues/5218)
    - **诉求分析**: 该问题严重性最高，用户报告当子Agent触发上下文压缩时，整个进程会完全冻结，只能通过手动重启恢复。这暴露了当前压缩机制存在**死锁或资源竞争**的严重缺陷，是影响生产环境稳定性的关键风险点。社区对此的关注度极高，期待核心团队尽快给出根本解决方案。
- **[Bug] #5171 - 上下文压缩保留缺少按条数保留/排除人设文件，导致信息完全丢失，任务中断 (8条评论)**
    - **链接**: [agentscope-ai/QwenPaw Issue #5171](https://github.com/agentscope-ai/QwenPaw/issues/5171)
    - **诉求分析**: 该问题进一步细化了上下文压缩的痛点。当人设文件超过保留阈值时，压缩策略会直接将所有上下文归零，导致Agent“失忆”和任务中断。社区诉求不仅是修复Bug，更是希望**引入更智能的压缩策略**（如按条数保留或排除关键文件）。
- **[Enhancement] #5063 - 集成Headroom作为可选上下文压缩层，以减少60-95%的Token消耗 (7条评论)**
    - **链接**: [agentscope-ai/QwenPaw Issue #5063](https://github.com/agentscope-ai/QwenPaw/issues/5063)
    - **诉求分析**: 该功能请求与前面两个Bug形成鲜明对比。社区在经历当前压缩机制带来的痛苦时，也积极寻求更优的替代方案。@K1-lihongrong 提出的集成 **Headroom** 的方案，承诺能大幅减少Token消耗，反映了用户对**降低成本和提升效率**的强烈需求。**值得注意的是，此功能已有一个对应的PR #5244 正在审核中**。

## 5. Bug 与稳定性

今日报告的Bug数量较多，但大部分已被快速修复。以下按严重程度排列关键Bug：

| 严重程度 | Issue/PR | 问题描述 | 当前状态 | 关联修复 |
| :--- | :--- | :--- | :--- | :--- |
| **严重** | [#5218](https://github.com/agentscope-ai/QwenPaw/issues/5218) | 子Agent触发上下文压缩时进程冻结无响应 | **开放中** | 暂无已合入的修复，是当前最严重的稳定性问题。 |
| **高** | [#5171](https://github.com/agentscope-ai/QwenPaw/issues/5171) | 上下文压缩策略不当，导致信息完全丢失 | **开放中** | 暂无已合入的修复，与#5218共同指向上下文管理机制的问题。 |
| **中** | [#5262](https://github.com/agentscope-ai/QwenPaw/issues/5262) | 软件升级后，被禁用的内置技能（如docx）会自动恢复启用 | **开放中** | 这是一个反复出现的回归问题，用户满意度低。 |
| **中** | [#5313](https://github.com/agentscope-ai/QwenPaw/issues/5313) | MCP `streamable_http` 的 `Authorization` header丢失"Bearer"前缀 | **已关闭** | 这是一个配置转换错误，已被快速处理。 |
| **低** | [#5317](https://github.com/agentscope-ai/QwenPaw/issues/5317) | Windows Tauri环境下，找不到Python，导致自定义Skill无法运行Python脚本 | **开放中** | 特定于Windows桌面版的打包问题。 |

此外，今日有多个修复Bug的PR被合并，如修复**钉钉频道 SSL 问题 (#5291)**、**构建验证SSL错误 (#5298)** 以及 **MCP server 资源泄露问题 (#4849)**，表明项目在持续快速迭代修补。

## 6. 功能请求与路线图信号

今日新提出的功能和社区讨论反映了项目未来可能的发展方向：

- **上下文管理革新 (高优先级)**: **Issue #5063** (Headroom集成) 获得了7条评论，且已有**PR #5244** (由 @K1-lihongrong 提交) 尝试实现。结合今日合并的**PR #5309** (迁移到AgentScope 2.0原生压缩)，可以预见，下一代上下文管理将成为近期开发的重点。
- **多模态与视觉模型路由 (长期需求)**: **Issue #3940** (支持为图片输入使用独立的视觉模型) 今日仍有讨论，这是一个用户呼声极高的功能，可能在未来版本中作为独立功能或Agent策略被实现。
- **任务进度可视化 (新需求)**: **Issue #5318** (由 @MoeShinX 创建) 建议添加一个类似TodoWrite的原生进度面板，以显示多步骤Agent任务的执行进度。这反映了用户对**工作流透明度和可控性**的更高要求。
- **记忆搜索优化 (新需求)**: **Issue #5316** (由 @hellozhouuu 创建) 建议为 `memory_search` 增加**时间感知排序**功能，使近期记忆在搜索结果中拥有更高权重。这表明在长期记忆管理上，社区希望寻求更精细化的控制。

## 7. 用户反馈摘要

从今日的Issue评论中，我们可以清晰地看到用户的真实痛点：

- **“失忆”与“假死”的焦虑**: 用户 @MCQSJ 和 @malongan 的反馈 (Issue #5218, #5171) 是最典型的负面体验。当Agent的核心能力（记忆和上下文）出现问题时，用户会感到极大的无助和挫败，尤其是当任务因信息丢失或进程冻结而中断时。
- **重复劳动与回归问题**: 用户 @daigoopautoy 在 Issue #5262 中表达了“每次升级都要手动禁用一批技能”的烦恼。这种反复出现的**配置不持久化**问题，消耗了用户的信任和精力。
- **渠道兼容性的“水土不服”**: 用户 @ardorleo 指出通过 `uv` 安装的QwenPaw钉钉频道无法使用，而官方安装包却可以。这种**不同安装方式的体验差异**给用户带来了困惑，并增加了学习成本。幸运的是，这个问题已在当日被PR #5291修复。
- **对更优方案的渴望**: 用户 @K1-lihongrong 提出的 Headroom 集成建议 (Issue #5063)，以及在Issue #3940中持续讨论的视觉模型路由，都表明了用户不仅希望修复现有问题，更在**积极寻求更高效、更灵活的新解决方案**。

## 8. 待处理积压

以下为长期未解决的重要Issue或PR，建议维护团队关注：

- **[Feature] #3940 - 支持为图片输入使用独立的视觉模型 (创建: 2026-04-29)**
    - **链接**: [agentscope-ai/QwenPaw Issue #3940](https://github.com/agentscope-ai/QwenPaw/issues/3940)
    - **分析**: 这是一个持续近2个月的高价值功能请求。用户需要在非视觉模型对话中也能分析图片，避免手动切换模型。尽管讨论热度不高，但呼声稳定，是完善多模态能力的重要一环。

- **[PR] #4622 - 插件(datapaw): 添加包含12个BI技能的数据分析插件 (创建: 2026-05-22)**
    - **链接**: [agentscope-ai/QwenPaw PR #4622](https://github.com/agentscope-ai/QwenPaw/pull/4622)
    - **分析**: 此PR提供了一个功能强大的数据分析插件，至今仍处于开放和审核状态。考虑到项目正在推进插件系统，尽快决策并合并/反馈此PR，将有助于生态建设。

- **[Bug] #5262 - 升级后禁用技能自动恢复 (创建: 2026-06-17)**
    - **链接**: [agentscope-ai/QwenPaw Issue #5262](https://github.com/agentscope-ai/QwenPaw/issues/5262)
    - **分析**: 这是一个影响用户升级体验的严重回归问题。用户此前已报告过类似问题 (#4807)，但未能彻底修复。此问题解决后，可作为提升用户满意度的典型案例。

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是根据您提供的 ZeroClaw 项目 GitHub 数据生成的 2026-06-18 项目动态日报。

---

## ZeroClaw 项目日报 | 日期: 2026-06-18

### 1. 今日速览

今日 ZeroClaw 项目保持**高活跃度**，开发与社区讨论均十分频繁。过去 24 小时内产生了 50 条 Pull Request（PR），显示了团队在多项功能与修复上的并行推进。**其中值得关注的是，41 条 PR 仍处于“待合并”状态，合并/关闭率仅为 18% (9/50)**，可能反映了代码审查流程的瓶颈或存在大量需要进一步讨论的工作。Issues 方面，12 条更新中有 11 条保持活跃状态，其中多个被标记为高风险的 Bug（如 #7871、#7804）和功能请求 (如 #2079) 成为社区焦点。**整体来看，项目正处于密集发力的开发周期，围绕 v0.8.1 集成的功能开发与稳定性修复是当前的核心。**

### 2. 版本发布
**无新版本发布。**

### 3. 项目进展
今日合并/关闭了 **9 个 PR**，其中以下工作尤其值得关注，它们推动了项目在稳定性、安全性与功能方面的进步：

- **稳定性与安全性修复：**
    - **[#7583] fix(runtime): honor profile tool iteration limits**：该 PR 修复了 cron/CLI 模式下工具迭代次数限制不生效的问题，确保了配置文件中的限制在不同运行场景下被正确执行，提升了系统的安全性与可预测性。
    - **[#7492] feat(cost): support cached input token pricing from OpenAI-compatible**: 这是一个大型功能，解析了来自 OpenAI 和 DeepSeek 等兼容接口的缓存 token 信息，使成本追踪器能更精确地计算费用，是成本管理方面的重要进展。
    - **[#7932] fix(docker): correct Node 24 digest pins**：修复了 Docker 构建镜像的 Node.js 版本安全问题，确保构建环境的一致性与安全性。
- **测试覆盖与代码质量：**
    - **[#7918] test: deterministic storage-reader timestamp and ordering regressions**：该 PR 针对 #7694 Issue 中提出的存储读取器边界情况（如同时间戳分页、排序等）添加了确定性回归测试。这直接响应了社区对代码健壮性的诉求，是提升项目长期稳定性的关键举措。
- **关键 Bug 修复：**
    - **[#7927] fix(providers): guard is_non_retryable against false-positive on 429 bodies**：修复了一个潜在的严重问题，即速率限制（429）响应体可能因包含关键词而被误判为永久性错误，导致不必要的重试放弃。**（注：该功能修复版本为 #7930 PR，目前处于待合并状态）**

### 4. 社区热点

今日社区讨论的热点集中在以下几个议题：

- **#2079 [Feature]: Restore GitHub as a native channel**
    - **讨论背景**：该 Issue 是在 2 月就已提出，但在今天仍有活跃评论，表明这是一项社区长期且强烈的诉求。主要讨论点是为何要“恢复”（Restore）以及如何将 GitHub 的原生功能（Issues, PRs 等）集成到 agent 的通道中，以形成统一的交互接口。
    - **背后诉求**：用户希望 Agent 能更深度地融入开发工作流，无需编写复杂的自定义代码（Webhook）即可自动响应和操作仓库事件，实现“代码与 Agent 的无缝互动”。

- **#7694 feat(memory): cover storage-reader timestamp and ordering edge cases**
    - **讨论背景**：作为 #7685 覆盖率跟踪器的一部分，此 Issue 聚焦于 memory/infra/log 组件的测试边界。PR #7916 和 #7918 直接响应了此问题。
    - **背后诉求**：社区贡献者和维护者非常重视系统的**确定性**和**鲁棒性**，特别是在数据持久化和检索场景下。通过对极端情况（如同时间戳）的测试，可防止在生产环境中出现难以排查的数据不一致问题。

- **#6916 feat: process-memory limits on shell/skill_tool subprocess execution**
    - **讨论背景**：由用户在生产环境中遇到 LLM 调用 `wkhtmltopdf` 导致 OOM 的问题而发起。此 Issue 的持续活跃表明这是用户在高强度使用 agent 时的**真实痛点**。
    - **背后诉求**：用户要求从“防止输出 OOM”升级到“**防止进程 OOM**”，即对子进程的内存消耗进行精确控制，确保 Agent 运行环境的稳定性，避免因为一个工具调用而影响整个系统。

### 5. Bug 与稳定性

今日报告了 7 个新 Bug（其中 6 个为 OPEN 状态），按严重程度排列如下：

- **S1 - 工作流阻塞 (工作流完全中断)**
    - **[#7871] shell tool can hang when grandchild processes inherit pipe handles**： `shell` 工具可能在子进程生成孙进程并继承管道句柄时**挂起**。这是一个高危错误，会直接导致 Agent 失去响应。**已有对应修复 PR #7934 被提出。**
    - **[#7804] Code history can send non-alternating Anthropic messages**：当对话历史较长或恢复会话时，向 Anthropic 发送的消息角色不交替，违反了 API 要求，导致 400 错误，**严重阻塞工作流**。**已有对应修复 PR #7931。**
    - **[#7756] native/MCP tools unavailable on OpenAI Responses/reasoning and Anthropic turns**：即使 MCP 服务器已成功连接并注册了工具，部分模型可能无法“接收”到这些工具，导致工具调用功能失效，严重依赖 MCP 的用户工作流将受阻。**已有对应修复 PR #7933。**

- **S2 - 功能退化 (核心功能异常)**
    - **[#7799] [已关闭] Resumed Code sessions reopen with a blank transcript**：恢复已保存的 Code 会话时，即使会话中存在消息，也可能显示空白记录。此问题**已于今日被关闭**，显示已被修复。

- **S3 - 次要问题 (可用性或用户体验受损)**
    - **[#7892] CLI approval prompt should read controlling terminal when stdin is detached**：当进程的标准输入被重定向时，CLI 审批提示仍从 stdin 读取，导致权限交互失败，用户无法进行操作。
    - **[#7917] i18n: file_download tool strings untranslated in all non-English locales**： `file_download` 工具的描述和错误信息在所有非英语区域设置均显示为英文，影响多语言用户的体验。**已有多个修复 PR（#7924、#7925）被提出。**

### 6. 功能请求与路线图信号

用户提出的新功能请求主要集中在以下方向，结合已有 PR 可判断其可能的纳入路线：

- **用户界面与体验统一化**：
    - **[#7929] Unify slash-command registries across web UI, zerocode TUI, and channel runtime**：用户提出将分布在三处（Web UI, TUI, 通道运行时）的斜杠命令注册中心统一为一个网关提供的目录服务。**此提议有望显著降低开发复杂度和体验不一的问题，可能会被纳入下一个次要版本规划。**
- **高级通道集成**：
    - **[#7922] feat(channels/discord): slash command localizations + guild scope**：此 PR 为 Discord 通道添加了本地化命令描述和公会范围注册功能，这是完善 Discord 集成的重要一步，**预计会被并入 v0.8.1 的集成队列中。**
- **平台扩展与生态支持**：
    - **[#7928] feat(wasi): initial WASM component-model plugin host code**：新增 WASM 组件模型插件宿主代码，这为 ZeroClaw 引入了一个轻量级、安全的**插件生态系统**。这是一个重要的基础设施级扩展，**可能标志着项目未来架构方向的关键一步。**
- **性能与资源管理**：
    - **[#7923] feat(auto-clean): Supports automatic clearing of temporary files**：添加了自动清理临时文件的功能，允许用户通过配置文件声明清理目标和阈值。**这是对@alex-nax 提出的内存限制诉求（#6916）的一个补充，共同组成了对 Agent 资源管理的更完整方案。**

### 7. 用户反馈摘要

从 Issues 的评论和描述中，可以提炼出以下用户的真实声音：

- **对安装和使用便利性的需求**：
    - “I am trying to install ZeroClaw on Android-Termux. / 我尝试在 Android-Termux 上安装 ZeroClaw。…I noticed that both trying to install the precompiled binary and compiling on the local device result in the unknown linux aarch64 binary being installed.” — **#7911** 用户反馈在 Android 的 Termux 环境中无法正确安装，表明项目在主流平台之外的支持仍有待加强。
- **对稳定性和安全性的高度关注**：
    - “This was observed in production when an LLM fell back to shell commands / 在生产环境中观察到，当 LLM 回退到 shell 命令时... consumed all available container memory” — **#6916** 用户报告了生产环境下的真实事故，显示出对 Agent 工具调用（特别是 Shell 工具）资源控制的**极度担忧**。
    - “The resume itself appears to attach to the... / 恢复本身似乎连接到了...” — **#7799** 用户对会话恢复功能失效表达了困惑，这是对用户体验的严重打击，好在此问题已被修复。
- **对功能集成的期待**：
    - “Integrating GitHub requires custom glue / 集成 GitHub 需要自定义胶水代码” — **#2079** 老用户再次强调了对原生 GitHub 通道的渴望，反映出社区希望平台能开箱即用地支持主流开发工具，而非需要自行搭建。

### 8. 待处理积压

以下 Issue 和 PR 已存在较长时间且未有关键响应，可能需要项目维护者关注：

- **[#2079] [Feature]: Restore GitHub as a native channel** (创建于 2026-02-27，今日仍有活跃讨论)
    - **链接**: [Issue #2079](zeroclaw-labs/zeroclaw Issue #2079)
    - **风险**: 作为一项社区呼声很高的**长期未决**功能请求，可能会影响一部分潜在用户（特别是开发者）的采用意愿。不过，PR #7162 和 #7170 等围绕通道的活跃工作，或许表明此问题正在被间接推进。
- **[#6970] [Tracker]: v0.8.1 integration/channel/provider/tool queue and history** (创建于 2026-05-27)
    - **链接**: [Issue #6970](zeroclaw-labs/zeroclaw Issue #6970)
    - **风险**: 这是一个重要的**版本跟踪器**，虽已合并一些相关 PR 并有关注，但作为一次小版本的里程碑，需要持续跟进以确保所有子任务协调推进，避免版本发布推迟。
- **[#7926] fix(runtime): restore SKILL.md always: true for compact prompt mode (#7904)** (创建于今日)
    - **链接**: [PR #7926](zeroclaw-labs/zeroclaw PR #7926)
    - **风险**: 该 PR 旨在修复一个在重构中引入的回归问题（#7904），即 `always: true` 功能失效。虽然今日已提交，但鉴于其主要在中国假期期间提交（提交者可能有时差），需要尽快安排 Review 以避免该功能在下一个版本中继续失效。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*