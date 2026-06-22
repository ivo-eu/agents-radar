# AI CLI 工具社区动态日报 2026-06-23

> 生成时间: 2026-06-22 17:18 UTC | 覆盖工具: 9 个

- [Claude Code](https://github.com/anthropics/claude-code)
- [OpenAI Codex](https://github.com/openai/codex)
- [Gemini CLI](https://github.com/google-gemini/gemini-cli)
- [GitHub Copilot CLI](https://github.com/github/copilot-cli)
- [Kimi Code CLI](https://github.com/MoonshotAI/kimi-cli)
- [OpenCode](https://github.com/anomalyco/opencode)
- [Pi](https://github.com/badlogic/pi-mono)
- [Qwen Code](https://github.com/QwenLM/qwen-code)
- [DeepSeek TUI](https://github.com/Hmbown/DeepSeek-TUI)
- [Claude Code Skills](https://github.com/anthropics/skills)

---

## 横向对比

好的，作为专注于 AI 开发工具生态的资深技术分析师，我已仔细审阅了 2026-06-23 各主流 AI CLI 工具的社区动态。现基于上述数据，为您生成一份横向对比分析报告。

---

### **AI CLI 工具生态横向对比分析报告 (2026-06-23)**

#### 1. 生态全景

当前 AI CLI 工具生态正处于 **从“可用性”向“可靠性”与“安全性”爬坡的激烈竞争阶段**。社区的热情已从单纯的模型能力接入，转向对生产级特性的务实需求，包括：**更强的权限与安全沙箱**、**稳定的会话与状态管理**、**灵活的 MCP（模型上下文协议）生态集成**，以及**可预测的计费与预算控制**。各工具在“成为开发者不可或缺的自动化副驾驶”这一目标上趋同，但在实现路径上呈现出差异化特色。整体生态活跃度高，Bug 修复与功能迭代高频并行，但也暴露出部分工具在跨平台兼容性和核心稳定性上仍有优化空间。

#### 2. 各工具活跃度对比

| 工具名称 | 活跃 Issues (精选) | 重要 PR/修复 | 版本发布 | 社区讨论热度 (今日速览) |
| :--- | :--- | :--- | :--- | :--- |
| **Claude Code** | 10 (高热度，以权限、MCP、会话污染为主) | 3 (以文档修复为主) | 无 | 高 (围绕核心功能增强与Bug) |
| **OpenAI Codex** | 10 (极高热度，以成本飙升、日志写入为核心) | 10 (涵盖架构重构、性能优化、功能增强) | 2个Alpha版本 | 极高 (成本与性能问题是引爆点) |
| **Gemini CLI** | 10 (以Agent稳定性、子代理Bug为主) | 10 (修复SIGINT取消、配置合并、文件损坏等) | 无 | 高 (Agent行为可靠性是核心议题) |
| **GitHub Copilot CLI**| 10 (以认证恢复、平台崩溃、插件结构为主) | 1 (无效PR) | 无 | 中 (问题聚焦，但讨论深度高) |
| **Kimi Code CLI** | 6 (以MCP配置管理、KOSONG兼容性为主) | 2 (发布新版本，修复核心逻辑) | **v1.48.0** | 中 (问题具体，开发者反馈集中) |
| **OpenCode** | 10 (以MCP参数传递、模型兼容、性能为主) | 10 (覆盖修复、新功能、RFC文档) | 无 | 高 (MCP和模型问题是核心痛点) |
| **Pi** | 10 (以连接可靠性、本地模型支持、扩展API为主) | 10 (修复安全、成本显示、嵌套Git问题) | **v0.79.10** | 高 (连接稳定性和安全隐私是焦点) |
| **Qwen Code** | 10 (以安全标签伪造、工具无限循环、配置容错为主) | 10 (修复安全、工具调用、CI/CD失败) | Nightly版本 | 中高 (侧重安全与工程化严谨性) |
| **CodeWhale** | 10 (以安全加固、Agent过度自主、沙箱限制为主) | 10 (修复安全、沙箱、CI/CD、UI改进) | **v0.8.64** | 高 (安全与Agent行为可控性是核心) |

#### 3. 共同关注的功能方向

多个工具社区不约而同地将焦点集中在以下方向，表明这些是当前 AI CLI 工具面临的共性挑战：

- **MCP 协议集成与增强**：
    - **Claude Code**：请求 `envFile` 支持、OAuth 注册优化。
    - **Kimi Code**：多起 MCP 服务器配置管理 Bug（自动发现残留、路径错误、ACP模式失效）。
    - **OpenCode**：核心 `object` 类型参数序列化错误，导致工具调用失败。
    - **Pi**：强调 MCP 在扩展生态中的重要性。
    - **趋势**：社区已不满足于“能用”，而是要求 MCP 集成在安全性、配置灵活性、协议兼容性上达到生产级标准。

- **权限与安全**：
    - **Claude Code**：复合命令权限拆分（#16561），是需求最热的特性之一。
    - **Gemini CLI**：信任对话框显示错误的 Hook 信息。
    - **OpenCode**：Bash 工具无防护，可执行破坏性命令。
    - **Qwen Code**：`autofix` 流程的标签可被 LLM 伪造，存在安全漏洞。
    - **CodeWhale**：安全加固是整个版本发布的核心，沙箱与 Git Worktree 冲突也是热点。
    - **趋势**：从“允许或拒绝”的二元权限，正走向**精细化、可审计、可预测的权限模型**，以防止 Agent 滥用能力。

- **会话管理与数据完整性**：
    - **Claude Code**：会话恢复时数据被重复记录污染（#69013），导致历史“消失”。
    - **Copilot CLI**：恢复会话时认证失败，`/restart` 消耗大量 Credits。
    - **CodeWhale**：讨论默认启用无缝自动压缩（#3363），以解决长对话的核心痛点。
    - **Kimi Code**：社区再次呼吁跨会话的持久化内存系统（#1283）。
    - **趋势**：稳定的会话状态（可恢复、无污染、可管理）是构建用户信任的基础；无缝的上下文压缩和长期记忆系统是提升工作效率的关键。

- **成本与计费透明度**：
    - **OpenAI Codex**：**#28879**（速率限制成本飙升）是今日热度最高的话题，引发大量用户共鸣。
    - **Claude Code**：Max 计划下 1M 上下文仍显示额外收费（#45390）。
    - **Copilot CLI**：`/restart` 操作异常消耗 AI Credits。
    - **Pi**：改进了 OpenRouter 实际费用的显示（#5950）。
    - **趋势**：随着模型使用量增加，用户对**每 Token 成本、预算预警、计费逻辑透明度**的敏感度急剧上升。价格稳定性与可预测性是留住用户的关键。

- **平台兼容性**：
    - **Claude Code**：Windows 平台 401 认证失败（#69706）和大小写敏感问题（#62288）。
    - **OpenAI Codex**：Intel Mac SIGTRAP 崩溃（#29000）、Windows 11 频繁卡顿（#20214）。
    - **Copilot CLI**：Windows ARM64 高负载崩溃（#3687）。
    - **趋势**：**Windows（尤其是 ARM64 版本）** 和 **Intel Mac** 成为跨平台兼容性的重灾区。开发者期望工具在所有主流开发环境上提供一致稳定的体验。

#### 4. 差异化定位分析

尽管功能方向趋同，各工具的社区反馈仍揭示了其独特的定位与用户群：

| 工具 | 核心定位与社区特质 | 差异化侧重 |
| :--- | :--- | :--- |
| **Claude Code** | **全能型选手**，社区对最前沿功能（如复合命令、MCP OAuth、会话管理）细节要求极高。 | 强调**权限管理的颗粒度**和**会话的零污染稳定性**，用户对计费模型（Max计划）最为敏感。 |
| **OpenAI Codex** | **成本敏感型**和**IDE深度集成型**工具。社区对价格和资源消耗（日志、配额）极度敏感。 | **技术架构的前瞻性**（世界状态重构、OTEL扩展）很强，同时反映出IDE集成和本地资源优化是提升体验的关键。 |
| **Gemini CLI** | **Agent行为可靠性专家**。社区讨论集中在子代理、任务完成判定、SIGINT取消等行为逻辑上。 | 聚焦于**Agent内部决策逻辑的健壮性**，如子代理误报成功、工具调用死锁、配置合并Bug等，追求算法上的鲁棒性。 |
| **GitHub Copilot CLI**| **微软生态的粘合剂**。社区关注点始终围绕与GitHub生态（认证、MCP Registry、Sandbox文档）的集成。 | **深度的生态绑定**和**企业级特性**，如WSL凭证管理、企业HTTP代理、沙盒per-host配置等。 |
| **Kimi Code CLI** | **追求稳定性的新锐**。社区问题集中，但没有太多宏大叙事，更关注当前版本的具体Bug。 | 侧重于**MCP配置管理的稳定性**和**推理层（KOSONG）的API兼容性**，以及**智能体“灵魂”模块的防死循环**。 |
| **OpenCode** | **开源社区驱动的全能工具**。社区反馈最为多元，从MCP到模型兼容性，再到UI细节。 | 强调**AI与GitHub工作流的深度对齐**（模板发现、PR安全），并积极探索**计算机操作（Computer Use）** 等前沿能力。 |
| **Pi** | **高度可扩展的平台**。社区围绕其扩展机制和本地模型支持展开热烈讨论。 | **以扩展为核心的差异化**，提供强大的API来自定义行为。同时非常注重**安全（机密泄露）** 和 **成本优化（自动路由）**。 |
| **Qwen Code** | **工程化严谨的挑战者**。社区反馈显示出对**安全规范、配置容错、CI/CD流程**的高度重视。 | **工程治理出色**，对环境变量类型验证、CI/CD流水线失败、安全标签伪造等细节问题的发现和修复速度很快。 |
| **CodeWhale** | **安全与可控的先行者**。几乎每个热点Issue都与安全、沙箱、Agent行为控制相关。 | **将安全与可控性作为核心卖点**，通过沙箱、护栏、审计门禁等机制确保Agent行为可预测、可干预。 |

#### 5. 社区热度与成熟度

- **最具热度与争议性**：**OpenAI Codex** 因“成本飙升”事件成为今日社区最沸沸扬扬的中心，讨论度高且情绪强烈。**Claude Code** 和 **OpenCode** 社区也非常活跃，但更多是针对具体功能缺陷的长期打磨。
- **快速迭代，积极修复**：**Gemini CLI**、**Pi**、**Qwen Code** 和 **CodeWhale** 表现出极强的版本迭代和 Bug 修复能力，PR 进展迅速。尤其是 **Qwen Code** 和 **CodeWhale**，其社区反馈与 PR 修复高度耦合，显示出高效的工程执行力。
- **沉稳但新锐**：**Kimi Code CLI** 和 **GitHub Copilot CLI** 社区规模相对较小，问题反馈更为集中。Kimi 发布了新版本，但问题依然突出；Copilot 虽然问题不多，但每个问题都指向其核心生态的病点（认证、平台、插件）。
- **成熟度判断**：
    - **成熟期**：**Claude Code** 和 **OpenAI Codex** 社区对功能的讨论已深入到非常细节的角落，但同时也暴露出其在规模化使用后的“老化”问题（如会话污染、成本失控）。
    - **成长期**：**Gemini CLI**、**Pi**、**OpenCode**、**Qwen Code** 和 **CodeWhale** 正处于功能膨胀和快速完善的阶段，社区问题多样，但解决效率高。
    - **早期追赶**：**Kimi Code CLI** 正在努力解决基础设施问题（如 MCP 管理），以确保其在功能上能跟上第一梯队。

#### 6. 值得关注的趋势信号

1.  **“护栏”与“信任”是核心矛盾**：社区对 Agent **“过度自主”**（CodeWhale #3275）和 **“破坏性行为”**（Gemini #22672）的担忧，已催生出一系列技术方案，如预审策略、编辑护栏、精细权限等。**未来的竞争点在于：如何在不牺牲效率的前提下，构建让开发者放心的“AI Copilot”。**

2.  **MCP 从“噱头”走向“基建”**：MCP 不再是简单的“可接入”，而是暴露出大量底层问题（参数序列化、路径解析、认证复用、配置文件管理）。这预示着 MCP 生态正在从“浅层集成”向“深度嵌入”转型，**谁能率先解决 MCP 的稳定性、安全性和易用性，谁就能在第三方工具生态中占据先机。**

3.  **“成本危机”是悬在所有工具头上的达摩克利斯之剑**：OpenAI Codex 的 #28879 是一个强烈信号。当模型使用成本变得不可预测时，用户会迅速失去信任。**未来，提供透明、可控的计费策略和预算工具，将成为 AI CLI 标配功能。** 这也催生了自动路由（如 Pi #5970）等成本优化方案的需求。

4.  **“环境变量类型检查”等微小工程细节，正成为质量的试金石**：Qwen Code 一口气修复了多个因错误地接受小数作为整数配置而引发的 Bug。这表明，在 AI 工具快速迭代中，**对“防御性编程”和“配置容错性”的忽视，会累积成大量的“隐形 Bug”，最终消磨开发者的信任。**

5.  **从“代码助手”到“开发环境管家”**：多个社区（CodeWhale、OpenCode）正在探索 **Web 自动化测试**、**本地开发服务器管理**、**计算机操作**等超出传统“写代码”范畴的能力。这预示着 AI CLI 的目标正从“辅助生成代码”，向**“全流程代理开发环境”** 演变。

**对开发者的建议**：在选择 AI CLI 工具时，除了模型能力，应优先关注其**安全模型、会话可靠性、成本透明度**和**跨平台兼容性**。目前来看，没有完美无缺的工具，但可以根据团队在“效率”、“安全”、“成本”和“生态”上的不同优先级，选择最匹配的选项。例如，追求极致安全可选 **CodeWhale**；追求成本控制需密切关注 **OpenAI Codex** 的动态；希望在微软生态内无缝衔接则 **GitHub Copilot CLI** 是不错选择。

---

## 各工具详细报告

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills 社区热点

> 数据来源: [anthropics/skills](https://github.com/anthropics/skills)

好的，作为一名专注于 Claude Code 生态的技术分析师，以下是基于您提供的数据生成的社区热点报告。

---

### Claude Code Skills 社区热点报告 (数据截止: 2026-06-23)

#### 1. 热门 Skills 排行

以下为社区关注度最高、讨论最激烈的 5 个 Skills（PR）：

1.  **skill-creator 修复：run_eval.py 触发率零问题**
    - **功能**: 修复 `run_eval.py` 在评估技能描述时总是报告 0% 召回率的根本性 Bug。此脚本是技能优化循环的核心，该 Bug 导致优化过程失效。
    - **社区讨论热点**: 社区围绕该 PR 进行了密集的协作，贡献了 10 余次独立复现及多个 PR (#1298, #1099, #1050)，共同排查 Windows 兼容性问题（`claude.cmd` 执行、管道读取崩溃等）和 YAML 解析问题。
    - **当前状态**: **OPEN**
    - **链接**: [PR #1298](https://github.com/anthropics/skills/pull/1298)

2.  **新增文档排版技能 (document-typography)**
    - **功能**: 解决 AI 生成文档中的常见排版问题，如孤行、寡段和编号错位，提升文档的专业性与可读性。
    - **社区讨论热点**: 社区普遍认为这是一个“价值洼地”技能，解决了用户“感觉不对但说不出来”的痛点。讨论集中在其极高的通用性和实用性上。
    - **当前状态**: **OPEN**
    - **链接**: [PR #514](https://github.com/anthropics/skills/pull/514)

3.  **新增 ODT 格式支持技能 (odt)**
    - **功能**: 让 Claude 能够创建、编辑、填充和转换 OpenDocument 格式文件（.odt, .ods），直接对标 LibreOffice。
    - **社区讨论热点**: 反映了社区对开源/标准办公格式的强烈需求，以及对办公软件生态从 Microsoft Office 向更开放生态拓展的期待。
    - **当前状态**: **OPEN**
    - **链接**: [PR #486](https://github.com/anthropics/skills/pull/486)

4.  **新增 ServiceNow 平台技能 (servicenow)**
    - **功能**: 一个全面的 ServiceNow 平台助手，覆盖 ITSM、ITOM、SecOps、HR 等多个模块，旨在成为企业级平台的通用接口。
    - **社区讨论热点**: 体现了企业用户对将 Claude 集成到自身 IT 和业务流程中的强烈需求，技能覆盖度广，但讨论也隐含了对技能复杂度与可维护性的关注。
    - **当前状态**: **OPEN**
    - **链接**: [PR #568](https://github.com/anthropics/skills/pull/568)

5.  **新增 Masonry AI 图像与视频生成技能**
    - **功能**: 集成 Masonry CLI，使 Claude 能够调用 Imagen 3.0、Veo 3.1 等模型进行文生图、文生视频及任务管理。
    - **社区讨论热点**: 社区对多模态能力非常兴奋，视其为 Claude 能力从文本扩展到视觉内容创作的关键一步。讨论集中在与现有工具的集成方式上。
    - **当前状态**: **OPEN**
    - **链接**: [PR #335](https://github.com/anthropics/skills/pull/335)

6.  **新增 AURELION 认知与记忆框架技能套件**
    - **功能**: 引入一套复杂的认知和记忆框架（内核、顾问、代理、记忆），旨在为 AI Agent 提供结构化思考和长期记忆能力。
    - **社区讨论热点**: 这是社区对“增强 Claude 原生能力”的深层探索，特别是关于 Agent 的持久上下文和思考深度。讨论焦点在于认知架构的设计及其对 Agent 行为的影响。
    - **当前状态**: **OPEN**
    - **链接**: [PR #444](https://github.com/anthropics/skills/pull/444)

#### 2. 社区需求趋势

从 Issues 中可提炼出三个最核心的社区需求方向：

1.  **工具链的稳定性与可靠性（压倒性需求）**:
    - **核心痛点**: 技能创建工具 `run_eval.py` 存在严重缺陷，导致整个技能优化流程失效（#556）。该问题引发了 10+ 条评论和多个修复 PR，是当前社区最关注的焦点。
    - **平台兼容性**: 技能创建工具对 Windows 用户不友好，存在子进程执行、编码处理等多处兼容性问题（#1061, #1169）。
    - **数据安全**: 用户对在技能中处理敏感数据（如 SharePoint 权限逻辑）的安全性和上下文窗口占用感到担忧（#1175）。

2.  **规模化部署与企业级功能**:
    - **组织级共享**: 社区强烈要求官方提供 Skills 的组织内共享机制（#228），而非目前的手动文件传输方式。这是社区呼声最高的特性请求之一。
    - **认证与安全**: 对社区技能可能冒充官方技能、突破信任边界的安全问题表达了担忧（#492），呼吁官方提供认证或沙箱机制。

3.  **技能生态的标准化与互操作性**:
    - **标准化描述**: 用户反馈 `skill-creator` 技能本身像一个开发文档，而非一个可执行的技能，要求其自身遵循最佳实践（#202）。
    - **MCP 集成**: 社区持续探索将 Skills 能力暴露为标准 MCP 服务的可能性（#16），以实现更广泛的互操作。

#### 3. 高潜力待合并 Skills

以下 PR 评论活跃、功能需求明确，具有较高落地潜力：

1.  **文档排版技能** (`document-typography`, PR #514): **极高适用性**。几乎所有生成文档的场景都需要，且功能点明确（孤行、寡段等），技术实现清晰。合并优先级极高。
2.  **ODT 格式支持** (`odt`, PR #486): **生态互补性强**。填补了 Skills 在开源办公软件生态中的空白，对 LibreOffice 用户是刚需。
3.  **ServiceNow 平台技能** (`servicenow`, PR #568): **企业市场潜力巨大**。直接切入企业级 ITSM 市场，有助于推动 Claude 在企业内部落地，具有战略价值。
4.  **测试模式技能** (`testing-patterns`, PR #723): **开发刚需**。测试是软件开发的核心环节，一个系统化的测试技能可以显著提升开发效率和质量，社区需求明确。
5.  **YAML 特殊字符检测修复** (PR #361, #539): **Bug 修复**。虽然是一个小功能，但它直接影向技能描述的正确解析，属于维护项目健康度的关键修复，应尽快合并。

#### 4. Skills 生态洞察

**一句话总结**: 当前社区最集中的诉求在于 **“夯实基础”** ——即修复技能创建工具链（`skill-creator`）的关键性能和兼容性缺陷，并建立机构级的共享与安全机制，以确保 Skills 生态从“可玩”迈向“可靠可用”。

---

# Claude Code 社区动态日报（2026-06-23）

**数据来源：** github.com/anthropics/claude-code

---

## 今日速览

过去24小时内，社区围绕**权限系统增强**与**MCP 认证流程优化**讨论最为活跃；同时多个 **Windows 平台 Bug** 及 **会话数据污染** 问题引发关注，开发者对计费异常和子代理模型继承的反馈较为集中。另外，文档修复类 PR 获少量合并。

---

## 版本发布

过去24小时内无新版本发布。

---

## 社区热点 Issues（10 条）

### 1. 复合 Bash 命令权限解析
- **#16561**：请求将 `&&`、`|`、`||` 等复合命令中的每个子命令单独与权限规则匹配，而非当作整体字符串判断。当前行为导致本应放行的安全命令被要求额外审批。
- 社区反响热烈：44 条评论，168 个 👍，仍在开放中。这是权限系统最受期待的功能改进之一。
- 🔗 https://github.com/anthropics/claude-code/issues/16561

### 2. Max 计划下 1M 上下文错误提示
- **#45390**：用户反馈，在 Max 计划中选用 Opus 4.6（1M 上下文）时仍显示“需额外使用量”，明显违反套餐权益。35 条评论，开放中。
- 影响面广，涉及计费与模型选择逻辑。
- 🔗 https://github.com/anthropics/claude-code/issues/45390

### 3. 源码仓库作为组织技能来源
- **#28729**：建议允许将 Git 仓库（如 GitHub 组织仓库）直接链接为 Skills 来源，而非仅支持本地文件。28 条评论，54 个 👍。
- 团队协作场景下的高频需求。
- 🔗 https://github.com/anthropics/claude-code/issues/28729

### 4. 自动模式分类器持久 HTTP 429
- **#60438**：用户描述自动模式分类器（xml_s1）反复返回 429 限流错误，即使重装也无法解决，怀疑账户端配置问题。15 条评论，已关闭（可能已修复或标记为已知问题）。
- 🔗 https://github.com/anthropics/claude-code/issues/60438

### 5. Windows 平台 401 认证失败
- **#69706**：Windows 用户频繁遇到 `API Error: 401 Invalid authentication credentials`，即使凭证正确。15 条评论，9 个 👍，开放中。
- 平台特定认证 Bug，影响较大。
- 🔗 https://github.com/anthropics/claude-code/issues/69706

### 6. Windows VS Code 扩展会话列表消失
- **#62288**：当项目工作目录大小写与记录的 cwd 不一致时（如 `D:\Project` vs `d:\project`），VS Code 扩展中历史会话不显示。7 条评论，开放中。
- 典型的大小写敏感问题，Windows 用户群广泛。
- 🔗 https://github.com/anthropics/claude-code/issues/62288

### 7. MCP 支持 `envFile` 加载环境变量
- **#28942**：请求在 `.mcp.json` 配置中增加 `envFile` 字段，以便在环境变量展开前从文件加载密钥。6 条评论，16 个 👍。
- 提升 MCP 服务器安全性，社区反响积极。
- 🔗 https://github.com/anthropics/claude-code/issues/28942

### 8. MCP OAuth 每次认证都重新注册客户端
- **#59460**：Claude Code 在 MCP 服务器认证时总是执行全新的动态客户端注册（DCR），导致旧 client_id 及其 refresh_token 被废弃。4 条评论，开放中。
- 影响自托管 MCP 服务器的持续使用。
- 🔗 https://github.com/anthropics/claude-code/issues/59460

### 9. 聊天记录重复附加致会话顺序错乱
- **#69939**：打开一个旧聊天会无条件向 JSONL 文件追加 `mode`、`custom-title` 等记录（即使值与已有记录相同），导致文件 mtime 更新、会话在“最近”列表中提前。3 条评论，刚刚提出。
- 新 Bug，直接影响会话管理体验。
- 🔗 https://github.com/anthropics/claude-code/issues/69939

### 10. 会话恢复污染文件致历史记录“消失”
- **#69013**：使用 `--continue` 或 `--resume` 恢复对话时，JSONL 文件被大量重复的 `mode` 和 `permission-mode` 记录污染，导致实际对话内容不可见。3 条评论，开放中。
- 严重影响工作连续性，需优先修复。
- 🔗 https://github.com/anthropics/claude-code/issues/69013

---

## 重要 PR 进展

### #70074 – 修复插件开发文档中过时的市场名称
- 作者：itxaiohanglover，于 2026-06-22 提出，无评论，0 👍。
- 将 `plugin-dev/README.md` 中三处 `claude-code-marketplace` 更新为 `claude-code-plugins`，修正仓库重命名后的遗留文本。
- 🔗 https://github.com/anthropics/claude-code/pull/70074

### #70066 – 更新插件市场安装文档
- 作者：abhi-zit77，于 2026-06-22 提出，无评论，0 👍。
- 更新安装说明，使用官方插件市场名称；将本地开发示例从 `cc --plugin-dir` 改为 `claude --plugin-dir`；澄清贡献指向。
- 🔗 https://github.com/anthropics/claude-code/pull/70066

### #69916 – 修复 `edit-issue-labels.sh` 静默退出问题
- 作者：Dreamstick9，于 2026-06-21 提出，无评论，0 👍。
- 当脚本未传入 `--add-label` 或 `--remove-label` 参数时，原逻辑 `exit 1` 未打印错误信息。现添加错误提示后再退出。
- 🔗 https://github.com/anthropics/claude-code/pull/69916

---

## 功能需求趋势

1. **权限系统精细化**  
   - 复合命令权限拆分（#16561）呼声最高，用户期望工具对每一个子命令进行独立匹配，减少不必要的审批弹窗。

2. **MCP 能力增强**  
   - `envFile` 支持（#28942）、OAuth 重复注册优化（#59460）是 MCP 生态成熟的关键障碍；开发者期望更安全、更高效的认证与配置方式。

3. **团队协作与技能管理**  
   - 支持从源码仓库链接组织技能（#28729），便于团队共享和版本控制既有提示库。

4. **会话与聊天体验改进**  
   - 会话恢复避免数据污染（#69013）、`/compact` 预配置提示（#70083）、多回复批量提交（#55152）等，反映用户对稳定、高效会话管理的强烈需求。

5. **跨平台兼容性**  
   - Windows 平台认证（#69706）、大小写敏感问题（#62288）持续被反馈，表明平台适配仍有较大优化空间。

6. **计费与用量透明度**  
   - 1M 上下文额外收费问题（#45390）、自动充电循环计费错误（#68773）凸显用户对价格策略透明度和账户控制机制的关注。

---

## 开发者关注点

- **认证与令牌管理**：401 错误、MCP OAuth 重复注册、Windows 平台认证失败，开发者普遍对认证流程的稳定性和可调试性感到困扰。
- **会话数据安全与完整性**：JSONL 文件被重复记录污染、聊天历史“消失”等问题被多次报告，直接动摇用户对工具的信任。
- **平台特定 Bug**：Windows 用户遭遇的问题（大小写、RPC 消息过大等）在社区中占比上升，但修复进度较为缓慢。
- **子代理行为异常**：在 Plan 模式中子代理继承模型错误（#67942）、读入 `node_modules` 而非源文件（#55146），影响复杂任务执行的可靠性。
- **计费争议**：自动充电循环产生高额费用且难以联系人工客服（#68773），用户对自助服务与退款机制提出质疑。

---

*日报整理于 2026-06-23，覆盖 GitHub 仓库过去 24 小时（截至 2026-06-22）的更新数据。*

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex 社区动态日报 — 2026-06-23

## 今日速览
- **速率限制成本激增**：Issue #28879 报告自 6 月 16 日起 `gpt-5.5` 每 token 消耗的配额暴涨 10–20 倍，Plus 用户 2–3 次 prompt 即可耗尽 5 小时预算，获得 231 👍 和 116 条讨论。
- **本地日志写入量惊人**：Issue #28224 指出 SQLite 反馈日志年写入量可达 ~640 TB，严重消耗 SSD 寿命，引发社区对日志机制的热议（214 👍）。
- **两个 Alpha 版本发布**：`rust-v0.142.0-alpha.10` 与 `rust-v0.142.0-alpha.11`，暂无详细变更说明。

---

## 版本发布
### `rust-v0.142.0-alpha.10` & `rust-v0.142.0-alpha.11`
- **发布时间**：2026-06-22（过去 24 小时内）
- **内容**：仅标注为 Alpha 版本，无配套 Release Notes。建议关注后续 commits 或相关 PR 了解具体变更。
- [查看 Release](https://github.com/openai/codex/releases)

---

## 社区热点 Issues（Top 10）

### 1️⃣ 速率限制成本飙升（#28879）
- **摘要**：`gpt-5.5` 在 Plus 套餐下，每 token 消耗的 rate-limit 配额上涨 10–20 倍，正常使用 2–3 次 prompt 即耗尽 5 小时预算。用户怀疑是 6 月 16 日后 API 计算方式变更或 bug。
- **社区反应**：231 👍，116 条评论，大量 Plus/Pro 用户复现并附日志截图，要求紧急修复或回调。
- [GitHub 链接](https://github.com/openai/codex/issues/28879)

### 2️⃣ SQLite 日志写入量过大（#28224）
- **摘要**：本地 `~/.codex/logs_2.sqlite` 等文件持续高强度写入，单机年写入量估算达 ~640 TB，严重影响 SSD 寿命。反馈日志包含大量 TRACE 级别数据。
- **社区反应**：214 👍，25 条评论，开发者提出限制日志级别、轮转策略等建议。
- [GitHub 链接](https://github.com/openai/codex/issues/28224)

### 3️⃣ IDE 集成 diff/审批（#2998）
- **摘要**：期望 Codex CLI 的终端 diff 审批流程能直接嵌入 IDE（如 VS Code）侧边栏，提供可视化红绿对比和一键接受/拒绝。
- **社区反应**：198 👍，62 条评论，被称为“替代 Claude Code 的关键缺失功能”。
- [GitHub 链接](https://github.com/openai/codex/issues/2998)

### 4️⃣ 企业 HTTP 代理支持（#4242）
- **摘要**：要求 Codex CLI 自动遵从 `HTTPS_PROXY`、`HTTP_PROXY` 等环境变量，避免逐个请求手动配置。当前企业用户需 patch 多个调用点。
- **社区反应**：80 👍，5 条评论，企业用户迫切需求。
- [GitHub 链接](https://github.com/openai/codex/issues/4242)

### 5️⃣ Windows 11 频繁卡顿（#20214）
- **摘要**：Codex App 在 Windows 11 Pro（Ryzen 5 5600 / 32 GB RAM）上频繁无响应/卡顿，系统资源充足但界面冻结。
- **社区反应**：31 👍，18 条评论，多人复现，怀疑与渲染线程或 GPU 加速有关。
- [GitHub 链接](https://github.com/openai/codex/issues/20214)

### 6️⃣ TUI 内联图片显示（#29451）
- **摘要**：希望在 Codex TUI 中能直接预览图片（如 `view_image` 输出的图片），使用终端原生渲染或 ASCII 降级。
- **社区反应**：29 👍，0 条评论但为今日新开，获大量 👍 表明社区呼声高。
- [GitHub 链接](https://github.com/openai/codex/issues/29451)

### 7️⃣ Intel macOS SIGTRAP 崩溃（#29000）
- **摘要**：Codex CLI 0.141.0 在 Intel Mac 上使用 `gpt-5.5` 时触发 SIGTRAP 崩溃，回溯指向 native 代码错误。
- **社区反应**：11 👍，9 条评论，Intel Mac 用户受影响。
- [GitHub 链接](https://github.com/openai/codex/issues/29000)

### 8️⃣ 远程 compact 任务流断开（#14559）
- **摘要**：执行远程 compact 任务时出现“stream disconnected before completion”，任务未完成即中断。
- **社区反应**：14 👍，23 条评论，涉及 Windows 平台 + PowerShell 环境。
- [GitHub 链接](https://github.com/openai/codex/issues/14559)

### 9️⃣ GPT 5.5 被错误报告为 5.4（#29362）
- **摘要**：用户选择 `gpt-5.5` 但 usage 统计中大量显示 `gpt-5.4` 消耗，疑似模型选择或计费映射错误，可能导致预算计算混乱。
- **社区反应**：3 条评论，但涉及模型识别准确性，需关注。
- [GitHub 链接](https://github.com/openai/codex/issues/29362)

### 🔟 SIGINT 取消后工具继续执行（#29439）
- **摘要**：按下 Ctrl+C 取消后，Codex CLI 仍继续执行已发起的工具调用，导致不希望的操作（如文件修改）在取消后仍然完成。
- **社区反应**：1 条评论，但属于安全/行为问题，开发者需修复取消语义。
- [GitHub 链接](https://github.com/openai/codex/issues/29439)

---

## 重要 PR 进展（Top 10）

### 1️⃣ 环境上下文迁移至模型世界状态（#29249）
- **摘要**：将模型可见的环境上下文从临时 turn 值迁移为类型化、可重放的“世界状态”表示，为后续可重复、可调试的上下文管理打基础。
- **状态**：已 code-review，接近合并。
- [GitHub 链接](https://github.com/openai/codex/pull/29249)

### 2️⃣ 实时上下文 diff 并入世界状态（#29282）
- **摘要**：将 settings diff 逻辑从 `context_manager/updates.rs` 移入世界状态模块，保证多迭代 turn 中环境变化的一致性渲染。
- **状态**：OPEN，依赖 #29249。
- [GitHub 链接](https://github.com/openai/codex/pull/29282)

### 3️⃣ 插件根路径 URI 本地化（#28918）
- **摘要**：要求执行器插件根路径使用 `file://` URI 格式（如 `file:///opt/plugins/foo`），以统一跨平台路径处理，消除 Windows 反斜杠问题。
- **状态**：OPEN，核心基础架构变更。
- [GitHub 链接](https://github.com/openai/codex/pull/28918)

### 4️⃣ 支持无邮箱 ChatGPT 账户（#28991）
- **摘要**：允许使用 Personal Access Token 登录且无邮箱的账户，修复 PAT 登录因缺少 email 字段而失败的 bug。
- **状态**：OPEN，直接影响使用 PAT 的企业用户。
- [GitHub 链接](https://github.com/openai/codex/pull/28991)

### 5️⃣ OTEL 暴露服务层与推理力度（#29155）
- **摘要**：在 Codex CLI 的 OpenTelemetry `response.completed` 事件中添加 `service_tier` 和 `model_reasoning_effort` 字段，满足 NVIDIA 等企业测量 Fast 模式使用情况的需求。
- **状态**：已 code-review。
- [GitHub 链接](https://github.com/openai/codex/pull/29155)

### 6️⃣ 使用 input items 传递 Responses Lite 工具（#27946）
- **摘要**：将 Responses Lite 模式下的工具定义从顶层 `tools` 数组改为通过 `additional_tools` + developer item 方式注入，保持与 API 1-to-1 映射，避免命名空间冲突。
- **状态**：OPEN，长期建设。
- [GitHub 链接](https://github.com/openai/codex/pull/27946)

### 7️⃣ 可配置预算提醒阈值（#29423）
- **摘要**：将固定的 `reminder_interval_tokens` 替换为用户可配置的剩余 token 提醒列表（如 `[65536, 32768, ...]`），提供更精细的预算预警。
- **状态**：OPEN，用户体验改进。
- [GitHub 链接](https://github.com/openai/codex/pull/29423)

### 8️⃣ 远程 sandbox 拒绝语义报告（#29424）
- **摘要**：使远程执行服务器能返回结构化的 sandbox 拒绝原因（如权限、网络限制），而非仅显示“Invalid request”，帮助用户理解并调整配置。
- **状态**：OPEN，安全与调试相关。
- [GitHub 链接](https://github.com/openai/codex/pull/29424)

### 9️⃣ 过滤持久日志中的噪声目标（#29457）
- **摘要**：限制本地 SQLite 日志只记录关键目标，避免 TRACE 级别的高频日志（如 WebSocket 事件、OTEL 镜像）填满分区，缓解 #28224 问题。
- **状态**：已合并（CLOSED），快速修复。
- [GitHub 链接](https://github.com/openai/codex/pull/29457)

### 🔟 缓存 Codex Apps 工具（#29003）
- **摘要**：将 `list_all_tools()` 从重复的磁盘读取改为内存快照，启动时从磁盘加载，之后使用缓存，显著减少工具枚举开销。
- **状态**：已 code-review。
- [GitHub 链接](https://github.com/openai/codex/pull/29003)

---

## 功能需求趋势
- **IDE 深度集成**：将终端 diff 审批移植到 IDE（VS Code、code-server）是最高频需求（#2998、#2932），并期望改进侧边栏渲染性能。
- **性能与资源优化**：日志写入量过大（#28224）、Windows 卡顿（#20214）、TUI 内联图片（#29451）等反映社区对轻量、流畅体验的追求。
- **企业级特性**：HTTP 代理支持（#4242）、远程 sandbox 精细控制（#29424）、无邮箱账户登录（#28991）持续成为企业用户痛点。
- **透明度与可控性**：速率成本飙升后，社区要求更好的预算提醒（#29423）、模型选择与用量映射正确性（#29362）。
- **跨平台稳定性**：Intel macOS 崩溃（#29000）、Windows SSH 超时（#29253）表明平台兼容性仍需加固。

---

## 开发者关注点
| 痛点 / 高频需求 | 相关 Issue |
|----------------|------------|
| **速率限制成本异常** - 每次 prompt 消耗 10-20x 配额，严重依赖每日预算 | #28879 |
| **本地日志写入灾难** - SQLite 年写入 640 TB，SSD 寿命告急 | #28224 |
| **IDE 体验缺失** - 无法在 VS Code 中可视化 diff 和审批 | #2998, #2932 |
| **Windows 卡顿 / 崩溃** - 系统资源充足但 UI 冻结 | #20214, #29253 |
| **Intel Mac 兼容性** - SIGTRAP 崩溃导致 CLI 不可用 | #29000 |
| **远程执行稳定性** - stream 断开、sandbox 错误信息不明确 | #14559, #29424 |
| **取消语义** - Ctrl+C 后工具仍继续执行，存在安全隐患 | #29439 |
| **模型选择与用量统计不一致** - 选 5.5 却被计为 5.4 | #29362 |
| **预算提醒不灵活** - 固定间隔望改为用户自定义剩余 token 阈值 | #29423 |
| **TUI 功能完善** - 希望内联显示图片、改进 @ 菜单导航 | #29451, #28500 |

> *日报数据采集时间：2026-06-23 00:00 ~ 2026-06-23 23:59 (UTC+8)*  
> *数据源：GitHub openai/codex 仓库*

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

好的，以下是为您生成的 2026-06-23 Gemini CLI 社区动态日报。

---

# Gemini CLI 社区动态日报 | 2026-06-23

## 今日速览

社区活跃度持续高涨，尽管无新版本发布，但修复与功能开发势头强劲。**Agent 稳定性**仍是核心议题，涉及通用代理挂起、子代理恢复逻辑错误等关键 Bug 修复。此外，**安全与隐私**（如 Auto Memory 日志处理）和 **MCP 协议**的深入集成成为新的焦点，显示出社区正从基础功能向生产环境的高可用性和安全性演进。

## 社区热点 Issues

1. **[#21409] Generalist agent hangs (通用代理挂起)**
   - **热度:** 👍 8 | 评论 7 | P1/Bug
   - **重要性:** 严重阻碍用户使用的问题。当 Gemini CLI 委托给通用代理时，会无限期挂起，用户只能手动取消。这直接影响了 CLI 的核心自动化体验。
   - **链接:** https://github.com/google-gemini/gemini-cli/issues/21409

2. **[#22323] Subagent recovery after MAX_TURNS is reported as GOAL success (子代理达到最大轮次后被错误报告为成功)**
   - **热度:** 👍 2 | 评论 7 | P1/Bug
   - **重要性:** 一个危险的逻辑漏洞。子代理在未完成任何分析、仅因超出执行轮次上限而中断时，主系统却错误地将其标记为“成功”，掩盖了实际的中断和失败，导致用户误以为任务完成。
   - **链接:** https://github.com/google-gemini/gemini-cli/issues/22323

3. **[#24353] Robust component level evaluations (稳健的组件级评估)**
   - **热度:** 评论 7 | P1/Feature
   - **重要性:** 这是提升 CLI 内部测试质量的顶层 EPIC。它为引入“行为评估”测试而设，目前已生成 76 个测试用例，并覆盖 6 个 Gemini 版本。这是保证 Agent 行为可靠性的基石。
   - **链接:** https://github.com/google-gemini/gemini-cli/issues/24353

4. **[#22745] Assess the impact of AST-aware file reads, search, and mapping (评估AST感知的文件读取、搜索和映射的影响)**
   - **热度:** 👍 1 | 评论 7 | P2/Feature
   - **重要性:** 探索前沿技术。社区在探索利用**抽象语法树 (AST)** 来使 Agent 更精确地理解和操作代码库，目标是减少不必要的 Token 消耗和错误读取，是提升 Agent 代码理解能力的长期研究方向。
   - **链接:** https://github.com/google-gemini/gemini-cli/issues/22745

5. **[#25166] Shell command execution gets stuck with "Waiting input" (Shell 命令执行卡在“等待输入”)**
   - **热度:** 👍 3 | 评论 4 | P1/Bug
   - **重要性:** 常见且烦人的 Bug。执行完一个简单的 Shell 命令后，CLI 会错误地认为需要等待用户输入而卡死。这破坏了命令执行流程的连续性，严重降低效率。
   - **链接:** https://github.com/google-gemini/gemini-cli/issues/25166

6. **[#26525] Add deterministic redaction and reduce Auto Memory logging (增加确定的编辑功能并减少自动记忆日志)**
   - **热度:** 评论 5 | P2/Bug/Security
   - **重要性:** 安全问题。Auto Memory 功能在读取本地对话记录时，会将内容发送给模型进行“编辑”后再存储。但现有流程存在隐患：机密信息在发送给模型时可能已经暴露。该 Issue 旨在从源头解决问题。
   - **链接:** https://github.com/google-gemini/gemini-cli/issues/26525

7. **[#26522] Stop Auto Memory from retrying low-signal sessions indefinitely (停止自动记忆无限重试低信号会话)**
   - **热度:** 评论 5 | P2/Bug
   - **重要性:** 性能与资源浪费。Auto Memory 后台提取器会不断尝试处理它认为“低信号”的会话，导致 CPU 和磁盘 I/O 的无谓消耗。该 Issue 要求修正重试逻辑，避免死循环。
   - **链接:** https://github.com/google-gemini/gemini-cli/issues/26522

8. **[#21968] Gemini does not use skills and sub-agents enough (Gemini 不充分利用技能和子代理)**
   - **热度:** 评论 6 | P2/Bug
   - **重要性:** 核心体验问题。用户反馈，即使在明确配置了自定义“技能”（如 Git、Gradle）的情况下，模型仍然不主动调用它们，而是采用更笨拙的方式完成任务。这表明 Agent 的工具选择策略仍需优化。
   - **链接:** https://github.com/google-gemini/gemini-cli/issues/21968

9. **[#24246] Gemini CLI encounters 400 error with > 128 tools (工具超过128个时触发400错误)**
   - **热度:** 评论 3 | P2/Bug
   - **重要性:** 限制了CLI的可扩展性。随着用户安装更多 MCP 工具或自定义技能，工具总数很容易超过 API 限制。社区期望 Agent 能智能选择核心工具，而不是全量发送。
   - **链接:** https://github.com/google-gemini/gemini-cli/issues/24246

10. **[#22672] Agent should stop/discourage destructive behavior (代理应停止/阻止破坏性行为)**
    - **热度:** 👍 1 | 评论 3 | P2/Customer Issue
    - **重要性:** 安全护栏需求。用户担忧 Agent 会在执行 Git 操作、数据库管理等任务时使用 `git reset --force` 等危险命令。该 Issue 寻求更智能的防护机制，比如在可能造成破坏前进行二次确认或寻找更安全的替代方案。
    - **链接:** https://github.com/google-gemini/gemini-cli/issues/22672

## 重要 PR 进展

1. **[#28096] fix(core): drop late tool calls after SIGINT cancellation (修复SIGINT取消后，丢弃滞后的工具调用)**
   - **功能:** 修复了一个关键的中断处理问题：用户按下 Ctrl+C 取消后，模型可能仍然执行之前发起的工具调用，导致意外副作用。此 PR 确保取消后立即丢弃所有待处理的工具调用。
   - **链接:** https://github.com/google-gemini/gemini-cli/pull/28096

2. **[#28094] fix(a2a-server): deep-merge user and workspace settings (修复A2A服务器：深度合并用户和工作区设置)**
   - **功能:** 修复了配置管理中一个隐蔽的 Bug。之前的浅合并（shallow merge）会导致工作区设置完全覆盖用户设置，使得部分配置项丢失。此 PR 改为深度合并，保证了配置的完整性。
   - **链接:** https://github.com/google-gemini/gemini-cli/pull/28094

3. **[#28093] fix(core): buffer chat compression telemetry until SDK is initialized (修复核心：缓存聊天压缩遥测直到SDK初始化)**
   - **功能:** 修复了启动阶段的遥测崩溃问题。在 SDK 完全初始化前，如果触发了遥测上报，会导致程序崩溃。此 PR 通过引入缓冲机制解决了启动时序问题。
   - **链接:** https://github.com/google-gemini/gemini-cli/pull/28093

4. **[#28089] feat(core): implement MCP elicitation (form + url) capability (特性：实现MCP表单和URL引导功能)**
   - **功能:** 这是一个重要的功能增强。实现了 MCP 协议中定义的“引导”（表单和URL模式）能力，使 CLI 能够与支持该功能的 MCP 服务器进行更复杂的交互。
   - **链接:** https://github.com/google-gemini/gemini-cli/pull/28089

5. **[#28068] fix(core): guard message inspectors against empty parts arrays (修复核心：防止消息检查器处理空内容数组)**
   - **功能:** 修复了一个因 JavaScript 特性导致的逻辑错误。当消息内容（`parts`）为空时，相关检查函数会错误地返回 `true`，可能引发后续处理异常。此 PR 增加了防御性检查。
   - **链接:** https://github.com/google-gemini/gemini-cli/pull/28068

6. **[#28053] fix(core-tools): resolve defensive path resolution for at-reference files (修复核心工具：解决@引用文件的防御性路径解析)**
   - **功能:** 修复了一个生产环境 Bug：当模型在路径前加入 `@` 符号（如 `@policies/new-policies.txt`）时，文件操作工具（如 `read_file`）会找不到文件。此 PR 实现了更健壮的路径解析逻辑。
   - **链接:** https://github.com/google-gemini/gemini-cli/pull/28053

7. **[#28000] fix(core-tools): resolve Jupyter Notebook and JSON corruption in write_file (修复核心工具：解决write_file损坏Jupyter Notebook和JSON文件)**
   - **功能:** 修复了一个严重的文件损坏 Bug。`write_file` 工具在处理 `.ipynb` 和 `.json` 文件时会静默地破坏其格式，导致文件无法使用。此 PR 预计将大幅改善数据科学工作流的稳定性。
   - **链接:** https://github.com/google-gemini/gemini-cli/pull/28000

8. **[#27915] fix(core): trust dialog discloses the hook shape that never runs (修复核心：信任对话框显示错误的Hook信息)**
   - **功能:** **安全修复**。工作区信任对话框显示的是“不会执行”的 Hook 命令，而非实际要执行的命令，给用户造成安全错觉。此 PR 修复了显示逻辑，确保用户看到的是将要执行的命令。
   - **链接:** https://github.com/google-gemini/gemini-cli/pull/27915

9. **[#27916] fix(core): validate GCP project ID format and prevent alias extraction in memory (修复核心：验证GCP项目ID格式并防止记忆提取别名)**
   - **功能:** 修复了 Auto Memory 在提取 GCP 项目信息时，错误地存储了显示名称或别名，而非有效的项目 ID，导致后续调用返回 403 错误。增加了格式校验。
   - **链接:** https://github.com/google-gemini/gemini-cli/pull/27916

10. **[#27910] fix(core): bound web search tool latency (修复核心：限制网络搜索工具延迟)**
    - **功能:** 为 `google_web_search` 工具增加了 120 秒的超时限制。解决了一些情况下搜索请求无限等待的问题，让 Agent 能迅速从超时中恢复并通知用户。
    - **链接:** https://github.com/google-gemini/gemini-cli/pull/27910

## 功能需求趋势

从近期的 Issues 中可以提炼出社区的三大关注方向：

1.  **Agent 行为的可靠性与安全性**：社区不再满足于“能用”，而是追求“可靠”和“安全”。这体现在对 Agent 滥用 `--force` 等危险命令的担忧（#22672），以及对 Auto Memory 功能可能泄露敏感信息（#26525）的警惕上。用户希望在赋予 Agent 强大能力的同时，有完善的护栏机制。

2.  **深入的代码理解与工具集成**：用户不再满足于简单的文件读写，而是希望 Agent 能“理解”代码。**AST（抽象语法树）** 的引入（#22745）是一个明确的信号，旨在让 Agent 能像开发者一样定位到精确的方法、类。同时，与 **MCP 协议**（#28089）的集成也表明，社区期望 CLI 能成为一个强大的平台，接入更多专业工具。

3.  **彻底的自动化与性能优化**：自动记忆（Auto Memory）反复重试低价值任务（#26522）和 Agent 不主动调用已配置的技能（#21968），这些 Issue 都指向了自动化的“智能”程度不足。社区期望真正的“放手”式自动化，并希望系统更高效地处理任务，避免资源浪费和无效操作。

## 开发者关注点

近期社区反馈的核心痛点和需求主要集中在：

- **Agent “卡住”和“死锁”问题**：无论是通用 Agent 挂起（#21409）还是 Shell 命令执行后卡住（#25166），都是开发者最常遇到的挫折。这直接破坏了工作流程，是用户流失的潜在风险点。
- **配置管理的“意外”行为**：不论是子代理错误报告成功（#22323），还是浏览器代理忽略 `settings.json` 配置（#22267），都暴露出配置系统不够可靠和一致。开发者希望配置即是事实，行为可预期。
- **“不智能”的自动化决策**：Agent 不主动使用定制技能（#21968）或者在随机位置创建临时脚本（#23571），让开发者觉得它“不够聪明”。这提示 Agent 的执行规划和工具选择模型需要大幅优化，以更贴近真实开发场景。
- **终端体验的细碎问题**：虽然不致命，但终端大小变化时的闪烁（#21924）、退出外部编辑器后屏幕损坏（#24935）、JSON 序列化错误（#22466）等问题持续存在，影响了日常使用的舒适度。
- **安全与权限的隐忧**：工作区信任对话框的 Bug（#27915）和自动记忆的问题（#26525, #26522）让开发者对 CLI 在敏感项目中的使用产生顾虑。透明的操作和明确的权限控制是赢得信任的关键。

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI 社区动态日报 | 2026-06-23

> 数据来源：github.com/github/copilot-cli | 统计截止 2026-06-23 08:00 UTC

---

## 今日速览

过去 24 小时内，Copilot CLI 仓库活跃度较高，共更新 22 条 Issue/PR。社区关注焦点集中在 **认证会话恢复失败**、**Windows ARM64 稳定性** 及 **MCP 服务器匹配误报** 三大问题上。此外，多项与插件结构优化、终端 UI 计时、沙盒文档不一致相关的功能请求也获得了较多讨论。

---

## 版本发布

过去 24 小时内无新版本发布。当前最新版本为 **v1.0.63**（发布于 2026-06-18）。

---

## 社区热点 Issues（10 项）

精选过去 24 小时内更新、讨论热度高或影响范围广的 Issue，按关注度排序。

### 🔥 1. #3596 – 恢复会话时出现 “Not authenticated” 错误（👍 11 | 💬 6）
**标签**：`area:authentication`, `area:sessions`, `area:models`  
**摘要**：用户在使用 `/resume` 恢复指定 session 后，`/model` 命令无法列出模型，报错 “Not authenticated”，但新建 session 则正常。触发多模型选择的用户频繁中招，严重影响工作流。  
➡️ [查看详情](https://github.com/github/copilot-cli/issues/3596)

### 🔥 2. #3162 – CLI 1.0.42 误报自定义 MCP 服务器被策略阻止（已关闭，👍 1 | 💬 7）
**标签**：`area:mcp`  
**摘要**：已存在于 MCP Registry 中的自定义服务器被 CLI 错误标记为 “blocked by policy”，属于 Registry 验证逻辑的误报。虽然已关闭（可能已修复），但暴露了 MCP 匹配策略的脆性。  
➡️ [查看详情](https://github.com/github/copilot-cli/issues/3162)

### 🔥 3. #3687 – Windows ARM64 下 copilot.exe 在高负载时崩溃（BEX64）（👍 1 | 💬 6）
**标签**：`area:sessions`, `area:platform-windows`  
**摘要**：重现率高：同时启动多个 session（如 Windows Terminal 标签页恢复）或内存压力下，`copilot.exe` 以 `0xc0000409` 硬错误退出，无法优雅关闭。ARM64 用户持续性受影响。  
➡️ [查看详情](https://github.com/github/copilot-cli/issues/3687)

### 🔥 4. #1632 – 支持技能（skills）子文件夹以更好组织（👍 20 | 💬 8）
**标签**：`area:plugins`  
**摘要**：当前 skills 文件夹只支持扁平结构，用户希望将 10+ 个自定义 skill 按测试、开发等分类放入子目录。社区点赞数最高，是插件体系的重要增强。  
➡️ [查看详情](https://github.com/github/copilot-cli/issues/1632)

### 🔥 5. #3854 – `@` 文件引用自动补全失效（👍 0 | 💬 2）
**标签**：`area:input-keyboard`  
**摘要**：自最近几个版本起，在允许目录内输入 `@` 加文件名时不再出现自动补全提示，只显示根目录和临时文件夹，严重影响引用文件效率。  
➡️ [查看详情](https://github.com/github/copilot-cli/issues/3854)

### 🔥 6. #2399 – 插件安装应使用稀疏检出（sparse checkout）（👍 0 | 💬 2）
**标签**：`area:plugins`, `area:installation`  
**摘要**：当前插件安装会完整 `git clone` 整个仓库，下载测试、CI 配置、图片等非必需文件。建议只拉取插件运行所需的资产，减少带宽和磁盘占用。  
➡️ [查看详情](https://github.com/github/copilot-cli/issues/2399)

### 🔥 7. #2337 – WSL 下应使用 git-credential-manager 安全存储 token（👍 2 | 💬 1）
**标签**：`area:platform-windows`  
**摘要**：WSL 上登录时显示警告，但用户已配置 git-credential-manager。CLI 当前未利用该工具存储 token，存在安全隐患，且用户体验割裂。  
➡️ [查看详情](https://github.com/github/copilot-cli/issues/2337)

### 🔥 8. #3861 – 沙盒文档与真实行为不符（per-host 过滤等无效）（👍 0 | 💬 1）
**标签**：`area:permissions`, `area:networking`  
**摘要**：文档声称支持 per-host 网络过滤（`allowedHosts`/`blockedHosts`）及跨平台隔离，但实际并未生效。用户请求更新文档或修复功能。  
➡️ [查看详情](https://github.com/github/copilot-cli/issues/3861)

### 🔥 9. #3886 – `/restart` 消耗大量 AI credits（👍 0 | 💬 0）
**标签**：`area:sessions`, `area:models`  
**摘要**：近几个版本中，`/restart`、`/resume` 甚至 `/update` 会固定消耗约 174 AI credits（而非预期的小额）。用户反馈配额被异常占用。  
➡️ [查看详情](https://github.com/github/copilot-cli/issues/3886)

### 🔥 10. #3883 – 功能请求：支持 i18n 国际化（👍 1 | 💬 0）
**标签**：`area:theming-accessibility`  
**摘要**：建议为 CLI 界面（菜单、提示、错误信息等）添加国际化支持，优先翻译全球使用人数最多的前 10 种语言，降低非英语用户门槛。  
➡️ [查看详情](https://github.com/github/copilot-cli/issues/3883)

---

## 重要 PR 进展

过去 24 小时内仅有 1 个 PR 更新：

### #3880 – [测试/无效 PR]（👍 0 | 💬 0）
**摘要**：该 PR 标题为 “beyond the streets of amaerica”，内容包含一个不相关的 React 组件示例。极可能为垃圾或测试 PR，无实际功能贡献。**今日无实质性 PR 进展。**  
➡️ [查看详情](https://github.com/github/copilot-cli/pull/3880)

---

## 功能需求趋势

综合所有 Issue，社区最关注的功能方向如下：

| 方向 | 代表 Issue | 说明 |
|------|------------|------|
| **插件系统增强** | #1632（子文件夹）、#2399（稀疏检出） | 插件/技能的组织与安装效率是当前最大诉求 |
| **MCP 互联生态** | #3162（误报）、#3638（VS Code 集成） | MCP 服务器配置在 CLI 与 IDE 之间的无缝共享 |
| **会话与认证稳定性** | #3596（认证失败）、#3886（信用消耗） | 恢复会话时的鉴权漏洞、AI 信用计费异常 |
| **Windows 平台兼容性** | #3687（ARM64 崩溃）、#2337（WSL 凭证） | ARM64 稳定性、WSL 密钥存储 |
| **终端 UI 与国际化** | #3111/#3278/#3055（计时）、#3885（滚动）、#3883（i18n） | 更透明的时间反馈、长文本滚动、多语言支持 |
| **沙盒与权限文档对齐** | #3861（沙盒功能无效）、#3884（企业策略） | 文档描述与实际行为不一致，企业用户需明确指引 |

---

## 开发者关注点

1. **认证会话脆弱**：`/resume` 后模型列表无法加载（#3596），且重启操作异常消耗大量 credits（#3886），严重影响多session工作流。
2. **MCP 验证逻辑待打磨**：自定义 MCP 服务器被误拦（#3162），MCP 初始化指令被忽略（#1579），降低 MCP 生态的可信度。
3. **输入交互退化**：`@` 文件引用自动补全失效（#3854），长文本无法滚动（#3885），基础编辑体验回退。
4. **Windows 用户痛点尖锐**：ARM64 高负载崩溃（#3687）、WSL 凭证存储不兼容（#2337），Windows 平台用户稳定性亟待提升。
5. **沙盒功能名不副实**：per-host 过滤无效（#3861），文档宣称的能力无法复现，企业管理员缺乏配置依据。
6. **计费透明度不足**：模型 multiplier 计算异常（#3881），重启消耗固定额度（#3886），用户难以预期信用消耗。

---

*以上分析基于仓库公开数据，如需查看完整列表，请访问 [github.com/github/copilot-cli/issues](https://github.com/github/copilot-cli/issues)。*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

好的，作为专注于 AI 开发工具的技术分析师，以下是根据最新 GitHub 数据生成的 **2026-06-23 Kimi Code CLI 社区动态日报**。

---

# Kimi Code CLI 社区动态日报 | 2026-06-23

## 今日速览

Kimi Code CLI 发布 **v1.48.0**，主要修复了 KOSONG 推理内容为空时的往返问题，并增强了“灵魂”模块对重复工具调用的终止与提醒机制。社区本月集中反馈了 **MCP 服务器配置管理** 的多项缺陷（自动发现、路径偏差、ACP 模式失效），以及 **KOSONG 推理参数** 的兼容性问题，开发者对内存系统和子进程稳定性也提出了明确需求。

---

## 版本发布

### v1.48.0（最新）
- **发布标签**: `1.48.0`（同时更新 KOSONG 至 `0.54.0`）
- **主要变更**:
  - `fix(kosong)`: 修复空推理内容往返时导致的问题 [#2446](https://github.com/MoonshotAI/kimi-cli/pull/2446)
  - `feat(soul)`: 当工具调用重复超过 3 次时，逐级注入提醒（r1/r2/r3），并在陷入死循环后强制停止当前处理 [#2466](https://github.com/MoonshotAI/kimi-cli/pull/2466)
  - 内部构建与依赖同步（KOSONG 0.54.0）

---

## 社区热点 Issues（共 6 条，全部附链接）

### [#1283] [增强] 功能请求：内存系统 —— 跨会话持久化上下文
- **作者**: CatKang | **创建**: 2026-02-27 | **更新**: 2026-06-22 | **评论**: 6
- **链接**: [Issue #1283](https://github.com/MoonshotAI/kimi-cli/issues/1283)
- **重要性**：虽为旧 Issue 但近期再次活跃，社区对跨会话上下文记忆的呼声很高。用户希望 CLI 能自动记住项目模式与用户偏好，甚至支持手动编辑记忆。
- **社区反应**：已有 6 条讨论，开发者未明确回复，但 Issue 被重新标记为 open 状态。

### [#2457] [Bug] Kimi Code CLI 自动发现已删除的 MCP 服务器，导致 400 错误无法修复
- **作者**: xavier2sy8827-cmyk | **创建**: 2026-06-16 | **更新**: 2026-06-22 | **评论**: 1
- **链接**: [#2457](https://github.com/MoonshotAI/kimi-cli/issues/2457)
- **重要性**：严重 bug，用户删除 MCP 配置后 CLI 仍自动加载已删除的服务器，导致反复 400 错误，且无法通过常规方式清除。
- **社区反应**：评论中建议手动清理配置文件，但用户表示无效，期待官方补丁。

### [#2469] [Bug] `kimi web` 从 CLI 安装目录启动 MCP 服务器，破坏工作区相对路径工具
- **作者**: Zehee | **创建**: 2026-06-22 | **更新**: 2026-06-22 | **评论**: 0
- **链接**: [#2469](https://github.com/MoonshotAI/kimi-cli/issues/2469)
- **重要性**：影响使用工作区相对路径的 MCP 工具（如 `./scripts/mcp-server.js`）。启动时 CWD 错误导致工具找不到资源。
- **社区反应**：刚提交，尚无回复。

### [#2468] [Bug] Kimi CLI 在分离子进程工具调用后挂起
- **作者**: N0zoM1z0 | **创建**: 2026-06-22 | **更新**: 2026-06-22 | **评论**: 0
- **链接**: [#2468](https://github.com/MoonshotAI/kimi-cli/issues/2468)
- **重要性**：影响使用本地 API 模拟提供者的用户，子进程分离后 CLI 不再响应，可能阻塞自动化流程。
- **社区反应**：暂无回复，但描述详细，包含复现步骤。

### [#2465] [Bug] KOSONG: OpenAILegacy 在思考关闭时发送 `reasoning_effort: null`，导致严格 API 报错且无法真正禁用推理
- **作者**: 0xbentang | **创建**: 2026-06-22 | **更新**: 2026-06-22 | **评论**: 0
- **链接**: [#2465](https://github.com/MoonshotAI/kimi-cli/issues/2465)
- **重要性**：严格兼容性 bug，违反 OpenAI 聊天补全 Schema（null 非法），且即使设置 off 也未能真正禁用思考，导致模型始终输出推理内容。
- **社区反应**：刚提交，暂无回复。

### [#2464] [Bug] `kimi acp` 不加载 MCP 服务器 —— ACP 模式下 MCP 工具缺失
- **作者**: Tasktivity | **创建**: 2026-06-22 | **更新**: 2026-06-22 | **评论**: 0
- **链接**: [#2464](https://github.com/MoonshotAI/kimi-cli/issues/2464)
- **重要性**：ACP（Agent Communication Protocol）模式完全忽略 `--mcp-config-file` 参数，导致该模式无法使用任何自定义 MCP 工具，影响集成开发。
- **社区反应**：新报告，开发团队暂未响应。

---

## 重要 PR 进展（共 2 条，全部附链接）

### [#2467] [已合并] chore(release): bump kimi-cli to 1.48.0 and kosong to 0.54.0
- **作者**: sailist | **创建/更新**: 2026-06-22
- **链接**: [PR #2467](https://github.com/MoonshotAI/kimi-cli/pull/2467)
- **功能**：版本号晋升及依赖同步，不包含公开 Changelog。
- **意义**：标志着 v1.48.0 正式发布，修复了先前版本的 KOSONG 和 Soul 问题。

### [#2466] [已合并] feat(soul): escalate repeated-tool-call reminders and force-stop on dead-end streak
- **作者**: jackfish212 | **创建/更新**: 2026-06-22
- **链接**: [PR #2466](https://github.com/MoonshotAI/kimi-cli/pull/2466)
- **功能**：从 kimi-code 移植重复工具调用处理逻辑：连续重复 3 次后注入逐步升级的提醒（r1/r2/r3），达到死胡同次数后强制停止本轮调用。
- **意义**：提升智能体稳定性，避免无限循环消耗 Token 或陷入无响应。

---

## 功能需求趋势

从近期 Issues 和社区反馈来看，开发者的核心关注方向集中在：

1. **MCP 服务器配置管理**（共 4 个 Bug 直接相关）：
   - 自动发现残留服务器（#2457）
   - 启动工作路径错误（#2469）
   - ACP 模式完全忽略 MCP 配置（#2464）
   - 部分用户还要求支持动态加载/卸载 MCP 工具。

2. **推理参数兼容性与灵活性**（#2465）：
   - 用户希望 KOSONG 的 `reasoning_effort` 能严格遵循 OpenAI 规范，同时真正关闭思考时不再发送多余参数。

3. **持久化记忆与上下文管理**（#1283）：
   - 虽然为较旧 Issue，但近期仍有讨论，表明用户对于“CLI 能记住项目习惯和偏好”的期待未减。

4. **子进程与系统稳定性**（#2468）：
   - 对于使用本地 Mock 或自定义后端的开发者，分离子进程后的挂起问题是阻碍自动化的痛点。

---

## 开发者关注点（痛点与高频需求）

| 痛点 / 需求 | 来源 | 严重程度 |
|-------------|------|----------|
| **ACP 模式完全不支持 MCP** | #2464 | 🚫 高 – 核心功能缺失 |
| **MCP 服务器自动装载残留** | #2457 | ⚠️ 中 – 错误不可修复 |
| **kimi web 启动 MCP 路径错位** | #2469 | ⚠️ 中 – 破坏工作区相对路径 |
| **KOSONG 发送不合规 null 值** | #2465 | ⚠️ 中 – 兼容性风险 |
| **子进程分离后 CLI 挂起** | #2468 | 🔴 高 – 阻塞自动化 |
| **跨会话记忆系统缺失** | #1283 | 💡 中等需求 |

> 总结：**MCP 基础设施的稳定性** 是当前社区最迫切的交火点，建议官方在下个迭代集中修复路径、加载生命周期和模式兼容性问题。同时 KOSONG 层的 API 严格合规性也需要尽快补丁，以免影响第三方集成。

---

*数据来源：[MoonshotAI/kimi-cli](https://github.com/MoonshotAI/kimi-cli) | 日报生成时间：2026-06-23*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode 社区动态日报 | 2026-06-23

---

## 1. 今日速览

截至2026年6月22日，社区共产生36条议题更新和50条PR更新。**MCP工具参数序列化**（#28472）和**Todo面板刷新问题**（#33063）成为讨论最多的bug；**DeepSeek V4 Pro官方提供者无响应**（#33395）引发用户困扰。PR方面，**子会话事件NDJSON流**（#33403）和**Codex模型列表过滤**（#33400）是两项关键修复。

---

## 2. 版本发布

无新版本。

---

## 3. 社区热点 Issues（精选10条）

### 🔥 #28472 MCP工具`object`类型参数被序列化为字符串
- **重要性**：核心MCP交互bug，导致所有含`body`（type: object）的MCP工具调用失败。
- **链接**：[Issue #28472](https://github.com/anomalyco/opencode/issues/28472)

### 🔥 #33395 DeepSeek V4 Pro（Max）在官方提供者中无响应
- **重要性**：用户反馈官方DeepSeek提供者输出为空，重装才能临时恢复，OpenRouter正常。社区暂无解决方案。
- **链接**：[Issue #33395](https://github.com/anomalyco/opencode/issues/33395)

### 🔥 #33063 Todo Dock UI在`todowrite`后不刷新
- **重要性**：SolidJS响应性缺陷，导致Todo面板始终显示初始数据。影响开发工作流。
- **链接**：[Issue #33063](https://github.com/anomalyco/opencode/issues/33063)

### 🔥 #32435 Codex (ChatGPT账号) 显示不可用的`-pro`模型
- **重要性**：界面提供`gpt-5.5-pro`但请求时被后端拒绝，用户体验差。已通过PR #33400修复。
- **链接**：[Issue #32435](https://github.com/anomalyco/opencode/issues/32435)

### 🔥 #25785 AI应尊重GitHub仓库模板并避免语言混杂
- **重要性**：用户期望AI在贡献时自动发现`PR_TEMPLATE.md`、`CONTRIBUTING.md`，并保持语言一致。PR #31989已部分解决。
- **链接**：[Issue #25785](https://github.com/anomalyco/opencode/issues/25785)

### 🔥 #25872 为已发送消息增加编辑、撤回、删除按钮
- **重要性**：高频需求，用户无法修正错误或撤回消息，影响聊天效率。
- **链接**：[Issue #25872](https://github.com/anomalyco/opencode/issues/25872)

### 🔥 #33077 Bash工具无防护，可执行`git reset --hard`等破坏性命令
- **重要性**：安全漏洞，用户误操作可能丢失工作。社区已在讨论权限模型改进。
- **链接**：[Issue #33077](https://github.com/anomalyco/opencode/issues/33077)

### 🔥 #25351 `/models`命令显示预设列表而非真实LM Studio模型
- **重要性**：本地提供者集成bug，开发者无法看到实际加载的模型。
- **链接**：[Issue #25351](https://github.com/anomalyco/opencode/issues/25351)

### 🔥 #33074 WebFetch Cloudflare回退使用`opencode` User-Agent
- **重要性**：反直觉设计，回退时使用非标准UA反而更容易被拦截，导致抓取失败。
- **链接**：[Issue #33074](https://github.com/anomalyco/opencode/issues/33074)

### 🔥 #33399 opencode进程随机占用99-100% CPU并失去响应
- **重要性**：性能严重退化，用户从v1.3.3开始遇到，CLI无法接受键盘输入。
- **链接**：[Issue #33399](https://github.com/anomalyco/opencode/issues/33399)

---

## 4. 重要 PR 进展（精选10条）

### 🚀 #33403 feat(run): 将子会话事件转发至NDJSON流
- **内容**：使`opencode run --format json`输出子agent的中间事件，提升可观测性。
- **链接**：[PR #33403](https://github.com/anomalyco/opencode/pull/33403)

### 🚀 #33400 fix: 使用Codex OAuth时不显示gpt-5.5-pro
- **内容**：直接修复#32435，从模型列表中排除ChatGPT账号不可用的`-pro`模型。
- **链接**：[PR #33400](https://github.com/anomalyco/opencode/pull/33400)

### 🚀 #33374 fix(prompt): 清理任务时保护备份/凭据不被删除
- **内容**：避免AI在“清理旧文件”时误删用户备份和凭据文件。
- **链接**：[PR #33374](https://github.com/anomalyco/opencode/pull/33374)

### 🚀 #33393 fix(core): 区分WebFetch URL错误与网络错误
- **内容**：修复URL验证失效的问题（#33073），正确区分格式错误与网络失败。
- **链接**：[PR #33393](https://github.com/anomalyco/opencode/pull/33393)

### 🚀 #33392 feat(llm): 在工具定义中传递`strict`参数以实现Codex对等
- **内容**：为所有函数工具默认设置`strict: false`，避免动态Schema被OpenAI API拒绝。
- **链接**：[PR #33392](https://github.com/anomalyco/opencode/pull/33392)

### 🚀 #33227 fix(opencode): 改进TUI上的Workspace功能
- **内容**：允许在运行会话时切换workspace，并修复名称显示问题。
- **链接**：[PR #33227](https://github.com/anomalyco/opencode/pull/33227)

### 🚀 #33160 fix(mcp): 防止OpenAI兼容提供者中MCP参数为null
- **内容**：当参数只有`description`无`type`时，MiniMax等模型会收到`null`，已修复。
- **链接**：[PR #33160](https://github.com/anomalyco/opencode/pull/33160)

### 🚀 #33082 docs(rfc): Computer Use for opencode
- **内容**：RFC请求设计对齐，探讨在opencode中实现计算机操作能力（如桌面自动化）。
- **链接**：[PR #33082](https://github.com/anomalyco/opencode/pull/33082)

### 🚀 #32075 feat(core): 添加可配置的计划提醒
- **内容**：允许用户覆盖“请先制定计划”的系统提示，满足不同工作流需求。
- **链接**：[PR #32075](https://github.com/anomalyco/opencode/pull/32075)

### 🚀 #29217 feat(tui): 在提示编辑器中添加内联`$skill`调用
- **内容**：输入`$`自动补全技能，并支持粘贴文本，极大提升便捷性。
- **链接**：[PR #29217](https://github.com/anomalyco/opencode/pull/29217)

---

## 5. 功能需求趋势

从近期议题来看，社区关注的功能方向集中为：

1. **MCP生态完善**：参数传递（object序列化、高精度数字丢失）、传输协议兼容性（Accept header）、远程服务器连接。
2. **安全与权限**：预执行校验、破坏性命令防护、目录遍历保护、权限级联逻辑优化。
3. **用户界面增强**：消息编辑/撤回、命令选择器、消息翻译、缩放适配、会话归档可视化。
4. **模型支持与兼容性**：DeepSeek V4 Pro官方提供者、ChatGPT账号模型列表过滤、LM Studio本地模型动态加载。
5. **工作流自动化**：预设指令系统、工厂重置/缓存清理、`/init`引导、子会话事件暴露、计划提醒可配置。
6. **性能与稳定性**：CPU占用异常、Todo Dock刷新、子进程清理TOCTOU竞争。

---

## 6. 开发者关注点

- **高频痛点**：MCP参数序列化错误直接阻塞工具调用；Todo面板不刷新导致状态不一致；DeepSeek官方提供者不可用；`/models`命令与真实模型脱节。
- **安全担忧**：Bash工具可执行任意命令、Glob/Grep允许目录遍历、WebFetch URL验证失效、权限拒绝级联到所有请求——这些问题被多位开发者指出并提交了对应的修复PR。
- **性能退化**：部分用户遭遇opencode进程100% CPU占用且无响应，影响开发效率。
- **交互缺失**：无法编辑/删除已发送消息、没有命令快速选择按钮、没有归档会话的查看入口——这些Feature请求获得社区共鸣。

---

*数据截止时间：2026-06-22 23:59 UTC。日报由AI助手自动生成，请以GitHub仓库实际状态为准。*

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

# 🐚 Pi 社区动态日报 | 2026-06-23

---

## 1. 今日速览

- **v0.79.10 发布**：扩展事件 `session_before_compact` 和 `session_compact` 新增 `reason` 与 `willRetry` 字段，让扩展能区分手动 `/compact`、阈值自动压缩和溢出重试场景。
- **连接可靠性问题持续发酵**：Issue #4945 关于 `openai-codex/gpt-5.5` 在交互 TUI 中卡在 `Working...` 的问题已获 64 条评论，成为社区最关注的热点。
- **安全与隐私议题升温**：多起报告指出代理可能意外泄露密钥文件（#5956）或在宽泛任务中复制敏感文件，已通过 PR #5955 修正默认系统提示。

---

## 2. 版本发布

### v0.79.10
**链接**：https://github.com/earendil-works/pi/releases/tag/v0.79.10

**主要特性**：
- **扩展压缩事件上下文增强**：`session_before_compact` 和 `session_compact` 事件现在包含 `reason`（`"manual" | "threshold" | "overflow"`）和 `willRetry`（布尔值）。扩展开发者可以据此精确判断压缩触发原因并调整行为。
- **关联 PR**：#5962、#5941 已合并关闭。

---

## 3. 社区热点 Issues（精选 10 个）

### 🔥 #4945 — `openai-codex` 连接可靠性问题
- **状态**：OPEN / inprogress  
- **评论数**：64 | 👍 30  
- **链接**：https://github.com/earendil-works/pi/issues/4945  
- **摘要**：`openai-codex/gpt-5.5` 交互 TUI 常卡在 `Working...`，无法接收流式文本、工具调用或错误信息，只能按 Esc 中止。问题在过去数天反复出现，社区高度关注，亟需解决。

### 🔥 #3357 — 官方本地 LLM 提供商扩展
- **状态**：OPEN  
- **评论数**：26 | 👍 36  
- **链接**：https://github.com/earendil-works/pi/issues/3357  
- **摘要**：希望 Pi 能动态从 `{baseUrl}/models` 拉取模型列表，以便无缝对接 llama.cpp / ollama / LM Studio 等本地推理引擎。社区呼声极高（36 个赞），是本地化部署的关键需求。

### 🔥 #5653 — 从 Shrinkwrap 迁移
- **状态**：OPEN / inprogress, to-discuss  
- **评论数**：15 | 👍 0  
- **链接**：https://github.com/earendil-works/pi/issues/5653  
- **摘要**：同时安装 `pi-ai` 和 `pi-coding-agent` 会导致两份相同的 `pi-ai` 副本（一个提升到根、一个嵌套），造成模块级 `Map` 分离，引发 API 提供者注册失效。建议依赖管理方案重构。

### 🔥 #5916 — 支持提供商扩展的模型别名并改进搜索
- **状态**：OPEN / bug, inprogress  
- **评论数**：11 | 👍 0  
- **链接**：https://github.com/earendil-works/pi/issues/5916  
- **摘要**：用户为 OpenRouter 配置 `modelOverrides` 后，模型别名在选择器中仍无显示，且无法搜索别名。需要 UI/搜索逻辑改进。

### 🔥 #5571 — `pi -p` 在非 TTY 管道未关闭时无限挂起
- **状态**：CLOSED  
- **评论数**：10 | 👍 0  
- **链接**：https://github.com/earendil-works/pi/issues/5571  
- **摘要**：当 stdin 是非 TTY 管道且不会关闭时，`pi -p` 由于默认提供商无凭证而永久挂起，而非快速失败。已修复（相关 PR 未列出，但 Issue 已关闭）。

### 🔥 #5778 — pi-agent-core 在无响应流或死锁时无限挂起
- **状态**：CLOSED  
- **评论数**：8 | 👍 0  
- **链接**：https://github.com/earendil-works/pi/issues/5778  
- **摘要**：代理循环中若 LLM 流断开未关闭迭代器，或工具 `execute()` 返回未 resolve 的 Promise，会导致永久挂起。已通过补丁修复。

### 🔥 #5939 — 让自动压缩变为可选且安全执行
- **状态**：CLOSED / no-action  
- **评论数**：7 | 👍 0  
- **链接**：https://github.com/earendil-works/pi/issues/5939  
- **摘要**：建议将自动压缩设为默认关闭，仅在显式启用后、在助手工具调用回合完成且下一个提供商请求开始前安全执行。虽未采纳，但反映了社区对上下文管理可控性的重视。

### 🔥 #4748 — TUI 键绑定单例问题破坏扩展
- **状态**：OPEN  
- **评论数**：5 | 👍 2  
- **链接**：https://github.com/earendil-works/pi/issues/4748  
- **摘要**：扩展加载时解析到独立的 `@earendil-works/pi-tui` 副本，造成 `getKeybindings()` 模块级单例失效。影响所有导入 `keyText` 的扩展。

### 🔥 #5263 — 使会话内模型和思考级别更改默认为临时
- **状态**：OPEN  
- **评论数**：4 | 👍 4  
- **链接**：https://github.com/earendil-works/pi/issues/5263  
- **摘要**：当前会话内切换模型会全局生效，用户期望仅影响当前会话。建议在 `/settings` 中增加“默认模型”选项，提供明确的全局覆盖入口。

### 🔥 #5871 — Anthropic OAuth token 检测硬编码
- **状态**：OPEN / inprogress  
- **评论数**：4 | 👍 0  
- **链接**：https://github.com/earendil-works/pi/issues/5871  
- **摘要**：Anthropic 提供者只通过 `sk-ant-oat` 前缀判断 OAuth token，不支持显式声明。导致使用非标准前缀的 token 被误判。

---

## 4. 重要 PR 进展（精选 10 个）

### 📌 #5262 — [feat] 添加 Anthropic Vertex 提供商
- **状态**：OPEN  
- **链接**：https://github.com/earendil-works/pi/pull/5262  
- **摘要**：新增内置 `anthropic-vertex` 提供者，用于 Google Cloud Vertex AI 上的 Claude。薄层适配器复用现有 Anthropic 流式处理路径，支持请求/流/工具/思考处理。对 GCP 用户意义重大。

### ✅ #5970 — [feat] DeepSeek V4 Pro/Flash 自动路由扩展
- **状态**：CLOSED  
- **链接**：https://github.com/earendil-works/pi/pull/5970  
- **摘要**：添加 `auto-router.ts` 扩展，根据 prompt 复杂度自动在 DeepSeek V4 Flash（简单任务）和 V4 Pro（复杂任务）之间路由，可节省 60-70% API 成本。社区期待的成本优化方案。

### ✅ #5962 — [feat] 给扩展压缩事件添加 reason 和 willRetry
- **状态**：CLOSED（合并到 v0.79.10）  
- **链接**：https://github.com/earendil-works/pi/pull/5962  
- **摘要**：实现 v0.79.10 核心特性。`session_before_compact` 和 `session_compact` 事件携带压缩原因（手动/阈值/溢出）和重试标记，与 RPC 协议对齐。关闭 #5217。

### ✅ #5963 — [fix] 拒绝格式错误的最终工具调用参数
- **状态**：CLOSED  
- **链接**：https://github.com/earendil-works/pi/pull/5963  
- **摘要**：在共享 AI 流路径中验证最终的工具调用参数 JSON。若参数格式错误（如字符串代替对象），则触发 `stopReason: "error"` 而非暴露非法工具调用。提升代理可靠性。

### ✅ #5941 — [fix] 给压缩扩展事件添加必要字段（与 #5962 相似但较早版本）
- **状态**：CLOSED  
- **链接**：https://github.com/earendil-works/pi/pull/5941  
- **摘要**：早期实现相同功能，后与 #5962 整合。确保扩展 API 一致。

### ✅ #5955 — [fix] 默认系统提示增加密钥泄露范围约束
- **状态**：CLOSED  
- **链接**：https://github.com/earendil-works/pi/pull/5955  
- **摘要**：解决 Issue #5956 报告的安全问题：代理执行“复制所有文件”任务时会将 `.env` 或私钥等秘密文件复制到公开目录。现在系统提示明确约束不可暴露机密文件。

### ✅ #5950 — [fix] 在底部栏使用 OpenRouter 实际费用
- **状态**：CLOSED  
- **链接**：https://github.com/earendil-works/pi/pull/5950  
- **摘要**：OpenRouter 响应中返回 `usage.cost` 真实费用，Pi 之前仅使用静态每 token 估算，导致显示偏差。现在改为显示 OpenRouter 实际收费，对 OpenRouter 用户更准确。

### ✅ #5966 — [fix] 修复删除包不生效的问题
- **状态**：CLOSED  
- **链接**：https://github.com/earendil-works/pi/pull/5966  
- **摘要**：`pi remove npm:@foo/bar` 执行后包仍出现在配置中。已修复移除逻辑。

### ✅ #5960 — [fix] `find` 工具忽略嵌套 git 仓库中的文件
- **状态**：CLOSED  
- **链接**：https://github.com/earendil-works/pi/pull/5960  
- **摘要**：当父目录有宽泛 `.gitignore` 时，`find` 工具会忽略嵌套 git 仓库中的文件（如 `submodule/important.txt`）。已修正搜索范围。

### ✅ #5967 — [feat] 支持 OpenCode Go GLM-5.2 最大推理努力
- **状态**：CLOSED  
- **链接**：https://github.com/earendil-works/pi/pull/5967  
- **摘要**：为 `opencode-go/glm-5.2` 模型添加 `xhigh` 思考级别映射到 OpenCode Go 的最大推理 tier。满足用户对该模型深度推理的需求。

---

## 5. 功能需求趋势

从近 24 小时更新的 Issues/PR 中，社区最关注的功能方向如下：

| 方向 | 代表 Issue/PR | 热度 |
|------|--------------|------|
| **本地/自托管 LLM 提供商支持** | #3357（动态模型列表）、#5262（Vertex）、#5916（OpenRouter 别名） | 🟢🟢🟢 |
| **会话管理与扩展 API 增强** | #5932（暴露 `navigateTree()`）、#5810（RPC 暴露会话条目）、#5912（扩展层会话切换）、#5952（安全会话替换 API） | 🟢🟢🟢 |
| **成本优化与自动路由** | #5970（DeepSeek 路由）、#5950（OpenRouter 实际费用） | 🟢🟢 |
| **上下文压缩可控性** | #5939（压缩选项）、#5962/#5941（压缩事件详情） | 🟢🟢 |
| **模型别名与搜索** | #5916（支持别名）、#5972（接受 `provider/model` 格式 `--model`） | 🟢 |
| **安全与隐私** | #5955（机密泄露防护）、#5956（信息披露报告） | 🟢 |
| **WSL/跨平台支持** | #5927（WSL2 路径问题）、#5960（嵌套 git repo） | 🟢 |
| **文档与测试** | #5959、#5957（文档断链修复）、#5974（多行 PR body 创建文档） | 🟢 |

---

## 6. 开发者关注点

- **连接稳定性**：#4945 的 `Working...` 卡死问题持续困扰用户，特别是使用 `openai-codex/gpt-5.5` 时。社区期待官方快速修复。
- **依赖管理冲突**：#5653 揭示的 `pi-ai` 重复安装问题影响扩展开发者，需要依赖提升或重构。
- **扩展开发体验**：#4748（键绑定单例）、#5932（接口不一致）等表明扩展 API 的可预测性和隔离性仍需加强。
- **安全合规**：#5956 的“全部复制”导致机密泄露是高风险场景，幸好已通过 #5955 系统提示修复，但用户仍需注意默认行为。
- **配置可发现性**：#5916 中模型别名无法搜索、#5871 中 OAuth token 检测硬编码，说明 UI 与配置灵活性有待提升。
- **WSL2 用户痛点**：#5927 显示在 Windows WSL2 下通过 `\\wsl.localhost\` 路径启动 Pi 会导致工作目录异常跳转，影响跨平台工作流。

---

## 📅 日报说明

- 数据来源：GitHub 仓库 [earendil-works/pi](https://github.com/earendil-works/pi)（镜像 `badlogic/pi-mono`）
- 数据时间范围：2026-06-22 00:00 UTC ~ 2026-06-23 00:00 UTC（基于 Issue/PR 最后更新时间）
- 统计：过去 24 小时内 1 个 Release、43 条活跃 Issues、7 条活跃 PRs

*本日报由 AI 技术分析师自动生成，供开发社区参考。*

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# Qwen Code 社区动态日报 | 2026-06-23

## 今日速览

昨日社区活跃度高，共产生 18 条 Issue 和 50 个 PR，主要围绕 **安全加固**、**工具调用稳定性** 以及 **配置解析容错性** 展开。三个版本发布（v0.19.0、v0.18.0-preview.3、v0.19.0-preview.0）均因集成测试或发布环节失败而中断，社区正在紧急修复。此外，多个关于 `fastModel` 和自定义 Provider 的改进 PR 进入 review 阶段，用户体验即将优化。

---

## 版本发布

- **v0.18.5-nightly.20260622.6bc3f853e**  
  常规夜间发布，包含 `chore(release)` 与 VSCode 伴侣插件自动发布流程的 CI 改进。  
  [查看发布说明](https://github.com/QwenLM/qwen-code/releases/tag/v0.18.5-nightly.20260622.6bc3f853e)

> ⚠️ 昨日三次正式发布尝试（v0.19.0、v0.18.0-preview.3、v0.19.0-preview.0）均因集成测试或 publish 阶段失败，社区已创建对应 Issue 跟踪（#5663、#5659、#5653），其中 #5663 的修复 PR #5676 已提交。

---

## 社区热点 Issues（10 个）

| Issue | 状态 | 优先级 | 类型 | 要点 | 社区反应 |
|-------|------|--------|------|------|----------|
| [#5634](https://github.com/QwenLM/qwen-code/issues/5634) | OPEN | P2 | bug/security | `autofix` 流程信任由 LLM 自动生成的 `ready-for-agent` 标签，攻击者可伪造标签触发特快修复，绕过人工审核。 | 4 条评论，已 PR #5662 尝试修复 |
| [#5641](https://github.com/QwenLM/qwen-code/issues/5641) | OPEN | P2 | bug/core | 当 provider 重复返回相同的 shell 工具调用结果时，Qwen Code 会陷入无限循环，不断提交已完成的工具结果。 | 4 条评论，附有可复现案例 |
| [#5611](https://github.com/QwenLM/qwen-code/issues/5611) | OPEN | P2 | bug/tools | `web_fetch` 工具仅发送 `text/*` Accept 头，无法获取 JSON API（如 GitHub REST API），收到 415 错误。 | 3 条评论，已有修复 PR #5660 |
| [#5675](https://github.com/QwenLM/qwen-code/issues/5675) | OPEN | P2 | bug/security | `QWEN_CODE_IDE_SERVER_PORT` 未做路径验证，可能被注入路径片段导致读取任意 lock 文件。 | 1 条评论，欢迎 PR |
| [#5664](https://github.com/QwenLM/qwen-code/issues/5664) | OPEN | P2 | bug/core | `stopHookBlockingCap` 环境变量接受小数，被 `Math.floor` 处理，导致阈值意外降低（如 1.5 变为 1）。 | 2 条评论，已有 PR #5667 |
| [#5665](https://github.com/QwenLM/qwen-code/issues/5665) | OPEN | P2 | enhancement | AI 辅助 PR 常更新产品代码和单元测试，但遗漏 `integration-tests/` 下的集成测试更新，直到发布才暴露问题。 | 2 条评论 |
| [#5673](https://github.com/QwenLM/qwen-code/issues/5673) | OPEN | P2 | feature-request | 建议发布工作流在 native 源未变更时复用已有的 `@qwen-code/audio-capture` 预构建制品，避免每次跑全平台构建矩阵。 | 1 条评论 |
| [#5670](https://github.com/QwenLM/qwen-code/issues/5670) | OPEN | type/bug | 当 `fastModel` 配置了 `authType` 前缀（如 `openai:qwen3.6-flash`），`/stats` 和退出摘要中 suggestion 的调用次数被重复计算（原始字符串和解析后名称各一次）。 | 0 条评论，需确认 |
| [#5656](https://github.com/QwenLM/qwen-code/issues/5656) | OPEN | P3 | feature-request | 建议将每次工具调用后的摘要标签（如“Searched in auth/”）从对话历史移至“加载中”指示器区域，减少对话噪音。 | 5 条评论，社区有共鸣 |
| [#5669](https://github.com/QwenLM/qwen-code/issues/5669) | OPEN | P3 | bug/core | `QWEN_CODE_MAX_INLINE_MEDIA_BYTES` 接受小数值，0.5 通过 `>0` 检查后 `Math.floor(0.5) = 0`，导致媒体截断失效。 | 1 条评论，已有 PR #5671 |

---

## 重要 PR 进展（10 个）

| PR | 状态 | 类型 | 核心变更 | 关联 Issue |
|----|------|------|----------|------------|
| [#5676](https://github.com/QwenLM/qwen-code/pull/5676) | OPEN | fix | 保持 settings v5 迁移的幂等性，修复 v0.19.0 发布失败（集成测试中因重复迁移报错） | #5663 |
| [#5662](https://github.com/QwenLM/qwen-code/pull/5662) | CLOSED | ci | 将 autofix 快速通道从 `status/ready-for-agent` 改为专用 `autofix/ready` 标签，防止 LLM 控制标签触发自动修复 | #5634 |
| [#5660](https://github.com/QwenLM/qwen-code/pull/5660) | OPEN | fix | `web_fetch` 在现有 Accept 头末尾添加低优先级 `*/*;q=0.1`，允许降级获取 JSON 响应 | #5611 |
| [#5667](https://github.com/QwenLM/qwen-code/pull/5667) | OPEN | fix | 要求 `QWEN_CODE_STOP_HOOK_BLOCK_CAP` 必须为正整数，小数回退默认值 | #5664 |
| [#5671](https://github.com/QwenLM/qwen-code/pull/5671) | OPEN | fix | 要求 `QWEN_CODE_MAX_INLINE_MEDIA_BYTES` 为正整数，小数回退默认值 | #5669 |
| [#5674](https://github.com/QwenLM/qwen-code/pull/5674) | OPEN | fix | 要求 `QWEN_CODE_MERMAID_RENDER_TIMEOUT_MS` 为正整数，小数回退默认值 | #5672 |
| [#5668](https://github.com/QwenLM/qwen-code/pull/5668) | OPEN | feat | 在加载指示器中实时显示模型的思考内容（`ThoughtSummary` 主题），替代随机励志短语 | #5656（部分） |
| [#5661](https://github.com/QwenLM/qwen-code/pull/5661) | OPEN | feat | 统一工具输出模式：已完成工具展示语义摘要（如“Read 3 files, edited 2 files”）而非原始结果，移除紧凑/普通双模式 | - |
| [#5632](https://github.com/QwenLM/qwen-code/pull/5632) | OPEN | feat | 为模型配置增加 `fastOnly` / `voiceOnly` 标志，可将模型从主模型列表中隐藏，仅在 fast/voice 选择器中显示 | - |
| [#4242](https://github.com/QwenLM/qwen-code/pull/4242) | OPEN | fix | 修复会话压缩后对话回退（rewind）目标映射错误，确保压缩后的 ACP 模型可正确回溯历史 | #4046 |

---

## 功能需求趋势

从近 24 小时 Issue 和 PR 中可以观察到社区最关注的方向：

1. **Provider 与模型管理**  
   - 解耦 Provider 身份与 SDK 协议（#5090 已关闭，但体现长期方向）  
   - 自定义 Provider 用户在 UI 中更方便地添加新模型（#4814）  
   - 通过 `fastOnly`/`voiceOnly` 标记隐藏不必要模型（#5632）

2. **安全性加固**  
   - 防止 LLM 控制修复流程标签（#5634）  
   - 环境变量值验证（端口、路径注入）  
   - 文件操作中暴露敏感内容的保护（#5550）

3. **工具系统体验打磨**  
   - 工具调用摘要移至加载指示器，减少对话噪音（#5656）  
   - 统一工具输出格式（#5661）  
   - 支持 JSON API 获取（#5611）  
   - 防止重复调用导致死循环（#5641）

4. **CI/CD 与构建效率**  
   - 复用 native 预构建制品（#5673）  
   - AI 辅助 PR 集成测试遗漏问题（#5665）  
   - 发布流程稳定性（多版本发布失败）

5. **配置与环境变量容错性**  
   - 多个环境变量接受小数导致意外行为（#5664、#5669、#5672、#5674）  
   - 请求社区对配置值做严格类型检查

---

## 开发者关注点

- **环境变量类型验证缺乏**：至少 4 个 Issue 报告了小数被误接受为整数配置的问题（超时、截断限制等），开发者希望框架对所有数值型配置做严格整数/正整数校验，避免隐形 bug。
- **工具调用稳定性**：重复提交已完成结果、Web 抓取无法处理 JSON 等问题直接影响日常开发流程，修复优先级高。
- **CI/CD 信任模型**：自动修复流程依赖 LLM 可以影响的标签，引发安全担忧。社区预期引入更严格的标签体系（如专用 `autofix/ready`）并限制其由人类 workflow 而非 issue 内容生成。
- **发布流程脆弱**：连续三次正式发布均失败（集成测试、publish 环节），暴露了 CI 中迁移脚本幂等性差、native 预构建冗余等工程债务，社区期望更彻底的 CI 重构。
- **AI 辅助开发协作缺漏**：AI 生成的 PR 常忘记更新集成测试，建议在 CI 中增加自动化检测或 prompt 改进。

---

*数据来源：GitHub [QwenLM/qwen-code](https://github.com/QwenLM/qwen-code)，截止 2026-06-23 09:00 UTC。日报由 AI 自动生成，如有疏漏欢迎指正。*

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

好的，作为专注于 AI 开发工具的技术分析师，我已根据您提供的 GitHub 数据，为您生成了 2026-06-23 的 CodeWhale（原 DeepSeek TUI）社区动态日报。

---

# 🚀 CodeWhale 社区日报 | 2026-06-23

## 📰 今日速览

1.  **品牌重塑加速落地**：`v0.8.64` 正式发布，项目已彻底从 `deepseek-tui` 迁移至 `CodeWhale`。社区安全加固（#3368）与 CI/CD 修复（#3374）成为本周发布的首要任务。
2.  **安全与可靠性成社区共识**：本周大量 Issue 围绕安全扫描修复、沙箱限制优化及自动化检查展开，表明开发者对 Agent 工具的稳定性和安全性要求日益严苛。
3.  **AI Agent 自主性争议持续**：用户反馈 Agent 过度自主、偏离用户意图的问题（#3275）依然活跃，社区正致力于通过“审阅门禁”（#3144）和“编辑护栏”（#3364）等技术手段平衡 Agent 效率与可控性。

---

## 🏷️ 版本发布

### v0.8.64: CodeWhale
-   **链接**: [查看发布页](https://github.com/Hmbown/DeepSeek-TUI/releases/tag/v0.8.64)
-   **核心摘要**:
    -   **品牌迁移**：项目正式更名为 **CodeWhale**，旧的 `deepseek-tui` 包已废弃。所有文档和资产均已完成更新。
    -   **安装**：[获取最新安装指南](https://github.com/Hmbown/CodeWhale/releases/tag/v0.8.64)

---

## 🔥 社区热点 Issues

### 1. #3368: 安全加固与代码扫描修复
-   “**重要性：极高**”，这是 v0.8.64 发布的安全门禁追踪器。
-   **链接**: [Hmbown/CodeWhale Issue #3368](https://github.com/Hmbown/CodeWhale/issues/3368)
-   **摘要**: 该 Issue 作为安全加固工作的公共追踪器，涵盖了 CodeQL 扫描、安全公告修复和本地集成提交。社区反应热烈，已有 **29条评论**，是过去24小时内讨论最活跃的 Issue。

### 2. #3275: CodeWhale 过度自主，偏离用户意图
-   “**核心痛点**”，直接关系到 Agent 的核心使用体验。
-   **链接**: [Hmbown/CodeWhale Issue #3275](https://github.com/Hmbown/CodeWhale/issues/3275)
-   **摘要**: 用户 `yekern` 报告了一个严重回归问题（来自 #3061），批评 CodeWhale 经常陷入“自问自答”的循环，自动执行超出用户请求范围的大量工作，让用户感到失控。**11条评论**显示其他用户对此深有共鸣。

### 3. #3144: 增加自然语言自动审查策略
-   “**解决 Agent 信任问题的钥匙**”。
-   **链接**: [Hmbown/CodeWhale Issue #3144](https://github.com/Hmbown/CodeWhale/issues/3144)
-   **摘要**: 受 Cursor 启发，提议在自动执行和手动确认之间建立一个“中间地带”：一个自然语言驱动的自动审查策略，能在代码提交前自动进行安全检查。这是解决 #3275 问题的关键方案之一，有 **12条评论**。

### 4. #3369: 恢复 Nightly 跨平台构建
-   “**工程可靠性基石**”。
-   **链接**: [Hmbown/CodeWhale Issue #3369](https://github.com/Hmbown/CodeWhale/issues/3369)
-   **摘要**: 维护者发现 Nightly 构建和自动标签工作流存在失败。直接影响 CI/CD 管道稳定性，是确保每次发布质量的基础。

### 5. #3222: 为 OpenAI 兼容 API 添加 `reasoning_style` 支持
-   “**新模型兼容性需求**”。
-   **链接**: [Hmbown/CodeWhale Issue #3222](https://github.com/Hmbown/CodeWhale/issues/3222)
-   **摘要**: 用户 `buko` 指出，CodeWhale 对 MiniMax M3、Qwen 等模型的思维链（Reasoning）内容解析存在问题。提议通过 `reasoning_style` 参数解决，表明社区对非 DeepSeek 模型的支持需求旺盛。

### 6. #3364: 添加“读后编辑”护栏和清晰的编辑失败提示
-   “**提升 Agent 可靠性的微观举措**”。
-   **链接**: [Hmbown/CodeWhale Issue #3364](https://github.com/Hmbown/CodeWhale/issues/3364)
-   **摘要**: 编辑错误是 Agent 不可靠感的主要来源。提议在执行编辑操作前强制要求模型读取最新文件内容，并让编辑失败的原因更具体、可操作。

### 7. #3363: 默认启用无缝自动压缩
-   “**解决长对话的核心痛点**”。
-   **链接**: [Hmbown/CodeWhale Issue #3363](https://github.com/Hmbown/CodeWhale/issues/3363)
-   **摘要**: 当模型接近上下文限制时，用户体验脆弱。提议通过携带前文摘要的方式，实现默认、自动、无缝的对话压缩，避免用户手动操作。

### 8. #3367: 支持用户自定义子代理角色
-   “**从工具到平台的进化**”。
-   **链接**: [Hmbown/CodeWhale Issue #3367](https://github.com/Hmbown/CodeWhale/issues/3367)
-   **摘要**: 提议在项目本地 `.codewhale/agents/` 目录下定义可复用的子代理角色（Persona），让开发者无需等待内置支持即可定制专属的代码审查、文档撰写等代理。

### 9. #3355: 沙箱阻止 Git 工作树（Worktree）写入操作
-   “**与 Git 高级用户工作流的冲突**”。
-   **链接**: [Hmbown/CodeWhale Issue #3355](https://github.com/Hmbown/CodeWhale/issues/3355)
-   **摘要**: 使用 Git Worktree 的用户发现，macOS 沙箱阻止了 `git add` 等操作，迫使他们切换到不安全的“信任模式”。此 Issue 已经通过 PR #3356 被修复（Closed）。

### 10. #3380: 让审批弹窗的按键提示更醒目
-   “**新用户体验的微小改进**”。
-   **链接**: [Hmbown/CodeWhale Issue #3380](https://github.com/Hmbown/CodeWhale/issues/3380)
-   **摘要**: 当前审批弹窗的按键提示（Enter/y/a/d）对比度低，新用户容易忽略。此 Issue 被标记为 **good first issue**，适合新手贡献者。

---

## 🤖 重要 PR 进展

### 1. #3373: v0.8.64 安全和发布集成
-   “**本周最核心的集成 PR**”。
-   **链接**: [Hmbown/CodeWhale PR #3373](https://github.com/Hmbown/CodeWhale/pull/3373)
-   **摘要**: 这是 v0.8.64 发布的集成分支，集成了所有安全加固、审核护栏、CI 修复和社区贡献。作为主合并请求，状态为 OPEN，需要密切关注其合并时间线。

### 2. #3356: 修复工作树沙箱写入
-   “**问题 #3355 的即时修复**”。
-   **链接**: [Hmbown/CodeWhale PR #3356](https://github.com/Hmbown/CodeWhale/pull/3356)
-   **摘要**: 通过检测链接的工作树 `.git` 指针文件，自动添加其指向的 `.git` 目录到沙箱白名单，从而允许 Git 写入操作。该 PR 已合并，问题得到解决。

### 3. #3374: 恢复 Nightly 跨平台构建
-   “**修复工程可靠性问题**”。
-   **链接**: [Hmbown/CodeWhale PR #3374](https://github.com/Hmbown/CodeWhale/pull/3374)
-   **摘要**: 针对 #3369 的修复，确保跨平台构建正常工作，并增加构建工件存在性检查以避免重复构建，提升 CI 稳定性。

### 4. #3350: 添加 `/model` 快捷命令和 CLI 模型设置
-   “**提升模型切换效率**”。
-   **链接**: [Hmbown/CodeWhale PR #3350](https://github.com/Hmbown/CodeWhale/pull/3350)
-   **摘要**: 新增 `pro` 和 `flash` 两个快捷别名，用户可以通过 `/model pro` 快速切换到 DeepSeek V4 Pro 模型，并通过 `codewhale model set` CLI 命令进行设置。

### 5. #3371: 降低侧边栏可见的最小终端宽度
-   “**提升小屏幕用户体验**”。
-   **链接**: [Hmbown/CodeWhale PR #3371](https://github.com/Hmbown/CodeWhale/pull/3371)
-   **摘要**: 修复了 #3328，将侧边栏显示的最小宽度从 100 列降低到更常用的值，使更多用户能在默认终端大小下看到侧边栏。

### 6. #3376: 添加开发服务器就绪检测工具
-   “**提升 Web 应用开发自动化能力**”。
-   **链接**: [Hmbown/CodeWhale PR #3376](https://github.com/Hmbown/CodeWhale/pull/3376)
-   **摘要**: 新增 `wait_for_dev_server` 工具，允许 Agent 在启动本地开发服务器后，轮询其 TCP/HTTP 端口以确认服务已就绪，避免 Web 测试中的竞态问题。

### 7. #3375: 抑制提供商等待时的闲置超时倒计时
-   “**减少 UI 噪音**”。
-   **链接**: [Hmbown/CodeWhale PR #3375](https://github.com/Hmbown/CodeWhale/pull/3375)
-   **摘要**: 修复 #3189，优化了数据流等待时的状态栏提示。当等待时间较短时（<60s），不再显示倒计时，减少不必要的视觉干扰；仅在接近超时阈值时显示详细时间。

### 8. #3377: 更新安全联系邮箱
-   “**安全合规的微小但重要操作**”。
-   **链接**: [Hmbown/CodeWhale PR #3377](https://github.com/Hmbown/CodeWhale/pull/3377)
-   **摘要**: 更新 `SECURITY.md` 中的安全联系邮箱，以匹配 CodeWhale 的全新品牌域名，确保漏洞报告能够正确送达。

### 9. #3345: 重构配置内联测试
-   “**提升代码可维护性**”。
-   **链接**: [Hmbown/CodeWhale PR #3345](https://github.com/Hmbown/CodeWhale/pull/3345)
-   **摘要**: 社区贡献者 `cyq1017` 将 `crates/config/src/lib.rs` 中的大型内联测试模块迁移到独立的 `tests.rs` 文件中，减少了生产代码的体积并降低了合并冲突的风险。

### 10. #3381: 内存标签功能
-   “**探讨中，可能为未来功能**”。
-   **链接**: [Hmbown/CodeWhale PR #3381](https://github.com/Hmbown/CodeWhale/pull/3381)
-   **摘要**: 一个关于“记忆标签”（Feat/memory tags）的新 PR 被创建。虽然描述尚不清晰，但结合社区对 Agent 长期记忆的需求，这可能是 CodeWhale 增强项目级上下文管理的一个重要信号。

---

## 📈 功能需求趋势

1.  **安全性、可靠性与沙箱**：超过一半的热点 Issue 与此相关，包括安全扫描（#3368）、沙箱限制（#3355）、审计轨道（#3144）。社区对 Agent 工具的信任度要求已达到生产级标准。
2.  **Agent 行为可控性**：如何避免 Agent “过度发挥”（#3275）是核心矛盾。社区共识是通过更强的 **护栏**（Guardrails），如预审策略（#3144）、“读前编辑”（Read-before-edit）等机制来约束 Agent。
3.  **Web 开发与自动化**：关于 Playwright 浏览器自动化（#3358）和开发服务器管理（#3360、#3376）的讨论表明，CodeWhale 正从纯 CLI 工具向全能型“编码代理”演进。
4.  **跨模型适配**：除了原生 DeepSeek 模型，社区强烈希望 CodeWhale 无缝支持 OpenAI、MiniMax、Qwen 等主流第三方模型，尤其对推理过程和工具调用的兼容性提出更高要求。
5.  **工程化与 CI/CD**：恢复 Nightly 构建（#3369）、分支卫生检查（#3348）等修复说明，随着项目复杂度增加，保障基础工程设施的稳定性成为开发团队的关注重点。

---

## 🔧 开发者关注点

1.  **Agent “过度自主” 是最普遍的负面体验** (#3275)：开发者抱怨 Agent 经常不按指令行事，进行不必要的“自问自答”，甚至主动修改无关文件。这要求必须提供更强力的控制开关和审核流程。
2.  **沙箱模式是双刃剑** (#3355)：沙箱提升了安全性，但与 Git 工作树等高级工作流冲突，导致用户被迫在“安全但不可用”和“可用但不安全”之间做选择。社区期待更智能、更精细的沙箱策略。
3.  **缺乏浏览器级测试能力**：`web_run` 工具只能抓取页面文本，无法执行复杂的 JavaScript、动态交互或截图，限制了其在前端迭代和 UI 自动化测试中的应用。
4.  **用户界面细节影响体验**：从侧边栏宽度（#3371）到审批按键提示（#3380），再到闲置倒计时提示（#3375），说明即使是资深用户也对 UI 的细节非常敏感，这些“噪音”会直接影响工作效率。
5.  **本地化配置的灵活性**：自定义子代理角色（#3367）的需求表明，开发者希望 Agent 能为不同场景（如代码审查、文档编写）自动切换不同的“专家人格”，而不是一个万能的通用模型。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*