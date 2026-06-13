# OpenClaw 生态日报 2026-06-13

> Issues: 147 | PRs: 500 | 覆盖项目: 13 个 | 生成时间: 2026-06-13 10:15 UTC

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

# OpenClaw 项目日报 – 2026-06-13

## 1. 今日速览

过去 24 小时内，OpenClaw 项目保持极高活跃度：共处理 147 条 Issue（新开/活跃 81，关闭 66）和 500 条 Pull Request（待合并 343，合并/关闭 157）。新发布两个 beta 版本（v2026.6.7-beta.1、v2026.6.6），重点强化了跨渠道投递一致性与安全边界。社区讨论集中在高风险 regressions（如 DeepSeek Prompt Cache 失效、memory_search 元数据丢失）和长期悬而未决的安全功能（私有网络访问、权限模型）。已合并的 PR 覆盖了 Discord 连接稳定性、子代理生命周期修复、模型缓存优化等关键改进，项目整体健康度良好，但部分 P0/P1 问题仍在等待维护者决断。

---

## 2. 版本发布

### v2026.6.7-beta.1

- **发布时间**：2026-06-13（基于数据中的最新 release）
- **亮点**：
  - 渠道投递一致性增强：Slack、Telegram、outbound media、silent replies、progress drafts 及分页 action results 中的同渠道 Slack finals 现正确持久化到会话记录中；顶层 `image` 消息工具可发送附件媒体；Telegram 支持可展开的 blockquote 和 spool。
- **破坏性变更**：暂未明确标注。
- **迁移建议**：若使用 Telegram blockquote 功能，需更新客户端版本以支持可展开样式；`image` 工具行为变更可能影响依赖旧有附件格式的自定义技能。

### v2026.6.6

- **发布时间**：2026-06-13
- **亮点**：
  - 安全边界大幅收紧：涉及会话记录、沙箱绑定、宿主环境继承、MCP stdio、Codex HTTP 访问、原生搜索策略、发送者校验（elevated sender checks）、已删除代理的 ACP 绕过、loopback 工具、Discord 审核及 Teams 组操作。所有安全敏感路径均增加了更严格的权限校验。
- **破坏性变更**：收紧的安全策略可能导致原有配置（如 MCP 或 Codex 外部访问）被拒绝，需管理员复查 `tools.web.fetch.allowPrivateNetwork` 等开关配置。
- **迁移注意**：升级后若遇到工具调用被拒绝，请检查沙箱权限、环境变量白名单及安全策略配置；Discord 审核功能需重新配置机器人权限。

---

## 3. 项目进展

以下为今日已合并/关闭的重要 Pull Request（基于展示且状态为 CLOSED 的高影响力 PR）：

