# AI CLI 工具社区动态日报 2026-06-30

> 生成时间: 2026-06-30 10:45 UTC | 覆盖工具: 9 个

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

# AI CLI 工具横向对比分析报告（2026-06-30）

---

## 1. 生态全景

当前 AI CLI 工具生态呈现 **“三足鼎立、多元深耕”** 的格局。以 Anthropic、OpenAI、Google 为代表的第一梯队（Claude Code、OpenAI Codex、Gemini CLI）在核心代理能力、长会话管理、企业级安全方面持续加码；GitHub Copilot CLI 背靠 GitHub 生态，在插件系统与 IDE 协同上快速迭代；开源/半开源项目（OpenCode、Pi、Qwen Code、DeepSeek TUI、Kimi Code）则围绕模型集成多样性、细粒度控制、社区驱动创新展开差异化竞争。整体趋势是：**从“能跑模型”转向“可靠执行、智能规划、安全合规”**，各工具均面临稳定性、用户体验与成本控制三大核心挑战。

---

## 2. 各工具活跃度对比（2026-06-30）

| 工具 | 今日 Issues 更新数 | 今日 PR 更新数 | 今日 Release | 社区讨论热度（评论总数） | 核心维护者活跃度 |
|------|-------------------|----------------|--------------|------------------------|------------------|
| Claude Code | 10 | 10 | 1（v2.1.196） | 高（5个Issue评论>20条） | 极高（官方持续迭代） |
| OpenAI Codex | 10 | 10 | 1（alpha.31） | 高（#30224 61评论） | 高（官方+外部PR） |
| Gemini CLI | 10 | 10 | 1（nightly） | 中（多数评论<10） | 高（官方修复密集） |
| GitHub Copilot CLI | 10（更新5条新issue） | 1 | 1（v1.0.66-2） | 中（#1665 17👍） | 中（仅1 PR） |
| Kimi Code | 0 | 2 | 0 | 低 | 低（仅2 PR合入） |
| OpenCode | 10+ | 10+ | 0 | 高（#7602 90👍, #8463 89👍） | 极高（Robin1987China批量合入） |
| Pi | 10 | 10 | 0 | 中（#5825 42评论） | 高（修复密集） |
| Qwen Code | 5 | 10（精选） | 1（nightly） | 中（#6049 中等讨论） | 高（频道/记忆等多方向） |
| DeepSeek TUI | 10 | 10 | 0（冲刺v0.8.66） | 中（#1177 24评论） | 极高（7个PR针对TUI冻结） |

**说明**：Issues/PR 数量基于日报列出或社区仓库当日更新情况估算。热度综合评论数和点赞数。

---

## 3. 共同关注的功能方向

### 3.1 多模型与模型路由
- **Claude Code**：组织默认模型、角色默认模型支持
- **OpenAI Codex**：GPT-5.5 容量不足导致模型无法切换（底层需求）
- **Gemini CLI**：MCP elicitation 扩展模型能力
- **GitHub Copilot CLI**：`explore` 工具硬编码 `gpt-5.4-mini`，用户期望自定义模型
- **OpenCode**：Model Fallback (#7602, 90👍)、Copilot 自动路由 (#20235, #25239)
- **Pi**：小米 MiMo 定价错误、Scaleway 提供商请求 (#6165)
- **Qwen Code**：`/model <id>` 一次性覆盖命令 (#6022)

**核心诉求**：用户希望工具能灵活选择/回退模型，避免单点故障，同时降低推理成本。

### 3.2 会话持久化与中断恢复
- **Claude Code**：中断后恢复执行 (#16607)
- **OpenAI Codex**：用户消息队列持久化、跨重启恢复（PR #28307 等）
- **Gemini CLI**：自动记忆系统优化（脱敏、停止无效重试）
- **Pi**：Session 文件夹冲突 (#4877)
- **Qwen Code**：会话存档 (#6057)、折叠预览 (#5759)、工作区无会话记忆 (#5884)
- **DeepSeek TUI**：会话永久损坏 (#3821，新 Bug)

**核心诉求**：大型任务（尤其是多步骤 Agent 流程）需要可靠的中间状态保存与恢复能力，减少重复劳动。

### 3.3 安全性 & 审计
- **Claude Code**：安全指导插件符号链接修复、`/btw` 安全上下文
- **Gemini CLI**：文件写入沙箱安全加固、思维泄露修复、代理劝阻破坏性行为
- **GitHub Copilot CLI**：Hook 拒绝策略失效修复 (#3874)
- **OpenCode**：YOLO 模式跳过权限 (#8463)、OAuth token 刷新
- **Qwen Code**：`deniedMcpServers` 白名单补充、TLS 支持、子代理结果净化
- **DeepSeek TUI**：MCP OAuth 认证 UX 问题 (#3819)

**核心诉求**：AI 代理拥有执行权限后，必须有细粒度的控制（禁止/允许）、中断机制和透明日志，防止信息泄露或误操作。

### 3.4 终端 UI/UX 精细打磨
- **Claude Code**：自动复制禁用 (#60755)、状态栏显示
- **OpenAI Codex**：复制对话为 Markdown (#2880, 71👍)
- **Gemini CLI**：多行编辑省略号、`camelToSpace` 修复
- **GitHub Copilot CLI**：alt-screen 可控性、触控板滚动异常、鼠标字符污染
- **Kimi Code**：高亮用户输入 + 分隔线 PR
- **Qwen Code**：emoji 转 Unicode、代码块复制去行号、标签化设置
- **DeepSeek TUI**：审批控件保持可见、Hotbar 默认隐藏
- **Pi**：禁用助手消息填充、流式滚动问题

**核心诉求**：终端已非“简陋界面”，用户期待类似 VSCode 的交互细节——可控的布局、清晰的视觉层次、标准化的快捷键与剪贴板行为。

### 3.5 内存/性能优化
- **Claude Code**：深度压缩 (#72461)
- **OpenAI Codex**：瞬时事件跳过持久化、重试机制
- **GitHub Copilot CLI**：`tgrep` OOM-kill 宿主机 (#3976)
- **Qwen Code**：避免全量历史克隆 (#6018)
- **DeepSeek TUI**：子 Agent 高扇出 TUI 冻结（锁竞争、非阻塞发送）
- **Pi**：流式传输错误处理、重试机制

**核心诉求**：大代码库、多 Agent、长上下文场景下，资源消耗（内存、Token、CPU）必须可控，否则生产环境不可用。

---

## 4. 差异化定位分析

| 工具 | 定位 | 核心差异化 | 目标用户 | 技术路线 |
|------|------|------------|----------|----------|
| **Claude Code** | 通用高智商 Agent | 深度推理（Opus）、长上下文、企业级安全（组织默认模型、安全插件） | 专业开发者、团队 | 闭源，优先安全与可靠性 |
| **OpenAI Codex** | 全栈开发助手（桌面+CLI） | GPT-5.x 专用、Chronicle 后台录制（计算机使用）、沙箱隔离 | 全栈开发者、AI 重度用户 | 闭源，强推桌面应用与平台化 |
| **Gemini CLI** | 智能代码代理 + 评估驱动 | 子代理架构、AST 感知、组件级评估体系、MCP 扩展 | Google 云生态开发者、研究类 | 开源（但核心模型闭源），强调代理规划 |
| **GitHub Copilot CLI** | GitHub 生态嵌入式终端 | 插件系统、项目级 scope、alt-screen、与 Copilot 模型路由集成 | GitHub 用户、企业团队 | 闭源，深度绑定 GitHub 工作流 |
| **OpenCode** | 开源通用 Code Agent | 模型 Fallback、YOLO 模式、MCP OAuth、Copilot 路由混用 | 开源用户、需要自定义模型者 | 开源，社区驱动，快速迭代 |
| **Pi** | 轻量级交互式 Agent | 多提供商、HITL 编程、Bedrock 稳定性、Redo 编辑 | 独立开发者、多模型实验者 | 开源，强调跨平台与 SDK 扩展 |
| **Qwen Code** | 阿里系模型驱动的 AI CLI | daemon 架构、频道/QQ 集成、记忆系统、auto-mode 分类 | 阿里云/Qwen 用户、亚洲市场 | 开源，大模型原生 + 消息渠道扩展 |
| **DeepSeek TUI** | 高性能 TUI Agent | 子 Agent 并行、缓存优化（当前弱项）、MCP 通配符 | 追求性价比和终端原生体验的开发者 | 开源，社区贡献活跃，性能敏感型 |
| **Kimi Code** | 极简交互式 Shell | 轻量、`--prompt-interactive`、视觉优化 | 轻量用户、初学者 | 开源，功能简单，社区不活跃 |

---

## 5. 社区热度与成熟度

### 高热度 + 高成熟度
- **Claude Code**：老牌主力，社区反馈专业，Bug 追踪系统完善，官方迭代稳定。适合生产环境。
- **OpenAI Codex**：虽然 Bug 多（容量、Windows 沙箱），但官方人力充足，PR 合并快，功能丰富度最高。适合胆大的早期采用者。

### 快速迭代 + 社区响应强
- **OpenCode**：近两日 Issue 获 90+ 赞，PR 由社区明星（Robin1987China）批量合入，功能与底层修复并进。是目前开源生态中最活跃的 AGI CLI。
- **DeepSeek TUI**：v0.8.66 冲刺阶段，一天 7 个 PR 修复 TUI 冻结，社区反馈（缓存、Token）强烈。适合性能敏感和社区参与意愿高的用户。
- **Qwen Code**：频道记忆、会话存档等创新多，PR 活跃度高，但 Issue 更新相对较少（5条），社区热度略低于 OpenCode。

### 中等活跃度
- **Gemini CLI**：功能丰富，但 Issues 更新数虽多，评论热度不高（多数 <10），说明用户群体较小或更偏向沉默。
- **Pi**：拥有稳定的社区（#5825 42评论），PR 集中在 Bedrock 和错误处理，但缺乏上游大力推动。
- **GitHub Copilot CLI**：用户基数大，但 PR 更新仅 1 条，显示官方投入有所放缓。Issue 数量不少，但部分核心 Bug（如 tgrep OOM）刚提交。

### 低活跃度 / 早期阶段
- **Kimi Code**：24小时内无 Issue 更新，仅 2 条 PR，社区不活跃。功能有限。
- **其他未提及工具**：如 Llama CLI 等不在本次分析内。

---

## 6. 值得关注的趋势信号

### 6.1 缓存与成本优化成为刚需
DeepSeek TUI 用户对比发现缓存命中率远低于竞品（95%+ vs 当前70-80%），直接反映在 Token 消耗和 API 账单上。**所有工具都必须配套透明化的缓存统计和可配置压缩策略**，否则用户将流向成本更低的方案。

### 6.2 AI “代理” 的可靠性标准急剧提升
过去“能跑就行”，现在用户要求 Agent 卡死后可恢复、子 Agent 状态不丢失、模型推理不误报（#22323 Gemini）、工具调用不输出原始 XML（#64108 Claude）。**“可中断、可恢复、可审计”将成为生产级工具的必要条件。**

### 6.3 MCP 生态从“可选项”变为“核心能力”
多个工具（OpenCode、Qwen Code、Gemini CLI、DeepSeek TUI）都在完善 MCP 集成，包括通配符工具过滤、环境变量注入、OAuth 认证优化。**MCP 协议正在成为 AI CLI 的“插件标准”**，未来可能取代部分专有扩展机制。

### 6.4 跨平台兼容性问题持续暴露
Windows 沙箱崩溃（OpenAI Codex #29072）、macOS SSL 证书失效（Claude Code #71663）、Linux Wayland 失败（Gemini #21983）、Windows 控制台弹出（DeepSeek #3823）…… **Linux 用户相对省心，macOS 和 Windows 仍需针对性投入**。对于面向全球市场的产品，跨平台 CI 必须加强。

### 6.5 隐私与安全从“警示”变成“默认要求”
自动记忆脱敏（Gemini #26525）、MCP OAuth 刷新（多个）、符号链接注入（Claude Code #68689）——**用户不再接受“事后检查”，而是要求 “Secure by Design”**。企业级用户尤其关注。

### 6.6 模型行为异常需透明化
OpenAI Codex GPT-5.5 推理 token 固定边界（#30364）、Claude Code 工具调用输出原始文本（#64108）、Gemini 子代理误报成功（#22323）——**当 AI 输出不符合预期时，用户需要可解释的日志和调试能力**。未来“思考链透明化”可能成为标配。

---

*报告基于 2026-06-30 各项目 GitHub 公开数据整理，部分数据经过估算，但仍具有代表性。*

---

## 各工具详细报告

<details>
<summary><strong>Claude Code</strong> — <a href="https://github.com/anthropics/claude-code">anthropics/claude-code</a></summary>

## Claude Code Skills 社区热点

> 数据来源: [anthropics/skills](https://github.com/anthropics/skills)

好的，作为专注于 Claude Code 生态的技术分析师，以下是根据截至2026年6月30日的 GitHub 数据生成的社区热点报告。

---

## Claude Code Skills 社区热点报告 (数据截止 2026-06-30)

### 1. 热门 Skills 排行 (基于 PR 评论与关注度)

以下列出了社区讨论热度最高、关注度最集中的 7 个 Skills 动态：

1.  **修复：`run_eval.py` 评分系统全面失效 (PR #1298)**
    *   **功能**: 修复了核心评估脚本 `run_eval.py` 在多个场景下始终报告0%召回率的严重Bug。
    *   **讨论热点**: 该问题在 Issue #556 中有“10+ 独立复现”的报告，社区广泛关注其导致整个 skill 优化循环“在噪音上优化”的严重后果。评论焦点集中在修复方案的全面性（Windows支持、触发检测等）。
    *   **状态**: **OPEN** (评论最多，至关重要)
    *   **链接**: [PR #1298](https://github.com/anthropics/skills/pull/1298)

2.  **新增：`document-typography` 文档排版质量技能 (PR #514)**
    *   **功能**: 自动检查和修正AI生成文档中的典型排版问题，如孤词、寡行和编号错位。
    *   **讨论热点**: 社区普遍认同这是AI生成内容中“高频但用户不常主动提出”的痛点。讨论焦点在于该技能的普适性和对文档美观度的实际提升效果。
    *   **状态**: **OPEN** (高价值实用技能)
    *   **链接**: [PR #514](https://github.com/anthropics/skills/pull/514)

3.  **新增：`odt` 开放文档格式支持 (PR #486)**
    *   **功能**: 支持创建、填充、读取和转换 `.odt`、`.ods` 等 OpenDocument 格式文件，与 LibreOffice 生态无缝对接。
    *   **讨论热点**: 社区对扩展 Claude 处理非微软 Office 格式的能力表示欢迎，讨论焦点集中在处理复杂表格、样式和模板填充的可靠性上。
    *   **状态**: **OPEN** (生态扩展方向)
    *   **链接**: [PR #486](https://github.com/anthropics/skills/pull/486)

4.  **改进：`frontend-design` 技能清晰度与可操作性 (PR #210)**
    *   **功能**: 重写前端设计技能，确保指令清晰、可操作，且 Claude 能在单次对话中有效执行。
    *   **讨论热点**: 社区讨论反映了对“如何编写高质量 Skill”的深度思考，包括指令具体性、行为引导和内部一致性。这是一个元层面的讨论，影响所有类型的 Skills。
    *   **状态**: **OPEN** (技能最佳实践讨论)
    *   **链接**: [PR #210](https://github.com/anthropics/skills/pull/210)

5.  **新增：元技能 `skill-quality-analyzer` & `skill-security-analyzer` (PR #83)**
    *   **功能**: 引入两个分析工具 Skill，分别用于评估其他 Skill 的质量（结构、文档、例子）和安全性。
    *   **讨论热点**: 社区高度关注该想法，认为这是建立 Skill 质量标准和安全审查体系的关键一步。讨论焦点在于分析维度的准确性和如何规避误判。
    *   **状态**: **OPEN** (杀手级辅助工具)
    *   **链接**: [PR #83](https://github.com/anthropics/skills/pull/83)

6.  **修复：`skill-creator` 描述中的 YAML 字符解析问题 (PR #361， #539)**
    *   **功能**: 检测并警告 Skill 描述 YAML 头信息中未加引号的特殊字符（如 `:`、`#`），防止静默解析失败。
    *   **讨论热点**: 这是一个开发者的“潜在陷阱”，社区反应积极，认为这能大幅减少因配置错误导致的 Skill 失效问题，是提升开发体验的重要细节。
    *   **状态**: **OPEN** (提升开发者体验)
    *   **链接**: [PR #361](https://github.com/anthropics/skills/pull/361) ， [PR #539](https://github.com/anthropics/skills/pull/539)

7.  **新增：`self-audit` 交付前推理质量关 (PR #1367)**
    *   **功能**: 一个通用 Skill，在 Claude 交付最终输出前，强制其对完整性、一致性、基础性和安全性进行四维自检。
    *   **讨论热点**: 这是一个非常新颖且富有争议的技能。讨论焦点集中在其是否过于通用、能否有效防止“幻觉”以及是否会引入额外的对话开销。
    *   **状态**: **OPEN** (新锐且争议性高)
    *   **链接**: [PR #1367](https://github.com/anthropics/skills/pull/1367)

### 2. 社区需求趋势 (从 Issues 提炼)

根据 Issue 的讨论热度，社区最期待的新 Skill 方向与核心关注点如下：

1.  **安全与信任 (Issue #492)**: 社区对 `anthropic/` 命名空间下混入社区贡献的技能表达了强烈的安全担忧，呼吁建立更明确的信任边界、来源审核和权限管理机制。这是当前最尖锐的社区诉求。
2.  **团队协作与共享 (Issue #228)**: 用户迫切需要在组织内直接共享 Skills，而非通过下载和手动上传的方式。一个集中的 Skill 共享库或直接分享链接成为强烈需求。
3.  **特定的文件格式与业务系统集成 (Issue #1175)**: 处理 SharePoint Online (SPO) 等企业级云盘文档的需求出现，并引发了关于权限、安全和上下文窗口控制的深度讨论。
4.  **技能创建工具链的稳定性与可靠性 (Issue #556, #1061, #1169)**: 大量 Issues 报告了 `skill-creator` 脚本在 Windows 平台、编码处理和评估循环中的崩溃和错误，表明社区对稳定、可用的开发工具有极高的期望。
5.  **平台兼容性 (Issue #1061)**: Windows 用户面临显著的兼容性问题，包括子进程调用、编码和管道读取失败，是社区的一个主要痛点。
6.  **新兴生态整合 (Issue #16)**: 将 Skills 暴露为 MCPs 的早期提议，显示了社区对构建更开放、可组合的 AI 生态系统的远见和兴趣。

### 3. 高潜力待合并 Skills (评论活跃但未合并)

以下 PR 因功能重要、讨论充分或修复核心问题，具有近期合并的高潜力：

*   **`run_eval.py` 修复系列 (PR #1298, #1099, #1050, #1323)**: 这些 PR 共同指向修复核心评估工具的稳定性与正确性。没有这个工具，所有 Skill 的优化循环都无法正常工作。这是最高优先级的合并对象。
*   **ODT 技能 (PR #486)**: 一个新生态点的成熟技能，社区需求明确。随着后续讨论的收敛，合并可能性很高。
*   **YAML 特殊字符检测 (PR #361)**: 一个简单有效的代码质量改进，对开发者体验提升明显，预计会较快合并。

### 4. Skills 生态洞察

**当前社区在 Skills 层面最集中的诉求是：稳定、可信、可协作 —— 即在修复核心开发工具链和评估体系（盈利/报销）Bug 的基础上，建立官方的安全信任机制和团队共享渠道，以支撑 Skills 生态的健康、规模化发展。** 单纯的新技能创意已开始被对工具链可靠性和生态治理的需求所覆盖。

---

好的，这是为你生成的 2026-06-30 的 Claude Code 社区动态日报。

---

# Claude Code 社区动态日报 | 2026-06-30

## 今日速览

今日社区动态活跃，**多账户切换**和**IDE功能对齐** (VS Code 扩展) 是讨论热度最高的两个话题。同时，Claude Code 发布了 v2.1.196 小版本，新增了**组织默认模型**支持并优化了会话识别体验。此外，多个关于**上下文压缩**和**SSL证书**的 Bug 被提出，值得关注。

## 版本发布

### v2.1.196 发布
- **核心更新**:
    1.  **组织默认模型**: 管理员可在组织控制台设置默认模型，用户在未手动选择时会看到“Org default”或“Role default”标识，简化了团队内的模型管理。
    2.  **会话命名优化**: 新的会话在启动时会自动生成易读的默认名称，方便在会话列表中快速识别和切换。

## 社区热点 Issues

1.  **[#36151] 多账户切换功能**
    - **重要性**: 社区呼声最高的功能请求 (👍401条) ，讨论超过百条。用户希望在 Claude 移动端（可能也隐含在工具中）实现无需共享邮箱的账户切换，对于有个人和工作账户的用户是刚需。
    - **社区反应**: 用户普遍认为这是当前工作流中的最大痛点，讨论集中在对账户隔离和安全性的要求上。
    - 链接: [Issue #36151](https://github.com/anthropics/claude-code/issues/36151)

2.  **[#37323] 在 VS Code 扩展中支持 `/btw` 命令**
    - **重要性**: 用户期望 CLI 和 IDE 扩展的功能对齐。`/btw` 命令可以快速处理侧线问题，提高与 AI 交互的效率。
    - **社区反应**: 开发者普遍认为这是一个缺失的核心功能，表示在 IDE 中工作流断裂，不得不切换回终端使用该命令。
    - 链接: [Issue #37323](https://github.com/anthropics/claude-code/issues/37323)

3.  **[#64108] 工具调用异常：模型执行失败，输出原始文本**
    - **重要性**: 影响核心功能的严重 Bug。模型在一个长会话中（特别是使用 Opus 模型时），将工具调用（如读/写文件）的 XML 指令作为普通文本输出，导致指令执行失败。
    - **社区反应**: 该问题影响用户信任，尤其是在处理大型、复杂项目时，看起来像个难以追踪的偶发故障。
    - 链接: [Issue #64108](https://github.com/anthropics/claude-code/issues/64108)

4.  **[#30958] 推理摘要 (thinking summaries) 在转录中为空**
    - **重要性**: 这是一个已持续数月的 Bug。自 v2.1.69 更新后，模型的“思考过程”摘要变得空白，影响用户理解和调试模型逻辑。
    - **社区反应**: 用户认为这是`回归`问题，指出该行为改变未经文档说明，对依赖这些内容进行“思考链”分析的开发者造成困扰。
    - 链接: [Issue #30958](https://github.com/anthropics/claude-code/issues/30958)

5.  **[#16607] 允许 Agent 中断后恢复执行**
    - **重要性**: 一个长期存在的功能请求。当开发者手动中断（Esc键）一个长时间运行的任务后，Agent 无法从中断点接续，导致重复劳动。
    - **社区反应**: 用户反馈在调试或执行耗时任务时这一功能缺失尤为明显，希望 Agent 能够记住已完成的部分并从“下一个步骤”开始。
    - 链接: [Issue #16607](https://github.com/anthropics/claude-code/issues/16607)

6.  **[#60755] 禁用 TUI 中的自动复制功能**
    - **重要性**: 社区对使用习惯的差异化诉求。当前在 Agent 管理视图中选中文本即自动复制，不符合所有用户的交互习惯。
    - **社区反应**: 讨论集中在与原生 macOS 终端行为的一致性上，不少用户希望在设置中添加一个开关来控制。
    - 链接: [Issue #60755](https://github.com/anthropics/claude-code/issues/60755)

7.  **[#63048] `/code-review` 代码审查范围错误**
    - **重要性**: 影响代码审查功能的可靠性。当分支没有跟踪远程分支时，该命令会回退到 `main...HEAD` 进行差异比较，这会将其他开发者的提交也纳入审查范围，产生错误结果。
    - **社区反应**: 开发者期待该命令能更智能地判断代码范围，例如使用更精确的 git 命令来定位当前功能分支的基础提交。
    - 链接: [Issue #63048](https://github.com/anthropics/claude-code/issues/63048)

8.  **[#62644] “购买积分”按钮被禁用 (Free 用户)**
    - **重要性**: 一个严重的账户和计费 Bug，导致免费用户无法正常付费升级。界面显示 $500 用量限制，但在尝试购买时遇到 HTTP 429 错误。
    - **社区反应**: 受影响的用户表示无法开始使用或继续使用付费功能，用户反馈中透露出较强的挫败感。
    - 链接: [Issue #62644](https://github.com/anthropics/claude-code/issues/62644)

9.  **[#72461] 为超长编码会话增加“深度压缩”功能**
    - **重要性**: 针对超长会话性能问题的解决方案。当前 Claude Code 的上下文压缩机制（上下文压缩）在处理深度嵌套的对话历史时效果不佳，导致性能下降和模型“忘记”早期内容。
    - **社区反应**: 该功能请求刚一发布，就获得了专业用户的关注，他们是当前上下文压缩机制的主要受困者。
    - 链接: [Issue #72461](https://github.com/anthropics/claude-code/issues/72461)

10. **[#71663] SSL 证书在 macOS 上失效 (v2.1.190+)**
    - **重要性**: 严重影响网络连接的 Bug。从 v2.1.190 版本开始，Claude Code 在 macOS 上发起的网络请求因 SSL 证书验证问题而失败。
    - **社区反应**: 受影响的用户很多，普遍要求紧急修复，因为它完全阻断了需要网络连接的模型和功能。
    - 链接: [Issue #71663](https://github.com/anthropics/claude-code/issues/71663)

## 重要 PR 进展

1.  **[#68699] 修复 hookify 插件在 Windows 上的路径问题**
    - **功能**: 修复了 Windows 系统下 `CLAUDE_PLUGIN_ROOT` 环境变量中的反斜杠导致 Hook 命令执行失败的问题，并添加了 Python 包装器来解决特定 Python 运行环境的问题。
    - 链接: [PR #68699](https://github.com/anthropics/claude-code/pull/68699)

2.  **[#68701] 修复 Windows 平台上的 CRLF 问题**
    - **功能**: 解决了 Python 版本检查时，Windows 换行符 `\r\n` 导致 Bash 字符串比较失败的问题。
    - 链接: [PR #68701](https://github.com/anthropics/claude-code/pull/68701)

3.  **[#68689] 修复安全指导插件的符号链接利用漏洞**
    - **功能**: 重要安全修复。修复了恶意仓库可通过符号链接文件（如 `claude-security-guidance.md`）指向任意本地文件，从而泄露敏感信息的风险。
    - 链接: [PR #68689](https://github.com/anthropics/claude-code/pull/68689)

4.  **[#68702] 修复 ralph-wiggum 插件在 macOS 上的 Bash 3.x 兼容性**
    - **功能**: 修复了在 macOS 系统（未安装新 Bash）上，由于 `set -u` 导致的对空数组扩展报错问题。
    - 链接: [PR #68702](https://github.com/anthropics/claude-code/pull/68702)

5.  **[#68693] 修复重复 Issue 标签处理脚本**
    - **功能**: 修复了关闭重复 Issue 时，`labels` 参数的设置会意外清空原有标签的问题。现在改为只添加 `duplicate` 标签。
    - 链接: [PR #68693](https://github.com/anthropics/claude-code/pull/68693)

6.  **[#68686] 修复 hookify 插件配置解析器 Bug**
    - **功能**: 解决了 Python 中的变量名作用域冲突（`field` 变量与 `dataclasses.field` 冲突）和解析内联字典时遇到的逗号处理问题。
    - 链接: [PR #68686](https://github.com/anthropics/claude-code/pull/68686)

7.  **[#68694] 修复安全指导插件在 Windows 上的路径问题**
    - **功能**: 与 #68699 类似，修复了该插件在 Windows 环境中因路径分隔符导致 hook 命令失败的问题。
    - 链接: [PR #68694](https://github.com/anthropics/claude-code/pull/68694)

8.  **[#68690] 修正 ralph-wiggum 插件文档错误**
    - **功能**: 修复了帮助文档中关于状态文件路径的错误说明，消除了文档与实际实现不一致的问题。
    - 链接: [PR #68690](https://github.com/anthropics/claude-code/pull/68690)

9.  **[#68707] 新增 `/bug` 命令 (bug-reporter 插件)**
    - **功能**: 一个非常实用的新插件。允许用户直接在 Claude Code 终端中通过 `/bug` 命令一键提交 GitHub Issue，简化了反馈流程。
    - 链接: [PR #68707](https://github.com/anthropics/claude-code/pull/68707)

10. **[#72451] 修复防火墙初始化脚本**
    - **功能**: 从启动脚本的允许列表中移除一个已失效的域名 (`statsig.anthropic.com`)，解决了 DevContainer 启动时因 DNS 解析失败而报错中断的问题。
    - 链接: [PR #72451](https://github.com/anthropics/claude-code/pull/72451)

## 功能需求趋势

- **多账户与跨项目支持**: 社区非常渴望能够无缝切换个人与工作账户（#36151），以及在 `--resume` 搜索中能跨项目检索历史会话，这表明开发者工作流日益复杂，单一项目模式已无法满足需求。
- **IDE 集成与功能对齐**: VS Code 扩展与 CLI 之间的功能差异（如 `/btw` 命令缺失）是明显的痛点。社区希望获得与终端体验完全一致的无缝集成体验。
- **会话管理与可靠性**: “中断后恢复”（#16607）和“深度上下文压缩”（#72461）反映了开发者对 Agent 的稳定性和持久性提出了更高要求，希望能处理更复杂的、长时间的工作流。
- **安全与合规**: 对插件安全性的关注度提升（通过 #68689 的修复可见），开发者越来越在意 AI 开发工具是否会引入新的安全风险。
- **UI/UX 微调**: 对自动复制、屏幕重绘（Ctrl+L）、状态栏显示等交互细节的反馈，表明社区已进入精细化打磨阶段，希望工具能更符合直觉和现有系统习惯。

## 开发者关注点

- **核心功能的稳定性和可靠性**: 工具调用异常（#64108）、SSL 证书失效（#71663）这类 Bug 会直接打断工作流，是开发者的首要痛点，也是紧急度最高的修复方向。
- **长会话与大型项目的性能**: 上下文模型“遗忘”、“思考过程”（Thinking）不透明（#30958）以及无法从中断点恢复等问题，严重影响了开发者在大型项目中使用 Claude Code 的信心和效率。
- **平台兼容性**: Windows 和 macOS（尤其是旧版 Bash）的兼容性问题是持续的关注点，很多开发者希望体验能与 Linux 保持一致。
- **代码审查等高级功能的准确性**: `/code-review` 设置错误的审查范围会输出无效信息，开发者希望这类高级功能足够智能和可靠，能真正辅助开发，而非制造噪音。
- **账户、网络与计费问题**: “购买按钮禁用”（#62644）和服务器限流（#71635）等故障影响了用户的基本使用和付费意愿，是社区信任度建设的关键环节。

</details>

<details>
<summary><strong>OpenAI Codex</strong> — <a href="https://github.com/openai/codex">openai/codex</a></summary>

# OpenAI Codex 社区动态日报 — 2026-06-30

---

## 今日速览

今日社区焦点集中在 **GPT-5.5 模型容量不足**和 **Windows 沙箱崩溃** 两大高频 Bug 上，多个用户报告“模型已满”错误中断工作流。同时，官方发布了 `rust-v0.143.0-alpha.31` 版本，并密集合并了一批围绕 **用户消息队列** 和 **会话持久化** 的架构改进。值得注意的是，一个新上报的 **Chronicle 后台屏幕录制消耗配额** 问题引发了广泛讨论，可能涉及隐私与资源管理。

---

## 版本发布

### rust-v0.143.0-alpha.31
- **版本号**: 0.143.0-alpha.31  
- **链接**: [GitHub Release](https://github.com/openai/codex/releases/tag/rust-v0.143.0-alpha.31)  
- **说明**: 本次发布为 Rust 工具链的预发布版本，更新内容暂无详细 changelog，但通常包含对底层引擎的修复和性能优化。

---

## 社区热点 Issues（精选 10 条）

### 1. [#30224](https://github.com/openai/codex/issues/30224) — `X-OpenAI-Internal-Codex-Responses-Lite` 报错“模型不支持”
- **标签**: bug / custom-model / app / config  
- **评论**: 61 | **赞**: 20  
- **摘要**: 用户使用 `X-OpenAI-Internal-Codex-Responses-Lite` 头部时遇到“This model is not supported”错误，疑似内部轻量化响应路径未兼容自定义模型。社区讨论热烈，推测与新版 App 的响应架构变化有关。

### 2. [#30364](https://github.com/openai/codex/issues/30364) — GPT-5.5 推理 token 出现 516/1034/1552 固定边界聚类，导致复杂任务性能下降
- **标签**: bug / model-behavior / rate-limits  
- **评论**: 22 | **赞**: 32  
- **摘要**: 用户发现 `gpt-5.5` 的 `reasoning_output_tokens` 高度集中在 516、1034、1552 等固定数值，且推理质量明显下降。该模式可能是模型内部 token 截断或配额约束的副作用，引发对模型行为异常的关注。

### 3. [#29760](https://github.com/openai/codex/issues/29760) — 模型容量已满，要求切换模型（CLI）
- **标签**: bug / rate-limits / CLI  
- **评论**: 29 | **赞**: 3  
- **摘要**: 用户使用 `codex-cli 0.142.0` 和 `gpt-5.4 high` 模型时频繁出现“Selected model is at capacity”错误，影响持续使用。该问题在近一周多起并发上报（如 #28507、#30575），表明容量问题已蔓延至 CLI 和桌面客户端。

### 4. [#29072](https://github.com/openai/codex/issues/29072) — Windows 下 `apply_patch` 因沙箱安装程序路径问题失败
- **标签**: bug / windows-os / sandbox / tool-calls / app  
- **评论**: 27 | **赞**: 19  
- **摘要**: Windows 版 Codex 桌面应用在调用 `apply_patch` 工具时，因 `codex-windows-sandbox-setup.exe` 无法从包路径正常启动导致补丁操作失败。该 Bug 已持续超过 10 天，社区已有多名用户复现。

### 5. [#30579](https://github.com/openai/codex/issues/30579) — GPT-5.5 容量错误中断正在进行的图像会话（Windows App 26.623.61825）
- **标签**: bug / windows-os / rate-limits / app  
- **评论**: 7 | **赞**: 0  
- **摘要**: 用户在使用基于图像的会话时，模型容量错误直接中断了运行中的任务，且无重试机制。此问题与 #29760 同属容量不足，但影响到了非文本工作流。

### 6. [#30639](https://github.com/openai/codex/issues/30639) — macOS 上 Chronicle 后台屏幕录制每 10 分钟消耗配额，关闭后仍不停止
- **标签**: bug / rate-limits / app / computer-use / memory  
- **评论**: 3 | **赞**: 0  
- **摘要**: 桌面应用内的“Chronicle”功能（可能是自动化屏幕摘要）在后台静默录制屏幕并生成摘要，大量消耗计划配额，且禁用后未能立即停止捕获。该问题引发对隐私和资源管理的担忧，刚上报即获关注。

### 7. [#25269](https://github.com/openai/codex/issues/25269) — macOS Appshot 截图失败：`captureNotFound` 错误
- **标签**: bug / app / computer-use  
- **评论**: 5 | **赞**: 1  
- **摘要**: 使用 Codex Desktop 的 Appshot（应用截图）功能时，即使启动成功，后续截图仍返回“未找到捕获”错误，影响 Computer Use 体验。该问题已存在超过一个月，至今未修复。

### 8. [#30440](https://github.com/openai/codex/issues/30440) — Codex 使用内置 pnpm 而非主机工具链，导致构建失败
- **标签**: bug / sandbox / app  
- **评论**: 6 | **赞**: 9  
- **摘要**: 用户发现 Codex 在执行项目构建时优先使用捆绑的 pnpm 版本，而非系统中安装的版本，导致与自定义构建脚本不兼容。虽然赞数不多，但反映了沙箱隔离策略对开发工具链的干扰。

### 9. [#30525](https://github.com/openai/codex/issues/30525) — 桌面应用空闲时意外消耗大量配额（已关闭）
- **标签**: bug / rate-limits / app  
- **评论**: 4 | **赞**: 0  
- **摘要**: 用户报告 Codex 桌面 App 在闲置状态下仍频繁消耗 API 配额，已确认与某个后台服务有关。该 Issue 已被关闭（可能已复现或修复），但未公开根因。

### 10. [#2880](https://github.com/openai/codex/issues/2880) — 新增功能：复制/导出消息为 Markdown
- **标签**: enhancement / TUI  
- **评论**: 24 | **赞**: 71  
- **摘要**: 该功能请求已持续近一年（创建于2025-08-29），获得 71 个赞。用户希望能在 TUI 中一键复制对话为 Markdown 格式，以便用于外部文档或 Issue 引用。至今未实现，社区呼声极高。

---

## 重要 PR 进展（精选 10 条）

### 1. [#28307](https://github.com/openai/codex/pull/28307) — 通过 app-server 队列 TUI 跟进消息（已合并）
- **作者**: efrazer-oai  
- **说明**: 将 TUI 中排队等待的普通用户消息通过 app-server 持久化，使其在 TUI 进程重启后仍能存活。这是用户消息队列架构的关键拼图。

### 2. [#28267](https://github.com/openai/codex/pull/28267) — 通过核心空闲扩展分发排队用户消息
- **作者**: efrazer-oai  
- **说明**: 将 queued user messages 接入核心的 `on_thread_idle` 扩展路径，确保队列中的消息能在线程空闲时按序处理，避免与目标（goals）冲突。

### 3. [#28265](https://github.com/openai/codex/pull/28265) — 在空闲轮次边界接受用户提交
- **作者**: efrazer-oai  
- **说明**: 允许排队用户消息在空闲时以原子方式进入系统，同时保留常规用户提交的行为（如上下文、Response API 元数据等），为消息队列功能提供基础。

### 4. [#28266](https://github.com/openai/codex/pull/28266) — 添加持久化用户消息队列存储
- **作者**: efrazer-oai  
- **说明**: 新增 `queue_1.sqlite` 数据库及事务性存储，确保队列消息在客户端或 app-server 重启后不会丢失，并支持跨进程的消费确认。

### 5. [#25283](https://github.com/openai/codex/pull/25283) — 同步运行时工作区根目录到线程设置
- **作者**: efrazer-oai  
- **说明**: 将 `runtimeWorkspaceRoots` 添加到线程设置中，使队列调度时能获取正确的上下文工作区，保证排队任务与直接任务看到相同的环境。

### 6. [#27932](https://github.com/openai/codex/pull/27932) — 允许模型更改工作目录（已合并）
- **作者**: fcoury-oai  
- **说明**: 支持长时间运行的任务中，模型自动切换工作目录（如切换到其他代码库或生成的工作区）。工具路径解析不再固定于启动目录，提升复杂任务灵活性。

### 7. [#25629](https://github.com/openai/codex/pull/25629) — 添加远程插件搜索工具（已合并）
- **作者**: adaley-openai  
- **说明**: 新增模型可见的 `search_plugins` 工具，允许模型通过自然语言查询 ChatGPT Plugin Store 发现全局插件。该功能被 `Plugins` 和 `RemotePlugin` 特性门控。

### 8. [#28378](https://github.com/openai/codex/pull/28378) — 重试模型列表获取失败
- **作者**: sayan-oai  
- **说明**: 修复 Rust 发布工作流中偶尔因网络瞬断导致模型目录获取失败的问题，添加重试机制，提高 CI/CD 稳定性。

### 9. [#28335](https://github.com/openai/codex/pull/28335) — 修复不兼容的状态迁移（已合并）
- **作者**: fcoury-oai  
- **说明**: 当 SQLx 检测到已应用的运行时数据库迁移被修改或部分应用时，Codex 无法启动。此 PR 增强了 `codex doctor` 的诊断和修复能力，能够自动识别并修复冲突的迁移状态。

### 10. [#28314](https://github.com/openai/codex/pull/28314) — 跳过瞬态事件的后台持久化
- **作者**: jif-oai  
- **说明**: 优化事件管道：在提交事件到后台存储前先检查 rollout 事件策略，过滤掉不必持久化的瞬态事件，减少 I/O 开销并提升高并发场景下的性能。

---

## 功能需求趋势

从所有 Issues 和 PR 中可提炼出当前社区最关注的四个方向：

1. **配额管理与容量透明度** — 用户希望：  
   - 提供配额重置机制（#30686）  
   - 清晰展示容量状态，避免模型“已满”错误打断工作流  
   - 后台活动（如 Chronicle）的配额消耗可视化  

2. **消息持久化与跨会话连续性** — 近期大量 PR 围绕用户消息队列展开，社区期待：  
   - CLI/桌面应用的“离线”排队功能  
   - 跨重启、跨客户端的消息恢复  
   - 更可靠的消息排序与原子提交  

3. **沙箱与工具链兼容性** — Windows 和 Linux 用户持续反馈：  
   - 沙箱安装程序路径问题（#29072、#29418）  
   - 内置 pnpm、bubblewrap 检测不准确等问题  
   - 需要更好的环境探测与回退策略  

4. **“零成本”功能改进** — 如复制为 Markdown（#2880）、服务层定价选项（#2916），虽非高频 Bug，但获得了极高的社区支持率，表明用户对轻量级生产力增强的强烈需求。

---

## 开发者关注点

- **模型容量错误成为最大痛点**：过去 24 小时内新增的 #30575、#30579、#30639 等 Issue 均指向“模型已满”或配额异常消耗，严重影响日常开发。开发者呼吁 OpenAI 提供更稳定的容量分配和错误重试机制。
- **Windows 平台稳定性亟待提升**：沙箱设置、补丁应用、Git 认证等基础功能在 Windows 上频繁故障（#29072、#29492、#29828），部分问题已存在数周未修复。
- **夜间活动消耗配额引发隐私担忧**：#30639 揭露的 Chronicle 后台屏幕录制功能，可能在用户不知情下耗尽计划配额。开发者希望增加明确的开关和配额消耗提示。
- **LLM 行为异常需透明化**：#30364 中 GPT-5.5 的推理 token 固定边界模式让用户怀疑模型输出被硬截断，影响复杂任务质量。社区要求 OpenAI 披露是否存在 token 预算限制或输出模板。
- **功能请求长期未回应**：#2880（复制 Markdown）创建近一年仍无进展，而 #2916（服务层支持）也停滞多月。用户对部分呼声极高的功能改进感到失望，希望官方加速路线图透明化。

---

*数据来源：GitHub openai/codex 仓库，截至 2026-06-30 18:00 UTC。*

</details>

<details>
<summary><strong>Gemini CLI</strong> — <a href="https://github.com/google-gemini/gemini-cli">google-gemini/gemini-cli</a></summary>

好的，各位开发者，以下是 2026 年 6 月 30 日的 Gemini CLI 社区动态日报。

---

# Gemini CLI 社区动态日报 | 2026-06-30

## 今日速览

今日社区动态聚焦于项目稳定性和智能性的提升。一方面，多个针对 **子代理行为异常** 和 **Shell 命令执行卡死** 的核心 Bug 正在被积极修复；另一方面，围绕 **AST（抽象语法树）感知** 和 **组件级评估** 的长期规划讨论热烈，标志着项目正从功能堆砌迈向精细化治理。此外，`settings.json` 的解析兼容性问题和小型 UI 优化也得到了解决。

## 版本发布

**`v0.51.0-nightly.20260630.gae0a3aa7b`**
- **链接**: [Release](https://github.com/google-gemini/gemini-cli/compare/v0.51.0-nightly.20260629.gae0a3aa7b...v0.51.0-nightly.20260630.gae0a3aa7b)
- **内容**: 常规晚间版本发布，包含今日合并的各项修复和改进。

## 社区热点 Issues

1.  **[Bug] 通用子代理挂起 (`#21409`)**
    - **摘要**: 当 CLI 将任务委托给通用子代理时，会无限期挂起，即使是创建文件夹这类简单操作也无法完成。用户手动禁止委托可绕过此问题。
    - **关注度**: 评论 7，👍 8 | [🔗](https://github.com/google-gemini/gemini-cli/issues/21409)
    - **重要性**: 这是一个严重影响核心“代理”功能可用性的 Bug，用户反馈强烈，是目前社区的最大痛点之一。

2.  **[Bug] 子代理在达到最大轮次后误报“成功” (`#22323`)**
    - **摘要**: 当`codebase_investigator`子代理因达到`MAX_TURNS`限制而未能完成任务时，系统却错误地将终止原因报告为“GOAL”（成功），掩盖了任务被中断的事实。
    - **关注度**: 评论 8，👍 2 | [🔗](https://github.com/google-gemini/gemini-cli/issues/22323)
    - **重要性**: 这是一个严重的代理状态管理问题，会导致用户误判任务执行情况，影响对自动化流程的信任。

3.  **[Epic] 稳健的组件级评估 (`#24353`)**
    - **摘要**: 此 EPIC 旨在建立更为精细的组件级行为评估体系，作为已有“行为评估”测试的后续。目前已生成76个测试用例，覆盖6个Gemini模型，目标是实现更小的测试单元和更快的反馈循环。
    - **关注度**: 评论 7 | [🔗](https://github.com/google-gemini/gemini-cli/issues/24353)
    - **重要性**: 这代表了项目质量控制策略从端到端测试向模块化、单元化演进，是提升CI/CD效率和模型迭代质量的关键举措。

4.  **[Epic] 探索 AST 感知的文件读取与搜索 (`#22745`)**
    - **摘要**: 探讨引入 AST（抽象语法树）感知能力，以实现更精确的方法边界读取、代码导航和代码库映射，从而减少不必要的 Token 消耗和修复因读取不匹配导致的工具调用失败问题。
    - **关注度**: 评论 7 | [🔗](https://github.com/google-gemini/gemini-cli/issues/22745)
    - **重要性**: 这是一个极具前瞻性的技术方向。若能实现，将极大提升 CLI 对大型复杂代码库的理解和处理效率，是当前“精准代码分析”需求的直接体现。

5.  **[Bug] Shell 命令执行后卡死 (`#25166`)**
    - **摘要**: 代理在执行完一个简单的 CLI 命令后，界面仍显示“等待输入”并持续挂起，尽管命令早已结束。
    - **关注度**: 评论 4，👍 3 | [🔗](https://github.com/google-gemini/gemini-cli/issues/25166)
    - **重要性**: 极差的用户体验问题，打断了自动化工作流，是仅次于“代理挂起”的交互性痛点。

6.  **[Bug] Gemini 不主动使用技能和子代理 (`#21968`)**
    - **摘要**: 即使用户自定义了技能（如 Gradle、Git）并提供了详细描述，Gemini 也很少主动调用它们，除非用户明确指令。
    - **关注度**: 评论 6 | [🔗](https://github.com/google-gemini/gemini-cli/issues/21968)
    - **重要性**: 直指代理规划能力（Agentic Planning）的薄弱环节，与用户对“智能助手”的期望存在差距。

7.  **[Bug] 自动记忆系统：确定性脱敏与日志优化 (`#26525`)**
    - **摘要**: Auto Memory 将内容发送至模型后才进行脱敏，存在安全隐患；同时，其日志记录可能泄露已保存的技能内容。
    - **关注度**: 评论 5 | [🔗](https://github.com/google-gemini/gemini-cli/issues/26525)
    - **重要性**: 涉及安全和隐私的核心问题，是项目从实验性走向生产环境必须解决的问题。

8.  **[Bug] 自动记忆应停止对低信号会话的无限重试 (`#26522`)**
    - **摘要**: Auto Memory 仅将成功读取的会话标记为“已处理”，若代理判断某个会话“低价值”而跳过，该会话将永远处于“未处理”状态并被反复尝试。
    - **关注度**: 评论 5 | [🔗](https://github.com/google-gemini/gemini-cli/issues/26522)
    - **重要性**: 这是典型的资源浪费和逻辑缺陷，影响自动记忆系统的效率和准确性。

9.  **[Feature] 代理应阻止/劝阻破坏性行为 (`#22672`)**
    - **摘要**: 在涉及`git reset`、`--force` 或数据库操作等场景时，模型有时会采取破坏性手段，而非选择更安全的替代方案。
    - **关注度**: 评论 3，👍 1 | [🔗](https://github.com/google-gemini/gemini-cli/issues/22672)
    - **重要性**: 直指 AI 安全的核心，用户需要的是一个能理解操作后果并优先选择安全路径的智能体。

10. **[Bug] 浏览器子代理在 Wayland 环境下运行失败 (`#21983`)**
    - **摘要**: 在 Linux Wayland 显示协议下，浏览器子代理会启动失败，无法正常工作。
    - **关注度**: 评论 4，👍 1 | [🔗](https://github.com/google-gemini/gemini-cli/issues/21983)
    - **重要性**: 特定平台兼容性问题，对于 Linux 用户社区来说是关键痛点。

## 重要 PR 进展

1.  **`fix(cli): parse commented settings.json in memory bootstrap` (`#28219`)**
    - **重要性**: 修复了因`settings.json`中包含注释（JSON5格式）导致内存模式启动失败的兼容性问题。
    - **状态**: 已开放 | [🔗](https://github.com/google-gemini/gemini-cli/pull/28219)

2.  **`fix(core): drop late tool calls after SIGINT cancellation` (`#28096`)**
    - **重要性**: 修复了用户发送中断信号（Ctrl+C）后，系统仍会执行延迟到达的旧工具调用结果的 Bug，防止状态混乱。
    - **状态**: 已开放 | [🔗](https://github.com/google-gemini/gemini-cli/pull/28096)

3.  **`fix(vscode-ide-companion): add missing activate() Disposables` (`#27936`)**
    - **重要性**: 修复了 VS Code 插件中因逗号操作符误用导致部分资源（Disposables）未被正确注册和释放的问题。
    - **状态**: **已关闭** | [🔗](https://github.com/google-gemini/gemini-cli/pull/27936)

4.  **`fix(core): remove leading space in camelToSpace for capitalized keys` (`#27942`)**
    - **重要性**: 修复了键名转换函数在遇到首字母大写的键（如“Id”）时，错误地插入前置空格的UI显示问题。
    - **状态**: **已关闭** | [🔗](https://github.com/google-gemini/gemini-cli/pull/27942)

5.  **`fix(vscode-ide-companion): register Disposables leaked by comma operators` (`#28100`)**
    - **重要性**: 与 #27936 类似，修复了VS Code插件中另一处因逗号操作符导致的资源泄露问题。
    - **状态**: 已开放 | [🔗](https://github.com/google-gemini/gemini-cli/pull/28100)

6.  **`feat(core): implement MCP elicitation (form + url) capability` (`#28089`)**
    - **重要性**: 实现了 MCP 协议中的表单和 URL 提权功能，是扩展 MCP 工具链能力的重要一步。
    - **状态**: 已开放 | [🔗](https://github.com/google-gemini/gemini-cli/pull/28089)

7.  **`fix(core): limit recursive reasoning turns per single user request` (`#28164`)**
    - **重要性**: 为了防止模型陷入无限递归推理循环，设置了每轮用户请求的上限（15次），保护本地资源和API配额。
    - **状态**: 已开放 | [🔗](https://github.com/google-gemini/gemini-cli/pull/28164)

8.  **`fix(core): strip thoughts from scrubbed history turns and resolve thought leakage` (`#27971`)**
    - **重要性**: 修复了“思维泄露”问题，即模型的内部推理过程被写入历史记录，导致后续对话出现幻觉或异常模仿。
    - **状态**: 已开放 | [🔗](https://github.com/google-gemini/gemini-cli/pull/27971)

9.  **`fix(core-tools): show ellipsis on multi-line edit snippets` (`#28126`)**
    - **重要性**: 优化了工具调用的 UI 展示，当编辑内容过长或多行时，使用省略号表示截断，避免显示不全。
    - **状态**: 已开放 | [🔗](https://github.com/google-gemini/gemini-cli/pull/28126)

10. **`Harden file-write scope: stop sandbox/auto-accept writes to .gemini and .gitconfig` (`#28215`)**
    - **重要性**: 安全增强，阻止代理在沙箱或自动确认模式下向`.gemini/`和`.gitconfig`等核心配置文件写入，防止提示注入攻击导致的权限逃逸。
    - **状态**: **已关闭** | [🔗](https://github.com/google-gemini/gemini-cli/pull/28215)

## 功能需求趋势

1.  **精准代码理解与导航**: 社区对 AST 感知的讨论热度很高，核心诉求是希望 CLI 能像资深开发者一样，通过理解代码结构（而非字符串匹配）来执行更高效、更精确的任务，如调用特定方法、搜索接口、映射项目依赖等。

2.  **评估体系精细化**: `#24353` 组件级评估 EPIC 的推进，表明社区和项目维护者都意识到仅靠端到端测试不够。未来趋势是建立更细粒度的行为评估矩阵，以量化衡量模型在特定子任务（如文件编辑、Shell执行）上的表现。

3.  **代理行为的安全与可预测性**: 从“阻止破坏性行为”到“子代理误报成功”，再到“思维泄露”，用户对代理行为的可预测性和安全性要求越来越高。让“AI 助手”在未知或高风险操作前进行确认，并准确报告其自身状态，是当前最紧迫的功能需求。

4.  **优化的“记忆”系统**: 对 Auto Memory 的改进集中于**安全性**（脱敏）、**效率**（停止无效重试）和**准确性**（过滤无效补丁）。社区希望一个既能记住关键信息，又不会泄露隐私或浪费资源的“记忆”机制。

## 开发者关注点

- **代理** (Agent) 是绝对核心的痛点与关注点，问题集中在 **挂起、误报状态、不主动使用工具** 上。这表明代理的规划、执行与状态管理能力仍有巨大提升空间。
- **Shell 命令执行** 的稳定性是自动化工作流的基础。执行后卡死、确认提示崩溃等问题严重影响了用户对 CLI 作为“自动化工具”的信心。
- **配置文件的健壮性** 成为隐性需求。`settings.json` 对注释的高敏性暴露了项目在处理非标准配置文件时的脆弱性，开发者期望一个更宽容、更稳定的配置解析系统。
- **安全问题** 正从“功能”变为“特性”。无论是“提示注入”导致权限逃逸，还是“自动记忆”的隐私风险，都表明社区开始要求 CLI 具备默认安全（Secure by Default）的特性。

</details>

<details>
<summary><strong>GitHub Copilot CLI</strong> — <a href="https://github.com/github/copilot-cli">github/copilot-cli</a></summary>

# GitHub Copilot CLI 社区动态日报 | 2026-06-30

## 今日速览
今日发布 **v1.0.66-2**，新增了插件同名共存、用户设置读写、LSP 日志查看等关键能力。社区活跃，共 23 条议题更新，其中 **#1665**（项目级插件作用域）已关闭并获 17 个 👍，**#3976**（`tgrep` 索引器 OOM 杀宿主机）成为新的高优先级 Bug。PR 方面，自动议题分类工作流被合并，将提升社区治理效率。

---

## 版本发布

### v1.0.66-2
- **允许不同插件注册同名 skill 共存**：解决多插件冲突，提高扩展性。
- **集成插件可读写 CLI 用户设置**：第三方集成能动态修改配置，扩展生态能力。
- **查看 LSP 服务日志**：通过 `/lsp logs` 和 `read_agent` 命令，便于调试 agent 行为。
- **缺失 gh CLI 时提示安装**：在 GitHub 仓库内检测并引导安装，降低上手门槛。
- **为 prompt 渲染增加 GitHub 附件变体**：对 GitHub 链接/附件提供更丰富的展示效果。

---

## 社区热点 Issues（精选 10 条）

### 1. [#1665 Support Copilot CLI Plugins Scoped to Project or Repository](https://github.com/github/copilot-cli/issues/1665) — **已关闭** ⭐17  
**重要性：** 当前插件为全局安装，无法按项目隔离。该需求被社区强烈呼吁，关闭后应已实现或进入路线图。  
**社区反应：** 10 条评论，89% 点赞支持。

### 2. [#1799 How to turn off alt-screen views?](https://github.com/github/copilot-cli/issues/1799) — **开放** ⭐7  
**重要性：** 新版 alt-screen 引发多用户不适，要求回归原始模式。反映终端渲染对用户习惯的冲击。  
**社区反应：** 10 条评论，持续讨论回退方案。

### 3. [#3874 `preToolUse` agent hook denial does not work](https://github.com/github/copilot-cli/issues/3874) — **开放**  
**重要性：** hook 系统核心功能失效，拒绝所有命令的 hook 无法生效，可能影响安全策略。  
**社区反应：** 2 条评论，值得安全团队立即关注。

### 4. [#3948 web_fetch: TypeError: fetch failed](https://github.com/github/copilot-cli/issues/3948) — **开放**  
**重要性：** 所有 `web_fetch` 工具调用均返回 `fetch failed`，即使网络正常。影响基于工具的 Agent 工作流。  
**社区反应：** 2 条评论，用户认为并非代理配置问题。

### 5. [#3957 Unable to scroll through history using trackpad on MBP](https://github.com/github/copilot-cli/issues/3957) — **已关闭** ⭐5  
**重要性：** 触控板滚动导致选中历史命令而非滚动窗口，严重影响 macOS 用户交互体验。  
**社区反应：** 1 条评论，已标记为关闭（推测已修复或临时方案）。

### 6. [#3936 Ctrl+G should expand paste tokens to full text in $EDITOR](https://github.com/github/copilot-cli/issues/3936) — **开放**  
**重要性：** 当 `compactPaste` 启用时，Ctrl+G 会写入粘贴 token 而非真实内容，导致编辑困难。与 Claude Code 行为对齐的需求。  
**社区反应：** 2 条评论，用户期待快捷键展开。

### 7. [#3971 Feature Request: Full file-tree browser for repository-backed sessions](https://github.com/github/copilot-cli/issues/3971) — **开放**  
**重要性：** 仓库会话仅显示 git Changes 面板，缺少文件夹会话中完整的文件树浏览器，影响项目导航效率。  
**社区反应：** 1 条评论，功能请求合理。

### 8. [#3976 native `tgrep` indexer OOM-kills the host on large monorepos](https://github.com/github/copilot-cli/issues/3976) — **开放**  
**重要性：** 原生 `tgrep` 索引器无内存上限，在大单体仓库中直接耗尽内存杀死宿主机。高危性能 Bug。  
**社区反应：** 0 条评论（刚提交），但影响范围极大。

### 9. [#3954 Bug: `explore` tool hardcodes model to `gpt-5.4-mini`](https://github.com/github/copilot-cli/issues/3954) — **开放**  
**重要性：** 使用自定义模型（如 DeepSeek）时，`explore` 工具硬编码为 `gpt-5.4-mini`，导致 API 调用失败。破坏多模型支持。  
**社区反应：** 0 条评论，但关联自定义模型配置，值得跟进。

### 10. [#3972 UI displaying continuous stream of characters representing mouse movements](https://github.com/github/copilot-cli/issues/3972) — **开放**  
**重要性：** 首次加载时终端输出大量鼠标移动字符，完全破坏 UI。属于严重渲染问题。  
**社区反应：** 0 条评论（新提交），疑似鼠标事件捕获异常。

---

## 重要 PR 进展

仅 1 条 PR 在统计时间内更新：

### [#2587 Add automated issue classification with GitHub Agentic Workflows](https://github.com/github/copilot-cli/pull/2587) — **已合并**  
**内容：** 引入基于 GitHub Agentic Workflows 的 AI 自动分类工作流。当 Issue 被创建或重新打开时，自动打上 `area:` 和 `triage` 标签。  
**价值：** 显著提升社区 Issue 治理效率，减少人工分类负担。该 PR 于今日合并，表明 Copilot CLI 团队正在利用自己的 ai 工具优化项目管理。

---

## 功能需求趋势

从近 24 小时更新的 Issues 中，社区最关注的功能方向为：

- **插件作用域与隔离**（#1665）：从用户全局扩展为项目/仓库级别。
- **终端渲染与交互优化**（#1799、#3957、#3972）：alt-screen 可控性、滚动行为、鼠标事件过滤。
- **输入编辑体验**（#3936）：Ctrl+G 展开粘贴 token，对标 Claude Code。
- **网络与工具**（#3948、#3976）：`web_fetch` 可靠性、大 repo 索引器内存控制。
- **多模型支持**（#3954）：防止工具硬编码特定模型，尊重用户自定义配置。
- **会话文件管理**（#3971）：仓库会话提供完整文件树。
- **MCP 配置稳定性**（#3893、#3973）：同名 MCP Server 加载策略、Windows OAuth 回环端口冲突。
- **桌面通知**（#2941）：CLI 在后台需用户输入时弹出系统通知。

---

## 开发者关注点

- **🎯 痛点高频出现**：  
  - **内存泄漏 / OOM**：`tgrep` 索引器无上限（#3976）是近期最严重的性能问题。  
  - **终端渲染回归**：alt-screen 模式强制启用、滚动异常、鼠标字符污染 UI（#1799、#3957、#3972）。  
  - **模型硬编码**：`explore` 工具无视用户配置（#3954），影响自定义模型用户。  
  - **文件拖拽失效**：macOS 上文件拖拽到 Copilot 应用无响应（#3955）。

- **📢 高频需求**：  
  - 项目级插件（#1665）得到广泛支持，预计会优先推进。  
  - 输入编辑体验改进（#3936）与 Claude Code 对齐的需求增长。  
  - 仓库会话文件树（#3971）解决工作流效率瓶颈。

- **⚠️ 安全与稳定性**：  
  - Hook 拒绝策略失效（#3874）可能被利用绕过安全控制。  
  - MCP OAuth 在 Windows 上反复重认证失败（#3973），影响企业部署。

---

*数据来源：[github/github/copilot-cli](https://github.com/github/copilot-cli) | 统计区间：2026-06-29 ～ 2026-06-30*

</details>

<details>
<summary><strong>Kimi Code CLI</strong> — <a href="https://github.com/MoonshotAI/kimi-cli">MoonshotAI/kimi-cli</a></summary>

# Kimi Code CLI 社区动态日报 | 2026-06-30

## 今日速览
今日社区没有新版本发布和 Issue 更新，主要关注两项 PR 进展：一项优化了 Shell UI 中用户输入的可视性（高亮蓝色 + 分隔线），另一项新增了 `--prompt-interactive` 选项，允许在启动 Shell 时传入初始提示并保持交互会话。这两项改动均聚焦于提升命令行交互体验。

---

## 版本发布
无（过去24小时内没有新 Releases）

---

## 社区热点 Issues
（今日无新增或更新的 Issue，暂无可推荐的热点）

---

## 重要 PR 进展
以下是过去24小时内更新的 2 条 Pull Request，均处于可见状态：

### 1. [#1600] feat(shell): highlight user input with bright blue and separator for better visibility  
**作者**: liuchong | **创建**: 2026-03-27 | **更新**: 2026-06-30 | **状态**: Open  
**摘要**: 改进 Shell UI 中用户消息的视觉区分：  
- 将用户回显文本设置为亮蓝色（`#007AFF`），与系统消息形成统一配色  
- 在每个用户输入下方添加一条全宽分隔线，增强视觉层次  
**链接**: [PR #1600](https://github.com/MoonshotAI/kimi-cli/pull/1600)

### 2. [#2246] feat(shell): add --prompt-interactive option  
**作者**: shuizhongyueming | **创建**: 2026-05-12 | **更新**: 2026-06-30 | **状态**: Closed（已合并）  
**摘要**: 新增 CLI 选项 `--prompt-interactive`（短名 `-P`），允许用户在启动 Shell 界面时传入初始提示（prompt），并在回答后保持交互式会话，方便连续追问。  
**相关关联**: 解决 Issue #2240  
**链接**: [PR #2246](https://github.com/MoonshotAI/kimi-cli/pull/2246)

---

## 功能需求趋势
从今日活跃的 PR 可以看出，社区对以下方向持续关注：

- **Shell 交互体验优化**：用户输入的可视性（颜色、分隔符）是提升日常使用效率的常见需求。  
- **CLI 选项扩展**：新增 `--prompt-interactive` 体现用户希望灵活控制会话启动方式，尤其适合脚本化和渐进式对话场景。  
- **可配置性与统一性**：颜色使用的统一规范（如 `#007AFF`）表明社区对界面一致性的重视。

---

## 开发者关注点
今日无 Issue 或 Comments 直接反映痛点，但从 PR 内容可推测：

- 用户可能希望在 Shell 中更清晰地区分自己的输入与 AI 回复，避免视觉混淆。  
- 需要一种在不牺牲交互性的情况下预先注入初始提示的能力，说明部分用户工作流依赖预设指令。  
- 缺少反馈渠道意味着当前版本稳定性较高，或社区活跃度暂时较低。建议开发者关注后续 Issue 中的实际使用反馈。

---

*数据来源: [MoonshotAI/kimi-cli](https://github.com/MoonshotAI/kimi-cli) | 采集时间: 2026-06-30*

</details>

<details>
<summary><strong>OpenCode</strong> — <a href="https://github.com/anomalyco/opencode">anomalyco/opencode</a></summary>

# OpenCode 社区动态日报 | 2026-06-30

## 📰 今日速览
- 社区对 **模型 fallback 机制**（#7602）和 **YOLO 模式**（#8463）的诉求持续升温，分别获 90👍 和 89👍，成为当前最受关注的功能方向。
- 多个 **MCP OAuth token 刷新** 问题浮出水面（#34582、#34592），开发者反馈 token 过期后无法自动续期，影响远程服务使用。
- Robin1987China 贡献的一批底层修复 PR 今日合入，覆盖 Bedrock DeepSeek 模型 ID 保留、frequencyPenalty/Gemini 参数传递、文件写入目录创建等关键缺陷。

---

## 🔥 社区热点 Issues（Top 10）

1. **[Feature] Native Model Fallback / Failover Support**  
   #7602 | 需求：支持在不同模型之间自动 fallback（如模型 A 超时 → 切换到模型 B），当前仅支持同一 provider 下的 provider fallback。社区反应强烈，90👍 表示这是提升 Agent 任务稳定性的必备能力。  
   → https://github.com/anomalyco/opencode/issues/7602

2. **[Feature] 添加 `--dangerously-skip-permissions`（YOLO 模式）**  
   #8463 | 需求：在受信任环境或自动化流程中跳过权限确认弹窗，由 surma 提出。89👍 表明开发者在 CI/CD 等场景中急需该能力。  
   → https://github.com/anomalyco/opencode/issues/8463

3. **[Feature] 请求 GitHub Copilot 自动模型路由 API 访问 + chat.model 插件 hook**  
   #20235 | 需求：获取 Copilot 的 `/models/session` 接口以实现类似 VS Code 的智能模型路由。25👍，显示用户希望 OpenCode 能复用 Copilot 的底层模型选择能力。  
   → https://github.com/anomalyco/opencode/issues/20235

4. **[Bug] Custom OpenAI-compatible provider options 未传递到 API 调用**  
   #5674 | 问题：`opencode.json` 中配置的 `baseURL`、`apiKey` 等 `options` 未被实际传递给 `@ai-sdk/openai-compatible` 的 API 调用。影响所有自定义兼容 provider 用户，13👍，25条评论。  
   → https://github.com/anomalyco/opencode/issues/5674

5. **[Feature] 在 model selector 中暴露 GitHub Copilot "Auto" 选项**  
   #25239 | 需求：让用户能在 OpenCode 的模型选择器中直接选择 Copilot 的自动路由模式。14👍，与 #20235 呼应。  
   → https://github.com/anomalyco/opencode/issues/25239

6. **[Bug] 通过 SSH + tmux 在 Ghostty 中复制到剪贴板失败**  
   #15907 | 问题：复制操作显示“已复制到剪贴板”但实际未更新系统剪贴板。影响远程开发场景，10👍，10条评论。  
   → https://github.com/anomalyco/opencode/issues/15907

7. **[Bug] OpenAI GPT 模型（ChatGPT Plus 浏览器认证）无响应**  
   #34250 | 问题：使用 ChatGPT Plus 浏览器认证时，所有 GPT 模型接收 prompt 但不返回任何响应（无错误提示）。新近报告，4条评论。  
   → https://github.com/anomalyco/opencode/issues/34250

8. **[Bug] Remote MCP OAuth：令牌过期后未使用 refresh token 刷新**  
   #34582 | 问题：OpenCode 成功存储 refresh token，但 access token 过期后未执行 refresh-grant 操作，导致请求 401。影响 OAuth 保护的 MCP 服务器。  
   → https://github.com/anomalyco/opencode/issues/34582

9. **[Bug] MCP OAuth token 刷新缺少 resource 参数 → Atlassian 返回 401**  
   #34592 | 问题：刷新 OAuth token 时未按 RFC 9728 添加 `resource` 参数，导致 Atlassian OAuth 服务返回 401。  
   → https://github.com/anomalyco/opencode/issues/34592

10. **[Bug] MiniMax-M3 模型无视用户配置的 `thinking.type`，强制使用 `adaptive`**  
    #34570 | 问题：无论用户设置 `"enabled"` 还是其他值，OpenCode 在代码中硬编码为 `"adaptive"`，绕过用户意图。  
    → https://github.com/anomalyco/opencode/issues/34570

---

## 🔧 重要 PR 进展（Top 10）

1. **fix(app): restore prompt cursor on focus**  
   #34175 | 已合入。修复 TUI/App 中 prompt 光标在 focus 时未恢复到之前位置的 bug，提升输入体验。  
   → https://github.com/anomalyco/opencode/pull/34175

2. **fix: respect agent temperature config when capabilities.temperature is false**  
   #34583 | 已合入。修复当模型声明的 `capabilities.temperature` 为 false 时，agent 级别 `temperature` 配置被丢弃的问题（影响自定义 OpenAI 兼容 provider）。  
   → https://github.com/anomalyco/opencode/pull/34583

3. **feat(app): autocomplete mcp resources**  
   #34597 | 进行中。为 TUI/App 的 prompt 添加 MCP resource 自动补全（在 agent 名称后 @）。提升 MCP 资源使用的便捷性。  
   → https://github.com/anomalyco/opencode/pull/34597

4. **fix: preserve Bedrock DeepSeek model ids**  
   #34441 | 进行中。修复 Bedrock 上的 DeepSeek 模型 ID（如 `deepseek.v3.2`）被错误当作跨区域模型 ID 处理导致无法使用的 bug。  
   → https://github.com/anomalyco/opencode/pull/34441

5. **fix(desktop): persist last active url**  
   #34595 | 已合入。桌面端页面路由状态持久化到 localStorage，重启后恢复上次打开的桌面路由，改善用户体验。  
   → https://github.com/anomalyco/opencode/pull/34595

6. **feat(tui): insert absolute path for unsupported pastes and toast on failure**  
   #34307 | 进行中。修复 Windows 下粘贴失败（clipboard.read 被拒绝）时无反馈的问题，改为插入绝对路径提示并 toast 异常。  
   → https://github.com/anomalyco/opencode/pull/34307

7. **fix(app): preserve macos titlebar inset**  
   #34594 | 已合入。确保 macOS 标题栏交通灯区域始终保留安全间距，窗口窄时也不遮挡原生控件。  
   → https://github.com/anomalyco/opencode/pull/34594

8. **fix(app): autocomplete configured references**  
   #34308 | 进行中。修复 desktop/app 的 @ 自动补全中未加载项目级 reference 的问题，现在 reference 能正确显示并插入为 directory file part。  
   → https://github.com/anomalyco/opencode/pull/34308

9. **fix(llm): forward bedrock.thinking provider option to Bedrock Converse**  
   #34590 | 进行中。让 Bedrock 用户可以通过 `providerOptions.bedrock.thinking` 配置扩展思考（extended thinking），修复该配置被静默忽略的 bug（对应 #33630）。  
   → https://github.com/anomalyco/opencode/pull/34590

10. **Multiple fixes by Robin1987China (合入)**  
    - #33330：修复 compaction 时 boundary 消息被重复计入 summary 的 bug  
    - #33326：将 `frequencyPenalty`、`presencePenalty`、`seed` 转发到 Gemini provider  
    - #33099：修复 `sampledChecksum` 最后采样窗口只覆盖 2KB 的问题  
    - #32919：恢复 Copilot 聊天流 chunk 类型安全  
    这些 PR 今日均已完成合入，显著提升了底层稳定性和多 provider 兼容性。  
    汇总链接：https://github.com/anomalyco/opencode/pulls?q=author%3ARobin1987China+is%3Apr

---

## 🧭 功能需求趋势
从今日活跃的 Issue 中提炼出社区最关注的四大方向：

- **模型与 Provider 增强**  
  Model Fallback（#7602）、Copilot 自动路由（#20235）、JetBrains 中推理 effort 选择（#34551）、/export 桌面端支持（#31453）

- **安全与权限优化**  
  YOLO 模式（#8463）、危险操作跳过权限提示

- **MCP 集成成熟度**  
  OAuth token 刷新（#34582、#34592）、MCP 资源自动补全（#34597）、LLM 丢失 MCP 列表（#34579）

- **多语言与国际化**  
  孟加拉语 UI 支持（#34593）、意大利语翻译合入（#30719）

- **Agent 与用户体验细节**  
  /insights 作为一等公民（#12981）、桌面端 session 标题 hover 编辑（#34589）

---

## 🐛 开发者关注点（痛点 / 高频问题）

- **配置与认证**  
  - Custom provider 的 `baseURL`/`apiKey` 未传递（#5674）  
  - ChatGPT Plus 浏览器认证无响应（#34250）  
  - Paid Go 订阅扣费未激活（#32420）——虽已关闭但仍有大量同类报告

- **MCP / 外部服务集成**  
  - OAuth token 过期后不刷新（#34582）  
  - Atlassian OAuth 因缺少 resource 参数返回 401（#34592）  
  - MCP 服务器列表突然丢失（#34579）  
  - Remote MCP with 9router：MiniMax-M3 不支持图片输入（#34596）

- **模型推理问题**  
  - MiniMax-M3 `thinking.type` 被硬编码覆盖（#34570）  
  - Bedrock 上扩展思考配置被忽略（#33630，PR #34590 已修复）  
  - 开源模型（Qwen、GLM、DeepSeek）零响应（#29605）

- **桌面与 TUI 功能性 bug**  
  - 剪贴板复制在 SSH+tmux 下无效（#15907）  
  - Windows ARM64 安装包缺少 OpenCode.exe（#34581）  
  - 使用 /undo 无法恢复 AI 编辑的文件（#34587）  
  - ACP 会话忽略客户端选择的模型（#13644）

- **其他**  
  - Ubuntu / Codium + Zen 订阅下开源模型无响应（#29605）

> 以上日报基于 GitHub 上 `anomalyco/opencode` 仓库截至 2026-06-30 的数据汇总生成。  
> 所有链接均可直接跳转对应的 Issue 或 PR 页面。

</details>

<details>
<summary><strong>Pi</strong> — <a href="https://github.com/badlogic/pi-mono">badlogic/pi-mono</a></summary>

好的，这是为您生成的 2026-06-30 Pi 社区动态日报。

---

# Pi 社区动态日报 | 2026-06-30

## 今日速览

今日，Pi 社区的技术讨论热度集中在**流式传输的稳定性**和 **Session 管理**上。多个关于 Agent 卡死、网络错误处理不当的 Bug 得到修复，同时社区对于 **新的 LLM 提供商（如 Scaleway）** 和 **企业级功能（如 HITL 审批）** 的呼声日益高涨。此外，**Redo 操作支持** 和 **管理配置** 等 PR 的合并，显示出 Pi 在提升用户体验和扩展性方面的积极进展。

## 社区热点 Issues

1.  **[Bug] 流式 Markdown 强制滚动至底部 (#5825)**
    - **重要性**: ⭐⭐⭐⭐⭐ (42 评论)
    - **摘要**: 当启用“Shrink on clear”设置时，AI 在流式输出 Markdown 内容时会强制将用户的阅读位置滚动到最底部，严重影响阅读体验。社区反响热烈，是该周期内最受关注的 Bug。
    - **链接**: [earendil-works/pi Issue #5825](https://earendil-works/pi/issues/5825)

2.  **[Bug] Session 文件夹冲突 (#4877)**
    - **重要性**: ⭐⭐⭐⭐ (20 评论)
    - **摘要**: 由于文件夹路径分隔符的编码方式，不同路径的 Session 可能被存储到同一个文件夹中，例如 `/a/b/c/d` 和 `/a-b/c-d`。虽然不影响核心功能，但存在潜在的命名冲突风险。
    - **链接**: [earendil-works/pi Issue #4877](https://earendil-works/pi/issues/4877)

3.  **[Bug] Agent 被卡住，声称“working”但无进展 (#4338)**
    - **重要性**: ⭐⭐⭐⭐ (6 评论)
    - **摘要**: 用户反馈 Agent 在执行任务时会陷入无限循环，只显示“working”状态，不输出任何内容或进展，需要手动重启 Session。这是一个影响开发效率的核心痛点。
    - **链接**: [earendil-works/pi Issue #4338](https://earendil-works/pi/issues/4338)

4.  **[Bug] OpenAI 流式重试错误未被正确处理 (#6019)**
    - **重要性**: ⭐⭐⭐⭐ (5 评论)
    - **摘要**: 当 OpenAI 在流式响应中途返回一个可重试的错误时（错误信息明确写着“可以重试”），Pi 没有自动进行重试，而是直接以错误终止了本次请求，导致资源浪费和体验不佳。
    - **链接**: [earendil-works/pi Issue #6019](https://earendil-works/pi/issues/6019)

5.  **[Bug] 提供商吞没 HTTP 错误体，导致网关/非 Schema 错误无法阅读 (#5763)**
    - **重要性**: ⭐⭐⭐ (5 评论)
    - **摘要**: 通过代理/网关访问 API 时，如果返回非 2xx 错误，Provider 会将详细的错误信息吞没，导致用户看到的错误信息毫无意义（如 `Unknown: UnknownError`）。这对调试网络或配置问题非常不利。
    - **链接**: [earendil-works/pi Issue #5763](https://earendil-works/pi/issues/5763)

6.  **[功能提案] 持久的“人在回路”（HITL）工具调用中断 (#5901)**
    - **重要性**: ⭐⭐⭐ (4 评论)
    - **摘要**: 开发者提议为 Pi SDK 增加类似 LangGraph 的持久化 HITL 功能，允许在 Headless 集成中，在执行敏感工具调用（如文件写入、部署）前插入人工审批环节。这代表了企业级需求。
    - **链接**: [earendil-works/pi Issue #5901](https://earendil-works/pi/issues/5901)

7.  **[Bug] 小米 MiMo 原生模型定价错误 (#6138)**
    - **重要性**: ⭐⭐⭐ (4 评论)
    - **摘要**: 社区发现 Pi 硬编码的小米 MiMo 模型（如 `mimo-v2.5-pro`）价格与官方定价不符，可能导致用户在使用该模型时预算超支。该问题已修复并关闭。
    - **链接**: [earendil-works/pi Issue #6138](https://earendil-works/pi/issues/6138)

8.  **[Bug] OpenAI Responses API 错误地将空工具结果标注为“(see attached image)” (#6103)**
    - **重要性**: ⭐⭐⭐ (3 评论)
    - **摘要**: 当扩展工具执行成功但返回空结果时，Agent 错误地将结果显示为“（请参阅附图）”，这会导致 Agent 逻辑混乱和 LLM 的错误理解。该 Bug 隐藏在核心逻辑中，但被特定扩展暴露出来。
    - **链接**: [earendil-works/pi Issue #6103](https://earendil-works/pi/issues/6103)

9.  **[Bug] 自动压缩在最终轮次后报错 (#5463)**
    - **重要性**: ⭐⭐⭐ (3 评论, 5 👍)
    - **摘要**: 在 Coding Agent 中，当一次对话的最后一个消息是 Assistant 消息时，自动上下文压缩（Auto-Compaction）会抛出一个未被处理的错误 `Cannot continue from message role: assistant`。这可能会导致一些流程异常结束。
    - **链接**: [earendil-works/pi Issue #5463](https://earendil-works/pi/issues/5463)

10. **[功能提案] 支持 `image_url` 内容类型 (#6151)**
    - **重要性**: ⭐⭐⭐ (2 评论)
    - **摘要**: 目前 Pi 将所有图像都转换为 base64 数据 URI 发送。社区提议支持直接传递 URL 的 `image_url` 类型，这可以显著减少 Token 消耗并提升图片发送效率，尤其是在处理大图或重复图片时。
    - **链接**: [earendil-works/pi Issue #6151](https://earendil-works/pi/issues/6151)

## 重要 PR 进展

1.  **feat(tui): 为编辑器添加重做 (Redo) 支持 (#6182)**
    - **内容**: 实现了编辑器中的 Redo 操作，补全了先前关于撤销 (Undo) 功能讨论的最后一块拼图。用户现在可以在 TUI 编辑器中执行标准的重做操作。
    - **链接**: [earendil-works/pi PR #6182](https://earendil-works/pi/pull/6182)

2.  **fix(tui): 禁用助手消息的填充 (Padding) (#6169)**
    - **内容**: 为 TUI 添加了禁用助手消息之间填充的配置，社区关于“去除填充”的呼声得到了响应。这改善了界面紧凑性和信息密度。
    - **链接**: [earendil-works/pi PR #6169](https://earendil-works/pi/pull/6169)

3.  **fix(coding-agent): 向扩展发出 Session 名称更改事件 (#6175)**
    - **内容**: 确保当 Session 名称被重命名时，这个状态变化能正确通知给所有已注册的扩展，解决了扩展状态与 Session 状态不一致的问题。
    - **链接**: [earendil-works/pi PR #6175](https://earendil-works/pi/pull/6175)

4.  **fix: 修复工具结果消息中未定义内容的防护 (#6178)**
    - **内容**: 当扩展工具（如 `get_kline`）返回结果为 `undefined` 时，此修复避免了因内容未定义而可能引发的运行时错误，增强了系统健壮性。
    - **链接**: [earendil-works/pi PR #6178](https://earendil-works/pi/pull/6178)

5.  **fix(ai): 暴露提供商 HTTP 错误体而非不透明的 SDK 消息 (#5832)**
    - **内容**: 针对 #5763 的修复，此 PR 改进了错误处理逻辑，现在能够将上游网关/代理返回的原始错误信息传递给用户，而不是显示无意义的 SDK 内部错误。
    - **链接**: [earendil-works/pi PR #5832](https://earendil-works/pi/pull/5832)

6.  **在同一个运行中，于下次提供商请求前应用扩展工具变更 (#6176)**
    - **内容**: 修复了 Session 运行时状态刷新问题。当扩展工具在执行过程中修改了激活的工具列表（例如调用 `pi.setActiveTools()`），下一次 Provider 请求将能立刻使用更新后的工具集。
    - **链接**: [earendil-works/pi PR #6176](https://earendil-works/pi/pull/6176)

7.  **避免重放历史内联图像 (#6170)**
    - **内容**: 优化了历史 Session 重建时的上下文处理，不再重放内联的终端图像转义负载，而是使用轻量级的 `[Image: ...]` 标签替代，从而节省 Token 并提升性能。
    - **链接**: [earendil-works/pi PR #6170](https://earendil-works/pi/pull/6170)

8.  **fix(ai): 从挂起的流中恢复并重试未建模的 Bedrock 错误 (#6051)**
    - **内容**: 大幅提升了与 Bedrock 和 Anthropic 的连接稳定性，通过引入空闲超时和连接超时机制，可重试并恢复半开连接上的挂起流，解决了“卡死”问题。
    - **链接**: [earendil-works/pi PR #6051](https://earendil-works/pi/pull/6051)

9.  **fix(tui): 稳定工作状态行 (#6026)**
    - **内容**: 旨在修复 #5825 相关问题，通过稳定 TUI 中的工作状态行显示，可能包含了修复滚动问题所需的基础逻辑。
    - **链接**: [earendil-works/pi PR #6026](https://earendil-works/pi/pull/6026)

10. **fix(ai): 将 Bedrock apiKey 认证映射到 Bearer Token 环境变量 (#6161)**
    - **内容**: 修复了亚马逊 Bedrock 提供商的认证问题，确保 `apiKey` 能够被正确映射并传递给 Bedrock Converse API 的 Bearer Token 环境中。
    - **链接**: [earendil-works/pi PR #6161](https://earendil-works/pi/pull/6161)

## 功能需求趋势

- **稳定性与错误处理优化**: 社区最强烈的需求是让 Agent 跑得更稳。这包括修复 Agent 卡死（#4338）、提高流式中断的恢复能力（#6019, #6051）、以及改进 Provider 连接和错误信息（#5763, #6133）。用户对“Agent 无响应”和“模棱两可的错误提示”的容忍度极低。
- **新模型与提供商支持**: 对更多模型的支持是持续的需求。本周有提案要求支持**Scaleway**（#6165）和**小米 MiMo**（#6138），同时也有关于Azure OpenAI新模型命名错误的报告（#6114）。这表明社区希望 Pi 能紧跟最新 AI 生态。
- **用户体验与编辑功能增强**: 除了核心的 Agent 能力，用户对“编辑器”本身也提出了更高要求。**Redo功能**（#6183）和**多行管道输入到扩展命令**（#6172）的提案表明，开发者社区希望 Pi 的 TUI 成为一个更强大的开发工具，达到类似 IDE 的操作标准。
- **可配置性与管理能力**: 用户不再满足于开箱即用。**Profile 支持**（#3966）用于隔离不同项目状态，以及**企业级管理配置**（#6159）的提出，标志着 Pi 正从个人工具向更复杂的团队协作环境演进。

## 开发者关注点

- **Session 冲突与状态管理**: 开发者注意到 Session 存储的潜在冲突（#4877）和 Agent 运行中状态同步的 Bug（#6176），这表明内部 Session 管理机制的鲁棒性仍有提升空间。
- **流式传输与连接中断**: 多个 Issue 指向了流式传输场景下的问题，包括不合理的滚动（#5825）、无法重试的错误（#6019）和网络中断导致的崩溃（#6133）。**开发者在网络不稳定或使用代理时体验尤为不佳**。
- **模型集成细节**: 模型集成不仅仅是“能用”，更需要“准确”。错误的定价（#6138）、错误的模型名（#6114）和错误的图片编码（#6164）都暴露出模型集成时细节验证的重要性。**开发者期望 Pi 能自动验证或同步官方数据**。
- **对“无填充”和自定义行为的强烈渴望**: 围绕“去除填充”（#6169）的讨论持续不断，开发者希望获得对 TUI 界面的更多控制权，拥有类似 **“无干扰”模式**的选项。
- **“人机交互”需求的抬头**: 尽管 HITL 提案（#5901）的评论数不多，但它代表了高级用户和团队用户对**控制权和安全性**的追求，希望 Agent 不是完全自主，而是可以被安全地“审核和暂停”。

---

</details>

<details>
<summary><strong>Qwen Code</strong> — <a href="https://github.com/QwenLM/qwen-code">QwenLM/qwen-code</a></summary>

# 2026-06-30 Qwen Code 社区动态日报

## 今日速览

今日发布 v0.19.3-nightly 版本，主要围绕 daemon 文档刷新和核心 auto‑ 功能配置。社区聚焦于会话管理（存档、折叠预览）、消息通道（worker、例行任务）以及安全与性能优化。多个高价值 PR 进入活跃开发阶段，值得关注。

## 版本发布

**v0.19.3-nightly.20260630.e00fe6a27**  
- 更新 daemon 相关文档（wave 2）  
- 添加可配置的 auto‑ 功能（实现细节待补全）  
🔗 [Release 页面](https://github.com/QwenLM/qwen-code/releases/tag/v0.19.3-nightly.20260630.e00fe6a27)

## 社区热点 Issues

由于过去 24 小时内仅更新了 5 条 Issue，以下全部列出：

1. **#6049 – generationConfig.timeout = 0 导致请求立即超时**  
   影响：用户将超时设为 0，期望不限制等待时间，结果反而立即超时。社区确认这是错误行为，期待修复。  
   🔗 [Issue](https://github.com/QwenLM/qwen-code/issues/6049)

2. **#5759 – 请求添加 `ui.history.collapsePreviewCount` 设置**  
   功能：当启用会话折叠时，允许在恢复时显示最后 N 条消息，避免完全隐藏历史。已关闭但设计值得参考。  
   🔗 [Issue](https://github.com/QwenLM/qwen-code/issues/5759)

3. **#6024 – 复制代码块时不应包含行号**  
   功能：终端 UI 中复制代码会附带左侧行号，影响粘贴体验。社区欢迎贡献者实现修复。  
   🔗 [Issue](https://github.com/QwenLM/qwen-code/issues/6024)

4. **#6057 – 添加 daemon 会话存档支持**  
   功能：允许将会话移出活动 `chats/` 目录而无需删除，归档后可恢复但不可加载/恢复。解决用户“轻量隐藏或保留旧会话”的需求。  
   🔗 [Issue](https://github.com/QwenLM/qwen-code/issues/6057)

5. **#4940 – 添加 `deniedMcpServers` 策略**  
   功能：补充现有的 `allowedMcpServers` 白名单，增加显式禁止的 MCP 服务器列表，提升安全性。已关闭，但思路值得后续版本参考。  
   🔗 [Issue](https://github.com/QwenLM/qwen-code/issues/4940)

## 重要 PR 进展

从 50 个 PR 中精选以下 10 个：

1. **#6018 – 避免在 OOM 高风险路径中克隆全量历史**  
   优化：API 错误报告改为发送紧凑摘要，forked‑agent 缓存快照避免深克隆，降低内存压力。  
   🔗 [PR](https://github.com/QwenLM/qwen-code/pull/6018)

2. **#6021 – 修复 ACP read_file 对托管本地路径的处理**  
   修复：使 ACP 后端能够读取技能指令、临时输出、子代理记录等本地托管路径，绕过工作区限制。  
   🔗 [PR](https://github.com/QwenLM/qwen-code/pull/6021)

3. **#5902 – QQ Bot 流式改进：空闲刷新、移除分片、消息 TTL、Markdown 管道**  
   功能：重写 QQ 机器人流式行为，增加 2 秒空闲刷新、移除 2000 字符限制、被动回复增加 5 分钟 TTL、修复 Markdown 表格。  
   🔗 [PR](https://github.com/QwenLM/qwen-code/pull/5902)

4. **#5884 – 添加会话无关的工作区记忆功能**  
   功能：daemon 新增 `workspace remember` API，允许在不创建用户可见会话的情况下存入托管记忆，用于无会话上下文存储。  
   🔗 [PR](https://github.com/QwenLM/qwen-code/pull/5884)

5. **#6027 – 净化子代理结果标签**  
   安全：子代理最终结果在返回父代理前移除 `<analysis>` 等内部标签，仅保留模型可见的干净摘要。  
   🔗 [PR](https://github.com/QwenLM/qwen-code/pull/6027)

6. **#6022 – 支持单次模型覆盖命令 `/model <id> <prompt>`**  
   功能：允许一次性使用其他模型执行提示，之后自动恢复会话原有模型，方便快速切换测试。  
   🔗 [PR](https://github.com/QwenLM/qwen-code/pull/6022)

7. **#6031 – daemon 管理的频道工作进程**  
   功能：`qwen serve` 支持 `--channel <name>` 或 `--channel all`，启动一个独立的频道工作进程，连接回当前 serve 实例。  
   🔗 [PR](https://github.com/QwenLM/qwen-code/pull/6031)

8. **#5999 – 将 TUI 中所有 emoji 替换为 Unicode 文字符号**  
   UI：完全替换终端 UI 中宽度不一的 emoji（如 💡→∴, ✅→✓），提升跨终端兼容性。  
   🔗 [PR](https://github.com/QwenLM/qwen-code/pull/5999)

9. **#6040 – 添加 `autoMode.classifyAllShell` 设置**  
   功能：启用后所有 shell 命令（包括只读命令如 `ls`）都经过自动模式的 LLM 分类器再执行，增加安全控制粒度。  
   🔗 [PR](https://github.com/QwenLM/qwen-code/pull/6040)

10. **#6032 – 支持 HTTPS/TLS（ `--tls-cert` / `--tls-key` ）**  
    安全：为 `qwen serve` 添加 PEM 证书和私钥参数，使 daemon 可运行 HTTPS 而非纯 HTTP。  
    🔗 [PR](https://github.com/QwenLM/qwen-code/pull/6032)

## 功能需求趋势

综合今日 Issues 与 PR，社区最关注的方向包括：

- **会话管理**：存档（#6057）、折叠预览（#5759）、无会话记忆（#5884）表明用户需要更灵活的会话生命周期控制。
- **消息渠道增强**：频道工作进程（#6031）、频道例行任务（#6038、#6054）、频道记忆（#6051）显示开发者正大力扩展多平台集成。
- **UI/UX 改进**：emoji 替换（#5999）、代码块复制去除行号（#6024）、移动端侧边栏（#6003）、标签化设置（#6044）持续打磨终端和 Web 体验。
- **安全与策略**：deniedMcpServers（#4940）、classifyAllShell（#6040）、TLS 支持（#6032）、子代理结果净化（#6027）反映出对安全边界的重视。
- **性能与稳定性**：避免全量克隆（#6018）、CI 稳定性（#6056）、serve health 路径优化（#6013）持续保障运行可靠性。

## 开发者关注点

从 Issue 评论和 PR 描述中提炼出的高频反馈：

- **配置行为意外**：`timeout = 0` 立即超时（#6049）并非直观行为，开发者期望 `0` 能表示“不限时”或“由服务端决定”。
- **终端复制体验差**：复制代码块携带行号严重影响日常使用（#6024），该议题获得 welcome-pr 标签，欢迎社区贡献。
- **会话管理粒度不足**：用户希望在不删除历史的情况下“存档”会话（#6057），以及恢复折叠会话时能预览最后几条消息（#5759），表明当前全部隐藏或全部显示的策略不够灵活。
- **安全策略扩展**：允许拒绝特定 MCP 服务器（#4940）是许多企业用户的安全刚需，社区希望该功能能尽快进入主版本。
- **移动端支持**：多个 PR 涉及移动端布局（#6003）和触控优化，开发者对跨设备使用 Qwen Code 的需求日益强烈。

---  
*数据截止 2026-06-30 23:59 UTC，所有链接均指向 GitHub 已公开页面。*

</details>

<details>
<summary><strong>DeepSeek TUI</strong> — <a href="https://github.com/Hmbown/DeepSeek-TUI">Hmbown/DeepSeek-TUI</a></summary>

好的，以下是为您生成的 **2026年6月30日 DeepSeek TUI (CodeWhale) 社区动态日报**。

---

# DeepSeek TUI (CodeWhale) 社区动态日报 | 2026-06-30

## 📰 今日速览

今日社区焦点集中于 **v0.8.66 版本发布前的最后冲刺**。项目核心维护者合并了多达 7 个关于子代理（Sub-agent）高并发下 TUI 冻结的修复 PR，并调整了 Hotbar 功能的默认隐藏策略。此外，社区报告了两个严重影响体验的 Bug：会话因长时间工具输出或认证超时而永久损坏，以及 MCP OAuth 认证流程的 UX 问题。

## 🔖 版本发布

*暂无新版本发布。社区正处于 v0.8.66 候选版本的密集测试与修复阶段。*

## 🔥 社区热点 Issues

| 序号 | Issue 标题 | 更新时间 | 热度 | 关键信息 |
| :--- | :--- | :--- | :--- | :--- |
| 1 | **输入缓存命中率太低了** [#1177](https://github.com/Hmbown/CodeWhale/issues/1177) | 2026-06-29 | 💬 **24** | **社区最高频痛点**。用户对比发现缓存命中率远低于竞品（如 DeepSeek-Reasonix 达 95%+），强烈要求优化，已连续多周保持高热度。 |
| 2 | **缓存命中方面似乎还是有些问题** [#1120](https://github.com/Hmbown/CodeWhale/issues/1120) | 2026-06-29 | 💬 **21** | 与 #1177 同属缓存问题，用户报告即使在更新版本后，对同一项目的修改仍存在缓存未命中，怀疑 Bug 并未完全修复。 |
| 3 | **Token消耗增大了很多** [#743](https://github.com/Hmbown/CodeWhale/issues/743) | 2026-06-29 | 💬 **13** | **性能与成本担忧**。用户报告半天消耗 4 亿 Token，请求过于密集，强烈建议优化对话交互信息，减少不必要的请求。 |
| 4 | **v0.8.65: MCP重复服务实例生命周期与doctor覆盖** [#3461](https://github.com/Hmbown/CodeWhale/issues/3461) | 2026-06-30 | 💬 **9** | **可靠性Bug（已关闭）**。发现单个 `mcp.json` 配置会启动两个 MCP 服务器进程，其中一个为孤儿进程，浪费内存并可能导致管道问题。 |
| 5 | **会话永久损坏 — 长工具输出 / 审批超时导致 Agent 停止响应** [#3821](https://github.com/Hmbown/CodeWhale/issues/3821) | 2026-06-30 | 💬 **1** | **严重性Bug**。今天新提交，用户反馈在长时间工具输出或审批对话框超时后，Agent 会话进入永久损坏状态，必须重启应用。 |
| 6 | **MCP OAuth认证 UX问题：token未自动刷新、静默错误、前台登录超时** [#3819](https://github.com/Hmbown/CodeWhale/issues/3819) | 2026-06-30 | 💬 **1** | **UX缺陷**。配置 OAuth 保护的 MCP 服务器后，token 无法自动刷新，错误提示不明确，前台登录流程容易超时，严重影响配置体验。 |
| 7 | **v0.9.0 HarnessPosture: 模型特定的上下文与子代理策略** [#2693](https://github.com/Hmbown/CodeWhale/issues/2693) | 2026-06-30 | 💬 **3** | **远期规划 (v0.9.0)**。维护者提出的未来方向，计划为不同模型（如 DeepSeek V4、小米MiMo）设置不同的上下文策略，以利用缓存和前缀稳定性。 |
| 8 | **v0.8.66: 多子 Agent 扇出冻结的发布门禁** [#3800](https://github.com/Hmbown/CodeWhale/issues/3800) | 2026-06-30 | 💬 **2** | **当前版本核心问题 (已关闭)**。当子 Agent 数量较多（约20个）时，TUI 会表现冻结。今日有多个 PR 针对此问题进行了修复。 |
| 9 | **执行大文本处理工程后会话中断卡死** [#1425](https://github.com/Hmbown/CodeWhale/issues/1425) | 2026-06-29 | 💬 **1** | **稳定性问题**。用户在尝试让 AI 分析 300 万字小说时，启动 10 个子 Agent 后因等待超时而导致会话卡死，反映了大规模任务下的处理瓶颈。 |
| 10 | **Agent模式：为工具调用失败添加回退策略** [#1641](https://github.com/Hmbown/CodeWhale/issues/1641) | 2026-06-29 | 💬 **3** | **功能需求**。用户希望 Agent 在执行依赖外部服务的任务时，如果工具调用因反爬或超时而失败，能够自动切换到替代方案或优雅降级，而非一直重试直到失败。 |

## 🚀 重要 PR 进展

| 序号 | PR 标题 | 更新时间 | 类型 | 核心内容 |
| :--- | :--- | :--- | :--- | :--- |
| 1 | **feat(mcp): 在MCP stdio配置中扩展${VAR}环境变量占位符** [#3825](https://github.com/Hmbown/CodeWhale/pull/3825) | 2026-06-30 | 功能 | **外部贡献**。允许在 `mcp.json` 中使用 `${VAR}` 引用外部环境变量（如 API Key），解决了因子进程环境变量限制而无法传递密钥的问题。 |
| 2 | **fix(engine): 支持通配符禁止工具前缀** [#3824](https://github.com/Hmbown/CodeWhale/pull/3824) | 2026-06-30 | 修复 | **外部贡献**。`disallowed_tools` 配置现在支持通配符匹配，用户可以更方便地禁用来自某个 MCP 服务器的所有工具。 |
| 3 | **fix: 在Windows上抑制后台控制台窗口弹出** [#3823](https://github.com/Hmbown/CodeWhale/pull/3823) | 2026-06-30 | 修复 | **外部贡献**。修复了 Windows 下频繁弹出控制台窗口的问题，这些窗口会遮挡 UI 并抢夺焦点，改善 Windows 用户体验。 |
| 4 | **fix(subagent): 在Manager写锁热路径之外持久化状态** [#3816](https://github.com/Hmbown/CodeWhale/pull/3816) | 2026-06-30 | 修复 | **性能优化**。将子代理状态的持久化（JSON序列化/写盘）操作移出写锁范围，显著减少高扇出场景下的并发争用，缓解 #3800 描述的性能问题。 |
| 5 | **fix(tui): 使用非阻塞发送来处理ListSubAgents刷新事件** [#3813](https://github.com/Hmbown/CodeWhale/pull/3813) | 2026-06-30 | 修复 | **性能优化**。将子代理列表刷新事件使用非阻塞发送，避免因事件通道充满而导致的引擎或UI事件循环阻塞，缓解 #3800 描述的性能问题。 |
| 6 | **fix(tui): 允许子Agent启动加入并行分发批次** [#3812](https://github.com/Hmbown/CodeWhale/pull/3812) | 2026-06-30 | 修复 | **性能优化**。允许 `agent` 工具的多个调用并行分发，解决了之前串行启动导致的高延迟和 TUI 冻结感。 |
| 7 | **fix(tui): 从只读快照渲染子Agent侧边栏** [#3809](https://github.com/Hmbown/CodeWhale/pull/3809) | 2026-06-30 | 修复 | **性能优化**。UI 刷新时不再对 `SubAgentManager` 上写锁，而是从只读快照渲染，避免与并发更新和持久化操作产生锁争用。 |
| 8 | **fix(tui): 在异步UI刷新路径中使用try_lock操作shell管理器** [#3808](https://github.com/Hmbown/CodeWhale/pull/3808) | 2026-06-30 | 修复 | **性能优化**。将UI刷新路径中对shell管理器的阻塞锁改为`try_lock`，防止后台shell任务长时间持有锁时阻塞UI线程。 |
| 9 | **fix(tui): 保持审批控件内联可见** [#3814](https://github.com/Hmbown/CodeWhale/pull/3814) | 2026-06-30 | 修复 | **UI修复**。修复了长审批提示导致交互按钮超出屏幕的问题，现在审批控件始终可见，改善使用体验。 |
| 10 | **fix(tui): 保留运行时延续的 YOLO 模式权限** [#3817](https://github.com/Hmbown/CodeWhale/pull/3817) | 2026-06-30 | 修复 | **Bug修复**。修复了在 YOLO 模式下，因运行时延续或子 Agent 交接导致的某些操作（如 git push）仍会弹出审批提示的问题。 |

## 💡 功能需求趋势

根据近期 Issue 分析，社区最关注的功能方向集中在：
1.  **性能与成本优化**：缓存命中率、Token消耗是压倒性的关注焦点，直接关系到用户的使用成本和响应速度。
2.  **UI/UX 与交互优化**：包括 Hotbar 自定义、TUI 布局溢出、审批流程优化、鼠标滚轮支持等，体现了对终端用户友好度的追求。
3.  **可靠性与稳定性**：子 Agent 高并发下的 TUI 僵死、会话永久损坏、大规模文本处理卡死等问题是当前开发的重点攻坚方向。
4.  **MCP 生态集成改进**：包括 OAuth 认证流程优化、环境变量注入、工具管理和错误处理等，MCP 正成为扩展功能的核心渠道。
5.  **文档与本地化**：随着功能增多，社区对更清晰、更全面的文档（尤其是针对新功能如 Constitution、Hotbar）以及非英语本地化的需求正在涌现。

## 🧑‍💻 开发者关注点

- **高扇出场景下的性能瓶颈**：当 Agent 同时启动多个子 Agent 时，TUI 的卡顿、冻结、锁竞争是开发者反馈的最主要痛点。今日大量的修复 PR 都与此相关。
- **缓存命中率远低于预期**：用户强烈对比并反馈，当前项目的缓存机制效率远低于同类竞品，这是导致高 Token 消耗和低响应速度的核心原因之一，亟待根本性优化。
- **会话不可恢复性崩溃**：新提交的关于“会话永久损坏”的话题 (#3821) 属于严重性 Bug，开发者应高度关注其根因分析和修复进度。
- **MCP 认证体验痛苦**：随着 MCP 服务接入增多，OAuth 认证流程的失败和糟糕的用户体验成为新用户配置的主要障碍。
- **旧版 Hotbar 的默认可见性**：社区对 Hotbar 功能的默认隐藏策略存在讨论，开发者需注意在稳定性和新功能推广之间找到平衡。
- **跨平台问题**：Windows 下控制台窗口闪烁的问题 (#3823) 反映了桌面工具跨平台适配中需要持续投入的细节。

</details>

---
*本日报由 [agents-radar](https://github.com/ivo-eu/agents-radar) 自动生成。*