| PR 编号 | 标题 | 影响域 | 说明 |
|--------|------|--------|------|
| [#89438](https://github.com/openclaw/openclaw/pull/89438) | fix(slack): warn when channels map is keyed by name | Slack 通道路由 | 当用户错误地使用通道名称而非 ID 配置 Slack 映射时，现在会发出警告，避免静默配置失效。 |
| [#90660](https://github.com/openclaw/openclaw/pull/90660) | fix(discord): filter startup slots by runtime-enabled predicate | Discord 启动 | 修复重复 token 账户导致启动槽位冲突的问题，使用运行时启用谓词过滤。 |
| [#92651](https://github.com/openclaw/openclaw/pull/92651) | fix: require admin for HTTP session kills | 网关安全 | 修复 HTTP 会话杀死端点 `POST /sessions/:sessionKey/kill` 的认证漏洞，现在要求 `operator.admin` 作用域。 |
| [#74914](https://github.com/openclaw/openclaw/pull/74914) | [codex] allow native google-vertex model provider configs | Google Vertex 集成 | 允许在配置中使用 `google-vertex` 作为 provider，并更新配置模式以正确验证。 |
| [#92652](https://github.com/openclaw/openclaw/pull/92652) | test(models): stabilize plugin auth marker fixtures | 测试基础设施 | 修复插件认证标记的单元测试，使其不依赖运行环境，提升测试可靠性。 |
| [#89835](https://github.com/openclaw/openclaw/pull/89835) | feat(usage): native templated /usage full footer renderer | 命令行/UI 体验 | 内联渲染 `/usage full` 脚注模板，替代原有跨组件依赖，提升响应速度。 |
| [#90686](https://github.com/openclaw/openclaw/pull/90686) | fix(gateway): honor profile auth for SecretRef model entries | 模型认证 | 修复 SecretRef 模型条目中认证配置文件未被正确使用的 bug。 |

此外，`#91778`（memory_search 元数据丢失，P0）和 `#91018`（DeepSeek Prompt Cache 失效，P1）已被标记为 CLOSED，但未见关联 PR 说明，推测通过配置调整或已在 v2026.6.7-beta.1 中修复。

**总结**：今日修复集中在渠道连通性（Slack、Discord）、网关安全（session kill 认证）、模型认证（Google Vertex、SecretRef）、性能（usage 模板渲染），项目在稳定性与安全性上取得显著进展。

---

## 4. 社区热点

以下为今日评论数最高的 Issues（截至展示）：

| Issue | 标题 | 评论数 | 核心诉求 |
|-------|------|-------|----------|
| [#54253](https://github.com/openclaw/openclaw/issues/54253) | [Bug]: OpenClaw returns "run Error : LLM Request Failed" on RISC-V64 System | **14** | 用户在 RISC-V64 架构上成功安装但运行时报错，推测为底层请求处理兼容性问题。 |
| [#39604](https://github.com/openclaw/openclaw/issues/39604) | [Feature]: Add tools.web.fetch.allowPrivateNetwork to allow private network access | **13** | 强烈要求增加私有网络访问控制开关，当前 `web_fetch` 默认禁止访问内网地址，影响开发/测试场景。 |
| [#39476](https://github.com/openclaw/openclaw/issues/39476) | A2A sessions_send: target agent can call sessions_send back, causing duplicate messages | **10** | 跨代理消息发送导致重复消息，暴露了 A2A 会话路由设计的缺陷。 |
| [#7707](https://github.com/openclaw/openclaw/issues/7707) | Feature Request: Memory Trust Tagging by Source | 9 | 内存来源信任标签功能，防止恶意内容通过网页或第三方技能投毒。 |
| [#41165](https://github.com/openclaw/openclaw/issues/41165) | [Bug]: Telegram DMs can still land in agent:main:main | 8 | 修复后仍有 Telegram 私信路由到错误会话的问题。 |
| [#91778](https://github.com/openclaw/openclaw/issues/91778) | memory_search cassé — index metadata is missing | 8 | P0 回归：memory_search 的向量索引元数据丢失，导致所有代理失明。 |

**分析**：  
- **RISC-V 兼容性** (#54253) 是架构扩展需求，虽非核心用户群，但呼声高，建议维护者提供日志收集指引或尝试在 CI 中增加交叉编译测试。  
- **私有网络访问** (#39604) 是安全与可用性的平衡问题，已有 PR #40311 提出类似功能（Brave Goggles），但用户希望更通用的 `allowPrivateNetwork` 配置。  
- **A2A 消息重复** (#39476) 是协议设计问题，影响多代理协作场景，需要核心开发人员介入重构路由逻辑。  
- **内存信任标签** (#7707) 已被标记为 `stale` 但评论始终活跃，说明社区对安全防御有长期需求。

---

## 5. Bug 与稳定性

按严重程度排列今日报告的关键 Bug（含过去 24 小时新活跃的 P0/P1 问题）：

| 严重性 | Issue | 标题 | 状态 | Fix PR |
|--------|-------|------|------|--------|
| **P0** | 无新 P0 | — | — | — |
| **P1** | [#91018](https://github.com/openclaw/openclaw/issues/91018) | ⚠️ 升级 2026.6.1 后 DeepSeek Prompt Cache 失效，一小时烧掉 ~$6 | CLOSED | 已修复（至少标记为关闭） |
| **P1** | [#91016](https://github.com/openclaw/openclaw/issues/91016) | 同 #91018 中文版 | OPEN | 无，等待关闭 |
| **P1** | [#41165](https://github.com/openclaw/openclaw/issues/41165) | Telegram DMs 仍可路由到错误会话 | OPEN | 无，需维护者复查 |
| **P1** | [#41346](https://github.com/openclaw/openclaw/issues/41346) | 自动注册的 cron 任务默认启用且无成本防护 | OPEN | 无 |
| **P1** | [#39476](https://github.com/openclaw/openclaw/issues/39476) | A2A sessions_send 消息重复 | OPEN | 无 |
| **P1** | [#66443](https://github.com/openclaw/openclaw/issues/66443) | 溢出恢复导致用户消息重复 | OPEN | 无 |
| **P1** | [#70903](https://github.com/openclaw/openclaw/issues/70903) | 文件型 provider 冷却持续数小时，即使账单已恢复 | OPEN | 无 |
| **P1** | [#92094](https://github.com/openclaw/openclaw/issues/92094) | message tool 对 Telegram 返回 "unsupported channel" | OPEN | 无 |
| **P1** | [#92479](https://github.com/openclaw/openclaw/issues/92479) | Zen provider 不包含模型目录，所有模型需手动注册 | OPEN | 无 |
| **P1** | [#38721](https://github.com/openclaw/openclaw/issues/38721) | 网关关闭超时、子进程句柄残留 | OPEN | 无 |
| **P1** | [#39847](https://github.com/openclaw/openclaw/issues/39847) | 内部元数据泄露（echo contamination） | OPEN | 无 |

**注**：今日没有新报告的 P0 问题，但多个 P1 问题长期存在且无对应 Fix PR，需要在下一版本中重点解决。

**稳定性要点**：  
- DeepSeek Prompt Cache 失效问题已标记关闭，但仍有中文 Issue #91016 处于 OPEN，建议同步更新。  
- 网关关闭超时（#38721）、cron 安全（#41346）、Telegram 路由（#41165）等是已知的稳定性短板。

---

## 6. 功能请求与路线图信号

以下为用户提出的高票功能需求，部分已有关联 PR，可能被纳入 v2026.6.x 后续版本：

| Issue | 标题 | 票数 👍 | 关联 PR | 优先级判断 |
|-------|------|---------|---------|-----------|
| [#39604](https://github.com/openclaw/openclaw/issues/39604) | 添加私有网络访问开关 | 9 | [#40311](https://github.com/openclaw/openclaw/pull/40311)（Brave Goggles） | 高：社区呼声强，PR 已有实现但未合并。 |
| [#7707](https://github.com/openclaw/openclaw/issues/7707) | 内存信任标签 | 0（但评论9） | 无 | 中：安全增强，但实现复杂，需要安全评审。 |
| [#39979](https://github.com/openclaw/openclaw/issues/39979) | 路径级 RWX 权限 | 0 | 无 | 中：替代二进制 exec 白名单，更符合 Unix DAC 原则。 |
| [#40418](https://github.com/openclaw/openclaw/issues/40418) | 自动会话记忆保存与综合 | 1 | 无 | 低：涉及核心会话管理，优先级暂被旧bug挤占。 |
| [#65403](https://github.com/openclaw/openclaw/issues/65403) | 代理电话会议原语 | 1 | 无 | 低：创意提议，但短期无实施计划。 |
| [#74261](https://github.com/openclaw/openclaw/issues/74261) | WhatsApp 创建群组工具 | 1 | 无 | 低：单一平台功能，需维护者评估 ROI。 |

**路线图信号**：  
- `tools.web.fetch.allowPrivateNetwork` 是安全与可用性矛盾的核心，已有实现但需维护者拍板。  
- 内存信任标签（#7707）与路径权限（#39979）均被标记为 `needs-security-review`，说明安全团队已纳入视野。  
- 会话记忆综合（#40418）和代理电话会议（#65403）属于长线创新，社区关注度较低，短期内不会进入主线。

---

## 7. 用户反馈摘要

从 Issue 评论中提炼的真实用户痛点与场景：

- **RISC-V 架构用户**：对项目支持感到兴奋，但安装成功后运行时报错，缺乏明确错误日志指引，社区需要排查环境 (Node.js 版本、系统调用兼容性) 差异。
- **成本敏感用户**：DeepSeek Prompt Cache 失效导致分钟级消费激增（$6/h），强烈要求建立成本断路器或缓存失效报警机制（#91016）。
- **多代理部署用户**：A2A 消息重复（#39476）和子代理输出丢失（#92631 PR 试图修复）导致业务流程不可靠，期望统一 session 路由语义。
- **Mac 桌面用户**：启动项设置更新后自动重置，且 Chrome 在后台下载大模型文件（~4GB），浪费磁盘与网络（#39588, #41088）。
- **Telegram 重度用户**：DM 路由异常（#41165）和 message tool 返回“unsupported channel”（#92094）严重影响日常使用。
- **开发者社区**：对 `web_fetch` 无法访问内网（#39604）表示不满，认为应提供可控开关；同时期待更简洁的权限模型（#39979）。

**满意点**：用户对项目的跨平台支持（x86、Apple Silicon、RISC-V）和频繁更新表示赞赏；新发布的 channel delivery 增强被社区认为是“一直以来缺失的功能”。

**不满/担忧**：部分 P1 问题长期存在（如 Telegram 路由、网关关闭超时），且修复经过多个版本仍未根除，引起社区对质量管理的质疑。

---

## 8. 待处理积压

以下为长期未响应或缺少维护者关注的重要 Issue/PR，需提醒团队优先处理：

### 高影响 Issue（超过 30 天未实质性推进）

| Issue | 标题 | 创建时间 | 上次更新 | 状态标签 | 优先原因 |
|-------|------|---------|---------|---------|----------|
| [#39604](https://github.com/openclaw/openclaw/issues/39604) | 私有网络访问开关 | 2026-03-08 | 2026-06-13 | `needs-product-decision` `needs-security-review` | 社区希望3个月，已有 PR #40311，但缺乏维护者决定。 |
| [#7707](https://github.com/openclaw/openclaw/issues/7707) | 内存信任标签 | 2026-02-03 | 2026-06-13 | `stale`, `needs-maintainer-review`, `needs-security-review` | 安全基础功能，若被攻破会影响大多数用户。 |
| [#39979](https://github.com/openclaw/openclaw/issues/39979) | 路径级 RWX 权限 | 2026-03-08 | 2026-06-13 | `stale`, `needs-maintainer-review`, `needs-security-review` | 与 #39604 同为安全替代方案。 |
| [#41346](https://github.com/openclaw/openclaw/issues/41346) | 自动 cron 任务默认启用无成本防护 | 2026-03-09 | 2026-06-13 | `needs-maintainer-review`, `needs-product-decision`, `needs-security-review` | 新增 cron 可能直接产生费用，属于设计缺陷。 |
| [#39476](https://github.com/openclaw/openclaw/issues/39476) | A2A 消息重复 | 2026-03-08 | 2026-06-13 | `needs-product-decision`, `linked-pr-open` | 影响多代理协作场景核心体验。 |
| [#41165](https://github.com/openclaw/openclaw/issues/41165) | Telegram DM 路由错误 | 2026-03-09 | 2026-06-13 | `needs-product-decision` | 修复后仍有残留问题。 |
| [#38721](https://github.com/openclaw/openclaw/issues/38721) | 网关关闭超时 | 2026-03-07 | 2026-06-13 | `needs-maintainer-review`, `needs-live-repro` | 生产环境常见问题，影响容器化部署稳定性。 |

### 需要关注的 PR

| PR | 标题 | 状态 | 原因 |
|----|------|------|------|
| [#40311](https://github.com/openclaw/openclaw/pull/40311) | feat(web-search): expose Brave Goggles | OPEN, `needs-real-behavior-proof` | 与 #39604 高度相关，可为私有网络访问提供实现参考。 |
| [#83933](https://github.com/openclaw/openclaw/pull/83933) | fix(cron): skip deleteAfterRun and preserve counters | OPEN, `waiting on author` | 修复 cron 手动运行破坏调度状态，已等待近一个月。 |
| [#92631](https://github.com/openclaw/openclaw/pull/92631) | fix(agents): pause yielded subagent runs | OPEN, `ready for maintainer look` | 今天新提交，直接修复 #92448 子代理 yield 崩溃问题，需快速审查。 |
| [#92639](https://github.com/openclaw/openclaw/pull/92639) | fix(memory): keep memory_search in transient qmd mode | OPEN, `waiting on author` | 修复 memory_search 性能问题，需作者补充测试。 |

**建议**：维护团队本周集中评审 #39604 (私有网络访问) 和 #92631 (子代理 fix)，这两个问题直接关系到用户体验与系统稳定性。长期积压问题如 #7707、#39979 可以在产品主管确认方向后，指定给安全团队制定实现规范。

---

*数据来源：OpenClaw GitHub 仓库，统计截止 2026-06-13 23:59 UTC。*

---

## 横向生态对比

好的，作为AI智能体与个人AI助手领域的开源项目分析师，以下是根据您提供的2026年6月13日各项目动态，生成的横向对比分析报告。

---

# 个人AI助手开源生态横向分析报告 (2026-06-13)

## 1. 生态全景

当前，个人AI助手/自主智能体开源生态整体呈现 **“核心平台成熟化，边缘创新碎片化”** 的态势。以OpenClaw为代表的超大规模项目已进入功能深化与安全加固并重的阶段，其社区围绕稳定性、成本控制和权限模型展开了激烈讨论。与此同时，以NanoBot、Hermes Agent为代表的新锐项目则在快速迭代中，通过引入TUI、TTS、多平台适配等特性，争夺特定场景下的用户体验高地。一个显著的趋势是，**社区对“静默失败”和“成本失控”的容忍度急剧降低**，这标志着整个生态正从“能运行”的早期阶段，转向“可靠、经济、安全”的成熟期。

## 2. 各项目活跃度对比

| 项目名称 | Issues (新开/活跃) | PR总数 (待合并/合并关闭) | 版本发布 | 健康度评估 |
| :--- | :--- | :--- | :--- | :--- |
| **OpenClaw** | 81 | 343 / 157 | v2026.6.7-beta.1, v2026.6.6 | **极高度活跃**，核心功能迭代快，但P0/P1 Bug积压，社区健康度良好但存在质量隐忧。 |
| **NanoBot** | 12 | 25 / 10 | 无 | **高度活跃**，功能迭代迅速（TTS、TUI、审计），社区响应积极，处于快速上升期。 |
| **Hermes Agent** | 12 | 45 / 5 | 无 | **贡献踊跃但审核瓶颈严重**，大量PR待合并，可能拖慢创新落地。 |
| **PicoClaw** | 7 | 7 / 3 | v0.2.9-nightly | **中等活跃**，集中于修复与通道扩展，Nightly版本稳定，适合早期用户尝鲜。 |
| **NanoClaw** | 11 | 12 / 10 | 无 | **活跃且高效**，合并率高，Signal、Ollama等关键功能落地迅速。 |
| **NullClaw** | 2 | 3 / 0 | 无 | **维护驱动**，活跃度偏低，核心Bug（定时任务）搁置已久。 |
| **IronClaw** | 4 | 50 (大量) | 无 | **极高度活跃**，核心贡献者主导开发，功能管线（附件）清晰，架构设计严谨。 |
| **LobsterAI** | 4 | 5 / 5 | 无 | **稳定开发**，合入了里程碑功能（Computer Use），但历史Bug积压未解。 |
| **CoPaw** | 11 | 3 / 1 | 无 | **中等活跃**，新Bug（Gemini回归）和性能退化问题突出，社区反馈强烈。 |
| **Moltis** | 2 | 0 / 0 | 无 | **偏低活跃**，仅有新Issue，无PR或版本产出，处于静默期。 |
| **ZeroClaw** | 2 | 50 / 26 | 无 | **高度活跃**，PR流动量大，但多数为大型重构PR，生态正酝酿重大变更。 |
| **TinyClaw / ZeptoClaw** | 0 | 0 / 0 | 无 | **无活动**，可能处于开发间歇期或社区维护力度不足。 |

## 3. OpenClaw 在生态中的定位

OpenClaw 在生态中扮演着 **“基础设施标准制定者”** 的角色，其定位与优势如下：

- **优势**：
    - **功能广度**：在渠道一致性、安全边界、模型供应商支持（包括Google Vertex）、测试基础设施等方面全面领先，是功能最完备的“一站式”解决方案。
    - **社区规模**：拥有最高的Issue和PR处理量，生态活跃度远超其他项目。其每日处理的PR数是大多数项目的5-10倍。
    - **版本成熟度**：拥有明确的beta版本发布策略，对破坏性变更的迁移路径有所提及，展现了成熟项目的规范。

- **差异化与瓶颈**：
    - **与NanoBot/CoPaw等新兴项目相比**：OpenClaw更侧重于**架构的稳健性和安全性**（如安全边界收紧），而NanoBot和CoPaw更侧重于**开发者的易用性和功能创新**（如TUI、插件生态、本地模型集成）。
    - **社区治理挑战**：极高的PR与Issue数量也带来了“合并瓶颈”，大量高票功能请求（如私有网络访问）和P1级Bug（如Telegram路由错误）长期存在，维护者的决策速度成为制约其发展的主要因素。

## 4. 共同关注的技术方向

多个项目不约而同地涌现出以下技术需求：

| 关注方向 | 涉及项目 | 具体诉求 |
| :--- | :--- | :--- |
| **成本控制与可观测性** | **OpenClaw**、**NanoBot**、**CoPaw** | DeepSeek Prompt Cache失效导致成本激增、API返回零Token统计、定时任务缺乏成本防护。 |
| **记忆与上下文管理** | **OpenClaw**、**ZeroClaw**、**CoPaw** | memory_search元数据丢失、Agent对话总结逻辑缺陷、SQLite向量搜索效率低、长对话卡死。 |
| **多模态与媒体集成** | **IronClaw**、**NanoClaw**、**PicoClaw**、**Hermes Agent** | 统一“附件”概念、支持图像模型输入、内联Data URL媒体提取、Discord附件读取。 |
| **安全边界与权限模型** | **OpenClaw**、**NanoClaw**、**CoPaw** | 私有网络访问、Agent容器安全加固、MCP工具权限隔离、技能信任标签。 |
| **跨平台/跨渠道连通性** | **OpenClaw**、**Hermes Agent**、**CoPaw**、**PicoClaw** | Telegram DM路由、Matrix加密、WhatsApp群组路由、Discord断连恢复、Zalo Bot支持。 |
| **开发者体验与CLI/UI** | **NanoBot**、**OpenClaw**、**ZeroClaw** | 终端TUI、WebUI与配置文件双向同步、反向代理支持、子路径部署、快捷方式本地化（macOS Cmd键）。 |

## 5. 差异化定位分析

| 项目 | 功能侧重 | 目标用户 | 技术架构/哲学 |
| :--- | :--- | :--- | :--- |
| **OpenClaw** | **通用AI助手平台** | 寻求一体化解决方案的高级用户和团队 | 功能全面、安全优先、架构稳重 |
| **NanoBot / CoPaw** | **开发者效率工具** | 追求快速上手和功能创新的开发者 | 插件/技能生态、WebUI友好、CLI创新、本地模型先行 |
| **Hermes Agent** | **私有化与安全通信** | 注重隐私和平台原生体验的用户 | 针对Matrix/Telegram等平台深度适配、端到端加密 |
| **IronClaw** | **可扩展的企业级Agent** | 有大模型集成和复杂工作流需求的企业团队 | 模块化设计、精细的设计回顾、长线追踪 |
| **ZeroClaw** | **前沿架构探索** | 关注Agent架构演进的研究型开发者 | 敢于进行核心重构（运行时）、引入“梦境”等创新概念 |
| **PicoClaw / NanoClaw** | **特定场景工具** | 轻量级部署、嵌入式场景或特定社区 | 核心功能精简、提供特有的通道（Signal）或轻量化部署 |

## 6. 社区热度与成熟度

基于数据分析，可将项目分为三个梯队：

- **第一梯队（超大规模，高速迭代）**：**OpenClaw**、**Hermes Agent**
    - **特征**：PR与Issue数均远超其他项目，社区讨论海量。但同时也面临“贡献多，审核难”的质量管理挑战。OpenClaw已进入“质量巩固”阶段，而Hermes Agent仍处于“快速迭代”阶段。

- **第二梯队（快速成长，创新涌现）**：**NanoBot**、**CoPaw**、**NanoClaw**、**IronClaw**、**ZeroClaw**
    - **特征**：社区活跃度稳定，功能迭代方向和路线图明确，用户反馈与开发响应较为迅速。NanoBot和CoPaw偏向易用性创新，IronClaw和ZeroClaw偏向架构深度。

- **第三梯队（市场细分，维护驱动）**：**PicoClaw**、**NullClaw**、**LobsterAI**、**Moltis**
    - **特征**：项目规模较小，活跃度波动大，可能聚焦于特定细分市场或处于维护模式。社区讨论密度低，Bug修复周期较长。

## 7. 值得关注的趋势信号

1.  **“成本恐慌”成为首要痛点**：OpenClaw用户报告“一小时烧掉$6”，CoPaw用户因Gemini回归导致工具调用“失效”。这表明**开发者对LLM API的成本敏感度极高**，缺乏成本断路器、缓存失效告警等机制的项目将面临用户流失。**建议开发者**：在项目中内置token用量统计和警告功能。

2.  **记忆与上下文管理的“第二次战役”**：多个项目同时出现memory_search元数据丢失、长对话死循环、向量搜索性能问题。这标志着，在解决基础对话功能后，**长期记忆、上下文摘要和高效的检索机制已成为下一代智能体的核心挑战**。**建议开发者**：关注SQLite ANN索引、选择性压缩等提升记忆系统可靠性和效率的方案。

3.  **“沉默是金”的终结**：无论是Agent定时任务未执行但显示完成，还是在预算耗尽时静默返回200，抑或短时间多轮响应被去重丢弃，**“静默失败”问题正在被社区视为“不可接受”的重大Bug**。**建议开发者**：在设计任务执行、资源限制、消息路由等关键路径时，确保有明确的错误反馈和超时机制，避免用户陷入“它到底干活了没？”的焦虑。

4.  **UI/UX正从“能用”走向“好用”**：NanoBot的TUI、CoPaw桌面端启动性能退化、ZeroClaw Shift+Enter换行问题，都指向了**开发者对非功能性需求（体验、性能、操作习惯）的日益重视**。**建议开发者**：在投入核心功能的同时，不要忽视CLI/WebUI的交互细节和性能基线，这往往是建立用户忠诚度的关键。

5.  **横向集成与平台扩展成为新战场**：从Hermes加大对Matrix加密的投入，到CoPaw涌现对Zalo Bot的需求，再到IronClaw对Slack授权的优化，可以看出，**项目们正努力成为用户已有沟通生态的一部分**，而不仅仅是另一个独立的App。**对开发者**：拥抱开放协议（如WebSocket、Signal、Matrix），并优先解决与主流平台的集成稳定性，将获得更强的市场竞争力。

---

## 同赛道项目详细报告

<details>
<summary><strong>NanoBot</strong> — <a href="https://github.com/HKUDS/nanobot">HKUDS/nanobot</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，根据您提供的 NanoBot 项目 GitHub 数据，我为您生成了 2026-06-13 的项目动态日报。

---

### NanoBot 项目动态日报 | 2026-06-13

---

#### 1. 今日速览

今日 NanoBot 项目社区**活跃度极高**，核心开发与社区贡献并行。总计 35 条 PR 更新，其中 10 条已合并或关闭，显示出上游核心团队与社区贡献者均在高频输出代码。Bug 修复集中于会话管理、WebUI 和服务端 API，同时涌现了多项重要的新功能提案，如 TUI 交互界面、TTS 配置系统和审计模块。项目整体处于功能快速迭代与稳定性增强并重的阶段。

---

#### 2. 版本发布

今日无新版本发布。

---

#### 3. 项目进展

今日有 10 个 Pull Requests 被合并或关闭，部分重要进展如下：

- **WebUI 功能与配置同步**：PR #4313 `Feat(webui): config.json/webui parity` 被关闭并合并。该 PR 显著增强了 WebUI 设置面板与 `config.json` 之间的同步性，新增了温度、工具限制、记忆等字段的写入端点，使 WebUI 几乎可以完全替代手动编辑配置文件。
- **审计模块（Audit Module）**：社区贡献者 `bjoshuanoah` 提交了多个关于审计功能的 PR（#4318, #4319, #4320）。其中 #4318 和 #4319 已被关闭（可能与 #4320 合并或取代），但#4320 仍为打开状态。这表明项目正在探索为 Agent 工具调用添加日志审计的能力，以提升可观测性。
- **MCP 服务器稳定性**：PR #4303 `fix(mcp): close tracked generators in _close_server to prevent GC crash` 修复了一个关键的 MCP 连接崩溃问题，该问题在会话重连时会导致 Python 运行时错误。此修复对依赖 MCP 的用户至关重要。
- **错误数据处理**：PR #4315 `fix(memory): ignore malformed history entries` 被提出，旨在使记忆系统更健壮，能够优雅地忽略 `history.jsonl` 中的损坏记录，防止读取时崩溃。
- **输入验证增强**：PR #4312 和 #4311 分别从媒体附件和文件分页参数两个方面加强了输入验证，防止因无效输入导致工具调用异常。

这些进展表明，NanoBot 在提升用户在 WebUI 上的体验、增强系统稳定性和健壮性方面取得了扎实的进展。

---

#### 4. 社区热点

今日讨论的热点集中在**基础设施与易用性**上。

- **最受关注的新功能提议**：PR #4329 `Nanobot TUI` 和 PR #4328 `webui: serve correctly under a path prefix (reverse proxy / sub-path)` 均是今日新提交的 PR，分别提供了全新的终端交互体验和解决反向代理部署场景下的兼容性问题。这两项改进直接回应了终端用户和运维人员的核心痛点，预计会引发大量讨论。
- **核心功能缺陷**：Issue #4203 `Bug: find_legal_message_start...` 尽管已于昨日关闭，但其描述的场景——用户消息后紧跟孤立工具结果导致所有消息被丢弃——是一个影响对话连贯性的严重逻辑缺陷，是社区关注的焦点。
- **基础设施需求**：Issue #4305 `[enhancement] Multiple custom providers` 讨论了如何支持多个自定义（Custom）提供商。这反映出用户有运行多个自建或非主流模型服务的需求，项目当前的单一自定义提供商配置模型限制了灵活性。

---

#### 5. Bug 与稳定性

以下为今日报告的 Bug，按严重程度排列：

- **严重：`NameError` 导致启动崩溃**
  - **描述**：Issue #4322 `NameError: 'session_key' is not defined in context.py after merge` 报告了一个因代码合并导致的回归问题，令 Agent 在启动时立即崩溃。这是一个严重的阻断了用户使用的 Bug。
  - **状态**：开放中，暂无关联 PR。

- **中：`idleCompact` 会话总结逻辑缺陷**
  - **描述**：Issue #4264 `idleCompact should use the complete session history...` 指出 `idleCompact` 机制存在逻辑缺陷，它没有将对话末尾的关键纠正/结果输入总结模型，可能导致历史记录存储错误的信息。
  - **状态**：已有修复 PR #4326 `fix(memory): summarize full session tail during idle compaction`，开发团队响应迅速。

- **中：API 返回零 Token 统计**
  - **描述**：Issue #4309 `nanobot serve: /v1/chat/completions always returns zero usage tokens` 指出 OpenAI 兼容 API 的返回值中，Token 使用量始终为零，这影响了需要精确计费或用量分析的用户。
  - **状态**：开放中，暂无关联 PR。

---

#### 6. 功能请求与路线图信号

- **终端用户界面（TUI）**：PR #4329 提交了全新的内联（非全屏）交互式 TUI，支持 Markdown 渲染、斜杠命令等功能。如果被合并，这将极大改善无图形界面环境下的用户体验。
- **多自定义提供商支持**：Issue #4305 提出的需求，已有多个开发者在评论中表示赞同。这是一个明确的社区需求信号，很可能在后续版本中被采纳。
- **文本转语音（TTS）**：PR #4316 `feat(tts): add TTS configuration system with multi-provider support` 为项目引入了 TTS 功能，支持 OpenAI、Groq 和 ElevenLabs。这表明项目正在向多模态交互扩展。
- **审计/可观测性**：PR #4320 `feat(audit): add tools.audit config and AuditTool...` 表明项目正在构建 Agent 行为的审计框架，这通常是生产环境部署和合规要求的必备能力。

---

#### 7. 用户反馈摘要

从今日的 Issues 和 PR 评论中，可以提炼出以下用户痛点与使用场景：

- **配置管理的局限性**：用户 `smurfix` 在 Issue #4305 中提出“我需要多个自定义提供商”，直接点明了当前配置模型无法满足部分高级用户需求。
- **反代部署的普遍需求**：PR #4328 的提交者 `niradler` 意在解决 WebUI 在反向代理或子路径（sub-path）下无法正常工作的普遍性问题。这反映出许多用户将 NanoBot 部署在企业内网或利用 Nginx 等工具进行托管。
- **对对话准确性的高要求**：用户 `imkuang` 在 Issue #4264 中详细描述了用户与模型多轮纠错并最终成功的场景，他担心 `idleCompact` 会“留下错误的记录和结论”，这体现了用户对对话历史和模型行为准确性的重视。

---

#### 8. 待处理积压

今日未发现长期未响应的重要 Issue 或 PR。以下两点建议维护者关注：

- **Issue #4309: `/v1/chat/completions` 返回零 Token**：这是一个标准 API 的行为缺陷，对集成方（如计费系统）影响直接，应尽快确认问题并安排修复。
- **Issue #4322: `NameError` 启动崩溃**：这是一个严重的阻拦性 BUG，已被标记为 OPEN 且无关联 PR，属于高优先级事件，需要立即介入排查。

</details>

<details>
<summary><strong>Hermes Agent</strong> — <a href="https://github.com/nousresearch/hermes-agent">nousresearch/hermes-agent</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，我将根据您提供的 Hermes Agent GitHub 数据，生成以下项目动态日报。

---

# Hermes Agent 项目动态日报 | 2026-06-13

## 1. 今日速览

今日 Hermes Agent 项目社区活跃度极高。在 24 小时内，社区共提交了 **12 个新 Issue**，但无任何 Issue 被关闭，且同时产生了 **50 个 Pull Request**，其中大部分（45 个）处于待合并状态。这反映出社区参与热情高涨，大量关于 Bug 修复、功能增强和平台适配的代码贡献正在等待审核。尽管没有新版本发布，但项目在 Matrix 加密、桌面端 TTS、Telegram 稳定性及新模型支持等多个关键领域都有显著的代码推进。当前项目状态可概括为：**“提交踊跃，审核吃紧”**，核心维护者需要加快 PR 合并速度以缓解积压。

## 2. 版本发布

无新版本发布。

## 3. 项目进展

今日共有 **5 个 PR** 被合并或关闭。虽然具体合并内容未在数据中详细列出，但结合近期活跃的 PR 趋势，可以推断项目在以下几个方面取得了关键进展：

- **安全性修复**：`#9833` 与 `#44753` 的持续推进，意味着核心的文件操作工具（如 `patch_tool`、`read_file`）的安全性和鲁棒性正在得到加强，这对于生产环境的部署至关重要。
- **稳定性提升**：多个针对不同平台和场景的修复 PR（如 Telegram 连接泄漏、Desktop 环境变量传递、Gateway 流控）正在密集提交，表明项目正在系统地解决已知的稳定性问题。

**项目整体前进方向**：Hermes Agent 正在从功能快速迭代阶段，转向一个更加注重 **稳定性、安全性、及平台兼容性** 的成熟阶段。尤其是对 Matrix、Telegram、Slack 等第三方平台适配的深度打磨，表明项目正在积极拓展其作为通用 AI 助手的应用场景。

## 4. 社区热点

今日讨论最活跃的 Issue/PR 主要集中在系统稳定性与配置兼容性上：

1.  **Issue #45508 [Hermes Agent 自我诊断‘自杀’且无法自动重启]**
    - **链接**: `NousResearch/hermes-agent Issue #45508`
    - **分析**: 这是一个严重的用户端问题，用户报告在更新到 v0.16.0 后，后端服务连接超时（15秒），导致桌面端 WebUI 虽然启动成功却无法正常工作。该 Issue 获得 1 条评论，反映了版本升级带来的兼容性阵痛，用户对此类“阻断性” Bug 反应强烈。

2.  **Issue #45506 [桌面语音模式使用默认配置文件的 TTS 配置]**
    - **链接**: `NousResearch/hermes-agent Issue #45506`
    - **分析**: 这是一个配置隔离性问题，当用户使用非默认 Profile 启动桌面端时，语音模式没有正确读取对应配置，而是“回退”到了默认配置。这反映出 Profile 功能在桌面端的实现尚不完善，影响了多用户/多配置场景下的使用体验。贡献者 `qdivan` 已经提交了修复 **PR #45522**，反应迅速。

3.  **Issue #45500 [Matrix: 文本消息绕过端到端加密]**
    - **链接**: `NousResearch/hermes-agent Issue #45500`
    - **分析**: 这是一项严重的安全 Bug，报告者指出在 Matrix 适配器中，附件传输能正确遵循加密逻辑，但文本消息路径却完全跳过了加密检查。在隐私意识日益增强的今天，该问题对 Matrix 平台用户构成重大安全风险。贡献者 `liuhao1024` 已提交修复 **PR #45518**，积极堵上安全漏洞。

## 5. Bug 与稳定性

今日报告的 Bug 中，按严重程度排列如下：

- **严重（Critical）**:
    - **#45508**: **服务启动后连接超时，导致前端不可用**。(无修复 PR)
        - 链接: `NousResearch/hermes-agent Issue #45508`

- **高（High）**:
    - **#45500**: **Matrix 文本消息绕过 E2EE 加密**。（已有修复 PR #45518）
        - 链接: `NousResearch/hermes-agent Issue #45500`
    - **#45496**: **`delegate_task` 子任务挂起导致父进程无响应**。（无修复 PR）
        - 链接: `NousResearch/hermes-agent Issue #45496`
    - **#45519**: **GLM-5.2 上下文长度检测错误，导致压缩阈值异常**。（无修复 PR，但有新模型支持 PR #45483）
        - 链接: `NousResearch/hermes-agent Issue #45519`

- **中（Medium）**:
    - **#18646**: **WhatsApp 群组消息路由错误**。（无修复 PR，是一个长期遗留 Bug）
        - 链接: `NousResearch/hermes-agent Issue #18646`
    - **#45506**: **桌面端 TTS 配置污染**。（已有修复 PR #45522）
        - 链接: `NousResearch/hermes-agent Issue #45506`
    - **#45493**: **Matrix 线程上下文中首条消息丢失**。（无修复 PR）
        - 链接: `NousResearch/hermes-agent Issue #45493`
    - **#37414**: **Codex 环境下执行输出未截断导致上下文溢出**。（无修复 PR）
        - 链接: `NousResearch/hermes-agent Issue #37414`

- **低（Low）**:
    - **#45499**: **API 网关调用时生成冗余告警日志**。（无修复 PR）
        - 链接: `NousResearch/hermes-agent Issue #45499`
    - **#45520**: **Linux VPS 无 GPU 环境下 Dashboard 无法渲染**。（无修复 PR）
        - 链接: `NousResearch/hermes-agent Issue #45520`

## 6. 功能请求与路线图信号

今日的功能请求信号主要来自新的 Pull Request，显示出以下开发方向：

1.  **新模型支持与上下文扩展 (已入 PR)**:
    - **PR #45483**: 贡献者 `potatogim` 提交了为 **Z.AI GLM-5.2** 模型提供支持的代码，该模型拥有 100 万 Token 的超大上下文窗口。这强烈暗示下一版本将支持该模型，以解决 #45519 中提到的上下文检测问题。

2.  **认证与凭证管理增强 (已入 PR)**:
    - **PR #45513**: 增加了 `hermes auth switch` 命令，允许用户在同类型凭证池中手动切换首选凭证。这表明项目在强化对企业级或者多账户场景的支持。

3.  **文件同步与平台集成 (长期待合并)**:
    - **PR #11460** & **#11459**: 这两个 PR 专注于为 Hermes Agent 增加 **Nextcloud** 文件双向同步和后台通知服务。尽管是 4 月提交的早期 PR，今日仍有更新，说明该项目维护者或社区成员仍在关注这个重要功能。这将是 Agent 环境感知能力的重要扩展。

## 7. 用户反馈摘要

从今日的 Issue 评论中可以提炼出以下用户痛点：

- **升级之痛**: 用户 `boosteryuge` 在 **#45508** 中描述了升级到 v0.16.0 后的完全无法使用的情况，这是一个典型的因升级导致的回退问题，严重影响用户体验和信任。
- **配置隔离不足**: 用户 `adamking77` 在 **#45506** 中抱怨使用非默认 Profile 时配置无法正确生效，这是对项目“Profile”功能在桌面端实现的质疑，表明其易用性和准确性有待提高。
- **子任务失控**: 用户 `sarv-projects` 在 **#45496** 中反馈，子任务（`delegate_task`）一旦卡死（如挂载、重试），父进程会完全挂起且无任何进度反馈，这是一个严重的可用性问题，让用户感觉 Agent 失去了控制。
- **平台适配细节**: 用户 `adaofeliz` 在 **#45500** 和 **#45493** 中连续指出了 Matrix 平台的加密和上下文处理问题，显示深度用户对特定平台的细节体验要求很高，而项目在这些细节上还存在不足。

## 8. 待处理积压

以下为长期未响应的重要 Issue 或 PR，需提醒维护者关注：

- **重要 Bug**:
    - **#18646**: **WhatsApp 群组路由错误** (创建于 2026-05-02): 这是一个已生成超过一个月的 P2 级别 Bug，至今未有关联的修复 PR，对于 WhatsApp 商业用户而言是严重障碍。
        - 链接: `NousResearch/hermes-agent Issue #18646`
    - **#15150**: **泰语翻译贡献** (创建于 2026-04-24): 社区用户 `nanobro` 提交了详细的文档翻译，但挂在 Issue 中长达近两个月未被合并。这不仅打击了贡献者的积极性，也不利于项目国际化。
        - 链接: `NousResearch/hermes-agent Issue #15150`

- **重要 PR (长期待合并)**:
    - **#11460** & **#11459**: **Nextcloud 集成** (创建于 2026-04-17): 这两个 P3 级别的特性 PR 由同一贡献者提交，且今日仍有更新，代表了社区对 Agent 和外部文件系统集成的强烈需求。长时间未合并可能会使贡献者失去耐心。
        - 链接: `NousResearch/hermes-agent PR #11460`
        - 链接: `NousResearch/hermes-agent PR #11459`

</details>

<details>
<summary><strong>PicoClaw</strong> — <a href="https://github.com/sipeed/picoclaw">sipeed/picoclaw</a></summary>

# PicoClaw 项目日报 — 2026-06-13

---

## 1. 今日速览

过去24小时项目活跃度较高：发布了一个自动化 nightly 版本，合并/关闭了3个 PR，同时有7个新 PR 处于待合并状态。社区讨论主要集中在 Evolution 功能下的异常 Token 消耗（Issue #3012）以及 WebSocket 协议中缺少明确结束信号（Issue #2984）。整体来看，项目在修复稳定性缺陷、扩展通道支持及完善协议信号方面取得了实质性进展，但仍有多个长期开放的功能性 PR 等待维护者评审。

---

## 2. 版本发布

**Nightly Build — v0.2.9-nightly.20260613.c362114c**  
- 这是一个每日自动构建版本，未包含破坏性变更或明确的迁移说明。  
- 建议仅用于测试环境，生产环境请等待稳定版本发布。  
- 完整变更日志： [v0.2.9...main](https://github.com/sipeed/picoclaw/compare/v0.2.9...main)

---

## 3. 项目进展

今日有3个 PR 被合并或关闭，标志着以下功能或修复已进入主分支：

| PR | 类型 | 描述 | 合并/关闭 |
|----|------|------|-----------|
| [#2551](https://github.com/sipeed/picoclaw/pull/2551) | refactor | 标准化通道标识，解耦通道名称与提供者类型（支持同一提供者多实例） | 已关闭 |
| [#3113](https://github.com/sipeed/picoclaw/pull/3113) | fix | 检查 `channel` 管理器中 JSON 序列化/反序列化错误（静默丢弃的问题） | 已关闭 |
| [#3112](https://github.com/sipeed/picoclaw/pull/3112) | fix | 修复工具循环中工具调用参数的 `json.Marshal` 错误被忽略的问题 | 已关闭 |

这些合并在健壮性（错误处理）和基础设施（通道架构）方面迈出了重要一步。

此外，还有7个待合并 PR 处于开放状态，涵盖以下重要特性：
- [#3117](https://github.com/sipeed/picoclaw/pull/3117) – 代理媒体走向图像模型（修复 Issue #3108）
- [#3118](https://github.com/sipeed/picoclaw/pull/3118) – 为 `picoclaw agent` 添加远程 WebSocket 模式
- [#3116](https://github.com/sipeed/picoclaw/pull/3116) – 补全 Pico 协议 `turn.done` 生命周期信号
- [#2964](https://github.com/sipeed/picoclaw/pull/2964) – 图像输入自动压缩功能
- [#2917](https://github.com/sipeed/picoclaw/pull/2917) – 新增 NEAR AI Cloud 提供商支持
- [#3063](https://github.com/sipeed/picoclaw/pull/3063) – DeltaChat 网关集成
- [#3115](https://github.com/sipeed/picoclaw/pull/3115) – 修复通用工具输出中内联 data URL 的媒体提取错误

---

## 4. 社区热点

**Issue #3012** — [Continuous consumption of tokens every minutes when evolution is enabled](https://github.com/sipeed/picoclaw/issues/3012)  
- 3条评论，0个赞。用户报告当启用 Evolution（进化模式）后，每分钟持续消耗 Token，即使没有新交互。影响环境为 FreeBSD + MiniMax 模型。该问题被标记为 `stale`，但昨日仍有更新，说明社区仍在关注。

**Issue #2984** — [Add explicit turn completion signal for Pico WebSocket clients](https://github.com/sipeed/picoclaw/issues/2984)  
- 2条评论，2个赞。该功能请求获得了较多社区认可。用户希望获得一个明确的 `turn.done` 事件，以便外部 WebSocket 客户端能精确判断代理是否已完整结束一轮处理。目前已有对应的修复 PR [#3116](https://github.com/sipeed/picoclaw/pull/3116) 被提出并处于待合并状态。

**Open PR #3117** — [fix(agent): route media turns to image models](https://github.com/sipeed/picoclaw/pull/3117)  
- 该 PR 针对 Issue #3108 中图像模型路由失败的问题，同时将 onboard workspace 从根目录嵌入。虽然尚无评论，但修复内容关乎视觉通道稳定性，预计将受到使用多模态模型的用户关注。

---

## 5. Bug 与稳定性

| 严重程度 | Issue / PR | 描述 | 当前状态 |
|----------|------------|------|----------|
| 高 | [#3012](https://github.com/sipeed/picoclaw/issues/3012) | 启用 Evolution 后每分鐘持续消耗 Token（疑似循环追问或轮询） | 开放，无 fix PR，标记为 stale |
| 中 | [#3115](https://github.com/sipeed/picoclaw/pull/3115) | 通用工具（如 `read_file`）返回的纯文本中若包含 `data:image/...;base64,...` 字符串，会被错误当作媒体附件，导致会话历史损坏 | 已提交 PR，未合并 |
| 低 | [#3113](https://github.com/sipeed/picoclaw/pull/3113) | channel 配置 JSON 序列化/反序列化错误被静默丢弃 | 已合并修复 |
| 低 | [#3112](https://github.com/sipeed/picoclaw/pull/3112) | ToolLoop 中工具调用参数的 JSON 序列化失败，导致调用数据丢失 | 已合并修复 |

> 建议维护者优先关注 #3012，该问题可能影响大量 Evolution 用户，且涉及计费资源消耗。

---

## 6. 功能请求与路线图信号

以下功能请求在今日数据中表现出较高社区需求或已有对应实现：

- **Pico WebSocket 明确回合结束信号** ([#2984](https://github.com/sipeed/picoclaw/issues/2984)) — 已由 PR #3116 实现，预计可随下次合并进入 nightly 版本。
- **远程代理模式** ([#3118](https://github.com/sipeed/picoclaw/pull/3118)) — 用户可通过 `--remote` 参数连接 WebSocket 服务端，为 remote agent 使用场景打开入口。
- **图像输入压缩** ([#2964](https://github.com/sipeed/picoclaw/pull/2964)) — 开放近三周，但无评论。支持多级配置压缩，降低视觉模型负载和 Token 成本。
- **NEAR AI Cloud 提供商** ([#2917](https://github.com/sipeed/picoclaw/pull/2917)) — 添加 OpenAI 兼容新 provider，适合 TEE 环境用户。
- **DeltaChat 网关** ([#3063](https://github.com/sipeed/picoclaw/pull/3063)) — 将 PicoClaw 接入 Decentralized Chat 生态，虽开放不久但功能较为完善。

综合来看，v0.2.9 之后的近期版本很可能会包含 WebSocket 信号完善、图像压缩以及新的外部集成提供者。

---

## 7. 用户反馈摘要

从今日 Issues 评论中提炼出以下用户声音：

- **Evolution 模式异常 Token 消耗** (#3012)：「即使没有用户输入，每分钟都会触发一次 token 消耗」。用户怀疑是 Evolution 的自动演变轮询导致的，建议增加触发频率控制或手动确认机制。
- **WebSocket 客户端缺乏确定性信号** (#2984)：「目前客户端能收到 `typing.start` / `typing.stop`，但没有明确的 ‘turn finished’ 事件，导致需要轮询猜时间」。用户希望获得类似 `turn.done` 的标准事件，以优化界面渲染和用户体验。

总体来看，用户对 PicoClaw 的实时性体验和资源控制有较高要求。

---

## 8. 待处理积压

以下 Issue 或 PR 已开放较长时间（超过2周），但尚未获得维护者响应或合并：

| 项目 | 类型 | 创建日期 | 最后更新 | 原因/备注 |
|------|------|----------|----------|-----------|
| [#2964](https://github.com/sipeed/picoclaw/pull/2964) | PR（Feat） | 2026-05-28 | 2026-06-12 | 图像压缩功能，无评论、无标签更新，可能被忽略 |
| [#2917](https://github.com/sipeed/picoclaw/pull/2917) | PR（Feat） | 2026-05-21 | 2026-06-12 | 新 provider 支持，无维护者 review |
| [#2984](https://github.com/sipeed/picoclaw/issues/2984) | Issue（Feature） | 2026-06-02 | 2026-06-12 | 已有关联 PR #3116，但 Issue 尚未关闭也未标记计划 |
| [#3012](https://github.com/sipeed/picoclaw/issues/3012) | Issue（Bug） | 2026-06-05 | 2026-06-13 | 虽标记 stale，但持续有更新，可能因复现环境特殊（FreeBSD）而未被优先处理 |

建议维护者优先 review #2964 和 #2917，这两个功能已具备完整实现，合并后可丰富 PicoClaw 的媒体处理能力和 provider 生态。

---

*数据来源：GitHub sipeed/picoclaw，截至 2026-06-13 23:59 UTC*

</details>

<details>
<summary><strong>NanoClaw</strong> — <a href="https://github.com/qwibitai/nanoclaw">qwibitai/nanoclaw</a></summary>

好的，这是为您生成的 NanoClaw 项目动态日报。

---

# NanoClaw 项目动态日报 | 2026-06-13

## 今日速览

今日项目活跃度表现强劲，尤其体现在 Pull Request 的合并与关闭上，共完成 10 个 PR，覆盖了从 Signal 原生集成、Ollama 多模态支持到关键 Bug 修复等多个方面，显示出开发节奏明显加快。Issues 方面，新提交的 Bug 修复 PR 与新增 Issue 均涉及核心稳定性与安全性问题，社区反馈质量较高。值得注意的是，今日无新版本发布，但大量 PR 的合并为下一版本积累了实质性内容。

## 版本发布

无

## 项目进展

今日共有 10 个 Pull Request 被合并或关闭，项目在以下几个关键领域取得了实质性的进步：

- **Signal 原生集成大幅增强**：系列 PR (#2040, #2070, #2071, #2203) 被合并，为 Signal 适配器添加了发出附件、接收任意文件类型附件以及双向消息反应（Reaction）的支持。这标志着 NanoClaw 在 Signal 渠道上的能力已趋于完整，能够处理更丰富的用户交互。
  - [PR #2203](nanocoai/nanoclaw PR #2203)
  - [PR #2071](nanocoai/nanoclaw PR #2071)
  - [PR #2070](nanocoai/nanoclaw PR #2070)
  - [PR #2040](nanocoai/nanoclaw PR #2040)

- **Ollama 支持多模态模型**：PR #2072 已合并。现在 `ollama_generate` 工具支持通过工作空间内路径传入图片，使 Agent 能够处理图片输入，显著扩展了与本地模型的交互能力。
  - [PR #2072](nanocoai/nanoclaw PR #2072)

- **系统稳定性与恢复机制完善**：PR #2670 修复了 Agent 因受损的会话转录而陷入崩溃循环的问题，引入了自我修复机制。同时，PR #2084 增加了每日项目备份与恢复功能，为灾难恢复提供了基础保障。
  - [PR #2670](nanocoai/nanoclaw PR #2670)
  - [PR #2084](nanocoai/nanoclaw PR #2084)

- **会话路由与分发逻辑改进**：PR #2277 和 PR #2267 分别修复了在后续消息中刷新路由上下文，以及 Agent 间（A2A）回复错误路由至最新会话而非原始会话的问题，保证了多会话场景下的对话一致性。
  - [PR #2277](nanocoai/nanoclaw PR #2277)
  - [PR #2267](nanocoai/nanoclaw PR #2267)

- **API 错误处理增强**：PR #2692 合并后，Agent 在面对临时性 5xx API 错误时将进行重试，仅在重试耗尽后通知用户，改善了在 LLM API 不稳定时的使用体验。
  - [PR #2692](nanocoai/nanoclaw PR #2692)

## 社区热点

今日最活跃的 Issue 是 **#2506**，讨论了 `send_message` 重复去重机制导致 Agent 响应被静默丢弃的问题。该 Issue 已存在一个月，有 3 条评论，但今日未出现新的讨论，目前仍处于开放状态。该问题描述了两种导致客户端超时的场景，是影响用户体验的关键 Bug，社区对此有较高关注度。
- [Issue #2506](nanocoai/nanoclaw Issue #2506)

## Bug 与稳定性

今日报告的 Bug 和稳定性问题较为集中，按严重程度排列如下：

- **严重**: #2751 (已关闭) - 当 LLM 预算耗尽时，Agent 会收到一个“伪造的”HTTP 200 响应，Agent SDK 将其视为正常成功响应处理，导致用户收不到任何回复。此问题已被识别并关闭，但未提及具体修复 PR。
  - [Issue #2751](nanocoai/nanoclaw Issue #2751)

- **严重**: #2506 (开放) - `send_message` 的去重机制在短时间内（<60秒）完成多个回合时，会错误地丢弃响应，导致客户端超时。目前尚无修复 PR。
  - [Issue #2506](nanocoai/nanoclaw Issue #2506)

- **严重**: #2668 (开放) - Agent 会话缺乏单次工具调用的超时机制。一个卡住的 MCP 工具会阻塞整个 Agent 会话长达 30 分钟，直到被强制杀死。此问题严重影响 Agent 的稳定性和响应速度。
  - [Issue #2668](nanocoai/nanoclaw Issue #2668)

- **中等**: #2711 (开放) - `create_agent` MCP 工具的权限控制存在漏洞，文档声称仅限管理员操作，但实际上任何容器都可以调用该工具创建新的 Agent 组，存在权限提升风险。
  - [Issue #2711](nanocoai/nanoclaw Issue #2711)

- **中等**: #2750 (待合并) - 此 PR 修复了在容器被杀死后 `outbound.db` 日志文件变为“陈旧”状态，以及轮询竞态条件的两个相关 Bug (#2516, #2640)。
  - [PR #2750](nanocoai/nanoclaw PR #2750)

## 功能请求与路线图信号

- **安全加固**：社区贡献者 `boazdori` 提交了两个关注安全性的 PR：`#2749` 要求通过 npm 安装的包需满足 3 天发布龄；`#2748` 建议对 Agent 容器进行安全加固，包括删除所有权限、禁止新权限生成、限制进程数。这表明社区和开发者已开始关注 Agent 容器的安全隔离与供应链安全，这些可能成为下一版本的安全基线。
  - [PR #2749](nanocoai/nanoclaw PR #2749)
  - [PR #2748](nanocoai/nanoclaw PR #2748)

- **Discord 附件支持**：PR #2752 旨在解决 Discord 渠道中附件（如文本和图片）无法被 Agent 读取的问题。这显示了社区对扩大平台支持范围的明确诉求。
  - [PR #2752](nanocoai/nanoclaw PR #2752)

- **v2 路线图澄清**：用户 `arthurkrupa` 在 Issue #2632 中询问了关于 v2 版本中 Telegram 多机器人（swarm）身份功能的明确状态，表明社区用户正在规划从 v1 向 v2 迁移，并期待官方对此功能的支持承诺。
  - [Issue #2632](nanocoai/nanoclaw Issue #2632)

## 用户反馈摘要

- **痛点**：多个 Issue 反映了 Agent 在特定场景下静默失败的问题（预算耗尽被忽略、短时间多轮对话响应被丢弃），这严重破坏了用户对 Agent 的信任。用户 `mshirel` 反馈，“Agent 静默丢弃响应，导致客户端超时”。
- **使用场景**：用户正积极探索 v2 的高级功能，如 Telegram 机器人 swarm 以及 MCP 工具的权限边界，显示出用户对生产级应用的部署需求。
- **改进方向**：PR #2753 修复了 `pnpm` 未在 PATH 中时 pre-commit 钩子失败的问题，虽然是一个小修补，但反映出对开发体验的打磨。同时，多个 PR 专注于安全性，表明用户对 Agent 沙箱的安全性有更高期待。

## 待处理积压

- **Issue #2506** (5月16日创建)：`send_message` 去重导致响应丢失的严重 Bug，至今无修复 PR，亟需关注。
  - [Issue #2506](nanocoai/nanoclaw Issue #2506)

- **Issue #2668** (6月1日创建)：Agent 会话因工具调用无超时而阻塞。这是一个影响深远的稳定性问题，建议尽快规划修复。
  - [Issue #2668](nanocoai/nanoclaw Issue #2668)

- **Issue #2632** (5月28日创建)：关于 v2 中 Telegram 多机器人功能的路线图请求，期待官方维护者给予明确回应。
  - [Issue #2632](nanocoai/nanoclaw Issue #2632)

- **PR #2750** (6月12日创建)：针对 `outbound.db` 日志损坏问题的修复，已标记为 “Fix”，建议尽快合并以解决关联的长期 Bug。
  - [PR #2750](nanocoai/nanoclaw PR #2750)

</details>

<details>
<summary><strong>NullClaw</strong> — <a href="https://github.com/nullclaw/nullclaw">nullclaw/nullclaw</a></summary>

好的，作为AI智能体与个人AI助手领域的开源项目分析师，以下是根据NullClaw项目在2026年6月13日的GitHub数据生成的动态日报。

---

# NullClaw 项目动态日报 | 2026-06-13

## 今日速览

- **活跃度评估**：项目处于中等活跃水平。过去24小时内无新版本发布，但社区提交了3个待合并的修复类PR，并有两项长期存在的Issue获得更新，表明维护和修复工作是当前重点。
- **重点动态**：核心功能`Agent`的定时任务（Cron Job）存在一个严重缺陷，导致子进程无法启动且消息无法投递，此问题已超过两周仍未被修复，用户呼声较高。
- **基础设施修复进行中**：多个来自贡献者`vernonstinebaker`的PR旨在提升系统配置灵活性和通道稳定性，特别是对Discord网关连接断开问题的修复，对提升用户体验至关重要。
- **新功能探索**：一项关于集成JIRA项目管理工具的功能请求持续获得关注，显示了用户对平台扩展至工作流自动化的明确需求。

## 版本发布

- **今日无新版本发布。**

## 项目进展

- **今日无合并或关闭的PR。**
- **待合并PR进展**：贡献者`vernonstinebaker`提交了3个重要的修复类PR，目前均处于开放状态，分别是：
    1.  **队列模式可配置化** ([PR #949](nullclaw/nullclaw PR #949))：旨在允许用户通过`config.json`配置文件为新的Agent会话设置默认队列模式，增加了系统的灵活性。
    2.  **抑制Agent失败时的错误日志误导** ([PR #951](nullclaw/nullclaw PR #951))：修复了当Agent子进程崩溃时，系统会将初始化阶段的日志错误地发送到用户频道的Bug，提升了错误报告的准确性。
    3.  **修复Discord网关连接恢复** ([PR #953](nullclaw/nullclaw PR #953))：解决了Discord通道在网关套接字关闭后无法稳定重连的问题，增加了心跳线程清理和连接状态健康检查，对保障聊天机器人服务的连续性至关重要。

## 社区热点

- **Issue #941：Agent定时任务无法执行** ([Issue #941](nullclaw/nullclaw Issue #941))：此问题目前是社区讨论的焦点。用户`weissfl`报告，当一个定时任务（`job_type: "agent"`）配置了通过Telegram发送消息时，任务状态会错误地标记为“完成”，但实际上Agent子进程根本没有启动，导致消息永远无法发送。该问题创建于5月31日，持续至今，严重影响了平台核心功能的可靠性。

## Bug 与稳定性

- **严重性：高 | Agent定时任务子进程未启动**：Bug [#941](nullclaw/nullclaw Issue #941) 描述了平台关键功能“定时Agent任务”的核心逻辑缺陷。任务被标记为完成但实际未执行，直接导致用户关键业务中断。目前尚无关联的修复PR。
- **严重性：中 | Agent失败时错误输出误导用户**：PR [#951](nullclaw/nullclaw PR #951) 针对一个用户感知层面的Bug —— 当Agent执行失败时，会将内部初始化日志（如内存计划、MCP寄存器信息）输出到用户频道，干扰用户判断。此PR是直接的修复方案。
- **严重性：中 | Discord网关断连后无法恢复**：PR [#953](nullclaw/nullclaw PR #953) 解决了Discord集成的一个稳定性问题。网关Socket关闭后，系统可能陷入无法恢复的无响应状态，影响基于Discord的Agent交互体验。

## 功能请求与路线图信号

- **JIRA集成工具** ([Issue #914](nullclaw/nullclaw Issue #914))：该功能请求已有近一个月，最新更新在今日。用户`sayjeyhi`提出了构建JIRA接入工具的需求，期望Agent能执行读取问题、创建工单等项目管理操作。这反映了用户希望将NullClaw Agent从个人助手扩展到团队协作和工作流自动化平台的强烈信号，有潜力成为下一个重要的功能方向。

## 用户反馈摘要

- 从Issue [#941](nullclaw/nullclaw Issue #941) 的讨论中可以提炼出用户**强烈的挫败感**：用户遵循了文档配置了所有参数，但系统表面的“成功”执行反馈与实际的功能缺失形成了巨大反差。这表明**任务执行状态的精确性和“最小可行性执行”的验证**对用户信任至关重要。
- 用户**对通道稳定性的担忧**：尽管没有直接评论，但针对Discord网关的修复PR [#953](nullclaw/nullclaw PR #953) 暗示了用户在使用聊天通道集成Agent时可能遇到了连接中断后无法恢复的痛点，影响了Agent作为“始终在线”助手的体验。

## 待处理积压

- **Issue #941：Agent定时任务未执行**：作为当前最严重且创建时间较长的Bug（6月3日至今），该项目应被列为最高优先级。项目维护者需要尽快响应，确认问题根因，并与社区沟通修复计划或临时工作区方案。
- **Issue #914：JIRA集成功能请求**：已开放一个月，尽管不是错误，但其活跃度和潜在价值使其成为需要纳入路线图讨论的重要事项。长期无官方回应可能会挫伤贡献者`sayjeyhi`的积极性。
- **PR #949, #951, #953**：这三个PR已经开放3-4天，且都涉及核心功能和稳定性。项目维护者应尽快进行代码审查，评估并合并，以消除用户痛点并保持社区的贡献热情。

</details>

<details>
<summary><strong>IronClaw</strong> — <a href="https://github.com/nearai/ironclaw">nearai/ironclaw</a></summary>

好的，作为 AI 智能体与个人 AI 助手领域开源项目分析师，以下是根据您提供的 IronClaw 项目数据生成的 2026-06-13 项目动态日报。

---

## IronClaw 项目动态日报 — 2026-06-13

### 1. 今日速览

过去 24 小时，IronClaw 项目展现出极高活跃度，核心围绕 **“附件支持 (Attachments)”** 和 **“DeferredBusy 消息排空机制”** 两个主题展开深度开发。PR 流动量巨大（50 条），核心贡献者主导了多项大尺寸（XL）PR 的提交和审核。系统稳定性方面，`#4817` 等长线追踪 issue 显示核心开发团队正对模块进行严格的设计回顾，并规划后续优化。总体而言，项目正处于密集的功能开发和架构优化阶段，健康度极高。

### 2. 版本发布

**无新版本发布。**

### 3. 项目进展

今日项目在多个功能管线上取得重大进展，多件核心 PR 被合并关闭，标志着关键里程碑的达成。

-   **附件支持（Attachments）管线**：这是当前最活跃的开发线。多件 PR 于今日被合并：
    -   `#4655` [CLOSED] 实现了附件引用（`AttachmentRef`）在 Reborn 对话记录（transcript）合约中的持久化，结束了文本-only 的时代。这是附件功能的核心基础。
    -   `#4654` [CLOSED] 创建了统一的附件格式注册表（`ironclaw_common`），消除了此前分散在多处的硬编码列表，从根源上解决格式支持漂移的 bug。
    -   `#4812` [CLOSED] 完成了 `DeferredBusy` 消息的自动排空（drain）功能，解决了线程被门控阻塞时，后续消息被永久吞没的问题，优化了用户体验。
-   **后续规划**：在 `#4812` 合并后，核心开发者迅速开启了 `#4817` 等系列 Issue，对 `DeferredBusy` drain 机制进行架构回顾。虽然已发布的版本是“照原样运行”（ships as-is），但团队已规划了**信任重新提交 seam**、**过时意图策略**和**启动时扫描**等后续优化，体现了严谨的开发态度。

### 4. 社区热点

今日社区热点主要集中在对关键设计决策的深入讨论和后续规划上。

-   **Issue `#4817` [OPEN]：`DeferredBusy` 排空机制后续追踪**
    -   **链接**: [Issue #4817](https://github.com/nearai/ironclaw/issues/4817)
    -   **讨论热度**: 3 条评论
    -   **分析**: 此 issue 并非报告 Bug，而是核心开发者在合并 `#4812` 后，主动发起的设计回顾清单。它列出了三个未在当前 PR 中解决的架构决策（`trusted-resubmit seam`、`stale-intent policy`、`startup sweep`）。这表明社区讨论的核心是**“如何让新功能更健壮、更优雅”**，而非简单的功能请求或 Bug 反馈，反映了项目的极高水平。

-   **PR `#4668` [OPEN] / `#4670` [OPEN] / `#4676` [OPEN] 等系列附件 PR**
    -   **链接**: [PR #4668](https://github.com/nearai/ironclaw/pull/4668), [PR #4670](https://github.com/nearai/ironclaw/pull/4670), [PR #4676](https://github.com/nearai/ironclaw/pull/4676) 等
    -   **讨论热度**: 系列 PR，持续更新
    -   **分析**: 由核心贡献者 `ilblackdragon` 主导的附件支持是一个庞大的“Track”系列。今日多条 PR 处于活跃状态，讨论集中在如何将字节存储、文本提取、上下文折叠等环节有效串联。这反映了社区对**多模态交互**和**文件处理能力**的强烈需求，是项目未来差异化竞争力的关键。

### 5. Bug 与稳定性

今日报告和处理的 Bug 主要集中在用户交互流程和系统可靠性上。

-   **高严重性（稳定性）**：
    -   `#4108` [OPEN]：**持续存在的 Nightly E2E 测试失败**。这表示主干分支可能长期存在未被发现的回归问题。**尚无 Fix PR**，项目维护者需优先关注。
-   **中严重性（用户体验）**：
    -   `#4839` [OPEN] / `#4840` [OPEN]：**Slack 重新授权循环** (Re-approval loop)。问题在于当能力调用同时需要一次性审批和长期凭证时，用户会陷入重复授权。这是典型的认证流程时序问题。**已有 Fix PR** (`#4839`, `#4840`)，核心维护者 `henrypark133` 已提交修复方案。
    -   `#4777` [OPEN]：**Slack 连接状态 UI 未正确显示**，导致用户误以为 Slack 未连接而触发重连循环。**已有 Fix PR**，通过查询真实连接状态而非静态显示来解决。

### 6. 功能请求与路线图信号

-   **运行时上下文感知**：`#4828` [OPEN] 提出了一个多功能请求，希望模型能“感知”到当前的连接通道、交付状态和运行来源。这旨在解决当 Slack 连接完成后，模型无法知晓该连接的 bug。**已有 Fix PR (`#4836`)**，状态为 OPEN，表明该功能很可能在下一版本中实现。
-   **“始终允许”审批持久化**：`#4825` [OPEN] 用户明确要求，在一次对话中选择“始终允许”（always allow）审批后，此设置应跨线程持久化。这是一个高价值的用户体验增强请求，核心开发者已跟进，并可能推动 `#4825` 成为下一版本的关键特性。

### 7. 用户反馈摘要

从 Issues 和 PR 的描述中提炼出的用户痛点：

-   **繁琐的重复操作**：用户对跨线程的重复授权审批感到沮丧（`#4825`）。用户期望系统能记住自己的偏好。
-   **不直观的错误反馈**：当线程被门控阻塞时，新消息被静默吞没或状态显示异常（如 Slack 连接显示为“断开”），使用户感到困惑并采取不必要的重试操作（`#4777`, `#4838`）。用户期望系统能给出明确、可操作的反馈。
-   **认证流程混乱**：在权限和凭证分离的场景下，用户被要求批准一个尚不能执行的操作，浪费了审批行为。用户期望系统在请求批准前先解决凭证问题（`#4840`）。

### 8. 待处理积压

-   **高优先级**：
    -   `#4108` [OPEN]：**Nightly E2E 测试持续失败**（创建于 5 月 27 日）。此 Issue 已超过两周未得到解决，是评估项目代码稳定性的一个关键风险点。**需要维护者紧急响应。**
-   **长线跟进**：
    -   `#3708` [OPEN]：**发布流程 PR**（创建于 5 月 16 日）。该 PR 已打开近一个月，阻塞了多个 crate 的版本发布。可能涉及复杂的依赖冲突或发布流程问题，需要项目核心团队做决策以推动发布。

</details>

<details>
<summary><strong>LobsterAI</strong> — <a href="https://github.com/netease-youdao/LobsterAI">netease-youdao/LobsterAI</a></summary>

# LobsterAI 项目动态日报 — 2026-06-13

## 今日速览
过去24小时内，项目共更新4条Issue和8条Pull Request，无新版本发布。所有活跃Issue均为4月3日创建的“stale”状态，今日因维护者回复或系统自动更新而出现活跃信号。3个PR在昨日被合并至主分支（#2158、#2156、#2157），涵盖Computer Use MVP、实时语音输入、图片保存修复等核心功能。整体来看，项目处于**稳定开发与维护期**，但社区反馈的多个历史Bug尚未被修复，活跃度中等偏下。

## 版本发布
无

## 项目进展
昨日（6月12日）有3个PR被合并，标志着项目向前迈进了重要一步：

- **#2158 [CLOSED] chore(release): merge release/2026.6.11 into main**  
  合并了2026.6.11版本的发布分支，带来三项重量级特性：
  - **Computer Use MVP** 及内建Computer Use工具包
  - 实时ASR语音输入（cowork提示框）
  - HTML Artifact公开分享模式选择、图片/SVG Artifact分享支持  
  [netease-youdao/LobsterAI PR #2158](https://github.com/netease-youdao/LobsterAI/pull/2158)

- **#2156 [CLOSED] fix(computer-use): bump runtime to 1.0.7**  
  升级Computer Use运行时至1.0.7，修复了辅助进程异常退出的诊断问题。  
  [netease-youdao/LobsterAI PR #2156](https://github.com/netease-youdao/LobsterAI/pull/2156)

- **#2157 [CLOSED] fix(media): 修正文生图保存图片的扩展名**  
  优先根据文件字节识别真实格式，修正服务端返回文件名中错误后缀（如PNG被存为.jpg），并补充回归测试。  
  [netease-youdao/LobsterAI PR #2157](https://github.com/netease-youdao/LobsterAI/pull/2157)

此外，两个较早的stale PR于今日关闭：

- **#1466 [CLOSED] fix(mcp): modal close button unreachable when content grows tall**  
  修复MCP表单弹窗因overflow-y-auto导致关闭/Cancel按钮不可点击的问题。  
  [netease-youdao/LobsterAI PR #1466](https://github.com/netease-youdao/LobsterAI/pull/1466)

- **#1467 [CLOSED] fix(shortcuts): display Cmd (⌘) instead of Ctrl on macOS**  
  修复macOS上快捷键设置面板未正确显示⌘键的问题。  
  [netease-youdao/LobsterAI PR #1467](https://github.com/netease-youdao/LobsterAI/pull/1467)

## 社区热点
今日所有Issue和PR的评论数均不超过2条，点赞数为0，整体讨论热度较低。相对而言，**#1443 关于openclaw新版本兼容性** 的问题获得2条评论，是唯一存在用户交互的Issue：

> “之前升级到openclaw v2026.3.24，官方的release说明有breaking change，我本地尝试了会报错，拉不起来。我们有计划适配openclaw新版本吗？”

该问题反映了用户在跟进上游依赖变更时遇到的障碍，社区关注点集中在外围生态兼容性上。

[netease-youdao/LobsterAI Issue #1443](https://github.com/netease-youdao/LobsterAI/issues/1443)

## Bug 与稳定性
以下为今日记录的未修复Bug，按严重程度排列：

| 严重程度 | 问题描述 | Issue | 状态 | 是否有fix PR |
|----------|----------|-------|------|--------------|
| 🔴 严重 | 已停用的技能在对话中仍然可以被调用，可能引发安全或行为异常 | [#1439](https://github.com/netease-youdao/LobsterAI/issues/1439) | OPEN | 无 |
| 🟠 中 | 创建定时任务时，选择“不重复”并清空日历后，点击“创建任务”按钮无反应，无报错提示 | [#1437](https://github.com/netease-youdao/LobsterAI/issues/1437) | OPEN | 无 |
| 🟡 低 | Agent添加技能后，对话中引用的技能不展示，重新切换Agent才恢复 | [#1442](https://github.com/netease-youdao/LobsterAI/issues/1442) | OPEN | 无 |

另外，两个已关闭的Bug修复（#1466、#1467）属于界面体验类问题，已通过合并PR解决。

## 功能请求与路线图信号
- **#1443 支持openclaw新版本** —— 用户请求适配openclaw的breaking change。目前项目仍依赖旧版，若上游强制升级则可能影响部署。该请求是项目未来的兼容性隐患。
- **#1440 feat(cowork): 将已选技能标签移至输入框内顶部展示** —— 由内部开发者提交的UI改进PR，解决技能标签与工具栏按钮混排导致的拥挤问题，目前仍为OPEN状态。
- **#1441 feat(artifacts): add extensible preview pipeline for HTML, React and Mermaid** —— 扩展Artifact预览管道的PR，支持HTML、React、Mermaid渲染，来自社区贡献，同样是OPEN状态。

这两项PR与近期发布的Artifact分享功能高度相关，可能被纳入下一版本迭代。

[netease-youdao/LobsterAI PR #1440](https://github.com/netease-youdao/LobsterAI/pull/1440)  
[netease-youdao/LobsterAI PR #1441](https://github.com/netease-youdao/LobsterAI/pull/1441)

## 用户反馈摘要
从仅有的2条Issue评论中，可以提炼以下用户痛点：

1. **依赖升级断裂**：“升级到openclaw v2026.3.24后报错拉不起来”——用户尝试跟随上游更新，但项目未及时适配，导致使用受阻。
2. **技能管理逻辑混乱**：“关闭技能后对话仍可调用”“添加技能后对话不展示引用”——表明技能启用/禁用状态与运行时行为不一致，用户对技能系统的工作机制存在疑问（“agent选择技能的作用是什么？只触发选择的技能？”）。
3. **重复导入无校验**（来自PR #1445描述）：“上传zip/文件夹/GitHub导入时遇到同名目录自动追加-1后缀，导致重复技能同时注入System Prompt”——影响模型路由稳定性，开发者已提交修复PR。

## 待处理积压
以下Issue和PR自4月3日创建后已停滞超过70天，至今无实质性进展，提醒维护者关注：

| 类型 | 编号 | 标题 | 创建时间 | 最后更新 | 链接 |
|------|------|------|----------|----------|------|
| Issue | #1443 | 有计划支持新版本的openclaw吗？ | 2026-04-03 | 2026-06-13 | [Link](https://github.com/netease-youdao/LobsterAI/issues/1443) |
| Issue | #1437 | 创建定时任务时点击创建无反应 | 2026-04-03 | 2026-06-13 | [Link](https://github.com/netease-youdao/LobsterAI/issues/1437) |
| Issue | #1439 | 上传技能已停用，对话中仍然可以调用 | 2026-04-03 | 2026-06-13 | [Link](https://github.com/netease-youdao/LobsterAI/issues/1439) |
| Issue | #1442 | Agent添加技能对话后引用不展示 | 2026-04-03 | 2026-06-13 | [Link](https://github.com/netease-youdao/LobsterAI/issues/1442) |
| PR | #1440 | feat(cowork): 已选技能标签移至输入框内顶部展示 | 2026-04-03 | 2026-06-13 | [Link](https://github.com/netease-youdao/LobsterAI/pull/1440) |
| PR | #1441 | feat(artifacts): 扩展HTML/React/Mermaid预览管道 | 2026-04-03 | 2026-06-13 | [Link](https://github.com/netease-youdao/LobsterAI/pull/1441) |
| PR | #1445 | fix(skills): 修复技能重复导入及zip目录名问题 | 2026-04-03 | 2026-06-13 | [Link](https://github.com/netease-youdao/LobsterAI/pull/1445) |

上述积压项涵盖功能需求、Bug修复和UI改进，建议项目维护者在下一个开发冲刺中优先评估其优先级。

</details>

<details>
<summary><strong>TinyClaw</strong> — <a href="https://github.com/TinyAGI/tinyagi">TinyAGI/tinyagi</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>Moltis</strong> — <a href="https://github.com/moltis-org/moltis">moltis-org/moltis</a></summary>

# Moltis 项目日报 | 2026-06-13

---

## 1. 今日速览

- 过去24小时项目活跃度较低：新增2条Issue，无Pull Request更新，无版本发布。
- 新提交的Issue涵盖1个Bug（Fastmail MCP授权问题）和1个功能请求（Kubernetes原生沙箱后端），社区讨论尚未深入，每条Issue仅1-2条评论。
- 无已合并或关闭的PR，项目整体推进停滞，依赖后续维护者对上述Issue的响应。
- 开发与维护节奏偏缓，建议关注新Issue中可能影响Agent安全性与扩展性的核心诉求。

---

## 2. 版本发布

（无新版本发布，本节略）

---

## 3. 项目进展

- **合并/关闭的PR**：今日无任何PR被合并或关闭，代码库无增量变更。
- **待办状态**：目前所有开放PR仍处于待处理阶段，无可见的功能推进或缺陷修复落地。

---

## 4. 社区热点

### #1115 [Bug] Fastmail MCP Authorisation
- **链接**：[Issue #1115](https://github.com/moltis-org/moltis/issues/1115)
- **评论数**：2 | 👍：0
- **分析**：用户报告在Fastmail MCP（Model Context Protocol）集成时出现授权问题。虽然评论数不高，但该Issue直接关系到Moltis连接到外部邮件服务的能力，属于集成兼容性关键问题。社区尚未给出临时解决方案，维护者应优先排查认证流程或提供环境配置建议。

### #1118 [Feature] Kubernetes-native sandbox backend with runtimeClassName support
- **链接**：[Issue #1118](https://github.com/moltis-org/moltis/issues/1118)
- **评论数**：1 | 👍：0
- **分析**：提出增加Kubernetes原生的沙箱后端，利用`runtimeClassName`支持Kata Containers、gVisor等VM级隔离。这是对Agent执行环境安全性的重要扩展，尤其针对LLM生成的非可信代码执行场景。尽管当前评论少，但该需求符合云原生趋势，可能成为未来版本的核心特性之一。

---

## 5. Bug 与稳定性

| 严重程度 | Issue | 摘要 | 是否有修复PR |
|---------|-------|------|-------------|
| 中 | [#1115](https://github.com/moltis-org/moltis/issues/1115) | Fastmail MCP授权失败，影响邮件服务集成 | 无 |
| 低 | — | 无其他崩溃或回归报告 | — |

**说明**：当前仅有一个Bug，未提供完整日志或复现步骤，建议维护者引导用户补充会话上下文以加速定位。

---

## 6. 功能请求与路线图信号

| 功能 | Issue | 分析 | 是否可能纳入下一版本 |
|------|-------|------|-------------------|
| Kubernetes原生沙箱后端 | [#1118](https://github.com/moltis-org/moltis/issues/1118) | 提供VM级隔离，适合企业级部署与多租户场景；需新增Kubernetes client依赖及Pod生命周期管理 | ⚠️ 可能性中等，需评估实现成本与现有Docker沙箱的互补性 |

**额外信号**：当前无其他新功能请求。该特性若被采纳，将显著增强Moltis在生产环境中的可信执行能力，与行业安全合规趋势一致。

---

## 7. 用户反馈摘要

- **#1115 评论中的隐含痛点**：用户可能已按照文档配置Fastmail MCP但失败，说明文档或默认配置存在盲区；用户未提供完整session context，暗示排查门槛较高。
- **#1118 提出的场景**：用户希望能在Kubernetes集群中直接运行Agent沙箱，避免额外依赖Docker daemon，并利用现有OCI runtime实现更强隔离。这意味着部分用户已开始将Moltis部署在容器编排环境中，对原生K8s集成的呼声开始显现。
- **整体满意度**：今日无正面或负面满意度信号，社区活跃度低可能反映用户对当前版本功能相对满意，或正处于功能等待期。

---

## 8. 待处理积压

当前无长期未响应的重大Issue或PR。所有开放Issue均为近48小时内新提交，但维护者尚未回复。建议：

- **尽快关注 #1115**：授权类Bug易引发用户流失，可先提供临时Workaround或要求补充日志。
- **对 #1118 给出初步态度**：即使暂不实现，也应回复设计思路或收集更多使用场景，避免用户感到被忽略。

> *注意：数据截止2026-06-13 12:00 UTC，项目动态存在延迟可能。*

</details>

<details>
<summary><strong>CoPaw</strong> — <a href="https://github.com/agentscope-ai/CoPaw">agentscope-ai/CoPaw</a></summary>

# CoPaw 项目动态日报 (2026-06-13)

> 数据来源：GitHub [`agentscope-ai/QwenPaw`](https://github.com/agentscope-ai/QwenPaw)（CoPaw 开源项目仓库）  
> 统计时段：2026-06-12 至 2026-06-13

---

## 1️⃣ 今日速览

过去 24 小时项目保持中等活跃度，共计 12 条 Issues 更新（新开/活跃 11 条、关闭 1 条）和 4 条 PR 更新（待合并 3 条、合并/关闭 1 条）。无新版本发布。社区反馈集中在 **定时任务触发异常、Tauri 桌面端启动性能退化、长对话卡死** 等稳定性问题，同时用户对新渠道（Zalo Bot）和集成体验（Kimi for Coding、Feishu 流式卡片优化）提出了明确需求。值得注意的是，1 个回归 Bug（Gemini 工具调用在 `v1.1.11.post2` 失效）引起团队关注，已有重现步骤。PR 方面，一个带有 **12 个 BI 技能的数据分析插件 `DataPaw`** 正在等待合并，呈现出生态扩展的积极信号。

---

## 2️⃣ 版本发布

**无**。过去 24 小时未发布任何新版本或预发布版本。

---

## 3️⃣ 项目进展

今日 **1 个 PR 被合并/关闭**，3 个 PR 仍处于开放审查状态。

### ✅ 已合并/关闭 PR

| PR | 描述 | 作者 | 状态 |
|----|------|------|------|
| [#4969 feat(skill): Add skill tag batch download](https://github.com/agentscope-ai/QwenPaw/pull/4969) | 支持按标签批量下载 Skill 到工作区 | Leirunlin | ✅ 合并 |

该 PR 自 6 月 5 日创建，历时 8 天完成合并，对应 Issue #2961，完善了 Skill 管理功能。

### 🔄 待合并 PR

| PR | 描述 | 作者 | 备注 |
|----|------|------|------|
| [#5088 feat: initial governance & sandbox interface disscussion](https://github.com/agentscope-ai/QwenPaw/pull/5088) | 治理与沙箱接口初始讨论 | weidankong | 进展较慢，已 3 天未更新 |
| [#5069 feat(agents): add visual model fallback for text-only primary models](https://github.com/agentscope-ai/QwenPaw/pull/5069) | 为仅文本模型增加视觉模型后备方案 | yuanxs21 | 可提升多模态处理鲁棒性 |
| [#4622 plugin(datapaw): add data-analysis plugin with 12 BI skills](https://github.com/agentscope-ai/QwenPaw/pull/4622) | 新增数据分析插件 DataPaw，内含 12 个 BI 技能 | EliasMei | **首次贡献者**，插件生态重要推进 |

项目整体向前推进了 **Skill 批量下载** 功能，同时在 **沙箱安全、视觉模型后备、数据分析插件** 3 个方向有明确进展。

---

## 4️⃣ 社区热点

今日 **评论最多** 的 Issue 反映了用户对核心功能的强烈关注：

- **🔥 [#5064 - 定时任务无法正常触发](https://github.com/agentscope-ai/QwenPaw/issues/5064)**  
  💬 11 条评论，热度最高。用户 `tina0501853` 指出 Agent 创建的定时任务在 `v1.1.10` 中能正常生成，但到达时间后不触发，且无法手动编辑。该问题直接影响用户的工作流自动化和日常使用，社区讨论集中在可能的原因（任务调度器、权限模型等），但尚未有明确方案。

- **🔥 [#5140 - v1.1.11.post2 附件下载 404 (已关闭)](https://github.com/agentscope-ai/QwenPaw/issues/5140)**  
  💬 6 条评论。用户报告纯文本可下载但 docx/pdf 报 404，该问题今日已关闭，预计已在开发分支修复。

- **🔥 [#5047 - Windows Tauri 桌面端启动极慢](https://github.com/agentscope-ai/QwenPaw/issues/5047)**  
  💬 3 条评论，但代表广泛痛点：从 Python 迁移到 Tauri 后启动时间从 1-2 分钟恶化到十几分钟，且在 Windows 11 上频繁无响应。用户 `moolawooda` 给出了详细的环境信息（CPU i7-11800H, 32GB RAM），排除了硬件瓶颈。

这些热点显示社区 **对稳定性和基础体验（定时、启动、附件）** 的诉求最为迫切。

---

## 5️⃣ Bug 与稳定性

今日报告的 Bug 按严重程度排列如下：

| 严重性 | Issue | 描述 | 状态 | 是否有关联修复 PR |
|--------|-------|------|------|-------------------|
| 🔴 严重回归 | [#5163 Gemini 工具调用失效](https://github.com/agentscope-ai/QwenPaw/issues/5163) | `v1.1.10` 正常，`v1.1.11.post2` 失败，Gemini 无法使用内置工具 | 开放，已确认回归版本区间 | 无 |
| 🔴 严重 | [#5064 Agent 定时任务不触发](https://github.com/agentscope-ai/QwenPaw/issues/5064) | 任务生成无报错，但到达时间不执行，且无法手动编辑 | 开放 | 无 |
| 🟡 中等 | [#5162 对话思考逻辑死循环](https://github.com/agentscope-ai/QwenPaw/issues/5162) | 长对话或特定提示词下进入无限思考 | 开放 | 无 |
| 🟡 中等 | [#5140 附件下载 404 (docx/pdf)](https://github.com/agentscope-ai/QwenPaw/issues/5140) | 已关闭，推测已修复 | ✅ 已关闭 | 可能已合并 |
| 🟡 中等 | [#5166 Python 3.13 安装 TeamChat 插件失败](https://github.com/agentscope-ai/QwenPaw/issues/5166) | 报错 `No module named 'imghdr'`（Python 3.13 移除） | 开放 | 无 |
| 🟢 低 | [#5165 打包后白屏](https://github.com/agentscope-ai/QwenPaw/issues/5165) | `build_win_pyinstaller.ps1` 引用不存在的模块导致 exe 白屏 | 开放 | 无 |

**关键问题**：`#5163` 确认了 `v1.1.11.post2` 引入了 Gemini 通道的回归，建议维护者优先定位。`#5064` 和 `#5162` 影响 Agent 自动化和对话可靠性，属于功能级 Bug。

---

## 6️⃣ 功能请求与路线图信号

| 功能需求 | Issue | 用户诉求 | 潜在纳入版本信号 |
|----------|-------|----------|------------------|
| **Zalo Bot 渠道支持** | [#5168](https://github.com/agentscope-ai/QwenPaw/issues/5168) | 越南最流行消息平台，连接 Agent 到 Zalo | 无现有 PR，但社区呼声高 |
| **Kimi for Coding / uv 白名单** | [#5156](https://github.com/agentscope-ai/QwenPaw/issues/5156) | 已订阅 Kimi coding 套餐的用户希望直接集成 | 属于 Provider 扩展，无对应 PR |
| **桌面端系统托盘、开机自启、后台常驻** | [#5164](https://github.com/agentscope-ai/QwenPaw/issues/5164) | 完善桌面应用的后台服务管理能力 | 无对应 PR |
| **飞书 CardKit 流式卡片长回复优化** | [#5167](https://github.com/agentscope-ai/QwenPaw/issues/5167) | 长回复时卡片刷新变慢，字字输出体验差 | 社区提供了竞品对比（非流式+分段更新），建议优化流式推送间隔 |
| **数据插件 DataPaw** | PR [#4622](https://github.com/agentscope-ai/QwenPaw/pull/4622) | 新增 12 个 BI 技能的数据分析插件 | **已验证通过**，等待合并，下一版本有望纳入 |

**路线图信号**：`DataPaw` 插件是近期最明确的功能扩展，体现了项目生态化演进。桌面端和渠道扩展（Zalo）的需求表明项目正在覆盖更多地区与使用场景。

---

## 7️⃣ 用户反馈摘要

从 Issues 评论及描述中提取的真实用户声音：

- **“定时任务生成正常，但到点不触发，也不能手动编辑，等于废了。”**  
  —— `tina0501853` 对 #5064 的反馈，情绪显露出挫败感，期待尽快修复。

- **“自从换成 Tauri，启动速度从一两分钟变成十几分钟，而且经常无响应。”**  
  —— `moolawooda` 在 #5047 中的描述，对桌面端迁移后的性能退化表示困扰。

- **“长对话后直接卡死，不回复任何内容，只能关闭重开。”**  
  —— `tecgic` 在 #5161 中反馈，影响重度用户的持续使用。

- **“已经付费订阅了 Kimi coding 套餐，但 QwenPaw 只支持官方 API，导致我的套餐能力浪费了。”**  
  —— `wjt0321` 在 #5156 中表达了对订阅权益无法迁移的痛点。

- **“流式卡片长回复时一个字一个字地吐，体验还不如非流式分段更新。”**  
  —— `wjt0321` 在 #5167 中基于实际使用体验提出优化建议，态度建设性。

- **“我是越南用户，非常希望 Zalo 能和 QwenPaw 打通。”**  
  —— `lamnguyen3119` 在 #5168 中为新渠道发声，体现国际化社区需求。

整体上，用户对项目功能深度表示认可，但对 **稳定性**（定时、卡顿、回归）和 **体验细节**（启动速度、流式刷新）的期待较高。

---

## 8️⃣ 待处理积压

以下为长期未响应或进展缓慢的重要 Issue/PR，建议维护者优先关注：

| 事项 | 创建时间 | 最后更新 | 重要性 | 原因 |
|------|----------|----------|--------|------|
| [#5047 Windows Tauri 启动极慢](https://github.com/agentscope-ai/QwenPaw/issues/5047) | 06-09 | 06-13 (仅评论未解决) | 🔴 高 | 严重影响桌面端迁移用户体验 |
| [#5064 Agent 定时任务不触发](https://github.com/agentscope-ai/QwenPaw/issues/5064) | 06-10 | 06-12 | 🔴 高 | 核心 Agent 能力缺陷，11 条评论无团队回复 |
| [#5163 Gemini 工具调用回归](https://github.com/agentscope-ai/QwenPaw/issues/5163) | 06-12 | 06-12 | 🔴 高 | 1 天内无维护者确认，已明确回归版本区间 |
| PR [#5088 治理与沙箱接口](https://github.com/agentscope-ai/QwenPaw/pull/5088) | 06-10 | 06-12 | 🟡 中 | 3 天未更新，无人审查 |
| PR [#4622 DataPaw 插件](https://github.com/agentscope-ai/QwenPaw/pull/4622) | 05-22 | 06-12 | 🟡 中 | 首次贡献者，等待合并已超过 3 周 |

> 建议：尽快对 `#5047`、`#5064`、`#5163` 给出官方回应或临时方案，对 `#4622` 进行代码审查并合并，以保持贡献者积极性。

---

**总结**：CoPaw 项目在功能扩展上保持良好节奏（Skill 下载、DataPaw 插件），但近期 Bug 回归和性能退化问题较为集中，社区对稳定性的关注度上升。建议下一阶段以 **“稳定优先”** 为基调，重点修复定时任务、桌面端启动和 Gemini 回归问题，同时加快 DataPaw 等生态插件的合并节奏。

*— 由 AI 智能体分析师自动生成，仅供参考。*

</details>

<details>
<summary><strong>ZeptoClaw</strong> — <a href="https://github.com/qhkm/zeptoclaw">qhkm/zeptoclaw</a></summary>

过去24小时无活动。

</details>

<details>
<summary><strong>ZeroClaw</strong> — <a href="https://github.com/zeroclaw-labs/zeroclaw">zeroclaw-labs/zeroclaw</a></summary>

好的，作为AI智能体与个人AI助手领域开源项目分析师，以下是基于您提供的ZeroClaw (zeroclaw-labs/zeroclaw) GitHub数据生成的2026年6月13日项目动态日报。

---

# ZeroClaw 项目动态日报 | 2026-06-13

## 今日速览

ZeroClaw 项目今日保持高活跃度，共处理了50条 Pull Request (PR)，其中26条已合并/关闭，显示出项目团队高效的代码审查与合并节奏。同时，社区在Bug反馈和功能增强方面也保持活跃。值得注意的是，今日有多个高风险、大规模的PR处于开放状态，涉及核心运行时引擎重构、技能系统大改和新的记忆模式，预示着项目正酝酿重大架构演进。Issues方面，今日新开2条，关闭3条，整体健康度良好。

## 版本发布

**无新版本发布。**

## 项目进展

今日项目完成了多项重要修复和功能合并，主要集中在提升安装可靠性、修复核心Bug以及推进配置管理优化。

- **关键修复合并：**
    - **[PR #7554] fix(install): install web dashboard on source installs:** 修复了源码安装后Web仪表盘静态文件未被复制到正确目录，导致服务无页面可用的严重Bug。该问题直接影响通过systemd启动的用户体验。 [查看PR](https://github.com/zeroclaw-labs/zeroclaw/pull/7554)
    - **[PR #7553] [Bug]: source install never copies web/dist to the data dir:** 对应的Bug报告已被关闭，标志着该问题已解决。 [查看Issue](https://github.com/zeroclaw-labs/zeroclaw/issues/7553)

- **功能与配置优化：**
    - **[PR #7555 & #7556] feat(config): declarative section grouping for the Config menu:** 为配置菜单引入了声明式的分区分组功能，将配置项逻辑归类（如“Foundation”、“Tools”等），提升了用户通过UI配置的便利性。 [查看PR #7555](https://github.com/zeroclaw-labs/zeroclaw/pull/7555) | [查看PR #7556](https://github.com/zeroclaw-labs/zeroclaw/pull/7556)

## 社区热点

今日社区讨论最集中的是围绕**运行时核心重构**和**技能系统**的大型PR，它们代表了项目的未来发展方向。

- **热点1：运行时引擎三合一重构**
    - **[PR #7540] refactor(runtime): consolidate the three agent turn engines onto run_tool_call_loop:** 这是一个XL规模、高风险的重构PR，旨在将三种不同的Agent交互循环引擎合并为一个统一的`run_tool_call_loop`。这是对核心架构的重大简化，影响深远，社区高度关注其稳定性和兼容性。 [查看PR](https://github.com/zeroclaw-labs/zeroclaw/pull/7540)

- **热点2：大型技能系统增强与集成**
    - **[PR #6667] feat(skills): background review fork + skill_manage tool:** 该项目历史上最大的PR之一，引入了后台技能审查分支和`skill_manage`工具，实现Agent自我改进和学习。社区对此功能的最终形态和应用场景有浓厚兴趣。 [查看PR](https://github.com/zeroclaw-labs/zeroclaw/pull/6667)
    - **[PR #6693] feat(memory): add dream mode for periodic memory consolidation:** 引入了“梦境模式”，用于定期的记忆整合。这是一个创新性的功能探索，社区正在热烈讨论其“本地优先”的设计哲学和对长期记忆能力的影响。 [查看PR](https://github.com/zeroclaw-labs/zeroclaw/pull/6693)

## Bug 与稳定性

今日报告的Bug数量较少，多为次要问题，且已有关联的修复PR。以下按严重程度排列：

- **S2 - 降级行为：**
    - **[Issue #5470] [Bug]: Multiple issues when running safely (已关闭):** 报告称运行`runtime/daemon`时，在Telegram通道上发现同一消息被重复保存的问题。该Issue已被关闭。 [查看Issue](https://github.com/zeroclaw-labs/zeroclaw/issues/5470)

- **S3 - 次要问题：**
    - **[Issue #7557] [Bug]: In zerocode Chat the Shift+Enter not works (新开):** 用户在`zerocode/tui`组件中发现无法使用Shift+Enter换行。这是一个影响聊天交互的可用性问题，尚未有修复PR。 [查看Issue](https://github.com/zeroclaw-labs/zeroclaw/issues/7557)
        - 目前已有关联的修复PR **(*暂无*)**。

- **其他已修复Bug：**
    - **[PR #7536] fix(channels/whatsapp-web): forward quoted media attachments:** 修复了WhatsApp频道中转发引用媒体附件（如图片、视频）时丢失的问题。 [查看PR](https://github.com/zeroclaw-labs/zeroclaw/pull/7536)
    - **[PR #7552] fix(skills): respect timeout_secs from SKILL.toml:** 修复了技能超时时间无法从配置文件读取，始终使用硬编码60秒限制的Bug。 [查看PR](https://github.com/zeroclaw-labs/zeroclaw/pull/7552)

## 功能请求与路线图信号

今日社区提交的新功能请求及开放中的大型增强PR，为未来版本绘制了清晰路线图。

- **v0.8.1 集成路线图：**
    - **[Issue #6970] [Tracker]: v0.8.1 integration/channel/provider/tool PR queue (开放中):** 这是一个追踪v0.8.1版本待合并PR队列的操作性问题，明确指示了项目下一阶段将重点推进集成、通道、提供商和工具的接入。 [查看Issue](https://github.com/zeroclaw-labs/zeroclaw/issues/6970)

- **核心功能增强候选：**
    - **[Issue #5570] [Enchancment]: Faster SQLite memory vector search works in O(n) (已关闭):** 用户请求将SQLite内存后端的向量搜索从暴力扫描(O(n))升级为近似最近邻索引(ANN)。该建议已被采纳并关闭，极有可能作为性能优化项进入后续开发计划。 [查看Issue](https://github.com/zeroclaw-labs/zeroclaw/issues/5570)
    - **（从PR看候选）**：如上文“社区热点”提到的 **[PR #7540]** 运行时引擎重构、**[PR #6667]** 技能系统大改、**[PR #6693]** 梦境记忆模式，这些都是明确的路线图信号，表明项目致力于深化Agent能力、提升架构健壮性和扩展长期记忆。

## 用户反馈摘要

从今日的Issues评论中，可以提炼出以下用户反馈：

- **安装体验痛点：** 用户`singlerider`报告了源码安装后Web仪表盘不工作的问题，这是一个影响开发者快速上手的实际障碍。该Bug已得到快速修复，体现了项目对开发者体验的重视。
- **聊天交互不便：** 用户`technologicalsingularity`反馈`zerocode Chat`中Shift+Enter无法换行，直接影响了其高效输入和使用习惯。这是一个典型的可用性反馈，表明即使在核心功能实现后，细节上的交互体验仍有优化空间。
- **功能性能期望：** 用户`ottogin`提出的SQLite向量搜索加速需求，反映了社区用户在处理大规模记忆时的性能痛点。用户期望以更低的成本（ANN）换取更快的搜索体验，而非满足于当前O(n)的线性搜索。

## 待处理积压

以下为长期未合并或未响应，但风险较高或重要性突出的Issue和PR，建议维护者重点关注：

- **高风险且长期未合并的PR：**
    - **[PR #6719] fix(runtime,channels): persist model_switch across turns in all paths:** 修复模型切换无法跨轮次保持的问题，风险高，已打开近一个月，当前标记为`needs-author-action`。 [查看PR](https://github.com/zeroclaw-labs/zeroclaw/pull/6719)
    - **[PR #6667] feat(skills): background review fork + skill_manage tool:** XL规模，高风险，同样标记为`needs-author-action`，距离提交已超过30天。 [查看PR](https://github.com/zeroclaw-labs/zeroclaw/pull/6667)
    - **[PR #6693] feat(memory): add dream mode for periodic memory consolidation:** 同样是XL规模，高风险，标记为`needs-author-action`。 [查看PR](https://github.com/zeroclaw-labs/zeroclaw/pull/6693)

- **需要关注的长期Issue：**
    - **[Issue #5470] [Bug]: Multiple issues when running safely (已关闭):** 尽管已关闭，但报告了**runtime/daemon**的S2降级行为，且涉及重复保存消息等不稳定状态，建议跟进确认复现并确保回归测试通过。 [查看Issue](https://github.com/zeroclaw-labs/zeroclaw/issues/5470)
    - **[Issue #5570] [Enchancment]: Faster SQLite memory vector search (已关闭):** 已关闭但功能实现尚未出现在任何PR中，需确认是否已排入内部开发计划。 [查看Issue](https://github.com/zeroclaw-labs/zeroclaw/issues/5570)

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*